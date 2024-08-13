// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || js

//line /usr/local/go/src/net/error_unix.go:7
package net

//line /usr/local/go/src/net/error_unix.go:7
import (
//line /usr/local/go/src/net/error_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/error_unix.go:7
)
//line /usr/local/go/src/net/error_unix.go:7
import (
//line /usr/local/go/src/net/error_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/error_unix.go:7
)

import "syscall"

func isConnError(err error) bool {
//line /usr/local/go/src/net/error_unix.go:11
	_go_fuzz_dep_.CoverTab[13822]++
						if se, ok := err.(syscall.Errno); ok {
//line /usr/local/go/src/net/error_unix.go:12
		_go_fuzz_dep_.CoverTab[13824]++
							return se == syscall.ECONNRESET || func() bool {
//line /usr/local/go/src/net/error_unix.go:13
			_go_fuzz_dep_.CoverTab[13825]++
//line /usr/local/go/src/net/error_unix.go:13
			return se == syscall.ECONNABORTED
//line /usr/local/go/src/net/error_unix.go:13
			// _ = "end of CoverTab[13825]"
//line /usr/local/go/src/net/error_unix.go:13
		}()
//line /usr/local/go/src/net/error_unix.go:13
		// _ = "end of CoverTab[13824]"
	} else {
//line /usr/local/go/src/net/error_unix.go:14
		_go_fuzz_dep_.CoverTab[13826]++
//line /usr/local/go/src/net/error_unix.go:14
		// _ = "end of CoverTab[13826]"
//line /usr/local/go/src/net/error_unix.go:14
	}
//line /usr/local/go/src/net/error_unix.go:14
	// _ = "end of CoverTab[13822]"
//line /usr/local/go/src/net/error_unix.go:14
	_go_fuzz_dep_.CoverTab[13823]++
						return false
//line /usr/local/go/src/net/error_unix.go:15
	// _ = "end of CoverTab[13823]"
}

//line /usr/local/go/src/net/error_unix.go:16
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/error_unix.go:16
var _ = _go_fuzz_dep_.CoverTab
