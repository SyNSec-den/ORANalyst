// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/strings/strings.go:5
// Package strings implements simple functions to manipulate UTF-8 encoded strings.
//line /snap/go/10455/src/strings/strings.go:5
//
//line /snap/go/10455/src/strings/strings.go:5
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
//line /snap/go/10455/src/strings/strings.go:8
package strings

//line /snap/go/10455/src/strings/strings.go:8
import (
//line /snap/go/10455/src/strings/strings.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/strings/strings.go:8
)
//line /snap/go/10455/src/strings/strings.go:8
import (
//line /snap/go/10455/src/strings/strings.go:8
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/strings/strings.go:8
)

import (
	"internal/bytealg"
	"unicode"
	"unicode/utf8"
)

const maxInt = int(^uint(0) >> 1)

// explode splits s into a slice of UTF-8 strings,
//line /snap/go/10455/src/strings/strings.go:18
// one string per Unicode character up to a maximum of n (n < 0 means no limit).
//line /snap/go/10455/src/strings/strings.go:18
// Invalid UTF-8 bytes are sliced individually.
//line /snap/go/10455/src/strings/strings.go:21
func explode(s string, n int) []string {
//line /snap/go/10455/src/strings/strings.go:21
	_go_fuzz_dep_.CoverTab[1208]++
							l := utf8.RuneCountInString(s)
							if n < 0 || func() bool {
//line /snap/go/10455/src/strings/strings.go:23
		_go_fuzz_dep_.CoverTab[1212]++
//line /snap/go/10455/src/strings/strings.go:23
		return n > l
//line /snap/go/10455/src/strings/strings.go:23
		// _ = "end of CoverTab[1212]"
//line /snap/go/10455/src/strings/strings.go:23
	}() {
//line /snap/go/10455/src/strings/strings.go:23
		_go_fuzz_dep_.CoverTab[525199]++
//line /snap/go/10455/src/strings/strings.go:23
		_go_fuzz_dep_.CoverTab[1213]++
								n = l
//line /snap/go/10455/src/strings/strings.go:24
		// _ = "end of CoverTab[1213]"
	} else {
//line /snap/go/10455/src/strings/strings.go:25
		_go_fuzz_dep_.CoverTab[525200]++
//line /snap/go/10455/src/strings/strings.go:25
		_go_fuzz_dep_.CoverTab[1214]++
//line /snap/go/10455/src/strings/strings.go:25
		// _ = "end of CoverTab[1214]"
//line /snap/go/10455/src/strings/strings.go:25
	}
//line /snap/go/10455/src/strings/strings.go:25
	// _ = "end of CoverTab[1208]"
//line /snap/go/10455/src/strings/strings.go:25
	_go_fuzz_dep_.CoverTab[1209]++
							a := make([]string, n)
//line /snap/go/10455/src/strings/strings.go:26
	_go_fuzz_dep_.CoverTab[786510] = 0
							for i := 0; i < n-1; i++ {
//line /snap/go/10455/src/strings/strings.go:27
		if _go_fuzz_dep_.CoverTab[786510] == 0 {
//line /snap/go/10455/src/strings/strings.go:27
			_go_fuzz_dep_.CoverTab[525495]++
//line /snap/go/10455/src/strings/strings.go:27
		} else {
//line /snap/go/10455/src/strings/strings.go:27
			_go_fuzz_dep_.CoverTab[525496]++
//line /snap/go/10455/src/strings/strings.go:27
		}
//line /snap/go/10455/src/strings/strings.go:27
		_go_fuzz_dep_.CoverTab[786510] = 1
//line /snap/go/10455/src/strings/strings.go:27
		_go_fuzz_dep_.CoverTab[1215]++
								_, size := utf8.DecodeRuneInString(s)
								a[i] = s[:size]
								s = s[size:]
//line /snap/go/10455/src/strings/strings.go:30
		// _ = "end of CoverTab[1215]"
	}
//line /snap/go/10455/src/strings/strings.go:31
	if _go_fuzz_dep_.CoverTab[786510] == 0 {
//line /snap/go/10455/src/strings/strings.go:31
		_go_fuzz_dep_.CoverTab[525497]++
//line /snap/go/10455/src/strings/strings.go:31
	} else {
//line /snap/go/10455/src/strings/strings.go:31
		_go_fuzz_dep_.CoverTab[525498]++
//line /snap/go/10455/src/strings/strings.go:31
	}
//line /snap/go/10455/src/strings/strings.go:31
	// _ = "end of CoverTab[1209]"
//line /snap/go/10455/src/strings/strings.go:31
	_go_fuzz_dep_.CoverTab[1210]++
							if n > 0 {
//line /snap/go/10455/src/strings/strings.go:32
		_go_fuzz_dep_.CoverTab[525201]++
//line /snap/go/10455/src/strings/strings.go:32
		_go_fuzz_dep_.CoverTab[1216]++
								a[n-1] = s
//line /snap/go/10455/src/strings/strings.go:33
		// _ = "end of CoverTab[1216]"
	} else {
//line /snap/go/10455/src/strings/strings.go:34
		_go_fuzz_dep_.CoverTab[525202]++
//line /snap/go/10455/src/strings/strings.go:34
		_go_fuzz_dep_.CoverTab[1217]++
//line /snap/go/10455/src/strings/strings.go:34
		// _ = "end of CoverTab[1217]"
//line /snap/go/10455/src/strings/strings.go:34
	}
//line /snap/go/10455/src/strings/strings.go:34
	// _ = "end of CoverTab[1210]"
//line /snap/go/10455/src/strings/strings.go:34
	_go_fuzz_dep_.CoverTab[1211]++
							return a
//line /snap/go/10455/src/strings/strings.go:35
	// _ = "end of CoverTab[1211]"
}

// Count counts the number of non-overlapping instances of substr in s.
//line /snap/go/10455/src/strings/strings.go:38
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
//line /snap/go/10455/src/strings/strings.go:40
func Count(s, substr string) int {
//line /snap/go/10455/src/strings/strings.go:40
	_go_fuzz_dep_.CoverTab[1218]++

							if len(substr) == 0 {
//line /snap/go/10455/src/strings/strings.go:42
		_go_fuzz_dep_.CoverTab[525203]++
//line /snap/go/10455/src/strings/strings.go:42
		_go_fuzz_dep_.CoverTab[1221]++
								return utf8.RuneCountInString(s) + 1
//line /snap/go/10455/src/strings/strings.go:43
		// _ = "end of CoverTab[1221]"
	} else {
//line /snap/go/10455/src/strings/strings.go:44
		_go_fuzz_dep_.CoverTab[525204]++
//line /snap/go/10455/src/strings/strings.go:44
		_go_fuzz_dep_.CoverTab[1222]++
//line /snap/go/10455/src/strings/strings.go:44
		// _ = "end of CoverTab[1222]"
//line /snap/go/10455/src/strings/strings.go:44
	}
//line /snap/go/10455/src/strings/strings.go:44
	// _ = "end of CoverTab[1218]"
//line /snap/go/10455/src/strings/strings.go:44
	_go_fuzz_dep_.CoverTab[1219]++
							if len(substr) == 1 {
//line /snap/go/10455/src/strings/strings.go:45
		_go_fuzz_dep_.CoverTab[525205]++
//line /snap/go/10455/src/strings/strings.go:45
		_go_fuzz_dep_.CoverTab[1223]++
								return bytealg.CountString(s, substr[0])
//line /snap/go/10455/src/strings/strings.go:46
		// _ = "end of CoverTab[1223]"
	} else {
//line /snap/go/10455/src/strings/strings.go:47
		_go_fuzz_dep_.CoverTab[525206]++
//line /snap/go/10455/src/strings/strings.go:47
		_go_fuzz_dep_.CoverTab[1224]++
//line /snap/go/10455/src/strings/strings.go:47
		// _ = "end of CoverTab[1224]"
//line /snap/go/10455/src/strings/strings.go:47
	}
//line /snap/go/10455/src/strings/strings.go:47
	// _ = "end of CoverTab[1219]"
//line /snap/go/10455/src/strings/strings.go:47
	_go_fuzz_dep_.CoverTab[1220]++
							n := 0
//line /snap/go/10455/src/strings/strings.go:48
	_go_fuzz_dep_.CoverTab[786511] = 0
							for {
//line /snap/go/10455/src/strings/strings.go:49
		if _go_fuzz_dep_.CoverTab[786511] == 0 {
//line /snap/go/10455/src/strings/strings.go:49
			_go_fuzz_dep_.CoverTab[525499]++
//line /snap/go/10455/src/strings/strings.go:49
		} else {
//line /snap/go/10455/src/strings/strings.go:49
			_go_fuzz_dep_.CoverTab[525500]++
//line /snap/go/10455/src/strings/strings.go:49
		}
//line /snap/go/10455/src/strings/strings.go:49
		_go_fuzz_dep_.CoverTab[786511] = 1
//line /snap/go/10455/src/strings/strings.go:49
		_go_fuzz_dep_.CoverTab[1225]++
								i := Index(s, substr)
								if i == -1 {
//line /snap/go/10455/src/strings/strings.go:51
			_go_fuzz_dep_.CoverTab[525207]++
//line /snap/go/10455/src/strings/strings.go:51
			_go_fuzz_dep_.CoverTab[1227]++
									return n
//line /snap/go/10455/src/strings/strings.go:52
			// _ = "end of CoverTab[1227]"
		} else {
//line /snap/go/10455/src/strings/strings.go:53
			_go_fuzz_dep_.CoverTab[525208]++
//line /snap/go/10455/src/strings/strings.go:53
			_go_fuzz_dep_.CoverTab[1228]++
//line /snap/go/10455/src/strings/strings.go:53
			// _ = "end of CoverTab[1228]"
//line /snap/go/10455/src/strings/strings.go:53
		}
//line /snap/go/10455/src/strings/strings.go:53
		// _ = "end of CoverTab[1225]"
//line /snap/go/10455/src/strings/strings.go:53
		_go_fuzz_dep_.CoverTab[1226]++
								n++
								s = s[i+len(substr):]
//line /snap/go/10455/src/strings/strings.go:55
		// _ = "end of CoverTab[1226]"
	}
//line /snap/go/10455/src/strings/strings.go:56
	// _ = "end of CoverTab[1220]"
}

// Contains reports whether substr is within s.
func Contains(s, substr string) bool {
//line /snap/go/10455/src/strings/strings.go:60
	_go_fuzz_dep_.CoverTab[1229]++
							return Index(s, substr) >= 0
//line /snap/go/10455/src/strings/strings.go:61
	// _ = "end of CoverTab[1229]"
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool {
//line /snap/go/10455/src/strings/strings.go:65
	_go_fuzz_dep_.CoverTab[1230]++
							return IndexAny(s, chars) >= 0
//line /snap/go/10455/src/strings/strings.go:66
	// _ = "end of CoverTab[1230]"
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool {
//line /snap/go/10455/src/strings/strings.go:70
	_go_fuzz_dep_.CoverTab[1231]++
							return IndexRune(s, r) >= 0
//line /snap/go/10455/src/strings/strings.go:71
	// _ = "end of CoverTab[1231]"
}

// ContainsFunc reports whether any Unicode code points r within s satisfy f(r).
func ContainsFunc(s string, f func(rune) bool) bool {
//line /snap/go/10455/src/strings/strings.go:75
	_go_fuzz_dep_.CoverTab[1232]++
							return IndexFunc(s, f) >= 0
//line /snap/go/10455/src/strings/strings.go:76
	// _ = "end of CoverTab[1232]"
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int {
//line /snap/go/10455/src/strings/strings.go:80
	_go_fuzz_dep_.CoverTab[1233]++
							n := len(substr)
							switch {
	case n == 0:
//line /snap/go/10455/src/strings/strings.go:83
		_go_fuzz_dep_.CoverTab[525209]++
//line /snap/go/10455/src/strings/strings.go:83
		_go_fuzz_dep_.CoverTab[1238]++
								return len(s)
//line /snap/go/10455/src/strings/strings.go:84
		// _ = "end of CoverTab[1238]"
	case n == 1:
//line /snap/go/10455/src/strings/strings.go:85
		_go_fuzz_dep_.CoverTab[525210]++
//line /snap/go/10455/src/strings/strings.go:85
		_go_fuzz_dep_.CoverTab[1239]++
								return LastIndexByte(s, substr[0])
//line /snap/go/10455/src/strings/strings.go:86
		// _ = "end of CoverTab[1239]"
	case n == len(s):
//line /snap/go/10455/src/strings/strings.go:87
		_go_fuzz_dep_.CoverTab[525211]++
//line /snap/go/10455/src/strings/strings.go:87
		_go_fuzz_dep_.CoverTab[1240]++
								if substr == s {
//line /snap/go/10455/src/strings/strings.go:88
			_go_fuzz_dep_.CoverTab[525214]++
//line /snap/go/10455/src/strings/strings.go:88
			_go_fuzz_dep_.CoverTab[1244]++
									return 0
//line /snap/go/10455/src/strings/strings.go:89
			// _ = "end of CoverTab[1244]"
		} else {
//line /snap/go/10455/src/strings/strings.go:90
			_go_fuzz_dep_.CoverTab[525215]++
//line /snap/go/10455/src/strings/strings.go:90
			_go_fuzz_dep_.CoverTab[1245]++
//line /snap/go/10455/src/strings/strings.go:90
			// _ = "end of CoverTab[1245]"
//line /snap/go/10455/src/strings/strings.go:90
		}
//line /snap/go/10455/src/strings/strings.go:90
		// _ = "end of CoverTab[1240]"
//line /snap/go/10455/src/strings/strings.go:90
		_go_fuzz_dep_.CoverTab[1241]++
								return -1
//line /snap/go/10455/src/strings/strings.go:91
		// _ = "end of CoverTab[1241]"
	case n > len(s):
//line /snap/go/10455/src/strings/strings.go:92
		_go_fuzz_dep_.CoverTab[525212]++
//line /snap/go/10455/src/strings/strings.go:92
		_go_fuzz_dep_.CoverTab[1242]++
								return -1
//line /snap/go/10455/src/strings/strings.go:93
		// _ = "end of CoverTab[1242]"
//line /snap/go/10455/src/strings/strings.go:93
	default:
//line /snap/go/10455/src/strings/strings.go:93
		_go_fuzz_dep_.CoverTab[525213]++
//line /snap/go/10455/src/strings/strings.go:93
		_go_fuzz_dep_.CoverTab[1243]++
//line /snap/go/10455/src/strings/strings.go:93
		// _ = "end of CoverTab[1243]"
	}
//line /snap/go/10455/src/strings/strings.go:94
	// _ = "end of CoverTab[1233]"
//line /snap/go/10455/src/strings/strings.go:94
	_go_fuzz_dep_.CoverTab[1234]++

							hashss, pow := bytealg.HashStrRev(substr)
							last := len(s) - n
							var h uint32
//line /snap/go/10455/src/strings/strings.go:98
	_go_fuzz_dep_.CoverTab[786512] = 0
							for i := len(s) - 1; i >= last; i-- {
//line /snap/go/10455/src/strings/strings.go:99
		if _go_fuzz_dep_.CoverTab[786512] == 0 {
//line /snap/go/10455/src/strings/strings.go:99
			_go_fuzz_dep_.CoverTab[525503]++
//line /snap/go/10455/src/strings/strings.go:99
		} else {
//line /snap/go/10455/src/strings/strings.go:99
			_go_fuzz_dep_.CoverTab[525504]++
//line /snap/go/10455/src/strings/strings.go:99
		}
//line /snap/go/10455/src/strings/strings.go:99
		_go_fuzz_dep_.CoverTab[786512] = 1
//line /snap/go/10455/src/strings/strings.go:99
		_go_fuzz_dep_.CoverTab[1246]++
								h = h*bytealg.PrimeRK + uint32(s[i])
//line /snap/go/10455/src/strings/strings.go:100
		// _ = "end of CoverTab[1246]"
	}
//line /snap/go/10455/src/strings/strings.go:101
	if _go_fuzz_dep_.CoverTab[786512] == 0 {
//line /snap/go/10455/src/strings/strings.go:101
		_go_fuzz_dep_.CoverTab[525505]++
//line /snap/go/10455/src/strings/strings.go:101
	} else {
//line /snap/go/10455/src/strings/strings.go:101
		_go_fuzz_dep_.CoverTab[525506]++
//line /snap/go/10455/src/strings/strings.go:101
	}
//line /snap/go/10455/src/strings/strings.go:101
	// _ = "end of CoverTab[1234]"
//line /snap/go/10455/src/strings/strings.go:101
	_go_fuzz_dep_.CoverTab[1235]++
							if h == hashss && func() bool {
//line /snap/go/10455/src/strings/strings.go:102
		_go_fuzz_dep_.CoverTab[1247]++
//line /snap/go/10455/src/strings/strings.go:102
		return s[last:] == substr
//line /snap/go/10455/src/strings/strings.go:102
		// _ = "end of CoverTab[1247]"
//line /snap/go/10455/src/strings/strings.go:102
	}() {
//line /snap/go/10455/src/strings/strings.go:102
		_go_fuzz_dep_.CoverTab[525216]++
//line /snap/go/10455/src/strings/strings.go:102
		_go_fuzz_dep_.CoverTab[1248]++
								return last
//line /snap/go/10455/src/strings/strings.go:103
		// _ = "end of CoverTab[1248]"
	} else {
//line /snap/go/10455/src/strings/strings.go:104
		_go_fuzz_dep_.CoverTab[525217]++
//line /snap/go/10455/src/strings/strings.go:104
		_go_fuzz_dep_.CoverTab[1249]++
//line /snap/go/10455/src/strings/strings.go:104
		// _ = "end of CoverTab[1249]"
//line /snap/go/10455/src/strings/strings.go:104
	}
//line /snap/go/10455/src/strings/strings.go:104
	// _ = "end of CoverTab[1235]"
//line /snap/go/10455/src/strings/strings.go:104
	_go_fuzz_dep_.CoverTab[1236]++
//line /snap/go/10455/src/strings/strings.go:104
	_go_fuzz_dep_.CoverTab[786513] = 0
							for i := last - 1; i >= 0; i-- {
//line /snap/go/10455/src/strings/strings.go:105
		if _go_fuzz_dep_.CoverTab[786513] == 0 {
//line /snap/go/10455/src/strings/strings.go:105
			_go_fuzz_dep_.CoverTab[525507]++
//line /snap/go/10455/src/strings/strings.go:105
		} else {
//line /snap/go/10455/src/strings/strings.go:105
			_go_fuzz_dep_.CoverTab[525508]++
//line /snap/go/10455/src/strings/strings.go:105
		}
//line /snap/go/10455/src/strings/strings.go:105
		_go_fuzz_dep_.CoverTab[786513] = 1
//line /snap/go/10455/src/strings/strings.go:105
		_go_fuzz_dep_.CoverTab[1250]++
								h *= bytealg.PrimeRK
								h += uint32(s[i])
								h -= pow * uint32(s[i+n])
								if h == hashss && func() bool {
//line /snap/go/10455/src/strings/strings.go:109
			_go_fuzz_dep_.CoverTab[1251]++
//line /snap/go/10455/src/strings/strings.go:109
			return s[i:i+n] == substr
//line /snap/go/10455/src/strings/strings.go:109
			// _ = "end of CoverTab[1251]"
//line /snap/go/10455/src/strings/strings.go:109
		}() {
//line /snap/go/10455/src/strings/strings.go:109
			_go_fuzz_dep_.CoverTab[525218]++
//line /snap/go/10455/src/strings/strings.go:109
			_go_fuzz_dep_.CoverTab[1252]++
									return i
//line /snap/go/10455/src/strings/strings.go:110
			// _ = "end of CoverTab[1252]"
		} else {
//line /snap/go/10455/src/strings/strings.go:111
			_go_fuzz_dep_.CoverTab[525219]++
//line /snap/go/10455/src/strings/strings.go:111
			_go_fuzz_dep_.CoverTab[1253]++
//line /snap/go/10455/src/strings/strings.go:111
			// _ = "end of CoverTab[1253]"
//line /snap/go/10455/src/strings/strings.go:111
		}
//line /snap/go/10455/src/strings/strings.go:111
		// _ = "end of CoverTab[1250]"
	}
//line /snap/go/10455/src/strings/strings.go:112
	if _go_fuzz_dep_.CoverTab[786513] == 0 {
//line /snap/go/10455/src/strings/strings.go:112
		_go_fuzz_dep_.CoverTab[525509]++
//line /snap/go/10455/src/strings/strings.go:112
	} else {
//line /snap/go/10455/src/strings/strings.go:112
		_go_fuzz_dep_.CoverTab[525510]++
//line /snap/go/10455/src/strings/strings.go:112
	}
//line /snap/go/10455/src/strings/strings.go:112
	// _ = "end of CoverTab[1236]"
//line /snap/go/10455/src/strings/strings.go:112
	_go_fuzz_dep_.CoverTab[1237]++
							return -1
//line /snap/go/10455/src/strings/strings.go:113
	// _ = "end of CoverTab[1237]"
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s string, c byte) int {
//line /snap/go/10455/src/strings/strings.go:117
	_go_fuzz_dep_.CoverTab[1254]++
							return bytealg.IndexByteString(s, c)
//line /snap/go/10455/src/strings/strings.go:118
	// _ = "end of CoverTab[1254]"
}

// IndexRune returns the index of the first instance of the Unicode code point
//line /snap/go/10455/src/strings/strings.go:121
// r, or -1 if rune is not present in s.
//line /snap/go/10455/src/strings/strings.go:121
// If r is utf8.RuneError, it returns the first instance of any
//line /snap/go/10455/src/strings/strings.go:121
// invalid UTF-8 byte sequence.
//line /snap/go/10455/src/strings/strings.go:125
func IndexRune(s string, r rune) int {
//line /snap/go/10455/src/strings/strings.go:125
	_go_fuzz_dep_.CoverTab[1255]++
							switch {
	case 0 <= r && func() bool {
//line /snap/go/10455/src/strings/strings.go:127
		_go_fuzz_dep_.CoverTab[1261]++
//line /snap/go/10455/src/strings/strings.go:127
		return r < utf8.RuneSelf
//line /snap/go/10455/src/strings/strings.go:127
		// _ = "end of CoverTab[1261]"
//line /snap/go/10455/src/strings/strings.go:127
	}():
//line /snap/go/10455/src/strings/strings.go:127
		_go_fuzz_dep_.CoverTab[525220]++
//line /snap/go/10455/src/strings/strings.go:127
		_go_fuzz_dep_.CoverTab[1256]++
								return IndexByte(s, byte(r))
//line /snap/go/10455/src/strings/strings.go:128
		// _ = "end of CoverTab[1256]"
	case r == utf8.RuneError:
//line /snap/go/10455/src/strings/strings.go:129
		_go_fuzz_dep_.CoverTab[525221]++
//line /snap/go/10455/src/strings/strings.go:129
		_go_fuzz_dep_.CoverTab[1257]++
								for i, r := range s {
//line /snap/go/10455/src/strings/strings.go:130
			_go_fuzz_dep_.CoverTab[1262]++
									if r == utf8.RuneError {
//line /snap/go/10455/src/strings/strings.go:131
				_go_fuzz_dep_.CoverTab[525224]++
//line /snap/go/10455/src/strings/strings.go:131
				_go_fuzz_dep_.CoverTab[1263]++
										return i
//line /snap/go/10455/src/strings/strings.go:132
				// _ = "end of CoverTab[1263]"
			} else {
//line /snap/go/10455/src/strings/strings.go:133
				_go_fuzz_dep_.CoverTab[525225]++
//line /snap/go/10455/src/strings/strings.go:133
				_go_fuzz_dep_.CoverTab[1264]++
//line /snap/go/10455/src/strings/strings.go:133
				// _ = "end of CoverTab[1264]"
//line /snap/go/10455/src/strings/strings.go:133
			}
//line /snap/go/10455/src/strings/strings.go:133
			// _ = "end of CoverTab[1262]"
		}
//line /snap/go/10455/src/strings/strings.go:134
		// _ = "end of CoverTab[1257]"
//line /snap/go/10455/src/strings/strings.go:134
		_go_fuzz_dep_.CoverTab[1258]++
								return -1
//line /snap/go/10455/src/strings/strings.go:135
		// _ = "end of CoverTab[1258]"
	case !utf8.ValidRune(r):
//line /snap/go/10455/src/strings/strings.go:136
		_go_fuzz_dep_.CoverTab[525222]++
//line /snap/go/10455/src/strings/strings.go:136
		_go_fuzz_dep_.CoverTab[1259]++
								return -1
//line /snap/go/10455/src/strings/strings.go:137
		// _ = "end of CoverTab[1259]"
	default:
//line /snap/go/10455/src/strings/strings.go:138
		_go_fuzz_dep_.CoverTab[525223]++
//line /snap/go/10455/src/strings/strings.go:138
		_go_fuzz_dep_.CoverTab[1260]++
								return Index(s, string(r))
//line /snap/go/10455/src/strings/strings.go:139
		// _ = "end of CoverTab[1260]"
	}
//line /snap/go/10455/src/strings/strings.go:140
	// _ = "end of CoverTab[1255]"
}

// IndexAny returns the index of the first instance of any Unicode code point
//line /snap/go/10455/src/strings/strings.go:143
// from chars in s, or -1 if no Unicode code point from chars is present in s.
//line /snap/go/10455/src/strings/strings.go:145
func IndexAny(s, chars string) int {
//line /snap/go/10455/src/strings/strings.go:145
	_go_fuzz_dep_.CoverTab[1265]++
							if chars == "" {
//line /snap/go/10455/src/strings/strings.go:146
		_go_fuzz_dep_.CoverTab[525226]++
//line /snap/go/10455/src/strings/strings.go:146
		_go_fuzz_dep_.CoverTab[1270]++

								return -1
//line /snap/go/10455/src/strings/strings.go:148
		// _ = "end of CoverTab[1270]"
	} else {
//line /snap/go/10455/src/strings/strings.go:149
		_go_fuzz_dep_.CoverTab[525227]++
//line /snap/go/10455/src/strings/strings.go:149
		_go_fuzz_dep_.CoverTab[1271]++
//line /snap/go/10455/src/strings/strings.go:149
		// _ = "end of CoverTab[1271]"
//line /snap/go/10455/src/strings/strings.go:149
	}
//line /snap/go/10455/src/strings/strings.go:149
	// _ = "end of CoverTab[1265]"
//line /snap/go/10455/src/strings/strings.go:149
	_go_fuzz_dep_.CoverTab[1266]++
							if len(chars) == 1 {
//line /snap/go/10455/src/strings/strings.go:150
		_go_fuzz_dep_.CoverTab[525228]++
//line /snap/go/10455/src/strings/strings.go:150
		_go_fuzz_dep_.CoverTab[1272]++

								r := rune(chars[0])
								if r >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:153
			_go_fuzz_dep_.CoverTab[525230]++
//line /snap/go/10455/src/strings/strings.go:153
			_go_fuzz_dep_.CoverTab[1274]++
									r = utf8.RuneError
//line /snap/go/10455/src/strings/strings.go:154
			// _ = "end of CoverTab[1274]"
		} else {
//line /snap/go/10455/src/strings/strings.go:155
			_go_fuzz_dep_.CoverTab[525231]++
//line /snap/go/10455/src/strings/strings.go:155
			_go_fuzz_dep_.CoverTab[1275]++
//line /snap/go/10455/src/strings/strings.go:155
			// _ = "end of CoverTab[1275]"
//line /snap/go/10455/src/strings/strings.go:155
		}
//line /snap/go/10455/src/strings/strings.go:155
		// _ = "end of CoverTab[1272]"
//line /snap/go/10455/src/strings/strings.go:155
		_go_fuzz_dep_.CoverTab[1273]++
								return IndexRune(s, r)
//line /snap/go/10455/src/strings/strings.go:156
		// _ = "end of CoverTab[1273]"
	} else {
//line /snap/go/10455/src/strings/strings.go:157
		_go_fuzz_dep_.CoverTab[525229]++
//line /snap/go/10455/src/strings/strings.go:157
		_go_fuzz_dep_.CoverTab[1276]++
//line /snap/go/10455/src/strings/strings.go:157
		// _ = "end of CoverTab[1276]"
//line /snap/go/10455/src/strings/strings.go:157
	}
//line /snap/go/10455/src/strings/strings.go:157
	// _ = "end of CoverTab[1266]"
//line /snap/go/10455/src/strings/strings.go:157
	_go_fuzz_dep_.CoverTab[1267]++
							if len(s) > 8 {
//line /snap/go/10455/src/strings/strings.go:158
		_go_fuzz_dep_.CoverTab[525232]++
//line /snap/go/10455/src/strings/strings.go:158
		_go_fuzz_dep_.CoverTab[1277]++
								if as, isASCII := makeASCIISet(chars); isASCII {
//line /snap/go/10455/src/strings/strings.go:159
			_go_fuzz_dep_.CoverTab[525234]++
//line /snap/go/10455/src/strings/strings.go:159
			_go_fuzz_dep_.CoverTab[1278]++
//line /snap/go/10455/src/strings/strings.go:159
			_go_fuzz_dep_.CoverTab[786515] = 0
									for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/strings.go:160
				if _go_fuzz_dep_.CoverTab[786515] == 0 {
//line /snap/go/10455/src/strings/strings.go:160
					_go_fuzz_dep_.CoverTab[525515]++
//line /snap/go/10455/src/strings/strings.go:160
				} else {
//line /snap/go/10455/src/strings/strings.go:160
					_go_fuzz_dep_.CoverTab[525516]++
//line /snap/go/10455/src/strings/strings.go:160
				}
//line /snap/go/10455/src/strings/strings.go:160
				_go_fuzz_dep_.CoverTab[786515] = 1
//line /snap/go/10455/src/strings/strings.go:160
				_go_fuzz_dep_.CoverTab[1280]++
										if as.contains(s[i]) {
//line /snap/go/10455/src/strings/strings.go:161
					_go_fuzz_dep_.CoverTab[525236]++
//line /snap/go/10455/src/strings/strings.go:161
					_go_fuzz_dep_.CoverTab[1281]++
											return i
//line /snap/go/10455/src/strings/strings.go:162
					// _ = "end of CoverTab[1281]"
				} else {
//line /snap/go/10455/src/strings/strings.go:163
					_go_fuzz_dep_.CoverTab[525237]++
//line /snap/go/10455/src/strings/strings.go:163
					_go_fuzz_dep_.CoverTab[1282]++
//line /snap/go/10455/src/strings/strings.go:163
					// _ = "end of CoverTab[1282]"
//line /snap/go/10455/src/strings/strings.go:163
				}
//line /snap/go/10455/src/strings/strings.go:163
				// _ = "end of CoverTab[1280]"
			}
//line /snap/go/10455/src/strings/strings.go:164
			if _go_fuzz_dep_.CoverTab[786515] == 0 {
//line /snap/go/10455/src/strings/strings.go:164
				_go_fuzz_dep_.CoverTab[525517]++
//line /snap/go/10455/src/strings/strings.go:164
			} else {
//line /snap/go/10455/src/strings/strings.go:164
				_go_fuzz_dep_.CoverTab[525518]++
//line /snap/go/10455/src/strings/strings.go:164
			}
//line /snap/go/10455/src/strings/strings.go:164
			// _ = "end of CoverTab[1278]"
//line /snap/go/10455/src/strings/strings.go:164
			_go_fuzz_dep_.CoverTab[1279]++
									return -1
//line /snap/go/10455/src/strings/strings.go:165
			// _ = "end of CoverTab[1279]"
		} else {
//line /snap/go/10455/src/strings/strings.go:166
			_go_fuzz_dep_.CoverTab[525235]++
//line /snap/go/10455/src/strings/strings.go:166
			_go_fuzz_dep_.CoverTab[1283]++
//line /snap/go/10455/src/strings/strings.go:166
			// _ = "end of CoverTab[1283]"
//line /snap/go/10455/src/strings/strings.go:166
		}
//line /snap/go/10455/src/strings/strings.go:166
		// _ = "end of CoverTab[1277]"
	} else {
//line /snap/go/10455/src/strings/strings.go:167
		_go_fuzz_dep_.CoverTab[525233]++
//line /snap/go/10455/src/strings/strings.go:167
		_go_fuzz_dep_.CoverTab[1284]++
//line /snap/go/10455/src/strings/strings.go:167
		// _ = "end of CoverTab[1284]"
//line /snap/go/10455/src/strings/strings.go:167
	}
//line /snap/go/10455/src/strings/strings.go:167
	// _ = "end of CoverTab[1267]"
//line /snap/go/10455/src/strings/strings.go:167
	_go_fuzz_dep_.CoverTab[1268]++
//line /snap/go/10455/src/strings/strings.go:167
	_go_fuzz_dep_.CoverTab[786514] = 0
							for i, c := range s {
//line /snap/go/10455/src/strings/strings.go:168
		if _go_fuzz_dep_.CoverTab[786514] == 0 {
//line /snap/go/10455/src/strings/strings.go:168
			_go_fuzz_dep_.CoverTab[525511]++
//line /snap/go/10455/src/strings/strings.go:168
		} else {
//line /snap/go/10455/src/strings/strings.go:168
			_go_fuzz_dep_.CoverTab[525512]++
//line /snap/go/10455/src/strings/strings.go:168
		}
//line /snap/go/10455/src/strings/strings.go:168
		_go_fuzz_dep_.CoverTab[786514] = 1
//line /snap/go/10455/src/strings/strings.go:168
		_go_fuzz_dep_.CoverTab[1285]++
								if IndexRune(chars, c) >= 0 {
//line /snap/go/10455/src/strings/strings.go:169
			_go_fuzz_dep_.CoverTab[525238]++
//line /snap/go/10455/src/strings/strings.go:169
			_go_fuzz_dep_.CoverTab[1286]++
									return i
//line /snap/go/10455/src/strings/strings.go:170
			// _ = "end of CoverTab[1286]"
		} else {
//line /snap/go/10455/src/strings/strings.go:171
			_go_fuzz_dep_.CoverTab[525239]++
//line /snap/go/10455/src/strings/strings.go:171
			_go_fuzz_dep_.CoverTab[1287]++
//line /snap/go/10455/src/strings/strings.go:171
			// _ = "end of CoverTab[1287]"
//line /snap/go/10455/src/strings/strings.go:171
		}
//line /snap/go/10455/src/strings/strings.go:171
		// _ = "end of CoverTab[1285]"
	}
//line /snap/go/10455/src/strings/strings.go:172
	if _go_fuzz_dep_.CoverTab[786514] == 0 {
//line /snap/go/10455/src/strings/strings.go:172
		_go_fuzz_dep_.CoverTab[525513]++
//line /snap/go/10455/src/strings/strings.go:172
	} else {
//line /snap/go/10455/src/strings/strings.go:172
		_go_fuzz_dep_.CoverTab[525514]++
//line /snap/go/10455/src/strings/strings.go:172
	}
//line /snap/go/10455/src/strings/strings.go:172
	// _ = "end of CoverTab[1268]"
//line /snap/go/10455/src/strings/strings.go:172
	_go_fuzz_dep_.CoverTab[1269]++
							return -1
//line /snap/go/10455/src/strings/strings.go:173
	// _ = "end of CoverTab[1269]"
}

// LastIndexAny returns the index of the last instance of any Unicode code
//line /snap/go/10455/src/strings/strings.go:176
// point from chars in s, or -1 if no Unicode code point from chars is
//line /snap/go/10455/src/strings/strings.go:176
// present in s.
//line /snap/go/10455/src/strings/strings.go:179
func LastIndexAny(s, chars string) int {
//line /snap/go/10455/src/strings/strings.go:179
	_go_fuzz_dep_.CoverTab[1288]++
							if chars == "" {
//line /snap/go/10455/src/strings/strings.go:180
		_go_fuzz_dep_.CoverTab[525240]++
//line /snap/go/10455/src/strings/strings.go:180
		_go_fuzz_dep_.CoverTab[1294]++

								return -1
//line /snap/go/10455/src/strings/strings.go:182
		// _ = "end of CoverTab[1294]"
	} else {
//line /snap/go/10455/src/strings/strings.go:183
		_go_fuzz_dep_.CoverTab[525241]++
//line /snap/go/10455/src/strings/strings.go:183
		_go_fuzz_dep_.CoverTab[1295]++
//line /snap/go/10455/src/strings/strings.go:183
		// _ = "end of CoverTab[1295]"
//line /snap/go/10455/src/strings/strings.go:183
	}
//line /snap/go/10455/src/strings/strings.go:183
	// _ = "end of CoverTab[1288]"
//line /snap/go/10455/src/strings/strings.go:183
	_go_fuzz_dep_.CoverTab[1289]++
							if len(s) == 1 {
//line /snap/go/10455/src/strings/strings.go:184
		_go_fuzz_dep_.CoverTab[525242]++
//line /snap/go/10455/src/strings/strings.go:184
		_go_fuzz_dep_.CoverTab[1296]++
								rc := rune(s[0])
								if rc >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:186
			_go_fuzz_dep_.CoverTab[525244]++
//line /snap/go/10455/src/strings/strings.go:186
			_go_fuzz_dep_.CoverTab[1299]++
									rc = utf8.RuneError
//line /snap/go/10455/src/strings/strings.go:187
			// _ = "end of CoverTab[1299]"
		} else {
//line /snap/go/10455/src/strings/strings.go:188
			_go_fuzz_dep_.CoverTab[525245]++
//line /snap/go/10455/src/strings/strings.go:188
			_go_fuzz_dep_.CoverTab[1300]++
//line /snap/go/10455/src/strings/strings.go:188
			// _ = "end of CoverTab[1300]"
//line /snap/go/10455/src/strings/strings.go:188
		}
//line /snap/go/10455/src/strings/strings.go:188
		// _ = "end of CoverTab[1296]"
//line /snap/go/10455/src/strings/strings.go:188
		_go_fuzz_dep_.CoverTab[1297]++
								if IndexRune(chars, rc) >= 0 {
//line /snap/go/10455/src/strings/strings.go:189
			_go_fuzz_dep_.CoverTab[525246]++
//line /snap/go/10455/src/strings/strings.go:189
			_go_fuzz_dep_.CoverTab[1301]++
									return 0
//line /snap/go/10455/src/strings/strings.go:190
			// _ = "end of CoverTab[1301]"
		} else {
//line /snap/go/10455/src/strings/strings.go:191
			_go_fuzz_dep_.CoverTab[525247]++
//line /snap/go/10455/src/strings/strings.go:191
			_go_fuzz_dep_.CoverTab[1302]++
//line /snap/go/10455/src/strings/strings.go:191
			// _ = "end of CoverTab[1302]"
//line /snap/go/10455/src/strings/strings.go:191
		}
//line /snap/go/10455/src/strings/strings.go:191
		// _ = "end of CoverTab[1297]"
//line /snap/go/10455/src/strings/strings.go:191
		_go_fuzz_dep_.CoverTab[1298]++
								return -1
//line /snap/go/10455/src/strings/strings.go:192
		// _ = "end of CoverTab[1298]"
	} else {
//line /snap/go/10455/src/strings/strings.go:193
		_go_fuzz_dep_.CoverTab[525243]++
//line /snap/go/10455/src/strings/strings.go:193
		_go_fuzz_dep_.CoverTab[1303]++
//line /snap/go/10455/src/strings/strings.go:193
		// _ = "end of CoverTab[1303]"
//line /snap/go/10455/src/strings/strings.go:193
	}
//line /snap/go/10455/src/strings/strings.go:193
	// _ = "end of CoverTab[1289]"
//line /snap/go/10455/src/strings/strings.go:193
	_go_fuzz_dep_.CoverTab[1290]++
							if len(s) > 8 {
//line /snap/go/10455/src/strings/strings.go:194
		_go_fuzz_dep_.CoverTab[525248]++
//line /snap/go/10455/src/strings/strings.go:194
		_go_fuzz_dep_.CoverTab[1304]++
								if as, isASCII := makeASCIISet(chars); isASCII {
//line /snap/go/10455/src/strings/strings.go:195
			_go_fuzz_dep_.CoverTab[525250]++
//line /snap/go/10455/src/strings/strings.go:195
			_go_fuzz_dep_.CoverTab[1305]++
//line /snap/go/10455/src/strings/strings.go:195
			_go_fuzz_dep_.CoverTab[786517] = 0
									for i := len(s) - 1; i >= 0; i-- {
//line /snap/go/10455/src/strings/strings.go:196
				if _go_fuzz_dep_.CoverTab[786517] == 0 {
//line /snap/go/10455/src/strings/strings.go:196
					_go_fuzz_dep_.CoverTab[525523]++
//line /snap/go/10455/src/strings/strings.go:196
				} else {
//line /snap/go/10455/src/strings/strings.go:196
					_go_fuzz_dep_.CoverTab[525524]++
//line /snap/go/10455/src/strings/strings.go:196
				}
//line /snap/go/10455/src/strings/strings.go:196
				_go_fuzz_dep_.CoverTab[786517] = 1
//line /snap/go/10455/src/strings/strings.go:196
				_go_fuzz_dep_.CoverTab[1307]++
										if as.contains(s[i]) {
//line /snap/go/10455/src/strings/strings.go:197
					_go_fuzz_dep_.CoverTab[525252]++
//line /snap/go/10455/src/strings/strings.go:197
					_go_fuzz_dep_.CoverTab[1308]++
											return i
//line /snap/go/10455/src/strings/strings.go:198
					// _ = "end of CoverTab[1308]"
				} else {
//line /snap/go/10455/src/strings/strings.go:199
					_go_fuzz_dep_.CoverTab[525253]++
//line /snap/go/10455/src/strings/strings.go:199
					_go_fuzz_dep_.CoverTab[1309]++
//line /snap/go/10455/src/strings/strings.go:199
					// _ = "end of CoverTab[1309]"
//line /snap/go/10455/src/strings/strings.go:199
				}
//line /snap/go/10455/src/strings/strings.go:199
				// _ = "end of CoverTab[1307]"
			}
//line /snap/go/10455/src/strings/strings.go:200
			if _go_fuzz_dep_.CoverTab[786517] == 0 {
//line /snap/go/10455/src/strings/strings.go:200
				_go_fuzz_dep_.CoverTab[525525]++
//line /snap/go/10455/src/strings/strings.go:200
			} else {
//line /snap/go/10455/src/strings/strings.go:200
				_go_fuzz_dep_.CoverTab[525526]++
//line /snap/go/10455/src/strings/strings.go:200
			}
//line /snap/go/10455/src/strings/strings.go:200
			// _ = "end of CoverTab[1305]"
//line /snap/go/10455/src/strings/strings.go:200
			_go_fuzz_dep_.CoverTab[1306]++
									return -1
//line /snap/go/10455/src/strings/strings.go:201
			// _ = "end of CoverTab[1306]"
		} else {
//line /snap/go/10455/src/strings/strings.go:202
			_go_fuzz_dep_.CoverTab[525251]++
//line /snap/go/10455/src/strings/strings.go:202
			_go_fuzz_dep_.CoverTab[1310]++
//line /snap/go/10455/src/strings/strings.go:202
			// _ = "end of CoverTab[1310]"
//line /snap/go/10455/src/strings/strings.go:202
		}
//line /snap/go/10455/src/strings/strings.go:202
		// _ = "end of CoverTab[1304]"
	} else {
//line /snap/go/10455/src/strings/strings.go:203
		_go_fuzz_dep_.CoverTab[525249]++
//line /snap/go/10455/src/strings/strings.go:203
		_go_fuzz_dep_.CoverTab[1311]++
//line /snap/go/10455/src/strings/strings.go:203
		// _ = "end of CoverTab[1311]"
//line /snap/go/10455/src/strings/strings.go:203
	}
//line /snap/go/10455/src/strings/strings.go:203
	// _ = "end of CoverTab[1290]"
//line /snap/go/10455/src/strings/strings.go:203
	_go_fuzz_dep_.CoverTab[1291]++
							if len(chars) == 1 {
//line /snap/go/10455/src/strings/strings.go:204
		_go_fuzz_dep_.CoverTab[525254]++
//line /snap/go/10455/src/strings/strings.go:204
		_go_fuzz_dep_.CoverTab[1312]++
								rc := rune(chars[0])
								if rc >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:206
			_go_fuzz_dep_.CoverTab[525256]++
//line /snap/go/10455/src/strings/strings.go:206
			_go_fuzz_dep_.CoverTab[1315]++
									rc = utf8.RuneError
//line /snap/go/10455/src/strings/strings.go:207
			// _ = "end of CoverTab[1315]"
		} else {
//line /snap/go/10455/src/strings/strings.go:208
			_go_fuzz_dep_.CoverTab[525257]++
//line /snap/go/10455/src/strings/strings.go:208
			_go_fuzz_dep_.CoverTab[1316]++
//line /snap/go/10455/src/strings/strings.go:208
			// _ = "end of CoverTab[1316]"
//line /snap/go/10455/src/strings/strings.go:208
		}
//line /snap/go/10455/src/strings/strings.go:208
		// _ = "end of CoverTab[1312]"
//line /snap/go/10455/src/strings/strings.go:208
		_go_fuzz_dep_.CoverTab[1313]++
//line /snap/go/10455/src/strings/strings.go:208
		_go_fuzz_dep_.CoverTab[786518] = 0
								for i := len(s); i > 0; {
//line /snap/go/10455/src/strings/strings.go:209
			if _go_fuzz_dep_.CoverTab[786518] == 0 {
//line /snap/go/10455/src/strings/strings.go:209
				_go_fuzz_dep_.CoverTab[525527]++
//line /snap/go/10455/src/strings/strings.go:209
			} else {
//line /snap/go/10455/src/strings/strings.go:209
				_go_fuzz_dep_.CoverTab[525528]++
//line /snap/go/10455/src/strings/strings.go:209
			}
//line /snap/go/10455/src/strings/strings.go:209
			_go_fuzz_dep_.CoverTab[786518] = 1
//line /snap/go/10455/src/strings/strings.go:209
			_go_fuzz_dep_.CoverTab[1317]++
									r, size := utf8.DecodeLastRuneInString(s[:i])
									i -= size
									if rc == r {
//line /snap/go/10455/src/strings/strings.go:212
				_go_fuzz_dep_.CoverTab[525258]++
//line /snap/go/10455/src/strings/strings.go:212
				_go_fuzz_dep_.CoverTab[1318]++
										return i
//line /snap/go/10455/src/strings/strings.go:213
				// _ = "end of CoverTab[1318]"
			} else {
//line /snap/go/10455/src/strings/strings.go:214
				_go_fuzz_dep_.CoverTab[525259]++
//line /snap/go/10455/src/strings/strings.go:214
				_go_fuzz_dep_.CoverTab[1319]++
//line /snap/go/10455/src/strings/strings.go:214
				// _ = "end of CoverTab[1319]"
//line /snap/go/10455/src/strings/strings.go:214
			}
//line /snap/go/10455/src/strings/strings.go:214
			// _ = "end of CoverTab[1317]"
		}
//line /snap/go/10455/src/strings/strings.go:215
		if _go_fuzz_dep_.CoverTab[786518] == 0 {
//line /snap/go/10455/src/strings/strings.go:215
			_go_fuzz_dep_.CoverTab[525529]++
//line /snap/go/10455/src/strings/strings.go:215
		} else {
//line /snap/go/10455/src/strings/strings.go:215
			_go_fuzz_dep_.CoverTab[525530]++
//line /snap/go/10455/src/strings/strings.go:215
		}
//line /snap/go/10455/src/strings/strings.go:215
		// _ = "end of CoverTab[1313]"
//line /snap/go/10455/src/strings/strings.go:215
		_go_fuzz_dep_.CoverTab[1314]++
								return -1
//line /snap/go/10455/src/strings/strings.go:216
		// _ = "end of CoverTab[1314]"
	} else {
//line /snap/go/10455/src/strings/strings.go:217
		_go_fuzz_dep_.CoverTab[525255]++
//line /snap/go/10455/src/strings/strings.go:217
		_go_fuzz_dep_.CoverTab[1320]++
//line /snap/go/10455/src/strings/strings.go:217
		// _ = "end of CoverTab[1320]"
//line /snap/go/10455/src/strings/strings.go:217
	}
