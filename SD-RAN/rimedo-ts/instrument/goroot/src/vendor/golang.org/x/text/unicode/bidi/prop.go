// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:5
package bidi

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:5
)

import "unicode/utf8"

// Properties provides access to BiDi properties of runes.
type Properties struct {
	entry	uint8
	last	uint8
}

var trie = newBidiTrie(0)

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:27
// Class returns the Bidi class for p.
func (p Properties) Class() Class {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:28
	_go_fuzz_dep_.CoverTab[32631]++
										c := Class(p.entry & 0x0F)
										if c == Control {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:30
		_go_fuzz_dep_.CoverTab[32633]++
											c = controlByteToClass[p.last&0xF]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:31
		// _ = "end of CoverTab[32633]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:32
		_go_fuzz_dep_.CoverTab[32634]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:32
		// _ = "end of CoverTab[32634]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:32
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:32
	// _ = "end of CoverTab[32631]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:32
	_go_fuzz_dep_.CoverTab[32632]++
										return c
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:33
	// _ = "end of CoverTab[32632]"
}

// IsBracket reports whether the rune is a bracket.
func (p Properties) IsBracket() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:37
	_go_fuzz_dep_.CoverTab[32635]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:37
	return p.entry&0xF0 != 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:37
	// _ = "end of CoverTab[32635]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:37
}

// IsOpeningBracket reports whether the rune is an opening bracket.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:39
// IsBracket must return true.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:41
func (p Properties) IsOpeningBracket() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:41
	_go_fuzz_dep_.CoverTab[32636]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:41
	return p.entry&openMask != 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:41
	// _ = "end of CoverTab[32636]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:41
}

// TODO: find a better API and expose.
func (p Properties) reverseBracket(r rune) rune {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:44
	_go_fuzz_dep_.CoverTab[32637]++
										return xorMasks[p.entry>>xorMaskShift] ^ r
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:45
	// _ = "end of CoverTab[32637]"
}

var controlByteToClass = [16]Class{
	0xD:	LRO,
	0xE:	RLO,
	0xA:	LRE,
	0xB:	RLE,
	0xC:	PDF,
	0x6:	LRI,
	0x7:	RLI,
	0x8:	FSI,
	0x9:	PDI,
}

// LookupRune returns properties for r.
func LookupRune(r rune) (p Properties, size int) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:61
	_go_fuzz_dep_.CoverTab[32638]++
										var buf [4]byte
										n := utf8.EncodeRune(buf[:], r)
										return Lookup(buf[:n])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:64
	// _ = "end of CoverTab[32638]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:82
