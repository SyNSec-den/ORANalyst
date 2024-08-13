// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/bytes/bytes.go:5
// Package bytes implements functions for the manipulation of byte slices.
//line /usr/local/go/src/bytes/bytes.go:5
// It is analogous to the facilities of the strings package.
//line /usr/local/go/src/bytes/bytes.go:7
package bytes

//line /usr/local/go/src/bytes/bytes.go:7
import (
//line /usr/local/go/src/bytes/bytes.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/bytes/bytes.go:7
)
//line /usr/local/go/src/bytes/bytes.go:7
import (
//line /usr/local/go/src/bytes/bytes.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/bytes/bytes.go:7
)

import (
	"internal/bytealg"
	"unicode"
	"unicode/utf8"
)

// Equal reports whether a and b
//line /usr/local/go/src/bytes/bytes.go:15
// are the same length and contain the same bytes.
//line /usr/local/go/src/bytes/bytes.go:15
// A nil argument is equivalent to an empty slice.
//line /usr/local/go/src/bytes/bytes.go:18
func Equal(a, b []byte) bool {
//line /usr/local/go/src/bytes/bytes.go:18
	_go_fuzz_dep_.CoverTab[142]++

						return string(a) == string(b)
//line /usr/local/go/src/bytes/bytes.go:20
	// _ = "end of CoverTab[142]"
}

// Compare returns an integer comparing two byte slices lexicographically.
//line /usr/local/go/src/bytes/bytes.go:23
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//line /usr/local/go/src/bytes/bytes.go:23
// A nil argument is equivalent to an empty slice.
//line /usr/local/go/src/bytes/bytes.go:26
func Compare(a, b []byte) int {
//line /usr/local/go/src/bytes/bytes.go:26
	_go_fuzz_dep_.CoverTab[143]++
						return bytealg.Compare(a, b)
//line /usr/local/go/src/bytes/bytes.go:27
	// _ = "end of CoverTab[143]"
}

// explode splits s into a slice of UTF-8 sequences, one per Unicode code point (still slices of bytes),
//line /usr/local/go/src/bytes/bytes.go:30
// up to a maximum of n byte slices. Invalid UTF-8 sequences are chopped into individual bytes.
//line /usr/local/go/src/bytes/bytes.go:32
func explode(s []byte, n int) [][]byte {
//line /usr/local/go/src/bytes/bytes.go:32
	_go_fuzz_dep_.CoverTab[144]++
						if n <= 0 || func() bool {
//line /usr/local/go/src/bytes/bytes.go:33
		_go_fuzz_dep_.CoverTab[147]++
//line /usr/local/go/src/bytes/bytes.go:33
		return n > len(s)
//line /usr/local/go/src/bytes/bytes.go:33
		// _ = "end of CoverTab[147]"
//line /usr/local/go/src/bytes/bytes.go:33
	}() {
//line /usr/local/go/src/bytes/bytes.go:33
		_go_fuzz_dep_.CoverTab[148]++
							n = len(s)
//line /usr/local/go/src/bytes/bytes.go:34
		// _ = "end of CoverTab[148]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:35
		_go_fuzz_dep_.CoverTab[149]++
//line /usr/local/go/src/bytes/bytes.go:35
		// _ = "end of CoverTab[149]"
//line /usr/local/go/src/bytes/bytes.go:35
	}
//line /usr/local/go/src/bytes/bytes.go:35
	// _ = "end of CoverTab[144]"
//line /usr/local/go/src/bytes/bytes.go:35
	_go_fuzz_dep_.CoverTab[145]++
						a := make([][]byte, n)
						var size int
						na := 0
						for len(s) > 0 {
//line /usr/local/go/src/bytes/bytes.go:39
		_go_fuzz_dep_.CoverTab[150]++
							if na+1 >= n {
//line /usr/local/go/src/bytes/bytes.go:40
			_go_fuzz_dep_.CoverTab[152]++
								a[na] = s
								na++
								break
//line /usr/local/go/src/bytes/bytes.go:43
			// _ = "end of CoverTab[152]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:44
			_go_fuzz_dep_.CoverTab[153]++
//line /usr/local/go/src/bytes/bytes.go:44
			// _ = "end of CoverTab[153]"
//line /usr/local/go/src/bytes/bytes.go:44
		}
//line /usr/local/go/src/bytes/bytes.go:44
		// _ = "end of CoverTab[150]"
//line /usr/local/go/src/bytes/bytes.go:44
		_go_fuzz_dep_.CoverTab[151]++
							_, size = utf8.DecodeRune(s)
							a[na] = s[0:size:size]
							s = s[size:]
							na++
//line /usr/local/go/src/bytes/bytes.go:48
		// _ = "end of CoverTab[151]"
	}
//line /usr/local/go/src/bytes/bytes.go:49
	// _ = "end of CoverTab[145]"
//line /usr/local/go/src/bytes/bytes.go:49
	_go_fuzz_dep_.CoverTab[146]++
						return a[0:na]
//line /usr/local/go/src/bytes/bytes.go:50
	// _ = "end of CoverTab[146]"
}

// Count counts the number of non-overlapping instances of sep in s.
//line /usr/local/go/src/bytes/bytes.go:53
// If sep is an empty slice, Count returns 1 + the number of UTF-8-encoded code points in s.
//line /usr/local/go/src/bytes/bytes.go:55
func Count(s, sep []byte) int {
//line /usr/local/go/src/bytes/bytes.go:55
	_go_fuzz_dep_.CoverTab[154]++

						if len(sep) == 0 {
//line /usr/local/go/src/bytes/bytes.go:57
		_go_fuzz_dep_.CoverTab[157]++
							return utf8.RuneCount(s) + 1
//line /usr/local/go/src/bytes/bytes.go:58
		// _ = "end of CoverTab[157]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:59
		_go_fuzz_dep_.CoverTab[158]++
//line /usr/local/go/src/bytes/bytes.go:59
		// _ = "end of CoverTab[158]"
//line /usr/local/go/src/bytes/bytes.go:59
	}
//line /usr/local/go/src/bytes/bytes.go:59
	// _ = "end of CoverTab[154]"
//line /usr/local/go/src/bytes/bytes.go:59
	_go_fuzz_dep_.CoverTab[155]++
						if len(sep) == 1 {
//line /usr/local/go/src/bytes/bytes.go:60
		_go_fuzz_dep_.CoverTab[159]++
							return bytealg.Count(s, sep[0])
//line /usr/local/go/src/bytes/bytes.go:61
		// _ = "end of CoverTab[159]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:62
		_go_fuzz_dep_.CoverTab[160]++
//line /usr/local/go/src/bytes/bytes.go:62
		// _ = "end of CoverTab[160]"
//line /usr/local/go/src/bytes/bytes.go:62
	}
//line /usr/local/go/src/bytes/bytes.go:62
	// _ = "end of CoverTab[155]"
//line /usr/local/go/src/bytes/bytes.go:62
	_go_fuzz_dep_.CoverTab[156]++
						n := 0
						for {
//line /usr/local/go/src/bytes/bytes.go:64
		_go_fuzz_dep_.CoverTab[161]++
							i := Index(s, sep)
							if i == -1 {
//line /usr/local/go/src/bytes/bytes.go:66
			_go_fuzz_dep_.CoverTab[163]++
								return n
//line /usr/local/go/src/bytes/bytes.go:67
			// _ = "end of CoverTab[163]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:68
			_go_fuzz_dep_.CoverTab[164]++
//line /usr/local/go/src/bytes/bytes.go:68
			// _ = "end of CoverTab[164]"
//line /usr/local/go/src/bytes/bytes.go:68
		}
//line /usr/local/go/src/bytes/bytes.go:68
		// _ = "end of CoverTab[161]"
//line /usr/local/go/src/bytes/bytes.go:68
		_go_fuzz_dep_.CoverTab[162]++
							n++
							s = s[i+len(sep):]
//line /usr/local/go/src/bytes/bytes.go:70
		// _ = "end of CoverTab[162]"
	}
//line /usr/local/go/src/bytes/bytes.go:71
	// _ = "end of CoverTab[156]"
}

// Contains reports whether subslice is within b.
func Contains(b, subslice []byte) bool {
//line /usr/local/go/src/bytes/bytes.go:75
	_go_fuzz_dep_.CoverTab[165]++
						return Index(b, subslice) != -1
//line /usr/local/go/src/bytes/bytes.go:76
	// _ = "end of CoverTab[165]"
}

// ContainsAny reports whether any of the UTF-8-encoded code points in chars are within b.
func ContainsAny(b []byte, chars string) bool {
//line /usr/local/go/src/bytes/bytes.go:80
	_go_fuzz_dep_.CoverTab[166]++
						return IndexAny(b, chars) >= 0
//line /usr/local/go/src/bytes/bytes.go:81
	// _ = "end of CoverTab[166]"
}

// ContainsRune reports whether the rune is contained in the UTF-8-encoded byte slice b.
func ContainsRune(b []byte, r rune) bool {
//line /usr/local/go/src/bytes/bytes.go:85
	_go_fuzz_dep_.CoverTab[167]++
						return IndexRune(b, r) >= 0
//line /usr/local/go/src/bytes/bytes.go:86
	// _ = "end of CoverTab[167]"
}

// IndexByte returns the index of the first instance of c in b, or -1 if c is not present in b.
func IndexByte(b []byte, c byte) int {
//line /usr/local/go/src/bytes/bytes.go:90
	_go_fuzz_dep_.CoverTab[168]++
						return bytealg.IndexByte(b, c)
//line /usr/local/go/src/bytes/bytes.go:91
	// _ = "end of CoverTab[168]"
}

func indexBytePortable(s []byte, c byte) int {
//line /usr/local/go/src/bytes/bytes.go:94
	_go_fuzz_dep_.CoverTab[169]++
						for i, b := range s {
//line /usr/local/go/src/bytes/bytes.go:95
		_go_fuzz_dep_.CoverTab[171]++
							if b == c {
//line /usr/local/go/src/bytes/bytes.go:96
			_go_fuzz_dep_.CoverTab[172]++
								return i
//line /usr/local/go/src/bytes/bytes.go:97
			// _ = "end of CoverTab[172]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:98
			_go_fuzz_dep_.CoverTab[173]++
//line /usr/local/go/src/bytes/bytes.go:98
			// _ = "end of CoverTab[173]"
//line /usr/local/go/src/bytes/bytes.go:98
		}
//line /usr/local/go/src/bytes/bytes.go:98
		// _ = "end of CoverTab[171]"
	}
//line /usr/local/go/src/bytes/bytes.go:99
	// _ = "end of CoverTab[169]"
//line /usr/local/go/src/bytes/bytes.go:99
	_go_fuzz_dep_.CoverTab[170]++
						return -1
//line /usr/local/go/src/bytes/bytes.go:100
	// _ = "end of CoverTab[170]"
}

// LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.
func LastIndex(s, sep []byte) int {
//line /usr/local/go/src/bytes/bytes.go:104
	_go_fuzz_dep_.CoverTab[174]++
						n := len(sep)
						switch {
	case n == 0:
//line /usr/local/go/src/bytes/bytes.go:107
		_go_fuzz_dep_.CoverTab[179]++
							return len(s)
//line /usr/local/go/src/bytes/bytes.go:108
		// _ = "end of CoverTab[179]"
	case n == 1:
//line /usr/local/go/src/bytes/bytes.go:109
		_go_fuzz_dep_.CoverTab[180]++
							return LastIndexByte(s, sep[0])
//line /usr/local/go/src/bytes/bytes.go:110
		// _ = "end of CoverTab[180]"
	case n == len(s):
//line /usr/local/go/src/bytes/bytes.go:111
		_go_fuzz_dep_.CoverTab[181]++
							if Equal(s, sep) {
//line /usr/local/go/src/bytes/bytes.go:112
			_go_fuzz_dep_.CoverTab[185]++
								return 0
//line /usr/local/go/src/bytes/bytes.go:113
			// _ = "end of CoverTab[185]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:114
			_go_fuzz_dep_.CoverTab[186]++
//line /usr/local/go/src/bytes/bytes.go:114
			// _ = "end of CoverTab[186]"
//line /usr/local/go/src/bytes/bytes.go:114
		}
//line /usr/local/go/src/bytes/bytes.go:114
		// _ = "end of CoverTab[181]"
//line /usr/local/go/src/bytes/bytes.go:114
		_go_fuzz_dep_.CoverTab[182]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:115
		// _ = "end of CoverTab[182]"
	case n > len(s):
//line /usr/local/go/src/bytes/bytes.go:116
		_go_fuzz_dep_.CoverTab[183]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:117
		// _ = "end of CoverTab[183]"
//line /usr/local/go/src/bytes/bytes.go:117
	default:
//line /usr/local/go/src/bytes/bytes.go:117
		_go_fuzz_dep_.CoverTab[184]++
//line /usr/local/go/src/bytes/bytes.go:117
		// _ = "end of CoverTab[184]"
	}
//line /usr/local/go/src/bytes/bytes.go:118
	// _ = "end of CoverTab[174]"
//line /usr/local/go/src/bytes/bytes.go:118
	_go_fuzz_dep_.CoverTab[175]++

						hashss, pow := bytealg.HashStrRevBytes(sep)
						last := len(s) - n
						var h uint32
						for i := len(s) - 1; i >= last; i-- {
//line /usr/local/go/src/bytes/bytes.go:123
		_go_fuzz_dep_.CoverTab[187]++
							h = h*bytealg.PrimeRK + uint32(s[i])
//line /usr/local/go/src/bytes/bytes.go:124
		// _ = "end of CoverTab[187]"
	}
//line /usr/local/go/src/bytes/bytes.go:125
	// _ = "end of CoverTab[175]"
//line /usr/local/go/src/bytes/bytes.go:125
	_go_fuzz_dep_.CoverTab[176]++
						if h == hashss && func() bool {
//line /usr/local/go/src/bytes/bytes.go:126
		_go_fuzz_dep_.CoverTab[188]++
//line /usr/local/go/src/bytes/bytes.go:126
		return Equal(s[last:], sep)
//line /usr/local/go/src/bytes/bytes.go:126
		// _ = "end of CoverTab[188]"
//line /usr/local/go/src/bytes/bytes.go:126
	}() {
//line /usr/local/go/src/bytes/bytes.go:126
		_go_fuzz_dep_.CoverTab[189]++
							return last
//line /usr/local/go/src/bytes/bytes.go:127
		// _ = "end of CoverTab[189]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:128
		_go_fuzz_dep_.CoverTab[190]++
//line /usr/local/go/src/bytes/bytes.go:128
		// _ = "end of CoverTab[190]"
//line /usr/local/go/src/bytes/bytes.go:128
	}
//line /usr/local/go/src/bytes/bytes.go:128
	// _ = "end of CoverTab[176]"
//line /usr/local/go/src/bytes/bytes.go:128
	_go_fuzz_dep_.CoverTab[177]++
						for i := last - 1; i >= 0; i-- {
//line /usr/local/go/src/bytes/bytes.go:129
		_go_fuzz_dep_.CoverTab[191]++
							h *= bytealg.PrimeRK
							h += uint32(s[i])
							h -= pow * uint32(s[i+n])
							if h == hashss && func() bool {
//line /usr/local/go/src/bytes/bytes.go:133
			_go_fuzz_dep_.CoverTab[192]++
//line /usr/local/go/src/bytes/bytes.go:133
			return Equal(s[i:i+n], sep)
//line /usr/local/go/src/bytes/bytes.go:133
			// _ = "end of CoverTab[192]"
//line /usr/local/go/src/bytes/bytes.go:133
		}() {
//line /usr/local/go/src/bytes/bytes.go:133
			_go_fuzz_dep_.CoverTab[193]++
								return i
//line /usr/local/go/src/bytes/bytes.go:134
			// _ = "end of CoverTab[193]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:135
			_go_fuzz_dep_.CoverTab[194]++
//line /usr/local/go/src/bytes/bytes.go:135
			// _ = "end of CoverTab[194]"
//line /usr/local/go/src/bytes/bytes.go:135
		}
//line /usr/local/go/src/bytes/bytes.go:135
		// _ = "end of CoverTab[191]"
	}
//line /usr/local/go/src/bytes/bytes.go:136
	// _ = "end of CoverTab[177]"
//line /usr/local/go/src/bytes/bytes.go:136
	_go_fuzz_dep_.CoverTab[178]++
						return -1
//line /usr/local/go/src/bytes/bytes.go:137
	// _ = "end of CoverTab[178]"
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte(s []byte, c byte) int {
//line /usr/local/go/src/bytes/bytes.go:141
	_go_fuzz_dep_.CoverTab[195]++
						for i := len(s) - 1; i >= 0; i-- {
//line /usr/local/go/src/bytes/bytes.go:142
		_go_fuzz_dep_.CoverTab[197]++
							if s[i] == c {
//line /usr/local/go/src/bytes/bytes.go:143
			_go_fuzz_dep_.CoverTab[198]++
								return i
//line /usr/local/go/src/bytes/bytes.go:144
			// _ = "end of CoverTab[198]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:145
			_go_fuzz_dep_.CoverTab[199]++
//line /usr/local/go/src/bytes/bytes.go:145
			// _ = "end of CoverTab[199]"
//line /usr/local/go/src/bytes/bytes.go:145
		}
//line /usr/local/go/src/bytes/bytes.go:145
		// _ = "end of CoverTab[197]"
	}
//line /usr/local/go/src/bytes/bytes.go:146
	// _ = "end of CoverTab[195]"
//line /usr/local/go/src/bytes/bytes.go:146
	_go_fuzz_dep_.CoverTab[196]++
						return -1
//line /usr/local/go/src/bytes/bytes.go:147
	// _ = "end of CoverTab[196]"
}

// IndexRune interprets s as a sequence of UTF-8-encoded code points.
//line /usr/local/go/src/bytes/bytes.go:150
// It returns the byte index of the first occurrence in s of the given rune.
//line /usr/local/go/src/bytes/bytes.go:150
// It returns -1 if rune is not present in s.
//line /usr/local/go/src/bytes/bytes.go:150
// If r is utf8.RuneError, it returns the first instance of any
//line /usr/local/go/src/bytes/bytes.go:150
// invalid UTF-8 byte sequence.
//line /usr/local/go/src/bytes/bytes.go:155
func IndexRune(s []byte, r rune) int {
//line /usr/local/go/src/bytes/bytes.go:155
	_go_fuzz_dep_.CoverTab[200]++
						switch {
	case 0 <= r && func() bool {
//line /usr/local/go/src/bytes/bytes.go:157
		_go_fuzz_dep_.CoverTab[206]++
//line /usr/local/go/src/bytes/bytes.go:157
		return r < utf8.RuneSelf
//line /usr/local/go/src/bytes/bytes.go:157
		// _ = "end of CoverTab[206]"
//line /usr/local/go/src/bytes/bytes.go:157
	}():
//line /usr/local/go/src/bytes/bytes.go:157
		_go_fuzz_dep_.CoverTab[201]++
							return IndexByte(s, byte(r))
//line /usr/local/go/src/bytes/bytes.go:158
		// _ = "end of CoverTab[201]"
	case r == utf8.RuneError:
//line /usr/local/go/src/bytes/bytes.go:159
		_go_fuzz_dep_.CoverTab[202]++
							for i := 0; i < len(s); {
//line /usr/local/go/src/bytes/bytes.go:160
			_go_fuzz_dep_.CoverTab[207]++
								r1, n := utf8.DecodeRune(s[i:])
								if r1 == utf8.RuneError {
//line /usr/local/go/src/bytes/bytes.go:162
				_go_fuzz_dep_.CoverTab[209]++
									return i
//line /usr/local/go/src/bytes/bytes.go:163
				// _ = "end of CoverTab[209]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:164
				_go_fuzz_dep_.CoverTab[210]++
//line /usr/local/go/src/bytes/bytes.go:164
				// _ = "end of CoverTab[210]"
//line /usr/local/go/src/bytes/bytes.go:164
			}
//line /usr/local/go/src/bytes/bytes.go:164
			// _ = "end of CoverTab[207]"
//line /usr/local/go/src/bytes/bytes.go:164
			_go_fuzz_dep_.CoverTab[208]++
								i += n
//line /usr/local/go/src/bytes/bytes.go:165
			// _ = "end of CoverTab[208]"
		}
//line /usr/local/go/src/bytes/bytes.go:166
		// _ = "end of CoverTab[202]"
//line /usr/local/go/src/bytes/bytes.go:166
		_go_fuzz_dep_.CoverTab[203]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:167
		// _ = "end of CoverTab[203]"
	case !utf8.ValidRune(r):
//line /usr/local/go/src/bytes/bytes.go:168
		_go_fuzz_dep_.CoverTab[204]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:169
		// _ = "end of CoverTab[204]"
	default:
//line /usr/local/go/src/bytes/bytes.go:170
		_go_fuzz_dep_.CoverTab[205]++
							var b [utf8.UTFMax]byte
							n := utf8.EncodeRune(b[:], r)
							return Index(s, b[:n])
//line /usr/local/go/src/bytes/bytes.go:173
		// _ = "end of CoverTab[205]"
	}
//line /usr/local/go/src/bytes/bytes.go:174
	// _ = "end of CoverTab[200]"
}

// IndexAny interprets s as a sequence of UTF-8-encoded Unicode code points.
//line /usr/local/go/src/bytes/bytes.go:177
// It returns the byte index of the first occurrence in s of any of the Unicode
//line /usr/local/go/src/bytes/bytes.go:177
// code points in chars. It returns -1 if chars is empty or if there is no code
//line /usr/local/go/src/bytes/bytes.go:177
// point in common.
//line /usr/local/go/src/bytes/bytes.go:181
func IndexAny(s []byte, chars string) int {
//line /usr/local/go/src/bytes/bytes.go:181
	_go_fuzz_dep_.CoverTab[211]++
						if chars == "" {
//line /usr/local/go/src/bytes/bytes.go:182
		_go_fuzz_dep_.CoverTab[217]++

							return -1
//line /usr/local/go/src/bytes/bytes.go:184
		// _ = "end of CoverTab[217]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:185
		_go_fuzz_dep_.CoverTab[218]++
//line /usr/local/go/src/bytes/bytes.go:185
		// _ = "end of CoverTab[218]"
//line /usr/local/go/src/bytes/bytes.go:185
	}
//line /usr/local/go/src/bytes/bytes.go:185
	// _ = "end of CoverTab[211]"
//line /usr/local/go/src/bytes/bytes.go:185
	_go_fuzz_dep_.CoverTab[212]++
						if len(s) == 1 {
//line /usr/local/go/src/bytes/bytes.go:186
		_go_fuzz_dep_.CoverTab[219]++
							r := rune(s[0])
							if r >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:188
			_go_fuzz_dep_.CoverTab[222]++

								for _, r = range chars {
//line /usr/local/go/src/bytes/bytes.go:190
				_go_fuzz_dep_.CoverTab[224]++
									if r == utf8.RuneError {
//line /usr/local/go/src/bytes/bytes.go:191
					_go_fuzz_dep_.CoverTab[225]++
										return 0
//line /usr/local/go/src/bytes/bytes.go:192
					// _ = "end of CoverTab[225]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:193
					_go_fuzz_dep_.CoverTab[226]++
//line /usr/local/go/src/bytes/bytes.go:193
					// _ = "end of CoverTab[226]"
//line /usr/local/go/src/bytes/bytes.go:193
				}
//line /usr/local/go/src/bytes/bytes.go:193
				// _ = "end of CoverTab[224]"
			}
//line /usr/local/go/src/bytes/bytes.go:194
			// _ = "end of CoverTab[222]"
//line /usr/local/go/src/bytes/bytes.go:194
			_go_fuzz_dep_.CoverTab[223]++
								return -1
//line /usr/local/go/src/bytes/bytes.go:195
			// _ = "end of CoverTab[223]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:196
			_go_fuzz_dep_.CoverTab[227]++
//line /usr/local/go/src/bytes/bytes.go:196
			// _ = "end of CoverTab[227]"
//line /usr/local/go/src/bytes/bytes.go:196
		}
//line /usr/local/go/src/bytes/bytes.go:196
		// _ = "end of CoverTab[219]"
//line /usr/local/go/src/bytes/bytes.go:196
		_go_fuzz_dep_.CoverTab[220]++
							if bytealg.IndexByteString(chars, s[0]) >= 0 {
//line /usr/local/go/src/bytes/bytes.go:197
			_go_fuzz_dep_.CoverTab[228]++
								return 0
//line /usr/local/go/src/bytes/bytes.go:198
			// _ = "end of CoverTab[228]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:199
			_go_fuzz_dep_.CoverTab[229]++
//line /usr/local/go/src/bytes/bytes.go:199
			// _ = "end of CoverTab[229]"
//line /usr/local/go/src/bytes/bytes.go:199
		}
//line /usr/local/go/src/bytes/bytes.go:199
		// _ = "end of CoverTab[220]"
//line /usr/local/go/src/bytes/bytes.go:199
		_go_fuzz_dep_.CoverTab[221]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:200
		// _ = "end of CoverTab[221]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:201
		_go_fuzz_dep_.CoverTab[230]++
//line /usr/local/go/src/bytes/bytes.go:201
		// _ = "end of CoverTab[230]"
//line /usr/local/go/src/bytes/bytes.go:201
	}
//line /usr/local/go/src/bytes/bytes.go:201
	// _ = "end of CoverTab[212]"
//line /usr/local/go/src/bytes/bytes.go:201
	_go_fuzz_dep_.CoverTab[213]++
						if len(chars) == 1 {
//line /usr/local/go/src/bytes/bytes.go:202
		_go_fuzz_dep_.CoverTab[231]++
							r := rune(chars[0])
							if r >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:204
			_go_fuzz_dep_.CoverTab[233]++
								r = utf8.RuneError
//line /usr/local/go/src/bytes/bytes.go:205
			// _ = "end of CoverTab[233]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:206
			_go_fuzz_dep_.CoverTab[234]++
//line /usr/local/go/src/bytes/bytes.go:206
			// _ = "end of CoverTab[234]"
//line /usr/local/go/src/bytes/bytes.go:206
		}
//line /usr/local/go/src/bytes/bytes.go:206
		// _ = "end of CoverTab[231]"
//line /usr/local/go/src/bytes/bytes.go:206
		_go_fuzz_dep_.CoverTab[232]++
							return IndexRune(s, r)
//line /usr/local/go/src/bytes/bytes.go:207
		// _ = "end of CoverTab[232]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:208
		_go_fuzz_dep_.CoverTab[235]++
//line /usr/local/go/src/bytes/bytes.go:208
		// _ = "end of CoverTab[235]"
//line /usr/local/go/src/bytes/bytes.go:208
	}
