package mocks

import (
	"reflect"

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

type MockResourceConverter interface {
	Convert(src, dst crd.SoloKitCrd) error
}

type mockResourceConverter struct {
	upConverter   UpConverter
	downConverter DownConverter
}

func NewMockResourceConverter(u UpConverter, d DownConverter) MockResourceConverter {
	return &mockResourceConverter{
		upConverter:   u,
		downConverter: d,
	}
}

func (c *mockResourceConverter) Convert(src, dst crd.SoloKitCrd) error {
	srcVersion, err := kubeapi.ParseVersion(src.GetObjectKind().GroupVersionKind().Version)
	if err != nil {
		return err
	}
	dstVersion, err := kubeapi.ParseVersion(dst.GetObjectKind().GroupVersionKind().Version)
	if err != nil {
		return err
	}

	if srcVersion.GreaterThan(dstVersion) {
		dst = c.convertDown(src, dst)
	} else if srcVersion.LessThan(dstVersion) {
		c.convertUp(src, dst)
	} else {
		dst = src
	}

	return nil
}

func (c *mockResourceConverter) convertDown(src, dst crd.SoloKitCrd) crd.SoloKitCrd {
	if reflect.TypeOf(src) == reflect.TypeOf(dst) {
		return src
	}

	switch t := src.(type) {
	case *v2alpha1.MockResource:
		return c.convertDown(c.downConverter.FromV2Alpha1ToV1(t), dst)
	case *v1.MockResource:
		return c.convertDown(c.downConverter.FromV1ToV1Alpha1(t), dst)
	}
	return nil
}

func (c *mockResourceConverter) convertUp(src, dst crd.SoloKitCrd) {
	if reflect.TypeOf(src) == reflect.TypeOf(dst) {
		dst = src
		return
	}

	switch t := src.(type) {
	case *v1alpha1.MockResource:
		c.convertUp(c.upConverter.FromV1Alpha1ToV1(t), dst)
	case *v1.MockResource:
		c.convertUp(c.upConverter.FromV1ToV2Alpha1(t), dst)
	}
}
