// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/httputil/httputil.go:5
// Package httputil provides HTTP utility functions, complementing the
//line /usr/local/go/src/net/http/httputil/httputil.go:5
// more common ones in the net/http package.
//line /usr/local/go/src/net/http/httputil/httputil.go:7
package httputil

//line /usr/local/go/src/net/http/httputil/httputil.go:7
import (
//line /usr/local/go/src/net/http/httputil/httputil.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/httputil/httputil.go:7
)
//line /usr/local/go/src/net/http/httputil/httputil.go:7
import (
//line /usr/local/go/src/net/http/httputil/httputil.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/httputil/httputil.go:7
)

import (
	"io"
	"net/http/internal"
)

// NewChunkedReader returns a new chunkedReader that translates the data read from r
//line /usr/local/go/src/net/http/httputil/httputil.go:14
// out of HTTP "chunked" format before returning it.
//line /usr/local/go/src/net/http/httputil/httputil.go:14
// The chunkedReader returns io.EOF when the final 0-length chunk is read.
//line /usr/local/go/src/net/http/httputil/httputil.go:14
//
//line /usr/local/go/src/net/http/httputil/httputil.go:14
// NewChunkedReader is not needed by normal applications. The http package
//line /usr/local/go/src/net/http/httputil/httputil.go:14
// automatically decodes chunking when reading response bodies.
//line /usr/local/go/src/net/http/httputil/httputil.go:20
func NewChunkedReader(r io.Reader) io.Reader {
//line /usr/local/go/src/net/http/httputil/httputil.go:20
	_go_fuzz_dep_.CoverTab[76171]++
								return internal.NewChunkedReader(r)
//line /usr/local/go/src/net/http/httputil/httputil.go:21
	// _ = "end of CoverTab[76171]"
}

// NewChunkedWriter returns a new chunkedWriter that translates writes into HTTP
//line /usr/local/go/src/net/http/httputil/httputil.go:24
// "chunked" format before writing them to w. Closing the returned chunkedWriter
//line /usr/local/go/src/net/http/httputil/httputil.go:24
// sends the final 0-length chunk that marks the end of the stream but does
//line /usr/local/go/src/net/http/httputil/httputil.go:24
// not send the final CRLF that appears after trailers; trailers and the last
//line /usr/local/go/src/net/http/httputil/httputil.go:24
// CRLF must be written separately.
//line /usr/local/go/src/net/http/httputil/httputil.go:24
//
//line /usr/local/go/src/net/http/httputil/httputil.go:24
// NewChunkedWriter is not needed by normal applications. The http
//line /usr/local/go/src/net/http/httputil/httputil.go:24
// package adds chunking automatically if handlers don't set a
//line /usr/local/go/src/net/http/httputil/httputil.go:24
// Content-Length header. Using NewChunkedWriter inside a handler
//line /usr/local/go/src/net/http/httputil/httputil.go:24
// would result in double chunking or chunking with a Content-Length
//line /usr/local/go/src/net/http/httputil/httputil.go:24
// length, both of which are wrong.
//line /usr/local/go/src/net/http/httputil/httputil.go:35
func NewChunkedWriter(w io.Writer) io.WriteCloser {
//line /usr/local/go/src/net/http/httputil/httputil.go:35
	_go_fuzz_dep_.CoverTab[76172]++
								return internal.NewChunkedWriter(w)
//line /usr/local/go/src/net/http/httputil/httputil.go:36
	// _ = "end of CoverTab[76172]"
}

// ErrLineTooLong is returned when reading malformed chunked data
//line /usr/local/go/src/net/http/httputil/httputil.go:39
// with lines that are too long.
//line /usr/local/go/src/net/http/httputil/httputil.go:41
var ErrLineTooLong = internal.ErrLineTooLong
//line /usr/local/go/src/net/http/httputil/httputil.go:41
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/httputil/httputil.go:41
var _ = _go_fuzz_dep_.CoverTab
