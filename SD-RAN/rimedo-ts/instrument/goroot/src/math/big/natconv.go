// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements nat-to-string conversion functions.

//line /usr/local/go/src/math/big/natconv.go:7
package big

//line /usr/local/go/src/math/big/natconv.go:7
import (
//line /usr/local/go/src/math/big/natconv.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/natconv.go:7
)
//line /usr/local/go/src/math/big/natconv.go:7
import (
//line /usr/local/go/src/math/big/natconv.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/natconv.go:7
)

import (
	"errors"
	"fmt"
	"io"
	"math"
	"math/bits"
	"sync"
)

const digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//line /usr/local/go/src/math/big/natconv.go:23
// MaxBase is the largest number base accepted for string conversions.
const MaxBase = 10 + ('z' - 'a' + 1) + ('Z' - 'A' + 1)
const maxBaseSmall = 10 + ('z' - 'a' + 1)

// maxPow returns (b**n, n) such that b**n is the largest power b**n <= _M.
//line /usr/local/go/src/math/big/natconv.go:27
// For instance maxPow(10) == (1e19, 19) for 19 decimal digits in a 64bit Word.
//line /usr/local/go/src/math/big/natconv.go:27
// In other words, at most n digits in base b fit into a Word.
//line /usr/local/go/src/math/big/natconv.go:27
// TODO(gri) replace this with a table, generated at build time.
//line /usr/local/go/src/math/big/natconv.go:31
func maxPow(b Word) (p Word, n int) {
//line /usr/local/go/src/math/big/natconv.go:31
	_go_fuzz_dep_.CoverTab[6157]++
							p, n = b, 1
							for max := _M / b; p <= max; {
//line /usr/local/go/src/math/big/natconv.go:33
		_go_fuzz_dep_.CoverTab[6159]++

								p *= b
								n++
//line /usr/local/go/src/math/big/natconv.go:36
		// _ = "end of CoverTab[6159]"
	}
//line /usr/local/go/src/math/big/natconv.go:37
	// _ = "end of CoverTab[6157]"
//line /usr/local/go/src/math/big/natconv.go:37
	_go_fuzz_dep_.CoverTab[6158]++

							return
//line /usr/local/go/src/math/big/natconv.go:39
	// _ = "end of CoverTab[6158]"
}

// pow returns x**n for n > 0, and 1 otherwise.
func pow(x Word, n int) (p Word) {
//line /usr/local/go/src/math/big/natconv.go:43
	_go_fuzz_dep_.CoverTab[6160]++

//line /usr/local/go/src/math/big/natconv.go:47
	p = 1
	for n > 0 {
//line /usr/local/go/src/math/big/natconv.go:48
		_go_fuzz_dep_.CoverTab[6162]++
								if n&1 != 0 {
//line /usr/local/go/src/math/big/natconv.go:49
			_go_fuzz_dep_.CoverTab[6164]++
									p *= x
//line /usr/local/go/src/math/big/natconv.go:50
			// _ = "end of CoverTab[6164]"
		} else {
//line /usr/local/go/src/math/big/natconv.go:51
			_go_fuzz_dep_.CoverTab[6165]++
//line /usr/local/go/src/math/big/natconv.go:51
			// _ = "end of CoverTab[6165]"
//line /usr/local/go/src/math/big/natconv.go:51
		}
//line /usr/local/go/src/math/big/natconv.go:51
		// _ = "end of CoverTab[6162]"
//line /usr/local/go/src/math/big/natconv.go:51
		_go_fuzz_dep_.CoverTab[6163]++
								x *= x
								n >>= 1
//line /usr/local/go/src/math/big/natconv.go:53
		// _ = "end of CoverTab[6163]"
	}
//line /usr/local/go/src/math/big/natconv.go:54
	// _ = "end of CoverTab[6160]"
//line /usr/local/go/src/math/big/natconv.go:54
	_go_fuzz_dep_.CoverTab[6161]++
							return
//line /usr/local/go/src/math/big/natconv.go:55
	// _ = "end of CoverTab[6161]"
}

// scan errors
var (
	errNoDigits	= errors.New("number has no digits")
	errInvalSep	= errors.New("'_' must separate successive digits")
)

