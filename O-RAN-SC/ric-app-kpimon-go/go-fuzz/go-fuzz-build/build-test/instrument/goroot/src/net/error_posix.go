// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

//line /snap/go/10455/src/net/error_posix.go:7
package net

//line /snap/go/10455/src/net/error_posix.go:7
import (
//line /snap/go/10455/src/net/error_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/error_posix.go:7
)
//line /snap/go/10455/src/net/error_posix.go:7
import (
//line /snap/go/10455/src/net/error_posix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/error_posix.go:7
)

import (
	"os"
	"syscall"
)

// wrapSyscallError takes an error and a syscall name. If the error is
//line /snap/go/10455/src/net/error_posix.go:14
// a syscall.Errno, it wraps it in an os.SyscallError using the syscall name.
//line /snap/go/10455/src/net/error_posix.go:16
func wrapSyscallError(name string, err error) error {
//line /snap/go/10455/src/net/error_posix.go:16
	_go_fuzz_dep_.CoverTab[5804]++
							if _, ok := err.(syscall.Errno); ok {
//line /snap/go/10455/src/net/error_posix.go:17
		_go_fuzz_dep_.CoverTab[528241]++
//line /snap/go/10455/src/net/error_posix.go:17
		_go_fuzz_dep_.CoverTab[5806]++
								err = os.NewSyscallError(name, err)
//line /snap/go/10455/src/net/error_posix.go:18
		// _ = "end of CoverTab[5806]"
	} else {
//line /snap/go/10455/src/net/error_posix.go:19
		_go_fuzz_dep_.CoverTab[528242]++
//line /snap/go/10455/src/net/error_posix.go:19
		_go_fuzz_dep_.CoverTab[5807]++
//line /snap/go/10455/src/net/error_posix.go:19
		// _ = "end of CoverTab[5807]"
//line /snap/go/10455/src/net/error_posix.go:19
	}
//line /snap/go/10455/src/net/error_posix.go:19
	// _ = "end of CoverTab[5804]"
//line /snap/go/10455/src/net/error_posix.go:19
	_go_fuzz_dep_.CoverTab[5805]++
							return err
//line /snap/go/10455/src/net/error_posix.go:20
	// _ = "end of CoverTab[5805]"
}

//line /snap/go/10455/src/net/error_posix.go:21
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/error_posix.go:21
var _ = _go_fuzz_dep_.CoverTab