//line /snap/go/10455/src/strings/strings.go:217
	// _ = "end of CoverTab[1291]"
//line /snap/go/10455/src/strings/strings.go:217
	_go_fuzz_dep_.CoverTab[1292]++
//line /snap/go/10455/src/strings/strings.go:217
	_go_fuzz_dep_.CoverTab[786516] = 0
							for i := len(s); i > 0; {
//line /snap/go/10455/src/strings/strings.go:218
		if _go_fuzz_dep_.CoverTab[786516] == 0 {
//line /snap/go/10455/src/strings/strings.go:218
			_go_fuzz_dep_.CoverTab[525519]++
//line /snap/go/10455/src/strings/strings.go:218
		} else {
//line /snap/go/10455/src/strings/strings.go:218
			_go_fuzz_dep_.CoverTab[525520]++
//line /snap/go/10455/src/strings/strings.go:218
		}
//line /snap/go/10455/src/strings/strings.go:218
		_go_fuzz_dep_.CoverTab[786516] = 1
//line /snap/go/10455/src/strings/strings.go:218
		_go_fuzz_dep_.CoverTab[1321]++
								r, size := utf8.DecodeLastRuneInString(s[:i])
								i -= size
								if IndexRune(chars, r) >= 0 {
//line /snap/go/10455/src/strings/strings.go:221
			_go_fuzz_dep_.CoverTab[525260]++
//line /snap/go/10455/src/strings/strings.go:221
			_go_fuzz_dep_.CoverTab[1322]++
									return i
//line /snap/go/10455/src/strings/strings.go:222
			// _ = "end of CoverTab[1322]"
		} else {
//line /snap/go/10455/src/strings/strings.go:223
			_go_fuzz_dep_.CoverTab[525261]++
//line /snap/go/10455/src/strings/strings.go:223
			_go_fuzz_dep_.CoverTab[1323]++
//line /snap/go/10455/src/strings/strings.go:223
			// _ = "end of CoverTab[1323]"
//line /snap/go/10455/src/strings/strings.go:223
		}
//line /snap/go/10455/src/strings/strings.go:223
		// _ = "end of CoverTab[1321]"
	}
//line /snap/go/10455/src/strings/strings.go:224
	if _go_fuzz_dep_.CoverTab[786516] == 0 {
//line /snap/go/10455/src/strings/strings.go:224
		_go_fuzz_dep_.CoverTab[525521]++
//line /snap/go/10455/src/strings/strings.go:224
	} else {
//line /snap/go/10455/src/strings/strings.go:224
		_go_fuzz_dep_.CoverTab[525522]++
//line /snap/go/10455/src/strings/strings.go:224
	}
//line /snap/go/10455/src/strings/strings.go:224
	// _ = "end of CoverTab[1292]"
//line /snap/go/10455/src/strings/strings.go:224
	_go_fuzz_dep_.CoverTab[1293]++
							return -1
//line /snap/go/10455/src/strings/strings.go:225
	// _ = "end of CoverTab[1293]"
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte(s string, c byte) int {
//line /snap/go/10455/src/strings/strings.go:229
	_go_fuzz_dep_.CoverTab[1324]++
//line /snap/go/10455/src/strings/strings.go:229
	_go_fuzz_dep_.CoverTab[786519] = 0
							for i := len(s) - 1; i >= 0; i-- {
//line /snap/go/10455/src/strings/strings.go:230
		if _go_fuzz_dep_.CoverTab[786519] == 0 {
//line /snap/go/10455/src/strings/strings.go:230
			_go_fuzz_dep_.CoverTab[525531]++
//line /snap/go/10455/src/strings/strings.go:230
		} else {
//line /snap/go/10455/src/strings/strings.go:230
			_go_fuzz_dep_.CoverTab[525532]++
//line /snap/go/10455/src/strings/strings.go:230
		}
//line /snap/go/10455/src/strings/strings.go:230
		_go_fuzz_dep_.CoverTab[786519] = 1
//line /snap/go/10455/src/strings/strings.go:230
		_go_fuzz_dep_.CoverTab[1326]++
								if s[i] == c {
//line /snap/go/10455/src/strings/strings.go:231
			_go_fuzz_dep_.CoverTab[525262]++
//line /snap/go/10455/src/strings/strings.go:231
			_go_fuzz_dep_.CoverTab[1327]++
									return i
//line /snap/go/10455/src/strings/strings.go:232
			// _ = "end of CoverTab[1327]"
		} else {
//line /snap/go/10455/src/strings/strings.go:233
			_go_fuzz_dep_.CoverTab[525263]++
//line /snap/go/10455/src/strings/strings.go:233
			_go_fuzz_dep_.CoverTab[1328]++
//line /snap/go/10455/src/strings/strings.go:233
			// _ = "end of CoverTab[1328]"
//line /snap/go/10455/src/strings/strings.go:233
		}
//line /snap/go/10455/src/strings/strings.go:233
		// _ = "end of CoverTab[1326]"
	}
//line /snap/go/10455/src/strings/strings.go:234
	if _go_fuzz_dep_.CoverTab[786519] == 0 {
//line /snap/go/10455/src/strings/strings.go:234
		_go_fuzz_dep_.CoverTab[525533]++
//line /snap/go/10455/src/strings/strings.go:234
	} else {
//line /snap/go/10455/src/strings/strings.go:234
		_go_fuzz_dep_.CoverTab[525534]++
//line /snap/go/10455/src/strings/strings.go:234
	}
//line /snap/go/10455/src/strings/strings.go:234
	// _ = "end of CoverTab[1324]"
//line /snap/go/10455/src/strings/strings.go:234
	_go_fuzz_dep_.CoverTab[1325]++
							return -1
//line /snap/go/10455/src/strings/strings.go:235
	// _ = "end of CoverTab[1325]"
}

// Generic split: splits after each instance of sep,
//line /snap/go/10455/src/strings/strings.go:238
// including sepSave bytes of sep in the subarrays.
//line /snap/go/10455/src/strings/strings.go:240
func genSplit(s, sep string, sepSave, n int) []string {
//line /snap/go/10455/src/strings/strings.go:240
	_go_fuzz_dep_.CoverTab[1329]++
							if n == 0 {
//line /snap/go/10455/src/strings/strings.go:241
		_go_fuzz_dep_.CoverTab[525264]++
//line /snap/go/10455/src/strings/strings.go:241
		_go_fuzz_dep_.CoverTab[1335]++
								return nil
//line /snap/go/10455/src/strings/strings.go:242
		// _ = "end of CoverTab[1335]"
	} else {
//line /snap/go/10455/src/strings/strings.go:243
		_go_fuzz_dep_.CoverTab[525265]++
//line /snap/go/10455/src/strings/strings.go:243
		_go_fuzz_dep_.CoverTab[1336]++
//line /snap/go/10455/src/strings/strings.go:243
		// _ = "end of CoverTab[1336]"
//line /snap/go/10455/src/strings/strings.go:243
	}
//line /snap/go/10455/src/strings/strings.go:243
	// _ = "end of CoverTab[1329]"
//line /snap/go/10455/src/strings/strings.go:243
	_go_fuzz_dep_.CoverTab[1330]++
							if sep == "" {
//line /snap/go/10455/src/strings/strings.go:244
		_go_fuzz_dep_.CoverTab[525266]++
//line /snap/go/10455/src/strings/strings.go:244
		_go_fuzz_dep_.CoverTab[1337]++
								return explode(s, n)
//line /snap/go/10455/src/strings/strings.go:245
		// _ = "end of CoverTab[1337]"
	} else {
//line /snap/go/10455/src/strings/strings.go:246
		_go_fuzz_dep_.CoverTab[525267]++
//line /snap/go/10455/src/strings/strings.go:246
		_go_fuzz_dep_.CoverTab[1338]++
//line /snap/go/10455/src/strings/strings.go:246
		// _ = "end of CoverTab[1338]"
//line /snap/go/10455/src/strings/strings.go:246
	}
//line /snap/go/10455/src/strings/strings.go:246
	// _ = "end of CoverTab[1330]"
//line /snap/go/10455/src/strings/strings.go:246
	_go_fuzz_dep_.CoverTab[1331]++
							if n < 0 {
//line /snap/go/10455/src/strings/strings.go:247
		_go_fuzz_dep_.CoverTab[525268]++
//line /snap/go/10455/src/strings/strings.go:247
		_go_fuzz_dep_.CoverTab[1339]++
								n = Count(s, sep) + 1
//line /snap/go/10455/src/strings/strings.go:248
		// _ = "end of CoverTab[1339]"
	} else {
//line /snap/go/10455/src/strings/strings.go:249
		_go_fuzz_dep_.CoverTab[525269]++
//line /snap/go/10455/src/strings/strings.go:249
		_go_fuzz_dep_.CoverTab[1340]++
//line /snap/go/10455/src/strings/strings.go:249
		// _ = "end of CoverTab[1340]"
//line /snap/go/10455/src/strings/strings.go:249
	}
//line /snap/go/10455/src/strings/strings.go:249
	// _ = "end of CoverTab[1331]"
//line /snap/go/10455/src/strings/strings.go:249
	_go_fuzz_dep_.CoverTab[1332]++

							if n > len(s)+1 {
//line /snap/go/10455/src/strings/strings.go:251
		_go_fuzz_dep_.CoverTab[525270]++
//line /snap/go/10455/src/strings/strings.go:251
		_go_fuzz_dep_.CoverTab[1341]++
								n = len(s) + 1
//line /snap/go/10455/src/strings/strings.go:252
		// _ = "end of CoverTab[1341]"
	} else {
//line /snap/go/10455/src/strings/strings.go:253
		_go_fuzz_dep_.CoverTab[525271]++
//line /snap/go/10455/src/strings/strings.go:253
		_go_fuzz_dep_.CoverTab[1342]++
//line /snap/go/10455/src/strings/strings.go:253
		// _ = "end of CoverTab[1342]"
//line /snap/go/10455/src/strings/strings.go:253
	}
//line /snap/go/10455/src/strings/strings.go:253
	// _ = "end of CoverTab[1332]"
//line /snap/go/10455/src/strings/strings.go:253
	_go_fuzz_dep_.CoverTab[1333]++
							a := make([]string, n)
							n--
							i := 0
//line /snap/go/10455/src/strings/strings.go:256
	_go_fuzz_dep_.CoverTab[786520] = 0
							for i < n {
//line /snap/go/10455/src/strings/strings.go:257
		if _go_fuzz_dep_.CoverTab[786520] == 0 {
//line /snap/go/10455/src/strings/strings.go:257
			_go_fuzz_dep_.CoverTab[525535]++
//line /snap/go/10455/src/strings/strings.go:257
		} else {
//line /snap/go/10455/src/strings/strings.go:257
			_go_fuzz_dep_.CoverTab[525536]++
//line /snap/go/10455/src/strings/strings.go:257
		}
//line /snap/go/10455/src/strings/strings.go:257
		_go_fuzz_dep_.CoverTab[786520] = 1
//line /snap/go/10455/src/strings/strings.go:257
		_go_fuzz_dep_.CoverTab[1343]++
								m := Index(s, sep)
								if m < 0 {
//line /snap/go/10455/src/strings/strings.go:259
			_go_fuzz_dep_.CoverTab[525272]++
//line /snap/go/10455/src/strings/strings.go:259
			_go_fuzz_dep_.CoverTab[1345]++
									break
//line /snap/go/10455/src/strings/strings.go:260
			// _ = "end of CoverTab[1345]"
		} else {
//line /snap/go/10455/src/strings/strings.go:261
			_go_fuzz_dep_.CoverTab[525273]++
//line /snap/go/10455/src/strings/strings.go:261
			_go_fuzz_dep_.CoverTab[1346]++
//line /snap/go/10455/src/strings/strings.go:261
			// _ = "end of CoverTab[1346]"
//line /snap/go/10455/src/strings/strings.go:261
		}
//line /snap/go/10455/src/strings/strings.go:261
		// _ = "end of CoverTab[1343]"
//line /snap/go/10455/src/strings/strings.go:261
		_go_fuzz_dep_.CoverTab[1344]++
								a[i] = s[:m+sepSave]
								s = s[m+len(sep):]
								i++
//line /snap/go/10455/src/strings/strings.go:264
		// _ = "end of CoverTab[1344]"
	}
//line /snap/go/10455/src/strings/strings.go:265
	if _go_fuzz_dep_.CoverTab[786520] == 0 {
//line /snap/go/10455/src/strings/strings.go:265
		_go_fuzz_dep_.CoverTab[525537]++
//line /snap/go/10455/src/strings/strings.go:265
	} else {
//line /snap/go/10455/src/strings/strings.go:265
		_go_fuzz_dep_.CoverTab[525538]++
//line /snap/go/10455/src/strings/strings.go:265
	}
//line /snap/go/10455/src/strings/strings.go:265
	// _ = "end of CoverTab[1333]"
//line /snap/go/10455/src/strings/strings.go:265
	_go_fuzz_dep_.CoverTab[1334]++
							a[i] = s
							return a[:i+1]
//line /snap/go/10455/src/strings/strings.go:267
	// _ = "end of CoverTab[1334]"
}

// SplitN slices s into substrings separated by sep and returns a slice of
//line /snap/go/10455/src/strings/strings.go:270
// the substrings between those separators.
//line /snap/go/10455/src/strings/strings.go:270
//
//line /snap/go/10455/src/strings/strings.go:270
// The count determines the number of substrings to return:
//line /snap/go/10455/src/strings/strings.go:270
//
//line /snap/go/10455/src/strings/strings.go:270
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//line /snap/go/10455/src/strings/strings.go:270
//	n == 0: the result is nil (zero substrings)
//line /snap/go/10455/src/strings/strings.go:270
//	n < 0: all substrings
//line /snap/go/10455/src/strings/strings.go:270
//
//line /snap/go/10455/src/strings/strings.go:270
// Edge cases for s and sep (for example, empty strings) are handled
//line /snap/go/10455/src/strings/strings.go:270
// as described in the documentation for Split.
//line /snap/go/10455/src/strings/strings.go:270
//
//line /snap/go/10455/src/strings/strings.go:270
// To split around the first instance of a separator, see Cut.
//line /snap/go/10455/src/strings/strings.go:283
func SplitN(s, sep string, n int) []string {
//line /snap/go/10455/src/strings/strings.go:283
	_go_fuzz_dep_.CoverTab[1347]++
//line /snap/go/10455/src/strings/strings.go:283
	return genSplit(s, sep, 0, n)
//line /snap/go/10455/src/strings/strings.go:283
	// _ = "end of CoverTab[1347]"
//line /snap/go/10455/src/strings/strings.go:283
}

// SplitAfterN slices s into substrings after each instance of sep and
//line /snap/go/10455/src/strings/strings.go:285
// returns a slice of those substrings.
//line /snap/go/10455/src/strings/strings.go:285
//
//line /snap/go/10455/src/strings/strings.go:285
// The count determines the number of substrings to return:
//line /snap/go/10455/src/strings/strings.go:285
//
//line /snap/go/10455/src/strings/strings.go:285
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//line /snap/go/10455/src/strings/strings.go:285
//	n == 0: the result is nil (zero substrings)
//line /snap/go/10455/src/strings/strings.go:285
//	n < 0: all substrings
//line /snap/go/10455/src/strings/strings.go:285
//
//line /snap/go/10455/src/strings/strings.go:285
// Edge cases for s and sep (for example, empty strings) are handled
//line /snap/go/10455/src/strings/strings.go:285
// as described in the documentation for SplitAfter.
//line /snap/go/10455/src/strings/strings.go:296
func SplitAfterN(s, sep string, n int) []string {
//line /snap/go/10455/src/strings/strings.go:296
	_go_fuzz_dep_.CoverTab[1348]++
							return genSplit(s, sep, len(sep), n)
//line /snap/go/10455/src/strings/strings.go:297
	// _ = "end of CoverTab[1348]"
}

// Split slices s into all substrings separated by sep and returns a slice of
//line /snap/go/10455/src/strings/strings.go:300
// the substrings between those separators.
//line /snap/go/10455/src/strings/strings.go:300
//
//line /snap/go/10455/src/strings/strings.go:300
// If s does not contain sep and sep is not empty, Split returns a
//line /snap/go/10455/src/strings/strings.go:300
// slice of length 1 whose only element is s.
//line /snap/go/10455/src/strings/strings.go:300
//
//line /snap/go/10455/src/strings/strings.go:300
// If sep is empty, Split splits after each UTF-8 sequence. If both s
//line /snap/go/10455/src/strings/strings.go:300
// and sep are empty, Split returns an empty slice.
//line /snap/go/10455/src/strings/strings.go:300
//
//line /snap/go/10455/src/strings/strings.go:300
// It is equivalent to SplitN with a count of -1.
//line /snap/go/10455/src/strings/strings.go:300
//
//line /snap/go/10455/src/strings/strings.go:300
// To split around the first instance of a separator, see Cut.
//line /snap/go/10455/src/strings/strings.go:312
func Split(s, sep string) []string {
//line /snap/go/10455/src/strings/strings.go:312
	_go_fuzz_dep_.CoverTab[1349]++
//line /snap/go/10455/src/strings/strings.go:312
	return genSplit(s, sep, 0, -1)
//line /snap/go/10455/src/strings/strings.go:312
	// _ = "end of CoverTab[1349]"
//line /snap/go/10455/src/strings/strings.go:312
}

