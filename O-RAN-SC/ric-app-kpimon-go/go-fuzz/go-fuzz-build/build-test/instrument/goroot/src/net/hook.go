// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/hook.go:5
package net

//line /snap/go/10455/src/net/hook.go:5
import (
//line /snap/go/10455/src/net/hook.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/hook.go:5
)
//line /snap/go/10455/src/net/hook.go:5
import (
//line /snap/go/10455/src/net/hook.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/hook.go:5
)

import (
	"context"
	"time"
)

var (
	// if non-nil, overrides dialTCP.
	testHookDialTCP	func(ctx context.Context, net string, laddr, raddr *TCPAddr) (*TCPConn, error)

	testHookHostsPath	= "/etc/hosts"
	testHookLookupIP	= func(
		ctx context.Context,
		fn func(context.Context, string, string) ([]IPAddr, error),
		network string,
		host string,
	) ([]IPAddr, error) {
//line /snap/go/10455/src/net/hook.go:22
		_go_fuzz_dep_.CoverTab[5981]++
							return fn(ctx, network, host)
//line /snap/go/10455/src/net/hook.go:23
		// _ = "end of CoverTab[5981]"
	}
	testHookSetKeepAlive	= func(time.Duration) { _go_fuzz_dep_.CoverTab[5982]++; // _ = "end of CoverTab[5982]" }
)
//line /snap/go/10455/src/net/hook.go:26
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/hook.go:26
var _ = _go_fuzz_dep_.CoverTab
