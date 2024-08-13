// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:5
package bidi

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:5
)

import "unicode/utf8"

// Properties provides access to BiDi properties of runes.
type Properties struct {
	entry	uint8
	last	uint8
}

var trie = newBidiTrie(0)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:27
// Class returns the Bidi class for p.
func (p Properties) Class() Class {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:28
	_go_fuzz_dep_.CoverTab[69912]++
											c := Class(p.entry & 0x0F)
											if c == Control {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:30
		_go_fuzz_dep_.CoverTab[69914]++
												c = controlByteToClass[p.last&0xF]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:31
		// _ = "end of CoverTab[69914]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:32
		_go_fuzz_dep_.CoverTab[69915]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:32
		// _ = "end of CoverTab[69915]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:32
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:32
	// _ = "end of CoverTab[69912]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:32
	_go_fuzz_dep_.CoverTab[69913]++
											return c
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:33
	// _ = "end of CoverTab[69913]"
}

// IsBracket reports whether the rune is a bracket.
func (p Properties) IsBracket() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:37
	_go_fuzz_dep_.CoverTab[69916]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:37
	return p.entry&0xF0 != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:37
	// _ = "end of CoverTab[69916]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:37
}

// IsOpeningBracket reports whether the rune is an opening bracket.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:39
// IsBracket must return true.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:41
func (p Properties) IsOpeningBracket() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:41
	_go_fuzz_dep_.CoverTab[69917]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:41
	return p.entry&openMask != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:41
	// _ = "end of CoverTab[69917]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:41
}

