// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/css.go:5
package template

//line /usr/local/go/src/html/template/css.go:5
import (
//line /usr/local/go/src/html/template/css.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/css.go:5
)
//line /usr/local/go/src/html/template/css.go:5
import (
//line /usr/local/go/src/html/template/css.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/css.go:5
)

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// endsWithCSSKeyword reports whether b ends with an ident that
//line /usr/local/go/src/html/template/css.go:15
// case-insensitively matches the lower-case kw.
//line /usr/local/go/src/html/template/css.go:17
func endsWithCSSKeyword(b []byte, kw string) bool {
//line /usr/local/go/src/html/template/css.go:17
	_go_fuzz_dep_.CoverTab[30732]++
							i := len(b) - len(kw)
							if i < 0 {
//line /usr/local/go/src/html/template/css.go:19
		_go_fuzz_dep_.CoverTab[30735]++

								return false
//line /usr/local/go/src/html/template/css.go:21
		// _ = "end of CoverTab[30735]"
	} else {
//line /usr/local/go/src/html/template/css.go:22
		_go_fuzz_dep_.CoverTab[30736]++
//line /usr/local/go/src/html/template/css.go:22
		// _ = "end of CoverTab[30736]"
//line /usr/local/go/src/html/template/css.go:22
	}
//line /usr/local/go/src/html/template/css.go:22
	// _ = "end of CoverTab[30732]"
//line /usr/local/go/src/html/template/css.go:22
	_go_fuzz_dep_.CoverTab[30733]++
							if i != 0 {
//line /usr/local/go/src/html/template/css.go:23
		_go_fuzz_dep_.CoverTab[30737]++
								r, _ := utf8.DecodeLastRune(b[:i])
								if isCSSNmchar(r) {
//line /usr/local/go/src/html/template/css.go:25
			_go_fuzz_dep_.CoverTab[30738]++

									return false
//line /usr/local/go/src/html/template/css.go:27
			// _ = "end of CoverTab[30738]"
		} else {
//line /usr/local/go/src/html/template/css.go:28
			_go_fuzz_dep_.CoverTab[30739]++
//line /usr/local/go/src/html/template/css.go:28
			// _ = "end of CoverTab[30739]"
//line /usr/local/go/src/html/template/css.go:28
		}
//line /usr/local/go/src/html/template/css.go:28
		// _ = "end of CoverTab[30737]"
	} else {
//line /usr/local/go/src/html/template/css.go:29
		_go_fuzz_dep_.CoverTab[30740]++
//line /usr/local/go/src/html/template/css.go:29
		// _ = "end of CoverTab[30740]"
//line /usr/local/go/src/html/template/css.go:29
	}
//line /usr/local/go/src/html/template/css.go:29
	// _ = "end of CoverTab[30733]"
//line /usr/local/go/src/html/template/css.go:29
	_go_fuzz_dep_.CoverTab[30734]++

//line /usr/local/go/src/html/template/css.go:35
	return string(bytes.ToLower(b[i:])) == kw
//line /usr/local/go/src/html/template/css.go:35
	// _ = "end of CoverTab[30734]"
}

