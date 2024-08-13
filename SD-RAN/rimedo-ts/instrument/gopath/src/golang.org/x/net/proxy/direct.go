// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:5
package proxy

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:5
)

import (
	"context"
	"net"
)

type direct struct{}

// Direct implements Dialer by making network connections directly using net.Dial or net.DialContext.
var Direct = direct{}

var (
	_	Dialer		= Direct
	_	ContextDialer	= Direct
)

// Dial directly invokes net.Dial with the supplied parameters.
func (direct) Dial(network, addr string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:23
	_go_fuzz_dep_.CoverTab[96947]++
										return net.Dial(network, addr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:24
	// _ = "end of CoverTab[96947]"
}

// DialContext instantiates a net.Dialer and invokes its DialContext receiver with the supplied parameters.
func (direct) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:28
	_go_fuzz_dep_.CoverTab[96948]++
										var d net.Dialer
										return d.DialContext(ctx, network, addr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:30
	// _ = "end of CoverTab[96948]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:31
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/direct.go:31
var _ = _go_fuzz_dep_.CoverTab
