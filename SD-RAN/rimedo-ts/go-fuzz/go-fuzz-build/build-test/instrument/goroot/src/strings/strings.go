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
	_go_fuzz_dep_.CoverTab[1195]++
						l := utf8.RuneCountInString(s)
						if n < 0 || func() bool {
//line /usr/local/go/src/strings/strings.go:21
		_go_fuzz_dep_.CoverTab[1199]++
//line /usr/local/go/src/strings/strings.go:21
		return n > l
//line /usr/local/go/src/strings/strings.go:21
		// _ = "end of CoverTab[1199]"
//line /usr/local/go/src/strings/strings.go:21
	}() {
//line /usr/local/go/src/strings/strings.go:21
		_go_fuzz_dep_.CoverTab[1200]++
							n = l
//line /usr/local/go/src/strings/strings.go:22
		// _ = "end of CoverTab[1200]"
	} else {
//line /usr/local/go/src/strings/strings.go:23
		_go_fuzz_dep_.CoverTab[1201]++
//line /usr/local/go/src/strings/strings.go:23
		// _ = "end of CoverTab[1201]"
//line /usr/local/go/src/strings/strings.go:23
	}
//line /usr/local/go/src/strings/strings.go:23
	// _ = "end of CoverTab[1195]"
//line /usr/local/go/src/strings/strings.go:23
	_go_fuzz_dep_.CoverTab[1196]++
						a := make([]string, n)
						for i := 0; i < n-1; i++ {
//line /usr/local/go/src/strings/strings.go:25
		_go_fuzz_dep_.CoverTab[1202]++
							_, size := utf8.DecodeRuneInString(s)
							a[i] = s[:size]
							s = s[size:]
//line /usr/local/go/src/strings/strings.go:28
		// _ = "end of CoverTab[1202]"
	}
//line /usr/local/go/src/strings/strings.go:29
	// _ = "end of CoverTab[1196]"
//line /usr/local/go/src/strings/strings.go:29
	_go_fuzz_dep_.CoverTab[1197]++
						if n > 0 {
//line /usr/local/go/src/strings/strings.go:30
		_go_fuzz_dep_.CoverTab[1203]++
							a[n-1] = s
//line /usr/local/go/src/strings/strings.go:31
		// _ = "end of CoverTab[1203]"
	} else {
//line /usr/local/go/src/strings/strings.go:32
		_go_fuzz_dep_.CoverTab[1204]++
//line /usr/local/go/src/strings/strings.go:32
		// _ = "end of CoverTab[1204]"
//line /usr/local/go/src/strings/strings.go:32
	}
//line /usr/local/go/src/strings/strings.go:32
	// _ = "end of CoverTab[1197]"
//line /usr/local/go/src/strings/strings.go:32
	_go_fuzz_dep_.CoverTab[1198]++
						return a
//line /usr/local/go/src/strings/strings.go:33
	// _ = "end of CoverTab[1198]"
}

// Count counts the number of non-overlapping instances of substr in s.
//line /usr/local/go/src/strings/strings.go:36
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
//line /usr/local/go/src/strings/strings.go:38
func Count(s, substr string) int {
//line /usr/local/go/src/strings/strings.go:38
	_go_fuzz_dep_.CoverTab[1205]++

						if len(substr) == 0 {
//line /usr/local/go/src/strings/strings.go:40
		_go_fuzz_dep_.CoverTab[1208]++
							return utf8.RuneCountInString(s) + 1
//line /usr/local/go/src/strings/strings.go:41
		// _ = "end of CoverTab[1208]"
	} else {
//line /usr/local/go/src/strings/strings.go:42
		_go_fuzz_dep_.CoverTab[1209]++
//line /usr/local/go/src/strings/strings.go:42
		// _ = "end of CoverTab[1209]"
//line /usr/local/go/src/strings/strings.go:42
	}
//line /usr/local/go/src/strings/strings.go:42
	// _ = "end of CoverTab[1205]"
//line /usr/local/go/src/strings/strings.go:42
	_go_fuzz_dep_.CoverTab[1206]++
						if len(substr) == 1 {
//line /usr/local/go/src/strings/strings.go:43
		_go_fuzz_dep_.CoverTab[1210]++
							return bytealg.CountString(s, substr[0])
//line /usr/local/go/src/strings/strings.go:44
		// _ = "end of CoverTab[1210]"
	} else {
//line /usr/local/go/src/strings/strings.go:45
		_go_fuzz_dep_.CoverTab[1211]++
//line /usr/local/go/src/strings/strings.go:45
		// _ = "end of CoverTab[1211]"
//line /usr/local/go/src/strings/strings.go:45
	}
//line /usr/local/go/src/strings/strings.go:45
	// _ = "end of CoverTab[1206]"
//line /usr/local/go/src/strings/strings.go:45
	_go_fuzz_dep_.CoverTab[1207]++
						n := 0
						for {
//line /usr/local/go/src/strings/strings.go:47
		_go_fuzz_dep_.CoverTab[1212]++
							i := Index(s, substr)
							if i == -1 {
//line /usr/local/go/src/strings/strings.go:49
			_go_fuzz_dep_.CoverTab[1214]++
								return n
//line /usr/local/go/src/strings/strings.go:50
			// _ = "end of CoverTab[1214]"
		} else {
//line /usr/local/go/src/strings/strings.go:51
			_go_fuzz_dep_.CoverTab[1215]++
//line /usr/local/go/src/strings/strings.go:51
			// _ = "end of CoverTab[1215]"
//line /usr/local/go/src/strings/strings.go:51
		}
//line /usr/local/go/src/strings/strings.go:51
		// _ = "end of CoverTab[1212]"
//line /usr/local/go/src/strings/strings.go:51
		_go_fuzz_dep_.CoverTab[1213]++
							n++
							s = s[i+len(substr):]
//line /usr/local/go/src/strings/strings.go:53
		// _ = "end of CoverTab[1213]"
	}
//line /usr/local/go/src/strings/strings.go:54
	// _ = "end of CoverTab[1207]"
}

// Contains reports whether substr is within s.
func Contains(s, substr string) bool {
//line /usr/local/go/src/strings/strings.go:58
	_go_fuzz_dep_.CoverTab[1216]++
						return Index(s, substr) >= 0
//line /usr/local/go/src/strings/strings.go:59
	// _ = "end of CoverTab[1216]"
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool {
//line /usr/local/go/src/strings/strings.go:63
	_go_fuzz_dep_.CoverTab[1217]++
						return IndexAny(s, chars) >= 0
//line /usr/local/go/src/strings/strings.go:64
	// _ = "end of CoverTab[1217]"
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool {
//line /usr/local/go/src/strings/strings.go:68
	_go_fuzz_dep_.CoverTab[1218]++
						return IndexRune(s, r) >= 0
//line /usr/local/go/src/strings/strings.go:69
	// _ = "end of CoverTab[1218]"
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int {
//line /usr/local/go/src/strings/strings.go:73
	_go_fuzz_dep_.CoverTab[1219]++
						n := len(substr)
						switch {
	case n == 0:
//line /usr/local/go/src/strings/strings.go:76
		_go_fuzz_dep_.CoverTab[1224]++
							return len(s)
//line /usr/local/go/src/strings/strings.go:77
		// _ = "end of CoverTab[1224]"
	case n == 1:
//line /usr/local/go/src/strings/strings.go:78
		_go_fuzz_dep_.CoverTab[1225]++
							return LastIndexByte(s, substr[0])
//line /usr/local/go/src/strings/strings.go:79
		// _ = "end of CoverTab[1225]"
	case n == len(s):
//line /usr/local/go/src/strings/strings.go:80
		_go_fuzz_dep_.CoverTab[1226]++
							if substr == s {
//line /usr/local/go/src/strings/strings.go:81
			_go_fuzz_dep_.CoverTab[1230]++
								return 0
//line /usr/local/go/src/strings/strings.go:82
			// _ = "end of CoverTab[1230]"
		} else {
//line /usr/local/go/src/strings/strings.go:83
			_go_fuzz_dep_.CoverTab[1231]++
//line /usr/local/go/src/strings/strings.go:83
			// _ = "end of CoverTab[1231]"
//line /usr/local/go/src/strings/strings.go:83
		}
//line /usr/local/go/src/strings/strings.go:83
		// _ = "end of CoverTab[1226]"
//line /usr/local/go/src/strings/strings.go:83
		_go_fuzz_dep_.CoverTab[1227]++
							return -1
//line /usr/local/go/src/strings/strings.go:84
		// _ = "end of CoverTab[1227]"
	case n > len(s):
//line /usr/local/go/src/strings/strings.go:85
		_go_fuzz_dep_.CoverTab[1228]++
							return -1
//line /usr/local/go/src/strings/strings.go:86
		// _ = "end of CoverTab[1228]"
//line /usr/local/go/src/strings/strings.go:86
	default:
//line /usr/local/go/src/strings/strings.go:86
		_go_fuzz_dep_.CoverTab[1229]++
//line /usr/local/go/src/strings/strings.go:86
		// _ = "end of CoverTab[1229]"
	}
//line /usr/local/go/src/strings/strings.go:87
	// _ = "end of CoverTab[1219]"
//line /usr/local/go/src/strings/strings.go:87
	_go_fuzz_dep_.CoverTab[1220]++

						hashss, pow := bytealg.HashStrRev(substr)
						last := len(s) - n
						var h uint32
						for i := len(s) - 1; i >= last; i-- {
//line /usr/local/go/src/strings/strings.go:92
		_go_fuzz_dep_.CoverTab[1232]++
							h = h*bytealg.PrimeRK + uint32(s[i])
//line /usr/local/go/src/strings/strings.go:93
		// _ = "end of CoverTab[1232]"
	}
//line /usr/local/go/src/strings/strings.go:94
	// _ = "end of CoverTab[1220]"
//line /usr/local/go/src/strings/strings.go:94
	_go_fuzz_dep_.CoverTab[1221]++
						if h == hashss && func() bool {
//line /usr/local/go/src/strings/strings.go:95
		_go_fuzz_dep_.CoverTab[1233]++
//line /usr/local/go/src/strings/strings.go:95
		return s[last:] == substr
//line /usr/local/go/src/strings/strings.go:95
		// _ = "end of CoverTab[1233]"
//line /usr/local/go/src/strings/strings.go:95
	}() {
//line /usr/local/go/src/strings/strings.go:95
		_go_fuzz_dep_.CoverTab[1234]++
							return last
//line /usr/local/go/src/strings/strings.go:96
		// _ = "end of CoverTab[1234]"
	} else {
//line /usr/local/go/src/strings/strings.go:97
		_go_fuzz_dep_.CoverTab[1235]++
//line /usr/local/go/src/strings/strings.go:97
		// _ = "end of CoverTab[1235]"
//line /usr/local/go/src/strings/strings.go:97
	}
//line /usr/local/go/src/strings/strings.go:97
	// _ = "end of CoverTab[1221]"
//line /usr/local/go/src/strings/strings.go:97
	_go_fuzz_dep_.CoverTab[1222]++
						for i := last - 1; i >= 0; i-- {
//line /usr/local/go/src/strings/strings.go:98
		_go_fuzz_dep_.CoverTab[1236]++
							h *= bytealg.PrimeRK
							h += uint32(s[i])
							h -= pow * uint32(s[i+n])
							if h == hashss && func() bool {
//line /usr/local/go/src/strings/strings.go:102
			_go_fuzz_dep_.CoverTab[1237]++
//line /usr/local/go/src/strings/strings.go:102
			return s[i:i+n] == substr
//line /usr/local/go/src/strings/strings.go:102
			// _ = "end of CoverTab[1237]"
//line /usr/local/go/src/strings/strings.go:102
		}() {
//line /usr/local/go/src/strings/strings.go:102
			_go_fuzz_dep_.CoverTab[1238]++
									return i
//line /usr/local/go/src/strings/strings.go:103
			// _ = "end of CoverTab[1238]"
		} else {
//line /usr/local/go/src/strings/strings.go:104
			_go_fuzz_dep_.CoverTab[1239]++
//line /usr/local/go/src/strings/strings.go:104
			// _ = "end of CoverTab[1239]"
//line /usr/local/go/src/strings/strings.go:104
		}
//line /usr/local/go/src/strings/strings.go:104
		// _ = "end of CoverTab[1236]"
	}
//line /usr/local/go/src/strings/strings.go:105
	// _ = "end of CoverTab[1222]"
//line /usr/local/go/src/strings/strings.go:105
	_go_fuzz_dep_.CoverTab[1223]++
							return -1
//line /usr/local/go/src/strings/strings.go:106
	// _ = "end of CoverTab[1223]"
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s string, c byte) int {
//line /usr/local/go/src/strings/strings.go:110
	_go_fuzz_dep_.CoverTab[1240]++
							return bytealg.IndexByteString(s, c)
//line /usr/local/go/src/strings/strings.go:111
	// _ = "end of CoverTab[1240]"
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
	_go_fuzz_dep_.CoverTab[1241]++
							switch {
	case 0 <= r && func() bool {
//line /usr/local/go/src/strings/strings.go:120
		_go_fuzz_dep_.CoverTab[1247]++
//line /usr/local/go/src/strings/strings.go:120
		return r < utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:120
		// _ = "end of CoverTab[1247]"
//line /usr/local/go/src/strings/strings.go:120
	}():
//line /usr/local/go/src/strings/strings.go:120
		_go_fuzz_dep_.CoverTab[1242]++
								return IndexByte(s, byte(r))
//line /usr/local/go/src/strings/strings.go:121
		// _ = "end of CoverTab[1242]"
	case r == utf8.RuneError:
//line /usr/local/go/src/strings/strings.go:122
		_go_fuzz_dep_.CoverTab[1243]++
								for i, r := range s {
//line /usr/local/go/src/strings/strings.go:123
			_go_fuzz_dep_.CoverTab[1248]++
									if r == utf8.RuneError {
//line /usr/local/go/src/strings/strings.go:124
				_go_fuzz_dep_.CoverTab[1249]++
										return i
//line /usr/local/go/src/strings/strings.go:125
				// _ = "end of CoverTab[1249]"
			} else {
//line /usr/local/go/src/strings/strings.go:126
				_go_fuzz_dep_.CoverTab[1250]++
//line /usr/local/go/src/strings/strings.go:126
				// _ = "end of CoverTab[1250]"
//line /usr/local/go/src/strings/strings.go:126
			}
//line /usr/local/go/src/strings/strings.go:126
			// _ = "end of CoverTab[1248]"
		}
//line /usr/local/go/src/strings/strings.go:127
		// _ = "end of CoverTab[1243]"
//line /usr/local/go/src/strings/strings.go:127
		_go_fuzz_dep_.CoverTab[1244]++
								return -1
//line /usr/local/go/src/strings/strings.go:128
		// _ = "end of CoverTab[1244]"
	case !utf8.ValidRune(r):
//line /usr/local/go/src/strings/strings.go:129
		_go_fuzz_dep_.CoverTab[1245]++
								return -1
//line /usr/local/go/src/strings/strings.go:130
		// _ = "end of CoverTab[1245]"
	default:
//line /usr/local/go/src/strings/strings.go:131
		_go_fuzz_dep_.CoverTab[1246]++
								return Index(s, string(r))
//line /usr/local/go/src/strings/strings.go:132
		// _ = "end of CoverTab[1246]"
	}
//line /usr/local/go/src/strings/strings.go:133
	// _ = "end of CoverTab[1241]"
}

// IndexAny returns the index of the first instance of any Unicode code point
//line /usr/local/go/src/strings/strings.go:136
// from chars in s, or -1 if no Unicode code point from chars is present in s.
//line /usr/local/go/src/strings/strings.go:138
func IndexAny(s, chars string) int {
//line /usr/local/go/src/strings/strings.go:138
	_go_fuzz_dep_.CoverTab[1251]++
							if chars == "" {
//line /usr/local/go/src/strings/strings.go:139
		_go_fuzz_dep_.CoverTab[1256]++

								return -1
//line /usr/local/go/src/strings/strings.go:141
		// _ = "end of CoverTab[1256]"
	} else {
//line /usr/local/go/src/strings/strings.go:142
		_go_fuzz_dep_.CoverTab[1257]++
//line /usr/local/go/src/strings/strings.go:142
		// _ = "end of CoverTab[1257]"
//line /usr/local/go/src/strings/strings.go:142
	}
//line /usr/local/go/src/strings/strings.go:142
	// _ = "end of CoverTab[1251]"
//line /usr/local/go/src/strings/strings.go:142
	_go_fuzz_dep_.CoverTab[1252]++
							if len(chars) == 1 {
//line /usr/local/go/src/strings/strings.go:143
		_go_fuzz_dep_.CoverTab[1258]++

								r := rune(chars[0])
								if r >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:146
			_go_fuzz_dep_.CoverTab[1260]++
									r = utf8.RuneError
//line /usr/local/go/src/strings/strings.go:147
			// _ = "end of CoverTab[1260]"
		} else {
//line /usr/local/go/src/strings/strings.go:148
			_go_fuzz_dep_.CoverTab[1261]++
//line /usr/local/go/src/strings/strings.go:148
			// _ = "end of CoverTab[1261]"
//line /usr/local/go/src/strings/strings.go:148
		}
//line /usr/local/go/src/strings/strings.go:148
		// _ = "end of CoverTab[1258]"
//line /usr/local/go/src/strings/strings.go:148
		_go_fuzz_dep_.CoverTab[1259]++
								return IndexRune(s, r)
//line /usr/local/go/src/strings/strings.go:149
		// _ = "end of CoverTab[1259]"
	} else {
//line /usr/local/go/src/strings/strings.go:150
		_go_fuzz_dep_.CoverTab[1262]++
//line /usr/local/go/src/strings/strings.go:150
		// _ = "end of CoverTab[1262]"
//line /usr/local/go/src/strings/strings.go:150
	}
//line /usr/local/go/src/strings/strings.go:150
	// _ = "end of CoverTab[1252]"
//line /usr/local/go/src/strings/strings.go:150
	_go_fuzz_dep_.CoverTab[1253]++
							if len(s) > 8 {
//line /usr/local/go/src/strings/strings.go:151
		_go_fuzz_dep_.CoverTab[1263]++
								if as, isASCII := makeASCIISet(chars); isASCII {
//line /usr/local/go/src/strings/strings.go:152
			_go_fuzz_dep_.CoverTab[1264]++
									for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:153
				_go_fuzz_dep_.CoverTab[1266]++
										if as.contains(s[i]) {
//line /usr/local/go/src/strings/strings.go:154
					_go_fuzz_dep_.CoverTab[1267]++
											return i
//line /usr/local/go/src/strings/strings.go:155
					// _ = "end of CoverTab[1267]"
				} else {
//line /usr/local/go/src/strings/strings.go:156
					_go_fuzz_dep_.CoverTab[1268]++
//line /usr/local/go/src/strings/strings.go:156
					// _ = "end of CoverTab[1268]"
//line /usr/local/go/src/strings/strings.go:156
				}
//line /usr/local/go/src/strings/strings.go:156
				// _ = "end of CoverTab[1266]"
			}
//line /usr/local/go/src/strings/strings.go:157
			// _ = "end of CoverTab[1264]"
//line /usr/local/go/src/strings/strings.go:157
			_go_fuzz_dep_.CoverTab[1265]++
									return -1
//line /usr/local/go/src/strings/strings.go:158
			// _ = "end of CoverTab[1265]"
		} else {
//line /usr/local/go/src/strings/strings.go:159
			_go_fuzz_dep_.CoverTab[1269]++
//line /usr/local/go/src/strings/strings.go:159
			// _ = "end of CoverTab[1269]"
//line /usr/local/go/src/strings/strings.go:159
		}
//line /usr/local/go/src/strings/strings.go:159
		// _ = "end of CoverTab[1263]"
	} else {
//line /usr/local/go/src/strings/strings.go:160
		_go_fuzz_dep_.CoverTab[1270]++
//line /usr/local/go/src/strings/strings.go:160
		// _ = "end of CoverTab[1270]"
//line /usr/local/go/src/strings/strings.go:160
	}
//line /usr/local/go/src/strings/strings.go:160
	// _ = "end of CoverTab[1253]"
//line /usr/local/go/src/strings/strings.go:160
	_go_fuzz_dep_.CoverTab[1254]++
							for i, c := range s {
//line /usr/local/go/src/strings/strings.go:161
		_go_fuzz_dep_.CoverTab[1271]++
								if IndexRune(chars, c) >= 0 {
//line /usr/local/go/src/strings/strings.go:162
			_go_fuzz_dep_.CoverTab[1272]++
									return i
//line /usr/local/go/src/strings/strings.go:163
			// _ = "end of CoverTab[1272]"
		} else {
//line /usr/local/go/src/strings/strings.go:164
			_go_fuzz_dep_.CoverTab[1273]++
//line /usr/local/go/src/strings/strings.go:164
			// _ = "end of CoverTab[1273]"
//line /usr/local/go/src/strings/strings.go:164
		}
//line /usr/local/go/src/strings/strings.go:164
		// _ = "end of CoverTab[1271]"
	}
//line /usr/local/go/src/strings/strings.go:165
	// _ = "end of CoverTab[1254]"
//line /usr/local/go/src/strings/strings.go:165
	_go_fuzz_dep_.CoverTab[1255]++
							return -1
//line /usr/local/go/src/strings/strings.go:166
	// _ = "end of CoverTab[1255]"
}

// LastIndexAny returns the index of the last instance of any Unicode code
//line /usr/local/go/src/strings/strings.go:169
// point from chars in s, or -1 if no Unicode code point from chars is
//line /usr/local/go/src/strings/strings.go:169
// present in s.
//line /usr/local/go/src/strings/strings.go:172
func LastIndexAny(s, chars string) int {
//line /usr/local/go/src/strings/strings.go:172
	_go_fuzz_dep_.CoverTab[1274]++
							if chars == "" {
//line /usr/local/go/src/strings/strings.go:173
		_go_fuzz_dep_.CoverTab[1280]++

								return -1
//line /usr/local/go/src/strings/strings.go:175
		// _ = "end of CoverTab[1280]"
	} else {
//line /usr/local/go/src/strings/strings.go:176
		_go_fuzz_dep_.CoverTab[1281]++
//line /usr/local/go/src/strings/strings.go:176
		// _ = "end of CoverTab[1281]"
//line /usr/local/go/src/strings/strings.go:176
	}
//line /usr/local/go/src/strings/strings.go:176
	// _ = "end of CoverTab[1274]"
//line /usr/local/go/src/strings/strings.go:176
	_go_fuzz_dep_.CoverTab[1275]++
							if len(s) == 1 {
//line /usr/local/go/src/strings/strings.go:177
		_go_fuzz_dep_.CoverTab[1282]++
								rc := rune(s[0])
								if rc >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:179
			_go_fuzz_dep_.CoverTab[1285]++
									rc = utf8.RuneError
//line /usr/local/go/src/strings/strings.go:180
			// _ = "end of CoverTab[1285]"
		} else {
//line /usr/local/go/src/strings/strings.go:181
			_go_fuzz_dep_.CoverTab[1286]++
//line /usr/local/go/src/strings/strings.go:181
			// _ = "end of CoverTab[1286]"
//line /usr/local/go/src/strings/strings.go:181
		}
//line /usr/local/go/src/strings/strings.go:181
		// _ = "end of CoverTab[1282]"
//line /usr/local/go/src/strings/strings.go:181
		_go_fuzz_dep_.CoverTab[1283]++
								if IndexRune(chars, rc) >= 0 {
//line /usr/local/go/src/strings/strings.go:182
			_go_fuzz_dep_.CoverTab[1287]++
									return 0
//line /usr/local/go/src/strings/strings.go:183
			// _ = "end of CoverTab[1287]"
		} else {
//line /usr/local/go/src/strings/strings.go:184
			_go_fuzz_dep_.CoverTab[1288]++
//line /usr/local/go/src/strings/strings.go:184
			// _ = "end of CoverTab[1288]"
//line /usr/local/go/src/strings/strings.go:184
		}
//line /usr/local/go/src/strings/strings.go:184
		// _ = "end of CoverTab[1283]"
//line /usr/local/go/src/strings/strings.go:184
		_go_fuzz_dep_.CoverTab[1284]++
								return -1
//line /usr/local/go/src/strings/strings.go:185
		// _ = "end of CoverTab[1284]"
	} else {
//line /usr/local/go/src/strings/strings.go:186
		_go_fuzz_dep_.CoverTab[1289]++
//line /usr/local/go/src/strings/strings.go:186
		// _ = "end of CoverTab[1289]"
//line /usr/local/go/src/strings/strings.go:186
	}
//line /usr/local/go/src/strings/strings.go:186
	// _ = "end of CoverTab[1275]"
//line /usr/local/go/src/strings/strings.go:186
	_go_fuzz_dep_.CoverTab[1276]++
							if len(s) > 8 {
//line /usr/local/go/src/strings/strings.go:187
		_go_fuzz_dep_.CoverTab[1290]++
								if as, isASCII := makeASCIISet(chars); isASCII {
//line /usr/local/go/src/strings/strings.go:188
			_go_fuzz_dep_.CoverTab[1291]++
									for i := len(s) - 1; i >= 0; i-- {
//line /usr/local/go/src/strings/strings.go:189
				_go_fuzz_dep_.CoverTab[1293]++
										if as.contains(s[i]) {
//line /usr/local/go/src/strings/strings.go:190
					_go_fuzz_dep_.CoverTab[1294]++
											return i
//line /usr/local/go/src/strings/strings.go:191
					// _ = "end of CoverTab[1294]"
				} else {
//line /usr/local/go/src/strings/strings.go:192
					_go_fuzz_dep_.CoverTab[1295]++
//line /usr/local/go/src/strings/strings.go:192
					// _ = "end of CoverTab[1295]"
//line /usr/local/go/src/strings/strings.go:192
				}
//line /usr/local/go/src/strings/strings.go:192
				// _ = "end of CoverTab[1293]"
			}
//line /usr/local/go/src/strings/strings.go:193
			// _ = "end of CoverTab[1291]"
//line /usr/local/go/src/strings/strings.go:193
			_go_fuzz_dep_.CoverTab[1292]++
									return -1
//line /usr/local/go/src/strings/strings.go:194
			// _ = "end of CoverTab[1292]"
		} else {
//line /usr/local/go/src/strings/strings.go:195
			_go_fuzz_dep_.CoverTab[1296]++
//line /usr/local/go/src/strings/strings.go:195
			// _ = "end of CoverTab[1296]"
//line /usr/local/go/src/strings/strings.go:195
		}
//line /usr/local/go/src/strings/strings.go:195
		// _ = "end of CoverTab[1290]"
	} else {
//line /usr/local/go/src/strings/strings.go:196
		_go_fuzz_dep_.CoverTab[1297]++
//line /usr/local/go/src/strings/strings.go:196
		// _ = "end of CoverTab[1297]"
//line /usr/local/go/src/strings/strings.go:196
	}
//line /usr/local/go/src/strings/strings.go:196
	// _ = "end of CoverTab[1276]"
//line /usr/local/go/src/strings/strings.go:196
	_go_fuzz_dep_.CoverTab[1277]++
							if len(chars) == 1 {
//line /usr/local/go/src/strings/strings.go:197
		_go_fuzz_dep_.CoverTab[1298]++
								rc := rune(chars[0])
								if rc >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:199
			_go_fuzz_dep_.CoverTab[1301]++
									rc = utf8.RuneError
//line /usr/local/go/src/strings/strings.go:200
			// _ = "end of CoverTab[1301]"
		} else {
//line /usr/local/go/src/strings/strings.go:201
			_go_fuzz_dep_.CoverTab[1302]++
//line /usr/local/go/src/strings/strings.go:201
			// _ = "end of CoverTab[1302]"
//line /usr/local/go/src/strings/strings.go:201
		}
//line /usr/local/go/src/strings/strings.go:201
		// _ = "end of CoverTab[1298]"
//line /usr/local/go/src/strings/strings.go:201
		_go_fuzz_dep_.CoverTab[1299]++
								for i := len(s); i > 0; {
//line /usr/local/go/src/strings/strings.go:202
			_go_fuzz_dep_.CoverTab[1303]++
									r, size := utf8.DecodeLastRuneInString(s[:i])
									i -= size
									if rc == r {
//line /usr/local/go/src/strings/strings.go:205
				_go_fuzz_dep_.CoverTab[1304]++
										return i
//line /usr/local/go/src/strings/strings.go:206
				// _ = "end of CoverTab[1304]"
			} else {
//line /usr/local/go/src/strings/strings.go:207
				_go_fuzz_dep_.CoverTab[1305]++
//line /usr/local/go/src/strings/strings.go:207
				// _ = "end of CoverTab[1305]"
//line /usr/local/go/src/strings/strings.go:207
			}
//line /usr/local/go/src/strings/strings.go:207
			// _ = "end of CoverTab[1303]"
		}
//line /usr/local/go/src/strings/strings.go:208
		// _ = "end of CoverTab[1299]"
//line /usr/local/go/src/strings/strings.go:208
		_go_fuzz_dep_.CoverTab[1300]++
								return -1
//line /usr/local/go/src/strings/strings.go:209
		// _ = "end of CoverTab[1300]"
	} else {
//line /usr/local/go/src/strings/strings.go:210
		_go_fuzz_dep_.CoverTab[1306]++
//line /usr/local/go/src/strings/strings.go:210
		// _ = "end of CoverTab[1306]"
//line /usr/local/go/src/strings/strings.go:210
	}
