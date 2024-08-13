// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/httputil/persist.go:5
package httputil

//line /usr/local/go/src/net/http/httputil/persist.go:5
import (
//line /usr/local/go/src/net/http/httputil/persist.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/httputil/persist.go:5
)
//line /usr/local/go/src/net/http/httputil/persist.go:5
import (
//line /usr/local/go/src/net/http/httputil/persist.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/httputil/persist.go:5
)

import (
	"bufio"
	"errors"
	"io"
	"net"
	"net/http"
	"net/textproto"
	"sync"
)

var (
	// Deprecated: No longer used.
	ErrPersistEOF	= &http.ProtocolError{ErrorString: "persistent connection closed"}

	// Deprecated: No longer used.
	ErrClosed	= &http.ProtocolError{ErrorString: "connection closed by user"}

	// Deprecated: No longer used.
	ErrPipeline	= &http.ProtocolError{ErrorString: "pipeline error"}
)

// This is an API usage error - the local side is closed.
//line /usr/local/go/src/net/http/httputil/persist.go:28
// ErrPersistEOF (above) reports that the remote side is closed.
//line /usr/local/go/src/net/http/httputil/persist.go:30
var errClosed = errors.New("i/o operation on closed connection")

// ServerConn is an artifact of Go's early HTTP implementation.
//line /usr/local/go/src/net/http/httputil/persist.go:32
// It is low-level, old, and unused by Go's current HTTP stack.
//line /usr/local/go/src/net/http/httputil/persist.go:32
// We should have deleted it before Go 1.
//line /usr/local/go/src/net/http/httputil/persist.go:32
//
//line /usr/local/go/src/net/http/httputil/persist.go:32
// Deprecated: Use the Server in package net/http instead.
//line /usr/local/go/src/net/http/httputil/persist.go:37
type ServerConn struct {
	mu		sync.Mutex	// read-write protects the following fields
	c		net.Conn
	r		*bufio.Reader
	re, we		error	// read/write errors
	lastbody	io.ReadCloser
	nread, nwritten	int
	pipereq		map[*http.Request]uint

	pipe	textproto.Pipeline
}

// NewServerConn is an artifact of Go's early HTTP implementation.
//line /usr/local/go/src/net/http/httputil/persist.go:49
// It is low-level, old, and unused by Go's current HTTP stack.
//line /usr/local/go/src/net/http/httputil/persist.go:49
// We should have deleted it before Go 1.
//line /usr/local/go/src/net/http/httputil/persist.go:49
//
//line /usr/local/go/src/net/http/httputil/persist.go:49
// Deprecated: Use the Server in package net/http instead.
//line /usr/local/go/src/net/http/httputil/persist.go:54
func NewServerConn(c net.Conn, r *bufio.Reader) *ServerConn {
//line /usr/local/go/src/net/http/httputil/persist.go:54
	_go_fuzz_dep_.CoverTab[76173]++
								if r == nil {
//line /usr/local/go/src/net/http/httputil/persist.go:55
		_go_fuzz_dep_.CoverTab[76175]++
									r = bufio.NewReader(c)
//line /usr/local/go/src/net/http/httputil/persist.go:56
		// _ = "end of CoverTab[76175]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:57
		_go_fuzz_dep_.CoverTab[76176]++
//line /usr/local/go/src/net/http/httputil/persist.go:57
		// _ = "end of CoverTab[76176]"
//line /usr/local/go/src/net/http/httputil/persist.go:57
	}
//line /usr/local/go/src/net/http/httputil/persist.go:57
	// _ = "end of CoverTab[76173]"
//line /usr/local/go/src/net/http/httputil/persist.go:57
	_go_fuzz_dep_.CoverTab[76174]++
								return &ServerConn{c: c, r: r, pipereq: make(map[*http.Request]uint)}
//line /usr/local/go/src/net/http/httputil/persist.go:58
	// _ = "end of CoverTab[76174]"
}

// Hijack detaches the ServerConn and returns the underlying connection as well
//line /usr/local/go/src/net/http/httputil/persist.go:61
// as the read-side bufio which may have some left over data. Hijack may be
//line /usr/local/go/src/net/http/httputil/persist.go:61
// called before Read has signaled the end of the keep-alive logic. The user
//line /usr/local/go/src/net/http/httputil/persist.go:61
// should not call Hijack while Read or Write is in progress.
//line /usr/local/go/src/net/http/httputil/persist.go:65
func (sc *ServerConn) Hijack() (net.Conn, *bufio.Reader) {
//line /usr/local/go/src/net/http/httputil/persist.go:65
	_go_fuzz_dep_.CoverTab[76177]++
								sc.mu.Lock()
								defer sc.mu.Unlock()
								c := sc.c
								r := sc.r
								sc.c = nil
								sc.r = nil
								return c, r
//line /usr/local/go/src/net/http/httputil/persist.go:72
	// _ = "end of CoverTab[76177]"
}

// Close calls Hijack and then also closes the underlying connection.
func (sc *ServerConn) Close() error {
//line /usr/local/go/src/net/http/httputil/persist.go:76
	_go_fuzz_dep_.CoverTab[76178]++
								c, _ := sc.Hijack()
								if c != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:78
		_go_fuzz_dep_.CoverTab[76180]++
									return c.Close()
//line /usr/local/go/src/net/http/httputil/persist.go:79
		// _ = "end of CoverTab[76180]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:80
		_go_fuzz_dep_.CoverTab[76181]++
//line /usr/local/go/src/net/http/httputil/persist.go:80
		// _ = "end of CoverTab[76181]"
//line /usr/local/go/src/net/http/httputil/persist.go:80
	}
