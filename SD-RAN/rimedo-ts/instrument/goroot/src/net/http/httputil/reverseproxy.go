// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP reverse proxy handler

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:7
package httputil

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:7
import (
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:7
)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:7
import (
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:7
)

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/http/internal/ascii"
	"net/textproto"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/http/httpguts"
)

// A ProxyRequest contains a request to be rewritten by a ReverseProxy.
type ProxyRequest struct {
	// In is the request received by the proxy.
	// The Rewrite function must not modify In.
	In	*http.Request

	// Out is the request which will be sent by the proxy.
	// The Rewrite function may modify or replace this request.
	// Hop-by-hop headers are removed from this request
	// before Rewrite is called.
	Out	*http.Request
}

// SetURL routes the outbound request to the scheme, host, and base path
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
// provided in target. If the target's path is "/base" and the incoming
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
// request was for "/dir", the target request will be for "/base/dir".
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
//
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
// SetURL rewrites the outbound Host header to match the target's host.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
// To preserve the inbound request's Host header (the default behavior
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
// of NewSingleHostReverseProxy):
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
//
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
//	rewriteFunc := func(r *httputil.ProxyRequest) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
//		r.SetURL(url)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
//		r.Out.Host = r.In.Host
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:42
//	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:54
func (r *ProxyRequest) SetURL(target *url.URL) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:54
	_go_fuzz_dep_.CoverTab[76285]++
								rewriteRequestURL(r.Out, target)
								r.Out.Host = ""
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:56
	// _ = "end of CoverTab[76285]"
}

// SetXForwarded sets the X-Forwarded-For, X-Forwarded-Host, and
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
// X-Forwarded-Proto headers of the outbound request.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//   - The X-Forwarded-For header is set to the client IP address.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//   - The X-Forwarded-Host header is set to the host name requested
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//     by the client.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//   - The X-Forwarded-Proto header is set to "http" or "https", depending
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//     on whether the inbound request was made on a TLS-enabled connection.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
// If the outbound request contains an existing X-Forwarded-For header,
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
// SetXForwarded appends the client IP address to it. To append to the
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
// inbound request's X-Forwarded-For header (the default behavior of
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
// ReverseProxy when using a Director function), copy the header
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
// from the inbound request before calling SetXForwarded:
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//	rewriteFunc := func(r *httputil.ProxyRequest) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//		r.Out.Header["X-Forwarded-For"] = r.In.Header["X-Forwarded-For"]
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//		r.SetXForwarded()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:59
//	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:78
func (r *ProxyRequest) SetXForwarded() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:78
	_go_fuzz_dep_.CoverTab[76286]++
								clientIP, _, err := net.SplitHostPort(r.In.RemoteAddr)
								if err == nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:80
		_go_fuzz_dep_.CoverTab[76288]++
									prior := r.Out.Header["X-Forwarded-For"]
									if len(prior) > 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:82
			_go_fuzz_dep_.CoverTab[76290]++
										clientIP = strings.Join(prior, ", ") + ", " + clientIP
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:83
			// _ = "end of CoverTab[76290]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:84
			_go_fuzz_dep_.CoverTab[76291]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:84
			// _ = "end of CoverTab[76291]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:84
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:84
		// _ = "end of CoverTab[76288]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:84
		_go_fuzz_dep_.CoverTab[76289]++
									r.Out.Header.Set("X-Forwarded-For", clientIP)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:85
		// _ = "end of CoverTab[76289]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:86
		_go_fuzz_dep_.CoverTab[76292]++
									r.Out.Header.Del("X-Forwarded-For")
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:87
		// _ = "end of CoverTab[76292]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:88
	// _ = "end of CoverTab[76286]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:88
	_go_fuzz_dep_.CoverTab[76287]++
								r.Out.Header.Set("X-Forwarded-Host", r.In.Host)
								if r.In.TLS == nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:90
		_go_fuzz_dep_.CoverTab[76293]++
									r.Out.Header.Set("X-Forwarded-Proto", "http")
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:91
		// _ = "end of CoverTab[76293]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:92
		_go_fuzz_dep_.CoverTab[76294]++
									r.Out.Header.Set("X-Forwarded-Proto", "https")
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:93
		// _ = "end of CoverTab[76294]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:94
	// _ = "end of CoverTab[76287]"
}

// ReverseProxy is an HTTP Handler that takes an incoming request and
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:97
// sends it to another server, proxying the response back to the
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:97
// client.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:97
//
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:97
// 1xx responses are forwarded to the client if the underlying
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:97
// transport supports ClientTrace.Got1xxResponse.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:103
type ReverseProxy struct {
	// Rewrite must be a function which modifies
	// the request into a new request to be sent
	// using Transport. Its response is then copied
	// back to the original client unmodified.
	// Rewrite must not access the provided ProxyRequest
	// or its contents after returning.
	//
	// The Forwarded, X-Forwarded, X-Forwarded-Host,
	// and X-Forwarded-Proto headers are removed from the
	// outbound request before Rewrite is called. See also
	// the ProxyRequest.SetXForwarded method.
	//
	// Unparsable query parameters are removed from the
	// outbound request before Rewrite is called.
	// The Rewrite function may copy the inbound URL's
	// RawQuery to the outbound URL to preserve the original
	// parameter string. Note that this can lead to security
	// issues if the proxy's interpretation of query parameters
	// does not match that of the downstream server.
	//
	// At most one of Rewrite or Director may be set.
	Rewrite	func(*ProxyRequest)

	// Director is a function which modifies
	// the request into a new request to be sent
	// using Transport. Its response is then copied
	// back to the original client unmodified.
	// Director must not access the provided Request
	// after returning.
	//
	// By default, the X-Forwarded-For header is set to the
	// value of the client IP address. If an X-Forwarded-For
	// header already exists, the client IP is appended to the
	// existing values. As a special case, if the header
	// exists in the Request.Header map but has a nil value
	// (such as when set by the Director func), the X-Forwarded-For
	// header is not modified.
	//
	// To prevent IP spoofing, be sure to delete any pre-existing
	// X-Forwarded-For header coming from the client or
	// an untrusted proxy.
	//
	// Hop-by-hop headers are removed from the request after
	// Director returns, which can remove headers added by
	// Director. Use a Rewrite function instead to ensure
	// modifications to the request are preserved.
	//
	// Unparsable query parameters are removed from the outbound
	// request if Request.Form is set after Director returns.
	//
	// At most one of Rewrite or Director may be set.
	Director	func(*http.Request)

	// The transport used to perform proxy requests.
	// If nil, http.DefaultTransport is used.
	Transport	http.RoundTripper

	// FlushInterval specifies the flush interval
	// to flush to the client while copying the
	// response body.
	// If zero, no periodic flushing is done.
	// A negative value means to flush immediately
	// after each write to the client.
	// The FlushInterval is ignored when ReverseProxy
	// recognizes a response as a streaming response, or
	// if its ContentLength is -1; for such responses, writes
	// are flushed to the client immediately.
	FlushInterval	time.Duration

	// ErrorLog specifies an optional logger for errors
	// that occur when attempting to proxy the request.
	// If nil, logging is done via the log package's standard logger.
	ErrorLog	*log.Logger

	// BufferPool optionally specifies a buffer pool to
	// get byte slices for use by io.CopyBuffer when
	// copying HTTP response bodies.
	BufferPool	BufferPool

	// ModifyResponse is an optional function that modifies the
	// Response from the backend. It is called if the backend
	// returns a response at all, with any HTTP status code.
	// If the backend is unreachable, the optional ErrorHandler is
	// called without any call to ModifyResponse.
	//
	// If ModifyResponse returns an error, ErrorHandler is called
	// with its error value. If ErrorHandler is nil, its default
	// implementation is used.
	ModifyResponse	func(*http.Response) error

	// ErrorHandler is an optional function that handles errors
	// reaching the backend or errors from ModifyResponse.
	//
	// If nil, the default is to log the provided error and return
	// a 502 Status Bad Gateway response.
	ErrorHandler	func(http.ResponseWriter, *http.Request, error)
}

