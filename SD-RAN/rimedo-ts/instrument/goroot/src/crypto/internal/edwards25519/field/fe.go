// Copyright (c) 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:5
// Package field implements fast arithmetic modulo 2^255-19.
package field

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:6
import (
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:6
)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:6
import (
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:6
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:6
)

import (
	"crypto/subtle"
	"encoding/binary"
	"errors"
	"math/bits"
)

// Element represents an element of the field GF(2^255-19). Note that this
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:15
// is not a cryptographically secure group, and should only be used to interact
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:15
// with edwards25519.Point coordinates.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:15
//
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:15
// This type works similarly to math/big.Int, and all arguments and receivers
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:15
// are allowed to alias.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:15
//
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:15
// The zero value is a valid zero element.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:23
type Element struct {
	// An element t represents the integer
	//     t.l0 + t.l1*2^51 + t.l2*2^102 + t.l3*2^153 + t.l4*2^204
	//
	// Between operations, all limbs are expected to be lower than 2^52.
	l0	uint64
	l1	uint64
	l2	uint64
	l3	uint64
	l4	uint64
}

const maskLow51Bits uint64 = (1 << 51) - 1

var feZero = &Element{0, 0, 0, 0, 0}

// Zero sets v = 0, and returns v.
func (v *Element) Zero() *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:40
	_go_fuzz_dep_.CoverTab[1986]++
									*v = *feZero
									return v
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:42
	// _ = "end of CoverTab[1986]"
}

var feOne = &Element{1, 0, 0, 0, 0}

// One sets v = 1, and returns v.
func (v *Element) One() *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:48
	_go_fuzz_dep_.CoverTab[1987]++
									*v = *feOne
									return v
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:50
	// _ = "end of CoverTab[1987]"
}

// reduce reduces v modulo 2^255 - 19 and returns it.
func (v *Element) reduce() *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:54
	_go_fuzz_dep_.CoverTab[1988]++
									v.carryPropagate()

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:62
	c := (v.l0 + 19) >> 51
									c = (v.l1 + c) >> 51
									c = (v.l2 + c) >> 51
									c = (v.l3 + c) >> 51
									c = (v.l4 + c) >> 51

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:70
	v.l0 += 19 * c

									v.l1 += v.l0 >> 51
									v.l0 = v.l0 & maskLow51Bits
									v.l2 += v.l1 >> 51
									v.l1 = v.l1 & maskLow51Bits
									v.l3 += v.l2 >> 51
									v.l2 = v.l2 & maskLow51Bits
									v.l4 += v.l3 >> 51
									v.l3 = v.l3 & maskLow51Bits

									v.l4 = v.l4 & maskLow51Bits

									return v
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:83
	// _ = "end of CoverTab[1988]"
}

// Add sets v = a + b, and returns v.
func (v *Element) Add(a, b *Element) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:87
	_go_fuzz_dep_.CoverTab[1989]++
									v.l0 = a.l0 + b.l0
									v.l1 = a.l1 + b.l1
									v.l2 = a.l2 + b.l2
									v.l3 = a.l3 + b.l3
									v.l4 = a.l4 + b.l4

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:97
	return v.carryPropagateGeneric()
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:97
	// _ = "end of CoverTab[1989]"
}

// Subtract sets v = a - b, and returns v.
func (v *Element) Subtract(a, b *Element) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:101
	_go_fuzz_dep_.CoverTab[1990]++

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:104
	v.l0 = (a.l0 + 0xFFFFFFFFFFFDA) - b.l0
									v.l1 = (a.l1 + 0xFFFFFFFFFFFFE) - b.l1
									v.l2 = (a.l2 + 0xFFFFFFFFFFFFE) - b.l2
									v.l3 = (a.l3 + 0xFFFFFFFFFFFFE) - b.l3
									v.l4 = (a.l4 + 0xFFFFFFFFFFFFE) - b.l4
									return v.carryPropagate()
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:109
	// _ = "end of CoverTab[1990]"
}

// Negate sets v = -a, and returns v.
func (v *Element) Negate(a *Element) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:113
	_go_fuzz_dep_.CoverTab[1991]++
									return v.Subtract(feZero, a)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:114
	// _ = "end of CoverTab[1991]"
}