//line /usr/local/go/src/strings/strings.go:210
	// _ = "end of CoverTab[1277]"
//line /usr/local/go/src/strings/strings.go:210
	_go_fuzz_dep_.CoverTab[1278]++
							for i := len(s); i > 0; {
//line /usr/local/go/src/strings/strings.go:211
		_go_fuzz_dep_.CoverTab[1307]++
								r, size := utf8.DecodeLastRuneInString(s[:i])
								i -= size
								if IndexRune(chars, r) >= 0 {
//line /usr/local/go/src/strings/strings.go:214
			_go_fuzz_dep_.CoverTab[1308]++
									return i
//line /usr/local/go/src/strings/strings.go:215
			// _ = "end of CoverTab[1308]"
		} else {
//line /usr/local/go/src/strings/strings.go:216
			_go_fuzz_dep_.CoverTab[1309]++
//line /usr/local/go/src/strings/strings.go:216
			// _ = "end of CoverTab[1309]"
//line /usr/local/go/src/strings/strings.go:216
		}
//line /usr/local/go/src/strings/strings.go:216
		// _ = "end of CoverTab[1307]"
	}
//line /usr/local/go/src/strings/strings.go:217
	// _ = "end of CoverTab[1278]"
//line /usr/local/go/src/strings/strings.go:217
	_go_fuzz_dep_.CoverTab[1279]++
							return -1
//line /usr/local/go/src/strings/strings.go:218
	// _ = "end of CoverTab[1279]"
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte(s string, c byte) int {
//line /usr/local/go/src/strings/strings.go:222
	_go_fuzz_dep_.CoverTab[1310]++
							for i := len(s) - 1; i >= 0; i-- {
//line /usr/local/go/src/strings/strings.go:223
		_go_fuzz_dep_.CoverTab[1312]++
								if s[i] == c {
//line /usr/local/go/src/strings/strings.go:224
			_go_fuzz_dep_.CoverTab[1313]++
									return i
//line /usr/local/go/src/strings/strings.go:225
			// _ = "end of CoverTab[1313]"
		} else {
//line /usr/local/go/src/strings/strings.go:226
			_go_fuzz_dep_.CoverTab[1314]++
//line /usr/local/go/src/strings/strings.go:226
			// _ = "end of CoverTab[1314]"
//line /usr/local/go/src/strings/strings.go:226
		}
//line /usr/local/go/src/strings/strings.go:226
		// _ = "end of CoverTab[1312]"
	}
//line /usr/local/go/src/strings/strings.go:227
	// _ = "end of CoverTab[1310]"
//line /usr/local/go/src/strings/strings.go:227
	_go_fuzz_dep_.CoverTab[1311]++
							return -1
//line /usr/local/go/src/strings/strings.go:228
	// _ = "end of CoverTab[1311]"
}

// Generic split: splits after each instance of sep,
//line /usr/local/go/src/strings/strings.go:231
// including sepSave bytes of sep in the subarrays.
//line /usr/local/go/src/strings/strings.go:233
func genSplit(s, sep string, sepSave, n int) []string {
//line /usr/local/go/src/strings/strings.go:233
	_go_fuzz_dep_.CoverTab[1315]++
							if n == 0 {
//line /usr/local/go/src/strings/strings.go:234
		_go_fuzz_dep_.CoverTab[1321]++
								return nil
//line /usr/local/go/src/strings/strings.go:235
		// _ = "end of CoverTab[1321]"
	} else {
//line /usr/local/go/src/strings/strings.go:236
		_go_fuzz_dep_.CoverTab[1322]++
//line /usr/local/go/src/strings/strings.go:236
		// _ = "end of CoverTab[1322]"
//line /usr/local/go/src/strings/strings.go:236
	}
//line /usr/local/go/src/strings/strings.go:236
	// _ = "end of CoverTab[1315]"
//line /usr/local/go/src/strings/strings.go:236
	_go_fuzz_dep_.CoverTab[1316]++
							if sep == "" {
//line /usr/local/go/src/strings/strings.go:237
		_go_fuzz_dep_.CoverTab[1323]++
								return explode(s, n)
//line /usr/local/go/src/strings/strings.go:238
		// _ = "end of CoverTab[1323]"
	} else {
//line /usr/local/go/src/strings/strings.go:239
		_go_fuzz_dep_.CoverTab[1324]++
//line /usr/local/go/src/strings/strings.go:239
		// _ = "end of CoverTab[1324]"
//line /usr/local/go/src/strings/strings.go:239
	}
//line /usr/local/go/src/strings/strings.go:239
	// _ = "end of CoverTab[1316]"
//line /usr/local/go/src/strings/strings.go:239
	_go_fuzz_dep_.CoverTab[1317]++
							if n < 0 {
//line /usr/local/go/src/strings/strings.go:240
		_go_fuzz_dep_.CoverTab[1325]++
								n = Count(s, sep) + 1
//line /usr/local/go/src/strings/strings.go:241
		// _ = "end of CoverTab[1325]"
	} else {
//line /usr/local/go/src/strings/strings.go:242
		_go_fuzz_dep_.CoverTab[1326]++
//line /usr/local/go/src/strings/strings.go:242
		// _ = "end of CoverTab[1326]"
//line /usr/local/go/src/strings/strings.go:242
	}
//line /usr/local/go/src/strings/strings.go:242
	// _ = "end of CoverTab[1317]"
//line /usr/local/go/src/strings/strings.go:242
	_go_fuzz_dep_.CoverTab[1318]++

							if n > len(s)+1 {
//line /usr/local/go/src/strings/strings.go:244
		_go_fuzz_dep_.CoverTab[1327]++
								n = len(s) + 1
//line /usr/local/go/src/strings/strings.go:245
		// _ = "end of CoverTab[1327]"
	} else {
//line /usr/local/go/src/strings/strings.go:246
		_go_fuzz_dep_.CoverTab[1328]++
//line /usr/local/go/src/strings/strings.go:246
		// _ = "end of CoverTab[1328]"
//line /usr/local/go/src/strings/strings.go:246
	}
//line /usr/local/go/src/strings/strings.go:246
	// _ = "end of CoverTab[1318]"
//line /usr/local/go/src/strings/strings.go:246
	_go_fuzz_dep_.CoverTab[1319]++
							a := make([]string, n)
							n--
							i := 0
							for i < n {
//line /usr/local/go/src/strings/strings.go:250
		_go_fuzz_dep_.CoverTab[1329]++
								m := Index(s, sep)
								if m < 0 {
//line /usr/local/go/src/strings/strings.go:252
			_go_fuzz_dep_.CoverTab[1331]++
									break
//line /usr/local/go/src/strings/strings.go:253
			// _ = "end of CoverTab[1331]"
		} else {
//line /usr/local/go/src/strings/strings.go:254
			_go_fuzz_dep_.CoverTab[1332]++
//line /usr/local/go/src/strings/strings.go:254
			// _ = "end of CoverTab[1332]"
//line /usr/local/go/src/strings/strings.go:254
		}
//line /usr/local/go/src/strings/strings.go:254
		// _ = "end of CoverTab[1329]"
//line /usr/local/go/src/strings/strings.go:254
		_go_fuzz_dep_.CoverTab[1330]++
								a[i] = s[:m+sepSave]
								s = s[m+len(sep):]
								i++
//line /usr/local/go/src/strings/strings.go:257
		// _ = "end of CoverTab[1330]"
	}
//line /usr/local/go/src/strings/strings.go:258
	// _ = "end of CoverTab[1319]"
//line /usr/local/go/src/strings/strings.go:258
	_go_fuzz_dep_.CoverTab[1320]++
							a[i] = s
							return a[:i+1]
//line /usr/local/go/src/strings/strings.go:260
	// _ = "end of CoverTab[1320]"
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
	_go_fuzz_dep_.CoverTab[1333]++
//line /usr/local/go/src/strings/strings.go:276
	return genSplit(s, sep, 0, n)
//line /usr/local/go/src/strings/strings.go:276
	// _ = "end of CoverTab[1333]"
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
	_go_fuzz_dep_.CoverTab[1334]++
							return genSplit(s, sep, len(sep), n)
//line /usr/local/go/src/strings/strings.go:290
	// _ = "end of CoverTab[1334]"
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
	_go_fuzz_dep_.CoverTab[1335]++
//line /usr/local/go/src/strings/strings.go:305
	return genSplit(s, sep, 0, -1)
//line /usr/local/go/src/strings/strings.go:305
	// _ = "end of CoverTab[1335]"
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
	_go_fuzz_dep_.CoverTab[1336]++
							return genSplit(s, sep, len(sep), -1)
//line /usr/local/go/src/strings/strings.go:318
	// _ = "end of CoverTab[1336]"
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
	_go_fuzz_dep_.CoverTab[1337]++

//line /usr/local/go/src/strings/strings.go:329
	n := 0
	wasSpace := 1

	setBits := uint8(0)
	for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:333
		_go_fuzz_dep_.CoverTab[1343]++
								r := s[i]
								setBits |= r
								isSpace := int(asciiSpace[r])
								n += wasSpace & ^isSpace
								wasSpace = isSpace
//line /usr/local/go/src/strings/strings.go:338
		// _ = "end of CoverTab[1343]"
	}
//line /usr/local/go/src/strings/strings.go:339
	// _ = "end of CoverTab[1337]"
//line /usr/local/go/src/strings/strings.go:339
	_go_fuzz_dep_.CoverTab[1338]++

							if setBits >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:341
		_go_fuzz_dep_.CoverTab[1344]++

								return FieldsFunc(s, unicode.IsSpace)
//line /usr/local/go/src/strings/strings.go:343
		// _ = "end of CoverTab[1344]"
	} else {
//line /usr/local/go/src/strings/strings.go:344
		_go_fuzz_dep_.CoverTab[1345]++
//line /usr/local/go/src/strings/strings.go:344
		// _ = "end of CoverTab[1345]"
//line /usr/local/go/src/strings/strings.go:344
	}
//line /usr/local/go/src/strings/strings.go:344
	// _ = "end of CoverTab[1338]"
//line /usr/local/go/src/strings/strings.go:344
	_go_fuzz_dep_.CoverTab[1339]++

							a := make([]string, n)
							na := 0
							fieldStart := 0
							i := 0

							for i < len(s) && func() bool {
//line /usr/local/go/src/strings/strings.go:351
		_go_fuzz_dep_.CoverTab[1346]++
//line /usr/local/go/src/strings/strings.go:351
		return asciiSpace[s[i]] != 0
//line /usr/local/go/src/strings/strings.go:351
		// _ = "end of CoverTab[1346]"
//line /usr/local/go/src/strings/strings.go:351
	}() {
//line /usr/local/go/src/strings/strings.go:351
		_go_fuzz_dep_.CoverTab[1347]++
								i++
//line /usr/local/go/src/strings/strings.go:352
		// _ = "end of CoverTab[1347]"
	}
//line /usr/local/go/src/strings/strings.go:353
	// _ = "end of CoverTab[1339]"
