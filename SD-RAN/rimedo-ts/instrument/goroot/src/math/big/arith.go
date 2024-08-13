// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides Go implementations of elementary multi-precision
// arithmetic operations on word vectors. These have the suffix _g.
// These are needed for platforms without assembly implementations of these routines.
// This file also contains elementary operations that can be implemented
// sufficiently efficiently in Go.

//line /usr/local/go/src/math/big/arith.go:11
package big

//line /usr/local/go/src/math/big/arith.go:11
import (
//line /usr/local/go/src/math/big/arith.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/arith.go:11
)
//line /usr/local/go/src/math/big/arith.go:11
import (
//line /usr/local/go/src/math/big/arith.go:11
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/arith.go:11
)

import "math/bits"

// A Word represents a single digit of a multi-precision unsigned integer.
type Word uint

const (
	_S	= _W / 8	// word size in bytes

	_W	= bits.UintSize	// word size in bits
	_B	= 1 << _W	// digit base
	_M	= _B - 1	// digit mask
)

//line /usr/local/go/src/math/big/arith.go:43
// z1<<_W + z0 = x*y
func mulWW(x, y Word) (z1, z0 Word) {
//line /usr/local/go/src/math/big/arith.go:44
	_go_fuzz_dep_.CoverTab[4045]++
						hi, lo := bits.Mul(uint(x), uint(y))
						return Word(hi), Word(lo)
//line /usr/local/go/src/math/big/arith.go:46
	// _ = "end of CoverTab[4045]"
}

// z1<<_W + z0 = x*y + c
func mulAddWWW_g(x, y, c Word) (z1, z0 Word) {
//line /usr/local/go/src/math/big/arith.go:50
	_go_fuzz_dep_.CoverTab[4046]++
						hi, lo := bits.Mul(uint(x), uint(y))
						var cc uint
						lo, cc = bits.Add(lo, uint(c), 0)
						return Word(hi + cc), Word(lo)
//line /usr/local/go/src/math/big/arith.go:54
	// _ = "end of CoverTab[4046]"
}

// nlz returns the number of leading zeros in x.
//line /usr/local/go/src/math/big/arith.go:57
// Wraps bits.LeadingZeros call for convenience.
//line /usr/local/go/src/math/big/arith.go:59
func nlz(x Word) uint {
//line /usr/local/go/src/math/big/arith.go:59
	_go_fuzz_dep_.CoverTab[4047]++
						return uint(bits.LeadingZeros(uint(x)))
//line /usr/local/go/src/math/big/arith.go:60
	// _ = "end of CoverTab[4047]"
}

// The resulting carry c is either 0 or 1.
func addVV_g(z, x, y []Word) (c Word) {
//line /usr/local/go/src/math/big/arith.go:64
	_go_fuzz_dep_.CoverTab[4048]++

						for i := 0; i < len(z) && func() bool {
//line /usr/local/go/src/math/big/arith.go:66
		_go_fuzz_dep_.CoverTab[4050]++
//line /usr/local/go/src/math/big/arith.go:66
		return i < len(x)
//line /usr/local/go/src/math/big/arith.go:66
		// _ = "end of CoverTab[4050]"
//line /usr/local/go/src/math/big/arith.go:66
	}() && func() bool {
//line /usr/local/go/src/math/big/arith.go:66
		_go_fuzz_dep_.CoverTab[4051]++
//line /usr/local/go/src/math/big/arith.go:66
		return i < len(y)
//line /usr/local/go/src/math/big/arith.go:66
		// _ = "end of CoverTab[4051]"
//line /usr/local/go/src/math/big/arith.go:66
	}(); i++ {
//line /usr/local/go/src/math/big/arith.go:66
		_go_fuzz_dep_.CoverTab[4052]++
							zi, cc := bits.Add(uint(x[i]), uint(y[i]), uint(c))
							z[i] = Word(zi)
							c = Word(cc)
//line /usr/local/go/src/math/big/arith.go:69
		// _ = "end of CoverTab[4052]"
	}
//line /usr/local/go/src/math/big/arith.go:70
	// _ = "end of CoverTab[4048]"
//line /usr/local/go/src/math/big/arith.go:70
	_go_fuzz_dep_.CoverTab[4049]++
						return
//line /usr/local/go/src/math/big/arith.go:71
	// _ = "end of CoverTab[4049]"
}

