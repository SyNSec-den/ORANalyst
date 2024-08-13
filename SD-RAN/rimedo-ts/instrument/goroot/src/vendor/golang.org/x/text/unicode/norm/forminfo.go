// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:5
package norm

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:5
)

import "encoding/binary"

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:33
const (
	qcInfoMask	= 0x3F	// to clear all but the relevant bits in a qcInfo
	headerLenMask	= 0x3F	// extract the length value from the header byte
	headerFlagsMask	= 0xC0	// extract the qcInfo bits from the header byte
)

// Properties provides access to normalization properties of a rune.
type Properties struct {
	pos	uint8	// start position in reorderBuffer; used in composition.go
	size	uint8	// length of UTF-8 encoding of this rune
	ccc	uint8	// leading canonical combining class (ccc if not decomposition)
	tccc	uint8	// trailing canonical combining class (ccc if not decomposition)
	nLead	uint8	// number of leading non-starters.
	flags	qcInfo	// quick check flags
	index	uint16
}

// functions dispatchable per form
type lookupFunc func(b input, i int) Properties

// formInfo holds Form-specific functions and tables.
type formInfo struct {
	form				Form
	composing, compatibility	bool	// form type
	info				lookupFunc
	nextMain			iterFunc
}

var formTable = []*formInfo{{
	form:		NFC,
	composing:	true,
	compatibility:	false,
	info:		lookupInfoNFC,
	nextMain:	nextComposed,
}, {
	form:		NFD,
	composing:	false,
	compatibility:	false,
	info:		lookupInfoNFC,
	nextMain:	nextDecomposed,
}, {
	form:		NFKC,
	composing:	true,
	compatibility:	true,
	info:		lookupInfoNFKC,
	nextMain:	nextComposed,
}, {
	form:		NFKD,
	composing:	false,
	compatibility:	true,
	info:		lookupInfoNFKC,
	nextMain:	nextDecomposed,
}}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:93
// BoundaryBefore returns true if this rune starts a new segment and
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:93
// cannot combine with any rune on the left.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:95
func (p Properties) BoundaryBefore() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:95
	_go_fuzz_dep_.CoverTab[33083]++
										if p.ccc == 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:96
		_go_fuzz_dep_.CoverTab[33085]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:96
		return !p.combinesBackward()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:96
		// _ = "end of CoverTab[33085]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:96
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:96
		_go_fuzz_dep_.CoverTab[33086]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:97
		// _ = "end of CoverTab[33086]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:98
		_go_fuzz_dep_.CoverTab[33087]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:98
		// _ = "end of CoverTab[33087]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:98
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:98
	// _ = "end of CoverTab[33083]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:98
	_go_fuzz_dep_.CoverTab[33084]++

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:102
	return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:102
	// _ = "end of CoverTab[33084]"
}

// BoundaryAfter returns true if runes cannot combine with or otherwise
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:105
// interact with this or previous runes.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:107
func (p Properties) BoundaryAfter() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:107
	_go_fuzz_dep_.CoverTab[33088]++

										return p.isInert()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:109
	// _ = "end of CoverTab[33088]"
}

// We pack quick check data in 4 bits:
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:112
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:112
//	5:    Combines forward  (0 == false, 1 == true)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:112
//	4..3: NFC_QC Yes(00), No (10), or Maybe (11)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:112
//	2:    NFD_QC Yes (0) or No (1). No also means there is a decomposition.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:112
//	1..0: Number of trailing non-starters.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:112
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:112
// When all 4 bits are zero, the character is inert, meaning it is never
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:112
// influenced by normalization.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:121
type qcInfo uint8

func (p Properties) isYesC() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:123
	_go_fuzz_dep_.CoverTab[33089]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:123
	return p.flags&0x10 == 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:123
	// _ = "end of CoverTab[33089]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:123
}
func (p Properties) isYesD() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:124
	_go_fuzz_dep_.CoverTab[33090]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:124
	return p.flags&0x4 == 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:124
	// _ = "end of CoverTab[33090]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:124
}

