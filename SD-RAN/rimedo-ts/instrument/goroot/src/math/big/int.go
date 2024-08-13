// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements signed multi-precision integers.

//line /usr/local/go/src/math/big/int.go:7
package big

//line /usr/local/go/src/math/big/int.go:7
import (
//line /usr/local/go/src/math/big/int.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/int.go:7
)
//line /usr/local/go/src/math/big/int.go:7
import (
//line /usr/local/go/src/math/big/int.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/int.go:7
)

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

// An Int represents a signed multi-precision integer.
//line /usr/local/go/src/math/big/int.go:16
// The zero value for an Int represents the value 0.
//line /usr/local/go/src/math/big/int.go:16
//
//line /usr/local/go/src/math/big/int.go:16
// Operations always take pointer arguments (*Int) rather
//line /usr/local/go/src/math/big/int.go:16
// than Int values, and each unique Int value requires
//line /usr/local/go/src/math/big/int.go:16
// its own unique *Int pointer. To "copy" an Int value,
//line /usr/local/go/src/math/big/int.go:16
// an existing (or newly allocated) Int must be set to
//line /usr/local/go/src/math/big/int.go:16
// a new value using the Int.Set method; shallow copies
//line /usr/local/go/src/math/big/int.go:16
// of Ints are not supported and may lead to errors.
//line /usr/local/go/src/math/big/int.go:25
type Int struct {
	neg	bool	// sign
	abs	nat	// absolute value of the integer
}

var intOne = &Int{false, natOne}

// Sign returns:
//line /usr/local/go/src/math/big/int.go:32
//
//line /usr/local/go/src/math/big/int.go:32
//	-1 if x <  0
//line /usr/local/go/src/math/big/int.go:32
//	 0 if x == 0
//line /usr/local/go/src/math/big/int.go:32
//	+1 if x >  0
//line /usr/local/go/src/math/big/int.go:37
func (x *Int) Sign() int {
//line /usr/local/go/src/math/big/int.go:37
	_go_fuzz_dep_.CoverTab[5156]++

//line /usr/local/go/src/math/big/int.go:41
	if len(x.abs) == 0 {
//line /usr/local/go/src/math/big/int.go:41
		_go_fuzz_dep_.CoverTab[5159]++
							return 0
//line /usr/local/go/src/math/big/int.go:42
		// _ = "end of CoverTab[5159]"
	} else {
//line /usr/local/go/src/math/big/int.go:43
		_go_fuzz_dep_.CoverTab[5160]++
//line /usr/local/go/src/math/big/int.go:43
		// _ = "end of CoverTab[5160]"
//line /usr/local/go/src/math/big/int.go:43
	}
//line /usr/local/go/src/math/big/int.go:43
	// _ = "end of CoverTab[5156]"
//line /usr/local/go/src/math/big/int.go:43
	_go_fuzz_dep_.CoverTab[5157]++
						if x.neg {
//line /usr/local/go/src/math/big/int.go:44
		_go_fuzz_dep_.CoverTab[5161]++
							return -1
//line /usr/local/go/src/math/big/int.go:45
		// _ = "end of CoverTab[5161]"
	} else {
//line /usr/local/go/src/math/big/int.go:46
		_go_fuzz_dep_.CoverTab[5162]++
//line /usr/local/go/src/math/big/int.go:46
		// _ = "end of CoverTab[5162]"
//line /usr/local/go/src/math/big/int.go:46
	}
//line /usr/local/go/src/math/big/int.go:46
	// _ = "end of CoverTab[5157]"
//line /usr/local/go/src/math/big/int.go:46
	_go_fuzz_dep_.CoverTab[5158]++
						return 1
//line /usr/local/go/src/math/big/int.go:47
	// _ = "end of CoverTab[5158]"
}

// SetInt64 sets z to x and returns z.
func (z *Int) SetInt64(x int64) *Int {
//line /usr/local/go/src/math/big/int.go:51
	_go_fuzz_dep_.CoverTab[5163]++
						neg := false
						if x < 0 {
//line /usr/local/go/src/math/big/int.go:53
		_go_fuzz_dep_.CoverTab[5165]++
							neg = true
							x = -x
//line /usr/local/go/src/math/big/int.go:55
		// _ = "end of CoverTab[5165]"
	} else {
//line /usr/local/go/src/math/big/int.go:56
		_go_fuzz_dep_.CoverTab[5166]++
//line /usr/local/go/src/math/big/int.go:56
		// _ = "end of CoverTab[5166]"
//line /usr/local/go/src/math/big/int.go:56
	}
//line /usr/local/go/src/math/big/int.go:56
	// _ = "end of CoverTab[5163]"
//line /usr/local/go/src/math/big/int.go:56
	_go_fuzz_dep_.CoverTab[5164]++
						z.abs = z.abs.setUint64(uint64(x))
						z.neg = neg
						return z
//line /usr/local/go/src/math/big/int.go:59
	// _ = "end of CoverTab[5164]"
}

// SetUint64 sets z to x and returns z.
func (z *Int) SetUint64(x uint64) *Int {
//line /usr/local/go/src/math/big/int.go:63
	_go_fuzz_dep_.CoverTab[5167]++
						z.abs = z.abs.setUint64(x)
						z.neg = false
						return z
//line /usr/local/go/src/math/big/int.go:66
	// _ = "end of CoverTab[5167]"
}

// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) *Int {
//line /usr/local/go/src/math/big/int.go:70
	_go_fuzz_dep_.CoverTab[5168]++

//line /usr/local/go/src/math/big/int.go:73
	u := uint64(x)
	if x < 0 {
//line /usr/local/go/src/math/big/int.go:74
		_go_fuzz_dep_.CoverTab[5171]++
							u = -u
//line /usr/local/go/src/math/big/int.go:75
		// _ = "end of CoverTab[5171]"
	} else {
//line /usr/local/go/src/math/big/int.go:76
		_go_fuzz_dep_.CoverTab[5172]++
//line /usr/local/go/src/math/big/int.go:76
		// _ = "end of CoverTab[5172]"
//line /usr/local/go/src/math/big/int.go:76
	}
//line /usr/local/go/src/math/big/int.go:76
	// _ = "end of CoverTab[5168]"
//line /usr/local/go/src/math/big/int.go:76
	_go_fuzz_dep_.CoverTab[5169]++
						var abs []Word
						if x == 0 {
//line /usr/local/go/src/math/big/int.go:78
		_go_fuzz_dep_.CoverTab[5173]++
//line /usr/local/go/src/math/big/int.go:78
		// _ = "end of CoverTab[5173]"
	} else {
//line /usr/local/go/src/math/big/int.go:79
		_go_fuzz_dep_.CoverTab[5174]++
//line /usr/local/go/src/math/big/int.go:79
		if _W == 32 && func() bool {
//line /usr/local/go/src/math/big/int.go:79
			_go_fuzz_dep_.CoverTab[5175]++
//line /usr/local/go/src/math/big/int.go:79
			return u>>32 != 0
//line /usr/local/go/src/math/big/int.go:79
			// _ = "end of CoverTab[5175]"
//line /usr/local/go/src/math/big/int.go:79
		}() {
//line /usr/local/go/src/math/big/int.go:79
			_go_fuzz_dep_.CoverTab[5176]++
								abs = []Word{Word(u), Word(u >> 32)}
//line /usr/local/go/src/math/big/int.go:80
			// _ = "end of CoverTab[5176]"
		} else {
//line /usr/local/go/src/math/big/int.go:81
			_go_fuzz_dep_.CoverTab[5177]++
								abs = []Word{Word(u)}
//line /usr/local/go/src/math/big/int.go:82
			// _ = "end of CoverTab[5177]"
		}
//line /usr/local/go/src/math/big/int.go:83
		// _ = "end of CoverTab[5174]"
//line /usr/local/go/src/math/big/int.go:83
	}
//line /usr/local/go/src/math/big/int.go:83
	// _ = "end of CoverTab[5169]"
//line /usr/local/go/src/math/big/int.go:83
	_go_fuzz_dep_.CoverTab[5170]++
						return &Int{neg: x < 0, abs: abs}
//line /usr/local/go/src/math/big/int.go:84
	// _ = "end of CoverTab[5170]"
}

// Set sets z to x and returns z.
func (z *Int) Set(x *Int) *Int {
//line /usr/local/go/src/math/big/int.go:88
	_go_fuzz_dep_.CoverTab[5178]++
						if z != x {
//line /usr/local/go/src/math/big/int.go:89
		_go_fuzz_dep_.CoverTab[5180]++
							z.abs = z.abs.set(x.abs)
							z.neg = x.neg
//line /usr/local/go/src/math/big/int.go:91
		// _ = "end of CoverTab[5180]"
	} else {
//line /usr/local/go/src/math/big/int.go:92
		_go_fuzz_dep_.CoverTab[5181]++
//line /usr/local/go/src/math/big/int.go:92
		// _ = "end of CoverTab[5181]"
//line /usr/local/go/src/math/big/int.go:92
	}
//line /usr/local/go/src/math/big/int.go:92
	// _ = "end of CoverTab[5178]"
//line /usr/local/go/src/math/big/int.go:92
	_go_fuzz_dep_.CoverTab[5179]++
						return z
//line /usr/local/go/src/math/big/int.go:93
	// _ = "end of CoverTab[5179]"
}

// Bits provides raw (unchecked but fast) access to x by returning its
//line /usr/local/go/src/math/big/int.go:96
// absolute value as a little-endian Word slice. The result and x share
//line /usr/local/go/src/math/big/int.go:96
// the same underlying array.
//line /usr/local/go/src/math/big/int.go:96
// Bits is intended to support implementation of missing low-level Int
//line /usr/local/go/src/math/big/int.go:96
// functionality outside this package; it should be avoided otherwise.
//line /usr/local/go/src/math/big/int.go:101
func (x *Int) Bits() []Word {
//line /usr/local/go/src/math/big/int.go:101
	_go_fuzz_dep_.CoverTab[5182]++

//line /usr/local/go/src/math/big/int.go:105
	return x.abs
//line /usr/local/go/src/math/big/int.go:105
	// _ = "end of CoverTab[5182]"
}

// SetBits provides raw (unchecked but fast) access to z by setting its
//line /usr/local/go/src/math/big/int.go:108
// value to abs, interpreted as a little-endian Word slice, and returning
//line /usr/local/go/src/math/big/int.go:108
// z. The result and abs share the same underlying array.
//line /usr/local/go/src/math/big/int.go:108
// SetBits is intended to support implementation of missing low-level Int
//line /usr/local/go/src/math/big/int.go:108
// functionality outside this package; it should be avoided otherwise.
//line /usr/local/go/src/math/big/int.go:113
func (z *Int) SetBits(abs []Word) *Int {
//line /usr/local/go/src/math/big/int.go:113
	_go_fuzz_dep_.CoverTab[5183]++
						z.abs = nat(abs).norm()
						z.neg = false
						return z
//line /usr/local/go/src/math/big/int.go:116
	// _ = "end of CoverTab[5183]"
}

// Abs sets z to |x| (the absolute value of x) and returns z.
func (z *Int) Abs(x *Int) *Int {
//line /usr/local/go/src/math/big/int.go:120
	_go_fuzz_dep_.CoverTab[5184]++
						z.Set(x)
						z.neg = false
						return z
//line /usr/local/go/src/math/big/int.go:123
	// _ = "end of CoverTab[5184]"
}

// Neg sets z to -x and returns z.
func (z *Int) Neg(x *Int) *Int {
//line /usr/local/go/src/math/big/int.go:127
	_go_fuzz_dep_.CoverTab[5185]++
						z.Set(x)
						z.neg = len(z.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:129
		_go_fuzz_dep_.CoverTab[5186]++
//line /usr/local/go/src/math/big/int.go:129
		return !z.neg
//line /usr/local/go/src/math/big/int.go:129
		// _ = "end of CoverTab[5186]"
//line /usr/local/go/src/math/big/int.go:129
	}()
						return z
//line /usr/local/go/src/math/big/int.go:130
	// _ = "end of CoverTab[5185]"
}

// Add sets z to the sum x+y and returns z.
func (z *Int) Add(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:134
	_go_fuzz_dep_.CoverTab[5187]++
						neg := x.neg
						if x.neg == y.neg {
//line /usr/local/go/src/math/big/int.go:136
		_go_fuzz_dep_.CoverTab[5189]++

//line /usr/local/go/src/math/big/int.go:139
		z.abs = z.abs.add(x.abs, y.abs)
//line /usr/local/go/src/math/big/int.go:139
		// _ = "end of CoverTab[5189]"
	} else {
//line /usr/local/go/src/math/big/int.go:140
		_go_fuzz_dep_.CoverTab[5190]++

//line /usr/local/go/src/math/big/int.go:143
		if x.abs.cmp(y.abs) >= 0 {
//line /usr/local/go/src/math/big/int.go:143
			_go_fuzz_dep_.CoverTab[5191]++
								z.abs = z.abs.sub(x.abs, y.abs)
//line /usr/local/go/src/math/big/int.go:144
			// _ = "end of CoverTab[5191]"
		} else {
//line /usr/local/go/src/math/big/int.go:145
			_go_fuzz_dep_.CoverTab[5192]++
								neg = !neg
								z.abs = z.abs.sub(y.abs, x.abs)
//line /usr/local/go/src/math/big/int.go:147
			// _ = "end of CoverTab[5192]"
		}
//line /usr/local/go/src/math/big/int.go:148
		// _ = "end of CoverTab[5190]"
	}
//line /usr/local/go/src/math/big/int.go:149
	// _ = "end of CoverTab[5187]"
//line /usr/local/go/src/math/big/int.go:149
	_go_fuzz_dep_.CoverTab[5188]++
						z.neg = len(z.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:150
		_go_fuzz_dep_.CoverTab[5193]++
//line /usr/local/go/src/math/big/int.go:150
		return neg
//line /usr/local/go/src/math/big/int.go:150
		// _ = "end of CoverTab[5193]"
//line /usr/local/go/src/math/big/int.go:150
	}()
						return z
//line /usr/local/go/src/math/big/int.go:151
	// _ = "end of CoverTab[5188]"
}

// Sub sets z to the difference x-y and returns z.
func (z *Int) Sub(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:155
	_go_fuzz_dep_.CoverTab[5194]++
						neg := x.neg
						if x.neg != y.neg {
//line /usr/local/go/src/math/big/int.go:157
		_go_fuzz_dep_.CoverTab[5196]++

//line /usr/local/go/src/math/big/int.go:160
		z.abs = z.abs.add(x.abs, y.abs)
//line /usr/local/go/src/math/big/int.go:160
		// _ = "end of CoverTab[5196]"
	} else {
//line /usr/local/go/src/math/big/int.go:161
		_go_fuzz_dep_.CoverTab[5197]++

//line /usr/local/go/src/math/big/int.go:164
		if x.abs.cmp(y.abs) >= 0 {
//line /usr/local/go/src/math/big/int.go:164
			_go_fuzz_dep_.CoverTab[5198]++
								z.abs = z.abs.sub(x.abs, y.abs)
//line /usr/local/go/src/math/big/int.go:165
			// _ = "end of CoverTab[5198]"
		} else {
//line /usr/local/go/src/math/big/int.go:166
			_go_fuzz_dep_.CoverTab[5199]++
								neg = !neg
								z.abs = z.abs.sub(y.abs, x.abs)
//line /usr/local/go/src/math/big/int.go:168
			// _ = "end of CoverTab[5199]"
		}
//line /usr/local/go/src/math/big/int.go:169
		// _ = "end of CoverTab[5197]"
	}
//line /usr/local/go/src/math/big/int.go:170
	// _ = "end of CoverTab[5194]"
//line /usr/local/go/src/math/big/int.go:170
	_go_fuzz_dep_.CoverTab[5195]++
						z.neg = len(z.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:171
		_go_fuzz_dep_.CoverTab[5200]++
//line /usr/local/go/src/math/big/int.go:171
		return neg
//line /usr/local/go/src/math/big/int.go:171
		// _ = "end of CoverTab[5200]"
//line /usr/local/go/src/math/big/int.go:171
	}()
						return z
//line /usr/local/go/src/math/big/int.go:172
	// _ = "end of CoverTab[5195]"
}

// Mul sets z to the product x*y and returns z.
func (z *Int) Mul(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:176
	_go_fuzz_dep_.CoverTab[5201]++

//line /usr/local/go/src/math/big/int.go:181
	if x == y {
//line /usr/local/go/src/math/big/int.go:181
		_go_fuzz_dep_.CoverTab[5203]++
							z.abs = z.abs.sqr(x.abs)
							z.neg = false
							return z
//line /usr/local/go/src/math/big/int.go:184
		// _ = "end of CoverTab[5203]"
	} else {
//line /usr/local/go/src/math/big/int.go:185
		_go_fuzz_dep_.CoverTab[5204]++
//line /usr/local/go/src/math/big/int.go:185
		// _ = "end of CoverTab[5204]"
//line /usr/local/go/src/math/big/int.go:185
	}
//line /usr/local/go/src/math/big/int.go:185
	// _ = "end of CoverTab[5201]"
//line /usr/local/go/src/math/big/int.go:185
	_go_fuzz_dep_.CoverTab[5202]++
						z.abs = z.abs.mul(x.abs, y.abs)
						z.neg = len(z.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:187
		_go_fuzz_dep_.CoverTab[5205]++
//line /usr/local/go/src/math/big/int.go:187
		return x.neg != y.neg
//line /usr/local/go/src/math/big/int.go:187
		// _ = "end of CoverTab[5205]"
//line /usr/local/go/src/math/big/int.go:187
	}()
						return z
//line /usr/local/go/src/math/big/int.go:188
	// _ = "end of CoverTab[5202]"
}

