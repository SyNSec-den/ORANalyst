// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements string-to-Float conversion functions.

//line /usr/local/go/src/math/big/floatconv.go:7
package big

//line /usr/local/go/src/math/big/floatconv.go:7
import (
//line /usr/local/go/src/math/big/floatconv.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/floatconv.go:7
)
//line /usr/local/go/src/math/big/floatconv.go:7
import (
//line /usr/local/go/src/math/big/floatconv.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/floatconv.go:7
)

import (
	"fmt"
	"io"
	"strings"
)

var floatZero Float

// SetString sets z to the value of s and returns z and a boolean indicating
//line /usr/local/go/src/math/big/floatconv.go:17
// success. s must be a floating-point number of the same format as accepted
//line /usr/local/go/src/math/big/floatconv.go:17
// by Parse, with base argument 0. The entire string (not just a prefix) must
//line /usr/local/go/src/math/big/floatconv.go:17
// be valid for success. If the operation failed, the value of z is undefined
//line /usr/local/go/src/math/big/floatconv.go:17
// but the returned value is nil.
//line /usr/local/go/src/math/big/floatconv.go:22
func (z *Float) SetString(s string) (*Float, bool) {
//line /usr/local/go/src/math/big/floatconv.go:22
	_go_fuzz_dep_.CoverTab[4832]++
							if f, _, err := z.Parse(s, 0); err == nil {
//line /usr/local/go/src/math/big/floatconv.go:23
		_go_fuzz_dep_.CoverTab[4834]++
								return f, true
//line /usr/local/go/src/math/big/floatconv.go:24
		// _ = "end of CoverTab[4834]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:25
		_go_fuzz_dep_.CoverTab[4835]++
//line /usr/local/go/src/math/big/floatconv.go:25
		// _ = "end of CoverTab[4835]"
//line /usr/local/go/src/math/big/floatconv.go:25
	}
//line /usr/local/go/src/math/big/floatconv.go:25
	// _ = "end of CoverTab[4832]"
//line /usr/local/go/src/math/big/floatconv.go:25
	_go_fuzz_dep_.CoverTab[4833]++
							return nil, false
//line /usr/local/go/src/math/big/floatconv.go:26
	// _ = "end of CoverTab[4833]"
}