// Invert sets v = 1/z mod p, and returns v.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:117
//
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:117
// If z == 0, Invert returns v = 0.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:120
func (v *Element) Invert(z *Element) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:120
	_go_fuzz_dep_.CoverTab[1992]++
	// Inversion is implemented as exponentiation with exponent p − 2. It uses the
	// same sequence of 255 squarings and 11 multiplications as [Curve25519].
	var z2, z9, z11, z2_5_0, z2_10_0, z2_20_0, z2_50_0, z2_100_0, t Element

	z2.Square(z)
	t.Square(&z2)
	t.Square(&t)
	z9.Multiply(&t, z)
	z11.Multiply(&z9, &z2)
	t.Square(&z11)
	z2_5_0.Multiply(&t, &z9)

	t.Square(&z2_5_0)
	for i := 0; i < 4; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:134
		_go_fuzz_dep_.CoverTab[2000]++
										t.Square(&t)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:135
		// _ = "end of CoverTab[2000]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:136
	// _ = "end of CoverTab[1992]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:136
	_go_fuzz_dep_.CoverTab[1993]++
									z2_10_0.Multiply(&t, &z2_5_0)

									t.Square(&z2_10_0)
									for i := 0; i < 9; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:140
		_go_fuzz_dep_.CoverTab[2001]++
										t.Square(&t)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:141
		// _ = "end of CoverTab[2001]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:142
	// _ = "end of CoverTab[1993]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:142
	_go_fuzz_dep_.CoverTab[1994]++
									z2_20_0.Multiply(&t, &z2_10_0)

									t.Square(&z2_20_0)
									for i := 0; i < 19; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:146
		_go_fuzz_dep_.CoverTab[2002]++
										t.Square(&t)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:147
		// _ = "end of CoverTab[2002]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:148
	// _ = "end of CoverTab[1994]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:148
	_go_fuzz_dep_.CoverTab[1995]++
									t.Multiply(&t, &z2_20_0)

									t.Square(&t)
									for i := 0; i < 9; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:152
		_go_fuzz_dep_.CoverTab[2003]++
										t.Square(&t)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:153
		// _ = "end of CoverTab[2003]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:154
	// _ = "end of CoverTab[1995]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:154
	_go_fuzz_dep_.CoverTab[1996]++
									z2_50_0.Multiply(&t, &z2_10_0)

									t.Square(&z2_50_0)
									for i := 0; i < 49; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:158
		_go_fuzz_dep_.CoverTab[2004]++
										t.Square(&t)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:159
		// _ = "end of CoverTab[2004]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:160
	// _ = "end of CoverTab[1996]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:160
	_go_fuzz_dep_.CoverTab[1997]++
									z2_100_0.Multiply(&t, &z2_50_0)

									t.Square(&z2_100_0)
									for i := 0; i < 99; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:164
		_go_fuzz_dep_.CoverTab[2005]++
										t.Square(&t)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:165
		// _ = "end of CoverTab[2005]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:166
	// _ = "end of CoverTab[1997]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:166
	_go_fuzz_dep_.CoverTab[1998]++
									t.Multiply(&t, &z2_100_0)

									t.Square(&t)
									for i := 0; i < 49; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:170
		_go_fuzz_dep_.CoverTab[2006]++
										t.Square(&t)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:171
		// _ = "end of CoverTab[2006]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:172
	// _ = "end of CoverTab[1998]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:172
	_go_fuzz_dep_.CoverTab[1999]++
									t.Multiply(&t, &z2_50_0)

									t.Square(&t)
									t.Square(&t)
									t.Square(&t)
									t.Square(&t)
									t.Square(&t)

									return v.Multiply(&t, &z11)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:181
	// _ = "end of CoverTab[1999]"
}

// Set sets v = a, and returns v.
func (v *Element) Set(a *Element) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:185
	_go_fuzz_dep_.CoverTab[2007]++
									*v = *a
									return v
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:187
	// _ = "end of CoverTab[2007]"
}

