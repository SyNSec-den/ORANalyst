// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements multi-precision decimal numbers.
// The implementation is for float to decimal conversion only;
// not general purpose use.
// The only operations are precise conversion from binary to
// decimal and rounding.
//
// The key observation and some code (shr) is borrowed from
// strconv/decimal.go: conversion of binary fractional values can be done
// precisely in multi-precision decimal because 2 divides 10 (required for
// >> of mantissa); but conversion of decimal floating-point values cannot
// be done precisely in binary representation.
//
// In contrast to strconv/decimal.go, only right shift is implemented in
// decimal format - left shift can be done precisely in binary format.

//line /usr/local/go/src/math/big/decimal.go:20
package big

//line /usr/local/go/src/math/big/decimal.go:20
import (
//line /usr/local/go/src/math/big/decimal.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/decimal.go:20
)
//line /usr/local/go/src/math/big/decimal.go:20
import (
//line /usr/local/go/src/math/big/decimal.go:20
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/decimal.go:20
)

// A decimal represents an unsigned floating-point number in decimal representation.
//line /usr/local/go/src/math/big/decimal.go:22
// The value of a non-zero decimal d is d.mant * 10**d.exp with 0.1 <= d.mant < 1,
//line /usr/local/go/src/math/big/decimal.go:22
// with the most-significant mantissa digit at index 0. For the zero decimal, the
//line /usr/local/go/src/math/big/decimal.go:22
// mantissa length and exponent are 0.
//line /usr/local/go/src/math/big/decimal.go:22
// The zero value for decimal represents a ready-to-use 0.0.
//line /usr/local/go/src/math/big/decimal.go:27
type decimal struct {
	mant	[]byte	// mantissa ASCII digits, big-endian
	exp	int	// exponent
}

// at returns the i'th mantissa digit, starting with the most significant digit at 0.
func (d *decimal) at(i int) byte {
//line /usr/local/go/src/math/big/decimal.go:33
	_go_fuzz_dep_.CoverTab[4121]++
							if 0 <= i && func() bool {
//line /usr/local/go/src/math/big/decimal.go:34
		_go_fuzz_dep_.CoverTab[4123]++
//line /usr/local/go/src/math/big/decimal.go:34
		return i < len(d.mant)
//line /usr/local/go/src/math/big/decimal.go:34
		// _ = "end of CoverTab[4123]"
//line /usr/local/go/src/math/big/decimal.go:34
	}() {
//line /usr/local/go/src/math/big/decimal.go:34
		_go_fuzz_dep_.CoverTab[4124]++
								return d.mant[i]
//line /usr/local/go/src/math/big/decimal.go:35
		// _ = "end of CoverTab[4124]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:36
		_go_fuzz_dep_.CoverTab[4125]++
//line /usr/local/go/src/math/big/decimal.go:36
		// _ = "end of CoverTab[4125]"
//line /usr/local/go/src/math/big/decimal.go:36
	}
//line /usr/local/go/src/math/big/decimal.go:36
	// _ = "end of CoverTab[4121]"
//line /usr/local/go/src/math/big/decimal.go:36
	_go_fuzz_dep_.CoverTab[4122]++
							return '0'
//line /usr/local/go/src/math/big/decimal.go:37
	// _ = "end of CoverTab[4122]"
}

// Maximum shift amount that can be done in one pass without overflow.
//line /usr/local/go/src/math/big/decimal.go:40
// A Word has _W bits and (1<<maxShift - 1)*10 + 9 must fit into Word.
//line /usr/local/go/src/math/big/decimal.go:42
const maxShift = _W - 4

