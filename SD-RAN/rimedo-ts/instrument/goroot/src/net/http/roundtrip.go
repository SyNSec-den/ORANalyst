// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js || !wasm

//line /usr/local/go/src/net/http/roundtrip.go:7
package http

//line /usr/local/go/src/net/http/roundtrip.go:7
import (
//line /usr/local/go/src/net/http/roundtrip.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/roundtrip.go:7
)
//line /usr/local/go/src/net/http/roundtrip.go:7
import (
//line /usr/local/go/src/net/http/roundtrip.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/roundtrip.go:7
)

// RoundTrip implements the RoundTripper interface.
//line /usr/local/go/src/net/http/roundtrip.go:9
//
//line /usr/local/go/src/net/http/roundtrip.go:9
// For higher-level HTTP client support (such as handling of cookies
//line /usr/local/go/src/net/http/roundtrip.go:9
// and redirects), see Get, Post, and the Client type.
//line /usr/local/go/src/net/http/roundtrip.go:9
//
//line /usr/local/go/src/net/http/roundtrip.go:9
// Like the RoundTripper interface, the error types returned
//line /usr/local/go/src/net/http/roundtrip.go:9
// by RoundTrip are unspecified.
//line /usr/local/go/src/net/http/roundtrip.go:16
func (t *Transport) RoundTrip(req *Request) (*Response, error) {
//line /usr/local/go/src/net/http/roundtrip.go:16
	_go_fuzz_dep_.CoverTab[42098]++
							return t.roundTrip(req)
//line /usr/local/go/src/net/http/roundtrip.go:17
	// _ = "end of CoverTab[42098]"
}

//line /usr/local/go/src/net/http/roundtrip.go:18
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/roundtrip.go:18
var _ = _go_fuzz_dep_.CoverTab
