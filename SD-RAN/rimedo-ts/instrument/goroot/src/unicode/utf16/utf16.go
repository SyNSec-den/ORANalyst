// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/unicode/utf16/utf16.go:5
// Package utf16 implements encoding and decoding of UTF-16 sequences.
package utf16

//line /usr/local/go/src/unicode/utf16/utf16.go:6
import (
//line /usr/local/go/src/unicode/utf16/utf16.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/unicode/utf16/utf16.go:6
)
//line /usr/local/go/src/unicode/utf16/utf16.go:6
import (
//line /usr/local/go/src/unicode/utf16/utf16.go:6
	_atomic_ "sync/atomic"
//line /usr/local/go/src/unicode/utf16/utf16.go:6
)

//line /usr/local/go/src/unicode/utf16/utf16.go:12
const (
	replacementChar	= '\uFFFD'	// Unicode replacement character
	maxRune		= '\U0010FFFF'	// Maximum valid Unicode code point.
)

const (
	// 0xd800-0xdc00 encodes the high 10 bits of a pair.
	// 0xdc00-0xe000 encodes the low 10 bits of a pair.
	// the value is those 20 bits plus 0x10000.
	surr1	= 0xd800
	surr2	= 0xdc00
	surr3	= 0xe000

	surrSelf	= 0x10000
)

// IsSurrogate reports whether the specified Unicode code point
//line /usr/local/go/src/unicode/utf16/utf16.go:28
// can appear in a surrogate pair.
//line /usr/local/go/src/unicode/utf16/utf16.go:30
func IsSurrogate(r rune) bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:30
	_go_fuzz_dep_.CoverTab[7440]++
							return surr1 <= r && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:31
		_go_fuzz_dep_.CoverTab[7441]++
//line /usr/local/go/src/unicode/utf16/utf16.go:31
		return r < surr3
//line /usr/local/go/src/unicode/utf16/utf16.go:31
		// _ = "end of CoverTab[7441]"
//line /usr/local/go/src/unicode/utf16/utf16.go:31
	}()
//line /usr/local/go/src/unicode/utf16/utf16.go:31
	// _ = "end of CoverTab[7440]"
}

// DecodeRune returns the UTF-16 decoding of a surrogate pair.
//line /usr/local/go/src/unicode/utf16/utf16.go:34
// If the pair is not a valid UTF-16 surrogate pair, DecodeRune returns
//line /usr/local/go/src/unicode/utf16/utf16.go:34
// the Unicode replacement code point U+FFFD.
//line /usr/local/go/src/unicode/utf16/utf16.go:37
func DecodeRune(r1, r2 rune) rune {
//line /usr/local/go/src/unicode/utf16/utf16.go:37
	_go_fuzz_dep_.CoverTab[7442]++
							if surr1 <= r1 && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		_go_fuzz_dep_.CoverTab[7444]++
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		return r1 < surr2
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		// _ = "end of CoverTab[7444]"
//line /usr/local/go/src/unicode/utf16/utf16.go:38
	}() && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		_go_fuzz_dep_.CoverTab[7445]++
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		return surr2 <= r2
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		// _ = "end of CoverTab[7445]"
//line /usr/local/go/src/unicode/utf16/utf16.go:38
	}() && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		_go_fuzz_dep_.CoverTab[7446]++
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		return r2 < surr3
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		// _ = "end of CoverTab[7446]"
//line /usr/local/go/src/unicode/utf16/utf16.go:38
	}() {
//line /usr/local/go/src/unicode/utf16/utf16.go:38
		_go_fuzz_dep_.CoverTab[7447]++
								return (r1-surr1)<<10 | (r2 - surr2) + surrSelf
//line /usr/local/go/src/unicode/utf16/utf16.go:39
		// _ = "end of CoverTab[7447]"
	} else {
//line /usr/local/go/src/unicode/utf16/utf16.go:40
		_go_fuzz_dep_.CoverTab[7448]++
//line /usr/local/go/src/unicode/utf16/utf16.go:40
		// _ = "end of CoverTab[7448]"
//line /usr/local/go/src/unicode/utf16/utf16.go:40
	}
//line /usr/local/go/src/unicode/utf16/utf16.go:40
	// _ = "end of CoverTab[7442]"
//line /usr/local/go/src/unicode/utf16/utf16.go:40
	_go_fuzz_dep_.CoverTab[7443]++
							return replacementChar
//line /usr/local/go/src/unicode/utf16/utf16.go:41
	// _ = "end of CoverTab[7443]"
}

