// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/path/filepath/match.go:5
package filepath

//line /usr/local/go/src/path/filepath/match.go:5
import (
//line /usr/local/go/src/path/filepath/match.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/path/filepath/match.go:5
)
//line /usr/local/go/src/path/filepath/match.go:5
import (
//line /usr/local/go/src/path/filepath/match.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/path/filepath/match.go:5
)

import (
	"errors"
	"os"
	"runtime"
	"sort"
	"strings"
	"unicode/utf8"
)

// ErrBadPattern indicates a pattern was malformed.
var ErrBadPattern = errors.New("syntax error in pattern")

// Match reports whether name matches the shell file name pattern.
//line /usr/local/go/src/path/filepath/match.go:19
// The pattern syntax is:
//line /usr/local/go/src/path/filepath/match.go:19
//
//line /usr/local/go/src/path/filepath/match.go:19
//	pattern:
//line /usr/local/go/src/path/filepath/match.go:19
//		{ term }
//line /usr/local/go/src/path/filepath/match.go:19
//	term:
//line /usr/local/go/src/path/filepath/match.go:19
//		'*'         matches any sequence of non-Separator characters
//line /usr/local/go/src/path/filepath/match.go:19
//		'?'         matches any single non-Separator character
//line /usr/local/go/src/path/filepath/match.go:19
//		'[' [ '^' ] { character-range } ']'
//line /usr/local/go/src/path/filepath/match.go:19
//		            character class (must be non-empty)
//line /usr/local/go/src/path/filepath/match.go:19
//		c           matches character c (c != '*', '?', '\\', '[')
//line /usr/local/go/src/path/filepath/match.go:19
//		'\\' c      matches character c
//line /usr/local/go/src/path/filepath/match.go:19
//
//line /usr/local/go/src/path/filepath/match.go:19
//	character-range:
//line /usr/local/go/src/path/filepath/match.go:19
//		c           matches character c (c != '\\', '-', ']')
//line /usr/local/go/src/path/filepath/match.go:19
//		'\\' c      matches character c
//line /usr/local/go/src/path/filepath/match.go:19
//		lo '-' hi   matches character c for lo <= c <= hi
//line /usr/local/go/src/path/filepath/match.go:19
//
//line /usr/local/go/src/path/filepath/match.go:19
// Match requires pattern to match all of name, not just a substring.
//line /usr/local/go/src/path/filepath/match.go:19
// The only possible returned error is ErrBadPattern, when pattern
//line /usr/local/go/src/path/filepath/match.go:19
// is malformed.
//line /usr/local/go/src/path/filepath/match.go:19
//
//line /usr/local/go/src/path/filepath/match.go:19
// On Windows, escaping is disabled. Instead, '\\' is treated as
//line /usr/local/go/src/path/filepath/match.go:19
// path separator.
//line /usr/local/go/src/path/filepath/match.go:43
func Match(pattern, name string) (matched bool, err error) {
//line /usr/local/go/src/path/filepath/match.go:43
	_go_fuzz_dep_.CoverTab[17768]++
Pattern:
	for len(pattern) > 0 {
//line /usr/local/go/src/path/filepath/match.go:45
		_go_fuzz_dep_.CoverTab[17770]++
								var star bool
								var chunk string
								star, chunk, pattern = scanChunk(pattern)
								if star && func() bool {
//line /usr/local/go/src/path/filepath/match.go:49
			_go_fuzz_dep_.CoverTab[17775]++
//line /usr/local/go/src/path/filepath/match.go:49
			return chunk == ""
//line /usr/local/go/src/path/filepath/match.go:49
			// _ = "end of CoverTab[17775]"
//line /usr/local/go/src/path/filepath/match.go:49
		}() {
//line /usr/local/go/src/path/filepath/match.go:49
			_go_fuzz_dep_.CoverTab[17776]++

									return !strings.Contains(name, string(Separator)), nil
//line /usr/local/go/src/path/filepath/match.go:51
			// _ = "end of CoverTab[17776]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:52
			_go_fuzz_dep_.CoverTab[17777]++
//line /usr/local/go/src/path/filepath/match.go:52
			// _ = "end of CoverTab[17777]"
//line /usr/local/go/src/path/filepath/match.go:52
		}
//line /usr/local/go/src/path/filepath/match.go:52
		// _ = "end of CoverTab[17770]"
//line /usr/local/go/src/path/filepath/match.go:52
		_go_fuzz_dep_.CoverTab[17771]++

								t, ok, err := matchChunk(chunk, name)

//line /usr/local/go/src/path/filepath/match.go:58
		if ok && func() bool {
//line /usr/local/go/src/path/filepath/match.go:58
			_go_fuzz_dep_.CoverTab[17778]++
//line /usr/local/go/src/path/filepath/match.go:58
			return (len(t) == 0 || func() bool {
//line /usr/local/go/src/path/filepath/match.go:58
				_go_fuzz_dep_.CoverTab[17779]++
//line /usr/local/go/src/path/filepath/match.go:58
				return len(pattern) > 0
//line /usr/local/go/src/path/filepath/match.go:58
				// _ = "end of CoverTab[17779]"
//line /usr/local/go/src/path/filepath/match.go:58
			}())
//line /usr/local/go/src/path/filepath/match.go:58
			// _ = "end of CoverTab[17778]"
//line /usr/local/go/src/path/filepath/match.go:58
		}() {
//line /usr/local/go/src/path/filepath/match.go:58
			_go_fuzz_dep_.CoverTab[17780]++
									name = t
									continue
//line /usr/local/go/src/path/filepath/match.go:60
			// _ = "end of CoverTab[17780]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:61
			_go_fuzz_dep_.CoverTab[17781]++
//line /usr/local/go/src/path/filepath/match.go:61
			// _ = "end of CoverTab[17781]"
//line /usr/local/go/src/path/filepath/match.go:61
		}
//line /usr/local/go/src/path/filepath/match.go:61
		// _ = "end of CoverTab[17771]"
//line /usr/local/go/src/path/filepath/match.go:61
		_go_fuzz_dep_.CoverTab[17772]++
								if err != nil {
//line /usr/local/go/src/path/filepath/match.go:62
			_go_fuzz_dep_.CoverTab[17782]++
									return false, err
//line /usr/local/go/src/path/filepath/match.go:63
			// _ = "end of CoverTab[17782]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:64
			_go_fuzz_dep_.CoverTab[17783]++
//line /usr/local/go/src/path/filepath/match.go:64
			// _ = "end of CoverTab[17783]"
//line /usr/local/go/src/path/filepath/match.go:64
		}
//line /usr/local/go/src/path/filepath/match.go:64
		// _ = "end of CoverTab[17772]"
//line /usr/local/go/src/path/filepath/match.go:64
		_go_fuzz_dep_.CoverTab[17773]++
								if star {
//line /usr/local/go/src/path/filepath/match.go:65
			_go_fuzz_dep_.CoverTab[17784]++

//line /usr/local/go/src/path/filepath/match.go:68
			for i := 0; i < len(name) && func() bool {
//line /usr/local/go/src/path/filepath/match.go:68
				_go_fuzz_dep_.CoverTab[17785]++
//line /usr/local/go/src/path/filepath/match.go:68
				return name[i] != Separator
//line /usr/local/go/src/path/filepath/match.go:68
				// _ = "end of CoverTab[17785]"
//line /usr/local/go/src/path/filepath/match.go:68
			}(); i++ {
//line /usr/local/go/src/path/filepath/match.go:68
				_go_fuzz_dep_.CoverTab[17786]++
										t, ok, err := matchChunk(chunk, name[i+1:])
										if ok {
//line /usr/local/go/src/path/filepath/match.go:70
					_go_fuzz_dep_.CoverTab[17788]++

											if len(pattern) == 0 && func() bool {
//line /usr/local/go/src/path/filepath/match.go:72
						_go_fuzz_dep_.CoverTab[17790]++
//line /usr/local/go/src/path/filepath/match.go:72
						return len(t) > 0
//line /usr/local/go/src/path/filepath/match.go:72
						// _ = "end of CoverTab[17790]"
//line /usr/local/go/src/path/filepath/match.go:72
					}() {
//line /usr/local/go/src/path/filepath/match.go:72
						_go_fuzz_dep_.CoverTab[17791]++
												continue
//line /usr/local/go/src/path/filepath/match.go:73
						// _ = "end of CoverTab[17791]"
					} else {
//line /usr/local/go/src/path/filepath/match.go:74
						_go_fuzz_dep_.CoverTab[17792]++
//line /usr/local/go/src/path/filepath/match.go:74
						// _ = "end of CoverTab[17792]"
//line /usr/local/go/src/path/filepath/match.go:74
					}
//line /usr/local/go/src/path/filepath/match.go:74
					// _ = "end of CoverTab[17788]"
//line /usr/local/go/src/path/filepath/match.go:74
					_go_fuzz_dep_.CoverTab[17789]++
											name = t
											continue Pattern
//line /usr/local/go/src/path/filepath/match.go:76
					// _ = "end of CoverTab[17789]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:77
					_go_fuzz_dep_.CoverTab[17793]++
//line /usr/local/go/src/path/filepath/match.go:77
					// _ = "end of CoverTab[17793]"
//line /usr/local/go/src/path/filepath/match.go:77
				}
//line /usr/local/go/src/path/filepath/match.go:77
				// _ = "end of CoverTab[17786]"
//line /usr/local/go/src/path/filepath/match.go:77
				_go_fuzz_dep_.CoverTab[17787]++
										if err != nil {
//line /usr/local/go/src/path/filepath/match.go:78
					_go_fuzz_dep_.CoverTab[17794]++
											return false, err
//line /usr/local/go/src/path/filepath/match.go:79
					// _ = "end of CoverTab[17794]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:80
					_go_fuzz_dep_.CoverTab[17795]++
//line /usr/local/go/src/path/filepath/match.go:80
					// _ = "end of CoverTab[17795]"
//line /usr/local/go/src/path/filepath/match.go:80
				}
//line /usr/local/go/src/path/filepath/match.go:80
				// _ = "end of CoverTab[17787]"
			}
//line /usr/local/go/src/path/filepath/match.go:81
			// _ = "end of CoverTab[17784]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:82
			_go_fuzz_dep_.CoverTab[17796]++
//line /usr/local/go/src/path/filepath/match.go:82
			// _ = "end of CoverTab[17796]"
//line /usr/local/go/src/path/filepath/match.go:82
		}
//line /usr/local/go/src/path/filepath/match.go:82
		// _ = "end of CoverTab[17773]"
//line /usr/local/go/src/path/filepath/match.go:82
		_go_fuzz_dep_.CoverTab[17774]++
								return false, nil
//line /usr/local/go/src/path/filepath/match.go:83
		// _ = "end of CoverTab[17774]"
	}
//line /usr/local/go/src/path/filepath/match.go:84
	// _ = "end of CoverTab[17768]"
//line /usr/local/go/src/path/filepath/match.go:84
	_go_fuzz_dep_.CoverTab[17769]++
							return len(name) == 0, nil
//line /usr/local/go/src/path/filepath/match.go:85
	// _ = "end of CoverTab[17769]"
}