//line /usr/local/go/src/bytes/bytes.go:208
	// _ = "end of CoverTab[213]"
//line /usr/local/go/src/bytes/bytes.go:208
	_go_fuzz_dep_.CoverTab[214]++
						if len(s) > 8 {
//line /usr/local/go/src/bytes/bytes.go:209
		_go_fuzz_dep_.CoverTab[236]++
							if as, isASCII := makeASCIISet(chars); isASCII {
//line /usr/local/go/src/bytes/bytes.go:210
			_go_fuzz_dep_.CoverTab[237]++
								for i, c := range s {
//line /usr/local/go/src/bytes/bytes.go:211
				_go_fuzz_dep_.CoverTab[239]++
									if as.contains(c) {
//line /usr/local/go/src/bytes/bytes.go:212
					_go_fuzz_dep_.CoverTab[240]++
										return i
//line /usr/local/go/src/bytes/bytes.go:213
					// _ = "end of CoverTab[240]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:214
					_go_fuzz_dep_.CoverTab[241]++
//line /usr/local/go/src/bytes/bytes.go:214
					// _ = "end of CoverTab[241]"
//line /usr/local/go/src/bytes/bytes.go:214
				}
//line /usr/local/go/src/bytes/bytes.go:214
				// _ = "end of CoverTab[239]"
			}
//line /usr/local/go/src/bytes/bytes.go:215
			// _ = "end of CoverTab[237]"
//line /usr/local/go/src/bytes/bytes.go:215
			_go_fuzz_dep_.CoverTab[238]++
								return -1
//line /usr/local/go/src/bytes/bytes.go:216
			// _ = "end of CoverTab[238]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:217
			_go_fuzz_dep_.CoverTab[242]++
//line /usr/local/go/src/bytes/bytes.go:217
			// _ = "end of CoverTab[242]"
//line /usr/local/go/src/bytes/bytes.go:217
		}
//line /usr/local/go/src/bytes/bytes.go:217
		// _ = "end of CoverTab[236]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:218
		_go_fuzz_dep_.CoverTab[243]++
//line /usr/local/go/src/bytes/bytes.go:218
		// _ = "end of CoverTab[243]"
//line /usr/local/go/src/bytes/bytes.go:218
	}
//line /usr/local/go/src/bytes/bytes.go:218
	// _ = "end of CoverTab[214]"
//line /usr/local/go/src/bytes/bytes.go:218
	_go_fuzz_dep_.CoverTab[215]++
						var width int
						for i := 0; i < len(s); i += width {
//line /usr/local/go/src/bytes/bytes.go:220
		_go_fuzz_dep_.CoverTab[244]++
							r := rune(s[i])
							if r < utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:222
			_go_fuzz_dep_.CoverTab[247]++
								if bytealg.IndexByteString(chars, s[i]) >= 0 {
//line /usr/local/go/src/bytes/bytes.go:223
				_go_fuzz_dep_.CoverTab[249]++
									return i
//line /usr/local/go/src/bytes/bytes.go:224
				// _ = "end of CoverTab[249]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:225
				_go_fuzz_dep_.CoverTab[250]++
//line /usr/local/go/src/bytes/bytes.go:225
				// _ = "end of CoverTab[250]"
//line /usr/local/go/src/bytes/bytes.go:225
			}
//line /usr/local/go/src/bytes/bytes.go:225
			// _ = "end of CoverTab[247]"
//line /usr/local/go/src/bytes/bytes.go:225
			_go_fuzz_dep_.CoverTab[248]++
								width = 1
								continue
//line /usr/local/go/src/bytes/bytes.go:227
			// _ = "end of CoverTab[248]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:228
			_go_fuzz_dep_.CoverTab[251]++
//line /usr/local/go/src/bytes/bytes.go:228
			// _ = "end of CoverTab[251]"
//line /usr/local/go/src/bytes/bytes.go:228
		}
//line /usr/local/go/src/bytes/bytes.go:228
		// _ = "end of CoverTab[244]"
//line /usr/local/go/src/bytes/bytes.go:228
		_go_fuzz_dep_.CoverTab[245]++
							r, width = utf8.DecodeRune(s[i:])
							if r != utf8.RuneError {
//line /usr/local/go/src/bytes/bytes.go:230
			_go_fuzz_dep_.CoverTab[252]++

								if len(chars) == width {
//line /usr/local/go/src/bytes/bytes.go:232
				_go_fuzz_dep_.CoverTab[254]++
									if chars == string(r) {
//line /usr/local/go/src/bytes/bytes.go:233
					_go_fuzz_dep_.CoverTab[256]++
										return i
//line /usr/local/go/src/bytes/bytes.go:234
					// _ = "end of CoverTab[256]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:235
					_go_fuzz_dep_.CoverTab[257]++
//line /usr/local/go/src/bytes/bytes.go:235
					// _ = "end of CoverTab[257]"
//line /usr/local/go/src/bytes/bytes.go:235
				}
//line /usr/local/go/src/bytes/bytes.go:235
				// _ = "end of CoverTab[254]"
//line /usr/local/go/src/bytes/bytes.go:235
				_go_fuzz_dep_.CoverTab[255]++
									continue
//line /usr/local/go/src/bytes/bytes.go:236
				// _ = "end of CoverTab[255]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:237
				_go_fuzz_dep_.CoverTab[258]++
//line /usr/local/go/src/bytes/bytes.go:237
				// _ = "end of CoverTab[258]"
//line /usr/local/go/src/bytes/bytes.go:237
			}
//line /usr/local/go/src/bytes/bytes.go:237
			// _ = "end of CoverTab[252]"
//line /usr/local/go/src/bytes/bytes.go:237
			_go_fuzz_dep_.CoverTab[253]++

								if bytealg.MaxLen >= width {
//line /usr/local/go/src/bytes/bytes.go:239
				_go_fuzz_dep_.CoverTab[259]++
									if bytealg.IndexString(chars, string(r)) >= 0 {
//line /usr/local/go/src/bytes/bytes.go:240
					_go_fuzz_dep_.CoverTab[261]++
										return i
//line /usr/local/go/src/bytes/bytes.go:241
					// _ = "end of CoverTab[261]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:242
					_go_fuzz_dep_.CoverTab[262]++
//line /usr/local/go/src/bytes/bytes.go:242
					// _ = "end of CoverTab[262]"
//line /usr/local/go/src/bytes/bytes.go:242
				}
//line /usr/local/go/src/bytes/bytes.go:242
				// _ = "end of CoverTab[259]"
//line /usr/local/go/src/bytes/bytes.go:242
				_go_fuzz_dep_.CoverTab[260]++
									continue
//line /usr/local/go/src/bytes/bytes.go:243
				// _ = "end of CoverTab[260]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:244
				_go_fuzz_dep_.CoverTab[263]++
//line /usr/local/go/src/bytes/bytes.go:244
				// _ = "end of CoverTab[263]"
//line /usr/local/go/src/bytes/bytes.go:244
			}
//line /usr/local/go/src/bytes/bytes.go:244
			// _ = "end of CoverTab[253]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:245
			_go_fuzz_dep_.CoverTab[264]++
//line /usr/local/go/src/bytes/bytes.go:245
			// _ = "end of CoverTab[264]"
//line /usr/local/go/src/bytes/bytes.go:245
		}
//line /usr/local/go/src/bytes/bytes.go:245
		// _ = "end of CoverTab[245]"
//line /usr/local/go/src/bytes/bytes.go:245
		_go_fuzz_dep_.CoverTab[246]++
							for _, ch := range chars {
//line /usr/local/go/src/bytes/bytes.go:246
			_go_fuzz_dep_.CoverTab[265]++
								if r == ch {
//line /usr/local/go/src/bytes/bytes.go:247
				_go_fuzz_dep_.CoverTab[266]++
									return i
//line /usr/local/go/src/bytes/bytes.go:248
				// _ = "end of CoverTab[266]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:249
				_go_fuzz_dep_.CoverTab[267]++
//line /usr/local/go/src/bytes/bytes.go:249
				// _ = "end of CoverTab[267]"
//line /usr/local/go/src/bytes/bytes.go:249
			}
//line /usr/local/go/src/bytes/bytes.go:249
			// _ = "end of CoverTab[265]"
		}
//line /usr/local/go/src/bytes/bytes.go:250
		// _ = "end of CoverTab[246]"
	}
//line /usr/local/go/src/bytes/bytes.go:251
	// _ = "end of CoverTab[215]"
//line /usr/local/go/src/bytes/bytes.go:251
	_go_fuzz_dep_.CoverTab[216]++
						return -1
//line /usr/local/go/src/bytes/bytes.go:252
	// _ = "end of CoverTab[216]"
}

// LastIndexAny interprets s as a sequence of UTF-8-encoded Unicode code
//line /usr/local/go/src/bytes/bytes.go:255
// points. It returns the byte index of the last occurrence in s of any of
//line /usr/local/go/src/bytes/bytes.go:255
// the Unicode code points in chars. It returns -1 if chars is empty or if
//line /usr/local/go/src/bytes/bytes.go:255
// there is no code point in common.
//line /usr/local/go/src/bytes/bytes.go:259
func LastIndexAny(s []byte, chars string) int {
//line /usr/local/go/src/bytes/bytes.go:259
	_go_fuzz_dep_.CoverTab[268]++
						if chars == "" {
//line /usr/local/go/src/bytes/bytes.go:260
		_go_fuzz_dep_.CoverTab[274]++

							return -1
//line /usr/local/go/src/bytes/bytes.go:262
		// _ = "end of CoverTab[274]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:263
		_go_fuzz_dep_.CoverTab[275]++
//line /usr/local/go/src/bytes/bytes.go:263
		// _ = "end of CoverTab[275]"
//line /usr/local/go/src/bytes/bytes.go:263
	}
//line /usr/local/go/src/bytes/bytes.go:263
	// _ = "end of CoverTab[268]"
//line /usr/local/go/src/bytes/bytes.go:263
	_go_fuzz_dep_.CoverTab[269]++
						if len(s) > 8 {
//line /usr/local/go/src/bytes/bytes.go:264
		_go_fuzz_dep_.CoverTab[276]++
							if as, isASCII := makeASCIISet(chars); isASCII {
//line /usr/local/go/src/bytes/bytes.go:265
			_go_fuzz_dep_.CoverTab[277]++
								for i := len(s) - 1; i >= 0; i-- {
//line /usr/local/go/src/bytes/bytes.go:266
				_go_fuzz_dep_.CoverTab[279]++
									if as.contains(s[i]) {
//line /usr/local/go/src/bytes/bytes.go:267
					_go_fuzz_dep_.CoverTab[280]++
										return i
//line /usr/local/go/src/bytes/bytes.go:268
					// _ = "end of CoverTab[280]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:269
					_go_fuzz_dep_.CoverTab[281]++
//line /usr/local/go/src/bytes/bytes.go:269
					// _ = "end of CoverTab[281]"
//line /usr/local/go/src/bytes/bytes.go:269
				}
//line /usr/local/go/src/bytes/bytes.go:269
				// _ = "end of CoverTab[279]"
			}
//line /usr/local/go/src/bytes/bytes.go:270
			// _ = "end of CoverTab[277]"
//line /usr/local/go/src/bytes/bytes.go:270
			_go_fuzz_dep_.CoverTab[278]++
								return -1
//line /usr/local/go/src/bytes/bytes.go:271
			// _ = "end of CoverTab[278]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:272
			_go_fuzz_dep_.CoverTab[282]++
//line /usr/local/go/src/bytes/bytes.go:272
			// _ = "end of CoverTab[282]"
//line /usr/local/go/src/bytes/bytes.go:272
		}
//line /usr/local/go/src/bytes/bytes.go:272
		// _ = "end of CoverTab[276]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:273
		_go_fuzz_dep_.CoverTab[283]++
//line /usr/local/go/src/bytes/bytes.go:273
		// _ = "end of CoverTab[283]"
//line /usr/local/go/src/bytes/bytes.go:273
	}
//line /usr/local/go/src/bytes/bytes.go:273
	// _ = "end of CoverTab[269]"
//line /usr/local/go/src/bytes/bytes.go:273
	_go_fuzz_dep_.CoverTab[270]++
						if len(s) == 1 {
//line /usr/local/go/src/bytes/bytes.go:274
		_go_fuzz_dep_.CoverTab[284]++
							r := rune(s[0])
							if r >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:276
			_go_fuzz_dep_.CoverTab[287]++
								for _, r = range chars {
//line /usr/local/go/src/bytes/bytes.go:277
				_go_fuzz_dep_.CoverTab[289]++
									if r == utf8.RuneError {
//line /usr/local/go/src/bytes/bytes.go:278
					_go_fuzz_dep_.CoverTab[290]++
										return 0
//line /usr/local/go/src/bytes/bytes.go:279
					// _ = "end of CoverTab[290]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:280
					_go_fuzz_dep_.CoverTab[291]++
//line /usr/local/go/src/bytes/bytes.go:280
					// _ = "end of CoverTab[291]"
//line /usr/local/go/src/bytes/bytes.go:280
				}
//line /usr/local/go/src/bytes/bytes.go:280
				// _ = "end of CoverTab[289]"
			}
//line /usr/local/go/src/bytes/bytes.go:281
			// _ = "end of CoverTab[287]"
//line /usr/local/go/src/bytes/bytes.go:281
			_go_fuzz_dep_.CoverTab[288]++
								return -1
//line /usr/local/go/src/bytes/bytes.go:282
			// _ = "end of CoverTab[288]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:283
			_go_fuzz_dep_.CoverTab[292]++
//line /usr/local/go/src/bytes/bytes.go:283
			// _ = "end of CoverTab[292]"
//line /usr/local/go/src/bytes/bytes.go:283
		}
//line /usr/local/go/src/bytes/bytes.go:283
		// _ = "end of CoverTab[284]"
//line /usr/local/go/src/bytes/bytes.go:283
		_go_fuzz_dep_.CoverTab[285]++
							if bytealg.IndexByteString(chars, s[0]) >= 0 {
//line /usr/local/go/src/bytes/bytes.go:284
			_go_fuzz_dep_.CoverTab[293]++
								return 0
//line /usr/local/go/src/bytes/bytes.go:285
			// _ = "end of CoverTab[293]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:286
			_go_fuzz_dep_.CoverTab[294]++
//line /usr/local/go/src/bytes/bytes.go:286
			// _ = "end of CoverTab[294]"
//line /usr/local/go/src/bytes/bytes.go:286
		}
//line /usr/local/go/src/bytes/bytes.go:286
		// _ = "end of CoverTab[285]"
//line /usr/local/go/src/bytes/bytes.go:286
		_go_fuzz_dep_.CoverTab[286]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:287
		// _ = "end of CoverTab[286]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:288
		_go_fuzz_dep_.CoverTab[295]++
//line /usr/local/go/src/bytes/bytes.go:288
		// _ = "end of CoverTab[295]"
//line /usr/local/go/src/bytes/bytes.go:288
	}
//line /usr/local/go/src/bytes/bytes.go:288
	// _ = "end of CoverTab[270]"
//line /usr/local/go/src/bytes/bytes.go:288
	_go_fuzz_dep_.CoverTab[271]++
						if len(chars) == 1 {
//line /usr/local/go/src/bytes/bytes.go:289
		_go_fuzz_dep_.CoverTab[296]++
							cr := rune(chars[0])
							if cr >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:291
			_go_fuzz_dep_.CoverTab[299]++
								cr = utf8.RuneError
//line /usr/local/go/src/bytes/bytes.go:292
			// _ = "end of CoverTab[299]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:293
			_go_fuzz_dep_.CoverTab[300]++
//line /usr/local/go/src/bytes/bytes.go:293
			// _ = "end of CoverTab[300]"
//line /usr/local/go/src/bytes/bytes.go:293
		}
//line /usr/local/go/src/bytes/bytes.go:293
		// _ = "end of CoverTab[296]"
//line /usr/local/go/src/bytes/bytes.go:293
		_go_fuzz_dep_.CoverTab[297]++
							for i := len(s); i > 0; {
//line /usr/local/go/src/bytes/bytes.go:294
			_go_fuzz_dep_.CoverTab[301]++
								r, size := utf8.DecodeLastRune(s[:i])
								i -= size
								if r == cr {
//line /usr/local/go/src/bytes/bytes.go:297
				_go_fuzz_dep_.CoverTab[302]++
									return i
//line /usr/local/go/src/bytes/bytes.go:298
				// _ = "end of CoverTab[302]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:299
				_go_fuzz_dep_.CoverTab[303]++
//line /usr/local/go/src/bytes/bytes.go:299
				// _ = "end of CoverTab[303]"
//line /usr/local/go/src/bytes/bytes.go:299
			}
//line /usr/local/go/src/bytes/bytes.go:299
			// _ = "end of CoverTab[301]"
		}
//line /usr/local/go/src/bytes/bytes.go:300
		// _ = "end of CoverTab[297]"
//line /usr/local/go/src/bytes/bytes.go:300
		_go_fuzz_dep_.CoverTab[298]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:301
		// _ = "end of CoverTab[298]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:302
		_go_fuzz_dep_.CoverTab[304]++
//line /usr/local/go/src/bytes/bytes.go:302
		// _ = "end of CoverTab[304]"
//line /usr/local/go/src/bytes/bytes.go:302
	}
//line /usr/local/go/src/bytes/bytes.go:302
	// _ = "end of CoverTab[271]"
//line /usr/local/go/src/bytes/bytes.go:302
	_go_fuzz_dep_.CoverTab[272]++
						for i := len(s); i > 0; {
//line /usr/local/go/src/bytes/bytes.go:303
		_go_fuzz_dep_.CoverTab[305]++
							r := rune(s[i-1])
							if r < utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:305
			_go_fuzz_dep_.CoverTab[308]++
								if bytealg.IndexByteString(chars, s[i-1]) >= 0 {
//line /usr/local/go/src/bytes/bytes.go:306
				_go_fuzz_dep_.CoverTab[310]++
									return i - 1
//line /usr/local/go/src/bytes/bytes.go:307
				// _ = "end of CoverTab[310]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:308
				_go_fuzz_dep_.CoverTab[311]++
//line /usr/local/go/src/bytes/bytes.go:308
				// _ = "end of CoverTab[311]"
//line /usr/local/go/src/bytes/bytes.go:308
			}
//line /usr/local/go/src/bytes/bytes.go:308
			// _ = "end of CoverTab[308]"
//line /usr/local/go/src/bytes/bytes.go:308
			_go_fuzz_dep_.CoverTab[309]++
								i--
								continue
//line /usr/local/go/src/bytes/bytes.go:310
			// _ = "end of CoverTab[309]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:311
			_go_fuzz_dep_.CoverTab[312]++
//line /usr/local/go/src/bytes/bytes.go:311
			// _ = "end of CoverTab[312]"
//line /usr/local/go/src/bytes/bytes.go:311
		}
//line /usr/local/go/src/bytes/bytes.go:311
		// _ = "end of CoverTab[305]"
//line /usr/local/go/src/bytes/bytes.go:311
		_go_fuzz_dep_.CoverTab[306]++
							r, size := utf8.DecodeLastRune(s[:i])
							i -= size
							if r != utf8.RuneError {
//line /usr/local/go/src/bytes/bytes.go:314
			_go_fuzz_dep_.CoverTab[313]++

								if len(chars) == size {
//line /usr/local/go/src/bytes/bytes.go:316
				_go_fuzz_dep_.CoverTab[315]++
									if chars == string(r) {
//line /usr/local/go/src/bytes/bytes.go:317
					_go_fuzz_dep_.CoverTab[317]++
										return i
//line /usr/local/go/src/bytes/bytes.go:318
					// _ = "end of CoverTab[317]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:319
					_go_fuzz_dep_.CoverTab[318]++
//line /usr/local/go/src/bytes/bytes.go:319
					// _ = "end of CoverTab[318]"
//line /usr/local/go/src/bytes/bytes.go:319
				}
//line /usr/local/go/src/bytes/bytes.go:319
				// _ = "end of CoverTab[315]"
//line /usr/local/go/src/bytes/bytes.go:319
				_go_fuzz_dep_.CoverTab[316]++
									continue
//line /usr/local/go/src/bytes/bytes.go:320
				// _ = "end of CoverTab[316]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:321
				_go_fuzz_dep_.CoverTab[319]++
//line /usr/local/go/src/bytes/bytes.go:321
				// _ = "end of CoverTab[319]"
//line /usr/local/go/src/bytes/bytes.go:321
			}
//line /usr/local/go/src/bytes/bytes.go:321
			// _ = "end of CoverTab[313]"
//line /usr/local/go/src/bytes/bytes.go:321
			_go_fuzz_dep_.CoverTab[314]++

								if bytealg.MaxLen >= size {
//line /usr/local/go/src/bytes/bytes.go:323
				_go_fuzz_dep_.CoverTab[320]++
									if bytealg.IndexString(chars, string(r)) >= 0 {
//line /usr/local/go/src/bytes/bytes.go:324
					_go_fuzz_dep_.CoverTab[322]++
										return i
//line /usr/local/go/src/bytes/bytes.go:325
					// _ = "end of CoverTab[322]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:326
					_go_fuzz_dep_.CoverTab[323]++
//line /usr/local/go/src/bytes/bytes.go:326
					// _ = "end of CoverTab[323]"
//line /usr/local/go/src/bytes/bytes.go:326
				}
//line /usr/local/go/src/bytes/bytes.go:326
				// _ = "end of CoverTab[320]"
//line /usr/local/go/src/bytes/bytes.go:326
				_go_fuzz_dep_.CoverTab[321]++
									continue
//line /usr/local/go/src/bytes/bytes.go:327
				// _ = "end of CoverTab[321]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:328
				_go_fuzz_dep_.CoverTab[324]++
//line /usr/local/go/src/bytes/bytes.go:328
				// _ = "end of CoverTab[324]"
//line /usr/local/go/src/bytes/bytes.go:328
			}
//line /usr/local/go/src/bytes/bytes.go:328
			// _ = "end of CoverTab[314]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:329
			_go_fuzz_dep_.CoverTab[325]++
//line /usr/local/go/src/bytes/bytes.go:329
			// _ = "end of CoverTab[325]"
//line /usr/local/go/src/bytes/bytes.go:329
		}
