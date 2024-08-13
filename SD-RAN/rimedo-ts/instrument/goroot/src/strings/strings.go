// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/strings/strings.go:5
// Package strings implements simple functions to manipulate UTF-8 encoded strings.
//line /usr/local/go/src/strings/strings.go:5
//
//line /usr/local/go/src/strings/strings.go:5
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
//line /usr/local/go/src/strings/strings.go:8
package strings

//line /usr/local/go/src/strings/strings.go:8
import (
//line /usr/local/go/src/strings/strings.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/strings/strings.go:8
)
//line /usr/local/go/src/strings/strings.go:8
import (
//line /usr/local/go/src/strings/strings.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/strings/strings.go:8
)

import (
	"internal/bytealg"
	"unicode"
	"unicode/utf8"
)

// explode splits s into a slice of UTF-8 strings,
//line /usr/local/go/src/strings/strings.go:16
// one string per Unicode character up to a maximum of n (n < 0 means no limit).
//line /usr/local/go/src/strings/strings.go:16
// Invalid UTF-8 bytes are sliced individually.
//line /usr/local/go/src/strings/strings.go:19
func explode(s string, n int) []string {
//line /usr/local/go/src/strings/strings.go:19
	_go_fuzz_dep_.CoverTab[3450]++
						l := utf8.RuneCountInString(s)
						if n < 0 || func() bool {
//line /usr/local/go/src/strings/strings.go:21
		_go_fuzz_dep_.CoverTab[3454]++
//line /usr/local/go/src/strings/strings.go:21
		return n > l
//line /usr/local/go/src/strings/strings.go:21
		// _ = "end of CoverTab[3454]"
//line /usr/local/go/src/strings/strings.go:21
	}() {
//line /usr/local/go/src/strings/strings.go:21
		_go_fuzz_dep_.CoverTab[3455]++
							n = l
//line /usr/local/go/src/strings/strings.go:22
		// _ = "end of CoverTab[3455]"
	} else {
//line /usr/local/go/src/strings/strings.go:23
		_go_fuzz_dep_.CoverTab[3456]++
//line /usr/local/go/src/strings/strings.go:23
		// _ = "end of CoverTab[3456]"
//line /usr/local/go/src/strings/strings.go:23
	}
//line /usr/local/go/src/strings/strings.go:23
	// _ = "end of CoverTab[3450]"
//line /usr/local/go/src/strings/strings.go:23
	_go_fuzz_dep_.CoverTab[3451]++
						a := make([]string, n)
						for i := 0; i < n-1; i++ {
//line /usr/local/go/src/strings/strings.go:25
		_go_fuzz_dep_.CoverTab[3457]++
							_, size := utf8.DecodeRuneInString(s)
							a[i] = s[:size]
							s = s[size:]
//line /usr/local/go/src/strings/strings.go:28
		// _ = "end of CoverTab[3457]"
	}
//line /usr/local/go/src/strings/strings.go:29
	// _ = "end of CoverTab[3451]"
//line /usr/local/go/src/strings/strings.go:29
	_go_fuzz_dep_.CoverTab[3452]++
						if n > 0 {
//line /usr/local/go/src/strings/strings.go:30
		_go_fuzz_dep_.CoverTab[3458]++
							a[n-1] = s
//line /usr/local/go/src/strings/strings.go:31
		// _ = "end of CoverTab[3458]"
	} else {
//line /usr/local/go/src/strings/strings.go:32
		_go_fuzz_dep_.CoverTab[3459]++
//line /usr/local/go/src/strings/strings.go:32
		// _ = "end of CoverTab[3459]"
//line /usr/local/go/src/strings/strings.go:32
	}
//line /usr/local/go/src/strings/strings.go:32
	// _ = "end of CoverTab[3452]"
//line /usr/local/go/src/strings/strings.go:32
	_go_fuzz_dep_.CoverTab[3453]++
						return a
//line /usr/local/go/src/strings/strings.go:33
	// _ = "end of CoverTab[3453]"
}

// Count counts the number of non-overlapping instances of substr in s.
//line /usr/local/go/src/strings/strings.go:36
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
//line /usr/local/go/src/strings/strings.go:38
func Count(s, substr string) int {
//line /usr/local/go/src/strings/strings.go:38
	_go_fuzz_dep_.CoverTab[3460]++

						if len(substr) == 0 {
//line /usr/local/go/src/strings/strings.go:40
		_go_fuzz_dep_.CoverTab[3463]++
							return utf8.RuneCountInString(s) + 1
//line /usr/local/go/src/strings/strings.go:41
		// _ = "end of CoverTab[3463]"
	} else {
//line /usr/local/go/src/strings/strings.go:42
		_go_fuzz_dep_.CoverTab[3464]++
//line /usr/local/go/src/strings/strings.go:42
		// _ = "end of CoverTab[3464]"
//line /usr/local/go/src/strings/strings.go:42
	}
//line /usr/local/go/src/strings/strings.go:42
	// _ = "end of CoverTab[3460]"
//line /usr/local/go/src/strings/strings.go:42
	_go_fuzz_dep_.CoverTab[3461]++
						if len(substr) == 1 {
//line /usr/local/go/src/strings/strings.go:43
		_go_fuzz_dep_.CoverTab[3465]++
							return bytealg.CountString(s, substr[0])
//line /usr/local/go/src/strings/strings.go:44
		// _ = "end of CoverTab[3465]"
	} else {
//line /usr/local/go/src/strings/strings.go:45
		_go_fuzz_dep_.CoverTab[3466]++
//line /usr/local/go/src/strings/strings.go:45
		// _ = "end of CoverTab[3466]"
//line /usr/local/go/src/strings/strings.go:45
	}
//line /usr/local/go/src/strings/strings.go:45
	// _ = "end of CoverTab[3461]"
//line /usr/local/go/src/strings/strings.go:45
	_go_fuzz_dep_.CoverTab[3462]++
						n := 0
						for {
//line /usr/local/go/src/strings/strings.go:47
		_go_fuzz_dep_.CoverTab[3467]++
							i := Index(s, substr)
							if i == -1 {
//line /usr/local/go/src/strings/strings.go:49
			_go_fuzz_dep_.CoverTab[3469]++
								return n
//line /usr/local/go/src/strings/strings.go:50
			// _ = "end of CoverTab[3469]"
		} else {
//line /usr/local/go/src/strings/strings.go:51
			_go_fuzz_dep_.CoverTab[3470]++
//line /usr/local/go/src/strings/strings.go:51
			// _ = "end of CoverTab[3470]"
//line /usr/local/go/src/strings/strings.go:51
		}
//line /usr/local/go/src/strings/strings.go:51
		// _ = "end of CoverTab[3467]"
//line /usr/local/go/src/strings/strings.go:51
		_go_fuzz_dep_.CoverTab[3468]++
							n++
							s = s[i+len(substr):]
//line /usr/local/go/src/strings/strings.go:53
		// _ = "end of CoverTab[3468]"
	}
//line /usr/local/go/src/strings/strings.go:54
	// _ = "end of CoverTab[3462]"
}

// Contains reports whether substr is within s.
func Contains(s, substr string) bool {
//line /usr/local/go/src/strings/strings.go:58
	_go_fuzz_dep_.CoverTab[3471]++
						return Index(s, substr) >= 0
//line /usr/local/go/src/strings/strings.go:59
	// _ = "end of CoverTab[3471]"
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool {
//line /usr/local/go/src/strings/strings.go:63
	_go_fuzz_dep_.CoverTab[3472]++
						return IndexAny(s, chars) >= 0
//line /usr/local/go/src/strings/strings.go:64
	// _ = "end of CoverTab[3472]"
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool {
//line /usr/local/go/src/strings/strings.go:68
	_go_fuzz_dep_.CoverTab[3473]++
						return IndexRune(s, r) >= 0
//line /usr/local/go/src/strings/strings.go:69
	// _ = "end of CoverTab[3473]"
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int {
//line /usr/local/go/src/strings/strings.go:73
	_go_fuzz_dep_.CoverTab[3474]++
						n := len(substr)
						switch {
	case n == 0:
//line /usr/local/go/src/strings/strings.go:76
		_go_fuzz_dep_.CoverTab[3479]++
							return len(s)
//line /usr/local/go/src/strings/strings.go:77
		// _ = "end of CoverTab[3479]"
	case n == 1:
//line /usr/local/go/src/strings/strings.go:78
		_go_fuzz_dep_.CoverTab[3480]++
							return LastIndexByte(s, substr[0])
//line /usr/local/go/src/strings/strings.go:79
		// _ = "end of CoverTab[3480]"
	case n == len(s):
//line /usr/local/go/src/strings/strings.go:80
		_go_fuzz_dep_.CoverTab[3481]++
							if substr == s {
//line /usr/local/go/src/strings/strings.go:81
			_go_fuzz_dep_.CoverTab[3485]++
								return 0
//line /usr/local/go/src/strings/strings.go:82
			// _ = "end of CoverTab[3485]"
		} else {
//line /usr/local/go/src/strings/strings.go:83
			_go_fuzz_dep_.CoverTab[3486]++
//line /usr/local/go/src/strings/strings.go:83
			// _ = "end of CoverTab[3486]"
//line /usr/local/go/src/strings/strings.go:83
		}
//line /usr/local/go/src/strings/strings.go:83
		// _ = "end of CoverTab[3481]"
//line /usr/local/go/src/strings/strings.go:83
		_go_fuzz_dep_.CoverTab[3482]++
							return -1
//line /usr/local/go/src/strings/strings.go:84
		// _ = "end of CoverTab[3482]"
	case n > len(s):
//line /usr/local/go/src/strings/strings.go:85
		_go_fuzz_dep_.CoverTab[3483]++
							return -1
//line /usr/local/go/src/strings/strings.go:86
		// _ = "end of CoverTab[3483]"
//line /usr/local/go/src/strings/strings.go:86
	default:
//line /usr/local/go/src/strings/strings.go:86
		_go_fuzz_dep_.CoverTab[3484]++
//line /usr/local/go/src/strings/strings.go:86
		// _ = "end of CoverTab[3484]"
	}
//line /usr/local/go/src/strings/strings.go:87
	// _ = "end of CoverTab[3474]"
//line /usr/local/go/src/strings/strings.go:87
	_go_fuzz_dep_.CoverTab[3475]++

						hashss, pow := bytealg.HashStrRev(substr)
						last := len(s) - n
						var h uint32
						for i := len(s) - 1; i >= last; i-- {
//line /usr/local/go/src/strings/strings.go:92
		_go_fuzz_dep_.CoverTab[3487]++
							h = h*bytealg.PrimeRK + uint32(s[i])
//line /usr/local/go/src/strings/strings.go:93
		// _ = "end of CoverTab[3487]"
	}
//line /usr/local/go/src/strings/strings.go:94
	// _ = "end of CoverTab[3475]"
//line /usr/local/go/src/strings/strings.go:94
	_go_fuzz_dep_.CoverTab[3476]++
						if h == hashss && func() bool {
//line /usr/local/go/src/strings/strings.go:95
		_go_fuzz_dep_.CoverTab[3488]++
//line /usr/local/go/src/strings/strings.go:95
		return s[last:] == substr
//line /usr/local/go/src/strings/strings.go:95
		// _ = "end of CoverTab[3488]"
//line /usr/local/go/src/strings/strings.go:95
	}() {
//line /usr/local/go/src/strings/strings.go:95
		_go_fuzz_dep_.CoverTab[3489]++
							return last
//line /usr/local/go/src/strings/strings.go:96
		// _ = "end of CoverTab[3489]"
	} else {
//line /usr/local/go/src/strings/strings.go:97
		_go_fuzz_dep_.CoverTab[3490]++
//line /usr/local/go/src/strings/strings.go:97
		// _ = "end of CoverTab[3490]"
//line /usr/local/go/src/strings/strings.go:97
	}
//line /usr/local/go/src/strings/strings.go:97
	// _ = "end of CoverTab[3476]"
//line /usr/local/go/src/strings/strings.go:97
	_go_fuzz_dep_.CoverTab[3477]++
						for i := last - 1; i >= 0; i-- {
//line /usr/local/go/src/strings/strings.go:98
		_go_fuzz_dep_.CoverTab[3491]++
							h *= bytealg.PrimeRK
							h += uint32(s[i])
							h -= pow * uint32(s[i+n])
							if h == hashss && func() bool {
//line /usr/local/go/src/strings/strings.go:102
			_go_fuzz_dep_.CoverTab[3492]++
//line /usr/local/go/src/strings/strings.go:102
			return s[i:i+n] == substr
//line /usr/local/go/src/strings/strings.go:102
			// _ = "end of CoverTab[3492]"
//line /usr/local/go/src/strings/strings.go:102
		}() {
//line /usr/local/go/src/strings/strings.go:102
			_go_fuzz_dep_.CoverTab[3493]++
									return i
//line /usr/local/go/src/strings/strings.go:103
			// _ = "end of CoverTab[3493]"
		} else {
//line /usr/local/go/src/strings/strings.go:104
			_go_fuzz_dep_.CoverTab[3494]++
//line /usr/local/go/src/strings/strings.go:104
			// _ = "end of CoverTab[3494]"
//line /usr/local/go/src/strings/strings.go:104
		}
//line /usr/local/go/src/strings/strings.go:104
		// _ = "end of CoverTab[3491]"
	}
//line /usr/local/go/src/strings/strings.go:105
	// _ = "end of CoverTab[3477]"
//line /usr/local/go/src/strings/strings.go:105
	_go_fuzz_dep_.CoverTab[3478]++
							return -1
//line /usr/local/go/src/strings/strings.go:106
	// _ = "end of CoverTab[3478]"
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s string, c byte) int {
//line /usr/local/go/src/strings/strings.go:110
	_go_fuzz_dep_.CoverTab[3495]++
							return bytealg.IndexByteString(s, c)
//line /usr/local/go/src/strings/strings.go:111
	// _ = "end of CoverTab[3495]"
}

// IndexRune returns the index of the first instance of the Unicode code point
//line /usr/local/go/src/strings/strings.go:114
// r, or -1 if rune is not present in s.
//line /usr/local/go/src/strings/strings.go:114
// If r is utf8.RuneError, it returns the first instance of any
//line /usr/local/go/src/strings/strings.go:114
// invalid UTF-8 byte sequence.
//line /usr/local/go/src/strings/strings.go:118
func IndexRune(s string, r rune) int {
//line /usr/local/go/src/strings/strings.go:118
	_go_fuzz_dep_.CoverTab[3496]++
							switch {
	case 0 <= r && func() bool {
//line /usr/local/go/src/strings/strings.go:120
		_go_fuzz_dep_.CoverTab[3502]++
//line /usr/local/go/src/strings/strings.go:120
		return r < utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:120
		// _ = "end of CoverTab[3502]"
//line /usr/local/go/src/strings/strings.go:120
	}():
//line /usr/local/go/src/strings/strings.go:120
		_go_fuzz_dep_.CoverTab[3497]++
								return IndexByte(s, byte(r))
//line /usr/local/go/src/strings/strings.go:121
		// _ = "end of CoverTab[3497]"
	case r == utf8.RuneError:
//line /usr/local/go/src/strings/strings.go:122
		_go_fuzz_dep_.CoverTab[3498]++
								for i, r := range s {
//line /usr/local/go/src/strings/strings.go:123
			_go_fuzz_dep_.CoverTab[3503]++
									if r == utf8.RuneError {
//line /usr/local/go/src/strings/strings.go:124
				_go_fuzz_dep_.CoverTab[3504]++
										return i
//line /usr/local/go/src/strings/strings.go:125
				// _ = "end of CoverTab[3504]"
			} else {
//line /usr/local/go/src/strings/strings.go:126
				_go_fuzz_dep_.CoverTab[3505]++
//line /usr/local/go/src/strings/strings.go:126
				// _ = "end of CoverTab[3505]"
//line /usr/local/go/src/strings/strings.go:126
			}
//line /usr/local/go/src/strings/strings.go:126
			// _ = "end of CoverTab[3503]"
		}
//line /usr/local/go/src/strings/strings.go:127
		// _ = "end of CoverTab[3498]"
//line /usr/local/go/src/strings/strings.go:127
		_go_fuzz_dep_.CoverTab[3499]++
								return -1
//line /usr/local/go/src/strings/strings.go:128
		// _ = "end of CoverTab[3499]"
	case !utf8.ValidRune(r):
//line /usr/local/go/src/strings/strings.go:129
		_go_fuzz_dep_.CoverTab[3500]++
								return -1
//line /usr/local/go/src/strings/strings.go:130
		// _ = "end of CoverTab[3500]"
	default:
//line /usr/local/go/src/strings/strings.go:131
		_go_fuzz_dep_.CoverTab[3501]++
								return Index(s, string(r))
//line /usr/local/go/src/strings/strings.go:132
		// _ = "end of CoverTab[3501]"
	}
//line /usr/local/go/src/strings/strings.go:133
	// _ = "end of CoverTab[3496]"
}

// IndexAny returns the index of the first instance of any Unicode code point
//line /usr/local/go/src/strings/strings.go:136
// from chars in s, or -1 if no Unicode code point from chars is present in s.
//line /usr/local/go/src/strings/strings.go:138
func IndexAny(s, chars string) int {
//line /usr/local/go/src/strings/strings.go:138
	_go_fuzz_dep_.CoverTab[3506]++
							if chars == "" {
//line /usr/local/go/src/strings/strings.go:139
		_go_fuzz_dep_.CoverTab[3511]++

								return -1
//line /usr/local/go/src/strings/strings.go:141
		// _ = "end of CoverTab[3511]"
	} else {
//line /usr/local/go/src/strings/strings.go:142
		_go_fuzz_dep_.CoverTab[3512]++
//line /usr/local/go/src/strings/strings.go:142
		// _ = "end of CoverTab[3512]"
//line /usr/local/go/src/strings/strings.go:142
	}
//line /usr/local/go/src/strings/strings.go:142
	// _ = "end of CoverTab[3506]"
//line /usr/local/go/src/strings/strings.go:142
	_go_fuzz_dep_.CoverTab[3507]++
							if len(chars) == 1 {
//line /usr/local/go/src/strings/strings.go:143
		_go_fuzz_dep_.CoverTab[3513]++

								r := rune(chars[0])
								if r >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:146
			_go_fuzz_dep_.CoverTab[3515]++
									r = utf8.RuneError
//line /usr/local/go/src/strings/strings.go:147
			// _ = "end of CoverTab[3515]"
		} else {
//line /usr/local/go/src/strings/strings.go:148
			_go_fuzz_dep_.CoverTab[3516]++
//line /usr/local/go/src/strings/strings.go:148
			// _ = "end of CoverTab[3516]"
//line /usr/local/go/src/strings/strings.go:148
		}
//line /usr/local/go/src/strings/strings.go:148
		// _ = "end of CoverTab[3513]"
//line /usr/local/go/src/strings/strings.go:148
		_go_fuzz_dep_.CoverTab[3514]++
								return IndexRune(s, r)
//line /usr/local/go/src/strings/strings.go:149
		// _ = "end of CoverTab[3514]"
	} else {
//line /usr/local/go/src/strings/strings.go:150
		_go_fuzz_dep_.CoverTab[3517]++
//line /usr/local/go/src/strings/strings.go:150
		// _ = "end of CoverTab[3517]"
//line /usr/local/go/src/strings/strings.go:150
	}
//line /usr/local/go/src/strings/strings.go:150
	// _ = "end of CoverTab[3507]"
//line /usr/local/go/src/strings/strings.go:150
	_go_fuzz_dep_.CoverTab[3508]++
							if len(s) > 8 {
//line /usr/local/go/src/strings/strings.go:151
		_go_fuzz_dep_.CoverTab[3518]++
								if as, isASCII := makeASCIISet(chars); isASCII {
//line /usr/local/go/src/strings/strings.go:152
			_go_fuzz_dep_.CoverTab[3519]++
									for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:153
				_go_fuzz_dep_.CoverTab[3521]++
										if as.contains(s[i]) {
//line /usr/local/go/src/strings/strings.go:154
					_go_fuzz_dep_.CoverTab[3522]++
											return i
//line /usr/local/go/src/strings/strings.go:155
					// _ = "end of CoverTab[3522]"
				} else {
//line /usr/local/go/src/strings/strings.go:156
					_go_fuzz_dep_.CoverTab[3523]++
//line /usr/local/go/src/strings/strings.go:156
					// _ = "end of CoverTab[3523]"
//line /usr/local/go/src/strings/strings.go:156
				}
//line /usr/local/go/src/strings/strings.go:156
				// _ = "end of CoverTab[3521]"
			}
//line /usr/local/go/src/strings/strings.go:157
			// _ = "end of CoverTab[3519]"
//line /usr/local/go/src/strings/strings.go:157
			_go_fuzz_dep_.CoverTab[3520]++
									return -1
//line /usr/local/go/src/strings/strings.go:158
			// _ = "end of CoverTab[3520]"
		} else {
//line /usr/local/go/src/strings/strings.go:159
			_go_fuzz_dep_.CoverTab[3524]++
//line /usr/local/go/src/strings/strings.go:159
			// _ = "end of CoverTab[3524]"
//line /usr/local/go/src/strings/strings.go:159
		}
//line /usr/local/go/src/strings/strings.go:159
		// _ = "end of CoverTab[3518]"
	} else {
//line /usr/local/go/src/strings/strings.go:160
		_go_fuzz_dep_.CoverTab[3525]++
//line /usr/local/go/src/strings/strings.go:160
		// _ = "end of CoverTab[3525]"
//line /usr/local/go/src/strings/strings.go:160
	}
//line /usr/local/go/src/strings/strings.go:160
	// _ = "end of CoverTab[3508]"
//line /usr/local/go/src/strings/strings.go:160
	_go_fuzz_dep_.CoverTab[3509]++
							for i, c := range s {
//line /usr/local/go/src/strings/strings.go:161
		_go_fuzz_dep_.CoverTab[3526]++
								if IndexRune(chars, c) >= 0 {
//line /usr/local/go/src/strings/strings.go:162
			_go_fuzz_dep_.CoverTab[3527]++
									return i
//line /usr/local/go/src/strings/strings.go:163
			// _ = "end of CoverTab[3527]"
		} else {
//line /usr/local/go/src/strings/strings.go:164
			_go_fuzz_dep_.CoverTab[3528]++
//line /usr/local/go/src/strings/strings.go:164
			// _ = "end of CoverTab[3528]"
//line /usr/local/go/src/strings/strings.go:164
		}
//line /usr/local/go/src/strings/strings.go:164
		// _ = "end of CoverTab[3526]"
	}
//line /usr/local/go/src/strings/strings.go:165
	// _ = "end of CoverTab[3509]"
//line /usr/local/go/src/strings/strings.go:165
	_go_fuzz_dep_.CoverTab[3510]++
							return -1
//line /usr/local/go/src/strings/strings.go:166
	// _ = "end of CoverTab[3510]"
}

// LastIndexAny returns the index of the last instance of any Unicode code
//line /usr/local/go/src/strings/strings.go:169
// point from chars in s, or -1 if no Unicode code point from chars is
//line /usr/local/go/src/strings/strings.go:169
// present in s.
//line /usr/local/go/src/strings/strings.go:172
func LastIndexAny(s, chars string) int {
//line /usr/local/go/src/strings/strings.go:172
	_go_fuzz_dep_.CoverTab[3529]++
							if chars == "" {
//line /usr/local/go/src/strings/strings.go:173
		_go_fuzz_dep_.CoverTab[3535]++

								return -1
//line /usr/local/go/src/strings/strings.go:175
		// _ = "end of CoverTab[3535]"
	} else {
//line /usr/local/go/src/strings/strings.go:176
		_go_fuzz_dep_.CoverTab[3536]++
//line /usr/local/go/src/strings/strings.go:176
		// _ = "end of CoverTab[3536]"
//line /usr/local/go/src/strings/strings.go:176
	}
//line /usr/local/go/src/strings/strings.go:176
	// _ = "end of CoverTab[3529]"
//line /usr/local/go/src/strings/strings.go:176
	_go_fuzz_dep_.CoverTab[3530]++
							if len(s) == 1 {
//line /usr/local/go/src/strings/strings.go:177
		_go_fuzz_dep_.CoverTab[3537]++
								rc := rune(s[0])
								if rc >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:179
			_go_fuzz_dep_.CoverTab[3540]++
									rc = utf8.RuneError
//line /usr/local/go/src/strings/strings.go:180
			// _ = "end of CoverTab[3540]"
		} else {
//line /usr/local/go/src/strings/strings.go:181
			_go_fuzz_dep_.CoverTab[3541]++
//line /usr/local/go/src/strings/strings.go:181
			// _ = "end of CoverTab[3541]"
//line /usr/local/go/src/strings/strings.go:181
		}
//line /usr/local/go/src/strings/strings.go:181
		// _ = "end of CoverTab[3537]"
//line /usr/local/go/src/strings/strings.go:181
		_go_fuzz_dep_.CoverTab[3538]++
								if IndexRune(chars, rc) >= 0 {
//line /usr/local/go/src/strings/strings.go:182
			_go_fuzz_dep_.CoverTab[3542]++
									return 0
//line /usr/local/go/src/strings/strings.go:183
			// _ = "end of CoverTab[3542]"
		} else {
//line /usr/local/go/src/strings/strings.go:184
			_go_fuzz_dep_.CoverTab[3543]++
//line /usr/local/go/src/strings/strings.go:184
			// _ = "end of CoverTab[3543]"
//line /usr/local/go/src/strings/strings.go:184
		}
//line /usr/local/go/src/strings/strings.go:184
		// _ = "end of CoverTab[3538]"
//line /usr/local/go/src/strings/strings.go:184
		_go_fuzz_dep_.CoverTab[3539]++
								return -1
//line /usr/local/go/src/strings/strings.go:185
		// _ = "end of CoverTab[3539]"
	} else {
//line /usr/local/go/src/strings/strings.go:186
		_go_fuzz_dep_.CoverTab[3544]++
//line /usr/local/go/src/strings/strings.go:186
		// _ = "end of CoverTab[3544]"
//line /usr/local/go/src/strings/strings.go:186
	}
//line /usr/local/go/src/strings/strings.go:186
	// _ = "end of CoverTab[3530]"
//line /usr/local/go/src/strings/strings.go:186
	_go_fuzz_dep_.CoverTab[3531]++
							if len(s) > 8 {
//line /usr/local/go/src/strings/strings.go:187
		_go_fuzz_dep_.CoverTab[3545]++
								if as, isASCII := makeASCIISet(chars); isASCII {
//line /usr/local/go/src/strings/strings.go:188
			_go_fuzz_dep_.CoverTab[3546]++
									for i := len(s) - 1; i >= 0; i-- {
//line /usr/local/go/src/strings/strings.go:189
				_go_fuzz_dep_.CoverTab[3548]++
										if as.contains(s[i]) {
//line /usr/local/go/src/strings/strings.go:190
					_go_fuzz_dep_.CoverTab[3549]++
											return i
//line /usr/local/go/src/strings/strings.go:191
					// _ = "end of CoverTab[3549]"
				} else {
//line /usr/local/go/src/strings/strings.go:192
					_go_fuzz_dep_.CoverTab[3550]++
//line /usr/local/go/src/strings/strings.go:192
					// _ = "end of CoverTab[3550]"
//line /usr/local/go/src/strings/strings.go:192
				}
//line /usr/local/go/src/strings/strings.go:192
				// _ = "end of CoverTab[3548]"
			}
//line /usr/local/go/src/strings/strings.go:193
			// _ = "end of CoverTab[3546]"
//line /usr/local/go/src/strings/strings.go:193
			_go_fuzz_dep_.CoverTab[3547]++
									return -1
//line /usr/local/go/src/strings/strings.go:194
			// _ = "end of CoverTab[3547]"
		} else {
//line /usr/local/go/src/strings/strings.go:195
			_go_fuzz_dep_.CoverTab[3551]++
//line /usr/local/go/src/strings/strings.go:195
			// _ = "end of CoverTab[3551]"
//line /usr/local/go/src/strings/strings.go:195
		}
//line /usr/local/go/src/strings/strings.go:195
		// _ = "end of CoverTab[3545]"
	} else {
//line /usr/local/go/src/strings/strings.go:196
		_go_fuzz_dep_.CoverTab[3552]++
//line /usr/local/go/src/strings/strings.go:196
		// _ = "end of CoverTab[3552]"
//line /usr/local/go/src/strings/strings.go:196
	}
//line /usr/local/go/src/strings/strings.go:196
	// _ = "end of CoverTab[3531]"
//line /usr/local/go/src/strings/strings.go:196
	_go_fuzz_dep_.CoverTab[3532]++
							if len(chars) == 1 {
//line /usr/local/go/src/strings/strings.go:197
		_go_fuzz_dep_.CoverTab[3553]++
								rc := rune(chars[0])
								if rc >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:199
			_go_fuzz_dep_.CoverTab[3556]++
									rc = utf8.RuneError
//line /usr/local/go/src/strings/strings.go:200
			// _ = "end of CoverTab[3556]"
		} else {
//line /usr/local/go/src/strings/strings.go:201
			_go_fuzz_dep_.CoverTab[3557]++
//line /usr/local/go/src/strings/strings.go:201
			// _ = "end of CoverTab[3557]"
//line /usr/local/go/src/strings/strings.go:201
		}
//line /usr/local/go/src/strings/strings.go:201
		// _ = "end of CoverTab[3553]"
//line /usr/local/go/src/strings/strings.go:201
		_go_fuzz_dep_.CoverTab[3554]++
								for i := len(s); i > 0; {
//line /usr/local/go/src/strings/strings.go:202
			_go_fuzz_dep_.CoverTab[3558]++
									r, size := utf8.DecodeLastRuneInString(s[:i])
									i -= size
									if rc == r {
//line /usr/local/go/src/strings/strings.go:205
				_go_fuzz_dep_.CoverTab[3559]++
										return i
//line /usr/local/go/src/strings/strings.go:206
				// _ = "end of CoverTab[3559]"
			} else {
//line /usr/local/go/src/strings/strings.go:207
				_go_fuzz_dep_.CoverTab[3560]++
//line /usr/local/go/src/strings/strings.go:207
				// _ = "end of CoverTab[3560]"
//line /usr/local/go/src/strings/strings.go:207
			}
//line /usr/local/go/src/strings/strings.go:207
			// _ = "end of CoverTab[3558]"
		}
//line /usr/local/go/src/strings/strings.go:208
		// _ = "end of CoverTab[3554]"
//line /usr/local/go/src/strings/strings.go:208
		_go_fuzz_dep_.CoverTab[3555]++
								return -1
//line /usr/local/go/src/strings/strings.go:209
		// _ = "end of CoverTab[3555]"
	} else {
//line /usr/local/go/src/strings/strings.go:210
		_go_fuzz_dep_.CoverTab[3561]++
//line /usr/local/go/src/strings/strings.go:210
		// _ = "end of CoverTab[3561]"
//line /usr/local/go/src/strings/strings.go:210
	}
