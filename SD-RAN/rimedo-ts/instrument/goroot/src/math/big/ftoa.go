// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements Float-to-string conversion functions.
// It is closely following the corresponding implementation
// in strconv/ftoa.go, but modified and simplified for Float.

//line /usr/local/go/src/math/big/ftoa.go:9
package big

//line /usr/local/go/src/math/big/ftoa.go:9
import (
//line /usr/local/go/src/math/big/ftoa.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/ftoa.go:9
)
//line /usr/local/go/src/math/big/ftoa.go:9
import (
//line /usr/local/go/src/math/big/ftoa.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/ftoa.go:9
)

import (
	"bytes"
	"fmt"
	"strconv"
)

// Text converts the floating-point number x to a string according
//line /usr/local/go/src/math/big/ftoa.go:17
// to the given format and precision prec. The format is one of:
//line /usr/local/go/src/math/big/ftoa.go:17
//
//line /usr/local/go/src/math/big/ftoa.go:17
//	'e'	-d.dddde±dd, decimal exponent, at least two (possibly 0) exponent digits
//line /usr/local/go/src/math/big/ftoa.go:17
//	'E'	-d.ddddE±dd, decimal exponent, at least two (possibly 0) exponent digits
//line /usr/local/go/src/math/big/ftoa.go:17
//	'f'	-ddddd.dddd, no exponent
//line /usr/local/go/src/math/big/ftoa.go:17
//	'g'	like 'e' for large exponents, like 'f' otherwise
//line /usr/local/go/src/math/big/ftoa.go:17
//	'G'	like 'E' for large exponents, like 'f' otherwise
//line /usr/local/go/src/math/big/ftoa.go:17
//	'x'	-0xd.dddddp±dd, hexadecimal mantissa, decimal power of two exponent
//line /usr/local/go/src/math/big/ftoa.go:17
//	'p'	-0x.dddp±dd, hexadecimal mantissa, decimal power of two exponent (non-standard)
//line /usr/local/go/src/math/big/ftoa.go:17
//	'b'	-ddddddp±dd, decimal mantissa, decimal power of two exponent (non-standard)
//line /usr/local/go/src/math/big/ftoa.go:17
//
//line /usr/local/go/src/math/big/ftoa.go:17
// For the power-of-two exponent formats, the mantissa is printed in normalized form:
//line /usr/local/go/src/math/big/ftoa.go:17
//
//line /usr/local/go/src/math/big/ftoa.go:17
//	'x'	hexadecimal mantissa in [1, 2), or 0
//line /usr/local/go/src/math/big/ftoa.go:17
//	'p'	hexadecimal mantissa in [½, 1), or 0
//line /usr/local/go/src/math/big/ftoa.go:17
//	'b'	decimal integer mantissa using x.Prec() bits, or 0
//line /usr/local/go/src/math/big/ftoa.go:17
//
//line /usr/local/go/src/math/big/ftoa.go:17
// Note that the 'x' form is the one used by most other languages and libraries.
//line /usr/local/go/src/math/big/ftoa.go:17
//
//line /usr/local/go/src/math/big/ftoa.go:17
// If format is a different character, Text returns a "%" followed by the
//line /usr/local/go/src/math/big/ftoa.go:17
// unrecognized format character.
//line /usr/local/go/src/math/big/ftoa.go:17
//
//line /usr/local/go/src/math/big/ftoa.go:17
// The precision prec controls the number of digits (excluding the exponent)
//line /usr/local/go/src/math/big/ftoa.go:17
// printed by the 'e', 'E', 'f', 'g', 'G', and 'x' formats.
//line /usr/local/go/src/math/big/ftoa.go:17
// For 'e', 'E', 'f', and 'x', it is the number of digits after the decimal point.
//line /usr/local/go/src/math/big/ftoa.go:17
// For 'g' and 'G' it is the total number of digits. A negative precision selects
//line /usr/local/go/src/math/big/ftoa.go:17
// the smallest number of decimal digits necessary to identify the value x uniquely
//line /usr/local/go/src/math/big/ftoa.go:17
// using x.Prec() mantissa bits.
//line /usr/local/go/src/math/big/ftoa.go:17
// The prec value is ignored for the 'b' and 'p' formats.
//line /usr/local/go/src/math/big/ftoa.go:47
func (x *Float) Text(format byte, prec int) string {
//line /usr/local/go/src/math/big/ftoa.go:47
	_go_fuzz_dep_.CoverTab[4949]++
						cap := 10
						if prec > 0 {
//line /usr/local/go/src/math/big/ftoa.go:49
		_go_fuzz_dep_.CoverTab[4951]++
							cap += prec
//line /usr/local/go/src/math/big/ftoa.go:50
		// _ = "end of CoverTab[4951]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:51
		_go_fuzz_dep_.CoverTab[4952]++
//line /usr/local/go/src/math/big/ftoa.go:51
		// _ = "end of CoverTab[4952]"
//line /usr/local/go/src/math/big/ftoa.go:51
	}
//line /usr/local/go/src/math/big/ftoa.go:51
	// _ = "end of CoverTab[4949]"
//line /usr/local/go/src/math/big/ftoa.go:51
	_go_fuzz_dep_.CoverTab[4950]++
						return string(x.Append(make([]byte, 0, cap), format, prec))
//line /usr/local/go/src/math/big/ftoa.go:52
	// _ = "end of CoverTab[4950]"
}

// String formats x like x.Text('g', 10).
//line /usr/local/go/src/math/big/ftoa.go:55
// (String must be called explicitly, Float.Format does not support %s verb.)
//line /usr/local/go/src/math/big/ftoa.go:57
func (x *Float) String() string {
//line /usr/local/go/src/math/big/ftoa.go:57
	_go_fuzz_dep_.CoverTab[4953]++
						return x.Text('g', 10)
//line /usr/local/go/src/math/big/ftoa.go:58
	// _ = "end of CoverTab[4953]"
}

