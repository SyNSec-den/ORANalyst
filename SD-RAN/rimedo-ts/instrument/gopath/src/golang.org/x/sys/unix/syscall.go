// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// Package unix contains an interface to the low-level operating system
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// primitives. OS details vary depending on the underlying system, and
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// by default, godoc will display OS-specific documentation for the current
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// system. If you want godoc to display OS documentation for another
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// system, set $GOOS and $GOARCH to the desired system. For example, if
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// you want to view documentation for freebsd/arm on linux/amd64, set $GOOS
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// to freebsd and $GOARCH to arm.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// The primary use of this package is inside other packages that provide a more
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// portable interface to the system, such as "os", "time" and "net".  Use
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// those packages rather than this one if you can.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// For details of the functions and data types in this package consult
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// the manuals for the appropriate operating system.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// These calls return err == nil to indicate success; otherwise
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// err represents an operating system error describing the failure and
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:8
// holds a value of type syscall.Errno.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:26
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:26
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:26
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:26
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:26
)

import (
	"bytes"
	"strings"
	"unsafe"
)

// ByteSliceFromString returns a NUL-terminated slice of bytes
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:34
// containing the text of s. If s contains a NUL byte at any
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:34
// location, it returns (nil, EINVAL).
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:37
func ByteSliceFromString(s string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:37
	_go_fuzz_dep_.CoverTab[45923]++
										if strings.IndexByte(s, 0) != -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:38
		_go_fuzz_dep_.CoverTab[45925]++
											return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:39
		// _ = "end of CoverTab[45925]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:40
		_go_fuzz_dep_.CoverTab[45926]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:40
		// _ = "end of CoverTab[45926]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:40
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:40
	// _ = "end of CoverTab[45923]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:40
	_go_fuzz_dep_.CoverTab[45924]++
										a := make([]byte, len(s)+1)
										copy(a, s)
										return a, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:43
	// _ = "end of CoverTab[45924]"
}

// BytePtrFromString returns a pointer to a NUL-terminated array of
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:46
// bytes containing the text of s. If s contains a NUL byte at any
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:46
// location, it returns (nil, EINVAL).
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:49
func BytePtrFromString(s string) (*byte, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:49
	_go_fuzz_dep_.CoverTab[45927]++
										a, err := ByteSliceFromString(s)
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:51
		_go_fuzz_dep_.CoverTab[45929]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:52
		// _ = "end of CoverTab[45929]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:53
		_go_fuzz_dep_.CoverTab[45930]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:53
		// _ = "end of CoverTab[45930]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:53
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:53
	// _ = "end of CoverTab[45927]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:53
	_go_fuzz_dep_.CoverTab[45928]++
										return &a[0], nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:54
	// _ = "end of CoverTab[45928]"
}

// ByteSliceToString returns a string form of the text represented by the slice s, with a terminating NUL and any
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:57
// bytes after the NUL removed.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:59
func ByteSliceToString(s []byte) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:59
	_go_fuzz_dep_.CoverTab[45931]++
										if i := bytes.IndexByte(s, 0); i != -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:60
		_go_fuzz_dep_.CoverTab[45933]++
											s = s[:i]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:61
		// _ = "end of CoverTab[45933]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:62
		_go_fuzz_dep_.CoverTab[45934]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:62
		// _ = "end of CoverTab[45934]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:62
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:62
	// _ = "end of CoverTab[45931]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:62
	_go_fuzz_dep_.CoverTab[45932]++
										return string(s)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:63
	// _ = "end of CoverTab[45932]"
}

// BytePtrToString takes a pointer to a sequence of text and returns the corresponding string.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:66
// If the pointer is nil, it returns the empty string. It assumes that the text sequence is terminated
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:66
// at a zero byte; if the zero byte is not present, the program may crash.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:69
func BytePtrToString(p *byte) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:69
	_go_fuzz_dep_.CoverTab[45935]++
										if p == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:70
		_go_fuzz_dep_.CoverTab[45939]++
											return ""
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:71
		// _ = "end of CoverTab[45939]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:72
		_go_fuzz_dep_.CoverTab[45940]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:72
		// _ = "end of CoverTab[45940]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:72
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:72
	// _ = "end of CoverTab[45935]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:72
	_go_fuzz_dep_.CoverTab[45936]++
										if *p == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:73
		_go_fuzz_dep_.CoverTab[45941]++
											return ""
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:74
		// _ = "end of CoverTab[45941]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:75
		_go_fuzz_dep_.CoverTab[45942]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:75
		// _ = "end of CoverTab[45942]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:75
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:75
	// _ = "end of CoverTab[45936]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:75
	_go_fuzz_dep_.CoverTab[45937]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:78
	n := 0
	for ptr := unsafe.Pointer(p); *(*byte)(ptr) != 0; n++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:79
		_go_fuzz_dep_.CoverTab[45943]++
											ptr = unsafe.Pointer(uintptr(ptr) + 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:80
		// _ = "end of CoverTab[45943]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:81
	// _ = "end of CoverTab[45937]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:81
	_go_fuzz_dep_.CoverTab[45938]++

										return string(unsafe.Slice(p, n))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:83
	// _ = "end of CoverTab[45938]"
}

// Single-word zero for use when we need a valid pointer to 0 bytes.
var _zero uintptr
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:87
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall.go:87
var _ = _go_fuzz_dep_.CoverTab
