package discovery_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/multicluster/discovery"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

var _ = Describe("Gke", func() {
	if os.Getenv("RUN_GKE_TESTS") != "1" {
		log.Printf("This test creates GKE resources and is disabled by default. To enable, set RUN_GKE_TESTS=1 in your env.")
		//return
	}
	Context("multiple clusters created", func() {
		BeforeEach(func() {

		})
		It("detects the clusters and creates kube configs from them", func() {

			ctx := context.TODO()
			creds := []byte(`{
  "type": "service_account",
  "project_id": "solo-public",
  "private_key_id": "e324178c1a4bb05d6097de3287848658609068e3",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCfUlq1f8FMx3NZ\nkXfHV70xCHcf54VArnPfnMJWbIwmsGHWE6oy2I4WKItFU1ens6TkLdZUjt9gsica\n5b+7cfPV1u9RCe8kbTavJPtqxBD1gY20x6CTgpfzkzRFVgesQQKmdNjLT3+gijBt\n1PCC1k2HR7JKS6FK8CKmOFj8l0bdorHoSy+mg70vY087gApBiencg15pCcXtYdpY\nfNtisAW/mgoaxfwiJ8bSGh1a02iopPJXaaEvXHyo1y7pd000ERbpBESN4JXi1oy6\nZx5rrx/HPAIEm0ksrRik+vMm4t0oZIOFdxBvPeUb9zj4VbGRH7WNH382mpn8Yp2s\niDKBs87BAgMBAAECggEAGSgjlxOeZd17fwnbcDBQORkVWEYSMiGpvcYOoJ9grO4+\ncJNn1UIBA4ow3Yg4p8wTrgz30h/CMU5IIvTiproqTppy7Oik6GtPTm4IPfZKGD31\n7nT0C767+BPHFeis6vvT6405OzcpF6QiXAFa3cnxcXo2colJJWBQFE+C65WGdNHv\nPh7Ermga4P906tTni+kr9tX9yxAYhFQ7WCwqukhDh+yWq2fzIY/ih9kkcYtc1Uz7\nJ18a1uPaQDqapWugZ0LuVkQDW/7V0Q8kDmWD28XwZoIfCIDRBOycw7l8FoD4Q9ik\n2gSBaSs/4skVQ8xVaqSkz9VTb6HAHFhokUsDo68pLQKBgQDLyvx8wJnx1nIcxRqX\nMOeI3jhLBYrUABB4bqGbbJ3Ui/yDYDlLrc/g/rThPC82/DfZ8y3df69wMePDLCYw\nJPsLbBWbvo+JRiWnl7BcyiswEnPBJOOREg8I2WHk2ycCrbZLElcEF30QJVEj8zzs\nhVVU288pZFk0MBt0jO0bEYceFQKBgQDIIuK7jV+Sx8S5wZyjSoErqcTnRCer/9dD\nSiT/a3YoqYWtEtBC8wQvxpOeU6xRRd1wplQKPZI3YLItG3VlmE1dJqxCG5ldT8Yf\nvw8wrMIB0OjLvYOk43pwlrHUc4yI0Nv7V3cEo3x0D9090qJN3nZDBPatIwc2ggsi\nhpJRN1fE/QKBgEc6aUlJIAVQyI2ZrpONej9yPAaspzs10ovlTwK90eRKETXx2dTD\nqVagb6QRnwb+3J6Gyk8So8T5CvxaX/aP1lbFrj/DOPPn2p1a/T9RQgsJSAAH3qoD\nv9F9+SM+HcJn6MEQZe1+MC4GfPetumuIpqyEL0HuWudMsSvpXa4KUEmJAoGBALLj\nU8yWsg3N2A69e1gNmWXAyop4xgkclnnEBUv07tmrpRutTE/7TguyMHJ9kfHXQ/aU\nBVxd6prrKHffKlEUEFqp3aD9cFkSnCH2Mgqs8ICVDfBGxiuVVPTcDbm7SqtkHK0N\nYgnYY76OC5Bd/MsjhIulHSRmweS72l4S8Sf9Esu1AoGBAKW9MMXX2KYHKLBwyfGt\npy34mYr74gaGlGEHEar9EgCGLOtCC0ddz13t6RAfxEugrWZdpUu47/vJdVHqdg1R\ns1p5npQSRjuBLQ19dKsACpelAIZz1KFWMsSmXGxahy2BK+jY5j23Pzmco/KL24l6\nENuQjE6sQ/qsYfLzPTFlcmz2\n-----END PRIVATE KEY-----\n",
  "client_email": "gke-cluster-discovery@solo-public.iam.gserviceaccount.com",
  "client_id": "107244506411818627755",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/gke-cluster-discovery%40solo-public.iam.gserviceaccount.com"
}
`)
			opts := discovery.ClusterDiscoveryOpts{
				ProjectID:       "solo-public",
				CredentialsJSON: creds,
			}
			kubeConfigs, err := discovery.DiscoverGkeClusters(ctx, "write-namespace", opts)
			Expect(err).NotTo(HaveOccurred())
			kc, err := kubeConfigs.Find("write-namespace", "solo-kit-multicluster-1")
			Expect(err).NotTo(HaveOccurred())

			raw, err := clientcmd.Write(kc.Config)
			Expect(err).NotTo(HaveOccurred())
			restCfg, err := clientcmd.RESTConfigFromKubeConfig(raw)
			Expect(err).NotTo(HaveOccurred())

			kube, err := kubernetes.NewForConfig(restCfg)
			Expect(err).NotTo(HaveOccurred())
			svc, err := kube.CoreV1().Services("").List(v1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(svc).NotTo(HaveOccurred())
		})
	})
})
