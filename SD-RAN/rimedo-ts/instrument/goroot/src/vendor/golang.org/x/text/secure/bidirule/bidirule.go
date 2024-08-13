// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:5
// Package bidirule implements the Bidi Rule defined by RFC 5893.
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:5
// This package is under development. The API may change without notice and
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:5
// without preserving backward compatibility.
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:9
package bidirule

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:9
)
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:9
)

import (
	"errors"
	"unicode/utf8"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/bidi"
)

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:45
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

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:68
	ruleInitial: {
		{ruleLTRFinal, 1 << bidi.L},
		{ruleRTLFinal, 1<<bidi.R | 1<<bidi.AL},
	},
	ruleRTL: {

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:76
		{ruleRTLFinal, 1<<bidi.R | 1<<bidi.AL | 1<<bidi.EN | 1<<bidi.AN},

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:81
		{ruleRTL, 1<<bidi.ES | 1<<bidi.CS | 1<<bidi.ET | 1<<bidi.ON | 1<<bidi.BN | 1<<bidi.NSM},
	},
	ruleRTLFinal: {

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:87
		{ruleRTLFinal, 1<<bidi.R | 1<<bidi.AL | 1<<bidi.EN | 1<<bidi.AN | 1<<bidi.NSM},

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:92
		{ruleRTL, 1<<bidi.ES | 1<<bidi.CS | 1<<bidi.ET | 1<<bidi.ON | 1<<bidi.BN},
	},
	ruleLTR: {

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:98
		{ruleLTRFinal, 1<<bidi.L | 1<<bidi.EN},

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:103
		{ruleLTR, 1<<bidi.ES | 1<<bidi.CS | 1<<bidi.ET | 1<<bidi.ON | 1<<bidi.BN | 1<<bidi.NSM},
	},
	ruleLTRFinal: {

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:109
		{ruleLTRFinal, 1<<bidi.L | 1<<bidi.EN | 1<<bidi.NSM},

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:114
		{ruleLTR, 1<<bidi.ES | 1<<bidi.CS | 1<<bidi.ET | 1<<bidi.ON | 1<<bidi.BN},
	},
	ruleInvalid: {
		{ruleInvalid, 0},
		{ruleInvalid, 0},
	},
}