// scan is like Parse but reads the longest possible prefix representing a valid
//line /usr/local/go/src/math/big/floatconv.go:29
// floating point number from an io.ByteScanner rather than a string. It serves
//line /usr/local/go/src/math/big/floatconv.go:29
// as the implementation of Parse. It does not recognize ±Inf and does not expect
//line /usr/local/go/src/math/big/floatconv.go:29
// EOF at the end.
//line /usr/local/go/src/math/big/floatconv.go:33
func (z *Float) scan(r io.ByteScanner, base int) (f *Float, b int, err error) {
//line /usr/local/go/src/math/big/floatconv.go:33
	_go_fuzz_dep_.CoverTab[4836]++
							prec := z.prec
							if prec == 0 {
//line /usr/local/go/src/math/big/floatconv.go:35
		_go_fuzz_dep_.CoverTab[4847]++
								prec = 64
//line /usr/local/go/src/math/big/floatconv.go:36
		// _ = "end of CoverTab[4847]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:37
		_go_fuzz_dep_.CoverTab[4848]++
//line /usr/local/go/src/math/big/floatconv.go:37
		// _ = "end of CoverTab[4848]"
//line /usr/local/go/src/math/big/floatconv.go:37
	}
//line /usr/local/go/src/math/big/floatconv.go:37
	// _ = "end of CoverTab[4836]"
//line /usr/local/go/src/math/big/floatconv.go:37
	_go_fuzz_dep_.CoverTab[4837]++

//line /usr/local/go/src/math/big/floatconv.go:40
	z.form = zero

//line /usr/local/go/src/math/big/floatconv.go:43
	z.neg, err = scanSign(r)
	if err != nil {
//line /usr/local/go/src/math/big/floatconv.go:44
		_go_fuzz_dep_.CoverTab[4849]++
								return
//line /usr/local/go/src/math/big/floatconv.go:45
		// _ = "end of CoverTab[4849]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:46
		_go_fuzz_dep_.CoverTab[4850]++
//line /usr/local/go/src/math/big/floatconv.go:46
		// _ = "end of CoverTab[4850]"
//line /usr/local/go/src/math/big/floatconv.go:46
	}
//line /usr/local/go/src/math/big/floatconv.go:46
	// _ = "end of CoverTab[4837]"
//line /usr/local/go/src/math/big/floatconv.go:46
	_go_fuzz_dep_.CoverTab[4838]++

	// mantissa
	var fcount int	// fractional digit count; valid if <= 0
	z.mant, b, fcount, err = z.mant.scan(r, base, true)
	if err != nil {
//line /usr/local/go/src/math/big/floatconv.go:51
		_go_fuzz_dep_.CoverTab[4851]++
								return
//line /usr/local/go/src/math/big/floatconv.go:52
		// _ = "end of CoverTab[4851]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:53
		_go_fuzz_dep_.CoverTab[4852]++
//line /usr/local/go/src/math/big/floatconv.go:53
		// _ = "end of CoverTab[4852]"
//line /usr/local/go/src/math/big/floatconv.go:53
	}
//line /usr/local/go/src/math/big/floatconv.go:53
	// _ = "end of CoverTab[4838]"
//line /usr/local/go/src/math/big/floatconv.go:53
	_go_fuzz_dep_.CoverTab[4839]++

	// exponent
	var exp int64
	var ebase int
	exp, ebase, err = scanExponent(r, true, base == 0)
	if err != nil {
//line /usr/local/go/src/math/big/floatconv.go:59
		_go_fuzz_dep_.CoverTab[4853]++
								return
//line /usr/local/go/src/math/big/floatconv.go:60
		// _ = "end of CoverTab[4853]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:61
		_go_fuzz_dep_.CoverTab[4854]++
//line /usr/local/go/src/math/big/floatconv.go:61
		// _ = "end of CoverTab[4854]"
//line /usr/local/go/src/math/big/floatconv.go:61
	}
//line /usr/local/go/src/math/big/floatconv.go:61
	// _ = "end of CoverTab[4839]"
//line /usr/local/go/src/math/big/floatconv.go:61
	_go_fuzz_dep_.CoverTab[4840]++

//line /usr/local/go/src/math/big/floatconv.go:64
	if len(z.mant) == 0 {
//line /usr/local/go/src/math/big/floatconv.go:64
		_go_fuzz_dep_.CoverTab[4855]++
								z.prec = prec
								z.acc = Exact
								z.form = zero
								f = z
								return
//line /usr/local/go/src/math/big/floatconv.go:69
		// _ = "end of CoverTab[4855]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:70
		_go_fuzz_dep_.CoverTab[4856]++
//line /usr/local/go/src/math/big/floatconv.go:70
		// _ = "end of CoverTab[4856]"
//line /usr/local/go/src/math/big/floatconv.go:70
	}
//line /usr/local/go/src/math/big/floatconv.go:70
	// _ = "end of CoverTab[4840]"
//line /usr/local/go/src/math/big/floatconv.go:70
	_go_fuzz_dep_.CoverTab[4841]++

//line /usr/local/go/src/math/big/floatconv.go:85
	exp2 := int64(len(z.mant))*_W - fnorm(z.mant)
							exp5 := int64(0)

//line /usr/local/go/src/math/big/floatconv.go:89
	if fcount < 0 {
//line /usr/local/go/src/math/big/floatconv.go:89
		_go_fuzz_dep_.CoverTab[4857]++

//line /usr/local/go/src/math/big/floatconv.go:93
		d := int64(fcount)
		switch b {
		case 10:
//line /usr/local/go/src/math/big/floatconv.go:95
			_go_fuzz_dep_.CoverTab[4858]++
									exp5 = d
									fallthrough
//line /usr/local/go/src/math/big/floatconv.go:97
			// _ = "end of CoverTab[4858]"
		case 2:
//line /usr/local/go/src/math/big/floatconv.go:98
			_go_fuzz_dep_.CoverTab[4859]++
									exp2 += d
//line /usr/local/go/src/math/big/floatconv.go:99
			// _ = "end of CoverTab[4859]"
		case 8:
//line /usr/local/go/src/math/big/floatconv.go:100
			_go_fuzz_dep_.CoverTab[4860]++
									exp2 += d * 3
//line /usr/local/go/src/math/big/floatconv.go:101
			// _ = "end of CoverTab[4860]"
		case 16:
//line /usr/local/go/src/math/big/floatconv.go:102
			_go_fuzz_dep_.CoverTab[4861]++
									exp2 += d * 4
//line /usr/local/go/src/math/big/floatconv.go:103
			// _ = "end of CoverTab[4861]"
		default:
//line /usr/local/go/src/math/big/floatconv.go:104
			_go_fuzz_dep_.CoverTab[4862]++
									panic("unexpected mantissa base")
//line /usr/local/go/src/math/big/floatconv.go:105
			// _ = "end of CoverTab[4862]"
		}
//line /usr/local/go/src/math/big/floatconv.go:106
		// _ = "end of CoverTab[4857]"

	} else {
//line /usr/local/go/src/math/big/floatconv.go:108
		_go_fuzz_dep_.CoverTab[4863]++
//line /usr/local/go/src/math/big/floatconv.go:108
		// _ = "end of CoverTab[4863]"
//line /usr/local/go/src/math/big/floatconv.go:108
	}
//line /usr/local/go/src/math/big/floatconv.go:108
	// _ = "end of CoverTab[4841]"
//line /usr/local/go/src/math/big/floatconv.go:108
	_go_fuzz_dep_.CoverTab[4842]++

//line /usr/local/go/src/math/big/floatconv.go:111
	switch ebase {
	case 10:
//line /usr/local/go/src/math/big/floatconv.go:112
		_go_fuzz_dep_.CoverTab[4864]++
								exp5 += exp
								fallthrough
//line /usr/local/go/src/math/big/floatconv.go:114
		// _ = "end of CoverTab[4864]"
	case 2:
//line /usr/local/go/src/math/big/floatconv.go:115
		_go_fuzz_dep_.CoverTab[4865]++
								exp2 += exp
//line /usr/local/go/src/math/big/floatconv.go:116
		// _ = "end of CoverTab[4865]"
	default:
//line /usr/local/go/src/math/big/floatconv.go:117
		_go_fuzz_dep_.CoverTab[4866]++
								panic("unexpected exponent base")
//line /usr/local/go/src/math/big/floatconv.go:118
		// _ = "end of CoverTab[4866]"
	}
//line /usr/local/go/src/math/big/floatconv.go:119
	// _ = "end of CoverTab[4842]"
//line /usr/local/go/src/math/big/floatconv.go:119
	_go_fuzz_dep_.CoverTab[4843]++

//line /usr/local/go/src/math/big/floatconv.go:123
	if MinExp <= exp2 && func() bool {
//line /usr/local/go/src/math/big/floatconv.go:123
		_go_fuzz_dep_.CoverTab[4867]++
//line /usr/local/go/src/math/big/floatconv.go:123
		return exp2 <= MaxExp
//line /usr/local/go/src/math/big/floatconv.go:123
		// _ = "end of CoverTab[4867]"
//line /usr/local/go/src/math/big/floatconv.go:123
	}() {
//line /usr/local/go/src/math/big/floatconv.go:123
		_go_fuzz_dep_.CoverTab[4868]++
								z.prec = prec
								z.form = finite
								z.exp = int32(exp2)
								f = z
//line /usr/local/go/src/math/big/floatconv.go:127
		// _ = "end of CoverTab[4868]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:128
		_go_fuzz_dep_.CoverTab[4869]++
								err = fmt.Errorf("exponent overflow")
								return
//line /usr/local/go/src/math/big/floatconv.go:130
		// _ = "end of CoverTab[4869]"
	}
//line /usr/local/go/src/math/big/floatconv.go:131
	// _ = "end of CoverTab[4843]"
//line /usr/local/go/src/math/big/floatconv.go:131
	_go_fuzz_dep_.CoverTab[4844]++

							if exp5 == 0 {
//line /usr/local/go/src/math/big/floatconv.go:133
		_go_fuzz_dep_.CoverTab[4870]++

								z.round(0)
								return
//line /usr/local/go/src/math/big/floatconv.go:136
		// _ = "end of CoverTab[4870]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:137
		_go_fuzz_dep_.CoverTab[4871]++
//line /usr/local/go/src/math/big/floatconv.go:137
		// _ = "end of CoverTab[4871]"
//line /usr/local/go/src/math/big/floatconv.go:137
	}
//line /usr/local/go/src/math/big/floatconv.go:137
	// _ = "end of CoverTab[4844]"
//line /usr/local/go/src/math/big/floatconv.go:137
	_go_fuzz_dep_.CoverTab[4845]++

//line /usr/local/go/src/math/big/floatconv.go:141
	p := new(Float).SetPrec(z.Prec() + 64)
	if exp5 < 0 {
//line /usr/local/go/src/math/big/floatconv.go:142
		_go_fuzz_dep_.CoverTab[4872]++
								z.Quo(z, p.pow5(uint64(-exp5)))
//line /usr/local/go/src/math/big/floatconv.go:143
		// _ = "end of CoverTab[4872]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:144
		_go_fuzz_dep_.CoverTab[4873]++
								z.Mul(z, p.pow5(uint64(exp5)))
//line /usr/local/go/src/math/big/floatconv.go:145
		// _ = "end of CoverTab[4873]"
	}
//line /usr/local/go/src/math/big/floatconv.go:146
	// _ = "end of CoverTab[4845]"
//line /usr/local/go/src/math/big/floatconv.go:146
	_go_fuzz_dep_.CoverTab[4846]++

							return
//line /usr/local/go/src/math/big/floatconv.go:148
	// _ = "end of CoverTab[4846]"
}

