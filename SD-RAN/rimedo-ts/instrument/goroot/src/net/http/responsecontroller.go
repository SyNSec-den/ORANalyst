// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/responsecontroller.go:5
package http

//line /usr/local/go/src/net/http/responsecontroller.go:5
import (
//line /usr/local/go/src/net/http/responsecontroller.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/responsecontroller.go:5
)
//line /usr/local/go/src/net/http/responsecontroller.go:5
import (
//line /usr/local/go/src/net/http/responsecontroller.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/responsecontroller.go:5
)

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// A ResponseController is used by an HTTP handler to control the response.
//line /usr/local/go/src/net/http/responsecontroller.go:14
//
//line /usr/local/go/src/net/http/responsecontroller.go:14
// A ResponseController may not be used after the Handler.ServeHTTP method has returned.
//line /usr/local/go/src/net/http/responsecontroller.go:17
type ResponseController struct {
	rw ResponseWriter
}

// NewResponseController creates a ResponseController for a request.
//line /usr/local/go/src/net/http/responsecontroller.go:21
//
//line /usr/local/go/src/net/http/responsecontroller.go:21
// The ResponseWriter should be the original value passed to the Handler.ServeHTTP method,
//line /usr/local/go/src/net/http/responsecontroller.go:21
// or have an Unwrap method returning the original ResponseWriter.
//line /usr/local/go/src/net/http/responsecontroller.go:21
//
//line /usr/local/go/src/net/http/responsecontroller.go:21
// If the ResponseWriter implements any of the following methods, the ResponseController
//line /usr/local/go/src/net/http/responsecontroller.go:21
// will call them as appropriate:
//line /usr/local/go/src/net/http/responsecontroller.go:21
//
//line /usr/local/go/src/net/http/responsecontroller.go:21
//	Flush()
//line /usr/local/go/src/net/http/responsecontroller.go:21
//	FlushError() error // alternative Flush returning an error
//line /usr/local/go/src/net/http/responsecontroller.go:21
//	Hijack() (net.Conn, *bufio.ReadWriter, error)
//line /usr/local/go/src/net/http/responsecontroller.go:21
//	SetReadDeadline(deadline time.Time) error
//line /usr/local/go/src/net/http/responsecontroller.go:21
//	SetWriteDeadline(deadline time.Time) error
//line /usr/local/go/src/net/http/responsecontroller.go:21
//
//line /usr/local/go/src/net/http/responsecontroller.go:21
// If the ResponseWriter does not support a method, ResponseController returns
//line /usr/local/go/src/net/http/responsecontroller.go:21
// an error matching ErrNotSupported.
//line /usr/local/go/src/net/http/responsecontroller.go:37
func NewResponseController(rw ResponseWriter) *ResponseController {
//line /usr/local/go/src/net/http/responsecontroller.go:37
	_go_fuzz_dep_.CoverTab[42075]++
								return &ResponseController{rw}
//line /usr/local/go/src/net/http/responsecontroller.go:38
	// _ = "end of CoverTab[42075]"
}

type rwUnwrapper interface {
	Unwrap() ResponseWriter
}

