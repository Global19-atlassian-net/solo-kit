package mocks_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/test/mocks"
	v1 "github.com/solo-io/solo-kit/test/mocks/v1"
	"github.com/solo-io/solo-kit/test/mocks/v1alpha1"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
)

var (
	converter mocks.MockResourceConverter
)

var _ = Describe("MockResourceConverter", func() {

	BeforeEach(func() {
		converter = mocks.NewMockResourceConverter(upConverter{}, downConverter{})
	})

	Describe("Convert", func() {
		It("works for noop conversions", func() {
			src := &v2alpha1.MockResource{}
			dst := &v2alpha1.MockResource{}

			err := converter.Convert(src, dst)

			Expect(err).NotTo(HaveOccurred())
			Expect(reflect.TypeOf(dst)).To(Equal(reflect.TypeOf(&v2alpha1.MockResource{})))
		})
	})

})

type upConverter struct{}

func (upConverter) FromV1Alpha1ToV1(src *v1alpha1.MockResource) *v1.MockResource {
	return &v1.MockResource{}
}

func (upConverter) FromV1ToV2Alpha1(src *v1.MockResource) *v2alpha1.MockResource {
	return &v2alpha1.MockResource{}
}

type downConverter struct{}

func (downConverter) FromV2Alpha1ToV1(src *v2alpha1.MockResource) *v1.MockResource {
	return &v1.MockResource{}
}

func (downConverter) FromV1ToV1Alpha1(src *v1.MockResource) *v1alpha1.MockResource {
	return &v1alpha1.MockResource{}
}