// These powers of 5 fit into a uint64.
//line /usr/local/go/src/math/big/floatconv.go:151
//
//line /usr/local/go/src/math/big/floatconv.go:151
//	for p, q := uint64(0), uint64(1); p < q; p, q = q, q*5 {
//line /usr/local/go/src/math/big/floatconv.go:151
//		fmt.Println(q)
//line /usr/local/go/src/math/big/floatconv.go:151
//	}
//line /usr/local/go/src/math/big/floatconv.go:156
var pow5tab = [...]uint64{
	1,
	5,
	25,
	125,
	625,
	3125,
	15625,
	78125,
	390625,
	1953125,
	9765625,
	48828125,
	244140625,
	1220703125,
	6103515625,
	30517578125,
	152587890625,
	762939453125,
	3814697265625,
	19073486328125,
	95367431640625,
	476837158203125,
	2384185791015625,
	11920928955078125,
	59604644775390625,
	298023223876953125,
	1490116119384765625,
	7450580596923828125,
}

// pow5 sets z to 5**n and returns z.
//line /usr/local/go/src/math/big/floatconv.go:187
// n must not be negative.
//line /usr/local/go/src/math/big/floatconv.go:189
func (z *Float) pow5(n uint64) *Float {
//line /usr/local/go/src/math/big/floatconv.go:189
	_go_fuzz_dep_.CoverTab[4874]++
							const m = uint64(len(pow5tab) - 1)
							if n <= m {
//line /usr/local/go/src/math/big/floatconv.go:191
		_go_fuzz_dep_.CoverTab[4877]++
								return z.SetUint64(pow5tab[n])
//line /usr/local/go/src/math/big/floatconv.go:192
		// _ = "end of CoverTab[4877]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:193
		_go_fuzz_dep_.CoverTab[4878]++
//line /usr/local/go/src/math/big/floatconv.go:193
		// _ = "end of CoverTab[4878]"
//line /usr/local/go/src/math/big/floatconv.go:193
	}
//line /usr/local/go/src/math/big/floatconv.go:193
	// _ = "end of CoverTab[4874]"
//line /usr/local/go/src/math/big/floatconv.go:193
	_go_fuzz_dep_.CoverTab[4875]++

//line /usr/local/go/src/math/big/floatconv.go:196
	z.SetUint64(pow5tab[m])
							n -= m

//line /usr/local/go/src/math/big/floatconv.go:201
	f := new(Float).SetPrec(z.Prec() + 64).SetUint64(5)

	for n > 0 {
//line /usr/local/go/src/math/big/floatconv.go:203
		_go_fuzz_dep_.CoverTab[4879]++
								if n&1 != 0 {
//line /usr/local/go/src/math/big/floatconv.go:204
			_go_fuzz_dep_.CoverTab[4881]++
									z.Mul(z, f)
//line /usr/local/go/src/math/big/floatconv.go:205
			// _ = "end of CoverTab[4881]"
		} else {
//line /usr/local/go/src/math/big/floatconv.go:206
			_go_fuzz_dep_.CoverTab[4882]++
//line /usr/local/go/src/math/big/floatconv.go:206
			// _ = "end of CoverTab[4882]"
//line /usr/local/go/src/math/big/floatconv.go:206
		}
//line /usr/local/go/src/math/big/floatconv.go:206
		// _ = "end of CoverTab[4879]"
//line /usr/local/go/src/math/big/floatconv.go:206
		_go_fuzz_dep_.CoverTab[4880]++
								f.Mul(f, f)
								n >>= 1
//line /usr/local/go/src/math/big/floatconv.go:208
		// _ = "end of CoverTab[4880]"
	}
//line /usr/local/go/src/math/big/floatconv.go:209
	// _ = "end of CoverTab[4875]"
//line /usr/local/go/src/math/big/floatconv.go:209
	_go_fuzz_dep_.CoverTab[4876]++

							return z
//line /usr/local/go/src/math/big/floatconv.go:211
	// _ = "end of CoverTab[4876]"
}