// A BufferPool is an interface for getting and returning temporary
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:202
// byte slices for use by io.CopyBuffer.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:204
type BufferPool interface {
	Get() []byte
	Put([]byte)
}

func singleJoiningSlash(a, b string) string {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:209
	_go_fuzz_dep_.CoverTab[76295]++
								aslash := strings.HasSuffix(a, "/")
								bslash := strings.HasPrefix(b, "/")
								switch {
	case aslash && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:213
		_go_fuzz_dep_.CoverTab[76300]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:213
		return bslash
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:213
		// _ = "end of CoverTab[76300]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:213
	}():
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:213
		_go_fuzz_dep_.CoverTab[76297]++
									return a + b[1:]
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:214
		// _ = "end of CoverTab[76297]"
	case !aslash && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:215
		_go_fuzz_dep_.CoverTab[76301]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:215
		return !bslash
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:215
		// _ = "end of CoverTab[76301]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:215
	}():
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:215
		_go_fuzz_dep_.CoverTab[76298]++
									return a + "/" + b
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:216
		// _ = "end of CoverTab[76298]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:216
	default:
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:216
		_go_fuzz_dep_.CoverTab[76299]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:216
		// _ = "end of CoverTab[76299]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:217
	// _ = "end of CoverTab[76295]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:217
	_go_fuzz_dep_.CoverTab[76296]++
								return a + b
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:218
	// _ = "end of CoverTab[76296]"
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:221
	_go_fuzz_dep_.CoverTab[76302]++
								if a.RawPath == "" && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:222
		_go_fuzz_dep_.CoverTab[76305]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:222
		return b.RawPath == ""
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:222
		// _ = "end of CoverTab[76305]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:222
	}() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:222
		_go_fuzz_dep_.CoverTab[76306]++
									return singleJoiningSlash(a.Path, b.Path), ""
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:223
		// _ = "end of CoverTab[76306]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:224
		_go_fuzz_dep_.CoverTab[76307]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:224
		// _ = "end of CoverTab[76307]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:224
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:224
	// _ = "end of CoverTab[76302]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:224
	_go_fuzz_dep_.CoverTab[76303]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:227
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:234
		_go_fuzz_dep_.CoverTab[76311]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:234
		return bslash
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:234
		// _ = "end of CoverTab[76311]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:234
	}():
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:234
		_go_fuzz_dep_.CoverTab[76308]++
									return a.Path + b.Path[1:], apath + bpath[1:]
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:235
		// _ = "end of CoverTab[76308]"
	case !aslash && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:236
		_go_fuzz_dep_.CoverTab[76312]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:236
		return !bslash
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:236
		// _ = "end of CoverTab[76312]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:236
	}():
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:236
		_go_fuzz_dep_.CoverTab[76309]++
									return a.Path + "/" + b.Path, apath + "/" + bpath
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:237
		// _ = "end of CoverTab[76309]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:237
	default:
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:237
		_go_fuzz_dep_.CoverTab[76310]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:237
		// _ = "end of CoverTab[76310]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:238
	// _ = "end of CoverTab[76303]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:238
	_go_fuzz_dep_.CoverTab[76304]++
								return a.Path + b.Path, apath + bpath
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:239
	// _ = "end of CoverTab[76304]"
}

// NewSingleHostReverseProxy returns a new ReverseProxy that routes
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// URLs to the scheme, host, and base path provided in target. If the
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// target's path is "/base" and the incoming request was for "/dir",
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// the target request will be for /base/dir.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
//
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// NewSingleHostReverseProxy does not rewrite the Host header.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
//
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// To customize the ReverseProxy behavior beyond what
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// NewSingleHostReverseProxy provides, use ReverseProxy directly
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// with a Rewrite function. The ProxyRequest SetURL method
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// may be used to route the outbound request. (Note that SetURL,
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// unlike NewSingleHostReverseProxy, rewrites the Host header
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
// of the outbound request by default.)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
//
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
//	proxy := &ReverseProxy{
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
//		Rewrite: func(r *ProxyRequest) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
//			r.SetURL(target)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
//			r.Out.Host = r.In.Host // if desired
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
//		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:242
//	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:262
func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:262
	_go_fuzz_dep_.CoverTab[76313]++
								director := func(req *http.Request) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:263
		_go_fuzz_dep_.CoverTab[76315]++
									rewriteRequestURL(req, target)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:264
		// _ = "end of CoverTab[76315]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:265
	// _ = "end of CoverTab[76313]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:265
	_go_fuzz_dep_.CoverTab[76314]++
								return &ReverseProxy{Director: director}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:266
	// _ = "end of CoverTab[76314]"
}

func rewriteRequestURL(req *http.Request, target *url.URL) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:269
	_go_fuzz_dep_.CoverTab[76316]++
								targetQuery := target.RawQuery
								req.URL.Scheme = target.Scheme
								req.URL.Host = target.Host
								req.URL.Path, req.URL.RawPath = joinURLPath(target, req.URL)
								if targetQuery == "" || func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:274
		_go_fuzz_dep_.CoverTab[76317]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:274
		return req.URL.RawQuery == ""
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:274
		// _ = "end of CoverTab[76317]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:274
	}() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:274
		_go_fuzz_dep_.CoverTab[76318]++
									req.URL.RawQuery = targetQuery + req.URL.RawQuery
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:275
		// _ = "end of CoverTab[76318]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:276
		_go_fuzz_dep_.CoverTab[76319]++
									req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:277
		// _ = "end of CoverTab[76319]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:278
	// _ = "end of CoverTab[76316]"
}

func copyHeader(dst, src http.Header) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:281
	_go_fuzz_dep_.CoverTab[76320]++
								for k, vv := range src {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:282
		_go_fuzz_dep_.CoverTab[76321]++
									for _, v := range vv {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:283
			_go_fuzz_dep_.CoverTab[76322]++
										dst.Add(k, v)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:284
			// _ = "end of CoverTab[76322]"
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:285
		// _ = "end of CoverTab[76321]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:286
	// _ = "end of CoverTab[76320]"
}

// Hop-by-hop headers. These are removed when sent to the backend.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:289
// As of RFC 7230, hop-by-hop headers are required to appear in the
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:289
// Connection header field. These are the headers defined by the
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:289
// obsoleted RFC 2616 (section 13.5.1) and are used for backward
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:289
// compatibility.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:294
var hopHeaders = []string{
	"Connection",
	"Proxy-Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailer",
	"Transfer-Encoding",
	"Upgrade",
}

func (p *ReverseProxy) defaultErrorHandler(rw http.ResponseWriter, req *http.Request, err error) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:306
	_go_fuzz_dep_.CoverTab[76323]++
								p.logf("http: proxy error: %v", err)
								rw.WriteHeader(http.StatusBadGateway)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:308
	// _ = "end of CoverTab[76323]"
}

