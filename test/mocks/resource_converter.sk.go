package converter

import (
	"errors"

	"github.com/solo-io/go-utils/versionutils/kubeapi"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"

	v1 "github.com/solo-io/solo-kit/test/mocks/v1"
	"github.com/solo-io/solo-kit/test/mocks/v1alpha1"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
)

type FakeResourceUpConverter interface {
	FromV1Alpha1ToV1(src *v1alpha1.FakeResource) *v1.FakeResource
	FromV1ToV2Alpha1(src *v1.FakeResource) *v2alpha1.FakeResource
}

type FakeResourceDownConverter interface {
	FromV1ToV1Alpha1(src *v1.FakeResource) *v1alpha1.FakeResource
}

type fakeResourceConverter struct {
	upConverter   FakeResourceUpConverter
	downConverter FakeResourceDownConverter
}

func NewFakeResourceConverter(u FakeResourceUpConverter, d FakeResourceDownConverter) crd.Converter {
	return &fakeResourceConverter{
		upConverter:   u,
		downConverter: d,
	}
}

func (c *fakeResourceConverter) Convert(src, dst crd.SoloKitCrd) error {
	srcVersion, err := kubeapi.ParseVersion(src.GetObjectKind().GroupVersionKind().Version)
	if err != nil {
		return err
	}
	dstVersion, err := kubeapi.ParseVersion(dst.GetObjectKind().GroupVersionKind().Version)
	if err != nil {
		return err
	}

	if srcVersion.GreaterThan(dstVersion) {
		return c.convertDown(src, dst)
	} else if srcVersion.LessThan(dstVersion) {
		return c.convertUp(src, dst)
	}
	return crd.Copy(src, dst)
}

func (c *fakeResourceConverter) convertUp(src, dst crd.SoloKitCrd) error {
	if src.GetObjectKind().GroupVersionKind().Version == dst.GetObjectKind().GroupVersionKind().Version {
		return crd.Copy(src, dst)
	}

	switch t := src.(type) {
	case *v1alpha1.FakeResource:
		return c.convertUp(c.upConverter.FromV1Alpha1ToV1(t), dst)
	case *v1.FakeResource:
		return c.convertUp(c.upConverter.FromV1ToV2Alpha1(t), dst)
	}
	return errors.New("unrecognized source type, this should never happen")
}

func (c *fakeResourceConverter) convertDown(src, dst crd.SoloKitCrd) error {
	if src.GetObjectKind().GroupVersionKind().Version == dst.GetObjectKind().GroupVersionKind().Version {
		return crd.Copy(src, dst)
	}

	switch t := src.(type) {
	case *v1.FakeResource:
		return c.convertUp(c.downConverter.FromV1ToV1Alpha1(t), dst)
	}
	return errors.New("unrecognized source type, this should never happen")
}

type MockResourceUpConverter interface {
	FromV1Alpha1ToV1(src *v1alpha1.MockResource) *v1.MockResource
	FromV1ToV2Alpha1(src *v1.MockResource) *v2alpha1.MockResource
}

type MockResourceDownConverter interface {
	FromV1ToV1Alpha1(src *v1.MockResource) *v1alpha1.MockResource
	FromV2Alpha1ToV1(src *v2alpha1.MockResource) *v1.MockResource
}

type mockResourceConverter struct {
	upConverter   MockResourceUpConverter
	downConverter MockResourceDownConverter
}

func NewMockResourceConverter(u MockResourceUpConverter, d MockResourceDownConverter) crd.Converter {
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
		return c.convertDown(src, dst)
	} else if srcVersion.LessThan(dstVersion) {
		return c.convertUp(src, dst)
	}
	return crd.Copy(src, dst)
}

func (c *mockResourceConverter) convertUp(src, dst crd.SoloKitCrd) error {
	if src.GetObjectKind().GroupVersionKind().Version == dst.GetObjectKind().GroupVersionKind().Version {
		return crd.Copy(src, dst)
	}

	switch t := src.(type) {
	case *v1alpha1.MockResource:
		return c.convertUp(c.upConverter.FromV1Alpha1ToV1(t), dst)
	case *v1.MockResource:
		return c.convertUp(c.upConverter.FromV1ToV2Alpha1(t), dst)
	}
	return errors.New("unrecognized source type, this should never happen")
}

func (c *mockResourceConverter) convertDown(src, dst crd.SoloKitCrd) error {
	if src.GetObjectKind().GroupVersionKind().Version == dst.GetObjectKind().GroupVersionKind().Version {
		return crd.Copy(src, dst)
	}

	switch t := src.(type) {
	case *v1.MockResource:
		return c.convertUp(c.downConverter.FromV1ToV1Alpha1(t), dst)
	case *v2alpha1.MockResource:
		return c.convertUp(c.downConverter.FromV2Alpha1ToV1(t), dst)
	}
	return errors.New("unrecognized source type, this should never happen")
}
