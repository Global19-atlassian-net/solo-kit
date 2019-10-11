// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1alpha1

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/wrapper"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/tests/typed"
)

var _ = Describe("FakeResourceMultiClusterClient", func() {
	var (
		namespace string
	)
	for _, test := range []typed.ResourceClientTester{
		&typed.KubeRcTester{Crd: FakeResourceCrd},
		&typed.ConsulRcTester{},
		&typed.FileRcTester{},
		&typed.MemoryRcTester{},
		&typed.VaultRcTester{},
		&typed.KubeSecretRcTester{},
		&typed.KubeConfigMapRcTester{},
	} {
		Context("multi cluster resource client backed by "+test.Description(), func() {
			var (
				client              FakeResourceMultiClusterClient
				name1, name2, name3 = "foo" + helpers.RandString(3), "boo" + helpers.RandString(3), "goo" + helpers.RandString(3)
			)

			BeforeEach(func() {
				namespace = helpers.RandString(6)
				test.Setup(namespace)
			})
			AfterEach(func() {
				test.Teardown(namespace)
			})
			It("CRUDs FakeResources "+test.Description(), func() {
				client = NewFakeResourceMultiClusterClient(test)
				FakeResourceMultiClusterClientTest(namespace, client, name1, name2, name3)
			})
			It("errors when no client exists for the given cluster "+test.Description(), func() {
				client = NewFakeResourceMultiClusterClient(test)
				FakeResourceMultiClusterClientCrudErrorsTest(client)
			})
			It("populates an aggregated watch "+test.Description(), func() {
				watchAggregator := wrapper.NewWatchAggregator()
				client = NewFakeResourceMultiClusterClientWithWatchAggregator(watchAggregator, test)
				FakeResourceMultiClusterClientWatchAggregationTest(client, watchAggregator, namespace)
			})
		})
	}
})

