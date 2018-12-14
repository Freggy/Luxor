package rest

import (
	"github.com/valyala/fasthttp"
	"strconv"
)

type Server struct {
	useHttps    bool
	port        int
	certFile    string
	keyFile     string
	handlerFunc fasthttp.RequestHandler
}

// NewServer creates a HTTP server running on the specified port.
func NewServer(port int) *Server {
	return &Server{
		false,
		port,
		"",
		"",
		nil,
	}
}

// NewServerTls creates a HTTPS server using the provided certificate and key on the given port.
func NewServerTls(port int, certFile string, keyFile string) *Server {
	return &Server{
		false,
		port,
		certFile,
		keyFile,
		nil,
	}
}

// SetHandler sets the handler that will be executed.
func (s *Server) SetHandler(handler fasthttp.RequestHandler) {
	s.handlerFunc = handler
}

// ListenAndServe starts the HTTP server and serves connections.
// if HTTPS is enabled this function will utilize it.
func (s *Server) ListenAndServe() error {
	if s.useHttps {
		return fasthttp.ListenAndServeTLS(":"+strconv.Itoa(s.port), s.certFile, s.keyFile, s.handlerFunc)
	} else {
		return fasthttp.ListenAndServe(":"+strconv.Itoa(s.port), s.handlerFunc)
	}
}