func (p Properties) combinesForward() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:126
	_go_fuzz_dep_.CoverTab[33091]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:126
	return p.flags&0x20 != 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:126
	// _ = "end of CoverTab[33091]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:126
}
func (p Properties) combinesBackward() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:127
	_go_fuzz_dep_.CoverTab[33092]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:127
	return p.flags&0x8 != 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:127
	// _ = "end of CoverTab[33092]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:127
}
func (p Properties) hasDecomposition() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:128
	_go_fuzz_dep_.CoverTab[33093]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:128
	return p.flags&0x4 != 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:128
	// _ = "end of CoverTab[33093]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:128
}

func (p Properties) isInert() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:130
	_go_fuzz_dep_.CoverTab[33094]++
										return p.flags&qcInfoMask == 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:131
		_go_fuzz_dep_.CoverTab[33095]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:131
		return p.ccc == 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:131
		// _ = "end of CoverTab[33095]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:131
	}()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:131
	// _ = "end of CoverTab[33094]"
}

func (p Properties) multiSegment() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:134
	_go_fuzz_dep_.CoverTab[33096]++
										return p.index >= firstMulti && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:135
		_go_fuzz_dep_.CoverTab[33097]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:135
		return p.index < endMulti
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:135
		// _ = "end of CoverTab[33097]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:135
	}()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:135
	// _ = "end of CoverTab[33096]"
}

func (p Properties) nLeadingNonStarters() uint8 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:138
	_go_fuzz_dep_.CoverTab[33098]++
										return p.nLead
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:139
	// _ = "end of CoverTab[33098]"
}

func (p Properties) nTrailingNonStarters() uint8 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:142
	_go_fuzz_dep_.CoverTab[33099]++
										return uint8(p.flags & 0x03)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:143
	// _ = "end of CoverTab[33099]"
}

// Decomposition returns the decomposition for the underlying rune
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:146
// or nil if there is none.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:148
func (p Properties) Decomposition() []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:148
	_go_fuzz_dep_.CoverTab[33100]++

										if p.index == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:150
		_go_fuzz_dep_.CoverTab[33102]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:151
		// _ = "end of CoverTab[33102]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:152
		_go_fuzz_dep_.CoverTab[33103]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:152
		// _ = "end of CoverTab[33103]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:152
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:152
	// _ = "end of CoverTab[33100]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:152
	_go_fuzz_dep_.CoverTab[33101]++
										i := p.index
										n := decomps[i] & headerLenMask
										i++
										return decomps[i : i+uint16(n)]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:156
	// _ = "end of CoverTab[33101]"
}

// Size returns the length of UTF-8 encoding of the rune.
func (p Properties) Size() int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:160
	_go_fuzz_dep_.CoverTab[33104]++
										return int(p.size)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:161
	// _ = "end of CoverTab[33104]"
}

// CCC returns the canonical combining class of the underlying rune.
func (p Properties) CCC() uint8 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:165
	_go_fuzz_dep_.CoverTab[33105]++
										if p.index >= firstCCCZeroExcept {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:166
		_go_fuzz_dep_.CoverTab[33107]++
											return 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:167
		// _ = "end of CoverTab[33107]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:168
		_go_fuzz_dep_.CoverTab[33108]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:168
		// _ = "end of CoverTab[33108]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:168
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:168
	// _ = "end of CoverTab[33105]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:168
	_go_fuzz_dep_.CoverTab[33106]++
										return ccc[p.ccc]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:169
	// _ = "end of CoverTab[33106]"
}

// LeadCCC returns the CCC of the first rune in the decomposition.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:172
// If there is no decomposition, LeadCCC equals CCC.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:174
func (p Properties) LeadCCC() uint8 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:174
	_go_fuzz_dep_.CoverTab[33109]++
										return ccc[p.ccc]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:175
	// _ = "end of CoverTab[33109]"
}

