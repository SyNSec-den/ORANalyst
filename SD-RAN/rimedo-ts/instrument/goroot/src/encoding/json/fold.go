// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/json/fold.go:5
package json

//line /usr/local/go/src/encoding/json/fold.go:5
import (
//line /usr/local/go/src/encoding/json/fold.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/json/fold.go:5
)
//line /usr/local/go/src/encoding/json/fold.go:5
import (
//line /usr/local/go/src/encoding/json/fold.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/json/fold.go:5
)

import (
	"bytes"
	"unicode/utf8"
)

const (
	caseMask	= ^byte(0x20)	// Mask to ignore case in ASCII.
	kelvin		= '\u212a'
	smallLongEss	= '\u017f'
)

// foldFunc returns one of four different case folding equivalence
//line /usr/local/go/src/encoding/json/fold.go:18
// functions, from most general (and slow) to fastest:
//line /usr/local/go/src/encoding/json/fold.go:18
//
//line /usr/local/go/src/encoding/json/fold.go:18
// 1) bytes.EqualFold, if the key s contains any non-ASCII UTF-8
//line /usr/local/go/src/encoding/json/fold.go:18
// 2) equalFoldRight, if s contains special folding ASCII ('k', 'K', 's', 'S')
//line /usr/local/go/src/encoding/json/fold.go:18
// 3) asciiEqualFold, no special, but includes non-letters (including _)
//line /usr/local/go/src/encoding/json/fold.go:18
// 4) simpleLetterEqualFold, no specials, no non-letters.
//line /usr/local/go/src/encoding/json/fold.go:18
//
//line /usr/local/go/src/encoding/json/fold.go:18
// The letters S and K are special because they map to 3 runes, not just 2:
//line /usr/local/go/src/encoding/json/fold.go:18
//   - S maps to s and to U+017F 'ſ' Latin small letter long s
//line /usr/local/go/src/encoding/json/fold.go:18
//   - k maps to K and to U+212A 'K' Kelvin sign
//line /usr/local/go/src/encoding/json/fold.go:18
//
//line /usr/local/go/src/encoding/json/fold.go:18
// See https://play.golang.org/p/tTxjOc0OGo
//line /usr/local/go/src/encoding/json/fold.go:18
//
//line /usr/local/go/src/encoding/json/fold.go:18
// The returned function is specialized for matching against s and
//line /usr/local/go/src/encoding/json/fold.go:18
// should only be given s. It's not curried for performance reasons.
//line /usr/local/go/src/encoding/json/fold.go:34
func foldFunc(s []byte) func(s, t []byte) bool {
//line /usr/local/go/src/encoding/json/fold.go:34
	_go_fuzz_dep_.CoverTab[27968]++
							nonLetter := false
							special := false
							for _, b := range s {
//line /usr/local/go/src/encoding/json/fold.go:37
		_go_fuzz_dep_.CoverTab[27972]++
								if b >= utf8.RuneSelf {
//line /usr/local/go/src/encoding/json/fold.go:38
			_go_fuzz_dep_.CoverTab[27974]++
									return bytes.EqualFold
//line /usr/local/go/src/encoding/json/fold.go:39
			// _ = "end of CoverTab[27974]"
		} else {
//line /usr/local/go/src/encoding/json/fold.go:40
			_go_fuzz_dep_.CoverTab[27975]++
//line /usr/local/go/src/encoding/json/fold.go:40
			// _ = "end of CoverTab[27975]"
//line /usr/local/go/src/encoding/json/fold.go:40
		}
//line /usr/local/go/src/encoding/json/fold.go:40
		// _ = "end of CoverTab[27972]"
//line /usr/local/go/src/encoding/json/fold.go:40
		_go_fuzz_dep_.CoverTab[27973]++
								upper := b & caseMask
								if upper < 'A' || func() bool {
//line /usr/local/go/src/encoding/json/fold.go:42
			_go_fuzz_dep_.CoverTab[27976]++
//line /usr/local/go/src/encoding/json/fold.go:42
			return upper > 'Z'
//line /usr/local/go/src/encoding/json/fold.go:42
			// _ = "end of CoverTab[27976]"
//line /usr/local/go/src/encoding/json/fold.go:42
		}() {
//line /usr/local/go/src/encoding/json/fold.go:42
			_go_fuzz_dep_.CoverTab[27977]++
									nonLetter = true
//line /usr/local/go/src/encoding/json/fold.go:43
			// _ = "end of CoverTab[27977]"
		} else {
//line /usr/local/go/src/encoding/json/fold.go:44
			_go_fuzz_dep_.CoverTab[27978]++
//line /usr/local/go/src/encoding/json/fold.go:44
			if upper == 'K' || func() bool {
//line /usr/local/go/src/encoding/json/fold.go:44
				_go_fuzz_dep_.CoverTab[27979]++
//line /usr/local/go/src/encoding/json/fold.go:44
				return upper == 'S'
//line /usr/local/go/src/encoding/json/fold.go:44
				// _ = "end of CoverTab[27979]"
//line /usr/local/go/src/encoding/json/fold.go:44
			}() {
//line /usr/local/go/src/encoding/json/fold.go:44
				_go_fuzz_dep_.CoverTab[27980]++

										special = true
//line /usr/local/go/src/encoding/json/fold.go:46
				// _ = "end of CoverTab[27980]"
			} else {
//line /usr/local/go/src/encoding/json/fold.go:47
				_go_fuzz_dep_.CoverTab[27981]++
//line /usr/local/go/src/encoding/json/fold.go:47
				// _ = "end of CoverTab[27981]"
//line /usr/local/go/src/encoding/json/fold.go:47
			}
//line /usr/local/go/src/encoding/json/fold.go:47
			// _ = "end of CoverTab[27978]"
//line /usr/local/go/src/encoding/json/fold.go:47
		}
//line /usr/local/go/src/encoding/json/fold.go:47
		// _ = "end of CoverTab[27973]"
	}
//line /usr/local/go/src/encoding/json/fold.go:48
	// _ = "end of CoverTab[27968]"
//line /usr/local/go/src/encoding/json/fold.go:48
	_go_fuzz_dep_.CoverTab[27969]++
							if special {
//line /usr/local/go/src/encoding/json/fold.go:49
		_go_fuzz_dep_.CoverTab[27982]++
								return equalFoldRight
//line /usr/local/go/src/encoding/json/fold.go:50
		// _ = "end of CoverTab[27982]"
	} else {
//line /usr/local/go/src/encoding/json/fold.go:51
		_go_fuzz_dep_.CoverTab[27983]++
//line /usr/local/go/src/encoding/json/fold.go:51
		// _ = "end of CoverTab[27983]"
//line /usr/local/go/src/encoding/json/fold.go:51
	}
//line /usr/local/go/src/encoding/json/fold.go:51
	// _ = "end of CoverTab[27969]"
//line /usr/local/go/src/encoding/json/fold.go:51
	_go_fuzz_dep_.CoverTab[27970]++
							if nonLetter {
//line /usr/local/go/src/encoding/json/fold.go:52
		_go_fuzz_dep_.CoverTab[27984]++
								return asciiEqualFold
//line /usr/local/go/src/encoding/json/fold.go:53
		// _ = "end of CoverTab[27984]"
	} else {
//line /usr/local/go/src/encoding/json/fold.go:54
		_go_fuzz_dep_.CoverTab[27985]++
//line /usr/local/go/src/encoding/json/fold.go:54
		// _ = "end of CoverTab[27985]"
//line /usr/local/go/src/encoding/json/fold.go:54
	}
//line /usr/local/go/src/encoding/json/fold.go:54
	// _ = "end of CoverTab[27970]"
//line /usr/local/go/src/encoding/json/fold.go:54
	_go_fuzz_dep_.CoverTab[27971]++
							return simpleLetterEqualFold
//line /usr/local/go/src/encoding/json/fold.go:55
	// _ = "end of CoverTab[27971]"
}