// SetBytes sets v to x, where x is a 32-byte little-endian encoding. If x is
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:190
// not of the right length, SetBytes returns nil and an error, and the
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:190
// receiver is unchanged.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:190
//
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:190
// Consistent with RFC 7748, the most significant bit (the high bit of the
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:190
// last byte) is ignored, and non-canonical values (2^255-19 through 2^255-1)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:190
// are accepted. Note that this is laxer than specified by RFC 8032, but
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:190
// consistent with most Ed25519 implementations.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:198
func (v *Element) SetBytes(x []byte) (*Element, error) {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:198
	_go_fuzz_dep_.CoverTab[2008]++
									if len(x) != 32 {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:199
		_go_fuzz_dep_.CoverTab[2010]++
										return nil, errors.New("edwards25519: invalid field element input size")
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:200
		// _ = "end of CoverTab[2010]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:201
		_go_fuzz_dep_.CoverTab[2011]++
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:201
		// _ = "end of CoverTab[2011]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:201
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:201
	// _ = "end of CoverTab[2008]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:201
	_go_fuzz_dep_.CoverTab[2009]++

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:204
	v.l0 = binary.LittleEndian.Uint64(x[0:8])
									v.l0 &= maskLow51Bits

									v.l1 = binary.LittleEndian.Uint64(x[6:14]) >> 3
									v.l1 &= maskLow51Bits

									v.l2 = binary.LittleEndian.Uint64(x[12:20]) >> 6
									v.l2 &= maskLow51Bits

									v.l3 = binary.LittleEndian.Uint64(x[19:27]) >> 1
									v.l3 &= maskLow51Bits

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:217
	v.l4 = binary.LittleEndian.Uint64(x[24:32]) >> 12
									v.l4 &= maskLow51Bits

									return v, nil
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:220
	// _ = "end of CoverTab[2009]"
}

// Bytes returns the canonical 32-byte little-endian encoding of v.
func (v *Element) Bytes() []byte {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:224
	_go_fuzz_dep_.CoverTab[2012]++
	// This function is outlined to make the allocations inline in the caller
									// rather than happen on the heap.
									var out [32]byte
									return v.bytes(&out)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:228
	// _ = "end of CoverTab[2012]"
}

func (v *Element) bytes(out *[32]byte) []byte {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:231
	_go_fuzz_dep_.CoverTab[2013]++
									t := *v
									t.reduce()

									var buf [8]byte
									for i, l := range [5]uint64{t.l0, t.l1, t.l2, t.l3, t.l4} {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:236
		_go_fuzz_dep_.CoverTab[2015]++
										bitsOffset := i * 51
										binary.LittleEndian.PutUint64(buf[:], l<<uint(bitsOffset%8))
										for i, bb := range buf {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:239
			_go_fuzz_dep_.CoverTab[2016]++
											off := bitsOffset/8 + i
											if off >= len(out) {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:241
				_go_fuzz_dep_.CoverTab[2018]++
												break
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:242
				// _ = "end of CoverTab[2018]"
			} else {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:243
				_go_fuzz_dep_.CoverTab[2019]++
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:243
				// _ = "end of CoverTab[2019]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:243
			}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:243
			// _ = "end of CoverTab[2016]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:243
			_go_fuzz_dep_.CoverTab[2017]++
											out[off] |= bb
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:244
			// _ = "end of CoverTab[2017]"
		}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:245
		// _ = "end of CoverTab[2015]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:246
	// _ = "end of CoverTab[2013]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:246
	_go_fuzz_dep_.CoverTab[2014]++

									return out[:]
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:248
	// _ = "end of CoverTab[2014]"
}

// Equal returns 1 if v and u are equal, and 0 otherwise.
func (v *Element) Equal(u *Element) int {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:252
	_go_fuzz_dep_.CoverTab[2020]++
									sa, sv := u.Bytes(), v.Bytes()
									return subtle.ConstantTimeCompare(sa, sv)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:254
	// _ = "end of CoverTab[2020]"
}

// mask64Bits returns 0xffffffff if cond is 1, and 0 otherwise.
func mask64Bits(cond int) uint64 {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:258
	_go_fuzz_dep_.CoverTab[2021]++
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:258
	return ^(uint64(cond) - 1)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:258
	// _ = "end of CoverTab[2021]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:258
}

// Select sets v to a if cond == 1, and to b if cond == 0.
func (v *Element) Select(a, b *Element, cond int) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:261
	_go_fuzz_dep_.CoverTab[2022]++
									m := mask64Bits(cond)
									v.l0 = (m & a.l0) | (^m & b.l0)
									v.l1 = (m & a.l1) | (^m & b.l1)
									v.l2 = (m & a.l2) | (^m & b.l2)
									v.l3 = (m & a.l3) | (^m & b.l3)
									v.l4 = (m & a.l4) | (^m & b.l4)
									return v
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:268
	// _ = "end of CoverTab[2022]"
}

// Swap swaps v and u if cond == 1 or leaves them unchanged if cond == 0, and returns v.
func (v *Element) Swap(u *Element, cond int) {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:272
	_go_fuzz_dep_.CoverTab[2023]++
									m := mask64Bits(cond)
									t := m & (v.l0 ^ u.l0)
									v.l0 ^= t
									u.l0 ^= t
									t = m & (v.l1 ^ u.l1)
									v.l1 ^= t
									u.l1 ^= t
									t = m & (v.l2 ^ u.l2)
									v.l2 ^= t
									u.l2 ^= t
									t = m & (v.l3 ^ u.l3)
									v.l3 ^= t
									u.l3 ^= t
									t = m & (v.l4 ^ u.l4)
									v.l4 ^= t
									u.l4 ^= t
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:288
	// _ = "end of CoverTab[2023]"
}

// IsNegative returns 1 if v is negative, and 0 otherwise.
func (v *Element) IsNegative() int {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:292
	_go_fuzz_dep_.CoverTab[2024]++
									return int(v.Bytes()[0] & 1)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:293
	// _ = "end of CoverTab[2024]"
}

// Absolute sets v to |u|, and returns v.
func (v *Element) Absolute(u *Element) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:297
	_go_fuzz_dep_.CoverTab[2025]++
									return v.Select(new(Element).Negate(u), u, u.IsNegative())
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:298
	// _ = "end of CoverTab[2025]"
}