// TrailCCC returns the CCC of the last rune in the decomposition.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:178
// If there is no decomposition, TrailCCC equals CCC.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:180
func (p Properties) TrailCCC() uint8 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:180
	_go_fuzz_dep_.CoverTab[33110]++
										return ccc[p.tccc]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:181
	// _ = "end of CoverTab[33110]"
}

func buildRecompMap() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:184
	_go_fuzz_dep_.CoverTab[33111]++
										recompMap = make(map[uint32]rune, len(recompMapPacked)/8)
										var buf [8]byte
										for i := 0; i < len(recompMapPacked); i += 8 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:187
		_go_fuzz_dep_.CoverTab[33112]++
											copy(buf[:], recompMapPacked[i:i+8])
											key := binary.BigEndian.Uint32(buf[:4])
											val := binary.BigEndian.Uint32(buf[4:])
											recompMap[key] = rune(val)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:191
		// _ = "end of CoverTab[33112]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:192
	// _ = "end of CoverTab[33111]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:202
// combine returns the combined rune or 0 if it doesn't exist.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:202
//
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:202
// The caller is responsible for calling
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:202
// recompMapOnce.Do(buildRecompMap) sometime before this is called.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:206
func combine(a, b rune) rune {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:206
	_go_fuzz_dep_.CoverTab[33113]++
										key := uint32(uint16(a))<<16 + uint32(uint16(b))
										if recompMap == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:208
		_go_fuzz_dep_.CoverTab[33115]++
											panic("caller error")
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:209
		// _ = "end of CoverTab[33115]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:210
		_go_fuzz_dep_.CoverTab[33116]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:210
		// _ = "end of CoverTab[33116]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:210
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:210
	// _ = "end of CoverTab[33113]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:210
	_go_fuzz_dep_.CoverTab[33114]++
										return recompMap[key]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:211
	// _ = "end of CoverTab[33114]"
}

func lookupInfoNFC(b input, i int) Properties {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:214
	_go_fuzz_dep_.CoverTab[33117]++
										v, sz := b.charinfoNFC(i)
										return compInfo(v, sz)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:216
	// _ = "end of CoverTab[33117]"
}

func lookupInfoNFKC(b input, i int) Properties {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:219
	_go_fuzz_dep_.CoverTab[33118]++
										v, sz := b.charinfoNFKC(i)
										return compInfo(v, sz)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:221
	// _ = "end of CoverTab[33118]"
}

// Properties returns properties for the first rune in s.
func (f Form) Properties(s []byte) Properties {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:225
	_go_fuzz_dep_.CoverTab[33119]++
										if f == NFC || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:226
		_go_fuzz_dep_.CoverTab[33121]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:226
		return f == NFD
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:226
		// _ = "end of CoverTab[33121]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:226
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:226
		_go_fuzz_dep_.CoverTab[33122]++
											return compInfo(nfcData.lookup(s))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:227
		// _ = "end of CoverTab[33122]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:228
		_go_fuzz_dep_.CoverTab[33123]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:228
		// _ = "end of CoverTab[33123]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:228
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:228
	// _ = "end of CoverTab[33119]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:228
	_go_fuzz_dep_.CoverTab[33120]++
										return compInfo(nfkcData.lookup(s))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:229
	// _ = "end of CoverTab[33120]"
}

// PropertiesString returns properties for the first rune in s.
func (f Form) PropertiesString(s string) Properties {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:233
	_go_fuzz_dep_.CoverTab[33124]++
										if f == NFC || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:234
		_go_fuzz_dep_.CoverTab[33126]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:234
		return f == NFD
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:234
		// _ = "end of CoverTab[33126]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:234
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:234
		_go_fuzz_dep_.CoverTab[33127]++
											return compInfo(nfcData.lookupString(s))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:235
		// _ = "end of CoverTab[33127]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:236
		_go_fuzz_dep_.CoverTab[33128]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:236
		// _ = "end of CoverTab[33128]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:236
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:236
	// _ = "end of CoverTab[33124]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:236
	_go_fuzz_dep_.CoverTab[33125]++
										return compInfo(nfkcData.lookupString(s))
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:237
	// _ = "end of CoverTab[33125]"
}