func (p *ReverseProxy) getErrorHandler() func(http.ResponseWriter, *http.Request, error) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:311
	_go_fuzz_dep_.CoverTab[76324]++
								if p.ErrorHandler != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:312
		_go_fuzz_dep_.CoverTab[76326]++
									return p.ErrorHandler
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:313
		// _ = "end of CoverTab[76326]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:314
		_go_fuzz_dep_.CoverTab[76327]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:314
		// _ = "end of CoverTab[76327]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:314
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:314
	// _ = "end of CoverTab[76324]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:314
	_go_fuzz_dep_.CoverTab[76325]++
								return p.defaultErrorHandler
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:315
	// _ = "end of CoverTab[76325]"
}

// modifyResponse conditionally runs the optional ModifyResponse hook
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:318
// and reports whether the request should proceed.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:320
func (p *ReverseProxy) modifyResponse(rw http.ResponseWriter, res *http.Response, req *http.Request) bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:320
	_go_fuzz_dep_.CoverTab[76328]++
								if p.ModifyResponse == nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:321
		_go_fuzz_dep_.CoverTab[76331]++
									return true
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:322
		// _ = "end of CoverTab[76331]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:323
		_go_fuzz_dep_.CoverTab[76332]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:323
		// _ = "end of CoverTab[76332]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:323
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:323
	// _ = "end of CoverTab[76328]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:323
	_go_fuzz_dep_.CoverTab[76329]++
								if err := p.ModifyResponse(res); err != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:324
		_go_fuzz_dep_.CoverTab[76333]++
									res.Body.Close()
									p.getErrorHandler()(rw, req, err)
									return false
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:327
		// _ = "end of CoverTab[76333]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:328
		_go_fuzz_dep_.CoverTab[76334]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:328
		// _ = "end of CoverTab[76334]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:328
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:328
	// _ = "end of CoverTab[76329]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:328
	_go_fuzz_dep_.CoverTab[76330]++
								return true
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:329
	// _ = "end of CoverTab[76330]"
}

