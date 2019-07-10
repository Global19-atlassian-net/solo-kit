package mocks

import (
	"github.com/solo-io/go-utils/versionutils/kubeapi"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	v1 "github.com/solo-io/solo-kit/test/mocks/v1"
	"github.com/solo-io/solo-kit/test/mocks/v1alpha1"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
)

type UpConverter interface {
	FromV1Alpha1ToV1(src *v1alpha1.MockResource) *v1.MockResource
	FromV1ToV2Alpha1(src *v1.MockResource) *v2alpha1.MockResource
}

type DownConverter interface {
	FromV2Alpha1ToV1(src *v2alpha1.MockResource) *v1.MockResource
	FromV1ToV1Alpha1(src *v1.MockResource) *v1alpha1.MockResource
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
		c.convertDown(src, dst)
	}
	if srcVersion.LessThan(dstVersion) {
		c.convertUp(src, dst)
	}

	return nil
}

func (c *MockResourceConverter) convertDown(src, dst crd.SoloKitCrd) {
	if src.GetObjectKind().GroupVersionKind().Version == dst.GetObjectKind().GroupVersionKind().Version {
		return
	}

	switch src.GetObjectKind().GroupVersionKind().Version {
	case "v2alpha1":
		src = c.DownConverter.FromV2Alpha1ToV1(src.(*v2alpha1.MockResource))
	case "v1":
		src = c.DownConverter.FromV1ToV1Alpha1(src.(*v1.MockResource))
	default:
		return
	}
	c.convertDown(src, dst)
}

func (c *MockResourceConverter) convertUp(src, dst crd.SoloKitCrd) {
	if src.GetObjectKind().GroupVersionKind().Version == dst.GetObjectKind().GroupVersionKind().Version {
		return
	}

	switch src.GetObjectKind().GroupVersionKind().Version {
	case "v1alpha1":
		src = c.UpConverter.FromV1Alpha1ToV1(src.(*v1alpha1.MockResource))
	case "v1":
		src = c.UpConverter.FromV1ToV2Alpha1(src.(*v1.MockResource))
	default:
		return
	}
	c.convertUp(src, dst)
}
