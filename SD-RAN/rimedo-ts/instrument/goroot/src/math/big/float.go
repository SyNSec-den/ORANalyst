// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements multi-precision floating-point numbers.
// Like in the GNU MPFR library (https://www.mpfr.org/), operands
// can be of mixed precision. Unlike MPFR, the rounding mode is
// not specified with each operation, but with each operand. The
// rounding mode of the result operand determines the rounding
// mode of an operation. This is a from-scratch implementation.

//line /usr/local/go/src/math/big/float.go:12
package big

//line /usr/local/go/src/math/big/float.go:12
import (
//line /usr/local/go/src/math/big/float.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/float.go:12
)
//line /usr/local/go/src/math/big/float.go:12
import (
//line /usr/local/go/src/math/big/float.go:12
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/float.go:12
)

import (
	"fmt"
	"math"
	"math/bits"
)

const debugFloat = false

//line /usr/local/go/src/math/big/float.go:65
type Float struct {
	prec	uint32
	mode	RoundingMode
	acc	Accuracy
	form	form
	neg	bool
	mant	nat
	exp	int32
}

//line /usr/local/go/src/math/big/float.go:77
type ErrNaN struct {
	msg string
}

func (err ErrNaN) Error() string {
//line /usr/local/go/src/math/big/float.go:81
	_go_fuzz_dep_.CoverTab[4188]++
						return err.msg
//line /usr/local/go/src/math/big/float.go:82
	// _ = "end of CoverTab[4188]"
}

//line /usr/local/go/src/math/big/float.go:88
func NewFloat(x float64) *Float {
//line /usr/local/go/src/math/big/float.go:88
	_go_fuzz_dep_.CoverTab[4189]++
						if math.IsNaN(x) {
//line /usr/local/go/src/math/big/float.go:89
		_go_fuzz_dep_.CoverTab[4191]++
							panic(ErrNaN{"NewFloat(NaN)"})
//line /usr/local/go/src/math/big/float.go:90
		// _ = "end of CoverTab[4191]"
	} else {
//line /usr/local/go/src/math/big/float.go:91
		_go_fuzz_dep_.CoverTab[4192]++
//line /usr/local/go/src/math/big/float.go:91
		// _ = "end of CoverTab[4192]"
//line /usr/local/go/src/math/big/float.go:91
	}
//line /usr/local/go/src/math/big/float.go:91
	// _ = "end of CoverTab[4189]"
//line /usr/local/go/src/math/big/float.go:91
	_go_fuzz_dep_.CoverTab[4190]++
						return new(Float).SetFloat64(x)
//line /usr/local/go/src/math/big/float.go:92
	// _ = "end of CoverTab[4190]"
}

//line /usr/local/go/src/math/big/float.go:96
const (
	MaxExp	= math.MaxInt32
	MinExp	= math.MinInt32
	MaxPrec	= math.MaxUint32
)

//line /usr/local/go/src/math/big/float.go:120
type form byte

//line /usr/local/go/src/math/big/float.go:123
const (
	zero	form	= iota
	finite
	inf
)

//line /usr/local/go/src/math/big/float.go:132
type RoundingMode byte

//line /usr/local/go/src/math/big/float.go:135
const (
	ToNearestEven	RoundingMode	= iota
	ToNearestAway
	ToZero
	AwayFromZero
	ToNegativeInf
	ToPositiveInf
)

//go:generate stringer -type=RoundingMode

//line /usr/local/go/src/math/big/float.go:148
type Accuracy int8

//line /usr/local/go/src/math/big/float.go:151
const (
	Below	Accuracy	= -1
	Exact	Accuracy	= 0
	Above	Accuracy	= +1
)

//go:generate stringer -type=Accuracy

//line /usr/local/go/src/math/big/float.go:164
func (z *Float) SetPrec(prec uint) *Float {
//line /usr/local/go/src/math/big/float.go:164
	_go_fuzz_dep_.CoverTab[4193]++
						z.acc = Exact

//line /usr/local/go/src/math/big/float.go:168
	if prec == 0 {
//line /usr/local/go/src/math/big/float.go:168
		_go_fuzz_dep_.CoverTab[4197]++
							z.prec = 0
							if z.form == finite {
//line /usr/local/go/src/math/big/float.go:170
			_go_fuzz_dep_.CoverTab[4199]++

								z.acc = makeAcc(z.neg)
								z.form = zero
//line /usr/local/go/src/math/big/float.go:173
			// _ = "end of CoverTab[4199]"
		} else {
//line /usr/local/go/src/math/big/float.go:174
			_go_fuzz_dep_.CoverTab[4200]++
//line /usr/local/go/src/math/big/float.go:174
			// _ = "end of CoverTab[4200]"
//line /usr/local/go/src/math/big/float.go:174
		}
//line /usr/local/go/src/math/big/float.go:174
		// _ = "end of CoverTab[4197]"
//line /usr/local/go/src/math/big/float.go:174
		_go_fuzz_dep_.CoverTab[4198]++
							return z
//line /usr/local/go/src/math/big/float.go:175
		// _ = "end of CoverTab[4198]"
	} else {
//line /usr/local/go/src/math/big/float.go:176
		_go_fuzz_dep_.CoverTab[4201]++
//line /usr/local/go/src/math/big/float.go:176
		// _ = "end of CoverTab[4201]"
//line /usr/local/go/src/math/big/float.go:176
	}
//line /usr/local/go/src/math/big/float.go:176
	// _ = "end of CoverTab[4193]"
//line /usr/local/go/src/math/big/float.go:176
	_go_fuzz_dep_.CoverTab[4194]++

//line /usr/local/go/src/math/big/float.go:179
	if prec > MaxPrec {
//line /usr/local/go/src/math/big/float.go:179
		_go_fuzz_dep_.CoverTab[4202]++
							prec = MaxPrec
//line /usr/local/go/src/math/big/float.go:180
		// _ = "end of CoverTab[4202]"
	} else {
//line /usr/local/go/src/math/big/float.go:181
		_go_fuzz_dep_.CoverTab[4203]++
//line /usr/local/go/src/math/big/float.go:181
		// _ = "end of CoverTab[4203]"
//line /usr/local/go/src/math/big/float.go:181
	}
//line /usr/local/go/src/math/big/float.go:181
	// _ = "end of CoverTab[4194]"
//line /usr/local/go/src/math/big/float.go:181
	_go_fuzz_dep_.CoverTab[4195]++
						old := z.prec
						z.prec = uint32(prec)
						if z.prec < old {
//line /usr/local/go/src/math/big/float.go:184
		_go_fuzz_dep_.CoverTab[4204]++
							z.round(0)
//line /usr/local/go/src/math/big/float.go:185
		// _ = "end of CoverTab[4204]"
	} else {
//line /usr/local/go/src/math/big/float.go:186
		_go_fuzz_dep_.CoverTab[4205]++
//line /usr/local/go/src/math/big/float.go:186
		// _ = "end of CoverTab[4205]"
//line /usr/local/go/src/math/big/float.go:186
	}
//line /usr/local/go/src/math/big/float.go:186
	// _ = "end of CoverTab[4195]"
//line /usr/local/go/src/math/big/float.go:186
	_go_fuzz_dep_.CoverTab[4196]++
						return z
//line /usr/local/go/src/math/big/float.go:187
	// _ = "end of CoverTab[4196]"
}

func makeAcc(above bool) Accuracy {
//line /usr/local/go/src/math/big/float.go:190
	_go_fuzz_dep_.CoverTab[4206]++
						if above {
//line /usr/local/go/src/math/big/float.go:191
		_go_fuzz_dep_.CoverTab[4208]++
							return Above
//line /usr/local/go/src/math/big/float.go:192
		// _ = "end of CoverTab[4208]"
	} else {
//line /usr/local/go/src/math/big/float.go:193
		_go_fuzz_dep_.CoverTab[4209]++
//line /usr/local/go/src/math/big/float.go:193
		// _ = "end of CoverTab[4209]"
//line /usr/local/go/src/math/big/float.go:193
	}
//line /usr/local/go/src/math/big/float.go:193
	// _ = "end of CoverTab[4206]"
//line /usr/local/go/src/math/big/float.go:193
	_go_fuzz_dep_.CoverTab[4207]++
						return Below
//line /usr/local/go/src/math/big/float.go:194
	// _ = "end of CoverTab[4207]"
}

//line /usr/local/go/src/math/big/float.go:200
func (z *Float) SetMode(mode RoundingMode) *Float {
//line /usr/local/go/src/math/big/float.go:200
	_go_fuzz_dep_.CoverTab[4210]++
						z.mode = mode
						z.acc = Exact
						return z
//line /usr/local/go/src/math/big/float.go:203
	// _ = "end of CoverTab[4210]"
}

//line /usr/local/go/src/math/big/float.go:208
func (x *Float) Prec() uint {
//line /usr/local/go/src/math/big/float.go:208
	_go_fuzz_dep_.CoverTab[4211]++
						return uint(x.prec)
//line /usr/local/go/src/math/big/float.go:209
	// _ = "end of CoverTab[4211]"
}

//line /usr/local/go/src/math/big/float.go:215
func (x *Float) MinPrec() uint {
//line /usr/local/go/src/math/big/float.go:215
	_go_fuzz_dep_.CoverTab[4212]++
						if x.form != finite {
//line /usr/local/go/src/math/big/float.go:216
		_go_fuzz_dep_.CoverTab[4214]++
							return 0
//line /usr/local/go/src/math/big/float.go:217
		// _ = "end of CoverTab[4214]"
	} else {
//line /usr/local/go/src/math/big/float.go:218
		_go_fuzz_dep_.CoverTab[4215]++
//line /usr/local/go/src/math/big/float.go:218
		// _ = "end of CoverTab[4215]"
//line /usr/local/go/src/math/big/float.go:218
	}
//line /usr/local/go/src/math/big/float.go:218
	// _ = "end of CoverTab[4212]"
//line /usr/local/go/src/math/big/float.go:218
	_go_fuzz_dep_.CoverTab[4213]++
						return uint(len(x.mant))*_W - x.mant.trailingZeroBits()
//line /usr/local/go/src/math/big/float.go:219
	// _ = "end of CoverTab[4213]"
}

//line /usr/local/go/src/math/big/float.go:223
func (x *Float) Mode() RoundingMode {
//line /usr/local/go/src/math/big/float.go:223
	_go_fuzz_dep_.CoverTab[4216]++
						return x.mode
//line /usr/local/go/src/math/big/float.go:224
	// _ = "end of CoverTab[4216]"
}

//line /usr/local/go/src/math/big/float.go:230
func (x *Float) Acc() Accuracy {
//line /usr/local/go/src/math/big/float.go:230
	_go_fuzz_dep_.CoverTab[4217]++
						return x.acc
//line /usr/local/go/src/math/big/float.go:231
	// _ = "end of CoverTab[4217]"
}

//line /usr/local/go/src/math/big/float.go:239
func (x *Float) Sign() int {
//line /usr/local/go/src/math/big/float.go:239
	_go_fuzz_dep_.CoverTab[4218]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:240
		_go_fuzz_dep_.CoverTab[4222]++
							x.validate()
//line /usr/local/go/src/math/big/float.go:241
		// _ = "end of CoverTab[4222]"
	} else {
//line /usr/local/go/src/math/big/float.go:242
		_go_fuzz_dep_.CoverTab[4223]++
//line /usr/local/go/src/math/big/float.go:242
		// _ = "end of CoverTab[4223]"
//line /usr/local/go/src/math/big/float.go:242
	}
//line /usr/local/go/src/math/big/float.go:242
	// _ = "end of CoverTab[4218]"
//line /usr/local/go/src/math/big/float.go:242
	_go_fuzz_dep_.CoverTab[4219]++
						if x.form == zero {
//line /usr/local/go/src/math/big/float.go:243
		_go_fuzz_dep_.CoverTab[4224]++
							return 0
//line /usr/local/go/src/math/big/float.go:244
		// _ = "end of CoverTab[4224]"
	} else {
//line /usr/local/go/src/math/big/float.go:245
		_go_fuzz_dep_.CoverTab[4225]++
//line /usr/local/go/src/math/big/float.go:245
		// _ = "end of CoverTab[4225]"
//line /usr/local/go/src/math/big/float.go:245
	}
//line /usr/local/go/src/math/big/float.go:245
	// _ = "end of CoverTab[4219]"
//line /usr/local/go/src/math/big/float.go:245
	_go_fuzz_dep_.CoverTab[4220]++
						if x.neg {
//line /usr/local/go/src/math/big/float.go:246
		_go_fuzz_dep_.CoverTab[4226]++
							return -1
//line /usr/local/go/src/math/big/float.go:247
		// _ = "end of CoverTab[4226]"
	} else {
//line /usr/local/go/src/math/big/float.go:248
		_go_fuzz_dep_.CoverTab[4227]++
//line /usr/local/go/src/math/big/float.go:248
		// _ = "end of CoverTab[4227]"
//line /usr/local/go/src/math/big/float.go:248
	}
//line /usr/local/go/src/math/big/float.go:248
	// _ = "end of CoverTab[4220]"
//line /usr/local/go/src/math/big/float.go:248
	_go_fuzz_dep_.CoverTab[4221]++
						return 1
//line /usr/local/go/src/math/big/float.go:249
	// _ = "end of CoverTab[4221]"
}

//line /usr/local/go/src/math/big/float.go:267
func (x *Float) MantExp(mant *Float) (exp int) {
//line /usr/local/go/src/math/big/float.go:267
	_go_fuzz_dep_.CoverTab[4228]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:268
		_go_fuzz_dep_.CoverTab[4232]++
							x.validate()
//line /usr/local/go/src/math/big/float.go:269
		// _ = "end of CoverTab[4232]"
	} else {
//line /usr/local/go/src/math/big/float.go:270
		_go_fuzz_dep_.CoverTab[4233]++
//line /usr/local/go/src/math/big/float.go:270
		// _ = "end of CoverTab[4233]"
//line /usr/local/go/src/math/big/float.go:270
	}
//line /usr/local/go/src/math/big/float.go:270
	// _ = "end of CoverTab[4228]"
//line /usr/local/go/src/math/big/float.go:270
	_go_fuzz_dep_.CoverTab[4229]++
						if x.form == finite {
//line /usr/local/go/src/math/big/float.go:271
		_go_fuzz_dep_.CoverTab[4234]++
							exp = int(x.exp)
//line /usr/local/go/src/math/big/float.go:272
		// _ = "end of CoverTab[4234]"
	} else {
//line /usr/local/go/src/math/big/float.go:273
		_go_fuzz_dep_.CoverTab[4235]++
//line /usr/local/go/src/math/big/float.go:273
		// _ = "end of CoverTab[4235]"
//line /usr/local/go/src/math/big/float.go:273
	}
//line /usr/local/go/src/math/big/float.go:273
	// _ = "end of CoverTab[4229]"
//line /usr/local/go/src/math/big/float.go:273
	_go_fuzz_dep_.CoverTab[4230]++
						if mant != nil {
//line /usr/local/go/src/math/big/float.go:274
		_go_fuzz_dep_.CoverTab[4236]++
							mant.Copy(x)
							if mant.form == finite {
//line /usr/local/go/src/math/big/float.go:276
			_go_fuzz_dep_.CoverTab[4237]++
								mant.exp = 0
//line /usr/local/go/src/math/big/float.go:277
			// _ = "end of CoverTab[4237]"
		} else {
//line /usr/local/go/src/math/big/float.go:278
			_go_fuzz_dep_.CoverTab[4238]++
//line /usr/local/go/src/math/big/float.go:278
			// _ = "end of CoverTab[4238]"
//line /usr/local/go/src/math/big/float.go:278
		}
//line /usr/local/go/src/math/big/float.go:278
		// _ = "end of CoverTab[4236]"
	} else {
//line /usr/local/go/src/math/big/float.go:279
		_go_fuzz_dep_.CoverTab[4239]++
//line /usr/local/go/src/math/big/float.go:279
		// _ = "end of CoverTab[4239]"
//line /usr/local/go/src/math/big/float.go:279
	}
//line /usr/local/go/src/math/big/float.go:279
	// _ = "end of CoverTab[4230]"
//line /usr/local/go/src/math/big/float.go:279
	_go_fuzz_dep_.CoverTab[4231]++
						return
//line /usr/local/go/src/math/big/float.go:280
	// _ = "end of CoverTab[4231]"
}

func (z *Float) setExpAndRound(exp int64, sbit uint) {
//line /usr/local/go/src/math/big/float.go:283
	_go_fuzz_dep_.CoverTab[4240]++
						if exp < MinExp {
//line /usr/local/go/src/math/big/float.go:284
		_go_fuzz_dep_.CoverTab[4243]++

							z.acc = makeAcc(z.neg)
							z.form = zero
							return
//line /usr/local/go/src/math/big/float.go:288
		// _ = "end of CoverTab[4243]"
	} else {
//line /usr/local/go/src/math/big/float.go:289
		_go_fuzz_dep_.CoverTab[4244]++
//line /usr/local/go/src/math/big/float.go:289
		// _ = "end of CoverTab[4244]"
//line /usr/local/go/src/math/big/float.go:289
	}
//line /usr/local/go/src/math/big/float.go:289
	// _ = "end of CoverTab[4240]"
//line /usr/local/go/src/math/big/float.go:289
	_go_fuzz_dep_.CoverTab[4241]++

						if exp > MaxExp {
//line /usr/local/go/src/math/big/float.go:291
		_go_fuzz_dep_.CoverTab[4245]++

							z.acc = makeAcc(!z.neg)
							z.form = inf
							return
//line /usr/local/go/src/math/big/float.go:295
		// _ = "end of CoverTab[4245]"
	} else {
//line /usr/local/go/src/math/big/float.go:296
		_go_fuzz_dep_.CoverTab[4246]++
//line /usr/local/go/src/math/big/float.go:296
		// _ = "end of CoverTab[4246]"
//line /usr/local/go/src/math/big/float.go:296
	}
//line /usr/local/go/src/math/big/float.go:296
	// _ = "end of CoverTab[4241]"
//line /usr/local/go/src/math/big/float.go:296
	_go_fuzz_dep_.CoverTab[4242]++

						z.form = finite
						z.exp = int32(exp)
						z.round(sbit)
//line /usr/local/go/src/math/big/float.go:300
	// _ = "end of CoverTab[4242]"
}

//line /usr/local/go/src/math/big/float.go:320
func (z *Float) SetMantExp(mant *Float, exp int) *Float {
//line /usr/local/go/src/math/big/float.go:320
	_go_fuzz_dep_.CoverTab[4247]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:321
		_go_fuzz_dep_.CoverTab[4250]++
							z.validate()
							mant.validate()
//line /usr/local/go/src/math/big/float.go:323
		// _ = "end of CoverTab[4250]"
	} else {
//line /usr/local/go/src/math/big/float.go:324
		_go_fuzz_dep_.CoverTab[4251]++
//line /usr/local/go/src/math/big/float.go:324
		// _ = "end of CoverTab[4251]"
//line /usr/local/go/src/math/big/float.go:324
	}
//line /usr/local/go/src/math/big/float.go:324
	// _ = "end of CoverTab[4247]"
//line /usr/local/go/src/math/big/float.go:324
	_go_fuzz_dep_.CoverTab[4248]++
						z.Copy(mant)

						if z.form == finite {
//line /usr/local/go/src/math/big/float.go:327
		_go_fuzz_dep_.CoverTab[4252]++

							z.setExpAndRound(int64(z.exp)+int64(exp), 0)
//line /usr/local/go/src/math/big/float.go:329
		// _ = "end of CoverTab[4252]"
	} else {
//line /usr/local/go/src/math/big/float.go:330
		_go_fuzz_dep_.CoverTab[4253]++
//line /usr/local/go/src/math/big/float.go:330
		// _ = "end of CoverTab[4253]"
//line /usr/local/go/src/math/big/float.go:330
	}
//line /usr/local/go/src/math/big/float.go:330
	// _ = "end of CoverTab[4248]"
//line /usr/local/go/src/math/big/float.go:330
	_go_fuzz_dep_.CoverTab[4249]++
						return z
//line /usr/local/go/src/math/big/float.go:331
	// _ = "end of CoverTab[4249]"
}

//line /usr/local/go/src/math/big/float.go:335
func (x *Float) Signbit() bool {
//line /usr/local/go/src/math/big/float.go:335
	_go_fuzz_dep_.CoverTab[4254]++
						return x.neg
//line /usr/local/go/src/math/big/float.go:336
	// _ = "end of CoverTab[4254]"
}

//line /usr/local/go/src/math/big/float.go:340
func (x *Float) IsInf() bool {
//line /usr/local/go/src/math/big/float.go:340
	_go_fuzz_dep_.CoverTab[4255]++
						return x.form == inf
//line /usr/local/go/src/math/big/float.go:341
	// _ = "end of CoverTab[4255]"
}

//line /usr/local/go/src/math/big/float.go:346
func (x *Float) IsInt() bool {
//line /usr/local/go/src/math/big/float.go:346
	_go_fuzz_dep_.CoverTab[4256]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:347
		_go_fuzz_dep_.CoverTab[4260]++
							x.validate()
//line /usr/local/go/src/math/big/float.go:348
		// _ = "end of CoverTab[4260]"
	} else {
//line /usr/local/go/src/math/big/float.go:349
		_go_fuzz_dep_.CoverTab[4261]++
//line /usr/local/go/src/math/big/float.go:349
		// _ = "end of CoverTab[4261]"
//line /usr/local/go/src/math/big/float.go:349
	}
//line /usr/local/go/src/math/big/float.go:349
	// _ = "end of CoverTab[4256]"
//line /usr/local/go/src/math/big/float.go:349
	_go_fuzz_dep_.CoverTab[4257]++

						if x.form != finite {
//line /usr/local/go/src/math/big/float.go:351
		_go_fuzz_dep_.CoverTab[4262]++
							return x.form == zero
//line /usr/local/go/src/math/big/float.go:352
		// _ = "end of CoverTab[4262]"
	} else {
//line /usr/local/go/src/math/big/float.go:353
		_go_fuzz_dep_.CoverTab[4263]++
//line /usr/local/go/src/math/big/float.go:353
		// _ = "end of CoverTab[4263]"
//line /usr/local/go/src/math/big/float.go:353
	}
//line /usr/local/go/src/math/big/float.go:353
	// _ = "end of CoverTab[4257]"
//line /usr/local/go/src/math/big/float.go:353
	_go_fuzz_dep_.CoverTab[4258]++

						if x.exp <= 0 {
//line /usr/local/go/src/math/big/float.go:355
		_go_fuzz_dep_.CoverTab[4264]++
							return false
//line /usr/local/go/src/math/big/float.go:356
		// _ = "end of CoverTab[4264]"
	} else {
//line /usr/local/go/src/math/big/float.go:357
		_go_fuzz_dep_.CoverTab[4265]++
//line /usr/local/go/src/math/big/float.go:357
		// _ = "end of CoverTab[4265]"
//line /usr/local/go/src/math/big/float.go:357
	}
//line /usr/local/go/src/math/big/float.go:357
	// _ = "end of CoverTab[4258]"
//line /usr/local/go/src/math/big/float.go:357
	_go_fuzz_dep_.CoverTab[4259]++

						return x.prec <= uint32(x.exp) || func() bool {
//line /usr/local/go/src/math/big/float.go:359
		_go_fuzz_dep_.CoverTab[4266]++
//line /usr/local/go/src/math/big/float.go:359
		return x.MinPrec() <= uint(x.exp)
//line /usr/local/go/src/math/big/float.go:359
		// _ = "end of CoverTab[4266]"
//line /usr/local/go/src/math/big/float.go:359
	}()
