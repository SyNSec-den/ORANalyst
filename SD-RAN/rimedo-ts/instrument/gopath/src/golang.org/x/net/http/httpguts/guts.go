// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:5
// Package httpguts provides functions implementing various details
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:5
// of the HTTP specification.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:5
// This package is shared by the standard library (which vendors it)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:5
// and x/net/http2. It comes with no API stability promise.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:10
package httpguts

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:10
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:10
)

import (
	"net/textproto"
	"strings"
)

// ValidTrailerHeader reports whether name is a valid header field name to appear
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:17
// in trailers.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:17
// See RFC 7230, Section 4.1.2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:20
func ValidTrailerHeader(name string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:20
	_go_fuzz_dep_.CoverTab[71784]++
											name = textproto.CanonicalMIMEHeaderKey(name)
											if strings.HasPrefix(name, "If-") || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:22
		_go_fuzz_dep_.CoverTab[71786]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:22
		return badTrailer[name]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:22
		// _ = "end of CoverTab[71786]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:22
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:22
		_go_fuzz_dep_.CoverTab[71787]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:23
		// _ = "end of CoverTab[71787]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:24
		_go_fuzz_dep_.CoverTab[71788]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:24
		// _ = "end of CoverTab[71788]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:24
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:24
	// _ = "end of CoverTab[71784]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:24
	_go_fuzz_dep_.CoverTab[71785]++
											return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:25
	// _ = "end of CoverTab[71785]"
}

var badTrailer = map[string]bool{
	"Authorization":	true,
	"Cache-Control":	true,
	"Connection":		true,
	"Content-Encoding":	true,
	"Content-Length":	true,
	"Content-Range":	true,
	"Content-Type":		true,
	"Expect":		true,
	"Host":			true,
	"Keep-Alive":		true,
	"Max-Forwards":		true,
	"Pragma":		true,
	"Proxy-Authenticate":	true,
	"Proxy-Authorization":	true,
	"Proxy-Connection":	true,
	"Range":		true,
	"Realm":		true,
	"Te":			true,
	"Trailer":		true,
	"Transfer-Encoding":	true,
	"Www-Authenticate":	true,
}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:50
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/guts.go:50
var _ = _go_fuzz_dep_.CoverTab
