// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/httptrace/trace.go:5
// Package httptrace provides mechanisms to trace the events within
//line /usr/local/go/src/net/http/httptrace/trace.go:5
// HTTP client requests.
//line /usr/local/go/src/net/http/httptrace/trace.go:7
package httptrace

//line /usr/local/go/src/net/http/httptrace/trace.go:7
import (
//line /usr/local/go/src/net/http/httptrace/trace.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/httptrace/trace.go:7
)
//line /usr/local/go/src/net/http/httptrace/trace.go:7
import (
//line /usr/local/go/src/net/http/httptrace/trace.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/httptrace/trace.go:7
)

import (
	"context"
	"crypto/tls"
	"internal/nettrace"
	"net"
	"net/textproto"
	"reflect"
	"time"
)

// unique type to prevent assignment.
type clientEventContextKey struct{}

// ContextClientTrace returns the ClientTrace associated with the
//line /usr/local/go/src/net/http/httptrace/trace.go:22
// provided context. If none, it returns nil.
//line /usr/local/go/src/net/http/httptrace/trace.go:24
func ContextClientTrace(ctx context.Context) *ClientTrace {
//line /usr/local/go/src/net/http/httptrace/trace.go:24
	_go_fuzz_dep_.CoverTab[36418]++
								trace, _ := ctx.Value(clientEventContextKey{}).(*ClientTrace)
								return trace
//line /usr/local/go/src/net/http/httptrace/trace.go:26
	// _ = "end of CoverTab[36418]"
}