//line /usr/local/go/src/net/http/httputil/persist.go:80
	// _ = "end of CoverTab[76178]"
//line /usr/local/go/src/net/http/httputil/persist.go:80
	_go_fuzz_dep_.CoverTab[76179]++
								return nil
//line /usr/local/go/src/net/http/httputil/persist.go:81
	// _ = "end of CoverTab[76179]"
}

// Read returns the next request on the wire. An ErrPersistEOF is returned if
//line /usr/local/go/src/net/http/httputil/persist.go:84
// it is gracefully determined that there are no more requests (e.g. after the
//line /usr/local/go/src/net/http/httputil/persist.go:84
// first request on an HTTP/1.0 connection, or after a Connection:close on a
//line /usr/local/go/src/net/http/httputil/persist.go:84
// HTTP/1.1 connection).
//line /usr/local/go/src/net/http/httputil/persist.go:88
func (sc *ServerConn) Read() (*http.Request, error) {
//line /usr/local/go/src/net/http/httputil/persist.go:88
	_go_fuzz_dep_.CoverTab[76182]++
								var req *http.Request
								var err error

//line /usr/local/go/src/net/http/httputil/persist.go:93
	id := sc.pipe.Next()
	sc.pipe.StartRequest(id)
	defer func() {
//line /usr/local/go/src/net/http/httputil/persist.go:95
		_go_fuzz_dep_.CoverTab[76190]++
									sc.pipe.EndRequest(id)
									if req == nil {
//line /usr/local/go/src/net/http/httputil/persist.go:97
			_go_fuzz_dep_.CoverTab[76191]++
										sc.pipe.StartResponse(id)
										sc.pipe.EndResponse(id)
//line /usr/local/go/src/net/http/httputil/persist.go:99
			// _ = "end of CoverTab[76191]"
		} else {
//line /usr/local/go/src/net/http/httputil/persist.go:100
			_go_fuzz_dep_.CoverTab[76192]++

										sc.mu.Lock()
										sc.pipereq[req] = id
										sc.mu.Unlock()
//line /usr/local/go/src/net/http/httputil/persist.go:104
			// _ = "end of CoverTab[76192]"
		}
//line /usr/local/go/src/net/http/httputil/persist.go:105
		// _ = "end of CoverTab[76190]"
	}()
//line /usr/local/go/src/net/http/httputil/persist.go:106
	// _ = "end of CoverTab[76182]"
//line /usr/local/go/src/net/http/httputil/persist.go:106
	_go_fuzz_dep_.CoverTab[76183]++

								sc.mu.Lock()
								if sc.we != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:109
		_go_fuzz_dep_.CoverTab[76193]++
									defer sc.mu.Unlock()
									return nil, sc.we
//line /usr/local/go/src/net/http/httputil/persist.go:111
		// _ = "end of CoverTab[76193]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:112
		_go_fuzz_dep_.CoverTab[76194]++
//line /usr/local/go/src/net/http/httputil/persist.go:112
		// _ = "end of CoverTab[76194]"
//line /usr/local/go/src/net/http/httputil/persist.go:112
	}
//line /usr/local/go/src/net/http/httputil/persist.go:112
	// _ = "end of CoverTab[76183]"
//line /usr/local/go/src/net/http/httputil/persist.go:112
	_go_fuzz_dep_.CoverTab[76184]++
								if sc.re != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:113
		_go_fuzz_dep_.CoverTab[76195]++
									defer sc.mu.Unlock()
									return nil, sc.re
//line /usr/local/go/src/net/http/httputil/persist.go:115
		// _ = "end of CoverTab[76195]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:116
		_go_fuzz_dep_.CoverTab[76196]++
//line /usr/local/go/src/net/http/httputil/persist.go:116
		// _ = "end of CoverTab[76196]"
//line /usr/local/go/src/net/http/httputil/persist.go:116
	}
//line /usr/local/go/src/net/http/httputil/persist.go:116
	// _ = "end of CoverTab[76184]"
//line /usr/local/go/src/net/http/httputil/persist.go:116
	_go_fuzz_dep_.CoverTab[76185]++
								if sc.r == nil {
//line /usr/local/go/src/net/http/httputil/persist.go:117
		_go_fuzz_dep_.CoverTab[76197]++
									defer sc.mu.Unlock()
									return nil, errClosed
//line /usr/local/go/src/net/http/httputil/persist.go:119
		// _ = "end of CoverTab[76197]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:120
		_go_fuzz_dep_.CoverTab[76198]++
//line /usr/local/go/src/net/http/httputil/persist.go:120
		// _ = "end of CoverTab[76198]"
//line /usr/local/go/src/net/http/httputil/persist.go:120
	}
//line /usr/local/go/src/net/http/httputil/persist.go:120
	// _ = "end of CoverTab[76185]"
//line /usr/local/go/src/net/http/httputil/persist.go:120
	_go_fuzz_dep_.CoverTab[76186]++
								r := sc.r
								lastbody := sc.lastbody
								sc.lastbody = nil
								sc.mu.Unlock()