// equalFoldRight is a specialization of bytes.EqualFold when s is
//line /usr/local/go/src/encoding/json/fold.go:58
// known to be all ASCII (including punctuation), but contains an 's',
//line /usr/local/go/src/encoding/json/fold.go:58
// 'S', 'k', or 'K', requiring a Unicode fold on the bytes in t.
//line /usr/local/go/src/encoding/json/fold.go:58
// See comments on foldFunc.
//line /usr/local/go/src/encoding/json/fold.go:62
func equalFoldRight(s, t []byte) bool {
//line /usr/local/go/src/encoding/json/fold.go:62
	_go_fuzz_dep_.CoverTab[27986]++
							for _, sb := range s {
//line /usr/local/go/src/encoding/json/fold.go:63
		_go_fuzz_dep_.CoverTab[27988]++
								if len(t) == 0 {
//line /usr/local/go/src/encoding/json/fold.go:64
			_go_fuzz_dep_.CoverTab[27992]++
									return false
//line /usr/local/go/src/encoding/json/fold.go:65
			// _ = "end of CoverTab[27992]"
		} else {
//line /usr/local/go/src/encoding/json/fold.go:66
			_go_fuzz_dep_.CoverTab[27993]++
//line /usr/local/go/src/encoding/json/fold.go:66
			// _ = "end of CoverTab[27993]"
//line /usr/local/go/src/encoding/json/fold.go:66
		}
//line /usr/local/go/src/encoding/json/fold.go:66
		// _ = "end of CoverTab[27988]"
//line /usr/local/go/src/encoding/json/fold.go:66
		_go_fuzz_dep_.CoverTab[27989]++
								tb := t[0]
								if tb < utf8.RuneSelf {
//line /usr/local/go/src/encoding/json/fold.go:68
			_go_fuzz_dep_.CoverTab[27994]++
									if sb != tb {
//line /usr/local/go/src/encoding/json/fold.go:69
				_go_fuzz_dep_.CoverTab[27996]++
										sbUpper := sb & caseMask
										if 'A' <= sbUpper && func() bool {
//line /usr/local/go/src/encoding/json/fold.go:71
					_go_fuzz_dep_.CoverTab[27997]++
//line /usr/local/go/src/encoding/json/fold.go:71
					return sbUpper <= 'Z'
//line /usr/local/go/src/encoding/json/fold.go:71
					// _ = "end of CoverTab[27997]"
//line /usr/local/go/src/encoding/json/fold.go:71
				}() {
//line /usr/local/go/src/encoding/json/fold.go:71
					_go_fuzz_dep_.CoverTab[27998]++
											if sbUpper != tb&caseMask {
//line /usr/local/go/src/encoding/json/fold.go:72
						_go_fuzz_dep_.CoverTab[27999]++
												return false
//line /usr/local/go/src/encoding/json/fold.go:73
						// _ = "end of CoverTab[27999]"
					} else {
//line /usr/local/go/src/encoding/json/fold.go:74
						_go_fuzz_dep_.CoverTab[28000]++
//line /usr/local/go/src/encoding/json/fold.go:74
						// _ = "end of CoverTab[28000]"
//line /usr/local/go/src/encoding/json/fold.go:74
					}
//line /usr/local/go/src/encoding/json/fold.go:74
					// _ = "end of CoverTab[27998]"
				} else {
//line /usr/local/go/src/encoding/json/fold.go:75
					_go_fuzz_dep_.CoverTab[28001]++
											return false
//line /usr/local/go/src/encoding/json/fold.go:76
					// _ = "end of CoverTab[28001]"
				}
//line /usr/local/go/src/encoding/json/fold.go:77
				// _ = "end of CoverTab[27996]"
			} else {
//line /usr/local/go/src/encoding/json/fold.go:78
				_go_fuzz_dep_.CoverTab[28002]++
//line /usr/local/go/src/encoding/json/fold.go:78
				// _ = "end of CoverTab[28002]"
//line /usr/local/go/src/encoding/json/fold.go:78
			}
//line /usr/local/go/src/encoding/json/fold.go:78
			// _ = "end of CoverTab[27994]"
//line /usr/local/go/src/encoding/json/fold.go:78
			_go_fuzz_dep_.CoverTab[27995]++
									t = t[1:]
									continue
//line /usr/local/go/src/encoding/json/fold.go:80
			// _ = "end of CoverTab[27995]"
		} else {
//line /usr/local/go/src/encoding/json/fold.go:81
			_go_fuzz_dep_.CoverTab[28003]++
//line /usr/local/go/src/encoding/json/fold.go:81
			// _ = "end of CoverTab[28003]"
//line /usr/local/go/src/encoding/json/fold.go:81
		}
//line /usr/local/go/src/encoding/json/fold.go:81
		// _ = "end of CoverTab[27989]"
//line /usr/local/go/src/encoding/json/fold.go:81
		_go_fuzz_dep_.CoverTab[27990]++

//line /usr/local/go/src/encoding/json/fold.go:84
		tr, size := utf8.DecodeRune(t)
		switch sb {
		case 's', 'S':
//line /usr/local/go/src/encoding/json/fold.go:86
			_go_fuzz_dep_.CoverTab[28004]++
									if tr != smallLongEss {
//line /usr/local/go/src/encoding/json/fold.go:87
				_go_fuzz_dep_.CoverTab[28007]++
										return false
//line /usr/local/go/src/encoding/json/fold.go:88
				// _ = "end of CoverTab[28007]"
			} else {
//line /usr/local/go/src/encoding/json/fold.go:89
				_go_fuzz_dep_.CoverTab[28008]++
//line /usr/local/go/src/encoding/json/fold.go:89
				// _ = "end of CoverTab[28008]"
//line /usr/local/go/src/encoding/json/fold.go:89
			}
//line /usr/local/go/src/encoding/json/fold.go:89
			// _ = "end of CoverTab[28004]"
		case 'k', 'K':
//line /usr/local/go/src/encoding/json/fold.go:90
			_go_fuzz_dep_.CoverTab[28005]++
									if tr != kelvin {
//line /usr/local/go/src/encoding/json/fold.go:91
				_go_fuzz_dep_.CoverTab[28009]++
										return false
//line /usr/local/go/src/encoding/json/fold.go:92
				// _ = "end of CoverTab[28009]"
			} else {
//line /usr/local/go/src/encoding/json/fold.go:93
				_go_fuzz_dep_.CoverTab[28010]++
//line /usr/local/go/src/encoding/json/fold.go:93
				// _ = "end of CoverTab[28010]"
//line /usr/local/go/src/encoding/json/fold.go:93
			}
//line /usr/local/go/src/encoding/json/fold.go:93
			// _ = "end of CoverTab[28005]"
		default:
//line /usr/local/go/src/encoding/json/fold.go:94
			_go_fuzz_dep_.CoverTab[28006]++
									return false
//line /usr/local/go/src/encoding/json/fold.go:95
			// _ = "end of CoverTab[28006]"
		}
//line /usr/local/go/src/encoding/json/fold.go:96
		// _ = "end of CoverTab[27990]"
//line /usr/local/go/src/encoding/json/fold.go:96
		_go_fuzz_dep_.CoverTab[27991]++
								t = t[size:]
//line /usr/local/go/src/encoding/json/fold.go:97
		// _ = "end of CoverTab[27991]"

	}
//line /usr/local/go/src/encoding/json/fold.go:99
	// _ = "end of CoverTab[27986]"
//line /usr/local/go/src/encoding/json/fold.go:99
	_go_fuzz_dep_.CoverTab[27987]++
							return len(t) == 0
//line /usr/local/go/src/encoding/json/fold.go:100
	// _ = "end of CoverTab[27987]"
}