// scan scans the number corresponding to the longest possible prefix
//line /usr/local/go/src/math/big/natconv.go:64
// from r representing an unsigned number in a given conversion base.
//line /usr/local/go/src/math/big/natconv.go:64
// scan returns the corresponding natural number res, the actual base b,
//line /usr/local/go/src/math/big/natconv.go:64
// a digit count, and a read or syntax error err, if any.
//line /usr/local/go/src/math/big/natconv.go:64
//
//line /usr/local/go/src/math/big/natconv.go:64
// For base 0, an underscore character “_” may appear between a base
//line /usr/local/go/src/math/big/natconv.go:64
// prefix and an adjacent digit, and between successive digits; such
//line /usr/local/go/src/math/big/natconv.go:64
// underscores do not change the value of the number, or the returned
//line /usr/local/go/src/math/big/natconv.go:64
// digit count. Incorrect placement of underscores is reported as an
//line /usr/local/go/src/math/big/natconv.go:64
// error if there are no other errors. If base != 0, underscores are
//line /usr/local/go/src/math/big/natconv.go:64
// not recognized and thus terminate scanning like any other character
//line /usr/local/go/src/math/big/natconv.go:64
// that is not a valid radix point or digit.
//line /usr/local/go/src/math/big/natconv.go:64
//
//line /usr/local/go/src/math/big/natconv.go:64
//	number    = mantissa | prefix pmantissa .
//line /usr/local/go/src/math/big/natconv.go:64
//	prefix    = "0" [ "b" | "B" | "o" | "O" | "x" | "X" ] .
//line /usr/local/go/src/math/big/natconv.go:64
//	mantissa  = digits "." [ digits ] | digits | "." digits .
//line /usr/local/go/src/math/big/natconv.go:64
//	pmantissa = [ "_" ] digits "." [ digits ] | [ "_" ] digits | "." digits .
//line /usr/local/go/src/math/big/natconv.go:64
//	digits    = digit { [ "_" ] digit } .
//line /usr/local/go/src/math/big/natconv.go:64
//	digit     = "0" ... "9" | "a" ... "z" | "A" ... "Z" .
//line /usr/local/go/src/math/big/natconv.go:64
//
//line /usr/local/go/src/math/big/natconv.go:64
// Unless fracOk is set, the base argument must be 0 or a value between
//line /usr/local/go/src/math/big/natconv.go:64
// 2 and MaxBase. If fracOk is set, the base argument must be one of
//line /usr/local/go/src/math/big/natconv.go:64
// 0, 2, 8, 10, or 16. Providing an invalid base argument leads to a run-
//line /usr/local/go/src/math/big/natconv.go:64
// time panic.
//line /usr/local/go/src/math/big/natconv.go:64
//
//line /usr/local/go/src/math/big/natconv.go:64
// For base 0, the number prefix determines the actual base: A prefix of
//line /usr/local/go/src/math/big/natconv.go:64
// “0b” or “0B” selects base 2, “0o” or “0O” selects base 8, and
//line /usr/local/go/src/math/big/natconv.go:64
// “0x” or “0X” selects base 16. If fracOk is false, a “0” prefix
//line /usr/local/go/src/math/big/natconv.go:64
// (immediately followed by digits) selects base 8 as well. Otherwise,
//line /usr/local/go/src/math/big/natconv.go:64
// the selected base is 10 and no prefix is accepted.
//line /usr/local/go/src/math/big/natconv.go:64
//
//line /usr/local/go/src/math/big/natconv.go:64
// If fracOk is set, a period followed by a fractional part is permitted.
//line /usr/local/go/src/math/big/natconv.go:64
// The result value is computed as if there were no period present; and
//line /usr/local/go/src/math/big/natconv.go:64
// the count value is used to determine the fractional part.
//line /usr/local/go/src/math/big/natconv.go:64
//
//line /usr/local/go/src/math/big/natconv.go:64
// For bases <= 36, lower and upper case letters are considered the same:
//line /usr/local/go/src/math/big/natconv.go:64
// The letters 'a' to 'z' and 'A' to 'Z' represent digit values 10 to 35.
//line /usr/local/go/src/math/big/natconv.go:64
// For bases > 36, the upper case letters 'A' to 'Z' represent the digit
//line /usr/local/go/src/math/big/natconv.go:64
// values 36 to 61.
//line /usr/local/go/src/math/big/natconv.go:64
//
//line /usr/local/go/src/math/big/natconv.go:64
// A result digit count > 0 corresponds to the number of (non-prefix) digits
//line /usr/local/go/src/math/big/natconv.go:64
// parsed. A digit count <= 0 indicates the presence of a period (if fracOk
//line /usr/local/go/src/math/big/natconv.go:64
// is set, only), and -count is the number of fractional digits found.
//line /usr/local/go/src/math/big/natconv.go:64
// In this case, the actual value of the scanned number is res * b**count.
//line /usr/local/go/src/math/big/natconv.go:108
func (z nat) scan(r io.ByteScanner, base int, fracOk bool) (res nat, b, count int, err error) {
//line /usr/local/go/src/math/big/natconv.go:108
	_go_fuzz_dep_.CoverTab[6166]++

							baseOk := base == 0 || func() bool {
//line /usr/local/go/src/math/big/natconv.go:110
		_go_fuzz_dep_.CoverTab[6175]++
//line /usr/local/go/src/math/big/natconv.go:110
		return !fracOk && func() bool {
									_go_fuzz_dep_.CoverTab[6176]++
//line /usr/local/go/src/math/big/natconv.go:111
			return 2 <= base
//line /usr/local/go/src/math/big/natconv.go:111
			// _ = "end of CoverTab[6176]"
//line /usr/local/go/src/math/big/natconv.go:111
		}() && func() bool {
//line /usr/local/go/src/math/big/natconv.go:111
			_go_fuzz_dep_.CoverTab[6177]++
//line /usr/local/go/src/math/big/natconv.go:111
			return base <= MaxBase
//line /usr/local/go/src/math/big/natconv.go:111
			// _ = "end of CoverTab[6177]"
//line /usr/local/go/src/math/big/natconv.go:111
		}()
//line /usr/local/go/src/math/big/natconv.go:111
		// _ = "end of CoverTab[6175]"
//line /usr/local/go/src/math/big/natconv.go:111
	}() || func() bool {
//line /usr/local/go/src/math/big/natconv.go:111
		_go_fuzz_dep_.CoverTab[6178]++
//line /usr/local/go/src/math/big/natconv.go:111
		return fracOk && func() bool {
									_go_fuzz_dep_.CoverTab[6179]++
//line /usr/local/go/src/math/big/natconv.go:112
			return (base == 2 || func() bool {
//line /usr/local/go/src/math/big/natconv.go:112
				_go_fuzz_dep_.CoverTab[6180]++
//line /usr/local/go/src/math/big/natconv.go:112
				return base == 8
//line /usr/local/go/src/math/big/natconv.go:112
				// _ = "end of CoverTab[6180]"
//line /usr/local/go/src/math/big/natconv.go:112
			}() || func() bool {
//line /usr/local/go/src/math/big/natconv.go:112
				_go_fuzz_dep_.CoverTab[6181]++
//line /usr/local/go/src/math/big/natconv.go:112
				return base == 10
//line /usr/local/go/src/math/big/natconv.go:112
				// _ = "end of CoverTab[6181]"
//line /usr/local/go/src/math/big/natconv.go:112
			}() || func() bool {
//line /usr/local/go/src/math/big/natconv.go:112
				_go_fuzz_dep_.CoverTab[6182]++
//line /usr/local/go/src/math/big/natconv.go:112
				return base == 16
//line /usr/local/go/src/math/big/natconv.go:112
				// _ = "end of CoverTab[6182]"
//line /usr/local/go/src/math/big/natconv.go:112
			}())
//line /usr/local/go/src/math/big/natconv.go:112
			// _ = "end of CoverTab[6179]"
//line /usr/local/go/src/math/big/natconv.go:112
		}()
//line /usr/local/go/src/math/big/natconv.go:112
		// _ = "end of CoverTab[6178]"
//line /usr/local/go/src/math/big/natconv.go:112
	}()
	if !baseOk {
//line /usr/local/go/src/math/big/natconv.go:113
		_go_fuzz_dep_.CoverTab[6183]++
								panic(fmt.Sprintf("invalid number base %d", base))
//line /usr/local/go/src/math/big/natconv.go:114
		// _ = "end of CoverTab[6183]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:115
		_go_fuzz_dep_.CoverTab[6184]++
//line /usr/local/go/src/math/big/natconv.go:115
		// _ = "end of CoverTab[6184]"
//line /usr/local/go/src/math/big/natconv.go:115
	}
//line /usr/local/go/src/math/big/natconv.go:115
	// _ = "end of CoverTab[6166]"
//line /usr/local/go/src/math/big/natconv.go:115
	_go_fuzz_dep_.CoverTab[6167]++

//line /usr/local/go/src/math/big/natconv.go:121
	prev := '.'
							invalSep := false

//line /usr/local/go/src/math/big/natconv.go:125
	ch, err := r.ReadByte()

//line /usr/local/go/src/math/big/natconv.go:128
	b, prefix := base, 0
	if base == 0 {
//line /usr/local/go/src/math/big/natconv.go:129
		_go_fuzz_dep_.CoverTab[6185]++

								b = 10
								if err == nil && func() bool {
//line /usr/local/go/src/math/big/natconv.go:132
			_go_fuzz_dep_.CoverTab[6186]++
//line /usr/local/go/src/math/big/natconv.go:132
			return ch == '0'
//line /usr/local/go/src/math/big/natconv.go:132
			// _ = "end of CoverTab[6186]"
//line /usr/local/go/src/math/big/natconv.go:132
		}() {
//line /usr/local/go/src/math/big/natconv.go:132
			_go_fuzz_dep_.CoverTab[6187]++
									prev = '0'
									count = 1
									ch, err = r.ReadByte()
									if err == nil {
//line /usr/local/go/src/math/big/natconv.go:136
				_go_fuzz_dep_.CoverTab[6188]++

										switch ch {
				case 'b', 'B':
//line /usr/local/go/src/math/big/natconv.go:139
					_go_fuzz_dep_.CoverTab[6190]++
											b, prefix = 2, 'b'
//line /usr/local/go/src/math/big/natconv.go:140
					// _ = "end of CoverTab[6190]"
				case 'o', 'O':
//line /usr/local/go/src/math/big/natconv.go:141
					_go_fuzz_dep_.CoverTab[6191]++
											b, prefix = 8, 'o'
//line /usr/local/go/src/math/big/natconv.go:142
					// _ = "end of CoverTab[6191]"
				case 'x', 'X':
//line /usr/local/go/src/math/big/natconv.go:143
					_go_fuzz_dep_.CoverTab[6192]++
											b, prefix = 16, 'x'
//line /usr/local/go/src/math/big/natconv.go:144
					// _ = "end of CoverTab[6192]"
				default:
//line /usr/local/go/src/math/big/natconv.go:145
					_go_fuzz_dep_.CoverTab[6193]++
											if !fracOk {
//line /usr/local/go/src/math/big/natconv.go:146
						_go_fuzz_dep_.CoverTab[6194]++
												b, prefix = 8, '0'
//line /usr/local/go/src/math/big/natconv.go:147
						// _ = "end of CoverTab[6194]"
					} else {
//line /usr/local/go/src/math/big/natconv.go:148
						_go_fuzz_dep_.CoverTab[6195]++
//line /usr/local/go/src/math/big/natconv.go:148
						// _ = "end of CoverTab[6195]"
//line /usr/local/go/src/math/big/natconv.go:148
					}
//line /usr/local/go/src/math/big/natconv.go:148
					// _ = "end of CoverTab[6193]"
				}
//line /usr/local/go/src/math/big/natconv.go:149
				// _ = "end of CoverTab[6188]"
//line /usr/local/go/src/math/big/natconv.go:149
				_go_fuzz_dep_.CoverTab[6189]++
										if prefix != 0 {
//line /usr/local/go/src/math/big/natconv.go:150
					_go_fuzz_dep_.CoverTab[6196]++
											count = 0
											if prefix != '0' {
//line /usr/local/go/src/math/big/natconv.go:152
						_go_fuzz_dep_.CoverTab[6197]++
												ch, err = r.ReadByte()
//line /usr/local/go/src/math/big/natconv.go:153
						// _ = "end of CoverTab[6197]"
					} else {
//line /usr/local/go/src/math/big/natconv.go:154
						_go_fuzz_dep_.CoverTab[6198]++
//line /usr/local/go/src/math/big/natconv.go:154
						// _ = "end of CoverTab[6198]"
//line /usr/local/go/src/math/big/natconv.go:154
					}
//line /usr/local/go/src/math/big/natconv.go:154
					// _ = "end of CoverTab[6196]"
				} else {
//line /usr/local/go/src/math/big/natconv.go:155
					_go_fuzz_dep_.CoverTab[6199]++
//line /usr/local/go/src/math/big/natconv.go:155
					// _ = "end of CoverTab[6199]"
//line /usr/local/go/src/math/big/natconv.go:155
				}
//line /usr/local/go/src/math/big/natconv.go:155
				// _ = "end of CoverTab[6189]"
			} else {
//line /usr/local/go/src/math/big/natconv.go:156
				_go_fuzz_dep_.CoverTab[6200]++
//line /usr/local/go/src/math/big/natconv.go:156
				// _ = "end of CoverTab[6200]"
//line /usr/local/go/src/math/big/natconv.go:156
			}
//line /usr/local/go/src/math/big/natconv.go:156
			// _ = "end of CoverTab[6187]"
		} else {
//line /usr/local/go/src/math/big/natconv.go:157
			_go_fuzz_dep_.CoverTab[6201]++
//line /usr/local/go/src/math/big/natconv.go:157
			// _ = "end of CoverTab[6201]"
//line /usr/local/go/src/math/big/natconv.go:157
		}
//line /usr/local/go/src/math/big/natconv.go:157
		// _ = "end of CoverTab[6185]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:158
		_go_fuzz_dep_.CoverTab[6202]++
//line /usr/local/go/src/math/big/natconv.go:158
		// _ = "end of CoverTab[6202]"
//line /usr/local/go/src/math/big/natconv.go:158
	}
//line /usr/local/go/src/math/big/natconv.go:158
	// _ = "end of CoverTab[6167]"
//line /usr/local/go/src/math/big/natconv.go:158
	_go_fuzz_dep_.CoverTab[6168]++

//line /usr/local/go/src/math/big/natconv.go:164
	z = z[:0]
	b1 := Word(b)
	bn, n := maxPow(b1)
	di := Word(0)
	i := 0
	dp := -1
	for err == nil {
//line /usr/local/go/src/math/big/natconv.go:170
		_go_fuzz_dep_.CoverTab[6203]++
								if ch == '.' && func() bool {
//line /usr/local/go/src/math/big/natconv.go:171
			_go_fuzz_dep_.CoverTab[6205]++
//line /usr/local/go/src/math/big/natconv.go:171
			return fracOk
//line /usr/local/go/src/math/big/natconv.go:171
			// _ = "end of CoverTab[6205]"
//line /usr/local/go/src/math/big/natconv.go:171
		}() {
//line /usr/local/go/src/math/big/natconv.go:171
			_go_fuzz_dep_.CoverTab[6206]++
									fracOk = false
									if prev == '_' {
//line /usr/local/go/src/math/big/natconv.go:173
				_go_fuzz_dep_.CoverTab[6208]++
										invalSep = true
//line /usr/local/go/src/math/big/natconv.go:174
				// _ = "end of CoverTab[6208]"
			} else {
//line /usr/local/go/src/math/big/natconv.go:175
				_go_fuzz_dep_.CoverTab[6209]++
//line /usr/local/go/src/math/big/natconv.go:175
				// _ = "end of CoverTab[6209]"
//line /usr/local/go/src/math/big/natconv.go:175
			}
//line /usr/local/go/src/math/big/natconv.go:175
			// _ = "end of CoverTab[6206]"
//line /usr/local/go/src/math/big/natconv.go:175
			_go_fuzz_dep_.CoverTab[6207]++
									prev = '.'
									dp = count
//line /usr/local/go/src/math/big/natconv.go:177
			// _ = "end of CoverTab[6207]"
		} else {
//line /usr/local/go/src/math/big/natconv.go:178
			_go_fuzz_dep_.CoverTab[6210]++
//line /usr/local/go/src/math/big/natconv.go:178
			if ch == '_' && func() bool {
//line /usr/local/go/src/math/big/natconv.go:178
				_go_fuzz_dep_.CoverTab[6211]++
//line /usr/local/go/src/math/big/natconv.go:178
				return base == 0
//line /usr/local/go/src/math/big/natconv.go:178
				// _ = "end of CoverTab[6211]"
//line /usr/local/go/src/math/big/natconv.go:178
			}() {
//line /usr/local/go/src/math/big/natconv.go:178
				_go_fuzz_dep_.CoverTab[6212]++
										if prev != '0' {
//line /usr/local/go/src/math/big/natconv.go:179
					_go_fuzz_dep_.CoverTab[6214]++
											invalSep = true
//line /usr/local/go/src/math/big/natconv.go:180
					// _ = "end of CoverTab[6214]"
				} else {
//line /usr/local/go/src/math/big/natconv.go:181
					_go_fuzz_dep_.CoverTab[6215]++
//line /usr/local/go/src/math/big/natconv.go:181
					// _ = "end of CoverTab[6215]"
//line /usr/local/go/src/math/big/natconv.go:181
				}
//line /usr/local/go/src/math/big/natconv.go:181
				// _ = "end of CoverTab[6212]"
//line /usr/local/go/src/math/big/natconv.go:181
				_go_fuzz_dep_.CoverTab[6213]++
										prev = '_'
//line /usr/local/go/src/math/big/natconv.go:182
				// _ = "end of CoverTab[6213]"
			} else {
//line /usr/local/go/src/math/big/natconv.go:183
				_go_fuzz_dep_.CoverTab[6216]++
				// convert rune into digit value d1
				var d1 Word
				switch {
				case '0' <= ch && func() bool {
//line /usr/local/go/src/math/big/natconv.go:187
					_go_fuzz_dep_.CoverTab[6223]++
//line /usr/local/go/src/math/big/natconv.go:187
					return ch <= '9'
//line /usr/local/go/src/math/big/natconv.go:187
					// _ = "end of CoverTab[6223]"
//line /usr/local/go/src/math/big/natconv.go:187
				}():
//line /usr/local/go/src/math/big/natconv.go:187
					_go_fuzz_dep_.CoverTab[6219]++
											d1 = Word(ch - '0')
//line /usr/local/go/src/math/big/natconv.go:188
					// _ = "end of CoverTab[6219]"
				case 'a' <= ch && func() bool {
//line /usr/local/go/src/math/big/natconv.go:189
					_go_fuzz_dep_.CoverTab[6224]++
//line /usr/local/go/src/math/big/natconv.go:189
					return ch <= 'z'
//line /usr/local/go/src/math/big/natconv.go:189
					// _ = "end of CoverTab[6224]"
//line /usr/local/go/src/math/big/natconv.go:189
				}():
//line /usr/local/go/src/math/big/natconv.go:189
					_go_fuzz_dep_.CoverTab[6220]++
											d1 = Word(ch - 'a' + 10)
//line /usr/local/go/src/math/big/natconv.go:190
					// _ = "end of CoverTab[6220]"
				case 'A' <= ch && func() bool {
//line /usr/local/go/src/math/big/natconv.go:191
					_go_fuzz_dep_.CoverTab[6225]++
//line /usr/local/go/src/math/big/natconv.go:191
					return ch <= 'Z'
//line /usr/local/go/src/math/big/natconv.go:191
					// _ = "end of CoverTab[6225]"
//line /usr/local/go/src/math/big/natconv.go:191
				}():
//line /usr/local/go/src/math/big/natconv.go:191
					_go_fuzz_dep_.CoverTab[6221]++
											if b <= maxBaseSmall {
//line /usr/local/go/src/math/big/natconv.go:192
						_go_fuzz_dep_.CoverTab[6226]++
												d1 = Word(ch - 'A' + 10)
//line /usr/local/go/src/math/big/natconv.go:193
						// _ = "end of CoverTab[6226]"
					} else {
//line /usr/local/go/src/math/big/natconv.go:194
						_go_fuzz_dep_.CoverTab[6227]++
												d1 = Word(ch - 'A' + maxBaseSmall)
//line /usr/local/go/src/math/big/natconv.go:195
						// _ = "end of CoverTab[6227]"
					}
//line /usr/local/go/src/math/big/natconv.go:196
					// _ = "end of CoverTab[6221]"
				default:
//line /usr/local/go/src/math/big/natconv.go:197
					_go_fuzz_dep_.CoverTab[6222]++
											d1 = MaxBase + 1
//line /usr/local/go/src/math/big/natconv.go:198
					// _ = "end of CoverTab[6222]"
				}
//line /usr/local/go/src/math/big/natconv.go:199
				// _ = "end of CoverTab[6216]"
//line /usr/local/go/src/math/big/natconv.go:199
				_go_fuzz_dep_.CoverTab[6217]++
										if d1 >= b1 {
//line /usr/local/go/src/math/big/natconv.go:200
					_go_fuzz_dep_.CoverTab[6228]++
											r.UnreadByte()
											break
//line /usr/local/go/src/math/big/natconv.go:202
					// _ = "end of CoverTab[6228]"
				} else {
//line /usr/local/go/src/math/big/natconv.go:203
					_go_fuzz_dep_.CoverTab[6229]++
//line /usr/local/go/src/math/big/natconv.go:203
					// _ = "end of CoverTab[6229]"
//line /usr/local/go/src/math/big/natconv.go:203
				}
//line /usr/local/go/src/math/big/natconv.go:203
				// _ = "end of CoverTab[6217]"
//line /usr/local/go/src/math/big/natconv.go:203
				_go_fuzz_dep_.CoverTab[6218]++
										prev = '0'
										count++

//line /usr/local/go/src/math/big/natconv.go:208
				di = di*b1 + d1
										i++

//line /usr/local/go/src/math/big/natconv.go:212
				if i == n {
//line /usr/local/go/src/math/big/natconv.go:212
					_go_fuzz_dep_.CoverTab[6230]++
											z = z.mulAddWW(z, bn, di)
											di = 0
											i = 0
//line /usr/local/go/src/math/big/natconv.go:215
					// _ = "end of CoverTab[6230]"
				} else {
//line /usr/local/go/src/math/big/natconv.go:216
					_go_fuzz_dep_.CoverTab[6231]++
//line /usr/local/go/src/math/big/natconv.go:216
					// _ = "end of CoverTab[6231]"
//line /usr/local/go/src/math/big/natconv.go:216
				}
//line /usr/local/go/src/math/big/natconv.go:216
				// _ = "end of CoverTab[6218]"
			}
//line /usr/local/go/src/math/big/natconv.go:217
			// _ = "end of CoverTab[6210]"
//line /usr/local/go/src/math/big/natconv.go:217
		}
//line /usr/local/go/src/math/big/natconv.go:217
		// _ = "end of CoverTab[6203]"
//line /usr/local/go/src/math/big/natconv.go:217
		_go_fuzz_dep_.CoverTab[6204]++

								ch, err = r.ReadByte()
//line /usr/local/go/src/math/big/natconv.go:219
		// _ = "end of CoverTab[6204]"
	}
//line /usr/local/go/src/math/big/natconv.go:220
	// _ = "end of CoverTab[6168]"
//line /usr/local/go/src/math/big/natconv.go:220
	_go_fuzz_dep_.CoverTab[6169]++

							if err == io.EOF {
//line /usr/local/go/src/math/big/natconv.go:222
		_go_fuzz_dep_.CoverTab[6232]++
								err = nil
//line /usr/local/go/src/math/big/natconv.go:223
		// _ = "end of CoverTab[6232]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:224
		_go_fuzz_dep_.CoverTab[6233]++
//line /usr/local/go/src/math/big/natconv.go:224
		// _ = "end of CoverTab[6233]"
//line /usr/local/go/src/math/big/natconv.go:224
	}
//line /usr/local/go/src/math/big/natconv.go:224
	// _ = "end of CoverTab[6169]"
//line /usr/local/go/src/math/big/natconv.go:224
	_go_fuzz_dep_.CoverTab[6170]++

//line /usr/local/go/src/math/big/natconv.go:227
	if err == nil && func() bool {
//line /usr/local/go/src/math/big/natconv.go:227
		_go_fuzz_dep_.CoverTab[6234]++
//line /usr/local/go/src/math/big/natconv.go:227
		return (invalSep || func() bool {
//line /usr/local/go/src/math/big/natconv.go:227
			_go_fuzz_dep_.CoverTab[6235]++
//line /usr/local/go/src/math/big/natconv.go:227
			return prev == '_'
//line /usr/local/go/src/math/big/natconv.go:227
			// _ = "end of CoverTab[6235]"
//line /usr/local/go/src/math/big/natconv.go:227
		}())
//line /usr/local/go/src/math/big/natconv.go:227
		// _ = "end of CoverTab[6234]"
//line /usr/local/go/src/math/big/natconv.go:227
	}() {
//line /usr/local/go/src/math/big/natconv.go:227
		_go_fuzz_dep_.CoverTab[6236]++
								err = errInvalSep
//line /usr/local/go/src/math/big/natconv.go:228
		// _ = "end of CoverTab[6236]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:229
		_go_fuzz_dep_.CoverTab[6237]++
//line /usr/local/go/src/math/big/natconv.go:229
		// _ = "end of CoverTab[6237]"
//line /usr/local/go/src/math/big/natconv.go:229
	}
//line /usr/local/go/src/math/big/natconv.go:229
	// _ = "end of CoverTab[6170]"
//line /usr/local/go/src/math/big/natconv.go:229
	_go_fuzz_dep_.CoverTab[6171]++

							if count == 0 {
//line /usr/local/go/src/math/big/natconv.go:231
		_go_fuzz_dep_.CoverTab[6238]++

								if prefix == '0' {
//line /usr/local/go/src/math/big/natconv.go:233
			_go_fuzz_dep_.CoverTab[6240]++

//line /usr/local/go/src/math/big/natconv.go:236
			return z[:0], 10, 1, err
//line /usr/local/go/src/math/big/natconv.go:236
			// _ = "end of CoverTab[6240]"
		} else {
//line /usr/local/go/src/math/big/natconv.go:237
			_go_fuzz_dep_.CoverTab[6241]++
//line /usr/local/go/src/math/big/natconv.go:237
			// _ = "end of CoverTab[6241]"
//line /usr/local/go/src/math/big/natconv.go:237
		}
//line /usr/local/go/src/math/big/natconv.go:237
		// _ = "end of CoverTab[6238]"
//line /usr/local/go/src/math/big/natconv.go:237
		_go_fuzz_dep_.CoverTab[6239]++
								err = errNoDigits
//line /usr/local/go/src/math/big/natconv.go:238
		// _ = "end of CoverTab[6239]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:239
		_go_fuzz_dep_.CoverTab[6242]++
//line /usr/local/go/src/math/big/natconv.go:239
		// _ = "end of CoverTab[6242]"
//line /usr/local/go/src/math/big/natconv.go:239
	}
//line /usr/local/go/src/math/big/natconv.go:239
	// _ = "end of CoverTab[6171]"
//line /usr/local/go/src/math/big/natconv.go:239
	_go_fuzz_dep_.CoverTab[6172]++

//line /usr/local/go/src/math/big/natconv.go:242
	if i > 0 {
//line /usr/local/go/src/math/big/natconv.go:242
		_go_fuzz_dep_.CoverTab[6243]++
								z = z.mulAddWW(z, pow(b1, i), di)
//line /usr/local/go/src/math/big/natconv.go:243
		// _ = "end of CoverTab[6243]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:244
		_go_fuzz_dep_.CoverTab[6244]++
//line /usr/local/go/src/math/big/natconv.go:244
		// _ = "end of CoverTab[6244]"
//line /usr/local/go/src/math/big/natconv.go:244
	}
//line /usr/local/go/src/math/big/natconv.go:244
	// _ = "end of CoverTab[6172]"
//line /usr/local/go/src/math/big/natconv.go:244
	_go_fuzz_dep_.CoverTab[6173]++
							res = z.norm()

//line /usr/local/go/src/math/big/natconv.go:248
	if dp >= 0 {
//line /usr/local/go/src/math/big/natconv.go:248
		_go_fuzz_dep_.CoverTab[6245]++

								count = dp - count
//line /usr/local/go/src/math/big/natconv.go:250
		// _ = "end of CoverTab[6245]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:251
		_go_fuzz_dep_.CoverTab[6246]++
//line /usr/local/go/src/math/big/natconv.go:251
		// _ = "end of CoverTab[6246]"
//line /usr/local/go/src/math/big/natconv.go:251
	}
//line /usr/local/go/src/math/big/natconv.go:251
	// _ = "end of CoverTab[6173]"
//line /usr/local/go/src/math/big/natconv.go:251
	_go_fuzz_dep_.CoverTab[6174]++

							return
//line /usr/local/go/src/math/big/natconv.go:253
	// _ = "end of CoverTab[6174]"
}