//line /usr/local/go/src/net/http/httputil/persist.go:127
	if lastbody != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:127
		_go_fuzz_dep_.CoverTab[76199]++

//line /usr/local/go/src/net/http/httputil/persist.go:131
		err = lastbody.Close()
		if err != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:132
			_go_fuzz_dep_.CoverTab[76200]++
										sc.mu.Lock()
										defer sc.mu.Unlock()
										sc.re = err
										return nil, err
//line /usr/local/go/src/net/http/httputil/persist.go:136
			// _ = "end of CoverTab[76200]"
		} else {
//line /usr/local/go/src/net/http/httputil/persist.go:137
			_go_fuzz_dep_.CoverTab[76201]++
//line /usr/local/go/src/net/http/httputil/persist.go:137
			// _ = "end of CoverTab[76201]"
//line /usr/local/go/src/net/http/httputil/persist.go:137
		}
//line /usr/local/go/src/net/http/httputil/persist.go:137
		// _ = "end of CoverTab[76199]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:138
		_go_fuzz_dep_.CoverTab[76202]++
//line /usr/local/go/src/net/http/httputil/persist.go:138
		// _ = "end of CoverTab[76202]"
//line /usr/local/go/src/net/http/httputil/persist.go:138
	}
//line /usr/local/go/src/net/http/httputil/persist.go:138
	// _ = "end of CoverTab[76186]"
//line /usr/local/go/src/net/http/httputil/persist.go:138
	_go_fuzz_dep_.CoverTab[76187]++

								req, err = http.ReadRequest(r)
								sc.mu.Lock()
								defer sc.mu.Unlock()
								if err != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:143
		_go_fuzz_dep_.CoverTab[76203]++
									if err == io.ErrUnexpectedEOF {
//line /usr/local/go/src/net/http/httputil/persist.go:144
			_go_fuzz_dep_.CoverTab[76204]++

//line /usr/local/go/src/net/http/httputil/persist.go:148
			sc.re = ErrPersistEOF
										return nil, sc.re
//line /usr/local/go/src/net/http/httputil/persist.go:149
			// _ = "end of CoverTab[76204]"
		} else {
//line /usr/local/go/src/net/http/httputil/persist.go:150
			_go_fuzz_dep_.CoverTab[76205]++
										sc.re = err
										return req, err
//line /usr/local/go/src/net/http/httputil/persist.go:152
			// _ = "end of CoverTab[76205]"
		}
//line /usr/local/go/src/net/http/httputil/persist.go:153
		// _ = "end of CoverTab[76203]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:154
		_go_fuzz_dep_.CoverTab[76206]++
//line /usr/local/go/src/net/http/httputil/persist.go:154
		// _ = "end of CoverTab[76206]"
//line /usr/local/go/src/net/http/httputil/persist.go:154
	}
//line /usr/local/go/src/net/http/httputil/persist.go:154
	// _ = "end of CoverTab[76187]"
//line /usr/local/go/src/net/http/httputil/persist.go:154
	_go_fuzz_dep_.CoverTab[76188]++
								sc.lastbody = req.Body
								sc.nread++
								if req.Close {
//line /usr/local/go/src/net/http/httputil/persist.go:157
		_go_fuzz_dep_.CoverTab[76207]++
									sc.re = ErrPersistEOF
									return req, sc.re
//line /usr/local/go/src/net/http/httputil/persist.go:159
		// _ = "end of CoverTab[76207]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:160
		_go_fuzz_dep_.CoverTab[76208]++
//line /usr/local/go/src/net/http/httputil/persist.go:160
		// _ = "end of CoverTab[76208]"
//line /usr/local/go/src/net/http/httputil/persist.go:160
	}
//line /usr/local/go/src/net/http/httputil/persist.go:160
	// _ = "end of CoverTab[76188]"
//line /usr/local/go/src/net/http/httputil/persist.go:160
	_go_fuzz_dep_.CoverTab[76189]++
								return req, err
//line /usr/local/go/src/net/http/httputil/persist.go:161
	// _ = "end of CoverTab[76189]"
}

// Pending returns the number of unanswered requests
//line /usr/local/go/src/net/http/httputil/persist.go:164
// that have been received on the connection.
//line /usr/local/go/src/net/http/httputil/persist.go:166
func (sc *ServerConn) Pending() int {
//line /usr/local/go/src/net/http/httputil/persist.go:166
	_go_fuzz_dep_.CoverTab[76209]++
								sc.mu.Lock()
								defer sc.mu.Unlock()
								return sc.nread - sc.nwritten
//line /usr/local/go/src/net/http/httputil/persist.go:169
	// _ = "end of CoverTab[76209]"
}

