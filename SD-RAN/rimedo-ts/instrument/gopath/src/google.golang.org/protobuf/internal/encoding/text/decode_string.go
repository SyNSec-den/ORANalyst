// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:5
package text

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:5
)

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/strs"
)

// parseStringValue parses string field token.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:18
// This differs from parseString since the text format allows
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:18
// multiple back-to-back string literals where they are semantically treated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:18
// as a single large string with all values concatenated.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:18
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:18
// E.g., `"foo" "bar" "baz"` => "foobarbaz"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:24
func (d *Decoder) parseStringValue() (Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:24
	_go_fuzz_dep_.CoverTab[49803]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:30
	in0 := d.in
	var ss []string
	for len(d.in) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:32
		_go_fuzz_dep_.CoverTab[49805]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:32
		return (d.in[0] == '"' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:32
			_go_fuzz_dep_.CoverTab[49806]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:32
			return d.in[0] == '\''
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:32
			// _ = "end of CoverTab[49806]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:32
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:32
		// _ = "end of CoverTab[49805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:32
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:32
		_go_fuzz_dep_.CoverTab[49807]++
																s, err := d.parseString()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:34
			_go_fuzz_dep_.CoverTab[49809]++
																	return Token{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:35
			// _ = "end of CoverTab[49809]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:36
			_go_fuzz_dep_.CoverTab[49810]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:36
			// _ = "end of CoverTab[49810]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:36
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:36
		// _ = "end of CoverTab[49807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:36
		_go_fuzz_dep_.CoverTab[49808]++
																ss = append(ss, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:37
		// _ = "end of CoverTab[49808]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:38
	// _ = "end of CoverTab[49803]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:38
	_go_fuzz_dep_.CoverTab[49804]++

															return Token{
		kind:	Scalar,
		attrs:	stringValue,
		pos:	len(d.orig) - len(in0),
		raw:	in0[:len(in0)-len(d.in)],
		str:	strings.Join(ss, ""),
	}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:46
	// _ = "end of CoverTab[49804]"
}

// parseString parses a string value enclosed in " or '.
func (d *Decoder) parseString() (string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:50
	_go_fuzz_dep_.CoverTab[49811]++
															in := d.in
															if len(in) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:52
		_go_fuzz_dep_.CoverTab[49814]++
																return "", ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:53
		// _ = "end of CoverTab[49814]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:54
		_go_fuzz_dep_.CoverTab[49815]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:54
		// _ = "end of CoverTab[49815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:54
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:54
	// _ = "end of CoverTab[49811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:54
	_go_fuzz_dep_.CoverTab[49812]++
															quote := in[0]
															in = in[1:]
															i := indexNeedEscapeInBytes(in)
															in, out := in[i:], in[:i:i]
															for len(in) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:59
		_go_fuzz_dep_.CoverTab[49816]++
																switch r, n := utf8.DecodeRune(in); {
		case r == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:61
			_go_fuzz_dep_.CoverTab[49823]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:61
			return n == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:61
			// _ = "end of CoverTab[49823]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:61
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:61
			_go_fuzz_dep_.CoverTab[49817]++
																	return "", d.newSyntaxError("invalid UTF-8 detected")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:62
			// _ = "end of CoverTab[49817]"
		case r == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:63
			_go_fuzz_dep_.CoverTab[49824]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:63
			return r == '\n'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:63
			// _ = "end of CoverTab[49824]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:63
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:63
			_go_fuzz_dep_.CoverTab[49818]++
																	return "", d.newSyntaxError("invalid character %q in string", r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:64
			// _ = "end of CoverTab[49818]"
		case r == rune(quote):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:65
			_go_fuzz_dep_.CoverTab[49819]++
																	in = in[1:]
																	d.consume(len(d.in) - len(in))
																	return string(out), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:68
			// _ = "end of CoverTab[49819]"
		case r == '\\':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:69
			_go_fuzz_dep_.CoverTab[49820]++
																	if len(in) < 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:70
				_go_fuzz_dep_.CoverTab[49825]++
																		return "", ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:71
				// _ = "end of CoverTab[49825]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:72
				_go_fuzz_dep_.CoverTab[49826]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:72
				// _ = "end of CoverTab[49826]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:72
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:72
			// _ = "end of CoverTab[49820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:72
			_go_fuzz_dep_.CoverTab[49821]++
																	switch r := in[1]; r {
			case '"', '\'', '\\', '?':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:74
				_go_fuzz_dep_.CoverTab[49827]++
																		in, out = in[2:], append(out, r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:75
				// _ = "end of CoverTab[49827]"
			case 'a':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:76
				_go_fuzz_dep_.CoverTab[49828]++
																		in, out = in[2:], append(out, '\a')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:77
				// _ = "end of CoverTab[49828]"
			case 'b':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:78
				_go_fuzz_dep_.CoverTab[49829]++
																		in, out = in[2:], append(out, '\b')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:79
				// _ = "end of CoverTab[49829]"
			case 'n':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:80
				_go_fuzz_dep_.CoverTab[49830]++
																		in, out = in[2:], append(out, '\n')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:81
				// _ = "end of CoverTab[49830]"
			case 'r':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:82
				_go_fuzz_dep_.CoverTab[49831]++
																		in, out = in[2:], append(out, '\r')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:83
				// _ = "end of CoverTab[49831]"
			case 't':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:84
				_go_fuzz_dep_.CoverTab[49832]++
																		in, out = in[2:], append(out, '\t')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:85
				// _ = "end of CoverTab[49832]"
			case 'v':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:86
				_go_fuzz_dep_.CoverTab[49833]++
																		in, out = in[2:], append(out, '\v')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:87
				// _ = "end of CoverTab[49833]"
			case 'f':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:88
				_go_fuzz_dep_.CoverTab[49834]++
																		in, out = in[2:], append(out, '\f')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:89
				// _ = "end of CoverTab[49834]"
			case '0', '1', '2', '3', '4', '5', '6', '7':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:90
				_go_fuzz_dep_.CoverTab[49835]++

																		n := len(in[1:]) - len(bytes.TrimLeft(in[1:], "01234567"))
																		if n > 3 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:93
					_go_fuzz_dep_.CoverTab[49847]++
																			n = 3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:94
					// _ = "end of CoverTab[49847]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:95
					_go_fuzz_dep_.CoverTab[49848]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:95
					// _ = "end of CoverTab[49848]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:95
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:95
				// _ = "end of CoverTab[49835]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:95
				_go_fuzz_dep_.CoverTab[49836]++
																		v, err := strconv.ParseUint(string(in[1:1+n]), 8, 8)
																		if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:97
					_go_fuzz_dep_.CoverTab[49849]++
																			return "", d.newSyntaxError("invalid octal escape code %q in string", in[:1+n])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:98
					// _ = "end of CoverTab[49849]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:99
					_go_fuzz_dep_.CoverTab[49850]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:99
					// _ = "end of CoverTab[49850]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:99
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:99
				// _ = "end of CoverTab[49836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:99
				_go_fuzz_dep_.CoverTab[49837]++
																		in, out = in[1+n:], append(out, byte(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:100
				// _ = "end of CoverTab[49837]"
			case 'x':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:101
				_go_fuzz_dep_.CoverTab[49838]++

																		n := len(in[2:]) - len(bytes.TrimLeft(in[2:], "0123456789abcdefABCDEF"))
																		if n > 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:104
					_go_fuzz_dep_.CoverTab[49851]++
																			n = 2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:105
					// _ = "end of CoverTab[49851]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:106
					_go_fuzz_dep_.CoverTab[49852]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:106
					// _ = "end of CoverTab[49852]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:106
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:106
				// _ = "end of CoverTab[49838]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:106
				_go_fuzz_dep_.CoverTab[49839]++
																		v, err := strconv.ParseUint(string(in[2:2+n]), 16, 8)
																		if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:108
					_go_fuzz_dep_.CoverTab[49853]++
																			return "", d.newSyntaxError("invalid hex escape code %q in string", in[:2+n])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:109
					// _ = "end of CoverTab[49853]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:110
					_go_fuzz_dep_.CoverTab[49854]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:110
					// _ = "end of CoverTab[49854]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:110
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:110
				// _ = "end of CoverTab[49839]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:110
				_go_fuzz_dep_.CoverTab[49840]++
																		in, out = in[2+n:], append(out, byte(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:111
				// _ = "end of CoverTab[49840]"
			case 'u', 'U':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:112
				_go_fuzz_dep_.CoverTab[49841]++

																		n := 6
																		if r == 'U' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:115
					_go_fuzz_dep_.CoverTab[49855]++
																			n = 10
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:116
					// _ = "end of CoverTab[49855]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:117
					_go_fuzz_dep_.CoverTab[49856]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:117
					// _ = "end of CoverTab[49856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:117
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:117
				// _ = "end of CoverTab[49841]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:117
				_go_fuzz_dep_.CoverTab[49842]++
																		if len(in) < n {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:118
					_go_fuzz_dep_.CoverTab[49857]++
																			return "", ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:119
					// _ = "end of CoverTab[49857]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:120
					_go_fuzz_dep_.CoverTab[49858]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:120
					// _ = "end of CoverTab[49858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:120
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:120
				// _ = "end of CoverTab[49842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:120
				_go_fuzz_dep_.CoverTab[49843]++
																		v, err := strconv.ParseUint(string(in[2:n]), 16, 32)
																		if utf8.MaxRune < v || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:122
					_go_fuzz_dep_.CoverTab[49859]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:122
					return err != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:122
					// _ = "end of CoverTab[49859]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:122
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:122
					_go_fuzz_dep_.CoverTab[49860]++
																			return "", d.newSyntaxError("invalid Unicode escape code %q in string", in[:n])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:123
					// _ = "end of CoverTab[49860]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:124
					_go_fuzz_dep_.CoverTab[49861]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:124
					// _ = "end of CoverTab[49861]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:124
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:124
				// _ = "end of CoverTab[49843]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:124
				_go_fuzz_dep_.CoverTab[49844]++
																		in = in[n:]

																		r := rune(v)
																		if utf16.IsSurrogate(r) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:128
					_go_fuzz_dep_.CoverTab[49862]++
																			if len(in) < 6 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:129
						_go_fuzz_dep_.CoverTab[49865]++
																				return "", ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:130
						// _ = "end of CoverTab[49865]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:131
						_go_fuzz_dep_.CoverTab[49866]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:131
						// _ = "end of CoverTab[49866]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:131
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:131
					// _ = "end of CoverTab[49862]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:131
					_go_fuzz_dep_.CoverTab[49863]++
																			v, err := strconv.ParseUint(string(in[2:6]), 16, 16)
																			r = utf16.DecodeRune(r, rune(v))
																			if in[0] != '\\' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						_go_fuzz_dep_.CoverTab[49867]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						return in[1] != 'u'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						// _ = "end of CoverTab[49867]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						_go_fuzz_dep_.CoverTab[49868]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						return r == unicode.ReplacementChar
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						// _ = "end of CoverTab[49868]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						_go_fuzz_dep_.CoverTab[49869]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						return err != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						// _ = "end of CoverTab[49869]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
					}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:134
						_go_fuzz_dep_.CoverTab[49870]++
																				return "", d.newSyntaxError("invalid Unicode escape code %q in string", in[:6])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:135
						// _ = "end of CoverTab[49870]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:136
						_go_fuzz_dep_.CoverTab[49871]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:136
						// _ = "end of CoverTab[49871]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:136
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:136
					// _ = "end of CoverTab[49863]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:136
					_go_fuzz_dep_.CoverTab[49864]++
																			in = in[6:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:137
					// _ = "end of CoverTab[49864]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:138
					_go_fuzz_dep_.CoverTab[49872]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:138
					// _ = "end of CoverTab[49872]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:138
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:138
				// _ = "end of CoverTab[49844]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:138
				_go_fuzz_dep_.CoverTab[49845]++
																		out = append(out, string(r)...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:139
				// _ = "end of CoverTab[49845]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:140
				_go_fuzz_dep_.CoverTab[49846]++
																		return "", d.newSyntaxError("invalid escape code %q in string", in[:2])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:141
				// _ = "end of CoverTab[49846]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:142
			// _ = "end of CoverTab[49821]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:143
			_go_fuzz_dep_.CoverTab[49822]++
																	i := indexNeedEscapeInBytes(in[n:])
																	in, out = in[n+i:], append(out, in[:n+i]...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:145
			// _ = "end of CoverTab[49822]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:146
		// _ = "end of CoverTab[49816]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:147
	// _ = "end of CoverTab[49812]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:147
	_go_fuzz_dep_.CoverTab[49813]++
															return "", ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:148
	// _ = "end of CoverTab[49813]"
}

// indexNeedEscapeInString returns the index of the character that needs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:151
// escaping. If no characters need escaping, this returns the input length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:153
func indexNeedEscapeInBytes(b []byte) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:153
	_go_fuzz_dep_.CoverTab[49873]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:153
	return indexNeedEscapeInString(strs.UnsafeString(b))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:153
	// _ = "end of CoverTab[49873]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:153
}

// UnmarshalString returns an unescaped string given a textproto string value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:155
// String value needs to contain single or double quotes. This is only used by
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:155
// internal/encoding/defval package for unmarshaling bytes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:158
func UnmarshalString(s string) (string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:158
	_go_fuzz_dep_.CoverTab[49874]++
															d := NewDecoder([]byte(s))
															return d.parseString()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:160
	// _ = "end of CoverTab[49874]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:161
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode_string.go:161
var _ = _go_fuzz_dep_.CoverTab