// Lookup returns properties for the first rune in s and the width in bytes of
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:82
// its encoding. The size will be 0 if s does not hold enough bytes to complete
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:82
// the encoding.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:85
func Lookup(s []byte) (p Properties, sz int) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:85
	_go_fuzz_dep_.CoverTab[32639]++
										c0 := s[0]
										switch {
	case c0 < 0x80:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:88
		_go_fuzz_dep_.CoverTab[32641]++
											return Properties{entry: bidiValues[c0]}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:89
		// _ = "end of CoverTab[32641]"
	case c0 < 0xC2:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:90
		_go_fuzz_dep_.CoverTab[32642]++
											return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:91
		// _ = "end of CoverTab[32642]"
	case c0 < 0xE0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:92
		_go_fuzz_dep_.CoverTab[32643]++
											if len(s) < 2 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:93
			_go_fuzz_dep_.CoverTab[32656]++
												return Properties{}, 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:94
			// _ = "end of CoverTab[32656]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:95
			_go_fuzz_dep_.CoverTab[32657]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:95
			// _ = "end of CoverTab[32657]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:95
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:95
		// _ = "end of CoverTab[32643]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:95
		_go_fuzz_dep_.CoverTab[32644]++
											i := bidiIndex[c0]
											c1 := s[1]
											if c1 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:98
			_go_fuzz_dep_.CoverTab[32658]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:98
			return 0xC0 <= c1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:98
			// _ = "end of CoverTab[32658]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:98
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:98
			_go_fuzz_dep_.CoverTab[32659]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:99
			// _ = "end of CoverTab[32659]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:100
			_go_fuzz_dep_.CoverTab[32660]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:100
			// _ = "end of CoverTab[32660]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:100
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:100
		// _ = "end of CoverTab[32644]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:100
		_go_fuzz_dep_.CoverTab[32645]++
											return Properties{entry: trie.lookupValue(uint32(i), c1)}, 2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:101
		// _ = "end of CoverTab[32645]"
	case c0 < 0xF0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:102
		_go_fuzz_dep_.CoverTab[32646]++
											if len(s) < 3 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:103
			_go_fuzz_dep_.CoverTab[32661]++
												return Properties{}, 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:104
			// _ = "end of CoverTab[32661]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:105
			_go_fuzz_dep_.CoverTab[32662]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:105
			// _ = "end of CoverTab[32662]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:105
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:105
		// _ = "end of CoverTab[32646]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:105
		_go_fuzz_dep_.CoverTab[32647]++
											i := bidiIndex[c0]
											c1 := s[1]
											if c1 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:108
			_go_fuzz_dep_.CoverTab[32663]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:108
			return 0xC0 <= c1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:108
			// _ = "end of CoverTab[32663]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:108
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:108
			_go_fuzz_dep_.CoverTab[32664]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:109
			// _ = "end of CoverTab[32664]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:110
			_go_fuzz_dep_.CoverTab[32665]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:110
			// _ = "end of CoverTab[32665]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:110
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:110
		// _ = "end of CoverTab[32647]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:110
		_go_fuzz_dep_.CoverTab[32648]++
											o := uint32(i)<<6 + uint32(c1)
											i = bidiIndex[o]
											c2 := s[2]
											if c2 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:114
			_go_fuzz_dep_.CoverTab[32666]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:114
			return 0xC0 <= c2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:114
			// _ = "end of CoverTab[32666]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:114
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:114
			_go_fuzz_dep_.CoverTab[32667]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:115
			// _ = "end of CoverTab[32667]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:116
			_go_fuzz_dep_.CoverTab[32668]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:116
			// _ = "end of CoverTab[32668]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:116
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:116
		// _ = "end of CoverTab[32648]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:116
		_go_fuzz_dep_.CoverTab[32649]++
											return Properties{entry: trie.lookupValue(uint32(i), c2), last: c2}, 3
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:117
		// _ = "end of CoverTab[32649]"
	case c0 < 0xF8:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:118
		_go_fuzz_dep_.CoverTab[32650]++
											if len(s) < 4 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:119
			_go_fuzz_dep_.CoverTab[32669]++
												return Properties{}, 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:120
			// _ = "end of CoverTab[32669]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:121
			_go_fuzz_dep_.CoverTab[32670]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:121
			// _ = "end of CoverTab[32670]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:121
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:121
		// _ = "end of CoverTab[32650]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:121
		_go_fuzz_dep_.CoverTab[32651]++
											i := bidiIndex[c0]
											c1 := s[1]
											if c1 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:124
			_go_fuzz_dep_.CoverTab[32671]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:124
			return 0xC0 <= c1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:124
			// _ = "end of CoverTab[32671]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:124
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:124
			_go_fuzz_dep_.CoverTab[32672]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:125
			// _ = "end of CoverTab[32672]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:126
			_go_fuzz_dep_.CoverTab[32673]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:126
			// _ = "end of CoverTab[32673]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:126
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:126
		// _ = "end of CoverTab[32651]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:126
		_go_fuzz_dep_.CoverTab[32652]++
											o := uint32(i)<<6 + uint32(c1)
											i = bidiIndex[o]
											c2 := s[2]
											if c2 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:130
			_go_fuzz_dep_.CoverTab[32674]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:130
			return 0xC0 <= c2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:130
			// _ = "end of CoverTab[32674]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:130
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:130
			_go_fuzz_dep_.CoverTab[32675]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:131
			// _ = "end of CoverTab[32675]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:132
			_go_fuzz_dep_.CoverTab[32676]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:132
			// _ = "end of CoverTab[32676]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:132
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:132
		// _ = "end of CoverTab[32652]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:132
		_go_fuzz_dep_.CoverTab[32653]++
											o = uint32(i)<<6 + uint32(c2)
											i = bidiIndex[o]
											c3 := s[3]
											if c3 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:136
			_go_fuzz_dep_.CoverTab[32677]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:136
			return 0xC0 <= c3
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:136
			// _ = "end of CoverTab[32677]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:136
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:136
			_go_fuzz_dep_.CoverTab[32678]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:137
			// _ = "end of CoverTab[32678]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:138
			_go_fuzz_dep_.CoverTab[32679]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:138
			// _ = "end of CoverTab[32679]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:138
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:138
		// _ = "end of CoverTab[32653]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:138
		_go_fuzz_dep_.CoverTab[32654]++
											return Properties{entry: trie.lookupValue(uint32(i), c3)}, 4
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:139
		// _ = "end of CoverTab[32654]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:139
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:139
		_go_fuzz_dep_.CoverTab[32655]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:139
		// _ = "end of CoverTab[32655]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:140
	// _ = "end of CoverTab[32639]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:140
	_go_fuzz_dep_.CoverTab[32640]++

										return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:142
	// _ = "end of CoverTab[32640]"
}