// EncodeRune returns the UTF-16 surrogate pair r1, r2 for the given rune.
//line /usr/local/go/src/unicode/utf16/utf16.go:44
// If the rune is not a valid Unicode code point or does not need encoding,
//line /usr/local/go/src/unicode/utf16/utf16.go:44
// EncodeRune returns U+FFFD, U+FFFD.
//line /usr/local/go/src/unicode/utf16/utf16.go:47
func EncodeRune(r rune) (r1, r2 rune) {
//line /usr/local/go/src/unicode/utf16/utf16.go:47
	_go_fuzz_dep_.CoverTab[7449]++
							if r < surrSelf || func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:48
		_go_fuzz_dep_.CoverTab[7451]++
//line /usr/local/go/src/unicode/utf16/utf16.go:48
		return r > maxRune
//line /usr/local/go/src/unicode/utf16/utf16.go:48
		// _ = "end of CoverTab[7451]"
//line /usr/local/go/src/unicode/utf16/utf16.go:48
	}() {
//line /usr/local/go/src/unicode/utf16/utf16.go:48
		_go_fuzz_dep_.CoverTab[7452]++
								return replacementChar, replacementChar
//line /usr/local/go/src/unicode/utf16/utf16.go:49
		// _ = "end of CoverTab[7452]"
	} else {
//line /usr/local/go/src/unicode/utf16/utf16.go:50
		_go_fuzz_dep_.CoverTab[7453]++
//line /usr/local/go/src/unicode/utf16/utf16.go:50
		// _ = "end of CoverTab[7453]"
//line /usr/local/go/src/unicode/utf16/utf16.go:50
	}
//line /usr/local/go/src/unicode/utf16/utf16.go:50
	// _ = "end of CoverTab[7449]"
//line /usr/local/go/src/unicode/utf16/utf16.go:50
	_go_fuzz_dep_.CoverTab[7450]++
							r -= surrSelf
							return surr1 + (r>>10)&0x3ff, surr2 + r&0x3ff
//line /usr/local/go/src/unicode/utf16/utf16.go:52
	// _ = "end of CoverTab[7450]"
}

// Encode returns the UTF-16 encoding of the Unicode code point sequence s.
func Encode(s []rune) []uint16 {
//line /usr/local/go/src/unicode/utf16/utf16.go:56
	_go_fuzz_dep_.CoverTab[7454]++
							n := len(s)
							for _, v := range s {
//line /usr/local/go/src/unicode/utf16/utf16.go:58
		_go_fuzz_dep_.CoverTab[7457]++
								if v >= surrSelf {
//line /usr/local/go/src/unicode/utf16/utf16.go:59
			_go_fuzz_dep_.CoverTab[7458]++
									n++
//line /usr/local/go/src/unicode/utf16/utf16.go:60
			// _ = "end of CoverTab[7458]"
		} else {
//line /usr/local/go/src/unicode/utf16/utf16.go:61
			_go_fuzz_dep_.CoverTab[7459]++
//line /usr/local/go/src/unicode/utf16/utf16.go:61
			// _ = "end of CoverTab[7459]"
//line /usr/local/go/src/unicode/utf16/utf16.go:61
		}
//line /usr/local/go/src/unicode/utf16/utf16.go:61
		// _ = "end of CoverTab[7457]"
	}
//line /usr/local/go/src/unicode/utf16/utf16.go:62
	// _ = "end of CoverTab[7454]"
//line /usr/local/go/src/unicode/utf16/utf16.go:62
	_go_fuzz_dep_.CoverTab[7455]++

							a := make([]uint16, n)
							n = 0
							for _, v := range s {
//line /usr/local/go/src/unicode/utf16/utf16.go:66
		_go_fuzz_dep_.CoverTab[7460]++
								switch {
		case 0 <= v && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:68
			_go_fuzz_dep_.CoverTab[7464]++
//line /usr/local/go/src/unicode/utf16/utf16.go:68
			return v < surr1
//line /usr/local/go/src/unicode/utf16/utf16.go:68
			// _ = "end of CoverTab[7464]"
//line /usr/local/go/src/unicode/utf16/utf16.go:68
		}(), surr3 <= v && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:68
			_go_fuzz_dep_.CoverTab[7465]++
//line /usr/local/go/src/unicode/utf16/utf16.go:68
			return v < surrSelf
//line /usr/local/go/src/unicode/utf16/utf16.go:68
			// _ = "end of CoverTab[7465]"
//line /usr/local/go/src/unicode/utf16/utf16.go:68
		}():
//line /usr/local/go/src/unicode/utf16/utf16.go:68
			_go_fuzz_dep_.CoverTab[7461]++

									a[n] = uint16(v)
									n++
//line /usr/local/go/src/unicode/utf16/utf16.go:71
			// _ = "end of CoverTab[7461]"
		case surrSelf <= v && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:72
			_go_fuzz_dep_.CoverTab[7466]++
//line /usr/local/go/src/unicode/utf16/utf16.go:72
			return v <= maxRune
//line /usr/local/go/src/unicode/utf16/utf16.go:72
			// _ = "end of CoverTab[7466]"
//line /usr/local/go/src/unicode/utf16/utf16.go:72
		}():
