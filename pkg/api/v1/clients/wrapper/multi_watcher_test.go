package wrapper_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	. "github.com/solo-io/solo-kit/pkg/api/v1/clients/wrapper"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	v1 "github.com/solo-io/solo-kit/test/mocks/v1"
	"github.com/solo-io/solo-kit/test/util"
)

var _ = Describe("MultiWatcher", func() {
	var base1, base2 *memory.ResourceClient
	var watcher *MultiWatcher
	clusterName1 := "clustr1"
	clusterName2 := "clustr2"
	BeforeEach(func() {
		base1 = memory.NewResourceClient(memory.NewInMemoryResourceCache(), &v1.MockResource{})
		base2 = memory.NewResourceClient(memory.NewInMemoryResourceCache(), &v1.MockResource{})
		cluster1 := NewClusterClient(base1, clusterName1)
		cluster2 := NewClusterClient(base2, clusterName2)
		watcher = &MultiWatcher{Watchers: []clients.ResourceWatcher{cluster1, cluster2}}
	})
	It("aggregates watches from multiple clients", func() {
		watch, errs, err := watcher.Watch("", clients.WatchOpts{RefreshRate: time.Millisecond})
		Expect(err).NotTo(HaveOccurred())
		wait := make(chan struct{})
		go func() {
			defer GinkgoRecover()
			defer func() {
				close(wait)
			}()
			_, err = base1.Write(v1.NewMockResource("a", "a"), clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())
			_, err = base1.Write(v1.NewMockResource("a", "b"), clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())
			_, err = base2.Write(v1.NewMockResource("a", "a"), clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())
			_, err = base2.Write(v1.NewMockResource("a", "b"), clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())
		}()
		select {
		case <-wait:
		case <-time.After(time.Second * 5):
			Fail("expected wait to be closed before 5s")
		}

		var list resources.ResourceList

		Eventually(func() resources.ResourceList {
			select {
			default:
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case list = <-watch:
				return list
			case <-time.After(time.Millisecond * 5):
				Fail("expected a message in channel")
			}
			return nil
		}, time.Second*200).Should(HaveLen(4))

		list.Each(util.ZeroResourceVersion)

		Expect(list).To(Equal(resources.ResourceList{
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "a", Cluster: "clustr1"}},
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "b", Cluster: "clustr1"}},
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "a", Cluster: "clustr2"}},
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "b", Cluster: "clustr2"}},
		}))
	})
})
