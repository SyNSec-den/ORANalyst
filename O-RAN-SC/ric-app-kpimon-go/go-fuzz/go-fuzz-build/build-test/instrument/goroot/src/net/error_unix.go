// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || js || wasip1

//line /snap/go/10455/src/net/error_unix.go:7
package net

//line /snap/go/10455/src/net/error_unix.go:7
import (
//line /snap/go/10455/src/net/error_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/error_unix.go:7
)
//line /snap/go/10455/src/net/error_unix.go:7
import (
//line /snap/go/10455/src/net/error_unix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/error_unix.go:7
)

import "syscall"

func isConnError(err error) bool {
//line /snap/go/10455/src/net/error_unix.go:11
	_go_fuzz_dep_.CoverTab[5808]++
						if se, ok := err.(syscall.Errno); ok {
//line /snap/go/10455/src/net/error_unix.go:12
		_go_fuzz_dep_.CoverTab[528243]++
//line /snap/go/10455/src/net/error_unix.go:12
		_go_fuzz_dep_.CoverTab[5810]++
							return se == syscall.ECONNRESET || func() bool {
//line /snap/go/10455/src/net/error_unix.go:13
			_go_fuzz_dep_.CoverTab[5811]++
//line /snap/go/10455/src/net/error_unix.go:13
			return se == syscall.ECONNABORTED
//line /snap/go/10455/src/net/error_unix.go:13
			// _ = "end of CoverTab[5811]"
//line /snap/go/10455/src/net/error_unix.go:13
		}()
//line /snap/go/10455/src/net/error_unix.go:13
		// _ = "end of CoverTab[5810]"
	} else {
//line /snap/go/10455/src/net/error_unix.go:14
		_go_fuzz_dep_.CoverTab[528244]++
//line /snap/go/10455/src/net/error_unix.go:14
		_go_fuzz_dep_.CoverTab[5812]++
//line /snap/go/10455/src/net/error_unix.go:14
		// _ = "end of CoverTab[5812]"
//line /snap/go/10455/src/net/error_unix.go:14
	}
//line /snap/go/10455/src/net/error_unix.go:14
	// _ = "end of CoverTab[5808]"
//line /snap/go/10455/src/net/error_unix.go:14
	_go_fuzz_dep_.CoverTab[5809]++
						return false
//line /snap/go/10455/src/net/error_unix.go:15
	// _ = "end of CoverTab[5809]"
}

//line /snap/go/10455/src/net/error_unix.go:16
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/error_unix.go:16
var _ = _go_fuzz_dep_.CoverTab
