// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/strings/search.go:5
package strings

//line /snap/go/10455/src/strings/search.go:5
import (
//line /snap/go/10455/src/strings/search.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/strings/search.go:5
)
//line /snap/go/10455/src/strings/search.go:5
import (
//line /snap/go/10455/src/strings/search.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/strings/search.go:5
)

// stringFinder efficiently finds strings in a source text. It's implemented
//line /snap/go/10455/src/strings/search.go:7
// using the Boyer-Moore string search algorithm:
//line /snap/go/10455/src/strings/search.go:7
// https://en.wikipedia.org/wiki/Boyer-Moore_string_search_algorithm
//line /snap/go/10455/src/strings/search.go:7
// https://www.cs.utexas.edu/~moore/publications/fstrpos.pdf (note: this aged
//line /snap/go/10455/src/strings/search.go:7
// document uses 1-based indexing)
//line /snap/go/10455/src/strings/search.go:12
type stringFinder struct {
	// pattern is the string that we are searching for in the text.
	pattern	string

	// badCharSkip[b] contains the distance between the last byte of pattern
	// and the rightmost occurrence of b in pattern. If b is not in pattern,
	// badCharSkip[b] is len(pattern).
	//
	// Whenever a mismatch is found with byte b in the text, we can safely
	// shift the matching frame at least badCharSkip[b] until the next time
	// the matching char could be in alignment.
	badCharSkip	[256]int

	// goodSuffixSkip[i] defines how far we can shift the matching frame given
	// that the suffix pattern[i+1:] matches, but the byte pattern[i] does
	// not. There are two cases to consider:
	//
	// 1. The matched suffix occurs elsewhere in pattern (with a different
	// byte preceding it that we might possibly match). In this case, we can
	// shift the matching frame to align with the next suffix chunk. For
	// example, the pattern "mississi" has the suffix "issi" next occurring
	// (in right-to-left order) at index 1, so goodSuffixSkip[3] ==
	// shift+len(suffix) == 3+4 == 7.
	//
	// 2. If the matched suffix does not occur elsewhere in pattern, then the
	// matching frame may share part of its prefix with the end of the
	// matching suffix. In this case, goodSuffixSkip[i] will contain how far
	// to shift the frame to align this portion of the prefix to the
	// suffix. For example, in the pattern "abcxxxabc", when the first
	// mismatch from the back is found to be in position 3, the matching
	// suffix "xxabc" is not found elsewhere in the pattern. However, its
	// rightmost "abc" (at position 6) is a prefix of the whole pattern, so
	// goodSuffixSkip[3] == shift+len(suffix) == 6+5 == 11.
	goodSuffixSkip	[]int
}

