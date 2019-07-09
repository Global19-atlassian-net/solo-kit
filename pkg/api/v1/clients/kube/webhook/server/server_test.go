package server

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockHandler struct {
}

func (m *MockHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("implement me")
}

var _ = Describe("", func() {
	var (
		srv *server
		// ctx context.Context
		// cancel context.CancelFunc
	)

	BeforeEach(func() {
		// ctx, cancel = context.WithCancel(context.Background())
		srv = &server{
			Port:       8080,
			webhooks:   map[string]http.Handler{},
			WebhookMux: http.NewServeMux(),
		}
	})

	It("can only register each path once", func() {
		Expect(srv.Register("hello", &MockHandler{})).NotTo(HaveOccurred())
		err := srv.Register("hello", &MockHandler{})
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(DuplicatePathError("hello")))
	})
	It("sets defaults properly", func() {

	})
})