// The resulting carry c is either 0 or 1.
func subVV_g(z, x, y []Word) (c Word) {
//line /usr/local/go/src/math/big/arith.go:75
	_go_fuzz_dep_.CoverTab[4053]++

						for i := 0; i < len(z) && func() bool {
//line /usr/local/go/src/math/big/arith.go:77
		_go_fuzz_dep_.CoverTab[4055]++
//line /usr/local/go/src/math/big/arith.go:77
		return i < len(x)
//line /usr/local/go/src/math/big/arith.go:77
		// _ = "end of CoverTab[4055]"
//line /usr/local/go/src/math/big/arith.go:77
	}() && func() bool {
//line /usr/local/go/src/math/big/arith.go:77
		_go_fuzz_dep_.CoverTab[4056]++
//line /usr/local/go/src/math/big/arith.go:77
		return i < len(y)
//line /usr/local/go/src/math/big/arith.go:77
		// _ = "end of CoverTab[4056]"
//line /usr/local/go/src/math/big/arith.go:77
	}(); i++ {
//line /usr/local/go/src/math/big/arith.go:77
		_go_fuzz_dep_.CoverTab[4057]++
							zi, cc := bits.Sub(uint(x[i]), uint(y[i]), uint(c))
							z[i] = Word(zi)
							c = Word(cc)
//line /usr/local/go/src/math/big/arith.go:80
		// _ = "end of CoverTab[4057]"
	}
//line /usr/local/go/src/math/big/arith.go:81
	// _ = "end of CoverTab[4053]"
//line /usr/local/go/src/math/big/arith.go:81
	_go_fuzz_dep_.CoverTab[4054]++
						return
//line /usr/local/go/src/math/big/arith.go:82
	// _ = "end of CoverTab[4054]"
}

// The resulting carry c is either 0 or 1.
func addVW_g(z, x []Word, y Word) (c Word) {
//line /usr/local/go/src/math/big/arith.go:86
	_go_fuzz_dep_.CoverTab[4058]++
						c = y

						for i := 0; i < len(z) && func() bool {
//line /usr/local/go/src/math/big/arith.go:89
		_go_fuzz_dep_.CoverTab[4060]++
//line /usr/local/go/src/math/big/arith.go:89
		return i < len(x)
//line /usr/local/go/src/math/big/arith.go:89
		// _ = "end of CoverTab[4060]"
//line /usr/local/go/src/math/big/arith.go:89
	}(); i++ {
//line /usr/local/go/src/math/big/arith.go:89
		_go_fuzz_dep_.CoverTab[4061]++
							zi, cc := bits.Add(uint(x[i]), uint(c), 0)
							z[i] = Word(zi)
							c = Word(cc)
//line /usr/local/go/src/math/big/arith.go:92
		// _ = "end of CoverTab[4061]"
	}
//line /usr/local/go/src/math/big/arith.go:93
	// _ = "end of CoverTab[4058]"
//line /usr/local/go/src/math/big/arith.go:93
	_go_fuzz_dep_.CoverTab[4059]++
						return
//line /usr/local/go/src/math/big/arith.go:94
	// _ = "end of CoverTab[4059]"
}

