// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Transport code's client connection pooling.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:7
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:7
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:7
)

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"sync"
)

// ClientConnPool manages a pool of HTTP/2 client connections.
type ClientConnPool interface {
	// GetClientConn returns a specific HTTP/2 connection (usually
	// a TLS-TCP connection) to an HTTP/2 server. On success, the
	// returned ClientConn accounts for the upcoming RoundTrip
	// call, so the caller should not omit it. If the caller needs
	// to, ClientConn.RoundTrip can be called with a bogus
	// new(http.Request) to release the stream reservation.
	GetClientConn(req *http.Request, addr string) (*ClientConn, error)
	MarkDead(*ClientConn)
}

// clientConnPoolIdleCloser is the interface implemented by ClientConnPool
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:29
// implementations which can close their idle connections.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:31
type clientConnPoolIdleCloser interface {
	ClientConnPool
	closeIdleConnections()
}

var (
	_	clientConnPoolIdleCloser	= (*clientConnPool)(nil)
	_	clientConnPoolIdleCloser	= noDialClientConnPool{}
)

// TODO: use singleflight for dialing and addConnCalls?
type clientConnPool struct {
	t	*Transport

	mu	sync.Mutex	// TODO: maybe switch to RWMutex
	// TODO: add support for sharing conns based on cert names
	// (e.g. share conn for googleapis.com and appspot.com)
	conns		map[string][]*ClientConn	// key is host:port
	dialing		map[string]*dialCall		// currently in-flight dials
	keys		map[*ClientConn][]string
	addConnCalls	map[string]*addConnCall	// in-flight addConnIfNeeded calls
}

func (p *clientConnPool) GetClientConn(req *http.Request, addr string) (*ClientConn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:54
	_go_fuzz_dep_.CoverTab[72241]++
												return p.getClientConn(req, addr, dialOnMiss)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:55
	// _ = "end of CoverTab[72241]"
}

const (
	dialOnMiss	= true
	noDialOnMiss	= false
)

