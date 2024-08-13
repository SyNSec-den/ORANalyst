// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:5
package norm

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:5
)

import "encoding/binary"

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:33
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

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:93
// BoundaryBefore returns true if this rune starts a new segment and
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:93
// cannot combine with any rune on the left.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:95
func (p Properties) BoundaryBefore() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:95
	_go_fuzz_dep_.CoverTab[70364]++
											if p.ccc == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:96
		_go_fuzz_dep_.CoverTab[70366]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:96
		return !p.combinesBackward()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:96
		// _ = "end of CoverTab[70366]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:96
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:96
		_go_fuzz_dep_.CoverTab[70367]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:97
		// _ = "end of CoverTab[70367]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:98
		_go_fuzz_dep_.CoverTab[70368]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:98
		// _ = "end of CoverTab[70368]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:98
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:98
		// _ = "end of CoverTab[70364]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:98
		_go_fuzz_dep_.CoverTab[70365]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:102
	return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:102
	// _ = "end of CoverTab[70365]"
}

// BoundaryAfter returns true if runes cannot combine with or otherwise
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:105
// interact with this or previous runes.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:107
func (p Properties) BoundaryAfter() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:107
	_go_fuzz_dep_.CoverTab[70369]++

												return p.isInert()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:109
	// _ = "end of CoverTab[70369]"
}

// We pack quick check data in 4 bits:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:112
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:112
//	5:    Combines forward  (0 == false, 1 == true)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:112
//	4..3: NFC_QC Yes(00), No (10), or Maybe (11)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:112
//	2:    NFD_QC Yes (0) or No (1). No also means there is a decomposition.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:112
//	1..0: Number of trailing non-starters.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:112
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:112
// When all 4 bits are zero, the character is inert, meaning it is never
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:112
// influenced by normalization.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:121
type qcInfo uint8

func (p Properties) isYesC() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:123
	_go_fuzz_dep_.CoverTab[70370]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:123
	return p.flags&0x10 == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:123
	// _ = "end of CoverTab[70370]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:123
}
func (p Properties) isYesD() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:124
	_go_fuzz_dep_.CoverTab[70371]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:124
	return p.flags&0x4 == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:124
	// _ = "end of CoverTab[70371]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:124
}

func (p Properties) combinesForward() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:126
	_go_fuzz_dep_.CoverTab[70372]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:126
	return p.flags&0x20 != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:126
	// _ = "end of CoverTab[70372]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:126
}
func (p Properties) combinesBackward() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:127
	_go_fuzz_dep_.CoverTab[70373]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:127
	return p.flags&0x8 != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:127
	// _ = "end of CoverTab[70373]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:127
}
func (p Properties) hasDecomposition() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:128
	_go_fuzz_dep_.CoverTab[70374]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:128
	return p.flags&0x4 != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:128
	// _ = "end of CoverTab[70374]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:128
}

func (p Properties) isInert() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:130
	_go_fuzz_dep_.CoverTab[70375]++
												return p.flags&qcInfoMask == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:131
		_go_fuzz_dep_.CoverTab[70376]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:131
		return p.ccc == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:131
		// _ = "end of CoverTab[70376]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:131
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:131
	// _ = "end of CoverTab[70375]"
}

func (p Properties) multiSegment() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:134
	_go_fuzz_dep_.CoverTab[70377]++
												return p.index >= firstMulti && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:135
		_go_fuzz_dep_.CoverTab[70378]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:135
		return p.index < endMulti
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:135
		// _ = "end of CoverTab[70378]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:135
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:135
	// _ = "end of CoverTab[70377]"
}

func (p Properties) nLeadingNonStarters() uint8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:138
	_go_fuzz_dep_.CoverTab[70379]++
												return p.nLead
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:139
	// _ = "end of CoverTab[70379]"
}

func (p Properties) nTrailingNonStarters() uint8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:142
	_go_fuzz_dep_.CoverTab[70380]++
												return uint8(p.flags & 0x03)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:143
	// _ = "end of CoverTab[70380]"
}

// Decomposition returns the decomposition for the underlying rune
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:146
// or nil if there is none.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:148
func (p Properties) Decomposition() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:148
	_go_fuzz_dep_.CoverTab[70381]++

												if p.index == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:150
		_go_fuzz_dep_.CoverTab[70383]++
													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:151
		// _ = "end of CoverTab[70383]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:152
		_go_fuzz_dep_.CoverTab[70384]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:152
		// _ = "end of CoverTab[70384]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:152
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:152
	// _ = "end of CoverTab[70381]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:152
	_go_fuzz_dep_.CoverTab[70382]++
												i := p.index
												n := decomps[i] & headerLenMask
												i++
												return decomps[i : i+uint16(n)]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:156
	// _ = "end of CoverTab[70382]"
}