// scanChunk gets the next segment of pattern, which is a non-star string
//line /usr/local/go/src/path/filepath/match.go:88
// possibly preceded by a star.
//line /usr/local/go/src/path/filepath/match.go:90
func scanChunk(pattern string) (star bool, chunk, rest string) {
//line /usr/local/go/src/path/filepath/match.go:90
	_go_fuzz_dep_.CoverTab[17797]++
							for len(pattern) > 0 && func() bool {
//line /usr/local/go/src/path/filepath/match.go:91
		_go_fuzz_dep_.CoverTab[17800]++
//line /usr/local/go/src/path/filepath/match.go:91
		return pattern[0] == '*'
//line /usr/local/go/src/path/filepath/match.go:91
		// _ = "end of CoverTab[17800]"
//line /usr/local/go/src/path/filepath/match.go:91
	}() {
//line /usr/local/go/src/path/filepath/match.go:91
		_go_fuzz_dep_.CoverTab[17801]++
								pattern = pattern[1:]
								star = true
//line /usr/local/go/src/path/filepath/match.go:93
		// _ = "end of CoverTab[17801]"
	}
//line /usr/local/go/src/path/filepath/match.go:94
	// _ = "end of CoverTab[17797]"
//line /usr/local/go/src/path/filepath/match.go:94
	_go_fuzz_dep_.CoverTab[17798]++
							inrange := false
							var i int
Scan:
	for i = 0; i < len(pattern); i++ {
//line /usr/local/go/src/path/filepath/match.go:98
		_go_fuzz_dep_.CoverTab[17802]++
								switch pattern[i] {
		case '\\':
//line /usr/local/go/src/path/filepath/match.go:100
			_go_fuzz_dep_.CoverTab[17803]++
									if runtime.GOOS != "windows" {
//line /usr/local/go/src/path/filepath/match.go:101
				_go_fuzz_dep_.CoverTab[17808]++

										if i+1 < len(pattern) {
//line /usr/local/go/src/path/filepath/match.go:103
					_go_fuzz_dep_.CoverTab[17809]++
											i++
//line /usr/local/go/src/path/filepath/match.go:104
					// _ = "end of CoverTab[17809]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:105
					_go_fuzz_dep_.CoverTab[17810]++
//line /usr/local/go/src/path/filepath/match.go:105
					// _ = "end of CoverTab[17810]"
//line /usr/local/go/src/path/filepath/match.go:105
				}
//line /usr/local/go/src/path/filepath/match.go:105
				// _ = "end of CoverTab[17808]"
			} else {
//line /usr/local/go/src/path/filepath/match.go:106
				_go_fuzz_dep_.CoverTab[17811]++
//line /usr/local/go/src/path/filepath/match.go:106
				// _ = "end of CoverTab[17811]"
//line /usr/local/go/src/path/filepath/match.go:106
			}
//line /usr/local/go/src/path/filepath/match.go:106
			// _ = "end of CoverTab[17803]"
		case '[':
//line /usr/local/go/src/path/filepath/match.go:107
			_go_fuzz_dep_.CoverTab[17804]++
									inrange = true
//line /usr/local/go/src/path/filepath/match.go:108
			// _ = "end of CoverTab[17804]"
		case ']':
//line /usr/local/go/src/path/filepath/match.go:109
			_go_fuzz_dep_.CoverTab[17805]++
									inrange = false
//line /usr/local/go/src/path/filepath/match.go:110
			// _ = "end of CoverTab[17805]"
		case '*':
//line /usr/local/go/src/path/filepath/match.go:111
			_go_fuzz_dep_.CoverTab[17806]++
									if !inrange {
//line /usr/local/go/src/path/filepath/match.go:112
				_go_fuzz_dep_.CoverTab[17812]++
										break Scan
//line /usr/local/go/src/path/filepath/match.go:113
				// _ = "end of CoverTab[17812]"
			} else {
//line /usr/local/go/src/path/filepath/match.go:114
				_go_fuzz_dep_.CoverTab[17813]++
//line /usr/local/go/src/path/filepath/match.go:114
				// _ = "end of CoverTab[17813]"
//line /usr/local/go/src/path/filepath/match.go:114
			}
//line /usr/local/go/src/path/filepath/match.go:114
			// _ = "end of CoverTab[17806]"
//line /usr/local/go/src/path/filepath/match.go:114
		default:
//line /usr/local/go/src/path/filepath/match.go:114
			_go_fuzz_dep_.CoverTab[17807]++
//line /usr/local/go/src/path/filepath/match.go:114
			// _ = "end of CoverTab[17807]"
		}
//line /usr/local/go/src/path/filepath/match.go:115
		// _ = "end of CoverTab[17802]"
	}
//line /usr/local/go/src/path/filepath/match.go:116
	// _ = "end of CoverTab[17798]"
//line /usr/local/go/src/path/filepath/match.go:116
	_go_fuzz_dep_.CoverTab[17799]++
							return star, pattern[0:i], pattern[i:]
//line /usr/local/go/src/path/filepath/match.go:117
	// _ = "end of CoverTab[17799]"
}

