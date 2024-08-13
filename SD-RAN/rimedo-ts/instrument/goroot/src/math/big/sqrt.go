// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/math/big/sqrt.go:5
package big

//line /usr/local/go/src/math/big/sqrt.go:5
import (
//line /usr/local/go/src/math/big/sqrt.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/sqrt.go:5
)
//line /usr/local/go/src/math/big/sqrt.go:5
import (
//line /usr/local/go/src/math/big/sqrt.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/sqrt.go:5
)

import (
	"math"
	"sync"
)

var threeOnce struct {
	sync.Once
	v	*Float
}

func three() *Float {
//line /usr/local/go/src/math/big/sqrt.go:17
	_go_fuzz_dep_.CoverTab[6912]++
						threeOnce.Do(func() {
//line /usr/local/go/src/math/big/sqrt.go:18
		_go_fuzz_dep_.CoverTab[6914]++
							threeOnce.v = NewFloat(3.0)
//line /usr/local/go/src/math/big/sqrt.go:19
		// _ = "end of CoverTab[6914]"
	})
//line /usr/local/go/src/math/big/sqrt.go:20
	// _ = "end of CoverTab[6912]"
//line /usr/local/go/src/math/big/sqrt.go:20
	_go_fuzz_dep_.CoverTab[6913]++
						return threeOnce.v
//line /usr/local/go/src/math/big/sqrt.go:21
	// _ = "end of CoverTab[6913]"
}

// Sqrt sets z to the rounded square root of x, and returns it.
//line /usr/local/go/src/math/big/sqrt.go:24
//
//line /usr/local/go/src/math/big/sqrt.go:24
// If z's precision is 0, it is changed to x's precision before the
//line /usr/local/go/src/math/big/sqrt.go:24
// operation. Rounding is performed according to z's precision and
//line /usr/local/go/src/math/big/sqrt.go:24
// rounding mode, but z's accuracy is not computed. Specifically, the
//line /usr/local/go/src/math/big/sqrt.go:24
// result of z.Acc() is undefined.
//line /usr/local/go/src/math/big/sqrt.go:24
//
//line /usr/local/go/src/math/big/sqrt.go:24
// The function panics if z < 0. The value of z is undefined in that
//line /usr/local/go/src/math/big/sqrt.go:24
// case.
//line /usr/local/go/src/math/big/sqrt.go:33
func (z *Float) Sqrt(x *Float) *Float {
//line /usr/local/go/src/math/big/sqrt.go:33
	_go_fuzz_dep_.CoverTab[6915]++
						if debugFloat {
//line /usr/local/go/src/math/big/sqrt.go:34
		_go_fuzz_dep_.CoverTab[6921]++
							x.validate()
//line /usr/local/go/src/math/big/sqrt.go:35
		// _ = "end of CoverTab[6921]"
	} else {
//line /usr/local/go/src/math/big/sqrt.go:36
		_go_fuzz_dep_.CoverTab[6922]++
//line /usr/local/go/src/math/big/sqrt.go:36
		// _ = "end of CoverTab[6922]"
//line /usr/local/go/src/math/big/sqrt.go:36
	}
//line /usr/local/go/src/math/big/sqrt.go:36
	// _ = "end of CoverTab[6915]"
//line /usr/local/go/src/math/big/sqrt.go:36
	_go_fuzz_dep_.CoverTab[6916]++

						if z.prec == 0 {
//line /usr/local/go/src/math/big/sqrt.go:38
		_go_fuzz_dep_.CoverTab[6923]++
							z.prec = x.prec
//line /usr/local/go/src/math/big/sqrt.go:39
		// _ = "end of CoverTab[6923]"
	} else {
//line /usr/local/go/src/math/big/sqrt.go:40
		_go_fuzz_dep_.CoverTab[6924]++
//line /usr/local/go/src/math/big/sqrt.go:40
		// _ = "end of CoverTab[6924]"
//line /usr/local/go/src/math/big/sqrt.go:40
	}
//line /usr/local/go/src/math/big/sqrt.go:40
	// _ = "end of CoverTab[6916]"
//line /usr/local/go/src/math/big/sqrt.go:40
	_go_fuzz_dep_.CoverTab[6917]++

						if x.Sign() == -1 {
//line /usr/local/go/src/math/big/sqrt.go:42
		_go_fuzz_dep_.CoverTab[6925]++

							panic(ErrNaN{"square root of negative operand"})
//line /usr/local/go/src/math/big/sqrt.go:44
		// _ = "end of CoverTab[6925]"
	} else {
//line /usr/local/go/src/math/big/sqrt.go:45
		_go_fuzz_dep_.CoverTab[6926]++
//line /usr/local/go/src/math/big/sqrt.go:45
		// _ = "end of CoverTab[6926]"
//line /usr/local/go/src/math/big/sqrt.go:45
	}
//line /usr/local/go/src/math/big/sqrt.go:45
	// _ = "end of CoverTab[6917]"
//line /usr/local/go/src/math/big/sqrt.go:45
	_go_fuzz_dep_.CoverTab[6918]++

//line /usr/local/go/src/math/big/sqrt.go:48
	if x.form != finite {
//line /usr/local/go/src/math/big/sqrt.go:48
		_go_fuzz_dep_.CoverTab[6927]++
							z.acc = Exact
							z.form = x.form
							z.neg = x.neg
							return z
//line /usr/local/go/src/math/big/sqrt.go:52
		// _ = "end of CoverTab[6927]"
	} else {
//line /usr/local/go/src/math/big/sqrt.go:53
		_go_fuzz_dep_.CoverTab[6928]++
//line /usr/local/go/src/math/big/sqrt.go:53
		// _ = "end of CoverTab[6928]"
//line /usr/local/go/src/math/big/sqrt.go:53
	}
//line /usr/local/go/src/math/big/sqrt.go:53
	// _ = "end of CoverTab[6918]"
//line /usr/local/go/src/math/big/sqrt.go:53
	_go_fuzz_dep_.CoverTab[6919]++

//line /usr/local/go/src/math/big/sqrt.go:58
	prec := z.prec
						b := x.MantExp(z)
						z.prec = prec

//line /usr/local/go/src/math/big/sqrt.go:66
	switch b % 2 {
	case 0:
//line /usr/local/go/src/math/big/sqrt.go:67
		_go_fuzz_dep_.CoverTab[6929]++
//line /usr/local/go/src/math/big/sqrt.go:67
		// _ = "end of CoverTab[6929]"

	case 1:
//line /usr/local/go/src/math/big/sqrt.go:69
		_go_fuzz_dep_.CoverTab[6930]++
							z.exp++
//line /usr/local/go/src/math/big/sqrt.go:70
		// _ = "end of CoverTab[6930]"
	case -1:
//line /usr/local/go/src/math/big/sqrt.go:71
		_go_fuzz_dep_.CoverTab[6931]++
							z.exp--
//line /usr/local/go/src/math/big/sqrt.go:72
		// _ = "end of CoverTab[6931]"
//line /usr/local/go/src/math/big/sqrt.go:72
	default:
//line /usr/local/go/src/math/big/sqrt.go:72
		_go_fuzz_dep_.CoverTab[6932]++
//line /usr/local/go/src/math/big/sqrt.go:72
		// _ = "end of CoverTab[6932]"
	}
//line /usr/local/go/src/math/big/sqrt.go:73
	// _ = "end of CoverTab[6919]"
//line /usr/local/go/src/math/big/sqrt.go:73
	_go_fuzz_dep_.CoverTab[6920]++

//line /usr/local/go/src/math/big/sqrt.go:78
	z.sqrtInverse(z)

//line /usr/local/go/src/math/big/sqrt.go:81
	return z.SetMantExp(z, b/2)
//line /usr/local/go/src/math/big/sqrt.go:81
	// _ = "end of CoverTab[6920]"
}