// isCSSNmchar reports whether rune is allowed anywhere in a CSS identifier.
func isCSSNmchar(r rune) bool {
//line /usr/local/go/src/html/template/css.go:39
	_go_fuzz_dep_.CoverTab[30741]++

//line /usr/local/go/src/html/template/css.go:43
	return 'a' <= r && func() bool {
//line /usr/local/go/src/html/template/css.go:43
		_go_fuzz_dep_.CoverTab[30742]++
//line /usr/local/go/src/html/template/css.go:43
		return r <= 'z'
//line /usr/local/go/src/html/template/css.go:43
		// _ = "end of CoverTab[30742]"
//line /usr/local/go/src/html/template/css.go:43
	}() || func() bool {
//line /usr/local/go/src/html/template/css.go:43
		_go_fuzz_dep_.CoverTab[30743]++
//line /usr/local/go/src/html/template/css.go:43
		return 'A' <= r && func() bool {
									_go_fuzz_dep_.CoverTab[30744]++
//line /usr/local/go/src/html/template/css.go:44
			return r <= 'Z'
//line /usr/local/go/src/html/template/css.go:44
			// _ = "end of CoverTab[30744]"
//line /usr/local/go/src/html/template/css.go:44
		}()
//line /usr/local/go/src/html/template/css.go:44
		// _ = "end of CoverTab[30743]"
//line /usr/local/go/src/html/template/css.go:44
	}() || func() bool {
//line /usr/local/go/src/html/template/css.go:44
		_go_fuzz_dep_.CoverTab[30745]++
//line /usr/local/go/src/html/template/css.go:44
		return '0' <= r && func() bool {
									_go_fuzz_dep_.CoverTab[30746]++
//line /usr/local/go/src/html/template/css.go:45
			return r <= '9'
//line /usr/local/go/src/html/template/css.go:45
			// _ = "end of CoverTab[30746]"
//line /usr/local/go/src/html/template/css.go:45
		}()
//line /usr/local/go/src/html/template/css.go:45
		// _ = "end of CoverTab[30745]"
//line /usr/local/go/src/html/template/css.go:45
	}() || func() bool {
//line /usr/local/go/src/html/template/css.go:45
		_go_fuzz_dep_.CoverTab[30747]++
//line /usr/local/go/src/html/template/css.go:45
		return r == '-'
								// _ = "end of CoverTab[30747]"
//line /usr/local/go/src/html/template/css.go:46
	}() || func() bool {
//line /usr/local/go/src/html/template/css.go:46
		_go_fuzz_dep_.CoverTab[30748]++
//line /usr/local/go/src/html/template/css.go:46
		return r == '_'
								// _ = "end of CoverTab[30748]"
//line /usr/local/go/src/html/template/css.go:47
	}() || func() bool {
//line /usr/local/go/src/html/template/css.go:47
		_go_fuzz_dep_.CoverTab[30749]++
//line /usr/local/go/src/html/template/css.go:47
		return 0x80 <= r && func() bool {
//line /usr/local/go/src/html/template/css.go:49
			_go_fuzz_dep_.CoverTab[30750]++
//line /usr/local/go/src/html/template/css.go:49
			return r <= 0xd7ff
//line /usr/local/go/src/html/template/css.go:49
			// _ = "end of CoverTab[30750]"
//line /usr/local/go/src/html/template/css.go:49
		}()
//line /usr/local/go/src/html/template/css.go:49
		// _ = "end of CoverTab[30749]"
//line /usr/local/go/src/html/template/css.go:49
	}() || func() bool {
//line /usr/local/go/src/html/template/css.go:49
		_go_fuzz_dep_.CoverTab[30751]++
//line /usr/local/go/src/html/template/css.go:49
		return 0xe000 <= r && func() bool {
									_go_fuzz_dep_.CoverTab[30752]++
//line /usr/local/go/src/html/template/css.go:50
			return r <= 0xfffd
//line /usr/local/go/src/html/template/css.go:50
			// _ = "end of CoverTab[30752]"
//line /usr/local/go/src/html/template/css.go:50
		}()
//line /usr/local/go/src/html/template/css.go:50
		// _ = "end of CoverTab[30751]"
//line /usr/local/go/src/html/template/css.go:50
	}() || func() bool {
//line /usr/local/go/src/html/template/css.go:50
		_go_fuzz_dep_.CoverTab[30753]++
//line /usr/local/go/src/html/template/css.go:50
		return 0x10000 <= r && func() bool {
									_go_fuzz_dep_.CoverTab[30754]++
//line /usr/local/go/src/html/template/css.go:51
			return r <= 0x10ffff
//line /usr/local/go/src/html/template/css.go:51
			// _ = "end of CoverTab[30754]"
//line /usr/local/go/src/html/template/css.go:51
		}()
//line /usr/local/go/src/html/template/css.go:51
		// _ = "end of CoverTab[30753]"
//line /usr/local/go/src/html/template/css.go:51
	}()
//line /usr/local/go/src/html/template/css.go:51
	// _ = "end of CoverTab[30741]"
}

