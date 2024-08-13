// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || windows

//line /usr/local/go/src/net/error_posix.go:7
package net

//line /usr/local/go/src/net/error_posix.go:7
import (
//line /usr/local/go/src/net/error_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/error_posix.go:7
)
//line /usr/local/go/src/net/error_posix.go:7
import (
//line /usr/local/go/src/net/error_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/error_posix.go:7
)

import (
	"os"
	"syscall"
)

// wrapSyscallError takes an error and a syscall name. If the error is
//line /usr/local/go/src/net/error_posix.go:14
// a syscall.Errno, it wraps it in a os.SyscallError using the syscall name.
//line /usr/local/go/src/net/error_posix.go:16
func wrapSyscallError(name string, err error) error {
//line /usr/local/go/src/net/error_posix.go:16
	_go_fuzz_dep_.CoverTab[13818]++
						if _, ok := err.(syscall.Errno); ok {
//line /usr/local/go/src/net/error_posix.go:17
		_go_fuzz_dep_.CoverTab[13820]++
							err = os.NewSyscallError(name, err)
//line /usr/local/go/src/net/error_posix.go:18
		// _ = "end of CoverTab[13820]"
	} else {
//line /usr/local/go/src/net/error_posix.go:19
		_go_fuzz_dep_.CoverTab[13821]++
//line /usr/local/go/src/net/error_posix.go:19
		// _ = "end of CoverTab[13821]"
//line /usr/local/go/src/net/error_posix.go:19
	}
//line /usr/local/go/src/net/error_posix.go:19
	// _ = "end of CoverTab[13818]"
//line /usr/local/go/src/net/error_posix.go:19
	_go_fuzz_dep_.CoverTab[13819]++
						return err
//line /usr/local/go/src/net/error_posix.go:20
	// _ = "end of CoverTab[13819]"
}

//line /usr/local/go/src/net/error_posix.go:21
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/error_posix.go:21
var _ = _go_fuzz_dep_.CoverTab