//line /usr/local/go/src/math/big/decimal.go:53
// Init initializes x to the decimal representation of m << shift (for
//line /usr/local/go/src/math/big/decimal.go:53
// shift >= 0), or m >> -shift (for shift < 0).
//line /usr/local/go/src/math/big/decimal.go:55
func (x *decimal) init(m nat, shift int) {

	if len(m) == 0 {
		x.mant = x.mant[:0]
		x.exp = 0
		return
	}

//line /usr/local/go/src/math/big/decimal.go:66
	if shift < 0 {
		ntz := m.trailingZeroBits()
		s := uint(-shift)
		if s >= ntz {
			s = ntz
		}
		m = nat(nil).shr(m, s)
		shift += int(s)
	}

//line /usr/local/go/src/math/big/decimal.go:77
	if shift > 0 {
		m = nat(nil).shl(m, uint(shift))
		shift = 0
	}

//line /usr/local/go/src/math/big/decimal.go:83
	s := m.utoa(10)
							n := len(s)
							x.exp = n

//line /usr/local/go/src/math/big/decimal.go:88
	for n > 0 && s[n-1] == '0' {
		n--
	}
							x.mant = append(x.mant[:0], s[:n]...)

//line /usr/local/go/src/math/big/decimal.go:94
	if shift < 0 {
		for shift < -maxShift {
			shr(x, maxShift)
			shift += maxShift
		}
		shr(x, uint(-shift))
	}
}

// shr implements x >> s, for s <= maxShift.
func shr(x *decimal, s uint) {
//line /usr/local/go/src/math/big/decimal.go:104
	_go_fuzz_dep_.CoverTab[4126]++

//line /usr/local/go/src/math/big/decimal.go:108
	r := 0
	var n Word
	for n>>s == 0 && func() bool {
//line /usr/local/go/src/math/big/decimal.go:110
		_go_fuzz_dep_.CoverTab[4133]++
//line /usr/local/go/src/math/big/decimal.go:110
		return r < len(x.mant)
//line /usr/local/go/src/math/big/decimal.go:110
		// _ = "end of CoverTab[4133]"
//line /usr/local/go/src/math/big/decimal.go:110
	}() {
//line /usr/local/go/src/math/big/decimal.go:110
		_go_fuzz_dep_.CoverTab[4134]++
								ch := Word(x.mant[r])
								r++
								n = n*10 + ch - '0'
//line /usr/local/go/src/math/big/decimal.go:113
		// _ = "end of CoverTab[4134]"
	}
//line /usr/local/go/src/math/big/decimal.go:114
	// _ = "end of CoverTab[4126]"
//line /usr/local/go/src/math/big/decimal.go:114
	_go_fuzz_dep_.CoverTab[4127]++
							if n == 0 {
//line /usr/local/go/src/math/big/decimal.go:115
		_go_fuzz_dep_.CoverTab[4135]++

								x.mant = x.mant[:0]
								return
//line /usr/local/go/src/math/big/decimal.go:118
		// _ = "end of CoverTab[4135]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:119
		_go_fuzz_dep_.CoverTab[4136]++
//line /usr/local/go/src/math/big/decimal.go:119
		// _ = "end of CoverTab[4136]"
//line /usr/local/go/src/math/big/decimal.go:119
	}
//line /usr/local/go/src/math/big/decimal.go:119
	// _ = "end of CoverTab[4127]"
//line /usr/local/go/src/math/big/decimal.go:119
	_go_fuzz_dep_.CoverTab[4128]++
							for n>>s == 0 {
//line /usr/local/go/src/math/big/decimal.go:120
		_go_fuzz_dep_.CoverTab[4137]++
								r++
								n *= 10
//line /usr/local/go/src/math/big/decimal.go:122
		// _ = "end of CoverTab[4137]"
	}
//line /usr/local/go/src/math/big/decimal.go:123
	// _ = "end of CoverTab[4128]"
//line /usr/local/go/src/math/big/decimal.go:123
	_go_fuzz_dep_.CoverTab[4129]++
							x.exp += 1 - r

//line /usr/local/go/src/math/big/decimal.go:127
	w := 0
	mask := Word(1)<<s - 1
	for r < len(x.mant) {
//line /usr/local/go/src/math/big/decimal.go:129
		_go_fuzz_dep_.CoverTab[4138]++
								ch := Word(x.mant[r])
								r++
								d := n >> s
								n &= mask
								x.mant[w] = byte(d + '0')
								w++
								n = n*10 + ch - '0'
//line /usr/local/go/src/math/big/decimal.go:136
		// _ = "end of CoverTab[4138]"
	}
//line /usr/local/go/src/math/big/decimal.go:137
	// _ = "end of CoverTab[4129]"
//line /usr/local/go/src/math/big/decimal.go:137
	_go_fuzz_dep_.CoverTab[4130]++

//line /usr/local/go/src/math/big/decimal.go:140
	for n > 0 && func() bool {
//line /usr/local/go/src/math/big/decimal.go:140
		_go_fuzz_dep_.CoverTab[4139]++
//line /usr/local/go/src/math/big/decimal.go:140
		return w < len(x.mant)
//line /usr/local/go/src/math/big/decimal.go:140
		// _ = "end of CoverTab[4139]"
//line /usr/local/go/src/math/big/decimal.go:140
	}() {
//line /usr/local/go/src/math/big/decimal.go:140
		_go_fuzz_dep_.CoverTab[4140]++
								d := n >> s
								n &= mask
								x.mant[w] = byte(d + '0')
								w++
								n = n * 10
//line /usr/local/go/src/math/big/decimal.go:145
		// _ = "end of CoverTab[4140]"
	}
//line /usr/local/go/src/math/big/decimal.go:146
	// _ = "end of CoverTab[4130]"
//line /usr/local/go/src/math/big/decimal.go:146
	_go_fuzz_dep_.CoverTab[4131]++
							x.mant = x.mant[:w]

//line /usr/local/go/src/math/big/decimal.go:150
	for n > 0 {
//line /usr/local/go/src/math/big/decimal.go:150
		_go_fuzz_dep_.CoverTab[4141]++
								d := n >> s
								n &= mask
								x.mant = append(x.mant, byte(d+'0'))
								n = n * 10
//line /usr/local/go/src/math/big/decimal.go:154
		// _ = "end of CoverTab[4141]"
	}
//line /usr/local/go/src/math/big/decimal.go:155
	// _ = "end of CoverTab[4131]"
//line /usr/local/go/src/math/big/decimal.go:155
	_go_fuzz_dep_.CoverTab[4132]++

							trim(x)
//line /usr/local/go/src/math/big/decimal.go:157
	// _ = "end of CoverTab[4132]"
}