//line /usr/local/go/src/math/big/float.go:359
	// _ = "end of CoverTab[4259]"
}

//line /usr/local/go/src/math/big/float.go:363
func (x *Float) validate() {
//line /usr/local/go/src/math/big/float.go:363
	_go_fuzz_dep_.CoverTab[4267]++
						if !debugFloat {
//line /usr/local/go/src/math/big/float.go:364
		_go_fuzz_dep_.CoverTab[4272]++

							panic("validate called but debugFloat is not set")
//line /usr/local/go/src/math/big/float.go:366
		// _ = "end of CoverTab[4272]"
	} else {
//line /usr/local/go/src/math/big/float.go:367
		_go_fuzz_dep_.CoverTab[4273]++
//line /usr/local/go/src/math/big/float.go:367
		// _ = "end of CoverTab[4273]"
//line /usr/local/go/src/math/big/float.go:367
	}
//line /usr/local/go/src/math/big/float.go:367
	// _ = "end of CoverTab[4267]"
//line /usr/local/go/src/math/big/float.go:367
	_go_fuzz_dep_.CoverTab[4268]++
						if x.form != finite {
//line /usr/local/go/src/math/big/float.go:368
		_go_fuzz_dep_.CoverTab[4274]++
							return
//line /usr/local/go/src/math/big/float.go:369
		// _ = "end of CoverTab[4274]"
	} else {
//line /usr/local/go/src/math/big/float.go:370
		_go_fuzz_dep_.CoverTab[4275]++
//line /usr/local/go/src/math/big/float.go:370
		// _ = "end of CoverTab[4275]"
//line /usr/local/go/src/math/big/float.go:370
	}
//line /usr/local/go/src/math/big/float.go:370
	// _ = "end of CoverTab[4268]"
//line /usr/local/go/src/math/big/float.go:370
	_go_fuzz_dep_.CoverTab[4269]++
						m := len(x.mant)
						if m == 0 {
//line /usr/local/go/src/math/big/float.go:372
		_go_fuzz_dep_.CoverTab[4276]++
							panic("nonzero finite number with empty mantissa")
//line /usr/local/go/src/math/big/float.go:373
		// _ = "end of CoverTab[4276]"
	} else {
//line /usr/local/go/src/math/big/float.go:374
		_go_fuzz_dep_.CoverTab[4277]++
//line /usr/local/go/src/math/big/float.go:374
		// _ = "end of CoverTab[4277]"
//line /usr/local/go/src/math/big/float.go:374
	}
//line /usr/local/go/src/math/big/float.go:374
	// _ = "end of CoverTab[4269]"
//line /usr/local/go/src/math/big/float.go:374
	_go_fuzz_dep_.CoverTab[4270]++
						const msb = 1 << (_W - 1)
						if x.mant[m-1]&msb == 0 {
//line /usr/local/go/src/math/big/float.go:376
		_go_fuzz_dep_.CoverTab[4278]++
							panic(fmt.Sprintf("msb not set in last word %#x of %s", x.mant[m-1], x.Text('p', 0)))
//line /usr/local/go/src/math/big/float.go:377
		// _ = "end of CoverTab[4278]"
	} else {
//line /usr/local/go/src/math/big/float.go:378
		_go_fuzz_dep_.CoverTab[4279]++
//line /usr/local/go/src/math/big/float.go:378
		// _ = "end of CoverTab[4279]"
//line /usr/local/go/src/math/big/float.go:378
	}
//line /usr/local/go/src/math/big/float.go:378
	// _ = "end of CoverTab[4270]"
//line /usr/local/go/src/math/big/float.go:378
	_go_fuzz_dep_.CoverTab[4271]++
						if x.prec == 0 {
//line /usr/local/go/src/math/big/float.go:379
		_go_fuzz_dep_.CoverTab[4280]++
							panic("zero precision finite number")
//line /usr/local/go/src/math/big/float.go:380
		// _ = "end of CoverTab[4280]"
	} else {
//line /usr/local/go/src/math/big/float.go:381
		_go_fuzz_dep_.CoverTab[4281]++
//line /usr/local/go/src/math/big/float.go:381
		// _ = "end of CoverTab[4281]"
//line /usr/local/go/src/math/big/float.go:381
	}
//line /usr/local/go/src/math/big/float.go:381
	// _ = "end of CoverTab[4271]"
}

//line /usr/local/go/src/math/big/float.go:392
func (z *Float) round(sbit uint) {
//line /usr/local/go/src/math/big/float.go:392
	_go_fuzz_dep_.CoverTab[4282]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:393
		_go_fuzz_dep_.CoverTab[4289]++
							z.validate()
//line /usr/local/go/src/math/big/float.go:394
		// _ = "end of CoverTab[4289]"
	} else {
//line /usr/local/go/src/math/big/float.go:395
		_go_fuzz_dep_.CoverTab[4290]++
//line /usr/local/go/src/math/big/float.go:395
		// _ = "end of CoverTab[4290]"
//line /usr/local/go/src/math/big/float.go:395
	}
//line /usr/local/go/src/math/big/float.go:395
	// _ = "end of CoverTab[4282]"
//line /usr/local/go/src/math/big/float.go:395
	_go_fuzz_dep_.CoverTab[4283]++

						z.acc = Exact
						if z.form != finite {
//line /usr/local/go/src/math/big/float.go:398
		_go_fuzz_dep_.CoverTab[4291]++

							return
//line /usr/local/go/src/math/big/float.go:400
		// _ = "end of CoverTab[4291]"
	} else {
//line /usr/local/go/src/math/big/float.go:401
		_go_fuzz_dep_.CoverTab[4292]++
//line /usr/local/go/src/math/big/float.go:401
		// _ = "end of CoverTab[4292]"
//line /usr/local/go/src/math/big/float.go:401
	}
//line /usr/local/go/src/math/big/float.go:401
	// _ = "end of CoverTab[4283]"
//line /usr/local/go/src/math/big/float.go:401
	_go_fuzz_dep_.CoverTab[4284]++

//line /usr/local/go/src/math/big/float.go:405
	m := uint32(len(z.mant))
	bits := m * _W
	if bits <= z.prec {
//line /usr/local/go/src/math/big/float.go:407
		_go_fuzz_dep_.CoverTab[4293]++

							return
//line /usr/local/go/src/math/big/float.go:409
		// _ = "end of CoverTab[4293]"
	} else {
//line /usr/local/go/src/math/big/float.go:410
		_go_fuzz_dep_.CoverTab[4294]++
//line /usr/local/go/src/math/big/float.go:410
		// _ = "end of CoverTab[4294]"
//line /usr/local/go/src/math/big/float.go:410
	}
//line /usr/local/go/src/math/big/float.go:410
	// _ = "end of CoverTab[4284]"
//line /usr/local/go/src/math/big/float.go:410
	_go_fuzz_dep_.CoverTab[4285]++

//line /usr/local/go/src/math/big/float.go:426
	r := uint(bits - z.prec - 1)
						rbit := z.mant.bit(r) & 1

//line /usr/local/go/src/math/big/float.go:430
	if sbit == 0 && func() bool {
//line /usr/local/go/src/math/big/float.go:430
		_go_fuzz_dep_.CoverTab[4295]++
//line /usr/local/go/src/math/big/float.go:430
		return (rbit == 0 || func() bool {
//line /usr/local/go/src/math/big/float.go:430
			_go_fuzz_dep_.CoverTab[4296]++
//line /usr/local/go/src/math/big/float.go:430
			return z.mode == ToNearestEven
//line /usr/local/go/src/math/big/float.go:430
			// _ = "end of CoverTab[4296]"
//line /usr/local/go/src/math/big/float.go:430
		}())
//line /usr/local/go/src/math/big/float.go:430
		// _ = "end of CoverTab[4295]"
//line /usr/local/go/src/math/big/float.go:430
	}() {
//line /usr/local/go/src/math/big/float.go:430
		_go_fuzz_dep_.CoverTab[4297]++
							sbit = z.mant.sticky(r)
//line /usr/local/go/src/math/big/float.go:431
		// _ = "end of CoverTab[4297]"
	} else {
//line /usr/local/go/src/math/big/float.go:432
		_go_fuzz_dep_.CoverTab[4298]++
//line /usr/local/go/src/math/big/float.go:432
		// _ = "end of CoverTab[4298]"
//line /usr/local/go/src/math/big/float.go:432
	}
//line /usr/local/go/src/math/big/float.go:432
	// _ = "end of CoverTab[4285]"
//line /usr/local/go/src/math/big/float.go:432
	_go_fuzz_dep_.CoverTab[4286]++
						sbit &= 1

//line /usr/local/go/src/math/big/float.go:436
	n := (z.prec + (_W - 1)) / _W
	if m > n {
//line /usr/local/go/src/math/big/float.go:437
		_go_fuzz_dep_.CoverTab[4299]++
							copy(z.mant, z.mant[m-n:])
							z.mant = z.mant[:n]
//line /usr/local/go/src/math/big/float.go:439
		// _ = "end of CoverTab[4299]"
	} else {
//line /usr/local/go/src/math/big/float.go:440
		_go_fuzz_dep_.CoverTab[4300]++
//line /usr/local/go/src/math/big/float.go:440
		// _ = "end of CoverTab[4300]"
//line /usr/local/go/src/math/big/float.go:440
	}
//line /usr/local/go/src/math/big/float.go:440
	// _ = "end of CoverTab[4286]"
//line /usr/local/go/src/math/big/float.go:440
	_go_fuzz_dep_.CoverTab[4287]++

//line /usr/local/go/src/math/big/float.go:443
	ntz := n*_W - z.prec
						lsb := Word(1) << ntz

//line /usr/local/go/src/math/big/float.go:447
	if rbit|sbit != 0 {
//line /usr/local/go/src/math/big/float.go:447
		_go_fuzz_dep_.CoverTab[4301]++

//line /usr/local/go/src/math/big/float.go:451
		inc := false
		switch z.mode {
		case ToNegativeInf:
//line /usr/local/go/src/math/big/float.go:453
			_go_fuzz_dep_.CoverTab[4303]++
								inc = z.neg
//line /usr/local/go/src/math/big/float.go:454
			// _ = "end of CoverTab[4303]"
		case ToZero:
//line /usr/local/go/src/math/big/float.go:455
			_go_fuzz_dep_.CoverTab[4304]++
//line /usr/local/go/src/math/big/float.go:455
			// _ = "end of CoverTab[4304]"

		case ToNearestEven:
//line /usr/local/go/src/math/big/float.go:457
			_go_fuzz_dep_.CoverTab[4305]++
								inc = rbit != 0 && func() bool {
//line /usr/local/go/src/math/big/float.go:458
				_go_fuzz_dep_.CoverTab[4310]++
//line /usr/local/go/src/math/big/float.go:458
				return (sbit != 0 || func() bool {
//line /usr/local/go/src/math/big/float.go:458
					_go_fuzz_dep_.CoverTab[4311]++
//line /usr/local/go/src/math/big/float.go:458
					return z.mant[0]&lsb != 0
//line /usr/local/go/src/math/big/float.go:458
					// _ = "end of CoverTab[4311]"
//line /usr/local/go/src/math/big/float.go:458
				}())
//line /usr/local/go/src/math/big/float.go:458
				// _ = "end of CoverTab[4310]"
//line /usr/local/go/src/math/big/float.go:458
			}()
//line /usr/local/go/src/math/big/float.go:458
			// _ = "end of CoverTab[4305]"
		case ToNearestAway:
//line /usr/local/go/src/math/big/float.go:459
			_go_fuzz_dep_.CoverTab[4306]++
								inc = rbit != 0
//line /usr/local/go/src/math/big/float.go:460
			// _ = "end of CoverTab[4306]"
		case AwayFromZero:
//line /usr/local/go/src/math/big/float.go:461
			_go_fuzz_dep_.CoverTab[4307]++
								inc = true
//line /usr/local/go/src/math/big/float.go:462
			// _ = "end of CoverTab[4307]"
		case ToPositiveInf:
//line /usr/local/go/src/math/big/float.go:463
			_go_fuzz_dep_.CoverTab[4308]++
								inc = !z.neg
//line /usr/local/go/src/math/big/float.go:464
			// _ = "end of CoverTab[4308]"
		default:
//line /usr/local/go/src/math/big/float.go:465
			_go_fuzz_dep_.CoverTab[4309]++
								panic("unreachable")
//line /usr/local/go/src/math/big/float.go:466
			// _ = "end of CoverTab[4309]"
		}
//line /usr/local/go/src/math/big/float.go:467
		// _ = "end of CoverTab[4301]"
//line /usr/local/go/src/math/big/float.go:467
		_go_fuzz_dep_.CoverTab[4302]++

//line /usr/local/go/src/math/big/float.go:472
		z.acc = makeAcc(inc != z.neg)

		if inc {
//line /usr/local/go/src/math/big/float.go:474
			_go_fuzz_dep_.CoverTab[4312]++

								if addVW(z.mant, z.mant, lsb) != 0 {
//line /usr/local/go/src/math/big/float.go:476
				_go_fuzz_dep_.CoverTab[4313]++

									if z.exp >= MaxExp {
//line /usr/local/go/src/math/big/float.go:478
					_go_fuzz_dep_.CoverTab[4315]++

										z.form = inf
										return
//line /usr/local/go/src/math/big/float.go:481
					// _ = "end of CoverTab[4315]"
				} else {
//line /usr/local/go/src/math/big/float.go:482
					_go_fuzz_dep_.CoverTab[4316]++
//line /usr/local/go/src/math/big/float.go:482
					// _ = "end of CoverTab[4316]"
//line /usr/local/go/src/math/big/float.go:482
				}
//line /usr/local/go/src/math/big/float.go:482
				// _ = "end of CoverTab[4313]"
//line /usr/local/go/src/math/big/float.go:482
				_go_fuzz_dep_.CoverTab[4314]++
									z.exp++

									shrVU(z.mant, z.mant, 1)

									const msb = 1 << (_W - 1)
									z.mant[n-1] |= msb
//line /usr/local/go/src/math/big/float.go:488
				// _ = "end of CoverTab[4314]"
			} else {
//line /usr/local/go/src/math/big/float.go:489
				_go_fuzz_dep_.CoverTab[4317]++
//line /usr/local/go/src/math/big/float.go:489
				// _ = "end of CoverTab[4317]"
//line /usr/local/go/src/math/big/float.go:489
			}
//line /usr/local/go/src/math/big/float.go:489
			// _ = "end of CoverTab[4312]"
		} else {
//line /usr/local/go/src/math/big/float.go:490
			_go_fuzz_dep_.CoverTab[4318]++
//line /usr/local/go/src/math/big/float.go:490
			// _ = "end of CoverTab[4318]"
//line /usr/local/go/src/math/big/float.go:490
		}
//line /usr/local/go/src/math/big/float.go:490
		// _ = "end of CoverTab[4302]"
	} else {
//line /usr/local/go/src/math/big/float.go:491
		_go_fuzz_dep_.CoverTab[4319]++
//line /usr/local/go/src/math/big/float.go:491
		// _ = "end of CoverTab[4319]"
//line /usr/local/go/src/math/big/float.go:491
	}
//line /usr/local/go/src/math/big/float.go:491
	// _ = "end of CoverTab[4287]"
//line /usr/local/go/src/math/big/float.go:491
	_go_fuzz_dep_.CoverTab[4288]++

//line /usr/local/go/src/math/big/float.go:494
	z.mant[0] &^= lsb - 1

	if debugFloat {
//line /usr/local/go/src/math/big/float.go:496
		_go_fuzz_dep_.CoverTab[4320]++
							z.validate()
//line /usr/local/go/src/math/big/float.go:497
		// _ = "end of CoverTab[4320]"
	} else {
//line /usr/local/go/src/math/big/float.go:498
		_go_fuzz_dep_.CoverTab[4321]++
//line /usr/local/go/src/math/big/float.go:498
		// _ = "end of CoverTab[4321]"
//line /usr/local/go/src/math/big/float.go:498
	}
//line /usr/local/go/src/math/big/float.go:498
	// _ = "end of CoverTab[4288]"
}

func (z *Float) setBits64(neg bool, x uint64) *Float {
//line /usr/local/go/src/math/big/float.go:501
	_go_fuzz_dep_.CoverTab[4322]++
						if z.prec == 0 {
//line /usr/local/go/src/math/big/float.go:502
		_go_fuzz_dep_.CoverTab[4326]++
							z.prec = 64
//line /usr/local/go/src/math/big/float.go:503
		// _ = "end of CoverTab[4326]"
	} else {
//line /usr/local/go/src/math/big/float.go:504
		_go_fuzz_dep_.CoverTab[4327]++
//line /usr/local/go/src/math/big/float.go:504
		// _ = "end of CoverTab[4327]"
//line /usr/local/go/src/math/big/float.go:504
	}
//line /usr/local/go/src/math/big/float.go:504
	// _ = "end of CoverTab[4322]"
//line /usr/local/go/src/math/big/float.go:504
	_go_fuzz_dep_.CoverTab[4323]++
						z.acc = Exact
						z.neg = neg
						if x == 0 {
//line /usr/local/go/src/math/big/float.go:507
		_go_fuzz_dep_.CoverTab[4328]++
							z.form = zero
							return z
//line /usr/local/go/src/math/big/float.go:509
		// _ = "end of CoverTab[4328]"
	} else {
//line /usr/local/go/src/math/big/float.go:510
		_go_fuzz_dep_.CoverTab[4329]++
//line /usr/local/go/src/math/big/float.go:510
		// _ = "end of CoverTab[4329]"
//line /usr/local/go/src/math/big/float.go:510
	}
//line /usr/local/go/src/math/big/float.go:510
	// _ = "end of CoverTab[4323]"
//line /usr/local/go/src/math/big/float.go:510
	_go_fuzz_dep_.CoverTab[4324]++

						z.form = finite
						s := bits.LeadingZeros64(x)
						z.mant = z.mant.setUint64(x << uint(s))
						z.exp = int32(64 - s)
						if z.prec < 64 {
//line /usr/local/go/src/math/big/float.go:516
		_go_fuzz_dep_.CoverTab[4330]++
							z.round(0)
//line /usr/local/go/src/math/big/float.go:517
		// _ = "end of CoverTab[4330]"
	} else {
//line /usr/local/go/src/math/big/float.go:518
		_go_fuzz_dep_.CoverTab[4331]++
//line /usr/local/go/src/math/big/float.go:518
		// _ = "end of CoverTab[4331]"
//line /usr/local/go/src/math/big/float.go:518
	}
//line /usr/local/go/src/math/big/float.go:518
	// _ = "end of CoverTab[4324]"
//line /usr/local/go/src/math/big/float.go:518
	_go_fuzz_dep_.CoverTab[4325]++
						return z
//line /usr/local/go/src/math/big/float.go:519
	// _ = "end of CoverTab[4325]"
}

//line /usr/local/go/src/math/big/float.go:525
func (z *Float) SetUint64(x uint64) *Float {
//line /usr/local/go/src/math/big/float.go:525
	_go_fuzz_dep_.CoverTab[4332]++
						return z.setBits64(false, x)
//line /usr/local/go/src/math/big/float.go:526
	// _ = "end of CoverTab[4332]"
}

//line /usr/local/go/src/math/big/float.go:532
func (z *Float) SetInt64(x int64) *Float {
//line /usr/local/go/src/math/big/float.go:532
	_go_fuzz_dep_.CoverTab[4333]++
						u := x
						if u < 0 {
//line /usr/local/go/src/math/big/float.go:534
		_go_fuzz_dep_.CoverTab[4335]++
							u = -u
//line /usr/local/go/src/math/big/float.go:535
		// _ = "end of CoverTab[4335]"
	} else {
//line /usr/local/go/src/math/big/float.go:536
		_go_fuzz_dep_.CoverTab[4336]++
//line /usr/local/go/src/math/big/float.go:536
		// _ = "end of CoverTab[4336]"
//line /usr/local/go/src/math/big/float.go:536
	}
//line /usr/local/go/src/math/big/float.go:536
	// _ = "end of CoverTab[4333]"
//line /usr/local/go/src/math/big/float.go:536
	_go_fuzz_dep_.CoverTab[4334]++

//line /usr/local/go/src/math/big/float.go:539
	return z.setBits64(x < 0, uint64(u))
//line /usr/local/go/src/math/big/float.go:539
	// _ = "end of CoverTab[4334]"
}

//line /usr/local/go/src/math/big/float.go:545
func (z *Float) SetFloat64(x float64) *Float {
//line /usr/local/go/src/math/big/float.go:545
	_go_fuzz_dep_.CoverTab[4337]++
						if z.prec == 0 {
//line /usr/local/go/src/math/big/float.go:546
		_go_fuzz_dep_.CoverTab[4343]++
							z.prec = 53
//line /usr/local/go/src/math/big/float.go:547
		// _ = "end of CoverTab[4343]"
	} else {
//line /usr/local/go/src/math/big/float.go:548
		_go_fuzz_dep_.CoverTab[4344]++
//line /usr/local/go/src/math/big/float.go:548
		// _ = "end of CoverTab[4344]"
//line /usr/local/go/src/math/big/float.go:548
	}
//line /usr/local/go/src/math/big/float.go:548
	// _ = "end of CoverTab[4337]"
//line /usr/local/go/src/math/big/float.go:548
	_go_fuzz_dep_.CoverTab[4338]++
						if math.IsNaN(x) {
//line /usr/local/go/src/math/big/float.go:549
		_go_fuzz_dep_.CoverTab[4345]++
							panic(ErrNaN{"Float.SetFloat64(NaN)"})
//line /usr/local/go/src/math/big/float.go:550
		// _ = "end of CoverTab[4345]"
	} else {
//line /usr/local/go/src/math/big/float.go:551
		_go_fuzz_dep_.CoverTab[4346]++
//line /usr/local/go/src/math/big/float.go:551
		// _ = "end of CoverTab[4346]"
//line /usr/local/go/src/math/big/float.go:551
	}
//line /usr/local/go/src/math/big/float.go:551
	// _ = "end of CoverTab[4338]"
//line /usr/local/go/src/math/big/float.go:551
	_go_fuzz_dep_.CoverTab[4339]++
						z.acc = Exact
						z.neg = math.Signbit(x)
						if x == 0 {
//line /usr/local/go/src/math/big/float.go:554
		_go_fuzz_dep_.CoverTab[4347]++
							z.form = zero
							return z
//line /usr/local/go/src/math/big/float.go:556
		// _ = "end of CoverTab[4347]"
	} else {
//line /usr/local/go/src/math/big/float.go:557
		_go_fuzz_dep_.CoverTab[4348]++
//line /usr/local/go/src/math/big/float.go:557
		// _ = "end of CoverTab[4348]"
//line /usr/local/go/src/math/big/float.go:557
	}
//line /usr/local/go/src/math/big/float.go:557
	// _ = "end of CoverTab[4339]"
//line /usr/local/go/src/math/big/float.go:557
	_go_fuzz_dep_.CoverTab[4340]++
						if math.IsInf(x, 0) {
//line /usr/local/go/src/math/big/float.go:558
		_go_fuzz_dep_.CoverTab[4349]++
							z.form = inf
							return z
//line /usr/local/go/src/math/big/float.go:560
		// _ = "end of CoverTab[4349]"
	} else {
//line /usr/local/go/src/math/big/float.go:561
		_go_fuzz_dep_.CoverTab[4350]++
//line /usr/local/go/src/math/big/float.go:561
		// _ = "end of CoverTab[4350]"
//line /usr/local/go/src/math/big/float.go:561
	}
//line /usr/local/go/src/math/big/float.go:561
	// _ = "end of CoverTab[4340]"
//line /usr/local/go/src/math/big/float.go:561
	_go_fuzz_dep_.CoverTab[4341]++

						z.form = finite
						fmant, exp := math.Frexp(x)
						z.mant = z.mant.setUint64(1<<63 | math.Float64bits(fmant)<<11)
						z.exp = int32(exp)
						if z.prec < 53 {
//line /usr/local/go/src/math/big/float.go:567
		_go_fuzz_dep_.CoverTab[4351]++
							z.round(0)
//line /usr/local/go/src/math/big/float.go:568
		// _ = "end of CoverTab[4351]"
	} else {
//line /usr/local/go/src/math/big/float.go:569
		_go_fuzz_dep_.CoverTab[4352]++
//line /usr/local/go/src/math/big/float.go:569
		// _ = "end of CoverTab[4352]"
//line /usr/local/go/src/math/big/float.go:569
	}
//line /usr/local/go/src/math/big/float.go:569
	// _ = "end of CoverTab[4341]"
//line /usr/local/go/src/math/big/float.go:569
	_go_fuzz_dep_.CoverTab[4342]++
						return z
//line /usr/local/go/src/math/big/float.go:570
	// _ = "end of CoverTab[4342]"
}