// Size returns the length of UTF-8 encoding of the rune.
func (p Properties) Size() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:160
	_go_fuzz_dep_.CoverTab[70385]++
												return int(p.size)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:161
	// _ = "end of CoverTab[70385]"
}

// CCC returns the canonical combining class of the underlying rune.
func (p Properties) CCC() uint8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:165
	_go_fuzz_dep_.CoverTab[70386]++
												if p.index >= firstCCCZeroExcept {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:166
		_go_fuzz_dep_.CoverTab[70388]++
													return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:167
		// _ = "end of CoverTab[70388]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:168
		_go_fuzz_dep_.CoverTab[70389]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:168
		// _ = "end of CoverTab[70389]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:168
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:168
	// _ = "end of CoverTab[70386]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:168
	_go_fuzz_dep_.CoverTab[70387]++
												return ccc[p.ccc]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:169
	// _ = "end of CoverTab[70387]"
}

// LeadCCC returns the CCC of the first rune in the decomposition.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:172
// If there is no decomposition, LeadCCC equals CCC.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:174
func (p Properties) LeadCCC() uint8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:174
	_go_fuzz_dep_.CoverTab[70390]++
												return ccc[p.ccc]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:175
	// _ = "end of CoverTab[70390]"
}

// TrailCCC returns the CCC of the last rune in the decomposition.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:178
// If there is no decomposition, TrailCCC equals CCC.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:180
func (p Properties) TrailCCC() uint8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:180
	_go_fuzz_dep_.CoverTab[70391]++
												return ccc[p.tccc]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:181
	// _ = "end of CoverTab[70391]"
}

func buildRecompMap() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:184
	_go_fuzz_dep_.CoverTab[70392]++
												recompMap = make(map[uint32]rune, len(recompMapPacked)/8)
												var buf [8]byte
												for i := 0; i < len(recompMapPacked); i += 8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:187
		_go_fuzz_dep_.CoverTab[70393]++
													copy(buf[:], recompMapPacked[i:i+8])
													key := binary.BigEndian.Uint32(buf[:4])
													val := binary.BigEndian.Uint32(buf[4:])
													recompMap[key] = rune(val)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:191
		// _ = "end of CoverTab[70393]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:192
	// _ = "end of CoverTab[70392]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:202
// combine returns the combined rune or 0 if it doesn't exist.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:202
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:202
// The caller is responsible for calling
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:202
// recompMapOnce.Do(buildRecompMap) sometime before this is called.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:206
func combine(a, b rune) rune {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:206
	_go_fuzz_dep_.CoverTab[70394]++
												key := uint32(uint16(a))<<16 + uint32(uint16(b))
												if recompMap == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:208
		_go_fuzz_dep_.CoverTab[70396]++
													panic("caller error")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:209
		// _ = "end of CoverTab[70396]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:210
		_go_fuzz_dep_.CoverTab[70397]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:210
		// _ = "end of CoverTab[70397]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:210
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:210
	// _ = "end of CoverTab[70394]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:210
	_go_fuzz_dep_.CoverTab[70395]++
												return recompMap[key]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:211
	// _ = "end of CoverTab[70395]"
}

func lookupInfoNFC(b input, i int) Properties {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:214
	_go_fuzz_dep_.CoverTab[70398]++
												v, sz := b.charinfoNFC(i)
												return compInfo(v, sz)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:216
	// _ = "end of CoverTab[70398]"
}

func lookupInfoNFKC(b input, i int) Properties {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:219
	_go_fuzz_dep_.CoverTab[70399]++
												v, sz := b.charinfoNFKC(i)
												return compInfo(v, sz)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:221
	// _ = "end of CoverTab[70399]"
}

// Properties returns properties for the first rune in s.
func (f Form) Properties(s []byte) Properties {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:225
	_go_fuzz_dep_.CoverTab[70400]++
												if f == NFC || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:226
		_go_fuzz_dep_.CoverTab[70402]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:226
		return f == NFD
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:226
		// _ = "end of CoverTab[70402]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:226
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:226
		_go_fuzz_dep_.CoverTab[70403]++
													return compInfo(nfcData.lookup(s))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:227
		// _ = "end of CoverTab[70403]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:228
		_go_fuzz_dep_.CoverTab[70404]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:228
		// _ = "end of CoverTab[70404]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:228
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:228
	// _ = "end of CoverTab[70400]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:228
	_go_fuzz_dep_.CoverTab[70401]++
												return compInfo(nfkcData.lookup(s))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:229
	// _ = "end of CoverTab[70401]"
}

