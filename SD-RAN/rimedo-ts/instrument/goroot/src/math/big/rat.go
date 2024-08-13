// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements multi-precision rational numbers.

//line /usr/local/go/src/math/big/rat.go:7
package big

//line /usr/local/go/src/math/big/rat.go:7
import (
//line /usr/local/go/src/math/big/rat.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/rat.go:7
)
//line /usr/local/go/src/math/big/rat.go:7
import (
//line /usr/local/go/src/math/big/rat.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/rat.go:7
)

import (
	"fmt"
	"math"
)

// A Rat represents a quotient a/b of arbitrary precision.
//line /usr/local/go/src/math/big/rat.go:14
// The zero value for a Rat represents the value 0.
//line /usr/local/go/src/math/big/rat.go:14
//
//line /usr/local/go/src/math/big/rat.go:14
// Operations always take pointer arguments (*Rat) rather
//line /usr/local/go/src/math/big/rat.go:14
// than Rat values, and each unique Rat value requires
//line /usr/local/go/src/math/big/rat.go:14
// its own unique *Rat pointer. To "copy" a Rat value,
//line /usr/local/go/src/math/big/rat.go:14
// an existing (or newly allocated) Rat must be set to
//line /usr/local/go/src/math/big/rat.go:14
// a new value using the Rat.Set method; shallow copies
//line /usr/local/go/src/math/big/rat.go:14
// of Rats are not supported and may lead to errors.
//line /usr/local/go/src/math/big/rat.go:23
type Rat struct {
	// To make zero values for Rat work w/o initialization,
	// a zero value of b (len(b) == 0) acts like b == 1. At
	// the earliest opportunity (when an assignment to the Rat
	// is made), such uninitialized denominators are set to 1.
	// a.neg determines the sign of the Rat, b.neg is ignored.
	a, b Int
}

// NewRat creates a new Rat with numerator a and denominator b.
func NewRat(a, b int64) *Rat {
//line /usr/local/go/src/math/big/rat.go:33
	_go_fuzz_dep_.CoverTab[6539]++
						return new(Rat).SetFrac64(a, b)
//line /usr/local/go/src/math/big/rat.go:34
	// _ = "end of CoverTab[6539]"
}

// SetFloat64 sets z to exactly f and returns z.
//line /usr/local/go/src/math/big/rat.go:37
// If f is not finite, SetFloat returns nil.
//line /usr/local/go/src/math/big/rat.go:39
func (z *Rat) SetFloat64(f float64) *Rat {
//line /usr/local/go/src/math/big/rat.go:39
	_go_fuzz_dep_.CoverTab[6540]++
						const expMask = 1<<11 - 1
						bits := math.Float64bits(f)
						mantissa := bits & (1<<52 - 1)
						exp := int((bits >> 52) & expMask)
						switch exp {
	case expMask:
//line /usr/local/go/src/math/big/rat.go:45
		_go_fuzz_dep_.CoverTab[6544]++
							return nil
//line /usr/local/go/src/math/big/rat.go:46
		// _ = "end of CoverTab[6544]"
	case 0:
//line /usr/local/go/src/math/big/rat.go:47
		_go_fuzz_dep_.CoverTab[6545]++
							exp -= 1022
//line /usr/local/go/src/math/big/rat.go:48
		// _ = "end of CoverTab[6545]"
	default:
//line /usr/local/go/src/math/big/rat.go:49
		_go_fuzz_dep_.CoverTab[6546]++
							mantissa |= 1 << 52
							exp -= 1023
//line /usr/local/go/src/math/big/rat.go:51
		// _ = "end of CoverTab[6546]"
	}
//line /usr/local/go/src/math/big/rat.go:52
	// _ = "end of CoverTab[6540]"
//line /usr/local/go/src/math/big/rat.go:52
	_go_fuzz_dep_.CoverTab[6541]++

						shift := 52 - exp

//line /usr/local/go/src/math/big/rat.go:57
	for mantissa&1 == 0 && func() bool {
//line /usr/local/go/src/math/big/rat.go:57
		_go_fuzz_dep_.CoverTab[6547]++
//line /usr/local/go/src/math/big/rat.go:57
		return shift > 0
//line /usr/local/go/src/math/big/rat.go:57
		// _ = "end of CoverTab[6547]"
//line /usr/local/go/src/math/big/rat.go:57
	}() {
//line /usr/local/go/src/math/big/rat.go:57
		_go_fuzz_dep_.CoverTab[6548]++
							mantissa >>= 1
							shift--
//line /usr/local/go/src/math/big/rat.go:59
		// _ = "end of CoverTab[6548]"
	}
//line /usr/local/go/src/math/big/rat.go:60
	// _ = "end of CoverTab[6541]"
//line /usr/local/go/src/math/big/rat.go:60
	_go_fuzz_dep_.CoverTab[6542]++

						z.a.SetUint64(mantissa)
						z.a.neg = f < 0
						z.b.Set(intOne)
						if shift > 0 {
//line /usr/local/go/src/math/big/rat.go:65
		_go_fuzz_dep_.CoverTab[6549]++
							z.b.Lsh(&z.b, uint(shift))
//line /usr/local/go/src/math/big/rat.go:66
		// _ = "end of CoverTab[6549]"
	} else {
//line /usr/local/go/src/math/big/rat.go:67
		_go_fuzz_dep_.CoverTab[6550]++
							z.a.Lsh(&z.a, uint(-shift))
//line /usr/local/go/src/math/big/rat.go:68
		// _ = "end of CoverTab[6550]"
	}
//line /usr/local/go/src/math/big/rat.go:69
	// _ = "end of CoverTab[6542]"
//line /usr/local/go/src/math/big/rat.go:69
	_go_fuzz_dep_.CoverTab[6543]++
						return z.norm()
//line /usr/local/go/src/math/big/rat.go:70
	// _ = "end of CoverTab[6543]"
}

