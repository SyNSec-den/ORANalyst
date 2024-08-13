// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the Go wrapper for the constant-time, 64-bit assembly
// implementation of P256. The optimizations performed here are described in
// detail in:
// S.Gueron and V.Krasnov, "Fast prime field elliptic-curve cryptography with
//                          256-bit primes"
// https://link.springer.com/article/10.1007%2Fs13389-014-0090-x
// https://eprint.iacr.org/2013/816.pdf

//go:build amd64 || arm64 || ppc64le || s390x

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:15
package nistec

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:15
import (
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:15
)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:15
import (
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:15
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:15
)

import (
	_ "embed"
	"encoding/binary"
	"errors"
	"math/bits"
	"runtime"
	"unsafe"
)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:28
type p256Element [4]uint64

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:31
var p256One = p256Element{0x0000000000000001, 0xffffffff00000000,
	0xffffffffffffffff, 0x00000000fffffffe}

var p256Zero = p256Element{}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:37
var p256P = p256Element{0xffffffffffffffff, 0x00000000ffffffff,
	0x0000000000000000, 0xffffffff00000001}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:42
type P256Point struct {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:45
	x, y, z p256Element
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:49
func NewP256Point() *P256Point {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:49
	_go_fuzz_dep_.CoverTab[2493]++
								return &P256Point{
		x:	p256One, y: p256One, z: p256Zero,
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:52
	// _ = "end of CoverTab[2493]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:56
func (p *P256Point) SetGenerator() *P256Point {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:56
	_go_fuzz_dep_.CoverTab[2494]++
								p.x = p256Element{0x79e730d418a9143c, 0x75ba95fc5fedb601,
		0x79fb732b77622510, 0x18905f76a53755c6}
	p.y = p256Element{0xddf25357ce95560a, 0x8b4ab8e4ba19e45c,
		0xd2e88688dd21f325, 0x8571ff1825885d85}
								p.z = p256One
								return p
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:62
	// _ = "end of CoverTab[2494]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:66
func (p *P256Point) Set(q *P256Point) *P256Point {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:66
	_go_fuzz_dep_.CoverTab[2495]++
								p.x, p.y, p.z = q.x, q.y, q.z
								return p
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:68
	// _ = "end of CoverTab[2495]"
}

const p256ElementLength = 32
const p256UncompressedLength = 1 + 2*p256ElementLength
const p256CompressedLength = 1 + p256ElementLength

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:79
func (p *P256Point) SetBytes(b []byte) (*P256Point, error) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:79
	_go_fuzz_dep_.CoverTab[2496]++

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:83
	rr := p256Element{0x0000000000000003, 0xfffffffbffffffff,
		0xfffffffffffffffe, 0x00000004fffffffd}

	switch {

	case len(b) == 1 && func() bool {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:88
		_go_fuzz_dep_.CoverTab[2505]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:88
		return b[0] == 0
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:88
		// _ = "end of CoverTab[2505]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:88
	}():
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:88
		_go_fuzz_dep_.CoverTab[2497]++
									return p.Set(NewP256Point()), nil
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:89
		// _ = "end of CoverTab[2497]"

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:92
	case len(b) == p256UncompressedLength && func() bool {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:92
		_go_fuzz_dep_.CoverTab[2506]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:92
		return b[0] == 4
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:92
		// _ = "end of CoverTab[2506]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:92
	}():
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:92
		_go_fuzz_dep_.CoverTab[2498]++
									var r P256Point
									p256BigToLittle(&r.x, (*[32]byte)(b[1:33]))
									p256BigToLittle(&r.y, (*[32]byte)(b[33:65]))
									if p256LessThanP(&r.x) == 0 || func() bool {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:96
			_go_fuzz_dep_.CoverTab[2507]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:96
			return p256LessThanP(&r.y) == 0
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:96
			// _ = "end of CoverTab[2507]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:96
		}() {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:96
			_go_fuzz_dep_.CoverTab[2508]++
										return nil, errors.New("invalid P256 element encoding")
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:97
			// _ = "end of CoverTab[2508]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:98
			_go_fuzz_dep_.CoverTab[2509]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:98
			// _ = "end of CoverTab[2509]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:98
		}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:98
		// _ = "end of CoverTab[2498]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:98
		_go_fuzz_dep_.CoverTab[2499]++
									p256Mul(&r.x, &r.x, &rr)
									p256Mul(&r.y, &r.y, &rr)
									if err := p256CheckOnCurve(&r.x, &r.y); err != nil {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:101
			_go_fuzz_dep_.CoverTab[2510]++
											return nil, err
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:102
			// _ = "end of CoverTab[2510]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:103
			_go_fuzz_dep_.CoverTab[2511]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:103
			// _ = "end of CoverTab[2511]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:103
		}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:103
		// _ = "end of CoverTab[2499]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:103
		_go_fuzz_dep_.CoverTab[2500]++
										r.z = p256One
										return p.Set(&r), nil
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:105
		// _ = "end of CoverTab[2500]"

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
	case len(b) == p256CompressedLength && func() bool {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
		_go_fuzz_dep_.CoverTab[2512]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
		return (b[0] == 2 || func() bool {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
			_go_fuzz_dep_.CoverTab[2513]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
			return b[0] == 3
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
			// _ = "end of CoverTab[2513]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
		}())
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
		// _ = "end of CoverTab[2512]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
	}():
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:108
		_go_fuzz_dep_.CoverTab[2501]++
										var r P256Point
										p256BigToLittle(&r.x, (*[32]byte)(b[1:33]))
										if p256LessThanP(&r.x) == 0 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:111
			_go_fuzz_dep_.CoverTab[2514]++
											return nil, errors.New("invalid P256 element encoding")
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:112
			// _ = "end of CoverTab[2514]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:113
			_go_fuzz_dep_.CoverTab[2515]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:113
			// _ = "end of CoverTab[2515]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:113
		}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:113
		// _ = "end of CoverTab[2501]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:113
		_go_fuzz_dep_.CoverTab[2502]++
										p256Mul(&r.x, &r.x, &rr)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:117
		p256Polynomial(&r.y, &r.x)
		if !p256Sqrt(&r.y, &r.y) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:118
			_go_fuzz_dep_.CoverTab[2516]++
											return nil, errors.New("invalid P256 compressed point encoding")
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:119
			// _ = "end of CoverTab[2516]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:120
			_go_fuzz_dep_.CoverTab[2517]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:120
			// _ = "end of CoverTab[2517]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:120
		}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:120
		// _ = "end of CoverTab[2502]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:120
		_go_fuzz_dep_.CoverTab[2503]++

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:124
		yy := new(p256Element)
										p256FromMont(yy, &r.y)
										cond := int(yy[0]&1) ^ int(b[0]&1)
										p256NegCond(&r.y, cond)

										r.z = p256One
										return p.Set(&r), nil
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:130
		// _ = "end of CoverTab[2503]"

	default:
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:132
		_go_fuzz_dep_.CoverTab[2504]++
										return nil, errors.New("invalid P256 point encoding")
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:133
		// _ = "end of CoverTab[2504]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:134
	// _ = "end of CoverTab[2496]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:138
func p256Polynomial(y2, x *p256Element) *p256Element {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:138
	_go_fuzz_dep_.CoverTab[2518]++
									x3 := new(p256Element)
									p256Sqr(x3, x, 1)
									p256Mul(x3, x3, x)

									threeX := new(p256Element)
									p256Add(threeX, x, x)
									p256Add(threeX, threeX, x)
									p256NegCond(threeX, 1)

									p256B := &p256Element{0xd89cdf6229c4bddf, 0xacf005cd78843090,
		0xe5a220abf7212ed6, 0xdc30061d04874834}

									p256Add(x3, x3, threeX)
									p256Add(x3, x3, p256B)

									*y2 = *x3
									return y2
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:155
	// _ = "end of CoverTab[2518]"
}

func p256CheckOnCurve(x, y *p256Element) error {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:158
	_go_fuzz_dep_.CoverTab[2519]++

									rhs := p256Polynomial(new(p256Element), x)
									lhs := new(p256Element)
									p256Sqr(lhs, y, 1)
									if p256Equal(lhs, rhs) != 1 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:163
		_go_fuzz_dep_.CoverTab[2521]++
										return errors.New("P256 point not on curve")
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:164
		// _ = "end of CoverTab[2521]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:165
		_go_fuzz_dep_.CoverTab[2522]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:165
		// _ = "end of CoverTab[2522]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:165
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:165
	// _ = "end of CoverTab[2519]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:165
	_go_fuzz_dep_.CoverTab[2520]++
									return nil
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:166
	// _ = "end of CoverTab[2520]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:172
func p256LessThanP(x *p256Element) int {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:172
	_go_fuzz_dep_.CoverTab[2523]++
									var b uint64
									_, b = bits.Sub64(x[0], p256P[0], b)
									_, b = bits.Sub64(x[1], p256P[1], b)
									_, b = bits.Sub64(x[2], p256P[2], b)
									_, b = bits.Sub64(x[3], p256P[3], b)
									return int(b)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:178
	// _ = "end of CoverTab[2523]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:182
func p256Add(res, x, y *p256Element) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:182
	_go_fuzz_dep_.CoverTab[2524]++
									var c, b uint64
									t1 := make([]uint64, 4)
									t1[0], c = bits.Add64(x[0], y[0], 0)
									t1[1], c = bits.Add64(x[1], y[1], c)
									t1[2], c = bits.Add64(x[2], y[2], c)
									t1[3], c = bits.Add64(x[3], y[3], c)
									t2 := make([]uint64, 4)
									t2[0], b = bits.Sub64(t1[0], p256P[0], 0)
									t2[1], b = bits.Sub64(t1[1], p256P[1], b)
									t2[2], b = bits.Sub64(t1[2], p256P[2], b)
									t2[3], b = bits.Sub64(t1[3], p256P[3], b)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:201
	t2Mask := (c ^ b) - 1
									res[0] = (t1[0] & ^t2Mask) | (t2[0] & t2Mask)
									res[1] = (t1[1] & ^t2Mask) | (t2[1] & t2Mask)
									res[2] = (t1[2] & ^t2Mask) | (t2[2] & t2Mask)
									res[3] = (t1[3] & ^t2Mask) | (t2[3] & t2Mask)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:205
	// _ = "end of CoverTab[2524]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:210
func p256Sqrt(e, x *p256Element) (isSquare bool) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:210
	_go_fuzz_dep_.CoverTab[2525]++
									t0, t1 := new(p256Element), new(p256Element)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:228
	p256Sqr(t0, x, 1)
	p256Mul(t0, x, t0)
	p256Sqr(t1, t0, 2)
	p256Mul(t0, t0, t1)
	p256Sqr(t1, t0, 4)
	p256Mul(t0, t0, t1)
	p256Sqr(t1, t0, 8)
	p256Mul(t0, t0, t1)
	p256Sqr(t1, t0, 16)
	p256Mul(t0, t0, t1)
	p256Sqr(t0, t0, 32)
	p256Mul(t0, x, t0)
	p256Sqr(t0, t0, 96)
	p256Mul(t0, x, t0)
	p256Sqr(t0, t0, 94)

	p256Sqr(t1, t0, 1)
	if p256Equal(t1, x) != 1 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:245
		_go_fuzz_dep_.CoverTab[2527]++
										return false
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:246
		// _ = "end of CoverTab[2527]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:247
		_go_fuzz_dep_.CoverTab[2528]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:247
		// _ = "end of CoverTab[2528]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:247
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:247
	// _ = "end of CoverTab[2525]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:247
	_go_fuzz_dep_.CoverTab[2526]++
									*e = *t0
									return true
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:249
	// _ = "end of CoverTab[2526]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:256
//go:noescape
func p256Mul(res, in1, in2 *p256Element)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:261
//go:noescape
func p256Sqr(res, in *p256Element, n int)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:267
//go:noescape
func p256FromMont(res, in *p256Element)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:272
//go:noescape
func p256NegCond(val *p256Element, cond int)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:277
//go:noescape
func p256MovCond(res, a, b *P256Point, cond int)

//go:noescape
func p256BigToLittle(res *p256Element, in *[32]byte)

//go:noescape
func p256LittleToBig(res *[32]byte, in *p256Element)

//go:noescape
func p256OrdBigToLittle(res *p256OrdElement, in *[32]byte)

//go:noescape
func p256OrdLittleToBig(res *[32]byte, in *p256OrdElement)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:295
type p256Table [16]P256Point

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:300
//go:noescape
func p256Select(res *P256Point, table *p256Table, idx int)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:305
type p256AffinePoint struct {
	x, y p256Element
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:311
type p256AffineTable [32]p256AffinePoint

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:320
var p256Precomputed *[43]p256AffineTable

//go:embed p256_asm_table.bin
var p256PrecomputedEmbed string

func init() {
	p256PrecomputedPtr := (*unsafe.Pointer)(unsafe.Pointer(&p256PrecomputedEmbed))
	if runtime.GOARCH == "s390x" {
		var newTable [43 * 32 * 2 * 4]uint64
		for i, x := range (*[43 * 32 * 2 * 4][8]byte)(*p256PrecomputedPtr) {
			newTable[i] = binary.LittleEndian.Uint64(x[:])
		}
		newTablePtr := unsafe.Pointer(&newTable)
		p256PrecomputedPtr = &newTablePtr
	}
	p256Precomputed = (*[43]p256AffineTable)(*p256PrecomputedPtr)
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:341
//go:noescape
func p256SelectAffine(res *p256AffinePoint, table *p256AffineTable, idx int)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:348
//go:noescape
func p256PointAddAffineAsm(res, in1 *P256Point, in2 *p256AffinePoint, sign, sel, zero int)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:355
//go:noescape
func p256PointAddAsm(res, in1, in2 *P256Point) int

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:360
//go:noescape
func p256PointDoubleAsm(res, in *P256Point)

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:365
type p256OrdElement [4]uint64

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:368
func p256OrdReduce(s *p256OrdElement) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:368
	_go_fuzz_dep_.CoverTab[2529]++

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:371
	t0, b := bits.Sub64(s[0], 0xf3b9cac2fc632551, 0)
									t1, b := bits.Sub64(s[1], 0xbce6faada7179e84, b)
									t2, b := bits.Sub64(s[2], 0xffffffffffffffff, b)
									t3, b := bits.Sub64(s[3], 0xffffffff00000000, b)
									tMask := b - 1
									s[0] ^= (t0 ^ s[0]) & tMask
									s[1] ^= (t1 ^ s[1]) & tMask
									s[2] ^= (t2 ^ s[2]) & tMask
									s[3] ^= (t3 ^ s[3]) & tMask
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:379
	// _ = "end of CoverTab[2529]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:383
func (q *P256Point) Add(r1, r2 *P256Point) *P256Point {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:383
	_go_fuzz_dep_.CoverTab[2530]++
									var sum, double P256Point
									r1IsInfinity := r1.isInfinity()
									r2IsInfinity := r2.isInfinity()
									pointsEqual := p256PointAddAsm(&sum, r1, r2)
									p256PointDoubleAsm(&double, r1)
									p256MovCond(&sum, &double, &sum, pointsEqual)
									p256MovCond(&sum, r1, &sum, r2IsInfinity)
									p256MovCond(&sum, r2, &sum, r1IsInfinity)
									return q.Set(&sum)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:392
	// _ = "end of CoverTab[2530]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:396
func (q *P256Point) Double(p *P256Point) *P256Point {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:396
	_go_fuzz_dep_.CoverTab[2531]++
									var double P256Point
									p256PointDoubleAsm(&double, p)
									return q.Set(&double)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:399
	// _ = "end of CoverTab[2531]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:405
func (r *P256Point) ScalarBaseMult(scalar []byte) (*P256Point, error) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:405
	_go_fuzz_dep_.CoverTab[2532]++
									if len(scalar) != 32 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:406
		_go_fuzz_dep_.CoverTab[2534]++
										return nil, errors.New("invalid scalar length")
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:407
		// _ = "end of CoverTab[2534]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:408
		_go_fuzz_dep_.CoverTab[2535]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:408
		// _ = "end of CoverTab[2535]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:408
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:408
	// _ = "end of CoverTab[2532]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:408
	_go_fuzz_dep_.CoverTab[2533]++
									scalarReversed := new(p256OrdElement)
									p256OrdBigToLittle(scalarReversed, (*[32]byte)(scalar))
									p256OrdReduce(scalarReversed)

									r.p256BaseMult(scalarReversed)
									return r, nil
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:414
	// _ = "end of CoverTab[2533]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:420
func (r *P256Point) ScalarMult(q *P256Point, scalar []byte) (*P256Point, error) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:420
	_go_fuzz_dep_.CoverTab[2536]++
									if len(scalar) != 32 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:421
		_go_fuzz_dep_.CoverTab[2538]++
										return nil, errors.New("invalid scalar length")
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:422
		// _ = "end of CoverTab[2538]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:423
		_go_fuzz_dep_.CoverTab[2539]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:423
		// _ = "end of CoverTab[2539]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:423
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:423
	// _ = "end of CoverTab[2536]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:423
	_go_fuzz_dep_.CoverTab[2537]++
									scalarReversed := new(p256OrdElement)
									p256OrdBigToLittle(scalarReversed, (*[32]byte)(scalar))
									p256OrdReduce(scalarReversed)

									r.Set(q).p256ScalarMult(scalarReversed)
									return r, nil
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:429
	// _ = "end of CoverTab[2537]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:433
func uint64IsZero(x uint64) int {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:433
	_go_fuzz_dep_.CoverTab[2540]++
									x = ^x
									x &= x >> 32
									x &= x >> 16
									x &= x >> 8
									x &= x >> 4
									x &= x >> 2
									x &= x >> 1
									return int(x & 1)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:441
	// _ = "end of CoverTab[2540]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:445
func p256Equal(a, b *p256Element) int {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:445
	_go_fuzz_dep_.CoverTab[2541]++
									var acc uint64
									for i := range a {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:447
		_go_fuzz_dep_.CoverTab[2543]++
										acc |= a[i] ^ b[i]
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:448
		// _ = "end of CoverTab[2543]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:449
	// _ = "end of CoverTab[2541]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:449
	_go_fuzz_dep_.CoverTab[2542]++
									return uint64IsZero(acc)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:450
	// _ = "end of CoverTab[2542]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:454
func (p *P256Point) isInfinity() int {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:454
	_go_fuzz_dep_.CoverTab[2544]++
									return p256Equal(&p.z, &p256Zero)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:455
	// _ = "end of CoverTab[2544]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:461
func (p *P256Point) Bytes() []byte {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:461
	_go_fuzz_dep_.CoverTab[2545]++

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:464
	var out [p256UncompressedLength]byte
									return p.bytes(&out)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:465
	// _ = "end of CoverTab[2545]"
}

func (p *P256Point) bytes(out *[p256UncompressedLength]byte) []byte {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:468
	_go_fuzz_dep_.CoverTab[2546]++

									if p.isInfinity() == 1 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:470
		_go_fuzz_dep_.CoverTab[2548]++
										return append(out[:0], 0)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:471
		// _ = "end of CoverTab[2548]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:472
		_go_fuzz_dep_.CoverTab[2549]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:472
		// _ = "end of CoverTab[2549]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:472
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:472
	// _ = "end of CoverTab[2546]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:472
	_go_fuzz_dep_.CoverTab[2547]++

									x, y := new(p256Element), new(p256Element)
									p.affineFromMont(x, y)

									out[0] = 4
									p256LittleToBig((*[32]byte)(out[1:33]), x)
									p256LittleToBig((*[32]byte)(out[33:65]), y)

									return out[:]
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:481
	// _ = "end of CoverTab[2547]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:486
func (p *P256Point) affineFromMont(x, y *p256Element) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:486
	_go_fuzz_dep_.CoverTab[2550]++
									p256Inverse(y, &p.z)
									p256Sqr(x, y, 1)
									p256Mul(y, y, x)

									p256Mul(x, &p.x, x)
									p256Mul(y, &p.y, y)

									p256FromMont(x, x)
									p256FromMont(y, y)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:495
	// _ = "end of CoverTab[2550]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:500
func (p *P256Point) BytesX() ([]byte, error) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:500
	_go_fuzz_dep_.CoverTab[2551]++

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:503
	var out [p256ElementLength]byte
									return p.bytesX(&out)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:504
	// _ = "end of CoverTab[2551]"
}

func (p *P256Point) bytesX(out *[p256ElementLength]byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:507
	_go_fuzz_dep_.CoverTab[2552]++
									if p.isInfinity() == 1 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:508
		_go_fuzz_dep_.CoverTab[2554]++
										return nil, errors.New("P256 point is the point at infinity")
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:509
		// _ = "end of CoverTab[2554]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:510
		_go_fuzz_dep_.CoverTab[2555]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:510
		// _ = "end of CoverTab[2555]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:510
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:510
	// _ = "end of CoverTab[2552]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:510
	_go_fuzz_dep_.CoverTab[2553]++

									x := new(p256Element)
									p256Inverse(x, &p.z)
									p256Sqr(x, x, 1)
									p256Mul(x, &p.x, x)
									p256FromMont(x, x)
									p256LittleToBig((*[32]byte)(out[:]), x)

									return out[:], nil
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:519
	// _ = "end of CoverTab[2553]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:525
func (p *P256Point) BytesCompressed() []byte {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:525
	_go_fuzz_dep_.CoverTab[2556]++

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:528
	var out [p256CompressedLength]byte
									return p.bytesCompressed(&out)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:529
	// _ = "end of CoverTab[2556]"
}

func (p *P256Point) bytesCompressed(out *[p256CompressedLength]byte) []byte {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:532
	_go_fuzz_dep_.CoverTab[2557]++
									if p.isInfinity() == 1 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:533
		_go_fuzz_dep_.CoverTab[2559]++
										return append(out[:0], 0)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:534
		// _ = "end of CoverTab[2559]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:535
		_go_fuzz_dep_.CoverTab[2560]++
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:535
		// _ = "end of CoverTab[2560]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:535
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:535
	// _ = "end of CoverTab[2557]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:535
	_go_fuzz_dep_.CoverTab[2558]++

									x, y := new(p256Element), new(p256Element)
									p.affineFromMont(x, y)

									out[0] = 2 | byte(y[0]&1)
									p256LittleToBig((*[32]byte)(out[1:33]), x)

									return out[:]
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:543
	// _ = "end of CoverTab[2558]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:547
func (q *P256Point) Select(p1, p2 *P256Point, cond int) *P256Point {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:547
	_go_fuzz_dep_.CoverTab[2561]++
									p256MovCond(q, p1, p2, cond)
									return q
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:549
	// _ = "end of CoverTab[2561]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:553
func p256Inverse(out, in *p256Element) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:553
	_go_fuzz_dep_.CoverTab[2562]++

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:576
	var z = new(p256Element)
									var t0 = new(p256Element)
									var t1 = new(p256Element)

									p256Sqr(z, in, 1)
									p256Mul(z, in, z)
									p256Sqr(z, z, 1)
									p256Mul(z, in, z)
									p256Sqr(t0, z, 3)
									p256Mul(t0, z, t0)
									p256Sqr(t1, t0, 6)
									p256Mul(t0, t0, t1)
									p256Sqr(t0, t0, 3)
									p256Mul(z, z, t0)
									p256Sqr(t0, z, 1)
									p256Mul(t0, in, t0)
									p256Sqr(t1, t0, 16)
									p256Mul(t0, t0, t1)
									p256Sqr(t0, t0, 15)
									p256Mul(z, z, t0)
									p256Sqr(t0, t0, 17)
									p256Mul(t0, in, t0)
									p256Sqr(t0, t0, 143)
									p256Mul(t0, z, t0)
									p256Sqr(t0, t0, 47)
									p256Mul(z, z, t0)
									p256Sqr(z, z, 2)
									p256Mul(out, in, z)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:603
	// _ = "end of CoverTab[2562]"
}

func boothW5(in uint) (int, int) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:606
	_go_fuzz_dep_.CoverTab[2563]++
									var s uint = ^((in >> 5) - 1)
									var d uint = (1 << 6) - in - 1
									d = (d & s) | (in & (^s))
									d = (d >> 1) + (d & 1)
									return int(d), int(s & 1)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:611
	// _ = "end of CoverTab[2563]"
}

func boothW6(in uint) (int, int) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:614
	_go_fuzz_dep_.CoverTab[2564]++
									var s uint = ^((in >> 6) - 1)
									var d uint = (1 << 7) - in - 1
									d = (d & s) | (in & (^s))
									d = (d >> 1) + (d & 1)
									return int(d), int(s & 1)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:619
	// _ = "end of CoverTab[2564]"
}

func (p *P256Point) p256BaseMult(scalar *p256OrdElement) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:622
	_go_fuzz_dep_.CoverTab[2565]++
									var t0 p256AffinePoint

									wvalue := (scalar[0] << 1) & 0x7f
									sel, sign := boothW6(uint(wvalue))
									p256SelectAffine(&t0, &p256Precomputed[0], sel)
									p.x, p.y, p.z = t0.x, t0.y, p256One
									p256NegCond(&p.y, sign)

									index := uint(5)
									zero := sel

									for i := 1; i < 43; i++ {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:634
		_go_fuzz_dep_.CoverTab[2567]++
										if index < 192 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:635
			_go_fuzz_dep_.CoverTab[2569]++
											wvalue = ((scalar[index/64] >> (index % 64)) + (scalar[index/64+1] << (64 - (index % 64)))) & 0x7f
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:636
			// _ = "end of CoverTab[2569]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:637
			_go_fuzz_dep_.CoverTab[2570]++
											wvalue = (scalar[index/64] >> (index % 64)) & 0x7f
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:638
			// _ = "end of CoverTab[2570]"
		}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:639
		// _ = "end of CoverTab[2567]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:639
		_go_fuzz_dep_.CoverTab[2568]++
										index += 6
										sel, sign = boothW6(uint(wvalue))
										p256SelectAffine(&t0, &p256Precomputed[i], sel)
										p256PointAddAffineAsm(p, p, &t0, sign, sel, zero)
										zero |= sel
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:644
		// _ = "end of CoverTab[2568]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:645
	// _ = "end of CoverTab[2565]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:645
	_go_fuzz_dep_.CoverTab[2566]++

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:648
	p256MovCond(p, p, NewP256Point(), zero)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:648
	// _ = "end of CoverTab[2566]"
}

func (p *P256Point) p256ScalarMult(scalar *p256OrdElement) {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:651
	_go_fuzz_dep_.CoverTab[2571]++

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:654
	var precomp p256Table
									var t0, t1, t2, t3 P256Point

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:658
	precomp[0] = *p

									p256PointDoubleAsm(&t0, p)
									p256PointDoubleAsm(&t1, &t0)
									p256PointDoubleAsm(&t2, &t1)
									p256PointDoubleAsm(&t3, &t2)
									precomp[1] = t0
									precomp[3] = t1
									precomp[7] = t2
									precomp[15] = t3

									p256PointAddAsm(&t0, &t0, p)
									p256PointAddAsm(&t1, &t1, p)
									p256PointAddAsm(&t2, &t2, p)
									precomp[2] = t0
									precomp[4] = t1
									precomp[8] = t2

									p256PointDoubleAsm(&t0, &t0)
									p256PointDoubleAsm(&t1, &t1)
									precomp[5] = t0
									precomp[9] = t1

									p256PointAddAsm(&t2, &t0, p)
									p256PointAddAsm(&t1, &t1, p)
									precomp[6] = t2
									precomp[10] = t1

									p256PointDoubleAsm(&t0, &t0)
									p256PointDoubleAsm(&t2, &t2)
									precomp[11] = t0
									precomp[13] = t2

									p256PointAddAsm(&t0, &t0, p)
									p256PointAddAsm(&t2, &t2, p)
									precomp[12] = t0
									precomp[14] = t2

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:697
	index := uint(254)
	var sel, sign int

	wvalue := (scalar[index/64] >> (index % 64)) & 0x3f
	sel, _ = boothW5(uint(wvalue))

	p256Select(p, &precomp, sel)
	zero := sel

	for index > 4 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:706
		_go_fuzz_dep_.CoverTab[2573]++
										index -= 5
										p256PointDoubleAsm(p, p)
										p256PointDoubleAsm(p, p)
										p256PointDoubleAsm(p, p)
										p256PointDoubleAsm(p, p)
										p256PointDoubleAsm(p, p)

										if index < 192 {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:714
			_go_fuzz_dep_.CoverTab[2575]++
											wvalue = ((scalar[index/64] >> (index % 64)) + (scalar[index/64+1] << (64 - (index % 64)))) & 0x3f
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:715
			// _ = "end of CoverTab[2575]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:716
			_go_fuzz_dep_.CoverTab[2576]++
											wvalue = (scalar[index/64] >> (index % 64)) & 0x3f
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:717
			// _ = "end of CoverTab[2576]"
		}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:718
		// _ = "end of CoverTab[2573]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:718
		_go_fuzz_dep_.CoverTab[2574]++

										sel, sign = boothW5(uint(wvalue))

										p256Select(&t0, &precomp, sel)
										p256NegCond(&t0.y, sign)
										p256PointAddAsm(&t1, p, &t0)
										p256MovCond(&t1, &t1, p, sel)
										p256MovCond(p, &t1, &t0, zero)
										zero |= sel
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:727
		// _ = "end of CoverTab[2574]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:728
	// _ = "end of CoverTab[2571]"
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:728
	_go_fuzz_dep_.CoverTab[2572]++

									p256PointDoubleAsm(p, p)
									p256PointDoubleAsm(p, p)
									p256PointDoubleAsm(p, p)
									p256PointDoubleAsm(p, p)
									p256PointDoubleAsm(p, p)

									wvalue = (scalar[0] << 1) & 0x3f
									sel, sign = boothW5(uint(wvalue))

									p256Select(&t0, &precomp, sel)
									p256NegCond(&t0.y, sign)
									p256PointAddAsm(&t1, p, &t0)
									p256MovCond(&t1, &t1, p, sel)
									p256MovCond(p, &t1, &t0, zero)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:743
	// _ = "end of CoverTab[2572]"
}

//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:744
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/nistec/p256_asm.go:744
var _ = _go_fuzz_dep_.CoverTab