// SplitAfter slices s into all substrings after each instance of sep and
//line /snap/go/10455/src/strings/strings.go:314
// returns a slice of those substrings.
//line /snap/go/10455/src/strings/strings.go:314
//
//line /snap/go/10455/src/strings/strings.go:314
// If s does not contain sep and sep is not empty, SplitAfter returns
//line /snap/go/10455/src/strings/strings.go:314
// a slice of length 1 whose only element is s.
//line /snap/go/10455/src/strings/strings.go:314
//
//line /snap/go/10455/src/strings/strings.go:314
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
//line /snap/go/10455/src/strings/strings.go:314
// both s and sep are empty, SplitAfter returns an empty slice.
//line /snap/go/10455/src/strings/strings.go:314
//
//line /snap/go/10455/src/strings/strings.go:314
// It is equivalent to SplitAfterN with a count of -1.
//line /snap/go/10455/src/strings/strings.go:324
func SplitAfter(s, sep string) []string {
//line /snap/go/10455/src/strings/strings.go:324
	_go_fuzz_dep_.CoverTab[1350]++
							return genSplit(s, sep, len(sep), -1)
//line /snap/go/10455/src/strings/strings.go:325
	// _ = "end of CoverTab[1350]"
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// Fields splits the string s around each instance of one or more consecutive white space
//line /snap/go/10455/src/strings/strings.go:330
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
//line /snap/go/10455/src/strings/strings.go:330
// empty slice if s contains only white space.
//line /snap/go/10455/src/strings/strings.go:333
func Fields(s string) []string {
//line /snap/go/10455/src/strings/strings.go:333
	_go_fuzz_dep_.CoverTab[1351]++

//line /snap/go/10455/src/strings/strings.go:336
	n := 0
							wasSpace := 1

							setBits := uint8(0)
//line /snap/go/10455/src/strings/strings.go:339
	_go_fuzz_dep_.CoverTab[786521] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/strings.go:340
		if _go_fuzz_dep_.CoverTab[786521] == 0 {
//line /snap/go/10455/src/strings/strings.go:340
			_go_fuzz_dep_.CoverTab[525539]++
//line /snap/go/10455/src/strings/strings.go:340
		} else {
//line /snap/go/10455/src/strings/strings.go:340
			_go_fuzz_dep_.CoverTab[525540]++
//line /snap/go/10455/src/strings/strings.go:340
		}
//line /snap/go/10455/src/strings/strings.go:340
		_go_fuzz_dep_.CoverTab[786521] = 1
//line /snap/go/10455/src/strings/strings.go:340
		_go_fuzz_dep_.CoverTab[1357]++
								r := s[i]
								setBits |= r
								isSpace := int(asciiSpace[r])
								n += wasSpace & ^isSpace
								wasSpace = isSpace
//line /snap/go/10455/src/strings/strings.go:345
		// _ = "end of CoverTab[1357]"
	}
//line /snap/go/10455/src/strings/strings.go:346
	if _go_fuzz_dep_.CoverTab[786521] == 0 {
//line /snap/go/10455/src/strings/strings.go:346
		_go_fuzz_dep_.CoverTab[525541]++
//line /snap/go/10455/src/strings/strings.go:346
	} else {
//line /snap/go/10455/src/strings/strings.go:346
		_go_fuzz_dep_.CoverTab[525542]++
//line /snap/go/10455/src/strings/strings.go:346
	}
//line /snap/go/10455/src/strings/strings.go:346
	// _ = "end of CoverTab[1351]"
//line /snap/go/10455/src/strings/strings.go:346
	_go_fuzz_dep_.CoverTab[1352]++

							if setBits >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:348
		_go_fuzz_dep_.CoverTab[525274]++
//line /snap/go/10455/src/strings/strings.go:348
		_go_fuzz_dep_.CoverTab[1358]++

								return FieldsFunc(s, unicode.IsSpace)
//line /snap/go/10455/src/strings/strings.go:350
		// _ = "end of CoverTab[1358]"
	} else {
//line /snap/go/10455/src/strings/strings.go:351
		_go_fuzz_dep_.CoverTab[525275]++
//line /snap/go/10455/src/strings/strings.go:351
		_go_fuzz_dep_.CoverTab[1359]++
//line /snap/go/10455/src/strings/strings.go:351
		// _ = "end of CoverTab[1359]"
//line /snap/go/10455/src/strings/strings.go:351
	}
//line /snap/go/10455/src/strings/strings.go:351
	// _ = "end of CoverTab[1352]"
//line /snap/go/10455/src/strings/strings.go:351
	_go_fuzz_dep_.CoverTab[1353]++

							a := make([]string, n)
							na := 0
							fieldStart := 0
							i := 0
//line /snap/go/10455/src/strings/strings.go:356
	_go_fuzz_dep_.CoverTab[786522] = 0

							for i < len(s) && func() bool {
//line /snap/go/10455/src/strings/strings.go:358
		_go_fuzz_dep_.CoverTab[1360]++
//line /snap/go/10455/src/strings/strings.go:358
		return asciiSpace[s[i]] != 0
//line /snap/go/10455/src/strings/strings.go:358
		// _ = "end of CoverTab[1360]"
//line /snap/go/10455/src/strings/strings.go:358
	}() {
//line /snap/go/10455/src/strings/strings.go:358
		if _go_fuzz_dep_.CoverTab[786522] == 0 {
//line /snap/go/10455/src/strings/strings.go:358
			_go_fuzz_dep_.CoverTab[525543]++
//line /snap/go/10455/src/strings/strings.go:358
		} else {
//line /snap/go/10455/src/strings/strings.go:358
			_go_fuzz_dep_.CoverTab[525544]++
//line /snap/go/10455/src/strings/strings.go:358
		}
//line /snap/go/10455/src/strings/strings.go:358
		_go_fuzz_dep_.CoverTab[786522] = 1
//line /snap/go/10455/src/strings/strings.go:358
		_go_fuzz_dep_.CoverTab[1361]++
								i++
//line /snap/go/10455/src/strings/strings.go:359
		// _ = "end of CoverTab[1361]"
	}
//line /snap/go/10455/src/strings/strings.go:360
	if _go_fuzz_dep_.CoverTab[786522] == 0 {
//line /snap/go/10455/src/strings/strings.go:360
		_go_fuzz_dep_.CoverTab[525545]++
//line /snap/go/10455/src/strings/strings.go:360
	} else {
//line /snap/go/10455/src/strings/strings.go:360
		_go_fuzz_dep_.CoverTab[525546]++
//line /snap/go/10455/src/strings/strings.go:360
	}
//line /snap/go/10455/src/strings/strings.go:360
	// _ = "end of CoverTab[1353]"
//line /snap/go/10455/src/strings/strings.go:360
	_go_fuzz_dep_.CoverTab[1354]++
							fieldStart = i
//line /snap/go/10455/src/strings/strings.go:361
	_go_fuzz_dep_.CoverTab[786523] = 0
							for i < len(s) {
//line /snap/go/10455/src/strings/strings.go:362
		if _go_fuzz_dep_.CoverTab[786523] == 0 {
//line /snap/go/10455/src/strings/strings.go:362
			_go_fuzz_dep_.CoverTab[525547]++
//line /snap/go/10455/src/strings/strings.go:362
		} else {
//line /snap/go/10455/src/strings/strings.go:362
			_go_fuzz_dep_.CoverTab[525548]++
//line /snap/go/10455/src/strings/strings.go:362
		}
//line /snap/go/10455/src/strings/strings.go:362
		_go_fuzz_dep_.CoverTab[786523] = 1
//line /snap/go/10455/src/strings/strings.go:362
		_go_fuzz_dep_.CoverTab[1362]++
								if asciiSpace[s[i]] == 0 {
//line /snap/go/10455/src/strings/strings.go:363
			_go_fuzz_dep_.CoverTab[525276]++
//line /snap/go/10455/src/strings/strings.go:363
			_go_fuzz_dep_.CoverTab[1365]++
									i++
									continue
//line /snap/go/10455/src/strings/strings.go:365
			// _ = "end of CoverTab[1365]"
		} else {
//line /snap/go/10455/src/strings/strings.go:366
			_go_fuzz_dep_.CoverTab[525277]++
//line /snap/go/10455/src/strings/strings.go:366
			_go_fuzz_dep_.CoverTab[1366]++
//line /snap/go/10455/src/strings/strings.go:366
			// _ = "end of CoverTab[1366]"
//line /snap/go/10455/src/strings/strings.go:366
		}
//line /snap/go/10455/src/strings/strings.go:366
		// _ = "end of CoverTab[1362]"
//line /snap/go/10455/src/strings/strings.go:366
		_go_fuzz_dep_.CoverTab[1363]++
								a[na] = s[fieldStart:i]
								na++
								i++
//line /snap/go/10455/src/strings/strings.go:369
		_go_fuzz_dep_.CoverTab[786524] = 0

								for i < len(s) && func() bool {
//line /snap/go/10455/src/strings/strings.go:371
			_go_fuzz_dep_.CoverTab[1367]++
//line /snap/go/10455/src/strings/strings.go:371
			return asciiSpace[s[i]] != 0
//line /snap/go/10455/src/strings/strings.go:371
			// _ = "end of CoverTab[1367]"
//line /snap/go/10455/src/strings/strings.go:371
		}() {
//line /snap/go/10455/src/strings/strings.go:371
			if _go_fuzz_dep_.CoverTab[786524] == 0 {
//line /snap/go/10455/src/strings/strings.go:371
				_go_fuzz_dep_.CoverTab[525551]++
//line /snap/go/10455/src/strings/strings.go:371
			} else {
//line /snap/go/10455/src/strings/strings.go:371
				_go_fuzz_dep_.CoverTab[525552]++
//line /snap/go/10455/src/strings/strings.go:371
			}
//line /snap/go/10455/src/strings/strings.go:371
			_go_fuzz_dep_.CoverTab[786524] = 1
//line /snap/go/10455/src/strings/strings.go:371
			_go_fuzz_dep_.CoverTab[1368]++
									i++
//line /snap/go/10455/src/strings/strings.go:372
			// _ = "end of CoverTab[1368]"
		}
//line /snap/go/10455/src/strings/strings.go:373
		if _go_fuzz_dep_.CoverTab[786524] == 0 {
//line /snap/go/10455/src/strings/strings.go:373
			_go_fuzz_dep_.CoverTab[525553]++
//line /snap/go/10455/src/strings/strings.go:373
		} else {
//line /snap/go/10455/src/strings/strings.go:373
			_go_fuzz_dep_.CoverTab[525554]++
//line /snap/go/10455/src/strings/strings.go:373
		}
//line /snap/go/10455/src/strings/strings.go:373
		// _ = "end of CoverTab[1363]"
//line /snap/go/10455/src/strings/strings.go:373
		_go_fuzz_dep_.CoverTab[1364]++
								fieldStart = i
//line /snap/go/10455/src/strings/strings.go:374
		// _ = "end of CoverTab[1364]"
	}
//line /snap/go/10455/src/strings/strings.go:375
	if _go_fuzz_dep_.CoverTab[786523] == 0 {
//line /snap/go/10455/src/strings/strings.go:375
		_go_fuzz_dep_.CoverTab[525549]++
//line /snap/go/10455/src/strings/strings.go:375
	} else {
//line /snap/go/10455/src/strings/strings.go:375
		_go_fuzz_dep_.CoverTab[525550]++
//line /snap/go/10455/src/strings/strings.go:375
	}
//line /snap/go/10455/src/strings/strings.go:375
	// _ = "end of CoverTab[1354]"
//line /snap/go/10455/src/strings/strings.go:375
	_go_fuzz_dep_.CoverTab[1355]++
							if fieldStart < len(s) {
//line /snap/go/10455/src/strings/strings.go:376
		_go_fuzz_dep_.CoverTab[525278]++
//line /snap/go/10455/src/strings/strings.go:376
		_go_fuzz_dep_.CoverTab[1369]++
								a[na] = s[fieldStart:]
//line /snap/go/10455/src/strings/strings.go:377
		// _ = "end of CoverTab[1369]"
	} else {
//line /snap/go/10455/src/strings/strings.go:378
		_go_fuzz_dep_.CoverTab[525279]++
//line /snap/go/10455/src/strings/strings.go:378
		_go_fuzz_dep_.CoverTab[1370]++
//line /snap/go/10455/src/strings/strings.go:378
		// _ = "end of CoverTab[1370]"
//line /snap/go/10455/src/strings/strings.go:378
	}
//line /snap/go/10455/src/strings/strings.go:378
	// _ = "end of CoverTab[1355]"
//line /snap/go/10455/src/strings/strings.go:378
	_go_fuzz_dep_.CoverTab[1356]++
							return a
//line /snap/go/10455/src/strings/strings.go:379
	// _ = "end of CoverTab[1356]"
}

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
//line /snap/go/10455/src/strings/strings.go:382
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
//line /snap/go/10455/src/strings/strings.go:382
// string is empty, an empty slice is returned.
//line /snap/go/10455/src/strings/strings.go:382
//
//line /snap/go/10455/src/strings/strings.go:382
// FieldsFunc makes no guarantees about the order in which it calls f(c)
//line /snap/go/10455/src/strings/strings.go:382
// and assumes that f always returns the same value for a given c.
//line /snap/go/10455/src/strings/strings.go:388
func FieldsFunc(s string, f func(rune) bool) []string {
//line /snap/go/10455/src/strings/strings.go:388
	_go_fuzz_dep_.CoverTab[1371]++
	// A span is used to record a slice of s of the form s[start:end].
	// The start index is inclusive and the end index is exclusive.
	type span struct {
		start	int
		end	int
	}
							spans := make([]span, 0, 32)

//line /snap/go/10455/src/strings/strings.go:401
	start := -1
//line /snap/go/10455/src/strings/strings.go:401
	_go_fuzz_dep_.CoverTab[786525] = 0
							for end, rune := range s {
//line /snap/go/10455/src/strings/strings.go:402
		if _go_fuzz_dep_.CoverTab[786525] == 0 {
//line /snap/go/10455/src/strings/strings.go:402
			_go_fuzz_dep_.CoverTab[525555]++
//line /snap/go/10455/src/strings/strings.go:402
		} else {
//line /snap/go/10455/src/strings/strings.go:402
			_go_fuzz_dep_.CoverTab[525556]++
//line /snap/go/10455/src/strings/strings.go:402
		}
//line /snap/go/10455/src/strings/strings.go:402
		_go_fuzz_dep_.CoverTab[786525] = 1
//line /snap/go/10455/src/strings/strings.go:402
		_go_fuzz_dep_.CoverTab[1375]++
								if f(rune) {
//line /snap/go/10455/src/strings/strings.go:403
			_go_fuzz_dep_.CoverTab[525280]++
//line /snap/go/10455/src/strings/strings.go:403
			_go_fuzz_dep_.CoverTab[1376]++
									if start >= 0 {
//line /snap/go/10455/src/strings/strings.go:404
				_go_fuzz_dep_.CoverTab[525282]++
//line /snap/go/10455/src/strings/strings.go:404
				_go_fuzz_dep_.CoverTab[1377]++
										spans = append(spans, span{start, end})

//line /snap/go/10455/src/strings/strings.go:409
				start = ^start
//line /snap/go/10455/src/strings/strings.go:409
				// _ = "end of CoverTab[1377]"
			} else {
//line /snap/go/10455/src/strings/strings.go:410
				_go_fuzz_dep_.CoverTab[525283]++
//line /snap/go/10455/src/strings/strings.go:410
				_go_fuzz_dep_.CoverTab[1378]++
//line /snap/go/10455/src/strings/strings.go:410
				// _ = "end of CoverTab[1378]"
//line /snap/go/10455/src/strings/strings.go:410
			}
//line /snap/go/10455/src/strings/strings.go:410
			// _ = "end of CoverTab[1376]"
		} else {
//line /snap/go/10455/src/strings/strings.go:411
			_go_fuzz_dep_.CoverTab[525281]++
//line /snap/go/10455/src/strings/strings.go:411
			_go_fuzz_dep_.CoverTab[1379]++
									if start < 0 {
//line /snap/go/10455/src/strings/strings.go:412
				_go_fuzz_dep_.CoverTab[525284]++
//line /snap/go/10455/src/strings/strings.go:412
				_go_fuzz_dep_.CoverTab[1380]++
										start = end
//line /snap/go/10455/src/strings/strings.go:413
				// _ = "end of CoverTab[1380]"
			} else {
//line /snap/go/10455/src/strings/strings.go:414
				_go_fuzz_dep_.CoverTab[525285]++
//line /snap/go/10455/src/strings/strings.go:414
				_go_fuzz_dep_.CoverTab[1381]++
//line /snap/go/10455/src/strings/strings.go:414
				// _ = "end of CoverTab[1381]"
//line /snap/go/10455/src/strings/strings.go:414
			}
//line /snap/go/10455/src/strings/strings.go:414
			// _ = "end of CoverTab[1379]"
		}
//line /snap/go/10455/src/strings/strings.go:415
		// _ = "end of CoverTab[1375]"
	}
//line /snap/go/10455/src/strings/strings.go:416
	if _go_fuzz_dep_.CoverTab[786525] == 0 {
//line /snap/go/10455/src/strings/strings.go:416
		_go_fuzz_dep_.CoverTab[525557]++
//line /snap/go/10455/src/strings/strings.go:416
	} else {
//line /snap/go/10455/src/strings/strings.go:416
		_go_fuzz_dep_.CoverTab[525558]++
//line /snap/go/10455/src/strings/strings.go:416
	}
//line /snap/go/10455/src/strings/strings.go:416
	// _ = "end of CoverTab[1371]"
//line /snap/go/10455/src/strings/strings.go:416
	_go_fuzz_dep_.CoverTab[1372]++

//line /snap/go/10455/src/strings/strings.go:419
	if start >= 0 {
//line /snap/go/10455/src/strings/strings.go:419
		_go_fuzz_dep_.CoverTab[525286]++
//line /snap/go/10455/src/strings/strings.go:419
		_go_fuzz_dep_.CoverTab[1382]++
								spans = append(spans, span{start, len(s)})
//line /snap/go/10455/src/strings/strings.go:420
		// _ = "end of CoverTab[1382]"
	} else {
//line /snap/go/10455/src/strings/strings.go:421
		_go_fuzz_dep_.CoverTab[525287]++
//line /snap/go/10455/src/strings/strings.go:421
		_go_fuzz_dep_.CoverTab[1383]++
//line /snap/go/10455/src/strings/strings.go:421
		// _ = "end of CoverTab[1383]"
//line /snap/go/10455/src/strings/strings.go:421
	}
//line /snap/go/10455/src/strings/strings.go:421
	// _ = "end of CoverTab[1372]"
//line /snap/go/10455/src/strings/strings.go:421
	_go_fuzz_dep_.CoverTab[1373]++

//line /snap/go/10455/src/strings/strings.go:424
	a := make([]string, len(spans))
//line /snap/go/10455/src/strings/strings.go:424
	_go_fuzz_dep_.CoverTab[786526] = 0
							for i, span := range spans {
//line /snap/go/10455/src/strings/strings.go:425
		if _go_fuzz_dep_.CoverTab[786526] == 0 {
//line /snap/go/10455/src/strings/strings.go:425
			_go_fuzz_dep_.CoverTab[525559]++
//line /snap/go/10455/src/strings/strings.go:425
		} else {
//line /snap/go/10455/src/strings/strings.go:425
			_go_fuzz_dep_.CoverTab[525560]++
//line /snap/go/10455/src/strings/strings.go:425
		}
//line /snap/go/10455/src/strings/strings.go:425
		_go_fuzz_dep_.CoverTab[786526] = 1
//line /snap/go/10455/src/strings/strings.go:425
		_go_fuzz_dep_.CoverTab[1384]++
								a[i] = s[span.start:span.end]
//line /snap/go/10455/src/strings/strings.go:426
		// _ = "end of CoverTab[1384]"
	}
//line /snap/go/10455/src/strings/strings.go:427
	if _go_fuzz_dep_.CoverTab[786526] == 0 {
//line /snap/go/10455/src/strings/strings.go:427
		_go_fuzz_dep_.CoverTab[525561]++
//line /snap/go/10455/src/strings/strings.go:427
	} else {
//line /snap/go/10455/src/strings/strings.go:427
		_go_fuzz_dep_.CoverTab[525562]++
//line /snap/go/10455/src/strings/strings.go:427
	}
//line /snap/go/10455/src/strings/strings.go:427
	// _ = "end of CoverTab[1373]"
//line /snap/go/10455/src/strings/strings.go:427
	_go_fuzz_dep_.CoverTab[1374]++

							return a
//line /snap/go/10455/src/strings/strings.go:429
	// _ = "end of CoverTab[1374]"
}

// Join concatenates the elements of its first argument to create a single string. The separator
//line /snap/go/10455/src/strings/strings.go:432
// string sep is placed between elements in the resulting string.
//line /snap/go/10455/src/strings/strings.go:434
func Join(elems []string, sep string) string {
//line /snap/go/10455/src/strings/strings.go:434
	_go_fuzz_dep_.CoverTab[1385]++
							switch len(elems) {
	case 0:
//line /snap/go/10455/src/strings/strings.go:436
		_go_fuzz_dep_.CoverTab[525288]++
//line /snap/go/10455/src/strings/strings.go:436
		_go_fuzz_dep_.CoverTab[1390]++
								return ""
//line /snap/go/10455/src/strings/strings.go:437
		// _ = "end of CoverTab[1390]"
	case 1:
//line /snap/go/10455/src/strings/strings.go:438
		_go_fuzz_dep_.CoverTab[525289]++
//line /snap/go/10455/src/strings/strings.go:438
		_go_fuzz_dep_.CoverTab[1391]++
								return elems[0]
//line /snap/go/10455/src/strings/strings.go:439
		// _ = "end of CoverTab[1391]"
//line /snap/go/10455/src/strings/strings.go:439
	default:
//line /snap/go/10455/src/strings/strings.go:439
		_go_fuzz_dep_.CoverTab[525290]++
//line /snap/go/10455/src/strings/strings.go:439
		_go_fuzz_dep_.CoverTab[1392]++
//line /snap/go/10455/src/strings/strings.go:439
		// _ = "end of CoverTab[1392]"
	}
//line /snap/go/10455/src/strings/strings.go:440
	// _ = "end of CoverTab[1385]"
//line /snap/go/10455/src/strings/strings.go:440
	_go_fuzz_dep_.CoverTab[1386]++

							var n int
							if len(sep) > 0 {
//line /snap/go/10455/src/strings/strings.go:443
		_go_fuzz_dep_.CoverTab[525291]++
//line /snap/go/10455/src/strings/strings.go:443
		_go_fuzz_dep_.CoverTab[1393]++
								if len(sep) >= maxInt/(len(elems)-1) {
//line /snap/go/10455/src/strings/strings.go:444
			_go_fuzz_dep_.CoverTab[525293]++
//line /snap/go/10455/src/strings/strings.go:444
			_go_fuzz_dep_.CoverTab[1395]++
									panic("strings: Join output length overflow")
//line /snap/go/10455/src/strings/strings.go:445
			// _ = "end of CoverTab[1395]"
		} else {
//line /snap/go/10455/src/strings/strings.go:446
			_go_fuzz_dep_.CoverTab[525294]++
//line /snap/go/10455/src/strings/strings.go:446
			_go_fuzz_dep_.CoverTab[1396]++
//line /snap/go/10455/src/strings/strings.go:446
			// _ = "end of CoverTab[1396]"
//line /snap/go/10455/src/strings/strings.go:446
		}
//line /snap/go/10455/src/strings/strings.go:446
		// _ = "end of CoverTab[1393]"
//line /snap/go/10455/src/strings/strings.go:446
		_go_fuzz_dep_.CoverTab[1394]++
								n += len(sep) * (len(elems) - 1)
//line /snap/go/10455/src/strings/strings.go:447
		// _ = "end of CoverTab[1394]"
	} else {
//line /snap/go/10455/src/strings/strings.go:448
		_go_fuzz_dep_.CoverTab[525292]++
//line /snap/go/10455/src/strings/strings.go:448
		_go_fuzz_dep_.CoverTab[1397]++
//line /snap/go/10455/src/strings/strings.go:448
		// _ = "end of CoverTab[1397]"
//line /snap/go/10455/src/strings/strings.go:448
	}
//line /snap/go/10455/src/strings/strings.go:448
	// _ = "end of CoverTab[1386]"
//line /snap/go/10455/src/strings/strings.go:448
	_go_fuzz_dep_.CoverTab[1387]++
//line /snap/go/10455/src/strings/strings.go:448
	_go_fuzz_dep_.CoverTab[786527] = 0
							for _, elem := range elems {
//line /snap/go/10455/src/strings/strings.go:449
		if _go_fuzz_dep_.CoverTab[786527] == 0 {
//line /snap/go/10455/src/strings/strings.go:449
			_go_fuzz_dep_.CoverTab[525563]++
//line /snap/go/10455/src/strings/strings.go:449
		} else {
//line /snap/go/10455/src/strings/strings.go:449
			_go_fuzz_dep_.CoverTab[525564]++
//line /snap/go/10455/src/strings/strings.go:449
		}
//line /snap/go/10455/src/strings/strings.go:449
		_go_fuzz_dep_.CoverTab[786527] = 1
//line /snap/go/10455/src/strings/strings.go:449
		_go_fuzz_dep_.CoverTab[1398]++
								if len(elem) > maxInt-n {
//line /snap/go/10455/src/strings/strings.go:450
			_go_fuzz_dep_.CoverTab[525295]++
//line /snap/go/10455/src/strings/strings.go:450
			_go_fuzz_dep_.CoverTab[1400]++
									panic("strings: Join output length overflow")
//line /snap/go/10455/src/strings/strings.go:451
			// _ = "end of CoverTab[1400]"
		} else {
//line /snap/go/10455/src/strings/strings.go:452
			_go_fuzz_dep_.CoverTab[525296]++
//line /snap/go/10455/src/strings/strings.go:452
			_go_fuzz_dep_.CoverTab[1401]++
//line /snap/go/10455/src/strings/strings.go:452
			// _ = "end of CoverTab[1401]"
//line /snap/go/10455/src/strings/strings.go:452
		}
//line /snap/go/10455/src/strings/strings.go:452
		// _ = "end of CoverTab[1398]"
//line /snap/go/10455/src/strings/strings.go:452
		_go_fuzz_dep_.CoverTab[1399]++
								n += len(elem)
//line /snap/go/10455/src/strings/strings.go:453
		// _ = "end of CoverTab[1399]"
	}
//line /snap/go/10455/src/strings/strings.go:454
	if _go_fuzz_dep_.CoverTab[786527] == 0 {
//line /snap/go/10455/src/strings/strings.go:454
		_go_fuzz_dep_.CoverTab[525565]++
//line /snap/go/10455/src/strings/strings.go:454
	} else {
//line /snap/go/10455/src/strings/strings.go:454
		_go_fuzz_dep_.CoverTab[525566]++
//line /snap/go/10455/src/strings/strings.go:454
	}
//line /snap/go/10455/src/strings/strings.go:454
	// _ = "end of CoverTab[1387]"
//line /snap/go/10455/src/strings/strings.go:454
	_go_fuzz_dep_.CoverTab[1388]++

							var b Builder
							b.Grow(n)
							b.WriteString(elems[0])
//line /snap/go/10455/src/strings/strings.go:458
	_go_fuzz_dep_.CoverTab[786528] = 0
							for _, s := range elems[1:] {
//line /snap/go/10455/src/strings/strings.go:459
		if _go_fuzz_dep_.CoverTab[786528] == 0 {
//line /snap/go/10455/src/strings/strings.go:459
			_go_fuzz_dep_.CoverTab[525567]++
//line /snap/go/10455/src/strings/strings.go:459
		} else {
//line /snap/go/10455/src/strings/strings.go:459
			_go_fuzz_dep_.CoverTab[525568]++
//line /snap/go/10455/src/strings/strings.go:459
		}
//line /snap/go/10455/src/strings/strings.go:459
		_go_fuzz_dep_.CoverTab[786528] = 1
//line /snap/go/10455/src/strings/strings.go:459
		_go_fuzz_dep_.CoverTab[1402]++
								b.WriteString(sep)
								b.WriteString(s)
//line /snap/go/10455/src/strings/strings.go:461
		// _ = "end of CoverTab[1402]"
	}
//line /snap/go/10455/src/strings/strings.go:462
	if _go_fuzz_dep_.CoverTab[786528] == 0 {
//line /snap/go/10455/src/strings/strings.go:462
		_go_fuzz_dep_.CoverTab[525569]++
//line /snap/go/10455/src/strings/strings.go:462
	} else {
//line /snap/go/10455/src/strings/strings.go:462
		_go_fuzz_dep_.CoverTab[525570]++
//line /snap/go/10455/src/strings/strings.go:462
	}
//line /snap/go/10455/src/strings/strings.go:462
	// _ = "end of CoverTab[1388]"
//line /snap/go/10455/src/strings/strings.go:462
	_go_fuzz_dep_.CoverTab[1389]++
							return b.String()
//line /snap/go/10455/src/strings/strings.go:463
	// _ = "end of CoverTab[1389]"
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool {
//line /snap/go/10455/src/strings/strings.go:467
	_go_fuzz_dep_.CoverTab[1403]++
							return len(s) >= len(prefix) && func() bool {
//line /snap/go/10455/src/strings/strings.go:468
		_go_fuzz_dep_.CoverTab[1404]++
//line /snap/go/10455/src/strings/strings.go:468
		return s[0:len(prefix)] == prefix
//line /snap/go/10455/src/strings/strings.go:468
		// _ = "end of CoverTab[1404]"
//line /snap/go/10455/src/strings/strings.go:468
	}()
//line /snap/go/10455/src/strings/strings.go:468
	// _ = "end of CoverTab[1403]"
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
//line /snap/go/10455/src/strings/strings.go:472
	_go_fuzz_dep_.CoverTab[1405]++
							return len(s) >= len(suffix) && func() bool {
//line /snap/go/10455/src/strings/strings.go:473
		_go_fuzz_dep_.CoverTab[1406]++
//line /snap/go/10455/src/strings/strings.go:473
		return s[len(s)-len(suffix):] == suffix
//line /snap/go/10455/src/strings/strings.go:473
		// _ = "end of CoverTab[1406]"
//line /snap/go/10455/src/strings/strings.go:473
	}()
//line /snap/go/10455/src/strings/strings.go:473
	// _ = "end of CoverTab[1405]"
}

// Map returns a copy of the string s with all its characters modified
//line /snap/go/10455/src/strings/strings.go:476
// according to the mapping function. If mapping returns a negative value, the character is
//line /snap/go/10455/src/strings/strings.go:476
// dropped from the string with no replacement.
//line /snap/go/10455/src/strings/strings.go:479
func Map(mapping func(rune) rune, s string) string {
//line /snap/go/10455/src/strings/strings.go:479
	_go_fuzz_dep_.CoverTab[1407]++

//line /snap/go/10455/src/strings/strings.go:484
	// The output buffer b is initialized on demand, the first
							// time a character differs.
							var b Builder
//line /snap/go/10455/src/strings/strings.go:486
	_go_fuzz_dep_.CoverTab[786529] = 0

							for i, c := range s {
//line /snap/go/10455/src/strings/strings.go:488
		if _go_fuzz_dep_.CoverTab[786529] == 0 {
//line /snap/go/10455/src/strings/strings.go:488
			_go_fuzz_dep_.CoverTab[525571]++
//line /snap/go/10455/src/strings/strings.go:488
		} else {
//line /snap/go/10455/src/strings/strings.go:488
			_go_fuzz_dep_.CoverTab[525572]++
//line /snap/go/10455/src/strings/strings.go:488
		}
//line /snap/go/10455/src/strings/strings.go:488
		_go_fuzz_dep_.CoverTab[786529] = 1
//line /snap/go/10455/src/strings/strings.go:488
		_go_fuzz_dep_.CoverTab[1411]++
								r := mapping(c)
								if r == c && func() bool {
//line /snap/go/10455/src/strings/strings.go:490
			_go_fuzz_dep_.CoverTab[1415]++
//line /snap/go/10455/src/strings/strings.go:490
			return c != utf8.RuneError
//line /snap/go/10455/src/strings/strings.go:490
			// _ = "end of CoverTab[1415]"
//line /snap/go/10455/src/strings/strings.go:490
		}() {
//line /snap/go/10455/src/strings/strings.go:490
			_go_fuzz_dep_.CoverTab[525297]++
//line /snap/go/10455/src/strings/strings.go:490
			_go_fuzz_dep_.CoverTab[1416]++
									continue
//line /snap/go/10455/src/strings/strings.go:491
			// _ = "end of CoverTab[1416]"
		} else {
//line /snap/go/10455/src/strings/strings.go:492
			_go_fuzz_dep_.CoverTab[525298]++
//line /snap/go/10455/src/strings/strings.go:492
			_go_fuzz_dep_.CoverTab[1417]++
//line /snap/go/10455/src/strings/strings.go:492
			// _ = "end of CoverTab[1417]"
//line /snap/go/10455/src/strings/strings.go:492
		}
//line /snap/go/10455/src/strings/strings.go:492
		// _ = "end of CoverTab[1411]"
//line /snap/go/10455/src/strings/strings.go:492
		_go_fuzz_dep_.CoverTab[1412]++

								var width int
								if c == utf8.RuneError {
//line /snap/go/10455/src/strings/strings.go:495
			_go_fuzz_dep_.CoverTab[525299]++
//line /snap/go/10455/src/strings/strings.go:495
			_go_fuzz_dep_.CoverTab[1418]++
									c, width = utf8.DecodeRuneInString(s[i:])
									if width != 1 && func() bool {
//line /snap/go/10455/src/strings/strings.go:497
				_go_fuzz_dep_.CoverTab[1419]++
//line /snap/go/10455/src/strings/strings.go:497
				return r == c
//line /snap/go/10455/src/strings/strings.go:497
				// _ = "end of CoverTab[1419]"
//line /snap/go/10455/src/strings/strings.go:497
			}() {
//line /snap/go/10455/src/strings/strings.go:497
				_go_fuzz_dep_.CoverTab[525301]++
//line /snap/go/10455/src/strings/strings.go:497
				_go_fuzz_dep_.CoverTab[1420]++
										continue
//line /snap/go/10455/src/strings/strings.go:498
				// _ = "end of CoverTab[1420]"
			} else {
//line /snap/go/10455/src/strings/strings.go:499
				_go_fuzz_dep_.CoverTab[525302]++
//line /snap/go/10455/src/strings/strings.go:499
				_go_fuzz_dep_.CoverTab[1421]++
//line /snap/go/10455/src/strings/strings.go:499
				// _ = "end of CoverTab[1421]"
//line /snap/go/10455/src/strings/strings.go:499
			}
//line /snap/go/10455/src/strings/strings.go:499
			// _ = "end of CoverTab[1418]"
		} else {
//line /snap/go/10455/src/strings/strings.go:500
			_go_fuzz_dep_.CoverTab[525300]++
//line /snap/go/10455/src/strings/strings.go:500
			_go_fuzz_dep_.CoverTab[1422]++
									width = utf8.RuneLen(c)
//line /snap/go/10455/src/strings/strings.go:501
			// _ = "end of CoverTab[1422]"
		}
//line /snap/go/10455/src/strings/strings.go:502
		// _ = "end of CoverTab[1412]"
//line /snap/go/10455/src/strings/strings.go:502
		_go_fuzz_dep_.CoverTab[1413]++

								b.Grow(len(s) + utf8.UTFMax)
								b.WriteString(s[:i])
								if r >= 0 {
//line /snap/go/10455/src/strings/strings.go:506
			_go_fuzz_dep_.CoverTab[525303]++
//line /snap/go/10455/src/strings/strings.go:506
			_go_fuzz_dep_.CoverTab[1423]++
									b.WriteRune(r)
//line /snap/go/10455/src/strings/strings.go:507
			// _ = "end of CoverTab[1423]"
		} else {
//line /snap/go/10455/src/strings/strings.go:508
			_go_fuzz_dep_.CoverTab[525304]++
//line /snap/go/10455/src/strings/strings.go:508
			_go_fuzz_dep_.CoverTab[1424]++
//line /snap/go/10455/src/strings/strings.go:508
			// _ = "end of CoverTab[1424]"
//line /snap/go/10455/src/strings/strings.go:508
		}
//line /snap/go/10455/src/strings/strings.go:508
		// _ = "end of CoverTab[1413]"
//line /snap/go/10455/src/strings/strings.go:508
		_go_fuzz_dep_.CoverTab[1414]++

								s = s[i+width:]
								break
//line /snap/go/10455/src/strings/strings.go:511
		// _ = "end of CoverTab[1414]"
	}
//line /snap/go/10455/src/strings/strings.go:512
	if _go_fuzz_dep_.CoverTab[786529] == 0 {
//line /snap/go/10455/src/strings/strings.go:512
		_go_fuzz_dep_.CoverTab[525573]++
//line /snap/go/10455/src/strings/strings.go:512
	} else {
//line /snap/go/10455/src/strings/strings.go:512
		_go_fuzz_dep_.CoverTab[525574]++
//line /snap/go/10455/src/strings/strings.go:512
	}
//line /snap/go/10455/src/strings/strings.go:512
	// _ = "end of CoverTab[1407]"
//line /snap/go/10455/src/strings/strings.go:512
	_go_fuzz_dep_.CoverTab[1408]++

//line /snap/go/10455/src/strings/strings.go:515
	if b.Cap() == 0 {
//line /snap/go/10455/src/strings/strings.go:515
		_go_fuzz_dep_.CoverTab[525305]++
//line /snap/go/10455/src/strings/strings.go:515
		_go_fuzz_dep_.CoverTab[1425]++
								return s
//line /snap/go/10455/src/strings/strings.go:516
		// _ = "end of CoverTab[1425]"
	} else {
//line /snap/go/10455/src/strings/strings.go:517
		_go_fuzz_dep_.CoverTab[525306]++
//line /snap/go/10455/src/strings/strings.go:517
		_go_fuzz_dep_.CoverTab[1426]++
//line /snap/go/10455/src/strings/strings.go:517
		// _ = "end of CoverTab[1426]"
//line /snap/go/10455/src/strings/strings.go:517
	}
//line /snap/go/10455/src/strings/strings.go:517
	// _ = "end of CoverTab[1408]"
//line /snap/go/10455/src/strings/strings.go:517
	_go_fuzz_dep_.CoverTab[1409]++
//line /snap/go/10455/src/strings/strings.go:517
	_go_fuzz_dep_.CoverTab[786530] = 0

							for _, c := range s {
//line /snap/go/10455/src/strings/strings.go:519
		if _go_fuzz_dep_.CoverTab[786530] == 0 {
//line /snap/go/10455/src/strings/strings.go:519
			_go_fuzz_dep_.CoverTab[525575]++
//line /snap/go/10455/src/strings/strings.go:519
		} else {
//line /snap/go/10455/src/strings/strings.go:519
			_go_fuzz_dep_.CoverTab[525576]++
//line /snap/go/10455/src/strings/strings.go:519
		}
//line /snap/go/10455/src/strings/strings.go:519
		_go_fuzz_dep_.CoverTab[786530] = 1
//line /snap/go/10455/src/strings/strings.go:519
		_go_fuzz_dep_.CoverTab[1427]++
								r := mapping(c)

								if r >= 0 {
//line /snap/go/10455/src/strings/strings.go:522
			_go_fuzz_dep_.CoverTab[525307]++
//line /snap/go/10455/src/strings/strings.go:522
			_go_fuzz_dep_.CoverTab[1428]++

//line /snap/go/10455/src/strings/strings.go:526
			if r < utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:526
				_go_fuzz_dep_.CoverTab[525309]++
//line /snap/go/10455/src/strings/strings.go:526
				_go_fuzz_dep_.CoverTab[1429]++
										b.WriteByte(byte(r))
//line /snap/go/10455/src/strings/strings.go:527
				// _ = "end of CoverTab[1429]"
			} else {
//line /snap/go/10455/src/strings/strings.go:528
				_go_fuzz_dep_.CoverTab[525310]++
//line /snap/go/10455/src/strings/strings.go:528
				_go_fuzz_dep_.CoverTab[1430]++

										b.WriteRune(r)
//line /snap/go/10455/src/strings/strings.go:530
				// _ = "end of CoverTab[1430]"
			}
//line /snap/go/10455/src/strings/strings.go:531
			// _ = "end of CoverTab[1428]"
		} else {
//line /snap/go/10455/src/strings/strings.go:532
			_go_fuzz_dep_.CoverTab[525308]++
//line /snap/go/10455/src/strings/strings.go:532
			_go_fuzz_dep_.CoverTab[1431]++
//line /snap/go/10455/src/strings/strings.go:532
			// _ = "end of CoverTab[1431]"
//line /snap/go/10455/src/strings/strings.go:532
		}
//line /snap/go/10455/src/strings/strings.go:532
		// _ = "end of CoverTab[1427]"
	}
//line /snap/go/10455/src/strings/strings.go:533
	if _go_fuzz_dep_.CoverTab[786530] == 0 {
//line /snap/go/10455/src/strings/strings.go:533
		_go_fuzz_dep_.CoverTab[525577]++
//line /snap/go/10455/src/strings/strings.go:533
	} else {
//line /snap/go/10455/src/strings/strings.go:533
		_go_fuzz_dep_.CoverTab[525578]++
//line /snap/go/10455/src/strings/strings.go:533
	}
//line /snap/go/10455/src/strings/strings.go:533
	// _ = "end of CoverTab[1409]"
//line /snap/go/10455/src/strings/strings.go:533
	_go_fuzz_dep_.CoverTab[1410]++

							return b.String()
//line /snap/go/10455/src/strings/strings.go:535
	// _ = "end of CoverTab[1410]"
}