// MulRange sets z to the product of all integers
//line /usr/local/go/src/math/big/int.go:191
// in the range [a, b] inclusively and returns z.
//line /usr/local/go/src/math/big/int.go:191
// If a > b (empty range), the result is 1.
//line /usr/local/go/src/math/big/int.go:194
func (z *Int) MulRange(a, b int64) *Int {
//line /usr/local/go/src/math/big/int.go:194
	_go_fuzz_dep_.CoverTab[5206]++
						switch {
	case a > b:
//line /usr/local/go/src/math/big/int.go:196
		_go_fuzz_dep_.CoverTab[5209]++
							return z.SetInt64(1)
//line /usr/local/go/src/math/big/int.go:197
		// _ = "end of CoverTab[5209]"
	case a <= 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:198
		_go_fuzz_dep_.CoverTab[5212]++
//line /usr/local/go/src/math/big/int.go:198
		return b >= 0
//line /usr/local/go/src/math/big/int.go:198
		// _ = "end of CoverTab[5212]"
//line /usr/local/go/src/math/big/int.go:198
	}():
//line /usr/local/go/src/math/big/int.go:198
		_go_fuzz_dep_.CoverTab[5210]++
							return z.SetInt64(0)
//line /usr/local/go/src/math/big/int.go:199
		// _ = "end of CoverTab[5210]"
//line /usr/local/go/src/math/big/int.go:199
	default:
//line /usr/local/go/src/math/big/int.go:199
		_go_fuzz_dep_.CoverTab[5211]++
//line /usr/local/go/src/math/big/int.go:199
		// _ = "end of CoverTab[5211]"
	}
//line /usr/local/go/src/math/big/int.go:200
	// _ = "end of CoverTab[5206]"
//line /usr/local/go/src/math/big/int.go:200
	_go_fuzz_dep_.CoverTab[5207]++

//line /usr/local/go/src/math/big/int.go:203
	neg := false
	if a < 0 {
//line /usr/local/go/src/math/big/int.go:204
		_go_fuzz_dep_.CoverTab[5213]++
							neg = (b-a)&1 == 0
							a, b = -b, -a
//line /usr/local/go/src/math/big/int.go:206
		// _ = "end of CoverTab[5213]"
	} else {
//line /usr/local/go/src/math/big/int.go:207
		_go_fuzz_dep_.CoverTab[5214]++
//line /usr/local/go/src/math/big/int.go:207
		// _ = "end of CoverTab[5214]"
//line /usr/local/go/src/math/big/int.go:207
	}
//line /usr/local/go/src/math/big/int.go:207
	// _ = "end of CoverTab[5207]"
//line /usr/local/go/src/math/big/int.go:207
	_go_fuzz_dep_.CoverTab[5208]++

						z.abs = z.abs.mulRange(uint64(a), uint64(b))
						z.neg = neg
						return z
//line /usr/local/go/src/math/big/int.go:211
	// _ = "end of CoverTab[5208]"
}

// Binomial sets z to the binomial coefficient C(n, k) and returns z.
func (z *Int) Binomial(n, k int64) *Int {
//line /usr/local/go/src/math/big/int.go:215
	_go_fuzz_dep_.CoverTab[5215]++
						if k > n {
//line /usr/local/go/src/math/big/int.go:216
		_go_fuzz_dep_.CoverTab[5219]++
							return z.SetInt64(0)
//line /usr/local/go/src/math/big/int.go:217
		// _ = "end of CoverTab[5219]"
	} else {
//line /usr/local/go/src/math/big/int.go:218
		_go_fuzz_dep_.CoverTab[5220]++
//line /usr/local/go/src/math/big/int.go:218
		// _ = "end of CoverTab[5220]"
//line /usr/local/go/src/math/big/int.go:218
	}
//line /usr/local/go/src/math/big/int.go:218
	// _ = "end of CoverTab[5215]"
//line /usr/local/go/src/math/big/int.go:218
	_go_fuzz_dep_.CoverTab[5216]++

						if k > n-k {
//line /usr/local/go/src/math/big/int.go:220
		_go_fuzz_dep_.CoverTab[5221]++
							k = n - k
//line /usr/local/go/src/math/big/int.go:221
		// _ = "end of CoverTab[5221]"
	} else {
//line /usr/local/go/src/math/big/int.go:222
		_go_fuzz_dep_.CoverTab[5222]++
//line /usr/local/go/src/math/big/int.go:222
		// _ = "end of CoverTab[5222]"
//line /usr/local/go/src/math/big/int.go:222
	}
//line /usr/local/go/src/math/big/int.go:222
	// _ = "end of CoverTab[5216]"
//line /usr/local/go/src/math/big/int.go:222
	_go_fuzz_dep_.CoverTab[5217]++
	// C(n, k) == n * (n-1) * ... * (n-k+1) / k * (k-1) * ... * 1
	//         == n * (n-1) * ... * (n-k+1) / 1 * (1+1) * ... * k
	//
	// Using the multiplicative formula produces smaller values
	// at each step, requiring fewer allocations and computations:
	//
	// z = 1
	// for i := 0; i < k; i = i+1 {
	//     z *= n-i
	//     z /= i+1
	// }
	//
	// finally to avoid computing i+1 twice per loop:
	//
	// z = 1
	// i := 0
	// for i < k {
	//     z *= n-i
	//     i++
	//     z /= i
	// }
	var N, K, i, t Int
	N.SetInt64(n)
	K.SetInt64(k)
	z.Set(intOne)
	for i.Cmp(&K) < 0 {
//line /usr/local/go/src/math/big/int.go:248
		_go_fuzz_dep_.CoverTab[5223]++
							z.Mul(z, t.Sub(&N, &i))
							i.Add(&i, intOne)
							z.Quo(z, &i)
//line /usr/local/go/src/math/big/int.go:251
		// _ = "end of CoverTab[5223]"
	}
//line /usr/local/go/src/math/big/int.go:252
	// _ = "end of CoverTab[5217]"
//line /usr/local/go/src/math/big/int.go:252
	_go_fuzz_dep_.CoverTab[5218]++
						return z
//line /usr/local/go/src/math/big/int.go:253
	// _ = "end of CoverTab[5218]"
}

// Quo sets z to the quotient x/y for y != 0 and returns z.
//line /usr/local/go/src/math/big/int.go:256
// If y == 0, a division-by-zero run-time panic occurs.
//line /usr/local/go/src/math/big/int.go:256
// Quo implements truncated division (like Go); see QuoRem for more details.
//line /usr/local/go/src/math/big/int.go:259
func (z *Int) Quo(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:259
	_go_fuzz_dep_.CoverTab[5224]++
						z.abs, _ = z.abs.div(nil, x.abs, y.abs)
						z.neg = len(z.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:261
		_go_fuzz_dep_.CoverTab[5225]++
//line /usr/local/go/src/math/big/int.go:261
		return x.neg != y.neg
//line /usr/local/go/src/math/big/int.go:261
		// _ = "end of CoverTab[5225]"
//line /usr/local/go/src/math/big/int.go:261
	}()
						return z
//line /usr/local/go/src/math/big/int.go:262
	// _ = "end of CoverTab[5224]"
}

// Rem sets z to the remainder x%y for y != 0 and returns z.
//line /usr/local/go/src/math/big/int.go:265
// If y == 0, a division-by-zero run-time panic occurs.
//line /usr/local/go/src/math/big/int.go:265
// Rem implements truncated modulus (like Go); see QuoRem for more details.
//line /usr/local/go/src/math/big/int.go:268
func (z *Int) Rem(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:268
	_go_fuzz_dep_.CoverTab[5226]++
						_, z.abs = nat(nil).div(z.abs, x.abs, y.abs)
						z.neg = len(z.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:270
		_go_fuzz_dep_.CoverTab[5227]++
//line /usr/local/go/src/math/big/int.go:270
		return x.neg
//line /usr/local/go/src/math/big/int.go:270
		// _ = "end of CoverTab[5227]"
//line /usr/local/go/src/math/big/int.go:270
	}()
						return z
//line /usr/local/go/src/math/big/int.go:271
	// _ = "end of CoverTab[5226]"
}

// QuoRem sets z to the quotient x/y and r to the remainder x%y
//line /usr/local/go/src/math/big/int.go:274
// and returns the pair (z, r) for y != 0.
//line /usr/local/go/src/math/big/int.go:274
// If y == 0, a division-by-zero run-time panic occurs.
//line /usr/local/go/src/math/big/int.go:274
//
//line /usr/local/go/src/math/big/int.go:274
// QuoRem implements T-division and modulus (like Go):
//line /usr/local/go/src/math/big/int.go:274
//
//line /usr/local/go/src/math/big/int.go:274
//	q = x/y      with the result truncated to zero
//line /usr/local/go/src/math/big/int.go:274
//	r = x - y*q
//line /usr/local/go/src/math/big/int.go:274
//
//line /usr/local/go/src/math/big/int.go:274
// (See Daan Leijen, “Division and Modulus for Computer Scientists”.)
//line /usr/local/go/src/math/big/int.go:274
// See DivMod for Euclidean division and modulus (unlike Go).
//line /usr/local/go/src/math/big/int.go:285
func (z *Int) QuoRem(x, y, r *Int) (*Int, *Int) {
//line /usr/local/go/src/math/big/int.go:285
	_go_fuzz_dep_.CoverTab[5228]++
						z.abs, r.abs = z.abs.div(r.abs, x.abs, y.abs)
						z.neg, r.neg = len(z.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:287
		_go_fuzz_dep_.CoverTab[5229]++
//line /usr/local/go/src/math/big/int.go:287
		return x.neg != y.neg
//line /usr/local/go/src/math/big/int.go:287
		// _ = "end of CoverTab[5229]"
//line /usr/local/go/src/math/big/int.go:287
	}(), len(r.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:287
		_go_fuzz_dep_.CoverTab[5230]++
//line /usr/local/go/src/math/big/int.go:287
		return x.neg
//line /usr/local/go/src/math/big/int.go:287
		// _ = "end of CoverTab[5230]"
//line /usr/local/go/src/math/big/int.go:287
	}()
						return z, r
//line /usr/local/go/src/math/big/int.go:288
	// _ = "end of CoverTab[5228]"
}

// Div sets z to the quotient x/y for y != 0 and returns z.
//line /usr/local/go/src/math/big/int.go:291
// If y == 0, a division-by-zero run-time panic occurs.
//line /usr/local/go/src/math/big/int.go:291
// Div implements Euclidean division (unlike Go); see DivMod for more details.
//line /usr/local/go/src/math/big/int.go:294
func (z *Int) Div(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:294
	_go_fuzz_dep_.CoverTab[5231]++
						y_neg := y.neg
						var r Int
						z.QuoRem(x, y, &r)
						if r.neg {
//line /usr/local/go/src/math/big/int.go:298
		_go_fuzz_dep_.CoverTab[5233]++
							if y_neg {
//line /usr/local/go/src/math/big/int.go:299
			_go_fuzz_dep_.CoverTab[5234]++
								z.Add(z, intOne)
//line /usr/local/go/src/math/big/int.go:300
			// _ = "end of CoverTab[5234]"
		} else {
//line /usr/local/go/src/math/big/int.go:301
			_go_fuzz_dep_.CoverTab[5235]++
								z.Sub(z, intOne)
//line /usr/local/go/src/math/big/int.go:302
			// _ = "end of CoverTab[5235]"
		}
//line /usr/local/go/src/math/big/int.go:303
		// _ = "end of CoverTab[5233]"
	} else {
//line /usr/local/go/src/math/big/int.go:304
		_go_fuzz_dep_.CoverTab[5236]++
//line /usr/local/go/src/math/big/int.go:304
		// _ = "end of CoverTab[5236]"
//line /usr/local/go/src/math/big/int.go:304
	}
//line /usr/local/go/src/math/big/int.go:304
	// _ = "end of CoverTab[5231]"
//line /usr/local/go/src/math/big/int.go:304
	_go_fuzz_dep_.CoverTab[5232]++
						return z
//line /usr/local/go/src/math/big/int.go:305
	// _ = "end of CoverTab[5232]"
}

// Mod sets z to the modulus x%y for y != 0 and returns z.
//line /usr/local/go/src/math/big/int.go:308
// If y == 0, a division-by-zero run-time panic occurs.
//line /usr/local/go/src/math/big/int.go:308
// Mod implements Euclidean modulus (unlike Go); see DivMod for more details.
//line /usr/local/go/src/math/big/int.go:311
func (z *Int) Mod(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:311
	_go_fuzz_dep_.CoverTab[5237]++
						y0 := y
						if z == y || func() bool {
//line /usr/local/go/src/math/big/int.go:313
		_go_fuzz_dep_.CoverTab[5240]++
//line /usr/local/go/src/math/big/int.go:313
		return alias(z.abs, y.abs)
//line /usr/local/go/src/math/big/int.go:313
		// _ = "end of CoverTab[5240]"
//line /usr/local/go/src/math/big/int.go:313
	}() {
//line /usr/local/go/src/math/big/int.go:313
		_go_fuzz_dep_.CoverTab[5241]++
							y0 = new(Int).Set(y)
//line /usr/local/go/src/math/big/int.go:314
		// _ = "end of CoverTab[5241]"
	} else {
//line /usr/local/go/src/math/big/int.go:315
		_go_fuzz_dep_.CoverTab[5242]++
//line /usr/local/go/src/math/big/int.go:315
		// _ = "end of CoverTab[5242]"
//line /usr/local/go/src/math/big/int.go:315
	}
//line /usr/local/go/src/math/big/int.go:315
	// _ = "end of CoverTab[5237]"
//line /usr/local/go/src/math/big/int.go:315
	_go_fuzz_dep_.CoverTab[5238]++
						var q Int
						q.QuoRem(x, y, z)
						if z.neg {
//line /usr/local/go/src/math/big/int.go:318
		_go_fuzz_dep_.CoverTab[5243]++
							if y0.neg {
//line /usr/local/go/src/math/big/int.go:319
			_go_fuzz_dep_.CoverTab[5244]++
								z.Sub(z, y0)
//line /usr/local/go/src/math/big/int.go:320
			// _ = "end of CoverTab[5244]"
		} else {
//line /usr/local/go/src/math/big/int.go:321
			_go_fuzz_dep_.CoverTab[5245]++
								z.Add(z, y0)
//line /usr/local/go/src/math/big/int.go:322
			// _ = "end of CoverTab[5245]"
		}
//line /usr/local/go/src/math/big/int.go:323
		// _ = "end of CoverTab[5243]"
	} else {
//line /usr/local/go/src/math/big/int.go:324
		_go_fuzz_dep_.CoverTab[5246]++
//line /usr/local/go/src/math/big/int.go:324
		// _ = "end of CoverTab[5246]"
//line /usr/local/go/src/math/big/int.go:324
	}
//line /usr/local/go/src/math/big/int.go:324
	// _ = "end of CoverTab[5238]"
//line /usr/local/go/src/math/big/int.go:324
	_go_fuzz_dep_.CoverTab[5239]++
						return z
//line /usr/local/go/src/math/big/int.go:325
	// _ = "end of CoverTab[5239]"
}

// DivMod sets z to the quotient x div y and m to the modulus x mod y
//line /usr/local/go/src/math/big/int.go:328
// and returns the pair (z, m) for y != 0.
//line /usr/local/go/src/math/big/int.go:328
// If y == 0, a division-by-zero run-time panic occurs.
//line /usr/local/go/src/math/big/int.go:328
//
//line /usr/local/go/src/math/big/int.go:328
// DivMod implements Euclidean division and modulus (unlike Go):
//line /usr/local/go/src/math/big/int.go:328
//
//line /usr/local/go/src/math/big/int.go:328
//	q = x div y  such that
//line /usr/local/go/src/math/big/int.go:328
//	m = x - y*q  with 0 <= m < |y|
//line /usr/local/go/src/math/big/int.go:328
//
//line /usr/local/go/src/math/big/int.go:328
// (See Raymond T. Boute, “The Euclidean definition of the functions
//line /usr/local/go/src/math/big/int.go:328
// div and mod”. ACM Transactions on Programming Languages and
//line /usr/local/go/src/math/big/int.go:328
// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
//line /usr/local/go/src/math/big/int.go:328
// ACM press.)
//line /usr/local/go/src/math/big/int.go:328
// See QuoRem for T-division and modulus (like Go).
//line /usr/local/go/src/math/big/int.go:342
func (z *Int) DivMod(x, y, m *Int) (*Int, *Int) {
//line /usr/local/go/src/math/big/int.go:342
	_go_fuzz_dep_.CoverTab[5247]++
						y0 := y
						if z == y || func() bool {
//line /usr/local/go/src/math/big/int.go:344
		_go_fuzz_dep_.CoverTab[5250]++
//line /usr/local/go/src/math/big/int.go:344
		return alias(z.abs, y.abs)
//line /usr/local/go/src/math/big/int.go:344
		// _ = "end of CoverTab[5250]"
//line /usr/local/go/src/math/big/int.go:344
	}() {
//line /usr/local/go/src/math/big/int.go:344
		_go_fuzz_dep_.CoverTab[5251]++
							y0 = new(Int).Set(y)
//line /usr/local/go/src/math/big/int.go:345
		// _ = "end of CoverTab[5251]"
	} else {
//line /usr/local/go/src/math/big/int.go:346
		_go_fuzz_dep_.CoverTab[5252]++
//line /usr/local/go/src/math/big/int.go:346
		// _ = "end of CoverTab[5252]"
//line /usr/local/go/src/math/big/int.go:346
	}
//line /usr/local/go/src/math/big/int.go:346
	// _ = "end of CoverTab[5247]"
//line /usr/local/go/src/math/big/int.go:346
	_go_fuzz_dep_.CoverTab[5248]++
						z.QuoRem(x, y, m)
						if m.neg {
//line /usr/local/go/src/math/big/int.go:348
		_go_fuzz_dep_.CoverTab[5253]++
							if y0.neg {
//line /usr/local/go/src/math/big/int.go:349
			_go_fuzz_dep_.CoverTab[5254]++
								z.Add(z, intOne)
								m.Sub(m, y0)
//line /usr/local/go/src/math/big/int.go:351
			// _ = "end of CoverTab[5254]"
		} else {
//line /usr/local/go/src/math/big/int.go:352
			_go_fuzz_dep_.CoverTab[5255]++
								z.Sub(z, intOne)
								m.Add(m, y0)
//line /usr/local/go/src/math/big/int.go:354
			// _ = "end of CoverTab[5255]"
		}
//line /usr/local/go/src/math/big/int.go:355
		// _ = "end of CoverTab[5253]"
	} else {
//line /usr/local/go/src/math/big/int.go:356
		_go_fuzz_dep_.CoverTab[5256]++
//line /usr/local/go/src/math/big/int.go:356
		// _ = "end of CoverTab[5256]"
//line /usr/local/go/src/math/big/int.go:356
	}
//line /usr/local/go/src/math/big/int.go:356
	// _ = "end of CoverTab[5248]"
//line /usr/local/go/src/math/big/int.go:356
	_go_fuzz_dep_.CoverTab[5249]++
						return z, m
//line /usr/local/go/src/math/big/int.go:357
	// _ = "end of CoverTab[5249]"
}

