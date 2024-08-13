// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:8
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:8
)

import (
	"crypto/tls"
	"net"
)

func tlsUnderlyingConn(tc *tls.Conn) net.Conn {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:15
	_go_fuzz_dep_.CoverTab[72979]++
										return tc.NetConn()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:16
	// _ = "end of CoverTab[72979]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:17
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go118.go:17
var _ = _go_fuzz_dep_.CoverTab
