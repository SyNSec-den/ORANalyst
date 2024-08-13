// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/hook.go:5
package net

//line /usr/local/go/src/net/hook.go:5
import (
//line /usr/local/go/src/net/hook.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/hook.go:5
)
//line /usr/local/go/src/net/hook.go:5
import (
//line /usr/local/go/src/net/hook.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/hook.go:5
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
//line /usr/local/go/src/net/hook.go:22
		_go_fuzz_dep_.CoverTab[5605]++
							return fn(ctx, network, host)
//line /usr/local/go/src/net/hook.go:23
		// _ = "end of CoverTab[5605]"
	}
	testHookSetKeepAlive	= func(time.Duration) { _go_fuzz_dep_.CoverTab[5606]++; // _ = "end of CoverTab[5606]" }
)
//line /usr/local/go/src/net/hook.go:26
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/hook.go:26
var _ = _go_fuzz_dep_.CoverTab