func FakeResourceMultiClusterClientTest(namespace string, client FakeResourceMultiClusterClient, name1, name2, name3 string) {
	cfg, err := kubeutils.GetConfig("", "")
	Expect(err).NotTo(HaveOccurred())
	client.ClusterAdded("", cfg)

	name := name1
	input := NewFakeResource(namespace, name)

	r1, err := client.Write(input, clients.WriteOpts{})
	Expect(err).NotTo(HaveOccurred())

	_, err = client.Write(input, clients.WriteOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsExist(err)).To(BeTrue())

	Expect(r1).To(BeAssignableToTypeOf(&FakeResource{}))
	Expect(r1.GetMetadata().Name).To(Equal(name))
	Expect(r1.GetMetadata().Namespace).To(Equal(namespace))
	Expect(r1.GetMetadata().ResourceVersion).NotTo(Equal(input.GetMetadata().ResourceVersion))
	Expect(r1.GetMetadata().Ref()).To(Equal(input.GetMetadata().Ref()))
	Expect(r1.Count).To(Equal(input.Count))

	_, err = client.Write(input, clients.WriteOpts{
		OverwriteExisting: true,
	})
	Expect(err).To(HaveOccurred())

	resources.UpdateMetadata(input, func(meta *core.Metadata) {
		meta.ResourceVersion = r1.GetMetadata().ResourceVersion
	})
	r1, err = client.Write(input, clients.WriteOpts{
		OverwriteExisting: true,
	})
	Expect(err).NotTo(HaveOccurred())
	read, err := client.Read(namespace, name, clients.ReadOpts{})
	Expect(err).NotTo(HaveOccurred())
	Expect(read).To(Equal(r1))
	_, err = client.Read("doesntexist", name, clients.ReadOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsNotExist(err)).To(BeTrue())

	name = name2
	input = &FakeResource{}

	input.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})

	r2, err := client.Write(input, clients.WriteOpts{})
	Expect(err).NotTo(HaveOccurred())
	list, err := client.List(namespace, clients.ListOpts{})
	Expect(err).NotTo(HaveOccurred())
	Expect(list).To(ContainElement(r1))
	Expect(list).To(ContainElement(r2))
	err = client.Delete(namespace, "adsfw", clients.DeleteOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsNotExist(err)).To(BeTrue())
	err = client.Delete(namespace, "adsfw", clients.DeleteOpts{
		IgnoreNotExist: true,
	})
	Expect(err).NotTo(HaveOccurred())
	err = client.Delete(namespace, r2.GetMetadata().Name, clients.DeleteOpts{})
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() FakeResourceList {
		list, err = client.List(namespace, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		return list
	}, time.Second*10).Should(ContainElement(r1))
	Eventually(func() FakeResourceList {
		list, err = client.List(namespace, clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		return list
	}, time.Second*10).ShouldNot(ContainElement(r2))
	w, errs, err := client.Watch(namespace, clients.WatchOpts{
		RefreshRate: time.Hour,
	})
	Expect(err).NotTo(HaveOccurred())

	var r3 resources.Resource
	wait := make(chan struct{})
	go func() {
		defer close(wait)
		defer GinkgoRecover()

		resources.UpdateMetadata(r2, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		r2, err = client.Write(r2, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())

		name = name3
		input = &FakeResource{}
		Expect(err).NotTo(HaveOccurred())
		input.SetMetadata(core.Metadata{
			Name:      name,
			Namespace: namespace,
		})

		r3, err = client.Write(input, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
	}()
	<-wait

	select {
	case err := <-errs:
		Expect(err).NotTo(HaveOccurred())
	case list = <-w:
	case <-time.After(time.Millisecond * 5):
		Fail("expected a message in channel")
	}

	go func() {
		defer GinkgoRecover()
		for {
			select {
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Second / 4):
				return
			}
		}
	}()

	Eventually(w, time.Second*5, time.Second/10).Should(Receive(And(ContainElement(r1), ContainElement(r3), ContainElement(r3))))
}

func FakeResourceMultiClusterClientCrudErrorsTest(client FakeResourceMultiClusterClient) {
	_, err := client.Read("foo", "bar", clients.ReadOpts{Cluster: "read"})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoFakeResourceClientForClusterError("read").Error()))
	_, err = client.List("foo", clients.ListOpts{Cluster: "list"})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoFakeResourceClientForClusterError("list").Error()))
	err = client.Delete("foo", "bar", clients.DeleteOpts{Cluster: "delete"})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoFakeResourceClientForClusterError("delete").Error()))

	input := &FakeResource{}
	input.SetMetadata(core.Metadata{
		Cluster:   "write",
		Name:      "bar",
		Namespace: namespace,
	})
	_, err = client.Write(input, clients.WriteOpts{})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoFakeResourceClientForClusterError("write").Error()))
	_, _, err = client.Watch("foo", clients.WatchOpts{Cluster: "watch"})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoFakeResourceClientForClusterError("watch").Error()))
}
func FakeResourceMultiClusterClientWatchAggregationTest(client FakeResourceMultiClusterClient, aggregator wrapper.WatchAggregator, namespace string) {
	w, errs, err := aggregator.Watch(namespace, clients.WatchOpts{})
	Expect(err).NotTo(HaveOccurred())
	go func() {
		defer GinkgoRecover()
		for {
			select {
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Second / 4):
				return
			}
		}
	}()

	cfg, err := kubeutils.GetConfig("", "")
	Expect(err).NotTo(HaveOccurred())
	client.ClusterAdded("", cfg)
	input := &FakeResource{}
	input.SetMetadata(core.Metadata{
		Cluster:   "write",
		Name:      "bar",
		Namespace: namespace,
	})
	_, err = client.Write(input, clients.WriteOpts{})
	written, err := client.Write(input, clients.WriteOpts{})
	Expect(err).NotTo(HaveOccurred())
	Eventually(w, time.Second*5, time.Second/10).Should(Receive(And(ContainElement(written))))
}
