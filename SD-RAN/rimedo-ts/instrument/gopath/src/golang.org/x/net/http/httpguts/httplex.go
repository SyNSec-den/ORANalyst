// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:5
package httpguts

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:5
)

import (
	"net"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/idna"
)

var isTokenTable = [127]bool{
	'!':	true,
	'#':	true,
	'$':	true,
	'%':	true,
	'&':	true,
	'\'':	true,
	'*':	true,
	'+':	true,
	'-':	true,
	'.':	true,
	'0':	true,
	'1':	true,
	'2':	true,
	'3':	true,
	'4':	true,
	'5':	true,
	'6':	true,
	'7':	true,
	'8':	true,
	'9':	true,
	'A':	true,
	'B':	true,
	'C':	true,
	'D':	true,
	'E':	true,
	'F':	true,
	'G':	true,
	'H':	true,
	'I':	true,
	'J':	true,
	'K':	true,
	'L':	true,
	'M':	true,
	'N':	true,
	'O':	true,
	'P':	true,
	'Q':	true,
	'R':	true,
	'S':	true,
	'T':	true,
	'U':	true,
	'W':	true,
	'V':	true,
	'X':	true,
	'Y':	true,
	'Z':	true,
	'^':	true,
	'_':	true,
	'`':	true,
	'a':	true,
	'b':	true,
	'c':	true,
	'd':	true,
	'e':	true,
	'f':	true,
	'g':	true,
	'h':	true,
	'i':	true,
	'j':	true,
	'k':	true,
	'l':	true,
	'm':	true,
	'n':	true,
	'o':	true,
	'p':	true,
	'q':	true,
	'r':	true,
	's':	true,
	't':	true,
	'u':	true,
	'v':	true,
	'w':	true,
	'x':	true,
	'y':	true,
	'z':	true,
	'|':	true,
	'~':	true,
}

func IsTokenRune(r rune) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:95
	_go_fuzz_dep_.CoverTab[71789]++
											i := int(r)
											return i < len(isTokenTable) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:97
		_go_fuzz_dep_.CoverTab[71790]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:97
		return isTokenTable[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:97
		// _ = "end of CoverTab[71790]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:97
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:97
	// _ = "end of CoverTab[71789]"
}

func isNotToken(r rune) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:100
	_go_fuzz_dep_.CoverTab[71791]++
												return !IsTokenRune(r)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:101
	// _ = "end of CoverTab[71791]"
}

// HeaderValuesContainsToken reports whether any string in values
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:104
// contains the provided token, ASCII case-insensitively.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:106
func HeaderValuesContainsToken(values []string, token string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:106
	_go_fuzz_dep_.CoverTab[71792]++
												for _, v := range values {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:107
		_go_fuzz_dep_.CoverTab[71794]++
													if headerValueContainsToken(v, token) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:108
			_go_fuzz_dep_.CoverTab[71795]++
														return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:109
			// _ = "end of CoverTab[71795]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:110
			_go_fuzz_dep_.CoverTab[71796]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:110
			// _ = "end of CoverTab[71796]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:110
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:110
		// _ = "end of CoverTab[71794]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:111
	// _ = "end of CoverTab[71792]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:111
	_go_fuzz_dep_.CoverTab[71793]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:112
	// _ = "end of CoverTab[71793]"
}

// isOWS reports whether b is an optional whitespace byte, as defined
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:115
// by RFC 7230 section 3.2.3.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:117
func isOWS(b byte) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:117
	_go_fuzz_dep_.CoverTab[71797]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:117
	return b == ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:117
		_go_fuzz_dep_.CoverTab[71798]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:117
		return b == '\t'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:117
		// _ = "end of CoverTab[71798]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:117
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:117
	// _ = "end of CoverTab[71797]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:117
}

// trimOWS returns x with all optional whitespace removes from the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:119
// beginning and end.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:121
func trimOWS(x string) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:121
	_go_fuzz_dep_.CoverTab[71799]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:126
	for len(x) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:126
		_go_fuzz_dep_.CoverTab[71802]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:126
		return isOWS(x[0])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:126
		// _ = "end of CoverTab[71802]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:126
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:126
		_go_fuzz_dep_.CoverTab[71803]++
													x = x[1:]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:127
		// _ = "end of CoverTab[71803]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:128
	// _ = "end of CoverTab[71799]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:128
	_go_fuzz_dep_.CoverTab[71800]++
												for len(x) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:129
		_go_fuzz_dep_.CoverTab[71804]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:129
		return isOWS(x[len(x)-1])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:129
		// _ = "end of CoverTab[71804]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:129
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:129
		_go_fuzz_dep_.CoverTab[71805]++
													x = x[:len(x)-1]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:130
		// _ = "end of CoverTab[71805]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:131
	// _ = "end of CoverTab[71800]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:131
	_go_fuzz_dep_.CoverTab[71801]++
												return x
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:132
	// _ = "end of CoverTab[71801]"
}

