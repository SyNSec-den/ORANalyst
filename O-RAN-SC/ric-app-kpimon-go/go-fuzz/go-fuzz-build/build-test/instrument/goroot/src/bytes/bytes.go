// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/bytes/bytes.go:5
// Package bytes implements functions for the manipulation of byte slices.
//line /snap/go/10455/src/bytes/bytes.go:5
// It is analogous to the facilities of the [strings] package.
//line /snap/go/10455/src/bytes/bytes.go:7
package bytes

//line /snap/go/10455/src/bytes/bytes.go:7
import (
//line /snap/go/10455/src/bytes/bytes.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/bytes/bytes.go:7
)
//line /snap/go/10455/src/bytes/bytes.go:7
import (
//line /snap/go/10455/src/bytes/bytes.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/bytes/bytes.go:7
)

import (
	"internal/bytealg"
	"unicode"
	"unicode/utf8"
)

// Equal reports whether a and b
//line /snap/go/10455/src/bytes/bytes.go:15
// are the same length and contain the same bytes.
//line /snap/go/10455/src/bytes/bytes.go:15
// A nil argument is equivalent to an empty slice.
//line /snap/go/10455/src/bytes/bytes.go:18
func Equal(a, b []byte) bool {
//line /snap/go/10455/src/bytes/bytes.go:18
	_go_fuzz_dep_.CoverTab[144]++

						return string(a) == string(b)
//line /snap/go/10455/src/bytes/bytes.go:20
	// _ = "end of CoverTab[144]"
}

// Compare returns an integer comparing two byte slices lexicographically.
//line /snap/go/10455/src/bytes/bytes.go:23
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//line /snap/go/10455/src/bytes/bytes.go:23
// A nil argument is equivalent to an empty slice.
//line /snap/go/10455/src/bytes/bytes.go:26
func Compare(a, b []byte) int {
//line /snap/go/10455/src/bytes/bytes.go:26
	_go_fuzz_dep_.CoverTab[145]++
						return bytealg.Compare(a, b)
//line /snap/go/10455/src/bytes/bytes.go:27
	// _ = "end of CoverTab[145]"
}

// explode splits s into a slice of UTF-8 sequences, one per Unicode code point (still slices of bytes),
//line /snap/go/10455/src/bytes/bytes.go:30
// up to a maximum of n byte slices. Invalid UTF-8 sequences are chopped into individual bytes.
//line /snap/go/10455/src/bytes/bytes.go:32
func explode(s []byte, n int) [][]byte {
//line /snap/go/10455/src/bytes/bytes.go:32
	_go_fuzz_dep_.CoverTab[146]++
						if n <= 0 || func() bool {
//line /snap/go/10455/src/bytes/bytes.go:33
		_go_fuzz_dep_.CoverTab[149]++
//line /snap/go/10455/src/bytes/bytes.go:33
		return n > len(s)
//line /snap/go/10455/src/bytes/bytes.go:33
		// _ = "end of CoverTab[149]"
//line /snap/go/10455/src/bytes/bytes.go:33
	}() {
//line /snap/go/10455/src/bytes/bytes.go:33
		_go_fuzz_dep_.CoverTab[524367]++
//line /snap/go/10455/src/bytes/bytes.go:33
		_go_fuzz_dep_.CoverTab[150]++
							n = len(s)
//line /snap/go/10455/src/bytes/bytes.go:34
		// _ = "end of CoverTab[150]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:35
		_go_fuzz_dep_.CoverTab[524368]++
//line /snap/go/10455/src/bytes/bytes.go:35
		_go_fuzz_dep_.CoverTab[151]++
//line /snap/go/10455/src/bytes/bytes.go:35
		// _ = "end of CoverTab[151]"
//line /snap/go/10455/src/bytes/bytes.go:35
	}
//line /snap/go/10455/src/bytes/bytes.go:35
	// _ = "end of CoverTab[146]"
//line /snap/go/10455/src/bytes/bytes.go:35
	_go_fuzz_dep_.CoverTab[147]++
						a := make([][]byte, n)
						var size int
						na := 0
//line /snap/go/10455/src/bytes/bytes.go:38
	_go_fuzz_dep_.CoverTab[786434] = 0
						for len(s) > 0 {
//line /snap/go/10455/src/bytes/bytes.go:39
		if _go_fuzz_dep_.CoverTab[786434] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:39
			_go_fuzz_dep_.CoverTab[524699]++
//line /snap/go/10455/src/bytes/bytes.go:39
		} else {
//line /snap/go/10455/src/bytes/bytes.go:39
			_go_fuzz_dep_.CoverTab[524700]++
//line /snap/go/10455/src/bytes/bytes.go:39
		}
//line /snap/go/10455/src/bytes/bytes.go:39
		_go_fuzz_dep_.CoverTab[786434] = 1
//line /snap/go/10455/src/bytes/bytes.go:39
		_go_fuzz_dep_.CoverTab[152]++
							if na+1 >= n {
//line /snap/go/10455/src/bytes/bytes.go:40
			_go_fuzz_dep_.CoverTab[524369]++
//line /snap/go/10455/src/bytes/bytes.go:40
			_go_fuzz_dep_.CoverTab[154]++
								a[na] = s
								na++
								break
//line /snap/go/10455/src/bytes/bytes.go:43
			// _ = "end of CoverTab[154]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:44
			_go_fuzz_dep_.CoverTab[524370]++
//line /snap/go/10455/src/bytes/bytes.go:44
			_go_fuzz_dep_.CoverTab[155]++
//line /snap/go/10455/src/bytes/bytes.go:44
			// _ = "end of CoverTab[155]"
//line /snap/go/10455/src/bytes/bytes.go:44
		}
//line /snap/go/10455/src/bytes/bytes.go:44
		// _ = "end of CoverTab[152]"
//line /snap/go/10455/src/bytes/bytes.go:44
		_go_fuzz_dep_.CoverTab[153]++
							_, size = utf8.DecodeRune(s)
							a[na] = s[0:size:size]
							s = s[size:]
							na++
//line /snap/go/10455/src/bytes/bytes.go:48
		// _ = "end of CoverTab[153]"
	}
//line /snap/go/10455/src/bytes/bytes.go:49
	if _go_fuzz_dep_.CoverTab[786434] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:49
		_go_fuzz_dep_.CoverTab[524701]++
//line /snap/go/10455/src/bytes/bytes.go:49
	} else {
//line /snap/go/10455/src/bytes/bytes.go:49
		_go_fuzz_dep_.CoverTab[524702]++
//line /snap/go/10455/src/bytes/bytes.go:49
	}
//line /snap/go/10455/src/bytes/bytes.go:49
	// _ = "end of CoverTab[147]"
//line /snap/go/10455/src/bytes/bytes.go:49
	_go_fuzz_dep_.CoverTab[148]++
						return a[0:na]
//line /snap/go/10455/src/bytes/bytes.go:50
	// _ = "end of CoverTab[148]"
}

// Count counts the number of non-overlapping instances of sep in s.
//line /snap/go/10455/src/bytes/bytes.go:53
// If sep is an empty slice, Count returns 1 + the number of UTF-8-encoded code points in s.
//line /snap/go/10455/src/bytes/bytes.go:55
func Count(s, sep []byte) int {
//line /snap/go/10455/src/bytes/bytes.go:55
	_go_fuzz_dep_.CoverTab[156]++

						if len(sep) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:57
		_go_fuzz_dep_.CoverTab[524371]++
//line /snap/go/10455/src/bytes/bytes.go:57
		_go_fuzz_dep_.CoverTab[159]++
							return utf8.RuneCount(s) + 1
//line /snap/go/10455/src/bytes/bytes.go:58
		// _ = "end of CoverTab[159]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:59
		_go_fuzz_dep_.CoverTab[524372]++
//line /snap/go/10455/src/bytes/bytes.go:59
		_go_fuzz_dep_.CoverTab[160]++
//line /snap/go/10455/src/bytes/bytes.go:59
		// _ = "end of CoverTab[160]"
//line /snap/go/10455/src/bytes/bytes.go:59
	}
//line /snap/go/10455/src/bytes/bytes.go:59
	// _ = "end of CoverTab[156]"
//line /snap/go/10455/src/bytes/bytes.go:59
	_go_fuzz_dep_.CoverTab[157]++
						if len(sep) == 1 {
//line /snap/go/10455/src/bytes/bytes.go:60
		_go_fuzz_dep_.CoverTab[524373]++
//line /snap/go/10455/src/bytes/bytes.go:60
		_go_fuzz_dep_.CoverTab[161]++
							return bytealg.Count(s, sep[0])
//line /snap/go/10455/src/bytes/bytes.go:61
		// _ = "end of CoverTab[161]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:62
		_go_fuzz_dep_.CoverTab[524374]++
//line /snap/go/10455/src/bytes/bytes.go:62
		_go_fuzz_dep_.CoverTab[162]++
//line /snap/go/10455/src/bytes/bytes.go:62
		// _ = "end of CoverTab[162]"
//line /snap/go/10455/src/bytes/bytes.go:62
	}
//line /snap/go/10455/src/bytes/bytes.go:62
	// _ = "end of CoverTab[157]"
//line /snap/go/10455/src/bytes/bytes.go:62
	_go_fuzz_dep_.CoverTab[158]++
						n := 0
//line /snap/go/10455/src/bytes/bytes.go:63
	_go_fuzz_dep_.CoverTab[786435] = 0
						for {
//line /snap/go/10455/src/bytes/bytes.go:64
		if _go_fuzz_dep_.CoverTab[786435] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:64
			_go_fuzz_dep_.CoverTab[524703]++
//line /snap/go/10455/src/bytes/bytes.go:64
		} else {
//line /snap/go/10455/src/bytes/bytes.go:64
			_go_fuzz_dep_.CoverTab[524704]++
//line /snap/go/10455/src/bytes/bytes.go:64
		}
//line /snap/go/10455/src/bytes/bytes.go:64
		_go_fuzz_dep_.CoverTab[786435] = 1
//line /snap/go/10455/src/bytes/bytes.go:64
		_go_fuzz_dep_.CoverTab[163]++
							i := Index(s, sep)
							if i == -1 {
//line /snap/go/10455/src/bytes/bytes.go:66
			_go_fuzz_dep_.CoverTab[524375]++
//line /snap/go/10455/src/bytes/bytes.go:66
			_go_fuzz_dep_.CoverTab[165]++
								return n
//line /snap/go/10455/src/bytes/bytes.go:67
			// _ = "end of CoverTab[165]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:68
			_go_fuzz_dep_.CoverTab[524376]++
//line /snap/go/10455/src/bytes/bytes.go:68
			_go_fuzz_dep_.CoverTab[166]++
//line /snap/go/10455/src/bytes/bytes.go:68
			// _ = "end of CoverTab[166]"
//line /snap/go/10455/src/bytes/bytes.go:68
		}
//line /snap/go/10455/src/bytes/bytes.go:68
		// _ = "end of CoverTab[163]"
//line /snap/go/10455/src/bytes/bytes.go:68
		_go_fuzz_dep_.CoverTab[164]++
							n++
							s = s[i+len(sep):]
//line /snap/go/10455/src/bytes/bytes.go:70
		// _ = "end of CoverTab[164]"
	}
//line /snap/go/10455/src/bytes/bytes.go:71
	// _ = "end of CoverTab[158]"
}

// Contains reports whether subslice is within b.
func Contains(b, subslice []byte) bool {
//line /snap/go/10455/src/bytes/bytes.go:75
	_go_fuzz_dep_.CoverTab[167]++
						return Index(b, subslice) != -1
//line /snap/go/10455/src/bytes/bytes.go:76
	// _ = "end of CoverTab[167]"
}

// ContainsAny reports whether any of the UTF-8-encoded code points in chars are within b.
func ContainsAny(b []byte, chars string) bool {
//line /snap/go/10455/src/bytes/bytes.go:80
	_go_fuzz_dep_.CoverTab[168]++
						return IndexAny(b, chars) >= 0
//line /snap/go/10455/src/bytes/bytes.go:81
	// _ = "end of CoverTab[168]"
}

// ContainsRune reports whether the rune is contained in the UTF-8-encoded byte slice b.
func ContainsRune(b []byte, r rune) bool {
//line /snap/go/10455/src/bytes/bytes.go:85
	_go_fuzz_dep_.CoverTab[169]++
						return IndexRune(b, r) >= 0
//line /snap/go/10455/src/bytes/bytes.go:86
	// _ = "end of CoverTab[169]"
}

// ContainsFunc reports whether any of the UTF-8-encoded code points r within b satisfy f(r).
func ContainsFunc(b []byte, f func(rune) bool) bool {
//line /snap/go/10455/src/bytes/bytes.go:90
	_go_fuzz_dep_.CoverTab[170]++
						return IndexFunc(b, f) >= 0
//line /snap/go/10455/src/bytes/bytes.go:91
	// _ = "end of CoverTab[170]"
}

// IndexByte returns the index of the first instance of c in b, or -1 if c is not present in b.
func IndexByte(b []byte, c byte) int {
//line /snap/go/10455/src/bytes/bytes.go:95
	_go_fuzz_dep_.CoverTab[171]++
						return bytealg.IndexByte(b, c)
//line /snap/go/10455/src/bytes/bytes.go:96
	// _ = "end of CoverTab[171]"
}

func indexBytePortable(s []byte, c byte) int {
//line /snap/go/10455/src/bytes/bytes.go:99
	_go_fuzz_dep_.CoverTab[172]++
//line /snap/go/10455/src/bytes/bytes.go:99
	_go_fuzz_dep_.CoverTab[786436] = 0
						for i, b := range s {
//line /snap/go/10455/src/bytes/bytes.go:100
		if _go_fuzz_dep_.CoverTab[786436] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:100
			_go_fuzz_dep_.CoverTab[524707]++
//line /snap/go/10455/src/bytes/bytes.go:100
		} else {
//line /snap/go/10455/src/bytes/bytes.go:100
			_go_fuzz_dep_.CoverTab[524708]++
//line /snap/go/10455/src/bytes/bytes.go:100
		}
//line /snap/go/10455/src/bytes/bytes.go:100
		_go_fuzz_dep_.CoverTab[786436] = 1
//line /snap/go/10455/src/bytes/bytes.go:100
		_go_fuzz_dep_.CoverTab[174]++
							if b == c {
//line /snap/go/10455/src/bytes/bytes.go:101
			_go_fuzz_dep_.CoverTab[524377]++
//line /snap/go/10455/src/bytes/bytes.go:101
			_go_fuzz_dep_.CoverTab[175]++
								return i
//line /snap/go/10455/src/bytes/bytes.go:102
			// _ = "end of CoverTab[175]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:103
			_go_fuzz_dep_.CoverTab[524378]++
//line /snap/go/10455/src/bytes/bytes.go:103
			_go_fuzz_dep_.CoverTab[176]++
//line /snap/go/10455/src/bytes/bytes.go:103
			// _ = "end of CoverTab[176]"
//line /snap/go/10455/src/bytes/bytes.go:103
		}
//line /snap/go/10455/src/bytes/bytes.go:103
		// _ = "end of CoverTab[174]"
	}
//line /snap/go/10455/src/bytes/bytes.go:104
	if _go_fuzz_dep_.CoverTab[786436] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:104
		_go_fuzz_dep_.CoverTab[524709]++
//line /snap/go/10455/src/bytes/bytes.go:104
	} else {
//line /snap/go/10455/src/bytes/bytes.go:104
		_go_fuzz_dep_.CoverTab[524710]++
//line /snap/go/10455/src/bytes/bytes.go:104
	}
//line /snap/go/10455/src/bytes/bytes.go:104
	// _ = "end of CoverTab[172]"
//line /snap/go/10455/src/bytes/bytes.go:104
	_go_fuzz_dep_.CoverTab[173]++
						return -1
//line /snap/go/10455/src/bytes/bytes.go:105
	// _ = "end of CoverTab[173]"
}

// LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.
func LastIndex(s, sep []byte) int {
//line /snap/go/10455/src/bytes/bytes.go:109
	_go_fuzz_dep_.CoverTab[177]++
						n := len(sep)
						switch {
	case n == 0:
//line /snap/go/10455/src/bytes/bytes.go:112
		_go_fuzz_dep_.CoverTab[524379]++
//line /snap/go/10455/src/bytes/bytes.go:112
		_go_fuzz_dep_.CoverTab[182]++
							return len(s)
//line /snap/go/10455/src/bytes/bytes.go:113
		// _ = "end of CoverTab[182]"
	case n == 1:
//line /snap/go/10455/src/bytes/bytes.go:114
		_go_fuzz_dep_.CoverTab[524380]++
//line /snap/go/10455/src/bytes/bytes.go:114
		_go_fuzz_dep_.CoverTab[183]++
							return LastIndexByte(s, sep[0])
//line /snap/go/10455/src/bytes/bytes.go:115
		// _ = "end of CoverTab[183]"
	case n == len(s):
//line /snap/go/10455/src/bytes/bytes.go:116
		_go_fuzz_dep_.CoverTab[524381]++
//line /snap/go/10455/src/bytes/bytes.go:116
		_go_fuzz_dep_.CoverTab[184]++
							if Equal(s, sep) {
//line /snap/go/10455/src/bytes/bytes.go:117
			_go_fuzz_dep_.CoverTab[524384]++
//line /snap/go/10455/src/bytes/bytes.go:117
			_go_fuzz_dep_.CoverTab[188]++
								return 0
//line /snap/go/10455/src/bytes/bytes.go:118
			// _ = "end of CoverTab[188]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:119
			_go_fuzz_dep_.CoverTab[524385]++
//line /snap/go/10455/src/bytes/bytes.go:119
			_go_fuzz_dep_.CoverTab[189]++
//line /snap/go/10455/src/bytes/bytes.go:119
			// _ = "end of CoverTab[189]"
//line /snap/go/10455/src/bytes/bytes.go:119
		}
//line /snap/go/10455/src/bytes/bytes.go:119
		// _ = "end of CoverTab[184]"
//line /snap/go/10455/src/bytes/bytes.go:119
		_go_fuzz_dep_.CoverTab[185]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:120
		// _ = "end of CoverTab[185]"
	case n > len(s):
//line /snap/go/10455/src/bytes/bytes.go:121
		_go_fuzz_dep_.CoverTab[524382]++
//line /snap/go/10455/src/bytes/bytes.go:121
		_go_fuzz_dep_.CoverTab[186]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:122
		// _ = "end of CoverTab[186]"
//line /snap/go/10455/src/bytes/bytes.go:122
	default:
//line /snap/go/10455/src/bytes/bytes.go:122
		_go_fuzz_dep_.CoverTab[524383]++
//line /snap/go/10455/src/bytes/bytes.go:122
		_go_fuzz_dep_.CoverTab[187]++
//line /snap/go/10455/src/bytes/bytes.go:122
		// _ = "end of CoverTab[187]"
	}
//line /snap/go/10455/src/bytes/bytes.go:123
	// _ = "end of CoverTab[177]"
//line /snap/go/10455/src/bytes/bytes.go:123
	_go_fuzz_dep_.CoverTab[178]++

						hashss, pow := bytealg.HashStrRevBytes(sep)
						last := len(s) - n
						var h uint32
//line /snap/go/10455/src/bytes/bytes.go:127
	_go_fuzz_dep_.CoverTab[786437] = 0
						for i := len(s) - 1; i >= last; i-- {
//line /snap/go/10455/src/bytes/bytes.go:128
		if _go_fuzz_dep_.CoverTab[786437] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:128
			_go_fuzz_dep_.CoverTab[524711]++
//line /snap/go/10455/src/bytes/bytes.go:128
		} else {
//line /snap/go/10455/src/bytes/bytes.go:128
			_go_fuzz_dep_.CoverTab[524712]++
//line /snap/go/10455/src/bytes/bytes.go:128
		}
//line /snap/go/10455/src/bytes/bytes.go:128
		_go_fuzz_dep_.CoverTab[786437] = 1
//line /snap/go/10455/src/bytes/bytes.go:128
		_go_fuzz_dep_.CoverTab[190]++
							h = h*bytealg.PrimeRK + uint32(s[i])
//line /snap/go/10455/src/bytes/bytes.go:129
		// _ = "end of CoverTab[190]"
	}
//line /snap/go/10455/src/bytes/bytes.go:130
	if _go_fuzz_dep_.CoverTab[786437] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:130
		_go_fuzz_dep_.CoverTab[524713]++
//line /snap/go/10455/src/bytes/bytes.go:130
	} else {
//line /snap/go/10455/src/bytes/bytes.go:130
		_go_fuzz_dep_.CoverTab[524714]++
//line /snap/go/10455/src/bytes/bytes.go:130
	}
//line /snap/go/10455/src/bytes/bytes.go:130
	// _ = "end of CoverTab[178]"
//line /snap/go/10455/src/bytes/bytes.go:130
	_go_fuzz_dep_.CoverTab[179]++
						if h == hashss && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:131
		_go_fuzz_dep_.CoverTab[191]++
//line /snap/go/10455/src/bytes/bytes.go:131
		return Equal(s[last:], sep)
//line /snap/go/10455/src/bytes/bytes.go:131
		// _ = "end of CoverTab[191]"
//line /snap/go/10455/src/bytes/bytes.go:131
	}() {
//line /snap/go/10455/src/bytes/bytes.go:131
		_go_fuzz_dep_.CoverTab[524386]++
//line /snap/go/10455/src/bytes/bytes.go:131
		_go_fuzz_dep_.CoverTab[192]++
							return last
//line /snap/go/10455/src/bytes/bytes.go:132
		// _ = "end of CoverTab[192]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:133
		_go_fuzz_dep_.CoverTab[524387]++
//line /snap/go/10455/src/bytes/bytes.go:133
		_go_fuzz_dep_.CoverTab[193]++
//line /snap/go/10455/src/bytes/bytes.go:133
		// _ = "end of CoverTab[193]"
//line /snap/go/10455/src/bytes/bytes.go:133
	}
//line /snap/go/10455/src/bytes/bytes.go:133
	// _ = "end of CoverTab[179]"
//line /snap/go/10455/src/bytes/bytes.go:133
	_go_fuzz_dep_.CoverTab[180]++
//line /snap/go/10455/src/bytes/bytes.go:133
	_go_fuzz_dep_.CoverTab[786438] = 0
						for i := last - 1; i >= 0; i-- {
//line /snap/go/10455/src/bytes/bytes.go:134
		if _go_fuzz_dep_.CoverTab[786438] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:134
			_go_fuzz_dep_.CoverTab[524715]++
//line /snap/go/10455/src/bytes/bytes.go:134
		} else {
//line /snap/go/10455/src/bytes/bytes.go:134
			_go_fuzz_dep_.CoverTab[524716]++
//line /snap/go/10455/src/bytes/bytes.go:134
		}
//line /snap/go/10455/src/bytes/bytes.go:134
		_go_fuzz_dep_.CoverTab[786438] = 1
//line /snap/go/10455/src/bytes/bytes.go:134
		_go_fuzz_dep_.CoverTab[194]++
							h *= bytealg.PrimeRK
							h += uint32(s[i])
							h -= pow * uint32(s[i+n])
							if h == hashss && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:138
			_go_fuzz_dep_.CoverTab[195]++
//line /snap/go/10455/src/bytes/bytes.go:138
			return Equal(s[i:i+n], sep)
//line /snap/go/10455/src/bytes/bytes.go:138
			// _ = "end of CoverTab[195]"
//line /snap/go/10455/src/bytes/bytes.go:138
		}() {
//line /snap/go/10455/src/bytes/bytes.go:138
			_go_fuzz_dep_.CoverTab[524388]++
//line /snap/go/10455/src/bytes/bytes.go:138
			_go_fuzz_dep_.CoverTab[196]++
								return i
//line /snap/go/10455/src/bytes/bytes.go:139
			// _ = "end of CoverTab[196]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:140
			_go_fuzz_dep_.CoverTab[524389]++
//line /snap/go/10455/src/bytes/bytes.go:140
			_go_fuzz_dep_.CoverTab[197]++
//line /snap/go/10455/src/bytes/bytes.go:140
			// _ = "end of CoverTab[197]"
//line /snap/go/10455/src/bytes/bytes.go:140
		}
//line /snap/go/10455/src/bytes/bytes.go:140
		// _ = "end of CoverTab[194]"
	}
//line /snap/go/10455/src/bytes/bytes.go:141
	if _go_fuzz_dep_.CoverTab[786438] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:141
		_go_fuzz_dep_.CoverTab[524717]++
//line /snap/go/10455/src/bytes/bytes.go:141
	} else {
//line /snap/go/10455/src/bytes/bytes.go:141
		_go_fuzz_dep_.CoverTab[524718]++
//line /snap/go/10455/src/bytes/bytes.go:141
	}
//line /snap/go/10455/src/bytes/bytes.go:141
	// _ = "end of CoverTab[180]"
//line /snap/go/10455/src/bytes/bytes.go:141
	_go_fuzz_dep_.CoverTab[181]++
						return -1
//line /snap/go/10455/src/bytes/bytes.go:142
	// _ = "end of CoverTab[181]"
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte(s []byte, c byte) int {
//line /snap/go/10455/src/bytes/bytes.go:146
	_go_fuzz_dep_.CoverTab[198]++
//line /snap/go/10455/src/bytes/bytes.go:146
	_go_fuzz_dep_.CoverTab[786439] = 0
						for i := len(s) - 1; i >= 0; i-- {
//line /snap/go/10455/src/bytes/bytes.go:147
		if _go_fuzz_dep_.CoverTab[786439] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:147
			_go_fuzz_dep_.CoverTab[524719]++
//line /snap/go/10455/src/bytes/bytes.go:147
		} else {
//line /snap/go/10455/src/bytes/bytes.go:147
			_go_fuzz_dep_.CoverTab[524720]++
//line /snap/go/10455/src/bytes/bytes.go:147
		}
//line /snap/go/10455/src/bytes/bytes.go:147
		_go_fuzz_dep_.CoverTab[786439] = 1
//line /snap/go/10455/src/bytes/bytes.go:147
		_go_fuzz_dep_.CoverTab[200]++
							if s[i] == c {
//line /snap/go/10455/src/bytes/bytes.go:148
			_go_fuzz_dep_.CoverTab[524390]++
//line /snap/go/10455/src/bytes/bytes.go:148
			_go_fuzz_dep_.CoverTab[201]++
								return i
//line /snap/go/10455/src/bytes/bytes.go:149
			// _ = "end of CoverTab[201]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:150
			_go_fuzz_dep_.CoverTab[524391]++
//line /snap/go/10455/src/bytes/bytes.go:150
			_go_fuzz_dep_.CoverTab[202]++
//line /snap/go/10455/src/bytes/bytes.go:150
			// _ = "end of CoverTab[202]"
//line /snap/go/10455/src/bytes/bytes.go:150
		}
//line /snap/go/10455/src/bytes/bytes.go:150
		// _ = "end of CoverTab[200]"
	}
//line /snap/go/10455/src/bytes/bytes.go:151
	if _go_fuzz_dep_.CoverTab[786439] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:151
		_go_fuzz_dep_.CoverTab[524721]++
//line /snap/go/10455/src/bytes/bytes.go:151
	} else {
//line /snap/go/10455/src/bytes/bytes.go:151
		_go_fuzz_dep_.CoverTab[524722]++
//line /snap/go/10455/src/bytes/bytes.go:151
	}
//line /snap/go/10455/src/bytes/bytes.go:151
	// _ = "end of CoverTab[198]"
//line /snap/go/10455/src/bytes/bytes.go:151
	_go_fuzz_dep_.CoverTab[199]++
						return -1
//line /snap/go/10455/src/bytes/bytes.go:152
	// _ = "end of CoverTab[199]"
}

// IndexRune interprets s as a sequence of UTF-8-encoded code points.
//line /snap/go/10455/src/bytes/bytes.go:155
// It returns the byte index of the first occurrence in s of the given rune.
//line /snap/go/10455/src/bytes/bytes.go:155
// It returns -1 if rune is not present in s.
//line /snap/go/10455/src/bytes/bytes.go:155
// If r is utf8.RuneError, it returns the first instance of any
//line /snap/go/10455/src/bytes/bytes.go:155
// invalid UTF-8 byte sequence.
//line /snap/go/10455/src/bytes/bytes.go:160
func IndexRune(s []byte, r rune) int {
//line /snap/go/10455/src/bytes/bytes.go:160
	_go_fuzz_dep_.CoverTab[203]++
						switch {
	case 0 <= r && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:162
		_go_fuzz_dep_.CoverTab[209]++
//line /snap/go/10455/src/bytes/bytes.go:162
		return r < utf8.RuneSelf
//line /snap/go/10455/src/bytes/bytes.go:162
		// _ = "end of CoverTab[209]"
//line /snap/go/10455/src/bytes/bytes.go:162
	}():
//line /snap/go/10455/src/bytes/bytes.go:162
		_go_fuzz_dep_.CoverTab[524392]++
//line /snap/go/10455/src/bytes/bytes.go:162
		_go_fuzz_dep_.CoverTab[204]++
							return IndexByte(s, byte(r))
//line /snap/go/10455/src/bytes/bytes.go:163
		// _ = "end of CoverTab[204]"
	case r == utf8.RuneError:
//line /snap/go/10455/src/bytes/bytes.go:164
		_go_fuzz_dep_.CoverTab[524393]++
//line /snap/go/10455/src/bytes/bytes.go:164
		_go_fuzz_dep_.CoverTab[205]++
							for i := 0; i < len(s); {
//line /snap/go/10455/src/bytes/bytes.go:165
			_go_fuzz_dep_.CoverTab[210]++
								r1, n := utf8.DecodeRune(s[i:])
								if r1 == utf8.RuneError {
//line /snap/go/10455/src/bytes/bytes.go:167
				_go_fuzz_dep_.CoverTab[524396]++
//line /snap/go/10455/src/bytes/bytes.go:167
				_go_fuzz_dep_.CoverTab[212]++
									return i
//line /snap/go/10455/src/bytes/bytes.go:168
				// _ = "end of CoverTab[212]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:169
				_go_fuzz_dep_.CoverTab[524397]++
//line /snap/go/10455/src/bytes/bytes.go:169
				_go_fuzz_dep_.CoverTab[213]++
//line /snap/go/10455/src/bytes/bytes.go:169
				// _ = "end of CoverTab[213]"
//line /snap/go/10455/src/bytes/bytes.go:169
			}
//line /snap/go/10455/src/bytes/bytes.go:169
			// _ = "end of CoverTab[210]"
//line /snap/go/10455/src/bytes/bytes.go:169
			_go_fuzz_dep_.CoverTab[211]++
								i += n
//line /snap/go/10455/src/bytes/bytes.go:170
			// _ = "end of CoverTab[211]"
		}
//line /snap/go/10455/src/bytes/bytes.go:171
		// _ = "end of CoverTab[205]"
//line /snap/go/10455/src/bytes/bytes.go:171
		_go_fuzz_dep_.CoverTab[206]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:172
		// _ = "end of CoverTab[206]"
	case !utf8.ValidRune(r):
//line /snap/go/10455/src/bytes/bytes.go:173
		_go_fuzz_dep_.CoverTab[524394]++
//line /snap/go/10455/src/bytes/bytes.go:173
		_go_fuzz_dep_.CoverTab[207]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:174
		// _ = "end of CoverTab[207]"
	default:
//line /snap/go/10455/src/bytes/bytes.go:175
		_go_fuzz_dep_.CoverTab[524395]++
//line /snap/go/10455/src/bytes/bytes.go:175
		_go_fuzz_dep_.CoverTab[208]++
							var b [utf8.UTFMax]byte
							n := utf8.EncodeRune(b[:], r)
							return Index(s, b[:n])
//line /snap/go/10455/src/bytes/bytes.go:178
		// _ = "end of CoverTab[208]"
	}
//line /snap/go/10455/src/bytes/bytes.go:179
	// _ = "end of CoverTab[203]"
}

// IndexAny interprets s as a sequence of UTF-8-encoded Unicode code points.
//line /snap/go/10455/src/bytes/bytes.go:182
// It returns the byte index of the first occurrence in s of any of the Unicode
//line /snap/go/10455/src/bytes/bytes.go:182
// code points in chars. It returns -1 if chars is empty or if there is no code
//line /snap/go/10455/src/bytes/bytes.go:182
// point in common.
//line /snap/go/10455/src/bytes/bytes.go:186
func IndexAny(s []byte, chars string) int {
//line /snap/go/10455/src/bytes/bytes.go:186
	_go_fuzz_dep_.CoverTab[214]++
						if chars == "" {
//line /snap/go/10455/src/bytes/bytes.go:187
		_go_fuzz_dep_.CoverTab[524398]++
//line /snap/go/10455/src/bytes/bytes.go:187
		_go_fuzz_dep_.CoverTab[220]++

							return -1
//line /snap/go/10455/src/bytes/bytes.go:189
		// _ = "end of CoverTab[220]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:190
		_go_fuzz_dep_.CoverTab[524399]++
//line /snap/go/10455/src/bytes/bytes.go:190
		_go_fuzz_dep_.CoverTab[221]++
//line /snap/go/10455/src/bytes/bytes.go:190
		// _ = "end of CoverTab[221]"
//line /snap/go/10455/src/bytes/bytes.go:190
	}
//line /snap/go/10455/src/bytes/bytes.go:190
	// _ = "end of CoverTab[214]"
//line /snap/go/10455/src/bytes/bytes.go:190
	_go_fuzz_dep_.CoverTab[215]++
						if len(s) == 1 {
//line /snap/go/10455/src/bytes/bytes.go:191
		_go_fuzz_dep_.CoverTab[524400]++
//line /snap/go/10455/src/bytes/bytes.go:191
		_go_fuzz_dep_.CoverTab[222]++
							r := rune(s[0])
							if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:193
			_go_fuzz_dep_.CoverTab[524402]++
//line /snap/go/10455/src/bytes/bytes.go:193
			_go_fuzz_dep_.CoverTab[225]++
//line /snap/go/10455/src/bytes/bytes.go:193
			_go_fuzz_dep_.CoverTab[786441] = 0

								for _, r = range chars {
//line /snap/go/10455/src/bytes/bytes.go:195
				if _go_fuzz_dep_.CoverTab[786441] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:195
					_go_fuzz_dep_.CoverTab[524727]++
//line /snap/go/10455/src/bytes/bytes.go:195
				} else {
//line /snap/go/10455/src/bytes/bytes.go:195
					_go_fuzz_dep_.CoverTab[524728]++
//line /snap/go/10455/src/bytes/bytes.go:195
				}
//line /snap/go/10455/src/bytes/bytes.go:195
				_go_fuzz_dep_.CoverTab[786441] = 1
//line /snap/go/10455/src/bytes/bytes.go:195
				_go_fuzz_dep_.CoverTab[227]++
									if r == utf8.RuneError {
//line /snap/go/10455/src/bytes/bytes.go:196
					_go_fuzz_dep_.CoverTab[524404]++
//line /snap/go/10455/src/bytes/bytes.go:196
					_go_fuzz_dep_.CoverTab[228]++
										return 0
//line /snap/go/10455/src/bytes/bytes.go:197
					// _ = "end of CoverTab[228]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:198
					_go_fuzz_dep_.CoverTab[524405]++
//line /snap/go/10455/src/bytes/bytes.go:198
					_go_fuzz_dep_.CoverTab[229]++
//line /snap/go/10455/src/bytes/bytes.go:198
					// _ = "end of CoverTab[229]"
//line /snap/go/10455/src/bytes/bytes.go:198
				}
//line /snap/go/10455/src/bytes/bytes.go:198
				// _ = "end of CoverTab[227]"
			}
//line /snap/go/10455/src/bytes/bytes.go:199
			if _go_fuzz_dep_.CoverTab[786441] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:199
				_go_fuzz_dep_.CoverTab[524729]++
//line /snap/go/10455/src/bytes/bytes.go:199
			} else {
//line /snap/go/10455/src/bytes/bytes.go:199
				_go_fuzz_dep_.CoverTab[524730]++
//line /snap/go/10455/src/bytes/bytes.go:199
			}