// Write writes resp in response to req. To close the connection gracefully, set the
//line /usr/local/go/src/net/http/httputil/persist.go:172
// Response.Close field to true. Write should be considered operational until
//line /usr/local/go/src/net/http/httputil/persist.go:172
// it returns an error, regardless of any errors returned on the Read side.
//line /usr/local/go/src/net/http/httputil/persist.go:175
func (sc *ServerConn) Write(req *http.Request, resp *http.Response) error {
//line /usr/local/go/src/net/http/httputil/persist.go:175
	_go_fuzz_dep_.CoverTab[76210]++

//line /usr/local/go/src/net/http/httputil/persist.go:178
	sc.mu.Lock()
	id, ok := sc.pipereq[req]
	delete(sc.pipereq, req)
	if !ok {
//line /usr/local/go/src/net/http/httputil/persist.go:181
		_go_fuzz_dep_.CoverTab[76217]++
									sc.mu.Unlock()
									return ErrPipeline
//line /usr/local/go/src/net/http/httputil/persist.go:183
		// _ = "end of CoverTab[76217]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:184
		_go_fuzz_dep_.CoverTab[76218]++
//line /usr/local/go/src/net/http/httputil/persist.go:184
		// _ = "end of CoverTab[76218]"
//line /usr/local/go/src/net/http/httputil/persist.go:184
	}
//line /usr/local/go/src/net/http/httputil/persist.go:184
	// _ = "end of CoverTab[76210]"
//line /usr/local/go/src/net/http/httputil/persist.go:184
	_go_fuzz_dep_.CoverTab[76211]++
								sc.mu.Unlock()

//line /usr/local/go/src/net/http/httputil/persist.go:188
	sc.pipe.StartResponse(id)
	defer sc.pipe.EndResponse(id)

	sc.mu.Lock()
	if sc.we != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:192
		_go_fuzz_dep_.CoverTab[76219]++
									defer sc.mu.Unlock()
									return sc.we
//line /usr/local/go/src/net/http/httputil/persist.go:194
		// _ = "end of CoverTab[76219]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:195
		_go_fuzz_dep_.CoverTab[76220]++
//line /usr/local/go/src/net/http/httputil/persist.go:195
		// _ = "end of CoverTab[76220]"
//line /usr/local/go/src/net/http/httputil/persist.go:195
	}
//line /usr/local/go/src/net/http/httputil/persist.go:195
	// _ = "end of CoverTab[76211]"
//line /usr/local/go/src/net/http/httputil/persist.go:195
	_go_fuzz_dep_.CoverTab[76212]++
								if sc.c == nil {
//line /usr/local/go/src/net/http/httputil/persist.go:196
		_go_fuzz_dep_.CoverTab[76221]++
									defer sc.mu.Unlock()
									return ErrClosed
//line /usr/local/go/src/net/http/httputil/persist.go:198
		// _ = "end of CoverTab[76221]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:199
		_go_fuzz_dep_.CoverTab[76222]++
//line /usr/local/go/src/net/http/httputil/persist.go:199
		// _ = "end of CoverTab[76222]"
//line /usr/local/go/src/net/http/httputil/persist.go:199
	}
//line /usr/local/go/src/net/http/httputil/persist.go:199
	// _ = "end of CoverTab[76212]"
//line /usr/local/go/src/net/http/httputil/persist.go:199
	_go_fuzz_dep_.CoverTab[76213]++
								c := sc.c
								if sc.nread <= sc.nwritten {
//line /usr/local/go/src/net/http/httputil/persist.go:201
		_go_fuzz_dep_.CoverTab[76223]++
									defer sc.mu.Unlock()
									return errors.New("persist server pipe count")
//line /usr/local/go/src/net/http/httputil/persist.go:203
		// _ = "end of CoverTab[76223]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:204
		_go_fuzz_dep_.CoverTab[76224]++
//line /usr/local/go/src/net/http/httputil/persist.go:204
		// _ = "end of CoverTab[76224]"
//line /usr/local/go/src/net/http/httputil/persist.go:204
	}
//line /usr/local/go/src/net/http/httputil/persist.go:204
	// _ = "end of CoverTab[76213]"
//line /usr/local/go/src/net/http/httputil/persist.go:204
	_go_fuzz_dep_.CoverTab[76214]++
								if resp.Close {
//line /usr/local/go/src/net/http/httputil/persist.go:205
		_go_fuzz_dep_.CoverTab[76225]++

//line /usr/local/go/src/net/http/httputil/persist.go:209
		sc.re = ErrPersistEOF
//line /usr/local/go/src/net/http/httputil/persist.go:209
		// _ = "end of CoverTab[76225]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:210
		_go_fuzz_dep_.CoverTab[76226]++
//line /usr/local/go/src/net/http/httputil/persist.go:210
		// _ = "end of CoverTab[76226]"
//line /usr/local/go/src/net/http/httputil/persist.go:210
	}
//line /usr/local/go/src/net/http/httputil/persist.go:210
	// _ = "end of CoverTab[76214]"
//line /usr/local/go/src/net/http/httputil/persist.go:210
	_go_fuzz_dep_.CoverTab[76215]++
								sc.mu.Unlock()

								err := resp.Write(c)
								sc.mu.Lock()
								defer sc.mu.Unlock()
								if err != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:216
		_go_fuzz_dep_.CoverTab[76227]++
									sc.we = err
									return err
//line /usr/local/go/src/net/http/httputil/persist.go:218
		// _ = "end of CoverTab[76227]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:219
		_go_fuzz_dep_.CoverTab[76228]++
//line /usr/local/go/src/net/http/httputil/persist.go:219
		// _ = "end of CoverTab[76228]"
//line /usr/local/go/src/net/http/httputil/persist.go:219
	}
//line /usr/local/go/src/net/http/httputil/persist.go:219
	// _ = "end of CoverTab[76215]"
//line /usr/local/go/src/net/http/httputil/persist.go:219
	_go_fuzz_dep_.CoverTab[76216]++
								sc.nwritten++

								return nil
//line /usr/local/go/src/net/http/httputil/persist.go:222
	// _ = "end of CoverTab[76216]"
}