// Append appends to buf the string form of the floating-point number x,
//line /usr/local/go/src/math/big/ftoa.go:61
// as generated by x.Text, and returns the extended buffer.
//line /usr/local/go/src/math/big/ftoa.go:63
func (x *Float) Append(buf []byte, fmt byte, prec int) []byte {
//line /usr/local/go/src/math/big/ftoa.go:63
	_go_fuzz_dep_.CoverTab[4954]++

						if x.neg {
//line /usr/local/go/src/math/big/ftoa.go:65
		_go_fuzz_dep_.CoverTab[4962]++
							buf = append(buf, '-')
//line /usr/local/go/src/math/big/ftoa.go:66
		// _ = "end of CoverTab[4962]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:67
		_go_fuzz_dep_.CoverTab[4963]++
//line /usr/local/go/src/math/big/ftoa.go:67
		// _ = "end of CoverTab[4963]"
//line /usr/local/go/src/math/big/ftoa.go:67
	}
//line /usr/local/go/src/math/big/ftoa.go:67
	// _ = "end of CoverTab[4954]"
//line /usr/local/go/src/math/big/ftoa.go:67
	_go_fuzz_dep_.CoverTab[4955]++

//line /usr/local/go/src/math/big/ftoa.go:70
	if x.form == inf {
//line /usr/local/go/src/math/big/ftoa.go:70
		_go_fuzz_dep_.CoverTab[4964]++
							if !x.neg {
//line /usr/local/go/src/math/big/ftoa.go:71
			_go_fuzz_dep_.CoverTab[4966]++
								buf = append(buf, '+')
//line /usr/local/go/src/math/big/ftoa.go:72
			// _ = "end of CoverTab[4966]"
		} else {
//line /usr/local/go/src/math/big/ftoa.go:73
			_go_fuzz_dep_.CoverTab[4967]++
//line /usr/local/go/src/math/big/ftoa.go:73
			// _ = "end of CoverTab[4967]"
//line /usr/local/go/src/math/big/ftoa.go:73
		}
//line /usr/local/go/src/math/big/ftoa.go:73
		// _ = "end of CoverTab[4964]"
//line /usr/local/go/src/math/big/ftoa.go:73
		_go_fuzz_dep_.CoverTab[4965]++
							return append(buf, "Inf"...)
//line /usr/local/go/src/math/big/ftoa.go:74
		// _ = "end of CoverTab[4965]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:75
		_go_fuzz_dep_.CoverTab[4968]++
//line /usr/local/go/src/math/big/ftoa.go:75
		// _ = "end of CoverTab[4968]"
//line /usr/local/go/src/math/big/ftoa.go:75
	}
//line /usr/local/go/src/math/big/ftoa.go:75
	// _ = "end of CoverTab[4955]"
//line /usr/local/go/src/math/big/ftoa.go:75
	_go_fuzz_dep_.CoverTab[4956]++

//line /usr/local/go/src/math/big/ftoa.go:78
	switch fmt {
	case 'b':
//line /usr/local/go/src/math/big/ftoa.go:79
		_go_fuzz_dep_.CoverTab[4969]++
							return x.fmtB(buf)
//line /usr/local/go/src/math/big/ftoa.go:80
		// _ = "end of CoverTab[4969]"
	case 'p':
//line /usr/local/go/src/math/big/ftoa.go:81
		_go_fuzz_dep_.CoverTab[4970]++
							return x.fmtP(buf)
//line /usr/local/go/src/math/big/ftoa.go:82
		// _ = "end of CoverTab[4970]"
	case 'x':
//line /usr/local/go/src/math/big/ftoa.go:83
		_go_fuzz_dep_.CoverTab[4971]++
							return x.fmtX(buf, prec)
//line /usr/local/go/src/math/big/ftoa.go:84
		// _ = "end of CoverTab[4971]"
//line /usr/local/go/src/math/big/ftoa.go:84
	default:
//line /usr/local/go/src/math/big/ftoa.go:84
		_go_fuzz_dep_.CoverTab[4972]++
//line /usr/local/go/src/math/big/ftoa.go:84
		// _ = "end of CoverTab[4972]"
	}
//line /usr/local/go/src/math/big/ftoa.go:85
	// _ = "end of CoverTab[4956]"
//line /usr/local/go/src/math/big/ftoa.go:85
	_go_fuzz_dep_.CoverTab[4957]++

//line /usr/local/go/src/math/big/ftoa.go:92
	// 1) convert Float to multiprecision decimal
						var d decimal	// == 0.0
						if x.form == finite {
//line /usr/local/go/src/math/big/ftoa.go:94
		_go_fuzz_dep_.CoverTab[4973]++

							d.init(x.mant, int(x.exp)-x.mant.bitLen())
//line /usr/local/go/src/math/big/ftoa.go:96
		// _ = "end of CoverTab[4973]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:97
		_go_fuzz_dep_.CoverTab[4974]++
//line /usr/local/go/src/math/big/ftoa.go:97
		// _ = "end of CoverTab[4974]"
//line /usr/local/go/src/math/big/ftoa.go:97
	}
//line /usr/local/go/src/math/big/ftoa.go:97
	// _ = "end of CoverTab[4957]"
//line /usr/local/go/src/math/big/ftoa.go:97
	_go_fuzz_dep_.CoverTab[4958]++

//line /usr/local/go/src/math/big/ftoa.go:100
	shortest := false
	if prec < 0 {
//line /usr/local/go/src/math/big/ftoa.go:101
		_go_fuzz_dep_.CoverTab[4975]++
							shortest = true
							roundShortest(&d, x)

							switch fmt {
		case 'e', 'E':
//line /usr/local/go/src/math/big/ftoa.go:106
			_go_fuzz_dep_.CoverTab[4976]++
								prec = len(d.mant) - 1
//line /usr/local/go/src/math/big/ftoa.go:107
			// _ = "end of CoverTab[4976]"
		case 'f':
//line /usr/local/go/src/math/big/ftoa.go:108
			_go_fuzz_dep_.CoverTab[4977]++
								prec = max(len(d.mant)-d.exp, 0)
//line /usr/local/go/src/math/big/ftoa.go:109
			// _ = "end of CoverTab[4977]"
		case 'g', 'G':
//line /usr/local/go/src/math/big/ftoa.go:110
			_go_fuzz_dep_.CoverTab[4978]++
								prec = len(d.mant)
//line /usr/local/go/src/math/big/ftoa.go:111
			// _ = "end of CoverTab[4978]"
//line /usr/local/go/src/math/big/ftoa.go:111
		default:
//line /usr/local/go/src/math/big/ftoa.go:111
			_go_fuzz_dep_.CoverTab[4979]++
//line /usr/local/go/src/math/big/ftoa.go:111
			// _ = "end of CoverTab[4979]"
		}
//line /usr/local/go/src/math/big/ftoa.go:112
		// _ = "end of CoverTab[4975]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:113
		_go_fuzz_dep_.CoverTab[4980]++

							switch fmt {
		case 'e', 'E':
//line /usr/local/go/src/math/big/ftoa.go:116
			_go_fuzz_dep_.CoverTab[4981]++

								d.round(1 + prec)
//line /usr/local/go/src/math/big/ftoa.go:118
			// _ = "end of CoverTab[4981]"
		case 'f':
//line /usr/local/go/src/math/big/ftoa.go:119
			_go_fuzz_dep_.CoverTab[4982]++

								d.round(d.exp + prec)
//line /usr/local/go/src/math/big/ftoa.go:121
			// _ = "end of CoverTab[4982]"
		case 'g', 'G':
//line /usr/local/go/src/math/big/ftoa.go:122
			_go_fuzz_dep_.CoverTab[4983]++
								if prec == 0 {
//line /usr/local/go/src/math/big/ftoa.go:123
				_go_fuzz_dep_.CoverTab[4986]++
									prec = 1
//line /usr/local/go/src/math/big/ftoa.go:124
				// _ = "end of CoverTab[4986]"
			} else {
//line /usr/local/go/src/math/big/ftoa.go:125
				_go_fuzz_dep_.CoverTab[4987]++
//line /usr/local/go/src/math/big/ftoa.go:125
				// _ = "end of CoverTab[4987]"
//line /usr/local/go/src/math/big/ftoa.go:125
			}
//line /usr/local/go/src/math/big/ftoa.go:125
			// _ = "end of CoverTab[4983]"
//line /usr/local/go/src/math/big/ftoa.go:125
			_go_fuzz_dep_.CoverTab[4984]++
								d.round(prec)
//line /usr/local/go/src/math/big/ftoa.go:126
			// _ = "end of CoverTab[4984]"
//line /usr/local/go/src/math/big/ftoa.go:126
		default:
//line /usr/local/go/src/math/big/ftoa.go:126
			_go_fuzz_dep_.CoverTab[4985]++
//line /usr/local/go/src/math/big/ftoa.go:126
			// _ = "end of CoverTab[4985]"
		}
//line /usr/local/go/src/math/big/ftoa.go:127
		// _ = "end of CoverTab[4980]"
	}
//line /usr/local/go/src/math/big/ftoa.go:128
	// _ = "end of CoverTab[4958]"
//line /usr/local/go/src/math/big/ftoa.go:128
	_go_fuzz_dep_.CoverTab[4959]++

//line /usr/local/go/src/math/big/ftoa.go:131
	switch fmt {
	case 'e', 'E':
//line /usr/local/go/src/math/big/ftoa.go:132
		_go_fuzz_dep_.CoverTab[4988]++
							return fmtE(buf, fmt, prec, d)
//line /usr/local/go/src/math/big/ftoa.go:133
		// _ = "end of CoverTab[4988]"
	case 'f':
//line /usr/local/go/src/math/big/ftoa.go:134
		_go_fuzz_dep_.CoverTab[4989]++
							return fmtF(buf, prec, d)
//line /usr/local/go/src/math/big/ftoa.go:135
		// _ = "end of CoverTab[4989]"
	case 'g', 'G':
//line /usr/local/go/src/math/big/ftoa.go:136
		_go_fuzz_dep_.CoverTab[4990]++

							eprec := prec
							if eprec > len(d.mant) && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:139
			_go_fuzz_dep_.CoverTab[4996]++
//line /usr/local/go/src/math/big/ftoa.go:139
			return len(d.mant) >= d.exp
//line /usr/local/go/src/math/big/ftoa.go:139
			// _ = "end of CoverTab[4996]"
//line /usr/local/go/src/math/big/ftoa.go:139
		}() {
//line /usr/local/go/src/math/big/ftoa.go:139
			_go_fuzz_dep_.CoverTab[4997]++
								eprec = len(d.mant)
//line /usr/local/go/src/math/big/ftoa.go:140
			// _ = "end of CoverTab[4997]"
		} else {
//line /usr/local/go/src/math/big/ftoa.go:141
			_go_fuzz_dep_.CoverTab[4998]++
//line /usr/local/go/src/math/big/ftoa.go:141
			// _ = "end of CoverTab[4998]"
//line /usr/local/go/src/math/big/ftoa.go:141
		}
//line /usr/local/go/src/math/big/ftoa.go:141
		// _ = "end of CoverTab[4990]"
//line /usr/local/go/src/math/big/ftoa.go:141
		_go_fuzz_dep_.CoverTab[4991]++

//line /usr/local/go/src/math/big/ftoa.go:146
		if shortest {
//line /usr/local/go/src/math/big/ftoa.go:146
			_go_fuzz_dep_.CoverTab[4999]++
								eprec = 6
//line /usr/local/go/src/math/big/ftoa.go:147
			// _ = "end of CoverTab[4999]"
		} else {
//line /usr/local/go/src/math/big/ftoa.go:148
			_go_fuzz_dep_.CoverTab[5000]++
//line /usr/local/go/src/math/big/ftoa.go:148
			// _ = "end of CoverTab[5000]"
//line /usr/local/go/src/math/big/ftoa.go:148
		}
//line /usr/local/go/src/math/big/ftoa.go:148
		// _ = "end of CoverTab[4991]"
//line /usr/local/go/src/math/big/ftoa.go:148
		_go_fuzz_dep_.CoverTab[4992]++
							exp := d.exp - 1
							if exp < -4 || func() bool {
//line /usr/local/go/src/math/big/ftoa.go:150
			_go_fuzz_dep_.CoverTab[5001]++
//line /usr/local/go/src/math/big/ftoa.go:150
			return exp >= eprec
//line /usr/local/go/src/math/big/ftoa.go:150
			// _ = "end of CoverTab[5001]"
//line /usr/local/go/src/math/big/ftoa.go:150
		}() {
//line /usr/local/go/src/math/big/ftoa.go:150
			_go_fuzz_dep_.CoverTab[5002]++
								if prec > len(d.mant) {
//line /usr/local/go/src/math/big/ftoa.go:151
				_go_fuzz_dep_.CoverTab[5004]++
									prec = len(d.mant)
//line /usr/local/go/src/math/big/ftoa.go:152
				// _ = "end of CoverTab[5004]"
			} else {
//line /usr/local/go/src/math/big/ftoa.go:153
				_go_fuzz_dep_.CoverTab[5005]++
//line /usr/local/go/src/math/big/ftoa.go:153
				// _ = "end of CoverTab[5005]"
//line /usr/local/go/src/math/big/ftoa.go:153
			}
//line /usr/local/go/src/math/big/ftoa.go:153
			// _ = "end of CoverTab[5002]"
//line /usr/local/go/src/math/big/ftoa.go:153
			_go_fuzz_dep_.CoverTab[5003]++
								return fmtE(buf, fmt+'e'-'g', prec-1, d)
//line /usr/local/go/src/math/big/ftoa.go:154
			// _ = "end of CoverTab[5003]"
		} else {
//line /usr/local/go/src/math/big/ftoa.go:155
			_go_fuzz_dep_.CoverTab[5006]++
//line /usr/local/go/src/math/big/ftoa.go:155
			// _ = "end of CoverTab[5006]"
//line /usr/local/go/src/math/big/ftoa.go:155
		}
//line /usr/local/go/src/math/big/ftoa.go:155
		// _ = "end of CoverTab[4992]"
//line /usr/local/go/src/math/big/ftoa.go:155
		_go_fuzz_dep_.CoverTab[4993]++
							if prec > d.exp {
//line /usr/local/go/src/math/big/ftoa.go:156
			_go_fuzz_dep_.CoverTab[5007]++
								prec = len(d.mant)
//line /usr/local/go/src/math/big/ftoa.go:157
			// _ = "end of CoverTab[5007]"
		} else {
//line /usr/local/go/src/math/big/ftoa.go:158
			_go_fuzz_dep_.CoverTab[5008]++
//line /usr/local/go/src/math/big/ftoa.go:158
			// _ = "end of CoverTab[5008]"
//line /usr/local/go/src/math/big/ftoa.go:158
		}
//line /usr/local/go/src/math/big/ftoa.go:158
		// _ = "end of CoverTab[4993]"
//line /usr/local/go/src/math/big/ftoa.go:158
		_go_fuzz_dep_.CoverTab[4994]++
							return fmtF(buf, max(prec-d.exp, 0), d)
//line /usr/local/go/src/math/big/ftoa.go:159
		// _ = "end of CoverTab[4994]"
//line /usr/local/go/src/math/big/ftoa.go:159
	default:
//line /usr/local/go/src/math/big/ftoa.go:159
		_go_fuzz_dep_.CoverTab[4995]++
//line /usr/local/go/src/math/big/ftoa.go:159
		// _ = "end of CoverTab[4995]"
	}