// Flush flushes buffered data to the client.
func (c *ResponseController) Flush() error {
//line /usr/local/go/src/net/http/responsecontroller.go:46
	_go_fuzz_dep_.CoverTab[42076]++
								rw := c.rw
								for {
//line /usr/local/go/src/net/http/responsecontroller.go:48
		_go_fuzz_dep_.CoverTab[42077]++
									switch t := rw.(type) {
		case interface{ FlushError() error }:
//line /usr/local/go/src/net/http/responsecontroller.go:50
			_go_fuzz_dep_.CoverTab[42078]++
										return t.FlushError()
//line /usr/local/go/src/net/http/responsecontroller.go:51
			// _ = "end of CoverTab[42078]"
		case Flusher:
//line /usr/local/go/src/net/http/responsecontroller.go:52
			_go_fuzz_dep_.CoverTab[42079]++
										t.Flush()
										return nil
//line /usr/local/go/src/net/http/responsecontroller.go:54
			// _ = "end of CoverTab[42079]"
		case rwUnwrapper:
//line /usr/local/go/src/net/http/responsecontroller.go:55
			_go_fuzz_dep_.CoverTab[42080]++
										rw = t.Unwrap()
//line /usr/local/go/src/net/http/responsecontroller.go:56
			// _ = "end of CoverTab[42080]"
		default:
//line /usr/local/go/src/net/http/responsecontroller.go:57
			_go_fuzz_dep_.CoverTab[42081]++
										return errNotSupported()
//line /usr/local/go/src/net/http/responsecontroller.go:58
			// _ = "end of CoverTab[42081]"
		}
//line /usr/local/go/src/net/http/responsecontroller.go:59
		// _ = "end of CoverTab[42077]"
	}
//line /usr/local/go/src/net/http/responsecontroller.go:60
	// _ = "end of CoverTab[42076]"
}

// Hijack lets the caller take over the connection.
//line /usr/local/go/src/net/http/responsecontroller.go:63
// See the Hijacker interface for details.
//line /usr/local/go/src/net/http/responsecontroller.go:65
func (c *ResponseController) Hijack() (net.Conn, *bufio.ReadWriter, error) {
//line /usr/local/go/src/net/http/responsecontroller.go:65
	_go_fuzz_dep_.CoverTab[42082]++
								rw := c.rw
								for {
//line /usr/local/go/src/net/http/responsecontroller.go:67
		_go_fuzz_dep_.CoverTab[42083]++
									switch t := rw.(type) {
		case Hijacker:
//line /usr/local/go/src/net/http/responsecontroller.go:69
			_go_fuzz_dep_.CoverTab[42084]++
										return t.Hijack()
//line /usr/local/go/src/net/http/responsecontroller.go:70
			// _ = "end of CoverTab[42084]"
		case rwUnwrapper:
//line /usr/local/go/src/net/http/responsecontroller.go:71
			_go_fuzz_dep_.CoverTab[42085]++
										rw = t.Unwrap()
//line /usr/local/go/src/net/http/responsecontroller.go:72
			// _ = "end of CoverTab[42085]"
		default:
//line /usr/local/go/src/net/http/responsecontroller.go:73
			_go_fuzz_dep_.CoverTab[42086]++
										return nil, nil, errNotSupported()
//line /usr/local/go/src/net/http/responsecontroller.go:74
			// _ = "end of CoverTab[42086]"
		}
//line /usr/local/go/src/net/http/responsecontroller.go:75
		// _ = "end of CoverTab[42083]"
	}
//line /usr/local/go/src/net/http/responsecontroller.go:76
	// _ = "end of CoverTab[42082]"
}

// SetReadDeadline sets the deadline for reading the entire request, including the body.
//line /usr/local/go/src/net/http/responsecontroller.go:79
// Reads from the request body after the deadline has been exceeded will return an error.
//line /usr/local/go/src/net/http/responsecontroller.go:79
// A zero value means no deadline.
//line /usr/local/go/src/net/http/responsecontroller.go:79
//
//line /usr/local/go/src/net/http/responsecontroller.go:79
// Setting the read deadline after it has been exceeded will not extend it.
//line /usr/local/go/src/net/http/responsecontroller.go:84
func (c *ResponseController) SetReadDeadline(deadline time.Time) error {
//line /usr/local/go/src/net/http/responsecontroller.go:84
	_go_fuzz_dep_.CoverTab[42087]++
								rw := c.rw
								for {
//line /usr/local/go/src/net/http/responsecontroller.go:86
		_go_fuzz_dep_.CoverTab[42088]++
									switch t := rw.(type) {
		case interface{ SetReadDeadline(time.Time) error }:
//line /usr/local/go/src/net/http/responsecontroller.go:88
			_go_fuzz_dep_.CoverTab[42089]++
										return t.SetReadDeadline(deadline)
//line /usr/local/go/src/net/http/responsecontroller.go:89
			// _ = "end of CoverTab[42089]"
		case rwUnwrapper:
//line /usr/local/go/src/net/http/responsecontroller.go:90
			_go_fuzz_dep_.CoverTab[42090]++
										rw = t.Unwrap()
//line /usr/local/go/src/net/http/responsecontroller.go:91
			// _ = "end of CoverTab[42090]"
		default:
//line /usr/local/go/src/net/http/responsecontroller.go:92
			_go_fuzz_dep_.CoverTab[42091]++
										return errNotSupported()
//line /usr/local/go/src/net/http/responsecontroller.go:93
			// _ = "end of CoverTab[42091]"
		}
//line /usr/local/go/src/net/http/responsecontroller.go:94
		// _ = "end of CoverTab[42088]"
	}
//line /usr/local/go/src/net/http/responsecontroller.go:95
	// _ = "end of CoverTab[42087]"
}