// Cmp compares x and y and returns:
//line /usr/local/go/src/math/big/int.go:360
//
//line /usr/local/go/src/math/big/int.go:360
//	-1 if x <  y
//line /usr/local/go/src/math/big/int.go:360
//	 0 if x == y
//line /usr/local/go/src/math/big/int.go:360
//	+1 if x >  y
//line /usr/local/go/src/math/big/int.go:365
func (x *Int) Cmp(y *Int) (r int) {
//line /usr/local/go/src/math/big/int.go:365
	_go_fuzz_dep_.CoverTab[5257]++

//line /usr/local/go/src/math/big/int.go:370
	switch {
	case x == y:
//line /usr/local/go/src/math/big/int.go:371
		_go_fuzz_dep_.CoverTab[5259]++
//line /usr/local/go/src/math/big/int.go:371
		// _ = "end of CoverTab[5259]"

	case x.neg == y.neg:
//line /usr/local/go/src/math/big/int.go:373
		_go_fuzz_dep_.CoverTab[5260]++
							r = x.abs.cmp(y.abs)
							if x.neg {
//line /usr/local/go/src/math/big/int.go:375
			_go_fuzz_dep_.CoverTab[5263]++
								r = -r
//line /usr/local/go/src/math/big/int.go:376
			// _ = "end of CoverTab[5263]"
		} else {
//line /usr/local/go/src/math/big/int.go:377
			_go_fuzz_dep_.CoverTab[5264]++
//line /usr/local/go/src/math/big/int.go:377
			// _ = "end of CoverTab[5264]"
//line /usr/local/go/src/math/big/int.go:377
		}
//line /usr/local/go/src/math/big/int.go:377
		// _ = "end of CoverTab[5260]"
	case x.neg:
//line /usr/local/go/src/math/big/int.go:378
		_go_fuzz_dep_.CoverTab[5261]++
							r = -1
//line /usr/local/go/src/math/big/int.go:379
		// _ = "end of CoverTab[5261]"
	default:
//line /usr/local/go/src/math/big/int.go:380
		_go_fuzz_dep_.CoverTab[5262]++
							r = 1
//line /usr/local/go/src/math/big/int.go:381
		// _ = "end of CoverTab[5262]"
	}
//line /usr/local/go/src/math/big/int.go:382
	// _ = "end of CoverTab[5257]"
//line /usr/local/go/src/math/big/int.go:382
	_go_fuzz_dep_.CoverTab[5258]++
						return
//line /usr/local/go/src/math/big/int.go:383
	// _ = "end of CoverTab[5258]"
}

// CmpAbs compares the absolute values of x and y and returns:
//line /usr/local/go/src/math/big/int.go:386
//
//line /usr/local/go/src/math/big/int.go:386
//	-1 if |x| <  |y|
//line /usr/local/go/src/math/big/int.go:386
//	 0 if |x| == |y|
//line /usr/local/go/src/math/big/int.go:386
//	+1 if |x| >  |y|
//line /usr/local/go/src/math/big/int.go:391
func (x *Int) CmpAbs(y *Int) int {
//line /usr/local/go/src/math/big/int.go:391
	_go_fuzz_dep_.CoverTab[5265]++
						return x.abs.cmp(y.abs)
//line /usr/local/go/src/math/big/int.go:392
	// _ = "end of CoverTab[5265]"
}

// low32 returns the least significant 32 bits of x.
func low32(x nat) uint32 {
//line /usr/local/go/src/math/big/int.go:396
	_go_fuzz_dep_.CoverTab[5266]++
						if len(x) == 0 {
//line /usr/local/go/src/math/big/int.go:397
		_go_fuzz_dep_.CoverTab[5268]++
							return 0
//line /usr/local/go/src/math/big/int.go:398
		// _ = "end of CoverTab[5268]"
	} else {
//line /usr/local/go/src/math/big/int.go:399
		_go_fuzz_dep_.CoverTab[5269]++
//line /usr/local/go/src/math/big/int.go:399
		// _ = "end of CoverTab[5269]"
//line /usr/local/go/src/math/big/int.go:399
	}
//line /usr/local/go/src/math/big/int.go:399
	// _ = "end of CoverTab[5266]"
//line /usr/local/go/src/math/big/int.go:399
	_go_fuzz_dep_.CoverTab[5267]++
						return uint32(x[0])
//line /usr/local/go/src/math/big/int.go:400
	// _ = "end of CoverTab[5267]"
}

// low64 returns the least significant 64 bits of x.
func low64(x nat) uint64 {
//line /usr/local/go/src/math/big/int.go:404
	_go_fuzz_dep_.CoverTab[5270]++
						if len(x) == 0 {
//line /usr/local/go/src/math/big/int.go:405
		_go_fuzz_dep_.CoverTab[5273]++
							return 0
//line /usr/local/go/src/math/big/int.go:406
		// _ = "end of CoverTab[5273]"
	} else {
//line /usr/local/go/src/math/big/int.go:407
		_go_fuzz_dep_.CoverTab[5274]++
//line /usr/local/go/src/math/big/int.go:407
		// _ = "end of CoverTab[5274]"
//line /usr/local/go/src/math/big/int.go:407
	}
//line /usr/local/go/src/math/big/int.go:407
	// _ = "end of CoverTab[5270]"
//line /usr/local/go/src/math/big/int.go:407
	_go_fuzz_dep_.CoverTab[5271]++
						v := uint64(x[0])
						if _W == 32 && func() bool {
//line /usr/local/go/src/math/big/int.go:409
		_go_fuzz_dep_.CoverTab[5275]++
//line /usr/local/go/src/math/big/int.go:409
		return len(x) > 1
//line /usr/local/go/src/math/big/int.go:409
		// _ = "end of CoverTab[5275]"
//line /usr/local/go/src/math/big/int.go:409
	}() {
//line /usr/local/go/src/math/big/int.go:409
		_go_fuzz_dep_.CoverTab[5276]++
							return uint64(x[1])<<32 | v
//line /usr/local/go/src/math/big/int.go:410
		// _ = "end of CoverTab[5276]"
	} else {
//line /usr/local/go/src/math/big/int.go:411
		_go_fuzz_dep_.CoverTab[5277]++
//line /usr/local/go/src/math/big/int.go:411
		// _ = "end of CoverTab[5277]"
//line /usr/local/go/src/math/big/int.go:411
	}
//line /usr/local/go/src/math/big/int.go:411
	// _ = "end of CoverTab[5271]"
//line /usr/local/go/src/math/big/int.go:411
	_go_fuzz_dep_.CoverTab[5272]++
						return v
//line /usr/local/go/src/math/big/int.go:412
	// _ = "end of CoverTab[5272]"
}

// Int64 returns the int64 representation of x.
//line /usr/local/go/src/math/big/int.go:415
// If x cannot be represented in an int64, the result is undefined.
//line /usr/local/go/src/math/big/int.go:417
func (x *Int) Int64() int64 {
//line /usr/local/go/src/math/big/int.go:417
	_go_fuzz_dep_.CoverTab[5278]++
						v := int64(low64(x.abs))
						if x.neg {
//line /usr/local/go/src/math/big/int.go:419
		_go_fuzz_dep_.CoverTab[5280]++
							v = -v
//line /usr/local/go/src/math/big/int.go:420
		// _ = "end of CoverTab[5280]"
	} else {
//line /usr/local/go/src/math/big/int.go:421
		_go_fuzz_dep_.CoverTab[5281]++
//line /usr/local/go/src/math/big/int.go:421
		// _ = "end of CoverTab[5281]"
//line /usr/local/go/src/math/big/int.go:421
	}
//line /usr/local/go/src/math/big/int.go:421
	// _ = "end of CoverTab[5278]"
//line /usr/local/go/src/math/big/int.go:421
	_go_fuzz_dep_.CoverTab[5279]++
						return v
//line /usr/local/go/src/math/big/int.go:422
	// _ = "end of CoverTab[5279]"
}

// Uint64 returns the uint64 representation of x.
//line /usr/local/go/src/math/big/int.go:425
// If x cannot be represented in a uint64, the result is undefined.
//line /usr/local/go/src/math/big/int.go:427
func (x *Int) Uint64() uint64 {
//line /usr/local/go/src/math/big/int.go:427
	_go_fuzz_dep_.CoverTab[5282]++
						return low64(x.abs)
//line /usr/local/go/src/math/big/int.go:428
	// _ = "end of CoverTab[5282]"
}

// IsInt64 reports whether x can be represented as an int64.
func (x *Int) IsInt64() bool {
//line /usr/local/go/src/math/big/int.go:432
	_go_fuzz_dep_.CoverTab[5283]++
						if len(x.abs) <= 64/_W {
//line /usr/local/go/src/math/big/int.go:433
		_go_fuzz_dep_.CoverTab[5285]++
							w := int64(low64(x.abs))
							return w >= 0 || func() bool {
//line /usr/local/go/src/math/big/int.go:435
			_go_fuzz_dep_.CoverTab[5286]++
//line /usr/local/go/src/math/big/int.go:435
			return x.neg && func() bool {
//line /usr/local/go/src/math/big/int.go:435
				_go_fuzz_dep_.CoverTab[5287]++
//line /usr/local/go/src/math/big/int.go:435
				return w == -w
//line /usr/local/go/src/math/big/int.go:435
				// _ = "end of CoverTab[5287]"
//line /usr/local/go/src/math/big/int.go:435
			}()
//line /usr/local/go/src/math/big/int.go:435
			// _ = "end of CoverTab[5286]"
//line /usr/local/go/src/math/big/int.go:435
		}()
//line /usr/local/go/src/math/big/int.go:435
		// _ = "end of CoverTab[5285]"
	} else {
//line /usr/local/go/src/math/big/int.go:436
		_go_fuzz_dep_.CoverTab[5288]++
//line /usr/local/go/src/math/big/int.go:436
		// _ = "end of CoverTab[5288]"
//line /usr/local/go/src/math/big/int.go:436
	}
//line /usr/local/go/src/math/big/int.go:436
	// _ = "end of CoverTab[5283]"
//line /usr/local/go/src/math/big/int.go:436
	_go_fuzz_dep_.CoverTab[5284]++
						return false
//line /usr/local/go/src/math/big/int.go:437
	// _ = "end of CoverTab[5284]"
}

// IsUint64 reports whether x can be represented as a uint64.
func (x *Int) IsUint64() bool {
//line /usr/local/go/src/math/big/int.go:441
	_go_fuzz_dep_.CoverTab[5289]++
						return !x.neg && func() bool {
//line /usr/local/go/src/math/big/int.go:442
		_go_fuzz_dep_.CoverTab[5290]++
//line /usr/local/go/src/math/big/int.go:442
		return len(x.abs) <= 64/_W
//line /usr/local/go/src/math/big/int.go:442
		// _ = "end of CoverTab[5290]"
//line /usr/local/go/src/math/big/int.go:442
	}()
//line /usr/local/go/src/math/big/int.go:442
	// _ = "end of CoverTab[5289]"
}

// SetString sets z to the value of s, interpreted in the given base,
//line /usr/local/go/src/math/big/int.go:445
// and returns z and a boolean indicating success. The entire string
//line /usr/local/go/src/math/big/int.go:445
// (not just a prefix) must be valid for success. If SetString fails,
//line /usr/local/go/src/math/big/int.go:445
// the value of z is undefined but the returned value is nil.
//line /usr/local/go/src/math/big/int.go:445
//
//line /usr/local/go/src/math/big/int.go:445
// The base argument must be 0 or a value between 2 and MaxBase.
//line /usr/local/go/src/math/big/int.go:445
// For base 0, the number prefix determines the actual base: A prefix of
//line /usr/local/go/src/math/big/int.go:445
// “0b” or “0B” selects base 2, “0”, “0o” or “0O” selects base 8,
//line /usr/local/go/src/math/big/int.go:445
// and “0x” or “0X” selects base 16. Otherwise, the selected base is 10
//line /usr/local/go/src/math/big/int.go:445
// and no prefix is accepted.
//line /usr/local/go/src/math/big/int.go:445
//
//line /usr/local/go/src/math/big/int.go:445
// For bases <= 36, lower and upper case letters are considered the same:
//line /usr/local/go/src/math/big/int.go:445
// The letters 'a' to 'z' and 'A' to 'Z' represent digit values 10 to 35.
//line /usr/local/go/src/math/big/int.go:445
// For bases > 36, the upper case letters 'A' to 'Z' represent the digit
//line /usr/local/go/src/math/big/int.go:445
// values 36 to 61.
//line /usr/local/go/src/math/big/int.go:445
//
//line /usr/local/go/src/math/big/int.go:445
// For base 0, an underscore character “_” may appear between a base
//line /usr/local/go/src/math/big/int.go:445
// prefix and an adjacent digit, and between successive digits; such
//line /usr/local/go/src/math/big/int.go:445
// underscores do not change the value of the number.
//line /usr/local/go/src/math/big/int.go:445
// Incorrect placement of underscores is reported as an error if there
//line /usr/local/go/src/math/big/int.go:445
// are no other errors. If base != 0, underscores are not recognized
//line /usr/local/go/src/math/big/int.go:445
// and act like any other character that is not a valid digit.
//line /usr/local/go/src/math/big/int.go:467
func (z *Int) SetString(s string, base int) (*Int, bool) {
//line /usr/local/go/src/math/big/int.go:467
	_go_fuzz_dep_.CoverTab[5291]++
						return z.setFromScanner(strings.NewReader(s), base)
//line /usr/local/go/src/math/big/int.go:468
	// _ = "end of CoverTab[5291]"
}

// setFromScanner implements SetString given an io.ByteScanner.
//line /usr/local/go/src/math/big/int.go:471
// For documentation see comments of SetString.
//line /usr/local/go/src/math/big/int.go:473
func (z *Int) setFromScanner(r io.ByteScanner, base int) (*Int, bool) {
//line /usr/local/go/src/math/big/int.go:473
	_go_fuzz_dep_.CoverTab[5292]++
						if _, _, err := z.scan(r, base); err != nil {
//line /usr/local/go/src/math/big/int.go:474
		_go_fuzz_dep_.CoverTab[5295]++
							return nil, false
//line /usr/local/go/src/math/big/int.go:475
		// _ = "end of CoverTab[5295]"
	} else {
//line /usr/local/go/src/math/big/int.go:476
		_go_fuzz_dep_.CoverTab[5296]++
//line /usr/local/go/src/math/big/int.go:476
		// _ = "end of CoverTab[5296]"
//line /usr/local/go/src/math/big/int.go:476
	}
//line /usr/local/go/src/math/big/int.go:476
	// _ = "end of CoverTab[5292]"
//line /usr/local/go/src/math/big/int.go:476
	_go_fuzz_dep_.CoverTab[5293]++

						if _, err := r.ReadByte(); err != io.EOF {
//line /usr/local/go/src/math/big/int.go:478
		_go_fuzz_dep_.CoverTab[5297]++
							return nil, false
//line /usr/local/go/src/math/big/int.go:479
		// _ = "end of CoverTab[5297]"
	} else {
//line /usr/local/go/src/math/big/int.go:480
		_go_fuzz_dep_.CoverTab[5298]++
//line /usr/local/go/src/math/big/int.go:480
		// _ = "end of CoverTab[5298]"
//line /usr/local/go/src/math/big/int.go:480
	}
//line /usr/local/go/src/math/big/int.go:480
	// _ = "end of CoverTab[5293]"
//line /usr/local/go/src/math/big/int.go:480
	_go_fuzz_dep_.CoverTab[5294]++
						return z, true
//line /usr/local/go/src/math/big/int.go:481
	// _ = "end of CoverTab[5294]"
}

// SetBytes interprets buf as the bytes of a big-endian unsigned
//line /usr/local/go/src/math/big/int.go:484
// integer, sets z to that value, and returns z.
//line /usr/local/go/src/math/big/int.go:486
func (z *Int) SetBytes(buf []byte) *Int {
//line /usr/local/go/src/math/big/int.go:486
	_go_fuzz_dep_.CoverTab[5299]++
						z.abs = z.abs.setBytes(buf)
						z.neg = false
						return z
//line /usr/local/go/src/math/big/int.go:489
	// _ = "end of CoverTab[5299]"
}

