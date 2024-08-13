// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:5
package httpguts

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:5
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
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:95
	_go_fuzz_dep_.CoverTab[34914]++
										i := int(r)
										return i < len(isTokenTable) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:97
		_go_fuzz_dep_.CoverTab[34915]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:97
		return isTokenTable[i]
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:97
		// _ = "end of CoverTab[34915]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:97
	}()
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:97
	// _ = "end of CoverTab[34914]"
}

func isNotToken(r rune) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:100
	_go_fuzz_dep_.CoverTab[34916]++
										return !IsTokenRune(r)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:101
	// _ = "end of CoverTab[34916]"
}

// HeaderValuesContainsToken reports whether any string in values
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:104
// contains the provided token, ASCII case-insensitively.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:106
func HeaderValuesContainsToken(values []string, token string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:106
	_go_fuzz_dep_.CoverTab[34917]++
										for _, v := range values {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:107
		_go_fuzz_dep_.CoverTab[34919]++
											if headerValueContainsToken(v, token) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:108
			_go_fuzz_dep_.CoverTab[34920]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:109
			// _ = "end of CoverTab[34920]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:110
			_go_fuzz_dep_.CoverTab[34921]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:110
			// _ = "end of CoverTab[34921]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:110
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:110
		// _ = "end of CoverTab[34919]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:111
	// _ = "end of CoverTab[34917]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:111
	_go_fuzz_dep_.CoverTab[34918]++
										return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:112
	// _ = "end of CoverTab[34918]"
}

// isOWS reports whether b is an optional whitespace byte, as defined
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:115
// by RFC 7230 section 3.2.3.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:117
func isOWS(b byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:117
	_go_fuzz_dep_.CoverTab[34922]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:117
	return b == ' ' || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:117
		_go_fuzz_dep_.CoverTab[34923]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:117
		return b == '\t'
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:117
		// _ = "end of CoverTab[34923]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:117
	}()
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:117
	// _ = "end of CoverTab[34922]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:117
}

// trimOWS returns x with all optional whitespace removes from the
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:119
// beginning and end.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:121
func trimOWS(x string) string {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:121
	_go_fuzz_dep_.CoverTab[34924]++

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:126
	for len(x) > 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:126
		_go_fuzz_dep_.CoverTab[34927]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:126
		return isOWS(x[0])
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:126
		// _ = "end of CoverTab[34927]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:126
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:126
		_go_fuzz_dep_.CoverTab[34928]++
											x = x[1:]
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:127
		// _ = "end of CoverTab[34928]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:128
	// _ = "end of CoverTab[34924]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:128
	_go_fuzz_dep_.CoverTab[34925]++
										for len(x) > 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:129
		_go_fuzz_dep_.CoverTab[34929]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:129
		return isOWS(x[len(x)-1])
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:129
		// _ = "end of CoverTab[34929]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:129
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:129
		_go_fuzz_dep_.CoverTab[34930]++
											x = x[:len(x)-1]
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:130
		// _ = "end of CoverTab[34930]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:131
	// _ = "end of CoverTab[34925]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:131
	_go_fuzz_dep_.CoverTab[34926]++
										return x
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:132
	// _ = "end of CoverTab[34926]"
}

