// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/escape.go:5
// Package html provides functions for escaping and unescaping HTML text.
package html

//line /usr/local/go/src/html/escape.go:6
import (
//line /usr/local/go/src/html/escape.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/escape.go:6
)
//line /usr/local/go/src/html/escape.go:6
import (
//line /usr/local/go/src/html/escape.go:6
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/escape.go:6
)

import (
	"strings"
	"unicode/utf8"
)

// These replacements permit compatibility with old numeric entities that
//line /usr/local/go/src/html/escape.go:13
// assumed Windows-1252 encoding.
//line /usr/local/go/src/html/escape.go:13
// https://html.spec.whatwg.org/multipage/parsing.html#numeric-character-reference-end-state
//line /usr/local/go/src/html/escape.go:16
var replacementTable = [...]rune{
	'\u20AC',
	'\u0081',
	'\u201A',
	'\u0192',
	'\u201E',
	'\u2026',
	'\u2020',
	'\u2021',
	'\u02C6',
	'\u2030',
	'\u0160',
	'\u2039',
	'\u0152',
	'\u008D',
	'\u017D',
	'\u008F',
	'\u0090',
	'\u2018',
	'\u2019',
	'\u201C',
	'\u201D',
	'\u2022',
	'\u2013',
	'\u2014',
	'\u02DC',
	'\u2122',
	'\u0161',
	'\u203A',
	'\u0153',
	'\u009D',
	'\u017E',
	'\u0178',
//line /usr/local/go/src/html/escape.go:51
}