// [2.4] In an RTL label, if an EN is present, no AN may be present, and
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:122
// vice versa.
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:124
const exclusiveRTL = uint16(1<<bidi.EN | 1<<bidi.AN)

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:132
// Direction reports the direction of the given label as defined by RFC 5893.
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:132
// The Bidi Rule does not have to be applied to labels of the category
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:132
// LeftToRight.
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:135
func Direction(b []byte) bidi.Direction {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:135
	_go_fuzz_dep_.CoverTab[32832]++
											for i := 0; i < len(b); {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:136
		_go_fuzz_dep_.CoverTab[32834]++
												e, sz := bidi.Lookup(b[i:])
												if sz == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:138
			_go_fuzz_dep_.CoverTab[32837]++
													i++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:139
			// _ = "end of CoverTab[32837]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:140
			_go_fuzz_dep_.CoverTab[32838]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:140
			// _ = "end of CoverTab[32838]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:140
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:140
		// _ = "end of CoverTab[32834]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:140
		_go_fuzz_dep_.CoverTab[32835]++
												c := e.Class()
												if c == bidi.R || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:142
			_go_fuzz_dep_.CoverTab[32839]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:142
			return c == bidi.AL
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:142
			// _ = "end of CoverTab[32839]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:142
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:142
			_go_fuzz_dep_.CoverTab[32840]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:142
			return c == bidi.AN
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:142
			// _ = "end of CoverTab[32840]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:142
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:142
			_go_fuzz_dep_.CoverTab[32841]++
													return bidi.RightToLeft
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:143
			// _ = "end of CoverTab[32841]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:144
			_go_fuzz_dep_.CoverTab[32842]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:144
			// _ = "end of CoverTab[32842]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:144
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:144
		// _ = "end of CoverTab[32835]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:144
		_go_fuzz_dep_.CoverTab[32836]++
												i += sz
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:145
		// _ = "end of CoverTab[32836]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:146
	// _ = "end of CoverTab[32832]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:146
	_go_fuzz_dep_.CoverTab[32833]++
											return bidi.LeftToRight
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:147
	// _ = "end of CoverTab[32833]"
}

// DirectionString reports the direction of the given label as defined by RFC
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:150
// 5893. The Bidi Rule does not have to be applied to labels of the category
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:150
// LeftToRight.
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:153
func DirectionString(s string) bidi.Direction {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:153
	_go_fuzz_dep_.CoverTab[32843]++
											for i := 0; i < len(s); {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:154
		_go_fuzz_dep_.CoverTab[32845]++
												e, sz := bidi.LookupString(s[i:])
												if sz == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:156
			_go_fuzz_dep_.CoverTab[32848]++
													i++
													continue
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:158
			// _ = "end of CoverTab[32848]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:159
			_go_fuzz_dep_.CoverTab[32849]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:159
			// _ = "end of CoverTab[32849]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:159
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:159
		// _ = "end of CoverTab[32845]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:159
		_go_fuzz_dep_.CoverTab[32846]++
												c := e.Class()
												if c == bidi.R || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:161
			_go_fuzz_dep_.CoverTab[32850]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:161
			return c == bidi.AL
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:161
			// _ = "end of CoverTab[32850]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:161
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:161
			_go_fuzz_dep_.CoverTab[32851]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:161
			return c == bidi.AN
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:161
			// _ = "end of CoverTab[32851]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:161
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:161
			_go_fuzz_dep_.CoverTab[32852]++
													return bidi.RightToLeft
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:162
			// _ = "end of CoverTab[32852]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:163
			_go_fuzz_dep_.CoverTab[32853]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:163
			// _ = "end of CoverTab[32853]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:163
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:163
		// _ = "end of CoverTab[32846]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:163
		_go_fuzz_dep_.CoverTab[32847]++
												i += sz
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:164
		// _ = "end of CoverTab[32847]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:165
	// _ = "end of CoverTab[32843]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:165
	_go_fuzz_dep_.CoverTab[32844]++
											return bidi.LeftToRight
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:166
	// _ = "end of CoverTab[32844]"
}

// Valid reports whether b conforms to the BiDi rule.
func Valid(b []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:170
	_go_fuzz_dep_.CoverTab[32854]++
											var t Transformer
											if n, ok := t.advance(b); !ok || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:172
		_go_fuzz_dep_.CoverTab[32856]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:172
		return n < len(b)
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:172
		// _ = "end of CoverTab[32856]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:172
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:172
		_go_fuzz_dep_.CoverTab[32857]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:173
		// _ = "end of CoverTab[32857]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:174
		_go_fuzz_dep_.CoverTab[32858]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:174
		// _ = "end of CoverTab[32858]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:174
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:174
	// _ = "end of CoverTab[32854]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:174
	_go_fuzz_dep_.CoverTab[32855]++
											return t.isFinal()
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:175
	// _ = "end of CoverTab[32855]"
}

// ValidString reports whether s conforms to the BiDi rule.
func ValidString(s string) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:179
	_go_fuzz_dep_.CoverTab[32859]++
											var t Transformer
											if n, ok := t.advanceString(s); !ok || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:181
		_go_fuzz_dep_.CoverTab[32861]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:181
		return n < len(s)
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:181
		// _ = "end of CoverTab[32861]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:181
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:181
		_go_fuzz_dep_.CoverTab[32862]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:182
		// _ = "end of CoverTab[32862]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:183
		_go_fuzz_dep_.CoverTab[32863]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:183
		// _ = "end of CoverTab[32863]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:183
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:183
	// _ = "end of CoverTab[32859]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:183
	_go_fuzz_dep_.CoverTab[32860]++
											return t.isFinal()
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:184
	// _ = "end of CoverTab[32860]"
}

// New returns a Transformer that verifies that input adheres to the Bidi Rule.
func New() *Transformer {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:188
	_go_fuzz_dep_.CoverTab[32864]++
											return &Transformer{}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:189
	// _ = "end of CoverTab[32864]"
}

// Transformer implements transform.Transform.
type Transformer struct {
	state	ruleState
	hasRTL	bool
	seen	uint16
}

// A rule can only be violated for "Bidi Domain names", meaning if one of the
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:199
// following categories has been observed.
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:201
func (t *Transformer) isRTL() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:201
	_go_fuzz_dep_.CoverTab[32865]++
											const isRTL = 1<<bidi.R | 1<<bidi.AL | 1<<bidi.AN
											return t.seen&isRTL != 0
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:203
	// _ = "end of CoverTab[32865]"
}

