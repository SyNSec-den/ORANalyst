// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:5
// Package httpguts provides functions implementing various details
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:5
// of the HTTP specification.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:5
// This package is shared by the standard library (which vendors it)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:5
// and x/net/http2. It comes with no API stability promise.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:10
package httpguts

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:10
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:10
)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:10
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:10
)

import (
	"net/textproto"
	"strings"
)

// ValidTrailerHeader reports whether name is a valid header field name to appear
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:17
// in trailers.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:17
// See RFC 7230, Section 4.1.2
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:20
func ValidTrailerHeader(name string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:20
	_go_fuzz_dep_.CoverTab[34909]++
										name = textproto.CanonicalMIMEHeaderKey(name)
										if strings.HasPrefix(name, "If-") || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:22
		_go_fuzz_dep_.CoverTab[34911]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:22
		return badTrailer[name]
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:22
		// _ = "end of CoverTab[34911]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:22
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:22
		_go_fuzz_dep_.CoverTab[34912]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:23
		// _ = "end of CoverTab[34912]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:24
		_go_fuzz_dep_.CoverTab[34913]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:24
		// _ = "end of CoverTab[34913]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:24
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:24
	// _ = "end of CoverTab[34909]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:24
	_go_fuzz_dep_.CoverTab[34910]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:25
	// _ = "end of CoverTab[34910]"
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
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:50
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/guts.go:50
var _ = _go_fuzz_dep_.CoverTab
