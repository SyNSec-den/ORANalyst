// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:5
// Package bidirule implements the Bidi Rule defined by RFC 5893.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:5
// This package is under development. The API may change without notice and
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:5
// without preserving backward compatibility.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:9
package bidirule

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:9
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:9
)

import (
	"errors"
	"unicode/utf8"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/bidi"
)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:45
// ErrInvalid indicates a label is invalid according to the Bidi Rule.
var ErrInvalid = errors.New("bidirule: failed Bidi Rule")

type ruleState uint8

const (
	ruleInitial	ruleState	= iota
	ruleLTR
	ruleLTRFinal
	ruleRTL
	ruleRTLFinal
	ruleInvalid
)

type ruleTransition struct {
	next	ruleState
	mask	uint16
}

var transitions = [...][2]ruleTransition{

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:68
	ruleInitial: {
		{ruleLTRFinal, 1 << bidi.L},
		{ruleRTLFinal, 1<<bidi.R | 1<<bidi.AL},
	},
	ruleRTL: {

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:76
		{ruleRTLFinal, 1<<bidi.R | 1<<bidi.AL | 1<<bidi.EN | 1<<bidi.AN},

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:81
		{ruleRTL, 1<<bidi.ES | 1<<bidi.CS | 1<<bidi.ET | 1<<bidi.ON | 1<<bidi.BN | 1<<bidi.NSM},
	},
	ruleRTLFinal: {

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:87
		{ruleRTLFinal, 1<<bidi.R | 1<<bidi.AL | 1<<bidi.EN | 1<<bidi.AN | 1<<bidi.NSM},

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:92
		{ruleRTL, 1<<bidi.ES | 1<<bidi.CS | 1<<bidi.ET | 1<<bidi.ON | 1<<bidi.BN},
	},
	ruleLTR: {

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:98
		{ruleLTRFinal, 1<<bidi.L | 1<<bidi.EN},

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:103
		{ruleLTR, 1<<bidi.ES | 1<<bidi.CS | 1<<bidi.ET | 1<<bidi.ON | 1<<bidi.BN | 1<<bidi.NSM},
	},
	ruleLTRFinal: {

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:109
		{ruleLTRFinal, 1<<bidi.L | 1<<bidi.EN | 1<<bidi.NSM},

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:114
		{ruleLTR, 1<<bidi.ES | 1<<bidi.CS | 1<<bidi.ET | 1<<bidi.ON | 1<<bidi.BN},
	},
	ruleInvalid: {
		{ruleInvalid, 0},
		{ruleInvalid, 0},
	},
}