// decodeCSS decodes CSS3 escapes given a sequence of stringchars.
//line /usr/local/go/src/html/template/css.go:54
// If there is no change, it returns the input, otherwise it returns a slice
//line /usr/local/go/src/html/template/css.go:54
// backed by a new array.
//line /usr/local/go/src/html/template/css.go:54
// https://www.w3.org/TR/css3-syntax/#SUBTOK-stringchar defines stringchar.
//line /usr/local/go/src/html/template/css.go:58
func decodeCSS(s []byte) []byte {
//line /usr/local/go/src/html/template/css.go:58
	_go_fuzz_dep_.CoverTab[30755]++
							i := bytes.IndexByte(s, '\\')
							if i == -1 {
//line /usr/local/go/src/html/template/css.go:60
		_go_fuzz_dep_.CoverTab[30758]++
								return s
//line /usr/local/go/src/html/template/css.go:61
		// _ = "end of CoverTab[30758]"
	} else {
//line /usr/local/go/src/html/template/css.go:62
		_go_fuzz_dep_.CoverTab[30759]++
//line /usr/local/go/src/html/template/css.go:62
		// _ = "end of CoverTab[30759]"
//line /usr/local/go/src/html/template/css.go:62
	}
//line /usr/local/go/src/html/template/css.go:62
	// _ = "end of CoverTab[30755]"
//line /usr/local/go/src/html/template/css.go:62
	_go_fuzz_dep_.CoverTab[30756]++

//line /usr/local/go/src/html/template/css.go:66
	b := make([]byte, 0, len(s))
	for len(s) != 0 {
//line /usr/local/go/src/html/template/css.go:67
		_go_fuzz_dep_.CoverTab[30760]++
								i := bytes.IndexByte(s, '\\')
								if i == -1 {
//line /usr/local/go/src/html/template/css.go:69
			_go_fuzz_dep_.CoverTab[30763]++
									i = len(s)
//line /usr/local/go/src/html/template/css.go:70
			// _ = "end of CoverTab[30763]"
		} else {
//line /usr/local/go/src/html/template/css.go:71
			_go_fuzz_dep_.CoverTab[30764]++
//line /usr/local/go/src/html/template/css.go:71
			// _ = "end of CoverTab[30764]"
//line /usr/local/go/src/html/template/css.go:71
		}
//line /usr/local/go/src/html/template/css.go:71
		// _ = "end of CoverTab[30760]"
//line /usr/local/go/src/html/template/css.go:71
		_go_fuzz_dep_.CoverTab[30761]++
								b, s = append(b, s[:i]...), s[i:]
								if len(s) < 2 {
//line /usr/local/go/src/html/template/css.go:73
			_go_fuzz_dep_.CoverTab[30765]++
									break
//line /usr/local/go/src/html/template/css.go:74
			// _ = "end of CoverTab[30765]"
		} else {
//line /usr/local/go/src/html/template/css.go:75
			_go_fuzz_dep_.CoverTab[30766]++
//line /usr/local/go/src/html/template/css.go:75
			// _ = "end of CoverTab[30766]"
//line /usr/local/go/src/html/template/css.go:75
		}
//line /usr/local/go/src/html/template/css.go:75
		// _ = "end of CoverTab[30761]"
//line /usr/local/go/src/html/template/css.go:75
		_go_fuzz_dep_.CoverTab[30762]++

//line /usr/local/go/src/html/template/css.go:78
		if isHex(s[1]) {
//line /usr/local/go/src/html/template/css.go:78
			_go_fuzz_dep_.CoverTab[30767]++

//line /usr/local/go/src/html/template/css.go:81
			j := 2
			for j < len(s) && func() bool {
//line /usr/local/go/src/html/template/css.go:82
				_go_fuzz_dep_.CoverTab[30770]++
//line /usr/local/go/src/html/template/css.go:82
				return j < 7
//line /usr/local/go/src/html/template/css.go:82
				// _ = "end of CoverTab[30770]"
//line /usr/local/go/src/html/template/css.go:82
			}() && func() bool {
//line /usr/local/go/src/html/template/css.go:82
				_go_fuzz_dep_.CoverTab[30771]++
//line /usr/local/go/src/html/template/css.go:82
				return isHex(s[j])
//line /usr/local/go/src/html/template/css.go:82
				// _ = "end of CoverTab[30771]"
//line /usr/local/go/src/html/template/css.go:82
			}() {
//line /usr/local/go/src/html/template/css.go:82
				_go_fuzz_dep_.CoverTab[30772]++
										j++
//line /usr/local/go/src/html/template/css.go:83
				// _ = "end of CoverTab[30772]"
			}
//line /usr/local/go/src/html/template/css.go:84
			// _ = "end of CoverTab[30767]"
//line /usr/local/go/src/html/template/css.go:84
			_go_fuzz_dep_.CoverTab[30768]++
									r := hexDecode(s[1:j])
									if r > unicode.MaxRune {
//line /usr/local/go/src/html/template/css.go:86
				_go_fuzz_dep_.CoverTab[30773]++
										r, j = r/16, j-1
//line /usr/local/go/src/html/template/css.go:87
				// _ = "end of CoverTab[30773]"
			} else {
//line /usr/local/go/src/html/template/css.go:88
				_go_fuzz_dep_.CoverTab[30774]++
//line /usr/local/go/src/html/template/css.go:88
				// _ = "end of CoverTab[30774]"
//line /usr/local/go/src/html/template/css.go:88
			}
//line /usr/local/go/src/html/template/css.go:88
			// _ = "end of CoverTab[30768]"
//line /usr/local/go/src/html/template/css.go:88
			_go_fuzz_dep_.CoverTab[30769]++
									n := utf8.EncodeRune(b[len(b):cap(b)], r)

//line /usr/local/go/src/html/template/css.go:93
			b, s = b[:len(b)+n], skipCSSSpace(s[j:])
//line /usr/local/go/src/html/template/css.go:93
			// _ = "end of CoverTab[30769]"
		} else {
//line /usr/local/go/src/html/template/css.go:94
			_go_fuzz_dep_.CoverTab[30775]++

									_, n := utf8.DecodeRune(s[1:])
									b, s = append(b, s[1:1+n]...), s[1+n:]
//line /usr/local/go/src/html/template/css.go:97
			// _ = "end of CoverTab[30775]"
		}
//line /usr/local/go/src/html/template/css.go:98
		// _ = "end of CoverTab[30762]"
	}
//line /usr/local/go/src/html/template/css.go:99
	// _ = "end of CoverTab[30756]"
//line /usr/local/go/src/html/template/css.go:99
	_go_fuzz_dep_.CoverTab[30757]++
							return b
//line /usr/local/go/src/html/template/css.go:100
	// _ = "end of CoverTab[30757]"
}