// matchChunk checks whether chunk matches the beginning of s.
//line /usr/local/go/src/path/filepath/match.go:120
// If so, it returns the remainder of s (after the match).
//line /usr/local/go/src/path/filepath/match.go:120
// Chunk is all single-character operators: literals, char classes, and ?.
//line /usr/local/go/src/path/filepath/match.go:123
func matchChunk(chunk, s string) (rest string, ok bool, err error) {
//line /usr/local/go/src/path/filepath/match.go:123
	_go_fuzz_dep_.CoverTab[17814]++

//line /usr/local/go/src/path/filepath/match.go:127
	failed := false
	for len(chunk) > 0 {
//line /usr/local/go/src/path/filepath/match.go:128
		_go_fuzz_dep_.CoverTab[17817]++
								if !failed && func() bool {
//line /usr/local/go/src/path/filepath/match.go:129
			_go_fuzz_dep_.CoverTab[17819]++
//line /usr/local/go/src/path/filepath/match.go:129
			return len(s) == 0
//line /usr/local/go/src/path/filepath/match.go:129
			// _ = "end of CoverTab[17819]"
//line /usr/local/go/src/path/filepath/match.go:129
		}() {
//line /usr/local/go/src/path/filepath/match.go:129
			_go_fuzz_dep_.CoverTab[17820]++
									failed = true
//line /usr/local/go/src/path/filepath/match.go:130
			// _ = "end of CoverTab[17820]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:131
			_go_fuzz_dep_.CoverTab[17821]++
//line /usr/local/go/src/path/filepath/match.go:131
			// _ = "end of CoverTab[17821]"
//line /usr/local/go/src/path/filepath/match.go:131
		}
//line /usr/local/go/src/path/filepath/match.go:131
		// _ = "end of CoverTab[17817]"
//line /usr/local/go/src/path/filepath/match.go:131
		_go_fuzz_dep_.CoverTab[17818]++
								switch chunk[0] {
		case '[':
//line /usr/local/go/src/path/filepath/match.go:133
			_go_fuzz_dep_.CoverTab[17822]++
			// character class
			var r rune
			if !failed {
//line /usr/local/go/src/path/filepath/match.go:136
				_go_fuzz_dep_.CoverTab[17832]++
										var n int
										r, n = utf8.DecodeRuneInString(s)
										s = s[n:]
//line /usr/local/go/src/path/filepath/match.go:139
				// _ = "end of CoverTab[17832]"
			} else {
//line /usr/local/go/src/path/filepath/match.go:140
				_go_fuzz_dep_.CoverTab[17833]++
//line /usr/local/go/src/path/filepath/match.go:140
				// _ = "end of CoverTab[17833]"
//line /usr/local/go/src/path/filepath/match.go:140
			}
//line /usr/local/go/src/path/filepath/match.go:140
			// _ = "end of CoverTab[17822]"
//line /usr/local/go/src/path/filepath/match.go:140
			_go_fuzz_dep_.CoverTab[17823]++
									chunk = chunk[1:]

									negated := false
									if len(chunk) > 0 && func() bool {
//line /usr/local/go/src/path/filepath/match.go:144
				_go_fuzz_dep_.CoverTab[17834]++
//line /usr/local/go/src/path/filepath/match.go:144
				return chunk[0] == '^'
//line /usr/local/go/src/path/filepath/match.go:144
				// _ = "end of CoverTab[17834]"
//line /usr/local/go/src/path/filepath/match.go:144
			}() {
//line /usr/local/go/src/path/filepath/match.go:144
				_go_fuzz_dep_.CoverTab[17835]++
										negated = true
										chunk = chunk[1:]
//line /usr/local/go/src/path/filepath/match.go:146
				// _ = "end of CoverTab[17835]"
			} else {
//line /usr/local/go/src/path/filepath/match.go:147
				_go_fuzz_dep_.CoverTab[17836]++
//line /usr/local/go/src/path/filepath/match.go:147
				// _ = "end of CoverTab[17836]"
//line /usr/local/go/src/path/filepath/match.go:147
			}
//line /usr/local/go/src/path/filepath/match.go:147
			// _ = "end of CoverTab[17823]"
//line /usr/local/go/src/path/filepath/match.go:147
			_go_fuzz_dep_.CoverTab[17824]++

									match := false
									nrange := 0
									for {
//line /usr/local/go/src/path/filepath/match.go:151
				_go_fuzz_dep_.CoverTab[17837]++
										if len(chunk) > 0 && func() bool {
//line /usr/local/go/src/path/filepath/match.go:152
					_go_fuzz_dep_.CoverTab[17842]++
//line /usr/local/go/src/path/filepath/match.go:152
					return chunk[0] == ']'
//line /usr/local/go/src/path/filepath/match.go:152
					// _ = "end of CoverTab[17842]"
//line /usr/local/go/src/path/filepath/match.go:152
				}() && func() bool {
//line /usr/local/go/src/path/filepath/match.go:152
					_go_fuzz_dep_.CoverTab[17843]++
//line /usr/local/go/src/path/filepath/match.go:152
					return nrange > 0
//line /usr/local/go/src/path/filepath/match.go:152
					// _ = "end of CoverTab[17843]"
//line /usr/local/go/src/path/filepath/match.go:152
				}() {
//line /usr/local/go/src/path/filepath/match.go:152
					_go_fuzz_dep_.CoverTab[17844]++
											chunk = chunk[1:]
											break
//line /usr/local/go/src/path/filepath/match.go:154
					// _ = "end of CoverTab[17844]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:155
					_go_fuzz_dep_.CoverTab[17845]++
//line /usr/local/go/src/path/filepath/match.go:155
					// _ = "end of CoverTab[17845]"
//line /usr/local/go/src/path/filepath/match.go:155
				}
//line /usr/local/go/src/path/filepath/match.go:155
				// _ = "end of CoverTab[17837]"
//line /usr/local/go/src/path/filepath/match.go:155
				_go_fuzz_dep_.CoverTab[17838]++
										var lo, hi rune
										if lo, chunk, err = getEsc(chunk); err != nil {
//line /usr/local/go/src/path/filepath/match.go:157
					_go_fuzz_dep_.CoverTab[17846]++
											return "", false, err
//line /usr/local/go/src/path/filepath/match.go:158
					// _ = "end of CoverTab[17846]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:159
					_go_fuzz_dep_.CoverTab[17847]++
//line /usr/local/go/src/path/filepath/match.go:159
					// _ = "end of CoverTab[17847]"
//line /usr/local/go/src/path/filepath/match.go:159
				}
//line /usr/local/go/src/path/filepath/match.go:159
				// _ = "end of CoverTab[17838]"
//line /usr/local/go/src/path/filepath/match.go:159
				_go_fuzz_dep_.CoverTab[17839]++
										hi = lo
										if chunk[0] == '-' {
//line /usr/local/go/src/path/filepath/match.go:161
					_go_fuzz_dep_.CoverTab[17848]++
											if hi, chunk, err = getEsc(chunk[1:]); err != nil {
//line /usr/local/go/src/path/filepath/match.go:162
						_go_fuzz_dep_.CoverTab[17849]++
												return "", false, err
//line /usr/local/go/src/path/filepath/match.go:163
						// _ = "end of CoverTab[17849]"
					} else {
//line /usr/local/go/src/path/filepath/match.go:164
						_go_fuzz_dep_.CoverTab[17850]++
//line /usr/local/go/src/path/filepath/match.go:164
						// _ = "end of CoverTab[17850]"
//line /usr/local/go/src/path/filepath/match.go:164
					}
//line /usr/local/go/src/path/filepath/match.go:164
					// _ = "end of CoverTab[17848]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:165
					_go_fuzz_dep_.CoverTab[17851]++
//line /usr/local/go/src/path/filepath/match.go:165
					// _ = "end of CoverTab[17851]"
//line /usr/local/go/src/path/filepath/match.go:165
				}
//line /usr/local/go/src/path/filepath/match.go:165
				// _ = "end of CoverTab[17839]"
//line /usr/local/go/src/path/filepath/match.go:165
				_go_fuzz_dep_.CoverTab[17840]++
										if lo <= r && func() bool {
//line /usr/local/go/src/path/filepath/match.go:166
					_go_fuzz_dep_.CoverTab[17852]++
//line /usr/local/go/src/path/filepath/match.go:166
					return r <= hi
//line /usr/local/go/src/path/filepath/match.go:166
					// _ = "end of CoverTab[17852]"
//line /usr/local/go/src/path/filepath/match.go:166
				}() {
//line /usr/local/go/src/path/filepath/match.go:166
					_go_fuzz_dep_.CoverTab[17853]++
											match = true
//line /usr/local/go/src/path/filepath/match.go:167
					// _ = "end of CoverTab[17853]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:168
					_go_fuzz_dep_.CoverTab[17854]++
//line /usr/local/go/src/path/filepath/match.go:168
					// _ = "end of CoverTab[17854]"
//line /usr/local/go/src/path/filepath/match.go:168
				}
//line /usr/local/go/src/path/filepath/match.go:168
				// _ = "end of CoverTab[17840]"
//line /usr/local/go/src/path/filepath/match.go:168
				_go_fuzz_dep_.CoverTab[17841]++
										nrange++
//line /usr/local/go/src/path/filepath/match.go:169
				// _ = "end of CoverTab[17841]"
			}
//line /usr/local/go/src/path/filepath/match.go:170
			// _ = "end of CoverTab[17824]"
//line /usr/local/go/src/path/filepath/match.go:170
			_go_fuzz_dep_.CoverTab[17825]++
									if match == negated {
//line /usr/local/go/src/path/filepath/match.go:171
				_go_fuzz_dep_.CoverTab[17855]++
										failed = true
//line /usr/local/go/src/path/filepath/match.go:172
				// _ = "end of CoverTab[17855]"
			} else {
//line /usr/local/go/src/path/filepath/match.go:173
				_go_fuzz_dep_.CoverTab[17856]++
//line /usr/local/go/src/path/filepath/match.go:173
				// _ = "end of CoverTab[17856]"
//line /usr/local/go/src/path/filepath/match.go:173
			}
//line /usr/local/go/src/path/filepath/match.go:173
			// _ = "end of CoverTab[17825]"

		case '?':
//line /usr/local/go/src/path/filepath/match.go:175
			_go_fuzz_dep_.CoverTab[17826]++
									if !failed {
//line /usr/local/go/src/path/filepath/match.go:176
				_go_fuzz_dep_.CoverTab[17857]++
										if s[0] == Separator {
//line /usr/local/go/src/path/filepath/match.go:177
					_go_fuzz_dep_.CoverTab[17859]++
											failed = true
//line /usr/local/go/src/path/filepath/match.go:178
					// _ = "end of CoverTab[17859]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:179
					_go_fuzz_dep_.CoverTab[17860]++
//line /usr/local/go/src/path/filepath/match.go:179
					// _ = "end of CoverTab[17860]"
//line /usr/local/go/src/path/filepath/match.go:179
				}
//line /usr/local/go/src/path/filepath/match.go:179
				// _ = "end of CoverTab[17857]"
//line /usr/local/go/src/path/filepath/match.go:179
				_go_fuzz_dep_.CoverTab[17858]++
										_, n := utf8.DecodeRuneInString(s)
										s = s[n:]
//line /usr/local/go/src/path/filepath/match.go:181
				// _ = "end of CoverTab[17858]"
			} else {
//line /usr/local/go/src/path/filepath/match.go:182
				_go_fuzz_dep_.CoverTab[17861]++
//line /usr/local/go/src/path/filepath/match.go:182
				// _ = "end of CoverTab[17861]"
//line /usr/local/go/src/path/filepath/match.go:182
			}
//line /usr/local/go/src/path/filepath/match.go:182
			// _ = "end of CoverTab[17826]"
//line /usr/local/go/src/path/filepath/match.go:182
			_go_fuzz_dep_.CoverTab[17827]++
									chunk = chunk[1:]
//line /usr/local/go/src/path/filepath/match.go:183
			// _ = "end of CoverTab[17827]"

		case '\\':
//line /usr/local/go/src/path/filepath/match.go:185
			_go_fuzz_dep_.CoverTab[17828]++
									if runtime.GOOS != "windows" {
//line /usr/local/go/src/path/filepath/match.go:186
				_go_fuzz_dep_.CoverTab[17862]++
										chunk = chunk[1:]
										if len(chunk) == 0 {
//line /usr/local/go/src/path/filepath/match.go:188
					_go_fuzz_dep_.CoverTab[17863]++
											return "", false, ErrBadPattern
//line /usr/local/go/src/path/filepath/match.go:189
					// _ = "end of CoverTab[17863]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:190
					_go_fuzz_dep_.CoverTab[17864]++
//line /usr/local/go/src/path/filepath/match.go:190
					// _ = "end of CoverTab[17864]"
//line /usr/local/go/src/path/filepath/match.go:190
				}
//line /usr/local/go/src/path/filepath/match.go:190
				// _ = "end of CoverTab[17862]"
			} else {
//line /usr/local/go/src/path/filepath/match.go:191
				_go_fuzz_dep_.CoverTab[17865]++
//line /usr/local/go/src/path/filepath/match.go:191
				// _ = "end of CoverTab[17865]"
//line /usr/local/go/src/path/filepath/match.go:191
			}
//line /usr/local/go/src/path/filepath/match.go:191
			// _ = "end of CoverTab[17828]"
//line /usr/local/go/src/path/filepath/match.go:191
			_go_fuzz_dep_.CoverTab[17829]++
									fallthrough
//line /usr/local/go/src/path/filepath/match.go:192
			// _ = "end of CoverTab[17829]"

		default:
//line /usr/local/go/src/path/filepath/match.go:194
			_go_fuzz_dep_.CoverTab[17830]++
									if !failed {
//line /usr/local/go/src/path/filepath/match.go:195
				_go_fuzz_dep_.CoverTab[17866]++
										if chunk[0] != s[0] {
//line /usr/local/go/src/path/filepath/match.go:196
					_go_fuzz_dep_.CoverTab[17868]++
											failed = true
//line /usr/local/go/src/path/filepath/match.go:197
					// _ = "end of CoverTab[17868]"
				} else {
//line /usr/local/go/src/path/filepath/match.go:198
					_go_fuzz_dep_.CoverTab[17869]++
//line /usr/local/go/src/path/filepath/match.go:198
					// _ = "end of CoverTab[17869]"
//line /usr/local/go/src/path/filepath/match.go:198
				}
//line /usr/local/go/src/path/filepath/match.go:198
				// _ = "end of CoverTab[17866]"
//line /usr/local/go/src/path/filepath/match.go:198
				_go_fuzz_dep_.CoverTab[17867]++
										s = s[1:]
//line /usr/local/go/src/path/filepath/match.go:199
				// _ = "end of CoverTab[17867]"
			} else {
//line /usr/local/go/src/path/filepath/match.go:200
				_go_fuzz_dep_.CoverTab[17870]++
//line /usr/local/go/src/path/filepath/match.go:200
				// _ = "end of CoverTab[17870]"
//line /usr/local/go/src/path/filepath/match.go:200
			}
//line /usr/local/go/src/path/filepath/match.go:200
			// _ = "end of CoverTab[17830]"
//line /usr/local/go/src/path/filepath/match.go:200
			_go_fuzz_dep_.CoverTab[17831]++
									chunk = chunk[1:]
//line /usr/local/go/src/path/filepath/match.go:201
			// _ = "end of CoverTab[17831]"
		}
//line /usr/local/go/src/path/filepath/match.go:202
		// _ = "end of CoverTab[17818]"
	}
//line /usr/local/go/src/path/filepath/match.go:203
	// _ = "end of CoverTab[17814]"
//line /usr/local/go/src/path/filepath/match.go:203
	_go_fuzz_dep_.CoverTab[17815]++
							if failed {
//line /usr/local/go/src/path/filepath/match.go:204
		_go_fuzz_dep_.CoverTab[17871]++
								return "", false, nil
//line /usr/local/go/src/path/filepath/match.go:205
		// _ = "end of CoverTab[17871]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:206
		_go_fuzz_dep_.CoverTab[17872]++
//line /usr/local/go/src/path/filepath/match.go:206
		// _ = "end of CoverTab[17872]"
//line /usr/local/go/src/path/filepath/match.go:206
	}
//line /usr/local/go/src/path/filepath/match.go:206
	// _ = "end of CoverTab[17815]"
//line /usr/local/go/src/path/filepath/match.go:206
	_go_fuzz_dep_.CoverTab[17816]++
							return s, true, nil
//line /usr/local/go/src/path/filepath/match.go:207
	// _ = "end of CoverTab[17816]"
}