// asciiEqualFold is a specialization of bytes.EqualFold for use when
//line /usr/local/go/src/encoding/json/fold.go:103
// s is all ASCII (but may contain non-letters) and contains no
//line /usr/local/go/src/encoding/json/fold.go:103
// special-folding letters.
//line /usr/local/go/src/encoding/json/fold.go:103
// See comments on foldFunc.
//line /usr/local/go/src/encoding/json/fold.go:107
func asciiEqualFold(s, t []byte) bool {
//line /usr/local/go/src/encoding/json/fold.go:107
	_go_fuzz_dep_.CoverTab[28011]++
							if len(s) != len(t) {
//line /usr/local/go/src/encoding/json/fold.go:108
		_go_fuzz_dep_.CoverTab[28014]++
								return false
//line /usr/local/go/src/encoding/json/fold.go:109
		// _ = "end of CoverTab[28014]"
	} else {
//line /usr/local/go/src/encoding/json/fold.go:110
		_go_fuzz_dep_.CoverTab[28015]++
//line /usr/local/go/src/encoding/json/fold.go:110
		// _ = "end of CoverTab[28015]"
//line /usr/local/go/src/encoding/json/fold.go:110
	}
//line /usr/local/go/src/encoding/json/fold.go:110
	// _ = "end of CoverTab[28011]"
//line /usr/local/go/src/encoding/json/fold.go:110
	_go_fuzz_dep_.CoverTab[28012]++
							for i, sb := range s {
//line /usr/local/go/src/encoding/json/fold.go:111
		_go_fuzz_dep_.CoverTab[28016]++
								tb := t[i]
								if sb == tb {
//line /usr/local/go/src/encoding/json/fold.go:113
			_go_fuzz_dep_.CoverTab[28018]++
									continue
//line /usr/local/go/src/encoding/json/fold.go:114
			// _ = "end of CoverTab[28018]"
		} else {
//line /usr/local/go/src/encoding/json/fold.go:115
			_go_fuzz_dep_.CoverTab[28019]++
//line /usr/local/go/src/encoding/json/fold.go:115
			// _ = "end of CoverTab[28019]"
//line /usr/local/go/src/encoding/json/fold.go:115
		}
//line /usr/local/go/src/encoding/json/fold.go:115
		// _ = "end of CoverTab[28016]"
//line /usr/local/go/src/encoding/json/fold.go:115
		_go_fuzz_dep_.CoverTab[28017]++
								if ('a' <= sb && func() bool {
//line /usr/local/go/src/encoding/json/fold.go:116
			_go_fuzz_dep_.CoverTab[28020]++
//line /usr/local/go/src/encoding/json/fold.go:116
			return sb <= 'z'
//line /usr/local/go/src/encoding/json/fold.go:116
			// _ = "end of CoverTab[28020]"
//line /usr/local/go/src/encoding/json/fold.go:116
		}()) || func() bool {
//line /usr/local/go/src/encoding/json/fold.go:116
			_go_fuzz_dep_.CoverTab[28021]++
//line /usr/local/go/src/encoding/json/fold.go:116
			return ('A' <= sb && func() bool {
//line /usr/local/go/src/encoding/json/fold.go:116
				_go_fuzz_dep_.CoverTab[28022]++
//line /usr/local/go/src/encoding/json/fold.go:116
				return sb <= 'Z'
//line /usr/local/go/src/encoding/json/fold.go:116
				// _ = "end of CoverTab[28022]"
//line /usr/local/go/src/encoding/json/fold.go:116
			}())
//line /usr/local/go/src/encoding/json/fold.go:116
			// _ = "end of CoverTab[28021]"
//line /usr/local/go/src/encoding/json/fold.go:116
		}() {
//line /usr/local/go/src/encoding/json/fold.go:116
			_go_fuzz_dep_.CoverTab[28023]++
									if sb&caseMask != tb&caseMask {
//line /usr/local/go/src/encoding/json/fold.go:117
				_go_fuzz_dep_.CoverTab[28024]++
										return false
//line /usr/local/go/src/encoding/json/fold.go:118
				// _ = "end of CoverTab[28024]"
			} else {
//line /usr/local/go/src/encoding/json/fold.go:119
				_go_fuzz_dep_.CoverTab[28025]++
//line /usr/local/go/src/encoding/json/fold.go:119
				// _ = "end of CoverTab[28025]"
//line /usr/local/go/src/encoding/json/fold.go:119
			}
//line /usr/local/go/src/encoding/json/fold.go:119
			// _ = "end of CoverTab[28023]"
		} else {
//line /usr/local/go/src/encoding/json/fold.go:120
			_go_fuzz_dep_.CoverTab[28026]++
									return false
//line /usr/local/go/src/encoding/json/fold.go:121
			// _ = "end of CoverTab[28026]"
		}
//line /usr/local/go/src/encoding/json/fold.go:122
		// _ = "end of CoverTab[28017]"
	}
//line /usr/local/go/src/encoding/json/fold.go:123
	// _ = "end of CoverTab[28012]"
//line /usr/local/go/src/encoding/json/fold.go:123
	_go_fuzz_dep_.CoverTab[28013]++
							return true
//line /usr/local/go/src/encoding/json/fold.go:124
	// _ = "end of CoverTab[28013]"
}