//line /usr/local/go/src/strings/strings.go:353
	_go_fuzz_dep_.CoverTab[1340]++
							fieldStart = i
							for i < len(s) {
//line /usr/local/go/src/strings/strings.go:355
		_go_fuzz_dep_.CoverTab[1348]++
								if asciiSpace[s[i]] == 0 {
//line /usr/local/go/src/strings/strings.go:356
			_go_fuzz_dep_.CoverTab[1351]++
									i++
									continue
//line /usr/local/go/src/strings/strings.go:358
			// _ = "end of CoverTab[1351]"
		} else {
//line /usr/local/go/src/strings/strings.go:359
			_go_fuzz_dep_.CoverTab[1352]++
//line /usr/local/go/src/strings/strings.go:359
			// _ = "end of CoverTab[1352]"
//line /usr/local/go/src/strings/strings.go:359
		}
//line /usr/local/go/src/strings/strings.go:359
		// _ = "end of CoverTab[1348]"
//line /usr/local/go/src/strings/strings.go:359
		_go_fuzz_dep_.CoverTab[1349]++
								a[na] = s[fieldStart:i]
								na++
								i++

								for i < len(s) && func() bool {
//line /usr/local/go/src/strings/strings.go:364
			_go_fuzz_dep_.CoverTab[1353]++
//line /usr/local/go/src/strings/strings.go:364
			return asciiSpace[s[i]] != 0
//line /usr/local/go/src/strings/strings.go:364
			// _ = "end of CoverTab[1353]"
//line /usr/local/go/src/strings/strings.go:364
		}() {
//line /usr/local/go/src/strings/strings.go:364
			_go_fuzz_dep_.CoverTab[1354]++
									i++
//line /usr/local/go/src/strings/strings.go:365
			// _ = "end of CoverTab[1354]"
		}
//line /usr/local/go/src/strings/strings.go:366
		// _ = "end of CoverTab[1349]"
//line /usr/local/go/src/strings/strings.go:366
		_go_fuzz_dep_.CoverTab[1350]++
								fieldStart = i
//line /usr/local/go/src/strings/strings.go:367
		// _ = "end of CoverTab[1350]"
	}
//line /usr/local/go/src/strings/strings.go:368
	// _ = "end of CoverTab[1340]"
//line /usr/local/go/src/strings/strings.go:368
	_go_fuzz_dep_.CoverTab[1341]++
							if fieldStart < len(s) {
//line /usr/local/go/src/strings/strings.go:369
		_go_fuzz_dep_.CoverTab[1355]++
								a[na] = s[fieldStart:]
//line /usr/local/go/src/strings/strings.go:370
		// _ = "end of CoverTab[1355]"
	} else {
//line /usr/local/go/src/strings/strings.go:371
		_go_fuzz_dep_.CoverTab[1356]++
//line /usr/local/go/src/strings/strings.go:371
		// _ = "end of CoverTab[1356]"
//line /usr/local/go/src/strings/strings.go:371
	}
//line /usr/local/go/src/strings/strings.go:371
	// _ = "end of CoverTab[1341]"
//line /usr/local/go/src/strings/strings.go:371
	_go_fuzz_dep_.CoverTab[1342]++
							return a
//line /usr/local/go/src/strings/strings.go:372
	// _ = "end of CoverTab[1342]"
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
	_go_fuzz_dep_.CoverTab[1357]++
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
		_go_fuzz_dep_.CoverTab[1361]++
								if f(rune) {
//line /usr/local/go/src/strings/strings.go:396
			_go_fuzz_dep_.CoverTab[1362]++
									if start >= 0 {
//line /usr/local/go/src/strings/strings.go:397
				_go_fuzz_dep_.CoverTab[1363]++
										spans = append(spans, span{start, end})

//line /usr/local/go/src/strings/strings.go:402
				start = ^start
//line /usr/local/go/src/strings/strings.go:402
				// _ = "end of CoverTab[1363]"
			} else {
//line /usr/local/go/src/strings/strings.go:403
				_go_fuzz_dep_.CoverTab[1364]++
//line /usr/local/go/src/strings/strings.go:403
				// _ = "end of CoverTab[1364]"
//line /usr/local/go/src/strings/strings.go:403
			}
//line /usr/local/go/src/strings/strings.go:403
			// _ = "end of CoverTab[1362]"
		} else {
//line /usr/local/go/src/strings/strings.go:404
			_go_fuzz_dep_.CoverTab[1365]++
									if start < 0 {
//line /usr/local/go/src/strings/strings.go:405
				_go_fuzz_dep_.CoverTab[1366]++
										start = end
//line /usr/local/go/src/strings/strings.go:406
				// _ = "end of CoverTab[1366]"
			} else {
//line /usr/local/go/src/strings/strings.go:407
				_go_fuzz_dep_.CoverTab[1367]++
//line /usr/local/go/src/strings/strings.go:407
				// _ = "end of CoverTab[1367]"
//line /usr/local/go/src/strings/strings.go:407
			}
//line /usr/local/go/src/strings/strings.go:407
			// _ = "end of CoverTab[1365]"
		}
//line /usr/local/go/src/strings/strings.go:408
		// _ = "end of CoverTab[1361]"
	}
//line /usr/local/go/src/strings/strings.go:409
	// _ = "end of CoverTab[1357]"
//line /usr/local/go/src/strings/strings.go:409
	_go_fuzz_dep_.CoverTab[1358]++

//line /usr/local/go/src/strings/strings.go:412
	if start >= 0 {
//line /usr/local/go/src/strings/strings.go:412
		_go_fuzz_dep_.CoverTab[1368]++
								spans = append(spans, span{start, len(s)})
//line /usr/local/go/src/strings/strings.go:413
		// _ = "end of CoverTab[1368]"
	} else {
//line /usr/local/go/src/strings/strings.go:414
		_go_fuzz_dep_.CoverTab[1369]++
//line /usr/local/go/src/strings/strings.go:414
		// _ = "end of CoverTab[1369]"
//line /usr/local/go/src/strings/strings.go:414
	}
//line /usr/local/go/src/strings/strings.go:414
	// _ = "end of CoverTab[1358]"
//line /usr/local/go/src/strings/strings.go:414
	_go_fuzz_dep_.CoverTab[1359]++

//line /usr/local/go/src/strings/strings.go:417
	a := make([]string, len(spans))
	for i, span := range spans {
//line /usr/local/go/src/strings/strings.go:418
		_go_fuzz_dep_.CoverTab[1370]++
								a[i] = s[span.start:span.end]
//line /usr/local/go/src/strings/strings.go:419
		// _ = "end of CoverTab[1370]"
	}
//line /usr/local/go/src/strings/strings.go:420
	// _ = "end of CoverTab[1359]"
//line /usr/local/go/src/strings/strings.go:420
	_go_fuzz_dep_.CoverTab[1360]++

							return a
//line /usr/local/go/src/strings/strings.go:422
	// _ = "end of CoverTab[1360]"
}

// Join concatenates the elements of its first argument to create a single string. The separator
//line /usr/local/go/src/strings/strings.go:425
// string sep is placed between elements in the resulting string.
//line /usr/local/go/src/strings/strings.go:427
func Join(elems []string, sep string) string {
//line /usr/local/go/src/strings/strings.go:427
	_go_fuzz_dep_.CoverTab[1371]++
							switch len(elems) {
	case 0:
//line /usr/local/go/src/strings/strings.go:429
		_go_fuzz_dep_.CoverTab[1375]++
								return ""
//line /usr/local/go/src/strings/strings.go:430
		// _ = "end of CoverTab[1375]"
	case 1:
//line /usr/local/go/src/strings/strings.go:431
		_go_fuzz_dep_.CoverTab[1376]++
								return elems[0]
//line /usr/local/go/src/strings/strings.go:432
		// _ = "end of CoverTab[1376]"
//line /usr/local/go/src/strings/strings.go:432
	default:
//line /usr/local/go/src/strings/strings.go:432
		_go_fuzz_dep_.CoverTab[1377]++
//line /usr/local/go/src/strings/strings.go:432
		// _ = "end of CoverTab[1377]"
	}
//line /usr/local/go/src/strings/strings.go:433
	// _ = "end of CoverTab[1371]"
//line /usr/local/go/src/strings/strings.go:433
	_go_fuzz_dep_.CoverTab[1372]++
							n := len(sep) * (len(elems) - 1)
							for i := 0; i < len(elems); i++ {
//line /usr/local/go/src/strings/strings.go:435
		_go_fuzz_dep_.CoverTab[1378]++
								n += len(elems[i])
//line /usr/local/go/src/strings/strings.go:436
		// _ = "end of CoverTab[1378]"
	}
//line /usr/local/go/src/strings/strings.go:437
	// _ = "end of CoverTab[1372]"
//line /usr/local/go/src/strings/strings.go:437
	_go_fuzz_dep_.CoverTab[1373]++

							var b Builder
							b.Grow(n)
							b.WriteString(elems[0])
							for _, s := range elems[1:] {
//line /usr/local/go/src/strings/strings.go:442
		_go_fuzz_dep_.CoverTab[1379]++
								b.WriteString(sep)
								b.WriteString(s)
//line /usr/local/go/src/strings/strings.go:444
		// _ = "end of CoverTab[1379]"
	}
//line /usr/local/go/src/strings/strings.go:445
	// _ = "end of CoverTab[1373]"
//line /usr/local/go/src/strings/strings.go:445
	_go_fuzz_dep_.CoverTab[1374]++
							return b.String()
//line /usr/local/go/src/strings/strings.go:446
	// _ = "end of CoverTab[1374]"
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool {
//line /usr/local/go/src/strings/strings.go:450
	_go_fuzz_dep_.CoverTab[1380]++
							return len(s) >= len(prefix) && func() bool {
//line /usr/local/go/src/strings/strings.go:451
		_go_fuzz_dep_.CoverTab[1381]++
//line /usr/local/go/src/strings/strings.go:451
		return s[0:len(prefix)] == prefix
//line /usr/local/go/src/strings/strings.go:451
		// _ = "end of CoverTab[1381]"
//line /usr/local/go/src/strings/strings.go:451
	}()
//line /usr/local/go/src/strings/strings.go:451
	// _ = "end of CoverTab[1380]"
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
//line /usr/local/go/src/strings/strings.go:455
	_go_fuzz_dep_.CoverTab[1382]++
							return len(s) >= len(suffix) && func() bool {
//line /usr/local/go/src/strings/strings.go:456
		_go_fuzz_dep_.CoverTab[1383]++
//line /usr/local/go/src/strings/strings.go:456
		return s[len(s)-len(suffix):] == suffix
//line /usr/local/go/src/strings/strings.go:456
		// _ = "end of CoverTab[1383]"
//line /usr/local/go/src/strings/strings.go:456
	}()
//line /usr/local/go/src/strings/strings.go:456
	// _ = "end of CoverTab[1382]"
}

// Map returns a copy of the string s with all its characters modified
//line /usr/local/go/src/strings/strings.go:459
// according to the mapping function. If mapping returns a negative value, the character is
//line /usr/local/go/src/strings/strings.go:459
// dropped from the string with no replacement.
//line /usr/local/go/src/strings/strings.go:462
func Map(mapping func(rune) rune, s string) string {
//line /usr/local/go/src/strings/strings.go:462
	_go_fuzz_dep_.CoverTab[1384]++

//line /usr/local/go/src/strings/strings.go:467
	// The output buffer b is initialized on demand, the first
	// time a character differs.
	var b Builder

	for i, c := range s {
//line /usr/local/go/src/strings/strings.go:471
		_go_fuzz_dep_.CoverTab[1388]++
								r := mapping(c)
								if r == c && func() bool {
//line /usr/local/go/src/strings/strings.go:473
			_go_fuzz_dep_.CoverTab[1392]++
//line /usr/local/go/src/strings/strings.go:473
			return c != utf8.RuneError
//line /usr/local/go/src/strings/strings.go:473
			// _ = "end of CoverTab[1392]"
//line /usr/local/go/src/strings/strings.go:473
		}() {
//line /usr/local/go/src/strings/strings.go:473
			_go_fuzz_dep_.CoverTab[1393]++
									continue
//line /usr/local/go/src/strings/strings.go:474
			// _ = "end of CoverTab[1393]"
		} else {
//line /usr/local/go/src/strings/strings.go:475
			_go_fuzz_dep_.CoverTab[1394]++
//line /usr/local/go/src/strings/strings.go:475
			// _ = "end of CoverTab[1394]"
//line /usr/local/go/src/strings/strings.go:475
		}
//line /usr/local/go/src/strings/strings.go:475
		// _ = "end of CoverTab[1388]"
//line /usr/local/go/src/strings/strings.go:475
		_go_fuzz_dep_.CoverTab[1389]++

								var width int
								if c == utf8.RuneError {
//line /usr/local/go/src/strings/strings.go:478
			_go_fuzz_dep_.CoverTab[1395]++
									c, width = utf8.DecodeRuneInString(s[i:])
									if width != 1 && func() bool {
//line /usr/local/go/src/strings/strings.go:480
				_go_fuzz_dep_.CoverTab[1396]++
//line /usr/local/go/src/strings/strings.go:480
				return r == c
//line /usr/local/go/src/strings/strings.go:480
				// _ = "end of CoverTab[1396]"
//line /usr/local/go/src/strings/strings.go:480
			}() {
//line /usr/local/go/src/strings/strings.go:480
				_go_fuzz_dep_.CoverTab[1397]++
										continue
//line /usr/local/go/src/strings/strings.go:481
				// _ = "end of CoverTab[1397]"
			} else {
//line /usr/local/go/src/strings/strings.go:482
				_go_fuzz_dep_.CoverTab[1398]++
//line /usr/local/go/src/strings/strings.go:482
				// _ = "end of CoverTab[1398]"
//line /usr/local/go/src/strings/strings.go:482
			}
//line /usr/local/go/src/strings/strings.go:482
			// _ = "end of CoverTab[1395]"
		} else {
//line /usr/local/go/src/strings/strings.go:483
			_go_fuzz_dep_.CoverTab[1399]++
									width = utf8.RuneLen(c)
//line /usr/local/go/src/strings/strings.go:484
			// _ = "end of CoverTab[1399]"
		}
//line /usr/local/go/src/strings/strings.go:485
		// _ = "end of CoverTab[1389]"
//line /usr/local/go/src/strings/strings.go:485
		_go_fuzz_dep_.CoverTab[1390]++

								b.Grow(len(s) + utf8.UTFMax)
								b.WriteString(s[:i])
								if r >= 0 {
//line /usr/local/go/src/strings/strings.go:489
			_go_fuzz_dep_.CoverTab[1400]++
									b.WriteRune(r)
//line /usr/local/go/src/strings/strings.go:490
			// _ = "end of CoverTab[1400]"
		} else {
//line /usr/local/go/src/strings/strings.go:491
			_go_fuzz_dep_.CoverTab[1401]++
//line /usr/local/go/src/strings/strings.go:491
			// _ = "end of CoverTab[1401]"
//line /usr/local/go/src/strings/strings.go:491
		}
//line /usr/local/go/src/strings/strings.go:491
		// _ = "end of CoverTab[1390]"
//line /usr/local/go/src/strings/strings.go:491
		_go_fuzz_dep_.CoverTab[1391]++

								s = s[i+width:]
								break
//line /usr/local/go/src/strings/strings.go:494
		// _ = "end of CoverTab[1391]"
	}
//line /usr/local/go/src/strings/strings.go:495
	// _ = "end of CoverTab[1384]"
//line /usr/local/go/src/strings/strings.go:495
	_go_fuzz_dep_.CoverTab[1385]++

//line /usr/local/go/src/strings/strings.go:498
	if b.Cap() == 0 {
//line /usr/local/go/src/strings/strings.go:498
		_go_fuzz_dep_.CoverTab[1402]++
								return s
//line /usr/local/go/src/strings/strings.go:499
		// _ = "end of CoverTab[1402]"
	} else {
//line /usr/local/go/src/strings/strings.go:500
		_go_fuzz_dep_.CoverTab[1403]++
//line /usr/local/go/src/strings/strings.go:500
		// _ = "end of CoverTab[1403]"
//line /usr/local/go/src/strings/strings.go:500
	}
//line /usr/local/go/src/strings/strings.go:500
	// _ = "end of CoverTab[1385]"
//line /usr/local/go/src/strings/strings.go:500
	_go_fuzz_dep_.CoverTab[1386]++

							for _, c := range s {
//line /usr/local/go/src/strings/strings.go:502
		_go_fuzz_dep_.CoverTab[1404]++
								r := mapping(c)

								if r >= 0 {
//line /usr/local/go/src/strings/strings.go:505
			_go_fuzz_dep_.CoverTab[1405]++

//line /usr/local/go/src/strings/strings.go:509
			if r < utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:509
				_go_fuzz_dep_.CoverTab[1406]++
										b.WriteByte(byte(r))
//line /usr/local/go/src/strings/strings.go:510
				// _ = "end of CoverTab[1406]"
			} else {
//line /usr/local/go/src/strings/strings.go:511
				_go_fuzz_dep_.CoverTab[1407]++

										b.WriteRune(r)
//line /usr/local/go/src/strings/strings.go:513
				// _ = "end of CoverTab[1407]"
			}
//line /usr/local/go/src/strings/strings.go:514
			// _ = "end of CoverTab[1405]"
		} else {
//line /usr/local/go/src/strings/strings.go:515
			_go_fuzz_dep_.CoverTab[1408]++
//line /usr/local/go/src/strings/strings.go:515
			// _ = "end of CoverTab[1408]"
//line /usr/local/go/src/strings/strings.go:515
		}
//line /usr/local/go/src/strings/strings.go:515
		// _ = "end of CoverTab[1404]"
	}
//line /usr/local/go/src/strings/strings.go:516
	// _ = "end of CoverTab[1386]"
//line /usr/local/go/src/strings/strings.go:516
	_go_fuzz_dep_.CoverTab[1387]++

							return b.String()
//line /usr/local/go/src/strings/strings.go:518
	// _ = "end of CoverTab[1387]"
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
	_go_fuzz_dep_.CoverTab[1409]++
							switch count {
	case 0:
//line /usr/local/go/src/strings/strings.go:527
		_go_fuzz_dep_.CoverTab[1415]++
								return ""
//line /usr/local/go/src/strings/strings.go:528
		// _ = "end of CoverTab[1415]"
	case 1:
//line /usr/local/go/src/strings/strings.go:529
		_go_fuzz_dep_.CoverTab[1416]++
								return s
//line /usr/local/go/src/strings/strings.go:530
		// _ = "end of CoverTab[1416]"
//line /usr/local/go/src/strings/strings.go:530
	default:
//line /usr/local/go/src/strings/strings.go:530
		_go_fuzz_dep_.CoverTab[1417]++
//line /usr/local/go/src/strings/strings.go:530
		// _ = "end of CoverTab[1417]"
	}
//line /usr/local/go/src/strings/strings.go:531
	// _ = "end of CoverTab[1409]"
//line /usr/local/go/src/strings/strings.go:531
	_go_fuzz_dep_.CoverTab[1410]++

//line /usr/local/go/src/strings/strings.go:537
	if count < 0 {
//line /usr/local/go/src/strings/strings.go:537
		_go_fuzz_dep_.CoverTab[1418]++
								panic("strings: negative Repeat count")
//line /usr/local/go/src/strings/strings.go:538
		// _ = "end of CoverTab[1418]"
	} else {
//line /usr/local/go/src/strings/strings.go:539
		_go_fuzz_dep_.CoverTab[1419]++
//line /usr/local/go/src/strings/strings.go:539
		if len(s)*count/count != len(s) {
//line /usr/local/go/src/strings/strings.go:539
			_go_fuzz_dep_.CoverTab[1420]++
									panic("strings: Repeat count causes overflow")
//line /usr/local/go/src/strings/strings.go:540
			// _ = "end of CoverTab[1420]"
		} else {
//line /usr/local/go/src/strings/strings.go:541
			_go_fuzz_dep_.CoverTab[1421]++
//line /usr/local/go/src/strings/strings.go:541
			// _ = "end of CoverTab[1421]"
//line /usr/local/go/src/strings/strings.go:541
		}
//line /usr/local/go/src/strings/strings.go:541
		// _ = "end of CoverTab[1419]"
//line /usr/local/go/src/strings/strings.go:541
	}
//line /usr/local/go/src/strings/strings.go:541
	// _ = "end of CoverTab[1410]"
//line /usr/local/go/src/strings/strings.go:541
	_go_fuzz_dep_.CoverTab[1411]++

							if len(s) == 0 {
//line /usr/local/go/src/strings/strings.go:543
		_go_fuzz_dep_.CoverTab[1422]++
								return ""
//line /usr/local/go/src/strings/strings.go:544
		// _ = "end of CoverTab[1422]"
	} else {
//line /usr/local/go/src/strings/strings.go:545
		_go_fuzz_dep_.CoverTab[1423]++
//line /usr/local/go/src/strings/strings.go:545
		// _ = "end of CoverTab[1423]"
//line /usr/local/go/src/strings/strings.go:545
	}
//line /usr/local/go/src/strings/strings.go:545
	// _ = "end of CoverTab[1411]"
//line /usr/local/go/src/strings/strings.go:545
	_go_fuzz_dep_.CoverTab[1412]++

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
		_go_fuzz_dep_.CoverTab[1424]++
								chunkMax = chunkLimit / len(s) * len(s)
								if chunkMax == 0 {
//line /usr/local/go/src/strings/strings.go:563
			_go_fuzz_dep_.CoverTab[1425]++
									chunkMax = len(s)
//line /usr/local/go/src/strings/strings.go:564
			// _ = "end of CoverTab[1425]"
		} else {
//line /usr/local/go/src/strings/strings.go:565
			_go_fuzz_dep_.CoverTab[1426]++
//line /usr/local/go/src/strings/strings.go:565
			// _ = "end of CoverTab[1426]"
//line /usr/local/go/src/strings/strings.go:565
		}
//line /usr/local/go/src/strings/strings.go:565
		// _ = "end of CoverTab[1424]"
	} else {
//line /usr/local/go/src/strings/strings.go:566
		_go_fuzz_dep_.CoverTab[1427]++
//line /usr/local/go/src/strings/strings.go:566
		// _ = "end of CoverTab[1427]"
//line /usr/local/go/src/strings/strings.go:566
	}