// utoa converts x to an ASCII representation in the given base;
//line /usr/local/go/src/math/big/natconv.go:256
// base must be between 2 and MaxBase, inclusive.
//line /usr/local/go/src/math/big/natconv.go:258
func (x nat) utoa(base int) []byte {
//line /usr/local/go/src/math/big/natconv.go:258
	_go_fuzz_dep_.CoverTab[6247]++
							return x.itoa(false, base)
//line /usr/local/go/src/math/big/natconv.go:259
	// _ = "end of CoverTab[6247]"
}

// itoa is like utoa but it prepends a '-' if neg && x != 0.
func (x nat) itoa(neg bool, base int) []byte {
//line /usr/local/go/src/math/big/natconv.go:263
	_go_fuzz_dep_.CoverTab[6248]++
							if base < 2 || func() bool {
//line /usr/local/go/src/math/big/natconv.go:264
		_go_fuzz_dep_.CoverTab[6254]++
//line /usr/local/go/src/math/big/natconv.go:264
		return base > MaxBase
//line /usr/local/go/src/math/big/natconv.go:264
		// _ = "end of CoverTab[6254]"
//line /usr/local/go/src/math/big/natconv.go:264
	}() {
//line /usr/local/go/src/math/big/natconv.go:264
		_go_fuzz_dep_.CoverTab[6255]++
								panic("invalid base")
//line /usr/local/go/src/math/big/natconv.go:265
		// _ = "end of CoverTab[6255]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:266
		_go_fuzz_dep_.CoverTab[6256]++
//line /usr/local/go/src/math/big/natconv.go:266
		// _ = "end of CoverTab[6256]"
//line /usr/local/go/src/math/big/natconv.go:266
	}
//line /usr/local/go/src/math/big/natconv.go:266
	// _ = "end of CoverTab[6248]"
//line /usr/local/go/src/math/big/natconv.go:266
	_go_fuzz_dep_.CoverTab[6249]++

//line /usr/local/go/src/math/big/natconv.go:269
	if len(x) == 0 {
//line /usr/local/go/src/math/big/natconv.go:269
		_go_fuzz_dep_.CoverTab[6257]++
								return []byte("0")
//line /usr/local/go/src/math/big/natconv.go:270
		// _ = "end of CoverTab[6257]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:271
		_go_fuzz_dep_.CoverTab[6258]++
//line /usr/local/go/src/math/big/natconv.go:271
		// _ = "end of CoverTab[6258]"
//line /usr/local/go/src/math/big/natconv.go:271
	}
//line /usr/local/go/src/math/big/natconv.go:271
	// _ = "end of CoverTab[6249]"
//line /usr/local/go/src/math/big/natconv.go:271
	_go_fuzz_dep_.CoverTab[6250]++

//line /usr/local/go/src/math/big/natconv.go:275
	i := int(float64(x.bitLen())/math.Log2(float64(base))) + 1
	if neg {
//line /usr/local/go/src/math/big/natconv.go:276
		_go_fuzz_dep_.CoverTab[6259]++
								i++
//line /usr/local/go/src/math/big/natconv.go:277
		// _ = "end of CoverTab[6259]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:278
		_go_fuzz_dep_.CoverTab[6260]++
//line /usr/local/go/src/math/big/natconv.go:278
		// _ = "end of CoverTab[6260]"
//line /usr/local/go/src/math/big/natconv.go:278
	}
//line /usr/local/go/src/math/big/natconv.go:278
	// _ = "end of CoverTab[6250]"
//line /usr/local/go/src/math/big/natconv.go:278
	_go_fuzz_dep_.CoverTab[6251]++
							s := make([]byte, i)

//line /usr/local/go/src/math/big/natconv.go:282
	if b := Word(base); b == b&-b {
//line /usr/local/go/src/math/big/natconv.go:282
		_go_fuzz_dep_.CoverTab[6261]++

								shift := uint(bits.TrailingZeros(uint(b)))
								mask := Word(1<<shift - 1)
								w := x[0]
								nbits := uint(_W)

//line /usr/local/go/src/math/big/natconv.go:290
		for k := 1; k < len(x); k++ {
//line /usr/local/go/src/math/big/natconv.go:290
			_go_fuzz_dep_.CoverTab[6263]++

									for nbits >= shift {
//line /usr/local/go/src/math/big/natconv.go:292
				_go_fuzz_dep_.CoverTab[6265]++
										i--
										s[i] = digits[w&mask]
										w >>= shift
										nbits -= shift
//line /usr/local/go/src/math/big/natconv.go:296
				// _ = "end of CoverTab[6265]"
			}
//line /usr/local/go/src/math/big/natconv.go:297
			// _ = "end of CoverTab[6263]"
//line /usr/local/go/src/math/big/natconv.go:297
			_go_fuzz_dep_.CoverTab[6264]++

//line /usr/local/go/src/math/big/natconv.go:300
			if nbits == 0 {
//line /usr/local/go/src/math/big/natconv.go:300
				_go_fuzz_dep_.CoverTab[6266]++

										w = x[k]
										nbits = _W
//line /usr/local/go/src/math/big/natconv.go:303
				// _ = "end of CoverTab[6266]"
			} else {
//line /usr/local/go/src/math/big/natconv.go:304
				_go_fuzz_dep_.CoverTab[6267]++

										w |= x[k] << nbits
										i--
										s[i] = digits[w&mask]

//line /usr/local/go/src/math/big/natconv.go:311
				w = x[k] >> (shift - nbits)
										nbits = _W - (shift - nbits)
//line /usr/local/go/src/math/big/natconv.go:312
				// _ = "end of CoverTab[6267]"
			}
//line /usr/local/go/src/math/big/natconv.go:313
			// _ = "end of CoverTab[6264]"
		}
//line /usr/local/go/src/math/big/natconv.go:314
		// _ = "end of CoverTab[6261]"
//line /usr/local/go/src/math/big/natconv.go:314
		_go_fuzz_dep_.CoverTab[6262]++

//line /usr/local/go/src/math/big/natconv.go:317
		for w != 0 {
//line /usr/local/go/src/math/big/natconv.go:317
			_go_fuzz_dep_.CoverTab[6268]++
									i--
									s[i] = digits[w&mask]
									w >>= shift
//line /usr/local/go/src/math/big/natconv.go:320
			// _ = "end of CoverTab[6268]"
		}
//line /usr/local/go/src/math/big/natconv.go:321
		// _ = "end of CoverTab[6262]"

	} else {
//line /usr/local/go/src/math/big/natconv.go:323
		_go_fuzz_dep_.CoverTab[6269]++
								bb, ndigits := maxPow(b)

//line /usr/local/go/src/math/big/natconv.go:328
		table := divisors(len(x), b, ndigits, bb)

//line /usr/local/go/src/math/big/natconv.go:331
		q := nat(nil).set(x)

//line /usr/local/go/src/math/big/natconv.go:334
		q.convertWords(s, b, ndigits, bb, table)

//line /usr/local/go/src/math/big/natconv.go:339
		i = 0
		for s[i] == '0' {
//line /usr/local/go/src/math/big/natconv.go:340
			_go_fuzz_dep_.CoverTab[6270]++
									i++
//line /usr/local/go/src/math/big/natconv.go:341
			// _ = "end of CoverTab[6270]"
		}
//line /usr/local/go/src/math/big/natconv.go:342
		// _ = "end of CoverTab[6269]"
	}
//line /usr/local/go/src/math/big/natconv.go:343
	// _ = "end of CoverTab[6251]"
//line /usr/local/go/src/math/big/natconv.go:343
	_go_fuzz_dep_.CoverTab[6252]++

							if neg {
//line /usr/local/go/src/math/big/natconv.go:345
		_go_fuzz_dep_.CoverTab[6271]++
								i--
								s[i] = '-'
//line /usr/local/go/src/math/big/natconv.go:347
		// _ = "end of CoverTab[6271]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:348
		_go_fuzz_dep_.CoverTab[6272]++
//line /usr/local/go/src/math/big/natconv.go:348
		// _ = "end of CoverTab[6272]"
//line /usr/local/go/src/math/big/natconv.go:348
	}
//line /usr/local/go/src/math/big/natconv.go:348
	// _ = "end of CoverTab[6252]"
//line /usr/local/go/src/math/big/natconv.go:348
	_go_fuzz_dep_.CoverTab[6253]++

							return s[i:]
//line /usr/local/go/src/math/big/natconv.go:350
	// _ = "end of CoverTab[6253]"
}