// quotToFloat32 returns the non-negative float32 value
//line /usr/local/go/src/math/big/rat.go:73
// nearest to the quotient a/b, using round-to-even in
//line /usr/local/go/src/math/big/rat.go:73
// halfway cases. It does not mutate its arguments.
//line /usr/local/go/src/math/big/rat.go:73
// Preconditions: b is non-zero; a and b have no common factors.
//line /usr/local/go/src/math/big/rat.go:77
func quotToFloat32(a, b nat) (f float32, exact bool) {
//line /usr/local/go/src/math/big/rat.go:77
	_go_fuzz_dep_.CoverTab[6551]++
						const (
		// float size in bits
		Fsize	= 32

		// mantissa
		Msize	= 23
		Msize1	= Msize + 1	// incl. implicit 1
		Msize2	= Msize1 + 1

		// exponent
		Esize	= Fsize - Msize1
		Ebias	= 1<<(Esize-1) - 1
		Emin	= 1 - Ebias
		Emax	= Ebias
	)

//line /usr/local/go/src/math/big/rat.go:95
	alen := a.bitLen()
	if alen == 0 {
//line /usr/local/go/src/math/big/rat.go:96
		_go_fuzz_dep_.CoverTab[6560]++
							return 0, true
//line /usr/local/go/src/math/big/rat.go:97
		// _ = "end of CoverTab[6560]"
	} else {
//line /usr/local/go/src/math/big/rat.go:98
		_go_fuzz_dep_.CoverTab[6561]++
//line /usr/local/go/src/math/big/rat.go:98
		// _ = "end of CoverTab[6561]"
//line /usr/local/go/src/math/big/rat.go:98
	}
//line /usr/local/go/src/math/big/rat.go:98
	// _ = "end of CoverTab[6551]"
//line /usr/local/go/src/math/big/rat.go:98
	_go_fuzz_dep_.CoverTab[6552]++
						blen := b.bitLen()
						if blen == 0 {
//line /usr/local/go/src/math/big/rat.go:100
		_go_fuzz_dep_.CoverTab[6562]++
							panic("division by zero")
//line /usr/local/go/src/math/big/rat.go:101
		// _ = "end of CoverTab[6562]"
	} else {
//line /usr/local/go/src/math/big/rat.go:102
		_go_fuzz_dep_.CoverTab[6563]++
//line /usr/local/go/src/math/big/rat.go:102
		// _ = "end of CoverTab[6563]"
//line /usr/local/go/src/math/big/rat.go:102
	}
//line /usr/local/go/src/math/big/rat.go:102
	// _ = "end of CoverTab[6552]"
//line /usr/local/go/src/math/big/rat.go:102
	_go_fuzz_dep_.CoverTab[6553]++

//line /usr/local/go/src/math/big/rat.go:110
	exp := alen - blen
	var a2, b2 nat
	a2 = a2.set(a)
	b2 = b2.set(b)
	if shift := Msize2 - exp; shift > 0 {
//line /usr/local/go/src/math/big/rat.go:114
		_go_fuzz_dep_.CoverTab[6564]++
							a2 = a2.shl(a2, uint(shift))
//line /usr/local/go/src/math/big/rat.go:115
		// _ = "end of CoverTab[6564]"
	} else {
//line /usr/local/go/src/math/big/rat.go:116
		_go_fuzz_dep_.CoverTab[6565]++
//line /usr/local/go/src/math/big/rat.go:116
		if shift < 0 {
//line /usr/local/go/src/math/big/rat.go:116
			_go_fuzz_dep_.CoverTab[6566]++
								b2 = b2.shl(b2, uint(-shift))
//line /usr/local/go/src/math/big/rat.go:117
			// _ = "end of CoverTab[6566]"
		} else {
//line /usr/local/go/src/math/big/rat.go:118
			_go_fuzz_dep_.CoverTab[6567]++
//line /usr/local/go/src/math/big/rat.go:118
			// _ = "end of CoverTab[6567]"
//line /usr/local/go/src/math/big/rat.go:118
		}
//line /usr/local/go/src/math/big/rat.go:118
		// _ = "end of CoverTab[6565]"
//line /usr/local/go/src/math/big/rat.go:118
	}
//line /usr/local/go/src/math/big/rat.go:118
	// _ = "end of CoverTab[6553]"
//line /usr/local/go/src/math/big/rat.go:118
	_go_fuzz_dep_.CoverTab[6554]++

	// 2. Compute quotient and remainder (q, r).  NB: due to the
	// extra shift, the low-order bit of q is logically the
						// high-order bit of r.
						var q nat
						q, r := q.div(a2, a2, b2)
						mantissa := low32(q)
						haveRem := len(r) > 0

//line /usr/local/go/src/math/big/rat.go:130
	if mantissa>>Msize2 == 1 {
//line /usr/local/go/src/math/big/rat.go:130
		_go_fuzz_dep_.CoverTab[6568]++
							if mantissa&1 == 1 {
//line /usr/local/go/src/math/big/rat.go:131
			_go_fuzz_dep_.CoverTab[6570]++
								haveRem = true
//line /usr/local/go/src/math/big/rat.go:132
			// _ = "end of CoverTab[6570]"
		} else {
//line /usr/local/go/src/math/big/rat.go:133
			_go_fuzz_dep_.CoverTab[6571]++
//line /usr/local/go/src/math/big/rat.go:133
			// _ = "end of CoverTab[6571]"
//line /usr/local/go/src/math/big/rat.go:133
		}
//line /usr/local/go/src/math/big/rat.go:133
		// _ = "end of CoverTab[6568]"
//line /usr/local/go/src/math/big/rat.go:133
		_go_fuzz_dep_.CoverTab[6569]++
							mantissa >>= 1
							exp++
//line /usr/local/go/src/math/big/rat.go:135
		// _ = "end of CoverTab[6569]"
	} else {
//line /usr/local/go/src/math/big/rat.go:136
		_go_fuzz_dep_.CoverTab[6572]++
//line /usr/local/go/src/math/big/rat.go:136
		// _ = "end of CoverTab[6572]"
//line /usr/local/go/src/math/big/rat.go:136
	}
//line /usr/local/go/src/math/big/rat.go:136
	// _ = "end of CoverTab[6554]"
//line /usr/local/go/src/math/big/rat.go:136
	_go_fuzz_dep_.CoverTab[6555]++
						if mantissa>>Msize1 != 1 {
//line /usr/local/go/src/math/big/rat.go:137
		_go_fuzz_dep_.CoverTab[6573]++
							panic(fmt.Sprintf("expected exactly %d bits of result", Msize2))
//line /usr/local/go/src/math/big/rat.go:138
		// _ = "end of CoverTab[6573]"
	} else {
//line /usr/local/go/src/math/big/rat.go:139
		_go_fuzz_dep_.CoverTab[6574]++
//line /usr/local/go/src/math/big/rat.go:139
		// _ = "end of CoverTab[6574]"
//line /usr/local/go/src/math/big/rat.go:139
	}
//line /usr/local/go/src/math/big/rat.go:139
	// _ = "end of CoverTab[6555]"
//line /usr/local/go/src/math/big/rat.go:139
	_go_fuzz_dep_.CoverTab[6556]++

//line /usr/local/go/src/math/big/rat.go:142
	if Emin-Msize <= exp && func() bool {
//line /usr/local/go/src/math/big/rat.go:142
		_go_fuzz_dep_.CoverTab[6575]++
//line /usr/local/go/src/math/big/rat.go:142
		return exp <= Emin
//line /usr/local/go/src/math/big/rat.go:142
		// _ = "end of CoverTab[6575]"
//line /usr/local/go/src/math/big/rat.go:142
	}() {
//line /usr/local/go/src/math/big/rat.go:142
		_go_fuzz_dep_.CoverTab[6576]++

							shift := uint(Emin - (exp - 1))
							lostbits := mantissa & (1<<shift - 1)
							haveRem = haveRem || func() bool {
//line /usr/local/go/src/math/big/rat.go:146
			_go_fuzz_dep_.CoverTab[6577]++
//line /usr/local/go/src/math/big/rat.go:146
			return lostbits != 0
//line /usr/local/go/src/math/big/rat.go:146
			// _ = "end of CoverTab[6577]"
//line /usr/local/go/src/math/big/rat.go:146
		}()
							mantissa >>= shift
							exp = 2 - Ebias
//line /usr/local/go/src/math/big/rat.go:148
		// _ = "end of CoverTab[6576]"
	} else {
//line /usr/local/go/src/math/big/rat.go:149
		_go_fuzz_dep_.CoverTab[6578]++
//line /usr/local/go/src/math/big/rat.go:149
		// _ = "end of CoverTab[6578]"
//line /usr/local/go/src/math/big/rat.go:149
	}
//line /usr/local/go/src/math/big/rat.go:149
	// _ = "end of CoverTab[6556]"
//line /usr/local/go/src/math/big/rat.go:149
	_go_fuzz_dep_.CoverTab[6557]++

						exact = !haveRem
						if mantissa&1 != 0 {
//line /usr/local/go/src/math/big/rat.go:152
		_go_fuzz_dep_.CoverTab[6579]++
							exact = false
							if haveRem || func() bool {
//line /usr/local/go/src/math/big/rat.go:154
			_go_fuzz_dep_.CoverTab[6580]++
//line /usr/local/go/src/math/big/rat.go:154
			return mantissa&2 != 0
//line /usr/local/go/src/math/big/rat.go:154
			// _ = "end of CoverTab[6580]"
//line /usr/local/go/src/math/big/rat.go:154
		}() {
//line /usr/local/go/src/math/big/rat.go:154
			_go_fuzz_dep_.CoverTab[6581]++
								if mantissa++; mantissa >= 1<<Msize2 {
//line /usr/local/go/src/math/big/rat.go:155
				_go_fuzz_dep_.CoverTab[6582]++

									mantissa >>= 1
									exp++
//line /usr/local/go/src/math/big/rat.go:158
				// _ = "end of CoverTab[6582]"
			} else {
//line /usr/local/go/src/math/big/rat.go:159
				_go_fuzz_dep_.CoverTab[6583]++
//line /usr/local/go/src/math/big/rat.go:159
				// _ = "end of CoverTab[6583]"
//line /usr/local/go/src/math/big/rat.go:159
			}
//line /usr/local/go/src/math/big/rat.go:159
			// _ = "end of CoverTab[6581]"
		} else {
//line /usr/local/go/src/math/big/rat.go:160
			_go_fuzz_dep_.CoverTab[6584]++
//line /usr/local/go/src/math/big/rat.go:160
			// _ = "end of CoverTab[6584]"
//line /usr/local/go/src/math/big/rat.go:160
		}
//line /usr/local/go/src/math/big/rat.go:160
		// _ = "end of CoverTab[6579]"
	} else {
//line /usr/local/go/src/math/big/rat.go:161
		_go_fuzz_dep_.CoverTab[6585]++
//line /usr/local/go/src/math/big/rat.go:161
		// _ = "end of CoverTab[6585]"
//line /usr/local/go/src/math/big/rat.go:161
	}
//line /usr/local/go/src/math/big/rat.go:161
	// _ = "end of CoverTab[6557]"
//line /usr/local/go/src/math/big/rat.go:161
	_go_fuzz_dep_.CoverTab[6558]++
						mantissa >>= 1

						f = float32(math.Ldexp(float64(mantissa), exp-Msize1))
						if math.IsInf(float64(f), 0) {
//line /usr/local/go/src/math/big/rat.go:165
		_go_fuzz_dep_.CoverTab[6586]++
							exact = false
//line /usr/local/go/src/math/big/rat.go:166
		// _ = "end of CoverTab[6586]"
	} else {
//line /usr/local/go/src/math/big/rat.go:167
		_go_fuzz_dep_.CoverTab[6587]++
//line /usr/local/go/src/math/big/rat.go:167
		// _ = "end of CoverTab[6587]"
//line /usr/local/go/src/math/big/rat.go:167
	}
//line /usr/local/go/src/math/big/rat.go:167
	// _ = "end of CoverTab[6558]"
//line /usr/local/go/src/math/big/rat.go:167
	_go_fuzz_dep_.CoverTab[6559]++
						return
//line /usr/local/go/src/math/big/rat.go:168
	// _ = "end of CoverTab[6559]"
}