//line /snap/go/10455/src/bytes/bytes.go:199
			// _ = "end of CoverTab[225]"
//line /snap/go/10455/src/bytes/bytes.go:199
			_go_fuzz_dep_.CoverTab[226]++
								return -1
//line /snap/go/10455/src/bytes/bytes.go:200
			// _ = "end of CoverTab[226]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:201
			_go_fuzz_dep_.CoverTab[524403]++
//line /snap/go/10455/src/bytes/bytes.go:201
			_go_fuzz_dep_.CoverTab[230]++
//line /snap/go/10455/src/bytes/bytes.go:201
			// _ = "end of CoverTab[230]"
//line /snap/go/10455/src/bytes/bytes.go:201
		}
//line /snap/go/10455/src/bytes/bytes.go:201
		// _ = "end of CoverTab[222]"
//line /snap/go/10455/src/bytes/bytes.go:201
		_go_fuzz_dep_.CoverTab[223]++
							if bytealg.IndexByteString(chars, s[0]) >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:202
			_go_fuzz_dep_.CoverTab[524406]++
//line /snap/go/10455/src/bytes/bytes.go:202
			_go_fuzz_dep_.CoverTab[231]++
								return 0
//line /snap/go/10455/src/bytes/bytes.go:203
			// _ = "end of CoverTab[231]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:204
			_go_fuzz_dep_.CoverTab[524407]++
//line /snap/go/10455/src/bytes/bytes.go:204
			_go_fuzz_dep_.CoverTab[232]++
//line /snap/go/10455/src/bytes/bytes.go:204
			// _ = "end of CoverTab[232]"
//line /snap/go/10455/src/bytes/bytes.go:204
		}
//line /snap/go/10455/src/bytes/bytes.go:204
		// _ = "end of CoverTab[223]"
//line /snap/go/10455/src/bytes/bytes.go:204
		_go_fuzz_dep_.CoverTab[224]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:205
		// _ = "end of CoverTab[224]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:206
		_go_fuzz_dep_.CoverTab[524401]++
//line /snap/go/10455/src/bytes/bytes.go:206
		_go_fuzz_dep_.CoverTab[233]++
//line /snap/go/10455/src/bytes/bytes.go:206
		// _ = "end of CoverTab[233]"
//line /snap/go/10455/src/bytes/bytes.go:206
	}
//line /snap/go/10455/src/bytes/bytes.go:206
	// _ = "end of CoverTab[215]"
//line /snap/go/10455/src/bytes/bytes.go:206
	_go_fuzz_dep_.CoverTab[216]++
						if len(chars) == 1 {
//line /snap/go/10455/src/bytes/bytes.go:207
		_go_fuzz_dep_.CoverTab[524408]++
//line /snap/go/10455/src/bytes/bytes.go:207
		_go_fuzz_dep_.CoverTab[234]++
							r := rune(chars[0])
							if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:209
			_go_fuzz_dep_.CoverTab[524410]++
//line /snap/go/10455/src/bytes/bytes.go:209
			_go_fuzz_dep_.CoverTab[236]++
								r = utf8.RuneError
//line /snap/go/10455/src/bytes/bytes.go:210
			// _ = "end of CoverTab[236]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:211
			_go_fuzz_dep_.CoverTab[524411]++
//line /snap/go/10455/src/bytes/bytes.go:211
			_go_fuzz_dep_.CoverTab[237]++
//line /snap/go/10455/src/bytes/bytes.go:211
			// _ = "end of CoverTab[237]"
//line /snap/go/10455/src/bytes/bytes.go:211
		}
//line /snap/go/10455/src/bytes/bytes.go:211
		// _ = "end of CoverTab[234]"
//line /snap/go/10455/src/bytes/bytes.go:211
		_go_fuzz_dep_.CoverTab[235]++
							return IndexRune(s, r)
//line /snap/go/10455/src/bytes/bytes.go:212
		// _ = "end of CoverTab[235]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:213
		_go_fuzz_dep_.CoverTab[524409]++
//line /snap/go/10455/src/bytes/bytes.go:213
		_go_fuzz_dep_.CoverTab[238]++
//line /snap/go/10455/src/bytes/bytes.go:213
		// _ = "end of CoverTab[238]"
//line /snap/go/10455/src/bytes/bytes.go:213
	}
//line /snap/go/10455/src/bytes/bytes.go:213
	// _ = "end of CoverTab[216]"
//line /snap/go/10455/src/bytes/bytes.go:213
	_go_fuzz_dep_.CoverTab[217]++
						if len(s) > 8 {
//line /snap/go/10455/src/bytes/bytes.go:214
		_go_fuzz_dep_.CoverTab[524412]++
//line /snap/go/10455/src/bytes/bytes.go:214
		_go_fuzz_dep_.CoverTab[239]++
							if as, isASCII := makeASCIISet(chars); isASCII {
//line /snap/go/10455/src/bytes/bytes.go:215
			_go_fuzz_dep_.CoverTab[524414]++
//line /snap/go/10455/src/bytes/bytes.go:215
			_go_fuzz_dep_.CoverTab[240]++
//line /snap/go/10455/src/bytes/bytes.go:215
			_go_fuzz_dep_.CoverTab[786442] = 0
								for i, c := range s {
//line /snap/go/10455/src/bytes/bytes.go:216
				if _go_fuzz_dep_.CoverTab[786442] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:216
					_go_fuzz_dep_.CoverTab[524731]++
//line /snap/go/10455/src/bytes/bytes.go:216
				} else {
//line /snap/go/10455/src/bytes/bytes.go:216
					_go_fuzz_dep_.CoverTab[524732]++
//line /snap/go/10455/src/bytes/bytes.go:216
				}
//line /snap/go/10455/src/bytes/bytes.go:216
				_go_fuzz_dep_.CoverTab[786442] = 1
//line /snap/go/10455/src/bytes/bytes.go:216
				_go_fuzz_dep_.CoverTab[242]++
									if as.contains(c) {
//line /snap/go/10455/src/bytes/bytes.go:217
					_go_fuzz_dep_.CoverTab[524416]++
//line /snap/go/10455/src/bytes/bytes.go:217
					_go_fuzz_dep_.CoverTab[243]++
										return i
//line /snap/go/10455/src/bytes/bytes.go:218
					// _ = "end of CoverTab[243]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:219
					_go_fuzz_dep_.CoverTab[524417]++
//line /snap/go/10455/src/bytes/bytes.go:219
					_go_fuzz_dep_.CoverTab[244]++
//line /snap/go/10455/src/bytes/bytes.go:219
					// _ = "end of CoverTab[244]"
//line /snap/go/10455/src/bytes/bytes.go:219
				}
//line /snap/go/10455/src/bytes/bytes.go:219
				// _ = "end of CoverTab[242]"
			}
//line /snap/go/10455/src/bytes/bytes.go:220
			if _go_fuzz_dep_.CoverTab[786442] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:220
				_go_fuzz_dep_.CoverTab[524733]++
//line /snap/go/10455/src/bytes/bytes.go:220
			} else {
//line /snap/go/10455/src/bytes/bytes.go:220
				_go_fuzz_dep_.CoverTab[524734]++
//line /snap/go/10455/src/bytes/bytes.go:220
			}
//line /snap/go/10455/src/bytes/bytes.go:220
			// _ = "end of CoverTab[240]"
//line /snap/go/10455/src/bytes/bytes.go:220
			_go_fuzz_dep_.CoverTab[241]++
								return -1
//line /snap/go/10455/src/bytes/bytes.go:221
			// _ = "end of CoverTab[241]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:222
			_go_fuzz_dep_.CoverTab[524415]++
//line /snap/go/10455/src/bytes/bytes.go:222
			_go_fuzz_dep_.CoverTab[245]++
//line /snap/go/10455/src/bytes/bytes.go:222
			// _ = "end of CoverTab[245]"
//line /snap/go/10455/src/bytes/bytes.go:222
		}
//line /snap/go/10455/src/bytes/bytes.go:222
		// _ = "end of CoverTab[239]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:223
		_go_fuzz_dep_.CoverTab[524413]++
//line /snap/go/10455/src/bytes/bytes.go:223
		_go_fuzz_dep_.CoverTab[246]++
//line /snap/go/10455/src/bytes/bytes.go:223
		// _ = "end of CoverTab[246]"
//line /snap/go/10455/src/bytes/bytes.go:223
	}
//line /snap/go/10455/src/bytes/bytes.go:223
	// _ = "end of CoverTab[217]"
//line /snap/go/10455/src/bytes/bytes.go:223
	_go_fuzz_dep_.CoverTab[218]++
						var width int
//line /snap/go/10455/src/bytes/bytes.go:224
	_go_fuzz_dep_.CoverTab[786440] = 0
						for i := 0; i < len(s); i += width {
//line /snap/go/10455/src/bytes/bytes.go:225
		if _go_fuzz_dep_.CoverTab[786440] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:225
			_go_fuzz_dep_.CoverTab[524723]++
//line /snap/go/10455/src/bytes/bytes.go:225
		} else {
//line /snap/go/10455/src/bytes/bytes.go:225
			_go_fuzz_dep_.CoverTab[524724]++
//line /snap/go/10455/src/bytes/bytes.go:225
		}
//line /snap/go/10455/src/bytes/bytes.go:225
		_go_fuzz_dep_.CoverTab[786440] = 1
//line /snap/go/10455/src/bytes/bytes.go:225
		_go_fuzz_dep_.CoverTab[247]++
							r := rune(s[i])
							if r < utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:227
			_go_fuzz_dep_.CoverTab[524418]++
//line /snap/go/10455/src/bytes/bytes.go:227
			_go_fuzz_dep_.CoverTab[250]++
								if bytealg.IndexByteString(chars, s[i]) >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:228
				_go_fuzz_dep_.CoverTab[524420]++
//line /snap/go/10455/src/bytes/bytes.go:228
				_go_fuzz_dep_.CoverTab[252]++
									return i
//line /snap/go/10455/src/bytes/bytes.go:229
				// _ = "end of CoverTab[252]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:230
				_go_fuzz_dep_.CoverTab[524421]++
//line /snap/go/10455/src/bytes/bytes.go:230
				_go_fuzz_dep_.CoverTab[253]++
//line /snap/go/10455/src/bytes/bytes.go:230
				// _ = "end of CoverTab[253]"
//line /snap/go/10455/src/bytes/bytes.go:230
			}
//line /snap/go/10455/src/bytes/bytes.go:230
			// _ = "end of CoverTab[250]"
//line /snap/go/10455/src/bytes/bytes.go:230
			_go_fuzz_dep_.CoverTab[251]++
								width = 1
								continue
//line /snap/go/10455/src/bytes/bytes.go:232
			// _ = "end of CoverTab[251]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:233
			_go_fuzz_dep_.CoverTab[524419]++
//line /snap/go/10455/src/bytes/bytes.go:233
			_go_fuzz_dep_.CoverTab[254]++
//line /snap/go/10455/src/bytes/bytes.go:233
			// _ = "end of CoverTab[254]"
//line /snap/go/10455/src/bytes/bytes.go:233
		}
//line /snap/go/10455/src/bytes/bytes.go:233
		// _ = "end of CoverTab[247]"
//line /snap/go/10455/src/bytes/bytes.go:233
		_go_fuzz_dep_.CoverTab[248]++
							r, width = utf8.DecodeRune(s[i:])
							if r != utf8.RuneError {
//line /snap/go/10455/src/bytes/bytes.go:235
			_go_fuzz_dep_.CoverTab[524422]++
//line /snap/go/10455/src/bytes/bytes.go:235
			_go_fuzz_dep_.CoverTab[255]++

								if len(chars) == width {
//line /snap/go/10455/src/bytes/bytes.go:237
				_go_fuzz_dep_.CoverTab[524424]++
//line /snap/go/10455/src/bytes/bytes.go:237
				_go_fuzz_dep_.CoverTab[257]++
									if chars == string(r) {
//line /snap/go/10455/src/bytes/bytes.go:238
					_go_fuzz_dep_.CoverTab[524426]++
//line /snap/go/10455/src/bytes/bytes.go:238
					_go_fuzz_dep_.CoverTab[259]++
										return i
//line /snap/go/10455/src/bytes/bytes.go:239
					// _ = "end of CoverTab[259]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:240
					_go_fuzz_dep_.CoverTab[524427]++
//line /snap/go/10455/src/bytes/bytes.go:240
					_go_fuzz_dep_.CoverTab[260]++
//line /snap/go/10455/src/bytes/bytes.go:240
					// _ = "end of CoverTab[260]"
//line /snap/go/10455/src/bytes/bytes.go:240
				}
//line /snap/go/10455/src/bytes/bytes.go:240
				// _ = "end of CoverTab[257]"
//line /snap/go/10455/src/bytes/bytes.go:240
				_go_fuzz_dep_.CoverTab[258]++
									continue
//line /snap/go/10455/src/bytes/bytes.go:241
				// _ = "end of CoverTab[258]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:242
				_go_fuzz_dep_.CoverTab[524425]++
//line /snap/go/10455/src/bytes/bytes.go:242
				_go_fuzz_dep_.CoverTab[261]++
//line /snap/go/10455/src/bytes/bytes.go:242
				// _ = "end of CoverTab[261]"
//line /snap/go/10455/src/bytes/bytes.go:242
			}
//line /snap/go/10455/src/bytes/bytes.go:242
			// _ = "end of CoverTab[255]"
//line /snap/go/10455/src/bytes/bytes.go:242
			_go_fuzz_dep_.CoverTab[256]++

								if bytealg.MaxLen >= width {
//line /snap/go/10455/src/bytes/bytes.go:244
				_go_fuzz_dep_.CoverTab[524428]++
//line /snap/go/10455/src/bytes/bytes.go:244
				_go_fuzz_dep_.CoverTab[262]++
									if bytealg.IndexString(chars, string(r)) >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:245
					_go_fuzz_dep_.CoverTab[524430]++
//line /snap/go/10455/src/bytes/bytes.go:245
					_go_fuzz_dep_.CoverTab[264]++
										return i
//line /snap/go/10455/src/bytes/bytes.go:246
					// _ = "end of CoverTab[264]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:247
					_go_fuzz_dep_.CoverTab[524431]++
//line /snap/go/10455/src/bytes/bytes.go:247
					_go_fuzz_dep_.CoverTab[265]++
//line /snap/go/10455/src/bytes/bytes.go:247
					// _ = "end of CoverTab[265]"
//line /snap/go/10455/src/bytes/bytes.go:247
				}
//line /snap/go/10455/src/bytes/bytes.go:247
				// _ = "end of CoverTab[262]"
//line /snap/go/10455/src/bytes/bytes.go:247
				_go_fuzz_dep_.CoverTab[263]++
									continue
//line /snap/go/10455/src/bytes/bytes.go:248
				// _ = "end of CoverTab[263]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:249
				_go_fuzz_dep_.CoverTab[524429]++
//line /snap/go/10455/src/bytes/bytes.go:249
				_go_fuzz_dep_.CoverTab[266]++
//line /snap/go/10455/src/bytes/bytes.go:249
				// _ = "end of CoverTab[266]"
//line /snap/go/10455/src/bytes/bytes.go:249
			}
//line /snap/go/10455/src/bytes/bytes.go:249
			// _ = "end of CoverTab[256]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:250
			_go_fuzz_dep_.CoverTab[524423]++
//line /snap/go/10455/src/bytes/bytes.go:250
			_go_fuzz_dep_.CoverTab[267]++
//line /snap/go/10455/src/bytes/bytes.go:250
			// _ = "end of CoverTab[267]"
//line /snap/go/10455/src/bytes/bytes.go:250
		}
//line /snap/go/10455/src/bytes/bytes.go:250
		// _ = "end of CoverTab[248]"
//line /snap/go/10455/src/bytes/bytes.go:250
		_go_fuzz_dep_.CoverTab[249]++
//line /snap/go/10455/src/bytes/bytes.go:250
		_go_fuzz_dep_.CoverTab[786443] = 0
							for _, ch := range chars {
//line /snap/go/10455/src/bytes/bytes.go:251
			if _go_fuzz_dep_.CoverTab[786443] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:251
				_go_fuzz_dep_.CoverTab[524735]++
//line /snap/go/10455/src/bytes/bytes.go:251
			} else {
//line /snap/go/10455/src/bytes/bytes.go:251
				_go_fuzz_dep_.CoverTab[524736]++
//line /snap/go/10455/src/bytes/bytes.go:251
			}
//line /snap/go/10455/src/bytes/bytes.go:251
			_go_fuzz_dep_.CoverTab[786443] = 1
//line /snap/go/10455/src/bytes/bytes.go:251
			_go_fuzz_dep_.CoverTab[268]++
								if r == ch {
//line /snap/go/10455/src/bytes/bytes.go:252
				_go_fuzz_dep_.CoverTab[524432]++
//line /snap/go/10455/src/bytes/bytes.go:252
				_go_fuzz_dep_.CoverTab[269]++
									return i
//line /snap/go/10455/src/bytes/bytes.go:253
				// _ = "end of CoverTab[269]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:254
				_go_fuzz_dep_.CoverTab[524433]++
//line /snap/go/10455/src/bytes/bytes.go:254
				_go_fuzz_dep_.CoverTab[270]++
//line /snap/go/10455/src/bytes/bytes.go:254
				// _ = "end of CoverTab[270]"
//line /snap/go/10455/src/bytes/bytes.go:254
			}
//line /snap/go/10455/src/bytes/bytes.go:254
			// _ = "end of CoverTab[268]"
		}
//line /snap/go/10455/src/bytes/bytes.go:255
		if _go_fuzz_dep_.CoverTab[786443] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:255
			_go_fuzz_dep_.CoverTab[524737]++
//line /snap/go/10455/src/bytes/bytes.go:255
		} else {
//line /snap/go/10455/src/bytes/bytes.go:255
			_go_fuzz_dep_.CoverTab[524738]++
//line /snap/go/10455/src/bytes/bytes.go:255
		}
//line /snap/go/10455/src/bytes/bytes.go:255
		// _ = "end of CoverTab[249]"
	}
//line /snap/go/10455/src/bytes/bytes.go:256
	if _go_fuzz_dep_.CoverTab[786440] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:256
		_go_fuzz_dep_.CoverTab[524725]++
//line /snap/go/10455/src/bytes/bytes.go:256
	} else {
//line /snap/go/10455/src/bytes/bytes.go:256
		_go_fuzz_dep_.CoverTab[524726]++
//line /snap/go/10455/src/bytes/bytes.go:256
	}
//line /snap/go/10455/src/bytes/bytes.go:256
	// _ = "end of CoverTab[218]"
//line /snap/go/10455/src/bytes/bytes.go:256
	_go_fuzz_dep_.CoverTab[219]++
						return -1
//line /snap/go/10455/src/bytes/bytes.go:257
	// _ = "end of CoverTab[219]"
}

// LastIndexAny interprets s as a sequence of UTF-8-encoded Unicode code
//line /snap/go/10455/src/bytes/bytes.go:260
// points. It returns the byte index of the last occurrence in s of any of
//line /snap/go/10455/src/bytes/bytes.go:260
// the Unicode code points in chars. It returns -1 if chars is empty or if
//line /snap/go/10455/src/bytes/bytes.go:260
// there is no code point in common.
//line /snap/go/10455/src/bytes/bytes.go:264
func LastIndexAny(s []byte, chars string) int {
//line /snap/go/10455/src/bytes/bytes.go:264
	_go_fuzz_dep_.CoverTab[271]++
						if chars == "" {
//line /snap/go/10455/src/bytes/bytes.go:265
		_go_fuzz_dep_.CoverTab[524434]++
//line /snap/go/10455/src/bytes/bytes.go:265
		_go_fuzz_dep_.CoverTab[277]++

							return -1
//line /snap/go/10455/src/bytes/bytes.go:267
		// _ = "end of CoverTab[277]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:268
		_go_fuzz_dep_.CoverTab[524435]++
//line /snap/go/10455/src/bytes/bytes.go:268
		_go_fuzz_dep_.CoverTab[278]++
//line /snap/go/10455/src/bytes/bytes.go:268
		// _ = "end of CoverTab[278]"
//line /snap/go/10455/src/bytes/bytes.go:268
	}
//line /snap/go/10455/src/bytes/bytes.go:268
	// _ = "end of CoverTab[271]"
//line /snap/go/10455/src/bytes/bytes.go:268
	_go_fuzz_dep_.CoverTab[272]++
						if len(s) > 8 {
//line /snap/go/10455/src/bytes/bytes.go:269
		_go_fuzz_dep_.CoverTab[524436]++
//line /snap/go/10455/src/bytes/bytes.go:269
		_go_fuzz_dep_.CoverTab[279]++
							if as, isASCII := makeASCIISet(chars); isASCII {
//line /snap/go/10455/src/bytes/bytes.go:270
			_go_fuzz_dep_.CoverTab[524438]++
//line /snap/go/10455/src/bytes/bytes.go:270
			_go_fuzz_dep_.CoverTab[280]++
//line /snap/go/10455/src/bytes/bytes.go:270
			_go_fuzz_dep_.CoverTab[786445] = 0
								for i := len(s) - 1; i >= 0; i-- {
//line /snap/go/10455/src/bytes/bytes.go:271
				if _go_fuzz_dep_.CoverTab[786445] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:271
					_go_fuzz_dep_.CoverTab[524743]++
//line /snap/go/10455/src/bytes/bytes.go:271
				} else {
//line /snap/go/10455/src/bytes/bytes.go:271
					_go_fuzz_dep_.CoverTab[524744]++
//line /snap/go/10455/src/bytes/bytes.go:271
				}
//line /snap/go/10455/src/bytes/bytes.go:271
				_go_fuzz_dep_.CoverTab[786445] = 1
//line /snap/go/10455/src/bytes/bytes.go:271
				_go_fuzz_dep_.CoverTab[282]++
									if as.contains(s[i]) {
//line /snap/go/10455/src/bytes/bytes.go:272
					_go_fuzz_dep_.CoverTab[524440]++
//line /snap/go/10455/src/bytes/bytes.go:272
					_go_fuzz_dep_.CoverTab[283]++
										return i
//line /snap/go/10455/src/bytes/bytes.go:273
					// _ = "end of CoverTab[283]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:274
					_go_fuzz_dep_.CoverTab[524441]++
//line /snap/go/10455/src/bytes/bytes.go:274
					_go_fuzz_dep_.CoverTab[284]++
//line /snap/go/10455/src/bytes/bytes.go:274
					// _ = "end of CoverTab[284]"
//line /snap/go/10455/src/bytes/bytes.go:274
				}
//line /snap/go/10455/src/bytes/bytes.go:274
				// _ = "end of CoverTab[282]"
			}
//line /snap/go/10455/src/bytes/bytes.go:275
			if _go_fuzz_dep_.CoverTab[786445] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:275
				_go_fuzz_dep_.CoverTab[524745]++
//line /snap/go/10455/src/bytes/bytes.go:275
			} else {
//line /snap/go/10455/src/bytes/bytes.go:275
				_go_fuzz_dep_.CoverTab[524746]++
//line /snap/go/10455/src/bytes/bytes.go:275
			}
//line /snap/go/10455/src/bytes/bytes.go:275
			// _ = "end of CoverTab[280]"
//line /snap/go/10455/src/bytes/bytes.go:275
			_go_fuzz_dep_.CoverTab[281]++
								return -1
//line /snap/go/10455/src/bytes/bytes.go:276
			// _ = "end of CoverTab[281]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:277
			_go_fuzz_dep_.CoverTab[524439]++
//line /snap/go/10455/src/bytes/bytes.go:277
			_go_fuzz_dep_.CoverTab[285]++
//line /snap/go/10455/src/bytes/bytes.go:277
			// _ = "end of CoverTab[285]"
//line /snap/go/10455/src/bytes/bytes.go:277
		}
//line /snap/go/10455/src/bytes/bytes.go:277
		// _ = "end of CoverTab[279]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:278
		_go_fuzz_dep_.CoverTab[524437]++
//line /snap/go/10455/src/bytes/bytes.go:278
		_go_fuzz_dep_.CoverTab[286]++
//line /snap/go/10455/src/bytes/bytes.go:278
		// _ = "end of CoverTab[286]"
//line /snap/go/10455/src/bytes/bytes.go:278
	}
//line /snap/go/10455/src/bytes/bytes.go:278
	// _ = "end of CoverTab[272]"
//line /snap/go/10455/src/bytes/bytes.go:278
	_go_fuzz_dep_.CoverTab[273]++
						if len(s) == 1 {
//line /snap/go/10455/src/bytes/bytes.go:279
		_go_fuzz_dep_.CoverTab[524442]++
//line /snap/go/10455/src/bytes/bytes.go:279
		_go_fuzz_dep_.CoverTab[287]++
							r := rune(s[0])
							if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:281
			_go_fuzz_dep_.CoverTab[524444]++
//line /snap/go/10455/src/bytes/bytes.go:281
			_go_fuzz_dep_.CoverTab[290]++
//line /snap/go/10455/src/bytes/bytes.go:281
			_go_fuzz_dep_.CoverTab[786446] = 0
								for _, r = range chars {
//line /snap/go/10455/src/bytes/bytes.go:282
				if _go_fuzz_dep_.CoverTab[786446] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:282
					_go_fuzz_dep_.CoverTab[524747]++
//line /snap/go/10455/src/bytes/bytes.go:282
				} else {
//line /snap/go/10455/src/bytes/bytes.go:282
					_go_fuzz_dep_.CoverTab[524748]++
//line /snap/go/10455/src/bytes/bytes.go:282
				}
//line /snap/go/10455/src/bytes/bytes.go:282
				_go_fuzz_dep_.CoverTab[786446] = 1
//line /snap/go/10455/src/bytes/bytes.go:282
				_go_fuzz_dep_.CoverTab[292]++
									if r == utf8.RuneError {
//line /snap/go/10455/src/bytes/bytes.go:283
					_go_fuzz_dep_.CoverTab[524446]++
//line /snap/go/10455/src/bytes/bytes.go:283
					_go_fuzz_dep_.CoverTab[293]++
										return 0
//line /snap/go/10455/src/bytes/bytes.go:284
					// _ = "end of CoverTab[293]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:285
					_go_fuzz_dep_.CoverTab[524447]++
//line /snap/go/10455/src/bytes/bytes.go:285
					_go_fuzz_dep_.CoverTab[294]++
//line /snap/go/10455/src/bytes/bytes.go:285
					// _ = "end of CoverTab[294]"
//line /snap/go/10455/src/bytes/bytes.go:285
				}
//line /snap/go/10455/src/bytes/bytes.go:285
				// _ = "end of CoverTab[292]"
			}
//line /snap/go/10455/src/bytes/bytes.go:286
			if _go_fuzz_dep_.CoverTab[786446] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:286
				_go_fuzz_dep_.CoverTab[524749]++
//line /snap/go/10455/src/bytes/bytes.go:286
			} else {
//line /snap/go/10455/src/bytes/bytes.go:286
				_go_fuzz_dep_.CoverTab[524750]++
//line /snap/go/10455/src/bytes/bytes.go:286
			}
//line /snap/go/10455/src/bytes/bytes.go:286
			// _ = "end of CoverTab[290]"
//line /snap/go/10455/src/bytes/bytes.go:286
			_go_fuzz_dep_.CoverTab[291]++
								return -1
//line /snap/go/10455/src/bytes/bytes.go:287
			// _ = "end of CoverTab[291]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:288
			_go_fuzz_dep_.CoverTab[524445]++
//line /snap/go/10455/src/bytes/bytes.go:288
			_go_fuzz_dep_.CoverTab[295]++
//line /snap/go/10455/src/bytes/bytes.go:288
			// _ = "end of CoverTab[295]"
//line /snap/go/10455/src/bytes/bytes.go:288
		}
//line /snap/go/10455/src/bytes/bytes.go:288
		// _ = "end of CoverTab[287]"
//line /snap/go/10455/src/bytes/bytes.go:288
		_go_fuzz_dep_.CoverTab[288]++
							if bytealg.IndexByteString(chars, s[0]) >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:289
			_go_fuzz_dep_.CoverTab[524448]++
//line /snap/go/10455/src/bytes/bytes.go:289
			_go_fuzz_dep_.CoverTab[296]++
								return 0
//line /snap/go/10455/src/bytes/bytes.go:290
			// _ = "end of CoverTab[296]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:291
			_go_fuzz_dep_.CoverTab[524449]++
//line /snap/go/10455/src/bytes/bytes.go:291
			_go_fuzz_dep_.CoverTab[297]++
//line /snap/go/10455/src/bytes/bytes.go:291
			// _ = "end of CoverTab[297]"
//line /snap/go/10455/src/bytes/bytes.go:291
		}
//line /snap/go/10455/src/bytes/bytes.go:291
		// _ = "end of CoverTab[288]"
//line /snap/go/10455/src/bytes/bytes.go:291
		_go_fuzz_dep_.CoverTab[289]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:292
		// _ = "end of CoverTab[289]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:293
		_go_fuzz_dep_.CoverTab[524443]++
//line /snap/go/10455/src/bytes/bytes.go:293
		_go_fuzz_dep_.CoverTab[298]++
//line /snap/go/10455/src/bytes/bytes.go:293
		// _ = "end of CoverTab[298]"
//line /snap/go/10455/src/bytes/bytes.go:293
	}
//line /snap/go/10455/src/bytes/bytes.go:293
	// _ = "end of CoverTab[273]"
//line /snap/go/10455/src/bytes/bytes.go:293
	_go_fuzz_dep_.CoverTab[274]++
						if len(chars) == 1 {
//line /snap/go/10455/src/bytes/bytes.go:294
		_go_fuzz_dep_.CoverTab[524450]++
//line /snap/go/10455/src/bytes/bytes.go:294
		_go_fuzz_dep_.CoverTab[299]++
							cr := rune(chars[0])
							if cr >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:296
			_go_fuzz_dep_.CoverTab[524452]++
//line /snap/go/10455/src/bytes/bytes.go:296
			_go_fuzz_dep_.CoverTab[302]++
								cr = utf8.RuneError
//line /snap/go/10455/src/bytes/bytes.go:297
			// _ = "end of CoverTab[302]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:298
			_go_fuzz_dep_.CoverTab[524453]++
//line /snap/go/10455/src/bytes/bytes.go:298
			_go_fuzz_dep_.CoverTab[303]++
//line /snap/go/10455/src/bytes/bytes.go:298
			// _ = "end of CoverTab[303]"
//line /snap/go/10455/src/bytes/bytes.go:298
		}
//line /snap/go/10455/src/bytes/bytes.go:298
		// _ = "end of CoverTab[299]"
//line /snap/go/10455/src/bytes/bytes.go:298
		_go_fuzz_dep_.CoverTab[300]++
//line /snap/go/10455/src/bytes/bytes.go:298
		_go_fuzz_dep_.CoverTab[786447] = 0
							for i := len(s); i > 0; {
//line /snap/go/10455/src/bytes/bytes.go:299
			if _go_fuzz_dep_.CoverTab[786447] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:299
				_go_fuzz_dep_.CoverTab[524751]++
//line /snap/go/10455/src/bytes/bytes.go:299
			} else {
//line /snap/go/10455/src/bytes/bytes.go:299
				_go_fuzz_dep_.CoverTab[524752]++
//line /snap/go/10455/src/bytes/bytes.go:299
			}
//line /snap/go/10455/src/bytes/bytes.go:299
			_go_fuzz_dep_.CoverTab[786447] = 1
//line /snap/go/10455/src/bytes/bytes.go:299
			_go_fuzz_dep_.CoverTab[304]++
								r, size := utf8.DecodeLastRune(s[:i])
								i -= size
								if r == cr {
//line /snap/go/10455/src/bytes/bytes.go:302
				_go_fuzz_dep_.CoverTab[524454]++
//line /snap/go/10455/src/bytes/bytes.go:302
				_go_fuzz_dep_.CoverTab[305]++
									return i
//line /snap/go/10455/src/bytes/bytes.go:303
				// _ = "end of CoverTab[305]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:304
				_go_fuzz_dep_.CoverTab[524455]++
//line /snap/go/10455/src/bytes/bytes.go:304
				_go_fuzz_dep_.CoverTab[306]++
//line /snap/go/10455/src/bytes/bytes.go:304
				// _ = "end of CoverTab[306]"
//line /snap/go/10455/src/bytes/bytes.go:304
			}
//line /snap/go/10455/src/bytes/bytes.go:304
			// _ = "end of CoverTab[304]"
		}
//line /snap/go/10455/src/bytes/bytes.go:305
		if _go_fuzz_dep_.CoverTab[786447] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:305
			_go_fuzz_dep_.CoverTab[524753]++
//line /snap/go/10455/src/bytes/bytes.go:305
		} else {
//line /snap/go/10455/src/bytes/bytes.go:305
			_go_fuzz_dep_.CoverTab[524754]++
//line /snap/go/10455/src/bytes/bytes.go:305
		}
//line /snap/go/10455/src/bytes/bytes.go:305
		// _ = "end of CoverTab[300]"
//line /snap/go/10455/src/bytes/bytes.go:305
		_go_fuzz_dep_.CoverTab[301]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:306
		// _ = "end of CoverTab[301]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:307
		_go_fuzz_dep_.CoverTab[524451]++
//line /snap/go/10455/src/bytes/bytes.go:307
		_go_fuzz_dep_.CoverTab[307]++
//line /snap/go/10455/src/bytes/bytes.go:307
		// _ = "end of CoverTab[307]"
//line /snap/go/10455/src/bytes/bytes.go:307
	}
//line /snap/go/10455/src/bytes/bytes.go:307
	// _ = "end of CoverTab[274]"
//line /snap/go/10455/src/bytes/bytes.go:307
	_go_fuzz_dep_.CoverTab[275]++
//line /snap/go/10455/src/bytes/bytes.go:307
	_go_fuzz_dep_.CoverTab[786444] = 0
						for i := len(s); i > 0; {
//line /snap/go/10455/src/bytes/bytes.go:308
		if _go_fuzz_dep_.CoverTab[786444] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:308
			_go_fuzz_dep_.CoverTab[524739]++
//line /snap/go/10455/src/bytes/bytes.go:308
		} else {
//line /snap/go/10455/src/bytes/bytes.go:308
			_go_fuzz_dep_.CoverTab[524740]++
//line /snap/go/10455/src/bytes/bytes.go:308
		}
//line /snap/go/10455/src/bytes/bytes.go:308
		_go_fuzz_dep_.CoverTab[786444] = 1
//line /snap/go/10455/src/bytes/bytes.go:308
		_go_fuzz_dep_.CoverTab[308]++
							r := rune(s[i-1])
							if r < utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:310
			_go_fuzz_dep_.CoverTab[524456]++
//line /snap/go/10455/src/bytes/bytes.go:310
			_go_fuzz_dep_.CoverTab[311]++
								if bytealg.IndexByteString(chars, s[i-1]) >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:311
				_go_fuzz_dep_.CoverTab[524458]++
//line /snap/go/10455/src/bytes/bytes.go:311
				_go_fuzz_dep_.CoverTab[313]++
									return i - 1
//line /snap/go/10455/src/bytes/bytes.go:312
				// _ = "end of CoverTab[313]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:313
				_go_fuzz_dep_.CoverTab[524459]++
//line /snap/go/10455/src/bytes/bytes.go:313
				_go_fuzz_dep_.CoverTab[314]++
//line /snap/go/10455/src/bytes/bytes.go:313
				// _ = "end of CoverTab[314]"
//line /snap/go/10455/src/bytes/bytes.go:313
			}
//line /snap/go/10455/src/bytes/bytes.go:313
			// _ = "end of CoverTab[311]"
//line /snap/go/10455/src/bytes/bytes.go:313
			_go_fuzz_dep_.CoverTab[312]++
								i--
								continue
//line /snap/go/10455/src/bytes/bytes.go:315
			// _ = "end of CoverTab[312]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:316
			_go_fuzz_dep_.CoverTab[524457]++
