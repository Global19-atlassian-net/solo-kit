package swagger

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/go-multierror"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/log"

	"github.com/pkg/errors"
	discovery "github.com/solo-io/solo-kit/projects/discovery/pkg"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins"
	rest_plugins "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/rest"
	transformation_plugins "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/transformation"
)

var commonSwaggerURIs = []string{
	"/swagger.json",
	"/swagger/docs/v1",
	"/swagger/docs/v2",
	"/v1/swagger",
	"/v2/swagger",
}

// TODO(yuval-k): run this in a back off for a limited amount of time, with high initial retry.
// maybe backoff with initial 1 minute a total of 10 minutes till giving up. this should probably be configurable

type SwaggerFuncitonDiscoveryFactory struct {
	DetectionTimeout   time.Duration
	DetectionRetryBase time.Duration
	FunctionPollTime   time.Duration
	swaggerUrisToTry   []string
}

func (f *SwaggerFuncitonDiscoveryFactory) NewFunctionDiscovery(u *v1.Upstream) discovery.UpstreamFunctionDiscovery {
	return &SwaggerFuncitonDiscovery{
		detectionTimeout:   f.DetectionTimeout,
		detectionRetryBase: f.DetectionRetryBase,
		functionPollTime:   f.FunctionPollTime,
		swaggerUrisToTry:   f.swaggerUrisToTry,
		upstream:           u,
	}
}

type SwaggerFuncitonDiscovery struct {
	detectionTimeout   time.Duration
	detectionRetryBase time.Duration
	functionPollTime   time.Duration
	upstream           *v1.Upstream
	swaggerUrisToTry   []string
}

type specable interface {
	GetServiceSpec() *plugins.ServiceSpec
}
type setspecable interface {
	specable
	SetServiceSpec(*plugins.ServiceSpec)
}

func getswagspec(u *v1.Upstream) *rest_plugins.ServiceSpec_SwaggerInfo {
	spec, ok := u.UpstreamSpec.UpstreamType.(specable)
	if !ok {
		return nil
	}
	restwrapper, ok := spec.GetServiceSpec().PluginType.(*plugins.ServiceSpec_Rest)
	if !ok {
		return nil
	}
	rest := restwrapper.Rest
	return rest.SwaggerInfo
}

func (d *SwaggerFuncitonDiscovery) IsFunctional() bool {
	return getswagspec(d.upstream) != nil
}

func (d *SwaggerFuncitonDiscovery) DetectType(ctx context.Context, baseurl *url.URL) (*plugins.ServiceSpec, error) {
	var spec *plugins.ServiceSpec

	err := contextutils.NewExponentioalBackoff(contextutils.ExponentioalBackoff{MaxDuration: &d.detectionTimeout}).Backoff(ctx, func(ctx context.Context) error {
		var err error
		spec, err = d.detectUpstreamTypeOnce(ctx, baseurl)
		return err
	})

	return spec, err
}

func (d *SwaggerFuncitonDiscovery) detectUpstreamTypeOnce(ctx context.Context, baseurl *url.URL) (*plugins.ServiceSpec, error) {
	// run detection and get functions
	var errs error
	log := contextutils.LoggerFrom(ctx)

	log.Debugf("attempting to detect swagger base url %v", baseurl)

	switch baseurl.Scheme {
	case "http":
		fallthrough
	case "https":
		// nothing to do as this baseurl already has an http address.
	case "tcp":
		// if it is a tcp address, assume it is plain http
		baseurl.Scheme = "http"
	default:
		return nil, fmt.Errorf("unsupported baseurl for swagger discovery %v", baseurl)
	}

	for _, uri := range d.swaggerUrisToTry {
		url := baseurl.ResolveReference(&url.URL{Path: uri}).String()
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "invalid url for request")
		}
		req.Header.Set("X-Gloo-Discovery", "Swagger-Discovery")

		req = req.WithContext(ctx)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			errs = multierror.Append(errs, errors.Wrapf(err, "could not perform HTTP GET on resolved addr: %v", url))
			continue
		}
		// might have found a swagger service
		if res.StatusCode == http.StatusOK {
			if _, err := RetrieveSwaggerDocFromUrl(ctx, url); err != nil {
				// first check if this is a context error
				if ctx.Err() != nil {
					return nil, ctx.Err()
				}
				errs = multierror.Append(errs, err)
				continue
			}
			// definitely found swagger
			log.Infof("swagger upstream detected: %v", url)
			svcInfo := &plugins.ServiceSpec{
				PluginType: &plugins.ServiceSpec_Rest{
					Rest: &rest_plugins.ServiceSpec{
						SwaggerInfo: &rest_plugins.ServiceSpec_SwaggerInfo{
							SwaggerSpec: &rest_plugins.ServiceSpec_SwaggerInfo_Url{
								Url: url,
							},
						},
					},
				},
			}
			return svcInfo, nil
		}
		errs = multierror.Append(errs, errors.Errorf("path: %v response code: %v headers: %v", uri, res.Status, res.Header))

	}
	log.Infof("failed to detect swagger for %s: %v", baseurl.String(), errs.Error())
	// not a swagger upstream
	return nil, errors.Wrapf(errs, "service at %s does not implement swagger at a known endpoint, "+
		"or was unreachable", baseurl.String())

}