func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:332
	_go_fuzz_dep_.CoverTab[76335]++
								transport := p.Transport
								if transport == nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:334
		_go_fuzz_dep_.CoverTab[76356]++
									transport = http.DefaultTransport
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:335
		// _ = "end of CoverTab[76356]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:336
		_go_fuzz_dep_.CoverTab[76357]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:336
		// _ = "end of CoverTab[76357]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:336
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:336
	// _ = "end of CoverTab[76335]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:336
	_go_fuzz_dep_.CoverTab[76336]++

								ctx := req.Context()
								if ctx.Done() != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:339
		_go_fuzz_dep_.CoverTab[76358]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:339
		// _ = "end of CoverTab[76358]"

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:350
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:350
		_go_fuzz_dep_.CoverTab[76359]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:350
		if cn, ok := rw.(http.CloseNotifier); ok {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:350
			_go_fuzz_dep_.CoverTab[76360]++
										var cancel context.CancelFunc
										ctx, cancel = context.WithCancel(ctx)
										defer cancel()
										notifyChan := cn.CloseNotify()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:354
			_curRoutineNum72_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:354
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum72_)
										go func() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:355
				_go_fuzz_dep_.CoverTab[76361]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:355
				defer func() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:355
					_go_fuzz_dep_.CoverTab[76362]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:355
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum72_)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:355
					// _ = "end of CoverTab[76362]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:355
				}()
											select {
				case <-notifyChan:
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:357
					_go_fuzz_dep_.CoverTab[76363]++
												cancel()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:358
					// _ = "end of CoverTab[76363]"
				case <-ctx.Done():
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:359
					_go_fuzz_dep_.CoverTab[76364]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:359
					// _ = "end of CoverTab[76364]"
				}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:360
				// _ = "end of CoverTab[76361]"
			}()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:361
			// _ = "end of CoverTab[76360]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:362
			_go_fuzz_dep_.CoverTab[76365]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:362
			// _ = "end of CoverTab[76365]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:362
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:362
		// _ = "end of CoverTab[76359]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:362
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:362
	// _ = "end of CoverTab[76336]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:362
	_go_fuzz_dep_.CoverTab[76337]++

								outreq := req.Clone(ctx)
								if req.ContentLength == 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:365
		_go_fuzz_dep_.CoverTab[76366]++
									outreq.Body = nil
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:366
		// _ = "end of CoverTab[76366]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:367
		_go_fuzz_dep_.CoverTab[76367]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:367
		// _ = "end of CoverTab[76367]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:367
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:367
	// _ = "end of CoverTab[76337]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:367
	_go_fuzz_dep_.CoverTab[76338]++
								if outreq.Body != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:368
		_go_fuzz_dep_.CoverTab[76368]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:375
		defer outreq.Body.Close()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:375
		// _ = "end of CoverTab[76368]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:376
		_go_fuzz_dep_.CoverTab[76369]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:376
		// _ = "end of CoverTab[76369]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:376
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:376
	// _ = "end of CoverTab[76338]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:376
	_go_fuzz_dep_.CoverTab[76339]++
								if outreq.Header == nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:377
		_go_fuzz_dep_.CoverTab[76370]++
									outreq.Header = make(http.Header)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:378
		// _ = "end of CoverTab[76370]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:379
		_go_fuzz_dep_.CoverTab[76371]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:379
		// _ = "end of CoverTab[76371]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:379
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:379
	// _ = "end of CoverTab[76339]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:379
	_go_fuzz_dep_.CoverTab[76340]++

								if (p.Director != nil) == (p.Rewrite != nil) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:381
		_go_fuzz_dep_.CoverTab[76372]++
									p.getErrorHandler()(rw, req, errors.New("ReverseProxy must have exactly one of Director or Rewrite set"))
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:383
		// _ = "end of CoverTab[76372]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:384
		_go_fuzz_dep_.CoverTab[76373]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:384
		// _ = "end of CoverTab[76373]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:384
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:384
	// _ = "end of CoverTab[76340]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:384
	_go_fuzz_dep_.CoverTab[76341]++

								if p.Director != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:386
		_go_fuzz_dep_.CoverTab[76374]++
									p.Director(outreq)
									if outreq.Form != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:388
			_go_fuzz_dep_.CoverTab[76375]++
										outreq.URL.RawQuery = cleanQueryParams(outreq.URL.RawQuery)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:389
			// _ = "end of CoverTab[76375]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:390
			_go_fuzz_dep_.CoverTab[76376]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:390
			// _ = "end of CoverTab[76376]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:390
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:390
		// _ = "end of CoverTab[76374]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:391
		_go_fuzz_dep_.CoverTab[76377]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:391
		// _ = "end of CoverTab[76377]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:391
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:391
	// _ = "end of CoverTab[76341]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:391
	_go_fuzz_dep_.CoverTab[76342]++
								outreq.Close = false

								reqUpType := upgradeType(outreq.Header)
								if !ascii.IsPrint(reqUpType) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:395
		_go_fuzz_dep_.CoverTab[76378]++
									p.getErrorHandler()(rw, req, fmt.Errorf("client tried to switch to invalid protocol %q", reqUpType))
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:397
		// _ = "end of CoverTab[76378]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:398
		_go_fuzz_dep_.CoverTab[76379]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:398
		// _ = "end of CoverTab[76379]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:398
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:398
	// _ = "end of CoverTab[76342]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:398
	_go_fuzz_dep_.CoverTab[76343]++
								removeHopByHopHeaders(outreq.Header)

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:406
	if httpguts.HeaderValuesContainsToken(req.Header["Te"], "trailers") {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:406
		_go_fuzz_dep_.CoverTab[76380]++
									outreq.Header.Set("Te", "trailers")
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:407
		// _ = "end of CoverTab[76380]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:408
		_go_fuzz_dep_.CoverTab[76381]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:408
		// _ = "end of CoverTab[76381]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:408
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:408
	// _ = "end of CoverTab[76343]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:408
	_go_fuzz_dep_.CoverTab[76344]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:412
	if reqUpType != "" {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:412
		_go_fuzz_dep_.CoverTab[76382]++
									outreq.Header.Set("Connection", "Upgrade")
									outreq.Header.Set("Upgrade", reqUpType)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:414
		// _ = "end of CoverTab[76382]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:415
		_go_fuzz_dep_.CoverTab[76383]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:415
		// _ = "end of CoverTab[76383]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:415
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:415
	// _ = "end of CoverTab[76344]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:415
	_go_fuzz_dep_.CoverTab[76345]++

								if p.Rewrite != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:417
		_go_fuzz_dep_.CoverTab[76384]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:421
		outreq.Header.Del("Forwarded")
									outreq.Header.Del("X-Forwarded-For")
									outreq.Header.Del("X-Forwarded-Host")
									outreq.Header.Del("X-Forwarded-Proto")

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:427
		outreq.URL.RawQuery = cleanQueryParams(outreq.URL.RawQuery)

		pr := &ProxyRequest{
			In:	req,
			Out:	outreq,
		}
									p.Rewrite(pr)
									outreq = pr.Out
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:434
		// _ = "end of CoverTab[76384]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:435
		_go_fuzz_dep_.CoverTab[76385]++
									if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:436
			_go_fuzz_dep_.CoverTab[76386]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:440
			prior, ok := outreq.Header["X-Forwarded-For"]
			omit := ok && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:441
				_go_fuzz_dep_.CoverTab[76388]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:441
				return prior == nil
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:441
				// _ = "end of CoverTab[76388]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:441
			}()
										if len(prior) > 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:442
				_go_fuzz_dep_.CoverTab[76389]++
											clientIP = strings.Join(prior, ", ") + ", " + clientIP
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:443
				// _ = "end of CoverTab[76389]"
			} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:444
				_go_fuzz_dep_.CoverTab[76390]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:444
				// _ = "end of CoverTab[76390]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:444
			}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:444
			// _ = "end of CoverTab[76386]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:444
			_go_fuzz_dep_.CoverTab[76387]++
										if !omit {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:445
				_go_fuzz_dep_.CoverTab[76391]++
											outreq.Header.Set("X-Forwarded-For", clientIP)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:446
				// _ = "end of CoverTab[76391]"
			} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:447
				_go_fuzz_dep_.CoverTab[76392]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:447
				// _ = "end of CoverTab[76392]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:447
			}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:447
			// _ = "end of CoverTab[76387]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:448
			_go_fuzz_dep_.CoverTab[76393]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:448
			// _ = "end of CoverTab[76393]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:448
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:448
		// _ = "end of CoverTab[76385]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:449
	// _ = "end of CoverTab[76345]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:449
	_go_fuzz_dep_.CoverTab[76346]++

								if _, ok := outreq.Header["User-Agent"]; !ok {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:451
		_go_fuzz_dep_.CoverTab[76394]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:454
		outreq.Header.Set("User-Agent", "")
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:454
		// _ = "end of CoverTab[76394]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:455
		_go_fuzz_dep_.CoverTab[76395]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:455
		// _ = "end of CoverTab[76395]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:455
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:455
	// _ = "end of CoverTab[76346]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:455
	_go_fuzz_dep_.CoverTab[76347]++

								trace := &httptrace.ClientTrace{
		Got1xxResponse: func(code int, header textproto.MIMEHeader) error {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:458
			_go_fuzz_dep_.CoverTab[76396]++
										h := rw.Header()
										copyHeader(h, http.Header(header))
										rw.WriteHeader(code)

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:464
			for k := range h {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:464
				_go_fuzz_dep_.CoverTab[76398]++
											delete(h, k)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:465
				// _ = "end of CoverTab[76398]"
			}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:466
			// _ = "end of CoverTab[76396]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:466
			_go_fuzz_dep_.CoverTab[76397]++

										return nil
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:468
			// _ = "end of CoverTab[76397]"
		},
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:470
	// _ = "end of CoverTab[76347]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:470
	_go_fuzz_dep_.CoverTab[76348]++
								outreq = outreq.WithContext(httptrace.WithClientTrace(outreq.Context(), trace))

								res, err := transport.RoundTrip(outreq)
								if err != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:474
		_go_fuzz_dep_.CoverTab[76399]++
									p.getErrorHandler()(rw, outreq, err)
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:476
		// _ = "end of CoverTab[76399]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:477
		_go_fuzz_dep_.CoverTab[76400]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:477
		// _ = "end of CoverTab[76400]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:477
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:477
	// _ = "end of CoverTab[76348]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:477
	_go_fuzz_dep_.CoverTab[76349]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:480
	if res.StatusCode == http.StatusSwitchingProtocols {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:480
		_go_fuzz_dep_.CoverTab[76401]++
									if !p.modifyResponse(rw, res, outreq) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:481
			_go_fuzz_dep_.CoverTab[76403]++
										return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:482
			// _ = "end of CoverTab[76403]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:483
			_go_fuzz_dep_.CoverTab[76404]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:483
			// _ = "end of CoverTab[76404]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:483
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:483
		// _ = "end of CoverTab[76401]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:483
		_go_fuzz_dep_.CoverTab[76402]++
									p.handleUpgradeResponse(rw, outreq, res)
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:485
		// _ = "end of CoverTab[76402]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:486
		_go_fuzz_dep_.CoverTab[76405]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:486
		// _ = "end of CoverTab[76405]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:486
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:486
	// _ = "end of CoverTab[76349]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:486
	_go_fuzz_dep_.CoverTab[76350]++

								removeHopByHopHeaders(res.Header)

								if !p.modifyResponse(rw, res, outreq) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:490
		_go_fuzz_dep_.CoverTab[76406]++
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:491
		// _ = "end of CoverTab[76406]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:492
		_go_fuzz_dep_.CoverTab[76407]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:492
		// _ = "end of CoverTab[76407]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:492
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:492
	// _ = "end of CoverTab[76350]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:492
	_go_fuzz_dep_.CoverTab[76351]++

								copyHeader(rw.Header(), res.Header)

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:498
	announcedTrailers := len(res.Trailer)
	if announcedTrailers > 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:499
		_go_fuzz_dep_.CoverTab[76408]++
									trailerKeys := make([]string, 0, len(res.Trailer))
									for k := range res.Trailer {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:501
			_go_fuzz_dep_.CoverTab[76410]++
										trailerKeys = append(trailerKeys, k)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:502
			// _ = "end of CoverTab[76410]"
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:503
		// _ = "end of CoverTab[76408]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:503
		_go_fuzz_dep_.CoverTab[76409]++
									rw.Header().Add("Trailer", strings.Join(trailerKeys, ", "))
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:504
		// _ = "end of CoverTab[76409]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:505
		_go_fuzz_dep_.CoverTab[76411]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:505
		// _ = "end of CoverTab[76411]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:505
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:505
	// _ = "end of CoverTab[76351]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:505
	_go_fuzz_dep_.CoverTab[76352]++

								rw.WriteHeader(res.StatusCode)

								err = p.copyResponse(rw, res.Body, p.flushInterval(res))
								if err != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:510
		_go_fuzz_dep_.CoverTab[76412]++
									defer res.Body.Close()

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:515
		if !shouldPanicOnCopyError(req) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:515
			_go_fuzz_dep_.CoverTab[76414]++
										p.logf("suppressing panic for copyResponse error in test; copy error: %v", err)
										return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:517
			// _ = "end of CoverTab[76414]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:518
			_go_fuzz_dep_.CoverTab[76415]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:518
			// _ = "end of CoverTab[76415]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:518
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:518
		// _ = "end of CoverTab[76412]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:518
		_go_fuzz_dep_.CoverTab[76413]++
									panic(http.ErrAbortHandler)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:519
		// _ = "end of CoverTab[76413]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:520
		_go_fuzz_dep_.CoverTab[76416]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:520
		// _ = "end of CoverTab[76416]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:520
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:520
	// _ = "end of CoverTab[76352]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:520
	_go_fuzz_dep_.CoverTab[76353]++
								res.Body.Close()

								if len(res.Trailer) > 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:523
		_go_fuzz_dep_.CoverTab[76417]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:527
		if fl, ok := rw.(http.Flusher); ok {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:527
			_go_fuzz_dep_.CoverTab[76418]++
										fl.Flush()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:528
			// _ = "end of CoverTab[76418]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:529
			_go_fuzz_dep_.CoverTab[76419]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:529
			// _ = "end of CoverTab[76419]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:529
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:529
		// _ = "end of CoverTab[76417]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:530
		_go_fuzz_dep_.CoverTab[76420]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:530
		// _ = "end of CoverTab[76420]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:530
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:530
	// _ = "end of CoverTab[76353]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:530
	_go_fuzz_dep_.CoverTab[76354]++

								if len(res.Trailer) == announcedTrailers {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:532
		_go_fuzz_dep_.CoverTab[76421]++
									copyHeader(rw.Header(), res.Trailer)
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:534
		// _ = "end of CoverTab[76421]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:535
		_go_fuzz_dep_.CoverTab[76422]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:535
		// _ = "end of CoverTab[76422]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:535
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:535
	// _ = "end of CoverTab[76354]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:535
	_go_fuzz_dep_.CoverTab[76355]++

								for k, vv := range res.Trailer {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:537
		_go_fuzz_dep_.CoverTab[76423]++
									k = http.TrailerPrefix + k
									for _, v := range vv {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:539
			_go_fuzz_dep_.CoverTab[76424]++
										rw.Header().Add(k, v)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:540
			// _ = "end of CoverTab[76424]"
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:541
		// _ = "end of CoverTab[76423]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:542
	// _ = "end of CoverTab[76355]"
}

var inOurTests bool	// whether we're in our own tests

// shouldPanicOnCopyError reports whether the reverse proxy should
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:547
// panic with http.ErrAbortHandler. This is the right thing to do by
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:547
// default, but Go 1.10 and earlier did not, so existing unit tests
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:547
// weren't expecting panics. Only panic in our own tests, or when
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:547
// running under the HTTP server.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:552
func shouldPanicOnCopyError(req *http.Request) bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:552
	_go_fuzz_dep_.CoverTab[76425]++
								if inOurTests {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:553
		_go_fuzz_dep_.CoverTab[76428]++

									return true
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:555
		// _ = "end of CoverTab[76428]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:556
		_go_fuzz_dep_.CoverTab[76429]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:556
		// _ = "end of CoverTab[76429]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:556
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:556
	// _ = "end of CoverTab[76425]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:556
	_go_fuzz_dep_.CoverTab[76426]++
								if req.Context().Value(http.ServerContextKey) != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:557
		_go_fuzz_dep_.CoverTab[76430]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:560
		return true
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:560
		// _ = "end of CoverTab[76430]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:561
		_go_fuzz_dep_.CoverTab[76431]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:561
		// _ = "end of CoverTab[76431]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:561
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:561
	// _ = "end of CoverTab[76426]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:561
	_go_fuzz_dep_.CoverTab[76427]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:564
	return false
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:564
	// _ = "end of CoverTab[76427]"
}

// removeHopByHopHeaders removes hop-by-hop headers.
func removeHopByHopHeaders(h http.Header) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:568
	_go_fuzz_dep_.CoverTab[76432]++

								for _, f := range h["Connection"] {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:570
		_go_fuzz_dep_.CoverTab[76434]++
									for _, sf := range strings.Split(f, ",") {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:571
			_go_fuzz_dep_.CoverTab[76435]++
										if sf = textproto.TrimString(sf); sf != "" {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:572
				_go_fuzz_dep_.CoverTab[76436]++
											h.Del(sf)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:573
				// _ = "end of CoverTab[76436]"
			} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:574
				_go_fuzz_dep_.CoverTab[76437]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:574
				// _ = "end of CoverTab[76437]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:574
			}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:574
			// _ = "end of CoverTab[76435]"
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:575
		// _ = "end of CoverTab[76434]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:576
	// _ = "end of CoverTab[76432]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:576
	_go_fuzz_dep_.CoverTab[76433]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:580
	for _, f := range hopHeaders {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:580
		_go_fuzz_dep_.CoverTab[76438]++
									h.Del(f)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:581
		// _ = "end of CoverTab[76438]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:582
	// _ = "end of CoverTab[76433]"
}

// flushInterval returns the p.FlushInterval value, conditionally
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:585
// overriding its value for a specific request/response.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:587
func (p *ReverseProxy) flushInterval(res *http.Response) time.Duration {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:587
	_go_fuzz_dep_.CoverTab[76439]++
								resCT := res.Header.Get("Content-Type")

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:592
	if baseCT, _, _ := mime.ParseMediaType(resCT); baseCT == "text/event-stream" {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:592
		_go_fuzz_dep_.CoverTab[76442]++
									return -1
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:593
		// _ = "end of CoverTab[76442]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:594
		_go_fuzz_dep_.CoverTab[76443]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:594
		// _ = "end of CoverTab[76443]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:594
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:594
	// _ = "end of CoverTab[76439]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:594
	_go_fuzz_dep_.CoverTab[76440]++

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:597
	if res.ContentLength == -1 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:597
		_go_fuzz_dep_.CoverTab[76444]++
									return -1
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:598
		// _ = "end of CoverTab[76444]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:599
		_go_fuzz_dep_.CoverTab[76445]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:599
		// _ = "end of CoverTab[76445]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:599
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:599
	// _ = "end of CoverTab[76440]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:599
	_go_fuzz_dep_.CoverTab[76441]++

								return p.FlushInterval
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:601
	// _ = "end of CoverTab[76441]"
}

func (p *ReverseProxy) copyResponse(dst io.Writer, src io.Reader, flushInterval time.Duration) error {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:604
	_go_fuzz_dep_.CoverTab[76446]++
								if flushInterval != 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:605
		_go_fuzz_dep_.CoverTab[76449]++
									if wf, ok := dst.(writeFlusher); ok {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:606
			_go_fuzz_dep_.CoverTab[76450]++
										mlw := &maxLatencyWriter{
				dst:		wf,
				latency:	flushInterval,
			}
										defer mlw.stop()

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:614
			mlw.flushPending = true
										mlw.t = time.AfterFunc(flushInterval, mlw.delayedFlush)

										dst = mlw
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:617
			// _ = "end of CoverTab[76450]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:618
			_go_fuzz_dep_.CoverTab[76451]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:618
			// _ = "end of CoverTab[76451]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:618
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:618
		// _ = "end of CoverTab[76449]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:619
		_go_fuzz_dep_.CoverTab[76452]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:619
		// _ = "end of CoverTab[76452]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:619
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:619
	// _ = "end of CoverTab[76446]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:619
	_go_fuzz_dep_.CoverTab[76447]++

								var buf []byte
								if p.BufferPool != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:622
		_go_fuzz_dep_.CoverTab[76453]++
									buf = p.BufferPool.Get()
									defer p.BufferPool.Put(buf)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:624
		// _ = "end of CoverTab[76453]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:625
		_go_fuzz_dep_.CoverTab[76454]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:625
		// _ = "end of CoverTab[76454]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:625
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:625
	// _ = "end of CoverTab[76447]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:625
	_go_fuzz_dep_.CoverTab[76448]++
								_, err := p.copyBuffer(dst, src, buf)
								return err
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:627
	// _ = "end of CoverTab[76448]"
}

// copyBuffer returns any write errors or non-EOF read errors, and the amount
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:630
// of bytes written.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:632
func (p *ReverseProxy) copyBuffer(dst io.Writer, src io.Reader, buf []byte) (int64, error) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:632
	_go_fuzz_dep_.CoverTab[76455]++
								if len(buf) == 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:633
		_go_fuzz_dep_.CoverTab[76457]++
									buf = make([]byte, 32*1024)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:634
		// _ = "end of CoverTab[76457]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:635
		_go_fuzz_dep_.CoverTab[76458]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:635
		// _ = "end of CoverTab[76458]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:635
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:635
	// _ = "end of CoverTab[76455]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:635
	_go_fuzz_dep_.CoverTab[76456]++
								var written int64
								for {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:637
		_go_fuzz_dep_.CoverTab[76459]++
									nr, rerr := src.Read(buf)
									if rerr != nil && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:639
			_go_fuzz_dep_.CoverTab[76462]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:639
			return rerr != io.EOF
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:639
			// _ = "end of CoverTab[76462]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:639
		}() && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:639
			_go_fuzz_dep_.CoverTab[76463]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:639
			return rerr != context.Canceled
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:639
			// _ = "end of CoverTab[76463]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:639
		}() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:639
			_go_fuzz_dep_.CoverTab[76464]++
										p.logf("httputil: ReverseProxy read error during body copy: %v", rerr)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:640
			// _ = "end of CoverTab[76464]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:641
			_go_fuzz_dep_.CoverTab[76465]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:641
			// _ = "end of CoverTab[76465]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:641
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:641
		// _ = "end of CoverTab[76459]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:641
		_go_fuzz_dep_.CoverTab[76460]++
									if nr > 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:642
			_go_fuzz_dep_.CoverTab[76466]++
										nw, werr := dst.Write(buf[:nr])
										if nw > 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:644
				_go_fuzz_dep_.CoverTab[76469]++
											written += int64(nw)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:645
				// _ = "end of CoverTab[76469]"
			} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:646
				_go_fuzz_dep_.CoverTab[76470]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:646
				// _ = "end of CoverTab[76470]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:646
			}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:646
			// _ = "end of CoverTab[76466]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:646
			_go_fuzz_dep_.CoverTab[76467]++
										if werr != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:647
				_go_fuzz_dep_.CoverTab[76471]++
											return written, werr
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:648
				// _ = "end of CoverTab[76471]"
			} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:649
				_go_fuzz_dep_.CoverTab[76472]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:649
				// _ = "end of CoverTab[76472]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:649
			}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:649
			// _ = "end of CoverTab[76467]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:649
			_go_fuzz_dep_.CoverTab[76468]++
										if nr != nw {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:650
				_go_fuzz_dep_.CoverTab[76473]++
											return written, io.ErrShortWrite
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:651
				// _ = "end of CoverTab[76473]"
			} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:652
				_go_fuzz_dep_.CoverTab[76474]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:652
				// _ = "end of CoverTab[76474]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:652
			}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:652
			// _ = "end of CoverTab[76468]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:653
			_go_fuzz_dep_.CoverTab[76475]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:653
			// _ = "end of CoverTab[76475]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:653
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:653
		// _ = "end of CoverTab[76460]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:653
		_go_fuzz_dep_.CoverTab[76461]++
									if rerr != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:654
			_go_fuzz_dep_.CoverTab[76476]++
										if rerr == io.EOF {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:655
				_go_fuzz_dep_.CoverTab[76478]++
											rerr = nil
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:656
				// _ = "end of CoverTab[76478]"
			} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:657
				_go_fuzz_dep_.CoverTab[76479]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:657
				// _ = "end of CoverTab[76479]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:657
			}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:657
			// _ = "end of CoverTab[76476]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:657
			_go_fuzz_dep_.CoverTab[76477]++
										return written, rerr
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:658
			// _ = "end of CoverTab[76477]"
		} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:659
			_go_fuzz_dep_.CoverTab[76480]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:659
			// _ = "end of CoverTab[76480]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:659
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:659
		// _ = "end of CoverTab[76461]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:660
	// _ = "end of CoverTab[76456]"
}

func (p *ReverseProxy) logf(format string, args ...any) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:663
	_go_fuzz_dep_.CoverTab[76481]++
								if p.ErrorLog != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:664
		_go_fuzz_dep_.CoverTab[76482]++
									p.ErrorLog.Printf(format, args...)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:665
		// _ = "end of CoverTab[76482]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:666
		_go_fuzz_dep_.CoverTab[76483]++
									log.Printf(format, args...)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:667
		// _ = "end of CoverTab[76483]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:668
	// _ = "end of CoverTab[76481]"
}