// isHex reports whether the given character is a hex digit.
func isHex(c byte) bool {
//line /usr/local/go/src/html/template/css.go:104
	_go_fuzz_dep_.CoverTab[30776]++
							return '0' <= c && func() bool {
//line /usr/local/go/src/html/template/css.go:105
		_go_fuzz_dep_.CoverTab[30777]++
//line /usr/local/go/src/html/template/css.go:105
		return c <= '9'
//line /usr/local/go/src/html/template/css.go:105
		// _ = "end of CoverTab[30777]"
//line /usr/local/go/src/html/template/css.go:105
	}() || func() bool {
//line /usr/local/go/src/html/template/css.go:105
		_go_fuzz_dep_.CoverTab[30778]++
//line /usr/local/go/src/html/template/css.go:105
		return 'a' <= c && func() bool {
//line /usr/local/go/src/html/template/css.go:105
			_go_fuzz_dep_.CoverTab[30779]++
//line /usr/local/go/src/html/template/css.go:105
			return c <= 'f'
//line /usr/local/go/src/html/template/css.go:105
			// _ = "end of CoverTab[30779]"
//line /usr/local/go/src/html/template/css.go:105
		}()
//line /usr/local/go/src/html/template/css.go:105
		// _ = "end of CoverTab[30778]"
//line /usr/local/go/src/html/template/css.go:105
	}() || func() bool {
//line /usr/local/go/src/html/template/css.go:105
		_go_fuzz_dep_.CoverTab[30780]++
//line /usr/local/go/src/html/template/css.go:105
		return 'A' <= c && func() bool {
//line /usr/local/go/src/html/template/css.go:105
			_go_fuzz_dep_.CoverTab[30781]++
//line /usr/local/go/src/html/template/css.go:105
			return c <= 'F'
//line /usr/local/go/src/html/template/css.go:105
			// _ = "end of CoverTab[30781]"
//line /usr/local/go/src/html/template/css.go:105
		}()
//line /usr/local/go/src/html/template/css.go:105
		// _ = "end of CoverTab[30780]"
//line /usr/local/go/src/html/template/css.go:105
	}()
//line /usr/local/go/src/html/template/css.go:105
	// _ = "end of CoverTab[30776]"
}

// hexDecode decodes a short hex digit sequence: "10" -> 16.
func hexDecode(s []byte) rune {
//line /usr/local/go/src/html/template/css.go:109
	_go_fuzz_dep_.CoverTab[30782]++
							n := '\x00'
							for _, c := range s {
//line /usr/local/go/src/html/template/css.go:111
		_go_fuzz_dep_.CoverTab[30784]++
								n <<= 4
								switch {
		case '0' <= c && func() bool {
//line /usr/local/go/src/html/template/css.go:114
			_go_fuzz_dep_.CoverTab[30789]++
//line /usr/local/go/src/html/template/css.go:114
			return c <= '9'
//line /usr/local/go/src/html/template/css.go:114
			// _ = "end of CoverTab[30789]"
//line /usr/local/go/src/html/template/css.go:114
		}():
//line /usr/local/go/src/html/template/css.go:114
			_go_fuzz_dep_.CoverTab[30785]++
									n |= rune(c - '0')
//line /usr/local/go/src/html/template/css.go:115
			// _ = "end of CoverTab[30785]"
		case 'a' <= c && func() bool {
//line /usr/local/go/src/html/template/css.go:116
			_go_fuzz_dep_.CoverTab[30790]++
//line /usr/local/go/src/html/template/css.go:116
			return c <= 'f'
//line /usr/local/go/src/html/template/css.go:116
			// _ = "end of CoverTab[30790]"
//line /usr/local/go/src/html/template/css.go:116
		}():
//line /usr/local/go/src/html/template/css.go:116
			_go_fuzz_dep_.CoverTab[30786]++
									n |= rune(c-'a') + 10
//line /usr/local/go/src/html/template/css.go:117
			// _ = "end of CoverTab[30786]"
		case 'A' <= c && func() bool {
//line /usr/local/go/src/html/template/css.go:118
			_go_fuzz_dep_.CoverTab[30791]++
//line /usr/local/go/src/html/template/css.go:118
			return c <= 'F'
//line /usr/local/go/src/html/template/css.go:118
			// _ = "end of CoverTab[30791]"
//line /usr/local/go/src/html/template/css.go:118
		}():
//line /usr/local/go/src/html/template/css.go:118
			_go_fuzz_dep_.CoverTab[30787]++
									n |= rune(c-'A') + 10
//line /usr/local/go/src/html/template/css.go:119
			// _ = "end of CoverTab[30787]"
		default:
//line /usr/local/go/src/html/template/css.go:120
			_go_fuzz_dep_.CoverTab[30788]++
									panic(fmt.Sprintf("Bad hex digit in %q", s))
//line /usr/local/go/src/html/template/css.go:121
			// _ = "end of CoverTab[30788]"
		}
//line /usr/local/go/src/html/template/css.go:122
		// _ = "end of CoverTab[30784]"
	}
//line /usr/local/go/src/html/template/css.go:123
	// _ = "end of CoverTab[30782]"
//line /usr/local/go/src/html/template/css.go:123
	_go_fuzz_dep_.CoverTab[30783]++
							return n
//line /usr/local/go/src/html/template/css.go:124
	// _ = "end of CoverTab[30783]"
}