// [2.4] In an RTL label, if an EN is present, no AN may be present, and
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:122
// vice versa.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:124
const exclusiveRTL = uint16(1<<bidi.EN | 1<<bidi.AN)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:132
// Direction reports the direction of the given label as defined by RFC 5893.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:132
// The Bidi Rule does not have to be applied to labels of the category
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:132
// LeftToRight.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:135
func Direction(b []byte) bidi.Direction {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:135
	_go_fuzz_dep_.CoverTab[70113]++
												for i := 0; i < len(b); {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:136
		_go_fuzz_dep_.CoverTab[70115]++
													e, sz := bidi.Lookup(b[i:])
													if sz == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:138
			_go_fuzz_dep_.CoverTab[70118]++
														i++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:139
			// _ = "end of CoverTab[70118]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:140
			_go_fuzz_dep_.CoverTab[70119]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:140
			// _ = "end of CoverTab[70119]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:140
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:140
		// _ = "end of CoverTab[70115]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:140
		_go_fuzz_dep_.CoverTab[70116]++
													c := e.Class()
													if c == bidi.R || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:142
			_go_fuzz_dep_.CoverTab[70120]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:142
			return c == bidi.AL
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:142
			// _ = "end of CoverTab[70120]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:142
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:142
			_go_fuzz_dep_.CoverTab[70121]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:142
			return c == bidi.AN
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:142
			// _ = "end of CoverTab[70121]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:142
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:142
			_go_fuzz_dep_.CoverTab[70122]++
														return bidi.RightToLeft
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:143
			// _ = "end of CoverTab[70122]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:144
			_go_fuzz_dep_.CoverTab[70123]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:144
			// _ = "end of CoverTab[70123]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:144
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:144
		// _ = "end of CoverTab[70116]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:144
		_go_fuzz_dep_.CoverTab[70117]++
													i += sz
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:145
		// _ = "end of CoverTab[70117]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:146
	// _ = "end of CoverTab[70113]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:146
	_go_fuzz_dep_.CoverTab[70114]++
												return bidi.LeftToRight
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:147
	// _ = "end of CoverTab[70114]"
}

// DirectionString reports the direction of the given label as defined by RFC
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:150
// 5893. The Bidi Rule does not have to be applied to labels of the category
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:150
// LeftToRight.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:153
func DirectionString(s string) bidi.Direction {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:153
	_go_fuzz_dep_.CoverTab[70124]++
												for i := 0; i < len(s); {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:154
		_go_fuzz_dep_.CoverTab[70126]++
													e, sz := bidi.LookupString(s[i:])
													if sz == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:156
			_go_fuzz_dep_.CoverTab[70129]++
														i++
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:158
			// _ = "end of CoverTab[70129]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:159
			_go_fuzz_dep_.CoverTab[70130]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:159
			// _ = "end of CoverTab[70130]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:159
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:159
		// _ = "end of CoverTab[70126]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:159
		_go_fuzz_dep_.CoverTab[70127]++
													c := e.Class()
													if c == bidi.R || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:161
			_go_fuzz_dep_.CoverTab[70131]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:161
			return c == bidi.AL
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:161
			// _ = "end of CoverTab[70131]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:161
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:161
			_go_fuzz_dep_.CoverTab[70132]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:161
			return c == bidi.AN
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:161
			// _ = "end of CoverTab[70132]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:161
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:161
			_go_fuzz_dep_.CoverTab[70133]++
														return bidi.RightToLeft
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:162
			// _ = "end of CoverTab[70133]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:163
			_go_fuzz_dep_.CoverTab[70134]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:163
			// _ = "end of CoverTab[70134]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:163
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:163
		// _ = "end of CoverTab[70127]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:163
		_go_fuzz_dep_.CoverTab[70128]++
													i += sz
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:164
		// _ = "end of CoverTab[70128]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:165
	// _ = "end of CoverTab[70124]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:165
	_go_fuzz_dep_.CoverTab[70125]++
												return bidi.LeftToRight
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:166
	// _ = "end of CoverTab[70125]"
}

// Valid reports whether b conforms to the BiDi rule.
func Valid(b []byte) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:170
	_go_fuzz_dep_.CoverTab[70135]++
												var t Transformer
												if n, ok := t.advance(b); !ok || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:172
		_go_fuzz_dep_.CoverTab[70137]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:172
		return n < len(b)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:172
		// _ = "end of CoverTab[70137]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:172
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:172
		_go_fuzz_dep_.CoverTab[70138]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:173
		// _ = "end of CoverTab[70138]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:174
		_go_fuzz_dep_.CoverTab[70139]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:174
		// _ = "end of CoverTab[70139]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:174
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:174
	// _ = "end of CoverTab[70135]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:174
	_go_fuzz_dep_.CoverTab[70136]++
												return t.isFinal()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:175
	// _ = "end of CoverTab[70136]"
}

// ValidString reports whether s conforms to the BiDi rule.
func ValidString(s string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:179
	_go_fuzz_dep_.CoverTab[70140]++
												var t Transformer
												if n, ok := t.advanceString(s); !ok || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:181
		_go_fuzz_dep_.CoverTab[70142]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:181
		return n < len(s)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:181
		// _ = "end of CoverTab[70142]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:181
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:181
		_go_fuzz_dep_.CoverTab[70143]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:182
		// _ = "end of CoverTab[70143]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:183
		_go_fuzz_dep_.CoverTab[70144]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:183
		// _ = "end of CoverTab[70144]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:183
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:183
	// _ = "end of CoverTab[70140]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:183
	_go_fuzz_dep_.CoverTab[70141]++
												return t.isFinal()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:184
	// _ = "end of CoverTab[70141]"
}

// New returns a Transformer that verifies that input adheres to the Bidi Rule.
func New() *Transformer {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:188
	_go_fuzz_dep_.CoverTab[70145]++
												return &Transformer{}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:189
	// _ = "end of CoverTab[70145]"
}

// Transformer implements transform.Transform.
type Transformer struct {
	state	ruleState
	hasRTL	bool
	seen	uint16
}

// A rule can only be violated for "Bidi Domain names", meaning if one of the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:199
// following categories has been observed.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:201
func (t *Transformer) isRTL() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:201
	_go_fuzz_dep_.CoverTab[70146]++
												const isRTL = 1<<bidi.R | 1<<bidi.AL | 1<<bidi.AN
												return t.seen&isRTL != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:203
	// _ = "end of CoverTab[70146]"
}