// Bytes returns the absolute value of x as a big-endian byte slice.
//line /usr/local/go/src/math/big/int.go:492
//
//line /usr/local/go/src/math/big/int.go:492
// To use a fixed length slice, or a preallocated one, use FillBytes.
//line /usr/local/go/src/math/big/int.go:495
func (x *Int) Bytes() []byte {
//line /usr/local/go/src/math/big/int.go:495
	_go_fuzz_dep_.CoverTab[5300]++

//line /usr/local/go/src/math/big/int.go:499
	buf := make([]byte, len(x.abs)*_S)
						return buf[x.abs.bytes(buf):]
//line /usr/local/go/src/math/big/int.go:500
	// _ = "end of CoverTab[5300]"
}

// FillBytes sets buf to the absolute value of x, storing it as a zero-extended
//line /usr/local/go/src/math/big/int.go:503
// big-endian byte slice, and returns buf.
//line /usr/local/go/src/math/big/int.go:503
//
//line /usr/local/go/src/math/big/int.go:503
// If the absolute value of x doesn't fit in buf, FillBytes will panic.
//line /usr/local/go/src/math/big/int.go:507
func (x *Int) FillBytes(buf []byte) []byte {
//line /usr/local/go/src/math/big/int.go:507
	_go_fuzz_dep_.CoverTab[5301]++

						for i := range buf {
//line /usr/local/go/src/math/big/int.go:509
		_go_fuzz_dep_.CoverTab[5303]++
							buf[i] = 0
//line /usr/local/go/src/math/big/int.go:510
		// _ = "end of CoverTab[5303]"
	}
//line /usr/local/go/src/math/big/int.go:511
	// _ = "end of CoverTab[5301]"
//line /usr/local/go/src/math/big/int.go:511
	_go_fuzz_dep_.CoverTab[5302]++
						x.abs.bytes(buf)
						return buf
//line /usr/local/go/src/math/big/int.go:513
	// _ = "end of CoverTab[5302]"
}

// BitLen returns the length of the absolute value of x in bits.
//line /usr/local/go/src/math/big/int.go:516
// The bit length of 0 is 0.
//line /usr/local/go/src/math/big/int.go:518
func (x *Int) BitLen() int {
//line /usr/local/go/src/math/big/int.go:518
	_go_fuzz_dep_.CoverTab[5304]++

//line /usr/local/go/src/math/big/int.go:522
	return x.abs.bitLen()
//line /usr/local/go/src/math/big/int.go:522
	// _ = "end of CoverTab[5304]"
}

// TrailingZeroBits returns the number of consecutive least significant zero
//line /usr/local/go/src/math/big/int.go:525
// bits of |x|.
//line /usr/local/go/src/math/big/int.go:527
func (x *Int) TrailingZeroBits() uint {
//line /usr/local/go/src/math/big/int.go:527
	_go_fuzz_dep_.CoverTab[5305]++
						return x.abs.trailingZeroBits()
//line /usr/local/go/src/math/big/int.go:528
	// _ = "end of CoverTab[5305]"
}

// Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
//line /usr/local/go/src/math/big/int.go:531
// If m == nil or m == 0, z = x**y unless y <= 0 then z = 1. If m != 0, y < 0,
//line /usr/local/go/src/math/big/int.go:531
// and x and m are not relatively prime, z is unchanged and nil is returned.
//line /usr/local/go/src/math/big/int.go:531
//
//line /usr/local/go/src/math/big/int.go:531
// Modular exponentiation of inputs of a particular size is not a
//line /usr/local/go/src/math/big/int.go:531
// cryptographically constant-time operation.
//line /usr/local/go/src/math/big/int.go:537
func (z *Int) Exp(x, y, m *Int) *Int {
//line /usr/local/go/src/math/big/int.go:537
	_go_fuzz_dep_.CoverTab[5306]++
						return z.exp(x, y, m, false)
//line /usr/local/go/src/math/big/int.go:538
	// _ = "end of CoverTab[5306]"
}

func (z *Int) expSlow(x, y, m *Int) *Int {
//line /usr/local/go/src/math/big/int.go:541
	_go_fuzz_dep_.CoverTab[5307]++
						return z.exp(x, y, m, true)
//line /usr/local/go/src/math/big/int.go:542
	// _ = "end of CoverTab[5307]"
}

func (z *Int) exp(x, y, m *Int, slow bool) *Int {
//line /usr/local/go/src/math/big/int.go:545
	_go_fuzz_dep_.CoverTab[5308]++

						xWords := x.abs
						if y.neg {
//line /usr/local/go/src/math/big/int.go:548
		_go_fuzz_dep_.CoverTab[5312]++
							if m == nil || func() bool {
//line /usr/local/go/src/math/big/int.go:549
			_go_fuzz_dep_.CoverTab[5315]++
//line /usr/local/go/src/math/big/int.go:549
			return len(m.abs) == 0
//line /usr/local/go/src/math/big/int.go:549
			// _ = "end of CoverTab[5315]"
//line /usr/local/go/src/math/big/int.go:549
		}() {
//line /usr/local/go/src/math/big/int.go:549
			_go_fuzz_dep_.CoverTab[5316]++
								return z.SetInt64(1)
//line /usr/local/go/src/math/big/int.go:550
			// _ = "end of CoverTab[5316]"
		} else {
//line /usr/local/go/src/math/big/int.go:551
			_go_fuzz_dep_.CoverTab[5317]++
//line /usr/local/go/src/math/big/int.go:551
			// _ = "end of CoverTab[5317]"
//line /usr/local/go/src/math/big/int.go:551
		}
//line /usr/local/go/src/math/big/int.go:551
		// _ = "end of CoverTab[5312]"
//line /usr/local/go/src/math/big/int.go:551
		_go_fuzz_dep_.CoverTab[5313]++

							inverse := new(Int).ModInverse(x, m)
							if inverse == nil {
//line /usr/local/go/src/math/big/int.go:554
			_go_fuzz_dep_.CoverTab[5318]++
								return nil
//line /usr/local/go/src/math/big/int.go:555
			// _ = "end of CoverTab[5318]"
		} else {
//line /usr/local/go/src/math/big/int.go:556
			_go_fuzz_dep_.CoverTab[5319]++
//line /usr/local/go/src/math/big/int.go:556
			// _ = "end of CoverTab[5319]"
//line /usr/local/go/src/math/big/int.go:556
		}
//line /usr/local/go/src/math/big/int.go:556
		// _ = "end of CoverTab[5313]"
//line /usr/local/go/src/math/big/int.go:556
		_go_fuzz_dep_.CoverTab[5314]++
							xWords = inverse.abs
//line /usr/local/go/src/math/big/int.go:557
		// _ = "end of CoverTab[5314]"
	} else {
//line /usr/local/go/src/math/big/int.go:558
		_go_fuzz_dep_.CoverTab[5320]++
//line /usr/local/go/src/math/big/int.go:558
		// _ = "end of CoverTab[5320]"
//line /usr/local/go/src/math/big/int.go:558
	}
//line /usr/local/go/src/math/big/int.go:558
	// _ = "end of CoverTab[5308]"
//line /usr/local/go/src/math/big/int.go:558
	_go_fuzz_dep_.CoverTab[5309]++
						yWords := y.abs

						var mWords nat
						if m != nil {
//line /usr/local/go/src/math/big/int.go:562
		_go_fuzz_dep_.CoverTab[5321]++
							if z == m || func() bool {
//line /usr/local/go/src/math/big/int.go:563
			_go_fuzz_dep_.CoverTab[5323]++
//line /usr/local/go/src/math/big/int.go:563
			return alias(z.abs, m.abs)
//line /usr/local/go/src/math/big/int.go:563
			// _ = "end of CoverTab[5323]"
//line /usr/local/go/src/math/big/int.go:563
		}() {
//line /usr/local/go/src/math/big/int.go:563
			_go_fuzz_dep_.CoverTab[5324]++
								m = new(Int).Set(m)
//line /usr/local/go/src/math/big/int.go:564
			// _ = "end of CoverTab[5324]"
		} else {
//line /usr/local/go/src/math/big/int.go:565
			_go_fuzz_dep_.CoverTab[5325]++
//line /usr/local/go/src/math/big/int.go:565
			// _ = "end of CoverTab[5325]"
//line /usr/local/go/src/math/big/int.go:565
		}
//line /usr/local/go/src/math/big/int.go:565
		// _ = "end of CoverTab[5321]"
//line /usr/local/go/src/math/big/int.go:565
		_go_fuzz_dep_.CoverTab[5322]++
							mWords = m.abs
//line /usr/local/go/src/math/big/int.go:566
		// _ = "end of CoverTab[5322]"
	} else {
//line /usr/local/go/src/math/big/int.go:567
		_go_fuzz_dep_.CoverTab[5326]++
//line /usr/local/go/src/math/big/int.go:567
		// _ = "end of CoverTab[5326]"
//line /usr/local/go/src/math/big/int.go:567
	}
//line /usr/local/go/src/math/big/int.go:567
	// _ = "end of CoverTab[5309]"
//line /usr/local/go/src/math/big/int.go:567
	_go_fuzz_dep_.CoverTab[5310]++

						z.abs = z.abs.expNN(xWords, yWords, mWords, slow)
						z.neg = len(z.abs) > 0 && func() bool {
//line /usr/local/go/src/math/big/int.go:570
		_go_fuzz_dep_.CoverTab[5327]++
//line /usr/local/go/src/math/big/int.go:570
		return x.neg
//line /usr/local/go/src/math/big/int.go:570
		// _ = "end of CoverTab[5327]"
//line /usr/local/go/src/math/big/int.go:570
	}() && func() bool {
//line /usr/local/go/src/math/big/int.go:570
		_go_fuzz_dep_.CoverTab[5328]++
//line /usr/local/go/src/math/big/int.go:570
		return len(yWords) > 0
//line /usr/local/go/src/math/big/int.go:570
		// _ = "end of CoverTab[5328]"
//line /usr/local/go/src/math/big/int.go:570
	}() && func() bool {
//line /usr/local/go/src/math/big/int.go:570
		_go_fuzz_dep_.CoverTab[5329]++
//line /usr/local/go/src/math/big/int.go:570
		return yWords[0]&1 == 1
//line /usr/local/go/src/math/big/int.go:570
		// _ = "end of CoverTab[5329]"
//line /usr/local/go/src/math/big/int.go:570
	}()
						if z.neg && func() bool {
//line /usr/local/go/src/math/big/int.go:571
		_go_fuzz_dep_.CoverTab[5330]++
//line /usr/local/go/src/math/big/int.go:571
		return len(mWords) > 0
//line /usr/local/go/src/math/big/int.go:571
		// _ = "end of CoverTab[5330]"
//line /usr/local/go/src/math/big/int.go:571
	}() {
//line /usr/local/go/src/math/big/int.go:571
		_go_fuzz_dep_.CoverTab[5331]++

							z.abs = z.abs.sub(mWords, z.abs)
							z.neg = false
//line /usr/local/go/src/math/big/int.go:574
		// _ = "end of CoverTab[5331]"
	} else {
//line /usr/local/go/src/math/big/int.go:575
		_go_fuzz_dep_.CoverTab[5332]++
//line /usr/local/go/src/math/big/int.go:575
		// _ = "end of CoverTab[5332]"
//line /usr/local/go/src/math/big/int.go:575
	}
//line /usr/local/go/src/math/big/int.go:575
	// _ = "end of CoverTab[5310]"
//line /usr/local/go/src/math/big/int.go:575
	_go_fuzz_dep_.CoverTab[5311]++

						return z
//line /usr/local/go/src/math/big/int.go:577
	// _ = "end of CoverTab[5311]"
}

// GCD sets z to the greatest common divisor of a and b and returns z.
//line /usr/local/go/src/math/big/int.go:580
// If x or y are not nil, GCD sets their value such that z = a*x + b*y.
//line /usr/local/go/src/math/big/int.go:580
//
//line /usr/local/go/src/math/big/int.go:580
// a and b may be positive, zero or negative. (Before Go 1.14 both had
//line /usr/local/go/src/math/big/int.go:580
// to be > 0.) Regardless of the signs of a and b, z is always >= 0.
//line /usr/local/go/src/math/big/int.go:580
//
//line /usr/local/go/src/math/big/int.go:580
// If a == b == 0, GCD sets z = x = y = 0.
//line /usr/local/go/src/math/big/int.go:580
//
//line /usr/local/go/src/math/big/int.go:580
// If a == 0 and b != 0, GCD sets z = |b|, x = 0, y = sign(b) * 1.
//line /usr/local/go/src/math/big/int.go:580
//
//line /usr/local/go/src/math/big/int.go:580
// If a != 0 and b == 0, GCD sets z = |a|, x = sign(a) * 1, y = 0.
//line /usr/local/go/src/math/big/int.go:591
func (z *Int) GCD(x, y, a, b *Int) *Int {
//line /usr/local/go/src/math/big/int.go:591
	_go_fuzz_dep_.CoverTab[5333]++
						if len(a.abs) == 0 || func() bool {
//line /usr/local/go/src/math/big/int.go:592
		_go_fuzz_dep_.CoverTab[5335]++
//line /usr/local/go/src/math/big/int.go:592
		return len(b.abs) == 0
//line /usr/local/go/src/math/big/int.go:592
		// _ = "end of CoverTab[5335]"
//line /usr/local/go/src/math/big/int.go:592
	}() {
//line /usr/local/go/src/math/big/int.go:592
		_go_fuzz_dep_.CoverTab[5336]++
							lenA, lenB, negA, negB := len(a.abs), len(b.abs), a.neg, b.neg
							if lenA == 0 {
//line /usr/local/go/src/math/big/int.go:594
			_go_fuzz_dep_.CoverTab[5340]++
								z.Set(b)
//line /usr/local/go/src/math/big/int.go:595
			// _ = "end of CoverTab[5340]"
		} else {
//line /usr/local/go/src/math/big/int.go:596
			_go_fuzz_dep_.CoverTab[5341]++
								z.Set(a)
//line /usr/local/go/src/math/big/int.go:597
			// _ = "end of CoverTab[5341]"
		}
//line /usr/local/go/src/math/big/int.go:598
		// _ = "end of CoverTab[5336]"
//line /usr/local/go/src/math/big/int.go:598
		_go_fuzz_dep_.CoverTab[5337]++
							z.neg = false
							if x != nil {
//line /usr/local/go/src/math/big/int.go:600
			_go_fuzz_dep_.CoverTab[5342]++
								if lenA == 0 {
//line /usr/local/go/src/math/big/int.go:601
				_go_fuzz_dep_.CoverTab[5343]++
									x.SetUint64(0)
//line /usr/local/go/src/math/big/int.go:602
				// _ = "end of CoverTab[5343]"
			} else {
//line /usr/local/go/src/math/big/int.go:603
				_go_fuzz_dep_.CoverTab[5344]++
									x.SetUint64(1)
									x.neg = negA
//line /usr/local/go/src/math/big/int.go:605
				// _ = "end of CoverTab[5344]"
			}
//line /usr/local/go/src/math/big/int.go:606
			// _ = "end of CoverTab[5342]"
		} else {
//line /usr/local/go/src/math/big/int.go:607
			_go_fuzz_dep_.CoverTab[5345]++
//line /usr/local/go/src/math/big/int.go:607
			// _ = "end of CoverTab[5345]"
//line /usr/local/go/src/math/big/int.go:607
		}
//line /usr/local/go/src/math/big/int.go:607
		// _ = "end of CoverTab[5337]"
//line /usr/local/go/src/math/big/int.go:607
		_go_fuzz_dep_.CoverTab[5338]++
							if y != nil {
//line /usr/local/go/src/math/big/int.go:608
			_go_fuzz_dep_.CoverTab[5346]++
								if lenB == 0 {
//line /usr/local/go/src/math/big/int.go:609
				_go_fuzz_dep_.CoverTab[5347]++
									y.SetUint64(0)
//line /usr/local/go/src/math/big/int.go:610
				// _ = "end of CoverTab[5347]"
			} else {
//line /usr/local/go/src/math/big/int.go:611
				_go_fuzz_dep_.CoverTab[5348]++
									y.SetUint64(1)
									y.neg = negB
//line /usr/local/go/src/math/big/int.go:613
				// _ = "end of CoverTab[5348]"
			}
//line /usr/local/go/src/math/big/int.go:614
			// _ = "end of CoverTab[5346]"
		} else {
//line /usr/local/go/src/math/big/int.go:615
			_go_fuzz_dep_.CoverTab[5349]++
//line /usr/local/go/src/math/big/int.go:615
			// _ = "end of CoverTab[5349]"
//line /usr/local/go/src/math/big/int.go:615
		}
//line /usr/local/go/src/math/big/int.go:615
		// _ = "end of CoverTab[5338]"
//line /usr/local/go/src/math/big/int.go:615
		_go_fuzz_dep_.CoverTab[5339]++
							return z
//line /usr/local/go/src/math/big/int.go:616
		// _ = "end of CoverTab[5339]"
	} else {
//line /usr/local/go/src/math/big/int.go:617
		_go_fuzz_dep_.CoverTab[5350]++
//line /usr/local/go/src/math/big/int.go:617
		// _ = "end of CoverTab[5350]"
//line /usr/local/go/src/math/big/int.go:617
	}
//line /usr/local/go/src/math/big/int.go:617
	// _ = "end of CoverTab[5333]"
//line /usr/local/go/src/math/big/int.go:617
	_go_fuzz_dep_.CoverTab[5334]++

						return z.lehmerGCD(x, y, a, b)
//line /usr/local/go/src/math/big/int.go:619
	// _ = "end of CoverTab[5334]"
}