func (x *decimal) String() string {
//line /usr/local/go/src/math/big/decimal.go:160
	_go_fuzz_dep_.CoverTab[4142]++
							if len(x.mant) == 0 {
//line /usr/local/go/src/math/big/decimal.go:161
		_go_fuzz_dep_.CoverTab[4145]++
								return "0"
//line /usr/local/go/src/math/big/decimal.go:162
		// _ = "end of CoverTab[4145]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:163
		_go_fuzz_dep_.CoverTab[4146]++
//line /usr/local/go/src/math/big/decimal.go:163
		// _ = "end of CoverTab[4146]"
//line /usr/local/go/src/math/big/decimal.go:163
	}
//line /usr/local/go/src/math/big/decimal.go:163
	// _ = "end of CoverTab[4142]"
//line /usr/local/go/src/math/big/decimal.go:163
	_go_fuzz_dep_.CoverTab[4143]++

							var buf []byte
							switch {
	case x.exp <= 0:
//line /usr/local/go/src/math/big/decimal.go:167
		_go_fuzz_dep_.CoverTab[4147]++

								buf = make([]byte, 0, 2+(-x.exp)+len(x.mant))
								buf = append(buf, "0."...)
								buf = appendZeros(buf, -x.exp)
								buf = append(buf, x.mant...)
//line /usr/local/go/src/math/big/decimal.go:172
		// _ = "end of CoverTab[4147]"

	case x.exp < len(x.mant):
//line /usr/local/go/src/math/big/decimal.go:174
		_go_fuzz_dep_.CoverTab[4148]++

								buf = make([]byte, 0, 1+len(x.mant))
								buf = append(buf, x.mant[:x.exp]...)
								buf = append(buf, '.')
								buf = append(buf, x.mant[x.exp:]...)
//line /usr/local/go/src/math/big/decimal.go:179
		// _ = "end of CoverTab[4148]"

	default:
//line /usr/local/go/src/math/big/decimal.go:181
		_go_fuzz_dep_.CoverTab[4149]++

								buf = make([]byte, 0, x.exp)
								buf = append(buf, x.mant...)
								buf = appendZeros(buf, x.exp-len(x.mant))
//line /usr/local/go/src/math/big/decimal.go:185
		// _ = "end of CoverTab[4149]"
	}
//line /usr/local/go/src/math/big/decimal.go:186
	// _ = "end of CoverTab[4143]"
//line /usr/local/go/src/math/big/decimal.go:186
	_go_fuzz_dep_.CoverTab[4144]++

							return string(buf)
//line /usr/local/go/src/math/big/decimal.go:188
	// _ = "end of CoverTab[4144]"
}