//line /usr/local/go/src/math/big/ftoa.go:160
	// _ = "end of CoverTab[4959]"
//line /usr/local/go/src/math/big/ftoa.go:160
	_go_fuzz_dep_.CoverTab[4960]++

//line /usr/local/go/src/math/big/ftoa.go:163
	if x.neg {
//line /usr/local/go/src/math/big/ftoa.go:163
		_go_fuzz_dep_.CoverTab[5009]++
							buf = buf[:len(buf)-1]
//line /usr/local/go/src/math/big/ftoa.go:164
		// _ = "end of CoverTab[5009]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:165
		_go_fuzz_dep_.CoverTab[5010]++
//line /usr/local/go/src/math/big/ftoa.go:165
		// _ = "end of CoverTab[5010]"
//line /usr/local/go/src/math/big/ftoa.go:165
	}
//line /usr/local/go/src/math/big/ftoa.go:165
	// _ = "end of CoverTab[4960]"
//line /usr/local/go/src/math/big/ftoa.go:165
	_go_fuzz_dep_.CoverTab[4961]++
						return append(buf, '%', fmt)
//line /usr/local/go/src/math/big/ftoa.go:166
	// _ = "end of CoverTab[4961]"
}

func roundShortest(d *decimal, x *Float) {
//line /usr/local/go/src/math/big/ftoa.go:169
	_go_fuzz_dep_.CoverTab[5011]++

						if len(d.mant) == 0 {
//line /usr/local/go/src/math/big/ftoa.go:171
		_go_fuzz_dep_.CoverTab[5014]++
							return
//line /usr/local/go/src/math/big/ftoa.go:172
		// _ = "end of CoverTab[5014]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:173
		_go_fuzz_dep_.CoverTab[5015]++
//line /usr/local/go/src/math/big/ftoa.go:173
		// _ = "end of CoverTab[5015]"
//line /usr/local/go/src/math/big/ftoa.go:173
	}
//line /usr/local/go/src/math/big/ftoa.go:173
	// _ = "end of CoverTab[5011]"
//line /usr/local/go/src/math/big/ftoa.go:173
	_go_fuzz_dep_.CoverTab[5012]++

//line /usr/local/go/src/math/big/ftoa.go:186
	mant := nat(nil).set(x.mant)
	exp := int(x.exp) - mant.bitLen()
	s := mant.bitLen() - int(x.prec+1)
	switch {
	case s < 0:
//line /usr/local/go/src/math/big/ftoa.go:190
		_go_fuzz_dep_.CoverTab[5016]++
							mant = mant.shl(mant, uint(-s))
//line /usr/local/go/src/math/big/ftoa.go:191
		// _ = "end of CoverTab[5016]"
	case s > 0:
//line /usr/local/go/src/math/big/ftoa.go:192
		_go_fuzz_dep_.CoverTab[5017]++
							mant = mant.shr(mant, uint(+s))
//line /usr/local/go/src/math/big/ftoa.go:193
		// _ = "end of CoverTab[5017]"
//line /usr/local/go/src/math/big/ftoa.go:193
	default:
//line /usr/local/go/src/math/big/ftoa.go:193
		_go_fuzz_dep_.CoverTab[5018]++
//line /usr/local/go/src/math/big/ftoa.go:193
		// _ = "end of CoverTab[5018]"
	}
//line /usr/local/go/src/math/big/ftoa.go:194
	// _ = "end of CoverTab[5012]"
//line /usr/local/go/src/math/big/ftoa.go:194
	_go_fuzz_dep_.CoverTab[5013]++
						exp += s

//line /usr/local/go/src/math/big/ftoa.go:198
	// 2) Compute lower bound by subtracting 1/2 ulp.
						var lower decimal
						var tmp nat
						lower.init(tmp.sub(mant, natOne), exp)

						// 3) Compute upper bound by adding 1/2 ulp.
						var upper decimal
						upper.init(tmp.add(mant, natOne), exp)

//line /usr/local/go/src/math/big/ftoa.go:210
	inclusive := mant[0]&2 == 0

//line /usr/local/go/src/math/big/ftoa.go:214
	for i, m := range d.mant {
//line /usr/local/go/src/math/big/ftoa.go:214
		_go_fuzz_dep_.CoverTab[5019]++
							l := lower.at(i)
							u := upper.at(i)

//line /usr/local/go/src/math/big/ftoa.go:221
		okdown := l != m || func() bool {
//line /usr/local/go/src/math/big/ftoa.go:221
			_go_fuzz_dep_.CoverTab[5020]++
//line /usr/local/go/src/math/big/ftoa.go:221
			return inclusive && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:221
				_go_fuzz_dep_.CoverTab[5021]++
//line /usr/local/go/src/math/big/ftoa.go:221
				return i+1 == len(lower.mant)
//line /usr/local/go/src/math/big/ftoa.go:221
				// _ = "end of CoverTab[5021]"
//line /usr/local/go/src/math/big/ftoa.go:221
			}()
//line /usr/local/go/src/math/big/ftoa.go:221
			// _ = "end of CoverTab[5020]"
//line /usr/local/go/src/math/big/ftoa.go:221
		}()

//line /usr/local/go/src/math/big/ftoa.go:225
		okup := m != u && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:225
			_go_fuzz_dep_.CoverTab[5022]++
//line /usr/local/go/src/math/big/ftoa.go:225
			return (inclusive || func() bool {
//line /usr/local/go/src/math/big/ftoa.go:225
				_go_fuzz_dep_.CoverTab[5023]++
//line /usr/local/go/src/math/big/ftoa.go:225
				return m+1 < u
//line /usr/local/go/src/math/big/ftoa.go:225
				// _ = "end of CoverTab[5023]"
//line /usr/local/go/src/math/big/ftoa.go:225
			}() || func() bool {
//line /usr/local/go/src/math/big/ftoa.go:225
				_go_fuzz_dep_.CoverTab[5024]++
//line /usr/local/go/src/math/big/ftoa.go:225
				return i+1 < len(upper.mant)
//line /usr/local/go/src/math/big/ftoa.go:225
				// _ = "end of CoverTab[5024]"
//line /usr/local/go/src/math/big/ftoa.go:225
			}())
//line /usr/local/go/src/math/big/ftoa.go:225
			// _ = "end of CoverTab[5022]"
//line /usr/local/go/src/math/big/ftoa.go:225
		}()