//line /usr/local/go/src/math/big/float.go:576
func fnorm(m nat) int64 {
//line /usr/local/go/src/math/big/float.go:576
	_go_fuzz_dep_.CoverTab[4353]++
						if debugFloat && func() bool {
//line /usr/local/go/src/math/big/float.go:577
		_go_fuzz_dep_.CoverTab[4356]++
//line /usr/local/go/src/math/big/float.go:577
		return (len(m) == 0 || func() bool {
//line /usr/local/go/src/math/big/float.go:577
			_go_fuzz_dep_.CoverTab[4357]++
//line /usr/local/go/src/math/big/float.go:577
			return m[len(m)-1] == 0
//line /usr/local/go/src/math/big/float.go:577
			// _ = "end of CoverTab[4357]"
//line /usr/local/go/src/math/big/float.go:577
		}())
//line /usr/local/go/src/math/big/float.go:577
		// _ = "end of CoverTab[4356]"
//line /usr/local/go/src/math/big/float.go:577
	}() {
//line /usr/local/go/src/math/big/float.go:577
		_go_fuzz_dep_.CoverTab[4358]++
							panic("msw of mantissa is 0")
//line /usr/local/go/src/math/big/float.go:578
		// _ = "end of CoverTab[4358]"
	} else {
//line /usr/local/go/src/math/big/float.go:579
		_go_fuzz_dep_.CoverTab[4359]++
//line /usr/local/go/src/math/big/float.go:579
		// _ = "end of CoverTab[4359]"
//line /usr/local/go/src/math/big/float.go:579
	}
//line /usr/local/go/src/math/big/float.go:579
	// _ = "end of CoverTab[4353]"
//line /usr/local/go/src/math/big/float.go:579
	_go_fuzz_dep_.CoverTab[4354]++
						s := nlz(m[len(m)-1])
						if s > 0 {
//line /usr/local/go/src/math/big/float.go:581
		_go_fuzz_dep_.CoverTab[4360]++
							c := shlVU(m, m, s)
							if debugFloat && func() bool {
//line /usr/local/go/src/math/big/float.go:583
			_go_fuzz_dep_.CoverTab[4361]++
//line /usr/local/go/src/math/big/float.go:583
			return c != 0
//line /usr/local/go/src/math/big/float.go:583
			// _ = "end of CoverTab[4361]"
//line /usr/local/go/src/math/big/float.go:583
		}() {
//line /usr/local/go/src/math/big/float.go:583
			_go_fuzz_dep_.CoverTab[4362]++
								panic("nlz or shlVU incorrect")
//line /usr/local/go/src/math/big/float.go:584
			// _ = "end of CoverTab[4362]"
		} else {
//line /usr/local/go/src/math/big/float.go:585
			_go_fuzz_dep_.CoverTab[4363]++
//line /usr/local/go/src/math/big/float.go:585
			// _ = "end of CoverTab[4363]"
//line /usr/local/go/src/math/big/float.go:585
		}
//line /usr/local/go/src/math/big/float.go:585
		// _ = "end of CoverTab[4360]"
	} else {
//line /usr/local/go/src/math/big/float.go:586
		_go_fuzz_dep_.CoverTab[4364]++
//line /usr/local/go/src/math/big/float.go:586
		// _ = "end of CoverTab[4364]"
//line /usr/local/go/src/math/big/float.go:586
	}
//line /usr/local/go/src/math/big/float.go:586
	// _ = "end of CoverTab[4354]"
//line /usr/local/go/src/math/big/float.go:586
	_go_fuzz_dep_.CoverTab[4355]++
						return int64(s)
//line /usr/local/go/src/math/big/float.go:587
	// _ = "end of CoverTab[4355]"
}

//line /usr/local/go/src/math/big/float.go:593
func (z *Float) SetInt(x *Int) *Float {
//line /usr/local/go/src/math/big/float.go:593
	_go_fuzz_dep_.CoverTab[4365]++

//line /usr/local/go/src/math/big/float.go:597
	bits := uint32(x.BitLen())
	if z.prec == 0 {
//line /usr/local/go/src/math/big/float.go:598
		_go_fuzz_dep_.CoverTab[4368]++
							z.prec = umax32(bits, 64)
//line /usr/local/go/src/math/big/float.go:599
		// _ = "end of CoverTab[4368]"
	} else {
//line /usr/local/go/src/math/big/float.go:600
		_go_fuzz_dep_.CoverTab[4369]++
//line /usr/local/go/src/math/big/float.go:600
		// _ = "end of CoverTab[4369]"
//line /usr/local/go/src/math/big/float.go:600
	}
//line /usr/local/go/src/math/big/float.go:600
	// _ = "end of CoverTab[4365]"
//line /usr/local/go/src/math/big/float.go:600
	_go_fuzz_dep_.CoverTab[4366]++
						z.acc = Exact
						z.neg = x.neg
						if len(x.abs) == 0 {
//line /usr/local/go/src/math/big/float.go:603
		_go_fuzz_dep_.CoverTab[4370]++
							z.form = zero
							return z
//line /usr/local/go/src/math/big/float.go:605
		// _ = "end of CoverTab[4370]"
	} else {
//line /usr/local/go/src/math/big/float.go:606
		_go_fuzz_dep_.CoverTab[4371]++
//line /usr/local/go/src/math/big/float.go:606
		// _ = "end of CoverTab[4371]"
//line /usr/local/go/src/math/big/float.go:606
	}
//line /usr/local/go/src/math/big/float.go:606
	// _ = "end of CoverTab[4366]"
//line /usr/local/go/src/math/big/float.go:606
	_go_fuzz_dep_.CoverTab[4367]++

						z.mant = z.mant.set(x.abs)
						fnorm(z.mant)
						z.setExpAndRound(int64(bits), 0)
						return z
//line /usr/local/go/src/math/big/float.go:611
	// _ = "end of CoverTab[4367]"
}

//line /usr/local/go/src/math/big/float.go:617
func (z *Float) SetRat(x *Rat) *Float {
//line /usr/local/go/src/math/big/float.go:617
	_go_fuzz_dep_.CoverTab[4372]++
						if x.IsInt() {
//line /usr/local/go/src/math/big/float.go:618
		_go_fuzz_dep_.CoverTab[4375]++
							return z.SetInt(x.Num())
//line /usr/local/go/src/math/big/float.go:619
		// _ = "end of CoverTab[4375]"
	} else {
//line /usr/local/go/src/math/big/float.go:620
		_go_fuzz_dep_.CoverTab[4376]++
//line /usr/local/go/src/math/big/float.go:620
		// _ = "end of CoverTab[4376]"
//line /usr/local/go/src/math/big/float.go:620
	}
//line /usr/local/go/src/math/big/float.go:620
	// _ = "end of CoverTab[4372]"
//line /usr/local/go/src/math/big/float.go:620
	_go_fuzz_dep_.CoverTab[4373]++
						var a, b Float
						a.SetInt(x.Num())
						b.SetInt(x.Denom())
						if z.prec == 0 {
//line /usr/local/go/src/math/big/float.go:624
		_go_fuzz_dep_.CoverTab[4377]++
							z.prec = umax32(a.prec, b.prec)
//line /usr/local/go/src/math/big/float.go:625
		// _ = "end of CoverTab[4377]"
	} else {
//line /usr/local/go/src/math/big/float.go:626
		_go_fuzz_dep_.CoverTab[4378]++
//line /usr/local/go/src/math/big/float.go:626
		// _ = "end of CoverTab[4378]"
//line /usr/local/go/src/math/big/float.go:626
	}
//line /usr/local/go/src/math/big/float.go:626
	// _ = "end of CoverTab[4373]"
//line /usr/local/go/src/math/big/float.go:626
	_go_fuzz_dep_.CoverTab[4374]++
						return z.Quo(&a, &b)
//line /usr/local/go/src/math/big/float.go:627
	// _ = "end of CoverTab[4374]"
}

//line /usr/local/go/src/math/big/float.go:634
func (z *Float) SetInf(signbit bool) *Float {
//line /usr/local/go/src/math/big/float.go:634
	_go_fuzz_dep_.CoverTab[4379]++
						z.acc = Exact
						z.form = inf
						z.neg = signbit
						return z
//line /usr/local/go/src/math/big/float.go:638
	// _ = "end of CoverTab[4379]"
}

//line /usr/local/go/src/math/big/float.go:647
func (z *Float) Set(x *Float) *Float {
//line /usr/local/go/src/math/big/float.go:647
	_go_fuzz_dep_.CoverTab[4380]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:648
		_go_fuzz_dep_.CoverTab[4383]++
							x.validate()
//line /usr/local/go/src/math/big/float.go:649
		// _ = "end of CoverTab[4383]"
	} else {
//line /usr/local/go/src/math/big/float.go:650
		_go_fuzz_dep_.CoverTab[4384]++
//line /usr/local/go/src/math/big/float.go:650
		// _ = "end of CoverTab[4384]"
//line /usr/local/go/src/math/big/float.go:650
	}
//line /usr/local/go/src/math/big/float.go:650
	// _ = "end of CoverTab[4380]"
//line /usr/local/go/src/math/big/float.go:650
	_go_fuzz_dep_.CoverTab[4381]++
						z.acc = Exact
						if z != x {
//line /usr/local/go/src/math/big/float.go:652
		_go_fuzz_dep_.CoverTab[4385]++
							z.form = x.form
							z.neg = x.neg
							if x.form == finite {
//line /usr/local/go/src/math/big/float.go:655
			_go_fuzz_dep_.CoverTab[4387]++
								z.exp = x.exp
								z.mant = z.mant.set(x.mant)
//line /usr/local/go/src/math/big/float.go:657
			// _ = "end of CoverTab[4387]"
		} else {
//line /usr/local/go/src/math/big/float.go:658
			_go_fuzz_dep_.CoverTab[4388]++
//line /usr/local/go/src/math/big/float.go:658
			// _ = "end of CoverTab[4388]"
//line /usr/local/go/src/math/big/float.go:658
		}
//line /usr/local/go/src/math/big/float.go:658
		// _ = "end of CoverTab[4385]"
//line /usr/local/go/src/math/big/float.go:658
		_go_fuzz_dep_.CoverTab[4386]++
							if z.prec == 0 {
//line /usr/local/go/src/math/big/float.go:659
			_go_fuzz_dep_.CoverTab[4389]++
								z.prec = x.prec
//line /usr/local/go/src/math/big/float.go:660
			// _ = "end of CoverTab[4389]"
		} else {
//line /usr/local/go/src/math/big/float.go:661
			_go_fuzz_dep_.CoverTab[4390]++
//line /usr/local/go/src/math/big/float.go:661
			if z.prec < x.prec {
//line /usr/local/go/src/math/big/float.go:661
				_go_fuzz_dep_.CoverTab[4391]++
									z.round(0)
//line /usr/local/go/src/math/big/float.go:662
				// _ = "end of CoverTab[4391]"
			} else {
//line /usr/local/go/src/math/big/float.go:663
				_go_fuzz_dep_.CoverTab[4392]++
//line /usr/local/go/src/math/big/float.go:663
				// _ = "end of CoverTab[4392]"
//line /usr/local/go/src/math/big/float.go:663
			}
//line /usr/local/go/src/math/big/float.go:663
			// _ = "end of CoverTab[4390]"
//line /usr/local/go/src/math/big/float.go:663
		}
//line /usr/local/go/src/math/big/float.go:663
		// _ = "end of CoverTab[4386]"
	} else {
//line /usr/local/go/src/math/big/float.go:664
		_go_fuzz_dep_.CoverTab[4393]++
//line /usr/local/go/src/math/big/float.go:664
		// _ = "end of CoverTab[4393]"
//line /usr/local/go/src/math/big/float.go:664
	}
//line /usr/local/go/src/math/big/float.go:664
	// _ = "end of CoverTab[4381]"
//line /usr/local/go/src/math/big/float.go:664
	_go_fuzz_dep_.CoverTab[4382]++
						return z
//line /usr/local/go/src/math/big/float.go:665
	// _ = "end of CoverTab[4382]"
}

//line /usr/local/go/src/math/big/float.go:671
func (z *Float) Copy(x *Float) *Float {
//line /usr/local/go/src/math/big/float.go:671
	_go_fuzz_dep_.CoverTab[4394]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:672
		_go_fuzz_dep_.CoverTab[4397]++
							x.validate()
//line /usr/local/go/src/math/big/float.go:673
		// _ = "end of CoverTab[4397]"
	} else {
//line /usr/local/go/src/math/big/float.go:674
		_go_fuzz_dep_.CoverTab[4398]++
//line /usr/local/go/src/math/big/float.go:674
		// _ = "end of CoverTab[4398]"
//line /usr/local/go/src/math/big/float.go:674
	}
//line /usr/local/go/src/math/big/float.go:674
	// _ = "end of CoverTab[4394]"
//line /usr/local/go/src/math/big/float.go:674
	_go_fuzz_dep_.CoverTab[4395]++
						if z != x {
//line /usr/local/go/src/math/big/float.go:675
		_go_fuzz_dep_.CoverTab[4399]++
							z.prec = x.prec
							z.mode = x.mode
							z.acc = x.acc
							z.form = x.form
							z.neg = x.neg
							if z.form == finite {
//line /usr/local/go/src/math/big/float.go:681
			_go_fuzz_dep_.CoverTab[4400]++
								z.mant = z.mant.set(x.mant)
								z.exp = x.exp
//line /usr/local/go/src/math/big/float.go:683
			// _ = "end of CoverTab[4400]"
		} else {
//line /usr/local/go/src/math/big/float.go:684
			_go_fuzz_dep_.CoverTab[4401]++
//line /usr/local/go/src/math/big/float.go:684
			// _ = "end of CoverTab[4401]"
//line /usr/local/go/src/math/big/float.go:684
		}
//line /usr/local/go/src/math/big/float.go:684
		// _ = "end of CoverTab[4399]"
	} else {
//line /usr/local/go/src/math/big/float.go:685
		_go_fuzz_dep_.CoverTab[4402]++
//line /usr/local/go/src/math/big/float.go:685
		// _ = "end of CoverTab[4402]"
//line /usr/local/go/src/math/big/float.go:685
	}
//line /usr/local/go/src/math/big/float.go:685
	// _ = "end of CoverTab[4395]"
//line /usr/local/go/src/math/big/float.go:685
	_go_fuzz_dep_.CoverTab[4396]++
						return z
//line /usr/local/go/src/math/big/float.go:686
	// _ = "end of CoverTab[4396]"
}

//line /usr/local/go/src/math/big/float.go:690
func msb32(x nat) uint32 {
//line /usr/local/go/src/math/big/float.go:690
	_go_fuzz_dep_.CoverTab[4403]++
						i := len(x) - 1
						if i < 0 {
//line /usr/local/go/src/math/big/float.go:692
		_go_fuzz_dep_.CoverTab[4407]++
							return 0
//line /usr/local/go/src/math/big/float.go:693
		// _ = "end of CoverTab[4407]"
	} else {
//line /usr/local/go/src/math/big/float.go:694
		_go_fuzz_dep_.CoverTab[4408]++
//line /usr/local/go/src/math/big/float.go:694
		// _ = "end of CoverTab[4408]"
//line /usr/local/go/src/math/big/float.go:694
	}
//line /usr/local/go/src/math/big/float.go:694
	// _ = "end of CoverTab[4403]"
//line /usr/local/go/src/math/big/float.go:694
	_go_fuzz_dep_.CoverTab[4404]++
						if debugFloat && func() bool {
//line /usr/local/go/src/math/big/float.go:695
		_go_fuzz_dep_.CoverTab[4409]++
//line /usr/local/go/src/math/big/float.go:695
		return x[i]&(1<<(_W-1)) == 0
//line /usr/local/go/src/math/big/float.go:695
		// _ = "end of CoverTab[4409]"
//line /usr/local/go/src/math/big/float.go:695
	}() {
//line /usr/local/go/src/math/big/float.go:695
		_go_fuzz_dep_.CoverTab[4410]++
							panic("x not normalized")
//line /usr/local/go/src/math/big/float.go:696
		// _ = "end of CoverTab[4410]"
	} else {
//line /usr/local/go/src/math/big/float.go:697
		_go_fuzz_dep_.CoverTab[4411]++
//line /usr/local/go/src/math/big/float.go:697
		// _ = "end of CoverTab[4411]"
//line /usr/local/go/src/math/big/float.go:697
	}
//line /usr/local/go/src/math/big/float.go:697
	// _ = "end of CoverTab[4404]"
//line /usr/local/go/src/math/big/float.go:697
	_go_fuzz_dep_.CoverTab[4405]++
						switch _W {
	case 32:
//line /usr/local/go/src/math/big/float.go:699
		_go_fuzz_dep_.CoverTab[4412]++
							return uint32(x[i])
//line /usr/local/go/src/math/big/float.go:700
		// _ = "end of CoverTab[4412]"
	case 64:
//line /usr/local/go/src/math/big/float.go:701
		_go_fuzz_dep_.CoverTab[4413]++
							return uint32(x[i] >> 32)
//line /usr/local/go/src/math/big/float.go:702
		// _ = "end of CoverTab[4413]"
//line /usr/local/go/src/math/big/float.go:702
	default:
//line /usr/local/go/src/math/big/float.go:702
		_go_fuzz_dep_.CoverTab[4414]++
//line /usr/local/go/src/math/big/float.go:702
		// _ = "end of CoverTab[4414]"
	}
//line /usr/local/go/src/math/big/float.go:703
	// _ = "end of CoverTab[4405]"
//line /usr/local/go/src/math/big/float.go:703
	_go_fuzz_dep_.CoverTab[4406]++
						panic("unreachable")
//line /usr/local/go/src/math/big/float.go:704
	// _ = "end of CoverTab[4406]"
}

//line /usr/local/go/src/math/big/float.go:708
func msb64(x nat) uint64 {
//line /usr/local/go/src/math/big/float.go:708
	_go_fuzz_dep_.CoverTab[4415]++
						i := len(x) - 1
						if i < 0 {
//line /usr/local/go/src/math/big/float.go:710
		_go_fuzz_dep_.CoverTab[4419]++
							return 0
//line /usr/local/go/src/math/big/float.go:711
		// _ = "end of CoverTab[4419]"
	} else {
//line /usr/local/go/src/math/big/float.go:712
		_go_fuzz_dep_.CoverTab[4420]++
//line /usr/local/go/src/math/big/float.go:712
		// _ = "end of CoverTab[4420]"
//line /usr/local/go/src/math/big/float.go:712
	}
//line /usr/local/go/src/math/big/float.go:712
	// _ = "end of CoverTab[4415]"
//line /usr/local/go/src/math/big/float.go:712
	_go_fuzz_dep_.CoverTab[4416]++
						if debugFloat && func() bool {
//line /usr/local/go/src/math/big/float.go:713
		_go_fuzz_dep_.CoverTab[4421]++
//line /usr/local/go/src/math/big/float.go:713
		return x[i]&(1<<(_W-1)) == 0
//line /usr/local/go/src/math/big/float.go:713
		// _ = "end of CoverTab[4421]"
//line /usr/local/go/src/math/big/float.go:713
	}() {
//line /usr/local/go/src/math/big/float.go:713
		_go_fuzz_dep_.CoverTab[4422]++
							panic("x not normalized")
//line /usr/local/go/src/math/big/float.go:714
		// _ = "end of CoverTab[4422]"
	} else {
//line /usr/local/go/src/math/big/float.go:715
		_go_fuzz_dep_.CoverTab[4423]++
//line /usr/local/go/src/math/big/float.go:715
		// _ = "end of CoverTab[4423]"
//line /usr/local/go/src/math/big/float.go:715
	}
//line /usr/local/go/src/math/big/float.go:715
	// _ = "end of CoverTab[4416]"
//line /usr/local/go/src/math/big/float.go:715
	_go_fuzz_dep_.CoverTab[4417]++
						switch _W {
	case 32:
//line /usr/local/go/src/math/big/float.go:717
		_go_fuzz_dep_.CoverTab[4424]++
							v := uint64(x[i]) << 32
							if i > 0 {
//line /usr/local/go/src/math/big/float.go:719
			_go_fuzz_dep_.CoverTab[4428]++
								v |= uint64(x[i-1])
//line /usr/local/go/src/math/big/float.go:720
			// _ = "end of CoverTab[4428]"
		} else {
//line /usr/local/go/src/math/big/float.go:721
			_go_fuzz_dep_.CoverTab[4429]++
//line /usr/local/go/src/math/big/float.go:721
			// _ = "end of CoverTab[4429]"
//line /usr/local/go/src/math/big/float.go:721
		}
//line /usr/local/go/src/math/big/float.go:721
		// _ = "end of CoverTab[4424]"
//line /usr/local/go/src/math/big/float.go:721
		_go_fuzz_dep_.CoverTab[4425]++
							return v
//line /usr/local/go/src/math/big/float.go:722
		// _ = "end of CoverTab[4425]"
	case 64:
//line /usr/local/go/src/math/big/float.go:723
		_go_fuzz_dep_.CoverTab[4426]++
							return uint64(x[i])
//line /usr/local/go/src/math/big/float.go:724
		// _ = "end of CoverTab[4426]"
//line /usr/local/go/src/math/big/float.go:724
	default:
//line /usr/local/go/src/math/big/float.go:724
		_go_fuzz_dep_.CoverTab[4427]++
//line /usr/local/go/src/math/big/float.go:724
		// _ = "end of CoverTab[4427]"
	}
//line /usr/local/go/src/math/big/float.go:725
	// _ = "end of CoverTab[4417]"
//line /usr/local/go/src/math/big/float.go:725
	_go_fuzz_dep_.CoverTab[4418]++
						panic("unreachable")
//line /usr/local/go/src/math/big/float.go:726
	// _ = "end of CoverTab[4418]"
}

