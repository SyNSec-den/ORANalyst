// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.15
// +build go1.15

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:8
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:8
)

import (
	"context"
	"crypto/tls"
)

// dialTLSWithContext uses tls.Dialer, added in Go 1.15, to open a TLS
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:15
// connection.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:17
func (t *Transport) dialTLSWithContext(ctx context.Context, network, addr string, cfg *tls.Config) (*tls.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:17
	_go_fuzz_dep_.CoverTab[72975]++
										dialer := &tls.Dialer{
		Config: cfg,
	}
	cn, err := dialer.DialContext(ctx, network, addr)
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:22
		_go_fuzz_dep_.CoverTab[72977]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:23
		// _ = "end of CoverTab[72977]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:24
		_go_fuzz_dep_.CoverTab[72978]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:24
		// _ = "end of CoverTab[72978]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:24
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:24
	// _ = "end of CoverTab[72975]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:24
	_go_fuzz_dep_.CoverTab[72976]++
										tlsCn := cn.(*tls.Conn)
										return tlsCn, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:26
	// _ = "end of CoverTab[72976]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:27
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go115.go:27
var _ = _go_fuzz_dep_.CoverTab