//line /snap/go/10455/src/bytes/bytes.go:316
			_go_fuzz_dep_.CoverTab[315]++
//line /snap/go/10455/src/bytes/bytes.go:316
			// _ = "end of CoverTab[315]"
//line /snap/go/10455/src/bytes/bytes.go:316
		}
//line /snap/go/10455/src/bytes/bytes.go:316
		// _ = "end of CoverTab[308]"
//line /snap/go/10455/src/bytes/bytes.go:316
		_go_fuzz_dep_.CoverTab[309]++
							r, size := utf8.DecodeLastRune(s[:i])
							i -= size
							if r != utf8.RuneError {
//line /snap/go/10455/src/bytes/bytes.go:319
			_go_fuzz_dep_.CoverTab[524460]++
//line /snap/go/10455/src/bytes/bytes.go:319
			_go_fuzz_dep_.CoverTab[316]++

								if len(chars) == size {
//line /snap/go/10455/src/bytes/bytes.go:321
				_go_fuzz_dep_.CoverTab[524462]++
//line /snap/go/10455/src/bytes/bytes.go:321
				_go_fuzz_dep_.CoverTab[318]++
									if chars == string(r) {
//line /snap/go/10455/src/bytes/bytes.go:322
					_go_fuzz_dep_.CoverTab[524464]++
//line /snap/go/10455/src/bytes/bytes.go:322
					_go_fuzz_dep_.CoverTab[320]++
										return i
//line /snap/go/10455/src/bytes/bytes.go:323
					// _ = "end of CoverTab[320]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:324
					_go_fuzz_dep_.CoverTab[524465]++
//line /snap/go/10455/src/bytes/bytes.go:324
					_go_fuzz_dep_.CoverTab[321]++
//line /snap/go/10455/src/bytes/bytes.go:324
					// _ = "end of CoverTab[321]"
//line /snap/go/10455/src/bytes/bytes.go:324
				}
//line /snap/go/10455/src/bytes/bytes.go:324
				// _ = "end of CoverTab[318]"
//line /snap/go/10455/src/bytes/bytes.go:324
				_go_fuzz_dep_.CoverTab[319]++
									continue
//line /snap/go/10455/src/bytes/bytes.go:325
				// _ = "end of CoverTab[319]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:326
				_go_fuzz_dep_.CoverTab[524463]++
//line /snap/go/10455/src/bytes/bytes.go:326
				_go_fuzz_dep_.CoverTab[322]++
//line /snap/go/10455/src/bytes/bytes.go:326
				// _ = "end of CoverTab[322]"
//line /snap/go/10455/src/bytes/bytes.go:326
			}
//line /snap/go/10455/src/bytes/bytes.go:326
			// _ = "end of CoverTab[316]"
//line /snap/go/10455/src/bytes/bytes.go:326
			_go_fuzz_dep_.CoverTab[317]++

								if bytealg.MaxLen >= size {
//line /snap/go/10455/src/bytes/bytes.go:328
				_go_fuzz_dep_.CoverTab[524466]++
//line /snap/go/10455/src/bytes/bytes.go:328
				_go_fuzz_dep_.CoverTab[323]++
									if bytealg.IndexString(chars, string(r)) >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:329
					_go_fuzz_dep_.CoverTab[524468]++
//line /snap/go/10455/src/bytes/bytes.go:329
					_go_fuzz_dep_.CoverTab[325]++
										return i
//line /snap/go/10455/src/bytes/bytes.go:330
					// _ = "end of CoverTab[325]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:331
					_go_fuzz_dep_.CoverTab[524469]++
//line /snap/go/10455/src/bytes/bytes.go:331
					_go_fuzz_dep_.CoverTab[326]++
//line /snap/go/10455/src/bytes/bytes.go:331
					// _ = "end of CoverTab[326]"
//line /snap/go/10455/src/bytes/bytes.go:331
				}
//line /snap/go/10455/src/bytes/bytes.go:331
				// _ = "end of CoverTab[323]"
//line /snap/go/10455/src/bytes/bytes.go:331
				_go_fuzz_dep_.CoverTab[324]++
									continue
//line /snap/go/10455/src/bytes/bytes.go:332
				// _ = "end of CoverTab[324]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:333
				_go_fuzz_dep_.CoverTab[524467]++
//line /snap/go/10455/src/bytes/bytes.go:333
				_go_fuzz_dep_.CoverTab[327]++
//line /snap/go/10455/src/bytes/bytes.go:333
				// _ = "end of CoverTab[327]"
//line /snap/go/10455/src/bytes/bytes.go:333
			}
//line /snap/go/10455/src/bytes/bytes.go:333
			// _ = "end of CoverTab[317]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:334
			_go_fuzz_dep_.CoverTab[524461]++
//line /snap/go/10455/src/bytes/bytes.go:334
			_go_fuzz_dep_.CoverTab[328]++
//line /snap/go/10455/src/bytes/bytes.go:334
			// _ = "end of CoverTab[328]"
//line /snap/go/10455/src/bytes/bytes.go:334
		}
//line /snap/go/10455/src/bytes/bytes.go:334
		// _ = "end of CoverTab[309]"
//line /snap/go/10455/src/bytes/bytes.go:334
		_go_fuzz_dep_.CoverTab[310]++
//line /snap/go/10455/src/bytes/bytes.go:334
		_go_fuzz_dep_.CoverTab[786448] = 0
							for _, ch := range chars {
//line /snap/go/10455/src/bytes/bytes.go:335
			if _go_fuzz_dep_.CoverTab[786448] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:335
				_go_fuzz_dep_.CoverTab[524755]++
//line /snap/go/10455/src/bytes/bytes.go:335
			} else {
//line /snap/go/10455/src/bytes/bytes.go:335
				_go_fuzz_dep_.CoverTab[524756]++
//line /snap/go/10455/src/bytes/bytes.go:335
			}
//line /snap/go/10455/src/bytes/bytes.go:335
			_go_fuzz_dep_.CoverTab[786448] = 1
//line /snap/go/10455/src/bytes/bytes.go:335
			_go_fuzz_dep_.CoverTab[329]++
								if r == ch {
//line /snap/go/10455/src/bytes/bytes.go:336
				_go_fuzz_dep_.CoverTab[524470]++
//line /snap/go/10455/src/bytes/bytes.go:336
				_go_fuzz_dep_.CoverTab[330]++
									return i
//line /snap/go/10455/src/bytes/bytes.go:337
				// _ = "end of CoverTab[330]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:338
				_go_fuzz_dep_.CoverTab[524471]++
//line /snap/go/10455/src/bytes/bytes.go:338
				_go_fuzz_dep_.CoverTab[331]++
//line /snap/go/10455/src/bytes/bytes.go:338
				// _ = "end of CoverTab[331]"
//line /snap/go/10455/src/bytes/bytes.go:338
			}
//line /snap/go/10455/src/bytes/bytes.go:338
			// _ = "end of CoverTab[329]"
		}
//line /snap/go/10455/src/bytes/bytes.go:339
		if _go_fuzz_dep_.CoverTab[786448] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:339
			_go_fuzz_dep_.CoverTab[524757]++
//line /snap/go/10455/src/bytes/bytes.go:339
		} else {
//line /snap/go/10455/src/bytes/bytes.go:339
			_go_fuzz_dep_.CoverTab[524758]++
//line /snap/go/10455/src/bytes/bytes.go:339
		}
//line /snap/go/10455/src/bytes/bytes.go:339
		// _ = "end of CoverTab[310]"
	}
//line /snap/go/10455/src/bytes/bytes.go:340
	if _go_fuzz_dep_.CoverTab[786444] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:340
		_go_fuzz_dep_.CoverTab[524741]++
//line /snap/go/10455/src/bytes/bytes.go:340
	} else {
//line /snap/go/10455/src/bytes/bytes.go:340
		_go_fuzz_dep_.CoverTab[524742]++
//line /snap/go/10455/src/bytes/bytes.go:340
	}
//line /snap/go/10455/src/bytes/bytes.go:340
	// _ = "end of CoverTab[275]"
//line /snap/go/10455/src/bytes/bytes.go:340
	_go_fuzz_dep_.CoverTab[276]++
						return -1
//line /snap/go/10455/src/bytes/bytes.go:341
	// _ = "end of CoverTab[276]"
}

// Generic split: splits after each instance of sep,
//line /snap/go/10455/src/bytes/bytes.go:344
// including sepSave bytes of sep in the subslices.
//line /snap/go/10455/src/bytes/bytes.go:346
func genSplit(s, sep []byte, sepSave, n int) [][]byte {
//line /snap/go/10455/src/bytes/bytes.go:346
	_go_fuzz_dep_.CoverTab[332]++
						if n == 0 {
//line /snap/go/10455/src/bytes/bytes.go:347
		_go_fuzz_dep_.CoverTab[524472]++
//line /snap/go/10455/src/bytes/bytes.go:347
		_go_fuzz_dep_.CoverTab[338]++
							return nil
//line /snap/go/10455/src/bytes/bytes.go:348
		// _ = "end of CoverTab[338]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:349
		_go_fuzz_dep_.CoverTab[524473]++
//line /snap/go/10455/src/bytes/bytes.go:349
		_go_fuzz_dep_.CoverTab[339]++
//line /snap/go/10455/src/bytes/bytes.go:349
		// _ = "end of CoverTab[339]"
//line /snap/go/10455/src/bytes/bytes.go:349
	}
//line /snap/go/10455/src/bytes/bytes.go:349
	// _ = "end of CoverTab[332]"
//line /snap/go/10455/src/bytes/bytes.go:349
	_go_fuzz_dep_.CoverTab[333]++
						if len(sep) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:350
		_go_fuzz_dep_.CoverTab[524474]++
//line /snap/go/10455/src/bytes/bytes.go:350
		_go_fuzz_dep_.CoverTab[340]++
							return explode(s, n)
//line /snap/go/10455/src/bytes/bytes.go:351
		// _ = "end of CoverTab[340]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:352
		_go_fuzz_dep_.CoverTab[524475]++
//line /snap/go/10455/src/bytes/bytes.go:352
		_go_fuzz_dep_.CoverTab[341]++
//line /snap/go/10455/src/bytes/bytes.go:352
		// _ = "end of CoverTab[341]"
//line /snap/go/10455/src/bytes/bytes.go:352
	}
//line /snap/go/10455/src/bytes/bytes.go:352
	// _ = "end of CoverTab[333]"
//line /snap/go/10455/src/bytes/bytes.go:352
	_go_fuzz_dep_.CoverTab[334]++
						if n < 0 {
//line /snap/go/10455/src/bytes/bytes.go:353
		_go_fuzz_dep_.CoverTab[524476]++
//line /snap/go/10455/src/bytes/bytes.go:353
		_go_fuzz_dep_.CoverTab[342]++
							n = Count(s, sep) + 1
//line /snap/go/10455/src/bytes/bytes.go:354
		// _ = "end of CoverTab[342]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:355
		_go_fuzz_dep_.CoverTab[524477]++
//line /snap/go/10455/src/bytes/bytes.go:355
		_go_fuzz_dep_.CoverTab[343]++
//line /snap/go/10455/src/bytes/bytes.go:355
		// _ = "end of CoverTab[343]"
//line /snap/go/10455/src/bytes/bytes.go:355
	}
//line /snap/go/10455/src/bytes/bytes.go:355
	// _ = "end of CoverTab[334]"
//line /snap/go/10455/src/bytes/bytes.go:355
	_go_fuzz_dep_.CoverTab[335]++
						if n > len(s)+1 {
//line /snap/go/10455/src/bytes/bytes.go:356
		_go_fuzz_dep_.CoverTab[524478]++
//line /snap/go/10455/src/bytes/bytes.go:356
		_go_fuzz_dep_.CoverTab[344]++
							n = len(s) + 1
//line /snap/go/10455/src/bytes/bytes.go:357
		// _ = "end of CoverTab[344]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:358
		_go_fuzz_dep_.CoverTab[524479]++
//line /snap/go/10455/src/bytes/bytes.go:358
		_go_fuzz_dep_.CoverTab[345]++
//line /snap/go/10455/src/bytes/bytes.go:358
		// _ = "end of CoverTab[345]"
//line /snap/go/10455/src/bytes/bytes.go:358
	}
//line /snap/go/10455/src/bytes/bytes.go:358
	// _ = "end of CoverTab[335]"
//line /snap/go/10455/src/bytes/bytes.go:358
	_go_fuzz_dep_.CoverTab[336]++

						a := make([][]byte, n)
						n--
						i := 0
//line /snap/go/10455/src/bytes/bytes.go:362
	_go_fuzz_dep_.CoverTab[786449] = 0
						for i < n {
//line /snap/go/10455/src/bytes/bytes.go:363
		if _go_fuzz_dep_.CoverTab[786449] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:363
			_go_fuzz_dep_.CoverTab[524759]++
//line /snap/go/10455/src/bytes/bytes.go:363
		} else {
//line /snap/go/10455/src/bytes/bytes.go:363
			_go_fuzz_dep_.CoverTab[524760]++
//line /snap/go/10455/src/bytes/bytes.go:363
		}
//line /snap/go/10455/src/bytes/bytes.go:363
		_go_fuzz_dep_.CoverTab[786449] = 1
//line /snap/go/10455/src/bytes/bytes.go:363
		_go_fuzz_dep_.CoverTab[346]++
							m := Index(s, sep)
							if m < 0 {
//line /snap/go/10455/src/bytes/bytes.go:365
			_go_fuzz_dep_.CoverTab[524480]++
//line /snap/go/10455/src/bytes/bytes.go:365
			_go_fuzz_dep_.CoverTab[348]++
								break
//line /snap/go/10455/src/bytes/bytes.go:366
			// _ = "end of CoverTab[348]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:367
			_go_fuzz_dep_.CoverTab[524481]++
//line /snap/go/10455/src/bytes/bytes.go:367
			_go_fuzz_dep_.CoverTab[349]++
//line /snap/go/10455/src/bytes/bytes.go:367
			// _ = "end of CoverTab[349]"
//line /snap/go/10455/src/bytes/bytes.go:367
		}
//line /snap/go/10455/src/bytes/bytes.go:367
		// _ = "end of CoverTab[346]"
//line /snap/go/10455/src/bytes/bytes.go:367
		_go_fuzz_dep_.CoverTab[347]++
							a[i] = s[: m+sepSave : m+sepSave]
							s = s[m+len(sep):]
							i++
//line /snap/go/10455/src/bytes/bytes.go:370
		// _ = "end of CoverTab[347]"
	}
//line /snap/go/10455/src/bytes/bytes.go:371
	if _go_fuzz_dep_.CoverTab[786449] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:371
		_go_fuzz_dep_.CoverTab[524761]++
//line /snap/go/10455/src/bytes/bytes.go:371
	} else {
//line /snap/go/10455/src/bytes/bytes.go:371
		_go_fuzz_dep_.CoverTab[524762]++
//line /snap/go/10455/src/bytes/bytes.go:371
	}
//line /snap/go/10455/src/bytes/bytes.go:371
	// _ = "end of CoverTab[336]"
//line /snap/go/10455/src/bytes/bytes.go:371
	_go_fuzz_dep_.CoverTab[337]++
						a[i] = s
						return a[:i+1]
//line /snap/go/10455/src/bytes/bytes.go:373
	// _ = "end of CoverTab[337]"
}

// SplitN slices s into subslices separated by sep and returns a slice of
//line /snap/go/10455/src/bytes/bytes.go:376
// the subslices between those separators.
//line /snap/go/10455/src/bytes/bytes.go:376
// If sep is empty, SplitN splits after each UTF-8 sequence.
//line /snap/go/10455/src/bytes/bytes.go:376
// The count determines the number of subslices to return:
//line /snap/go/10455/src/bytes/bytes.go:376
//
//line /snap/go/10455/src/bytes/bytes.go:376
//	n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//line /snap/go/10455/src/bytes/bytes.go:376
//	n == 0: the result is nil (zero subslices)
//line /snap/go/10455/src/bytes/bytes.go:376
//	n < 0: all subslices
//line /snap/go/10455/src/bytes/bytes.go:376
//
//line /snap/go/10455/src/bytes/bytes.go:376
// To split around the first instance of a separator, see Cut.
//line /snap/go/10455/src/bytes/bytes.go:386
func SplitN(s, sep []byte, n int) [][]byte {
//line /snap/go/10455/src/bytes/bytes.go:386
	_go_fuzz_dep_.CoverTab[350]++
//line /snap/go/10455/src/bytes/bytes.go:386
	return genSplit(s, sep, 0, n)
//line /snap/go/10455/src/bytes/bytes.go:386
	// _ = "end of CoverTab[350]"
//line /snap/go/10455/src/bytes/bytes.go:386
}

// SplitAfterN slices s into subslices after each instance of sep and
//line /snap/go/10455/src/bytes/bytes.go:388
// returns a slice of those subslices.
//line /snap/go/10455/src/bytes/bytes.go:388
// If sep is empty, SplitAfterN splits after each UTF-8 sequence.
//line /snap/go/10455/src/bytes/bytes.go:388
// The count determines the number of subslices to return:
//line /snap/go/10455/src/bytes/bytes.go:388
//
//line /snap/go/10455/src/bytes/bytes.go:388
//	n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//line /snap/go/10455/src/bytes/bytes.go:388
//	n == 0: the result is nil (zero subslices)
//line /snap/go/10455/src/bytes/bytes.go:388
//	n < 0: all subslices
//line /snap/go/10455/src/bytes/bytes.go:396
func SplitAfterN(s, sep []byte, n int) [][]byte {
//line /snap/go/10455/src/bytes/bytes.go:396
	_go_fuzz_dep_.CoverTab[351]++
						return genSplit(s, sep, len(sep), n)
//line /snap/go/10455/src/bytes/bytes.go:397
	// _ = "end of CoverTab[351]"
}

// Split slices s into all subslices separated by sep and returns a slice of
//line /snap/go/10455/src/bytes/bytes.go:400
// the subslices between those separators.
//line /snap/go/10455/src/bytes/bytes.go:400
// If sep is empty, Split splits after each UTF-8 sequence.
//line /snap/go/10455/src/bytes/bytes.go:400
// It is equivalent to SplitN with a count of -1.
//line /snap/go/10455/src/bytes/bytes.go:400
//
//line /snap/go/10455/src/bytes/bytes.go:400
// To split around the first instance of a separator, see Cut.
//line /snap/go/10455/src/bytes/bytes.go:406
func Split(s, sep []byte) [][]byte {
//line /snap/go/10455/src/bytes/bytes.go:406
	_go_fuzz_dep_.CoverTab[352]++
//line /snap/go/10455/src/bytes/bytes.go:406
	return genSplit(s, sep, 0, -1)
//line /snap/go/10455/src/bytes/bytes.go:406
	// _ = "end of CoverTab[352]"
//line /snap/go/10455/src/bytes/bytes.go:406
}

// SplitAfter slices s into all subslices after each instance of sep and
//line /snap/go/10455/src/bytes/bytes.go:408
// returns a slice of those subslices.
//line /snap/go/10455/src/bytes/bytes.go:408
// If sep is empty, SplitAfter splits after each UTF-8 sequence.
//line /snap/go/10455/src/bytes/bytes.go:408
// It is equivalent to SplitAfterN with a count of -1.
//line /snap/go/10455/src/bytes/bytes.go:412
func SplitAfter(s, sep []byte) [][]byte {
//line /snap/go/10455/src/bytes/bytes.go:412
	_go_fuzz_dep_.CoverTab[353]++
						return genSplit(s, sep, len(sep), -1)
//line /snap/go/10455/src/bytes/bytes.go:413
	// _ = "end of CoverTab[353]"
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// Fields interprets s as a sequence of UTF-8-encoded code points.
//line /snap/go/10455/src/bytes/bytes.go:418
// It splits the slice s around each instance of one or more consecutive white space
//line /snap/go/10455/src/bytes/bytes.go:418
// characters, as defined by unicode.IsSpace, returning a slice of subslices of s or an
//line /snap/go/10455/src/bytes/bytes.go:418
// empty slice if s contains only white space.
//line /snap/go/10455/src/bytes/bytes.go:422
func Fields(s []byte) [][]byte {
//line /snap/go/10455/src/bytes/bytes.go:422
	_go_fuzz_dep_.CoverTab[354]++

//line /snap/go/10455/src/bytes/bytes.go:425
	n := 0
						wasSpace := 1

						setBits := uint8(0)
//line /snap/go/10455/src/bytes/bytes.go:428
	_go_fuzz_dep_.CoverTab[786450] = 0
						for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/bytes/bytes.go:429
		if _go_fuzz_dep_.CoverTab[786450] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:429
			_go_fuzz_dep_.CoverTab[524763]++
//line /snap/go/10455/src/bytes/bytes.go:429
		} else {
//line /snap/go/10455/src/bytes/bytes.go:429
			_go_fuzz_dep_.CoverTab[524764]++
//line /snap/go/10455/src/bytes/bytes.go:429
		}
//line /snap/go/10455/src/bytes/bytes.go:429
		_go_fuzz_dep_.CoverTab[786450] = 1
//line /snap/go/10455/src/bytes/bytes.go:429
		_go_fuzz_dep_.CoverTab[360]++
							r := s[i]
							setBits |= r
							isSpace := int(asciiSpace[r])
							n += wasSpace & ^isSpace
							wasSpace = isSpace
//line /snap/go/10455/src/bytes/bytes.go:434
		// _ = "end of CoverTab[360]"
	}
//line /snap/go/10455/src/bytes/bytes.go:435
	if _go_fuzz_dep_.CoverTab[786450] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:435
		_go_fuzz_dep_.CoverTab[524765]++
//line /snap/go/10455/src/bytes/bytes.go:435
	} else {
//line /snap/go/10455/src/bytes/bytes.go:435
		_go_fuzz_dep_.CoverTab[524766]++
//line /snap/go/10455/src/bytes/bytes.go:435
	}
//line /snap/go/10455/src/bytes/bytes.go:435
	// _ = "end of CoverTab[354]"
//line /snap/go/10455/src/bytes/bytes.go:435
	_go_fuzz_dep_.CoverTab[355]++

						if setBits >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:437
		_go_fuzz_dep_.CoverTab[524482]++
//line /snap/go/10455/src/bytes/bytes.go:437
		_go_fuzz_dep_.CoverTab[361]++

							return FieldsFunc(s, unicode.IsSpace)
//line /snap/go/10455/src/bytes/bytes.go:439
		// _ = "end of CoverTab[361]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:440
		_go_fuzz_dep_.CoverTab[524483]++
//line /snap/go/10455/src/bytes/bytes.go:440
		_go_fuzz_dep_.CoverTab[362]++
//line /snap/go/10455/src/bytes/bytes.go:440
		// _ = "end of CoverTab[362]"
//line /snap/go/10455/src/bytes/bytes.go:440
	}
//line /snap/go/10455/src/bytes/bytes.go:440
	// _ = "end of CoverTab[355]"
//line /snap/go/10455/src/bytes/bytes.go:440
	_go_fuzz_dep_.CoverTab[356]++

//line /snap/go/10455/src/bytes/bytes.go:443
	a := make([][]byte, n)
						na := 0
						fieldStart := 0
						i := 0
//line /snap/go/10455/src/bytes/bytes.go:446
	_go_fuzz_dep_.CoverTab[786451] = 0

						for i < len(s) && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:448
		_go_fuzz_dep_.CoverTab[363]++
//line /snap/go/10455/src/bytes/bytes.go:448
		return asciiSpace[s[i]] != 0
//line /snap/go/10455/src/bytes/bytes.go:448
		// _ = "end of CoverTab[363]"
//line /snap/go/10455/src/bytes/bytes.go:448
	}() {
//line /snap/go/10455/src/bytes/bytes.go:448
		if _go_fuzz_dep_.CoverTab[786451] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:448
			_go_fuzz_dep_.CoverTab[524767]++
//line /snap/go/10455/src/bytes/bytes.go:448
		} else {
//line /snap/go/10455/src/bytes/bytes.go:448
			_go_fuzz_dep_.CoverTab[524768]++
//line /snap/go/10455/src/bytes/bytes.go:448
		}
//line /snap/go/10455/src/bytes/bytes.go:448
		_go_fuzz_dep_.CoverTab[786451] = 1
//line /snap/go/10455/src/bytes/bytes.go:448
		_go_fuzz_dep_.CoverTab[364]++
							i++
//line /snap/go/10455/src/bytes/bytes.go:449
		// _ = "end of CoverTab[364]"
	}
//line /snap/go/10455/src/bytes/bytes.go:450
	if _go_fuzz_dep_.CoverTab[786451] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:450
		_go_fuzz_dep_.CoverTab[524769]++
//line /snap/go/10455/src/bytes/bytes.go:450
	} else {
//line /snap/go/10455/src/bytes/bytes.go:450
		_go_fuzz_dep_.CoverTab[524770]++
//line /snap/go/10455/src/bytes/bytes.go:450
	}
//line /snap/go/10455/src/bytes/bytes.go:450
	// _ = "end of CoverTab[356]"
//line /snap/go/10455/src/bytes/bytes.go:450
	_go_fuzz_dep_.CoverTab[357]++
						fieldStart = i
//line /snap/go/10455/src/bytes/bytes.go:451
	_go_fuzz_dep_.CoverTab[786452] = 0
						for i < len(s) {
//line /snap/go/10455/src/bytes/bytes.go:452
		if _go_fuzz_dep_.CoverTab[786452] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:452
			_go_fuzz_dep_.CoverTab[524771]++
//line /snap/go/10455/src/bytes/bytes.go:452
		} else {
//line /snap/go/10455/src/bytes/bytes.go:452
			_go_fuzz_dep_.CoverTab[524772]++
//line /snap/go/10455/src/bytes/bytes.go:452
		}
//line /snap/go/10455/src/bytes/bytes.go:452
		_go_fuzz_dep_.CoverTab[786452] = 1
//line /snap/go/10455/src/bytes/bytes.go:452
		_go_fuzz_dep_.CoverTab[365]++
							if asciiSpace[s[i]] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:453
			_go_fuzz_dep_.CoverTab[524484]++
//line /snap/go/10455/src/bytes/bytes.go:453
			_go_fuzz_dep_.CoverTab[368]++
								i++
								continue
//line /snap/go/10455/src/bytes/bytes.go:455
			// _ = "end of CoverTab[368]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:456
			_go_fuzz_dep_.CoverTab[524485]++
//line /snap/go/10455/src/bytes/bytes.go:456
			_go_fuzz_dep_.CoverTab[369]++
//line /snap/go/10455/src/bytes/bytes.go:456
			// _ = "end of CoverTab[369]"
//line /snap/go/10455/src/bytes/bytes.go:456
		}
//line /snap/go/10455/src/bytes/bytes.go:456
		// _ = "end of CoverTab[365]"
//line /snap/go/10455/src/bytes/bytes.go:456
		_go_fuzz_dep_.CoverTab[366]++
							a[na] = s[fieldStart:i:i]
							na++
							i++
//line /snap/go/10455/src/bytes/bytes.go:459
		_go_fuzz_dep_.CoverTab[786453] = 0

							for i < len(s) && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:461
			_go_fuzz_dep_.CoverTab[370]++
//line /snap/go/10455/src/bytes/bytes.go:461
			return asciiSpace[s[i]] != 0
//line /snap/go/10455/src/bytes/bytes.go:461
			// _ = "end of CoverTab[370]"
//line /snap/go/10455/src/bytes/bytes.go:461
		}() {
//line /snap/go/10455/src/bytes/bytes.go:461
			if _go_fuzz_dep_.CoverTab[786453] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:461
				_go_fuzz_dep_.CoverTab[524775]++
//line /snap/go/10455/src/bytes/bytes.go:461
			} else {
//line /snap/go/10455/src/bytes/bytes.go:461
				_go_fuzz_dep_.CoverTab[524776]++
//line /snap/go/10455/src/bytes/bytes.go:461
			}
//line /snap/go/10455/src/bytes/bytes.go:461
			_go_fuzz_dep_.CoverTab[786453] = 1
//line /snap/go/10455/src/bytes/bytes.go:461
			_go_fuzz_dep_.CoverTab[371]++
								i++
//line /snap/go/10455/src/bytes/bytes.go:462
			// _ = "end of CoverTab[371]"
		}
//line /snap/go/10455/src/bytes/bytes.go:463
		if _go_fuzz_dep_.CoverTab[786453] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:463
			_go_fuzz_dep_.CoverTab[524777]++
//line /snap/go/10455/src/bytes/bytes.go:463
		} else {
//line /snap/go/10455/src/bytes/bytes.go:463
			_go_fuzz_dep_.CoverTab[524778]++
//line /snap/go/10455/src/bytes/bytes.go:463
		}
//line /snap/go/10455/src/bytes/bytes.go:463
		// _ = "end of CoverTab[366]"
//line /snap/go/10455/src/bytes/bytes.go:463
		_go_fuzz_dep_.CoverTab[367]++
							fieldStart = i
//line /snap/go/10455/src/bytes/bytes.go:464
		// _ = "end of CoverTab[367]"
	}
//line /snap/go/10455/src/bytes/bytes.go:465
	if _go_fuzz_dep_.CoverTab[786452] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:465
		_go_fuzz_dep_.CoverTab[524773]++
//line /snap/go/10455/src/bytes/bytes.go:465
	} else {
//line /snap/go/10455/src/bytes/bytes.go:465
		_go_fuzz_dep_.CoverTab[524774]++
//line /snap/go/10455/src/bytes/bytes.go:465
	}
//line /snap/go/10455/src/bytes/bytes.go:465
	// _ = "end of CoverTab[357]"
//line /snap/go/10455/src/bytes/bytes.go:465
	_go_fuzz_dep_.CoverTab[358]++
						if fieldStart < len(s) {
//line /snap/go/10455/src/bytes/bytes.go:466
		_go_fuzz_dep_.CoverTab[524486]++
//line /snap/go/10455/src/bytes/bytes.go:466
		_go_fuzz_dep_.CoverTab[372]++
							a[na] = s[fieldStart:len(s):len(s)]
//line /snap/go/10455/src/bytes/bytes.go:467
		// _ = "end of CoverTab[372]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:468
		_go_fuzz_dep_.CoverTab[524487]++
//line /snap/go/10455/src/bytes/bytes.go:468
		_go_fuzz_dep_.CoverTab[373]++
//line /snap/go/10455/src/bytes/bytes.go:468
		// _ = "end of CoverTab[373]"
//line /snap/go/10455/src/bytes/bytes.go:468
	}
//line /snap/go/10455/src/bytes/bytes.go:468
	// _ = "end of CoverTab[358]"
//line /snap/go/10455/src/bytes/bytes.go:468
	_go_fuzz_dep_.CoverTab[359]++
						return a
//line /snap/go/10455/src/bytes/bytes.go:469
	// _ = "end of CoverTab[359]"
}

// FieldsFunc interprets s as a sequence of UTF-8-encoded code points.
//line /snap/go/10455/src/bytes/bytes.go:472
// It splits the slice s at each run of code points c satisfying f(c) and
//line /snap/go/10455/src/bytes/bytes.go:472
// returns a slice of subslices of s. If all code points in s satisfy f(c), or
//line /snap/go/10455/src/bytes/bytes.go:472
// len(s) == 0, an empty slice is returned.
//line /snap/go/10455/src/bytes/bytes.go:472
//
//line /snap/go/10455/src/bytes/bytes.go:472
// FieldsFunc makes no guarantees about the order in which it calls f(c)
//line /snap/go/10455/src/bytes/bytes.go:472
// and assumes that f always returns the same value for a given c.
//line /snap/go/10455/src/bytes/bytes.go:479
func FieldsFunc(s []byte, f func(rune) bool) [][]byte {
//line /snap/go/10455/src/bytes/bytes.go:479
	_go_fuzz_dep_.CoverTab[374]++
	// A span is used to record a slice of s of the form s[start:end].
	// The start index is inclusive and the end index is exclusive.
	type span struct {
		start	int
		end	int
	}
						spans := make([]span, 0, 32)

//line /snap/go/10455/src/bytes/bytes.go:492
	start := -1
//line /snap/go/10455/src/bytes/bytes.go:492
	_go_fuzz_dep_.CoverTab[786454] = 0
						for i := 0; i < len(s); {
//line /snap/go/10455/src/bytes/bytes.go:493
		if _go_fuzz_dep_.CoverTab[786454] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:493
			_go_fuzz_dep_.CoverTab[524779]++
//line /snap/go/10455/src/bytes/bytes.go:493
		} else {
//line /snap/go/10455/src/bytes/bytes.go:493
			_go_fuzz_dep_.CoverTab[524780]++
//line /snap/go/10455/src/bytes/bytes.go:493
		}
//line /snap/go/10455/src/bytes/bytes.go:493
		_go_fuzz_dep_.CoverTab[786454] = 1
//line /snap/go/10455/src/bytes/bytes.go:493
		_go_fuzz_dep_.CoverTab[378]++
							size := 1
							r := rune(s[i])
							if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:496
			_go_fuzz_dep_.CoverTab[524488]++
//line /snap/go/10455/src/bytes/bytes.go:496
			_go_fuzz_dep_.CoverTab[381]++
								r, size = utf8.DecodeRune(s[i:])
//line /snap/go/10455/src/bytes/bytes.go:497
			// _ = "end of CoverTab[381]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:498
			_go_fuzz_dep_.CoverTab[524489]++
//line /snap/go/10455/src/bytes/bytes.go:498
			_go_fuzz_dep_.CoverTab[382]++
//line /snap/go/10455/src/bytes/bytes.go:498
			// _ = "end of CoverTab[382]"
//line /snap/go/10455/src/bytes/bytes.go:498
		}
//line /snap/go/10455/src/bytes/bytes.go:498
		// _ = "end of CoverTab[378]"
//line /snap/go/10455/src/bytes/bytes.go:498
		_go_fuzz_dep_.CoverTab[379]++
							if f(r) {
//line /snap/go/10455/src/bytes/bytes.go:499
			_go_fuzz_dep_.CoverTab[524490]++
//line /snap/go/10455/src/bytes/bytes.go:499
			_go_fuzz_dep_.CoverTab[383]++
								if start >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:500
				_go_fuzz_dep_.CoverTab[524492]++
//line /snap/go/10455/src/bytes/bytes.go:500
				_go_fuzz_dep_.CoverTab[384]++
									spans = append(spans, span{start, i})
									start = -1
//line /snap/go/10455/src/bytes/bytes.go:502
				// _ = "end of CoverTab[384]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:503
				_go_fuzz_dep_.CoverTab[524493]++
//line /snap/go/10455/src/bytes/bytes.go:503
				_go_fuzz_dep_.CoverTab[385]++
//line /snap/go/10455/src/bytes/bytes.go:503
				// _ = "end of CoverTab[385]"
//line /snap/go/10455/src/bytes/bytes.go:503
			}
//line /snap/go/10455/src/bytes/bytes.go:503
			// _ = "end of CoverTab[383]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:504
			_go_fuzz_dep_.CoverTab[524491]++
//line /snap/go/10455/src/bytes/bytes.go:504
			_go_fuzz_dep_.CoverTab[386]++
								if start < 0 {
//line /snap/go/10455/src/bytes/bytes.go:505
				_go_fuzz_dep_.CoverTab[524494]++
//line /snap/go/10455/src/bytes/bytes.go:505
				_go_fuzz_dep_.CoverTab[387]++
									start = i
//line /snap/go/10455/src/bytes/bytes.go:506
				// _ = "end of CoverTab[387]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:507
				_go_fuzz_dep_.CoverTab[524495]++
//line /snap/go/10455/src/bytes/bytes.go:507
				_go_fuzz_dep_.CoverTab[388]++
//line /snap/go/10455/src/bytes/bytes.go:507
				// _ = "end of CoverTab[388]"
//line /snap/go/10455/src/bytes/bytes.go:507
			}
//line /snap/go/10455/src/bytes/bytes.go:507
			// _ = "end of CoverTab[386]"
		}
//line /snap/go/10455/src/bytes/bytes.go:508
		// _ = "end of CoverTab[379]"