// ClientConn is an artifact of Go's early HTTP implementation.
//line /usr/local/go/src/net/http/httputil/persist.go:225
// It is low-level, old, and unused by Go's current HTTP stack.
//line /usr/local/go/src/net/http/httputil/persist.go:225
// We should have deleted it before Go 1.
//line /usr/local/go/src/net/http/httputil/persist.go:225
//
//line /usr/local/go/src/net/http/httputil/persist.go:225
// Deprecated: Use Client or Transport in package net/http instead.
//line /usr/local/go/src/net/http/httputil/persist.go:230
type ClientConn struct {
	mu		sync.Mutex	// read-write protects the following fields
	c		net.Conn
	r		*bufio.Reader
	re, we		error	// read/write errors
	lastbody	io.ReadCloser
	nread, nwritten	int
	pipereq		map[*http.Request]uint

	pipe		textproto.Pipeline
	writeReq	func(*http.Request, io.Writer) error
}

// NewClientConn is an artifact of Go's early HTTP implementation.
//line /usr/local/go/src/net/http/httputil/persist.go:243
// It is low-level, old, and unused by Go's current HTTP stack.
//line /usr/local/go/src/net/http/httputil/persist.go:243
// We should have deleted it before Go 1.
//line /usr/local/go/src/net/http/httputil/persist.go:243
//
//line /usr/local/go/src/net/http/httputil/persist.go:243
// Deprecated: Use the Client or Transport in package net/http instead.
//line /usr/local/go/src/net/http/httputil/persist.go:248
func NewClientConn(c net.Conn, r *bufio.Reader) *ClientConn {
//line /usr/local/go/src/net/http/httputil/persist.go:248
	_go_fuzz_dep_.CoverTab[76229]++
								if r == nil {
//line /usr/local/go/src/net/http/httputil/persist.go:249
		_go_fuzz_dep_.CoverTab[76231]++
									r = bufio.NewReader(c)
//line /usr/local/go/src/net/http/httputil/persist.go:250
		// _ = "end of CoverTab[76231]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:251
		_go_fuzz_dep_.CoverTab[76232]++
//line /usr/local/go/src/net/http/httputil/persist.go:251
		// _ = "end of CoverTab[76232]"
//line /usr/local/go/src/net/http/httputil/persist.go:251
	}
//line /usr/local/go/src/net/http/httputil/persist.go:251
	// _ = "end of CoverTab[76229]"
//line /usr/local/go/src/net/http/httputil/persist.go:251
	_go_fuzz_dep_.CoverTab[76230]++
								return &ClientConn{
		c:		c,
		r:		r,
		pipereq:	make(map[*http.Request]uint),
		writeReq:	(*http.Request).Write,
	}
//line /usr/local/go/src/net/http/httputil/persist.go:257
	// _ = "end of CoverTab[76230]"
}

// NewProxyClientConn is an artifact of Go's early HTTP implementation.
//line /usr/local/go/src/net/http/httputil/persist.go:260
// It is low-level, old, and unused by Go's current HTTP stack.
//line /usr/local/go/src/net/http/httputil/persist.go:260
// We should have deleted it before Go 1.
//line /usr/local/go/src/net/http/httputil/persist.go:260
//
//line /usr/local/go/src/net/http/httputil/persist.go:260
// Deprecated: Use the Client or Transport in package net/http instead.
//line /usr/local/go/src/net/http/httputil/persist.go:265
func NewProxyClientConn(c net.Conn, r *bufio.Reader) *ClientConn {
//line /usr/local/go/src/net/http/httputil/persist.go:265
	_go_fuzz_dep_.CoverTab[76233]++
								cc := NewClientConn(c, r)
								cc.writeReq = (*http.Request).WriteProxy
								return cc
//line /usr/local/go/src/net/http/httputil/persist.go:268
	// _ = "end of CoverTab[76233]"
}

// Hijack detaches the ClientConn and returns the underlying connection as well
//line /usr/local/go/src/net/http/httputil/persist.go:271
// as the read-side bufio which may have some left over data. Hijack may be
//line /usr/local/go/src/net/http/httputil/persist.go:271
// called before the user or Read have signaled the end of the keep-alive
//line /usr/local/go/src/net/http/httputil/persist.go:271
// logic. The user should not call Hijack while Read or Write is in progress.
//line /usr/local/go/src/net/http/httputil/persist.go:275
func (cc *ClientConn) Hijack() (c net.Conn, r *bufio.Reader) {
//line /usr/local/go/src/net/http/httputil/persist.go:275
	_go_fuzz_dep_.CoverTab[76234]++
								cc.mu.Lock()
								defer cc.mu.Unlock()
								c = cc.c
								r = cc.r
								cc.c = nil
								cc.r = nil
								return
//line /usr/local/go/src/net/http/httputil/persist.go:282
	// _ = "end of CoverTab[76234]"
}

// Close calls Hijack and then also closes the underlying connection.
func (cc *ClientConn) Close() error {
//line /usr/local/go/src/net/http/httputil/persist.go:286
	_go_fuzz_dep_.CoverTab[76235]++
								c, _ := cc.Hijack()
								if c != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:288
		_go_fuzz_dep_.CoverTab[76237]++
									return c.Close()
//line /usr/local/go/src/net/http/httputil/persist.go:289
		// _ = "end of CoverTab[76237]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:290
		_go_fuzz_dep_.CoverTab[76238]++
//line /usr/local/go/src/net/http/httputil/persist.go:290
		// _ = "end of CoverTab[76238]"
//line /usr/local/go/src/net/http/httputil/persist.go:290
	}