//line /usr/local/go/src/strings/strings.go:566
	// _ = "end of CoverTab[1412]"
//line /usr/local/go/src/strings/strings.go:566
	_go_fuzz_dep_.CoverTab[1413]++

							var b Builder
							b.Grow(n)
							b.WriteString(s)
							for b.Len() < n {
//line /usr/local/go/src/strings/strings.go:571
		_go_fuzz_dep_.CoverTab[1428]++
								chunk := n - b.Len()
								if chunk > b.Len() {
//line /usr/local/go/src/strings/strings.go:573
			_go_fuzz_dep_.CoverTab[1431]++
									chunk = b.Len()
//line /usr/local/go/src/strings/strings.go:574
			// _ = "end of CoverTab[1431]"
		} else {
//line /usr/local/go/src/strings/strings.go:575
			_go_fuzz_dep_.CoverTab[1432]++
//line /usr/local/go/src/strings/strings.go:575
			// _ = "end of CoverTab[1432]"
//line /usr/local/go/src/strings/strings.go:575
		}
//line /usr/local/go/src/strings/strings.go:575
		// _ = "end of CoverTab[1428]"
//line /usr/local/go/src/strings/strings.go:575
		_go_fuzz_dep_.CoverTab[1429]++
								if chunk > chunkMax {
//line /usr/local/go/src/strings/strings.go:576
			_go_fuzz_dep_.CoverTab[1433]++
									chunk = chunkMax
//line /usr/local/go/src/strings/strings.go:577
			// _ = "end of CoverTab[1433]"
		} else {
//line /usr/local/go/src/strings/strings.go:578
			_go_fuzz_dep_.CoverTab[1434]++
//line /usr/local/go/src/strings/strings.go:578
			// _ = "end of CoverTab[1434]"
//line /usr/local/go/src/strings/strings.go:578
		}
//line /usr/local/go/src/strings/strings.go:578
		// _ = "end of CoverTab[1429]"
//line /usr/local/go/src/strings/strings.go:578
		_go_fuzz_dep_.CoverTab[1430]++
								b.WriteString(b.String()[:chunk])
//line /usr/local/go/src/strings/strings.go:579
		// _ = "end of CoverTab[1430]"
	}
//line /usr/local/go/src/strings/strings.go:580
	// _ = "end of CoverTab[1413]"
//line /usr/local/go/src/strings/strings.go:580
	_go_fuzz_dep_.CoverTab[1414]++
							return b.String()
//line /usr/local/go/src/strings/strings.go:581
	// _ = "end of CoverTab[1414]"
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
//line /usr/local/go/src/strings/strings.go:585
	_go_fuzz_dep_.CoverTab[1435]++
							isASCII, hasLower := true, false
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:587
		_go_fuzz_dep_.CoverTab[1438]++
								c := s[i]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:589
			_go_fuzz_dep_.CoverTab[1440]++
									isASCII = false
									break
//line /usr/local/go/src/strings/strings.go:591
			// _ = "end of CoverTab[1440]"
		} else {
//line /usr/local/go/src/strings/strings.go:592
			_go_fuzz_dep_.CoverTab[1441]++
//line /usr/local/go/src/strings/strings.go:592
			// _ = "end of CoverTab[1441]"
//line /usr/local/go/src/strings/strings.go:592
		}
//line /usr/local/go/src/strings/strings.go:592
		// _ = "end of CoverTab[1438]"
//line /usr/local/go/src/strings/strings.go:592
		_go_fuzz_dep_.CoverTab[1439]++
								hasLower = hasLower || func() bool {
//line /usr/local/go/src/strings/strings.go:593
			_go_fuzz_dep_.CoverTab[1442]++
//line /usr/local/go/src/strings/strings.go:593
			return ('a' <= c && func() bool {
//line /usr/local/go/src/strings/strings.go:593
				_go_fuzz_dep_.CoverTab[1443]++
//line /usr/local/go/src/strings/strings.go:593
				return c <= 'z'
//line /usr/local/go/src/strings/strings.go:593
				// _ = "end of CoverTab[1443]"
//line /usr/local/go/src/strings/strings.go:593
			}())
//line /usr/local/go/src/strings/strings.go:593
			// _ = "end of CoverTab[1442]"
//line /usr/local/go/src/strings/strings.go:593
		}()
//line /usr/local/go/src/strings/strings.go:593
		// _ = "end of CoverTab[1439]"
	}
//line /usr/local/go/src/strings/strings.go:594
	// _ = "end of CoverTab[1435]"
//line /usr/local/go/src/strings/strings.go:594
	_go_fuzz_dep_.CoverTab[1436]++

							if isASCII {
//line /usr/local/go/src/strings/strings.go:596
		_go_fuzz_dep_.CoverTab[1444]++
								if !hasLower {
//line /usr/local/go/src/strings/strings.go:597
			_go_fuzz_dep_.CoverTab[1448]++
									return s
//line /usr/local/go/src/strings/strings.go:598
			// _ = "end of CoverTab[1448]"
		} else {
//line /usr/local/go/src/strings/strings.go:599
			_go_fuzz_dep_.CoverTab[1449]++
//line /usr/local/go/src/strings/strings.go:599
			// _ = "end of CoverTab[1449]"
//line /usr/local/go/src/strings/strings.go:599
		}
//line /usr/local/go/src/strings/strings.go:599
		// _ = "end of CoverTab[1444]"
//line /usr/local/go/src/strings/strings.go:599
		_go_fuzz_dep_.CoverTab[1445]++
								var (
			b	Builder
			pos	int
		)
		b.Grow(len(s))
		for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:605
			_go_fuzz_dep_.CoverTab[1450]++
									c := s[i]
									if 'a' <= c && func() bool {
//line /usr/local/go/src/strings/strings.go:607
				_go_fuzz_dep_.CoverTab[1451]++
//line /usr/local/go/src/strings/strings.go:607
				return c <= 'z'
//line /usr/local/go/src/strings/strings.go:607
				// _ = "end of CoverTab[1451]"
//line /usr/local/go/src/strings/strings.go:607
			}() {
//line /usr/local/go/src/strings/strings.go:607
				_go_fuzz_dep_.CoverTab[1452]++
										c -= 'a' - 'A'
										if pos < i {
//line /usr/local/go/src/strings/strings.go:609
					_go_fuzz_dep_.CoverTab[1454]++
											b.WriteString(s[pos:i])
//line /usr/local/go/src/strings/strings.go:610
					// _ = "end of CoverTab[1454]"
				} else {
//line /usr/local/go/src/strings/strings.go:611
					_go_fuzz_dep_.CoverTab[1455]++
//line /usr/local/go/src/strings/strings.go:611
					// _ = "end of CoverTab[1455]"
//line /usr/local/go/src/strings/strings.go:611
				}
//line /usr/local/go/src/strings/strings.go:611
				// _ = "end of CoverTab[1452]"
//line /usr/local/go/src/strings/strings.go:611
				_go_fuzz_dep_.CoverTab[1453]++
										b.WriteByte(c)
										pos = i + 1
//line /usr/local/go/src/strings/strings.go:613
				// _ = "end of CoverTab[1453]"
			} else {
//line /usr/local/go/src/strings/strings.go:614
				_go_fuzz_dep_.CoverTab[1456]++
//line /usr/local/go/src/strings/strings.go:614
				// _ = "end of CoverTab[1456]"
//line /usr/local/go/src/strings/strings.go:614
			}
//line /usr/local/go/src/strings/strings.go:614
			// _ = "end of CoverTab[1450]"
		}
//line /usr/local/go/src/strings/strings.go:615
		// _ = "end of CoverTab[1445]"
//line /usr/local/go/src/strings/strings.go:615
		_go_fuzz_dep_.CoverTab[1446]++
								if pos < len(s) {
//line /usr/local/go/src/strings/strings.go:616
			_go_fuzz_dep_.CoverTab[1457]++
									b.WriteString(s[pos:])
//line /usr/local/go/src/strings/strings.go:617
			// _ = "end of CoverTab[1457]"
		} else {
//line /usr/local/go/src/strings/strings.go:618
			_go_fuzz_dep_.CoverTab[1458]++
//line /usr/local/go/src/strings/strings.go:618
			// _ = "end of CoverTab[1458]"
//line /usr/local/go/src/strings/strings.go:618
		}
//line /usr/local/go/src/strings/strings.go:618
		// _ = "end of CoverTab[1446]"
//line /usr/local/go/src/strings/strings.go:618
		_go_fuzz_dep_.CoverTab[1447]++
								return b.String()
//line /usr/local/go/src/strings/strings.go:619
		// _ = "end of CoverTab[1447]"
	} else {
//line /usr/local/go/src/strings/strings.go:620
		_go_fuzz_dep_.CoverTab[1459]++
//line /usr/local/go/src/strings/strings.go:620
		// _ = "end of CoverTab[1459]"
//line /usr/local/go/src/strings/strings.go:620
	}
//line /usr/local/go/src/strings/strings.go:620
	// _ = "end of CoverTab[1436]"
//line /usr/local/go/src/strings/strings.go:620
	_go_fuzz_dep_.CoverTab[1437]++
							return Map(unicode.ToUpper, s)
//line /usr/local/go/src/strings/strings.go:621
	// _ = "end of CoverTab[1437]"
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
//line /usr/local/go/src/strings/strings.go:625
	_go_fuzz_dep_.CoverTab[1460]++
							isASCII, hasUpper := true, false
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:627
		_go_fuzz_dep_.CoverTab[1463]++
								c := s[i]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:629
			_go_fuzz_dep_.CoverTab[1465]++
									isASCII = false
									break
//line /usr/local/go/src/strings/strings.go:631
			// _ = "end of CoverTab[1465]"
		} else {
//line /usr/local/go/src/strings/strings.go:632
			_go_fuzz_dep_.CoverTab[1466]++
//line /usr/local/go/src/strings/strings.go:632
			// _ = "end of CoverTab[1466]"
//line /usr/local/go/src/strings/strings.go:632
		}
//line /usr/local/go/src/strings/strings.go:632
		// _ = "end of CoverTab[1463]"
//line /usr/local/go/src/strings/strings.go:632
		_go_fuzz_dep_.CoverTab[1464]++
								hasUpper = hasUpper || func() bool {
//line /usr/local/go/src/strings/strings.go:633
			_go_fuzz_dep_.CoverTab[1467]++
//line /usr/local/go/src/strings/strings.go:633
			return ('A' <= c && func() bool {
//line /usr/local/go/src/strings/strings.go:633
				_go_fuzz_dep_.CoverTab[1468]++
//line /usr/local/go/src/strings/strings.go:633
				return c <= 'Z'
//line /usr/local/go/src/strings/strings.go:633
				// _ = "end of CoverTab[1468]"
//line /usr/local/go/src/strings/strings.go:633
			}())
//line /usr/local/go/src/strings/strings.go:633
			// _ = "end of CoverTab[1467]"
//line /usr/local/go/src/strings/strings.go:633
		}()
//line /usr/local/go/src/strings/strings.go:633
		// _ = "end of CoverTab[1464]"
	}
//line /usr/local/go/src/strings/strings.go:634
	// _ = "end of CoverTab[1460]"
//line /usr/local/go/src/strings/strings.go:634
	_go_fuzz_dep_.CoverTab[1461]++

							if isASCII {
//line /usr/local/go/src/strings/strings.go:636
		_go_fuzz_dep_.CoverTab[1469]++
								if !hasUpper {
//line /usr/local/go/src/strings/strings.go:637
			_go_fuzz_dep_.CoverTab[1473]++
									return s
//line /usr/local/go/src/strings/strings.go:638
			// _ = "end of CoverTab[1473]"
		} else {
//line /usr/local/go/src/strings/strings.go:639
			_go_fuzz_dep_.CoverTab[1474]++
//line /usr/local/go/src/strings/strings.go:639
			// _ = "end of CoverTab[1474]"
//line /usr/local/go/src/strings/strings.go:639
		}
//line /usr/local/go/src/strings/strings.go:639
		// _ = "end of CoverTab[1469]"
//line /usr/local/go/src/strings/strings.go:639
		_go_fuzz_dep_.CoverTab[1470]++
								var (
			b	Builder
			pos	int
		)
		b.Grow(len(s))
		for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/strings.go:645
			_go_fuzz_dep_.CoverTab[1475]++
									c := s[i]
									if 'A' <= c && func() bool {
//line /usr/local/go/src/strings/strings.go:647
				_go_fuzz_dep_.CoverTab[1476]++
//line /usr/local/go/src/strings/strings.go:647
				return c <= 'Z'
//line /usr/local/go/src/strings/strings.go:647
				// _ = "end of CoverTab[1476]"
//line /usr/local/go/src/strings/strings.go:647
			}() {
//line /usr/local/go/src/strings/strings.go:647
				_go_fuzz_dep_.CoverTab[1477]++
										c += 'a' - 'A'
										if pos < i {
//line /usr/local/go/src/strings/strings.go:649
					_go_fuzz_dep_.CoverTab[1479]++
											b.WriteString(s[pos:i])
//line /usr/local/go/src/strings/strings.go:650
					// _ = "end of CoverTab[1479]"
				} else {
//line /usr/local/go/src/strings/strings.go:651
					_go_fuzz_dep_.CoverTab[1480]++
//line /usr/local/go/src/strings/strings.go:651
					// _ = "end of CoverTab[1480]"
//line /usr/local/go/src/strings/strings.go:651
				}
//line /usr/local/go/src/strings/strings.go:651
				// _ = "end of CoverTab[1477]"
//line /usr/local/go/src/strings/strings.go:651
				_go_fuzz_dep_.CoverTab[1478]++
										b.WriteByte(c)
										pos = i + 1
//line /usr/local/go/src/strings/strings.go:653
				// _ = "end of CoverTab[1478]"
			} else {
//line /usr/local/go/src/strings/strings.go:654
				_go_fuzz_dep_.CoverTab[1481]++
//line /usr/local/go/src/strings/strings.go:654
				// _ = "end of CoverTab[1481]"
//line /usr/local/go/src/strings/strings.go:654
			}
//line /usr/local/go/src/strings/strings.go:654
			// _ = "end of CoverTab[1475]"
		}
//line /usr/local/go/src/strings/strings.go:655
		// _ = "end of CoverTab[1470]"
//line /usr/local/go/src/strings/strings.go:655
		_go_fuzz_dep_.CoverTab[1471]++
								if pos < len(s) {
//line /usr/local/go/src/strings/strings.go:656
			_go_fuzz_dep_.CoverTab[1482]++
									b.WriteString(s[pos:])
//line /usr/local/go/src/strings/strings.go:657
			// _ = "end of CoverTab[1482]"
		} else {
//line /usr/local/go/src/strings/strings.go:658
			_go_fuzz_dep_.CoverTab[1483]++
//line /usr/local/go/src/strings/strings.go:658
			// _ = "end of CoverTab[1483]"
//line /usr/local/go/src/strings/strings.go:658
		}
//line /usr/local/go/src/strings/strings.go:658
		// _ = "end of CoverTab[1471]"
//line /usr/local/go/src/strings/strings.go:658
		_go_fuzz_dep_.CoverTab[1472]++
								return b.String()
//line /usr/local/go/src/strings/strings.go:659
		// _ = "end of CoverTab[1472]"
	} else {
//line /usr/local/go/src/strings/strings.go:660
		_go_fuzz_dep_.CoverTab[1484]++
//line /usr/local/go/src/strings/strings.go:660
		// _ = "end of CoverTab[1484]"
//line /usr/local/go/src/strings/strings.go:660
	}
//line /usr/local/go/src/strings/strings.go:660
	// _ = "end of CoverTab[1461]"
//line /usr/local/go/src/strings/strings.go:660
	_go_fuzz_dep_.CoverTab[1462]++
							return Map(unicode.ToLower, s)
//line /usr/local/go/src/strings/strings.go:661
	// _ = "end of CoverTab[1462]"
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to
//line /usr/local/go/src/strings/strings.go:664
// their Unicode title case.
//line /usr/local/go/src/strings/strings.go:666
func ToTitle(s string) string {
//line /usr/local/go/src/strings/strings.go:666
	_go_fuzz_dep_.CoverTab[1485]++
//line /usr/local/go/src/strings/strings.go:666
	return Map(unicode.ToTitle, s)
//line /usr/local/go/src/strings/strings.go:666
	// _ = "end of CoverTab[1485]"
//line /usr/local/go/src/strings/strings.go:666
}

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
//line /usr/local/go/src/strings/strings.go:668
// upper case using the case mapping specified by c.
//line /usr/local/go/src/strings/strings.go:670
func ToUpperSpecial(c unicode.SpecialCase, s string) string {
//line /usr/local/go/src/strings/strings.go:670
	_go_fuzz_dep_.CoverTab[1486]++
							return Map(c.ToUpper, s)
//line /usr/local/go/src/strings/strings.go:671
	// _ = "end of CoverTab[1486]"
}

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
//line /usr/local/go/src/strings/strings.go:674
// lower case using the case mapping specified by c.
//line /usr/local/go/src/strings/strings.go:676
func ToLowerSpecial(c unicode.SpecialCase, s string) string {
//line /usr/local/go/src/strings/strings.go:676
	_go_fuzz_dep_.CoverTab[1487]++
							return Map(c.ToLower, s)
//line /usr/local/go/src/strings/strings.go:677
	// _ = "end of CoverTab[1487]"
}

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
//line /usr/local/go/src/strings/strings.go:680
// Unicode title case, giving priority to the special casing rules.
//line /usr/local/go/src/strings/strings.go:682
func ToTitleSpecial(c unicode.SpecialCase, s string) string {
//line /usr/local/go/src/strings/strings.go:682
	_go_fuzz_dep_.CoverTab[1488]++
							return Map(c.ToTitle, s)
//line /usr/local/go/src/strings/strings.go:683
	// _ = "end of CoverTab[1488]"
}

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
//line /usr/local/go/src/strings/strings.go:686
// replaced by the replacement string, which may be empty.
//line /usr/local/go/src/strings/strings.go:688
func ToValidUTF8(s, replacement string) string {
//line /usr/local/go/src/strings/strings.go:688
	_go_fuzz_dep_.CoverTab[1489]++
							var b Builder

							for i, c := range s {
//line /usr/local/go/src/strings/strings.go:691
		_go_fuzz_dep_.CoverTab[1493]++
								if c != utf8.RuneError {
//line /usr/local/go/src/strings/strings.go:692
			_go_fuzz_dep_.CoverTab[1495]++
									continue
//line /usr/local/go/src/strings/strings.go:693
			// _ = "end of CoverTab[1495]"
		} else {
//line /usr/local/go/src/strings/strings.go:694
			_go_fuzz_dep_.CoverTab[1496]++
//line /usr/local/go/src/strings/strings.go:694
			// _ = "end of CoverTab[1496]"
//line /usr/local/go/src/strings/strings.go:694
		}
//line /usr/local/go/src/strings/strings.go:694
		// _ = "end of CoverTab[1493]"
//line /usr/local/go/src/strings/strings.go:694
		_go_fuzz_dep_.CoverTab[1494]++

								_, wid := utf8.DecodeRuneInString(s[i:])
								if wid == 1 {
//line /usr/local/go/src/strings/strings.go:697
			_go_fuzz_dep_.CoverTab[1497]++
									b.Grow(len(s) + len(replacement))
									b.WriteString(s[:i])
									s = s[i:]
									break
//line /usr/local/go/src/strings/strings.go:701
			// _ = "end of CoverTab[1497]"
		} else {
//line /usr/local/go/src/strings/strings.go:702
			_go_fuzz_dep_.CoverTab[1498]++
//line /usr/local/go/src/strings/strings.go:702
			// _ = "end of CoverTab[1498]"
//line /usr/local/go/src/strings/strings.go:702
		}
//line /usr/local/go/src/strings/strings.go:702
		// _ = "end of CoverTab[1494]"
	}
//line /usr/local/go/src/strings/strings.go:703
	// _ = "end of CoverTab[1489]"
//line /usr/local/go/src/strings/strings.go:703
	_go_fuzz_dep_.CoverTab[1490]++

//line /usr/local/go/src/strings/strings.go:706
	if b.Cap() == 0 {
//line /usr/local/go/src/strings/strings.go:706
		_go_fuzz_dep_.CoverTab[1499]++
								return s
//line /usr/local/go/src/strings/strings.go:707
		// _ = "end of CoverTab[1499]"
	} else {
//line /usr/local/go/src/strings/strings.go:708
		_go_fuzz_dep_.CoverTab[1500]++
//line /usr/local/go/src/strings/strings.go:708
		// _ = "end of CoverTab[1500]"
//line /usr/local/go/src/strings/strings.go:708
	}
//line /usr/local/go/src/strings/strings.go:708
	// _ = "end of CoverTab[1490]"
//line /usr/local/go/src/strings/strings.go:708
	_go_fuzz_dep_.CoverTab[1491]++

							invalid := false
							for i := 0; i < len(s); {
//line /usr/local/go/src/strings/strings.go:711
		_go_fuzz_dep_.CoverTab[1501]++
								c := s[i]
								if c < utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:713
			_go_fuzz_dep_.CoverTab[1504]++
									i++
									invalid = false
									b.WriteByte(c)
									continue
//line /usr/local/go/src/strings/strings.go:717
			// _ = "end of CoverTab[1504]"
		} else {
//line /usr/local/go/src/strings/strings.go:718
			_go_fuzz_dep_.CoverTab[1505]++
//line /usr/local/go/src/strings/strings.go:718
			// _ = "end of CoverTab[1505]"
//line /usr/local/go/src/strings/strings.go:718
		}
//line /usr/local/go/src/strings/strings.go:718
		// _ = "end of CoverTab[1501]"
//line /usr/local/go/src/strings/strings.go:718
		_go_fuzz_dep_.CoverTab[1502]++
								_, wid := utf8.DecodeRuneInString(s[i:])
								if wid == 1 {
//line /usr/local/go/src/strings/strings.go:720
			_go_fuzz_dep_.CoverTab[1506]++
									i++
									if !invalid {
//line /usr/local/go/src/strings/strings.go:722
				_go_fuzz_dep_.CoverTab[1508]++
										invalid = true
										b.WriteString(replacement)
//line /usr/local/go/src/strings/strings.go:724
				// _ = "end of CoverTab[1508]"
			} else {
//line /usr/local/go/src/strings/strings.go:725
				_go_fuzz_dep_.CoverTab[1509]++
//line /usr/local/go/src/strings/strings.go:725
				// _ = "end of CoverTab[1509]"
//line /usr/local/go/src/strings/strings.go:725
			}
//line /usr/local/go/src/strings/strings.go:725
			// _ = "end of CoverTab[1506]"
//line /usr/local/go/src/strings/strings.go:725
			_go_fuzz_dep_.CoverTab[1507]++
									continue
//line /usr/local/go/src/strings/strings.go:726
			// _ = "end of CoverTab[1507]"
		} else {
//line /usr/local/go/src/strings/strings.go:727
			_go_fuzz_dep_.CoverTab[1510]++
//line /usr/local/go/src/strings/strings.go:727
			// _ = "end of CoverTab[1510]"
//line /usr/local/go/src/strings/strings.go:727
		}
//line /usr/local/go/src/strings/strings.go:727
		// _ = "end of CoverTab[1502]"
//line /usr/local/go/src/strings/strings.go:727
		_go_fuzz_dep_.CoverTab[1503]++
								invalid = false
								b.WriteString(s[i : i+wid])
								i += wid
//line /usr/local/go/src/strings/strings.go:730
		// _ = "end of CoverTab[1503]"
	}