// skipCSSSpace returns a suffix of c, skipping over a single space.
func skipCSSSpace(c []byte) []byte {
//line /usr/local/go/src/html/template/css.go:128
	_go_fuzz_dep_.CoverTab[30792]++
							if len(c) == 0 {
//line /usr/local/go/src/html/template/css.go:129
		_go_fuzz_dep_.CoverTab[30795]++
								return c
//line /usr/local/go/src/html/template/css.go:130
		// _ = "end of CoverTab[30795]"
	} else {
//line /usr/local/go/src/html/template/css.go:131
		_go_fuzz_dep_.CoverTab[30796]++
//line /usr/local/go/src/html/template/css.go:131
		// _ = "end of CoverTab[30796]"
//line /usr/local/go/src/html/template/css.go:131
	}
//line /usr/local/go/src/html/template/css.go:131
	// _ = "end of CoverTab[30792]"
//line /usr/local/go/src/html/template/css.go:131
	_go_fuzz_dep_.CoverTab[30793]++

							switch c[0] {
	case '\t', '\n', '\f', ' ':
//line /usr/local/go/src/html/template/css.go:134
		_go_fuzz_dep_.CoverTab[30797]++
								return c[1:]
//line /usr/local/go/src/html/template/css.go:135
		// _ = "end of CoverTab[30797]"
	case '\r':
//line /usr/local/go/src/html/template/css.go:136
		_go_fuzz_dep_.CoverTab[30798]++

//line /usr/local/go/src/html/template/css.go:140
		if len(c) >= 2 && func() bool {
//line /usr/local/go/src/html/template/css.go:140
			_go_fuzz_dep_.CoverTab[30801]++
//line /usr/local/go/src/html/template/css.go:140
			return c[1] == '\n'
//line /usr/local/go/src/html/template/css.go:140
			// _ = "end of CoverTab[30801]"
//line /usr/local/go/src/html/template/css.go:140
		}() {
//line /usr/local/go/src/html/template/css.go:140
			_go_fuzz_dep_.CoverTab[30802]++
									return c[2:]
//line /usr/local/go/src/html/template/css.go:141
			// _ = "end of CoverTab[30802]"
		} else {
//line /usr/local/go/src/html/template/css.go:142
			_go_fuzz_dep_.CoverTab[30803]++
//line /usr/local/go/src/html/template/css.go:142
			// _ = "end of CoverTab[30803]"
//line /usr/local/go/src/html/template/css.go:142
		}
//line /usr/local/go/src/html/template/css.go:142
		// _ = "end of CoverTab[30798]"
//line /usr/local/go/src/html/template/css.go:142
		_go_fuzz_dep_.CoverTab[30799]++
								return c[1:]
//line /usr/local/go/src/html/template/css.go:143
		// _ = "end of CoverTab[30799]"
//line /usr/local/go/src/html/template/css.go:143
	default:
//line /usr/local/go/src/html/template/css.go:143
		_go_fuzz_dep_.CoverTab[30800]++
//line /usr/local/go/src/html/template/css.go:143
		// _ = "end of CoverTab[30800]"
	}
//line /usr/local/go/src/html/template/css.go:144
	// _ = "end of CoverTab[30793]"
//line /usr/local/go/src/html/template/css.go:144
	_go_fuzz_dep_.CoverTab[30794]++
							return c
//line /usr/local/go/src/html/template/css.go:145
	// _ = "end of CoverTab[30794]"
}