//line /usr/local/go/src/net/http/httputil/persist.go:290
	// _ = "end of CoverTab[76235]"
//line /usr/local/go/src/net/http/httputil/persist.go:290
	_go_fuzz_dep_.CoverTab[76236]++
								return nil
//line /usr/local/go/src/net/http/httputil/persist.go:291
	// _ = "end of CoverTab[76236]"
}

// Write writes a request. An ErrPersistEOF error is returned if the connection
//line /usr/local/go/src/net/http/httputil/persist.go:294
// has been closed in an HTTP keep-alive sense. If req.Close equals true, the
//line /usr/local/go/src/net/http/httputil/persist.go:294
// keep-alive connection is logically closed after this request and the opposing
//line /usr/local/go/src/net/http/httputil/persist.go:294
// server is informed. An ErrUnexpectedEOF indicates the remote closed the
//line /usr/local/go/src/net/http/httputil/persist.go:294
// underlying TCP connection, which is usually considered as graceful close.
//line /usr/local/go/src/net/http/httputil/persist.go:299
func (cc *ClientConn) Write(req *http.Request) error {
//line /usr/local/go/src/net/http/httputil/persist.go:299
	_go_fuzz_dep_.CoverTab[76239]++
								var err error

//line /usr/local/go/src/net/http/httputil/persist.go:303
	id := cc.pipe.Next()
	cc.pipe.StartRequest(id)
	defer func() {
//line /usr/local/go/src/net/http/httputil/persist.go:305
		_go_fuzz_dep_.CoverTab[76246]++
									cc.pipe.EndRequest(id)
									if err != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:307
			_go_fuzz_dep_.CoverTab[76247]++
										cc.pipe.StartResponse(id)
										cc.pipe.EndResponse(id)
//line /usr/local/go/src/net/http/httputil/persist.go:309
			// _ = "end of CoverTab[76247]"
		} else {
//line /usr/local/go/src/net/http/httputil/persist.go:310
			_go_fuzz_dep_.CoverTab[76248]++

										cc.mu.Lock()
										cc.pipereq[req] = id
										cc.mu.Unlock()
//line /usr/local/go/src/net/http/httputil/persist.go:314
			// _ = "end of CoverTab[76248]"
		}
//line /usr/local/go/src/net/http/httputil/persist.go:315
		// _ = "end of CoverTab[76246]"
	}()
//line /usr/local/go/src/net/http/httputil/persist.go:316
	// _ = "end of CoverTab[76239]"
//line /usr/local/go/src/net/http/httputil/persist.go:316
	_go_fuzz_dep_.CoverTab[76240]++

								cc.mu.Lock()
								if cc.re != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:319
		_go_fuzz_dep_.CoverTab[76249]++
									defer cc.mu.Unlock()
									return cc.re
//line /usr/local/go/src/net/http/httputil/persist.go:321
		// _ = "end of CoverTab[76249]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:322
		_go_fuzz_dep_.CoverTab[76250]++
//line /usr/local/go/src/net/http/httputil/persist.go:322
		// _ = "end of CoverTab[76250]"
//line /usr/local/go/src/net/http/httputil/persist.go:322
	}
//line /usr/local/go/src/net/http/httputil/persist.go:322
	// _ = "end of CoverTab[76240]"
//line /usr/local/go/src/net/http/httputil/persist.go:322
	_go_fuzz_dep_.CoverTab[76241]++
								if cc.we != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:323
		_go_fuzz_dep_.CoverTab[76251]++
									defer cc.mu.Unlock()
									return cc.we
//line /usr/local/go/src/net/http/httputil/persist.go:325
		// _ = "end of CoverTab[76251]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:326
		_go_fuzz_dep_.CoverTab[76252]++
//line /usr/local/go/src/net/http/httputil/persist.go:326
		// _ = "end of CoverTab[76252]"
//line /usr/local/go/src/net/http/httputil/persist.go:326
	}
//line /usr/local/go/src/net/http/httputil/persist.go:326
	// _ = "end of CoverTab[76241]"
//line /usr/local/go/src/net/http/httputil/persist.go:326
	_go_fuzz_dep_.CoverTab[76242]++
								if cc.c == nil {
//line /usr/local/go/src/net/http/httputil/persist.go:327
		_go_fuzz_dep_.CoverTab[76253]++
									defer cc.mu.Unlock()
									return errClosed
//line /usr/local/go/src/net/http/httputil/persist.go:329
		// _ = "end of CoverTab[76253]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:330
		_go_fuzz_dep_.CoverTab[76254]++
//line /usr/local/go/src/net/http/httputil/persist.go:330
		// _ = "end of CoverTab[76254]"
//line /usr/local/go/src/net/http/httputil/persist.go:330
	}
//line /usr/local/go/src/net/http/httputil/persist.go:330
	// _ = "end of CoverTab[76242]"
