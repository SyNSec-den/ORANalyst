// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/strings/search.go:5
package strings

//line /usr/local/go/src/strings/search.go:5
import (
//line /usr/local/go/src/strings/search.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/strings/search.go:5
)
//line /usr/local/go/src/strings/search.go:5
import (
//line /usr/local/go/src/strings/search.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/strings/search.go:5
)

// stringFinder efficiently finds strings in a source text. It's implemented
//line /usr/local/go/src/strings/search.go:7
// using the Boyer-Moore string search algorithm:
//line /usr/local/go/src/strings/search.go:7
// https://en.wikipedia.org/wiki/Boyer-Moore_string_search_algorithm
//line /usr/local/go/src/strings/search.go:7
// https://www.cs.utexas.edu/~moore/publications/fstrpos.pdf (note: this aged
//line /usr/local/go/src/strings/search.go:7
// document uses 1-based indexing)
//line /usr/local/go/src/strings/search.go:12
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
//line /usr/local/go/src/strings/search.go:48
	_go_fuzz_dep_.CoverTab[1162]++
						f := &stringFinder{
		pattern:	pattern,
		goodSuffixSkip:	make([]int, len(pattern)),
	}

						last := len(pattern) - 1

//line /usr/local/go/src/strings/search.go:58
	for i := range f.badCharSkip {
//line /usr/local/go/src/strings/search.go:58
		_go_fuzz_dep_.CoverTab[1167]++
							f.badCharSkip[i] = len(pattern)
//line /usr/local/go/src/strings/search.go:59
		// _ = "end of CoverTab[1167]"
	}
//line /usr/local/go/src/strings/search.go:60
	// _ = "end of CoverTab[1162]"
//line /usr/local/go/src/strings/search.go:60
	_go_fuzz_dep_.CoverTab[1163]++

//line /usr/local/go/src/strings/search.go:64
	for i := 0; i < last; i++ {
//line /usr/local/go/src/strings/search.go:64
		_go_fuzz_dep_.CoverTab[1168]++
							f.badCharSkip[pattern[i]] = last - i
//line /usr/local/go/src/strings/search.go:65
		// _ = "end of CoverTab[1168]"
	}
//line /usr/local/go/src/strings/search.go:66
	// _ = "end of CoverTab[1163]"
//line /usr/local/go/src/strings/search.go:66
	_go_fuzz_dep_.CoverTab[1164]++

//line /usr/local/go/src/strings/search.go:71
	lastPrefix := last
	for i := last; i >= 0; i-- {
//line /usr/local/go/src/strings/search.go:72
		_go_fuzz_dep_.CoverTab[1169]++
							if HasPrefix(pattern, pattern[i+1:]) {
//line /usr/local/go/src/strings/search.go:73
			_go_fuzz_dep_.CoverTab[1171]++
								lastPrefix = i + 1
//line /usr/local/go/src/strings/search.go:74
			// _ = "end of CoverTab[1171]"
		} else {
//line /usr/local/go/src/strings/search.go:75
			_go_fuzz_dep_.CoverTab[1172]++
//line /usr/local/go/src/strings/search.go:75
			// _ = "end of CoverTab[1172]"
//line /usr/local/go/src/strings/search.go:75
		}
//line /usr/local/go/src/strings/search.go:75
		// _ = "end of CoverTab[1169]"
//line /usr/local/go/src/strings/search.go:75
		_go_fuzz_dep_.CoverTab[1170]++

							f.goodSuffixSkip[i] = lastPrefix + last - i
//line /usr/local/go/src/strings/search.go:77
		// _ = "end of CoverTab[1170]"
	}
//line /usr/local/go/src/strings/search.go:78
	// _ = "end of CoverTab[1164]"
//line /usr/local/go/src/strings/search.go:78
	_go_fuzz_dep_.CoverTab[1165]++

						for i := 0; i < last; i++ {
//line /usr/local/go/src/strings/search.go:80
		_go_fuzz_dep_.CoverTab[1173]++
							lenSuffix := longestCommonSuffix(pattern, pattern[1:i+1])
							if pattern[i-lenSuffix] != pattern[last-lenSuffix] {
//line /usr/local/go/src/strings/search.go:82
			_go_fuzz_dep_.CoverTab[1174]++

								f.goodSuffixSkip[last-lenSuffix] = lenSuffix + last - i
//line /usr/local/go/src/strings/search.go:84
			// _ = "end of CoverTab[1174]"
		} else {
//line /usr/local/go/src/strings/search.go:85
			_go_fuzz_dep_.CoverTab[1175]++
//line /usr/local/go/src/strings/search.go:85
			// _ = "end of CoverTab[1175]"
//line /usr/local/go/src/strings/search.go:85
		}
//line /usr/local/go/src/strings/search.go:85
		// _ = "end of CoverTab[1173]"
	}
//line /usr/local/go/src/strings/search.go:86
	// _ = "end of CoverTab[1165]"
//line /usr/local/go/src/strings/search.go:86
	_go_fuzz_dep_.CoverTab[1166]++

						return f
//line /usr/local/go/src/strings/search.go:88
	// _ = "end of CoverTab[1166]"
}