// addVWlarge is addVW, but intended for large z.
//line /usr/local/go/src/math/big/arith.go:97
// The only difference is that we check on every iteration
//line /usr/local/go/src/math/big/arith.go:97
// whether we are done with carries,
//line /usr/local/go/src/math/big/arith.go:97
// and if so, switch to a much faster copy instead.
//line /usr/local/go/src/math/big/arith.go:97
// This is only a good idea for large z,
//line /usr/local/go/src/math/big/arith.go:97
// because the overhead of the check and the function call
//line /usr/local/go/src/math/big/arith.go:97
// outweigh the benefits when z is small.
//line /usr/local/go/src/math/big/arith.go:104
func addVWlarge(z, x []Word, y Word) (c Word) {
//line /usr/local/go/src/math/big/arith.go:104
	_go_fuzz_dep_.CoverTab[4062]++
						c = y

						for i := 0; i < len(z) && func() bool {
//line /usr/local/go/src/math/big/arith.go:107
		_go_fuzz_dep_.CoverTab[4064]++
//line /usr/local/go/src/math/big/arith.go:107
		return i < len(x)
//line /usr/local/go/src/math/big/arith.go:107
		// _ = "end of CoverTab[4064]"
//line /usr/local/go/src/math/big/arith.go:107
	}(); i++ {
//line /usr/local/go/src/math/big/arith.go:107
		_go_fuzz_dep_.CoverTab[4065]++
							if c == 0 {
//line /usr/local/go/src/math/big/arith.go:108
			_go_fuzz_dep_.CoverTab[4067]++
								copy(z[i:], x[i:])
								return
//line /usr/local/go/src/math/big/arith.go:110
			// _ = "end of CoverTab[4067]"
		} else {
//line /usr/local/go/src/math/big/arith.go:111
			_go_fuzz_dep_.CoverTab[4068]++
//line /usr/local/go/src/math/big/arith.go:111
			// _ = "end of CoverTab[4068]"
//line /usr/local/go/src/math/big/arith.go:111
		}
//line /usr/local/go/src/math/big/arith.go:111
		// _ = "end of CoverTab[4065]"
//line /usr/local/go/src/math/big/arith.go:111
		_go_fuzz_dep_.CoverTab[4066]++
							zi, cc := bits.Add(uint(x[i]), uint(c), 0)
							z[i] = Word(zi)
							c = Word(cc)
//line /usr/local/go/src/math/big/arith.go:114
		// _ = "end of CoverTab[4066]"
	}
//line /usr/local/go/src/math/big/arith.go:115
	// _ = "end of CoverTab[4062]"
//line /usr/local/go/src/math/big/arith.go:115
	_go_fuzz_dep_.CoverTab[4063]++
						return
//line /usr/local/go/src/math/big/arith.go:116
	// _ = "end of CoverTab[4063]"
}

func subVW_g(z, x []Word, y Word) (c Word) {
//line /usr/local/go/src/math/big/arith.go:119
	_go_fuzz_dep_.CoverTab[4069]++
						c = y

						for i := 0; i < len(z) && func() bool {
//line /usr/local/go/src/math/big/arith.go:122
		_go_fuzz_dep_.CoverTab[4071]++
//line /usr/local/go/src/math/big/arith.go:122
		return i < len(x)
//line /usr/local/go/src/math/big/arith.go:122
		// _ = "end of CoverTab[4071]"
//line /usr/local/go/src/math/big/arith.go:122
	}(); i++ {
//line /usr/local/go/src/math/big/arith.go:122
		_go_fuzz_dep_.CoverTab[4072]++
							zi, cc := bits.Sub(uint(x[i]), uint(c), 0)
							z[i] = Word(zi)
							c = Word(cc)
//line /usr/local/go/src/math/big/arith.go:125
		// _ = "end of CoverTab[4072]"
	}
//line /usr/local/go/src/math/big/arith.go:126
	// _ = "end of CoverTab[4069]"
//line /usr/local/go/src/math/big/arith.go:126
	_go_fuzz_dep_.CoverTab[4070]++
						return
//line /usr/local/go/src/math/big/arith.go:127
	// _ = "end of CoverTab[4070]"
}

