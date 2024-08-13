// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego
// +build !purego

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:8
// Package alias implements memory aliasing tests.
package alias

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:9
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:9
)

import "unsafe"

// AnyOverlap reports whether x and y share memory at any (not necessarily
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:13
// corresponding) index. The memory beyond the slice length is ignored.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:15
func AnyOverlap(x, y []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:15
	_go_fuzz_dep_.CoverTab[20647]++
										return len(x) > 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:16
		_go_fuzz_dep_.CoverTab[20648]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:16
		return len(y) > 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:16
		// _ = "end of CoverTab[20648]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:16
	}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:16
		_go_fuzz_dep_.CoverTab[20649]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:16
		return uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1]))
											// _ = "end of CoverTab[20649]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:17
	}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:17
		_go_fuzz_dep_.CoverTab[20650]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:17
		return uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1]))
											// _ = "end of CoverTab[20650]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:18
	}()
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:18
	// _ = "end of CoverTab[20647]"
}

// InexactOverlap reports whether x and y share memory at any non-corresponding
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:21
// index. The memory beyond the slice length is ignored. Note that x and y can
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:21
// have different lengths and still not have any inexact overlap.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:21
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:21
// InexactOverlap can be used to implement the requirements of the crypto/cipher
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:21
// AEAD, Block, BlockMode and Stream interfaces.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:27
func InexactOverlap(x, y []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:27
	_go_fuzz_dep_.CoverTab[20651]++
										if len(x) == 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:28
		_go_fuzz_dep_.CoverTab[20653]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:28
		return len(y) == 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:28
		// _ = "end of CoverTab[20653]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:28
	}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:28
		_go_fuzz_dep_.CoverTab[20654]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:28
		return &x[0] == &y[0]
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:28
		// _ = "end of CoverTab[20654]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:28
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:28
		_go_fuzz_dep_.CoverTab[20655]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:29
		// _ = "end of CoverTab[20655]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:30
		_go_fuzz_dep_.CoverTab[20656]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:30
		// _ = "end of CoverTab[20656]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:30
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:30
	// _ = "end of CoverTab[20651]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:30
	_go_fuzz_dep_.CoverTab[20652]++
										return AnyOverlap(x, y)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:31
	// _ = "end of CoverTab[20652]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:32
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/alias/alias.go:32
var _ = _go_fuzz_dep_.CoverTab