// unescapeEntity reads an entity like "&lt;" from b[src:] and writes the
//line /usr/local/go/src/html/escape.go:53
// corresponding "<" to b[dst:], returning the incremented dst and src cursors.
//line /usr/local/go/src/html/escape.go:53
// Precondition: b[src] == '&' && dst <= src.
//line /usr/local/go/src/html/escape.go:56
func unescapeEntity(b []byte, dst, src int) (dst1, src1 int) {
//line /usr/local/go/src/html/escape.go:56
	_go_fuzz_dep_.CoverTab[28815]++
						const attribute = false

//line /usr/local/go/src/html/escape.go:62
	i, s := 1, b[src:]

	if len(s) <= 1 {
//line /usr/local/go/src/html/escape.go:64
		_go_fuzz_dep_.CoverTab[28820]++
							b[dst] = b[src]
							return dst + 1, src + 1
//line /usr/local/go/src/html/escape.go:66
		// _ = "end of CoverTab[28820]"
	} else {
//line /usr/local/go/src/html/escape.go:67
		_go_fuzz_dep_.CoverTab[28821]++
//line /usr/local/go/src/html/escape.go:67
		// _ = "end of CoverTab[28821]"
//line /usr/local/go/src/html/escape.go:67
	}
//line /usr/local/go/src/html/escape.go:67
	// _ = "end of CoverTab[28815]"
//line /usr/local/go/src/html/escape.go:67
	_go_fuzz_dep_.CoverTab[28816]++

						if s[i] == '#' {
//line /usr/local/go/src/html/escape.go:69
		_go_fuzz_dep_.CoverTab[28822]++
							if len(s) <= 3 {
//line /usr/local/go/src/html/escape.go:70
			_go_fuzz_dep_.CoverTab[28828]++
								b[dst] = b[src]
								return dst + 1, src + 1
//line /usr/local/go/src/html/escape.go:72
			// _ = "end of CoverTab[28828]"
		} else {
//line /usr/local/go/src/html/escape.go:73
			_go_fuzz_dep_.CoverTab[28829]++
//line /usr/local/go/src/html/escape.go:73
			// _ = "end of CoverTab[28829]"
//line /usr/local/go/src/html/escape.go:73
		}
//line /usr/local/go/src/html/escape.go:73
		// _ = "end of CoverTab[28822]"
//line /usr/local/go/src/html/escape.go:73
		_go_fuzz_dep_.CoverTab[28823]++
							i++
							c := s[i]
							hex := false
							if c == 'x' || func() bool {
//line /usr/local/go/src/html/escape.go:77
			_go_fuzz_dep_.CoverTab[28830]++
//line /usr/local/go/src/html/escape.go:77
			return c == 'X'
//line /usr/local/go/src/html/escape.go:77
			// _ = "end of CoverTab[28830]"
//line /usr/local/go/src/html/escape.go:77
		}() {
//line /usr/local/go/src/html/escape.go:77
			_go_fuzz_dep_.CoverTab[28831]++
								hex = true
								i++
//line /usr/local/go/src/html/escape.go:79
			// _ = "end of CoverTab[28831]"
		} else {
//line /usr/local/go/src/html/escape.go:80
			_go_fuzz_dep_.CoverTab[28832]++
//line /usr/local/go/src/html/escape.go:80
			// _ = "end of CoverTab[28832]"
//line /usr/local/go/src/html/escape.go:80
		}
//line /usr/local/go/src/html/escape.go:80
		// _ = "end of CoverTab[28823]"
//line /usr/local/go/src/html/escape.go:80
		_go_fuzz_dep_.CoverTab[28824]++

							x := '\x00'
							for i < len(s) {
//line /usr/local/go/src/html/escape.go:83
			_go_fuzz_dep_.CoverTab[28833]++
								c = s[i]
								i++
								if hex {
//line /usr/local/go/src/html/escape.go:86
				_go_fuzz_dep_.CoverTab[28836]++
									if '0' <= c && func() bool {
//line /usr/local/go/src/html/escape.go:87
					_go_fuzz_dep_.CoverTab[28837]++
//line /usr/local/go/src/html/escape.go:87
					return c <= '9'
//line /usr/local/go/src/html/escape.go:87
					// _ = "end of CoverTab[28837]"
//line /usr/local/go/src/html/escape.go:87
				}() {
//line /usr/local/go/src/html/escape.go:87
					_go_fuzz_dep_.CoverTab[28838]++
										x = 16*x + rune(c) - '0'
										continue
//line /usr/local/go/src/html/escape.go:89
					// _ = "end of CoverTab[28838]"
				} else {
//line /usr/local/go/src/html/escape.go:90
					_go_fuzz_dep_.CoverTab[28839]++
//line /usr/local/go/src/html/escape.go:90
					if 'a' <= c && func() bool {
//line /usr/local/go/src/html/escape.go:90
						_go_fuzz_dep_.CoverTab[28840]++
//line /usr/local/go/src/html/escape.go:90
						return c <= 'f'
//line /usr/local/go/src/html/escape.go:90
						// _ = "end of CoverTab[28840]"
//line /usr/local/go/src/html/escape.go:90
					}() {
//line /usr/local/go/src/html/escape.go:90
						_go_fuzz_dep_.CoverTab[28841]++
											x = 16*x + rune(c) - 'a' + 10
											continue
//line /usr/local/go/src/html/escape.go:92
						// _ = "end of CoverTab[28841]"
					} else {
//line /usr/local/go/src/html/escape.go:93
						_go_fuzz_dep_.CoverTab[28842]++
//line /usr/local/go/src/html/escape.go:93
						if 'A' <= c && func() bool {
//line /usr/local/go/src/html/escape.go:93
							_go_fuzz_dep_.CoverTab[28843]++
//line /usr/local/go/src/html/escape.go:93
							return c <= 'F'
//line /usr/local/go/src/html/escape.go:93
							// _ = "end of CoverTab[28843]"
//line /usr/local/go/src/html/escape.go:93
						}() {
//line /usr/local/go/src/html/escape.go:93
							_go_fuzz_dep_.CoverTab[28844]++
												x = 16*x + rune(c) - 'A' + 10
												continue
//line /usr/local/go/src/html/escape.go:95
							// _ = "end of CoverTab[28844]"
						} else {
//line /usr/local/go/src/html/escape.go:96
							_go_fuzz_dep_.CoverTab[28845]++
//line /usr/local/go/src/html/escape.go:96
							// _ = "end of CoverTab[28845]"
//line /usr/local/go/src/html/escape.go:96
						}
//line /usr/local/go/src/html/escape.go:96
						// _ = "end of CoverTab[28842]"
//line /usr/local/go/src/html/escape.go:96
					}
//line /usr/local/go/src/html/escape.go:96
					// _ = "end of CoverTab[28839]"
//line /usr/local/go/src/html/escape.go:96
				}
//line /usr/local/go/src/html/escape.go:96
				// _ = "end of CoverTab[28836]"
			} else {
//line /usr/local/go/src/html/escape.go:97
				_go_fuzz_dep_.CoverTab[28846]++
//line /usr/local/go/src/html/escape.go:97
				if '0' <= c && func() bool {
//line /usr/local/go/src/html/escape.go:97
					_go_fuzz_dep_.CoverTab[28847]++
//line /usr/local/go/src/html/escape.go:97
					return c <= '9'
//line /usr/local/go/src/html/escape.go:97
					// _ = "end of CoverTab[28847]"
//line /usr/local/go/src/html/escape.go:97
				}() {
//line /usr/local/go/src/html/escape.go:97
					_go_fuzz_dep_.CoverTab[28848]++
										x = 10*x + rune(c) - '0'
										continue
//line /usr/local/go/src/html/escape.go:99
					// _ = "end of CoverTab[28848]"
				} else {
//line /usr/local/go/src/html/escape.go:100
					_go_fuzz_dep_.CoverTab[28849]++
//line /usr/local/go/src/html/escape.go:100
					// _ = "end of CoverTab[28849]"
//line /usr/local/go/src/html/escape.go:100
				}
//line /usr/local/go/src/html/escape.go:100
				// _ = "end of CoverTab[28846]"
//line /usr/local/go/src/html/escape.go:100
			}
//line /usr/local/go/src/html/escape.go:100
			// _ = "end of CoverTab[28833]"
//line /usr/local/go/src/html/escape.go:100
			_go_fuzz_dep_.CoverTab[28834]++
								if c != ';' {
//line /usr/local/go/src/html/escape.go:101
				_go_fuzz_dep_.CoverTab[28850]++
									i--
//line /usr/local/go/src/html/escape.go:102
				// _ = "end of CoverTab[28850]"
			} else {
//line /usr/local/go/src/html/escape.go:103
				_go_fuzz_dep_.CoverTab[28851]++
//line /usr/local/go/src/html/escape.go:103
				// _ = "end of CoverTab[28851]"
//line /usr/local/go/src/html/escape.go:103
			}
//line /usr/local/go/src/html/escape.go:103
			// _ = "end of CoverTab[28834]"
//line /usr/local/go/src/html/escape.go:103
			_go_fuzz_dep_.CoverTab[28835]++
								break
//line /usr/local/go/src/html/escape.go:104
			// _ = "end of CoverTab[28835]"
		}
//line /usr/local/go/src/html/escape.go:105
		// _ = "end of CoverTab[28824]"
//line /usr/local/go/src/html/escape.go:105
		_go_fuzz_dep_.CoverTab[28825]++

							if i <= 3 {
//line /usr/local/go/src/html/escape.go:107
			_go_fuzz_dep_.CoverTab[28852]++
								b[dst] = b[src]
								return dst + 1, src + 1
//line /usr/local/go/src/html/escape.go:109
			// _ = "end of CoverTab[28852]"
		} else {
//line /usr/local/go/src/html/escape.go:110
			_go_fuzz_dep_.CoverTab[28853]++
//line /usr/local/go/src/html/escape.go:110
			// _ = "end of CoverTab[28853]"
//line /usr/local/go/src/html/escape.go:110
		}
//line /usr/local/go/src/html/escape.go:110
		// _ = "end of CoverTab[28825]"
//line /usr/local/go/src/html/escape.go:110
		_go_fuzz_dep_.CoverTab[28826]++

							if 0x80 <= x && func() bool {
//line /usr/local/go/src/html/escape.go:112
			_go_fuzz_dep_.CoverTab[28854]++
//line /usr/local/go/src/html/escape.go:112
			return x <= 0x9F
//line /usr/local/go/src/html/escape.go:112
			// _ = "end of CoverTab[28854]"
//line /usr/local/go/src/html/escape.go:112
		}() {
//line /usr/local/go/src/html/escape.go:112
			_go_fuzz_dep_.CoverTab[28855]++

								x = replacementTable[x-0x80]
//line /usr/local/go/src/html/escape.go:114
			// _ = "end of CoverTab[28855]"
		} else {
//line /usr/local/go/src/html/escape.go:115
			_go_fuzz_dep_.CoverTab[28856]++
//line /usr/local/go/src/html/escape.go:115
			if x == 0 || func() bool {
//line /usr/local/go/src/html/escape.go:115
				_go_fuzz_dep_.CoverTab[28857]++
//line /usr/local/go/src/html/escape.go:115
				return (0xD800 <= x && func() bool {
//line /usr/local/go/src/html/escape.go:115
					_go_fuzz_dep_.CoverTab[28858]++
//line /usr/local/go/src/html/escape.go:115
					return x <= 0xDFFF
//line /usr/local/go/src/html/escape.go:115
					// _ = "end of CoverTab[28858]"
//line /usr/local/go/src/html/escape.go:115
				}())
//line /usr/local/go/src/html/escape.go:115
				// _ = "end of CoverTab[28857]"
//line /usr/local/go/src/html/escape.go:115
			}() || func() bool {
//line /usr/local/go/src/html/escape.go:115
				_go_fuzz_dep_.CoverTab[28859]++
//line /usr/local/go/src/html/escape.go:115
				return x > 0x10FFFF
//line /usr/local/go/src/html/escape.go:115
				// _ = "end of CoverTab[28859]"
//line /usr/local/go/src/html/escape.go:115
			}() {
//line /usr/local/go/src/html/escape.go:115
				_go_fuzz_dep_.CoverTab[28860]++

									x = '\uFFFD'
//line /usr/local/go/src/html/escape.go:117
				// _ = "end of CoverTab[28860]"
			} else {
//line /usr/local/go/src/html/escape.go:118
				_go_fuzz_dep_.CoverTab[28861]++
//line /usr/local/go/src/html/escape.go:118
				// _ = "end of CoverTab[28861]"
//line /usr/local/go/src/html/escape.go:118
			}
//line /usr/local/go/src/html/escape.go:118
			// _ = "end of CoverTab[28856]"
//line /usr/local/go/src/html/escape.go:118
		}
//line /usr/local/go/src/html/escape.go:118
		// _ = "end of CoverTab[28826]"
//line /usr/local/go/src/html/escape.go:118
		_go_fuzz_dep_.CoverTab[28827]++

							return dst + utf8.EncodeRune(b[dst:], x), src + i
//line /usr/local/go/src/html/escape.go:120
		// _ = "end of CoverTab[28827]"
	} else {
//line /usr/local/go/src/html/escape.go:121
		_go_fuzz_dep_.CoverTab[28862]++
//line /usr/local/go/src/html/escape.go:121
		// _ = "end of CoverTab[28862]"
//line /usr/local/go/src/html/escape.go:121
	}
//line /usr/local/go/src/html/escape.go:121
	// _ = "end of CoverTab[28816]"
//line /usr/local/go/src/html/escape.go:121
	_go_fuzz_dep_.CoverTab[28817]++

//line /usr/local/go/src/html/escape.go:126
	for i < len(s) {
//line /usr/local/go/src/html/escape.go:126
		_go_fuzz_dep_.CoverTab[28863]++
							c := s[i]
							i++

							if 'a' <= c && func() bool {
//line /usr/local/go/src/html/escape.go:130
			_go_fuzz_dep_.CoverTab[28866]++
//line /usr/local/go/src/html/escape.go:130
			return c <= 'z'
//line /usr/local/go/src/html/escape.go:130
			// _ = "end of CoverTab[28866]"
//line /usr/local/go/src/html/escape.go:130
		}() || func() bool {
//line /usr/local/go/src/html/escape.go:130
			_go_fuzz_dep_.CoverTab[28867]++
//line /usr/local/go/src/html/escape.go:130
			return 'A' <= c && func() bool {
//line /usr/local/go/src/html/escape.go:130
				_go_fuzz_dep_.CoverTab[28868]++
//line /usr/local/go/src/html/escape.go:130
				return c <= 'Z'
//line /usr/local/go/src/html/escape.go:130
				// _ = "end of CoverTab[28868]"
//line /usr/local/go/src/html/escape.go:130
			}()
//line /usr/local/go/src/html/escape.go:130
			// _ = "end of CoverTab[28867]"
//line /usr/local/go/src/html/escape.go:130
		}() || func() bool {
//line /usr/local/go/src/html/escape.go:130
			_go_fuzz_dep_.CoverTab[28869]++
//line /usr/local/go/src/html/escape.go:130
			return '0' <= c && func() bool {
//line /usr/local/go/src/html/escape.go:130
				_go_fuzz_dep_.CoverTab[28870]++
//line /usr/local/go/src/html/escape.go:130
				return c <= '9'
//line /usr/local/go/src/html/escape.go:130
				// _ = "end of CoverTab[28870]"
//line /usr/local/go/src/html/escape.go:130
			}()
//line /usr/local/go/src/html/escape.go:130
			// _ = "end of CoverTab[28869]"
//line /usr/local/go/src/html/escape.go:130
		}() {
//line /usr/local/go/src/html/escape.go:130
			_go_fuzz_dep_.CoverTab[28871]++
								continue
//line /usr/local/go/src/html/escape.go:131
			// _ = "end of CoverTab[28871]"
		} else {
//line /usr/local/go/src/html/escape.go:132
			_go_fuzz_dep_.CoverTab[28872]++
//line /usr/local/go/src/html/escape.go:132
			// _ = "end of CoverTab[28872]"
//line /usr/local/go/src/html/escape.go:132
		}
//line /usr/local/go/src/html/escape.go:132
		// _ = "end of CoverTab[28863]"
//line /usr/local/go/src/html/escape.go:132
		_go_fuzz_dep_.CoverTab[28864]++
							if c != ';' {
//line /usr/local/go/src/html/escape.go:133
			_go_fuzz_dep_.CoverTab[28873]++
								i--
//line /usr/local/go/src/html/escape.go:134
			// _ = "end of CoverTab[28873]"
		} else {
//line /usr/local/go/src/html/escape.go:135
			_go_fuzz_dep_.CoverTab[28874]++
//line /usr/local/go/src/html/escape.go:135
			// _ = "end of CoverTab[28874]"
//line /usr/local/go/src/html/escape.go:135
		}
//line /usr/local/go/src/html/escape.go:135
		// _ = "end of CoverTab[28864]"
//line /usr/local/go/src/html/escape.go:135
		_go_fuzz_dep_.CoverTab[28865]++
							break
//line /usr/local/go/src/html/escape.go:136
		// _ = "end of CoverTab[28865]"
	}
//line /usr/local/go/src/html/escape.go:137
	// _ = "end of CoverTab[28817]"
//line /usr/local/go/src/html/escape.go:137
	_go_fuzz_dep_.CoverTab[28818]++

						entityName := s[1:i]
						if len(entityName) == 0 {
//line /usr/local/go/src/html/escape.go:140
		_go_fuzz_dep_.CoverTab[28875]++
//line /usr/local/go/src/html/escape.go:140
		// _ = "end of CoverTab[28875]"

	} else {
//line /usr/local/go/src/html/escape.go:142
		_go_fuzz_dep_.CoverTab[28876]++
//line /usr/local/go/src/html/escape.go:142
		if attribute && func() bool {
//line /usr/local/go/src/html/escape.go:142
			_go_fuzz_dep_.CoverTab[28877]++
//line /usr/local/go/src/html/escape.go:142
			return entityName[len(entityName)-1] != ';'
//line /usr/local/go/src/html/escape.go:142
			// _ = "end of CoverTab[28877]"
//line /usr/local/go/src/html/escape.go:142
		}() && func() bool {
//line /usr/local/go/src/html/escape.go:142
			_go_fuzz_dep_.CoverTab[28878]++
//line /usr/local/go/src/html/escape.go:142
			return len(s) > i
//line /usr/local/go/src/html/escape.go:142
			// _ = "end of CoverTab[28878]"
//line /usr/local/go/src/html/escape.go:142
		}() && func() bool {
//line /usr/local/go/src/html/escape.go:142
			_go_fuzz_dep_.CoverTab[28879]++
//line /usr/local/go/src/html/escape.go:142
			return s[i] == '='
//line /usr/local/go/src/html/escape.go:142
			// _ = "end of CoverTab[28879]"
//line /usr/local/go/src/html/escape.go:142
		}() {
//line /usr/local/go/src/html/escape.go:142
			_go_fuzz_dep_.CoverTab[28880]++
//line /usr/local/go/src/html/escape.go:142
			// _ = "end of CoverTab[28880]"

		} else {
//line /usr/local/go/src/html/escape.go:144
			_go_fuzz_dep_.CoverTab[28881]++
//line /usr/local/go/src/html/escape.go:144
			if x := entity[string(entityName)]; x != 0 {
//line /usr/local/go/src/html/escape.go:144
				_go_fuzz_dep_.CoverTab[28882]++
									return dst + utf8.EncodeRune(b[dst:], x), src + i
//line /usr/local/go/src/html/escape.go:145
				// _ = "end of CoverTab[28882]"
			} else {
//line /usr/local/go/src/html/escape.go:146
				_go_fuzz_dep_.CoverTab[28883]++
//line /usr/local/go/src/html/escape.go:146
				if x := entity2[string(entityName)]; x[0] != 0 {
//line /usr/local/go/src/html/escape.go:146
					_go_fuzz_dep_.CoverTab[28884]++
										dst1 := dst + utf8.EncodeRune(b[dst:], x[0])
										return dst1 + utf8.EncodeRune(b[dst1:], x[1]), src + i
//line /usr/local/go/src/html/escape.go:148
					// _ = "end of CoverTab[28884]"
				} else {
//line /usr/local/go/src/html/escape.go:149
					_go_fuzz_dep_.CoverTab[28885]++
//line /usr/local/go/src/html/escape.go:149
					if !attribute {
//line /usr/local/go/src/html/escape.go:149
						_go_fuzz_dep_.CoverTab[28886]++
											maxLen := len(entityName) - 1
											if maxLen > longestEntityWithoutSemicolon {
//line /usr/local/go/src/html/escape.go:151
							_go_fuzz_dep_.CoverTab[28888]++
												maxLen = longestEntityWithoutSemicolon
//line /usr/local/go/src/html/escape.go:152
							// _ = "end of CoverTab[28888]"
						} else {
//line /usr/local/go/src/html/escape.go:153
							_go_fuzz_dep_.CoverTab[28889]++
//line /usr/local/go/src/html/escape.go:153
							// _ = "end of CoverTab[28889]"
//line /usr/local/go/src/html/escape.go:153
						}
//line /usr/local/go/src/html/escape.go:153
						// _ = "end of CoverTab[28886]"
//line /usr/local/go/src/html/escape.go:153
						_go_fuzz_dep_.CoverTab[28887]++
											for j := maxLen; j > 1; j-- {
//line /usr/local/go/src/html/escape.go:154
							_go_fuzz_dep_.CoverTab[28890]++
												if x := entity[string(entityName[:j])]; x != 0 {
//line /usr/local/go/src/html/escape.go:155
								_go_fuzz_dep_.CoverTab[28891]++
													return dst + utf8.EncodeRune(b[dst:], x), src + j + 1
//line /usr/local/go/src/html/escape.go:156
								// _ = "end of CoverTab[28891]"
							} else {
//line /usr/local/go/src/html/escape.go:157
								_go_fuzz_dep_.CoverTab[28892]++
//line /usr/local/go/src/html/escape.go:157
								// _ = "end of CoverTab[28892]"
//line /usr/local/go/src/html/escape.go:157
							}
//line /usr/local/go/src/html/escape.go:157
							// _ = "end of CoverTab[28890]"
						}
//line /usr/local/go/src/html/escape.go:158
						// _ = "end of CoverTab[28887]"
					} else {
//line /usr/local/go/src/html/escape.go:159
						_go_fuzz_dep_.CoverTab[28893]++
//line /usr/local/go/src/html/escape.go:159
						// _ = "end of CoverTab[28893]"
//line /usr/local/go/src/html/escape.go:159
					}
//line /usr/local/go/src/html/escape.go:159
					// _ = "end of CoverTab[28885]"
//line /usr/local/go/src/html/escape.go:159
				}
//line /usr/local/go/src/html/escape.go:159
				// _ = "end of CoverTab[28883]"
//line /usr/local/go/src/html/escape.go:159
			}
//line /usr/local/go/src/html/escape.go:159
			// _ = "end of CoverTab[28881]"
//line /usr/local/go/src/html/escape.go:159
		}
//line /usr/local/go/src/html/escape.go:159
		// _ = "end of CoverTab[28876]"
//line /usr/local/go/src/html/escape.go:159
	}
//line /usr/local/go/src/html/escape.go:159
	// _ = "end of CoverTab[28818]"
//line /usr/local/go/src/html/escape.go:159
	_go_fuzz_dep_.CoverTab[28819]++

						dst1, src1 = dst+i, src+i
						copy(b[dst:dst1], b[src:src1])
						return dst1, src1
//line /usr/local/go/src/html/escape.go:163
	// _ = "end of CoverTab[28819]"
}