// quotToFloat64 returns the non-negative float64 value
//line /usr/local/go/src/math/big/rat.go:171
// nearest to the quotient a/b, using round-to-even in
//line /usr/local/go/src/math/big/rat.go:171
// halfway cases. It does not mutate its arguments.
//line /usr/local/go/src/math/big/rat.go:171
// Preconditions: b is non-zero; a and b have no common factors.
//line /usr/local/go/src/math/big/rat.go:175
func quotToFloat64(a, b nat) (f float64, exact bool) {
//line /usr/local/go/src/math/big/rat.go:175
	_go_fuzz_dep_.CoverTab[6588]++
						const (
		// float size in bits
		Fsize	= 64

		// mantissa
		Msize	= 52
		Msize1	= Msize + 1	// incl. implicit 1
		Msize2	= Msize1 + 1

		// exponent
		Esize	= Fsize - Msize1
		Ebias	= 1<<(Esize-1) - 1
		Emin	= 1 - Ebias
		Emax	= Ebias
	)

//line /usr/local/go/src/math/big/rat.go:193
	alen := a.bitLen()
	if alen == 0 {
//line /usr/local/go/src/math/big/rat.go:194
		_go_fuzz_dep_.CoverTab[6597]++
							return 0, true
//line /usr/local/go/src/math/big/rat.go:195
		// _ = "end of CoverTab[6597]"
	} else {
//line /usr/local/go/src/math/big/rat.go:196
		_go_fuzz_dep_.CoverTab[6598]++
//line /usr/local/go/src/math/big/rat.go:196
		// _ = "end of CoverTab[6598]"
//line /usr/local/go/src/math/big/rat.go:196
	}
//line /usr/local/go/src/math/big/rat.go:196
	// _ = "end of CoverTab[6588]"
//line /usr/local/go/src/math/big/rat.go:196
	_go_fuzz_dep_.CoverTab[6589]++
						blen := b.bitLen()
						if blen == 0 {
//line /usr/local/go/src/math/big/rat.go:198
		_go_fuzz_dep_.CoverTab[6599]++
							panic("division by zero")
//line /usr/local/go/src/math/big/rat.go:199
		// _ = "end of CoverTab[6599]"
	} else {
//line /usr/local/go/src/math/big/rat.go:200
		_go_fuzz_dep_.CoverTab[6600]++
//line /usr/local/go/src/math/big/rat.go:200
		// _ = "end of CoverTab[6600]"
//line /usr/local/go/src/math/big/rat.go:200
	}
//line /usr/local/go/src/math/big/rat.go:200
	// _ = "end of CoverTab[6589]"
//line /usr/local/go/src/math/big/rat.go:200
	_go_fuzz_dep_.CoverTab[6590]++

//line /usr/local/go/src/math/big/rat.go:208
	exp := alen - blen
	var a2, b2 nat
	a2 = a2.set(a)
	b2 = b2.set(b)
	if shift := Msize2 - exp; shift > 0 {
//line /usr/local/go/src/math/big/rat.go:212
		_go_fuzz_dep_.CoverTab[6601]++
							a2 = a2.shl(a2, uint(shift))
//line /usr/local/go/src/math/big/rat.go:213
		// _ = "end of CoverTab[6601]"
	} else {
//line /usr/local/go/src/math/big/rat.go:214
		_go_fuzz_dep_.CoverTab[6602]++
//line /usr/local/go/src/math/big/rat.go:214
		if shift < 0 {
//line /usr/local/go/src/math/big/rat.go:214
			_go_fuzz_dep_.CoverTab[6603]++
								b2 = b2.shl(b2, uint(-shift))
//line /usr/local/go/src/math/big/rat.go:215
			// _ = "end of CoverTab[6603]"
		} else {
//line /usr/local/go/src/math/big/rat.go:216
			_go_fuzz_dep_.CoverTab[6604]++
//line /usr/local/go/src/math/big/rat.go:216
			// _ = "end of CoverTab[6604]"
//line /usr/local/go/src/math/big/rat.go:216
		}
//line /usr/local/go/src/math/big/rat.go:216
		// _ = "end of CoverTab[6602]"
//line /usr/local/go/src/math/big/rat.go:216
	}
//line /usr/local/go/src/math/big/rat.go:216
	// _ = "end of CoverTab[6590]"
//line /usr/local/go/src/math/big/rat.go:216
	_go_fuzz_dep_.CoverTab[6591]++

	// 2. Compute quotient and remainder (q, r).  NB: due to the
	// extra shift, the low-order bit of q is logically the
						// high-order bit of r.
						var q nat
						q, r := q.div(a2, a2, b2)
						mantissa := low64(q)
						haveRem := len(r) > 0

//line /usr/local/go/src/math/big/rat.go:228
	if mantissa>>Msize2 == 1 {
//line /usr/local/go/src/math/big/rat.go:228
		_go_fuzz_dep_.CoverTab[6605]++
							if mantissa&1 == 1 {
//line /usr/local/go/src/math/big/rat.go:229
			_go_fuzz_dep_.CoverTab[6607]++
								haveRem = true
//line /usr/local/go/src/math/big/rat.go:230
			// _ = "end of CoverTab[6607]"
		} else {
//line /usr/local/go/src/math/big/rat.go:231
			_go_fuzz_dep_.CoverTab[6608]++
//line /usr/local/go/src/math/big/rat.go:231
			// _ = "end of CoverTab[6608]"
//line /usr/local/go/src/math/big/rat.go:231
		}
//line /usr/local/go/src/math/big/rat.go:231
		// _ = "end of CoverTab[6605]"
//line /usr/local/go/src/math/big/rat.go:231
		_go_fuzz_dep_.CoverTab[6606]++
							mantissa >>= 1
							exp++
//line /usr/local/go/src/math/big/rat.go:233
		// _ = "end of CoverTab[6606]"
	} else {
//line /usr/local/go/src/math/big/rat.go:234
		_go_fuzz_dep_.CoverTab[6609]++
//line /usr/local/go/src/math/big/rat.go:234
		// _ = "end of CoverTab[6609]"
//line /usr/local/go/src/math/big/rat.go:234
	}
//line /usr/local/go/src/math/big/rat.go:234
	// _ = "end of CoverTab[6591]"
//line /usr/local/go/src/math/big/rat.go:234
	_go_fuzz_dep_.CoverTab[6592]++
						if mantissa>>Msize1 != 1 {
//line /usr/local/go/src/math/big/rat.go:235
		_go_fuzz_dep_.CoverTab[6610]++
							panic(fmt.Sprintf("expected exactly %d bits of result", Msize2))
//line /usr/local/go/src/math/big/rat.go:236
		// _ = "end of CoverTab[6610]"
	} else {
//line /usr/local/go/src/math/big/rat.go:237
		_go_fuzz_dep_.CoverTab[6611]++
//line /usr/local/go/src/math/big/rat.go:237
		// _ = "end of CoverTab[6611]"
//line /usr/local/go/src/math/big/rat.go:237
	}
//line /usr/local/go/src/math/big/rat.go:237
	// _ = "end of CoverTab[6592]"
//line /usr/local/go/src/math/big/rat.go:237
	_go_fuzz_dep_.CoverTab[6593]++

//line /usr/local/go/src/math/big/rat.go:240
	if Emin-Msize <= exp && func() bool {
//line /usr/local/go/src/math/big/rat.go:240
		_go_fuzz_dep_.CoverTab[6612]++
//line /usr/local/go/src/math/big/rat.go:240
		return exp <= Emin
//line /usr/local/go/src/math/big/rat.go:240
		// _ = "end of CoverTab[6612]"
//line /usr/local/go/src/math/big/rat.go:240
	}() {
//line /usr/local/go/src/math/big/rat.go:240
		_go_fuzz_dep_.CoverTab[6613]++

							shift := uint(Emin - (exp - 1))
							lostbits := mantissa & (1<<shift - 1)
							haveRem = haveRem || func() bool {
//line /usr/local/go/src/math/big/rat.go:244
			_go_fuzz_dep_.CoverTab[6614]++
//line /usr/local/go/src/math/big/rat.go:244
			return lostbits != 0
//line /usr/local/go/src/math/big/rat.go:244
			// _ = "end of CoverTab[6614]"
//line /usr/local/go/src/math/big/rat.go:244
		}()
							mantissa >>= shift
							exp = 2 - Ebias
//line /usr/local/go/src/math/big/rat.go:246
		// _ = "end of CoverTab[6613]"
	} else {
//line /usr/local/go/src/math/big/rat.go:247
		_go_fuzz_dep_.CoverTab[6615]++
//line /usr/local/go/src/math/big/rat.go:247
		// _ = "end of CoverTab[6615]"
//line /usr/local/go/src/math/big/rat.go:247
	}
//line /usr/local/go/src/math/big/rat.go:247
	// _ = "end of CoverTab[6593]"
//line /usr/local/go/src/math/big/rat.go:247
	_go_fuzz_dep_.CoverTab[6594]++

						exact = !haveRem
						if mantissa&1 != 0 {
//line /usr/local/go/src/math/big/rat.go:250
		_go_fuzz_dep_.CoverTab[6616]++
							exact = false
							if haveRem || func() bool {
//line /usr/local/go/src/math/big/rat.go:252
			_go_fuzz_dep_.CoverTab[6617]++
//line /usr/local/go/src/math/big/rat.go:252
			return mantissa&2 != 0
//line /usr/local/go/src/math/big/rat.go:252
			// _ = "end of CoverTab[6617]"
//line /usr/local/go/src/math/big/rat.go:252
		}() {
//line /usr/local/go/src/math/big/rat.go:252
			_go_fuzz_dep_.CoverTab[6618]++
								if mantissa++; mantissa >= 1<<Msize2 {
//line /usr/local/go/src/math/big/rat.go:253
				_go_fuzz_dep_.CoverTab[6619]++

									mantissa >>= 1
									exp++
//line /usr/local/go/src/math/big/rat.go:256
				// _ = "end of CoverTab[6619]"
			} else {
//line /usr/local/go/src/math/big/rat.go:257
				_go_fuzz_dep_.CoverTab[6620]++
//line /usr/local/go/src/math/big/rat.go:257
				// _ = "end of CoverTab[6620]"
//line /usr/local/go/src/math/big/rat.go:257
			}
//line /usr/local/go/src/math/big/rat.go:257
			// _ = "end of CoverTab[6618]"
		} else {
//line /usr/local/go/src/math/big/rat.go:258
			_go_fuzz_dep_.CoverTab[6621]++
//line /usr/local/go/src/math/big/rat.go:258
			// _ = "end of CoverTab[6621]"
//line /usr/local/go/src/math/big/rat.go:258
		}
//line /usr/local/go/src/math/big/rat.go:258
		// _ = "end of CoverTab[6616]"
	} else {
//line /usr/local/go/src/math/big/rat.go:259
		_go_fuzz_dep_.CoverTab[6622]++
//line /usr/local/go/src/math/big/rat.go:259
		// _ = "end of CoverTab[6622]"
//line /usr/local/go/src/math/big/rat.go:259
	}
//line /usr/local/go/src/math/big/rat.go:259
	// _ = "end of CoverTab[6594]"
//line /usr/local/go/src/math/big/rat.go:259
	_go_fuzz_dep_.CoverTab[6595]++
						mantissa >>= 1

						f = math.Ldexp(float64(mantissa), exp-Msize1)
						if math.IsInf(f, 0) {
//line /usr/local/go/src/math/big/rat.go:263
		_go_fuzz_dep_.CoverTab[6623]++
							exact = false
//line /usr/local/go/src/math/big/rat.go:264
		// _ = "end of CoverTab[6623]"
	} else {
//line /usr/local/go/src/math/big/rat.go:265
		_go_fuzz_dep_.CoverTab[6624]++
//line /usr/local/go/src/math/big/rat.go:265
		// _ = "end of CoverTab[6624]"
//line /usr/local/go/src/math/big/rat.go:265
	}
//line /usr/local/go/src/math/big/rat.go:265
	// _ = "end of CoverTab[6595]"
//line /usr/local/go/src/math/big/rat.go:265
	_go_fuzz_dep_.CoverTab[6596]++
						return
//line /usr/local/go/src/math/big/rat.go:266
	// _ = "end of CoverTab[6596]"
}