// getEsc gets a possibly-escaped character from chunk, for a character class.
func getEsc(chunk string) (r rune, nchunk string, err error) {
//line /usr/local/go/src/path/filepath/match.go:211
	_go_fuzz_dep_.CoverTab[17873]++
							if len(chunk) == 0 || func() bool {
//line /usr/local/go/src/path/filepath/match.go:212
		_go_fuzz_dep_.CoverTab[17878]++
//line /usr/local/go/src/path/filepath/match.go:212
		return chunk[0] == '-'
//line /usr/local/go/src/path/filepath/match.go:212
		// _ = "end of CoverTab[17878]"
//line /usr/local/go/src/path/filepath/match.go:212
	}() || func() bool {
//line /usr/local/go/src/path/filepath/match.go:212
		_go_fuzz_dep_.CoverTab[17879]++
//line /usr/local/go/src/path/filepath/match.go:212
		return chunk[0] == ']'
//line /usr/local/go/src/path/filepath/match.go:212
		// _ = "end of CoverTab[17879]"
//line /usr/local/go/src/path/filepath/match.go:212
	}() {
//line /usr/local/go/src/path/filepath/match.go:212
		_go_fuzz_dep_.CoverTab[17880]++
								err = ErrBadPattern
								return
//line /usr/local/go/src/path/filepath/match.go:214
		// _ = "end of CoverTab[17880]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:215
		_go_fuzz_dep_.CoverTab[17881]++
//line /usr/local/go/src/path/filepath/match.go:215
		// _ = "end of CoverTab[17881]"
//line /usr/local/go/src/path/filepath/match.go:215
	}
//line /usr/local/go/src/path/filepath/match.go:215
	// _ = "end of CoverTab[17873]"
//line /usr/local/go/src/path/filepath/match.go:215
	_go_fuzz_dep_.CoverTab[17874]++
							if chunk[0] == '\\' && func() bool {
//line /usr/local/go/src/path/filepath/match.go:216
		_go_fuzz_dep_.CoverTab[17882]++
//line /usr/local/go/src/path/filepath/match.go:216
		return runtime.GOOS != "windows"
//line /usr/local/go/src/path/filepath/match.go:216
		// _ = "end of CoverTab[17882]"
//line /usr/local/go/src/path/filepath/match.go:216
	}() {
//line /usr/local/go/src/path/filepath/match.go:216
		_go_fuzz_dep_.CoverTab[17883]++
								chunk = chunk[1:]
								if len(chunk) == 0 {
//line /usr/local/go/src/path/filepath/match.go:218
			_go_fuzz_dep_.CoverTab[17884]++
									err = ErrBadPattern
									return
//line /usr/local/go/src/path/filepath/match.go:220
			// _ = "end of CoverTab[17884]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:221
			_go_fuzz_dep_.CoverTab[17885]++
//line /usr/local/go/src/path/filepath/match.go:221
			// _ = "end of CoverTab[17885]"
//line /usr/local/go/src/path/filepath/match.go:221
		}
//line /usr/local/go/src/path/filepath/match.go:221
		// _ = "end of CoverTab[17883]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:222
		_go_fuzz_dep_.CoverTab[17886]++
//line /usr/local/go/src/path/filepath/match.go:222
		// _ = "end of CoverTab[17886]"
//line /usr/local/go/src/path/filepath/match.go:222
	}
//line /usr/local/go/src/path/filepath/match.go:222
	// _ = "end of CoverTab[17874]"
//line /usr/local/go/src/path/filepath/match.go:222
	_go_fuzz_dep_.CoverTab[17875]++
							r, n := utf8.DecodeRuneInString(chunk)
							if r == utf8.RuneError && func() bool {
//line /usr/local/go/src/path/filepath/match.go:224
		_go_fuzz_dep_.CoverTab[17887]++
//line /usr/local/go/src/path/filepath/match.go:224
		return n == 1
//line /usr/local/go/src/path/filepath/match.go:224
		// _ = "end of CoverTab[17887]"
//line /usr/local/go/src/path/filepath/match.go:224
	}() {
//line /usr/local/go/src/path/filepath/match.go:224
		_go_fuzz_dep_.CoverTab[17888]++
								err = ErrBadPattern
//line /usr/local/go/src/path/filepath/match.go:225
		// _ = "end of CoverTab[17888]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:226
		_go_fuzz_dep_.CoverTab[17889]++
//line /usr/local/go/src/path/filepath/match.go:226
		// _ = "end of CoverTab[17889]"
//line /usr/local/go/src/path/filepath/match.go:226
	}
//line /usr/local/go/src/path/filepath/match.go:226
	// _ = "end of CoverTab[17875]"
//line /usr/local/go/src/path/filepath/match.go:226
	_go_fuzz_dep_.CoverTab[17876]++
							nchunk = chunk[n:]
							if len(nchunk) == 0 {
//line /usr/local/go/src/path/filepath/match.go:228
		_go_fuzz_dep_.CoverTab[17890]++
								err = ErrBadPattern
//line /usr/local/go/src/path/filepath/match.go:229
		// _ = "end of CoverTab[17890]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:230
		_go_fuzz_dep_.CoverTab[17891]++
//line /usr/local/go/src/path/filepath/match.go:230
		// _ = "end of CoverTab[17891]"
//line /usr/local/go/src/path/filepath/match.go:230
	}
//line /usr/local/go/src/path/filepath/match.go:230
	// _ = "end of CoverTab[17876]"
//line /usr/local/go/src/path/filepath/match.go:230
	_go_fuzz_dep_.CoverTab[17877]++
							return
//line /usr/local/go/src/path/filepath/match.go:231
	// _ = "end of CoverTab[17877]"
}