// Convert words of q to base b digits in s. If q is large, it is recursively "split in half"
//line /usr/local/go/src/math/big/natconv.go:353
// by nat/nat division using tabulated divisors. Otherwise, it is converted iteratively using
//line /usr/local/go/src/math/big/natconv.go:353
// repeated nat/Word division.
//line /usr/local/go/src/math/big/natconv.go:353
//
//line /usr/local/go/src/math/big/natconv.go:353
// The iterative method processes n Words by n divW() calls, each of which visits every Word in the
//line /usr/local/go/src/math/big/natconv.go:353
// incrementally shortened q for a total of n + (n-1) + (n-2) ... + 2 + 1, or n(n+1)/2 divW()'s.
//line /usr/local/go/src/math/big/natconv.go:353
// Recursive conversion divides q by its approximate square root, yielding two parts, each half
//line /usr/local/go/src/math/big/natconv.go:353
// the size of q. Using the iterative method on both halves means 2 * (n/2)(n/2 + 1)/2 divW()'s
//line /usr/local/go/src/math/big/natconv.go:353
// plus the expensive long div(). Asymptotically, the ratio is favorable at 1/2 the divW()'s, and
//line /usr/local/go/src/math/big/natconv.go:353
// is made better by splitting the subblocks recursively. Best is to split blocks until one more
//line /usr/local/go/src/math/big/natconv.go:353
// split would take longer (because of the nat/nat div()) than the twice as many divW()'s of the
//line /usr/local/go/src/math/big/natconv.go:353
// iterative approach. This threshold is represented by leafSize. Benchmarking of leafSize in the
//line /usr/local/go/src/math/big/natconv.go:353
// range 2..64 shows that values of 8 and 16 work well, with a 4x speedup at medium lengths and
//line /usr/local/go/src/math/big/natconv.go:353
// ~30x for 20000 digits. Use nat_test.go's BenchmarkLeafSize tests to optimize leafSize for
//line /usr/local/go/src/math/big/natconv.go:353
// specific hardware.
//line /usr/local/go/src/math/big/natconv.go:368
func (q nat) convertWords(s []byte, b Word, ndigits int, bb Word, table []divisor) {
//line /usr/local/go/src/math/big/natconv.go:368
	_go_fuzz_dep_.CoverTab[6273]++

							if table != nil {
//line /usr/local/go/src/math/big/natconv.go:370
		_go_fuzz_dep_.CoverTab[6276]++
		// len(q) > leafSize > 0
		var r nat
		index := len(table) - 1
		for len(q) > leafSize {
//line /usr/local/go/src/math/big/natconv.go:374
			_go_fuzz_dep_.CoverTab[6277]++

									maxLength := q.bitLen()
									minLength := maxLength >> 1
									for index > 0 && func() bool {
//line /usr/local/go/src/math/big/natconv.go:378
				_go_fuzz_dep_.CoverTab[6280]++
//line /usr/local/go/src/math/big/natconv.go:378
				return table[index-1].nbits > minLength
//line /usr/local/go/src/math/big/natconv.go:378
				// _ = "end of CoverTab[6280]"
//line /usr/local/go/src/math/big/natconv.go:378
			}() {
//line /usr/local/go/src/math/big/natconv.go:378
				_go_fuzz_dep_.CoverTab[6281]++
										index--
//line /usr/local/go/src/math/big/natconv.go:379
				// _ = "end of CoverTab[6281]"
			}
//line /usr/local/go/src/math/big/natconv.go:380
			// _ = "end of CoverTab[6277]"
//line /usr/local/go/src/math/big/natconv.go:380
			_go_fuzz_dep_.CoverTab[6278]++
									if table[index].nbits >= maxLength && func() bool {
//line /usr/local/go/src/math/big/natconv.go:381
				_go_fuzz_dep_.CoverTab[6282]++
//line /usr/local/go/src/math/big/natconv.go:381
				return table[index].bbb.cmp(q) >= 0
//line /usr/local/go/src/math/big/natconv.go:381
				// _ = "end of CoverTab[6282]"
//line /usr/local/go/src/math/big/natconv.go:381
			}() {
//line /usr/local/go/src/math/big/natconv.go:381
				_go_fuzz_dep_.CoverTab[6283]++
										index--
										if index < 0 {
//line /usr/local/go/src/math/big/natconv.go:383
					_go_fuzz_dep_.CoverTab[6284]++
											panic("internal inconsistency")
//line /usr/local/go/src/math/big/natconv.go:384
					// _ = "end of CoverTab[6284]"
				} else {
//line /usr/local/go/src/math/big/natconv.go:385
					_go_fuzz_dep_.CoverTab[6285]++
//line /usr/local/go/src/math/big/natconv.go:385
					// _ = "end of CoverTab[6285]"
//line /usr/local/go/src/math/big/natconv.go:385
				}
//line /usr/local/go/src/math/big/natconv.go:385
				// _ = "end of CoverTab[6283]"
			} else {
//line /usr/local/go/src/math/big/natconv.go:386
				_go_fuzz_dep_.CoverTab[6286]++
//line /usr/local/go/src/math/big/natconv.go:386
				// _ = "end of CoverTab[6286]"
//line /usr/local/go/src/math/big/natconv.go:386
			}
//line /usr/local/go/src/math/big/natconv.go:386
			// _ = "end of CoverTab[6278]"
//line /usr/local/go/src/math/big/natconv.go:386
			_go_fuzz_dep_.CoverTab[6279]++

//line /usr/local/go/src/math/big/natconv.go:389
			q, r = q.div(r, q, table[index].bbb)

//line /usr/local/go/src/math/big/natconv.go:392
			h := len(s) - table[index].ndigits
									r.convertWords(s[h:], b, ndigits, bb, table[0:index])
									s = s[:h]
//line /usr/local/go/src/math/big/natconv.go:394
			// _ = "end of CoverTab[6279]"
		}
//line /usr/local/go/src/math/big/natconv.go:395
		// _ = "end of CoverTab[6276]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:396
		_go_fuzz_dep_.CoverTab[6287]++
//line /usr/local/go/src/math/big/natconv.go:396
		// _ = "end of CoverTab[6287]"
//line /usr/local/go/src/math/big/natconv.go:396
	}
//line /usr/local/go/src/math/big/natconv.go:396
	// _ = "end of CoverTab[6273]"
//line /usr/local/go/src/math/big/natconv.go:396
	_go_fuzz_dep_.CoverTab[6274]++

//line /usr/local/go/src/math/big/natconv.go:399
	i := len(s)
	var r Word
	if b == 10 {
//line /usr/local/go/src/math/big/natconv.go:401
		_go_fuzz_dep_.CoverTab[6288]++

								for len(q) > 0 {
//line /usr/local/go/src/math/big/natconv.go:403
			_go_fuzz_dep_.CoverTab[6289]++

									q, r = q.divW(q, bb)
									for j := 0; j < ndigits && func() bool {
//line /usr/local/go/src/math/big/natconv.go:406
				_go_fuzz_dep_.CoverTab[6290]++
//line /usr/local/go/src/math/big/natconv.go:406
				return i > 0
//line /usr/local/go/src/math/big/natconv.go:406
				// _ = "end of CoverTab[6290]"
//line /usr/local/go/src/math/big/natconv.go:406
			}(); j++ {
//line /usr/local/go/src/math/big/natconv.go:406
				_go_fuzz_dep_.CoverTab[6291]++
										i--

//line /usr/local/go/src/math/big/natconv.go:411
				t := r / 10
										s[i] = '0' + byte(r-t*10)
										r = t
//line /usr/local/go/src/math/big/natconv.go:413
				// _ = "end of CoverTab[6291]"
			}
//line /usr/local/go/src/math/big/natconv.go:414
			// _ = "end of CoverTab[6289]"
		}
//line /usr/local/go/src/math/big/natconv.go:415
		// _ = "end of CoverTab[6288]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:416
		_go_fuzz_dep_.CoverTab[6292]++
								for len(q) > 0 {
//line /usr/local/go/src/math/big/natconv.go:417
			_go_fuzz_dep_.CoverTab[6293]++

									q, r = q.divW(q, bb)
									for j := 0; j < ndigits && func() bool {
//line /usr/local/go/src/math/big/natconv.go:420
				_go_fuzz_dep_.CoverTab[6294]++
//line /usr/local/go/src/math/big/natconv.go:420
				return i > 0
//line /usr/local/go/src/math/big/natconv.go:420
				// _ = "end of CoverTab[6294]"
//line /usr/local/go/src/math/big/natconv.go:420
			}(); j++ {
//line /usr/local/go/src/math/big/natconv.go:420
				_go_fuzz_dep_.CoverTab[6295]++
										i--
										s[i] = digits[r%b]
										r /= b
//line /usr/local/go/src/math/big/natconv.go:423
				// _ = "end of CoverTab[6295]"
			}
//line /usr/local/go/src/math/big/natconv.go:424
			// _ = "end of CoverTab[6293]"
		}
//line /usr/local/go/src/math/big/natconv.go:425
		// _ = "end of CoverTab[6292]"
	}
//line /usr/local/go/src/math/big/natconv.go:426
	// _ = "end of CoverTab[6274]"
//line /usr/local/go/src/math/big/natconv.go:426
	_go_fuzz_dep_.CoverTab[6275]++

//line /usr/local/go/src/math/big/natconv.go:429
	for i > 0 {
//line /usr/local/go/src/math/big/natconv.go:429
		_go_fuzz_dep_.CoverTab[6296]++
								i--
								s[i] = '0'
//line /usr/local/go/src/math/big/natconv.go:431
		// _ = "end of CoverTab[6296]"
	}
//line /usr/local/go/src/math/big/natconv.go:432
	// _ = "end of CoverTab[6275]"
}