//line /usr/local/go/src/math/big/ftoa.go:229
		switch {
		case okdown && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:230
			_go_fuzz_dep_.CoverTab[5029]++
//line /usr/local/go/src/math/big/ftoa.go:230
			return okup
//line /usr/local/go/src/math/big/ftoa.go:230
			// _ = "end of CoverTab[5029]"
//line /usr/local/go/src/math/big/ftoa.go:230
		}():
//line /usr/local/go/src/math/big/ftoa.go:230
			_go_fuzz_dep_.CoverTab[5025]++
								d.round(i + 1)
								return
//line /usr/local/go/src/math/big/ftoa.go:232
			// _ = "end of CoverTab[5025]"
		case okdown:
//line /usr/local/go/src/math/big/ftoa.go:233
			_go_fuzz_dep_.CoverTab[5026]++
								d.roundDown(i + 1)
								return
//line /usr/local/go/src/math/big/ftoa.go:235
			// _ = "end of CoverTab[5026]"
		case okup:
//line /usr/local/go/src/math/big/ftoa.go:236
			_go_fuzz_dep_.CoverTab[5027]++
								d.roundUp(i + 1)
								return
//line /usr/local/go/src/math/big/ftoa.go:238
			// _ = "end of CoverTab[5027]"
//line /usr/local/go/src/math/big/ftoa.go:238
		default:
//line /usr/local/go/src/math/big/ftoa.go:238
			_go_fuzz_dep_.CoverTab[5028]++
//line /usr/local/go/src/math/big/ftoa.go:238
			// _ = "end of CoverTab[5028]"
		}
//line /usr/local/go/src/math/big/ftoa.go:239
		// _ = "end of CoverTab[5019]"
	}
//line /usr/local/go/src/math/big/ftoa.go:240
	// _ = "end of CoverTab[5013]"
}

// %e: d.ddddde±dd
func fmtE(buf []byte, fmt byte, prec int, d decimal) []byte {
//line /usr/local/go/src/math/big/ftoa.go:244
	_go_fuzz_dep_.CoverTab[5030]++

						ch := byte('0')
						if len(d.mant) > 0 {
//line /usr/local/go/src/math/big/ftoa.go:247
		_go_fuzz_dep_.CoverTab[5036]++
							ch = d.mant[0]
//line /usr/local/go/src/math/big/ftoa.go:248
		// _ = "end of CoverTab[5036]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:249
		_go_fuzz_dep_.CoverTab[5037]++
//line /usr/local/go/src/math/big/ftoa.go:249
		// _ = "end of CoverTab[5037]"
//line /usr/local/go/src/math/big/ftoa.go:249
	}
//line /usr/local/go/src/math/big/ftoa.go:249
	// _ = "end of CoverTab[5030]"
//line /usr/local/go/src/math/big/ftoa.go:249
	_go_fuzz_dep_.CoverTab[5031]++
						buf = append(buf, ch)

//line /usr/local/go/src/math/big/ftoa.go:253
	if prec > 0 {
//line /usr/local/go/src/math/big/ftoa.go:253
		_go_fuzz_dep_.CoverTab[5038]++
							buf = append(buf, '.')
							i := 1
							m := min(len(d.mant), prec+1)
							if i < m {
//line /usr/local/go/src/math/big/ftoa.go:257
			_go_fuzz_dep_.CoverTab[5040]++
								buf = append(buf, d.mant[i:m]...)
								i = m
//line /usr/local/go/src/math/big/ftoa.go:259
			// _ = "end of CoverTab[5040]"
		} else {
//line /usr/local/go/src/math/big/ftoa.go:260
			_go_fuzz_dep_.CoverTab[5041]++
//line /usr/local/go/src/math/big/ftoa.go:260
			// _ = "end of CoverTab[5041]"
//line /usr/local/go/src/math/big/ftoa.go:260
		}
//line /usr/local/go/src/math/big/ftoa.go:260
		// _ = "end of CoverTab[5038]"
//line /usr/local/go/src/math/big/ftoa.go:260
		_go_fuzz_dep_.CoverTab[5039]++
							for ; i <= prec; i++ {
//line /usr/local/go/src/math/big/ftoa.go:261
			_go_fuzz_dep_.CoverTab[5042]++
								buf = append(buf, '0')
//line /usr/local/go/src/math/big/ftoa.go:262
			// _ = "end of CoverTab[5042]"
		}
//line /usr/local/go/src/math/big/ftoa.go:263
		// _ = "end of CoverTab[5039]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:264
		_go_fuzz_dep_.CoverTab[5043]++
//line /usr/local/go/src/math/big/ftoa.go:264
		// _ = "end of CoverTab[5043]"
//line /usr/local/go/src/math/big/ftoa.go:264
	}
//line /usr/local/go/src/math/big/ftoa.go:264
	// _ = "end of CoverTab[5031]"
//line /usr/local/go/src/math/big/ftoa.go:264
	_go_fuzz_dep_.CoverTab[5032]++

//line /usr/local/go/src/math/big/ftoa.go:267
	buf = append(buf, fmt)
	var exp int64
	if len(d.mant) > 0 {
//line /usr/local/go/src/math/big/ftoa.go:269
		_go_fuzz_dep_.CoverTab[5044]++
							exp = int64(d.exp) - 1
//line /usr/local/go/src/math/big/ftoa.go:270
		// _ = "end of CoverTab[5044]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:271
		_go_fuzz_dep_.CoverTab[5045]++
//line /usr/local/go/src/math/big/ftoa.go:271
		// _ = "end of CoverTab[5045]"
//line /usr/local/go/src/math/big/ftoa.go:271
	}
//line /usr/local/go/src/math/big/ftoa.go:271
	// _ = "end of CoverTab[5032]"
//line /usr/local/go/src/math/big/ftoa.go:271
	_go_fuzz_dep_.CoverTab[5033]++
						if exp < 0 {
//line /usr/local/go/src/math/big/ftoa.go:272
		_go_fuzz_dep_.CoverTab[5046]++
							ch = '-'
							exp = -exp
//line /usr/local/go/src/math/big/ftoa.go:274
		// _ = "end of CoverTab[5046]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:275
		_go_fuzz_dep_.CoverTab[5047]++
							ch = '+'
//line /usr/local/go/src/math/big/ftoa.go:276
		// _ = "end of CoverTab[5047]"
	}
//line /usr/local/go/src/math/big/ftoa.go:277
	// _ = "end of CoverTab[5033]"
//line /usr/local/go/src/math/big/ftoa.go:277
	_go_fuzz_dep_.CoverTab[5034]++
						buf = append(buf, ch)

//line /usr/local/go/src/math/big/ftoa.go:281
	if exp < 10 {
//line /usr/local/go/src/math/big/ftoa.go:281
		_go_fuzz_dep_.CoverTab[5048]++
							buf = append(buf, '0')
//line /usr/local/go/src/math/big/ftoa.go:282
		// _ = "end of CoverTab[5048]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:283
		_go_fuzz_dep_.CoverTab[5049]++
//line /usr/local/go/src/math/big/ftoa.go:283
		// _ = "end of CoverTab[5049]"
//line /usr/local/go/src/math/big/ftoa.go:283
	}
//line /usr/local/go/src/math/big/ftoa.go:283
	// _ = "end of CoverTab[5034]"
//line /usr/local/go/src/math/big/ftoa.go:283
	_go_fuzz_dep_.CoverTab[5035]++
						return strconv.AppendInt(buf, exp, 10)
//line /usr/local/go/src/math/big/ftoa.go:284
	// _ = "end of CoverTab[5035]"
}

