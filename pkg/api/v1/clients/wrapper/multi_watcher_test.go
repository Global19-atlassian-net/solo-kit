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

var _ = Describe("watchAggregator", func() {
	var cluster1, cluster2, cluster3 *Client // add / remove later
	var watcher WatchAggregator
	clusterName1 := "clustr1"
	clusterName2 := "clustr2"
	clusterName3 := "clustr3"
	BeforeEach(func() {
		base1 := memory.NewResourceClient(memory.NewInMemoryResourceCache(), &v1.MockResource{})
		base2 := memory.NewResourceClient(memory.NewInMemoryResourceCache(), &v1.MockResource{})
		base3 := memory.NewResourceClient(memory.NewInMemoryResourceCache(), &v1.MockResource{})
		cluster1 = NewClusterClient(base1, clusterName1)
		cluster2 = NewClusterClient(base2, clusterName2)
		cluster3 = NewClusterClient(base3, clusterName3)

		watcher = NewWatchAggregator()
		err := watcher.AddWatch(cluster1)
		Expect(err).NotTo(HaveOccurred())
		err = watcher.AddWatch(cluster2)
		Expect(err).NotTo(HaveOccurred())
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
			_, err = cluster1.Write(v1.NewMockResource("a", "a"), clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())
			_, err = cluster1.Write(v1.NewMockResource("a", "b"), clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())
			_, err = cluster2.Write(v1.NewMockResource("a", "a"), clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())
			_, err = cluster2.Write(v1.NewMockResource("a", "b"), clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			err = watcher.AddWatch(cluster3)
			Expect(err).NotTo(HaveOccurred())

			_, err = cluster3.Write(v1.NewMockResource("a", "a"), clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())
			_, err = cluster3.Write(v1.NewMockResource("a", "b"), clients.WriteOpts{})
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
		}, time.Second*200).Should(HaveLen(6))

		list.Each(util.ZeroResourceVersion)

		Expect(list).To(Equal(resources.ResourceList{
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "a", Cluster: "clustr1"}},
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "b", Cluster: "clustr1"}},
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "a", Cluster: "clustr2"}},
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "b", Cluster: "clustr2"}},
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "a", Cluster: "clustr3"}},
			&v1.MockResource{Metadata: core.Metadata{Namespace: "a", Name: "b", Cluster: "clustr3"}},
		}))
	})
})
