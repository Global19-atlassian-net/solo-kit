package translator

import (
	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	defaultv1 "github.com/solo-io/gloo/pkg/api/defaults/v1"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins"
	"github.com/gogo/protobuf/types"
)

type reportErr func(error error, format string, args ...interface{})

func (t *translator) computeRouteConfig(proxy *v1.Proxy, listener *v1.Listener, routeCfgName string, snap *v1.Snapshot, resourceErrs reporter.ResourceErrors) *envoyapi.RouteConfiguration {
	report := func(err error, format string, args ...interface{}) {
		resourceErrs.AddError(proxy, errors.Wrapf(err, format, args...))
	}

	virtualHosts := t.computeVirtualHosts(listener, snap, report)

	// validate ssl config if the listener specifies any
	if err := validateListenerSslConfig(listener, snap.SecretList); err != nil {
		report(err, "invalid listener %v", listener.Name)
	}

	return &envoyapi.RouteConfiguration{
		Name:         routeCfgName,
		VirtualHosts: virtualHosts,
	}
}

func (t *translator) computeVirtualHosts(listener *v1.Listener, snap *v1.Snapshot, report reportErr) []envoyroute.VirtualHost {
	httpListener, ok := listener.ListenerType.(*v1.Listener_HttpListener)
	if !ok {
		panic("non-HTTP listeners are not currently supported in Gloo")
	}
	virtualHosts := httpListener.HttpListener.VirtualHosts
	if err := validateVirtualHostDomains(virtualHosts); err != nil {
		report(err, "invalid listener %v", listener.Name)
	}
	requireTls := len(listener.SslConfiguations) > 0
	var envoyVirtualHosts []envoyroute.VirtualHost
	for _, virtualHost := range virtualHosts {
		envoyVirtualHosts = append(envoyVirtualHosts, t.computeVirtualHost(virtualHost, requireTls, snap, report))
	}
	return envoyVirtualHosts
}

func (t *translator) computeVirtualHost(virtualHost *v1.VirtualHost, requireTls bool, snap *v1.Snapshot, report reportErr) envoyroute.VirtualHost {
	var envoyRoutes []envoyroute.Route
	for _, route := range virtualHost.Routes {
		envoyRoute := t.envoyRoute(snap, report, route)
		envoyRoutes = append(envoyRoutes, envoyRoute)
	}
	domains := virtualHost.Domains
	if len(domains) == 0 || (len(domains) == 1 && domains[0] == "") {
		domains = []string{"*"}
	}
	var envoyRequireTls envoyroute.VirtualHost_TlsRequirementType
	if requireTls {
		// TODO (ilackarms): support external-only TLS
		envoyRequireTls = envoyroute.VirtualHost_ALL
	}

	return envoyroute.VirtualHost{
		Name:       virtualHost.Name,
		Domains:    domains,
		Routes:     envoyRoutes,
		RequireTls: envoyRequireTls,
		//TODO (ilackarms): plugins for these
		//VirtualClusters: nil,
		//RateLimits: nil,
		//RequestHeadersToAdd: nil,
		//ResponseHeadersToRemove: nil,
		//Cors: nil,
		//Auth: nil,
	}
}

func (t *translator) envoyRoute(snap *v1.Snapshot, report reportErr, in *v1.Route) envoyroute.Route {
	out := &envoyroute.Route{}

	setMatch(in, out)

	setAction(snap, report, in, out)

	// run the plugins
	params := plugins.PluginParams{
		Snapshot: snap,
	}
	for _, plug := range t.plugins {
		routePlugin, ok := plug.(plugins.RoutePlugin)
		if !ok {
			continue
		}
		if err := routePlugin.ProcessRoute(params, in, out); err != nil {
			report(err, "invalid route")
		}
	}
	return *out
}

func setMatch(in *v1.Route, out *envoyroute.Route) {
	match := envoyroute.RouteMatch{
		Headers:         envoyHeaderMatcher(in.Matcher.Headers),
		QueryParameters: envoyQueryMatcher(in.Matcher.QueryParameters),
	}
	// need to do this because Go's proto implementation makes oneofs private
	// which genius thought of that?
	setEnvoyPathMatcher(in.Matcher, &match)

	out.Match = match
}

