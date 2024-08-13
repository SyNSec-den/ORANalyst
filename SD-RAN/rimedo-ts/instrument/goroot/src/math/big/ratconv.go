// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements rat-to-string conversion functions.

//line /usr/local/go/src/math/big/ratconv.go:7
package big

//line /usr/local/go/src/math/big/ratconv.go:7
import (
//line /usr/local/go/src/math/big/ratconv.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/ratconv.go:7
)
//line /usr/local/go/src/math/big/ratconv.go:7
import (
//line /usr/local/go/src/math/big/ratconv.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/ratconv.go:7
)

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ratTok(ch rune) bool {
//line /usr/local/go/src/math/big/ratconv.go:17
	_go_fuzz_dep_.CoverTab[6712]++
							return strings.ContainsRune("+-/0123456789.eE", ch)
//line /usr/local/go/src/math/big/ratconv.go:18
	// _ = "end of CoverTab[6712]"
}

var ratZero Rat
var _ fmt.Scanner = &ratZero	// *Rat must implement fmt.Scanner

// Scan is a support routine for fmt.Scanner. It accepts the formats
//line /usr/local/go/src/math/big/ratconv.go:24
// 'e', 'E', 'f', 'F', 'g', 'G', and 'v'. All formats are equivalent.
//line /usr/local/go/src/math/big/ratconv.go:26
func (z *Rat) Scan(s fmt.ScanState, ch rune) error {
//line /usr/local/go/src/math/big/ratconv.go:26
	_go_fuzz_dep_.CoverTab[6713]++
							tok, err := s.Token(true, ratTok)
							if err != nil {
//line /usr/local/go/src/math/big/ratconv.go:28
		_go_fuzz_dep_.CoverTab[6717]++
								return err
//line /usr/local/go/src/math/big/ratconv.go:29
		// _ = "end of CoverTab[6717]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:30
		_go_fuzz_dep_.CoverTab[6718]++
//line /usr/local/go/src/math/big/ratconv.go:30
		// _ = "end of CoverTab[6718]"
//line /usr/local/go/src/math/big/ratconv.go:30
	}
//line /usr/local/go/src/math/big/ratconv.go:30
	// _ = "end of CoverTab[6713]"
//line /usr/local/go/src/math/big/ratconv.go:30
	_go_fuzz_dep_.CoverTab[6714]++
							if !strings.ContainsRune("efgEFGv", ch) {
//line /usr/local/go/src/math/big/ratconv.go:31
		_go_fuzz_dep_.CoverTab[6719]++
								return errors.New("Rat.Scan: invalid verb")
//line /usr/local/go/src/math/big/ratconv.go:32
		// _ = "end of CoverTab[6719]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:33
		_go_fuzz_dep_.CoverTab[6720]++
//line /usr/local/go/src/math/big/ratconv.go:33
		// _ = "end of CoverTab[6720]"
//line /usr/local/go/src/math/big/ratconv.go:33
	}
//line /usr/local/go/src/math/big/ratconv.go:33
	// _ = "end of CoverTab[6714]"
//line /usr/local/go/src/math/big/ratconv.go:33
	_go_fuzz_dep_.CoverTab[6715]++
							if _, ok := z.SetString(string(tok)); !ok {
//line /usr/local/go/src/math/big/ratconv.go:34
		_go_fuzz_dep_.CoverTab[6721]++
								return errors.New("Rat.Scan: invalid syntax")
//line /usr/local/go/src/math/big/ratconv.go:35
		// _ = "end of CoverTab[6721]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:36
		_go_fuzz_dep_.CoverTab[6722]++
//line /usr/local/go/src/math/big/ratconv.go:36
		// _ = "end of CoverTab[6722]"
//line /usr/local/go/src/math/big/ratconv.go:36
	}
//line /usr/local/go/src/math/big/ratconv.go:36
	// _ = "end of CoverTab[6715]"
//line /usr/local/go/src/math/big/ratconv.go:36
	_go_fuzz_dep_.CoverTab[6716]++
							return nil
//line /usr/local/go/src/math/big/ratconv.go:37
	// _ = "end of CoverTab[6716]"
}