type writeFlusher interface {
	io.Writer
	http.Flusher
}

type maxLatencyWriter struct {
	dst	writeFlusher
	latency	time.Duration	// non-zero; negative means to flush immediately

	mu		sync.Mutex	// protects t, flushPending, and dst.Flush
	t		*time.Timer
	flushPending	bool
}

func (m *maxLatencyWriter) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:685
	_go_fuzz_dep_.CoverTab[76484]++
								m.mu.Lock()
								defer m.mu.Unlock()
								n, err = m.dst.Write(p)
								if m.latency < 0 {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:689
		_go_fuzz_dep_.CoverTab[76488]++
									m.dst.Flush()
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:691
		// _ = "end of CoverTab[76488]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:692
		_go_fuzz_dep_.CoverTab[76489]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:692
		// _ = "end of CoverTab[76489]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:692
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:692
	// _ = "end of CoverTab[76484]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:692
	_go_fuzz_dep_.CoverTab[76485]++
								if m.flushPending {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:693
		_go_fuzz_dep_.CoverTab[76490]++
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:694
		// _ = "end of CoverTab[76490]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:695
		_go_fuzz_dep_.CoverTab[76491]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:695
		// _ = "end of CoverTab[76491]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:695
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:695
	// _ = "end of CoverTab[76485]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:695
	_go_fuzz_dep_.CoverTab[76486]++
								if m.t == nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:696
		_go_fuzz_dep_.CoverTab[76492]++
									m.t = time.AfterFunc(m.latency, m.delayedFlush)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:697
		// _ = "end of CoverTab[76492]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:698
		_go_fuzz_dep_.CoverTab[76493]++
									m.t.Reset(m.latency)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:699
		// _ = "end of CoverTab[76493]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:700
	// _ = "end of CoverTab[76486]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:700
	_go_fuzz_dep_.CoverTab[76487]++
								m.flushPending = true
								return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:702
	// _ = "end of CoverTab[76487]"
}

func (m *maxLatencyWriter) delayedFlush() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:705
	_go_fuzz_dep_.CoverTab[76494]++
								m.mu.Lock()
								defer m.mu.Unlock()
								if !m.flushPending {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:708
		_go_fuzz_dep_.CoverTab[76496]++
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:709
		// _ = "end of CoverTab[76496]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:710
		_go_fuzz_dep_.CoverTab[76497]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:710
		// _ = "end of CoverTab[76497]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:710
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:710
	// _ = "end of CoverTab[76494]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:710
	_go_fuzz_dep_.CoverTab[76495]++
								m.dst.Flush()
								m.flushPending = false
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:712
	// _ = "end of CoverTab[76495]"
}

func (m *maxLatencyWriter) stop() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:715
	_go_fuzz_dep_.CoverTab[76498]++
								m.mu.Lock()
								defer m.mu.Unlock()
								m.flushPending = false
								if m.t != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:719
		_go_fuzz_dep_.CoverTab[76499]++
									m.t.Stop()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:720
		// _ = "end of CoverTab[76499]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:721
		_go_fuzz_dep_.CoverTab[76500]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:721
		// _ = "end of CoverTab[76500]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:721
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:721
	// _ = "end of CoverTab[76498]"
}