// PropertiesString returns properties for the first rune in s.
func (f Form) PropertiesString(s string) Properties {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:233
	_go_fuzz_dep_.CoverTab[70405]++
												if f == NFC || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:234
		_go_fuzz_dep_.CoverTab[70407]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:234
		return f == NFD
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:234
		// _ = "end of CoverTab[70407]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:234
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:234
		_go_fuzz_dep_.CoverTab[70408]++
													return compInfo(nfcData.lookupString(s))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:235
		// _ = "end of CoverTab[70408]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:236
		_go_fuzz_dep_.CoverTab[70409]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:236
		// _ = "end of CoverTab[70409]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:236
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:236
	// _ = "end of CoverTab[70405]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:236
	_go_fuzz_dep_.CoverTab[70406]++
												return compInfo(nfkcData.lookupString(s))
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:237
	// _ = "end of CoverTab[70406]"
}

// compInfo converts the information contained in v and sz
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:240
// to a Properties.  See the comment at the top of the file
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:240
// for more information on the format.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:243
func compInfo(v uint16, sz int) Properties {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:243
	_go_fuzz_dep_.CoverTab[70410]++
												if v == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:244
		_go_fuzz_dep_.CoverTab[70413]++
													return Properties{size: uint8(sz)}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:245
		// _ = "end of CoverTab[70413]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:246
		_go_fuzz_dep_.CoverTab[70414]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:246
		if v >= 0x8000 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:246
			_go_fuzz_dep_.CoverTab[70415]++
														p := Properties{
				size:	uint8(sz),
				ccc:	uint8(v),
				tccc:	uint8(v),
				flags:	qcInfo(v >> 8),
			}
			if p.ccc > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:253
				_go_fuzz_dep_.CoverTab[70417]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:253
				return p.combinesBackward()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:253
				// _ = "end of CoverTab[70417]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:253
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:253
				_go_fuzz_dep_.CoverTab[70418]++
															p.nLead = uint8(p.flags & 0x3)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:254
				// _ = "end of CoverTab[70418]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:255
				_go_fuzz_dep_.CoverTab[70419]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:255
				// _ = "end of CoverTab[70419]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:255
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:255
			// _ = "end of CoverTab[70415]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:255
			_go_fuzz_dep_.CoverTab[70416]++
														return p
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:256
			// _ = "end of CoverTab[70416]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:257
			_go_fuzz_dep_.CoverTab[70420]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:257
			// _ = "end of CoverTab[70420]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:257
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:257
		// _ = "end of CoverTab[70414]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:257
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:257
	// _ = "end of CoverTab[70410]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:257
	_go_fuzz_dep_.CoverTab[70411]++

												h := decomps[v]
												f := (qcInfo(h&headerFlagsMask) >> 2) | 0x4
												p := Properties{size: uint8(sz), flags: f, index: v}
												if v >= firstCCC {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:262
		_go_fuzz_dep_.CoverTab[70421]++
													v += uint16(h&headerLenMask) + 1
													c := decomps[v]
													p.tccc = c >> 2
													p.flags |= qcInfo(c & 0x3)
													if v >= firstLeadingCCC {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:267
			_go_fuzz_dep_.CoverTab[70422]++
														p.nLead = c & 0x3
														if v >= firstStarterWithNLead {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:269
				_go_fuzz_dep_.CoverTab[70424]++

															p.flags &= 0x03
															p.index = 0
															return p
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:273
				// _ = "end of CoverTab[70424]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:274
				_go_fuzz_dep_.CoverTab[70425]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:274
				// _ = "end of CoverTab[70425]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:274
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:274
			// _ = "end of CoverTab[70422]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:274
			_go_fuzz_dep_.CoverTab[70423]++
														p.ccc = decomps[v+1]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:275
			// _ = "end of CoverTab[70423]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:276
			_go_fuzz_dep_.CoverTab[70426]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:276
			// _ = "end of CoverTab[70426]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:276
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:276
		// _ = "end of CoverTab[70421]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:277
		_go_fuzz_dep_.CoverTab[70427]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:277
		// _ = "end of CoverTab[70427]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:277
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:277
	// _ = "end of CoverTab[70411]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:277
	_go_fuzz_dep_.CoverTab[70412]++
												return p
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:278
	// _ = "end of CoverTab[70412]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:279
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/forminfo.go:279
var _ = _go_fuzz_dep_.CoverTab