//line /usr/local/go/src/math/big/float.go:734
func (x *Float) Uint64() (uint64, Accuracy) {
//line /usr/local/go/src/math/big/float.go:734
	_go_fuzz_dep_.CoverTab[4430]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:735
		_go_fuzz_dep_.CoverTab[4433]++
							x.validate()
//line /usr/local/go/src/math/big/float.go:736
		// _ = "end of CoverTab[4433]"
	} else {
//line /usr/local/go/src/math/big/float.go:737
		_go_fuzz_dep_.CoverTab[4434]++
//line /usr/local/go/src/math/big/float.go:737
		// _ = "end of CoverTab[4434]"
//line /usr/local/go/src/math/big/float.go:737
	}
//line /usr/local/go/src/math/big/float.go:737
	// _ = "end of CoverTab[4430]"
//line /usr/local/go/src/math/big/float.go:737
	_go_fuzz_dep_.CoverTab[4431]++

						switch x.form {
	case finite:
//line /usr/local/go/src/math/big/float.go:740
		_go_fuzz_dep_.CoverTab[4435]++
							if x.neg {
//line /usr/local/go/src/math/big/float.go:741
			_go_fuzz_dep_.CoverTab[4443]++
								return 0, Above
//line /usr/local/go/src/math/big/float.go:742
			// _ = "end of CoverTab[4443]"
		} else {
//line /usr/local/go/src/math/big/float.go:743
			_go_fuzz_dep_.CoverTab[4444]++
//line /usr/local/go/src/math/big/float.go:743
			// _ = "end of CoverTab[4444]"
//line /usr/local/go/src/math/big/float.go:743
		}
//line /usr/local/go/src/math/big/float.go:743
		// _ = "end of CoverTab[4435]"
//line /usr/local/go/src/math/big/float.go:743
		_go_fuzz_dep_.CoverTab[4436]++

							if x.exp <= 0 {
//line /usr/local/go/src/math/big/float.go:745
			_go_fuzz_dep_.CoverTab[4445]++

								return 0, Below
//line /usr/local/go/src/math/big/float.go:747
			// _ = "end of CoverTab[4445]"
		} else {
//line /usr/local/go/src/math/big/float.go:748
			_go_fuzz_dep_.CoverTab[4446]++
//line /usr/local/go/src/math/big/float.go:748
			// _ = "end of CoverTab[4446]"
//line /usr/local/go/src/math/big/float.go:748
		}
//line /usr/local/go/src/math/big/float.go:748
		// _ = "end of CoverTab[4436]"
//line /usr/local/go/src/math/big/float.go:748
		_go_fuzz_dep_.CoverTab[4437]++

							if x.exp <= 64 {
//line /usr/local/go/src/math/big/float.go:750
			_go_fuzz_dep_.CoverTab[4447]++

								u := msb64(x.mant) >> (64 - uint32(x.exp))
								if x.MinPrec() <= 64 {
//line /usr/local/go/src/math/big/float.go:753
				_go_fuzz_dep_.CoverTab[4449]++
									return u, Exact
//line /usr/local/go/src/math/big/float.go:754
				// _ = "end of CoverTab[4449]"
			} else {
//line /usr/local/go/src/math/big/float.go:755
				_go_fuzz_dep_.CoverTab[4450]++
//line /usr/local/go/src/math/big/float.go:755
				// _ = "end of CoverTab[4450]"
//line /usr/local/go/src/math/big/float.go:755
			}
//line /usr/local/go/src/math/big/float.go:755
			// _ = "end of CoverTab[4447]"
//line /usr/local/go/src/math/big/float.go:755
			_go_fuzz_dep_.CoverTab[4448]++
								return u, Below
//line /usr/local/go/src/math/big/float.go:756
			// _ = "end of CoverTab[4448]"
		} else {
//line /usr/local/go/src/math/big/float.go:757
			_go_fuzz_dep_.CoverTab[4451]++
//line /usr/local/go/src/math/big/float.go:757
			// _ = "end of CoverTab[4451]"
//line /usr/local/go/src/math/big/float.go:757
		}
//line /usr/local/go/src/math/big/float.go:757
		// _ = "end of CoverTab[4437]"
//line /usr/local/go/src/math/big/float.go:757
		_go_fuzz_dep_.CoverTab[4438]++

							return math.MaxUint64, Below
//line /usr/local/go/src/math/big/float.go:759
		// _ = "end of CoverTab[4438]"

	case zero:
//line /usr/local/go/src/math/big/float.go:761
		_go_fuzz_dep_.CoverTab[4439]++
							return 0, Exact
//line /usr/local/go/src/math/big/float.go:762
		// _ = "end of CoverTab[4439]"

	case inf:
//line /usr/local/go/src/math/big/float.go:764
		_go_fuzz_dep_.CoverTab[4440]++
							if x.neg {
//line /usr/local/go/src/math/big/float.go:765
			_go_fuzz_dep_.CoverTab[4452]++
								return 0, Above
//line /usr/local/go/src/math/big/float.go:766
			// _ = "end of CoverTab[4452]"
		} else {
//line /usr/local/go/src/math/big/float.go:767
			_go_fuzz_dep_.CoverTab[4453]++
//line /usr/local/go/src/math/big/float.go:767
			// _ = "end of CoverTab[4453]"
//line /usr/local/go/src/math/big/float.go:767
		}
//line /usr/local/go/src/math/big/float.go:767
		// _ = "end of CoverTab[4440]"
//line /usr/local/go/src/math/big/float.go:767
		_go_fuzz_dep_.CoverTab[4441]++
							return math.MaxUint64, Below
//line /usr/local/go/src/math/big/float.go:768
		// _ = "end of CoverTab[4441]"
//line /usr/local/go/src/math/big/float.go:768
	default:
//line /usr/local/go/src/math/big/float.go:768
		_go_fuzz_dep_.CoverTab[4442]++
//line /usr/local/go/src/math/big/float.go:768
		// _ = "end of CoverTab[4442]"
	}
//line /usr/local/go/src/math/big/float.go:769
	// _ = "end of CoverTab[4431]"
//line /usr/local/go/src/math/big/float.go:769
	_go_fuzz_dep_.CoverTab[4432]++

						panic("unreachable")
//line /usr/local/go/src/math/big/float.go:771
	// _ = "end of CoverTab[4432]"
}

//line /usr/local/go/src/math/big/float.go:779
func (x *Float) Int64() (int64, Accuracy) {
//line /usr/local/go/src/math/big/float.go:779
	_go_fuzz_dep_.CoverTab[4454]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:780
		_go_fuzz_dep_.CoverTab[4457]++
							x.validate()
//line /usr/local/go/src/math/big/float.go:781
		// _ = "end of CoverTab[4457]"
	} else {
//line /usr/local/go/src/math/big/float.go:782
		_go_fuzz_dep_.CoverTab[4458]++
//line /usr/local/go/src/math/big/float.go:782
		// _ = "end of CoverTab[4458]"
//line /usr/local/go/src/math/big/float.go:782
	}
//line /usr/local/go/src/math/big/float.go:782
	// _ = "end of CoverTab[4454]"
//line /usr/local/go/src/math/big/float.go:782
	_go_fuzz_dep_.CoverTab[4455]++

						switch x.form {
	case finite:
//line /usr/local/go/src/math/big/float.go:785
		_go_fuzz_dep_.CoverTab[4459]++

							acc := makeAcc(x.neg)
							if x.exp <= 0 {
//line /usr/local/go/src/math/big/float.go:788
			_go_fuzz_dep_.CoverTab[4467]++

								return 0, acc
//line /usr/local/go/src/math/big/float.go:790
			// _ = "end of CoverTab[4467]"
		} else {
//line /usr/local/go/src/math/big/float.go:791
			_go_fuzz_dep_.CoverTab[4468]++
//line /usr/local/go/src/math/big/float.go:791
			// _ = "end of CoverTab[4468]"
//line /usr/local/go/src/math/big/float.go:791
		}
//line /usr/local/go/src/math/big/float.go:791
		// _ = "end of CoverTab[4459]"
//line /usr/local/go/src/math/big/float.go:791
		_go_fuzz_dep_.CoverTab[4460]++

//line /usr/local/go/src/math/big/float.go:795
		if x.exp <= 63 {
//line /usr/local/go/src/math/big/float.go:795
			_go_fuzz_dep_.CoverTab[4469]++

								i := int64(msb64(x.mant) >> (64 - uint32(x.exp)))
								if x.neg {
//line /usr/local/go/src/math/big/float.go:798
				_go_fuzz_dep_.CoverTab[4472]++
									i = -i
//line /usr/local/go/src/math/big/float.go:799
				// _ = "end of CoverTab[4472]"
			} else {
//line /usr/local/go/src/math/big/float.go:800
				_go_fuzz_dep_.CoverTab[4473]++
//line /usr/local/go/src/math/big/float.go:800
				// _ = "end of CoverTab[4473]"
//line /usr/local/go/src/math/big/float.go:800
			}
//line /usr/local/go/src/math/big/float.go:800
			// _ = "end of CoverTab[4469]"
//line /usr/local/go/src/math/big/float.go:800
			_go_fuzz_dep_.CoverTab[4470]++
								if x.MinPrec() <= uint(x.exp) {
//line /usr/local/go/src/math/big/float.go:801
				_go_fuzz_dep_.CoverTab[4474]++
									return i, Exact
//line /usr/local/go/src/math/big/float.go:802
				// _ = "end of CoverTab[4474]"
			} else {
//line /usr/local/go/src/math/big/float.go:803
				_go_fuzz_dep_.CoverTab[4475]++
//line /usr/local/go/src/math/big/float.go:803
				// _ = "end of CoverTab[4475]"
//line /usr/local/go/src/math/big/float.go:803
			}
//line /usr/local/go/src/math/big/float.go:803
			// _ = "end of CoverTab[4470]"
//line /usr/local/go/src/math/big/float.go:803
			_go_fuzz_dep_.CoverTab[4471]++
								return i, acc
//line /usr/local/go/src/math/big/float.go:804
			// _ = "end of CoverTab[4471]"
		} else {
//line /usr/local/go/src/math/big/float.go:805
			_go_fuzz_dep_.CoverTab[4476]++
//line /usr/local/go/src/math/big/float.go:805
			// _ = "end of CoverTab[4476]"
//line /usr/local/go/src/math/big/float.go:805
		}
//line /usr/local/go/src/math/big/float.go:805
		// _ = "end of CoverTab[4460]"
//line /usr/local/go/src/math/big/float.go:805
		_go_fuzz_dep_.CoverTab[4461]++
							if x.neg {
//line /usr/local/go/src/math/big/float.go:806
			_go_fuzz_dep_.CoverTab[4477]++

								if x.exp == 64 && func() bool {
//line /usr/local/go/src/math/big/float.go:808
				_go_fuzz_dep_.CoverTab[4479]++
//line /usr/local/go/src/math/big/float.go:808
				return x.MinPrec() == 1
//line /usr/local/go/src/math/big/float.go:808
				// _ = "end of CoverTab[4479]"
//line /usr/local/go/src/math/big/float.go:808
			}() {
//line /usr/local/go/src/math/big/float.go:808
				_go_fuzz_dep_.CoverTab[4480]++
									acc = Exact
//line /usr/local/go/src/math/big/float.go:809
				// _ = "end of CoverTab[4480]"
			} else {
//line /usr/local/go/src/math/big/float.go:810
				_go_fuzz_dep_.CoverTab[4481]++
//line /usr/local/go/src/math/big/float.go:810
				// _ = "end of CoverTab[4481]"
//line /usr/local/go/src/math/big/float.go:810
			}
//line /usr/local/go/src/math/big/float.go:810
			// _ = "end of CoverTab[4477]"
//line /usr/local/go/src/math/big/float.go:810
			_go_fuzz_dep_.CoverTab[4478]++
								return math.MinInt64, acc
//line /usr/local/go/src/math/big/float.go:811
			// _ = "end of CoverTab[4478]"
		} else {
//line /usr/local/go/src/math/big/float.go:812
			_go_fuzz_dep_.CoverTab[4482]++
//line /usr/local/go/src/math/big/float.go:812
			// _ = "end of CoverTab[4482]"
//line /usr/local/go/src/math/big/float.go:812
		}
//line /usr/local/go/src/math/big/float.go:812
		// _ = "end of CoverTab[4461]"
//line /usr/local/go/src/math/big/float.go:812
		_go_fuzz_dep_.CoverTab[4462]++

							return math.MaxInt64, Below
//line /usr/local/go/src/math/big/float.go:814
		// _ = "end of CoverTab[4462]"

	case zero:
//line /usr/local/go/src/math/big/float.go:816
		_go_fuzz_dep_.CoverTab[4463]++
							return 0, Exact
//line /usr/local/go/src/math/big/float.go:817
		// _ = "end of CoverTab[4463]"

	case inf:
//line /usr/local/go/src/math/big/float.go:819
		_go_fuzz_dep_.CoverTab[4464]++
							if x.neg {
//line /usr/local/go/src/math/big/float.go:820
			_go_fuzz_dep_.CoverTab[4483]++
								return math.MinInt64, Above
//line /usr/local/go/src/math/big/float.go:821
			// _ = "end of CoverTab[4483]"
		} else {
//line /usr/local/go/src/math/big/float.go:822
			_go_fuzz_dep_.CoverTab[4484]++
//line /usr/local/go/src/math/big/float.go:822
			// _ = "end of CoverTab[4484]"
//line /usr/local/go/src/math/big/float.go:822
		}
//line /usr/local/go/src/math/big/float.go:822
		// _ = "end of CoverTab[4464]"
//line /usr/local/go/src/math/big/float.go:822
		_go_fuzz_dep_.CoverTab[4465]++
							return math.MaxInt64, Below
//line /usr/local/go/src/math/big/float.go:823
		// _ = "end of CoverTab[4465]"
//line /usr/local/go/src/math/big/float.go:823
	default:
//line /usr/local/go/src/math/big/float.go:823
		_go_fuzz_dep_.CoverTab[4466]++
//line /usr/local/go/src/math/big/float.go:823
		// _ = "end of CoverTab[4466]"
	}
//line /usr/local/go/src/math/big/float.go:824
	// _ = "end of CoverTab[4455]"
//line /usr/local/go/src/math/big/float.go:824
	_go_fuzz_dep_.CoverTab[4456]++

						panic("unreachable")
//line /usr/local/go/src/math/big/float.go:826
	// _ = "end of CoverTab[4456]"
}

//line /usr/local/go/src/math/big/float.go:834
func (x *Float) Float32() (float32, Accuracy) {
//line /usr/local/go/src/math/big/float.go:834
	_go_fuzz_dep_.CoverTab[4485]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:835
		_go_fuzz_dep_.CoverTab[4488]++
							x.validate()
//line /usr/local/go/src/math/big/float.go:836
		// _ = "end of CoverTab[4488]"
	} else {
//line /usr/local/go/src/math/big/float.go:837
		_go_fuzz_dep_.CoverTab[4489]++
//line /usr/local/go/src/math/big/float.go:837
		// _ = "end of CoverTab[4489]"
//line /usr/local/go/src/math/big/float.go:837
	}
//line /usr/local/go/src/math/big/float.go:837
	// _ = "end of CoverTab[4485]"
//line /usr/local/go/src/math/big/float.go:837
	_go_fuzz_dep_.CoverTab[4486]++

						switch x.form {
	case finite:
//line /usr/local/go/src/math/big/float.go:840
		_go_fuzz_dep_.CoverTab[4490]++

//line /usr/local/go/src/math/big/float.go:843
		const (
			fbits	= 32
			mbits	= 23
			ebits	= fbits - mbits - 1
			bias	= 1<<(ebits-1) - 1
			dmin	= 1 - bias - mbits
			emin	= 1 - bias
			emax	= bias
		)

//line /usr/local/go/src/math/big/float.go:854
		e := x.exp - 1

//line /usr/local/go/src/math/big/float.go:860
		p := mbits + 1
		if e < emin {
//line /usr/local/go/src/math/big/float.go:861
			_go_fuzz_dep_.CoverTab[4500]++

								p = mbits + 1 - emin + int(e)

//line /usr/local/go/src/math/big/float.go:871
			if p < 0 || func() bool {
//line /usr/local/go/src/math/big/float.go:871
				_go_fuzz_dep_.CoverTab[4502]++
//line /usr/local/go/src/math/big/float.go:871
				return p == 0 && func() bool {
//line /usr/local/go/src/math/big/float.go:871
					_go_fuzz_dep_.CoverTab[4503]++
//line /usr/local/go/src/math/big/float.go:871
					return x.mant.sticky(uint(len(x.mant))*_W-1) == 0
//line /usr/local/go/src/math/big/float.go:871
					// _ = "end of CoverTab[4503]"
//line /usr/local/go/src/math/big/float.go:871
				}()
//line /usr/local/go/src/math/big/float.go:871
				// _ = "end of CoverTab[4502]"
//line /usr/local/go/src/math/big/float.go:871
			}() {
//line /usr/local/go/src/math/big/float.go:871
				_go_fuzz_dep_.CoverTab[4504]++

									if x.neg {
//line /usr/local/go/src/math/big/float.go:873
					_go_fuzz_dep_.CoverTab[4506]++
										var z float32
										return -z, Above
//line /usr/local/go/src/math/big/float.go:875
					// _ = "end of CoverTab[4506]"
				} else {
//line /usr/local/go/src/math/big/float.go:876
					_go_fuzz_dep_.CoverTab[4507]++
//line /usr/local/go/src/math/big/float.go:876
					// _ = "end of CoverTab[4507]"
//line /usr/local/go/src/math/big/float.go:876
				}
//line /usr/local/go/src/math/big/float.go:876
				// _ = "end of CoverTab[4504]"
//line /usr/local/go/src/math/big/float.go:876
				_go_fuzz_dep_.CoverTab[4505]++
									return 0.0, Below
//line /usr/local/go/src/math/big/float.go:877
				// _ = "end of CoverTab[4505]"
			} else {
//line /usr/local/go/src/math/big/float.go:878
				_go_fuzz_dep_.CoverTab[4508]++
//line /usr/local/go/src/math/big/float.go:878
				// _ = "end of CoverTab[4508]"
//line /usr/local/go/src/math/big/float.go:878
			}
//line /usr/local/go/src/math/big/float.go:878
			// _ = "end of CoverTab[4500]"
//line /usr/local/go/src/math/big/float.go:878
			_go_fuzz_dep_.CoverTab[4501]++

//line /usr/local/go/src/math/big/float.go:882
			if p == 0 {
//line /usr/local/go/src/math/big/float.go:882
				_go_fuzz_dep_.CoverTab[4509]++
									if x.neg {
//line /usr/local/go/src/math/big/float.go:883
					_go_fuzz_dep_.CoverTab[4511]++
										return -math.SmallestNonzeroFloat32, Below
//line /usr/local/go/src/math/big/float.go:884
					// _ = "end of CoverTab[4511]"
				} else {
//line /usr/local/go/src/math/big/float.go:885
					_go_fuzz_dep_.CoverTab[4512]++
//line /usr/local/go/src/math/big/float.go:885
					// _ = "end of CoverTab[4512]"
//line /usr/local/go/src/math/big/float.go:885
				}
//line /usr/local/go/src/math/big/float.go:885
				// _ = "end of CoverTab[4509]"
//line /usr/local/go/src/math/big/float.go:885
				_go_fuzz_dep_.CoverTab[4510]++
									return math.SmallestNonzeroFloat32, Above
//line /usr/local/go/src/math/big/float.go:886
				// _ = "end of CoverTab[4510]"
			} else {
//line /usr/local/go/src/math/big/float.go:887
				_go_fuzz_dep_.CoverTab[4513]++
//line /usr/local/go/src/math/big/float.go:887
				// _ = "end of CoverTab[4513]"
//line /usr/local/go/src/math/big/float.go:887
			}
//line /usr/local/go/src/math/big/float.go:887
			// _ = "end of CoverTab[4501]"
		} else {
//line /usr/local/go/src/math/big/float.go:888
			_go_fuzz_dep_.CoverTab[4514]++
//line /usr/local/go/src/math/big/float.go:888
			// _ = "end of CoverTab[4514]"
//line /usr/local/go/src/math/big/float.go:888
		}
//line /usr/local/go/src/math/big/float.go:888
		// _ = "end of CoverTab[4490]"
//line /usr/local/go/src/math/big/float.go:888
		_go_fuzz_dep_.CoverTab[4491]++

//line /usr/local/go/src/math/big/float.go:892
		var r Float
							r.prec = uint32(p)
							r.Set(x)
							e = r.exp - 1

//line /usr/local/go/src/math/big/float.go:900
		if r.form == inf || func() bool {
//line /usr/local/go/src/math/big/float.go:900
			_go_fuzz_dep_.CoverTab[4515]++
//line /usr/local/go/src/math/big/float.go:900
			return e > emax
//line /usr/local/go/src/math/big/float.go:900
			// _ = "end of CoverTab[4515]"
//line /usr/local/go/src/math/big/float.go:900
		}() {
//line /usr/local/go/src/math/big/float.go:900
			_go_fuzz_dep_.CoverTab[4516]++

								if x.neg {
//line /usr/local/go/src/math/big/float.go:902
				_go_fuzz_dep_.CoverTab[4518]++
									return float32(math.Inf(-1)), Below
//line /usr/local/go/src/math/big/float.go:903
				// _ = "end of CoverTab[4518]"
			} else {
//line /usr/local/go/src/math/big/float.go:904
				_go_fuzz_dep_.CoverTab[4519]++
//line /usr/local/go/src/math/big/float.go:904
				// _ = "end of CoverTab[4519]"
//line /usr/local/go/src/math/big/float.go:904
			}
//line /usr/local/go/src/math/big/float.go:904
			// _ = "end of CoverTab[4516]"
//line /usr/local/go/src/math/big/float.go:904
			_go_fuzz_dep_.CoverTab[4517]++
								return float32(math.Inf(+1)), Above
//line /usr/local/go/src/math/big/float.go:905
			// _ = "end of CoverTab[4517]"
		} else {
//line /usr/local/go/src/math/big/float.go:906
			_go_fuzz_dep_.CoverTab[4520]++
//line /usr/local/go/src/math/big/float.go:906
			// _ = "end of CoverTab[4520]"
//line /usr/local/go/src/math/big/float.go:906
		}
//line /usr/local/go/src/math/big/float.go:906
		// _ = "end of CoverTab[4491]"
//line /usr/local/go/src/math/big/float.go:906
		_go_fuzz_dep_.CoverTab[4492]++

//line /usr/local/go/src/math/big/float.go:910
		var sign, bexp, mant uint32
		if x.neg {
//line /usr/local/go/src/math/big/float.go:911
			_go_fuzz_dep_.CoverTab[4521]++
								sign = 1 << (fbits - 1)
//line /usr/local/go/src/math/big/float.go:912
			// _ = "end of CoverTab[4521]"
		} else {
//line /usr/local/go/src/math/big/float.go:913
			_go_fuzz_dep_.CoverTab[4522]++
//line /usr/local/go/src/math/big/float.go:913
			// _ = "end of CoverTab[4522]"
//line /usr/local/go/src/math/big/float.go:913
		}
//line /usr/local/go/src/math/big/float.go:913
		// _ = "end of CoverTab[4492]"
//line /usr/local/go/src/math/big/float.go:913
		_go_fuzz_dep_.CoverTab[4493]++

//line /usr/local/go/src/math/big/float.go:917
		if e < emin {
//line /usr/local/go/src/math/big/float.go:917
			_go_fuzz_dep_.CoverTab[4523]++

//line /usr/local/go/src/math/big/float.go:922
			p = mbits + 1 - emin + int(e)
								mant = msb32(r.mant) >> uint(fbits-p)
//line /usr/local/go/src/math/big/float.go:923
			// _ = "end of CoverTab[4523]"
		} else {
//line /usr/local/go/src/math/big/float.go:924
			_go_fuzz_dep_.CoverTab[4524]++

								bexp = uint32(e+bias) << mbits
								mant = msb32(r.mant) >> ebits & (1<<mbits - 1)
//line /usr/local/go/src/math/big/float.go:927
			// _ = "end of CoverTab[4524]"
		}
//line /usr/local/go/src/math/big/float.go:928
		// _ = "end of CoverTab[4493]"
//line /usr/local/go/src/math/big/float.go:928
		_go_fuzz_dep_.CoverTab[4494]++

							return math.Float32frombits(sign | bexp | mant), r.acc
//line /usr/local/go/src/math/big/float.go:930
		// _ = "end of CoverTab[4494]"

	case zero:
//line /usr/local/go/src/math/big/float.go:932
		_go_fuzz_dep_.CoverTab[4495]++
							if x.neg {
//line /usr/local/go/src/math/big/float.go:933
			_go_fuzz_dep_.CoverTab[4525]++
								var z float32
								return -z, Exact
//line /usr/local/go/src/math/big/float.go:935
			// _ = "end of CoverTab[4525]"
		} else {
//line /usr/local/go/src/math/big/float.go:936
			_go_fuzz_dep_.CoverTab[4526]++
//line /usr/local/go/src/math/big/float.go:936
			// _ = "end of CoverTab[4526]"
//line /usr/local/go/src/math/big/float.go:936
		}
//line /usr/local/go/src/math/big/float.go:936
		// _ = "end of CoverTab[4495]"
//line /usr/local/go/src/math/big/float.go:936
		_go_fuzz_dep_.CoverTab[4496]++
							return 0.0, Exact
//line /usr/local/go/src/math/big/float.go:937
		// _ = "end of CoverTab[4496]"

	case inf:
//line /usr/local/go/src/math/big/float.go:939
		_go_fuzz_dep_.CoverTab[4497]++
							if x.neg {
//line /usr/local/go/src/math/big/float.go:940
			_go_fuzz_dep_.CoverTab[4527]++
								return float32(math.Inf(-1)), Exact
//line /usr/local/go/src/math/big/float.go:941
			// _ = "end of CoverTab[4527]"
		} else {
//line /usr/local/go/src/math/big/float.go:942
			_go_fuzz_dep_.CoverTab[4528]++
//line /usr/local/go/src/math/big/float.go:942
			// _ = "end of CoverTab[4528]"
//line /usr/local/go/src/math/big/float.go:942
		}
//line /usr/local/go/src/math/big/float.go:942
		// _ = "end of CoverTab[4497]"
//line /usr/local/go/src/math/big/float.go:942
		_go_fuzz_dep_.CoverTab[4498]++
							return float32(math.Inf(+1)), Exact
//line /usr/local/go/src/math/big/float.go:943
		// _ = "end of CoverTab[4498]"
//line /usr/local/go/src/math/big/float.go:943
	default:
//line /usr/local/go/src/math/big/float.go:943
		_go_fuzz_dep_.CoverTab[4499]++
//line /usr/local/go/src/math/big/float.go:943
		// _ = "end of CoverTab[4499]"
	}