//line /usr/local/go/src/bytes/bytes.go:329
		// _ = "end of CoverTab[306]"
//line /usr/local/go/src/bytes/bytes.go:329
		_go_fuzz_dep_.CoverTab[307]++
							for _, ch := range chars {
//line /usr/local/go/src/bytes/bytes.go:330
			_go_fuzz_dep_.CoverTab[326]++
								if r == ch {
//line /usr/local/go/src/bytes/bytes.go:331
				_go_fuzz_dep_.CoverTab[327]++
									return i
//line /usr/local/go/src/bytes/bytes.go:332
				// _ = "end of CoverTab[327]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:333
				_go_fuzz_dep_.CoverTab[328]++
//line /usr/local/go/src/bytes/bytes.go:333
				// _ = "end of CoverTab[328]"
//line /usr/local/go/src/bytes/bytes.go:333
			}
//line /usr/local/go/src/bytes/bytes.go:333
			// _ = "end of CoverTab[326]"
		}
//line /usr/local/go/src/bytes/bytes.go:334
		// _ = "end of CoverTab[307]"
	}
//line /usr/local/go/src/bytes/bytes.go:335
	// _ = "end of CoverTab[272]"
//line /usr/local/go/src/bytes/bytes.go:335
	_go_fuzz_dep_.CoverTab[273]++
						return -1
//line /usr/local/go/src/bytes/bytes.go:336
	// _ = "end of CoverTab[273]"
}

// Generic split: splits after each instance of sep,
//line /usr/local/go/src/bytes/bytes.go:339
// including sepSave bytes of sep in the subslices.
//line /usr/local/go/src/bytes/bytes.go:341
func genSplit(s, sep []byte, sepSave, n int) [][]byte {
//line /usr/local/go/src/bytes/bytes.go:341
	_go_fuzz_dep_.CoverTab[329]++
						if n == 0 {
//line /usr/local/go/src/bytes/bytes.go:342
		_go_fuzz_dep_.CoverTab[335]++
							return nil
//line /usr/local/go/src/bytes/bytes.go:343
		// _ = "end of CoverTab[335]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:344
		_go_fuzz_dep_.CoverTab[336]++
//line /usr/local/go/src/bytes/bytes.go:344
		// _ = "end of CoverTab[336]"
//line /usr/local/go/src/bytes/bytes.go:344
	}
//line /usr/local/go/src/bytes/bytes.go:344
	// _ = "end of CoverTab[329]"
//line /usr/local/go/src/bytes/bytes.go:344
	_go_fuzz_dep_.CoverTab[330]++
						if len(sep) == 0 {
//line /usr/local/go/src/bytes/bytes.go:345
		_go_fuzz_dep_.CoverTab[337]++
							return explode(s, n)
//line /usr/local/go/src/bytes/bytes.go:346
		// _ = "end of CoverTab[337]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:347
		_go_fuzz_dep_.CoverTab[338]++
//line /usr/local/go/src/bytes/bytes.go:347
		// _ = "end of CoverTab[338]"
//line /usr/local/go/src/bytes/bytes.go:347
	}
//line /usr/local/go/src/bytes/bytes.go:347
	// _ = "end of CoverTab[330]"
//line /usr/local/go/src/bytes/bytes.go:347
	_go_fuzz_dep_.CoverTab[331]++
						if n < 0 {
//line /usr/local/go/src/bytes/bytes.go:348
		_go_fuzz_dep_.CoverTab[339]++
							n = Count(s, sep) + 1
//line /usr/local/go/src/bytes/bytes.go:349
		// _ = "end of CoverTab[339]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:350
		_go_fuzz_dep_.CoverTab[340]++
//line /usr/local/go/src/bytes/bytes.go:350
		// _ = "end of CoverTab[340]"
//line /usr/local/go/src/bytes/bytes.go:350
	}
//line /usr/local/go/src/bytes/bytes.go:350
	// _ = "end of CoverTab[331]"
//line /usr/local/go/src/bytes/bytes.go:350
	_go_fuzz_dep_.CoverTab[332]++
						if n > len(s)+1 {
//line /usr/local/go/src/bytes/bytes.go:351
		_go_fuzz_dep_.CoverTab[341]++
							n = len(s) + 1
//line /usr/local/go/src/bytes/bytes.go:352
		// _ = "end of CoverTab[341]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:353
		_go_fuzz_dep_.CoverTab[342]++
//line /usr/local/go/src/bytes/bytes.go:353
		// _ = "end of CoverTab[342]"
//line /usr/local/go/src/bytes/bytes.go:353
	}
//line /usr/local/go/src/bytes/bytes.go:353
	// _ = "end of CoverTab[332]"
//line /usr/local/go/src/bytes/bytes.go:353
	_go_fuzz_dep_.CoverTab[333]++

						a := make([][]byte, n)
						n--
						i := 0
						for i < n {
//line /usr/local/go/src/bytes/bytes.go:358
		_go_fuzz_dep_.CoverTab[343]++
							m := Index(s, sep)
							if m < 0 {
//line /usr/local/go/src/bytes/bytes.go:360
			_go_fuzz_dep_.CoverTab[345]++
								break
//line /usr/local/go/src/bytes/bytes.go:361
			// _ = "end of CoverTab[345]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:362
			_go_fuzz_dep_.CoverTab[346]++
//line /usr/local/go/src/bytes/bytes.go:362
			// _ = "end of CoverTab[346]"
//line /usr/local/go/src/bytes/bytes.go:362
		}
//line /usr/local/go/src/bytes/bytes.go:362
		// _ = "end of CoverTab[343]"
//line /usr/local/go/src/bytes/bytes.go:362
		_go_fuzz_dep_.CoverTab[344]++
							a[i] = s[: m+sepSave : m+sepSave]
							s = s[m+len(sep):]
							i++
//line /usr/local/go/src/bytes/bytes.go:365
		// _ = "end of CoverTab[344]"
	}
//line /usr/local/go/src/bytes/bytes.go:366
	// _ = "end of CoverTab[333]"
//line /usr/local/go/src/bytes/bytes.go:366
	_go_fuzz_dep_.CoverTab[334]++
						a[i] = s
						return a[:i+1]
//line /usr/local/go/src/bytes/bytes.go:368
	// _ = "end of CoverTab[334]"
}

// SplitN slices s into subslices separated by sep and returns a slice of
//line /usr/local/go/src/bytes/bytes.go:371
// the subslices between those separators.
//line /usr/local/go/src/bytes/bytes.go:371
// If sep is empty, SplitN splits after each UTF-8 sequence.
//line /usr/local/go/src/bytes/bytes.go:371
// The count determines the number of subslices to return:
//line /usr/local/go/src/bytes/bytes.go:371
//
//line /usr/local/go/src/bytes/bytes.go:371
//	n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//line /usr/local/go/src/bytes/bytes.go:371
//	n == 0: the result is nil (zero subslices)
//line /usr/local/go/src/bytes/bytes.go:371
//	n < 0: all subslices
//line /usr/local/go/src/bytes/bytes.go:371
//
//line /usr/local/go/src/bytes/bytes.go:371
// To split around the first instance of a separator, see Cut.
//line /usr/local/go/src/bytes/bytes.go:381
func SplitN(s, sep []byte, n int) [][]byte {
//line /usr/local/go/src/bytes/bytes.go:381
	_go_fuzz_dep_.CoverTab[347]++
//line /usr/local/go/src/bytes/bytes.go:381
	return genSplit(s, sep, 0, n)
//line /usr/local/go/src/bytes/bytes.go:381
	// _ = "end of CoverTab[347]"
//line /usr/local/go/src/bytes/bytes.go:381
}

// SplitAfterN slices s into subslices after each instance of sep and
//line /usr/local/go/src/bytes/bytes.go:383
// returns a slice of those subslices.
//line /usr/local/go/src/bytes/bytes.go:383
// If sep is empty, SplitAfterN splits after each UTF-8 sequence.
//line /usr/local/go/src/bytes/bytes.go:383
// The count determines the number of subslices to return:
//line /usr/local/go/src/bytes/bytes.go:383
//
//line /usr/local/go/src/bytes/bytes.go:383
//	n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//line /usr/local/go/src/bytes/bytes.go:383
//	n == 0: the result is nil (zero subslices)
//line /usr/local/go/src/bytes/bytes.go:383
//	n < 0: all subslices
//line /usr/local/go/src/bytes/bytes.go:391
func SplitAfterN(s, sep []byte, n int) [][]byte {
//line /usr/local/go/src/bytes/bytes.go:391
	_go_fuzz_dep_.CoverTab[348]++
						return genSplit(s, sep, len(sep), n)
//line /usr/local/go/src/bytes/bytes.go:392
	// _ = "end of CoverTab[348]"
}

// Split slices s into all subslices separated by sep and returns a slice of
//line /usr/local/go/src/bytes/bytes.go:395
// the subslices between those separators.
//line /usr/local/go/src/bytes/bytes.go:395
// If sep is empty, Split splits after each UTF-8 sequence.
//line /usr/local/go/src/bytes/bytes.go:395
// It is equivalent to SplitN with a count of -1.
//line /usr/local/go/src/bytes/bytes.go:395
//
//line /usr/local/go/src/bytes/bytes.go:395
// To split around the first instance of a separator, see Cut.
//line /usr/local/go/src/bytes/bytes.go:401
func Split(s, sep []byte) [][]byte {
//line /usr/local/go/src/bytes/bytes.go:401
	_go_fuzz_dep_.CoverTab[349]++
//line /usr/local/go/src/bytes/bytes.go:401
	return genSplit(s, sep, 0, -1)
//line /usr/local/go/src/bytes/bytes.go:401
	// _ = "end of CoverTab[349]"
//line /usr/local/go/src/bytes/bytes.go:401
}

// SplitAfter slices s into all subslices after each instance of sep and
//line /usr/local/go/src/bytes/bytes.go:403
// returns a slice of those subslices.
//line /usr/local/go/src/bytes/bytes.go:403
// If sep is empty, SplitAfter splits after each UTF-8 sequence.
//line /usr/local/go/src/bytes/bytes.go:403
// It is equivalent to SplitAfterN with a count of -1.
//line /usr/local/go/src/bytes/bytes.go:407
func SplitAfter(s, sep []byte) [][]byte {
//line /usr/local/go/src/bytes/bytes.go:407
	_go_fuzz_dep_.CoverTab[350]++
						return genSplit(s, sep, len(sep), -1)
//line /usr/local/go/src/bytes/bytes.go:408
	// _ = "end of CoverTab[350]"
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// Fields interprets s as a sequence of UTF-8-encoded code points.
//line /usr/local/go/src/bytes/bytes.go:413
// It splits the slice s around each instance of one or more consecutive white space
//line /usr/local/go/src/bytes/bytes.go:413
// characters, as defined by unicode.IsSpace, returning a slice of subslices of s or an
//line /usr/local/go/src/bytes/bytes.go:413
// empty slice if s contains only white space.
//line /usr/local/go/src/bytes/bytes.go:417
func Fields(s []byte) [][]byte {
//line /usr/local/go/src/bytes/bytes.go:417
	_go_fuzz_dep_.CoverTab[351]++

//line /usr/local/go/src/bytes/bytes.go:420
	n := 0
	wasSpace := 1

	setBits := uint8(0)
	for i := 0; i < len(s); i++ {
//line /usr/local/go/src/bytes/bytes.go:424
		_go_fuzz_dep_.CoverTab[357]++
							r := s[i]
							setBits |= r
							isSpace := int(asciiSpace[r])
							n += wasSpace & ^isSpace
							wasSpace = isSpace
//line /usr/local/go/src/bytes/bytes.go:429
		// _ = "end of CoverTab[357]"
	}
//line /usr/local/go/src/bytes/bytes.go:430
	// _ = "end of CoverTab[351]"
//line /usr/local/go/src/bytes/bytes.go:430
	_go_fuzz_dep_.CoverTab[352]++

						if setBits >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:432
		_go_fuzz_dep_.CoverTab[358]++

							return FieldsFunc(s, unicode.IsSpace)
//line /usr/local/go/src/bytes/bytes.go:434
		// _ = "end of CoverTab[358]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:435
		_go_fuzz_dep_.CoverTab[359]++
//line /usr/local/go/src/bytes/bytes.go:435
		// _ = "end of CoverTab[359]"
//line /usr/local/go/src/bytes/bytes.go:435
	}
//line /usr/local/go/src/bytes/bytes.go:435
	// _ = "end of CoverTab[352]"
//line /usr/local/go/src/bytes/bytes.go:435
	_go_fuzz_dep_.CoverTab[353]++

//line /usr/local/go/src/bytes/bytes.go:438
	a := make([][]byte, n)
	na := 0
	fieldStart := 0
	i := 0

	for i < len(s) && func() bool {
//line /usr/local/go/src/bytes/bytes.go:443
		_go_fuzz_dep_.CoverTab[360]++
//line /usr/local/go/src/bytes/bytes.go:443
		return asciiSpace[s[i]] != 0
//line /usr/local/go/src/bytes/bytes.go:443
		// _ = "end of CoverTab[360]"
//line /usr/local/go/src/bytes/bytes.go:443
	}() {
//line /usr/local/go/src/bytes/bytes.go:443
		_go_fuzz_dep_.CoverTab[361]++
							i++
//line /usr/local/go/src/bytes/bytes.go:444
		// _ = "end of CoverTab[361]"
	}
//line /usr/local/go/src/bytes/bytes.go:445
	// _ = "end of CoverTab[353]"
//line /usr/local/go/src/bytes/bytes.go:445
	_go_fuzz_dep_.CoverTab[354]++
						fieldStart = i
						for i < len(s) {
//line /usr/local/go/src/bytes/bytes.go:447
		_go_fuzz_dep_.CoverTab[362]++
							if asciiSpace[s[i]] == 0 {
//line /usr/local/go/src/bytes/bytes.go:448
			_go_fuzz_dep_.CoverTab[365]++
								i++
								continue
//line /usr/local/go/src/bytes/bytes.go:450
			// _ = "end of CoverTab[365]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:451
			_go_fuzz_dep_.CoverTab[366]++
//line /usr/local/go/src/bytes/bytes.go:451
			// _ = "end of CoverTab[366]"
//line /usr/local/go/src/bytes/bytes.go:451
		}
//line /usr/local/go/src/bytes/bytes.go:451
		// _ = "end of CoverTab[362]"
//line /usr/local/go/src/bytes/bytes.go:451
		_go_fuzz_dep_.CoverTab[363]++
							a[na] = s[fieldStart:i:i]
							na++
							i++

							for i < len(s) && func() bool {
//line /usr/local/go/src/bytes/bytes.go:456
			_go_fuzz_dep_.CoverTab[367]++
//line /usr/local/go/src/bytes/bytes.go:456
			return asciiSpace[s[i]] != 0
//line /usr/local/go/src/bytes/bytes.go:456
			// _ = "end of CoverTab[367]"
//line /usr/local/go/src/bytes/bytes.go:456
		}() {
//line /usr/local/go/src/bytes/bytes.go:456
			_go_fuzz_dep_.CoverTab[368]++
								i++
//line /usr/local/go/src/bytes/bytes.go:457
			// _ = "end of CoverTab[368]"
		}
//line /usr/local/go/src/bytes/bytes.go:458
		// _ = "end of CoverTab[363]"
//line /usr/local/go/src/bytes/bytes.go:458
		_go_fuzz_dep_.CoverTab[364]++
							fieldStart = i
//line /usr/local/go/src/bytes/bytes.go:459
		// _ = "end of CoverTab[364]"
	}
//line /usr/local/go/src/bytes/bytes.go:460
	// _ = "end of CoverTab[354]"
//line /usr/local/go/src/bytes/bytes.go:460
	_go_fuzz_dep_.CoverTab[355]++
						if fieldStart < len(s) {
//line /usr/local/go/src/bytes/bytes.go:461
		_go_fuzz_dep_.CoverTab[369]++
							a[na] = s[fieldStart:len(s):len(s)]
//line /usr/local/go/src/bytes/bytes.go:462
		// _ = "end of CoverTab[369]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:463
		_go_fuzz_dep_.CoverTab[370]++
//line /usr/local/go/src/bytes/bytes.go:463
		// _ = "end of CoverTab[370]"
//line /usr/local/go/src/bytes/bytes.go:463
	}
//line /usr/local/go/src/bytes/bytes.go:463
	// _ = "end of CoverTab[355]"
//line /usr/local/go/src/bytes/bytes.go:463
	_go_fuzz_dep_.CoverTab[356]++
						return a
//line /usr/local/go/src/bytes/bytes.go:464
	// _ = "end of CoverTab[356]"
}

// FieldsFunc interprets s as a sequence of UTF-8-encoded code points.
//line /usr/local/go/src/bytes/bytes.go:467
// It splits the slice s at each run of code points c satisfying f(c) and
//line /usr/local/go/src/bytes/bytes.go:467
// returns a slice of subslices of s. If all code points in s satisfy f(c), or
//line /usr/local/go/src/bytes/bytes.go:467
// len(s) == 0, an empty slice is returned.
//line /usr/local/go/src/bytes/bytes.go:467
//
//line /usr/local/go/src/bytes/bytes.go:467
// FieldsFunc makes no guarantees about the order in which it calls f(c)
//line /usr/local/go/src/bytes/bytes.go:467
// and assumes that f always returns the same value for a given c.
//line /usr/local/go/src/bytes/bytes.go:474
func FieldsFunc(s []byte, f func(rune) bool) [][]byte {
//line /usr/local/go/src/bytes/bytes.go:474
	_go_fuzz_dep_.CoverTab[371]++
	// A span is used to record a slice of s of the form s[start:end].
	// The start index is inclusive and the end index is exclusive.
	type span struct {
		start	int
		end	int
	}
						spans := make([]span, 0, 32)

//line /usr/local/go/src/bytes/bytes.go:487
	start := -1
	for i := 0; i < len(s); {
//line /usr/local/go/src/bytes/bytes.go:488
		_go_fuzz_dep_.CoverTab[375]++
							size := 1
							r := rune(s[i])
							if r >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:491
			_go_fuzz_dep_.CoverTab[378]++
								r, size = utf8.DecodeRune(s[i:])
//line /usr/local/go/src/bytes/bytes.go:492
			// _ = "end of CoverTab[378]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:493
			_go_fuzz_dep_.CoverTab[379]++
//line /usr/local/go/src/bytes/bytes.go:493
			// _ = "end of CoverTab[379]"
//line /usr/local/go/src/bytes/bytes.go:493
		}
//line /usr/local/go/src/bytes/bytes.go:493
		// _ = "end of CoverTab[375]"
//line /usr/local/go/src/bytes/bytes.go:493
		_go_fuzz_dep_.CoverTab[376]++
							if f(r) {
//line /usr/local/go/src/bytes/bytes.go:494
			_go_fuzz_dep_.CoverTab[380]++
								if start >= 0 {
//line /usr/local/go/src/bytes/bytes.go:495
				_go_fuzz_dep_.CoverTab[381]++
									spans = append(spans, span{start, i})
									start = -1
//line /usr/local/go/src/bytes/bytes.go:497
				// _ = "end of CoverTab[381]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:498
				_go_fuzz_dep_.CoverTab[382]++
//line /usr/local/go/src/bytes/bytes.go:498
				// _ = "end of CoverTab[382]"
//line /usr/local/go/src/bytes/bytes.go:498
			}
//line /usr/local/go/src/bytes/bytes.go:498
			// _ = "end of CoverTab[380]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:499
			_go_fuzz_dep_.CoverTab[383]++
								if start < 0 {
//line /usr/local/go/src/bytes/bytes.go:500
				_go_fuzz_dep_.CoverTab[384]++
									start = i
//line /usr/local/go/src/bytes/bytes.go:501
				// _ = "end of CoverTab[384]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:502
				_go_fuzz_dep_.CoverTab[385]++
//line /usr/local/go/src/bytes/bytes.go:502
				// _ = "end of CoverTab[385]"
//line /usr/local/go/src/bytes/bytes.go:502
			}
//line /usr/local/go/src/bytes/bytes.go:502
			// _ = "end of CoverTab[383]"
		}
//line /usr/local/go/src/bytes/bytes.go:503
		// _ = "end of CoverTab[376]"
//line /usr/local/go/src/bytes/bytes.go:503
		_go_fuzz_dep_.CoverTab[377]++
							i += size
//line /usr/local/go/src/bytes/bytes.go:504
		// _ = "end of CoverTab[377]"
	}
//line /usr/local/go/src/bytes/bytes.go:505
	// _ = "end of CoverTab[371]"
//line /usr/local/go/src/bytes/bytes.go:505
	_go_fuzz_dep_.CoverTab[372]++

//line /usr/local/go/src/bytes/bytes.go:508
	if start >= 0 {
//line /usr/local/go/src/bytes/bytes.go:508
		_go_fuzz_dep_.CoverTab[386]++
							spans = append(spans, span{start, len(s)})
//line /usr/local/go/src/bytes/bytes.go:509
		// _ = "end of CoverTab[386]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:510
		_go_fuzz_dep_.CoverTab[387]++
//line /usr/local/go/src/bytes/bytes.go:510
		// _ = "end of CoverTab[387]"
//line /usr/local/go/src/bytes/bytes.go:510
	}
//line /usr/local/go/src/bytes/bytes.go:510
	// _ = "end of CoverTab[372]"
//line /usr/local/go/src/bytes/bytes.go:510
	_go_fuzz_dep_.CoverTab[373]++

//line /usr/local/go/src/bytes/bytes.go:513
	a := make([][]byte, len(spans))
	for i, span := range spans {
//line /usr/local/go/src/bytes/bytes.go:514
		_go_fuzz_dep_.CoverTab[388]++
							a[i] = s[span.start:span.end:span.end]
//line /usr/local/go/src/bytes/bytes.go:515
		// _ = "end of CoverTab[388]"
	}
//line /usr/local/go/src/bytes/bytes.go:516
	// _ = "end of CoverTab[373]"
//line /usr/local/go/src/bytes/bytes.go:516
	_go_fuzz_dep_.CoverTab[374]++

						return a
//line /usr/local/go/src/bytes/bytes.go:518
	// _ = "end of CoverTab[374]"
}