// Repeat returns a new string consisting of count copies of the string s.
//line /snap/go/10455/src/strings/strings.go:538
//
//line /snap/go/10455/src/strings/strings.go:538
// It panics if count is negative or if the result of (len(s) * count)
//line /snap/go/10455/src/strings/strings.go:538
// overflows.
//line /snap/go/10455/src/strings/strings.go:542
func Repeat(s string, count int) string {
//line /snap/go/10455/src/strings/strings.go:542
	_go_fuzz_dep_.CoverTab[1432]++
							switch count {
	case 0:
//line /snap/go/10455/src/strings/strings.go:544
		_go_fuzz_dep_.CoverTab[525311]++
//line /snap/go/10455/src/strings/strings.go:544
		_go_fuzz_dep_.CoverTab[1439]++
								return ""
//line /snap/go/10455/src/strings/strings.go:545
		// _ = "end of CoverTab[1439]"
	case 1:
//line /snap/go/10455/src/strings/strings.go:546
		_go_fuzz_dep_.CoverTab[525312]++
//line /snap/go/10455/src/strings/strings.go:546
		_go_fuzz_dep_.CoverTab[1440]++
								return s
//line /snap/go/10455/src/strings/strings.go:547
		// _ = "end of CoverTab[1440]"
//line /snap/go/10455/src/strings/strings.go:547
	default:
//line /snap/go/10455/src/strings/strings.go:547
		_go_fuzz_dep_.CoverTab[525313]++
//line /snap/go/10455/src/strings/strings.go:547
		_go_fuzz_dep_.CoverTab[1441]++
//line /snap/go/10455/src/strings/strings.go:547
		// _ = "end of CoverTab[1441]"
	}
//line /snap/go/10455/src/strings/strings.go:548
	// _ = "end of CoverTab[1432]"
//line /snap/go/10455/src/strings/strings.go:548
	_go_fuzz_dep_.CoverTab[1433]++

//line /snap/go/10455/src/strings/strings.go:553
	if count < 0 {
//line /snap/go/10455/src/strings/strings.go:553
		_go_fuzz_dep_.CoverTab[525314]++
//line /snap/go/10455/src/strings/strings.go:553
		_go_fuzz_dep_.CoverTab[1442]++
								panic("strings: negative Repeat count")
//line /snap/go/10455/src/strings/strings.go:554
		// _ = "end of CoverTab[1442]"
	} else {
//line /snap/go/10455/src/strings/strings.go:555
		_go_fuzz_dep_.CoverTab[525315]++
//line /snap/go/10455/src/strings/strings.go:555
		_go_fuzz_dep_.CoverTab[1443]++
//line /snap/go/10455/src/strings/strings.go:555
		// _ = "end of CoverTab[1443]"
//line /snap/go/10455/src/strings/strings.go:555
	}
//line /snap/go/10455/src/strings/strings.go:555
	// _ = "end of CoverTab[1433]"
//line /snap/go/10455/src/strings/strings.go:555
	_go_fuzz_dep_.CoverTab[1434]++
							if len(s) >= maxInt/count {
//line /snap/go/10455/src/strings/strings.go:556
		_go_fuzz_dep_.CoverTab[525316]++
//line /snap/go/10455/src/strings/strings.go:556
		_go_fuzz_dep_.CoverTab[1444]++
								panic("strings: Repeat output length overflow")
//line /snap/go/10455/src/strings/strings.go:557
		// _ = "end of CoverTab[1444]"
	} else {
//line /snap/go/10455/src/strings/strings.go:558
		_go_fuzz_dep_.CoverTab[525317]++
//line /snap/go/10455/src/strings/strings.go:558
		_go_fuzz_dep_.CoverTab[1445]++
//line /snap/go/10455/src/strings/strings.go:558
		// _ = "end of CoverTab[1445]"
//line /snap/go/10455/src/strings/strings.go:558
	}
//line /snap/go/10455/src/strings/strings.go:558
	// _ = "end of CoverTab[1434]"
//line /snap/go/10455/src/strings/strings.go:558
	_go_fuzz_dep_.CoverTab[1435]++
							n := len(s) * count

							if len(s) == 0 {
//line /snap/go/10455/src/strings/strings.go:561
		_go_fuzz_dep_.CoverTab[525318]++
//line /snap/go/10455/src/strings/strings.go:561
		_go_fuzz_dep_.CoverTab[1446]++
								return ""
//line /snap/go/10455/src/strings/strings.go:562
		// _ = "end of CoverTab[1446]"
	} else {
//line /snap/go/10455/src/strings/strings.go:563
		_go_fuzz_dep_.CoverTab[525319]++
//line /snap/go/10455/src/strings/strings.go:563
		_go_fuzz_dep_.CoverTab[1447]++
//line /snap/go/10455/src/strings/strings.go:563
		// _ = "end of CoverTab[1447]"
//line /snap/go/10455/src/strings/strings.go:563
	}
//line /snap/go/10455/src/strings/strings.go:563
	// _ = "end of CoverTab[1435]"
//line /snap/go/10455/src/strings/strings.go:563
	_go_fuzz_dep_.CoverTab[1436]++

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
//line /snap/go/10455/src/strings/strings.go:577
		_go_fuzz_dep_.CoverTab[525320]++
//line /snap/go/10455/src/strings/strings.go:577
		_go_fuzz_dep_.CoverTab[1448]++
								chunkMax = chunkLimit / len(s) * len(s)
								if chunkMax == 0 {
//line /snap/go/10455/src/strings/strings.go:579
			_go_fuzz_dep_.CoverTab[525322]++
//line /snap/go/10455/src/strings/strings.go:579
			_go_fuzz_dep_.CoverTab[1449]++
									chunkMax = len(s)
//line /snap/go/10455/src/strings/strings.go:580
			// _ = "end of CoverTab[1449]"
		} else {
//line /snap/go/10455/src/strings/strings.go:581
			_go_fuzz_dep_.CoverTab[525323]++
//line /snap/go/10455/src/strings/strings.go:581
			_go_fuzz_dep_.CoverTab[1450]++
//line /snap/go/10455/src/strings/strings.go:581
			// _ = "end of CoverTab[1450]"
//line /snap/go/10455/src/strings/strings.go:581
		}
//line /snap/go/10455/src/strings/strings.go:581
		// _ = "end of CoverTab[1448]"
	} else {
//line /snap/go/10455/src/strings/strings.go:582
		_go_fuzz_dep_.CoverTab[525321]++
//line /snap/go/10455/src/strings/strings.go:582
		_go_fuzz_dep_.CoverTab[1451]++
//line /snap/go/10455/src/strings/strings.go:582
		// _ = "end of CoverTab[1451]"
//line /snap/go/10455/src/strings/strings.go:582
	}
//line /snap/go/10455/src/strings/strings.go:582
	// _ = "end of CoverTab[1436]"
//line /snap/go/10455/src/strings/strings.go:582
	_go_fuzz_dep_.CoverTab[1437]++

							var b Builder
							b.Grow(n)
							b.WriteString(s)
//line /snap/go/10455/src/strings/strings.go:586
	_go_fuzz_dep_.CoverTab[786531] = 0
							for b.Len() < n {
//line /snap/go/10455/src/strings/strings.go:587
		if _go_fuzz_dep_.CoverTab[786531] == 0 {
//line /snap/go/10455/src/strings/strings.go:587
			_go_fuzz_dep_.CoverTab[525579]++
//line /snap/go/10455/src/strings/strings.go:587
		} else {
//line /snap/go/10455/src/strings/strings.go:587
			_go_fuzz_dep_.CoverTab[525580]++
//line /snap/go/10455/src/strings/strings.go:587
		}
//line /snap/go/10455/src/strings/strings.go:587
		_go_fuzz_dep_.CoverTab[786531] = 1
//line /snap/go/10455/src/strings/strings.go:587
		_go_fuzz_dep_.CoverTab[1452]++
								chunk := n - b.Len()
								if chunk > b.Len() {
//line /snap/go/10455/src/strings/strings.go:589
			_go_fuzz_dep_.CoverTab[525324]++
//line /snap/go/10455/src/strings/strings.go:589
			_go_fuzz_dep_.CoverTab[1455]++
									chunk = b.Len()
//line /snap/go/10455/src/strings/strings.go:590
			// _ = "end of CoverTab[1455]"
		} else {
//line /snap/go/10455/src/strings/strings.go:591
			_go_fuzz_dep_.CoverTab[525325]++
//line /snap/go/10455/src/strings/strings.go:591
			_go_fuzz_dep_.CoverTab[1456]++
//line /snap/go/10455/src/strings/strings.go:591
			// _ = "end of CoverTab[1456]"
//line /snap/go/10455/src/strings/strings.go:591
		}
//line /snap/go/10455/src/strings/strings.go:591
		// _ = "end of CoverTab[1452]"
//line /snap/go/10455/src/strings/strings.go:591
		_go_fuzz_dep_.CoverTab[1453]++
								if chunk > chunkMax {
//line /snap/go/10455/src/strings/strings.go:592
			_go_fuzz_dep_.CoverTab[525326]++
//line /snap/go/10455/src/strings/strings.go:592
			_go_fuzz_dep_.CoverTab[1457]++
									chunk = chunkMax
//line /snap/go/10455/src/strings/strings.go:593
			// _ = "end of CoverTab[1457]"
		} else {
//line /snap/go/10455/src/strings/strings.go:594
			_go_fuzz_dep_.CoverTab[525327]++
//line /snap/go/10455/src/strings/strings.go:594
			_go_fuzz_dep_.CoverTab[1458]++
//line /snap/go/10455/src/strings/strings.go:594
			// _ = "end of CoverTab[1458]"
//line /snap/go/10455/src/strings/strings.go:594
		}
//line /snap/go/10455/src/strings/strings.go:594
		// _ = "end of CoverTab[1453]"
//line /snap/go/10455/src/strings/strings.go:594
		_go_fuzz_dep_.CoverTab[1454]++
								b.WriteString(b.String()[:chunk])
//line /snap/go/10455/src/strings/strings.go:595
		// _ = "end of CoverTab[1454]"
	}
//line /snap/go/10455/src/strings/strings.go:596
	if _go_fuzz_dep_.CoverTab[786531] == 0 {
//line /snap/go/10455/src/strings/strings.go:596
		_go_fuzz_dep_.CoverTab[525581]++
//line /snap/go/10455/src/strings/strings.go:596
	} else {
//line /snap/go/10455/src/strings/strings.go:596
		_go_fuzz_dep_.CoverTab[525582]++
//line /snap/go/10455/src/strings/strings.go:596
	}
//line /snap/go/10455/src/strings/strings.go:596
	// _ = "end of CoverTab[1437]"
//line /snap/go/10455/src/strings/strings.go:596
	_go_fuzz_dep_.CoverTab[1438]++
							return b.String()
//line /snap/go/10455/src/strings/strings.go:597
	// _ = "end of CoverTab[1438]"
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
//line /snap/go/10455/src/strings/strings.go:601
	_go_fuzz_dep_.CoverTab[1459]++
							isASCII, hasLower := true, false
//line /snap/go/10455/src/strings/strings.go:602
	_go_fuzz_dep_.CoverTab[786532] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/strings.go:603
		if _go_fuzz_dep_.CoverTab[786532] == 0 {
//line /snap/go/10455/src/strings/strings.go:603
			_go_fuzz_dep_.CoverTab[525583]++
//line /snap/go/10455/src/strings/strings.go:603
		} else {
//line /snap/go/10455/src/strings/strings.go:603
			_go_fuzz_dep_.CoverTab[525584]++
//line /snap/go/10455/src/strings/strings.go:603
		}
//line /snap/go/10455/src/strings/strings.go:603
		_go_fuzz_dep_.CoverTab[786532] = 1
//line /snap/go/10455/src/strings/strings.go:603
		_go_fuzz_dep_.CoverTab[1462]++
								c := s[i]
								if c >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:605
			_go_fuzz_dep_.CoverTab[525328]++
//line /snap/go/10455/src/strings/strings.go:605
			_go_fuzz_dep_.CoverTab[1464]++
									isASCII = false
									break
//line /snap/go/10455/src/strings/strings.go:607
			// _ = "end of CoverTab[1464]"
		} else {
//line /snap/go/10455/src/strings/strings.go:608
			_go_fuzz_dep_.CoverTab[525329]++
//line /snap/go/10455/src/strings/strings.go:608
			_go_fuzz_dep_.CoverTab[1465]++
//line /snap/go/10455/src/strings/strings.go:608
			// _ = "end of CoverTab[1465]"
//line /snap/go/10455/src/strings/strings.go:608
		}
//line /snap/go/10455/src/strings/strings.go:608
		// _ = "end of CoverTab[1462]"
//line /snap/go/10455/src/strings/strings.go:608
		_go_fuzz_dep_.CoverTab[1463]++
								hasLower = hasLower || func() bool {
//line /snap/go/10455/src/strings/strings.go:609
			_go_fuzz_dep_.CoverTab[1466]++
//line /snap/go/10455/src/strings/strings.go:609
			return ('a' <= c && func() bool {
//line /snap/go/10455/src/strings/strings.go:609
				_go_fuzz_dep_.CoverTab[1467]++
//line /snap/go/10455/src/strings/strings.go:609
				return c <= 'z'
//line /snap/go/10455/src/strings/strings.go:609
				// _ = "end of CoverTab[1467]"
//line /snap/go/10455/src/strings/strings.go:609
			}())
//line /snap/go/10455/src/strings/strings.go:609
			// _ = "end of CoverTab[1466]"
//line /snap/go/10455/src/strings/strings.go:609
		}()
//line /snap/go/10455/src/strings/strings.go:609
		// _ = "end of CoverTab[1463]"
	}
//line /snap/go/10455/src/strings/strings.go:610
	if _go_fuzz_dep_.CoverTab[786532] == 0 {
//line /snap/go/10455/src/strings/strings.go:610
		_go_fuzz_dep_.CoverTab[525585]++
//line /snap/go/10455/src/strings/strings.go:610
	} else {
//line /snap/go/10455/src/strings/strings.go:610
		_go_fuzz_dep_.CoverTab[525586]++
//line /snap/go/10455/src/strings/strings.go:610
	}
//line /snap/go/10455/src/strings/strings.go:610
	// _ = "end of CoverTab[1459]"
//line /snap/go/10455/src/strings/strings.go:610
	_go_fuzz_dep_.CoverTab[1460]++

							if isASCII {
//line /snap/go/10455/src/strings/strings.go:612
		_go_fuzz_dep_.CoverTab[525330]++
//line /snap/go/10455/src/strings/strings.go:612
		_go_fuzz_dep_.CoverTab[1468]++
								if !hasLower {
//line /snap/go/10455/src/strings/strings.go:613
			_go_fuzz_dep_.CoverTab[525332]++
//line /snap/go/10455/src/strings/strings.go:613
			_go_fuzz_dep_.CoverTab[1472]++
									return s
//line /snap/go/10455/src/strings/strings.go:614
			// _ = "end of CoverTab[1472]"
		} else {
//line /snap/go/10455/src/strings/strings.go:615
			_go_fuzz_dep_.CoverTab[525333]++
//line /snap/go/10455/src/strings/strings.go:615
			_go_fuzz_dep_.CoverTab[1473]++
//line /snap/go/10455/src/strings/strings.go:615
			// _ = "end of CoverTab[1473]"
//line /snap/go/10455/src/strings/strings.go:615
		}
//line /snap/go/10455/src/strings/strings.go:615
		// _ = "end of CoverTab[1468]"
//line /snap/go/10455/src/strings/strings.go:615
		_go_fuzz_dep_.CoverTab[1469]++
								var (
			b	Builder
			pos	int
		)
								b.Grow(len(s))
//line /snap/go/10455/src/strings/strings.go:620
		_go_fuzz_dep_.CoverTab[786533] = 0
								for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/strings.go:621
			if _go_fuzz_dep_.CoverTab[786533] == 0 {
//line /snap/go/10455/src/strings/strings.go:621
				_go_fuzz_dep_.CoverTab[525587]++
//line /snap/go/10455/src/strings/strings.go:621
			} else {
//line /snap/go/10455/src/strings/strings.go:621
				_go_fuzz_dep_.CoverTab[525588]++
//line /snap/go/10455/src/strings/strings.go:621
			}
//line /snap/go/10455/src/strings/strings.go:621
			_go_fuzz_dep_.CoverTab[786533] = 1
//line /snap/go/10455/src/strings/strings.go:621
			_go_fuzz_dep_.CoverTab[1474]++
									c := s[i]
									if 'a' <= c && func() bool {
//line /snap/go/10455/src/strings/strings.go:623
				_go_fuzz_dep_.CoverTab[1475]++
//line /snap/go/10455/src/strings/strings.go:623
				return c <= 'z'
//line /snap/go/10455/src/strings/strings.go:623
				// _ = "end of CoverTab[1475]"
//line /snap/go/10455/src/strings/strings.go:623
			}() {
//line /snap/go/10455/src/strings/strings.go:623
				_go_fuzz_dep_.CoverTab[525334]++
//line /snap/go/10455/src/strings/strings.go:623
				_go_fuzz_dep_.CoverTab[1476]++
										c -= 'a' - 'A'
										if pos < i {
//line /snap/go/10455/src/strings/strings.go:625
					_go_fuzz_dep_.CoverTab[525336]++
//line /snap/go/10455/src/strings/strings.go:625
					_go_fuzz_dep_.CoverTab[1478]++
											b.WriteString(s[pos:i])
//line /snap/go/10455/src/strings/strings.go:626
					// _ = "end of CoverTab[1478]"
				} else {
//line /snap/go/10455/src/strings/strings.go:627
					_go_fuzz_dep_.CoverTab[525337]++
//line /snap/go/10455/src/strings/strings.go:627
					_go_fuzz_dep_.CoverTab[1479]++
//line /snap/go/10455/src/strings/strings.go:627
					// _ = "end of CoverTab[1479]"
//line /snap/go/10455/src/strings/strings.go:627
				}
//line /snap/go/10455/src/strings/strings.go:627
				// _ = "end of CoverTab[1476]"
//line /snap/go/10455/src/strings/strings.go:627
				_go_fuzz_dep_.CoverTab[1477]++
										b.WriteByte(c)
										pos = i + 1
//line /snap/go/10455/src/strings/strings.go:629
				// _ = "end of CoverTab[1477]"
			} else {
//line /snap/go/10455/src/strings/strings.go:630
				_go_fuzz_dep_.CoverTab[525335]++
//line /snap/go/10455/src/strings/strings.go:630
				_go_fuzz_dep_.CoverTab[1480]++
//line /snap/go/10455/src/strings/strings.go:630
				// _ = "end of CoverTab[1480]"
//line /snap/go/10455/src/strings/strings.go:630
			}
//line /snap/go/10455/src/strings/strings.go:630
			// _ = "end of CoverTab[1474]"
		}
//line /snap/go/10455/src/strings/strings.go:631
		if _go_fuzz_dep_.CoverTab[786533] == 0 {
//line /snap/go/10455/src/strings/strings.go:631
			_go_fuzz_dep_.CoverTab[525589]++
//line /snap/go/10455/src/strings/strings.go:631
		} else {
//line /snap/go/10455/src/strings/strings.go:631
			_go_fuzz_dep_.CoverTab[525590]++
//line /snap/go/10455/src/strings/strings.go:631
		}
//line /snap/go/10455/src/strings/strings.go:631
		// _ = "end of CoverTab[1469]"
//line /snap/go/10455/src/strings/strings.go:631
		_go_fuzz_dep_.CoverTab[1470]++
								if pos < len(s) {
//line /snap/go/10455/src/strings/strings.go:632
			_go_fuzz_dep_.CoverTab[525338]++
//line /snap/go/10455/src/strings/strings.go:632
			_go_fuzz_dep_.CoverTab[1481]++
									b.WriteString(s[pos:])
//line /snap/go/10455/src/strings/strings.go:633
			// _ = "end of CoverTab[1481]"
		} else {
//line /snap/go/10455/src/strings/strings.go:634
			_go_fuzz_dep_.CoverTab[525339]++
//line /snap/go/10455/src/strings/strings.go:634
			_go_fuzz_dep_.CoverTab[1482]++
//line /snap/go/10455/src/strings/strings.go:634
			// _ = "end of CoverTab[1482]"
//line /snap/go/10455/src/strings/strings.go:634
		}
//line /snap/go/10455/src/strings/strings.go:634
		// _ = "end of CoverTab[1470]"
//line /snap/go/10455/src/strings/strings.go:634
		_go_fuzz_dep_.CoverTab[1471]++
								return b.String()
//line /snap/go/10455/src/strings/strings.go:635
		// _ = "end of CoverTab[1471]"
	} else {
//line /snap/go/10455/src/strings/strings.go:636
		_go_fuzz_dep_.CoverTab[525331]++
//line /snap/go/10455/src/strings/strings.go:636
		_go_fuzz_dep_.CoverTab[1483]++
//line /snap/go/10455/src/strings/strings.go:636
		// _ = "end of CoverTab[1483]"
//line /snap/go/10455/src/strings/strings.go:636
	}
//line /snap/go/10455/src/strings/strings.go:636
	// _ = "end of CoverTab[1460]"
//line /snap/go/10455/src/strings/strings.go:636
	_go_fuzz_dep_.CoverTab[1461]++
							return Map(unicode.ToUpper, s)
//line /snap/go/10455/src/strings/strings.go:637
	// _ = "end of CoverTab[1461]"
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
//line /snap/go/10455/src/strings/strings.go:641
	_go_fuzz_dep_.CoverTab[1484]++
							isASCII, hasUpper := true, false
//line /snap/go/10455/src/strings/strings.go:642
	_go_fuzz_dep_.CoverTab[786534] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/strings.go:643
		if _go_fuzz_dep_.CoverTab[786534] == 0 {
//line /snap/go/10455/src/strings/strings.go:643
			_go_fuzz_dep_.CoverTab[525591]++
//line /snap/go/10455/src/strings/strings.go:643
		} else {
//line /snap/go/10455/src/strings/strings.go:643
			_go_fuzz_dep_.CoverTab[525592]++
//line /snap/go/10455/src/strings/strings.go:643
		}
//line /snap/go/10455/src/strings/strings.go:643
		_go_fuzz_dep_.CoverTab[786534] = 1
//line /snap/go/10455/src/strings/strings.go:643
		_go_fuzz_dep_.CoverTab[1487]++
								c := s[i]
								if c >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:645
			_go_fuzz_dep_.CoverTab[525340]++
//line /snap/go/10455/src/strings/strings.go:645
			_go_fuzz_dep_.CoverTab[1489]++
									isASCII = false
									break
//line /snap/go/10455/src/strings/strings.go:647
			// _ = "end of CoverTab[1489]"
		} else {
//line /snap/go/10455/src/strings/strings.go:648
			_go_fuzz_dep_.CoverTab[525341]++
//line /snap/go/10455/src/strings/strings.go:648
			_go_fuzz_dep_.CoverTab[1490]++
//line /snap/go/10455/src/strings/strings.go:648
			// _ = "end of CoverTab[1490]"
//line /snap/go/10455/src/strings/strings.go:648
		}
//line /snap/go/10455/src/strings/strings.go:648
		// _ = "end of CoverTab[1487]"
//line /snap/go/10455/src/strings/strings.go:648
		_go_fuzz_dep_.CoverTab[1488]++
								hasUpper = hasUpper || func() bool {
//line /snap/go/10455/src/strings/strings.go:649
			_go_fuzz_dep_.CoverTab[1491]++
//line /snap/go/10455/src/strings/strings.go:649
			return ('A' <= c && func() bool {
//line /snap/go/10455/src/strings/strings.go:649
				_go_fuzz_dep_.CoverTab[1492]++
//line /snap/go/10455/src/strings/strings.go:649
				return c <= 'Z'
//line /snap/go/10455/src/strings/strings.go:649
				// _ = "end of CoverTab[1492]"
//line /snap/go/10455/src/strings/strings.go:649
			}())
//line /snap/go/10455/src/strings/strings.go:649
			// _ = "end of CoverTab[1491]"
//line /snap/go/10455/src/strings/strings.go:649
		}()
//line /snap/go/10455/src/strings/strings.go:649
		// _ = "end of CoverTab[1488]"
	}
//line /snap/go/10455/src/strings/strings.go:650
	if _go_fuzz_dep_.CoverTab[786534] == 0 {
//line /snap/go/10455/src/strings/strings.go:650
		_go_fuzz_dep_.CoverTab[525593]++
//line /snap/go/10455/src/strings/strings.go:650
	} else {
//line /snap/go/10455/src/strings/strings.go:650
		_go_fuzz_dep_.CoverTab[525594]++
//line /snap/go/10455/src/strings/strings.go:650
	}
//line /snap/go/10455/src/strings/strings.go:650
	// _ = "end of CoverTab[1484]"
//line /snap/go/10455/src/strings/strings.go:650
	_go_fuzz_dep_.CoverTab[1485]++

							if isASCII {
//line /snap/go/10455/src/strings/strings.go:652
		_go_fuzz_dep_.CoverTab[525342]++
//line /snap/go/10455/src/strings/strings.go:652
		_go_fuzz_dep_.CoverTab[1493]++
								if !hasUpper {
//line /snap/go/10455/src/strings/strings.go:653
			_go_fuzz_dep_.CoverTab[525344]++
//line /snap/go/10455/src/strings/strings.go:653
			_go_fuzz_dep_.CoverTab[1497]++
									return s
//line /snap/go/10455/src/strings/strings.go:654
			// _ = "end of CoverTab[1497]"
		} else {
//line /snap/go/10455/src/strings/strings.go:655
			_go_fuzz_dep_.CoverTab[525345]++
//line /snap/go/10455/src/strings/strings.go:655
			_go_fuzz_dep_.CoverTab[1498]++
//line /snap/go/10455/src/strings/strings.go:655
			// _ = "end of CoverTab[1498]"
//line /snap/go/10455/src/strings/strings.go:655
		}
//line /snap/go/10455/src/strings/strings.go:655
		// _ = "end of CoverTab[1493]"
//line /snap/go/10455/src/strings/strings.go:655
		_go_fuzz_dep_.CoverTab[1494]++
								var (
			b	Builder
			pos	int
		)
								b.Grow(len(s))
//line /snap/go/10455/src/strings/strings.go:660
		_go_fuzz_dep_.CoverTab[786535] = 0
								for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/strings.go:661
			if _go_fuzz_dep_.CoverTab[786535] == 0 {
//line /snap/go/10455/src/strings/strings.go:661
				_go_fuzz_dep_.CoverTab[525595]++
//line /snap/go/10455/src/strings/strings.go:661
			} else {
//line /snap/go/10455/src/strings/strings.go:661
				_go_fuzz_dep_.CoverTab[525596]++
//line /snap/go/10455/src/strings/strings.go:661
			}
//line /snap/go/10455/src/strings/strings.go:661
			_go_fuzz_dep_.CoverTab[786535] = 1
//line /snap/go/10455/src/strings/strings.go:661
			_go_fuzz_dep_.CoverTab[1499]++
									c := s[i]
									if 'A' <= c && func() bool {
//line /snap/go/10455/src/strings/strings.go:663
				_go_fuzz_dep_.CoverTab[1500]++
//line /snap/go/10455/src/strings/strings.go:663
				return c <= 'Z'
//line /snap/go/10455/src/strings/strings.go:663
				// _ = "end of CoverTab[1500]"
//line /snap/go/10455/src/strings/strings.go:663
			}() {
//line /snap/go/10455/src/strings/strings.go:663
				_go_fuzz_dep_.CoverTab[525346]++
//line /snap/go/10455/src/strings/strings.go:663
				_go_fuzz_dep_.CoverTab[1501]++
										c += 'a' - 'A'
										if pos < i {
//line /snap/go/10455/src/strings/strings.go:665
					_go_fuzz_dep_.CoverTab[525348]++
//line /snap/go/10455/src/strings/strings.go:665
					_go_fuzz_dep_.CoverTab[1503]++
											b.WriteString(s[pos:i])
//line /snap/go/10455/src/strings/strings.go:666
					// _ = "end of CoverTab[1503]"
				} else {
//line /snap/go/10455/src/strings/strings.go:667
					_go_fuzz_dep_.CoverTab[525349]++
//line /snap/go/10455/src/strings/strings.go:667
					_go_fuzz_dep_.CoverTab[1504]++
//line /snap/go/10455/src/strings/strings.go:667
					// _ = "end of CoverTab[1504]"
//line /snap/go/10455/src/strings/strings.go:667
				}
//line /snap/go/10455/src/strings/strings.go:667
				// _ = "end of CoverTab[1501]"
//line /snap/go/10455/src/strings/strings.go:667
				_go_fuzz_dep_.CoverTab[1502]++
										b.WriteByte(c)
										pos = i + 1
//line /snap/go/10455/src/strings/strings.go:669
				// _ = "end of CoverTab[1502]"
			} else {
//line /snap/go/10455/src/strings/strings.go:670
				_go_fuzz_dep_.CoverTab[525347]++
//line /snap/go/10455/src/strings/strings.go:670
				_go_fuzz_dep_.CoverTab[1505]++
//line /snap/go/10455/src/strings/strings.go:670
				// _ = "end of CoverTab[1505]"
//line /snap/go/10455/src/strings/strings.go:670
			}
//line /snap/go/10455/src/strings/strings.go:670
			// _ = "end of CoverTab[1499]"
		}
