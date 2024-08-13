// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/sock_linux.go:5
package net

//line /usr/local/go/src/net/sock_linux.go:5
import (
//line /usr/local/go/src/net/sock_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/sock_linux.go:5
)
//line /usr/local/go/src/net/sock_linux.go:5
import (
//line /usr/local/go/src/net/sock_linux.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/sock_linux.go:5
)

import (
	"internal/syscall/unix"
	"syscall"
)

// Linux stores the backlog as:
//line /usr/local/go/src/net/sock_linux.go:12
//
//line /usr/local/go/src/net/sock_linux.go:12
//   - uint16 in kernel version < 4.1,
//line /usr/local/go/src/net/sock_linux.go:12
//   - uint32 in kernel version >= 4.1
//line /usr/local/go/src/net/sock_linux.go:12
//
//line /usr/local/go/src/net/sock_linux.go:12
// Truncate number to avoid wrapping.
//line /usr/local/go/src/net/sock_linux.go:12
//
//line /usr/local/go/src/net/sock_linux.go:12
// See issue 5030 and 41470.
//line /usr/local/go/src/net/sock_linux.go:20
func maxAckBacklog(n int) int {
//line /usr/local/go/src/net/sock_linux.go:20
	_go_fuzz_dep_.CoverTab[16115]++
						major, minor := unix.KernelVersion()
						size := 16
						if major > 4 || func() bool {
//line /usr/local/go/src/net/sock_linux.go:23
		_go_fuzz_dep_.CoverTab[16118]++
//line /usr/local/go/src/net/sock_linux.go:23
		return (major == 4 && func() bool {
//line /usr/local/go/src/net/sock_linux.go:23
			_go_fuzz_dep_.CoverTab[16119]++
//line /usr/local/go/src/net/sock_linux.go:23
			return minor >= 1
//line /usr/local/go/src/net/sock_linux.go:23
			// _ = "end of CoverTab[16119]"
//line /usr/local/go/src/net/sock_linux.go:23
		}())
//line /usr/local/go/src/net/sock_linux.go:23
		// _ = "end of CoverTab[16118]"
//line /usr/local/go/src/net/sock_linux.go:23
	}() {
//line /usr/local/go/src/net/sock_linux.go:23
		_go_fuzz_dep_.CoverTab[16120]++
							size = 32
//line /usr/local/go/src/net/sock_linux.go:24
		// _ = "end of CoverTab[16120]"
	} else {
//line /usr/local/go/src/net/sock_linux.go:25
		_go_fuzz_dep_.CoverTab[16121]++
//line /usr/local/go/src/net/sock_linux.go:25
		// _ = "end of CoverTab[16121]"
//line /usr/local/go/src/net/sock_linux.go:25
	}
//line /usr/local/go/src/net/sock_linux.go:25
	// _ = "end of CoverTab[16115]"
//line /usr/local/go/src/net/sock_linux.go:25
	_go_fuzz_dep_.CoverTab[16116]++

						var max uint = 1<<size - 1
						if uint(n) > max {
//line /usr/local/go/src/net/sock_linux.go:28
		_go_fuzz_dep_.CoverTab[16122]++
							n = int(max)
//line /usr/local/go/src/net/sock_linux.go:29
		// _ = "end of CoverTab[16122]"
	} else {
//line /usr/local/go/src/net/sock_linux.go:30
		_go_fuzz_dep_.CoverTab[16123]++
//line /usr/local/go/src/net/sock_linux.go:30
		// _ = "end of CoverTab[16123]"
//line /usr/local/go/src/net/sock_linux.go:30
	}
//line /usr/local/go/src/net/sock_linux.go:30
	// _ = "end of CoverTab[16116]"
//line /usr/local/go/src/net/sock_linux.go:30
	_go_fuzz_dep_.CoverTab[16117]++
						return n
//line /usr/local/go/src/net/sock_linux.go:31
	// _ = "end of CoverTab[16117]"
}

