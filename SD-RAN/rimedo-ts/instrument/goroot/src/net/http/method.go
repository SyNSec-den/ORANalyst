// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/method.go:5
package http

//line /usr/local/go/src/net/http/method.go:5
import (
//line /usr/local/go/src/net/http/method.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/method.go:5
)
//line /usr/local/go/src/net/http/method.go:5
import (
//line /usr/local/go/src/net/http/method.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/method.go:5
)

// Common HTTP methods.
//line /usr/local/go/src/net/http/method.go:7
//
//line /usr/local/go/src/net/http/method.go:7
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
//line /usr/local/go/src/net/http/method.go:10
const (
	MethodGet	= "GET"
	MethodHead	= "HEAD"
	MethodPost	= "POST"
	MethodPut	= "PUT"
	MethodPatch	= "PATCH"	// RFC 5789
	MethodDelete	= "DELETE"
	MethodConnect	= "CONNECT"
	MethodOptions	= "OPTIONS"
	MethodTrace	= "TRACE"
)

//line /usr/local/go/src/net/http/method.go:20
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/method.go:20
var _ = _go_fuzz_dep_.CoverTab