// subVWlarge is to subVW as addVWlarge is to addVW.
func subVWlarge(z, x []Word, y Word) (c Word) {
//line /usr/local/go/src/math/big/arith.go:131
	_go_fuzz_dep_.CoverTab[4073]++
						c = y

						for i := 0; i < len(z) && func() bool {
//line /usr/local/go/src/math/big/arith.go:134
		_go_fuzz_dep_.CoverTab[4075]++
//line /usr/local/go/src/math/big/arith.go:134
		return i < len(x)
//line /usr/local/go/src/math/big/arith.go:134
		// _ = "end of CoverTab[4075]"
//line /usr/local/go/src/math/big/arith.go:134
	}(); i++ {
//line /usr/local/go/src/math/big/arith.go:134
		_go_fuzz_dep_.CoverTab[4076]++
							if c == 0 {
//line /usr/local/go/src/math/big/arith.go:135
			_go_fuzz_dep_.CoverTab[4078]++
								copy(z[i:], x[i:])
								return
//line /usr/local/go/src/math/big/arith.go:137
			// _ = "end of CoverTab[4078]"
		} else {
//line /usr/local/go/src/math/big/arith.go:138
			_go_fuzz_dep_.CoverTab[4079]++
//line /usr/local/go/src/math/big/arith.go:138
			// _ = "end of CoverTab[4079]"
//line /usr/local/go/src/math/big/arith.go:138
		}
//line /usr/local/go/src/math/big/arith.go:138
		// _ = "end of CoverTab[4076]"
//line /usr/local/go/src/math/big/arith.go:138
		_go_fuzz_dep_.CoverTab[4077]++
							zi, cc := bits.Sub(uint(x[i]), uint(c), 0)
							z[i] = Word(zi)
							c = Word(cc)
//line /usr/local/go/src/math/big/arith.go:141
		// _ = "end of CoverTab[4077]"
	}
//line /usr/local/go/src/math/big/arith.go:142
	// _ = "end of CoverTab[4073]"
//line /usr/local/go/src/math/big/arith.go:142
	_go_fuzz_dep_.CoverTab[4074]++
						return
//line /usr/local/go/src/math/big/arith.go:143
	// _ = "end of CoverTab[4074]"
}

func shlVU_g(z, x []Word, s uint) (c Word) {
//line /usr/local/go/src/math/big/arith.go:146
	_go_fuzz_dep_.CoverTab[4080]++
						if s == 0 {
//line /usr/local/go/src/math/big/arith.go:147
		_go_fuzz_dep_.CoverTab[4084]++
							copy(z, x)
							return
//line /usr/local/go/src/math/big/arith.go:149
		// _ = "end of CoverTab[4084]"
	} else {
//line /usr/local/go/src/math/big/arith.go:150
		_go_fuzz_dep_.CoverTab[4085]++
//line /usr/local/go/src/math/big/arith.go:150
		// _ = "end of CoverTab[4085]"
//line /usr/local/go/src/math/big/arith.go:150
	}
//line /usr/local/go/src/math/big/arith.go:150
	// _ = "end of CoverTab[4080]"
//line /usr/local/go/src/math/big/arith.go:150
	_go_fuzz_dep_.CoverTab[4081]++
						if len(z) == 0 {
//line /usr/local/go/src/math/big/arith.go:151
		_go_fuzz_dep_.CoverTab[4086]++
							return
//line /usr/local/go/src/math/big/arith.go:152
		// _ = "end of CoverTab[4086]"
	} else {
//line /usr/local/go/src/math/big/arith.go:153
		_go_fuzz_dep_.CoverTab[4087]++
//line /usr/local/go/src/math/big/arith.go:153
		// _ = "end of CoverTab[4087]"
//line /usr/local/go/src/math/big/arith.go:153
	}
//line /usr/local/go/src/math/big/arith.go:153
	// _ = "end of CoverTab[4081]"
//line /usr/local/go/src/math/big/arith.go:153
	_go_fuzz_dep_.CoverTab[4082]++
						s &= _W - 1
						ŝ := _W - s
						ŝ &= _W - 1
						c = x[len(z)-1] >> ŝ
						for i := len(z) - 1; i > 0; i-- {
//line /usr/local/go/src/math/big/arith.go:158
		_go_fuzz_dep_.CoverTab[4088]++
							z[i] = x[i]<<s | x[i-1]>>ŝ
//line /usr/local/go/src/math/big/arith.go:159
		// _ = "end of CoverTab[4088]"
	}
//line /usr/local/go/src/math/big/arith.go:160
	// _ = "end of CoverTab[4082]"
//line /usr/local/go/src/math/big/arith.go:160
	_go_fuzz_dep_.CoverTab[4083]++
						z[0] = x[0] << s
						return
//line /usr/local/go/src/math/big/arith.go:162
	// _ = "end of CoverTab[4083]"
}