//line /usr/local/go/src/strings/strings.go:731
	// _ = "end of CoverTab[1491]"
//line /usr/local/go/src/strings/strings.go:731
	_go_fuzz_dep_.CoverTab[1492]++

							return b.String()
//line /usr/local/go/src/strings/strings.go:733
	// _ = "end of CoverTab[1492]"
}

// isSeparator reports whether the rune could mark a word boundary.
//line /usr/local/go/src/strings/strings.go:736
// TODO: update when package unicode captures more of the properties.
//line /usr/local/go/src/strings/strings.go:738
func isSeparator(r rune) bool {
//line /usr/local/go/src/strings/strings.go:738
	_go_fuzz_dep_.CoverTab[1511]++

							if r <= 0x7F {
//line /usr/local/go/src/strings/strings.go:740
		_go_fuzz_dep_.CoverTab[1514]++
								switch {
		case '0' <= r && func() bool {
//line /usr/local/go/src/strings/strings.go:742
			_go_fuzz_dep_.CoverTab[1521]++
//line /usr/local/go/src/strings/strings.go:742
			return r <= '9'
//line /usr/local/go/src/strings/strings.go:742
			// _ = "end of CoverTab[1521]"
//line /usr/local/go/src/strings/strings.go:742
		}():
//line /usr/local/go/src/strings/strings.go:742
			_go_fuzz_dep_.CoverTab[1516]++
									return false
//line /usr/local/go/src/strings/strings.go:743
			// _ = "end of CoverTab[1516]"
		case 'a' <= r && func() bool {
//line /usr/local/go/src/strings/strings.go:744
			_go_fuzz_dep_.CoverTab[1522]++
//line /usr/local/go/src/strings/strings.go:744
			return r <= 'z'
//line /usr/local/go/src/strings/strings.go:744
			// _ = "end of CoverTab[1522]"
//line /usr/local/go/src/strings/strings.go:744
		}():
//line /usr/local/go/src/strings/strings.go:744
			_go_fuzz_dep_.CoverTab[1517]++
									return false
//line /usr/local/go/src/strings/strings.go:745
			// _ = "end of CoverTab[1517]"
		case 'A' <= r && func() bool {
//line /usr/local/go/src/strings/strings.go:746
			_go_fuzz_dep_.CoverTab[1523]++
//line /usr/local/go/src/strings/strings.go:746
			return r <= 'Z'
//line /usr/local/go/src/strings/strings.go:746
			// _ = "end of CoverTab[1523]"
//line /usr/local/go/src/strings/strings.go:746
		}():
//line /usr/local/go/src/strings/strings.go:746
			_go_fuzz_dep_.CoverTab[1518]++
									return false
//line /usr/local/go/src/strings/strings.go:747
			// _ = "end of CoverTab[1518]"
		case r == '_':
//line /usr/local/go/src/strings/strings.go:748
			_go_fuzz_dep_.CoverTab[1519]++
									return false
//line /usr/local/go/src/strings/strings.go:749
			// _ = "end of CoverTab[1519]"
//line /usr/local/go/src/strings/strings.go:749
		default:
//line /usr/local/go/src/strings/strings.go:749
			_go_fuzz_dep_.CoverTab[1520]++
//line /usr/local/go/src/strings/strings.go:749
			// _ = "end of CoverTab[1520]"
		}
//line /usr/local/go/src/strings/strings.go:750
		// _ = "end of CoverTab[1514]"
//line /usr/local/go/src/strings/strings.go:750
		_go_fuzz_dep_.CoverTab[1515]++
								return true
//line /usr/local/go/src/strings/strings.go:751
		// _ = "end of CoverTab[1515]"
	} else {
//line /usr/local/go/src/strings/strings.go:752
		_go_fuzz_dep_.CoverTab[1524]++
//line /usr/local/go/src/strings/strings.go:752
		// _ = "end of CoverTab[1524]"
//line /usr/local/go/src/strings/strings.go:752
	}
//line /usr/local/go/src/strings/strings.go:752
	// _ = "end of CoverTab[1511]"
//line /usr/local/go/src/strings/strings.go:752
	_go_fuzz_dep_.CoverTab[1512]++

							if unicode.IsLetter(r) || func() bool {
//line /usr/local/go/src/strings/strings.go:754
		_go_fuzz_dep_.CoverTab[1525]++
//line /usr/local/go/src/strings/strings.go:754
		return unicode.IsDigit(r)
//line /usr/local/go/src/strings/strings.go:754
		// _ = "end of CoverTab[1525]"
//line /usr/local/go/src/strings/strings.go:754
	}() {
//line /usr/local/go/src/strings/strings.go:754
		_go_fuzz_dep_.CoverTab[1526]++
								return false
//line /usr/local/go/src/strings/strings.go:755
		// _ = "end of CoverTab[1526]"
	} else {
//line /usr/local/go/src/strings/strings.go:756
		_go_fuzz_dep_.CoverTab[1527]++
//line /usr/local/go/src/strings/strings.go:756
		// _ = "end of CoverTab[1527]"
//line /usr/local/go/src/strings/strings.go:756
	}
//line /usr/local/go/src/strings/strings.go:756
	// _ = "end of CoverTab[1512]"
//line /usr/local/go/src/strings/strings.go:756
	_go_fuzz_dep_.CoverTab[1513]++

							return unicode.IsSpace(r)
//line /usr/local/go/src/strings/strings.go:758
	// _ = "end of CoverTab[1513]"
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
	_go_fuzz_dep_.CoverTab[1528]++

//line /usr/local/go/src/strings/strings.go:770
	prev := ' '
	return Map(
		func(r rune) rune {
//line /usr/local/go/src/strings/strings.go:772
			_go_fuzz_dep_.CoverTab[1529]++
									if isSeparator(prev) {
//line /usr/local/go/src/strings/strings.go:773
				_go_fuzz_dep_.CoverTab[1531]++
										prev = r
										return unicode.ToTitle(r)
//line /usr/local/go/src/strings/strings.go:775
				// _ = "end of CoverTab[1531]"
			} else {
//line /usr/local/go/src/strings/strings.go:776
				_go_fuzz_dep_.CoverTab[1532]++
//line /usr/local/go/src/strings/strings.go:776
				// _ = "end of CoverTab[1532]"
//line /usr/local/go/src/strings/strings.go:776
			}
//line /usr/local/go/src/strings/strings.go:776
			// _ = "end of CoverTab[1529]"
//line /usr/local/go/src/strings/strings.go:776
			_go_fuzz_dep_.CoverTab[1530]++
									prev = r
									return r
//line /usr/local/go/src/strings/strings.go:778
			// _ = "end of CoverTab[1530]"
		},
		s)
//line /usr/local/go/src/strings/strings.go:780
	// _ = "end of CoverTab[1528]"
}

// TrimLeftFunc returns a slice of the string s with all leading
//line /usr/local/go/src/strings/strings.go:783
// Unicode code points c satisfying f(c) removed.
//line /usr/local/go/src/strings/strings.go:785
func TrimLeftFunc(s string, f func(rune) bool) string {
//line /usr/local/go/src/strings/strings.go:785
	_go_fuzz_dep_.CoverTab[1533]++
							i := indexFunc(s, f, false)
							if i == -1 {
//line /usr/local/go/src/strings/strings.go:787
		_go_fuzz_dep_.CoverTab[1535]++
								return ""
//line /usr/local/go/src/strings/strings.go:788
		// _ = "end of CoverTab[1535]"
	} else {
//line /usr/local/go/src/strings/strings.go:789
		_go_fuzz_dep_.CoverTab[1536]++
//line /usr/local/go/src/strings/strings.go:789
		// _ = "end of CoverTab[1536]"
//line /usr/local/go/src/strings/strings.go:789
	}
//line /usr/local/go/src/strings/strings.go:789
	// _ = "end of CoverTab[1533]"
//line /usr/local/go/src/strings/strings.go:789
	_go_fuzz_dep_.CoverTab[1534]++
							return s[i:]
//line /usr/local/go/src/strings/strings.go:790
	// _ = "end of CoverTab[1534]"
}

// TrimRightFunc returns a slice of the string s with all trailing
//line /usr/local/go/src/strings/strings.go:793
// Unicode code points c satisfying f(c) removed.
//line /usr/local/go/src/strings/strings.go:795
func TrimRightFunc(s string, f func(rune) bool) string {
//line /usr/local/go/src/strings/strings.go:795
	_go_fuzz_dep_.CoverTab[1537]++
							i := lastIndexFunc(s, f, false)
							if i >= 0 && func() bool {
//line /usr/local/go/src/strings/strings.go:797
		_go_fuzz_dep_.CoverTab[1539]++
//line /usr/local/go/src/strings/strings.go:797
		return s[i] >= utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:797
		// _ = "end of CoverTab[1539]"
//line /usr/local/go/src/strings/strings.go:797
	}() {
//line /usr/local/go/src/strings/strings.go:797
		_go_fuzz_dep_.CoverTab[1540]++
								_, wid := utf8.DecodeRuneInString(s[i:])
								i += wid
//line /usr/local/go/src/strings/strings.go:799
		// _ = "end of CoverTab[1540]"
	} else {
//line /usr/local/go/src/strings/strings.go:800
		_go_fuzz_dep_.CoverTab[1541]++
								i++
//line /usr/local/go/src/strings/strings.go:801
		// _ = "end of CoverTab[1541]"
	}
//line /usr/local/go/src/strings/strings.go:802
	// _ = "end of CoverTab[1537]"
//line /usr/local/go/src/strings/strings.go:802
	_go_fuzz_dep_.CoverTab[1538]++
							return s[0:i]
//line /usr/local/go/src/strings/strings.go:803
	// _ = "end of CoverTab[1538]"
}

// TrimFunc returns a slice of the string s with all leading
//line /usr/local/go/src/strings/strings.go:806
// and trailing Unicode code points c satisfying f(c) removed.
//line /usr/local/go/src/strings/strings.go:808
func TrimFunc(s string, f func(rune) bool) string {
//line /usr/local/go/src/strings/strings.go:808
	_go_fuzz_dep_.CoverTab[1542]++
							return TrimRightFunc(TrimLeftFunc(s, f), f)
//line /usr/local/go/src/strings/strings.go:809
	// _ = "end of CoverTab[1542]"
}

// IndexFunc returns the index into s of the first Unicode
//line /usr/local/go/src/strings/strings.go:812
// code point satisfying f(c), or -1 if none do.
//line /usr/local/go/src/strings/strings.go:814
func IndexFunc(s string, f func(rune) bool) int {
//line /usr/local/go/src/strings/strings.go:814
	_go_fuzz_dep_.CoverTab[1543]++
							return indexFunc(s, f, true)
//line /usr/local/go/src/strings/strings.go:815
	// _ = "end of CoverTab[1543]"
}

// LastIndexFunc returns the index into s of the last
//line /usr/local/go/src/strings/strings.go:818
// Unicode code point satisfying f(c), or -1 if none do.
//line /usr/local/go/src/strings/strings.go:820
func LastIndexFunc(s string, f func(rune) bool) int {
//line /usr/local/go/src/strings/strings.go:820
	_go_fuzz_dep_.CoverTab[1544]++
							return lastIndexFunc(s, f, true)
//line /usr/local/go/src/strings/strings.go:821
	// _ = "end of CoverTab[1544]"
}

// indexFunc is the same as IndexFunc except that if
//line /usr/local/go/src/strings/strings.go:824
// truth==false, the sense of the predicate function is
//line /usr/local/go/src/strings/strings.go:824
// inverted.
//line /usr/local/go/src/strings/strings.go:827
func indexFunc(s string, f func(rune) bool, truth bool) int {
//line /usr/local/go/src/strings/strings.go:827
	_go_fuzz_dep_.CoverTab[1545]++
							for i, r := range s {
//line /usr/local/go/src/strings/strings.go:828
		_go_fuzz_dep_.CoverTab[1547]++
								if f(r) == truth {
//line /usr/local/go/src/strings/strings.go:829
			_go_fuzz_dep_.CoverTab[1548]++
									return i
//line /usr/local/go/src/strings/strings.go:830
			// _ = "end of CoverTab[1548]"
		} else {
//line /usr/local/go/src/strings/strings.go:831
			_go_fuzz_dep_.CoverTab[1549]++
//line /usr/local/go/src/strings/strings.go:831
			// _ = "end of CoverTab[1549]"
//line /usr/local/go/src/strings/strings.go:831
		}
//line /usr/local/go/src/strings/strings.go:831
		// _ = "end of CoverTab[1547]"
	}
//line /usr/local/go/src/strings/strings.go:832
	// _ = "end of CoverTab[1545]"
//line /usr/local/go/src/strings/strings.go:832
	_go_fuzz_dep_.CoverTab[1546]++
							return -1
//line /usr/local/go/src/strings/strings.go:833
	// _ = "end of CoverTab[1546]"
}

// lastIndexFunc is the same as LastIndexFunc except that if
//line /usr/local/go/src/strings/strings.go:836
// truth==false, the sense of the predicate function is
//line /usr/local/go/src/strings/strings.go:836
// inverted.
//line /usr/local/go/src/strings/strings.go:839
func lastIndexFunc(s string, f func(rune) bool, truth bool) int {
//line /usr/local/go/src/strings/strings.go:839
	_go_fuzz_dep_.CoverTab[1550]++
							for i := len(s); i > 0; {
//line /usr/local/go/src/strings/strings.go:840
		_go_fuzz_dep_.CoverTab[1552]++
								r, size := utf8.DecodeLastRuneInString(s[0:i])
								i -= size
								if f(r) == truth {
//line /usr/local/go/src/strings/strings.go:843
			_go_fuzz_dep_.CoverTab[1553]++
									return i
//line /usr/local/go/src/strings/strings.go:844
			// _ = "end of CoverTab[1553]"
		} else {
//line /usr/local/go/src/strings/strings.go:845
			_go_fuzz_dep_.CoverTab[1554]++
//line /usr/local/go/src/strings/strings.go:845
			// _ = "end of CoverTab[1554]"
//line /usr/local/go/src/strings/strings.go:845
		}
//line /usr/local/go/src/strings/strings.go:845
		// _ = "end of CoverTab[1552]"
	}
//line /usr/local/go/src/strings/strings.go:846
	// _ = "end of CoverTab[1550]"
//line /usr/local/go/src/strings/strings.go:846
	_go_fuzz_dep_.CoverTab[1551]++
							return -1
//line /usr/local/go/src/strings/strings.go:847
	// _ = "end of CoverTab[1551]"
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
	_go_fuzz_dep_.CoverTab[1555]++
							for i := 0; i < len(chars); i++ {
//line /usr/local/go/src/strings/strings.go:863
		_go_fuzz_dep_.CoverTab[1557]++
								c := chars[i]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:865
			_go_fuzz_dep_.CoverTab[1559]++
									return as, false
//line /usr/local/go/src/strings/strings.go:866
			// _ = "end of CoverTab[1559]"
		} else {
//line /usr/local/go/src/strings/strings.go:867
			_go_fuzz_dep_.CoverTab[1560]++
//line /usr/local/go/src/strings/strings.go:867
			// _ = "end of CoverTab[1560]"
//line /usr/local/go/src/strings/strings.go:867
		}
//line /usr/local/go/src/strings/strings.go:867
		// _ = "end of CoverTab[1557]"
//line /usr/local/go/src/strings/strings.go:867
		_go_fuzz_dep_.CoverTab[1558]++
								as[c/32] |= 1 << (c % 32)
//line /usr/local/go/src/strings/strings.go:868
		// _ = "end of CoverTab[1558]"
	}
//line /usr/local/go/src/strings/strings.go:869
	// _ = "end of CoverTab[1555]"
//line /usr/local/go/src/strings/strings.go:869
	_go_fuzz_dep_.CoverTab[1556]++
							return as, true
//line /usr/local/go/src/strings/strings.go:870
	// _ = "end of CoverTab[1556]"
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c byte) bool {
//line /usr/local/go/src/strings/strings.go:874
	_go_fuzz_dep_.CoverTab[1561]++
							return (as[c/32] & (1 << (c % 32))) != 0
//line /usr/local/go/src/strings/strings.go:875
	// _ = "end of CoverTab[1561]"
}