var htmlEscaper = strings.NewReplacer(
	`&`, "&amp;",
	`'`, "&#39;",
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&#34;",
)

// EscapeString escapes special characters like "<" to become "&lt;". It
//line /usr/local/go/src/html/escape.go:174
// escapes only five such characters: <, >, &, ' and ".
//line /usr/local/go/src/html/escape.go:174
// UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
//line /usr/local/go/src/html/escape.go:174
// always true.
//line /usr/local/go/src/html/escape.go:178
func EscapeString(s string) string {
//line /usr/local/go/src/html/escape.go:178
	_go_fuzz_dep_.CoverTab[28894]++
						return htmlEscaper.Replace(s)
//line /usr/local/go/src/html/escape.go:179
	// _ = "end of CoverTab[28894]"
}

// UnescapeString unescapes entities like "&lt;" to become "<". It unescapes a
//line /usr/local/go/src/html/escape.go:182
// larger range of entities than EscapeString escapes. For example, "&aacute;"
//line /usr/local/go/src/html/escape.go:182
// unescapes to "á", as does "&#225;" and "&#xE1;".
//line /usr/local/go/src/html/escape.go:182
// UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
//line /usr/local/go/src/html/escape.go:182
// always true.
//line /usr/local/go/src/html/escape.go:187
func UnescapeString(s string) string {
//line /usr/local/go/src/html/escape.go:187
	_go_fuzz_dep_.CoverTab[28895]++
						populateMapsOnce.Do(populateMaps)
						i := strings.IndexByte(s, '&')

						if i < 0 {
//line /usr/local/go/src/html/escape.go:191
		_go_fuzz_dep_.CoverTab[28898]++
							return s
//line /usr/local/go/src/html/escape.go:192
		// _ = "end of CoverTab[28898]"
	} else {
//line /usr/local/go/src/html/escape.go:193
		_go_fuzz_dep_.CoverTab[28899]++
//line /usr/local/go/src/html/escape.go:193
		// _ = "end of CoverTab[28899]"
//line /usr/local/go/src/html/escape.go:193
	}
//line /usr/local/go/src/html/escape.go:193
	// _ = "end of CoverTab[28895]"
//line /usr/local/go/src/html/escape.go:193
	_go_fuzz_dep_.CoverTab[28896]++

						b := []byte(s)
						dst, src := unescapeEntity(b, i, i)
						for len(s[src:]) > 0 {
//line /usr/local/go/src/html/escape.go:197
		_go_fuzz_dep_.CoverTab[28900]++
							if s[src] == '&' {
//line /usr/local/go/src/html/escape.go:198
			_go_fuzz_dep_.CoverTab[28904]++
								i = 0
//line /usr/local/go/src/html/escape.go:199
			// _ = "end of CoverTab[28904]"
		} else {
//line /usr/local/go/src/html/escape.go:200
			_go_fuzz_dep_.CoverTab[28905]++
								i = strings.IndexByte(s[src:], '&')
//line /usr/local/go/src/html/escape.go:201
			// _ = "end of CoverTab[28905]"
		}
//line /usr/local/go/src/html/escape.go:202
		// _ = "end of CoverTab[28900]"
//line /usr/local/go/src/html/escape.go:202
		_go_fuzz_dep_.CoverTab[28901]++
							if i < 0 {
//line /usr/local/go/src/html/escape.go:203
			_go_fuzz_dep_.CoverTab[28906]++
								dst += copy(b[dst:], s[src:])
								break
//line /usr/local/go/src/html/escape.go:205
			// _ = "end of CoverTab[28906]"
		} else {
//line /usr/local/go/src/html/escape.go:206
			_go_fuzz_dep_.CoverTab[28907]++
//line /usr/local/go/src/html/escape.go:206
			// _ = "end of CoverTab[28907]"
//line /usr/local/go/src/html/escape.go:206
		}
//line /usr/local/go/src/html/escape.go:206
		// _ = "end of CoverTab[28901]"
//line /usr/local/go/src/html/escape.go:206
		_go_fuzz_dep_.CoverTab[28902]++

							if i > 0 {
//line /usr/local/go/src/html/escape.go:208
			_go_fuzz_dep_.CoverTab[28908]++
								copy(b[dst:], s[src:src+i])
//line /usr/local/go/src/html/escape.go:209
			// _ = "end of CoverTab[28908]"
		} else {
//line /usr/local/go/src/html/escape.go:210
			_go_fuzz_dep_.CoverTab[28909]++
//line /usr/local/go/src/html/escape.go:210
			// _ = "end of CoverTab[28909]"
//line /usr/local/go/src/html/escape.go:210
		}
//line /usr/local/go/src/html/escape.go:210
		// _ = "end of CoverTab[28902]"
//line /usr/local/go/src/html/escape.go:210
		_go_fuzz_dep_.CoverTab[28903]++
							dst, src = unescapeEntity(b, dst+i, src+i)
//line /usr/local/go/src/html/escape.go:211
		// _ = "end of CoverTab[28903]"
	}
//line /usr/local/go/src/html/escape.go:212
	// _ = "end of CoverTab[28896]"
//line /usr/local/go/src/html/escape.go:212
	_go_fuzz_dep_.CoverTab[28897]++
						return string(b[:dst])
//line /usr/local/go/src/html/escape.go:213
	// _ = "end of CoverTab[28897]"
}

//line /usr/local/go/src/html/escape.go:214
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/escape.go:214
var _ = _go_fuzz_dep_.CoverTab