// Float32 returns the nearest float32 value for x and a bool indicating
//line /usr/local/go/src/math/big/rat.go:269
// whether f represents x exactly. If the magnitude of x is too large to
//line /usr/local/go/src/math/big/rat.go:269
// be represented by a float32, f is an infinity and exact is false.
//line /usr/local/go/src/math/big/rat.go:269
// The sign of f always matches the sign of x, even if f == 0.
//line /usr/local/go/src/math/big/rat.go:273
func (x *Rat) Float32() (f float32, exact bool) {
//line /usr/local/go/src/math/big/rat.go:273
	_go_fuzz_dep_.CoverTab[6625]++
						b := x.b.abs
						if len(b) == 0 {
//line /usr/local/go/src/math/big/rat.go:275
		_go_fuzz_dep_.CoverTab[6628]++
							b = natOne
//line /usr/local/go/src/math/big/rat.go:276
		// _ = "end of CoverTab[6628]"
	} else {
//line /usr/local/go/src/math/big/rat.go:277
		_go_fuzz_dep_.CoverTab[6629]++
//line /usr/local/go/src/math/big/rat.go:277
		// _ = "end of CoverTab[6629]"
//line /usr/local/go/src/math/big/rat.go:277
	}
//line /usr/local/go/src/math/big/rat.go:277
	// _ = "end of CoverTab[6625]"
//line /usr/local/go/src/math/big/rat.go:277
	_go_fuzz_dep_.CoverTab[6626]++
						f, exact = quotToFloat32(x.a.abs, b)
						if x.a.neg {
//line /usr/local/go/src/math/big/rat.go:279
		_go_fuzz_dep_.CoverTab[6630]++
							f = -f
//line /usr/local/go/src/math/big/rat.go:280
		// _ = "end of CoverTab[6630]"
	} else {
//line /usr/local/go/src/math/big/rat.go:281
		_go_fuzz_dep_.CoverTab[6631]++
//line /usr/local/go/src/math/big/rat.go:281
		// _ = "end of CoverTab[6631]"
//line /usr/local/go/src/math/big/rat.go:281
	}
//line /usr/local/go/src/math/big/rat.go:281
	// _ = "end of CoverTab[6626]"
//line /usr/local/go/src/math/big/rat.go:281
	_go_fuzz_dep_.CoverTab[6627]++
						return
//line /usr/local/go/src/math/big/rat.go:282
	// _ = "end of CoverTab[6627]"
}