//line /usr/local/go/src/strings/strings.go:210
	// _ = "end of CoverTab[3532]"
//line /usr/local/go/src/strings/strings.go:210
	_go_fuzz_dep_.CoverTab[3533]++
							for i := len(s); i > 0; {
//line /usr/local/go/src/strings/strings.go:211
		_go_fuzz_dep_.CoverTab[3562]++
								r, size := utf8.DecodeLastRuneInString(s[:i])
								i -= size
								if IndexRune(chars, r) >= 0 {
//line /usr/local/go/src/strings/strings.go:214
			_go_fuzz_dep_.CoverTab[3563]++
									return i
//line /usr/local/go/src/strings/strings.go:215
			// _ = "end of CoverTab[3563]"
		} else {
//line /usr/local/go/src/strings/strings.go:216
			_go_fuzz_dep_.CoverTab[3564]++
//line /usr/local/go/src/strings/strings.go:216
			// _ = "end of CoverTab[3564]"
//line /usr/local/go/src/strings/strings.go:216
		}
//line /usr/local/go/src/strings/strings.go:216
		// _ = "end of CoverTab[3562]"
	}
//line /usr/local/go/src/strings/strings.go:217
	// _ = "end of CoverTab[3533]"
//line /usr/local/go/src/strings/strings.go:217
	_go_fuzz_dep_.CoverTab[3534]++
							return -1
//line /usr/local/go/src/strings/strings.go:218
	// _ = "end of CoverTab[3534]"
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte(s string, c byte) int {
//line /usr/local/go/src/strings/strings.go:222
	_go_fuzz_dep_.CoverTab[3565]++
							for i := len(s) - 1; i >= 0; i-- {
//line /usr/local/go/src/strings/strings.go:223
		_go_fuzz_dep_.CoverTab[3567]++
								if s[i] == c {
//line /usr/local/go/src/strings/strings.go:224
			_go_fuzz_dep_.CoverTab[3568]++
									return i
//line /usr/local/go/src/strings/strings.go:225
			// _ = "end of CoverTab[3568]"
		} else {
//line /usr/local/go/src/strings/strings.go:226
			_go_fuzz_dep_.CoverTab[3569]++
//line /usr/local/go/src/strings/strings.go:226
			// _ = "end of CoverTab[3569]"
//line /usr/local/go/src/strings/strings.go:226
		}
//line /usr/local/go/src/strings/strings.go:226
		// _ = "end of CoverTab[3567]"
	}
//line /usr/local/go/src/strings/strings.go:227
	// _ = "end of CoverTab[3565]"
//line /usr/local/go/src/strings/strings.go:227
	_go_fuzz_dep_.CoverTab[3566]++
							return -1
//line /usr/local/go/src/strings/strings.go:228
	// _ = "end of CoverTab[3566]"
}

// Generic split: splits after each instance of sep,
//line /usr/local/go/src/strings/strings.go:231
// including sepSave bytes of sep in the subarrays.
//line /usr/local/go/src/strings/strings.go:233
func genSplit(s, sep string, sepSave, n int) []string {
//line /usr/local/go/src/strings/strings.go:233
	_go_fuzz_dep_.CoverTab[3570]++
							if n == 0 {
//line /usr/local/go/src/strings/strings.go:234
		_go_fuzz_dep_.CoverTab[3576]++
								return nil
//line /usr/local/go/src/strings/strings.go:235
		// _ = "end of CoverTab[3576]"
	} else {
//line /usr/local/go/src/strings/strings.go:236
		_go_fuzz_dep_.CoverTab[3577]++
//line /usr/local/go/src/strings/strings.go:236
		// _ = "end of CoverTab[3577]"
//line /usr/local/go/src/strings/strings.go:236
	}
//line /usr/local/go/src/strings/strings.go:236
	// _ = "end of CoverTab[3570]"
//line /usr/local/go/src/strings/strings.go:236
	_go_fuzz_dep_.CoverTab[3571]++
							if sep == "" {
//line /usr/local/go/src/strings/strings.go:237
		_go_fuzz_dep_.CoverTab[3578]++
								return explode(s, n)
//line /usr/local/go/src/strings/strings.go:238
		// _ = "end of CoverTab[3578]"
	} else {
//line /usr/local/go/src/strings/strings.go:239
		_go_fuzz_dep_.CoverTab[3579]++
//line /usr/local/go/src/strings/strings.go:239
		// _ = "end of CoverTab[3579]"
//line /usr/local/go/src/strings/strings.go:239
	}
//line /usr/local/go/src/strings/strings.go:239
	// _ = "end of CoverTab[3571]"
//line /usr/local/go/src/strings/strings.go:239
	_go_fuzz_dep_.CoverTab[3572]++
							if n < 0 {
//line /usr/local/go/src/strings/strings.go:240
		_go_fuzz_dep_.CoverTab[3580]++
								n = Count(s, sep) + 1
//line /usr/local/go/src/strings/strings.go:241
		// _ = "end of CoverTab[3580]"
	} else {
//line /usr/local/go/src/strings/strings.go:242
		_go_fuzz_dep_.CoverTab[3581]++
//line /usr/local/go/src/strings/strings.go:242
		// _ = "end of CoverTab[3581]"
//line /usr/local/go/src/strings/strings.go:242
	}
//line /usr/local/go/src/strings/strings.go:242
	// _ = "end of CoverTab[3572]"
//line /usr/local/go/src/strings/strings.go:242
	_go_fuzz_dep_.CoverTab[3573]++

							if n > len(s)+1 {
//line /usr/local/go/src/strings/strings.go:244
		_go_fuzz_dep_.CoverTab[3582]++
								n = len(s) + 1
//line /usr/local/go/src/strings/strings.go:245
		// _ = "end of CoverTab[3582]"
	} else {
//line /usr/local/go/src/strings/strings.go:246
		_go_fuzz_dep_.CoverTab[3583]++
//line /usr/local/go/src/strings/strings.go:246
		// _ = "end of CoverTab[3583]"
//line /usr/local/go/src/strings/strings.go:246
	}
//line /usr/local/go/src/strings/strings.go:246
	// _ = "end of CoverTab[3573]"
//line /usr/local/go/src/strings/strings.go:246
	_go_fuzz_dep_.CoverTab[3574]++
							a := make([]string, n)
							n--
							i := 0
							for i < n {
//line /usr/local/go/src/strings/strings.go:250
		_go_fuzz_dep_.CoverTab[3584]++
								m := Index(s, sep)
								if m < 0 {
//line /usr/local/go/src/strings/strings.go:252
			_go_fuzz_dep_.CoverTab[3586]++
									break
//line /usr/local/go/src/strings/strings.go:253
			// _ = "end of CoverTab[3586]"
		} else {
//line /usr/local/go/src/strings/strings.go:254
			_go_fuzz_dep_.CoverTab[3587]++
//line /usr/local/go/src/strings/strings.go:254
			// _ = "end of CoverTab[3587]"
//line /usr/local/go/src/strings/strings.go:254
		}
//line /usr/local/go/src/strings/strings.go:254
		// _ = "end of CoverTab[3584]"
//line /usr/local/go/src/strings/strings.go:254
		_go_fuzz_dep_.CoverTab[3585]++
								a[i] = s[:m+sepSave]
								s = s[m+len(sep):]
								i++
//line /usr/local/go/src/strings/strings.go:257
		// _ = "end of CoverTab[3585]"
	}
//line /usr/local/go/src/strings/strings.go:258
	// _ = "end of CoverTab[3574]"
//line /usr/local/go/src/strings/strings.go:258
	_go_fuzz_dep_.CoverTab[3575]++
							a[i] = s
							return a[:i+1]
//line /usr/local/go/src/strings/strings.go:260
	// _ = "end of CoverTab[3575]"
}

// SplitN slices s into substrings separated by sep and returns a slice of
//line /usr/local/go/src/strings/strings.go:263
// the substrings between those separators.
//line /usr/local/go/src/strings/strings.go:263
//
//line /usr/local/go/src/strings/strings.go:263
// The count determines the number of substrings to return:
//line /usr/local/go/src/strings/strings.go:263
//
//line /usr/local/go/src/strings/strings.go:263
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//line /usr/local/go/src/strings/strings.go:263
//	n == 0: the result is nil (zero substrings)
//line /usr/local/go/src/strings/strings.go:263
//	n < 0: all substrings
//line /usr/local/go/src/strings/strings.go:263
//
//line /usr/local/go/src/strings/strings.go:263
// Edge cases for s and sep (for example, empty strings) are handled
//line /usr/local/go/src/strings/strings.go:263
// as described in the documentation for Split.
//line /usr/local/go/src/strings/strings.go:263
//
//line /usr/local/go/src/strings/strings.go:263
// To split around the first instance of a separator, see Cut.
//line /usr/local/go/src/strings/strings.go:276
func SplitN(s, sep string, n int) []string {
//line /usr/local/go/src/strings/strings.go:276
	_go_fuzz_dep_.CoverTab[3588]++
//line /usr/local/go/src/strings/strings.go:276
	return genSplit(s, sep, 0, n)
//line /usr/local/go/src/strings/strings.go:276
	// _ = "end of CoverTab[3588]"
//line /usr/local/go/src/strings/strings.go:276
}

// SplitAfterN slices s into substrings after each instance of sep and
//line /usr/local/go/src/strings/strings.go:278
// returns a slice of those substrings.
//line /usr/local/go/src/strings/strings.go:278
//
//line /usr/local/go/src/strings/strings.go:278
// The count determines the number of substrings to return:
//line /usr/local/go/src/strings/strings.go:278
//
//line /usr/local/go/src/strings/strings.go:278
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//line /usr/local/go/src/strings/strings.go:278
//	n == 0: the result is nil (zero substrings)
//line /usr/local/go/src/strings/strings.go:278
//	n < 0: all substrings
//line /usr/local/go/src/strings/strings.go:278
//
//line /usr/local/go/src/strings/strings.go:278
// Edge cases for s and sep (for example, empty strings) are handled
//line /usr/local/go/src/strings/strings.go:278
// as described in the documentation for SplitAfter.
//line /usr/local/go/src/strings/strings.go:289
func SplitAfterN(s, sep string, n int) []string {
//line /usr/local/go/src/strings/strings.go:289
	_go_fuzz_dep_.CoverTab[3589]++
							return genSplit(s, sep, len(sep), n)
//line /usr/local/go/src/strings/strings.go:290
	// _ = "end of CoverTab[3589]"
}

// Split slices s into all substrings separated by sep and returns a slice of
//line /usr/local/go/src/strings/strings.go:293
// the substrings between those separators.
//line /usr/local/go/src/strings/strings.go:293
//
//line /usr/local/go/src/strings/strings.go:293
// If s does not contain sep and sep is not empty, Split returns a
//line /usr/local/go/src/strings/strings.go:293
// slice of length 1 whose only element is s.
//line /usr/local/go/src/strings/strings.go:293
//
//line /usr/local/go/src/strings/strings.go:293
// If sep is empty, Split splits after each UTF-8 sequence. If both s
//line /usr/local/go/src/strings/strings.go:293
// and sep are empty, Split returns an empty slice.
//line /usr/local/go/src/strings/strings.go:293
//
//line /usr/local/go/src/strings/strings.go:293
// It is equivalent to SplitN with a count of -1.
//line /usr/local/go/src/strings/strings.go:293
//
//line /usr/local/go/src/strings/strings.go:293
// To split around the first instance of a separator, see Cut.
//line /usr/local/go/src/strings/strings.go:305
func Split(s, sep string) []string {
//line /usr/local/go/src/strings/strings.go:305
	_go_fuzz_dep_.CoverTab[3590]++
//line /usr/local/go/src/strings/strings.go:305
	return genSplit(s, sep, 0, -1)
//line /usr/local/go/src/strings/strings.go:305
	// _ = "end of CoverTab[3590]"
//line /usr/local/go/src/strings/strings.go:305
}

// SplitAfter slices s into all substrings after each instance of sep and
//line /usr/local/go/src/strings/strings.go:307
// returns a slice of those substrings.
//line /usr/local/go/src/strings/strings.go:307
//
//line /usr/local/go/src/strings/strings.go:307
// If s does not contain sep and sep is not empty, SplitAfter returns
//line /usr/local/go/src/strings/strings.go:307
// a slice of length 1 whose only element is s.
//line /usr/local/go/src/strings/strings.go:307
//
//line /usr/local/go/src/strings/strings.go:307
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
//line /usr/local/go/src/strings/strings.go:307
// both s and sep are empty, SplitAfter returns an empty slice.
//line /usr/local/go/src/strings/strings.go:307
//
//line /usr/local/go/src/strings/strings.go:307
// It is equivalent to SplitAfterN with a count of -1.
//line /usr/local/go/src/strings/strings.go:317
func SplitAfter(s, sep string) []string {
//line /usr/local/go/src/strings/strings.go:317
	_go_fuzz_dep_.CoverTab[3591]++
							return genSplit(s, sep, len(sep), -1)
//line /usr/local/go/src/strings/strings.go:318
	// _ = "end of CoverTab[3591]"
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// Fields splits the string s around each instance of one or more consecutive white space
//line /usr/local/go/src/strings/strings.go:323
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
//line /usr/local/go/src/strings/strings.go:323
// empty slice if s contains only white space.
//line /usr/local/go/src/strings/strings.go:326
func Fields(s string) []string {
//line /usr/local/go/src/strings/strings.go:326
	_go_fuzz_dep_.CoverTab[3592]++

//line /usr/local/go/src/strings/strings.go:329
	n := 0
	wasSpace := 1

	setBits := uint8(0)
	for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:333
		_go_fuzz_dep_.CoverTab[3598]++
								r := s[i]
								setBits |= r
								isSpace := int(asciiSpace[r])
								n += wasSpace & ^isSpace
								wasSpace = isSpace
//line /usr/local/go/src/strings/strings.go:338
		// _ = "end of CoverTab[3598]"
	}
//line /usr/local/go/src/strings/strings.go:339
	// _ = "end of CoverTab[3592]"
//line /usr/local/go/src/strings/strings.go:339
	_go_fuzz_dep_.CoverTab[3593]++

							if setBits >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:341
		_go_fuzz_dep_.CoverTab[3599]++

								return FieldsFunc(s, unicode.IsSpace)
//line /usr/local/go/src/strings/strings.go:343
		// _ = "end of CoverTab[3599]"
	} else {
//line /usr/local/go/src/strings/strings.go:344
		_go_fuzz_dep_.CoverTab[3600]++
//line /usr/local/go/src/strings/strings.go:344
		// _ = "end of CoverTab[3600]"
//line /usr/local/go/src/strings/strings.go:344
	}