// Compute √x (to z.prec precision) by solving
//line /usr/local/go/src/math/big/sqrt.go:84
//
//line /usr/local/go/src/math/big/sqrt.go:84
//	1/t² - x = 0
//line /usr/local/go/src/math/big/sqrt.go:84
//
//line /usr/local/go/src/math/big/sqrt.go:84
// for t (using Newton's method), and then inverting.
//line /usr/local/go/src/math/big/sqrt.go:89
func (z *Float) sqrtInverse(x *Float) {
//line /usr/local/go/src/math/big/sqrt.go:89
	_go_fuzz_dep_.CoverTab[6933]++

//line /usr/local/go/src/math/big/sqrt.go:96
	u := newFloat(z.prec)
	v := newFloat(z.prec)
	three := three()
	ng := func(t *Float) *Float {
//line /usr/local/go/src/math/big/sqrt.go:99
		_go_fuzz_dep_.CoverTab[6936]++
							u.prec = t.prec
							v.prec = t.prec
							u.Mul(t, t)
							u.Mul(x, u)
							v.Sub(three, u)
							u.Mul(t, v)
							u.exp--
							return t.Set(u)
//line /usr/local/go/src/math/big/sqrt.go:107
		// _ = "end of CoverTab[6936]"
	}
//line /usr/local/go/src/math/big/sqrt.go:108
	// _ = "end of CoverTab[6933]"
//line /usr/local/go/src/math/big/sqrt.go:108
	_go_fuzz_dep_.CoverTab[6934]++

						xf, _ := x.Float64()
						sqi := newFloat(z.prec)
						sqi.SetFloat64(1 / math.Sqrt(xf))
						for prec := z.prec + 32; sqi.prec < prec; {
//line /usr/local/go/src/math/big/sqrt.go:113
		_go_fuzz_dep_.CoverTab[6937]++
							sqi.prec *= 2
							sqi = ng(sqi)
//line /usr/local/go/src/math/big/sqrt.go:115
		// _ = "end of CoverTab[6937]"
	}
//line /usr/local/go/src/math/big/sqrt.go:116
	// _ = "end of CoverTab[6934]"
//line /usr/local/go/src/math/big/sqrt.go:116
	_go_fuzz_dep_.CoverTab[6935]++

//line /usr/local/go/src/math/big/sqrt.go:120
	z.Mul(x, sqi)
//line /usr/local/go/src/math/big/sqrt.go:120
	// _ = "end of CoverTab[6935]"
}

// newFloat returns a new *Float with space for twice the given
//line /usr/local/go/src/math/big/sqrt.go:123
// precision.
//line /usr/local/go/src/math/big/sqrt.go:125
func newFloat(prec2 uint32) *Float {
//line /usr/local/go/src/math/big/sqrt.go:125
	_go_fuzz_dep_.CoverTab[6938]++
						z := new(Float)

						z.mant = z.mant.make(int(prec2/_W) * 2)
						return z
//line /usr/local/go/src/math/big/sqrt.go:129
	// _ = "end of CoverTab[6938]"
}

//line /usr/local/go/src/math/big/sqrt.go:130
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/sqrt.go:130
var _ = _go_fuzz_dep_.CoverTab