//line /snap/go/10455/src/bytes/bytes.go:508
		_go_fuzz_dep_.CoverTab[380]++
							i += size
//line /snap/go/10455/src/bytes/bytes.go:509
		// _ = "end of CoverTab[380]"
	}
//line /snap/go/10455/src/bytes/bytes.go:510
	if _go_fuzz_dep_.CoverTab[786454] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:510
		_go_fuzz_dep_.CoverTab[524781]++
//line /snap/go/10455/src/bytes/bytes.go:510
	} else {
//line /snap/go/10455/src/bytes/bytes.go:510
		_go_fuzz_dep_.CoverTab[524782]++
//line /snap/go/10455/src/bytes/bytes.go:510
	}
//line /snap/go/10455/src/bytes/bytes.go:510
	// _ = "end of CoverTab[374]"
//line /snap/go/10455/src/bytes/bytes.go:510
	_go_fuzz_dep_.CoverTab[375]++

//line /snap/go/10455/src/bytes/bytes.go:513
	if start >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:513
		_go_fuzz_dep_.CoverTab[524496]++
//line /snap/go/10455/src/bytes/bytes.go:513
		_go_fuzz_dep_.CoverTab[389]++
							spans = append(spans, span{start, len(s)})
//line /snap/go/10455/src/bytes/bytes.go:514
		// _ = "end of CoverTab[389]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:515
		_go_fuzz_dep_.CoverTab[524497]++
//line /snap/go/10455/src/bytes/bytes.go:515
		_go_fuzz_dep_.CoverTab[390]++
//line /snap/go/10455/src/bytes/bytes.go:515
		// _ = "end of CoverTab[390]"
//line /snap/go/10455/src/bytes/bytes.go:515
	}
//line /snap/go/10455/src/bytes/bytes.go:515
	// _ = "end of CoverTab[375]"
//line /snap/go/10455/src/bytes/bytes.go:515
	_go_fuzz_dep_.CoverTab[376]++

//line /snap/go/10455/src/bytes/bytes.go:518
	a := make([][]byte, len(spans))
//line /snap/go/10455/src/bytes/bytes.go:518
	_go_fuzz_dep_.CoverTab[786455] = 0
						for i, span := range spans {
//line /snap/go/10455/src/bytes/bytes.go:519
		if _go_fuzz_dep_.CoverTab[786455] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:519
			_go_fuzz_dep_.CoverTab[524783]++
//line /snap/go/10455/src/bytes/bytes.go:519
		} else {
//line /snap/go/10455/src/bytes/bytes.go:519
			_go_fuzz_dep_.CoverTab[524784]++
//line /snap/go/10455/src/bytes/bytes.go:519
		}
//line /snap/go/10455/src/bytes/bytes.go:519
		_go_fuzz_dep_.CoverTab[786455] = 1
//line /snap/go/10455/src/bytes/bytes.go:519
		_go_fuzz_dep_.CoverTab[391]++
							a[i] = s[span.start:span.end:span.end]
//line /snap/go/10455/src/bytes/bytes.go:520
		// _ = "end of CoverTab[391]"
	}
//line /snap/go/10455/src/bytes/bytes.go:521
	if _go_fuzz_dep_.CoverTab[786455] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:521
		_go_fuzz_dep_.CoverTab[524785]++
//line /snap/go/10455/src/bytes/bytes.go:521
	} else {
//line /snap/go/10455/src/bytes/bytes.go:521
		_go_fuzz_dep_.CoverTab[524786]++
//line /snap/go/10455/src/bytes/bytes.go:521
	}
//line /snap/go/10455/src/bytes/bytes.go:521
	// _ = "end of CoverTab[376]"
//line /snap/go/10455/src/bytes/bytes.go:521
	_go_fuzz_dep_.CoverTab[377]++

						return a
//line /snap/go/10455/src/bytes/bytes.go:523
	// _ = "end of CoverTab[377]"
}

// Join concatenates the elements of s to create a new byte slice. The separator
//line /snap/go/10455/src/bytes/bytes.go:526
// sep is placed between elements in the resulting slice.
//line /snap/go/10455/src/bytes/bytes.go:528
func Join(s [][]byte, sep []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:528
	_go_fuzz_dep_.CoverTab[392]++
						if len(s) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:529
		_go_fuzz_dep_.CoverTab[524498]++
//line /snap/go/10455/src/bytes/bytes.go:529
		_go_fuzz_dep_.CoverTab[398]++
							return []byte{}
//line /snap/go/10455/src/bytes/bytes.go:530
		// _ = "end of CoverTab[398]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:531
		_go_fuzz_dep_.CoverTab[524499]++
//line /snap/go/10455/src/bytes/bytes.go:531
		_go_fuzz_dep_.CoverTab[399]++
//line /snap/go/10455/src/bytes/bytes.go:531
		// _ = "end of CoverTab[399]"
//line /snap/go/10455/src/bytes/bytes.go:531
	}
//line /snap/go/10455/src/bytes/bytes.go:531
	// _ = "end of CoverTab[392]"
//line /snap/go/10455/src/bytes/bytes.go:531
	_go_fuzz_dep_.CoverTab[393]++
						if len(s) == 1 {
//line /snap/go/10455/src/bytes/bytes.go:532
		_go_fuzz_dep_.CoverTab[524500]++
//line /snap/go/10455/src/bytes/bytes.go:532
		_go_fuzz_dep_.CoverTab[400]++

							return append([]byte(nil), s[0]...)
//line /snap/go/10455/src/bytes/bytes.go:534
		// _ = "end of CoverTab[400]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:535
		_go_fuzz_dep_.CoverTab[524501]++
//line /snap/go/10455/src/bytes/bytes.go:535
		_go_fuzz_dep_.CoverTab[401]++
//line /snap/go/10455/src/bytes/bytes.go:535
		// _ = "end of CoverTab[401]"
//line /snap/go/10455/src/bytes/bytes.go:535
	}
//line /snap/go/10455/src/bytes/bytes.go:535
	// _ = "end of CoverTab[393]"
//line /snap/go/10455/src/bytes/bytes.go:535
	_go_fuzz_dep_.CoverTab[394]++

						var n int
						if len(sep) > 0 {
//line /snap/go/10455/src/bytes/bytes.go:538
		_go_fuzz_dep_.CoverTab[524502]++
//line /snap/go/10455/src/bytes/bytes.go:538
		_go_fuzz_dep_.CoverTab[402]++
							if len(sep) >= maxInt/(len(s)-1) {
//line /snap/go/10455/src/bytes/bytes.go:539
			_go_fuzz_dep_.CoverTab[524504]++
//line /snap/go/10455/src/bytes/bytes.go:539
			_go_fuzz_dep_.CoverTab[404]++
								panic("bytes: Join output length overflow")
//line /snap/go/10455/src/bytes/bytes.go:540
			// _ = "end of CoverTab[404]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:541
			_go_fuzz_dep_.CoverTab[524505]++
//line /snap/go/10455/src/bytes/bytes.go:541
			_go_fuzz_dep_.CoverTab[405]++
//line /snap/go/10455/src/bytes/bytes.go:541
			// _ = "end of CoverTab[405]"
//line /snap/go/10455/src/bytes/bytes.go:541
		}
//line /snap/go/10455/src/bytes/bytes.go:541
		// _ = "end of CoverTab[402]"
//line /snap/go/10455/src/bytes/bytes.go:541
		_go_fuzz_dep_.CoverTab[403]++
							n += len(sep) * (len(s) - 1)
//line /snap/go/10455/src/bytes/bytes.go:542
		// _ = "end of CoverTab[403]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:543
		_go_fuzz_dep_.CoverTab[524503]++
//line /snap/go/10455/src/bytes/bytes.go:543
		_go_fuzz_dep_.CoverTab[406]++
//line /snap/go/10455/src/bytes/bytes.go:543
		// _ = "end of CoverTab[406]"
//line /snap/go/10455/src/bytes/bytes.go:543
	}
//line /snap/go/10455/src/bytes/bytes.go:543
	// _ = "end of CoverTab[394]"
//line /snap/go/10455/src/bytes/bytes.go:543
	_go_fuzz_dep_.CoverTab[395]++
//line /snap/go/10455/src/bytes/bytes.go:543
	_go_fuzz_dep_.CoverTab[786456] = 0
						for _, v := range s {
//line /snap/go/10455/src/bytes/bytes.go:544
		if _go_fuzz_dep_.CoverTab[786456] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:544
			_go_fuzz_dep_.CoverTab[524787]++
//line /snap/go/10455/src/bytes/bytes.go:544
		} else {
//line /snap/go/10455/src/bytes/bytes.go:544
			_go_fuzz_dep_.CoverTab[524788]++
//line /snap/go/10455/src/bytes/bytes.go:544
		}
//line /snap/go/10455/src/bytes/bytes.go:544
		_go_fuzz_dep_.CoverTab[786456] = 1
//line /snap/go/10455/src/bytes/bytes.go:544
		_go_fuzz_dep_.CoverTab[407]++
							if len(v) > maxInt-n {
//line /snap/go/10455/src/bytes/bytes.go:545
			_go_fuzz_dep_.CoverTab[524506]++
//line /snap/go/10455/src/bytes/bytes.go:545
			_go_fuzz_dep_.CoverTab[409]++
								panic("bytes: Join output length overflow")
//line /snap/go/10455/src/bytes/bytes.go:546
			// _ = "end of CoverTab[409]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:547
			_go_fuzz_dep_.CoverTab[524507]++
//line /snap/go/10455/src/bytes/bytes.go:547
			_go_fuzz_dep_.CoverTab[410]++
//line /snap/go/10455/src/bytes/bytes.go:547
			// _ = "end of CoverTab[410]"
//line /snap/go/10455/src/bytes/bytes.go:547
		}
//line /snap/go/10455/src/bytes/bytes.go:547
		// _ = "end of CoverTab[407]"
//line /snap/go/10455/src/bytes/bytes.go:547
		_go_fuzz_dep_.CoverTab[408]++
							n += len(v)
//line /snap/go/10455/src/bytes/bytes.go:548
		// _ = "end of CoverTab[408]"
	}
//line /snap/go/10455/src/bytes/bytes.go:549
	if _go_fuzz_dep_.CoverTab[786456] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:549
		_go_fuzz_dep_.CoverTab[524789]++
//line /snap/go/10455/src/bytes/bytes.go:549
	} else {
//line /snap/go/10455/src/bytes/bytes.go:549
		_go_fuzz_dep_.CoverTab[524790]++
//line /snap/go/10455/src/bytes/bytes.go:549
	}
//line /snap/go/10455/src/bytes/bytes.go:549
	// _ = "end of CoverTab[395]"
//line /snap/go/10455/src/bytes/bytes.go:549
	_go_fuzz_dep_.CoverTab[396]++

						b := bytealg.MakeNoZero(n)
						bp := copy(b, s[0])
//line /snap/go/10455/src/bytes/bytes.go:552
	_go_fuzz_dep_.CoverTab[786457] = 0
						for _, v := range s[1:] {
//line /snap/go/10455/src/bytes/bytes.go:553
		if _go_fuzz_dep_.CoverTab[786457] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:553
			_go_fuzz_dep_.CoverTab[524791]++
//line /snap/go/10455/src/bytes/bytes.go:553
		} else {
//line /snap/go/10455/src/bytes/bytes.go:553
			_go_fuzz_dep_.CoverTab[524792]++
//line /snap/go/10455/src/bytes/bytes.go:553
		}
//line /snap/go/10455/src/bytes/bytes.go:553
		_go_fuzz_dep_.CoverTab[786457] = 1
//line /snap/go/10455/src/bytes/bytes.go:553
		_go_fuzz_dep_.CoverTab[411]++
							bp += copy(b[bp:], sep)
							bp += copy(b[bp:], v)
//line /snap/go/10455/src/bytes/bytes.go:555
		// _ = "end of CoverTab[411]"
	}
//line /snap/go/10455/src/bytes/bytes.go:556
	if _go_fuzz_dep_.CoverTab[786457] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:556
		_go_fuzz_dep_.CoverTab[524793]++
//line /snap/go/10455/src/bytes/bytes.go:556
	} else {
//line /snap/go/10455/src/bytes/bytes.go:556
		_go_fuzz_dep_.CoverTab[524794]++
//line /snap/go/10455/src/bytes/bytes.go:556
	}
//line /snap/go/10455/src/bytes/bytes.go:556
	// _ = "end of CoverTab[396]"
//line /snap/go/10455/src/bytes/bytes.go:556
	_go_fuzz_dep_.CoverTab[397]++
						return b
//line /snap/go/10455/src/bytes/bytes.go:557
	// _ = "end of CoverTab[397]"
}

// HasPrefix tests whether the byte slice s begins with prefix.
func HasPrefix(s, prefix []byte) bool {
//line /snap/go/10455/src/bytes/bytes.go:561
	_go_fuzz_dep_.CoverTab[412]++
						return len(s) >= len(prefix) && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:562
		_go_fuzz_dep_.CoverTab[413]++
//line /snap/go/10455/src/bytes/bytes.go:562
		return Equal(s[0:len(prefix)], prefix)
//line /snap/go/10455/src/bytes/bytes.go:562
		// _ = "end of CoverTab[413]"
//line /snap/go/10455/src/bytes/bytes.go:562
	}()
//line /snap/go/10455/src/bytes/bytes.go:562
	// _ = "end of CoverTab[412]"
}

// HasSuffix tests whether the byte slice s ends with suffix.
func HasSuffix(s, suffix []byte) bool {
//line /snap/go/10455/src/bytes/bytes.go:566
	_go_fuzz_dep_.CoverTab[414]++
						return len(s) >= len(suffix) && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:567
		_go_fuzz_dep_.CoverTab[415]++
//line /snap/go/10455/src/bytes/bytes.go:567
		return Equal(s[len(s)-len(suffix):], suffix)
//line /snap/go/10455/src/bytes/bytes.go:567
		// _ = "end of CoverTab[415]"
//line /snap/go/10455/src/bytes/bytes.go:567
	}()
//line /snap/go/10455/src/bytes/bytes.go:567
	// _ = "end of CoverTab[414]"
}

// Map returns a copy of the byte slice s with all its characters modified
//line /snap/go/10455/src/bytes/bytes.go:570
// according to the mapping function. If mapping returns a negative value, the character is
//line /snap/go/10455/src/bytes/bytes.go:570
// dropped from the byte slice with no replacement. The characters in s and the
//line /snap/go/10455/src/bytes/bytes.go:570
// output are interpreted as UTF-8-encoded code points.
//line /snap/go/10455/src/bytes/bytes.go:574
func Map(mapping func(r rune) rune, s []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:574
	_go_fuzz_dep_.CoverTab[416]++

//line /snap/go/10455/src/bytes/bytes.go:578
	b := make([]byte, 0, len(s))
//line /snap/go/10455/src/bytes/bytes.go:578
	_go_fuzz_dep_.CoverTab[786458] = 0
						for i := 0; i < len(s); {
//line /snap/go/10455/src/bytes/bytes.go:579
		if _go_fuzz_dep_.CoverTab[786458] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:579
			_go_fuzz_dep_.CoverTab[524795]++
//line /snap/go/10455/src/bytes/bytes.go:579
		} else {
//line /snap/go/10455/src/bytes/bytes.go:579
			_go_fuzz_dep_.CoverTab[524796]++
//line /snap/go/10455/src/bytes/bytes.go:579
		}
//line /snap/go/10455/src/bytes/bytes.go:579
		_go_fuzz_dep_.CoverTab[786458] = 1
//line /snap/go/10455/src/bytes/bytes.go:579
		_go_fuzz_dep_.CoverTab[418]++
							wid := 1
							r := rune(s[i])
							if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:582
			_go_fuzz_dep_.CoverTab[524508]++
//line /snap/go/10455/src/bytes/bytes.go:582
			_go_fuzz_dep_.CoverTab[421]++
								r, wid = utf8.DecodeRune(s[i:])
//line /snap/go/10455/src/bytes/bytes.go:583
			// _ = "end of CoverTab[421]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:584
			_go_fuzz_dep_.CoverTab[524509]++
//line /snap/go/10455/src/bytes/bytes.go:584
			_go_fuzz_dep_.CoverTab[422]++
//line /snap/go/10455/src/bytes/bytes.go:584
			// _ = "end of CoverTab[422]"
//line /snap/go/10455/src/bytes/bytes.go:584
		}
//line /snap/go/10455/src/bytes/bytes.go:584
		// _ = "end of CoverTab[418]"
//line /snap/go/10455/src/bytes/bytes.go:584
		_go_fuzz_dep_.CoverTab[419]++
							r = mapping(r)
							if r >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:586
			_go_fuzz_dep_.CoverTab[524510]++
//line /snap/go/10455/src/bytes/bytes.go:586
			_go_fuzz_dep_.CoverTab[423]++
								b = utf8.AppendRune(b, r)
//line /snap/go/10455/src/bytes/bytes.go:587
			// _ = "end of CoverTab[423]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:588
			_go_fuzz_dep_.CoverTab[524511]++
//line /snap/go/10455/src/bytes/bytes.go:588
			_go_fuzz_dep_.CoverTab[424]++
//line /snap/go/10455/src/bytes/bytes.go:588
			// _ = "end of CoverTab[424]"
//line /snap/go/10455/src/bytes/bytes.go:588
		}
//line /snap/go/10455/src/bytes/bytes.go:588
		// _ = "end of CoverTab[419]"
//line /snap/go/10455/src/bytes/bytes.go:588
		_go_fuzz_dep_.CoverTab[420]++
							i += wid
//line /snap/go/10455/src/bytes/bytes.go:589
		// _ = "end of CoverTab[420]"
	}
//line /snap/go/10455/src/bytes/bytes.go:590
	if _go_fuzz_dep_.CoverTab[786458] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:590
		_go_fuzz_dep_.CoverTab[524797]++
//line /snap/go/10455/src/bytes/bytes.go:590
	} else {
//line /snap/go/10455/src/bytes/bytes.go:590
		_go_fuzz_dep_.CoverTab[524798]++
//line /snap/go/10455/src/bytes/bytes.go:590
	}
//line /snap/go/10455/src/bytes/bytes.go:590
	// _ = "end of CoverTab[416]"
//line /snap/go/10455/src/bytes/bytes.go:590
	_go_fuzz_dep_.CoverTab[417]++
						return b
//line /snap/go/10455/src/bytes/bytes.go:591
	// _ = "end of CoverTab[417]"
}

// Repeat returns a new byte slice consisting of count copies of b.
//line /snap/go/10455/src/bytes/bytes.go:594
//
//line /snap/go/10455/src/bytes/bytes.go:594
// It panics if count is negative or if the result of (len(b) * count)
//line /snap/go/10455/src/bytes/bytes.go:594
// overflows.
//line /snap/go/10455/src/bytes/bytes.go:598
func Repeat(b []byte, count int) []byte {
//line /snap/go/10455/src/bytes/bytes.go:598
	_go_fuzz_dep_.CoverTab[425]++
						if count == 0 {
//line /snap/go/10455/src/bytes/bytes.go:599
		_go_fuzz_dep_.CoverTab[524512]++
//line /snap/go/10455/src/bytes/bytes.go:599
		_go_fuzz_dep_.CoverTab[432]++
							return []byte{}
//line /snap/go/10455/src/bytes/bytes.go:600
		// _ = "end of CoverTab[432]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:601
		_go_fuzz_dep_.CoverTab[524513]++
//line /snap/go/10455/src/bytes/bytes.go:601
		_go_fuzz_dep_.CoverTab[433]++
//line /snap/go/10455/src/bytes/bytes.go:601
		// _ = "end of CoverTab[433]"
//line /snap/go/10455/src/bytes/bytes.go:601
	}
//line /snap/go/10455/src/bytes/bytes.go:601
	// _ = "end of CoverTab[425]"
//line /snap/go/10455/src/bytes/bytes.go:601
	_go_fuzz_dep_.CoverTab[426]++

//line /snap/go/10455/src/bytes/bytes.go:606
	if count < 0 {
//line /snap/go/10455/src/bytes/bytes.go:606
		_go_fuzz_dep_.CoverTab[524514]++
//line /snap/go/10455/src/bytes/bytes.go:606
		_go_fuzz_dep_.CoverTab[434]++
							panic("bytes: negative Repeat count")
//line /snap/go/10455/src/bytes/bytes.go:607
		// _ = "end of CoverTab[434]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:608
		_go_fuzz_dep_.CoverTab[524515]++
//line /snap/go/10455/src/bytes/bytes.go:608
		_go_fuzz_dep_.CoverTab[435]++
//line /snap/go/10455/src/bytes/bytes.go:608
		// _ = "end of CoverTab[435]"
//line /snap/go/10455/src/bytes/bytes.go:608
	}
//line /snap/go/10455/src/bytes/bytes.go:608
	// _ = "end of CoverTab[426]"
//line /snap/go/10455/src/bytes/bytes.go:608
	_go_fuzz_dep_.CoverTab[427]++
						if len(b) >= maxInt/count {
//line /snap/go/10455/src/bytes/bytes.go:609
		_go_fuzz_dep_.CoverTab[524516]++
//line /snap/go/10455/src/bytes/bytes.go:609
		_go_fuzz_dep_.CoverTab[436]++
							panic("bytes: Repeat output length overflow")
//line /snap/go/10455/src/bytes/bytes.go:610
		// _ = "end of CoverTab[436]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:611
		_go_fuzz_dep_.CoverTab[524517]++
//line /snap/go/10455/src/bytes/bytes.go:611
		_go_fuzz_dep_.CoverTab[437]++
//line /snap/go/10455/src/bytes/bytes.go:611
		// _ = "end of CoverTab[437]"
//line /snap/go/10455/src/bytes/bytes.go:611
	}
//line /snap/go/10455/src/bytes/bytes.go:611
	// _ = "end of CoverTab[427]"
//line /snap/go/10455/src/bytes/bytes.go:611
	_go_fuzz_dep_.CoverTab[428]++
						n := len(b) * count

						if len(b) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:614
		_go_fuzz_dep_.CoverTab[524518]++
//line /snap/go/10455/src/bytes/bytes.go:614
		_go_fuzz_dep_.CoverTab[438]++
							return []byte{}
//line /snap/go/10455/src/bytes/bytes.go:615
		// _ = "end of CoverTab[438]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:616
		_go_fuzz_dep_.CoverTab[524519]++
//line /snap/go/10455/src/bytes/bytes.go:616
		_go_fuzz_dep_.CoverTab[439]++
//line /snap/go/10455/src/bytes/bytes.go:616
		// _ = "end of CoverTab[439]"
//line /snap/go/10455/src/bytes/bytes.go:616
	}
//line /snap/go/10455/src/bytes/bytes.go:616
	// _ = "end of CoverTab[428]"
//line /snap/go/10455/src/bytes/bytes.go:616
	_go_fuzz_dep_.CoverTab[429]++

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
//line /snap/go/10455/src/bytes/bytes.go:630
		_go_fuzz_dep_.CoverTab[524520]++
//line /snap/go/10455/src/bytes/bytes.go:630
		_go_fuzz_dep_.CoverTab[440]++
							chunkMax = chunkLimit / len(b) * len(b)
							if chunkMax == 0 {
//line /snap/go/10455/src/bytes/bytes.go:632
			_go_fuzz_dep_.CoverTab[524522]++
//line /snap/go/10455/src/bytes/bytes.go:632
			_go_fuzz_dep_.CoverTab[441]++
								chunkMax = len(b)
//line /snap/go/10455/src/bytes/bytes.go:633
			// _ = "end of CoverTab[441]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:634
			_go_fuzz_dep_.CoverTab[524523]++
//line /snap/go/10455/src/bytes/bytes.go:634
			_go_fuzz_dep_.CoverTab[442]++
//line /snap/go/10455/src/bytes/bytes.go:634
			// _ = "end of CoverTab[442]"
//line /snap/go/10455/src/bytes/bytes.go:634
		}
//line /snap/go/10455/src/bytes/bytes.go:634
		// _ = "end of CoverTab[440]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:635
		_go_fuzz_dep_.CoverTab[524521]++
//line /snap/go/10455/src/bytes/bytes.go:635
		_go_fuzz_dep_.CoverTab[443]++
//line /snap/go/10455/src/bytes/bytes.go:635
		// _ = "end of CoverTab[443]"
//line /snap/go/10455/src/bytes/bytes.go:635
	}
//line /snap/go/10455/src/bytes/bytes.go:635
	// _ = "end of CoverTab[429]"
//line /snap/go/10455/src/bytes/bytes.go:635
	_go_fuzz_dep_.CoverTab[430]++
						nb := bytealg.MakeNoZero(n)
						bp := copy(nb, b)
//line /snap/go/10455/src/bytes/bytes.go:637
	_go_fuzz_dep_.CoverTab[786459] = 0
						for bp < n {
//line /snap/go/10455/src/bytes/bytes.go:638
		if _go_fuzz_dep_.CoverTab[786459] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:638
			_go_fuzz_dep_.CoverTab[524799]++
//line /snap/go/10455/src/bytes/bytes.go:638
		} else {
//line /snap/go/10455/src/bytes/bytes.go:638
			_go_fuzz_dep_.CoverTab[524800]++
//line /snap/go/10455/src/bytes/bytes.go:638
		}
//line /snap/go/10455/src/bytes/bytes.go:638
		_go_fuzz_dep_.CoverTab[786459] = 1
//line /snap/go/10455/src/bytes/bytes.go:638
		_go_fuzz_dep_.CoverTab[444]++
							chunk := bp
							if chunk > chunkMax {
//line /snap/go/10455/src/bytes/bytes.go:640
			_go_fuzz_dep_.CoverTab[524524]++
//line /snap/go/10455/src/bytes/bytes.go:640
			_go_fuzz_dep_.CoverTab[446]++
								chunk = chunkMax
//line /snap/go/10455/src/bytes/bytes.go:641
			// _ = "end of CoverTab[446]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:642
			_go_fuzz_dep_.CoverTab[524525]++
//line /snap/go/10455/src/bytes/bytes.go:642
			_go_fuzz_dep_.CoverTab[447]++
//line /snap/go/10455/src/bytes/bytes.go:642
			// _ = "end of CoverTab[447]"
//line /snap/go/10455/src/bytes/bytes.go:642
		}
//line /snap/go/10455/src/bytes/bytes.go:642
		// _ = "end of CoverTab[444]"
//line /snap/go/10455/src/bytes/bytes.go:642
		_go_fuzz_dep_.CoverTab[445]++
							bp += copy(nb[bp:], nb[:chunk])
//line /snap/go/10455/src/bytes/bytes.go:643
		// _ = "end of CoverTab[445]"
	}
//line /snap/go/10455/src/bytes/bytes.go:644
	if _go_fuzz_dep_.CoverTab[786459] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:644
		_go_fuzz_dep_.CoverTab[524801]++
//line /snap/go/10455/src/bytes/bytes.go:644
	} else {
//line /snap/go/10455/src/bytes/bytes.go:644
		_go_fuzz_dep_.CoverTab[524802]++
//line /snap/go/10455/src/bytes/bytes.go:644
	}
//line /snap/go/10455/src/bytes/bytes.go:644
	// _ = "end of CoverTab[430]"
//line /snap/go/10455/src/bytes/bytes.go:644
	_go_fuzz_dep_.CoverTab[431]++
						return nb
//line /snap/go/10455/src/bytes/bytes.go:645
	// _ = "end of CoverTab[431]"
}

// ToUpper returns a copy of the byte slice s with all Unicode letters mapped to
//line /snap/go/10455/src/bytes/bytes.go:648
// their upper case.
//line /snap/go/10455/src/bytes/bytes.go:650
func ToUpper(s []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:650
	_go_fuzz_dep_.CoverTab[448]++
						isASCII, hasLower := true, false
//line /snap/go/10455/src/bytes/bytes.go:651
	_go_fuzz_dep_.CoverTab[786460] = 0
						for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/bytes/bytes.go:652
		if _go_fuzz_dep_.CoverTab[786460] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:652
			_go_fuzz_dep_.CoverTab[524803]++
//line /snap/go/10455/src/bytes/bytes.go:652
		} else {
//line /snap/go/10455/src/bytes/bytes.go:652
			_go_fuzz_dep_.CoverTab[524804]++
//line /snap/go/10455/src/bytes/bytes.go:652
		}
//line /snap/go/10455/src/bytes/bytes.go:652
		_go_fuzz_dep_.CoverTab[786460] = 1
//line /snap/go/10455/src/bytes/bytes.go:652
		_go_fuzz_dep_.CoverTab[451]++
							c := s[i]
							if c >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:654
			_go_fuzz_dep_.CoverTab[524526]++
//line /snap/go/10455/src/bytes/bytes.go:654
			_go_fuzz_dep_.CoverTab[453]++
								isASCII = false
								break
//line /snap/go/10455/src/bytes/bytes.go:656
			// _ = "end of CoverTab[453]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:657
			_go_fuzz_dep_.CoverTab[524527]++
//line /snap/go/10455/src/bytes/bytes.go:657
			_go_fuzz_dep_.CoverTab[454]++
//line /snap/go/10455/src/bytes/bytes.go:657
			// _ = "end of CoverTab[454]"
//line /snap/go/10455/src/bytes/bytes.go:657
		}
//line /snap/go/10455/src/bytes/bytes.go:657
		// _ = "end of CoverTab[451]"
//line /snap/go/10455/src/bytes/bytes.go:657
		_go_fuzz_dep_.CoverTab[452]++
							hasLower = hasLower || func() bool {
//line /snap/go/10455/src/bytes/bytes.go:658
			_go_fuzz_dep_.CoverTab[455]++
//line /snap/go/10455/src/bytes/bytes.go:658
			return ('a' <= c && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:658
				_go_fuzz_dep_.CoverTab[456]++
//line /snap/go/10455/src/bytes/bytes.go:658
				return c <= 'z'
//line /snap/go/10455/src/bytes/bytes.go:658
				// _ = "end of CoverTab[456]"
//line /snap/go/10455/src/bytes/bytes.go:658
			}())
//line /snap/go/10455/src/bytes/bytes.go:658
			// _ = "end of CoverTab[455]"
//line /snap/go/10455/src/bytes/bytes.go:658
		}()
//line /snap/go/10455/src/bytes/bytes.go:658
		// _ = "end of CoverTab[452]"
	}
//line /snap/go/10455/src/bytes/bytes.go:659
	if _go_fuzz_dep_.CoverTab[786460] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:659
		_go_fuzz_dep_.CoverTab[524805]++
//line /snap/go/10455/src/bytes/bytes.go:659
	} else {
//line /snap/go/10455/src/bytes/bytes.go:659
		_go_fuzz_dep_.CoverTab[524806]++
//line /snap/go/10455/src/bytes/bytes.go:659
	}
//line /snap/go/10455/src/bytes/bytes.go:659
	// _ = "end of CoverTab[448]"
//line /snap/go/10455/src/bytes/bytes.go:659
	_go_fuzz_dep_.CoverTab[449]++

						if isASCII {
//line /snap/go/10455/src/bytes/bytes.go:661
		_go_fuzz_dep_.CoverTab[524528]++
//line /snap/go/10455/src/bytes/bytes.go:661
		_go_fuzz_dep_.CoverTab[457]++
							if !hasLower {
//line /snap/go/10455/src/bytes/bytes.go:662
			_go_fuzz_dep_.CoverTab[524530]++
//line /snap/go/10455/src/bytes/bytes.go:662
			_go_fuzz_dep_.CoverTab[460]++

								return append([]byte(""), s...)
//line /snap/go/10455/src/bytes/bytes.go:664
			// _ = "end of CoverTab[460]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:665
			_go_fuzz_dep_.CoverTab[524531]++
//line /snap/go/10455/src/bytes/bytes.go:665
			_go_fuzz_dep_.CoverTab[461]++
//line /snap/go/10455/src/bytes/bytes.go:665
			// _ = "end of CoverTab[461]"
//line /snap/go/10455/src/bytes/bytes.go:665
		}
//line /snap/go/10455/src/bytes/bytes.go:665
		// _ = "end of CoverTab[457]"
//line /snap/go/10455/src/bytes/bytes.go:665
		_go_fuzz_dep_.CoverTab[458]++
							b := bytealg.MakeNoZero(len(s))
//line /snap/go/10455/src/bytes/bytes.go:666
		_go_fuzz_dep_.CoverTab[786461] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/bytes/bytes.go:667
			if _go_fuzz_dep_.CoverTab[786461] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:667
				_go_fuzz_dep_.CoverTab[524807]++
//line /snap/go/10455/src/bytes/bytes.go:667
			} else {
//line /snap/go/10455/src/bytes/bytes.go:667
				_go_fuzz_dep_.CoverTab[524808]++
//line /snap/go/10455/src/bytes/bytes.go:667
			}
//line /snap/go/10455/src/bytes/bytes.go:667
			_go_fuzz_dep_.CoverTab[786461] = 1
//line /snap/go/10455/src/bytes/bytes.go:667
			_go_fuzz_dep_.CoverTab[462]++
								c := s[i]
								if 'a' <= c && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:669
				_go_fuzz_dep_.CoverTab[464]++
//line /snap/go/10455/src/bytes/bytes.go:669
				return c <= 'z'
//line /snap/go/10455/src/bytes/bytes.go:669
				// _ = "end of CoverTab[464]"
//line /snap/go/10455/src/bytes/bytes.go:669
			}() {
//line /snap/go/10455/src/bytes/bytes.go:669
				_go_fuzz_dep_.CoverTab[524532]++
//line /snap/go/10455/src/bytes/bytes.go:669
				_go_fuzz_dep_.CoverTab[465]++
									c -= 'a' - 'A'
//line /snap/go/10455/src/bytes/bytes.go:670
				// _ = "end of CoverTab[465]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:671
				_go_fuzz_dep_.CoverTab[524533]++
//line /snap/go/10455/src/bytes/bytes.go:671
				_go_fuzz_dep_.CoverTab[466]++
//line /snap/go/10455/src/bytes/bytes.go:671
				// _ = "end of CoverTab[466]"
//line /snap/go/10455/src/bytes/bytes.go:671
			}
//line /snap/go/10455/src/bytes/bytes.go:671
			// _ = "end of CoverTab[462]"
//line /snap/go/10455/src/bytes/bytes.go:671
			_go_fuzz_dep_.CoverTab[463]++
								b[i] = c
//line /snap/go/10455/src/bytes/bytes.go:672
			// _ = "end of CoverTab[463]"
		}
//line /snap/go/10455/src/bytes/bytes.go:673
		if _go_fuzz_dep_.CoverTab[786461] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:673
			_go_fuzz_dep_.CoverTab[524809]++
//line /snap/go/10455/src/bytes/bytes.go:673
		} else {
//line /snap/go/10455/src/bytes/bytes.go:673
			_go_fuzz_dep_.CoverTab[524810]++
//line /snap/go/10455/src/bytes/bytes.go:673
		}
//line /snap/go/10455/src/bytes/bytes.go:673
		// _ = "end of CoverTab[458]"
//line /snap/go/10455/src/bytes/bytes.go:673
		_go_fuzz_dep_.CoverTab[459]++
							return b
//line /snap/go/10455/src/bytes/bytes.go:674
		// _ = "end of CoverTab[459]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:675
		_go_fuzz_dep_.CoverTab[524529]++
//line /snap/go/10455/src/bytes/bytes.go:675
		_go_fuzz_dep_.CoverTab[467]++
//line /snap/go/10455/src/bytes/bytes.go:675
		// _ = "end of CoverTab[467]"
//line /snap/go/10455/src/bytes/bytes.go:675
	}
//line /snap/go/10455/src/bytes/bytes.go:675
	// _ = "end of CoverTab[449]"
//line /snap/go/10455/src/bytes/bytes.go:675
	_go_fuzz_dep_.CoverTab[450]++
						return Map(unicode.ToUpper, s)
//line /snap/go/10455/src/bytes/bytes.go:676
	// _ = "end of CoverTab[450]"
}