// Glob returns the names of all files matching pattern or nil
//line /usr/local/go/src/path/filepath/match.go:234
// if there is no matching file. The syntax of patterns is the same
//line /usr/local/go/src/path/filepath/match.go:234
// as in Match. The pattern may describe hierarchical names such as
//line /usr/local/go/src/path/filepath/match.go:234
// /usr/*/bin/ed (assuming the Separator is '/').
//line /usr/local/go/src/path/filepath/match.go:234
//
//line /usr/local/go/src/path/filepath/match.go:234
// Glob ignores file system errors such as I/O errors reading directories.
//line /usr/local/go/src/path/filepath/match.go:234
// The only possible returned error is ErrBadPattern, when pattern
//line /usr/local/go/src/path/filepath/match.go:234
// is malformed.
//line /usr/local/go/src/path/filepath/match.go:242
func Glob(pattern string) (matches []string, err error) {
//line /usr/local/go/src/path/filepath/match.go:242
	_go_fuzz_dep_.CoverTab[17892]++
							return globWithLimit(pattern, 0)
//line /usr/local/go/src/path/filepath/match.go:243
	// _ = "end of CoverTab[17892]"
}

func globWithLimit(pattern string, depth int) (matches []string, err error) {
//line /usr/local/go/src/path/filepath/match.go:246
	_go_fuzz_dep_.CoverTab[17893]++
	// This limit is used prevent stack exhaustion issues. See CVE-2022-30632.
	const pathSeparatorsLimit = 10000
	if depth == pathSeparatorsLimit {
//line /usr/local/go/src/path/filepath/match.go:249
		_go_fuzz_dep_.CoverTab[17902]++
								return nil, ErrBadPattern
//line /usr/local/go/src/path/filepath/match.go:250
		// _ = "end of CoverTab[17902]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:251
		_go_fuzz_dep_.CoverTab[17903]++
//line /usr/local/go/src/path/filepath/match.go:251
		// _ = "end of CoverTab[17903]"
//line /usr/local/go/src/path/filepath/match.go:251
	}
//line /usr/local/go/src/path/filepath/match.go:251
	// _ = "end of CoverTab[17893]"
//line /usr/local/go/src/path/filepath/match.go:251
	_go_fuzz_dep_.CoverTab[17894]++

//line /usr/local/go/src/path/filepath/match.go:254
	if _, err := Match(pattern, ""); err != nil {
//line /usr/local/go/src/path/filepath/match.go:254
		_go_fuzz_dep_.CoverTab[17904]++
								return nil, err
//line /usr/local/go/src/path/filepath/match.go:255
		// _ = "end of CoverTab[17904]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:256
		_go_fuzz_dep_.CoverTab[17905]++
//line /usr/local/go/src/path/filepath/match.go:256
		// _ = "end of CoverTab[17905]"
//line /usr/local/go/src/path/filepath/match.go:256
	}
//line /usr/local/go/src/path/filepath/match.go:256
	// _ = "end of CoverTab[17894]"
//line /usr/local/go/src/path/filepath/match.go:256
	_go_fuzz_dep_.CoverTab[17895]++
							if !hasMeta(pattern) {
//line /usr/local/go/src/path/filepath/match.go:257
		_go_fuzz_dep_.CoverTab[17906]++
								if _, err = os.Lstat(pattern); err != nil {
//line /usr/local/go/src/path/filepath/match.go:258
			_go_fuzz_dep_.CoverTab[17908]++
									return nil, nil
//line /usr/local/go/src/path/filepath/match.go:259
			// _ = "end of CoverTab[17908]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:260
			_go_fuzz_dep_.CoverTab[17909]++
//line /usr/local/go/src/path/filepath/match.go:260
			// _ = "end of CoverTab[17909]"
//line /usr/local/go/src/path/filepath/match.go:260
		}
//line /usr/local/go/src/path/filepath/match.go:260
		// _ = "end of CoverTab[17906]"
//line /usr/local/go/src/path/filepath/match.go:260
		_go_fuzz_dep_.CoverTab[17907]++
								return []string{pattern}, nil
//line /usr/local/go/src/path/filepath/match.go:261
		// _ = "end of CoverTab[17907]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:262
		_go_fuzz_dep_.CoverTab[17910]++
//line /usr/local/go/src/path/filepath/match.go:262
		// _ = "end of CoverTab[17910]"
//line /usr/local/go/src/path/filepath/match.go:262
	}
//line /usr/local/go/src/path/filepath/match.go:262
	// _ = "end of CoverTab[17895]"
//line /usr/local/go/src/path/filepath/match.go:262
	_go_fuzz_dep_.CoverTab[17896]++

							dir, file := Split(pattern)
							volumeLen := 0
							if runtime.GOOS == "windows" {
//line /usr/local/go/src/path/filepath/match.go:266
		_go_fuzz_dep_.CoverTab[17911]++
								volumeLen, dir = cleanGlobPathWindows(dir)
//line /usr/local/go/src/path/filepath/match.go:267
		// _ = "end of CoverTab[17911]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:268
		_go_fuzz_dep_.CoverTab[17912]++
								dir = cleanGlobPath(dir)
//line /usr/local/go/src/path/filepath/match.go:269
		// _ = "end of CoverTab[17912]"
	}
//line /usr/local/go/src/path/filepath/match.go:270
	// _ = "end of CoverTab[17896]"
//line /usr/local/go/src/path/filepath/match.go:270
	_go_fuzz_dep_.CoverTab[17897]++

							if !hasMeta(dir[volumeLen:]) {
//line /usr/local/go/src/path/filepath/match.go:272
		_go_fuzz_dep_.CoverTab[17913]++
								return glob(dir, file, nil)
//line /usr/local/go/src/path/filepath/match.go:273
		// _ = "end of CoverTab[17913]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:274
		_go_fuzz_dep_.CoverTab[17914]++
//line /usr/local/go/src/path/filepath/match.go:274
		// _ = "end of CoverTab[17914]"
//line /usr/local/go/src/path/filepath/match.go:274
	}
//line /usr/local/go/src/path/filepath/match.go:274
	// _ = "end of CoverTab[17897]"
//line /usr/local/go/src/path/filepath/match.go:274
	_go_fuzz_dep_.CoverTab[17898]++

//line /usr/local/go/src/path/filepath/match.go:277
	if dir == pattern {
//line /usr/local/go/src/path/filepath/match.go:277
		_go_fuzz_dep_.CoverTab[17915]++
								return nil, ErrBadPattern
//line /usr/local/go/src/path/filepath/match.go:278
		// _ = "end of CoverTab[17915]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:279
		_go_fuzz_dep_.CoverTab[17916]++
//line /usr/local/go/src/path/filepath/match.go:279
		// _ = "end of CoverTab[17916]"
//line /usr/local/go/src/path/filepath/match.go:279
	}
//line /usr/local/go/src/path/filepath/match.go:279
	// _ = "end of CoverTab[17898]"
//line /usr/local/go/src/path/filepath/match.go:279
	_go_fuzz_dep_.CoverTab[17899]++

							var m []string
							m, err = globWithLimit(dir, depth+1)
							if err != nil {
//line /usr/local/go/src/path/filepath/match.go:283
		_go_fuzz_dep_.CoverTab[17917]++
								return
//line /usr/local/go/src/path/filepath/match.go:284
		// _ = "end of CoverTab[17917]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:285
		_go_fuzz_dep_.CoverTab[17918]++
//line /usr/local/go/src/path/filepath/match.go:285
		// _ = "end of CoverTab[17918]"
//line /usr/local/go/src/path/filepath/match.go:285
	}
//line /usr/local/go/src/path/filepath/match.go:285
	// _ = "end of CoverTab[17899]"
//line /usr/local/go/src/path/filepath/match.go:285
	_go_fuzz_dep_.CoverTab[17900]++
							for _, d := range m {
//line /usr/local/go/src/path/filepath/match.go:286
		_go_fuzz_dep_.CoverTab[17919]++
								matches, err = glob(d, file, matches)
								if err != nil {
//line /usr/local/go/src/path/filepath/match.go:288
			_go_fuzz_dep_.CoverTab[17920]++
									return
//line /usr/local/go/src/path/filepath/match.go:289
			// _ = "end of CoverTab[17920]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:290
			_go_fuzz_dep_.CoverTab[17921]++
//line /usr/local/go/src/path/filepath/match.go:290
			// _ = "end of CoverTab[17921]"
//line /usr/local/go/src/path/filepath/match.go:290
		}
//line /usr/local/go/src/path/filepath/match.go:290
		// _ = "end of CoverTab[17919]"
	}
//line /usr/local/go/src/path/filepath/match.go:291
	// _ = "end of CoverTab[17900]"
//line /usr/local/go/src/path/filepath/match.go:291
	_go_fuzz_dep_.CoverTab[17901]++
							return
//line /usr/local/go/src/path/filepath/match.go:292
	// _ = "end of CoverTab[17901]"
}