//line /usr/local/go/src/net/http/httputil/persist.go:330
	_go_fuzz_dep_.CoverTab[76243]++
								c := cc.c
								if req.Close {
//line /usr/local/go/src/net/http/httputil/persist.go:332
		_go_fuzz_dep_.CoverTab[76255]++

//line /usr/local/go/src/net/http/httputil/persist.go:335
		cc.we = ErrPersistEOF
//line /usr/local/go/src/net/http/httputil/persist.go:335
		// _ = "end of CoverTab[76255]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:336
		_go_fuzz_dep_.CoverTab[76256]++
//line /usr/local/go/src/net/http/httputil/persist.go:336
		// _ = "end of CoverTab[76256]"
//line /usr/local/go/src/net/http/httputil/persist.go:336
	}
//line /usr/local/go/src/net/http/httputil/persist.go:336
	// _ = "end of CoverTab[76243]"
//line /usr/local/go/src/net/http/httputil/persist.go:336
	_go_fuzz_dep_.CoverTab[76244]++
								cc.mu.Unlock()

								err = cc.writeReq(req, c)
								cc.mu.Lock()
								defer cc.mu.Unlock()
								if err != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:342
		_go_fuzz_dep_.CoverTab[76257]++
									cc.we = err
									return err
//line /usr/local/go/src/net/http/httputil/persist.go:344
		// _ = "end of CoverTab[76257]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:345
		_go_fuzz_dep_.CoverTab[76258]++
//line /usr/local/go/src/net/http/httputil/persist.go:345
		// _ = "end of CoverTab[76258]"
//line /usr/local/go/src/net/http/httputil/persist.go:345
	}
//line /usr/local/go/src/net/http/httputil/persist.go:345
	// _ = "end of CoverTab[76244]"
//line /usr/local/go/src/net/http/httputil/persist.go:345
	_go_fuzz_dep_.CoverTab[76245]++
								cc.nwritten++

								return nil
//line /usr/local/go/src/net/http/httputil/persist.go:348
	// _ = "end of CoverTab[76245]"
}

// Pending returns the number of unanswered requests
//line /usr/local/go/src/net/http/httputil/persist.go:351
// that have been sent on the connection.
//line /usr/local/go/src/net/http/httputil/persist.go:353
func (cc *ClientConn) Pending() int {
//line /usr/local/go/src/net/http/httputil/persist.go:353
	_go_fuzz_dep_.CoverTab[76259]++
								cc.mu.Lock()
								defer cc.mu.Unlock()
								return cc.nwritten - cc.nread
//line /usr/local/go/src/net/http/httputil/persist.go:356
	// _ = "end of CoverTab[76259]"
}

// Read reads the next response from the wire. A valid response might be
//line /usr/local/go/src/net/http/httputil/persist.go:359
// returned together with an ErrPersistEOF, which means that the remote
//line /usr/local/go/src/net/http/httputil/persist.go:359
// requested that this be the last request serviced. Read can be called
//line /usr/local/go/src/net/http/httputil/persist.go:359
// concurrently with Write, but not with another Read.
//line /usr/local/go/src/net/http/httputil/persist.go:363
func (cc *ClientConn) Read(req *http.Request) (resp *http.Response, err error) {
//line /usr/local/go/src/net/http/httputil/persist.go:363
	_go_fuzz_dep_.CoverTab[76260]++

								cc.mu.Lock()
								id, ok := cc.pipereq[req]
								delete(cc.pipereq, req)
								if !ok {
//line /usr/local/go/src/net/http/httputil/persist.go:368
		_go_fuzz_dep_.CoverTab[76267]++
									cc.mu.Unlock()
									return nil, ErrPipeline
//line /usr/local/go/src/net/http/httputil/persist.go:370
		// _ = "end of CoverTab[76267]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:371
		_go_fuzz_dep_.CoverTab[76268]++
//line /usr/local/go/src/net/http/httputil/persist.go:371
		// _ = "end of CoverTab[76268]"
//line /usr/local/go/src/net/http/httputil/persist.go:371
	}
//line /usr/local/go/src/net/http/httputil/persist.go:371
	// _ = "end of CoverTab[76260]"
//line /usr/local/go/src/net/http/httputil/persist.go:371
	_go_fuzz_dep_.CoverTab[76261]++
								cc.mu.Unlock()

//line /usr/local/go/src/net/http/httputil/persist.go:375
	cc.pipe.StartResponse(id)
	defer cc.pipe.EndResponse(id)

	cc.mu.Lock()
	if cc.re != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:379
		_go_fuzz_dep_.CoverTab[76269]++
									defer cc.mu.Unlock()
									return nil, cc.re
//line /usr/local/go/src/net/http/httputil/persist.go:381
		// _ = "end of CoverTab[76269]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:382
		_go_fuzz_dep_.CoverTab[76270]++
//line /usr/local/go/src/net/http/httputil/persist.go:382
		// _ = "end of CoverTab[76270]"
//line /usr/local/go/src/net/http/httputil/persist.go:382
	}
//line /usr/local/go/src/net/http/httputil/persist.go:382
	// _ = "end of CoverTab[76261]"
//line /usr/local/go/src/net/http/httputil/persist.go:382
	_go_fuzz_dep_.CoverTab[76262]++
								if cc.r == nil {
//line /usr/local/go/src/net/http/httputil/persist.go:383
		_go_fuzz_dep_.CoverTab[76271]++
									defer cc.mu.Unlock()
									return nil, errClosed
//line /usr/local/go/src/net/http/httputil/persist.go:385
		// _ = "end of CoverTab[76271]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:386
		_go_fuzz_dep_.CoverTab[76272]++
//line /usr/local/go/src/net/http/httputil/persist.go:386
		// _ = "end of CoverTab[76272]"
//line /usr/local/go/src/net/http/httputil/persist.go:386
	}