func makeStringFinder(pattern string) *stringFinder {
//line /snap/go/10455/src/strings/search.go:48
	_go_fuzz_dep_.CoverTab[1175]++
						f := &stringFinder{
		pattern:	pattern,
		goodSuffixSkip:	make([]int, len(pattern)),
	}

						last := len(pattern) - 1
//line /snap/go/10455/src/strings/search.go:54
	_go_fuzz_dep_.CoverTab[786503] = 0

//line /snap/go/10455/src/strings/search.go:58
	for i := range f.badCharSkip {
//line /snap/go/10455/src/strings/search.go:58
		if _go_fuzz_dep_.CoverTab[786503] == 0 {
//line /snap/go/10455/src/strings/search.go:58
			_go_fuzz_dep_.CoverTab[525171]++
//line /snap/go/10455/src/strings/search.go:58
		} else {
//line /snap/go/10455/src/strings/search.go:58
			_go_fuzz_dep_.CoverTab[525172]++
//line /snap/go/10455/src/strings/search.go:58
		}
//line /snap/go/10455/src/strings/search.go:58
		_go_fuzz_dep_.CoverTab[786503] = 1
//line /snap/go/10455/src/strings/search.go:58
		_go_fuzz_dep_.CoverTab[1180]++
							f.badCharSkip[i] = len(pattern)
//line /snap/go/10455/src/strings/search.go:59
		// _ = "end of CoverTab[1180]"
	}
//line /snap/go/10455/src/strings/search.go:60
	if _go_fuzz_dep_.CoverTab[786503] == 0 {
//line /snap/go/10455/src/strings/search.go:60
		_go_fuzz_dep_.CoverTab[525173]++
//line /snap/go/10455/src/strings/search.go:60
	} else {
//line /snap/go/10455/src/strings/search.go:60
		_go_fuzz_dep_.CoverTab[525174]++
//line /snap/go/10455/src/strings/search.go:60
	}
//line /snap/go/10455/src/strings/search.go:60
	// _ = "end of CoverTab[1175]"
//line /snap/go/10455/src/strings/search.go:60
	_go_fuzz_dep_.CoverTab[1176]++
//line /snap/go/10455/src/strings/search.go:60
	_go_fuzz_dep_.CoverTab[786504] = 0

//line /snap/go/10455/src/strings/search.go:64
	for i := 0; i < last; i++ {
//line /snap/go/10455/src/strings/search.go:64
		if _go_fuzz_dep_.CoverTab[786504] == 0 {
//line /snap/go/10455/src/strings/search.go:64
			_go_fuzz_dep_.CoverTab[525175]++
//line /snap/go/10455/src/strings/search.go:64
		} else {
//line /snap/go/10455/src/strings/search.go:64
			_go_fuzz_dep_.CoverTab[525176]++
//line /snap/go/10455/src/strings/search.go:64
		}
//line /snap/go/10455/src/strings/search.go:64
		_go_fuzz_dep_.CoverTab[786504] = 1
//line /snap/go/10455/src/strings/search.go:64
		_go_fuzz_dep_.CoverTab[1181]++
							f.badCharSkip[pattern[i]] = last - i
//line /snap/go/10455/src/strings/search.go:65
		// _ = "end of CoverTab[1181]"
	}
//line /snap/go/10455/src/strings/search.go:66
	if _go_fuzz_dep_.CoverTab[786504] == 0 {
//line /snap/go/10455/src/strings/search.go:66
		_go_fuzz_dep_.CoverTab[525177]++
//line /snap/go/10455/src/strings/search.go:66
	} else {
//line /snap/go/10455/src/strings/search.go:66
		_go_fuzz_dep_.CoverTab[525178]++
//line /snap/go/10455/src/strings/search.go:66
	}
//line /snap/go/10455/src/strings/search.go:66
	// _ = "end of CoverTab[1176]"
//line /snap/go/10455/src/strings/search.go:66
	_go_fuzz_dep_.CoverTab[1177]++

//line /snap/go/10455/src/strings/search.go:71
	lastPrefix := last
//line /snap/go/10455/src/strings/search.go:71
	_go_fuzz_dep_.CoverTab[786505] = 0
						for i := last; i >= 0; i-- {
//line /snap/go/10455/src/strings/search.go:72
		if _go_fuzz_dep_.CoverTab[786505] == 0 {
//line /snap/go/10455/src/strings/search.go:72
			_go_fuzz_dep_.CoverTab[525179]++
//line /snap/go/10455/src/strings/search.go:72
		} else {
//line /snap/go/10455/src/strings/search.go:72
			_go_fuzz_dep_.CoverTab[525180]++
//line /snap/go/10455/src/strings/search.go:72
		}
//line /snap/go/10455/src/strings/search.go:72
		_go_fuzz_dep_.CoverTab[786505] = 1
//line /snap/go/10455/src/strings/search.go:72
		_go_fuzz_dep_.CoverTab[1182]++
							if HasPrefix(pattern, pattern[i+1:]) {
//line /snap/go/10455/src/strings/search.go:73
			_go_fuzz_dep_.CoverTab[525161]++
//line /snap/go/10455/src/strings/search.go:73
			_go_fuzz_dep_.CoverTab[1184]++
								lastPrefix = i + 1
//line /snap/go/10455/src/strings/search.go:74
			// _ = "end of CoverTab[1184]"
		} else {
//line /snap/go/10455/src/strings/search.go:75
			_go_fuzz_dep_.CoverTab[525162]++
//line /snap/go/10455/src/strings/search.go:75
			_go_fuzz_dep_.CoverTab[1185]++
//line /snap/go/10455/src/strings/search.go:75
			// _ = "end of CoverTab[1185]"
//line /snap/go/10455/src/strings/search.go:75
		}
//line /snap/go/10455/src/strings/search.go:75
		// _ = "end of CoverTab[1182]"
//line /snap/go/10455/src/strings/search.go:75
		_go_fuzz_dep_.CoverTab[1183]++

							f.goodSuffixSkip[i] = lastPrefix + last - i
//line /snap/go/10455/src/strings/search.go:77
		// _ = "end of CoverTab[1183]"
	}
//line /snap/go/10455/src/strings/search.go:78
	if _go_fuzz_dep_.CoverTab[786505] == 0 {
//line /snap/go/10455/src/strings/search.go:78
		_go_fuzz_dep_.CoverTab[525181]++
//line /snap/go/10455/src/strings/search.go:78
	} else {
//line /snap/go/10455/src/strings/search.go:78
		_go_fuzz_dep_.CoverTab[525182]++
//line /snap/go/10455/src/strings/search.go:78
	}
//line /snap/go/10455/src/strings/search.go:78
	// _ = "end of CoverTab[1177]"
//line /snap/go/10455/src/strings/search.go:78
	_go_fuzz_dep_.CoverTab[1178]++
//line /snap/go/10455/src/strings/search.go:78
	_go_fuzz_dep_.CoverTab[786506] = 0

						for i := 0; i < last; i++ {
//line /snap/go/10455/src/strings/search.go:80
		if _go_fuzz_dep_.CoverTab[786506] == 0 {
//line /snap/go/10455/src/strings/search.go:80
			_go_fuzz_dep_.CoverTab[525183]++
//line /snap/go/10455/src/strings/search.go:80
		} else {
//line /snap/go/10455/src/strings/search.go:80
			_go_fuzz_dep_.CoverTab[525184]++
//line /snap/go/10455/src/strings/search.go:80
		}
//line /snap/go/10455/src/strings/search.go:80
		_go_fuzz_dep_.CoverTab[786506] = 1
//line /snap/go/10455/src/strings/search.go:80
		_go_fuzz_dep_.CoverTab[1186]++
							lenSuffix := longestCommonSuffix(pattern, pattern[1:i+1])
							if pattern[i-lenSuffix] != pattern[last-lenSuffix] {
//line /snap/go/10455/src/strings/search.go:82
			_go_fuzz_dep_.CoverTab[525163]++
//line /snap/go/10455/src/strings/search.go:82
			_go_fuzz_dep_.CoverTab[1187]++

								f.goodSuffixSkip[last-lenSuffix] = lenSuffix + last - i
//line /snap/go/10455/src/strings/search.go:84
			// _ = "end of CoverTab[1187]"
		} else {
//line /snap/go/10455/src/strings/search.go:85
			_go_fuzz_dep_.CoverTab[525164]++
//line /snap/go/10455/src/strings/search.go:85
			_go_fuzz_dep_.CoverTab[1188]++
//line /snap/go/10455/src/strings/search.go:85
			// _ = "end of CoverTab[1188]"
//line /snap/go/10455/src/strings/search.go:85
		}
//line /snap/go/10455/src/strings/search.go:85
		// _ = "end of CoverTab[1186]"
	}
//line /snap/go/10455/src/strings/search.go:86
	if _go_fuzz_dep_.CoverTab[786506] == 0 {
//line /snap/go/10455/src/strings/search.go:86
		_go_fuzz_dep_.CoverTab[525185]++
//line /snap/go/10455/src/strings/search.go:86
	} else {
//line /snap/go/10455/src/strings/search.go:86
		_go_fuzz_dep_.CoverTab[525186]++
//line /snap/go/10455/src/strings/search.go:86
	}
//line /snap/go/10455/src/strings/search.go:86
	// _ = "end of CoverTab[1178]"
//line /snap/go/10455/src/strings/search.go:86
	_go_fuzz_dep_.CoverTab[1179]++

						return f
//line /snap/go/10455/src/strings/search.go:88
	// _ = "end of CoverTab[1179]"
}