// Trim returns a slice of the string s with all leading and
//line /usr/local/go/src/strings/strings.go:878
// trailing Unicode code points contained in cutset removed.
//line /usr/local/go/src/strings/strings.go:880
func Trim(s, cutset string) string {
//line /usr/local/go/src/strings/strings.go:880
	_go_fuzz_dep_.CoverTab[1562]++
							if s == "" || func() bool {
//line /usr/local/go/src/strings/strings.go:881
		_go_fuzz_dep_.CoverTab[1566]++
//line /usr/local/go/src/strings/strings.go:881
		return cutset == ""
//line /usr/local/go/src/strings/strings.go:881
		// _ = "end of CoverTab[1566]"
//line /usr/local/go/src/strings/strings.go:881
	}() {
//line /usr/local/go/src/strings/strings.go:881
		_go_fuzz_dep_.CoverTab[1567]++
								return s
//line /usr/local/go/src/strings/strings.go:882
		// _ = "end of CoverTab[1567]"
	} else {
//line /usr/local/go/src/strings/strings.go:883
		_go_fuzz_dep_.CoverTab[1568]++
//line /usr/local/go/src/strings/strings.go:883
		// _ = "end of CoverTab[1568]"
//line /usr/local/go/src/strings/strings.go:883
	}
//line /usr/local/go/src/strings/strings.go:883
	// _ = "end of CoverTab[1562]"
//line /usr/local/go/src/strings/strings.go:883
	_go_fuzz_dep_.CoverTab[1563]++
							if len(cutset) == 1 && func() bool {
//line /usr/local/go/src/strings/strings.go:884
		_go_fuzz_dep_.CoverTab[1569]++
//line /usr/local/go/src/strings/strings.go:884
		return cutset[0] < utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:884
		// _ = "end of CoverTab[1569]"
//line /usr/local/go/src/strings/strings.go:884
	}() {
//line /usr/local/go/src/strings/strings.go:884
		_go_fuzz_dep_.CoverTab[1570]++
								return trimLeftByte(trimRightByte(s, cutset[0]), cutset[0])
//line /usr/local/go/src/strings/strings.go:885
		// _ = "end of CoverTab[1570]"
	} else {
//line /usr/local/go/src/strings/strings.go:886
		_go_fuzz_dep_.CoverTab[1571]++
//line /usr/local/go/src/strings/strings.go:886
		// _ = "end of CoverTab[1571]"
//line /usr/local/go/src/strings/strings.go:886
	}
//line /usr/local/go/src/strings/strings.go:886
	// _ = "end of CoverTab[1563]"
//line /usr/local/go/src/strings/strings.go:886
	_go_fuzz_dep_.CoverTab[1564]++
							if as, ok := makeASCIISet(cutset); ok {
//line /usr/local/go/src/strings/strings.go:887
		_go_fuzz_dep_.CoverTab[1572]++
								return trimLeftASCII(trimRightASCII(s, &as), &as)
//line /usr/local/go/src/strings/strings.go:888
		// _ = "end of CoverTab[1572]"
	} else {
//line /usr/local/go/src/strings/strings.go:889
		_go_fuzz_dep_.CoverTab[1573]++
//line /usr/local/go/src/strings/strings.go:889
		// _ = "end of CoverTab[1573]"
//line /usr/local/go/src/strings/strings.go:889
	}
//line /usr/local/go/src/strings/strings.go:889
	// _ = "end of CoverTab[1564]"
//line /usr/local/go/src/strings/strings.go:889
	_go_fuzz_dep_.CoverTab[1565]++
							return trimLeftUnicode(trimRightUnicode(s, cutset), cutset)
//line /usr/local/go/src/strings/strings.go:890
	// _ = "end of CoverTab[1565]"
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
	_go_fuzz_dep_.CoverTab[1574]++
							if s == "" || func() bool {
//line /usr/local/go/src/strings/strings.go:898
		_go_fuzz_dep_.CoverTab[1578]++
//line /usr/local/go/src/strings/strings.go:898
		return cutset == ""
//line /usr/local/go/src/strings/strings.go:898
		// _ = "end of CoverTab[1578]"
//line /usr/local/go/src/strings/strings.go:898
	}() {
//line /usr/local/go/src/strings/strings.go:898
		_go_fuzz_dep_.CoverTab[1579]++
								return s
//line /usr/local/go/src/strings/strings.go:899
		// _ = "end of CoverTab[1579]"
	} else {
//line /usr/local/go/src/strings/strings.go:900
		_go_fuzz_dep_.CoverTab[1580]++
//line /usr/local/go/src/strings/strings.go:900
		// _ = "end of CoverTab[1580]"
//line /usr/local/go/src/strings/strings.go:900
	}
//line /usr/local/go/src/strings/strings.go:900
	// _ = "end of CoverTab[1574]"
//line /usr/local/go/src/strings/strings.go:900
	_go_fuzz_dep_.CoverTab[1575]++
							if len(cutset) == 1 && func() bool {
//line /usr/local/go/src/strings/strings.go:901
		_go_fuzz_dep_.CoverTab[1581]++
//line /usr/local/go/src/strings/strings.go:901
		return cutset[0] < utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:901
		// _ = "end of CoverTab[1581]"
//line /usr/local/go/src/strings/strings.go:901
	}() {
//line /usr/local/go/src/strings/strings.go:901
		_go_fuzz_dep_.CoverTab[1582]++
								return trimLeftByte(s, cutset[0])
//line /usr/local/go/src/strings/strings.go:902
		// _ = "end of CoverTab[1582]"
	} else {
//line /usr/local/go/src/strings/strings.go:903
		_go_fuzz_dep_.CoverTab[1583]++
//line /usr/local/go/src/strings/strings.go:903
		// _ = "end of CoverTab[1583]"
//line /usr/local/go/src/strings/strings.go:903
	}
//line /usr/local/go/src/strings/strings.go:903
	// _ = "end of CoverTab[1575]"
//line /usr/local/go/src/strings/strings.go:903
	_go_fuzz_dep_.CoverTab[1576]++
							if as, ok := makeASCIISet(cutset); ok {
//line /usr/local/go/src/strings/strings.go:904
		_go_fuzz_dep_.CoverTab[1584]++
								return trimLeftASCII(s, &as)
//line /usr/local/go/src/strings/strings.go:905
		// _ = "end of CoverTab[1584]"
	} else {
//line /usr/local/go/src/strings/strings.go:906
		_go_fuzz_dep_.CoverTab[1585]++
//line /usr/local/go/src/strings/strings.go:906
		// _ = "end of CoverTab[1585]"
//line /usr/local/go/src/strings/strings.go:906
	}
//line /usr/local/go/src/strings/strings.go:906
	// _ = "end of CoverTab[1576]"
//line /usr/local/go/src/strings/strings.go:906
	_go_fuzz_dep_.CoverTab[1577]++
							return trimLeftUnicode(s, cutset)
//line /usr/local/go/src/strings/strings.go:907
	// _ = "end of CoverTab[1577]"
}

func trimLeftByte(s string, c byte) string {
//line /usr/local/go/src/strings/strings.go:910
	_go_fuzz_dep_.CoverTab[1586]++
							for len(s) > 0 && func() bool {
//line /usr/local/go/src/strings/strings.go:911
		_go_fuzz_dep_.CoverTab[1588]++
//line /usr/local/go/src/strings/strings.go:911
		return s[0] == c
//line /usr/local/go/src/strings/strings.go:911
		// _ = "end of CoverTab[1588]"
//line /usr/local/go/src/strings/strings.go:911
	}() {
//line /usr/local/go/src/strings/strings.go:911
		_go_fuzz_dep_.CoverTab[1589]++
								s = s[1:]
//line /usr/local/go/src/strings/strings.go:912
		// _ = "end of CoverTab[1589]"
	}
//line /usr/local/go/src/strings/strings.go:913
	// _ = "end of CoverTab[1586]"
//line /usr/local/go/src/strings/strings.go:913
	_go_fuzz_dep_.CoverTab[1587]++
							return s
//line /usr/local/go/src/strings/strings.go:914
	// _ = "end of CoverTab[1587]"
}

func trimLeftASCII(s string, as *asciiSet) string {
//line /usr/local/go/src/strings/strings.go:917
	_go_fuzz_dep_.CoverTab[1590]++
							for len(s) > 0 {
//line /usr/local/go/src/strings/strings.go:918
		_go_fuzz_dep_.CoverTab[1592]++
								if !as.contains(s[0]) {
//line /usr/local/go/src/strings/strings.go:919
			_go_fuzz_dep_.CoverTab[1594]++
									break
//line /usr/local/go/src/strings/strings.go:920
			// _ = "end of CoverTab[1594]"
		} else {
//line /usr/local/go/src/strings/strings.go:921
			_go_fuzz_dep_.CoverTab[1595]++
//line /usr/local/go/src/strings/strings.go:921
			// _ = "end of CoverTab[1595]"
//line /usr/local/go/src/strings/strings.go:921
		}
//line /usr/local/go/src/strings/strings.go:921
		// _ = "end of CoverTab[1592]"
//line /usr/local/go/src/strings/strings.go:921
		_go_fuzz_dep_.CoverTab[1593]++
								s = s[1:]
//line /usr/local/go/src/strings/strings.go:922
		// _ = "end of CoverTab[1593]"
	}
//line /usr/local/go/src/strings/strings.go:923
	// _ = "end of CoverTab[1590]"
//line /usr/local/go/src/strings/strings.go:923
	_go_fuzz_dep_.CoverTab[1591]++
							return s
//line /usr/local/go/src/strings/strings.go:924
	// _ = "end of CoverTab[1591]"
}

func trimLeftUnicode(s, cutset string) string {
//line /usr/local/go/src/strings/strings.go:927
	_go_fuzz_dep_.CoverTab[1596]++
							for len(s) > 0 {
//line /usr/local/go/src/strings/strings.go:928
		_go_fuzz_dep_.CoverTab[1598]++
								r, n := rune(s[0]), 1
								if r >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:930
			_go_fuzz_dep_.CoverTab[1601]++
									r, n = utf8.DecodeRuneInString(s)
//line /usr/local/go/src/strings/strings.go:931
			// _ = "end of CoverTab[1601]"
		} else {
//line /usr/local/go/src/strings/strings.go:932
			_go_fuzz_dep_.CoverTab[1602]++
//line /usr/local/go/src/strings/strings.go:932
			// _ = "end of CoverTab[1602]"
//line /usr/local/go/src/strings/strings.go:932
		}
//line /usr/local/go/src/strings/strings.go:932
		// _ = "end of CoverTab[1598]"
//line /usr/local/go/src/strings/strings.go:932
		_go_fuzz_dep_.CoverTab[1599]++
								if !ContainsRune(cutset, r) {
//line /usr/local/go/src/strings/strings.go:933
			_go_fuzz_dep_.CoverTab[1603]++
									break
//line /usr/local/go/src/strings/strings.go:934
			// _ = "end of CoverTab[1603]"
		} else {
//line /usr/local/go/src/strings/strings.go:935
			_go_fuzz_dep_.CoverTab[1604]++
//line /usr/local/go/src/strings/strings.go:935
			// _ = "end of CoverTab[1604]"
//line /usr/local/go/src/strings/strings.go:935
		}
//line /usr/local/go/src/strings/strings.go:935
		// _ = "end of CoverTab[1599]"
//line /usr/local/go/src/strings/strings.go:935
		_go_fuzz_dep_.CoverTab[1600]++
								s = s[n:]
//line /usr/local/go/src/strings/strings.go:936
		// _ = "end of CoverTab[1600]"
	}
//line /usr/local/go/src/strings/strings.go:937
	// _ = "end of CoverTab[1596]"
//line /usr/local/go/src/strings/strings.go:937
	_go_fuzz_dep_.CoverTab[1597]++
							return s
//line /usr/local/go/src/strings/strings.go:938
	// _ = "end of CoverTab[1597]"
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
	_go_fuzz_dep_.CoverTab[1605]++
							if s == "" || func() bool {
//line /usr/local/go/src/strings/strings.go:946
		_go_fuzz_dep_.CoverTab[1609]++
//line /usr/local/go/src/strings/strings.go:946
		return cutset == ""
//line /usr/local/go/src/strings/strings.go:946
		// _ = "end of CoverTab[1609]"
//line /usr/local/go/src/strings/strings.go:946
	}() {
//line /usr/local/go/src/strings/strings.go:946
		_go_fuzz_dep_.CoverTab[1610]++
								return s
//line /usr/local/go/src/strings/strings.go:947
		// _ = "end of CoverTab[1610]"
	} else {
//line /usr/local/go/src/strings/strings.go:948
		_go_fuzz_dep_.CoverTab[1611]++
//line /usr/local/go/src/strings/strings.go:948
		// _ = "end of CoverTab[1611]"
//line /usr/local/go/src/strings/strings.go:948
	}
//line /usr/local/go/src/strings/strings.go:948
	// _ = "end of CoverTab[1605]"
//line /usr/local/go/src/strings/strings.go:948
	_go_fuzz_dep_.CoverTab[1606]++
							if len(cutset) == 1 && func() bool {
//line /usr/local/go/src/strings/strings.go:949
		_go_fuzz_dep_.CoverTab[1612]++
//line /usr/local/go/src/strings/strings.go:949
		return cutset[0] < utf8.RuneSelf
//line /usr/local/go/src/strings/strings.go:949
		// _ = "end of CoverTab[1612]"
//line /usr/local/go/src/strings/strings.go:949
	}() {
//line /usr/local/go/src/strings/strings.go:949
		_go_fuzz_dep_.CoverTab[1613]++
								return trimRightByte(s, cutset[0])
//line /usr/local/go/src/strings/strings.go:950
		// _ = "end of CoverTab[1613]"
	} else {
//line /usr/local/go/src/strings/strings.go:951
		_go_fuzz_dep_.CoverTab[1614]++
//line /usr/local/go/src/strings/strings.go:951
		// _ = "end of CoverTab[1614]"
//line /usr/local/go/src/strings/strings.go:951
	}
//line /usr/local/go/src/strings/strings.go:951
	// _ = "end of CoverTab[1606]"
//line /usr/local/go/src/strings/strings.go:951
	_go_fuzz_dep_.CoverTab[1607]++
							if as, ok := makeASCIISet(cutset); ok {
//line /usr/local/go/src/strings/strings.go:952
		_go_fuzz_dep_.CoverTab[1615]++
								return trimRightASCII(s, &as)
//line /usr/local/go/src/strings/strings.go:953
		// _ = "end of CoverTab[1615]"
	} else {
//line /usr/local/go/src/strings/strings.go:954
		_go_fuzz_dep_.CoverTab[1616]++
//line /usr/local/go/src/strings/strings.go:954
		// _ = "end of CoverTab[1616]"
//line /usr/local/go/src/strings/strings.go:954
	}
//line /usr/local/go/src/strings/strings.go:954
	// _ = "end of CoverTab[1607]"
//line /usr/local/go/src/strings/strings.go:954
	_go_fuzz_dep_.CoverTab[1608]++
							return trimRightUnicode(s, cutset)
//line /usr/local/go/src/strings/strings.go:955
	// _ = "end of CoverTab[1608]"
}

func trimRightByte(s string, c byte) string {
//line /usr/local/go/src/strings/strings.go:958
	_go_fuzz_dep_.CoverTab[1617]++
							for len(s) > 0 && func() bool {
//line /usr/local/go/src/strings/strings.go:959
		_go_fuzz_dep_.CoverTab[1619]++
//line /usr/local/go/src/strings/strings.go:959
		return s[len(s)-1] == c
//line /usr/local/go/src/strings/strings.go:959
		// _ = "end of CoverTab[1619]"
//line /usr/local/go/src/strings/strings.go:959
	}() {
//line /usr/local/go/src/strings/strings.go:959
		_go_fuzz_dep_.CoverTab[1620]++
								s = s[:len(s)-1]
//line /usr/local/go/src/strings/strings.go:960
		// _ = "end of CoverTab[1620]"
	}
//line /usr/local/go/src/strings/strings.go:961
	// _ = "end of CoverTab[1617]"
//line /usr/local/go/src/strings/strings.go:961
	_go_fuzz_dep_.CoverTab[1618]++
							return s
//line /usr/local/go/src/strings/strings.go:962
	// _ = "end of CoverTab[1618]"
}

func trimRightASCII(s string, as *asciiSet) string {
//line /usr/local/go/src/strings/strings.go:965
	_go_fuzz_dep_.CoverTab[1621]++
							for len(s) > 0 {
//line /usr/local/go/src/strings/strings.go:966
		_go_fuzz_dep_.CoverTab[1623]++
								if !as.contains(s[len(s)-1]) {
//line /usr/local/go/src/strings/strings.go:967
			_go_fuzz_dep_.CoverTab[1625]++
									break
//line /usr/local/go/src/strings/strings.go:968
			// _ = "end of CoverTab[1625]"
		} else {
//line /usr/local/go/src/strings/strings.go:969
			_go_fuzz_dep_.CoverTab[1626]++
//line /usr/local/go/src/strings/strings.go:969
			// _ = "end of CoverTab[1626]"
//line /usr/local/go/src/strings/strings.go:969
		}
//line /usr/local/go/src/strings/strings.go:969
		// _ = "end of CoverTab[1623]"
//line /usr/local/go/src/strings/strings.go:969
		_go_fuzz_dep_.CoverTab[1624]++
								s = s[:len(s)-1]
//line /usr/local/go/src/strings/strings.go:970
		// _ = "end of CoverTab[1624]"
	}
//line /usr/local/go/src/strings/strings.go:971
	// _ = "end of CoverTab[1621]"
//line /usr/local/go/src/strings/strings.go:971
	_go_fuzz_dep_.CoverTab[1622]++
							return s
//line /usr/local/go/src/strings/strings.go:972
	// _ = "end of CoverTab[1622]"
}

func trimRightUnicode(s, cutset string) string {
//line /usr/local/go/src/strings/strings.go:975
	_go_fuzz_dep_.CoverTab[1627]++
							for len(s) > 0 {
//line /usr/local/go/src/strings/strings.go:976
		_go_fuzz_dep_.CoverTab[1629]++
								r, n := rune(s[len(s)-1]), 1
								if r >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:978
			_go_fuzz_dep_.CoverTab[1632]++
									r, n = utf8.DecodeLastRuneInString(s)
//line /usr/local/go/src/strings/strings.go:979
			// _ = "end of CoverTab[1632]"
		} else {
//line /usr/local/go/src/strings/strings.go:980
			_go_fuzz_dep_.CoverTab[1633]++
//line /usr/local/go/src/strings/strings.go:980
			// _ = "end of CoverTab[1633]"
//line /usr/local/go/src/strings/strings.go:980
		}
//line /usr/local/go/src/strings/strings.go:980
		// _ = "end of CoverTab[1629]"
//line /usr/local/go/src/strings/strings.go:980
		_go_fuzz_dep_.CoverTab[1630]++
								if !ContainsRune(cutset, r) {
//line /usr/local/go/src/strings/strings.go:981
			_go_fuzz_dep_.CoverTab[1634]++
									break
//line /usr/local/go/src/strings/strings.go:982
			// _ = "end of CoverTab[1634]"
		} else {
//line /usr/local/go/src/strings/strings.go:983
			_go_fuzz_dep_.CoverTab[1635]++
//line /usr/local/go/src/strings/strings.go:983
			// _ = "end of CoverTab[1635]"
//line /usr/local/go/src/strings/strings.go:983
		}
//line /usr/local/go/src/strings/strings.go:983
		// _ = "end of CoverTab[1630]"
//line /usr/local/go/src/strings/strings.go:983
		_go_fuzz_dep_.CoverTab[1631]++
								s = s[:len(s)-n]
//line /usr/local/go/src/strings/strings.go:984
		// _ = "end of CoverTab[1631]"
	}
//line /usr/local/go/src/strings/strings.go:985
	// _ = "end of CoverTab[1627]"
//line /usr/local/go/src/strings/strings.go:985
	_go_fuzz_dep_.CoverTab[1628]++
							return s
//line /usr/local/go/src/strings/strings.go:986
	// _ = "end of CoverTab[1628]"
}

// TrimSpace returns a slice of the string s, with all leading
//line /usr/local/go/src/strings/strings.go:989
// and trailing white space removed, as defined by Unicode.
//line /usr/local/go/src/strings/strings.go:991
func TrimSpace(s string) string {
//line /usr/local/go/src/strings/strings.go:991
	_go_fuzz_dep_.CoverTab[1636]++

							start := 0
							for ; start < len(s); start++ {
//line /usr/local/go/src/strings/strings.go:994
		_go_fuzz_dep_.CoverTab[1639]++
								c := s[start]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:996
			_go_fuzz_dep_.CoverTab[1641]++

//line /usr/local/go/src/strings/strings.go:999
			return TrimFunc(s[start:], unicode.IsSpace)
//line /usr/local/go/src/strings/strings.go:999
			// _ = "end of CoverTab[1641]"
		} else {
//line /usr/local/go/src/strings/strings.go:1000
			_go_fuzz_dep_.CoverTab[1642]++
//line /usr/local/go/src/strings/strings.go:1000
			// _ = "end of CoverTab[1642]"
//line /usr/local/go/src/strings/strings.go:1000
		}
//line /usr/local/go/src/strings/strings.go:1000
		// _ = "end of CoverTab[1639]"
//line /usr/local/go/src/strings/strings.go:1000
		_go_fuzz_dep_.CoverTab[1640]++
								if asciiSpace[c] == 0 {
//line /usr/local/go/src/strings/strings.go:1001
			_go_fuzz_dep_.CoverTab[1643]++
									break
//line /usr/local/go/src/strings/strings.go:1002
			// _ = "end of CoverTab[1643]"
		} else {
//line /usr/local/go/src/strings/strings.go:1003
			_go_fuzz_dep_.CoverTab[1644]++
//line /usr/local/go/src/strings/strings.go:1003
			// _ = "end of CoverTab[1644]"
//line /usr/local/go/src/strings/strings.go:1003
		}
//line /usr/local/go/src/strings/strings.go:1003
		// _ = "end of CoverTab[1640]"
	}