// Float64 returns the nearest float64 value for x and a bool indicating
//line /usr/local/go/src/math/big/rat.go:285
// whether f represents x exactly. If the magnitude of x is too large to
//line /usr/local/go/src/math/big/rat.go:285
// be represented by a float64, f is an infinity and exact is false.
//line /usr/local/go/src/math/big/rat.go:285
// The sign of f always matches the sign of x, even if f == 0.
//line /usr/local/go/src/math/big/rat.go:289
func (x *Rat) Float64() (f float64, exact bool) {
//line /usr/local/go/src/math/big/rat.go:289
	_go_fuzz_dep_.CoverTab[6632]++
						b := x.b.abs
						if len(b) == 0 {
//line /usr/local/go/src/math/big/rat.go:291
		_go_fuzz_dep_.CoverTab[6635]++
							b = natOne
//line /usr/local/go/src/math/big/rat.go:292
		// _ = "end of CoverTab[6635]"
	} else {
//line /usr/local/go/src/math/big/rat.go:293
		_go_fuzz_dep_.CoverTab[6636]++
//line /usr/local/go/src/math/big/rat.go:293
		// _ = "end of CoverTab[6636]"
//line /usr/local/go/src/math/big/rat.go:293
	}
//line /usr/local/go/src/math/big/rat.go:293
	// _ = "end of CoverTab[6632]"
//line /usr/local/go/src/math/big/rat.go:293
	_go_fuzz_dep_.CoverTab[6633]++
						f, exact = quotToFloat64(x.a.abs, b)
						if x.a.neg {
//line /usr/local/go/src/math/big/rat.go:295
		_go_fuzz_dep_.CoverTab[6637]++
							f = -f
//line /usr/local/go/src/math/big/rat.go:296
		// _ = "end of CoverTab[6637]"
	} else {
//line /usr/local/go/src/math/big/rat.go:297
		_go_fuzz_dep_.CoverTab[6638]++
//line /usr/local/go/src/math/big/rat.go:297
		// _ = "end of CoverTab[6638]"
//line /usr/local/go/src/math/big/rat.go:297
	}
//line /usr/local/go/src/math/big/rat.go:297
	// _ = "end of CoverTab[6633]"
//line /usr/local/go/src/math/big/rat.go:297
	_go_fuzz_dep_.CoverTab[6634]++
						return
//line /usr/local/go/src/math/big/rat.go:298
	// _ = "end of CoverTab[6634]"
}

// SetFrac sets z to a/b and returns z.
//line /usr/local/go/src/math/big/rat.go:301
// If b == 0, SetFrac panics.
//line /usr/local/go/src/math/big/rat.go:303
func (z *Rat) SetFrac(a, b *Int) *Rat {
//line /usr/local/go/src/math/big/rat.go:303
	_go_fuzz_dep_.CoverTab[6639]++
						z.a.neg = a.neg != b.neg
						babs := b.abs
						if len(babs) == 0 {
//line /usr/local/go/src/math/big/rat.go:306
		_go_fuzz_dep_.CoverTab[6642]++
							panic("division by zero")
//line /usr/local/go/src/math/big/rat.go:307
		// _ = "end of CoverTab[6642]"
	} else {
//line /usr/local/go/src/math/big/rat.go:308
		_go_fuzz_dep_.CoverTab[6643]++
//line /usr/local/go/src/math/big/rat.go:308
		// _ = "end of CoverTab[6643]"
//line /usr/local/go/src/math/big/rat.go:308
	}
//line /usr/local/go/src/math/big/rat.go:308
	// _ = "end of CoverTab[6639]"
//line /usr/local/go/src/math/big/rat.go:308
	_go_fuzz_dep_.CoverTab[6640]++
						if &z.a == b || func() bool {
//line /usr/local/go/src/math/big/rat.go:309
		_go_fuzz_dep_.CoverTab[6644]++
//line /usr/local/go/src/math/big/rat.go:309
		return alias(z.a.abs, babs)
//line /usr/local/go/src/math/big/rat.go:309
		// _ = "end of CoverTab[6644]"
//line /usr/local/go/src/math/big/rat.go:309
	}() {
//line /usr/local/go/src/math/big/rat.go:309
		_go_fuzz_dep_.CoverTab[6645]++
							babs = nat(nil).set(babs)
//line /usr/local/go/src/math/big/rat.go:310
		// _ = "end of CoverTab[6645]"
	} else {
//line /usr/local/go/src/math/big/rat.go:311
		_go_fuzz_dep_.CoverTab[6646]++
//line /usr/local/go/src/math/big/rat.go:311
		// _ = "end of CoverTab[6646]"
//line /usr/local/go/src/math/big/rat.go:311
	}
//line /usr/local/go/src/math/big/rat.go:311
	// _ = "end of CoverTab[6640]"
//line /usr/local/go/src/math/big/rat.go:311
	_go_fuzz_dep_.CoverTab[6641]++
						z.a.abs = z.a.abs.set(a.abs)
						z.b.abs = z.b.abs.set(babs)
						return z.norm()
//line /usr/local/go/src/math/big/rat.go:314
	// _ = "end of CoverTab[6641]"
}

