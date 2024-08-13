// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/sock_linux.go:5
package net

//line /snap/go/10455/src/net/sock_linux.go:5
import (
//line /snap/go/10455/src/net/sock_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/sock_linux.go:5
)
//line /snap/go/10455/src/net/sock_linux.go:5
import (
//line /snap/go/10455/src/net/sock_linux.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/sock_linux.go:5
)

import (
	"internal/syscall/unix"
	"syscall"
)

// Linux stores the backlog as:
//line /snap/go/10455/src/net/sock_linux.go:12
//
//line /snap/go/10455/src/net/sock_linux.go:12
//   - uint16 in kernel version < 4.1,
//line /snap/go/10455/src/net/sock_linux.go:12
//   - uint32 in kernel version >= 4.1
//line /snap/go/10455/src/net/sock_linux.go:12
//
//line /snap/go/10455/src/net/sock_linux.go:12
// Truncate number to avoid wrapping.
//line /snap/go/10455/src/net/sock_linux.go:12
//
//line /snap/go/10455/src/net/sock_linux.go:12
// See issue 5030 and 41470.
//line /snap/go/10455/src/net/sock_linux.go:20
func maxAckBacklog(n int) int {
//line /snap/go/10455/src/net/sock_linux.go:20
	_go_fuzz_dep_.CoverTab[8019]++
						major, minor := unix.KernelVersion()
						size := 16
						if major > 4 || func() bool {
//line /snap/go/10455/src/net/sock_linux.go:23
		_go_fuzz_dep_.CoverTab[8022]++
//line /snap/go/10455/src/net/sock_linux.go:23
		return (major == 4 && func() bool {
//line /snap/go/10455/src/net/sock_linux.go:23
			_go_fuzz_dep_.CoverTab[8023]++
//line /snap/go/10455/src/net/sock_linux.go:23
			return minor >= 1
//line /snap/go/10455/src/net/sock_linux.go:23
			// _ = "end of CoverTab[8023]"
//line /snap/go/10455/src/net/sock_linux.go:23
		}())
//line /snap/go/10455/src/net/sock_linux.go:23
		// _ = "end of CoverTab[8022]"
//line /snap/go/10455/src/net/sock_linux.go:23
	}() {
//line /snap/go/10455/src/net/sock_linux.go:23
		_go_fuzz_dep_.CoverTab[529617]++
//line /snap/go/10455/src/net/sock_linux.go:23
		_go_fuzz_dep_.CoverTab[8024]++
							size = 32
//line /snap/go/10455/src/net/sock_linux.go:24
		// _ = "end of CoverTab[8024]"
	} else {
//line /snap/go/10455/src/net/sock_linux.go:25
		_go_fuzz_dep_.CoverTab[529618]++
//line /snap/go/10455/src/net/sock_linux.go:25
		_go_fuzz_dep_.CoverTab[8025]++
//line /snap/go/10455/src/net/sock_linux.go:25
		// _ = "end of CoverTab[8025]"
//line /snap/go/10455/src/net/sock_linux.go:25
	}
//line /snap/go/10455/src/net/sock_linux.go:25
	// _ = "end of CoverTab[8019]"
//line /snap/go/10455/src/net/sock_linux.go:25
	_go_fuzz_dep_.CoverTab[8020]++

						var max uint = 1<<size - 1
						if uint(n) > max {
//line /snap/go/10455/src/net/sock_linux.go:28
		_go_fuzz_dep_.CoverTab[529619]++
//line /snap/go/10455/src/net/sock_linux.go:28
		_go_fuzz_dep_.CoverTab[8026]++
							n = int(max)
//line /snap/go/10455/src/net/sock_linux.go:29
		// _ = "end of CoverTab[8026]"
	} else {
//line /snap/go/10455/src/net/sock_linux.go:30
		_go_fuzz_dep_.CoverTab[529620]++
//line /snap/go/10455/src/net/sock_linux.go:30
		_go_fuzz_dep_.CoverTab[8027]++
//line /snap/go/10455/src/net/sock_linux.go:30
		// _ = "end of CoverTab[8027]"
//line /snap/go/10455/src/net/sock_linux.go:30
	}
//line /snap/go/10455/src/net/sock_linux.go:30
	// _ = "end of CoverTab[8020]"
//line /snap/go/10455/src/net/sock_linux.go:30
	_go_fuzz_dep_.CoverTab[8021]++
						return n
//line /snap/go/10455/src/net/sock_linux.go:31
	// _ = "end of CoverTab[8021]"
}