// Split blocks greater than leafSize Words (or set to 0 to disable recursive conversion)
//line /usr/local/go/src/math/big/natconv.go:435
// Benchmark and configure leafSize using: go test -bench="Leaf"
//line /usr/local/go/src/math/big/natconv.go:435
//
//line /usr/local/go/src/math/big/natconv.go:435
//	8 and 16 effective on 3.0 GHz Xeon "Clovertown" CPU (128 byte cache lines)
//line /usr/local/go/src/math/big/natconv.go:435
//	8 and 16 effective on 2.66 GHz Core 2 Duo "Penryn" CPU
//line /usr/local/go/src/math/big/natconv.go:440
var leafSize int = 8	// number of Word-size binary values treat as a monolithic block

type divisor struct {
	bbb	nat	// divisor
	nbits	int	// bit length of divisor (discounting leading zeros) ~= log2(bbb)
	ndigits	int	// digit length of divisor in terms of output base digits
}

var cacheBase10 struct {
	sync.Mutex
	table	[64]divisor	// cached divisors for base 10
}

// expWW computes x**y
func (z nat) expWW(x, y Word) nat {
//line /usr/local/go/src/math/big/natconv.go:454
	_go_fuzz_dep_.CoverTab[6297]++
							return z.expNN(nat(nil).setWord(x), nat(nil).setWord(y), nil, false)
//line /usr/local/go/src/math/big/natconv.go:455
	// _ = "end of CoverTab[6297]"
}

