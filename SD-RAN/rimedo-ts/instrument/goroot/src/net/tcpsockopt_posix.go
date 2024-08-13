// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

//line /usr/local/go/src/net/tcpsockopt_posix.go:7
package net

//line /usr/local/go/src/net/tcpsockopt_posix.go:7
import (
//line /usr/local/go/src/net/tcpsockopt_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/tcpsockopt_posix.go:7
)
//line /usr/local/go/src/net/tcpsockopt_posix.go:7
import (
//line /usr/local/go/src/net/tcpsockopt_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/tcpsockopt_posix.go:7
)

import (
	"runtime"
	"syscall"
)

func setNoDelay(fd *netFD, noDelay bool) error {
//line /usr/local/go/src/net/tcpsockopt_posix.go:14
	_go_fuzz_dep_.CoverTab[16629]++
							err := fd.pfd.SetsockoptInt(syscall.IPPROTO_TCP, syscall.TCP_NODELAY, boolint(noDelay))
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/tcpsockopt_posix.go:17
	// _ = "end of CoverTab[16629]"
}

//line /usr/local/go/src/net/tcpsockopt_posix.go:18
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/tcpsockopt_posix.go:18
var _ = _go_fuzz_dep_.CoverTab