// SetFrac64 sets z to a/b and returns z.
//line /usr/local/go/src/math/big/rat.go:317
// If b == 0, SetFrac64 panics.
//line /usr/local/go/src/math/big/rat.go:319
func (z *Rat) SetFrac64(a, b int64) *Rat {
//line /usr/local/go/src/math/big/rat.go:319
	_go_fuzz_dep_.CoverTab[6647]++
						if b == 0 {
//line /usr/local/go/src/math/big/rat.go:320
		_go_fuzz_dep_.CoverTab[6650]++
							panic("division by zero")
//line /usr/local/go/src/math/big/rat.go:321
		// _ = "end of CoverTab[6650]"
	} else {
//line /usr/local/go/src/math/big/rat.go:322
		_go_fuzz_dep_.CoverTab[6651]++
//line /usr/local/go/src/math/big/rat.go:322
		// _ = "end of CoverTab[6651]"
//line /usr/local/go/src/math/big/rat.go:322
	}
//line /usr/local/go/src/math/big/rat.go:322
	// _ = "end of CoverTab[6647]"
//line /usr/local/go/src/math/big/rat.go:322
	_go_fuzz_dep_.CoverTab[6648]++
						z.a.SetInt64(a)
						if b < 0 {
//line /usr/local/go/src/math/big/rat.go:324
		_go_fuzz_dep_.CoverTab[6652]++
							b = -b
							z.a.neg = !z.a.neg
//line /usr/local/go/src/math/big/rat.go:326
		// _ = "end of CoverTab[6652]"
	} else {
//line /usr/local/go/src/math/big/rat.go:327
		_go_fuzz_dep_.CoverTab[6653]++
//line /usr/local/go/src/math/big/rat.go:327
		// _ = "end of CoverTab[6653]"
//line /usr/local/go/src/math/big/rat.go:327
	}
//line /usr/local/go/src/math/big/rat.go:327
	// _ = "end of CoverTab[6648]"
//line /usr/local/go/src/math/big/rat.go:327
	_go_fuzz_dep_.CoverTab[6649]++
						z.b.abs = z.b.abs.setUint64(uint64(b))
						return z.norm()
//line /usr/local/go/src/math/big/rat.go:329
	// _ = "end of CoverTab[6649]"
}

// SetInt sets z to x (by making a copy of x) and returns z.
func (z *Rat) SetInt(x *Int) *Rat {
//line /usr/local/go/src/math/big/rat.go:333
	_go_fuzz_dep_.CoverTab[6654]++
						z.a.Set(x)
						z.b.abs = z.b.abs.setWord(1)
						return z
//line /usr/local/go/src/math/big/rat.go:336
	// _ = "end of CoverTab[6654]"
}

// SetInt64 sets z to x and returns z.
func (z *Rat) SetInt64(x int64) *Rat {
//line /usr/local/go/src/math/big/rat.go:340
	_go_fuzz_dep_.CoverTab[6655]++
						z.a.SetInt64(x)
						z.b.abs = z.b.abs.setWord(1)
						return z
//line /usr/local/go/src/math/big/rat.go:343
	// _ = "end of CoverTab[6655]"
}

// SetUint64 sets z to x and returns z.
func (z *Rat) SetUint64(x uint64) *Rat {
//line /usr/local/go/src/math/big/rat.go:347
	_go_fuzz_dep_.CoverTab[6656]++
						z.a.SetUint64(x)
						z.b.abs = z.b.abs.setWord(1)
						return z
//line /usr/local/go/src/math/big/rat.go:350
	// _ = "end of CoverTab[6656]"
}

// Set sets z to x (by making a copy of x) and returns z.
func (z *Rat) Set(x *Rat) *Rat {
//line /usr/local/go/src/math/big/rat.go:354
	_go_fuzz_dep_.CoverTab[6657]++
						if z != x {
//line /usr/local/go/src/math/big/rat.go:355
		_go_fuzz_dep_.CoverTab[6660]++
							z.a.Set(&x.a)
							z.b.Set(&x.b)
//line /usr/local/go/src/math/big/rat.go:357
		// _ = "end of CoverTab[6660]"
	} else {
//line /usr/local/go/src/math/big/rat.go:358
		_go_fuzz_dep_.CoverTab[6661]++
//line /usr/local/go/src/math/big/rat.go:358
		// _ = "end of CoverTab[6661]"
//line /usr/local/go/src/math/big/rat.go:358
	}
//line /usr/local/go/src/math/big/rat.go:358
	// _ = "end of CoverTab[6657]"
//line /usr/local/go/src/math/big/rat.go:358
	_go_fuzz_dep_.CoverTab[6658]++
						if len(z.b.abs) == 0 {
//line /usr/local/go/src/math/big/rat.go:359
		_go_fuzz_dep_.CoverTab[6662]++
							z.b.abs = z.b.abs.setWord(1)
//line /usr/local/go/src/math/big/rat.go:360
		// _ = "end of CoverTab[6662]"
	} else {
//line /usr/local/go/src/math/big/rat.go:361
		_go_fuzz_dep_.CoverTab[6663]++
//line /usr/local/go/src/math/big/rat.go:361
		// _ = "end of CoverTab[6663]"
//line /usr/local/go/src/math/big/rat.go:361
	}
//line /usr/local/go/src/math/big/rat.go:361
	// _ = "end of CoverTab[6658]"
//line /usr/local/go/src/math/big/rat.go:361
	_go_fuzz_dep_.CoverTab[6659]++
						return z
//line /usr/local/go/src/math/big/rat.go:362
	// _ = "end of CoverTab[6659]"
}

// Abs sets z to |x| (the absolute value of x) and returns z.
func (z *Rat) Abs(x *Rat) *Rat {
//line /usr/local/go/src/math/big/rat.go:366
	_go_fuzz_dep_.CoverTab[6664]++
						z.Set(x)
						z.a.neg = false
						return z
//line /usr/local/go/src/math/big/rat.go:369
	// _ = "end of CoverTab[6664]"
}

// Neg sets z to -x and returns z.
func (z *Rat) Neg(x *Rat) *Rat {
//line /usr/local/go/src/math/big/rat.go:373
	_go_fuzz_dep_.CoverTab[6665]++
						z.Set(x)
						z.a.neg = len(z.a.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/rat.go:375
		_go_fuzz_dep_.CoverTab[6666]++
//line /usr/local/go/src/math/big/rat.go:375
		return !z.a.neg
//line /usr/local/go/src/math/big/rat.go:375
		// _ = "end of CoverTab[6666]"
//line /usr/local/go/src/math/big/rat.go:375
	}()
						return z
//line /usr/local/go/src/math/big/rat.go:376
	// _ = "end of CoverTab[6665]"
}

// Inv sets z to 1/x and returns z.
//line /usr/local/go/src/math/big/rat.go:379
// If x == 0, Inv panics.
//line /usr/local/go/src/math/big/rat.go:381
func (z *Rat) Inv(x *Rat) *Rat {
//line /usr/local/go/src/math/big/rat.go:381
	_go_fuzz_dep_.CoverTab[6667]++
						if len(x.a.abs) == 0 {
//line /usr/local/go/src/math/big/rat.go:382
		_go_fuzz_dep_.CoverTab[6669]++
							panic("division by zero")
//line /usr/local/go/src/math/big/rat.go:383
		// _ = "end of CoverTab[6669]"
	} else {
//line /usr/local/go/src/math/big/rat.go:384
		_go_fuzz_dep_.CoverTab[6670]++
//line /usr/local/go/src/math/big/rat.go:384
		// _ = "end of CoverTab[6670]"
//line /usr/local/go/src/math/big/rat.go:384
	}
//line /usr/local/go/src/math/big/rat.go:384
	// _ = "end of CoverTab[6667]"
//line /usr/local/go/src/math/big/rat.go:384
	_go_fuzz_dep_.CoverTab[6668]++
						z.Set(x)
						z.a.abs, z.b.abs = z.b.abs, z.a.abs
						return z
//line /usr/local/go/src/math/big/rat.go:387
	// _ = "end of CoverTab[6668]"
}

// Sign returns:
//line /usr/local/go/src/math/big/rat.go:390
//
//line /usr/local/go/src/math/big/rat.go:390
//	-1 if x <  0
//line /usr/local/go/src/math/big/rat.go:390
//	 0 if x == 0
//line /usr/local/go/src/math/big/rat.go:390
//	+1 if x >  0
//line /usr/local/go/src/math/big/rat.go:395
func (x *Rat) Sign() int {
//line /usr/local/go/src/math/big/rat.go:395
	_go_fuzz_dep_.CoverTab[6671]++
						return x.a.Sign()
//line /usr/local/go/src/math/big/rat.go:396
	// _ = "end of CoverTab[6671]"
}