func (p *clientConnPool) getClientConn(req *http.Request, addr string, dialOnMiss bool) (*ClientConn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:63
	_go_fuzz_dep_.CoverTab[72242]++

												if isConnectionCloseRequest(req) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:65
		_go_fuzz_dep_.CoverTab[72244]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:65
		return dialOnMiss
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:65
		// _ = "end of CoverTab[72244]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:65
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:65
		_go_fuzz_dep_.CoverTab[72245]++

													traceGetConn(req, addr)
													const singleUse = true
													cc, err := p.t.dialClientConn(req.Context(), addr, singleUse)
													if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:70
			_go_fuzz_dep_.CoverTab[72247]++
														return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:71
			// _ = "end of CoverTab[72247]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:72
			_go_fuzz_dep_.CoverTab[72248]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:72
			// _ = "end of CoverTab[72248]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:72
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:72
		// _ = "end of CoverTab[72245]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:72
		_go_fuzz_dep_.CoverTab[72246]++
													return cc, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:73
		// _ = "end of CoverTab[72246]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:74
		_go_fuzz_dep_.CoverTab[72249]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:74
		// _ = "end of CoverTab[72249]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:74
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:74
	// _ = "end of CoverTab[72242]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:74
	_go_fuzz_dep_.CoverTab[72243]++
												for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:75
		_go_fuzz_dep_.CoverTab[72250]++
													p.mu.Lock()
													for _, cc := range p.conns[addr] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:77
			_go_fuzz_dep_.CoverTab[72255]++
														if cc.ReserveNewRequest() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:78
				_go_fuzz_dep_.CoverTab[72256]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:82
				if !cc.getConnCalled {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:82
					_go_fuzz_dep_.CoverTab[72258]++
																traceGetConn(req, addr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:83
					// _ = "end of CoverTab[72258]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:84
					_go_fuzz_dep_.CoverTab[72259]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:84
					// _ = "end of CoverTab[72259]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:84
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:84
				// _ = "end of CoverTab[72256]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:84
				_go_fuzz_dep_.CoverTab[72257]++
															cc.getConnCalled = false
															p.mu.Unlock()
															return cc, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:87
				// _ = "end of CoverTab[72257]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:88
				_go_fuzz_dep_.CoverTab[72260]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:88
				// _ = "end of CoverTab[72260]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:88
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:88
			// _ = "end of CoverTab[72255]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:89
		// _ = "end of CoverTab[72250]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:89
		_go_fuzz_dep_.CoverTab[72251]++
													if !dialOnMiss {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:90
			_go_fuzz_dep_.CoverTab[72261]++
														p.mu.Unlock()
														return nil, ErrNoCachedConn
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:92
			// _ = "end of CoverTab[72261]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:93
			_go_fuzz_dep_.CoverTab[72262]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:93
			// _ = "end of CoverTab[72262]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:93
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:93
		// _ = "end of CoverTab[72251]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:93
		_go_fuzz_dep_.CoverTab[72252]++
													traceGetConn(req, addr)
													call := p.getStartDialLocked(req.Context(), addr)
													p.mu.Unlock()
													<-call.done
													if shouldRetryDial(call, req) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:98
			_go_fuzz_dep_.CoverTab[72263]++
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:99
			// _ = "end of CoverTab[72263]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:100
			_go_fuzz_dep_.CoverTab[72264]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:100
			// _ = "end of CoverTab[72264]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:100
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:100
		// _ = "end of CoverTab[72252]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:100
		_go_fuzz_dep_.CoverTab[72253]++
													cc, err := call.res, call.err
													if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:102
			_go_fuzz_dep_.CoverTab[72265]++
														return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:103
			// _ = "end of CoverTab[72265]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:104
			_go_fuzz_dep_.CoverTab[72266]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:104
			// _ = "end of CoverTab[72266]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:104
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:104
		// _ = "end of CoverTab[72253]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:104
		_go_fuzz_dep_.CoverTab[72254]++
													if cc.ReserveNewRequest() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:105
			_go_fuzz_dep_.CoverTab[72267]++
														return cc, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:106
			// _ = "end of CoverTab[72267]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:107
			_go_fuzz_dep_.CoverTab[72268]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:107
			// _ = "end of CoverTab[72268]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:107
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:107
		// _ = "end of CoverTab[72254]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:108
	// _ = "end of CoverTab[72243]"
}

// dialCall is an in-flight Transport dial call to a host.
type dialCall struct {
	_	incomparable
	p	*clientConnPool
	// the context associated with the request
	// that created this dialCall
	ctx	context.Context
	done	chan struct{}	// closed when done
	res	*ClientConn	// valid after done is closed
	err	error		// valid after done is closed
}

// requires p.mu is held.
func (p *clientConnPool) getStartDialLocked(ctx context.Context, addr string) *dialCall {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:124
	_go_fuzz_dep_.CoverTab[72269]++
												if call, ok := p.dialing[addr]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:125
		_go_fuzz_dep_.CoverTab[72272]++

													return call
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:127
		// _ = "end of CoverTab[72272]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:128
		_go_fuzz_dep_.CoverTab[72273]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:128
		// _ = "end of CoverTab[72273]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:128
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:128
	// _ = "end of CoverTab[72269]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:128
	_go_fuzz_dep_.CoverTab[72270]++
												call := &dialCall{p: p, done: make(chan struct{}), ctx: ctx}
												if p.dialing == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:130
		_go_fuzz_dep_.CoverTab[72274]++
													p.dialing = make(map[string]*dialCall)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:131
		// _ = "end of CoverTab[72274]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:132
		_go_fuzz_dep_.CoverTab[72275]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:132
		// _ = "end of CoverTab[72275]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:132
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:132
	// _ = "end of CoverTab[72270]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:132
	_go_fuzz_dep_.CoverTab[72271]++
												p.dialing[addr] = call
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:133
	_curRoutineNum54_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:133
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum54_)
												go call.dial(call.ctx, addr)
												return call
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:135
	// _ = "end of CoverTab[72271]"
}

// run in its own goroutine.
func (c *dialCall) dial(ctx context.Context, addr string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:139
	_go_fuzz_dep_.CoverTab[72276]++
												const singleUse = false	// shared conn
												c.res, c.err = c.p.t.dialClientConn(ctx, addr, singleUse)

												c.p.mu.Lock()
												delete(c.p.dialing, addr)
												if c.err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:145
		_go_fuzz_dep_.CoverTab[72278]++
													c.p.addConnLocked(addr, c.res)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:146
		// _ = "end of CoverTab[72278]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:147
		_go_fuzz_dep_.CoverTab[72279]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:147
		// _ = "end of CoverTab[72279]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:147
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:147
	// _ = "end of CoverTab[72276]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:147
	_go_fuzz_dep_.CoverTab[72277]++
												c.p.mu.Unlock()

												close(c.done)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:150
	// _ = "end of CoverTab[72277]"
}

// addConnIfNeeded makes a NewClientConn out of c if a connection for key doesn't
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:153
// already exist. It coalesces concurrent calls with the same key.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:153
// This is used by the http1 Transport code when it creates a new connection. Because
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:153
// the http1 Transport doesn't de-dup TCP dials to outbound hosts (because it doesn't know
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:153
// the protocol), it can get into a situation where it has multiple TLS connections.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:153
// This code decides which ones live or die.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:153
// The return value used is whether c was used.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:153
// c is never closed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:161
func (p *clientConnPool) addConnIfNeeded(key string, t *Transport, c *tls.Conn) (used bool, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:161
	_go_fuzz_dep_.CoverTab[72280]++
												p.mu.Lock()
												for _, cc := range p.conns[key] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:163
		_go_fuzz_dep_.CoverTab[72284]++
													if cc.CanTakeNewRequest() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:164
			_go_fuzz_dep_.CoverTab[72285]++
														p.mu.Unlock()
														return false, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:166
			// _ = "end of CoverTab[72285]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:167
			_go_fuzz_dep_.CoverTab[72286]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:167
			// _ = "end of CoverTab[72286]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:167
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:167
		// _ = "end of CoverTab[72284]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:168
	// _ = "end of CoverTab[72280]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:168
	_go_fuzz_dep_.CoverTab[72281]++
												call, dup := p.addConnCalls[key]
												if !dup {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:170
		_go_fuzz_dep_.CoverTab[72287]++
													if p.addConnCalls == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:171
			_go_fuzz_dep_.CoverTab[72289]++
														p.addConnCalls = make(map[string]*addConnCall)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:172
			// _ = "end of CoverTab[72289]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:173
			_go_fuzz_dep_.CoverTab[72290]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:173
			// _ = "end of CoverTab[72290]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:173
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:173
		// _ = "end of CoverTab[72287]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:173
		_go_fuzz_dep_.CoverTab[72288]++
													call = &addConnCall{
			p:	p,
			done:	make(chan struct{}),
		}
													p.addConnCalls[key] = call
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:178
		_curRoutineNum55_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:178
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum55_)
													go call.run(t, key, c)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:179
		// _ = "end of CoverTab[72288]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:180
		_go_fuzz_dep_.CoverTab[72291]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:180
		// _ = "end of CoverTab[72291]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:180
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:180
	// _ = "end of CoverTab[72281]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:180
	_go_fuzz_dep_.CoverTab[72282]++
												p.mu.Unlock()

												<-call.done
												if call.err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:184
		_go_fuzz_dep_.CoverTab[72292]++
													return false, call.err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:185
		// _ = "end of CoverTab[72292]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:186
		_go_fuzz_dep_.CoverTab[72293]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:186
		// _ = "end of CoverTab[72293]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:186
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:186
	// _ = "end of CoverTab[72282]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:186
	_go_fuzz_dep_.CoverTab[72283]++
												return !dup, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:187
	// _ = "end of CoverTab[72283]"
}

type addConnCall struct {
	_	incomparable
	p	*clientConnPool
	done	chan struct{}	// closed when done
	err	error
}

func (c *addConnCall) run(t *Transport, key string, tc *tls.Conn) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:197
	_go_fuzz_dep_.CoverTab[72294]++
												cc, err := t.NewClientConn(tc)

												p := c.p
												p.mu.Lock()
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:202
		_go_fuzz_dep_.CoverTab[72296]++
													c.err = err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:203
		// _ = "end of CoverTab[72296]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:204
		_go_fuzz_dep_.CoverTab[72297]++
													cc.getConnCalled = true
													p.addConnLocked(key, cc)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:206
		// _ = "end of CoverTab[72297]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:207
	// _ = "end of CoverTab[72294]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:207
	_go_fuzz_dep_.CoverTab[72295]++
												delete(p.addConnCalls, key)
												p.mu.Unlock()
												close(c.done)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:210
	// _ = "end of CoverTab[72295]"
}

// p.mu must be held
func (p *clientConnPool) addConnLocked(key string, cc *ClientConn) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:214
	_go_fuzz_dep_.CoverTab[72298]++
												for _, v := range p.conns[key] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:215
		_go_fuzz_dep_.CoverTab[72302]++
													if v == cc {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:216
			_go_fuzz_dep_.CoverTab[72303]++
														return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:217
			// _ = "end of CoverTab[72303]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:218
			_go_fuzz_dep_.CoverTab[72304]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:218
			// _ = "end of CoverTab[72304]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:218
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:218
		// _ = "end of CoverTab[72302]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:219
	// _ = "end of CoverTab[72298]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:219
	_go_fuzz_dep_.CoverTab[72299]++
												if p.conns == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:220
		_go_fuzz_dep_.CoverTab[72305]++
													p.conns = make(map[string][]*ClientConn)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:221
		// _ = "end of CoverTab[72305]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:222
		_go_fuzz_dep_.CoverTab[72306]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:222
		// _ = "end of CoverTab[72306]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:222
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:222
	// _ = "end of CoverTab[72299]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:222
	_go_fuzz_dep_.CoverTab[72300]++
												if p.keys == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:223
		_go_fuzz_dep_.CoverTab[72307]++
													p.keys = make(map[*ClientConn][]string)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:224
		// _ = "end of CoverTab[72307]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:225
		_go_fuzz_dep_.CoverTab[72308]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:225
		// _ = "end of CoverTab[72308]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:225
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:225
	// _ = "end of CoverTab[72300]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:225
	_go_fuzz_dep_.CoverTab[72301]++
												p.conns[key] = append(p.conns[key], cc)
												p.keys[cc] = append(p.keys[cc], key)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:227
	// _ = "end of CoverTab[72301]"
}

func (p *clientConnPool) MarkDead(cc *ClientConn) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:230
	_go_fuzz_dep_.CoverTab[72309]++
												p.mu.Lock()
												defer p.mu.Unlock()
												for _, key := range p.keys[cc] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:233
		_go_fuzz_dep_.CoverTab[72311]++
													vv, ok := p.conns[key]
													if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:235
			_go_fuzz_dep_.CoverTab[72313]++
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:236
			// _ = "end of CoverTab[72313]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:237
			_go_fuzz_dep_.CoverTab[72314]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:237
			// _ = "end of CoverTab[72314]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:237
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:237
		// _ = "end of CoverTab[72311]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:237
		_go_fuzz_dep_.CoverTab[72312]++
													newList := filterOutClientConn(vv, cc)
													if len(newList) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:239
			_go_fuzz_dep_.CoverTab[72315]++
														p.conns[key] = newList
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:240
			// _ = "end of CoverTab[72315]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:241
			_go_fuzz_dep_.CoverTab[72316]++
														delete(p.conns, key)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:242
			// _ = "end of CoverTab[72316]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:243
		// _ = "end of CoverTab[72312]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:244
	// _ = "end of CoverTab[72309]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:244
	_go_fuzz_dep_.CoverTab[72310]++
												delete(p.keys, cc)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:245
	// _ = "end of CoverTab[72310]"
}

