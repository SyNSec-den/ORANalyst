// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || freebsd || linux || netbsd

//line /usr/local/go/src/net/tcpsockopt_unix.go:7
package net

//line /usr/local/go/src/net/tcpsockopt_unix.go:7
import (
//line /usr/local/go/src/net/tcpsockopt_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/tcpsockopt_unix.go:7
)
//line /usr/local/go/src/net/tcpsockopt_unix.go:7
import (
//line /usr/local/go/src/net/tcpsockopt_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/tcpsockopt_unix.go:7
)

import (
	"runtime"
	"syscall"
	"time"
)

func setKeepAlivePeriod(fd *netFD, d time.Duration) error {
//line /usr/local/go/src/net/tcpsockopt_unix.go:15
	_go_fuzz_dep_.CoverTab[16630]++

							secs := int(roundDurationUp(d, time.Second))
							if err := fd.pfd.SetsockoptInt(syscall.IPPROTO_TCP, syscall.TCP_KEEPINTVL, secs); err != nil {
//line /usr/local/go/src/net/tcpsockopt_unix.go:18
		_go_fuzz_dep_.CoverTab[16632]++
								return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/tcpsockopt_unix.go:19
		// _ = "end of CoverTab[16632]"
	} else {
//line /usr/local/go/src/net/tcpsockopt_unix.go:20
		_go_fuzz_dep_.CoverTab[16633]++
//line /usr/local/go/src/net/tcpsockopt_unix.go:20
		// _ = "end of CoverTab[16633]"
//line /usr/local/go/src/net/tcpsockopt_unix.go:20
	}
//line /usr/local/go/src/net/tcpsockopt_unix.go:20
	// _ = "end of CoverTab[16630]"
//line /usr/local/go/src/net/tcpsockopt_unix.go:20
	_go_fuzz_dep_.CoverTab[16631]++
							err := fd.pfd.SetsockoptInt(syscall.IPPROTO_TCP, syscall.TCP_KEEPIDLE, secs)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/tcpsockopt_unix.go:23
	// _ = "end of CoverTab[16631]"
}

//line /usr/local/go/src/net/tcpsockopt_unix.go:24
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/tcpsockopt_unix.go:24
var _ = _go_fuzz_dep_.CoverTab