// isCSSSpace reports whether b is a CSS space char as defined in wc.
func isCSSSpace(b byte) bool {
//line /usr/local/go/src/html/template/css.go:149
	_go_fuzz_dep_.CoverTab[30804]++
							switch b {
	case '\t', '\n', '\f', '\r', ' ':
//line /usr/local/go/src/html/template/css.go:151
		_go_fuzz_dep_.CoverTab[30806]++
								return true
//line /usr/local/go/src/html/template/css.go:152
		// _ = "end of CoverTab[30806]"
//line /usr/local/go/src/html/template/css.go:152
	default:
//line /usr/local/go/src/html/template/css.go:152
		_go_fuzz_dep_.CoverTab[30807]++
//line /usr/local/go/src/html/template/css.go:152
		// _ = "end of CoverTab[30807]"
	}
//line /usr/local/go/src/html/template/css.go:153
	// _ = "end of CoverTab[30804]"
//line /usr/local/go/src/html/template/css.go:153
	_go_fuzz_dep_.CoverTab[30805]++
							return false
//line /usr/local/go/src/html/template/css.go:154
	// _ = "end of CoverTab[30805]"
}

// cssEscaper escapes HTML and CSS special characters using \<hex>+ escapes.
func cssEscaper(args ...any) string {
//line /usr/local/go/src/html/template/css.go:158
	_go_fuzz_dep_.CoverTab[30808]++
							s, _ := stringify(args...)
							var b strings.Builder
							r, w, written := rune(0), 0, 0
							for i := 0; i < len(s); i += w {
//line /usr/local/go/src/html/template/css.go:162
		_go_fuzz_dep_.CoverTab[30811]++

								r, w = utf8.DecodeRuneInString(s[i:])
								var repl string
								switch {
		case int(r) < len(cssReplacementTable) && func() bool {
//line /usr/local/go/src/html/template/css.go:167
			_go_fuzz_dep_.CoverTab[30816]++
//line /usr/local/go/src/html/template/css.go:167
			return cssReplacementTable[r] != ""
//line /usr/local/go/src/html/template/css.go:167
			// _ = "end of CoverTab[30816]"
//line /usr/local/go/src/html/template/css.go:167
		}():
//line /usr/local/go/src/html/template/css.go:167
			_go_fuzz_dep_.CoverTab[30814]++
									repl = cssReplacementTable[r]
//line /usr/local/go/src/html/template/css.go:168
			// _ = "end of CoverTab[30814]"
		default:
//line /usr/local/go/src/html/template/css.go:169
			_go_fuzz_dep_.CoverTab[30815]++
									continue
//line /usr/local/go/src/html/template/css.go:170
			// _ = "end of CoverTab[30815]"
		}
//line /usr/local/go/src/html/template/css.go:171
		// _ = "end of CoverTab[30811]"
//line /usr/local/go/src/html/template/css.go:171
		_go_fuzz_dep_.CoverTab[30812]++
								if written == 0 {
//line /usr/local/go/src/html/template/css.go:172
			_go_fuzz_dep_.CoverTab[30817]++
									b.Grow(len(s))
//line /usr/local/go/src/html/template/css.go:173
			// _ = "end of CoverTab[30817]"
		} else {
//line /usr/local/go/src/html/template/css.go:174
			_go_fuzz_dep_.CoverTab[30818]++
//line /usr/local/go/src/html/template/css.go:174
			// _ = "end of CoverTab[30818]"
//line /usr/local/go/src/html/template/css.go:174
		}
//line /usr/local/go/src/html/template/css.go:174
		// _ = "end of CoverTab[30812]"
//line /usr/local/go/src/html/template/css.go:174
		_go_fuzz_dep_.CoverTab[30813]++
								b.WriteString(s[written:i])
								b.WriteString(repl)
								written = i + w
								if repl != `\\` && func() bool {
//line /usr/local/go/src/html/template/css.go:178
			_go_fuzz_dep_.CoverTab[30819]++
//line /usr/local/go/src/html/template/css.go:178
			return (written == len(s) || func() bool {
//line /usr/local/go/src/html/template/css.go:178
				_go_fuzz_dep_.CoverTab[30820]++
//line /usr/local/go/src/html/template/css.go:178
				return isHex(s[written])
//line /usr/local/go/src/html/template/css.go:178
				// _ = "end of CoverTab[30820]"
//line /usr/local/go/src/html/template/css.go:178
			}() || func() bool {
//line /usr/local/go/src/html/template/css.go:178
				_go_fuzz_dep_.CoverTab[30821]++
//line /usr/local/go/src/html/template/css.go:178
				return isCSSSpace(s[written])
//line /usr/local/go/src/html/template/css.go:178
				// _ = "end of CoverTab[30821]"
//line /usr/local/go/src/html/template/css.go:178
			}())
//line /usr/local/go/src/html/template/css.go:178
			// _ = "end of CoverTab[30819]"
//line /usr/local/go/src/html/template/css.go:178
		}() {
//line /usr/local/go/src/html/template/css.go:178
			_go_fuzz_dep_.CoverTab[30822]++
									b.WriteByte(' ')
//line /usr/local/go/src/html/template/css.go:179
			// _ = "end of CoverTab[30822]"
		} else {
//line /usr/local/go/src/html/template/css.go:180
			_go_fuzz_dep_.CoverTab[30823]++
//line /usr/local/go/src/html/template/css.go:180
			// _ = "end of CoverTab[30823]"
//line /usr/local/go/src/html/template/css.go:180
		}
//line /usr/local/go/src/html/template/css.go:180
		// _ = "end of CoverTab[30813]"
	}
//line /usr/local/go/src/html/template/css.go:181
	// _ = "end of CoverTab[30808]"
//line /usr/local/go/src/html/template/css.go:181
	_go_fuzz_dep_.CoverTab[30809]++
							if written == 0 {
//line /usr/local/go/src/html/template/css.go:182
		_go_fuzz_dep_.CoverTab[30824]++
								return s
//line /usr/local/go/src/html/template/css.go:183
		// _ = "end of CoverTab[30824]"
	} else {
//line /usr/local/go/src/html/template/css.go:184
		_go_fuzz_dep_.CoverTab[30825]++
//line /usr/local/go/src/html/template/css.go:184
		// _ = "end of CoverTab[30825]"
//line /usr/local/go/src/html/template/css.go:184
	}
//line /usr/local/go/src/html/template/css.go:184
	// _ = "end of CoverTab[30809]"
//line /usr/local/go/src/html/template/css.go:184
	_go_fuzz_dep_.CoverTab[30810]++
							b.WriteString(s[written:])
							return b.String()
//line /usr/local/go/src/html/template/css.go:186
	// _ = "end of CoverTab[30810]"
}