// WithClientTrace returns a new context based on the provided parent
//line /usr/local/go/src/net/http/httptrace/trace.go:29
// ctx. HTTP client requests made with the returned context will use
//line /usr/local/go/src/net/http/httptrace/trace.go:29
// the provided trace hooks, in addition to any previous hooks
//line /usr/local/go/src/net/http/httptrace/trace.go:29
// registered with ctx. Any hooks defined in the provided trace will
//line /usr/local/go/src/net/http/httptrace/trace.go:29
// be called first.
//line /usr/local/go/src/net/http/httptrace/trace.go:34
func WithClientTrace(ctx context.Context, trace *ClientTrace) context.Context {
//line /usr/local/go/src/net/http/httptrace/trace.go:34
	_go_fuzz_dep_.CoverTab[36419]++
								if trace == nil {
//line /usr/local/go/src/net/http/httptrace/trace.go:35
		_go_fuzz_dep_.CoverTab[36422]++
									panic("nil trace")
//line /usr/local/go/src/net/http/httptrace/trace.go:36
		// _ = "end of CoverTab[36422]"
	} else {
//line /usr/local/go/src/net/http/httptrace/trace.go:37
		_go_fuzz_dep_.CoverTab[36423]++
//line /usr/local/go/src/net/http/httptrace/trace.go:37
		// _ = "end of CoverTab[36423]"
//line /usr/local/go/src/net/http/httptrace/trace.go:37
	}
//line /usr/local/go/src/net/http/httptrace/trace.go:37
	// _ = "end of CoverTab[36419]"
//line /usr/local/go/src/net/http/httptrace/trace.go:37
	_go_fuzz_dep_.CoverTab[36420]++
								old := ContextClientTrace(ctx)
								trace.compose(old)

								ctx = context.WithValue(ctx, clientEventContextKey{}, trace)
								if trace.hasNetHooks() {
//line /usr/local/go/src/net/http/httptrace/trace.go:42
		_go_fuzz_dep_.CoverTab[36424]++
									nt := &nettrace.Trace{
			ConnectStart:	trace.ConnectStart,
			ConnectDone:	trace.ConnectDone,
		}
		if trace.DNSStart != nil {
//line /usr/local/go/src/net/http/httptrace/trace.go:47
			_go_fuzz_dep_.CoverTab[36427]++
										nt.DNSStart = func(name string) {
//line /usr/local/go/src/net/http/httptrace/trace.go:48
				_go_fuzz_dep_.CoverTab[36428]++
											trace.DNSStart(DNSStartInfo{Host: name})
//line /usr/local/go/src/net/http/httptrace/trace.go:49
				// _ = "end of CoverTab[36428]"
			}
//line /usr/local/go/src/net/http/httptrace/trace.go:50
			// _ = "end of CoverTab[36427]"
		} else {
//line /usr/local/go/src/net/http/httptrace/trace.go:51
			_go_fuzz_dep_.CoverTab[36429]++
//line /usr/local/go/src/net/http/httptrace/trace.go:51
			// _ = "end of CoverTab[36429]"
//line /usr/local/go/src/net/http/httptrace/trace.go:51
		}
//line /usr/local/go/src/net/http/httptrace/trace.go:51
		// _ = "end of CoverTab[36424]"
//line /usr/local/go/src/net/http/httptrace/trace.go:51
		_go_fuzz_dep_.CoverTab[36425]++
									if trace.DNSDone != nil {
//line /usr/local/go/src/net/http/httptrace/trace.go:52
			_go_fuzz_dep_.CoverTab[36430]++
										nt.DNSDone = func(netIPs []any, coalesced bool, err error) {
//line /usr/local/go/src/net/http/httptrace/trace.go:53
				_go_fuzz_dep_.CoverTab[36431]++
											addrs := make([]net.IPAddr, len(netIPs))
											for i, ip := range netIPs {
//line /usr/local/go/src/net/http/httptrace/trace.go:55
					_go_fuzz_dep_.CoverTab[36433]++
												addrs[i] = ip.(net.IPAddr)
//line /usr/local/go/src/net/http/httptrace/trace.go:56
					// _ = "end of CoverTab[36433]"
				}
//line /usr/local/go/src/net/http/httptrace/trace.go:57
				// _ = "end of CoverTab[36431]"
//line /usr/local/go/src/net/http/httptrace/trace.go:57
				_go_fuzz_dep_.CoverTab[36432]++
											trace.DNSDone(DNSDoneInfo{
					Addrs:		addrs,
					Coalesced:	coalesced,
					Err:		err,
				})
//line /usr/local/go/src/net/http/httptrace/trace.go:62
				// _ = "end of CoverTab[36432]"
			}
//line /usr/local/go/src/net/http/httptrace/trace.go:63
			// _ = "end of CoverTab[36430]"
		} else {
//line /usr/local/go/src/net/http/httptrace/trace.go:64
			_go_fuzz_dep_.CoverTab[36434]++
//line /usr/local/go/src/net/http/httptrace/trace.go:64
			// _ = "end of CoverTab[36434]"
//line /usr/local/go/src/net/http/httptrace/trace.go:64
		}
//line /usr/local/go/src/net/http/httptrace/trace.go:64
		// _ = "end of CoverTab[36425]"
//line /usr/local/go/src/net/http/httptrace/trace.go:64
		_go_fuzz_dep_.CoverTab[36426]++
									ctx = context.WithValue(ctx, nettrace.TraceKey{}, nt)
//line /usr/local/go/src/net/http/httptrace/trace.go:65
		// _ = "end of CoverTab[36426]"
	} else {
//line /usr/local/go/src/net/http/httptrace/trace.go:66
		_go_fuzz_dep_.CoverTab[36435]++
//line /usr/local/go/src/net/http/httptrace/trace.go:66
		// _ = "end of CoverTab[36435]"
//line /usr/local/go/src/net/http/httptrace/trace.go:66
	}
//line /usr/local/go/src/net/http/httptrace/trace.go:66
	// _ = "end of CoverTab[36420]"
//line /usr/local/go/src/net/http/httptrace/trace.go:66
	_go_fuzz_dep_.CoverTab[36421]++
								return ctx
//line /usr/local/go/src/net/http/httptrace/trace.go:67
	// _ = "end of CoverTab[36421]"
}