//line /usr/local/go/src/strings/strings.go:1004
	// _ = "end of CoverTab[1636]"
//line /usr/local/go/src/strings/strings.go:1004
	_go_fuzz_dep_.CoverTab[1637]++

//line /usr/local/go/src/strings/strings.go:1007
	stop := len(s)
	for ; stop > start; stop-- {
//line /usr/local/go/src/strings/strings.go:1008
		_go_fuzz_dep_.CoverTab[1645]++
								c := s[stop-1]
								if c >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:1010
			_go_fuzz_dep_.CoverTab[1647]++

									return TrimRightFunc(s[start:stop], unicode.IsSpace)
//line /usr/local/go/src/strings/strings.go:1012
			// _ = "end of CoverTab[1647]"
		} else {
//line /usr/local/go/src/strings/strings.go:1013
			_go_fuzz_dep_.CoverTab[1648]++
//line /usr/local/go/src/strings/strings.go:1013
			// _ = "end of CoverTab[1648]"
//line /usr/local/go/src/strings/strings.go:1013
		}
//line /usr/local/go/src/strings/strings.go:1013
		// _ = "end of CoverTab[1645]"
//line /usr/local/go/src/strings/strings.go:1013
		_go_fuzz_dep_.CoverTab[1646]++
								if asciiSpace[c] == 0 {
//line /usr/local/go/src/strings/strings.go:1014
			_go_fuzz_dep_.CoverTab[1649]++
									break
//line /usr/local/go/src/strings/strings.go:1015
			// _ = "end of CoverTab[1649]"
		} else {
//line /usr/local/go/src/strings/strings.go:1016
			_go_fuzz_dep_.CoverTab[1650]++
//line /usr/local/go/src/strings/strings.go:1016
			// _ = "end of CoverTab[1650]"
//line /usr/local/go/src/strings/strings.go:1016
		}
//line /usr/local/go/src/strings/strings.go:1016
		// _ = "end of CoverTab[1646]"
	}
//line /usr/local/go/src/strings/strings.go:1017
	// _ = "end of CoverTab[1637]"
//line /usr/local/go/src/strings/strings.go:1017
	_go_fuzz_dep_.CoverTab[1638]++

//line /usr/local/go/src/strings/strings.go:1022
	return s[start:stop]
//line /usr/local/go/src/strings/strings.go:1022
	// _ = "end of CoverTab[1638]"
}

// TrimPrefix returns s without the provided leading prefix string.
//line /usr/local/go/src/strings/strings.go:1025
// If s doesn't start with prefix, s is returned unchanged.
//line /usr/local/go/src/strings/strings.go:1027
func TrimPrefix(s, prefix string) string {
//line /usr/local/go/src/strings/strings.go:1027
	_go_fuzz_dep_.CoverTab[1651]++
							if HasPrefix(s, prefix) {
//line /usr/local/go/src/strings/strings.go:1028
		_go_fuzz_dep_.CoverTab[1653]++
								return s[len(prefix):]
//line /usr/local/go/src/strings/strings.go:1029
		// _ = "end of CoverTab[1653]"
	} else {
//line /usr/local/go/src/strings/strings.go:1030
		_go_fuzz_dep_.CoverTab[1654]++
//line /usr/local/go/src/strings/strings.go:1030
		// _ = "end of CoverTab[1654]"
//line /usr/local/go/src/strings/strings.go:1030
	}
//line /usr/local/go/src/strings/strings.go:1030
	// _ = "end of CoverTab[1651]"
//line /usr/local/go/src/strings/strings.go:1030
	_go_fuzz_dep_.CoverTab[1652]++
							return s
//line /usr/local/go/src/strings/strings.go:1031
	// _ = "end of CoverTab[1652]"
}

// TrimSuffix returns s without the provided trailing suffix string.
//line /usr/local/go/src/strings/strings.go:1034
// If s doesn't end with suffix, s is returned unchanged.
//line /usr/local/go/src/strings/strings.go:1036
func TrimSuffix(s, suffix string) string {
//line /usr/local/go/src/strings/strings.go:1036
	_go_fuzz_dep_.CoverTab[1655]++
							if HasSuffix(s, suffix) {
//line /usr/local/go/src/strings/strings.go:1037
		_go_fuzz_dep_.CoverTab[1657]++
								return s[:len(s)-len(suffix)]
//line /usr/local/go/src/strings/strings.go:1038
		// _ = "end of CoverTab[1657]"
	} else {
//line /usr/local/go/src/strings/strings.go:1039
		_go_fuzz_dep_.CoverTab[1658]++
//line /usr/local/go/src/strings/strings.go:1039
		// _ = "end of CoverTab[1658]"
//line /usr/local/go/src/strings/strings.go:1039
	}
//line /usr/local/go/src/strings/strings.go:1039
	// _ = "end of CoverTab[1655]"
//line /usr/local/go/src/strings/strings.go:1039
	_go_fuzz_dep_.CoverTab[1656]++
							return s
//line /usr/local/go/src/strings/strings.go:1040
	// _ = "end of CoverTab[1656]"
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
	_go_fuzz_dep_.CoverTab[1659]++
							if old == new || func() bool {
//line /usr/local/go/src/strings/strings.go:1050
		_go_fuzz_dep_.CoverTab[1663]++
//line /usr/local/go/src/strings/strings.go:1050
		return n == 0
//line /usr/local/go/src/strings/strings.go:1050
		// _ = "end of CoverTab[1663]"
//line /usr/local/go/src/strings/strings.go:1050
	}() {
//line /usr/local/go/src/strings/strings.go:1050
		_go_fuzz_dep_.CoverTab[1664]++
								return s
//line /usr/local/go/src/strings/strings.go:1051
		// _ = "end of CoverTab[1664]"
	} else {
//line /usr/local/go/src/strings/strings.go:1052
		_go_fuzz_dep_.CoverTab[1665]++
//line /usr/local/go/src/strings/strings.go:1052
		// _ = "end of CoverTab[1665]"
//line /usr/local/go/src/strings/strings.go:1052
	}
//line /usr/local/go/src/strings/strings.go:1052
	// _ = "end of CoverTab[1659]"
//line /usr/local/go/src/strings/strings.go:1052
	_go_fuzz_dep_.CoverTab[1660]++

//line /usr/local/go/src/strings/strings.go:1055
	if m := Count(s, old); m == 0 {
//line /usr/local/go/src/strings/strings.go:1055
		_go_fuzz_dep_.CoverTab[1666]++
								return s
//line /usr/local/go/src/strings/strings.go:1056
		// _ = "end of CoverTab[1666]"
	} else {
//line /usr/local/go/src/strings/strings.go:1057
		_go_fuzz_dep_.CoverTab[1667]++
//line /usr/local/go/src/strings/strings.go:1057
		if n < 0 || func() bool {
//line /usr/local/go/src/strings/strings.go:1057
			_go_fuzz_dep_.CoverTab[1668]++
//line /usr/local/go/src/strings/strings.go:1057
			return m < n
//line /usr/local/go/src/strings/strings.go:1057
			// _ = "end of CoverTab[1668]"
//line /usr/local/go/src/strings/strings.go:1057
		}() {
//line /usr/local/go/src/strings/strings.go:1057
			_go_fuzz_dep_.CoverTab[1669]++
									n = m
//line /usr/local/go/src/strings/strings.go:1058
			// _ = "end of CoverTab[1669]"
		} else {
//line /usr/local/go/src/strings/strings.go:1059
			_go_fuzz_dep_.CoverTab[1670]++
//line /usr/local/go/src/strings/strings.go:1059
			// _ = "end of CoverTab[1670]"
//line /usr/local/go/src/strings/strings.go:1059
		}
//line /usr/local/go/src/strings/strings.go:1059
		// _ = "end of CoverTab[1667]"
//line /usr/local/go/src/strings/strings.go:1059
	}
//line /usr/local/go/src/strings/strings.go:1059
	// _ = "end of CoverTab[1660]"
//line /usr/local/go/src/strings/strings.go:1059
	_go_fuzz_dep_.CoverTab[1661]++

	// Apply replacements to buffer.
	var b Builder
	b.Grow(len(s) + n*(len(new)-len(old)))
	start := 0
	for i := 0; i < n; i++ {
//line /usr/local/go/src/strings/strings.go:1065
		_go_fuzz_dep_.CoverTab[1671]++
								j := start
								if len(old) == 0 {
//line /usr/local/go/src/strings/strings.go:1067
			_go_fuzz_dep_.CoverTab[1673]++
									if i > 0 {
//line /usr/local/go/src/strings/strings.go:1068
				_go_fuzz_dep_.CoverTab[1674]++
										_, wid := utf8.DecodeRuneInString(s[start:])
										j += wid
//line /usr/local/go/src/strings/strings.go:1070
				// _ = "end of CoverTab[1674]"
			} else {
//line /usr/local/go/src/strings/strings.go:1071
				_go_fuzz_dep_.CoverTab[1675]++
//line /usr/local/go/src/strings/strings.go:1071
				// _ = "end of CoverTab[1675]"
//line /usr/local/go/src/strings/strings.go:1071
			}
//line /usr/local/go/src/strings/strings.go:1071
			// _ = "end of CoverTab[1673]"
		} else {
//line /usr/local/go/src/strings/strings.go:1072
			_go_fuzz_dep_.CoverTab[1676]++
									j += Index(s[start:], old)
//line /usr/local/go/src/strings/strings.go:1073
			// _ = "end of CoverTab[1676]"
		}
//line /usr/local/go/src/strings/strings.go:1074
		// _ = "end of CoverTab[1671]"
//line /usr/local/go/src/strings/strings.go:1074
		_go_fuzz_dep_.CoverTab[1672]++
								b.WriteString(s[start:j])
								b.WriteString(new)
								start = j + len(old)
//line /usr/local/go/src/strings/strings.go:1077
		// _ = "end of CoverTab[1672]"
	}
//line /usr/local/go/src/strings/strings.go:1078
	// _ = "end of CoverTab[1661]"
//line /usr/local/go/src/strings/strings.go:1078
	_go_fuzz_dep_.CoverTab[1662]++
							b.WriteString(s[start:])
							return b.String()
//line /usr/local/go/src/strings/strings.go:1080
	// _ = "end of CoverTab[1662]"
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
	_go_fuzz_dep_.CoverTab[1677]++
							return Replace(s, old, new, -1)
//line /usr/local/go/src/strings/strings.go:1089
	// _ = "end of CoverTab[1677]"
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
//line /usr/local/go/src/strings/strings.go:1092
// are equal under simple Unicode case-folding, which is a more general
//line /usr/local/go/src/strings/strings.go:1092
// form of case-insensitivity.
//line /usr/local/go/src/strings/strings.go:1095
func EqualFold(s, t string) bool {
//line /usr/local/go/src/strings/strings.go:1095
	_go_fuzz_dep_.CoverTab[1678]++

							i := 0
							for ; i < len(s) && func() bool {
//line /usr/local/go/src/strings/strings.go:1098
		_go_fuzz_dep_.CoverTab[1681]++
//line /usr/local/go/src/strings/strings.go:1098
		return i < len(t)
//line /usr/local/go/src/strings/strings.go:1098
		// _ = "end of CoverTab[1681]"
//line /usr/local/go/src/strings/strings.go:1098
	}(); i++ {
//line /usr/local/go/src/strings/strings.go:1098
		_go_fuzz_dep_.CoverTab[1682]++
								sr := s[i]
								tr := t[i]
								if sr|tr >= utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:1101
			_go_fuzz_dep_.CoverTab[1687]++
									goto hasUnicode
//line /usr/local/go/src/strings/strings.go:1102
			// _ = "end of CoverTab[1687]"
		} else {
//line /usr/local/go/src/strings/strings.go:1103
			_go_fuzz_dep_.CoverTab[1688]++
//line /usr/local/go/src/strings/strings.go:1103
			// _ = "end of CoverTab[1688]"
//line /usr/local/go/src/strings/strings.go:1103
		}
//line /usr/local/go/src/strings/strings.go:1103
		// _ = "end of CoverTab[1682]"
//line /usr/local/go/src/strings/strings.go:1103
		_go_fuzz_dep_.CoverTab[1683]++

//line /usr/local/go/src/strings/strings.go:1106
		if tr == sr {
//line /usr/local/go/src/strings/strings.go:1106
			_go_fuzz_dep_.CoverTab[1689]++
									continue
//line /usr/local/go/src/strings/strings.go:1107
			// _ = "end of CoverTab[1689]"
		} else {
//line /usr/local/go/src/strings/strings.go:1108
			_go_fuzz_dep_.CoverTab[1690]++
//line /usr/local/go/src/strings/strings.go:1108
			// _ = "end of CoverTab[1690]"
//line /usr/local/go/src/strings/strings.go:1108
		}
//line /usr/local/go/src/strings/strings.go:1108
		// _ = "end of CoverTab[1683]"
//line /usr/local/go/src/strings/strings.go:1108
		_go_fuzz_dep_.CoverTab[1684]++

//line /usr/local/go/src/strings/strings.go:1111
		if tr < sr {
//line /usr/local/go/src/strings/strings.go:1111
			_go_fuzz_dep_.CoverTab[1691]++
									tr, sr = sr, tr
//line /usr/local/go/src/strings/strings.go:1112
			// _ = "end of CoverTab[1691]"
		} else {
//line /usr/local/go/src/strings/strings.go:1113
			_go_fuzz_dep_.CoverTab[1692]++
//line /usr/local/go/src/strings/strings.go:1113
			// _ = "end of CoverTab[1692]"
//line /usr/local/go/src/strings/strings.go:1113
		}
//line /usr/local/go/src/strings/strings.go:1113
		// _ = "end of CoverTab[1684]"
//line /usr/local/go/src/strings/strings.go:1113
		_go_fuzz_dep_.CoverTab[1685]++

								if 'A' <= sr && func() bool {
//line /usr/local/go/src/strings/strings.go:1115
			_go_fuzz_dep_.CoverTab[1693]++
//line /usr/local/go/src/strings/strings.go:1115
			return sr <= 'Z'
//line /usr/local/go/src/strings/strings.go:1115
			// _ = "end of CoverTab[1693]"
//line /usr/local/go/src/strings/strings.go:1115
		}() && func() bool {
//line /usr/local/go/src/strings/strings.go:1115
			_go_fuzz_dep_.CoverTab[1694]++
//line /usr/local/go/src/strings/strings.go:1115
			return tr == sr+'a'-'A'
//line /usr/local/go/src/strings/strings.go:1115
			// _ = "end of CoverTab[1694]"
//line /usr/local/go/src/strings/strings.go:1115
		}() {
//line /usr/local/go/src/strings/strings.go:1115
			_go_fuzz_dep_.CoverTab[1695]++
									continue
//line /usr/local/go/src/strings/strings.go:1116
			// _ = "end of CoverTab[1695]"
		} else {
//line /usr/local/go/src/strings/strings.go:1117
			_go_fuzz_dep_.CoverTab[1696]++
//line /usr/local/go/src/strings/strings.go:1117
			// _ = "end of CoverTab[1696]"
//line /usr/local/go/src/strings/strings.go:1117
		}
//line /usr/local/go/src/strings/strings.go:1117
		// _ = "end of CoverTab[1685]"
//line /usr/local/go/src/strings/strings.go:1117
		_go_fuzz_dep_.CoverTab[1686]++
								return false
//line /usr/local/go/src/strings/strings.go:1118
		// _ = "end of CoverTab[1686]"
	}
//line /usr/local/go/src/strings/strings.go:1119
	// _ = "end of CoverTab[1678]"
//line /usr/local/go/src/strings/strings.go:1119
	_go_fuzz_dep_.CoverTab[1679]++

							return len(s) == len(t)

hasUnicode:
	s = s[i:]
	t = t[i:]
	for _, sr := range s {
//line /usr/local/go/src/strings/strings.go:1126
		_go_fuzz_dep_.CoverTab[1697]++

								if len(t) == 0 {
//line /usr/local/go/src/strings/strings.go:1128
			_go_fuzz_dep_.CoverTab[1705]++
									return false
//line /usr/local/go/src/strings/strings.go:1129
			// _ = "end of CoverTab[1705]"
		} else {
//line /usr/local/go/src/strings/strings.go:1130
			_go_fuzz_dep_.CoverTab[1706]++
//line /usr/local/go/src/strings/strings.go:1130
			// _ = "end of CoverTab[1706]"
//line /usr/local/go/src/strings/strings.go:1130
		}
//line /usr/local/go/src/strings/strings.go:1130
		// _ = "end of CoverTab[1697]"
//line /usr/local/go/src/strings/strings.go:1130
		_go_fuzz_dep_.CoverTab[1698]++

		// Extract first rune from second string.
		var tr rune
		if t[0] < utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:1134
			_go_fuzz_dep_.CoverTab[1707]++
									tr, t = rune(t[0]), t[1:]
//line /usr/local/go/src/strings/strings.go:1135
			// _ = "end of CoverTab[1707]"
		} else {
//line /usr/local/go/src/strings/strings.go:1136
			_go_fuzz_dep_.CoverTab[1708]++
									r, size := utf8.DecodeRuneInString(t)
									tr, t = r, t[size:]
//line /usr/local/go/src/strings/strings.go:1138
			// _ = "end of CoverTab[1708]"
		}
//line /usr/local/go/src/strings/strings.go:1139
		// _ = "end of CoverTab[1698]"
//line /usr/local/go/src/strings/strings.go:1139
		_go_fuzz_dep_.CoverTab[1699]++

//line /usr/local/go/src/strings/strings.go:1144
		if tr == sr {
//line /usr/local/go/src/strings/strings.go:1144
			_go_fuzz_dep_.CoverTab[1709]++
									continue
//line /usr/local/go/src/strings/strings.go:1145
			// _ = "end of CoverTab[1709]"
		} else {
//line /usr/local/go/src/strings/strings.go:1146
			_go_fuzz_dep_.CoverTab[1710]++
//line /usr/local/go/src/strings/strings.go:1146
			// _ = "end of CoverTab[1710]"
//line /usr/local/go/src/strings/strings.go:1146
		}
//line /usr/local/go/src/strings/strings.go:1146
		// _ = "end of CoverTab[1699]"
//line /usr/local/go/src/strings/strings.go:1146
		_go_fuzz_dep_.CoverTab[1700]++

//line /usr/local/go/src/strings/strings.go:1149
		if tr < sr {
//line /usr/local/go/src/strings/strings.go:1149
			_go_fuzz_dep_.CoverTab[1711]++
									tr, sr = sr, tr
//line /usr/local/go/src/strings/strings.go:1150
			// _ = "end of CoverTab[1711]"
		} else {
//line /usr/local/go/src/strings/strings.go:1151
			_go_fuzz_dep_.CoverTab[1712]++
//line /usr/local/go/src/strings/strings.go:1151
			// _ = "end of CoverTab[1712]"
//line /usr/local/go/src/strings/strings.go:1151
		}
//line /usr/local/go/src/strings/strings.go:1151
		// _ = "end of CoverTab[1700]"
//line /usr/local/go/src/strings/strings.go:1151
		_go_fuzz_dep_.CoverTab[1701]++

								if tr < utf8.RuneSelf {
//line /usr/local/go/src/strings/strings.go:1153
			_go_fuzz_dep_.CoverTab[1713]++

									if 'A' <= sr && func() bool {
//line /usr/local/go/src/strings/strings.go:1155
				_go_fuzz_dep_.CoverTab[1715]++
//line /usr/local/go/src/strings/strings.go:1155
				return sr <= 'Z'
//line /usr/local/go/src/strings/strings.go:1155
				// _ = "end of CoverTab[1715]"
//line /usr/local/go/src/strings/strings.go:1155
			}() && func() bool {
//line /usr/local/go/src/strings/strings.go:1155
				_go_fuzz_dep_.CoverTab[1716]++
//line /usr/local/go/src/strings/strings.go:1155
				return tr == sr+'a'-'A'
//line /usr/local/go/src/strings/strings.go:1155
				// _ = "end of CoverTab[1716]"
//line /usr/local/go/src/strings/strings.go:1155
			}() {
//line /usr/local/go/src/strings/strings.go:1155
				_go_fuzz_dep_.CoverTab[1717]++
										continue
//line /usr/local/go/src/strings/strings.go:1156
				// _ = "end of CoverTab[1717]"
			} else {
//line /usr/local/go/src/strings/strings.go:1157
				_go_fuzz_dep_.CoverTab[1718]++
//line /usr/local/go/src/strings/strings.go:1157
				// _ = "end of CoverTab[1718]"
//line /usr/local/go/src/strings/strings.go:1157
			}
//line /usr/local/go/src/strings/strings.go:1157
			// _ = "end of CoverTab[1713]"
//line /usr/local/go/src/strings/strings.go:1157
			_go_fuzz_dep_.CoverTab[1714]++
									return false
//line /usr/local/go/src/strings/strings.go:1158
			// _ = "end of CoverTab[1714]"
		} else {
//line /usr/local/go/src/strings/strings.go:1159
			_go_fuzz_dep_.CoverTab[1719]++
//line /usr/local/go/src/strings/strings.go:1159
			// _ = "end of CoverTab[1719]"
//line /usr/local/go/src/strings/strings.go:1159
		}
//line /usr/local/go/src/strings/strings.go:1159
		// _ = "end of CoverTab[1701]"
//line /usr/local/go/src/strings/strings.go:1159
		_go_fuzz_dep_.CoverTab[1702]++

//line /usr/local/go/src/strings/strings.go:1163
		r := unicode.SimpleFold(sr)
		for r != sr && func() bool {
//line /usr/local/go/src/strings/strings.go:1164
			_go_fuzz_dep_.CoverTab[1720]++
//line /usr/local/go/src/strings/strings.go:1164
			return r < tr
//line /usr/local/go/src/strings/strings.go:1164
			// _ = "end of CoverTab[1720]"
//line /usr/local/go/src/strings/strings.go:1164
		}() {
//line /usr/local/go/src/strings/strings.go:1164
			_go_fuzz_dep_.CoverTab[1721]++
									r = unicode.SimpleFold(r)
//line /usr/local/go/src/strings/strings.go:1165
			// _ = "end of CoverTab[1721]"
		}
//line /usr/local/go/src/strings/strings.go:1166
		// _ = "end of CoverTab[1702]"
//line /usr/local/go/src/strings/strings.go:1166
		_go_fuzz_dep_.CoverTab[1703]++
								if r == tr {
//line /usr/local/go/src/strings/strings.go:1167
			_go_fuzz_dep_.CoverTab[1722]++
									continue
//line /usr/local/go/src/strings/strings.go:1168
			// _ = "end of CoverTab[1722]"
		} else {
//line /usr/local/go/src/strings/strings.go:1169
			_go_fuzz_dep_.CoverTab[1723]++
//line /usr/local/go/src/strings/strings.go:1169
			// _ = "end of CoverTab[1723]"
//line /usr/local/go/src/strings/strings.go:1169
		}
//line /usr/local/go/src/strings/strings.go:1169
		// _ = "end of CoverTab[1703]"
//line /usr/local/go/src/strings/strings.go:1169
		_go_fuzz_dep_.CoverTab[1704]++
								return false
//line /usr/local/go/src/strings/strings.go:1170
		// _ = "end of CoverTab[1704]"
	}