// %f: ddddddd.ddddd
func fmtF(buf []byte, prec int, d decimal) []byte {
//line /usr/local/go/src/math/big/ftoa.go:288
	_go_fuzz_dep_.CoverTab[5050]++

						if d.exp > 0 {
//line /usr/local/go/src/math/big/ftoa.go:290
		_go_fuzz_dep_.CoverTab[5053]++
							m := min(len(d.mant), d.exp)
							buf = append(buf, d.mant[:m]...)
							for ; m < d.exp; m++ {
//line /usr/local/go/src/math/big/ftoa.go:293
			_go_fuzz_dep_.CoverTab[5054]++
								buf = append(buf, '0')
//line /usr/local/go/src/math/big/ftoa.go:294
			// _ = "end of CoverTab[5054]"
		}
//line /usr/local/go/src/math/big/ftoa.go:295
		// _ = "end of CoverTab[5053]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:296
		_go_fuzz_dep_.CoverTab[5055]++
							buf = append(buf, '0')
//line /usr/local/go/src/math/big/ftoa.go:297
		// _ = "end of CoverTab[5055]"
	}
//line /usr/local/go/src/math/big/ftoa.go:298
	// _ = "end of CoverTab[5050]"
//line /usr/local/go/src/math/big/ftoa.go:298
	_go_fuzz_dep_.CoverTab[5051]++

//line /usr/local/go/src/math/big/ftoa.go:301
	if prec > 0 {
//line /usr/local/go/src/math/big/ftoa.go:301
		_go_fuzz_dep_.CoverTab[5056]++
							buf = append(buf, '.')
							for i := 0; i < prec; i++ {
//line /usr/local/go/src/math/big/ftoa.go:303
			_go_fuzz_dep_.CoverTab[5057]++
								buf = append(buf, d.at(d.exp+i))
//line /usr/local/go/src/math/big/ftoa.go:304
			// _ = "end of CoverTab[5057]"
		}
//line /usr/local/go/src/math/big/ftoa.go:305
		// _ = "end of CoverTab[5056]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:306
		_go_fuzz_dep_.CoverTab[5058]++
//line /usr/local/go/src/math/big/ftoa.go:306
		// _ = "end of CoverTab[5058]"
//line /usr/local/go/src/math/big/ftoa.go:306
	}
//line /usr/local/go/src/math/big/ftoa.go:306
	// _ = "end of CoverTab[5051]"
//line /usr/local/go/src/math/big/ftoa.go:306
	_go_fuzz_dep_.CoverTab[5052]++

						return buf
//line /usr/local/go/src/math/big/ftoa.go:308
	// _ = "end of CoverTab[5052]"
}

// fmtB appends the string of x in the format mantissa "p" exponent
//line /usr/local/go/src/math/big/ftoa.go:311
// with a decimal mantissa and a binary exponent, or 0" if x is zero,
//line /usr/local/go/src/math/big/ftoa.go:311
// and returns the extended buffer.
//line /usr/local/go/src/math/big/ftoa.go:311
// The mantissa is normalized such that is uses x.Prec() bits in binary
//line /usr/local/go/src/math/big/ftoa.go:311
// representation.
//line /usr/local/go/src/math/big/ftoa.go:311
// The sign of x is ignored, and x must not be an Inf.
//line /usr/local/go/src/math/big/ftoa.go:311
// (The caller handles Inf before invoking fmtB.)
//line /usr/local/go/src/math/big/ftoa.go:318
func (x *Float) fmtB(buf []byte) []byte {
//line /usr/local/go/src/math/big/ftoa.go:318
	_go_fuzz_dep_.CoverTab[5059]++
						if x.form == zero {
//line /usr/local/go/src/math/big/ftoa.go:319
		_go_fuzz_dep_.CoverTab[5064]++
							return append(buf, '0')
//line /usr/local/go/src/math/big/ftoa.go:320
		// _ = "end of CoverTab[5064]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:321
		_go_fuzz_dep_.CoverTab[5065]++
//line /usr/local/go/src/math/big/ftoa.go:321
		// _ = "end of CoverTab[5065]"
//line /usr/local/go/src/math/big/ftoa.go:321
	}
//line /usr/local/go/src/math/big/ftoa.go:321
	// _ = "end of CoverTab[5059]"
//line /usr/local/go/src/math/big/ftoa.go:321
	_go_fuzz_dep_.CoverTab[5060]++

						if debugFloat && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:323
		_go_fuzz_dep_.CoverTab[5066]++
//line /usr/local/go/src/math/big/ftoa.go:323
		return x.form != finite
//line /usr/local/go/src/math/big/ftoa.go:323
		// _ = "end of CoverTab[5066]"
//line /usr/local/go/src/math/big/ftoa.go:323
	}() {
//line /usr/local/go/src/math/big/ftoa.go:323
		_go_fuzz_dep_.CoverTab[5067]++
							panic("non-finite float")
//line /usr/local/go/src/math/big/ftoa.go:324
		// _ = "end of CoverTab[5067]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:325
		_go_fuzz_dep_.CoverTab[5068]++
//line /usr/local/go/src/math/big/ftoa.go:325
		// _ = "end of CoverTab[5068]"
//line /usr/local/go/src/math/big/ftoa.go:325
	}
//line /usr/local/go/src/math/big/ftoa.go:325
	// _ = "end of CoverTab[5060]"
//line /usr/local/go/src/math/big/ftoa.go:325
	_go_fuzz_dep_.CoverTab[5061]++

//line /usr/local/go/src/math/big/ftoa.go:329
	m := x.mant
	switch w := uint32(len(x.mant)) * _W; {
	case w < x.prec:
//line /usr/local/go/src/math/big/ftoa.go:331
		_go_fuzz_dep_.CoverTab[5069]++
							m = nat(nil).shl(m, uint(x.prec-w))
//line /usr/local/go/src/math/big/ftoa.go:332
		// _ = "end of CoverTab[5069]"
	case w > x.prec:
//line /usr/local/go/src/math/big/ftoa.go:333
		_go_fuzz_dep_.CoverTab[5070]++
							m = nat(nil).shr(m, uint(w-x.prec))
//line /usr/local/go/src/math/big/ftoa.go:334
		// _ = "end of CoverTab[5070]"
//line /usr/local/go/src/math/big/ftoa.go:334
	default:
//line /usr/local/go/src/math/big/ftoa.go:334
		_go_fuzz_dep_.CoverTab[5071]++
//line /usr/local/go/src/math/big/ftoa.go:334
		// _ = "end of CoverTab[5071]"
	}
//line /usr/local/go/src/math/big/ftoa.go:335
	// _ = "end of CoverTab[5061]"
//line /usr/local/go/src/math/big/ftoa.go:335
	_go_fuzz_dep_.CoverTab[5062]++

						buf = append(buf, m.utoa(10)...)
						buf = append(buf, 'p')
						e := int64(x.exp) - int64(x.prec)
						if e >= 0 {
//line /usr/local/go/src/math/big/ftoa.go:340
		_go_fuzz_dep_.CoverTab[5072]++
							buf = append(buf, '+')
//line /usr/local/go/src/math/big/ftoa.go:341
		// _ = "end of CoverTab[5072]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:342
		_go_fuzz_dep_.CoverTab[5073]++
//line /usr/local/go/src/math/big/ftoa.go:342
		// _ = "end of CoverTab[5073]"
//line /usr/local/go/src/math/big/ftoa.go:342
	}
//line /usr/local/go/src/math/big/ftoa.go:342
	// _ = "end of CoverTab[5062]"
//line /usr/local/go/src/math/big/ftoa.go:342
	_go_fuzz_dep_.CoverTab[5063]++
						return strconv.AppendInt(buf, e, 10)
//line /usr/local/go/src/math/big/ftoa.go:343
	// _ = "end of CoverTab[5063]"
}

