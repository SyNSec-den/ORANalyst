// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || freebsd || linux || netbsd || openbsd
// +build dragonfly freebsd linux netbsd openbsd

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:8
)

import "unsafe"

// fcntl64Syscall is usually SYS_FCNTL, but is overridden on 32-bit Linux
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:12
// systems by fcntl_linux_32bit.go to be SYS_FCNTL64.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:14
var fcntl64Syscall uintptr = SYS_FCNTL

func fcntl(fd int, cmd, arg int) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:16
	_go_fuzz_dep_.CoverTab[45772]++
										valptr, _, errno := Syscall(fcntl64Syscall, uintptr(fd), uintptr(cmd), uintptr(arg))
										var err error
										if errno != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:19
		_go_fuzz_dep_.CoverTab[45774]++
											err = errno
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:20
		// _ = "end of CoverTab[45774]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:21
		_go_fuzz_dep_.CoverTab[45775]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:21
		// _ = "end of CoverTab[45775]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:21
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:21
	// _ = "end of CoverTab[45772]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:21
	_go_fuzz_dep_.CoverTab[45773]++
										return int(valptr), err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:22
	// _ = "end of CoverTab[45773]"
}

// FcntlInt performs a fcntl syscall on fd with the provided command and argument.
func FcntlInt(fd uintptr, cmd, arg int) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:26
	_go_fuzz_dep_.CoverTab[45776]++
										return fcntl(int(fd), cmd, arg)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:27
	// _ = "end of CoverTab[45776]"
}

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:31
	_go_fuzz_dep_.CoverTab[45777]++
										_, _, errno := Syscall(fcntl64Syscall, fd, uintptr(cmd), uintptr(unsafe.Pointer(lk)))
										if errno == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:33
		_go_fuzz_dep_.CoverTab[45779]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:34
		// _ = "end of CoverTab[45779]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:35
		_go_fuzz_dep_.CoverTab[45780]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:35
		// _ = "end of CoverTab[45780]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:35
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:35
	// _ = "end of CoverTab[45777]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:35
	_go_fuzz_dep_.CoverTab[45778]++
										return errno
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:36
	// _ = "end of CoverTab[45778]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:37
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/fcntl.go:37
var _ = _go_fuzz_dep_.CoverTab