// Multiply sets v = x * y, and returns v.
func (v *Element) Multiply(x, y *Element) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:302
	_go_fuzz_dep_.CoverTab[2026]++
									feMul(v, x, y)
									return v
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:304
	// _ = "end of CoverTab[2026]"
}

// Square sets v = x * x, and returns v.
func (v *Element) Square(x *Element) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:308
	_go_fuzz_dep_.CoverTab[2027]++
									feSquare(v, x)
									return v
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:310
	// _ = "end of CoverTab[2027]"
}

// Mult32 sets v = x * y, and returns v.
func (v *Element) Mult32(x *Element, y uint32) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:314
	_go_fuzz_dep_.CoverTab[2028]++
									x0lo, x0hi := mul51(x.l0, y)
									x1lo, x1hi := mul51(x.l1, y)
									x2lo, x2hi := mul51(x.l2, y)
									x3lo, x3hi := mul51(x.l3, y)
									x4lo, x4hi := mul51(x.l4, y)
									v.l0 = x0lo + 19*x4hi
									v.l1 = x1lo + x0hi
									v.l2 = x2lo + x1hi
									v.l3 = x3lo + x2hi
									v.l4 = x4lo + x3hi

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:327
	return v
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:327
	// _ = "end of CoverTab[2028]"
}

// mul51 returns lo + hi * 2⁵¹ = a * b.
func mul51(a uint64, b uint32) (lo uint64, hi uint64) {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:331
	_go_fuzz_dep_.CoverTab[2029]++
									mh, ml := bits.Mul64(a, uint64(b))
									lo = ml & maskLow51Bits
									hi = (mh << 13) | (ml >> 51)
									return
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:335
	// _ = "end of CoverTab[2029]"
}