func longestCommonSuffix(a, b string) (i int) {
//line /snap/go/10455/src/strings/search.go:91
	_go_fuzz_dep_.CoverTab[1189]++
//line /snap/go/10455/src/strings/search.go:91
	_go_fuzz_dep_.CoverTab[786507] = 0
						for ; i < len(a) && func() bool {
//line /snap/go/10455/src/strings/search.go:92
		_go_fuzz_dep_.CoverTab[1191]++
//line /snap/go/10455/src/strings/search.go:92
		return i < len(b)
//line /snap/go/10455/src/strings/search.go:92
		// _ = "end of CoverTab[1191]"
//line /snap/go/10455/src/strings/search.go:92
	}(); i++ {
//line /snap/go/10455/src/strings/search.go:92
		if _go_fuzz_dep_.CoverTab[786507] == 0 {
//line /snap/go/10455/src/strings/search.go:92
			_go_fuzz_dep_.CoverTab[525187]++
//line /snap/go/10455/src/strings/search.go:92
		} else {
//line /snap/go/10455/src/strings/search.go:92
			_go_fuzz_dep_.CoverTab[525188]++
//line /snap/go/10455/src/strings/search.go:92
		}
//line /snap/go/10455/src/strings/search.go:92
		_go_fuzz_dep_.CoverTab[786507] = 1
//line /snap/go/10455/src/strings/search.go:92
		_go_fuzz_dep_.CoverTab[1192]++
							if a[len(a)-1-i] != b[len(b)-1-i] {
//line /snap/go/10455/src/strings/search.go:93
			_go_fuzz_dep_.CoverTab[525165]++
//line /snap/go/10455/src/strings/search.go:93
			_go_fuzz_dep_.CoverTab[1193]++
								break
//line /snap/go/10455/src/strings/search.go:94
			// _ = "end of CoverTab[1193]"
		} else {
//line /snap/go/10455/src/strings/search.go:95
			_go_fuzz_dep_.CoverTab[525166]++
//line /snap/go/10455/src/strings/search.go:95
			_go_fuzz_dep_.CoverTab[1194]++
//line /snap/go/10455/src/strings/search.go:95
			// _ = "end of CoverTab[1194]"
//line /snap/go/10455/src/strings/search.go:95
		}
//line /snap/go/10455/src/strings/search.go:95
		// _ = "end of CoverTab[1192]"
	}
//line /snap/go/10455/src/strings/search.go:96
	if _go_fuzz_dep_.CoverTab[786507] == 0 {
//line /snap/go/10455/src/strings/search.go:96
		_go_fuzz_dep_.CoverTab[525189]++
//line /snap/go/10455/src/strings/search.go:96
	} else {
//line /snap/go/10455/src/strings/search.go:96
		_go_fuzz_dep_.CoverTab[525190]++
//line /snap/go/10455/src/strings/search.go:96
	}
//line /snap/go/10455/src/strings/search.go:96
	// _ = "end of CoverTab[1189]"
//line /snap/go/10455/src/strings/search.go:96
	_go_fuzz_dep_.CoverTab[1190]++
						return
//line /snap/go/10455/src/strings/search.go:97
	// _ = "end of CoverTab[1190]"
}

