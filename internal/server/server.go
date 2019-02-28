package server

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	// http2MinFrameSize https://tools.ietf.org/html/rfc7540#section-4.2
	http2MinFrameSize = 16 * 1024

	// Streams over this limit will be queued
	maxConcurrentStreams = 1000 // "infinite", per spec. 1000 seems good enough
)

// Config holds proxy server configuration
type Config struct {
	Addr      string
	Port      int
}

// Run executes the forward proxy server
func Run(cfg *Config) error {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		if host := req.Header.Get("H2-Host-Override"); host != "" {
			req.URL.Host = host
			req.Host = host
		}
		if cfg.Port > 0 {
			req.URL.Host = fmt.Sprintf("%s:%d", strings.Split(req.Host, ":")[0], cfg.Port)
		}
	}

	proxy := &httputil.ReverseProxy{
		Director:      director,
		Transport:     &h2cTransport{},
		FlushInterval: 50 * time.Millisecond,
	}

	handler := h2c.NewHandler(proxy, &http2.Server{
		MaxConcurrentStreams: maxConcurrentStreams,
		MaxReadFrameSize:     http2MinFrameSize,
	})

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: handler,
	}

	return server.ListenAndServe()
}

// h2cTransport implements http.Roundtripper creating a new http2.Transport on each call
type h2cTransport struct {
}

// RoundTrip implements http.Roundtripper.Roundtrip
func (h *h2cTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}
	return transport.RoundTrip(req)
}