//line /usr/local/go/src/strings/strings.go:344
	// _ = "end of CoverTab[3593]"
//line /usr/local/go/src/strings/strings.go:344
	_go_fuzz_dep_.CoverTab[3594]++

							a := make([]string, n)
							na := 0
							fieldStart := 0
							i := 0

							for i < len(s) && func() bool {
//line /usr/local/go/src/strings/strings.go:351
		_go_fuzz_dep_.CoverTab[3601]++
//line /usr/local/go/src/strings/strings.go:351
		return asciiSpace[s[i]] != 0
//line /usr/local/go/src/strings/strings.go:351
		// _ = "end of CoverTab[3601]"
//line /usr/local/go/src/strings/strings.go:351
	}() {
//line /usr/local/go/src/strings/strings.go:351
		_go_fuzz_dep_.CoverTab[3602]++
								i++
//line /usr/local/go/src/strings/strings.go:352
		// _ = "end of CoverTab[3602]"
	}
//line /usr/local/go/src/strings/strings.go:353
	// _ = "end of CoverTab[3594]"
//line /usr/local/go/src/strings/strings.go:353
	_go_fuzz_dep_.CoverTab[3595]++
							fieldStart = i
							for i < len(s) {
//line /usr/local/go/src/strings/strings.go:355
		_go_fuzz_dep_.CoverTab[3603]++
								if asciiSpace[s[i]] == 0 {
//line /usr/local/go/src/strings/strings.go:356
			_go_fuzz_dep_.CoverTab[3606]++
									i++
									continue
//line /usr/local/go/src/strings/strings.go:358
			// _ = "end of CoverTab[3606]"
		} else {
//line /usr/local/go/src/strings/strings.go:359
			_go_fuzz_dep_.CoverTab[3607]++
//line /usr/local/go/src/strings/strings.go:359
			// _ = "end of CoverTab[3607]"
//line /usr/local/go/src/strings/strings.go:359
		}
//line /usr/local/go/src/strings/strings.go:359
		// _ = "end of CoverTab[3603]"
//line /usr/local/go/src/strings/strings.go:359
		_go_fuzz_dep_.CoverTab[3604]++
								a[na] = s[fieldStart:i]
								na++
								i++

								for i < len(s) && func() bool {
//line /usr/local/go/src/strings/strings.go:364
			_go_fuzz_dep_.CoverTab[3608]++
//line /usr/local/go/src/strings/strings.go:364
			return asciiSpace[s[i]] != 0
//line /usr/local/go/src/strings/strings.go:364
			// _ = "end of CoverTab[3608]"
//line /usr/local/go/src/strings/strings.go:364
		}() {
//line /usr/local/go/src/strings/strings.go:364
			_go_fuzz_dep_.CoverTab[3609]++
									i++
//line /usr/local/go/src/strings/strings.go:365
			// _ = "end of CoverTab[3609]"
		}
//line /usr/local/go/src/strings/strings.go:366
		// _ = "end of CoverTab[3604]"
//line /usr/local/go/src/strings/strings.go:366
		_go_fuzz_dep_.CoverTab[3605]++
								fieldStart = i
//line /usr/local/go/src/strings/strings.go:367
		// _ = "end of CoverTab[3605]"
	}
//line /usr/local/go/src/strings/strings.go:368
	// _ = "end of CoverTab[3595]"
//line /usr/local/go/src/strings/strings.go:368
	_go_fuzz_dep_.CoverTab[3596]++
							if fieldStart < len(s) {
//line /usr/local/go/src/strings/strings.go:369
		_go_fuzz_dep_.CoverTab[3610]++
								a[na] = s[fieldStart:]
//line /usr/local/go/src/strings/strings.go:370
		// _ = "end of CoverTab[3610]"
	} else {
//line /usr/local/go/src/strings/strings.go:371
		_go_fuzz_dep_.CoverTab[3611]++
//line /usr/local/go/src/strings/strings.go:371
		// _ = "end of CoverTab[3611]"
//line /usr/local/go/src/strings/strings.go:371
	}
//line /usr/local/go/src/strings/strings.go:371
	// _ = "end of CoverTab[3596]"
//line /usr/local/go/src/strings/strings.go:371
	_go_fuzz_dep_.CoverTab[3597]++
							return a
//line /usr/local/go/src/strings/strings.go:372
	// _ = "end of CoverTab[3597]"
}

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
//line /usr/local/go/src/strings/strings.go:375
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
//line /usr/local/go/src/strings/strings.go:375
// string is empty, an empty slice is returned.
//line /usr/local/go/src/strings/strings.go:375
//
//line /usr/local/go/src/strings/strings.go:375
// FieldsFunc makes no guarantees about the order in which it calls f(c)
//line /usr/local/go/src/strings/strings.go:375
// and assumes that f always returns the same value for a given c.
//line /usr/local/go/src/strings/strings.go:381
func FieldsFunc(s string, f func(rune) bool) []string {
//line /usr/local/go/src/strings/strings.go:381
	_go_fuzz_dep_.CoverTab[3612]++
	// A span is used to record a slice of s of the form s[start:end].
	// The start index is inclusive and the end index is exclusive.
	type span struct {
		start	int
		end	int
	}
							spans := make([]span, 0, 32)

//line /usr/local/go/src/strings/strings.go:394
	start := -1
	for end, rune := range s {
//line /usr/local/go/src/strings/strings.go:395
		_go_fuzz_dep_.CoverTab[3616]++
								if f(rune) {
//line /usr/local/go/src/strings/strings.go:396
			_go_fuzz_dep_.CoverTab[3617]++
									if start >= 0 {
//line /usr/local/go/src/strings/strings.go:397
				_go_fuzz_dep_.CoverTab[3618]++
										spans = append(spans, span{start, end})

//line /usr/local/go/src/strings/strings.go:402
				start = ^start
//line /usr/local/go/src/strings/strings.go:402
				// _ = "end of CoverTab[3618]"
			} else {
//line /usr/local/go/src/strings/strings.go:403
				_go_fuzz_dep_.CoverTab[3619]++
//line /usr/local/go/src/strings/strings.go:403
				// _ = "end of CoverTab[3619]"
//line /usr/local/go/src/strings/strings.go:403
			}
//line /usr/local/go/src/strings/strings.go:403
			// _ = "end of CoverTab[3617]"
		} else {
//line /usr/local/go/src/strings/strings.go:404
			_go_fuzz_dep_.CoverTab[3620]++
									if start < 0 {
//line /usr/local/go/src/strings/strings.go:405
				_go_fuzz_dep_.CoverTab[3621]++
										start = end
//line /usr/local/go/src/strings/strings.go:406
				// _ = "end of CoverTab[3621]"
			} else {
//line /usr/local/go/src/strings/strings.go:407
				_go_fuzz_dep_.CoverTab[3622]++
//line /usr/local/go/src/strings/strings.go:407
				// _ = "end of CoverTab[3622]"
//line /usr/local/go/src/strings/strings.go:407
			}
//line /usr/local/go/src/strings/strings.go:407
			// _ = "end of CoverTab[3620]"
		}
//line /usr/local/go/src/strings/strings.go:408
		// _ = "end of CoverTab[3616]"
	}
//line /usr/local/go/src/strings/strings.go:409
	// _ = "end of CoverTab[3612]"
//line /usr/local/go/src/strings/strings.go:409
	_go_fuzz_dep_.CoverTab[3613]++

//line /usr/local/go/src/strings/strings.go:412
	if start >= 0 {
//line /usr/local/go/src/strings/strings.go:412
		_go_fuzz_dep_.CoverTab[3623]++
								spans = append(spans, span{start, len(s)})
//line /usr/local/go/src/strings/strings.go:413
		// _ = "end of CoverTab[3623]"
	} else {
//line /usr/local/go/src/strings/strings.go:414
		_go_fuzz_dep_.CoverTab[3624]++
//line /usr/local/go/src/strings/strings.go:414
		// _ = "end of CoverTab[3624]"
//line /usr/local/go/src/strings/strings.go:414
	}
//line /usr/local/go/src/strings/strings.go:414
	// _ = "end of CoverTab[3613]"
//line /usr/local/go/src/strings/strings.go:414
	_go_fuzz_dep_.CoverTab[3614]++

//line /usr/local/go/src/strings/strings.go:417
	a := make([]string, len(spans))
	for i, span := range spans {
//line /usr/local/go/src/strings/strings.go:418
		_go_fuzz_dep_.CoverTab[3625]++
								a[i] = s[span.start:span.end]
//line /usr/local/go/src/strings/strings.go:419
		// _ = "end of CoverTab[3625]"
	}
//line /usr/local/go/src/strings/strings.go:420
	// _ = "end of CoverTab[3614]"
//line /usr/local/go/src/strings/strings.go:420
	_go_fuzz_dep_.CoverTab[3615]++

							return a
//line /usr/local/go/src/strings/strings.go:422
	// _ = "end of CoverTab[3615]"
}

// Join concatenates the elements of its first argument to create a single string. The separator
//line /usr/local/go/src/strings/strings.go:425
// string sep is placed between elements in the resulting string.
//line /usr/local/go/src/strings/strings.go:427
func Join(elems []string, sep string) string {
//line /usr/local/go/src/strings/strings.go:427
	_go_fuzz_dep_.CoverTab[3626]++
							switch len(elems) {
	case 0:
//line /usr/local/go/src/strings/strings.go:429
		_go_fuzz_dep_.CoverTab[3630]++
								return ""
//line /usr/local/go/src/strings/strings.go:430
		// _ = "end of CoverTab[3630]"
	case 1:
//line /usr/local/go/src/strings/strings.go:431
		_go_fuzz_dep_.CoverTab[3631]++
								return elems[0]
//line /usr/local/go/src/strings/strings.go:432
		// _ = "end of CoverTab[3631]"
//line /usr/local/go/src/strings/strings.go:432
	default:
//line /usr/local/go/src/strings/strings.go:432
		_go_fuzz_dep_.CoverTab[3632]++
//line /usr/local/go/src/strings/strings.go:432
		// _ = "end of CoverTab[3632]"
	}
//line /usr/local/go/src/strings/strings.go:433
	// _ = "end of CoverTab[3626]"
//line /usr/local/go/src/strings/strings.go:433
	_go_fuzz_dep_.CoverTab[3627]++
							n := len(sep) * (len(elems) - 1)
							for i := 0; i < len(elems); i++ {
//line /usr/local/go/src/strings/strings.go:435
		_go_fuzz_dep_.CoverTab[3633]++
								n += len(elems[i])
//line /usr/local/go/src/strings/strings.go:436
		// _ = "end of CoverTab[3633]"
	}
//line /usr/local/go/src/strings/strings.go:437
	// _ = "end of CoverTab[3627]"
//line /usr/local/go/src/strings/strings.go:437
	_go_fuzz_dep_.CoverTab[3628]++

							var b Builder
							b.Grow(n)
							b.WriteString(elems[0])
							for _, s := range elems[1:] {
//line /usr/local/go/src/strings/strings.go:442
		_go_fuzz_dep_.CoverTab[3634]++
								b.WriteString(sep)
								b.WriteString(s)
//line /usr/local/go/src/strings/strings.go:444
		// _ = "end of CoverTab[3634]"
	}
//line /usr/local/go/src/strings/strings.go:445
	// _ = "end of CoverTab[3628]"
//line /usr/local/go/src/strings/strings.go:445
	_go_fuzz_dep_.CoverTab[3629]++
							return b.String()
//line /usr/local/go/src/strings/strings.go:446
	// _ = "end of CoverTab[3629]"
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool {
//line /usr/local/go/src/strings/strings.go:450
	_go_fuzz_dep_.CoverTab[3635]++
							return len(s) >= len(prefix) && func() bool {
//line /usr/local/go/src/strings/strings.go:451
		_go_fuzz_dep_.CoverTab[3636]++
//line /usr/local/go/src/strings/strings.go:451
		return s[0:len(prefix)] == prefix
//line /usr/local/go/src/strings/strings.go:451
		// _ = "end of CoverTab[3636]"
//line /usr/local/go/src/strings/strings.go:451
	}()
//line /usr/local/go/src/strings/strings.go:451
	// _ = "end of CoverTab[3635]"
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
//line /usr/local/go/src/strings/strings.go:455
	_go_fuzz_dep_.CoverTab[3637]++
							return len(s) >= len(suffix) && func() bool {
//line /usr/local/go/src/strings/strings.go:456
		_go_fuzz_dep_.CoverTab[3638]++
//line /usr/local/go/src/strings/strings.go:456
		return s[len(s)-len(suffix):] == suffix
//line /usr/local/go/src/strings/strings.go:456
		// _ = "end of CoverTab[3638]"
//line /usr/local/go/src/strings/strings.go:456
	}()
//line /usr/local/go/src/strings/strings.go:456
	// _ = "end of CoverTab[3637]"
}

// Map returns a copy of the string s with all its characters modified
//line /usr/local/go/src/strings/strings.go:459
// according to the mapping function. If mapping returns a negative value, the character is
//line /usr/local/go/src/strings/strings.go:459
// dropped from the string with no replacement.
//line /usr/local/go/src/strings/strings.go:462
func Map(mapping func(rune) rune, s string) string {
//line /usr/local/go/src/strings/strings.go:462
	_go_fuzz_dep_.CoverTab[3639]++

//line /usr/local/go/src/strings/strings.go:467
	// The output buffer b is initialized on demand, the first
	// time a character differs.
	var b Builder

	for i, c := range s {
//line /usr/local/go/src/strings/strings.go:471
		_go_fuzz_dep_.CoverTab[3643]++
								r := mapping(c)
								if r == c && func() bool {
//line /usr/local/go/src/strings/strings.go:473
			_go_fuzz_dep_.CoverTab[3647]++
//line /usr/local/go/src/strings/strings.go:473
			return c != utf8.RuneError
//line /usr/local/go/src/strings/strings.go:473
			// _ = "end of CoverTab[3647]"
//line /usr/local/go/src/strings/strings.go:473
		}() {
//line /usr/local/go/src/strings/strings.go:473
			_go_fuzz_dep_.CoverTab[3648]++
									continue
//line /usr/local/go/src/strings/strings.go:474
			// _ = "end of CoverTab[3648]"
		} else {
//line /usr/local/go/src/strings/strings.go:475
			_go_fuzz_dep_.CoverTab[3649]++
//line /usr/local/go/src/strings/strings.go:475
			// _ = "end of CoverTab[3649]"
//line /usr/local/go/src/strings/strings.go:475
		}
//line /usr/local/go/src/strings/strings.go:475
		// _ = "end of CoverTab[3643]"
//line /usr/local/go/src/strings/strings.go:475
		_go_fuzz_dep_.CoverTab[3644]++

								var width int
								if c == utf8.RuneError {
//line /usr/local/go/src/strings/strings.go:478
			_go_fuzz_dep_.CoverTab[3650]++
									c, width = utf8.DecodeRuneInString(s[i:])
									if width != 1 && func() bool {
//line /usr/local/go/src/strings/strings.go:480
				_go_fuzz_dep_.CoverTab[3651]++
//line /usr/local/go/src/strings/strings.go:480
				return r == c
//line /usr/local/go/src/strings/strings.go:480
				// _ = "end of CoverTab[3651]"
//line /usr/local/go/src/strings/strings.go:480
			}() {
//line /usr/local/go/src/strings/strings.go:480
				_go_fuzz_dep_.CoverTab[3652]++
										continue
//line /usr/local/go/src/strings/strings.go:481
				// _ = "end of CoverTab[3652]"
			} else {
//line /usr/local/go/src/strings/strings.go:482
				_go_fuzz_dep_.CoverTab[3653]++
//line /usr/local/go/src/strings/strings.go:482
				// _ = "end of CoverTab[3653]"
//line /usr/local/go/src/strings/strings.go:482
			}
//line /usr/local/go/src/strings/strings.go:482
			// _ = "end of CoverTab[3650]"
		} else {
//line /usr/local/go/src/strings/strings.go:483
			_go_fuzz_dep_.CoverTab[3654]++
									width = utf8.RuneLen(c)
//line /usr/local/go/src/strings/strings.go:484
			// _ = "end of CoverTab[3654]"
		}
//line /usr/local/go/src/strings/strings.go:485
		// _ = "end of CoverTab[3644]"
//line /usr/local/go/src/strings/strings.go:485
		_go_fuzz_dep_.CoverTab[3645]++

								b.Grow(len(s) + utf8.UTFMax)
								b.WriteString(s[:i])
								if r >= 0 {
//line /usr/local/go/src/strings/strings.go:489
			_go_fuzz_dep_.CoverTab[3655]++
									b.WriteRune(r)
//line /usr/local/go/src/strings/strings.go:490
			// _ = "end of CoverTab[3655]"
		} else {
//line /usr/local/go/src/strings/strings.go:491
			_go_fuzz_dep_.CoverTab[3656]++
//line /usr/local/go/src/strings/strings.go:491
			// _ = "end of CoverTab[3656]"
//line /usr/local/go/src/strings/strings.go:491
		}
//line /usr/local/go/src/strings/strings.go:491
		// _ = "end of CoverTab[3645]"
//line /usr/local/go/src/strings/strings.go:491
		_go_fuzz_dep_.CoverTab[3646]++

								s = s[i+width:]
								break
//line /usr/local/go/src/strings/strings.go:494
		// _ = "end of CoverTab[3646]"
	}
//line /usr/local/go/src/strings/strings.go:495
	// _ = "end of CoverTab[3639]"
//line /usr/local/go/src/strings/strings.go:495
	_go_fuzz_dep_.CoverTab[3640]++

//line /usr/local/go/src/strings/strings.go:498
	if b.Cap() == 0 {
//line /usr/local/go/src/strings/strings.go:498
		_go_fuzz_dep_.CoverTab[3657]++
								return s
//line /usr/local/go/src/strings/strings.go:499
		// _ = "end of CoverTab[3657]"
	} else {
//line /usr/local/go/src/strings/strings.go:500
		_go_fuzz_dep_.CoverTab[3658]++
//line /usr/local/go/src/strings/strings.go:500
		// _ = "end of CoverTab[3658]"
//line /usr/local/go/src/strings/strings.go:500
	}
//line /usr/local/go/src/strings/strings.go:500
	// _ = "end of CoverTab[3640]"
//line /usr/local/go/src/strings/strings.go:500
	_go_fuzz_dep_.CoverTab[3641]++

							for _, c := range s {
//line /usr/local/go/src/strings/strings.go:502
		_go_fuzz_dep_.CoverTab[3659]++
								r := mapping(c)

								if r >= 0 {
//line /usr/local/go/src/strings/strings.go:505
			_go_fuzz_dep_.CoverTab[3660]++

//line /usr/local/go/src/strings/strings.go:509
			if r < utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:509
				_go_fuzz_dep_.CoverTab[3661]++
										b.WriteByte(byte(r))
//line /usr/local/go/src/strings/strings.go:510
				// _ = "end of CoverTab[3661]"
			} else {
//line /usr/local/go/src/strings/strings.go:511
				_go_fuzz_dep_.CoverTab[3662]++

										b.WriteRune(r)
//line /usr/local/go/src/strings/strings.go:513
				// _ = "end of CoverTab[3662]"
			}
//line /usr/local/go/src/strings/strings.go:514
			// _ = "end of CoverTab[3660]"
		} else {
//line /usr/local/go/src/strings/strings.go:515
			_go_fuzz_dep_.CoverTab[3663]++
//line /usr/local/go/src/strings/strings.go:515
			// _ = "end of CoverTab[3663]"
//line /usr/local/go/src/strings/strings.go:515
		}
//line /usr/local/go/src/strings/strings.go:515
		// _ = "end of CoverTab[3659]"
	}
//line /usr/local/go/src/strings/strings.go:516
	// _ = "end of CoverTab[3641]"
//line /usr/local/go/src/strings/strings.go:516
	_go_fuzz_dep_.CoverTab[3642]++

							return b.String()
//line /usr/local/go/src/strings/strings.go:518
	// _ = "end of CoverTab[3642]"
}