// cleanGlobPath prepares path for glob matching.
func cleanGlobPath(path string) string {
//line /usr/local/go/src/path/filepath/match.go:296
	_go_fuzz_dep_.CoverTab[17922]++
							switch path {
	case "":
//line /usr/local/go/src/path/filepath/match.go:298
		_go_fuzz_dep_.CoverTab[17923]++
								return "."
//line /usr/local/go/src/path/filepath/match.go:299
		// _ = "end of CoverTab[17923]"
	case string(Separator):
//line /usr/local/go/src/path/filepath/match.go:300
		_go_fuzz_dep_.CoverTab[17924]++

								return path
//line /usr/local/go/src/path/filepath/match.go:302
		// _ = "end of CoverTab[17924]"
	default:
//line /usr/local/go/src/path/filepath/match.go:303
		_go_fuzz_dep_.CoverTab[17925]++
								return path[0 : len(path)-1]
//line /usr/local/go/src/path/filepath/match.go:304
		// _ = "end of CoverTab[17925]"
	}
//line /usr/local/go/src/path/filepath/match.go:305
	// _ = "end of CoverTab[17922]"
}

// cleanGlobPathWindows is windows version of cleanGlobPath.
func cleanGlobPathWindows(path string) (prefixLen int, cleaned string) {
//line /usr/local/go/src/path/filepath/match.go:309
	_go_fuzz_dep_.CoverTab[17926]++
							vollen := volumeNameLen(path)
							switch {
	case path == "":
//line /usr/local/go/src/path/filepath/match.go:312
		_go_fuzz_dep_.CoverTab[17927]++
								return 0, "."
//line /usr/local/go/src/path/filepath/match.go:313
		// _ = "end of CoverTab[17927]"
	case vollen+1 == len(path) && func() bool {
//line /usr/local/go/src/path/filepath/match.go:314
		_go_fuzz_dep_.CoverTab[17932]++
//line /usr/local/go/src/path/filepath/match.go:314
		return os.IsPathSeparator(path[len(path)-1])
//line /usr/local/go/src/path/filepath/match.go:314
		// _ = "end of CoverTab[17932]"
//line /usr/local/go/src/path/filepath/match.go:314
	}():
//line /usr/local/go/src/path/filepath/match.go:314
		_go_fuzz_dep_.CoverTab[17928]++

								return vollen + 1, path
//line /usr/local/go/src/path/filepath/match.go:316
		// _ = "end of CoverTab[17928]"
	case vollen == len(path) && func() bool {
//line /usr/local/go/src/path/filepath/match.go:317
		_go_fuzz_dep_.CoverTab[17933]++
//line /usr/local/go/src/path/filepath/match.go:317
		return len(path) == 2
//line /usr/local/go/src/path/filepath/match.go:317
		// _ = "end of CoverTab[17933]"
//line /usr/local/go/src/path/filepath/match.go:317
	}():
//line /usr/local/go/src/path/filepath/match.go:317
		_go_fuzz_dep_.CoverTab[17929]++
								return vollen, path + "."
//line /usr/local/go/src/path/filepath/match.go:318
		// _ = "end of CoverTab[17929]"
	default:
//line /usr/local/go/src/path/filepath/match.go:319
		_go_fuzz_dep_.CoverTab[17930]++
								if vollen >= len(path) {
//line /usr/local/go/src/path/filepath/match.go:320
			_go_fuzz_dep_.CoverTab[17934]++
									vollen = len(path) - 1
//line /usr/local/go/src/path/filepath/match.go:321
			// _ = "end of CoverTab[17934]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:322
			_go_fuzz_dep_.CoverTab[17935]++
//line /usr/local/go/src/path/filepath/match.go:322
			// _ = "end of CoverTab[17935]"
//line /usr/local/go/src/path/filepath/match.go:322
		}
//line /usr/local/go/src/path/filepath/match.go:322
		// _ = "end of CoverTab[17930]"
//line /usr/local/go/src/path/filepath/match.go:322
		_go_fuzz_dep_.CoverTab[17931]++
								return vollen, path[0 : len(path)-1]
//line /usr/local/go/src/path/filepath/match.go:323
		// _ = "end of CoverTab[17931]"
	}
//line /usr/local/go/src/path/filepath/match.go:324
	// _ = "end of CoverTab[17926]"
}