// Reset implements transform.Transformer.
func (t *Transformer) Reset() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:207
	_go_fuzz_dep_.CoverTab[70147]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:207
	*t = Transformer{}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:207
	// _ = "end of CoverTab[70147]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:207
}

// Transform implements transform.Transformer. This Transformer has state and
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:209
// needs to be reset between uses.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:211
func (t *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:211
	_go_fuzz_dep_.CoverTab[70148]++
												if len(dst) < len(src) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:212
		_go_fuzz_dep_.CoverTab[70151]++
													src = src[:len(dst)]
													atEOF = false
													err = transform.ErrShortDst
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:215
		// _ = "end of CoverTab[70151]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:216
		_go_fuzz_dep_.CoverTab[70152]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:216
		// _ = "end of CoverTab[70152]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:216
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:216
	// _ = "end of CoverTab[70148]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:216
	_go_fuzz_dep_.CoverTab[70149]++
												n, err1 := t.Span(src, atEOF)
												copy(dst, src[:n])
												if err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:219
		_go_fuzz_dep_.CoverTab[70153]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:219
		return err1 != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:219
			_go_fuzz_dep_.CoverTab[70154]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:219
			return err1 != transform.ErrShortSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:219
			// _ = "end of CoverTab[70154]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:219
		}()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:219
		// _ = "end of CoverTab[70153]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:219
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:219
		_go_fuzz_dep_.CoverTab[70155]++
													err = err1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:220
		// _ = "end of CoverTab[70155]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:221
		_go_fuzz_dep_.CoverTab[70156]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:221
		// _ = "end of CoverTab[70156]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:221
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:221
	// _ = "end of CoverTab[70149]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:221
	_go_fuzz_dep_.CoverTab[70150]++
												return n, n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:222
	// _ = "end of CoverTab[70150]"
}

