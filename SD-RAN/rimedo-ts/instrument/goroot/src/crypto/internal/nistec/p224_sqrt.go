// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:5
package nistec

//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:5
import (
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:5
)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:5
import (
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:5
)

import (
	"crypto/internal/nistec/fiat"
	"sync"
)

var p224GG *[96]fiat.P224Element
var p224GGOnce sync.Once

// p224SqrtCandidate sets r to a square root candidate for x. r and x must not overlap.
func p224SqrtCandidate(r, x *fiat.P224Element) {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:16
	_go_fuzz_dep_.CoverTab[2470]++

//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:27
	p224GGOnce.Do(func() {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:27
		_go_fuzz_dep_.CoverTab[2479]++
										p224GG = new([96]fiat.P224Element)
										for i := range p224GG {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:29
			_go_fuzz_dep_.CoverTab[2480]++
											if i == 0 {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:30
				_go_fuzz_dep_.CoverTab[2481]++
												p224GG[i].SetBytes([]byte{0x6a, 0x0f, 0xec, 0x67,
					0x85, 0x98, 0xa7, 0x92, 0x0c, 0x55, 0xb2, 0xd4,
					0x0b, 0x2d, 0x6f, 0xfb, 0xbe, 0xa3, 0xd8, 0xce,
					0xf3, 0xfb, 0x36, 0x32, 0xdc, 0x69, 0x1b, 0x74})
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:34
				// _ = "end of CoverTab[2481]"
			} else {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:35
				_go_fuzz_dep_.CoverTab[2482]++
												p224GG[i].Square(&p224GG[i-1])
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:36
				// _ = "end of CoverTab[2482]"
			}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:37
			// _ = "end of CoverTab[2480]"
		}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:38
		// _ = "end of CoverTab[2479]"
	})
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:39
	// _ = "end of CoverTab[2470]"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:39
	_go_fuzz_dep_.CoverTab[2471]++

//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:44
	// Compute x^(2^127-1) first.
	//
	// The sequence of 10 multiplications and 126 squarings is derived from the
	// following addition chain generated with github.com/mmcloughlin/addchain v0.4.0.
	//
	//	_10      = 2*1
	//	_11      = 1 + _10
	//	_110     = 2*_11
	//	_111     = 1 + _110
	//	_111000  = _111 << 3
	//	_111111  = _111 + _111000
	//	_1111110 = 2*_111111
	//	_1111111 = 1 + _1111110
	//	x12      = _1111110 << 5 + _111111
	//	x24      = x12 << 12 + x12
	//	i36      = x24 << 7
	//	x31      = _1111111 + i36
	//	x48      = i36 << 17 + x24
	//	x96      = x48 << 48 + x48
	//	return     x96 << 31 + x31
	//
	var t0 = new(fiat.P224Element)
	var t1 = new(fiat.P224Element)

	r.Square(x)
	r.Mul(x, r)
	r.Square(r)
	r.Mul(x, r)
	t0.Square(r)
	for s := 1; s < 3; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:73
		_go_fuzz_dep_.CoverTab[2483]++
										t0.Square(t0)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:74
		// _ = "end of CoverTab[2483]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:75
	// _ = "end of CoverTab[2471]"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:75
	_go_fuzz_dep_.CoverTab[2472]++
									t0.Mul(r, t0)
									t1.Square(t0)
									r.Mul(x, t1)
									for s := 0; s < 5; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:79
		_go_fuzz_dep_.CoverTab[2484]++
										t1.Square(t1)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:80
		// _ = "end of CoverTab[2484]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:81
	// _ = "end of CoverTab[2472]"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:81
	_go_fuzz_dep_.CoverTab[2473]++
									t0.Mul(t0, t1)
									t1.Square(t0)
									for s := 1; s < 12; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:84
		_go_fuzz_dep_.CoverTab[2485]++
										t1.Square(t1)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:85
		// _ = "end of CoverTab[2485]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:86
	// _ = "end of CoverTab[2473]"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:86
	_go_fuzz_dep_.CoverTab[2474]++
									t0.Mul(t0, t1)
									t1.Square(t0)
									for s := 1; s < 7; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:89
		_go_fuzz_dep_.CoverTab[2486]++
										t1.Square(t1)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:90
		// _ = "end of CoverTab[2486]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:91
	// _ = "end of CoverTab[2474]"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:91
	_go_fuzz_dep_.CoverTab[2475]++
									r.Mul(r, t1)
									for s := 0; s < 17; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:93
		_go_fuzz_dep_.CoverTab[2487]++
										t1.Square(t1)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:94
		// _ = "end of CoverTab[2487]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:95
	// _ = "end of CoverTab[2475]"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:95
	_go_fuzz_dep_.CoverTab[2476]++
									t0.Mul(t0, t1)
									t1.Square(t0)
									for s := 1; s < 48; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:98
		_go_fuzz_dep_.CoverTab[2488]++
										t1.Square(t1)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:99
		// _ = "end of CoverTab[2488]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:100
	// _ = "end of CoverTab[2476]"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:100
	_go_fuzz_dep_.CoverTab[2477]++
									t0.Mul(t0, t1)
									for s := 0; s < 31; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:102
		_go_fuzz_dep_.CoverTab[2489]++
										t0.Square(t0)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:103
		// _ = "end of CoverTab[2489]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:104
	// _ = "end of CoverTab[2477]"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:104
	_go_fuzz_dep_.CoverTab[2478]++
									r.Mul(r, t0)

//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:108
	v := new(fiat.P224Element).Square(r)
									v.Mul(v, x)

//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:112
	r.Mul(r, x)

//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:120
	var p224MinusOne = new(fiat.P224Element).Sub(
		new(fiat.P224Element), new(fiat.P224Element).One())

	for i := 96 - 1; i >= 1; i-- {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:123
		_go_fuzz_dep_.CoverTab[2490]++
										w := new(fiat.P224Element).Set(v)
										for j := 0; j < i-1; j++ {
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:125
			_go_fuzz_dep_.CoverTab[2492]++
											w.Square(w)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:126
			// _ = "end of CoverTab[2492]"
		}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:127
		// _ = "end of CoverTab[2490]"
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:127
		_go_fuzz_dep_.CoverTab[2491]++
										cond := w.Equal(p224MinusOne)
										v.Select(t0.Mul(v, &p224GG[96-i]), v, cond)
										r.Select(t0.Mul(r, &p224GG[96-i-1]), r, cond)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:130
		// _ = "end of CoverTab[2491]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:131
	// _ = "end of CoverTab[2478]"
}

//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:132
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/nistec/p224_sqrt.go:132
var _ = _go_fuzz_dep_.CoverTab