// appendZeros appends n 0 digits to buf and returns buf.
func appendZeros(buf []byte, n int) []byte {
//line /usr/local/go/src/math/big/decimal.go:192
	_go_fuzz_dep_.CoverTab[4150]++
							for ; n > 0; n-- {
//line /usr/local/go/src/math/big/decimal.go:193
		_go_fuzz_dep_.CoverTab[4152]++
								buf = append(buf, '0')
//line /usr/local/go/src/math/big/decimal.go:194
		// _ = "end of CoverTab[4152]"
	}
//line /usr/local/go/src/math/big/decimal.go:195
	// _ = "end of CoverTab[4150]"
//line /usr/local/go/src/math/big/decimal.go:195
	_go_fuzz_dep_.CoverTab[4151]++
							return buf
//line /usr/local/go/src/math/big/decimal.go:196
	// _ = "end of CoverTab[4151]"
}

// shouldRoundUp reports if x should be rounded up
//line /usr/local/go/src/math/big/decimal.go:199
// if shortened to n digits. n must be a valid index
//line /usr/local/go/src/math/big/decimal.go:199
// for x.mant.
//line /usr/local/go/src/math/big/decimal.go:202
func shouldRoundUp(x *decimal, n int) bool {
//line /usr/local/go/src/math/big/decimal.go:202
	_go_fuzz_dep_.CoverTab[4153]++
							if x.mant[n] == '5' && func() bool {
//line /usr/local/go/src/math/big/decimal.go:203
		_go_fuzz_dep_.CoverTab[4155]++
//line /usr/local/go/src/math/big/decimal.go:203
		return n+1 == len(x.mant)
//line /usr/local/go/src/math/big/decimal.go:203
		// _ = "end of CoverTab[4155]"
//line /usr/local/go/src/math/big/decimal.go:203
	}() {
//line /usr/local/go/src/math/big/decimal.go:203
		_go_fuzz_dep_.CoverTab[4156]++

								return n > 0 && func() bool {
//line /usr/local/go/src/math/big/decimal.go:205
			_go_fuzz_dep_.CoverTab[4157]++
//line /usr/local/go/src/math/big/decimal.go:205
			return (x.mant[n-1]-'0')&1 != 0
//line /usr/local/go/src/math/big/decimal.go:205
			// _ = "end of CoverTab[4157]"
//line /usr/local/go/src/math/big/decimal.go:205
		}()
//line /usr/local/go/src/math/big/decimal.go:205
		// _ = "end of CoverTab[4156]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:206
		_go_fuzz_dep_.CoverTab[4158]++
//line /usr/local/go/src/math/big/decimal.go:206
		// _ = "end of CoverTab[4158]"
//line /usr/local/go/src/math/big/decimal.go:206
	}
//line /usr/local/go/src/math/big/decimal.go:206
	// _ = "end of CoverTab[4153]"
//line /usr/local/go/src/math/big/decimal.go:206
	_go_fuzz_dep_.CoverTab[4154]++

							return x.mant[n] >= '5'
//line /usr/local/go/src/math/big/decimal.go:208
	// _ = "end of CoverTab[4154]"
}