// ClientTrace is a set of hooks to run at various stages of an outgoing
//line /usr/local/go/src/net/http/httptrace/trace.go:70
// HTTP request. Any particular hook may be nil. Functions may be
//line /usr/local/go/src/net/http/httptrace/trace.go:70
// called concurrently from different goroutines and some may be called
//line /usr/local/go/src/net/http/httptrace/trace.go:70
// after the request has completed or failed.
//line /usr/local/go/src/net/http/httptrace/trace.go:70
//
//line /usr/local/go/src/net/http/httptrace/trace.go:70
// ClientTrace currently traces a single HTTP request & response
//line /usr/local/go/src/net/http/httptrace/trace.go:70
// during a single round trip and has no hooks that span a series
//line /usr/local/go/src/net/http/httptrace/trace.go:70
// of redirected requests.
//line /usr/local/go/src/net/http/httptrace/trace.go:70
//
//line /usr/local/go/src/net/http/httptrace/trace.go:70
// See https://blog.golang.org/http-tracing for more.
//line /usr/local/go/src/net/http/httptrace/trace.go:80
type ClientTrace struct {
	// GetConn is called before a connection is created or
	// retrieved from an idle pool. The hostPort is the
	// "host:port" of the target or proxy. GetConn is called even
	// if there's already an idle cached connection available.
	GetConn	func(hostPort string)

	// GotConn is called after a successful connection is
	// obtained. There is no hook for failure to obtain a
	// connection; instead, use the error from
	// Transport.RoundTrip.
	GotConn	func(GotConnInfo)

	// PutIdleConn is called when the connection is returned to
	// the idle pool. If err is nil, the connection was
	// successfully returned to the idle pool. If err is non-nil,
	// it describes why not. PutIdleConn is not called if
	// connection reuse is disabled via Transport.DisableKeepAlives.
	// PutIdleConn is called before the caller's Response.Body.Close
	// call returns.
	// For HTTP/2, this hook is not currently used.
	PutIdleConn	func(err error)

	// GotFirstResponseByte is called when the first byte of the response
	// headers is available.
	GotFirstResponseByte	func()

	// Got100Continue is called if the server replies with a "100
	// Continue" response.
	Got100Continue	func()

	// Got1xxResponse is called for each 1xx informational response header
	// returned before the final non-1xx response. Got1xxResponse is called
	// for "100 Continue" responses, even if Got100Continue is also defined.
	// If it returns an error, the client request is aborted with that error value.
	Got1xxResponse	func(code int, header textproto.MIMEHeader) error

	// DNSStart is called when a DNS lookup begins.
	DNSStart	func(DNSStartInfo)

	// DNSDone is called when a DNS lookup ends.
	DNSDone	func(DNSDoneInfo)

	// ConnectStart is called when a new connection's Dial begins.
	// If net.Dialer.DualStack (IPv6 "Happy Eyeballs") support is
	// enabled, this may be called multiple times.
	ConnectStart	func(network, addr string)

	// ConnectDone is called when a new connection's Dial
	// completes. The provided err indicates whether the
	// connection completed successfully.
	// If net.Dialer.DualStack ("Happy Eyeballs") support is
	// enabled, this may be called multiple times.
	ConnectDone	func(network, addr string, err error)

	// TLSHandshakeStart is called when the TLS handshake is started. When
	// connecting to an HTTPS site via an HTTP proxy, the handshake happens
	// after the CONNECT request is processed by the proxy.
	TLSHandshakeStart	func()

	// TLSHandshakeDone is called after the TLS handshake with either the
	// successful handshake's connection state, or a non-nil error on handshake
	// failure.
	TLSHandshakeDone	func(tls.ConnectionState, error)

	// WroteHeaderField is called after the Transport has written
	// each request header. At the time of this call the values
	// might be buffered and not yet written to the network.
	WroteHeaderField	func(key string, value []string)

	// WroteHeaders is called after the Transport has written
	// all request headers.
	WroteHeaders	func()

	// Wait100Continue is called if the Request specified
	// "Expect: 100-continue" and the Transport has written the
	// request headers but is waiting for "100 Continue" from the
	// server before writing the request body.
	Wait100Continue	func()

	// WroteRequest is called with the result of writing the
	// request and any body. It may be called multiple times
	// in the case of retried requests.
	WroteRequest	func(WroteRequestInfo)
}

// WroteRequestInfo contains information provided to the WroteRequest
//line /usr/local/go/src/net/http/httptrace/trace.go:166
// hook.
//line /usr/local/go/src/net/http/httptrace/trace.go:168
type WroteRequestInfo struct {
	// Err is any error encountered while writing the Request.
	Err error
}