// Reset implements transform.Transformer.
func (t *Transformer) Reset() {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:207
	_go_fuzz_dep_.CoverTab[32866]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:207
	*t = Transformer{}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:207
	// _ = "end of CoverTab[32866]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:207
}

// Transform implements transform.Transformer. This Transformer has state and
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:209
// needs to be reset between uses.
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:211
func (t *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:211
	_go_fuzz_dep_.CoverTab[32867]++
											if len(dst) < len(src) {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:212
		_go_fuzz_dep_.CoverTab[32870]++
												src = src[:len(dst)]
												atEOF = false
												err = transform.ErrShortDst
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:215
		// _ = "end of CoverTab[32870]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:216
		_go_fuzz_dep_.CoverTab[32871]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:216
		// _ = "end of CoverTab[32871]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:216
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:216
	// _ = "end of CoverTab[32867]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:216
	_go_fuzz_dep_.CoverTab[32868]++
											n, err1 := t.Span(src, atEOF)
											copy(dst, src[:n])
											if err == nil || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:219
		_go_fuzz_dep_.CoverTab[32872]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:219
		return err1 != nil && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:219
			_go_fuzz_dep_.CoverTab[32873]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:219
			return err1 != transform.ErrShortSrc
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:219
			// _ = "end of CoverTab[32873]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:219
		}()
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:219
		// _ = "end of CoverTab[32872]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:219
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:219
		_go_fuzz_dep_.CoverTab[32874]++
												err = err1
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:220
		// _ = "end of CoverTab[32874]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:221
		_go_fuzz_dep_.CoverTab[32875]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:221
		// _ = "end of CoverTab[32875]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:221
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:221
	// _ = "end of CoverTab[32868]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:221
	_go_fuzz_dep_.CoverTab[32869]++
											return n, n, err
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:222
	// _ = "end of CoverTab[32869]"
}