// lehmerSimulate attempts to simulate several Euclidean update steps
//line /usr/local/go/src/math/big/int.go:622
// using the leading digits of A and B.  It returns u0, u1, v0, v1
//line /usr/local/go/src/math/big/int.go:622
// such that A and B can be updated as:
//line /usr/local/go/src/math/big/int.go:622
//
//line /usr/local/go/src/math/big/int.go:622
//	A = u0*A + v0*B
//line /usr/local/go/src/math/big/int.go:622
//	B = u1*A + v1*B
//line /usr/local/go/src/math/big/int.go:622
//
//line /usr/local/go/src/math/big/int.go:622
// Requirements: A >= B and len(B.abs) >= 2
//line /usr/local/go/src/math/big/int.go:622
// Since we are calculating with full words to avoid overflow,
//line /usr/local/go/src/math/big/int.go:622
// we use 'even' to track the sign of the cosequences.
//line /usr/local/go/src/math/big/int.go:622
// For even iterations: u0, v1 >= 0 && u1, v0 <= 0
//line /usr/local/go/src/math/big/int.go:622
// For odd  iterations: u0, v1 <= 0 && u1, v0 >= 0
//line /usr/local/go/src/math/big/int.go:634
func lehmerSimulate(A, B *Int) (u0, u1, v0, v1 Word, even bool) {
//line /usr/local/go/src/math/big/int.go:634
	_go_fuzz_dep_.CoverTab[5351]++
						// initialize the digits
						var a1, a2, u2, v2 Word

						m := len(B.abs)
						n := len(A.abs)

//line /usr/local/go/src/math/big/int.go:642
	h := nlz(A.abs[n-1])
	a1 = A.abs[n-1]<<h | A.abs[n-2]>>(_W-h)

	switch {
	case n == m:
//line /usr/local/go/src/math/big/int.go:646
		_go_fuzz_dep_.CoverTab[5354]++
							a2 = B.abs[n-1]<<h | B.abs[n-2]>>(_W-h)
//line /usr/local/go/src/math/big/int.go:647
		// _ = "end of CoverTab[5354]"
	case n == m+1:
//line /usr/local/go/src/math/big/int.go:648
		_go_fuzz_dep_.CoverTab[5355]++
							a2 = B.abs[n-2] >> (_W - h)
//line /usr/local/go/src/math/big/int.go:649
		// _ = "end of CoverTab[5355]"
	default:
//line /usr/local/go/src/math/big/int.go:650
		_go_fuzz_dep_.CoverTab[5356]++
							a2 = 0
//line /usr/local/go/src/math/big/int.go:651
		// _ = "end of CoverTab[5356]"
	}
//line /usr/local/go/src/math/big/int.go:652
	// _ = "end of CoverTab[5351]"
//line /usr/local/go/src/math/big/int.go:652
	_go_fuzz_dep_.CoverTab[5352]++

//line /usr/local/go/src/math/big/int.go:659
	even = false

						u0, u1, u2 = 0, 1, 0
						v0, v1, v2 = 0, 0, 1

//line /usr/local/go/src/math/big/int.go:668
	for a2 >= v2 && func() bool {
//line /usr/local/go/src/math/big/int.go:668
		_go_fuzz_dep_.CoverTab[5357]++
//line /usr/local/go/src/math/big/int.go:668
		return a1-a2 >= v1+v2
//line /usr/local/go/src/math/big/int.go:668
		// _ = "end of CoverTab[5357]"
//line /usr/local/go/src/math/big/int.go:668
	}() {
//line /usr/local/go/src/math/big/int.go:668
		_go_fuzz_dep_.CoverTab[5358]++
							q, r := a1/a2, a1%a2
							a1, a2 = a2, r
							u0, u1, u2 = u1, u2, u1+q*u2
							v0, v1, v2 = v1, v2, v1+q*v2
							even = !even
//line /usr/local/go/src/math/big/int.go:673
		// _ = "end of CoverTab[5358]"
	}
//line /usr/local/go/src/math/big/int.go:674
	// _ = "end of CoverTab[5352]"
//line /usr/local/go/src/math/big/int.go:674
	_go_fuzz_dep_.CoverTab[5353]++
						return
//line /usr/local/go/src/math/big/int.go:675
	// _ = "end of CoverTab[5353]"
}

// lehmerUpdate updates the inputs A and B such that:
//line /usr/local/go/src/math/big/int.go:678
//
//line /usr/local/go/src/math/big/int.go:678
//	A = u0*A + v0*B
//line /usr/local/go/src/math/big/int.go:678
//	B = u1*A + v1*B
//line /usr/local/go/src/math/big/int.go:678
//
//line /usr/local/go/src/math/big/int.go:678
// where the signs of u0, u1, v0, v1 are given by even
//line /usr/local/go/src/math/big/int.go:678
// For even == true: u0, v1 >= 0 && u1, v0 <= 0
//line /usr/local/go/src/math/big/int.go:678
// For even == false: u0, v1 <= 0 && u1, v0 >= 0
//line /usr/local/go/src/math/big/int.go:678
// q, r, s, t are temporary variables to avoid allocations in the multiplication.
//line /usr/local/go/src/math/big/int.go:687
func lehmerUpdate(A, B, q, r, s, t *Int, u0, u1, v0, v1 Word, even bool) {
//line /usr/local/go/src/math/big/int.go:687
	_go_fuzz_dep_.CoverTab[5359]++

						t.abs = t.abs.setWord(u0)
						s.abs = s.abs.setWord(v0)
						t.neg = !even
						s.neg = even

						t.Mul(A, t)
						s.Mul(B, s)

						r.abs = r.abs.setWord(u1)
						q.abs = q.abs.setWord(v1)
						r.neg = even
						q.neg = !even

						r.Mul(A, r)
						q.Mul(B, q)

						A.Add(t, s)
						B.Add(r, q)
//line /usr/local/go/src/math/big/int.go:706
	// _ = "end of CoverTab[5359]"
}

// euclidUpdate performs a single step of the Euclidean GCD algorithm
//line /usr/local/go/src/math/big/int.go:709
// if extended is true, it also updates the cosequence Ua, Ub.
//line /usr/local/go/src/math/big/int.go:711
func euclidUpdate(A, B, Ua, Ub, q, r, s, t *Int, extended bool) {
//line /usr/local/go/src/math/big/int.go:711
	_go_fuzz_dep_.CoverTab[5360]++
						q, r = q.QuoRem(A, B, r)

						*A, *B, *r = *B, *r, *A

						if extended {
//line /usr/local/go/src/math/big/int.go:716
		_go_fuzz_dep_.CoverTab[5361]++

							t.Set(Ub)
							s.Mul(Ub, q)
							Ub.Sub(Ua, s)
							Ua.Set(t)
//line /usr/local/go/src/math/big/int.go:721
		// _ = "end of CoverTab[5361]"
	} else {
//line /usr/local/go/src/math/big/int.go:722
		_go_fuzz_dep_.CoverTab[5362]++
//line /usr/local/go/src/math/big/int.go:722
		// _ = "end of CoverTab[5362]"
//line /usr/local/go/src/math/big/int.go:722
	}
//line /usr/local/go/src/math/big/int.go:722
	// _ = "end of CoverTab[5360]"
}

// lehmerGCD sets z to the greatest common divisor of a and b,
//line /usr/local/go/src/math/big/int.go:725
// which both must be != 0, and returns z.
//line /usr/local/go/src/math/big/int.go:725
// If x or y are not nil, their values are set such that z = a*x + b*y.
//line /usr/local/go/src/math/big/int.go:725
// See Knuth, The Art of Computer Programming, Vol. 2, Section 4.5.2, Algorithm L.
//line /usr/local/go/src/math/big/int.go:725
// This implementation uses the improved condition by Collins requiring only one
//line /usr/local/go/src/math/big/int.go:725
// quotient and avoiding the possibility of single Word overflow.
//line /usr/local/go/src/math/big/int.go:725
// See Jebelean, "Improving the multiprecision Euclidean algorithm",
//line /usr/local/go/src/math/big/int.go:725
// Design and Implementation of Symbolic Computation Systems, pp 45-58.
//line /usr/local/go/src/math/big/int.go:725
// The cosequences are updated according to Algorithm 10.45 from
//line /usr/local/go/src/math/big/int.go:725
// Cohen et al. "Handbook of Elliptic and Hyperelliptic Curve Cryptography" pp 192.
//line /usr/local/go/src/math/big/int.go:735
func (z *Int) lehmerGCD(x, y, a, b *Int) *Int {
//line /usr/local/go/src/math/big/int.go:735
	_go_fuzz_dep_.CoverTab[5363]++
						var A, B, Ua, Ub *Int

						A = new(Int).Abs(a)
						B = new(Int).Abs(b)

						extended := x != nil || func() bool {
//line /usr/local/go/src/math/big/int.go:741
		_go_fuzz_dep_.CoverTab[5370]++
//line /usr/local/go/src/math/big/int.go:741
		return y != nil
//line /usr/local/go/src/math/big/int.go:741
		// _ = "end of CoverTab[5370]"
//line /usr/local/go/src/math/big/int.go:741
	}()

						if extended {
//line /usr/local/go/src/math/big/int.go:743
		_go_fuzz_dep_.CoverTab[5371]++

							Ua = new(Int).SetInt64(1)
							Ub = new(Int)
//line /usr/local/go/src/math/big/int.go:746
		// _ = "end of CoverTab[5371]"
	} else {
//line /usr/local/go/src/math/big/int.go:747
		_go_fuzz_dep_.CoverTab[5372]++
//line /usr/local/go/src/math/big/int.go:747
		// _ = "end of CoverTab[5372]"
//line /usr/local/go/src/math/big/int.go:747
	}
//line /usr/local/go/src/math/big/int.go:747
	// _ = "end of CoverTab[5363]"
//line /usr/local/go/src/math/big/int.go:747
	_go_fuzz_dep_.CoverTab[5364]++

//line /usr/local/go/src/math/big/int.go:750
	q := new(Int)
						r := new(Int)
						s := new(Int)
						t := new(Int)

//line /usr/local/go/src/math/big/int.go:756
	if A.abs.cmp(B.abs) < 0 {
//line /usr/local/go/src/math/big/int.go:756
		_go_fuzz_dep_.CoverTab[5373]++
							A, B = B, A
							Ub, Ua = Ua, Ub
//line /usr/local/go/src/math/big/int.go:758
		// _ = "end of CoverTab[5373]"
	} else {
//line /usr/local/go/src/math/big/int.go:759
		_go_fuzz_dep_.CoverTab[5374]++
//line /usr/local/go/src/math/big/int.go:759
		// _ = "end of CoverTab[5374]"
//line /usr/local/go/src/math/big/int.go:759
	}
//line /usr/local/go/src/math/big/int.go:759
	// _ = "end of CoverTab[5364]"
//line /usr/local/go/src/math/big/int.go:759
	_go_fuzz_dep_.CoverTab[5365]++

//line /usr/local/go/src/math/big/int.go:762
	for len(B.abs) > 1 {
//line /usr/local/go/src/math/big/int.go:762
		_go_fuzz_dep_.CoverTab[5375]++

							u0, u1, v0, v1, even := lehmerSimulate(A, B)

//line /usr/local/go/src/math/big/int.go:767
		if v0 != 0 {
//line /usr/local/go/src/math/big/int.go:767
			_go_fuzz_dep_.CoverTab[5376]++

//line /usr/local/go/src/math/big/int.go:771
			lehmerUpdate(A, B, q, r, s, t, u0, u1, v0, v1, even)

			if extended {
//line /usr/local/go/src/math/big/int.go:773
				_go_fuzz_dep_.CoverTab[5377]++

//line /usr/local/go/src/math/big/int.go:776
				lehmerUpdate(Ua, Ub, q, r, s, t, u0, u1, v0, v1, even)
//line /usr/local/go/src/math/big/int.go:776
				// _ = "end of CoverTab[5377]"
			} else {
//line /usr/local/go/src/math/big/int.go:777
				_go_fuzz_dep_.CoverTab[5378]++
//line /usr/local/go/src/math/big/int.go:777
				// _ = "end of CoverTab[5378]"
//line /usr/local/go/src/math/big/int.go:777
			}
//line /usr/local/go/src/math/big/int.go:777
			// _ = "end of CoverTab[5376]"

		} else {
//line /usr/local/go/src/math/big/int.go:779
			_go_fuzz_dep_.CoverTab[5379]++

//line /usr/local/go/src/math/big/int.go:782
			euclidUpdate(A, B, Ua, Ub, q, r, s, t, extended)
//line /usr/local/go/src/math/big/int.go:782
			// _ = "end of CoverTab[5379]"
		}
//line /usr/local/go/src/math/big/int.go:783
		// _ = "end of CoverTab[5375]"
	}
//line /usr/local/go/src/math/big/int.go:784
	// _ = "end of CoverTab[5365]"
//line /usr/local/go/src/math/big/int.go:784
	_go_fuzz_dep_.CoverTab[5366]++

						if len(B.abs) > 0 {
//line /usr/local/go/src/math/big/int.go:786
		_go_fuzz_dep_.CoverTab[5380]++

							if len(A.abs) > 1 {
//line /usr/local/go/src/math/big/int.go:788
			_go_fuzz_dep_.CoverTab[5382]++

								euclidUpdate(A, B, Ua, Ub, q, r, s, t, extended)
//line /usr/local/go/src/math/big/int.go:790
			// _ = "end of CoverTab[5382]"
		} else {
//line /usr/local/go/src/math/big/int.go:791
			_go_fuzz_dep_.CoverTab[5383]++
//line /usr/local/go/src/math/big/int.go:791
			// _ = "end of CoverTab[5383]"
//line /usr/local/go/src/math/big/int.go:791
		}
//line /usr/local/go/src/math/big/int.go:791
		// _ = "end of CoverTab[5380]"
//line /usr/local/go/src/math/big/int.go:791
		_go_fuzz_dep_.CoverTab[5381]++
							if len(B.abs) > 0 {
//line /usr/local/go/src/math/big/int.go:792
			_go_fuzz_dep_.CoverTab[5384]++

								aWord, bWord := A.abs[0], B.abs[0]
								if extended {
//line /usr/local/go/src/math/big/int.go:795
				_go_fuzz_dep_.CoverTab[5386]++
									var ua, ub, va, vb Word
									ua, ub = 1, 0
									va, vb = 0, 1
									even := true
									for bWord != 0 {
//line /usr/local/go/src/math/big/int.go:800
					_go_fuzz_dep_.CoverTab[5388]++
										q, r := aWord/bWord, aWord%bWord
										aWord, bWord = bWord, r
										ua, ub = ub, ua+q*ub
										va, vb = vb, va+q*vb
										even = !even
//line /usr/local/go/src/math/big/int.go:805
					// _ = "end of CoverTab[5388]"
				}
//line /usr/local/go/src/math/big/int.go:806
				// _ = "end of CoverTab[5386]"
//line /usr/local/go/src/math/big/int.go:806
				_go_fuzz_dep_.CoverTab[5387]++

									t.abs = t.abs.setWord(ua)
									s.abs = s.abs.setWord(va)
									t.neg = !even
									s.neg = even

									t.Mul(Ua, t)
									s.Mul(Ub, s)

									Ua.Add(t, s)
//line /usr/local/go/src/math/big/int.go:816
				// _ = "end of CoverTab[5387]"
			} else {
//line /usr/local/go/src/math/big/int.go:817
				_go_fuzz_dep_.CoverTab[5389]++
									for bWord != 0 {
//line /usr/local/go/src/math/big/int.go:818
					_go_fuzz_dep_.CoverTab[5390]++
										aWord, bWord = bWord, aWord%bWord
//line /usr/local/go/src/math/big/int.go:819
					// _ = "end of CoverTab[5390]"
				}
//line /usr/local/go/src/math/big/int.go:820
				// _ = "end of CoverTab[5389]"
			}
//line /usr/local/go/src/math/big/int.go:821
			// _ = "end of CoverTab[5384]"
//line /usr/local/go/src/math/big/int.go:821
			_go_fuzz_dep_.CoverTab[5385]++
								A.abs[0] = aWord
//line /usr/local/go/src/math/big/int.go:822
			// _ = "end of CoverTab[5385]"
		} else {
//line /usr/local/go/src/math/big/int.go:823
			_go_fuzz_dep_.CoverTab[5391]++
//line /usr/local/go/src/math/big/int.go:823
			// _ = "end of CoverTab[5391]"
//line /usr/local/go/src/math/big/int.go:823
		}
//line /usr/local/go/src/math/big/int.go:823
		// _ = "end of CoverTab[5381]"
	} else {
//line /usr/local/go/src/math/big/int.go:824
		_go_fuzz_dep_.CoverTab[5392]++
//line /usr/local/go/src/math/big/int.go:824
		// _ = "end of CoverTab[5392]"
//line /usr/local/go/src/math/big/int.go:824
	}
//line /usr/local/go/src/math/big/int.go:824
	// _ = "end of CoverTab[5366]"
//line /usr/local/go/src/math/big/int.go:824
	_go_fuzz_dep_.CoverTab[5367]++
						negA := a.neg
						if y != nil {
//line /usr/local/go/src/math/big/int.go:826
		_go_fuzz_dep_.CoverTab[5393]++

							if y == b {
//line /usr/local/go/src/math/big/int.go:828
			_go_fuzz_dep_.CoverTab[5396]++
								B.Set(b)
//line /usr/local/go/src/math/big/int.go:829
			// _ = "end of CoverTab[5396]"
		} else {
//line /usr/local/go/src/math/big/int.go:830
			_go_fuzz_dep_.CoverTab[5397]++
								B = b
//line /usr/local/go/src/math/big/int.go:831
			// _ = "end of CoverTab[5397]"
		}
//line /usr/local/go/src/math/big/int.go:832
		// _ = "end of CoverTab[5393]"
//line /usr/local/go/src/math/big/int.go:832
		_go_fuzz_dep_.CoverTab[5394]++

							y.Mul(a, Ua)
							if negA {
//line /usr/local/go/src/math/big/int.go:835
			_go_fuzz_dep_.CoverTab[5398]++
								y.neg = !y.neg
//line /usr/local/go/src/math/big/int.go:836
			// _ = "end of CoverTab[5398]"
		} else {
//line /usr/local/go/src/math/big/int.go:837
			_go_fuzz_dep_.CoverTab[5399]++
//line /usr/local/go/src/math/big/int.go:837
			// _ = "end of CoverTab[5399]"
//line /usr/local/go/src/math/big/int.go:837
		}
//line /usr/local/go/src/math/big/int.go:837
		// _ = "end of CoverTab[5394]"
//line /usr/local/go/src/math/big/int.go:837
		_go_fuzz_dep_.CoverTab[5395]++
							y.Sub(A, y)
							y.Div(y, B)
//line /usr/local/go/src/math/big/int.go:839
		// _ = "end of CoverTab[5395]"
	} else {
//line /usr/local/go/src/math/big/int.go:840
		_go_fuzz_dep_.CoverTab[5400]++
//line /usr/local/go/src/math/big/int.go:840
		// _ = "end of CoverTab[5400]"
//line /usr/local/go/src/math/big/int.go:840
	}
//line /usr/local/go/src/math/big/int.go:840
	// _ = "end of CoverTab[5367]"
//line /usr/local/go/src/math/big/int.go:840
	_go_fuzz_dep_.CoverTab[5368]++

						if x != nil {
//line /usr/local/go/src/math/big/int.go:842
		_go_fuzz_dep_.CoverTab[5401]++
							*x = *Ua
							if negA {
//line /usr/local/go/src/math/big/int.go:844
			_go_fuzz_dep_.CoverTab[5402]++
								x.neg = !x.neg
//line /usr/local/go/src/math/big/int.go:845
			// _ = "end of CoverTab[5402]"
		} else {
//line /usr/local/go/src/math/big/int.go:846
			_go_fuzz_dep_.CoverTab[5403]++
//line /usr/local/go/src/math/big/int.go:846
			// _ = "end of CoverTab[5403]"
//line /usr/local/go/src/math/big/int.go:846
		}
//line /usr/local/go/src/math/big/int.go:846
		// _ = "end of CoverTab[5401]"
	} else {
//line /usr/local/go/src/math/big/int.go:847
		_go_fuzz_dep_.CoverTab[5404]++
//line /usr/local/go/src/math/big/int.go:847
		// _ = "end of CoverTab[5404]"
//line /usr/local/go/src/math/big/int.go:847
	}
//line /usr/local/go/src/math/big/int.go:847
	// _ = "end of CoverTab[5368]"
//line /usr/local/go/src/math/big/int.go:847
	_go_fuzz_dep_.CoverTab[5369]++

						*z = *A

						return z
//line /usr/local/go/src/math/big/int.go:851
	// _ = "end of CoverTab[5369]"
}