// round sets x to (at most) n mantissa digits by rounding it
//line /usr/local/go/src/math/big/decimal.go:211
// to the nearest even value with n (or fever) mantissa digits.
//line /usr/local/go/src/math/big/decimal.go:211
// If n < 0, x remains unchanged.
//line /usr/local/go/src/math/big/decimal.go:214
func (x *decimal) round(n int) {
//line /usr/local/go/src/math/big/decimal.go:214
	_go_fuzz_dep_.CoverTab[4159]++
							if n < 0 || func() bool {
//line /usr/local/go/src/math/big/decimal.go:215
		_go_fuzz_dep_.CoverTab[4161]++
//line /usr/local/go/src/math/big/decimal.go:215
		return n >= len(x.mant)
//line /usr/local/go/src/math/big/decimal.go:215
		// _ = "end of CoverTab[4161]"
//line /usr/local/go/src/math/big/decimal.go:215
	}() {
//line /usr/local/go/src/math/big/decimal.go:215
		_go_fuzz_dep_.CoverTab[4162]++
								return
//line /usr/local/go/src/math/big/decimal.go:216
		// _ = "end of CoverTab[4162]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:217
		_go_fuzz_dep_.CoverTab[4163]++
//line /usr/local/go/src/math/big/decimal.go:217
		// _ = "end of CoverTab[4163]"
//line /usr/local/go/src/math/big/decimal.go:217
	}
//line /usr/local/go/src/math/big/decimal.go:217
	// _ = "end of CoverTab[4159]"
//line /usr/local/go/src/math/big/decimal.go:217
	_go_fuzz_dep_.CoverTab[4160]++

							if shouldRoundUp(x, n) {
//line /usr/local/go/src/math/big/decimal.go:219
		_go_fuzz_dep_.CoverTab[4164]++
								x.roundUp(n)
//line /usr/local/go/src/math/big/decimal.go:220
		// _ = "end of CoverTab[4164]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:221
		_go_fuzz_dep_.CoverTab[4165]++
								x.roundDown(n)
//line /usr/local/go/src/math/big/decimal.go:222
		// _ = "end of CoverTab[4165]"
	}
//line /usr/local/go/src/math/big/decimal.go:223
	// _ = "end of CoverTab[4160]"
}

func (x *decimal) roundUp(n int) {
//line /usr/local/go/src/math/big/decimal.go:226
	_go_fuzz_dep_.CoverTab[4166]++
							if n < 0 || func() bool {
//line /usr/local/go/src/math/big/decimal.go:227
		_go_fuzz_dep_.CoverTab[4170]++
//line /usr/local/go/src/math/big/decimal.go:227
		return n >= len(x.mant)
//line /usr/local/go/src/math/big/decimal.go:227
		// _ = "end of CoverTab[4170]"
//line /usr/local/go/src/math/big/decimal.go:227
	}() {
//line /usr/local/go/src/math/big/decimal.go:227
		_go_fuzz_dep_.CoverTab[4171]++
								return
//line /usr/local/go/src/math/big/decimal.go:228
		// _ = "end of CoverTab[4171]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:229
		_go_fuzz_dep_.CoverTab[4172]++
//line /usr/local/go/src/math/big/decimal.go:229
		// _ = "end of CoverTab[4172]"
//line /usr/local/go/src/math/big/decimal.go:229
	}
//line /usr/local/go/src/math/big/decimal.go:229
	// _ = "end of CoverTab[4166]"
//line /usr/local/go/src/math/big/decimal.go:229
	_go_fuzz_dep_.CoverTab[4167]++

//line /usr/local/go/src/math/big/decimal.go:233
	for n > 0 && func() bool {
//line /usr/local/go/src/math/big/decimal.go:233
		_go_fuzz_dep_.CoverTab[4173]++
//line /usr/local/go/src/math/big/decimal.go:233
		return x.mant[n-1] >= '9'
//line /usr/local/go/src/math/big/decimal.go:233
		// _ = "end of CoverTab[4173]"
//line /usr/local/go/src/math/big/decimal.go:233
	}() {
//line /usr/local/go/src/math/big/decimal.go:233
		_go_fuzz_dep_.CoverTab[4174]++
								n--
//line /usr/local/go/src/math/big/decimal.go:234
		// _ = "end of CoverTab[4174]"
	}
//line /usr/local/go/src/math/big/decimal.go:235
	// _ = "end of CoverTab[4167]"
//line /usr/local/go/src/math/big/decimal.go:235
	_go_fuzz_dep_.CoverTab[4168]++

							if n == 0 {
//line /usr/local/go/src/math/big/decimal.go:237
		_go_fuzz_dep_.CoverTab[4175]++

								x.mant[0] = '1'
								x.mant = x.mant[:1]
								x.exp++
								return
//line /usr/local/go/src/math/big/decimal.go:242
		// _ = "end of CoverTab[4175]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:243
		_go_fuzz_dep_.CoverTab[4176]++
//line /usr/local/go/src/math/big/decimal.go:243
		// _ = "end of CoverTab[4176]"
//line /usr/local/go/src/math/big/decimal.go:243
	}
//line /usr/local/go/src/math/big/decimal.go:243
	// _ = "end of CoverTab[4168]"
//line /usr/local/go/src/math/big/decimal.go:243
	_go_fuzz_dep_.CoverTab[4169]++

//line /usr/local/go/src/math/big/decimal.go:246
	x.mant[n-1]++
							x.mant = x.mant[:n]
//line /usr/local/go/src/math/big/decimal.go:247
	// _ = "end of CoverTab[4169]"

}

