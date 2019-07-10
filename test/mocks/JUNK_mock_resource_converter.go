package mocks

import (
	"errors"
	"reflect"

	"github.com/solo-io/go-utils/versionutils/kubeapi"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
	v1 "github.com/solo-io/solo-kit/test/mocks/v1"
	"github.com/solo-io/solo-kit/test/mocks/v1alpha1"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
)

var unrecognizedTypeError = errors.New("unrecognized source type. Should never happen")

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

type Converter interface {
	Convert(src, dst crd.SoloKitCrd) (crd.SoloKitCrd, error)
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
		return c.convertDown(src, dst)
	} else if srcVersion.LessThan(dstVersion) {
		return c.convertUp(src, dst)
	}
	return writeSrcToDst(src, dst)
}

func (c *mockResourceConverter) convertDown(src, dst crd.SoloKitCrd) error {
	if reflect.TypeOf(src) == reflect.TypeOf(dst) {
		return writeSrcToDst(src, dst)
	}

	switch t := src.(type) {
	case *v2alpha1.MockResource:
		return c.convertDown(c.downConverter.FromV2Alpha1ToV1(t), dst)
	case *v1.MockResource:
		return c.convertDown(c.downConverter.FromV1ToV1Alpha1(t), dst)
	}
	return unrecognizedTypeError
}

func (c *mockResourceConverter) convertUp(src, dst crd.SoloKitCrd) error {
	if reflect.TypeOf(src) == reflect.TypeOf(dst) {
		return writeSrcToDst(src, dst)
	}

	switch t := src.(type) {
	case *v1alpha1.MockResource:
		return c.convertUp(c.upConverter.FromV1Alpha1ToV1(t), dst)
	case *v1.MockResource:
		return c.convertUp(c.upConverter.FromV1ToV2Alpha1(t), dst)
	}
	return unrecognizedTypeError
}

func writeSrcToDst(src, dst crd.SoloKitCrd) error {
	srcBytes, err := protoutils.MarshalBytes(src)
	if err != nil {
		return err
	}
	err = protoutils.UnmarshalBytes(srcBytes, dst)
	if err != nil {
		return err
	}
	return nil
}
