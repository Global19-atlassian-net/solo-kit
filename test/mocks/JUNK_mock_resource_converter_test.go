package mocks_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/test/mocks"
	v1 "github.com/solo-io/solo-kit/test/mocks/v1"
	"github.com/solo-io/solo-kit/test/mocks/v1alpha1"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
)

var (
	converter crd.Converter
)

var _ = Describe("MockResourceConverter", func() {
	BeforeEach(func() {
		converter = mocks.NewMockResourceConverter(upConverter{}, downConverter{})
	})

	Describe("Convert", func() {
		It("works for noop conversions", func() {
			src := &v2alpha1.MockResource{Metadata: core.Metadata{Name: "name"}}
			dst := &v2alpha1.MockResource{}
			err := converter.Convert(src, dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.GetMetadata().Name).To(Equal("name"))
		})

		It("converts all the way up", func() {
			src := &v1alpha1.MockResource{}
			dst := &v2alpha1.MockResource{}
			err := converter.Convert(src, dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.GetMetadata().Name).To(Equal("v2alpha1"))
		})

		It("converts all the way down", func() {
			src := &v2alpha1.MockResource{}
			dst := &v1alpha1.MockResource{}
			err := converter.Convert(src, dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.GetMetadata().Name).To(Equal("v1alpha1"))
		})
	})
})

type upConverter struct{}

func (upConverter) FromV1Alpha1ToV1(src *v1alpha1.MockResource) *v1.MockResource {
	return &v1.MockResource{Metadata: core.Metadata{Name: "v1"}}
}

func (upConverter) FromV1ToV2Alpha1(src *v1.MockResource) *v2alpha1.MockResource {
	return &v2alpha1.MockResource{Metadata: core.Metadata{Name: "v2alpha1"}}
}

type downConverter struct{}

func (downConverter) FromV2Alpha1ToV1(src *v2alpha1.MockResource) *v1.MockResource {
	return &v1.MockResource{Metadata: core.Metadata{Name: "v1"}}
}

func (downConverter) FromV1ToV1Alpha1(src *v1.MockResource) *v1alpha1.MockResource {
	return &v1alpha1.MockResource{Metadata: core.Metadata{Name: "v1alpha1"}}
}