func shrVU_g(z, x []Word, s uint) (c Word) {
//line /usr/local/go/src/math/big/arith.go:165
	_go_fuzz_dep_.CoverTab[4089]++
						if s == 0 {
//line /usr/local/go/src/math/big/arith.go:166
		_go_fuzz_dep_.CoverTab[4094]++
							copy(z, x)
							return
//line /usr/local/go/src/math/big/arith.go:168
		// _ = "end of CoverTab[4094]"
	} else {
//line /usr/local/go/src/math/big/arith.go:169
		_go_fuzz_dep_.CoverTab[4095]++
//line /usr/local/go/src/math/big/arith.go:169
		// _ = "end of CoverTab[4095]"
//line /usr/local/go/src/math/big/arith.go:169
	}
//line /usr/local/go/src/math/big/arith.go:169
	// _ = "end of CoverTab[4089]"
//line /usr/local/go/src/math/big/arith.go:169
	_go_fuzz_dep_.CoverTab[4090]++
						if len(z) == 0 {
//line /usr/local/go/src/math/big/arith.go:170
		_go_fuzz_dep_.CoverTab[4096]++
							return
//line /usr/local/go/src/math/big/arith.go:171
		// _ = "end of CoverTab[4096]"
	} else {
//line /usr/local/go/src/math/big/arith.go:172
		_go_fuzz_dep_.CoverTab[4097]++
//line /usr/local/go/src/math/big/arith.go:172
		// _ = "end of CoverTab[4097]"
//line /usr/local/go/src/math/big/arith.go:172
	}
//line /usr/local/go/src/math/big/arith.go:172
	// _ = "end of CoverTab[4090]"
//line /usr/local/go/src/math/big/arith.go:172
	_go_fuzz_dep_.CoverTab[4091]++
						if len(x) != len(z) {
//line /usr/local/go/src/math/big/arith.go:173
		_go_fuzz_dep_.CoverTab[4098]++

							panic("len(x) != len(z)")
//line /usr/local/go/src/math/big/arith.go:175
		// _ = "end of CoverTab[4098]"
	} else {
//line /usr/local/go/src/math/big/arith.go:176
		_go_fuzz_dep_.CoverTab[4099]++
//line /usr/local/go/src/math/big/arith.go:176
		// _ = "end of CoverTab[4099]"
//line /usr/local/go/src/math/big/arith.go:176
	}
//line /usr/local/go/src/math/big/arith.go:176
	// _ = "end of CoverTab[4091]"
//line /usr/local/go/src/math/big/arith.go:176
	_go_fuzz_dep_.CoverTab[4092]++
						s &= _W - 1
						ŝ := _W - s
						ŝ &= _W - 1
						c = x[0] << ŝ
						for i := 1; i < len(z); i++ {
//line /usr/local/go/src/math/big/arith.go:181
		_go_fuzz_dep_.CoverTab[4100]++
							z[i-1] = x[i-1]>>s | x[i]<<ŝ
//line /usr/local/go/src/math/big/arith.go:182
		// _ = "end of CoverTab[4100]"
	}
//line /usr/local/go/src/math/big/arith.go:183
	// _ = "end of CoverTab[4092]"
//line /usr/local/go/src/math/big/arith.go:183
	_go_fuzz_dep_.CoverTab[4093]++
						z[len(z)-1] = x[len(z)-1] >> s
						return
//line /usr/local/go/src/math/big/arith.go:185
	// _ = "end of CoverTab[4093]"
}