// Repeat returns a new string consisting of count copies of the string s.
//line /usr/local/go/src/strings/strings.go:521
//
//line /usr/local/go/src/strings/strings.go:521
// It panics if count is negative or if the result of (len(s) * count)
//line /usr/local/go/src/strings/strings.go:521
// overflows.
//line /usr/local/go/src/strings/strings.go:525
func Repeat(s string, count int) string {
//line /usr/local/go/src/strings/strings.go:525
	_go_fuzz_dep_.CoverTab[3664]++
							switch count {
	case 0:
//line /usr/local/go/src/strings/strings.go:527
		_go_fuzz_dep_.CoverTab[3670]++
								return ""
//line /usr/local/go/src/strings/strings.go:528
		// _ = "end of CoverTab[3670]"
	case 1:
//line /usr/local/go/src/strings/strings.go:529
		_go_fuzz_dep_.CoverTab[3671]++
								return s
//line /usr/local/go/src/strings/strings.go:530
		// _ = "end of CoverTab[3671]"
//line /usr/local/go/src/strings/strings.go:530
	default:
//line /usr/local/go/src/strings/strings.go:530
		_go_fuzz_dep_.CoverTab[3672]++
//line /usr/local/go/src/strings/strings.go:530
		// _ = "end of CoverTab[3672]"
	}
//line /usr/local/go/src/strings/strings.go:531
	// _ = "end of CoverTab[3664]"
//line /usr/local/go/src/strings/strings.go:531
	_go_fuzz_dep_.CoverTab[3665]++

//line /usr/local/go/src/strings/strings.go:537
	if count < 0 {
//line /usr/local/go/src/strings/strings.go:537
		_go_fuzz_dep_.CoverTab[3673]++
								panic("strings: negative Repeat count")
//line /usr/local/go/src/strings/strings.go:538
		// _ = "end of CoverTab[3673]"
	} else {
//line /usr/local/go/src/strings/strings.go:539
		_go_fuzz_dep_.CoverTab[3674]++
//line /usr/local/go/src/strings/strings.go:539
		if len(s)*count/count != len(s) {
//line /usr/local/go/src/strings/strings.go:539
			_go_fuzz_dep_.CoverTab[3675]++
									panic("strings: Repeat count causes overflow")
//line /usr/local/go/src/strings/strings.go:540
			// _ = "end of CoverTab[3675]"
		} else {
//line /usr/local/go/src/strings/strings.go:541
			_go_fuzz_dep_.CoverTab[3676]++
//line /usr/local/go/src/strings/strings.go:541
			// _ = "end of CoverTab[3676]"
//line /usr/local/go/src/strings/strings.go:541
		}
//line /usr/local/go/src/strings/strings.go:541
		// _ = "end of CoverTab[3674]"
//line /usr/local/go/src/strings/strings.go:541
	}
//line /usr/local/go/src/strings/strings.go:541
	// _ = "end of CoverTab[3665]"
//line /usr/local/go/src/strings/strings.go:541
	_go_fuzz_dep_.CoverTab[3666]++

							if len(s) == 0 {
//line /usr/local/go/src/strings/strings.go:543
		_go_fuzz_dep_.CoverTab[3677]++
								return ""
//line /usr/local/go/src/strings/strings.go:544
		// _ = "end of CoverTab[3677]"
	} else {
//line /usr/local/go/src/strings/strings.go:545
		_go_fuzz_dep_.CoverTab[3678]++
//line /usr/local/go/src/strings/strings.go:545
		// _ = "end of CoverTab[3678]"
//line /usr/local/go/src/strings/strings.go:545
	}
//line /usr/local/go/src/strings/strings.go:545
	// _ = "end of CoverTab[3666]"
//line /usr/local/go/src/strings/strings.go:545
	_go_fuzz_dep_.CoverTab[3667]++

							n := len(s) * count

	// Past a certain chunk size it is counterproductive to use
	// larger chunks as the source of the write, as when the source
	// is too large we are basically just thrashing the CPU D-cache.
	// So if the result length is larger than an empirically-found
	// limit (8KB), we stop growing the source string once the limit
	// is reached and keep reusing the same source string - that
	// should therefore be always resident in the L1 cache - until we
	// have completed the construction of the result.
	// This yields significant speedups (up to +100%) in cases where
	// the result length is large (roughly, over L2 cache size).
	const chunkLimit = 8 * 1024
	chunkMax := n
	if n > chunkLimit {
//line /usr/local/go/src/strings/strings.go:561
		_go_fuzz_dep_.CoverTab[3679]++
								chunkMax = chunkLimit / len(s) * len(s)
								if chunkMax == 0 {
//line /usr/local/go/src/strings/strings.go:563
			_go_fuzz_dep_.CoverTab[3680]++
									chunkMax = len(s)
//line /usr/local/go/src/strings/strings.go:564
			// _ = "end of CoverTab[3680]"
		} else {
//line /usr/local/go/src/strings/strings.go:565
			_go_fuzz_dep_.CoverTab[3681]++
//line /usr/local/go/src/strings/strings.go:565
			// _ = "end of CoverTab[3681]"
//line /usr/local/go/src/strings/strings.go:565
		}
//line /usr/local/go/src/strings/strings.go:565
		// _ = "end of CoverTab[3679]"
	} else {
//line /usr/local/go/src/strings/strings.go:566
		_go_fuzz_dep_.CoverTab[3682]++
//line /usr/local/go/src/strings/strings.go:566
		// _ = "end of CoverTab[3682]"
//line /usr/local/go/src/strings/strings.go:566
	}
//line /usr/local/go/src/strings/strings.go:566
	// _ = "end of CoverTab[3667]"
//line /usr/local/go/src/strings/strings.go:566
	_go_fuzz_dep_.CoverTab[3668]++

							var b Builder
							b.Grow(n)
							b.WriteString(s)
							for b.Len() < n {
//line /usr/local/go/src/strings/strings.go:571
		_go_fuzz_dep_.CoverTab[3683]++
								chunk := n - b.Len()
								if chunk > b.Len() {
//line /usr/local/go/src/strings/strings.go:573
			_go_fuzz_dep_.CoverTab[3686]++
									chunk = b.Len()
//line /usr/local/go/src/strings/strings.go:574
			// _ = "end of CoverTab[3686]"
		} else {
//line /usr/local/go/src/strings/strings.go:575
			_go_fuzz_dep_.CoverTab[3687]++
//line /usr/local/go/src/strings/strings.go:575
			// _ = "end of CoverTab[3687]"
//line /usr/local/go/src/strings/strings.go:575
		}
//line /usr/local/go/src/strings/strings.go:575
		// _ = "end of CoverTab[3683]"
//line /usr/local/go/src/strings/strings.go:575
		_go_fuzz_dep_.CoverTab[3684]++
								if chunk > chunkMax {
//line /usr/local/go/src/strings/strings.go:576
			_go_fuzz_dep_.CoverTab[3688]++
									chunk = chunkMax
//line /usr/local/go/src/strings/strings.go:577
			// _ = "end of CoverTab[3688]"
		} else {
//line /usr/local/go/src/strings/strings.go:578
			_go_fuzz_dep_.CoverTab[3689]++
//line /usr/local/go/src/strings/strings.go:578
			// _ = "end of CoverTab[3689]"
//line /usr/local/go/src/strings/strings.go:578
		}
//line /usr/local/go/src/strings/strings.go:578
		// _ = "end of CoverTab[3684]"
//line /usr/local/go/src/strings/strings.go:578
		_go_fuzz_dep_.CoverTab[3685]++
								b.WriteString(b.String()[:chunk])
//line /usr/local/go/src/strings/strings.go:579
		// _ = "end of CoverTab[3685]"
	}
//line /usr/local/go/src/strings/strings.go:580
	// _ = "end of CoverTab[3668]"
//line /usr/local/go/src/strings/strings.go:580
	_go_fuzz_dep_.CoverTab[3669]++
							return b.String()
//line /usr/local/go/src/strings/strings.go:581
	// _ = "end of CoverTab[3669]"
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
//line /usr/local/go/src/strings/strings.go:585
	_go_fuzz_dep_.CoverTab[3690]++
							isASCII, hasLower := true, false
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:587
		_go_fuzz_dep_.CoverTab[3693]++
								c := s[i]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:589
			_go_fuzz_dep_.CoverTab[3695]++
									isASCII = false
									break
//line /usr/local/go/src/strings/strings.go:591
			// _ = "end of CoverTab[3695]"
		} else {
//line /usr/local/go/src/strings/strings.go:592
			_go_fuzz_dep_.CoverTab[3696]++
//line /usr/local/go/src/strings/strings.go:592
			// _ = "end of CoverTab[3696]"
//line /usr/local/go/src/strings/strings.go:592
		}
//line /usr/local/go/src/strings/strings.go:592
		// _ = "end of CoverTab[3693]"
//line /usr/local/go/src/strings/strings.go:592
		_go_fuzz_dep_.CoverTab[3694]++
								hasLower = hasLower || func() bool {
//line /usr/local/go/src/strings/strings.go:593
			_go_fuzz_dep_.CoverTab[3697]++
//line /usr/local/go/src/strings/strings.go:593
			return ('a' <= c && func() bool {
//line /usr/local/go/src/strings/strings.go:593
				_go_fuzz_dep_.CoverTab[3698]++
//line /usr/local/go/src/strings/strings.go:593
				return c <= 'z'
//line /usr/local/go/src/strings/strings.go:593
				// _ = "end of CoverTab[3698]"
//line /usr/local/go/src/strings/strings.go:593
			}())
//line /usr/local/go/src/strings/strings.go:593
			// _ = "end of CoverTab[3697]"
//line /usr/local/go/src/strings/strings.go:593
		}()
//line /usr/local/go/src/strings/strings.go:593
		// _ = "end of CoverTab[3694]"
	}
//line /usr/local/go/src/strings/strings.go:594
	// _ = "end of CoverTab[3690]"
//line /usr/local/go/src/strings/strings.go:594
	_go_fuzz_dep_.CoverTab[3691]++

							if isASCII {
//line /usr/local/go/src/strings/strings.go:596
		_go_fuzz_dep_.CoverTab[3699]++
								if !hasLower {
//line /usr/local/go/src/strings/strings.go:597
			_go_fuzz_dep_.CoverTab[3703]++
									return s
//line /usr/local/go/src/strings/strings.go:598
			// _ = "end of CoverTab[3703]"
		} else {
//line /usr/local/go/src/strings/strings.go:599
			_go_fuzz_dep_.CoverTab[3704]++
//line /usr/local/go/src/strings/strings.go:599
			// _ = "end of CoverTab[3704]"
//line /usr/local/go/src/strings/strings.go:599
		}
//line /usr/local/go/src/strings/strings.go:599
		// _ = "end of CoverTab[3699]"
//line /usr/local/go/src/strings/strings.go:599
		_go_fuzz_dep_.CoverTab[3700]++
								var (
			b	Builder
			pos	int
		)
		b.Grow(len(s))
		for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:605
			_go_fuzz_dep_.CoverTab[3705]++
									c := s[i]
									if 'a' <= c && func() bool {
//line /usr/local/go/src/strings/strings.go:607
				_go_fuzz_dep_.CoverTab[3706]++
//line /usr/local/go/src/strings/strings.go:607
				return c <= 'z'
//line /usr/local/go/src/strings/strings.go:607
				// _ = "end of CoverTab[3706]"
//line /usr/local/go/src/strings/strings.go:607
			}() {
//line /usr/local/go/src/strings/strings.go:607
				_go_fuzz_dep_.CoverTab[3707]++
										c -= 'a' - 'A'
										if pos < i {
//line /usr/local/go/src/strings/strings.go:609
					_go_fuzz_dep_.CoverTab[3709]++
											b.WriteString(s[pos:i])
//line /usr/local/go/src/strings/strings.go:610
					// _ = "end of CoverTab[3709]"
				} else {
//line /usr/local/go/src/strings/strings.go:611
					_go_fuzz_dep_.CoverTab[3710]++
//line /usr/local/go/src/strings/strings.go:611
					// _ = "end of CoverTab[3710]"
//line /usr/local/go/src/strings/strings.go:611
				}
//line /usr/local/go/src/strings/strings.go:611
				// _ = "end of CoverTab[3707]"
//line /usr/local/go/src/strings/strings.go:611
				_go_fuzz_dep_.CoverTab[3708]++
										b.WriteByte(c)
										pos = i + 1
//line /usr/local/go/src/strings/strings.go:613
				// _ = "end of CoverTab[3708]"
			} else {
//line /usr/local/go/src/strings/strings.go:614
				_go_fuzz_dep_.CoverTab[3711]++
//line /usr/local/go/src/strings/strings.go:614
				// _ = "end of CoverTab[3711]"
//line /usr/local/go/src/strings/strings.go:614
			}
//line /usr/local/go/src/strings/strings.go:614
			// _ = "end of CoverTab[3705]"
		}
//line /usr/local/go/src/strings/strings.go:615
		// _ = "end of CoverTab[3700]"
//line /usr/local/go/src/strings/strings.go:615
		_go_fuzz_dep_.CoverTab[3701]++
								if pos < len(s) {
//line /usr/local/go/src/strings/strings.go:616
			_go_fuzz_dep_.CoverTab[3712]++
									b.WriteString(s[pos:])
//line /usr/local/go/src/strings/strings.go:617
			// _ = "end of CoverTab[3712]"
		} else {
//line /usr/local/go/src/strings/strings.go:618
			_go_fuzz_dep_.CoverTab[3713]++
//line /usr/local/go/src/strings/strings.go:618
			// _ = "end of CoverTab[3713]"
//line /usr/local/go/src/strings/strings.go:618
		}
//line /usr/local/go/src/strings/strings.go:618
		// _ = "end of CoverTab[3701]"
//line /usr/local/go/src/strings/strings.go:618
		_go_fuzz_dep_.CoverTab[3702]++
								return b.String()
//line /usr/local/go/src/strings/strings.go:619
		// _ = "end of CoverTab[3702]"
	} else {
//line /usr/local/go/src/strings/strings.go:620
		_go_fuzz_dep_.CoverTab[3714]++
//line /usr/local/go/src/strings/strings.go:620
		// _ = "end of CoverTab[3714]"
//line /usr/local/go/src/strings/strings.go:620
	}
//line /usr/local/go/src/strings/strings.go:620
	// _ = "end of CoverTab[3691]"
//line /usr/local/go/src/strings/strings.go:620
	_go_fuzz_dep_.CoverTab[3692]++
							return Map(unicode.ToUpper, s)
//line /usr/local/go/src/strings/strings.go:621
	// _ = "end of CoverTab[3692]"
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
//line /usr/local/go/src/strings/strings.go:625
	_go_fuzz_dep_.CoverTab[3715]++
							isASCII, hasUpper := true, false
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:627
		_go_fuzz_dep_.CoverTab[3718]++
								c := s[i]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:629
			_go_fuzz_dep_.CoverTab[3720]++
									isASCII = false
									break
//line /usr/local/go/src/strings/strings.go:631
			// _ = "end of CoverTab[3720]"
		} else {
//line /usr/local/go/src/strings/strings.go:632
			_go_fuzz_dep_.CoverTab[3721]++
//line /usr/local/go/src/strings/strings.go:632
			// _ = "end of CoverTab[3721]"
//line /usr/local/go/src/strings/strings.go:632
		}
//line /usr/local/go/src/strings/strings.go:632
		// _ = "end of CoverTab[3718]"
//line /usr/local/go/src/strings/strings.go:632
		_go_fuzz_dep_.CoverTab[3719]++
								hasUpper = hasUpper || func() bool {
//line /usr/local/go/src/strings/strings.go:633
			_go_fuzz_dep_.CoverTab[3722]++
//line /usr/local/go/src/strings/strings.go:633
			return ('A' <= c && func() bool {
//line /usr/local/go/src/strings/strings.go:633
				_go_fuzz_dep_.CoverTab[3723]++
//line /usr/local/go/src/strings/strings.go:633
				return c <= 'Z'
//line /usr/local/go/src/strings/strings.go:633
				// _ = "end of CoverTab[3723]"
//line /usr/local/go/src/strings/strings.go:633
			}())
//line /usr/local/go/src/strings/strings.go:633
			// _ = "end of CoverTab[3722]"
//line /usr/local/go/src/strings/strings.go:633
		}()
//line /usr/local/go/src/strings/strings.go:633
		// _ = "end of CoverTab[3719]"
	}
//line /usr/local/go/src/strings/strings.go:634
	// _ = "end of CoverTab[3715]"
//line /usr/local/go/src/strings/strings.go:634
	_go_fuzz_dep_.CoverTab[3716]++

							if isASCII {
//line /usr/local/go/src/strings/strings.go:636
		_go_fuzz_dep_.CoverTab[3724]++
								if !hasUpper {
//line /usr/local/go/src/strings/strings.go:637
			_go_fuzz_dep_.CoverTab[3728]++
									return s
//line /usr/local/go/src/strings/strings.go:638
			// _ = "end of CoverTab[3728]"
		} else {
//line /usr/local/go/src/strings/strings.go:639
			_go_fuzz_dep_.CoverTab[3729]++
//line /usr/local/go/src/strings/strings.go:639
			// _ = "end of CoverTab[3729]"
//line /usr/local/go/src/strings/strings.go:639
		}
//line /usr/local/go/src/strings/strings.go:639
		// _ = "end of CoverTab[3724]"
//line /usr/local/go/src/strings/strings.go:639
		_go_fuzz_dep_.CoverTab[3725]++
								var (
			b	Builder
			pos	int
		)
		b.Grow(len(s))
		for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:645
			_go_fuzz_dep_.CoverTab[3730]++
									c := s[i]
									if 'A' <= c && func() bool {
//line /usr/local/go/src/strings/strings.go:647
				_go_fuzz_dep_.CoverTab[3731]++
//line /usr/local/go/src/strings/strings.go:647
				return c <= 'Z'
//line /usr/local/go/src/strings/strings.go:647
				// _ = "end of CoverTab[3731]"
//line /usr/local/go/src/strings/strings.go:647
			}() {
//line /usr/local/go/src/strings/strings.go:647
				_go_fuzz_dep_.CoverTab[3732]++
										c += 'a' - 'A'
										if pos < i {
//line /usr/local/go/src/strings/strings.go:649
					_go_fuzz_dep_.CoverTab[3734]++
											b.WriteString(s[pos:i])
//line /usr/local/go/src/strings/strings.go:650
					// _ = "end of CoverTab[3734]"
				} else {
//line /usr/local/go/src/strings/strings.go:651
					_go_fuzz_dep_.CoverTab[3735]++
//line /usr/local/go/src/strings/strings.go:651
					// _ = "end of CoverTab[3735]"
//line /usr/local/go/src/strings/strings.go:651
				}
//line /usr/local/go/src/strings/strings.go:651
				// _ = "end of CoverTab[3732]"
//line /usr/local/go/src/strings/strings.go:651
				_go_fuzz_dep_.CoverTab[3733]++
										b.WriteByte(c)
										pos = i + 1
//line /usr/local/go/src/strings/strings.go:653
				// _ = "end of CoverTab[3733]"
			} else {
//line /usr/local/go/src/strings/strings.go:654
				_go_fuzz_dep_.CoverTab[3736]++
//line /usr/local/go/src/strings/strings.go:654
				// _ = "end of CoverTab[3736]"
//line /usr/local/go/src/strings/strings.go:654
			}
//line /usr/local/go/src/strings/strings.go:654
			// _ = "end of CoverTab[3730]"
		}
//line /usr/local/go/src/strings/strings.go:655
		// _ = "end of CoverTab[3725]"
//line /usr/local/go/src/strings/strings.go:655
		_go_fuzz_dep_.CoverTab[3726]++
								if pos < len(s) {
//line /usr/local/go/src/strings/strings.go:656
			_go_fuzz_dep_.CoverTab[3737]++
									b.WriteString(s[pos:])
//line /usr/local/go/src/strings/strings.go:657
			// _ = "end of CoverTab[3737]"
		} else {
//line /usr/local/go/src/strings/strings.go:658
			_go_fuzz_dep_.CoverTab[3738]++
//line /usr/local/go/src/strings/strings.go:658
			// _ = "end of CoverTab[3738]"
//line /usr/local/go/src/strings/strings.go:658
		}
//line /usr/local/go/src/strings/strings.go:658
		// _ = "end of CoverTab[3726]"
//line /usr/local/go/src/strings/strings.go:658
		_go_fuzz_dep_.CoverTab[3727]++
								return b.String()
//line /usr/local/go/src/strings/strings.go:659
		// _ = "end of CoverTab[3727]"
	} else {
//line /usr/local/go/src/strings/strings.go:660
		_go_fuzz_dep_.CoverTab[3739]++
//line /usr/local/go/src/strings/strings.go:660
		// _ = "end of CoverTab[3739]"
//line /usr/local/go/src/strings/strings.go:660
	}
//line /usr/local/go/src/strings/strings.go:660
	// _ = "end of CoverTab[3716]"
//line /usr/local/go/src/strings/strings.go:660
	_go_fuzz_dep_.CoverTab[3717]++
							return Map(unicode.ToLower, s)