// headerValueContainsToken reports whether v (assumed to be a
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:135
// 0#element, in the ABNF extension described in RFC 7230 section 7)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:135
// contains token amongst its comma-separated tokens, ASCII
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:135
// case-insensitively.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:139
func headerValueContainsToken(v string, token string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:139
	_go_fuzz_dep_.CoverTab[34931]++
										for comma := strings.IndexByte(v, ','); comma != -1; comma = strings.IndexByte(v, ',') {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:140
		_go_fuzz_dep_.CoverTab[34933]++
											if tokenEqual(trimOWS(v[:comma]), token) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:141
			_go_fuzz_dep_.CoverTab[34935]++
												return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:142
			// _ = "end of CoverTab[34935]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:143
			_go_fuzz_dep_.CoverTab[34936]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:143
			// _ = "end of CoverTab[34936]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:143
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:143
		// _ = "end of CoverTab[34933]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:143
		_go_fuzz_dep_.CoverTab[34934]++
											v = v[comma+1:]
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:144
		// _ = "end of CoverTab[34934]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:145
	// _ = "end of CoverTab[34931]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:145
	_go_fuzz_dep_.CoverTab[34932]++
										return tokenEqual(trimOWS(v), token)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:146
	// _ = "end of CoverTab[34932]"
}

// lowerASCII returns the ASCII lowercase version of b.
func lowerASCII(b byte) byte {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:150
	_go_fuzz_dep_.CoverTab[34937]++
										if 'A' <= b && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:151
		_go_fuzz_dep_.CoverTab[34939]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:151
		return b <= 'Z'
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:151
		// _ = "end of CoverTab[34939]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:151
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:151
		_go_fuzz_dep_.CoverTab[34940]++
											return b + ('a' - 'A')
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:152
		// _ = "end of CoverTab[34940]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:153
		_go_fuzz_dep_.CoverTab[34941]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:153
		// _ = "end of CoverTab[34941]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:153
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:153
	// _ = "end of CoverTab[34937]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:153
	_go_fuzz_dep_.CoverTab[34938]++
										return b
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:154
	// _ = "end of CoverTab[34938]"
}

// tokenEqual reports whether t1 and t2 are equal, ASCII case-insensitively.
func tokenEqual(t1, t2 string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:158
	_go_fuzz_dep_.CoverTab[34942]++
										if len(t1) != len(t2) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:159
		_go_fuzz_dep_.CoverTab[34945]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:160
		// _ = "end of CoverTab[34945]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:161
		_go_fuzz_dep_.CoverTab[34946]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:161
		// _ = "end of CoverTab[34946]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:161
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:161
	// _ = "end of CoverTab[34942]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:161
	_go_fuzz_dep_.CoverTab[34943]++
										for i, b := range t1 {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:162
		_go_fuzz_dep_.CoverTab[34947]++
											if b >= utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:163
			_go_fuzz_dep_.CoverTab[34949]++

												return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:165
			// _ = "end of CoverTab[34949]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:166
			_go_fuzz_dep_.CoverTab[34950]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:166
			// _ = "end of CoverTab[34950]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:166
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:166
		// _ = "end of CoverTab[34947]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:166
		_go_fuzz_dep_.CoverTab[34948]++
											if lowerASCII(byte(b)) != lowerASCII(t2[i]) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:167
			_go_fuzz_dep_.CoverTab[34951]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:168
			// _ = "end of CoverTab[34951]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:169
			_go_fuzz_dep_.CoverTab[34952]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:169
			// _ = "end of CoverTab[34952]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:169
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:169
		// _ = "end of CoverTab[34948]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:170
	// _ = "end of CoverTab[34943]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:170
	_go_fuzz_dep_.CoverTab[34944]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:171
	// _ = "end of CoverTab[34944]"
}

// isLWS reports whether b is linear white space, according
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:174
// to http://www.w3.org/Protocols/rfc2616/rfc2616-sec2.html#sec2.2
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:174
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:174
//	LWS            = [CRLF] 1*( SP | HT )
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:178
func isLWS(b byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:178
	_go_fuzz_dep_.CoverTab[34953]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:178
	return b == ' ' || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:178
		_go_fuzz_dep_.CoverTab[34954]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:178
		return b == '\t'
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:178
		// _ = "end of CoverTab[34954]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:178
	}()
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:178
	// _ = "end of CoverTab[34953]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:178
}

// isCTL reports whether b is a control byte, according
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:180
// to http://www.w3.org/Protocols/rfc2616/rfc2616-sec2.html#sec2.2
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:180
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:180
//	CTL            = <any US-ASCII control character
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:180
//	                 (octets 0 - 31) and DEL (127)>
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:185
func isCTL(b byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:185
	_go_fuzz_dep_.CoverTab[34955]++
										const del = 0x7f	// a CTL
										return b < ' ' || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:187
		_go_fuzz_dep_.CoverTab[34956]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:187
		return b == del
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:187
		// _ = "end of CoverTab[34956]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:187
	}()
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:187
	// _ = "end of CoverTab[34955]"
}

