// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
)

var _ = Describe("TestingEventLoop", func() {
	var (
		namespace string
		emitter   TestingEmitter
		err       error
	)

	BeforeEach(func() {

		mockResourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		mockResourceClient, err := NewMockResourceClient(mockResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		fakeResourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		fakeResourceClient, err := NewFakeResourceClient(fakeResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		anotherMockResourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		anotherMockResourceClient, err := NewAnotherMockResourceClient(anotherMockResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		emitter = NewTestingEmitter(mockResourceClient, fakeResourceClient, anotherMockResourceClient)
	})
	It("runs sync function on a new snapshot", func() {
		_, err = emitter.MockResource().Write(NewMockResource(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.FakeResource().Write(NewFakeResource(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.AnotherMockResource().Write(NewAnotherMockResource(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		sync := &mockTestingSyncer{}
		el := NewTestingEventLoop(emitter, sync)
		_, err := el.Run([]string{namespace}, clients.WatchOpts{})
		Expect(err).NotTo(HaveOccurred())
		Eventually(sync.Synced, 5*time.Second).Should(BeTrue())
	})
})

type mockTestingSyncer struct {
	synced bool
	mutex  sync.Mutex
}

func (s *mockTestingSyncer) Synced() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.synced
}

func (s *mockTestingSyncer) Sync(ctx context.Context, snap *TestingSnapshot) error {
	s.mutex.Lock()
	s.synced = true
	s.mutex.Unlock()
	return nil
}