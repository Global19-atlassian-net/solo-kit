package server

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/webhook/certwatcher"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
)

//go:generate mockgen -destination=./mocks/client_interface.go github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/webhook/server Webhook,KubeWebhook

const (
	certName = "tls.crt"
	keyName  = "tls.key"
)

// DefaultPort is the default port that the webhook server serves.
var (
	DefaultPort = 443

	DuplicatePathError = func(path string) error {
		return errors.Errorf("can't register duplicate path: %v", path)
	}
)

type KubeWebhook interface {
	http.Handler
	InjectScheme(s *runtime.Scheme) error
	Path() string
}

type Webhook interface {
	Register(path string, hook http.Handler) error
	Start(ctx context.Context) error
}

// Server is an admission webhook server that can serve traffic and
// generates related k8s resources for deploying.
type server struct {
	// Host is the address that the server will listen on.
	// Defaults to "" - all addresses.
	Host string

	// Port is the port number that the server will serve.
	// It will be defaulted to 443 if unspecified.
	Port int

	// CertDir is the directory that contains the server key and certificate.
	// If using FSCertWriter in Provisioner, the server itself will provision the certificate and
	// store it in this directory.
	// If using SecretCertWriter in Provisioner, the server will provision the certificate in a secret,
	// the user is responsible to mount the secret to the this location for the server to consume.
	CertDir string

	// WebhookMux is the multiplexer that handles different webhooks.
	WebhookMux *http.ServeMux

	// webhooks keep track of all registered webhooks for dependency injection,
	// and to provide better panic messages on duplicate webhook registration.
	webhooks map[string]http.Handler

	// defaultingOnce ensures that the default fields are only ever set once.
	defaultingOnce sync.Once
	mu             sync.Mutex
}

func NewWebhook() *server {
	return &server{}
}

// setDefaults does defaulting for the server.
func (s *server) setDefaults() {
	s.webhooks = map[string]http.Handler{}
	if s.WebhookMux == nil {
		s.WebhookMux = http.NewServeMux()
	}

	if s.Port <= 0 {
		s.Port = DefaultPort
	}

	if len(s.CertDir) == 0 {
		s.CertDir = path.Join("/tmp", "k8s-webhook-server", "serving-certs")
	}
}

// Register marks the given webhook as being served at the given path.
func (s *server) Register(path string, hook http.Handler) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.defaultingOnce.Do(s.setDefaults)
	_, found := s.webhooks[path]
	if found {
		return DuplicatePathError(path)
	}
	// TODO(directxman12): call setfields if we've already started the server
	s.webhooks[path] = hook
	s.WebhookMux.Handle(path, hook)
	return nil
}

// Start runs the server.
// It will install the webhook related resources depend on the server configuration.
func (s *server) Start(ctx context.Context) error {
	s.defaultingOnce.Do(s.setDefaults)

	baseHookLog := contextutils.LoggerFrom(ctx).With(zap.String("webhook", "server"))
	certPath := filepath.Join(s.CertDir, certName)
	keyPath := filepath.Join(s.CertDir, keyName)

	certWatcher, err := certwatcher.New(ctx, certPath, keyPath)
	if err != nil {
		return err
	}
	go func() {
		if err := certWatcher.Start(ctx); err != nil {
			baseHookLog.Error(err, "certificate watcher error")
		}
	}()

	cfg := &tls.Config{
		NextProtos:     []string{"h2"},
		GetCertificate: certWatcher.GetCertificate,
	}

	listener, err := tls.Listen("tcp", net.JoinHostPort(s.Host, strconv.Itoa(int(s.Port))), cfg)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Handler: s.WebhookMux,
	}

	err = srv.Serve(listener)
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	select {
	case <-ctx.Done():
		timeout, _ := context.WithTimeout(context.Background(), time.Second*10)
		if err := srv.Shutdown(timeout); err != nil {
			// Error from closing listeners, or context timeout
			baseHookLog.Error(err, "error shutting down the HTTP server")
		}
	}
	return nil
}