func mulAddVWW_g(z, x []Word, y, r Word) (c Word) {
//line /usr/local/go/src/math/big/arith.go:188
	_go_fuzz_dep_.CoverTab[4101]++
						c = r

						for i := 0; i < len(z) && func() bool {
//line /usr/local/go/src/math/big/arith.go:191
		_go_fuzz_dep_.CoverTab[4103]++
//line /usr/local/go/src/math/big/arith.go:191
		return i < len(x)
//line /usr/local/go/src/math/big/arith.go:191
		// _ = "end of CoverTab[4103]"
//line /usr/local/go/src/math/big/arith.go:191
	}(); i++ {
//line /usr/local/go/src/math/big/arith.go:191
		_go_fuzz_dep_.CoverTab[4104]++
							c, z[i] = mulAddWWW_g(x[i], y, c)
//line /usr/local/go/src/math/big/arith.go:192
		// _ = "end of CoverTab[4104]"
	}
//line /usr/local/go/src/math/big/arith.go:193
	// _ = "end of CoverTab[4101]"
//line /usr/local/go/src/math/big/arith.go:193
	_go_fuzz_dep_.CoverTab[4102]++
						return
//line /usr/local/go/src/math/big/arith.go:194
	// _ = "end of CoverTab[4102]"
}

func addMulVVW_g(z, x []Word, y Word) (c Word) {
//line /usr/local/go/src/math/big/arith.go:197
	_go_fuzz_dep_.CoverTab[4105]++

						for i := 0; i < len(z) && func() bool {
//line /usr/local/go/src/math/big/arith.go:199
		_go_fuzz_dep_.CoverTab[4107]++
//line /usr/local/go/src/math/big/arith.go:199
		return i < len(x)
//line /usr/local/go/src/math/big/arith.go:199
		// _ = "end of CoverTab[4107]"
//line /usr/local/go/src/math/big/arith.go:199
	}(); i++ {
//line /usr/local/go/src/math/big/arith.go:199
		_go_fuzz_dep_.CoverTab[4108]++
							z1, z0 := mulAddWWW_g(x[i], y, z[i])
							lo, cc := bits.Add(uint(z0), uint(c), 0)
							c, z[i] = Word(cc), Word(lo)
							c += z1
//line /usr/local/go/src/math/big/arith.go:203
		// _ = "end of CoverTab[4108]"
	}
//line /usr/local/go/src/math/big/arith.go:204
	// _ = "end of CoverTab[4105]"
//line /usr/local/go/src/math/big/arith.go:204
	_go_fuzz_dep_.CoverTab[4106]++
						return
//line /usr/local/go/src/math/big/arith.go:205
	// _ = "end of CoverTab[4106]"
}