// Join concatenates the elements of s to create a new byte slice. The separator
//line /usr/local/go/src/bytes/bytes.go:521
// sep is placed between elements in the resulting slice.
//line /usr/local/go/src/bytes/bytes.go:523
func Join(s [][]byte, sep []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:523
	_go_fuzz_dep_.CoverTab[389]++
						if len(s) == 0 {
//line /usr/local/go/src/bytes/bytes.go:524
		_go_fuzz_dep_.CoverTab[394]++
							return []byte{}
//line /usr/local/go/src/bytes/bytes.go:525
		// _ = "end of CoverTab[394]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:526
		_go_fuzz_dep_.CoverTab[395]++
//line /usr/local/go/src/bytes/bytes.go:526
		// _ = "end of CoverTab[395]"
//line /usr/local/go/src/bytes/bytes.go:526
	}
//line /usr/local/go/src/bytes/bytes.go:526
	// _ = "end of CoverTab[389]"
//line /usr/local/go/src/bytes/bytes.go:526
	_go_fuzz_dep_.CoverTab[390]++
						if len(s) == 1 {
//line /usr/local/go/src/bytes/bytes.go:527
		_go_fuzz_dep_.CoverTab[396]++

							return append([]byte(nil), s[0]...)
//line /usr/local/go/src/bytes/bytes.go:529
		// _ = "end of CoverTab[396]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:530
		_go_fuzz_dep_.CoverTab[397]++
//line /usr/local/go/src/bytes/bytes.go:530
		// _ = "end of CoverTab[397]"
//line /usr/local/go/src/bytes/bytes.go:530
	}
//line /usr/local/go/src/bytes/bytes.go:530
	// _ = "end of CoverTab[390]"
//line /usr/local/go/src/bytes/bytes.go:530
	_go_fuzz_dep_.CoverTab[391]++
						n := len(sep) * (len(s) - 1)
						for _, v := range s {
//line /usr/local/go/src/bytes/bytes.go:532
		_go_fuzz_dep_.CoverTab[398]++
							n += len(v)
//line /usr/local/go/src/bytes/bytes.go:533
		// _ = "end of CoverTab[398]"
	}
//line /usr/local/go/src/bytes/bytes.go:534
	// _ = "end of CoverTab[391]"
//line /usr/local/go/src/bytes/bytes.go:534
	_go_fuzz_dep_.CoverTab[392]++

						b := make([]byte, n)
						bp := copy(b, s[0])
						for _, v := range s[1:] {
//line /usr/local/go/src/bytes/bytes.go:538
		_go_fuzz_dep_.CoverTab[399]++
							bp += copy(b[bp:], sep)
							bp += copy(b[bp:], v)
//line /usr/local/go/src/bytes/bytes.go:540
		// _ = "end of CoverTab[399]"
	}
//line /usr/local/go/src/bytes/bytes.go:541
	// _ = "end of CoverTab[392]"
//line /usr/local/go/src/bytes/bytes.go:541
	_go_fuzz_dep_.CoverTab[393]++
						return b
//line /usr/local/go/src/bytes/bytes.go:542
	// _ = "end of CoverTab[393]"
}

// HasPrefix tests whether the byte slice s begins with prefix.
func HasPrefix(s, prefix []byte) bool {
//line /usr/local/go/src/bytes/bytes.go:546
	_go_fuzz_dep_.CoverTab[400]++
						return len(s) >= len(prefix) && func() bool {
//line /usr/local/go/src/bytes/bytes.go:547
		_go_fuzz_dep_.CoverTab[401]++
//line /usr/local/go/src/bytes/bytes.go:547
		return Equal(s[0:len(prefix)], prefix)
//line /usr/local/go/src/bytes/bytes.go:547
		// _ = "end of CoverTab[401]"
//line /usr/local/go/src/bytes/bytes.go:547
	}()
//line /usr/local/go/src/bytes/bytes.go:547
	// _ = "end of CoverTab[400]"
}

// HasSuffix tests whether the byte slice s ends with suffix.
func HasSuffix(s, suffix []byte) bool {
//line /usr/local/go/src/bytes/bytes.go:551
	_go_fuzz_dep_.CoverTab[402]++
						return len(s) >= len(suffix) && func() bool {
//line /usr/local/go/src/bytes/bytes.go:552
		_go_fuzz_dep_.CoverTab[403]++
//line /usr/local/go/src/bytes/bytes.go:552
		return Equal(s[len(s)-len(suffix):], suffix)
//line /usr/local/go/src/bytes/bytes.go:552
		// _ = "end of CoverTab[403]"
//line /usr/local/go/src/bytes/bytes.go:552
	}()
//line /usr/local/go/src/bytes/bytes.go:552
	// _ = "end of CoverTab[402]"
}

// Map returns a copy of the byte slice s with all its characters modified
//line /usr/local/go/src/bytes/bytes.go:555
// according to the mapping function. If mapping returns a negative value, the character is
//line /usr/local/go/src/bytes/bytes.go:555
// dropped from the byte slice with no replacement. The characters in s and the
//line /usr/local/go/src/bytes/bytes.go:555
// output are interpreted as UTF-8-encoded code points.
//line /usr/local/go/src/bytes/bytes.go:559
func Map(mapping func(r rune) rune, s []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:559
	_go_fuzz_dep_.CoverTab[404]++

//line /usr/local/go/src/bytes/bytes.go:563
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); {
//line /usr/local/go/src/bytes/bytes.go:564
		_go_fuzz_dep_.CoverTab[406]++
							wid := 1
							r := rune(s[i])
							if r >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:567
			_go_fuzz_dep_.CoverTab[409]++
								r, wid = utf8.DecodeRune(s[i:])
//line /usr/local/go/src/bytes/bytes.go:568
			// _ = "end of CoverTab[409]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:569
			_go_fuzz_dep_.CoverTab[410]++
//line /usr/local/go/src/bytes/bytes.go:569
			// _ = "end of CoverTab[410]"
//line /usr/local/go/src/bytes/bytes.go:569
		}
//line /usr/local/go/src/bytes/bytes.go:569
		// _ = "end of CoverTab[406]"
//line /usr/local/go/src/bytes/bytes.go:569
		_go_fuzz_dep_.CoverTab[407]++
							r = mapping(r)
							if r >= 0 {
//line /usr/local/go/src/bytes/bytes.go:571
			_go_fuzz_dep_.CoverTab[411]++
								b = utf8.AppendRune(b, r)
//line /usr/local/go/src/bytes/bytes.go:572
			// _ = "end of CoverTab[411]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:573
			_go_fuzz_dep_.CoverTab[412]++
//line /usr/local/go/src/bytes/bytes.go:573
			// _ = "end of CoverTab[412]"
//line /usr/local/go/src/bytes/bytes.go:573
		}
//line /usr/local/go/src/bytes/bytes.go:573
		// _ = "end of CoverTab[407]"
//line /usr/local/go/src/bytes/bytes.go:573
		_go_fuzz_dep_.CoverTab[408]++
							i += wid
//line /usr/local/go/src/bytes/bytes.go:574
		// _ = "end of CoverTab[408]"
	}
//line /usr/local/go/src/bytes/bytes.go:575
	// _ = "end of CoverTab[404]"
//line /usr/local/go/src/bytes/bytes.go:575
	_go_fuzz_dep_.CoverTab[405]++
						return b
//line /usr/local/go/src/bytes/bytes.go:576
	// _ = "end of CoverTab[405]"
}

// Repeat returns a new byte slice consisting of count copies of b.
//line /usr/local/go/src/bytes/bytes.go:579
//
//line /usr/local/go/src/bytes/bytes.go:579
// It panics if count is negative or if the result of (len(b) * count)
//line /usr/local/go/src/bytes/bytes.go:579
// overflows.
//line /usr/local/go/src/bytes/bytes.go:583
func Repeat(b []byte, count int) []byte {
//line /usr/local/go/src/bytes/bytes.go:583
	_go_fuzz_dep_.CoverTab[413]++
						if count == 0 {
//line /usr/local/go/src/bytes/bytes.go:584
		_go_fuzz_dep_.CoverTab[419]++
							return []byte{}
//line /usr/local/go/src/bytes/bytes.go:585
		// _ = "end of CoverTab[419]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:586
		_go_fuzz_dep_.CoverTab[420]++
//line /usr/local/go/src/bytes/bytes.go:586
		// _ = "end of CoverTab[420]"
//line /usr/local/go/src/bytes/bytes.go:586
	}
//line /usr/local/go/src/bytes/bytes.go:586
	// _ = "end of CoverTab[413]"
//line /usr/local/go/src/bytes/bytes.go:586
	_go_fuzz_dep_.CoverTab[414]++

//line /usr/local/go/src/bytes/bytes.go:591
	if count < 0 {
//line /usr/local/go/src/bytes/bytes.go:591
		_go_fuzz_dep_.CoverTab[421]++
							panic("bytes: negative Repeat count")
//line /usr/local/go/src/bytes/bytes.go:592
		// _ = "end of CoverTab[421]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:593
		_go_fuzz_dep_.CoverTab[422]++
//line /usr/local/go/src/bytes/bytes.go:593
		if len(b)*count/count != len(b) {
//line /usr/local/go/src/bytes/bytes.go:593
			_go_fuzz_dep_.CoverTab[423]++
								panic("bytes: Repeat count causes overflow")
//line /usr/local/go/src/bytes/bytes.go:594
			// _ = "end of CoverTab[423]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:595
			_go_fuzz_dep_.CoverTab[424]++
//line /usr/local/go/src/bytes/bytes.go:595
			// _ = "end of CoverTab[424]"
//line /usr/local/go/src/bytes/bytes.go:595
		}
//line /usr/local/go/src/bytes/bytes.go:595
		// _ = "end of CoverTab[422]"
//line /usr/local/go/src/bytes/bytes.go:595
	}
//line /usr/local/go/src/bytes/bytes.go:595
	// _ = "end of CoverTab[414]"
//line /usr/local/go/src/bytes/bytes.go:595
	_go_fuzz_dep_.CoverTab[415]++

						if len(b) == 0 {
//line /usr/local/go/src/bytes/bytes.go:597
		_go_fuzz_dep_.CoverTab[425]++
							return []byte{}
//line /usr/local/go/src/bytes/bytes.go:598
		// _ = "end of CoverTab[425]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:599
		_go_fuzz_dep_.CoverTab[426]++
//line /usr/local/go/src/bytes/bytes.go:599
		// _ = "end of CoverTab[426]"
//line /usr/local/go/src/bytes/bytes.go:599
	}
//line /usr/local/go/src/bytes/bytes.go:599
	// _ = "end of CoverTab[415]"
//line /usr/local/go/src/bytes/bytes.go:599
	_go_fuzz_dep_.CoverTab[416]++

						n := len(b) * count

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
	if chunkMax > chunkLimit {
//line /usr/local/go/src/bytes/bytes.go:615
		_go_fuzz_dep_.CoverTab[427]++
							chunkMax = chunkLimit / len(b) * len(b)
							if chunkMax == 0 {
//line /usr/local/go/src/bytes/bytes.go:617
			_go_fuzz_dep_.CoverTab[428]++
								chunkMax = len(b)
//line /usr/local/go/src/bytes/bytes.go:618
			// _ = "end of CoverTab[428]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:619
			_go_fuzz_dep_.CoverTab[429]++
//line /usr/local/go/src/bytes/bytes.go:619
			// _ = "end of CoverTab[429]"
//line /usr/local/go/src/bytes/bytes.go:619
		}
//line /usr/local/go/src/bytes/bytes.go:619
		// _ = "end of CoverTab[427]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:620
		_go_fuzz_dep_.CoverTab[430]++
//line /usr/local/go/src/bytes/bytes.go:620
		// _ = "end of CoverTab[430]"
//line /usr/local/go/src/bytes/bytes.go:620
	}
//line /usr/local/go/src/bytes/bytes.go:620
	// _ = "end of CoverTab[416]"
//line /usr/local/go/src/bytes/bytes.go:620
	_go_fuzz_dep_.CoverTab[417]++
						nb := make([]byte, n)
						bp := copy(nb, b)
						for bp < len(nb) {
//line /usr/local/go/src/bytes/bytes.go:623
		_go_fuzz_dep_.CoverTab[431]++
							chunk := bp
							if chunk > chunkMax {
//line /usr/local/go/src/bytes/bytes.go:625
			_go_fuzz_dep_.CoverTab[433]++
								chunk = chunkMax
//line /usr/local/go/src/bytes/bytes.go:626
			// _ = "end of CoverTab[433]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:627
			_go_fuzz_dep_.CoverTab[434]++
//line /usr/local/go/src/bytes/bytes.go:627
			// _ = "end of CoverTab[434]"
//line /usr/local/go/src/bytes/bytes.go:627
		}
//line /usr/local/go/src/bytes/bytes.go:627
		// _ = "end of CoverTab[431]"
//line /usr/local/go/src/bytes/bytes.go:627
		_go_fuzz_dep_.CoverTab[432]++
							bp += copy(nb[bp:], nb[:chunk])
//line /usr/local/go/src/bytes/bytes.go:628
		// _ = "end of CoverTab[432]"
	}
//line /usr/local/go/src/bytes/bytes.go:629
	// _ = "end of CoverTab[417]"
//line /usr/local/go/src/bytes/bytes.go:629
	_go_fuzz_dep_.CoverTab[418]++
						return nb
//line /usr/local/go/src/bytes/bytes.go:630
	// _ = "end of CoverTab[418]"
}

// ToUpper returns a copy of the byte slice s with all Unicode letters mapped to
//line /usr/local/go/src/bytes/bytes.go:633
// their upper case.
//line /usr/local/go/src/bytes/bytes.go:635
func ToUpper(s []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:635
	_go_fuzz_dep_.CoverTab[435]++
						isASCII, hasLower := true, false
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/bytes/bytes.go:637
		_go_fuzz_dep_.CoverTab[438]++
							c := s[i]
							if c >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:639
			_go_fuzz_dep_.CoverTab[440]++
								isASCII = false
								break
//line /usr/local/go/src/bytes/bytes.go:641
			// _ = "end of CoverTab[440]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:642
			_go_fuzz_dep_.CoverTab[441]++
//line /usr/local/go/src/bytes/bytes.go:642
			// _ = "end of CoverTab[441]"
//line /usr/local/go/src/bytes/bytes.go:642
		}
//line /usr/local/go/src/bytes/bytes.go:642
		// _ = "end of CoverTab[438]"
//line /usr/local/go/src/bytes/bytes.go:642
		_go_fuzz_dep_.CoverTab[439]++
							hasLower = hasLower || func() bool {
//line /usr/local/go/src/bytes/bytes.go:643
			_go_fuzz_dep_.CoverTab[442]++
//line /usr/local/go/src/bytes/bytes.go:643
			return ('a' <= c && func() bool {
//line /usr/local/go/src/bytes/bytes.go:643
				_go_fuzz_dep_.CoverTab[443]++
//line /usr/local/go/src/bytes/bytes.go:643
				return c <= 'z'
//line /usr/local/go/src/bytes/bytes.go:643
				// _ = "end of CoverTab[443]"
//line /usr/local/go/src/bytes/bytes.go:643
			}())
//line /usr/local/go/src/bytes/bytes.go:643
			// _ = "end of CoverTab[442]"
//line /usr/local/go/src/bytes/bytes.go:643
		}()
//line /usr/local/go/src/bytes/bytes.go:643
		// _ = "end of CoverTab[439]"
	}
//line /usr/local/go/src/bytes/bytes.go:644
	// _ = "end of CoverTab[435]"
//line /usr/local/go/src/bytes/bytes.go:644
	_go_fuzz_dep_.CoverTab[436]++

						if isASCII {
//line /usr/local/go/src/bytes/bytes.go:646
		_go_fuzz_dep_.CoverTab[444]++
							if !hasLower {
//line /usr/local/go/src/bytes/bytes.go:647
			_go_fuzz_dep_.CoverTab[447]++

								return append([]byte(""), s...)
//line /usr/local/go/src/bytes/bytes.go:649
			// _ = "end of CoverTab[447]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:650
			_go_fuzz_dep_.CoverTab[448]++
//line /usr/local/go/src/bytes/bytes.go:650
			// _ = "end of CoverTab[448]"
//line /usr/local/go/src/bytes/bytes.go:650
		}
//line /usr/local/go/src/bytes/bytes.go:650
		// _ = "end of CoverTab[444]"
//line /usr/local/go/src/bytes/bytes.go:650
		_go_fuzz_dep_.CoverTab[445]++
							b := make([]byte, len(s))
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/bytes/bytes.go:652
			_go_fuzz_dep_.CoverTab[449]++
								c := s[i]
								if 'a' <= c && func() bool {
//line /usr/local/go/src/bytes/bytes.go:654
				_go_fuzz_dep_.CoverTab[451]++
//line /usr/local/go/src/bytes/bytes.go:654
				return c <= 'z'
//line /usr/local/go/src/bytes/bytes.go:654
				// _ = "end of CoverTab[451]"
//line /usr/local/go/src/bytes/bytes.go:654
			}() {
//line /usr/local/go/src/bytes/bytes.go:654
				_go_fuzz_dep_.CoverTab[452]++
									c -= 'a' - 'A'
//line /usr/local/go/src/bytes/bytes.go:655
				// _ = "end of CoverTab[452]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:656
				_go_fuzz_dep_.CoverTab[453]++
//line /usr/local/go/src/bytes/bytes.go:656
				// _ = "end of CoverTab[453]"
//line /usr/local/go/src/bytes/bytes.go:656
			}
//line /usr/local/go/src/bytes/bytes.go:656
			// _ = "end of CoverTab[449]"
//line /usr/local/go/src/bytes/bytes.go:656
			_go_fuzz_dep_.CoverTab[450]++
								b[i] = c
//line /usr/local/go/src/bytes/bytes.go:657
			// _ = "end of CoverTab[450]"
		}
//line /usr/local/go/src/bytes/bytes.go:658
		// _ = "end of CoverTab[445]"
//line /usr/local/go/src/bytes/bytes.go:658
		_go_fuzz_dep_.CoverTab[446]++
							return b
//line /usr/local/go/src/bytes/bytes.go:659
		// _ = "end of CoverTab[446]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:660
		_go_fuzz_dep_.CoverTab[454]++
//line /usr/local/go/src/bytes/bytes.go:660
		// _ = "end of CoverTab[454]"
//line /usr/local/go/src/bytes/bytes.go:660
	}
//line /usr/local/go/src/bytes/bytes.go:660
	// _ = "end of CoverTab[436]"
//line /usr/local/go/src/bytes/bytes.go:660
	_go_fuzz_dep_.CoverTab[437]++
						return Map(unicode.ToUpper, s)
//line /usr/local/go/src/bytes/bytes.go:661
	// _ = "end of CoverTab[437]"
}

// ToLower returns a copy of the byte slice s with all Unicode letters mapped to
//line /usr/local/go/src/bytes/bytes.go:664
// their lower case.
//line /usr/local/go/src/bytes/bytes.go:666
func ToLower(s []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:666
	_go_fuzz_dep_.CoverTab[455]++
						isASCII, hasUpper := true, false
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/bytes/bytes.go:668
		_go_fuzz_dep_.CoverTab[458]++
							c := s[i]
							if c >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:670
			_go_fuzz_dep_.CoverTab[460]++
								isASCII = false
								break
//line /usr/local/go/src/bytes/bytes.go:672
			// _ = "end of CoverTab[460]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:673
			_go_fuzz_dep_.CoverTab[461]++
//line /usr/local/go/src/bytes/bytes.go:673
			// _ = "end of CoverTab[461]"
//line /usr/local/go/src/bytes/bytes.go:673
		}
//line /usr/local/go/src/bytes/bytes.go:673
		// _ = "end of CoverTab[458]"
//line /usr/local/go/src/bytes/bytes.go:673
		_go_fuzz_dep_.CoverTab[459]++
							hasUpper = hasUpper || func() bool {
//line /usr/local/go/src/bytes/bytes.go:674
			_go_fuzz_dep_.CoverTab[462]++
//line /usr/local/go/src/bytes/bytes.go:674
			return ('A' <= c && func() bool {
//line /usr/local/go/src/bytes/bytes.go:674
				_go_fuzz_dep_.CoverTab[463]++
//line /usr/local/go/src/bytes/bytes.go:674
				return c <= 'Z'
//line /usr/local/go/src/bytes/bytes.go:674
				// _ = "end of CoverTab[463]"
//line /usr/local/go/src/bytes/bytes.go:674
			}())
//line /usr/local/go/src/bytes/bytes.go:674
			// _ = "end of CoverTab[462]"
//line /usr/local/go/src/bytes/bytes.go:674
		}()
//line /usr/local/go/src/bytes/bytes.go:674
		// _ = "end of CoverTab[459]"
	}
//line /usr/local/go/src/bytes/bytes.go:675
	// _ = "end of CoverTab[455]"
//line /usr/local/go/src/bytes/bytes.go:675
	_go_fuzz_dep_.CoverTab[456]++

						if isASCII {
//line /usr/local/go/src/bytes/bytes.go:677
		_go_fuzz_dep_.CoverTab[464]++
							if !hasUpper {
//line /usr/local/go/src/bytes/bytes.go:678
			_go_fuzz_dep_.CoverTab[467]++
								return append([]byte(""), s...)
//line /usr/local/go/src/bytes/bytes.go:679
			// _ = "end of CoverTab[467]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:680
			_go_fuzz_dep_.CoverTab[468]++
//line /usr/local/go/src/bytes/bytes.go:680
			// _ = "end of CoverTab[468]"
//line /usr/local/go/src/bytes/bytes.go:680
		}
//line /usr/local/go/src/bytes/bytes.go:680
		// _ = "end of CoverTab[464]"
//line /usr/local/go/src/bytes/bytes.go:680
		_go_fuzz_dep_.CoverTab[465]++
							b := make([]byte, len(s))
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/bytes/bytes.go:682
			_go_fuzz_dep_.CoverTab[469]++
								c := s[i]
								if 'A' <= c && func() bool {
//line /usr/local/go/src/bytes/bytes.go:684
				_go_fuzz_dep_.CoverTab[471]++
//line /usr/local/go/src/bytes/bytes.go:684
				return c <= 'Z'
//line /usr/local/go/src/bytes/bytes.go:684
				// _ = "end of CoverTab[471]"
//line /usr/local/go/src/bytes/bytes.go:684
			}() {
//line /usr/local/go/src/bytes/bytes.go:684
				_go_fuzz_dep_.CoverTab[472]++
									c += 'a' - 'A'
//line /usr/local/go/src/bytes/bytes.go:685
				// _ = "end of CoverTab[472]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:686
				_go_fuzz_dep_.CoverTab[473]++
//line /usr/local/go/src/bytes/bytes.go:686
				// _ = "end of CoverTab[473]"
//line /usr/local/go/src/bytes/bytes.go:686
			}
//line /usr/local/go/src/bytes/bytes.go:686
			// _ = "end of CoverTab[469]"
//line /usr/local/go/src/bytes/bytes.go:686
			_go_fuzz_dep_.CoverTab[470]++
								b[i] = c
//line /usr/local/go/src/bytes/bytes.go:687
			// _ = "end of CoverTab[470]"
		}
//line /usr/local/go/src/bytes/bytes.go:688
		// _ = "end of CoverTab[465]"
//line /usr/local/go/src/bytes/bytes.go:688
		_go_fuzz_dep_.CoverTab[466]++
							return b
//line /usr/local/go/src/bytes/bytes.go:689
		// _ = "end of CoverTab[466]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:690
		_go_fuzz_dep_.CoverTab[474]++
//line /usr/local/go/src/bytes/bytes.go:690
		// _ = "end of CoverTab[474]"
//line /usr/local/go/src/bytes/bytes.go:690
	}
//line /usr/local/go/src/bytes/bytes.go:690
	// _ = "end of CoverTab[456]"
//line /usr/local/go/src/bytes/bytes.go:690
	_go_fuzz_dep_.CoverTab[457]++
						return Map(unicode.ToLower, s)
//line /usr/local/go/src/bytes/bytes.go:691
	// _ = "end of CoverTab[457]"
}

// ToTitle treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their title case.
func ToTitle(s []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:695
	_go_fuzz_dep_.CoverTab[475]++
//line /usr/local/go/src/bytes/bytes.go:695
	return Map(unicode.ToTitle, s)
//line /usr/local/go/src/bytes/bytes.go:695
	// _ = "end of CoverTab[475]"
//line /usr/local/go/src/bytes/bytes.go:695
}

// ToUpperSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
//line /usr/local/go/src/bytes/bytes.go:697
// upper case, giving priority to the special casing rules.
//line /usr/local/go/src/bytes/bytes.go:699
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:699
	_go_fuzz_dep_.CoverTab[476]++
						return Map(c.ToUpper, s)
//line /usr/local/go/src/bytes/bytes.go:700
	// _ = "end of CoverTab[476]"
}

// ToLowerSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
//line /usr/local/go/src/bytes/bytes.go:703
// lower case, giving priority to the special casing rules.
//line /usr/local/go/src/bytes/bytes.go:705
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:705
	_go_fuzz_dep_.CoverTab[477]++
						return Map(c.ToLower, s)
//line /usr/local/go/src/bytes/bytes.go:706
	// _ = "end of CoverTab[477]"
}

// ToTitleSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
//line /usr/local/go/src/bytes/bytes.go:709
// title case, giving priority to the special casing rules.
//line /usr/local/go/src/bytes/bytes.go:711
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:711
	_go_fuzz_dep_.CoverTab[478]++
						return Map(c.ToTitle, s)
//line /usr/local/go/src/bytes/bytes.go:712
	// _ = "end of CoverTab[478]"
}

// ToValidUTF8 treats s as UTF-8-encoded bytes and returns a copy with each run of bytes
//line /usr/local/go/src/bytes/bytes.go:715
// representing invalid UTF-8 replaced with the bytes in replacement, which may be empty.
//line /usr/local/go/src/bytes/bytes.go:717
func ToValidUTF8(s, replacement []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:717
	_go_fuzz_dep_.CoverTab[479]++
						b := make([]byte, 0, len(s)+len(replacement))
						invalid := false
						for i := 0; i < len(s); {
//line /usr/local/go/src/bytes/bytes.go:720
		_go_fuzz_dep_.CoverTab[481]++
							c := s[i]
							if c < utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:722
			_go_fuzz_dep_.CoverTab[484]++
								i++
								invalid = false
								b = append(b, c)
								continue
//line /usr/local/go/src/bytes/bytes.go:726
			// _ = "end of CoverTab[484]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:727
			_go_fuzz_dep_.CoverTab[485]++
//line /usr/local/go/src/bytes/bytes.go:727
			// _ = "end of CoverTab[485]"
//line /usr/local/go/src/bytes/bytes.go:727
		}
//line /usr/local/go/src/bytes/bytes.go:727
		// _ = "end of CoverTab[481]"
//line /usr/local/go/src/bytes/bytes.go:727
		_go_fuzz_dep_.CoverTab[482]++
							_, wid := utf8.DecodeRune(s[i:])
							if wid == 1 {
//line /usr/local/go/src/bytes/bytes.go:729
			_go_fuzz_dep_.CoverTab[486]++
								i++
								if !invalid {
//line /usr/local/go/src/bytes/bytes.go:731
				_go_fuzz_dep_.CoverTab[488]++
									invalid = true
									b = append(b, replacement...)
//line /usr/local/go/src/bytes/bytes.go:733
				// _ = "end of CoverTab[488]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:734
				_go_fuzz_dep_.CoverTab[489]++
//line /usr/local/go/src/bytes/bytes.go:734
				// _ = "end of CoverTab[489]"
//line /usr/local/go/src/bytes/bytes.go:734
			}
//line /usr/local/go/src/bytes/bytes.go:734
			// _ = "end of CoverTab[486]"
//line /usr/local/go/src/bytes/bytes.go:734
			_go_fuzz_dep_.CoverTab[487]++
								continue
//line /usr/local/go/src/bytes/bytes.go:735
			// _ = "end of CoverTab[487]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:736
			_go_fuzz_dep_.CoverTab[490]++
//line /usr/local/go/src/bytes/bytes.go:736
			// _ = "end of CoverTab[490]"
//line /usr/local/go/src/bytes/bytes.go:736
		}
//line /usr/local/go/src/bytes/bytes.go:736
		// _ = "end of CoverTab[482]"
//line /usr/local/go/src/bytes/bytes.go:736
		_go_fuzz_dep_.CoverTab[483]++
							invalid = false
							b = append(b, s[i:i+wid]...)
							i += wid
//line /usr/local/go/src/bytes/bytes.go:739
		// _ = "end of CoverTab[483]"
	}
//line /usr/local/go/src/bytes/bytes.go:740
	// _ = "end of CoverTab[479]"
//line /usr/local/go/src/bytes/bytes.go:740
	_go_fuzz_dep_.CoverTab[480]++
						return b
//line /usr/local/go/src/bytes/bytes.go:741
	// _ = "end of CoverTab[480]"
}

// isSeparator reports whether the rune could mark a word boundary.
//line /usr/local/go/src/bytes/bytes.go:744
// TODO: update when package unicode captures more of the properties.
//line /usr/local/go/src/bytes/bytes.go:746
func isSeparator(r rune) bool {
//line /usr/local/go/src/bytes/bytes.go:746
	_go_fuzz_dep_.CoverTab[491]++

						if r <= 0x7F {
//line /usr/local/go/src/bytes/bytes.go:748
		_go_fuzz_dep_.CoverTab[494]++
							switch {
		case '0' <= r && func() bool {
//line /usr/local/go/src/bytes/bytes.go:750
			_go_fuzz_dep_.CoverTab[501]++
//line /usr/local/go/src/bytes/bytes.go:750
			return r <= '9'
//line /usr/local/go/src/bytes/bytes.go:750
			// _ = "end of CoverTab[501]"
//line /usr/local/go/src/bytes/bytes.go:750
		}():
//line /usr/local/go/src/bytes/bytes.go:750
			_go_fuzz_dep_.CoverTab[496]++
								return false
//line /usr/local/go/src/bytes/bytes.go:751
			// _ = "end of CoverTab[496]"
		case 'a' <= r && func() bool {
//line /usr/local/go/src/bytes/bytes.go:752
			_go_fuzz_dep_.CoverTab[502]++
//line /usr/local/go/src/bytes/bytes.go:752
			return r <= 'z'
//line /usr/local/go/src/bytes/bytes.go:752
			// _ = "end of CoverTab[502]"
//line /usr/local/go/src/bytes/bytes.go:752
		}():
//line /usr/local/go/src/bytes/bytes.go:752
			_go_fuzz_dep_.CoverTab[497]++
								return false
//line /usr/local/go/src/bytes/bytes.go:753
			// _ = "end of CoverTab[497]"
		case 'A' <= r && func() bool {
//line /usr/local/go/src/bytes/bytes.go:754
			_go_fuzz_dep_.CoverTab[503]++
//line /usr/local/go/src/bytes/bytes.go:754
			return r <= 'Z'
//line /usr/local/go/src/bytes/bytes.go:754
			// _ = "end of CoverTab[503]"
//line /usr/local/go/src/bytes/bytes.go:754
		}():
//line /usr/local/go/src/bytes/bytes.go:754
			_go_fuzz_dep_.CoverTab[498]++
								return false
//line /usr/local/go/src/bytes/bytes.go:755
			// _ = "end of CoverTab[498]"
		case r == '_':
//line /usr/local/go/src/bytes/bytes.go:756
			_go_fuzz_dep_.CoverTab[499]++
								return false
//line /usr/local/go/src/bytes/bytes.go:757
			// _ = "end of CoverTab[499]"
//line /usr/local/go/src/bytes/bytes.go:757
		default:
//line /usr/local/go/src/bytes/bytes.go:757
			_go_fuzz_dep_.CoverTab[500]++
//line /usr/local/go/src/bytes/bytes.go:757
			// _ = "end of CoverTab[500]"
		}
//line /usr/local/go/src/bytes/bytes.go:758
		// _ = "end of CoverTab[494]"
//line /usr/local/go/src/bytes/bytes.go:758
		_go_fuzz_dep_.CoverTab[495]++
							return true
//line /usr/local/go/src/bytes/bytes.go:759
		// _ = "end of CoverTab[495]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:760
		_go_fuzz_dep_.CoverTab[504]++
//line /usr/local/go/src/bytes/bytes.go:760
		// _ = "end of CoverTab[504]"
//line /usr/local/go/src/bytes/bytes.go:760
	}
//line /usr/local/go/src/bytes/bytes.go:760
	// _ = "end of CoverTab[491]"
//line /usr/local/go/src/bytes/bytes.go:760
	_go_fuzz_dep_.CoverTab[492]++

						if unicode.IsLetter(r) || func() bool {
//line /usr/local/go/src/bytes/bytes.go:762
		_go_fuzz_dep_.CoverTab[505]++
//line /usr/local/go/src/bytes/bytes.go:762
		return unicode.IsDigit(r)
//line /usr/local/go/src/bytes/bytes.go:762
		// _ = "end of CoverTab[505]"
//line /usr/local/go/src/bytes/bytes.go:762
	}() {
//line /usr/local/go/src/bytes/bytes.go:762
		_go_fuzz_dep_.CoverTab[506]++
							return false
//line /usr/local/go/src/bytes/bytes.go:763
		// _ = "end of CoverTab[506]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:764
		_go_fuzz_dep_.CoverTab[507]++
//line /usr/local/go/src/bytes/bytes.go:764
		// _ = "end of CoverTab[507]"
//line /usr/local/go/src/bytes/bytes.go:764
	}
//line /usr/local/go/src/bytes/bytes.go:764
	// _ = "end of CoverTab[492]"
//line /usr/local/go/src/bytes/bytes.go:764
	_go_fuzz_dep_.CoverTab[493]++

						return unicode.IsSpace(r)
//line /usr/local/go/src/bytes/bytes.go:766
	// _ = "end of CoverTab[493]"
}

// Title treats s as UTF-8-encoded bytes and returns a copy with all Unicode letters that begin
//line /usr/local/go/src/bytes/bytes.go:769
// words mapped to their title case.
//line /usr/local/go/src/bytes/bytes.go:769
//
//line /usr/local/go/src/bytes/bytes.go:769
// Deprecated: The rule Title uses for word boundaries does not handle Unicode
//line /usr/local/go/src/bytes/bytes.go:769
// punctuation properly. Use golang.org/x/text/cases instead.
//line /usr/local/go/src/bytes/bytes.go:774
func Title(s []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:774
	_go_fuzz_dep_.CoverTab[508]++

//line /usr/local/go/src/bytes/bytes.go:778
	prev := ' '
	return Map(
		func(r rune) rune {
//line /usr/local/go/src/bytes/bytes.go:780
			_go_fuzz_dep_.CoverTab[509]++
								if isSeparator(prev) {
//line /usr/local/go/src/bytes/bytes.go:781
				_go_fuzz_dep_.CoverTab[511]++
									prev = r
									return unicode.ToTitle(r)
//line /usr/local/go/src/bytes/bytes.go:783
				// _ = "end of CoverTab[511]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:784
				_go_fuzz_dep_.CoverTab[512]++
//line /usr/local/go/src/bytes/bytes.go:784
				// _ = "end of CoverTab[512]"
//line /usr/local/go/src/bytes/bytes.go:784
			}
//line /usr/local/go/src/bytes/bytes.go:784
			// _ = "end of CoverTab[509]"
//line /usr/local/go/src/bytes/bytes.go:784
			_go_fuzz_dep_.CoverTab[510]++
								prev = r
								return r
//line /usr/local/go/src/bytes/bytes.go:786
			// _ = "end of CoverTab[510]"
		},
		s)
//line /usr/local/go/src/bytes/bytes.go:788
	// _ = "end of CoverTab[508]"
}

// TrimLeftFunc treats s as UTF-8-encoded bytes and returns a subslice of s by slicing off
//line /usr/local/go/src/bytes/bytes.go:791
// all leading UTF-8-encoded code points c that satisfy f(c).
//line /usr/local/go/src/bytes/bytes.go:793
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte {
//line /usr/local/go/src/bytes/bytes.go:793
	_go_fuzz_dep_.CoverTab[513]++
						i := indexFunc(s, f, false)
						if i == -1 {
//line /usr/local/go/src/bytes/bytes.go:795
		_go_fuzz_dep_.CoverTab[515]++
							return nil
//line /usr/local/go/src/bytes/bytes.go:796
		// _ = "end of CoverTab[515]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:797
		_go_fuzz_dep_.CoverTab[516]++
//line /usr/local/go/src/bytes/bytes.go:797
		// _ = "end of CoverTab[516]"
//line /usr/local/go/src/bytes/bytes.go:797
	}
//line /usr/local/go/src/bytes/bytes.go:797
	// _ = "end of CoverTab[513]"
//line /usr/local/go/src/bytes/bytes.go:797
	_go_fuzz_dep_.CoverTab[514]++
						return s[i:]
//line /usr/local/go/src/bytes/bytes.go:798
	// _ = "end of CoverTab[514]"
}

// TrimRightFunc returns a subslice of s by slicing off all trailing
//line /usr/local/go/src/bytes/bytes.go:801
// UTF-8-encoded code points c that satisfy f(c).
//line /usr/local/go/src/bytes/bytes.go:803
func TrimRightFunc(s []byte, f func(r rune) bool) []byte {
//line /usr/local/go/src/bytes/bytes.go:803
	_go_fuzz_dep_.CoverTab[517]++
						i := lastIndexFunc(s, f, false)
						if i >= 0 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:805
		_go_fuzz_dep_.CoverTab[519]++
//line /usr/local/go/src/bytes/bytes.go:805
		return s[i] >= utf8.RuneSelf
//line /usr/local/go/src/bytes/bytes.go:805
		// _ = "end of CoverTab[519]"
//line /usr/local/go/src/bytes/bytes.go:805
	}() {
//line /usr/local/go/src/bytes/bytes.go:805
		_go_fuzz_dep_.CoverTab[520]++
							_, wid := utf8.DecodeRune(s[i:])
							i += wid
//line /usr/local/go/src/bytes/bytes.go:807
		// _ = "end of CoverTab[520]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:808
		_go_fuzz_dep_.CoverTab[521]++
							i++
//line /usr/local/go/src/bytes/bytes.go:809
		// _ = "end of CoverTab[521]"
	}
//line /usr/local/go/src/bytes/bytes.go:810
	// _ = "end of CoverTab[517]"
//line /usr/local/go/src/bytes/bytes.go:810
	_go_fuzz_dep_.CoverTab[518]++
						return s[0:i]
//line /usr/local/go/src/bytes/bytes.go:811
	// _ = "end of CoverTab[518]"
}

// TrimFunc returns a subslice of s by slicing off all leading and trailing
//line /usr/local/go/src/bytes/bytes.go:814
// UTF-8-encoded code points c that satisfy f(c).
//line /usr/local/go/src/bytes/bytes.go:816
func TrimFunc(s []byte, f func(r rune) bool) []byte {
//line /usr/local/go/src/bytes/bytes.go:816
	_go_fuzz_dep_.CoverTab[522]++
						return TrimRightFunc(TrimLeftFunc(s, f), f)
//line /usr/local/go/src/bytes/bytes.go:817
	// _ = "end of CoverTab[522]"
}

// TrimPrefix returns s without the provided leading prefix string.
//line /usr/local/go/src/bytes/bytes.go:820
// If s doesn't start with prefix, s is returned unchanged.
//line /usr/local/go/src/bytes/bytes.go:822
func TrimPrefix(s, prefix []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:822
	_go_fuzz_dep_.CoverTab[523]++
						if HasPrefix(s, prefix) {
//line /usr/local/go/src/bytes/bytes.go:823
		_go_fuzz_dep_.CoverTab[525]++
							return s[len(prefix):]
//line /usr/local/go/src/bytes/bytes.go:824
		// _ = "end of CoverTab[525]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:825
		_go_fuzz_dep_.CoverTab[526]++
//line /usr/local/go/src/bytes/bytes.go:825
		// _ = "end of CoverTab[526]"
//line /usr/local/go/src/bytes/bytes.go:825
	}
//line /usr/local/go/src/bytes/bytes.go:825
	// _ = "end of CoverTab[523]"
//line /usr/local/go/src/bytes/bytes.go:825
	_go_fuzz_dep_.CoverTab[524]++
						return s
//line /usr/local/go/src/bytes/bytes.go:826
	// _ = "end of CoverTab[524]"
}

// TrimSuffix returns s without the provided trailing suffix string.
//line /usr/local/go/src/bytes/bytes.go:829
// If s doesn't end with suffix, s is returned unchanged.
//line /usr/local/go/src/bytes/bytes.go:831
func TrimSuffix(s, suffix []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:831
	_go_fuzz_dep_.CoverTab[527]++
						if HasSuffix(s, suffix) {
//line /usr/local/go/src/bytes/bytes.go:832
		_go_fuzz_dep_.CoverTab[529]++
							return s[:len(s)-len(suffix)]
//line /usr/local/go/src/bytes/bytes.go:833
		// _ = "end of CoverTab[529]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:834
		_go_fuzz_dep_.CoverTab[530]++
//line /usr/local/go/src/bytes/bytes.go:834
		// _ = "end of CoverTab[530]"
//line /usr/local/go/src/bytes/bytes.go:834
	}
//line /usr/local/go/src/bytes/bytes.go:834
	// _ = "end of CoverTab[527]"
//line /usr/local/go/src/bytes/bytes.go:834
	_go_fuzz_dep_.CoverTab[528]++
						return s
//line /usr/local/go/src/bytes/bytes.go:835
	// _ = "end of CoverTab[528]"
}

// IndexFunc interprets s as a sequence of UTF-8-encoded code points.
//line /usr/local/go/src/bytes/bytes.go:838
// It returns the byte index in s of the first Unicode
//line /usr/local/go/src/bytes/bytes.go:838
// code point satisfying f(c), or -1 if none do.
//line /usr/local/go/src/bytes/bytes.go:841
func IndexFunc(s []byte, f func(r rune) bool) int {
//line /usr/local/go/src/bytes/bytes.go:841
	_go_fuzz_dep_.CoverTab[531]++
						return indexFunc(s, f, true)
//line /usr/local/go/src/bytes/bytes.go:842
	// _ = "end of CoverTab[531]"
}

// LastIndexFunc interprets s as a sequence of UTF-8-encoded code points.
//line /usr/local/go/src/bytes/bytes.go:845
// It returns the byte index in s of the last Unicode
//line /usr/local/go/src/bytes/bytes.go:845
// code point satisfying f(c), or -1 if none do.
//line /usr/local/go/src/bytes/bytes.go:848
func LastIndexFunc(s []byte, f func(r rune) bool) int {
//line /usr/local/go/src/bytes/bytes.go:848
	_go_fuzz_dep_.CoverTab[532]++
						return lastIndexFunc(s, f, true)
//line /usr/local/go/src/bytes/bytes.go:849
	// _ = "end of CoverTab[532]"
}

// indexFunc is the same as IndexFunc except that if
//line /usr/local/go/src/bytes/bytes.go:852
// truth==false, the sense of the predicate function is
//line /usr/local/go/src/bytes/bytes.go:852
// inverted.
//line /usr/local/go/src/bytes/bytes.go:855
func indexFunc(s []byte, f func(r rune) bool, truth bool) int {
//line /usr/local/go/src/bytes/bytes.go:855
	_go_fuzz_dep_.CoverTab[533]++
						start := 0
						for start < len(s) {
//line /usr/local/go/src/bytes/bytes.go:857
		_go_fuzz_dep_.CoverTab[535]++
							wid := 1
							r := rune(s[start])
							if r >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:860
			_go_fuzz_dep_.CoverTab[538]++
								r, wid = utf8.DecodeRune(s[start:])
//line /usr/local/go/src/bytes/bytes.go:861
			// _ = "end of CoverTab[538]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:862
			_go_fuzz_dep_.CoverTab[539]++
//line /usr/local/go/src/bytes/bytes.go:862
			// _ = "end of CoverTab[539]"
//line /usr/local/go/src/bytes/bytes.go:862
		}
//line /usr/local/go/src/bytes/bytes.go:862
		// _ = "end of CoverTab[535]"
//line /usr/local/go/src/bytes/bytes.go:862
		_go_fuzz_dep_.CoverTab[536]++
							if f(r) == truth {
//line /usr/local/go/src/bytes/bytes.go:863
			_go_fuzz_dep_.CoverTab[540]++
								return start
//line /usr/local/go/src/bytes/bytes.go:864
			// _ = "end of CoverTab[540]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:865
			_go_fuzz_dep_.CoverTab[541]++
//line /usr/local/go/src/bytes/bytes.go:865
			// _ = "end of CoverTab[541]"
//line /usr/local/go/src/bytes/bytes.go:865
		}
//line /usr/local/go/src/bytes/bytes.go:865
		// _ = "end of CoverTab[536]"
//line /usr/local/go/src/bytes/bytes.go:865
		_go_fuzz_dep_.CoverTab[537]++
							start += wid
//line /usr/local/go/src/bytes/bytes.go:866
		// _ = "end of CoverTab[537]"
	}
//line /usr/local/go/src/bytes/bytes.go:867
	// _ = "end of CoverTab[533]"
//line /usr/local/go/src/bytes/bytes.go:867
	_go_fuzz_dep_.CoverTab[534]++
						return -1
//line /usr/local/go/src/bytes/bytes.go:868
	// _ = "end of CoverTab[534]"
}

// lastIndexFunc is the same as LastIndexFunc except that if
//line /usr/local/go/src/bytes/bytes.go:871
// truth==false, the sense of the predicate function is
//line /usr/local/go/src/bytes/bytes.go:871
// inverted.
//line /usr/local/go/src/bytes/bytes.go:874
func lastIndexFunc(s []byte, f func(r rune) bool, truth bool) int {
//line /usr/local/go/src/bytes/bytes.go:874
	_go_fuzz_dep_.CoverTab[542]++
						for i := len(s); i > 0; {
//line /usr/local/go/src/bytes/bytes.go:875
		_go_fuzz_dep_.CoverTab[544]++
							r, size := rune(s[i-1]), 1
							if r >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:877
			_go_fuzz_dep_.CoverTab[546]++
								r, size = utf8.DecodeLastRune(s[0:i])
//line /usr/local/go/src/bytes/bytes.go:878
			// _ = "end of CoverTab[546]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:879
			_go_fuzz_dep_.CoverTab[547]++
//line /usr/local/go/src/bytes/bytes.go:879
			// _ = "end of CoverTab[547]"
//line /usr/local/go/src/bytes/bytes.go:879
		}
//line /usr/local/go/src/bytes/bytes.go:879
		// _ = "end of CoverTab[544]"
//line /usr/local/go/src/bytes/bytes.go:879
		_go_fuzz_dep_.CoverTab[545]++
							i -= size
							if f(r) == truth {
//line /usr/local/go/src/bytes/bytes.go:881
			_go_fuzz_dep_.CoverTab[548]++
								return i
//line /usr/local/go/src/bytes/bytes.go:882
			// _ = "end of CoverTab[548]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:883
			_go_fuzz_dep_.CoverTab[549]++
//line /usr/local/go/src/bytes/bytes.go:883
			// _ = "end of CoverTab[549]"
//line /usr/local/go/src/bytes/bytes.go:883
		}
//line /usr/local/go/src/bytes/bytes.go:883
		// _ = "end of CoverTab[545]"
	}
//line /usr/local/go/src/bytes/bytes.go:884
	// _ = "end of CoverTab[542]"
//line /usr/local/go/src/bytes/bytes.go:884
	_go_fuzz_dep_.CoverTab[543]++
						return -1
//line /usr/local/go/src/bytes/bytes.go:885
	// _ = "end of CoverTab[543]"
}