// Rand sets z to a pseudo-random number in [0, n) and returns z.
//line /usr/local/go/src/math/big/int.go:854
//
//line /usr/local/go/src/math/big/int.go:854
// As this uses the math/rand package, it must not be used for
//line /usr/local/go/src/math/big/int.go:854
// security-sensitive work. Use crypto/rand.Int instead.
//line /usr/local/go/src/math/big/int.go:858
func (z *Int) Rand(rnd *rand.Rand, n *Int) *Int {
//line /usr/local/go/src/math/big/int.go:858
	_go_fuzz_dep_.CoverTab[5405]++

						if n.neg || func() bool {
//line /usr/local/go/src/math/big/int.go:860
		_go_fuzz_dep_.CoverTab[5407]++
//line /usr/local/go/src/math/big/int.go:860
		return len(n.abs) == 0
//line /usr/local/go/src/math/big/int.go:860
		// _ = "end of CoverTab[5407]"
//line /usr/local/go/src/math/big/int.go:860
	}() {
//line /usr/local/go/src/math/big/int.go:860
		_go_fuzz_dep_.CoverTab[5408]++
							z.neg = false
							z.abs = nil
							return z
//line /usr/local/go/src/math/big/int.go:863
		// _ = "end of CoverTab[5408]"
	} else {
//line /usr/local/go/src/math/big/int.go:864
		_go_fuzz_dep_.CoverTab[5409]++
//line /usr/local/go/src/math/big/int.go:864
		// _ = "end of CoverTab[5409]"
//line /usr/local/go/src/math/big/int.go:864
	}
//line /usr/local/go/src/math/big/int.go:864
	// _ = "end of CoverTab[5405]"
//line /usr/local/go/src/math/big/int.go:864
	_go_fuzz_dep_.CoverTab[5406]++
						z.neg = false
						z.abs = z.abs.random(rnd, n.abs, n.abs.bitLen())
						return z
//line /usr/local/go/src/math/big/int.go:867
	// _ = "end of CoverTab[5406]"
}

// ModInverse sets z to the multiplicative inverse of g in the ring ℤ/nℤ
//line /usr/local/go/src/math/big/int.go:870
// and returns z. If g and n are not relatively prime, g has no multiplicative
//line /usr/local/go/src/math/big/int.go:870
// inverse in the ring ℤ/nℤ.  In this case, z is unchanged and the return value
//line /usr/local/go/src/math/big/int.go:870
// is nil. If n == 0, a division-by-zero run-time panic occurs.
//line /usr/local/go/src/math/big/int.go:874
func (z *Int) ModInverse(g, n *Int) *Int {
//line /usr/local/go/src/math/big/int.go:874
	_go_fuzz_dep_.CoverTab[5410]++

						if n.neg {
//line /usr/local/go/src/math/big/int.go:876
		_go_fuzz_dep_.CoverTab[5415]++
							var n2 Int
							n = n2.Neg(n)
//line /usr/local/go/src/math/big/int.go:878
		// _ = "end of CoverTab[5415]"
	} else {
//line /usr/local/go/src/math/big/int.go:879
		_go_fuzz_dep_.CoverTab[5416]++
//line /usr/local/go/src/math/big/int.go:879
		// _ = "end of CoverTab[5416]"
//line /usr/local/go/src/math/big/int.go:879
	}
//line /usr/local/go/src/math/big/int.go:879
	// _ = "end of CoverTab[5410]"
//line /usr/local/go/src/math/big/int.go:879
	_go_fuzz_dep_.CoverTab[5411]++
						if g.neg {
//line /usr/local/go/src/math/big/int.go:880
		_go_fuzz_dep_.CoverTab[5417]++
							var g2 Int
							g = g2.Mod(g, n)
//line /usr/local/go/src/math/big/int.go:882
		// _ = "end of CoverTab[5417]"
	} else {
//line /usr/local/go/src/math/big/int.go:883
		_go_fuzz_dep_.CoverTab[5418]++
//line /usr/local/go/src/math/big/int.go:883
		// _ = "end of CoverTab[5418]"
//line /usr/local/go/src/math/big/int.go:883
	}
//line /usr/local/go/src/math/big/int.go:883
	// _ = "end of CoverTab[5411]"
//line /usr/local/go/src/math/big/int.go:883
	_go_fuzz_dep_.CoverTab[5412]++
						var d, x Int
						d.GCD(&x, nil, g, n)

//line /usr/local/go/src/math/big/int.go:888
	if d.Cmp(intOne) != 0 {
//line /usr/local/go/src/math/big/int.go:888
		_go_fuzz_dep_.CoverTab[5419]++
							return nil
//line /usr/local/go/src/math/big/int.go:889
		// _ = "end of CoverTab[5419]"
	} else {
//line /usr/local/go/src/math/big/int.go:890
		_go_fuzz_dep_.CoverTab[5420]++
//line /usr/local/go/src/math/big/int.go:890
		// _ = "end of CoverTab[5420]"
//line /usr/local/go/src/math/big/int.go:890
	}
//line /usr/local/go/src/math/big/int.go:890
	// _ = "end of CoverTab[5412]"
//line /usr/local/go/src/math/big/int.go:890
	_go_fuzz_dep_.CoverTab[5413]++

//line /usr/local/go/src/math/big/int.go:894
	if x.neg {
//line /usr/local/go/src/math/big/int.go:894
		_go_fuzz_dep_.CoverTab[5421]++
							z.Add(&x, n)
//line /usr/local/go/src/math/big/int.go:895
		// _ = "end of CoverTab[5421]"
	} else {
//line /usr/local/go/src/math/big/int.go:896
		_go_fuzz_dep_.CoverTab[5422]++
							z.Set(&x)
//line /usr/local/go/src/math/big/int.go:897
		// _ = "end of CoverTab[5422]"
	}
//line /usr/local/go/src/math/big/int.go:898
	// _ = "end of CoverTab[5413]"
//line /usr/local/go/src/math/big/int.go:898
	_go_fuzz_dep_.CoverTab[5414]++
						return z
//line /usr/local/go/src/math/big/int.go:899
	// _ = "end of CoverTab[5414]"
}

func (z nat) modInverse(g, n nat) nat {
//line /usr/local/go/src/math/big/int.go:902
	_go_fuzz_dep_.CoverTab[5423]++

						return (&Int{abs: z}).ModInverse(&Int{abs: g}, &Int{abs: n}).abs
//line /usr/local/go/src/math/big/int.go:904
	// _ = "end of CoverTab[5423]"
}

// Jacobi returns the Jacobi symbol (x/y), either +1, -1, or 0.
//line /usr/local/go/src/math/big/int.go:907
// The y argument must be an odd integer.
//line /usr/local/go/src/math/big/int.go:909
func Jacobi(x, y *Int) int {
//line /usr/local/go/src/math/big/int.go:909
	_go_fuzz_dep_.CoverTab[5424]++
						if len(y.abs) == 0 || func() bool {
//line /usr/local/go/src/math/big/int.go:910
		_go_fuzz_dep_.CoverTab[5427]++
//line /usr/local/go/src/math/big/int.go:910
		return y.abs[0]&1 == 0
//line /usr/local/go/src/math/big/int.go:910
		// _ = "end of CoverTab[5427]"
//line /usr/local/go/src/math/big/int.go:910
	}() {
//line /usr/local/go/src/math/big/int.go:910
		_go_fuzz_dep_.CoverTab[5428]++
							panic(fmt.Sprintf("big: invalid 2nd argument to Int.Jacobi: need odd integer but got %s", y.String()))
//line /usr/local/go/src/math/big/int.go:911
		// _ = "end of CoverTab[5428]"
	} else {
//line /usr/local/go/src/math/big/int.go:912
		_go_fuzz_dep_.CoverTab[5429]++
//line /usr/local/go/src/math/big/int.go:912
		// _ = "end of CoverTab[5429]"
//line /usr/local/go/src/math/big/int.go:912
	}
//line /usr/local/go/src/math/big/int.go:912
	// _ = "end of CoverTab[5424]"
//line /usr/local/go/src/math/big/int.go:912
	_go_fuzz_dep_.CoverTab[5425]++

//line /usr/local/go/src/math/big/int.go:918
	var a, b, c Int
	a.Set(x)
	b.Set(y)
	j := 1

	if b.neg {
//line /usr/local/go/src/math/big/int.go:923
		_go_fuzz_dep_.CoverTab[5430]++
							if a.neg {
//line /usr/local/go/src/math/big/int.go:924
			_go_fuzz_dep_.CoverTab[5432]++
								j = -1
//line /usr/local/go/src/math/big/int.go:925
			// _ = "end of CoverTab[5432]"
		} else {
//line /usr/local/go/src/math/big/int.go:926
			_go_fuzz_dep_.CoverTab[5433]++
//line /usr/local/go/src/math/big/int.go:926
			// _ = "end of CoverTab[5433]"
//line /usr/local/go/src/math/big/int.go:926
		}
//line /usr/local/go/src/math/big/int.go:926
		// _ = "end of CoverTab[5430]"
//line /usr/local/go/src/math/big/int.go:926
		_go_fuzz_dep_.CoverTab[5431]++
							b.neg = false
//line /usr/local/go/src/math/big/int.go:927
		// _ = "end of CoverTab[5431]"
	} else {
//line /usr/local/go/src/math/big/int.go:928
		_go_fuzz_dep_.CoverTab[5434]++
//line /usr/local/go/src/math/big/int.go:928
		// _ = "end of CoverTab[5434]"
//line /usr/local/go/src/math/big/int.go:928
	}
//line /usr/local/go/src/math/big/int.go:928
	// _ = "end of CoverTab[5425]"
//line /usr/local/go/src/math/big/int.go:928
	_go_fuzz_dep_.CoverTab[5426]++

						for {
//line /usr/local/go/src/math/big/int.go:930
		_go_fuzz_dep_.CoverTab[5435]++
							if b.Cmp(intOne) == 0 {
//line /usr/local/go/src/math/big/int.go:931
			_go_fuzz_dep_.CoverTab[5441]++
								return j
//line /usr/local/go/src/math/big/int.go:932
			// _ = "end of CoverTab[5441]"
		} else {
//line /usr/local/go/src/math/big/int.go:933
			_go_fuzz_dep_.CoverTab[5442]++
//line /usr/local/go/src/math/big/int.go:933
			// _ = "end of CoverTab[5442]"
//line /usr/local/go/src/math/big/int.go:933
		}
//line /usr/local/go/src/math/big/int.go:933
		// _ = "end of CoverTab[5435]"
//line /usr/local/go/src/math/big/int.go:933
		_go_fuzz_dep_.CoverTab[5436]++
							if len(a.abs) == 0 {
//line /usr/local/go/src/math/big/int.go:934
			_go_fuzz_dep_.CoverTab[5443]++
								return 0
//line /usr/local/go/src/math/big/int.go:935
			// _ = "end of CoverTab[5443]"
		} else {
//line /usr/local/go/src/math/big/int.go:936
			_go_fuzz_dep_.CoverTab[5444]++
//line /usr/local/go/src/math/big/int.go:936
			// _ = "end of CoverTab[5444]"
//line /usr/local/go/src/math/big/int.go:936
		}
//line /usr/local/go/src/math/big/int.go:936
		// _ = "end of CoverTab[5436]"
//line /usr/local/go/src/math/big/int.go:936
		_go_fuzz_dep_.CoverTab[5437]++
							a.Mod(&a, &b)
							if len(a.abs) == 0 {
//line /usr/local/go/src/math/big/int.go:938
			_go_fuzz_dep_.CoverTab[5445]++
								return 0
//line /usr/local/go/src/math/big/int.go:939
			// _ = "end of CoverTab[5445]"
		} else {
//line /usr/local/go/src/math/big/int.go:940
			_go_fuzz_dep_.CoverTab[5446]++
//line /usr/local/go/src/math/big/int.go:940
			// _ = "end of CoverTab[5446]"
//line /usr/local/go/src/math/big/int.go:940
		}
//line /usr/local/go/src/math/big/int.go:940
		// _ = "end of CoverTab[5437]"
//line /usr/local/go/src/math/big/int.go:940
		_go_fuzz_dep_.CoverTab[5438]++

//line /usr/local/go/src/math/big/int.go:944
		s := a.abs.trailingZeroBits()
		if s&1 != 0 {
//line /usr/local/go/src/math/big/int.go:945
			_go_fuzz_dep_.CoverTab[5447]++
								bmod8 := b.abs[0] & 7
								if bmod8 == 3 || func() bool {
//line /usr/local/go/src/math/big/int.go:947
				_go_fuzz_dep_.CoverTab[5448]++
//line /usr/local/go/src/math/big/int.go:947
				return bmod8 == 5
//line /usr/local/go/src/math/big/int.go:947
				// _ = "end of CoverTab[5448]"
//line /usr/local/go/src/math/big/int.go:947
			}() {
//line /usr/local/go/src/math/big/int.go:947
				_go_fuzz_dep_.CoverTab[5449]++
									j = -j
//line /usr/local/go/src/math/big/int.go:948
				// _ = "end of CoverTab[5449]"
			} else {
//line /usr/local/go/src/math/big/int.go:949
				_go_fuzz_dep_.CoverTab[5450]++
//line /usr/local/go/src/math/big/int.go:949
				// _ = "end of CoverTab[5450]"
//line /usr/local/go/src/math/big/int.go:949
			}
//line /usr/local/go/src/math/big/int.go:949
			// _ = "end of CoverTab[5447]"
		} else {
//line /usr/local/go/src/math/big/int.go:950
			_go_fuzz_dep_.CoverTab[5451]++
//line /usr/local/go/src/math/big/int.go:950
			// _ = "end of CoverTab[5451]"
//line /usr/local/go/src/math/big/int.go:950
		}
//line /usr/local/go/src/math/big/int.go:950
		// _ = "end of CoverTab[5438]"
//line /usr/local/go/src/math/big/int.go:950
		_go_fuzz_dep_.CoverTab[5439]++
							c.Rsh(&a, s)

//line /usr/local/go/src/math/big/int.go:954
		if b.abs[0]&3 == 3 && func() bool {
//line /usr/local/go/src/math/big/int.go:954
			_go_fuzz_dep_.CoverTab[5452]++
//line /usr/local/go/src/math/big/int.go:954
			return c.abs[0]&3 == 3
//line /usr/local/go/src/math/big/int.go:954
			// _ = "end of CoverTab[5452]"
//line /usr/local/go/src/math/big/int.go:954
		}() {
//line /usr/local/go/src/math/big/int.go:954
			_go_fuzz_dep_.CoverTab[5453]++
								j = -j
//line /usr/local/go/src/math/big/int.go:955
			// _ = "end of CoverTab[5453]"
		} else {
//line /usr/local/go/src/math/big/int.go:956
			_go_fuzz_dep_.CoverTab[5454]++
//line /usr/local/go/src/math/big/int.go:956
			// _ = "end of CoverTab[5454]"
//line /usr/local/go/src/math/big/int.go:956
		}
//line /usr/local/go/src/math/big/int.go:956
		// _ = "end of CoverTab[5439]"
//line /usr/local/go/src/math/big/int.go:956
		_go_fuzz_dep_.CoverTab[5440]++
							a.Set(&b)
							b.Set(&c)
//line /usr/local/go/src/math/big/int.go:958
		// _ = "end of CoverTab[5440]"
	}
//line /usr/local/go/src/math/big/int.go:959
	// _ = "end of CoverTab[5426]"
}