// fmtX appends the string of x in the format "0x1." mantissa "p" exponent
//line /usr/local/go/src/math/big/ftoa.go:346
// with a hexadecimal mantissa and a binary exponent, or "0x0p0" if x is zero,
//line /usr/local/go/src/math/big/ftoa.go:346
// and returns the extended buffer.
//line /usr/local/go/src/math/big/ftoa.go:346
// A non-zero mantissa is normalized such that 1.0 <= mantissa < 2.0.
//line /usr/local/go/src/math/big/ftoa.go:346
// The sign of x is ignored, and x must not be an Inf.
//line /usr/local/go/src/math/big/ftoa.go:346
// (The caller handles Inf before invoking fmtX.)
//line /usr/local/go/src/math/big/ftoa.go:352
func (x *Float) fmtX(buf []byte, prec int) []byte {
//line /usr/local/go/src/math/big/ftoa.go:352
	_go_fuzz_dep_.CoverTab[5074]++
						if x.form == zero {
//line /usr/local/go/src/math/big/ftoa.go:353
		_go_fuzz_dep_.CoverTab[5083]++
							buf = append(buf, "0x0"...)
							if prec > 0 {
//line /usr/local/go/src/math/big/ftoa.go:355
			_go_fuzz_dep_.CoverTab[5085]++
								buf = append(buf, '.')
								for i := 0; i < prec; i++ {
//line /usr/local/go/src/math/big/ftoa.go:357
				_go_fuzz_dep_.CoverTab[5086]++
									buf = append(buf, '0')
//line /usr/local/go/src/math/big/ftoa.go:358
				// _ = "end of CoverTab[5086]"
			}
//line /usr/local/go/src/math/big/ftoa.go:359
			// _ = "end of CoverTab[5085]"
		} else {
//line /usr/local/go/src/math/big/ftoa.go:360
			_go_fuzz_dep_.CoverTab[5087]++
//line /usr/local/go/src/math/big/ftoa.go:360
			// _ = "end of CoverTab[5087]"
//line /usr/local/go/src/math/big/ftoa.go:360
		}
//line /usr/local/go/src/math/big/ftoa.go:360
		// _ = "end of CoverTab[5083]"
//line /usr/local/go/src/math/big/ftoa.go:360
		_go_fuzz_dep_.CoverTab[5084]++
							buf = append(buf, "p+00"...)
							return buf
//line /usr/local/go/src/math/big/ftoa.go:362
		// _ = "end of CoverTab[5084]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:363
		_go_fuzz_dep_.CoverTab[5088]++
//line /usr/local/go/src/math/big/ftoa.go:363
		// _ = "end of CoverTab[5088]"
//line /usr/local/go/src/math/big/ftoa.go:363
	}
//line /usr/local/go/src/math/big/ftoa.go:363
	// _ = "end of CoverTab[5074]"
//line /usr/local/go/src/math/big/ftoa.go:363
	_go_fuzz_dep_.CoverTab[5075]++

						if debugFloat && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:365
		_go_fuzz_dep_.CoverTab[5089]++
//line /usr/local/go/src/math/big/ftoa.go:365
		return x.form != finite
//line /usr/local/go/src/math/big/ftoa.go:365
		// _ = "end of CoverTab[5089]"
//line /usr/local/go/src/math/big/ftoa.go:365
	}() {
//line /usr/local/go/src/math/big/ftoa.go:365
		_go_fuzz_dep_.CoverTab[5090]++
							panic("non-finite float")
//line /usr/local/go/src/math/big/ftoa.go:366
		// _ = "end of CoverTab[5090]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:367
		_go_fuzz_dep_.CoverTab[5091]++
//line /usr/local/go/src/math/big/ftoa.go:367
		// _ = "end of CoverTab[5091]"
//line /usr/local/go/src/math/big/ftoa.go:367
	}
//line /usr/local/go/src/math/big/ftoa.go:367
	// _ = "end of CoverTab[5075]"
//line /usr/local/go/src/math/big/ftoa.go:367
	_go_fuzz_dep_.CoverTab[5076]++

	// round mantissa to n bits
	var n uint
	if prec < 0 {
//line /usr/local/go/src/math/big/ftoa.go:371
		_go_fuzz_dep_.CoverTab[5092]++
							n = 1 + (x.MinPrec()-1+3)/4*4
//line /usr/local/go/src/math/big/ftoa.go:372
		// _ = "end of CoverTab[5092]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:373
		_go_fuzz_dep_.CoverTab[5093]++
							n = 1 + 4*uint(prec)
//line /usr/local/go/src/math/big/ftoa.go:374
		// _ = "end of CoverTab[5093]"
	}
//line /usr/local/go/src/math/big/ftoa.go:375
	// _ = "end of CoverTab[5076]"
//line /usr/local/go/src/math/big/ftoa.go:375
	_go_fuzz_dep_.CoverTab[5077]++

						x = new(Float).SetPrec(n).SetMode(x.mode).Set(x)

//line /usr/local/go/src/math/big/ftoa.go:380
	m := x.mant
	switch w := uint(len(x.mant)) * _W; {
	case w < n:
//line /usr/local/go/src/math/big/ftoa.go:382
		_go_fuzz_dep_.CoverTab[5094]++
							m = nat(nil).shl(m, n-w)
//line /usr/local/go/src/math/big/ftoa.go:383
		// _ = "end of CoverTab[5094]"
	case w > n:
//line /usr/local/go/src/math/big/ftoa.go:384
		_go_fuzz_dep_.CoverTab[5095]++
							m = nat(nil).shr(m, w-n)
//line /usr/local/go/src/math/big/ftoa.go:385
		// _ = "end of CoverTab[5095]"
//line /usr/local/go/src/math/big/ftoa.go:385
	default:
//line /usr/local/go/src/math/big/ftoa.go:385
		_go_fuzz_dep_.CoverTab[5096]++
//line /usr/local/go/src/math/big/ftoa.go:385
		// _ = "end of CoverTab[5096]"
	}
//line /usr/local/go/src/math/big/ftoa.go:386
	// _ = "end of CoverTab[5077]"
//line /usr/local/go/src/math/big/ftoa.go:386
	_go_fuzz_dep_.CoverTab[5078]++
						exp64 := int64(x.exp) - 1

						hm := m.utoa(16)
						if debugFloat && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:390
		_go_fuzz_dep_.CoverTab[5097]++
//line /usr/local/go/src/math/big/ftoa.go:390
		return hm[0] != '1'
//line /usr/local/go/src/math/big/ftoa.go:390
		// _ = "end of CoverTab[5097]"
//line /usr/local/go/src/math/big/ftoa.go:390
	}() {
//line /usr/local/go/src/math/big/ftoa.go:390
		_go_fuzz_dep_.CoverTab[5098]++
							panic("incorrect mantissa: " + string(hm))
//line /usr/local/go/src/math/big/ftoa.go:391
		// _ = "end of CoverTab[5098]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:392
		_go_fuzz_dep_.CoverTab[5099]++
//line /usr/local/go/src/math/big/ftoa.go:392
		// _ = "end of CoverTab[5099]"
//line /usr/local/go/src/math/big/ftoa.go:392
	}
//line /usr/local/go/src/math/big/ftoa.go:392
	// _ = "end of CoverTab[5078]"
//line /usr/local/go/src/math/big/ftoa.go:392
	_go_fuzz_dep_.CoverTab[5079]++
						buf = append(buf, "0x1"...)
						if len(hm) > 1 {
//line /usr/local/go/src/math/big/ftoa.go:394
		_go_fuzz_dep_.CoverTab[5100]++
							buf = append(buf, '.')
							buf = append(buf, hm[1:]...)
//line /usr/local/go/src/math/big/ftoa.go:396
		// _ = "end of CoverTab[5100]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:397
		_go_fuzz_dep_.CoverTab[5101]++
//line /usr/local/go/src/math/big/ftoa.go:397
		// _ = "end of CoverTab[5101]"
//line /usr/local/go/src/math/big/ftoa.go:397
	}
//line /usr/local/go/src/math/big/ftoa.go:397
	// _ = "end of CoverTab[5079]"
//line /usr/local/go/src/math/big/ftoa.go:397
	_go_fuzz_dep_.CoverTab[5080]++

						buf = append(buf, 'p')
						if exp64 >= 0 {
//line /usr/local/go/src/math/big/ftoa.go:400
		_go_fuzz_dep_.CoverTab[5102]++
							buf = append(buf, '+')
//line /usr/local/go/src/math/big/ftoa.go:401
		// _ = "end of CoverTab[5102]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:402
		_go_fuzz_dep_.CoverTab[5103]++
							exp64 = -exp64
							buf = append(buf, '-')
//line /usr/local/go/src/math/big/ftoa.go:404
		// _ = "end of CoverTab[5103]"
	}
//line /usr/local/go/src/math/big/ftoa.go:405
	// _ = "end of CoverTab[5080]"