// ToLower returns a copy of the byte slice s with all Unicode letters mapped to
//line /snap/go/10455/src/bytes/bytes.go:679
// their lower case.
//line /snap/go/10455/src/bytes/bytes.go:681
func ToLower(s []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:681
	_go_fuzz_dep_.CoverTab[468]++
						isASCII, hasUpper := true, false
//line /snap/go/10455/src/bytes/bytes.go:682
	_go_fuzz_dep_.CoverTab[786462] = 0
						for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/bytes/bytes.go:683
		if _go_fuzz_dep_.CoverTab[786462] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:683
			_go_fuzz_dep_.CoverTab[524811]++
//line /snap/go/10455/src/bytes/bytes.go:683
		} else {
//line /snap/go/10455/src/bytes/bytes.go:683
			_go_fuzz_dep_.CoverTab[524812]++
//line /snap/go/10455/src/bytes/bytes.go:683
		}
//line /snap/go/10455/src/bytes/bytes.go:683
		_go_fuzz_dep_.CoverTab[786462] = 1
//line /snap/go/10455/src/bytes/bytes.go:683
		_go_fuzz_dep_.CoverTab[471]++
							c := s[i]
							if c >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:685
			_go_fuzz_dep_.CoverTab[524534]++
//line /snap/go/10455/src/bytes/bytes.go:685
			_go_fuzz_dep_.CoverTab[473]++
								isASCII = false
								break
//line /snap/go/10455/src/bytes/bytes.go:687
			// _ = "end of CoverTab[473]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:688
			_go_fuzz_dep_.CoverTab[524535]++
//line /snap/go/10455/src/bytes/bytes.go:688
			_go_fuzz_dep_.CoverTab[474]++
//line /snap/go/10455/src/bytes/bytes.go:688
			// _ = "end of CoverTab[474]"
//line /snap/go/10455/src/bytes/bytes.go:688
		}
//line /snap/go/10455/src/bytes/bytes.go:688
		// _ = "end of CoverTab[471]"
//line /snap/go/10455/src/bytes/bytes.go:688
		_go_fuzz_dep_.CoverTab[472]++
							hasUpper = hasUpper || func() bool {
//line /snap/go/10455/src/bytes/bytes.go:689
			_go_fuzz_dep_.CoverTab[475]++
//line /snap/go/10455/src/bytes/bytes.go:689
			return ('A' <= c && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:689
				_go_fuzz_dep_.CoverTab[476]++
//line /snap/go/10455/src/bytes/bytes.go:689
				return c <= 'Z'
//line /snap/go/10455/src/bytes/bytes.go:689
				// _ = "end of CoverTab[476]"
//line /snap/go/10455/src/bytes/bytes.go:689
			}())
//line /snap/go/10455/src/bytes/bytes.go:689
			// _ = "end of CoverTab[475]"
//line /snap/go/10455/src/bytes/bytes.go:689
		}()
//line /snap/go/10455/src/bytes/bytes.go:689
		// _ = "end of CoverTab[472]"
	}
//line /snap/go/10455/src/bytes/bytes.go:690
	if _go_fuzz_dep_.CoverTab[786462] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:690
		_go_fuzz_dep_.CoverTab[524813]++
//line /snap/go/10455/src/bytes/bytes.go:690
	} else {
//line /snap/go/10455/src/bytes/bytes.go:690
		_go_fuzz_dep_.CoverTab[524814]++
//line /snap/go/10455/src/bytes/bytes.go:690
	}
//line /snap/go/10455/src/bytes/bytes.go:690
	// _ = "end of CoverTab[468]"
//line /snap/go/10455/src/bytes/bytes.go:690
	_go_fuzz_dep_.CoverTab[469]++

						if isASCII {
//line /snap/go/10455/src/bytes/bytes.go:692
		_go_fuzz_dep_.CoverTab[524536]++
//line /snap/go/10455/src/bytes/bytes.go:692
		_go_fuzz_dep_.CoverTab[477]++
							if !hasUpper {
//line /snap/go/10455/src/bytes/bytes.go:693
			_go_fuzz_dep_.CoverTab[524538]++
//line /snap/go/10455/src/bytes/bytes.go:693
			_go_fuzz_dep_.CoverTab[480]++
								return append([]byte(""), s...)
//line /snap/go/10455/src/bytes/bytes.go:694
			// _ = "end of CoverTab[480]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:695
			_go_fuzz_dep_.CoverTab[524539]++
//line /snap/go/10455/src/bytes/bytes.go:695
			_go_fuzz_dep_.CoverTab[481]++
//line /snap/go/10455/src/bytes/bytes.go:695
			// _ = "end of CoverTab[481]"
//line /snap/go/10455/src/bytes/bytes.go:695
		}
//line /snap/go/10455/src/bytes/bytes.go:695
		// _ = "end of CoverTab[477]"
//line /snap/go/10455/src/bytes/bytes.go:695
		_go_fuzz_dep_.CoverTab[478]++
							b := bytealg.MakeNoZero(len(s))
//line /snap/go/10455/src/bytes/bytes.go:696
		_go_fuzz_dep_.CoverTab[786463] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/bytes/bytes.go:697
			if _go_fuzz_dep_.CoverTab[786463] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:697
				_go_fuzz_dep_.CoverTab[524815]++
//line /snap/go/10455/src/bytes/bytes.go:697
			} else {
//line /snap/go/10455/src/bytes/bytes.go:697
				_go_fuzz_dep_.CoverTab[524816]++
//line /snap/go/10455/src/bytes/bytes.go:697
			}
//line /snap/go/10455/src/bytes/bytes.go:697
			_go_fuzz_dep_.CoverTab[786463] = 1
//line /snap/go/10455/src/bytes/bytes.go:697
			_go_fuzz_dep_.CoverTab[482]++
								c := s[i]
								if 'A' <= c && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:699
				_go_fuzz_dep_.CoverTab[484]++
//line /snap/go/10455/src/bytes/bytes.go:699
				return c <= 'Z'
//line /snap/go/10455/src/bytes/bytes.go:699
				// _ = "end of CoverTab[484]"
//line /snap/go/10455/src/bytes/bytes.go:699
			}() {
//line /snap/go/10455/src/bytes/bytes.go:699
				_go_fuzz_dep_.CoverTab[524540]++
//line /snap/go/10455/src/bytes/bytes.go:699
				_go_fuzz_dep_.CoverTab[485]++
									c += 'a' - 'A'
//line /snap/go/10455/src/bytes/bytes.go:700
				// _ = "end of CoverTab[485]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:701
				_go_fuzz_dep_.CoverTab[524541]++
//line /snap/go/10455/src/bytes/bytes.go:701
				_go_fuzz_dep_.CoverTab[486]++
//line /snap/go/10455/src/bytes/bytes.go:701
				// _ = "end of CoverTab[486]"
//line /snap/go/10455/src/bytes/bytes.go:701
			}
//line /snap/go/10455/src/bytes/bytes.go:701
			// _ = "end of CoverTab[482]"
//line /snap/go/10455/src/bytes/bytes.go:701
			_go_fuzz_dep_.CoverTab[483]++
								b[i] = c
//line /snap/go/10455/src/bytes/bytes.go:702
			// _ = "end of CoverTab[483]"
		}
//line /snap/go/10455/src/bytes/bytes.go:703
		if _go_fuzz_dep_.CoverTab[786463] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:703
			_go_fuzz_dep_.CoverTab[524817]++
//line /snap/go/10455/src/bytes/bytes.go:703
		} else {
//line /snap/go/10455/src/bytes/bytes.go:703
			_go_fuzz_dep_.CoverTab[524818]++
//line /snap/go/10455/src/bytes/bytes.go:703
		}
//line /snap/go/10455/src/bytes/bytes.go:703
		// _ = "end of CoverTab[478]"
//line /snap/go/10455/src/bytes/bytes.go:703
		_go_fuzz_dep_.CoverTab[479]++
							return b
//line /snap/go/10455/src/bytes/bytes.go:704
		// _ = "end of CoverTab[479]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:705
		_go_fuzz_dep_.CoverTab[524537]++
//line /snap/go/10455/src/bytes/bytes.go:705
		_go_fuzz_dep_.CoverTab[487]++
//line /snap/go/10455/src/bytes/bytes.go:705
		// _ = "end of CoverTab[487]"
//line /snap/go/10455/src/bytes/bytes.go:705
	}
//line /snap/go/10455/src/bytes/bytes.go:705
	// _ = "end of CoverTab[469]"
//line /snap/go/10455/src/bytes/bytes.go:705
	_go_fuzz_dep_.CoverTab[470]++
						return Map(unicode.ToLower, s)
//line /snap/go/10455/src/bytes/bytes.go:706
	// _ = "end of CoverTab[470]"
}

// ToTitle treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their title case.
func ToTitle(s []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:710
	_go_fuzz_dep_.CoverTab[488]++
//line /snap/go/10455/src/bytes/bytes.go:710
	return Map(unicode.ToTitle, s)
//line /snap/go/10455/src/bytes/bytes.go:710
	// _ = "end of CoverTab[488]"
//line /snap/go/10455/src/bytes/bytes.go:710
}

// ToUpperSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
//line /snap/go/10455/src/bytes/bytes.go:712
// upper case, giving priority to the special casing rules.
//line /snap/go/10455/src/bytes/bytes.go:714
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:714
	_go_fuzz_dep_.CoverTab[489]++
						return Map(c.ToUpper, s)
//line /snap/go/10455/src/bytes/bytes.go:715
	// _ = "end of CoverTab[489]"
}

// ToLowerSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
//line /snap/go/10455/src/bytes/bytes.go:718
// lower case, giving priority to the special casing rules.
//line /snap/go/10455/src/bytes/bytes.go:720
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:720
	_go_fuzz_dep_.CoverTab[490]++
						return Map(c.ToLower, s)
//line /snap/go/10455/src/bytes/bytes.go:721
	// _ = "end of CoverTab[490]"
}

// ToTitleSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
//line /snap/go/10455/src/bytes/bytes.go:724
// title case, giving priority to the special casing rules.
//line /snap/go/10455/src/bytes/bytes.go:726
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:726
	_go_fuzz_dep_.CoverTab[491]++
						return Map(c.ToTitle, s)
//line /snap/go/10455/src/bytes/bytes.go:727
	// _ = "end of CoverTab[491]"
}

// ToValidUTF8 treats s as UTF-8-encoded bytes and returns a copy with each run of bytes
//line /snap/go/10455/src/bytes/bytes.go:730
// representing invalid UTF-8 replaced with the bytes in replacement, which may be empty.
//line /snap/go/10455/src/bytes/bytes.go:732
func ToValidUTF8(s, replacement []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:732
	_go_fuzz_dep_.CoverTab[492]++
						b := make([]byte, 0, len(s)+len(replacement))
						invalid := false
//line /snap/go/10455/src/bytes/bytes.go:734
	_go_fuzz_dep_.CoverTab[786464] = 0
						for i := 0; i < len(s); {
//line /snap/go/10455/src/bytes/bytes.go:735
		if _go_fuzz_dep_.CoverTab[786464] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:735
			_go_fuzz_dep_.CoverTab[524819]++
//line /snap/go/10455/src/bytes/bytes.go:735
		} else {
//line /snap/go/10455/src/bytes/bytes.go:735
			_go_fuzz_dep_.CoverTab[524820]++
//line /snap/go/10455/src/bytes/bytes.go:735
		}
//line /snap/go/10455/src/bytes/bytes.go:735
		_go_fuzz_dep_.CoverTab[786464] = 1
//line /snap/go/10455/src/bytes/bytes.go:735
		_go_fuzz_dep_.CoverTab[494]++
							c := s[i]
							if c < utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:737
			_go_fuzz_dep_.CoverTab[524542]++
//line /snap/go/10455/src/bytes/bytes.go:737
			_go_fuzz_dep_.CoverTab[497]++
								i++
								invalid = false
								b = append(b, c)
								continue
//line /snap/go/10455/src/bytes/bytes.go:741
			// _ = "end of CoverTab[497]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:742
			_go_fuzz_dep_.CoverTab[524543]++
//line /snap/go/10455/src/bytes/bytes.go:742
			_go_fuzz_dep_.CoverTab[498]++
//line /snap/go/10455/src/bytes/bytes.go:742
			// _ = "end of CoverTab[498]"
//line /snap/go/10455/src/bytes/bytes.go:742
		}
//line /snap/go/10455/src/bytes/bytes.go:742
		// _ = "end of CoverTab[494]"
//line /snap/go/10455/src/bytes/bytes.go:742
		_go_fuzz_dep_.CoverTab[495]++
							_, wid := utf8.DecodeRune(s[i:])
							if wid == 1 {
//line /snap/go/10455/src/bytes/bytes.go:744
			_go_fuzz_dep_.CoverTab[524544]++
//line /snap/go/10455/src/bytes/bytes.go:744
			_go_fuzz_dep_.CoverTab[499]++
								i++
								if !invalid {
//line /snap/go/10455/src/bytes/bytes.go:746
				_go_fuzz_dep_.CoverTab[524546]++
//line /snap/go/10455/src/bytes/bytes.go:746
				_go_fuzz_dep_.CoverTab[501]++
									invalid = true
									b = append(b, replacement...)
//line /snap/go/10455/src/bytes/bytes.go:748
				// _ = "end of CoverTab[501]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:749
				_go_fuzz_dep_.CoverTab[524547]++
//line /snap/go/10455/src/bytes/bytes.go:749
				_go_fuzz_dep_.CoverTab[502]++
//line /snap/go/10455/src/bytes/bytes.go:749
				// _ = "end of CoverTab[502]"
//line /snap/go/10455/src/bytes/bytes.go:749
			}
//line /snap/go/10455/src/bytes/bytes.go:749
			// _ = "end of CoverTab[499]"
//line /snap/go/10455/src/bytes/bytes.go:749
			_go_fuzz_dep_.CoverTab[500]++
								continue
//line /snap/go/10455/src/bytes/bytes.go:750
			// _ = "end of CoverTab[500]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:751
			_go_fuzz_dep_.CoverTab[524545]++
//line /snap/go/10455/src/bytes/bytes.go:751
			_go_fuzz_dep_.CoverTab[503]++
//line /snap/go/10455/src/bytes/bytes.go:751
			// _ = "end of CoverTab[503]"
//line /snap/go/10455/src/bytes/bytes.go:751
		}
//line /snap/go/10455/src/bytes/bytes.go:751
		// _ = "end of CoverTab[495]"
//line /snap/go/10455/src/bytes/bytes.go:751
		_go_fuzz_dep_.CoverTab[496]++
							invalid = false
							b = append(b, s[i:i+wid]...)
							i += wid
//line /snap/go/10455/src/bytes/bytes.go:754
		// _ = "end of CoverTab[496]"
	}
//line /snap/go/10455/src/bytes/bytes.go:755
	if _go_fuzz_dep_.CoverTab[786464] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:755
		_go_fuzz_dep_.CoverTab[524821]++
//line /snap/go/10455/src/bytes/bytes.go:755
	} else {
//line /snap/go/10455/src/bytes/bytes.go:755
		_go_fuzz_dep_.CoverTab[524822]++
//line /snap/go/10455/src/bytes/bytes.go:755
	}
//line /snap/go/10455/src/bytes/bytes.go:755
	// _ = "end of CoverTab[492]"
//line /snap/go/10455/src/bytes/bytes.go:755
	_go_fuzz_dep_.CoverTab[493]++
						return b
//line /snap/go/10455/src/bytes/bytes.go:756
	// _ = "end of CoverTab[493]"
}

// isSeparator reports whether the rune could mark a word boundary.
//line /snap/go/10455/src/bytes/bytes.go:759
// TODO: update when package unicode captures more of the properties.
//line /snap/go/10455/src/bytes/bytes.go:761
func isSeparator(r rune) bool {
//line /snap/go/10455/src/bytes/bytes.go:761
	_go_fuzz_dep_.CoverTab[504]++

						if r <= 0x7F {
//line /snap/go/10455/src/bytes/bytes.go:763
		_go_fuzz_dep_.CoverTab[524548]++
//line /snap/go/10455/src/bytes/bytes.go:763
		_go_fuzz_dep_.CoverTab[507]++
							switch {
		case '0' <= r && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:765
			_go_fuzz_dep_.CoverTab[514]++
//line /snap/go/10455/src/bytes/bytes.go:765
			return r <= '9'
//line /snap/go/10455/src/bytes/bytes.go:765
			// _ = "end of CoverTab[514]"
//line /snap/go/10455/src/bytes/bytes.go:765
		}():
//line /snap/go/10455/src/bytes/bytes.go:765
			_go_fuzz_dep_.CoverTab[524550]++
//line /snap/go/10455/src/bytes/bytes.go:765
			_go_fuzz_dep_.CoverTab[509]++
								return false
//line /snap/go/10455/src/bytes/bytes.go:766
			// _ = "end of CoverTab[509]"
		case 'a' <= r && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:767
			_go_fuzz_dep_.CoverTab[515]++
//line /snap/go/10455/src/bytes/bytes.go:767
			return r <= 'z'
//line /snap/go/10455/src/bytes/bytes.go:767
			// _ = "end of CoverTab[515]"
//line /snap/go/10455/src/bytes/bytes.go:767
		}():
//line /snap/go/10455/src/bytes/bytes.go:767
			_go_fuzz_dep_.CoverTab[524551]++
//line /snap/go/10455/src/bytes/bytes.go:767
			_go_fuzz_dep_.CoverTab[510]++
								return false
//line /snap/go/10455/src/bytes/bytes.go:768
			// _ = "end of CoverTab[510]"
		case 'A' <= r && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:769
			_go_fuzz_dep_.CoverTab[516]++
//line /snap/go/10455/src/bytes/bytes.go:769
			return r <= 'Z'
//line /snap/go/10455/src/bytes/bytes.go:769
			// _ = "end of CoverTab[516]"
//line /snap/go/10455/src/bytes/bytes.go:769
		}():
//line /snap/go/10455/src/bytes/bytes.go:769
			_go_fuzz_dep_.CoverTab[524552]++
//line /snap/go/10455/src/bytes/bytes.go:769
			_go_fuzz_dep_.CoverTab[511]++
								return false
//line /snap/go/10455/src/bytes/bytes.go:770
			// _ = "end of CoverTab[511]"
		case r == '_':
//line /snap/go/10455/src/bytes/bytes.go:771
			_go_fuzz_dep_.CoverTab[524553]++
//line /snap/go/10455/src/bytes/bytes.go:771
			_go_fuzz_dep_.CoverTab[512]++
								return false
//line /snap/go/10455/src/bytes/bytes.go:772
			// _ = "end of CoverTab[512]"
//line /snap/go/10455/src/bytes/bytes.go:772
		default:
//line /snap/go/10455/src/bytes/bytes.go:772
			_go_fuzz_dep_.CoverTab[524554]++
//line /snap/go/10455/src/bytes/bytes.go:772
			_go_fuzz_dep_.CoverTab[513]++
//line /snap/go/10455/src/bytes/bytes.go:772
			// _ = "end of CoverTab[513]"
		}
//line /snap/go/10455/src/bytes/bytes.go:773
		// _ = "end of CoverTab[507]"
//line /snap/go/10455/src/bytes/bytes.go:773
		_go_fuzz_dep_.CoverTab[508]++
							return true
//line /snap/go/10455/src/bytes/bytes.go:774
		// _ = "end of CoverTab[508]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:775
		_go_fuzz_dep_.CoverTab[524549]++
//line /snap/go/10455/src/bytes/bytes.go:775
		_go_fuzz_dep_.CoverTab[517]++
//line /snap/go/10455/src/bytes/bytes.go:775
		// _ = "end of CoverTab[517]"
//line /snap/go/10455/src/bytes/bytes.go:775
	}
//line /snap/go/10455/src/bytes/bytes.go:775
	// _ = "end of CoverTab[504]"
//line /snap/go/10455/src/bytes/bytes.go:775
	_go_fuzz_dep_.CoverTab[505]++

						if unicode.IsLetter(r) || func() bool {
//line /snap/go/10455/src/bytes/bytes.go:777
		_go_fuzz_dep_.CoverTab[518]++
//line /snap/go/10455/src/bytes/bytes.go:777
		return unicode.IsDigit(r)
//line /snap/go/10455/src/bytes/bytes.go:777
		// _ = "end of CoverTab[518]"
//line /snap/go/10455/src/bytes/bytes.go:777
	}() {
//line /snap/go/10455/src/bytes/bytes.go:777
		_go_fuzz_dep_.CoverTab[524555]++
//line /snap/go/10455/src/bytes/bytes.go:777
		_go_fuzz_dep_.CoverTab[519]++
							return false
//line /snap/go/10455/src/bytes/bytes.go:778
		// _ = "end of CoverTab[519]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:779
		_go_fuzz_dep_.CoverTab[524556]++
//line /snap/go/10455/src/bytes/bytes.go:779
		_go_fuzz_dep_.CoverTab[520]++
//line /snap/go/10455/src/bytes/bytes.go:779
		// _ = "end of CoverTab[520]"
//line /snap/go/10455/src/bytes/bytes.go:779
	}
//line /snap/go/10455/src/bytes/bytes.go:779
	// _ = "end of CoverTab[505]"
//line /snap/go/10455/src/bytes/bytes.go:779
	_go_fuzz_dep_.CoverTab[506]++

						return unicode.IsSpace(r)
//line /snap/go/10455/src/bytes/bytes.go:781
	// _ = "end of CoverTab[506]"
}

// Title treats s as UTF-8-encoded bytes and returns a copy with all Unicode letters that begin
//line /snap/go/10455/src/bytes/bytes.go:784
// words mapped to their title case.
//line /snap/go/10455/src/bytes/bytes.go:784
//
//line /snap/go/10455/src/bytes/bytes.go:784
// Deprecated: The rule Title uses for word boundaries does not handle Unicode
//line /snap/go/10455/src/bytes/bytes.go:784
// punctuation properly. Use golang.org/x/text/cases instead.
//line /snap/go/10455/src/bytes/bytes.go:789
func Title(s []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:789
	_go_fuzz_dep_.CoverTab[521]++

//line /snap/go/10455/src/bytes/bytes.go:793
	prev := ' '
	return Map(
		func(r rune) rune {
//line /snap/go/10455/src/bytes/bytes.go:795
			_go_fuzz_dep_.CoverTab[522]++
								if isSeparator(prev) {
//line /snap/go/10455/src/bytes/bytes.go:796
				_go_fuzz_dep_.CoverTab[524557]++
//line /snap/go/10455/src/bytes/bytes.go:796
				_go_fuzz_dep_.CoverTab[524]++
									prev = r
									return unicode.ToTitle(r)
//line /snap/go/10455/src/bytes/bytes.go:798
				// _ = "end of CoverTab[524]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:799
				_go_fuzz_dep_.CoverTab[524558]++
//line /snap/go/10455/src/bytes/bytes.go:799
				_go_fuzz_dep_.CoverTab[525]++
//line /snap/go/10455/src/bytes/bytes.go:799
				// _ = "end of CoverTab[525]"
//line /snap/go/10455/src/bytes/bytes.go:799
			}
//line /snap/go/10455/src/bytes/bytes.go:799
			// _ = "end of CoverTab[522]"
//line /snap/go/10455/src/bytes/bytes.go:799
			_go_fuzz_dep_.CoverTab[523]++
								prev = r
								return r
//line /snap/go/10455/src/bytes/bytes.go:801
			// _ = "end of CoverTab[523]"
		},
		s)
//line /snap/go/10455/src/bytes/bytes.go:803
	// _ = "end of CoverTab[521]"
}

// TrimLeftFunc treats s as UTF-8-encoded bytes and returns a subslice of s by slicing off
//line /snap/go/10455/src/bytes/bytes.go:806
// all leading UTF-8-encoded code points c that satisfy f(c).
//line /snap/go/10455/src/bytes/bytes.go:808
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte {
//line /snap/go/10455/src/bytes/bytes.go:808
	_go_fuzz_dep_.CoverTab[526]++
						i := indexFunc(s, f, false)
						if i == -1 {
//line /snap/go/10455/src/bytes/bytes.go:810
		_go_fuzz_dep_.CoverTab[524559]++
//line /snap/go/10455/src/bytes/bytes.go:810
		_go_fuzz_dep_.CoverTab[528]++
							return nil
//line /snap/go/10455/src/bytes/bytes.go:811
		// _ = "end of CoverTab[528]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:812
		_go_fuzz_dep_.CoverTab[524560]++
//line /snap/go/10455/src/bytes/bytes.go:812
		_go_fuzz_dep_.CoverTab[529]++
//line /snap/go/10455/src/bytes/bytes.go:812
		// _ = "end of CoverTab[529]"
//line /snap/go/10455/src/bytes/bytes.go:812
	}
//line /snap/go/10455/src/bytes/bytes.go:812
	// _ = "end of CoverTab[526]"
//line /snap/go/10455/src/bytes/bytes.go:812
	_go_fuzz_dep_.CoverTab[527]++
						return s[i:]
//line /snap/go/10455/src/bytes/bytes.go:813
	// _ = "end of CoverTab[527]"
}

// TrimRightFunc returns a subslice of s by slicing off all trailing
//line /snap/go/10455/src/bytes/bytes.go:816
// UTF-8-encoded code points c that satisfy f(c).
//line /snap/go/10455/src/bytes/bytes.go:818
func TrimRightFunc(s []byte, f func(r rune) bool) []byte {
//line /snap/go/10455/src/bytes/bytes.go:818
	_go_fuzz_dep_.CoverTab[530]++
						i := lastIndexFunc(s, f, false)
						if i >= 0 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:820
		_go_fuzz_dep_.CoverTab[532]++
//line /snap/go/10455/src/bytes/bytes.go:820
		return s[i] >= utf8.RuneSelf
//line /snap/go/10455/src/bytes/bytes.go:820
		// _ = "end of CoverTab[532]"
//line /snap/go/10455/src/bytes/bytes.go:820
	}() {
//line /snap/go/10455/src/bytes/bytes.go:820
		_go_fuzz_dep_.CoverTab[524561]++
//line /snap/go/10455/src/bytes/bytes.go:820
		_go_fuzz_dep_.CoverTab[533]++
							_, wid := utf8.DecodeRune(s[i:])
							i += wid
//line /snap/go/10455/src/bytes/bytes.go:822
		// _ = "end of CoverTab[533]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:823
		_go_fuzz_dep_.CoverTab[524562]++
//line /snap/go/10455/src/bytes/bytes.go:823
		_go_fuzz_dep_.CoverTab[534]++
							i++
//line /snap/go/10455/src/bytes/bytes.go:824
		// _ = "end of CoverTab[534]"
	}
//line /snap/go/10455/src/bytes/bytes.go:825
	// _ = "end of CoverTab[530]"
//line /snap/go/10455/src/bytes/bytes.go:825
	_go_fuzz_dep_.CoverTab[531]++
						return s[0:i]
//line /snap/go/10455/src/bytes/bytes.go:826
	// _ = "end of CoverTab[531]"
}

// TrimFunc returns a subslice of s by slicing off all leading and trailing
//line /snap/go/10455/src/bytes/bytes.go:829
// UTF-8-encoded code points c that satisfy f(c).
//line /snap/go/10455/src/bytes/bytes.go:831
func TrimFunc(s []byte, f func(r rune) bool) []byte {
//line /snap/go/10455/src/bytes/bytes.go:831
	_go_fuzz_dep_.CoverTab[535]++
						return TrimRightFunc(TrimLeftFunc(s, f), f)
//line /snap/go/10455/src/bytes/bytes.go:832
	// _ = "end of CoverTab[535]"
}

// TrimPrefix returns s without the provided leading prefix string.
//line /snap/go/10455/src/bytes/bytes.go:835
// If s doesn't start with prefix, s is returned unchanged.
//line /snap/go/10455/src/bytes/bytes.go:837
func TrimPrefix(s, prefix []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:837
	_go_fuzz_dep_.CoverTab[536]++
						if HasPrefix(s, prefix) {
//line /snap/go/10455/src/bytes/bytes.go:838
		_go_fuzz_dep_.CoverTab[524563]++
//line /snap/go/10455/src/bytes/bytes.go:838
		_go_fuzz_dep_.CoverTab[538]++
							return s[len(prefix):]
//line /snap/go/10455/src/bytes/bytes.go:839
		// _ = "end of CoverTab[538]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:840
		_go_fuzz_dep_.CoverTab[524564]++
//line /snap/go/10455/src/bytes/bytes.go:840
		_go_fuzz_dep_.CoverTab[539]++
//line /snap/go/10455/src/bytes/bytes.go:840
		// _ = "end of CoverTab[539]"
//line /snap/go/10455/src/bytes/bytes.go:840
	}
//line /snap/go/10455/src/bytes/bytes.go:840
	// _ = "end of CoverTab[536]"
//line /snap/go/10455/src/bytes/bytes.go:840
	_go_fuzz_dep_.CoverTab[537]++
						return s
//line /snap/go/10455/src/bytes/bytes.go:841
	// _ = "end of CoverTab[537]"
}

// TrimSuffix returns s without the provided trailing suffix string.
//line /snap/go/10455/src/bytes/bytes.go:844
// If s doesn't end with suffix, s is returned unchanged.
//line /snap/go/10455/src/bytes/bytes.go:846
func TrimSuffix(s, suffix []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:846
	_go_fuzz_dep_.CoverTab[540]++
						if HasSuffix(s, suffix) {
//line /snap/go/10455/src/bytes/bytes.go:847
		_go_fuzz_dep_.CoverTab[524565]++
//line /snap/go/10455/src/bytes/bytes.go:847
		_go_fuzz_dep_.CoverTab[542]++
							return s[:len(s)-len(suffix)]
//line /snap/go/10455/src/bytes/bytes.go:848
		// _ = "end of CoverTab[542]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:849
		_go_fuzz_dep_.CoverTab[524566]++
//line /snap/go/10455/src/bytes/bytes.go:849
		_go_fuzz_dep_.CoverTab[543]++
//line /snap/go/10455/src/bytes/bytes.go:849
		// _ = "end of CoverTab[543]"
//line /snap/go/10455/src/bytes/bytes.go:849
	}
//line /snap/go/10455/src/bytes/bytes.go:849
	// _ = "end of CoverTab[540]"
//line /snap/go/10455/src/bytes/bytes.go:849
	_go_fuzz_dep_.CoverTab[541]++
						return s
//line /snap/go/10455/src/bytes/bytes.go:850
	// _ = "end of CoverTab[541]"
}

// IndexFunc interprets s as a sequence of UTF-8-encoded code points.
//line /snap/go/10455/src/bytes/bytes.go:853
// It returns the byte index in s of the first Unicode
//line /snap/go/10455/src/bytes/bytes.go:853
// code point satisfying f(c), or -1 if none do.
//line /snap/go/10455/src/bytes/bytes.go:856
func IndexFunc(s []byte, f func(r rune) bool) int {
//line /snap/go/10455/src/bytes/bytes.go:856
	_go_fuzz_dep_.CoverTab[544]++
						return indexFunc(s, f, true)
//line /snap/go/10455/src/bytes/bytes.go:857
	// _ = "end of CoverTab[544]"
}

// LastIndexFunc interprets s as a sequence of UTF-8-encoded code points.
//line /snap/go/10455/src/bytes/bytes.go:860
// It returns the byte index in s of the last Unicode
//line /snap/go/10455/src/bytes/bytes.go:860
// code point satisfying f(c), or -1 if none do.
//line /snap/go/10455/src/bytes/bytes.go:863
func LastIndexFunc(s []byte, f func(r rune) bool) int {
//line /snap/go/10455/src/bytes/bytes.go:863
	_go_fuzz_dep_.CoverTab[545]++
						return lastIndexFunc(s, f, true)
//line /snap/go/10455/src/bytes/bytes.go:864
	// _ = "end of CoverTab[545]"
}

// indexFunc is the same as IndexFunc except that if
//line /snap/go/10455/src/bytes/bytes.go:867
// truth==false, the sense of the predicate function is
//line /snap/go/10455/src/bytes/bytes.go:867
// inverted.
//line /snap/go/10455/src/bytes/bytes.go:870
func indexFunc(s []byte, f func(r rune) bool, truth bool) int {
//line /snap/go/10455/src/bytes/bytes.go:870
	_go_fuzz_dep_.CoverTab[546]++
						start := 0
//line /snap/go/10455/src/bytes/bytes.go:871
	_go_fuzz_dep_.CoverTab[786465] = 0
						for start < len(s) {
//line /snap/go/10455/src/bytes/bytes.go:872
		if _go_fuzz_dep_.CoverTab[786465] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:872
			_go_fuzz_dep_.CoverTab[524823]++
//line /snap/go/10455/src/bytes/bytes.go:872
		} else {
//line /snap/go/10455/src/bytes/bytes.go:872
			_go_fuzz_dep_.CoverTab[524824]++
//line /snap/go/10455/src/bytes/bytes.go:872
		}
//line /snap/go/10455/src/bytes/bytes.go:872
		_go_fuzz_dep_.CoverTab[786465] = 1
//line /snap/go/10455/src/bytes/bytes.go:872
		_go_fuzz_dep_.CoverTab[548]++
							wid := 1
							r := rune(s[start])
							if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:875
			_go_fuzz_dep_.CoverTab[524567]++
//line /snap/go/10455/src/bytes/bytes.go:875
			_go_fuzz_dep_.CoverTab[551]++
								r, wid = utf8.DecodeRune(s[start:])
//line /snap/go/10455/src/bytes/bytes.go:876
			// _ = "end of CoverTab[551]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:877
			_go_fuzz_dep_.CoverTab[524568]++
//line /snap/go/10455/src/bytes/bytes.go:877
			_go_fuzz_dep_.CoverTab[552]++
//line /snap/go/10455/src/bytes/bytes.go:877
			// _ = "end of CoverTab[552]"
//line /snap/go/10455/src/bytes/bytes.go:877
		}
//line /snap/go/10455/src/bytes/bytes.go:877
		// _ = "end of CoverTab[548]"
//line /snap/go/10455/src/bytes/bytes.go:877
		_go_fuzz_dep_.CoverTab[549]++
							if f(r) == truth {
//line /snap/go/10455/src/bytes/bytes.go:878
			_go_fuzz_dep_.CoverTab[524569]++
//line /snap/go/10455/src/bytes/bytes.go:878
			_go_fuzz_dep_.CoverTab[553]++
								return start
//line /snap/go/10455/src/bytes/bytes.go:879
			// _ = "end of CoverTab[553]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:880
			_go_fuzz_dep_.CoverTab[524570]++
//line /snap/go/10455/src/bytes/bytes.go:880
			_go_fuzz_dep_.CoverTab[554]++
//line /snap/go/10455/src/bytes/bytes.go:880
			// _ = "end of CoverTab[554]"
//line /snap/go/10455/src/bytes/bytes.go:880
		}
//line /snap/go/10455/src/bytes/bytes.go:880
		// _ = "end of CoverTab[549]"
//line /snap/go/10455/src/bytes/bytes.go:880
		_go_fuzz_dep_.CoverTab[550]++
							start += wid
//line /snap/go/10455/src/bytes/bytes.go:881
		// _ = "end of CoverTab[550]"
	}
//line /snap/go/10455/src/bytes/bytes.go:882
	if _go_fuzz_dep_.CoverTab[786465] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:882
		_go_fuzz_dep_.CoverTab[524825]++
//line /snap/go/10455/src/bytes/bytes.go:882
	} else {
//line /snap/go/10455/src/bytes/bytes.go:882
		_go_fuzz_dep_.CoverTab[524826]++
//line /snap/go/10455/src/bytes/bytes.go:882
	}
//line /snap/go/10455/src/bytes/bytes.go:882
	// _ = "end of CoverTab[546]"
//line /snap/go/10455/src/bytes/bytes.go:882
	_go_fuzz_dep_.CoverTab[547]++
						return -1
//line /snap/go/10455/src/bytes/bytes.go:883
	// _ = "end of CoverTab[547]"
}

