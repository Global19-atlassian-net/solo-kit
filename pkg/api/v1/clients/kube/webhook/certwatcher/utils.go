package certwatcher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"k8s.io/client-go/util/cert"
)

type Generator struct {
}

type certificates struct {
	// PEM-encoded CA certificate that has been used to sign the server certificate
	CaCertificate []byte
	// PEM-encoded server certificate
	ServerCertificate []byte
	// PEM-encoded private key that has been used to sign the server certificate
	ServerCertKey []byte
}

func GenerateSelfSignedPodCerts(metadata core.Metadata) (*certificates, error) {
	return generateSelfSignedCertificate(cert.Config{
		CommonName:   fmt.Sprintf("%s.%s", metadata.Name, metadata.Namespace),
		Organization: []string{"solo.io"},
		AltNames: cert.AltNames{
			DNSNames: []string{
				metadata.Name,
				fmt.Sprintf("%s.%s", metadata.Name, metadata.Namespace),
				fmt.Sprintf("%s.%s.svc", metadata.Name, metadata.Namespace),
				fmt.Sprintf("%s.%s.svc.cluster.local", metadata.Name, metadata.Namespace),
			},
		},
		Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	})
}

// This function generates a self-signed TLS certificate
func generateSelfSignedCertificate(config cert.Config) (*certificates, error) {

	// Generate the CA certificate that will be used to sign the webhook server certificate
	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create CA private key")
	}
	caCert, err := cert.NewSelfSignedCACert(cert.Config{CommonName: "supergloo-webhook-cert-ca"}, caPrivateKey)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create CA certificate")
	}

	// Generate webhook server certificate
	serverCertPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create server cert private key")
	}
	signedServerCert, err := cert.NewSignedCert(config, serverCertPrivateKey, caCert, caPrivateKey)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create server cert")
	}

	serverCertPrivateKeyPEM, err := cert.MarshalPrivateKeyToPEM(serverCertPrivateKey)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to convert server cert private key to PEM")
	}

	return &certificates{
		CaCertificate:     cert.EncodeCertPEM(caCert),
		ServerCertificate: cert.EncodeCertPEM(signedServerCert),
		ServerCertKey:     serverCertPrivateKeyPEM,
	}, nil
}