// SetWriteDeadline sets the deadline for writing the response.
//line /usr/local/go/src/net/http/responsecontroller.go:98
// Writes to the response body after the deadline has been exceeded will not block,
//line /usr/local/go/src/net/http/responsecontroller.go:98
// but may succeed if the data has been buffered.
//line /usr/local/go/src/net/http/responsecontroller.go:98
// A zero value means no deadline.
//line /usr/local/go/src/net/http/responsecontroller.go:98
//
//line /usr/local/go/src/net/http/responsecontroller.go:98
// Setting the write deadline after it has been exceeded will not extend it.
//line /usr/local/go/src/net/http/responsecontroller.go:104
func (c *ResponseController) SetWriteDeadline(deadline time.Time) error {
//line /usr/local/go/src/net/http/responsecontroller.go:104
	_go_fuzz_dep_.CoverTab[42092]++
								rw := c.rw
								for {
//line /usr/local/go/src/net/http/responsecontroller.go:106
		_go_fuzz_dep_.CoverTab[42093]++
									switch t := rw.(type) {
		case interface{ SetWriteDeadline(time.Time) error }:
//line /usr/local/go/src/net/http/responsecontroller.go:108
			_go_fuzz_dep_.CoverTab[42094]++
										return t.SetWriteDeadline(deadline)
//line /usr/local/go/src/net/http/responsecontroller.go:109
			// _ = "end of CoverTab[42094]"
		case rwUnwrapper:
//line /usr/local/go/src/net/http/responsecontroller.go:110
			_go_fuzz_dep_.CoverTab[42095]++
										rw = t.Unwrap()
//line /usr/local/go/src/net/http/responsecontroller.go:111
			// _ = "end of CoverTab[42095]"
		default:
//line /usr/local/go/src/net/http/responsecontroller.go:112
			_go_fuzz_dep_.CoverTab[42096]++
										return errNotSupported()
//line /usr/local/go/src/net/http/responsecontroller.go:113
			// _ = "end of CoverTab[42096]"
		}
//line /usr/local/go/src/net/http/responsecontroller.go:114
		// _ = "end of CoverTab[42093]"
	}
//line /usr/local/go/src/net/http/responsecontroller.go:115
	// _ = "end of CoverTab[42092]"
}

// errNotSupported returns an error that Is ErrNotSupported,
//line /usr/local/go/src/net/http/responsecontroller.go:118
// but is not == to it.
//line /usr/local/go/src/net/http/responsecontroller.go:120
func errNotSupported() error {
//line /usr/local/go/src/net/http/responsecontroller.go:120
	_go_fuzz_dep_.CoverTab[42097]++
								return fmt.Errorf("%w", ErrNotSupported)
//line /usr/local/go/src/net/http/responsecontroller.go:121
	// _ = "end of CoverTab[42097]"
}

//line /usr/local/go/src/net/http/responsecontroller.go:122
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/responsecontroller.go:122
var _ = _go_fuzz_dep_.CoverTab
