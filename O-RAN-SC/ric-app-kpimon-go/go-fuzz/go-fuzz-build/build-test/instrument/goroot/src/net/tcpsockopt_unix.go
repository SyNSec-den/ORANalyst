// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || freebsd || linux || netbsd

//line /snap/go/10455/src/net/tcpsockopt_unix.go:7
package net

//line /snap/go/10455/src/net/tcpsockopt_unix.go:7
import (
//line /snap/go/10455/src/net/tcpsockopt_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/tcpsockopt_unix.go:7
)
//line /snap/go/10455/src/net/tcpsockopt_unix.go:7
import (
//line /snap/go/10455/src/net/tcpsockopt_unix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/tcpsockopt_unix.go:7
)

import (
	"runtime"
	"syscall"
	"time"
)

func setKeepAlivePeriod(fd *netFD, d time.Duration) error {
//line /snap/go/10455/src/net/tcpsockopt_unix.go:15
	_go_fuzz_dep_.CoverTab[8540]++

							secs := int(roundDurationUp(d, time.Second))
							if err := fd.pfd.SetsockoptInt(syscall.IPPROTO_TCP, syscall.TCP_KEEPINTVL, secs); err != nil {
//line /snap/go/10455/src/net/tcpsockopt_unix.go:18
		_go_fuzz_dep_.CoverTab[529936]++
//line /snap/go/10455/src/net/tcpsockopt_unix.go:18
		_go_fuzz_dep_.CoverTab[8542]++
								return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/tcpsockopt_unix.go:19
		// _ = "end of CoverTab[8542]"
	} else {
//line /snap/go/10455/src/net/tcpsockopt_unix.go:20
		_go_fuzz_dep_.CoverTab[529937]++
//line /snap/go/10455/src/net/tcpsockopt_unix.go:20
		_go_fuzz_dep_.CoverTab[8543]++
//line /snap/go/10455/src/net/tcpsockopt_unix.go:20
		// _ = "end of CoverTab[8543]"
//line /snap/go/10455/src/net/tcpsockopt_unix.go:20
	}
//line /snap/go/10455/src/net/tcpsockopt_unix.go:20
	// _ = "end of CoverTab[8540]"
//line /snap/go/10455/src/net/tcpsockopt_unix.go:20
	_go_fuzz_dep_.CoverTab[8541]++
							err := fd.pfd.SetsockoptInt(syscall.IPPROTO_TCP, syscall.TCP_KEEPIDLE, secs)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/tcpsockopt_unix.go:23
	// _ = "end of CoverTab[8541]"
}

//line /snap/go/10455/src/net/tcpsockopt_unix.go:24
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/tcpsockopt_unix.go:24
var _ = _go_fuzz_dep_.CoverTab