// asciiSet is a 32-byte value, where each bit represents the presence of a
//line /usr/local/go/src/bytes/bytes.go:888
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
//line /usr/local/go/src/bytes/bytes.go:888
// starting with the least-significant bit of the lowest word to the
//line /usr/local/go/src/bytes/bytes.go:888
// most-significant bit of the highest word, map to the full range of all
//line /usr/local/go/src/bytes/bytes.go:888
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
//line /usr/local/go/src/bytes/bytes.go:888
// ensuring that any non-ASCII character will be reported as not in the set.
//line /usr/local/go/src/bytes/bytes.go:888
// This allocates a total of 32 bytes even though the upper half
//line /usr/local/go/src/bytes/bytes.go:888
// is unused to avoid bounds checks in asciiSet.contains.
//line /usr/local/go/src/bytes/bytes.go:896
type asciiSet [8]uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
//line /usr/local/go/src/bytes/bytes.go:898
// characters in chars are ASCII.
//line /usr/local/go/src/bytes/bytes.go:900
func makeASCIISet(chars string) (as asciiSet, ok bool) {
//line /usr/local/go/src/bytes/bytes.go:900
	_go_fuzz_dep_.CoverTab[550]++
						for i := 0; i < len(chars); i++ {
//line /usr/local/go/src/bytes/bytes.go:901
		_go_fuzz_dep_.CoverTab[552]++
							c := chars[i]
							if c >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:903
			_go_fuzz_dep_.CoverTab[554]++
								return as, false
//line /usr/local/go/src/bytes/bytes.go:904
			// _ = "end of CoverTab[554]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:905
			_go_fuzz_dep_.CoverTab[555]++
//line /usr/local/go/src/bytes/bytes.go:905
			// _ = "end of CoverTab[555]"
//line /usr/local/go/src/bytes/bytes.go:905
		}
//line /usr/local/go/src/bytes/bytes.go:905
		// _ = "end of CoverTab[552]"
//line /usr/local/go/src/bytes/bytes.go:905
		_go_fuzz_dep_.CoverTab[553]++
							as[c/32] |= 1 << (c % 32)
//line /usr/local/go/src/bytes/bytes.go:906
		// _ = "end of CoverTab[553]"
	}
//line /usr/local/go/src/bytes/bytes.go:907
	// _ = "end of CoverTab[550]"
//line /usr/local/go/src/bytes/bytes.go:907
	_go_fuzz_dep_.CoverTab[551]++
						return as, true
//line /usr/local/go/src/bytes/bytes.go:908
	// _ = "end of CoverTab[551]"
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c byte) bool {
//line /usr/local/go/src/bytes/bytes.go:912
	_go_fuzz_dep_.CoverTab[556]++
						return (as[c/32] & (1 << (c % 32))) != 0
//line /usr/local/go/src/bytes/bytes.go:913
	// _ = "end of CoverTab[556]"
}

// containsRune is a simplified version of strings.ContainsRune
//line /usr/local/go/src/bytes/bytes.go:916
// to avoid importing the strings package.
//line /usr/local/go/src/bytes/bytes.go:916
// We avoid bytes.ContainsRune to avoid allocating a temporary copy of s.
//line /usr/local/go/src/bytes/bytes.go:919
func containsRune(s string, r rune) bool {
//line /usr/local/go/src/bytes/bytes.go:919
	_go_fuzz_dep_.CoverTab[557]++
						for _, c := range s {
//line /usr/local/go/src/bytes/bytes.go:920
		_go_fuzz_dep_.CoverTab[559]++
							if c == r {
//line /usr/local/go/src/bytes/bytes.go:921
			_go_fuzz_dep_.CoverTab[560]++
								return true
//line /usr/local/go/src/bytes/bytes.go:922
			// _ = "end of CoverTab[560]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:923
			_go_fuzz_dep_.CoverTab[561]++
//line /usr/local/go/src/bytes/bytes.go:923
			// _ = "end of CoverTab[561]"
//line /usr/local/go/src/bytes/bytes.go:923
		}
//line /usr/local/go/src/bytes/bytes.go:923
		// _ = "end of CoverTab[559]"
	}
//line /usr/local/go/src/bytes/bytes.go:924
	// _ = "end of CoverTab[557]"
//line /usr/local/go/src/bytes/bytes.go:924
	_go_fuzz_dep_.CoverTab[558]++
						return false
//line /usr/local/go/src/bytes/bytes.go:925
	// _ = "end of CoverTab[558]"
}

// Trim returns a subslice of s by slicing off all leading and
//line /usr/local/go/src/bytes/bytes.go:928
// trailing UTF-8-encoded code points contained in cutset.
//line /usr/local/go/src/bytes/bytes.go:930
func Trim(s []byte, cutset string) []byte {
//line /usr/local/go/src/bytes/bytes.go:930
	_go_fuzz_dep_.CoverTab[562]++
						if len(s) == 0 {
//line /usr/local/go/src/bytes/bytes.go:931
		_go_fuzz_dep_.CoverTab[567]++

							return nil
//line /usr/local/go/src/bytes/bytes.go:933
		// _ = "end of CoverTab[567]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:934
		_go_fuzz_dep_.CoverTab[568]++
//line /usr/local/go/src/bytes/bytes.go:934
		// _ = "end of CoverTab[568]"
//line /usr/local/go/src/bytes/bytes.go:934
	}
//line /usr/local/go/src/bytes/bytes.go:934
	// _ = "end of CoverTab[562]"
//line /usr/local/go/src/bytes/bytes.go:934
	_go_fuzz_dep_.CoverTab[563]++
						if cutset == "" {
//line /usr/local/go/src/bytes/bytes.go:935
		_go_fuzz_dep_.CoverTab[569]++
							return s
//line /usr/local/go/src/bytes/bytes.go:936
		// _ = "end of CoverTab[569]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:937
		_go_fuzz_dep_.CoverTab[570]++
//line /usr/local/go/src/bytes/bytes.go:937
		// _ = "end of CoverTab[570]"
//line /usr/local/go/src/bytes/bytes.go:937
	}
//line /usr/local/go/src/bytes/bytes.go:937
	// _ = "end of CoverTab[563]"
//line /usr/local/go/src/bytes/bytes.go:937
	_go_fuzz_dep_.CoverTab[564]++
						if len(cutset) == 1 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:938
		_go_fuzz_dep_.CoverTab[571]++
//line /usr/local/go/src/bytes/bytes.go:938
		return cutset[0] < utf8.RuneSelf
//line /usr/local/go/src/bytes/bytes.go:938
		// _ = "end of CoverTab[571]"
//line /usr/local/go/src/bytes/bytes.go:938
	}() {
//line /usr/local/go/src/bytes/bytes.go:938
		_go_fuzz_dep_.CoverTab[572]++
							return trimLeftByte(trimRightByte(s, cutset[0]), cutset[0])
//line /usr/local/go/src/bytes/bytes.go:939
		// _ = "end of CoverTab[572]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:940
		_go_fuzz_dep_.CoverTab[573]++
//line /usr/local/go/src/bytes/bytes.go:940
		// _ = "end of CoverTab[573]"
//line /usr/local/go/src/bytes/bytes.go:940
	}
//line /usr/local/go/src/bytes/bytes.go:940
	// _ = "end of CoverTab[564]"
//line /usr/local/go/src/bytes/bytes.go:940
	_go_fuzz_dep_.CoverTab[565]++
						if as, ok := makeASCIISet(cutset); ok {
//line /usr/local/go/src/bytes/bytes.go:941
		_go_fuzz_dep_.CoverTab[574]++
							return trimLeftASCII(trimRightASCII(s, &as), &as)
//line /usr/local/go/src/bytes/bytes.go:942
		// _ = "end of CoverTab[574]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:943
		_go_fuzz_dep_.CoverTab[575]++
//line /usr/local/go/src/bytes/bytes.go:943
		// _ = "end of CoverTab[575]"
//line /usr/local/go/src/bytes/bytes.go:943
	}
//line /usr/local/go/src/bytes/bytes.go:943
	// _ = "end of CoverTab[565]"
//line /usr/local/go/src/bytes/bytes.go:943
	_go_fuzz_dep_.CoverTab[566]++
						return trimLeftUnicode(trimRightUnicode(s, cutset), cutset)
//line /usr/local/go/src/bytes/bytes.go:944
	// _ = "end of CoverTab[566]"
}

// TrimLeft returns a subslice of s by slicing off all leading
//line /usr/local/go/src/bytes/bytes.go:947
// UTF-8-encoded code points contained in cutset.
//line /usr/local/go/src/bytes/bytes.go:949
func TrimLeft(s []byte, cutset string) []byte {
//line /usr/local/go/src/bytes/bytes.go:949
	_go_fuzz_dep_.CoverTab[576]++
						if len(s) == 0 {
//line /usr/local/go/src/bytes/bytes.go:950
		_go_fuzz_dep_.CoverTab[581]++

							return nil
//line /usr/local/go/src/bytes/bytes.go:952
		// _ = "end of CoverTab[581]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:953
		_go_fuzz_dep_.CoverTab[582]++
//line /usr/local/go/src/bytes/bytes.go:953
		// _ = "end of CoverTab[582]"
//line /usr/local/go/src/bytes/bytes.go:953
	}
//line /usr/local/go/src/bytes/bytes.go:953
	// _ = "end of CoverTab[576]"
//line /usr/local/go/src/bytes/bytes.go:953
	_go_fuzz_dep_.CoverTab[577]++
						if cutset == "" {
//line /usr/local/go/src/bytes/bytes.go:954
		_go_fuzz_dep_.CoverTab[583]++
							return s
//line /usr/local/go/src/bytes/bytes.go:955
		// _ = "end of CoverTab[583]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:956
		_go_fuzz_dep_.CoverTab[584]++
//line /usr/local/go/src/bytes/bytes.go:956
		// _ = "end of CoverTab[584]"
//line /usr/local/go/src/bytes/bytes.go:956
	}
//line /usr/local/go/src/bytes/bytes.go:956
	// _ = "end of CoverTab[577]"
//line /usr/local/go/src/bytes/bytes.go:956
	_go_fuzz_dep_.CoverTab[578]++
						if len(cutset) == 1 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:957
		_go_fuzz_dep_.CoverTab[585]++
//line /usr/local/go/src/bytes/bytes.go:957
		return cutset[0] < utf8.RuneSelf
//line /usr/local/go/src/bytes/bytes.go:957
		// _ = "end of CoverTab[585]"
//line /usr/local/go/src/bytes/bytes.go:957
	}() {
//line /usr/local/go/src/bytes/bytes.go:957
		_go_fuzz_dep_.CoverTab[586]++
							return trimLeftByte(s, cutset[0])
//line /usr/local/go/src/bytes/bytes.go:958
		// _ = "end of CoverTab[586]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:959
		_go_fuzz_dep_.CoverTab[587]++
//line /usr/local/go/src/bytes/bytes.go:959
		// _ = "end of CoverTab[587]"
//line /usr/local/go/src/bytes/bytes.go:959
	}
//line /usr/local/go/src/bytes/bytes.go:959
	// _ = "end of CoverTab[578]"
//line /usr/local/go/src/bytes/bytes.go:959
	_go_fuzz_dep_.CoverTab[579]++
						if as, ok := makeASCIISet(cutset); ok {
//line /usr/local/go/src/bytes/bytes.go:960
		_go_fuzz_dep_.CoverTab[588]++
							return trimLeftASCII(s, &as)
//line /usr/local/go/src/bytes/bytes.go:961
		// _ = "end of CoverTab[588]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:962
		_go_fuzz_dep_.CoverTab[589]++
//line /usr/local/go/src/bytes/bytes.go:962
		// _ = "end of CoverTab[589]"
//line /usr/local/go/src/bytes/bytes.go:962
	}
//line /usr/local/go/src/bytes/bytes.go:962
	// _ = "end of CoverTab[579]"
//line /usr/local/go/src/bytes/bytes.go:962
	_go_fuzz_dep_.CoverTab[580]++
						return trimLeftUnicode(s, cutset)
//line /usr/local/go/src/bytes/bytes.go:963
	// _ = "end of CoverTab[580]"
}

func trimLeftByte(s []byte, c byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:966
	_go_fuzz_dep_.CoverTab[590]++
						for len(s) > 0 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:967
		_go_fuzz_dep_.CoverTab[593]++
//line /usr/local/go/src/bytes/bytes.go:967
		return s[0] == c
//line /usr/local/go/src/bytes/bytes.go:967
		// _ = "end of CoverTab[593]"
//line /usr/local/go/src/bytes/bytes.go:967
	}() {
//line /usr/local/go/src/bytes/bytes.go:967
		_go_fuzz_dep_.CoverTab[594]++
							s = s[1:]
//line /usr/local/go/src/bytes/bytes.go:968
		// _ = "end of CoverTab[594]"
	}
//line /usr/local/go/src/bytes/bytes.go:969
	// _ = "end of CoverTab[590]"
//line /usr/local/go/src/bytes/bytes.go:969
	_go_fuzz_dep_.CoverTab[591]++
						if len(s) == 0 {
//line /usr/local/go/src/bytes/bytes.go:970
		_go_fuzz_dep_.CoverTab[595]++

							return nil
//line /usr/local/go/src/bytes/bytes.go:972
		// _ = "end of CoverTab[595]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:973
		_go_fuzz_dep_.CoverTab[596]++
//line /usr/local/go/src/bytes/bytes.go:973
		// _ = "end of CoverTab[596]"
//line /usr/local/go/src/bytes/bytes.go:973
	}
//line /usr/local/go/src/bytes/bytes.go:973
	// _ = "end of CoverTab[591]"
//line /usr/local/go/src/bytes/bytes.go:973
	_go_fuzz_dep_.CoverTab[592]++
						return s
//line /usr/local/go/src/bytes/bytes.go:974
	// _ = "end of CoverTab[592]"
}

func trimLeftASCII(s []byte, as *asciiSet) []byte {
//line /usr/local/go/src/bytes/bytes.go:977
	_go_fuzz_dep_.CoverTab[597]++
						for len(s) > 0 {
//line /usr/local/go/src/bytes/bytes.go:978
		_go_fuzz_dep_.CoverTab[600]++
							if !as.contains(s[0]) {
//line /usr/local/go/src/bytes/bytes.go:979
			_go_fuzz_dep_.CoverTab[602]++
								break
//line /usr/local/go/src/bytes/bytes.go:980
			// _ = "end of CoverTab[602]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:981
			_go_fuzz_dep_.CoverTab[603]++
//line /usr/local/go/src/bytes/bytes.go:981
			// _ = "end of CoverTab[603]"
//line /usr/local/go/src/bytes/bytes.go:981
		}
//line /usr/local/go/src/bytes/bytes.go:981
		// _ = "end of CoverTab[600]"
//line /usr/local/go/src/bytes/bytes.go:981
		_go_fuzz_dep_.CoverTab[601]++
							s = s[1:]
//line /usr/local/go/src/bytes/bytes.go:982
		// _ = "end of CoverTab[601]"
	}
//line /usr/local/go/src/bytes/bytes.go:983
	// _ = "end of CoverTab[597]"
//line /usr/local/go/src/bytes/bytes.go:983
	_go_fuzz_dep_.CoverTab[598]++
						if len(s) == 0 {
//line /usr/local/go/src/bytes/bytes.go:984
		_go_fuzz_dep_.CoverTab[604]++

							return nil
//line /usr/local/go/src/bytes/bytes.go:986
		// _ = "end of CoverTab[604]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:987
		_go_fuzz_dep_.CoverTab[605]++
//line /usr/local/go/src/bytes/bytes.go:987
		// _ = "end of CoverTab[605]"
//line /usr/local/go/src/bytes/bytes.go:987
	}
//line /usr/local/go/src/bytes/bytes.go:987
	// _ = "end of CoverTab[598]"
//line /usr/local/go/src/bytes/bytes.go:987
	_go_fuzz_dep_.CoverTab[599]++
						return s
//line /usr/local/go/src/bytes/bytes.go:988
	// _ = "end of CoverTab[599]"
}

func trimLeftUnicode(s []byte, cutset string) []byte {
//line /usr/local/go/src/bytes/bytes.go:991
	_go_fuzz_dep_.CoverTab[606]++
						for len(s) > 0 {
//line /usr/local/go/src/bytes/bytes.go:992
		_go_fuzz_dep_.CoverTab[609]++
							r, n := rune(s[0]), 1
							if r >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:994
			_go_fuzz_dep_.CoverTab[612]++
								r, n = utf8.DecodeRune(s)
//line /usr/local/go/src/bytes/bytes.go:995
			// _ = "end of CoverTab[612]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:996
			_go_fuzz_dep_.CoverTab[613]++
//line /usr/local/go/src/bytes/bytes.go:996
			// _ = "end of CoverTab[613]"
//line /usr/local/go/src/bytes/bytes.go:996
		}
//line /usr/local/go/src/bytes/bytes.go:996
		// _ = "end of CoverTab[609]"
//line /usr/local/go/src/bytes/bytes.go:996
		_go_fuzz_dep_.CoverTab[610]++
							if !containsRune(cutset, r) {
//line /usr/local/go/src/bytes/bytes.go:997
			_go_fuzz_dep_.CoverTab[614]++
								break
//line /usr/local/go/src/bytes/bytes.go:998
			// _ = "end of CoverTab[614]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:999
			_go_fuzz_dep_.CoverTab[615]++
//line /usr/local/go/src/bytes/bytes.go:999
			// _ = "end of CoverTab[615]"
//line /usr/local/go/src/bytes/bytes.go:999
		}
//line /usr/local/go/src/bytes/bytes.go:999
		// _ = "end of CoverTab[610]"
//line /usr/local/go/src/bytes/bytes.go:999
		_go_fuzz_dep_.CoverTab[611]++
							s = s[n:]
//line /usr/local/go/src/bytes/bytes.go:1000
		// _ = "end of CoverTab[611]"
	}
//line /usr/local/go/src/bytes/bytes.go:1001
	// _ = "end of CoverTab[606]"
//line /usr/local/go/src/bytes/bytes.go:1001
	_go_fuzz_dep_.CoverTab[607]++
						if len(s) == 0 {
//line /usr/local/go/src/bytes/bytes.go:1002
		_go_fuzz_dep_.CoverTab[616]++

							return nil
//line /usr/local/go/src/bytes/bytes.go:1004
		// _ = "end of CoverTab[616]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1005
		_go_fuzz_dep_.CoverTab[617]++
//line /usr/local/go/src/bytes/bytes.go:1005
		// _ = "end of CoverTab[617]"
//line /usr/local/go/src/bytes/bytes.go:1005
	}
//line /usr/local/go/src/bytes/bytes.go:1005
	// _ = "end of CoverTab[607]"
//line /usr/local/go/src/bytes/bytes.go:1005
	_go_fuzz_dep_.CoverTab[608]++
						return s
//line /usr/local/go/src/bytes/bytes.go:1006
	// _ = "end of CoverTab[608]"
}

// TrimRight returns a subslice of s by slicing off all trailing
//line /usr/local/go/src/bytes/bytes.go:1009
// UTF-8-encoded code points that are contained in cutset.
//line /usr/local/go/src/bytes/bytes.go:1011
func TrimRight(s []byte, cutset string) []byte {
//line /usr/local/go/src/bytes/bytes.go:1011
	_go_fuzz_dep_.CoverTab[618]++
						if len(s) == 0 || func() bool {
//line /usr/local/go/src/bytes/bytes.go:1012
		_go_fuzz_dep_.CoverTab[622]++
//line /usr/local/go/src/bytes/bytes.go:1012
		return cutset == ""
//line /usr/local/go/src/bytes/bytes.go:1012
		// _ = "end of CoverTab[622]"
//line /usr/local/go/src/bytes/bytes.go:1012
	}() {
//line /usr/local/go/src/bytes/bytes.go:1012
		_go_fuzz_dep_.CoverTab[623]++
							return s
//line /usr/local/go/src/bytes/bytes.go:1013
		// _ = "end of CoverTab[623]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1014
		_go_fuzz_dep_.CoverTab[624]++
//line /usr/local/go/src/bytes/bytes.go:1014
		// _ = "end of CoverTab[624]"
//line /usr/local/go/src/bytes/bytes.go:1014
	}
//line /usr/local/go/src/bytes/bytes.go:1014
	// _ = "end of CoverTab[618]"
//line /usr/local/go/src/bytes/bytes.go:1014
	_go_fuzz_dep_.CoverTab[619]++
						if len(cutset) == 1 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1015
		_go_fuzz_dep_.CoverTab[625]++
//line /usr/local/go/src/bytes/bytes.go:1015
		return cutset[0] < utf8.RuneSelf
//line /usr/local/go/src/bytes/bytes.go:1015
		// _ = "end of CoverTab[625]"
//line /usr/local/go/src/bytes/bytes.go:1015
	}() {
//line /usr/local/go/src/bytes/bytes.go:1015
		_go_fuzz_dep_.CoverTab[626]++
							return trimRightByte(s, cutset[0])
//line /usr/local/go/src/bytes/bytes.go:1016
		// _ = "end of CoverTab[626]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1017
		_go_fuzz_dep_.CoverTab[627]++
//line /usr/local/go/src/bytes/bytes.go:1017
		// _ = "end of CoverTab[627]"
//line /usr/local/go/src/bytes/bytes.go:1017
	}
//line /usr/local/go/src/bytes/bytes.go:1017
	// _ = "end of CoverTab[619]"
//line /usr/local/go/src/bytes/bytes.go:1017
	_go_fuzz_dep_.CoverTab[620]++
						if as, ok := makeASCIISet(cutset); ok {
//line /usr/local/go/src/bytes/bytes.go:1018
		_go_fuzz_dep_.CoverTab[628]++
							return trimRightASCII(s, &as)
//line /usr/local/go/src/bytes/bytes.go:1019
		// _ = "end of CoverTab[628]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1020
		_go_fuzz_dep_.CoverTab[629]++
//line /usr/local/go/src/bytes/bytes.go:1020
		// _ = "end of CoverTab[629]"
//line /usr/local/go/src/bytes/bytes.go:1020
	}
//line /usr/local/go/src/bytes/bytes.go:1020
	// _ = "end of CoverTab[620]"
//line /usr/local/go/src/bytes/bytes.go:1020
	_go_fuzz_dep_.CoverTab[621]++
						return trimRightUnicode(s, cutset)
//line /usr/local/go/src/bytes/bytes.go:1021
	// _ = "end of CoverTab[621]"
}

func trimRightByte(s []byte, c byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:1024
	_go_fuzz_dep_.CoverTab[630]++
						for len(s) > 0 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1025
		_go_fuzz_dep_.CoverTab[632]++
//line /usr/local/go/src/bytes/bytes.go:1025
		return s[len(s)-1] == c
//line /usr/local/go/src/bytes/bytes.go:1025
		// _ = "end of CoverTab[632]"
//line /usr/local/go/src/bytes/bytes.go:1025
	}() {
//line /usr/local/go/src/bytes/bytes.go:1025
		_go_fuzz_dep_.CoverTab[633]++
							s = s[:len(s)-1]
//line /usr/local/go/src/bytes/bytes.go:1026
		// _ = "end of CoverTab[633]"
	}
//line /usr/local/go/src/bytes/bytes.go:1027
	// _ = "end of CoverTab[630]"
//line /usr/local/go/src/bytes/bytes.go:1027
	_go_fuzz_dep_.CoverTab[631]++
						return s
//line /usr/local/go/src/bytes/bytes.go:1028
	// _ = "end of CoverTab[631]"
}

func trimRightASCII(s []byte, as *asciiSet) []byte {
//line /usr/local/go/src/bytes/bytes.go:1031
	_go_fuzz_dep_.CoverTab[634]++
						for len(s) > 0 {
//line /usr/local/go/src/bytes/bytes.go:1032
		_go_fuzz_dep_.CoverTab[636]++
							if !as.contains(s[len(s)-1]) {
//line /usr/local/go/src/bytes/bytes.go:1033
			_go_fuzz_dep_.CoverTab[638]++
								break
//line /usr/local/go/src/bytes/bytes.go:1034
			// _ = "end of CoverTab[638]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1035
			_go_fuzz_dep_.CoverTab[639]++
//line /usr/local/go/src/bytes/bytes.go:1035
			// _ = "end of CoverTab[639]"
//line /usr/local/go/src/bytes/bytes.go:1035
		}
//line /usr/local/go/src/bytes/bytes.go:1035
		// _ = "end of CoverTab[636]"
//line /usr/local/go/src/bytes/bytes.go:1035
		_go_fuzz_dep_.CoverTab[637]++
							s = s[:len(s)-1]
//line /usr/local/go/src/bytes/bytes.go:1036
		// _ = "end of CoverTab[637]"
	}
//line /usr/local/go/src/bytes/bytes.go:1037
	// _ = "end of CoverTab[634]"
//line /usr/local/go/src/bytes/bytes.go:1037
	_go_fuzz_dep_.CoverTab[635]++
						return s
//line /usr/local/go/src/bytes/bytes.go:1038
	// _ = "end of CoverTab[635]"
}