func setAction(snap *v1.Snapshot, report reportErr, in *v1.Route, out *envoyroute.Route) {
	switch action := in.Action.(type) {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//
	// TODO(ilackarms): rethink this block of code
	//
	case *v1.Route_RouteAction:
		if err := validateRouteDestinations(snap.UpstreamList, action.RouteAction); err != nil {
			report(err, "invalid route")
		}
	case *v1.Route_DirectResponseAction:
		out.Action = &envoyroute.Route_DirectResponse{
			DirectResponse: &envoyroute.DirectResponseAction{
				Status: action.DirectResponseAction.Status,
				Body:   DataSourceFromString(action.DirectResponseAction.Body),
			},
		}
	case *v1.Route_RedirectAction:
		out.Action = &envoyroute.Route_Redirect{
			Redirect: &envoyroute.RedirectAction{
				HostRedirect:  action.RedirectAction.HostRedirect,
				ResponseCode:  envoyroute.RedirectAction_RedirectResponseCode(action.RedirectAction.ResponseCode),
				HttpsRedirect: action.RedirectAction.HttpsRedirect,
				StripQuery:    action.RedirectAction.StripQuery,
			},
		}

		switch pathRewrite := action.RedirectAction.PathRewriteSpecifier.(type) {
		case *v1.RedirectAction_PathRedirect:
			out.Action.(*envoyroute.Route_Redirect).Redirect.PathRewriteSpecifier = &envoyroute.RedirectAction_PathRedirect{
				PathRedirect: pathRewrite.PathRedirect,
			}
		case *v1.RedirectAction_PrefixRewrite:
			out.Action.(*envoyroute.Route_Redirect).Redirect.PathRewriteSpecifier = &envoyroute.RedirectAction_PrefixRewrite{
				PrefixRewrite: pathRewrite.PrefixRewrite,
			}
		}
	}
	//
	//
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////

}

func setEnvoyPathMatcher(in *v1.RouteMatcher, out *envoyroute.RouteMatch) {
	switch path := in.PathSpecifier.(type) {
	case *v1.RouteMatcher_Exact:
		out.PathSpecifier = &envoyroute.RouteMatch_Path{
			Path: path.Exact,
		}
	case *v1.RouteMatcher_Regex:
		out.PathSpecifier = &envoyroute.RouteMatch_Regex{
			Regex: path.Regex,
		}
	case *v1.RouteMatcher_Prefix:
		out.PathSpecifier = &envoyroute.RouteMatch_Prefix{
			Prefix: path.Prefix,
		}
	}
}

func envoyHeaderMatcher(in []*v1.HeaderMatcher) []*envoyroute.HeaderMatcher {
	var out []*envoyroute.HeaderMatcher
	for _, matcher := range in {
		envoyMatch := &envoyroute.HeaderMatcher{
			Name:  matcher.Name,
			Value: matcher.Value,
			Regex: &types.BoolValue{
				Value: matcher.Regex,
			},
		}
		out = append(out, envoyMatch)
	}
	return out
}

func envoyQueryMatcher(in []*v1.QueryParameterMatcher) []*envoyroute.QueryParameterMatcher {
	var out []*envoyroute.QueryParameterMatcher
	for _, matcher := range in {
		envoyMatch := &envoyroute.QueryParameterMatcher{
			Name:  matcher.Name,
			Value: matcher.Value,
			Regex: &types.BoolValue{
				Value: matcher.Regex,
			},
		}
		out = append(out, envoyMatch)
	}
	return out
}