// lastIndexFunc is the same as LastIndexFunc except that if
//line /snap/go/10455/src/bytes/bytes.go:886
// truth==false, the sense of the predicate function is
//line /snap/go/10455/src/bytes/bytes.go:886
// inverted.
//line /snap/go/10455/src/bytes/bytes.go:889
func lastIndexFunc(s []byte, f func(r rune) bool, truth bool) int {
//line /snap/go/10455/src/bytes/bytes.go:889
	_go_fuzz_dep_.CoverTab[555]++
//line /snap/go/10455/src/bytes/bytes.go:889
	_go_fuzz_dep_.CoverTab[786466] = 0
						for i := len(s); i > 0; {
//line /snap/go/10455/src/bytes/bytes.go:890
		if _go_fuzz_dep_.CoverTab[786466] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:890
			_go_fuzz_dep_.CoverTab[524827]++
//line /snap/go/10455/src/bytes/bytes.go:890
		} else {
//line /snap/go/10455/src/bytes/bytes.go:890
			_go_fuzz_dep_.CoverTab[524828]++
//line /snap/go/10455/src/bytes/bytes.go:890
		}
//line /snap/go/10455/src/bytes/bytes.go:890
		_go_fuzz_dep_.CoverTab[786466] = 1
//line /snap/go/10455/src/bytes/bytes.go:890
		_go_fuzz_dep_.CoverTab[557]++
							r, size := rune(s[i-1]), 1
							if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:892
			_go_fuzz_dep_.CoverTab[524571]++
//line /snap/go/10455/src/bytes/bytes.go:892
			_go_fuzz_dep_.CoverTab[559]++
								r, size = utf8.DecodeLastRune(s[0:i])
//line /snap/go/10455/src/bytes/bytes.go:893
			// _ = "end of CoverTab[559]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:894
			_go_fuzz_dep_.CoverTab[524572]++
//line /snap/go/10455/src/bytes/bytes.go:894
			_go_fuzz_dep_.CoverTab[560]++
//line /snap/go/10455/src/bytes/bytes.go:894
			// _ = "end of CoverTab[560]"
//line /snap/go/10455/src/bytes/bytes.go:894
		}
//line /snap/go/10455/src/bytes/bytes.go:894
		// _ = "end of CoverTab[557]"
//line /snap/go/10455/src/bytes/bytes.go:894
		_go_fuzz_dep_.CoverTab[558]++
							i -= size
							if f(r) == truth {
//line /snap/go/10455/src/bytes/bytes.go:896
			_go_fuzz_dep_.CoverTab[524573]++
//line /snap/go/10455/src/bytes/bytes.go:896
			_go_fuzz_dep_.CoverTab[561]++
								return i
//line /snap/go/10455/src/bytes/bytes.go:897
			// _ = "end of CoverTab[561]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:898
			_go_fuzz_dep_.CoverTab[524574]++
//line /snap/go/10455/src/bytes/bytes.go:898
			_go_fuzz_dep_.CoverTab[562]++
//line /snap/go/10455/src/bytes/bytes.go:898
			// _ = "end of CoverTab[562]"
//line /snap/go/10455/src/bytes/bytes.go:898
		}
//line /snap/go/10455/src/bytes/bytes.go:898
		// _ = "end of CoverTab[558]"
	}
//line /snap/go/10455/src/bytes/bytes.go:899
	if _go_fuzz_dep_.CoverTab[786466] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:899
		_go_fuzz_dep_.CoverTab[524829]++
//line /snap/go/10455/src/bytes/bytes.go:899
	} else {
//line /snap/go/10455/src/bytes/bytes.go:899
		_go_fuzz_dep_.CoverTab[524830]++
//line /snap/go/10455/src/bytes/bytes.go:899
	}
//line /snap/go/10455/src/bytes/bytes.go:899
	// _ = "end of CoverTab[555]"
//line /snap/go/10455/src/bytes/bytes.go:899
	_go_fuzz_dep_.CoverTab[556]++
						return -1
//line /snap/go/10455/src/bytes/bytes.go:900
	// _ = "end of CoverTab[556]"
}

// asciiSet is a 32-byte value, where each bit represents the presence of a
//line /snap/go/10455/src/bytes/bytes.go:903
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
//line /snap/go/10455/src/bytes/bytes.go:903
// starting with the least-significant bit of the lowest word to the
//line /snap/go/10455/src/bytes/bytes.go:903
// most-significant bit of the highest word, map to the full range of all
//line /snap/go/10455/src/bytes/bytes.go:903
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
//line /snap/go/10455/src/bytes/bytes.go:903
// ensuring that any non-ASCII character will be reported as not in the set.
//line /snap/go/10455/src/bytes/bytes.go:903
// This allocates a total of 32 bytes even though the upper half
//line /snap/go/10455/src/bytes/bytes.go:903
// is unused to avoid bounds checks in asciiSet.contains.
//line /snap/go/10455/src/bytes/bytes.go:911
type asciiSet [8]uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
//line /snap/go/10455/src/bytes/bytes.go:913
// characters in chars are ASCII.
//line /snap/go/10455/src/bytes/bytes.go:915
func makeASCIISet(chars string) (as asciiSet, ok bool) {
//line /snap/go/10455/src/bytes/bytes.go:915
	_go_fuzz_dep_.CoverTab[563]++
//line /snap/go/10455/src/bytes/bytes.go:915
	_go_fuzz_dep_.CoverTab[786467] = 0
						for i := 0; i < len(chars); i++ {
//line /snap/go/10455/src/bytes/bytes.go:916
		if _go_fuzz_dep_.CoverTab[786467] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:916
			_go_fuzz_dep_.CoverTab[524831]++
//line /snap/go/10455/src/bytes/bytes.go:916
		} else {
//line /snap/go/10455/src/bytes/bytes.go:916
			_go_fuzz_dep_.CoverTab[524832]++
//line /snap/go/10455/src/bytes/bytes.go:916
		}
//line /snap/go/10455/src/bytes/bytes.go:916
		_go_fuzz_dep_.CoverTab[786467] = 1
//line /snap/go/10455/src/bytes/bytes.go:916
		_go_fuzz_dep_.CoverTab[565]++
							c := chars[i]
							if c >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:918
			_go_fuzz_dep_.CoverTab[524575]++
//line /snap/go/10455/src/bytes/bytes.go:918
			_go_fuzz_dep_.CoverTab[567]++
								return as, false
//line /snap/go/10455/src/bytes/bytes.go:919
			// _ = "end of CoverTab[567]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:920
			_go_fuzz_dep_.CoverTab[524576]++
//line /snap/go/10455/src/bytes/bytes.go:920
			_go_fuzz_dep_.CoverTab[568]++
//line /snap/go/10455/src/bytes/bytes.go:920
			// _ = "end of CoverTab[568]"
//line /snap/go/10455/src/bytes/bytes.go:920
		}
//line /snap/go/10455/src/bytes/bytes.go:920
		// _ = "end of CoverTab[565]"
//line /snap/go/10455/src/bytes/bytes.go:920
		_go_fuzz_dep_.CoverTab[566]++
							as[c/32] |= 1 << (c % 32)
//line /snap/go/10455/src/bytes/bytes.go:921
		// _ = "end of CoverTab[566]"
	}
//line /snap/go/10455/src/bytes/bytes.go:922
	if _go_fuzz_dep_.CoverTab[786467] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:922
		_go_fuzz_dep_.CoverTab[524833]++
//line /snap/go/10455/src/bytes/bytes.go:922
	} else {
//line /snap/go/10455/src/bytes/bytes.go:922
		_go_fuzz_dep_.CoverTab[524834]++
//line /snap/go/10455/src/bytes/bytes.go:922
	}
//line /snap/go/10455/src/bytes/bytes.go:922
	// _ = "end of CoverTab[563]"
//line /snap/go/10455/src/bytes/bytes.go:922
	_go_fuzz_dep_.CoverTab[564]++
						return as, true
//line /snap/go/10455/src/bytes/bytes.go:923
	// _ = "end of CoverTab[564]"
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c byte) bool {
//line /snap/go/10455/src/bytes/bytes.go:927
	_go_fuzz_dep_.CoverTab[569]++
						return (as[c/32] & (1 << (c % 32))) != 0
//line /snap/go/10455/src/bytes/bytes.go:928
	// _ = "end of CoverTab[569]"
}

// containsRune is a simplified version of strings.ContainsRune
//line /snap/go/10455/src/bytes/bytes.go:931
// to avoid importing the strings package.
//line /snap/go/10455/src/bytes/bytes.go:931
// We avoid bytes.ContainsRune to avoid allocating a temporary copy of s.
//line /snap/go/10455/src/bytes/bytes.go:934
func containsRune(s string, r rune) bool {
//line /snap/go/10455/src/bytes/bytes.go:934
	_go_fuzz_dep_.CoverTab[570]++
//line /snap/go/10455/src/bytes/bytes.go:934
	_go_fuzz_dep_.CoverTab[786468] = 0
						for _, c := range s {
//line /snap/go/10455/src/bytes/bytes.go:935
		if _go_fuzz_dep_.CoverTab[786468] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:935
			_go_fuzz_dep_.CoverTab[524835]++
//line /snap/go/10455/src/bytes/bytes.go:935
		} else {
//line /snap/go/10455/src/bytes/bytes.go:935
			_go_fuzz_dep_.CoverTab[524836]++
//line /snap/go/10455/src/bytes/bytes.go:935
		}
//line /snap/go/10455/src/bytes/bytes.go:935
		_go_fuzz_dep_.CoverTab[786468] = 1
//line /snap/go/10455/src/bytes/bytes.go:935
		_go_fuzz_dep_.CoverTab[572]++
							if c == r {
//line /snap/go/10455/src/bytes/bytes.go:936
			_go_fuzz_dep_.CoverTab[524577]++
//line /snap/go/10455/src/bytes/bytes.go:936
			_go_fuzz_dep_.CoverTab[573]++
								return true
//line /snap/go/10455/src/bytes/bytes.go:937
			// _ = "end of CoverTab[573]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:938
			_go_fuzz_dep_.CoverTab[524578]++
//line /snap/go/10455/src/bytes/bytes.go:938
			_go_fuzz_dep_.CoverTab[574]++
//line /snap/go/10455/src/bytes/bytes.go:938
			// _ = "end of CoverTab[574]"
//line /snap/go/10455/src/bytes/bytes.go:938
		}
//line /snap/go/10455/src/bytes/bytes.go:938
		// _ = "end of CoverTab[572]"
	}
//line /snap/go/10455/src/bytes/bytes.go:939
	if _go_fuzz_dep_.CoverTab[786468] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:939
		_go_fuzz_dep_.CoverTab[524837]++
//line /snap/go/10455/src/bytes/bytes.go:939
	} else {
//line /snap/go/10455/src/bytes/bytes.go:939
		_go_fuzz_dep_.CoverTab[524838]++
//line /snap/go/10455/src/bytes/bytes.go:939
	}
//line /snap/go/10455/src/bytes/bytes.go:939
	// _ = "end of CoverTab[570]"
//line /snap/go/10455/src/bytes/bytes.go:939
	_go_fuzz_dep_.CoverTab[571]++
						return false
//line /snap/go/10455/src/bytes/bytes.go:940
	// _ = "end of CoverTab[571]"
}

// Trim returns a subslice of s by slicing off all leading and
//line /snap/go/10455/src/bytes/bytes.go:943
// trailing UTF-8-encoded code points contained in cutset.
//line /snap/go/10455/src/bytes/bytes.go:945
func Trim(s []byte, cutset string) []byte {
//line /snap/go/10455/src/bytes/bytes.go:945
	_go_fuzz_dep_.CoverTab[575]++
						if len(s) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:946
		_go_fuzz_dep_.CoverTab[524579]++
//line /snap/go/10455/src/bytes/bytes.go:946
		_go_fuzz_dep_.CoverTab[580]++

							return nil
//line /snap/go/10455/src/bytes/bytes.go:948
		// _ = "end of CoverTab[580]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:949
		_go_fuzz_dep_.CoverTab[524580]++
//line /snap/go/10455/src/bytes/bytes.go:949
		_go_fuzz_dep_.CoverTab[581]++
//line /snap/go/10455/src/bytes/bytes.go:949
		// _ = "end of CoverTab[581]"
//line /snap/go/10455/src/bytes/bytes.go:949
	}
//line /snap/go/10455/src/bytes/bytes.go:949
	// _ = "end of CoverTab[575]"
//line /snap/go/10455/src/bytes/bytes.go:949
	_go_fuzz_dep_.CoverTab[576]++
						if cutset == "" {
//line /snap/go/10455/src/bytes/bytes.go:950
		_go_fuzz_dep_.CoverTab[524581]++
//line /snap/go/10455/src/bytes/bytes.go:950
		_go_fuzz_dep_.CoverTab[582]++
							return s
//line /snap/go/10455/src/bytes/bytes.go:951
		// _ = "end of CoverTab[582]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:952
		_go_fuzz_dep_.CoverTab[524582]++
//line /snap/go/10455/src/bytes/bytes.go:952
		_go_fuzz_dep_.CoverTab[583]++
//line /snap/go/10455/src/bytes/bytes.go:952
		// _ = "end of CoverTab[583]"
//line /snap/go/10455/src/bytes/bytes.go:952
	}
//line /snap/go/10455/src/bytes/bytes.go:952
	// _ = "end of CoverTab[576]"
//line /snap/go/10455/src/bytes/bytes.go:952
	_go_fuzz_dep_.CoverTab[577]++
						if len(cutset) == 1 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:953
		_go_fuzz_dep_.CoverTab[584]++
//line /snap/go/10455/src/bytes/bytes.go:953
		return cutset[0] < utf8.RuneSelf
//line /snap/go/10455/src/bytes/bytes.go:953
		// _ = "end of CoverTab[584]"
//line /snap/go/10455/src/bytes/bytes.go:953
	}() {
//line /snap/go/10455/src/bytes/bytes.go:953
		_go_fuzz_dep_.CoverTab[524583]++
//line /snap/go/10455/src/bytes/bytes.go:953
		_go_fuzz_dep_.CoverTab[585]++
							return trimLeftByte(trimRightByte(s, cutset[0]), cutset[0])
//line /snap/go/10455/src/bytes/bytes.go:954
		// _ = "end of CoverTab[585]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:955
		_go_fuzz_dep_.CoverTab[524584]++
//line /snap/go/10455/src/bytes/bytes.go:955
		_go_fuzz_dep_.CoverTab[586]++
//line /snap/go/10455/src/bytes/bytes.go:955
		// _ = "end of CoverTab[586]"
//line /snap/go/10455/src/bytes/bytes.go:955
	}
//line /snap/go/10455/src/bytes/bytes.go:955
	// _ = "end of CoverTab[577]"
//line /snap/go/10455/src/bytes/bytes.go:955
	_go_fuzz_dep_.CoverTab[578]++
						if as, ok := makeASCIISet(cutset); ok {
//line /snap/go/10455/src/bytes/bytes.go:956
		_go_fuzz_dep_.CoverTab[524585]++
//line /snap/go/10455/src/bytes/bytes.go:956
		_go_fuzz_dep_.CoverTab[587]++
							return trimLeftASCII(trimRightASCII(s, &as), &as)
//line /snap/go/10455/src/bytes/bytes.go:957
		// _ = "end of CoverTab[587]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:958
		_go_fuzz_dep_.CoverTab[524586]++
//line /snap/go/10455/src/bytes/bytes.go:958
		_go_fuzz_dep_.CoverTab[588]++
//line /snap/go/10455/src/bytes/bytes.go:958
		// _ = "end of CoverTab[588]"
//line /snap/go/10455/src/bytes/bytes.go:958
	}
//line /snap/go/10455/src/bytes/bytes.go:958
	// _ = "end of CoverTab[578]"
//line /snap/go/10455/src/bytes/bytes.go:958
	_go_fuzz_dep_.CoverTab[579]++
						return trimLeftUnicode(trimRightUnicode(s, cutset), cutset)
//line /snap/go/10455/src/bytes/bytes.go:959
	// _ = "end of CoverTab[579]"
}

// TrimLeft returns a subslice of s by slicing off all leading
//line /snap/go/10455/src/bytes/bytes.go:962
// UTF-8-encoded code points contained in cutset.
//line /snap/go/10455/src/bytes/bytes.go:964
func TrimLeft(s []byte, cutset string) []byte {
//line /snap/go/10455/src/bytes/bytes.go:964
	_go_fuzz_dep_.CoverTab[589]++
						if len(s) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:965
		_go_fuzz_dep_.CoverTab[524587]++
//line /snap/go/10455/src/bytes/bytes.go:965
		_go_fuzz_dep_.CoverTab[594]++

							return nil
//line /snap/go/10455/src/bytes/bytes.go:967
		// _ = "end of CoverTab[594]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:968
		_go_fuzz_dep_.CoverTab[524588]++
//line /snap/go/10455/src/bytes/bytes.go:968
		_go_fuzz_dep_.CoverTab[595]++
//line /snap/go/10455/src/bytes/bytes.go:968
		// _ = "end of CoverTab[595]"
//line /snap/go/10455/src/bytes/bytes.go:968
	}
//line /snap/go/10455/src/bytes/bytes.go:968
	// _ = "end of CoverTab[589]"
//line /snap/go/10455/src/bytes/bytes.go:968
	_go_fuzz_dep_.CoverTab[590]++
						if cutset == "" {
//line /snap/go/10455/src/bytes/bytes.go:969
		_go_fuzz_dep_.CoverTab[524589]++
//line /snap/go/10455/src/bytes/bytes.go:969
		_go_fuzz_dep_.CoverTab[596]++
							return s
//line /snap/go/10455/src/bytes/bytes.go:970
		// _ = "end of CoverTab[596]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:971
		_go_fuzz_dep_.CoverTab[524590]++
//line /snap/go/10455/src/bytes/bytes.go:971
		_go_fuzz_dep_.CoverTab[597]++
//line /snap/go/10455/src/bytes/bytes.go:971
		// _ = "end of CoverTab[597]"
//line /snap/go/10455/src/bytes/bytes.go:971
	}
//line /snap/go/10455/src/bytes/bytes.go:971
	// _ = "end of CoverTab[590]"
//line /snap/go/10455/src/bytes/bytes.go:971
	_go_fuzz_dep_.CoverTab[591]++
						if len(cutset) == 1 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:972
		_go_fuzz_dep_.CoverTab[598]++
//line /snap/go/10455/src/bytes/bytes.go:972
		return cutset[0] < utf8.RuneSelf
//line /snap/go/10455/src/bytes/bytes.go:972
		// _ = "end of CoverTab[598]"
//line /snap/go/10455/src/bytes/bytes.go:972
	}() {
//line /snap/go/10455/src/bytes/bytes.go:972
		_go_fuzz_dep_.CoverTab[524591]++
//line /snap/go/10455/src/bytes/bytes.go:972
		_go_fuzz_dep_.CoverTab[599]++
							return trimLeftByte(s, cutset[0])
//line /snap/go/10455/src/bytes/bytes.go:973
		// _ = "end of CoverTab[599]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:974
		_go_fuzz_dep_.CoverTab[524592]++
//line /snap/go/10455/src/bytes/bytes.go:974
		_go_fuzz_dep_.CoverTab[600]++
//line /snap/go/10455/src/bytes/bytes.go:974
		// _ = "end of CoverTab[600]"
//line /snap/go/10455/src/bytes/bytes.go:974
	}
//line /snap/go/10455/src/bytes/bytes.go:974
	// _ = "end of CoverTab[591]"
//line /snap/go/10455/src/bytes/bytes.go:974
	_go_fuzz_dep_.CoverTab[592]++
						if as, ok := makeASCIISet(cutset); ok {
//line /snap/go/10455/src/bytes/bytes.go:975
		_go_fuzz_dep_.CoverTab[524593]++
//line /snap/go/10455/src/bytes/bytes.go:975
		_go_fuzz_dep_.CoverTab[601]++
							return trimLeftASCII(s, &as)
//line /snap/go/10455/src/bytes/bytes.go:976
		// _ = "end of CoverTab[601]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:977
		_go_fuzz_dep_.CoverTab[524594]++
//line /snap/go/10455/src/bytes/bytes.go:977
		_go_fuzz_dep_.CoverTab[602]++
//line /snap/go/10455/src/bytes/bytes.go:977
		// _ = "end of CoverTab[602]"
//line /snap/go/10455/src/bytes/bytes.go:977
	}
//line /snap/go/10455/src/bytes/bytes.go:977
	// _ = "end of CoverTab[592]"
//line /snap/go/10455/src/bytes/bytes.go:977
	_go_fuzz_dep_.CoverTab[593]++
						return trimLeftUnicode(s, cutset)
//line /snap/go/10455/src/bytes/bytes.go:978
	// _ = "end of CoverTab[593]"
}

func trimLeftByte(s []byte, c byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:981
	_go_fuzz_dep_.CoverTab[603]++
//line /snap/go/10455/src/bytes/bytes.go:981
	_go_fuzz_dep_.CoverTab[786469] = 0
						for len(s) > 0 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:982
		_go_fuzz_dep_.CoverTab[606]++
//line /snap/go/10455/src/bytes/bytes.go:982
		return s[0] == c
//line /snap/go/10455/src/bytes/bytes.go:982
		// _ = "end of CoverTab[606]"
//line /snap/go/10455/src/bytes/bytes.go:982
	}() {
//line /snap/go/10455/src/bytes/bytes.go:982
		if _go_fuzz_dep_.CoverTab[786469] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:982
			_go_fuzz_dep_.CoverTab[524839]++
//line /snap/go/10455/src/bytes/bytes.go:982
		} else {
//line /snap/go/10455/src/bytes/bytes.go:982
			_go_fuzz_dep_.CoverTab[524840]++
//line /snap/go/10455/src/bytes/bytes.go:982
		}
//line /snap/go/10455/src/bytes/bytes.go:982
		_go_fuzz_dep_.CoverTab[786469] = 1
//line /snap/go/10455/src/bytes/bytes.go:982
		_go_fuzz_dep_.CoverTab[607]++
							s = s[1:]
//line /snap/go/10455/src/bytes/bytes.go:983
		// _ = "end of CoverTab[607]"
	}
//line /snap/go/10455/src/bytes/bytes.go:984
	if _go_fuzz_dep_.CoverTab[786469] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:984
		_go_fuzz_dep_.CoverTab[524841]++
//line /snap/go/10455/src/bytes/bytes.go:984
	} else {
//line /snap/go/10455/src/bytes/bytes.go:984
		_go_fuzz_dep_.CoverTab[524842]++
//line /snap/go/10455/src/bytes/bytes.go:984
	}
//line /snap/go/10455/src/bytes/bytes.go:984
	// _ = "end of CoverTab[603]"
//line /snap/go/10455/src/bytes/bytes.go:984
	_go_fuzz_dep_.CoverTab[604]++
						if len(s) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:985
		_go_fuzz_dep_.CoverTab[524595]++
//line /snap/go/10455/src/bytes/bytes.go:985
		_go_fuzz_dep_.CoverTab[608]++

							return nil
//line /snap/go/10455/src/bytes/bytes.go:987
		// _ = "end of CoverTab[608]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:988
		_go_fuzz_dep_.CoverTab[524596]++
//line /snap/go/10455/src/bytes/bytes.go:988
		_go_fuzz_dep_.CoverTab[609]++
//line /snap/go/10455/src/bytes/bytes.go:988
		// _ = "end of CoverTab[609]"
//line /snap/go/10455/src/bytes/bytes.go:988
	}
//line /snap/go/10455/src/bytes/bytes.go:988
	// _ = "end of CoverTab[604]"
//line /snap/go/10455/src/bytes/bytes.go:988
	_go_fuzz_dep_.CoverTab[605]++
						return s
//line /snap/go/10455/src/bytes/bytes.go:989
	// _ = "end of CoverTab[605]"
}

func trimLeftASCII(s []byte, as *asciiSet) []byte {
//line /snap/go/10455/src/bytes/bytes.go:992
	_go_fuzz_dep_.CoverTab[610]++
//line /snap/go/10455/src/bytes/bytes.go:992
	_go_fuzz_dep_.CoverTab[786470] = 0
						for len(s) > 0 {
//line /snap/go/10455/src/bytes/bytes.go:993
		if _go_fuzz_dep_.CoverTab[786470] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:993
			_go_fuzz_dep_.CoverTab[524843]++
//line /snap/go/10455/src/bytes/bytes.go:993
		} else {
//line /snap/go/10455/src/bytes/bytes.go:993
			_go_fuzz_dep_.CoverTab[524844]++
//line /snap/go/10455/src/bytes/bytes.go:993
		}
//line /snap/go/10455/src/bytes/bytes.go:993
		_go_fuzz_dep_.CoverTab[786470] = 1
//line /snap/go/10455/src/bytes/bytes.go:993
		_go_fuzz_dep_.CoverTab[613]++
							if !as.contains(s[0]) {
//line /snap/go/10455/src/bytes/bytes.go:994
			_go_fuzz_dep_.CoverTab[524597]++
//line /snap/go/10455/src/bytes/bytes.go:994
			_go_fuzz_dep_.CoverTab[615]++
								break
//line /snap/go/10455/src/bytes/bytes.go:995
			// _ = "end of CoverTab[615]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:996
			_go_fuzz_dep_.CoverTab[524598]++
//line /snap/go/10455/src/bytes/bytes.go:996
			_go_fuzz_dep_.CoverTab[616]++
//line /snap/go/10455/src/bytes/bytes.go:996
			// _ = "end of CoverTab[616]"
//line /snap/go/10455/src/bytes/bytes.go:996
		}
//line /snap/go/10455/src/bytes/bytes.go:996
		// _ = "end of CoverTab[613]"
//line /snap/go/10455/src/bytes/bytes.go:996
		_go_fuzz_dep_.CoverTab[614]++
							s = s[1:]
//line /snap/go/10455/src/bytes/bytes.go:997
		// _ = "end of CoverTab[614]"
	}
//line /snap/go/10455/src/bytes/bytes.go:998
	if _go_fuzz_dep_.CoverTab[786470] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:998
		_go_fuzz_dep_.CoverTab[524845]++
//line /snap/go/10455/src/bytes/bytes.go:998
	} else {
//line /snap/go/10455/src/bytes/bytes.go:998
		_go_fuzz_dep_.CoverTab[524846]++
//line /snap/go/10455/src/bytes/bytes.go:998
	}
//line /snap/go/10455/src/bytes/bytes.go:998
	// _ = "end of CoverTab[610]"
//line /snap/go/10455/src/bytes/bytes.go:998
	_go_fuzz_dep_.CoverTab[611]++
						if len(s) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:999
		_go_fuzz_dep_.CoverTab[524599]++
//line /snap/go/10455/src/bytes/bytes.go:999
		_go_fuzz_dep_.CoverTab[617]++

							return nil
//line /snap/go/10455/src/bytes/bytes.go:1001
		// _ = "end of CoverTab[617]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1002
		_go_fuzz_dep_.CoverTab[524600]++
//line /snap/go/10455/src/bytes/bytes.go:1002
		_go_fuzz_dep_.CoverTab[618]++
//line /snap/go/10455/src/bytes/bytes.go:1002
		// _ = "end of CoverTab[618]"
//line /snap/go/10455/src/bytes/bytes.go:1002
	}
//line /snap/go/10455/src/bytes/bytes.go:1002
	// _ = "end of CoverTab[611]"
//line /snap/go/10455/src/bytes/bytes.go:1002
	_go_fuzz_dep_.CoverTab[612]++
						return s
//line /snap/go/10455/src/bytes/bytes.go:1003
	// _ = "end of CoverTab[612]"
}

func trimLeftUnicode(s []byte, cutset string) []byte {
//line /snap/go/10455/src/bytes/bytes.go:1006
	_go_fuzz_dep_.CoverTab[619]++
//line /snap/go/10455/src/bytes/bytes.go:1006
	_go_fuzz_dep_.CoverTab[786471] = 0
						for len(s) > 0 {
//line /snap/go/10455/src/bytes/bytes.go:1007
		if _go_fuzz_dep_.CoverTab[786471] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1007
			_go_fuzz_dep_.CoverTab[524847]++
//line /snap/go/10455/src/bytes/bytes.go:1007
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1007
			_go_fuzz_dep_.CoverTab[524848]++
//line /snap/go/10455/src/bytes/bytes.go:1007
		}
//line /snap/go/10455/src/bytes/bytes.go:1007
		_go_fuzz_dep_.CoverTab[786471] = 1
//line /snap/go/10455/src/bytes/bytes.go:1007
		_go_fuzz_dep_.CoverTab[622]++
							r, n := rune(s[0]), 1
							if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:1009
			_go_fuzz_dep_.CoverTab[524601]++
//line /snap/go/10455/src/bytes/bytes.go:1009
			_go_fuzz_dep_.CoverTab[625]++
								r, n = utf8.DecodeRune(s)
//line /snap/go/10455/src/bytes/bytes.go:1010
			// _ = "end of CoverTab[625]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1011
			_go_fuzz_dep_.CoverTab[524602]++
//line /snap/go/10455/src/bytes/bytes.go:1011
			_go_fuzz_dep_.CoverTab[626]++
//line /snap/go/10455/src/bytes/bytes.go:1011
			// _ = "end of CoverTab[626]"
//line /snap/go/10455/src/bytes/bytes.go:1011
		}
//line /snap/go/10455/src/bytes/bytes.go:1011
		// _ = "end of CoverTab[622]"
//line /snap/go/10455/src/bytes/bytes.go:1011
		_go_fuzz_dep_.CoverTab[623]++
							if !containsRune(cutset, r) {
//line /snap/go/10455/src/bytes/bytes.go:1012
			_go_fuzz_dep_.CoverTab[524603]++
//line /snap/go/10455/src/bytes/bytes.go:1012
			_go_fuzz_dep_.CoverTab[627]++
								break
//line /snap/go/10455/src/bytes/bytes.go:1013
			// _ = "end of CoverTab[627]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1014
			_go_fuzz_dep_.CoverTab[524604]++
//line /snap/go/10455/src/bytes/bytes.go:1014
			_go_fuzz_dep_.CoverTab[628]++
//line /snap/go/10455/src/bytes/bytes.go:1014
			// _ = "end of CoverTab[628]"
//line /snap/go/10455/src/bytes/bytes.go:1014
		}
//line /snap/go/10455/src/bytes/bytes.go:1014
		// _ = "end of CoverTab[623]"
//line /snap/go/10455/src/bytes/bytes.go:1014
		_go_fuzz_dep_.CoverTab[624]++
							s = s[n:]
//line /snap/go/10455/src/bytes/bytes.go:1015
		// _ = "end of CoverTab[624]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1016
	if _go_fuzz_dep_.CoverTab[786471] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1016
		_go_fuzz_dep_.CoverTab[524849]++
//line /snap/go/10455/src/bytes/bytes.go:1016
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1016
		_go_fuzz_dep_.CoverTab[524850]++
//line /snap/go/10455/src/bytes/bytes.go:1016
	}
//line /snap/go/10455/src/bytes/bytes.go:1016
	// _ = "end of CoverTab[619]"
//line /snap/go/10455/src/bytes/bytes.go:1016
	_go_fuzz_dep_.CoverTab[620]++
						if len(s) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1017
		_go_fuzz_dep_.CoverTab[524605]++
//line /snap/go/10455/src/bytes/bytes.go:1017
		_go_fuzz_dep_.CoverTab[629]++

							return nil
//line /snap/go/10455/src/bytes/bytes.go:1019
		// _ = "end of CoverTab[629]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1020
		_go_fuzz_dep_.CoverTab[524606]++
//line /snap/go/10455/src/bytes/bytes.go:1020
		_go_fuzz_dep_.CoverTab[630]++
//line /snap/go/10455/src/bytes/bytes.go:1020
		// _ = "end of CoverTab[630]"
//line /snap/go/10455/src/bytes/bytes.go:1020
	}
//line /snap/go/10455/src/bytes/bytes.go:1020
	// _ = "end of CoverTab[620]"
//line /snap/go/10455/src/bytes/bytes.go:1020
	_go_fuzz_dep_.CoverTab[621]++
						return s
//line /snap/go/10455/src/bytes/bytes.go:1021
	// _ = "end of CoverTab[621]"
}

// TrimRight returns a subslice of s by slicing off all trailing
//line /snap/go/10455/src/bytes/bytes.go:1024
// UTF-8-encoded code points that are contained in cutset.
//line /snap/go/10455/src/bytes/bytes.go:1026
func TrimRight(s []byte, cutset string) []byte {
//line /snap/go/10455/src/bytes/bytes.go:1026
	_go_fuzz_dep_.CoverTab[631]++
						if len(s) == 0 || func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1027
		_go_fuzz_dep_.CoverTab[635]++
//line /snap/go/10455/src/bytes/bytes.go:1027
		return cutset == ""
//line /snap/go/10455/src/bytes/bytes.go:1027
		// _ = "end of CoverTab[635]"
//line /snap/go/10455/src/bytes/bytes.go:1027
	}() {
//line /snap/go/10455/src/bytes/bytes.go:1027
		_go_fuzz_dep_.CoverTab[524607]++
//line /snap/go/10455/src/bytes/bytes.go:1027
		_go_fuzz_dep_.CoverTab[636]++
							return s
//line /snap/go/10455/src/bytes/bytes.go:1028
		// _ = "end of CoverTab[636]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1029
		_go_fuzz_dep_.CoverTab[524608]++
//line /snap/go/10455/src/bytes/bytes.go:1029
		_go_fuzz_dep_.CoverTab[637]++
//line /snap/go/10455/src/bytes/bytes.go:1029
		// _ = "end of CoverTab[637]"
//line /snap/go/10455/src/bytes/bytes.go:1029
	}
//line /snap/go/10455/src/bytes/bytes.go:1029
	// _ = "end of CoverTab[631]"
//line /snap/go/10455/src/bytes/bytes.go:1029
	_go_fuzz_dep_.CoverTab[632]++
						if len(cutset) == 1 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1030
		_go_fuzz_dep_.CoverTab[638]++
//line /snap/go/10455/src/bytes/bytes.go:1030
		return cutset[0] < utf8.RuneSelf
//line /snap/go/10455/src/bytes/bytes.go:1030
		// _ = "end of CoverTab[638]"
//line /snap/go/10455/src/bytes/bytes.go:1030
	}() {
//line /snap/go/10455/src/bytes/bytes.go:1030
		_go_fuzz_dep_.CoverTab[524609]++
//line /snap/go/10455/src/bytes/bytes.go:1030
		_go_fuzz_dep_.CoverTab[639]++
							return trimRightByte(s, cutset[0])
//line /snap/go/10455/src/bytes/bytes.go:1031
		// _ = "end of CoverTab[639]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1032
		_go_fuzz_dep_.CoverTab[524610]++
//line /snap/go/10455/src/bytes/bytes.go:1032
		_go_fuzz_dep_.CoverTab[640]++
//line /snap/go/10455/src/bytes/bytes.go:1032
		// _ = "end of CoverTab[640]"
//line /snap/go/10455/src/bytes/bytes.go:1032
	}
//line /snap/go/10455/src/bytes/bytes.go:1032
	// _ = "end of CoverTab[632]"
//line /snap/go/10455/src/bytes/bytes.go:1032
	_go_fuzz_dep_.CoverTab[633]++
						if as, ok := makeASCIISet(cutset); ok {
//line /snap/go/10455/src/bytes/bytes.go:1033
		_go_fuzz_dep_.CoverTab[524611]++
//line /snap/go/10455/src/bytes/bytes.go:1033
		_go_fuzz_dep_.CoverTab[641]++
							return trimRightASCII(s, &as)
//line /snap/go/10455/src/bytes/bytes.go:1034
		// _ = "end of CoverTab[641]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1035
		_go_fuzz_dep_.CoverTab[524612]++
//line /snap/go/10455/src/bytes/bytes.go:1035
		_go_fuzz_dep_.CoverTab[642]++
//line /snap/go/10455/src/bytes/bytes.go:1035
		// _ = "end of CoverTab[642]"
//line /snap/go/10455/src/bytes/bytes.go:1035
	}