func trimRightUnicode(s []byte, cutset string) []byte {
//line /usr/local/go/src/bytes/bytes.go:1041
	_go_fuzz_dep_.CoverTab[640]++
						for len(s) > 0 {
//line /usr/local/go/src/bytes/bytes.go:1042
		_go_fuzz_dep_.CoverTab[642]++
							r, n := rune(s[len(s)-1]), 1
							if r >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:1044
			_go_fuzz_dep_.CoverTab[645]++
								r, n = utf8.DecodeLastRune(s)
//line /usr/local/go/src/bytes/bytes.go:1045
			// _ = "end of CoverTab[645]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1046
			_go_fuzz_dep_.CoverTab[646]++
//line /usr/local/go/src/bytes/bytes.go:1046
			// _ = "end of CoverTab[646]"
//line /usr/local/go/src/bytes/bytes.go:1046
		}
//line /usr/local/go/src/bytes/bytes.go:1046
		// _ = "end of CoverTab[642]"
//line /usr/local/go/src/bytes/bytes.go:1046
		_go_fuzz_dep_.CoverTab[643]++
							if !containsRune(cutset, r) {
//line /usr/local/go/src/bytes/bytes.go:1047
			_go_fuzz_dep_.CoverTab[647]++
								break
//line /usr/local/go/src/bytes/bytes.go:1048
			// _ = "end of CoverTab[647]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1049
			_go_fuzz_dep_.CoverTab[648]++
//line /usr/local/go/src/bytes/bytes.go:1049
			// _ = "end of CoverTab[648]"
//line /usr/local/go/src/bytes/bytes.go:1049
		}
//line /usr/local/go/src/bytes/bytes.go:1049
		// _ = "end of CoverTab[643]"
//line /usr/local/go/src/bytes/bytes.go:1049
		_go_fuzz_dep_.CoverTab[644]++
							s = s[:len(s)-n]
//line /usr/local/go/src/bytes/bytes.go:1050
		// _ = "end of CoverTab[644]"
	}
//line /usr/local/go/src/bytes/bytes.go:1051
	// _ = "end of CoverTab[640]"
//line /usr/local/go/src/bytes/bytes.go:1051
	_go_fuzz_dep_.CoverTab[641]++
						return s
//line /usr/local/go/src/bytes/bytes.go:1052
	// _ = "end of CoverTab[641]"
}

// TrimSpace returns a subslice of s by slicing off all leading and
//line /usr/local/go/src/bytes/bytes.go:1055
// trailing white space, as defined by Unicode.
//line /usr/local/go/src/bytes/bytes.go:1057
func TrimSpace(s []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:1057
	_go_fuzz_dep_.CoverTab[649]++

						start := 0
						for ; start < len(s); start++ {
//line /usr/local/go/src/bytes/bytes.go:1060
		_go_fuzz_dep_.CoverTab[653]++
							c := s[start]
							if c >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:1062
			_go_fuzz_dep_.CoverTab[655]++

//line /usr/local/go/src/bytes/bytes.go:1065
			return TrimFunc(s[start:], unicode.IsSpace)
//line /usr/local/go/src/bytes/bytes.go:1065
			// _ = "end of CoverTab[655]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1066
			_go_fuzz_dep_.CoverTab[656]++
//line /usr/local/go/src/bytes/bytes.go:1066
			// _ = "end of CoverTab[656]"
//line /usr/local/go/src/bytes/bytes.go:1066
		}
//line /usr/local/go/src/bytes/bytes.go:1066
		// _ = "end of CoverTab[653]"
//line /usr/local/go/src/bytes/bytes.go:1066
		_go_fuzz_dep_.CoverTab[654]++
							if asciiSpace[c] == 0 {
//line /usr/local/go/src/bytes/bytes.go:1067
			_go_fuzz_dep_.CoverTab[657]++
								break
//line /usr/local/go/src/bytes/bytes.go:1068
			// _ = "end of CoverTab[657]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1069
			_go_fuzz_dep_.CoverTab[658]++
//line /usr/local/go/src/bytes/bytes.go:1069
			// _ = "end of CoverTab[658]"
//line /usr/local/go/src/bytes/bytes.go:1069
		}
//line /usr/local/go/src/bytes/bytes.go:1069
		// _ = "end of CoverTab[654]"
	}
//line /usr/local/go/src/bytes/bytes.go:1070
	// _ = "end of CoverTab[649]"
//line /usr/local/go/src/bytes/bytes.go:1070
	_go_fuzz_dep_.CoverTab[650]++

//line /usr/local/go/src/bytes/bytes.go:1073
	stop := len(s)
	for ; stop > start; stop-- {
//line /usr/local/go/src/bytes/bytes.go:1074
		_go_fuzz_dep_.CoverTab[659]++
							c := s[stop-1]
							if c >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:1076
			_go_fuzz_dep_.CoverTab[661]++
								return TrimFunc(s[start:stop], unicode.IsSpace)
//line /usr/local/go/src/bytes/bytes.go:1077
			// _ = "end of CoverTab[661]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1078
			_go_fuzz_dep_.CoverTab[662]++
//line /usr/local/go/src/bytes/bytes.go:1078
			// _ = "end of CoverTab[662]"
//line /usr/local/go/src/bytes/bytes.go:1078
		}
//line /usr/local/go/src/bytes/bytes.go:1078
		// _ = "end of CoverTab[659]"
//line /usr/local/go/src/bytes/bytes.go:1078
		_go_fuzz_dep_.CoverTab[660]++
							if asciiSpace[c] == 0 {
//line /usr/local/go/src/bytes/bytes.go:1079
			_go_fuzz_dep_.CoverTab[663]++
								break
//line /usr/local/go/src/bytes/bytes.go:1080
			// _ = "end of CoverTab[663]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1081
			_go_fuzz_dep_.CoverTab[664]++
//line /usr/local/go/src/bytes/bytes.go:1081
			// _ = "end of CoverTab[664]"
//line /usr/local/go/src/bytes/bytes.go:1081
		}
//line /usr/local/go/src/bytes/bytes.go:1081
		// _ = "end of CoverTab[660]"
	}
//line /usr/local/go/src/bytes/bytes.go:1082
	// _ = "end of CoverTab[650]"
//line /usr/local/go/src/bytes/bytes.go:1082
	_go_fuzz_dep_.CoverTab[651]++

//line /usr/local/go/src/bytes/bytes.go:1087
	if start == stop {
//line /usr/local/go/src/bytes/bytes.go:1087
		_go_fuzz_dep_.CoverTab[665]++

//line /usr/local/go/src/bytes/bytes.go:1090
		return nil
//line /usr/local/go/src/bytes/bytes.go:1090
		// _ = "end of CoverTab[665]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1091
		_go_fuzz_dep_.CoverTab[666]++
//line /usr/local/go/src/bytes/bytes.go:1091
		// _ = "end of CoverTab[666]"
//line /usr/local/go/src/bytes/bytes.go:1091
	}
//line /usr/local/go/src/bytes/bytes.go:1091
	// _ = "end of CoverTab[651]"
//line /usr/local/go/src/bytes/bytes.go:1091
	_go_fuzz_dep_.CoverTab[652]++
						return s[start:stop]
//line /usr/local/go/src/bytes/bytes.go:1092
	// _ = "end of CoverTab[652]"
}

// Runes interprets s as a sequence of UTF-8-encoded code points.
//line /usr/local/go/src/bytes/bytes.go:1095
// It returns a slice of runes (Unicode code points) equivalent to s.
//line /usr/local/go/src/bytes/bytes.go:1097
func Runes(s []byte) []rune {
//line /usr/local/go/src/bytes/bytes.go:1097
	_go_fuzz_dep_.CoverTab[667]++
						t := make([]rune, utf8.RuneCount(s))
						i := 0
						for len(s) > 0 {
//line /usr/local/go/src/bytes/bytes.go:1100
		_go_fuzz_dep_.CoverTab[669]++
							r, l := utf8.DecodeRune(s)
							t[i] = r
							i++
							s = s[l:]
//line /usr/local/go/src/bytes/bytes.go:1104
		// _ = "end of CoverTab[669]"
	}
//line /usr/local/go/src/bytes/bytes.go:1105
	// _ = "end of CoverTab[667]"
//line /usr/local/go/src/bytes/bytes.go:1105
	_go_fuzz_dep_.CoverTab[668]++
						return t
//line /usr/local/go/src/bytes/bytes.go:1106
	// _ = "end of CoverTab[668]"
}

// Replace returns a copy of the slice s with the first n
//line /usr/local/go/src/bytes/bytes.go:1109
// non-overlapping instances of old replaced by new.
//line /usr/local/go/src/bytes/bytes.go:1109
// If old is empty, it matches at the beginning of the slice
//line /usr/local/go/src/bytes/bytes.go:1109
// and after each UTF-8 sequence, yielding up to k+1 replacements
//line /usr/local/go/src/bytes/bytes.go:1109
// for a k-rune slice.
//line /usr/local/go/src/bytes/bytes.go:1109
// If n < 0, there is no limit on the number of replacements.
//line /usr/local/go/src/bytes/bytes.go:1115
func Replace(s, old, new []byte, n int) []byte {
//line /usr/local/go/src/bytes/bytes.go:1115
	_go_fuzz_dep_.CoverTab[670]++
						m := 0
						if n != 0 {
//line /usr/local/go/src/bytes/bytes.go:1117
		_go_fuzz_dep_.CoverTab[675]++

							m = Count(s, old)
//line /usr/local/go/src/bytes/bytes.go:1119
		// _ = "end of CoverTab[675]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1120
		_go_fuzz_dep_.CoverTab[676]++
//line /usr/local/go/src/bytes/bytes.go:1120
		// _ = "end of CoverTab[676]"
//line /usr/local/go/src/bytes/bytes.go:1120
	}
//line /usr/local/go/src/bytes/bytes.go:1120
	// _ = "end of CoverTab[670]"
//line /usr/local/go/src/bytes/bytes.go:1120
	_go_fuzz_dep_.CoverTab[671]++
						if m == 0 {
//line /usr/local/go/src/bytes/bytes.go:1121
		_go_fuzz_dep_.CoverTab[677]++

							return append([]byte(nil), s...)
//line /usr/local/go/src/bytes/bytes.go:1123
		// _ = "end of CoverTab[677]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1124
		_go_fuzz_dep_.CoverTab[678]++
//line /usr/local/go/src/bytes/bytes.go:1124
		// _ = "end of CoverTab[678]"
//line /usr/local/go/src/bytes/bytes.go:1124
	}
//line /usr/local/go/src/bytes/bytes.go:1124
	// _ = "end of CoverTab[671]"
//line /usr/local/go/src/bytes/bytes.go:1124
	_go_fuzz_dep_.CoverTab[672]++
						if n < 0 || func() bool {
//line /usr/local/go/src/bytes/bytes.go:1125
		_go_fuzz_dep_.CoverTab[679]++
//line /usr/local/go/src/bytes/bytes.go:1125
		return m < n
//line /usr/local/go/src/bytes/bytes.go:1125
		// _ = "end of CoverTab[679]"
//line /usr/local/go/src/bytes/bytes.go:1125
	}() {
//line /usr/local/go/src/bytes/bytes.go:1125
		_go_fuzz_dep_.CoverTab[680]++
							n = m
//line /usr/local/go/src/bytes/bytes.go:1126
		// _ = "end of CoverTab[680]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1127
		_go_fuzz_dep_.CoverTab[681]++
//line /usr/local/go/src/bytes/bytes.go:1127
		// _ = "end of CoverTab[681]"
//line /usr/local/go/src/bytes/bytes.go:1127
	}
//line /usr/local/go/src/bytes/bytes.go:1127
	// _ = "end of CoverTab[672]"
//line /usr/local/go/src/bytes/bytes.go:1127
	_go_fuzz_dep_.CoverTab[673]++

//line /usr/local/go/src/bytes/bytes.go:1130
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
//line /usr/local/go/src/bytes/bytes.go:1133
		_go_fuzz_dep_.CoverTab[682]++
							j := start
							if len(old) == 0 {
//line /usr/local/go/src/bytes/bytes.go:1135
			_go_fuzz_dep_.CoverTab[684]++
								if i > 0 {
//line /usr/local/go/src/bytes/bytes.go:1136
				_go_fuzz_dep_.CoverTab[685]++
									_, wid := utf8.DecodeRune(s[start:])
									j += wid
//line /usr/local/go/src/bytes/bytes.go:1138
				// _ = "end of CoverTab[685]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:1139
				_go_fuzz_dep_.CoverTab[686]++
//line /usr/local/go/src/bytes/bytes.go:1139
				// _ = "end of CoverTab[686]"
//line /usr/local/go/src/bytes/bytes.go:1139
			}
//line /usr/local/go/src/bytes/bytes.go:1139
			// _ = "end of CoverTab[684]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1140
			_go_fuzz_dep_.CoverTab[687]++
								j += Index(s[start:], old)
//line /usr/local/go/src/bytes/bytes.go:1141
			// _ = "end of CoverTab[687]"
		}
//line /usr/local/go/src/bytes/bytes.go:1142
		// _ = "end of CoverTab[682]"
//line /usr/local/go/src/bytes/bytes.go:1142
		_go_fuzz_dep_.CoverTab[683]++
							w += copy(t[w:], s[start:j])
							w += copy(t[w:], new)
							start = j + len(old)
//line /usr/local/go/src/bytes/bytes.go:1145
		// _ = "end of CoverTab[683]"
	}
//line /usr/local/go/src/bytes/bytes.go:1146
	// _ = "end of CoverTab[673]"
//line /usr/local/go/src/bytes/bytes.go:1146
	_go_fuzz_dep_.CoverTab[674]++
						w += copy(t[w:], s[start:])
						return t[0:w]
//line /usr/local/go/src/bytes/bytes.go:1148
	// _ = "end of CoverTab[674]"
}

// ReplaceAll returns a copy of the slice s with all
//line /usr/local/go/src/bytes/bytes.go:1151
// non-overlapping instances of old replaced by new.
//line /usr/local/go/src/bytes/bytes.go:1151
// If old is empty, it matches at the beginning of the slice
//line /usr/local/go/src/bytes/bytes.go:1151
// and after each UTF-8 sequence, yielding up to k+1 replacements
//line /usr/local/go/src/bytes/bytes.go:1151
// for a k-rune slice.
//line /usr/local/go/src/bytes/bytes.go:1156
func ReplaceAll(s, old, new []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:1156
	_go_fuzz_dep_.CoverTab[688]++
						return Replace(s, old, new, -1)
//line /usr/local/go/src/bytes/bytes.go:1157
	// _ = "end of CoverTab[688]"
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
//line /usr/local/go/src/bytes/bytes.go:1160
// are equal under simple Unicode case-folding, which is a more general
//line /usr/local/go/src/bytes/bytes.go:1160
// form of case-insensitivity.
//line /usr/local/go/src/bytes/bytes.go:1163
func EqualFold(s, t []byte) bool {
//line /usr/local/go/src/bytes/bytes.go:1163
	_go_fuzz_dep_.CoverTab[689]++

						i := 0
						for ; i < len(s) && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1166
		_go_fuzz_dep_.CoverTab[692]++
//line /usr/local/go/src/bytes/bytes.go:1166
		return i < len(t)
//line /usr/local/go/src/bytes/bytes.go:1166
		// _ = "end of CoverTab[692]"
//line /usr/local/go/src/bytes/bytes.go:1166
	}(); i++ {
//line /usr/local/go/src/bytes/bytes.go:1166
		_go_fuzz_dep_.CoverTab[693]++
							sr := s[i]
							tr := t[i]
							if sr|tr >= utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:1169
			_go_fuzz_dep_.CoverTab[698]++
								goto hasUnicode
//line /usr/local/go/src/bytes/bytes.go:1170
			// _ = "end of CoverTab[698]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1171
			_go_fuzz_dep_.CoverTab[699]++
//line /usr/local/go/src/bytes/bytes.go:1171
			// _ = "end of CoverTab[699]"
//line /usr/local/go/src/bytes/bytes.go:1171
		}
//line /usr/local/go/src/bytes/bytes.go:1171
		// _ = "end of CoverTab[693]"
//line /usr/local/go/src/bytes/bytes.go:1171
		_go_fuzz_dep_.CoverTab[694]++

//line /usr/local/go/src/bytes/bytes.go:1174
		if tr == sr {
//line /usr/local/go/src/bytes/bytes.go:1174
			_go_fuzz_dep_.CoverTab[700]++
								continue
//line /usr/local/go/src/bytes/bytes.go:1175
			// _ = "end of CoverTab[700]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1176
			_go_fuzz_dep_.CoverTab[701]++
//line /usr/local/go/src/bytes/bytes.go:1176
			// _ = "end of CoverTab[701]"
//line /usr/local/go/src/bytes/bytes.go:1176
		}
//line /usr/local/go/src/bytes/bytes.go:1176
		// _ = "end of CoverTab[694]"
//line /usr/local/go/src/bytes/bytes.go:1176
		_go_fuzz_dep_.CoverTab[695]++

//line /usr/local/go/src/bytes/bytes.go:1179
		if tr < sr {
//line /usr/local/go/src/bytes/bytes.go:1179
			_go_fuzz_dep_.CoverTab[702]++
								tr, sr = sr, tr
//line /usr/local/go/src/bytes/bytes.go:1180
			// _ = "end of CoverTab[702]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1181
			_go_fuzz_dep_.CoverTab[703]++
//line /usr/local/go/src/bytes/bytes.go:1181
			// _ = "end of CoverTab[703]"
//line /usr/local/go/src/bytes/bytes.go:1181
		}
//line /usr/local/go/src/bytes/bytes.go:1181
		// _ = "end of CoverTab[695]"
//line /usr/local/go/src/bytes/bytes.go:1181
		_go_fuzz_dep_.CoverTab[696]++

							if 'A' <= sr && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1183
			_go_fuzz_dep_.CoverTab[704]++
//line /usr/local/go/src/bytes/bytes.go:1183
			return sr <= 'Z'
//line /usr/local/go/src/bytes/bytes.go:1183
			// _ = "end of CoverTab[704]"
//line /usr/local/go/src/bytes/bytes.go:1183
		}() && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1183
			_go_fuzz_dep_.CoverTab[705]++
//line /usr/local/go/src/bytes/bytes.go:1183
			return tr == sr+'a'-'A'
//line /usr/local/go/src/bytes/bytes.go:1183
			// _ = "end of CoverTab[705]"
//line /usr/local/go/src/bytes/bytes.go:1183
		}() {
//line /usr/local/go/src/bytes/bytes.go:1183
			_go_fuzz_dep_.CoverTab[706]++
								continue
//line /usr/local/go/src/bytes/bytes.go:1184
			// _ = "end of CoverTab[706]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1185
			_go_fuzz_dep_.CoverTab[707]++
//line /usr/local/go/src/bytes/bytes.go:1185
			// _ = "end of CoverTab[707]"
//line /usr/local/go/src/bytes/bytes.go:1185
		}
//line /usr/local/go/src/bytes/bytes.go:1185
		// _ = "end of CoverTab[696]"
//line /usr/local/go/src/bytes/bytes.go:1185
		_go_fuzz_dep_.CoverTab[697]++
							return false
//line /usr/local/go/src/bytes/bytes.go:1186
		// _ = "end of CoverTab[697]"
	}
//line /usr/local/go/src/bytes/bytes.go:1187
	// _ = "end of CoverTab[689]"
//line /usr/local/go/src/bytes/bytes.go:1187
	_go_fuzz_dep_.CoverTab[690]++

						return len(s) == len(t)

hasUnicode:
	s = s[i:]
	t = t[i:]
	for len(s) != 0 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1194
		_go_fuzz_dep_.CoverTab[708]++
//line /usr/local/go/src/bytes/bytes.go:1194
		return len(t) != 0
//line /usr/local/go/src/bytes/bytes.go:1194
		// _ = "end of CoverTab[708]"
//line /usr/local/go/src/bytes/bytes.go:1194
	}() {
//line /usr/local/go/src/bytes/bytes.go:1194
		_go_fuzz_dep_.CoverTab[709]++
		// Extract first rune from each.
		var sr, tr rune
		if s[0] < utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:1197
			_go_fuzz_dep_.CoverTab[717]++
								sr, s = rune(s[0]), s[1:]
//line /usr/local/go/src/bytes/bytes.go:1198
			// _ = "end of CoverTab[717]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1199
			_go_fuzz_dep_.CoverTab[718]++
								r, size := utf8.DecodeRune(s)
								sr, s = r, s[size:]
//line /usr/local/go/src/bytes/bytes.go:1201
			// _ = "end of CoverTab[718]"
		}
//line /usr/local/go/src/bytes/bytes.go:1202
		// _ = "end of CoverTab[709]"
//line /usr/local/go/src/bytes/bytes.go:1202
		_go_fuzz_dep_.CoverTab[710]++
							if t[0] < utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:1203
			_go_fuzz_dep_.CoverTab[719]++
								tr, t = rune(t[0]), t[1:]
//line /usr/local/go/src/bytes/bytes.go:1204
			// _ = "end of CoverTab[719]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1205
			_go_fuzz_dep_.CoverTab[720]++
								r, size := utf8.DecodeRune(t)
								tr, t = r, t[size:]
//line /usr/local/go/src/bytes/bytes.go:1207
			// _ = "end of CoverTab[720]"
		}
//line /usr/local/go/src/bytes/bytes.go:1208
		// _ = "end of CoverTab[710]"
//line /usr/local/go/src/bytes/bytes.go:1208
		_go_fuzz_dep_.CoverTab[711]++

//line /usr/local/go/src/bytes/bytes.go:1213
		if tr == sr {
//line /usr/local/go/src/bytes/bytes.go:1213
			_go_fuzz_dep_.CoverTab[721]++
								continue
//line /usr/local/go/src/bytes/bytes.go:1214
			// _ = "end of CoverTab[721]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1215
			_go_fuzz_dep_.CoverTab[722]++
//line /usr/local/go/src/bytes/bytes.go:1215
			// _ = "end of CoverTab[722]"
//line /usr/local/go/src/bytes/bytes.go:1215
		}
//line /usr/local/go/src/bytes/bytes.go:1215
		// _ = "end of CoverTab[711]"
//line /usr/local/go/src/bytes/bytes.go:1215
		_go_fuzz_dep_.CoverTab[712]++

//line /usr/local/go/src/bytes/bytes.go:1218
		if tr < sr {
//line /usr/local/go/src/bytes/bytes.go:1218
			_go_fuzz_dep_.CoverTab[723]++
								tr, sr = sr, tr
//line /usr/local/go/src/bytes/bytes.go:1219
			// _ = "end of CoverTab[723]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1220
			_go_fuzz_dep_.CoverTab[724]++
//line /usr/local/go/src/bytes/bytes.go:1220
			// _ = "end of CoverTab[724]"
//line /usr/local/go/src/bytes/bytes.go:1220
		}
//line /usr/local/go/src/bytes/bytes.go:1220
		// _ = "end of CoverTab[712]"
