package crd_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/client/clientset/versioned/scheme"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	// "sigs.k8s.io/controller-runtime/pkg/conversion"
	// "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrl "sigs.k8s.io/controller-runtime"
)

var _ = FDescribe("resources", func() {

	var (
		stopCh chan struct{}
		mgr    manager.Manager
	)

	BeforeEach(func() {
		stopCh = make(chan struct{})
		cfg, err := kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		// apiexts, err := clientset.NewForConfig(cfg)
		// Expect(err).NotTo(HaveOccurred())
		// err = v2alpha1.MockResourceCrd.Register(apiexts)
		// Expect(err).NotTo(HaveOccurred())

		// err = v2alpha1.MockResourceAddToScheme(scheme.Scheme)
		mgr, err = ctrl.NewManager(cfg, ctrl.Options{
			Scheme: scheme.Scheme,
		})
		go func() {
			defer GinkgoRecover()
			err := mgr.Start(stopCh)
			Expect(err).NotTo(HaveOccurred())
		}()
	})

	AfterEach(func() {
		select {
		case stopCh <- struct{}{}:
		case <-time.After(time.Second * 5):
			Fail("could not stop manager")

		}
	})

	It("can do stuff", func() {
		var err error
		Expect(err).NotTo(HaveOccurred())
		var item v2alpha1.MockResource
		time.Sleep(1 * time.Second)
		err = mgr.GetClient().Get(context.TODO(), types.NamespacedName{Name: "one", Namespace: "default"}, &item)
		Expect(err).NotTo(HaveOccurred())

	})
})