// SetString sets z to the value of s and returns z and a boolean indicating
//line /usr/local/go/src/math/big/ratconv.go:40
// success. s can be given as a (possibly signed) fraction "a/b", or as a
//line /usr/local/go/src/math/big/ratconv.go:40
// floating-point number optionally followed by an exponent.
//line /usr/local/go/src/math/big/ratconv.go:40
// If a fraction is provided, both the dividend and the divisor may be a
//line /usr/local/go/src/math/big/ratconv.go:40
// decimal integer or independently use a prefix of “0b”, “0” or “0o”,
//line /usr/local/go/src/math/big/ratconv.go:40
// or “0x” (or their upper-case variants) to denote a binary, octal, or
//line /usr/local/go/src/math/big/ratconv.go:40
// hexadecimal integer, respectively. The divisor may not be signed.
//line /usr/local/go/src/math/big/ratconv.go:40
// If a floating-point number is provided, it may be in decimal form or
//line /usr/local/go/src/math/big/ratconv.go:40
// use any of the same prefixes as above but for “0” to denote a non-decimal
//line /usr/local/go/src/math/big/ratconv.go:40
// mantissa. A leading “0” is considered a decimal leading 0; it does not
//line /usr/local/go/src/math/big/ratconv.go:40
// indicate octal representation in this case.
//line /usr/local/go/src/math/big/ratconv.go:40
// An optional base-10 “e” or base-2 “p” (or their upper-case variants)
//line /usr/local/go/src/math/big/ratconv.go:40
// exponent may be provided as well, except for hexadecimal floats which
//line /usr/local/go/src/math/big/ratconv.go:40
// only accept an (optional) “p” exponent (because an “e” or “E” cannot
//line /usr/local/go/src/math/big/ratconv.go:40
// be distinguished from a mantissa digit). If the exponent's absolute value
//line /usr/local/go/src/math/big/ratconv.go:40
// is too large, the operation may fail.
//line /usr/local/go/src/math/big/ratconv.go:40
// The entire string, not just a prefix, must be valid for success. If the
//line /usr/local/go/src/math/big/ratconv.go:40
// operation failed, the value of z is undefined but the returned value is nil.
//line /usr/local/go/src/math/big/ratconv.go:58
func (z *Rat) SetString(s string) (*Rat, bool) {
//line /usr/local/go/src/math/big/ratconv.go:58
	_go_fuzz_dep_.CoverTab[6723]++
							if len(s) == 0 {
//line /usr/local/go/src/math/big/ratconv.go:59
		_go_fuzz_dep_.CoverTab[6736]++
								return nil, false
//line /usr/local/go/src/math/big/ratconv.go:60
		// _ = "end of CoverTab[6736]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:61
		_go_fuzz_dep_.CoverTab[6737]++
//line /usr/local/go/src/math/big/ratconv.go:61
		// _ = "end of CoverTab[6737]"
//line /usr/local/go/src/math/big/ratconv.go:61
	}
//line /usr/local/go/src/math/big/ratconv.go:61
	// _ = "end of CoverTab[6723]"
//line /usr/local/go/src/math/big/ratconv.go:61
	_go_fuzz_dep_.CoverTab[6724]++

//line /usr/local/go/src/math/big/ratconv.go:65
	if sep := strings.Index(s, "/"); sep >= 0 {
//line /usr/local/go/src/math/big/ratconv.go:65
		_go_fuzz_dep_.CoverTab[6738]++
								if _, ok := z.a.SetString(s[:sep], 0); !ok {
//line /usr/local/go/src/math/big/ratconv.go:66
			_go_fuzz_dep_.CoverTab[6743]++
									return nil, false
//line /usr/local/go/src/math/big/ratconv.go:67
			// _ = "end of CoverTab[6743]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:68
			_go_fuzz_dep_.CoverTab[6744]++
//line /usr/local/go/src/math/big/ratconv.go:68
			// _ = "end of CoverTab[6744]"
//line /usr/local/go/src/math/big/ratconv.go:68
		}
//line /usr/local/go/src/math/big/ratconv.go:68
		// _ = "end of CoverTab[6738]"
//line /usr/local/go/src/math/big/ratconv.go:68
		_go_fuzz_dep_.CoverTab[6739]++
								r := strings.NewReader(s[sep+1:])
								var err error
								if z.b.abs, _, _, err = z.b.abs.scan(r, 0, false); err != nil {
//line /usr/local/go/src/math/big/ratconv.go:71
			_go_fuzz_dep_.CoverTab[6745]++
									return nil, false
//line /usr/local/go/src/math/big/ratconv.go:72
			// _ = "end of CoverTab[6745]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:73
			_go_fuzz_dep_.CoverTab[6746]++
//line /usr/local/go/src/math/big/ratconv.go:73
			// _ = "end of CoverTab[6746]"
//line /usr/local/go/src/math/big/ratconv.go:73
		}
//line /usr/local/go/src/math/big/ratconv.go:73
		// _ = "end of CoverTab[6739]"
//line /usr/local/go/src/math/big/ratconv.go:73
		_go_fuzz_dep_.CoverTab[6740]++

								if _, err = r.ReadByte(); err != io.EOF {
//line /usr/local/go/src/math/big/ratconv.go:75
			_go_fuzz_dep_.CoverTab[6747]++
									return nil, false
//line /usr/local/go/src/math/big/ratconv.go:76
			// _ = "end of CoverTab[6747]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:77
			_go_fuzz_dep_.CoverTab[6748]++
//line /usr/local/go/src/math/big/ratconv.go:77
			// _ = "end of CoverTab[6748]"
//line /usr/local/go/src/math/big/ratconv.go:77
		}
//line /usr/local/go/src/math/big/ratconv.go:77
		// _ = "end of CoverTab[6740]"
//line /usr/local/go/src/math/big/ratconv.go:77
		_go_fuzz_dep_.CoverTab[6741]++
								if len(z.b.abs) == 0 {
//line /usr/local/go/src/math/big/ratconv.go:78
			_go_fuzz_dep_.CoverTab[6749]++
									return nil, false
//line /usr/local/go/src/math/big/ratconv.go:79
			// _ = "end of CoverTab[6749]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:80
			_go_fuzz_dep_.CoverTab[6750]++
//line /usr/local/go/src/math/big/ratconv.go:80
			// _ = "end of CoverTab[6750]"
//line /usr/local/go/src/math/big/ratconv.go:80
		}
//line /usr/local/go/src/math/big/ratconv.go:80
		// _ = "end of CoverTab[6741]"
//line /usr/local/go/src/math/big/ratconv.go:80
		_go_fuzz_dep_.CoverTab[6742]++
								return z.norm(), true
//line /usr/local/go/src/math/big/ratconv.go:81
		// _ = "end of CoverTab[6742]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:82
		_go_fuzz_dep_.CoverTab[6751]++
//line /usr/local/go/src/math/big/ratconv.go:82
		// _ = "end of CoverTab[6751]"
//line /usr/local/go/src/math/big/ratconv.go:82
	}
//line /usr/local/go/src/math/big/ratconv.go:82
	// _ = "end of CoverTab[6724]"
//line /usr/local/go/src/math/big/ratconv.go:82
	_go_fuzz_dep_.CoverTab[6725]++

//line /usr/local/go/src/math/big/ratconv.go:85
	r := strings.NewReader(s)

//line /usr/local/go/src/math/big/ratconv.go:88
	neg, err := scanSign(r)
	if err != nil {
//line /usr/local/go/src/math/big/ratconv.go:89
		_go_fuzz_dep_.CoverTab[6752]++
								return nil, false
//line /usr/local/go/src/math/big/ratconv.go:90
		// _ = "end of CoverTab[6752]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:91
		_go_fuzz_dep_.CoverTab[6753]++
//line /usr/local/go/src/math/big/ratconv.go:91
		// _ = "end of CoverTab[6753]"
//line /usr/local/go/src/math/big/ratconv.go:91
	}
//line /usr/local/go/src/math/big/ratconv.go:91
	// _ = "end of CoverTab[6725]"
//line /usr/local/go/src/math/big/ratconv.go:91
	_go_fuzz_dep_.CoverTab[6726]++

	// mantissa
	var base int
	var fcount int	// fractional digit count; valid if <= 0
	z.a.abs, base, fcount, err = z.a.abs.scan(r, 0, true)
	if err != nil {
//line /usr/local/go/src/math/big/ratconv.go:97
		_go_fuzz_dep_.CoverTab[6754]++
								return nil, false
//line /usr/local/go/src/math/big/ratconv.go:98
		// _ = "end of CoverTab[6754]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:99
		_go_fuzz_dep_.CoverTab[6755]++
//line /usr/local/go/src/math/big/ratconv.go:99
		// _ = "end of CoverTab[6755]"
//line /usr/local/go/src/math/big/ratconv.go:99
	}
//line /usr/local/go/src/math/big/ratconv.go:99
	// _ = "end of CoverTab[6726]"
//line /usr/local/go/src/math/big/ratconv.go:99
	_go_fuzz_dep_.CoverTab[6727]++

	// exponent
	var exp int64
	var ebase int
	exp, ebase, err = scanExponent(r, true, true)
	if err != nil {
//line /usr/local/go/src/math/big/ratconv.go:105
		_go_fuzz_dep_.CoverTab[6756]++
								return nil, false
//line /usr/local/go/src/math/big/ratconv.go:106
		// _ = "end of CoverTab[6756]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:107
		_go_fuzz_dep_.CoverTab[6757]++
//line /usr/local/go/src/math/big/ratconv.go:107
		// _ = "end of CoverTab[6757]"
//line /usr/local/go/src/math/big/ratconv.go:107
	}
//line /usr/local/go/src/math/big/ratconv.go:107
	// _ = "end of CoverTab[6727]"
//line /usr/local/go/src/math/big/ratconv.go:107
	_go_fuzz_dep_.CoverTab[6728]++

//line /usr/local/go/src/math/big/ratconv.go:110
	if _, err = r.ReadByte(); err != io.EOF {
//line /usr/local/go/src/math/big/ratconv.go:110
		_go_fuzz_dep_.CoverTab[6758]++
								return nil, false
//line /usr/local/go/src/math/big/ratconv.go:111
		// _ = "end of CoverTab[6758]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:112
		_go_fuzz_dep_.CoverTab[6759]++
//line /usr/local/go/src/math/big/ratconv.go:112
		// _ = "end of CoverTab[6759]"
//line /usr/local/go/src/math/big/ratconv.go:112
	}
//line /usr/local/go/src/math/big/ratconv.go:112
	// _ = "end of CoverTab[6728]"
//line /usr/local/go/src/math/big/ratconv.go:112
	_go_fuzz_dep_.CoverTab[6729]++

//line /usr/local/go/src/math/big/ratconv.go:115
	if len(z.a.abs) == 0 {
//line /usr/local/go/src/math/big/ratconv.go:115
		_go_fuzz_dep_.CoverTab[6760]++
								return z.norm(), true
//line /usr/local/go/src/math/big/ratconv.go:116
		// _ = "end of CoverTab[6760]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:117
		_go_fuzz_dep_.CoverTab[6761]++
//line /usr/local/go/src/math/big/ratconv.go:117
		// _ = "end of CoverTab[6761]"
//line /usr/local/go/src/math/big/ratconv.go:117
	}
//line /usr/local/go/src/math/big/ratconv.go:117
	// _ = "end of CoverTab[6729]"
//line /usr/local/go/src/math/big/ratconv.go:117
	_go_fuzz_dep_.CoverTab[6730]++

//line /usr/local/go/src/math/big/ratconv.go:131
	// determine binary or decimal exponent contribution of radix point
							var exp2, exp5 int64
							if fcount < 0 {
//line /usr/local/go/src/math/big/ratconv.go:133
		_go_fuzz_dep_.CoverTab[6762]++

//line /usr/local/go/src/math/big/ratconv.go:137
		d := int64(fcount)
		switch base {
		case 10:
//line /usr/local/go/src/math/big/ratconv.go:139
			_go_fuzz_dep_.CoverTab[6763]++
									exp5 = d
									fallthrough
//line /usr/local/go/src/math/big/ratconv.go:141
			// _ = "end of CoverTab[6763]"
		case 2:
//line /usr/local/go/src/math/big/ratconv.go:142
			_go_fuzz_dep_.CoverTab[6764]++
									exp2 = d
//line /usr/local/go/src/math/big/ratconv.go:143
			// _ = "end of CoverTab[6764]"
		case 8:
//line /usr/local/go/src/math/big/ratconv.go:144
			_go_fuzz_dep_.CoverTab[6765]++
									exp2 = d * 3
//line /usr/local/go/src/math/big/ratconv.go:145
			// _ = "end of CoverTab[6765]"
		case 16:
//line /usr/local/go/src/math/big/ratconv.go:146
			_go_fuzz_dep_.CoverTab[6766]++
									exp2 = d * 4
//line /usr/local/go/src/math/big/ratconv.go:147
			// _ = "end of CoverTab[6766]"
		default:
//line /usr/local/go/src/math/big/ratconv.go:148
			_go_fuzz_dep_.CoverTab[6767]++
									panic("unexpected mantissa base")
//line /usr/local/go/src/math/big/ratconv.go:149
			// _ = "end of CoverTab[6767]"
		}
//line /usr/local/go/src/math/big/ratconv.go:150
		// _ = "end of CoverTab[6762]"

	} else {
//line /usr/local/go/src/math/big/ratconv.go:152
		_go_fuzz_dep_.CoverTab[6768]++
//line /usr/local/go/src/math/big/ratconv.go:152
		// _ = "end of CoverTab[6768]"
//line /usr/local/go/src/math/big/ratconv.go:152
	}
//line /usr/local/go/src/math/big/ratconv.go:152
	// _ = "end of CoverTab[6730]"
//line /usr/local/go/src/math/big/ratconv.go:152
	_go_fuzz_dep_.CoverTab[6731]++

//line /usr/local/go/src/math/big/ratconv.go:155
	switch ebase {
	case 10:
//line /usr/local/go/src/math/big/ratconv.go:156
		_go_fuzz_dep_.CoverTab[6769]++
								exp5 += exp
								fallthrough
//line /usr/local/go/src/math/big/ratconv.go:158
		// _ = "end of CoverTab[6769]"
	case 2:
//line /usr/local/go/src/math/big/ratconv.go:159
		_go_fuzz_dep_.CoverTab[6770]++
								exp2 += exp
//line /usr/local/go/src/math/big/ratconv.go:160
		// _ = "end of CoverTab[6770]"
	default:
//line /usr/local/go/src/math/big/ratconv.go:161
		_go_fuzz_dep_.CoverTab[6771]++
								panic("unexpected exponent base")
//line /usr/local/go/src/math/big/ratconv.go:162
		// _ = "end of CoverTab[6771]"
	}
//line /usr/local/go/src/math/big/ratconv.go:163
	// _ = "end of CoverTab[6731]"
//line /usr/local/go/src/math/big/ratconv.go:163
	_go_fuzz_dep_.CoverTab[6732]++

//line /usr/local/go/src/math/big/ratconv.go:168
	if exp5 != 0 {
//line /usr/local/go/src/math/big/ratconv.go:168
		_go_fuzz_dep_.CoverTab[6772]++
								n := exp5
								if n < 0 {
//line /usr/local/go/src/math/big/ratconv.go:170
			_go_fuzz_dep_.CoverTab[6775]++
									n = -n
									if n < 0 {
//line /usr/local/go/src/math/big/ratconv.go:172
				_go_fuzz_dep_.CoverTab[6776]++

//line /usr/local/go/src/math/big/ratconv.go:175
				return nil, false
//line /usr/local/go/src/math/big/ratconv.go:175
				// _ = "end of CoverTab[6776]"
			} else {
//line /usr/local/go/src/math/big/ratconv.go:176
				_go_fuzz_dep_.CoverTab[6777]++
//line /usr/local/go/src/math/big/ratconv.go:176
				// _ = "end of CoverTab[6777]"
//line /usr/local/go/src/math/big/ratconv.go:176
			}
//line /usr/local/go/src/math/big/ratconv.go:176
			// _ = "end of CoverTab[6775]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:177
			_go_fuzz_dep_.CoverTab[6778]++
//line /usr/local/go/src/math/big/ratconv.go:177
			// _ = "end of CoverTab[6778]"
//line /usr/local/go/src/math/big/ratconv.go:177
		}
//line /usr/local/go/src/math/big/ratconv.go:177
		// _ = "end of CoverTab[6772]"
//line /usr/local/go/src/math/big/ratconv.go:177
		_go_fuzz_dep_.CoverTab[6773]++
								if n > 1e6 {
//line /usr/local/go/src/math/big/ratconv.go:178
			_go_fuzz_dep_.CoverTab[6779]++
									return nil, false
//line /usr/local/go/src/math/big/ratconv.go:179
			// _ = "end of CoverTab[6779]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:180
			_go_fuzz_dep_.CoverTab[6780]++
//line /usr/local/go/src/math/big/ratconv.go:180
			// _ = "end of CoverTab[6780]"
//line /usr/local/go/src/math/big/ratconv.go:180
		}
//line /usr/local/go/src/math/big/ratconv.go:180
		// _ = "end of CoverTab[6773]"
//line /usr/local/go/src/math/big/ratconv.go:180
		_go_fuzz_dep_.CoverTab[6774]++
								pow5 := z.b.abs.expNN(natFive, nat(nil).setWord(Word(n)), nil, false)
								if exp5 > 0 {
//line /usr/local/go/src/math/big/ratconv.go:182
			_go_fuzz_dep_.CoverTab[6781]++
									z.a.abs = z.a.abs.mul(z.a.abs, pow5)
									z.b.abs = z.b.abs.setWord(1)
//line /usr/local/go/src/math/big/ratconv.go:184
			// _ = "end of CoverTab[6781]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:185
			_go_fuzz_dep_.CoverTab[6782]++
									z.b.abs = pow5
//line /usr/local/go/src/math/big/ratconv.go:186
			// _ = "end of CoverTab[6782]"
		}
//line /usr/local/go/src/math/big/ratconv.go:187
		// _ = "end of CoverTab[6774]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:188
		_go_fuzz_dep_.CoverTab[6783]++
								z.b.abs = z.b.abs.setWord(1)
//line /usr/local/go/src/math/big/ratconv.go:189
		// _ = "end of CoverTab[6783]"
	}
//line /usr/local/go/src/math/big/ratconv.go:190
	// _ = "end of CoverTab[6732]"
//line /usr/local/go/src/math/big/ratconv.go:190
	_go_fuzz_dep_.CoverTab[6733]++

//line /usr/local/go/src/math/big/ratconv.go:193
	if exp2 < -1e7 || func() bool {
//line /usr/local/go/src/math/big/ratconv.go:193
		_go_fuzz_dep_.CoverTab[6784]++
//line /usr/local/go/src/math/big/ratconv.go:193
		return exp2 > 1e7
//line /usr/local/go/src/math/big/ratconv.go:193
		// _ = "end of CoverTab[6784]"
//line /usr/local/go/src/math/big/ratconv.go:193
	}() {
//line /usr/local/go/src/math/big/ratconv.go:193
		_go_fuzz_dep_.CoverTab[6785]++
								return nil, false
//line /usr/local/go/src/math/big/ratconv.go:194
		// _ = "end of CoverTab[6785]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:195
		_go_fuzz_dep_.CoverTab[6786]++
//line /usr/local/go/src/math/big/ratconv.go:195
		// _ = "end of CoverTab[6786]"
//line /usr/local/go/src/math/big/ratconv.go:195
	}
//line /usr/local/go/src/math/big/ratconv.go:195
	// _ = "end of CoverTab[6733]"
//line /usr/local/go/src/math/big/ratconv.go:195
	_go_fuzz_dep_.CoverTab[6734]++
							if exp2 > 0 {
//line /usr/local/go/src/math/big/ratconv.go:196
		_go_fuzz_dep_.CoverTab[6787]++
								z.a.abs = z.a.abs.shl(z.a.abs, uint(exp2))
//line /usr/local/go/src/math/big/ratconv.go:197
		// _ = "end of CoverTab[6787]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:198
		_go_fuzz_dep_.CoverTab[6788]++
//line /usr/local/go/src/math/big/ratconv.go:198
		if exp2 < 0 {
//line /usr/local/go/src/math/big/ratconv.go:198
			_go_fuzz_dep_.CoverTab[6789]++
									z.b.abs = z.b.abs.shl(z.b.abs, uint(-exp2))
//line /usr/local/go/src/math/big/ratconv.go:199
			// _ = "end of CoverTab[6789]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:200
			_go_fuzz_dep_.CoverTab[6790]++
//line /usr/local/go/src/math/big/ratconv.go:200
			// _ = "end of CoverTab[6790]"
//line /usr/local/go/src/math/big/ratconv.go:200
		}
//line /usr/local/go/src/math/big/ratconv.go:200
		// _ = "end of CoverTab[6788]"
//line /usr/local/go/src/math/big/ratconv.go:200
	}
//line /usr/local/go/src/math/big/ratconv.go:200
	// _ = "end of CoverTab[6734]"
//line /usr/local/go/src/math/big/ratconv.go:200
	_go_fuzz_dep_.CoverTab[6735]++

							z.a.neg = neg && func() bool {
//line /usr/local/go/src/math/big/ratconv.go:202
		_go_fuzz_dep_.CoverTab[6791]++
//line /usr/local/go/src/math/big/ratconv.go:202
		return len(z.a.abs) > 0
//line /usr/local/go/src/math/big/ratconv.go:202
		// _ = "end of CoverTab[6791]"
//line /usr/local/go/src/math/big/ratconv.go:202
	}()

							return z.norm(), true
//line /usr/local/go/src/math/big/ratconv.go:204
	// _ = "end of CoverTab[6735]"
}