func upgradeType(h http.Header) string {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:724
	_go_fuzz_dep_.CoverTab[76501]++
								if !httpguts.HeaderValuesContainsToken(h["Connection"], "Upgrade") {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:725
		_go_fuzz_dep_.CoverTab[76503]++
									return ""
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:726
		// _ = "end of CoverTab[76503]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:727
		_go_fuzz_dep_.CoverTab[76504]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:727
		// _ = "end of CoverTab[76504]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:727
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:727
	// _ = "end of CoverTab[76501]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:727
	_go_fuzz_dep_.CoverTab[76502]++
								return h.Get("Upgrade")
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:728
	// _ = "end of CoverTab[76502]"
}

func (p *ReverseProxy) handleUpgradeResponse(rw http.ResponseWriter, req *http.Request, res *http.Response) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:731
	_go_fuzz_dep_.CoverTab[76505]++
								reqUpType := upgradeType(req.Header)
								resUpType := upgradeType(res.Header)
								if !ascii.IsPrint(resUpType) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:734
		_go_fuzz_dep_.CoverTab[76514]++
									p.getErrorHandler()(rw, req, fmt.Errorf("backend tried to switch to invalid protocol %q", resUpType))
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:735
		// _ = "end of CoverTab[76514]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:736
		_go_fuzz_dep_.CoverTab[76515]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:736
		// _ = "end of CoverTab[76515]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:736
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:736
	// _ = "end of CoverTab[76505]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:736
	_go_fuzz_dep_.CoverTab[76506]++
								if !ascii.EqualFold(reqUpType, resUpType) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:737
		_go_fuzz_dep_.CoverTab[76516]++
									p.getErrorHandler()(rw, req, fmt.Errorf("backend tried to switch protocol %q when %q was requested", resUpType, reqUpType))
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:739
		// _ = "end of CoverTab[76516]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:740
		_go_fuzz_dep_.CoverTab[76517]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:740
		// _ = "end of CoverTab[76517]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:740
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:740
	// _ = "end of CoverTab[76506]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:740
	_go_fuzz_dep_.CoverTab[76507]++

								hj, ok := rw.(http.Hijacker)
								if !ok {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:743
		_go_fuzz_dep_.CoverTab[76518]++
									p.getErrorHandler()(rw, req, fmt.Errorf("can't switch protocols using non-Hijacker ResponseWriter type %T", rw))
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:745
		// _ = "end of CoverTab[76518]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:746
		_go_fuzz_dep_.CoverTab[76519]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:746
		// _ = "end of CoverTab[76519]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:746
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:746
	// _ = "end of CoverTab[76507]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:746
	_go_fuzz_dep_.CoverTab[76508]++
								backConn, ok := res.Body.(io.ReadWriteCloser)
								if !ok {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:748
		_go_fuzz_dep_.CoverTab[76520]++
									p.getErrorHandler()(rw, req, fmt.Errorf("internal error: 101 switching protocols response with non-writable body"))
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:750
		// _ = "end of CoverTab[76520]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:751
		_go_fuzz_dep_.CoverTab[76521]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:751
		// _ = "end of CoverTab[76521]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:751
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:751
	// _ = "end of CoverTab[76508]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:751
	_go_fuzz_dep_.CoverTab[76509]++

								backConnCloseCh := make(chan bool)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:753
	_curRoutineNum73_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:753
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum73_)
								go func() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:754
		_go_fuzz_dep_.CoverTab[76522]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:754
		defer func() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:754
			_go_fuzz_dep_.CoverTab[76524]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:754
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum73_)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:754
			// _ = "end of CoverTab[76524]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:754
		}()

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:757
		select {
		case <-req.Context().Done():
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:758
			_go_fuzz_dep_.CoverTab[76525]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:758
			// _ = "end of CoverTab[76525]"
		case <-backConnCloseCh:
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:759
			_go_fuzz_dep_.CoverTab[76526]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:759
			// _ = "end of CoverTab[76526]"
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:760
		// _ = "end of CoverTab[76522]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:760
		_go_fuzz_dep_.CoverTab[76523]++
									backConn.Close()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:761
		// _ = "end of CoverTab[76523]"
	}()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:762
	// _ = "end of CoverTab[76509]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:762
	_go_fuzz_dep_.CoverTab[76510]++

								defer close(backConnCloseCh)

								conn, brw, err := hj.Hijack()
								if err != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:767
		_go_fuzz_dep_.CoverTab[76527]++
									p.getErrorHandler()(rw, req, fmt.Errorf("Hijack failed on protocol switch: %v", err))
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:769
		// _ = "end of CoverTab[76527]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:770
		_go_fuzz_dep_.CoverTab[76528]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:770
		// _ = "end of CoverTab[76528]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:770
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:770
	// _ = "end of CoverTab[76510]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:770
	_go_fuzz_dep_.CoverTab[76511]++
								defer conn.Close()

								copyHeader(rw.Header(), res.Header)

								res.Header = rw.Header()
								res.Body = nil
								if err := res.Write(brw); err != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:777
		_go_fuzz_dep_.CoverTab[76529]++
									p.getErrorHandler()(rw, req, fmt.Errorf("response write: %v", err))
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:779
		// _ = "end of CoverTab[76529]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:780
		_go_fuzz_dep_.CoverTab[76530]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:780
		// _ = "end of CoverTab[76530]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:780
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:780
	// _ = "end of CoverTab[76511]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:780
	_go_fuzz_dep_.CoverTab[76512]++
								if err := brw.Flush(); err != nil {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:781
		_go_fuzz_dep_.CoverTab[76531]++
									p.getErrorHandler()(rw, req, fmt.Errorf("response flush: %v", err))
									return
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:783
		// _ = "end of CoverTab[76531]"
	} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:784
		_go_fuzz_dep_.CoverTab[76532]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:784
		// _ = "end of CoverTab[76532]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:784
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:784
	// _ = "end of CoverTab[76512]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:784
	_go_fuzz_dep_.CoverTab[76513]++
								errc := make(chan error, 1)
								spc := switchProtocolCopier{user: conn, backend: backConn}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:786
	_curRoutineNum74_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:786
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum74_)
								go spc.copyToBackend(errc)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:787
	_curRoutineNum75_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:787
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum75_)
								go spc.copyFromBackend(errc)
								<-errc
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:789
	// _ = "end of CoverTab[76513]"
}