//line /usr/local/go/src/bytes/bytes.go:1220
		_go_fuzz_dep_.CoverTab[713]++

							if tr < utf8.RuneSelf {
//line /usr/local/go/src/bytes/bytes.go:1222
			_go_fuzz_dep_.CoverTab[725]++

								if 'A' <= sr && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1224
				_go_fuzz_dep_.CoverTab[727]++
//line /usr/local/go/src/bytes/bytes.go:1224
				return sr <= 'Z'
//line /usr/local/go/src/bytes/bytes.go:1224
				// _ = "end of CoverTab[727]"
//line /usr/local/go/src/bytes/bytes.go:1224
			}() && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1224
				_go_fuzz_dep_.CoverTab[728]++
//line /usr/local/go/src/bytes/bytes.go:1224
				return tr == sr+'a'-'A'
//line /usr/local/go/src/bytes/bytes.go:1224
				// _ = "end of CoverTab[728]"
//line /usr/local/go/src/bytes/bytes.go:1224
			}() {
//line /usr/local/go/src/bytes/bytes.go:1224
				_go_fuzz_dep_.CoverTab[729]++
									continue
//line /usr/local/go/src/bytes/bytes.go:1225
				// _ = "end of CoverTab[729]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:1226
				_go_fuzz_dep_.CoverTab[730]++
//line /usr/local/go/src/bytes/bytes.go:1226
				// _ = "end of CoverTab[730]"
//line /usr/local/go/src/bytes/bytes.go:1226
			}
//line /usr/local/go/src/bytes/bytes.go:1226
			// _ = "end of CoverTab[725]"
//line /usr/local/go/src/bytes/bytes.go:1226
			_go_fuzz_dep_.CoverTab[726]++
								return false
//line /usr/local/go/src/bytes/bytes.go:1227
			// _ = "end of CoverTab[726]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1228
			_go_fuzz_dep_.CoverTab[731]++
//line /usr/local/go/src/bytes/bytes.go:1228
			// _ = "end of CoverTab[731]"
//line /usr/local/go/src/bytes/bytes.go:1228
		}
//line /usr/local/go/src/bytes/bytes.go:1228
		// _ = "end of CoverTab[713]"
//line /usr/local/go/src/bytes/bytes.go:1228
		_go_fuzz_dep_.CoverTab[714]++

//line /usr/local/go/src/bytes/bytes.go:1232
		r := unicode.SimpleFold(sr)
		for r != sr && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1233
			_go_fuzz_dep_.CoverTab[732]++
//line /usr/local/go/src/bytes/bytes.go:1233
			return r < tr
//line /usr/local/go/src/bytes/bytes.go:1233
			// _ = "end of CoverTab[732]"
//line /usr/local/go/src/bytes/bytes.go:1233
		}() {
//line /usr/local/go/src/bytes/bytes.go:1233
			_go_fuzz_dep_.CoverTab[733]++
								r = unicode.SimpleFold(r)
//line /usr/local/go/src/bytes/bytes.go:1234
			// _ = "end of CoverTab[733]"
		}
//line /usr/local/go/src/bytes/bytes.go:1235
		// _ = "end of CoverTab[714]"
//line /usr/local/go/src/bytes/bytes.go:1235
		_go_fuzz_dep_.CoverTab[715]++
							if r == tr {
//line /usr/local/go/src/bytes/bytes.go:1236
			_go_fuzz_dep_.CoverTab[734]++
								continue
//line /usr/local/go/src/bytes/bytes.go:1237
			// _ = "end of CoverTab[734]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1238
			_go_fuzz_dep_.CoverTab[735]++
//line /usr/local/go/src/bytes/bytes.go:1238
			// _ = "end of CoverTab[735]"
//line /usr/local/go/src/bytes/bytes.go:1238
		}
//line /usr/local/go/src/bytes/bytes.go:1238
		// _ = "end of CoverTab[715]"
//line /usr/local/go/src/bytes/bytes.go:1238
		_go_fuzz_dep_.CoverTab[716]++
							return false
//line /usr/local/go/src/bytes/bytes.go:1239
		// _ = "end of CoverTab[716]"
	}
//line /usr/local/go/src/bytes/bytes.go:1240
	// _ = "end of CoverTab[690]"
//line /usr/local/go/src/bytes/bytes.go:1240
	_go_fuzz_dep_.CoverTab[691]++

//line /usr/local/go/src/bytes/bytes.go:1243
	return len(s) == len(t)
//line /usr/local/go/src/bytes/bytes.go:1243
	// _ = "end of CoverTab[691]"
}

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func Index(s, sep []byte) int {
//line /usr/local/go/src/bytes/bytes.go:1247
	_go_fuzz_dep_.CoverTab[736]++
						n := len(sep)
						switch {
	case n == 0:
//line /usr/local/go/src/bytes/bytes.go:1250
		_go_fuzz_dep_.CoverTab[739]++
							return 0
//line /usr/local/go/src/bytes/bytes.go:1251
		// _ = "end of CoverTab[739]"
	case n == 1:
//line /usr/local/go/src/bytes/bytes.go:1252
		_go_fuzz_dep_.CoverTab[740]++
							return IndexByte(s, sep[0])
//line /usr/local/go/src/bytes/bytes.go:1253
		// _ = "end of CoverTab[740]"
	case n == len(s):
//line /usr/local/go/src/bytes/bytes.go:1254
		_go_fuzz_dep_.CoverTab[741]++
							if Equal(sep, s) {
//line /usr/local/go/src/bytes/bytes.go:1255
			_go_fuzz_dep_.CoverTab[748]++
								return 0
//line /usr/local/go/src/bytes/bytes.go:1256
			// _ = "end of CoverTab[748]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1257
			_go_fuzz_dep_.CoverTab[749]++
//line /usr/local/go/src/bytes/bytes.go:1257
			// _ = "end of CoverTab[749]"
//line /usr/local/go/src/bytes/bytes.go:1257
		}
//line /usr/local/go/src/bytes/bytes.go:1257
		// _ = "end of CoverTab[741]"
//line /usr/local/go/src/bytes/bytes.go:1257
		_go_fuzz_dep_.CoverTab[742]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:1258
		// _ = "end of CoverTab[742]"
	case n > len(s):
//line /usr/local/go/src/bytes/bytes.go:1259
		_go_fuzz_dep_.CoverTab[743]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:1260
		// _ = "end of CoverTab[743]"
	case n <= bytealg.MaxLen:
//line /usr/local/go/src/bytes/bytes.go:1261
		_go_fuzz_dep_.CoverTab[744]++

							if len(s) <= bytealg.MaxBruteForce {
//line /usr/local/go/src/bytes/bytes.go:1263
			_go_fuzz_dep_.CoverTab[750]++
								return bytealg.Index(s, sep)
//line /usr/local/go/src/bytes/bytes.go:1264
			// _ = "end of CoverTab[750]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1265
			_go_fuzz_dep_.CoverTab[751]++
//line /usr/local/go/src/bytes/bytes.go:1265
			// _ = "end of CoverTab[751]"
//line /usr/local/go/src/bytes/bytes.go:1265
		}
//line /usr/local/go/src/bytes/bytes.go:1265
		// _ = "end of CoverTab[744]"
//line /usr/local/go/src/bytes/bytes.go:1265
		_go_fuzz_dep_.CoverTab[745]++
							c0 := sep[0]
							c1 := sep[1]
							i := 0
							t := len(s) - n + 1
							fails := 0
							for i < t {
//line /usr/local/go/src/bytes/bytes.go:1271
			_go_fuzz_dep_.CoverTab[752]++
								if s[i] != c0 {
//line /usr/local/go/src/bytes/bytes.go:1272
				_go_fuzz_dep_.CoverTab[755]++

//line /usr/local/go/src/bytes/bytes.go:1275
				o := IndexByte(s[i+1:t], c0)
				if o < 0 {
//line /usr/local/go/src/bytes/bytes.go:1276
					_go_fuzz_dep_.CoverTab[757]++
										return -1
//line /usr/local/go/src/bytes/bytes.go:1277
					// _ = "end of CoverTab[757]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:1278
					_go_fuzz_dep_.CoverTab[758]++
//line /usr/local/go/src/bytes/bytes.go:1278
					// _ = "end of CoverTab[758]"
//line /usr/local/go/src/bytes/bytes.go:1278
				}
//line /usr/local/go/src/bytes/bytes.go:1278
				// _ = "end of CoverTab[755]"
//line /usr/local/go/src/bytes/bytes.go:1278
				_go_fuzz_dep_.CoverTab[756]++
									i += o + 1
//line /usr/local/go/src/bytes/bytes.go:1279
				// _ = "end of CoverTab[756]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:1280
				_go_fuzz_dep_.CoverTab[759]++
//line /usr/local/go/src/bytes/bytes.go:1280
				// _ = "end of CoverTab[759]"
//line /usr/local/go/src/bytes/bytes.go:1280
			}
//line /usr/local/go/src/bytes/bytes.go:1280
			// _ = "end of CoverTab[752]"
//line /usr/local/go/src/bytes/bytes.go:1280
			_go_fuzz_dep_.CoverTab[753]++
								if s[i+1] == c1 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1281
				_go_fuzz_dep_.CoverTab[760]++
//line /usr/local/go/src/bytes/bytes.go:1281
				return Equal(s[i:i+n], sep)
//line /usr/local/go/src/bytes/bytes.go:1281
				// _ = "end of CoverTab[760]"
//line /usr/local/go/src/bytes/bytes.go:1281
			}() {
//line /usr/local/go/src/bytes/bytes.go:1281
				_go_fuzz_dep_.CoverTab[761]++
									return i
//line /usr/local/go/src/bytes/bytes.go:1282
				// _ = "end of CoverTab[761]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:1283
				_go_fuzz_dep_.CoverTab[762]++
//line /usr/local/go/src/bytes/bytes.go:1283
				// _ = "end of CoverTab[762]"
//line /usr/local/go/src/bytes/bytes.go:1283
			}
//line /usr/local/go/src/bytes/bytes.go:1283
			// _ = "end of CoverTab[753]"
//line /usr/local/go/src/bytes/bytes.go:1283
			_go_fuzz_dep_.CoverTab[754]++
								fails++
								i++

								if fails > bytealg.Cutover(i) {
//line /usr/local/go/src/bytes/bytes.go:1287
				_go_fuzz_dep_.CoverTab[763]++
									r := bytealg.Index(s[i:], sep)
									if r >= 0 {
//line /usr/local/go/src/bytes/bytes.go:1289
					_go_fuzz_dep_.CoverTab[765]++
										return r + i
//line /usr/local/go/src/bytes/bytes.go:1290
					// _ = "end of CoverTab[765]"
				} else {
//line /usr/local/go/src/bytes/bytes.go:1291
					_go_fuzz_dep_.CoverTab[766]++
//line /usr/local/go/src/bytes/bytes.go:1291
					// _ = "end of CoverTab[766]"
//line /usr/local/go/src/bytes/bytes.go:1291
				}
//line /usr/local/go/src/bytes/bytes.go:1291
				// _ = "end of CoverTab[763]"
//line /usr/local/go/src/bytes/bytes.go:1291
				_go_fuzz_dep_.CoverTab[764]++
									return -1
//line /usr/local/go/src/bytes/bytes.go:1292
				// _ = "end of CoverTab[764]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:1293
				_go_fuzz_dep_.CoverTab[767]++
//line /usr/local/go/src/bytes/bytes.go:1293
				// _ = "end of CoverTab[767]"
//line /usr/local/go/src/bytes/bytes.go:1293
			}
//line /usr/local/go/src/bytes/bytes.go:1293
			// _ = "end of CoverTab[754]"
		}
//line /usr/local/go/src/bytes/bytes.go:1294
		// _ = "end of CoverTab[745]"
//line /usr/local/go/src/bytes/bytes.go:1294
		_go_fuzz_dep_.CoverTab[746]++
							return -1
//line /usr/local/go/src/bytes/bytes.go:1295
		// _ = "end of CoverTab[746]"
//line /usr/local/go/src/bytes/bytes.go:1295
	default:
//line /usr/local/go/src/bytes/bytes.go:1295
		_go_fuzz_dep_.CoverTab[747]++
//line /usr/local/go/src/bytes/bytes.go:1295
		// _ = "end of CoverTab[747]"
	}
//line /usr/local/go/src/bytes/bytes.go:1296
	// _ = "end of CoverTab[736]"
//line /usr/local/go/src/bytes/bytes.go:1296
	_go_fuzz_dep_.CoverTab[737]++
						c0 := sep[0]
						c1 := sep[1]
						i := 0
						fails := 0
						t := len(s) - n + 1
						for i < t {
//line /usr/local/go/src/bytes/bytes.go:1302
		_go_fuzz_dep_.CoverTab[768]++
							if s[i] != c0 {
//line /usr/local/go/src/bytes/bytes.go:1303
			_go_fuzz_dep_.CoverTab[771]++
								o := IndexByte(s[i+1:t], c0)
								if o < 0 {
//line /usr/local/go/src/bytes/bytes.go:1305
				_go_fuzz_dep_.CoverTab[773]++
									break
//line /usr/local/go/src/bytes/bytes.go:1306
				// _ = "end of CoverTab[773]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:1307
				_go_fuzz_dep_.CoverTab[774]++
//line /usr/local/go/src/bytes/bytes.go:1307
				// _ = "end of CoverTab[774]"
//line /usr/local/go/src/bytes/bytes.go:1307
			}
//line /usr/local/go/src/bytes/bytes.go:1307
			// _ = "end of CoverTab[771]"
//line /usr/local/go/src/bytes/bytes.go:1307
			_go_fuzz_dep_.CoverTab[772]++
								i += o + 1
//line /usr/local/go/src/bytes/bytes.go:1308
			// _ = "end of CoverTab[772]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1309
			_go_fuzz_dep_.CoverTab[775]++
//line /usr/local/go/src/bytes/bytes.go:1309
			// _ = "end of CoverTab[775]"
//line /usr/local/go/src/bytes/bytes.go:1309
		}
//line /usr/local/go/src/bytes/bytes.go:1309
		// _ = "end of CoverTab[768]"
//line /usr/local/go/src/bytes/bytes.go:1309
		_go_fuzz_dep_.CoverTab[769]++
							if s[i+1] == c1 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1310
			_go_fuzz_dep_.CoverTab[776]++
//line /usr/local/go/src/bytes/bytes.go:1310
			return Equal(s[i:i+n], sep)
//line /usr/local/go/src/bytes/bytes.go:1310
			// _ = "end of CoverTab[776]"
//line /usr/local/go/src/bytes/bytes.go:1310
		}() {
//line /usr/local/go/src/bytes/bytes.go:1310
			_go_fuzz_dep_.CoverTab[777]++
								return i
//line /usr/local/go/src/bytes/bytes.go:1311
			// _ = "end of CoverTab[777]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1312
			_go_fuzz_dep_.CoverTab[778]++
//line /usr/local/go/src/bytes/bytes.go:1312
			// _ = "end of CoverTab[778]"
//line /usr/local/go/src/bytes/bytes.go:1312
		}
//line /usr/local/go/src/bytes/bytes.go:1312
		// _ = "end of CoverTab[769]"
//line /usr/local/go/src/bytes/bytes.go:1312
		_go_fuzz_dep_.CoverTab[770]++
							i++
							fails++
							if fails >= 4+i>>4 && func() bool {
//line /usr/local/go/src/bytes/bytes.go:1315
			_go_fuzz_dep_.CoverTab[779]++
//line /usr/local/go/src/bytes/bytes.go:1315
			return i < t
//line /usr/local/go/src/bytes/bytes.go:1315
			// _ = "end of CoverTab[779]"
//line /usr/local/go/src/bytes/bytes.go:1315
		}() {
//line /usr/local/go/src/bytes/bytes.go:1315
			_go_fuzz_dep_.CoverTab[780]++

//line /usr/local/go/src/bytes/bytes.go:1324
			j := bytealg.IndexRabinKarpBytes(s[i:], sep)
			if j < 0 {
//line /usr/local/go/src/bytes/bytes.go:1325
				_go_fuzz_dep_.CoverTab[782]++
									return -1
//line /usr/local/go/src/bytes/bytes.go:1326
				// _ = "end of CoverTab[782]"
			} else {
//line /usr/local/go/src/bytes/bytes.go:1327
				_go_fuzz_dep_.CoverTab[783]++
//line /usr/local/go/src/bytes/bytes.go:1327
				// _ = "end of CoverTab[783]"
//line /usr/local/go/src/bytes/bytes.go:1327
			}
//line /usr/local/go/src/bytes/bytes.go:1327
			// _ = "end of CoverTab[780]"
//line /usr/local/go/src/bytes/bytes.go:1327
			_go_fuzz_dep_.CoverTab[781]++
								return i + j
//line /usr/local/go/src/bytes/bytes.go:1328
			// _ = "end of CoverTab[781]"
		} else {
//line /usr/local/go/src/bytes/bytes.go:1329
			_go_fuzz_dep_.CoverTab[784]++
//line /usr/local/go/src/bytes/bytes.go:1329
			// _ = "end of CoverTab[784]"
//line /usr/local/go/src/bytes/bytes.go:1329
		}
//line /usr/local/go/src/bytes/bytes.go:1329
		// _ = "end of CoverTab[770]"
	}
//line /usr/local/go/src/bytes/bytes.go:1330
	// _ = "end of CoverTab[737]"
//line /usr/local/go/src/bytes/bytes.go:1330
	_go_fuzz_dep_.CoverTab[738]++
						return -1
//line /usr/local/go/src/bytes/bytes.go:1331
	// _ = "end of CoverTab[738]"
}

// Cut slices s around the first instance of sep,
//line /usr/local/go/src/bytes/bytes.go:1334
// returning the text before and after sep.
//line /usr/local/go/src/bytes/bytes.go:1334
// The found result reports whether sep appears in s.
//line /usr/local/go/src/bytes/bytes.go:1334
// If sep does not appear in s, cut returns s, nil, false.
//line /usr/local/go/src/bytes/bytes.go:1334
//
//line /usr/local/go/src/bytes/bytes.go:1334
// Cut returns slices of the original slice s, not copies.
//line /usr/local/go/src/bytes/bytes.go:1340
func Cut(s, sep []byte) (before, after []byte, found bool) {
//line /usr/local/go/src/bytes/bytes.go:1340
	_go_fuzz_dep_.CoverTab[785]++
						if i := Index(s, sep); i >= 0 {
//line /usr/local/go/src/bytes/bytes.go:1341
		_go_fuzz_dep_.CoverTab[787]++
							return s[:i], s[i+len(sep):], true
//line /usr/local/go/src/bytes/bytes.go:1342
		// _ = "end of CoverTab[787]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1343
		_go_fuzz_dep_.CoverTab[788]++
//line /usr/local/go/src/bytes/bytes.go:1343
		// _ = "end of CoverTab[788]"
//line /usr/local/go/src/bytes/bytes.go:1343
	}
//line /usr/local/go/src/bytes/bytes.go:1343
	// _ = "end of CoverTab[785]"
//line /usr/local/go/src/bytes/bytes.go:1343
	_go_fuzz_dep_.CoverTab[786]++
						return s, nil, false
//line /usr/local/go/src/bytes/bytes.go:1344
	// _ = "end of CoverTab[786]"
}

// Clone returns a copy of b[:len(b)].
//line /usr/local/go/src/bytes/bytes.go:1347
// The result may have additional unused capacity.
//line /usr/local/go/src/bytes/bytes.go:1347
// Clone(nil) returns nil.
//line /usr/local/go/src/bytes/bytes.go:1350
func Clone(b []byte) []byte {
//line /usr/local/go/src/bytes/bytes.go:1350
	_go_fuzz_dep_.CoverTab[789]++
						if b == nil {
//line /usr/local/go/src/bytes/bytes.go:1351
		_go_fuzz_dep_.CoverTab[791]++
							return nil
//line /usr/local/go/src/bytes/bytes.go:1352
		// _ = "end of CoverTab[791]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1353
		_go_fuzz_dep_.CoverTab[792]++
//line /usr/local/go/src/bytes/bytes.go:1353
		// _ = "end of CoverTab[792]"
//line /usr/local/go/src/bytes/bytes.go:1353
	}
//line /usr/local/go/src/bytes/bytes.go:1353
	// _ = "end of CoverTab[789]"
//line /usr/local/go/src/bytes/bytes.go:1353
	_go_fuzz_dep_.CoverTab[790]++
						return append([]byte{}, b...)
//line /usr/local/go/src/bytes/bytes.go:1354
	// _ = "end of CoverTab[790]"
}

// CutPrefix returns s without the provided leading prefix byte slice
//line /usr/local/go/src/bytes/bytes.go:1357
// and reports whether it found the prefix.
//line /usr/local/go/src/bytes/bytes.go:1357
// If s doesn't start with prefix, CutPrefix returns s, false.
//line /usr/local/go/src/bytes/bytes.go:1357
// If prefix is the empty byte slice, CutPrefix returns s, true.
//line /usr/local/go/src/bytes/bytes.go:1357
//
//line /usr/local/go/src/bytes/bytes.go:1357
// CutPrefix returns slices of the original slice s, not copies.
//line /usr/local/go/src/bytes/bytes.go:1363
func CutPrefix(s, prefix []byte) (after []byte, found bool) {
//line /usr/local/go/src/bytes/bytes.go:1363
	_go_fuzz_dep_.CoverTab[793]++
						if !HasPrefix(s, prefix) {
//line /usr/local/go/src/bytes/bytes.go:1364
		_go_fuzz_dep_.CoverTab[795]++
							return s, false
//line /usr/local/go/src/bytes/bytes.go:1365
		// _ = "end of CoverTab[795]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1366
		_go_fuzz_dep_.CoverTab[796]++
//line /usr/local/go/src/bytes/bytes.go:1366
		// _ = "end of CoverTab[796]"
//line /usr/local/go/src/bytes/bytes.go:1366
	}
//line /usr/local/go/src/bytes/bytes.go:1366
	// _ = "end of CoverTab[793]"
//line /usr/local/go/src/bytes/bytes.go:1366
	_go_fuzz_dep_.CoverTab[794]++
						return s[len(prefix):], true
//line /usr/local/go/src/bytes/bytes.go:1367
	// _ = "end of CoverTab[794]"
}

// CutSuffix returns s without the provided ending suffix byte slice
//line /usr/local/go/src/bytes/bytes.go:1370
// and reports whether it found the suffix.
//line /usr/local/go/src/bytes/bytes.go:1370
// If s doesn't end with suffix, CutSuffix returns s, false.
//line /usr/local/go/src/bytes/bytes.go:1370
// If suffix is the empty byte slice, CutSuffix returns s, true.
//line /usr/local/go/src/bytes/bytes.go:1370
//
//line /usr/local/go/src/bytes/bytes.go:1370
// CutSuffix returns slices of the original slice s, not copies.
//line /usr/local/go/src/bytes/bytes.go:1376
func CutSuffix(s, suffix []byte) (before []byte, found bool) {
//line /usr/local/go/src/bytes/bytes.go:1376
	_go_fuzz_dep_.CoverTab[797]++
						if !HasSuffix(s, suffix) {
//line /usr/local/go/src/bytes/bytes.go:1377
		_go_fuzz_dep_.CoverTab[799]++
							return s, false
//line /usr/local/go/src/bytes/bytes.go:1378
		// _ = "end of CoverTab[799]"
	} else {
//line /usr/local/go/src/bytes/bytes.go:1379
		_go_fuzz_dep_.CoverTab[800]++
//line /usr/local/go/src/bytes/bytes.go:1379
		// _ = "end of CoverTab[800]"
//line /usr/local/go/src/bytes/bytes.go:1379
	}
//line /usr/local/go/src/bytes/bytes.go:1379
	// _ = "end of CoverTab[797]"
//line /usr/local/go/src/bytes/bytes.go:1379
	_go_fuzz_dep_.CoverTab[798]++
						return s[:len(s)-len(suffix)], true
//line /usr/local/go/src/bytes/bytes.go:1380
	// _ = "end of CoverTab[798]"
}

//line /usr/local/go/src/bytes/bytes.go:1381
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/bytes/bytes.go:1381
var _ = _go_fuzz_dep_.CoverTab