//line /usr/local/go/src/math/big/float.go:944
	// _ = "end of CoverTab[4486]"
//line /usr/local/go/src/math/big/float.go:944
	_go_fuzz_dep_.CoverTab[4487]++

						panic("unreachable")
//line /usr/local/go/src/math/big/float.go:946
	// _ = "end of CoverTab[4487]"
}

//line /usr/local/go/src/math/big/float.go:954
func (x *Float) Float64() (float64, Accuracy) {
//line /usr/local/go/src/math/big/float.go:954
	_go_fuzz_dep_.CoverTab[4529]++
						if debugFloat {
//line /usr/local/go/src/math/big/float.go:955
		_go_fuzz_dep_.CoverTab[4532]++
							x.validate()
//line /usr/local/go/src/math/big/float.go:956
		// _ = "end of CoverTab[4532]"
	} else {
//line /usr/local/go/src/math/big/float.go:957
		_go_fuzz_dep_.CoverTab[4533]++
//line /usr/local/go/src/math/big/float.go:957
		// _ = "end of CoverTab[4533]"
//line /usr/local/go/src/math/big/float.go:957
	}
//line /usr/local/go/src/math/big/float.go:957
	// _ = "end of CoverTab[4529]"
//line /usr/local/go/src/math/big/float.go:957
	_go_fuzz_dep_.CoverTab[4530]++

						switch x.form {
	case finite:
//line /usr/local/go/src/math/big/float.go:960
		_go_fuzz_dep_.CoverTab[4534]++

//line /usr/local/go/src/math/big/float.go:963
		const (
			fbits	= 64
			mbits	= 52
			ebits	= fbits - mbits - 1
			bias	= 1<<(ebits-1) - 1
			dmin	= 1 - bias - mbits
			emin	= 1 - bias
			emax	= bias
		)

//line /usr/local/go/src/math/big/float.go:974
		e := x.exp - 1

//line /usr/local/go/src/math/big/float.go:980
		p := mbits + 1
		if e < emin {
//line /usr/local/go/src/math/big/float.go:981
			_go_fuzz_dep_.CoverTab[4544]++

								p = mbits + 1 - emin + int(e)

//line /usr/local/go/src/math/big/float.go:991
			if p < 0 || func() bool {
//line /usr/local/go/src/math/big/float.go:991
				_go_fuzz_dep_.CoverTab[4546]++
//line /usr/local/go/src/math/big/float.go:991
				return p == 0 && func() bool {
//line /usr/local/go/src/math/big/float.go:991
					_go_fuzz_dep_.CoverTab[4547]++
//line /usr/local/go/src/math/big/float.go:991
					return x.mant.sticky(uint(len(x.mant))*_W-1) == 0
//line /usr/local/go/src/math/big/float.go:991
					// _ = "end of CoverTab[4547]"
//line /usr/local/go/src/math/big/float.go:991
				}()
//line /usr/local/go/src/math/big/float.go:991
				// _ = "end of CoverTab[4546]"
//line /usr/local/go/src/math/big/float.go:991
			}() {
//line /usr/local/go/src/math/big/float.go:991
				_go_fuzz_dep_.CoverTab[4548]++

									if x.neg {
//line /usr/local/go/src/math/big/float.go:993
					_go_fuzz_dep_.CoverTab[4550]++
										var z float64
										return -z, Above
//line /usr/local/go/src/math/big/float.go:995
					// _ = "end of CoverTab[4550]"
				} else {
//line /usr/local/go/src/math/big/float.go:996
					_go_fuzz_dep_.CoverTab[4551]++
//line /usr/local/go/src/math/big/float.go:996
					// _ = "end of CoverTab[4551]"
//line /usr/local/go/src/math/big/float.go:996
				}
//line /usr/local/go/src/math/big/float.go:996
				// _ = "end of CoverTab[4548]"
//line /usr/local/go/src/math/big/float.go:996
				_go_fuzz_dep_.CoverTab[4549]++
									return 0.0, Below
//line /usr/local/go/src/math/big/float.go:997
				// _ = "end of CoverTab[4549]"
			} else {
//line /usr/local/go/src/math/big/float.go:998
				_go_fuzz_dep_.CoverTab[4552]++
//line /usr/local/go/src/math/big/float.go:998
				// _ = "end of CoverTab[4552]"
//line /usr/local/go/src/math/big/float.go:998
			}
//line /usr/local/go/src/math/big/float.go:998
				// _ = "end of CoverTab[4544]"
//line /usr/local/go/src/math/big/float.go:998
				_go_fuzz_dep_.CoverTab[4545]++

//line /usr/local/go/src/math/big/float.go:1002
			if p == 0 {
//line /usr/local/go/src/math/big/float.go:1002
				_go_fuzz_dep_.CoverTab[4553]++
										if x.neg {
//line /usr/local/go/src/math/big/float.go:1003
					_go_fuzz_dep_.CoverTab[4555]++
											return -math.SmallestNonzeroFloat64, Below
//line /usr/local/go/src/math/big/float.go:1004
					// _ = "end of CoverTab[4555]"
				} else {
//line /usr/local/go/src/math/big/float.go:1005
					_go_fuzz_dep_.CoverTab[4556]++
//line /usr/local/go/src/math/big/float.go:1005
					// _ = "end of CoverTab[4556]"
//line /usr/local/go/src/math/big/float.go:1005
				}
//line /usr/local/go/src/math/big/float.go:1005
				// _ = "end of CoverTab[4553]"
//line /usr/local/go/src/math/big/float.go:1005
				_go_fuzz_dep_.CoverTab[4554]++
										return math.SmallestNonzeroFloat64, Above
//line /usr/local/go/src/math/big/float.go:1006
				// _ = "end of CoverTab[4554]"
			} else {
//line /usr/local/go/src/math/big/float.go:1007
				_go_fuzz_dep_.CoverTab[4557]++
//line /usr/local/go/src/math/big/float.go:1007
				// _ = "end of CoverTab[4557]"
//line /usr/local/go/src/math/big/float.go:1007
			}
//line /usr/local/go/src/math/big/float.go:1007
			// _ = "end of CoverTab[4545]"
		} else {
//line /usr/local/go/src/math/big/float.go:1008
			_go_fuzz_dep_.CoverTab[4558]++
//line /usr/local/go/src/math/big/float.go:1008
			// _ = "end of CoverTab[4558]"
//line /usr/local/go/src/math/big/float.go:1008
		}
//line /usr/local/go/src/math/big/float.go:1008
		// _ = "end of CoverTab[4534]"
//line /usr/local/go/src/math/big/float.go:1008
		_go_fuzz_dep_.CoverTab[4535]++

//line /usr/local/go/src/math/big/float.go:1012
		var r Float
								r.prec = uint32(p)
								r.Set(x)
								e = r.exp - 1

//line /usr/local/go/src/math/big/float.go:1020
		if r.form == inf || func() bool {
//line /usr/local/go/src/math/big/float.go:1020
			_go_fuzz_dep_.CoverTab[4559]++
//line /usr/local/go/src/math/big/float.go:1020
			return e > emax
//line /usr/local/go/src/math/big/float.go:1020
			// _ = "end of CoverTab[4559]"
//line /usr/local/go/src/math/big/float.go:1020
		}() {
//line /usr/local/go/src/math/big/float.go:1020
			_go_fuzz_dep_.CoverTab[4560]++

									if x.neg {
//line /usr/local/go/src/math/big/float.go:1022
				_go_fuzz_dep_.CoverTab[4562]++
										return math.Inf(-1), Below
//line /usr/local/go/src/math/big/float.go:1023
				// _ = "end of CoverTab[4562]"
			} else {
//line /usr/local/go/src/math/big/float.go:1024
				_go_fuzz_dep_.CoverTab[4563]++
//line /usr/local/go/src/math/big/float.go:1024
				// _ = "end of CoverTab[4563]"
//line /usr/local/go/src/math/big/float.go:1024
			}
//line /usr/local/go/src/math/big/float.go:1024
			// _ = "end of CoverTab[4560]"
//line /usr/local/go/src/math/big/float.go:1024
			_go_fuzz_dep_.CoverTab[4561]++
									return math.Inf(+1), Above
//line /usr/local/go/src/math/big/float.go:1025
			// _ = "end of CoverTab[4561]"
		} else {
//line /usr/local/go/src/math/big/float.go:1026
			_go_fuzz_dep_.CoverTab[4564]++
//line /usr/local/go/src/math/big/float.go:1026
			// _ = "end of CoverTab[4564]"
//line /usr/local/go/src/math/big/float.go:1026
		}
//line /usr/local/go/src/math/big/float.go:1026
		// _ = "end of CoverTab[4535]"
//line /usr/local/go/src/math/big/float.go:1026
		_go_fuzz_dep_.CoverTab[4536]++

//line /usr/local/go/src/math/big/float.go:1030
		var sign, bexp, mant uint64
		if x.neg {
//line /usr/local/go/src/math/big/float.go:1031
			_go_fuzz_dep_.CoverTab[4565]++
									sign = 1 << (fbits - 1)
//line /usr/local/go/src/math/big/float.go:1032
			// _ = "end of CoverTab[4565]"
		} else {
//line /usr/local/go/src/math/big/float.go:1033
			_go_fuzz_dep_.CoverTab[4566]++
//line /usr/local/go/src/math/big/float.go:1033
			// _ = "end of CoverTab[4566]"
//line /usr/local/go/src/math/big/float.go:1033
		}
//line /usr/local/go/src/math/big/float.go:1033
		// _ = "end of CoverTab[4536]"
//line /usr/local/go/src/math/big/float.go:1033
		_go_fuzz_dep_.CoverTab[4537]++

//line /usr/local/go/src/math/big/float.go:1037
		if e < emin {
//line /usr/local/go/src/math/big/float.go:1037
			_go_fuzz_dep_.CoverTab[4567]++

//line /usr/local/go/src/math/big/float.go:1042
			p = mbits + 1 - emin + int(e)
									mant = msb64(r.mant) >> uint(fbits-p)
//line /usr/local/go/src/math/big/float.go:1043
			// _ = "end of CoverTab[4567]"
		} else {
//line /usr/local/go/src/math/big/float.go:1044
			_go_fuzz_dep_.CoverTab[4568]++

									bexp = uint64(e+bias) << mbits
									mant = msb64(r.mant) >> ebits & (1<<mbits - 1)
//line /usr/local/go/src/math/big/float.go:1047
			// _ = "end of CoverTab[4568]"
		}
//line /usr/local/go/src/math/big/float.go:1048
		// _ = "end of CoverTab[4537]"
//line /usr/local/go/src/math/big/float.go:1048
		_go_fuzz_dep_.CoverTab[4538]++

								return math.Float64frombits(sign | bexp | mant), r.acc
//line /usr/local/go/src/math/big/float.go:1050
		// _ = "end of CoverTab[4538]"

	case zero:
//line /usr/local/go/src/math/big/float.go:1052
		_go_fuzz_dep_.CoverTab[4539]++
								if x.neg {
//line /usr/local/go/src/math/big/float.go:1053
			_go_fuzz_dep_.CoverTab[4569]++
									var z float64
									return -z, Exact
//line /usr/local/go/src/math/big/float.go:1055
			// _ = "end of CoverTab[4569]"
		} else {
//line /usr/local/go/src/math/big/float.go:1056
			_go_fuzz_dep_.CoverTab[4570]++
//line /usr/local/go/src/math/big/float.go:1056
			// _ = "end of CoverTab[4570]"
//line /usr/local/go/src/math/big/float.go:1056
		}
//line /usr/local/go/src/math/big/float.go:1056
		// _ = "end of CoverTab[4539]"
//line /usr/local/go/src/math/big/float.go:1056
		_go_fuzz_dep_.CoverTab[4540]++
								return 0.0, Exact
//line /usr/local/go/src/math/big/float.go:1057
		// _ = "end of CoverTab[4540]"

	case inf:
//line /usr/local/go/src/math/big/float.go:1059
		_go_fuzz_dep_.CoverTab[4541]++
								if x.neg {
//line /usr/local/go/src/math/big/float.go:1060
			_go_fuzz_dep_.CoverTab[4571]++
									return math.Inf(-1), Exact
//line /usr/local/go/src/math/big/float.go:1061
			// _ = "end of CoverTab[4571]"
		} else {
//line /usr/local/go/src/math/big/float.go:1062
			_go_fuzz_dep_.CoverTab[4572]++
//line /usr/local/go/src/math/big/float.go:1062
			// _ = "end of CoverTab[4572]"
//line /usr/local/go/src/math/big/float.go:1062
		}
//line /usr/local/go/src/math/big/float.go:1062
		// _ = "end of CoverTab[4541]"
//line /usr/local/go/src/math/big/float.go:1062
		_go_fuzz_dep_.CoverTab[4542]++
								return math.Inf(+1), Exact
//line /usr/local/go/src/math/big/float.go:1063
		// _ = "end of CoverTab[4542]"
//line /usr/local/go/src/math/big/float.go:1063
	default:
//line /usr/local/go/src/math/big/float.go:1063
		_go_fuzz_dep_.CoverTab[4543]++
//line /usr/local/go/src/math/big/float.go:1063
		// _ = "end of CoverTab[4543]"
	}
//line /usr/local/go/src/math/big/float.go:1064
	// _ = "end of CoverTab[4530]"
//line /usr/local/go/src/math/big/float.go:1064
	_go_fuzz_dep_.CoverTab[4531]++

							panic("unreachable")
//line /usr/local/go/src/math/big/float.go:1066
	// _ = "end of CoverTab[4531]"
}

//line /usr/local/go/src/math/big/float.go:1075
func (x *Float) Int(z *Int) (*Int, Accuracy) {
//line /usr/local/go/src/math/big/float.go:1075
	_go_fuzz_dep_.CoverTab[4573]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1076
		_go_fuzz_dep_.CoverTab[4577]++
								x.validate()
//line /usr/local/go/src/math/big/float.go:1077
		// _ = "end of CoverTab[4577]"
	} else {
//line /usr/local/go/src/math/big/float.go:1078
		_go_fuzz_dep_.CoverTab[4578]++
//line /usr/local/go/src/math/big/float.go:1078
		// _ = "end of CoverTab[4578]"
//line /usr/local/go/src/math/big/float.go:1078
	}
//line /usr/local/go/src/math/big/float.go:1078
	// _ = "end of CoverTab[4573]"
//line /usr/local/go/src/math/big/float.go:1078
	_go_fuzz_dep_.CoverTab[4574]++

							if z == nil && func() bool {
//line /usr/local/go/src/math/big/float.go:1080
		_go_fuzz_dep_.CoverTab[4579]++
//line /usr/local/go/src/math/big/float.go:1080
		return x.form <= finite
//line /usr/local/go/src/math/big/float.go:1080
		// _ = "end of CoverTab[4579]"
//line /usr/local/go/src/math/big/float.go:1080
	}() {
//line /usr/local/go/src/math/big/float.go:1080
		_go_fuzz_dep_.CoverTab[4580]++
								z = new(Int)
//line /usr/local/go/src/math/big/float.go:1081
		// _ = "end of CoverTab[4580]"
	} else {
//line /usr/local/go/src/math/big/float.go:1082
		_go_fuzz_dep_.CoverTab[4581]++
//line /usr/local/go/src/math/big/float.go:1082
		// _ = "end of CoverTab[4581]"
//line /usr/local/go/src/math/big/float.go:1082
	}
//line /usr/local/go/src/math/big/float.go:1082
	// _ = "end of CoverTab[4574]"
//line /usr/local/go/src/math/big/float.go:1082
	_go_fuzz_dep_.CoverTab[4575]++

							switch x.form {
	case finite:
//line /usr/local/go/src/math/big/float.go:1085
		_go_fuzz_dep_.CoverTab[4582]++

								acc := makeAcc(x.neg)
								if x.exp <= 0 {
//line /usr/local/go/src/math/big/float.go:1088
			_go_fuzz_dep_.CoverTab[4590]++

									return z.SetInt64(0), acc
//line /usr/local/go/src/math/big/float.go:1090
			// _ = "end of CoverTab[4590]"
		} else {
//line /usr/local/go/src/math/big/float.go:1091
			_go_fuzz_dep_.CoverTab[4591]++
//line /usr/local/go/src/math/big/float.go:1091
			// _ = "end of CoverTab[4591]"
//line /usr/local/go/src/math/big/float.go:1091
		}
//line /usr/local/go/src/math/big/float.go:1091
		// _ = "end of CoverTab[4582]"
//line /usr/local/go/src/math/big/float.go:1091
		_go_fuzz_dep_.CoverTab[4583]++

//line /usr/local/go/src/math/big/float.go:1096
		allBits := uint(len(x.mant)) * _W
		exp := uint(x.exp)
		if x.MinPrec() <= exp {
//line /usr/local/go/src/math/big/float.go:1098
			_go_fuzz_dep_.CoverTab[4592]++
									acc = Exact
//line /usr/local/go/src/math/big/float.go:1099
			// _ = "end of CoverTab[4592]"
		} else {
//line /usr/local/go/src/math/big/float.go:1100
			_go_fuzz_dep_.CoverTab[4593]++
//line /usr/local/go/src/math/big/float.go:1100
			// _ = "end of CoverTab[4593]"
//line /usr/local/go/src/math/big/float.go:1100
		}
//line /usr/local/go/src/math/big/float.go:1100
		// _ = "end of CoverTab[4583]"
//line /usr/local/go/src/math/big/float.go:1100
		_go_fuzz_dep_.CoverTab[4584]++

								if z == nil {
//line /usr/local/go/src/math/big/float.go:1102
			_go_fuzz_dep_.CoverTab[4594]++
									z = new(Int)
//line /usr/local/go/src/math/big/float.go:1103
			// _ = "end of CoverTab[4594]"
		} else {
//line /usr/local/go/src/math/big/float.go:1104
			_go_fuzz_dep_.CoverTab[4595]++
//line /usr/local/go/src/math/big/float.go:1104
			// _ = "end of CoverTab[4595]"
//line /usr/local/go/src/math/big/float.go:1104
		}
//line /usr/local/go/src/math/big/float.go:1104
		// _ = "end of CoverTab[4584]"
//line /usr/local/go/src/math/big/float.go:1104
		_go_fuzz_dep_.CoverTab[4585]++
								z.neg = x.neg
								switch {
		case exp > allBits:
//line /usr/local/go/src/math/big/float.go:1107
			_go_fuzz_dep_.CoverTab[4596]++
									z.abs = z.abs.shl(x.mant, exp-allBits)
//line /usr/local/go/src/math/big/float.go:1108
			// _ = "end of CoverTab[4596]"
		default:
//line /usr/local/go/src/math/big/float.go:1109
			_go_fuzz_dep_.CoverTab[4597]++
									z.abs = z.abs.set(x.mant)
//line /usr/local/go/src/math/big/float.go:1110
			// _ = "end of CoverTab[4597]"
		case exp < allBits:
//line /usr/local/go/src/math/big/float.go:1111
			_go_fuzz_dep_.CoverTab[4598]++
									z.abs = z.abs.shr(x.mant, allBits-exp)
//line /usr/local/go/src/math/big/float.go:1112
			// _ = "end of CoverTab[4598]"
		}
//line /usr/local/go/src/math/big/float.go:1113
		// _ = "end of CoverTab[4585]"
//line /usr/local/go/src/math/big/float.go:1113
		_go_fuzz_dep_.CoverTab[4586]++
								return z, acc
//line /usr/local/go/src/math/big/float.go:1114
		// _ = "end of CoverTab[4586]"

	case zero:
//line /usr/local/go/src/math/big/float.go:1116
		_go_fuzz_dep_.CoverTab[4587]++
								return z.SetInt64(0), Exact
//line /usr/local/go/src/math/big/float.go:1117
		// _ = "end of CoverTab[4587]"

	case inf:
//line /usr/local/go/src/math/big/float.go:1119
		_go_fuzz_dep_.CoverTab[4588]++
								return nil, makeAcc(x.neg)
//line /usr/local/go/src/math/big/float.go:1120
		// _ = "end of CoverTab[4588]"
//line /usr/local/go/src/math/big/float.go:1120
	default:
//line /usr/local/go/src/math/big/float.go:1120
		_go_fuzz_dep_.CoverTab[4589]++
//line /usr/local/go/src/math/big/float.go:1120
		// _ = "end of CoverTab[4589]"
	}