//line /snap/go/10455/src/bytes/bytes.go:1035
	// _ = "end of CoverTab[633]"
//line /snap/go/10455/src/bytes/bytes.go:1035
	_go_fuzz_dep_.CoverTab[634]++
						return trimRightUnicode(s, cutset)
//line /snap/go/10455/src/bytes/bytes.go:1036
	// _ = "end of CoverTab[634]"
}

func trimRightByte(s []byte, c byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:1039
	_go_fuzz_dep_.CoverTab[643]++
//line /snap/go/10455/src/bytes/bytes.go:1039
	_go_fuzz_dep_.CoverTab[786472] = 0
						for len(s) > 0 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1040
		_go_fuzz_dep_.CoverTab[645]++
//line /snap/go/10455/src/bytes/bytes.go:1040
		return s[len(s)-1] == c
//line /snap/go/10455/src/bytes/bytes.go:1040
		// _ = "end of CoverTab[645]"
//line /snap/go/10455/src/bytes/bytes.go:1040
	}() {
//line /snap/go/10455/src/bytes/bytes.go:1040
		if _go_fuzz_dep_.CoverTab[786472] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1040
			_go_fuzz_dep_.CoverTab[524851]++
//line /snap/go/10455/src/bytes/bytes.go:1040
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1040
			_go_fuzz_dep_.CoverTab[524852]++
//line /snap/go/10455/src/bytes/bytes.go:1040
		}
//line /snap/go/10455/src/bytes/bytes.go:1040
		_go_fuzz_dep_.CoverTab[786472] = 1
//line /snap/go/10455/src/bytes/bytes.go:1040
		_go_fuzz_dep_.CoverTab[646]++
							s = s[:len(s)-1]
//line /snap/go/10455/src/bytes/bytes.go:1041
		// _ = "end of CoverTab[646]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1042
	if _go_fuzz_dep_.CoverTab[786472] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1042
		_go_fuzz_dep_.CoverTab[524853]++
//line /snap/go/10455/src/bytes/bytes.go:1042
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1042
		_go_fuzz_dep_.CoverTab[524854]++
//line /snap/go/10455/src/bytes/bytes.go:1042
	}
//line /snap/go/10455/src/bytes/bytes.go:1042
	// _ = "end of CoverTab[643]"
//line /snap/go/10455/src/bytes/bytes.go:1042
	_go_fuzz_dep_.CoverTab[644]++
						return s
//line /snap/go/10455/src/bytes/bytes.go:1043
	// _ = "end of CoverTab[644]"
}

func trimRightASCII(s []byte, as *asciiSet) []byte {
//line /snap/go/10455/src/bytes/bytes.go:1046
	_go_fuzz_dep_.CoverTab[647]++
//line /snap/go/10455/src/bytes/bytes.go:1046
	_go_fuzz_dep_.CoverTab[786473] = 0
						for len(s) > 0 {
//line /snap/go/10455/src/bytes/bytes.go:1047
		if _go_fuzz_dep_.CoverTab[786473] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1047
			_go_fuzz_dep_.CoverTab[524855]++
//line /snap/go/10455/src/bytes/bytes.go:1047
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1047
			_go_fuzz_dep_.CoverTab[524856]++
//line /snap/go/10455/src/bytes/bytes.go:1047
		}
//line /snap/go/10455/src/bytes/bytes.go:1047
		_go_fuzz_dep_.CoverTab[786473] = 1
//line /snap/go/10455/src/bytes/bytes.go:1047
		_go_fuzz_dep_.CoverTab[649]++
							if !as.contains(s[len(s)-1]) {
//line /snap/go/10455/src/bytes/bytes.go:1048
			_go_fuzz_dep_.CoverTab[524613]++
//line /snap/go/10455/src/bytes/bytes.go:1048
			_go_fuzz_dep_.CoverTab[651]++
								break
//line /snap/go/10455/src/bytes/bytes.go:1049
			// _ = "end of CoverTab[651]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1050
			_go_fuzz_dep_.CoverTab[524614]++
//line /snap/go/10455/src/bytes/bytes.go:1050
			_go_fuzz_dep_.CoverTab[652]++
//line /snap/go/10455/src/bytes/bytes.go:1050
			// _ = "end of CoverTab[652]"
//line /snap/go/10455/src/bytes/bytes.go:1050
		}
//line /snap/go/10455/src/bytes/bytes.go:1050
		// _ = "end of CoverTab[649]"
//line /snap/go/10455/src/bytes/bytes.go:1050
		_go_fuzz_dep_.CoverTab[650]++
							s = s[:len(s)-1]
//line /snap/go/10455/src/bytes/bytes.go:1051
		// _ = "end of CoverTab[650]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1052
	if _go_fuzz_dep_.CoverTab[786473] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1052
		_go_fuzz_dep_.CoverTab[524857]++
//line /snap/go/10455/src/bytes/bytes.go:1052
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1052
		_go_fuzz_dep_.CoverTab[524858]++
//line /snap/go/10455/src/bytes/bytes.go:1052
	}
//line /snap/go/10455/src/bytes/bytes.go:1052
	// _ = "end of CoverTab[647]"
//line /snap/go/10455/src/bytes/bytes.go:1052
	_go_fuzz_dep_.CoverTab[648]++
						return s
//line /snap/go/10455/src/bytes/bytes.go:1053
	// _ = "end of CoverTab[648]"
}

func trimRightUnicode(s []byte, cutset string) []byte {
//line /snap/go/10455/src/bytes/bytes.go:1056
	_go_fuzz_dep_.CoverTab[653]++
//line /snap/go/10455/src/bytes/bytes.go:1056
	_go_fuzz_dep_.CoverTab[786474] = 0
						for len(s) > 0 {
//line /snap/go/10455/src/bytes/bytes.go:1057
		if _go_fuzz_dep_.CoverTab[786474] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1057
			_go_fuzz_dep_.CoverTab[524859]++
//line /snap/go/10455/src/bytes/bytes.go:1057
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1057
			_go_fuzz_dep_.CoverTab[524860]++
//line /snap/go/10455/src/bytes/bytes.go:1057
		}
//line /snap/go/10455/src/bytes/bytes.go:1057
		_go_fuzz_dep_.CoverTab[786474] = 1
//line /snap/go/10455/src/bytes/bytes.go:1057
		_go_fuzz_dep_.CoverTab[655]++
							r, n := rune(s[len(s)-1]), 1
							if r >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:1059
			_go_fuzz_dep_.CoverTab[524615]++
//line /snap/go/10455/src/bytes/bytes.go:1059
			_go_fuzz_dep_.CoverTab[658]++
								r, n = utf8.DecodeLastRune(s)
//line /snap/go/10455/src/bytes/bytes.go:1060
			// _ = "end of CoverTab[658]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1061
			_go_fuzz_dep_.CoverTab[524616]++
//line /snap/go/10455/src/bytes/bytes.go:1061
			_go_fuzz_dep_.CoverTab[659]++
//line /snap/go/10455/src/bytes/bytes.go:1061
			// _ = "end of CoverTab[659]"
//line /snap/go/10455/src/bytes/bytes.go:1061
		}
//line /snap/go/10455/src/bytes/bytes.go:1061
		// _ = "end of CoverTab[655]"
//line /snap/go/10455/src/bytes/bytes.go:1061
		_go_fuzz_dep_.CoverTab[656]++
							if !containsRune(cutset, r) {
//line /snap/go/10455/src/bytes/bytes.go:1062
			_go_fuzz_dep_.CoverTab[524617]++
//line /snap/go/10455/src/bytes/bytes.go:1062
			_go_fuzz_dep_.CoverTab[660]++
								break
//line /snap/go/10455/src/bytes/bytes.go:1063
			// _ = "end of CoverTab[660]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1064
			_go_fuzz_dep_.CoverTab[524618]++
//line /snap/go/10455/src/bytes/bytes.go:1064
			_go_fuzz_dep_.CoverTab[661]++
//line /snap/go/10455/src/bytes/bytes.go:1064
			// _ = "end of CoverTab[661]"
//line /snap/go/10455/src/bytes/bytes.go:1064
		}
//line /snap/go/10455/src/bytes/bytes.go:1064
		// _ = "end of CoverTab[656]"
//line /snap/go/10455/src/bytes/bytes.go:1064
		_go_fuzz_dep_.CoverTab[657]++
							s = s[:len(s)-n]
//line /snap/go/10455/src/bytes/bytes.go:1065
		// _ = "end of CoverTab[657]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1066
	if _go_fuzz_dep_.CoverTab[786474] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1066
		_go_fuzz_dep_.CoverTab[524861]++
//line /snap/go/10455/src/bytes/bytes.go:1066
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1066
		_go_fuzz_dep_.CoverTab[524862]++
//line /snap/go/10455/src/bytes/bytes.go:1066
	}
//line /snap/go/10455/src/bytes/bytes.go:1066
	// _ = "end of CoverTab[653]"
//line /snap/go/10455/src/bytes/bytes.go:1066
	_go_fuzz_dep_.CoverTab[654]++
						return s
//line /snap/go/10455/src/bytes/bytes.go:1067
	// _ = "end of CoverTab[654]"
}

// TrimSpace returns a subslice of s by slicing off all leading and
//line /snap/go/10455/src/bytes/bytes.go:1070
// trailing white space, as defined by Unicode.
//line /snap/go/10455/src/bytes/bytes.go:1072
func TrimSpace(s []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:1072
	_go_fuzz_dep_.CoverTab[662]++

						start := 0
//line /snap/go/10455/src/bytes/bytes.go:1074
	_go_fuzz_dep_.CoverTab[786475] = 0
						for ; start < len(s); start++ {
//line /snap/go/10455/src/bytes/bytes.go:1075
		if _go_fuzz_dep_.CoverTab[786475] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1075
			_go_fuzz_dep_.CoverTab[524863]++
//line /snap/go/10455/src/bytes/bytes.go:1075
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1075
			_go_fuzz_dep_.CoverTab[524864]++
//line /snap/go/10455/src/bytes/bytes.go:1075
		}
//line /snap/go/10455/src/bytes/bytes.go:1075
		_go_fuzz_dep_.CoverTab[786475] = 1
//line /snap/go/10455/src/bytes/bytes.go:1075
		_go_fuzz_dep_.CoverTab[666]++
							c := s[start]
							if c >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:1077
			_go_fuzz_dep_.CoverTab[524619]++
//line /snap/go/10455/src/bytes/bytes.go:1077
			_go_fuzz_dep_.CoverTab[668]++

//line /snap/go/10455/src/bytes/bytes.go:1080
			return TrimFunc(s[start:], unicode.IsSpace)
//line /snap/go/10455/src/bytes/bytes.go:1080
			// _ = "end of CoverTab[668]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1081
			_go_fuzz_dep_.CoverTab[524620]++
//line /snap/go/10455/src/bytes/bytes.go:1081
			_go_fuzz_dep_.CoverTab[669]++
//line /snap/go/10455/src/bytes/bytes.go:1081
			// _ = "end of CoverTab[669]"
//line /snap/go/10455/src/bytes/bytes.go:1081
		}
//line /snap/go/10455/src/bytes/bytes.go:1081
		// _ = "end of CoverTab[666]"
//line /snap/go/10455/src/bytes/bytes.go:1081
		_go_fuzz_dep_.CoverTab[667]++
							if asciiSpace[c] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1082
			_go_fuzz_dep_.CoverTab[524621]++
//line /snap/go/10455/src/bytes/bytes.go:1082
			_go_fuzz_dep_.CoverTab[670]++
								break
//line /snap/go/10455/src/bytes/bytes.go:1083
			// _ = "end of CoverTab[670]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1084
			_go_fuzz_dep_.CoverTab[524622]++
//line /snap/go/10455/src/bytes/bytes.go:1084
			_go_fuzz_dep_.CoverTab[671]++
//line /snap/go/10455/src/bytes/bytes.go:1084
			// _ = "end of CoverTab[671]"
//line /snap/go/10455/src/bytes/bytes.go:1084
		}
//line /snap/go/10455/src/bytes/bytes.go:1084
		// _ = "end of CoverTab[667]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1085
	if _go_fuzz_dep_.CoverTab[786475] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1085
		_go_fuzz_dep_.CoverTab[524865]++
//line /snap/go/10455/src/bytes/bytes.go:1085
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1085
		_go_fuzz_dep_.CoverTab[524866]++
//line /snap/go/10455/src/bytes/bytes.go:1085
	}
//line /snap/go/10455/src/bytes/bytes.go:1085
	// _ = "end of CoverTab[662]"
//line /snap/go/10455/src/bytes/bytes.go:1085
	_go_fuzz_dep_.CoverTab[663]++

//line /snap/go/10455/src/bytes/bytes.go:1088
	stop := len(s)
//line /snap/go/10455/src/bytes/bytes.go:1088
	_go_fuzz_dep_.CoverTab[786476] = 0
						for ; stop > start; stop-- {
//line /snap/go/10455/src/bytes/bytes.go:1089
		if _go_fuzz_dep_.CoverTab[786476] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1089
			_go_fuzz_dep_.CoverTab[524867]++
//line /snap/go/10455/src/bytes/bytes.go:1089
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1089
			_go_fuzz_dep_.CoverTab[524868]++
//line /snap/go/10455/src/bytes/bytes.go:1089
		}
//line /snap/go/10455/src/bytes/bytes.go:1089
		_go_fuzz_dep_.CoverTab[786476] = 1
//line /snap/go/10455/src/bytes/bytes.go:1089
		_go_fuzz_dep_.CoverTab[672]++
							c := s[stop-1]
							if c >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:1091
			_go_fuzz_dep_.CoverTab[524623]++
//line /snap/go/10455/src/bytes/bytes.go:1091
			_go_fuzz_dep_.CoverTab[674]++
								return TrimFunc(s[start:stop], unicode.IsSpace)
//line /snap/go/10455/src/bytes/bytes.go:1092
			// _ = "end of CoverTab[674]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1093
			_go_fuzz_dep_.CoverTab[524624]++
//line /snap/go/10455/src/bytes/bytes.go:1093
			_go_fuzz_dep_.CoverTab[675]++
//line /snap/go/10455/src/bytes/bytes.go:1093
			// _ = "end of CoverTab[675]"
//line /snap/go/10455/src/bytes/bytes.go:1093
		}
//line /snap/go/10455/src/bytes/bytes.go:1093
		// _ = "end of CoverTab[672]"
//line /snap/go/10455/src/bytes/bytes.go:1093
		_go_fuzz_dep_.CoverTab[673]++
							if asciiSpace[c] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1094
			_go_fuzz_dep_.CoverTab[524625]++
//line /snap/go/10455/src/bytes/bytes.go:1094
			_go_fuzz_dep_.CoverTab[676]++
								break
//line /snap/go/10455/src/bytes/bytes.go:1095
			// _ = "end of CoverTab[676]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1096
			_go_fuzz_dep_.CoverTab[524626]++
//line /snap/go/10455/src/bytes/bytes.go:1096
			_go_fuzz_dep_.CoverTab[677]++
//line /snap/go/10455/src/bytes/bytes.go:1096
			// _ = "end of CoverTab[677]"
//line /snap/go/10455/src/bytes/bytes.go:1096
		}
//line /snap/go/10455/src/bytes/bytes.go:1096
		// _ = "end of CoverTab[673]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1097
	if _go_fuzz_dep_.CoverTab[786476] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1097
		_go_fuzz_dep_.CoverTab[524869]++
//line /snap/go/10455/src/bytes/bytes.go:1097
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1097
		_go_fuzz_dep_.CoverTab[524870]++
//line /snap/go/10455/src/bytes/bytes.go:1097
	}
//line /snap/go/10455/src/bytes/bytes.go:1097
	// _ = "end of CoverTab[663]"
//line /snap/go/10455/src/bytes/bytes.go:1097
	_go_fuzz_dep_.CoverTab[664]++

//line /snap/go/10455/src/bytes/bytes.go:1102
	if start == stop {
//line /snap/go/10455/src/bytes/bytes.go:1102
		_go_fuzz_dep_.CoverTab[524627]++
//line /snap/go/10455/src/bytes/bytes.go:1102
		_go_fuzz_dep_.CoverTab[678]++

//line /snap/go/10455/src/bytes/bytes.go:1105
		return nil
//line /snap/go/10455/src/bytes/bytes.go:1105
		// _ = "end of CoverTab[678]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1106
		_go_fuzz_dep_.CoverTab[524628]++
//line /snap/go/10455/src/bytes/bytes.go:1106
		_go_fuzz_dep_.CoverTab[679]++
//line /snap/go/10455/src/bytes/bytes.go:1106
		// _ = "end of CoverTab[679]"
//line /snap/go/10455/src/bytes/bytes.go:1106
	}
//line /snap/go/10455/src/bytes/bytes.go:1106
	// _ = "end of CoverTab[664]"
//line /snap/go/10455/src/bytes/bytes.go:1106
	_go_fuzz_dep_.CoverTab[665]++
						return s[start:stop]
//line /snap/go/10455/src/bytes/bytes.go:1107
	// _ = "end of CoverTab[665]"
}

// Runes interprets s as a sequence of UTF-8-encoded code points.
//line /snap/go/10455/src/bytes/bytes.go:1110
// It returns a slice of runes (Unicode code points) equivalent to s.
//line /snap/go/10455/src/bytes/bytes.go:1112
func Runes(s []byte) []rune {
//line /snap/go/10455/src/bytes/bytes.go:1112
	_go_fuzz_dep_.CoverTab[680]++
						t := make([]rune, utf8.RuneCount(s))
						i := 0
//line /snap/go/10455/src/bytes/bytes.go:1114
	_go_fuzz_dep_.CoverTab[786477] = 0
						for len(s) > 0 {
//line /snap/go/10455/src/bytes/bytes.go:1115
		if _go_fuzz_dep_.CoverTab[786477] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1115
			_go_fuzz_dep_.CoverTab[524871]++
//line /snap/go/10455/src/bytes/bytes.go:1115
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1115
			_go_fuzz_dep_.CoverTab[524872]++
//line /snap/go/10455/src/bytes/bytes.go:1115
		}
//line /snap/go/10455/src/bytes/bytes.go:1115
		_go_fuzz_dep_.CoverTab[786477] = 1
//line /snap/go/10455/src/bytes/bytes.go:1115
		_go_fuzz_dep_.CoverTab[682]++
							r, l := utf8.DecodeRune(s)
							t[i] = r
							i++
							s = s[l:]
//line /snap/go/10455/src/bytes/bytes.go:1119
		// _ = "end of CoverTab[682]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1120
	if _go_fuzz_dep_.CoverTab[786477] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1120
		_go_fuzz_dep_.CoverTab[524873]++
//line /snap/go/10455/src/bytes/bytes.go:1120
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1120
		_go_fuzz_dep_.CoverTab[524874]++
//line /snap/go/10455/src/bytes/bytes.go:1120
	}
//line /snap/go/10455/src/bytes/bytes.go:1120
	// _ = "end of CoverTab[680]"
//line /snap/go/10455/src/bytes/bytes.go:1120
	_go_fuzz_dep_.CoverTab[681]++
						return t
//line /snap/go/10455/src/bytes/bytes.go:1121
	// _ = "end of CoverTab[681]"
}

// Replace returns a copy of the slice s with the first n
//line /snap/go/10455/src/bytes/bytes.go:1124
// non-overlapping instances of old replaced by new.
//line /snap/go/10455/src/bytes/bytes.go:1124
// If old is empty, it matches at the beginning of the slice
//line /snap/go/10455/src/bytes/bytes.go:1124
// and after each UTF-8 sequence, yielding up to k+1 replacements
//line /snap/go/10455/src/bytes/bytes.go:1124
// for a k-rune slice.
//line /snap/go/10455/src/bytes/bytes.go:1124
// If n < 0, there is no limit on the number of replacements.
//line /snap/go/10455/src/bytes/bytes.go:1130
func Replace(s, old, new []byte, n int) []byte {
//line /snap/go/10455/src/bytes/bytes.go:1130
	_go_fuzz_dep_.CoverTab[683]++
						m := 0
						if n != 0 {
//line /snap/go/10455/src/bytes/bytes.go:1132
		_go_fuzz_dep_.CoverTab[524629]++
//line /snap/go/10455/src/bytes/bytes.go:1132
		_go_fuzz_dep_.CoverTab[688]++

							m = Count(s, old)
//line /snap/go/10455/src/bytes/bytes.go:1134
		// _ = "end of CoverTab[688]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1135
		_go_fuzz_dep_.CoverTab[524630]++
//line /snap/go/10455/src/bytes/bytes.go:1135
		_go_fuzz_dep_.CoverTab[689]++
//line /snap/go/10455/src/bytes/bytes.go:1135
		// _ = "end of CoverTab[689]"
//line /snap/go/10455/src/bytes/bytes.go:1135
	}
//line /snap/go/10455/src/bytes/bytes.go:1135
	// _ = "end of CoverTab[683]"
//line /snap/go/10455/src/bytes/bytes.go:1135
	_go_fuzz_dep_.CoverTab[684]++
						if m == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1136
		_go_fuzz_dep_.CoverTab[524631]++
//line /snap/go/10455/src/bytes/bytes.go:1136
		_go_fuzz_dep_.CoverTab[690]++

							return append([]byte(nil), s...)
//line /snap/go/10455/src/bytes/bytes.go:1138
		// _ = "end of CoverTab[690]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1139
		_go_fuzz_dep_.CoverTab[524632]++
//line /snap/go/10455/src/bytes/bytes.go:1139
		_go_fuzz_dep_.CoverTab[691]++
//line /snap/go/10455/src/bytes/bytes.go:1139
		// _ = "end of CoverTab[691]"
//line /snap/go/10455/src/bytes/bytes.go:1139
	}
//line /snap/go/10455/src/bytes/bytes.go:1139
	// _ = "end of CoverTab[684]"
//line /snap/go/10455/src/bytes/bytes.go:1139
	_go_fuzz_dep_.CoverTab[685]++
						if n < 0 || func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1140
		_go_fuzz_dep_.CoverTab[692]++
//line /snap/go/10455/src/bytes/bytes.go:1140
		return m < n
//line /snap/go/10455/src/bytes/bytes.go:1140
		// _ = "end of CoverTab[692]"
//line /snap/go/10455/src/bytes/bytes.go:1140
	}() {
//line /snap/go/10455/src/bytes/bytes.go:1140
		_go_fuzz_dep_.CoverTab[524633]++
//line /snap/go/10455/src/bytes/bytes.go:1140
		_go_fuzz_dep_.CoverTab[693]++
							n = m
//line /snap/go/10455/src/bytes/bytes.go:1141
		// _ = "end of CoverTab[693]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1142
		_go_fuzz_dep_.CoverTab[524634]++
//line /snap/go/10455/src/bytes/bytes.go:1142
		_go_fuzz_dep_.CoverTab[694]++
//line /snap/go/10455/src/bytes/bytes.go:1142
		// _ = "end of CoverTab[694]"
//line /snap/go/10455/src/bytes/bytes.go:1142
	}
//line /snap/go/10455/src/bytes/bytes.go:1142
	// _ = "end of CoverTab[685]"
//line /snap/go/10455/src/bytes/bytes.go:1142
	_go_fuzz_dep_.CoverTab[686]++

//line /snap/go/10455/src/bytes/bytes.go:1145
	t := make([]byte, len(s)+n*(len(new)-len(old)))
						w := 0
						start := 0
//line /snap/go/10455/src/bytes/bytes.go:1147
	_go_fuzz_dep_.CoverTab[786478] = 0
						for i := 0; i < n; i++ {
//line /snap/go/10455/src/bytes/bytes.go:1148
		if _go_fuzz_dep_.CoverTab[786478] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1148
			_go_fuzz_dep_.CoverTab[524875]++
//line /snap/go/10455/src/bytes/bytes.go:1148
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1148
			_go_fuzz_dep_.CoverTab[524876]++
//line /snap/go/10455/src/bytes/bytes.go:1148
		}
//line /snap/go/10455/src/bytes/bytes.go:1148
		_go_fuzz_dep_.CoverTab[786478] = 1
//line /snap/go/10455/src/bytes/bytes.go:1148
		_go_fuzz_dep_.CoverTab[695]++
							j := start
							if len(old) == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1150
			_go_fuzz_dep_.CoverTab[524635]++
//line /snap/go/10455/src/bytes/bytes.go:1150
			_go_fuzz_dep_.CoverTab[697]++
								if i > 0 {
//line /snap/go/10455/src/bytes/bytes.go:1151
				_go_fuzz_dep_.CoverTab[524637]++
//line /snap/go/10455/src/bytes/bytes.go:1151
				_go_fuzz_dep_.CoverTab[698]++
									_, wid := utf8.DecodeRune(s[start:])
									j += wid
//line /snap/go/10455/src/bytes/bytes.go:1153
				// _ = "end of CoverTab[698]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:1154
				_go_fuzz_dep_.CoverTab[524638]++
//line /snap/go/10455/src/bytes/bytes.go:1154
				_go_fuzz_dep_.CoverTab[699]++
//line /snap/go/10455/src/bytes/bytes.go:1154
				// _ = "end of CoverTab[699]"
//line /snap/go/10455/src/bytes/bytes.go:1154
			}
//line /snap/go/10455/src/bytes/bytes.go:1154
			// _ = "end of CoverTab[697]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1155
			_go_fuzz_dep_.CoverTab[524636]++
//line /snap/go/10455/src/bytes/bytes.go:1155
			_go_fuzz_dep_.CoverTab[700]++
								j += Index(s[start:], old)
//line /snap/go/10455/src/bytes/bytes.go:1156
			// _ = "end of CoverTab[700]"
		}
//line /snap/go/10455/src/bytes/bytes.go:1157
		// _ = "end of CoverTab[695]"
//line /snap/go/10455/src/bytes/bytes.go:1157
		_go_fuzz_dep_.CoverTab[696]++
							w += copy(t[w:], s[start:j])
							w += copy(t[w:], new)
							start = j + len(old)
//line /snap/go/10455/src/bytes/bytes.go:1160
		// _ = "end of CoverTab[696]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1161
	if _go_fuzz_dep_.CoverTab[786478] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1161
		_go_fuzz_dep_.CoverTab[524877]++
//line /snap/go/10455/src/bytes/bytes.go:1161
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1161
		_go_fuzz_dep_.CoverTab[524878]++
//line /snap/go/10455/src/bytes/bytes.go:1161
	}
//line /snap/go/10455/src/bytes/bytes.go:1161
	// _ = "end of CoverTab[686]"
//line /snap/go/10455/src/bytes/bytes.go:1161
	_go_fuzz_dep_.CoverTab[687]++
						w += copy(t[w:], s[start:])
						return t[0:w]
//line /snap/go/10455/src/bytes/bytes.go:1163
	// _ = "end of CoverTab[687]"
}

// ReplaceAll returns a copy of the slice s with all
//line /snap/go/10455/src/bytes/bytes.go:1166
// non-overlapping instances of old replaced by new.
//line /snap/go/10455/src/bytes/bytes.go:1166
// If old is empty, it matches at the beginning of the slice
//line /snap/go/10455/src/bytes/bytes.go:1166
// and after each UTF-8 sequence, yielding up to k+1 replacements
//line /snap/go/10455/src/bytes/bytes.go:1166
// for a k-rune slice.
//line /snap/go/10455/src/bytes/bytes.go:1171
func ReplaceAll(s, old, new []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:1171
	_go_fuzz_dep_.CoverTab[701]++
						return Replace(s, old, new, -1)
//line /snap/go/10455/src/bytes/bytes.go:1172
	// _ = "end of CoverTab[701]"
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
//line /snap/go/10455/src/bytes/bytes.go:1175
// are equal under simple Unicode case-folding, which is a more general
//line /snap/go/10455/src/bytes/bytes.go:1175
// form of case-insensitivity.
//line /snap/go/10455/src/bytes/bytes.go:1178
func EqualFold(s, t []byte) bool {
//line /snap/go/10455/src/bytes/bytes.go:1178
	_go_fuzz_dep_.CoverTab[702]++

						i := 0
//line /snap/go/10455/src/bytes/bytes.go:1180
	_go_fuzz_dep_.CoverTab[786479] = 0
						for ; i < len(s) && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1181
		_go_fuzz_dep_.CoverTab[705]++
//line /snap/go/10455/src/bytes/bytes.go:1181
		return i < len(t)
//line /snap/go/10455/src/bytes/bytes.go:1181
		// _ = "end of CoverTab[705]"
//line /snap/go/10455/src/bytes/bytes.go:1181
	}(); i++ {
//line /snap/go/10455/src/bytes/bytes.go:1181
		if _go_fuzz_dep_.CoverTab[786479] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1181
			_go_fuzz_dep_.CoverTab[524879]++
//line /snap/go/10455/src/bytes/bytes.go:1181
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1181
			_go_fuzz_dep_.CoverTab[524880]++
//line /snap/go/10455/src/bytes/bytes.go:1181
		}
//line /snap/go/10455/src/bytes/bytes.go:1181
		_go_fuzz_dep_.CoverTab[786479] = 1
//line /snap/go/10455/src/bytes/bytes.go:1181
		_go_fuzz_dep_.CoverTab[706]++
							sr := s[i]
							tr := t[i]
							if sr|tr >= utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:1184
			_go_fuzz_dep_.CoverTab[524639]++
//line /snap/go/10455/src/bytes/bytes.go:1184
			_go_fuzz_dep_.CoverTab[711]++
								goto hasUnicode
//line /snap/go/10455/src/bytes/bytes.go:1185
			// _ = "end of CoverTab[711]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1186
			_go_fuzz_dep_.CoverTab[524640]++
//line /snap/go/10455/src/bytes/bytes.go:1186
			_go_fuzz_dep_.CoverTab[712]++
//line /snap/go/10455/src/bytes/bytes.go:1186
			// _ = "end of CoverTab[712]"
//line /snap/go/10455/src/bytes/bytes.go:1186
		}
//line /snap/go/10455/src/bytes/bytes.go:1186
		// _ = "end of CoverTab[706]"
//line /snap/go/10455/src/bytes/bytes.go:1186
		_go_fuzz_dep_.CoverTab[707]++

//line /snap/go/10455/src/bytes/bytes.go:1189
		if tr == sr {
//line /snap/go/10455/src/bytes/bytes.go:1189
			_go_fuzz_dep_.CoverTab[524641]++
//line /snap/go/10455/src/bytes/bytes.go:1189
			_go_fuzz_dep_.CoverTab[713]++
								continue
//line /snap/go/10455/src/bytes/bytes.go:1190
			// _ = "end of CoverTab[713]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1191
			_go_fuzz_dep_.CoverTab[524642]++
//line /snap/go/10455/src/bytes/bytes.go:1191
			_go_fuzz_dep_.CoverTab[714]++
//line /snap/go/10455/src/bytes/bytes.go:1191
			// _ = "end of CoverTab[714]"
//line /snap/go/10455/src/bytes/bytes.go:1191
		}
//line /snap/go/10455/src/bytes/bytes.go:1191
		// _ = "end of CoverTab[707]"
//line /snap/go/10455/src/bytes/bytes.go:1191
		_go_fuzz_dep_.CoverTab[708]++

//line /snap/go/10455/src/bytes/bytes.go:1194
		if tr < sr {
//line /snap/go/10455/src/bytes/bytes.go:1194
			_go_fuzz_dep_.CoverTab[524643]++
//line /snap/go/10455/src/bytes/bytes.go:1194
			_go_fuzz_dep_.CoverTab[715]++
								tr, sr = sr, tr
//line /snap/go/10455/src/bytes/bytes.go:1195
			// _ = "end of CoverTab[715]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1196
			_go_fuzz_dep_.CoverTab[524644]++
//line /snap/go/10455/src/bytes/bytes.go:1196
			_go_fuzz_dep_.CoverTab[716]++
//line /snap/go/10455/src/bytes/bytes.go:1196
			// _ = "end of CoverTab[716]"
//line /snap/go/10455/src/bytes/bytes.go:1196
		}
//line /snap/go/10455/src/bytes/bytes.go:1196
		// _ = "end of CoverTab[708]"
//line /snap/go/10455/src/bytes/bytes.go:1196
		_go_fuzz_dep_.CoverTab[709]++

							if 'A' <= sr && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1198
			_go_fuzz_dep_.CoverTab[717]++
//line /snap/go/10455/src/bytes/bytes.go:1198
			return sr <= 'Z'
//line /snap/go/10455/src/bytes/bytes.go:1198
			// _ = "end of CoverTab[717]"
//line /snap/go/10455/src/bytes/bytes.go:1198
		}() && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1198
			_go_fuzz_dep_.CoverTab[718]++
//line /snap/go/10455/src/bytes/bytes.go:1198
			return tr == sr+'a'-'A'
//line /snap/go/10455/src/bytes/bytes.go:1198
			// _ = "end of CoverTab[718]"
//line /snap/go/10455/src/bytes/bytes.go:1198
		}() {
//line /snap/go/10455/src/bytes/bytes.go:1198
			_go_fuzz_dep_.CoverTab[524645]++
//line /snap/go/10455/src/bytes/bytes.go:1198
			_go_fuzz_dep_.CoverTab[719]++
								continue
//line /snap/go/10455/src/bytes/bytes.go:1199
			// _ = "end of CoverTab[719]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1200
			_go_fuzz_dep_.CoverTab[524646]++
//line /snap/go/10455/src/bytes/bytes.go:1200
			_go_fuzz_dep_.CoverTab[720]++
//line /snap/go/10455/src/bytes/bytes.go:1200
			// _ = "end of CoverTab[720]"
//line /snap/go/10455/src/bytes/bytes.go:1200
		}
//line /snap/go/10455/src/bytes/bytes.go:1200
		// _ = "end of CoverTab[709]"
//line /snap/go/10455/src/bytes/bytes.go:1200
		_go_fuzz_dep_.CoverTab[710]++
							return false
//line /snap/go/10455/src/bytes/bytes.go:1201
		// _ = "end of CoverTab[710]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1202
	if _go_fuzz_dep_.CoverTab[786479] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1202
		_go_fuzz_dep_.CoverTab[524881]++
//line /snap/go/10455/src/bytes/bytes.go:1202
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1202
		_go_fuzz_dep_.CoverTab[524882]++
//line /snap/go/10455/src/bytes/bytes.go:1202
	}
//line /snap/go/10455/src/bytes/bytes.go:1202
	// _ = "end of CoverTab[702]"
//line /snap/go/10455/src/bytes/bytes.go:1202
	_go_fuzz_dep_.CoverTab[703]++

						return len(s) == len(t)

hasUnicode:
						s = s[i:]
						t = t[i:]
//line /snap/go/10455/src/bytes/bytes.go:1208
	_go_fuzz_dep_.CoverTab[786480] = 0
						for len(s) != 0 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1209
		_go_fuzz_dep_.CoverTab[721]++
//line /snap/go/10455/src/bytes/bytes.go:1209
		return len(t) != 0
//line /snap/go/10455/src/bytes/bytes.go:1209
		// _ = "end of CoverTab[721]"
//line /snap/go/10455/src/bytes/bytes.go:1209
	}() {
//line /snap/go/10455/src/bytes/bytes.go:1209
		if _go_fuzz_dep_.CoverTab[786480] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1209
			_go_fuzz_dep_.CoverTab[524883]++
//line /snap/go/10455/src/bytes/bytes.go:1209
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1209
			_go_fuzz_dep_.CoverTab[524884]++
//line /snap/go/10455/src/bytes/bytes.go:1209
		}
//line /snap/go/10455/src/bytes/bytes.go:1209
		_go_fuzz_dep_.CoverTab[786480] = 1
//line /snap/go/10455/src/bytes/bytes.go:1209
		_go_fuzz_dep_.CoverTab[722]++
		// Extract first rune from each.
		var sr, tr rune
		if s[0] < utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:1212
			_go_fuzz_dep_.CoverTab[524647]++
//line /snap/go/10455/src/bytes/bytes.go:1212
			_go_fuzz_dep_.CoverTab[730]++
								sr, s = rune(s[0]), s[1:]