//line /snap/go/10455/src/strings/strings.go:671
		if _go_fuzz_dep_.CoverTab[786535] == 0 {
//line /snap/go/10455/src/strings/strings.go:671
			_go_fuzz_dep_.CoverTab[525597]++
//line /snap/go/10455/src/strings/strings.go:671
		} else {
//line /snap/go/10455/src/strings/strings.go:671
			_go_fuzz_dep_.CoverTab[525598]++
//line /snap/go/10455/src/strings/strings.go:671
		}
//line /snap/go/10455/src/strings/strings.go:671
		// _ = "end of CoverTab[1494]"
//line /snap/go/10455/src/strings/strings.go:671
		_go_fuzz_dep_.CoverTab[1495]++
								if pos < len(s) {
//line /snap/go/10455/src/strings/strings.go:672
			_go_fuzz_dep_.CoverTab[525350]++
//line /snap/go/10455/src/strings/strings.go:672
			_go_fuzz_dep_.CoverTab[1506]++
									b.WriteString(s[pos:])
//line /snap/go/10455/src/strings/strings.go:673
			// _ = "end of CoverTab[1506]"
		} else {
//line /snap/go/10455/src/strings/strings.go:674
			_go_fuzz_dep_.CoverTab[525351]++
//line /snap/go/10455/src/strings/strings.go:674
			_go_fuzz_dep_.CoverTab[1507]++
//line /snap/go/10455/src/strings/strings.go:674
			// _ = "end of CoverTab[1507]"
//line /snap/go/10455/src/strings/strings.go:674
		}
//line /snap/go/10455/src/strings/strings.go:674
		// _ = "end of CoverTab[1495]"
//line /snap/go/10455/src/strings/strings.go:674
		_go_fuzz_dep_.CoverTab[1496]++
								return b.String()
//line /snap/go/10455/src/strings/strings.go:675
		// _ = "end of CoverTab[1496]"
	} else {
//line /snap/go/10455/src/strings/strings.go:676
		_go_fuzz_dep_.CoverTab[525343]++
//line /snap/go/10455/src/strings/strings.go:676
		_go_fuzz_dep_.CoverTab[1508]++
//line /snap/go/10455/src/strings/strings.go:676
		// _ = "end of CoverTab[1508]"
//line /snap/go/10455/src/strings/strings.go:676
	}
//line /snap/go/10455/src/strings/strings.go:676
	// _ = "end of CoverTab[1485]"
//line /snap/go/10455/src/strings/strings.go:676
	_go_fuzz_dep_.CoverTab[1486]++
							return Map(unicode.ToLower, s)
//line /snap/go/10455/src/strings/strings.go:677
	// _ = "end of CoverTab[1486]"
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to
//line /snap/go/10455/src/strings/strings.go:680
// their Unicode title case.
//line /snap/go/10455/src/strings/strings.go:682
func ToTitle(s string) string {
//line /snap/go/10455/src/strings/strings.go:682
	_go_fuzz_dep_.CoverTab[1509]++
//line /snap/go/10455/src/strings/strings.go:682
	return Map(unicode.ToTitle, s)
//line /snap/go/10455/src/strings/strings.go:682
	// _ = "end of CoverTab[1509]"
//line /snap/go/10455/src/strings/strings.go:682
}

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
//line /snap/go/10455/src/strings/strings.go:684
// upper case using the case mapping specified by c.
//line /snap/go/10455/src/strings/strings.go:686
func ToUpperSpecial(c unicode.SpecialCase, s string) string {
//line /snap/go/10455/src/strings/strings.go:686
	_go_fuzz_dep_.CoverTab[1510]++
							return Map(c.ToUpper, s)
//line /snap/go/10455/src/strings/strings.go:687
	// _ = "end of CoverTab[1510]"
}

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
//line /snap/go/10455/src/strings/strings.go:690
// lower case using the case mapping specified by c.
//line /snap/go/10455/src/strings/strings.go:692
func ToLowerSpecial(c unicode.SpecialCase, s string) string {
//line /snap/go/10455/src/strings/strings.go:692
	_go_fuzz_dep_.CoverTab[1511]++
							return Map(c.ToLower, s)
//line /snap/go/10455/src/strings/strings.go:693
	// _ = "end of CoverTab[1511]"
}

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
//line /snap/go/10455/src/strings/strings.go:696
// Unicode title case, giving priority to the special casing rules.
//line /snap/go/10455/src/strings/strings.go:698
func ToTitleSpecial(c unicode.SpecialCase, s string) string {
//line /snap/go/10455/src/strings/strings.go:698
	_go_fuzz_dep_.CoverTab[1512]++
							return Map(c.ToTitle, s)
//line /snap/go/10455/src/strings/strings.go:699
	// _ = "end of CoverTab[1512]"
}

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
//line /snap/go/10455/src/strings/strings.go:702
// replaced by the replacement string, which may be empty.
//line /snap/go/10455/src/strings/strings.go:704
func ToValidUTF8(s, replacement string) string {
//line /snap/go/10455/src/strings/strings.go:704
	_go_fuzz_dep_.CoverTab[1513]++
							var b Builder
//line /snap/go/10455/src/strings/strings.go:705
	_go_fuzz_dep_.CoverTab[786536] = 0

							for i, c := range s {
//line /snap/go/10455/src/strings/strings.go:707
		if _go_fuzz_dep_.CoverTab[786536] == 0 {
//line /snap/go/10455/src/strings/strings.go:707
			_go_fuzz_dep_.CoverTab[525599]++
//line /snap/go/10455/src/strings/strings.go:707
		} else {
//line /snap/go/10455/src/strings/strings.go:707
			_go_fuzz_dep_.CoverTab[525600]++
//line /snap/go/10455/src/strings/strings.go:707
		}
//line /snap/go/10455/src/strings/strings.go:707
		_go_fuzz_dep_.CoverTab[786536] = 1
//line /snap/go/10455/src/strings/strings.go:707
		_go_fuzz_dep_.CoverTab[1517]++
								if c != utf8.RuneError {
//line /snap/go/10455/src/strings/strings.go:708
			_go_fuzz_dep_.CoverTab[525352]++
//line /snap/go/10455/src/strings/strings.go:708
			_go_fuzz_dep_.CoverTab[1519]++
									continue
//line /snap/go/10455/src/strings/strings.go:709
			// _ = "end of CoverTab[1519]"
		} else {
//line /snap/go/10455/src/strings/strings.go:710
			_go_fuzz_dep_.CoverTab[525353]++
//line /snap/go/10455/src/strings/strings.go:710
			_go_fuzz_dep_.CoverTab[1520]++
//line /snap/go/10455/src/strings/strings.go:710
			// _ = "end of CoverTab[1520]"
//line /snap/go/10455/src/strings/strings.go:710
		}
//line /snap/go/10455/src/strings/strings.go:710
		// _ = "end of CoverTab[1517]"
//line /snap/go/10455/src/strings/strings.go:710
		_go_fuzz_dep_.CoverTab[1518]++

								_, wid := utf8.DecodeRuneInString(s[i:])
								if wid == 1 {
//line /snap/go/10455/src/strings/strings.go:713
			_go_fuzz_dep_.CoverTab[525354]++
//line /snap/go/10455/src/strings/strings.go:713
			_go_fuzz_dep_.CoverTab[1521]++
									b.Grow(len(s) + len(replacement))
									b.WriteString(s[:i])
									s = s[i:]
									break
//line /snap/go/10455/src/strings/strings.go:717
			// _ = "end of CoverTab[1521]"
		} else {
//line /snap/go/10455/src/strings/strings.go:718
			_go_fuzz_dep_.CoverTab[525355]++
//line /snap/go/10455/src/strings/strings.go:718
			_go_fuzz_dep_.CoverTab[1522]++
//line /snap/go/10455/src/strings/strings.go:718
			// _ = "end of CoverTab[1522]"
//line /snap/go/10455/src/strings/strings.go:718
		}
//line /snap/go/10455/src/strings/strings.go:718
		// _ = "end of CoverTab[1518]"
	}
//line /snap/go/10455/src/strings/strings.go:719
	if _go_fuzz_dep_.CoverTab[786536] == 0 {
//line /snap/go/10455/src/strings/strings.go:719
		_go_fuzz_dep_.CoverTab[525601]++
//line /snap/go/10455/src/strings/strings.go:719
	} else {
//line /snap/go/10455/src/strings/strings.go:719
		_go_fuzz_dep_.CoverTab[525602]++
//line /snap/go/10455/src/strings/strings.go:719
	}
//line /snap/go/10455/src/strings/strings.go:719
	// _ = "end of CoverTab[1513]"
//line /snap/go/10455/src/strings/strings.go:719
	_go_fuzz_dep_.CoverTab[1514]++

//line /snap/go/10455/src/strings/strings.go:722
	if b.Cap() == 0 {
//line /snap/go/10455/src/strings/strings.go:722
		_go_fuzz_dep_.CoverTab[525356]++
//line /snap/go/10455/src/strings/strings.go:722
		_go_fuzz_dep_.CoverTab[1523]++
								return s
//line /snap/go/10455/src/strings/strings.go:723
		// _ = "end of CoverTab[1523]"
	} else {
//line /snap/go/10455/src/strings/strings.go:724
		_go_fuzz_dep_.CoverTab[525357]++
//line /snap/go/10455/src/strings/strings.go:724
		_go_fuzz_dep_.CoverTab[1524]++
//line /snap/go/10455/src/strings/strings.go:724
		// _ = "end of CoverTab[1524]"
//line /snap/go/10455/src/strings/strings.go:724
	}
//line /snap/go/10455/src/strings/strings.go:724
	// _ = "end of CoverTab[1514]"
//line /snap/go/10455/src/strings/strings.go:724
	_go_fuzz_dep_.CoverTab[1515]++

							invalid := false
//line /snap/go/10455/src/strings/strings.go:726
	_go_fuzz_dep_.CoverTab[786537] = 0
							for i := 0; i < len(s); {
//line /snap/go/10455/src/strings/strings.go:727
		if _go_fuzz_dep_.CoverTab[786537] == 0 {
//line /snap/go/10455/src/strings/strings.go:727
			_go_fuzz_dep_.CoverTab[525603]++
//line /snap/go/10455/src/strings/strings.go:727
		} else {
//line /snap/go/10455/src/strings/strings.go:727
			_go_fuzz_dep_.CoverTab[525604]++
//line /snap/go/10455/src/strings/strings.go:727
		}
//line /snap/go/10455/src/strings/strings.go:727
		_go_fuzz_dep_.CoverTab[786537] = 1
//line /snap/go/10455/src/strings/strings.go:727
		_go_fuzz_dep_.CoverTab[1525]++
								c := s[i]
								if c < utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:729
			_go_fuzz_dep_.CoverTab[525358]++
//line /snap/go/10455/src/strings/strings.go:729
			_go_fuzz_dep_.CoverTab[1528]++
									i++
									invalid = false
									b.WriteByte(c)
									continue
//line /snap/go/10455/src/strings/strings.go:733
			// _ = "end of CoverTab[1528]"
		} else {
//line /snap/go/10455/src/strings/strings.go:734
			_go_fuzz_dep_.CoverTab[525359]++
//line /snap/go/10455/src/strings/strings.go:734
			_go_fuzz_dep_.CoverTab[1529]++
//line /snap/go/10455/src/strings/strings.go:734
			// _ = "end of CoverTab[1529]"
//line /snap/go/10455/src/strings/strings.go:734
		}
//line /snap/go/10455/src/strings/strings.go:734
		// _ = "end of CoverTab[1525]"
//line /snap/go/10455/src/strings/strings.go:734
		_go_fuzz_dep_.CoverTab[1526]++
								_, wid := utf8.DecodeRuneInString(s[i:])
								if wid == 1 {
//line /snap/go/10455/src/strings/strings.go:736
			_go_fuzz_dep_.CoverTab[525360]++
//line /snap/go/10455/src/strings/strings.go:736
			_go_fuzz_dep_.CoverTab[1530]++
									i++
									if !invalid {
//line /snap/go/10455/src/strings/strings.go:738
				_go_fuzz_dep_.CoverTab[525362]++
//line /snap/go/10455/src/strings/strings.go:738
				_go_fuzz_dep_.CoverTab[1532]++
										invalid = true
										b.WriteString(replacement)
//line /snap/go/10455/src/strings/strings.go:740
				// _ = "end of CoverTab[1532]"
			} else {
//line /snap/go/10455/src/strings/strings.go:741
				_go_fuzz_dep_.CoverTab[525363]++
//line /snap/go/10455/src/strings/strings.go:741
				_go_fuzz_dep_.CoverTab[1533]++
//line /snap/go/10455/src/strings/strings.go:741
				// _ = "end of CoverTab[1533]"
//line /snap/go/10455/src/strings/strings.go:741
			}
//line /snap/go/10455/src/strings/strings.go:741
			// _ = "end of CoverTab[1530]"
//line /snap/go/10455/src/strings/strings.go:741
			_go_fuzz_dep_.CoverTab[1531]++
									continue
//line /snap/go/10455/src/strings/strings.go:742
			// _ = "end of CoverTab[1531]"
		} else {
//line /snap/go/10455/src/strings/strings.go:743
			_go_fuzz_dep_.CoverTab[525361]++
//line /snap/go/10455/src/strings/strings.go:743
			_go_fuzz_dep_.CoverTab[1534]++
//line /snap/go/10455/src/strings/strings.go:743
			// _ = "end of CoverTab[1534]"
//line /snap/go/10455/src/strings/strings.go:743
		}
//line /snap/go/10455/src/strings/strings.go:743
		// _ = "end of CoverTab[1526]"
//line /snap/go/10455/src/strings/strings.go:743
		_go_fuzz_dep_.CoverTab[1527]++
								invalid = false
								b.WriteString(s[i : i+wid])
								i += wid
//line /snap/go/10455/src/strings/strings.go:746
		// _ = "end of CoverTab[1527]"
	}
//line /snap/go/10455/src/strings/strings.go:747
	if _go_fuzz_dep_.CoverTab[786537] == 0 {
//line /snap/go/10455/src/strings/strings.go:747
		_go_fuzz_dep_.CoverTab[525605]++
//line /snap/go/10455/src/strings/strings.go:747
	} else {
//line /snap/go/10455/src/strings/strings.go:747
		_go_fuzz_dep_.CoverTab[525606]++
//line /snap/go/10455/src/strings/strings.go:747
	}
//line /snap/go/10455/src/strings/strings.go:747
	// _ = "end of CoverTab[1515]"
//line /snap/go/10455/src/strings/strings.go:747
	_go_fuzz_dep_.CoverTab[1516]++

							return b.String()
//line /snap/go/10455/src/strings/strings.go:749
	// _ = "end of CoverTab[1516]"
}

// isSeparator reports whether the rune could mark a word boundary.
//line /snap/go/10455/src/strings/strings.go:752
// TODO: update when package unicode captures more of the properties.
//line /snap/go/10455/src/strings/strings.go:754
func isSeparator(r rune) bool {
//line /snap/go/10455/src/strings/strings.go:754
	_go_fuzz_dep_.CoverTab[1535]++

							if r <= 0x7F {
//line /snap/go/10455/src/strings/strings.go:756
		_go_fuzz_dep_.CoverTab[525364]++
//line /snap/go/10455/src/strings/strings.go:756
		_go_fuzz_dep_.CoverTab[1538]++
								switch {
		case '0' <= r && func() bool {
//line /snap/go/10455/src/strings/strings.go:758
			_go_fuzz_dep_.CoverTab[1545]++
//line /snap/go/10455/src/strings/strings.go:758
			return r <= '9'
//line /snap/go/10455/src/strings/strings.go:758
			// _ = "end of CoverTab[1545]"
//line /snap/go/10455/src/strings/strings.go:758
		}():
//line /snap/go/10455/src/strings/strings.go:758
			_go_fuzz_dep_.CoverTab[525366]++
//line /snap/go/10455/src/strings/strings.go:758
			_go_fuzz_dep_.CoverTab[1540]++
									return false
//line /snap/go/10455/src/strings/strings.go:759
			// _ = "end of CoverTab[1540]"
		case 'a' <= r && func() bool {
//line /snap/go/10455/src/strings/strings.go:760
			_go_fuzz_dep_.CoverTab[1546]++
//line /snap/go/10455/src/strings/strings.go:760
			return r <= 'z'
//line /snap/go/10455/src/strings/strings.go:760
			// _ = "end of CoverTab[1546]"
//line /snap/go/10455/src/strings/strings.go:760
		}():
//line /snap/go/10455/src/strings/strings.go:760
			_go_fuzz_dep_.CoverTab[525367]++
//line /snap/go/10455/src/strings/strings.go:760
			_go_fuzz_dep_.CoverTab[1541]++
									return false
//line /snap/go/10455/src/strings/strings.go:761
			// _ = "end of CoverTab[1541]"
		case 'A' <= r && func() bool {
//line /snap/go/10455/src/strings/strings.go:762
			_go_fuzz_dep_.CoverTab[1547]++
//line /snap/go/10455/src/strings/strings.go:762
			return r <= 'Z'
//line /snap/go/10455/src/strings/strings.go:762
			// _ = "end of CoverTab[1547]"
//line /snap/go/10455/src/strings/strings.go:762
		}():
//line /snap/go/10455/src/strings/strings.go:762
			_go_fuzz_dep_.CoverTab[525368]++
//line /snap/go/10455/src/strings/strings.go:762
			_go_fuzz_dep_.CoverTab[1542]++
									return false
//line /snap/go/10455/src/strings/strings.go:763
			// _ = "end of CoverTab[1542]"
		case r == '_':
//line /snap/go/10455/src/strings/strings.go:764
			_go_fuzz_dep_.CoverTab[525369]++
//line /snap/go/10455/src/strings/strings.go:764
			_go_fuzz_dep_.CoverTab[1543]++
									return false
//line /snap/go/10455/src/strings/strings.go:765
			// _ = "end of CoverTab[1543]"
//line /snap/go/10455/src/strings/strings.go:765
		default:
//line /snap/go/10455/src/strings/strings.go:765
			_go_fuzz_dep_.CoverTab[525370]++
//line /snap/go/10455/src/strings/strings.go:765
			_go_fuzz_dep_.CoverTab[1544]++
//line /snap/go/10455/src/strings/strings.go:765
			// _ = "end of CoverTab[1544]"
		}
//line /snap/go/10455/src/strings/strings.go:766
		// _ = "end of CoverTab[1538]"
//line /snap/go/10455/src/strings/strings.go:766
		_go_fuzz_dep_.CoverTab[1539]++
								return true
//line /snap/go/10455/src/strings/strings.go:767
		// _ = "end of CoverTab[1539]"
	} else {
//line /snap/go/10455/src/strings/strings.go:768
		_go_fuzz_dep_.CoverTab[525365]++
//line /snap/go/10455/src/strings/strings.go:768
		_go_fuzz_dep_.CoverTab[1548]++
//line /snap/go/10455/src/strings/strings.go:768
		// _ = "end of CoverTab[1548]"
//line /snap/go/10455/src/strings/strings.go:768
	}
//line /snap/go/10455/src/strings/strings.go:768
	// _ = "end of CoverTab[1535]"
//line /snap/go/10455/src/strings/strings.go:768
	_go_fuzz_dep_.CoverTab[1536]++

							if unicode.IsLetter(r) || func() bool {
//line /snap/go/10455/src/strings/strings.go:770
		_go_fuzz_dep_.CoverTab[1549]++
//line /snap/go/10455/src/strings/strings.go:770
		return unicode.IsDigit(r)
//line /snap/go/10455/src/strings/strings.go:770
		// _ = "end of CoverTab[1549]"
//line /snap/go/10455/src/strings/strings.go:770
	}() {
//line /snap/go/10455/src/strings/strings.go:770
		_go_fuzz_dep_.CoverTab[525371]++
//line /snap/go/10455/src/strings/strings.go:770
		_go_fuzz_dep_.CoverTab[1550]++
								return false
//line /snap/go/10455/src/strings/strings.go:771
		// _ = "end of CoverTab[1550]"
	} else {
//line /snap/go/10455/src/strings/strings.go:772
		_go_fuzz_dep_.CoverTab[525372]++
//line /snap/go/10455/src/strings/strings.go:772
		_go_fuzz_dep_.CoverTab[1551]++
//line /snap/go/10455/src/strings/strings.go:772
		// _ = "end of CoverTab[1551]"
//line /snap/go/10455/src/strings/strings.go:772
	}
//line /snap/go/10455/src/strings/strings.go:772
	// _ = "end of CoverTab[1536]"
//line /snap/go/10455/src/strings/strings.go:772
	_go_fuzz_dep_.CoverTab[1537]++

							return unicode.IsSpace(r)
//line /snap/go/10455/src/strings/strings.go:774
	// _ = "end of CoverTab[1537]"
}

// Title returns a copy of the string s with all Unicode letters that begin words
//line /snap/go/10455/src/strings/strings.go:777
// mapped to their Unicode title case.
//line /snap/go/10455/src/strings/strings.go:777
//
//line /snap/go/10455/src/strings/strings.go:777
// Deprecated: The rule Title uses for word boundaries does not handle Unicode
//line /snap/go/10455/src/strings/strings.go:777
// punctuation properly. Use golang.org/x/text/cases instead.
//line /snap/go/10455/src/strings/strings.go:782
func Title(s string) string {
//line /snap/go/10455/src/strings/strings.go:782
	_go_fuzz_dep_.CoverTab[1552]++

//line /snap/go/10455/src/strings/strings.go:786
	prev := ' '
	return Map(
		func(r rune) rune {
//line /snap/go/10455/src/strings/strings.go:788
			_go_fuzz_dep_.CoverTab[1553]++
									if isSeparator(prev) {
//line /snap/go/10455/src/strings/strings.go:789
				_go_fuzz_dep_.CoverTab[525373]++
//line /snap/go/10455/src/strings/strings.go:789
				_go_fuzz_dep_.CoverTab[1555]++
										prev = r
										return unicode.ToTitle(r)
//line /snap/go/10455/src/strings/strings.go:791
				// _ = "end of CoverTab[1555]"
			} else {
//line /snap/go/10455/src/strings/strings.go:792
				_go_fuzz_dep_.CoverTab[525374]++
//line /snap/go/10455/src/strings/strings.go:792
				_go_fuzz_dep_.CoverTab[1556]++
//line /snap/go/10455/src/strings/strings.go:792
				// _ = "end of CoverTab[1556]"
//line /snap/go/10455/src/strings/strings.go:792
			}
//line /snap/go/10455/src/strings/strings.go:792
			// _ = "end of CoverTab[1553]"
//line /snap/go/10455/src/strings/strings.go:792
			_go_fuzz_dep_.CoverTab[1554]++
									prev = r
									return r
//line /snap/go/10455/src/strings/strings.go:794
			// _ = "end of CoverTab[1554]"
		},
		s)
//line /snap/go/10455/src/strings/strings.go:796
	// _ = "end of CoverTab[1552]"
}

// TrimLeftFunc returns a slice of the string s with all leading
//line /snap/go/10455/src/strings/strings.go:799
// Unicode code points c satisfying f(c) removed.
//line /snap/go/10455/src/strings/strings.go:801
func TrimLeftFunc(s string, f func(rune) bool) string {
//line /snap/go/10455/src/strings/strings.go:801
	_go_fuzz_dep_.CoverTab[1557]++
							i := indexFunc(s, f, false)
							if i == -1 {
//line /snap/go/10455/src/strings/strings.go:803
		_go_fuzz_dep_.CoverTab[525375]++
//line /snap/go/10455/src/strings/strings.go:803
		_go_fuzz_dep_.CoverTab[1559]++
								return ""
//line /snap/go/10455/src/strings/strings.go:804
		// _ = "end of CoverTab[1559]"
	} else {
//line /snap/go/10455/src/strings/strings.go:805
		_go_fuzz_dep_.CoverTab[525376]++
//line /snap/go/10455/src/strings/strings.go:805
		_go_fuzz_dep_.CoverTab[1560]++
//line /snap/go/10455/src/strings/strings.go:805
		// _ = "end of CoverTab[1560]"
//line /snap/go/10455/src/strings/strings.go:805
	}
//line /snap/go/10455/src/strings/strings.go:805
	// _ = "end of CoverTab[1557]"
//line /snap/go/10455/src/strings/strings.go:805
	_go_fuzz_dep_.CoverTab[1558]++
							return s[i:]
//line /snap/go/10455/src/strings/strings.go:806
	// _ = "end of CoverTab[1558]"
}

// TrimRightFunc returns a slice of the string s with all trailing
//line /snap/go/10455/src/strings/strings.go:809
// Unicode code points c satisfying f(c) removed.
//line /snap/go/10455/src/strings/strings.go:811
func TrimRightFunc(s string, f func(rune) bool) string {
//line /snap/go/10455/src/strings/strings.go:811
	_go_fuzz_dep_.CoverTab[1561]++
							i := lastIndexFunc(s, f, false)
							if i >= 0 && func() bool {
//line /snap/go/10455/src/strings/strings.go:813
		_go_fuzz_dep_.CoverTab[1563]++
//line /snap/go/10455/src/strings/strings.go:813
		return s[i] >= utf8.RuneSelf
//line /snap/go/10455/src/strings/strings.go:813
		// _ = "end of CoverTab[1563]"
//line /snap/go/10455/src/strings/strings.go:813
	}() {
//line /snap/go/10455/src/strings/strings.go:813
		_go_fuzz_dep_.CoverTab[525377]++
//line /snap/go/10455/src/strings/strings.go:813
		_go_fuzz_dep_.CoverTab[1564]++
								_, wid := utf8.DecodeRuneInString(s[i:])
								i += wid
//line /snap/go/10455/src/strings/strings.go:815
		// _ = "end of CoverTab[1564]"
	} else {
//line /snap/go/10455/src/strings/strings.go:816
		_go_fuzz_dep_.CoverTab[525378]++
//line /snap/go/10455/src/strings/strings.go:816
		_go_fuzz_dep_.CoverTab[1565]++
								i++
//line /snap/go/10455/src/strings/strings.go:817
		// _ = "end of CoverTab[1565]"
	}
//line /snap/go/10455/src/strings/strings.go:818
	// _ = "end of CoverTab[1561]"
//line /snap/go/10455/src/strings/strings.go:818
	_go_fuzz_dep_.CoverTab[1562]++
							return s[0:i]
//line /snap/go/10455/src/strings/strings.go:819
	// _ = "end of CoverTab[1562]"
}

// TrimFunc returns a slice of the string s with all leading
//line /snap/go/10455/src/strings/strings.go:822
// and trailing Unicode code points c satisfying f(c) removed.
//line /snap/go/10455/src/strings/strings.go:824
func TrimFunc(s string, f func(rune) bool) string {
//line /snap/go/10455/src/strings/strings.go:824
	_go_fuzz_dep_.CoverTab[1566]++
							return TrimRightFunc(TrimLeftFunc(s, f), f)
//line /snap/go/10455/src/strings/strings.go:825
	// _ = "end of CoverTab[1566]"
}

// IndexFunc returns the index into s of the first Unicode
//line /snap/go/10455/src/strings/strings.go:828
// code point satisfying f(c), or -1 if none do.
//line /snap/go/10455/src/strings/strings.go:830
func IndexFunc(s string, f func(rune) bool) int {
//line /snap/go/10455/src/strings/strings.go:830
	_go_fuzz_dep_.CoverTab[1567]++
							return indexFunc(s, f, true)
//line /snap/go/10455/src/strings/strings.go:831
	// _ = "end of CoverTab[1567]"
}

// LastIndexFunc returns the index into s of the last
//line /snap/go/10455/src/strings/strings.go:834
// Unicode code point satisfying f(c), or -1 if none do.
//line /snap/go/10455/src/strings/strings.go:836
func LastIndexFunc(s string, f func(rune) bool) int {
//line /snap/go/10455/src/strings/strings.go:836
	_go_fuzz_dep_.CoverTab[1568]++
							return lastIndexFunc(s, f, true)
//line /snap/go/10455/src/strings/strings.go:837
	// _ = "end of CoverTab[1568]"
}

// indexFunc is the same as IndexFunc except that if
//line /snap/go/10455/src/strings/strings.go:840
// truth==false, the sense of the predicate function is
//line /snap/go/10455/src/strings/strings.go:840
// inverted.
//line /snap/go/10455/src/strings/strings.go:843
func indexFunc(s string, f func(rune) bool, truth bool) int {
//line /snap/go/10455/src/strings/strings.go:843
	_go_fuzz_dep_.CoverTab[1569]++
//line /snap/go/10455/src/strings/strings.go:843
	_go_fuzz_dep_.CoverTab[786538] = 0
							for i, r := range s {
//line /snap/go/10455/src/strings/strings.go:844
		if _go_fuzz_dep_.CoverTab[786538] == 0 {
//line /snap/go/10455/src/strings/strings.go:844
			_go_fuzz_dep_.CoverTab[525607]++
//line /snap/go/10455/src/strings/strings.go:844
		} else {
//line /snap/go/10455/src/strings/strings.go:844
			_go_fuzz_dep_.CoverTab[525608]++
//line /snap/go/10455/src/strings/strings.go:844
		}
//line /snap/go/10455/src/strings/strings.go:844
		_go_fuzz_dep_.CoverTab[786538] = 1
//line /snap/go/10455/src/strings/strings.go:844
		_go_fuzz_dep_.CoverTab[1571]++
								if f(r) == truth {
//line /snap/go/10455/src/strings/strings.go:845
			_go_fuzz_dep_.CoverTab[525379]++
//line /snap/go/10455/src/strings/strings.go:845
			_go_fuzz_dep_.CoverTab[1572]++
									return i
//line /snap/go/10455/src/strings/strings.go:846
			// _ = "end of CoverTab[1572]"
		} else {
//line /snap/go/10455/src/strings/strings.go:847
			_go_fuzz_dep_.CoverTab[525380]++
//line /snap/go/10455/src/strings/strings.go:847
			_go_fuzz_dep_.CoverTab[1573]++
//line /snap/go/10455/src/strings/strings.go:847
			// _ = "end of CoverTab[1573]"
//line /snap/go/10455/src/strings/strings.go:847
		}
//line /snap/go/10455/src/strings/strings.go:847
		// _ = "end of CoverTab[1571]"
	}
//line /snap/go/10455/src/strings/strings.go:848
	if _go_fuzz_dep_.CoverTab[786538] == 0 {
//line /snap/go/10455/src/strings/strings.go:848
		_go_fuzz_dep_.CoverTab[525609]++
//line /snap/go/10455/src/strings/strings.go:848
	} else {
//line /snap/go/10455/src/strings/strings.go:848
		_go_fuzz_dep_.CoverTab[525610]++
//line /snap/go/10455/src/strings/strings.go:848
	}
//line /snap/go/10455/src/strings/strings.go:848
	// _ = "end of CoverTab[1569]"
//line /snap/go/10455/src/strings/strings.go:848
	_go_fuzz_dep_.CoverTab[1570]++
							return -1
//line /snap/go/10455/src/strings/strings.go:849
	// _ = "end of CoverTab[1570]"
}

