// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/alias/alias.go:5
// Package alias implements memory alaising tests.
//line /usr/local/go/src/crypto/internal/alias/alias.go:5
// This code also exists as golang.org/x/crypto/internal/alias.
//line /usr/local/go/src/crypto/internal/alias/alias.go:7
package alias

//line /usr/local/go/src/crypto/internal/alias/alias.go:7
import (
//line /usr/local/go/src/crypto/internal/alias/alias.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/alias/alias.go:7
)
//line /usr/local/go/src/crypto/internal/alias/alias.go:7
import (
//line /usr/local/go/src/crypto/internal/alias/alias.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/alias/alias.go:7
)

import "unsafe"

// AnyOverlap reports whether x and y share memory at any (not necessarily
//line /usr/local/go/src/crypto/internal/alias/alias.go:11
// corresponding) index. The memory beyond the slice length is ignored.
//line /usr/local/go/src/crypto/internal/alias/alias.go:13
func AnyOverlap(x, y []byte) bool {
//line /usr/local/go/src/crypto/internal/alias/alias.go:13
	_go_fuzz_dep_.CoverTab[1135]++
								return len(x) > 0 && func() bool {
//line /usr/local/go/src/crypto/internal/alias/alias.go:14
		_go_fuzz_dep_.CoverTab[1136]++
//line /usr/local/go/src/crypto/internal/alias/alias.go:14
		return len(y) > 0
//line /usr/local/go/src/crypto/internal/alias/alias.go:14
		// _ = "end of CoverTab[1136]"
//line /usr/local/go/src/crypto/internal/alias/alias.go:14
	}() && func() bool {
//line /usr/local/go/src/crypto/internal/alias/alias.go:14
		_go_fuzz_dep_.CoverTab[1137]++
//line /usr/local/go/src/crypto/internal/alias/alias.go:14
		return uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1]))
									// _ = "end of CoverTab[1137]"
//line /usr/local/go/src/crypto/internal/alias/alias.go:15
	}() && func() bool {
//line /usr/local/go/src/crypto/internal/alias/alias.go:15
		_go_fuzz_dep_.CoverTab[1138]++
//line /usr/local/go/src/crypto/internal/alias/alias.go:15
		return uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1]))
									// _ = "end of CoverTab[1138]"
//line /usr/local/go/src/crypto/internal/alias/alias.go:16
	}()
//line /usr/local/go/src/crypto/internal/alias/alias.go:16
	// _ = "end of CoverTab[1135]"
}

// InexactOverlap reports whether x and y share memory at any non-corresponding
//line /usr/local/go/src/crypto/internal/alias/alias.go:19
// index. The memory beyond the slice length is ignored. Note that x and y can
//line /usr/local/go/src/crypto/internal/alias/alias.go:19
// have different lengths and still not have any inexact overlap.
//line /usr/local/go/src/crypto/internal/alias/alias.go:19
//
//line /usr/local/go/src/crypto/internal/alias/alias.go:19
// InexactOverlap can be used to implement the requirements of the crypto/cipher
//line /usr/local/go/src/crypto/internal/alias/alias.go:19
// AEAD, Block, BlockMode and Stream interfaces.
//line /usr/local/go/src/crypto/internal/alias/alias.go:25
func InexactOverlap(x, y []byte) bool {
//line /usr/local/go/src/crypto/internal/alias/alias.go:25
	_go_fuzz_dep_.CoverTab[1139]++
								if len(x) == 0 || func() bool {
//line /usr/local/go/src/crypto/internal/alias/alias.go:26
		_go_fuzz_dep_.CoverTab[1141]++
//line /usr/local/go/src/crypto/internal/alias/alias.go:26
		return len(y) == 0
//line /usr/local/go/src/crypto/internal/alias/alias.go:26
		// _ = "end of CoverTab[1141]"
//line /usr/local/go/src/crypto/internal/alias/alias.go:26
	}() || func() bool {
//line /usr/local/go/src/crypto/internal/alias/alias.go:26
		_go_fuzz_dep_.CoverTab[1142]++
//line /usr/local/go/src/crypto/internal/alias/alias.go:26
		return &x[0] == &y[0]
//line /usr/local/go/src/crypto/internal/alias/alias.go:26
		// _ = "end of CoverTab[1142]"
//line /usr/local/go/src/crypto/internal/alias/alias.go:26
	}() {
//line /usr/local/go/src/crypto/internal/alias/alias.go:26
		_go_fuzz_dep_.CoverTab[1143]++
									return false
//line /usr/local/go/src/crypto/internal/alias/alias.go:27
		// _ = "end of CoverTab[1143]"
	} else {
//line /usr/local/go/src/crypto/internal/alias/alias.go:28
		_go_fuzz_dep_.CoverTab[1144]++
//line /usr/local/go/src/crypto/internal/alias/alias.go:28
		// _ = "end of CoverTab[1144]"
//line /usr/local/go/src/crypto/internal/alias/alias.go:28
	}
//line /usr/local/go/src/crypto/internal/alias/alias.go:28
	// _ = "end of CoverTab[1139]"
//line /usr/local/go/src/crypto/internal/alias/alias.go:28
	_go_fuzz_dep_.CoverTab[1140]++
								return AnyOverlap(x, y)
//line /usr/local/go/src/crypto/internal/alias/alias.go:29
	// _ = "end of CoverTab[1140]"
}

//line /usr/local/go/src/crypto/internal/alias/alias.go:30
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/alias/alias.go:30
var _ = _go_fuzz_dep_.CoverTab