func maxListenerBacklog() int {
//line /usr/local/go/src/net/sock_linux.go:34
	_go_fuzz_dep_.CoverTab[16124]++
						fd, err := open("/proc/sys/net/core/somaxconn")
						if err != nil {
//line /usr/local/go/src/net/sock_linux.go:36
		_go_fuzz_dep_.CoverTab[16129]++
							return syscall.SOMAXCONN
//line /usr/local/go/src/net/sock_linux.go:37
		// _ = "end of CoverTab[16129]"
	} else {
//line /usr/local/go/src/net/sock_linux.go:38
		_go_fuzz_dep_.CoverTab[16130]++
//line /usr/local/go/src/net/sock_linux.go:38
		// _ = "end of CoverTab[16130]"
//line /usr/local/go/src/net/sock_linux.go:38
	}
//line /usr/local/go/src/net/sock_linux.go:38
	// _ = "end of CoverTab[16124]"
//line /usr/local/go/src/net/sock_linux.go:38
	_go_fuzz_dep_.CoverTab[16125]++
						defer fd.close()
						l, ok := fd.readLine()
						if !ok {
//line /usr/local/go/src/net/sock_linux.go:41
		_go_fuzz_dep_.CoverTab[16131]++
							return syscall.SOMAXCONN
//line /usr/local/go/src/net/sock_linux.go:42
		// _ = "end of CoverTab[16131]"
	} else {
//line /usr/local/go/src/net/sock_linux.go:43
		_go_fuzz_dep_.CoverTab[16132]++
//line /usr/local/go/src/net/sock_linux.go:43
		// _ = "end of CoverTab[16132]"
//line /usr/local/go/src/net/sock_linux.go:43
	}
//line /usr/local/go/src/net/sock_linux.go:43
	// _ = "end of CoverTab[16125]"
//line /usr/local/go/src/net/sock_linux.go:43
	_go_fuzz_dep_.CoverTab[16126]++
						f := getFields(l)
						n, _, ok := dtoi(f[0])
						if n == 0 || func() bool {
//line /usr/local/go/src/net/sock_linux.go:46
		_go_fuzz_dep_.CoverTab[16133]++
//line /usr/local/go/src/net/sock_linux.go:46
		return !ok
//line /usr/local/go/src/net/sock_linux.go:46
		// _ = "end of CoverTab[16133]"
//line /usr/local/go/src/net/sock_linux.go:46
	}() {
//line /usr/local/go/src/net/sock_linux.go:46
		_go_fuzz_dep_.CoverTab[16134]++
							return syscall.SOMAXCONN
//line /usr/local/go/src/net/sock_linux.go:47
		// _ = "end of CoverTab[16134]"
	} else {
//line /usr/local/go/src/net/sock_linux.go:48
		_go_fuzz_dep_.CoverTab[16135]++
//line /usr/local/go/src/net/sock_linux.go:48
		// _ = "end of CoverTab[16135]"
//line /usr/local/go/src/net/sock_linux.go:48
	}
//line /usr/local/go/src/net/sock_linux.go:48
	// _ = "end of CoverTab[16126]"
//line /usr/local/go/src/net/sock_linux.go:48
	_go_fuzz_dep_.CoverTab[16127]++

						if n > 1<<16-1 {
//line /usr/local/go/src/net/sock_linux.go:50
		_go_fuzz_dep_.CoverTab[16136]++
							return maxAckBacklog(n)
//line /usr/local/go/src/net/sock_linux.go:51
		// _ = "end of CoverTab[16136]"
	} else {
//line /usr/local/go/src/net/sock_linux.go:52
		_go_fuzz_dep_.CoverTab[16137]++
//line /usr/local/go/src/net/sock_linux.go:52
		// _ = "end of CoverTab[16137]"
//line /usr/local/go/src/net/sock_linux.go:52
	}
//line /usr/local/go/src/net/sock_linux.go:52
	// _ = "end of CoverTab[16127]"
//line /usr/local/go/src/net/sock_linux.go:52
	_go_fuzz_dep_.CoverTab[16128]++
						return n
//line /usr/local/go/src/net/sock_linux.go:53
	// _ = "end of CoverTab[16128]"
}

//line /usr/local/go/src/net/sock_linux.go:54
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sock_linux.go:54
var _ = _go_fuzz_dep_.CoverTab