// glob searches for files matching pattern in the directory dir
//line /usr/local/go/src/path/filepath/match.go:327
// and appends them to matches. If the directory cannot be
//line /usr/local/go/src/path/filepath/match.go:327
// opened, it returns the existing matches. New matches are
//line /usr/local/go/src/path/filepath/match.go:327
// added in lexicographical order.
//line /usr/local/go/src/path/filepath/match.go:331
func glob(dir, pattern string, matches []string) (m []string, e error) {
//line /usr/local/go/src/path/filepath/match.go:331
	_go_fuzz_dep_.CoverTab[17936]++
							m = matches
							fi, err := os.Stat(dir)
							if err != nil {
//line /usr/local/go/src/path/filepath/match.go:334
		_go_fuzz_dep_.CoverTab[17941]++
								return
//line /usr/local/go/src/path/filepath/match.go:335
		// _ = "end of CoverTab[17941]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:336
		_go_fuzz_dep_.CoverTab[17942]++
//line /usr/local/go/src/path/filepath/match.go:336
		// _ = "end of CoverTab[17942]"
//line /usr/local/go/src/path/filepath/match.go:336
	}
//line /usr/local/go/src/path/filepath/match.go:336
	// _ = "end of CoverTab[17936]"
//line /usr/local/go/src/path/filepath/match.go:336
	_go_fuzz_dep_.CoverTab[17937]++
							if !fi.IsDir() {
//line /usr/local/go/src/path/filepath/match.go:337
		_go_fuzz_dep_.CoverTab[17943]++
								return
//line /usr/local/go/src/path/filepath/match.go:338
		// _ = "end of CoverTab[17943]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:339
		_go_fuzz_dep_.CoverTab[17944]++
//line /usr/local/go/src/path/filepath/match.go:339
		// _ = "end of CoverTab[17944]"
//line /usr/local/go/src/path/filepath/match.go:339
	}
//line /usr/local/go/src/path/filepath/match.go:339
	// _ = "end of CoverTab[17937]"
//line /usr/local/go/src/path/filepath/match.go:339
	_go_fuzz_dep_.CoverTab[17938]++
							d, err := os.Open(dir)
							if err != nil {
//line /usr/local/go/src/path/filepath/match.go:341
		_go_fuzz_dep_.CoverTab[17945]++
								return
//line /usr/local/go/src/path/filepath/match.go:342
		// _ = "end of CoverTab[17945]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:343
		_go_fuzz_dep_.CoverTab[17946]++
//line /usr/local/go/src/path/filepath/match.go:343
		// _ = "end of CoverTab[17946]"
//line /usr/local/go/src/path/filepath/match.go:343
	}
//line /usr/local/go/src/path/filepath/match.go:343
	// _ = "end of CoverTab[17938]"
//line /usr/local/go/src/path/filepath/match.go:343
	_go_fuzz_dep_.CoverTab[17939]++
							defer d.Close()

							names, _ := d.Readdirnames(-1)
							sort.Strings(names)

							for _, n := range names {
//line /usr/local/go/src/path/filepath/match.go:349
		_go_fuzz_dep_.CoverTab[17947]++
								matched, err := Match(pattern, n)
								if err != nil {
//line /usr/local/go/src/path/filepath/match.go:351
			_go_fuzz_dep_.CoverTab[17949]++
									return m, err
//line /usr/local/go/src/path/filepath/match.go:352
			// _ = "end of CoverTab[17949]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:353
			_go_fuzz_dep_.CoverTab[17950]++
//line /usr/local/go/src/path/filepath/match.go:353
			// _ = "end of CoverTab[17950]"
//line /usr/local/go/src/path/filepath/match.go:353
		}
//line /usr/local/go/src/path/filepath/match.go:353
		// _ = "end of CoverTab[17947]"
//line /usr/local/go/src/path/filepath/match.go:353
		_go_fuzz_dep_.CoverTab[17948]++
								if matched {
//line /usr/local/go/src/path/filepath/match.go:354
			_go_fuzz_dep_.CoverTab[17951]++
									m = append(m, Join(dir, n))
//line /usr/local/go/src/path/filepath/match.go:355
			// _ = "end of CoverTab[17951]"
		} else {
//line /usr/local/go/src/path/filepath/match.go:356
			_go_fuzz_dep_.CoverTab[17952]++
//line /usr/local/go/src/path/filepath/match.go:356
			// _ = "end of CoverTab[17952]"
//line /usr/local/go/src/path/filepath/match.go:356
		}
//line /usr/local/go/src/path/filepath/match.go:356
		// _ = "end of CoverTab[17948]"
	}
//line /usr/local/go/src/path/filepath/match.go:357
	// _ = "end of CoverTab[17939]"
//line /usr/local/go/src/path/filepath/match.go:357
	_go_fuzz_dep_.CoverTab[17940]++
							return
//line /usr/local/go/src/path/filepath/match.go:358
	// _ = "end of CoverTab[17940]"
}