func (f *SwaggerFuncitonDiscovery) DetectFunctions(ctx context.Context, secrets func() v1.SecretList, updatecb func(discovery.UpstreamMutator) error) error {
	in := f.upstream
	spec := getswagspec(in)
	if spec == nil || spec.SwaggerSpec == nil {
		// TODO: make this a fatal error that avoids restarts?
		return errors.New("upstream doesn't have a swagger spec")
	}
	switch document := spec.SwaggerSpec.(type) {
	case *rest_plugins.ServiceSpec_SwaggerInfo_Url:
		return f.detectFunctionsFromUrl(ctx, document.Url, in, updatecb)
	case *rest_plugins.ServiceSpec_SwaggerInfo_Inline:
		return f.detectFunctionsFromInline(ctx, document.Inline, in, updatecb)
	}

	return errors.New("upstream doesn't have a swagger source")
}

func (f *SwaggerFuncitonDiscovery) detectFunctionsFromUrl(ctx context.Context, url string, in *v1.Upstream, updatecb func(discovery.UpstreamMutator) error) error {
	for {

		err := contextutils.NewExponentioalBackoff(contextutils.ExponentioalBackoff{}).Backoff(ctx, func(ctx context.Context) error {

			spec, err := RetrieveSwaggerDocFromUrl(ctx, url)
			if err != nil {
				return err
			}
			err = f.detectFunctionsFromSpec(ctx, spec, in, updatecb)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			// ignore other erros as we would like to continue forever.
		}

		if err := contextutils.Sleep(ctx, f.functionPollTime); err != nil {
			return err
		}
	}

}

func (f *SwaggerFuncitonDiscovery) detectFunctionsFromInline(ctx context.Context, document string, in *v1.Upstream, updatecb func(discovery.UpstreamMutator) error) error {
	spec, err := parseSwaggerDoc([]byte(document))
	if err != nil {
		return err
	}
	return f.detectFunctionsFromSpec(ctx, spec, in, updatecb)
}

func (f *SwaggerFuncitonDiscovery) detectFunctionsFromSpec(ctx context.Context, swaggerSpec *spec.Swagger, in *v1.Upstream, updatecb func(discovery.UpstreamMutator) error) error {
	var consumesJson bool
	if len(swaggerSpec.Consumes) == 0 {
		consumesJson = true
	}
	for _, contentType := range swaggerSpec.Consumes {
		if contentType == "application/json" {
			consumesJson = true
			break
		}
	}
	if !consumesJson {
		return errors.Errorf("swagger function discovery uses content type application/json; "+
			"available: %v", swaggerSpec.Consumes)
	}
	// TODO: when response transformation is done, look at produces as well

	funcs := make(map[string]*transformation_plugins.TransformationTemplate)

	for functionPath, pathItem := range swaggerSpec.Paths.Paths {
		createFunctionsForPath(funcs, swaggerSpec.BasePath, functionPath, pathItem.PathItemProps, swaggerSpec.Definitions)
	}

	return updatecb(func(u *v1.Upstream) error {
		upstremaspec, ok := u.UpstreamSpec.UpstreamType.(setspecable)
		if !ok {
			return errors.New("not a valid upstream")
		}
		spec := upstremaspec.GetServiceSpec()
		if spec == nil {
			spec = &plugins.ServiceSpec{}
		}
		restspec, ok := spec.PluginType.(*plugins.ServiceSpec_Rest)
		if !ok {
			restspec = &plugins.ServiceSpec_Rest{
				Rest: &rest_plugins.ServiceSpec{},
			}
		}

		restspec.Rest.Transformations = funcs
		spec.PluginType = restspec

		upstremaspec.SetServiceSpec(spec)
		return nil
	})
}

func RetrieveSwaggerDocFromUrl(ctx context.Context, url string) (*spec.Swagger, error) {
	docBytes, err := LoadFromFileOrHTTP(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "loading swagger doc from url")
	}
	return parseSwaggerDoc(docBytes)
}

func LoadFromFileOrHTTP(ctx context.Context, url string) ([]byte, error) {
	return swag.LoadStrategy(url, ioutil.ReadFile, loadHTTPBytes(ctx))(url)
}

func loadHTTPBytes(ctx context.Context) func(path string) ([]byte, error) {
	return func(path string) ([]byte, error) {
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			return nil, err
		}
		req = req.WithContext(ctx)
		resp, err := http.DefaultClient.Do(req)
		defer func() {
			if resp != nil {
				if e := resp.Body.Close(); e != nil {
					contextutils.LoggerFrom(ctx).Debug(e)
				}
			}
		}()
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("could not access document at %q [%s] ", path, resp.Status)
		}

		return ioutil.ReadAll(resp.Body)
	}
}

func parseSwaggerDoc(docBytes []byte) (*spec.Swagger, error) {
	doc, err := loads.Analyzed(docBytes, "")
	if err != nil {
		log.Warnf("parsing doc as json failed, falling back to yaml")
		jsn, err := swag.YAMLToJSON(docBytes)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert yaml to json (after falling back to yaml parsing)")
		}
		doc, err = loads.Analyzed(jsn, "")
		if err != nil {
			return nil, errors.Wrap(err, "invalid swagger doc")
		}
	}
	return doc.Spec(), nil
}