//line /usr/local/go/src/unicode/utf16/utf16.go:72
			_go_fuzz_dep_.CoverTab[7462]++

									r1, r2 := EncodeRune(v)
									a[n] = uint16(r1)
									a[n+1] = uint16(r2)
									n += 2
//line /usr/local/go/src/unicode/utf16/utf16.go:77
			// _ = "end of CoverTab[7462]"
		default:
//line /usr/local/go/src/unicode/utf16/utf16.go:78
			_go_fuzz_dep_.CoverTab[7463]++
									a[n] = uint16(replacementChar)
									n++
//line /usr/local/go/src/unicode/utf16/utf16.go:80
			// _ = "end of CoverTab[7463]"
		}
//line /usr/local/go/src/unicode/utf16/utf16.go:81
		// _ = "end of CoverTab[7460]"
	}
//line /usr/local/go/src/unicode/utf16/utf16.go:82
	// _ = "end of CoverTab[7455]"
//line /usr/local/go/src/unicode/utf16/utf16.go:82
	_go_fuzz_dep_.CoverTab[7456]++
							return a[:n]
//line /usr/local/go/src/unicode/utf16/utf16.go:83
	// _ = "end of CoverTab[7456]"
}

// AppendRune appends the UTF-16 encoding of the Unicode code point r
//line /usr/local/go/src/unicode/utf16/utf16.go:86
// to the end of p and returns the extended buffer. If the rune is not
//line /usr/local/go/src/unicode/utf16/utf16.go:86
// a valid Unicode code point, it appends the encoding of U+FFFD.
//line /usr/local/go/src/unicode/utf16/utf16.go:89
func AppendRune(a []uint16, r rune) []uint16 {
//line /usr/local/go/src/unicode/utf16/utf16.go:89
	_go_fuzz_dep_.CoverTab[7467]++

							switch {
	case 0 <= r && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:92
		_go_fuzz_dep_.CoverTab[7472]++
//line /usr/local/go/src/unicode/utf16/utf16.go:92
		return r < surr1
//line /usr/local/go/src/unicode/utf16/utf16.go:92
		// _ = "end of CoverTab[7472]"
//line /usr/local/go/src/unicode/utf16/utf16.go:92
	}(), surr3 <= r && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:92
		_go_fuzz_dep_.CoverTab[7473]++
//line /usr/local/go/src/unicode/utf16/utf16.go:92
		return r < surrSelf
//line /usr/local/go/src/unicode/utf16/utf16.go:92
		// _ = "end of CoverTab[7473]"
//line /usr/local/go/src/unicode/utf16/utf16.go:92
	}():
//line /usr/local/go/src/unicode/utf16/utf16.go:92
		_go_fuzz_dep_.CoverTab[7469]++

								return append(a, uint16(r))
//line /usr/local/go/src/unicode/utf16/utf16.go:94
		// _ = "end of CoverTab[7469]"
	case surrSelf <= r && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:95
		_go_fuzz_dep_.CoverTab[7474]++
//line /usr/local/go/src/unicode/utf16/utf16.go:95
		return r <= maxRune
//line /usr/local/go/src/unicode/utf16/utf16.go:95
		// _ = "end of CoverTab[7474]"
//line /usr/local/go/src/unicode/utf16/utf16.go:95
	}():
//line /usr/local/go/src/unicode/utf16/utf16.go:95
		_go_fuzz_dep_.CoverTab[7470]++

								r1, r2 := EncodeRune(r)
								return append(a, uint16(r1), uint16(r2))
//line /usr/local/go/src/unicode/utf16/utf16.go:98
		// _ = "end of CoverTab[7470]"
//line /usr/local/go/src/unicode/utf16/utf16.go:98
	default:
//line /usr/local/go/src/unicode/utf16/utf16.go:98
		_go_fuzz_dep_.CoverTab[7471]++