// construct table of powers of bb*leafSize to use in subdivisions.
func divisors(m int, b Word, ndigits int, bb Word) []divisor {
//line /usr/local/go/src/math/big/natconv.go:459
	_go_fuzz_dep_.CoverTab[6298]++

							if leafSize == 0 || func() bool {
//line /usr/local/go/src/math/big/natconv.go:461
		_go_fuzz_dep_.CoverTab[6304]++
//line /usr/local/go/src/math/big/natconv.go:461
		return m <= leafSize
//line /usr/local/go/src/math/big/natconv.go:461
		// _ = "end of CoverTab[6304]"
//line /usr/local/go/src/math/big/natconv.go:461
	}() {
//line /usr/local/go/src/math/big/natconv.go:461
		_go_fuzz_dep_.CoverTab[6305]++
								return nil
//line /usr/local/go/src/math/big/natconv.go:462
		// _ = "end of CoverTab[6305]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:463
		_go_fuzz_dep_.CoverTab[6306]++
//line /usr/local/go/src/math/big/natconv.go:463
		// _ = "end of CoverTab[6306]"
//line /usr/local/go/src/math/big/natconv.go:463
	}
//line /usr/local/go/src/math/big/natconv.go:463
	// _ = "end of CoverTab[6298]"
//line /usr/local/go/src/math/big/natconv.go:463
	_go_fuzz_dep_.CoverTab[6299]++

//line /usr/local/go/src/math/big/natconv.go:466
	k := 1
	for words := leafSize; words < m>>1 && func() bool {
//line /usr/local/go/src/math/big/natconv.go:467
		_go_fuzz_dep_.CoverTab[6307]++
//line /usr/local/go/src/math/big/natconv.go:467
		return k < len(cacheBase10.table)
//line /usr/local/go/src/math/big/natconv.go:467
		// _ = "end of CoverTab[6307]"
//line /usr/local/go/src/math/big/natconv.go:467
	}(); words <<= 1 {
//line /usr/local/go/src/math/big/natconv.go:467
		_go_fuzz_dep_.CoverTab[6308]++
								k++
//line /usr/local/go/src/math/big/natconv.go:468
		// _ = "end of CoverTab[6308]"
	}
//line /usr/local/go/src/math/big/natconv.go:469
	// _ = "end of CoverTab[6299]"
//line /usr/local/go/src/math/big/natconv.go:469
	_go_fuzz_dep_.CoverTab[6300]++

	// reuse and extend existing table of divisors or create new table as appropriate
	var table []divisor	// for b == 10, table overlaps with cacheBase10.table
	if b == 10 {
//line /usr/local/go/src/math/big/natconv.go:473
		_go_fuzz_dep_.CoverTab[6309]++
								cacheBase10.Lock()
								table = cacheBase10.table[0:k]
//line /usr/local/go/src/math/big/natconv.go:475
		// _ = "end of CoverTab[6309]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:476
		_go_fuzz_dep_.CoverTab[6310]++
								table = make([]divisor, k)
//line /usr/local/go/src/math/big/natconv.go:477
		// _ = "end of CoverTab[6310]"
	}
//line /usr/local/go/src/math/big/natconv.go:478
	// _ = "end of CoverTab[6300]"
//line /usr/local/go/src/math/big/natconv.go:478
	_go_fuzz_dep_.CoverTab[6301]++

//line /usr/local/go/src/math/big/natconv.go:481
	if table[k-1].ndigits == 0 {
//line /usr/local/go/src/math/big/natconv.go:481
		_go_fuzz_dep_.CoverTab[6311]++
		// add new entries as needed
		var larger nat
		for i := 0; i < k; i++ {
//line /usr/local/go/src/math/big/natconv.go:484
			_go_fuzz_dep_.CoverTab[6312]++
									if table[i].ndigits == 0 {
//line /usr/local/go/src/math/big/natconv.go:485
				_go_fuzz_dep_.CoverTab[6313]++
										if i == 0 {
//line /usr/local/go/src/math/big/natconv.go:486
					_go_fuzz_dep_.CoverTab[6316]++
											table[0].bbb = nat(nil).expWW(bb, Word(leafSize))
											table[0].ndigits = ndigits * leafSize
//line /usr/local/go/src/math/big/natconv.go:488
					// _ = "end of CoverTab[6316]"
				} else {
//line /usr/local/go/src/math/big/natconv.go:489
					_go_fuzz_dep_.CoverTab[6317]++
											table[i].bbb = nat(nil).sqr(table[i-1].bbb)
											table[i].ndigits = 2 * table[i-1].ndigits
//line /usr/local/go/src/math/big/natconv.go:491
					// _ = "end of CoverTab[6317]"
				}
//line /usr/local/go/src/math/big/natconv.go:492
				// _ = "end of CoverTab[6313]"
//line /usr/local/go/src/math/big/natconv.go:492
				_go_fuzz_dep_.CoverTab[6314]++

//line /usr/local/go/src/math/big/natconv.go:495
				larger = nat(nil).set(table[i].bbb)
				for mulAddVWW(larger, larger, b, 0) == 0 {
//line /usr/local/go/src/math/big/natconv.go:496
					_go_fuzz_dep_.CoverTab[6318]++
											table[i].bbb = table[i].bbb.set(larger)
											table[i].ndigits++
//line /usr/local/go/src/math/big/natconv.go:498
					// _ = "end of CoverTab[6318]"
				}
//line /usr/local/go/src/math/big/natconv.go:499
				// _ = "end of CoverTab[6314]"
//line /usr/local/go/src/math/big/natconv.go:499
				_go_fuzz_dep_.CoverTab[6315]++

										table[i].nbits = table[i].bbb.bitLen()
//line /usr/local/go/src/math/big/natconv.go:501
				// _ = "end of CoverTab[6315]"
			} else {
//line /usr/local/go/src/math/big/natconv.go:502
				_go_fuzz_dep_.CoverTab[6319]++
//line /usr/local/go/src/math/big/natconv.go:502
				// _ = "end of CoverTab[6319]"
//line /usr/local/go/src/math/big/natconv.go:502
			}
//line /usr/local/go/src/math/big/natconv.go:502
			// _ = "end of CoverTab[6312]"
		}
//line /usr/local/go/src/math/big/natconv.go:503
		// _ = "end of CoverTab[6311]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:504
		_go_fuzz_dep_.CoverTab[6320]++
//line /usr/local/go/src/math/big/natconv.go:504
		// _ = "end of CoverTab[6320]"
//line /usr/local/go/src/math/big/natconv.go:504
	}
//line /usr/local/go/src/math/big/natconv.go:504
	// _ = "end of CoverTab[6301]"
//line /usr/local/go/src/math/big/natconv.go:504
	_go_fuzz_dep_.CoverTab[6302]++

							if b == 10 {
//line /usr/local/go/src/math/big/natconv.go:506
		_go_fuzz_dep_.CoverTab[6321]++
								cacheBase10.Unlock()
//line /usr/local/go/src/math/big/natconv.go:507
		// _ = "end of CoverTab[6321]"
	} else {
//line /usr/local/go/src/math/big/natconv.go:508
		_go_fuzz_dep_.CoverTab[6322]++
//line /usr/local/go/src/math/big/natconv.go:508
		// _ = "end of CoverTab[6322]"
//line /usr/local/go/src/math/big/natconv.go:508
	}
//line /usr/local/go/src/math/big/natconv.go:508
	// _ = "end of CoverTab[6302]"
//line /usr/local/go/src/math/big/natconv.go:508
	_go_fuzz_dep_.CoverTab[6303]++

							return table
//line /usr/local/go/src/math/big/natconv.go:510
	// _ = "end of CoverTab[6303]"
}

//line /usr/local/go/src/math/big/natconv.go:511
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/natconv.go:511
var _ = _go_fuzz_dep_.CoverTab