//line /usr/local/go/src/math/big/float.go:1121
	// _ = "end of CoverTab[4575]"
//line /usr/local/go/src/math/big/float.go:1121
	_go_fuzz_dep_.CoverTab[4576]++

							panic("unreachable")
//line /usr/local/go/src/math/big/float.go:1123
	// _ = "end of CoverTab[4576]"
}

//line /usr/local/go/src/math/big/float.go:1131
func (x *Float) Rat(z *Rat) (*Rat, Accuracy) {
//line /usr/local/go/src/math/big/float.go:1131
	_go_fuzz_dep_.CoverTab[4599]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1132
		_go_fuzz_dep_.CoverTab[4603]++
								x.validate()
//line /usr/local/go/src/math/big/float.go:1133
		// _ = "end of CoverTab[4603]"
	} else {
//line /usr/local/go/src/math/big/float.go:1134
		_go_fuzz_dep_.CoverTab[4604]++
//line /usr/local/go/src/math/big/float.go:1134
		// _ = "end of CoverTab[4604]"
//line /usr/local/go/src/math/big/float.go:1134
	}
//line /usr/local/go/src/math/big/float.go:1134
	// _ = "end of CoverTab[4599]"
//line /usr/local/go/src/math/big/float.go:1134
	_go_fuzz_dep_.CoverTab[4600]++

							if z == nil && func() bool {
//line /usr/local/go/src/math/big/float.go:1136
		_go_fuzz_dep_.CoverTab[4605]++
//line /usr/local/go/src/math/big/float.go:1136
		return x.form <= finite
//line /usr/local/go/src/math/big/float.go:1136
		// _ = "end of CoverTab[4605]"
//line /usr/local/go/src/math/big/float.go:1136
	}() {
//line /usr/local/go/src/math/big/float.go:1136
		_go_fuzz_dep_.CoverTab[4606]++
								z = new(Rat)
//line /usr/local/go/src/math/big/float.go:1137
		// _ = "end of CoverTab[4606]"
	} else {
//line /usr/local/go/src/math/big/float.go:1138
		_go_fuzz_dep_.CoverTab[4607]++
//line /usr/local/go/src/math/big/float.go:1138
		// _ = "end of CoverTab[4607]"
//line /usr/local/go/src/math/big/float.go:1138
	}
//line /usr/local/go/src/math/big/float.go:1138
	// _ = "end of CoverTab[4600]"
//line /usr/local/go/src/math/big/float.go:1138
	_go_fuzz_dep_.CoverTab[4601]++

							switch x.form {
	case finite:
//line /usr/local/go/src/math/big/float.go:1141
		_go_fuzz_dep_.CoverTab[4608]++

								allBits := int32(len(x.mant)) * _W

								z.a.neg = x.neg
								switch {
		case x.exp > allBits:
//line /usr/local/go/src/math/big/float.go:1147
			_go_fuzz_dep_.CoverTab[4613]++
									z.a.abs = z.a.abs.shl(x.mant, uint(x.exp-allBits))
									z.b.abs = z.b.abs[:0]
//line /usr/local/go/src/math/big/float.go:1149
			// _ = "end of CoverTab[4613]"

		default:
//line /usr/local/go/src/math/big/float.go:1151
			_go_fuzz_dep_.CoverTab[4614]++
									z.a.abs = z.a.abs.set(x.mant)
									z.b.abs = z.b.abs[:0]
//line /usr/local/go/src/math/big/float.go:1153
			// _ = "end of CoverTab[4614]"

		case x.exp < allBits:
//line /usr/local/go/src/math/big/float.go:1155
			_go_fuzz_dep_.CoverTab[4615]++
									z.a.abs = z.a.abs.set(x.mant)
									t := z.b.abs.setUint64(1)
									z.b.abs = t.shl(t, uint(allBits-x.exp))
									z.norm()
//line /usr/local/go/src/math/big/float.go:1159
			// _ = "end of CoverTab[4615]"
		}
//line /usr/local/go/src/math/big/float.go:1160
		// _ = "end of CoverTab[4608]"
//line /usr/local/go/src/math/big/float.go:1160
		_go_fuzz_dep_.CoverTab[4609]++
								return z, Exact
//line /usr/local/go/src/math/big/float.go:1161
		// _ = "end of CoverTab[4609]"

	case zero:
//line /usr/local/go/src/math/big/float.go:1163
		_go_fuzz_dep_.CoverTab[4610]++
								return z.SetInt64(0), Exact
//line /usr/local/go/src/math/big/float.go:1164
		// _ = "end of CoverTab[4610]"

	case inf:
//line /usr/local/go/src/math/big/float.go:1166
		_go_fuzz_dep_.CoverTab[4611]++
								return nil, makeAcc(x.neg)
//line /usr/local/go/src/math/big/float.go:1167
		// _ = "end of CoverTab[4611]"
//line /usr/local/go/src/math/big/float.go:1167
	default:
//line /usr/local/go/src/math/big/float.go:1167
		_go_fuzz_dep_.CoverTab[4612]++
//line /usr/local/go/src/math/big/float.go:1167
		// _ = "end of CoverTab[4612]"
	}
//line /usr/local/go/src/math/big/float.go:1168
	// _ = "end of CoverTab[4601]"
//line /usr/local/go/src/math/big/float.go:1168
	_go_fuzz_dep_.CoverTab[4602]++

							panic("unreachable")
//line /usr/local/go/src/math/big/float.go:1170
	// _ = "end of CoverTab[4602]"
}

//line /usr/local/go/src/math/big/float.go:1175
func (z *Float) Abs(x *Float) *Float {
//line /usr/local/go/src/math/big/float.go:1175
	_go_fuzz_dep_.CoverTab[4616]++
							z.Set(x)
							z.neg = false
							return z
//line /usr/local/go/src/math/big/float.go:1178
	// _ = "end of CoverTab[4616]"
}

//line /usr/local/go/src/math/big/float.go:1183
func (z *Float) Neg(x *Float) *Float {
//line /usr/local/go/src/math/big/float.go:1183
	_go_fuzz_dep_.CoverTab[4617]++
							z.Set(x)
							z.neg = !z.neg
							return z
//line /usr/local/go/src/math/big/float.go:1186
	// _ = "end of CoverTab[4617]"
}

func validateBinaryOperands(x, y *Float) {
//line /usr/local/go/src/math/big/float.go:1189
	_go_fuzz_dep_.CoverTab[4618]++
							if !debugFloat {
//line /usr/local/go/src/math/big/float.go:1190
		_go_fuzz_dep_.CoverTab[4621]++

								panic("validateBinaryOperands called but debugFloat is not set")
//line /usr/local/go/src/math/big/float.go:1192
		// _ = "end of CoverTab[4621]"
	} else {
//line /usr/local/go/src/math/big/float.go:1193
		_go_fuzz_dep_.CoverTab[4622]++
//line /usr/local/go/src/math/big/float.go:1193
		// _ = "end of CoverTab[4622]"
//line /usr/local/go/src/math/big/float.go:1193
	}
//line /usr/local/go/src/math/big/float.go:1193
	// _ = "end of CoverTab[4618]"
//line /usr/local/go/src/math/big/float.go:1193
	_go_fuzz_dep_.CoverTab[4619]++
							if len(x.mant) == 0 {
//line /usr/local/go/src/math/big/float.go:1194
		_go_fuzz_dep_.CoverTab[4623]++
								panic("empty mantissa for x")
//line /usr/local/go/src/math/big/float.go:1195
		// _ = "end of CoverTab[4623]"
	} else {
//line /usr/local/go/src/math/big/float.go:1196
		_go_fuzz_dep_.CoverTab[4624]++
//line /usr/local/go/src/math/big/float.go:1196
		// _ = "end of CoverTab[4624]"
//line /usr/local/go/src/math/big/float.go:1196
	}
//line /usr/local/go/src/math/big/float.go:1196
	// _ = "end of CoverTab[4619]"
//line /usr/local/go/src/math/big/float.go:1196
	_go_fuzz_dep_.CoverTab[4620]++
							if len(y.mant) == 0 {
//line /usr/local/go/src/math/big/float.go:1197
		_go_fuzz_dep_.CoverTab[4625]++
								panic("empty mantissa for y")
//line /usr/local/go/src/math/big/float.go:1198
		// _ = "end of CoverTab[4625]"
	} else {
//line /usr/local/go/src/math/big/float.go:1199
		_go_fuzz_dep_.CoverTab[4626]++
//line /usr/local/go/src/math/big/float.go:1199
		// _ = "end of CoverTab[4626]"
//line /usr/local/go/src/math/big/float.go:1199
	}
//line /usr/local/go/src/math/big/float.go:1199
	// _ = "end of CoverTab[4620]"
}

//line /usr/local/go/src/math/big/float.go:1205
func (z *Float) uadd(x, y *Float) {
//line /usr/local/go/src/math/big/float.go:1205
	_go_fuzz_dep_.CoverTab[4627]++

//line /usr/local/go/src/math/big/float.go:1216
	if debugFloat {
//line /usr/local/go/src/math/big/float.go:1216
		_go_fuzz_dep_.CoverTab[4630]++
								validateBinaryOperands(x, y)
//line /usr/local/go/src/math/big/float.go:1217
		// _ = "end of CoverTab[4630]"
	} else {
//line /usr/local/go/src/math/big/float.go:1218
		_go_fuzz_dep_.CoverTab[4631]++
//line /usr/local/go/src/math/big/float.go:1218
		// _ = "end of CoverTab[4631]"
//line /usr/local/go/src/math/big/float.go:1218
	}
//line /usr/local/go/src/math/big/float.go:1218
	// _ = "end of CoverTab[4627]"
//line /usr/local/go/src/math/big/float.go:1218
	_go_fuzz_dep_.CoverTab[4628]++

//line /usr/local/go/src/math/big/float.go:1222
	ex := int64(x.exp) - int64(len(x.mant))*_W
	ey := int64(y.exp) - int64(len(y.mant))*_W

	al := alias(z.mant, x.mant) || func() bool {
//line /usr/local/go/src/math/big/float.go:1225
		_go_fuzz_dep_.CoverTab[4632]++
//line /usr/local/go/src/math/big/float.go:1225
		return alias(z.mant, y.mant)
//line /usr/local/go/src/math/big/float.go:1225
		// _ = "end of CoverTab[4632]"
//line /usr/local/go/src/math/big/float.go:1225
	}()

//line /usr/local/go/src/math/big/float.go:1229
	switch {
	case ex < ey:
//line /usr/local/go/src/math/big/float.go:1230
		_go_fuzz_dep_.CoverTab[4633]++
								if al {
//line /usr/local/go/src/math/big/float.go:1231
			_go_fuzz_dep_.CoverTab[4637]++
									t := nat(nil).shl(y.mant, uint(ey-ex))
									z.mant = z.mant.add(x.mant, t)
//line /usr/local/go/src/math/big/float.go:1233
			// _ = "end of CoverTab[4637]"
		} else {
//line /usr/local/go/src/math/big/float.go:1234
			_go_fuzz_dep_.CoverTab[4638]++
									z.mant = z.mant.shl(y.mant, uint(ey-ex))
									z.mant = z.mant.add(x.mant, z.mant)
//line /usr/local/go/src/math/big/float.go:1236
			// _ = "end of CoverTab[4638]"
		}
//line /usr/local/go/src/math/big/float.go:1237
		// _ = "end of CoverTab[4633]"
	default:
//line /usr/local/go/src/math/big/float.go:1238
		_go_fuzz_dep_.CoverTab[4634]++

								z.mant = z.mant.add(x.mant, y.mant)
//line /usr/local/go/src/math/big/float.go:1240
		// _ = "end of CoverTab[4634]"
	case ex > ey:
//line /usr/local/go/src/math/big/float.go:1241
		_go_fuzz_dep_.CoverTab[4635]++
								if al {
//line /usr/local/go/src/math/big/float.go:1242
			_go_fuzz_dep_.CoverTab[4639]++
									t := nat(nil).shl(x.mant, uint(ex-ey))
									z.mant = z.mant.add(t, y.mant)
//line /usr/local/go/src/math/big/float.go:1244
			// _ = "end of CoverTab[4639]"
		} else {
//line /usr/local/go/src/math/big/float.go:1245
			_go_fuzz_dep_.CoverTab[4640]++
									z.mant = z.mant.shl(x.mant, uint(ex-ey))
									z.mant = z.mant.add(z.mant, y.mant)
//line /usr/local/go/src/math/big/float.go:1247
			// _ = "end of CoverTab[4640]"
		}
//line /usr/local/go/src/math/big/float.go:1248
		// _ = "end of CoverTab[4635]"
//line /usr/local/go/src/math/big/float.go:1248
		_go_fuzz_dep_.CoverTab[4636]++
								ex = ey
//line /usr/local/go/src/math/big/float.go:1249
		// _ = "end of CoverTab[4636]"
	}
//line /usr/local/go/src/math/big/float.go:1250
	// _ = "end of CoverTab[4628]"
//line /usr/local/go/src/math/big/float.go:1250
	_go_fuzz_dep_.CoverTab[4629]++

//line /usr/local/go/src/math/big/float.go:1253
	z.setExpAndRound(ex+int64(len(z.mant))*_W-fnorm(z.mant), 0)
//line /usr/local/go/src/math/big/float.go:1253
	// _ = "end of CoverTab[4629]"
}

//line /usr/local/go/src/math/big/float.go:1259
func (z *Float) usub(x, y *Float) {
//line /usr/local/go/src/math/big/float.go:1259
	_go_fuzz_dep_.CoverTab[4641]++

//line /usr/local/go/src/math/big/float.go:1265
	if debugFloat {
//line /usr/local/go/src/math/big/float.go:1265
		_go_fuzz_dep_.CoverTab[4645]++
								validateBinaryOperands(x, y)
//line /usr/local/go/src/math/big/float.go:1266
		// _ = "end of CoverTab[4645]"
	} else {
//line /usr/local/go/src/math/big/float.go:1267
		_go_fuzz_dep_.CoverTab[4646]++
//line /usr/local/go/src/math/big/float.go:1267
		// _ = "end of CoverTab[4646]"
//line /usr/local/go/src/math/big/float.go:1267
	}
//line /usr/local/go/src/math/big/float.go:1267
	// _ = "end of CoverTab[4641]"
//line /usr/local/go/src/math/big/float.go:1267
	_go_fuzz_dep_.CoverTab[4642]++

							ex := int64(x.exp) - int64(len(x.mant))*_W
							ey := int64(y.exp) - int64(len(y.mant))*_W

							al := alias(z.mant, x.mant) || func() bool {
//line /usr/local/go/src/math/big/float.go:1272
		_go_fuzz_dep_.CoverTab[4647]++
//line /usr/local/go/src/math/big/float.go:1272
		return alias(z.mant, y.mant)
//line /usr/local/go/src/math/big/float.go:1272
		// _ = "end of CoverTab[4647]"
//line /usr/local/go/src/math/big/float.go:1272
	}()

							switch {
	case ex < ey:
//line /usr/local/go/src/math/big/float.go:1275
		_go_fuzz_dep_.CoverTab[4648]++
								if al {
//line /usr/local/go/src/math/big/float.go:1276
			_go_fuzz_dep_.CoverTab[4652]++
									t := nat(nil).shl(y.mant, uint(ey-ex))
									z.mant = t.sub(x.mant, t)
//line /usr/local/go/src/math/big/float.go:1278
			// _ = "end of CoverTab[4652]"
		} else {
//line /usr/local/go/src/math/big/float.go:1279
			_go_fuzz_dep_.CoverTab[4653]++
									z.mant = z.mant.shl(y.mant, uint(ey-ex))
									z.mant = z.mant.sub(x.mant, z.mant)
//line /usr/local/go/src/math/big/float.go:1281
			// _ = "end of CoverTab[4653]"
		}
//line /usr/local/go/src/math/big/float.go:1282
		// _ = "end of CoverTab[4648]"
	default:
//line /usr/local/go/src/math/big/float.go:1283
		_go_fuzz_dep_.CoverTab[4649]++

								z.mant = z.mant.sub(x.mant, y.mant)
//line /usr/local/go/src/math/big/float.go:1285
		// _ = "end of CoverTab[4649]"
	case ex > ey:
//line /usr/local/go/src/math/big/float.go:1286
		_go_fuzz_dep_.CoverTab[4650]++
								if al {
//line /usr/local/go/src/math/big/float.go:1287
			_go_fuzz_dep_.CoverTab[4654]++
									t := nat(nil).shl(x.mant, uint(ex-ey))
									z.mant = t.sub(t, y.mant)
//line /usr/local/go/src/math/big/float.go:1289
			// _ = "end of CoverTab[4654]"
		} else {
//line /usr/local/go/src/math/big/float.go:1290
			_go_fuzz_dep_.CoverTab[4655]++
									z.mant = z.mant.shl(x.mant, uint(ex-ey))
									z.mant = z.mant.sub(z.mant, y.mant)
//line /usr/local/go/src/math/big/float.go:1292
			// _ = "end of CoverTab[4655]"
		}
//line /usr/local/go/src/math/big/float.go:1293
		// _ = "end of CoverTab[4650]"
//line /usr/local/go/src/math/big/float.go:1293
		_go_fuzz_dep_.CoverTab[4651]++
								ex = ey
//line /usr/local/go/src/math/big/float.go:1294
		// _ = "end of CoverTab[4651]"
	}
//line /usr/local/go/src/math/big/float.go:1295
	// _ = "end of CoverTab[4642]"
//line /usr/local/go/src/math/big/float.go:1295
	_go_fuzz_dep_.CoverTab[4643]++

//line /usr/local/go/src/math/big/float.go:1298
	if len(z.mant) == 0 {
//line /usr/local/go/src/math/big/float.go:1298
		_go_fuzz_dep_.CoverTab[4656]++
								z.acc = Exact
								z.form = zero
								z.neg = false
								return
//line /usr/local/go/src/math/big/float.go:1302
		// _ = "end of CoverTab[4656]"
	} else {
//line /usr/local/go/src/math/big/float.go:1303
		_go_fuzz_dep_.CoverTab[4657]++
//line /usr/local/go/src/math/big/float.go:1303
		// _ = "end of CoverTab[4657]"
//line /usr/local/go/src/math/big/float.go:1303
	}
//line /usr/local/go/src/math/big/float.go:1303
	// _ = "end of CoverTab[4643]"
//line /usr/local/go/src/math/big/float.go:1303
	_go_fuzz_dep_.CoverTab[4644]++

//line /usr/local/go/src/math/big/float.go:1306
	z.setExpAndRound(ex+int64(len(z.mant))*_W-fnorm(z.mant), 0)
//line /usr/local/go/src/math/big/float.go:1306
	// _ = "end of CoverTab[4644]"
}

//line /usr/local/go/src/math/big/float.go:1312
func (z *Float) umul(x, y *Float) {
//line /usr/local/go/src/math/big/float.go:1312
	_go_fuzz_dep_.CoverTab[4658]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1313
		_go_fuzz_dep_.CoverTab[4661]++
								validateBinaryOperands(x, y)
//line /usr/local/go/src/math/big/float.go:1314
		// _ = "end of CoverTab[4661]"
	} else {
//line /usr/local/go/src/math/big/float.go:1315
		_go_fuzz_dep_.CoverTab[4662]++
//line /usr/local/go/src/math/big/float.go:1315
		// _ = "end of CoverTab[4662]"
//line /usr/local/go/src/math/big/float.go:1315
	}
//line /usr/local/go/src/math/big/float.go:1315
	// _ = "end of CoverTab[4658]"
//line /usr/local/go/src/math/big/float.go:1315
	_go_fuzz_dep_.CoverTab[4659]++

//line /usr/local/go/src/math/big/float.go:1323
	e := int64(x.exp) + int64(y.exp)
	if x == y {
//line /usr/local/go/src/math/big/float.go:1324
		_go_fuzz_dep_.CoverTab[4663]++
								z.mant = z.mant.sqr(x.mant)
//line /usr/local/go/src/math/big/float.go:1325
		// _ = "end of CoverTab[4663]"
	} else {
//line /usr/local/go/src/math/big/float.go:1326
		_go_fuzz_dep_.CoverTab[4664]++
								z.mant = z.mant.mul(x.mant, y.mant)
//line /usr/local/go/src/math/big/float.go:1327
		// _ = "end of CoverTab[4664]"
	}
//line /usr/local/go/src/math/big/float.go:1328
	// _ = "end of CoverTab[4659]"
//line /usr/local/go/src/math/big/float.go:1328
	_go_fuzz_dep_.CoverTab[4660]++
							z.setExpAndRound(e-fnorm(z.mant), 0)
//line /usr/local/go/src/math/big/float.go:1329
	// _ = "end of CoverTab[4660]"
}

//line /usr/local/go/src/math/big/float.go:1335
func (z *Float) uquo(x, y *Float) {
//line /usr/local/go/src/math/big/float.go:1335
	_go_fuzz_dep_.CoverTab[4665]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1336
		_go_fuzz_dep_.CoverTab[4669]++
								validateBinaryOperands(x, y)
//line /usr/local/go/src/math/big/float.go:1337
		// _ = "end of CoverTab[4669]"
	} else {
//line /usr/local/go/src/math/big/float.go:1338
		_go_fuzz_dep_.CoverTab[4670]++
//line /usr/local/go/src/math/big/float.go:1338
		// _ = "end of CoverTab[4670]"
//line /usr/local/go/src/math/big/float.go:1338
	}
//line /usr/local/go/src/math/big/float.go:1338
	// _ = "end of CoverTab[4665]"
//line /usr/local/go/src/math/big/float.go:1338
	_go_fuzz_dep_.CoverTab[4666]++

//line /usr/local/go/src/math/big/float.go:1343
	n := int(z.prec/_W) + 1

//line /usr/local/go/src/math/big/float.go:1346
	xadj := x.mant
	if d := n - len(x.mant) + len(y.mant); d > 0 {
//line /usr/local/go/src/math/big/float.go:1347
		_go_fuzz_dep_.CoverTab[4671]++

								xadj = make(nat, len(x.mant)+d)
								copy(xadj[d:], x.mant)
//line /usr/local/go/src/math/big/float.go:1350
		// _ = "end of CoverTab[4671]"
	} else {
//line /usr/local/go/src/math/big/float.go:1351
		_go_fuzz_dep_.CoverTab[4672]++
//line /usr/local/go/src/math/big/float.go:1351
		// _ = "end of CoverTab[4672]"
//line /usr/local/go/src/math/big/float.go:1351
	}
//line /usr/local/go/src/math/big/float.go:1351
	// _ = "end of CoverTab[4666]"
//line /usr/local/go/src/math/big/float.go:1351
	_go_fuzz_dep_.CoverTab[4667]++

//line /usr/local/go/src/math/big/float.go:1358
	d := len(xadj) - len(y.mant)

//line /usr/local/go/src/math/big/float.go:1361
	var r nat
							z.mant, r = z.mant.div(nil, xadj, y.mant)
							e := int64(x.exp) - int64(y.exp) - int64(d-len(z.mant))*_W

//line /usr/local/go/src/math/big/float.go:1369
	var sbit uint
	if len(r) > 0 {
//line /usr/local/go/src/math/big/float.go:1370
		_go_fuzz_dep_.CoverTab[4673]++
								sbit = 1
//line /usr/local/go/src/math/big/float.go:1371
		// _ = "end of CoverTab[4673]"
	} else {
//line /usr/local/go/src/math/big/float.go:1372
		_go_fuzz_dep_.CoverTab[4674]++
//line /usr/local/go/src/math/big/float.go:1372
		// _ = "end of CoverTab[4674]"
//line /usr/local/go/src/math/big/float.go:1372
	}
