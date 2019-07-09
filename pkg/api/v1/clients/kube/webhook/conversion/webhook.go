package conversion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	v1 "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/solo.io/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/webhook/server"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
	apix "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type kubeWebhook struct {
	decoder *server.Decoder
	scheme  *runtime.Scheme

	ctx      context.Context
	server   *server.Server
	resource *crd.MultiVersionCrd

	converter crd.Converter
}

func NewKubeWebhook(ctx context.Context, server *server.Server, gk schema.GroupKind, converter crd.Converter) (*kubeWebhook, error) {
	resource, err := crd.GetMultiVersionCrd(gk)
	if err != nil {
		return nil, err
	}
	kw := &kubeWebhook{
		server:    server,
		resource:  &resource,
		ctx:       ctx,
		converter: converter,
	}
	return kw, nil
}

// InjectScheme injects a scheme into the webhook, in order to construct a Decoder.
func (k *kubeWebhook) InjectScheme(s *runtime.Scheme) error {
	var err error
	k.scheme = s
	k.decoder, err = server.NewDecoder(s)
	if err != nil {
		return err
	}

	return nil
}

func (k *kubeWebhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := contextutils.LoggerFrom(k.ctx)
	var convertReview apix.ConversionReview
	err := json.NewDecoder(r.Body).Decode(&convertReview)
	if err != nil {
		logger.Error(err, "failed to read conversion request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := k.handleConvertRequest(convertReview.Request)
	if err != nil {
		logger.Error(err, "failed to convert", "request", convertReview.Request.UID)
		convertReview.Response = errored(err)
	} else {
		convertReview.Response = resp
	}
	convertReview.Response.UID = convertReview.Request.UID
	convertReview.Request = nil

	err = json.NewEncoder(w).Encode(convertReview)
	if err != nil {
		logger.Error(err, "failed to write response")
		return
	}
}

func (k *kubeWebhook) handleConvertRequest(req *apix.ConversionRequest) (*apix.ConversionResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("conversion request is nil")
	}

	desiredGv, err := schema.ParseGroupVersion(req.DesiredAPIVersion)
	if err != nil {
		return nil, err
	}

	var objects []runtime.RawExtension

	for _, obj := range req.Objects {
		src, gvk, err := k.translateSrcObj(obj.Raw)
		if err != nil {
			return nil, err
		}

		resourceDst, err := k.translateDstObj(desiredGv)
		if err != nil {
			return nil, err
		}
		if err := k.converter.Convert(src, resourceDst); err != nil {
			return nil, err
		}

		dst, err := k.allocateDstObject(resourceDst, req.DesiredAPIVersion, gvk.Kind)
		if err != nil {
			return nil, err
		}
		objects = append(objects, runtime.RawExtension{Object: dst})
	}
	return &apix.ConversionResponse{
		UID:              req.UID,
		ConvertedObjects: objects,
		Result: metav1.Status{
			Status: metav1.StatusSuccess,
		},
	}, nil
}

// allocateDstObject returns an instance for a given GVK.
func (k *kubeWebhook) allocateDstObject(resource crd.SoloKitCrd, apiVersion, kind string) (runtime.Object, error) {
	gvk := schema.FromAPIVersionAndKind(apiVersion, kind)

	obj, err := k.scheme.New(gvk)
	if err != nil {
		return obj, err
	}

	t, err := meta.TypeAccessor(obj)
	if err != nil {
		return obj, err
	}

	t.SetAPIVersion(apiVersion)
	t.SetKind(kind)

	res := crd.KubeResource(resource)
	return res, nil
}

func (k *kubeWebhook) translateDstObj(desiredGv schema.GroupVersion) (crd.SoloKitCrd, error) {
	resourceVersion, err := k.resource.GetVersion(desiredGv.Version)
	if err != nil {
		return nil, err
	}
	resource := resources.Clone(resourceVersion.Type)

	soloKitCrd, ok := resource.(crd.SoloKitCrd)
	if !ok {
		return nil, errors.New("could not translate to solo-kit crd type")
	}

	return soloKitCrd, nil
}

func (k *kubeWebhook) translateSrcObj(byt []byte) (crd.SoloKitCrd, *schema.GroupVersionKind, error) {
	src, gvk, err := k.decoder.Decode(byt)
	if err != nil {
		return nil, nil, err
	}

	resourceCrd, ok := src.(*v1.Resource)
	if !ok {
		return nil, nil, errors.New("could not translate to resource type")
	}

	resourceVersion, err := k.resource.GetVersion(gvk.Version)
	if err != nil {
		return nil, nil, err
	}

	resource := resources.Clone(resourceVersion.Type)

	if resourceCrd.Spec != nil {
		if err := protoutils.UnmarshalMap(*resourceCrd.Spec, resource); err != nil {
			return nil, nil, errors.Wrapf(err, "reading crd spec into %v", src.GetObjectKind())
		}
	}

	resource.SetMetadata(kubeutils.FromKubeMeta(resourceCrd.ObjectMeta))
	if withStatus, ok := resource.(resources.InputResource); ok {
		resources.UpdateStatus(withStatus, func(status *core.Status) {
			*status = resourceCrd.Status
		})
	}

	soloKitCrd, ok := resource.(crd.SoloKitCrd)
	if !ok {
		return nil, nil, errors.New("could not translate to solo-kit crd type")
	}

	return soloKitCrd, gvk, nil
}

// helper to construct error response.
func errored(err error) *apix.ConversionResponse {
	return &apix.ConversionResponse{
		Result: metav1.Status{
			Status:  metav1.StatusFailure,
			Message: err.Error(),
		},
	}
}