//line /usr/local/go/src/strings/strings.go:661
	// _ = "end of CoverTab[3717]"
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to
//line /usr/local/go/src/strings/strings.go:664
// their Unicode title case.
//line /usr/local/go/src/strings/strings.go:666
func ToTitle(s string) string {
//line /usr/local/go/src/strings/strings.go:666
	_go_fuzz_dep_.CoverTab[3740]++
//line /usr/local/go/src/strings/strings.go:666
	return Map(unicode.ToTitle, s)
//line /usr/local/go/src/strings/strings.go:666
	// _ = "end of CoverTab[3740]"
//line /usr/local/go/src/strings/strings.go:666
}

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
//line /usr/local/go/src/strings/strings.go:668
// upper case using the case mapping specified by c.
//line /usr/local/go/src/strings/strings.go:670
func ToUpperSpecial(c unicode.SpecialCase, s string) string {
//line /usr/local/go/src/strings/strings.go:670
	_go_fuzz_dep_.CoverTab[3741]++
							return Map(c.ToUpper, s)
//line /usr/local/go/src/strings/strings.go:671
	// _ = "end of CoverTab[3741]"
}

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
//line /usr/local/go/src/strings/strings.go:674
// lower case using the case mapping specified by c.
//line /usr/local/go/src/strings/strings.go:676
func ToLowerSpecial(c unicode.SpecialCase, s string) string {
//line /usr/local/go/src/strings/strings.go:676
	_go_fuzz_dep_.CoverTab[3742]++
							return Map(c.ToLower, s)
//line /usr/local/go/src/strings/strings.go:677
	// _ = "end of CoverTab[3742]"
}

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
//line /usr/local/go/src/strings/strings.go:680
// Unicode title case, giving priority to the special casing rules.
//line /usr/local/go/src/strings/strings.go:682
func ToTitleSpecial(c unicode.SpecialCase, s string) string {
//line /usr/local/go/src/strings/strings.go:682
	_go_fuzz_dep_.CoverTab[3743]++
							return Map(c.ToTitle, s)
//line /usr/local/go/src/strings/strings.go:683
	// _ = "end of CoverTab[3743]"
}

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
//line /usr/local/go/src/strings/strings.go:686
// replaced by the replacement string, which may be empty.
//line /usr/local/go/src/strings/strings.go:688
func ToValidUTF8(s, replacement string) string {
//line /usr/local/go/src/strings/strings.go:688
	_go_fuzz_dep_.CoverTab[3744]++
							var b Builder

							for i, c := range s {
//line /usr/local/go/src/strings/strings.go:691
		_go_fuzz_dep_.CoverTab[3748]++
								if c != utf8.RuneError {
//line /usr/local/go/src/strings/strings.go:692
			_go_fuzz_dep_.CoverTab[3750]++
									continue
//line /usr/local/go/src/strings/strings.go:693
			// _ = "end of CoverTab[3750]"
		} else {
//line /usr/local/go/src/strings/strings.go:694
			_go_fuzz_dep_.CoverTab[3751]++
//line /usr/local/go/src/strings/strings.go:694
			// _ = "end of CoverTab[3751]"
//line /usr/local/go/src/strings/strings.go:694
		}
//line /usr/local/go/src/strings/strings.go:694
		// _ = "end of CoverTab[3748]"
//line /usr/local/go/src/strings/strings.go:694
		_go_fuzz_dep_.CoverTab[3749]++

								_, wid := utf8.DecodeRuneInString(s[i:])
								if wid == 1 {
//line /usr/local/go/src/strings/strings.go:697
			_go_fuzz_dep_.CoverTab[3752]++
									b.Grow(len(s) + len(replacement))
									b.WriteString(s[:i])
									s = s[i:]
									break
//line /usr/local/go/src/strings/strings.go:701
			// _ = "end of CoverTab[3752]"
		} else {
//line /usr/local/go/src/strings/strings.go:702
			_go_fuzz_dep_.CoverTab[3753]++
//line /usr/local/go/src/strings/strings.go:702
			// _ = "end of CoverTab[3753]"
//line /usr/local/go/src/strings/strings.go:702
		}
//line /usr/local/go/src/strings/strings.go:702
		// _ = "end of CoverTab[3749]"
	}
//line /usr/local/go/src/strings/strings.go:703
	// _ = "end of CoverTab[3744]"
//line /usr/local/go/src/strings/strings.go:703
	_go_fuzz_dep_.CoverTab[3745]++

//line /usr/local/go/src/strings/strings.go:706
	if b.Cap() == 0 {
//line /usr/local/go/src/strings/strings.go:706
		_go_fuzz_dep_.CoverTab[3754]++
								return s
//line /usr/local/go/src/strings/strings.go:707
		// _ = "end of CoverTab[3754]"
	} else {
//line /usr/local/go/src/strings/strings.go:708
		_go_fuzz_dep_.CoverTab[3755]++
//line /usr/local/go/src/strings/strings.go:708
		// _ = "end of CoverTab[3755]"
//line /usr/local/go/src/strings/strings.go:708
	}
//line /usr/local/go/src/strings/strings.go:708
	// _ = "end of CoverTab[3745]"
//line /usr/local/go/src/strings/strings.go:708
	_go_fuzz_dep_.CoverTab[3746]++

							invalid := false
							for i := 0; i < len(s); {
//line /usr/local/go/src/strings/strings.go:711
		_go_fuzz_dep_.CoverTab[3756]++
								c := s[i]
								if c < utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:713
			_go_fuzz_dep_.CoverTab[3759]++
									i++
									invalid = false
									b.WriteByte(c)
									continue
//line /usr/local/go/src/strings/strings.go:717
			// _ = "end of CoverTab[3759]"
		} else {
//line /usr/local/go/src/strings/strings.go:718
			_go_fuzz_dep_.CoverTab[3760]++
//line /usr/local/go/src/strings/strings.go:718
			// _ = "end of CoverTab[3760]"
//line /usr/local/go/src/strings/strings.go:718
		}
//line /usr/local/go/src/strings/strings.go:718
		// _ = "end of CoverTab[3756]"
//line /usr/local/go/src/strings/strings.go:718
		_go_fuzz_dep_.CoverTab[3757]++
								_, wid := utf8.DecodeRuneInString(s[i:])
								if wid == 1 {
//line /usr/local/go/src/strings/strings.go:720
			_go_fuzz_dep_.CoverTab[3761]++
									i++
									if !invalid {
//line /usr/local/go/src/strings/strings.go:722
				_go_fuzz_dep_.CoverTab[3763]++
										invalid = true
										b.WriteString(replacement)
//line /usr/local/go/src/strings/strings.go:724
				// _ = "end of CoverTab[3763]"
			} else {
//line /usr/local/go/src/strings/strings.go:725
				_go_fuzz_dep_.CoverTab[3764]++
//line /usr/local/go/src/strings/strings.go:725
				// _ = "end of CoverTab[3764]"
//line /usr/local/go/src/strings/strings.go:725
			}
//line /usr/local/go/src/strings/strings.go:725
			// _ = "end of CoverTab[3761]"
//line /usr/local/go/src/strings/strings.go:725
			_go_fuzz_dep_.CoverTab[3762]++
									continue
//line /usr/local/go/src/strings/strings.go:726
			// _ = "end of CoverTab[3762]"
		} else {
//line /usr/local/go/src/strings/strings.go:727
			_go_fuzz_dep_.CoverTab[3765]++
//line /usr/local/go/src/strings/strings.go:727
			// _ = "end of CoverTab[3765]"
//line /usr/local/go/src/strings/strings.go:727
		}
//line /usr/local/go/src/strings/strings.go:727
		// _ = "end of CoverTab[3757]"
//line /usr/local/go/src/strings/strings.go:727
		_go_fuzz_dep_.CoverTab[3758]++
								invalid = false
								b.WriteString(s[i : i+wid])
								i += wid
//line /usr/local/go/src/strings/strings.go:730
		// _ = "end of CoverTab[3758]"
	}
//line /usr/local/go/src/strings/strings.go:731
	// _ = "end of CoverTab[3746]"
//line /usr/local/go/src/strings/strings.go:731
	_go_fuzz_dep_.CoverTab[3747]++

							return b.String()
//line /usr/local/go/src/strings/strings.go:733
	// _ = "end of CoverTab[3747]"
}

// isSeparator reports whether the rune could mark a word boundary.
//line /usr/local/go/src/strings/strings.go:736
// TODO: update when package unicode captures more of the properties.
//line /usr/local/go/src/strings/strings.go:738
func isSeparator(r rune) bool {
//line /usr/local/go/src/strings/strings.go:738
	_go_fuzz_dep_.CoverTab[3766]++

							if r <= 0x7F {
//line /usr/local/go/src/strings/strings.go:740
		_go_fuzz_dep_.CoverTab[3769]++
								switch {
		case '0' <= r && func() bool {
//line /usr/local/go/src/strings/strings.go:742
			_go_fuzz_dep_.CoverTab[3776]++
//line /usr/local/go/src/strings/strings.go:742
			return r <= '9'
//line /usr/local/go/src/strings/strings.go:742
			// _ = "end of CoverTab[3776]"
//line /usr/local/go/src/strings/strings.go:742
		}():
//line /usr/local/go/src/strings/strings.go:742
			_go_fuzz_dep_.CoverTab[3771]++
									return false
//line /usr/local/go/src/strings/strings.go:743
			// _ = "end of CoverTab[3771]"
		case 'a' <= r && func() bool {
//line /usr/local/go/src/strings/strings.go:744
			_go_fuzz_dep_.CoverTab[3777]++
//line /usr/local/go/src/strings/strings.go:744
			return r <= 'z'
//line /usr/local/go/src/strings/strings.go:744
			// _ = "end of CoverTab[3777]"
//line /usr/local/go/src/strings/strings.go:744
		}():
//line /usr/local/go/src/strings/strings.go:744
			_go_fuzz_dep_.CoverTab[3772]++
									return false
//line /usr/local/go/src/strings/strings.go:745
			// _ = "end of CoverTab[3772]"
		case 'A' <= r && func() bool {
//line /usr/local/go/src/strings/strings.go:746
			_go_fuzz_dep_.CoverTab[3778]++
//line /usr/local/go/src/strings/strings.go:746
			return r <= 'Z'
//line /usr/local/go/src/strings/strings.go:746
			// _ = "end of CoverTab[3778]"
//line /usr/local/go/src/strings/strings.go:746
		}():
//line /usr/local/go/src/strings/strings.go:746
			_go_fuzz_dep_.CoverTab[3773]++
									return false
//line /usr/local/go/src/strings/strings.go:747
			// _ = "end of CoverTab[3773]"
		case r == '_':
//line /usr/local/go/src/strings/strings.go:748
			_go_fuzz_dep_.CoverTab[3774]++
									return false
//line /usr/local/go/src/strings/strings.go:749
			// _ = "end of CoverTab[3774]"
//line /usr/local/go/src/strings/strings.go:749
		default:
//line /usr/local/go/src/strings/strings.go:749
			_go_fuzz_dep_.CoverTab[3775]++
//line /usr/local/go/src/strings/strings.go:749
			// _ = "end of CoverTab[3775]"
		}
//line /usr/local/go/src/strings/strings.go:750
		// _ = "end of CoverTab[3769]"
//line /usr/local/go/src/strings/strings.go:750
		_go_fuzz_dep_.CoverTab[3770]++
								return true
//line /usr/local/go/src/strings/strings.go:751
		// _ = "end of CoverTab[3770]"
	} else {
//line /usr/local/go/src/strings/strings.go:752
		_go_fuzz_dep_.CoverTab[3779]++
//line /usr/local/go/src/strings/strings.go:752
		// _ = "end of CoverTab[3779]"
//line /usr/local/go/src/strings/strings.go:752
	}
//line /usr/local/go/src/strings/strings.go:752
	// _ = "end of CoverTab[3766]"
//line /usr/local/go/src/strings/strings.go:752
	_go_fuzz_dep_.CoverTab[3767]++

							if unicode.IsLetter(r) || func() bool {
//line /usr/local/go/src/strings/strings.go:754
		_go_fuzz_dep_.CoverTab[3780]++
//line /usr/local/go/src/strings/strings.go:754
		return unicode.IsDigit(r)
//line /usr/local/go/src/strings/strings.go:754
		// _ = "end of CoverTab[3780]"
//line /usr/local/go/src/strings/strings.go:754
	}() {
//line /usr/local/go/src/strings/strings.go:754
		_go_fuzz_dep_.CoverTab[3781]++
								return false
//line /usr/local/go/src/strings/strings.go:755
		// _ = "end of CoverTab[3781]"
	} else {
//line /usr/local/go/src/strings/strings.go:756
		_go_fuzz_dep_.CoverTab[3782]++
//line /usr/local/go/src/strings/strings.go:756
		// _ = "end of CoverTab[3782]"
//line /usr/local/go/src/strings/strings.go:756
	}
//line /usr/local/go/src/strings/strings.go:756
	// _ = "end of CoverTab[3767]"
//line /usr/local/go/src/strings/strings.go:756
	_go_fuzz_dep_.CoverTab[3768]++

							return unicode.IsSpace(r)
//line /usr/local/go/src/strings/strings.go:758
	// _ = "end of CoverTab[3768]"
}

// Title returns a copy of the string s with all Unicode letters that begin words
//line /usr/local/go/src/strings/strings.go:761
// mapped to their Unicode title case.
//line /usr/local/go/src/strings/strings.go:761
//
//line /usr/local/go/src/strings/strings.go:761
// Deprecated: The rule Title uses for word boundaries does not handle Unicode
//line /usr/local/go/src/strings/strings.go:761
// punctuation properly. Use golang.org/x/text/cases instead.
//line /usr/local/go/src/strings/strings.go:766
func Title(s string) string {
//line /usr/local/go/src/strings/strings.go:766
	_go_fuzz_dep_.CoverTab[3783]++

//line /usr/local/go/src/strings/strings.go:770
	prev := ' '
	return Map(
		func(r rune) rune {
//line /usr/local/go/src/strings/strings.go:772
			_go_fuzz_dep_.CoverTab[3784]++
									if isSeparator(prev) {
//line /usr/local/go/src/strings/strings.go:773
				_go_fuzz_dep_.CoverTab[3786]++
										prev = r
										return unicode.ToTitle(r)
//line /usr/local/go/src/strings/strings.go:775
				// _ = "end of CoverTab[3786]"
			} else {
//line /usr/local/go/src/strings/strings.go:776
				_go_fuzz_dep_.CoverTab[3787]++
//line /usr/local/go/src/strings/strings.go:776
				// _ = "end of CoverTab[3787]"
//line /usr/local/go/src/strings/strings.go:776
			}
//line /usr/local/go/src/strings/strings.go:776
			// _ = "end of CoverTab[3784]"
//line /usr/local/go/src/strings/strings.go:776
			_go_fuzz_dep_.CoverTab[3785]++
									prev = r
									return r
//line /usr/local/go/src/strings/strings.go:778
			// _ = "end of CoverTab[3785]"
		},
		s)
//line /usr/local/go/src/strings/strings.go:780
	// _ = "end of CoverTab[3783]"
}

// TrimLeftFunc returns a slice of the string s with all leading
//line /usr/local/go/src/strings/strings.go:783
// Unicode code points c satisfying f(c) removed.
//line /usr/local/go/src/strings/strings.go:785
func TrimLeftFunc(s string, f func(rune) bool) string {
//line /usr/local/go/src/strings/strings.go:785
	_go_fuzz_dep_.CoverTab[3788]++
							i := indexFunc(s, f, false)
							if i == -1 {
//line /usr/local/go/src/strings/strings.go:787
		_go_fuzz_dep_.CoverTab[3790]++
								return ""
//line /usr/local/go/src/strings/strings.go:788
		// _ = "end of CoverTab[3790]"
	} else {
//line /usr/local/go/src/strings/strings.go:789
		_go_fuzz_dep_.CoverTab[3791]++
//line /usr/local/go/src/strings/strings.go:789
		// _ = "end of CoverTab[3791]"
//line /usr/local/go/src/strings/strings.go:789
	}
//line /usr/local/go/src/strings/strings.go:789
	// _ = "end of CoverTab[3788]"
//line /usr/local/go/src/strings/strings.go:789
	_go_fuzz_dep_.CoverTab[3789]++
							return s[i:]
//line /usr/local/go/src/strings/strings.go:790
	// _ = "end of CoverTab[3789]"
}

// TrimRightFunc returns a slice of the string s with all trailing
//line /usr/local/go/src/strings/strings.go:793
// Unicode code points c satisfying f(c) removed.
//line /usr/local/go/src/strings/strings.go:795
func TrimRightFunc(s string, f func(rune) bool) string {
//line /usr/local/go/src/strings/strings.go:795
	_go_fuzz_dep_.CoverTab[3792]++
							i := lastIndexFunc(s, f, false)
							if i >= 0 && func() bool {
//line /usr/local/go/src/strings/strings.go:797
		_go_fuzz_dep_.CoverTab[3794]++
//line /usr/local/go/src/strings/strings.go:797
		return s[i] >= utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:797
		// _ = "end of CoverTab[3794]"
//line /usr/local/go/src/strings/strings.go:797
	}() {
//line /usr/local/go/src/strings/strings.go:797
		_go_fuzz_dep_.CoverTab[3795]++
								_, wid := utf8.DecodeRuneInString(s[i:])
								i += wid
//line /usr/local/go/src/strings/strings.go:799
		// _ = "end of CoverTab[3795]"
	} else {
//line /usr/local/go/src/strings/strings.go:800
		_go_fuzz_dep_.CoverTab[3796]++
								i++
//line /usr/local/go/src/strings/strings.go:801
		// _ = "end of CoverTab[3796]"
	}
//line /usr/local/go/src/strings/strings.go:802
	// _ = "end of CoverTab[3792]"
//line /usr/local/go/src/strings/strings.go:802
	_go_fuzz_dep_.CoverTab[3793]++
							return s[0:i]
//line /usr/local/go/src/strings/strings.go:803
	// _ = "end of CoverTab[3793]"
}

// TrimFunc returns a slice of the string s with all leading
//line /usr/local/go/src/strings/strings.go:806
// and trailing Unicode code points c satisfying f(c) removed.
//line /usr/local/go/src/strings/strings.go:808
func TrimFunc(s string, f func(rune) bool) string {
//line /usr/local/go/src/strings/strings.go:808
	_go_fuzz_dep_.CoverTab[3797]++
							return TrimRightFunc(TrimLeftFunc(s, f), f)
//line /usr/local/go/src/strings/strings.go:809
	// _ = "end of CoverTab[3797]"
}

// IndexFunc returns the index into s of the first Unicode
//line /usr/local/go/src/strings/strings.go:812
// code point satisfying f(c), or -1 if none do.
//line /usr/local/go/src/strings/strings.go:814
func IndexFunc(s string, f func(rune) bool) int {
//line /usr/local/go/src/strings/strings.go:814
	_go_fuzz_dep_.CoverTab[3798]++
							return indexFunc(s, f, true)
//line /usr/local/go/src/strings/strings.go:815
	// _ = "end of CoverTab[3798]"
}

// LastIndexFunc returns the index into s of the last
//line /usr/local/go/src/strings/strings.go:818
// Unicode code point satisfying f(c), or -1 if none do.
//line /usr/local/go/src/strings/strings.go:820
func LastIndexFunc(s string, f func(rune) bool) int {
//line /usr/local/go/src/strings/strings.go:820
	_go_fuzz_dep_.CoverTab[3799]++
							return lastIndexFunc(s, f, true)
//line /usr/local/go/src/strings/strings.go:821
	// _ = "end of CoverTab[3799]"
}

// indexFunc is the same as IndexFunc except that if
//line /usr/local/go/src/strings/strings.go:824
// truth==false, the sense of the predicate function is
//line /usr/local/go/src/strings/strings.go:824
// inverted.
//line /usr/local/go/src/strings/strings.go:827
func indexFunc(s string, f func(rune) bool, truth bool) int {
//line /usr/local/go/src/strings/strings.go:827
	_go_fuzz_dep_.CoverTab[3800]++
							for i, r := range s {
//line /usr/local/go/src/strings/strings.go:828
		_go_fuzz_dep_.CoverTab[3802]++
								if f(r) == truth {
//line /usr/local/go/src/strings/strings.go:829
			_go_fuzz_dep_.CoverTab[3803]++
									return i
//line /usr/local/go/src/strings/strings.go:830
			// _ = "end of CoverTab[3803]"
		} else {
//line /usr/local/go/src/strings/strings.go:831
			_go_fuzz_dep_.CoverTab[3804]++
//line /usr/local/go/src/strings/strings.go:831
			// _ = "end of CoverTab[3804]"
//line /usr/local/go/src/strings/strings.go:831
		}
//line /usr/local/go/src/strings/strings.go:831
		// _ = "end of CoverTab[3802]"
	}
//line /usr/local/go/src/strings/strings.go:832
	// _ = "end of CoverTab[3800]"
//line /usr/local/go/src/strings/strings.go:832
	_go_fuzz_dep_.CoverTab[3801]++
							return -1
//line /usr/local/go/src/strings/strings.go:833
	// _ = "end of CoverTab[3801]"
}

// lastIndexFunc is the same as LastIndexFunc except that if
//line /usr/local/go/src/strings/strings.go:836
// truth==false, the sense of the predicate function is
//line /usr/local/go/src/strings/strings.go:836
// inverted.
//line /usr/local/go/src/strings/strings.go:839
func lastIndexFunc(s string, f func(rune) bool, truth bool) int {
//line /usr/local/go/src/strings/strings.go:839
	_go_fuzz_dep_.CoverTab[3805]++
							for i := len(s); i > 0; {
//line /usr/local/go/src/strings/strings.go:840
		_go_fuzz_dep_.CoverTab[3807]++
								r, size := utf8.DecodeLastRuneInString(s[0:i])
								i -= size
								if f(r) == truth {
//line /usr/local/go/src/strings/strings.go:843
			_go_fuzz_dep_.CoverTab[3808]++
									return i
//line /usr/local/go/src/strings/strings.go:844
			// _ = "end of CoverTab[3808]"
		} else {
//line /usr/local/go/src/strings/strings.go:845
			_go_fuzz_dep_.CoverTab[3809]++
//line /usr/local/go/src/strings/strings.go:845
			// _ = "end of CoverTab[3809]"
//line /usr/local/go/src/strings/strings.go:845
		}
//line /usr/local/go/src/strings/strings.go:845
		// _ = "end of CoverTab[3807]"
	}
//line /usr/local/go/src/strings/strings.go:846
	// _ = "end of CoverTab[3805]"
//line /usr/local/go/src/strings/strings.go:846
	_go_fuzz_dep_.CoverTab[3806]++
							return -1
//line /usr/local/go/src/strings/strings.go:847
	// _ = "end of CoverTab[3806]"
}

// asciiSet is a 32-byte value, where each bit represents the presence of a
//line /usr/local/go/src/strings/strings.go:850
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
//line /usr/local/go/src/strings/strings.go:850
// starting with the least-significant bit of the lowest word to the
//line /usr/local/go/src/strings/strings.go:850
// most-significant bit of the highest word, map to the full range of all
//line /usr/local/go/src/strings/strings.go:850
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
//line /usr/local/go/src/strings/strings.go:850
// ensuring that any non-ASCII character will be reported as not in the set.
//line /usr/local/go/src/strings/strings.go:850
// This allocates a total of 32 bytes even though the upper half
//line /usr/local/go/src/strings/strings.go:850
// is unused to avoid bounds checks in asciiSet.contains.
//line /usr/local/go/src/strings/strings.go:858
type asciiSet [8]uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
//line /usr/local/go/src/strings/strings.go:860
// characters in chars are ASCII.
//line /usr/local/go/src/strings/strings.go:862
func makeASCIISet(chars string) (as asciiSet, ok bool) {
//line /usr/local/go/src/strings/strings.go:862
	_go_fuzz_dep_.CoverTab[3810]++
							for i := 0; i < len(chars); i++ {
//line /usr/local/go/src/strings/strings.go:863
		_go_fuzz_dep_.CoverTab[3812]++
								c := chars[i]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:865
			_go_fuzz_dep_.CoverTab[3814]++
									return as, false
//line /usr/local/go/src/strings/strings.go:866
			// _ = "end of CoverTab[3814]"
		} else {
//line /usr/local/go/src/strings/strings.go:867
			_go_fuzz_dep_.CoverTab[3815]++
//line /usr/local/go/src/strings/strings.go:867
			// _ = "end of CoverTab[3815]"
//line /usr/local/go/src/strings/strings.go:867
		}
//line /usr/local/go/src/strings/strings.go:867
		// _ = "end of CoverTab[3812]"
//line /usr/local/go/src/strings/strings.go:867
		_go_fuzz_dep_.CoverTab[3813]++
								as[c/32] |= 1 << (c % 32)
//line /usr/local/go/src/strings/strings.go:868
		// _ = "end of CoverTab[3813]"
	}
