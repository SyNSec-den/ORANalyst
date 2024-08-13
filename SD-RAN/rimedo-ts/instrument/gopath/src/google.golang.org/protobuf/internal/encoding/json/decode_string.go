// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:5
package json

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:5
)

import (
	"strconv"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/strs"
)

func (d *Decoder) parseString(in []byte) (string, int, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:16
	_go_fuzz_dep_.CoverTab[65706]++
															in0 := in
															if len(in) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:18
		_go_fuzz_dep_.CoverTab[65710]++
																return "", 0, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:19
		// _ = "end of CoverTab[65710]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:20
		_go_fuzz_dep_.CoverTab[65711]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:20
		// _ = "end of CoverTab[65711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:20
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:20
	// _ = "end of CoverTab[65706]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:20
	_go_fuzz_dep_.CoverTab[65707]++
															if in[0] != '"' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:21
		_go_fuzz_dep_.CoverTab[65712]++
																return "", 0, d.newSyntaxError(d.currPos(), "invalid character %q at start of string", in[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:22
		// _ = "end of CoverTab[65712]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:23
		_go_fuzz_dep_.CoverTab[65713]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:23
		// _ = "end of CoverTab[65713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:23
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:23
	// _ = "end of CoverTab[65707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:23
	_go_fuzz_dep_.CoverTab[65708]++
															in = in[1:]
															i := indexNeedEscapeInBytes(in)
															in, out := in[i:], in[:i:i]
															for len(in) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:27
		_go_fuzz_dep_.CoverTab[65714]++
																switch r, n := utf8.DecodeRune(in); {
		case r == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:29
			_go_fuzz_dep_.CoverTab[65721]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:29
			return n == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:29
			// _ = "end of CoverTab[65721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:29
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:29
			_go_fuzz_dep_.CoverTab[65715]++
																	return "", 0, d.newSyntaxError(d.currPos(), "invalid UTF-8 in string")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:30
			// _ = "end of CoverTab[65715]"
		case r < ' ':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:31
			_go_fuzz_dep_.CoverTab[65716]++
																	return "", 0, d.newSyntaxError(d.currPos(), "invalid character %q in string", r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:32
			// _ = "end of CoverTab[65716]"
		case r == '"':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:33
			_go_fuzz_dep_.CoverTab[65717]++
																	in = in[1:]
																	n := len(in0) - len(in)
																	return string(out), n, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:36
			// _ = "end of CoverTab[65717]"
		case r == '\\':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:37
			_go_fuzz_dep_.CoverTab[65718]++
																	if len(in) < 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:38
				_go_fuzz_dep_.CoverTab[65722]++
																		return "", 0, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:39
				// _ = "end of CoverTab[65722]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:40
				_go_fuzz_dep_.CoverTab[65723]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:40
				// _ = "end of CoverTab[65723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:40
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:40
			// _ = "end of CoverTab[65718]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:40
			_go_fuzz_dep_.CoverTab[65719]++
																	switch r := in[1]; r {
			case '"', '\\', '/':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:42
				_go_fuzz_dep_.CoverTab[65724]++
																		in, out = in[2:], append(out, r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:43
				// _ = "end of CoverTab[65724]"
			case 'b':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:44
				_go_fuzz_dep_.CoverTab[65725]++
																		in, out = in[2:], append(out, '\b')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:45
				// _ = "end of CoverTab[65725]"
			case 'f':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:46
				_go_fuzz_dep_.CoverTab[65726]++
																		in, out = in[2:], append(out, '\f')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:47
				// _ = "end of CoverTab[65726]"
			case 'n':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:48
				_go_fuzz_dep_.CoverTab[65727]++
																		in, out = in[2:], append(out, '\n')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:49
				// _ = "end of CoverTab[65727]"
			case 'r':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:50
				_go_fuzz_dep_.CoverTab[65728]++
																		in, out = in[2:], append(out, '\r')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:51
				// _ = "end of CoverTab[65728]"
			case 't':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:52
				_go_fuzz_dep_.CoverTab[65729]++
																		in, out = in[2:], append(out, '\t')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:53
				// _ = "end of CoverTab[65729]"
			case 'u':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:54
				_go_fuzz_dep_.CoverTab[65730]++
																		if len(in) < 6 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:55
					_go_fuzz_dep_.CoverTab[65735]++
																			return "", 0, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:56
					// _ = "end of CoverTab[65735]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:57
					_go_fuzz_dep_.CoverTab[65736]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:57
					// _ = "end of CoverTab[65736]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:57
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:57
				// _ = "end of CoverTab[65730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:57
				_go_fuzz_dep_.CoverTab[65731]++
																		v, err := strconv.ParseUint(string(in[2:6]), 16, 16)
																		if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:59
					_go_fuzz_dep_.CoverTab[65737]++
																			return "", 0, d.newSyntaxError(d.currPos(), "invalid escape code %q in string", in[:6])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:60
					// _ = "end of CoverTab[65737]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:61
					_go_fuzz_dep_.CoverTab[65738]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:61
					// _ = "end of CoverTab[65738]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:61
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:61
				// _ = "end of CoverTab[65731]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:61
				_go_fuzz_dep_.CoverTab[65732]++
																		in = in[6:]

																		r := rune(v)
																		if utf16.IsSurrogate(r) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:65
					_go_fuzz_dep_.CoverTab[65739]++
																			if len(in) < 6 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:66
						_go_fuzz_dep_.CoverTab[65742]++
																				return "", 0, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:67
						// _ = "end of CoverTab[65742]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:68
						_go_fuzz_dep_.CoverTab[65743]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:68
						// _ = "end of CoverTab[65743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:68
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:68
					// _ = "end of CoverTab[65739]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:68
					_go_fuzz_dep_.CoverTab[65740]++
																			v, err := strconv.ParseUint(string(in[2:6]), 16, 16)
																			r = utf16.DecodeRune(r, rune(v))
																			if in[0] != '\\' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:71
						_go_fuzz_dep_.CoverTab[65744]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:71
						return in[1] != 'u'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:71
						// _ = "end of CoverTab[65744]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:71
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:71
						_go_fuzz_dep_.CoverTab[65745]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:71
						return r == unicode.ReplacementChar
																				// _ = "end of CoverTab[65745]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:72
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:72
						_go_fuzz_dep_.CoverTab[65746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:72
						return err != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:72
						// _ = "end of CoverTab[65746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:72
					}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:72
						_go_fuzz_dep_.CoverTab[65747]++
																				return "", 0, d.newSyntaxError(d.currPos(), "invalid escape code %q in string", in[:6])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:73
						// _ = "end of CoverTab[65747]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:74
						_go_fuzz_dep_.CoverTab[65748]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:74
						// _ = "end of CoverTab[65748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:74
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:74
					// _ = "end of CoverTab[65740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:74
					_go_fuzz_dep_.CoverTab[65741]++
																			in = in[6:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:75
					// _ = "end of CoverTab[65741]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:76
					_go_fuzz_dep_.CoverTab[65749]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:76
					// _ = "end of CoverTab[65749]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:76
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:76
				// _ = "end of CoverTab[65732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:76
				_go_fuzz_dep_.CoverTab[65733]++
																		out = append(out, string(r)...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:77
				// _ = "end of CoverTab[65733]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:78
				_go_fuzz_dep_.CoverTab[65734]++
																		return "", 0, d.newSyntaxError(d.currPos(), "invalid escape code %q in string", in[:2])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:79
				// _ = "end of CoverTab[65734]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:80
			// _ = "end of CoverTab[65719]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:81
			_go_fuzz_dep_.CoverTab[65720]++
																	i := indexNeedEscapeInBytes(in[n:])
																	in, out = in[n+i:], append(out, in[:n+i]...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:83
			// _ = "end of CoverTab[65720]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:84
		// _ = "end of CoverTab[65714]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:85
	// _ = "end of CoverTab[65708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:85
	_go_fuzz_dep_.CoverTab[65709]++
															return "", 0, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:86
	// _ = "end of CoverTab[65709]"
}

// indexNeedEscapeInBytes returns the index of the character that needs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:89
// escaping. If no characters need escaping, this returns the input length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:91
func indexNeedEscapeInBytes(b []byte) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:91
	_go_fuzz_dep_.CoverTab[65750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:91
	return indexNeedEscapeInString(strs.UnsafeString(b))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:91
	// _ = "end of CoverTab[65750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:91
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:91
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode_string.go:91
var _ = _go_fuzz_dep_.CoverTab