// modSqrt3Mod4 uses the identity
//line /usr/local/go/src/math/big/int.go:962
//
//line /usr/local/go/src/math/big/int.go:962
//	   (a^((p+1)/4))^2  mod p
//line /usr/local/go/src/math/big/int.go:962
//	== u^(p+1)          mod p
//line /usr/local/go/src/math/big/int.go:962
//	== u^2              mod p
//line /usr/local/go/src/math/big/int.go:962
//
//line /usr/local/go/src/math/big/int.go:962
// to calculate the square root of any quadratic residue mod p quickly for 3
//line /usr/local/go/src/math/big/int.go:962
// mod 4 primes.
//line /usr/local/go/src/math/big/int.go:970
func (z *Int) modSqrt3Mod4Prime(x, p *Int) *Int {
//line /usr/local/go/src/math/big/int.go:970
	_go_fuzz_dep_.CoverTab[5455]++
						e := new(Int).Add(p, intOne)
						e.Rsh(e, 2)
						z.Exp(x, e, p)
						return z
//line /usr/local/go/src/math/big/int.go:974
	// _ = "end of CoverTab[5455]"
}

// modSqrt5Mod8 uses Atkin's observation that 2 is not a square mod p
//line /usr/local/go/src/math/big/int.go:977
//
//line /usr/local/go/src/math/big/int.go:977
//	alpha ==  (2*a)^((p-5)/8)    mod p
//line /usr/local/go/src/math/big/int.go:977
//	beta  ==  2*a*alpha^2        mod p  is a square root of -1
//line /usr/local/go/src/math/big/int.go:977
//	b     ==  a*alpha*(beta-1)   mod p  is a square root of a
//line /usr/local/go/src/math/big/int.go:977
//
//line /usr/local/go/src/math/big/int.go:977
// to calculate the square root of any quadratic residue mod p quickly for 5
//line /usr/local/go/src/math/big/int.go:977
// mod 8 primes.
//line /usr/local/go/src/math/big/int.go:985
func (z *Int) modSqrt5Mod8Prime(x, p *Int) *Int {
//line /usr/local/go/src/math/big/int.go:985
	_go_fuzz_dep_.CoverTab[5456]++

//line /usr/local/go/src/math/big/int.go:988
	e := new(Int).Rsh(p, 3)
						tx := new(Int).Lsh(x, 1)
						alpha := new(Int).Exp(tx, e, p)
						beta := new(Int).Mul(alpha, alpha)
						beta.Mod(beta, p)
						beta.Mul(beta, tx)
						beta.Mod(beta, p)
						beta.Sub(beta, intOne)
						beta.Mul(beta, x)
						beta.Mod(beta, p)
						beta.Mul(beta, alpha)
						z.Mod(beta, p)
						return z
//line /usr/local/go/src/math/big/int.go:1000
	// _ = "end of CoverTab[5456]"
}

// modSqrtTonelliShanks uses the Tonelli-Shanks algorithm to find the square
//line /usr/local/go/src/math/big/int.go:1003
// root of a quadratic residue modulo any prime.
//line /usr/local/go/src/math/big/int.go:1005
func (z *Int) modSqrtTonelliShanks(x, p *Int) *Int {
//line /usr/local/go/src/math/big/int.go:1005
	_go_fuzz_dep_.CoverTab[5457]++
	// Break p-1 into s*2^e such that s is odd.
	var s Int
	s.Sub(p, intOne)
	e := s.abs.trailingZeroBits()
	s.Rsh(&s, e)

	// find some non-square n
	var n Int
	n.SetInt64(2)
	for Jacobi(&n, p) != -1 {
//line /usr/local/go/src/math/big/int.go:1015
		_go_fuzz_dep_.CoverTab[5459]++
							n.Add(&n, intOne)
//line /usr/local/go/src/math/big/int.go:1016
		// _ = "end of CoverTab[5459]"
	}
//line /usr/local/go/src/math/big/int.go:1017
	// _ = "end of CoverTab[5457]"
//line /usr/local/go/src/math/big/int.go:1017
	_go_fuzz_dep_.CoverTab[5458]++

	// Core of the Tonelli-Shanks algorithm. Follows the description in
	// section 6 of "Square roots from 1; 24, 51, 10 to Dan Shanks" by Ezra
	// Brown:
	// https://www.maa.org/sites/default/files/pdf/upload_library/22/Polya/07468342.di020786.02p0470a.pdf
	var y, b, g, t Int
	y.Add(&s, intOne)
	y.Rsh(&y, 1)
	y.Exp(x, &y, p)
	b.Exp(x, &s, p)
	g.Exp(&n, &s, p)
	r := e
	for {
//line /usr/local/go/src/math/big/int.go:1030
		_go_fuzz_dep_.CoverTab[5460]++
		// find the least m such that ord_p(b) = 2^m
		var m uint
		t.Set(&b)
		for t.Cmp(intOne) != 0 {
//line /usr/local/go/src/math/big/int.go:1034
			_go_fuzz_dep_.CoverTab[5463]++
								t.Mul(&t, &t).Mod(&t, p)
								m++
//line /usr/local/go/src/math/big/int.go:1036
			// _ = "end of CoverTab[5463]"
		}
//line /usr/local/go/src/math/big/int.go:1037
		// _ = "end of CoverTab[5460]"
//line /usr/local/go/src/math/big/int.go:1037
		_go_fuzz_dep_.CoverTab[5461]++

							if m == 0 {
//line /usr/local/go/src/math/big/int.go:1039
			_go_fuzz_dep_.CoverTab[5464]++
								return z.Set(&y)
//line /usr/local/go/src/math/big/int.go:1040
			// _ = "end of CoverTab[5464]"
		} else {
//line /usr/local/go/src/math/big/int.go:1041
			_go_fuzz_dep_.CoverTab[5465]++
//line /usr/local/go/src/math/big/int.go:1041
			// _ = "end of CoverTab[5465]"
//line /usr/local/go/src/math/big/int.go:1041
		}
//line /usr/local/go/src/math/big/int.go:1041
		// _ = "end of CoverTab[5461]"
//line /usr/local/go/src/math/big/int.go:1041
		_go_fuzz_dep_.CoverTab[5462]++

							t.SetInt64(0).SetBit(&t, int(r-m-1), 1).Exp(&g, &t, p)

							g.Mul(&t, &t).Mod(&g, p)
							y.Mul(&y, &t).Mod(&y, p)
							b.Mul(&b, &g).Mod(&b, p)
							r = m
//line /usr/local/go/src/math/big/int.go:1048
		// _ = "end of CoverTab[5462]"
	}
//line /usr/local/go/src/math/big/int.go:1049
	// _ = "end of CoverTab[5458]"
}

// ModSqrt sets z to a square root of x mod p if such a square root exists, and
//line /usr/local/go/src/math/big/int.go:1052
// returns z. The modulus p must be an odd prime. If x is not a square mod p,
//line /usr/local/go/src/math/big/int.go:1052
// ModSqrt leaves z unchanged and returns nil. This function panics if p is
//line /usr/local/go/src/math/big/int.go:1052
// not an odd integer, its behavior is undefined if p is odd but not prime.
//line /usr/local/go/src/math/big/int.go:1056
func (z *Int) ModSqrt(x, p *Int) *Int {
//line /usr/local/go/src/math/big/int.go:1056
	_go_fuzz_dep_.CoverTab[5466]++
						switch Jacobi(x, p) {
	case -1:
//line /usr/local/go/src/math/big/int.go:1058
		_go_fuzz_dep_.CoverTab[5469]++
							return nil
//line /usr/local/go/src/math/big/int.go:1059
		// _ = "end of CoverTab[5469]"
	case 0:
//line /usr/local/go/src/math/big/int.go:1060
		_go_fuzz_dep_.CoverTab[5470]++
							return z.SetInt64(0)
//line /usr/local/go/src/math/big/int.go:1061
		// _ = "end of CoverTab[5470]"
	case 1:
//line /usr/local/go/src/math/big/int.go:1062
		_go_fuzz_dep_.CoverTab[5471]++
							break
//line /usr/local/go/src/math/big/int.go:1063
		// _ = "end of CoverTab[5471]"
//line /usr/local/go/src/math/big/int.go:1063
	default:
//line /usr/local/go/src/math/big/int.go:1063
		_go_fuzz_dep_.CoverTab[5472]++
//line /usr/local/go/src/math/big/int.go:1063
		// _ = "end of CoverTab[5472]"
	}
//line /usr/local/go/src/math/big/int.go:1064
	// _ = "end of CoverTab[5466]"
//line /usr/local/go/src/math/big/int.go:1064
	_go_fuzz_dep_.CoverTab[5467]++
						if x.neg || func() bool {
//line /usr/local/go/src/math/big/int.go:1065
		_go_fuzz_dep_.CoverTab[5473]++
//line /usr/local/go/src/math/big/int.go:1065
		return x.Cmp(p) >= 0
//line /usr/local/go/src/math/big/int.go:1065
		// _ = "end of CoverTab[5473]"
//line /usr/local/go/src/math/big/int.go:1065
	}() {
//line /usr/local/go/src/math/big/int.go:1065
		_go_fuzz_dep_.CoverTab[5474]++
							x = new(Int).Mod(x, p)
//line /usr/local/go/src/math/big/int.go:1066
		// _ = "end of CoverTab[5474]"
	} else {
//line /usr/local/go/src/math/big/int.go:1067
		_go_fuzz_dep_.CoverTab[5475]++
//line /usr/local/go/src/math/big/int.go:1067
		// _ = "end of CoverTab[5475]"
//line /usr/local/go/src/math/big/int.go:1067
	}
//line /usr/local/go/src/math/big/int.go:1067
	// _ = "end of CoverTab[5467]"
//line /usr/local/go/src/math/big/int.go:1067
	_go_fuzz_dep_.CoverTab[5468]++

						switch {
	case p.abs[0]%4 == 3:
//line /usr/local/go/src/math/big/int.go:1070
		_go_fuzz_dep_.CoverTab[5476]++

							return z.modSqrt3Mod4Prime(x, p)
//line /usr/local/go/src/math/big/int.go:1072
		// _ = "end of CoverTab[5476]"
	case p.abs[0]%8 == 5:
//line /usr/local/go/src/math/big/int.go:1073
		_go_fuzz_dep_.CoverTab[5477]++

							return z.modSqrt5Mod8Prime(x, p)
//line /usr/local/go/src/math/big/int.go:1075
		// _ = "end of CoverTab[5477]"
	default:
//line /usr/local/go/src/math/big/int.go:1076
		_go_fuzz_dep_.CoverTab[5478]++

							return z.modSqrtTonelliShanks(x, p)
//line /usr/local/go/src/math/big/int.go:1078
		// _ = "end of CoverTab[5478]"
	}
//line /usr/local/go/src/math/big/int.go:1079
	// _ = "end of CoverTab[5468]"
}

// Lsh sets z = x << n and returns z.
func (z *Int) Lsh(x *Int, n uint) *Int {
//line /usr/local/go/src/math/big/int.go:1083
	_go_fuzz_dep_.CoverTab[5479]++
						z.abs = z.abs.shl(x.abs, n)
						z.neg = x.neg
						return z
//line /usr/local/go/src/math/big/int.go:1086
	// _ = "end of CoverTab[5479]"
}

// Rsh sets z = x >> n and returns z.
func (z *Int) Rsh(x *Int, n uint) *Int {
//line /usr/local/go/src/math/big/int.go:1090
	_go_fuzz_dep_.CoverTab[5480]++
						if x.neg {
//line /usr/local/go/src/math/big/int.go:1091
		_go_fuzz_dep_.CoverTab[5482]++

							t := z.abs.sub(x.abs, natOne)
							t = t.shr(t, n)
							z.abs = t.add(t, natOne)
							z.neg = true
							return z
//line /usr/local/go/src/math/big/int.go:1097
		// _ = "end of CoverTab[5482]"
	} else {
//line /usr/local/go/src/math/big/int.go:1098
		_go_fuzz_dep_.CoverTab[5483]++
//line /usr/local/go/src/math/big/int.go:1098
		// _ = "end of CoverTab[5483]"
//line /usr/local/go/src/math/big/int.go:1098
	}
//line /usr/local/go/src/math/big/int.go:1098
	// _ = "end of CoverTab[5480]"
//line /usr/local/go/src/math/big/int.go:1098
	_go_fuzz_dep_.CoverTab[5481]++

						z.abs = z.abs.shr(x.abs, n)
						z.neg = false
						return z
//line /usr/local/go/src/math/big/int.go:1102
	// _ = "end of CoverTab[5481]"
}

// Bit returns the value of the i'th bit of x. That is, it
//line /usr/local/go/src/math/big/int.go:1105
// returns (x>>i)&1. The bit index i must be >= 0.
//line /usr/local/go/src/math/big/int.go:1107
func (x *Int) Bit(i int) uint {
//line /usr/local/go/src/math/big/int.go:1107
	_go_fuzz_dep_.CoverTab[5484]++
						if i == 0 {
//line /usr/local/go/src/math/big/int.go:1108
		_go_fuzz_dep_.CoverTab[5488]++

							if len(x.abs) > 0 {
//line /usr/local/go/src/math/big/int.go:1110
			_go_fuzz_dep_.CoverTab[5490]++
								return uint(x.abs[0] & 1)
//line /usr/local/go/src/math/big/int.go:1111
			// _ = "end of CoverTab[5490]"
		} else {
//line /usr/local/go/src/math/big/int.go:1112
			_go_fuzz_dep_.CoverTab[5491]++
//line /usr/local/go/src/math/big/int.go:1112
			// _ = "end of CoverTab[5491]"
//line /usr/local/go/src/math/big/int.go:1112
		}
//line /usr/local/go/src/math/big/int.go:1112
		// _ = "end of CoverTab[5488]"
//line /usr/local/go/src/math/big/int.go:1112
		_go_fuzz_dep_.CoverTab[5489]++
							return 0
//line /usr/local/go/src/math/big/int.go:1113
		// _ = "end of CoverTab[5489]"
	} else {
//line /usr/local/go/src/math/big/int.go:1114
		_go_fuzz_dep_.CoverTab[5492]++
//line /usr/local/go/src/math/big/int.go:1114
		// _ = "end of CoverTab[5492]"
//line /usr/local/go/src/math/big/int.go:1114
	}
//line /usr/local/go/src/math/big/int.go:1114
	// _ = "end of CoverTab[5484]"
//line /usr/local/go/src/math/big/int.go:1114
	_go_fuzz_dep_.CoverTab[5485]++
						if i < 0 {
//line /usr/local/go/src/math/big/int.go:1115
		_go_fuzz_dep_.CoverTab[5493]++
							panic("negative bit index")
//line /usr/local/go/src/math/big/int.go:1116
		// _ = "end of CoverTab[5493]"
	} else {
//line /usr/local/go/src/math/big/int.go:1117
		_go_fuzz_dep_.CoverTab[5494]++
//line /usr/local/go/src/math/big/int.go:1117
		// _ = "end of CoverTab[5494]"
//line /usr/local/go/src/math/big/int.go:1117
	}
//line /usr/local/go/src/math/big/int.go:1117
	// _ = "end of CoverTab[5485]"
//line /usr/local/go/src/math/big/int.go:1117
	_go_fuzz_dep_.CoverTab[5486]++
						if x.neg {
//line /usr/local/go/src/math/big/int.go:1118
		_go_fuzz_dep_.CoverTab[5495]++
							t := nat(nil).sub(x.abs, natOne)
							return t.bit(uint(i)) ^ 1
//line /usr/local/go/src/math/big/int.go:1120
		// _ = "end of CoverTab[5495]"
	} else {
//line /usr/local/go/src/math/big/int.go:1121
		_go_fuzz_dep_.CoverTab[5496]++
//line /usr/local/go/src/math/big/int.go:1121
		// _ = "end of CoverTab[5496]"
//line /usr/local/go/src/math/big/int.go:1121
	}
//line /usr/local/go/src/math/big/int.go:1121
	// _ = "end of CoverTab[5486]"
//line /usr/local/go/src/math/big/int.go:1121
	_go_fuzz_dep_.CoverTab[5487]++

						return x.abs.bit(uint(i))
//line /usr/local/go/src/math/big/int.go:1123
	// _ = "end of CoverTab[5487]"
}