// lastIndexFunc is the same as LastIndexFunc except that if
//line /snap/go/10455/src/strings/strings.go:852
// truth==false, the sense of the predicate function is
//line /snap/go/10455/src/strings/strings.go:852
// inverted.
//line /snap/go/10455/src/strings/strings.go:855
func lastIndexFunc(s string, f func(rune) bool, truth bool) int {
//line /snap/go/10455/src/strings/strings.go:855
	_go_fuzz_dep_.CoverTab[1574]++
//line /snap/go/10455/src/strings/strings.go:855
	_go_fuzz_dep_.CoverTab[786539] = 0
							for i := len(s); i > 0; {
//line /snap/go/10455/src/strings/strings.go:856
		if _go_fuzz_dep_.CoverTab[786539] == 0 {
//line /snap/go/10455/src/strings/strings.go:856
			_go_fuzz_dep_.CoverTab[525611]++
//line /snap/go/10455/src/strings/strings.go:856
		} else {
//line /snap/go/10455/src/strings/strings.go:856
			_go_fuzz_dep_.CoverTab[525612]++
//line /snap/go/10455/src/strings/strings.go:856
		}
//line /snap/go/10455/src/strings/strings.go:856
		_go_fuzz_dep_.CoverTab[786539] = 1
//line /snap/go/10455/src/strings/strings.go:856
		_go_fuzz_dep_.CoverTab[1576]++
								r, size := utf8.DecodeLastRuneInString(s[0:i])
								i -= size
								if f(r) == truth {
//line /snap/go/10455/src/strings/strings.go:859
			_go_fuzz_dep_.CoverTab[525381]++
//line /snap/go/10455/src/strings/strings.go:859
			_go_fuzz_dep_.CoverTab[1577]++
									return i
//line /snap/go/10455/src/strings/strings.go:860
			// _ = "end of CoverTab[1577]"
		} else {
//line /snap/go/10455/src/strings/strings.go:861
			_go_fuzz_dep_.CoverTab[525382]++
//line /snap/go/10455/src/strings/strings.go:861
			_go_fuzz_dep_.CoverTab[1578]++
//line /snap/go/10455/src/strings/strings.go:861
			// _ = "end of CoverTab[1578]"
//line /snap/go/10455/src/strings/strings.go:861
		}
//line /snap/go/10455/src/strings/strings.go:861
		// _ = "end of CoverTab[1576]"
	}
//line /snap/go/10455/src/strings/strings.go:862
	if _go_fuzz_dep_.CoverTab[786539] == 0 {
//line /snap/go/10455/src/strings/strings.go:862
		_go_fuzz_dep_.CoverTab[525613]++
//line /snap/go/10455/src/strings/strings.go:862
	} else {
//line /snap/go/10455/src/strings/strings.go:862
		_go_fuzz_dep_.CoverTab[525614]++
//line /snap/go/10455/src/strings/strings.go:862
	}
//line /snap/go/10455/src/strings/strings.go:862
	// _ = "end of CoverTab[1574]"
//line /snap/go/10455/src/strings/strings.go:862
	_go_fuzz_dep_.CoverTab[1575]++
							return -1
//line /snap/go/10455/src/strings/strings.go:863
	// _ = "end of CoverTab[1575]"
}

// asciiSet is a 32-byte value, where each bit represents the presence of a
//line /snap/go/10455/src/strings/strings.go:866
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
//line /snap/go/10455/src/strings/strings.go:866
// starting with the least-significant bit of the lowest word to the
//line /snap/go/10455/src/strings/strings.go:866
// most-significant bit of the highest word, map to the full range of all
//line /snap/go/10455/src/strings/strings.go:866
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
//line /snap/go/10455/src/strings/strings.go:866
// ensuring that any non-ASCII character will be reported as not in the set.
//line /snap/go/10455/src/strings/strings.go:866
// This allocates a total of 32 bytes even though the upper half
//line /snap/go/10455/src/strings/strings.go:866
// is unused to avoid bounds checks in asciiSet.contains.
//line /snap/go/10455/src/strings/strings.go:874
type asciiSet [8]uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
//line /snap/go/10455/src/strings/strings.go:876
// characters in chars are ASCII.
//line /snap/go/10455/src/strings/strings.go:878
func makeASCIISet(chars string) (as asciiSet, ok bool) {
//line /snap/go/10455/src/strings/strings.go:878
	_go_fuzz_dep_.CoverTab[1579]++
//line /snap/go/10455/src/strings/strings.go:878
	_go_fuzz_dep_.CoverTab[786540] = 0
							for i := 0; i < len(chars); i++ {
//line /snap/go/10455/src/strings/strings.go:879
		if _go_fuzz_dep_.CoverTab[786540] == 0 {
//line /snap/go/10455/src/strings/strings.go:879
			_go_fuzz_dep_.CoverTab[525615]++
//line /snap/go/10455/src/strings/strings.go:879
		} else {
//line /snap/go/10455/src/strings/strings.go:879
			_go_fuzz_dep_.CoverTab[525616]++
//line /snap/go/10455/src/strings/strings.go:879
		}
//line /snap/go/10455/src/strings/strings.go:879
		_go_fuzz_dep_.CoverTab[786540] = 1
//line /snap/go/10455/src/strings/strings.go:879
		_go_fuzz_dep_.CoverTab[1581]++
								c := chars[i]
								if c >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:881
			_go_fuzz_dep_.CoverTab[525383]++
//line /snap/go/10455/src/strings/strings.go:881
			_go_fuzz_dep_.CoverTab[1583]++
									return as, false
//line /snap/go/10455/src/strings/strings.go:882
			// _ = "end of CoverTab[1583]"
		} else {
//line /snap/go/10455/src/strings/strings.go:883
			_go_fuzz_dep_.CoverTab[525384]++
//line /snap/go/10455/src/strings/strings.go:883
			_go_fuzz_dep_.CoverTab[1584]++
//line /snap/go/10455/src/strings/strings.go:883
			// _ = "end of CoverTab[1584]"
//line /snap/go/10455/src/strings/strings.go:883
		}
//line /snap/go/10455/src/strings/strings.go:883
		// _ = "end of CoverTab[1581]"
//line /snap/go/10455/src/strings/strings.go:883
		_go_fuzz_dep_.CoverTab[1582]++
								as[c/32] |= 1 << (c % 32)
//line /snap/go/10455/src/strings/strings.go:884
		// _ = "end of CoverTab[1582]"
	}
//line /snap/go/10455/src/strings/strings.go:885
	if _go_fuzz_dep_.CoverTab[786540] == 0 {
//line /snap/go/10455/src/strings/strings.go:885
		_go_fuzz_dep_.CoverTab[525617]++
//line /snap/go/10455/src/strings/strings.go:885
	} else {
//line /snap/go/10455/src/strings/strings.go:885
		_go_fuzz_dep_.CoverTab[525618]++
//line /snap/go/10455/src/strings/strings.go:885
	}
//line /snap/go/10455/src/strings/strings.go:885
	// _ = "end of CoverTab[1579]"
//line /snap/go/10455/src/strings/strings.go:885
	_go_fuzz_dep_.CoverTab[1580]++
							return as, true
//line /snap/go/10455/src/strings/strings.go:886
	// _ = "end of CoverTab[1580]"
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c byte) bool {
//line /snap/go/10455/src/strings/strings.go:890
	_go_fuzz_dep_.CoverTab[1585]++
							return (as[c/32] & (1 << (c % 32))) != 0
//line /snap/go/10455/src/strings/strings.go:891
	// _ = "end of CoverTab[1585]"
}

// Trim returns a slice of the string s with all leading and
//line /snap/go/10455/src/strings/strings.go:894
// trailing Unicode code points contained in cutset removed.
//line /snap/go/10455/src/strings/strings.go:896
func Trim(s, cutset string) string {
//line /snap/go/10455/src/strings/strings.go:896
	_go_fuzz_dep_.CoverTab[1586]++
							if s == "" || func() bool {
//line /snap/go/10455/src/strings/strings.go:897
		_go_fuzz_dep_.CoverTab[1590]++
//line /snap/go/10455/src/strings/strings.go:897
		return cutset == ""
//line /snap/go/10455/src/strings/strings.go:897
		// _ = "end of CoverTab[1590]"
//line /snap/go/10455/src/strings/strings.go:897
	}() {
//line /snap/go/10455/src/strings/strings.go:897
		_go_fuzz_dep_.CoverTab[525385]++
//line /snap/go/10455/src/strings/strings.go:897
		_go_fuzz_dep_.CoverTab[1591]++
								return s
//line /snap/go/10455/src/strings/strings.go:898
		// _ = "end of CoverTab[1591]"
	} else {
//line /snap/go/10455/src/strings/strings.go:899
		_go_fuzz_dep_.CoverTab[525386]++
//line /snap/go/10455/src/strings/strings.go:899
		_go_fuzz_dep_.CoverTab[1592]++
//line /snap/go/10455/src/strings/strings.go:899
		// _ = "end of CoverTab[1592]"
//line /snap/go/10455/src/strings/strings.go:899
	}
//line /snap/go/10455/src/strings/strings.go:899
	// _ = "end of CoverTab[1586]"
//line /snap/go/10455/src/strings/strings.go:899
	_go_fuzz_dep_.CoverTab[1587]++
							if len(cutset) == 1 && func() bool {
//line /snap/go/10455/src/strings/strings.go:900
		_go_fuzz_dep_.CoverTab[1593]++
//line /snap/go/10455/src/strings/strings.go:900
		return cutset[0] < utf8.RuneSelf
//line /snap/go/10455/src/strings/strings.go:900
		// _ = "end of CoverTab[1593]"
//line /snap/go/10455/src/strings/strings.go:900
	}() {
//line /snap/go/10455/src/strings/strings.go:900
		_go_fuzz_dep_.CoverTab[525387]++
//line /snap/go/10455/src/strings/strings.go:900
		_go_fuzz_dep_.CoverTab[1594]++
								return trimLeftByte(trimRightByte(s, cutset[0]), cutset[0])
//line /snap/go/10455/src/strings/strings.go:901
		// _ = "end of CoverTab[1594]"
	} else {
//line /snap/go/10455/src/strings/strings.go:902
		_go_fuzz_dep_.CoverTab[525388]++
//line /snap/go/10455/src/strings/strings.go:902
		_go_fuzz_dep_.CoverTab[1595]++
//line /snap/go/10455/src/strings/strings.go:902
		// _ = "end of CoverTab[1595]"
//line /snap/go/10455/src/strings/strings.go:902
	}
//line /snap/go/10455/src/strings/strings.go:902
	// _ = "end of CoverTab[1587]"
//line /snap/go/10455/src/strings/strings.go:902
	_go_fuzz_dep_.CoverTab[1588]++
							if as, ok := makeASCIISet(cutset); ok {
//line /snap/go/10455/src/strings/strings.go:903
		_go_fuzz_dep_.CoverTab[525389]++
//line /snap/go/10455/src/strings/strings.go:903
		_go_fuzz_dep_.CoverTab[1596]++
								return trimLeftASCII(trimRightASCII(s, &as), &as)
//line /snap/go/10455/src/strings/strings.go:904
		// _ = "end of CoverTab[1596]"
	} else {
//line /snap/go/10455/src/strings/strings.go:905
		_go_fuzz_dep_.CoverTab[525390]++
//line /snap/go/10455/src/strings/strings.go:905
		_go_fuzz_dep_.CoverTab[1597]++
//line /snap/go/10455/src/strings/strings.go:905
		// _ = "end of CoverTab[1597]"
//line /snap/go/10455/src/strings/strings.go:905
	}
//line /snap/go/10455/src/strings/strings.go:905
	// _ = "end of CoverTab[1588]"
//line /snap/go/10455/src/strings/strings.go:905
	_go_fuzz_dep_.CoverTab[1589]++
							return trimLeftUnicode(trimRightUnicode(s, cutset), cutset)
//line /snap/go/10455/src/strings/strings.go:906
	// _ = "end of CoverTab[1589]"
}

// TrimLeft returns a slice of the string s with all leading
//line /snap/go/10455/src/strings/strings.go:909
// Unicode code points contained in cutset removed.
//line /snap/go/10455/src/strings/strings.go:909
//
//line /snap/go/10455/src/strings/strings.go:909
// To remove a prefix, use TrimPrefix instead.
//line /snap/go/10455/src/strings/strings.go:913
func TrimLeft(s, cutset string) string {
//line /snap/go/10455/src/strings/strings.go:913
	_go_fuzz_dep_.CoverTab[1598]++
							if s == "" || func() bool {
//line /snap/go/10455/src/strings/strings.go:914
		_go_fuzz_dep_.CoverTab[1602]++
//line /snap/go/10455/src/strings/strings.go:914
		return cutset == ""
//line /snap/go/10455/src/strings/strings.go:914
		// _ = "end of CoverTab[1602]"
//line /snap/go/10455/src/strings/strings.go:914
	}() {
//line /snap/go/10455/src/strings/strings.go:914
		_go_fuzz_dep_.CoverTab[525391]++
//line /snap/go/10455/src/strings/strings.go:914
		_go_fuzz_dep_.CoverTab[1603]++
								return s
//line /snap/go/10455/src/strings/strings.go:915
		// _ = "end of CoverTab[1603]"
	} else {
//line /snap/go/10455/src/strings/strings.go:916
		_go_fuzz_dep_.CoverTab[525392]++
//line /snap/go/10455/src/strings/strings.go:916
		_go_fuzz_dep_.CoverTab[1604]++
//line /snap/go/10455/src/strings/strings.go:916
		// _ = "end of CoverTab[1604]"
//line /snap/go/10455/src/strings/strings.go:916
	}
//line /snap/go/10455/src/strings/strings.go:916
	// _ = "end of CoverTab[1598]"
//line /snap/go/10455/src/strings/strings.go:916
	_go_fuzz_dep_.CoverTab[1599]++
							if len(cutset) == 1 && func() bool {
//line /snap/go/10455/src/strings/strings.go:917
		_go_fuzz_dep_.CoverTab[1605]++
//line /snap/go/10455/src/strings/strings.go:917
		return cutset[0] < utf8.RuneSelf
//line /snap/go/10455/src/strings/strings.go:917
		// _ = "end of CoverTab[1605]"
//line /snap/go/10455/src/strings/strings.go:917
	}() {
//line /snap/go/10455/src/strings/strings.go:917
		_go_fuzz_dep_.CoverTab[525393]++
//line /snap/go/10455/src/strings/strings.go:917
		_go_fuzz_dep_.CoverTab[1606]++
								return trimLeftByte(s, cutset[0])
//line /snap/go/10455/src/strings/strings.go:918
		// _ = "end of CoverTab[1606]"
	} else {
//line /snap/go/10455/src/strings/strings.go:919
		_go_fuzz_dep_.CoverTab[525394]++
//line /snap/go/10455/src/strings/strings.go:919
		_go_fuzz_dep_.CoverTab[1607]++
//line /snap/go/10455/src/strings/strings.go:919
		// _ = "end of CoverTab[1607]"
//line /snap/go/10455/src/strings/strings.go:919
	}
//line /snap/go/10455/src/strings/strings.go:919
	// _ = "end of CoverTab[1599]"
//line /snap/go/10455/src/strings/strings.go:919
	_go_fuzz_dep_.CoverTab[1600]++
							if as, ok := makeASCIISet(cutset); ok {
//line /snap/go/10455/src/strings/strings.go:920
		_go_fuzz_dep_.CoverTab[525395]++
//line /snap/go/10455/src/strings/strings.go:920
		_go_fuzz_dep_.CoverTab[1608]++
								return trimLeftASCII(s, &as)
//line /snap/go/10455/src/strings/strings.go:921
		// _ = "end of CoverTab[1608]"
	} else {
//line /snap/go/10455/src/strings/strings.go:922
		_go_fuzz_dep_.CoverTab[525396]++
//line /snap/go/10455/src/strings/strings.go:922
		_go_fuzz_dep_.CoverTab[1609]++
//line /snap/go/10455/src/strings/strings.go:922
		// _ = "end of CoverTab[1609]"
//line /snap/go/10455/src/strings/strings.go:922
	}
//line /snap/go/10455/src/strings/strings.go:922
	// _ = "end of CoverTab[1600]"
//line /snap/go/10455/src/strings/strings.go:922
	_go_fuzz_dep_.CoverTab[1601]++
							return trimLeftUnicode(s, cutset)
//line /snap/go/10455/src/strings/strings.go:923
	// _ = "end of CoverTab[1601]"
}

func trimLeftByte(s string, c byte) string {
//line /snap/go/10455/src/strings/strings.go:926
	_go_fuzz_dep_.CoverTab[1610]++
//line /snap/go/10455/src/strings/strings.go:926
	_go_fuzz_dep_.CoverTab[786541] = 0
							for len(s) > 0 && func() bool {
//line /snap/go/10455/src/strings/strings.go:927
		_go_fuzz_dep_.CoverTab[1612]++
//line /snap/go/10455/src/strings/strings.go:927
		return s[0] == c
//line /snap/go/10455/src/strings/strings.go:927
		// _ = "end of CoverTab[1612]"
//line /snap/go/10455/src/strings/strings.go:927
	}() {
//line /snap/go/10455/src/strings/strings.go:927
		if _go_fuzz_dep_.CoverTab[786541] == 0 {
//line /snap/go/10455/src/strings/strings.go:927
			_go_fuzz_dep_.CoverTab[525619]++
//line /snap/go/10455/src/strings/strings.go:927
		} else {
//line /snap/go/10455/src/strings/strings.go:927
			_go_fuzz_dep_.CoverTab[525620]++
//line /snap/go/10455/src/strings/strings.go:927
		}
//line /snap/go/10455/src/strings/strings.go:927
		_go_fuzz_dep_.CoverTab[786541] = 1
//line /snap/go/10455/src/strings/strings.go:927
		_go_fuzz_dep_.CoverTab[1613]++
								s = s[1:]
//line /snap/go/10455/src/strings/strings.go:928
		// _ = "end of CoverTab[1613]"
	}
//line /snap/go/10455/src/strings/strings.go:929
	if _go_fuzz_dep_.CoverTab[786541] == 0 {
//line /snap/go/10455/src/strings/strings.go:929
		_go_fuzz_dep_.CoverTab[525621]++
//line /snap/go/10455/src/strings/strings.go:929
	} else {
//line /snap/go/10455/src/strings/strings.go:929
		_go_fuzz_dep_.CoverTab[525622]++
//line /snap/go/10455/src/strings/strings.go:929
	}
//line /snap/go/10455/src/strings/strings.go:929
	// _ = "end of CoverTab[1610]"
//line /snap/go/10455/src/strings/strings.go:929
	_go_fuzz_dep_.CoverTab[1611]++
							return s
//line /snap/go/10455/src/strings/strings.go:930
	// _ = "end of CoverTab[1611]"
}

func trimLeftASCII(s string, as *asciiSet) string {
//line /snap/go/10455/src/strings/strings.go:933
	_go_fuzz_dep_.CoverTab[1614]++
//line /snap/go/10455/src/strings/strings.go:933
	_go_fuzz_dep_.CoverTab[786542] = 0
							for len(s) > 0 {
//line /snap/go/10455/src/strings/strings.go:934
		if _go_fuzz_dep_.CoverTab[786542] == 0 {
//line /snap/go/10455/src/strings/strings.go:934
			_go_fuzz_dep_.CoverTab[525623]++
//line /snap/go/10455/src/strings/strings.go:934
		} else {
//line /snap/go/10455/src/strings/strings.go:934
			_go_fuzz_dep_.CoverTab[525624]++
//line /snap/go/10455/src/strings/strings.go:934
		}
//line /snap/go/10455/src/strings/strings.go:934
		_go_fuzz_dep_.CoverTab[786542] = 1
//line /snap/go/10455/src/strings/strings.go:934
		_go_fuzz_dep_.CoverTab[1616]++
								if !as.contains(s[0]) {
//line /snap/go/10455/src/strings/strings.go:935
			_go_fuzz_dep_.CoverTab[525397]++
//line /snap/go/10455/src/strings/strings.go:935
			_go_fuzz_dep_.CoverTab[1618]++
									break
//line /snap/go/10455/src/strings/strings.go:936
			// _ = "end of CoverTab[1618]"
		} else {
//line /snap/go/10455/src/strings/strings.go:937
			_go_fuzz_dep_.CoverTab[525398]++
//line /snap/go/10455/src/strings/strings.go:937
			_go_fuzz_dep_.CoverTab[1619]++
//line /snap/go/10455/src/strings/strings.go:937
			// _ = "end of CoverTab[1619]"
//line /snap/go/10455/src/strings/strings.go:937
		}
//line /snap/go/10455/src/strings/strings.go:937
		// _ = "end of CoverTab[1616]"
//line /snap/go/10455/src/strings/strings.go:937
		_go_fuzz_dep_.CoverTab[1617]++
								s = s[1:]
//line /snap/go/10455/src/strings/strings.go:938
		// _ = "end of CoverTab[1617]"
	}
//line /snap/go/10455/src/strings/strings.go:939
	if _go_fuzz_dep_.CoverTab[786542] == 0 {
//line /snap/go/10455/src/strings/strings.go:939
		_go_fuzz_dep_.CoverTab[525625]++
//line /snap/go/10455/src/strings/strings.go:939
	} else {
//line /snap/go/10455/src/strings/strings.go:939
		_go_fuzz_dep_.CoverTab[525626]++
//line /snap/go/10455/src/strings/strings.go:939
	}
//line /snap/go/10455/src/strings/strings.go:939
	// _ = "end of CoverTab[1614]"
//line /snap/go/10455/src/strings/strings.go:939
	_go_fuzz_dep_.CoverTab[1615]++
							return s
//line /snap/go/10455/src/strings/strings.go:940
	// _ = "end of CoverTab[1615]"
}

func trimLeftUnicode(s, cutset string) string {
//line /snap/go/10455/src/strings/strings.go:943
	_go_fuzz_dep_.CoverTab[1620]++
//line /snap/go/10455/src/strings/strings.go:943
	_go_fuzz_dep_.CoverTab[786543] = 0
							for len(s) > 0 {
//line /snap/go/10455/src/strings/strings.go:944
		if _go_fuzz_dep_.CoverTab[786543] == 0 {
//line /snap/go/10455/src/strings/strings.go:944
			_go_fuzz_dep_.CoverTab[525627]++
//line /snap/go/10455/src/strings/strings.go:944
		} else {
//line /snap/go/10455/src/strings/strings.go:944
			_go_fuzz_dep_.CoverTab[525628]++
//line /snap/go/10455/src/strings/strings.go:944
		}
//line /snap/go/10455/src/strings/strings.go:944
		_go_fuzz_dep_.CoverTab[786543] = 1
//line /snap/go/10455/src/strings/strings.go:944
		_go_fuzz_dep_.CoverTab[1622]++
								r, n := rune(s[0]), 1
								if r >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:946
			_go_fuzz_dep_.CoverTab[525399]++
//line /snap/go/10455/src/strings/strings.go:946
			_go_fuzz_dep_.CoverTab[1625]++
									r, n = utf8.DecodeRuneInString(s)
//line /snap/go/10455/src/strings/strings.go:947
			// _ = "end of CoverTab[1625]"
		} else {
//line /snap/go/10455/src/strings/strings.go:948
			_go_fuzz_dep_.CoverTab[525400]++
//line /snap/go/10455/src/strings/strings.go:948
			_go_fuzz_dep_.CoverTab[1626]++
//line /snap/go/10455/src/strings/strings.go:948
			// _ = "end of CoverTab[1626]"
//line /snap/go/10455/src/strings/strings.go:948
		}
//line /snap/go/10455/src/strings/strings.go:948
		// _ = "end of CoverTab[1622]"
//line /snap/go/10455/src/strings/strings.go:948
		_go_fuzz_dep_.CoverTab[1623]++
								if !ContainsRune(cutset, r) {
//line /snap/go/10455/src/strings/strings.go:949
			_go_fuzz_dep_.CoverTab[525401]++
//line /snap/go/10455/src/strings/strings.go:949
			_go_fuzz_dep_.CoverTab[1627]++
									break
//line /snap/go/10455/src/strings/strings.go:950
			// _ = "end of CoverTab[1627]"
		} else {
//line /snap/go/10455/src/strings/strings.go:951
			_go_fuzz_dep_.CoverTab[525402]++
//line /snap/go/10455/src/strings/strings.go:951
			_go_fuzz_dep_.CoverTab[1628]++
//line /snap/go/10455/src/strings/strings.go:951
			// _ = "end of CoverTab[1628]"
//line /snap/go/10455/src/strings/strings.go:951
		}
//line /snap/go/10455/src/strings/strings.go:951
		// _ = "end of CoverTab[1623]"
//line /snap/go/10455/src/strings/strings.go:951
		_go_fuzz_dep_.CoverTab[1624]++
								s = s[n:]
//line /snap/go/10455/src/strings/strings.go:952
		// _ = "end of CoverTab[1624]"
	}
//line /snap/go/10455/src/strings/strings.go:953
	if _go_fuzz_dep_.CoverTab[786543] == 0 {
//line /snap/go/10455/src/strings/strings.go:953
		_go_fuzz_dep_.CoverTab[525629]++
//line /snap/go/10455/src/strings/strings.go:953
	} else {
//line /snap/go/10455/src/strings/strings.go:953
		_go_fuzz_dep_.CoverTab[525630]++
//line /snap/go/10455/src/strings/strings.go:953
	}
//line /snap/go/10455/src/strings/strings.go:953
	// _ = "end of CoverTab[1620]"
//line /snap/go/10455/src/strings/strings.go:953
	_go_fuzz_dep_.CoverTab[1621]++
							return s
//line /snap/go/10455/src/strings/strings.go:954
	// _ = "end of CoverTab[1621]"
}

// TrimRight returns a slice of the string s, with all trailing
//line /snap/go/10455/src/strings/strings.go:957
// Unicode code points contained in cutset removed.
//line /snap/go/10455/src/strings/strings.go:957
//
//line /snap/go/10455/src/strings/strings.go:957
// To remove a suffix, use TrimSuffix instead.
//line /snap/go/10455/src/strings/strings.go:961
func TrimRight(s, cutset string) string {
//line /snap/go/10455/src/strings/strings.go:961
	_go_fuzz_dep_.CoverTab[1629]++
							if s == "" || func() bool {
//line /snap/go/10455/src/strings/strings.go:962
		_go_fuzz_dep_.CoverTab[1633]++
//line /snap/go/10455/src/strings/strings.go:962
		return cutset == ""
//line /snap/go/10455/src/strings/strings.go:962
		// _ = "end of CoverTab[1633]"
//line /snap/go/10455/src/strings/strings.go:962
	}() {
//line /snap/go/10455/src/strings/strings.go:962
		_go_fuzz_dep_.CoverTab[525403]++
//line /snap/go/10455/src/strings/strings.go:962
		_go_fuzz_dep_.CoverTab[1634]++
								return s
//line /snap/go/10455/src/strings/strings.go:963
		// _ = "end of CoverTab[1634]"
	} else {
//line /snap/go/10455/src/strings/strings.go:964
		_go_fuzz_dep_.CoverTab[525404]++
//line /snap/go/10455/src/strings/strings.go:964
		_go_fuzz_dep_.CoverTab[1635]++
//line /snap/go/10455/src/strings/strings.go:964
		// _ = "end of CoverTab[1635]"
//line /snap/go/10455/src/strings/strings.go:964
	}
//line /snap/go/10455/src/strings/strings.go:964
	// _ = "end of CoverTab[1629]"
//line /snap/go/10455/src/strings/strings.go:964
	_go_fuzz_dep_.CoverTab[1630]++
							if len(cutset) == 1 && func() bool {
//line /snap/go/10455/src/strings/strings.go:965
		_go_fuzz_dep_.CoverTab[1636]++
//line /snap/go/10455/src/strings/strings.go:965
		return cutset[0] < utf8.RuneSelf
//line /snap/go/10455/src/strings/strings.go:965
		// _ = "end of CoverTab[1636]"
//line /snap/go/10455/src/strings/strings.go:965
	}() {
//line /snap/go/10455/src/strings/strings.go:965
		_go_fuzz_dep_.CoverTab[525405]++
//line /snap/go/10455/src/strings/strings.go:965
		_go_fuzz_dep_.CoverTab[1637]++
								return trimRightByte(s, cutset[0])
//line /snap/go/10455/src/strings/strings.go:966
		// _ = "end of CoverTab[1637]"
	} else {
//line /snap/go/10455/src/strings/strings.go:967
		_go_fuzz_dep_.CoverTab[525406]++
//line /snap/go/10455/src/strings/strings.go:967
		_go_fuzz_dep_.CoverTab[1638]++
//line /snap/go/10455/src/strings/strings.go:967
		// _ = "end of CoverTab[1638]"
//line /snap/go/10455/src/strings/strings.go:967
	}
//line /snap/go/10455/src/strings/strings.go:967
	// _ = "end of CoverTab[1630]"
//line /snap/go/10455/src/strings/strings.go:967
	_go_fuzz_dep_.CoverTab[1631]++
							if as, ok := makeASCIISet(cutset); ok {
//line /snap/go/10455/src/strings/strings.go:968
		_go_fuzz_dep_.CoverTab[525407]++
//line /snap/go/10455/src/strings/strings.go:968
		_go_fuzz_dep_.CoverTab[1639]++
								return trimRightASCII(s, &as)
//line /snap/go/10455/src/strings/strings.go:969
		// _ = "end of CoverTab[1639]"
	} else {
//line /snap/go/10455/src/strings/strings.go:970
		_go_fuzz_dep_.CoverTab[525408]++
//line /snap/go/10455/src/strings/strings.go:970
		_go_fuzz_dep_.CoverTab[1640]++
//line /snap/go/10455/src/strings/strings.go:970
		// _ = "end of CoverTab[1640]"
//line /snap/go/10455/src/strings/strings.go:970
	}
//line /snap/go/10455/src/strings/strings.go:970
	// _ = "end of CoverTab[1631]"
//line /snap/go/10455/src/strings/strings.go:970
	_go_fuzz_dep_.CoverTab[1632]++
							return trimRightUnicode(s, cutset)
//line /snap/go/10455/src/strings/strings.go:971
	// _ = "end of CoverTab[1632]"
}

func trimRightByte(s string, c byte) string {
//line /snap/go/10455/src/strings/strings.go:974
	_go_fuzz_dep_.CoverTab[1641]++
//line /snap/go/10455/src/strings/strings.go:974
	_go_fuzz_dep_.CoverTab[786544] = 0
							for len(s) > 0 && func() bool {
//line /snap/go/10455/src/strings/strings.go:975
		_go_fuzz_dep_.CoverTab[1643]++
//line /snap/go/10455/src/strings/strings.go:975
		return s[len(s)-1] == c
//line /snap/go/10455/src/strings/strings.go:975
		// _ = "end of CoverTab[1643]"
//line /snap/go/10455/src/strings/strings.go:975
	}() {
//line /snap/go/10455/src/strings/strings.go:975
		if _go_fuzz_dep_.CoverTab[786544] == 0 {
//line /snap/go/10455/src/strings/strings.go:975
			_go_fuzz_dep_.CoverTab[525631]++
//line /snap/go/10455/src/strings/strings.go:975
		} else {
//line /snap/go/10455/src/strings/strings.go:975
			_go_fuzz_dep_.CoverTab[525632]++
//line /snap/go/10455/src/strings/strings.go:975
		}
//line /snap/go/10455/src/strings/strings.go:975
		_go_fuzz_dep_.CoverTab[786544] = 1
//line /snap/go/10455/src/strings/strings.go:975
		_go_fuzz_dep_.CoverTab[1644]++
								s = s[:len(s)-1]
//line /snap/go/10455/src/strings/strings.go:976
		// _ = "end of CoverTab[1644]"
	}
//line /snap/go/10455/src/strings/strings.go:977
	if _go_fuzz_dep_.CoverTab[786544] == 0 {
//line /snap/go/10455/src/strings/strings.go:977
		_go_fuzz_dep_.CoverTab[525633]++
//line /snap/go/10455/src/strings/strings.go:977
	} else {
//line /snap/go/10455/src/strings/strings.go:977
		_go_fuzz_dep_.CoverTab[525634]++
//line /snap/go/10455/src/strings/strings.go:977
	}
//line /snap/go/10455/src/strings/strings.go:977
	// _ = "end of CoverTab[1641]"
//line /snap/go/10455/src/strings/strings.go:977
	_go_fuzz_dep_.CoverTab[1642]++
							return s
//line /snap/go/10455/src/strings/strings.go:978
	// _ = "end of CoverTab[1642]"
}

func trimRightASCII(s string, as *asciiSet) string {
//line /snap/go/10455/src/strings/strings.go:981
	_go_fuzz_dep_.CoverTab[1645]++
//line /snap/go/10455/src/strings/strings.go:981
	_go_fuzz_dep_.CoverTab[786545] = 0
							for len(s) > 0 {
//line /snap/go/10455/src/strings/strings.go:982
		if _go_fuzz_dep_.CoverTab[786545] == 0 {
//line /snap/go/10455/src/strings/strings.go:982
			_go_fuzz_dep_.CoverTab[525635]++
//line /snap/go/10455/src/strings/strings.go:982
		} else {
//line /snap/go/10455/src/strings/strings.go:982
			_go_fuzz_dep_.CoverTab[525636]++
//line /snap/go/10455/src/strings/strings.go:982
		}
//line /snap/go/10455/src/strings/strings.go:982
		_go_fuzz_dep_.CoverTab[786545] = 1
//line /snap/go/10455/src/strings/strings.go:982
		_go_fuzz_dep_.CoverTab[1647]++
								if !as.contains(s[len(s)-1]) {
//line /snap/go/10455/src/strings/strings.go:983
			_go_fuzz_dep_.CoverTab[525409]++
//line /snap/go/10455/src/strings/strings.go:983
			_go_fuzz_dep_.CoverTab[1649]++
									break
//line /snap/go/10455/src/strings/strings.go:984
			// _ = "end of CoverTab[1649]"
		} else {
//line /snap/go/10455/src/strings/strings.go:985
			_go_fuzz_dep_.CoverTab[525410]++
//line /snap/go/10455/src/strings/strings.go:985
			_go_fuzz_dep_.CoverTab[1650]++
//line /snap/go/10455/src/strings/strings.go:985
			// _ = "end of CoverTab[1650]"
//line /snap/go/10455/src/strings/strings.go:985
		}
//line /snap/go/10455/src/strings/strings.go:985
		// _ = "end of CoverTab[1647]"
//line /snap/go/10455/src/strings/strings.go:985
		_go_fuzz_dep_.CoverTab[1648]++
								s = s[:len(s)-1]
//line /snap/go/10455/src/strings/strings.go:986
		// _ = "end of CoverTab[1648]"
	}
//line /snap/go/10455/src/strings/strings.go:987
	if _go_fuzz_dep_.CoverTab[786545] == 0 {
//line /snap/go/10455/src/strings/strings.go:987
		_go_fuzz_dep_.CoverTab[525637]++
//line /snap/go/10455/src/strings/strings.go:987
	} else {
//line /snap/go/10455/src/strings/strings.go:987
		_go_fuzz_dep_.CoverTab[525638]++
//line /snap/go/10455/src/strings/strings.go:987
	}
//line /snap/go/10455/src/strings/strings.go:987
	// _ = "end of CoverTab[1645]"
//line /snap/go/10455/src/strings/strings.go:987
	_go_fuzz_dep_.CoverTab[1646]++
							return s
//line /snap/go/10455/src/strings/strings.go:988
	// _ = "end of CoverTab[1646]"
}