//line /usr/local/go/src/strings/strings.go:869
	// _ = "end of CoverTab[3810]"
//line /usr/local/go/src/strings/strings.go:869
	_go_fuzz_dep_.CoverTab[3811]++
							return as, true
//line /usr/local/go/src/strings/strings.go:870
	// _ = "end of CoverTab[3811]"
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c byte) bool {
//line /usr/local/go/src/strings/strings.go:874
	_go_fuzz_dep_.CoverTab[3816]++
							return (as[c/32] & (1 << (c % 32))) != 0
//line /usr/local/go/src/strings/strings.go:875
	// _ = "end of CoverTab[3816]"
}

// Trim returns a slice of the string s with all leading and
//line /usr/local/go/src/strings/strings.go:878
// trailing Unicode code points contained in cutset removed.
//line /usr/local/go/src/strings/strings.go:880
func Trim(s, cutset string) string {
//line /usr/local/go/src/strings/strings.go:880
	_go_fuzz_dep_.CoverTab[3817]++
							if s == "" || func() bool {
//line /usr/local/go/src/strings/strings.go:881
		_go_fuzz_dep_.CoverTab[3821]++
//line /usr/local/go/src/strings/strings.go:881
		return cutset == ""
//line /usr/local/go/src/strings/strings.go:881
		// _ = "end of CoverTab[3821]"
//line /usr/local/go/src/strings/strings.go:881
	}() {
//line /usr/local/go/src/strings/strings.go:881
		_go_fuzz_dep_.CoverTab[3822]++
								return s
//line /usr/local/go/src/strings/strings.go:882
		// _ = "end of CoverTab[3822]"
	} else {
//line /usr/local/go/src/strings/strings.go:883
		_go_fuzz_dep_.CoverTab[3823]++
//line /usr/local/go/src/strings/strings.go:883
		// _ = "end of CoverTab[3823]"
//line /usr/local/go/src/strings/strings.go:883
	}
//line /usr/local/go/src/strings/strings.go:883
	// _ = "end of CoverTab[3817]"
//line /usr/local/go/src/strings/strings.go:883
	_go_fuzz_dep_.CoverTab[3818]++
							if len(cutset) == 1 && func() bool {
//line /usr/local/go/src/strings/strings.go:884
		_go_fuzz_dep_.CoverTab[3824]++
//line /usr/local/go/src/strings/strings.go:884
		return cutset[0] < utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:884
		// _ = "end of CoverTab[3824]"
//line /usr/local/go/src/strings/strings.go:884
	}() {
//line /usr/local/go/src/strings/strings.go:884
		_go_fuzz_dep_.CoverTab[3825]++
								return trimLeftByte(trimRightByte(s, cutset[0]), cutset[0])
//line /usr/local/go/src/strings/strings.go:885
		// _ = "end of CoverTab[3825]"
	} else {
//line /usr/local/go/src/strings/strings.go:886
		_go_fuzz_dep_.CoverTab[3826]++
//line /usr/local/go/src/strings/strings.go:886
		// _ = "end of CoverTab[3826]"
//line /usr/local/go/src/strings/strings.go:886
	}
//line /usr/local/go/src/strings/strings.go:886
	// _ = "end of CoverTab[3818]"
//line /usr/local/go/src/strings/strings.go:886
	_go_fuzz_dep_.CoverTab[3819]++
							if as, ok := makeASCIISet(cutset); ok {
//line /usr/local/go/src/strings/strings.go:887
		_go_fuzz_dep_.CoverTab[3827]++
								return trimLeftASCII(trimRightASCII(s, &as), &as)
//line /usr/local/go/src/strings/strings.go:888
		// _ = "end of CoverTab[3827]"
	} else {
//line /usr/local/go/src/strings/strings.go:889
		_go_fuzz_dep_.CoverTab[3828]++
//line /usr/local/go/src/strings/strings.go:889
		// _ = "end of CoverTab[3828]"
//line /usr/local/go/src/strings/strings.go:889
	}
//line /usr/local/go/src/strings/strings.go:889
	// _ = "end of CoverTab[3819]"
//line /usr/local/go/src/strings/strings.go:889
	_go_fuzz_dep_.CoverTab[3820]++
							return trimLeftUnicode(trimRightUnicode(s, cutset), cutset)
//line /usr/local/go/src/strings/strings.go:890
	// _ = "end of CoverTab[3820]"
}

// TrimLeft returns a slice of the string s with all leading
//line /usr/local/go/src/strings/strings.go:893
// Unicode code points contained in cutset removed.
//line /usr/local/go/src/strings/strings.go:893
//
//line /usr/local/go/src/strings/strings.go:893
// To remove a prefix, use TrimPrefix instead.
//line /usr/local/go/src/strings/strings.go:897
func TrimLeft(s, cutset string) string {
//line /usr/local/go/src/strings/strings.go:897
	_go_fuzz_dep_.CoverTab[3829]++
							if s == "" || func() bool {
//line /usr/local/go/src/strings/strings.go:898
		_go_fuzz_dep_.CoverTab[3833]++
//line /usr/local/go/src/strings/strings.go:898
		return cutset == ""
//line /usr/local/go/src/strings/strings.go:898
		// _ = "end of CoverTab[3833]"
//line /usr/local/go/src/strings/strings.go:898
	}() {
//line /usr/local/go/src/strings/strings.go:898
		_go_fuzz_dep_.CoverTab[3834]++
								return s
//line /usr/local/go/src/strings/strings.go:899
		// _ = "end of CoverTab[3834]"
	} else {
//line /usr/local/go/src/strings/strings.go:900
		_go_fuzz_dep_.CoverTab[3835]++
//line /usr/local/go/src/strings/strings.go:900
		// _ = "end of CoverTab[3835]"
//line /usr/local/go/src/strings/strings.go:900
	}
//line /usr/local/go/src/strings/strings.go:900
	// _ = "end of CoverTab[3829]"
//line /usr/local/go/src/strings/strings.go:900
	_go_fuzz_dep_.CoverTab[3830]++
							if len(cutset) == 1 && func() bool {
//line /usr/local/go/src/strings/strings.go:901
		_go_fuzz_dep_.CoverTab[3836]++
//line /usr/local/go/src/strings/strings.go:901
		return cutset[0] < utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:901
		// _ = "end of CoverTab[3836]"
//line /usr/local/go/src/strings/strings.go:901
	}() {
//line /usr/local/go/src/strings/strings.go:901
		_go_fuzz_dep_.CoverTab[3837]++
								return trimLeftByte(s, cutset[0])
//line /usr/local/go/src/strings/strings.go:902
		// _ = "end of CoverTab[3837]"
	} else {
//line /usr/local/go/src/strings/strings.go:903
		_go_fuzz_dep_.CoverTab[3838]++
//line /usr/local/go/src/strings/strings.go:903
		// _ = "end of CoverTab[3838]"
//line /usr/local/go/src/strings/strings.go:903
	}
//line /usr/local/go/src/strings/strings.go:903
	// _ = "end of CoverTab[3830]"
//line /usr/local/go/src/strings/strings.go:903
	_go_fuzz_dep_.CoverTab[3831]++
							if as, ok := makeASCIISet(cutset); ok {
//line /usr/local/go/src/strings/strings.go:904
		_go_fuzz_dep_.CoverTab[3839]++
								return trimLeftASCII(s, &as)
//line /usr/local/go/src/strings/strings.go:905
		// _ = "end of CoverTab[3839]"
	} else {
//line /usr/local/go/src/strings/strings.go:906
		_go_fuzz_dep_.CoverTab[3840]++
//line /usr/local/go/src/strings/strings.go:906
		// _ = "end of CoverTab[3840]"
//line /usr/local/go/src/strings/strings.go:906
	}
//line /usr/local/go/src/strings/strings.go:906
	// _ = "end of CoverTab[3831]"
//line /usr/local/go/src/strings/strings.go:906
	_go_fuzz_dep_.CoverTab[3832]++
							return trimLeftUnicode(s, cutset)
//line /usr/local/go/src/strings/strings.go:907
	// _ = "end of CoverTab[3832]"
}

func trimLeftByte(s string, c byte) string {
//line /usr/local/go/src/strings/strings.go:910
	_go_fuzz_dep_.CoverTab[3841]++
							for len(s) > 0 && func() bool {
//line /usr/local/go/src/strings/strings.go:911
		_go_fuzz_dep_.CoverTab[3843]++
//line /usr/local/go/src/strings/strings.go:911
		return s[0] == c
//line /usr/local/go/src/strings/strings.go:911
		// _ = "end of CoverTab[3843]"
//line /usr/local/go/src/strings/strings.go:911
	}() {
//line /usr/local/go/src/strings/strings.go:911
		_go_fuzz_dep_.CoverTab[3844]++
								s = s[1:]
//line /usr/local/go/src/strings/strings.go:912
		// _ = "end of CoverTab[3844]"
	}
//line /usr/local/go/src/strings/strings.go:913
	// _ = "end of CoverTab[3841]"
//line /usr/local/go/src/strings/strings.go:913
	_go_fuzz_dep_.CoverTab[3842]++
							return s
//line /usr/local/go/src/strings/strings.go:914
	// _ = "end of CoverTab[3842]"
}

func trimLeftASCII(s string, as *asciiSet) string {
//line /usr/local/go/src/strings/strings.go:917
	_go_fuzz_dep_.CoverTab[3845]++
							for len(s) > 0 {
//line /usr/local/go/src/strings/strings.go:918
		_go_fuzz_dep_.CoverTab[3847]++
								if !as.contains(s[0]) {
//line /usr/local/go/src/strings/strings.go:919
			_go_fuzz_dep_.CoverTab[3849]++
									break
//line /usr/local/go/src/strings/strings.go:920
			// _ = "end of CoverTab[3849]"
		} else {
//line /usr/local/go/src/strings/strings.go:921
			_go_fuzz_dep_.CoverTab[3850]++
//line /usr/local/go/src/strings/strings.go:921
			// _ = "end of CoverTab[3850]"
//line /usr/local/go/src/strings/strings.go:921
		}
//line /usr/local/go/src/strings/strings.go:921
		// _ = "end of CoverTab[3847]"
//line /usr/local/go/src/strings/strings.go:921
		_go_fuzz_dep_.CoverTab[3848]++
								s = s[1:]
//line /usr/local/go/src/strings/strings.go:922
		// _ = "end of CoverTab[3848]"
	}
//line /usr/local/go/src/strings/strings.go:923
	// _ = "end of CoverTab[3845]"
//line /usr/local/go/src/strings/strings.go:923
	_go_fuzz_dep_.CoverTab[3846]++
							return s
//line /usr/local/go/src/strings/strings.go:924
	// _ = "end of CoverTab[3846]"
}

func trimLeftUnicode(s, cutset string) string {
//line /usr/local/go/src/strings/strings.go:927
	_go_fuzz_dep_.CoverTab[3851]++
							for len(s) > 0 {
//line /usr/local/go/src/strings/strings.go:928
		_go_fuzz_dep_.CoverTab[3853]++
								r, n := rune(s[0]), 1
								if r >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:930
			_go_fuzz_dep_.CoverTab[3856]++
									r, n = utf8.DecodeRuneInString(s)
//line /usr/local/go/src/strings/strings.go:931
			// _ = "end of CoverTab[3856]"
		} else {
//line /usr/local/go/src/strings/strings.go:932
			_go_fuzz_dep_.CoverTab[3857]++
//line /usr/local/go/src/strings/strings.go:932
			// _ = "end of CoverTab[3857]"
//line /usr/local/go/src/strings/strings.go:932
		}
//line /usr/local/go/src/strings/strings.go:932
		// _ = "end of CoverTab[3853]"
//line /usr/local/go/src/strings/strings.go:932
		_go_fuzz_dep_.CoverTab[3854]++
								if !ContainsRune(cutset, r) {
//line /usr/local/go/src/strings/strings.go:933
			_go_fuzz_dep_.CoverTab[3858]++
									break
//line /usr/local/go/src/strings/strings.go:934
			// _ = "end of CoverTab[3858]"
		} else {
//line /usr/local/go/src/strings/strings.go:935
			_go_fuzz_dep_.CoverTab[3859]++
//line /usr/local/go/src/strings/strings.go:935
			// _ = "end of CoverTab[3859]"
//line /usr/local/go/src/strings/strings.go:935
		}
//line /usr/local/go/src/strings/strings.go:935
		// _ = "end of CoverTab[3854]"
//line /usr/local/go/src/strings/strings.go:935
		_go_fuzz_dep_.CoverTab[3855]++
								s = s[n:]
//line /usr/local/go/src/strings/strings.go:936
		// _ = "end of CoverTab[3855]"
	}
//line /usr/local/go/src/strings/strings.go:937
	// _ = "end of CoverTab[3851]"
//line /usr/local/go/src/strings/strings.go:937
	_go_fuzz_dep_.CoverTab[3852]++
							return s
//line /usr/local/go/src/strings/strings.go:938
	// _ = "end of CoverTab[3852]"
}

// TrimRight returns a slice of the string s, with all trailing
//line /usr/local/go/src/strings/strings.go:941
// Unicode code points contained in cutset removed.
//line /usr/local/go/src/strings/strings.go:941
//
//line /usr/local/go/src/strings/strings.go:941
// To remove a suffix, use TrimSuffix instead.
//line /usr/local/go/src/strings/strings.go:945
func TrimRight(s, cutset string) string {
//line /usr/local/go/src/strings/strings.go:945
	_go_fuzz_dep_.CoverTab[3860]++
							if s == "" || func() bool {
//line /usr/local/go/src/strings/strings.go:946
		_go_fuzz_dep_.CoverTab[3864]++
//line /usr/local/go/src/strings/strings.go:946
		return cutset == ""
//line /usr/local/go/src/strings/strings.go:946
		// _ = "end of CoverTab[3864]"
//line /usr/local/go/src/strings/strings.go:946
	}() {
//line /usr/local/go/src/strings/strings.go:946
		_go_fuzz_dep_.CoverTab[3865]++
								return s
//line /usr/local/go/src/strings/strings.go:947
		// _ = "end of CoverTab[3865]"
	} else {
//line /usr/local/go/src/strings/strings.go:948
		_go_fuzz_dep_.CoverTab[3866]++
//line /usr/local/go/src/strings/strings.go:948
		// _ = "end of CoverTab[3866]"
//line /usr/local/go/src/strings/strings.go:948
	}
//line /usr/local/go/src/strings/strings.go:948
	// _ = "end of CoverTab[3860]"
//line /usr/local/go/src/strings/strings.go:948
	_go_fuzz_dep_.CoverTab[3861]++
							if len(cutset) == 1 && func() bool {
//line /usr/local/go/src/strings/strings.go:949
		_go_fuzz_dep_.CoverTab[3867]++
//line /usr/local/go/src/strings/strings.go:949
		return cutset[0] < utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:949
		// _ = "end of CoverTab[3867]"
//line /usr/local/go/src/strings/strings.go:949
	}() {
//line /usr/local/go/src/strings/strings.go:949
		_go_fuzz_dep_.CoverTab[3868]++
								return trimRightByte(s, cutset[0])
//line /usr/local/go/src/strings/strings.go:950
		// _ = "end of CoverTab[3868]"
	} else {
//line /usr/local/go/src/strings/strings.go:951
		_go_fuzz_dep_.CoverTab[3869]++
//line /usr/local/go/src/strings/strings.go:951
		// _ = "end of CoverTab[3869]"
//line /usr/local/go/src/strings/strings.go:951
	}
//line /usr/local/go/src/strings/strings.go:951
	// _ = "end of CoverTab[3861]"
//line /usr/local/go/src/strings/strings.go:951
	_go_fuzz_dep_.CoverTab[3862]++
							if as, ok := makeASCIISet(cutset); ok {
//line /usr/local/go/src/strings/strings.go:952
		_go_fuzz_dep_.CoverTab[3870]++
								return trimRightASCII(s, &as)
//line /usr/local/go/src/strings/strings.go:953
		// _ = "end of CoverTab[3870]"
	} else {
//line /usr/local/go/src/strings/strings.go:954
		_go_fuzz_dep_.CoverTab[3871]++
//line /usr/local/go/src/strings/strings.go:954
		// _ = "end of CoverTab[3871]"
//line /usr/local/go/src/strings/strings.go:954
	}
//line /usr/local/go/src/strings/strings.go:954
	// _ = "end of CoverTab[3862]"
//line /usr/local/go/src/strings/strings.go:954
	_go_fuzz_dep_.CoverTab[3863]++
							return trimRightUnicode(s, cutset)
//line /usr/local/go/src/strings/strings.go:955
	// _ = "end of CoverTab[3863]"
}

func trimRightByte(s string, c byte) string {
//line /usr/local/go/src/strings/strings.go:958
	_go_fuzz_dep_.CoverTab[3872]++
							for len(s) > 0 && func() bool {
//line /usr/local/go/src/strings/strings.go:959
		_go_fuzz_dep_.CoverTab[3874]++
//line /usr/local/go/src/strings/strings.go:959
		return s[len(s)-1] == c
//line /usr/local/go/src/strings/strings.go:959
		// _ = "end of CoverTab[3874]"
//line /usr/local/go/src/strings/strings.go:959
	}() {
//line /usr/local/go/src/strings/strings.go:959
		_go_fuzz_dep_.CoverTab[3875]++
								s = s[:len(s)-1]
//line /usr/local/go/src/strings/strings.go:960
		// _ = "end of CoverTab[3875]"
	}
//line /usr/local/go/src/strings/strings.go:961
	// _ = "end of CoverTab[3872]"
//line /usr/local/go/src/strings/strings.go:961
	_go_fuzz_dep_.CoverTab[3873]++
							return s
//line /usr/local/go/src/strings/strings.go:962
	// _ = "end of CoverTab[3873]"
}

func trimRightASCII(s string, as *asciiSet) string {
//line /usr/local/go/src/strings/strings.go:965
	_go_fuzz_dep_.CoverTab[3876]++
							for len(s) > 0 {
//line /usr/local/go/src/strings/strings.go:966
		_go_fuzz_dep_.CoverTab[3878]++
								if !as.contains(s[len(s)-1]) {
//line /usr/local/go/src/strings/strings.go:967
			_go_fuzz_dep_.CoverTab[3880]++
									break
//line /usr/local/go/src/strings/strings.go:968
			// _ = "end of CoverTab[3880]"
		} else {
//line /usr/local/go/src/strings/strings.go:969
			_go_fuzz_dep_.CoverTab[3881]++
//line /usr/local/go/src/strings/strings.go:969
			// _ = "end of CoverTab[3881]"
//line /usr/local/go/src/strings/strings.go:969
		}
//line /usr/local/go/src/strings/strings.go:969
		// _ = "end of CoverTab[3878]"
//line /usr/local/go/src/strings/strings.go:969
		_go_fuzz_dep_.CoverTab[3879]++
								s = s[:len(s)-1]
//line /usr/local/go/src/strings/strings.go:970
		// _ = "end of CoverTab[3879]"
	}
//line /usr/local/go/src/strings/strings.go:971
	// _ = "end of CoverTab[3876]"
//line /usr/local/go/src/strings/strings.go:971
	_go_fuzz_dep_.CoverTab[3877]++
							return s
//line /usr/local/go/src/strings/strings.go:972
	// _ = "end of CoverTab[3877]"
}

func trimRightUnicode(s, cutset string) string {
//line /usr/local/go/src/strings/strings.go:975
	_go_fuzz_dep_.CoverTab[3882]++
							for len(s) > 0 {
//line /usr/local/go/src/strings/strings.go:976
		_go_fuzz_dep_.CoverTab[3884]++
								r, n := rune(s[len(s)-1]), 1
								if r >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:978
			_go_fuzz_dep_.CoverTab[3887]++
									r, n = utf8.DecodeLastRuneInString(s)
//line /usr/local/go/src/strings/strings.go:979
			// _ = "end of CoverTab[3887]"
		} else {
//line /usr/local/go/src/strings/strings.go:980
			_go_fuzz_dep_.CoverTab[3888]++
//line /usr/local/go/src/strings/strings.go:980
			// _ = "end of CoverTab[3888]"
//line /usr/local/go/src/strings/strings.go:980
		}
//line /usr/local/go/src/strings/strings.go:980
		// _ = "end of CoverTab[3884]"
//line /usr/local/go/src/strings/strings.go:980
		_go_fuzz_dep_.CoverTab[3885]++
								if !ContainsRune(cutset, r) {
//line /usr/local/go/src/strings/strings.go:981
			_go_fuzz_dep_.CoverTab[3889]++
									break
//line /usr/local/go/src/strings/strings.go:982
			// _ = "end of CoverTab[3889]"
		} else {
//line /usr/local/go/src/strings/strings.go:983
			_go_fuzz_dep_.CoverTab[3890]++
//line /usr/local/go/src/strings/strings.go:983
			// _ = "end of CoverTab[3890]"
//line /usr/local/go/src/strings/strings.go:983
		}
//line /usr/local/go/src/strings/strings.go:983
		// _ = "end of CoverTab[3885]"
//line /usr/local/go/src/strings/strings.go:983
		_go_fuzz_dep_.CoverTab[3886]++
								s = s[:len(s)-n]
//line /usr/local/go/src/strings/strings.go:984
		// _ = "end of CoverTab[3886]"
	}
//line /usr/local/go/src/strings/strings.go:985
	// _ = "end of CoverTab[3882]"
//line /usr/local/go/src/strings/strings.go:985
	_go_fuzz_dep_.CoverTab[3883]++
							return s
//line /usr/local/go/src/strings/strings.go:986
	// _ = "end of CoverTab[3883]"
}