//line /usr/local/go/src/math/big/ftoa.go:405
	_go_fuzz_dep_.CoverTab[5081]++

						if exp64 < 10 {
//line /usr/local/go/src/math/big/ftoa.go:407
		_go_fuzz_dep_.CoverTab[5104]++
							buf = append(buf, '0')
//line /usr/local/go/src/math/big/ftoa.go:408
		// _ = "end of CoverTab[5104]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:409
		_go_fuzz_dep_.CoverTab[5105]++
//line /usr/local/go/src/math/big/ftoa.go:409
		// _ = "end of CoverTab[5105]"
//line /usr/local/go/src/math/big/ftoa.go:409
	}
//line /usr/local/go/src/math/big/ftoa.go:409
	// _ = "end of CoverTab[5081]"
//line /usr/local/go/src/math/big/ftoa.go:409
	_go_fuzz_dep_.CoverTab[5082]++
						return strconv.AppendInt(buf, exp64, 10)
//line /usr/local/go/src/math/big/ftoa.go:410
	// _ = "end of CoverTab[5082]"
}

// fmtP appends the string of x in the format "0x." mantissa "p" exponent
//line /usr/local/go/src/math/big/ftoa.go:413
// with a hexadecimal mantissa and a binary exponent, or "0" if x is zero,
//line /usr/local/go/src/math/big/ftoa.go:413
// and returns the extended buffer.
//line /usr/local/go/src/math/big/ftoa.go:413
// The mantissa is normalized such that 0.5 <= 0.mantissa < 1.0.
//line /usr/local/go/src/math/big/ftoa.go:413
// The sign of x is ignored, and x must not be an Inf.
//line /usr/local/go/src/math/big/ftoa.go:413
// (The caller handles Inf before invoking fmtP.)
//line /usr/local/go/src/math/big/ftoa.go:419
func (x *Float) fmtP(buf []byte) []byte {
//line /usr/local/go/src/math/big/ftoa.go:419
	_go_fuzz_dep_.CoverTab[5106]++
						if x.form == zero {
//line /usr/local/go/src/math/big/ftoa.go:420
		_go_fuzz_dep_.CoverTab[5111]++
							return append(buf, '0')
//line /usr/local/go/src/math/big/ftoa.go:421
		// _ = "end of CoverTab[5111]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:422
		_go_fuzz_dep_.CoverTab[5112]++
//line /usr/local/go/src/math/big/ftoa.go:422
		// _ = "end of CoverTab[5112]"
//line /usr/local/go/src/math/big/ftoa.go:422
	}
//line /usr/local/go/src/math/big/ftoa.go:422
	// _ = "end of CoverTab[5106]"
//line /usr/local/go/src/math/big/ftoa.go:422
	_go_fuzz_dep_.CoverTab[5107]++

						if debugFloat && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:424
		_go_fuzz_dep_.CoverTab[5113]++
//line /usr/local/go/src/math/big/ftoa.go:424
		return x.form != finite
//line /usr/local/go/src/math/big/ftoa.go:424
		// _ = "end of CoverTab[5113]"
//line /usr/local/go/src/math/big/ftoa.go:424
	}() {
//line /usr/local/go/src/math/big/ftoa.go:424
		_go_fuzz_dep_.CoverTab[5114]++
							panic("non-finite float")
//line /usr/local/go/src/math/big/ftoa.go:425
		// _ = "end of CoverTab[5114]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:426
		_go_fuzz_dep_.CoverTab[5115]++
//line /usr/local/go/src/math/big/ftoa.go:426
		// _ = "end of CoverTab[5115]"
//line /usr/local/go/src/math/big/ftoa.go:426
	}
//line /usr/local/go/src/math/big/ftoa.go:426
	// _ = "end of CoverTab[5107]"
//line /usr/local/go/src/math/big/ftoa.go:426
	_go_fuzz_dep_.CoverTab[5108]++

//line /usr/local/go/src/math/big/ftoa.go:431
	m := x.mant
	i := 0
	for i < len(m) && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:433
		_go_fuzz_dep_.CoverTab[5116]++
//line /usr/local/go/src/math/big/ftoa.go:433
		return m[i] == 0
//line /usr/local/go/src/math/big/ftoa.go:433
		// _ = "end of CoverTab[5116]"
//line /usr/local/go/src/math/big/ftoa.go:433
	}() {
//line /usr/local/go/src/math/big/ftoa.go:433
		_go_fuzz_dep_.CoverTab[5117]++
							i++
//line /usr/local/go/src/math/big/ftoa.go:434
		// _ = "end of CoverTab[5117]"
	}
//line /usr/local/go/src/math/big/ftoa.go:435
	// _ = "end of CoverTab[5108]"
//line /usr/local/go/src/math/big/ftoa.go:435
	_go_fuzz_dep_.CoverTab[5109]++
						m = m[i:]

						buf = append(buf, "0x."...)
						buf = append(buf, bytes.TrimRight(m.utoa(16), "0")...)
						buf = append(buf, 'p')
						if x.exp >= 0 {
//line /usr/local/go/src/math/big/ftoa.go:441
		_go_fuzz_dep_.CoverTab[5118]++
							buf = append(buf, '+')
//line /usr/local/go/src/math/big/ftoa.go:442
		// _ = "end of CoverTab[5118]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:443
		_go_fuzz_dep_.CoverTab[5119]++
//line /usr/local/go/src/math/big/ftoa.go:443
		// _ = "end of CoverTab[5119]"
//line /usr/local/go/src/math/big/ftoa.go:443
	}
//line /usr/local/go/src/math/big/ftoa.go:443
	// _ = "end of CoverTab[5109]"
//line /usr/local/go/src/math/big/ftoa.go:443
	_go_fuzz_dep_.CoverTab[5110]++
						return strconv.AppendInt(buf, int64(x.exp), 10)
//line /usr/local/go/src/math/big/ftoa.go:444
	// _ = "end of CoverTab[5110]"
}

func min(x, y int) int {
//line /usr/local/go/src/math/big/ftoa.go:447
	_go_fuzz_dep_.CoverTab[5120]++
						if x < y {
//line /usr/local/go/src/math/big/ftoa.go:448
		_go_fuzz_dep_.CoverTab[5122]++
							return x
//line /usr/local/go/src/math/big/ftoa.go:449
		// _ = "end of CoverTab[5122]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:450
		_go_fuzz_dep_.CoverTab[5123]++
//line /usr/local/go/src/math/big/ftoa.go:450
		// _ = "end of CoverTab[5123]"
//line /usr/local/go/src/math/big/ftoa.go:450
	}
//line /usr/local/go/src/math/big/ftoa.go:450
	// _ = "end of CoverTab[5120]"
//line /usr/local/go/src/math/big/ftoa.go:450
	_go_fuzz_dep_.CoverTab[5121]++
						return y
//line /usr/local/go/src/math/big/ftoa.go:451
	// _ = "end of CoverTab[5121]"
}

var _ fmt.Formatter = &floatZero	// *Float must implement fmt.Formatter