var cssReplacementTable = []string{
							0:	`\0`,
							'\t':	`\9`,
							'\n':	`\a`,
							'\f':	`\c`,
							'\r':	`\d`,

//line /usr/local/go/src/html/template/css.go:197
	'"':	`\22`,
							'&':	`\26`,
							'\'':	`\27`,
							'(':	`\28`,
							')':	`\29`,
							'+':	`\2b`,
							'/':	`\2f`,
							':':	`\3a`,
							';':	`\3b`,
							'<':	`\3c`,
							'>':	`\3e`,
							'\\':	`\\`,
							'{':	`\7b`,
							'}':	`\7d`,
}

var expressionBytes = []byte("expression")
var mozBindingBytes = []byte("mozbinding")

// cssValueFilter allows innocuous CSS values in the output including CSS
//line /usr/local/go/src/html/template/css.go:216
// quantities (10px or 25%), ID or class literals (#foo, .bar), keyword values
//line /usr/local/go/src/html/template/css.go:216
// (inherit, blue), and colors (#888).
//line /usr/local/go/src/html/template/css.go:216
// It filters out unsafe values, such as those that affect token boundaries,
//line /usr/local/go/src/html/template/css.go:216
// and anything that might execute scripts.
//line /usr/local/go/src/html/template/css.go:221
func cssValueFilter(args ...any) string {
//line /usr/local/go/src/html/template/css.go:221
	_go_fuzz_dep_.CoverTab[30826]++
							s, t := stringify(args...)
							if t == contentTypeCSS {
//line /usr/local/go/src/html/template/css.go:223
		_go_fuzz_dep_.CoverTab[30830]++
								return s
//line /usr/local/go/src/html/template/css.go:224
		// _ = "end of CoverTab[30830]"
	} else {
//line /usr/local/go/src/html/template/css.go:225
		_go_fuzz_dep_.CoverTab[30831]++
//line /usr/local/go/src/html/template/css.go:225
		// _ = "end of CoverTab[30831]"
//line /usr/local/go/src/html/template/css.go:225
	}
//line /usr/local/go/src/html/template/css.go:225
	// _ = "end of CoverTab[30826]"
//line /usr/local/go/src/html/template/css.go:225
	_go_fuzz_dep_.CoverTab[30827]++
							b, id := decodeCSS([]byte(s)), make([]byte, 0, 64)

//line /usr/local/go/src/html/template/css.go:239
	for i, c := range b {
//line /usr/local/go/src/html/template/css.go:239
		_go_fuzz_dep_.CoverTab[30832]++
								switch c {
		case 0, '"', '\'', '(', ')', '/', ';', '@', '[', '\\', ']', '`', '{', '}', '<', '>':
//line /usr/local/go/src/html/template/css.go:241
			_go_fuzz_dep_.CoverTab[30833]++
									return filterFailsafe
//line /usr/local/go/src/html/template/css.go:242
			// _ = "end of CoverTab[30833]"
		case '-':
//line /usr/local/go/src/html/template/css.go:243
			_go_fuzz_dep_.CoverTab[30834]++

//line /usr/local/go/src/html/template/css.go:246
			if i != 0 && func() bool {
//line /usr/local/go/src/html/template/css.go:246
				_go_fuzz_dep_.CoverTab[30836]++
//line /usr/local/go/src/html/template/css.go:246
				return b[i-1] == '-'
//line /usr/local/go/src/html/template/css.go:246
				// _ = "end of CoverTab[30836]"
//line /usr/local/go/src/html/template/css.go:246
			}() {
//line /usr/local/go/src/html/template/css.go:246
				_go_fuzz_dep_.CoverTab[30837]++
										return filterFailsafe
//line /usr/local/go/src/html/template/css.go:247
				// _ = "end of CoverTab[30837]"
			} else {
//line /usr/local/go/src/html/template/css.go:248
				_go_fuzz_dep_.CoverTab[30838]++
//line /usr/local/go/src/html/template/css.go:248
				// _ = "end of CoverTab[30838]"
//line /usr/local/go/src/html/template/css.go:248
			}
//line /usr/local/go/src/html/template/css.go:248
			// _ = "end of CoverTab[30834]"
		default:
//line /usr/local/go/src/html/template/css.go:249
			_go_fuzz_dep_.CoverTab[30835]++
									if c < utf8.RuneSelf && func() bool {
//line /usr/local/go/src/html/template/css.go:250
				_go_fuzz_dep_.CoverTab[30839]++
//line /usr/local/go/src/html/template/css.go:250
				return isCSSNmchar(rune(c))
//line /usr/local/go/src/html/template/css.go:250
				// _ = "end of CoverTab[30839]"
//line /usr/local/go/src/html/template/css.go:250
			}() {
//line /usr/local/go/src/html/template/css.go:250
				_go_fuzz_dep_.CoverTab[30840]++
										id = append(id, c)
//line /usr/local/go/src/html/template/css.go:251
				// _ = "end of CoverTab[30840]"
			} else {
//line /usr/local/go/src/html/template/css.go:252
				_go_fuzz_dep_.CoverTab[30841]++
//line /usr/local/go/src/html/template/css.go:252
				// _ = "end of CoverTab[30841]"
//line /usr/local/go/src/html/template/css.go:252
			}
//line /usr/local/go/src/html/template/css.go:252
			// _ = "end of CoverTab[30835]"
		}
//line /usr/local/go/src/html/template/css.go:253
		// _ = "end of CoverTab[30832]"
	}
//line /usr/local/go/src/html/template/css.go:254
	// _ = "end of CoverTab[30827]"
//line /usr/local/go/src/html/template/css.go:254
	_go_fuzz_dep_.CoverTab[30828]++
							id = bytes.ToLower(id)
							if bytes.Contains(id, expressionBytes) || func() bool {
//line /usr/local/go/src/html/template/css.go:256
		_go_fuzz_dep_.CoverTab[30842]++
//line /usr/local/go/src/html/template/css.go:256
		return bytes.Contains(id, mozBindingBytes)
//line /usr/local/go/src/html/template/css.go:256
		// _ = "end of CoverTab[30842]"
//line /usr/local/go/src/html/template/css.go:256
	}() {
//line /usr/local/go/src/html/template/css.go:256
		_go_fuzz_dep_.CoverTab[30843]++
								return filterFailsafe
//line /usr/local/go/src/html/template/css.go:257
		// _ = "end of CoverTab[30843]"
	} else {
//line /usr/local/go/src/html/template/css.go:258
		_go_fuzz_dep_.CoverTab[30844]++
//line /usr/local/go/src/html/template/css.go:258
		// _ = "end of CoverTab[30844]"
//line /usr/local/go/src/html/template/css.go:258
	}
//line /usr/local/go/src/html/template/css.go:258
	// _ = "end of CoverTab[30828]"
//line /usr/local/go/src/html/template/css.go:258
	_go_fuzz_dep_.CoverTab[30829]++
							return string(b)
//line /usr/local/go/src/html/template/css.go:259
	// _ = "end of CoverTab[30829]"
}

//line /usr/local/go/src/html/template/css.go:260
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/css.go:260
var _ = _go_fuzz_dep_.CoverTab