// next returns the index in text of the first occurrence of the pattern. If
//line /snap/go/10455/src/strings/search.go:100
// the pattern is not found, it returns -1.
//line /snap/go/10455/src/strings/search.go:102
func (f *stringFinder) next(text string) int {
//line /snap/go/10455/src/strings/search.go:102
	_go_fuzz_dep_.CoverTab[1195]++
							i := len(f.pattern) - 1
//line /snap/go/10455/src/strings/search.go:103
	_go_fuzz_dep_.CoverTab[786508] = 0
							for i < len(text) {
//line /snap/go/10455/src/strings/search.go:104
		if _go_fuzz_dep_.CoverTab[786508] == 0 {
//line /snap/go/10455/src/strings/search.go:104
			_go_fuzz_dep_.CoverTab[525191]++
//line /snap/go/10455/src/strings/search.go:104
		} else {
//line /snap/go/10455/src/strings/search.go:104
			_go_fuzz_dep_.CoverTab[525192]++
//line /snap/go/10455/src/strings/search.go:104
		}
//line /snap/go/10455/src/strings/search.go:104
		_go_fuzz_dep_.CoverTab[786508] = 1
//line /snap/go/10455/src/strings/search.go:104
		_go_fuzz_dep_.CoverTab[1197]++

								j := len(f.pattern) - 1
//line /snap/go/10455/src/strings/search.go:106
		_go_fuzz_dep_.CoverTab[786509] = 0
								for j >= 0 && func() bool {
//line /snap/go/10455/src/strings/search.go:107
			_go_fuzz_dep_.CoverTab[1200]++
//line /snap/go/10455/src/strings/search.go:107
			return text[i] == f.pattern[j]
//line /snap/go/10455/src/strings/search.go:107
			// _ = "end of CoverTab[1200]"
//line /snap/go/10455/src/strings/search.go:107
		}() {
//line /snap/go/10455/src/strings/search.go:107
			if _go_fuzz_dep_.CoverTab[786509] == 0 {
//line /snap/go/10455/src/strings/search.go:107
				_go_fuzz_dep_.CoverTab[525195]++
//line /snap/go/10455/src/strings/search.go:107
			} else {
//line /snap/go/10455/src/strings/search.go:107
				_go_fuzz_dep_.CoverTab[525196]++
//line /snap/go/10455/src/strings/search.go:107
			}
//line /snap/go/10455/src/strings/search.go:107
			_go_fuzz_dep_.CoverTab[786509] = 1
//line /snap/go/10455/src/strings/search.go:107
			_go_fuzz_dep_.CoverTab[1201]++
									i--
									j--
//line /snap/go/10455/src/strings/search.go:109
			// _ = "end of CoverTab[1201]"
		}
//line /snap/go/10455/src/strings/search.go:110
		if _go_fuzz_dep_.CoverTab[786509] == 0 {
//line /snap/go/10455/src/strings/search.go:110
			_go_fuzz_dep_.CoverTab[525197]++
//line /snap/go/10455/src/strings/search.go:110
		} else {
//line /snap/go/10455/src/strings/search.go:110
			_go_fuzz_dep_.CoverTab[525198]++
//line /snap/go/10455/src/strings/search.go:110
		}
//line /snap/go/10455/src/strings/search.go:110
		// _ = "end of CoverTab[1197]"
//line /snap/go/10455/src/strings/search.go:110
		_go_fuzz_dep_.CoverTab[1198]++
								if j < 0 {
//line /snap/go/10455/src/strings/search.go:111
			_go_fuzz_dep_.CoverTab[525167]++
//line /snap/go/10455/src/strings/search.go:111
			_go_fuzz_dep_.CoverTab[1202]++
									return i + 1
//line /snap/go/10455/src/strings/search.go:112
			// _ = "end of CoverTab[1202]"
		} else {
//line /snap/go/10455/src/strings/search.go:113
			_go_fuzz_dep_.CoverTab[525168]++
//line /snap/go/10455/src/strings/search.go:113
			_go_fuzz_dep_.CoverTab[1203]++
//line /snap/go/10455/src/strings/search.go:113
			// _ = "end of CoverTab[1203]"
//line /snap/go/10455/src/strings/search.go:113
		}
//line /snap/go/10455/src/strings/search.go:113
		// _ = "end of CoverTab[1198]"
//line /snap/go/10455/src/strings/search.go:113
		_go_fuzz_dep_.CoverTab[1199]++
								i += max(f.badCharSkip[text[i]], f.goodSuffixSkip[j])
//line /snap/go/10455/src/strings/search.go:114
		// _ = "end of CoverTab[1199]"
	}
//line /snap/go/10455/src/strings/search.go:115
	if _go_fuzz_dep_.CoverTab[786508] == 0 {
//line /snap/go/10455/src/strings/search.go:115
		_go_fuzz_dep_.CoverTab[525193]++
//line /snap/go/10455/src/strings/search.go:115
	} else {
//line /snap/go/10455/src/strings/search.go:115
		_go_fuzz_dep_.CoverTab[525194]++
//line /snap/go/10455/src/strings/search.go:115
	}
//line /snap/go/10455/src/strings/search.go:115
	// _ = "end of CoverTab[1195]"
//line /snap/go/10455/src/strings/search.go:115
	_go_fuzz_dep_.CoverTab[1196]++
							return -1
//line /snap/go/10455/src/strings/search.go:116
	// _ = "end of CoverTab[1196]"
}