// Parse parses s which must contain a text representation of a floating-
//line /usr/local/go/src/math/big/floatconv.go:214
// point number with a mantissa in the given conversion base (the exponent
//line /usr/local/go/src/math/big/floatconv.go:214
// is always a decimal number), or a string representing an infinite value.
//line /usr/local/go/src/math/big/floatconv.go:214
//
//line /usr/local/go/src/math/big/floatconv.go:214
// For base 0, an underscore character “_” may appear between a base
//line /usr/local/go/src/math/big/floatconv.go:214
// prefix and an adjacent digit, and between successive digits; such
//line /usr/local/go/src/math/big/floatconv.go:214
// underscores do not change the value of the number, or the returned
//line /usr/local/go/src/math/big/floatconv.go:214
// digit count. Incorrect placement of underscores is reported as an
//line /usr/local/go/src/math/big/floatconv.go:214
// error if there are no other errors. If base != 0, underscores are
//line /usr/local/go/src/math/big/floatconv.go:214
// not recognized and thus terminate scanning like any other character
//line /usr/local/go/src/math/big/floatconv.go:214
// that is not a valid radix point or digit.
//line /usr/local/go/src/math/big/floatconv.go:214
//
//line /usr/local/go/src/math/big/floatconv.go:214
// It sets z to the (possibly rounded) value of the corresponding floating-
//line /usr/local/go/src/math/big/floatconv.go:214
// point value, and returns z, the actual base b, and an error err, if any.
//line /usr/local/go/src/math/big/floatconv.go:214
// The entire string (not just a prefix) must be consumed for success.
//line /usr/local/go/src/math/big/floatconv.go:214
// If z's precision is 0, it is changed to 64 before rounding takes effect.
//line /usr/local/go/src/math/big/floatconv.go:214
// The number must be of the form:
//line /usr/local/go/src/math/big/floatconv.go:214
//
//line /usr/local/go/src/math/big/floatconv.go:214
//	number    = [ sign ] ( float | "inf" | "Inf" ) .
//line /usr/local/go/src/math/big/floatconv.go:214
//	sign      = "+" | "-" .
//line /usr/local/go/src/math/big/floatconv.go:214
//	float     = ( mantissa | prefix pmantissa ) [ exponent ] .
//line /usr/local/go/src/math/big/floatconv.go:214
//	prefix    = "0" [ "b" | "B" | "o" | "O" | "x" | "X" ] .
//line /usr/local/go/src/math/big/floatconv.go:214
//	mantissa  = digits "." [ digits ] | digits | "." digits .
//line /usr/local/go/src/math/big/floatconv.go:214
//	pmantissa = [ "_" ] digits "." [ digits ] | [ "_" ] digits | "." digits .
//line /usr/local/go/src/math/big/floatconv.go:214
//	exponent  = ( "e" | "E" | "p" | "P" ) [ sign ] digits .
//line /usr/local/go/src/math/big/floatconv.go:214
//	digits    = digit { [ "_" ] digit } .
//line /usr/local/go/src/math/big/floatconv.go:214
//	digit     = "0" ... "9" | "a" ... "z" | "A" ... "Z" .
//line /usr/local/go/src/math/big/floatconv.go:214
//
//line /usr/local/go/src/math/big/floatconv.go:214
// The base argument must be 0, 2, 8, 10, or 16. Providing an invalid base
//line /usr/local/go/src/math/big/floatconv.go:214
// argument will lead to a run-time panic.
//line /usr/local/go/src/math/big/floatconv.go:214
//
//line /usr/local/go/src/math/big/floatconv.go:214
// For base 0, the number prefix determines the actual base: A prefix of
//line /usr/local/go/src/math/big/floatconv.go:214
// “0b” or “0B” selects base 2, “0o” or “0O” selects base 8, and
//line /usr/local/go/src/math/big/floatconv.go:214
// “0x” or “0X” selects base 16. Otherwise, the actual base is 10 and
//line /usr/local/go/src/math/big/floatconv.go:214
// no prefix is accepted. The octal prefix "0" is not supported (a leading
//line /usr/local/go/src/math/big/floatconv.go:214
// "0" is simply considered a "0").
//line /usr/local/go/src/math/big/floatconv.go:214
//
//line /usr/local/go/src/math/big/floatconv.go:214
// A "p" or "P" exponent indicates a base 2 (rather then base 10) exponent;
//line /usr/local/go/src/math/big/floatconv.go:214
// for instance, "0x1.fffffffffffffp1023" (using base 0) represents the
//line /usr/local/go/src/math/big/floatconv.go:214
// maximum float64 value. For hexadecimal mantissae, the exponent character
//line /usr/local/go/src/math/big/floatconv.go:214
// must be one of 'p' or 'P', if present (an "e" or "E" exponent indicator
//line /usr/local/go/src/math/big/floatconv.go:214
// cannot be distinguished from a mantissa digit).
//line /usr/local/go/src/math/big/floatconv.go:214
//
//line /usr/local/go/src/math/big/floatconv.go:214
// The returned *Float f is nil and the value of z is valid but not
//line /usr/local/go/src/math/big/floatconv.go:214
// defined if an error is reported.
//line /usr/local/go/src/math/big/floatconv.go:259
func (z *Float) Parse(s string, base int) (f *Float, b int, err error) {
//line /usr/local/go/src/math/big/floatconv.go:259
	_go_fuzz_dep_.CoverTab[4883]++

							if len(s) == 3 && func() bool {
//line /usr/local/go/src/math/big/floatconv.go:261
		_go_fuzz_dep_.CoverTab[4888]++
//line /usr/local/go/src/math/big/floatconv.go:261
		return (s == "Inf" || func() bool {
//line /usr/local/go/src/math/big/floatconv.go:261
			_go_fuzz_dep_.CoverTab[4889]++
//line /usr/local/go/src/math/big/floatconv.go:261
			return s == "inf"
//line /usr/local/go/src/math/big/floatconv.go:261
			// _ = "end of CoverTab[4889]"
//line /usr/local/go/src/math/big/floatconv.go:261
		}())
//line /usr/local/go/src/math/big/floatconv.go:261
		// _ = "end of CoverTab[4888]"
//line /usr/local/go/src/math/big/floatconv.go:261
	}() {
//line /usr/local/go/src/math/big/floatconv.go:261
		_go_fuzz_dep_.CoverTab[4890]++
								f = z.SetInf(false)
								return
//line /usr/local/go/src/math/big/floatconv.go:263
		// _ = "end of CoverTab[4890]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:264
		_go_fuzz_dep_.CoverTab[4891]++
//line /usr/local/go/src/math/big/floatconv.go:264
		// _ = "end of CoverTab[4891]"
//line /usr/local/go/src/math/big/floatconv.go:264
	}
//line /usr/local/go/src/math/big/floatconv.go:264
	// _ = "end of CoverTab[4883]"
//line /usr/local/go/src/math/big/floatconv.go:264
	_go_fuzz_dep_.CoverTab[4884]++
							if len(s) == 4 && func() bool {
//line /usr/local/go/src/math/big/floatconv.go:265
		_go_fuzz_dep_.CoverTab[4892]++
//line /usr/local/go/src/math/big/floatconv.go:265
		return (s[0] == '+' || func() bool {
//line /usr/local/go/src/math/big/floatconv.go:265
			_go_fuzz_dep_.CoverTab[4893]++
//line /usr/local/go/src/math/big/floatconv.go:265
			return s[0] == '-'
//line /usr/local/go/src/math/big/floatconv.go:265
			// _ = "end of CoverTab[4893]"
//line /usr/local/go/src/math/big/floatconv.go:265
		}())
//line /usr/local/go/src/math/big/floatconv.go:265
		// _ = "end of CoverTab[4892]"
//line /usr/local/go/src/math/big/floatconv.go:265
	}() && func() bool {
//line /usr/local/go/src/math/big/floatconv.go:265
		_go_fuzz_dep_.CoverTab[4894]++
//line /usr/local/go/src/math/big/floatconv.go:265
		return (s[1:] == "Inf" || func() bool {
//line /usr/local/go/src/math/big/floatconv.go:265
			_go_fuzz_dep_.CoverTab[4895]++
//line /usr/local/go/src/math/big/floatconv.go:265
			return s[1:] == "inf"
//line /usr/local/go/src/math/big/floatconv.go:265
			// _ = "end of CoverTab[4895]"
//line /usr/local/go/src/math/big/floatconv.go:265
		}())
//line /usr/local/go/src/math/big/floatconv.go:265
		// _ = "end of CoverTab[4894]"
//line /usr/local/go/src/math/big/floatconv.go:265
	}() {
//line /usr/local/go/src/math/big/floatconv.go:265
		_go_fuzz_dep_.CoverTab[4896]++
								f = z.SetInf(s[0] == '-')
								return
//line /usr/local/go/src/math/big/floatconv.go:267
		// _ = "end of CoverTab[4896]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:268
		_go_fuzz_dep_.CoverTab[4897]++
//line /usr/local/go/src/math/big/floatconv.go:268
		// _ = "end of CoverTab[4897]"
//line /usr/local/go/src/math/big/floatconv.go:268
	}
//line /usr/local/go/src/math/big/floatconv.go:268
	// _ = "end of CoverTab[4884]"
//line /usr/local/go/src/math/big/floatconv.go:268
	_go_fuzz_dep_.CoverTab[4885]++

							r := strings.NewReader(s)
							if f, b, err = z.scan(r, base); err != nil {
//line /usr/local/go/src/math/big/floatconv.go:271
		_go_fuzz_dep_.CoverTab[4898]++
								return
//line /usr/local/go/src/math/big/floatconv.go:272
		// _ = "end of CoverTab[4898]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:273
		_go_fuzz_dep_.CoverTab[4899]++
//line /usr/local/go/src/math/big/floatconv.go:273
		// _ = "end of CoverTab[4899]"
//line /usr/local/go/src/math/big/floatconv.go:273
	}
//line /usr/local/go/src/math/big/floatconv.go:273
	// _ = "end of CoverTab[4885]"
//line /usr/local/go/src/math/big/floatconv.go:273
	_go_fuzz_dep_.CoverTab[4886]++

//line /usr/local/go/src/math/big/floatconv.go:276
	if ch, err2 := r.ReadByte(); err2 == nil {
//line /usr/local/go/src/math/big/floatconv.go:276
		_go_fuzz_dep_.CoverTab[4900]++
								err = fmt.Errorf("expected end of string, found %q", ch)
//line /usr/local/go/src/math/big/floatconv.go:277
		// _ = "end of CoverTab[4900]"
	} else {
//line /usr/local/go/src/math/big/floatconv.go:278
		_go_fuzz_dep_.CoverTab[4901]++
//line /usr/local/go/src/math/big/floatconv.go:278
		if err2 != io.EOF {
//line /usr/local/go/src/math/big/floatconv.go:278
			_go_fuzz_dep_.CoverTab[4902]++
									err = err2
//line /usr/local/go/src/math/big/floatconv.go:279
			// _ = "end of CoverTab[4902]"
		} else {
//line /usr/local/go/src/math/big/floatconv.go:280
			_go_fuzz_dep_.CoverTab[4903]++
//line /usr/local/go/src/math/big/floatconv.go:280
			// _ = "end of CoverTab[4903]"
//line /usr/local/go/src/math/big/floatconv.go:280
		}
//line /usr/local/go/src/math/big/floatconv.go:280
		// _ = "end of CoverTab[4901]"
//line /usr/local/go/src/math/big/floatconv.go:280
	}
//line /usr/local/go/src/math/big/floatconv.go:280
	// _ = "end of CoverTab[4886]"
//line /usr/local/go/src/math/big/floatconv.go:280
	_go_fuzz_dep_.CoverTab[4887]++

							return
//line /usr/local/go/src/math/big/floatconv.go:282
	// _ = "end of CoverTab[4887]"
}