// Pow22523 set v = x^((p-5)/8), and returns v. (p-5)/8 is 2^252-3.
func (v *Element) Pow22523(x *Element) *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:339
	_go_fuzz_dep_.CoverTab[2030]++
									var t0, t1, t2 Element

									t0.Square(x)
									t1.Square(&t0)
									t1.Square(&t1)
									t1.Multiply(x, &t1)
									t0.Multiply(&t0, &t1)
									t0.Square(&t0)
									t0.Multiply(&t1, &t0)
									t1.Square(&t0)
									for i := 1; i < 5; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:350
		_go_fuzz_dep_.CoverTab[2038]++
										t1.Square(&t1)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:351
		// _ = "end of CoverTab[2038]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:352
	// _ = "end of CoverTab[2030]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:352
	_go_fuzz_dep_.CoverTab[2031]++
									t0.Multiply(&t1, &t0)
									t1.Square(&t0)
									for i := 1; i < 10; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:355
		_go_fuzz_dep_.CoverTab[2039]++
										t1.Square(&t1)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:356
		// _ = "end of CoverTab[2039]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:357
	// _ = "end of CoverTab[2031]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:357
	_go_fuzz_dep_.CoverTab[2032]++
									t1.Multiply(&t1, &t0)
									t2.Square(&t1)
									for i := 1; i < 20; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:360
		_go_fuzz_dep_.CoverTab[2040]++
										t2.Square(&t2)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:361
		// _ = "end of CoverTab[2040]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:362
	// _ = "end of CoverTab[2032]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:362
	_go_fuzz_dep_.CoverTab[2033]++
									t1.Multiply(&t2, &t1)
									t1.Square(&t1)
									for i := 1; i < 10; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:365
		_go_fuzz_dep_.CoverTab[2041]++
										t1.Square(&t1)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:366
		// _ = "end of CoverTab[2041]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:367
	// _ = "end of CoverTab[2033]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:367
	_go_fuzz_dep_.CoverTab[2034]++
									t0.Multiply(&t1, &t0)
									t1.Square(&t0)
									for i := 1; i < 50; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:370
		_go_fuzz_dep_.CoverTab[2042]++
										t1.Square(&t1)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:371
		// _ = "end of CoverTab[2042]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:372
	// _ = "end of CoverTab[2034]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:372
	_go_fuzz_dep_.CoverTab[2035]++
									t1.Multiply(&t1, &t0)
									t2.Square(&t1)
									for i := 1; i < 100; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:375
		_go_fuzz_dep_.CoverTab[2043]++
										t2.Square(&t2)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:376
		// _ = "end of CoverTab[2043]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:377
	// _ = "end of CoverTab[2035]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:377
	_go_fuzz_dep_.CoverTab[2036]++
									t1.Multiply(&t2, &t1)
									t1.Square(&t1)
									for i := 1; i < 50; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:380
		_go_fuzz_dep_.CoverTab[2044]++
										t1.Square(&t1)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:381
		// _ = "end of CoverTab[2044]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:382
	// _ = "end of CoverTab[2036]"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:382
	_go_fuzz_dep_.CoverTab[2037]++
									t0.Multiply(&t1, &t0)
									t0.Square(&t0)
									t0.Square(&t0)
									return v.Multiply(&t0, x)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:386
	// _ = "end of CoverTab[2037]"
}

// sqrtM1 is 2^((p-1)/4), which squared is equal to -1 by Euler's Criterion.
var sqrtM1 = &Element{1718705420411056, 234908883556509,
	2233514472574048, 2117202627021982, 765476049583133}

// SqrtRatio sets r to the non-negative square root of the ratio of u and v.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:393
//
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:393
// If u/v is square, SqrtRatio returns r and 1. If u/v is not square, SqrtRatio
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:393
// sets r according to Section 4.3 of draft-irtf-cfrg-ristretto255-decaf448-00,
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:393
// and returns r and 0.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:398
func (r *Element) SqrtRatio(u, v *Element) (R *Element, wasSquare int) {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:398
	_go_fuzz_dep_.CoverTab[2045]++
									t0 := new(Element)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:402
	v2 := new(Element).Square(v)
									uv3 := new(Element).Multiply(u, t0.Multiply(v2, v))
									uv7 := new(Element).Multiply(uv3, t0.Square(v2))
									rr := new(Element).Multiply(uv3, t0.Pow22523(uv7))

									check := new(Element).Multiply(v, t0.Square(rr))

									uNeg := new(Element).Negate(u)
									correctSignSqrt := check.Equal(u)
									flippedSignSqrt := check.Equal(uNeg)
									flippedSignSqrtI := check.Equal(t0.Multiply(uNeg, sqrtM1))

									rPrime := new(Element).Multiply(rr, sqrtM1)

									rr.Select(rPrime, rr, flippedSignSqrt|flippedSignSqrtI)

									r.Absolute(rr)
									return r, correctSignSqrt | flippedSignSqrt
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:419
	// _ = "end of CoverTab[2045]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:420
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe.go:420
var _ = _go_fuzz_dep_.CoverTab