func trimRightUnicode(s, cutset string) string {
//line /snap/go/10455/src/strings/strings.go:991
	_go_fuzz_dep_.CoverTab[1651]++
//line /snap/go/10455/src/strings/strings.go:991
	_go_fuzz_dep_.CoverTab[786546] = 0
							for len(s) > 0 {
//line /snap/go/10455/src/strings/strings.go:992
		if _go_fuzz_dep_.CoverTab[786546] == 0 {
//line /snap/go/10455/src/strings/strings.go:992
			_go_fuzz_dep_.CoverTab[525639]++
//line /snap/go/10455/src/strings/strings.go:992
		} else {
//line /snap/go/10455/src/strings/strings.go:992
			_go_fuzz_dep_.CoverTab[525640]++
//line /snap/go/10455/src/strings/strings.go:992
		}
//line /snap/go/10455/src/strings/strings.go:992
		_go_fuzz_dep_.CoverTab[786546] = 1
//line /snap/go/10455/src/strings/strings.go:992
		_go_fuzz_dep_.CoverTab[1653]++
								r, n := rune(s[len(s)-1]), 1
								if r >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:994
			_go_fuzz_dep_.CoverTab[525411]++
//line /snap/go/10455/src/strings/strings.go:994
			_go_fuzz_dep_.CoverTab[1656]++
									r, n = utf8.DecodeLastRuneInString(s)
//line /snap/go/10455/src/strings/strings.go:995
			// _ = "end of CoverTab[1656]"
		} else {
//line /snap/go/10455/src/strings/strings.go:996
			_go_fuzz_dep_.CoverTab[525412]++
//line /snap/go/10455/src/strings/strings.go:996
			_go_fuzz_dep_.CoverTab[1657]++
//line /snap/go/10455/src/strings/strings.go:996
			// _ = "end of CoverTab[1657]"
//line /snap/go/10455/src/strings/strings.go:996
		}
//line /snap/go/10455/src/strings/strings.go:996
		// _ = "end of CoverTab[1653]"
//line /snap/go/10455/src/strings/strings.go:996
		_go_fuzz_dep_.CoverTab[1654]++
								if !ContainsRune(cutset, r) {
//line /snap/go/10455/src/strings/strings.go:997
			_go_fuzz_dep_.CoverTab[525413]++
//line /snap/go/10455/src/strings/strings.go:997
			_go_fuzz_dep_.CoverTab[1658]++
									break
//line /snap/go/10455/src/strings/strings.go:998
			// _ = "end of CoverTab[1658]"
		} else {
//line /snap/go/10455/src/strings/strings.go:999
			_go_fuzz_dep_.CoverTab[525414]++
//line /snap/go/10455/src/strings/strings.go:999
			_go_fuzz_dep_.CoverTab[1659]++
//line /snap/go/10455/src/strings/strings.go:999
			// _ = "end of CoverTab[1659]"
//line /snap/go/10455/src/strings/strings.go:999
		}
//line /snap/go/10455/src/strings/strings.go:999
		// _ = "end of CoverTab[1654]"
//line /snap/go/10455/src/strings/strings.go:999
		_go_fuzz_dep_.CoverTab[1655]++
								s = s[:len(s)-n]
//line /snap/go/10455/src/strings/strings.go:1000
		// _ = "end of CoverTab[1655]"
	}
//line /snap/go/10455/src/strings/strings.go:1001
	if _go_fuzz_dep_.CoverTab[786546] == 0 {
//line /snap/go/10455/src/strings/strings.go:1001
		_go_fuzz_dep_.CoverTab[525641]++
//line /snap/go/10455/src/strings/strings.go:1001
	} else {
//line /snap/go/10455/src/strings/strings.go:1001
		_go_fuzz_dep_.CoverTab[525642]++
//line /snap/go/10455/src/strings/strings.go:1001
	}
//line /snap/go/10455/src/strings/strings.go:1001
	// _ = "end of CoverTab[1651]"
//line /snap/go/10455/src/strings/strings.go:1001
	_go_fuzz_dep_.CoverTab[1652]++
							return s
//line /snap/go/10455/src/strings/strings.go:1002
	// _ = "end of CoverTab[1652]"
}

// TrimSpace returns a slice of the string s, with all leading
//line /snap/go/10455/src/strings/strings.go:1005
// and trailing white space removed, as defined by Unicode.
//line /snap/go/10455/src/strings/strings.go:1007
func TrimSpace(s string) string {
//line /snap/go/10455/src/strings/strings.go:1007
	_go_fuzz_dep_.CoverTab[1660]++

							start := 0
//line /snap/go/10455/src/strings/strings.go:1009
	_go_fuzz_dep_.CoverTab[786547] = 0
							for ; start < len(s); start++ {
//line /snap/go/10455/src/strings/strings.go:1010
		if _go_fuzz_dep_.CoverTab[786547] == 0 {
//line /snap/go/10455/src/strings/strings.go:1010
			_go_fuzz_dep_.CoverTab[525643]++
//line /snap/go/10455/src/strings/strings.go:1010
		} else {
//line /snap/go/10455/src/strings/strings.go:1010
			_go_fuzz_dep_.CoverTab[525644]++
//line /snap/go/10455/src/strings/strings.go:1010
		}
//line /snap/go/10455/src/strings/strings.go:1010
		_go_fuzz_dep_.CoverTab[786547] = 1
//line /snap/go/10455/src/strings/strings.go:1010
		_go_fuzz_dep_.CoverTab[1663]++
								c := s[start]
								if c >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:1012
			_go_fuzz_dep_.CoverTab[525415]++
//line /snap/go/10455/src/strings/strings.go:1012
			_go_fuzz_dep_.CoverTab[1665]++

//line /snap/go/10455/src/strings/strings.go:1015
			return TrimFunc(s[start:], unicode.IsSpace)
//line /snap/go/10455/src/strings/strings.go:1015
			// _ = "end of CoverTab[1665]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1016
			_go_fuzz_dep_.CoverTab[525416]++
//line /snap/go/10455/src/strings/strings.go:1016
			_go_fuzz_dep_.CoverTab[1666]++
//line /snap/go/10455/src/strings/strings.go:1016
			// _ = "end of CoverTab[1666]"
//line /snap/go/10455/src/strings/strings.go:1016
		}
//line /snap/go/10455/src/strings/strings.go:1016
		// _ = "end of CoverTab[1663]"
//line /snap/go/10455/src/strings/strings.go:1016
		_go_fuzz_dep_.CoverTab[1664]++
								if asciiSpace[c] == 0 {
//line /snap/go/10455/src/strings/strings.go:1017
			_go_fuzz_dep_.CoverTab[525417]++
//line /snap/go/10455/src/strings/strings.go:1017
			_go_fuzz_dep_.CoverTab[1667]++
									break
//line /snap/go/10455/src/strings/strings.go:1018
			// _ = "end of CoverTab[1667]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1019
			_go_fuzz_dep_.CoverTab[525418]++
//line /snap/go/10455/src/strings/strings.go:1019
			_go_fuzz_dep_.CoverTab[1668]++
//line /snap/go/10455/src/strings/strings.go:1019
			// _ = "end of CoverTab[1668]"
//line /snap/go/10455/src/strings/strings.go:1019
		}
//line /snap/go/10455/src/strings/strings.go:1019
		// _ = "end of CoverTab[1664]"
	}
//line /snap/go/10455/src/strings/strings.go:1020
	if _go_fuzz_dep_.CoverTab[786547] == 0 {
//line /snap/go/10455/src/strings/strings.go:1020
		_go_fuzz_dep_.CoverTab[525645]++
//line /snap/go/10455/src/strings/strings.go:1020
	} else {
//line /snap/go/10455/src/strings/strings.go:1020
		_go_fuzz_dep_.CoverTab[525646]++
//line /snap/go/10455/src/strings/strings.go:1020
	}
//line /snap/go/10455/src/strings/strings.go:1020
	// _ = "end of CoverTab[1660]"
//line /snap/go/10455/src/strings/strings.go:1020
	_go_fuzz_dep_.CoverTab[1661]++

//line /snap/go/10455/src/strings/strings.go:1023
	stop := len(s)
//line /snap/go/10455/src/strings/strings.go:1023
	_go_fuzz_dep_.CoverTab[786548] = 0
							for ; stop > start; stop-- {
//line /snap/go/10455/src/strings/strings.go:1024
		if _go_fuzz_dep_.CoverTab[786548] == 0 {
//line /snap/go/10455/src/strings/strings.go:1024
			_go_fuzz_dep_.CoverTab[525647]++
//line /snap/go/10455/src/strings/strings.go:1024
		} else {
//line /snap/go/10455/src/strings/strings.go:1024
			_go_fuzz_dep_.CoverTab[525648]++
//line /snap/go/10455/src/strings/strings.go:1024
		}
//line /snap/go/10455/src/strings/strings.go:1024
		_go_fuzz_dep_.CoverTab[786548] = 1
//line /snap/go/10455/src/strings/strings.go:1024
		_go_fuzz_dep_.CoverTab[1669]++
								c := s[stop-1]
								if c >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:1026
			_go_fuzz_dep_.CoverTab[525419]++
//line /snap/go/10455/src/strings/strings.go:1026
			_go_fuzz_dep_.CoverTab[1671]++

									return TrimRightFunc(s[start:stop], unicode.IsSpace)
//line /snap/go/10455/src/strings/strings.go:1028
			// _ = "end of CoverTab[1671]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1029
			_go_fuzz_dep_.CoverTab[525420]++
//line /snap/go/10455/src/strings/strings.go:1029
			_go_fuzz_dep_.CoverTab[1672]++
//line /snap/go/10455/src/strings/strings.go:1029
			// _ = "end of CoverTab[1672]"
//line /snap/go/10455/src/strings/strings.go:1029
		}
//line /snap/go/10455/src/strings/strings.go:1029
		// _ = "end of CoverTab[1669]"
//line /snap/go/10455/src/strings/strings.go:1029
		_go_fuzz_dep_.CoverTab[1670]++
								if asciiSpace[c] == 0 {
//line /snap/go/10455/src/strings/strings.go:1030
			_go_fuzz_dep_.CoverTab[525421]++
//line /snap/go/10455/src/strings/strings.go:1030
			_go_fuzz_dep_.CoverTab[1673]++
									break
//line /snap/go/10455/src/strings/strings.go:1031
			// _ = "end of CoverTab[1673]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1032
			_go_fuzz_dep_.CoverTab[525422]++
//line /snap/go/10455/src/strings/strings.go:1032
			_go_fuzz_dep_.CoverTab[1674]++
//line /snap/go/10455/src/strings/strings.go:1032
			// _ = "end of CoverTab[1674]"
//line /snap/go/10455/src/strings/strings.go:1032
		}
//line /snap/go/10455/src/strings/strings.go:1032
		// _ = "end of CoverTab[1670]"
	}
//line /snap/go/10455/src/strings/strings.go:1033
	if _go_fuzz_dep_.CoverTab[786548] == 0 {
//line /snap/go/10455/src/strings/strings.go:1033
		_go_fuzz_dep_.CoverTab[525649]++
//line /snap/go/10455/src/strings/strings.go:1033
	} else {
//line /snap/go/10455/src/strings/strings.go:1033
		_go_fuzz_dep_.CoverTab[525650]++
//line /snap/go/10455/src/strings/strings.go:1033
	}
//line /snap/go/10455/src/strings/strings.go:1033
	// _ = "end of CoverTab[1661]"
//line /snap/go/10455/src/strings/strings.go:1033
	_go_fuzz_dep_.CoverTab[1662]++

//line /snap/go/10455/src/strings/strings.go:1038
	return s[start:stop]
//line /snap/go/10455/src/strings/strings.go:1038
	// _ = "end of CoverTab[1662]"
}

// TrimPrefix returns s without the provided leading prefix string.
//line /snap/go/10455/src/strings/strings.go:1041
// If s doesn't start with prefix, s is returned unchanged.
//line /snap/go/10455/src/strings/strings.go:1043
func TrimPrefix(s, prefix string) string {
//line /snap/go/10455/src/strings/strings.go:1043
	_go_fuzz_dep_.CoverTab[1675]++
							if HasPrefix(s, prefix) {
//line /snap/go/10455/src/strings/strings.go:1044
		_go_fuzz_dep_.CoverTab[525423]++
//line /snap/go/10455/src/strings/strings.go:1044
		_go_fuzz_dep_.CoverTab[1677]++
								return s[len(prefix):]
//line /snap/go/10455/src/strings/strings.go:1045
		// _ = "end of CoverTab[1677]"
	} else {
//line /snap/go/10455/src/strings/strings.go:1046
		_go_fuzz_dep_.CoverTab[525424]++
//line /snap/go/10455/src/strings/strings.go:1046
		_go_fuzz_dep_.CoverTab[1678]++
//line /snap/go/10455/src/strings/strings.go:1046
		// _ = "end of CoverTab[1678]"
//line /snap/go/10455/src/strings/strings.go:1046
	}
//line /snap/go/10455/src/strings/strings.go:1046
	// _ = "end of CoverTab[1675]"
//line /snap/go/10455/src/strings/strings.go:1046
	_go_fuzz_dep_.CoverTab[1676]++
							return s
//line /snap/go/10455/src/strings/strings.go:1047
	// _ = "end of CoverTab[1676]"
}

// TrimSuffix returns s without the provided trailing suffix string.
//line /snap/go/10455/src/strings/strings.go:1050
// If s doesn't end with suffix, s is returned unchanged.
//line /snap/go/10455/src/strings/strings.go:1052
func TrimSuffix(s, suffix string) string {
//line /snap/go/10455/src/strings/strings.go:1052
	_go_fuzz_dep_.CoverTab[1679]++
							if HasSuffix(s, suffix) {
//line /snap/go/10455/src/strings/strings.go:1053
		_go_fuzz_dep_.CoverTab[525425]++
//line /snap/go/10455/src/strings/strings.go:1053
		_go_fuzz_dep_.CoverTab[1681]++
								return s[:len(s)-len(suffix)]
//line /snap/go/10455/src/strings/strings.go:1054
		// _ = "end of CoverTab[1681]"
	} else {
//line /snap/go/10455/src/strings/strings.go:1055
		_go_fuzz_dep_.CoverTab[525426]++
//line /snap/go/10455/src/strings/strings.go:1055
		_go_fuzz_dep_.CoverTab[1682]++
//line /snap/go/10455/src/strings/strings.go:1055
		// _ = "end of CoverTab[1682]"
//line /snap/go/10455/src/strings/strings.go:1055
	}
//line /snap/go/10455/src/strings/strings.go:1055
	// _ = "end of CoverTab[1679]"
//line /snap/go/10455/src/strings/strings.go:1055
	_go_fuzz_dep_.CoverTab[1680]++
							return s
//line /snap/go/10455/src/strings/strings.go:1056
	// _ = "end of CoverTab[1680]"
}

// Replace returns a copy of the string s with the first n
//line /snap/go/10455/src/strings/strings.go:1059
// non-overlapping instances of old replaced by new.
//line /snap/go/10455/src/strings/strings.go:1059
// If old is empty, it matches at the beginning of the string
//line /snap/go/10455/src/strings/strings.go:1059
// and after each UTF-8 sequence, yielding up to k+1 replacements
//line /snap/go/10455/src/strings/strings.go:1059
// for a k-rune string.
//line /snap/go/10455/src/strings/strings.go:1059
// If n < 0, there is no limit on the number of replacements.
//line /snap/go/10455/src/strings/strings.go:1065
func Replace(s, old, new string, n int) string {
//line /snap/go/10455/src/strings/strings.go:1065
	_go_fuzz_dep_.CoverTab[1683]++
							if old == new || func() bool {
//line /snap/go/10455/src/strings/strings.go:1066
		_go_fuzz_dep_.CoverTab[1687]++
//line /snap/go/10455/src/strings/strings.go:1066
		return n == 0
//line /snap/go/10455/src/strings/strings.go:1066
		// _ = "end of CoverTab[1687]"
//line /snap/go/10455/src/strings/strings.go:1066
	}() {
//line /snap/go/10455/src/strings/strings.go:1066
		_go_fuzz_dep_.CoverTab[525427]++
//line /snap/go/10455/src/strings/strings.go:1066
		_go_fuzz_dep_.CoverTab[1688]++
								return s
//line /snap/go/10455/src/strings/strings.go:1067
		// _ = "end of CoverTab[1688]"
	} else {
//line /snap/go/10455/src/strings/strings.go:1068
		_go_fuzz_dep_.CoverTab[525428]++
//line /snap/go/10455/src/strings/strings.go:1068
		_go_fuzz_dep_.CoverTab[1689]++
//line /snap/go/10455/src/strings/strings.go:1068
		// _ = "end of CoverTab[1689]"
//line /snap/go/10455/src/strings/strings.go:1068
	}
//line /snap/go/10455/src/strings/strings.go:1068
	// _ = "end of CoverTab[1683]"
//line /snap/go/10455/src/strings/strings.go:1068
	_go_fuzz_dep_.CoverTab[1684]++

//line /snap/go/10455/src/strings/strings.go:1071
	if m := Count(s, old); m == 0 {
//line /snap/go/10455/src/strings/strings.go:1071
		_go_fuzz_dep_.CoverTab[525429]++
//line /snap/go/10455/src/strings/strings.go:1071
		_go_fuzz_dep_.CoverTab[1690]++
								return s
//line /snap/go/10455/src/strings/strings.go:1072
		// _ = "end of CoverTab[1690]"
	} else {
//line /snap/go/10455/src/strings/strings.go:1073
		_go_fuzz_dep_.CoverTab[525430]++
//line /snap/go/10455/src/strings/strings.go:1073
		_go_fuzz_dep_.CoverTab[1691]++
//line /snap/go/10455/src/strings/strings.go:1073
		if n < 0 || func() bool {
//line /snap/go/10455/src/strings/strings.go:1073
			_go_fuzz_dep_.CoverTab[1692]++
//line /snap/go/10455/src/strings/strings.go:1073
			return m < n
//line /snap/go/10455/src/strings/strings.go:1073
			// _ = "end of CoverTab[1692]"
//line /snap/go/10455/src/strings/strings.go:1073
		}() {
//line /snap/go/10455/src/strings/strings.go:1073
			_go_fuzz_dep_.CoverTab[525431]++
//line /snap/go/10455/src/strings/strings.go:1073
			_go_fuzz_dep_.CoverTab[1693]++
									n = m
//line /snap/go/10455/src/strings/strings.go:1074
			// _ = "end of CoverTab[1693]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1075
			_go_fuzz_dep_.CoverTab[525432]++
//line /snap/go/10455/src/strings/strings.go:1075
			_go_fuzz_dep_.CoverTab[1694]++
//line /snap/go/10455/src/strings/strings.go:1075
			// _ = "end of CoverTab[1694]"
//line /snap/go/10455/src/strings/strings.go:1075
		}
//line /snap/go/10455/src/strings/strings.go:1075
		// _ = "end of CoverTab[1691]"
//line /snap/go/10455/src/strings/strings.go:1075
	}
//line /snap/go/10455/src/strings/strings.go:1075
	// _ = "end of CoverTab[1684]"
//line /snap/go/10455/src/strings/strings.go:1075
	_go_fuzz_dep_.CoverTab[1685]++

							// Apply replacements to buffer.
							var b Builder
							b.Grow(len(s) + n*(len(new)-len(old)))
							start := 0
//line /snap/go/10455/src/strings/strings.go:1080
	_go_fuzz_dep_.CoverTab[786549] = 0
							for i := 0; i < n; i++ {
//line /snap/go/10455/src/strings/strings.go:1081
		if _go_fuzz_dep_.CoverTab[786549] == 0 {
//line /snap/go/10455/src/strings/strings.go:1081
			_go_fuzz_dep_.CoverTab[525651]++
//line /snap/go/10455/src/strings/strings.go:1081
		} else {
//line /snap/go/10455/src/strings/strings.go:1081
			_go_fuzz_dep_.CoverTab[525652]++
//line /snap/go/10455/src/strings/strings.go:1081
		}
//line /snap/go/10455/src/strings/strings.go:1081
		_go_fuzz_dep_.CoverTab[786549] = 1
//line /snap/go/10455/src/strings/strings.go:1081
		_go_fuzz_dep_.CoverTab[1695]++
								j := start
								if len(old) == 0 {
//line /snap/go/10455/src/strings/strings.go:1083
			_go_fuzz_dep_.CoverTab[525433]++
//line /snap/go/10455/src/strings/strings.go:1083
			_go_fuzz_dep_.CoverTab[1697]++
									if i > 0 {
//line /snap/go/10455/src/strings/strings.go:1084
				_go_fuzz_dep_.CoverTab[525435]++
//line /snap/go/10455/src/strings/strings.go:1084
				_go_fuzz_dep_.CoverTab[1698]++
										_, wid := utf8.DecodeRuneInString(s[start:])
										j += wid
//line /snap/go/10455/src/strings/strings.go:1086
				// _ = "end of CoverTab[1698]"
			} else {
//line /snap/go/10455/src/strings/strings.go:1087
				_go_fuzz_dep_.CoverTab[525436]++
//line /snap/go/10455/src/strings/strings.go:1087
				_go_fuzz_dep_.CoverTab[1699]++
//line /snap/go/10455/src/strings/strings.go:1087
				// _ = "end of CoverTab[1699]"
//line /snap/go/10455/src/strings/strings.go:1087
			}
//line /snap/go/10455/src/strings/strings.go:1087
			// _ = "end of CoverTab[1697]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1088
			_go_fuzz_dep_.CoverTab[525434]++
//line /snap/go/10455/src/strings/strings.go:1088
			_go_fuzz_dep_.CoverTab[1700]++
									j += Index(s[start:], old)
//line /snap/go/10455/src/strings/strings.go:1089
			// _ = "end of CoverTab[1700]"
		}
//line /snap/go/10455/src/strings/strings.go:1090
		// _ = "end of CoverTab[1695]"
//line /snap/go/10455/src/strings/strings.go:1090
		_go_fuzz_dep_.CoverTab[1696]++
								b.WriteString(s[start:j])
								b.WriteString(new)
								start = j + len(old)
//line /snap/go/10455/src/strings/strings.go:1093
		// _ = "end of CoverTab[1696]"
	}
//line /snap/go/10455/src/strings/strings.go:1094
	if _go_fuzz_dep_.CoverTab[786549] == 0 {
//line /snap/go/10455/src/strings/strings.go:1094
		_go_fuzz_dep_.CoverTab[525653]++
//line /snap/go/10455/src/strings/strings.go:1094
	} else {
//line /snap/go/10455/src/strings/strings.go:1094
		_go_fuzz_dep_.CoverTab[525654]++
//line /snap/go/10455/src/strings/strings.go:1094
	}
//line /snap/go/10455/src/strings/strings.go:1094
	// _ = "end of CoverTab[1685]"
//line /snap/go/10455/src/strings/strings.go:1094
	_go_fuzz_dep_.CoverTab[1686]++
							b.WriteString(s[start:])
							return b.String()
//line /snap/go/10455/src/strings/strings.go:1096
	// _ = "end of CoverTab[1686]"
}

// ReplaceAll returns a copy of the string s with all
//line /snap/go/10455/src/strings/strings.go:1099
// non-overlapping instances of old replaced by new.
//line /snap/go/10455/src/strings/strings.go:1099
// If old is empty, it matches at the beginning of the string
//line /snap/go/10455/src/strings/strings.go:1099
// and after each UTF-8 sequence, yielding up to k+1 replacements
//line /snap/go/10455/src/strings/strings.go:1099
// for a k-rune string.
//line /snap/go/10455/src/strings/strings.go:1104
func ReplaceAll(s, old, new string) string {
//line /snap/go/10455/src/strings/strings.go:1104
	_go_fuzz_dep_.CoverTab[1701]++
							return Replace(s, old, new, -1)
//line /snap/go/10455/src/strings/strings.go:1105
	// _ = "end of CoverTab[1701]"
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
//line /snap/go/10455/src/strings/strings.go:1108
// are equal under simple Unicode case-folding, which is a more general
//line /snap/go/10455/src/strings/strings.go:1108
// form of case-insensitivity.
//line /snap/go/10455/src/strings/strings.go:1111
func EqualFold(s, t string) bool {
//line /snap/go/10455/src/strings/strings.go:1111
	_go_fuzz_dep_.CoverTab[1702]++

							i := 0
//line /snap/go/10455/src/strings/strings.go:1113
	_go_fuzz_dep_.CoverTab[786550] = 0
							for ; i < len(s) && func() bool {
//line /snap/go/10455/src/strings/strings.go:1114
		_go_fuzz_dep_.CoverTab[1705]++
//line /snap/go/10455/src/strings/strings.go:1114
		return i < len(t)
//line /snap/go/10455/src/strings/strings.go:1114
		// _ = "end of CoverTab[1705]"
//line /snap/go/10455/src/strings/strings.go:1114
	}(); i++ {
//line /snap/go/10455/src/strings/strings.go:1114
		if _go_fuzz_dep_.CoverTab[786550] == 0 {
//line /snap/go/10455/src/strings/strings.go:1114
			_go_fuzz_dep_.CoverTab[525655]++
//line /snap/go/10455/src/strings/strings.go:1114
		} else {
//line /snap/go/10455/src/strings/strings.go:1114
			_go_fuzz_dep_.CoverTab[525656]++
//line /snap/go/10455/src/strings/strings.go:1114
		}
//line /snap/go/10455/src/strings/strings.go:1114
		_go_fuzz_dep_.CoverTab[786550] = 1
//line /snap/go/10455/src/strings/strings.go:1114
		_go_fuzz_dep_.CoverTab[1706]++
								sr := s[i]
								tr := t[i]
								if sr|tr >= utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:1117
			_go_fuzz_dep_.CoverTab[525437]++
//line /snap/go/10455/src/strings/strings.go:1117
			_go_fuzz_dep_.CoverTab[1711]++
									goto hasUnicode
//line /snap/go/10455/src/strings/strings.go:1118
			// _ = "end of CoverTab[1711]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1119
			_go_fuzz_dep_.CoverTab[525438]++
//line /snap/go/10455/src/strings/strings.go:1119
			_go_fuzz_dep_.CoverTab[1712]++
//line /snap/go/10455/src/strings/strings.go:1119
			// _ = "end of CoverTab[1712]"
//line /snap/go/10455/src/strings/strings.go:1119
		}
//line /snap/go/10455/src/strings/strings.go:1119
		// _ = "end of CoverTab[1706]"
//line /snap/go/10455/src/strings/strings.go:1119
		_go_fuzz_dep_.CoverTab[1707]++

//line /snap/go/10455/src/strings/strings.go:1122
		if tr == sr {
//line /snap/go/10455/src/strings/strings.go:1122
			_go_fuzz_dep_.CoverTab[525439]++
//line /snap/go/10455/src/strings/strings.go:1122
			_go_fuzz_dep_.CoverTab[1713]++
									continue
//line /snap/go/10455/src/strings/strings.go:1123
			// _ = "end of CoverTab[1713]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1124
			_go_fuzz_dep_.CoverTab[525440]++
//line /snap/go/10455/src/strings/strings.go:1124
			_go_fuzz_dep_.CoverTab[1714]++
//line /snap/go/10455/src/strings/strings.go:1124
			// _ = "end of CoverTab[1714]"
//line /snap/go/10455/src/strings/strings.go:1124
		}
//line /snap/go/10455/src/strings/strings.go:1124
		// _ = "end of CoverTab[1707]"
//line /snap/go/10455/src/strings/strings.go:1124
		_go_fuzz_dep_.CoverTab[1708]++

//line /snap/go/10455/src/strings/strings.go:1127
		if tr < sr {
//line /snap/go/10455/src/strings/strings.go:1127
			_go_fuzz_dep_.CoverTab[525441]++
//line /snap/go/10455/src/strings/strings.go:1127
			_go_fuzz_dep_.CoverTab[1715]++
									tr, sr = sr, tr
//line /snap/go/10455/src/strings/strings.go:1128
			// _ = "end of CoverTab[1715]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1129
			_go_fuzz_dep_.CoverTab[525442]++
//line /snap/go/10455/src/strings/strings.go:1129
			_go_fuzz_dep_.CoverTab[1716]++
//line /snap/go/10455/src/strings/strings.go:1129
			// _ = "end of CoverTab[1716]"
//line /snap/go/10455/src/strings/strings.go:1129
		}
//line /snap/go/10455/src/strings/strings.go:1129
		// _ = "end of CoverTab[1708]"
//line /snap/go/10455/src/strings/strings.go:1129
		_go_fuzz_dep_.CoverTab[1709]++

								if 'A' <= sr && func() bool {
//line /snap/go/10455/src/strings/strings.go:1131
			_go_fuzz_dep_.CoverTab[1717]++
//line /snap/go/10455/src/strings/strings.go:1131
			return sr <= 'Z'
//line /snap/go/10455/src/strings/strings.go:1131
			// _ = "end of CoverTab[1717]"
//line /snap/go/10455/src/strings/strings.go:1131
		}() && func() bool {
//line /snap/go/10455/src/strings/strings.go:1131
			_go_fuzz_dep_.CoverTab[1718]++
//line /snap/go/10455/src/strings/strings.go:1131
			return tr == sr+'a'-'A'
//line /snap/go/10455/src/strings/strings.go:1131
			// _ = "end of CoverTab[1718]"
//line /snap/go/10455/src/strings/strings.go:1131
		}() {
//line /snap/go/10455/src/strings/strings.go:1131
			_go_fuzz_dep_.CoverTab[525443]++
//line /snap/go/10455/src/strings/strings.go:1131
			_go_fuzz_dep_.CoverTab[1719]++
									continue
//line /snap/go/10455/src/strings/strings.go:1132
			// _ = "end of CoverTab[1719]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1133
			_go_fuzz_dep_.CoverTab[525444]++
//line /snap/go/10455/src/strings/strings.go:1133
			_go_fuzz_dep_.CoverTab[1720]++
//line /snap/go/10455/src/strings/strings.go:1133
			// _ = "end of CoverTab[1720]"
//line /snap/go/10455/src/strings/strings.go:1133
		}
//line /snap/go/10455/src/strings/strings.go:1133
		// _ = "end of CoverTab[1709]"
//line /snap/go/10455/src/strings/strings.go:1133
		_go_fuzz_dep_.CoverTab[1710]++
								return false
//line /snap/go/10455/src/strings/strings.go:1134
		// _ = "end of CoverTab[1710]"
	}
//line /snap/go/10455/src/strings/strings.go:1135
	if _go_fuzz_dep_.CoverTab[786550] == 0 {
//line /snap/go/10455/src/strings/strings.go:1135
		_go_fuzz_dep_.CoverTab[525657]++
//line /snap/go/10455/src/strings/strings.go:1135
	} else {
//line /snap/go/10455/src/strings/strings.go:1135
		_go_fuzz_dep_.CoverTab[525658]++
//line /snap/go/10455/src/strings/strings.go:1135
	}
//line /snap/go/10455/src/strings/strings.go:1135
	// _ = "end of CoverTab[1702]"
//line /snap/go/10455/src/strings/strings.go:1135
	_go_fuzz_dep_.CoverTab[1703]++

							return len(s) == len(t)

hasUnicode:
							s = s[i:]
							t = t[i:]
//line /snap/go/10455/src/strings/strings.go:1141
	_go_fuzz_dep_.CoverTab[786551] = 0
							for _, sr := range s {
//line /snap/go/10455/src/strings/strings.go:1142
		if _go_fuzz_dep_.CoverTab[786551] == 0 {
//line /snap/go/10455/src/strings/strings.go:1142
			_go_fuzz_dep_.CoverTab[525659]++
//line /snap/go/10455/src/strings/strings.go:1142
		} else {
//line /snap/go/10455/src/strings/strings.go:1142
			_go_fuzz_dep_.CoverTab[525660]++
//line /snap/go/10455/src/strings/strings.go:1142
		}
//line /snap/go/10455/src/strings/strings.go:1142
		_go_fuzz_dep_.CoverTab[786551] = 1
//line /snap/go/10455/src/strings/strings.go:1142
		_go_fuzz_dep_.CoverTab[1721]++

								if len(t) == 0 {
//line /snap/go/10455/src/strings/strings.go:1144
			_go_fuzz_dep_.CoverTab[525445]++
//line /snap/go/10455/src/strings/strings.go:1144
			_go_fuzz_dep_.CoverTab[1729]++
									return false
//line /snap/go/10455/src/strings/strings.go:1145
			// _ = "end of CoverTab[1729]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1146
			_go_fuzz_dep_.CoverTab[525446]++
//line /snap/go/10455/src/strings/strings.go:1146
			_go_fuzz_dep_.CoverTab[1730]++
//line /snap/go/10455/src/strings/strings.go:1146
			// _ = "end of CoverTab[1730]"
//line /snap/go/10455/src/strings/strings.go:1146
		}
//line /snap/go/10455/src/strings/strings.go:1146
		// _ = "end of CoverTab[1721]"
//line /snap/go/10455/src/strings/strings.go:1146
		_go_fuzz_dep_.CoverTab[1722]++

		// Extract first rune from second string.
		var tr rune
		if t[0] < utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:1150
			_go_fuzz_dep_.CoverTab[525447]++
//line /snap/go/10455/src/strings/strings.go:1150
			_go_fuzz_dep_.CoverTab[1731]++
									tr, t = rune(t[0]), t[1:]
//line /snap/go/10455/src/strings/strings.go:1151
			// _ = "end of CoverTab[1731]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1152
			_go_fuzz_dep_.CoverTab[525448]++
//line /snap/go/10455/src/strings/strings.go:1152
			_go_fuzz_dep_.CoverTab[1732]++
									r, size := utf8.DecodeRuneInString(t)
									tr, t = r, t[size:]
//line /snap/go/10455/src/strings/strings.go:1154
			// _ = "end of CoverTab[1732]"
		}
//line /snap/go/10455/src/strings/strings.go:1155
		// _ = "end of CoverTab[1722]"
//line /snap/go/10455/src/strings/strings.go:1155
		_go_fuzz_dep_.CoverTab[1723]++

//line /snap/go/10455/src/strings/strings.go:1160
		if tr == sr {
//line /snap/go/10455/src/strings/strings.go:1160
			_go_fuzz_dep_.CoverTab[525449]++
//line /snap/go/10455/src/strings/strings.go:1160
			_go_fuzz_dep_.CoverTab[1733]++
									continue
//line /snap/go/10455/src/strings/strings.go:1161
			// _ = "end of CoverTab[1733]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1162
			_go_fuzz_dep_.CoverTab[525450]++
//line /snap/go/10455/src/strings/strings.go:1162
			_go_fuzz_dep_.CoverTab[1734]++
//line /snap/go/10455/src/strings/strings.go:1162
			// _ = "end of CoverTab[1734]"
//line /snap/go/10455/src/strings/strings.go:1162
		}
//line /snap/go/10455/src/strings/strings.go:1162
		// _ = "end of CoverTab[1723]"
//line /snap/go/10455/src/strings/strings.go:1162
		_go_fuzz_dep_.CoverTab[1724]++

//line /snap/go/10455/src/strings/strings.go:1165
		if tr < sr {
//line /snap/go/10455/src/strings/strings.go:1165
			_go_fuzz_dep_.CoverTab[525451]++
//line /snap/go/10455/src/strings/strings.go:1165
			_go_fuzz_dep_.CoverTab[1735]++
									tr, sr = sr, tr
//line /snap/go/10455/src/strings/strings.go:1166
			// _ = "end of CoverTab[1735]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1167
			_go_fuzz_dep_.CoverTab[525452]++
//line /snap/go/10455/src/strings/strings.go:1167
			_go_fuzz_dep_.CoverTab[1736]++
//line /snap/go/10455/src/strings/strings.go:1167
			// _ = "end of CoverTab[1736]"
//line /snap/go/10455/src/strings/strings.go:1167
		}
//line /snap/go/10455/src/strings/strings.go:1167
		// _ = "end of CoverTab[1724]"
//line /snap/go/10455/src/strings/strings.go:1167
		_go_fuzz_dep_.CoverTab[1725]++

								if tr < utf8.RuneSelf {
//line /snap/go/10455/src/strings/strings.go:1169
			_go_fuzz_dep_.CoverTab[525453]++
//line /snap/go/10455/src/strings/strings.go:1169
			_go_fuzz_dep_.CoverTab[1737]++

									if 'A' <= sr && func() bool {
//line /snap/go/10455/src/strings/strings.go:1171
				_go_fuzz_dep_.CoverTab[1739]++
//line /snap/go/10455/src/strings/strings.go:1171
				return sr <= 'Z'
//line /snap/go/10455/src/strings/strings.go:1171
				// _ = "end of CoverTab[1739]"
//line /snap/go/10455/src/strings/strings.go:1171
			}() && func() bool {
//line /snap/go/10455/src/strings/strings.go:1171
				_go_fuzz_dep_.CoverTab[1740]++
//line /snap/go/10455/src/strings/strings.go:1171
				return tr == sr+'a'-'A'
//line /snap/go/10455/src/strings/strings.go:1171
				// _ = "end of CoverTab[1740]"
//line /snap/go/10455/src/strings/strings.go:1171
			}() {
//line /snap/go/10455/src/strings/strings.go:1171
				_go_fuzz_dep_.CoverTab[525455]++
//line /snap/go/10455/src/strings/strings.go:1171
				_go_fuzz_dep_.CoverTab[1741]++
										continue
//line /snap/go/10455/src/strings/strings.go:1172
				// _ = "end of CoverTab[1741]"
			} else {
//line /snap/go/10455/src/strings/strings.go:1173
				_go_fuzz_dep_.CoverTab[525456]++
//line /snap/go/10455/src/strings/strings.go:1173
				_go_fuzz_dep_.CoverTab[1742]++
//line /snap/go/10455/src/strings/strings.go:1173
				// _ = "end of CoverTab[1742]"
//line /snap/go/10455/src/strings/strings.go:1173
			}
//line /snap/go/10455/src/strings/strings.go:1173
			// _ = "end of CoverTab[1737]"
//line /snap/go/10455/src/strings/strings.go:1173
			_go_fuzz_dep_.CoverTab[1738]++
									return false
//line /snap/go/10455/src/strings/strings.go:1174
			// _ = "end of CoverTab[1738]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1175
			_go_fuzz_dep_.CoverTab[525454]++
//line /snap/go/10455/src/strings/strings.go:1175
			_go_fuzz_dep_.CoverTab[1743]++
//line /snap/go/10455/src/strings/strings.go:1175
			// _ = "end of CoverTab[1743]"
//line /snap/go/10455/src/strings/strings.go:1175
		}
//line /snap/go/10455/src/strings/strings.go:1175
		// _ = "end of CoverTab[1725]"
//line /snap/go/10455/src/strings/strings.go:1175
		_go_fuzz_dep_.CoverTab[1726]++

//line /snap/go/10455/src/strings/strings.go:1179
		r := unicode.SimpleFold(sr)
//line /snap/go/10455/src/strings/strings.go:1179
		_go_fuzz_dep_.CoverTab[786552] = 0
								for r != sr && func() bool {
//line /snap/go/10455/src/strings/strings.go:1180
			_go_fuzz_dep_.CoverTab[1744]++
//line /snap/go/10455/src/strings/strings.go:1180
			return r < tr
//line /snap/go/10455/src/strings/strings.go:1180
			// _ = "end of CoverTab[1744]"
//line /snap/go/10455/src/strings/strings.go:1180
		}() {
//line /snap/go/10455/src/strings/strings.go:1180
			if _go_fuzz_dep_.CoverTab[786552] == 0 {
//line /snap/go/10455/src/strings/strings.go:1180
				_go_fuzz_dep_.CoverTab[525663]++
//line /snap/go/10455/src/strings/strings.go:1180
			} else {
//line /snap/go/10455/src/strings/strings.go:1180
				_go_fuzz_dep_.CoverTab[525664]++
//line /snap/go/10455/src/strings/strings.go:1180
			}
//line /snap/go/10455/src/strings/strings.go:1180
			_go_fuzz_dep_.CoverTab[786552] = 1
//line /snap/go/10455/src/strings/strings.go:1180
			_go_fuzz_dep_.CoverTab[1745]++
									r = unicode.SimpleFold(r)
//line /snap/go/10455/src/strings/strings.go:1181
			// _ = "end of CoverTab[1745]"
		}
//line /snap/go/10455/src/strings/strings.go:1182
		if _go_fuzz_dep_.CoverTab[786552] == 0 {
//line /snap/go/10455/src/strings/strings.go:1182
			_go_fuzz_dep_.CoverTab[525665]++
//line /snap/go/10455/src/strings/strings.go:1182
		} else {
//line /snap/go/10455/src/strings/strings.go:1182
			_go_fuzz_dep_.CoverTab[525666]++
//line /snap/go/10455/src/strings/strings.go:1182
		}
//line /snap/go/10455/src/strings/strings.go:1182
		// _ = "end of CoverTab[1726]"
//line /snap/go/10455/src/strings/strings.go:1182
		_go_fuzz_dep_.CoverTab[1727]++
								if r == tr {
//line /snap/go/10455/src/strings/strings.go:1183
			_go_fuzz_dep_.CoverTab[525457]++
//line /snap/go/10455/src/strings/strings.go:1183
			_go_fuzz_dep_.CoverTab[1746]++
									continue
//line /snap/go/10455/src/strings/strings.go:1184
			// _ = "end of CoverTab[1746]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1185
			_go_fuzz_dep_.CoverTab[525458]++
//line /snap/go/10455/src/strings/strings.go:1185
			_go_fuzz_dep_.CoverTab[1747]++
//line /snap/go/10455/src/strings/strings.go:1185
			// _ = "end of CoverTab[1747]"
//line /snap/go/10455/src/strings/strings.go:1185
		}
//line /snap/go/10455/src/strings/strings.go:1185
		// _ = "end of CoverTab[1727]"
//line /snap/go/10455/src/strings/strings.go:1185
		_go_fuzz_dep_.CoverTab[1728]++
								return false
//line /snap/go/10455/src/strings/strings.go:1186
		// _ = "end of CoverTab[1728]"
	}