// ParseFloat is like f.Parse(s, base) with f set to the given precision
//line /usr/local/go/src/math/big/floatconv.go:285
// and rounding mode.
//line /usr/local/go/src/math/big/floatconv.go:287
func ParseFloat(s string, base int, prec uint, mode RoundingMode) (f *Float, b int, err error) {
//line /usr/local/go/src/math/big/floatconv.go:287
	_go_fuzz_dep_.CoverTab[4904]++
							return new(Float).SetPrec(prec).SetMode(mode).Parse(s, base)
//line /usr/local/go/src/math/big/floatconv.go:288
	// _ = "end of CoverTab[4904]"
}

var _ fmt.Scanner = (*Float)(nil)	// *Float must implement fmt.Scanner

// Scan is a support routine for fmt.Scanner; it sets z to the value of
//line /usr/local/go/src/math/big/floatconv.go:293
// the scanned number. It accepts formats whose verbs are supported by
//line /usr/local/go/src/math/big/floatconv.go:293
// fmt.Scan for floating point values, which are:
//line /usr/local/go/src/math/big/floatconv.go:293
// 'b' (binary), 'e', 'E', 'f', 'F', 'g' and 'G'.
//line /usr/local/go/src/math/big/floatconv.go:293
// Scan doesn't handle ±Inf.
//line /usr/local/go/src/math/big/floatconv.go:298
func (z *Float) Scan(s fmt.ScanState, ch rune) error {
//line /usr/local/go/src/math/big/floatconv.go:298
	_go_fuzz_dep_.CoverTab[4905]++
							s.SkipSpace()
							_, _, err := z.scan(byteReader{s}, 0)
							return err
//line /usr/local/go/src/math/big/floatconv.go:301
	// _ = "end of CoverTab[4905]"
}

//line /usr/local/go/src/math/big/floatconv.go:302
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/floatconv.go:302
var _ = _go_fuzz_dep_.CoverTab