// q = ( x1 << _W + x0 - r)/y. m = floor(( _B^2 - 1 ) / d - _B). Requiring x1<y.
//line /usr/local/go/src/math/big/arith.go:208
// An approximate reciprocal with a reference to "Improved Division by Invariant Integers
//line /usr/local/go/src/math/big/arith.go:208
// (IEEE Transactions on Computers, 11 Jun. 2010)"
//line /usr/local/go/src/math/big/arith.go:211
func divWW(x1, x0, y, m Word) (q, r Word) {
//line /usr/local/go/src/math/big/arith.go:211
	_go_fuzz_dep_.CoverTab[4109]++
						s := nlz(y)
						if s != 0 {
//line /usr/local/go/src/math/big/arith.go:213
		_go_fuzz_dep_.CoverTab[4113]++
							x1 = x1<<s | x0>>(_W-s)
							x0 <<= s
							y <<= s
//line /usr/local/go/src/math/big/arith.go:216
		// _ = "end of CoverTab[4113]"
	} else {
//line /usr/local/go/src/math/big/arith.go:217
		_go_fuzz_dep_.CoverTab[4114]++
//line /usr/local/go/src/math/big/arith.go:217
		// _ = "end of CoverTab[4114]"
//line /usr/local/go/src/math/big/arith.go:217
	}
//line /usr/local/go/src/math/big/arith.go:217
	// _ = "end of CoverTab[4109]"
//line /usr/local/go/src/math/big/arith.go:217
	_go_fuzz_dep_.CoverTab[4110]++
						d := uint(y)

//line /usr/local/go/src/math/big/arith.go:231
	t1, t0 := bits.Mul(uint(m), uint(x1))
						_, c := bits.Add(t0, uint(x0), 0)
						t1, _ = bits.Add(t1, uint(x1), c)

//line /usr/local/go/src/math/big/arith.go:236
	qq := t1

						dq1, dq0 := bits.Mul(d, qq)
						r0, b := bits.Sub(uint(x0), dq0, 0)
						r1, _ := bits.Sub(uint(x1), dq1, b)

//line /usr/local/go/src/math/big/arith.go:258
	if r1 != 0 {
//line /usr/local/go/src/math/big/arith.go:258
		_go_fuzz_dep_.CoverTab[4115]++
							qq++
							r0 -= d
//line /usr/local/go/src/math/big/arith.go:260
		// _ = "end of CoverTab[4115]"
	} else {
//line /usr/local/go/src/math/big/arith.go:261
		_go_fuzz_dep_.CoverTab[4116]++
//line /usr/local/go/src/math/big/arith.go:261
		// _ = "end of CoverTab[4116]"
//line /usr/local/go/src/math/big/arith.go:261
	}
//line /usr/local/go/src/math/big/arith.go:261
	// _ = "end of CoverTab[4110]"
//line /usr/local/go/src/math/big/arith.go:261
	_go_fuzz_dep_.CoverTab[4111]++

						if r0 >= d {
//line /usr/local/go/src/math/big/arith.go:263
		_go_fuzz_dep_.CoverTab[4117]++
							qq++
							r0 -= d
//line /usr/local/go/src/math/big/arith.go:265
		// _ = "end of CoverTab[4117]"
	} else {
//line /usr/local/go/src/math/big/arith.go:266
		_go_fuzz_dep_.CoverTab[4118]++
//line /usr/local/go/src/math/big/arith.go:266
		// _ = "end of CoverTab[4118]"
//line /usr/local/go/src/math/big/arith.go:266
	}
//line /usr/local/go/src/math/big/arith.go:266
	// _ = "end of CoverTab[4111]"
//line /usr/local/go/src/math/big/arith.go:266
	_go_fuzz_dep_.CoverTab[4112]++
						return Word(qq), Word(r0 >> s)
//line /usr/local/go/src/math/big/arith.go:267
	// _ = "end of CoverTab[4112]"
}

// reciprocalWord return the reciprocal of the divisor. rec = floor(( _B^2 - 1 ) / u - _B). u = d1 << nlz(d1).
func reciprocalWord(d1 Word) Word {
//line /usr/local/go/src/math/big/arith.go:271
	_go_fuzz_dep_.CoverTab[4119]++
						u := uint(d1 << nlz(d1))
						x1 := ^u
						x0 := uint(_M)
						rec, _ := bits.Div(x1, x0, u)
						return Word(rec)
//line /usr/local/go/src/math/big/arith.go:276
	// _ = "end of CoverTab[4119]"
}

//line /usr/local/go/src/math/big/arith.go:277
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/arith.go:277
var _ = _go_fuzz_dep_.CoverTab
