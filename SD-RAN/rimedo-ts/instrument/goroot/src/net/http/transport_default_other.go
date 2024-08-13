// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !(js && wasm)
// +build !js !wasm

//line /usr/local/go/src/net/http/transport_default_other.go:8
package http

//line /usr/local/go/src/net/http/transport_default_other.go:8
import (
//line /usr/local/go/src/net/http/transport_default_other.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/transport_default_other.go:8
)
//line /usr/local/go/src/net/http/transport_default_other.go:8
import (
//line /usr/local/go/src/net/http/transport_default_other.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/transport_default_other.go:8
)

import (
	"context"
	"net"
)

func defaultTransportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
//line /usr/local/go/src/net/http/transport_default_other.go:15
	_go_fuzz_dep_.CoverTab[45084]++
									return dialer.DialContext
//line /usr/local/go/src/net/http/transport_default_other.go:16
	// _ = "end of CoverTab[45084]"
}

//line /usr/local/go/src/net/http/transport_default_other.go:17
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/transport_default_other.go:17
var _ = _go_fuzz_dep_.CoverTab