func maxListenerBacklog() int {
//line /snap/go/10455/src/net/sock_linux.go:34
	_go_fuzz_dep_.CoverTab[8028]++
						fd, err := open("/proc/sys/net/core/somaxconn")
						if err != nil {
//line /snap/go/10455/src/net/sock_linux.go:36
		_go_fuzz_dep_.CoverTab[529621]++
//line /snap/go/10455/src/net/sock_linux.go:36
		_go_fuzz_dep_.CoverTab[8033]++
							return syscall.SOMAXCONN
//line /snap/go/10455/src/net/sock_linux.go:37
		// _ = "end of CoverTab[8033]"
	} else {
//line /snap/go/10455/src/net/sock_linux.go:38
		_go_fuzz_dep_.CoverTab[529622]++
//line /snap/go/10455/src/net/sock_linux.go:38
		_go_fuzz_dep_.CoverTab[8034]++
//line /snap/go/10455/src/net/sock_linux.go:38
		// _ = "end of CoverTab[8034]"
//line /snap/go/10455/src/net/sock_linux.go:38
	}
//line /snap/go/10455/src/net/sock_linux.go:38
	// _ = "end of CoverTab[8028]"
//line /snap/go/10455/src/net/sock_linux.go:38
	_go_fuzz_dep_.CoverTab[8029]++
						defer fd.close()
						l, ok := fd.readLine()
						if !ok {
//line /snap/go/10455/src/net/sock_linux.go:41
		_go_fuzz_dep_.CoverTab[529623]++
//line /snap/go/10455/src/net/sock_linux.go:41
		_go_fuzz_dep_.CoverTab[8035]++
							return syscall.SOMAXCONN
//line /snap/go/10455/src/net/sock_linux.go:42
		// _ = "end of CoverTab[8035]"
	} else {
//line /snap/go/10455/src/net/sock_linux.go:43
		_go_fuzz_dep_.CoverTab[529624]++
//line /snap/go/10455/src/net/sock_linux.go:43
		_go_fuzz_dep_.CoverTab[8036]++
//line /snap/go/10455/src/net/sock_linux.go:43
		// _ = "end of CoverTab[8036]"
//line /snap/go/10455/src/net/sock_linux.go:43
	}
//line /snap/go/10455/src/net/sock_linux.go:43
	// _ = "end of CoverTab[8029]"
//line /snap/go/10455/src/net/sock_linux.go:43
	_go_fuzz_dep_.CoverTab[8030]++
						f := getFields(l)
						n, _, ok := dtoi(f[0])
						if n == 0 || func() bool {
//line /snap/go/10455/src/net/sock_linux.go:46
		_go_fuzz_dep_.CoverTab[8037]++
//line /snap/go/10455/src/net/sock_linux.go:46
		return !ok
//line /snap/go/10455/src/net/sock_linux.go:46
		// _ = "end of CoverTab[8037]"
//line /snap/go/10455/src/net/sock_linux.go:46
	}() {
//line /snap/go/10455/src/net/sock_linux.go:46
		_go_fuzz_dep_.CoverTab[529625]++
//line /snap/go/10455/src/net/sock_linux.go:46
		_go_fuzz_dep_.CoverTab[8038]++
							return syscall.SOMAXCONN
//line /snap/go/10455/src/net/sock_linux.go:47
		// _ = "end of CoverTab[8038]"
	} else {
//line /snap/go/10455/src/net/sock_linux.go:48
		_go_fuzz_dep_.CoverTab[529626]++
//line /snap/go/10455/src/net/sock_linux.go:48
		_go_fuzz_dep_.CoverTab[8039]++
//line /snap/go/10455/src/net/sock_linux.go:48
		// _ = "end of CoverTab[8039]"
//line /snap/go/10455/src/net/sock_linux.go:48
	}
//line /snap/go/10455/src/net/sock_linux.go:48
	// _ = "end of CoverTab[8030]"
//line /snap/go/10455/src/net/sock_linux.go:48
	_go_fuzz_dep_.CoverTab[8031]++

						if n > 1<<16-1 {
//line /snap/go/10455/src/net/sock_linux.go:50
		_go_fuzz_dep_.CoverTab[529627]++
//line /snap/go/10455/src/net/sock_linux.go:50
		_go_fuzz_dep_.CoverTab[8040]++
							return maxAckBacklog(n)
//line /snap/go/10455/src/net/sock_linux.go:51
		// _ = "end of CoverTab[8040]"
	} else {
//line /snap/go/10455/src/net/sock_linux.go:52
		_go_fuzz_dep_.CoverTab[529628]++
//line /snap/go/10455/src/net/sock_linux.go:52
		_go_fuzz_dep_.CoverTab[8041]++
//line /snap/go/10455/src/net/sock_linux.go:52
		// _ = "end of CoverTab[8041]"
//line /snap/go/10455/src/net/sock_linux.go:52
	}
//line /snap/go/10455/src/net/sock_linux.go:52
	// _ = "end of CoverTab[8031]"
//line /snap/go/10455/src/net/sock_linux.go:52
	_go_fuzz_dep_.CoverTab[8032]++
						return n
//line /snap/go/10455/src/net/sock_linux.go:53
	// _ = "end of CoverTab[8032]"
}

//line /snap/go/10455/src/net/sock_linux.go:54
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/sock_linux.go:54
var _ = _go_fuzz_dep_.CoverTab