// Format implements fmt.Formatter. It accepts all the regular
//line /usr/local/go/src/math/big/ftoa.go:456
// formats for floating-point numbers ('b', 'e', 'E', 'f', 'F',
//line /usr/local/go/src/math/big/ftoa.go:456
// 'g', 'G', 'x') as well as 'p' and 'v'. See (*Float).Text for the
//line /usr/local/go/src/math/big/ftoa.go:456
// interpretation of 'p'. The 'v' format is handled like 'g'.
//line /usr/local/go/src/math/big/ftoa.go:456
// Format also supports specification of the minimum precision
//line /usr/local/go/src/math/big/ftoa.go:456
// in digits, the output field width, as well as the format flags
//line /usr/local/go/src/math/big/ftoa.go:456
// '+' and ' ' for sign control, '0' for space or zero padding,
//line /usr/local/go/src/math/big/ftoa.go:456
// and '-' for left or right justification. See the fmt package
//line /usr/local/go/src/math/big/ftoa.go:456
// for details.
//line /usr/local/go/src/math/big/ftoa.go:465
func (x *Float) Format(s fmt.State, format rune) {
//line /usr/local/go/src/math/big/ftoa.go:465
	_go_fuzz_dep_.CoverTab[5124]++
						prec, hasPrec := s.Precision()
						if !hasPrec {
//line /usr/local/go/src/math/big/ftoa.go:467
		_go_fuzz_dep_.CoverTab[5130]++
							prec = 6
//line /usr/local/go/src/math/big/ftoa.go:468
		// _ = "end of CoverTab[5130]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:469
		_go_fuzz_dep_.CoverTab[5131]++
//line /usr/local/go/src/math/big/ftoa.go:469
		// _ = "end of CoverTab[5131]"
//line /usr/local/go/src/math/big/ftoa.go:469
	}
//line /usr/local/go/src/math/big/ftoa.go:469
	// _ = "end of CoverTab[5124]"
//line /usr/local/go/src/math/big/ftoa.go:469
	_go_fuzz_dep_.CoverTab[5125]++

						switch format {
	case 'e', 'E', 'f', 'b', 'p', 'x':
//line /usr/local/go/src/math/big/ftoa.go:472
		_go_fuzz_dep_.CoverTab[5132]++
//line /usr/local/go/src/math/big/ftoa.go:472
		// _ = "end of CoverTab[5132]"

	case 'F':
//line /usr/local/go/src/math/big/ftoa.go:474
		_go_fuzz_dep_.CoverTab[5133]++

							format = 'f'
//line /usr/local/go/src/math/big/ftoa.go:476
		// _ = "end of CoverTab[5133]"
	case 'v':
//line /usr/local/go/src/math/big/ftoa.go:477
		_go_fuzz_dep_.CoverTab[5134]++

							format = 'g'
							fallthrough
//line /usr/local/go/src/math/big/ftoa.go:480
		// _ = "end of CoverTab[5134]"
	case 'g', 'G':
//line /usr/local/go/src/math/big/ftoa.go:481
		_go_fuzz_dep_.CoverTab[5135]++
							if !hasPrec {
//line /usr/local/go/src/math/big/ftoa.go:482
			_go_fuzz_dep_.CoverTab[5137]++
								prec = -1
//line /usr/local/go/src/math/big/ftoa.go:483
			// _ = "end of CoverTab[5137]"
		} else {
//line /usr/local/go/src/math/big/ftoa.go:484
			_go_fuzz_dep_.CoverTab[5138]++
//line /usr/local/go/src/math/big/ftoa.go:484
			// _ = "end of CoverTab[5138]"
//line /usr/local/go/src/math/big/ftoa.go:484
		}
//line /usr/local/go/src/math/big/ftoa.go:484
		// _ = "end of CoverTab[5135]"
	default:
//line /usr/local/go/src/math/big/ftoa.go:485
		_go_fuzz_dep_.CoverTab[5136]++
							fmt.Fprintf(s, "%%!%c(*big.Float=%s)", format, x.String())
							return
//line /usr/local/go/src/math/big/ftoa.go:487
		// _ = "end of CoverTab[5136]"
	}
//line /usr/local/go/src/math/big/ftoa.go:488
	// _ = "end of CoverTab[5125]"
//line /usr/local/go/src/math/big/ftoa.go:488
	_go_fuzz_dep_.CoverTab[5126]++
						var buf []byte
						buf = x.Append(buf, byte(format), prec)
						if len(buf) == 0 {
//line /usr/local/go/src/math/big/ftoa.go:491
		_go_fuzz_dep_.CoverTab[5139]++
							buf = []byte("?")
//line /usr/local/go/src/math/big/ftoa.go:492
		// _ = "end of CoverTab[5139]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:493
		_go_fuzz_dep_.CoverTab[5140]++
//line /usr/local/go/src/math/big/ftoa.go:493
		// _ = "end of CoverTab[5140]"
//line /usr/local/go/src/math/big/ftoa.go:493
	}
//line /usr/local/go/src/math/big/ftoa.go:493
	// _ = "end of CoverTab[5126]"
//line /usr/local/go/src/math/big/ftoa.go:493
	_go_fuzz_dep_.CoverTab[5127]++

//line /usr/local/go/src/math/big/ftoa.go:496
	var sign string
	switch {
	case buf[0] == '-':
//line /usr/local/go/src/math/big/ftoa.go:498
		_go_fuzz_dep_.CoverTab[5141]++
							sign = "-"
							buf = buf[1:]
//line /usr/local/go/src/math/big/ftoa.go:500
		// _ = "end of CoverTab[5141]"
	case buf[0] == '+':
//line /usr/local/go/src/math/big/ftoa.go:501
		_go_fuzz_dep_.CoverTab[5142]++

							sign = "+"
							if s.Flag(' ') {
//line /usr/local/go/src/math/big/ftoa.go:504
			_go_fuzz_dep_.CoverTab[5147]++
								sign = " "
//line /usr/local/go/src/math/big/ftoa.go:505
			// _ = "end of CoverTab[5147]"
		} else {
//line /usr/local/go/src/math/big/ftoa.go:506
			_go_fuzz_dep_.CoverTab[5148]++
//line /usr/local/go/src/math/big/ftoa.go:506
			// _ = "end of CoverTab[5148]"
//line /usr/local/go/src/math/big/ftoa.go:506
		}
//line /usr/local/go/src/math/big/ftoa.go:506
		// _ = "end of CoverTab[5142]"
//line /usr/local/go/src/math/big/ftoa.go:506
		_go_fuzz_dep_.CoverTab[5143]++
							buf = buf[1:]
//line /usr/local/go/src/math/big/ftoa.go:507
		// _ = "end of CoverTab[5143]"
	case s.Flag('+'):
//line /usr/local/go/src/math/big/ftoa.go:508
		_go_fuzz_dep_.CoverTab[5144]++
							sign = "+"
//line /usr/local/go/src/math/big/ftoa.go:509
		// _ = "end of CoverTab[5144]"
	case s.Flag(' '):
//line /usr/local/go/src/math/big/ftoa.go:510
		_go_fuzz_dep_.CoverTab[5145]++
							sign = " "
//line /usr/local/go/src/math/big/ftoa.go:511
		// _ = "end of CoverTab[5145]"
//line /usr/local/go/src/math/big/ftoa.go:511
	default:
//line /usr/local/go/src/math/big/ftoa.go:511
		_go_fuzz_dep_.CoverTab[5146]++
//line /usr/local/go/src/math/big/ftoa.go:511
		// _ = "end of CoverTab[5146]"
	}
//line /usr/local/go/src/math/big/ftoa.go:512
	// _ = "end of CoverTab[5127]"
//line /usr/local/go/src/math/big/ftoa.go:512
	_go_fuzz_dep_.CoverTab[5128]++

						var padding int
						if width, hasWidth := s.Width(); hasWidth && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:515
		_go_fuzz_dep_.CoverTab[5149]++
//line /usr/local/go/src/math/big/ftoa.go:515
		return width > len(sign)+len(buf)
//line /usr/local/go/src/math/big/ftoa.go:515
		// _ = "end of CoverTab[5149]"
//line /usr/local/go/src/math/big/ftoa.go:515
	}() {
//line /usr/local/go/src/math/big/ftoa.go:515
		_go_fuzz_dep_.CoverTab[5150]++
							padding = width - len(sign) - len(buf)
//line /usr/local/go/src/math/big/ftoa.go:516
		// _ = "end of CoverTab[5150]"
	} else {
//line /usr/local/go/src/math/big/ftoa.go:517
		_go_fuzz_dep_.CoverTab[5151]++
//line /usr/local/go/src/math/big/ftoa.go:517
		// _ = "end of CoverTab[5151]"
//line /usr/local/go/src/math/big/ftoa.go:517
	}
//line /usr/local/go/src/math/big/ftoa.go:517
	// _ = "end of CoverTab[5128]"
//line /usr/local/go/src/math/big/ftoa.go:517
	_go_fuzz_dep_.CoverTab[5129]++

						switch {
	case s.Flag('0') && func() bool {
//line /usr/local/go/src/math/big/ftoa.go:520
		_go_fuzz_dep_.CoverTab[5155]++
//line /usr/local/go/src/math/big/ftoa.go:520
		return !x.IsInf()
//line /usr/local/go/src/math/big/ftoa.go:520
		// _ = "end of CoverTab[5155]"
//line /usr/local/go/src/math/big/ftoa.go:520
	}():
//line /usr/local/go/src/math/big/ftoa.go:520
		_go_fuzz_dep_.CoverTab[5152]++

							writeMultiple(s, sign, 1)
							writeMultiple(s, "0", padding)
							s.Write(buf)
//line /usr/local/go/src/math/big/ftoa.go:524
		// _ = "end of CoverTab[5152]"
	case s.Flag('-'):
//line /usr/local/go/src/math/big/ftoa.go:525
		_go_fuzz_dep_.CoverTab[5153]++

							writeMultiple(s, sign, 1)
							s.Write(buf)
							writeMultiple(s, " ", padding)
//line /usr/local/go/src/math/big/ftoa.go:529
		// _ = "end of CoverTab[5153]"
	default:
//line /usr/local/go/src/math/big/ftoa.go:530
		_go_fuzz_dep_.CoverTab[5154]++

							writeMultiple(s, " ", padding)
							writeMultiple(s, sign, 1)
							s.Write(buf)
//line /usr/local/go/src/math/big/ftoa.go:534
		// _ = "end of CoverTab[5154]"
	}
//line /usr/local/go/src/math/big/ftoa.go:535
	// _ = "end of CoverTab[5129]"
}

//line /usr/local/go/src/math/big/ftoa.go:536
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/ftoa.go:536
var _ = _go_fuzz_dep_.CoverTab
