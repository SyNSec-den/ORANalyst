// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64

//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:7
package nistec

//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:7
import (
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:7
)
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:7
import (
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:7
)

import "errors"

//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:13
//go:noescape
func p256OrdMul(res, in1, in2 *p256OrdElement)

//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:18
//go:noescape
func p256OrdSqr(res, in *p256OrdElement, n int)

func P256OrdInverse(k []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:21
	_go_fuzz_dep_.CoverTab[2577]++
									if len(k) != 32 {
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:22
		_go_fuzz_dep_.CoverTab[2580]++
										return nil, errors.New("invalid scalar length")
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:23
		// _ = "end of CoverTab[2580]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:24
		_go_fuzz_dep_.CoverTab[2581]++
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:24
		// _ = "end of CoverTab[2581]"
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:24
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:24
	// _ = "end of CoverTab[2577]"
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:24
	_go_fuzz_dep_.CoverTab[2578]++

									x := new(p256OrdElement)
									p256OrdBigToLittle(x, (*[32]byte)(k))
									p256OrdReduce(x)

//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:34
	_1 := new(p256OrdElement)
									_11 := new(p256OrdElement)
									_101 := new(p256OrdElement)
									_111 := new(p256OrdElement)
									_1111 := new(p256OrdElement)
									_10101 := new(p256OrdElement)
									_101111 := new(p256OrdElement)
									t := new(p256OrdElement)

//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:48
	RR := &p256OrdElement{0x83244c95be79eea2, 0x4699799c49bd6fa6,
		0x2845b2392b6bec59, 0x66e12d94f3d95620}

	p256OrdMul(_1, x, RR)
	p256OrdSqr(x, _1, 1)
	p256OrdMul(_11, x, _1)
	p256OrdMul(_101, x, _11)
	p256OrdMul(_111, x, _101)
	p256OrdSqr(x, _101, 1)
	p256OrdMul(_1111, _101, x)

	p256OrdSqr(t, x, 1)
	p256OrdMul(_10101, t, _1)
	p256OrdSqr(x, _10101, 1)
	p256OrdMul(_101111, _101, x)
	p256OrdMul(x, _10101, x)
	p256OrdSqr(t, x, 2)
	p256OrdMul(t, t, _11)
	p256OrdSqr(x, t, 8)
	p256OrdMul(x, x, t)
	p256OrdSqr(t, x, 16)
	p256OrdMul(t, t, x)

	p256OrdSqr(x, t, 64)
	p256OrdMul(x, x, t)
	p256OrdSqr(x, x, 32)
	p256OrdMul(x, x, t)

	sqrs := []int{
		6, 5, 4, 5, 5,
		4, 3, 3, 5, 9,
		6, 2, 5, 6, 5,
		4, 5, 5, 3, 10,
		2, 5, 5, 3, 7, 6}
	muls := []*p256OrdElement{
		_101111, _111, _11, _1111, _10101,
		_101, _101, _101, _111, _101111,
		_1111, _1, _1, _1111, _111,
		_111, _111, _101, _11, _101111,
		_11, _11, _11, _1, _10101, _1111}

	for i, s := range sqrs {
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:89
		_go_fuzz_dep_.CoverTab[2582]++
										p256OrdSqr(x, x, s)
										p256OrdMul(x, x, muls[i])
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:91
		// _ = "end of CoverTab[2582]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:92
	// _ = "end of CoverTab[2578]"
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:92
	_go_fuzz_dep_.CoverTab[2579]++

//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:96
	one := &p256OrdElement{1}
									p256OrdMul(x, x, one)

									var xOut [32]byte
									p256OrdLittleToBig(&xOut, x)
									return xOut[:], nil
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:101
	// _ = "end of CoverTab[2579]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:102
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/nistec/p256_ordinv.go:102
var _ = _go_fuzz_dep_.CoverTab
