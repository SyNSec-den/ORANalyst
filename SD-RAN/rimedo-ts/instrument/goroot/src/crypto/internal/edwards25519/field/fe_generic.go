// Copyright (c) 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:5
package field

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:5
)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:5
)

import "math/bits"

// uint128 holds a 128-bit number as two 64-bit limbs, for use with the
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:9
// bits.Mul64 and bits.Add64 intrinsics.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:11
type uint128 struct {
	lo, hi uint64
}

// mul64 returns a * b.
func mul64(a, b uint64) uint128 {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:16
	_go_fuzz_dep_.CoverTab[2047]++
										hi, lo := bits.Mul64(a, b)
										return uint128{lo, hi}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:18
	// _ = "end of CoverTab[2047]"
}

// addMul64 returns v + a * b.
func addMul64(v uint128, a, b uint64) uint128 {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:22
	_go_fuzz_dep_.CoverTab[2048]++
										hi, lo := bits.Mul64(a, b)
										lo, c := bits.Add64(lo, v.lo, 0)
										hi, _ = bits.Add64(hi, v.hi, c)
										return uint128{lo, hi}
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:26
	// _ = "end of CoverTab[2048]"
}

// shiftRightBy51 returns a >> 51. a is assumed to be at most 115 bits.
func shiftRightBy51(a uint128) uint64 {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:30
	_go_fuzz_dep_.CoverTab[2049]++
										return (a.hi << (64 - 51)) | (a.lo >> 51)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:31
	// _ = "end of CoverTab[2049]"
}

func feMulGeneric(v, a, b *Element) {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:34
	_go_fuzz_dep_.CoverTab[2050]++
										a0 := a.l0
										a1 := a.l1
										a2 := a.l2
										a3 := a.l3
										a4 := a.l4

										b0 := b.l0
										b1 := b.l1
										b2 := b.l2
										b3 := b.l3
										b4 := b.l4

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:79
	a1_19 := a1 * 19
										a2_19 := a2 * 19
										a3_19 := a3 * 19
										a4_19 := a4 * 19

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:85
	r0 := mul64(a0, b0)
										r0 = addMul64(r0, a1_19, b4)
										r0 = addMul64(r0, a2_19, b3)
										r0 = addMul64(r0, a3_19, b2)
										r0 = addMul64(r0, a4_19, b1)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:92
	r1 := mul64(a0, b1)
										r1 = addMul64(r1, a1, b0)
										r1 = addMul64(r1, a2_19, b4)
										r1 = addMul64(r1, a3_19, b3)
										r1 = addMul64(r1, a4_19, b2)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:99
	r2 := mul64(a0, b2)
										r2 = addMul64(r2, a1, b1)
										r2 = addMul64(r2, a2, b0)
										r2 = addMul64(r2, a3_19, b4)
										r2 = addMul64(r2, a4_19, b3)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:106
	r3 := mul64(a0, b3)
										r3 = addMul64(r3, a1, b2)
										r3 = addMul64(r3, a2, b1)
										r3 = addMul64(r3, a3, b0)
										r3 = addMul64(r3, a4_19, b4)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:113
	r4 := mul64(a0, b4)
										r4 = addMul64(r4, a1, b3)
										r4 = addMul64(r4, a2, b2)
										r4 = addMul64(r4, a3, b1)
										r4 = addMul64(r4, a4, b0)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:146
	c0 := shiftRightBy51(r0)
										c1 := shiftRightBy51(r1)
										c2 := shiftRightBy51(r2)
										c3 := shiftRightBy51(r3)
										c4 := shiftRightBy51(r4)

										rr0 := r0.lo&maskLow51Bits + c4*19
										rr1 := r1.lo&maskLow51Bits + c0
										rr2 := r2.lo&maskLow51Bits + c1
										rr3 := r3.lo&maskLow51Bits + c2
										rr4 := r4.lo&maskLow51Bits + c3

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:161
	*v = Element{rr0, rr1, rr2, rr3, rr4}
										v.carryPropagate()
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:162
	// _ = "end of CoverTab[2050]"
}

func feSquareGeneric(v, a *Element) {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:165
	_go_fuzz_dep_.CoverTab[2051]++
										l0 := a.l0
										l1 := a.l1
										l2 := a.l2
										l3 := a.l3
										l4 := a.l4

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:197
	l0_2 := l0 * 2
										l1_2 := l1 * 2

										l1_38 := l1 * 38
										l2_38 := l2 * 38
										l3_38 := l3 * 38

										l3_19 := l3 * 19
										l4_19 := l4 * 19

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:208
	r0 := mul64(l0, l0)
										r0 = addMul64(r0, l1_38, l4)
										r0 = addMul64(r0, l2_38, l3)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:213
	r1 := mul64(l0_2, l1)
										r1 = addMul64(r1, l2_38, l4)
										r1 = addMul64(r1, l3_19, l3)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:218
	r2 := mul64(l0_2, l2)
										r2 = addMul64(r2, l1, l1)
										r2 = addMul64(r2, l3_38, l4)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:223
	r3 := mul64(l0_2, l3)
										r3 = addMul64(r3, l1_2, l2)
										r3 = addMul64(r3, l4_19, l4)

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:228
	r4 := mul64(l0_2, l4)
										r4 = addMul64(r4, l1_2, l3)
										r4 = addMul64(r4, l2, l2)

										c0 := shiftRightBy51(r0)
										c1 := shiftRightBy51(r1)
										c2 := shiftRightBy51(r2)
										c3 := shiftRightBy51(r3)
										c4 := shiftRightBy51(r4)

										rr0 := r0.lo&maskLow51Bits + c4*19
										rr1 := r1.lo&maskLow51Bits + c0
										rr2 := r2.lo&maskLow51Bits + c1
										rr3 := r3.lo&maskLow51Bits + c2
										rr4 := r4.lo&maskLow51Bits + c3

										*v = Element{rr0, rr1, rr2, rr3, rr4}
										v.carryPropagate()
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:245
	// _ = "end of CoverTab[2051]"
}

// carryPropagate brings the limbs below 52 bits by applying the reduction
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:248
// identity (a * 2²⁵⁵ + b = a * 19 + b) to the l4 carry.
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:250
func (v *Element) carryPropagateGeneric() *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:250
	_go_fuzz_dep_.CoverTab[2052]++
										c0 := v.l0 >> 51
										c1 := v.l1 >> 51
										c2 := v.l2 >> 51
										c3 := v.l3 >> 51
										c4 := v.l4 >> 51

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:259
	v.l0 = v.l0&maskLow51Bits + c4*19
										v.l1 = v.l1&maskLow51Bits + c0
										v.l2 = v.l2&maskLow51Bits + c1
										v.l3 = v.l3&maskLow51Bits + c2
										v.l4 = v.l4&maskLow51Bits + c3

										return v
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:265
	// _ = "end of CoverTab[2052]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:266
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_generic.go:266
var _ = _go_fuzz_dep_.CoverTab
