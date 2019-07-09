package webhook

import (
	"encoding/base64"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/webhook/certwatcher"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
	"gopkg.in/yaml.v2"
)

var _ = Describe("webhook e2e", func() {

	var (
		meta = core.Metadata{
			Namespace: "default",
			Name:      "crd-converter",
		}
	)
	It("can print CRD", func() {
		custom, err := crd.GetRegistry().GetKubeCrd(v2alpha1.MockResourceCrd.GroupVersionKind())
		Expect(err).NotTo(HaveOccurred())
		byt, err := yaml.Marshal(custom)
		Expect(err).NotTo(HaveOccurred())
		fmt.Println(string(byt))
	})

	FIt("can print certificates", func() {
		certs, err := certwatcher.GenerateSelfSignedPodCerts(meta)
		Expect(err).NotTo(HaveOccurred())
		fmt.Println(base64.StdEncoding.EncodeToString(certs.CaCertificate))
		fmt.Println(base64.StdEncoding.EncodeToString(certs.ServerCertificate))
		fmt.Println(base64.StdEncoding.EncodeToString(certs.ServerCertKey))
	})
})
