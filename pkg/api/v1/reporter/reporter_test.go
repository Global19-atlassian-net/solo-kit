package reporter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"
	"fmt"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	rep "github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/test/mocks"
)

var _ = Describe("Reporter", func() {
	var (
		reporter                               rep.Reporter
		mockResourceClient, fakeResourceClient clients.ResourceClient
	)
	BeforeEach(func() {
		mockResourceClient = memory.NewResourceClient(&mocks.MockResource{})
		fakeResourceClient = memory.NewResourceClient(&mocks.FakeResource{})
		reporter = rep.NewReporter(mockResourceClient, fakeResourceClient)
	})
	AfterEach(func() {
	})
	It("reports errors for resources", func() {
		r1, err := mockResourceClient.Write(mocks.NewMockResource("", "mocky"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		r2, err := mockResourceClient.Write(mocks.NewMockResource("", "fakey"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		resourceErrs := rep.ResourceErrors{
			r1: fmt.Errorf("everyone makes mistakes"),
			r2: fmt.Errorf("try your best"),
		}
		err = reporter.WriteReports(context.TODO(), resourceErrs)
		Expect(err).NotTo(HaveOccurred())

		r1, err = mockResourceClient.Read(r1.GetMetadata().Name, clients.ReadOpts{
			Namespace: r1.GetMetadata().Namespace,
		})
		Expect(err).NotTo(HaveOccurred())
		r2, err = mockResourceClient.Read(r2.GetMetadata().Name, clients.ReadOpts{
			Namespace: r2.GetMetadata().Namespace,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(r1.GetStatus()).To(Equal(core.Status{
			State:  2,
			Reason: "everyone makes mistakes",
		}))
		Expect(r2.GetStatus()).To(Equal(core.Status{
			State:  2,
			Reason: "try your best",
		}))
	})
})