// scanExponent scans the longest possible prefix of r representing a base 10
//line /usr/local/go/src/math/big/ratconv.go:207
// (“e”, “E”) or a base 2 (“p”, “P”) exponent, if any. It returns the
//line /usr/local/go/src/math/big/ratconv.go:207
// exponent, the exponent base (10 or 2), or a read or syntax error, if any.
//line /usr/local/go/src/math/big/ratconv.go:207
//
//line /usr/local/go/src/math/big/ratconv.go:207
// If sepOk is set, an underscore character “_” may appear between successive
//line /usr/local/go/src/math/big/ratconv.go:207
// exponent digits; such underscores do not change the value of the exponent.
//line /usr/local/go/src/math/big/ratconv.go:207
// Incorrect placement of underscores is reported as an error if there are no
//line /usr/local/go/src/math/big/ratconv.go:207
// other errors. If sepOk is not set, underscores are not recognized and thus
//line /usr/local/go/src/math/big/ratconv.go:207
// terminate scanning like any other character that is not a valid digit.
//line /usr/local/go/src/math/big/ratconv.go:207
//
//line /usr/local/go/src/math/big/ratconv.go:207
//	exponent = ( "e" | "E" | "p" | "P" ) [ sign ] digits .
//line /usr/local/go/src/math/big/ratconv.go:207
//	sign     = "+" | "-" .
//line /usr/local/go/src/math/big/ratconv.go:207
//	digits   = digit { [ '_' ] digit } .
//line /usr/local/go/src/math/big/ratconv.go:207
//	digit    = "0" ... "9" .
//line /usr/local/go/src/math/big/ratconv.go:207
//
//line /usr/local/go/src/math/big/ratconv.go:207
// A base 2 exponent is only permitted if base2ok is set.
//line /usr/local/go/src/math/big/ratconv.go:223
func scanExponent(r io.ByteScanner, base2ok, sepOk bool) (exp int64, base int, err error) {
//line /usr/local/go/src/math/big/ratconv.go:223
	_go_fuzz_dep_.CoverTab[6792]++

							ch, err := r.ReadByte()
							if err != nil {
//line /usr/local/go/src/math/big/ratconv.go:226
		_go_fuzz_dep_.CoverTab[6801]++
								if err == io.EOF {
//line /usr/local/go/src/math/big/ratconv.go:227
			_go_fuzz_dep_.CoverTab[6803]++
									err = nil
//line /usr/local/go/src/math/big/ratconv.go:228
			// _ = "end of CoverTab[6803]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:229
			_go_fuzz_dep_.CoverTab[6804]++
//line /usr/local/go/src/math/big/ratconv.go:229
			// _ = "end of CoverTab[6804]"
//line /usr/local/go/src/math/big/ratconv.go:229
		}
//line /usr/local/go/src/math/big/ratconv.go:229
		// _ = "end of CoverTab[6801]"
//line /usr/local/go/src/math/big/ratconv.go:229
		_go_fuzz_dep_.CoverTab[6802]++
								return 0, 10, err
//line /usr/local/go/src/math/big/ratconv.go:230
		// _ = "end of CoverTab[6802]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:231
		_go_fuzz_dep_.CoverTab[6805]++
//line /usr/local/go/src/math/big/ratconv.go:231
		// _ = "end of CoverTab[6805]"
//line /usr/local/go/src/math/big/ratconv.go:231
	}
//line /usr/local/go/src/math/big/ratconv.go:231
	// _ = "end of CoverTab[6792]"
//line /usr/local/go/src/math/big/ratconv.go:231
	_go_fuzz_dep_.CoverTab[6793]++

//line /usr/local/go/src/math/big/ratconv.go:234
	switch ch {
	case 'e', 'E':
//line /usr/local/go/src/math/big/ratconv.go:235
		_go_fuzz_dep_.CoverTab[6806]++
								base = 10
//line /usr/local/go/src/math/big/ratconv.go:236
		// _ = "end of CoverTab[6806]"
	case 'p', 'P':
//line /usr/local/go/src/math/big/ratconv.go:237
		_go_fuzz_dep_.CoverTab[6807]++
								if base2ok {
//line /usr/local/go/src/math/big/ratconv.go:238
			_go_fuzz_dep_.CoverTab[6810]++
									base = 2
									break
//line /usr/local/go/src/math/big/ratconv.go:240
			// _ = "end of CoverTab[6810]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:241
			_go_fuzz_dep_.CoverTab[6811]++
//line /usr/local/go/src/math/big/ratconv.go:241
			// _ = "end of CoverTab[6811]"
//line /usr/local/go/src/math/big/ratconv.go:241
		}
//line /usr/local/go/src/math/big/ratconv.go:241
		// _ = "end of CoverTab[6807]"
//line /usr/local/go/src/math/big/ratconv.go:241
		_go_fuzz_dep_.CoverTab[6808]++
								fallthrough
//line /usr/local/go/src/math/big/ratconv.go:242
		// _ = "end of CoverTab[6808]"
	default:
//line /usr/local/go/src/math/big/ratconv.go:243
		_go_fuzz_dep_.CoverTab[6809]++
								r.UnreadByte()
								return 0, 10, nil
//line /usr/local/go/src/math/big/ratconv.go:245
		// _ = "end of CoverTab[6809]"
	}
//line /usr/local/go/src/math/big/ratconv.go:246
	// _ = "end of CoverTab[6793]"
//line /usr/local/go/src/math/big/ratconv.go:246
	_go_fuzz_dep_.CoverTab[6794]++

	// sign
	var digits []byte
	ch, err = r.ReadByte()
	if err == nil && func() bool {
//line /usr/local/go/src/math/big/ratconv.go:251
		_go_fuzz_dep_.CoverTab[6812]++
//line /usr/local/go/src/math/big/ratconv.go:251
		return (ch == '+' || func() bool {
//line /usr/local/go/src/math/big/ratconv.go:251
			_go_fuzz_dep_.CoverTab[6813]++
//line /usr/local/go/src/math/big/ratconv.go:251
			return ch == '-'
//line /usr/local/go/src/math/big/ratconv.go:251
			// _ = "end of CoverTab[6813]"
//line /usr/local/go/src/math/big/ratconv.go:251
		}())
//line /usr/local/go/src/math/big/ratconv.go:251
		// _ = "end of CoverTab[6812]"
//line /usr/local/go/src/math/big/ratconv.go:251
	}() {
//line /usr/local/go/src/math/big/ratconv.go:251
		_go_fuzz_dep_.CoverTab[6814]++
								if ch == '-' {
//line /usr/local/go/src/math/big/ratconv.go:252
			_go_fuzz_dep_.CoverTab[6816]++
									digits = append(digits, '-')
//line /usr/local/go/src/math/big/ratconv.go:253
			// _ = "end of CoverTab[6816]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:254
			_go_fuzz_dep_.CoverTab[6817]++
//line /usr/local/go/src/math/big/ratconv.go:254
			// _ = "end of CoverTab[6817]"
//line /usr/local/go/src/math/big/ratconv.go:254
		}
//line /usr/local/go/src/math/big/ratconv.go:254
		// _ = "end of CoverTab[6814]"
//line /usr/local/go/src/math/big/ratconv.go:254
		_go_fuzz_dep_.CoverTab[6815]++
								ch, err = r.ReadByte()
//line /usr/local/go/src/math/big/ratconv.go:255
		// _ = "end of CoverTab[6815]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:256
		_go_fuzz_dep_.CoverTab[6818]++
//line /usr/local/go/src/math/big/ratconv.go:256
		// _ = "end of CoverTab[6818]"
//line /usr/local/go/src/math/big/ratconv.go:256
	}
//line /usr/local/go/src/math/big/ratconv.go:256
	// _ = "end of CoverTab[6794]"
//line /usr/local/go/src/math/big/ratconv.go:256
	_go_fuzz_dep_.CoverTab[6795]++

//line /usr/local/go/src/math/big/ratconv.go:261
	prev := '.'
							invalSep := false

//line /usr/local/go/src/math/big/ratconv.go:265
	hasDigits := false
	for err == nil {
//line /usr/local/go/src/math/big/ratconv.go:266
		_go_fuzz_dep_.CoverTab[6819]++
								if '0' <= ch && func() bool {
//line /usr/local/go/src/math/big/ratconv.go:267
			_go_fuzz_dep_.CoverTab[6821]++
//line /usr/local/go/src/math/big/ratconv.go:267
			return ch <= '9'
//line /usr/local/go/src/math/big/ratconv.go:267
			// _ = "end of CoverTab[6821]"
//line /usr/local/go/src/math/big/ratconv.go:267
		}() {
//line /usr/local/go/src/math/big/ratconv.go:267
			_go_fuzz_dep_.CoverTab[6822]++
									digits = append(digits, ch)
									prev = '0'
									hasDigits = true
//line /usr/local/go/src/math/big/ratconv.go:270
			// _ = "end of CoverTab[6822]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:271
			_go_fuzz_dep_.CoverTab[6823]++
//line /usr/local/go/src/math/big/ratconv.go:271
			if ch == '_' && func() bool {
//line /usr/local/go/src/math/big/ratconv.go:271
				_go_fuzz_dep_.CoverTab[6824]++
//line /usr/local/go/src/math/big/ratconv.go:271
				return sepOk
//line /usr/local/go/src/math/big/ratconv.go:271
				// _ = "end of CoverTab[6824]"
//line /usr/local/go/src/math/big/ratconv.go:271
			}() {
//line /usr/local/go/src/math/big/ratconv.go:271
				_go_fuzz_dep_.CoverTab[6825]++
										if prev != '0' {
//line /usr/local/go/src/math/big/ratconv.go:272
					_go_fuzz_dep_.CoverTab[6827]++
											invalSep = true
//line /usr/local/go/src/math/big/ratconv.go:273
					// _ = "end of CoverTab[6827]"
				} else {
//line /usr/local/go/src/math/big/ratconv.go:274
					_go_fuzz_dep_.CoverTab[6828]++
//line /usr/local/go/src/math/big/ratconv.go:274
					// _ = "end of CoverTab[6828]"
//line /usr/local/go/src/math/big/ratconv.go:274
				}
//line /usr/local/go/src/math/big/ratconv.go:274
				// _ = "end of CoverTab[6825]"
//line /usr/local/go/src/math/big/ratconv.go:274
				_go_fuzz_dep_.CoverTab[6826]++
										prev = '_'
//line /usr/local/go/src/math/big/ratconv.go:275
				// _ = "end of CoverTab[6826]"
			} else {
//line /usr/local/go/src/math/big/ratconv.go:276
				_go_fuzz_dep_.CoverTab[6829]++
										r.UnreadByte()
										break
//line /usr/local/go/src/math/big/ratconv.go:278
				// _ = "end of CoverTab[6829]"
			}
//line /usr/local/go/src/math/big/ratconv.go:279
			// _ = "end of CoverTab[6823]"
//line /usr/local/go/src/math/big/ratconv.go:279
		}
//line /usr/local/go/src/math/big/ratconv.go:279
		// _ = "end of CoverTab[6819]"
//line /usr/local/go/src/math/big/ratconv.go:279
		_go_fuzz_dep_.CoverTab[6820]++
								ch, err = r.ReadByte()
//line /usr/local/go/src/math/big/ratconv.go:280
		// _ = "end of CoverTab[6820]"
	}
//line /usr/local/go/src/math/big/ratconv.go:281
	// _ = "end of CoverTab[6795]"
//line /usr/local/go/src/math/big/ratconv.go:281
	_go_fuzz_dep_.CoverTab[6796]++

							if err == io.EOF {
//line /usr/local/go/src/math/big/ratconv.go:283
		_go_fuzz_dep_.CoverTab[6830]++
								err = nil
//line /usr/local/go/src/math/big/ratconv.go:284
		// _ = "end of CoverTab[6830]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:285
		_go_fuzz_dep_.CoverTab[6831]++
//line /usr/local/go/src/math/big/ratconv.go:285
		// _ = "end of CoverTab[6831]"
//line /usr/local/go/src/math/big/ratconv.go:285
	}
//line /usr/local/go/src/math/big/ratconv.go:285
	// _ = "end of CoverTab[6796]"
//line /usr/local/go/src/math/big/ratconv.go:285
	_go_fuzz_dep_.CoverTab[6797]++
							if err == nil && func() bool {
//line /usr/local/go/src/math/big/ratconv.go:286
		_go_fuzz_dep_.CoverTab[6832]++
//line /usr/local/go/src/math/big/ratconv.go:286
		return !hasDigits
//line /usr/local/go/src/math/big/ratconv.go:286
		// _ = "end of CoverTab[6832]"
//line /usr/local/go/src/math/big/ratconv.go:286
	}() {
//line /usr/local/go/src/math/big/ratconv.go:286
		_go_fuzz_dep_.CoverTab[6833]++
								err = errNoDigits
//line /usr/local/go/src/math/big/ratconv.go:287
		// _ = "end of CoverTab[6833]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:288
		_go_fuzz_dep_.CoverTab[6834]++
//line /usr/local/go/src/math/big/ratconv.go:288
		// _ = "end of CoverTab[6834]"
//line /usr/local/go/src/math/big/ratconv.go:288
	}
//line /usr/local/go/src/math/big/ratconv.go:288
	// _ = "end of CoverTab[6797]"
//line /usr/local/go/src/math/big/ratconv.go:288
	_go_fuzz_dep_.CoverTab[6798]++
							if err == nil {
//line /usr/local/go/src/math/big/ratconv.go:289
		_go_fuzz_dep_.CoverTab[6835]++
								exp, err = strconv.ParseInt(string(digits), 10, 64)
//line /usr/local/go/src/math/big/ratconv.go:290
		// _ = "end of CoverTab[6835]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:291
		_go_fuzz_dep_.CoverTab[6836]++
//line /usr/local/go/src/math/big/ratconv.go:291
		// _ = "end of CoverTab[6836]"
//line /usr/local/go/src/math/big/ratconv.go:291
	}
//line /usr/local/go/src/math/big/ratconv.go:291
	// _ = "end of CoverTab[6798]"
//line /usr/local/go/src/math/big/ratconv.go:291
	_go_fuzz_dep_.CoverTab[6799]++

							if err == nil && func() bool {
//line /usr/local/go/src/math/big/ratconv.go:293
		_go_fuzz_dep_.CoverTab[6837]++
//line /usr/local/go/src/math/big/ratconv.go:293
		return (invalSep || func() bool {
//line /usr/local/go/src/math/big/ratconv.go:293
			_go_fuzz_dep_.CoverTab[6838]++
//line /usr/local/go/src/math/big/ratconv.go:293
			return prev == '_'
//line /usr/local/go/src/math/big/ratconv.go:293
			// _ = "end of CoverTab[6838]"
//line /usr/local/go/src/math/big/ratconv.go:293
		}())
//line /usr/local/go/src/math/big/ratconv.go:293
		// _ = "end of CoverTab[6837]"
//line /usr/local/go/src/math/big/ratconv.go:293
	}() {
//line /usr/local/go/src/math/big/ratconv.go:293
		_go_fuzz_dep_.CoverTab[6839]++
								err = errInvalSep
//line /usr/local/go/src/math/big/ratconv.go:294
		// _ = "end of CoverTab[6839]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:295
		_go_fuzz_dep_.CoverTab[6840]++
//line /usr/local/go/src/math/big/ratconv.go:295
		// _ = "end of CoverTab[6840]"
//line /usr/local/go/src/math/big/ratconv.go:295
	}
//line /usr/local/go/src/math/big/ratconv.go:295
	// _ = "end of CoverTab[6799]"
//line /usr/local/go/src/math/big/ratconv.go:295
	_go_fuzz_dep_.CoverTab[6800]++

							return
//line /usr/local/go/src/math/big/ratconv.go:297
	// _ = "end of CoverTab[6800]"
}

