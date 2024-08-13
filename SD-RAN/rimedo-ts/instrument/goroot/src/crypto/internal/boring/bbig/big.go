// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:5
package bbig

//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:5
import (
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:5
)
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:5
import (
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:5
)

import (
	"crypto/internal/boring"
	"math/big"
	"unsafe"
)

func Enc(b *big.Int) boring.BigInt {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:13
	_go_fuzz_dep_.CoverTab[7319]++
								if b == nil {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:14
		_go_fuzz_dep_.CoverTab[7322]++
									return nil
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:15
		// _ = "end of CoverTab[7322]"
	} else {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:16
		_go_fuzz_dep_.CoverTab[7323]++
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:16
		// _ = "end of CoverTab[7323]"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:16
	}
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:16
	// _ = "end of CoverTab[7319]"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:16
	_go_fuzz_dep_.CoverTab[7320]++
								x := b.Bits()
								if len(x) == 0 {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:18
		_go_fuzz_dep_.CoverTab[7324]++
									return boring.BigInt{}
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:19
		// _ = "end of CoverTab[7324]"
	} else {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:20
		_go_fuzz_dep_.CoverTab[7325]++
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:20
		// _ = "end of CoverTab[7325]"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:20
	}
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:20
	// _ = "end of CoverTab[7320]"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:20
	_go_fuzz_dep_.CoverTab[7321]++
								return unsafe.Slice((*uint)(&x[0]), len(x))
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:21
	// _ = "end of CoverTab[7321]"
}

func Dec(b boring.BigInt) *big.Int {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:24
	_go_fuzz_dep_.CoverTab[7326]++
								if b == nil {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:25
		_go_fuzz_dep_.CoverTab[7329]++
									return nil
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:26
		// _ = "end of CoverTab[7329]"
	} else {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:27
		_go_fuzz_dep_.CoverTab[7330]++
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:27
		// _ = "end of CoverTab[7330]"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:27
	}
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:27
	// _ = "end of CoverTab[7326]"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:27
	_go_fuzz_dep_.CoverTab[7327]++
								if len(b) == 0 {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:28
		_go_fuzz_dep_.CoverTab[7331]++
									return new(big.Int)
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:29
		// _ = "end of CoverTab[7331]"
	} else {
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:30
		_go_fuzz_dep_.CoverTab[7332]++
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:30
		// _ = "end of CoverTab[7332]"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:30
	}
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:30
	// _ = "end of CoverTab[7327]"
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:30
	_go_fuzz_dep_.CoverTab[7328]++
								x := unsafe.Slice((*big.Word)(&b[0]), len(b))
								return new(big.Int).SetBits(x)
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:32
	// _ = "end of CoverTab[7328]"
}

//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:33
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/boring/bbig/big.go:33
var _ = _go_fuzz_dep_.CoverTab
