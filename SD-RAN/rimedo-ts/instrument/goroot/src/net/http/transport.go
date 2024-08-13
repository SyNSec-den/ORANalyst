// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP client implementation. See RFC 7230 through 7235.
//
// This is the low-level Transport implementation of RoundTripper.
// The high-level interface is in client.go.

//line /usr/local/go/src/net/http/transport.go:10
package http

//line /usr/local/go/src/net/http/transport.go:10
import (
//line /usr/local/go/src/net/http/transport.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/transport.go:10
)
//line /usr/local/go/src/net/http/transport.go:10
import (
//line /usr/local/go/src/net/http/transport.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/transport.go:10
)

import (
	"bufio"
	"compress/gzip"
	"container/list"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"internal/godebug"
	"io"
	"log"
	"net"
	"net/http/httptrace"
	"net/http/internal/ascii"
	"net/textproto"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/http/httpguts"
	"golang.org/x/net/http/httpproxy"
)

// DefaultTransport is the default implementation of Transport and is
//line /usr/local/go/src/net/http/transport.go:38
// used by DefaultClient. It establishes network connections as needed
//line /usr/local/go/src/net/http/transport.go:38
// and caches them for reuse by subsequent calls. It uses HTTP proxies
//line /usr/local/go/src/net/http/transport.go:38
// as directed by the environment variables HTTP_PROXY, HTTPS_PROXY
//line /usr/local/go/src/net/http/transport.go:38
// and NO_PROXY (or the lowercase versions thereof).
//line /usr/local/go/src/net/http/transport.go:43
var DefaultTransport RoundTripper = &Transport{
	Proxy:	ProxyFromEnvironment,
	DialContext: defaultTransportDialContext(&net.Dialer{
		Timeout:	30 * time.Second,
		KeepAlive:	30 * time.Second,
	}),
	ForceAttemptHTTP2:	true,
	MaxIdleConns:		100,
	IdleConnTimeout:	90 * time.Second,
	TLSHandshakeTimeout:	10 * time.Second,
	ExpectContinueTimeout:	1 * time.Second,
}

// DefaultMaxIdleConnsPerHost is the default value of Transport's
//line /usr/local/go/src/net/http/transport.go:56
// MaxIdleConnsPerHost.
//line /usr/local/go/src/net/http/transport.go:58
const DefaultMaxIdleConnsPerHost = 2

// Transport is an implementation of RoundTripper that supports HTTP,
//line /usr/local/go/src/net/http/transport.go:60
// HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).
//line /usr/local/go/src/net/http/transport.go:60
//
//line /usr/local/go/src/net/http/transport.go:60
// By default, Transport caches connections for future re-use.
//line /usr/local/go/src/net/http/transport.go:60
// This may leave many open connections when accessing many hosts.
//line /usr/local/go/src/net/http/transport.go:60
// This behavior can be managed using Transport's CloseIdleConnections method
//line /usr/local/go/src/net/http/transport.go:60
// and the MaxIdleConnsPerHost and DisableKeepAlives fields.
//line /usr/local/go/src/net/http/transport.go:60
//
//line /usr/local/go/src/net/http/transport.go:60
// Transports should be reused instead of created as needed.
//line /usr/local/go/src/net/http/transport.go:60
// Transports are safe for concurrent use by multiple goroutines.
//line /usr/local/go/src/net/http/transport.go:60
//
//line /usr/local/go/src/net/http/transport.go:60
// A Transport is a low-level primitive for making HTTP and HTTPS requests.
//line /usr/local/go/src/net/http/transport.go:60
// For high-level functionality, such as cookies and redirects, see Client.
//line /usr/local/go/src/net/http/transport.go:60
//
//line /usr/local/go/src/net/http/transport.go:60
// Transport uses HTTP/1.1 for HTTP URLs and either HTTP/1.1 or HTTP/2
//line /usr/local/go/src/net/http/transport.go:60
// for HTTPS URLs, depending on whether the server supports HTTP/2,
//line /usr/local/go/src/net/http/transport.go:60
// and how the Transport is configured. The DefaultTransport supports HTTP/2.
//line /usr/local/go/src/net/http/transport.go:60
// To explicitly enable HTTP/2 on a transport, use golang.org/x/net/http2
//line /usr/local/go/src/net/http/transport.go:60
// and call ConfigureTransport. See the package docs for more about HTTP/2.
//line /usr/local/go/src/net/http/transport.go:60
//
//line /usr/local/go/src/net/http/transport.go:60
// Responses with status codes in the 1xx range are either handled
//line /usr/local/go/src/net/http/transport.go:60
// automatically (100 expect-continue) or ignored. The one
//line /usr/local/go/src/net/http/transport.go:60
// exception is HTTP status code 101 (Switching Protocols), which is
//line /usr/local/go/src/net/http/transport.go:60
// considered a terminal status and returned by RoundTrip. To see the
//line /usr/local/go/src/net/http/transport.go:60
// ignored 1xx responses, use the httptrace trace package's
//line /usr/local/go/src/net/http/transport.go:60
// ClientTrace.Got1xxResponse.
//line /usr/local/go/src/net/http/transport.go:60
//
//line /usr/local/go/src/net/http/transport.go:60
// Transport only retries a request upon encountering a network error
//line /usr/local/go/src/net/http/transport.go:60
// if the request is idempotent and either has no body or has its
//line /usr/local/go/src/net/http/transport.go:60
// Request.GetBody defined. HTTP requests are considered idempotent if
//line /usr/local/go/src/net/http/transport.go:60
// they have HTTP methods GET, HEAD, OPTIONS, or TRACE; or if their
//line /usr/local/go/src/net/http/transport.go:60
// Header map contains an "Idempotency-Key" or "X-Idempotency-Key"
//line /usr/local/go/src/net/http/transport.go:60
// entry. If the idempotency key value is a zero-length slice, the
//line /usr/local/go/src/net/http/transport.go:60
// request is treated as idempotent but the header is not sent on the
//line /usr/local/go/src/net/http/transport.go:60
// wire.
//line /usr/local/go/src/net/http/transport.go:95
type Transport struct {
	idleMu		sync.Mutex
	closeIdle	bool					// user has requested to close all idle conns
	idleConn	map[connectMethodKey][]*persistConn	// most recently used at end
	idleConnWait	map[connectMethodKey]wantConnQueue	// waiting getConns
	idleLRU		connLRU

	reqMu		sync.Mutex
	reqCanceler	map[cancelKey]func(error)

	altMu		sync.Mutex	// guards changing altProto only
	altProto	atomic.Value	// of nil or map[string]RoundTripper, key is URI scheme

	connsPerHostMu		sync.Mutex
	connsPerHost		map[connectMethodKey]int
	connsPerHostWait	map[connectMethodKey]wantConnQueue	// waiting getConns

	// Proxy specifies a function to return a proxy for a given
	// Request. If the function returns a non-nil error, the
	// request is aborted with the provided error.
	//
	// The proxy type is determined by the URL scheme. "http",
	// "https", and "socks5" are supported. If the scheme is empty,
	// "http" is assumed.
	//
	// If Proxy is nil or returns a nil *URL, no proxy is used.
	Proxy	func(*Request) (*url.URL, error)

	// OnProxyConnectResponse is called when the Transport gets an HTTP response from
	// a proxy for a CONNECT request. It's called before the check for a 200 OK response.
	// If it returns an error, the request fails with that error.
	OnProxyConnectResponse	func(ctx context.Context, proxyURL *url.URL, connectReq *Request, connectRes *Response) error

	// DialContext specifies the dial function for creating unencrypted TCP connections.
	// If DialContext is nil (and the deprecated Dial below is also nil),
	// then the transport dials using package net.
	//
	// DialContext runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using
	// a connection dialed previously when the earlier connection
	// becomes idle before the later DialContext completes.
	DialContext	func(ctx context.Context, network, addr string) (net.Conn, error)

	// Dial specifies the dial function for creating unencrypted TCP connections.
	//
	// Dial runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using
	// a connection dialed previously when the earlier connection
	// becomes idle before the later Dial completes.
	//
	// Deprecated: Use DialContext instead, which allows the transport
	// to cancel dials as soon as they are no longer needed.
	// If both are set, DialContext takes priority.
	Dial	func(network, addr string) (net.Conn, error)

	// DialTLSContext specifies an optional dial function for creating
	// TLS connections for non-proxied HTTPS requests.
	//
	// If DialTLSContext is nil (and the deprecated DialTLS below is also nil),
	// DialContext and TLSClientConfig are used.
	//
	// If DialTLSContext is set, the Dial and DialContext hooks are not used for HTTPS
	// requests and the TLSClientConfig and TLSHandshakeTimeout
	// are ignored. The returned net.Conn is assumed to already be
	// past the TLS handshake.
	DialTLSContext	func(ctx context.Context, network, addr string) (net.Conn, error)

	// DialTLS specifies an optional dial function for creating
	// TLS connections for non-proxied HTTPS requests.
	//
	// Deprecated: Use DialTLSContext instead, which allows the transport
	// to cancel dials as soon as they are no longer needed.
	// If both are set, DialTLSContext takes priority.
	DialTLS	func(network, addr string) (net.Conn, error)

	// TLSClientConfig specifies the TLS configuration to use with
	// tls.Client.
	// If nil, the default configuration is used.
	// If non-nil, HTTP/2 support may not be enabled by default.
	TLSClientConfig	*tls.Config

	// TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake. Zero means no timeout.
	TLSHandshakeTimeout	time.Duration

	// DisableKeepAlives, if true, disables HTTP keep-alives and
	// will only use the connection to the server for a single
	// HTTP request.
	//
	// This is unrelated to the similarly named TCP keep-alives.
	DisableKeepAlives	bool

	// DisableCompression, if true, prevents the Transport from
	// requesting compression with an "Accept-Encoding: gzip"
	// request header when the Request contains no existing
	// Accept-Encoding value. If the Transport requests gzip on
	// its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body. However, if the user
	// explicitly requested gzip it is not automatically
	// uncompressed.
	DisableCompression	bool

	// MaxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConns	int

	// MaxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// DefaultMaxIdleConnsPerHost is used.
	MaxIdleConnsPerHost	int

	// MaxConnsPerHost optionally limits the total number of
	// connections per host, including connections in the dialing,
	// active, and idle states. On limit violation, dials will block.
	//
	// Zero means no limit.
	MaxConnsPerHost	int

	// IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	IdleConnTimeout	time.Duration

	// ResponseHeaderTimeout, if non-zero, specifies the amount of
	// time to wait for a server's response headers after fully
	// writing the request (including its body, if any). This
	// time does not include the time to read the response body.
	ResponseHeaderTimeout	time.Duration

	// ExpectContinueTimeout, if non-zero, specifies the amount of
	// time to wait for a server's first response headers after fully
	// writing the request headers if the request has an
	// "Expect: 100-continue" header. Zero means no timeout and
	// causes the body to be sent immediately, without
	// waiting for the server to approve.
	// This time does not include the time to send the request header.
	ExpectContinueTimeout	time.Duration

	// TLSNextProto specifies how the Transport switches to an
	// alternate protocol (such as HTTP/2) after a TLS ALPN
	// protocol negotiation. If Transport dials an TLS connection
	// with a non-empty protocol name and TLSNextProto contains a
	// map entry for that key (such as "h2"), then the func is
	// called with the request's authority (such as "example.com"
	// or "example.com:1234") and the TLS connection. The function
	// must return a RoundTripper that then handles the request.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	TLSNextProto	map[string]func(authority string, c *tls.Conn) RoundTripper

	// ProxyConnectHeader optionally specifies headers to send to
	// proxies during CONNECT requests.
	// To set the header dynamically, see GetProxyConnectHeader.
	ProxyConnectHeader	Header

	// GetProxyConnectHeader optionally specifies a func to return
	// headers to send to proxyURL during a CONNECT request to the
	// ip:port target.
	// If it returns an error, the Transport's RoundTrip fails with
	// that error. It can return (nil, nil) to not add headers.
	// If GetProxyConnectHeader is non-nil, ProxyConnectHeader is
	// ignored.
	GetProxyConnectHeader	func(ctx context.Context, proxyURL *url.URL, target string) (Header, error)

	// MaxResponseHeaderBytes specifies a limit on how many
	// response bytes are allowed in the server's response
	// header.
	//
	// Zero means to use a default limit.
	MaxResponseHeaderBytes	int64

	// WriteBufferSize specifies the size of the write buffer used
	// when writing to the transport.
	// If zero, a default (currently 4KB) is used.
	WriteBufferSize	int

	// ReadBufferSize specifies the size of the read buffer used
	// when reading from the transport.
	// If zero, a default (currently 4KB) is used.
	ReadBufferSize	int

	// nextProtoOnce guards initialization of TLSNextProto and
	// h2transport (via onceSetNextProtoDefaults)
	nextProtoOnce		sync.Once
	h2transport		h2Transport	// non-nil if http2 wired up
	tlsNextProtoWasNil	bool		// whether TLSNextProto was nil when the Once fired

	// ForceAttemptHTTP2 controls whether HTTP/2 is enabled when a non-zero
	// Dial, DialTLS, or DialContext func or TLSClientConfig is provided.
	// By default, use of any those fields conservatively disables HTTP/2.
	// To use a custom dialer or TLS config and still attempt HTTP/2
	// upgrades, set this to true.
	ForceAttemptHTTP2	bool
}

// A cancelKey is the key of the reqCanceler map.
//line /usr/local/go/src/net/http/transport.go:291
// We wrap the *Request in this type since we want to use the original request,
//line /usr/local/go/src/net/http/transport.go:291
// not any transient one created by roundTrip.
//line /usr/local/go/src/net/http/transport.go:294
type cancelKey struct {
	req *Request
}

func (t *Transport) writeBufferSize() int {
//line /usr/local/go/src/net/http/transport.go:298
	_go_fuzz_dep_.CoverTab[44042]++
							if t.WriteBufferSize > 0 {
//line /usr/local/go/src/net/http/transport.go:299
		_go_fuzz_dep_.CoverTab[44044]++
								return t.WriteBufferSize
//line /usr/local/go/src/net/http/transport.go:300
		// _ = "end of CoverTab[44044]"
	} else {
//line /usr/local/go/src/net/http/transport.go:301
		_go_fuzz_dep_.CoverTab[44045]++
//line /usr/local/go/src/net/http/transport.go:301
		// _ = "end of CoverTab[44045]"
//line /usr/local/go/src/net/http/transport.go:301
	}
//line /usr/local/go/src/net/http/transport.go:301
	// _ = "end of CoverTab[44042]"
//line /usr/local/go/src/net/http/transport.go:301
	_go_fuzz_dep_.CoverTab[44043]++
							return 4 << 10
//line /usr/local/go/src/net/http/transport.go:302
	// _ = "end of CoverTab[44043]"
}

func (t *Transport) readBufferSize() int {
//line /usr/local/go/src/net/http/transport.go:305
	_go_fuzz_dep_.CoverTab[44046]++
							if t.ReadBufferSize > 0 {
//line /usr/local/go/src/net/http/transport.go:306
		_go_fuzz_dep_.CoverTab[44048]++
								return t.ReadBufferSize
//line /usr/local/go/src/net/http/transport.go:307
		// _ = "end of CoverTab[44048]"
	} else {
//line /usr/local/go/src/net/http/transport.go:308
		_go_fuzz_dep_.CoverTab[44049]++
//line /usr/local/go/src/net/http/transport.go:308
		// _ = "end of CoverTab[44049]"
//line /usr/local/go/src/net/http/transport.go:308
	}
//line /usr/local/go/src/net/http/transport.go:308
	// _ = "end of CoverTab[44046]"
//line /usr/local/go/src/net/http/transport.go:308
	_go_fuzz_dep_.CoverTab[44047]++
							return 4 << 10
//line /usr/local/go/src/net/http/transport.go:309
	// _ = "end of CoverTab[44047]"
}

// Clone returns a deep copy of t's exported fields.
func (t *Transport) Clone() *Transport {
//line /usr/local/go/src/net/http/transport.go:313
	_go_fuzz_dep_.CoverTab[44050]++
							t.nextProtoOnce.Do(t.onceSetNextProtoDefaults)
							t2 := &Transport{
		Proxy:			t.Proxy,
		OnProxyConnectResponse:	t.OnProxyConnectResponse,
		DialContext:		t.DialContext,
		Dial:			t.Dial,
		DialTLS:		t.DialTLS,
		DialTLSContext:		t.DialTLSContext,
		TLSHandshakeTimeout:	t.TLSHandshakeTimeout,
		DisableKeepAlives:	t.DisableKeepAlives,
		DisableCompression:	t.DisableCompression,
		MaxIdleConns:		t.MaxIdleConns,
		MaxIdleConnsPerHost:	t.MaxIdleConnsPerHost,
		MaxConnsPerHost:	t.MaxConnsPerHost,
		IdleConnTimeout:	t.IdleConnTimeout,
		ResponseHeaderTimeout:	t.ResponseHeaderTimeout,
		ExpectContinueTimeout:	t.ExpectContinueTimeout,
		ProxyConnectHeader:	t.ProxyConnectHeader.Clone(),
		GetProxyConnectHeader:	t.GetProxyConnectHeader,
		MaxResponseHeaderBytes:	t.MaxResponseHeaderBytes,
		ForceAttemptHTTP2:	t.ForceAttemptHTTP2,
		WriteBufferSize:	t.WriteBufferSize,
		ReadBufferSize:		t.ReadBufferSize,
	}
	if t.TLSClientConfig != nil {
//line /usr/local/go/src/net/http/transport.go:338
		_go_fuzz_dep_.CoverTab[44053]++
								t2.TLSClientConfig = t.TLSClientConfig.Clone()
//line /usr/local/go/src/net/http/transport.go:339
		// _ = "end of CoverTab[44053]"
	} else {
//line /usr/local/go/src/net/http/transport.go:340
		_go_fuzz_dep_.CoverTab[44054]++
//line /usr/local/go/src/net/http/transport.go:340
		// _ = "end of CoverTab[44054]"
//line /usr/local/go/src/net/http/transport.go:340
	}
//line /usr/local/go/src/net/http/transport.go:340
	// _ = "end of CoverTab[44050]"
//line /usr/local/go/src/net/http/transport.go:340
	_go_fuzz_dep_.CoverTab[44051]++
							if !t.tlsNextProtoWasNil {
//line /usr/local/go/src/net/http/transport.go:341
		_go_fuzz_dep_.CoverTab[44055]++
								npm := map[string]func(authority string, c *tls.Conn) RoundTripper{}
								for k, v := range t.TLSNextProto {
//line /usr/local/go/src/net/http/transport.go:343
			_go_fuzz_dep_.CoverTab[44057]++
									npm[k] = v
//line /usr/local/go/src/net/http/transport.go:344
			// _ = "end of CoverTab[44057]"
		}
//line /usr/local/go/src/net/http/transport.go:345
		// _ = "end of CoverTab[44055]"
//line /usr/local/go/src/net/http/transport.go:345
		_go_fuzz_dep_.CoverTab[44056]++
								t2.TLSNextProto = npm
//line /usr/local/go/src/net/http/transport.go:346
		// _ = "end of CoverTab[44056]"
	} else {
//line /usr/local/go/src/net/http/transport.go:347
		_go_fuzz_dep_.CoverTab[44058]++
//line /usr/local/go/src/net/http/transport.go:347
		// _ = "end of CoverTab[44058]"
//line /usr/local/go/src/net/http/transport.go:347
	}
//line /usr/local/go/src/net/http/transport.go:347
	// _ = "end of CoverTab[44051]"
//line /usr/local/go/src/net/http/transport.go:347
	_go_fuzz_dep_.CoverTab[44052]++
							return t2
//line /usr/local/go/src/net/http/transport.go:348
	// _ = "end of CoverTab[44052]"
}

// h2Transport is the interface we expect to be able to call from
//line /usr/local/go/src/net/http/transport.go:351
// net/http against an *http2.Transport that's either bundled into
//line /usr/local/go/src/net/http/transport.go:351
// h2_bundle.go or supplied by the user via x/net/http2.
//line /usr/local/go/src/net/http/transport.go:351
//
//line /usr/local/go/src/net/http/transport.go:351
// We name it with the "h2" prefix to stay out of the "http2" prefix
//line /usr/local/go/src/net/http/transport.go:351
// namespace used by x/tools/cmd/bundle for h2_bundle.go.
//line /usr/local/go/src/net/http/transport.go:357
type h2Transport interface {
	CloseIdleConnections()
}

func (t *Transport) hasCustomTLSDialer() bool {
//line /usr/local/go/src/net/http/transport.go:361
	_go_fuzz_dep_.CoverTab[44059]++
							return t.DialTLS != nil || func() bool {
//line /usr/local/go/src/net/http/transport.go:362
		_go_fuzz_dep_.CoverTab[44060]++
//line /usr/local/go/src/net/http/transport.go:362
		return t.DialTLSContext != nil
//line /usr/local/go/src/net/http/transport.go:362
		// _ = "end of CoverTab[44060]"
//line /usr/local/go/src/net/http/transport.go:362
	}()
//line /usr/local/go/src/net/http/transport.go:362
	// _ = "end of CoverTab[44059]"
}

var http2client = godebug.New("http2client")

// onceSetNextProtoDefaults initializes TLSNextProto.
//line /usr/local/go/src/net/http/transport.go:367
// It must be called via t.nextProtoOnce.Do.
//line /usr/local/go/src/net/http/transport.go:369
func (t *Transport) onceSetNextProtoDefaults() {
//line /usr/local/go/src/net/http/transport.go:369
	_go_fuzz_dep_.CoverTab[44061]++
							t.tlsNextProtoWasNil = (t.TLSNextProto == nil)
							if http2client.Value() == "0" {
//line /usr/local/go/src/net/http/transport.go:371
		_go_fuzz_dep_.CoverTab[44068]++
								return
//line /usr/local/go/src/net/http/transport.go:372
		// _ = "end of CoverTab[44068]"
	} else {
//line /usr/local/go/src/net/http/transport.go:373
		_go_fuzz_dep_.CoverTab[44069]++
//line /usr/local/go/src/net/http/transport.go:373
		// _ = "end of CoverTab[44069]"
//line /usr/local/go/src/net/http/transport.go:373
	}
//line /usr/local/go/src/net/http/transport.go:373
	// _ = "end of CoverTab[44061]"
//line /usr/local/go/src/net/http/transport.go:373
	_go_fuzz_dep_.CoverTab[44062]++

//line /usr/local/go/src/net/http/transport.go:380
	altProto, _ := t.altProto.Load().(map[string]RoundTripper)
	if rv := reflect.ValueOf(altProto["https"]); rv.IsValid() && func() bool {
//line /usr/local/go/src/net/http/transport.go:381
		_go_fuzz_dep_.CoverTab[44070]++
//line /usr/local/go/src/net/http/transport.go:381
		return rv.Type().Kind() == reflect.Struct
//line /usr/local/go/src/net/http/transport.go:381
		// _ = "end of CoverTab[44070]"
//line /usr/local/go/src/net/http/transport.go:381
	}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:381
		_go_fuzz_dep_.CoverTab[44071]++
//line /usr/local/go/src/net/http/transport.go:381
		return rv.Type().NumField() == 1
//line /usr/local/go/src/net/http/transport.go:381
		// _ = "end of CoverTab[44071]"
//line /usr/local/go/src/net/http/transport.go:381
	}() {
//line /usr/local/go/src/net/http/transport.go:381
		_go_fuzz_dep_.CoverTab[44072]++
								if v := rv.Field(0); v.CanInterface() {
//line /usr/local/go/src/net/http/transport.go:382
			_go_fuzz_dep_.CoverTab[44073]++
									if h2i, ok := v.Interface().(h2Transport); ok {
//line /usr/local/go/src/net/http/transport.go:383
				_go_fuzz_dep_.CoverTab[44074]++
										t.h2transport = h2i
										return
//line /usr/local/go/src/net/http/transport.go:385
				// _ = "end of CoverTab[44074]"
			} else {
//line /usr/local/go/src/net/http/transport.go:386
				_go_fuzz_dep_.CoverTab[44075]++
//line /usr/local/go/src/net/http/transport.go:386
				// _ = "end of CoverTab[44075]"
//line /usr/local/go/src/net/http/transport.go:386
			}
//line /usr/local/go/src/net/http/transport.go:386
			// _ = "end of CoverTab[44073]"
		} else {
//line /usr/local/go/src/net/http/transport.go:387
			_go_fuzz_dep_.CoverTab[44076]++
//line /usr/local/go/src/net/http/transport.go:387
			// _ = "end of CoverTab[44076]"
//line /usr/local/go/src/net/http/transport.go:387
		}
//line /usr/local/go/src/net/http/transport.go:387
		// _ = "end of CoverTab[44072]"
	} else {
//line /usr/local/go/src/net/http/transport.go:388
		_go_fuzz_dep_.CoverTab[44077]++
//line /usr/local/go/src/net/http/transport.go:388
		// _ = "end of CoverTab[44077]"
//line /usr/local/go/src/net/http/transport.go:388
	}
//line /usr/local/go/src/net/http/transport.go:388
	// _ = "end of CoverTab[44062]"
//line /usr/local/go/src/net/http/transport.go:388
	_go_fuzz_dep_.CoverTab[44063]++

							if t.TLSNextProto != nil {
//line /usr/local/go/src/net/http/transport.go:390
		_go_fuzz_dep_.CoverTab[44078]++

//line /usr/local/go/src/net/http/transport.go:393
		return
//line /usr/local/go/src/net/http/transport.go:393
		// _ = "end of CoverTab[44078]"
	} else {
//line /usr/local/go/src/net/http/transport.go:394
		_go_fuzz_dep_.CoverTab[44079]++
//line /usr/local/go/src/net/http/transport.go:394
		// _ = "end of CoverTab[44079]"
//line /usr/local/go/src/net/http/transport.go:394
	}
//line /usr/local/go/src/net/http/transport.go:394
	// _ = "end of CoverTab[44063]"
//line /usr/local/go/src/net/http/transport.go:394
	_go_fuzz_dep_.CoverTab[44064]++
							if !t.ForceAttemptHTTP2 && func() bool {
//line /usr/local/go/src/net/http/transport.go:395
		_go_fuzz_dep_.CoverTab[44080]++
//line /usr/local/go/src/net/http/transport.go:395
		return (t.TLSClientConfig != nil || func() bool {
//line /usr/local/go/src/net/http/transport.go:395
			_go_fuzz_dep_.CoverTab[44081]++
//line /usr/local/go/src/net/http/transport.go:395
			return t.Dial != nil
//line /usr/local/go/src/net/http/transport.go:395
			// _ = "end of CoverTab[44081]"
//line /usr/local/go/src/net/http/transport.go:395
		}() || func() bool {
//line /usr/local/go/src/net/http/transport.go:395
			_go_fuzz_dep_.CoverTab[44082]++
//line /usr/local/go/src/net/http/transport.go:395
			return t.DialContext != nil
//line /usr/local/go/src/net/http/transport.go:395
			// _ = "end of CoverTab[44082]"
//line /usr/local/go/src/net/http/transport.go:395
		}() || func() bool {
//line /usr/local/go/src/net/http/transport.go:395
			_go_fuzz_dep_.CoverTab[44083]++
//line /usr/local/go/src/net/http/transport.go:395
			return t.hasCustomTLSDialer()
//line /usr/local/go/src/net/http/transport.go:395
			// _ = "end of CoverTab[44083]"
//line /usr/local/go/src/net/http/transport.go:395
		}())
//line /usr/local/go/src/net/http/transport.go:395
		// _ = "end of CoverTab[44080]"
//line /usr/local/go/src/net/http/transport.go:395
	}() {
//line /usr/local/go/src/net/http/transport.go:395
		_go_fuzz_dep_.CoverTab[44084]++

//line /usr/local/go/src/net/http/transport.go:402
		return
//line /usr/local/go/src/net/http/transport.go:402
		// _ = "end of CoverTab[44084]"
	} else {
//line /usr/local/go/src/net/http/transport.go:403
		_go_fuzz_dep_.CoverTab[44085]++
//line /usr/local/go/src/net/http/transport.go:403
		// _ = "end of CoverTab[44085]"
//line /usr/local/go/src/net/http/transport.go:403
	}
//line /usr/local/go/src/net/http/transport.go:403
	// _ = "end of CoverTab[44064]"
//line /usr/local/go/src/net/http/transport.go:403
	_go_fuzz_dep_.CoverTab[44065]++
							if omitBundledHTTP2 {
//line /usr/local/go/src/net/http/transport.go:404
		_go_fuzz_dep_.CoverTab[44086]++
								return
//line /usr/local/go/src/net/http/transport.go:405
		// _ = "end of CoverTab[44086]"
	} else {
//line /usr/local/go/src/net/http/transport.go:406
		_go_fuzz_dep_.CoverTab[44087]++
//line /usr/local/go/src/net/http/transport.go:406
		// _ = "end of CoverTab[44087]"
//line /usr/local/go/src/net/http/transport.go:406
	}
//line /usr/local/go/src/net/http/transport.go:406
	// _ = "end of CoverTab[44065]"
//line /usr/local/go/src/net/http/transport.go:406
	_go_fuzz_dep_.CoverTab[44066]++
							t2, err := http2configureTransports(t)
							if err != nil {
//line /usr/local/go/src/net/http/transport.go:408
		_go_fuzz_dep_.CoverTab[44088]++
								log.Printf("Error enabling Transport HTTP/2 support: %v", err)
								return
//line /usr/local/go/src/net/http/transport.go:410
		// _ = "end of CoverTab[44088]"
	} else {
//line /usr/local/go/src/net/http/transport.go:411
		_go_fuzz_dep_.CoverTab[44089]++
//line /usr/local/go/src/net/http/transport.go:411
		// _ = "end of CoverTab[44089]"
//line /usr/local/go/src/net/http/transport.go:411
	}
//line /usr/local/go/src/net/http/transport.go:411
	// _ = "end of CoverTab[44066]"
//line /usr/local/go/src/net/http/transport.go:411
	_go_fuzz_dep_.CoverTab[44067]++
							t.h2transport = t2

//line /usr/local/go/src/net/http/transport.go:420
	if limit1 := t.MaxResponseHeaderBytes; limit1 != 0 && func() bool {
//line /usr/local/go/src/net/http/transport.go:420
		_go_fuzz_dep_.CoverTab[44090]++
//line /usr/local/go/src/net/http/transport.go:420
		return t2.MaxHeaderListSize == 0
//line /usr/local/go/src/net/http/transport.go:420
		// _ = "end of CoverTab[44090]"
//line /usr/local/go/src/net/http/transport.go:420
	}() {
//line /usr/local/go/src/net/http/transport.go:420
		_go_fuzz_dep_.CoverTab[44091]++
								const h2max = 1<<32 - 1
								if limit1 >= h2max {
//line /usr/local/go/src/net/http/transport.go:422
			_go_fuzz_dep_.CoverTab[44092]++
									t2.MaxHeaderListSize = h2max
//line /usr/local/go/src/net/http/transport.go:423
			// _ = "end of CoverTab[44092]"
		} else {
//line /usr/local/go/src/net/http/transport.go:424
			_go_fuzz_dep_.CoverTab[44093]++
									t2.MaxHeaderListSize = uint32(limit1)
//line /usr/local/go/src/net/http/transport.go:425
			// _ = "end of CoverTab[44093]"
		}
//line /usr/local/go/src/net/http/transport.go:426
		// _ = "end of CoverTab[44091]"
	} else {
//line /usr/local/go/src/net/http/transport.go:427
		_go_fuzz_dep_.CoverTab[44094]++
//line /usr/local/go/src/net/http/transport.go:427
		// _ = "end of CoverTab[44094]"
//line /usr/local/go/src/net/http/transport.go:427
	}
//line /usr/local/go/src/net/http/transport.go:427
	// _ = "end of CoverTab[44067]"
}

// ProxyFromEnvironment returns the URL of the proxy to use for a
//line /usr/local/go/src/net/http/transport.go:430
// given request, as indicated by the environment variables
//line /usr/local/go/src/net/http/transport.go:430
// HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions
//line /usr/local/go/src/net/http/transport.go:430
// thereof). Requests use the proxy from the environment variable
//line /usr/local/go/src/net/http/transport.go:430
// matching their scheme, unless excluded by NO_PROXY.
//line /usr/local/go/src/net/http/transport.go:430
//
//line /usr/local/go/src/net/http/transport.go:430
// The environment values may be either a complete URL or a
//line /usr/local/go/src/net/http/transport.go:430
// "host[:port]", in which case the "http" scheme is assumed.
//line /usr/local/go/src/net/http/transport.go:430
// The schemes "http", "https", and "socks5" are supported.
//line /usr/local/go/src/net/http/transport.go:430
// An error is returned if the value is a different form.
//line /usr/local/go/src/net/http/transport.go:430
//
//line /usr/local/go/src/net/http/transport.go:430
// A nil URL and nil error are returned if no proxy is defined in the
//line /usr/local/go/src/net/http/transport.go:430
// environment, or a proxy should not be used for the given request,
//line /usr/local/go/src/net/http/transport.go:430
// as defined by NO_PROXY.
//line /usr/local/go/src/net/http/transport.go:430
//
//line /usr/local/go/src/net/http/transport.go:430
// As a special case, if req.URL.Host is "localhost" (with or without
//line /usr/local/go/src/net/http/transport.go:430
// a port number), then a nil URL and nil error will be returned.
//line /usr/local/go/src/net/http/transport.go:447
func ProxyFromEnvironment(req *Request) (*url.URL, error) {
//line /usr/local/go/src/net/http/transport.go:447
	_go_fuzz_dep_.CoverTab[44095]++
							return envProxyFunc()(req.URL)
//line /usr/local/go/src/net/http/transport.go:448
	// _ = "end of CoverTab[44095]"
}

// ProxyURL returns a proxy function (for use in a Transport)
//line /usr/local/go/src/net/http/transport.go:451
// that always returns the same URL.
//line /usr/local/go/src/net/http/transport.go:453
func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error) {
//line /usr/local/go/src/net/http/transport.go:453
	_go_fuzz_dep_.CoverTab[44096]++
							return func(*Request) (*url.URL, error) {
//line /usr/local/go/src/net/http/transport.go:454
		_go_fuzz_dep_.CoverTab[44097]++
								return fixedURL, nil
//line /usr/local/go/src/net/http/transport.go:455
		// _ = "end of CoverTab[44097]"
	}
//line /usr/local/go/src/net/http/transport.go:456
	// _ = "end of CoverTab[44096]"
}

// transportRequest is a wrapper around a *Request that adds
//line /usr/local/go/src/net/http/transport.go:459
// optional extra headers to write and stores any error to return
//line /usr/local/go/src/net/http/transport.go:459
// from roundTrip.
//line /usr/local/go/src/net/http/transport.go:462
type transportRequest struct {
	*Request				// original request, not to be mutated
	extra		Header			// extra headers to write, or nil
	trace		*httptrace.ClientTrace	// optional
	cancelKey	cancelKey

	mu	sync.Mutex	// guards err
	err	error		// first setError value for mapRoundTripError to consider
}

func (tr *transportRequest) extraHeaders() Header {
//line /usr/local/go/src/net/http/transport.go:472
	_go_fuzz_dep_.CoverTab[44098]++
							if tr.extra == nil {
//line /usr/local/go/src/net/http/transport.go:473
		_go_fuzz_dep_.CoverTab[44100]++
								tr.extra = make(Header)
//line /usr/local/go/src/net/http/transport.go:474
		// _ = "end of CoverTab[44100]"
	} else {
//line /usr/local/go/src/net/http/transport.go:475
		_go_fuzz_dep_.CoverTab[44101]++
//line /usr/local/go/src/net/http/transport.go:475
		// _ = "end of CoverTab[44101]"
//line /usr/local/go/src/net/http/transport.go:475
	}
//line /usr/local/go/src/net/http/transport.go:475
	// _ = "end of CoverTab[44098]"
//line /usr/local/go/src/net/http/transport.go:475
	_go_fuzz_dep_.CoverTab[44099]++
							return tr.extra
//line /usr/local/go/src/net/http/transport.go:476
	// _ = "end of CoverTab[44099]"
}

func (tr *transportRequest) setError(err error) {
//line /usr/local/go/src/net/http/transport.go:479
	_go_fuzz_dep_.CoverTab[44102]++
							tr.mu.Lock()
							if tr.err == nil {
//line /usr/local/go/src/net/http/transport.go:481
		_go_fuzz_dep_.CoverTab[44104]++
								tr.err = err
//line /usr/local/go/src/net/http/transport.go:482
		// _ = "end of CoverTab[44104]"
	} else {
//line /usr/local/go/src/net/http/transport.go:483
		_go_fuzz_dep_.CoverTab[44105]++
//line /usr/local/go/src/net/http/transport.go:483
		// _ = "end of CoverTab[44105]"
//line /usr/local/go/src/net/http/transport.go:483
	}
//line /usr/local/go/src/net/http/transport.go:483
	// _ = "end of CoverTab[44102]"
//line /usr/local/go/src/net/http/transport.go:483
	_go_fuzz_dep_.CoverTab[44103]++
							tr.mu.Unlock()
//line /usr/local/go/src/net/http/transport.go:484
	// _ = "end of CoverTab[44103]"
}

// useRegisteredProtocol reports whether an alternate protocol (as registered
//line /usr/local/go/src/net/http/transport.go:487
// with Transport.RegisterProtocol) should be respected for this request.
//line /usr/local/go/src/net/http/transport.go:489
func (t *Transport) useRegisteredProtocol(req *Request) bool {
//line /usr/local/go/src/net/http/transport.go:489
	_go_fuzz_dep_.CoverTab[44106]++
							if req.URL.Scheme == "https" && func() bool {
//line /usr/local/go/src/net/http/transport.go:490
		_go_fuzz_dep_.CoverTab[44108]++
//line /usr/local/go/src/net/http/transport.go:490
		return req.requiresHTTP1()
//line /usr/local/go/src/net/http/transport.go:490
		// _ = "end of CoverTab[44108]"
//line /usr/local/go/src/net/http/transport.go:490
	}() {
//line /usr/local/go/src/net/http/transport.go:490
		_go_fuzz_dep_.CoverTab[44109]++

//line /usr/local/go/src/net/http/transport.go:495
		return false
//line /usr/local/go/src/net/http/transport.go:495
		// _ = "end of CoverTab[44109]"
	} else {
//line /usr/local/go/src/net/http/transport.go:496
		_go_fuzz_dep_.CoverTab[44110]++
//line /usr/local/go/src/net/http/transport.go:496
		// _ = "end of CoverTab[44110]"
//line /usr/local/go/src/net/http/transport.go:496
	}
//line /usr/local/go/src/net/http/transport.go:496
	// _ = "end of CoverTab[44106]"
//line /usr/local/go/src/net/http/transport.go:496
	_go_fuzz_dep_.CoverTab[44107]++
							return true
//line /usr/local/go/src/net/http/transport.go:497
	// _ = "end of CoverTab[44107]"
}

// alternateRoundTripper returns the alternate RoundTripper to use
//line /usr/local/go/src/net/http/transport.go:500
// for this request if the Request's URL scheme requires one,
//line /usr/local/go/src/net/http/transport.go:500
// or nil for the normal case of using the Transport.
//line /usr/local/go/src/net/http/transport.go:503
func (t *Transport) alternateRoundTripper(req *Request) RoundTripper {
//line /usr/local/go/src/net/http/transport.go:503
	_go_fuzz_dep_.CoverTab[44111]++
							if !t.useRegisteredProtocol(req) {
//line /usr/local/go/src/net/http/transport.go:504
		_go_fuzz_dep_.CoverTab[44113]++
								return nil
//line /usr/local/go/src/net/http/transport.go:505
		// _ = "end of CoverTab[44113]"
	} else {
//line /usr/local/go/src/net/http/transport.go:506
		_go_fuzz_dep_.CoverTab[44114]++
//line /usr/local/go/src/net/http/transport.go:506
		// _ = "end of CoverTab[44114]"
//line /usr/local/go/src/net/http/transport.go:506
	}
//line /usr/local/go/src/net/http/transport.go:506
	// _ = "end of CoverTab[44111]"
//line /usr/local/go/src/net/http/transport.go:506
	_go_fuzz_dep_.CoverTab[44112]++
							altProto, _ := t.altProto.Load().(map[string]RoundTripper)
							return altProto[req.URL.Scheme]
//line /usr/local/go/src/net/http/transport.go:508
	// _ = "end of CoverTab[44112]"
}

// roundTrip implements a RoundTripper over HTTP.
func (t *Transport) roundTrip(req *Request) (*Response, error) {
//line /usr/local/go/src/net/http/transport.go:512
	_go_fuzz_dep_.CoverTab[44115]++
							t.nextProtoOnce.Do(t.onceSetNextProtoDefaults)
							ctx := req.Context()
							trace := httptrace.ContextClientTrace(ctx)

							if req.URL == nil {
//line /usr/local/go/src/net/http/transport.go:517
		_go_fuzz_dep_.CoverTab[44123]++
								req.closeBody()
								return nil, errors.New("http: nil Request.URL")
//line /usr/local/go/src/net/http/transport.go:519
		// _ = "end of CoverTab[44123]"
	} else {
//line /usr/local/go/src/net/http/transport.go:520
		_go_fuzz_dep_.CoverTab[44124]++
//line /usr/local/go/src/net/http/transport.go:520
		// _ = "end of CoverTab[44124]"
//line /usr/local/go/src/net/http/transport.go:520
	}
//line /usr/local/go/src/net/http/transport.go:520
	// _ = "end of CoverTab[44115]"
//line /usr/local/go/src/net/http/transport.go:520
	_go_fuzz_dep_.CoverTab[44116]++
							if req.Header == nil {
//line /usr/local/go/src/net/http/transport.go:521
		_go_fuzz_dep_.CoverTab[44125]++
								req.closeBody()
								return nil, errors.New("http: nil Request.Header")
//line /usr/local/go/src/net/http/transport.go:523
		// _ = "end of CoverTab[44125]"
	} else {
//line /usr/local/go/src/net/http/transport.go:524
		_go_fuzz_dep_.CoverTab[44126]++
//line /usr/local/go/src/net/http/transport.go:524
		// _ = "end of CoverTab[44126]"
//line /usr/local/go/src/net/http/transport.go:524
	}
//line /usr/local/go/src/net/http/transport.go:524
	// _ = "end of CoverTab[44116]"
//line /usr/local/go/src/net/http/transport.go:524
	_go_fuzz_dep_.CoverTab[44117]++
							scheme := req.URL.Scheme
							isHTTP := scheme == "http" || func() bool {
//line /usr/local/go/src/net/http/transport.go:526
		_go_fuzz_dep_.CoverTab[44127]++
//line /usr/local/go/src/net/http/transport.go:526
		return scheme == "https"
//line /usr/local/go/src/net/http/transport.go:526
		// _ = "end of CoverTab[44127]"
//line /usr/local/go/src/net/http/transport.go:526
	}()
							if isHTTP {
//line /usr/local/go/src/net/http/transport.go:527
		_go_fuzz_dep_.CoverTab[44128]++
								for k, vv := range req.Header {
//line /usr/local/go/src/net/http/transport.go:528
			_go_fuzz_dep_.CoverTab[44129]++
									if !httpguts.ValidHeaderFieldName(k) {
//line /usr/local/go/src/net/http/transport.go:529
				_go_fuzz_dep_.CoverTab[44131]++
										req.closeBody()
										return nil, fmt.Errorf("net/http: invalid header field name %q", k)
//line /usr/local/go/src/net/http/transport.go:531
				// _ = "end of CoverTab[44131]"
			} else {
//line /usr/local/go/src/net/http/transport.go:532
				_go_fuzz_dep_.CoverTab[44132]++
//line /usr/local/go/src/net/http/transport.go:532
				// _ = "end of CoverTab[44132]"
//line /usr/local/go/src/net/http/transport.go:532
			}
//line /usr/local/go/src/net/http/transport.go:532
			// _ = "end of CoverTab[44129]"
//line /usr/local/go/src/net/http/transport.go:532
			_go_fuzz_dep_.CoverTab[44130]++
									for _, v := range vv {
//line /usr/local/go/src/net/http/transport.go:533
				_go_fuzz_dep_.CoverTab[44133]++
										if !httpguts.ValidHeaderFieldValue(v) {
//line /usr/local/go/src/net/http/transport.go:534
					_go_fuzz_dep_.CoverTab[44134]++
											req.closeBody()

											return nil, fmt.Errorf("net/http: invalid header field value for %q", k)
//line /usr/local/go/src/net/http/transport.go:537
					// _ = "end of CoverTab[44134]"
				} else {
//line /usr/local/go/src/net/http/transport.go:538
					_go_fuzz_dep_.CoverTab[44135]++
//line /usr/local/go/src/net/http/transport.go:538
					// _ = "end of CoverTab[44135]"
//line /usr/local/go/src/net/http/transport.go:538
				}
//line /usr/local/go/src/net/http/transport.go:538
				// _ = "end of CoverTab[44133]"
			}
//line /usr/local/go/src/net/http/transport.go:539
			// _ = "end of CoverTab[44130]"
		}
//line /usr/local/go/src/net/http/transport.go:540
		// _ = "end of CoverTab[44128]"
	} else {
//line /usr/local/go/src/net/http/transport.go:541
		_go_fuzz_dep_.CoverTab[44136]++
//line /usr/local/go/src/net/http/transport.go:541
		// _ = "end of CoverTab[44136]"
//line /usr/local/go/src/net/http/transport.go:541
	}
//line /usr/local/go/src/net/http/transport.go:541
	// _ = "end of CoverTab[44117]"
//line /usr/local/go/src/net/http/transport.go:541
	_go_fuzz_dep_.CoverTab[44118]++

							origReq := req
							cancelKey := cancelKey{origReq}
							req = setupRewindBody(req)

							if altRT := t.alternateRoundTripper(req); altRT != nil {
//line /usr/local/go/src/net/http/transport.go:547
		_go_fuzz_dep_.CoverTab[44137]++
								if resp, err := altRT.RoundTrip(req); err != ErrSkipAltProtocol {
//line /usr/local/go/src/net/http/transport.go:548
			_go_fuzz_dep_.CoverTab[44139]++
									return resp, err
//line /usr/local/go/src/net/http/transport.go:549
			// _ = "end of CoverTab[44139]"
		} else {
//line /usr/local/go/src/net/http/transport.go:550
			_go_fuzz_dep_.CoverTab[44140]++
//line /usr/local/go/src/net/http/transport.go:550
			// _ = "end of CoverTab[44140]"
//line /usr/local/go/src/net/http/transport.go:550
		}
//line /usr/local/go/src/net/http/transport.go:550
		// _ = "end of CoverTab[44137]"
//line /usr/local/go/src/net/http/transport.go:550
		_go_fuzz_dep_.CoverTab[44138]++
								var err error
								req, err = rewindBody(req)
								if err != nil {
//line /usr/local/go/src/net/http/transport.go:553
			_go_fuzz_dep_.CoverTab[44141]++
									return nil, err
//line /usr/local/go/src/net/http/transport.go:554
			// _ = "end of CoverTab[44141]"
		} else {
//line /usr/local/go/src/net/http/transport.go:555
			_go_fuzz_dep_.CoverTab[44142]++
//line /usr/local/go/src/net/http/transport.go:555
			// _ = "end of CoverTab[44142]"
//line /usr/local/go/src/net/http/transport.go:555
		}
//line /usr/local/go/src/net/http/transport.go:555
		// _ = "end of CoverTab[44138]"
	} else {
//line /usr/local/go/src/net/http/transport.go:556
		_go_fuzz_dep_.CoverTab[44143]++
//line /usr/local/go/src/net/http/transport.go:556
		// _ = "end of CoverTab[44143]"
//line /usr/local/go/src/net/http/transport.go:556
	}
//line /usr/local/go/src/net/http/transport.go:556
	// _ = "end of CoverTab[44118]"
//line /usr/local/go/src/net/http/transport.go:556
	_go_fuzz_dep_.CoverTab[44119]++
							if !isHTTP {
//line /usr/local/go/src/net/http/transport.go:557
		_go_fuzz_dep_.CoverTab[44144]++
								req.closeBody()
								return nil, badStringError("unsupported protocol scheme", scheme)
//line /usr/local/go/src/net/http/transport.go:559
		// _ = "end of CoverTab[44144]"
	} else {
//line /usr/local/go/src/net/http/transport.go:560
		_go_fuzz_dep_.CoverTab[44145]++
//line /usr/local/go/src/net/http/transport.go:560
		// _ = "end of CoverTab[44145]"
//line /usr/local/go/src/net/http/transport.go:560
	}
//line /usr/local/go/src/net/http/transport.go:560
	// _ = "end of CoverTab[44119]"
//line /usr/local/go/src/net/http/transport.go:560
	_go_fuzz_dep_.CoverTab[44120]++
							if req.Method != "" && func() bool {
//line /usr/local/go/src/net/http/transport.go:561
		_go_fuzz_dep_.CoverTab[44146]++
//line /usr/local/go/src/net/http/transport.go:561
		return !validMethod(req.Method)
//line /usr/local/go/src/net/http/transport.go:561
		// _ = "end of CoverTab[44146]"
//line /usr/local/go/src/net/http/transport.go:561
	}() {
//line /usr/local/go/src/net/http/transport.go:561
		_go_fuzz_dep_.CoverTab[44147]++
								req.closeBody()
								return nil, fmt.Errorf("net/http: invalid method %q", req.Method)
//line /usr/local/go/src/net/http/transport.go:563
		// _ = "end of CoverTab[44147]"
	} else {
//line /usr/local/go/src/net/http/transport.go:564
		_go_fuzz_dep_.CoverTab[44148]++
//line /usr/local/go/src/net/http/transport.go:564
		// _ = "end of CoverTab[44148]"
//line /usr/local/go/src/net/http/transport.go:564
	}
//line /usr/local/go/src/net/http/transport.go:564
	// _ = "end of CoverTab[44120]"
//line /usr/local/go/src/net/http/transport.go:564
	_go_fuzz_dep_.CoverTab[44121]++
							if req.URL.Host == "" {
//line /usr/local/go/src/net/http/transport.go:565
		_go_fuzz_dep_.CoverTab[44149]++
								req.closeBody()
								return nil, errors.New("http: no Host in request URL")
//line /usr/local/go/src/net/http/transport.go:567
		// _ = "end of CoverTab[44149]"
	} else {
//line /usr/local/go/src/net/http/transport.go:568
		_go_fuzz_dep_.CoverTab[44150]++
//line /usr/local/go/src/net/http/transport.go:568
		// _ = "end of CoverTab[44150]"
//line /usr/local/go/src/net/http/transport.go:568
	}
//line /usr/local/go/src/net/http/transport.go:568
	// _ = "end of CoverTab[44121]"
//line /usr/local/go/src/net/http/transport.go:568
	_go_fuzz_dep_.CoverTab[44122]++

							for {
//line /usr/local/go/src/net/http/transport.go:570
		_go_fuzz_dep_.CoverTab[44151]++
								select {
		case <-ctx.Done():
//line /usr/local/go/src/net/http/transport.go:572
			_go_fuzz_dep_.CoverTab[44158]++
									req.closeBody()
									return nil, ctx.Err()
//line /usr/local/go/src/net/http/transport.go:574
			// _ = "end of CoverTab[44158]"
		default:
//line /usr/local/go/src/net/http/transport.go:575
			_go_fuzz_dep_.CoverTab[44159]++
//line /usr/local/go/src/net/http/transport.go:575
			// _ = "end of CoverTab[44159]"
		}
//line /usr/local/go/src/net/http/transport.go:576
		// _ = "end of CoverTab[44151]"
//line /usr/local/go/src/net/http/transport.go:576
		_go_fuzz_dep_.CoverTab[44152]++

//line /usr/local/go/src/net/http/transport.go:579
		treq := &transportRequest{Request: req, trace: trace, cancelKey: cancelKey}
		cm, err := t.connectMethodForRequest(treq)
		if err != nil {
//line /usr/local/go/src/net/http/transport.go:581
			_go_fuzz_dep_.CoverTab[44160]++
									req.closeBody()
									return nil, err
//line /usr/local/go/src/net/http/transport.go:583
			// _ = "end of CoverTab[44160]"
		} else {
//line /usr/local/go/src/net/http/transport.go:584
			_go_fuzz_dep_.CoverTab[44161]++
//line /usr/local/go/src/net/http/transport.go:584
			// _ = "end of CoverTab[44161]"
//line /usr/local/go/src/net/http/transport.go:584
		}
//line /usr/local/go/src/net/http/transport.go:584
		// _ = "end of CoverTab[44152]"
//line /usr/local/go/src/net/http/transport.go:584
		_go_fuzz_dep_.CoverTab[44153]++

//line /usr/local/go/src/net/http/transport.go:590
		pconn, err := t.getConn(treq, cm)
		if err != nil {
//line /usr/local/go/src/net/http/transport.go:591
			_go_fuzz_dep_.CoverTab[44162]++
									t.setReqCanceler(cancelKey, nil)
									req.closeBody()
									return nil, err
//line /usr/local/go/src/net/http/transport.go:594
			// _ = "end of CoverTab[44162]"
		} else {
//line /usr/local/go/src/net/http/transport.go:595
			_go_fuzz_dep_.CoverTab[44163]++
//line /usr/local/go/src/net/http/transport.go:595
			// _ = "end of CoverTab[44163]"
//line /usr/local/go/src/net/http/transport.go:595
		}
//line /usr/local/go/src/net/http/transport.go:595
		// _ = "end of CoverTab[44153]"
//line /usr/local/go/src/net/http/transport.go:595
		_go_fuzz_dep_.CoverTab[44154]++

								var resp *Response
								if pconn.alt != nil {
//line /usr/local/go/src/net/http/transport.go:598
			_go_fuzz_dep_.CoverTab[44164]++

									t.setReqCanceler(cancelKey, nil)
									resp, err = pconn.alt.RoundTrip(req)
//line /usr/local/go/src/net/http/transport.go:601
			// _ = "end of CoverTab[44164]"
		} else {
//line /usr/local/go/src/net/http/transport.go:602
			_go_fuzz_dep_.CoverTab[44165]++
									resp, err = pconn.roundTrip(treq)
//line /usr/local/go/src/net/http/transport.go:603
			// _ = "end of CoverTab[44165]"
		}
//line /usr/local/go/src/net/http/transport.go:604
		// _ = "end of CoverTab[44154]"
//line /usr/local/go/src/net/http/transport.go:604
		_go_fuzz_dep_.CoverTab[44155]++
								if err == nil {
//line /usr/local/go/src/net/http/transport.go:605
			_go_fuzz_dep_.CoverTab[44166]++
									resp.Request = origReq
									return resp, nil
//line /usr/local/go/src/net/http/transport.go:607
			// _ = "end of CoverTab[44166]"
		} else {
//line /usr/local/go/src/net/http/transport.go:608
			_go_fuzz_dep_.CoverTab[44167]++
//line /usr/local/go/src/net/http/transport.go:608
			// _ = "end of CoverTab[44167]"
//line /usr/local/go/src/net/http/transport.go:608
		}
//line /usr/local/go/src/net/http/transport.go:608
		// _ = "end of CoverTab[44155]"
//line /usr/local/go/src/net/http/transport.go:608
		_go_fuzz_dep_.CoverTab[44156]++

//line /usr/local/go/src/net/http/transport.go:611
		if http2isNoCachedConnError(err) {
//line /usr/local/go/src/net/http/transport.go:611
			_go_fuzz_dep_.CoverTab[44168]++
									if t.removeIdleConn(pconn) {
//line /usr/local/go/src/net/http/transport.go:612
				_go_fuzz_dep_.CoverTab[44169]++
										t.decConnsPerHost(pconn.cacheKey)
//line /usr/local/go/src/net/http/transport.go:613
				// _ = "end of CoverTab[44169]"
			} else {
//line /usr/local/go/src/net/http/transport.go:614
				_go_fuzz_dep_.CoverTab[44170]++
//line /usr/local/go/src/net/http/transport.go:614
				// _ = "end of CoverTab[44170]"
//line /usr/local/go/src/net/http/transport.go:614
			}
//line /usr/local/go/src/net/http/transport.go:614
			// _ = "end of CoverTab[44168]"
		} else {
//line /usr/local/go/src/net/http/transport.go:615
			_go_fuzz_dep_.CoverTab[44171]++
//line /usr/local/go/src/net/http/transport.go:615
			if !pconn.shouldRetryRequest(req, err) {
//line /usr/local/go/src/net/http/transport.go:615
				_go_fuzz_dep_.CoverTab[44172]++

//line /usr/local/go/src/net/http/transport.go:618
				if e, ok := err.(nothingWrittenError); ok {
//line /usr/local/go/src/net/http/transport.go:618
					_go_fuzz_dep_.CoverTab[44175]++
											err = e.error
//line /usr/local/go/src/net/http/transport.go:619
					// _ = "end of CoverTab[44175]"
				} else {
//line /usr/local/go/src/net/http/transport.go:620
					_go_fuzz_dep_.CoverTab[44176]++
//line /usr/local/go/src/net/http/transport.go:620
					// _ = "end of CoverTab[44176]"
//line /usr/local/go/src/net/http/transport.go:620
				}
//line /usr/local/go/src/net/http/transport.go:620
				// _ = "end of CoverTab[44172]"
//line /usr/local/go/src/net/http/transport.go:620
				_go_fuzz_dep_.CoverTab[44173]++
										if e, ok := err.(transportReadFromServerError); ok {
//line /usr/local/go/src/net/http/transport.go:621
					_go_fuzz_dep_.CoverTab[44177]++
											err = e.err
//line /usr/local/go/src/net/http/transport.go:622
					// _ = "end of CoverTab[44177]"
				} else {
//line /usr/local/go/src/net/http/transport.go:623
					_go_fuzz_dep_.CoverTab[44178]++
//line /usr/local/go/src/net/http/transport.go:623
					// _ = "end of CoverTab[44178]"
//line /usr/local/go/src/net/http/transport.go:623
				}
//line /usr/local/go/src/net/http/transport.go:623
				// _ = "end of CoverTab[44173]"
//line /usr/local/go/src/net/http/transport.go:623
				_go_fuzz_dep_.CoverTab[44174]++
										return nil, err
//line /usr/local/go/src/net/http/transport.go:624
				// _ = "end of CoverTab[44174]"
			} else {
//line /usr/local/go/src/net/http/transport.go:625
				_go_fuzz_dep_.CoverTab[44179]++
//line /usr/local/go/src/net/http/transport.go:625
				// _ = "end of CoverTab[44179]"
//line /usr/local/go/src/net/http/transport.go:625
			}
//line /usr/local/go/src/net/http/transport.go:625
			// _ = "end of CoverTab[44171]"
//line /usr/local/go/src/net/http/transport.go:625
		}
//line /usr/local/go/src/net/http/transport.go:625
		// _ = "end of CoverTab[44156]"
//line /usr/local/go/src/net/http/transport.go:625
		_go_fuzz_dep_.CoverTab[44157]++
								testHookRoundTripRetried()

//line /usr/local/go/src/net/http/transport.go:629
		req, err = rewindBody(req)
		if err != nil {
//line /usr/local/go/src/net/http/transport.go:630
			_go_fuzz_dep_.CoverTab[44180]++
									return nil, err
//line /usr/local/go/src/net/http/transport.go:631
			// _ = "end of CoverTab[44180]"
		} else {
//line /usr/local/go/src/net/http/transport.go:632
			_go_fuzz_dep_.CoverTab[44181]++
//line /usr/local/go/src/net/http/transport.go:632
			// _ = "end of CoverTab[44181]"
//line /usr/local/go/src/net/http/transport.go:632
		}
//line /usr/local/go/src/net/http/transport.go:632
		// _ = "end of CoverTab[44157]"
	}
//line /usr/local/go/src/net/http/transport.go:633
	// _ = "end of CoverTab[44122]"
}

var errCannotRewind = errors.New("net/http: cannot rewind body after connection loss")

type readTrackingBody struct {
	io.ReadCloser
	didRead		bool
	didClose	bool
}

func (r *readTrackingBody) Read(data []byte) (int, error) {
//line /usr/local/go/src/net/http/transport.go:644
	_go_fuzz_dep_.CoverTab[44182]++
							r.didRead = true
							return r.ReadCloser.Read(data)
//line /usr/local/go/src/net/http/transport.go:646
	// _ = "end of CoverTab[44182]"
}

func (r *readTrackingBody) Close() error {
//line /usr/local/go/src/net/http/transport.go:649
	_go_fuzz_dep_.CoverTab[44183]++
							r.didClose = true
							return r.ReadCloser.Close()
//line /usr/local/go/src/net/http/transport.go:651
	// _ = "end of CoverTab[44183]"
}

// setupRewindBody returns a new request with a custom body wrapper
//line /usr/local/go/src/net/http/transport.go:654
// that can report whether the body needs rewinding.
//line /usr/local/go/src/net/http/transport.go:654
// This lets rewindBody avoid an error result when the request
//line /usr/local/go/src/net/http/transport.go:654
// does not have GetBody but the body hasn't been read at all yet.
//line /usr/local/go/src/net/http/transport.go:658
func setupRewindBody(req *Request) *Request {
//line /usr/local/go/src/net/http/transport.go:658
	_go_fuzz_dep_.CoverTab[44184]++
							if req.Body == nil || func() bool {
//line /usr/local/go/src/net/http/transport.go:659
		_go_fuzz_dep_.CoverTab[44186]++
//line /usr/local/go/src/net/http/transport.go:659
		return req.Body == NoBody
//line /usr/local/go/src/net/http/transport.go:659
		// _ = "end of CoverTab[44186]"
//line /usr/local/go/src/net/http/transport.go:659
	}() {
//line /usr/local/go/src/net/http/transport.go:659
		_go_fuzz_dep_.CoverTab[44187]++
								return req
//line /usr/local/go/src/net/http/transport.go:660
		// _ = "end of CoverTab[44187]"
	} else {
//line /usr/local/go/src/net/http/transport.go:661
		_go_fuzz_dep_.CoverTab[44188]++
//line /usr/local/go/src/net/http/transport.go:661
		// _ = "end of CoverTab[44188]"
//line /usr/local/go/src/net/http/transport.go:661
	}
//line /usr/local/go/src/net/http/transport.go:661
	// _ = "end of CoverTab[44184]"
//line /usr/local/go/src/net/http/transport.go:661
	_go_fuzz_dep_.CoverTab[44185]++
							newReq := *req
							newReq.Body = &readTrackingBody{ReadCloser: req.Body}
							return &newReq
//line /usr/local/go/src/net/http/transport.go:664
	// _ = "end of CoverTab[44185]"
}

// rewindBody returns a new request with the body rewound.
//line /usr/local/go/src/net/http/transport.go:667
// It returns req unmodified if the body does not need rewinding.
//line /usr/local/go/src/net/http/transport.go:667
// rewindBody takes care of closing req.Body when appropriate
//line /usr/local/go/src/net/http/transport.go:667
// (in all cases except when rewindBody returns req unmodified).
//line /usr/local/go/src/net/http/transport.go:671
func rewindBody(req *Request) (rewound *Request, err error) {
//line /usr/local/go/src/net/http/transport.go:671
	_go_fuzz_dep_.CoverTab[44189]++
							if req.Body == nil || func() bool {
//line /usr/local/go/src/net/http/transport.go:672
		_go_fuzz_dep_.CoverTab[44194]++
//line /usr/local/go/src/net/http/transport.go:672
		return req.Body == NoBody
//line /usr/local/go/src/net/http/transport.go:672
		// _ = "end of CoverTab[44194]"
//line /usr/local/go/src/net/http/transport.go:672
	}() || func() bool {
//line /usr/local/go/src/net/http/transport.go:672
		_go_fuzz_dep_.CoverTab[44195]++
//line /usr/local/go/src/net/http/transport.go:672
		return (!req.Body.(*readTrackingBody).didRead && func() bool {
//line /usr/local/go/src/net/http/transport.go:672
			_go_fuzz_dep_.CoverTab[44196]++
//line /usr/local/go/src/net/http/transport.go:672
			return !req.Body.(*readTrackingBody).didClose
//line /usr/local/go/src/net/http/transport.go:672
			// _ = "end of CoverTab[44196]"
//line /usr/local/go/src/net/http/transport.go:672
		}())
//line /usr/local/go/src/net/http/transport.go:672
		// _ = "end of CoverTab[44195]"
//line /usr/local/go/src/net/http/transport.go:672
	}() {
//line /usr/local/go/src/net/http/transport.go:672
		_go_fuzz_dep_.CoverTab[44197]++
								return req, nil
//line /usr/local/go/src/net/http/transport.go:673
		// _ = "end of CoverTab[44197]"
	} else {
//line /usr/local/go/src/net/http/transport.go:674
		_go_fuzz_dep_.CoverTab[44198]++
//line /usr/local/go/src/net/http/transport.go:674
		// _ = "end of CoverTab[44198]"
//line /usr/local/go/src/net/http/transport.go:674
	}
//line /usr/local/go/src/net/http/transport.go:674
	// _ = "end of CoverTab[44189]"
//line /usr/local/go/src/net/http/transport.go:674
	_go_fuzz_dep_.CoverTab[44190]++
							if !req.Body.(*readTrackingBody).didClose {
//line /usr/local/go/src/net/http/transport.go:675
		_go_fuzz_dep_.CoverTab[44199]++
								req.closeBody()
//line /usr/local/go/src/net/http/transport.go:676
		// _ = "end of CoverTab[44199]"
	} else {
//line /usr/local/go/src/net/http/transport.go:677
		_go_fuzz_dep_.CoverTab[44200]++
//line /usr/local/go/src/net/http/transport.go:677
		// _ = "end of CoverTab[44200]"
//line /usr/local/go/src/net/http/transport.go:677
	}
//line /usr/local/go/src/net/http/transport.go:677
	// _ = "end of CoverTab[44190]"
//line /usr/local/go/src/net/http/transport.go:677
	_go_fuzz_dep_.CoverTab[44191]++
							if req.GetBody == nil {
//line /usr/local/go/src/net/http/transport.go:678
		_go_fuzz_dep_.CoverTab[44201]++
								return nil, errCannotRewind
//line /usr/local/go/src/net/http/transport.go:679
		// _ = "end of CoverTab[44201]"
	} else {
//line /usr/local/go/src/net/http/transport.go:680
		_go_fuzz_dep_.CoverTab[44202]++
//line /usr/local/go/src/net/http/transport.go:680
		// _ = "end of CoverTab[44202]"
//line /usr/local/go/src/net/http/transport.go:680
	}
//line /usr/local/go/src/net/http/transport.go:680
	// _ = "end of CoverTab[44191]"
//line /usr/local/go/src/net/http/transport.go:680
	_go_fuzz_dep_.CoverTab[44192]++
							body, err := req.GetBody()
							if err != nil {
//line /usr/local/go/src/net/http/transport.go:682
		_go_fuzz_dep_.CoverTab[44203]++
								return nil, err
//line /usr/local/go/src/net/http/transport.go:683
		// _ = "end of CoverTab[44203]"
	} else {
//line /usr/local/go/src/net/http/transport.go:684
		_go_fuzz_dep_.CoverTab[44204]++
//line /usr/local/go/src/net/http/transport.go:684
		// _ = "end of CoverTab[44204]"
//line /usr/local/go/src/net/http/transport.go:684
	}
//line /usr/local/go/src/net/http/transport.go:684
	// _ = "end of CoverTab[44192]"
//line /usr/local/go/src/net/http/transport.go:684
	_go_fuzz_dep_.CoverTab[44193]++
							newReq := *req
							newReq.Body = &readTrackingBody{ReadCloser: body}
							return &newReq, nil
//line /usr/local/go/src/net/http/transport.go:687
	// _ = "end of CoverTab[44193]"
}

// shouldRetryRequest reports whether we should retry sending a failed
//line /usr/local/go/src/net/http/transport.go:690
// HTTP request on a new connection. The non-nil input error is the
//line /usr/local/go/src/net/http/transport.go:690
// error from roundTrip.
//line /usr/local/go/src/net/http/transport.go:693
func (pc *persistConn) shouldRetryRequest(req *Request, err error) bool {
//line /usr/local/go/src/net/http/transport.go:693
	_go_fuzz_dep_.CoverTab[44205]++
							if http2isNoCachedConnError(err) {
//line /usr/local/go/src/net/http/transport.go:694
		_go_fuzz_dep_.CoverTab[44213]++

//line /usr/local/go/src/net/http/transport.go:701
		return true
//line /usr/local/go/src/net/http/transport.go:701
		// _ = "end of CoverTab[44213]"
	} else {
//line /usr/local/go/src/net/http/transport.go:702
		_go_fuzz_dep_.CoverTab[44214]++
//line /usr/local/go/src/net/http/transport.go:702
		// _ = "end of CoverTab[44214]"
//line /usr/local/go/src/net/http/transport.go:702
	}
//line /usr/local/go/src/net/http/transport.go:702
	// _ = "end of CoverTab[44205]"
//line /usr/local/go/src/net/http/transport.go:702
	_go_fuzz_dep_.CoverTab[44206]++
							if err == errMissingHost {
//line /usr/local/go/src/net/http/transport.go:703
		_go_fuzz_dep_.CoverTab[44215]++

								return false
//line /usr/local/go/src/net/http/transport.go:705
		// _ = "end of CoverTab[44215]"
	} else {
//line /usr/local/go/src/net/http/transport.go:706
		_go_fuzz_dep_.CoverTab[44216]++
//line /usr/local/go/src/net/http/transport.go:706
		// _ = "end of CoverTab[44216]"
//line /usr/local/go/src/net/http/transport.go:706
	}
//line /usr/local/go/src/net/http/transport.go:706
	// _ = "end of CoverTab[44206]"
//line /usr/local/go/src/net/http/transport.go:706
	_go_fuzz_dep_.CoverTab[44207]++
							if !pc.isReused() {
//line /usr/local/go/src/net/http/transport.go:707
		_go_fuzz_dep_.CoverTab[44217]++

//line /usr/local/go/src/net/http/transport.go:715
		return false
//line /usr/local/go/src/net/http/transport.go:715
		// _ = "end of CoverTab[44217]"
	} else {
//line /usr/local/go/src/net/http/transport.go:716
		_go_fuzz_dep_.CoverTab[44218]++
//line /usr/local/go/src/net/http/transport.go:716
		// _ = "end of CoverTab[44218]"
//line /usr/local/go/src/net/http/transport.go:716
	}
//line /usr/local/go/src/net/http/transport.go:716
	// _ = "end of CoverTab[44207]"
//line /usr/local/go/src/net/http/transport.go:716
	_go_fuzz_dep_.CoverTab[44208]++
							if _, ok := err.(nothingWrittenError); ok {
//line /usr/local/go/src/net/http/transport.go:717
		_go_fuzz_dep_.CoverTab[44219]++

//line /usr/local/go/src/net/http/transport.go:720
		return req.outgoingLength() == 0 || func() bool {
//line /usr/local/go/src/net/http/transport.go:720
			_go_fuzz_dep_.CoverTab[44220]++
//line /usr/local/go/src/net/http/transport.go:720
			return req.GetBody != nil
//line /usr/local/go/src/net/http/transport.go:720
			// _ = "end of CoverTab[44220]"
//line /usr/local/go/src/net/http/transport.go:720
		}()
//line /usr/local/go/src/net/http/transport.go:720
		// _ = "end of CoverTab[44219]"
	} else {
//line /usr/local/go/src/net/http/transport.go:721
		_go_fuzz_dep_.CoverTab[44221]++
//line /usr/local/go/src/net/http/transport.go:721
		// _ = "end of CoverTab[44221]"
//line /usr/local/go/src/net/http/transport.go:721
	}
//line /usr/local/go/src/net/http/transport.go:721
	// _ = "end of CoverTab[44208]"
//line /usr/local/go/src/net/http/transport.go:721
	_go_fuzz_dep_.CoverTab[44209]++
							if !req.isReplayable() {
//line /usr/local/go/src/net/http/transport.go:722
		_go_fuzz_dep_.CoverTab[44222]++

								return false
//line /usr/local/go/src/net/http/transport.go:724
		// _ = "end of CoverTab[44222]"
	} else {
//line /usr/local/go/src/net/http/transport.go:725
		_go_fuzz_dep_.CoverTab[44223]++
//line /usr/local/go/src/net/http/transport.go:725
		// _ = "end of CoverTab[44223]"
//line /usr/local/go/src/net/http/transport.go:725
	}
//line /usr/local/go/src/net/http/transport.go:725
	// _ = "end of CoverTab[44209]"
//line /usr/local/go/src/net/http/transport.go:725
	_go_fuzz_dep_.CoverTab[44210]++
							if _, ok := err.(transportReadFromServerError); ok {
//line /usr/local/go/src/net/http/transport.go:726
		_go_fuzz_dep_.CoverTab[44224]++

//line /usr/local/go/src/net/http/transport.go:729
		return true
//line /usr/local/go/src/net/http/transport.go:729
		// _ = "end of CoverTab[44224]"
	} else {
//line /usr/local/go/src/net/http/transport.go:730
		_go_fuzz_dep_.CoverTab[44225]++
//line /usr/local/go/src/net/http/transport.go:730
		// _ = "end of CoverTab[44225]"
//line /usr/local/go/src/net/http/transport.go:730
	}
//line /usr/local/go/src/net/http/transport.go:730
	// _ = "end of CoverTab[44210]"
//line /usr/local/go/src/net/http/transport.go:730
	_go_fuzz_dep_.CoverTab[44211]++
							if err == errServerClosedIdle {
//line /usr/local/go/src/net/http/transport.go:731
		_go_fuzz_dep_.CoverTab[44226]++

//line /usr/local/go/src/net/http/transport.go:735
		return true
//line /usr/local/go/src/net/http/transport.go:735
		// _ = "end of CoverTab[44226]"
	} else {
//line /usr/local/go/src/net/http/transport.go:736
		_go_fuzz_dep_.CoverTab[44227]++
//line /usr/local/go/src/net/http/transport.go:736
		// _ = "end of CoverTab[44227]"
//line /usr/local/go/src/net/http/transport.go:736
	}
//line /usr/local/go/src/net/http/transport.go:736
	// _ = "end of CoverTab[44211]"
//line /usr/local/go/src/net/http/transport.go:736
	_go_fuzz_dep_.CoverTab[44212]++
							return false
//line /usr/local/go/src/net/http/transport.go:737
	// _ = "end of CoverTab[44212]"
}

// ErrSkipAltProtocol is a sentinel error value defined by Transport.RegisterProtocol.
var ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")

// RegisterProtocol registers a new protocol with scheme.
//line /usr/local/go/src/net/http/transport.go:743
// The Transport will pass requests using the given scheme to rt.
//line /usr/local/go/src/net/http/transport.go:743
// It is rt's responsibility to simulate HTTP request semantics.
//line /usr/local/go/src/net/http/transport.go:743
//
//line /usr/local/go/src/net/http/transport.go:743
// RegisterProtocol can be used by other packages to provide
//line /usr/local/go/src/net/http/transport.go:743
// implementations of protocol schemes like "ftp" or "file".
//line /usr/local/go/src/net/http/transport.go:743
//
//line /usr/local/go/src/net/http/transport.go:743
// If rt.RoundTrip returns ErrSkipAltProtocol, the Transport will
//line /usr/local/go/src/net/http/transport.go:743
// handle the RoundTrip itself for that one request, as if the
//line /usr/local/go/src/net/http/transport.go:743
// protocol were not registered.
//line /usr/local/go/src/net/http/transport.go:753
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper) {
//line /usr/local/go/src/net/http/transport.go:753
	_go_fuzz_dep_.CoverTab[44228]++
							t.altMu.Lock()
							defer t.altMu.Unlock()
							oldMap, _ := t.altProto.Load().(map[string]RoundTripper)
							if _, exists := oldMap[scheme]; exists {
//line /usr/local/go/src/net/http/transport.go:757
		_go_fuzz_dep_.CoverTab[44231]++
								panic("protocol " + scheme + " already registered")
//line /usr/local/go/src/net/http/transport.go:758
		// _ = "end of CoverTab[44231]"
	} else {
//line /usr/local/go/src/net/http/transport.go:759
		_go_fuzz_dep_.CoverTab[44232]++
//line /usr/local/go/src/net/http/transport.go:759
		// _ = "end of CoverTab[44232]"
//line /usr/local/go/src/net/http/transport.go:759
	}
//line /usr/local/go/src/net/http/transport.go:759
	// _ = "end of CoverTab[44228]"
//line /usr/local/go/src/net/http/transport.go:759
	_go_fuzz_dep_.CoverTab[44229]++
							newMap := make(map[string]RoundTripper)
							for k, v := range oldMap {
//line /usr/local/go/src/net/http/transport.go:761
		_go_fuzz_dep_.CoverTab[44233]++
								newMap[k] = v
//line /usr/local/go/src/net/http/transport.go:762
		// _ = "end of CoverTab[44233]"
	}
//line /usr/local/go/src/net/http/transport.go:763
	// _ = "end of CoverTab[44229]"
//line /usr/local/go/src/net/http/transport.go:763
	_go_fuzz_dep_.CoverTab[44230]++
							newMap[scheme] = rt
							t.altProto.Store(newMap)
//line /usr/local/go/src/net/http/transport.go:765
	// _ = "end of CoverTab[44230]"
}

// CloseIdleConnections closes any connections which were previously
//line /usr/local/go/src/net/http/transport.go:768
// connected from previous requests but are now sitting idle in
//line /usr/local/go/src/net/http/transport.go:768
// a "keep-alive" state. It does not interrupt any connections currently
//line /usr/local/go/src/net/http/transport.go:768
// in use.
//line /usr/local/go/src/net/http/transport.go:772
func (t *Transport) CloseIdleConnections() {
//line /usr/local/go/src/net/http/transport.go:772
	_go_fuzz_dep_.CoverTab[44234]++
							t.nextProtoOnce.Do(t.onceSetNextProtoDefaults)
							t.idleMu.Lock()
							m := t.idleConn
							t.idleConn = nil
							t.closeIdle = true
							t.idleLRU = connLRU{}
							t.idleMu.Unlock()
							for _, conns := range m {
//line /usr/local/go/src/net/http/transport.go:780
		_go_fuzz_dep_.CoverTab[44236]++
								for _, pconn := range conns {
//line /usr/local/go/src/net/http/transport.go:781
			_go_fuzz_dep_.CoverTab[44237]++
									pconn.close(errCloseIdleConns)
//line /usr/local/go/src/net/http/transport.go:782
			// _ = "end of CoverTab[44237]"
		}
//line /usr/local/go/src/net/http/transport.go:783
		// _ = "end of CoverTab[44236]"
	}
//line /usr/local/go/src/net/http/transport.go:784
	// _ = "end of CoverTab[44234]"
//line /usr/local/go/src/net/http/transport.go:784
	_go_fuzz_dep_.CoverTab[44235]++
							if t2 := t.h2transport; t2 != nil {
//line /usr/local/go/src/net/http/transport.go:785
		_go_fuzz_dep_.CoverTab[44238]++
								t2.CloseIdleConnections()
//line /usr/local/go/src/net/http/transport.go:786
		// _ = "end of CoverTab[44238]"
	} else {
//line /usr/local/go/src/net/http/transport.go:787
		_go_fuzz_dep_.CoverTab[44239]++
//line /usr/local/go/src/net/http/transport.go:787
		// _ = "end of CoverTab[44239]"
//line /usr/local/go/src/net/http/transport.go:787
	}
//line /usr/local/go/src/net/http/transport.go:787
	// _ = "end of CoverTab[44235]"
}

// CancelRequest cancels an in-flight request by closing its connection.
//line /usr/local/go/src/net/http/transport.go:790
// CancelRequest should only be called after RoundTrip has returned.
//line /usr/local/go/src/net/http/transport.go:790
//
//line /usr/local/go/src/net/http/transport.go:790
// Deprecated: Use Request.WithContext to create a request with a
//line /usr/local/go/src/net/http/transport.go:790
// cancelable context instead. CancelRequest cannot cancel HTTP/2
//line /usr/local/go/src/net/http/transport.go:790
// requests.
//line /usr/local/go/src/net/http/transport.go:796
func (t *Transport) CancelRequest(req *Request) {
//line /usr/local/go/src/net/http/transport.go:796
	_go_fuzz_dep_.CoverTab[44240]++
							t.cancelRequest(cancelKey{req}, errRequestCanceled)
//line /usr/local/go/src/net/http/transport.go:797
	// _ = "end of CoverTab[44240]"
}

// Cancel an in-flight request, recording the error value.
//line /usr/local/go/src/net/http/transport.go:800
// Returns whether the request was canceled.
//line /usr/local/go/src/net/http/transport.go:802
func (t *Transport) cancelRequest(key cancelKey, err error) bool {
//line /usr/local/go/src/net/http/transport.go:802
	_go_fuzz_dep_.CoverTab[44241]++

//line /usr/local/go/src/net/http/transport.go:805
	t.reqMu.Lock()
	defer t.reqMu.Unlock()
	cancel := t.reqCanceler[key]
	delete(t.reqCanceler, key)
	if cancel != nil {
//line /usr/local/go/src/net/http/transport.go:809
		_go_fuzz_dep_.CoverTab[44243]++
								cancel(err)
//line /usr/local/go/src/net/http/transport.go:810
		// _ = "end of CoverTab[44243]"
	} else {
//line /usr/local/go/src/net/http/transport.go:811
		_go_fuzz_dep_.CoverTab[44244]++
//line /usr/local/go/src/net/http/transport.go:811
		// _ = "end of CoverTab[44244]"
//line /usr/local/go/src/net/http/transport.go:811
	}
//line /usr/local/go/src/net/http/transport.go:811
	// _ = "end of CoverTab[44241]"
//line /usr/local/go/src/net/http/transport.go:811
	_go_fuzz_dep_.CoverTab[44242]++

							return cancel != nil
//line /usr/local/go/src/net/http/transport.go:813
	// _ = "end of CoverTab[44242]"
}

//line /usr/local/go/src/net/http/transport.go:820
var (
	envProxyOnce		sync.Once
	envProxyFuncValue	func(*url.URL) (*url.URL, error)
)

// envProxyFunc returns a function that reads the
//line /usr/local/go/src/net/http/transport.go:825
// environment variable to determine the proxy address.
//line /usr/local/go/src/net/http/transport.go:827
func envProxyFunc() func(*url.URL) (*url.URL, error) {
//line /usr/local/go/src/net/http/transport.go:827
	_go_fuzz_dep_.CoverTab[44245]++
							envProxyOnce.Do(func() {
//line /usr/local/go/src/net/http/transport.go:828
		_go_fuzz_dep_.CoverTab[44247]++
								envProxyFuncValue = httpproxy.FromEnvironment().ProxyFunc()
//line /usr/local/go/src/net/http/transport.go:829
		// _ = "end of CoverTab[44247]"
	})
//line /usr/local/go/src/net/http/transport.go:830
	// _ = "end of CoverTab[44245]"
//line /usr/local/go/src/net/http/transport.go:830
	_go_fuzz_dep_.CoverTab[44246]++
							return envProxyFuncValue
//line /usr/local/go/src/net/http/transport.go:831
	// _ = "end of CoverTab[44246]"
}

// resetProxyConfig is used by tests.
func resetProxyConfig() {
//line /usr/local/go/src/net/http/transport.go:835
	_go_fuzz_dep_.CoverTab[44248]++
							envProxyOnce = sync.Once{}
							envProxyFuncValue = nil
//line /usr/local/go/src/net/http/transport.go:837
	// _ = "end of CoverTab[44248]"
}

func (t *Transport) connectMethodForRequest(treq *transportRequest) (cm connectMethod, err error) {
//line /usr/local/go/src/net/http/transport.go:840
	_go_fuzz_dep_.CoverTab[44249]++
							cm.targetScheme = treq.URL.Scheme
							cm.targetAddr = canonicalAddr(treq.URL)
							if t.Proxy != nil {
//line /usr/local/go/src/net/http/transport.go:843
		_go_fuzz_dep_.CoverTab[44251]++
								cm.proxyURL, err = t.Proxy(treq.Request)
//line /usr/local/go/src/net/http/transport.go:844
		// _ = "end of CoverTab[44251]"
	} else {
//line /usr/local/go/src/net/http/transport.go:845
		_go_fuzz_dep_.CoverTab[44252]++
//line /usr/local/go/src/net/http/transport.go:845
		// _ = "end of CoverTab[44252]"
//line /usr/local/go/src/net/http/transport.go:845
	}
//line /usr/local/go/src/net/http/transport.go:845
	// _ = "end of CoverTab[44249]"
//line /usr/local/go/src/net/http/transport.go:845
	_go_fuzz_dep_.CoverTab[44250]++
							cm.onlyH1 = treq.requiresHTTP1()
							return cm, err
//line /usr/local/go/src/net/http/transport.go:847
	// _ = "end of CoverTab[44250]"
}

// proxyAuth returns the Proxy-Authorization header to set
//line /usr/local/go/src/net/http/transport.go:850
// on requests, if applicable.
//line /usr/local/go/src/net/http/transport.go:852
func (cm *connectMethod) proxyAuth() string {
//line /usr/local/go/src/net/http/transport.go:852
	_go_fuzz_dep_.CoverTab[44253]++
							if cm.proxyURL == nil {
//line /usr/local/go/src/net/http/transport.go:853
		_go_fuzz_dep_.CoverTab[44256]++
								return ""
//line /usr/local/go/src/net/http/transport.go:854
		// _ = "end of CoverTab[44256]"
	} else {
//line /usr/local/go/src/net/http/transport.go:855
		_go_fuzz_dep_.CoverTab[44257]++
//line /usr/local/go/src/net/http/transport.go:855
		// _ = "end of CoverTab[44257]"
//line /usr/local/go/src/net/http/transport.go:855
	}
//line /usr/local/go/src/net/http/transport.go:855
	// _ = "end of CoverTab[44253]"
//line /usr/local/go/src/net/http/transport.go:855
	_go_fuzz_dep_.CoverTab[44254]++
							if u := cm.proxyURL.User; u != nil {
//line /usr/local/go/src/net/http/transport.go:856
		_go_fuzz_dep_.CoverTab[44258]++
								username := u.Username()
								password, _ := u.Password()
								return "Basic " + basicAuth(username, password)
//line /usr/local/go/src/net/http/transport.go:859
		// _ = "end of CoverTab[44258]"
	} else {
//line /usr/local/go/src/net/http/transport.go:860
		_go_fuzz_dep_.CoverTab[44259]++
//line /usr/local/go/src/net/http/transport.go:860
		// _ = "end of CoverTab[44259]"
//line /usr/local/go/src/net/http/transport.go:860
	}
//line /usr/local/go/src/net/http/transport.go:860
	// _ = "end of CoverTab[44254]"
//line /usr/local/go/src/net/http/transport.go:860
	_go_fuzz_dep_.CoverTab[44255]++
							return ""
//line /usr/local/go/src/net/http/transport.go:861
	// _ = "end of CoverTab[44255]"
}

// error values for debugging and testing, not seen by users.
var (
	errKeepAlivesDisabled	= errors.New("http: putIdleConn: keep alives disabled")
	errConnBroken		= errors.New("http: putIdleConn: connection is in bad state")
	errCloseIdle		= errors.New("http: putIdleConn: CloseIdleConnections was called")
	errTooManyIdle		= errors.New("http: putIdleConn: too many idle connections")
	errTooManyIdleHost	= errors.New("http: putIdleConn: too many idle connections for host")
	errCloseIdleConns	= errors.New("http: CloseIdleConnections called")
	errReadLoopExiting	= errors.New("http: persistConn.readLoop exiting")
	errIdleConnTimeout	= errors.New("http: idle connection timeout")

	// errServerClosedIdle is not seen by users for idempotent requests, but may be
	// seen by a user if the server shuts down an idle connection and sends its FIN
	// in flight with already-written POST body bytes from the client.
	// See https://github.com/golang/go/issues/19943#issuecomment-355607646
	errServerClosedIdle	= errors.New("http: server closed idle connection")
)

// transportReadFromServerError is used by Transport.readLoop when the
//line /usr/local/go/src/net/http/transport.go:882
// 1 byte peek read fails and we're actually anticipating a response.
//line /usr/local/go/src/net/http/transport.go:882
// Usually this is just due to the inherent keep-alive shut down race,
//line /usr/local/go/src/net/http/transport.go:882
// where the server closed the connection at the same time the client
//line /usr/local/go/src/net/http/transport.go:882
// wrote. The underlying err field is usually io.EOF or some
//line /usr/local/go/src/net/http/transport.go:882
// ECONNRESET sort of thing which varies by platform. But it might be
//line /usr/local/go/src/net/http/transport.go:882
// the user's custom net.Conn.Read error too, so we carry it along for
//line /usr/local/go/src/net/http/transport.go:882
// them to return from Transport.RoundTrip.
//line /usr/local/go/src/net/http/transport.go:890
type transportReadFromServerError struct {
	err error
}

func (e transportReadFromServerError) Unwrap() error {
//line /usr/local/go/src/net/http/transport.go:894
	_go_fuzz_dep_.CoverTab[44260]++
//line /usr/local/go/src/net/http/transport.go:894
	return e.err
//line /usr/local/go/src/net/http/transport.go:894
	// _ = "end of CoverTab[44260]"
//line /usr/local/go/src/net/http/transport.go:894
}

func (e transportReadFromServerError) Error() string {
//line /usr/local/go/src/net/http/transport.go:896
	_go_fuzz_dep_.CoverTab[44261]++
							return fmt.Sprintf("net/http: Transport failed to read from server: %v", e.err)
//line /usr/local/go/src/net/http/transport.go:897
	// _ = "end of CoverTab[44261]"
}

func (t *Transport) putOrCloseIdleConn(pconn *persistConn) {
//line /usr/local/go/src/net/http/transport.go:900
	_go_fuzz_dep_.CoverTab[44262]++
							if err := t.tryPutIdleConn(pconn); err != nil {
//line /usr/local/go/src/net/http/transport.go:901
		_go_fuzz_dep_.CoverTab[44263]++
								pconn.close(err)
//line /usr/local/go/src/net/http/transport.go:902
		// _ = "end of CoverTab[44263]"
	} else {
//line /usr/local/go/src/net/http/transport.go:903
		_go_fuzz_dep_.CoverTab[44264]++
//line /usr/local/go/src/net/http/transport.go:903
		// _ = "end of CoverTab[44264]"
//line /usr/local/go/src/net/http/transport.go:903
	}
//line /usr/local/go/src/net/http/transport.go:903
	// _ = "end of CoverTab[44262]"
}

func (t *Transport) maxIdleConnsPerHost() int {
//line /usr/local/go/src/net/http/transport.go:906
	_go_fuzz_dep_.CoverTab[44265]++
							if v := t.MaxIdleConnsPerHost; v != 0 {
//line /usr/local/go/src/net/http/transport.go:907
		_go_fuzz_dep_.CoverTab[44267]++
								return v
//line /usr/local/go/src/net/http/transport.go:908
		// _ = "end of CoverTab[44267]"
	} else {
//line /usr/local/go/src/net/http/transport.go:909
		_go_fuzz_dep_.CoverTab[44268]++
//line /usr/local/go/src/net/http/transport.go:909
		// _ = "end of CoverTab[44268]"
//line /usr/local/go/src/net/http/transport.go:909
	}
//line /usr/local/go/src/net/http/transport.go:909
	// _ = "end of CoverTab[44265]"
//line /usr/local/go/src/net/http/transport.go:909
	_go_fuzz_dep_.CoverTab[44266]++
							return DefaultMaxIdleConnsPerHost
//line /usr/local/go/src/net/http/transport.go:910
	// _ = "end of CoverTab[44266]"
}

// tryPutIdleConn adds pconn to the list of idle persistent connections awaiting
//line /usr/local/go/src/net/http/transport.go:913
// a new request.
//line /usr/local/go/src/net/http/transport.go:913
// If pconn is no longer needed or not in a good state, tryPutIdleConn returns
//line /usr/local/go/src/net/http/transport.go:913
// an error explaining why it wasn't registered.
//line /usr/local/go/src/net/http/transport.go:913
// tryPutIdleConn does not close pconn. Use putOrCloseIdleConn instead for that.
//line /usr/local/go/src/net/http/transport.go:918
func (t *Transport) tryPutIdleConn(pconn *persistConn) error {
//line /usr/local/go/src/net/http/transport.go:918
	_go_fuzz_dep_.CoverTab[44269]++
							if t.DisableKeepAlives || func() bool {
//line /usr/local/go/src/net/http/transport.go:919
		_go_fuzz_dep_.CoverTab[44280]++
//line /usr/local/go/src/net/http/transport.go:919
		return t.MaxIdleConnsPerHost < 0
//line /usr/local/go/src/net/http/transport.go:919
		// _ = "end of CoverTab[44280]"
//line /usr/local/go/src/net/http/transport.go:919
	}() {
//line /usr/local/go/src/net/http/transport.go:919
		_go_fuzz_dep_.CoverTab[44281]++
								return errKeepAlivesDisabled
//line /usr/local/go/src/net/http/transport.go:920
		// _ = "end of CoverTab[44281]"
	} else {
//line /usr/local/go/src/net/http/transport.go:921
		_go_fuzz_dep_.CoverTab[44282]++
//line /usr/local/go/src/net/http/transport.go:921
		// _ = "end of CoverTab[44282]"
//line /usr/local/go/src/net/http/transport.go:921
	}
//line /usr/local/go/src/net/http/transport.go:921
	// _ = "end of CoverTab[44269]"
//line /usr/local/go/src/net/http/transport.go:921
	_go_fuzz_dep_.CoverTab[44270]++
							if pconn.isBroken() {
//line /usr/local/go/src/net/http/transport.go:922
		_go_fuzz_dep_.CoverTab[44283]++
								return errConnBroken
//line /usr/local/go/src/net/http/transport.go:923
		// _ = "end of CoverTab[44283]"
	} else {
//line /usr/local/go/src/net/http/transport.go:924
		_go_fuzz_dep_.CoverTab[44284]++
//line /usr/local/go/src/net/http/transport.go:924
		// _ = "end of CoverTab[44284]"
//line /usr/local/go/src/net/http/transport.go:924
	}
//line /usr/local/go/src/net/http/transport.go:924
	// _ = "end of CoverTab[44270]"
//line /usr/local/go/src/net/http/transport.go:924
	_go_fuzz_dep_.CoverTab[44271]++
							pconn.markReused()

							t.idleMu.Lock()
							defer t.idleMu.Unlock()

//line /usr/local/go/src/net/http/transport.go:933
	if pconn.alt != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:933
		_go_fuzz_dep_.CoverTab[44285]++
//line /usr/local/go/src/net/http/transport.go:933
		return t.idleLRU.m[pconn] != nil
//line /usr/local/go/src/net/http/transport.go:933
		// _ = "end of CoverTab[44285]"
//line /usr/local/go/src/net/http/transport.go:933
	}() {
//line /usr/local/go/src/net/http/transport.go:933
		_go_fuzz_dep_.CoverTab[44286]++
								return nil
//line /usr/local/go/src/net/http/transport.go:934
		// _ = "end of CoverTab[44286]"
	} else {
//line /usr/local/go/src/net/http/transport.go:935
		_go_fuzz_dep_.CoverTab[44287]++
//line /usr/local/go/src/net/http/transport.go:935
		// _ = "end of CoverTab[44287]"
//line /usr/local/go/src/net/http/transport.go:935
	}
//line /usr/local/go/src/net/http/transport.go:935
	// _ = "end of CoverTab[44271]"
//line /usr/local/go/src/net/http/transport.go:935
	_go_fuzz_dep_.CoverTab[44272]++

//line /usr/local/go/src/net/http/transport.go:941
	key := pconn.cacheKey
	if q, ok := t.idleConnWait[key]; ok {
//line /usr/local/go/src/net/http/transport.go:942
		_go_fuzz_dep_.CoverTab[44288]++
								done := false
								if pconn.alt == nil {
//line /usr/local/go/src/net/http/transport.go:944
			_go_fuzz_dep_.CoverTab[44291]++

//line /usr/local/go/src/net/http/transport.go:947
			for q.len() > 0 {
//line /usr/local/go/src/net/http/transport.go:947
				_go_fuzz_dep_.CoverTab[44292]++
										w := q.popFront()
										if w.tryDeliver(pconn, nil) {
//line /usr/local/go/src/net/http/transport.go:949
					_go_fuzz_dep_.CoverTab[44293]++
											done = true
											break
//line /usr/local/go/src/net/http/transport.go:951
					// _ = "end of CoverTab[44293]"
				} else {
//line /usr/local/go/src/net/http/transport.go:952
					_go_fuzz_dep_.CoverTab[44294]++
//line /usr/local/go/src/net/http/transport.go:952
					// _ = "end of CoverTab[44294]"
//line /usr/local/go/src/net/http/transport.go:952
				}
//line /usr/local/go/src/net/http/transport.go:952
				// _ = "end of CoverTab[44292]"
			}
//line /usr/local/go/src/net/http/transport.go:953
			// _ = "end of CoverTab[44291]"
		} else {
//line /usr/local/go/src/net/http/transport.go:954
			_go_fuzz_dep_.CoverTab[44295]++

//line /usr/local/go/src/net/http/transport.go:959
			for q.len() > 0 {
//line /usr/local/go/src/net/http/transport.go:959
				_go_fuzz_dep_.CoverTab[44296]++
										w := q.popFront()
										w.tryDeliver(pconn, nil)
//line /usr/local/go/src/net/http/transport.go:961
				// _ = "end of CoverTab[44296]"
			}
//line /usr/local/go/src/net/http/transport.go:962
			// _ = "end of CoverTab[44295]"
		}
//line /usr/local/go/src/net/http/transport.go:963
		// _ = "end of CoverTab[44288]"
//line /usr/local/go/src/net/http/transport.go:963
		_go_fuzz_dep_.CoverTab[44289]++
								if q.len() == 0 {
//line /usr/local/go/src/net/http/transport.go:964
			_go_fuzz_dep_.CoverTab[44297]++
									delete(t.idleConnWait, key)
//line /usr/local/go/src/net/http/transport.go:965
			// _ = "end of CoverTab[44297]"
		} else {
//line /usr/local/go/src/net/http/transport.go:966
			_go_fuzz_dep_.CoverTab[44298]++
									t.idleConnWait[key] = q
//line /usr/local/go/src/net/http/transport.go:967
			// _ = "end of CoverTab[44298]"
		}
//line /usr/local/go/src/net/http/transport.go:968
		// _ = "end of CoverTab[44289]"
//line /usr/local/go/src/net/http/transport.go:968
		_go_fuzz_dep_.CoverTab[44290]++
								if done {
//line /usr/local/go/src/net/http/transport.go:969
			_go_fuzz_dep_.CoverTab[44299]++
									return nil
//line /usr/local/go/src/net/http/transport.go:970
			// _ = "end of CoverTab[44299]"
		} else {
//line /usr/local/go/src/net/http/transport.go:971
			_go_fuzz_dep_.CoverTab[44300]++
//line /usr/local/go/src/net/http/transport.go:971
			// _ = "end of CoverTab[44300]"
//line /usr/local/go/src/net/http/transport.go:971
		}
//line /usr/local/go/src/net/http/transport.go:971
		// _ = "end of CoverTab[44290]"
	} else {
//line /usr/local/go/src/net/http/transport.go:972
		_go_fuzz_dep_.CoverTab[44301]++
//line /usr/local/go/src/net/http/transport.go:972
		// _ = "end of CoverTab[44301]"
//line /usr/local/go/src/net/http/transport.go:972
	}
//line /usr/local/go/src/net/http/transport.go:972
	// _ = "end of CoverTab[44272]"
//line /usr/local/go/src/net/http/transport.go:972
	_go_fuzz_dep_.CoverTab[44273]++

							if t.closeIdle {
//line /usr/local/go/src/net/http/transport.go:974
		_go_fuzz_dep_.CoverTab[44302]++
								return errCloseIdle
//line /usr/local/go/src/net/http/transport.go:975
		// _ = "end of CoverTab[44302]"
	} else {
//line /usr/local/go/src/net/http/transport.go:976
		_go_fuzz_dep_.CoverTab[44303]++
//line /usr/local/go/src/net/http/transport.go:976
		// _ = "end of CoverTab[44303]"
//line /usr/local/go/src/net/http/transport.go:976
	}
//line /usr/local/go/src/net/http/transport.go:976
	// _ = "end of CoverTab[44273]"
//line /usr/local/go/src/net/http/transport.go:976
	_go_fuzz_dep_.CoverTab[44274]++
							if t.idleConn == nil {
//line /usr/local/go/src/net/http/transport.go:977
		_go_fuzz_dep_.CoverTab[44304]++
								t.idleConn = make(map[connectMethodKey][]*persistConn)
//line /usr/local/go/src/net/http/transport.go:978
		// _ = "end of CoverTab[44304]"
	} else {
//line /usr/local/go/src/net/http/transport.go:979
		_go_fuzz_dep_.CoverTab[44305]++
//line /usr/local/go/src/net/http/transport.go:979
		// _ = "end of CoverTab[44305]"
//line /usr/local/go/src/net/http/transport.go:979
	}
//line /usr/local/go/src/net/http/transport.go:979
	// _ = "end of CoverTab[44274]"
//line /usr/local/go/src/net/http/transport.go:979
	_go_fuzz_dep_.CoverTab[44275]++
							idles := t.idleConn[key]
							if len(idles) >= t.maxIdleConnsPerHost() {
//line /usr/local/go/src/net/http/transport.go:981
		_go_fuzz_dep_.CoverTab[44306]++
								return errTooManyIdleHost
//line /usr/local/go/src/net/http/transport.go:982
		// _ = "end of CoverTab[44306]"
	} else {
//line /usr/local/go/src/net/http/transport.go:983
		_go_fuzz_dep_.CoverTab[44307]++
//line /usr/local/go/src/net/http/transport.go:983
		// _ = "end of CoverTab[44307]"
//line /usr/local/go/src/net/http/transport.go:983
	}
//line /usr/local/go/src/net/http/transport.go:983
	// _ = "end of CoverTab[44275]"
//line /usr/local/go/src/net/http/transport.go:983
	_go_fuzz_dep_.CoverTab[44276]++
							for _, exist := range idles {
//line /usr/local/go/src/net/http/transport.go:984
		_go_fuzz_dep_.CoverTab[44308]++
								if exist == pconn {
//line /usr/local/go/src/net/http/transport.go:985
			_go_fuzz_dep_.CoverTab[44309]++
									log.Fatalf("dup idle pconn %p in freelist", pconn)
//line /usr/local/go/src/net/http/transport.go:986
			// _ = "end of CoverTab[44309]"
		} else {
//line /usr/local/go/src/net/http/transport.go:987
			_go_fuzz_dep_.CoverTab[44310]++
//line /usr/local/go/src/net/http/transport.go:987
			// _ = "end of CoverTab[44310]"
//line /usr/local/go/src/net/http/transport.go:987
		}
//line /usr/local/go/src/net/http/transport.go:987
		// _ = "end of CoverTab[44308]"
	}
//line /usr/local/go/src/net/http/transport.go:988
	// _ = "end of CoverTab[44276]"
//line /usr/local/go/src/net/http/transport.go:988
	_go_fuzz_dep_.CoverTab[44277]++
							t.idleConn[key] = append(idles, pconn)
							t.idleLRU.add(pconn)
							if t.MaxIdleConns != 0 && func() bool {
//line /usr/local/go/src/net/http/transport.go:991
		_go_fuzz_dep_.CoverTab[44311]++
//line /usr/local/go/src/net/http/transport.go:991
		return t.idleLRU.len() > t.MaxIdleConns
//line /usr/local/go/src/net/http/transport.go:991
		// _ = "end of CoverTab[44311]"
//line /usr/local/go/src/net/http/transport.go:991
	}() {
//line /usr/local/go/src/net/http/transport.go:991
		_go_fuzz_dep_.CoverTab[44312]++
								oldest := t.idleLRU.removeOldest()
								oldest.close(errTooManyIdle)
								t.removeIdleConnLocked(oldest)
//line /usr/local/go/src/net/http/transport.go:994
		// _ = "end of CoverTab[44312]"
	} else {
//line /usr/local/go/src/net/http/transport.go:995
		_go_fuzz_dep_.CoverTab[44313]++
//line /usr/local/go/src/net/http/transport.go:995
		// _ = "end of CoverTab[44313]"
//line /usr/local/go/src/net/http/transport.go:995
	}
//line /usr/local/go/src/net/http/transport.go:995
	// _ = "end of CoverTab[44277]"
//line /usr/local/go/src/net/http/transport.go:995
	_go_fuzz_dep_.CoverTab[44278]++

//line /usr/local/go/src/net/http/transport.go:1000
	if t.IdleConnTimeout > 0 && func() bool {
//line /usr/local/go/src/net/http/transport.go:1000
		_go_fuzz_dep_.CoverTab[44314]++
//line /usr/local/go/src/net/http/transport.go:1000
		return pconn.alt == nil
//line /usr/local/go/src/net/http/transport.go:1000
		// _ = "end of CoverTab[44314]"
//line /usr/local/go/src/net/http/transport.go:1000
	}() {
//line /usr/local/go/src/net/http/transport.go:1000
		_go_fuzz_dep_.CoverTab[44315]++
								if pconn.idleTimer != nil {
//line /usr/local/go/src/net/http/transport.go:1001
			_go_fuzz_dep_.CoverTab[44316]++
									pconn.idleTimer.Reset(t.IdleConnTimeout)
//line /usr/local/go/src/net/http/transport.go:1002
			// _ = "end of CoverTab[44316]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1003
			_go_fuzz_dep_.CoverTab[44317]++
									pconn.idleTimer = time.AfterFunc(t.IdleConnTimeout, pconn.closeConnIfStillIdle)
//line /usr/local/go/src/net/http/transport.go:1004
			// _ = "end of CoverTab[44317]"
		}
//line /usr/local/go/src/net/http/transport.go:1005
		// _ = "end of CoverTab[44315]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1006
		_go_fuzz_dep_.CoverTab[44318]++
//line /usr/local/go/src/net/http/transport.go:1006
		// _ = "end of CoverTab[44318]"
//line /usr/local/go/src/net/http/transport.go:1006
	}
//line /usr/local/go/src/net/http/transport.go:1006
	// _ = "end of CoverTab[44278]"
//line /usr/local/go/src/net/http/transport.go:1006
	_go_fuzz_dep_.CoverTab[44279]++
							pconn.idleAt = time.Now()
							return nil
//line /usr/local/go/src/net/http/transport.go:1008
	// _ = "end of CoverTab[44279]"
}

// queueForIdleConn queues w to receive the next idle connection for w.cm.
//line /usr/local/go/src/net/http/transport.go:1011
// As an optimization hint to the caller, queueForIdleConn reports whether
//line /usr/local/go/src/net/http/transport.go:1011
// it successfully delivered an already-idle connection.
//line /usr/local/go/src/net/http/transport.go:1014
func (t *Transport) queueForIdleConn(w *wantConn) (delivered bool) {
//line /usr/local/go/src/net/http/transport.go:1014
	_go_fuzz_dep_.CoverTab[44319]++
							if t.DisableKeepAlives {
//line /usr/local/go/src/net/http/transport.go:1015
		_go_fuzz_dep_.CoverTab[44325]++
								return false
//line /usr/local/go/src/net/http/transport.go:1016
		// _ = "end of CoverTab[44325]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1017
		_go_fuzz_dep_.CoverTab[44326]++
//line /usr/local/go/src/net/http/transport.go:1017
		// _ = "end of CoverTab[44326]"
//line /usr/local/go/src/net/http/transport.go:1017
	}
//line /usr/local/go/src/net/http/transport.go:1017
	// _ = "end of CoverTab[44319]"
//line /usr/local/go/src/net/http/transport.go:1017
	_go_fuzz_dep_.CoverTab[44320]++

							t.idleMu.Lock()
							defer t.idleMu.Unlock()

//line /usr/local/go/src/net/http/transport.go:1024
	t.closeIdle = false

	if w == nil {
//line /usr/local/go/src/net/http/transport.go:1026
		_go_fuzz_dep_.CoverTab[44327]++

								return false
//line /usr/local/go/src/net/http/transport.go:1028
		// _ = "end of CoverTab[44327]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1029
		_go_fuzz_dep_.CoverTab[44328]++
//line /usr/local/go/src/net/http/transport.go:1029
		// _ = "end of CoverTab[44328]"
//line /usr/local/go/src/net/http/transport.go:1029
	}
//line /usr/local/go/src/net/http/transport.go:1029
	// _ = "end of CoverTab[44320]"
//line /usr/local/go/src/net/http/transport.go:1029
	_go_fuzz_dep_.CoverTab[44321]++

	// If IdleConnTimeout is set, calculate the oldest
	// persistConn.idleAt time we're willing to use a cached idle
	// conn.
	var oldTime time.Time
	if t.IdleConnTimeout > 0 {
//line /usr/local/go/src/net/http/transport.go:1035
		_go_fuzz_dep_.CoverTab[44329]++
								oldTime = time.Now().Add(-t.IdleConnTimeout)
//line /usr/local/go/src/net/http/transport.go:1036
		// _ = "end of CoverTab[44329]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1037
		_go_fuzz_dep_.CoverTab[44330]++
//line /usr/local/go/src/net/http/transport.go:1037
		// _ = "end of CoverTab[44330]"
//line /usr/local/go/src/net/http/transport.go:1037
	}
//line /usr/local/go/src/net/http/transport.go:1037
	// _ = "end of CoverTab[44321]"
//line /usr/local/go/src/net/http/transport.go:1037
	_go_fuzz_dep_.CoverTab[44322]++

//line /usr/local/go/src/net/http/transport.go:1040
	if list, ok := t.idleConn[w.key]; ok {
//line /usr/local/go/src/net/http/transport.go:1040
		_go_fuzz_dep_.CoverTab[44331]++
								stop := false
								delivered := false
								for len(list) > 0 && func() bool {
//line /usr/local/go/src/net/http/transport.go:1043
			_go_fuzz_dep_.CoverTab[44334]++
//line /usr/local/go/src/net/http/transport.go:1043
			return !stop
//line /usr/local/go/src/net/http/transport.go:1043
			// _ = "end of CoverTab[44334]"
//line /usr/local/go/src/net/http/transport.go:1043
		}() {
//line /usr/local/go/src/net/http/transport.go:1043
			_go_fuzz_dep_.CoverTab[44335]++
									pconn := list[len(list)-1]

//line /usr/local/go/src/net/http/transport.go:1049
			tooOld := !oldTime.IsZero() && func() bool {
//line /usr/local/go/src/net/http/transport.go:1049
				_go_fuzz_dep_.CoverTab[44339]++
//line /usr/local/go/src/net/http/transport.go:1049
				return pconn.idleAt.Round(0).Before(oldTime)
//line /usr/local/go/src/net/http/transport.go:1049
				// _ = "end of CoverTab[44339]"
//line /usr/local/go/src/net/http/transport.go:1049
			}()
			if tooOld {
//line /usr/local/go/src/net/http/transport.go:1050
				_go_fuzz_dep_.CoverTab[44340]++
//line /usr/local/go/src/net/http/transport.go:1050
				_curRoutineNum37_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/transport.go:1050
				_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum37_)

//line /usr/local/go/src/net/http/transport.go:1054
				go pconn.closeConnIfStillIdle()
//line /usr/local/go/src/net/http/transport.go:1054
				// _ = "end of CoverTab[44340]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1055
				_go_fuzz_dep_.CoverTab[44341]++
//line /usr/local/go/src/net/http/transport.go:1055
				// _ = "end of CoverTab[44341]"
//line /usr/local/go/src/net/http/transport.go:1055
			}
//line /usr/local/go/src/net/http/transport.go:1055
			// _ = "end of CoverTab[44335]"
//line /usr/local/go/src/net/http/transport.go:1055
			_go_fuzz_dep_.CoverTab[44336]++
									if pconn.isBroken() || func() bool {
//line /usr/local/go/src/net/http/transport.go:1056
				_go_fuzz_dep_.CoverTab[44342]++
//line /usr/local/go/src/net/http/transport.go:1056
				return tooOld
//line /usr/local/go/src/net/http/transport.go:1056
				// _ = "end of CoverTab[44342]"
//line /usr/local/go/src/net/http/transport.go:1056
			}() {
//line /usr/local/go/src/net/http/transport.go:1056
				_go_fuzz_dep_.CoverTab[44343]++

//line /usr/local/go/src/net/http/transport.go:1062
				list = list[:len(list)-1]
										continue
//line /usr/local/go/src/net/http/transport.go:1063
				// _ = "end of CoverTab[44343]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1064
				_go_fuzz_dep_.CoverTab[44344]++
//line /usr/local/go/src/net/http/transport.go:1064
				// _ = "end of CoverTab[44344]"
//line /usr/local/go/src/net/http/transport.go:1064
			}
//line /usr/local/go/src/net/http/transport.go:1064
			// _ = "end of CoverTab[44336]"
//line /usr/local/go/src/net/http/transport.go:1064
			_go_fuzz_dep_.CoverTab[44337]++
									delivered = w.tryDeliver(pconn, nil)
									if delivered {
//line /usr/local/go/src/net/http/transport.go:1066
				_go_fuzz_dep_.CoverTab[44345]++
										if pconn.alt != nil {
//line /usr/local/go/src/net/http/transport.go:1067
					_go_fuzz_dep_.CoverTab[44346]++
//line /usr/local/go/src/net/http/transport.go:1067
					// _ = "end of CoverTab[44346]"

//line /usr/local/go/src/net/http/transport.go:1070
				} else {
//line /usr/local/go/src/net/http/transport.go:1070
					_go_fuzz_dep_.CoverTab[44347]++

//line /usr/local/go/src/net/http/transport.go:1073
					t.idleLRU.remove(pconn)
											list = list[:len(list)-1]
//line /usr/local/go/src/net/http/transport.go:1074
					// _ = "end of CoverTab[44347]"
				}
//line /usr/local/go/src/net/http/transport.go:1075
				// _ = "end of CoverTab[44345]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1076
				_go_fuzz_dep_.CoverTab[44348]++
//line /usr/local/go/src/net/http/transport.go:1076
				// _ = "end of CoverTab[44348]"
//line /usr/local/go/src/net/http/transport.go:1076
			}
//line /usr/local/go/src/net/http/transport.go:1076
			// _ = "end of CoverTab[44337]"
//line /usr/local/go/src/net/http/transport.go:1076
			_go_fuzz_dep_.CoverTab[44338]++
									stop = true
//line /usr/local/go/src/net/http/transport.go:1077
			// _ = "end of CoverTab[44338]"
		}
//line /usr/local/go/src/net/http/transport.go:1078
		// _ = "end of CoverTab[44331]"
//line /usr/local/go/src/net/http/transport.go:1078
		_go_fuzz_dep_.CoverTab[44332]++
								if len(list) > 0 {
//line /usr/local/go/src/net/http/transport.go:1079
			_go_fuzz_dep_.CoverTab[44349]++
									t.idleConn[w.key] = list
//line /usr/local/go/src/net/http/transport.go:1080
			// _ = "end of CoverTab[44349]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1081
			_go_fuzz_dep_.CoverTab[44350]++
									delete(t.idleConn, w.key)
//line /usr/local/go/src/net/http/transport.go:1082
			// _ = "end of CoverTab[44350]"
		}
//line /usr/local/go/src/net/http/transport.go:1083
		// _ = "end of CoverTab[44332]"
//line /usr/local/go/src/net/http/transport.go:1083
		_go_fuzz_dep_.CoverTab[44333]++
								if stop {
//line /usr/local/go/src/net/http/transport.go:1084
			_go_fuzz_dep_.CoverTab[44351]++
									return delivered
//line /usr/local/go/src/net/http/transport.go:1085
			// _ = "end of CoverTab[44351]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1086
			_go_fuzz_dep_.CoverTab[44352]++
//line /usr/local/go/src/net/http/transport.go:1086
			// _ = "end of CoverTab[44352]"
//line /usr/local/go/src/net/http/transport.go:1086
		}
//line /usr/local/go/src/net/http/transport.go:1086
		// _ = "end of CoverTab[44333]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1087
		_go_fuzz_dep_.CoverTab[44353]++
//line /usr/local/go/src/net/http/transport.go:1087
		// _ = "end of CoverTab[44353]"
//line /usr/local/go/src/net/http/transport.go:1087
	}
//line /usr/local/go/src/net/http/transport.go:1087
	// _ = "end of CoverTab[44322]"
//line /usr/local/go/src/net/http/transport.go:1087
	_go_fuzz_dep_.CoverTab[44323]++

//line /usr/local/go/src/net/http/transport.go:1090
	if t.idleConnWait == nil {
//line /usr/local/go/src/net/http/transport.go:1090
		_go_fuzz_dep_.CoverTab[44354]++
								t.idleConnWait = make(map[connectMethodKey]wantConnQueue)
//line /usr/local/go/src/net/http/transport.go:1091
		// _ = "end of CoverTab[44354]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1092
		_go_fuzz_dep_.CoverTab[44355]++
//line /usr/local/go/src/net/http/transport.go:1092
		// _ = "end of CoverTab[44355]"
//line /usr/local/go/src/net/http/transport.go:1092
	}
//line /usr/local/go/src/net/http/transport.go:1092
	// _ = "end of CoverTab[44323]"
//line /usr/local/go/src/net/http/transport.go:1092
	_go_fuzz_dep_.CoverTab[44324]++
							q := t.idleConnWait[w.key]
							q.cleanFront()
							q.pushBack(w)
							t.idleConnWait[w.key] = q
							return false
//line /usr/local/go/src/net/http/transport.go:1097
	// _ = "end of CoverTab[44324]"
}

// removeIdleConn marks pconn as dead.
func (t *Transport) removeIdleConn(pconn *persistConn) bool {
//line /usr/local/go/src/net/http/transport.go:1101
	_go_fuzz_dep_.CoverTab[44356]++
							t.idleMu.Lock()
							defer t.idleMu.Unlock()
							return t.removeIdleConnLocked(pconn)
//line /usr/local/go/src/net/http/transport.go:1104
	// _ = "end of CoverTab[44356]"
}

// t.idleMu must be held.
func (t *Transport) removeIdleConnLocked(pconn *persistConn) bool {
//line /usr/local/go/src/net/http/transport.go:1108
	_go_fuzz_dep_.CoverTab[44357]++
							if pconn.idleTimer != nil {
//line /usr/local/go/src/net/http/transport.go:1109
		_go_fuzz_dep_.CoverTab[44360]++
								pconn.idleTimer.Stop()
//line /usr/local/go/src/net/http/transport.go:1110
		// _ = "end of CoverTab[44360]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1111
		_go_fuzz_dep_.CoverTab[44361]++
//line /usr/local/go/src/net/http/transport.go:1111
		// _ = "end of CoverTab[44361]"
//line /usr/local/go/src/net/http/transport.go:1111
	}
//line /usr/local/go/src/net/http/transport.go:1111
	// _ = "end of CoverTab[44357]"
//line /usr/local/go/src/net/http/transport.go:1111
	_go_fuzz_dep_.CoverTab[44358]++
							t.idleLRU.remove(pconn)
							key := pconn.cacheKey
							pconns := t.idleConn[key]
							var removed bool
							switch len(pconns) {
	case 0:
//line /usr/local/go/src/net/http/transport.go:1117
		_go_fuzz_dep_.CoverTab[44362]++
//line /usr/local/go/src/net/http/transport.go:1117
		// _ = "end of CoverTab[44362]"

	case 1:
//line /usr/local/go/src/net/http/transport.go:1119
		_go_fuzz_dep_.CoverTab[44363]++
								if pconns[0] == pconn {
//line /usr/local/go/src/net/http/transport.go:1120
			_go_fuzz_dep_.CoverTab[44365]++
									delete(t.idleConn, key)
									removed = true
//line /usr/local/go/src/net/http/transport.go:1122
			// _ = "end of CoverTab[44365]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1123
			_go_fuzz_dep_.CoverTab[44366]++
//line /usr/local/go/src/net/http/transport.go:1123
			// _ = "end of CoverTab[44366]"
//line /usr/local/go/src/net/http/transport.go:1123
		}
//line /usr/local/go/src/net/http/transport.go:1123
		// _ = "end of CoverTab[44363]"
	default:
//line /usr/local/go/src/net/http/transport.go:1124
		_go_fuzz_dep_.CoverTab[44364]++
								for i, v := range pconns {
//line /usr/local/go/src/net/http/transport.go:1125
			_go_fuzz_dep_.CoverTab[44367]++
									if v != pconn {
//line /usr/local/go/src/net/http/transport.go:1126
				_go_fuzz_dep_.CoverTab[44369]++
										continue
//line /usr/local/go/src/net/http/transport.go:1127
				// _ = "end of CoverTab[44369]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1128
				_go_fuzz_dep_.CoverTab[44370]++
//line /usr/local/go/src/net/http/transport.go:1128
				// _ = "end of CoverTab[44370]"
//line /usr/local/go/src/net/http/transport.go:1128
			}
//line /usr/local/go/src/net/http/transport.go:1128
			// _ = "end of CoverTab[44367]"
//line /usr/local/go/src/net/http/transport.go:1128
			_go_fuzz_dep_.CoverTab[44368]++

//line /usr/local/go/src/net/http/transport.go:1131
			copy(pconns[i:], pconns[i+1:])
									t.idleConn[key] = pconns[:len(pconns)-1]
									removed = true
									break
//line /usr/local/go/src/net/http/transport.go:1134
			// _ = "end of CoverTab[44368]"
		}
//line /usr/local/go/src/net/http/transport.go:1135
		// _ = "end of CoverTab[44364]"
	}
//line /usr/local/go/src/net/http/transport.go:1136
	// _ = "end of CoverTab[44358]"
//line /usr/local/go/src/net/http/transport.go:1136
	_go_fuzz_dep_.CoverTab[44359]++
							return removed
//line /usr/local/go/src/net/http/transport.go:1137
	// _ = "end of CoverTab[44359]"
}

func (t *Transport) setReqCanceler(key cancelKey, fn func(error)) {
//line /usr/local/go/src/net/http/transport.go:1140
	_go_fuzz_dep_.CoverTab[44371]++
							t.reqMu.Lock()
							defer t.reqMu.Unlock()
							if t.reqCanceler == nil {
//line /usr/local/go/src/net/http/transport.go:1143
		_go_fuzz_dep_.CoverTab[44373]++
								t.reqCanceler = make(map[cancelKey]func(error))
//line /usr/local/go/src/net/http/transport.go:1144
		// _ = "end of CoverTab[44373]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1145
		_go_fuzz_dep_.CoverTab[44374]++
//line /usr/local/go/src/net/http/transport.go:1145
		// _ = "end of CoverTab[44374]"
//line /usr/local/go/src/net/http/transport.go:1145
	}
//line /usr/local/go/src/net/http/transport.go:1145
	// _ = "end of CoverTab[44371]"
//line /usr/local/go/src/net/http/transport.go:1145
	_go_fuzz_dep_.CoverTab[44372]++
							if fn != nil {
//line /usr/local/go/src/net/http/transport.go:1146
		_go_fuzz_dep_.CoverTab[44375]++
								t.reqCanceler[key] = fn
//line /usr/local/go/src/net/http/transport.go:1147
		// _ = "end of CoverTab[44375]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1148
		_go_fuzz_dep_.CoverTab[44376]++
								delete(t.reqCanceler, key)
//line /usr/local/go/src/net/http/transport.go:1149
		// _ = "end of CoverTab[44376]"
	}
//line /usr/local/go/src/net/http/transport.go:1150
	// _ = "end of CoverTab[44372]"
}

// replaceReqCanceler replaces an existing cancel function. If there is no cancel function
//line /usr/local/go/src/net/http/transport.go:1153
// for the request, we don't set the function and return false.
//line /usr/local/go/src/net/http/transport.go:1153
// Since CancelRequest will clear the canceler, we can use the return value to detect if
//line /usr/local/go/src/net/http/transport.go:1153
// the request was canceled since the last setReqCancel call.
//line /usr/local/go/src/net/http/transport.go:1157
func (t *Transport) replaceReqCanceler(key cancelKey, fn func(error)) bool {
//line /usr/local/go/src/net/http/transport.go:1157
	_go_fuzz_dep_.CoverTab[44377]++
							t.reqMu.Lock()
							defer t.reqMu.Unlock()
							_, ok := t.reqCanceler[key]
							if !ok {
//line /usr/local/go/src/net/http/transport.go:1161
		_go_fuzz_dep_.CoverTab[44380]++
								return false
//line /usr/local/go/src/net/http/transport.go:1162
		// _ = "end of CoverTab[44380]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1163
		_go_fuzz_dep_.CoverTab[44381]++
//line /usr/local/go/src/net/http/transport.go:1163
		// _ = "end of CoverTab[44381]"
//line /usr/local/go/src/net/http/transport.go:1163
	}
//line /usr/local/go/src/net/http/transport.go:1163
	// _ = "end of CoverTab[44377]"
//line /usr/local/go/src/net/http/transport.go:1163
	_go_fuzz_dep_.CoverTab[44378]++
							if fn != nil {
//line /usr/local/go/src/net/http/transport.go:1164
		_go_fuzz_dep_.CoverTab[44382]++
								t.reqCanceler[key] = fn
//line /usr/local/go/src/net/http/transport.go:1165
		// _ = "end of CoverTab[44382]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1166
		_go_fuzz_dep_.CoverTab[44383]++
								delete(t.reqCanceler, key)
//line /usr/local/go/src/net/http/transport.go:1167
		// _ = "end of CoverTab[44383]"
	}
//line /usr/local/go/src/net/http/transport.go:1168
	// _ = "end of CoverTab[44378]"
//line /usr/local/go/src/net/http/transport.go:1168
	_go_fuzz_dep_.CoverTab[44379]++
							return true
//line /usr/local/go/src/net/http/transport.go:1169
	// _ = "end of CoverTab[44379]"
}

var zeroDialer net.Dialer

func (t *Transport) dial(ctx context.Context, network, addr string) (net.Conn, error) {
//line /usr/local/go/src/net/http/transport.go:1174
	_go_fuzz_dep_.CoverTab[44384]++
							if t.DialContext != nil {
//line /usr/local/go/src/net/http/transport.go:1175
		_go_fuzz_dep_.CoverTab[44387]++
								return t.DialContext(ctx, network, addr)
//line /usr/local/go/src/net/http/transport.go:1176
		// _ = "end of CoverTab[44387]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1177
		_go_fuzz_dep_.CoverTab[44388]++
//line /usr/local/go/src/net/http/transport.go:1177
		// _ = "end of CoverTab[44388]"
//line /usr/local/go/src/net/http/transport.go:1177
	}
//line /usr/local/go/src/net/http/transport.go:1177
	// _ = "end of CoverTab[44384]"
//line /usr/local/go/src/net/http/transport.go:1177
	_go_fuzz_dep_.CoverTab[44385]++
							if t.Dial != nil {
//line /usr/local/go/src/net/http/transport.go:1178
		_go_fuzz_dep_.CoverTab[44389]++
								c, err := t.Dial(network, addr)
								if c == nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1180
			_go_fuzz_dep_.CoverTab[44391]++
//line /usr/local/go/src/net/http/transport.go:1180
			return err == nil
//line /usr/local/go/src/net/http/transport.go:1180
			// _ = "end of CoverTab[44391]"
//line /usr/local/go/src/net/http/transport.go:1180
		}() {
//line /usr/local/go/src/net/http/transport.go:1180
			_go_fuzz_dep_.CoverTab[44392]++
									err = errors.New("net/http: Transport.Dial hook returned (nil, nil)")
//line /usr/local/go/src/net/http/transport.go:1181
			// _ = "end of CoverTab[44392]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1182
			_go_fuzz_dep_.CoverTab[44393]++
//line /usr/local/go/src/net/http/transport.go:1182
			// _ = "end of CoverTab[44393]"
//line /usr/local/go/src/net/http/transport.go:1182
		}
//line /usr/local/go/src/net/http/transport.go:1182
		// _ = "end of CoverTab[44389]"
//line /usr/local/go/src/net/http/transport.go:1182
		_go_fuzz_dep_.CoverTab[44390]++
								return c, err
//line /usr/local/go/src/net/http/transport.go:1183
		// _ = "end of CoverTab[44390]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1184
		_go_fuzz_dep_.CoverTab[44394]++
//line /usr/local/go/src/net/http/transport.go:1184
		// _ = "end of CoverTab[44394]"
//line /usr/local/go/src/net/http/transport.go:1184
	}
//line /usr/local/go/src/net/http/transport.go:1184
	// _ = "end of CoverTab[44385]"
//line /usr/local/go/src/net/http/transport.go:1184
	_go_fuzz_dep_.CoverTab[44386]++
							return zeroDialer.DialContext(ctx, network, addr)
//line /usr/local/go/src/net/http/transport.go:1185
	// _ = "end of CoverTab[44386]"
}

// A wantConn records state about a wanted connection
//line /usr/local/go/src/net/http/transport.go:1188
// (that is, an active call to getConn).
//line /usr/local/go/src/net/http/transport.go:1188
// The conn may be gotten by dialing or by finding an idle connection,
//line /usr/local/go/src/net/http/transport.go:1188
// or a cancellation may make the conn no longer wanted.
//line /usr/local/go/src/net/http/transport.go:1188
// These three options are racing against each other and use
//line /usr/local/go/src/net/http/transport.go:1188
// wantConn to coordinate and agree about the winning outcome.
//line /usr/local/go/src/net/http/transport.go:1194
type wantConn struct {
	cm	connectMethod
	key	connectMethodKey	// cm.key()
	ctx	context.Context		// context for dial
	ready	chan struct{}		// closed when pc, err pair is delivered

	// hooks for testing to know when dials are done
	// beforeDial is called in the getConn goroutine when the dial is queued.
	// afterDial is called when the dial is completed or canceled.
	beforeDial	func()
	afterDial	func()

	mu	sync.Mutex	// protects pc, err, close(ready)
	pc	*persistConn
	err	error
}

// waiting reports whether w is still waiting for an answer (connection or error).
func (w *wantConn) waiting() bool {
//line /usr/local/go/src/net/http/transport.go:1212
	_go_fuzz_dep_.CoverTab[44395]++
							select {
	case <-w.ready:
//line /usr/local/go/src/net/http/transport.go:1214
		_go_fuzz_dep_.CoverTab[44396]++
								return false
//line /usr/local/go/src/net/http/transport.go:1215
		// _ = "end of CoverTab[44396]"
	default:
//line /usr/local/go/src/net/http/transport.go:1216
		_go_fuzz_dep_.CoverTab[44397]++
								return true
//line /usr/local/go/src/net/http/transport.go:1217
		// _ = "end of CoverTab[44397]"
	}
//line /usr/local/go/src/net/http/transport.go:1218
	// _ = "end of CoverTab[44395]"
}

// tryDeliver attempts to deliver pc, err to w and reports whether it succeeded.
func (w *wantConn) tryDeliver(pc *persistConn, err error) bool {
//line /usr/local/go/src/net/http/transport.go:1222
	_go_fuzz_dep_.CoverTab[44398]++
							w.mu.Lock()
							defer w.mu.Unlock()

							if w.pc != nil || func() bool {
//line /usr/local/go/src/net/http/transport.go:1226
		_go_fuzz_dep_.CoverTab[44401]++
//line /usr/local/go/src/net/http/transport.go:1226
		return w.err != nil
//line /usr/local/go/src/net/http/transport.go:1226
		// _ = "end of CoverTab[44401]"
//line /usr/local/go/src/net/http/transport.go:1226
	}() {
//line /usr/local/go/src/net/http/transport.go:1226
		_go_fuzz_dep_.CoverTab[44402]++
								return false
//line /usr/local/go/src/net/http/transport.go:1227
		// _ = "end of CoverTab[44402]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1228
		_go_fuzz_dep_.CoverTab[44403]++
//line /usr/local/go/src/net/http/transport.go:1228
		// _ = "end of CoverTab[44403]"
//line /usr/local/go/src/net/http/transport.go:1228
	}
//line /usr/local/go/src/net/http/transport.go:1228
	// _ = "end of CoverTab[44398]"
//line /usr/local/go/src/net/http/transport.go:1228
	_go_fuzz_dep_.CoverTab[44399]++

							w.pc = pc
							w.err = err
							if w.pc == nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1232
		_go_fuzz_dep_.CoverTab[44404]++
//line /usr/local/go/src/net/http/transport.go:1232
		return w.err == nil
//line /usr/local/go/src/net/http/transport.go:1232
		// _ = "end of CoverTab[44404]"
//line /usr/local/go/src/net/http/transport.go:1232
	}() {
//line /usr/local/go/src/net/http/transport.go:1232
		_go_fuzz_dep_.CoverTab[44405]++
								panic("net/http: internal error: misuse of tryDeliver")
//line /usr/local/go/src/net/http/transport.go:1233
		// _ = "end of CoverTab[44405]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1234
		_go_fuzz_dep_.CoverTab[44406]++
//line /usr/local/go/src/net/http/transport.go:1234
		// _ = "end of CoverTab[44406]"
//line /usr/local/go/src/net/http/transport.go:1234
	}
//line /usr/local/go/src/net/http/transport.go:1234
	// _ = "end of CoverTab[44399]"
//line /usr/local/go/src/net/http/transport.go:1234
	_go_fuzz_dep_.CoverTab[44400]++
							close(w.ready)
							return true
//line /usr/local/go/src/net/http/transport.go:1236
	// _ = "end of CoverTab[44400]"
}

// cancel marks w as no longer wanting a result (for example, due to cancellation).
//line /usr/local/go/src/net/http/transport.go:1239
// If a connection has been delivered already, cancel returns it with t.putOrCloseIdleConn.
//line /usr/local/go/src/net/http/transport.go:1241
func (w *wantConn) cancel(t *Transport, err error) {
//line /usr/local/go/src/net/http/transport.go:1241
	_go_fuzz_dep_.CoverTab[44407]++
							w.mu.Lock()
							if w.pc == nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1243
		_go_fuzz_dep_.CoverTab[44409]++
//line /usr/local/go/src/net/http/transport.go:1243
		return w.err == nil
//line /usr/local/go/src/net/http/transport.go:1243
		// _ = "end of CoverTab[44409]"
//line /usr/local/go/src/net/http/transport.go:1243
	}() {
//line /usr/local/go/src/net/http/transport.go:1243
		_go_fuzz_dep_.CoverTab[44410]++
								close(w.ready)
//line /usr/local/go/src/net/http/transport.go:1244
		// _ = "end of CoverTab[44410]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1245
		_go_fuzz_dep_.CoverTab[44411]++
//line /usr/local/go/src/net/http/transport.go:1245
		// _ = "end of CoverTab[44411]"
//line /usr/local/go/src/net/http/transport.go:1245
	}
//line /usr/local/go/src/net/http/transport.go:1245
	// _ = "end of CoverTab[44407]"
//line /usr/local/go/src/net/http/transport.go:1245
	_go_fuzz_dep_.CoverTab[44408]++
							pc := w.pc
							w.pc = nil
							w.err = err
							w.mu.Unlock()

							if pc != nil {
//line /usr/local/go/src/net/http/transport.go:1251
		_go_fuzz_dep_.CoverTab[44412]++
								t.putOrCloseIdleConn(pc)
//line /usr/local/go/src/net/http/transport.go:1252
		// _ = "end of CoverTab[44412]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1253
		_go_fuzz_dep_.CoverTab[44413]++
//line /usr/local/go/src/net/http/transport.go:1253
		// _ = "end of CoverTab[44413]"
//line /usr/local/go/src/net/http/transport.go:1253
	}
//line /usr/local/go/src/net/http/transport.go:1253
	// _ = "end of CoverTab[44408]"
}

// A wantConnQueue is a queue of wantConns.
type wantConnQueue struct {
	// This is a queue, not a deque.
	// It is split into two stages - head[headPos:] and tail.
	// popFront is trivial (headPos++) on the first stage, and
	// pushBack is trivial (append) on the second stage.
	// If the first stage is empty, popFront can swap the
	// first and second stages to remedy the situation.
	//
	// This two-stage split is analogous to the use of two lists
	// in Okasaki's purely functional queue but without the
	// overhead of reversing the list when swapping stages.
	head	[]*wantConn
	headPos	int
	tail	[]*wantConn
}

// len returns the number of items in the queue.
func (q *wantConnQueue) len() int {
//line /usr/local/go/src/net/http/transport.go:1274
	_go_fuzz_dep_.CoverTab[44414]++
							return len(q.head) - q.headPos + len(q.tail)
//line /usr/local/go/src/net/http/transport.go:1275
	// _ = "end of CoverTab[44414]"
}

// pushBack adds w to the back of the queue.
func (q *wantConnQueue) pushBack(w *wantConn) {
//line /usr/local/go/src/net/http/transport.go:1279
	_go_fuzz_dep_.CoverTab[44415]++
							q.tail = append(q.tail, w)
//line /usr/local/go/src/net/http/transport.go:1280
	// _ = "end of CoverTab[44415]"
}

// popFront removes and returns the wantConn at the front of the queue.
func (q *wantConnQueue) popFront() *wantConn {
//line /usr/local/go/src/net/http/transport.go:1284
	_go_fuzz_dep_.CoverTab[44416]++
							if q.headPos >= len(q.head) {
//line /usr/local/go/src/net/http/transport.go:1285
		_go_fuzz_dep_.CoverTab[44418]++
								if len(q.tail) == 0 {
//line /usr/local/go/src/net/http/transport.go:1286
			_go_fuzz_dep_.CoverTab[44420]++
									return nil
//line /usr/local/go/src/net/http/transport.go:1287
			// _ = "end of CoverTab[44420]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1288
			_go_fuzz_dep_.CoverTab[44421]++
//line /usr/local/go/src/net/http/transport.go:1288
			// _ = "end of CoverTab[44421]"
//line /usr/local/go/src/net/http/transport.go:1288
		}
//line /usr/local/go/src/net/http/transport.go:1288
		// _ = "end of CoverTab[44418]"
//line /usr/local/go/src/net/http/transport.go:1288
		_go_fuzz_dep_.CoverTab[44419]++

								q.head, q.headPos, q.tail = q.tail, 0, q.head[:0]
//line /usr/local/go/src/net/http/transport.go:1290
		// _ = "end of CoverTab[44419]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1291
		_go_fuzz_dep_.CoverTab[44422]++
//line /usr/local/go/src/net/http/transport.go:1291
		// _ = "end of CoverTab[44422]"
//line /usr/local/go/src/net/http/transport.go:1291
	}
//line /usr/local/go/src/net/http/transport.go:1291
	// _ = "end of CoverTab[44416]"
//line /usr/local/go/src/net/http/transport.go:1291
	_go_fuzz_dep_.CoverTab[44417]++
							w := q.head[q.headPos]
							q.head[q.headPos] = nil
							q.headPos++
							return w
//line /usr/local/go/src/net/http/transport.go:1295
	// _ = "end of CoverTab[44417]"
}

// peekFront returns the wantConn at the front of the queue without removing it.
func (q *wantConnQueue) peekFront() *wantConn {
//line /usr/local/go/src/net/http/transport.go:1299
	_go_fuzz_dep_.CoverTab[44423]++
							if q.headPos < len(q.head) {
//line /usr/local/go/src/net/http/transport.go:1300
		_go_fuzz_dep_.CoverTab[44426]++
								return q.head[q.headPos]
//line /usr/local/go/src/net/http/transport.go:1301
		// _ = "end of CoverTab[44426]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1302
		_go_fuzz_dep_.CoverTab[44427]++
//line /usr/local/go/src/net/http/transport.go:1302
		// _ = "end of CoverTab[44427]"
//line /usr/local/go/src/net/http/transport.go:1302
	}
//line /usr/local/go/src/net/http/transport.go:1302
	// _ = "end of CoverTab[44423]"
//line /usr/local/go/src/net/http/transport.go:1302
	_go_fuzz_dep_.CoverTab[44424]++
							if len(q.tail) > 0 {
//line /usr/local/go/src/net/http/transport.go:1303
		_go_fuzz_dep_.CoverTab[44428]++
								return q.tail[0]
//line /usr/local/go/src/net/http/transport.go:1304
		// _ = "end of CoverTab[44428]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1305
		_go_fuzz_dep_.CoverTab[44429]++
//line /usr/local/go/src/net/http/transport.go:1305
		// _ = "end of CoverTab[44429]"
//line /usr/local/go/src/net/http/transport.go:1305
	}
//line /usr/local/go/src/net/http/transport.go:1305
	// _ = "end of CoverTab[44424]"
//line /usr/local/go/src/net/http/transport.go:1305
	_go_fuzz_dep_.CoverTab[44425]++
							return nil
//line /usr/local/go/src/net/http/transport.go:1306
	// _ = "end of CoverTab[44425]"
}

// cleanFront pops any wantConns that are no longer waiting from the head of the
//line /usr/local/go/src/net/http/transport.go:1309
// queue, reporting whether any were popped.
//line /usr/local/go/src/net/http/transport.go:1311
func (q *wantConnQueue) cleanFront() (cleaned bool) {
//line /usr/local/go/src/net/http/transport.go:1311
	_go_fuzz_dep_.CoverTab[44430]++
							for {
//line /usr/local/go/src/net/http/transport.go:1312
		_go_fuzz_dep_.CoverTab[44431]++
								w := q.peekFront()
								if w == nil || func() bool {
//line /usr/local/go/src/net/http/transport.go:1314
			_go_fuzz_dep_.CoverTab[44433]++
//line /usr/local/go/src/net/http/transport.go:1314
			return w.waiting()
//line /usr/local/go/src/net/http/transport.go:1314
			// _ = "end of CoverTab[44433]"
//line /usr/local/go/src/net/http/transport.go:1314
		}() {
//line /usr/local/go/src/net/http/transport.go:1314
			_go_fuzz_dep_.CoverTab[44434]++
									return cleaned
//line /usr/local/go/src/net/http/transport.go:1315
			// _ = "end of CoverTab[44434]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1316
			_go_fuzz_dep_.CoverTab[44435]++
//line /usr/local/go/src/net/http/transport.go:1316
			// _ = "end of CoverTab[44435]"
//line /usr/local/go/src/net/http/transport.go:1316
		}
//line /usr/local/go/src/net/http/transport.go:1316
		// _ = "end of CoverTab[44431]"
//line /usr/local/go/src/net/http/transport.go:1316
		_go_fuzz_dep_.CoverTab[44432]++
								q.popFront()
								cleaned = true
//line /usr/local/go/src/net/http/transport.go:1318
		// _ = "end of CoverTab[44432]"
	}
//line /usr/local/go/src/net/http/transport.go:1319
	// _ = "end of CoverTab[44430]"
}

func (t *Transport) customDialTLS(ctx context.Context, network, addr string) (conn net.Conn, err error) {
//line /usr/local/go/src/net/http/transport.go:1322
	_go_fuzz_dep_.CoverTab[44436]++
							if t.DialTLSContext != nil {
//line /usr/local/go/src/net/http/transport.go:1323
		_go_fuzz_dep_.CoverTab[44439]++
								conn, err = t.DialTLSContext(ctx, network, addr)
//line /usr/local/go/src/net/http/transport.go:1324
		// _ = "end of CoverTab[44439]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1325
		_go_fuzz_dep_.CoverTab[44440]++
								conn, err = t.DialTLS(network, addr)
//line /usr/local/go/src/net/http/transport.go:1326
		// _ = "end of CoverTab[44440]"
	}
//line /usr/local/go/src/net/http/transport.go:1327
	// _ = "end of CoverTab[44436]"
//line /usr/local/go/src/net/http/transport.go:1327
	_go_fuzz_dep_.CoverTab[44437]++
							if conn == nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1328
		_go_fuzz_dep_.CoverTab[44441]++
//line /usr/local/go/src/net/http/transport.go:1328
		return err == nil
//line /usr/local/go/src/net/http/transport.go:1328
		// _ = "end of CoverTab[44441]"
//line /usr/local/go/src/net/http/transport.go:1328
	}() {
//line /usr/local/go/src/net/http/transport.go:1328
		_go_fuzz_dep_.CoverTab[44442]++
								err = errors.New("net/http: Transport.DialTLS or DialTLSContext returned (nil, nil)")
//line /usr/local/go/src/net/http/transport.go:1329
		// _ = "end of CoverTab[44442]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1330
		_go_fuzz_dep_.CoverTab[44443]++
//line /usr/local/go/src/net/http/transport.go:1330
		// _ = "end of CoverTab[44443]"
//line /usr/local/go/src/net/http/transport.go:1330
	}
//line /usr/local/go/src/net/http/transport.go:1330
	// _ = "end of CoverTab[44437]"
//line /usr/local/go/src/net/http/transport.go:1330
	_go_fuzz_dep_.CoverTab[44438]++
							return
//line /usr/local/go/src/net/http/transport.go:1331
	// _ = "end of CoverTab[44438]"
}

// getConn dials and creates a new persistConn to the target as
//line /usr/local/go/src/net/http/transport.go:1334
// specified in the connectMethod. This includes doing a proxy CONNECT
//line /usr/local/go/src/net/http/transport.go:1334
// and/or setting up TLS.  If this doesn't return an error, the persistConn
//line /usr/local/go/src/net/http/transport.go:1334
// is ready to write requests to.
//line /usr/local/go/src/net/http/transport.go:1338
func (t *Transport) getConn(treq *transportRequest, cm connectMethod) (pc *persistConn, err error) {
//line /usr/local/go/src/net/http/transport.go:1338
	_go_fuzz_dep_.CoverTab[44444]++
							req := treq.Request
							trace := treq.trace
							ctx := req.Context()
							if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1342
		_go_fuzz_dep_.CoverTab[44449]++
//line /usr/local/go/src/net/http/transport.go:1342
		return trace.GetConn != nil
//line /usr/local/go/src/net/http/transport.go:1342
		// _ = "end of CoverTab[44449]"
//line /usr/local/go/src/net/http/transport.go:1342
	}() {
//line /usr/local/go/src/net/http/transport.go:1342
		_go_fuzz_dep_.CoverTab[44450]++
								trace.GetConn(cm.addr())
//line /usr/local/go/src/net/http/transport.go:1343
		// _ = "end of CoverTab[44450]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1344
		_go_fuzz_dep_.CoverTab[44451]++
//line /usr/local/go/src/net/http/transport.go:1344
		// _ = "end of CoverTab[44451]"
//line /usr/local/go/src/net/http/transport.go:1344
	}
//line /usr/local/go/src/net/http/transport.go:1344
	// _ = "end of CoverTab[44444]"
//line /usr/local/go/src/net/http/transport.go:1344
	_go_fuzz_dep_.CoverTab[44445]++

							w := &wantConn{
		cm:		cm,
		key:		cm.key(),
		ctx:		ctx,
		ready:		make(chan struct{}, 1),
		beforeDial:	testHookPrePendingDial,
		afterDial:	testHookPostPendingDial,
	}
	defer func() {
//line /usr/local/go/src/net/http/transport.go:1354
		_go_fuzz_dep_.CoverTab[44452]++
								if err != nil {
//line /usr/local/go/src/net/http/transport.go:1355
			_go_fuzz_dep_.CoverTab[44453]++
									w.cancel(t, err)
//line /usr/local/go/src/net/http/transport.go:1356
			// _ = "end of CoverTab[44453]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1357
			_go_fuzz_dep_.CoverTab[44454]++
//line /usr/local/go/src/net/http/transport.go:1357
			// _ = "end of CoverTab[44454]"
//line /usr/local/go/src/net/http/transport.go:1357
		}
//line /usr/local/go/src/net/http/transport.go:1357
		// _ = "end of CoverTab[44452]"
	}()
//line /usr/local/go/src/net/http/transport.go:1358
	// _ = "end of CoverTab[44445]"
//line /usr/local/go/src/net/http/transport.go:1358
	_go_fuzz_dep_.CoverTab[44446]++

//line /usr/local/go/src/net/http/transport.go:1361
	if delivered := t.queueForIdleConn(w); delivered {
//line /usr/local/go/src/net/http/transport.go:1361
		_go_fuzz_dep_.CoverTab[44455]++
								pc := w.pc

//line /usr/local/go/src/net/http/transport.go:1365
		if pc.alt == nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1365
			_go_fuzz_dep_.CoverTab[44458]++
//line /usr/local/go/src/net/http/transport.go:1365
			return trace != nil
//line /usr/local/go/src/net/http/transport.go:1365
			// _ = "end of CoverTab[44458]"
//line /usr/local/go/src/net/http/transport.go:1365
		}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:1365
			_go_fuzz_dep_.CoverTab[44459]++
//line /usr/local/go/src/net/http/transport.go:1365
			return trace.GotConn != nil
//line /usr/local/go/src/net/http/transport.go:1365
			// _ = "end of CoverTab[44459]"
//line /usr/local/go/src/net/http/transport.go:1365
		}() {
//line /usr/local/go/src/net/http/transport.go:1365
			_go_fuzz_dep_.CoverTab[44460]++
									trace.GotConn(pc.gotIdleConnTrace(pc.idleAt))
//line /usr/local/go/src/net/http/transport.go:1366
			// _ = "end of CoverTab[44460]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1367
			_go_fuzz_dep_.CoverTab[44461]++
//line /usr/local/go/src/net/http/transport.go:1367
			// _ = "end of CoverTab[44461]"
//line /usr/local/go/src/net/http/transport.go:1367
		}
//line /usr/local/go/src/net/http/transport.go:1367
		// _ = "end of CoverTab[44455]"
//line /usr/local/go/src/net/http/transport.go:1367
		_go_fuzz_dep_.CoverTab[44456]++

//line /usr/local/go/src/net/http/transport.go:1371
		t.setReqCanceler(treq.cancelKey, func(error) { _go_fuzz_dep_.CoverTab[44462]++; // _ = "end of CoverTab[44462]" })
//line /usr/local/go/src/net/http/transport.go:1371
		// _ = "end of CoverTab[44456]"
//line /usr/local/go/src/net/http/transport.go:1371
		_go_fuzz_dep_.CoverTab[44457]++
								return pc, nil
//line /usr/local/go/src/net/http/transport.go:1372
		// _ = "end of CoverTab[44457]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1373
		_go_fuzz_dep_.CoverTab[44463]++
//line /usr/local/go/src/net/http/transport.go:1373
		// _ = "end of CoverTab[44463]"
//line /usr/local/go/src/net/http/transport.go:1373
	}
//line /usr/local/go/src/net/http/transport.go:1373
	// _ = "end of CoverTab[44446]"
//line /usr/local/go/src/net/http/transport.go:1373
	_go_fuzz_dep_.CoverTab[44447]++

							cancelc := make(chan error, 1)
							t.setReqCanceler(treq.cancelKey, func(err error) { _go_fuzz_dep_.CoverTab[44464]++; cancelc <- err; // _ = "end of CoverTab[44464]" })
//line /usr/local/go/src/net/http/transport.go:1376
	// _ = "end of CoverTab[44447]"
//line /usr/local/go/src/net/http/transport.go:1376
	_go_fuzz_dep_.CoverTab[44448]++

//line /usr/local/go/src/net/http/transport.go:1379
	t.queueForDial(w)

//line /usr/local/go/src/net/http/transport.go:1382
	select {
	case <-w.ready:
//line /usr/local/go/src/net/http/transport.go:1383
		_go_fuzz_dep_.CoverTab[44465]++

//line /usr/local/go/src/net/http/transport.go:1386
		if w.pc != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1386
			_go_fuzz_dep_.CoverTab[44472]++
//line /usr/local/go/src/net/http/transport.go:1386
			return w.pc.alt == nil
//line /usr/local/go/src/net/http/transport.go:1386
			// _ = "end of CoverTab[44472]"
//line /usr/local/go/src/net/http/transport.go:1386
		}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:1386
			_go_fuzz_dep_.CoverTab[44473]++
//line /usr/local/go/src/net/http/transport.go:1386
			return trace != nil
//line /usr/local/go/src/net/http/transport.go:1386
			// _ = "end of CoverTab[44473]"
//line /usr/local/go/src/net/http/transport.go:1386
		}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:1386
			_go_fuzz_dep_.CoverTab[44474]++
//line /usr/local/go/src/net/http/transport.go:1386
			return trace.GotConn != nil
//line /usr/local/go/src/net/http/transport.go:1386
			// _ = "end of CoverTab[44474]"
//line /usr/local/go/src/net/http/transport.go:1386
		}() {
//line /usr/local/go/src/net/http/transport.go:1386
			_go_fuzz_dep_.CoverTab[44475]++
									trace.GotConn(httptrace.GotConnInfo{Conn: w.pc.conn, Reused: w.pc.isReused()})
//line /usr/local/go/src/net/http/transport.go:1387
			// _ = "end of CoverTab[44475]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1388
			_go_fuzz_dep_.CoverTab[44476]++
//line /usr/local/go/src/net/http/transport.go:1388
			// _ = "end of CoverTab[44476]"
//line /usr/local/go/src/net/http/transport.go:1388
		}
//line /usr/local/go/src/net/http/transport.go:1388
		// _ = "end of CoverTab[44465]"
//line /usr/local/go/src/net/http/transport.go:1388
		_go_fuzz_dep_.CoverTab[44466]++
								if w.err != nil {
//line /usr/local/go/src/net/http/transport.go:1389
			_go_fuzz_dep_.CoverTab[44477]++

//line /usr/local/go/src/net/http/transport.go:1393
			select {
			case <-req.Cancel:
//line /usr/local/go/src/net/http/transport.go:1394
				_go_fuzz_dep_.CoverTab[44478]++
										return nil, errRequestCanceledConn
//line /usr/local/go/src/net/http/transport.go:1395
				// _ = "end of CoverTab[44478]"
			case <-req.Context().Done():
//line /usr/local/go/src/net/http/transport.go:1396
				_go_fuzz_dep_.CoverTab[44479]++
										return nil, req.Context().Err()
//line /usr/local/go/src/net/http/transport.go:1397
				// _ = "end of CoverTab[44479]"
			case err := <-cancelc:
//line /usr/local/go/src/net/http/transport.go:1398
				_go_fuzz_dep_.CoverTab[44480]++
										if err == errRequestCanceled {
//line /usr/local/go/src/net/http/transport.go:1399
					_go_fuzz_dep_.CoverTab[44483]++
											err = errRequestCanceledConn
//line /usr/local/go/src/net/http/transport.go:1400
					// _ = "end of CoverTab[44483]"
				} else {
//line /usr/local/go/src/net/http/transport.go:1401
					_go_fuzz_dep_.CoverTab[44484]++
//line /usr/local/go/src/net/http/transport.go:1401
					// _ = "end of CoverTab[44484]"
//line /usr/local/go/src/net/http/transport.go:1401
				}
//line /usr/local/go/src/net/http/transport.go:1401
				// _ = "end of CoverTab[44480]"
//line /usr/local/go/src/net/http/transport.go:1401
				_go_fuzz_dep_.CoverTab[44481]++
										return nil, err
//line /usr/local/go/src/net/http/transport.go:1402
				// _ = "end of CoverTab[44481]"
			default:
//line /usr/local/go/src/net/http/transport.go:1403
				_go_fuzz_dep_.CoverTab[44482]++
//line /usr/local/go/src/net/http/transport.go:1403
				// _ = "end of CoverTab[44482]"

			}
//line /usr/local/go/src/net/http/transport.go:1405
			// _ = "end of CoverTab[44477]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1406
			_go_fuzz_dep_.CoverTab[44485]++
//line /usr/local/go/src/net/http/transport.go:1406
			// _ = "end of CoverTab[44485]"
//line /usr/local/go/src/net/http/transport.go:1406
		}
//line /usr/local/go/src/net/http/transport.go:1406
		// _ = "end of CoverTab[44466]"
//line /usr/local/go/src/net/http/transport.go:1406
		_go_fuzz_dep_.CoverTab[44467]++
								return w.pc, w.err
//line /usr/local/go/src/net/http/transport.go:1407
		// _ = "end of CoverTab[44467]"
	case <-req.Cancel:
//line /usr/local/go/src/net/http/transport.go:1408
		_go_fuzz_dep_.CoverTab[44468]++
								return nil, errRequestCanceledConn
//line /usr/local/go/src/net/http/transport.go:1409
		// _ = "end of CoverTab[44468]"
	case <-req.Context().Done():
//line /usr/local/go/src/net/http/transport.go:1410
		_go_fuzz_dep_.CoverTab[44469]++
								return nil, req.Context().Err()
//line /usr/local/go/src/net/http/transport.go:1411
		// _ = "end of CoverTab[44469]"
	case err := <-cancelc:
//line /usr/local/go/src/net/http/transport.go:1412
		_go_fuzz_dep_.CoverTab[44470]++
								if err == errRequestCanceled {
//line /usr/local/go/src/net/http/transport.go:1413
			_go_fuzz_dep_.CoverTab[44486]++
									err = errRequestCanceledConn
//line /usr/local/go/src/net/http/transport.go:1414
			// _ = "end of CoverTab[44486]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1415
			_go_fuzz_dep_.CoverTab[44487]++
//line /usr/local/go/src/net/http/transport.go:1415
			// _ = "end of CoverTab[44487]"
//line /usr/local/go/src/net/http/transport.go:1415
		}
//line /usr/local/go/src/net/http/transport.go:1415
		// _ = "end of CoverTab[44470]"
//line /usr/local/go/src/net/http/transport.go:1415
		_go_fuzz_dep_.CoverTab[44471]++
								return nil, err
//line /usr/local/go/src/net/http/transport.go:1416
		// _ = "end of CoverTab[44471]"
	}
//line /usr/local/go/src/net/http/transport.go:1417
	// _ = "end of CoverTab[44448]"
}

// queueForDial queues w to wait for permission to begin dialing.
//line /usr/local/go/src/net/http/transport.go:1420
// Once w receives permission to dial, it will do so in a separate goroutine.
//line /usr/local/go/src/net/http/transport.go:1422
func (t *Transport) queueForDial(w *wantConn) {
//line /usr/local/go/src/net/http/transport.go:1422
	_go_fuzz_dep_.CoverTab[44488]++
							w.beforeDial()
							if t.MaxConnsPerHost <= 0 {
//line /usr/local/go/src/net/http/transport.go:1424
		_go_fuzz_dep_.CoverTab[44492]++
//line /usr/local/go/src/net/http/transport.go:1424
		_curRoutineNum38_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/transport.go:1424
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum38_)
								go t.dialConnFor(w)
								return
//line /usr/local/go/src/net/http/transport.go:1426
		// _ = "end of CoverTab[44492]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1427
		_go_fuzz_dep_.CoverTab[44493]++
//line /usr/local/go/src/net/http/transport.go:1427
		// _ = "end of CoverTab[44493]"
//line /usr/local/go/src/net/http/transport.go:1427
	}
//line /usr/local/go/src/net/http/transport.go:1427
	// _ = "end of CoverTab[44488]"
//line /usr/local/go/src/net/http/transport.go:1427
	_go_fuzz_dep_.CoverTab[44489]++

							t.connsPerHostMu.Lock()
							defer t.connsPerHostMu.Unlock()

							if n := t.connsPerHost[w.key]; n < t.MaxConnsPerHost {
//line /usr/local/go/src/net/http/transport.go:1432
		_go_fuzz_dep_.CoverTab[44494]++
								if t.connsPerHost == nil {
//line /usr/local/go/src/net/http/transport.go:1433
			_go_fuzz_dep_.CoverTab[44496]++
									t.connsPerHost = make(map[connectMethodKey]int)
//line /usr/local/go/src/net/http/transport.go:1434
			// _ = "end of CoverTab[44496]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1435
			_go_fuzz_dep_.CoverTab[44497]++
//line /usr/local/go/src/net/http/transport.go:1435
			// _ = "end of CoverTab[44497]"
//line /usr/local/go/src/net/http/transport.go:1435
		}
//line /usr/local/go/src/net/http/transport.go:1435
		// _ = "end of CoverTab[44494]"
//line /usr/local/go/src/net/http/transport.go:1435
		_go_fuzz_dep_.CoverTab[44495]++
								t.connsPerHost[w.key] = n + 1
//line /usr/local/go/src/net/http/transport.go:1436
		_curRoutineNum39_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/transport.go:1436
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum39_)
								go t.dialConnFor(w)
								return
//line /usr/local/go/src/net/http/transport.go:1438
		// _ = "end of CoverTab[44495]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1439
		_go_fuzz_dep_.CoverTab[44498]++
//line /usr/local/go/src/net/http/transport.go:1439
		// _ = "end of CoverTab[44498]"
//line /usr/local/go/src/net/http/transport.go:1439
	}
//line /usr/local/go/src/net/http/transport.go:1439
	// _ = "end of CoverTab[44489]"
//line /usr/local/go/src/net/http/transport.go:1439
	_go_fuzz_dep_.CoverTab[44490]++

							if t.connsPerHostWait == nil {
//line /usr/local/go/src/net/http/transport.go:1441
		_go_fuzz_dep_.CoverTab[44499]++
								t.connsPerHostWait = make(map[connectMethodKey]wantConnQueue)
//line /usr/local/go/src/net/http/transport.go:1442
		// _ = "end of CoverTab[44499]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1443
		_go_fuzz_dep_.CoverTab[44500]++
//line /usr/local/go/src/net/http/transport.go:1443
		// _ = "end of CoverTab[44500]"
//line /usr/local/go/src/net/http/transport.go:1443
	}
//line /usr/local/go/src/net/http/transport.go:1443
	// _ = "end of CoverTab[44490]"
//line /usr/local/go/src/net/http/transport.go:1443
	_go_fuzz_dep_.CoverTab[44491]++
							q := t.connsPerHostWait[w.key]
							q.cleanFront()
							q.pushBack(w)
							t.connsPerHostWait[w.key] = q
//line /usr/local/go/src/net/http/transport.go:1447
	// _ = "end of CoverTab[44491]"
}

// dialConnFor dials on behalf of w and delivers the result to w.
//line /usr/local/go/src/net/http/transport.go:1450
// dialConnFor has received permission to dial w.cm and is counted in t.connCount[w.cm.key()].
//line /usr/local/go/src/net/http/transport.go:1450
// If the dial is canceled or unsuccessful, dialConnFor decrements t.connCount[w.cm.key()].
//line /usr/local/go/src/net/http/transport.go:1453
func (t *Transport) dialConnFor(w *wantConn) {
//line /usr/local/go/src/net/http/transport.go:1453
	_go_fuzz_dep_.CoverTab[44501]++
							defer w.afterDial()

							pc, err := t.dialConn(w.ctx, w.cm)
							delivered := w.tryDeliver(pc, err)
							if err == nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1458
		_go_fuzz_dep_.CoverTab[44503]++
//line /usr/local/go/src/net/http/transport.go:1458
		return (!delivered || func() bool {
//line /usr/local/go/src/net/http/transport.go:1458
			_go_fuzz_dep_.CoverTab[44504]++
//line /usr/local/go/src/net/http/transport.go:1458
			return pc.alt != nil
//line /usr/local/go/src/net/http/transport.go:1458
			// _ = "end of CoverTab[44504]"
//line /usr/local/go/src/net/http/transport.go:1458
		}())
//line /usr/local/go/src/net/http/transport.go:1458
		// _ = "end of CoverTab[44503]"
//line /usr/local/go/src/net/http/transport.go:1458
	}() {
//line /usr/local/go/src/net/http/transport.go:1458
		_go_fuzz_dep_.CoverTab[44505]++

//line /usr/local/go/src/net/http/transport.go:1462
		t.putOrCloseIdleConn(pc)
//line /usr/local/go/src/net/http/transport.go:1462
		// _ = "end of CoverTab[44505]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1463
		_go_fuzz_dep_.CoverTab[44506]++
//line /usr/local/go/src/net/http/transport.go:1463
		// _ = "end of CoverTab[44506]"
//line /usr/local/go/src/net/http/transport.go:1463
	}
//line /usr/local/go/src/net/http/transport.go:1463
	// _ = "end of CoverTab[44501]"
//line /usr/local/go/src/net/http/transport.go:1463
	_go_fuzz_dep_.CoverTab[44502]++
							if err != nil {
//line /usr/local/go/src/net/http/transport.go:1464
		_go_fuzz_dep_.CoverTab[44507]++
								t.decConnsPerHost(w.key)
//line /usr/local/go/src/net/http/transport.go:1465
		// _ = "end of CoverTab[44507]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1466
		_go_fuzz_dep_.CoverTab[44508]++
//line /usr/local/go/src/net/http/transport.go:1466
		// _ = "end of CoverTab[44508]"
//line /usr/local/go/src/net/http/transport.go:1466
	}
//line /usr/local/go/src/net/http/transport.go:1466
	// _ = "end of CoverTab[44502]"
}

// decConnsPerHost decrements the per-host connection count for key,
//line /usr/local/go/src/net/http/transport.go:1469
// which may in turn give a different waiting goroutine permission to dial.
//line /usr/local/go/src/net/http/transport.go:1471
func (t *Transport) decConnsPerHost(key connectMethodKey) {
//line /usr/local/go/src/net/http/transport.go:1471
	_go_fuzz_dep_.CoverTab[44509]++
							if t.MaxConnsPerHost <= 0 {
//line /usr/local/go/src/net/http/transport.go:1472
		_go_fuzz_dep_.CoverTab[44513]++
								return
//line /usr/local/go/src/net/http/transport.go:1473
		// _ = "end of CoverTab[44513]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1474
		_go_fuzz_dep_.CoverTab[44514]++
//line /usr/local/go/src/net/http/transport.go:1474
		// _ = "end of CoverTab[44514]"
//line /usr/local/go/src/net/http/transport.go:1474
	}
//line /usr/local/go/src/net/http/transport.go:1474
	// _ = "end of CoverTab[44509]"
//line /usr/local/go/src/net/http/transport.go:1474
	_go_fuzz_dep_.CoverTab[44510]++

							t.connsPerHostMu.Lock()
							defer t.connsPerHostMu.Unlock()
							n := t.connsPerHost[key]
							if n == 0 {
//line /usr/local/go/src/net/http/transport.go:1479
		_go_fuzz_dep_.CoverTab[44515]++

//line /usr/local/go/src/net/http/transport.go:1482
		panic("net/http: internal error: connCount underflow")
//line /usr/local/go/src/net/http/transport.go:1482
		// _ = "end of CoverTab[44515]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1483
		_go_fuzz_dep_.CoverTab[44516]++
//line /usr/local/go/src/net/http/transport.go:1483
		// _ = "end of CoverTab[44516]"
//line /usr/local/go/src/net/http/transport.go:1483
	}
//line /usr/local/go/src/net/http/transport.go:1483
	// _ = "end of CoverTab[44510]"
//line /usr/local/go/src/net/http/transport.go:1483
	_go_fuzz_dep_.CoverTab[44511]++

//line /usr/local/go/src/net/http/transport.go:1489
	if q := t.connsPerHostWait[key]; q.len() > 0 {
//line /usr/local/go/src/net/http/transport.go:1489
		_go_fuzz_dep_.CoverTab[44517]++
								done := false
								for q.len() > 0 {
//line /usr/local/go/src/net/http/transport.go:1491
			_go_fuzz_dep_.CoverTab[44520]++
									w := q.popFront()
									if w.waiting() {
//line /usr/local/go/src/net/http/transport.go:1493
				_go_fuzz_dep_.CoverTab[44521]++
//line /usr/local/go/src/net/http/transport.go:1493
				_curRoutineNum40_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/transport.go:1493
				_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum40_)
										go t.dialConnFor(w)
										done = true
										break
//line /usr/local/go/src/net/http/transport.go:1496
				// _ = "end of CoverTab[44521]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1497
				_go_fuzz_dep_.CoverTab[44522]++
//line /usr/local/go/src/net/http/transport.go:1497
				// _ = "end of CoverTab[44522]"
//line /usr/local/go/src/net/http/transport.go:1497
			}
//line /usr/local/go/src/net/http/transport.go:1497
			// _ = "end of CoverTab[44520]"
		}
//line /usr/local/go/src/net/http/transport.go:1498
		// _ = "end of CoverTab[44517]"
//line /usr/local/go/src/net/http/transport.go:1498
		_go_fuzz_dep_.CoverTab[44518]++
								if q.len() == 0 {
//line /usr/local/go/src/net/http/transport.go:1499
			_go_fuzz_dep_.CoverTab[44523]++
									delete(t.connsPerHostWait, key)
//line /usr/local/go/src/net/http/transport.go:1500
			// _ = "end of CoverTab[44523]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1501
			_go_fuzz_dep_.CoverTab[44524]++

//line /usr/local/go/src/net/http/transport.go:1504
			t.connsPerHostWait[key] = q
//line /usr/local/go/src/net/http/transport.go:1504
			// _ = "end of CoverTab[44524]"
		}
//line /usr/local/go/src/net/http/transport.go:1505
		// _ = "end of CoverTab[44518]"
//line /usr/local/go/src/net/http/transport.go:1505
		_go_fuzz_dep_.CoverTab[44519]++
								if done {
//line /usr/local/go/src/net/http/transport.go:1506
			_go_fuzz_dep_.CoverTab[44525]++
									return
//line /usr/local/go/src/net/http/transport.go:1507
			// _ = "end of CoverTab[44525]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1508
			_go_fuzz_dep_.CoverTab[44526]++
//line /usr/local/go/src/net/http/transport.go:1508
			// _ = "end of CoverTab[44526]"
//line /usr/local/go/src/net/http/transport.go:1508
		}
//line /usr/local/go/src/net/http/transport.go:1508
		// _ = "end of CoverTab[44519]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1509
		_go_fuzz_dep_.CoverTab[44527]++
//line /usr/local/go/src/net/http/transport.go:1509
		// _ = "end of CoverTab[44527]"
//line /usr/local/go/src/net/http/transport.go:1509
	}
//line /usr/local/go/src/net/http/transport.go:1509
	// _ = "end of CoverTab[44511]"
//line /usr/local/go/src/net/http/transport.go:1509
	_go_fuzz_dep_.CoverTab[44512]++

//line /usr/local/go/src/net/http/transport.go:1512
	if n--; n == 0 {
//line /usr/local/go/src/net/http/transport.go:1512
		_go_fuzz_dep_.CoverTab[44528]++
								delete(t.connsPerHost, key)
//line /usr/local/go/src/net/http/transport.go:1513
		// _ = "end of CoverTab[44528]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1514
		_go_fuzz_dep_.CoverTab[44529]++
								t.connsPerHost[key] = n
//line /usr/local/go/src/net/http/transport.go:1515
		// _ = "end of CoverTab[44529]"
	}
//line /usr/local/go/src/net/http/transport.go:1516
	// _ = "end of CoverTab[44512]"
}

// Add TLS to a persistent connection, i.e. negotiate a TLS session. If pconn is already a TLS
//line /usr/local/go/src/net/http/transport.go:1519
// tunnel, this function establishes a nested TLS session inside the encrypted channel.
//line /usr/local/go/src/net/http/transport.go:1519
// The remote endpoint's name may be overridden by TLSClientConfig.ServerName.
//line /usr/local/go/src/net/http/transport.go:1522
func (pconn *persistConn) addTLS(ctx context.Context, name string, trace *httptrace.ClientTrace) error {
//line /usr/local/go/src/net/http/transport.go:1522
	_go_fuzz_dep_.CoverTab[44530]++

							cfg := cloneTLSConfig(pconn.t.TLSClientConfig)
							if cfg.ServerName == "" {
//line /usr/local/go/src/net/http/transport.go:1525
		_go_fuzz_dep_.CoverTab[44537]++
								cfg.ServerName = name
//line /usr/local/go/src/net/http/transport.go:1526
		// _ = "end of CoverTab[44537]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1527
		_go_fuzz_dep_.CoverTab[44538]++
//line /usr/local/go/src/net/http/transport.go:1527
		// _ = "end of CoverTab[44538]"
//line /usr/local/go/src/net/http/transport.go:1527
	}
//line /usr/local/go/src/net/http/transport.go:1527
	// _ = "end of CoverTab[44530]"
//line /usr/local/go/src/net/http/transport.go:1527
	_go_fuzz_dep_.CoverTab[44531]++
							if pconn.cacheKey.onlyH1 {
//line /usr/local/go/src/net/http/transport.go:1528
		_go_fuzz_dep_.CoverTab[44539]++
								cfg.NextProtos = nil
//line /usr/local/go/src/net/http/transport.go:1529
		// _ = "end of CoverTab[44539]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1530
		_go_fuzz_dep_.CoverTab[44540]++
//line /usr/local/go/src/net/http/transport.go:1530
		// _ = "end of CoverTab[44540]"
//line /usr/local/go/src/net/http/transport.go:1530
	}
//line /usr/local/go/src/net/http/transport.go:1530
	// _ = "end of CoverTab[44531]"
//line /usr/local/go/src/net/http/transport.go:1530
	_go_fuzz_dep_.CoverTab[44532]++
							plainConn := pconn.conn
							tlsConn := tls.Client(plainConn, cfg)
							errc := make(chan error, 2)
							var timer *time.Timer	// for canceling TLS handshake
							if d := pconn.t.TLSHandshakeTimeout; d != 0 {
//line /usr/local/go/src/net/http/transport.go:1535
		_go_fuzz_dep_.CoverTab[44541]++
								timer = time.AfterFunc(d, func() {
//line /usr/local/go/src/net/http/transport.go:1536
			_go_fuzz_dep_.CoverTab[44542]++
									errc <- tlsHandshakeTimeoutError{}
//line /usr/local/go/src/net/http/transport.go:1537
			// _ = "end of CoverTab[44542]"
		})
//line /usr/local/go/src/net/http/transport.go:1538
		// _ = "end of CoverTab[44541]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1539
		_go_fuzz_dep_.CoverTab[44543]++
//line /usr/local/go/src/net/http/transport.go:1539
		// _ = "end of CoverTab[44543]"
//line /usr/local/go/src/net/http/transport.go:1539
	}
//line /usr/local/go/src/net/http/transport.go:1539
	// _ = "end of CoverTab[44532]"
//line /usr/local/go/src/net/http/transport.go:1539
	_go_fuzz_dep_.CoverTab[44533]++
//line /usr/local/go/src/net/http/transport.go:1539
	_curRoutineNum41_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/transport.go:1539
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum41_)
							go func() {
//line /usr/local/go/src/net/http/transport.go:1540
		_go_fuzz_dep_.CoverTab[44544]++
//line /usr/local/go/src/net/http/transport.go:1540
		defer func() {
//line /usr/local/go/src/net/http/transport.go:1540
			_go_fuzz_dep_.CoverTab[44547]++
//line /usr/local/go/src/net/http/transport.go:1540
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum41_)
//line /usr/local/go/src/net/http/transport.go:1540
			// _ = "end of CoverTab[44547]"
//line /usr/local/go/src/net/http/transport.go:1540
		}()
								if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1541
			_go_fuzz_dep_.CoverTab[44548]++
//line /usr/local/go/src/net/http/transport.go:1541
			return trace.TLSHandshakeStart != nil
//line /usr/local/go/src/net/http/transport.go:1541
			// _ = "end of CoverTab[44548]"
//line /usr/local/go/src/net/http/transport.go:1541
		}() {
//line /usr/local/go/src/net/http/transport.go:1541
			_go_fuzz_dep_.CoverTab[44549]++
									trace.TLSHandshakeStart()
//line /usr/local/go/src/net/http/transport.go:1542
			// _ = "end of CoverTab[44549]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1543
			_go_fuzz_dep_.CoverTab[44550]++
//line /usr/local/go/src/net/http/transport.go:1543
			// _ = "end of CoverTab[44550]"
//line /usr/local/go/src/net/http/transport.go:1543
		}
//line /usr/local/go/src/net/http/transport.go:1543
		// _ = "end of CoverTab[44544]"
//line /usr/local/go/src/net/http/transport.go:1543
		_go_fuzz_dep_.CoverTab[44545]++
								err := tlsConn.HandshakeContext(ctx)
								if timer != nil {
//line /usr/local/go/src/net/http/transport.go:1545
			_go_fuzz_dep_.CoverTab[44551]++
									timer.Stop()
//line /usr/local/go/src/net/http/transport.go:1546
			// _ = "end of CoverTab[44551]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1547
			_go_fuzz_dep_.CoverTab[44552]++
//line /usr/local/go/src/net/http/transport.go:1547
			// _ = "end of CoverTab[44552]"
//line /usr/local/go/src/net/http/transport.go:1547
		}
//line /usr/local/go/src/net/http/transport.go:1547
		// _ = "end of CoverTab[44545]"
//line /usr/local/go/src/net/http/transport.go:1547
		_go_fuzz_dep_.CoverTab[44546]++
								errc <- err
//line /usr/local/go/src/net/http/transport.go:1548
		// _ = "end of CoverTab[44546]"
	}()
//line /usr/local/go/src/net/http/transport.go:1549
	// _ = "end of CoverTab[44533]"
//line /usr/local/go/src/net/http/transport.go:1549
	_go_fuzz_dep_.CoverTab[44534]++
							if err := <-errc; err != nil {
//line /usr/local/go/src/net/http/transport.go:1550
		_go_fuzz_dep_.CoverTab[44553]++
								plainConn.Close()
								if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1552
			_go_fuzz_dep_.CoverTab[44555]++
//line /usr/local/go/src/net/http/transport.go:1552
			return trace.TLSHandshakeDone != nil
//line /usr/local/go/src/net/http/transport.go:1552
			// _ = "end of CoverTab[44555]"
//line /usr/local/go/src/net/http/transport.go:1552
		}() {
//line /usr/local/go/src/net/http/transport.go:1552
			_go_fuzz_dep_.CoverTab[44556]++
									trace.TLSHandshakeDone(tls.ConnectionState{}, err)
//line /usr/local/go/src/net/http/transport.go:1553
			// _ = "end of CoverTab[44556]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1554
			_go_fuzz_dep_.CoverTab[44557]++
//line /usr/local/go/src/net/http/transport.go:1554
			// _ = "end of CoverTab[44557]"
//line /usr/local/go/src/net/http/transport.go:1554
		}
//line /usr/local/go/src/net/http/transport.go:1554
		// _ = "end of CoverTab[44553]"
//line /usr/local/go/src/net/http/transport.go:1554
		_go_fuzz_dep_.CoverTab[44554]++
								return err
//line /usr/local/go/src/net/http/transport.go:1555
		// _ = "end of CoverTab[44554]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1556
		_go_fuzz_dep_.CoverTab[44558]++
//line /usr/local/go/src/net/http/transport.go:1556
		// _ = "end of CoverTab[44558]"
//line /usr/local/go/src/net/http/transport.go:1556
	}
//line /usr/local/go/src/net/http/transport.go:1556
	// _ = "end of CoverTab[44534]"
//line /usr/local/go/src/net/http/transport.go:1556
	_go_fuzz_dep_.CoverTab[44535]++
							cs := tlsConn.ConnectionState()
							if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1558
		_go_fuzz_dep_.CoverTab[44559]++
//line /usr/local/go/src/net/http/transport.go:1558
		return trace.TLSHandshakeDone != nil
//line /usr/local/go/src/net/http/transport.go:1558
		// _ = "end of CoverTab[44559]"
//line /usr/local/go/src/net/http/transport.go:1558
	}() {
//line /usr/local/go/src/net/http/transport.go:1558
		_go_fuzz_dep_.CoverTab[44560]++
								trace.TLSHandshakeDone(cs, nil)
//line /usr/local/go/src/net/http/transport.go:1559
		// _ = "end of CoverTab[44560]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1560
		_go_fuzz_dep_.CoverTab[44561]++
//line /usr/local/go/src/net/http/transport.go:1560
		// _ = "end of CoverTab[44561]"
//line /usr/local/go/src/net/http/transport.go:1560
	}
//line /usr/local/go/src/net/http/transport.go:1560
	// _ = "end of CoverTab[44535]"
//line /usr/local/go/src/net/http/transport.go:1560
	_go_fuzz_dep_.CoverTab[44536]++
							pconn.tlsState = &cs
							pconn.conn = tlsConn
							return nil
//line /usr/local/go/src/net/http/transport.go:1563
	// _ = "end of CoverTab[44536]"
}

type erringRoundTripper interface {
	RoundTripErr() error
}

func (t *Transport) dialConn(ctx context.Context, cm connectMethod) (pconn *persistConn, err error) {
//line /usr/local/go/src/net/http/transport.go:1570
	_go_fuzz_dep_.CoverTab[44562]++
							pconn = &persistConn{
		t:		t,
		cacheKey:	cm.key(),
		reqch:		make(chan requestAndChan, 1),
		writech:	make(chan writeRequest, 1),
		closech:	make(chan struct{}),
		writeErrCh:	make(chan error, 1),
		writeLoopDone:	make(chan struct{}),
	}
	trace := httptrace.ContextClientTrace(ctx)
	wrapErr := func(err error) error {
//line /usr/local/go/src/net/http/transport.go:1581
		_go_fuzz_dep_.CoverTab[44568]++
								if cm.proxyURL != nil {
//line /usr/local/go/src/net/http/transport.go:1582
			_go_fuzz_dep_.CoverTab[44570]++

									return &net.OpError{Op: "proxyconnect", Net: "tcp", Err: err}
//line /usr/local/go/src/net/http/transport.go:1584
			// _ = "end of CoverTab[44570]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1585
			_go_fuzz_dep_.CoverTab[44571]++
//line /usr/local/go/src/net/http/transport.go:1585
			// _ = "end of CoverTab[44571]"
//line /usr/local/go/src/net/http/transport.go:1585
		}
//line /usr/local/go/src/net/http/transport.go:1585
		// _ = "end of CoverTab[44568]"
//line /usr/local/go/src/net/http/transport.go:1585
		_go_fuzz_dep_.CoverTab[44569]++
								return err
//line /usr/local/go/src/net/http/transport.go:1586
		// _ = "end of CoverTab[44569]"
	}
//line /usr/local/go/src/net/http/transport.go:1587
	// _ = "end of CoverTab[44562]"
//line /usr/local/go/src/net/http/transport.go:1587
	_go_fuzz_dep_.CoverTab[44563]++
							if cm.scheme() == "https" && func() bool {
//line /usr/local/go/src/net/http/transport.go:1588
		_go_fuzz_dep_.CoverTab[44572]++
//line /usr/local/go/src/net/http/transport.go:1588
		return t.hasCustomTLSDialer()
//line /usr/local/go/src/net/http/transport.go:1588
		// _ = "end of CoverTab[44572]"
//line /usr/local/go/src/net/http/transport.go:1588
	}() {
//line /usr/local/go/src/net/http/transport.go:1588
		_go_fuzz_dep_.CoverTab[44573]++
								var err error
								pconn.conn, err = t.customDialTLS(ctx, "tcp", cm.addr())
								if err != nil {
//line /usr/local/go/src/net/http/transport.go:1591
			_go_fuzz_dep_.CoverTab[44575]++
									return nil, wrapErr(err)
//line /usr/local/go/src/net/http/transport.go:1592
			// _ = "end of CoverTab[44575]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1593
			_go_fuzz_dep_.CoverTab[44576]++
//line /usr/local/go/src/net/http/transport.go:1593
			// _ = "end of CoverTab[44576]"
//line /usr/local/go/src/net/http/transport.go:1593
		}
//line /usr/local/go/src/net/http/transport.go:1593
		// _ = "end of CoverTab[44573]"
//line /usr/local/go/src/net/http/transport.go:1593
		_go_fuzz_dep_.CoverTab[44574]++
								if tc, ok := pconn.conn.(*tls.Conn); ok {
//line /usr/local/go/src/net/http/transport.go:1594
			_go_fuzz_dep_.CoverTab[44577]++

//line /usr/local/go/src/net/http/transport.go:1597
			if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1597
				_go_fuzz_dep_.CoverTab[44581]++
//line /usr/local/go/src/net/http/transport.go:1597
				return trace.TLSHandshakeStart != nil
//line /usr/local/go/src/net/http/transport.go:1597
				// _ = "end of CoverTab[44581]"
//line /usr/local/go/src/net/http/transport.go:1597
			}() {
//line /usr/local/go/src/net/http/transport.go:1597
				_go_fuzz_dep_.CoverTab[44582]++
										trace.TLSHandshakeStart()
//line /usr/local/go/src/net/http/transport.go:1598
				// _ = "end of CoverTab[44582]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1599
				_go_fuzz_dep_.CoverTab[44583]++
//line /usr/local/go/src/net/http/transport.go:1599
				// _ = "end of CoverTab[44583]"
//line /usr/local/go/src/net/http/transport.go:1599
			}
//line /usr/local/go/src/net/http/transport.go:1599
			// _ = "end of CoverTab[44577]"
//line /usr/local/go/src/net/http/transport.go:1599
			_go_fuzz_dep_.CoverTab[44578]++
									if err := tc.HandshakeContext(ctx); err != nil {
//line /usr/local/go/src/net/http/transport.go:1600
				_go_fuzz_dep_.CoverTab[44584]++
//line /usr/local/go/src/net/http/transport.go:1600
				_curRoutineNum44_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/transport.go:1600
				_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum44_)
										go pconn.conn.Close()
										if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1602
					_go_fuzz_dep_.CoverTab[44586]++
//line /usr/local/go/src/net/http/transport.go:1602
					return trace.TLSHandshakeDone != nil
//line /usr/local/go/src/net/http/transport.go:1602
					// _ = "end of CoverTab[44586]"
//line /usr/local/go/src/net/http/transport.go:1602
				}() {
//line /usr/local/go/src/net/http/transport.go:1602
					_go_fuzz_dep_.CoverTab[44587]++
											trace.TLSHandshakeDone(tls.ConnectionState{}, err)
//line /usr/local/go/src/net/http/transport.go:1603
					// _ = "end of CoverTab[44587]"
				} else {
//line /usr/local/go/src/net/http/transport.go:1604
					_go_fuzz_dep_.CoverTab[44588]++
//line /usr/local/go/src/net/http/transport.go:1604
					// _ = "end of CoverTab[44588]"
//line /usr/local/go/src/net/http/transport.go:1604
				}
//line /usr/local/go/src/net/http/transport.go:1604
				// _ = "end of CoverTab[44584]"
//line /usr/local/go/src/net/http/transport.go:1604
				_go_fuzz_dep_.CoverTab[44585]++
										return nil, err
//line /usr/local/go/src/net/http/transport.go:1605
				// _ = "end of CoverTab[44585]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1606
				_go_fuzz_dep_.CoverTab[44589]++
//line /usr/local/go/src/net/http/transport.go:1606
				// _ = "end of CoverTab[44589]"
//line /usr/local/go/src/net/http/transport.go:1606
			}
//line /usr/local/go/src/net/http/transport.go:1606
			// _ = "end of CoverTab[44578]"
//line /usr/local/go/src/net/http/transport.go:1606
			_go_fuzz_dep_.CoverTab[44579]++
									cs := tc.ConnectionState()
									if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1608
				_go_fuzz_dep_.CoverTab[44590]++
//line /usr/local/go/src/net/http/transport.go:1608
				return trace.TLSHandshakeDone != nil
//line /usr/local/go/src/net/http/transport.go:1608
				// _ = "end of CoverTab[44590]"
//line /usr/local/go/src/net/http/transport.go:1608
			}() {
//line /usr/local/go/src/net/http/transport.go:1608
				_go_fuzz_dep_.CoverTab[44591]++
										trace.TLSHandshakeDone(cs, nil)
//line /usr/local/go/src/net/http/transport.go:1609
				// _ = "end of CoverTab[44591]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1610
				_go_fuzz_dep_.CoverTab[44592]++
//line /usr/local/go/src/net/http/transport.go:1610
				// _ = "end of CoverTab[44592]"
//line /usr/local/go/src/net/http/transport.go:1610
			}
//line /usr/local/go/src/net/http/transport.go:1610
			// _ = "end of CoverTab[44579]"
//line /usr/local/go/src/net/http/transport.go:1610
			_go_fuzz_dep_.CoverTab[44580]++
									pconn.tlsState = &cs
//line /usr/local/go/src/net/http/transport.go:1611
			// _ = "end of CoverTab[44580]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1612
			_go_fuzz_dep_.CoverTab[44593]++
//line /usr/local/go/src/net/http/transport.go:1612
			// _ = "end of CoverTab[44593]"
//line /usr/local/go/src/net/http/transport.go:1612
		}
//line /usr/local/go/src/net/http/transport.go:1612
		// _ = "end of CoverTab[44574]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1613
		_go_fuzz_dep_.CoverTab[44594]++
								conn, err := t.dial(ctx, "tcp", cm.addr())
								if err != nil {
//line /usr/local/go/src/net/http/transport.go:1615
			_go_fuzz_dep_.CoverTab[44596]++
									return nil, wrapErr(err)
//line /usr/local/go/src/net/http/transport.go:1616
			// _ = "end of CoverTab[44596]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1617
			_go_fuzz_dep_.CoverTab[44597]++
//line /usr/local/go/src/net/http/transport.go:1617
			// _ = "end of CoverTab[44597]"
//line /usr/local/go/src/net/http/transport.go:1617
		}
//line /usr/local/go/src/net/http/transport.go:1617
		// _ = "end of CoverTab[44594]"
//line /usr/local/go/src/net/http/transport.go:1617
		_go_fuzz_dep_.CoverTab[44595]++
								pconn.conn = conn
								if cm.scheme() == "https" {
//line /usr/local/go/src/net/http/transport.go:1619
			_go_fuzz_dep_.CoverTab[44598]++
									var firstTLSHost string
									if firstTLSHost, _, err = net.SplitHostPort(cm.addr()); err != nil {
//line /usr/local/go/src/net/http/transport.go:1621
				_go_fuzz_dep_.CoverTab[44600]++
										return nil, wrapErr(err)
//line /usr/local/go/src/net/http/transport.go:1622
				// _ = "end of CoverTab[44600]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1623
				_go_fuzz_dep_.CoverTab[44601]++
//line /usr/local/go/src/net/http/transport.go:1623
				// _ = "end of CoverTab[44601]"
//line /usr/local/go/src/net/http/transport.go:1623
			}
//line /usr/local/go/src/net/http/transport.go:1623
			// _ = "end of CoverTab[44598]"
//line /usr/local/go/src/net/http/transport.go:1623
			_go_fuzz_dep_.CoverTab[44599]++
									if err = pconn.addTLS(ctx, firstTLSHost, trace); err != nil {
//line /usr/local/go/src/net/http/transport.go:1624
				_go_fuzz_dep_.CoverTab[44602]++
										return nil, wrapErr(err)
//line /usr/local/go/src/net/http/transport.go:1625
				// _ = "end of CoverTab[44602]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1626
				_go_fuzz_dep_.CoverTab[44603]++
//line /usr/local/go/src/net/http/transport.go:1626
				// _ = "end of CoverTab[44603]"
//line /usr/local/go/src/net/http/transport.go:1626
			}
//line /usr/local/go/src/net/http/transport.go:1626
			// _ = "end of CoverTab[44599]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1627
			_go_fuzz_dep_.CoverTab[44604]++
//line /usr/local/go/src/net/http/transport.go:1627
			// _ = "end of CoverTab[44604]"
//line /usr/local/go/src/net/http/transport.go:1627
		}
//line /usr/local/go/src/net/http/transport.go:1627
		// _ = "end of CoverTab[44595]"
	}
//line /usr/local/go/src/net/http/transport.go:1628
	// _ = "end of CoverTab[44563]"
//line /usr/local/go/src/net/http/transport.go:1628
	_go_fuzz_dep_.CoverTab[44564]++

//line /usr/local/go/src/net/http/transport.go:1631
	switch {
	case cm.proxyURL == nil:
//line /usr/local/go/src/net/http/transport.go:1632
		_go_fuzz_dep_.CoverTab[44605]++
//line /usr/local/go/src/net/http/transport.go:1632
		// _ = "end of CoverTab[44605]"

	case cm.proxyURL.Scheme == "socks5":
//line /usr/local/go/src/net/http/transport.go:1634
		_go_fuzz_dep_.CoverTab[44606]++
								conn := pconn.conn
								d := socksNewDialer("tcp", conn.RemoteAddr().String())
								if u := cm.proxyURL.User; u != nil {
//line /usr/local/go/src/net/http/transport.go:1637
			_go_fuzz_dep_.CoverTab[44619]++
									auth := &socksUsernamePassword{
				Username: u.Username(),
			}
			auth.Password, _ = u.Password()
			d.AuthMethods = []socksAuthMethod{
				socksAuthMethodNotRequired,
				socksAuthMethodUsernamePassword,
			}
									d.Authenticate = auth.Authenticate
//line /usr/local/go/src/net/http/transport.go:1646
			// _ = "end of CoverTab[44619]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1647
			_go_fuzz_dep_.CoverTab[44620]++
//line /usr/local/go/src/net/http/transport.go:1647
			// _ = "end of CoverTab[44620]"
//line /usr/local/go/src/net/http/transport.go:1647
		}
//line /usr/local/go/src/net/http/transport.go:1647
		// _ = "end of CoverTab[44606]"
//line /usr/local/go/src/net/http/transport.go:1647
		_go_fuzz_dep_.CoverTab[44607]++
								if _, err := d.DialWithConn(ctx, conn, "tcp", cm.targetAddr); err != nil {
//line /usr/local/go/src/net/http/transport.go:1648
			_go_fuzz_dep_.CoverTab[44621]++
									conn.Close()
									return nil, err
//line /usr/local/go/src/net/http/transport.go:1650
			// _ = "end of CoverTab[44621]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1651
			_go_fuzz_dep_.CoverTab[44622]++
//line /usr/local/go/src/net/http/transport.go:1651
			// _ = "end of CoverTab[44622]"
//line /usr/local/go/src/net/http/transport.go:1651
		}
//line /usr/local/go/src/net/http/transport.go:1651
		// _ = "end of CoverTab[44607]"
	case cm.targetScheme == "http":
//line /usr/local/go/src/net/http/transport.go:1652
		_go_fuzz_dep_.CoverTab[44608]++
								pconn.isProxy = true
								if pa := cm.proxyAuth(); pa != "" {
//line /usr/local/go/src/net/http/transport.go:1654
			_go_fuzz_dep_.CoverTab[44623]++
									pconn.mutateHeaderFunc = func(h Header) {
//line /usr/local/go/src/net/http/transport.go:1655
				_go_fuzz_dep_.CoverTab[44624]++
										h.Set("Proxy-Authorization", pa)
//line /usr/local/go/src/net/http/transport.go:1656
				// _ = "end of CoverTab[44624]"
			}
//line /usr/local/go/src/net/http/transport.go:1657
			// _ = "end of CoverTab[44623]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1658
			_go_fuzz_dep_.CoverTab[44625]++
//line /usr/local/go/src/net/http/transport.go:1658
			// _ = "end of CoverTab[44625]"
//line /usr/local/go/src/net/http/transport.go:1658
		}
//line /usr/local/go/src/net/http/transport.go:1658
		// _ = "end of CoverTab[44608]"
	case cm.targetScheme == "https":
//line /usr/local/go/src/net/http/transport.go:1659
		_go_fuzz_dep_.CoverTab[44609]++
								conn := pconn.conn
								var hdr Header
								if t.GetProxyConnectHeader != nil {
//line /usr/local/go/src/net/http/transport.go:1662
			_go_fuzz_dep_.CoverTab[44626]++
									var err error
									hdr, err = t.GetProxyConnectHeader(ctx, cm.proxyURL, cm.targetAddr)
									if err != nil {
//line /usr/local/go/src/net/http/transport.go:1665
				_go_fuzz_dep_.CoverTab[44627]++
										conn.Close()
										return nil, err
//line /usr/local/go/src/net/http/transport.go:1667
				// _ = "end of CoverTab[44627]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1668
				_go_fuzz_dep_.CoverTab[44628]++
//line /usr/local/go/src/net/http/transport.go:1668
				// _ = "end of CoverTab[44628]"
//line /usr/local/go/src/net/http/transport.go:1668
			}
//line /usr/local/go/src/net/http/transport.go:1668
			// _ = "end of CoverTab[44626]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1669
			_go_fuzz_dep_.CoverTab[44629]++
									hdr = t.ProxyConnectHeader
//line /usr/local/go/src/net/http/transport.go:1670
			// _ = "end of CoverTab[44629]"
		}
//line /usr/local/go/src/net/http/transport.go:1671
		// _ = "end of CoverTab[44609]"
//line /usr/local/go/src/net/http/transport.go:1671
		_go_fuzz_dep_.CoverTab[44610]++
								if hdr == nil {
//line /usr/local/go/src/net/http/transport.go:1672
			_go_fuzz_dep_.CoverTab[44630]++
									hdr = make(Header)
//line /usr/local/go/src/net/http/transport.go:1673
			// _ = "end of CoverTab[44630]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1674
			_go_fuzz_dep_.CoverTab[44631]++
//line /usr/local/go/src/net/http/transport.go:1674
			// _ = "end of CoverTab[44631]"
//line /usr/local/go/src/net/http/transport.go:1674
		}
//line /usr/local/go/src/net/http/transport.go:1674
		// _ = "end of CoverTab[44610]"
//line /usr/local/go/src/net/http/transport.go:1674
		_go_fuzz_dep_.CoverTab[44611]++
								if pa := cm.proxyAuth(); pa != "" {
//line /usr/local/go/src/net/http/transport.go:1675
			_go_fuzz_dep_.CoverTab[44632]++
									hdr = hdr.Clone()
									hdr.Set("Proxy-Authorization", pa)
//line /usr/local/go/src/net/http/transport.go:1677
			// _ = "end of CoverTab[44632]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1678
			_go_fuzz_dep_.CoverTab[44633]++
//line /usr/local/go/src/net/http/transport.go:1678
			// _ = "end of CoverTab[44633]"
//line /usr/local/go/src/net/http/transport.go:1678
		}
//line /usr/local/go/src/net/http/transport.go:1678
		// _ = "end of CoverTab[44611]"
//line /usr/local/go/src/net/http/transport.go:1678
		_go_fuzz_dep_.CoverTab[44612]++
								connectReq := &Request{
			Method:	"CONNECT",
			URL:	&url.URL{Opaque: cm.targetAddr},
			Host:	cm.targetAddr,
			Header:	hdr,
		}

//line /usr/local/go/src/net/http/transport.go:1691
		connectCtx := ctx
		if ctx.Done() == nil {
//line /usr/local/go/src/net/http/transport.go:1692
			_go_fuzz_dep_.CoverTab[44634]++
									newCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
									defer cancel()
									connectCtx = newCtx
//line /usr/local/go/src/net/http/transport.go:1695
			// _ = "end of CoverTab[44634]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1696
			_go_fuzz_dep_.CoverTab[44635]++
//line /usr/local/go/src/net/http/transport.go:1696
			// _ = "end of CoverTab[44635]"
//line /usr/local/go/src/net/http/transport.go:1696
		}
//line /usr/local/go/src/net/http/transport.go:1696
		// _ = "end of CoverTab[44612]"
//line /usr/local/go/src/net/http/transport.go:1696
		_go_fuzz_dep_.CoverTab[44613]++

								didReadResponse := make(chan struct{})
								var (
			resp	*Response
			err	error	// write or read error
		)

		go func() {
//line /usr/local/go/src/net/http/transport.go:1704
			_go_fuzz_dep_.CoverTab[44636]++
									defer close(didReadResponse)
									err = connectReq.Write(conn)
									if err != nil {
//line /usr/local/go/src/net/http/transport.go:1707
				_go_fuzz_dep_.CoverTab[44638]++
										return
//line /usr/local/go/src/net/http/transport.go:1708
				// _ = "end of CoverTab[44638]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1709
				_go_fuzz_dep_.CoverTab[44639]++
//line /usr/local/go/src/net/http/transport.go:1709
				// _ = "end of CoverTab[44639]"
//line /usr/local/go/src/net/http/transport.go:1709
			}
//line /usr/local/go/src/net/http/transport.go:1709
			// _ = "end of CoverTab[44636]"
//line /usr/local/go/src/net/http/transport.go:1709
			_go_fuzz_dep_.CoverTab[44637]++

//line /usr/local/go/src/net/http/transport.go:1712
			br := bufio.NewReader(conn)
									resp, err = ReadResponse(br, connectReq)
//line /usr/local/go/src/net/http/transport.go:1713
			// _ = "end of CoverTab[44637]"
		}()
//line /usr/local/go/src/net/http/transport.go:1714
		// _ = "end of CoverTab[44613]"
//line /usr/local/go/src/net/http/transport.go:1714
		_go_fuzz_dep_.CoverTab[44614]++
								select {
		case <-connectCtx.Done():
//line /usr/local/go/src/net/http/transport.go:1716
			_go_fuzz_dep_.CoverTab[44640]++
									conn.Close()
									<-didReadResponse
									return nil, connectCtx.Err()
//line /usr/local/go/src/net/http/transport.go:1719
			// _ = "end of CoverTab[44640]"
		case <-didReadResponse:
//line /usr/local/go/src/net/http/transport.go:1720
			_go_fuzz_dep_.CoverTab[44641]++
//line /usr/local/go/src/net/http/transport.go:1720
			// _ = "end of CoverTab[44641]"

		}
//line /usr/local/go/src/net/http/transport.go:1722
		// _ = "end of CoverTab[44614]"
//line /usr/local/go/src/net/http/transport.go:1722
		_go_fuzz_dep_.CoverTab[44615]++
								if err != nil {
//line /usr/local/go/src/net/http/transport.go:1723
			_go_fuzz_dep_.CoverTab[44642]++
									conn.Close()
									return nil, err
//line /usr/local/go/src/net/http/transport.go:1725
			// _ = "end of CoverTab[44642]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1726
			_go_fuzz_dep_.CoverTab[44643]++
//line /usr/local/go/src/net/http/transport.go:1726
			// _ = "end of CoverTab[44643]"
//line /usr/local/go/src/net/http/transport.go:1726
		}
//line /usr/local/go/src/net/http/transport.go:1726
		// _ = "end of CoverTab[44615]"
//line /usr/local/go/src/net/http/transport.go:1726
		_go_fuzz_dep_.CoverTab[44616]++

								if t.OnProxyConnectResponse != nil {
//line /usr/local/go/src/net/http/transport.go:1728
			_go_fuzz_dep_.CoverTab[44644]++
									err = t.OnProxyConnectResponse(ctx, cm.proxyURL, connectReq, resp)
									if err != nil {
//line /usr/local/go/src/net/http/transport.go:1730
				_go_fuzz_dep_.CoverTab[44645]++
										return nil, err
//line /usr/local/go/src/net/http/transport.go:1731
				// _ = "end of CoverTab[44645]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1732
				_go_fuzz_dep_.CoverTab[44646]++
//line /usr/local/go/src/net/http/transport.go:1732
				// _ = "end of CoverTab[44646]"
//line /usr/local/go/src/net/http/transport.go:1732
			}
//line /usr/local/go/src/net/http/transport.go:1732
			// _ = "end of CoverTab[44644]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1733
			_go_fuzz_dep_.CoverTab[44647]++
//line /usr/local/go/src/net/http/transport.go:1733
			// _ = "end of CoverTab[44647]"
//line /usr/local/go/src/net/http/transport.go:1733
		}
//line /usr/local/go/src/net/http/transport.go:1733
		// _ = "end of CoverTab[44616]"
//line /usr/local/go/src/net/http/transport.go:1733
		_go_fuzz_dep_.CoverTab[44617]++

								if resp.StatusCode != 200 {
//line /usr/local/go/src/net/http/transport.go:1735
			_go_fuzz_dep_.CoverTab[44648]++
									_, text, ok := strings.Cut(resp.Status, " ")
									conn.Close()
									if !ok {
//line /usr/local/go/src/net/http/transport.go:1738
				_go_fuzz_dep_.CoverTab[44650]++
										return nil, errors.New("unknown status code")
//line /usr/local/go/src/net/http/transport.go:1739
				// _ = "end of CoverTab[44650]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1740
				_go_fuzz_dep_.CoverTab[44651]++
//line /usr/local/go/src/net/http/transport.go:1740
				// _ = "end of CoverTab[44651]"
//line /usr/local/go/src/net/http/transport.go:1740
			}
//line /usr/local/go/src/net/http/transport.go:1740
			// _ = "end of CoverTab[44648]"
//line /usr/local/go/src/net/http/transport.go:1740
			_go_fuzz_dep_.CoverTab[44649]++
									return nil, errors.New(text)
//line /usr/local/go/src/net/http/transport.go:1741
			// _ = "end of CoverTab[44649]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1742
			_go_fuzz_dep_.CoverTab[44652]++
//line /usr/local/go/src/net/http/transport.go:1742
			// _ = "end of CoverTab[44652]"
//line /usr/local/go/src/net/http/transport.go:1742
		}
//line /usr/local/go/src/net/http/transport.go:1742
		// _ = "end of CoverTab[44617]"
//line /usr/local/go/src/net/http/transport.go:1742
	default:
//line /usr/local/go/src/net/http/transport.go:1742
		_go_fuzz_dep_.CoverTab[44618]++
//line /usr/local/go/src/net/http/transport.go:1742
		// _ = "end of CoverTab[44618]"
	}
//line /usr/local/go/src/net/http/transport.go:1743
	// _ = "end of CoverTab[44564]"
//line /usr/local/go/src/net/http/transport.go:1743
	_go_fuzz_dep_.CoverTab[44565]++

							if cm.proxyURL != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1745
		_go_fuzz_dep_.CoverTab[44653]++
//line /usr/local/go/src/net/http/transport.go:1745
		return cm.targetScheme == "https"
//line /usr/local/go/src/net/http/transport.go:1745
		// _ = "end of CoverTab[44653]"
//line /usr/local/go/src/net/http/transport.go:1745
	}() {
//line /usr/local/go/src/net/http/transport.go:1745
		_go_fuzz_dep_.CoverTab[44654]++
								if err := pconn.addTLS(ctx, cm.tlsHost(), trace); err != nil {
//line /usr/local/go/src/net/http/transport.go:1746
			_go_fuzz_dep_.CoverTab[44655]++
									return nil, err
//line /usr/local/go/src/net/http/transport.go:1747
			// _ = "end of CoverTab[44655]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1748
			_go_fuzz_dep_.CoverTab[44656]++
//line /usr/local/go/src/net/http/transport.go:1748
			// _ = "end of CoverTab[44656]"
//line /usr/local/go/src/net/http/transport.go:1748
		}
//line /usr/local/go/src/net/http/transport.go:1748
		// _ = "end of CoverTab[44654]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1749
		_go_fuzz_dep_.CoverTab[44657]++
//line /usr/local/go/src/net/http/transport.go:1749
		// _ = "end of CoverTab[44657]"
//line /usr/local/go/src/net/http/transport.go:1749
	}
//line /usr/local/go/src/net/http/transport.go:1749
	// _ = "end of CoverTab[44565]"
//line /usr/local/go/src/net/http/transport.go:1749
	_go_fuzz_dep_.CoverTab[44566]++

							if s := pconn.tlsState; s != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:1751
		_go_fuzz_dep_.CoverTab[44658]++
//line /usr/local/go/src/net/http/transport.go:1751
		return s.NegotiatedProtocolIsMutual
//line /usr/local/go/src/net/http/transport.go:1751
		// _ = "end of CoverTab[44658]"
//line /usr/local/go/src/net/http/transport.go:1751
	}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:1751
		_go_fuzz_dep_.CoverTab[44659]++
//line /usr/local/go/src/net/http/transport.go:1751
		return s.NegotiatedProtocol != ""
//line /usr/local/go/src/net/http/transport.go:1751
		// _ = "end of CoverTab[44659]"
//line /usr/local/go/src/net/http/transport.go:1751
	}() {
//line /usr/local/go/src/net/http/transport.go:1751
		_go_fuzz_dep_.CoverTab[44660]++
								if next, ok := t.TLSNextProto[s.NegotiatedProtocol]; ok {
//line /usr/local/go/src/net/http/transport.go:1752
			_go_fuzz_dep_.CoverTab[44661]++
									alt := next(cm.targetAddr, pconn.conn.(*tls.Conn))
									if e, ok := alt.(erringRoundTripper); ok {
//line /usr/local/go/src/net/http/transport.go:1754
				_go_fuzz_dep_.CoverTab[44663]++

										return nil, e.RoundTripErr()
//line /usr/local/go/src/net/http/transport.go:1756
				// _ = "end of CoverTab[44663]"
			} else {
//line /usr/local/go/src/net/http/transport.go:1757
				_go_fuzz_dep_.CoverTab[44664]++
//line /usr/local/go/src/net/http/transport.go:1757
				// _ = "end of CoverTab[44664]"
//line /usr/local/go/src/net/http/transport.go:1757
			}
//line /usr/local/go/src/net/http/transport.go:1757
			// _ = "end of CoverTab[44661]"
//line /usr/local/go/src/net/http/transport.go:1757
			_go_fuzz_dep_.CoverTab[44662]++
									return &persistConn{t: t, cacheKey: pconn.cacheKey, alt: alt}, nil
//line /usr/local/go/src/net/http/transport.go:1758
			// _ = "end of CoverTab[44662]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1759
			_go_fuzz_dep_.CoverTab[44665]++
//line /usr/local/go/src/net/http/transport.go:1759
			// _ = "end of CoverTab[44665]"
//line /usr/local/go/src/net/http/transport.go:1759
		}
//line /usr/local/go/src/net/http/transport.go:1759
		// _ = "end of CoverTab[44660]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1760
		_go_fuzz_dep_.CoverTab[44666]++
//line /usr/local/go/src/net/http/transport.go:1760
		// _ = "end of CoverTab[44666]"
//line /usr/local/go/src/net/http/transport.go:1760
	}
//line /usr/local/go/src/net/http/transport.go:1760
	// _ = "end of CoverTab[44566]"
//line /usr/local/go/src/net/http/transport.go:1760
	_go_fuzz_dep_.CoverTab[44567]++

							pconn.br = bufio.NewReaderSize(pconn, t.readBufferSize())
							pconn.bw = bufio.NewWriterSize(persistConnWriter{pconn}, t.writeBufferSize())
//line /usr/local/go/src/net/http/transport.go:1763
	_curRoutineNum42_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/transport.go:1763
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum42_)

							go pconn.readLoop()
//line /usr/local/go/src/net/http/transport.go:1765
	_curRoutineNum43_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/transport.go:1765
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum43_)
							go pconn.writeLoop()
							return pconn, nil
//line /usr/local/go/src/net/http/transport.go:1767
	// _ = "end of CoverTab[44567]"
}

// persistConnWriter is the io.Writer written to by pc.bw.
//line /usr/local/go/src/net/http/transport.go:1770
// It accumulates the number of bytes written to the underlying conn,
//line /usr/local/go/src/net/http/transport.go:1770
// so the retry logic can determine whether any bytes made it across
//line /usr/local/go/src/net/http/transport.go:1770
// the wire.
//line /usr/local/go/src/net/http/transport.go:1770
// This is exactly 1 pointer field wide so it can go into an interface
//line /usr/local/go/src/net/http/transport.go:1770
// without allocation.
//line /usr/local/go/src/net/http/transport.go:1776
type persistConnWriter struct {
	pc *persistConn
}

func (w persistConnWriter) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transport.go:1780
	_go_fuzz_dep_.CoverTab[44667]++
							n, err = w.pc.conn.Write(p)
							w.pc.nwrite += int64(n)
							return
//line /usr/local/go/src/net/http/transport.go:1783
	// _ = "end of CoverTab[44667]"
}

// ReadFrom exposes persistConnWriter's underlying Conn to io.Copy and if
//line /usr/local/go/src/net/http/transport.go:1786
// the Conn implements io.ReaderFrom, it can take advantage of optimizations
//line /usr/local/go/src/net/http/transport.go:1786
// such as sendfile.
//line /usr/local/go/src/net/http/transport.go:1789
func (w persistConnWriter) ReadFrom(r io.Reader) (n int64, err error) {
//line /usr/local/go/src/net/http/transport.go:1789
	_go_fuzz_dep_.CoverTab[44668]++
							n, err = io.Copy(w.pc.conn, r)
							w.pc.nwrite += n
							return
//line /usr/local/go/src/net/http/transport.go:1792
	// _ = "end of CoverTab[44668]"
}

var _ io.ReaderFrom = (*persistConnWriter)(nil)

// connectMethod is the map key (in its String form) for keeping persistent
//line /usr/local/go/src/net/http/transport.go:1797
// TCP connections alive for subsequent HTTP requests.
//line /usr/local/go/src/net/http/transport.go:1797
//
//line /usr/local/go/src/net/http/transport.go:1797
// A connect method may be of the following types:
//line /usr/local/go/src/net/http/transport.go:1797
//
//line /usr/local/go/src/net/http/transport.go:1797
//	connectMethod.key().String()      Description
//line /usr/local/go/src/net/http/transport.go:1797
//	------------------------------    -------------------------
//line /usr/local/go/src/net/http/transport.go:1797
//	|http|foo.com                     http directly to server, no proxy
//line /usr/local/go/src/net/http/transport.go:1797
//	|https|foo.com                    https directly to server, no proxy
//line /usr/local/go/src/net/http/transport.go:1797
//	|https,h1|foo.com                 https directly to server w/o HTTP/2, no proxy
//line /usr/local/go/src/net/http/transport.go:1797
//	http://proxy.com|https|foo.com    http to proxy, then CONNECT to foo.com
//line /usr/local/go/src/net/http/transport.go:1797
//	http://proxy.com|http             http to proxy, http to anywhere after that
//line /usr/local/go/src/net/http/transport.go:1797
//	socks5://proxy.com|http|foo.com   socks5 to proxy, then http to foo.com
//line /usr/local/go/src/net/http/transport.go:1797
//	socks5://proxy.com|https|foo.com  socks5 to proxy, then https to foo.com
//line /usr/local/go/src/net/http/transport.go:1797
//	https://proxy.com|https|foo.com   https to proxy, then CONNECT to foo.com
//line /usr/local/go/src/net/http/transport.go:1797
//	https://proxy.com|http            https to proxy, http to anywhere after that
//line /usr/local/go/src/net/http/transport.go:1813
type connectMethod struct {
	_		incomparable
	proxyURL	*url.URL	// nil for no proxy, else full proxy URL
	targetScheme	string		// "http" or "https"
	// If proxyURL specifies an http or https proxy, and targetScheme is http (not https),
	// then targetAddr is not included in the connect method key, because the socket can
	// be reused for different targetAddr values.
	targetAddr	string
	onlyH1		bool	// whether to disable HTTP/2 and force HTTP/1
}

func (cm *connectMethod) key() connectMethodKey {
//line /usr/local/go/src/net/http/transport.go:1824
	_go_fuzz_dep_.CoverTab[44669]++
							proxyStr := ""
							targetAddr := cm.targetAddr
							if cm.proxyURL != nil {
//line /usr/local/go/src/net/http/transport.go:1827
		_go_fuzz_dep_.CoverTab[44671]++
								proxyStr = cm.proxyURL.String()
								if (cm.proxyURL.Scheme == "http" || func() bool {
//line /usr/local/go/src/net/http/transport.go:1829
			_go_fuzz_dep_.CoverTab[44672]++
//line /usr/local/go/src/net/http/transport.go:1829
			return cm.proxyURL.Scheme == "https"
//line /usr/local/go/src/net/http/transport.go:1829
			// _ = "end of CoverTab[44672]"
//line /usr/local/go/src/net/http/transport.go:1829
		}()) && func() bool {
//line /usr/local/go/src/net/http/transport.go:1829
			_go_fuzz_dep_.CoverTab[44673]++
//line /usr/local/go/src/net/http/transport.go:1829
			return cm.targetScheme == "http"
//line /usr/local/go/src/net/http/transport.go:1829
			// _ = "end of CoverTab[44673]"
//line /usr/local/go/src/net/http/transport.go:1829
		}() {
//line /usr/local/go/src/net/http/transport.go:1829
			_go_fuzz_dep_.CoverTab[44674]++
									targetAddr = ""
//line /usr/local/go/src/net/http/transport.go:1830
			// _ = "end of CoverTab[44674]"
		} else {
//line /usr/local/go/src/net/http/transport.go:1831
			_go_fuzz_dep_.CoverTab[44675]++
//line /usr/local/go/src/net/http/transport.go:1831
			// _ = "end of CoverTab[44675]"
//line /usr/local/go/src/net/http/transport.go:1831
		}
//line /usr/local/go/src/net/http/transport.go:1831
		// _ = "end of CoverTab[44671]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1832
		_go_fuzz_dep_.CoverTab[44676]++
//line /usr/local/go/src/net/http/transport.go:1832
		// _ = "end of CoverTab[44676]"
//line /usr/local/go/src/net/http/transport.go:1832
	}
//line /usr/local/go/src/net/http/transport.go:1832
	// _ = "end of CoverTab[44669]"
//line /usr/local/go/src/net/http/transport.go:1832
	_go_fuzz_dep_.CoverTab[44670]++
							return connectMethodKey{
		proxy:	proxyStr,
		scheme:	cm.targetScheme,
		addr:	targetAddr,
		onlyH1:	cm.onlyH1,
	}
//line /usr/local/go/src/net/http/transport.go:1838
	// _ = "end of CoverTab[44670]"
}

// scheme returns the first hop scheme: http, https, or socks5
func (cm *connectMethod) scheme() string {
//line /usr/local/go/src/net/http/transport.go:1842
	_go_fuzz_dep_.CoverTab[44677]++
							if cm.proxyURL != nil {
//line /usr/local/go/src/net/http/transport.go:1843
		_go_fuzz_dep_.CoverTab[44679]++
								return cm.proxyURL.Scheme
//line /usr/local/go/src/net/http/transport.go:1844
		// _ = "end of CoverTab[44679]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1845
		_go_fuzz_dep_.CoverTab[44680]++
//line /usr/local/go/src/net/http/transport.go:1845
		// _ = "end of CoverTab[44680]"
//line /usr/local/go/src/net/http/transport.go:1845
	}
//line /usr/local/go/src/net/http/transport.go:1845
	// _ = "end of CoverTab[44677]"
//line /usr/local/go/src/net/http/transport.go:1845
	_go_fuzz_dep_.CoverTab[44678]++
							return cm.targetScheme
//line /usr/local/go/src/net/http/transport.go:1846
	// _ = "end of CoverTab[44678]"
}

// addr returns the first hop "host:port" to which we need to TCP connect.
func (cm *connectMethod) addr() string {
//line /usr/local/go/src/net/http/transport.go:1850
	_go_fuzz_dep_.CoverTab[44681]++
							if cm.proxyURL != nil {
//line /usr/local/go/src/net/http/transport.go:1851
		_go_fuzz_dep_.CoverTab[44683]++
								return canonicalAddr(cm.proxyURL)
//line /usr/local/go/src/net/http/transport.go:1852
		// _ = "end of CoverTab[44683]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1853
		_go_fuzz_dep_.CoverTab[44684]++
//line /usr/local/go/src/net/http/transport.go:1853
		// _ = "end of CoverTab[44684]"
//line /usr/local/go/src/net/http/transport.go:1853
	}
//line /usr/local/go/src/net/http/transport.go:1853
	// _ = "end of CoverTab[44681]"
//line /usr/local/go/src/net/http/transport.go:1853
	_go_fuzz_dep_.CoverTab[44682]++
							return cm.targetAddr
//line /usr/local/go/src/net/http/transport.go:1854
	// _ = "end of CoverTab[44682]"
}

// tlsHost returns the host name to match against the peer's
//line /usr/local/go/src/net/http/transport.go:1857
// TLS certificate.
//line /usr/local/go/src/net/http/transport.go:1859
func (cm *connectMethod) tlsHost() string {
//line /usr/local/go/src/net/http/transport.go:1859
	_go_fuzz_dep_.CoverTab[44685]++
							h := cm.targetAddr
							if hasPort(h) {
//line /usr/local/go/src/net/http/transport.go:1861
		_go_fuzz_dep_.CoverTab[44687]++
								h = h[:strings.LastIndex(h, ":")]
//line /usr/local/go/src/net/http/transport.go:1862
		// _ = "end of CoverTab[44687]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1863
		_go_fuzz_dep_.CoverTab[44688]++
//line /usr/local/go/src/net/http/transport.go:1863
		// _ = "end of CoverTab[44688]"
//line /usr/local/go/src/net/http/transport.go:1863
	}
//line /usr/local/go/src/net/http/transport.go:1863
	// _ = "end of CoverTab[44685]"
//line /usr/local/go/src/net/http/transport.go:1863
	_go_fuzz_dep_.CoverTab[44686]++
							return h
//line /usr/local/go/src/net/http/transport.go:1864
	// _ = "end of CoverTab[44686]"
}

// connectMethodKey is the map key version of connectMethod, with a
//line /usr/local/go/src/net/http/transport.go:1867
// stringified proxy URL (or the empty string) instead of a pointer to
//line /usr/local/go/src/net/http/transport.go:1867
// a URL.
//line /usr/local/go/src/net/http/transport.go:1870
type connectMethodKey struct {
	proxy, scheme, addr	string
	onlyH1			bool
}

func (k connectMethodKey) String() string {
//line /usr/local/go/src/net/http/transport.go:1875
	_go_fuzz_dep_.CoverTab[44689]++
	// Only used by tests.
	var h1 string
	if k.onlyH1 {
//line /usr/local/go/src/net/http/transport.go:1878
		_go_fuzz_dep_.CoverTab[44691]++
								h1 = ",h1"
//line /usr/local/go/src/net/http/transport.go:1879
		// _ = "end of CoverTab[44691]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1880
		_go_fuzz_dep_.CoverTab[44692]++
//line /usr/local/go/src/net/http/transport.go:1880
		// _ = "end of CoverTab[44692]"
//line /usr/local/go/src/net/http/transport.go:1880
	}
//line /usr/local/go/src/net/http/transport.go:1880
	// _ = "end of CoverTab[44689]"
//line /usr/local/go/src/net/http/transport.go:1880
	_go_fuzz_dep_.CoverTab[44690]++
							return fmt.Sprintf("%s|%s%s|%s", k.proxy, k.scheme, h1, k.addr)
//line /usr/local/go/src/net/http/transport.go:1881
	// _ = "end of CoverTab[44690]"
}

// persistConn wraps a connection, usually a persistent one
//line /usr/local/go/src/net/http/transport.go:1884
// (but may be used for non-keep-alive requests as well)
//line /usr/local/go/src/net/http/transport.go:1886
type persistConn struct {
	// alt optionally specifies the TLS NextProto RoundTripper.
	// This is used for HTTP/2 today and future protocols later.
	// If it's non-nil, the rest of the fields are unused.
	alt	RoundTripper

	t		*Transport
	cacheKey	connectMethodKey
	conn		net.Conn
	tlsState	*tls.ConnectionState
	br		*bufio.Reader		// from conn
	bw		*bufio.Writer		// to conn
	nwrite		int64			// bytes written
	reqch		chan requestAndChan	// written by roundTrip; read by readLoop
	writech		chan writeRequest	// written by roundTrip; read by writeLoop
	closech		chan struct{}		// closed when conn closed
	isProxy		bool
	sawEOF		bool	// whether we've seen EOF from conn; owned by readLoop
	readLimit	int64	// bytes allowed to be read; owned by readLoop
	// writeErrCh passes the request write error (usually nil)
	// from the writeLoop goroutine to the readLoop which passes
	// it off to the res.Body reader, which then uses it to decide
	// whether or not a connection can be reused. Issue 7569.
	writeErrCh	chan error

	writeLoopDone	chan struct{}	// closed when write loop ends

	// Both guarded by Transport.idleMu:
	idleAt		time.Time	// time it last become idle
	idleTimer	*time.Timer	// holding an AfterFunc to close it

	mu			sync.Mutex	// guards following fields
	numExpectedResponses	int
	closed			error	// set non-nil when conn is closed, before closech is closed
	canceledErr		error	// set non-nil if conn is canceled
	broken			bool	// an error has happened on this connection; marked broken so it's not reused.
	reused			bool	// whether conn has had successful request/response and is being reused.
	// mutateHeaderFunc is an optional func to modify extra
	// headers on each outbound request before it's written. (the
	// original Request given to RoundTrip is not modified)
	mutateHeaderFunc	func(Header)
}

func (pc *persistConn) maxHeaderResponseSize() int64 {
//line /usr/local/go/src/net/http/transport.go:1929
	_go_fuzz_dep_.CoverTab[44693]++
							if v := pc.t.MaxResponseHeaderBytes; v != 0 {
//line /usr/local/go/src/net/http/transport.go:1930
		_go_fuzz_dep_.CoverTab[44695]++
								return v
//line /usr/local/go/src/net/http/transport.go:1931
		// _ = "end of CoverTab[44695]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1932
		_go_fuzz_dep_.CoverTab[44696]++
//line /usr/local/go/src/net/http/transport.go:1932
		// _ = "end of CoverTab[44696]"
//line /usr/local/go/src/net/http/transport.go:1932
	}
//line /usr/local/go/src/net/http/transport.go:1932
	// _ = "end of CoverTab[44693]"
//line /usr/local/go/src/net/http/transport.go:1932
	_go_fuzz_dep_.CoverTab[44694]++
							return 10 << 20
//line /usr/local/go/src/net/http/transport.go:1933
	// _ = "end of CoverTab[44694]"
}

func (pc *persistConn) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transport.go:1936
	_go_fuzz_dep_.CoverTab[44697]++
							if pc.readLimit <= 0 {
//line /usr/local/go/src/net/http/transport.go:1937
		_go_fuzz_dep_.CoverTab[44701]++
								return 0, fmt.Errorf("read limit of %d bytes exhausted", pc.maxHeaderResponseSize())
//line /usr/local/go/src/net/http/transport.go:1938
		// _ = "end of CoverTab[44701]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1939
		_go_fuzz_dep_.CoverTab[44702]++
//line /usr/local/go/src/net/http/transport.go:1939
		// _ = "end of CoverTab[44702]"
//line /usr/local/go/src/net/http/transport.go:1939
	}
//line /usr/local/go/src/net/http/transport.go:1939
	// _ = "end of CoverTab[44697]"
//line /usr/local/go/src/net/http/transport.go:1939
	_go_fuzz_dep_.CoverTab[44698]++
							if int64(len(p)) > pc.readLimit {
//line /usr/local/go/src/net/http/transport.go:1940
		_go_fuzz_dep_.CoverTab[44703]++
								p = p[:pc.readLimit]
//line /usr/local/go/src/net/http/transport.go:1941
		// _ = "end of CoverTab[44703]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1942
		_go_fuzz_dep_.CoverTab[44704]++
//line /usr/local/go/src/net/http/transport.go:1942
		// _ = "end of CoverTab[44704]"
//line /usr/local/go/src/net/http/transport.go:1942
	}
//line /usr/local/go/src/net/http/transport.go:1942
	// _ = "end of CoverTab[44698]"
//line /usr/local/go/src/net/http/transport.go:1942
	_go_fuzz_dep_.CoverTab[44699]++
							n, err = pc.conn.Read(p)
							if err == io.EOF {
//line /usr/local/go/src/net/http/transport.go:1944
		_go_fuzz_dep_.CoverTab[44705]++
								pc.sawEOF = true
//line /usr/local/go/src/net/http/transport.go:1945
		// _ = "end of CoverTab[44705]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1946
		_go_fuzz_dep_.CoverTab[44706]++
//line /usr/local/go/src/net/http/transport.go:1946
		// _ = "end of CoverTab[44706]"
//line /usr/local/go/src/net/http/transport.go:1946
	}
//line /usr/local/go/src/net/http/transport.go:1946
	// _ = "end of CoverTab[44699]"
//line /usr/local/go/src/net/http/transport.go:1946
	_go_fuzz_dep_.CoverTab[44700]++
							pc.readLimit -= int64(n)
							return
//line /usr/local/go/src/net/http/transport.go:1948
	// _ = "end of CoverTab[44700]"
}

// isBroken reports whether this connection is in a known broken state.
func (pc *persistConn) isBroken() bool {
//line /usr/local/go/src/net/http/transport.go:1952
	_go_fuzz_dep_.CoverTab[44707]++
							pc.mu.Lock()
							b := pc.closed != nil
							pc.mu.Unlock()
							return b
//line /usr/local/go/src/net/http/transport.go:1956
	// _ = "end of CoverTab[44707]"
}

// canceled returns non-nil if the connection was closed due to
//line /usr/local/go/src/net/http/transport.go:1959
// CancelRequest or due to context cancellation.
//line /usr/local/go/src/net/http/transport.go:1961
func (pc *persistConn) canceled() error {
//line /usr/local/go/src/net/http/transport.go:1961
	_go_fuzz_dep_.CoverTab[44708]++
							pc.mu.Lock()
							defer pc.mu.Unlock()
							return pc.canceledErr
//line /usr/local/go/src/net/http/transport.go:1964
	// _ = "end of CoverTab[44708]"
}

// isReused reports whether this connection has been used before.
func (pc *persistConn) isReused() bool {
//line /usr/local/go/src/net/http/transport.go:1968
	_go_fuzz_dep_.CoverTab[44709]++
							pc.mu.Lock()
							r := pc.reused
							pc.mu.Unlock()
							return r
//line /usr/local/go/src/net/http/transport.go:1972
	// _ = "end of CoverTab[44709]"
}

func (pc *persistConn) gotIdleConnTrace(idleAt time.Time) (t httptrace.GotConnInfo) {
//line /usr/local/go/src/net/http/transport.go:1975
	_go_fuzz_dep_.CoverTab[44710]++
							pc.mu.Lock()
							defer pc.mu.Unlock()
							t.Reused = pc.reused
							t.Conn = pc.conn
							t.WasIdle = true
							if !idleAt.IsZero() {
//line /usr/local/go/src/net/http/transport.go:1981
		_go_fuzz_dep_.CoverTab[44712]++
								t.IdleTime = time.Since(idleAt)
//line /usr/local/go/src/net/http/transport.go:1982
		// _ = "end of CoverTab[44712]"
	} else {
//line /usr/local/go/src/net/http/transport.go:1983
		_go_fuzz_dep_.CoverTab[44713]++
//line /usr/local/go/src/net/http/transport.go:1983
		// _ = "end of CoverTab[44713]"
//line /usr/local/go/src/net/http/transport.go:1983
	}
//line /usr/local/go/src/net/http/transport.go:1983
	// _ = "end of CoverTab[44710]"
//line /usr/local/go/src/net/http/transport.go:1983
	_go_fuzz_dep_.CoverTab[44711]++
							return
//line /usr/local/go/src/net/http/transport.go:1984
	// _ = "end of CoverTab[44711]"
}

func (pc *persistConn) cancelRequest(err error) {
//line /usr/local/go/src/net/http/transport.go:1987
	_go_fuzz_dep_.CoverTab[44714]++
							pc.mu.Lock()
							defer pc.mu.Unlock()
							pc.canceledErr = err
							pc.closeLocked(errRequestCanceled)
//line /usr/local/go/src/net/http/transport.go:1991
	// _ = "end of CoverTab[44714]"
}

// closeConnIfStillIdle closes the connection if it's still sitting idle.
//line /usr/local/go/src/net/http/transport.go:1994
// This is what's called by the persistConn's idleTimer, and is run in its
//line /usr/local/go/src/net/http/transport.go:1994
// own goroutine.
//line /usr/local/go/src/net/http/transport.go:1997
func (pc *persistConn) closeConnIfStillIdle() {
//line /usr/local/go/src/net/http/transport.go:1997
	_go_fuzz_dep_.CoverTab[44715]++
							t := pc.t
							t.idleMu.Lock()
							defer t.idleMu.Unlock()
							if _, ok := t.idleLRU.m[pc]; !ok {
//line /usr/local/go/src/net/http/transport.go:2001
		_go_fuzz_dep_.CoverTab[44717]++

								return
//line /usr/local/go/src/net/http/transport.go:2003
		// _ = "end of CoverTab[44717]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2004
		_go_fuzz_dep_.CoverTab[44718]++
//line /usr/local/go/src/net/http/transport.go:2004
		// _ = "end of CoverTab[44718]"
//line /usr/local/go/src/net/http/transport.go:2004
	}
//line /usr/local/go/src/net/http/transport.go:2004
	// _ = "end of CoverTab[44715]"
//line /usr/local/go/src/net/http/transport.go:2004
	_go_fuzz_dep_.CoverTab[44716]++
							t.removeIdleConnLocked(pc)
							pc.close(errIdleConnTimeout)
//line /usr/local/go/src/net/http/transport.go:2006
	// _ = "end of CoverTab[44716]"
}

// mapRoundTripError returns the appropriate error value for
//line /usr/local/go/src/net/http/transport.go:2009
// persistConn.roundTrip.
//line /usr/local/go/src/net/http/transport.go:2009
//
//line /usr/local/go/src/net/http/transport.go:2009
// The provided err is the first error that (*persistConn).roundTrip
//line /usr/local/go/src/net/http/transport.go:2009
// happened to receive from its select statement.
//line /usr/local/go/src/net/http/transport.go:2009
//
//line /usr/local/go/src/net/http/transport.go:2009
// The startBytesWritten value should be the value of pc.nwrite before the roundTrip
//line /usr/local/go/src/net/http/transport.go:2009
// started writing the request.
//line /usr/local/go/src/net/http/transport.go:2017
func (pc *persistConn) mapRoundTripError(req *transportRequest, startBytesWritten int64, err error) error {
//line /usr/local/go/src/net/http/transport.go:2017
	_go_fuzz_dep_.CoverTab[44719]++
							if err == nil {
//line /usr/local/go/src/net/http/transport.go:2018
		_go_fuzz_dep_.CoverTab[44726]++
								return nil
//line /usr/local/go/src/net/http/transport.go:2019
		// _ = "end of CoverTab[44726]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2020
		_go_fuzz_dep_.CoverTab[44727]++
//line /usr/local/go/src/net/http/transport.go:2020
		// _ = "end of CoverTab[44727]"
//line /usr/local/go/src/net/http/transport.go:2020
	}
//line /usr/local/go/src/net/http/transport.go:2020
	// _ = "end of CoverTab[44719]"
//line /usr/local/go/src/net/http/transport.go:2020
	_go_fuzz_dep_.CoverTab[44720]++

//line /usr/local/go/src/net/http/transport.go:2029
	<-pc.writeLoopDone

//line /usr/local/go/src/net/http/transport.go:2034
	if cerr := pc.canceled(); cerr != nil {
//line /usr/local/go/src/net/http/transport.go:2034
		_go_fuzz_dep_.CoverTab[44728]++
								return cerr
//line /usr/local/go/src/net/http/transport.go:2035
		// _ = "end of CoverTab[44728]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2036
		_go_fuzz_dep_.CoverTab[44729]++
//line /usr/local/go/src/net/http/transport.go:2036
		// _ = "end of CoverTab[44729]"
//line /usr/local/go/src/net/http/transport.go:2036
	}
//line /usr/local/go/src/net/http/transport.go:2036
	// _ = "end of CoverTab[44720]"
//line /usr/local/go/src/net/http/transport.go:2036
	_go_fuzz_dep_.CoverTab[44721]++

//line /usr/local/go/src/net/http/transport.go:2039
	req.mu.Lock()
	reqErr := req.err
	req.mu.Unlock()
	if reqErr != nil {
//line /usr/local/go/src/net/http/transport.go:2042
		_go_fuzz_dep_.CoverTab[44730]++
								return reqErr
//line /usr/local/go/src/net/http/transport.go:2043
		// _ = "end of CoverTab[44730]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2044
		_go_fuzz_dep_.CoverTab[44731]++
//line /usr/local/go/src/net/http/transport.go:2044
		// _ = "end of CoverTab[44731]"
//line /usr/local/go/src/net/http/transport.go:2044
	}
//line /usr/local/go/src/net/http/transport.go:2044
	// _ = "end of CoverTab[44721]"
//line /usr/local/go/src/net/http/transport.go:2044
	_go_fuzz_dep_.CoverTab[44722]++

							if err == errServerClosedIdle {
//line /usr/local/go/src/net/http/transport.go:2046
		_go_fuzz_dep_.CoverTab[44732]++

								return err
//line /usr/local/go/src/net/http/transport.go:2048
		// _ = "end of CoverTab[44732]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2049
		_go_fuzz_dep_.CoverTab[44733]++
//line /usr/local/go/src/net/http/transport.go:2049
		// _ = "end of CoverTab[44733]"
//line /usr/local/go/src/net/http/transport.go:2049
	}
//line /usr/local/go/src/net/http/transport.go:2049
	// _ = "end of CoverTab[44722]"
//line /usr/local/go/src/net/http/transport.go:2049
	_go_fuzz_dep_.CoverTab[44723]++

							if _, ok := err.(transportReadFromServerError); ok {
//line /usr/local/go/src/net/http/transport.go:2051
		_go_fuzz_dep_.CoverTab[44734]++
								if pc.nwrite == startBytesWritten {
//line /usr/local/go/src/net/http/transport.go:2052
			_go_fuzz_dep_.CoverTab[44736]++
									return nothingWrittenError{err}
//line /usr/local/go/src/net/http/transport.go:2053
			// _ = "end of CoverTab[44736]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2054
			_go_fuzz_dep_.CoverTab[44737]++
//line /usr/local/go/src/net/http/transport.go:2054
			// _ = "end of CoverTab[44737]"
//line /usr/local/go/src/net/http/transport.go:2054
		}
//line /usr/local/go/src/net/http/transport.go:2054
		// _ = "end of CoverTab[44734]"
//line /usr/local/go/src/net/http/transport.go:2054
		_go_fuzz_dep_.CoverTab[44735]++

								return err
//line /usr/local/go/src/net/http/transport.go:2056
		// _ = "end of CoverTab[44735]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2057
		_go_fuzz_dep_.CoverTab[44738]++
//line /usr/local/go/src/net/http/transport.go:2057
		// _ = "end of CoverTab[44738]"
//line /usr/local/go/src/net/http/transport.go:2057
	}
//line /usr/local/go/src/net/http/transport.go:2057
	// _ = "end of CoverTab[44723]"
//line /usr/local/go/src/net/http/transport.go:2057
	_go_fuzz_dep_.CoverTab[44724]++
							if pc.isBroken() {
//line /usr/local/go/src/net/http/transport.go:2058
		_go_fuzz_dep_.CoverTab[44739]++
								if pc.nwrite == startBytesWritten {
//line /usr/local/go/src/net/http/transport.go:2059
			_go_fuzz_dep_.CoverTab[44741]++
									return nothingWrittenError{err}
//line /usr/local/go/src/net/http/transport.go:2060
			// _ = "end of CoverTab[44741]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2061
			_go_fuzz_dep_.CoverTab[44742]++
//line /usr/local/go/src/net/http/transport.go:2061
			// _ = "end of CoverTab[44742]"
//line /usr/local/go/src/net/http/transport.go:2061
		}
//line /usr/local/go/src/net/http/transport.go:2061
		// _ = "end of CoverTab[44739]"
//line /usr/local/go/src/net/http/transport.go:2061
		_go_fuzz_dep_.CoverTab[44740]++
								return fmt.Errorf("net/http: HTTP/1.x transport connection broken: %w", err)
//line /usr/local/go/src/net/http/transport.go:2062
		// _ = "end of CoverTab[44740]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2063
		_go_fuzz_dep_.CoverTab[44743]++
//line /usr/local/go/src/net/http/transport.go:2063
		// _ = "end of CoverTab[44743]"
//line /usr/local/go/src/net/http/transport.go:2063
	}
//line /usr/local/go/src/net/http/transport.go:2063
	// _ = "end of CoverTab[44724]"
//line /usr/local/go/src/net/http/transport.go:2063
	_go_fuzz_dep_.CoverTab[44725]++
							return err
//line /usr/local/go/src/net/http/transport.go:2064
	// _ = "end of CoverTab[44725]"
}

// errCallerOwnsConn is an internal sentinel error used when we hand
//line /usr/local/go/src/net/http/transport.go:2067
// off a writable response.Body to the caller. We use this to prevent
//line /usr/local/go/src/net/http/transport.go:2067
// closing a net.Conn that is now owned by the caller.
//line /usr/local/go/src/net/http/transport.go:2070
var errCallerOwnsConn = errors.New("read loop ending; caller owns writable underlying conn")

func (pc *persistConn) readLoop() {
//line /usr/local/go/src/net/http/transport.go:2072
	_go_fuzz_dep_.CoverTab[44744]++
							closeErr := errReadLoopExiting
							defer func() {
//line /usr/local/go/src/net/http/transport.go:2074
		_go_fuzz_dep_.CoverTab[44747]++
								pc.close(closeErr)
								pc.t.removeIdleConn(pc)
//line /usr/local/go/src/net/http/transport.go:2076
		// _ = "end of CoverTab[44747]"
	}()
//line /usr/local/go/src/net/http/transport.go:2077
	// _ = "end of CoverTab[44744]"
//line /usr/local/go/src/net/http/transport.go:2077
	_go_fuzz_dep_.CoverTab[44745]++

							tryPutIdleConn := func(trace *httptrace.ClientTrace) bool {
//line /usr/local/go/src/net/http/transport.go:2079
		_go_fuzz_dep_.CoverTab[44748]++
								if err := pc.t.tryPutIdleConn(pc); err != nil {
//line /usr/local/go/src/net/http/transport.go:2080
			_go_fuzz_dep_.CoverTab[44751]++
									closeErr = err
									if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:2082
				_go_fuzz_dep_.CoverTab[44753]++
//line /usr/local/go/src/net/http/transport.go:2082
				return trace.PutIdleConn != nil
//line /usr/local/go/src/net/http/transport.go:2082
				// _ = "end of CoverTab[44753]"
//line /usr/local/go/src/net/http/transport.go:2082
			}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2082
				_go_fuzz_dep_.CoverTab[44754]++
//line /usr/local/go/src/net/http/transport.go:2082
				return err != errKeepAlivesDisabled
//line /usr/local/go/src/net/http/transport.go:2082
				// _ = "end of CoverTab[44754]"
//line /usr/local/go/src/net/http/transport.go:2082
			}() {
//line /usr/local/go/src/net/http/transport.go:2082
				_go_fuzz_dep_.CoverTab[44755]++
										trace.PutIdleConn(err)
//line /usr/local/go/src/net/http/transport.go:2083
				// _ = "end of CoverTab[44755]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2084
				_go_fuzz_dep_.CoverTab[44756]++
//line /usr/local/go/src/net/http/transport.go:2084
				// _ = "end of CoverTab[44756]"
//line /usr/local/go/src/net/http/transport.go:2084
			}
//line /usr/local/go/src/net/http/transport.go:2084
			// _ = "end of CoverTab[44751]"
//line /usr/local/go/src/net/http/transport.go:2084
			_go_fuzz_dep_.CoverTab[44752]++
									return false
//line /usr/local/go/src/net/http/transport.go:2085
			// _ = "end of CoverTab[44752]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2086
			_go_fuzz_dep_.CoverTab[44757]++
//line /usr/local/go/src/net/http/transport.go:2086
			// _ = "end of CoverTab[44757]"
//line /usr/local/go/src/net/http/transport.go:2086
		}
//line /usr/local/go/src/net/http/transport.go:2086
		// _ = "end of CoverTab[44748]"
//line /usr/local/go/src/net/http/transport.go:2086
		_go_fuzz_dep_.CoverTab[44749]++
								if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:2087
			_go_fuzz_dep_.CoverTab[44758]++
//line /usr/local/go/src/net/http/transport.go:2087
			return trace.PutIdleConn != nil
//line /usr/local/go/src/net/http/transport.go:2087
			// _ = "end of CoverTab[44758]"
//line /usr/local/go/src/net/http/transport.go:2087
		}() {
//line /usr/local/go/src/net/http/transport.go:2087
			_go_fuzz_dep_.CoverTab[44759]++
									trace.PutIdleConn(nil)
//line /usr/local/go/src/net/http/transport.go:2088
			// _ = "end of CoverTab[44759]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2089
			_go_fuzz_dep_.CoverTab[44760]++
//line /usr/local/go/src/net/http/transport.go:2089
			// _ = "end of CoverTab[44760]"
//line /usr/local/go/src/net/http/transport.go:2089
		}
//line /usr/local/go/src/net/http/transport.go:2089
		// _ = "end of CoverTab[44749]"
//line /usr/local/go/src/net/http/transport.go:2089
		_go_fuzz_dep_.CoverTab[44750]++
								return true
//line /usr/local/go/src/net/http/transport.go:2090
		// _ = "end of CoverTab[44750]"
	}
//line /usr/local/go/src/net/http/transport.go:2091
	// _ = "end of CoverTab[44745]"
//line /usr/local/go/src/net/http/transport.go:2091
	_go_fuzz_dep_.CoverTab[44746]++

//line /usr/local/go/src/net/http/transport.go:2096
	eofc := make(chan struct{})
							defer close(eofc)

//line /usr/local/go/src/net/http/transport.go:2100
	testHookMu.Lock()
	testHookReadLoopBeforeNextRead := testHookReadLoopBeforeNextRead
	testHookMu.Unlock()

	alive := true
	for alive {
//line /usr/local/go/src/net/http/transport.go:2105
		_go_fuzz_dep_.CoverTab[44761]++
								pc.readLimit = pc.maxHeaderResponseSize()
								_, err := pc.br.Peek(1)

								pc.mu.Lock()
								if pc.numExpectedResponses == 0 {
//line /usr/local/go/src/net/http/transport.go:2110
			_go_fuzz_dep_.CoverTab[44771]++
									pc.readLoopPeekFailLocked(err)
									pc.mu.Unlock()
									return
//line /usr/local/go/src/net/http/transport.go:2113
			// _ = "end of CoverTab[44771]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2114
			_go_fuzz_dep_.CoverTab[44772]++
//line /usr/local/go/src/net/http/transport.go:2114
			// _ = "end of CoverTab[44772]"
//line /usr/local/go/src/net/http/transport.go:2114
		}
//line /usr/local/go/src/net/http/transport.go:2114
		// _ = "end of CoverTab[44761]"
//line /usr/local/go/src/net/http/transport.go:2114
		_go_fuzz_dep_.CoverTab[44762]++
								pc.mu.Unlock()

								rc := <-pc.reqch
								trace := httptrace.ContextClientTrace(rc.req.Context())

								var resp *Response
								if err == nil {
//line /usr/local/go/src/net/http/transport.go:2121
			_go_fuzz_dep_.CoverTab[44773]++
									resp, err = pc.readResponse(rc, trace)
//line /usr/local/go/src/net/http/transport.go:2122
			// _ = "end of CoverTab[44773]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2123
			_go_fuzz_dep_.CoverTab[44774]++
									err = transportReadFromServerError{err}
									closeErr = err
//line /usr/local/go/src/net/http/transport.go:2125
			// _ = "end of CoverTab[44774]"
		}
//line /usr/local/go/src/net/http/transport.go:2126
		// _ = "end of CoverTab[44762]"
//line /usr/local/go/src/net/http/transport.go:2126
		_go_fuzz_dep_.CoverTab[44763]++

								if err != nil {
//line /usr/local/go/src/net/http/transport.go:2128
			_go_fuzz_dep_.CoverTab[44775]++
									if pc.readLimit <= 0 {
//line /usr/local/go/src/net/http/transport.go:2129
				_go_fuzz_dep_.CoverTab[44778]++
										err = fmt.Errorf("net/http: server response headers exceeded %d bytes; aborted", pc.maxHeaderResponseSize())
//line /usr/local/go/src/net/http/transport.go:2130
				// _ = "end of CoverTab[44778]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2131
				_go_fuzz_dep_.CoverTab[44779]++
//line /usr/local/go/src/net/http/transport.go:2131
				// _ = "end of CoverTab[44779]"
//line /usr/local/go/src/net/http/transport.go:2131
			}
//line /usr/local/go/src/net/http/transport.go:2131
			// _ = "end of CoverTab[44775]"
//line /usr/local/go/src/net/http/transport.go:2131
			_go_fuzz_dep_.CoverTab[44776]++

									select {
			case rc.ch <- responseAndError{err: err}:
//line /usr/local/go/src/net/http/transport.go:2134
				_go_fuzz_dep_.CoverTab[44780]++
//line /usr/local/go/src/net/http/transport.go:2134
				// _ = "end of CoverTab[44780]"
			case <-rc.callerGone:
//line /usr/local/go/src/net/http/transport.go:2135
				_go_fuzz_dep_.CoverTab[44781]++
										return
//line /usr/local/go/src/net/http/transport.go:2136
				// _ = "end of CoverTab[44781]"
			}
//line /usr/local/go/src/net/http/transport.go:2137
			// _ = "end of CoverTab[44776]"
//line /usr/local/go/src/net/http/transport.go:2137
			_go_fuzz_dep_.CoverTab[44777]++
									return
//line /usr/local/go/src/net/http/transport.go:2138
			// _ = "end of CoverTab[44777]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2139
			_go_fuzz_dep_.CoverTab[44782]++
//line /usr/local/go/src/net/http/transport.go:2139
			// _ = "end of CoverTab[44782]"
//line /usr/local/go/src/net/http/transport.go:2139
		}
//line /usr/local/go/src/net/http/transport.go:2139
		// _ = "end of CoverTab[44763]"
//line /usr/local/go/src/net/http/transport.go:2139
		_go_fuzz_dep_.CoverTab[44764]++
								pc.readLimit = maxInt64

								pc.mu.Lock()
								pc.numExpectedResponses--
								pc.mu.Unlock()

								bodyWritable := resp.bodyIsWritable()
								hasBody := rc.req.Method != "HEAD" && func() bool {
//line /usr/local/go/src/net/http/transport.go:2147
			_go_fuzz_dep_.CoverTab[44783]++
//line /usr/local/go/src/net/http/transport.go:2147
			return resp.ContentLength != 0
//line /usr/local/go/src/net/http/transport.go:2147
			// _ = "end of CoverTab[44783]"
//line /usr/local/go/src/net/http/transport.go:2147
		}()

								if resp.Close || func() bool {
//line /usr/local/go/src/net/http/transport.go:2149
			_go_fuzz_dep_.CoverTab[44784]++
//line /usr/local/go/src/net/http/transport.go:2149
			return rc.req.Close
//line /usr/local/go/src/net/http/transport.go:2149
			// _ = "end of CoverTab[44784]"
//line /usr/local/go/src/net/http/transport.go:2149
		}() || func() bool {
//line /usr/local/go/src/net/http/transport.go:2149
			_go_fuzz_dep_.CoverTab[44785]++
//line /usr/local/go/src/net/http/transport.go:2149
			return resp.StatusCode <= 199
//line /usr/local/go/src/net/http/transport.go:2149
			// _ = "end of CoverTab[44785]"
//line /usr/local/go/src/net/http/transport.go:2149
		}() || func() bool {
//line /usr/local/go/src/net/http/transport.go:2149
			_go_fuzz_dep_.CoverTab[44786]++
//line /usr/local/go/src/net/http/transport.go:2149
			return bodyWritable
//line /usr/local/go/src/net/http/transport.go:2149
			// _ = "end of CoverTab[44786]"
//line /usr/local/go/src/net/http/transport.go:2149
		}() {
//line /usr/local/go/src/net/http/transport.go:2149
			_go_fuzz_dep_.CoverTab[44787]++

//line /usr/local/go/src/net/http/transport.go:2153
			alive = false
//line /usr/local/go/src/net/http/transport.go:2153
			// _ = "end of CoverTab[44787]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2154
			_go_fuzz_dep_.CoverTab[44788]++
//line /usr/local/go/src/net/http/transport.go:2154
			// _ = "end of CoverTab[44788]"
//line /usr/local/go/src/net/http/transport.go:2154
		}
//line /usr/local/go/src/net/http/transport.go:2154
		// _ = "end of CoverTab[44764]"
//line /usr/local/go/src/net/http/transport.go:2154
		_go_fuzz_dep_.CoverTab[44765]++

								if !hasBody || func() bool {
//line /usr/local/go/src/net/http/transport.go:2156
			_go_fuzz_dep_.CoverTab[44789]++
//line /usr/local/go/src/net/http/transport.go:2156
			return bodyWritable
//line /usr/local/go/src/net/http/transport.go:2156
			// _ = "end of CoverTab[44789]"
//line /usr/local/go/src/net/http/transport.go:2156
		}() {
//line /usr/local/go/src/net/http/transport.go:2156
			_go_fuzz_dep_.CoverTab[44790]++
									replaced := pc.t.replaceReqCanceler(rc.cancelKey, nil)

//line /usr/local/go/src/net/http/transport.go:2164
			alive = alive && func() bool {
//line /usr/local/go/src/net/http/transport.go:2164
				_go_fuzz_dep_.CoverTab[44793]++
//line /usr/local/go/src/net/http/transport.go:2164
				return !pc.sawEOF
										// _ = "end of CoverTab[44793]"
//line /usr/local/go/src/net/http/transport.go:2165
			}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2165
				_go_fuzz_dep_.CoverTab[44794]++
//line /usr/local/go/src/net/http/transport.go:2165
				return pc.wroteRequest()
										// _ = "end of CoverTab[44794]"
//line /usr/local/go/src/net/http/transport.go:2166
			}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2166
				_go_fuzz_dep_.CoverTab[44795]++
//line /usr/local/go/src/net/http/transport.go:2166
				return replaced
										// _ = "end of CoverTab[44795]"
//line /usr/local/go/src/net/http/transport.go:2167
			}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2167
				_go_fuzz_dep_.CoverTab[44796]++
//line /usr/local/go/src/net/http/transport.go:2167
				return tryPutIdleConn(trace)
//line /usr/local/go/src/net/http/transport.go:2167
				// _ = "end of CoverTab[44796]"
//line /usr/local/go/src/net/http/transport.go:2167
			}()

			if bodyWritable {
//line /usr/local/go/src/net/http/transport.go:2169
				_go_fuzz_dep_.CoverTab[44797]++
										closeErr = errCallerOwnsConn
//line /usr/local/go/src/net/http/transport.go:2170
				// _ = "end of CoverTab[44797]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2171
				_go_fuzz_dep_.CoverTab[44798]++
//line /usr/local/go/src/net/http/transport.go:2171
				// _ = "end of CoverTab[44798]"
//line /usr/local/go/src/net/http/transport.go:2171
			}
//line /usr/local/go/src/net/http/transport.go:2171
			// _ = "end of CoverTab[44790]"
//line /usr/local/go/src/net/http/transport.go:2171
			_go_fuzz_dep_.CoverTab[44791]++

									select {
			case rc.ch <- responseAndError{res: resp}:
//line /usr/local/go/src/net/http/transport.go:2174
				_go_fuzz_dep_.CoverTab[44799]++
//line /usr/local/go/src/net/http/transport.go:2174
				// _ = "end of CoverTab[44799]"
			case <-rc.callerGone:
//line /usr/local/go/src/net/http/transport.go:2175
				_go_fuzz_dep_.CoverTab[44800]++
										return
//line /usr/local/go/src/net/http/transport.go:2176
				// _ = "end of CoverTab[44800]"
			}
//line /usr/local/go/src/net/http/transport.go:2177
			// _ = "end of CoverTab[44791]"
//line /usr/local/go/src/net/http/transport.go:2177
			_go_fuzz_dep_.CoverTab[44792]++

//line /usr/local/go/src/net/http/transport.go:2182
			testHookReadLoopBeforeNextRead()
									continue
//line /usr/local/go/src/net/http/transport.go:2183
			// _ = "end of CoverTab[44792]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2184
			_go_fuzz_dep_.CoverTab[44801]++
//line /usr/local/go/src/net/http/transport.go:2184
			// _ = "end of CoverTab[44801]"
//line /usr/local/go/src/net/http/transport.go:2184
		}
//line /usr/local/go/src/net/http/transport.go:2184
		// _ = "end of CoverTab[44765]"
//line /usr/local/go/src/net/http/transport.go:2184
		_go_fuzz_dep_.CoverTab[44766]++

								waitForBodyRead := make(chan bool, 2)
								body := &bodyEOFSignal{
			body:	resp.Body,
			earlyCloseFn: func() error {
//line /usr/local/go/src/net/http/transport.go:2189
				_go_fuzz_dep_.CoverTab[44802]++
										waitForBodyRead <- false
										<-eofc
										return nil
//line /usr/local/go/src/net/http/transport.go:2192
				// _ = "end of CoverTab[44802]"

			},
			fn: func(err error) error {
//line /usr/local/go/src/net/http/transport.go:2195
				_go_fuzz_dep_.CoverTab[44803]++
										isEOF := err == io.EOF
										waitForBodyRead <- isEOF
										if isEOF {
//line /usr/local/go/src/net/http/transport.go:2198
					_go_fuzz_dep_.CoverTab[44805]++
											<-eofc
//line /usr/local/go/src/net/http/transport.go:2199
					// _ = "end of CoverTab[44805]"
				} else {
//line /usr/local/go/src/net/http/transport.go:2200
					_go_fuzz_dep_.CoverTab[44806]++
//line /usr/local/go/src/net/http/transport.go:2200
					if err != nil {
//line /usr/local/go/src/net/http/transport.go:2200
						_go_fuzz_dep_.CoverTab[44807]++
												if cerr := pc.canceled(); cerr != nil {
//line /usr/local/go/src/net/http/transport.go:2201
							_go_fuzz_dep_.CoverTab[44808]++
													return cerr
//line /usr/local/go/src/net/http/transport.go:2202
							// _ = "end of CoverTab[44808]"
						} else {
//line /usr/local/go/src/net/http/transport.go:2203
							_go_fuzz_dep_.CoverTab[44809]++
//line /usr/local/go/src/net/http/transport.go:2203
							// _ = "end of CoverTab[44809]"
//line /usr/local/go/src/net/http/transport.go:2203
						}
//line /usr/local/go/src/net/http/transport.go:2203
						// _ = "end of CoverTab[44807]"
					} else {
//line /usr/local/go/src/net/http/transport.go:2204
						_go_fuzz_dep_.CoverTab[44810]++
//line /usr/local/go/src/net/http/transport.go:2204
						// _ = "end of CoverTab[44810]"
//line /usr/local/go/src/net/http/transport.go:2204
					}
//line /usr/local/go/src/net/http/transport.go:2204
					// _ = "end of CoverTab[44806]"
//line /usr/local/go/src/net/http/transport.go:2204
				}
//line /usr/local/go/src/net/http/transport.go:2204
				// _ = "end of CoverTab[44803]"
//line /usr/local/go/src/net/http/transport.go:2204
				_go_fuzz_dep_.CoverTab[44804]++
										return err
//line /usr/local/go/src/net/http/transport.go:2205
				// _ = "end of CoverTab[44804]"
			},
		}
//line /usr/local/go/src/net/http/transport.go:2207
		// _ = "end of CoverTab[44766]"
//line /usr/local/go/src/net/http/transport.go:2207
		_go_fuzz_dep_.CoverTab[44767]++

								resp.Body = body
								if rc.addedGzip && func() bool {
//line /usr/local/go/src/net/http/transport.go:2210
			_go_fuzz_dep_.CoverTab[44811]++
//line /usr/local/go/src/net/http/transport.go:2210
			return ascii.EqualFold(resp.Header.Get("Content-Encoding"), "gzip")
//line /usr/local/go/src/net/http/transport.go:2210
			// _ = "end of CoverTab[44811]"
//line /usr/local/go/src/net/http/transport.go:2210
		}() {
//line /usr/local/go/src/net/http/transport.go:2210
			_go_fuzz_dep_.CoverTab[44812]++
									resp.Body = &gzipReader{body: body}
									resp.Header.Del("Content-Encoding")
									resp.Header.Del("Content-Length")
									resp.ContentLength = -1
									resp.Uncompressed = true
//line /usr/local/go/src/net/http/transport.go:2215
			// _ = "end of CoverTab[44812]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2216
			_go_fuzz_dep_.CoverTab[44813]++
//line /usr/local/go/src/net/http/transport.go:2216
			// _ = "end of CoverTab[44813]"
//line /usr/local/go/src/net/http/transport.go:2216
		}
//line /usr/local/go/src/net/http/transport.go:2216
		// _ = "end of CoverTab[44767]"
//line /usr/local/go/src/net/http/transport.go:2216
		_go_fuzz_dep_.CoverTab[44768]++

								select {
		case rc.ch <- responseAndError{res: resp}:
//line /usr/local/go/src/net/http/transport.go:2219
			_go_fuzz_dep_.CoverTab[44814]++
//line /usr/local/go/src/net/http/transport.go:2219
			// _ = "end of CoverTab[44814]"
		case <-rc.callerGone:
//line /usr/local/go/src/net/http/transport.go:2220
			_go_fuzz_dep_.CoverTab[44815]++
									return
//line /usr/local/go/src/net/http/transport.go:2221
			// _ = "end of CoverTab[44815]"
		}
//line /usr/local/go/src/net/http/transport.go:2222
		// _ = "end of CoverTab[44768]"
//line /usr/local/go/src/net/http/transport.go:2222
		_go_fuzz_dep_.CoverTab[44769]++

//line /usr/local/go/src/net/http/transport.go:2227
		select {
		case bodyEOF := <-waitForBodyRead:
//line /usr/local/go/src/net/http/transport.go:2228
			_go_fuzz_dep_.CoverTab[44816]++
									replaced := pc.t.replaceReqCanceler(rc.cancelKey, nil)
									alive = alive && func() bool {
//line /usr/local/go/src/net/http/transport.go:2230
				_go_fuzz_dep_.CoverTab[44820]++
//line /usr/local/go/src/net/http/transport.go:2230
				return bodyEOF
										// _ = "end of CoverTab[44820]"
//line /usr/local/go/src/net/http/transport.go:2231
			}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2231
				_go_fuzz_dep_.CoverTab[44821]++
//line /usr/local/go/src/net/http/transport.go:2231
				return !pc.sawEOF
										// _ = "end of CoverTab[44821]"
//line /usr/local/go/src/net/http/transport.go:2232
			}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2232
				_go_fuzz_dep_.CoverTab[44822]++
//line /usr/local/go/src/net/http/transport.go:2232
				return pc.wroteRequest()
										// _ = "end of CoverTab[44822]"
//line /usr/local/go/src/net/http/transport.go:2233
			}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2233
				_go_fuzz_dep_.CoverTab[44823]++
//line /usr/local/go/src/net/http/transport.go:2233
				return replaced
										// _ = "end of CoverTab[44823]"
//line /usr/local/go/src/net/http/transport.go:2234
			}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2234
				_go_fuzz_dep_.CoverTab[44824]++
//line /usr/local/go/src/net/http/transport.go:2234
				return tryPutIdleConn(trace)
//line /usr/local/go/src/net/http/transport.go:2234
				// _ = "end of CoverTab[44824]"
//line /usr/local/go/src/net/http/transport.go:2234
			}()
			if bodyEOF {
//line /usr/local/go/src/net/http/transport.go:2235
				_go_fuzz_dep_.CoverTab[44825]++
										eofc <- struct{}{}
//line /usr/local/go/src/net/http/transport.go:2236
				// _ = "end of CoverTab[44825]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2237
				_go_fuzz_dep_.CoverTab[44826]++
//line /usr/local/go/src/net/http/transport.go:2237
				// _ = "end of CoverTab[44826]"
//line /usr/local/go/src/net/http/transport.go:2237
			}
//line /usr/local/go/src/net/http/transport.go:2237
			// _ = "end of CoverTab[44816]"
		case <-rc.req.Cancel:
//line /usr/local/go/src/net/http/transport.go:2238
			_go_fuzz_dep_.CoverTab[44817]++
									alive = false
									pc.t.CancelRequest(rc.req)
//line /usr/local/go/src/net/http/transport.go:2240
			// _ = "end of CoverTab[44817]"
		case <-rc.req.Context().Done():
//line /usr/local/go/src/net/http/transport.go:2241
			_go_fuzz_dep_.CoverTab[44818]++
									alive = false
									pc.t.cancelRequest(rc.cancelKey, rc.req.Context().Err())
//line /usr/local/go/src/net/http/transport.go:2243
			// _ = "end of CoverTab[44818]"
		case <-pc.closech:
//line /usr/local/go/src/net/http/transport.go:2244
			_go_fuzz_dep_.CoverTab[44819]++
									alive = false
//line /usr/local/go/src/net/http/transport.go:2245
			// _ = "end of CoverTab[44819]"
		}
//line /usr/local/go/src/net/http/transport.go:2246
		// _ = "end of CoverTab[44769]"
//line /usr/local/go/src/net/http/transport.go:2246
		_go_fuzz_dep_.CoverTab[44770]++

								testHookReadLoopBeforeNextRead()
//line /usr/local/go/src/net/http/transport.go:2248
		// _ = "end of CoverTab[44770]"
	}
//line /usr/local/go/src/net/http/transport.go:2249
	// _ = "end of CoverTab[44746]"
}

func (pc *persistConn) readLoopPeekFailLocked(peekErr error) {
//line /usr/local/go/src/net/http/transport.go:2252
	_go_fuzz_dep_.CoverTab[44827]++
							if pc.closed != nil {
//line /usr/local/go/src/net/http/transport.go:2253
		_go_fuzz_dep_.CoverTab[44830]++
								return
//line /usr/local/go/src/net/http/transport.go:2254
		// _ = "end of CoverTab[44830]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2255
		_go_fuzz_dep_.CoverTab[44831]++
//line /usr/local/go/src/net/http/transport.go:2255
		// _ = "end of CoverTab[44831]"
//line /usr/local/go/src/net/http/transport.go:2255
	}
//line /usr/local/go/src/net/http/transport.go:2255
	// _ = "end of CoverTab[44827]"
//line /usr/local/go/src/net/http/transport.go:2255
	_go_fuzz_dep_.CoverTab[44828]++
							if n := pc.br.Buffered(); n > 0 {
//line /usr/local/go/src/net/http/transport.go:2256
		_go_fuzz_dep_.CoverTab[44832]++
								buf, _ := pc.br.Peek(n)
								if is408Message(buf) {
//line /usr/local/go/src/net/http/transport.go:2258
			_go_fuzz_dep_.CoverTab[44833]++
									pc.closeLocked(errServerClosedIdle)
									return
//line /usr/local/go/src/net/http/transport.go:2260
			// _ = "end of CoverTab[44833]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2261
			_go_fuzz_dep_.CoverTab[44834]++
									log.Printf("Unsolicited response received on idle HTTP channel starting with %q; err=%v", buf, peekErr)
//line /usr/local/go/src/net/http/transport.go:2262
			// _ = "end of CoverTab[44834]"
		}
//line /usr/local/go/src/net/http/transport.go:2263
		// _ = "end of CoverTab[44832]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2264
		_go_fuzz_dep_.CoverTab[44835]++
//line /usr/local/go/src/net/http/transport.go:2264
		// _ = "end of CoverTab[44835]"
//line /usr/local/go/src/net/http/transport.go:2264
	}
//line /usr/local/go/src/net/http/transport.go:2264
	// _ = "end of CoverTab[44828]"
//line /usr/local/go/src/net/http/transport.go:2264
	_go_fuzz_dep_.CoverTab[44829]++
							if peekErr == io.EOF {
//line /usr/local/go/src/net/http/transport.go:2265
		_go_fuzz_dep_.CoverTab[44836]++

								pc.closeLocked(errServerClosedIdle)
//line /usr/local/go/src/net/http/transport.go:2267
		// _ = "end of CoverTab[44836]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2268
		_go_fuzz_dep_.CoverTab[44837]++
								pc.closeLocked(fmt.Errorf("readLoopPeekFailLocked: %w", peekErr))
//line /usr/local/go/src/net/http/transport.go:2269
		// _ = "end of CoverTab[44837]"
	}
//line /usr/local/go/src/net/http/transport.go:2270
	// _ = "end of CoverTab[44829]"
}

// is408Message reports whether buf has the prefix of an
//line /usr/local/go/src/net/http/transport.go:2273
// HTTP 408 Request Timeout response.
//line /usr/local/go/src/net/http/transport.go:2273
// See golang.org/issue/32310.
//line /usr/local/go/src/net/http/transport.go:2276
func is408Message(buf []byte) bool {
//line /usr/local/go/src/net/http/transport.go:2276
	_go_fuzz_dep_.CoverTab[44838]++
							if len(buf) < len("HTTP/1.x 408") {
//line /usr/local/go/src/net/http/transport.go:2277
		_go_fuzz_dep_.CoverTab[44841]++
								return false
//line /usr/local/go/src/net/http/transport.go:2278
		// _ = "end of CoverTab[44841]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2279
		_go_fuzz_dep_.CoverTab[44842]++
//line /usr/local/go/src/net/http/transport.go:2279
		// _ = "end of CoverTab[44842]"
//line /usr/local/go/src/net/http/transport.go:2279
	}
//line /usr/local/go/src/net/http/transport.go:2279
	// _ = "end of CoverTab[44838]"
//line /usr/local/go/src/net/http/transport.go:2279
	_go_fuzz_dep_.CoverTab[44839]++
							if string(buf[:7]) != "HTTP/1." {
//line /usr/local/go/src/net/http/transport.go:2280
		_go_fuzz_dep_.CoverTab[44843]++
								return false
//line /usr/local/go/src/net/http/transport.go:2281
		// _ = "end of CoverTab[44843]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2282
		_go_fuzz_dep_.CoverTab[44844]++
//line /usr/local/go/src/net/http/transport.go:2282
		// _ = "end of CoverTab[44844]"
//line /usr/local/go/src/net/http/transport.go:2282
	}
//line /usr/local/go/src/net/http/transport.go:2282
	// _ = "end of CoverTab[44839]"
//line /usr/local/go/src/net/http/transport.go:2282
	_go_fuzz_dep_.CoverTab[44840]++
							return string(buf[8:12]) == " 408"
//line /usr/local/go/src/net/http/transport.go:2283
	// _ = "end of CoverTab[44840]"
}

// readResponse reads an HTTP response (or two, in the case of "Expect:
//line /usr/local/go/src/net/http/transport.go:2286
// 100-continue") from the server. It returns the final non-100 one.
//line /usr/local/go/src/net/http/transport.go:2286
// trace is optional.
//line /usr/local/go/src/net/http/transport.go:2289
func (pc *persistConn) readResponse(rc requestAndChan, trace *httptrace.ClientTrace) (resp *Response, err error) {
//line /usr/local/go/src/net/http/transport.go:2289
	_go_fuzz_dep_.CoverTab[44845]++
							if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:2290
		_go_fuzz_dep_.CoverTab[44849]++
//line /usr/local/go/src/net/http/transport.go:2290
		return trace.GotFirstResponseByte != nil
//line /usr/local/go/src/net/http/transport.go:2290
		// _ = "end of CoverTab[44849]"
//line /usr/local/go/src/net/http/transport.go:2290
	}() {
//line /usr/local/go/src/net/http/transport.go:2290
		_go_fuzz_dep_.CoverTab[44850]++
								if peek, err := pc.br.Peek(1); err == nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:2291
			_go_fuzz_dep_.CoverTab[44851]++
//line /usr/local/go/src/net/http/transport.go:2291
			return len(peek) == 1
//line /usr/local/go/src/net/http/transport.go:2291
			// _ = "end of CoverTab[44851]"
//line /usr/local/go/src/net/http/transport.go:2291
		}() {
//line /usr/local/go/src/net/http/transport.go:2291
			_go_fuzz_dep_.CoverTab[44852]++
									trace.GotFirstResponseByte()
//line /usr/local/go/src/net/http/transport.go:2292
			// _ = "end of CoverTab[44852]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2293
			_go_fuzz_dep_.CoverTab[44853]++
//line /usr/local/go/src/net/http/transport.go:2293
			// _ = "end of CoverTab[44853]"
//line /usr/local/go/src/net/http/transport.go:2293
		}
//line /usr/local/go/src/net/http/transport.go:2293
		// _ = "end of CoverTab[44850]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2294
		_go_fuzz_dep_.CoverTab[44854]++
//line /usr/local/go/src/net/http/transport.go:2294
		// _ = "end of CoverTab[44854]"
//line /usr/local/go/src/net/http/transport.go:2294
	}
//line /usr/local/go/src/net/http/transport.go:2294
	// _ = "end of CoverTab[44845]"
//line /usr/local/go/src/net/http/transport.go:2294
	_go_fuzz_dep_.CoverTab[44846]++
							num1xx := 0
							const max1xxResponses = 5	// arbitrary bound on number of informational responses

							continueCh := rc.continueCh
							for {
//line /usr/local/go/src/net/http/transport.go:2299
		_go_fuzz_dep_.CoverTab[44855]++
								resp, err = ReadResponse(pc.br, rc.req)
								if err != nil {
//line /usr/local/go/src/net/http/transport.go:2301
			_go_fuzz_dep_.CoverTab[44859]++
									return
//line /usr/local/go/src/net/http/transport.go:2302
			// _ = "end of CoverTab[44859]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2303
			_go_fuzz_dep_.CoverTab[44860]++
//line /usr/local/go/src/net/http/transport.go:2303
			// _ = "end of CoverTab[44860]"
//line /usr/local/go/src/net/http/transport.go:2303
		}
//line /usr/local/go/src/net/http/transport.go:2303
		// _ = "end of CoverTab[44855]"
//line /usr/local/go/src/net/http/transport.go:2303
		_go_fuzz_dep_.CoverTab[44856]++
								resCode := resp.StatusCode
								if continueCh != nil {
//line /usr/local/go/src/net/http/transport.go:2305
			_go_fuzz_dep_.CoverTab[44861]++
									if resCode == 100 {
//line /usr/local/go/src/net/http/transport.go:2306
				_go_fuzz_dep_.CoverTab[44862]++
										if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:2307
					_go_fuzz_dep_.CoverTab[44864]++
//line /usr/local/go/src/net/http/transport.go:2307
					return trace.Got100Continue != nil
//line /usr/local/go/src/net/http/transport.go:2307
					// _ = "end of CoverTab[44864]"
//line /usr/local/go/src/net/http/transport.go:2307
				}() {
//line /usr/local/go/src/net/http/transport.go:2307
					_go_fuzz_dep_.CoverTab[44865]++
											trace.Got100Continue()
//line /usr/local/go/src/net/http/transport.go:2308
					// _ = "end of CoverTab[44865]"
				} else {
//line /usr/local/go/src/net/http/transport.go:2309
					_go_fuzz_dep_.CoverTab[44866]++
//line /usr/local/go/src/net/http/transport.go:2309
					// _ = "end of CoverTab[44866]"
//line /usr/local/go/src/net/http/transport.go:2309
				}
//line /usr/local/go/src/net/http/transport.go:2309
				// _ = "end of CoverTab[44862]"
//line /usr/local/go/src/net/http/transport.go:2309
				_go_fuzz_dep_.CoverTab[44863]++
										continueCh <- struct{}{}
										continueCh = nil
//line /usr/local/go/src/net/http/transport.go:2311
				// _ = "end of CoverTab[44863]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2312
				_go_fuzz_dep_.CoverTab[44867]++
//line /usr/local/go/src/net/http/transport.go:2312
				if resCode >= 200 {
//line /usr/local/go/src/net/http/transport.go:2312
					_go_fuzz_dep_.CoverTab[44868]++
											close(continueCh)
											continueCh = nil
//line /usr/local/go/src/net/http/transport.go:2314
					// _ = "end of CoverTab[44868]"
				} else {
//line /usr/local/go/src/net/http/transport.go:2315
					_go_fuzz_dep_.CoverTab[44869]++
//line /usr/local/go/src/net/http/transport.go:2315
					// _ = "end of CoverTab[44869]"
//line /usr/local/go/src/net/http/transport.go:2315
				}
//line /usr/local/go/src/net/http/transport.go:2315
				// _ = "end of CoverTab[44867]"
//line /usr/local/go/src/net/http/transport.go:2315
			}
//line /usr/local/go/src/net/http/transport.go:2315
			// _ = "end of CoverTab[44861]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2316
			_go_fuzz_dep_.CoverTab[44870]++
//line /usr/local/go/src/net/http/transport.go:2316
			// _ = "end of CoverTab[44870]"
//line /usr/local/go/src/net/http/transport.go:2316
		}
//line /usr/local/go/src/net/http/transport.go:2316
		// _ = "end of CoverTab[44856]"
//line /usr/local/go/src/net/http/transport.go:2316
		_go_fuzz_dep_.CoverTab[44857]++
								is1xx := 100 <= resCode && func() bool {
//line /usr/local/go/src/net/http/transport.go:2317
			_go_fuzz_dep_.CoverTab[44871]++
//line /usr/local/go/src/net/http/transport.go:2317
			return resCode <= 199
//line /usr/local/go/src/net/http/transport.go:2317
			// _ = "end of CoverTab[44871]"
//line /usr/local/go/src/net/http/transport.go:2317
		}()

								is1xxNonTerminal := is1xx && func() bool {
//line /usr/local/go/src/net/http/transport.go:2319
			_go_fuzz_dep_.CoverTab[44872]++
//line /usr/local/go/src/net/http/transport.go:2319
			return resCode != StatusSwitchingProtocols
//line /usr/local/go/src/net/http/transport.go:2319
			// _ = "end of CoverTab[44872]"
//line /usr/local/go/src/net/http/transport.go:2319
		}()
								if is1xxNonTerminal {
//line /usr/local/go/src/net/http/transport.go:2320
			_go_fuzz_dep_.CoverTab[44873]++
									num1xx++
									if num1xx > max1xxResponses {
//line /usr/local/go/src/net/http/transport.go:2322
				_go_fuzz_dep_.CoverTab[44876]++
										return nil, errors.New("net/http: too many 1xx informational responses")
//line /usr/local/go/src/net/http/transport.go:2323
				// _ = "end of CoverTab[44876]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2324
				_go_fuzz_dep_.CoverTab[44877]++
//line /usr/local/go/src/net/http/transport.go:2324
				// _ = "end of CoverTab[44877]"
//line /usr/local/go/src/net/http/transport.go:2324
			}
//line /usr/local/go/src/net/http/transport.go:2324
			// _ = "end of CoverTab[44873]"
//line /usr/local/go/src/net/http/transport.go:2324
			_go_fuzz_dep_.CoverTab[44874]++
									pc.readLimit = pc.maxHeaderResponseSize()
									if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:2326
				_go_fuzz_dep_.CoverTab[44878]++
//line /usr/local/go/src/net/http/transport.go:2326
				return trace.Got1xxResponse != nil
//line /usr/local/go/src/net/http/transport.go:2326
				// _ = "end of CoverTab[44878]"
//line /usr/local/go/src/net/http/transport.go:2326
			}() {
//line /usr/local/go/src/net/http/transport.go:2326
				_go_fuzz_dep_.CoverTab[44879]++
										if err := trace.Got1xxResponse(resCode, textproto.MIMEHeader(resp.Header)); err != nil {
//line /usr/local/go/src/net/http/transport.go:2327
					_go_fuzz_dep_.CoverTab[44880]++
											return nil, err
//line /usr/local/go/src/net/http/transport.go:2328
					// _ = "end of CoverTab[44880]"
				} else {
//line /usr/local/go/src/net/http/transport.go:2329
					_go_fuzz_dep_.CoverTab[44881]++
//line /usr/local/go/src/net/http/transport.go:2329
					// _ = "end of CoverTab[44881]"
//line /usr/local/go/src/net/http/transport.go:2329
				}
//line /usr/local/go/src/net/http/transport.go:2329
				// _ = "end of CoverTab[44879]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2330
				_go_fuzz_dep_.CoverTab[44882]++
//line /usr/local/go/src/net/http/transport.go:2330
				// _ = "end of CoverTab[44882]"
//line /usr/local/go/src/net/http/transport.go:2330
			}
//line /usr/local/go/src/net/http/transport.go:2330
			// _ = "end of CoverTab[44874]"
//line /usr/local/go/src/net/http/transport.go:2330
			_go_fuzz_dep_.CoverTab[44875]++
									continue
//line /usr/local/go/src/net/http/transport.go:2331
			// _ = "end of CoverTab[44875]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2332
			_go_fuzz_dep_.CoverTab[44883]++
//line /usr/local/go/src/net/http/transport.go:2332
			// _ = "end of CoverTab[44883]"
//line /usr/local/go/src/net/http/transport.go:2332
		}
//line /usr/local/go/src/net/http/transport.go:2332
		// _ = "end of CoverTab[44857]"
//line /usr/local/go/src/net/http/transport.go:2332
		_go_fuzz_dep_.CoverTab[44858]++
								break
//line /usr/local/go/src/net/http/transport.go:2333
		// _ = "end of CoverTab[44858]"
	}
//line /usr/local/go/src/net/http/transport.go:2334
	// _ = "end of CoverTab[44846]"
//line /usr/local/go/src/net/http/transport.go:2334
	_go_fuzz_dep_.CoverTab[44847]++
							if resp.isProtocolSwitch() {
//line /usr/local/go/src/net/http/transport.go:2335
		_go_fuzz_dep_.CoverTab[44884]++
								resp.Body = newReadWriteCloserBody(pc.br, pc.conn)
//line /usr/local/go/src/net/http/transport.go:2336
		// _ = "end of CoverTab[44884]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2337
		_go_fuzz_dep_.CoverTab[44885]++
//line /usr/local/go/src/net/http/transport.go:2337
		// _ = "end of CoverTab[44885]"
//line /usr/local/go/src/net/http/transport.go:2337
	}
//line /usr/local/go/src/net/http/transport.go:2337
	// _ = "end of CoverTab[44847]"
//line /usr/local/go/src/net/http/transport.go:2337
	_go_fuzz_dep_.CoverTab[44848]++

							resp.TLS = pc.tlsState
							return
//line /usr/local/go/src/net/http/transport.go:2340
	// _ = "end of CoverTab[44848]"
}

// waitForContinue returns the function to block until
//line /usr/local/go/src/net/http/transport.go:2343
// any response, timeout or connection close. After any of them,
//line /usr/local/go/src/net/http/transport.go:2343
// the function returns a bool which indicates if the body should be sent.
//line /usr/local/go/src/net/http/transport.go:2346
func (pc *persistConn) waitForContinue(continueCh <-chan struct{}) func() bool {
//line /usr/local/go/src/net/http/transport.go:2346
	_go_fuzz_dep_.CoverTab[44886]++
							if continueCh == nil {
//line /usr/local/go/src/net/http/transport.go:2347
		_go_fuzz_dep_.CoverTab[44888]++
								return nil
//line /usr/local/go/src/net/http/transport.go:2348
		// _ = "end of CoverTab[44888]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2349
		_go_fuzz_dep_.CoverTab[44889]++
//line /usr/local/go/src/net/http/transport.go:2349
		// _ = "end of CoverTab[44889]"
//line /usr/local/go/src/net/http/transport.go:2349
	}
//line /usr/local/go/src/net/http/transport.go:2349
	// _ = "end of CoverTab[44886]"
//line /usr/local/go/src/net/http/transport.go:2349
	_go_fuzz_dep_.CoverTab[44887]++
							return func() bool {
//line /usr/local/go/src/net/http/transport.go:2350
		_go_fuzz_dep_.CoverTab[44890]++
								timer := time.NewTimer(pc.t.ExpectContinueTimeout)
								defer timer.Stop()

								select {
		case _, ok := <-continueCh:
//line /usr/local/go/src/net/http/transport.go:2355
			_go_fuzz_dep_.CoverTab[44891]++
									return ok
//line /usr/local/go/src/net/http/transport.go:2356
			// _ = "end of CoverTab[44891]"
		case <-timer.C:
//line /usr/local/go/src/net/http/transport.go:2357
			_go_fuzz_dep_.CoverTab[44892]++
									return true
//line /usr/local/go/src/net/http/transport.go:2358
			// _ = "end of CoverTab[44892]"
		case <-pc.closech:
//line /usr/local/go/src/net/http/transport.go:2359
			_go_fuzz_dep_.CoverTab[44893]++
									return false
//line /usr/local/go/src/net/http/transport.go:2360
			// _ = "end of CoverTab[44893]"
		}
//line /usr/local/go/src/net/http/transport.go:2361
		// _ = "end of CoverTab[44890]"
	}
//line /usr/local/go/src/net/http/transport.go:2362
	// _ = "end of CoverTab[44887]"
}

func newReadWriteCloserBody(br *bufio.Reader, rwc io.ReadWriteCloser) io.ReadWriteCloser {
//line /usr/local/go/src/net/http/transport.go:2365
	_go_fuzz_dep_.CoverTab[44894]++
							body := &readWriteCloserBody{ReadWriteCloser: rwc}
							if br.Buffered() != 0 {
//line /usr/local/go/src/net/http/transport.go:2367
		_go_fuzz_dep_.CoverTab[44896]++
								body.br = br
//line /usr/local/go/src/net/http/transport.go:2368
		// _ = "end of CoverTab[44896]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2369
		_go_fuzz_dep_.CoverTab[44897]++
//line /usr/local/go/src/net/http/transport.go:2369
		// _ = "end of CoverTab[44897]"
//line /usr/local/go/src/net/http/transport.go:2369
	}
//line /usr/local/go/src/net/http/transport.go:2369
	// _ = "end of CoverTab[44894]"
//line /usr/local/go/src/net/http/transport.go:2369
	_go_fuzz_dep_.CoverTab[44895]++
							return body
//line /usr/local/go/src/net/http/transport.go:2370
	// _ = "end of CoverTab[44895]"
}

// readWriteCloserBody is the Response.Body type used when we want to
//line /usr/local/go/src/net/http/transport.go:2373
// give users write access to the Body through the underlying
//line /usr/local/go/src/net/http/transport.go:2373
// connection (TCP, unless using custom dialers). This is then
//line /usr/local/go/src/net/http/transport.go:2373
// the concrete type for a Response.Body on the 101 Switching
//line /usr/local/go/src/net/http/transport.go:2373
// Protocols response, as used by WebSockets, h2c, etc.
//line /usr/local/go/src/net/http/transport.go:2378
type readWriteCloserBody struct {
	_	incomparable
	br	*bufio.Reader	// used until empty
	io.ReadWriteCloser
}

func (b *readWriteCloserBody) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transport.go:2384
	_go_fuzz_dep_.CoverTab[44898]++
							if b.br != nil {
//line /usr/local/go/src/net/http/transport.go:2385
		_go_fuzz_dep_.CoverTab[44900]++
								if n := b.br.Buffered(); len(p) > n {
//line /usr/local/go/src/net/http/transport.go:2386
			_go_fuzz_dep_.CoverTab[44903]++
									p = p[:n]
//line /usr/local/go/src/net/http/transport.go:2387
			// _ = "end of CoverTab[44903]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2388
			_go_fuzz_dep_.CoverTab[44904]++
//line /usr/local/go/src/net/http/transport.go:2388
			// _ = "end of CoverTab[44904]"
//line /usr/local/go/src/net/http/transport.go:2388
		}
//line /usr/local/go/src/net/http/transport.go:2388
		// _ = "end of CoverTab[44900]"
//line /usr/local/go/src/net/http/transport.go:2388
		_go_fuzz_dep_.CoverTab[44901]++
								n, err = b.br.Read(p)
								if b.br.Buffered() == 0 {
//line /usr/local/go/src/net/http/transport.go:2390
			_go_fuzz_dep_.CoverTab[44905]++
									b.br = nil
//line /usr/local/go/src/net/http/transport.go:2391
			// _ = "end of CoverTab[44905]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2392
			_go_fuzz_dep_.CoverTab[44906]++
//line /usr/local/go/src/net/http/transport.go:2392
			// _ = "end of CoverTab[44906]"
//line /usr/local/go/src/net/http/transport.go:2392
		}
//line /usr/local/go/src/net/http/transport.go:2392
		// _ = "end of CoverTab[44901]"
//line /usr/local/go/src/net/http/transport.go:2392
		_go_fuzz_dep_.CoverTab[44902]++
								return n, err
//line /usr/local/go/src/net/http/transport.go:2393
		// _ = "end of CoverTab[44902]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2394
		_go_fuzz_dep_.CoverTab[44907]++
//line /usr/local/go/src/net/http/transport.go:2394
		// _ = "end of CoverTab[44907]"
//line /usr/local/go/src/net/http/transport.go:2394
	}
//line /usr/local/go/src/net/http/transport.go:2394
	// _ = "end of CoverTab[44898]"
//line /usr/local/go/src/net/http/transport.go:2394
	_go_fuzz_dep_.CoverTab[44899]++
							return b.ReadWriteCloser.Read(p)
//line /usr/local/go/src/net/http/transport.go:2395
	// _ = "end of CoverTab[44899]"
}

// nothingWrittenError wraps a write errors which ended up writing zero bytes.
type nothingWrittenError struct {
	error
}

func (nwe nothingWrittenError) Unwrap() error {
//line /usr/local/go/src/net/http/transport.go:2403
	_go_fuzz_dep_.CoverTab[44908]++
							return nwe.error
//line /usr/local/go/src/net/http/transport.go:2404
	// _ = "end of CoverTab[44908]"
}

func (pc *persistConn) writeLoop() {
//line /usr/local/go/src/net/http/transport.go:2407
	_go_fuzz_dep_.CoverTab[44909]++
							defer close(pc.writeLoopDone)
							for {
//line /usr/local/go/src/net/http/transport.go:2409
		_go_fuzz_dep_.CoverTab[44910]++
								select {
		case wr := <-pc.writech:
//line /usr/local/go/src/net/http/transport.go:2411
			_go_fuzz_dep_.CoverTab[44911]++
									startBytesWritten := pc.nwrite
									err := wr.req.Request.write(pc.bw, pc.isProxy, wr.req.extra, pc.waitForContinue(wr.continueCh))
									if bre, ok := err.(requestBodyReadError); ok {
//line /usr/local/go/src/net/http/transport.go:2414
				_go_fuzz_dep_.CoverTab[44916]++
										err = bre.error

//line /usr/local/go/src/net/http/transport.go:2423
				wr.req.setError(err)
//line /usr/local/go/src/net/http/transport.go:2423
				// _ = "end of CoverTab[44916]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2424
				_go_fuzz_dep_.CoverTab[44917]++
//line /usr/local/go/src/net/http/transport.go:2424
				// _ = "end of CoverTab[44917]"
//line /usr/local/go/src/net/http/transport.go:2424
			}
//line /usr/local/go/src/net/http/transport.go:2424
			// _ = "end of CoverTab[44911]"
//line /usr/local/go/src/net/http/transport.go:2424
			_go_fuzz_dep_.CoverTab[44912]++
									if err == nil {
//line /usr/local/go/src/net/http/transport.go:2425
				_go_fuzz_dep_.CoverTab[44918]++
										err = pc.bw.Flush()
//line /usr/local/go/src/net/http/transport.go:2426
				// _ = "end of CoverTab[44918]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2427
				_go_fuzz_dep_.CoverTab[44919]++
//line /usr/local/go/src/net/http/transport.go:2427
				// _ = "end of CoverTab[44919]"
//line /usr/local/go/src/net/http/transport.go:2427
			}
//line /usr/local/go/src/net/http/transport.go:2427
			// _ = "end of CoverTab[44912]"
//line /usr/local/go/src/net/http/transport.go:2427
			_go_fuzz_dep_.CoverTab[44913]++
									if err != nil {
//line /usr/local/go/src/net/http/transport.go:2428
				_go_fuzz_dep_.CoverTab[44920]++
										if pc.nwrite == startBytesWritten {
//line /usr/local/go/src/net/http/transport.go:2429
					_go_fuzz_dep_.CoverTab[44921]++
											err = nothingWrittenError{err}
//line /usr/local/go/src/net/http/transport.go:2430
					// _ = "end of CoverTab[44921]"
				} else {
//line /usr/local/go/src/net/http/transport.go:2431
					_go_fuzz_dep_.CoverTab[44922]++
//line /usr/local/go/src/net/http/transport.go:2431
					// _ = "end of CoverTab[44922]"
//line /usr/local/go/src/net/http/transport.go:2431
				}
//line /usr/local/go/src/net/http/transport.go:2431
				// _ = "end of CoverTab[44920]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2432
				_go_fuzz_dep_.CoverTab[44923]++
//line /usr/local/go/src/net/http/transport.go:2432
				// _ = "end of CoverTab[44923]"
//line /usr/local/go/src/net/http/transport.go:2432
			}
//line /usr/local/go/src/net/http/transport.go:2432
			// _ = "end of CoverTab[44913]"
//line /usr/local/go/src/net/http/transport.go:2432
			_go_fuzz_dep_.CoverTab[44914]++
									pc.writeErrCh <- err
									wr.ch <- err
									if err != nil {
//line /usr/local/go/src/net/http/transport.go:2435
				_go_fuzz_dep_.CoverTab[44924]++
										pc.close(err)
										return
//line /usr/local/go/src/net/http/transport.go:2437
				// _ = "end of CoverTab[44924]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2438
				_go_fuzz_dep_.CoverTab[44925]++
//line /usr/local/go/src/net/http/transport.go:2438
				// _ = "end of CoverTab[44925]"
//line /usr/local/go/src/net/http/transport.go:2438
			}
//line /usr/local/go/src/net/http/transport.go:2438
			// _ = "end of CoverTab[44914]"
		case <-pc.closech:
//line /usr/local/go/src/net/http/transport.go:2439
			_go_fuzz_dep_.CoverTab[44915]++
									return
//line /usr/local/go/src/net/http/transport.go:2440
			// _ = "end of CoverTab[44915]"
		}
//line /usr/local/go/src/net/http/transport.go:2441
		// _ = "end of CoverTab[44910]"
	}
//line /usr/local/go/src/net/http/transport.go:2442
	// _ = "end of CoverTab[44909]"
}

// maxWriteWaitBeforeConnReuse is how long the a Transport RoundTrip
//line /usr/local/go/src/net/http/transport.go:2445
// will wait to see the Request's Body.Write result after getting a
//line /usr/local/go/src/net/http/transport.go:2445
// response from the server. See comments in (*persistConn).wroteRequest.
//line /usr/local/go/src/net/http/transport.go:2448
const maxWriteWaitBeforeConnReuse = 50 * time.Millisecond

// wroteRequest is a check before recycling a connection that the previous write
//line /usr/local/go/src/net/http/transport.go:2450
// (from writeLoop above) happened and was successful.
//line /usr/local/go/src/net/http/transport.go:2452
func (pc *persistConn) wroteRequest() bool {
//line /usr/local/go/src/net/http/transport.go:2452
	_go_fuzz_dep_.CoverTab[44926]++
							select {
	case err := <-pc.writeErrCh:
//line /usr/local/go/src/net/http/transport.go:2454
		_go_fuzz_dep_.CoverTab[44927]++

//line /usr/local/go/src/net/http/transport.go:2457
		return err == nil
//line /usr/local/go/src/net/http/transport.go:2457
		// _ = "end of CoverTab[44927]"
	default:
//line /usr/local/go/src/net/http/transport.go:2458
		_go_fuzz_dep_.CoverTab[44928]++

//line /usr/local/go/src/net/http/transport.go:2469
		t := time.NewTimer(maxWriteWaitBeforeConnReuse)
		defer t.Stop()
		select {
		case err := <-pc.writeErrCh:
//line /usr/local/go/src/net/http/transport.go:2472
			_go_fuzz_dep_.CoverTab[44929]++
									return err == nil
//line /usr/local/go/src/net/http/transport.go:2473
			// _ = "end of CoverTab[44929]"
		case <-t.C:
//line /usr/local/go/src/net/http/transport.go:2474
			_go_fuzz_dep_.CoverTab[44930]++
									return false
//line /usr/local/go/src/net/http/transport.go:2475
			// _ = "end of CoverTab[44930]"
		}
//line /usr/local/go/src/net/http/transport.go:2476
		// _ = "end of CoverTab[44928]"
	}
//line /usr/local/go/src/net/http/transport.go:2477
	// _ = "end of CoverTab[44926]"
}

// responseAndError is how the goroutine reading from an HTTP/1 server
//line /usr/local/go/src/net/http/transport.go:2480
// communicates with the goroutine doing the RoundTrip.
//line /usr/local/go/src/net/http/transport.go:2482
type responseAndError struct {
	_	incomparable
	res	*Response	// else use this response (see res method)
	err	error
}

type requestAndChan struct {
	_		incomparable
	req		*Request
	cancelKey	cancelKey
	ch		chan responseAndError	// unbuffered; always send in select on callerGone

	// whether the Transport (as opposed to the user client code)
	// added the Accept-Encoding gzip header. If the Transport
	// set it, only then do we transparently decode the gzip.
	addedGzip	bool

	// Optional blocking chan for Expect: 100-continue (for send).
	// If the request has an "Expect: 100-continue" header and
	// the server responds 100 Continue, readLoop send a value
	// to writeLoop via this chan.
	continueCh	chan<- struct{}

	callerGone	<-chan struct{}	// closed when roundTrip caller has returned
}

// A writeRequest is sent by the caller's goroutine to the
//line /usr/local/go/src/net/http/transport.go:2508
// writeLoop's goroutine to write a request while the read loop
//line /usr/local/go/src/net/http/transport.go:2508
// concurrently waits on both the write response and the server's
//line /usr/local/go/src/net/http/transport.go:2508
// reply.
//line /usr/local/go/src/net/http/transport.go:2512
type writeRequest struct {
	req	*transportRequest
	ch	chan<- error

	// Optional blocking chan for Expect: 100-continue (for receive).
	// If not nil, writeLoop blocks sending request body until
	// it receives from this chan.
	continueCh	<-chan struct{}
}

type httpError struct {
	err	string
	timeout	bool
}

func (e *httpError) Error() string {
//line /usr/local/go/src/net/http/transport.go:2527
	_go_fuzz_dep_.CoverTab[44931]++
//line /usr/local/go/src/net/http/transport.go:2527
	return e.err
//line /usr/local/go/src/net/http/transport.go:2527
	// _ = "end of CoverTab[44931]"
//line /usr/local/go/src/net/http/transport.go:2527
}
func (e *httpError) Timeout() bool {
//line /usr/local/go/src/net/http/transport.go:2528
	_go_fuzz_dep_.CoverTab[44932]++
//line /usr/local/go/src/net/http/transport.go:2528
	return e.timeout
//line /usr/local/go/src/net/http/transport.go:2528
	// _ = "end of CoverTab[44932]"
//line /usr/local/go/src/net/http/transport.go:2528
}
func (e *httpError) Temporary() bool {
//line /usr/local/go/src/net/http/transport.go:2529
	_go_fuzz_dep_.CoverTab[44933]++
//line /usr/local/go/src/net/http/transport.go:2529
	return true
//line /usr/local/go/src/net/http/transport.go:2529
	// _ = "end of CoverTab[44933]"
//line /usr/local/go/src/net/http/transport.go:2529
}

var errTimeout error = &httpError{err: "net/http: timeout awaiting response headers", timeout: true}

// errRequestCanceled is set to be identical to the one from h2 to facilitate
//line /usr/local/go/src/net/http/transport.go:2533
// testing.
//line /usr/local/go/src/net/http/transport.go:2535
var errRequestCanceled = http2errRequestCanceled
var errRequestCanceledConn = errors.New("net/http: request canceled while waiting for connection")	// TODO: unify?

func nop()	{ _go_fuzz_dep_.CoverTab[44934]++; // _ = "end of CoverTab[44934]" }

// testHooks. Always non-nil.
var (
	testHookEnterRoundTrip		= nop
	testHookWaitResLoop		= nop
	testHookRoundTripRetried	= nop
	testHookPrePendingDial		= nop
	testHookPostPendingDial		= nop

	testHookMu			sync.Locker	= fakeLocker{}	// guards following
	testHookReadLoopBeforeNextRead			= nop
)

func (pc *persistConn) roundTrip(req *transportRequest) (resp *Response, err error) {
//line /usr/local/go/src/net/http/transport.go:2552
	_go_fuzz_dep_.CoverTab[44935]++
							testHookEnterRoundTrip()
							if !pc.t.replaceReqCanceler(req.cancelKey, pc.cancelRequest) {
//line /usr/local/go/src/net/http/transport.go:2554
		_go_fuzz_dep_.CoverTab[44942]++
								pc.t.putOrCloseIdleConn(pc)
								return nil, errRequestCanceled
//line /usr/local/go/src/net/http/transport.go:2556
		// _ = "end of CoverTab[44942]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2557
		_go_fuzz_dep_.CoverTab[44943]++
//line /usr/local/go/src/net/http/transport.go:2557
		// _ = "end of CoverTab[44943]"
//line /usr/local/go/src/net/http/transport.go:2557
	}
//line /usr/local/go/src/net/http/transport.go:2557
	// _ = "end of CoverTab[44935]"
//line /usr/local/go/src/net/http/transport.go:2557
	_go_fuzz_dep_.CoverTab[44936]++
							pc.mu.Lock()
							pc.numExpectedResponses++
							headerFn := pc.mutateHeaderFunc
							pc.mu.Unlock()

							if headerFn != nil {
//line /usr/local/go/src/net/http/transport.go:2563
		_go_fuzz_dep_.CoverTab[44944]++
								headerFn(req.extraHeaders())
//line /usr/local/go/src/net/http/transport.go:2564
		// _ = "end of CoverTab[44944]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2565
		_go_fuzz_dep_.CoverTab[44945]++
//line /usr/local/go/src/net/http/transport.go:2565
		// _ = "end of CoverTab[44945]"
//line /usr/local/go/src/net/http/transport.go:2565
	}
//line /usr/local/go/src/net/http/transport.go:2565
	// _ = "end of CoverTab[44936]"
//line /usr/local/go/src/net/http/transport.go:2565
	_go_fuzz_dep_.CoverTab[44937]++

//line /usr/local/go/src/net/http/transport.go:2571
	requestedGzip := false
	if !pc.t.DisableCompression && func() bool {
//line /usr/local/go/src/net/http/transport.go:2572
		_go_fuzz_dep_.CoverTab[44946]++
//line /usr/local/go/src/net/http/transport.go:2572
		return req.Header.Get("Accept-Encoding") == ""
								// _ = "end of CoverTab[44946]"
//line /usr/local/go/src/net/http/transport.go:2573
	}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2573
		_go_fuzz_dep_.CoverTab[44947]++
//line /usr/local/go/src/net/http/transport.go:2573
		return req.Header.Get("Range") == ""
								// _ = "end of CoverTab[44947]"
//line /usr/local/go/src/net/http/transport.go:2574
	}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2574
		_go_fuzz_dep_.CoverTab[44948]++
//line /usr/local/go/src/net/http/transport.go:2574
		return req.Method != "HEAD"
								// _ = "end of CoverTab[44948]"
//line /usr/local/go/src/net/http/transport.go:2575
	}() {
//line /usr/local/go/src/net/http/transport.go:2575
		_go_fuzz_dep_.CoverTab[44949]++

//line /usr/local/go/src/net/http/transport.go:2588
		requestedGzip = true
								req.extraHeaders().Set("Accept-Encoding", "gzip")
//line /usr/local/go/src/net/http/transport.go:2589
		// _ = "end of CoverTab[44949]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2590
		_go_fuzz_dep_.CoverTab[44950]++
//line /usr/local/go/src/net/http/transport.go:2590
		// _ = "end of CoverTab[44950]"
//line /usr/local/go/src/net/http/transport.go:2590
	}
//line /usr/local/go/src/net/http/transport.go:2590
	// _ = "end of CoverTab[44937]"
//line /usr/local/go/src/net/http/transport.go:2590
	_go_fuzz_dep_.CoverTab[44938]++

							var continueCh chan struct{}
							if req.ProtoAtLeast(1, 1) && func() bool {
//line /usr/local/go/src/net/http/transport.go:2593
		_go_fuzz_dep_.CoverTab[44951]++
//line /usr/local/go/src/net/http/transport.go:2593
		return req.Body != nil
//line /usr/local/go/src/net/http/transport.go:2593
		// _ = "end of CoverTab[44951]"
//line /usr/local/go/src/net/http/transport.go:2593
	}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2593
		_go_fuzz_dep_.CoverTab[44952]++
//line /usr/local/go/src/net/http/transport.go:2593
		return req.expectsContinue()
//line /usr/local/go/src/net/http/transport.go:2593
		// _ = "end of CoverTab[44952]"
//line /usr/local/go/src/net/http/transport.go:2593
	}() {
//line /usr/local/go/src/net/http/transport.go:2593
		_go_fuzz_dep_.CoverTab[44953]++
								continueCh = make(chan struct{}, 1)
//line /usr/local/go/src/net/http/transport.go:2594
		// _ = "end of CoverTab[44953]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2595
		_go_fuzz_dep_.CoverTab[44954]++
//line /usr/local/go/src/net/http/transport.go:2595
		// _ = "end of CoverTab[44954]"
//line /usr/local/go/src/net/http/transport.go:2595
	}
//line /usr/local/go/src/net/http/transport.go:2595
	// _ = "end of CoverTab[44938]"
//line /usr/local/go/src/net/http/transport.go:2595
	_go_fuzz_dep_.CoverTab[44939]++

							if pc.t.DisableKeepAlives && func() bool {
//line /usr/local/go/src/net/http/transport.go:2597
		_go_fuzz_dep_.CoverTab[44955]++
//line /usr/local/go/src/net/http/transport.go:2597
		return !req.wantsClose()
								// _ = "end of CoverTab[44955]"
//line /usr/local/go/src/net/http/transport.go:2598
	}() && func() bool {
//line /usr/local/go/src/net/http/transport.go:2598
		_go_fuzz_dep_.CoverTab[44956]++
//line /usr/local/go/src/net/http/transport.go:2598
		return !isProtocolSwitchHeader(req.Header)
								// _ = "end of CoverTab[44956]"
//line /usr/local/go/src/net/http/transport.go:2599
	}() {
//line /usr/local/go/src/net/http/transport.go:2599
		_go_fuzz_dep_.CoverTab[44957]++
								req.extraHeaders().Set("Connection", "close")
//line /usr/local/go/src/net/http/transport.go:2600
		// _ = "end of CoverTab[44957]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2601
		_go_fuzz_dep_.CoverTab[44958]++
//line /usr/local/go/src/net/http/transport.go:2601
		// _ = "end of CoverTab[44958]"
//line /usr/local/go/src/net/http/transport.go:2601
	}
//line /usr/local/go/src/net/http/transport.go:2601
	// _ = "end of CoverTab[44939]"
//line /usr/local/go/src/net/http/transport.go:2601
	_go_fuzz_dep_.CoverTab[44940]++

							gone := make(chan struct{})
							defer close(gone)

							defer func() {
//line /usr/local/go/src/net/http/transport.go:2606
		_go_fuzz_dep_.CoverTab[44959]++
								if err != nil {
//line /usr/local/go/src/net/http/transport.go:2607
			_go_fuzz_dep_.CoverTab[44960]++
									pc.t.setReqCanceler(req.cancelKey, nil)
//line /usr/local/go/src/net/http/transport.go:2608
			// _ = "end of CoverTab[44960]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2609
			_go_fuzz_dep_.CoverTab[44961]++
//line /usr/local/go/src/net/http/transport.go:2609
			// _ = "end of CoverTab[44961]"
//line /usr/local/go/src/net/http/transport.go:2609
		}
//line /usr/local/go/src/net/http/transport.go:2609
		// _ = "end of CoverTab[44959]"
	}()
//line /usr/local/go/src/net/http/transport.go:2610
	// _ = "end of CoverTab[44940]"
//line /usr/local/go/src/net/http/transport.go:2610
	_go_fuzz_dep_.CoverTab[44941]++

							const debugRoundTrip = false

//line /usr/local/go/src/net/http/transport.go:2617
	startBytesWritten := pc.nwrite
	writeErrCh := make(chan error, 1)
	pc.writech <- writeRequest{req, writeErrCh, continueCh}

	resc := make(chan responseAndError)
	pc.reqch <- requestAndChan{
		req:		req.Request,
		cancelKey:	req.cancelKey,
		ch:		resc,
		addedGzip:	requestedGzip,
		continueCh:	continueCh,
		callerGone:	gone,
	}

	var respHeaderTimer <-chan time.Time
	cancelChan := req.Request.Cancel
	ctxDoneChan := req.Context().Done()
	pcClosed := pc.closech
	canceled := false
	for {
//line /usr/local/go/src/net/http/transport.go:2636
		_go_fuzz_dep_.CoverTab[44962]++
								testHookWaitResLoop()
								select {
		case err := <-writeErrCh:
//line /usr/local/go/src/net/http/transport.go:2639
			_go_fuzz_dep_.CoverTab[44963]++
									if debugRoundTrip {
//line /usr/local/go/src/net/http/transport.go:2640
				_go_fuzz_dep_.CoverTab[44975]++
										req.logf("writeErrCh resv: %T/%#v", err, err)
//line /usr/local/go/src/net/http/transport.go:2641
				// _ = "end of CoverTab[44975]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2642
				_go_fuzz_dep_.CoverTab[44976]++
//line /usr/local/go/src/net/http/transport.go:2642
				// _ = "end of CoverTab[44976]"
//line /usr/local/go/src/net/http/transport.go:2642
			}
//line /usr/local/go/src/net/http/transport.go:2642
			// _ = "end of CoverTab[44963]"
//line /usr/local/go/src/net/http/transport.go:2642
			_go_fuzz_dep_.CoverTab[44964]++
									if err != nil {
//line /usr/local/go/src/net/http/transport.go:2643
				_go_fuzz_dep_.CoverTab[44977]++
										pc.close(fmt.Errorf("write error: %w", err))
										return nil, pc.mapRoundTripError(req, startBytesWritten, err)
//line /usr/local/go/src/net/http/transport.go:2645
				// _ = "end of CoverTab[44977]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2646
				_go_fuzz_dep_.CoverTab[44978]++
//line /usr/local/go/src/net/http/transport.go:2646
				// _ = "end of CoverTab[44978]"
//line /usr/local/go/src/net/http/transport.go:2646
			}
//line /usr/local/go/src/net/http/transport.go:2646
			// _ = "end of CoverTab[44964]"
//line /usr/local/go/src/net/http/transport.go:2646
			_go_fuzz_dep_.CoverTab[44965]++
									if d := pc.t.ResponseHeaderTimeout; d > 0 {
//line /usr/local/go/src/net/http/transport.go:2647
				_go_fuzz_dep_.CoverTab[44979]++
										if debugRoundTrip {
//line /usr/local/go/src/net/http/transport.go:2648
					_go_fuzz_dep_.CoverTab[44981]++
											req.logf("starting timer for %v", d)
//line /usr/local/go/src/net/http/transport.go:2649
					// _ = "end of CoverTab[44981]"
				} else {
//line /usr/local/go/src/net/http/transport.go:2650
					_go_fuzz_dep_.CoverTab[44982]++
//line /usr/local/go/src/net/http/transport.go:2650
					// _ = "end of CoverTab[44982]"
//line /usr/local/go/src/net/http/transport.go:2650
				}
//line /usr/local/go/src/net/http/transport.go:2650
				// _ = "end of CoverTab[44979]"
//line /usr/local/go/src/net/http/transport.go:2650
				_go_fuzz_dep_.CoverTab[44980]++
										timer := time.NewTimer(d)
										defer timer.Stop()
										respHeaderTimer = timer.C
//line /usr/local/go/src/net/http/transport.go:2653
				// _ = "end of CoverTab[44980]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2654
				_go_fuzz_dep_.CoverTab[44983]++
//line /usr/local/go/src/net/http/transport.go:2654
				// _ = "end of CoverTab[44983]"
//line /usr/local/go/src/net/http/transport.go:2654
			}
//line /usr/local/go/src/net/http/transport.go:2654
			// _ = "end of CoverTab[44965]"
		case <-pcClosed:
//line /usr/local/go/src/net/http/transport.go:2655
			_go_fuzz_dep_.CoverTab[44966]++
									pcClosed = nil
									if canceled || func() bool {
//line /usr/local/go/src/net/http/transport.go:2657
				_go_fuzz_dep_.CoverTab[44984]++
//line /usr/local/go/src/net/http/transport.go:2657
				return pc.t.replaceReqCanceler(req.cancelKey, nil)
//line /usr/local/go/src/net/http/transport.go:2657
				// _ = "end of CoverTab[44984]"
//line /usr/local/go/src/net/http/transport.go:2657
			}() {
//line /usr/local/go/src/net/http/transport.go:2657
				_go_fuzz_dep_.CoverTab[44985]++
										if debugRoundTrip {
//line /usr/local/go/src/net/http/transport.go:2658
					_go_fuzz_dep_.CoverTab[44987]++
											req.logf("closech recv: %T %#v", pc.closed, pc.closed)
//line /usr/local/go/src/net/http/transport.go:2659
					// _ = "end of CoverTab[44987]"
				} else {
//line /usr/local/go/src/net/http/transport.go:2660
					_go_fuzz_dep_.CoverTab[44988]++
//line /usr/local/go/src/net/http/transport.go:2660
					// _ = "end of CoverTab[44988]"
//line /usr/local/go/src/net/http/transport.go:2660
				}
//line /usr/local/go/src/net/http/transport.go:2660
				// _ = "end of CoverTab[44985]"
//line /usr/local/go/src/net/http/transport.go:2660
				_go_fuzz_dep_.CoverTab[44986]++
										return nil, pc.mapRoundTripError(req, startBytesWritten, pc.closed)
//line /usr/local/go/src/net/http/transport.go:2661
				// _ = "end of CoverTab[44986]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2662
				_go_fuzz_dep_.CoverTab[44989]++
//line /usr/local/go/src/net/http/transport.go:2662
				// _ = "end of CoverTab[44989]"
//line /usr/local/go/src/net/http/transport.go:2662
			}
//line /usr/local/go/src/net/http/transport.go:2662
			// _ = "end of CoverTab[44966]"
		case <-respHeaderTimer:
//line /usr/local/go/src/net/http/transport.go:2663
			_go_fuzz_dep_.CoverTab[44967]++
									if debugRoundTrip {
//line /usr/local/go/src/net/http/transport.go:2664
				_go_fuzz_dep_.CoverTab[44990]++
										req.logf("timeout waiting for response headers.")
//line /usr/local/go/src/net/http/transport.go:2665
				// _ = "end of CoverTab[44990]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2666
				_go_fuzz_dep_.CoverTab[44991]++
//line /usr/local/go/src/net/http/transport.go:2666
				// _ = "end of CoverTab[44991]"
//line /usr/local/go/src/net/http/transport.go:2666
			}
//line /usr/local/go/src/net/http/transport.go:2666
			// _ = "end of CoverTab[44967]"
//line /usr/local/go/src/net/http/transport.go:2666
			_go_fuzz_dep_.CoverTab[44968]++
									pc.close(errTimeout)
									return nil, errTimeout
//line /usr/local/go/src/net/http/transport.go:2668
			// _ = "end of CoverTab[44968]"
		case re := <-resc:
//line /usr/local/go/src/net/http/transport.go:2669
			_go_fuzz_dep_.CoverTab[44969]++
									if (re.res == nil) == (re.err == nil) {
//line /usr/local/go/src/net/http/transport.go:2670
				_go_fuzz_dep_.CoverTab[44992]++
										panic(fmt.Sprintf("internal error: exactly one of res or err should be set; nil=%v", re.res == nil))
//line /usr/local/go/src/net/http/transport.go:2671
				// _ = "end of CoverTab[44992]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2672
				_go_fuzz_dep_.CoverTab[44993]++
//line /usr/local/go/src/net/http/transport.go:2672
				// _ = "end of CoverTab[44993]"
//line /usr/local/go/src/net/http/transport.go:2672
			}
//line /usr/local/go/src/net/http/transport.go:2672
			// _ = "end of CoverTab[44969]"
//line /usr/local/go/src/net/http/transport.go:2672
			_go_fuzz_dep_.CoverTab[44970]++
									if debugRoundTrip {
//line /usr/local/go/src/net/http/transport.go:2673
				_go_fuzz_dep_.CoverTab[44994]++
										req.logf("resc recv: %p, %T/%#v", re.res, re.err, re.err)
//line /usr/local/go/src/net/http/transport.go:2674
				// _ = "end of CoverTab[44994]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2675
				_go_fuzz_dep_.CoverTab[44995]++
//line /usr/local/go/src/net/http/transport.go:2675
				// _ = "end of CoverTab[44995]"
//line /usr/local/go/src/net/http/transport.go:2675
			}
//line /usr/local/go/src/net/http/transport.go:2675
			// _ = "end of CoverTab[44970]"
//line /usr/local/go/src/net/http/transport.go:2675
			_go_fuzz_dep_.CoverTab[44971]++
									if re.err != nil {
//line /usr/local/go/src/net/http/transport.go:2676
				_go_fuzz_dep_.CoverTab[44996]++
										return nil, pc.mapRoundTripError(req, startBytesWritten, re.err)
//line /usr/local/go/src/net/http/transport.go:2677
				// _ = "end of CoverTab[44996]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2678
				_go_fuzz_dep_.CoverTab[44997]++
//line /usr/local/go/src/net/http/transport.go:2678
				// _ = "end of CoverTab[44997]"
//line /usr/local/go/src/net/http/transport.go:2678
			}
//line /usr/local/go/src/net/http/transport.go:2678
			// _ = "end of CoverTab[44971]"
//line /usr/local/go/src/net/http/transport.go:2678
			_go_fuzz_dep_.CoverTab[44972]++
									return re.res, nil
//line /usr/local/go/src/net/http/transport.go:2679
			// _ = "end of CoverTab[44972]"
		case <-cancelChan:
//line /usr/local/go/src/net/http/transport.go:2680
			_go_fuzz_dep_.CoverTab[44973]++
									canceled = pc.t.cancelRequest(req.cancelKey, errRequestCanceled)
									cancelChan = nil
//line /usr/local/go/src/net/http/transport.go:2682
			// _ = "end of CoverTab[44973]"
		case <-ctxDoneChan:
//line /usr/local/go/src/net/http/transport.go:2683
			_go_fuzz_dep_.CoverTab[44974]++
									canceled = pc.t.cancelRequest(req.cancelKey, req.Context().Err())
									cancelChan = nil
									ctxDoneChan = nil
//line /usr/local/go/src/net/http/transport.go:2686
			// _ = "end of CoverTab[44974]"
		}
//line /usr/local/go/src/net/http/transport.go:2687
		// _ = "end of CoverTab[44962]"
	}
//line /usr/local/go/src/net/http/transport.go:2688
	// _ = "end of CoverTab[44941]"
}

// tLogKey is a context WithValue key for test debugging contexts containing
//line /usr/local/go/src/net/http/transport.go:2691
// a t.Logf func. See export_test.go's Request.WithT method.
//line /usr/local/go/src/net/http/transport.go:2693
type tLogKey struct{}

func (tr *transportRequest) logf(format string, args ...any) {
//line /usr/local/go/src/net/http/transport.go:2695
	_go_fuzz_dep_.CoverTab[44998]++
							if logf, ok := tr.Request.Context().Value(tLogKey{}).(func(string, ...any)); ok {
//line /usr/local/go/src/net/http/transport.go:2696
		_go_fuzz_dep_.CoverTab[44999]++
								logf(time.Now().Format(time.RFC3339Nano)+": "+format, args...)
//line /usr/local/go/src/net/http/transport.go:2697
		// _ = "end of CoverTab[44999]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2698
		_go_fuzz_dep_.CoverTab[45000]++
//line /usr/local/go/src/net/http/transport.go:2698
		// _ = "end of CoverTab[45000]"
//line /usr/local/go/src/net/http/transport.go:2698
	}
//line /usr/local/go/src/net/http/transport.go:2698
	// _ = "end of CoverTab[44998]"
}

// markReused marks this connection as having been successfully used for a
//line /usr/local/go/src/net/http/transport.go:2701
// request and response.
//line /usr/local/go/src/net/http/transport.go:2703
func (pc *persistConn) markReused() {
//line /usr/local/go/src/net/http/transport.go:2703
	_go_fuzz_dep_.CoverTab[45001]++
							pc.mu.Lock()
							pc.reused = true
							pc.mu.Unlock()
//line /usr/local/go/src/net/http/transport.go:2706
	// _ = "end of CoverTab[45001]"
}

// close closes the underlying TCP connection and closes
//line /usr/local/go/src/net/http/transport.go:2709
// the pc.closech channel.
//line /usr/local/go/src/net/http/transport.go:2709
//
//line /usr/local/go/src/net/http/transport.go:2709
// The provided err is only for testing and debugging; in normal
//line /usr/local/go/src/net/http/transport.go:2709
// circumstances it should never be seen by users.
//line /usr/local/go/src/net/http/transport.go:2714
func (pc *persistConn) close(err error) {
//line /usr/local/go/src/net/http/transport.go:2714
	_go_fuzz_dep_.CoverTab[45002]++
							pc.mu.Lock()
							defer pc.mu.Unlock()
							pc.closeLocked(err)
//line /usr/local/go/src/net/http/transport.go:2717
	// _ = "end of CoverTab[45002]"
}

func (pc *persistConn) closeLocked(err error) {
//line /usr/local/go/src/net/http/transport.go:2720
	_go_fuzz_dep_.CoverTab[45003]++
							if err == nil {
//line /usr/local/go/src/net/http/transport.go:2721
		_go_fuzz_dep_.CoverTab[45006]++
								panic("nil error")
//line /usr/local/go/src/net/http/transport.go:2722
		// _ = "end of CoverTab[45006]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2723
		_go_fuzz_dep_.CoverTab[45007]++
//line /usr/local/go/src/net/http/transport.go:2723
		// _ = "end of CoverTab[45007]"
//line /usr/local/go/src/net/http/transport.go:2723
	}
//line /usr/local/go/src/net/http/transport.go:2723
	// _ = "end of CoverTab[45003]"
//line /usr/local/go/src/net/http/transport.go:2723
	_go_fuzz_dep_.CoverTab[45004]++
							pc.broken = true
							if pc.closed == nil {
//line /usr/local/go/src/net/http/transport.go:2725
		_go_fuzz_dep_.CoverTab[45008]++
								pc.closed = err
								pc.t.decConnsPerHost(pc.cacheKey)

//line /usr/local/go/src/net/http/transport.go:2730
		if pc.alt == nil {
//line /usr/local/go/src/net/http/transport.go:2730
			_go_fuzz_dep_.CoverTab[45009]++
									if err != errCallerOwnsConn {
//line /usr/local/go/src/net/http/transport.go:2731
				_go_fuzz_dep_.CoverTab[45011]++
										pc.conn.Close()
//line /usr/local/go/src/net/http/transport.go:2732
				// _ = "end of CoverTab[45011]"
			} else {
//line /usr/local/go/src/net/http/transport.go:2733
				_go_fuzz_dep_.CoverTab[45012]++
//line /usr/local/go/src/net/http/transport.go:2733
				// _ = "end of CoverTab[45012]"
//line /usr/local/go/src/net/http/transport.go:2733
			}
//line /usr/local/go/src/net/http/transport.go:2733
			// _ = "end of CoverTab[45009]"
//line /usr/local/go/src/net/http/transport.go:2733
			_go_fuzz_dep_.CoverTab[45010]++
									close(pc.closech)
//line /usr/local/go/src/net/http/transport.go:2734
			// _ = "end of CoverTab[45010]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2735
			_go_fuzz_dep_.CoverTab[45013]++
//line /usr/local/go/src/net/http/transport.go:2735
			// _ = "end of CoverTab[45013]"
//line /usr/local/go/src/net/http/transport.go:2735
		}
//line /usr/local/go/src/net/http/transport.go:2735
		// _ = "end of CoverTab[45008]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2736
		_go_fuzz_dep_.CoverTab[45014]++
//line /usr/local/go/src/net/http/transport.go:2736
		// _ = "end of CoverTab[45014]"
//line /usr/local/go/src/net/http/transport.go:2736
	}
//line /usr/local/go/src/net/http/transport.go:2736
	// _ = "end of CoverTab[45004]"
//line /usr/local/go/src/net/http/transport.go:2736
	_go_fuzz_dep_.CoverTab[45005]++
							pc.mutateHeaderFunc = nil
//line /usr/local/go/src/net/http/transport.go:2737
	// _ = "end of CoverTab[45005]"
}

var portMap = map[string]string{
	"http":		"80",
	"https":	"443",
	"socks5":	"1080",
}

// canonicalAddr returns url.Host but always with a ":port" suffix.
func canonicalAddr(url *url.URL) string {
//line /usr/local/go/src/net/http/transport.go:2747
	_go_fuzz_dep_.CoverTab[45015]++
							addr := url.Hostname()
							if v, err := idnaASCII(addr); err == nil {
//line /usr/local/go/src/net/http/transport.go:2749
		_go_fuzz_dep_.CoverTab[45018]++
								addr = v
//line /usr/local/go/src/net/http/transport.go:2750
		// _ = "end of CoverTab[45018]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2751
		_go_fuzz_dep_.CoverTab[45019]++
//line /usr/local/go/src/net/http/transport.go:2751
		// _ = "end of CoverTab[45019]"
//line /usr/local/go/src/net/http/transport.go:2751
	}
//line /usr/local/go/src/net/http/transport.go:2751
	// _ = "end of CoverTab[45015]"
//line /usr/local/go/src/net/http/transport.go:2751
	_go_fuzz_dep_.CoverTab[45016]++
							port := url.Port()
							if port == "" {
//line /usr/local/go/src/net/http/transport.go:2753
		_go_fuzz_dep_.CoverTab[45020]++
								port = portMap[url.Scheme]
//line /usr/local/go/src/net/http/transport.go:2754
		// _ = "end of CoverTab[45020]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2755
		_go_fuzz_dep_.CoverTab[45021]++
//line /usr/local/go/src/net/http/transport.go:2755
		// _ = "end of CoverTab[45021]"
//line /usr/local/go/src/net/http/transport.go:2755
	}
//line /usr/local/go/src/net/http/transport.go:2755
	// _ = "end of CoverTab[45016]"
//line /usr/local/go/src/net/http/transport.go:2755
	_go_fuzz_dep_.CoverTab[45017]++
							return net.JoinHostPort(addr, port)
//line /usr/local/go/src/net/http/transport.go:2756
	// _ = "end of CoverTab[45017]"
}

// bodyEOFSignal is used by the HTTP/1 transport when reading response
//line /usr/local/go/src/net/http/transport.go:2759
// bodies to make sure we see the end of a response body before
//line /usr/local/go/src/net/http/transport.go:2759
// proceeding and reading on the connection again.
//line /usr/local/go/src/net/http/transport.go:2759
//
//line /usr/local/go/src/net/http/transport.go:2759
// It wraps a ReadCloser but runs fn (if non-nil) at most
//line /usr/local/go/src/net/http/transport.go:2759
// once, right before its final (error-producing) Read or Close call
//line /usr/local/go/src/net/http/transport.go:2759
// returns. fn should return the new error to return from Read or Close.
//line /usr/local/go/src/net/http/transport.go:2759
//
//line /usr/local/go/src/net/http/transport.go:2759
// If earlyCloseFn is non-nil and Close is called before io.EOF is
//line /usr/local/go/src/net/http/transport.go:2759
// seen, earlyCloseFn is called instead of fn, and its return value is
//line /usr/local/go/src/net/http/transport.go:2759
// the return value from Close.
//line /usr/local/go/src/net/http/transport.go:2770
type bodyEOFSignal struct {
	body		io.ReadCloser
	mu		sync.Mutex		// guards following 4 fields
	closed		bool			// whether Close has been called
	rerr		error			// sticky Read error
	fn		func(error) error	// err will be nil on Read io.EOF
	earlyCloseFn	func() error		// optional alt Close func used if io.EOF not seen
}

var errReadOnClosedResBody = errors.New("http: read on closed response body")

func (es *bodyEOFSignal) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transport.go:2781
	_go_fuzz_dep_.CoverTab[45022]++
							es.mu.Lock()
							closed, rerr := es.closed, es.rerr
							es.mu.Unlock()
							if closed {
//line /usr/local/go/src/net/http/transport.go:2785
		_go_fuzz_dep_.CoverTab[45026]++
								return 0, errReadOnClosedResBody
//line /usr/local/go/src/net/http/transport.go:2786
		// _ = "end of CoverTab[45026]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2787
		_go_fuzz_dep_.CoverTab[45027]++
//line /usr/local/go/src/net/http/transport.go:2787
		// _ = "end of CoverTab[45027]"
//line /usr/local/go/src/net/http/transport.go:2787
	}
//line /usr/local/go/src/net/http/transport.go:2787
	// _ = "end of CoverTab[45022]"
//line /usr/local/go/src/net/http/transport.go:2787
	_go_fuzz_dep_.CoverTab[45023]++
							if rerr != nil {
//line /usr/local/go/src/net/http/transport.go:2788
		_go_fuzz_dep_.CoverTab[45028]++
								return 0, rerr
//line /usr/local/go/src/net/http/transport.go:2789
		// _ = "end of CoverTab[45028]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2790
		_go_fuzz_dep_.CoverTab[45029]++
//line /usr/local/go/src/net/http/transport.go:2790
		// _ = "end of CoverTab[45029]"
//line /usr/local/go/src/net/http/transport.go:2790
	}
//line /usr/local/go/src/net/http/transport.go:2790
	// _ = "end of CoverTab[45023]"
//line /usr/local/go/src/net/http/transport.go:2790
	_go_fuzz_dep_.CoverTab[45024]++

							n, err = es.body.Read(p)
							if err != nil {
//line /usr/local/go/src/net/http/transport.go:2793
		_go_fuzz_dep_.CoverTab[45030]++
								es.mu.Lock()
								defer es.mu.Unlock()
								if es.rerr == nil {
//line /usr/local/go/src/net/http/transport.go:2796
			_go_fuzz_dep_.CoverTab[45032]++
									es.rerr = err
//line /usr/local/go/src/net/http/transport.go:2797
			// _ = "end of CoverTab[45032]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2798
			_go_fuzz_dep_.CoverTab[45033]++
//line /usr/local/go/src/net/http/transport.go:2798
			// _ = "end of CoverTab[45033]"
//line /usr/local/go/src/net/http/transport.go:2798
		}
//line /usr/local/go/src/net/http/transport.go:2798
		// _ = "end of CoverTab[45030]"
//line /usr/local/go/src/net/http/transport.go:2798
		_go_fuzz_dep_.CoverTab[45031]++
								err = es.condfn(err)
//line /usr/local/go/src/net/http/transport.go:2799
		// _ = "end of CoverTab[45031]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2800
		_go_fuzz_dep_.CoverTab[45034]++
//line /usr/local/go/src/net/http/transport.go:2800
		// _ = "end of CoverTab[45034]"
//line /usr/local/go/src/net/http/transport.go:2800
	}
//line /usr/local/go/src/net/http/transport.go:2800
	// _ = "end of CoverTab[45024]"
//line /usr/local/go/src/net/http/transport.go:2800
	_go_fuzz_dep_.CoverTab[45025]++
							return
//line /usr/local/go/src/net/http/transport.go:2801
	// _ = "end of CoverTab[45025]"
}

func (es *bodyEOFSignal) Close() error {
//line /usr/local/go/src/net/http/transport.go:2804
	_go_fuzz_dep_.CoverTab[45035]++
							es.mu.Lock()
							defer es.mu.Unlock()
							if es.closed {
//line /usr/local/go/src/net/http/transport.go:2807
		_go_fuzz_dep_.CoverTab[45038]++
								return nil
//line /usr/local/go/src/net/http/transport.go:2808
		// _ = "end of CoverTab[45038]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2809
		_go_fuzz_dep_.CoverTab[45039]++
//line /usr/local/go/src/net/http/transport.go:2809
		// _ = "end of CoverTab[45039]"
//line /usr/local/go/src/net/http/transport.go:2809
	}
//line /usr/local/go/src/net/http/transport.go:2809
	// _ = "end of CoverTab[45035]"
//line /usr/local/go/src/net/http/transport.go:2809
	_go_fuzz_dep_.CoverTab[45036]++
							es.closed = true
							if es.earlyCloseFn != nil && func() bool {
//line /usr/local/go/src/net/http/transport.go:2811
		_go_fuzz_dep_.CoverTab[45040]++
//line /usr/local/go/src/net/http/transport.go:2811
		return es.rerr != io.EOF
//line /usr/local/go/src/net/http/transport.go:2811
		// _ = "end of CoverTab[45040]"
//line /usr/local/go/src/net/http/transport.go:2811
	}() {
//line /usr/local/go/src/net/http/transport.go:2811
		_go_fuzz_dep_.CoverTab[45041]++
								return es.earlyCloseFn()
//line /usr/local/go/src/net/http/transport.go:2812
		// _ = "end of CoverTab[45041]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2813
		_go_fuzz_dep_.CoverTab[45042]++
//line /usr/local/go/src/net/http/transport.go:2813
		// _ = "end of CoverTab[45042]"
//line /usr/local/go/src/net/http/transport.go:2813
	}
//line /usr/local/go/src/net/http/transport.go:2813
	// _ = "end of CoverTab[45036]"
//line /usr/local/go/src/net/http/transport.go:2813
	_go_fuzz_dep_.CoverTab[45037]++
							err := es.body.Close()
							return es.condfn(err)
//line /usr/local/go/src/net/http/transport.go:2815
	// _ = "end of CoverTab[45037]"
}

// caller must hold es.mu.
func (es *bodyEOFSignal) condfn(err error) error {
//line /usr/local/go/src/net/http/transport.go:2819
	_go_fuzz_dep_.CoverTab[45043]++
							if es.fn == nil {
//line /usr/local/go/src/net/http/transport.go:2820
		_go_fuzz_dep_.CoverTab[45045]++
								return err
//line /usr/local/go/src/net/http/transport.go:2821
		// _ = "end of CoverTab[45045]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2822
		_go_fuzz_dep_.CoverTab[45046]++
//line /usr/local/go/src/net/http/transport.go:2822
		// _ = "end of CoverTab[45046]"
//line /usr/local/go/src/net/http/transport.go:2822
	}
//line /usr/local/go/src/net/http/transport.go:2822
	// _ = "end of CoverTab[45043]"
//line /usr/local/go/src/net/http/transport.go:2822
	_go_fuzz_dep_.CoverTab[45044]++
							err = es.fn(err)
							es.fn = nil
							return err
//line /usr/local/go/src/net/http/transport.go:2825
	// _ = "end of CoverTab[45044]"
}

// gzipReader wraps a response body so it can lazily
//line /usr/local/go/src/net/http/transport.go:2828
// call gzip.NewReader on the first call to Read
//line /usr/local/go/src/net/http/transport.go:2830
type gzipReader struct {
	_	incomparable
	body	*bodyEOFSignal	// underlying HTTP/1 response body framing
	zr	*gzip.Reader	// lazily-initialized gzip reader
	zerr	error		// any error from gzip.NewReader; sticky
}

func (gz *gzipReader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transport.go:2837
	_go_fuzz_dep_.CoverTab[45047]++
							if gz.zr == nil {
//line /usr/local/go/src/net/http/transport.go:2838
		_go_fuzz_dep_.CoverTab[45051]++
								if gz.zerr == nil {
//line /usr/local/go/src/net/http/transport.go:2839
			_go_fuzz_dep_.CoverTab[45053]++
									gz.zr, gz.zerr = gzip.NewReader(gz.body)
//line /usr/local/go/src/net/http/transport.go:2840
			// _ = "end of CoverTab[45053]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2841
			_go_fuzz_dep_.CoverTab[45054]++
//line /usr/local/go/src/net/http/transport.go:2841
			// _ = "end of CoverTab[45054]"
//line /usr/local/go/src/net/http/transport.go:2841
		}
//line /usr/local/go/src/net/http/transport.go:2841
		// _ = "end of CoverTab[45051]"
//line /usr/local/go/src/net/http/transport.go:2841
		_go_fuzz_dep_.CoverTab[45052]++
								if gz.zerr != nil {
//line /usr/local/go/src/net/http/transport.go:2842
			_go_fuzz_dep_.CoverTab[45055]++
									return 0, gz.zerr
//line /usr/local/go/src/net/http/transport.go:2843
			// _ = "end of CoverTab[45055]"
		} else {
//line /usr/local/go/src/net/http/transport.go:2844
			_go_fuzz_dep_.CoverTab[45056]++
//line /usr/local/go/src/net/http/transport.go:2844
			// _ = "end of CoverTab[45056]"
//line /usr/local/go/src/net/http/transport.go:2844
		}
//line /usr/local/go/src/net/http/transport.go:2844
		// _ = "end of CoverTab[45052]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2845
		_go_fuzz_dep_.CoverTab[45057]++
//line /usr/local/go/src/net/http/transport.go:2845
		// _ = "end of CoverTab[45057]"
//line /usr/local/go/src/net/http/transport.go:2845
	}
//line /usr/local/go/src/net/http/transport.go:2845
	// _ = "end of CoverTab[45047]"
//line /usr/local/go/src/net/http/transport.go:2845
	_go_fuzz_dep_.CoverTab[45048]++

							gz.body.mu.Lock()
							if gz.body.closed {
//line /usr/local/go/src/net/http/transport.go:2848
		_go_fuzz_dep_.CoverTab[45058]++
								err = errReadOnClosedResBody
//line /usr/local/go/src/net/http/transport.go:2849
		// _ = "end of CoverTab[45058]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2850
		_go_fuzz_dep_.CoverTab[45059]++
//line /usr/local/go/src/net/http/transport.go:2850
		// _ = "end of CoverTab[45059]"
//line /usr/local/go/src/net/http/transport.go:2850
	}
//line /usr/local/go/src/net/http/transport.go:2850
	// _ = "end of CoverTab[45048]"
//line /usr/local/go/src/net/http/transport.go:2850
	_go_fuzz_dep_.CoverTab[45049]++
							gz.body.mu.Unlock()

							if err != nil {
//line /usr/local/go/src/net/http/transport.go:2853
		_go_fuzz_dep_.CoverTab[45060]++
								return 0, err
//line /usr/local/go/src/net/http/transport.go:2854
		// _ = "end of CoverTab[45060]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2855
		_go_fuzz_dep_.CoverTab[45061]++
//line /usr/local/go/src/net/http/transport.go:2855
		// _ = "end of CoverTab[45061]"
//line /usr/local/go/src/net/http/transport.go:2855
	}
//line /usr/local/go/src/net/http/transport.go:2855
	// _ = "end of CoverTab[45049]"
//line /usr/local/go/src/net/http/transport.go:2855
	_go_fuzz_dep_.CoverTab[45050]++
							return gz.zr.Read(p)
//line /usr/local/go/src/net/http/transport.go:2856
	// _ = "end of CoverTab[45050]"
}

func (gz *gzipReader) Close() error {
//line /usr/local/go/src/net/http/transport.go:2859
	_go_fuzz_dep_.CoverTab[45062]++
							return gz.body.Close()
//line /usr/local/go/src/net/http/transport.go:2860
	// _ = "end of CoverTab[45062]"
}

type tlsHandshakeTimeoutError struct{}

func (tlsHandshakeTimeoutError) Timeout() bool {
//line /usr/local/go/src/net/http/transport.go:2865
	_go_fuzz_dep_.CoverTab[45063]++
//line /usr/local/go/src/net/http/transport.go:2865
	return true
//line /usr/local/go/src/net/http/transport.go:2865
	// _ = "end of CoverTab[45063]"
//line /usr/local/go/src/net/http/transport.go:2865
}
func (tlsHandshakeTimeoutError) Temporary() bool {
//line /usr/local/go/src/net/http/transport.go:2866
	_go_fuzz_dep_.CoverTab[45064]++
//line /usr/local/go/src/net/http/transport.go:2866
	return true
//line /usr/local/go/src/net/http/transport.go:2866
	// _ = "end of CoverTab[45064]"
//line /usr/local/go/src/net/http/transport.go:2866
}
func (tlsHandshakeTimeoutError) Error() string {
//line /usr/local/go/src/net/http/transport.go:2867
	_go_fuzz_dep_.CoverTab[45065]++
//line /usr/local/go/src/net/http/transport.go:2867
	return "net/http: TLS handshake timeout"
//line /usr/local/go/src/net/http/transport.go:2867
	// _ = "end of CoverTab[45065]"
//line /usr/local/go/src/net/http/transport.go:2867
}

// fakeLocker is a sync.Locker which does nothing. It's used to guard
//line /usr/local/go/src/net/http/transport.go:2869
// test-only fields when not under test, to avoid runtime atomic
//line /usr/local/go/src/net/http/transport.go:2869
// overhead.
//line /usr/local/go/src/net/http/transport.go:2872
type fakeLocker struct{}

func (fakeLocker) Lock()	{ _go_fuzz_dep_.CoverTab[45066]++; // _ = "end of CoverTab[45066]" }
func (fakeLocker) Unlock()	{ _go_fuzz_dep_.CoverTab[45067]++; // _ = "end of CoverTab[45067]" }

// cloneTLSConfig returns a shallow clone of cfg, or a new zero tls.Config if
//line /usr/local/go/src/net/http/transport.go:2877
// cfg is nil. This is safe to call even if cfg is in active use by a TLS
//line /usr/local/go/src/net/http/transport.go:2877
// client or server.
//line /usr/local/go/src/net/http/transport.go:2880
func cloneTLSConfig(cfg *tls.Config) *tls.Config {
//line /usr/local/go/src/net/http/transport.go:2880
	_go_fuzz_dep_.CoverTab[45068]++
							if cfg == nil {
//line /usr/local/go/src/net/http/transport.go:2881
		_go_fuzz_dep_.CoverTab[45070]++
								return &tls.Config{}
//line /usr/local/go/src/net/http/transport.go:2882
		// _ = "end of CoverTab[45070]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2883
		_go_fuzz_dep_.CoverTab[45071]++
//line /usr/local/go/src/net/http/transport.go:2883
		// _ = "end of CoverTab[45071]"
//line /usr/local/go/src/net/http/transport.go:2883
	}
//line /usr/local/go/src/net/http/transport.go:2883
	// _ = "end of CoverTab[45068]"
//line /usr/local/go/src/net/http/transport.go:2883
	_go_fuzz_dep_.CoverTab[45069]++
							return cfg.Clone()
//line /usr/local/go/src/net/http/transport.go:2884
	// _ = "end of CoverTab[45069]"
}

type connLRU struct {
	ll	*list.List	// list.Element.Value type is of *persistConn
	m	map[*persistConn]*list.Element
}

// add adds pc to the head of the linked list.
func (cl *connLRU) add(pc *persistConn) {
//line /usr/local/go/src/net/http/transport.go:2893
	_go_fuzz_dep_.CoverTab[45072]++
							if cl.ll == nil {
//line /usr/local/go/src/net/http/transport.go:2894
		_go_fuzz_dep_.CoverTab[45075]++
								cl.ll = list.New()
								cl.m = make(map[*persistConn]*list.Element)
//line /usr/local/go/src/net/http/transport.go:2896
		// _ = "end of CoverTab[45075]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2897
		_go_fuzz_dep_.CoverTab[45076]++
//line /usr/local/go/src/net/http/transport.go:2897
		// _ = "end of CoverTab[45076]"
//line /usr/local/go/src/net/http/transport.go:2897
	}
//line /usr/local/go/src/net/http/transport.go:2897
	// _ = "end of CoverTab[45072]"
//line /usr/local/go/src/net/http/transport.go:2897
	_go_fuzz_dep_.CoverTab[45073]++
							ele := cl.ll.PushFront(pc)
							if _, ok := cl.m[pc]; ok {
//line /usr/local/go/src/net/http/transport.go:2899
		_go_fuzz_dep_.CoverTab[45077]++
								panic("persistConn was already in LRU")
//line /usr/local/go/src/net/http/transport.go:2900
		// _ = "end of CoverTab[45077]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2901
		_go_fuzz_dep_.CoverTab[45078]++
//line /usr/local/go/src/net/http/transport.go:2901
		// _ = "end of CoverTab[45078]"
//line /usr/local/go/src/net/http/transport.go:2901
	}
//line /usr/local/go/src/net/http/transport.go:2901
	// _ = "end of CoverTab[45073]"
//line /usr/local/go/src/net/http/transport.go:2901
	_go_fuzz_dep_.CoverTab[45074]++
							cl.m[pc] = ele
//line /usr/local/go/src/net/http/transport.go:2902
	// _ = "end of CoverTab[45074]"
}

func (cl *connLRU) removeOldest() *persistConn {
//line /usr/local/go/src/net/http/transport.go:2905
	_go_fuzz_dep_.CoverTab[45079]++
							ele := cl.ll.Back()
							pc := ele.Value.(*persistConn)
							cl.ll.Remove(ele)
							delete(cl.m, pc)
							return pc
//line /usr/local/go/src/net/http/transport.go:2910
	// _ = "end of CoverTab[45079]"
}

// remove removes pc from cl.
func (cl *connLRU) remove(pc *persistConn) {
//line /usr/local/go/src/net/http/transport.go:2914
	_go_fuzz_dep_.CoverTab[45080]++
							if ele, ok := cl.m[pc]; ok {
//line /usr/local/go/src/net/http/transport.go:2915
		_go_fuzz_dep_.CoverTab[45081]++
								cl.ll.Remove(ele)
								delete(cl.m, pc)
//line /usr/local/go/src/net/http/transport.go:2917
		// _ = "end of CoverTab[45081]"
	} else {
//line /usr/local/go/src/net/http/transport.go:2918
		_go_fuzz_dep_.CoverTab[45082]++
//line /usr/local/go/src/net/http/transport.go:2918
		// _ = "end of CoverTab[45082]"
//line /usr/local/go/src/net/http/transport.go:2918
	}
//line /usr/local/go/src/net/http/transport.go:2918
	// _ = "end of CoverTab[45080]"
}

// len returns the number of items in the cache.
func (cl *connLRU) len() int {
//line /usr/local/go/src/net/http/transport.go:2922
	_go_fuzz_dep_.CoverTab[45083]++
							return len(cl.m)
//line /usr/local/go/src/net/http/transport.go:2923
	// _ = "end of CoverTab[45083]"
}

//line /usr/local/go/src/net/http/transport.go:2924
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/transport.go:2924
var _ = _go_fuzz_dep_.CoverTab