// String returns a string representation of x in the form "a/b" (even if b == 1).
func (x *Rat) String() string {
//line /usr/local/go/src/math/big/ratconv.go:301
	_go_fuzz_dep_.CoverTab[6841]++
							return string(x.marshal())
//line /usr/local/go/src/math/big/ratconv.go:302
	// _ = "end of CoverTab[6841]"
}

// marshal implements String returning a slice of bytes
func (x *Rat) marshal() []byte {
//line /usr/local/go/src/math/big/ratconv.go:306
	_go_fuzz_dep_.CoverTab[6842]++
							var buf []byte
							buf = x.a.Append(buf, 10)
							buf = append(buf, '/')
							if len(x.b.abs) != 0 {
//line /usr/local/go/src/math/big/ratconv.go:310
		_go_fuzz_dep_.CoverTab[6844]++
								buf = x.b.Append(buf, 10)
//line /usr/local/go/src/math/big/ratconv.go:311
		// _ = "end of CoverTab[6844]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:312
		_go_fuzz_dep_.CoverTab[6845]++
								buf = append(buf, '1')
//line /usr/local/go/src/math/big/ratconv.go:313
		// _ = "end of CoverTab[6845]"
	}
//line /usr/local/go/src/math/big/ratconv.go:314
	// _ = "end of CoverTab[6842]"
//line /usr/local/go/src/math/big/ratconv.go:314
	_go_fuzz_dep_.CoverTab[6843]++
							return buf
//line /usr/local/go/src/math/big/ratconv.go:315
	// _ = "end of CoverTab[6843]"
}

