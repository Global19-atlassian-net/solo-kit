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
	)

	BeforeEach(func() {
		srv = &server{}
	})

	It("can only register each path once", func() {
		Expect(srv.Register("hello", &MockHandler{})).NotTo(HaveOccurred())
		err := srv.Register("hello", &MockHandler{})
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(DuplicatePathError("hello")))
	})
	It("sets defaults properly", func() {
		Expect(srv.Register("hello", &MockHandler{})).NotTo(HaveOccurred())
		Expect(srv.Port).To(Equal(DefaultPort))
		Expect(srv.CertDir).To(Equal(DefaultCertPath))
		Expect(srv.Host).To(Equal(""))
		Expect(srv.webhooks).To(HaveLen(1))
	})
})