//line /snap/go/10455/src/bytes/bytes.go:1213
			// _ = "end of CoverTab[730]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1214
			_go_fuzz_dep_.CoverTab[524648]++
//line /snap/go/10455/src/bytes/bytes.go:1214
			_go_fuzz_dep_.CoverTab[731]++
								r, size := utf8.DecodeRune(s)
								sr, s = r, s[size:]
//line /snap/go/10455/src/bytes/bytes.go:1216
			// _ = "end of CoverTab[731]"
		}
//line /snap/go/10455/src/bytes/bytes.go:1217
		// _ = "end of CoverTab[722]"
//line /snap/go/10455/src/bytes/bytes.go:1217
		_go_fuzz_dep_.CoverTab[723]++
							if t[0] < utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:1218
			_go_fuzz_dep_.CoverTab[524649]++
//line /snap/go/10455/src/bytes/bytes.go:1218
			_go_fuzz_dep_.CoverTab[732]++
								tr, t = rune(t[0]), t[1:]
//line /snap/go/10455/src/bytes/bytes.go:1219
			// _ = "end of CoverTab[732]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1220
			_go_fuzz_dep_.CoverTab[524650]++
//line /snap/go/10455/src/bytes/bytes.go:1220
			_go_fuzz_dep_.CoverTab[733]++
								r, size := utf8.DecodeRune(t)
								tr, t = r, t[size:]
//line /snap/go/10455/src/bytes/bytes.go:1222
			// _ = "end of CoverTab[733]"
		}
//line /snap/go/10455/src/bytes/bytes.go:1223
		// _ = "end of CoverTab[723]"
//line /snap/go/10455/src/bytes/bytes.go:1223
		_go_fuzz_dep_.CoverTab[724]++

//line /snap/go/10455/src/bytes/bytes.go:1228
		if tr == sr {
//line /snap/go/10455/src/bytes/bytes.go:1228
			_go_fuzz_dep_.CoverTab[524651]++
//line /snap/go/10455/src/bytes/bytes.go:1228
			_go_fuzz_dep_.CoverTab[734]++
								continue
//line /snap/go/10455/src/bytes/bytes.go:1229
			// _ = "end of CoverTab[734]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1230
			_go_fuzz_dep_.CoverTab[524652]++
//line /snap/go/10455/src/bytes/bytes.go:1230
			_go_fuzz_dep_.CoverTab[735]++
//line /snap/go/10455/src/bytes/bytes.go:1230
			// _ = "end of CoverTab[735]"
//line /snap/go/10455/src/bytes/bytes.go:1230
		}
//line /snap/go/10455/src/bytes/bytes.go:1230
		// _ = "end of CoverTab[724]"
//line /snap/go/10455/src/bytes/bytes.go:1230
		_go_fuzz_dep_.CoverTab[725]++

//line /snap/go/10455/src/bytes/bytes.go:1233
		if tr < sr {
//line /snap/go/10455/src/bytes/bytes.go:1233
			_go_fuzz_dep_.CoverTab[524653]++
//line /snap/go/10455/src/bytes/bytes.go:1233
			_go_fuzz_dep_.CoverTab[736]++
								tr, sr = sr, tr
//line /snap/go/10455/src/bytes/bytes.go:1234
			// _ = "end of CoverTab[736]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1235
			_go_fuzz_dep_.CoverTab[524654]++
//line /snap/go/10455/src/bytes/bytes.go:1235
			_go_fuzz_dep_.CoverTab[737]++
//line /snap/go/10455/src/bytes/bytes.go:1235
			// _ = "end of CoverTab[737]"
//line /snap/go/10455/src/bytes/bytes.go:1235
		}
//line /snap/go/10455/src/bytes/bytes.go:1235
		// _ = "end of CoverTab[725]"
//line /snap/go/10455/src/bytes/bytes.go:1235
		_go_fuzz_dep_.CoverTab[726]++

							if tr < utf8.RuneSelf {
//line /snap/go/10455/src/bytes/bytes.go:1237
			_go_fuzz_dep_.CoverTab[524655]++
//line /snap/go/10455/src/bytes/bytes.go:1237
			_go_fuzz_dep_.CoverTab[738]++

								if 'A' <= sr && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1239
				_go_fuzz_dep_.CoverTab[740]++
//line /snap/go/10455/src/bytes/bytes.go:1239
				return sr <= 'Z'
//line /snap/go/10455/src/bytes/bytes.go:1239
				// _ = "end of CoverTab[740]"
//line /snap/go/10455/src/bytes/bytes.go:1239
			}() && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1239
				_go_fuzz_dep_.CoverTab[741]++
//line /snap/go/10455/src/bytes/bytes.go:1239
				return tr == sr+'a'-'A'
//line /snap/go/10455/src/bytes/bytes.go:1239
				// _ = "end of CoverTab[741]"
//line /snap/go/10455/src/bytes/bytes.go:1239
			}() {
//line /snap/go/10455/src/bytes/bytes.go:1239
				_go_fuzz_dep_.CoverTab[524657]++
//line /snap/go/10455/src/bytes/bytes.go:1239
				_go_fuzz_dep_.CoverTab[742]++
									continue
//line /snap/go/10455/src/bytes/bytes.go:1240
				// _ = "end of CoverTab[742]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:1241
				_go_fuzz_dep_.CoverTab[524658]++
//line /snap/go/10455/src/bytes/bytes.go:1241
				_go_fuzz_dep_.CoverTab[743]++
//line /snap/go/10455/src/bytes/bytes.go:1241
				// _ = "end of CoverTab[743]"
//line /snap/go/10455/src/bytes/bytes.go:1241
			}
//line /snap/go/10455/src/bytes/bytes.go:1241
			// _ = "end of CoverTab[738]"
//line /snap/go/10455/src/bytes/bytes.go:1241
			_go_fuzz_dep_.CoverTab[739]++
								return false
//line /snap/go/10455/src/bytes/bytes.go:1242
			// _ = "end of CoverTab[739]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1243
			_go_fuzz_dep_.CoverTab[524656]++
//line /snap/go/10455/src/bytes/bytes.go:1243
			_go_fuzz_dep_.CoverTab[744]++
//line /snap/go/10455/src/bytes/bytes.go:1243
			// _ = "end of CoverTab[744]"
//line /snap/go/10455/src/bytes/bytes.go:1243
		}
//line /snap/go/10455/src/bytes/bytes.go:1243
		// _ = "end of CoverTab[726]"
//line /snap/go/10455/src/bytes/bytes.go:1243
		_go_fuzz_dep_.CoverTab[727]++

//line /snap/go/10455/src/bytes/bytes.go:1247
		r := unicode.SimpleFold(sr)
//line /snap/go/10455/src/bytes/bytes.go:1247
		_go_fuzz_dep_.CoverTab[786481] = 0
							for r != sr && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1248
			_go_fuzz_dep_.CoverTab[745]++
//line /snap/go/10455/src/bytes/bytes.go:1248
			return r < tr
//line /snap/go/10455/src/bytes/bytes.go:1248
			// _ = "end of CoverTab[745]"
//line /snap/go/10455/src/bytes/bytes.go:1248
		}() {
//line /snap/go/10455/src/bytes/bytes.go:1248
			if _go_fuzz_dep_.CoverTab[786481] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1248
				_go_fuzz_dep_.CoverTab[524887]++
//line /snap/go/10455/src/bytes/bytes.go:1248
			} else {
//line /snap/go/10455/src/bytes/bytes.go:1248
				_go_fuzz_dep_.CoverTab[524888]++
//line /snap/go/10455/src/bytes/bytes.go:1248
			}
//line /snap/go/10455/src/bytes/bytes.go:1248
			_go_fuzz_dep_.CoverTab[786481] = 1
//line /snap/go/10455/src/bytes/bytes.go:1248
			_go_fuzz_dep_.CoverTab[746]++
								r = unicode.SimpleFold(r)
//line /snap/go/10455/src/bytes/bytes.go:1249
			// _ = "end of CoverTab[746]"
		}
//line /snap/go/10455/src/bytes/bytes.go:1250
		if _go_fuzz_dep_.CoverTab[786481] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1250
			_go_fuzz_dep_.CoverTab[524889]++
//line /snap/go/10455/src/bytes/bytes.go:1250
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1250
			_go_fuzz_dep_.CoverTab[524890]++
//line /snap/go/10455/src/bytes/bytes.go:1250
		}
//line /snap/go/10455/src/bytes/bytes.go:1250
		// _ = "end of CoverTab[727]"
//line /snap/go/10455/src/bytes/bytes.go:1250
		_go_fuzz_dep_.CoverTab[728]++
							if r == tr {
//line /snap/go/10455/src/bytes/bytes.go:1251
			_go_fuzz_dep_.CoverTab[524659]++
//line /snap/go/10455/src/bytes/bytes.go:1251
			_go_fuzz_dep_.CoverTab[747]++
								continue
//line /snap/go/10455/src/bytes/bytes.go:1252
			// _ = "end of CoverTab[747]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1253
			_go_fuzz_dep_.CoverTab[524660]++
//line /snap/go/10455/src/bytes/bytes.go:1253
			_go_fuzz_dep_.CoverTab[748]++
//line /snap/go/10455/src/bytes/bytes.go:1253
			// _ = "end of CoverTab[748]"
//line /snap/go/10455/src/bytes/bytes.go:1253
		}
//line /snap/go/10455/src/bytes/bytes.go:1253
		// _ = "end of CoverTab[728]"
//line /snap/go/10455/src/bytes/bytes.go:1253
		_go_fuzz_dep_.CoverTab[729]++
							return false
//line /snap/go/10455/src/bytes/bytes.go:1254
		// _ = "end of CoverTab[729]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1255
	if _go_fuzz_dep_.CoverTab[786480] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1255
		_go_fuzz_dep_.CoverTab[524885]++
//line /snap/go/10455/src/bytes/bytes.go:1255
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1255
		_go_fuzz_dep_.CoverTab[524886]++
//line /snap/go/10455/src/bytes/bytes.go:1255
	}
//line /snap/go/10455/src/bytes/bytes.go:1255
	// _ = "end of CoverTab[703]"
//line /snap/go/10455/src/bytes/bytes.go:1255
	_go_fuzz_dep_.CoverTab[704]++

//line /snap/go/10455/src/bytes/bytes.go:1258
	return len(s) == len(t)
//line /snap/go/10455/src/bytes/bytes.go:1258
	// _ = "end of CoverTab[704]"
}

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func Index(s, sep []byte) int {
//line /snap/go/10455/src/bytes/bytes.go:1262
	_go_fuzz_dep_.CoverTab[749]++
						n := len(sep)
						switch {
	case n == 0:
//line /snap/go/10455/src/bytes/bytes.go:1265
		_go_fuzz_dep_.CoverTab[524661]++
//line /snap/go/10455/src/bytes/bytes.go:1265
		_go_fuzz_dep_.CoverTab[752]++
							return 0
//line /snap/go/10455/src/bytes/bytes.go:1266
		// _ = "end of CoverTab[752]"
	case n == 1:
//line /snap/go/10455/src/bytes/bytes.go:1267
		_go_fuzz_dep_.CoverTab[524662]++
//line /snap/go/10455/src/bytes/bytes.go:1267
		_go_fuzz_dep_.CoverTab[753]++
							return IndexByte(s, sep[0])
//line /snap/go/10455/src/bytes/bytes.go:1268
		// _ = "end of CoverTab[753]"
	case n == len(s):
//line /snap/go/10455/src/bytes/bytes.go:1269
		_go_fuzz_dep_.CoverTab[524663]++
//line /snap/go/10455/src/bytes/bytes.go:1269
		_go_fuzz_dep_.CoverTab[754]++
							if Equal(sep, s) {
//line /snap/go/10455/src/bytes/bytes.go:1270
			_go_fuzz_dep_.CoverTab[524667]++
//line /snap/go/10455/src/bytes/bytes.go:1270
			_go_fuzz_dep_.CoverTab[761]++
								return 0
//line /snap/go/10455/src/bytes/bytes.go:1271
			// _ = "end of CoverTab[761]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1272
			_go_fuzz_dep_.CoverTab[524668]++
//line /snap/go/10455/src/bytes/bytes.go:1272
			_go_fuzz_dep_.CoverTab[762]++
//line /snap/go/10455/src/bytes/bytes.go:1272
			// _ = "end of CoverTab[762]"
//line /snap/go/10455/src/bytes/bytes.go:1272
		}
//line /snap/go/10455/src/bytes/bytes.go:1272
		// _ = "end of CoverTab[754]"
//line /snap/go/10455/src/bytes/bytes.go:1272
		_go_fuzz_dep_.CoverTab[755]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:1273
		// _ = "end of CoverTab[755]"
	case n > len(s):
//line /snap/go/10455/src/bytes/bytes.go:1274
		_go_fuzz_dep_.CoverTab[524664]++
//line /snap/go/10455/src/bytes/bytes.go:1274
		_go_fuzz_dep_.CoverTab[756]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:1275
		// _ = "end of CoverTab[756]"
	case n <= bytealg.MaxLen:
//line /snap/go/10455/src/bytes/bytes.go:1276
		_go_fuzz_dep_.CoverTab[524665]++
//line /snap/go/10455/src/bytes/bytes.go:1276
		_go_fuzz_dep_.CoverTab[757]++

							if len(s) <= bytealg.MaxBruteForce {
//line /snap/go/10455/src/bytes/bytes.go:1278
			_go_fuzz_dep_.CoverTab[524669]++
//line /snap/go/10455/src/bytes/bytes.go:1278
			_go_fuzz_dep_.CoverTab[763]++
								return bytealg.Index(s, sep)
//line /snap/go/10455/src/bytes/bytes.go:1279
			// _ = "end of CoverTab[763]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1280
			_go_fuzz_dep_.CoverTab[524670]++
//line /snap/go/10455/src/bytes/bytes.go:1280
			_go_fuzz_dep_.CoverTab[764]++
//line /snap/go/10455/src/bytes/bytes.go:1280
			// _ = "end of CoverTab[764]"
//line /snap/go/10455/src/bytes/bytes.go:1280
		}
//line /snap/go/10455/src/bytes/bytes.go:1280
		// _ = "end of CoverTab[757]"
//line /snap/go/10455/src/bytes/bytes.go:1280
		_go_fuzz_dep_.CoverTab[758]++
							c0 := sep[0]
							c1 := sep[1]
							i := 0
							t := len(s) - n + 1
							fails := 0
							for i < t {
//line /snap/go/10455/src/bytes/bytes.go:1286
			_go_fuzz_dep_.CoverTab[765]++
								if s[i] != c0 {
//line /snap/go/10455/src/bytes/bytes.go:1287
				_go_fuzz_dep_.CoverTab[524671]++
//line /snap/go/10455/src/bytes/bytes.go:1287
				_go_fuzz_dep_.CoverTab[768]++

//line /snap/go/10455/src/bytes/bytes.go:1290
				o := IndexByte(s[i+1:t], c0)
				if o < 0 {
//line /snap/go/10455/src/bytes/bytes.go:1291
					_go_fuzz_dep_.CoverTab[524673]++
//line /snap/go/10455/src/bytes/bytes.go:1291
					_go_fuzz_dep_.CoverTab[770]++
										return -1
//line /snap/go/10455/src/bytes/bytes.go:1292
					// _ = "end of CoverTab[770]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:1293
					_go_fuzz_dep_.CoverTab[524674]++
//line /snap/go/10455/src/bytes/bytes.go:1293
					_go_fuzz_dep_.CoverTab[771]++
//line /snap/go/10455/src/bytes/bytes.go:1293
					// _ = "end of CoverTab[771]"
//line /snap/go/10455/src/bytes/bytes.go:1293
				}
//line /snap/go/10455/src/bytes/bytes.go:1293
				// _ = "end of CoverTab[768]"
//line /snap/go/10455/src/bytes/bytes.go:1293
				_go_fuzz_dep_.CoverTab[769]++
									i += o + 1
//line /snap/go/10455/src/bytes/bytes.go:1294
				// _ = "end of CoverTab[769]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:1295
				_go_fuzz_dep_.CoverTab[524672]++
//line /snap/go/10455/src/bytes/bytes.go:1295
				_go_fuzz_dep_.CoverTab[772]++
//line /snap/go/10455/src/bytes/bytes.go:1295
				// _ = "end of CoverTab[772]"
//line /snap/go/10455/src/bytes/bytes.go:1295
			}
//line /snap/go/10455/src/bytes/bytes.go:1295
			// _ = "end of CoverTab[765]"
//line /snap/go/10455/src/bytes/bytes.go:1295
			_go_fuzz_dep_.CoverTab[766]++
								if s[i+1] == c1 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1296
				_go_fuzz_dep_.CoverTab[773]++
//line /snap/go/10455/src/bytes/bytes.go:1296
				return Equal(s[i:i+n], sep)
//line /snap/go/10455/src/bytes/bytes.go:1296
				// _ = "end of CoverTab[773]"
//line /snap/go/10455/src/bytes/bytes.go:1296
			}() {
//line /snap/go/10455/src/bytes/bytes.go:1296
				_go_fuzz_dep_.CoverTab[524675]++
//line /snap/go/10455/src/bytes/bytes.go:1296
				_go_fuzz_dep_.CoverTab[774]++
									return i
//line /snap/go/10455/src/bytes/bytes.go:1297
				// _ = "end of CoverTab[774]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:1298
				_go_fuzz_dep_.CoverTab[524676]++
//line /snap/go/10455/src/bytes/bytes.go:1298
				_go_fuzz_dep_.CoverTab[775]++
//line /snap/go/10455/src/bytes/bytes.go:1298
				// _ = "end of CoverTab[775]"
//line /snap/go/10455/src/bytes/bytes.go:1298
			}
//line /snap/go/10455/src/bytes/bytes.go:1298
			// _ = "end of CoverTab[766]"
//line /snap/go/10455/src/bytes/bytes.go:1298
			_go_fuzz_dep_.CoverTab[767]++
								fails++
								i++

								if fails > bytealg.Cutover(i) {
//line /snap/go/10455/src/bytes/bytes.go:1302
				_go_fuzz_dep_.CoverTab[524677]++
//line /snap/go/10455/src/bytes/bytes.go:1302
				_go_fuzz_dep_.CoverTab[776]++
									r := bytealg.Index(s[i:], sep)
									if r >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:1304
					_go_fuzz_dep_.CoverTab[524679]++
//line /snap/go/10455/src/bytes/bytes.go:1304
					_go_fuzz_dep_.CoverTab[778]++
										return r + i
//line /snap/go/10455/src/bytes/bytes.go:1305
					// _ = "end of CoverTab[778]"
				} else {
//line /snap/go/10455/src/bytes/bytes.go:1306
					_go_fuzz_dep_.CoverTab[524680]++
//line /snap/go/10455/src/bytes/bytes.go:1306
					_go_fuzz_dep_.CoverTab[779]++
//line /snap/go/10455/src/bytes/bytes.go:1306
					// _ = "end of CoverTab[779]"
//line /snap/go/10455/src/bytes/bytes.go:1306
				}
//line /snap/go/10455/src/bytes/bytes.go:1306
				// _ = "end of CoverTab[776]"
//line /snap/go/10455/src/bytes/bytes.go:1306
				_go_fuzz_dep_.CoverTab[777]++
									return -1
//line /snap/go/10455/src/bytes/bytes.go:1307
				// _ = "end of CoverTab[777]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:1308
				_go_fuzz_dep_.CoverTab[524678]++
//line /snap/go/10455/src/bytes/bytes.go:1308
				_go_fuzz_dep_.CoverTab[780]++
//line /snap/go/10455/src/bytes/bytes.go:1308
				// _ = "end of CoverTab[780]"
//line /snap/go/10455/src/bytes/bytes.go:1308
			}
//line /snap/go/10455/src/bytes/bytes.go:1308
			// _ = "end of CoverTab[767]"
		}
//line /snap/go/10455/src/bytes/bytes.go:1309
		// _ = "end of CoverTab[758]"
//line /snap/go/10455/src/bytes/bytes.go:1309
		_go_fuzz_dep_.CoverTab[759]++
							return -1
//line /snap/go/10455/src/bytes/bytes.go:1310
		// _ = "end of CoverTab[759]"
//line /snap/go/10455/src/bytes/bytes.go:1310
	default:
//line /snap/go/10455/src/bytes/bytes.go:1310
		_go_fuzz_dep_.CoverTab[524666]++
//line /snap/go/10455/src/bytes/bytes.go:1310
		_go_fuzz_dep_.CoverTab[760]++
//line /snap/go/10455/src/bytes/bytes.go:1310
		// _ = "end of CoverTab[760]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1311
	// _ = "end of CoverTab[749]"
//line /snap/go/10455/src/bytes/bytes.go:1311
	_go_fuzz_dep_.CoverTab[750]++
						c0 := sep[0]
						c1 := sep[1]
						i := 0
						fails := 0
						t := len(s) - n + 1
//line /snap/go/10455/src/bytes/bytes.go:1316
	_go_fuzz_dep_.CoverTab[786482] = 0
						for i < t {
//line /snap/go/10455/src/bytes/bytes.go:1317
		if _go_fuzz_dep_.CoverTab[786482] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1317
			_go_fuzz_dep_.CoverTab[524891]++
//line /snap/go/10455/src/bytes/bytes.go:1317
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1317
			_go_fuzz_dep_.CoverTab[524892]++
//line /snap/go/10455/src/bytes/bytes.go:1317
		}
//line /snap/go/10455/src/bytes/bytes.go:1317
		_go_fuzz_dep_.CoverTab[786482] = 1
//line /snap/go/10455/src/bytes/bytes.go:1317
		_go_fuzz_dep_.CoverTab[781]++
							if s[i] != c0 {
//line /snap/go/10455/src/bytes/bytes.go:1318
			_go_fuzz_dep_.CoverTab[524681]++
//line /snap/go/10455/src/bytes/bytes.go:1318
			_go_fuzz_dep_.CoverTab[784]++
								o := IndexByte(s[i+1:t], c0)
								if o < 0 {
//line /snap/go/10455/src/bytes/bytes.go:1320
				_go_fuzz_dep_.CoverTab[524683]++
//line /snap/go/10455/src/bytes/bytes.go:1320
				_go_fuzz_dep_.CoverTab[786]++
									break
//line /snap/go/10455/src/bytes/bytes.go:1321
				// _ = "end of CoverTab[786]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:1322
				_go_fuzz_dep_.CoverTab[524684]++
//line /snap/go/10455/src/bytes/bytes.go:1322
				_go_fuzz_dep_.CoverTab[787]++
//line /snap/go/10455/src/bytes/bytes.go:1322
				// _ = "end of CoverTab[787]"
//line /snap/go/10455/src/bytes/bytes.go:1322
			}
//line /snap/go/10455/src/bytes/bytes.go:1322
			// _ = "end of CoverTab[784]"
//line /snap/go/10455/src/bytes/bytes.go:1322
			_go_fuzz_dep_.CoverTab[785]++
								i += o + 1
//line /snap/go/10455/src/bytes/bytes.go:1323
			// _ = "end of CoverTab[785]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1324
			_go_fuzz_dep_.CoverTab[524682]++
//line /snap/go/10455/src/bytes/bytes.go:1324
			_go_fuzz_dep_.CoverTab[788]++
//line /snap/go/10455/src/bytes/bytes.go:1324
			// _ = "end of CoverTab[788]"
//line /snap/go/10455/src/bytes/bytes.go:1324
		}
//line /snap/go/10455/src/bytes/bytes.go:1324
		// _ = "end of CoverTab[781]"
//line /snap/go/10455/src/bytes/bytes.go:1324
		_go_fuzz_dep_.CoverTab[782]++
							if s[i+1] == c1 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1325
			_go_fuzz_dep_.CoverTab[789]++
//line /snap/go/10455/src/bytes/bytes.go:1325
			return Equal(s[i:i+n], sep)
//line /snap/go/10455/src/bytes/bytes.go:1325
			// _ = "end of CoverTab[789]"
//line /snap/go/10455/src/bytes/bytes.go:1325
		}() {
//line /snap/go/10455/src/bytes/bytes.go:1325
			_go_fuzz_dep_.CoverTab[524685]++
//line /snap/go/10455/src/bytes/bytes.go:1325
			_go_fuzz_dep_.CoverTab[790]++
								return i
//line /snap/go/10455/src/bytes/bytes.go:1326
			// _ = "end of CoverTab[790]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1327
			_go_fuzz_dep_.CoverTab[524686]++
//line /snap/go/10455/src/bytes/bytes.go:1327
			_go_fuzz_dep_.CoverTab[791]++
//line /snap/go/10455/src/bytes/bytes.go:1327
			// _ = "end of CoverTab[791]"
//line /snap/go/10455/src/bytes/bytes.go:1327
		}
//line /snap/go/10455/src/bytes/bytes.go:1327
		// _ = "end of CoverTab[782]"
//line /snap/go/10455/src/bytes/bytes.go:1327
		_go_fuzz_dep_.CoverTab[783]++
							i++
							fails++
							if fails >= 4+i>>4 && func() bool {
//line /snap/go/10455/src/bytes/bytes.go:1330
			_go_fuzz_dep_.CoverTab[792]++
//line /snap/go/10455/src/bytes/bytes.go:1330
			return i < t
//line /snap/go/10455/src/bytes/bytes.go:1330
			// _ = "end of CoverTab[792]"
//line /snap/go/10455/src/bytes/bytes.go:1330
		}() {
//line /snap/go/10455/src/bytes/bytes.go:1330
			_go_fuzz_dep_.CoverTab[524687]++
//line /snap/go/10455/src/bytes/bytes.go:1330
			_go_fuzz_dep_.CoverTab[793]++

//line /snap/go/10455/src/bytes/bytes.go:1339
			j := bytealg.IndexRabinKarpBytes(s[i:], sep)
			if j < 0 {
//line /snap/go/10455/src/bytes/bytes.go:1340
				_go_fuzz_dep_.CoverTab[524689]++
//line /snap/go/10455/src/bytes/bytes.go:1340
				_go_fuzz_dep_.CoverTab[795]++
									return -1
//line /snap/go/10455/src/bytes/bytes.go:1341
				// _ = "end of CoverTab[795]"
			} else {
//line /snap/go/10455/src/bytes/bytes.go:1342
				_go_fuzz_dep_.CoverTab[524690]++
//line /snap/go/10455/src/bytes/bytes.go:1342
				_go_fuzz_dep_.CoverTab[796]++
//line /snap/go/10455/src/bytes/bytes.go:1342
				// _ = "end of CoverTab[796]"
//line /snap/go/10455/src/bytes/bytes.go:1342
			}
//line /snap/go/10455/src/bytes/bytes.go:1342
			// _ = "end of CoverTab[793]"
//line /snap/go/10455/src/bytes/bytes.go:1342
			_go_fuzz_dep_.CoverTab[794]++
								return i + j
//line /snap/go/10455/src/bytes/bytes.go:1343
			// _ = "end of CoverTab[794]"
		} else {
//line /snap/go/10455/src/bytes/bytes.go:1344
			_go_fuzz_dep_.CoverTab[524688]++
//line /snap/go/10455/src/bytes/bytes.go:1344
			_go_fuzz_dep_.CoverTab[797]++
//line /snap/go/10455/src/bytes/bytes.go:1344
			// _ = "end of CoverTab[797]"
//line /snap/go/10455/src/bytes/bytes.go:1344
		}
//line /snap/go/10455/src/bytes/bytes.go:1344
		// _ = "end of CoverTab[783]"
	}
//line /snap/go/10455/src/bytes/bytes.go:1345
	if _go_fuzz_dep_.CoverTab[786482] == 0 {
//line /snap/go/10455/src/bytes/bytes.go:1345
		_go_fuzz_dep_.CoverTab[524893]++
//line /snap/go/10455/src/bytes/bytes.go:1345
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1345
		_go_fuzz_dep_.CoverTab[524894]++
//line /snap/go/10455/src/bytes/bytes.go:1345
	}
//line /snap/go/10455/src/bytes/bytes.go:1345
	// _ = "end of CoverTab[750]"
//line /snap/go/10455/src/bytes/bytes.go:1345
	_go_fuzz_dep_.CoverTab[751]++
						return -1
//line /snap/go/10455/src/bytes/bytes.go:1346
	// _ = "end of CoverTab[751]"
}

// Cut slices s around the first instance of sep,
//line /snap/go/10455/src/bytes/bytes.go:1349
// returning the text before and after sep.
//line /snap/go/10455/src/bytes/bytes.go:1349
// The found result reports whether sep appears in s.
//line /snap/go/10455/src/bytes/bytes.go:1349
// If sep does not appear in s, cut returns s, nil, false.
//line /snap/go/10455/src/bytes/bytes.go:1349
//
//line /snap/go/10455/src/bytes/bytes.go:1349
// Cut returns slices of the original slice s, not copies.
//line /snap/go/10455/src/bytes/bytes.go:1355
func Cut(s, sep []byte) (before, after []byte, found bool) {
//line /snap/go/10455/src/bytes/bytes.go:1355
	_go_fuzz_dep_.CoverTab[798]++
						if i := Index(s, sep); i >= 0 {
//line /snap/go/10455/src/bytes/bytes.go:1356
		_go_fuzz_dep_.CoverTab[524691]++
//line /snap/go/10455/src/bytes/bytes.go:1356
		_go_fuzz_dep_.CoverTab[800]++
							return s[:i], s[i+len(sep):], true
//line /snap/go/10455/src/bytes/bytes.go:1357
		// _ = "end of CoverTab[800]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1358
		_go_fuzz_dep_.CoverTab[524692]++
//line /snap/go/10455/src/bytes/bytes.go:1358
		_go_fuzz_dep_.CoverTab[801]++
//line /snap/go/10455/src/bytes/bytes.go:1358
		// _ = "end of CoverTab[801]"
//line /snap/go/10455/src/bytes/bytes.go:1358
	}
//line /snap/go/10455/src/bytes/bytes.go:1358
	// _ = "end of CoverTab[798]"
//line /snap/go/10455/src/bytes/bytes.go:1358
	_go_fuzz_dep_.CoverTab[799]++
						return s, nil, false
//line /snap/go/10455/src/bytes/bytes.go:1359
	// _ = "end of CoverTab[799]"
}

// Clone returns a copy of b[:len(b)].
//line /snap/go/10455/src/bytes/bytes.go:1362
// The result may have additional unused capacity.
//line /snap/go/10455/src/bytes/bytes.go:1362
// Clone(nil) returns nil.
//line /snap/go/10455/src/bytes/bytes.go:1365
func Clone(b []byte) []byte {
//line /snap/go/10455/src/bytes/bytes.go:1365
	_go_fuzz_dep_.CoverTab[802]++
						if b == nil {
//line /snap/go/10455/src/bytes/bytes.go:1366
		_go_fuzz_dep_.CoverTab[524693]++
//line /snap/go/10455/src/bytes/bytes.go:1366
		_go_fuzz_dep_.CoverTab[804]++
							return nil
//line /snap/go/10455/src/bytes/bytes.go:1367
		// _ = "end of CoverTab[804]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1368
		_go_fuzz_dep_.CoverTab[524694]++
//line /snap/go/10455/src/bytes/bytes.go:1368
		_go_fuzz_dep_.CoverTab[805]++
//line /snap/go/10455/src/bytes/bytes.go:1368
		// _ = "end of CoverTab[805]"
//line /snap/go/10455/src/bytes/bytes.go:1368
	}
//line /snap/go/10455/src/bytes/bytes.go:1368
	// _ = "end of CoverTab[802]"
//line /snap/go/10455/src/bytes/bytes.go:1368
	_go_fuzz_dep_.CoverTab[803]++
						return append([]byte{}, b...)
//line /snap/go/10455/src/bytes/bytes.go:1369
	// _ = "end of CoverTab[803]"
}

// CutPrefix returns s without the provided leading prefix byte slice
//line /snap/go/10455/src/bytes/bytes.go:1372
// and reports whether it found the prefix.
//line /snap/go/10455/src/bytes/bytes.go:1372
// If s doesn't start with prefix, CutPrefix returns s, false.
//line /snap/go/10455/src/bytes/bytes.go:1372
// If prefix is the empty byte slice, CutPrefix returns s, true.
//line /snap/go/10455/src/bytes/bytes.go:1372
//
//line /snap/go/10455/src/bytes/bytes.go:1372
// CutPrefix returns slices of the original slice s, not copies.
//line /snap/go/10455/src/bytes/bytes.go:1378
func CutPrefix(s, prefix []byte) (after []byte, found bool) {
//line /snap/go/10455/src/bytes/bytes.go:1378
	_go_fuzz_dep_.CoverTab[806]++
						if !HasPrefix(s, prefix) {
//line /snap/go/10455/src/bytes/bytes.go:1379
		_go_fuzz_dep_.CoverTab[524695]++
//line /snap/go/10455/src/bytes/bytes.go:1379
		_go_fuzz_dep_.CoverTab[808]++
							return s, false
//line /snap/go/10455/src/bytes/bytes.go:1380
		// _ = "end of CoverTab[808]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1381
		_go_fuzz_dep_.CoverTab[524696]++
//line /snap/go/10455/src/bytes/bytes.go:1381
		_go_fuzz_dep_.CoverTab[809]++
//line /snap/go/10455/src/bytes/bytes.go:1381
		// _ = "end of CoverTab[809]"
//line /snap/go/10455/src/bytes/bytes.go:1381
	}
//line /snap/go/10455/src/bytes/bytes.go:1381
	// _ = "end of CoverTab[806]"
//line /snap/go/10455/src/bytes/bytes.go:1381
	_go_fuzz_dep_.CoverTab[807]++
						return s[len(prefix):], true
//line /snap/go/10455/src/bytes/bytes.go:1382
	// _ = "end of CoverTab[807]"
}

// CutSuffix returns s without the provided ending suffix byte slice
//line /snap/go/10455/src/bytes/bytes.go:1385
// and reports whether it found the suffix.
//line /snap/go/10455/src/bytes/bytes.go:1385
// If s doesn't end with suffix, CutSuffix returns s, false.
//line /snap/go/10455/src/bytes/bytes.go:1385
// If suffix is the empty byte slice, CutSuffix returns s, true.
//line /snap/go/10455/src/bytes/bytes.go:1385
//
//line /snap/go/10455/src/bytes/bytes.go:1385
// CutSuffix returns slices of the original slice s, not copies.
//line /snap/go/10455/src/bytes/bytes.go:1391
func CutSuffix(s, suffix []byte) (before []byte, found bool) {
//line /snap/go/10455/src/bytes/bytes.go:1391
	_go_fuzz_dep_.CoverTab[810]++
						if !HasSuffix(s, suffix) {
//line /snap/go/10455/src/bytes/bytes.go:1392
		_go_fuzz_dep_.CoverTab[524697]++
//line /snap/go/10455/src/bytes/bytes.go:1392
		_go_fuzz_dep_.CoverTab[812]++
							return s, false
//line /snap/go/10455/src/bytes/bytes.go:1393
		// _ = "end of CoverTab[812]"
	} else {
//line /snap/go/10455/src/bytes/bytes.go:1394
		_go_fuzz_dep_.CoverTab[524698]++
//line /snap/go/10455/src/bytes/bytes.go:1394
		_go_fuzz_dep_.CoverTab[813]++
//line /snap/go/10455/src/bytes/bytes.go:1394
		// _ = "end of CoverTab[813]"
//line /snap/go/10455/src/bytes/bytes.go:1394
	}
//line /snap/go/10455/src/bytes/bytes.go:1394
	// _ = "end of CoverTab[810]"
//line /snap/go/10455/src/bytes/bytes.go:1394
	_go_fuzz_dep_.CoverTab[811]++
						return s[:len(s)-len(suffix)], true
//line /snap/go/10455/src/bytes/bytes.go:1395
	// _ = "end of CoverTab[811]"
}

//line /snap/go/10455/src/bytes/bytes.go:1396
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/bytes/bytes.go:1396
var _ = _go_fuzz_dep_.CoverTab