//line /usr/local/go/src/math/big/float.go:1372
	// _ = "end of CoverTab[4667]"
//line /usr/local/go/src/math/big/float.go:1372
	_go_fuzz_dep_.CoverTab[4668]++

							z.setExpAndRound(e-fnorm(z.mant), sbit)
//line /usr/local/go/src/math/big/float.go:1374
	// _ = "end of CoverTab[4668]"
}

//line /usr/local/go/src/math/big/float.go:1380
func (x *Float) ucmp(y *Float) int {
//line /usr/local/go/src/math/big/float.go:1380
	_go_fuzz_dep_.CoverTab[4675]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1381
		_go_fuzz_dep_.CoverTab[4679]++
								validateBinaryOperands(x, y)
//line /usr/local/go/src/math/big/float.go:1382
		// _ = "end of CoverTab[4679]"
	} else {
//line /usr/local/go/src/math/big/float.go:1383
		_go_fuzz_dep_.CoverTab[4680]++
//line /usr/local/go/src/math/big/float.go:1383
		// _ = "end of CoverTab[4680]"
//line /usr/local/go/src/math/big/float.go:1383
	}
//line /usr/local/go/src/math/big/float.go:1383
	// _ = "end of CoverTab[4675]"
//line /usr/local/go/src/math/big/float.go:1383
	_go_fuzz_dep_.CoverTab[4676]++

							switch {
	case x.exp < y.exp:
//line /usr/local/go/src/math/big/float.go:1386
		_go_fuzz_dep_.CoverTab[4681]++
								return -1
//line /usr/local/go/src/math/big/float.go:1387
		// _ = "end of CoverTab[4681]"
	case x.exp > y.exp:
//line /usr/local/go/src/math/big/float.go:1388
		_go_fuzz_dep_.CoverTab[4682]++
								return +1
//line /usr/local/go/src/math/big/float.go:1389
		// _ = "end of CoverTab[4682]"
//line /usr/local/go/src/math/big/float.go:1389
	default:
//line /usr/local/go/src/math/big/float.go:1389
		_go_fuzz_dep_.CoverTab[4683]++
//line /usr/local/go/src/math/big/float.go:1389
		// _ = "end of CoverTab[4683]"
	}
//line /usr/local/go/src/math/big/float.go:1390
	// _ = "end of CoverTab[4676]"
//line /usr/local/go/src/math/big/float.go:1390
	_go_fuzz_dep_.CoverTab[4677]++

//line /usr/local/go/src/math/big/float.go:1394
	i := len(x.mant)
	j := len(y.mant)
	for i > 0 || func() bool {
//line /usr/local/go/src/math/big/float.go:1396
		_go_fuzz_dep_.CoverTab[4684]++
//line /usr/local/go/src/math/big/float.go:1396
		return j > 0
//line /usr/local/go/src/math/big/float.go:1396
		// _ = "end of CoverTab[4684]"
//line /usr/local/go/src/math/big/float.go:1396
	}() {
//line /usr/local/go/src/math/big/float.go:1396
		_go_fuzz_dep_.CoverTab[4685]++
								var xm, ym Word
								if i > 0 {
//line /usr/local/go/src/math/big/float.go:1398
			_go_fuzz_dep_.CoverTab[4688]++
									i--
									xm = x.mant[i]
//line /usr/local/go/src/math/big/float.go:1400
			// _ = "end of CoverTab[4688]"
		} else {
//line /usr/local/go/src/math/big/float.go:1401
			_go_fuzz_dep_.CoverTab[4689]++
//line /usr/local/go/src/math/big/float.go:1401
			// _ = "end of CoverTab[4689]"
//line /usr/local/go/src/math/big/float.go:1401
		}
//line /usr/local/go/src/math/big/float.go:1401
		// _ = "end of CoverTab[4685]"
//line /usr/local/go/src/math/big/float.go:1401
		_go_fuzz_dep_.CoverTab[4686]++
								if j > 0 {
//line /usr/local/go/src/math/big/float.go:1402
			_go_fuzz_dep_.CoverTab[4690]++
									j--
									ym = y.mant[j]
//line /usr/local/go/src/math/big/float.go:1404
			// _ = "end of CoverTab[4690]"
		} else {
//line /usr/local/go/src/math/big/float.go:1405
			_go_fuzz_dep_.CoverTab[4691]++
//line /usr/local/go/src/math/big/float.go:1405
			// _ = "end of CoverTab[4691]"
//line /usr/local/go/src/math/big/float.go:1405
		}
//line /usr/local/go/src/math/big/float.go:1405
		// _ = "end of CoverTab[4686]"
//line /usr/local/go/src/math/big/float.go:1405
		_go_fuzz_dep_.CoverTab[4687]++
								switch {
		case xm < ym:
//line /usr/local/go/src/math/big/float.go:1407
			_go_fuzz_dep_.CoverTab[4692]++
									return -1
//line /usr/local/go/src/math/big/float.go:1408
			// _ = "end of CoverTab[4692]"
		case xm > ym:
//line /usr/local/go/src/math/big/float.go:1409
			_go_fuzz_dep_.CoverTab[4693]++
									return +1
//line /usr/local/go/src/math/big/float.go:1410
			// _ = "end of CoverTab[4693]"
//line /usr/local/go/src/math/big/float.go:1410
		default:
//line /usr/local/go/src/math/big/float.go:1410
			_go_fuzz_dep_.CoverTab[4694]++
//line /usr/local/go/src/math/big/float.go:1410
			// _ = "end of CoverTab[4694]"
		}
//line /usr/local/go/src/math/big/float.go:1411
		// _ = "end of CoverTab[4687]"
	}
//line /usr/local/go/src/math/big/float.go:1412
	// _ = "end of CoverTab[4677]"
//line /usr/local/go/src/math/big/float.go:1412
	_go_fuzz_dep_.CoverTab[4678]++

							return 0
//line /usr/local/go/src/math/big/float.go:1414
	// _ = "end of CoverTab[4678]"
}

//line /usr/local/go/src/math/big/float.go:1442
func (z *Float) Add(x, y *Float) *Float {
//line /usr/local/go/src/math/big/float.go:1442
	_go_fuzz_dep_.CoverTab[4695]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1443
		_go_fuzz_dep_.CoverTab[4702]++
								x.validate()
								y.validate()
//line /usr/local/go/src/math/big/float.go:1445
		// _ = "end of CoverTab[4702]"
	} else {
//line /usr/local/go/src/math/big/float.go:1446
		_go_fuzz_dep_.CoverTab[4703]++
//line /usr/local/go/src/math/big/float.go:1446
		// _ = "end of CoverTab[4703]"
//line /usr/local/go/src/math/big/float.go:1446
	}
//line /usr/local/go/src/math/big/float.go:1446
	// _ = "end of CoverTab[4695]"
//line /usr/local/go/src/math/big/float.go:1446
	_go_fuzz_dep_.CoverTab[4696]++

							if z.prec == 0 {
//line /usr/local/go/src/math/big/float.go:1448
		_go_fuzz_dep_.CoverTab[4704]++
								z.prec = umax32(x.prec, y.prec)
//line /usr/local/go/src/math/big/float.go:1449
		// _ = "end of CoverTab[4704]"
	} else {
//line /usr/local/go/src/math/big/float.go:1450
		_go_fuzz_dep_.CoverTab[4705]++
//line /usr/local/go/src/math/big/float.go:1450
		// _ = "end of CoverTab[4705]"
//line /usr/local/go/src/math/big/float.go:1450
	}
//line /usr/local/go/src/math/big/float.go:1450
	// _ = "end of CoverTab[4696]"
//line /usr/local/go/src/math/big/float.go:1450
	_go_fuzz_dep_.CoverTab[4697]++

							if x.form == finite && func() bool {
//line /usr/local/go/src/math/big/float.go:1452
		_go_fuzz_dep_.CoverTab[4706]++
//line /usr/local/go/src/math/big/float.go:1452
		return y.form == finite
//line /usr/local/go/src/math/big/float.go:1452
		// _ = "end of CoverTab[4706]"
//line /usr/local/go/src/math/big/float.go:1452
	}() {
//line /usr/local/go/src/math/big/float.go:1452
		_go_fuzz_dep_.CoverTab[4707]++

//line /usr/local/go/src/math/big/float.go:1460
		yneg := y.neg

		z.neg = x.neg
		if x.neg == yneg {
//line /usr/local/go/src/math/big/float.go:1463
			_go_fuzz_dep_.CoverTab[4710]++

//line /usr/local/go/src/math/big/float.go:1466
			z.uadd(x, y)
//line /usr/local/go/src/math/big/float.go:1466
			// _ = "end of CoverTab[4710]"
		} else {
//line /usr/local/go/src/math/big/float.go:1467
			_go_fuzz_dep_.CoverTab[4711]++

//line /usr/local/go/src/math/big/float.go:1470
			if x.ucmp(y) > 0 {
//line /usr/local/go/src/math/big/float.go:1470
				_go_fuzz_dep_.CoverTab[4712]++
										z.usub(x, y)
//line /usr/local/go/src/math/big/float.go:1471
				// _ = "end of CoverTab[4712]"
			} else {
//line /usr/local/go/src/math/big/float.go:1472
				_go_fuzz_dep_.CoverTab[4713]++
										z.neg = !z.neg
										z.usub(y, x)
//line /usr/local/go/src/math/big/float.go:1474
				// _ = "end of CoverTab[4713]"
			}
//line /usr/local/go/src/math/big/float.go:1475
			// _ = "end of CoverTab[4711]"
		}
//line /usr/local/go/src/math/big/float.go:1476
		// _ = "end of CoverTab[4707]"
//line /usr/local/go/src/math/big/float.go:1476
		_go_fuzz_dep_.CoverTab[4708]++
								if z.form == zero && func() bool {
//line /usr/local/go/src/math/big/float.go:1477
			_go_fuzz_dep_.CoverTab[4714]++
//line /usr/local/go/src/math/big/float.go:1477
			return z.mode == ToNegativeInf
//line /usr/local/go/src/math/big/float.go:1477
			// _ = "end of CoverTab[4714]"
//line /usr/local/go/src/math/big/float.go:1477
		}() && func() bool {
//line /usr/local/go/src/math/big/float.go:1477
			_go_fuzz_dep_.CoverTab[4715]++
//line /usr/local/go/src/math/big/float.go:1477
			return z.acc == Exact
//line /usr/local/go/src/math/big/float.go:1477
			// _ = "end of CoverTab[4715]"
//line /usr/local/go/src/math/big/float.go:1477
		}() {
//line /usr/local/go/src/math/big/float.go:1477
			_go_fuzz_dep_.CoverTab[4716]++
									z.neg = true
//line /usr/local/go/src/math/big/float.go:1478
			// _ = "end of CoverTab[4716]"
		} else {
//line /usr/local/go/src/math/big/float.go:1479
			_go_fuzz_dep_.CoverTab[4717]++
//line /usr/local/go/src/math/big/float.go:1479
			// _ = "end of CoverTab[4717]"
//line /usr/local/go/src/math/big/float.go:1479
		}
//line /usr/local/go/src/math/big/float.go:1479
		// _ = "end of CoverTab[4708]"
//line /usr/local/go/src/math/big/float.go:1479
		_go_fuzz_dep_.CoverTab[4709]++
								return z
//line /usr/local/go/src/math/big/float.go:1480
		// _ = "end of CoverTab[4709]"
	} else {
//line /usr/local/go/src/math/big/float.go:1481
		_go_fuzz_dep_.CoverTab[4718]++
//line /usr/local/go/src/math/big/float.go:1481
		// _ = "end of CoverTab[4718]"
//line /usr/local/go/src/math/big/float.go:1481
	}
//line /usr/local/go/src/math/big/float.go:1481
	// _ = "end of CoverTab[4697]"
//line /usr/local/go/src/math/big/float.go:1481
	_go_fuzz_dep_.CoverTab[4698]++

							if x.form == inf && func() bool {
//line /usr/local/go/src/math/big/float.go:1483
		_go_fuzz_dep_.CoverTab[4719]++
//line /usr/local/go/src/math/big/float.go:1483
		return y.form == inf
//line /usr/local/go/src/math/big/float.go:1483
		// _ = "end of CoverTab[4719]"
//line /usr/local/go/src/math/big/float.go:1483
	}() && func() bool {
//line /usr/local/go/src/math/big/float.go:1483
		_go_fuzz_dep_.CoverTab[4720]++
//line /usr/local/go/src/math/big/float.go:1483
		return x.neg != y.neg
//line /usr/local/go/src/math/big/float.go:1483
		// _ = "end of CoverTab[4720]"
//line /usr/local/go/src/math/big/float.go:1483
	}() {
//line /usr/local/go/src/math/big/float.go:1483
		_go_fuzz_dep_.CoverTab[4721]++

//line /usr/local/go/src/math/big/float.go:1487
		z.acc = Exact
								z.form = zero
								z.neg = false
								panic(ErrNaN{"addition of infinities with opposite signs"})
//line /usr/local/go/src/math/big/float.go:1490
		// _ = "end of CoverTab[4721]"
	} else {
//line /usr/local/go/src/math/big/float.go:1491
		_go_fuzz_dep_.CoverTab[4722]++
//line /usr/local/go/src/math/big/float.go:1491
		// _ = "end of CoverTab[4722]"
//line /usr/local/go/src/math/big/float.go:1491
	}
//line /usr/local/go/src/math/big/float.go:1491
	// _ = "end of CoverTab[4698]"
//line /usr/local/go/src/math/big/float.go:1491
	_go_fuzz_dep_.CoverTab[4699]++

							if x.form == zero && func() bool {
//line /usr/local/go/src/math/big/float.go:1493
		_go_fuzz_dep_.CoverTab[4723]++
//line /usr/local/go/src/math/big/float.go:1493
		return y.form == zero
//line /usr/local/go/src/math/big/float.go:1493
		// _ = "end of CoverTab[4723]"
//line /usr/local/go/src/math/big/float.go:1493
	}() {
//line /usr/local/go/src/math/big/float.go:1493
		_go_fuzz_dep_.CoverTab[4724]++

								z.acc = Exact
								z.form = zero
								z.neg = x.neg && func() bool {
//line /usr/local/go/src/math/big/float.go:1497
			_go_fuzz_dep_.CoverTab[4725]++
//line /usr/local/go/src/math/big/float.go:1497
			return y.neg
//line /usr/local/go/src/math/big/float.go:1497
			// _ = "end of CoverTab[4725]"
//line /usr/local/go/src/math/big/float.go:1497
		}()
								return z
//line /usr/local/go/src/math/big/float.go:1498
		// _ = "end of CoverTab[4724]"
	} else {
//line /usr/local/go/src/math/big/float.go:1499
		_go_fuzz_dep_.CoverTab[4726]++
//line /usr/local/go/src/math/big/float.go:1499
		// _ = "end of CoverTab[4726]"
//line /usr/local/go/src/math/big/float.go:1499
	}
//line /usr/local/go/src/math/big/float.go:1499
	// _ = "end of CoverTab[4699]"
//line /usr/local/go/src/math/big/float.go:1499
	_go_fuzz_dep_.CoverTab[4700]++

							if x.form == inf || func() bool {
//line /usr/local/go/src/math/big/float.go:1501
		_go_fuzz_dep_.CoverTab[4727]++
//line /usr/local/go/src/math/big/float.go:1501
		return y.form == zero
//line /usr/local/go/src/math/big/float.go:1501
		// _ = "end of CoverTab[4727]"
//line /usr/local/go/src/math/big/float.go:1501
	}() {
//line /usr/local/go/src/math/big/float.go:1501
		_go_fuzz_dep_.CoverTab[4728]++

//line /usr/local/go/src/math/big/float.go:1504
		return z.Set(x)
//line /usr/local/go/src/math/big/float.go:1504
		// _ = "end of CoverTab[4728]"
	} else {
//line /usr/local/go/src/math/big/float.go:1505
		_go_fuzz_dep_.CoverTab[4729]++
//line /usr/local/go/src/math/big/float.go:1505
		// _ = "end of CoverTab[4729]"
//line /usr/local/go/src/math/big/float.go:1505
	}
//line /usr/local/go/src/math/big/float.go:1505
	// _ = "end of CoverTab[4700]"
//line /usr/local/go/src/math/big/float.go:1505
	_go_fuzz_dep_.CoverTab[4701]++

//line /usr/local/go/src/math/big/float.go:1509
	return z.Set(y)
//line /usr/local/go/src/math/big/float.go:1509
	// _ = "end of CoverTab[4701]"
}

//line /usr/local/go/src/math/big/float.go:1516
func (z *Float) Sub(x, y *Float) *Float {
//line /usr/local/go/src/math/big/float.go:1516
	_go_fuzz_dep_.CoverTab[4730]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1517
		_go_fuzz_dep_.CoverTab[4737]++
								x.validate()
								y.validate()
//line /usr/local/go/src/math/big/float.go:1519
		// _ = "end of CoverTab[4737]"
	} else {
//line /usr/local/go/src/math/big/float.go:1520
		_go_fuzz_dep_.CoverTab[4738]++
//line /usr/local/go/src/math/big/float.go:1520
		// _ = "end of CoverTab[4738]"
//line /usr/local/go/src/math/big/float.go:1520
	}
//line /usr/local/go/src/math/big/float.go:1520
	// _ = "end of CoverTab[4730]"
//line /usr/local/go/src/math/big/float.go:1520
	_go_fuzz_dep_.CoverTab[4731]++

							if z.prec == 0 {
//line /usr/local/go/src/math/big/float.go:1522
		_go_fuzz_dep_.CoverTab[4739]++
								z.prec = umax32(x.prec, y.prec)
//line /usr/local/go/src/math/big/float.go:1523
		// _ = "end of CoverTab[4739]"
	} else {
//line /usr/local/go/src/math/big/float.go:1524
		_go_fuzz_dep_.CoverTab[4740]++
//line /usr/local/go/src/math/big/float.go:1524
		// _ = "end of CoverTab[4740]"
//line /usr/local/go/src/math/big/float.go:1524
	}
//line /usr/local/go/src/math/big/float.go:1524
	// _ = "end of CoverTab[4731]"
//line /usr/local/go/src/math/big/float.go:1524
	_go_fuzz_dep_.CoverTab[4732]++

							if x.form == finite && func() bool {
//line /usr/local/go/src/math/big/float.go:1526
		_go_fuzz_dep_.CoverTab[4741]++
//line /usr/local/go/src/math/big/float.go:1526
		return y.form == finite
//line /usr/local/go/src/math/big/float.go:1526
		// _ = "end of CoverTab[4741]"
//line /usr/local/go/src/math/big/float.go:1526
	}() {
//line /usr/local/go/src/math/big/float.go:1526
		_go_fuzz_dep_.CoverTab[4742]++

								yneg := y.neg
								z.neg = x.neg
								if x.neg != yneg {
//line /usr/local/go/src/math/big/float.go:1530
			_go_fuzz_dep_.CoverTab[4745]++

//line /usr/local/go/src/math/big/float.go:1533
			z.uadd(x, y)
//line /usr/local/go/src/math/big/float.go:1533
			// _ = "end of CoverTab[4745]"
		} else {
//line /usr/local/go/src/math/big/float.go:1534
			_go_fuzz_dep_.CoverTab[4746]++

//line /usr/local/go/src/math/big/float.go:1537
			if x.ucmp(y) > 0 {
//line /usr/local/go/src/math/big/float.go:1537
				_go_fuzz_dep_.CoverTab[4747]++
										z.usub(x, y)
//line /usr/local/go/src/math/big/float.go:1538
				// _ = "end of CoverTab[4747]"
			} else {
//line /usr/local/go/src/math/big/float.go:1539
				_go_fuzz_dep_.CoverTab[4748]++
										z.neg = !z.neg
										z.usub(y, x)
//line /usr/local/go/src/math/big/float.go:1541
				// _ = "end of CoverTab[4748]"
			}
//line /usr/local/go/src/math/big/float.go:1542
			// _ = "end of CoverTab[4746]"
		}
//line /usr/local/go/src/math/big/float.go:1543
		// _ = "end of CoverTab[4742]"
//line /usr/local/go/src/math/big/float.go:1543
		_go_fuzz_dep_.CoverTab[4743]++
								if z.form == zero && func() bool {
//line /usr/local/go/src/math/big/float.go:1544
			_go_fuzz_dep_.CoverTab[4749]++
//line /usr/local/go/src/math/big/float.go:1544
			return z.mode == ToNegativeInf
//line /usr/local/go/src/math/big/float.go:1544
			// _ = "end of CoverTab[4749]"
//line /usr/local/go/src/math/big/float.go:1544
		}() && func() bool {
//line /usr/local/go/src/math/big/float.go:1544
			_go_fuzz_dep_.CoverTab[4750]++
//line /usr/local/go/src/math/big/float.go:1544
			return z.acc == Exact
//line /usr/local/go/src/math/big/float.go:1544
			// _ = "end of CoverTab[4750]"
//line /usr/local/go/src/math/big/float.go:1544
		}() {
//line /usr/local/go/src/math/big/float.go:1544
			_go_fuzz_dep_.CoverTab[4751]++
									z.neg = true
//line /usr/local/go/src/math/big/float.go:1545
			// _ = "end of CoverTab[4751]"
		} else {
//line /usr/local/go/src/math/big/float.go:1546
			_go_fuzz_dep_.CoverTab[4752]++
//line /usr/local/go/src/math/big/float.go:1546
			// _ = "end of CoverTab[4752]"
//line /usr/local/go/src/math/big/float.go:1546
		}
//line /usr/local/go/src/math/big/float.go:1546
		// _ = "end of CoverTab[4743]"
//line /usr/local/go/src/math/big/float.go:1546
		_go_fuzz_dep_.CoverTab[4744]++
								return z
//line /usr/local/go/src/math/big/float.go:1547
		// _ = "end of CoverTab[4744]"
	} else {
//line /usr/local/go/src/math/big/float.go:1548
		_go_fuzz_dep_.CoverTab[4753]++
//line /usr/local/go/src/math/big/float.go:1548
		// _ = "end of CoverTab[4753]"
//line /usr/local/go/src/math/big/float.go:1548
	}
//line /usr/local/go/src/math/big/float.go:1548
	// _ = "end of CoverTab[4732]"
//line /usr/local/go/src/math/big/float.go:1548
	_go_fuzz_dep_.CoverTab[4733]++

							if x.form == inf && func() bool {
//line /usr/local/go/src/math/big/float.go:1550
		_go_fuzz_dep_.CoverTab[4754]++
//line /usr/local/go/src/math/big/float.go:1550
		return y.form == inf
//line /usr/local/go/src/math/big/float.go:1550
		// _ = "end of CoverTab[4754]"
//line /usr/local/go/src/math/big/float.go:1550
	}() && func() bool {
//line /usr/local/go/src/math/big/float.go:1550
		_go_fuzz_dep_.CoverTab[4755]++
//line /usr/local/go/src/math/big/float.go:1550
		return x.neg == y.neg
//line /usr/local/go/src/math/big/float.go:1550
		// _ = "end of CoverTab[4755]"
//line /usr/local/go/src/math/big/float.go:1550
	}() {
//line /usr/local/go/src/math/big/float.go:1550
		_go_fuzz_dep_.CoverTab[4756]++

//line /usr/local/go/src/math/big/float.go:1554
		z.acc = Exact
								z.form = zero
								z.neg = false
								panic(ErrNaN{"subtraction of infinities with equal signs"})
//line /usr/local/go/src/math/big/float.go:1557
		// _ = "end of CoverTab[4756]"
	} else {
//line /usr/local/go/src/math/big/float.go:1558
		_go_fuzz_dep_.CoverTab[4757]++
//line /usr/local/go/src/math/big/float.go:1558
		// _ = "end of CoverTab[4757]"
//line /usr/local/go/src/math/big/float.go:1558
	}