// simpleLetterEqualFold is a specialization of bytes.EqualFold for
//line /usr/local/go/src/encoding/json/fold.go:127
// use when s is all ASCII letters (no underscores, etc) and also
//line /usr/local/go/src/encoding/json/fold.go:127
// doesn't contain 'k', 'K', 's', or 'S'.
//line /usr/local/go/src/encoding/json/fold.go:127
// See comments on foldFunc.
//line /usr/local/go/src/encoding/json/fold.go:131
func simpleLetterEqualFold(s, t []byte) bool {
//line /usr/local/go/src/encoding/json/fold.go:131
	_go_fuzz_dep_.CoverTab[28027]++
							if len(s) != len(t) {
//line /usr/local/go/src/encoding/json/fold.go:132
		_go_fuzz_dep_.CoverTab[28030]++
								return false
//line /usr/local/go/src/encoding/json/fold.go:133
		// _ = "end of CoverTab[28030]"
	} else {
//line /usr/local/go/src/encoding/json/fold.go:134
		_go_fuzz_dep_.CoverTab[28031]++
//line /usr/local/go/src/encoding/json/fold.go:134
		// _ = "end of CoverTab[28031]"
//line /usr/local/go/src/encoding/json/fold.go:134
	}
//line /usr/local/go/src/encoding/json/fold.go:134
	// _ = "end of CoverTab[28027]"
//line /usr/local/go/src/encoding/json/fold.go:134
	_go_fuzz_dep_.CoverTab[28028]++
							for i, b := range s {
//line /usr/local/go/src/encoding/json/fold.go:135
		_go_fuzz_dep_.CoverTab[28032]++
								if b&caseMask != t[i]&caseMask {
//line /usr/local/go/src/encoding/json/fold.go:136
			_go_fuzz_dep_.CoverTab[28033]++
									return false
//line /usr/local/go/src/encoding/json/fold.go:137
			// _ = "end of CoverTab[28033]"
		} else {
//line /usr/local/go/src/encoding/json/fold.go:138
			_go_fuzz_dep_.CoverTab[28034]++
//line /usr/local/go/src/encoding/json/fold.go:138
			// _ = "end of CoverTab[28034]"
//line /usr/local/go/src/encoding/json/fold.go:138
		}
//line /usr/local/go/src/encoding/json/fold.go:138
		// _ = "end of CoverTab[28032]"
	}
//line /usr/local/go/src/encoding/json/fold.go:139
	// _ = "end of CoverTab[28028]"
//line /usr/local/go/src/encoding/json/fold.go:139
	_go_fuzz_dep_.CoverTab[28029]++
							return true
//line /usr/local/go/src/encoding/json/fold.go:140
	// _ = "end of CoverTab[28029]"
}

//line /usr/local/go/src/encoding/json/fold.go:141
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/json/fold.go:141
var _ = _go_fuzz_dep_.CoverTab