// RatString returns a string representation of x in the form "a/b" if b != 1,
//line /usr/local/go/src/math/big/ratconv.go:318
// and in the form "a" if b == 1.
//line /usr/local/go/src/math/big/ratconv.go:320
func (x *Rat) RatString() string {
//line /usr/local/go/src/math/big/ratconv.go:320
	_go_fuzz_dep_.CoverTab[6846]++
							if x.IsInt() {
//line /usr/local/go/src/math/big/ratconv.go:321
		_go_fuzz_dep_.CoverTab[6848]++
								return x.a.String()
//line /usr/local/go/src/math/big/ratconv.go:322
		// _ = "end of CoverTab[6848]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:323
		_go_fuzz_dep_.CoverTab[6849]++
//line /usr/local/go/src/math/big/ratconv.go:323
		// _ = "end of CoverTab[6849]"
//line /usr/local/go/src/math/big/ratconv.go:323
	}
//line /usr/local/go/src/math/big/ratconv.go:323
	// _ = "end of CoverTab[6846]"
//line /usr/local/go/src/math/big/ratconv.go:323
	_go_fuzz_dep_.CoverTab[6847]++
							return x.String()
//line /usr/local/go/src/math/big/ratconv.go:324
	// _ = "end of CoverTab[6847]"
}