func (p *clientConnPool) closeIdleConnections() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:248
	_go_fuzz_dep_.CoverTab[72317]++
												p.mu.Lock()
												defer p.mu.Unlock()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:257
	for _, vv := range p.conns {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:257
		_go_fuzz_dep_.CoverTab[72318]++
													for _, cc := range vv {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:258
			_go_fuzz_dep_.CoverTab[72319]++
														cc.closeIfIdle()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:259
			// _ = "end of CoverTab[72319]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:260
		// _ = "end of CoverTab[72318]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:261
	// _ = "end of CoverTab[72317]"
}

func filterOutClientConn(in []*ClientConn, exclude *ClientConn) []*ClientConn {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:264
	_go_fuzz_dep_.CoverTab[72320]++
												out := in[:0]
												for _, v := range in {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:266
		_go_fuzz_dep_.CoverTab[72323]++
													if v != exclude {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:267
			_go_fuzz_dep_.CoverTab[72324]++
														out = append(out, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:268
			// _ = "end of CoverTab[72324]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:269
			_go_fuzz_dep_.CoverTab[72325]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:269
			// _ = "end of CoverTab[72325]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:269
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:269
		// _ = "end of CoverTab[72323]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:270
	// _ = "end of CoverTab[72320]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:270
	_go_fuzz_dep_.CoverTab[72321]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:273
	if len(in) != len(out) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:273
		_go_fuzz_dep_.CoverTab[72326]++
													in[len(in)-1] = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:274
		// _ = "end of CoverTab[72326]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:275
		_go_fuzz_dep_.CoverTab[72327]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:275
		// _ = "end of CoverTab[72327]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:275
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:275
	// _ = "end of CoverTab[72321]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:275
	_go_fuzz_dep_.CoverTab[72322]++
												return out
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:276
	// _ = "end of CoverTab[72322]"
}

// noDialClientConnPool is an implementation of http2.ClientConnPool
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:279
// which never dials. We let the HTTP/1.1 client dial and use its TLS
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:279
// connection instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:282
type noDialClientConnPool struct{ *clientConnPool }

func (p noDialClientConnPool) GetClientConn(req *http.Request, addr string) (*ClientConn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:284
	_go_fuzz_dep_.CoverTab[72328]++
												return p.getClientConn(req, addr, noDialOnMiss)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:285
	// _ = "end of CoverTab[72328]"
}

// shouldRetryDial reports whether the current request should
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:288
// retry dialing after the call finished unsuccessfully, for example
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:288
// if the dial was canceled because of a context cancellation or
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:288
// deadline expiry.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:292
func shouldRetryDial(call *dialCall, req *http.Request) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:292
	_go_fuzz_dep_.CoverTab[72329]++
												if call.err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:293
		_go_fuzz_dep_.CoverTab[72333]++

													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:295
		// _ = "end of CoverTab[72333]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:296
		_go_fuzz_dep_.CoverTab[72334]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:296
		// _ = "end of CoverTab[72334]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:296
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:296
	// _ = "end of CoverTab[72329]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:296
	_go_fuzz_dep_.CoverTab[72330]++
												if call.ctx == req.Context() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:297
		_go_fuzz_dep_.CoverTab[72335]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:301
		return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:301
		// _ = "end of CoverTab[72335]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:302
		_go_fuzz_dep_.CoverTab[72336]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:302
		// _ = "end of CoverTab[72336]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:302
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:302
	// _ = "end of CoverTab[72330]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:302
	_go_fuzz_dep_.CoverTab[72331]++
												if !errors.Is(call.err, context.Canceled) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:303
		_go_fuzz_dep_.CoverTab[72337]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:303
		return !errors.Is(call.err, context.DeadlineExceeded)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:303
		// _ = "end of CoverTab[72337]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:303
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:303
		_go_fuzz_dep_.CoverTab[72338]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:306
		return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:306
		// _ = "end of CoverTab[72338]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:307
		_go_fuzz_dep_.CoverTab[72339]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:307
		// _ = "end of CoverTab[72339]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:307
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:307
	// _ = "end of CoverTab[72331]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:307
	_go_fuzz_dep_.CoverTab[72332]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:310
	return call.ctx.Err() != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:310
	// _ = "end of CoverTab[72332]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:311
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/client_conn_pool.go:311
var _ = _go_fuzz_dep_.CoverTab