func longestCommonSuffix(a, b string) (i int) {
//line /usr/local/go/src/strings/search.go:91
	_go_fuzz_dep_.CoverTab[1176]++
						for ; i < len(a) && func() bool {
//line /usr/local/go/src/strings/search.go:92
		_go_fuzz_dep_.CoverTab[1178]++
//line /usr/local/go/src/strings/search.go:92
		return i < len(b)
//line /usr/local/go/src/strings/search.go:92
		// _ = "end of CoverTab[1178]"
//line /usr/local/go/src/strings/search.go:92
	}(); i++ {
//line /usr/local/go/src/strings/search.go:92
		_go_fuzz_dep_.CoverTab[1179]++
							if a[len(a)-1-i] != b[len(b)-1-i] {
//line /usr/local/go/src/strings/search.go:93
			_go_fuzz_dep_.CoverTab[1180]++
								break
//line /usr/local/go/src/strings/search.go:94
			// _ = "end of CoverTab[1180]"
		} else {
//line /usr/local/go/src/strings/search.go:95
			_go_fuzz_dep_.CoverTab[1181]++
//line /usr/local/go/src/strings/search.go:95
			// _ = "end of CoverTab[1181]"
//line /usr/local/go/src/strings/search.go:95
		}
//line /usr/local/go/src/strings/search.go:95
		// _ = "end of CoverTab[1179]"
	}
//line /usr/local/go/src/strings/search.go:96
	// _ = "end of CoverTab[1176]"
//line /usr/local/go/src/strings/search.go:96
	_go_fuzz_dep_.CoverTab[1177]++
						return
//line /usr/local/go/src/strings/search.go:97
	// _ = "end of CoverTab[1177]"
}

// next returns the index in text of the first occurrence of the pattern. If
//line /usr/local/go/src/strings/search.go:100
// the pattern is not found, it returns -1.
//line /usr/local/go/src/strings/search.go:102
func (f *stringFinder) next(text string) int {
//line /usr/local/go/src/strings/search.go:102
	_go_fuzz_dep_.CoverTab[1182]++
						i := len(f.pattern) - 1
						for i < len(text) {
//line /usr/local/go/src/strings/search.go:104
		_go_fuzz_dep_.CoverTab[1184]++

							j := len(f.pattern) - 1
							for j >= 0 && func() bool {
//line /usr/local/go/src/strings/search.go:107
			_go_fuzz_dep_.CoverTab[1187]++
//line /usr/local/go/src/strings/search.go:107
			return text[i] == f.pattern[j]
//line /usr/local/go/src/strings/search.go:107
			// _ = "end of CoverTab[1187]"
//line /usr/local/go/src/strings/search.go:107
		}() {
//line /usr/local/go/src/strings/search.go:107
			_go_fuzz_dep_.CoverTab[1188]++
								i--
								j--
//line /usr/local/go/src/strings/search.go:109
			// _ = "end of CoverTab[1188]"
		}
//line /usr/local/go/src/strings/search.go:110
		// _ = "end of CoverTab[1184]"
//line /usr/local/go/src/strings/search.go:110
		_go_fuzz_dep_.CoverTab[1185]++
							if j < 0 {
//line /usr/local/go/src/strings/search.go:111
			_go_fuzz_dep_.CoverTab[1189]++
								return i + 1
//line /usr/local/go/src/strings/search.go:112
			// _ = "end of CoverTab[1189]"
		} else {
//line /usr/local/go/src/strings/search.go:113
			_go_fuzz_dep_.CoverTab[1190]++
//line /usr/local/go/src/strings/search.go:113
			// _ = "end of CoverTab[1190]"
//line /usr/local/go/src/strings/search.go:113
		}
//line /usr/local/go/src/strings/search.go:113
		// _ = "end of CoverTab[1185]"
//line /usr/local/go/src/strings/search.go:113
		_go_fuzz_dep_.CoverTab[1186]++
							i += max(f.badCharSkip[text[i]], f.goodSuffixSkip[j])
//line /usr/local/go/src/strings/search.go:114
		// _ = "end of CoverTab[1186]"
	}
//line /usr/local/go/src/strings/search.go:115
	// _ = "end of CoverTab[1182]"
//line /usr/local/go/src/strings/search.go:115
	_go_fuzz_dep_.CoverTab[1183]++
						return -1
//line /usr/local/go/src/strings/search.go:116
	// _ = "end of CoverTab[1183]"
}

func max(a, b int) int {
//line /usr/local/go/src/strings/search.go:119
	_go_fuzz_dep_.CoverTab[1191]++
						if a > b {
//line /usr/local/go/src/strings/search.go:120
		_go_fuzz_dep_.CoverTab[1193]++
							return a
//line /usr/local/go/src/strings/search.go:121
		// _ = "end of CoverTab[1193]"
	} else {
//line /usr/local/go/src/strings/search.go:122
		_go_fuzz_dep_.CoverTab[1194]++
//line /usr/local/go/src/strings/search.go:122
		// _ = "end of CoverTab[1194]"
//line /usr/local/go/src/strings/search.go:122
	}
//line /usr/local/go/src/strings/search.go:122
	// _ = "end of CoverTab[1191]"
//line /usr/local/go/src/strings/search.go:122
	_go_fuzz_dep_.CoverTab[1192]++
						return b
//line /usr/local/go/src/strings/search.go:123
	// _ = "end of CoverTab[1192]"
}

//line /usr/local/go/src/strings/search.go:124
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/search.go:124
var _ = _go_fuzz_dep_.CoverTab
