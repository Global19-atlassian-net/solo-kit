package mocks

import (
	"github.com/solo-io/go-utils/versionutils/kubeapi"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/errors"
	v1 "github.com/solo-io/solo-kit/test/mocks/v1"
	"github.com/solo-io/solo-kit/test/mocks/v1alpha1"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
)

type UpConverter interface {
	FromV1Alpha1ToV1(src *v1alpha1.MockResource, dst *v1.MockResource) error
	FromV1ToV2Alpha1(src *v1.MockResource, dst *v2alpha1.MockResource) error
}

type DownConverter interface {
	FromV2Alpha1ToV1(src *v2alpha1.MockResource, dst *v1.MockResource) error
	FromV1ToV1Alpha1(src *v1.MockResource, dst *v1alpha1.MockResource) error
}

type MockResourceConverter struct {
	UpConverter   UpConverter
	DownConverter DownConverter
}

func (c *MockResourceConverter) Convert(src, dst crd.SoloKitCrd) error {
	srcVersion, err := kubeapi.ParseVersion(src.GetObjectKind().GroupVersionKind().Version)
	if err != nil {
		return err
	}

	dstVersion, err := kubeapi.ParseVersion(dst.GetObjectKind().GroupVersionKind().Version)
	if err != nil {
		return err
	}

	if srcVersion.GreaterThan(dstVersion) {
		switch src.GetObjectKind().GroupVersionKind().Version {
		case "v2alpha1":
			err = c.DownConverter.FromV2Alpha1ToV1(src.(*v2alpha1.MockResource), dst.(*v1.MockResource))
			if err != nil {
				return err
			}
		case "v1":
			err = c.DownConverter.FromV1ToV1Alpha1(src.(*v1.MockResource), dst.(*v1alpha1.MockResource))
			if err != nil {
				return err
			}
		default:
			return errors.Errorf("Cannot convert up from %v", src.GetObjectKind().GroupVersionKind().Version)
		}
		return c.Convert(src, dst)
	}

	if srcVersion.LessThan(dstVersion) {
		switch src.GetObjectKind().GroupVersionKind().Version {
		case "v1alpha1":
			err = c.UpConverter.FromV1Alpha1ToV1(src.(*v1alpha1.MockResource), dst.(*v1.MockResource))
			if err != nil {
				return err
			}
		case "v1":
			err = c.UpConverter.FromV1ToV2Alpha1(src.(*v1.MockResource), dst.(*v2alpha1.MockResource))
			if err != nil {
				return err
			}
		default:
			return errors.Errorf("Cannot convert up from %v", src.GetObjectKind().GroupVersionKind().Version)
		}
		return c.Convert(src, dst)
	}

	return nil
}