// hasMeta reports whether path contains any of the magic characters
//line /usr/local/go/src/path/filepath/match.go:361
// recognized by Match.
//line /usr/local/go/src/path/filepath/match.go:363
func hasMeta(path string) bool {
//line /usr/local/go/src/path/filepath/match.go:363
	_go_fuzz_dep_.CoverTab[17953]++
							magicChars := `*?[`
							if runtime.GOOS != "windows" {
//line /usr/local/go/src/path/filepath/match.go:365
		_go_fuzz_dep_.CoverTab[17955]++
								magicChars = `*?[\`
//line /usr/local/go/src/path/filepath/match.go:366
		// _ = "end of CoverTab[17955]"
	} else {
//line /usr/local/go/src/path/filepath/match.go:367
		_go_fuzz_dep_.CoverTab[17956]++
//line /usr/local/go/src/path/filepath/match.go:367
		// _ = "end of CoverTab[17956]"
//line /usr/local/go/src/path/filepath/match.go:367
	}
//line /usr/local/go/src/path/filepath/match.go:367
	// _ = "end of CoverTab[17953]"
//line /usr/local/go/src/path/filepath/match.go:367
	_go_fuzz_dep_.CoverTab[17954]++
							return strings.ContainsAny(path, magicChars)
//line /usr/local/go/src/path/filepath/match.go:368
	// _ = "end of CoverTab[17954]"
}

//line /usr/local/go/src/path/filepath/match.go:369
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/path/filepath/match.go:369
var _ = _go_fuzz_dep_.CoverTab