// LookupString returns properties for the first rune in s and the width in
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:145
// bytes of its encoding. The size will be 0 if s does not hold enough bytes to
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:145
// complete the encoding.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:148
func LookupString(s string) (p Properties, sz int) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:148
	_go_fuzz_dep_.CoverTab[32680]++
										c0 := s[0]
										switch {
	case c0 < 0x80:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:151
		_go_fuzz_dep_.CoverTab[32682]++
											return Properties{entry: bidiValues[c0]}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:152
		// _ = "end of CoverTab[32682]"
	case c0 < 0xC2:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:153
		_go_fuzz_dep_.CoverTab[32683]++
											return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:154
		// _ = "end of CoverTab[32683]"
	case c0 < 0xE0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:155
		_go_fuzz_dep_.CoverTab[32684]++
											if len(s) < 2 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:156
			_go_fuzz_dep_.CoverTab[32697]++
												return Properties{}, 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:157
			// _ = "end of CoverTab[32697]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:158
			_go_fuzz_dep_.CoverTab[32698]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:158
			// _ = "end of CoverTab[32698]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:158
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:158
		// _ = "end of CoverTab[32684]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:158
		_go_fuzz_dep_.CoverTab[32685]++
											i := bidiIndex[c0]
											c1 := s[1]
											if c1 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:161
			_go_fuzz_dep_.CoverTab[32699]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:161
			return 0xC0 <= c1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:161
			// _ = "end of CoverTab[32699]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:161
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:161
			_go_fuzz_dep_.CoverTab[32700]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:162
			// _ = "end of CoverTab[32700]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:163
			_go_fuzz_dep_.CoverTab[32701]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:163
			// _ = "end of CoverTab[32701]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:163
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:163
		// _ = "end of CoverTab[32685]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:163
		_go_fuzz_dep_.CoverTab[32686]++
											return Properties{entry: trie.lookupValue(uint32(i), c1)}, 2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:164
		// _ = "end of CoverTab[32686]"
	case c0 < 0xF0:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:165
		_go_fuzz_dep_.CoverTab[32687]++
											if len(s) < 3 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:166
			_go_fuzz_dep_.CoverTab[32702]++
												return Properties{}, 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:167
			// _ = "end of CoverTab[32702]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:168
			_go_fuzz_dep_.CoverTab[32703]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:168
			// _ = "end of CoverTab[32703]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:168
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:168
		// _ = "end of CoverTab[32687]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:168
		_go_fuzz_dep_.CoverTab[32688]++
											i := bidiIndex[c0]
											c1 := s[1]
											if c1 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:171
			_go_fuzz_dep_.CoverTab[32704]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:171
			return 0xC0 <= c1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:171
			// _ = "end of CoverTab[32704]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:171
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:171
			_go_fuzz_dep_.CoverTab[32705]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:172
			// _ = "end of CoverTab[32705]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:173
			_go_fuzz_dep_.CoverTab[32706]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:173
			// _ = "end of CoverTab[32706]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:173
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:173
		// _ = "end of CoverTab[32688]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:173
		_go_fuzz_dep_.CoverTab[32689]++
											o := uint32(i)<<6 + uint32(c1)
											i = bidiIndex[o]
											c2 := s[2]
											if c2 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:177
			_go_fuzz_dep_.CoverTab[32707]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:177
			return 0xC0 <= c2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:177
			// _ = "end of CoverTab[32707]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:177
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:177
			_go_fuzz_dep_.CoverTab[32708]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:178
			// _ = "end of CoverTab[32708]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:179
			_go_fuzz_dep_.CoverTab[32709]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:179
			// _ = "end of CoverTab[32709]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:179
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:179
		// _ = "end of CoverTab[32689]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:179
		_go_fuzz_dep_.CoverTab[32690]++
											return Properties{entry: trie.lookupValue(uint32(i), c2), last: c2}, 3
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:180
		// _ = "end of CoverTab[32690]"
	case c0 < 0xF8:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:181
		_go_fuzz_dep_.CoverTab[32691]++
											if len(s) < 4 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:182
			_go_fuzz_dep_.CoverTab[32710]++
												return Properties{}, 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:183
			// _ = "end of CoverTab[32710]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:184
			_go_fuzz_dep_.CoverTab[32711]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:184
			// _ = "end of CoverTab[32711]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:184
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:184
		// _ = "end of CoverTab[32691]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:184
		_go_fuzz_dep_.CoverTab[32692]++
											i := bidiIndex[c0]
											c1 := s[1]
											if c1 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:187
			_go_fuzz_dep_.CoverTab[32712]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:187
			return 0xC0 <= c1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:187
			// _ = "end of CoverTab[32712]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:187
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:187
			_go_fuzz_dep_.CoverTab[32713]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:188
			// _ = "end of CoverTab[32713]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:189
			_go_fuzz_dep_.CoverTab[32714]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:189
			// _ = "end of CoverTab[32714]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:189
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:189
		// _ = "end of CoverTab[32692]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:189
		_go_fuzz_dep_.CoverTab[32693]++
											o := uint32(i)<<6 + uint32(c1)
											i = bidiIndex[o]
											c2 := s[2]
											if c2 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:193
			_go_fuzz_dep_.CoverTab[32715]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:193
			return 0xC0 <= c2
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:193
			// _ = "end of CoverTab[32715]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:193
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:193
			_go_fuzz_dep_.CoverTab[32716]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:194
			// _ = "end of CoverTab[32716]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:195
			_go_fuzz_dep_.CoverTab[32717]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:195
			// _ = "end of CoverTab[32717]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:195
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:195
		// _ = "end of CoverTab[32693]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:195
		_go_fuzz_dep_.CoverTab[32694]++
											o = uint32(i)<<6 + uint32(c2)
											i = bidiIndex[o]
											c3 := s[3]
											if c3 < 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:199
			_go_fuzz_dep_.CoverTab[32718]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:199
			return 0xC0 <= c3
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:199
			// _ = "end of CoverTab[32718]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:199
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:199
			_go_fuzz_dep_.CoverTab[32719]++
												return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:200
			// _ = "end of CoverTab[32719]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:201
			_go_fuzz_dep_.CoverTab[32720]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:201
			// _ = "end of CoverTab[32720]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:201
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:201
		// _ = "end of CoverTab[32694]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:201
		_go_fuzz_dep_.CoverTab[32695]++
											return Properties{entry: trie.lookupValue(uint32(i), c3)}, 4
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:202
		// _ = "end of CoverTab[32695]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:202
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:202
		_go_fuzz_dep_.CoverTab[32696]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:202
		// _ = "end of CoverTab[32696]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:203
	// _ = "end of CoverTab[32680]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:203
	_go_fuzz_dep_.CoverTab[32681]++

										return Properties{}, 1
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:205
	// _ = "end of CoverTab[32681]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:206
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/bidi/prop.go:206
var _ = _go_fuzz_dep_.CoverTab