// returns an error if any of the virtualhost domains overlap
func validateVirtualHostDomains(virtualHosts []*v1.VirtualHost) error {
	// this shouldbe a 1-1 mapping
	// if len(domainsToVirtualHosts[domain]) > 1, it's an error
	domainsToVirtualHosts := make(map[string][]string)
	for _, vHost := range virtualHosts {
		if len(vHost.Domains) == 0 {
			// default virtualhost
			domainsToVirtualHosts["*"] = append(domainsToVirtualHosts["*"], vHost.Name)
		}
		for _, domain := range vHost.Domains {
			// default virtualhost can be specified with empty string
			if domain == "" {
				domain = "*"
			}
			domainsToVirtualHosts[domain] = append(domainsToVirtualHosts[domain], vHost.Name)
		}
	}
	var domainErrors error
	// see if we found any conflicts, if so, write reports
	for domain, vHosts := range domainsToVirtualHosts {
		if len(vHosts) > 1 {
			domainErrors = multierror.Append(domainErrors, errors.Errorf("domain %v is "+
				"shared by the following virtual hosts: %v", domain, vHosts))
		}
	}
	return domainErrors
}

func validateRouteDestinations(upstreams []*v1.Upstream, action *v1.RouteAction) error {
	// make sure the destination itself has the right structure
	switch dest := action.Destination.(type) {
	case *v1.RouteAction_Single:
		return validateSingleDestination(upstreams, dest.Single)
	case *v1.RouteAction_Multi:
		return validateMultiDestination(upstreams, dest.Multi.Destinations)
	}
	return errors.Errorf("must specify either 'single_destination' or 'multiple_destinations' for action")
}

func validateMultiDestination(upstreams []*v1.Upstream, destinations []*v1.WeightedDestination) error {
	for _, dest := range destinations {
		if err := validateSingleDestination(upstreams, dest.Destination); err != nil {
			return errors.Wrap(err, "invalid destination in weighted destination list")
		}
	}
	return nil
}

func validateSingleDestination(upstreams []*v1.Upstream, destination *v1.Destination) error {
	upstreamName := destination.UpstreamName
	for _, us := range upstreams {
		if us.Metadata.Name == upstreamName {
			return nil
		}
	}
	return errors.Errorf("upstream %v was not found or had errors for upstream destination", upstreamName)
}

func validateListenerSslConfig(listener *v1.Listener, secrets []*v1.Secret) error {
	for _, ssl := range listener.SslConfiguations {
		switch secret := ssl.SslSecrets.(type) {
		case *v1.SSLConfig_SecretRef:
			if _, _, _, err := GetSslSecrets(secret.SecretRef, secrets); err != nil {
				return err
			}
		case *v1.SSLConfig_SslFiles:
			// TODO(ilackarms): validate SslFiles
		}
	}
	return nil
}

//TODO(ilackarms): move to a common utils file
const (
	deprecatedSslCertificateChainKey = "ca_chain"
	deprecatedSslPrivateKeyKey       = "private_key"
)

func GetSslSecrets(ref string, secrets []*v1.Secret) (string, string, string, error) {
	var sslSecret *v1.Secret
	for _, sec := range secrets {
		if sec.Metadata.Name == ref {
			sslSecret = sec
			break
		}
	}

	if sslSecret == nil {
		return "", "", "", errors.Errorf("ssl secret not found for ref %v", ref)
	}

	certChain, ok := sslSecret.Data[defaultv1.SslCertificateChainKey]
	if !ok {
		certChain, ok = sslSecret.Data[deprecatedSslCertificateChainKey]
		if !ok {
			return "", "", "", errors.Errorf("neither %v nor %v key not found in ssl secrets",
				defaultv1.SslCertificateChainKey, deprecatedSslCertificateChainKey)
		}
	}

	privateKey, ok := sslSecret.Data[defaultv1.SslPrivateKeyKey]
	if !ok {
		privateKey, ok = sslSecret.Data[deprecatedSslPrivateKeyKey]
		if !ok {
			return "", "", "", errors.Errorf("neither %v nor %v key not found in ssl secrets",
				defaultv1.SslPrivateKeyKey, deprecatedSslPrivateKeyKey)
		}
	}

	rootCa := sslSecret.Data[defaultv1.SslRootCaKey]
	return certChain, privateKey, rootCa, nil
}

func DataSourceFromString(str string) *envoycore.DataSource {
	return &envoycore.DataSource{
		Specifier: &envoycore.DataSource_InlineString{
			InlineString: str,
		},
	}
}