//line /usr/local/go/src/unicode/utf16/utf16.go:98
		// _ = "end of CoverTab[7471]"
	}
//line /usr/local/go/src/unicode/utf16/utf16.go:99
	// _ = "end of CoverTab[7467]"
//line /usr/local/go/src/unicode/utf16/utf16.go:99
	_go_fuzz_dep_.CoverTab[7468]++
							return append(a, replacementChar)
//line /usr/local/go/src/unicode/utf16/utf16.go:100
	// _ = "end of CoverTab[7468]"
}

// Decode returns the Unicode code point sequence represented
//line /usr/local/go/src/unicode/utf16/utf16.go:103
// by the UTF-16 encoding s.
//line /usr/local/go/src/unicode/utf16/utf16.go:105
func Decode(s []uint16) []rune {
//line /usr/local/go/src/unicode/utf16/utf16.go:105
	_go_fuzz_dep_.CoverTab[7475]++
							a := make([]rune, len(s))
							n := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/unicode/utf16/utf16.go:108
		_go_fuzz_dep_.CoverTab[7477]++
								switch r := s[i]; {
		case r < surr1, surr3 <= r:
//line /usr/local/go/src/unicode/utf16/utf16.go:110
			_go_fuzz_dep_.CoverTab[7479]++

									a[n] = rune(r)
//line /usr/local/go/src/unicode/utf16/utf16.go:112
			// _ = "end of CoverTab[7479]"
		case surr1 <= r && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:113
			_go_fuzz_dep_.CoverTab[7482]++
//line /usr/local/go/src/unicode/utf16/utf16.go:113
			return r < surr2
//line /usr/local/go/src/unicode/utf16/utf16.go:113
			// _ = "end of CoverTab[7482]"
//line /usr/local/go/src/unicode/utf16/utf16.go:113
		}() && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:113
			_go_fuzz_dep_.CoverTab[7483]++
//line /usr/local/go/src/unicode/utf16/utf16.go:113
			return i+1 < len(s)
//line /usr/local/go/src/unicode/utf16/utf16.go:113
			// _ = "end of CoverTab[7483]"
//line /usr/local/go/src/unicode/utf16/utf16.go:113
		}() && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:113
			_go_fuzz_dep_.CoverTab[7484]++
//line /usr/local/go/src/unicode/utf16/utf16.go:113
			return surr2 <= s[i+1]
									// _ = "end of CoverTab[7484]"
//line /usr/local/go/src/unicode/utf16/utf16.go:114
		}() && func() bool {
//line /usr/local/go/src/unicode/utf16/utf16.go:114
			_go_fuzz_dep_.CoverTab[7485]++
//line /usr/local/go/src/unicode/utf16/utf16.go:114
			return s[i+1] < surr3
//line /usr/local/go/src/unicode/utf16/utf16.go:114
			// _ = "end of CoverTab[7485]"
//line /usr/local/go/src/unicode/utf16/utf16.go:114
		}():
//line /usr/local/go/src/unicode/utf16/utf16.go:114
			_go_fuzz_dep_.CoverTab[7480]++

									a[n] = DecodeRune(rune(r), rune(s[i+1]))
									i++
//line /usr/local/go/src/unicode/utf16/utf16.go:117
			// _ = "end of CoverTab[7480]"
		default:
//line /usr/local/go/src/unicode/utf16/utf16.go:118
			_go_fuzz_dep_.CoverTab[7481]++

									a[n] = replacementChar
//line /usr/local/go/src/unicode/utf16/utf16.go:120
			// _ = "end of CoverTab[7481]"
		}
//line /usr/local/go/src/unicode/utf16/utf16.go:121
		// _ = "end of CoverTab[7477]"
//line /usr/local/go/src/unicode/utf16/utf16.go:121
		_go_fuzz_dep_.CoverTab[7478]++
								n++
//line /usr/local/go/src/unicode/utf16/utf16.go:122
		// _ = "end of CoverTab[7478]"
	}
//line /usr/local/go/src/unicode/utf16/utf16.go:123
	// _ = "end of CoverTab[7475]"
//line /usr/local/go/src/unicode/utf16/utf16.go:123
	_go_fuzz_dep_.CoverTab[7476]++
							return a[:n]
//line /usr/local/go/src/unicode/utf16/utf16.go:124
	// _ = "end of CoverTab[7476]"
}

//line /usr/local/go/src/unicode/utf16/utf16.go:125
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/unicode/utf16/utf16.go:125
var _ = _go_fuzz_dep_.CoverTab
