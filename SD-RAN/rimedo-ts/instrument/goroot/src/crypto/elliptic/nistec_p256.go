// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64

//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:7
package elliptic

//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:7
import (
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:7
)
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:7
import (
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:7
)

import (
	"crypto/internal/nistec"
	"math/big"
)

func (c p256Curve) Inverse(k *big.Int) *big.Int {
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:14
	_go_fuzz_dep_.CoverTab[7094]++
								if k.Sign() < 0 {
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:15
		_go_fuzz_dep_.CoverTab[7098]++

									k = new(big.Int).Neg(k)
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:17
		// _ = "end of CoverTab[7098]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:18
		_go_fuzz_dep_.CoverTab[7099]++
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:18
		// _ = "end of CoverTab[7099]"
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:18
	}
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:18
	// _ = "end of CoverTab[7094]"
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:18
	_go_fuzz_dep_.CoverTab[7095]++
								if k.Cmp(c.params.N) >= 0 {
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:19
		_go_fuzz_dep_.CoverTab[7100]++

									k = new(big.Int).Mod(k, c.params.N)
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:21
		// _ = "end of CoverTab[7100]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:22
		_go_fuzz_dep_.CoverTab[7101]++
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:22
		// _ = "end of CoverTab[7101]"
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:22
	}
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:22
	// _ = "end of CoverTab[7095]"
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:22
	_go_fuzz_dep_.CoverTab[7096]++
								scalar := k.FillBytes(make([]byte, 32))
								inverse, err := nistec.P256OrdInverse(scalar)
								if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:25
		_go_fuzz_dep_.CoverTab[7102]++
									panic("crypto/elliptic: nistec rejected normalized scalar")
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:26
		// _ = "end of CoverTab[7102]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:27
		_go_fuzz_dep_.CoverTab[7103]++
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:27
		// _ = "end of CoverTab[7103]"
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:27
	}
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:27
	// _ = "end of CoverTab[7096]"
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:27
	_go_fuzz_dep_.CoverTab[7097]++
								return new(big.Int).SetBytes(inverse)
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:28
	// _ = "end of CoverTab[7097]"
}

//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/elliptic/nistec_p256.go:29
var _ = _go_fuzz_dep_.CoverTab