// IsInt reports whether the denominator of x is 1.
func (x *Rat) IsInt() bool {
//line /usr/local/go/src/math/big/rat.go:400
	_go_fuzz_dep_.CoverTab[6672]++
						return len(x.b.abs) == 0 || func() bool {
//line /usr/local/go/src/math/big/rat.go:401
		_go_fuzz_dep_.CoverTab[6673]++
//line /usr/local/go/src/math/big/rat.go:401
		return x.b.abs.cmp(natOne) == 0
//line /usr/local/go/src/math/big/rat.go:401
		// _ = "end of CoverTab[6673]"
//line /usr/local/go/src/math/big/rat.go:401
	}()
//line /usr/local/go/src/math/big/rat.go:401
	// _ = "end of CoverTab[6672]"
}

// Num returns the numerator of x; it may be <= 0.
//line /usr/local/go/src/math/big/rat.go:404
// The result is a reference to x's numerator; it
//line /usr/local/go/src/math/big/rat.go:404
// may change if a new value is assigned to x, and vice versa.
//line /usr/local/go/src/math/big/rat.go:404
// The sign of the numerator corresponds to the sign of x.
//line /usr/local/go/src/math/big/rat.go:408
func (x *Rat) Num() *Int {
//line /usr/local/go/src/math/big/rat.go:408
	_go_fuzz_dep_.CoverTab[6674]++
						return &x.a
//line /usr/local/go/src/math/big/rat.go:409
	// _ = "end of CoverTab[6674]"
}

// Denom returns the denominator of x; it is always > 0.
//line /usr/local/go/src/math/big/rat.go:412
// The result is a reference to x's denominator, unless
//line /usr/local/go/src/math/big/rat.go:412
// x is an uninitialized (zero value) Rat, in which case
//line /usr/local/go/src/math/big/rat.go:412
// the result is a new Int of value 1. (To initialize x,
//line /usr/local/go/src/math/big/rat.go:412
// any operation that sets x will do, including x.Set(x).)
//line /usr/local/go/src/math/big/rat.go:412
// If the result is a reference to x's denominator it
//line /usr/local/go/src/math/big/rat.go:412
// may change if a new value is assigned to x, and vice versa.
//line /usr/local/go/src/math/big/rat.go:419
func (x *Rat) Denom() *Int {
//line /usr/local/go/src/math/big/rat.go:419
	_go_fuzz_dep_.CoverTab[6675]++

						if len(x.b.abs) == 0 {
//line /usr/local/go/src/math/big/rat.go:421
		_go_fuzz_dep_.CoverTab[6677]++

//line /usr/local/go/src/math/big/rat.go:425
		return &Int{abs: nat{1}}
//line /usr/local/go/src/math/big/rat.go:425
		// _ = "end of CoverTab[6677]"
	} else {
//line /usr/local/go/src/math/big/rat.go:426
		_go_fuzz_dep_.CoverTab[6678]++
//line /usr/local/go/src/math/big/rat.go:426
		// _ = "end of CoverTab[6678]"
//line /usr/local/go/src/math/big/rat.go:426
	}
//line /usr/local/go/src/math/big/rat.go:426
	// _ = "end of CoverTab[6675]"
//line /usr/local/go/src/math/big/rat.go:426
	_go_fuzz_dep_.CoverTab[6676]++
						return &x.b
//line /usr/local/go/src/math/big/rat.go:427
	// _ = "end of CoverTab[6676]"
}

func (z *Rat) norm() *Rat {
//line /usr/local/go/src/math/big/rat.go:430
	_go_fuzz_dep_.CoverTab[6679]++
						switch {
	case len(z.a.abs) == 0:
//line /usr/local/go/src/math/big/rat.go:432
		_go_fuzz_dep_.CoverTab[6681]++

							z.a.neg = false
							fallthrough
//line /usr/local/go/src/math/big/rat.go:435
		// _ = "end of CoverTab[6681]"
	case len(z.b.abs) == 0:
//line /usr/local/go/src/math/big/rat.go:436
		_go_fuzz_dep_.CoverTab[6682]++

							z.b.abs = z.b.abs.setWord(1)
//line /usr/local/go/src/math/big/rat.go:438
		// _ = "end of CoverTab[6682]"
	default:
//line /usr/local/go/src/math/big/rat.go:439
		_go_fuzz_dep_.CoverTab[6683]++

							neg := z.a.neg
							z.a.neg = false
							z.b.neg = false
							if f := NewInt(0).lehmerGCD(nil, nil, &z.a, &z.b); f.Cmp(intOne) != 0 {
//line /usr/local/go/src/math/big/rat.go:444
			_go_fuzz_dep_.CoverTab[6685]++
								z.a.abs, _ = z.a.abs.div(nil, z.a.abs, f.abs)
								z.b.abs, _ = z.b.abs.div(nil, z.b.abs, f.abs)
//line /usr/local/go/src/math/big/rat.go:446
			// _ = "end of CoverTab[6685]"
		} else {
//line /usr/local/go/src/math/big/rat.go:447
			_go_fuzz_dep_.CoverTab[6686]++
//line /usr/local/go/src/math/big/rat.go:447
			// _ = "end of CoverTab[6686]"
//line /usr/local/go/src/math/big/rat.go:447
		}
//line /usr/local/go/src/math/big/rat.go:447
		// _ = "end of CoverTab[6683]"
//line /usr/local/go/src/math/big/rat.go:447
		_go_fuzz_dep_.CoverTab[6684]++
							z.a.neg = neg
//line /usr/local/go/src/math/big/rat.go:448
		// _ = "end of CoverTab[6684]"
	}
//line /usr/local/go/src/math/big/rat.go:449
	// _ = "end of CoverTab[6679]"
//line /usr/local/go/src/math/big/rat.go:449
	_go_fuzz_dep_.CoverTab[6680]++
						return z
//line /usr/local/go/src/math/big/rat.go:450
	// _ = "end of CoverTab[6680]"
}

// mulDenom sets z to the denominator product x*y (by taking into
//line /usr/local/go/src/math/big/rat.go:453
// account that 0 values for x or y must be interpreted as 1) and
//line /usr/local/go/src/math/big/rat.go:453
// returns z.
//line /usr/local/go/src/math/big/rat.go:456
func mulDenom(z, x, y nat) nat {
//line /usr/local/go/src/math/big/rat.go:456
	_go_fuzz_dep_.CoverTab[6687]++
						switch {
	case len(x) == 0 && func() bool {
//line /usr/local/go/src/math/big/rat.go:458
		_go_fuzz_dep_.CoverTab[6693]++
//line /usr/local/go/src/math/big/rat.go:458
		return len(y) == 0
//line /usr/local/go/src/math/big/rat.go:458
		// _ = "end of CoverTab[6693]"
//line /usr/local/go/src/math/big/rat.go:458
	}():
//line /usr/local/go/src/math/big/rat.go:458
		_go_fuzz_dep_.CoverTab[6689]++
							return z.setWord(1)
//line /usr/local/go/src/math/big/rat.go:459
		// _ = "end of CoverTab[6689]"
	case len(x) == 0:
//line /usr/local/go/src/math/big/rat.go:460
		_go_fuzz_dep_.CoverTab[6690]++
							return z.set(y)
//line /usr/local/go/src/math/big/rat.go:461
		// _ = "end of CoverTab[6690]"
	case len(y) == 0:
//line /usr/local/go/src/math/big/rat.go:462
		_go_fuzz_dep_.CoverTab[6691]++
							return z.set(x)
//line /usr/local/go/src/math/big/rat.go:463
		// _ = "end of CoverTab[6691]"
//line /usr/local/go/src/math/big/rat.go:463
	default:
//line /usr/local/go/src/math/big/rat.go:463
		_go_fuzz_dep_.CoverTab[6692]++
//line /usr/local/go/src/math/big/rat.go:463
		// _ = "end of CoverTab[6692]"
	}
//line /usr/local/go/src/math/big/rat.go:464
	// _ = "end of CoverTab[6687]"
//line /usr/local/go/src/math/big/rat.go:464
	_go_fuzz_dep_.CoverTab[6688]++
						return z.mul(x, y)
//line /usr/local/go/src/math/big/rat.go:465
	// _ = "end of CoverTab[6688]"
}