//line /snap/go/10455/src/strings/strings.go:1187
	if _go_fuzz_dep_.CoverTab[786551] == 0 {
//line /snap/go/10455/src/strings/strings.go:1187
		_go_fuzz_dep_.CoverTab[525661]++
//line /snap/go/10455/src/strings/strings.go:1187
	} else {
//line /snap/go/10455/src/strings/strings.go:1187
		_go_fuzz_dep_.CoverTab[525662]++
//line /snap/go/10455/src/strings/strings.go:1187
	}
//line /snap/go/10455/src/strings/strings.go:1187
	// _ = "end of CoverTab[1703]"
//line /snap/go/10455/src/strings/strings.go:1187
	_go_fuzz_dep_.CoverTab[1704]++

//line /snap/go/10455/src/strings/strings.go:1190
	return len(t) == 0
//line /snap/go/10455/src/strings/strings.go:1190
	// _ = "end of CoverTab[1704]"
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
//line /snap/go/10455/src/strings/strings.go:1194
	_go_fuzz_dep_.CoverTab[1748]++
							n := len(substr)
							switch {
	case n == 0:
//line /snap/go/10455/src/strings/strings.go:1197
		_go_fuzz_dep_.CoverTab[525459]++
//line /snap/go/10455/src/strings/strings.go:1197
		_go_fuzz_dep_.CoverTab[1751]++
								return 0
//line /snap/go/10455/src/strings/strings.go:1198
		// _ = "end of CoverTab[1751]"
	case n == 1:
//line /snap/go/10455/src/strings/strings.go:1199
		_go_fuzz_dep_.CoverTab[525460]++
//line /snap/go/10455/src/strings/strings.go:1199
		_go_fuzz_dep_.CoverTab[1752]++
								return IndexByte(s, substr[0])
//line /snap/go/10455/src/strings/strings.go:1200
		// _ = "end of CoverTab[1752]"
	case n == len(s):
//line /snap/go/10455/src/strings/strings.go:1201
		_go_fuzz_dep_.CoverTab[525461]++
//line /snap/go/10455/src/strings/strings.go:1201
		_go_fuzz_dep_.CoverTab[1753]++
								if substr == s {
//line /snap/go/10455/src/strings/strings.go:1202
			_go_fuzz_dep_.CoverTab[525465]++
//line /snap/go/10455/src/strings/strings.go:1202
			_go_fuzz_dep_.CoverTab[1760]++
									return 0
//line /snap/go/10455/src/strings/strings.go:1203
			// _ = "end of CoverTab[1760]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1204
			_go_fuzz_dep_.CoverTab[525466]++
//line /snap/go/10455/src/strings/strings.go:1204
			_go_fuzz_dep_.CoverTab[1761]++
//line /snap/go/10455/src/strings/strings.go:1204
			// _ = "end of CoverTab[1761]"
//line /snap/go/10455/src/strings/strings.go:1204
		}
//line /snap/go/10455/src/strings/strings.go:1204
		// _ = "end of CoverTab[1753]"
//line /snap/go/10455/src/strings/strings.go:1204
		_go_fuzz_dep_.CoverTab[1754]++
								return -1
//line /snap/go/10455/src/strings/strings.go:1205
		// _ = "end of CoverTab[1754]"
	case n > len(s):
//line /snap/go/10455/src/strings/strings.go:1206
		_go_fuzz_dep_.CoverTab[525462]++
//line /snap/go/10455/src/strings/strings.go:1206
		_go_fuzz_dep_.CoverTab[1755]++
								return -1
//line /snap/go/10455/src/strings/strings.go:1207
		// _ = "end of CoverTab[1755]"
	case n <= bytealg.MaxLen:
//line /snap/go/10455/src/strings/strings.go:1208
		_go_fuzz_dep_.CoverTab[525463]++
//line /snap/go/10455/src/strings/strings.go:1208
		_go_fuzz_dep_.CoverTab[1756]++

								if len(s) <= bytealg.MaxBruteForce {
//line /snap/go/10455/src/strings/strings.go:1210
			_go_fuzz_dep_.CoverTab[525467]++
//line /snap/go/10455/src/strings/strings.go:1210
			_go_fuzz_dep_.CoverTab[1762]++
									return bytealg.IndexString(s, substr)
//line /snap/go/10455/src/strings/strings.go:1211
			// _ = "end of CoverTab[1762]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1212
			_go_fuzz_dep_.CoverTab[525468]++
//line /snap/go/10455/src/strings/strings.go:1212
			_go_fuzz_dep_.CoverTab[1763]++
//line /snap/go/10455/src/strings/strings.go:1212
			// _ = "end of CoverTab[1763]"
//line /snap/go/10455/src/strings/strings.go:1212
		}
//line /snap/go/10455/src/strings/strings.go:1212
		// _ = "end of CoverTab[1756]"
//line /snap/go/10455/src/strings/strings.go:1212
		_go_fuzz_dep_.CoverTab[1757]++
								c0 := substr[0]
								c1 := substr[1]
								i := 0
								t := len(s) - n + 1
								fails := 0
								for i < t {
//line /snap/go/10455/src/strings/strings.go:1218
			_go_fuzz_dep_.CoverTab[1764]++
									if s[i] != c0 {
//line /snap/go/10455/src/strings/strings.go:1219
				_go_fuzz_dep_.CoverTab[525469]++
//line /snap/go/10455/src/strings/strings.go:1219
				_go_fuzz_dep_.CoverTab[1767]++

//line /snap/go/10455/src/strings/strings.go:1222
				o := IndexByte(s[i+1:t], c0)
				if o < 0 {
//line /snap/go/10455/src/strings/strings.go:1223
					_go_fuzz_dep_.CoverTab[525471]++
//line /snap/go/10455/src/strings/strings.go:1223
					_go_fuzz_dep_.CoverTab[1769]++
											return -1
//line /snap/go/10455/src/strings/strings.go:1224
					// _ = "end of CoverTab[1769]"
				} else {
//line /snap/go/10455/src/strings/strings.go:1225
					_go_fuzz_dep_.CoverTab[525472]++
//line /snap/go/10455/src/strings/strings.go:1225
					_go_fuzz_dep_.CoverTab[1770]++
//line /snap/go/10455/src/strings/strings.go:1225
					// _ = "end of CoverTab[1770]"
//line /snap/go/10455/src/strings/strings.go:1225
				}
//line /snap/go/10455/src/strings/strings.go:1225
				// _ = "end of CoverTab[1767]"
//line /snap/go/10455/src/strings/strings.go:1225
				_go_fuzz_dep_.CoverTab[1768]++
										i += o + 1
//line /snap/go/10455/src/strings/strings.go:1226
				// _ = "end of CoverTab[1768]"
			} else {
//line /snap/go/10455/src/strings/strings.go:1227
				_go_fuzz_dep_.CoverTab[525470]++
//line /snap/go/10455/src/strings/strings.go:1227
				_go_fuzz_dep_.CoverTab[1771]++
//line /snap/go/10455/src/strings/strings.go:1227
				// _ = "end of CoverTab[1771]"
//line /snap/go/10455/src/strings/strings.go:1227
			}
//line /snap/go/10455/src/strings/strings.go:1227
			// _ = "end of CoverTab[1764]"
//line /snap/go/10455/src/strings/strings.go:1227
			_go_fuzz_dep_.CoverTab[1765]++
									if s[i+1] == c1 && func() bool {
//line /snap/go/10455/src/strings/strings.go:1228
				_go_fuzz_dep_.CoverTab[1772]++
//line /snap/go/10455/src/strings/strings.go:1228
				return s[i:i+n] == substr
//line /snap/go/10455/src/strings/strings.go:1228
				// _ = "end of CoverTab[1772]"
//line /snap/go/10455/src/strings/strings.go:1228
			}() {
//line /snap/go/10455/src/strings/strings.go:1228
				_go_fuzz_dep_.CoverTab[525473]++
//line /snap/go/10455/src/strings/strings.go:1228
				_go_fuzz_dep_.CoverTab[1773]++
										return i
//line /snap/go/10455/src/strings/strings.go:1229
				// _ = "end of CoverTab[1773]"
			} else {
//line /snap/go/10455/src/strings/strings.go:1230
				_go_fuzz_dep_.CoverTab[525474]++
//line /snap/go/10455/src/strings/strings.go:1230
				_go_fuzz_dep_.CoverTab[1774]++
//line /snap/go/10455/src/strings/strings.go:1230
				// _ = "end of CoverTab[1774]"
//line /snap/go/10455/src/strings/strings.go:1230
			}
//line /snap/go/10455/src/strings/strings.go:1230
			// _ = "end of CoverTab[1765]"
//line /snap/go/10455/src/strings/strings.go:1230
			_go_fuzz_dep_.CoverTab[1766]++
									fails++
									i++

									if fails > bytealg.Cutover(i) {
//line /snap/go/10455/src/strings/strings.go:1234
				_go_fuzz_dep_.CoverTab[525475]++
//line /snap/go/10455/src/strings/strings.go:1234
				_go_fuzz_dep_.CoverTab[1775]++
										r := bytealg.IndexString(s[i:], substr)
										if r >= 0 {
//line /snap/go/10455/src/strings/strings.go:1236
					_go_fuzz_dep_.CoverTab[525477]++
//line /snap/go/10455/src/strings/strings.go:1236
					_go_fuzz_dep_.CoverTab[1777]++
											return r + i
//line /snap/go/10455/src/strings/strings.go:1237
					// _ = "end of CoverTab[1777]"
				} else {
//line /snap/go/10455/src/strings/strings.go:1238
					_go_fuzz_dep_.CoverTab[525478]++
//line /snap/go/10455/src/strings/strings.go:1238
					_go_fuzz_dep_.CoverTab[1778]++
//line /snap/go/10455/src/strings/strings.go:1238
					// _ = "end of CoverTab[1778]"
//line /snap/go/10455/src/strings/strings.go:1238
				}
//line /snap/go/10455/src/strings/strings.go:1238
				// _ = "end of CoverTab[1775]"
//line /snap/go/10455/src/strings/strings.go:1238
				_go_fuzz_dep_.CoverTab[1776]++
										return -1
//line /snap/go/10455/src/strings/strings.go:1239
				// _ = "end of CoverTab[1776]"
			} else {
//line /snap/go/10455/src/strings/strings.go:1240
				_go_fuzz_dep_.CoverTab[525476]++
//line /snap/go/10455/src/strings/strings.go:1240
				_go_fuzz_dep_.CoverTab[1779]++
//line /snap/go/10455/src/strings/strings.go:1240
				// _ = "end of CoverTab[1779]"
//line /snap/go/10455/src/strings/strings.go:1240
			}
//line /snap/go/10455/src/strings/strings.go:1240
			// _ = "end of CoverTab[1766]"
		}
//line /snap/go/10455/src/strings/strings.go:1241
		// _ = "end of CoverTab[1757]"
//line /snap/go/10455/src/strings/strings.go:1241
		_go_fuzz_dep_.CoverTab[1758]++
								return -1
//line /snap/go/10455/src/strings/strings.go:1242
		// _ = "end of CoverTab[1758]"
//line /snap/go/10455/src/strings/strings.go:1242
	default:
//line /snap/go/10455/src/strings/strings.go:1242
		_go_fuzz_dep_.CoverTab[525464]++
//line /snap/go/10455/src/strings/strings.go:1242
		_go_fuzz_dep_.CoverTab[1759]++
//line /snap/go/10455/src/strings/strings.go:1242
		// _ = "end of CoverTab[1759]"
	}
//line /snap/go/10455/src/strings/strings.go:1243
	// _ = "end of CoverTab[1748]"
//line /snap/go/10455/src/strings/strings.go:1243
	_go_fuzz_dep_.CoverTab[1749]++
							c0 := substr[0]
							c1 := substr[1]
							i := 0
							t := len(s) - n + 1
							fails := 0
//line /snap/go/10455/src/strings/strings.go:1248
	_go_fuzz_dep_.CoverTab[786553] = 0
							for i < t {
//line /snap/go/10455/src/strings/strings.go:1249
		if _go_fuzz_dep_.CoverTab[786553] == 0 {
//line /snap/go/10455/src/strings/strings.go:1249
			_go_fuzz_dep_.CoverTab[525667]++
//line /snap/go/10455/src/strings/strings.go:1249
		} else {
//line /snap/go/10455/src/strings/strings.go:1249
			_go_fuzz_dep_.CoverTab[525668]++
//line /snap/go/10455/src/strings/strings.go:1249
		}
//line /snap/go/10455/src/strings/strings.go:1249
		_go_fuzz_dep_.CoverTab[786553] = 1
//line /snap/go/10455/src/strings/strings.go:1249
		_go_fuzz_dep_.CoverTab[1780]++
								if s[i] != c0 {
//line /snap/go/10455/src/strings/strings.go:1250
			_go_fuzz_dep_.CoverTab[525479]++
//line /snap/go/10455/src/strings/strings.go:1250
			_go_fuzz_dep_.CoverTab[1783]++
									o := IndexByte(s[i+1:t], c0)
									if o < 0 {
//line /snap/go/10455/src/strings/strings.go:1252
				_go_fuzz_dep_.CoverTab[525481]++
//line /snap/go/10455/src/strings/strings.go:1252
				_go_fuzz_dep_.CoverTab[1785]++
										return -1
//line /snap/go/10455/src/strings/strings.go:1253
				// _ = "end of CoverTab[1785]"
			} else {
//line /snap/go/10455/src/strings/strings.go:1254
				_go_fuzz_dep_.CoverTab[525482]++
//line /snap/go/10455/src/strings/strings.go:1254
				_go_fuzz_dep_.CoverTab[1786]++
//line /snap/go/10455/src/strings/strings.go:1254
				// _ = "end of CoverTab[1786]"
//line /snap/go/10455/src/strings/strings.go:1254
			}
//line /snap/go/10455/src/strings/strings.go:1254
			// _ = "end of CoverTab[1783]"
//line /snap/go/10455/src/strings/strings.go:1254
			_go_fuzz_dep_.CoverTab[1784]++
									i += o + 1
//line /snap/go/10455/src/strings/strings.go:1255
			// _ = "end of CoverTab[1784]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1256
			_go_fuzz_dep_.CoverTab[525480]++
//line /snap/go/10455/src/strings/strings.go:1256
			_go_fuzz_dep_.CoverTab[1787]++
//line /snap/go/10455/src/strings/strings.go:1256
			// _ = "end of CoverTab[1787]"
//line /snap/go/10455/src/strings/strings.go:1256
		}
//line /snap/go/10455/src/strings/strings.go:1256
		// _ = "end of CoverTab[1780]"
//line /snap/go/10455/src/strings/strings.go:1256
		_go_fuzz_dep_.CoverTab[1781]++
								if s[i+1] == c1 && func() bool {
//line /snap/go/10455/src/strings/strings.go:1257
			_go_fuzz_dep_.CoverTab[1788]++
//line /snap/go/10455/src/strings/strings.go:1257
			return s[i:i+n] == substr
//line /snap/go/10455/src/strings/strings.go:1257
			// _ = "end of CoverTab[1788]"
//line /snap/go/10455/src/strings/strings.go:1257
		}() {
//line /snap/go/10455/src/strings/strings.go:1257
			_go_fuzz_dep_.CoverTab[525483]++
//line /snap/go/10455/src/strings/strings.go:1257
			_go_fuzz_dep_.CoverTab[1789]++
									return i
//line /snap/go/10455/src/strings/strings.go:1258
			// _ = "end of CoverTab[1789]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1259
			_go_fuzz_dep_.CoverTab[525484]++
//line /snap/go/10455/src/strings/strings.go:1259
			_go_fuzz_dep_.CoverTab[1790]++
//line /snap/go/10455/src/strings/strings.go:1259
			// _ = "end of CoverTab[1790]"
//line /snap/go/10455/src/strings/strings.go:1259
		}
//line /snap/go/10455/src/strings/strings.go:1259
		// _ = "end of CoverTab[1781]"
//line /snap/go/10455/src/strings/strings.go:1259
		_go_fuzz_dep_.CoverTab[1782]++
								i++
								fails++
								if fails >= 4+i>>4 && func() bool {
//line /snap/go/10455/src/strings/strings.go:1262
			_go_fuzz_dep_.CoverTab[1791]++
//line /snap/go/10455/src/strings/strings.go:1262
			return i < t
//line /snap/go/10455/src/strings/strings.go:1262
			// _ = "end of CoverTab[1791]"
//line /snap/go/10455/src/strings/strings.go:1262
		}() {
//line /snap/go/10455/src/strings/strings.go:1262
			_go_fuzz_dep_.CoverTab[525485]++
//line /snap/go/10455/src/strings/strings.go:1262
			_go_fuzz_dep_.CoverTab[1792]++

									j := bytealg.IndexRabinKarp(s[i:], substr)
									if j < 0 {
//line /snap/go/10455/src/strings/strings.go:1265
				_go_fuzz_dep_.CoverTab[525487]++
//line /snap/go/10455/src/strings/strings.go:1265
				_go_fuzz_dep_.CoverTab[1794]++
										return -1
//line /snap/go/10455/src/strings/strings.go:1266
				// _ = "end of CoverTab[1794]"
			} else {
//line /snap/go/10455/src/strings/strings.go:1267
				_go_fuzz_dep_.CoverTab[525488]++
//line /snap/go/10455/src/strings/strings.go:1267
				_go_fuzz_dep_.CoverTab[1795]++
//line /snap/go/10455/src/strings/strings.go:1267
				// _ = "end of CoverTab[1795]"
//line /snap/go/10455/src/strings/strings.go:1267
			}
//line /snap/go/10455/src/strings/strings.go:1267
			// _ = "end of CoverTab[1792]"
//line /snap/go/10455/src/strings/strings.go:1267
			_go_fuzz_dep_.CoverTab[1793]++
									return i + j
//line /snap/go/10455/src/strings/strings.go:1268
			// _ = "end of CoverTab[1793]"
		} else {
//line /snap/go/10455/src/strings/strings.go:1269
			_go_fuzz_dep_.CoverTab[525486]++
//line /snap/go/10455/src/strings/strings.go:1269
			_go_fuzz_dep_.CoverTab[1796]++
//line /snap/go/10455/src/strings/strings.go:1269
			// _ = "end of CoverTab[1796]"
//line /snap/go/10455/src/strings/strings.go:1269
		}
//line /snap/go/10455/src/strings/strings.go:1269
		// _ = "end of CoverTab[1782]"
	}
//line /snap/go/10455/src/strings/strings.go:1270
	if _go_fuzz_dep_.CoverTab[786553] == 0 {
//line /snap/go/10455/src/strings/strings.go:1270
		_go_fuzz_dep_.CoverTab[525669]++
//line /snap/go/10455/src/strings/strings.go:1270
	} else {
//line /snap/go/10455/src/strings/strings.go:1270
		_go_fuzz_dep_.CoverTab[525670]++
//line /snap/go/10455/src/strings/strings.go:1270
	}
//line /snap/go/10455/src/strings/strings.go:1270
	// _ = "end of CoverTab[1749]"
//line /snap/go/10455/src/strings/strings.go:1270
	_go_fuzz_dep_.CoverTab[1750]++
							return -1
//line /snap/go/10455/src/strings/strings.go:1271
	// _ = "end of CoverTab[1750]"
}

// Cut slices s around the first instance of sep,
//line /snap/go/10455/src/strings/strings.go:1274
// returning the text before and after sep.
//line /snap/go/10455/src/strings/strings.go:1274
// The found result reports whether sep appears in s.
//line /snap/go/10455/src/strings/strings.go:1274
// If sep does not appear in s, cut returns s, "", false.
//line /snap/go/10455/src/strings/strings.go:1278
func Cut(s, sep string) (before, after string, found bool) {
//line /snap/go/10455/src/strings/strings.go:1278
	_go_fuzz_dep_.CoverTab[1797]++
							if i := Index(s, sep); i >= 0 {
//line /snap/go/10455/src/strings/strings.go:1279
		_go_fuzz_dep_.CoverTab[525489]++
//line /snap/go/10455/src/strings/strings.go:1279
		_go_fuzz_dep_.CoverTab[1799]++
								return s[:i], s[i+len(sep):], true
//line /snap/go/10455/src/strings/strings.go:1280
		// _ = "end of CoverTab[1799]"
	} else {
//line /snap/go/10455/src/strings/strings.go:1281
		_go_fuzz_dep_.CoverTab[525490]++
//line /snap/go/10455/src/strings/strings.go:1281
		_go_fuzz_dep_.CoverTab[1800]++
//line /snap/go/10455/src/strings/strings.go:1281
		// _ = "end of CoverTab[1800]"
//line /snap/go/10455/src/strings/strings.go:1281
	}
//line /snap/go/10455/src/strings/strings.go:1281
	// _ = "end of CoverTab[1797]"
//line /snap/go/10455/src/strings/strings.go:1281
	_go_fuzz_dep_.CoverTab[1798]++
							return s, "", false
//line /snap/go/10455/src/strings/strings.go:1282
	// _ = "end of CoverTab[1798]"
}

// CutPrefix returns s without the provided leading prefix string
//line /snap/go/10455/src/strings/strings.go:1285
// and reports whether it found the prefix.
//line /snap/go/10455/src/strings/strings.go:1285
// If s doesn't start with prefix, CutPrefix returns s, false.
//line /snap/go/10455/src/strings/strings.go:1285
// If prefix is the empty string, CutPrefix returns s, true.
//line /snap/go/10455/src/strings/strings.go:1289
func CutPrefix(s, prefix string) (after string, found bool) {
//line /snap/go/10455/src/strings/strings.go:1289
	_go_fuzz_dep_.CoverTab[1801]++
							if !HasPrefix(s, prefix) {
//line /snap/go/10455/src/strings/strings.go:1290
		_go_fuzz_dep_.CoverTab[525491]++
//line /snap/go/10455/src/strings/strings.go:1290
		_go_fuzz_dep_.CoverTab[1803]++
								return s, false
//line /snap/go/10455/src/strings/strings.go:1291
		// _ = "end of CoverTab[1803]"
	} else {
//line /snap/go/10455/src/strings/strings.go:1292
		_go_fuzz_dep_.CoverTab[525492]++
//line /snap/go/10455/src/strings/strings.go:1292
		_go_fuzz_dep_.CoverTab[1804]++
//line /snap/go/10455/src/strings/strings.go:1292
		// _ = "end of CoverTab[1804]"
//line /snap/go/10455/src/strings/strings.go:1292
	}
//line /snap/go/10455/src/strings/strings.go:1292
	// _ = "end of CoverTab[1801]"
//line /snap/go/10455/src/strings/strings.go:1292
	_go_fuzz_dep_.CoverTab[1802]++
							return s[len(prefix):], true
//line /snap/go/10455/src/strings/strings.go:1293
	// _ = "end of CoverTab[1802]"
}

// CutSuffix returns s without the provided ending suffix string
//line /snap/go/10455/src/strings/strings.go:1296
// and reports whether it found the suffix.
//line /snap/go/10455/src/strings/strings.go:1296
// If s doesn't end with suffix, CutSuffix returns s, false.
//line /snap/go/10455/src/strings/strings.go:1296
// If suffix is the empty string, CutSuffix returns s, true.
//line /snap/go/10455/src/strings/strings.go:1300
func CutSuffix(s, suffix string) (before string, found bool) {
//line /snap/go/10455/src/strings/strings.go:1300
	_go_fuzz_dep_.CoverTab[1805]++
							if !HasSuffix(s, suffix) {
//line /snap/go/10455/src/strings/strings.go:1301
		_go_fuzz_dep_.CoverTab[525493]++
//line /snap/go/10455/src/strings/strings.go:1301
		_go_fuzz_dep_.CoverTab[1807]++
								return s, false
//line /snap/go/10455/src/strings/strings.go:1302
		// _ = "end of CoverTab[1807]"
	} else {
//line /snap/go/10455/src/strings/strings.go:1303
		_go_fuzz_dep_.CoverTab[525494]++
//line /snap/go/10455/src/strings/strings.go:1303
		_go_fuzz_dep_.CoverTab[1808]++
//line /snap/go/10455/src/strings/strings.go:1303
		// _ = "end of CoverTab[1808]"
//line /snap/go/10455/src/strings/strings.go:1303
	}
//line /snap/go/10455/src/strings/strings.go:1303
	// _ = "end of CoverTab[1805]"
//line /snap/go/10455/src/strings/strings.go:1303
	_go_fuzz_dep_.CoverTab[1806]++
							return s[:len(s)-len(suffix)], true
//line /snap/go/10455/src/strings/strings.go:1304
	// _ = "end of CoverTab[1806]"
}

//line /snap/go/10455/src/strings/strings.go:1305
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/strings/strings.go:1305
var _ = _go_fuzz_dep_.CoverTab