// compose modifies t such that it respects the previously-registered hooks in old,
//line /usr/local/go/src/net/http/httptrace/trace.go:173
// subject to the composition policy requested in t.Compose.
//line /usr/local/go/src/net/http/httptrace/trace.go:175
func (t *ClientTrace) compose(old *ClientTrace) {
//line /usr/local/go/src/net/http/httptrace/trace.go:175
	_go_fuzz_dep_.CoverTab[36436]++
								if old == nil {
//line /usr/local/go/src/net/http/httptrace/trace.go:176
		_go_fuzz_dep_.CoverTab[36438]++
									return
//line /usr/local/go/src/net/http/httptrace/trace.go:177
		// _ = "end of CoverTab[36438]"
	} else {
//line /usr/local/go/src/net/http/httptrace/trace.go:178
		_go_fuzz_dep_.CoverTab[36439]++
//line /usr/local/go/src/net/http/httptrace/trace.go:178
		// _ = "end of CoverTab[36439]"
//line /usr/local/go/src/net/http/httptrace/trace.go:178
	}
//line /usr/local/go/src/net/http/httptrace/trace.go:178
	// _ = "end of CoverTab[36436]"
//line /usr/local/go/src/net/http/httptrace/trace.go:178
	_go_fuzz_dep_.CoverTab[36437]++
								tv := reflect.ValueOf(t).Elem()
								ov := reflect.ValueOf(old).Elem()
								structType := tv.Type()
								for i := 0; i < structType.NumField(); i++ {
//line /usr/local/go/src/net/http/httptrace/trace.go:182
		_go_fuzz_dep_.CoverTab[36440]++
									tf := tv.Field(i)
									hookType := tf.Type()
									if hookType.Kind() != reflect.Func {
//line /usr/local/go/src/net/http/httptrace/trace.go:185
			_go_fuzz_dep_.CoverTab[36445]++
										continue
//line /usr/local/go/src/net/http/httptrace/trace.go:186
			// _ = "end of CoverTab[36445]"
		} else {
//line /usr/local/go/src/net/http/httptrace/trace.go:187
			_go_fuzz_dep_.CoverTab[36446]++
//line /usr/local/go/src/net/http/httptrace/trace.go:187
			// _ = "end of CoverTab[36446]"
//line /usr/local/go/src/net/http/httptrace/trace.go:187
		}
//line /usr/local/go/src/net/http/httptrace/trace.go:187
		// _ = "end of CoverTab[36440]"
//line /usr/local/go/src/net/http/httptrace/trace.go:187
		_go_fuzz_dep_.CoverTab[36441]++
									of := ov.Field(i)
									if of.IsNil() {
//line /usr/local/go/src/net/http/httptrace/trace.go:189
			_go_fuzz_dep_.CoverTab[36447]++
										continue
//line /usr/local/go/src/net/http/httptrace/trace.go:190
			// _ = "end of CoverTab[36447]"
		} else {
//line /usr/local/go/src/net/http/httptrace/trace.go:191
			_go_fuzz_dep_.CoverTab[36448]++
//line /usr/local/go/src/net/http/httptrace/trace.go:191
			// _ = "end of CoverTab[36448]"
//line /usr/local/go/src/net/http/httptrace/trace.go:191
		}
//line /usr/local/go/src/net/http/httptrace/trace.go:191
		// _ = "end of CoverTab[36441]"
//line /usr/local/go/src/net/http/httptrace/trace.go:191
		_go_fuzz_dep_.CoverTab[36442]++
									if tf.IsNil() {
//line /usr/local/go/src/net/http/httptrace/trace.go:192
			_go_fuzz_dep_.CoverTab[36449]++
										tf.Set(of)
										continue
//line /usr/local/go/src/net/http/httptrace/trace.go:194
			// _ = "end of CoverTab[36449]"
		} else {
//line /usr/local/go/src/net/http/httptrace/trace.go:195
			_go_fuzz_dep_.CoverTab[36450]++
//line /usr/local/go/src/net/http/httptrace/trace.go:195
			// _ = "end of CoverTab[36450]"
//line /usr/local/go/src/net/http/httptrace/trace.go:195
		}
//line /usr/local/go/src/net/http/httptrace/trace.go:195
		// _ = "end of CoverTab[36442]"
//line /usr/local/go/src/net/http/httptrace/trace.go:195
		_go_fuzz_dep_.CoverTab[36443]++

//line /usr/local/go/src/net/http/httptrace/trace.go:199
		tfCopy := reflect.ValueOf(tf.Interface())

//line /usr/local/go/src/net/http/httptrace/trace.go:202
		newFunc := reflect.MakeFunc(hookType, func(args []reflect.Value) []reflect.Value {
//line /usr/local/go/src/net/http/httptrace/trace.go:202
			_go_fuzz_dep_.CoverTab[36451]++
										tfCopy.Call(args)
										return of.Call(args)
//line /usr/local/go/src/net/http/httptrace/trace.go:204
			// _ = "end of CoverTab[36451]"
		})
//line /usr/local/go/src/net/http/httptrace/trace.go:205
		// _ = "end of CoverTab[36443]"
//line /usr/local/go/src/net/http/httptrace/trace.go:205
		_go_fuzz_dep_.CoverTab[36444]++
									tv.Field(i).Set(newFunc)
//line /usr/local/go/src/net/http/httptrace/trace.go:206
		// _ = "end of CoverTab[36444]"
	}
//line /usr/local/go/src/net/http/httptrace/trace.go:207
	// _ = "end of CoverTab[36437]"
}