// headerValueContainsToken reports whether v (assumed to be a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:135
// 0#element, in the ABNF extension described in RFC 7230 section 7)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:135
// contains token amongst its comma-separated tokens, ASCII
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:135
// case-insensitively.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:139
func headerValueContainsToken(v string, token string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:139
	_go_fuzz_dep_.CoverTab[71806]++
												for comma := strings.IndexByte(v, ','); comma != -1; comma = strings.IndexByte(v, ',') {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:140
		_go_fuzz_dep_.CoverTab[71808]++
													if tokenEqual(trimOWS(v[:comma]), token) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:141
			_go_fuzz_dep_.CoverTab[71810]++
														return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:142
			// _ = "end of CoverTab[71810]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:143
			_go_fuzz_dep_.CoverTab[71811]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:143
			// _ = "end of CoverTab[71811]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:143
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:143
		// _ = "end of CoverTab[71808]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:143
		_go_fuzz_dep_.CoverTab[71809]++
													v = v[comma+1:]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:144
		// _ = "end of CoverTab[71809]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:145
	// _ = "end of CoverTab[71806]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:145
	_go_fuzz_dep_.CoverTab[71807]++
												return tokenEqual(trimOWS(v), token)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:146
	// _ = "end of CoverTab[71807]"
}

// lowerASCII returns the ASCII lowercase version of b.
func lowerASCII(b byte) byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:150
	_go_fuzz_dep_.CoverTab[71812]++
												if 'A' <= b && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:151
		_go_fuzz_dep_.CoverTab[71814]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:151
		return b <= 'Z'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:151
		// _ = "end of CoverTab[71814]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:151
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:151
		_go_fuzz_dep_.CoverTab[71815]++
													return b + ('a' - 'A')
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:152
		// _ = "end of CoverTab[71815]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:153
		_go_fuzz_dep_.CoverTab[71816]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:153
		// _ = "end of CoverTab[71816]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:153
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:153
	// _ = "end of CoverTab[71812]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:153
	_go_fuzz_dep_.CoverTab[71813]++
												return b
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:154
	// _ = "end of CoverTab[71813]"
}

// tokenEqual reports whether t1 and t2 are equal, ASCII case-insensitively.
func tokenEqual(t1, t2 string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:158
	_go_fuzz_dep_.CoverTab[71817]++
												if len(t1) != len(t2) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:159
		_go_fuzz_dep_.CoverTab[71820]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:160
		// _ = "end of CoverTab[71820]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:161
		_go_fuzz_dep_.CoverTab[71821]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:161
		// _ = "end of CoverTab[71821]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:161
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:161
	// _ = "end of CoverTab[71817]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:161
	_go_fuzz_dep_.CoverTab[71818]++
												for i, b := range t1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:162
		_go_fuzz_dep_.CoverTab[71822]++
													if b >= utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:163
			_go_fuzz_dep_.CoverTab[71824]++

														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:165
			// _ = "end of CoverTab[71824]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:166
			_go_fuzz_dep_.CoverTab[71825]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:166
			// _ = "end of CoverTab[71825]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:166
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:166
		// _ = "end of CoverTab[71822]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:166
		_go_fuzz_dep_.CoverTab[71823]++
													if lowerASCII(byte(b)) != lowerASCII(t2[i]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:167
			_go_fuzz_dep_.CoverTab[71826]++
														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:168
			// _ = "end of CoverTab[71826]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:169
			_go_fuzz_dep_.CoverTab[71827]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:169
			// _ = "end of CoverTab[71827]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:169
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:169
		// _ = "end of CoverTab[71823]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:170
	// _ = "end of CoverTab[71818]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:170
	_go_fuzz_dep_.CoverTab[71819]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:171
	// _ = "end of CoverTab[71819]"
}

// isLWS reports whether b is linear white space, according
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:174
// to http://www.w3.org/Protocols/rfc2616/rfc2616-sec2.html#sec2.2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:174
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:174
//	LWS            = [CRLF] 1*( SP | HT )
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:178
func isLWS(b byte) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:178
	_go_fuzz_dep_.CoverTab[71828]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:178
	return b == ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:178
		_go_fuzz_dep_.CoverTab[71829]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:178
		return b == '\t'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:178
		// _ = "end of CoverTab[71829]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:178
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:178
	// _ = "end of CoverTab[71828]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:178
}