// scaleDenom sets z to the product x*f.
//line /usr/local/go/src/math/big/rat.go:468
// If f == 0 (zero value of denominator), z is set to (a copy of) x.
//line /usr/local/go/src/math/big/rat.go:470
func (z *Int) scaleDenom(x *Int, f nat) {
//line /usr/local/go/src/math/big/rat.go:470
	_go_fuzz_dep_.CoverTab[6694]++
						if len(f) == 0 {
//line /usr/local/go/src/math/big/rat.go:471
		_go_fuzz_dep_.CoverTab[6696]++
							z.Set(x)
							return
//line /usr/local/go/src/math/big/rat.go:473
		// _ = "end of CoverTab[6696]"
	} else {
//line /usr/local/go/src/math/big/rat.go:474
		_go_fuzz_dep_.CoverTab[6697]++
//line /usr/local/go/src/math/big/rat.go:474
		// _ = "end of CoverTab[6697]"
//line /usr/local/go/src/math/big/rat.go:474
	}
//line /usr/local/go/src/math/big/rat.go:474
	// _ = "end of CoverTab[6694]"
//line /usr/local/go/src/math/big/rat.go:474
	_go_fuzz_dep_.CoverTab[6695]++
						z.abs = z.abs.mul(x.abs, f)
						z.neg = x.neg
//line /usr/local/go/src/math/big/rat.go:476
	// _ = "end of CoverTab[6695]"
}

// Cmp compares x and y and returns:
//line /usr/local/go/src/math/big/rat.go:479
//
//line /usr/local/go/src/math/big/rat.go:479
//	-1 if x <  y
//line /usr/local/go/src/math/big/rat.go:479
//	 0 if x == y
//line /usr/local/go/src/math/big/rat.go:479
//	+1 if x >  y
//line /usr/local/go/src/math/big/rat.go:484
func (x *Rat) Cmp(y *Rat) int {
//line /usr/local/go/src/math/big/rat.go:484
	_go_fuzz_dep_.CoverTab[6698]++
						var a, b Int
						a.scaleDenom(&x.a, y.b.abs)
						b.scaleDenom(&y.a, x.b.abs)
						return a.Cmp(&b)
//line /usr/local/go/src/math/big/rat.go:488
	// _ = "end of CoverTab[6698]"
}

// Add sets z to the sum x+y and returns z.
func (z *Rat) Add(x, y *Rat) *Rat {
//line /usr/local/go/src/math/big/rat.go:492
	_go_fuzz_dep_.CoverTab[6699]++
						var a1, a2 Int
						a1.scaleDenom(&x.a, y.b.abs)
						a2.scaleDenom(&y.a, x.b.abs)
						z.a.Add(&a1, &a2)
						z.b.abs = mulDenom(z.b.abs, x.b.abs, y.b.abs)
						return z.norm()
//line /usr/local/go/src/math/big/rat.go:498
	// _ = "end of CoverTab[6699]"
}

// Sub sets z to the difference x-y and returns z.
func (z *Rat) Sub(x, y *Rat) *Rat {
//line /usr/local/go/src/math/big/rat.go:502
	_go_fuzz_dep_.CoverTab[6700]++
						var a1, a2 Int
						a1.scaleDenom(&x.a, y.b.abs)
						a2.scaleDenom(&y.a, x.b.abs)
						z.a.Sub(&a1, &a2)
						z.b.abs = mulDenom(z.b.abs, x.b.abs, y.b.abs)
						return z.norm()
//line /usr/local/go/src/math/big/rat.go:508
	// _ = "end of CoverTab[6700]"
}

// Mul sets z to the product x*y and returns z.
func (z *Rat) Mul(x, y *Rat) *Rat {
//line /usr/local/go/src/math/big/rat.go:512
	_go_fuzz_dep_.CoverTab[6701]++
						if x == y {
//line /usr/local/go/src/math/big/rat.go:513
		_go_fuzz_dep_.CoverTab[6703]++

							z.a.neg = false
							z.a.abs = z.a.abs.sqr(x.a.abs)
							if len(x.b.abs) == 0 {
//line /usr/local/go/src/math/big/rat.go:517
			_go_fuzz_dep_.CoverTab[6705]++
								z.b.abs = z.b.abs.setWord(1)
//line /usr/local/go/src/math/big/rat.go:518
			// _ = "end of CoverTab[6705]"
		} else {
//line /usr/local/go/src/math/big/rat.go:519
			_go_fuzz_dep_.CoverTab[6706]++
								z.b.abs = z.b.abs.sqr(x.b.abs)
//line /usr/local/go/src/math/big/rat.go:520
			// _ = "end of CoverTab[6706]"
		}
//line /usr/local/go/src/math/big/rat.go:521
		// _ = "end of CoverTab[6703]"
//line /usr/local/go/src/math/big/rat.go:521
		_go_fuzz_dep_.CoverTab[6704]++
							return z
//line /usr/local/go/src/math/big/rat.go:522
		// _ = "end of CoverTab[6704]"
	} else {
//line /usr/local/go/src/math/big/rat.go:523
		_go_fuzz_dep_.CoverTab[6707]++
//line /usr/local/go/src/math/big/rat.go:523
		// _ = "end of CoverTab[6707]"
//line /usr/local/go/src/math/big/rat.go:523
	}
//line /usr/local/go/src/math/big/rat.go:523
	// _ = "end of CoverTab[6701]"
//line /usr/local/go/src/math/big/rat.go:523
	_go_fuzz_dep_.CoverTab[6702]++
						z.a.Mul(&x.a, &y.a)
						z.b.abs = mulDenom(z.b.abs, x.b.abs, y.b.abs)
						return z.norm()
//line /usr/local/go/src/math/big/rat.go:526
	// _ = "end of CoverTab[6702]"
}

// Quo sets z to the quotient x/y and returns z.
//line /usr/local/go/src/math/big/rat.go:529
// If y == 0, Quo panics.
//line /usr/local/go/src/math/big/rat.go:531
func (z *Rat) Quo(x, y *Rat) *Rat {
//line /usr/local/go/src/math/big/rat.go:531
	_go_fuzz_dep_.CoverTab[6708]++
						if len(y.a.abs) == 0 {
//line /usr/local/go/src/math/big/rat.go:532
		_go_fuzz_dep_.CoverTab[6710]++
							panic("division by zero")
//line /usr/local/go/src/math/big/rat.go:533
		// _ = "end of CoverTab[6710]"
	} else {
//line /usr/local/go/src/math/big/rat.go:534
		_go_fuzz_dep_.CoverTab[6711]++
//line /usr/local/go/src/math/big/rat.go:534
		// _ = "end of CoverTab[6711]"
//line /usr/local/go/src/math/big/rat.go:534
	}
//line /usr/local/go/src/math/big/rat.go:534
	// _ = "end of CoverTab[6708]"
//line /usr/local/go/src/math/big/rat.go:534
	_go_fuzz_dep_.CoverTab[6709]++
						var a, b Int
						a.scaleDenom(&x.a, y.b.abs)
						b.scaleDenom(&y.a, x.b.abs)
						z.a.abs = a.abs
						z.b.abs = b.abs
						z.a.neg = a.neg != b.neg
						return z.norm()
//line /usr/local/go/src/math/big/rat.go:541
	// _ = "end of CoverTab[6709]"
}

//line /usr/local/go/src/math/big/rat.go:542
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/rat.go:542
var _ = _go_fuzz_dep_.CoverTab