// TODO: find a better API and expose.
func (p Properties) reverseBracket(r rune) rune {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:44
	_go_fuzz_dep_.CoverTab[69918]++
											return xorMasks[p.entry>>xorMaskShift] ^ r
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:45
	// _ = "end of CoverTab[69918]"
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
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:61
	_go_fuzz_dep_.CoverTab[69919]++
											var buf [4]byte
											n := utf8.EncodeRune(buf[:], r)
											return Lookup(buf[:n])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:64
	// _ = "end of CoverTab[69919]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:82
// Lookup returns properties for the first rune in s and the width in bytes of
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:82
// its encoding. The size will be 0 if s does not hold enough bytes to complete
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:82
// the encoding.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:85
func Lookup(s []byte) (p Properties, sz int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:85
	_go_fuzz_dep_.CoverTab[69920]++
											c0 := s[0]
											switch {
	case c0 < 0x80:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:88
		_go_fuzz_dep_.CoverTab[69922]++
												return Properties{entry: bidiValues[c0]}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:89
		// _ = "end of CoverTab[69922]"
	case c0 < 0xC2:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:90
		_go_fuzz_dep_.CoverTab[69923]++
												return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:91
		// _ = "end of CoverTab[69923]"
	case c0 < 0xE0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:92
		_go_fuzz_dep_.CoverTab[69924]++
												if len(s) < 2 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:93
			_go_fuzz_dep_.CoverTab[69937]++
													return Properties{}, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:94
			// _ = "end of CoverTab[69937]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:95
			_go_fuzz_dep_.CoverTab[69938]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:95
			// _ = "end of CoverTab[69938]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:95
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:95
		// _ = "end of CoverTab[69924]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:95
		_go_fuzz_dep_.CoverTab[69925]++
												i := bidiIndex[c0]
												c1 := s[1]
												if c1 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:98
			_go_fuzz_dep_.CoverTab[69939]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:98
			return 0xC0 <= c1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:98
			// _ = "end of CoverTab[69939]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:98
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:98
			_go_fuzz_dep_.CoverTab[69940]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:99
			// _ = "end of CoverTab[69940]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:100
			_go_fuzz_dep_.CoverTab[69941]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:100
			// _ = "end of CoverTab[69941]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:100
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:100
		// _ = "end of CoverTab[69925]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:100
		_go_fuzz_dep_.CoverTab[69926]++
												return Properties{entry: trie.lookupValue(uint32(i), c1)}, 2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:101
		// _ = "end of CoverTab[69926]"
	case c0 < 0xF0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:102
		_go_fuzz_dep_.CoverTab[69927]++
												if len(s) < 3 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:103
			_go_fuzz_dep_.CoverTab[69942]++
													return Properties{}, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:104
			// _ = "end of CoverTab[69942]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:105
			_go_fuzz_dep_.CoverTab[69943]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:105
			// _ = "end of CoverTab[69943]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:105
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:105
		// _ = "end of CoverTab[69927]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:105
		_go_fuzz_dep_.CoverTab[69928]++
												i := bidiIndex[c0]
												c1 := s[1]
												if c1 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:108
			_go_fuzz_dep_.CoverTab[69944]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:108
			return 0xC0 <= c1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:108
			// _ = "end of CoverTab[69944]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:108
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:108
			_go_fuzz_dep_.CoverTab[69945]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:109
			// _ = "end of CoverTab[69945]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:110
			_go_fuzz_dep_.CoverTab[69946]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:110
			// _ = "end of CoverTab[69946]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:110
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:110
		// _ = "end of CoverTab[69928]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:110
		_go_fuzz_dep_.CoverTab[69929]++
												o := uint32(i)<<6 + uint32(c1)
												i = bidiIndex[o]
												c2 := s[2]
												if c2 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:114
			_go_fuzz_dep_.CoverTab[69947]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:114
			return 0xC0 <= c2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:114
			// _ = "end of CoverTab[69947]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:114
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:114
			_go_fuzz_dep_.CoverTab[69948]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:115
			// _ = "end of CoverTab[69948]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:116
			_go_fuzz_dep_.CoverTab[69949]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:116
			// _ = "end of CoverTab[69949]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:116
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:116
		// _ = "end of CoverTab[69929]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:116
		_go_fuzz_dep_.CoverTab[69930]++
												return Properties{entry: trie.lookupValue(uint32(i), c2), last: c2}, 3
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:117
		// _ = "end of CoverTab[69930]"
	case c0 < 0xF8:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:118
		_go_fuzz_dep_.CoverTab[69931]++
												if len(s) < 4 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:119
			_go_fuzz_dep_.CoverTab[69950]++
													return Properties{}, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:120
			// _ = "end of CoverTab[69950]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:121
			_go_fuzz_dep_.CoverTab[69951]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:121
			// _ = "end of CoverTab[69951]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:121
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:121
		// _ = "end of CoverTab[69931]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:121
		_go_fuzz_dep_.CoverTab[69932]++
												i := bidiIndex[c0]
												c1 := s[1]
												if c1 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:124
			_go_fuzz_dep_.CoverTab[69952]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:124
			return 0xC0 <= c1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:124
			// _ = "end of CoverTab[69952]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:124
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:124
			_go_fuzz_dep_.CoverTab[69953]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:125
			// _ = "end of CoverTab[69953]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:126
			_go_fuzz_dep_.CoverTab[69954]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:126
			// _ = "end of CoverTab[69954]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:126
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:126
		// _ = "end of CoverTab[69932]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:126
		_go_fuzz_dep_.CoverTab[69933]++
												o := uint32(i)<<6 + uint32(c1)
												i = bidiIndex[o]
												c2 := s[2]
												if c2 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:130
			_go_fuzz_dep_.CoverTab[69955]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:130
			return 0xC0 <= c2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:130
			// _ = "end of CoverTab[69955]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:130
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:130
			_go_fuzz_dep_.CoverTab[69956]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:131
			// _ = "end of CoverTab[69956]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:132
			_go_fuzz_dep_.CoverTab[69957]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:132
			// _ = "end of CoverTab[69957]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:132
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:132
		// _ = "end of CoverTab[69933]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:132
		_go_fuzz_dep_.CoverTab[69934]++
												o = uint32(i)<<6 + uint32(c2)
												i = bidiIndex[o]
												c3 := s[3]
												if c3 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:136
			_go_fuzz_dep_.CoverTab[69958]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:136
			return 0xC0 <= c3
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:136
			// _ = "end of CoverTab[69958]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:136
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:136
			_go_fuzz_dep_.CoverTab[69959]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:137
			// _ = "end of CoverTab[69959]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:138
			_go_fuzz_dep_.CoverTab[69960]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:138
			// _ = "end of CoverTab[69960]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:138
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:138
		// _ = "end of CoverTab[69934]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:138
		_go_fuzz_dep_.CoverTab[69935]++
												return Properties{entry: trie.lookupValue(uint32(i), c3)}, 4
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:139
		// _ = "end of CoverTab[69935]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:139
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:139
		_go_fuzz_dep_.CoverTab[69936]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:139
		// _ = "end of CoverTab[69936]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:140
	// _ = "end of CoverTab[69920]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:140
	_go_fuzz_dep_.CoverTab[69921]++

											return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:142
	// _ = "end of CoverTab[69921]"
}

// LookupString returns properties for the first rune in s and the width in
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:145
// bytes of its encoding. The size will be 0 if s does not hold enough bytes to
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:145
// complete the encoding.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:148
func LookupString(s string) (p Properties, sz int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:148
	_go_fuzz_dep_.CoverTab[69961]++
											c0 := s[0]
											switch {
	case c0 < 0x80:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:151
		_go_fuzz_dep_.CoverTab[69963]++
												return Properties{entry: bidiValues[c0]}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:152
		// _ = "end of CoverTab[69963]"
	case c0 < 0xC2:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:153
		_go_fuzz_dep_.CoverTab[69964]++
												return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:154
		// _ = "end of CoverTab[69964]"
	case c0 < 0xE0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:155
		_go_fuzz_dep_.CoverTab[69965]++
												if len(s) < 2 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:156
			_go_fuzz_dep_.CoverTab[69978]++
													return Properties{}, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:157
			// _ = "end of CoverTab[69978]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:158
			_go_fuzz_dep_.CoverTab[69979]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:158
			// _ = "end of CoverTab[69979]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:158
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:158
		// _ = "end of CoverTab[69965]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:158
		_go_fuzz_dep_.CoverTab[69966]++
												i := bidiIndex[c0]
												c1 := s[1]
												if c1 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:161
			_go_fuzz_dep_.CoverTab[69980]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:161
			return 0xC0 <= c1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:161
			// _ = "end of CoverTab[69980]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:161
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:161
			_go_fuzz_dep_.CoverTab[69981]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:162
			// _ = "end of CoverTab[69981]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:163
			_go_fuzz_dep_.CoverTab[69982]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:163
			// _ = "end of CoverTab[69982]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:163
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:163
		// _ = "end of CoverTab[69966]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:163
		_go_fuzz_dep_.CoverTab[69967]++
												return Properties{entry: trie.lookupValue(uint32(i), c1)}, 2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:164
		// _ = "end of CoverTab[69967]"
	case c0 < 0xF0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:165
		_go_fuzz_dep_.CoverTab[69968]++
												if len(s) < 3 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:166
			_go_fuzz_dep_.CoverTab[69983]++
													return Properties{}, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:167
			// _ = "end of CoverTab[69983]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:168
			_go_fuzz_dep_.CoverTab[69984]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:168
			// _ = "end of CoverTab[69984]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:168
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:168
		// _ = "end of CoverTab[69968]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:168
		_go_fuzz_dep_.CoverTab[69969]++
												i := bidiIndex[c0]
												c1 := s[1]
												if c1 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:171
			_go_fuzz_dep_.CoverTab[69985]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:171
			return 0xC0 <= c1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:171
			// _ = "end of CoverTab[69985]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:171
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:171
			_go_fuzz_dep_.CoverTab[69986]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:172
			// _ = "end of CoverTab[69986]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:173
			_go_fuzz_dep_.CoverTab[69987]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:173
			// _ = "end of CoverTab[69987]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:173
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:173
		// _ = "end of CoverTab[69969]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:173
		_go_fuzz_dep_.CoverTab[69970]++
												o := uint32(i)<<6 + uint32(c1)
												i = bidiIndex[o]
												c2 := s[2]
												if c2 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:177
			_go_fuzz_dep_.CoverTab[69988]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:177
			return 0xC0 <= c2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:177
			// _ = "end of CoverTab[69988]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:177
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:177
			_go_fuzz_dep_.CoverTab[69989]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:178
			// _ = "end of CoverTab[69989]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:179
			_go_fuzz_dep_.CoverTab[69990]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:179
			// _ = "end of CoverTab[69990]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:179
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:179
		// _ = "end of CoverTab[69970]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:179
		_go_fuzz_dep_.CoverTab[69971]++
												return Properties{entry: trie.lookupValue(uint32(i), c2), last: c2}, 3
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:180
		// _ = "end of CoverTab[69971]"
	case c0 < 0xF8:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:181
		_go_fuzz_dep_.CoverTab[69972]++
												if len(s) < 4 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:182
			_go_fuzz_dep_.CoverTab[69991]++
													return Properties{}, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:183
			// _ = "end of CoverTab[69991]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:184
			_go_fuzz_dep_.CoverTab[69992]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:184
			// _ = "end of CoverTab[69992]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:184
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:184
		// _ = "end of CoverTab[69972]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:184
		_go_fuzz_dep_.CoverTab[69973]++
												i := bidiIndex[c0]
												c1 := s[1]
												if c1 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:187
			_go_fuzz_dep_.CoverTab[69993]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:187
			return 0xC0 <= c1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:187
			// _ = "end of CoverTab[69993]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:187
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:187
			_go_fuzz_dep_.CoverTab[69994]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:188
			// _ = "end of CoverTab[69994]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:189
			_go_fuzz_dep_.CoverTab[69995]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:189
			// _ = "end of CoverTab[69995]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:189
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:189
		// _ = "end of CoverTab[69973]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:189
		_go_fuzz_dep_.CoverTab[69974]++
												o := uint32(i)<<6 + uint32(c1)
												i = bidiIndex[o]
												c2 := s[2]
												if c2 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:193
			_go_fuzz_dep_.CoverTab[69996]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:193
			return 0xC0 <= c2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:193
			// _ = "end of CoverTab[69996]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:193
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:193
			_go_fuzz_dep_.CoverTab[69997]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:194
			// _ = "end of CoverTab[69997]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:195
			_go_fuzz_dep_.CoverTab[69998]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:195
			// _ = "end of CoverTab[69998]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:195
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:195
		// _ = "end of CoverTab[69974]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:195
		_go_fuzz_dep_.CoverTab[69975]++
												o = uint32(i)<<6 + uint32(c2)
												i = bidiIndex[o]
												c3 := s[3]
												if c3 < 0x80 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:199
			_go_fuzz_dep_.CoverTab[69999]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:199
			return 0xC0 <= c3
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:199
			// _ = "end of CoverTab[69999]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:199
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:199
			_go_fuzz_dep_.CoverTab[70000]++
													return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:200
			// _ = "end of CoverTab[70000]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:201
			_go_fuzz_dep_.CoverTab[70001]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:201
			// _ = "end of CoverTab[70001]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:201
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:201
		// _ = "end of CoverTab[69975]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:201
		_go_fuzz_dep_.CoverTab[69976]++
												return Properties{entry: trie.lookupValue(uint32(i), c3)}, 4
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:202
		// _ = "end of CoverTab[69976]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:202
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:202
		_go_fuzz_dep_.CoverTab[69977]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:202
		// _ = "end of CoverTab[69977]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:203
	// _ = "end of CoverTab[69961]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:203
	_go_fuzz_dep_.CoverTab[69962]++

											return Properties{}, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:205
	// _ = "end of CoverTab[69962]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:206
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/bidi/prop.go:206
var _ = _go_fuzz_dep_.CoverTab