//line /usr/local/go/src/math/big/float.go:1558
	// _ = "end of CoverTab[4733]"
//line /usr/local/go/src/math/big/float.go:1558
	_go_fuzz_dep_.CoverTab[4734]++

							if x.form == zero && func() bool {
//line /usr/local/go/src/math/big/float.go:1560
		_go_fuzz_dep_.CoverTab[4758]++
//line /usr/local/go/src/math/big/float.go:1560
		return y.form == zero
//line /usr/local/go/src/math/big/float.go:1560
		// _ = "end of CoverTab[4758]"
//line /usr/local/go/src/math/big/float.go:1560
	}() {
//line /usr/local/go/src/math/big/float.go:1560
		_go_fuzz_dep_.CoverTab[4759]++

								z.acc = Exact
								z.form = zero
								z.neg = x.neg && func() bool {
//line /usr/local/go/src/math/big/float.go:1564
			_go_fuzz_dep_.CoverTab[4760]++
//line /usr/local/go/src/math/big/float.go:1564
			return !y.neg
//line /usr/local/go/src/math/big/float.go:1564
			// _ = "end of CoverTab[4760]"
//line /usr/local/go/src/math/big/float.go:1564
		}()
								return z
//line /usr/local/go/src/math/big/float.go:1565
		// _ = "end of CoverTab[4759]"
	} else {
//line /usr/local/go/src/math/big/float.go:1566
		_go_fuzz_dep_.CoverTab[4761]++
//line /usr/local/go/src/math/big/float.go:1566
		// _ = "end of CoverTab[4761]"
//line /usr/local/go/src/math/big/float.go:1566
	}
//line /usr/local/go/src/math/big/float.go:1566
	// _ = "end of CoverTab[4734]"
//line /usr/local/go/src/math/big/float.go:1566
	_go_fuzz_dep_.CoverTab[4735]++

							if x.form == inf || func() bool {
//line /usr/local/go/src/math/big/float.go:1568
		_go_fuzz_dep_.CoverTab[4762]++
//line /usr/local/go/src/math/big/float.go:1568
		return y.form == zero
//line /usr/local/go/src/math/big/float.go:1568
		// _ = "end of CoverTab[4762]"
//line /usr/local/go/src/math/big/float.go:1568
	}() {
//line /usr/local/go/src/math/big/float.go:1568
		_go_fuzz_dep_.CoverTab[4763]++

//line /usr/local/go/src/math/big/float.go:1571
		return z.Set(x)
//line /usr/local/go/src/math/big/float.go:1571
		// _ = "end of CoverTab[4763]"
	} else {
//line /usr/local/go/src/math/big/float.go:1572
		_go_fuzz_dep_.CoverTab[4764]++
//line /usr/local/go/src/math/big/float.go:1572
		// _ = "end of CoverTab[4764]"
//line /usr/local/go/src/math/big/float.go:1572
	}
//line /usr/local/go/src/math/big/float.go:1572
	// _ = "end of CoverTab[4735]"
//line /usr/local/go/src/math/big/float.go:1572
	_go_fuzz_dep_.CoverTab[4736]++

//line /usr/local/go/src/math/big/float.go:1576
	return z.Neg(y)
//line /usr/local/go/src/math/big/float.go:1576
	// _ = "end of CoverTab[4736]"
}

//line /usr/local/go/src/math/big/float.go:1583
func (z *Float) Mul(x, y *Float) *Float {
//line /usr/local/go/src/math/big/float.go:1583
	_go_fuzz_dep_.CoverTab[4765]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1584
		_go_fuzz_dep_.CoverTab[4771]++
								x.validate()
								y.validate()
//line /usr/local/go/src/math/big/float.go:1586
		// _ = "end of CoverTab[4771]"
	} else {
//line /usr/local/go/src/math/big/float.go:1587
		_go_fuzz_dep_.CoverTab[4772]++
//line /usr/local/go/src/math/big/float.go:1587
		// _ = "end of CoverTab[4772]"
//line /usr/local/go/src/math/big/float.go:1587
	}
//line /usr/local/go/src/math/big/float.go:1587
	// _ = "end of CoverTab[4765]"
//line /usr/local/go/src/math/big/float.go:1587
	_go_fuzz_dep_.CoverTab[4766]++

							if z.prec == 0 {
//line /usr/local/go/src/math/big/float.go:1589
		_go_fuzz_dep_.CoverTab[4773]++
								z.prec = umax32(x.prec, y.prec)
//line /usr/local/go/src/math/big/float.go:1590
		// _ = "end of CoverTab[4773]"
	} else {
//line /usr/local/go/src/math/big/float.go:1591
		_go_fuzz_dep_.CoverTab[4774]++
//line /usr/local/go/src/math/big/float.go:1591
		// _ = "end of CoverTab[4774]"
//line /usr/local/go/src/math/big/float.go:1591
	}
//line /usr/local/go/src/math/big/float.go:1591
	// _ = "end of CoverTab[4766]"
//line /usr/local/go/src/math/big/float.go:1591
	_go_fuzz_dep_.CoverTab[4767]++

							z.neg = x.neg != y.neg

							if x.form == finite && func() bool {
//line /usr/local/go/src/math/big/float.go:1595
		_go_fuzz_dep_.CoverTab[4775]++
//line /usr/local/go/src/math/big/float.go:1595
		return y.form == finite
//line /usr/local/go/src/math/big/float.go:1595
		// _ = "end of CoverTab[4775]"
//line /usr/local/go/src/math/big/float.go:1595
	}() {
//line /usr/local/go/src/math/big/float.go:1595
		_go_fuzz_dep_.CoverTab[4776]++

								z.umul(x, y)
								return z
//line /usr/local/go/src/math/big/float.go:1598
		// _ = "end of CoverTab[4776]"
	} else {
//line /usr/local/go/src/math/big/float.go:1599
		_go_fuzz_dep_.CoverTab[4777]++
//line /usr/local/go/src/math/big/float.go:1599
		// _ = "end of CoverTab[4777]"
//line /usr/local/go/src/math/big/float.go:1599
	}
//line /usr/local/go/src/math/big/float.go:1599
	// _ = "end of CoverTab[4767]"
//line /usr/local/go/src/math/big/float.go:1599
	_go_fuzz_dep_.CoverTab[4768]++

							z.acc = Exact
							if x.form == zero && func() bool {
//line /usr/local/go/src/math/big/float.go:1602
		_go_fuzz_dep_.CoverTab[4778]++
//line /usr/local/go/src/math/big/float.go:1602
		return y.form == inf
//line /usr/local/go/src/math/big/float.go:1602
		// _ = "end of CoverTab[4778]"
//line /usr/local/go/src/math/big/float.go:1602
	}() || func() bool {
//line /usr/local/go/src/math/big/float.go:1602
		_go_fuzz_dep_.CoverTab[4779]++
//line /usr/local/go/src/math/big/float.go:1602
		return x.form == inf && func() bool {
//line /usr/local/go/src/math/big/float.go:1602
			_go_fuzz_dep_.CoverTab[4780]++
//line /usr/local/go/src/math/big/float.go:1602
			return y.form == zero
//line /usr/local/go/src/math/big/float.go:1602
			// _ = "end of CoverTab[4780]"
//line /usr/local/go/src/math/big/float.go:1602
		}()
//line /usr/local/go/src/math/big/float.go:1602
		// _ = "end of CoverTab[4779]"
//line /usr/local/go/src/math/big/float.go:1602
	}() {
//line /usr/local/go/src/math/big/float.go:1602
		_go_fuzz_dep_.CoverTab[4781]++

//line /usr/local/go/src/math/big/float.go:1606
		z.form = zero
								z.neg = false
								panic(ErrNaN{"multiplication of zero with infinity"})
//line /usr/local/go/src/math/big/float.go:1608
		// _ = "end of CoverTab[4781]"
	} else {
//line /usr/local/go/src/math/big/float.go:1609
		_go_fuzz_dep_.CoverTab[4782]++
//line /usr/local/go/src/math/big/float.go:1609
		// _ = "end of CoverTab[4782]"
//line /usr/local/go/src/math/big/float.go:1609
	}
//line /usr/local/go/src/math/big/float.go:1609
	// _ = "end of CoverTab[4768]"
//line /usr/local/go/src/math/big/float.go:1609
	_go_fuzz_dep_.CoverTab[4769]++

							if x.form == inf || func() bool {
//line /usr/local/go/src/math/big/float.go:1611
		_go_fuzz_dep_.CoverTab[4783]++
//line /usr/local/go/src/math/big/float.go:1611
		return y.form == inf
//line /usr/local/go/src/math/big/float.go:1611
		// _ = "end of CoverTab[4783]"
//line /usr/local/go/src/math/big/float.go:1611
	}() {
//line /usr/local/go/src/math/big/float.go:1611
		_go_fuzz_dep_.CoverTab[4784]++

//line /usr/local/go/src/math/big/float.go:1614
		z.form = inf
								return z
//line /usr/local/go/src/math/big/float.go:1615
		// _ = "end of CoverTab[4784]"
	} else {
//line /usr/local/go/src/math/big/float.go:1616
		_go_fuzz_dep_.CoverTab[4785]++
//line /usr/local/go/src/math/big/float.go:1616
		// _ = "end of CoverTab[4785]"
//line /usr/local/go/src/math/big/float.go:1616
	}
//line /usr/local/go/src/math/big/float.go:1616
	// _ = "end of CoverTab[4769]"
//line /usr/local/go/src/math/big/float.go:1616
	_go_fuzz_dep_.CoverTab[4770]++

//line /usr/local/go/src/math/big/float.go:1620
	z.form = zero
							return z
//line /usr/local/go/src/math/big/float.go:1621
	// _ = "end of CoverTab[4770]"
}

//line /usr/local/go/src/math/big/float.go:1628
func (z *Float) Quo(x, y *Float) *Float {
//line /usr/local/go/src/math/big/float.go:1628
	_go_fuzz_dep_.CoverTab[4786]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1629
		_go_fuzz_dep_.CoverTab[4792]++
								x.validate()
								y.validate()
//line /usr/local/go/src/math/big/float.go:1631
		// _ = "end of CoverTab[4792]"
	} else {
//line /usr/local/go/src/math/big/float.go:1632
		_go_fuzz_dep_.CoverTab[4793]++
//line /usr/local/go/src/math/big/float.go:1632
		// _ = "end of CoverTab[4793]"
//line /usr/local/go/src/math/big/float.go:1632
	}
//line /usr/local/go/src/math/big/float.go:1632
	// _ = "end of CoverTab[4786]"
//line /usr/local/go/src/math/big/float.go:1632
	_go_fuzz_dep_.CoverTab[4787]++

							if z.prec == 0 {
//line /usr/local/go/src/math/big/float.go:1634
		_go_fuzz_dep_.CoverTab[4794]++
								z.prec = umax32(x.prec, y.prec)
//line /usr/local/go/src/math/big/float.go:1635
		// _ = "end of CoverTab[4794]"
	} else {
//line /usr/local/go/src/math/big/float.go:1636
		_go_fuzz_dep_.CoverTab[4795]++
//line /usr/local/go/src/math/big/float.go:1636
		// _ = "end of CoverTab[4795]"
//line /usr/local/go/src/math/big/float.go:1636
	}
//line /usr/local/go/src/math/big/float.go:1636
	// _ = "end of CoverTab[4787]"
//line /usr/local/go/src/math/big/float.go:1636
	_go_fuzz_dep_.CoverTab[4788]++

							z.neg = x.neg != y.neg

							if x.form == finite && func() bool {
//line /usr/local/go/src/math/big/float.go:1640
		_go_fuzz_dep_.CoverTab[4796]++
//line /usr/local/go/src/math/big/float.go:1640
		return y.form == finite
//line /usr/local/go/src/math/big/float.go:1640
		// _ = "end of CoverTab[4796]"
//line /usr/local/go/src/math/big/float.go:1640
	}() {
//line /usr/local/go/src/math/big/float.go:1640
		_go_fuzz_dep_.CoverTab[4797]++

								z.uquo(x, y)
								return z
//line /usr/local/go/src/math/big/float.go:1643
		// _ = "end of CoverTab[4797]"
	} else {
//line /usr/local/go/src/math/big/float.go:1644
		_go_fuzz_dep_.CoverTab[4798]++
//line /usr/local/go/src/math/big/float.go:1644
		// _ = "end of CoverTab[4798]"
//line /usr/local/go/src/math/big/float.go:1644
	}
//line /usr/local/go/src/math/big/float.go:1644
	// _ = "end of CoverTab[4788]"
//line /usr/local/go/src/math/big/float.go:1644
	_go_fuzz_dep_.CoverTab[4789]++

							z.acc = Exact
							if x.form == zero && func() bool {
//line /usr/local/go/src/math/big/float.go:1647
		_go_fuzz_dep_.CoverTab[4799]++
//line /usr/local/go/src/math/big/float.go:1647
		return y.form == zero
//line /usr/local/go/src/math/big/float.go:1647
		// _ = "end of CoverTab[4799]"
//line /usr/local/go/src/math/big/float.go:1647
	}() || func() bool {
//line /usr/local/go/src/math/big/float.go:1647
		_go_fuzz_dep_.CoverTab[4800]++
//line /usr/local/go/src/math/big/float.go:1647
		return x.form == inf && func() bool {
//line /usr/local/go/src/math/big/float.go:1647
			_go_fuzz_dep_.CoverTab[4801]++
//line /usr/local/go/src/math/big/float.go:1647
			return y.form == inf
//line /usr/local/go/src/math/big/float.go:1647
			// _ = "end of CoverTab[4801]"
//line /usr/local/go/src/math/big/float.go:1647
		}()
//line /usr/local/go/src/math/big/float.go:1647
		// _ = "end of CoverTab[4800]"
//line /usr/local/go/src/math/big/float.go:1647
	}() {
//line /usr/local/go/src/math/big/float.go:1647
		_go_fuzz_dep_.CoverTab[4802]++

//line /usr/local/go/src/math/big/float.go:1651
		z.form = zero
								z.neg = false
								panic(ErrNaN{"division of zero by zero or infinity by infinity"})
//line /usr/local/go/src/math/big/float.go:1653
		// _ = "end of CoverTab[4802]"
	} else {
//line /usr/local/go/src/math/big/float.go:1654
		_go_fuzz_dep_.CoverTab[4803]++
//line /usr/local/go/src/math/big/float.go:1654
		// _ = "end of CoverTab[4803]"
//line /usr/local/go/src/math/big/float.go:1654
	}
//line /usr/local/go/src/math/big/float.go:1654
	// _ = "end of CoverTab[4789]"
//line /usr/local/go/src/math/big/float.go:1654
	_go_fuzz_dep_.CoverTab[4790]++

							if x.form == zero || func() bool {
//line /usr/local/go/src/math/big/float.go:1656
		_go_fuzz_dep_.CoverTab[4804]++
//line /usr/local/go/src/math/big/float.go:1656
		return y.form == inf
//line /usr/local/go/src/math/big/float.go:1656
		// _ = "end of CoverTab[4804]"
//line /usr/local/go/src/math/big/float.go:1656
	}() {
//line /usr/local/go/src/math/big/float.go:1656
		_go_fuzz_dep_.CoverTab[4805]++

//line /usr/local/go/src/math/big/float.go:1659
		z.form = zero
								return z
//line /usr/local/go/src/math/big/float.go:1660
		// _ = "end of CoverTab[4805]"
	} else {
//line /usr/local/go/src/math/big/float.go:1661
		_go_fuzz_dep_.CoverTab[4806]++
//line /usr/local/go/src/math/big/float.go:1661
		// _ = "end of CoverTab[4806]"
//line /usr/local/go/src/math/big/float.go:1661
	}
//line /usr/local/go/src/math/big/float.go:1661
	// _ = "end of CoverTab[4790]"
//line /usr/local/go/src/math/big/float.go:1661
	_go_fuzz_dep_.CoverTab[4791]++

//line /usr/local/go/src/math/big/float.go:1665
	z.form = inf
							return z
//line /usr/local/go/src/math/big/float.go:1666
	// _ = "end of CoverTab[4791]"
}

//line /usr/local/go/src/math/big/float.go:1674
func (x *Float) Cmp(y *Float) int {
//line /usr/local/go/src/math/big/float.go:1674
	_go_fuzz_dep_.CoverTab[4807]++
							if debugFloat {
//line /usr/local/go/src/math/big/float.go:1675
		_go_fuzz_dep_.CoverTab[4811]++
								x.validate()
								y.validate()
//line /usr/local/go/src/math/big/float.go:1677
		// _ = "end of CoverTab[4811]"
	} else {
//line /usr/local/go/src/math/big/float.go:1678
		_go_fuzz_dep_.CoverTab[4812]++
//line /usr/local/go/src/math/big/float.go:1678
		// _ = "end of CoverTab[4812]"
//line /usr/local/go/src/math/big/float.go:1678
	}
//line /usr/local/go/src/math/big/float.go:1678
	// _ = "end of CoverTab[4807]"
//line /usr/local/go/src/math/big/float.go:1678
	_go_fuzz_dep_.CoverTab[4808]++

							mx := x.ord()
							my := y.ord()
							switch {
	case mx < my:
//line /usr/local/go/src/math/big/float.go:1683
		_go_fuzz_dep_.CoverTab[4813]++
								return -1
//line /usr/local/go/src/math/big/float.go:1684
		// _ = "end of CoverTab[4813]"
	case mx > my:
//line /usr/local/go/src/math/big/float.go:1685
		_go_fuzz_dep_.CoverTab[4814]++
								return +1
//line /usr/local/go/src/math/big/float.go:1686
		// _ = "end of CoverTab[4814]"
//line /usr/local/go/src/math/big/float.go:1686
	default:
//line /usr/local/go/src/math/big/float.go:1686
		_go_fuzz_dep_.CoverTab[4815]++
//line /usr/local/go/src/math/big/float.go:1686
		// _ = "end of CoverTab[4815]"
	}
//line /usr/local/go/src/math/big/float.go:1687
	// _ = "end of CoverTab[4808]"
//line /usr/local/go/src/math/big/float.go:1687
	_go_fuzz_dep_.CoverTab[4809]++

//line /usr/local/go/src/math/big/float.go:1691
	switch mx {
	case -1:
//line /usr/local/go/src/math/big/float.go:1692
		_go_fuzz_dep_.CoverTab[4816]++
								return y.ucmp(x)
//line /usr/local/go/src/math/big/float.go:1693
		// _ = "end of CoverTab[4816]"
	case +1:
//line /usr/local/go/src/math/big/float.go:1694
		_go_fuzz_dep_.CoverTab[4817]++
								return x.ucmp(y)
//line /usr/local/go/src/math/big/float.go:1695
		// _ = "end of CoverTab[4817]"
//line /usr/local/go/src/math/big/float.go:1695
	default:
//line /usr/local/go/src/math/big/float.go:1695
		_go_fuzz_dep_.CoverTab[4818]++
//line /usr/local/go/src/math/big/float.go:1695
		// _ = "end of CoverTab[4818]"
	}
//line /usr/local/go/src/math/big/float.go:1696
	// _ = "end of CoverTab[4809]"
//line /usr/local/go/src/math/big/float.go:1696
	_go_fuzz_dep_.CoverTab[4810]++

							return 0
//line /usr/local/go/src/math/big/float.go:1698
	// _ = "end of CoverTab[4810]"
}

//line /usr/local/go/src/math/big/float.go:1708
func (x *Float) ord() int {
//line /usr/local/go/src/math/big/float.go:1708
	_go_fuzz_dep_.CoverTab[4819]++
							var m int
							switch x.form {
	case finite:
//line /usr/local/go/src/math/big/float.go:1711
		_go_fuzz_dep_.CoverTab[4822]++
								m = 1
//line /usr/local/go/src/math/big/float.go:1712
		// _ = "end of CoverTab[4822]"
	case zero:
//line /usr/local/go/src/math/big/float.go:1713
		_go_fuzz_dep_.CoverTab[4823]++
								return 0
//line /usr/local/go/src/math/big/float.go:1714
		// _ = "end of CoverTab[4823]"
	case inf:
//line /usr/local/go/src/math/big/float.go:1715
		_go_fuzz_dep_.CoverTab[4824]++
								m = 2
//line /usr/local/go/src/math/big/float.go:1716
		// _ = "end of CoverTab[4824]"
//line /usr/local/go/src/math/big/float.go:1716
	default:
//line /usr/local/go/src/math/big/float.go:1716
		_go_fuzz_dep_.CoverTab[4825]++
//line /usr/local/go/src/math/big/float.go:1716
		// _ = "end of CoverTab[4825]"
	}
//line /usr/local/go/src/math/big/float.go:1717
	// _ = "end of CoverTab[4819]"
//line /usr/local/go/src/math/big/float.go:1717
	_go_fuzz_dep_.CoverTab[4820]++
							if x.neg {
//line /usr/local/go/src/math/big/float.go:1718
		_go_fuzz_dep_.CoverTab[4826]++
								m = -m
//line /usr/local/go/src/math/big/float.go:1719
		// _ = "end of CoverTab[4826]"
	} else {
//line /usr/local/go/src/math/big/float.go:1720
		_go_fuzz_dep_.CoverTab[4827]++
//line /usr/local/go/src/math/big/float.go:1720
		// _ = "end of CoverTab[4827]"
//line /usr/local/go/src/math/big/float.go:1720
	}
//line /usr/local/go/src/math/big/float.go:1720
	// _ = "end of CoverTab[4820]"
//line /usr/local/go/src/math/big/float.go:1720
	_go_fuzz_dep_.CoverTab[4821]++
							return m
//line /usr/local/go/src/math/big/float.go:1721
	// _ = "end of CoverTab[4821]"
}

func umax32(x, y uint32) uint32 {
//line /usr/local/go/src/math/big/float.go:1724
	_go_fuzz_dep_.CoverTab[4828]++
							if x > y {
//line /usr/local/go/src/math/big/float.go:1725
		_go_fuzz_dep_.CoverTab[4830]++
								return x
//line /usr/local/go/src/math/big/float.go:1726
		// _ = "end of CoverTab[4830]"
	} else {
//line /usr/local/go/src/math/big/float.go:1727
		_go_fuzz_dep_.CoverTab[4831]++
//line /usr/local/go/src/math/big/float.go:1727
		// _ = "end of CoverTab[4831]"
//line /usr/local/go/src/math/big/float.go:1727
	}
//line /usr/local/go/src/math/big/float.go:1727
	// _ = "end of CoverTab[4828]"
//line /usr/local/go/src/math/big/float.go:1727
	_go_fuzz_dep_.CoverTab[4829]++
							return y
//line /usr/local/go/src/math/big/float.go:1728
	// _ = "end of CoverTab[4829]"
}

//line /usr/local/go/src/math/big/float.go:1729
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/float.go:1729
var _ = _go_fuzz_dep_.CoverTab