func max(a, b int) int {
//line /snap/go/10455/src/strings/search.go:119
	_go_fuzz_dep_.CoverTab[1204]++
							if a > b {
//line /snap/go/10455/src/strings/search.go:120
		_go_fuzz_dep_.CoverTab[525169]++
//line /snap/go/10455/src/strings/search.go:120
		_go_fuzz_dep_.CoverTab[1206]++
								return a
//line /snap/go/10455/src/strings/search.go:121
		// _ = "end of CoverTab[1206]"
	} else {
//line /snap/go/10455/src/strings/search.go:122
		_go_fuzz_dep_.CoverTab[525170]++
//line /snap/go/10455/src/strings/search.go:122
		_go_fuzz_dep_.CoverTab[1207]++
//line /snap/go/10455/src/strings/search.go:122
		// _ = "end of CoverTab[1207]"
//line /snap/go/10455/src/strings/search.go:122
	}
//line /snap/go/10455/src/strings/search.go:122
	// _ = "end of CoverTab[1204]"
//line /snap/go/10455/src/strings/search.go:122
	_go_fuzz_dep_.CoverTab[1205]++
							return b
//line /snap/go/10455/src/strings/search.go:123
	// _ = "end of CoverTab[1205]"
}

//line /snap/go/10455/src/strings/search.go:124
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/strings/search.go:124
var _ = _go_fuzz_dep_.CoverTab