// compInfo converts the information contained in v and sz
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:240
// to a Properties.  See the comment at the top of the file
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:240
// for more information on the format.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:243
func compInfo(v uint16, sz int) Properties {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:243
	_go_fuzz_dep_.CoverTab[33129]++
										if v == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:244
		_go_fuzz_dep_.CoverTab[33132]++
											return Properties{size: uint8(sz)}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:245
		// _ = "end of CoverTab[33132]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:246
		_go_fuzz_dep_.CoverTab[33133]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:246
		if v >= 0x8000 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:246
			_go_fuzz_dep_.CoverTab[33134]++
												p := Properties{
				size:	uint8(sz),
				ccc:	uint8(v),
				tccc:	uint8(v),
				flags:	qcInfo(v >> 8),
			}
			if p.ccc > 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:253
				_go_fuzz_dep_.CoverTab[33136]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:253
				return p.combinesBackward()
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:253
				// _ = "end of CoverTab[33136]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:253
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:253
				_go_fuzz_dep_.CoverTab[33137]++
													p.nLead = uint8(p.flags & 0x3)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:254
				// _ = "end of CoverTab[33137]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:255
				_go_fuzz_dep_.CoverTab[33138]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:255
				// _ = "end of CoverTab[33138]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:255
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:255
			// _ = "end of CoverTab[33134]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:255
			_go_fuzz_dep_.CoverTab[33135]++
												return p
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:256
			// _ = "end of CoverTab[33135]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:257
			_go_fuzz_dep_.CoverTab[33139]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:257
			// _ = "end of CoverTab[33139]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:257
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:257
		// _ = "end of CoverTab[33133]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:257
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:257
	// _ = "end of CoverTab[33129]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:257
	_go_fuzz_dep_.CoverTab[33130]++

										h := decomps[v]
										f := (qcInfo(h&headerFlagsMask) >> 2) | 0x4
										p := Properties{size: uint8(sz), flags: f, index: v}
										if v >= firstCCC {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:262
		_go_fuzz_dep_.CoverTab[33140]++
											v += uint16(h&headerLenMask) + 1
											c := decomps[v]
											p.tccc = c >> 2
											p.flags |= qcInfo(c & 0x3)
											if v >= firstLeadingCCC {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:267
			_go_fuzz_dep_.CoverTab[33141]++
												p.nLead = c & 0x3
												if v >= firstStarterWithNLead {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:269
				_go_fuzz_dep_.CoverTab[33143]++

													p.flags &= 0x03
													p.index = 0
													return p
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:273
				// _ = "end of CoverTab[33143]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:274
				_go_fuzz_dep_.CoverTab[33144]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:274
				// _ = "end of CoverTab[33144]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:274
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:274
			// _ = "end of CoverTab[33141]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:274
			_go_fuzz_dep_.CoverTab[33142]++
												p.ccc = decomps[v+1]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:275
			// _ = "end of CoverTab[33142]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:276
			_go_fuzz_dep_.CoverTab[33145]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:276
			// _ = "end of CoverTab[33145]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:276
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:276
		// _ = "end of CoverTab[33140]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:277
		_go_fuzz_dep_.CoverTab[33146]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:277
		// _ = "end of CoverTab[33146]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:277
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:277
	// _ = "end of CoverTab[33130]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:277
	_go_fuzz_dep_.CoverTab[33131]++
										return p
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:278
	// _ = "end of CoverTab[33131]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:279
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/forminfo.go:279
var _ = _go_fuzz_dep_.CoverTab