func (x *decimal) roundDown(n int) {
//line /usr/local/go/src/math/big/decimal.go:251
	_go_fuzz_dep_.CoverTab[4177]++
							if n < 0 || func() bool {
//line /usr/local/go/src/math/big/decimal.go:252
		_go_fuzz_dep_.CoverTab[4179]++
//line /usr/local/go/src/math/big/decimal.go:252
		return n >= len(x.mant)
//line /usr/local/go/src/math/big/decimal.go:252
		// _ = "end of CoverTab[4179]"
//line /usr/local/go/src/math/big/decimal.go:252
	}() {
//line /usr/local/go/src/math/big/decimal.go:252
		_go_fuzz_dep_.CoverTab[4180]++
								return
//line /usr/local/go/src/math/big/decimal.go:253
		// _ = "end of CoverTab[4180]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:254
		_go_fuzz_dep_.CoverTab[4181]++
//line /usr/local/go/src/math/big/decimal.go:254
		// _ = "end of CoverTab[4181]"
//line /usr/local/go/src/math/big/decimal.go:254
	}
//line /usr/local/go/src/math/big/decimal.go:254
	// _ = "end of CoverTab[4177]"
//line /usr/local/go/src/math/big/decimal.go:254
	_go_fuzz_dep_.CoverTab[4178]++
							x.mant = x.mant[:n]
							trim(x)
//line /usr/local/go/src/math/big/decimal.go:256
	// _ = "end of CoverTab[4178]"
}

// trim cuts off any trailing zeros from x's mantissa;
//line /usr/local/go/src/math/big/decimal.go:259
// they are meaningless for the value of x.
//line /usr/local/go/src/math/big/decimal.go:261
func trim(x *decimal) {
//line /usr/local/go/src/math/big/decimal.go:261
	_go_fuzz_dep_.CoverTab[4182]++
							i := len(x.mant)
							for i > 0 && func() bool {
//line /usr/local/go/src/math/big/decimal.go:263
		_go_fuzz_dep_.CoverTab[4184]++
//line /usr/local/go/src/math/big/decimal.go:263
		return x.mant[i-1] == '0'
//line /usr/local/go/src/math/big/decimal.go:263
		// _ = "end of CoverTab[4184]"
//line /usr/local/go/src/math/big/decimal.go:263
	}() {
//line /usr/local/go/src/math/big/decimal.go:263
		_go_fuzz_dep_.CoverTab[4185]++
								i--
//line /usr/local/go/src/math/big/decimal.go:264
		// _ = "end of CoverTab[4185]"
	}
//line /usr/local/go/src/math/big/decimal.go:265
	// _ = "end of CoverTab[4182]"
//line /usr/local/go/src/math/big/decimal.go:265
	_go_fuzz_dep_.CoverTab[4183]++
							x.mant = x.mant[:i]
							if i == 0 {
//line /usr/local/go/src/math/big/decimal.go:267
		_go_fuzz_dep_.CoverTab[4186]++
								x.exp = 0
//line /usr/local/go/src/math/big/decimal.go:268
		// _ = "end of CoverTab[4186]"
	} else {
//line /usr/local/go/src/math/big/decimal.go:269
		_go_fuzz_dep_.CoverTab[4187]++
//line /usr/local/go/src/math/big/decimal.go:269
		// _ = "end of CoverTab[4187]"
//line /usr/local/go/src/math/big/decimal.go:269
	}
//line /usr/local/go/src/math/big/decimal.go:269
	// _ = "end of CoverTab[4183]"
}

//line /usr/local/go/src/math/big/decimal.go:270
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/decimal.go:270
var _ = _go_fuzz_dep_.CoverTab