// DNSStartInfo contains information about a DNS request.
type DNSStartInfo struct {
	Host string
}

// DNSDoneInfo contains information about the results of a DNS lookup.
type DNSDoneInfo struct {
	// Addrs are the IPv4 and/or IPv6 addresses found in the DNS
	// lookup. The contents of the slice should not be mutated.
	Addrs	[]net.IPAddr

	// Err is any error that occurred during the DNS lookup.
	Err	error

	// Coalesced is whether the Addrs were shared with another
	// caller who was doing the same DNS lookup concurrently.
	Coalesced	bool
}

func (t *ClientTrace) hasNetHooks() bool {
//line /usr/local/go/src/net/http/httptrace/trace.go:229
	_go_fuzz_dep_.CoverTab[36452]++
								if t == nil {
//line /usr/local/go/src/net/http/httptrace/trace.go:230
		_go_fuzz_dep_.CoverTab[36454]++
									return false
//line /usr/local/go/src/net/http/httptrace/trace.go:231
		// _ = "end of CoverTab[36454]"
	} else {
//line /usr/local/go/src/net/http/httptrace/trace.go:232
		_go_fuzz_dep_.CoverTab[36455]++
//line /usr/local/go/src/net/http/httptrace/trace.go:232
		// _ = "end of CoverTab[36455]"
//line /usr/local/go/src/net/http/httptrace/trace.go:232
	}
//line /usr/local/go/src/net/http/httptrace/trace.go:232
	// _ = "end of CoverTab[36452]"
//line /usr/local/go/src/net/http/httptrace/trace.go:232
	_go_fuzz_dep_.CoverTab[36453]++
								return t.DNSStart != nil || func() bool {
//line /usr/local/go/src/net/http/httptrace/trace.go:233
		_go_fuzz_dep_.CoverTab[36456]++
//line /usr/local/go/src/net/http/httptrace/trace.go:233
		return t.DNSDone != nil
//line /usr/local/go/src/net/http/httptrace/trace.go:233
		// _ = "end of CoverTab[36456]"
//line /usr/local/go/src/net/http/httptrace/trace.go:233
	}() || func() bool {
//line /usr/local/go/src/net/http/httptrace/trace.go:233
		_go_fuzz_dep_.CoverTab[36457]++
//line /usr/local/go/src/net/http/httptrace/trace.go:233
		return t.ConnectStart != nil
//line /usr/local/go/src/net/http/httptrace/trace.go:233
		// _ = "end of CoverTab[36457]"
//line /usr/local/go/src/net/http/httptrace/trace.go:233
	}() || func() bool {
//line /usr/local/go/src/net/http/httptrace/trace.go:233
		_go_fuzz_dep_.CoverTab[36458]++
//line /usr/local/go/src/net/http/httptrace/trace.go:233
		return t.ConnectDone != nil
//line /usr/local/go/src/net/http/httptrace/trace.go:233
		// _ = "end of CoverTab[36458]"
//line /usr/local/go/src/net/http/httptrace/trace.go:233
	}()
//line /usr/local/go/src/net/http/httptrace/trace.go:233
	// _ = "end of CoverTab[36453]"
}

// GotConnInfo is the argument to the ClientTrace.GotConn function and
//line /usr/local/go/src/net/http/httptrace/trace.go:236
// contains information about the obtained connection.
//line /usr/local/go/src/net/http/httptrace/trace.go:238
type GotConnInfo struct {
	// Conn is the connection that was obtained. It is owned by
	// the http.Transport and should not be read, written or
	// closed by users of ClientTrace.
	Conn	net.Conn

	// Reused is whether this connection has been previously
	// used for another HTTP request.
	Reused	bool

	// WasIdle is whether this connection was obtained from an
	// idle pool.
	WasIdle	bool

	// IdleTime reports how long the connection was previously
	// idle, if WasIdle is true.
	IdleTime	time.Duration
}

//line /usr/local/go/src/net/http/httptrace/trace.go:255
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/httptrace/trace.go:255
var _ = _go_fuzz_dep_.CoverTab
