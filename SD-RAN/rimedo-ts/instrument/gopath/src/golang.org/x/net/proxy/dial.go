// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:5
package proxy

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:5
)

import (
	"context"
	"net"
)

// A ContextDialer dials using a context.
type ContextDialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

// Dial works like DialContext on net.Dialer but using a dialer returned by FromEnvironment.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:17
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:17
// The passed ctx is only used for returning the Conn, not the lifetime of the Conn.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:17
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:17
// Custom dialers (registered via RegisterDialerType) that do not implement ContextDialer
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:17
// can leak a goroutine for as long as it takes the underlying Dialer implementation to timeout.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:17
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:17
// A Conn returned from a successful Dial after the context has been cancelled will be immediately closed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:25
func Dial(ctx context.Context, network, address string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:25
	_go_fuzz_dep_.CoverTab[96933]++
										d := FromEnvironment()
										if xd, ok := d.(ContextDialer); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:27
		_go_fuzz_dep_.CoverTab[96935]++
											return xd.DialContext(ctx, network, address)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:28
		// _ = "end of CoverTab[96935]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:29
		_go_fuzz_dep_.CoverTab[96936]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:29
		// _ = "end of CoverTab[96936]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:29
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:29
	// _ = "end of CoverTab[96933]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:29
	_go_fuzz_dep_.CoverTab[96934]++
										return dialContext(ctx, d, network, address)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:30
	// _ = "end of CoverTab[96934]"
}

// WARNING: this can leak a goroutine for as long as the underlying Dialer implementation takes to timeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:33
// A Conn returned from a successful Dial after the context has been cancelled will be immediately closed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:35
func dialContext(ctx context.Context, d Dialer, network, address string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:35
	_go_fuzz_dep_.CoverTab[96937]++
										var (
		conn	net.Conn
		done	= make(chan struct{}, 1)
		err	error
	)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:40
	_curRoutineNum113_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:40
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum113_)
										go func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:41
		_go_fuzz_dep_.CoverTab[96940]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:41
		defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:41
			_go_fuzz_dep_.CoverTab[96941]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:41
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum113_)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:41
			// _ = "end of CoverTab[96941]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:41
		}()
											conn, err = d.Dial(network, address)
											close(done)
											if conn != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:44
			_go_fuzz_dep_.CoverTab[96942]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:44
			return ctx.Err() != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:44
			// _ = "end of CoverTab[96942]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:44
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:44
			_go_fuzz_dep_.CoverTab[96943]++
												conn.Close()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:45
			// _ = "end of CoverTab[96943]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:46
			_go_fuzz_dep_.CoverTab[96944]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:46
			// _ = "end of CoverTab[96944]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:46
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:46
		// _ = "end of CoverTab[96940]"
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:47
	// _ = "end of CoverTab[96937]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:47
	_go_fuzz_dep_.CoverTab[96938]++
										select {
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:49
		_go_fuzz_dep_.CoverTab[96945]++
											err = ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:50
		// _ = "end of CoverTab[96945]"
	case <-done:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:51
		_go_fuzz_dep_.CoverTab[96946]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:51
		// _ = "end of CoverTab[96946]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:52
	// _ = "end of CoverTab[96938]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:52
	_go_fuzz_dep_.CoverTab[96939]++
										return conn, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:53
	// _ = "end of CoverTab[96939]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:54
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/dial.go:54
var _ = _go_fuzz_dep_.CoverTab