//line /usr/local/go/src/net/http/httputil/persist.go:386
	// _ = "end of CoverTab[76262]"
//line /usr/local/go/src/net/http/httputil/persist.go:386
	_go_fuzz_dep_.CoverTab[76263]++
								r := cc.r
								lastbody := cc.lastbody
								cc.lastbody = nil
								cc.mu.Unlock()

//line /usr/local/go/src/net/http/httputil/persist.go:393
	if lastbody != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:393
		_go_fuzz_dep_.CoverTab[76273]++

//line /usr/local/go/src/net/http/httputil/persist.go:397
		err = lastbody.Close()
		if err != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:398
			_go_fuzz_dep_.CoverTab[76274]++
										cc.mu.Lock()
										defer cc.mu.Unlock()
										cc.re = err
										return nil, err
//line /usr/local/go/src/net/http/httputil/persist.go:402
			// _ = "end of CoverTab[76274]"
		} else {
//line /usr/local/go/src/net/http/httputil/persist.go:403
			_go_fuzz_dep_.CoverTab[76275]++
//line /usr/local/go/src/net/http/httputil/persist.go:403
			// _ = "end of CoverTab[76275]"
//line /usr/local/go/src/net/http/httputil/persist.go:403
		}
//line /usr/local/go/src/net/http/httputil/persist.go:403
		// _ = "end of CoverTab[76273]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:404
		_go_fuzz_dep_.CoverTab[76276]++
//line /usr/local/go/src/net/http/httputil/persist.go:404
		// _ = "end of CoverTab[76276]"
//line /usr/local/go/src/net/http/httputil/persist.go:404
	}
//line /usr/local/go/src/net/http/httputil/persist.go:404
	// _ = "end of CoverTab[76263]"
//line /usr/local/go/src/net/http/httputil/persist.go:404
	_go_fuzz_dep_.CoverTab[76264]++

								resp, err = http.ReadResponse(r, req)
								cc.mu.Lock()
								defer cc.mu.Unlock()
								if err != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:409
		_go_fuzz_dep_.CoverTab[76277]++
									cc.re = err
									return resp, err
//line /usr/local/go/src/net/http/httputil/persist.go:411
		// _ = "end of CoverTab[76277]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:412
		_go_fuzz_dep_.CoverTab[76278]++
//line /usr/local/go/src/net/http/httputil/persist.go:412
		// _ = "end of CoverTab[76278]"
//line /usr/local/go/src/net/http/httputil/persist.go:412
	}
//line /usr/local/go/src/net/http/httputil/persist.go:412
	// _ = "end of CoverTab[76264]"
//line /usr/local/go/src/net/http/httputil/persist.go:412
	_go_fuzz_dep_.CoverTab[76265]++
								cc.lastbody = resp.Body

								cc.nread++

								if resp.Close {
//line /usr/local/go/src/net/http/httputil/persist.go:417
		_go_fuzz_dep_.CoverTab[76279]++
									cc.re = ErrPersistEOF
									return resp, cc.re
//line /usr/local/go/src/net/http/httputil/persist.go:419
		// _ = "end of CoverTab[76279]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:420
		_go_fuzz_dep_.CoverTab[76280]++
//line /usr/local/go/src/net/http/httputil/persist.go:420
		// _ = "end of CoverTab[76280]"
//line /usr/local/go/src/net/http/httputil/persist.go:420
	}
//line /usr/local/go/src/net/http/httputil/persist.go:420
	// _ = "end of CoverTab[76265]"
//line /usr/local/go/src/net/http/httputil/persist.go:420
	_go_fuzz_dep_.CoverTab[76266]++
								return resp, err
//line /usr/local/go/src/net/http/httputil/persist.go:421
	// _ = "end of CoverTab[76266]"
}

// Do is convenience method that writes a request and reads a response.
func (cc *ClientConn) Do(req *http.Request) (*http.Response, error) {
//line /usr/local/go/src/net/http/httputil/persist.go:425
	_go_fuzz_dep_.CoverTab[76281]++
								err := cc.Write(req)
								if err != nil {
//line /usr/local/go/src/net/http/httputil/persist.go:427
		_go_fuzz_dep_.CoverTab[76283]++
									return nil, err
//line /usr/local/go/src/net/http/httputil/persist.go:428
		// _ = "end of CoverTab[76283]"
	} else {
//line /usr/local/go/src/net/http/httputil/persist.go:429
		_go_fuzz_dep_.CoverTab[76284]++
//line /usr/local/go/src/net/http/httputil/persist.go:429
		// _ = "end of CoverTab[76284]"
//line /usr/local/go/src/net/http/httputil/persist.go:429
	}
//line /usr/local/go/src/net/http/httputil/persist.go:429
	// _ = "end of CoverTab[76281]"
//line /usr/local/go/src/net/http/httputil/persist.go:429
	_go_fuzz_dep_.CoverTab[76282]++
								return cc.Read(req)
//line /usr/local/go/src/net/http/httputil/persist.go:430
	// _ = "end of CoverTab[76282]"
}

//line /usr/local/go/src/net/http/httputil/persist.go:431
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/httputil/persist.go:431
var _ = _go_fuzz_dep_.CoverTab
