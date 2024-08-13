// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.13
// +build go1.13

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:8
package poly1305

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:8
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:8
)

import "math/bits"

func bitsAdd64(x, y, carry uint64) (sum, carryOut uint64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:12
	_go_fuzz_dep_.CoverTab[20736]++
												return bits.Add64(x, y, carry)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:13
	// _ = "end of CoverTab[20736]"
}

func bitsSub64(x, y, borrow uint64) (diff, borrowOut uint64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:16
	_go_fuzz_dep_.CoverTab[20737]++
												return bits.Sub64(x, y, borrow)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:17
	// _ = "end of CoverTab[20737]"
}

func bitsMul64(x, y uint64) (hi, lo uint64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:20
	_go_fuzz_dep_.CoverTab[20738]++
												return bits.Mul64(x, y)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:21
	// _ = "end of CoverTab[20738]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:22
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/bits_go1.13.go:22
var _ = _go_fuzz_dep_.CoverTab