// ValidHeaderFieldName reports whether v is a valid HTTP/1.x header name.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
// HTTP/2 imposes the additional restriction that uppercase ASCII
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
// letters are not allowed.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
// RFC 7230 says:
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
//	header-field   = field-name ":" OWS field-value OWS
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
//	field-name     = token
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
//	token          = 1*tchar
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
//	tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." /
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:190
//	        "^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:201
func ValidHeaderFieldName(v string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:201
	_go_fuzz_dep_.CoverTab[34957]++
										if len(v) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:202
		_go_fuzz_dep_.CoverTab[34960]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:203
		// _ = "end of CoverTab[34960]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:204
		_go_fuzz_dep_.CoverTab[34961]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:204
		// _ = "end of CoverTab[34961]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:204
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:204
	// _ = "end of CoverTab[34957]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:204
	_go_fuzz_dep_.CoverTab[34958]++
										for _, r := range v {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:205
		_go_fuzz_dep_.CoverTab[34962]++
											if !IsTokenRune(r) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:206
			_go_fuzz_dep_.CoverTab[34963]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:207
			// _ = "end of CoverTab[34963]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:208
			_go_fuzz_dep_.CoverTab[34964]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:208
			// _ = "end of CoverTab[34964]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:208
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:208
		// _ = "end of CoverTab[34962]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:209
	// _ = "end of CoverTab[34958]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:209
	_go_fuzz_dep_.CoverTab[34959]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:210
	// _ = "end of CoverTab[34959]"
}

// ValidHostHeader reports whether h is a valid host header.
func ValidHostHeader(h string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:214
	_go_fuzz_dep_.CoverTab[34965]++

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:226
	for i := 0; i < len(h); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:226
		_go_fuzz_dep_.CoverTab[34967]++
											if !validHostByte[h[i]] {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:227
			_go_fuzz_dep_.CoverTab[34968]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:228
			// _ = "end of CoverTab[34968]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:229
			_go_fuzz_dep_.CoverTab[34969]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:229
			// _ = "end of CoverTab[34969]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:229
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:229
		// _ = "end of CoverTab[34967]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:230
	// _ = "end of CoverTab[34965]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:230
	_go_fuzz_dep_.CoverTab[34966]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:231
	// _ = "end of CoverTab[34966]"
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
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2 :
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	message-header = field-name ":" [ field-value ]
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	field-value    = *( field-content | LWS )
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	field-content  = <the OCTETs making up the field-value
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	                 and consisting of either *TEXT or combinations
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	                 of token, separators, and quoted-string>
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec2.html#sec2.2 :
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	TEXT           = <any OCTET except CTLs,
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	                  but including LWS>
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	LWS            = [CRLF] 1*( SP | HT )
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	CTL            = <any US-ASCII control character
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	                 (octets 0 - 31) and DEL (127)>
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// RFC 7230 says:
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	field-value    = *( field-content / obs-fold )
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	obj-fold       =  N/A to http2, and deprecated
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	field-content  = field-vchar [ 1*( SP / HTAB ) field-vchar ]
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	field-vchar    = VCHAR / obs-text
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	obs-text       = %x80-FF
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//	VCHAR          = "any visible [USASCII] character"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// http2 further says: "Similarly, HTTP/2 allows header field values
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// that are not valid. While most of the values that can be encoded
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// will not alter header field parsing, carriage return (CR, ASCII
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// 0xd), line feed (LF, ASCII 0xa), and the zero character (NUL, ASCII
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// 0x0) might be exploited by an attacker if they are translated
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// verbatim. Any request or response that contains a character not
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// permitted in a header field value MUST be treated as malformed
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// (Section 8.1.2.6). Valid characters are defined by the
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// field-content ABNF rule in Section 3.2 of [RFC7230]."
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// This function does not (yet?) properly handle the rejection of
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:270
// strings that begin or end with SP or HTAB.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:308
func ValidHeaderFieldValue(v string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:308
	_go_fuzz_dep_.CoverTab[34970]++
										for i := 0; i < len(v); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:309
		_go_fuzz_dep_.CoverTab[34972]++
											b := v[i]
											if isCTL(b) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:311
			_go_fuzz_dep_.CoverTab[34973]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:311
			return !isLWS(b)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:311
			// _ = "end of CoverTab[34973]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:311
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:311
			_go_fuzz_dep_.CoverTab[34974]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:312
			// _ = "end of CoverTab[34974]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:313
			_go_fuzz_dep_.CoverTab[34975]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:313
			// _ = "end of CoverTab[34975]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:313
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:313
		// _ = "end of CoverTab[34972]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:314
	// _ = "end of CoverTab[34970]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:314
	_go_fuzz_dep_.CoverTab[34971]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:315
	// _ = "end of CoverTab[34971]"
}

func isASCII(s string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:318
	_go_fuzz_dep_.CoverTab[34976]++
										for i := 0; i < len(s); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:319
		_go_fuzz_dep_.CoverTab[34978]++
											if s[i] >= utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:320
			_go_fuzz_dep_.CoverTab[34979]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:321
			// _ = "end of CoverTab[34979]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:322
			_go_fuzz_dep_.CoverTab[34980]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:322
			// _ = "end of CoverTab[34980]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:322
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:322
		// _ = "end of CoverTab[34978]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:323
	// _ = "end of CoverTab[34976]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:323
	_go_fuzz_dep_.CoverTab[34977]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:324
	// _ = "end of CoverTab[34977]"
}

// PunycodeHostPort returns the IDNA Punycode version
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:327
// of the provided "host" or "host:port" string.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:329
func PunycodeHostPort(v string) (string, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:329
	_go_fuzz_dep_.CoverTab[34981]++
										if isASCII(v) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:330
		_go_fuzz_dep_.CoverTab[34986]++
											return v, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:331
		// _ = "end of CoverTab[34986]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:332
		_go_fuzz_dep_.CoverTab[34987]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:332
		// _ = "end of CoverTab[34987]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:332
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:332
	// _ = "end of CoverTab[34981]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:332
	_go_fuzz_dep_.CoverTab[34982]++

										host, port, err := net.SplitHostPort(v)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:335
		_go_fuzz_dep_.CoverTab[34988]++

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:339
		host = v
											port = ""
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:340
		// _ = "end of CoverTab[34988]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:341
		_go_fuzz_dep_.CoverTab[34989]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:341
		// _ = "end of CoverTab[34989]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:341
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:341
	// _ = "end of CoverTab[34982]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:341
	_go_fuzz_dep_.CoverTab[34983]++
										host, err = idna.ToASCII(host)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:343
		_go_fuzz_dep_.CoverTab[34990]++

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:346
		return "", err
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:346
		// _ = "end of CoverTab[34990]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:347
		_go_fuzz_dep_.CoverTab[34991]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:347
		// _ = "end of CoverTab[34991]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:347
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:347
	// _ = "end of CoverTab[34983]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:347
	_go_fuzz_dep_.CoverTab[34984]++
										if port == "" {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:348
		_go_fuzz_dep_.CoverTab[34992]++
											return host, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:349
		// _ = "end of CoverTab[34992]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:350
		_go_fuzz_dep_.CoverTab[34993]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:350
		// _ = "end of CoverTab[34993]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:350
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:350
	// _ = "end of CoverTab[34984]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:350
	_go_fuzz_dep_.CoverTab[34985]++
										return net.JoinHostPort(host, port), nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:351
	// _ = "end of CoverTab[34985]"
}

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:352
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpguts/httplex.go:352
var _ = _go_fuzz_dep_.CoverTab