// FloatString returns a string representation of x in decimal form with prec
//line /usr/local/go/src/math/big/ratconv.go:327
// digits of precision after the radix point. The last digit is rounded to
//line /usr/local/go/src/math/big/ratconv.go:327
// nearest, with halves rounded away from zero.
//line /usr/local/go/src/math/big/ratconv.go:330
func (x *Rat) FloatString(prec int) string {
//line /usr/local/go/src/math/big/ratconv.go:330
	_go_fuzz_dep_.CoverTab[6850]++
							var buf []byte

							if x.IsInt() {
//line /usr/local/go/src/math/big/ratconv.go:333
		_go_fuzz_dep_.CoverTab[6856]++
								buf = x.a.Append(buf, 10)
								if prec > 0 {
//line /usr/local/go/src/math/big/ratconv.go:335
			_go_fuzz_dep_.CoverTab[6858]++
									buf = append(buf, '.')
									for i := prec; i > 0; i-- {
//line /usr/local/go/src/math/big/ratconv.go:337
				_go_fuzz_dep_.CoverTab[6859]++
										buf = append(buf, '0')
//line /usr/local/go/src/math/big/ratconv.go:338
				// _ = "end of CoverTab[6859]"
			}
//line /usr/local/go/src/math/big/ratconv.go:339
			// _ = "end of CoverTab[6858]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:340
			_go_fuzz_dep_.CoverTab[6860]++
//line /usr/local/go/src/math/big/ratconv.go:340
			// _ = "end of CoverTab[6860]"
//line /usr/local/go/src/math/big/ratconv.go:340
		}
//line /usr/local/go/src/math/big/ratconv.go:340
		// _ = "end of CoverTab[6856]"
//line /usr/local/go/src/math/big/ratconv.go:340
		_go_fuzz_dep_.CoverTab[6857]++
								return string(buf)
//line /usr/local/go/src/math/big/ratconv.go:341
		// _ = "end of CoverTab[6857]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:342
		_go_fuzz_dep_.CoverTab[6861]++
//line /usr/local/go/src/math/big/ratconv.go:342
		// _ = "end of CoverTab[6861]"
//line /usr/local/go/src/math/big/ratconv.go:342
	}
//line /usr/local/go/src/math/big/ratconv.go:342
	// _ = "end of CoverTab[6850]"
//line /usr/local/go/src/math/big/ratconv.go:342
	_go_fuzz_dep_.CoverTab[6851]++

//line /usr/local/go/src/math/big/ratconv.go:345
	q, r := nat(nil).div(nat(nil), x.a.abs, x.b.abs)

	p := natOne
	if prec > 0 {
//line /usr/local/go/src/math/big/ratconv.go:348
		_go_fuzz_dep_.CoverTab[6862]++
								p = nat(nil).expNN(natTen, nat(nil).setUint64(uint64(prec)), nil, false)
//line /usr/local/go/src/math/big/ratconv.go:349
		// _ = "end of CoverTab[6862]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:350
		_go_fuzz_dep_.CoverTab[6863]++
//line /usr/local/go/src/math/big/ratconv.go:350
		// _ = "end of CoverTab[6863]"
//line /usr/local/go/src/math/big/ratconv.go:350
	}
//line /usr/local/go/src/math/big/ratconv.go:350
	// _ = "end of CoverTab[6851]"
//line /usr/local/go/src/math/big/ratconv.go:350
	_go_fuzz_dep_.CoverTab[6852]++

							r = r.mul(r, p)
							r, r2 := r.div(nat(nil), r, x.b.abs)

//line /usr/local/go/src/math/big/ratconv.go:356
	r2 = r2.add(r2, r2)
	if x.b.abs.cmp(r2) <= 0 {
//line /usr/local/go/src/math/big/ratconv.go:357
		_go_fuzz_dep_.CoverTab[6864]++
								r = r.add(r, natOne)
								if r.cmp(p) >= 0 {
//line /usr/local/go/src/math/big/ratconv.go:359
			_go_fuzz_dep_.CoverTab[6865]++
									q = nat(nil).add(q, natOne)
									r = nat(nil).sub(r, p)
//line /usr/local/go/src/math/big/ratconv.go:361
			// _ = "end of CoverTab[6865]"
		} else {
//line /usr/local/go/src/math/big/ratconv.go:362
			_go_fuzz_dep_.CoverTab[6866]++
//line /usr/local/go/src/math/big/ratconv.go:362
			// _ = "end of CoverTab[6866]"
//line /usr/local/go/src/math/big/ratconv.go:362
		}
//line /usr/local/go/src/math/big/ratconv.go:362
		// _ = "end of CoverTab[6864]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:363
		_go_fuzz_dep_.CoverTab[6867]++
//line /usr/local/go/src/math/big/ratconv.go:363
		// _ = "end of CoverTab[6867]"
//line /usr/local/go/src/math/big/ratconv.go:363
	}
//line /usr/local/go/src/math/big/ratconv.go:363
	// _ = "end of CoverTab[6852]"
//line /usr/local/go/src/math/big/ratconv.go:363
	_go_fuzz_dep_.CoverTab[6853]++

							if x.a.neg {
//line /usr/local/go/src/math/big/ratconv.go:365
		_go_fuzz_dep_.CoverTab[6868]++
								buf = append(buf, '-')
//line /usr/local/go/src/math/big/ratconv.go:366
		// _ = "end of CoverTab[6868]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:367
		_go_fuzz_dep_.CoverTab[6869]++
//line /usr/local/go/src/math/big/ratconv.go:367
		// _ = "end of CoverTab[6869]"
//line /usr/local/go/src/math/big/ratconv.go:367
	}
//line /usr/local/go/src/math/big/ratconv.go:367
	// _ = "end of CoverTab[6853]"
//line /usr/local/go/src/math/big/ratconv.go:367
	_go_fuzz_dep_.CoverTab[6854]++
							buf = append(buf, q.utoa(10)...)

							if prec > 0 {
//line /usr/local/go/src/math/big/ratconv.go:370
		_go_fuzz_dep_.CoverTab[6870]++
								buf = append(buf, '.')
								rs := r.utoa(10)
								for i := prec - len(rs); i > 0; i-- {
//line /usr/local/go/src/math/big/ratconv.go:373
			_go_fuzz_dep_.CoverTab[6872]++
									buf = append(buf, '0')
//line /usr/local/go/src/math/big/ratconv.go:374
			// _ = "end of CoverTab[6872]"
		}
//line /usr/local/go/src/math/big/ratconv.go:375
		// _ = "end of CoverTab[6870]"
//line /usr/local/go/src/math/big/ratconv.go:375
		_go_fuzz_dep_.CoverTab[6871]++
								buf = append(buf, rs...)
//line /usr/local/go/src/math/big/ratconv.go:376
		// _ = "end of CoverTab[6871]"
	} else {
//line /usr/local/go/src/math/big/ratconv.go:377
		_go_fuzz_dep_.CoverTab[6873]++
//line /usr/local/go/src/math/big/ratconv.go:377
		// _ = "end of CoverTab[6873]"
//line /usr/local/go/src/math/big/ratconv.go:377
	}
//line /usr/local/go/src/math/big/ratconv.go:377
	// _ = "end of CoverTab[6854]"
//line /usr/local/go/src/math/big/ratconv.go:377
	_go_fuzz_dep_.CoverTab[6855]++

							return string(buf)
//line /usr/local/go/src/math/big/ratconv.go:379
	// _ = "end of CoverTab[6855]"
}

//line /usr/local/go/src/math/big/ratconv.go:380
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/ratconv.go:380
var _ = _go_fuzz_dep_.CoverTab