// switchProtocolCopier exists so goroutines proxying data back and
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:792
// forth have nice names in stacks.
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:794
type switchProtocolCopier struct {
	user, backend io.ReadWriter
}

func (c switchProtocolCopier) copyFromBackend(errc chan<- error) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:798
	_go_fuzz_dep_.CoverTab[76533]++
								_, err := io.Copy(c.user, c.backend)
								errc <- err
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:800
	// _ = "end of CoverTab[76533]"
}

func (c switchProtocolCopier) copyToBackend(errc chan<- error) {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:803
	_go_fuzz_dep_.CoverTab[76534]++
								_, err := io.Copy(c.backend, c.user)
								errc <- err
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:805
	// _ = "end of CoverTab[76534]"
}

func cleanQueryParams(s string) string {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:808
	_go_fuzz_dep_.CoverTab[76535]++
								reencode := func(s string) string {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:809
		_go_fuzz_dep_.CoverTab[76538]++
									v, _ := url.ParseQuery(s)
									return v.Encode()
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:811
		// _ = "end of CoverTab[76538]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:812
	// _ = "end of CoverTab[76535]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:812
	_go_fuzz_dep_.CoverTab[76536]++
								for i := 0; i < len(s); {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:813
		_go_fuzz_dep_.CoverTab[76539]++
									switch s[i] {
		case ';':
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:815
			_go_fuzz_dep_.CoverTab[76540]++
										return reencode(s)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:816
			// _ = "end of CoverTab[76540]"
		case '%':
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:817
			_go_fuzz_dep_.CoverTab[76541]++
										if i+2 >= len(s) || func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:818
				_go_fuzz_dep_.CoverTab[76544]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:818
				return !ishex(s[i+1])
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:818
				// _ = "end of CoverTab[76544]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:818
			}() || func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:818
				_go_fuzz_dep_.CoverTab[76545]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:818
				return !ishex(s[i+2])
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:818
				// _ = "end of CoverTab[76545]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:818
			}() {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:818
				_go_fuzz_dep_.CoverTab[76546]++
											return reencode(s)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:819
				// _ = "end of CoverTab[76546]"
			} else {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:820
				_go_fuzz_dep_.CoverTab[76547]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:820
				// _ = "end of CoverTab[76547]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:820
			}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:820
			// _ = "end of CoverTab[76541]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:820
			_go_fuzz_dep_.CoverTab[76542]++
										i += 3
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:821
			// _ = "end of CoverTab[76542]"
		default:
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:822
			_go_fuzz_dep_.CoverTab[76543]++
										i++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:823
			// _ = "end of CoverTab[76543]"
		}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:824
		// _ = "end of CoverTab[76539]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:825
	// _ = "end of CoverTab[76536]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:825
	_go_fuzz_dep_.CoverTab[76537]++
								return s
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:826
	// _ = "end of CoverTab[76537]"
}