// SetBit sets z to x, with x's i'th bit set to b (0 or 1).
//line /usr/local/go/src/math/big/int.go:1126
// That is, if b is 1 SetBit sets z = x | (1 << i);
//line /usr/local/go/src/math/big/int.go:1126
// if b is 0 SetBit sets z = x &^ (1 << i). If b is not 0 or 1,
//line /usr/local/go/src/math/big/int.go:1126
// SetBit will panic.
//line /usr/local/go/src/math/big/int.go:1130
func (z *Int) SetBit(x *Int, i int, b uint) *Int {
//line /usr/local/go/src/math/big/int.go:1130
	_go_fuzz_dep_.CoverTab[5497]++
						if i < 0 {
//line /usr/local/go/src/math/big/int.go:1131
		_go_fuzz_dep_.CoverTab[5500]++
							panic("negative bit index")
//line /usr/local/go/src/math/big/int.go:1132
		// _ = "end of CoverTab[5500]"
	} else {
//line /usr/local/go/src/math/big/int.go:1133
		_go_fuzz_dep_.CoverTab[5501]++
//line /usr/local/go/src/math/big/int.go:1133
		// _ = "end of CoverTab[5501]"
//line /usr/local/go/src/math/big/int.go:1133
	}
//line /usr/local/go/src/math/big/int.go:1133
	// _ = "end of CoverTab[5497]"
//line /usr/local/go/src/math/big/int.go:1133
	_go_fuzz_dep_.CoverTab[5498]++
						if x.neg {
//line /usr/local/go/src/math/big/int.go:1134
		_go_fuzz_dep_.CoverTab[5502]++
							t := z.abs.sub(x.abs, natOne)
							t = t.setBit(t, uint(i), b^1)
							z.abs = t.add(t, natOne)
							z.neg = len(z.abs) > 0
							return z
//line /usr/local/go/src/math/big/int.go:1139
		// _ = "end of CoverTab[5502]"
	} else {
//line /usr/local/go/src/math/big/int.go:1140
		_go_fuzz_dep_.CoverTab[5503]++
//line /usr/local/go/src/math/big/int.go:1140
		// _ = "end of CoverTab[5503]"
//line /usr/local/go/src/math/big/int.go:1140
	}
//line /usr/local/go/src/math/big/int.go:1140
	// _ = "end of CoverTab[5498]"
//line /usr/local/go/src/math/big/int.go:1140
	_go_fuzz_dep_.CoverTab[5499]++
						z.abs = z.abs.setBit(x.abs, uint(i), b)
						z.neg = false
						return z
//line /usr/local/go/src/math/big/int.go:1143
	// _ = "end of CoverTab[5499]"
}

// And sets z = x & y and returns z.
func (z *Int) And(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:1147
	_go_fuzz_dep_.CoverTab[5504]++
						if x.neg == y.neg {
//line /usr/local/go/src/math/big/int.go:1148
		_go_fuzz_dep_.CoverTab[5507]++
							if x.neg {
//line /usr/local/go/src/math/big/int.go:1149
			_go_fuzz_dep_.CoverTab[5509]++

								x1 := nat(nil).sub(x.abs, natOne)
								y1 := nat(nil).sub(y.abs, natOne)
								z.abs = z.abs.add(z.abs.or(x1, y1), natOne)
								z.neg = true
								return z
//line /usr/local/go/src/math/big/int.go:1155
			// _ = "end of CoverTab[5509]"
		} else {
//line /usr/local/go/src/math/big/int.go:1156
			_go_fuzz_dep_.CoverTab[5510]++
//line /usr/local/go/src/math/big/int.go:1156
			// _ = "end of CoverTab[5510]"
//line /usr/local/go/src/math/big/int.go:1156
		}
//line /usr/local/go/src/math/big/int.go:1156
		// _ = "end of CoverTab[5507]"
//line /usr/local/go/src/math/big/int.go:1156
		_go_fuzz_dep_.CoverTab[5508]++

//line /usr/local/go/src/math/big/int.go:1159
		z.abs = z.abs.and(x.abs, y.abs)
							z.neg = false
							return z
//line /usr/local/go/src/math/big/int.go:1161
		// _ = "end of CoverTab[5508]"
	} else {
//line /usr/local/go/src/math/big/int.go:1162
		_go_fuzz_dep_.CoverTab[5511]++
//line /usr/local/go/src/math/big/int.go:1162
		// _ = "end of CoverTab[5511]"
//line /usr/local/go/src/math/big/int.go:1162
	}
//line /usr/local/go/src/math/big/int.go:1162
	// _ = "end of CoverTab[5504]"
//line /usr/local/go/src/math/big/int.go:1162
	_go_fuzz_dep_.CoverTab[5505]++

//line /usr/local/go/src/math/big/int.go:1165
	if x.neg {
//line /usr/local/go/src/math/big/int.go:1165
		_go_fuzz_dep_.CoverTab[5512]++
							x, y = y, x
//line /usr/local/go/src/math/big/int.go:1166
		// _ = "end of CoverTab[5512]"
	} else {
//line /usr/local/go/src/math/big/int.go:1167
		_go_fuzz_dep_.CoverTab[5513]++
//line /usr/local/go/src/math/big/int.go:1167
		// _ = "end of CoverTab[5513]"
//line /usr/local/go/src/math/big/int.go:1167
	}
//line /usr/local/go/src/math/big/int.go:1167
	// _ = "end of CoverTab[5505]"
//line /usr/local/go/src/math/big/int.go:1167
	_go_fuzz_dep_.CoverTab[5506]++

//line /usr/local/go/src/math/big/int.go:1170
	y1 := nat(nil).sub(y.abs, natOne)
						z.abs = z.abs.andNot(x.abs, y1)
						z.neg = false
						return z
//line /usr/local/go/src/math/big/int.go:1173
	// _ = "end of CoverTab[5506]"
}

// AndNot sets z = x &^ y and returns z.
func (z *Int) AndNot(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:1177
	_go_fuzz_dep_.CoverTab[5514]++
						if x.neg == y.neg {
//line /usr/local/go/src/math/big/int.go:1178
		_go_fuzz_dep_.CoverTab[5517]++
							if x.neg {
//line /usr/local/go/src/math/big/int.go:1179
			_go_fuzz_dep_.CoverTab[5519]++

								x1 := nat(nil).sub(x.abs, natOne)
								y1 := nat(nil).sub(y.abs, natOne)
								z.abs = z.abs.andNot(y1, x1)
								z.neg = false
								return z
//line /usr/local/go/src/math/big/int.go:1185
			// _ = "end of CoverTab[5519]"
		} else {
//line /usr/local/go/src/math/big/int.go:1186
			_go_fuzz_dep_.CoverTab[5520]++
//line /usr/local/go/src/math/big/int.go:1186
			// _ = "end of CoverTab[5520]"
//line /usr/local/go/src/math/big/int.go:1186
		}
//line /usr/local/go/src/math/big/int.go:1186
		// _ = "end of CoverTab[5517]"
//line /usr/local/go/src/math/big/int.go:1186
		_go_fuzz_dep_.CoverTab[5518]++

//line /usr/local/go/src/math/big/int.go:1189
		z.abs = z.abs.andNot(x.abs, y.abs)
							z.neg = false
							return z
//line /usr/local/go/src/math/big/int.go:1191
		// _ = "end of CoverTab[5518]"
	} else {
//line /usr/local/go/src/math/big/int.go:1192
		_go_fuzz_dep_.CoverTab[5521]++
//line /usr/local/go/src/math/big/int.go:1192
		// _ = "end of CoverTab[5521]"
//line /usr/local/go/src/math/big/int.go:1192
	}
//line /usr/local/go/src/math/big/int.go:1192
	// _ = "end of CoverTab[5514]"
//line /usr/local/go/src/math/big/int.go:1192
	_go_fuzz_dep_.CoverTab[5515]++

						if x.neg {
//line /usr/local/go/src/math/big/int.go:1194
		_go_fuzz_dep_.CoverTab[5522]++

							x1 := nat(nil).sub(x.abs, natOne)
							z.abs = z.abs.add(z.abs.or(x1, y.abs), natOne)
							z.neg = true
							return z
//line /usr/local/go/src/math/big/int.go:1199
		// _ = "end of CoverTab[5522]"
	} else {
//line /usr/local/go/src/math/big/int.go:1200
		_go_fuzz_dep_.CoverTab[5523]++
//line /usr/local/go/src/math/big/int.go:1200
		// _ = "end of CoverTab[5523]"
//line /usr/local/go/src/math/big/int.go:1200
	}
//line /usr/local/go/src/math/big/int.go:1200
	// _ = "end of CoverTab[5515]"
//line /usr/local/go/src/math/big/int.go:1200
	_go_fuzz_dep_.CoverTab[5516]++

//line /usr/local/go/src/math/big/int.go:1203
	y1 := nat(nil).sub(y.abs, natOne)
						z.abs = z.abs.and(x.abs, y1)
						z.neg = false
						return z
//line /usr/local/go/src/math/big/int.go:1206
	// _ = "end of CoverTab[5516]"
}

// Or sets z = x | y and returns z.
func (z *Int) Or(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:1210
	_go_fuzz_dep_.CoverTab[5524]++
						if x.neg == y.neg {
//line /usr/local/go/src/math/big/int.go:1211
		_go_fuzz_dep_.CoverTab[5527]++
							if x.neg {
//line /usr/local/go/src/math/big/int.go:1212
			_go_fuzz_dep_.CoverTab[5529]++

								x1 := nat(nil).sub(x.abs, natOne)
								y1 := nat(nil).sub(y.abs, natOne)
								z.abs = z.abs.add(z.abs.and(x1, y1), natOne)
								z.neg = true
								return z
//line /usr/local/go/src/math/big/int.go:1218
			// _ = "end of CoverTab[5529]"
		} else {
//line /usr/local/go/src/math/big/int.go:1219
			_go_fuzz_dep_.CoverTab[5530]++
//line /usr/local/go/src/math/big/int.go:1219
			// _ = "end of CoverTab[5530]"
//line /usr/local/go/src/math/big/int.go:1219
		}
//line /usr/local/go/src/math/big/int.go:1219
		// _ = "end of CoverTab[5527]"
//line /usr/local/go/src/math/big/int.go:1219
		_go_fuzz_dep_.CoverTab[5528]++

//line /usr/local/go/src/math/big/int.go:1222
		z.abs = z.abs.or(x.abs, y.abs)
							z.neg = false
							return z
//line /usr/local/go/src/math/big/int.go:1224
		// _ = "end of CoverTab[5528]"
	} else {
//line /usr/local/go/src/math/big/int.go:1225
		_go_fuzz_dep_.CoverTab[5531]++
//line /usr/local/go/src/math/big/int.go:1225
		// _ = "end of CoverTab[5531]"
//line /usr/local/go/src/math/big/int.go:1225
	}
//line /usr/local/go/src/math/big/int.go:1225
	// _ = "end of CoverTab[5524]"
//line /usr/local/go/src/math/big/int.go:1225
	_go_fuzz_dep_.CoverTab[5525]++

//line /usr/local/go/src/math/big/int.go:1228
	if x.neg {
//line /usr/local/go/src/math/big/int.go:1228
		_go_fuzz_dep_.CoverTab[5532]++
							x, y = y, x
//line /usr/local/go/src/math/big/int.go:1229
		// _ = "end of CoverTab[5532]"
	} else {
//line /usr/local/go/src/math/big/int.go:1230
		_go_fuzz_dep_.CoverTab[5533]++
//line /usr/local/go/src/math/big/int.go:1230
		// _ = "end of CoverTab[5533]"
//line /usr/local/go/src/math/big/int.go:1230
	}
//line /usr/local/go/src/math/big/int.go:1230
	// _ = "end of CoverTab[5525]"
//line /usr/local/go/src/math/big/int.go:1230
	_go_fuzz_dep_.CoverTab[5526]++

//line /usr/local/go/src/math/big/int.go:1233
	y1 := nat(nil).sub(y.abs, natOne)
						z.abs = z.abs.add(z.abs.andNot(y1, x.abs), natOne)
						z.neg = true
						return z
//line /usr/local/go/src/math/big/int.go:1236
	// _ = "end of CoverTab[5526]"
}

// Xor sets z = x ^ y and returns z.
func (z *Int) Xor(x, y *Int) *Int {
//line /usr/local/go/src/math/big/int.go:1240
	_go_fuzz_dep_.CoverTab[5534]++
						if x.neg == y.neg {
//line /usr/local/go/src/math/big/int.go:1241
		_go_fuzz_dep_.CoverTab[5537]++
							if x.neg {
//line /usr/local/go/src/math/big/int.go:1242
			_go_fuzz_dep_.CoverTab[5539]++

								x1 := nat(nil).sub(x.abs, natOne)
								y1 := nat(nil).sub(y.abs, natOne)
								z.abs = z.abs.xor(x1, y1)
								z.neg = false
								return z
//line /usr/local/go/src/math/big/int.go:1248
			// _ = "end of CoverTab[5539]"
		} else {
//line /usr/local/go/src/math/big/int.go:1249
			_go_fuzz_dep_.CoverTab[5540]++
//line /usr/local/go/src/math/big/int.go:1249
			// _ = "end of CoverTab[5540]"
//line /usr/local/go/src/math/big/int.go:1249
		}
//line /usr/local/go/src/math/big/int.go:1249
		// _ = "end of CoverTab[5537]"
//line /usr/local/go/src/math/big/int.go:1249
		_go_fuzz_dep_.CoverTab[5538]++

//line /usr/local/go/src/math/big/int.go:1252
		z.abs = z.abs.xor(x.abs, y.abs)
							z.neg = false
							return z
//line /usr/local/go/src/math/big/int.go:1254
		// _ = "end of CoverTab[5538]"
	} else {
//line /usr/local/go/src/math/big/int.go:1255
		_go_fuzz_dep_.CoverTab[5541]++
//line /usr/local/go/src/math/big/int.go:1255
		// _ = "end of CoverTab[5541]"
//line /usr/local/go/src/math/big/int.go:1255
	}
//line /usr/local/go/src/math/big/int.go:1255
	// _ = "end of CoverTab[5534]"
//line /usr/local/go/src/math/big/int.go:1255
	_go_fuzz_dep_.CoverTab[5535]++

//line /usr/local/go/src/math/big/int.go:1258
	if x.neg {
//line /usr/local/go/src/math/big/int.go:1258
		_go_fuzz_dep_.CoverTab[5542]++
							x, y = y, x
//line /usr/local/go/src/math/big/int.go:1259
		// _ = "end of CoverTab[5542]"
	} else {
//line /usr/local/go/src/math/big/int.go:1260
		_go_fuzz_dep_.CoverTab[5543]++
//line /usr/local/go/src/math/big/int.go:1260
		// _ = "end of CoverTab[5543]"
//line /usr/local/go/src/math/big/int.go:1260
	}
//line /usr/local/go/src/math/big/int.go:1260
	// _ = "end of CoverTab[5535]"
//line /usr/local/go/src/math/big/int.go:1260
	_go_fuzz_dep_.CoverTab[5536]++

//line /usr/local/go/src/math/big/int.go:1263
	y1 := nat(nil).sub(y.abs, natOne)
						z.abs = z.abs.add(z.abs.xor(x.abs, y1), natOne)
						z.neg = true
						return z
//line /usr/local/go/src/math/big/int.go:1266
	// _ = "end of CoverTab[5536]"
}

// Not sets z = ^x and returns z.
func (z *Int) Not(x *Int) *Int {
//line /usr/local/go/src/math/big/int.go:1270
	_go_fuzz_dep_.CoverTab[5544]++
						if x.neg {
//line /usr/local/go/src/math/big/int.go:1271
		_go_fuzz_dep_.CoverTab[5546]++

							z.abs = z.abs.sub(x.abs, natOne)
							z.neg = false
							return z
//line /usr/local/go/src/math/big/int.go:1275
		// _ = "end of CoverTab[5546]"
	} else {
//line /usr/local/go/src/math/big/int.go:1276
		_go_fuzz_dep_.CoverTab[5547]++
//line /usr/local/go/src/math/big/int.go:1276
		// _ = "end of CoverTab[5547]"
//line /usr/local/go/src/math/big/int.go:1276
	}
//line /usr/local/go/src/math/big/int.go:1276
	// _ = "end of CoverTab[5544]"
//line /usr/local/go/src/math/big/int.go:1276
	_go_fuzz_dep_.CoverTab[5545]++

//line /usr/local/go/src/math/big/int.go:1279
	z.abs = z.abs.add(x.abs, natOne)
						z.neg = true
						return z
//line /usr/local/go/src/math/big/int.go:1281
	// _ = "end of CoverTab[5545]"
}

// Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z.
//line /usr/local/go/src/math/big/int.go:1284
// It panics if x is negative.
//line /usr/local/go/src/math/big/int.go:1286
func (z *Int) Sqrt(x *Int) *Int {
//line /usr/local/go/src/math/big/int.go:1286
	_go_fuzz_dep_.CoverTab[5548]++
						if x.neg {
//line /usr/local/go/src/math/big/int.go:1287
		_go_fuzz_dep_.CoverTab[5550]++
							panic("square root of negative number")
//line /usr/local/go/src/math/big/int.go:1288
		// _ = "end of CoverTab[5550]"
	} else {
//line /usr/local/go/src/math/big/int.go:1289
		_go_fuzz_dep_.CoverTab[5551]++
//line /usr/local/go/src/math/big/int.go:1289
		// _ = "end of CoverTab[5551]"
//line /usr/local/go/src/math/big/int.go:1289
	}
//line /usr/local/go/src/math/big/int.go:1289
	// _ = "end of CoverTab[5548]"
//line /usr/local/go/src/math/big/int.go:1289
	_go_fuzz_dep_.CoverTab[5549]++
						z.neg = false
						z.abs = z.abs.sqrt(x.abs)
						return z
//line /usr/local/go/src/math/big/int.go:1292
	// _ = "end of CoverTab[5549]"
}

//line /usr/local/go/src/math/big/int.go:1293
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/int.go:1293
var _ = _go_fuzz_dep_.CoverTab