// Span returns the first n bytes of src that conform to the Bidi rule.
func (t *Transformer) Span(src []byte, atEOF bool) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:226
	_go_fuzz_dep_.CoverTab[32876]++
											if t.state == ruleInvalid && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:227
		_go_fuzz_dep_.CoverTab[32879]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:227
		return t.isRTL()
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:227
		// _ = "end of CoverTab[32879]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:227
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:227
		_go_fuzz_dep_.CoverTab[32880]++
												return 0, ErrInvalid
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:228
		// _ = "end of CoverTab[32880]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:229
		_go_fuzz_dep_.CoverTab[32881]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:229
		// _ = "end of CoverTab[32881]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:229
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:229
	// _ = "end of CoverTab[32876]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:229
	_go_fuzz_dep_.CoverTab[32877]++
											n, ok := t.advance(src)
											switch {
	case !ok:
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:232
		_go_fuzz_dep_.CoverTab[32882]++
												err = ErrInvalid
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:233
		// _ = "end of CoverTab[32882]"
	case n < len(src):
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:234
		_go_fuzz_dep_.CoverTab[32883]++
												if !atEOF {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:235
			_go_fuzz_dep_.CoverTab[32887]++
													err = transform.ErrShortSrc
													break
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:237
			// _ = "end of CoverTab[32887]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:238
			_go_fuzz_dep_.CoverTab[32888]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:238
			// _ = "end of CoverTab[32888]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:238
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:238
		// _ = "end of CoverTab[32883]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:238
		_go_fuzz_dep_.CoverTab[32884]++
												err = ErrInvalid
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:239
		// _ = "end of CoverTab[32884]"
	case !t.isFinal():
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:240
		_go_fuzz_dep_.CoverTab[32885]++
												err = ErrInvalid
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:241
		// _ = "end of CoverTab[32885]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:241
	default:
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:241
		_go_fuzz_dep_.CoverTab[32886]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:241
		// _ = "end of CoverTab[32886]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:242
	// _ = "end of CoverTab[32877]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:242
	_go_fuzz_dep_.CoverTab[32878]++
											return n, err
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:243
	// _ = "end of CoverTab[32878]"
}

// Precomputing the ASCII values decreases running time for the ASCII fast path
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:246
// by about 30%.
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:248
var asciiTable [128]bidi.Properties

func init() {
	for i := range asciiTable {
		p, _ := bidi.LookupRune(rune(i))
		asciiTable[i] = p
	}
}

func (t *Transformer) advance(s []byte) (n int, ok bool) {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:257
	_go_fuzz_dep_.CoverTab[32889]++
											var e bidi.Properties
											var sz int
											for n < len(s) {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:260
		_go_fuzz_dep_.CoverTab[32891]++
												if s[n] < utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:261
			_go_fuzz_dep_.CoverTab[32895]++
													e, sz = asciiTable[s[n]], 1
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:262
			// _ = "end of CoverTab[32895]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:263
			_go_fuzz_dep_.CoverTab[32896]++
													e, sz = bidi.Lookup(s[n:])
													if sz <= 1 {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:265
				_go_fuzz_dep_.CoverTab[32897]++
														if sz == 1 {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:266
					_go_fuzz_dep_.CoverTab[32899]++

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:270
					return n, false
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:270
					// _ = "end of CoverTab[32899]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:271
					_go_fuzz_dep_.CoverTab[32900]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:271
					// _ = "end of CoverTab[32900]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:271
				}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:271
				// _ = "end of CoverTab[32897]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:271
				_go_fuzz_dep_.CoverTab[32898]++
														return n, true
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:272
				// _ = "end of CoverTab[32898]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:273
				_go_fuzz_dep_.CoverTab[32901]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:273
				// _ = "end of CoverTab[32901]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:273
			}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:273
			// _ = "end of CoverTab[32896]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:274
		// _ = "end of CoverTab[32891]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:274
		_go_fuzz_dep_.CoverTab[32892]++

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:277
		c := uint16(1 << e.Class())
		t.seen |= c
		if t.seen&exclusiveRTL == exclusiveRTL {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:279
			_go_fuzz_dep_.CoverTab[32902]++
													t.state = ruleInvalid
													return n, false
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:281
			// _ = "end of CoverTab[32902]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:282
			_go_fuzz_dep_.CoverTab[32903]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:282
			// _ = "end of CoverTab[32903]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:282
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:282
		// _ = "end of CoverTab[32892]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:282
		_go_fuzz_dep_.CoverTab[32893]++
												switch tr := transitions[t.state]; {
		case tr[0].mask&c != 0:
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:284
			_go_fuzz_dep_.CoverTab[32904]++
													t.state = tr[0].next
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:285
			// _ = "end of CoverTab[32904]"
		case tr[1].mask&c != 0:
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:286
			_go_fuzz_dep_.CoverTab[32905]++
													t.state = tr[1].next
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:287
			// _ = "end of CoverTab[32905]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:288
			_go_fuzz_dep_.CoverTab[32906]++
													t.state = ruleInvalid
													if t.isRTL() {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:290
				_go_fuzz_dep_.CoverTab[32907]++
														return n, false
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:291
				// _ = "end of CoverTab[32907]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:292
				_go_fuzz_dep_.CoverTab[32908]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:292
				// _ = "end of CoverTab[32908]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:292
			}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:292
			// _ = "end of CoverTab[32906]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:293
		// _ = "end of CoverTab[32893]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:293
		_go_fuzz_dep_.CoverTab[32894]++
												n += sz
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:294
		// _ = "end of CoverTab[32894]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:295
	// _ = "end of CoverTab[32889]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:295
	_go_fuzz_dep_.CoverTab[32890]++
											return n, true
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:296
	// _ = "end of CoverTab[32890]"
}

func (t *Transformer) advanceString(s string) (n int, ok bool) {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:299
	_go_fuzz_dep_.CoverTab[32909]++
											var e bidi.Properties
											var sz int
											for n < len(s) {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:302
		_go_fuzz_dep_.CoverTab[32911]++
												if s[n] < utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:303
			_go_fuzz_dep_.CoverTab[32915]++
													e, sz = asciiTable[s[n]], 1
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:304
			// _ = "end of CoverTab[32915]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:305
			_go_fuzz_dep_.CoverTab[32916]++
													e, sz = bidi.LookupString(s[n:])
													if sz <= 1 {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:307
				_go_fuzz_dep_.CoverTab[32917]++
														if sz == 1 {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:308
					_go_fuzz_dep_.CoverTab[32919]++
															return n, false
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:309
					// _ = "end of CoverTab[32919]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:310
					_go_fuzz_dep_.CoverTab[32920]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:310
					// _ = "end of CoverTab[32920]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:310
				}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:310
				// _ = "end of CoverTab[32917]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:310
				_go_fuzz_dep_.CoverTab[32918]++
														return n, true
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:311
				// _ = "end of CoverTab[32918]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:312
				_go_fuzz_dep_.CoverTab[32921]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:312
				// _ = "end of CoverTab[32921]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:312
			}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:312
			// _ = "end of CoverTab[32916]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:313
		// _ = "end of CoverTab[32911]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:313
		_go_fuzz_dep_.CoverTab[32912]++

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:316
		c := uint16(1 << e.Class())
		t.seen |= c
		if t.seen&exclusiveRTL == exclusiveRTL {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:318
			_go_fuzz_dep_.CoverTab[32922]++
													t.state = ruleInvalid
													return n, false
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:320
			// _ = "end of CoverTab[32922]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:321
			_go_fuzz_dep_.CoverTab[32923]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:321
			// _ = "end of CoverTab[32923]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:321
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:321
		// _ = "end of CoverTab[32912]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:321
		_go_fuzz_dep_.CoverTab[32913]++
												switch tr := transitions[t.state]; {
		case tr[0].mask&c != 0:
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:323
			_go_fuzz_dep_.CoverTab[32924]++
													t.state = tr[0].next
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:324
			// _ = "end of CoverTab[32924]"
		case tr[1].mask&c != 0:
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:325
			_go_fuzz_dep_.CoverTab[32925]++
													t.state = tr[1].next
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:326
			// _ = "end of CoverTab[32925]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:327
			_go_fuzz_dep_.CoverTab[32926]++
													t.state = ruleInvalid
													if t.isRTL() {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:329
				_go_fuzz_dep_.CoverTab[32927]++
														return n, false
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:330
				// _ = "end of CoverTab[32927]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:331
				_go_fuzz_dep_.CoverTab[32928]++
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:331
				// _ = "end of CoverTab[32928]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:331
			}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:331
			// _ = "end of CoverTab[32926]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:332
		// _ = "end of CoverTab[32913]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:332
		_go_fuzz_dep_.CoverTab[32914]++
												n += sz
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:333
		// _ = "end of CoverTab[32914]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:334
	// _ = "end of CoverTab[32909]"
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:334
	_go_fuzz_dep_.CoverTab[32910]++
											return n, true
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:335
	// _ = "end of CoverTab[32910]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:336
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/secure/bidirule/bidirule.go:336
var _ = _go_fuzz_dep_.CoverTab