func ishex(c byte) bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:829
	_go_fuzz_dep_.CoverTab[76548]++
								switch {
	case '0' <= c && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:831
		_go_fuzz_dep_.CoverTab[76554]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:831
		return c <= '9'
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:831
		// _ = "end of CoverTab[76554]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:831
	}():
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:831
		_go_fuzz_dep_.CoverTab[76550]++
									return true
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:832
		// _ = "end of CoverTab[76550]"
	case 'a' <= c && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:833
		_go_fuzz_dep_.CoverTab[76555]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:833
		return c <= 'f'
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:833
		// _ = "end of CoverTab[76555]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:833
	}():
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:833
		_go_fuzz_dep_.CoverTab[76551]++
									return true
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:834
		// _ = "end of CoverTab[76551]"
	case 'A' <= c && func() bool {
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:835
		_go_fuzz_dep_.CoverTab[76556]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:835
		return c <= 'F'
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:835
		// _ = "end of CoverTab[76556]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:835
	}():
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:835
		_go_fuzz_dep_.CoverTab[76552]++
									return true
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:836
		// _ = "end of CoverTab[76552]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:836
	default:
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:836
		_go_fuzz_dep_.CoverTab[76553]++
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:836
		// _ = "end of CoverTab[76553]"
	}
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:837
	// _ = "end of CoverTab[76548]"
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:837
	_go_fuzz_dep_.CoverTab[76549]++
								return false
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:838
	// _ = "end of CoverTab[76549]"
}

//line /usr/local/go/src/net/http/httputil/reverseproxy.go:839
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/httputil/reverseproxy.go:839
var _ = _go_fuzz_dep_.CoverTab