// TrimSpace returns a slice of the string s, with all leading
//line /usr/local/go/src/strings/strings.go:989
// and trailing white space removed, as defined by Unicode.
//line /usr/local/go/src/strings/strings.go:991
func TrimSpace(s string) string {
//line /usr/local/go/src/strings/strings.go:991
	_go_fuzz_dep_.CoverTab[3891]++

							start := 0
							for ; start < len(s); start++ {
//line /usr/local/go/src/strings/strings.go:994
		_go_fuzz_dep_.CoverTab[3894]++
								c := s[start]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:996
			_go_fuzz_dep_.CoverTab[3896]++

//line /usr/local/go/src/strings/strings.go:999
			return TrimFunc(s[start:], unicode.IsSpace)
//line /usr/local/go/src/strings/strings.go:999
			// _ = "end of CoverTab[3896]"
		} else {
//line /usr/local/go/src/strings/strings.go:1000
			_go_fuzz_dep_.CoverTab[3897]++
//line /usr/local/go/src/strings/strings.go:1000
			// _ = "end of CoverTab[3897]"
//line /usr/local/go/src/strings/strings.go:1000
		}
//line /usr/local/go/src/strings/strings.go:1000
		// _ = "end of CoverTab[3894]"
//line /usr/local/go/src/strings/strings.go:1000
		_go_fuzz_dep_.CoverTab[3895]++
								if asciiSpace[c] == 0 {
//line /usr/local/go/src/strings/strings.go:1001
			_go_fuzz_dep_.CoverTab[3898]++
									break
//line /usr/local/go/src/strings/strings.go:1002
			// _ = "end of CoverTab[3898]"
		} else {
//line /usr/local/go/src/strings/strings.go:1003
			_go_fuzz_dep_.CoverTab[3899]++
//line /usr/local/go/src/strings/strings.go:1003
			// _ = "end of CoverTab[3899]"
//line /usr/local/go/src/strings/strings.go:1003
		}
//line /usr/local/go/src/strings/strings.go:1003
		// _ = "end of CoverTab[3895]"
	}
//line /usr/local/go/src/strings/strings.go:1004
	// _ = "end of CoverTab[3891]"
//line /usr/local/go/src/strings/strings.go:1004
	_go_fuzz_dep_.CoverTab[3892]++

//line /usr/local/go/src/strings/strings.go:1007
	stop := len(s)
	for ; stop > start; stop-- {
//line /usr/local/go/src/strings/strings.go:1008
		_go_fuzz_dep_.CoverTab[3900]++
								c := s[stop-1]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:1010
			_go_fuzz_dep_.CoverTab[3902]++

									return TrimRightFunc(s[start:stop], unicode.IsSpace)
//line /usr/local/go/src/strings/strings.go:1012
			// _ = "end of CoverTab[3902]"
		} else {
//line /usr/local/go/src/strings/strings.go:1013
			_go_fuzz_dep_.CoverTab[3903]++
//line /usr/local/go/src/strings/strings.go:1013
			// _ = "end of CoverTab[3903]"
//line /usr/local/go/src/strings/strings.go:1013
		}
//line /usr/local/go/src/strings/strings.go:1013
		// _ = "end of CoverTab[3900]"
//line /usr/local/go/src/strings/strings.go:1013
		_go_fuzz_dep_.CoverTab[3901]++
								if asciiSpace[c] == 0 {
//line /usr/local/go/src/strings/strings.go:1014
			_go_fuzz_dep_.CoverTab[3904]++
									break
//line /usr/local/go/src/strings/strings.go:1015
			// _ = "end of CoverTab[3904]"
		} else {
//line /usr/local/go/src/strings/strings.go:1016
			_go_fuzz_dep_.CoverTab[3905]++
//line /usr/local/go/src/strings/strings.go:1016
			// _ = "end of CoverTab[3905]"
//line /usr/local/go/src/strings/strings.go:1016
		}
//line /usr/local/go/src/strings/strings.go:1016
		// _ = "end of CoverTab[3901]"
	}
//line /usr/local/go/src/strings/strings.go:1017
	// _ = "end of CoverTab[3892]"
//line /usr/local/go/src/strings/strings.go:1017
	_go_fuzz_dep_.CoverTab[3893]++

//line /usr/local/go/src/strings/strings.go:1022
	return s[start:stop]
//line /usr/local/go/src/strings/strings.go:1022
	// _ = "end of CoverTab[3893]"
}

// TrimPrefix returns s without the provided leading prefix string.
//line /usr/local/go/src/strings/strings.go:1025
// If s doesn't start with prefix, s is returned unchanged.
//line /usr/local/go/src/strings/strings.go:1027
func TrimPrefix(s, prefix string) string {
//line /usr/local/go/src/strings/strings.go:1027
	_go_fuzz_dep_.CoverTab[3906]++
							if HasPrefix(s, prefix) {
//line /usr/local/go/src/strings/strings.go:1028
		_go_fuzz_dep_.CoverTab[3908]++
								return s[len(prefix):]
//line /usr/local/go/src/strings/strings.go:1029
		// _ = "end of CoverTab[3908]"
	} else {
//line /usr/local/go/src/strings/strings.go:1030
		_go_fuzz_dep_.CoverTab[3909]++
//line /usr/local/go/src/strings/strings.go:1030
		// _ = "end of CoverTab[3909]"
//line /usr/local/go/src/strings/strings.go:1030
	}
//line /usr/local/go/src/strings/strings.go:1030
	// _ = "end of CoverTab[3906]"
//line /usr/local/go/src/strings/strings.go:1030
	_go_fuzz_dep_.CoverTab[3907]++
							return s
//line /usr/local/go/src/strings/strings.go:1031
	// _ = "end of CoverTab[3907]"
}

// TrimSuffix returns s without the provided trailing suffix string.
//line /usr/local/go/src/strings/strings.go:1034
// If s doesn't end with suffix, s is returned unchanged.
//line /usr/local/go/src/strings/strings.go:1036
func TrimSuffix(s, suffix string) string {
//line /usr/local/go/src/strings/strings.go:1036
	_go_fuzz_dep_.CoverTab[3910]++
							if HasSuffix(s, suffix) {
//line /usr/local/go/src/strings/strings.go:1037
		_go_fuzz_dep_.CoverTab[3912]++
								return s[:len(s)-len(suffix)]
//line /usr/local/go/src/strings/strings.go:1038
		// _ = "end of CoverTab[3912]"
	} else {
//line /usr/local/go/src/strings/strings.go:1039
		_go_fuzz_dep_.CoverTab[3913]++
//line /usr/local/go/src/strings/strings.go:1039
		// _ = "end of CoverTab[3913]"
//line /usr/local/go/src/strings/strings.go:1039
	}
//line /usr/local/go/src/strings/strings.go:1039
	// _ = "end of CoverTab[3910]"
//line /usr/local/go/src/strings/strings.go:1039
	_go_fuzz_dep_.CoverTab[3911]++
							return s
//line /usr/local/go/src/strings/strings.go:1040
	// _ = "end of CoverTab[3911]"
}

// Replace returns a copy of the string s with the first n
//line /usr/local/go/src/strings/strings.go:1043
// non-overlapping instances of old replaced by new.
//line /usr/local/go/src/strings/strings.go:1043
// If old is empty, it matches at the beginning of the string
//line /usr/local/go/src/strings/strings.go:1043
// and after each UTF-8 sequence, yielding up to k+1 replacements
//line /usr/local/go/src/strings/strings.go:1043
// for a k-rune string.
//line /usr/local/go/src/strings/strings.go:1043
// If n < 0, there is no limit on the number of replacements.
//line /usr/local/go/src/strings/strings.go:1049
func Replace(s, old, new string, n int) string {
//line /usr/local/go/src/strings/strings.go:1049
	_go_fuzz_dep_.CoverTab[3914]++
							if old == new || func() bool {
//line /usr/local/go/src/strings/strings.go:1050
		_go_fuzz_dep_.CoverTab[3918]++
//line /usr/local/go/src/strings/strings.go:1050
		return n == 0
//line /usr/local/go/src/strings/strings.go:1050
		// _ = "end of CoverTab[3918]"
//line /usr/local/go/src/strings/strings.go:1050
	}() {
//line /usr/local/go/src/strings/strings.go:1050
		_go_fuzz_dep_.CoverTab[3919]++
								return s
//line /usr/local/go/src/strings/strings.go:1051
		// _ = "end of CoverTab[3919]"
	} else {
//line /usr/local/go/src/strings/strings.go:1052
		_go_fuzz_dep_.CoverTab[3920]++
//line /usr/local/go/src/strings/strings.go:1052
		// _ = "end of CoverTab[3920]"
//line /usr/local/go/src/strings/strings.go:1052
	}
//line /usr/local/go/src/strings/strings.go:1052
	// _ = "end of CoverTab[3914]"
//line /usr/local/go/src/strings/strings.go:1052
	_go_fuzz_dep_.CoverTab[3915]++

//line /usr/local/go/src/strings/strings.go:1055
	if m := Count(s, old); m == 0 {
//line /usr/local/go/src/strings/strings.go:1055
		_go_fuzz_dep_.CoverTab[3921]++
								return s
//line /usr/local/go/src/strings/strings.go:1056
		// _ = "end of CoverTab[3921]"
	} else {
//line /usr/local/go/src/strings/strings.go:1057
		_go_fuzz_dep_.CoverTab[3922]++
//line /usr/local/go/src/strings/strings.go:1057
		if n < 0 || func() bool {
//line /usr/local/go/src/strings/strings.go:1057
			_go_fuzz_dep_.CoverTab[3923]++
//line /usr/local/go/src/strings/strings.go:1057
			return m < n
//line /usr/local/go/src/strings/strings.go:1057
			// _ = "end of CoverTab[3923]"
//line /usr/local/go/src/strings/strings.go:1057
		}() {
//line /usr/local/go/src/strings/strings.go:1057
			_go_fuzz_dep_.CoverTab[3924]++
									n = m
//line /usr/local/go/src/strings/strings.go:1058
			// _ = "end of CoverTab[3924]"
		} else {
//line /usr/local/go/src/strings/strings.go:1059
			_go_fuzz_dep_.CoverTab[3925]++
//line /usr/local/go/src/strings/strings.go:1059
			// _ = "end of CoverTab[3925]"
//line /usr/local/go/src/strings/strings.go:1059
		}
//line /usr/local/go/src/strings/strings.go:1059
		// _ = "end of CoverTab[3922]"
//line /usr/local/go/src/strings/strings.go:1059
	}
//line /usr/local/go/src/strings/strings.go:1059
	// _ = "end of CoverTab[3915]"
//line /usr/local/go/src/strings/strings.go:1059
	_go_fuzz_dep_.CoverTab[3916]++

	// Apply replacements to buffer.
	var b Builder
	b.Grow(len(s) + n*(len(new)-len(old)))
	start := 0
	for i := 0; i < n; i++ {
//line /usr/local/go/src/strings/strings.go:1065
		_go_fuzz_dep_.CoverTab[3926]++
								j := start
								if len(old) == 0 {
//line /usr/local/go/src/strings/strings.go:1067
			_go_fuzz_dep_.CoverTab[3928]++
									if i > 0 {
//line /usr/local/go/src/strings/strings.go:1068
				_go_fuzz_dep_.CoverTab[3929]++
										_, wid := utf8.DecodeRuneInString(s[start:])
										j += wid
//line /usr/local/go/src/strings/strings.go:1070
				// _ = "end of CoverTab[3929]"
			} else {
//line /usr/local/go/src/strings/strings.go:1071
				_go_fuzz_dep_.CoverTab[3930]++
//line /usr/local/go/src/strings/strings.go:1071
				// _ = "end of CoverTab[3930]"
//line /usr/local/go/src/strings/strings.go:1071
			}
//line /usr/local/go/src/strings/strings.go:1071
			// _ = "end of CoverTab[3928]"
		} else {
//line /usr/local/go/src/strings/strings.go:1072
			_go_fuzz_dep_.CoverTab[3931]++
									j += Index(s[start:], old)
//line /usr/local/go/src/strings/strings.go:1073
			// _ = "end of CoverTab[3931]"
		}
//line /usr/local/go/src/strings/strings.go:1074
		// _ = "end of CoverTab[3926]"
//line /usr/local/go/src/strings/strings.go:1074
		_go_fuzz_dep_.CoverTab[3927]++
								b.WriteString(s[start:j])
								b.WriteString(new)
								start = j + len(old)
//line /usr/local/go/src/strings/strings.go:1077
		// _ = "end of CoverTab[3927]"
	}
//line /usr/local/go/src/strings/strings.go:1078
	// _ = "end of CoverTab[3916]"
//line /usr/local/go/src/strings/strings.go:1078
	_go_fuzz_dep_.CoverTab[3917]++
							b.WriteString(s[start:])
							return b.String()
//line /usr/local/go/src/strings/strings.go:1080
	// _ = "end of CoverTab[3917]"
}

// ReplaceAll returns a copy of the string s with all
//line /usr/local/go/src/strings/strings.go:1083
// non-overlapping instances of old replaced by new.
//line /usr/local/go/src/strings/strings.go:1083
// If old is empty, it matches at the beginning of the string
//line /usr/local/go/src/strings/strings.go:1083
// and after each UTF-8 sequence, yielding up to k+1 replacements
//line /usr/local/go/src/strings/strings.go:1083
// for a k-rune string.
//line /usr/local/go/src/strings/strings.go:1088
func ReplaceAll(s, old, new string) string {
//line /usr/local/go/src/strings/strings.go:1088
	_go_fuzz_dep_.CoverTab[3932]++
							return Replace(s, old, new, -1)
//line /usr/local/go/src/strings/strings.go:1089
	// _ = "end of CoverTab[3932]"
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
//line /usr/local/go/src/strings/strings.go:1092
// are equal under simple Unicode case-folding, which is a more general
//line /usr/local/go/src/strings/strings.go:1092
// form of case-insensitivity.
//line /usr/local/go/src/strings/strings.go:1095
func EqualFold(s, t string) bool {
//line /usr/local/go/src/strings/strings.go:1095
	_go_fuzz_dep_.CoverTab[3933]++

							i := 0
							for ; i < len(s) && func() bool {
//line /usr/local/go/src/strings/strings.go:1098
		_go_fuzz_dep_.CoverTab[3936]++
//line /usr/local/go/src/strings/strings.go:1098
		return i < len(t)
//line /usr/local/go/src/strings/strings.go:1098
		// _ = "end of CoverTab[3936]"
//line /usr/local/go/src/strings/strings.go:1098
	}(); i++ {
//line /usr/local/go/src/strings/strings.go:1098
		_go_fuzz_dep_.CoverTab[3937]++
								sr := s[i]
								tr := t[i]
								if sr|tr >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:1101
			_go_fuzz_dep_.CoverTab[3942]++
									goto hasUnicode
//line /usr/local/go/src/strings/strings.go:1102
			// _ = "end of CoverTab[3942]"
		} else {
//line /usr/local/go/src/strings/strings.go:1103
			_go_fuzz_dep_.CoverTab[3943]++
//line /usr/local/go/src/strings/strings.go:1103
			// _ = "end of CoverTab[3943]"
//line /usr/local/go/src/strings/strings.go:1103
		}
//line /usr/local/go/src/strings/strings.go:1103
		// _ = "end of CoverTab[3937]"
//line /usr/local/go/src/strings/strings.go:1103
		_go_fuzz_dep_.CoverTab[3938]++

//line /usr/local/go/src/strings/strings.go:1106
		if tr == sr {
//line /usr/local/go/src/strings/strings.go:1106
			_go_fuzz_dep_.CoverTab[3944]++
									continue
//line /usr/local/go/src/strings/strings.go:1107
			// _ = "end of CoverTab[3944]"
		} else {
//line /usr/local/go/src/strings/strings.go:1108
			_go_fuzz_dep_.CoverTab[3945]++
//line /usr/local/go/src/strings/strings.go:1108
			// _ = "end of CoverTab[3945]"
//line /usr/local/go/src/strings/strings.go:1108
		}
//line /usr/local/go/src/strings/strings.go:1108
		// _ = "end of CoverTab[3938]"
//line /usr/local/go/src/strings/strings.go:1108
		_go_fuzz_dep_.CoverTab[3939]++

//line /usr/local/go/src/strings/strings.go:1111
		if tr < sr {
//line /usr/local/go/src/strings/strings.go:1111
			_go_fuzz_dep_.CoverTab[3946]++
									tr, sr = sr, tr
//line /usr/local/go/src/strings/strings.go:1112
			// _ = "end of CoverTab[3946]"
		} else {
//line /usr/local/go/src/strings/strings.go:1113
			_go_fuzz_dep_.CoverTab[3947]++
//line /usr/local/go/src/strings/strings.go:1113
			// _ = "end of CoverTab[3947]"
//line /usr/local/go/src/strings/strings.go:1113
		}
//line /usr/local/go/src/strings/strings.go:1113
		// _ = "end of CoverTab[3939]"
//line /usr/local/go/src/strings/strings.go:1113
		_go_fuzz_dep_.CoverTab[3940]++

								if 'A' <= sr && func() bool {
//line /usr/local/go/src/strings/strings.go:1115
			_go_fuzz_dep_.CoverTab[3948]++
//line /usr/local/go/src/strings/strings.go:1115
			return sr <= 'Z'
//line /usr/local/go/src/strings/strings.go:1115
			// _ = "end of CoverTab[3948]"
//line /usr/local/go/src/strings/strings.go:1115
		}() && func() bool {
//line /usr/local/go/src/strings/strings.go:1115
			_go_fuzz_dep_.CoverTab[3949]++
//line /usr/local/go/src/strings/strings.go:1115
			return tr == sr+'a'-'A'
//line /usr/local/go/src/strings/strings.go:1115
			// _ = "end of CoverTab[3949]"
//line /usr/local/go/src/strings/strings.go:1115
		}() {
//line /usr/local/go/src/strings/strings.go:1115
			_go_fuzz_dep_.CoverTab[3950]++
									continue
//line /usr/local/go/src/strings/strings.go:1116
			// _ = "end of CoverTab[3950]"
		} else {
//line /usr/local/go/src/strings/strings.go:1117
			_go_fuzz_dep_.CoverTab[3951]++
//line /usr/local/go/src/strings/strings.go:1117
			// _ = "end of CoverTab[3951]"
//line /usr/local/go/src/strings/strings.go:1117
		}
//line /usr/local/go/src/strings/strings.go:1117
		// _ = "end of CoverTab[3940]"
//line /usr/local/go/src/strings/strings.go:1117
		_go_fuzz_dep_.CoverTab[3941]++
								return false
//line /usr/local/go/src/strings/strings.go:1118
		// _ = "end of CoverTab[3941]"
	}
//line /usr/local/go/src/strings/strings.go:1119
	// _ = "end of CoverTab[3933]"
//line /usr/local/go/src/strings/strings.go:1119
	_go_fuzz_dep_.CoverTab[3934]++

							return len(s) == len(t)

hasUnicode:
	s = s[i:]
	t = t[i:]
	for _, sr := range s {
//line /usr/local/go/src/strings/strings.go:1126
		_go_fuzz_dep_.CoverTab[3952]++

								if len(t) == 0 {
//line /usr/local/go/src/strings/strings.go:1128
			_go_fuzz_dep_.CoverTab[3960]++
									return false
//line /usr/local/go/src/strings/strings.go:1129
			// _ = "end of CoverTab[3960]"
		} else {
//line /usr/local/go/src/strings/strings.go:1130
			_go_fuzz_dep_.CoverTab[3961]++
//line /usr/local/go/src/strings/strings.go:1130
			// _ = "end of CoverTab[3961]"
//line /usr/local/go/src/strings/strings.go:1130
		}
//line /usr/local/go/src/strings/strings.go:1130
		// _ = "end of CoverTab[3952]"
//line /usr/local/go/src/strings/strings.go:1130
		_go_fuzz_dep_.CoverTab[3953]++

		// Extract first rune from second string.
		var tr rune
		if t[0] < utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:1134
			_go_fuzz_dep_.CoverTab[3962]++
									tr, t = rune(t[0]), t[1:]
//line /usr/local/go/src/strings/strings.go:1135
			// _ = "end of CoverTab[3962]"
		} else {
//line /usr/local/go/src/strings/strings.go:1136
			_go_fuzz_dep_.CoverTab[3963]++
									r, size := utf8.DecodeRuneInString(t)
									tr, t = r, t[size:]
//line /usr/local/go/src/strings/strings.go:1138
			// _ = "end of CoverTab[3963]"
		}
//line /usr/local/go/src/strings/strings.go:1139
		// _ = "end of CoverTab[3953]"
//line /usr/local/go/src/strings/strings.go:1139
		_go_fuzz_dep_.CoverTab[3954]++

//line /usr/local/go/src/strings/strings.go:1144
		if tr == sr {
//line /usr/local/go/src/strings/strings.go:1144
			_go_fuzz_dep_.CoverTab[3964]++
									continue
//line /usr/local/go/src/strings/strings.go:1145
			// _ = "end of CoverTab[3964]"
		} else {
//line /usr/local/go/src/strings/strings.go:1146
			_go_fuzz_dep_.CoverTab[3965]++
//line /usr/local/go/src/strings/strings.go:1146
			// _ = "end of CoverTab[3965]"
//line /usr/local/go/src/strings/strings.go:1146
		}
//line /usr/local/go/src/strings/strings.go:1146
		// _ = "end of CoverTab[3954]"
//line /usr/local/go/src/strings/strings.go:1146
		_go_fuzz_dep_.CoverTab[3955]++

//line /usr/local/go/src/strings/strings.go:1149
		if tr < sr {
//line /usr/local/go/src/strings/strings.go:1149
			_go_fuzz_dep_.CoverTab[3966]++
									tr, sr = sr, tr
//line /usr/local/go/src/strings/strings.go:1150
			// _ = "end of CoverTab[3966]"
		} else {
//line /usr/local/go/src/strings/strings.go:1151
			_go_fuzz_dep_.CoverTab[3967]++
//line /usr/local/go/src/strings/strings.go:1151
			// _ = "end of CoverTab[3967]"
//line /usr/local/go/src/strings/strings.go:1151
		}
//line /usr/local/go/src/strings/strings.go:1151
		// _ = "end of CoverTab[3955]"
//line /usr/local/go/src/strings/strings.go:1151
		_go_fuzz_dep_.CoverTab[3956]++

								if tr < utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:1153
			_go_fuzz_dep_.CoverTab[3968]++

									if 'A' <= sr && func() bool {
//line /usr/local/go/src/strings/strings.go:1155
				_go_fuzz_dep_.CoverTab[3970]++
//line /usr/local/go/src/strings/strings.go:1155
				return sr <= 'Z'
//line /usr/local/go/src/strings/strings.go:1155
				// _ = "end of CoverTab[3970]"
//line /usr/local/go/src/strings/strings.go:1155
			}() && func() bool {
//line /usr/local/go/src/strings/strings.go:1155
				_go_fuzz_dep_.CoverTab[3971]++
//line /usr/local/go/src/strings/strings.go:1155
				return tr == sr+'a'-'A'
//line /usr/local/go/src/strings/strings.go:1155
				// _ = "end of CoverTab[3971]"
//line /usr/local/go/src/strings/strings.go:1155
			}() {
//line /usr/local/go/src/strings/strings.go:1155
				_go_fuzz_dep_.CoverTab[3972]++
										continue
//line /usr/local/go/src/strings/strings.go:1156
				// _ = "end of CoverTab[3972]"
			} else {
//line /usr/local/go/src/strings/strings.go:1157
				_go_fuzz_dep_.CoverTab[3973]++
//line /usr/local/go/src/strings/strings.go:1157
				// _ = "end of CoverTab[3973]"
//line /usr/local/go/src/strings/strings.go:1157
			}
//line /usr/local/go/src/strings/strings.go:1157
			// _ = "end of CoverTab[3968]"
//line /usr/local/go/src/strings/strings.go:1157
			_go_fuzz_dep_.CoverTab[3969]++
									return false
//line /usr/local/go/src/strings/strings.go:1158
			// _ = "end of CoverTab[3969]"
		} else {
//line /usr/local/go/src/strings/strings.go:1159
			_go_fuzz_dep_.CoverTab[3974]++
//line /usr/local/go/src/strings/strings.go:1159
			// _ = "end of CoverTab[3974]"
//line /usr/local/go/src/strings/strings.go:1159
		}
