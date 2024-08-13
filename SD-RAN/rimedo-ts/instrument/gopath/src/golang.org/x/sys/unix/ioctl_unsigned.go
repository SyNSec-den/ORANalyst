// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || hurd || linux || netbsd || openbsd
// +build darwin dragonfly freebsd hurd linux netbsd openbsd

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:8
)

import (
	"unsafe"
)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:17
// IoctlSetInt performs an ioctl operation which sets an integer value
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:17
// on fd, using the specified request number.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:19
func IoctlSetInt(fd int, req uint, value int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:19
	_go_fuzz_dep_.CoverTab[45846]++
											return ioctl(fd, req, uintptr(value))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:20
	// _ = "end of CoverTab[45846]"
}

// IoctlSetPointerInt performs an ioctl operation which sets an
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:23
// integer value on fd, using the specified request number. The ioctl
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:23
// argument is called with a pointer to the integer value, rather than
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:23
// passing the integer value directly.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:27
func IoctlSetPointerInt(fd int, req uint, value int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:27
	_go_fuzz_dep_.CoverTab[45847]++
											v := int32(value)
											return ioctlPtr(fd, req, unsafe.Pointer(&v))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:29
	// _ = "end of CoverTab[45847]"
}

// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:32
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:32
// To change fd's window size, the req argument should be TIOCSWINSZ.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:35
func IoctlSetWinsize(fd int, req uint, value *Winsize) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:35
	_go_fuzz_dep_.CoverTab[45848]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:38
	return ioctlPtr(fd, req, unsafe.Pointer(value))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:38
	// _ = "end of CoverTab[45848]"
}

// IoctlSetTermios performs an ioctl on fd with a *Termios.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:41
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:41
// The req value will usually be TCSETA or TIOCSETA.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:44
func IoctlSetTermios(fd int, req uint, value *Termios) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:44
	_go_fuzz_dep_.CoverTab[45849]++

											return ioctlPtr(fd, req, unsafe.Pointer(value))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:46
	// _ = "end of CoverTab[45849]"
}

// IoctlGetInt performs an ioctl operation which gets an integer value
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:49
// from fd, using the specified request number.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:49
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:49
// A few ioctl requests use the return value as an output parameter;
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:49
// for those, IoctlRetInt should be used instead of this function.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:54
func IoctlGetInt(fd int, req uint) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:54
	_go_fuzz_dep_.CoverTab[45850]++
											var value int
											err := ioctlPtr(fd, req, unsafe.Pointer(&value))
											return value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:57
	// _ = "end of CoverTab[45850]"
}

func IoctlGetWinsize(fd int, req uint) (*Winsize, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:60
	_go_fuzz_dep_.CoverTab[45851]++
											var value Winsize
											err := ioctlPtr(fd, req, unsafe.Pointer(&value))
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:63
	// _ = "end of CoverTab[45851]"
}

func IoctlGetTermios(fd int, req uint) (*Termios, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:66
	_go_fuzz_dep_.CoverTab[45852]++
											var value Termios
											err := ioctlPtr(fd, req, unsafe.Pointer(&value))
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:69
	// _ = "end of CoverTab[45852]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:70
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_unsigned.go:70
var _ = _go_fuzz_dep_.CoverTab