// isCTL reports whether b is a control byte, according
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:180
// to http://www.w3.org/Protocols/rfc2616/rfc2616-sec2.html#sec2.2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:180
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:180
//	CTL            = <any US-ASCII control character
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:180
//	                 (octets 0 - 31) and DEL (127)>
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:185
func isCTL(b byte) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:185
	_go_fuzz_dep_.CoverTab[71830]++
												const del = 0x7f	// a CTL
												return b < ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:187
		_go_fuzz_dep_.CoverTab[71831]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:187
		return b == del
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:187
		// _ = "end of CoverTab[71831]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:187
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:187
	// _ = "end of CoverTab[71830]"
}

// ValidHeaderFieldName reports whether v is a valid HTTP/1.x header name.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
// HTTP/2 imposes the additional restriction that uppercase ASCII
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
// letters are not allowed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
// RFC 7230 says:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
//	header-field   = field-name ":" OWS field-value OWS
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
//	field-name     = token
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
//	token          = 1*tchar
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
//	tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." /
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:190
//	        "^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:201
func ValidHeaderFieldName(v string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:201
	_go_fuzz_dep_.CoverTab[71832]++
												if len(v) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:202
		_go_fuzz_dep_.CoverTab[71835]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:203
		// _ = "end of CoverTab[71835]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:204
		_go_fuzz_dep_.CoverTab[71836]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:204
		// _ = "end of CoverTab[71836]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:204
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:204
	// _ = "end of CoverTab[71832]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:204
	_go_fuzz_dep_.CoverTab[71833]++
												for _, r := range v {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:205
		_go_fuzz_dep_.CoverTab[71837]++
													if !IsTokenRune(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:206
			_go_fuzz_dep_.CoverTab[71838]++
														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:207
			// _ = "end of CoverTab[71838]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:208
			_go_fuzz_dep_.CoverTab[71839]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:208
			// _ = "end of CoverTab[71839]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:208
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:208
		// _ = "end of CoverTab[71837]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:209
	// _ = "end of CoverTab[71833]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:209
	_go_fuzz_dep_.CoverTab[71834]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:210
	// _ = "end of CoverTab[71834]"
}

// ValidHostHeader reports whether h is a valid host header.
func ValidHostHeader(h string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:214
	_go_fuzz_dep_.CoverTab[71840]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:226
	for i := 0; i < len(h); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:226
		_go_fuzz_dep_.CoverTab[71842]++
													if !validHostByte[h[i]] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:227
			_go_fuzz_dep_.CoverTab[71843]++
														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:228
			// _ = "end of CoverTab[71843]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:229
			_go_fuzz_dep_.CoverTab[71844]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:229
			// _ = "end of CoverTab[71844]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:229
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:229
		// _ = "end of CoverTab[71842]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:230
	// _ = "end of CoverTab[71840]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:230
	_go_fuzz_dep_.CoverTab[71841]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:231
	// _ = "end of CoverTab[71841]"
}

// See the validHostHeader comment.
var validHostByte = [256]bool{
	'0':	true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true,
	'8':	true, '9': true,

	'a':	true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true, 'h': true,
	'i':	true, 'j': true, 'k': true, 'l': true, 'm': true, 'n': true, 'o': true, 'p': true,
	'q':	true, 'r': true, 's': true, 't': true, 'u': true, 'v': true, 'w': true, 'x': true,
	'y':	true, 'z': true,

	'A':	true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true, 'H': true,
	'I':	true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true, 'O': true, 'P': true,
	'Q':	true, 'R': true, 'S': true, 'T': true, 'U': true, 'V': true, 'W': true, 'X': true,
	'Y':	true, 'Z': true,

	'!':	true,
	'$':	true,
	'%':	true,
	'&':	true,
	'(':	true,
	')':	true,
	'*':	true,
	'+':	true,
	',':	true,
	'-':	true,
	'.':	true,
	':':	true,
	';':	true,
	'=':	true,
	'[':	true,
	'\'':	true,
	']':	true,
	'_':	true,
	'~':	true,
}

// ValidHeaderFieldValue reports whether v is a valid "field-value" according to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2 :
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	message-header = field-name ":" [ field-value ]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	field-value    = *( field-content | LWS )
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	field-content  = <the OCTETs making up the field-value
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	                 and consisting of either *TEXT or combinations
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	                 of token, separators, and quoted-string>
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec2.html#sec2.2 :
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	TEXT           = <any OCTET except CTLs,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	                  but including LWS>
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	LWS            = [CRLF] 1*( SP | HT )
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	CTL            = <any US-ASCII control character
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	                 (octets 0 - 31) and DEL (127)>
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// RFC 7230 says:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	field-value    = *( field-content / obs-fold )
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	obj-fold       =  N/A to http2, and deprecated
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	field-content  = field-vchar [ 1*( SP / HTAB ) field-vchar ]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	field-vchar    = VCHAR / obs-text
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	obs-text       = %x80-FF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//	VCHAR          = "any visible [USASCII] character"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// http2 further says: "Similarly, HTTP/2 allows header field values
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// that are not valid. While most of the values that can be encoded
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// will not alter header field parsing, carriage return (CR, ASCII
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// 0xd), line feed (LF, ASCII 0xa), and the zero character (NUL, ASCII
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// 0x0) might be exploited by an attacker if they are translated
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// verbatim. Any request or response that contains a character not
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// permitted in a header field value MUST be treated as malformed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// (Section 8.1.2.6). Valid characters are defined by the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// field-content ABNF rule in Section 3.2 of [RFC7230]."
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// This function does not (yet?) properly handle the rejection of
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:270
// strings that begin or end with SP or HTAB.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:308
func ValidHeaderFieldValue(v string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:308
	_go_fuzz_dep_.CoverTab[71845]++
												for i := 0; i < len(v); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:309
		_go_fuzz_dep_.CoverTab[71847]++
													b := v[i]
													if isCTL(b) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:311
			_go_fuzz_dep_.CoverTab[71848]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:311
			return !isLWS(b)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:311
			// _ = "end of CoverTab[71848]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:311
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:311
			_go_fuzz_dep_.CoverTab[71849]++
														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:312
			// _ = "end of CoverTab[71849]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:313
			_go_fuzz_dep_.CoverTab[71850]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:313
			// _ = "end of CoverTab[71850]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:313
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:313
		// _ = "end of CoverTab[71847]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:314
	// _ = "end of CoverTab[71845]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:314
	_go_fuzz_dep_.CoverTab[71846]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:315
	// _ = "end of CoverTab[71846]"
}

func isASCII(s string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:318
	_go_fuzz_dep_.CoverTab[71851]++
												for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:319
		_go_fuzz_dep_.CoverTab[71853]++
													if s[i] >= utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:320
			_go_fuzz_dep_.CoverTab[71854]++
														return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:321
			// _ = "end of CoverTab[71854]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:322
			_go_fuzz_dep_.CoverTab[71855]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:322
			// _ = "end of CoverTab[71855]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:322
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:322
		// _ = "end of CoverTab[71853]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:323
	// _ = "end of CoverTab[71851]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:323
	_go_fuzz_dep_.CoverTab[71852]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:324
	// _ = "end of CoverTab[71852]"
}

// PunycodeHostPort returns the IDNA Punycode version
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:327
// of the provided "host" or "host:port" string.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:329
func PunycodeHostPort(v string) (string, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:329
	_go_fuzz_dep_.CoverTab[71856]++
												if isASCII(v) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:330
		_go_fuzz_dep_.CoverTab[71861]++
													return v, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:331
		// _ = "end of CoverTab[71861]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:332
		_go_fuzz_dep_.CoverTab[71862]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:332
		// _ = "end of CoverTab[71862]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:332
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:332
	// _ = "end of CoverTab[71856]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:332
	_go_fuzz_dep_.CoverTab[71857]++

												host, port, err := net.SplitHostPort(v)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:335
		_go_fuzz_dep_.CoverTab[71863]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:339
		host = v
													port = ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:340
		// _ = "end of CoverTab[71863]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:341
		_go_fuzz_dep_.CoverTab[71864]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:341
		// _ = "end of CoverTab[71864]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:341
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:341
	// _ = "end of CoverTab[71857]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:341
	_go_fuzz_dep_.CoverTab[71858]++
												host, err = idna.ToASCII(host)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:343
		_go_fuzz_dep_.CoverTab[71865]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:346
		return "", err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:346
		// _ = "end of CoverTab[71865]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:347
		_go_fuzz_dep_.CoverTab[71866]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:347
		// _ = "end of CoverTab[71866]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:347
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:347
	// _ = "end of CoverTab[71858]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:347
	_go_fuzz_dep_.CoverTab[71859]++
												if port == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:348
		_go_fuzz_dep_.CoverTab[71867]++
													return host, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:349
		// _ = "end of CoverTab[71867]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:350
		_go_fuzz_dep_.CoverTab[71868]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:350
		// _ = "end of CoverTab[71868]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:350
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:350
	// _ = "end of CoverTab[71859]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:350
	_go_fuzz_dep_.CoverTab[71860]++
												return net.JoinHostPort(host, port), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:351
	// _ = "end of CoverTab[71860]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:352
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http/httpguts/httplex.go:352
var _ = _go_fuzz_dep_.CoverTab