//line /usr/local/go/src/strings/strings.go:1159
		// _ = "end of CoverTab[3956]"
//line /usr/local/go/src/strings/strings.go:1159
		_go_fuzz_dep_.CoverTab[3957]++

//line /usr/local/go/src/strings/strings.go:1163
		r := unicode.SimpleFold(sr)
		for r != sr && func() bool {
//line /usr/local/go/src/strings/strings.go:1164
			_go_fuzz_dep_.CoverTab[3975]++
//line /usr/local/go/src/strings/strings.go:1164
			return r < tr
//line /usr/local/go/src/strings/strings.go:1164
			// _ = "end of CoverTab[3975]"
//line /usr/local/go/src/strings/strings.go:1164
		}() {
//line /usr/local/go/src/strings/strings.go:1164
			_go_fuzz_dep_.CoverTab[3976]++
									r = unicode.SimpleFold(r)
//line /usr/local/go/src/strings/strings.go:1165
			// _ = "end of CoverTab[3976]"
		}
//line /usr/local/go/src/strings/strings.go:1166
		// _ = "end of CoverTab[3957]"
//line /usr/local/go/src/strings/strings.go:1166
		_go_fuzz_dep_.CoverTab[3958]++
								if r == tr {
//line /usr/local/go/src/strings/strings.go:1167
			_go_fuzz_dep_.CoverTab[3977]++
									continue
//line /usr/local/go/src/strings/strings.go:1168
			// _ = "end of CoverTab[3977]"
		} else {
//line /usr/local/go/src/strings/strings.go:1169
			_go_fuzz_dep_.CoverTab[3978]++
//line /usr/local/go/src/strings/strings.go:1169
			// _ = "end of CoverTab[3978]"
//line /usr/local/go/src/strings/strings.go:1169
		}
//line /usr/local/go/src/strings/strings.go:1169
		// _ = "end of CoverTab[3958]"
//line /usr/local/go/src/strings/strings.go:1169
		_go_fuzz_dep_.CoverTab[3959]++
								return false
//line /usr/local/go/src/strings/strings.go:1170
		// _ = "end of CoverTab[3959]"
	}
//line /usr/local/go/src/strings/strings.go:1171
	// _ = "end of CoverTab[3934]"
//line /usr/local/go/src/strings/strings.go:1171
	_go_fuzz_dep_.CoverTab[3935]++

//line /usr/local/go/src/strings/strings.go:1174
	return len(t) == 0
//line /usr/local/go/src/strings/strings.go:1174
	// _ = "end of CoverTab[3935]"
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
//line /usr/local/go/src/strings/strings.go:1178
	_go_fuzz_dep_.CoverTab[3979]++
							n := len(substr)
							switch {
	case n == 0:
//line /usr/local/go/src/strings/strings.go:1181
		_go_fuzz_dep_.CoverTab[3982]++
								return 0
//line /usr/local/go/src/strings/strings.go:1182
		// _ = "end of CoverTab[3982]"
	case n == 1:
//line /usr/local/go/src/strings/strings.go:1183
		_go_fuzz_dep_.CoverTab[3983]++
								return IndexByte(s, substr[0])
//line /usr/local/go/src/strings/strings.go:1184
		// _ = "end of CoverTab[3983]"
	case n == len(s):
//line /usr/local/go/src/strings/strings.go:1185
		_go_fuzz_dep_.CoverTab[3984]++
								if substr == s {
//line /usr/local/go/src/strings/strings.go:1186
			_go_fuzz_dep_.CoverTab[3991]++
									return 0
//line /usr/local/go/src/strings/strings.go:1187
			// _ = "end of CoverTab[3991]"
		} else {
//line /usr/local/go/src/strings/strings.go:1188
			_go_fuzz_dep_.CoverTab[3992]++
//line /usr/local/go/src/strings/strings.go:1188
			// _ = "end of CoverTab[3992]"
//line /usr/local/go/src/strings/strings.go:1188
		}
//line /usr/local/go/src/strings/strings.go:1188
		// _ = "end of CoverTab[3984]"
//line /usr/local/go/src/strings/strings.go:1188
		_go_fuzz_dep_.CoverTab[3985]++
								return -1
//line /usr/local/go/src/strings/strings.go:1189
		// _ = "end of CoverTab[3985]"
	case n > len(s):
//line /usr/local/go/src/strings/strings.go:1190
		_go_fuzz_dep_.CoverTab[3986]++
								return -1
//line /usr/local/go/src/strings/strings.go:1191
		// _ = "end of CoverTab[3986]"
	case n <= bytealg.MaxLen:
//line /usr/local/go/src/strings/strings.go:1192
		_go_fuzz_dep_.CoverTab[3987]++

								if len(s) <= bytealg.MaxBruteForce {
//line /usr/local/go/src/strings/strings.go:1194
			_go_fuzz_dep_.CoverTab[3993]++
									return bytealg.IndexString(s, substr)
//line /usr/local/go/src/strings/strings.go:1195
			// _ = "end of CoverTab[3993]"
		} else {
//line /usr/local/go/src/strings/strings.go:1196
			_go_fuzz_dep_.CoverTab[3994]++
//line /usr/local/go/src/strings/strings.go:1196
			// _ = "end of CoverTab[3994]"
//line /usr/local/go/src/strings/strings.go:1196
		}
//line /usr/local/go/src/strings/strings.go:1196
		// _ = "end of CoverTab[3987]"
//line /usr/local/go/src/strings/strings.go:1196
		_go_fuzz_dep_.CoverTab[3988]++
								c0 := substr[0]
								c1 := substr[1]
								i := 0
								t := len(s) - n + 1
								fails := 0
								for i < t {
//line /usr/local/go/src/strings/strings.go:1202
			_go_fuzz_dep_.CoverTab[3995]++
									if s[i] != c0 {
//line /usr/local/go/src/strings/strings.go:1203
				_go_fuzz_dep_.CoverTab[3998]++

//line /usr/local/go/src/strings/strings.go:1206
				o := IndexByte(s[i+1:t], c0)
				if o < 0 {
//line /usr/local/go/src/strings/strings.go:1207
					_go_fuzz_dep_.CoverTab[4000]++
											return -1
//line /usr/local/go/src/strings/strings.go:1208
					// _ = "end of CoverTab[4000]"
				} else {
//line /usr/local/go/src/strings/strings.go:1209
					_go_fuzz_dep_.CoverTab[4001]++
//line /usr/local/go/src/strings/strings.go:1209
					// _ = "end of CoverTab[4001]"
//line /usr/local/go/src/strings/strings.go:1209
				}
//line /usr/local/go/src/strings/strings.go:1209
				// _ = "end of CoverTab[3998]"
//line /usr/local/go/src/strings/strings.go:1209
				_go_fuzz_dep_.CoverTab[3999]++
										i += o + 1
//line /usr/local/go/src/strings/strings.go:1210
				// _ = "end of CoverTab[3999]"
			} else {
//line /usr/local/go/src/strings/strings.go:1211
				_go_fuzz_dep_.CoverTab[4002]++
//line /usr/local/go/src/strings/strings.go:1211
				// _ = "end of CoverTab[4002]"
//line /usr/local/go/src/strings/strings.go:1211
			}
//line /usr/local/go/src/strings/strings.go:1211
			// _ = "end of CoverTab[3995]"
//line /usr/local/go/src/strings/strings.go:1211
			_go_fuzz_dep_.CoverTab[3996]++
									if s[i+1] == c1 && func() bool {
//line /usr/local/go/src/strings/strings.go:1212
				_go_fuzz_dep_.CoverTab[4003]++
//line /usr/local/go/src/strings/strings.go:1212
				return s[i:i+n] == substr
//line /usr/local/go/src/strings/strings.go:1212
				// _ = "end of CoverTab[4003]"
//line /usr/local/go/src/strings/strings.go:1212
			}() {
//line /usr/local/go/src/strings/strings.go:1212
				_go_fuzz_dep_.CoverTab[4004]++
										return i
//line /usr/local/go/src/strings/strings.go:1213
				// _ = "end of CoverTab[4004]"
			} else {
//line /usr/local/go/src/strings/strings.go:1214
				_go_fuzz_dep_.CoverTab[4005]++
//line /usr/local/go/src/strings/strings.go:1214
				// _ = "end of CoverTab[4005]"
//line /usr/local/go/src/strings/strings.go:1214
			}
//line /usr/local/go/src/strings/strings.go:1214
			// _ = "end of CoverTab[3996]"
//line /usr/local/go/src/strings/strings.go:1214
			_go_fuzz_dep_.CoverTab[3997]++
									fails++
									i++

									if fails > bytealg.Cutover(i) {
//line /usr/local/go/src/strings/strings.go:1218
				_go_fuzz_dep_.CoverTab[4006]++
										r := bytealg.IndexString(s[i:], substr)
										if r >= 0 {
//line /usr/local/go/src/strings/strings.go:1220
					_go_fuzz_dep_.CoverTab[4008]++
											return r + i
//line /usr/local/go/src/strings/strings.go:1221
					// _ = "end of CoverTab[4008]"
				} else {
//line /usr/local/go/src/strings/strings.go:1222
					_go_fuzz_dep_.CoverTab[4009]++
//line /usr/local/go/src/strings/strings.go:1222
					// _ = "end of CoverTab[4009]"
//line /usr/local/go/src/strings/strings.go:1222
				}
//line /usr/local/go/src/strings/strings.go:1222
				// _ = "end of CoverTab[4006]"
//line /usr/local/go/src/strings/strings.go:1222
				_go_fuzz_dep_.CoverTab[4007]++
										return -1
//line /usr/local/go/src/strings/strings.go:1223
				// _ = "end of CoverTab[4007]"
			} else {
//line /usr/local/go/src/strings/strings.go:1224
				_go_fuzz_dep_.CoverTab[4010]++
//line /usr/local/go/src/strings/strings.go:1224
				// _ = "end of CoverTab[4010]"
//line /usr/local/go/src/strings/strings.go:1224
			}
//line /usr/local/go/src/strings/strings.go:1224
			// _ = "end of CoverTab[3997]"
		}
//line /usr/local/go/src/strings/strings.go:1225
		// _ = "end of CoverTab[3988]"
//line /usr/local/go/src/strings/strings.go:1225
		_go_fuzz_dep_.CoverTab[3989]++
								return -1
//line /usr/local/go/src/strings/strings.go:1226
		// _ = "end of CoverTab[3989]"
//line /usr/local/go/src/strings/strings.go:1226
	default:
//line /usr/local/go/src/strings/strings.go:1226
		_go_fuzz_dep_.CoverTab[3990]++
//line /usr/local/go/src/strings/strings.go:1226
		// _ = "end of CoverTab[3990]"
	}
//line /usr/local/go/src/strings/strings.go:1227
	// _ = "end of CoverTab[3979]"
//line /usr/local/go/src/strings/strings.go:1227
	_go_fuzz_dep_.CoverTab[3980]++
							c0 := substr[0]
							c1 := substr[1]
							i := 0
							t := len(s) - n + 1
							fails := 0
							for i < t {
//line /usr/local/go/src/strings/strings.go:1233
		_go_fuzz_dep_.CoverTab[4011]++
								if s[i] != c0 {
//line /usr/local/go/src/strings/strings.go:1234
			_go_fuzz_dep_.CoverTab[4014]++
									o := IndexByte(s[i+1:t], c0)
									if o < 0 {
//line /usr/local/go/src/strings/strings.go:1236
				_go_fuzz_dep_.CoverTab[4016]++
										return -1
//line /usr/local/go/src/strings/strings.go:1237
				// _ = "end of CoverTab[4016]"
			} else {
//line /usr/local/go/src/strings/strings.go:1238
				_go_fuzz_dep_.CoverTab[4017]++
//line /usr/local/go/src/strings/strings.go:1238
				// _ = "end of CoverTab[4017]"
//line /usr/local/go/src/strings/strings.go:1238
			}
//line /usr/local/go/src/strings/strings.go:1238
			// _ = "end of CoverTab[4014]"
//line /usr/local/go/src/strings/strings.go:1238
			_go_fuzz_dep_.CoverTab[4015]++
									i += o + 1
//line /usr/local/go/src/strings/strings.go:1239
			// _ = "end of CoverTab[4015]"
		} else {
//line /usr/local/go/src/strings/strings.go:1240
			_go_fuzz_dep_.CoverTab[4018]++
//line /usr/local/go/src/strings/strings.go:1240
			// _ = "end of CoverTab[4018]"
//line /usr/local/go/src/strings/strings.go:1240
		}
//line /usr/local/go/src/strings/strings.go:1240
		// _ = "end of CoverTab[4011]"
//line /usr/local/go/src/strings/strings.go:1240
		_go_fuzz_dep_.CoverTab[4012]++
								if s[i+1] == c1 && func() bool {
//line /usr/local/go/src/strings/strings.go:1241
			_go_fuzz_dep_.CoverTab[4019]++
//line /usr/local/go/src/strings/strings.go:1241
			return s[i:i+n] == substr
//line /usr/local/go/src/strings/strings.go:1241
			// _ = "end of CoverTab[4019]"
//line /usr/local/go/src/strings/strings.go:1241
		}() {
//line /usr/local/go/src/strings/strings.go:1241
			_go_fuzz_dep_.CoverTab[4020]++
									return i
//line /usr/local/go/src/strings/strings.go:1242
			// _ = "end of CoverTab[4020]"
		} else {
//line /usr/local/go/src/strings/strings.go:1243
			_go_fuzz_dep_.CoverTab[4021]++
//line /usr/local/go/src/strings/strings.go:1243
			// _ = "end of CoverTab[4021]"
//line /usr/local/go/src/strings/strings.go:1243
		}
//line /usr/local/go/src/strings/strings.go:1243
		// _ = "end of CoverTab[4012]"
//line /usr/local/go/src/strings/strings.go:1243
		_go_fuzz_dep_.CoverTab[4013]++
								i++
								fails++
								if fails >= 4+i>>4 && func() bool {
//line /usr/local/go/src/strings/strings.go:1246
			_go_fuzz_dep_.CoverTab[4022]++
//line /usr/local/go/src/strings/strings.go:1246
			return i < t
//line /usr/local/go/src/strings/strings.go:1246
			// _ = "end of CoverTab[4022]"
//line /usr/local/go/src/strings/strings.go:1246
		}() {
//line /usr/local/go/src/strings/strings.go:1246
			_go_fuzz_dep_.CoverTab[4023]++

									j := bytealg.IndexRabinKarp(s[i:], substr)
									if j < 0 {
//line /usr/local/go/src/strings/strings.go:1249
				_go_fuzz_dep_.CoverTab[4025]++
										return -1
//line /usr/local/go/src/strings/strings.go:1250
				// _ = "end of CoverTab[4025]"
			} else {
//line /usr/local/go/src/strings/strings.go:1251
				_go_fuzz_dep_.CoverTab[4026]++
//line /usr/local/go/src/strings/strings.go:1251
				// _ = "end of CoverTab[4026]"
//line /usr/local/go/src/strings/strings.go:1251
			}
//line /usr/local/go/src/strings/strings.go:1251
			// _ = "end of CoverTab[4023]"
//line /usr/local/go/src/strings/strings.go:1251
			_go_fuzz_dep_.CoverTab[4024]++
									return i + j
//line /usr/local/go/src/strings/strings.go:1252
			// _ = "end of CoverTab[4024]"
		} else {
//line /usr/local/go/src/strings/strings.go:1253
			_go_fuzz_dep_.CoverTab[4027]++
//line /usr/local/go/src/strings/strings.go:1253
			// _ = "end of CoverTab[4027]"
//line /usr/local/go/src/strings/strings.go:1253
		}
//line /usr/local/go/src/strings/strings.go:1253
		// _ = "end of CoverTab[4013]"
	}
//line /usr/local/go/src/strings/strings.go:1254
	// _ = "end of CoverTab[3980]"
//line /usr/local/go/src/strings/strings.go:1254
	_go_fuzz_dep_.CoverTab[3981]++
							return -1
//line /usr/local/go/src/strings/strings.go:1255
	// _ = "end of CoverTab[3981]"
}

// Cut slices s around the first instance of sep,
//line /usr/local/go/src/strings/strings.go:1258
// returning the text before and after sep.
//line /usr/local/go/src/strings/strings.go:1258
// The found result reports whether sep appears in s.
//line /usr/local/go/src/strings/strings.go:1258
// If sep does not appear in s, cut returns s, "", false.
//line /usr/local/go/src/strings/strings.go:1262
func Cut(s, sep string) (before, after string, found bool) {
//line /usr/local/go/src/strings/strings.go:1262
	_go_fuzz_dep_.CoverTab[4028]++
							if i := Index(s, sep); i >= 0 {
//line /usr/local/go/src/strings/strings.go:1263
		_go_fuzz_dep_.CoverTab[4030]++
								return s[:i], s[i+len(sep):], true
//line /usr/local/go/src/strings/strings.go:1264
		// _ = "end of CoverTab[4030]"
	} else {
//line /usr/local/go/src/strings/strings.go:1265
		_go_fuzz_dep_.CoverTab[4031]++
//line /usr/local/go/src/strings/strings.go:1265
		// _ = "end of CoverTab[4031]"
//line /usr/local/go/src/strings/strings.go:1265
	}
//line /usr/local/go/src/strings/strings.go:1265
	// _ = "end of CoverTab[4028]"
//line /usr/local/go/src/strings/strings.go:1265
	_go_fuzz_dep_.CoverTab[4029]++
							return s, "", false
//line /usr/local/go/src/strings/strings.go:1266
	// _ = "end of CoverTab[4029]"
}

// CutPrefix returns s without the provided leading prefix string
//line /usr/local/go/src/strings/strings.go:1269
// and reports whether it found the prefix.
//line /usr/local/go/src/strings/strings.go:1269
// If s doesn't start with prefix, CutPrefix returns s, false.
//line /usr/local/go/src/strings/strings.go:1269
// If prefix is the empty string, CutPrefix returns s, true.
//line /usr/local/go/src/strings/strings.go:1273
func CutPrefix(s, prefix string) (after string, found bool) {
//line /usr/local/go/src/strings/strings.go:1273
	_go_fuzz_dep_.CoverTab[4032]++
							if !HasPrefix(s, prefix) {
//line /usr/local/go/src/strings/strings.go:1274
		_go_fuzz_dep_.CoverTab[4034]++
								return s, false
//line /usr/local/go/src/strings/strings.go:1275
		// _ = "end of CoverTab[4034]"
	} else {
//line /usr/local/go/src/strings/strings.go:1276
		_go_fuzz_dep_.CoverTab[4035]++
//line /usr/local/go/src/strings/strings.go:1276
		// _ = "end of CoverTab[4035]"
//line /usr/local/go/src/strings/strings.go:1276
	}
//line /usr/local/go/src/strings/strings.go:1276
	// _ = "end of CoverTab[4032]"
//line /usr/local/go/src/strings/strings.go:1276
	_go_fuzz_dep_.CoverTab[4033]++
							return s[len(prefix):], true
//line /usr/local/go/src/strings/strings.go:1277
	// _ = "end of CoverTab[4033]"
}

// CutSuffix returns s without the provided ending suffix string
//line /usr/local/go/src/strings/strings.go:1280
// and reports whether it found the suffix.
//line /usr/local/go/src/strings/strings.go:1280
// If s doesn't end with suffix, CutSuffix returns s, false.
//line /usr/local/go/src/strings/strings.go:1280
// If suffix is the empty string, CutSuffix returns s, true.
//line /usr/local/go/src/strings/strings.go:1284
func CutSuffix(s, suffix string) (before string, found bool) {
//line /usr/local/go/src/strings/strings.go:1284
	_go_fuzz_dep_.CoverTab[4036]++
							if !HasSuffix(s, suffix) {
//line /usr/local/go/src/strings/strings.go:1285
		_go_fuzz_dep_.CoverTab[4038]++
								return s, false
//line /usr/local/go/src/strings/strings.go:1286
		// _ = "end of CoverTab[4038]"
	} else {
//line /usr/local/go/src/strings/strings.go:1287
		_go_fuzz_dep_.CoverTab[4039]++
//line /usr/local/go/src/strings/strings.go:1287
		// _ = "end of CoverTab[4039]"
//line /usr/local/go/src/strings/strings.go:1287
	}
//line /usr/local/go/src/strings/strings.go:1287
	// _ = "end of CoverTab[4036]"
//line /usr/local/go/src/strings/strings.go:1287
	_go_fuzz_dep_.CoverTab[4037]++
							return s[:len(s)-len(suffix)], true
//line /usr/local/go/src/strings/strings.go:1288
	// _ = "end of CoverTab[4037]"
}

//line /usr/local/go/src/strings/strings.go:1289
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/strings.go:1289
var _ = _go_fuzz_dep_.CoverTab