//line /usr/local/go/src/strings/strings.go:1171
	// _ = "end of CoverTab[1679]"
//line /usr/local/go/src/strings/strings.go:1171
	_go_fuzz_dep_.CoverTab[1680]++

//line /usr/local/go/src/strings/strings.go:1174
	return len(t) == 0
//line /usr/local/go/src/strings/strings.go:1174
	// _ = "end of CoverTab[1680]"
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
//line /usr/local/go/src/strings/strings.go:1178
	_go_fuzz_dep_.CoverTab[1724]++
							n := len(substr)
							switch {
	case n == 0:
//line /usr/local/go/src/strings/strings.go:1181
		_go_fuzz_dep_.CoverTab[1727]++
								return 0
//line /usr/local/go/src/strings/strings.go:1182
		// _ = "end of CoverTab[1727]"
	case n == 1:
//line /usr/local/go/src/strings/strings.go:1183
		_go_fuzz_dep_.CoverTab[1728]++
								return IndexByte(s, substr[0])
//line /usr/local/go/src/strings/strings.go:1184
		// _ = "end of CoverTab[1728]"
	case n == len(s):
//line /usr/local/go/src/strings/strings.go:1185
		_go_fuzz_dep_.CoverTab[1729]++
								if substr == s {
//line /usr/local/go/src/strings/strings.go:1186
			_go_fuzz_dep_.CoverTab[1736]++
									return 0
//line /usr/local/go/src/strings/strings.go:1187
			// _ = "end of CoverTab[1736]"
		} else {
//line /usr/local/go/src/strings/strings.go:1188
			_go_fuzz_dep_.CoverTab[1737]++
//line /usr/local/go/src/strings/strings.go:1188
			// _ = "end of CoverTab[1737]"
//line /usr/local/go/src/strings/strings.go:1188
		}
//line /usr/local/go/src/strings/strings.go:1188
		// _ = "end of CoverTab[1729]"
//line /usr/local/go/src/strings/strings.go:1188
		_go_fuzz_dep_.CoverTab[1730]++
								return -1
//line /usr/local/go/src/strings/strings.go:1189
		// _ = "end of CoverTab[1730]"
	case n > len(s):
//line /usr/local/go/src/strings/strings.go:1190
		_go_fuzz_dep_.CoverTab[1731]++
								return -1
//line /usr/local/go/src/strings/strings.go:1191
		// _ = "end of CoverTab[1731]"
	case n <= bytealg.MaxLen:
//line /usr/local/go/src/strings/strings.go:1192
		_go_fuzz_dep_.CoverTab[1732]++

								if len(s) <= bytealg.MaxBruteForce {
//line /usr/local/go/src/strings/strings.go:1194
			_go_fuzz_dep_.CoverTab[1738]++
									return bytealg.IndexString(s, substr)
//line /usr/local/go/src/strings/strings.go:1195
			// _ = "end of CoverTab[1738]"
		} else {
//line /usr/local/go/src/strings/strings.go:1196
			_go_fuzz_dep_.CoverTab[1739]++
//line /usr/local/go/src/strings/strings.go:1196
			// _ = "end of CoverTab[1739]"
//line /usr/local/go/src/strings/strings.go:1196
		}
//line /usr/local/go/src/strings/strings.go:1196
		// _ = "end of CoverTab[1732]"
//line /usr/local/go/src/strings/strings.go:1196
		_go_fuzz_dep_.CoverTab[1733]++
								c0 := substr[0]
								c1 := substr[1]
								i := 0
								t := len(s) - n + 1
								fails := 0
								for i < t {
//line /usr/local/go/src/strings/strings.go:1202
			_go_fuzz_dep_.CoverTab[1740]++
									if s[i] != c0 {
//line /usr/local/go/src/strings/strings.go:1203
				_go_fuzz_dep_.CoverTab[1743]++

//line /usr/local/go/src/strings/strings.go:1206
				o := IndexByte(s[i+1:t], c0)
				if o < 0 {
//line /usr/local/go/src/strings/strings.go:1207
					_go_fuzz_dep_.CoverTab[1745]++
											return -1
//line /usr/local/go/src/strings/strings.go:1208
					// _ = "end of CoverTab[1745]"
				} else {
//line /usr/local/go/src/strings/strings.go:1209
					_go_fuzz_dep_.CoverTab[1746]++
//line /usr/local/go/src/strings/strings.go:1209
					// _ = "end of CoverTab[1746]"
//line /usr/local/go/src/strings/strings.go:1209
				}
//line /usr/local/go/src/strings/strings.go:1209
				// _ = "end of CoverTab[1743]"
//line /usr/local/go/src/strings/strings.go:1209
				_go_fuzz_dep_.CoverTab[1744]++
										i += o + 1
//line /usr/local/go/src/strings/strings.go:1210
				// _ = "end of CoverTab[1744]"
			} else {
//line /usr/local/go/src/strings/strings.go:1211
				_go_fuzz_dep_.CoverTab[1747]++
//line /usr/local/go/src/strings/strings.go:1211
				// _ = "end of CoverTab[1747]"
//line /usr/local/go/src/strings/strings.go:1211
			}
//line /usr/local/go/src/strings/strings.go:1211
			// _ = "end of CoverTab[1740]"
//line /usr/local/go/src/strings/strings.go:1211
			_go_fuzz_dep_.CoverTab[1741]++
									if s[i+1] == c1 && func() bool {
//line /usr/local/go/src/strings/strings.go:1212
				_go_fuzz_dep_.CoverTab[1748]++
//line /usr/local/go/src/strings/strings.go:1212
				return s[i:i+n] == substr
//line /usr/local/go/src/strings/strings.go:1212
				// _ = "end of CoverTab[1748]"
//line /usr/local/go/src/strings/strings.go:1212
			}() {
//line /usr/local/go/src/strings/strings.go:1212
				_go_fuzz_dep_.CoverTab[1749]++
										return i
//line /usr/local/go/src/strings/strings.go:1213
				// _ = "end of CoverTab[1749]"
			} else {
//line /usr/local/go/src/strings/strings.go:1214
				_go_fuzz_dep_.CoverTab[1750]++
//line /usr/local/go/src/strings/strings.go:1214
				// _ = "end of CoverTab[1750]"
//line /usr/local/go/src/strings/strings.go:1214
			}
//line /usr/local/go/src/strings/strings.go:1214
			// _ = "end of CoverTab[1741]"
//line /usr/local/go/src/strings/strings.go:1214
			_go_fuzz_dep_.CoverTab[1742]++
									fails++
									i++

									if fails > bytealg.Cutover(i) {
//line /usr/local/go/src/strings/strings.go:1218
				_go_fuzz_dep_.CoverTab[1751]++
										r := bytealg.IndexString(s[i:], substr)
										if r >= 0 {
//line /usr/local/go/src/strings/strings.go:1220
					_go_fuzz_dep_.CoverTab[1753]++
											return r + i
//line /usr/local/go/src/strings/strings.go:1221
					// _ = "end of CoverTab[1753]"
				} else {
//line /usr/local/go/src/strings/strings.go:1222
					_go_fuzz_dep_.CoverTab[1754]++
//line /usr/local/go/src/strings/strings.go:1222
					// _ = "end of CoverTab[1754]"
//line /usr/local/go/src/strings/strings.go:1222
				}
//line /usr/local/go/src/strings/strings.go:1222
				// _ = "end of CoverTab[1751]"
//line /usr/local/go/src/strings/strings.go:1222
				_go_fuzz_dep_.CoverTab[1752]++
										return -1
//line /usr/local/go/src/strings/strings.go:1223
				// _ = "end of CoverTab[1752]"
			} else {
//line /usr/local/go/src/strings/strings.go:1224
				_go_fuzz_dep_.CoverTab[1755]++
//line /usr/local/go/src/strings/strings.go:1224
				// _ = "end of CoverTab[1755]"
//line /usr/local/go/src/strings/strings.go:1224
			}
//line /usr/local/go/src/strings/strings.go:1224
			// _ = "end of CoverTab[1742]"
		}
//line /usr/local/go/src/strings/strings.go:1225
		// _ = "end of CoverTab[1733]"
//line /usr/local/go/src/strings/strings.go:1225
		_go_fuzz_dep_.CoverTab[1734]++
								return -1
//line /usr/local/go/src/strings/strings.go:1226
		// _ = "end of CoverTab[1734]"
//line /usr/local/go/src/strings/strings.go:1226
	default:
//line /usr/local/go/src/strings/strings.go:1226
		_go_fuzz_dep_.CoverTab[1735]++
//line /usr/local/go/src/strings/strings.go:1226
		// _ = "end of CoverTab[1735]"
	}
//line /usr/local/go/src/strings/strings.go:1227
	// _ = "end of CoverTab[1724]"
//line /usr/local/go/src/strings/strings.go:1227
	_go_fuzz_dep_.CoverTab[1725]++
							c0 := substr[0]
							c1 := substr[1]
							i := 0
							t := len(s) - n + 1
							fails := 0
							for i < t {
//line /usr/local/go/src/strings/strings.go:1233
		_go_fuzz_dep_.CoverTab[1756]++
								if s[i] != c0 {
//line /usr/local/go/src/strings/strings.go:1234
			_go_fuzz_dep_.CoverTab[1759]++
									o := IndexByte(s[i+1:t], c0)
									if o < 0 {
//line /usr/local/go/src/strings/strings.go:1236
				_go_fuzz_dep_.CoverTab[1761]++
										return -1
//line /usr/local/go/src/strings/strings.go:1237
				// _ = "end of CoverTab[1761]"
			} else {
//line /usr/local/go/src/strings/strings.go:1238
				_go_fuzz_dep_.CoverTab[1762]++
//line /usr/local/go/src/strings/strings.go:1238
				// _ = "end of CoverTab[1762]"
//line /usr/local/go/src/strings/strings.go:1238
			}
//line /usr/local/go/src/strings/strings.go:1238
			// _ = "end of CoverTab[1759]"
//line /usr/local/go/src/strings/strings.go:1238
			_go_fuzz_dep_.CoverTab[1760]++
									i += o + 1
//line /usr/local/go/src/strings/strings.go:1239
			// _ = "end of CoverTab[1760]"
		} else {
//line /usr/local/go/src/strings/strings.go:1240
			_go_fuzz_dep_.CoverTab[1763]++
//line /usr/local/go/src/strings/strings.go:1240
			// _ = "end of CoverTab[1763]"
//line /usr/local/go/src/strings/strings.go:1240
		}
//line /usr/local/go/src/strings/strings.go:1240
		// _ = "end of CoverTab[1756]"
//line /usr/local/go/src/strings/strings.go:1240
		_go_fuzz_dep_.CoverTab[1757]++
								if s[i+1] == c1 && func() bool {
//line /usr/local/go/src/strings/strings.go:1241
			_go_fuzz_dep_.CoverTab[1764]++
//line /usr/local/go/src/strings/strings.go:1241
			return s[i:i+n] == substr
//line /usr/local/go/src/strings/strings.go:1241
			// _ = "end of CoverTab[1764]"
//line /usr/local/go/src/strings/strings.go:1241
		}() {
//line /usr/local/go/src/strings/strings.go:1241
			_go_fuzz_dep_.CoverTab[1765]++
									return i
//line /usr/local/go/src/strings/strings.go:1242
			// _ = "end of CoverTab[1765]"
		} else {
//line /usr/local/go/src/strings/strings.go:1243
			_go_fuzz_dep_.CoverTab[1766]++
//line /usr/local/go/src/strings/strings.go:1243
			// _ = "end of CoverTab[1766]"
//line /usr/local/go/src/strings/strings.go:1243
		}
//line /usr/local/go/src/strings/strings.go:1243
		// _ = "end of CoverTab[1757]"
//line /usr/local/go/src/strings/strings.go:1243
		_go_fuzz_dep_.CoverTab[1758]++
								i++
								fails++
								if fails >= 4+i>>4 && func() bool {
//line /usr/local/go/src/strings/strings.go:1246
			_go_fuzz_dep_.CoverTab[1767]++
//line /usr/local/go/src/strings/strings.go:1246
			return i < t
//line /usr/local/go/src/strings/strings.go:1246
			// _ = "end of CoverTab[1767]"
//line /usr/local/go/src/strings/strings.go:1246
		}() {
//line /usr/local/go/src/strings/strings.go:1246
			_go_fuzz_dep_.CoverTab[1768]++

									j := bytealg.IndexRabinKarp(s[i:], substr)
									if j < 0 {
//line /usr/local/go/src/strings/strings.go:1249
				_go_fuzz_dep_.CoverTab[1770]++
										return -1
//line /usr/local/go/src/strings/strings.go:1250
				// _ = "end of CoverTab[1770]"
			} else {
//line /usr/local/go/src/strings/strings.go:1251
				_go_fuzz_dep_.CoverTab[1771]++
//line /usr/local/go/src/strings/strings.go:1251
				// _ = "end of CoverTab[1771]"
//line /usr/local/go/src/strings/strings.go:1251
			}
//line /usr/local/go/src/strings/strings.go:1251
			// _ = "end of CoverTab[1768]"
//line /usr/local/go/src/strings/strings.go:1251
			_go_fuzz_dep_.CoverTab[1769]++
									return i + j
//line /usr/local/go/src/strings/strings.go:1252
			// _ = "end of CoverTab[1769]"
		} else {
//line /usr/local/go/src/strings/strings.go:1253
			_go_fuzz_dep_.CoverTab[1772]++
//line /usr/local/go/src/strings/strings.go:1253
			// _ = "end of CoverTab[1772]"
//line /usr/local/go/src/strings/strings.go:1253
		}
//line /usr/local/go/src/strings/strings.go:1253
		// _ = "end of CoverTab[1758]"
	}
//line /usr/local/go/src/strings/strings.go:1254
	// _ = "end of CoverTab[1725]"
//line /usr/local/go/src/strings/strings.go:1254
	_go_fuzz_dep_.CoverTab[1726]++
							return -1
//line /usr/local/go/src/strings/strings.go:1255
	// _ = "end of CoverTab[1726]"
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
	_go_fuzz_dep_.CoverTab[1773]++
							if i := Index(s, sep); i >= 0 {
//line /usr/local/go/src/strings/strings.go:1263
		_go_fuzz_dep_.CoverTab[1775]++
								return s[:i], s[i+len(sep):], true
//line /usr/local/go/src/strings/strings.go:1264
		// _ = "end of CoverTab[1775]"
	} else {
//line /usr/local/go/src/strings/strings.go:1265
		_go_fuzz_dep_.CoverTab[1776]++
//line /usr/local/go/src/strings/strings.go:1265
		// _ = "end of CoverTab[1776]"
//line /usr/local/go/src/strings/strings.go:1265
	}
//line /usr/local/go/src/strings/strings.go:1265
	// _ = "end of CoverTab[1773]"
//line /usr/local/go/src/strings/strings.go:1265
	_go_fuzz_dep_.CoverTab[1774]++
							return s, "", false
//line /usr/local/go/src/strings/strings.go:1266
	// _ = "end of CoverTab[1774]"
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
	_go_fuzz_dep_.CoverTab[1777]++
							if !HasPrefix(s, prefix) {
//line /usr/local/go/src/strings/strings.go:1274
		_go_fuzz_dep_.CoverTab[1779]++
								return s, false
//line /usr/local/go/src/strings/strings.go:1275
		// _ = "end of CoverTab[1779]"
	} else {
//line /usr/local/go/src/strings/strings.go:1276
		_go_fuzz_dep_.CoverTab[1780]++
//line /usr/local/go/src/strings/strings.go:1276
		// _ = "end of CoverTab[1780]"
//line /usr/local/go/src/strings/strings.go:1276
	}
//line /usr/local/go/src/strings/strings.go:1276
	// _ = "end of CoverTab[1777]"
//line /usr/local/go/src/strings/strings.go:1276
	_go_fuzz_dep_.CoverTab[1778]++
							return s[len(prefix):], true
//line /usr/local/go/src/strings/strings.go:1277
	// _ = "end of CoverTab[1778]"
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
	_go_fuzz_dep_.CoverTab[1781]++
							if !HasSuffix(s, suffix) {
//line /usr/local/go/src/strings/strings.go:1285
		_go_fuzz_dep_.CoverTab[1783]++
								return s, false
//line /usr/local/go/src/strings/strings.go:1286
		// _ = "end of CoverTab[1783]"
	} else {
//line /usr/local/go/src/strings/strings.go:1287
		_go_fuzz_dep_.CoverTab[1784]++
//line /usr/local/go/src/strings/strings.go:1287
		// _ = "end of CoverTab[1784]"
//line /usr/local/go/src/strings/strings.go:1287
	}
//line /usr/local/go/src/strings/strings.go:1287
	// _ = "end of CoverTab[1781]"
//line /usr/local/go/src/strings/strings.go:1287
	_go_fuzz_dep_.CoverTab[1782]++
							return s[:len(s)-len(suffix)], true
//line /usr/local/go/src/strings/strings.go:1288
	// _ = "end of CoverTab[1782]"
}

//line /usr/local/go/src/strings/strings.go:1289
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/strings.go:1289
var _ = _go_fuzz_dep_.CoverTab
