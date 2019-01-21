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

// Run executes the forward proxy server
func Run(addr string, port int) error {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		if port > 0 {
			req.URL.Host = fmt.Sprintf("%s:%d", strings.Split(req.Host, ":")[0], port)
		}
	}

	transport := &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}

	proxy := &httputil.ReverseProxy{
		Director:      director,
		Transport:     transport,
		FlushInterval: 50 * time.Millisecond,
	}

	handler := h2c.NewHandler(proxy, &http2.Server{
		MaxConcurrentStreams: maxConcurrentStreams,
		MaxReadFrameSize:     http2MinFrameSize,
	})

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return server.ListenAndServe()
}