// Span returns the first n bytes of src that conform to the Bidi rule.
func (t *Transformer) Span(src []byte, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:226
	_go_fuzz_dep_.CoverTab[70157]++
												if t.state == ruleInvalid && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:227
		_go_fuzz_dep_.CoverTab[70160]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:227
		return t.isRTL()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:227
		// _ = "end of CoverTab[70160]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:227
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:227
		_go_fuzz_dep_.CoverTab[70161]++
													return 0, ErrInvalid
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:228
		// _ = "end of CoverTab[70161]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:229
		_go_fuzz_dep_.CoverTab[70162]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:229
		// _ = "end of CoverTab[70162]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:229
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:229
	// _ = "end of CoverTab[70157]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:229
	_go_fuzz_dep_.CoverTab[70158]++
												n, ok := t.advance(src)
												switch {
	case !ok:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:232
		_go_fuzz_dep_.CoverTab[70163]++
													err = ErrInvalid
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:233
		// _ = "end of CoverTab[70163]"
	case n < len(src):
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:234
		_go_fuzz_dep_.CoverTab[70164]++
													if !atEOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:235
			_go_fuzz_dep_.CoverTab[70168]++
														err = transform.ErrShortSrc
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:237
			// _ = "end of CoverTab[70168]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:238
			_go_fuzz_dep_.CoverTab[70169]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:238
			// _ = "end of CoverTab[70169]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:238
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:238
		// _ = "end of CoverTab[70164]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:238
		_go_fuzz_dep_.CoverTab[70165]++
													err = ErrInvalid
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:239
		// _ = "end of CoverTab[70165]"
	case !t.isFinal():
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:240
		_go_fuzz_dep_.CoverTab[70166]++
													err = ErrInvalid
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:241
		// _ = "end of CoverTab[70166]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:241
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:241
		_go_fuzz_dep_.CoverTab[70167]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:241
		// _ = "end of CoverTab[70167]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:242
	// _ = "end of CoverTab[70158]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:242
	_go_fuzz_dep_.CoverTab[70159]++
												return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:243
	// _ = "end of CoverTab[70159]"
}

// Precomputing the ASCII values decreases running time for the ASCII fast path
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:246
// by about 30%.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:248
var asciiTable [128]bidi.Properties

func init() {
	for i := range asciiTable {
		p, _ := bidi.LookupRune(rune(i))
		asciiTable[i] = p
	}
}

func (t *Transformer) advance(s []byte) (n int, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:257
	_go_fuzz_dep_.CoverTab[70170]++
												var e bidi.Properties
												var sz int
												for n < len(s) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:260
		_go_fuzz_dep_.CoverTab[70172]++
													if s[n] < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:261
			_go_fuzz_dep_.CoverTab[70176]++
														e, sz = asciiTable[s[n]], 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:262
			// _ = "end of CoverTab[70176]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:263
			_go_fuzz_dep_.CoverTab[70177]++
														e, sz = bidi.Lookup(s[n:])
														if sz <= 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:265
				_go_fuzz_dep_.CoverTab[70178]++
															if sz == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:266
					_go_fuzz_dep_.CoverTab[70180]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:270
					return n, false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:270
					// _ = "end of CoverTab[70180]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:271
					_go_fuzz_dep_.CoverTab[70181]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:271
					// _ = "end of CoverTab[70181]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:271
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:271
				// _ = "end of CoverTab[70178]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:271
				_go_fuzz_dep_.CoverTab[70179]++
															return n, true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:272
				// _ = "end of CoverTab[70179]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:273
				_go_fuzz_dep_.CoverTab[70182]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:273
				// _ = "end of CoverTab[70182]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:273
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:273
			// _ = "end of CoverTab[70177]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:274
		// _ = "end of CoverTab[70172]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:274
		_go_fuzz_dep_.CoverTab[70173]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:277
		c := uint16(1 << e.Class())
		t.seen |= c
		if t.seen&exclusiveRTL == exclusiveRTL {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:279
			_go_fuzz_dep_.CoverTab[70183]++
														t.state = ruleInvalid
														return n, false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:281
			// _ = "end of CoverTab[70183]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:282
			_go_fuzz_dep_.CoverTab[70184]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:282
			// _ = "end of CoverTab[70184]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:282
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:282
		// _ = "end of CoverTab[70173]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:282
		_go_fuzz_dep_.CoverTab[70174]++
													switch tr := transitions[t.state]; {
		case tr[0].mask&c != 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:284
			_go_fuzz_dep_.CoverTab[70185]++
														t.state = tr[0].next
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:285
			// _ = "end of CoverTab[70185]"
		case tr[1].mask&c != 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:286
			_go_fuzz_dep_.CoverTab[70186]++
														t.state = tr[1].next
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:287
			// _ = "end of CoverTab[70186]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:288
			_go_fuzz_dep_.CoverTab[70187]++
														t.state = ruleInvalid
														if t.isRTL() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:290
				_go_fuzz_dep_.CoverTab[70188]++
															return n, false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:291
				// _ = "end of CoverTab[70188]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:292
				_go_fuzz_dep_.CoverTab[70189]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:292
				// _ = "end of CoverTab[70189]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:292
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:292
			// _ = "end of CoverTab[70187]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:293
		// _ = "end of CoverTab[70174]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:293
		_go_fuzz_dep_.CoverTab[70175]++
													n += sz
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:294
		// _ = "end of CoverTab[70175]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:295
	// _ = "end of CoverTab[70170]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:295
	_go_fuzz_dep_.CoverTab[70171]++
												return n, true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:296
	// _ = "end of CoverTab[70171]"
}

func (t *Transformer) advanceString(s string) (n int, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:299
	_go_fuzz_dep_.CoverTab[70190]++
												var e bidi.Properties
												var sz int
												for n < len(s) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:302
		_go_fuzz_dep_.CoverTab[70192]++
													if s[n] < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:303
			_go_fuzz_dep_.CoverTab[70196]++
														e, sz = asciiTable[s[n]], 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:304
			// _ = "end of CoverTab[70196]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:305
			_go_fuzz_dep_.CoverTab[70197]++
														e, sz = bidi.LookupString(s[n:])
														if sz <= 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:307
				_go_fuzz_dep_.CoverTab[70198]++
															if sz == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:308
					_go_fuzz_dep_.CoverTab[70200]++
																return n, false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:309
					// _ = "end of CoverTab[70200]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:310
					_go_fuzz_dep_.CoverTab[70201]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:310
					// _ = "end of CoverTab[70201]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:310
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:310
				// _ = "end of CoverTab[70198]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:310
				_go_fuzz_dep_.CoverTab[70199]++
															return n, true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:311
				// _ = "end of CoverTab[70199]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:312
				_go_fuzz_dep_.CoverTab[70202]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:312
				// _ = "end of CoverTab[70202]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:312
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:312
			// _ = "end of CoverTab[70197]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:313
		// _ = "end of CoverTab[70192]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:313
		_go_fuzz_dep_.CoverTab[70193]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:316
		c := uint16(1 << e.Class())
		t.seen |= c
		if t.seen&exclusiveRTL == exclusiveRTL {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:318
			_go_fuzz_dep_.CoverTab[70203]++
														t.state = ruleInvalid
														return n, false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:320
			// _ = "end of CoverTab[70203]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:321
			_go_fuzz_dep_.CoverTab[70204]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:321
			// _ = "end of CoverTab[70204]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:321
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:321
		// _ = "end of CoverTab[70193]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:321
		_go_fuzz_dep_.CoverTab[70194]++
													switch tr := transitions[t.state]; {
		case tr[0].mask&c != 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:323
			_go_fuzz_dep_.CoverTab[70205]++
														t.state = tr[0].next
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:324
			// _ = "end of CoverTab[70205]"
		case tr[1].mask&c != 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:325
			_go_fuzz_dep_.CoverTab[70206]++
														t.state = tr[1].next
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:326
			// _ = "end of CoverTab[70206]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:327
			_go_fuzz_dep_.CoverTab[70207]++
														t.state = ruleInvalid
														if t.isRTL() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:329
				_go_fuzz_dep_.CoverTab[70208]++
															return n, false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:330
				// _ = "end of CoverTab[70208]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:331
				_go_fuzz_dep_.CoverTab[70209]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:331
				// _ = "end of CoverTab[70209]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:331
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:331
			// _ = "end of CoverTab[70207]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:332
		// _ = "end of CoverTab[70194]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:332
		_go_fuzz_dep_.CoverTab[70195]++
													n += sz
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:333
		// _ = "end of CoverTab[70195]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:334
	// _ = "end of CoverTab[70190]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:334
	_go_fuzz_dep_.CoverTab[70191]++
												return n, true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:335
	// _ = "end of CoverTab[70191]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:336
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/secure/bidirule/bidirule.go:336
var _ = _go_fuzz_dep_.CoverTab
