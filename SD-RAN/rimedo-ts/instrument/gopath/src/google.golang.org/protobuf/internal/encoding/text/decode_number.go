// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:5
package text

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:5
)

// parseNumberValue parses a number from the input and returns a Token object.
func (d *Decoder) parseNumberValue() (Token, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:8
	_go_fuzz_dep_.CoverTab[49706]++
														in := d.in
														num := parseNumber(in)
														if num.size == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:11
		_go_fuzz_dep_.CoverTab[49710]++
																return Token{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:12
		// _ = "end of CoverTab[49710]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:13
		_go_fuzz_dep_.CoverTab[49711]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:13
		// _ = "end of CoverTab[49711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:13
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:13
	// _ = "end of CoverTab[49706]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:13
	_go_fuzz_dep_.CoverTab[49707]++
															numAttrs := num.kind
															if num.neg {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:15
		_go_fuzz_dep_.CoverTab[49712]++
																numAttrs |= isNegative
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:16
		// _ = "end of CoverTab[49712]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:17
		_go_fuzz_dep_.CoverTab[49713]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:17
		// _ = "end of CoverTab[49713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:17
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:17
	// _ = "end of CoverTab[49707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:17
	_go_fuzz_dep_.CoverTab[49708]++
															strSize := num.size
															last := num.size - 1
															if num.kind == numFloat && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:20
		_go_fuzz_dep_.CoverTab[49714]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:20
		return (d.in[last] == 'f' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:20
			_go_fuzz_dep_.CoverTab[49715]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:20
			return d.in[last] == 'F'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:20
			// _ = "end of CoverTab[49715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:20
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:20
		// _ = "end of CoverTab[49714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:20
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:20
		_go_fuzz_dep_.CoverTab[49716]++
																strSize = last
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:21
		// _ = "end of CoverTab[49716]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:22
		_go_fuzz_dep_.CoverTab[49717]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:22
		// _ = "end of CoverTab[49717]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:22
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:22
	// _ = "end of CoverTab[49708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:22
	_go_fuzz_dep_.CoverTab[49709]++
															tok := Token{
		kind:		Scalar,
		attrs:		numberValue,
		pos:		len(d.orig) - len(d.in),
		raw:		d.in[:num.size],
		str:		string(d.in[:strSize]),
		numAttrs:	numAttrs,
	}
															d.consume(num.size)
															return tok, true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:32
	// _ = "end of CoverTab[49709]"
}

const (
	numDec	uint8	= (1 << iota) / 2
	numHex
	numOct
	numFloat
)

// number is the result of parsing out a valid number from parseNumber. It
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:42
// contains data for doing float or integer conversion via the strconv package
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:42
// in conjunction with the input bytes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:45
type number struct {
	kind	uint8
	neg	bool
	size	int
}

// parseNumber constructs a number object from given input. It allows for the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:51
// following patterns:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:51
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:51
//	integer: ^-?([1-9][0-9]*|0[xX][0-9a-fA-F]+|0[0-7]*)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:51
//	float: ^-?((0|[1-9][0-9]*)?([.][0-9]*)?([eE][+-]?[0-9]+)?[fF]?)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:51
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:51
// It also returns the number of parsed bytes for the given number, 0 if it is
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:51
// not a number.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:59
func parseNumber(input []byte) number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:59
	_go_fuzz_dep_.CoverTab[49718]++
															kind := numDec
															var size int
															var neg bool

															s := input
															if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:65
		_go_fuzz_dep_.CoverTab[49726]++
																return number{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:66
		// _ = "end of CoverTab[49726]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:67
		_go_fuzz_dep_.CoverTab[49727]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:67
		// _ = "end of CoverTab[49727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:67
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:67
	// _ = "end of CoverTab[49718]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:67
	_go_fuzz_dep_.CoverTab[49719]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:70
	if s[0] == '-' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:70
		_go_fuzz_dep_.CoverTab[49728]++
																neg = true
																s = s[1:]
																size++
																if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:74
			_go_fuzz_dep_.CoverTab[49729]++
																	return number{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:75
			// _ = "end of CoverTab[49729]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:76
			_go_fuzz_dep_.CoverTab[49730]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:76
			// _ = "end of CoverTab[49730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:76
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:76
		// _ = "end of CoverTab[49728]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:77
		_go_fuzz_dep_.CoverTab[49731]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:77
		// _ = "end of CoverTab[49731]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:77
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:77
	// _ = "end of CoverTab[49719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:77
	_go_fuzz_dep_.CoverTab[49720]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:83
	switch {
	case s[0] == '0':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:84
		_go_fuzz_dep_.CoverTab[49732]++
																if len(s) > 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:85
			_go_fuzz_dep_.CoverTab[49738]++
																	switch {
			case s[1] == 'x' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:87
				_go_fuzz_dep_.CoverTab[49746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:87
				return s[1] == 'X'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:87
				// _ = "end of CoverTab[49746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:87
			}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:87
				_go_fuzz_dep_.CoverTab[49740]++

																		kind = numHex
																		n := 2
																		s = s[2:]
																		for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:92
					_go_fuzz_dep_.CoverTab[49747]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:92
					return (('0' <= s[0] && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:92
						_go_fuzz_dep_.CoverTab[49748]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:92
						return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:92
						// _ = "end of CoverTab[49748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:92
					}()) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:92
						_go_fuzz_dep_.CoverTab[49749]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:92
						return ('a' <= s[0] && func() bool {
																					_go_fuzz_dep_.CoverTab[49750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:93
							return s[0] <= 'f'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:93
							// _ = "end of CoverTab[49750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:93
						}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:93
						// _ = "end of CoverTab[49749]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:93
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:93
						_go_fuzz_dep_.CoverTab[49751]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:93
						return ('A' <= s[0] && func() bool {
																					_go_fuzz_dep_.CoverTab[49752]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:94
							return s[0] <= 'F'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:94
							// _ = "end of CoverTab[49752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:94
						}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:94
						// _ = "end of CoverTab[49751]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:94
					}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:94
					// _ = "end of CoverTab[49747]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:94
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:94
					_go_fuzz_dep_.CoverTab[49753]++
																			s = s[1:]
																			n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:96
					// _ = "end of CoverTab[49753]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:97
				// _ = "end of CoverTab[49740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:97
				_go_fuzz_dep_.CoverTab[49741]++
																		if n == 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:98
					_go_fuzz_dep_.CoverTab[49754]++
																			return number{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:99
					// _ = "end of CoverTab[49754]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:100
					_go_fuzz_dep_.CoverTab[49755]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:100
					// _ = "end of CoverTab[49755]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:100
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:100
				// _ = "end of CoverTab[49741]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:100
				_go_fuzz_dep_.CoverTab[49742]++
																		size += n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:101
				// _ = "end of CoverTab[49742]"

			case '0' <= s[1] && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:103
				_go_fuzz_dep_.CoverTab[49756]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:103
				return s[1] <= '7'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:103
				// _ = "end of CoverTab[49756]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:103
			}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:103
				_go_fuzz_dep_.CoverTab[49743]++

																		kind = numOct
																		n := 2
																		s = s[2:]
																		for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:108
					_go_fuzz_dep_.CoverTab[49757]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:108
					return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:108
					// _ = "end of CoverTab[49757]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:108
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:108
					_go_fuzz_dep_.CoverTab[49758]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:108
					return s[0] <= '7'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:108
					// _ = "end of CoverTab[49758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:108
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:108
					_go_fuzz_dep_.CoverTab[49759]++
																			s = s[1:]
																			n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:110
					// _ = "end of CoverTab[49759]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:111
				// _ = "end of CoverTab[49743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:111
				_go_fuzz_dep_.CoverTab[49744]++
																		size += n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:112
				// _ = "end of CoverTab[49744]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:112
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:112
				_go_fuzz_dep_.CoverTab[49745]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:112
				// _ = "end of CoverTab[49745]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:113
			// _ = "end of CoverTab[49738]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:113
			_go_fuzz_dep_.CoverTab[49739]++

																	if kind&(numHex|numOct) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:115
				_go_fuzz_dep_.CoverTab[49760]++
																		if len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:116
					_go_fuzz_dep_.CoverTab[49762]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:116
					return !isDelim(s[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:116
					// _ = "end of CoverTab[49762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:116
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:116
					_go_fuzz_dep_.CoverTab[49763]++
																			return number{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:117
					// _ = "end of CoverTab[49763]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:118
					_go_fuzz_dep_.CoverTab[49764]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:118
					// _ = "end of CoverTab[49764]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:118
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:118
				// _ = "end of CoverTab[49760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:118
				_go_fuzz_dep_.CoverTab[49761]++
																		return number{kind: kind, neg: neg, size: size}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:119
				// _ = "end of CoverTab[49761]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:120
				_go_fuzz_dep_.CoverTab[49765]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:120
				// _ = "end of CoverTab[49765]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:120
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:120
			// _ = "end of CoverTab[49739]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:121
			_go_fuzz_dep_.CoverTab[49766]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:121
			// _ = "end of CoverTab[49766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:121
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:121
		// _ = "end of CoverTab[49732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:121
		_go_fuzz_dep_.CoverTab[49733]++
																s = s[1:]
																size++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:123
		// _ = "end of CoverTab[49733]"

	case '1' <= s[0] && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:125
		_go_fuzz_dep_.CoverTab[49767]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:125
		return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:125
		// _ = "end of CoverTab[49767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:125
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:125
		_go_fuzz_dep_.CoverTab[49734]++
																n := 1
																s = s[1:]
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:128
			_go_fuzz_dep_.CoverTab[49768]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:128
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:128
			// _ = "end of CoverTab[49768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:128
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:128
			_go_fuzz_dep_.CoverTab[49769]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:128
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:128
			// _ = "end of CoverTab[49769]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:128
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:128
			_go_fuzz_dep_.CoverTab[49770]++
																	s = s[1:]
																	n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:130
			// _ = "end of CoverTab[49770]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:131
		// _ = "end of CoverTab[49734]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:131
		_go_fuzz_dep_.CoverTab[49735]++
																size += n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:132
		// _ = "end of CoverTab[49735]"

	case s[0] == '.':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:134
		_go_fuzz_dep_.CoverTab[49736]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:137
		kind = numFloat
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:137
		// _ = "end of CoverTab[49736]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:139
		_go_fuzz_dep_.CoverTab[49737]++
																return number{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:140
		// _ = "end of CoverTab[49737]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:141
	// _ = "end of CoverTab[49720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:141
	_go_fuzz_dep_.CoverTab[49721]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:144
	if len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:144
		_go_fuzz_dep_.CoverTab[49771]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:144
		return s[0] == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:144
		// _ = "end of CoverTab[49771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:144
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:144
		_go_fuzz_dep_.CoverTab[49772]++
																n := 1
																s = s[1:]

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:149
		if len(s) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:149
			_go_fuzz_dep_.CoverTab[49775]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:149
			return kind == numFloat
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:149
			// _ = "end of CoverTab[49775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:149
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:149
			_go_fuzz_dep_.CoverTab[49776]++
																	return number{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:150
			// _ = "end of CoverTab[49776]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:151
			_go_fuzz_dep_.CoverTab[49777]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:151
			// _ = "end of CoverTab[49777]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:151
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:151
		// _ = "end of CoverTab[49772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:151
		_go_fuzz_dep_.CoverTab[49773]++
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:152
			_go_fuzz_dep_.CoverTab[49778]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:152
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:152
			// _ = "end of CoverTab[49778]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:152
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:152
			_go_fuzz_dep_.CoverTab[49779]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:152
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:152
			// _ = "end of CoverTab[49779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:152
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:152
			_go_fuzz_dep_.CoverTab[49780]++
																	s = s[1:]
																	n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:154
			// _ = "end of CoverTab[49780]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:155
		// _ = "end of CoverTab[49773]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:155
		_go_fuzz_dep_.CoverTab[49774]++
																size += n
																kind = numFloat
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:157
		// _ = "end of CoverTab[49774]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:158
		_go_fuzz_dep_.CoverTab[49781]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:158
		// _ = "end of CoverTab[49781]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:158
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:158
	// _ = "end of CoverTab[49721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:158
	_go_fuzz_dep_.CoverTab[49722]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
	if len(s) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
		_go_fuzz_dep_.CoverTab[49782]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
		return (s[0] == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
			_go_fuzz_dep_.CoverTab[49783]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
			return s[0] == 'E'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
			// _ = "end of CoverTab[49783]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
		// _ = "end of CoverTab[49782]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:161
		_go_fuzz_dep_.CoverTab[49784]++
																kind = numFloat
																s = s[1:]
																n := 1
																if s[0] == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:165
			_go_fuzz_dep_.CoverTab[49787]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:165
			return s[0] == '-'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:165
			// _ = "end of CoverTab[49787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:165
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:165
			_go_fuzz_dep_.CoverTab[49788]++
																	s = s[1:]
																	n++
																	if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:168
				_go_fuzz_dep_.CoverTab[49789]++
																		return number{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:169
				// _ = "end of CoverTab[49789]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:170
				_go_fuzz_dep_.CoverTab[49790]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:170
				// _ = "end of CoverTab[49790]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:170
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:170
			// _ = "end of CoverTab[49788]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:171
			_go_fuzz_dep_.CoverTab[49791]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:171
			// _ = "end of CoverTab[49791]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:171
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:171
		// _ = "end of CoverTab[49784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:171
		_go_fuzz_dep_.CoverTab[49785]++
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:172
			_go_fuzz_dep_.CoverTab[49792]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:172
			return '0' <= s[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:172
			// _ = "end of CoverTab[49792]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:172
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:172
			_go_fuzz_dep_.CoverTab[49793]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:172
			return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:172
			// _ = "end of CoverTab[49793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:172
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:172
			_go_fuzz_dep_.CoverTab[49794]++
																	s = s[1:]
																	n++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:174
			// _ = "end of CoverTab[49794]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:175
		// _ = "end of CoverTab[49785]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:175
		_go_fuzz_dep_.CoverTab[49786]++
																size += n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:176
		// _ = "end of CoverTab[49786]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:177
		_go_fuzz_dep_.CoverTab[49795]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:177
		// _ = "end of CoverTab[49795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:177
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:177
	// _ = "end of CoverTab[49722]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:177
	_go_fuzz_dep_.CoverTab[49723]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
	if len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
		_go_fuzz_dep_.CoverTab[49796]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
		return (s[0] == 'f' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
			_go_fuzz_dep_.CoverTab[49797]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
			return s[0] == 'F'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
			// _ = "end of CoverTab[49797]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
		// _ = "end of CoverTab[49796]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:180
		_go_fuzz_dep_.CoverTab[49798]++
																kind = numFloat
																s = s[1:]
																size++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:183
		// _ = "end of CoverTab[49798]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:184
		_go_fuzz_dep_.CoverTab[49799]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:184
		// _ = "end of CoverTab[49799]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:184
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:184
	// _ = "end of CoverTab[49723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:184
	_go_fuzz_dep_.CoverTab[49724]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:187
	if len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:187
		_go_fuzz_dep_.CoverTab[49800]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:187
		return !isDelim(s[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:187
		// _ = "end of CoverTab[49800]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:187
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:187
		_go_fuzz_dep_.CoverTab[49801]++
																return number{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:188
		// _ = "end of CoverTab[49801]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:189
		_go_fuzz_dep_.CoverTab[49802]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:189
		// _ = "end of CoverTab[49802]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:189
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:189
	// _ = "end of CoverTab[49724]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:189
	_go_fuzz_dep_.CoverTab[49725]++

															return number{kind: kind, neg: neg, size: size}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:191
	// _ = "end of CoverTab[49725]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:192
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_number.go:192
var _ = _go_fuzz_dep_.CoverTab
