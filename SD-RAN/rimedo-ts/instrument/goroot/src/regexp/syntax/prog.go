// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/regexp/syntax/prog.go:5
package syntax

//line /usr/local/go/src/regexp/syntax/prog.go:5
import (
//line /usr/local/go/src/regexp/syntax/prog.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/regexp/syntax/prog.go:5
)
//line /usr/local/go/src/regexp/syntax/prog.go:5
import (
//line /usr/local/go/src/regexp/syntax/prog.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/regexp/syntax/prog.go:5
)

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

//line /usr/local/go/src/regexp/syntax/prog.go:17
// A Prog is a compiled regular expression program.
type Prog struct {
	Inst	[]Inst
	Start	int	// index of start instruction
	NumCap	int	// number of InstCapture insts in re
}

// An InstOp is an instruction opcode.
type InstOp uint8

const (
	InstAlt	InstOp	= iota
	InstAltMatch
	InstCapture
	InstEmptyWidth
	InstMatch
	InstFail
	InstNop
	InstRune
	InstRune1
	InstRuneAny
	InstRuneAnyNotNL
)

var instOpNames = []string{
	"InstAlt",
	"InstAltMatch",
	"InstCapture",
	"InstEmptyWidth",
	"InstMatch",
	"InstFail",
	"InstNop",
	"InstRune",
	"InstRune1",
	"InstRuneAny",
	"InstRuneAnyNotNL",
}

func (i InstOp) String() string {
//line /usr/local/go/src/regexp/syntax/prog.go:55
	_go_fuzz_dep_.CoverTab[64063]++
							if uint(i) >= uint(len(instOpNames)) {
//line /usr/local/go/src/regexp/syntax/prog.go:56
		_go_fuzz_dep_.CoverTab[64065]++
								return ""
//line /usr/local/go/src/regexp/syntax/prog.go:57
		// _ = "end of CoverTab[64065]"
	} else {
//line /usr/local/go/src/regexp/syntax/prog.go:58
		_go_fuzz_dep_.CoverTab[64066]++
//line /usr/local/go/src/regexp/syntax/prog.go:58
		// _ = "end of CoverTab[64066]"
//line /usr/local/go/src/regexp/syntax/prog.go:58
	}
//line /usr/local/go/src/regexp/syntax/prog.go:58
	// _ = "end of CoverTab[64063]"
//line /usr/local/go/src/regexp/syntax/prog.go:58
	_go_fuzz_dep_.CoverTab[64064]++
							return instOpNames[i]
//line /usr/local/go/src/regexp/syntax/prog.go:59
	// _ = "end of CoverTab[64064]"
}

// An EmptyOp specifies a kind or mixture of zero-width assertions.
type EmptyOp uint8

const (
	EmptyBeginLine	EmptyOp	= 1 << iota
	EmptyEndLine
	EmptyBeginText
	EmptyEndText
	EmptyWordBoundary
	EmptyNoWordBoundary
)

// EmptyOpContext returns the zero-width assertions
//line /usr/local/go/src/regexp/syntax/prog.go:74
// satisfied at the position between the runes r1 and r2.
//line /usr/local/go/src/regexp/syntax/prog.go:74
// Passing r1 == -1 indicates that the position is
//line /usr/local/go/src/regexp/syntax/prog.go:74
// at the beginning of the text.
//line /usr/local/go/src/regexp/syntax/prog.go:74
// Passing r2 == -1 indicates that the position is
//line /usr/local/go/src/regexp/syntax/prog.go:74
// at the end of the text.
//line /usr/local/go/src/regexp/syntax/prog.go:80
func EmptyOpContext(r1, r2 rune) EmptyOp {
//line /usr/local/go/src/regexp/syntax/prog.go:80
	_go_fuzz_dep_.CoverTab[64067]++
							var op EmptyOp = EmptyNoWordBoundary
							var boundary byte
							switch {
	case IsWordChar(r1):
//line /usr/local/go/src/regexp/syntax/prog.go:84
		_go_fuzz_dep_.CoverTab[64071]++
								boundary = 1
//line /usr/local/go/src/regexp/syntax/prog.go:85
		// _ = "end of CoverTab[64071]"
	case r1 == '\n':
//line /usr/local/go/src/regexp/syntax/prog.go:86
		_go_fuzz_dep_.CoverTab[64072]++
								op |= EmptyBeginLine
//line /usr/local/go/src/regexp/syntax/prog.go:87
		// _ = "end of CoverTab[64072]"
	case r1 < 0:
//line /usr/local/go/src/regexp/syntax/prog.go:88
		_go_fuzz_dep_.CoverTab[64073]++
								op |= EmptyBeginText | EmptyBeginLine
//line /usr/local/go/src/regexp/syntax/prog.go:89
		// _ = "end of CoverTab[64073]"
//line /usr/local/go/src/regexp/syntax/prog.go:89
	default:
//line /usr/local/go/src/regexp/syntax/prog.go:89
		_go_fuzz_dep_.CoverTab[64074]++
//line /usr/local/go/src/regexp/syntax/prog.go:89
		// _ = "end of CoverTab[64074]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:90
	// _ = "end of CoverTab[64067]"
//line /usr/local/go/src/regexp/syntax/prog.go:90
	_go_fuzz_dep_.CoverTab[64068]++
							switch {
	case IsWordChar(r2):
//line /usr/local/go/src/regexp/syntax/prog.go:92
		_go_fuzz_dep_.CoverTab[64075]++
								boundary ^= 1
//line /usr/local/go/src/regexp/syntax/prog.go:93
		// _ = "end of CoverTab[64075]"
	case r2 == '\n':
//line /usr/local/go/src/regexp/syntax/prog.go:94
		_go_fuzz_dep_.CoverTab[64076]++
								op |= EmptyEndLine
//line /usr/local/go/src/regexp/syntax/prog.go:95
		// _ = "end of CoverTab[64076]"
	case r2 < 0:
//line /usr/local/go/src/regexp/syntax/prog.go:96
		_go_fuzz_dep_.CoverTab[64077]++
								op |= EmptyEndText | EmptyEndLine
//line /usr/local/go/src/regexp/syntax/prog.go:97
		// _ = "end of CoverTab[64077]"
//line /usr/local/go/src/regexp/syntax/prog.go:97
	default:
//line /usr/local/go/src/regexp/syntax/prog.go:97
		_go_fuzz_dep_.CoverTab[64078]++
//line /usr/local/go/src/regexp/syntax/prog.go:97
		// _ = "end of CoverTab[64078]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:98
	// _ = "end of CoverTab[64068]"
//line /usr/local/go/src/regexp/syntax/prog.go:98
	_go_fuzz_dep_.CoverTab[64069]++
							if boundary != 0 {
//line /usr/local/go/src/regexp/syntax/prog.go:99
		_go_fuzz_dep_.CoverTab[64079]++
								op ^= (EmptyWordBoundary | EmptyNoWordBoundary)
//line /usr/local/go/src/regexp/syntax/prog.go:100
		// _ = "end of CoverTab[64079]"
	} else {
//line /usr/local/go/src/regexp/syntax/prog.go:101
		_go_fuzz_dep_.CoverTab[64080]++
//line /usr/local/go/src/regexp/syntax/prog.go:101
		// _ = "end of CoverTab[64080]"
//line /usr/local/go/src/regexp/syntax/prog.go:101
	}
//line /usr/local/go/src/regexp/syntax/prog.go:101
	// _ = "end of CoverTab[64069]"
//line /usr/local/go/src/regexp/syntax/prog.go:101
	_go_fuzz_dep_.CoverTab[64070]++
							return op
//line /usr/local/go/src/regexp/syntax/prog.go:102
	// _ = "end of CoverTab[64070]"
}

// IsWordChar reports whether r is considered a “word character”
//line /usr/local/go/src/regexp/syntax/prog.go:105
// during the evaluation of the \b and \B zero-width assertions.
//line /usr/local/go/src/regexp/syntax/prog.go:105
// These assertions are ASCII-only: the word characters are [A-Za-z0-9_].
//line /usr/local/go/src/regexp/syntax/prog.go:108
func IsWordChar(r rune) bool {
//line /usr/local/go/src/regexp/syntax/prog.go:108
	_go_fuzz_dep_.CoverTab[64081]++
							return 'A' <= r && func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:109
		_go_fuzz_dep_.CoverTab[64082]++
//line /usr/local/go/src/regexp/syntax/prog.go:109
		return r <= 'Z'
//line /usr/local/go/src/regexp/syntax/prog.go:109
		// _ = "end of CoverTab[64082]"
//line /usr/local/go/src/regexp/syntax/prog.go:109
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:109
		_go_fuzz_dep_.CoverTab[64083]++
//line /usr/local/go/src/regexp/syntax/prog.go:109
		return 'a' <= r && func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:109
			_go_fuzz_dep_.CoverTab[64084]++
//line /usr/local/go/src/regexp/syntax/prog.go:109
			return r <= 'z'
//line /usr/local/go/src/regexp/syntax/prog.go:109
			// _ = "end of CoverTab[64084]"
//line /usr/local/go/src/regexp/syntax/prog.go:109
		}()
//line /usr/local/go/src/regexp/syntax/prog.go:109
		// _ = "end of CoverTab[64083]"
//line /usr/local/go/src/regexp/syntax/prog.go:109
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:109
		_go_fuzz_dep_.CoverTab[64085]++
//line /usr/local/go/src/regexp/syntax/prog.go:109
		return '0' <= r && func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:109
			_go_fuzz_dep_.CoverTab[64086]++
//line /usr/local/go/src/regexp/syntax/prog.go:109
			return r <= '9'
//line /usr/local/go/src/regexp/syntax/prog.go:109
			// _ = "end of CoverTab[64086]"
//line /usr/local/go/src/regexp/syntax/prog.go:109
		}()
//line /usr/local/go/src/regexp/syntax/prog.go:109
		// _ = "end of CoverTab[64085]"
//line /usr/local/go/src/regexp/syntax/prog.go:109
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:109
		_go_fuzz_dep_.CoverTab[64087]++
//line /usr/local/go/src/regexp/syntax/prog.go:109
		return r == '_'
//line /usr/local/go/src/regexp/syntax/prog.go:109
		// _ = "end of CoverTab[64087]"
//line /usr/local/go/src/regexp/syntax/prog.go:109
	}()
//line /usr/local/go/src/regexp/syntax/prog.go:109
	// _ = "end of CoverTab[64081]"
}

// An Inst is a single instruction in a regular expression program.
type Inst struct {
	Op	InstOp
	Out	uint32	// all but InstMatch, InstFail
	Arg	uint32	// InstAlt, InstAltMatch, InstCapture, InstEmptyWidth
	Rune	[]rune
}

func (p *Prog) String() string {
//line /usr/local/go/src/regexp/syntax/prog.go:120
	_go_fuzz_dep_.CoverTab[64088]++
							var b strings.Builder
							dumpProg(&b, p)
							return b.String()
//line /usr/local/go/src/regexp/syntax/prog.go:123
	// _ = "end of CoverTab[64088]"
}

// skipNop follows any no-op or capturing instructions.
func (p *Prog) skipNop(pc uint32) *Inst {
//line /usr/local/go/src/regexp/syntax/prog.go:127
	_go_fuzz_dep_.CoverTab[64089]++
							i := &p.Inst[pc]
							for i.Op == InstNop || func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:129
		_go_fuzz_dep_.CoverTab[64091]++
//line /usr/local/go/src/regexp/syntax/prog.go:129
		return i.Op == InstCapture
//line /usr/local/go/src/regexp/syntax/prog.go:129
		// _ = "end of CoverTab[64091]"
//line /usr/local/go/src/regexp/syntax/prog.go:129
	}() {
//line /usr/local/go/src/regexp/syntax/prog.go:129
		_go_fuzz_dep_.CoverTab[64092]++
								i = &p.Inst[i.Out]
//line /usr/local/go/src/regexp/syntax/prog.go:130
		// _ = "end of CoverTab[64092]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:131
	// _ = "end of CoverTab[64089]"
//line /usr/local/go/src/regexp/syntax/prog.go:131
	_go_fuzz_dep_.CoverTab[64090]++
							return i
//line /usr/local/go/src/regexp/syntax/prog.go:132
	// _ = "end of CoverTab[64090]"
}

// op returns i.Op but merges all the Rune special cases into InstRune
func (i *Inst) op() InstOp {
//line /usr/local/go/src/regexp/syntax/prog.go:136
	_go_fuzz_dep_.CoverTab[64093]++
							op := i.Op
							switch op {
	case InstRune1, InstRuneAny, InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/syntax/prog.go:139
		_go_fuzz_dep_.CoverTab[64095]++
								op = InstRune
//line /usr/local/go/src/regexp/syntax/prog.go:140
		// _ = "end of CoverTab[64095]"
//line /usr/local/go/src/regexp/syntax/prog.go:140
	default:
//line /usr/local/go/src/regexp/syntax/prog.go:140
		_go_fuzz_dep_.CoverTab[64096]++
//line /usr/local/go/src/regexp/syntax/prog.go:140
		// _ = "end of CoverTab[64096]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:141
	// _ = "end of CoverTab[64093]"
//line /usr/local/go/src/regexp/syntax/prog.go:141
	_go_fuzz_dep_.CoverTab[64094]++
							return op
//line /usr/local/go/src/regexp/syntax/prog.go:142
	// _ = "end of CoverTab[64094]"
}

// Prefix returns a literal string that all matches for the
//line /usr/local/go/src/regexp/syntax/prog.go:145
// regexp must start with. Complete is true if the prefix
//line /usr/local/go/src/regexp/syntax/prog.go:145
// is the entire match.
//line /usr/local/go/src/regexp/syntax/prog.go:148
func (p *Prog) Prefix() (prefix string, complete bool) {
//line /usr/local/go/src/regexp/syntax/prog.go:148
	_go_fuzz_dep_.CoverTab[64097]++
							i := p.skipNop(uint32(p.Start))

//line /usr/local/go/src/regexp/syntax/prog.go:152
	if i.op() != InstRune || func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:152
		_go_fuzz_dep_.CoverTab[64100]++
//line /usr/local/go/src/regexp/syntax/prog.go:152
		return len(i.Rune) != 1
//line /usr/local/go/src/regexp/syntax/prog.go:152
		// _ = "end of CoverTab[64100]"
//line /usr/local/go/src/regexp/syntax/prog.go:152
	}() {
//line /usr/local/go/src/regexp/syntax/prog.go:152
		_go_fuzz_dep_.CoverTab[64101]++
								return "", i.Op == InstMatch
//line /usr/local/go/src/regexp/syntax/prog.go:153
		// _ = "end of CoverTab[64101]"
	} else {
//line /usr/local/go/src/regexp/syntax/prog.go:154
		_go_fuzz_dep_.CoverTab[64102]++
//line /usr/local/go/src/regexp/syntax/prog.go:154
		// _ = "end of CoverTab[64102]"
//line /usr/local/go/src/regexp/syntax/prog.go:154
	}
//line /usr/local/go/src/regexp/syntax/prog.go:154
	// _ = "end of CoverTab[64097]"
//line /usr/local/go/src/regexp/syntax/prog.go:154
	_go_fuzz_dep_.CoverTab[64098]++

	// Have prefix; gather characters.
	var buf strings.Builder
	for i.op() == InstRune && func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:158
		_go_fuzz_dep_.CoverTab[64103]++
//line /usr/local/go/src/regexp/syntax/prog.go:158
		return len(i.Rune) == 1
//line /usr/local/go/src/regexp/syntax/prog.go:158
		// _ = "end of CoverTab[64103]"
//line /usr/local/go/src/regexp/syntax/prog.go:158
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:158
		_go_fuzz_dep_.CoverTab[64104]++
//line /usr/local/go/src/regexp/syntax/prog.go:158
		return Flags(i.Arg)&FoldCase == 0
//line /usr/local/go/src/regexp/syntax/prog.go:158
		// _ = "end of CoverTab[64104]"
//line /usr/local/go/src/regexp/syntax/prog.go:158
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:158
		_go_fuzz_dep_.CoverTab[64105]++
//line /usr/local/go/src/regexp/syntax/prog.go:158
		return i.Rune[0] != utf8.RuneError
//line /usr/local/go/src/regexp/syntax/prog.go:158
		// _ = "end of CoverTab[64105]"
//line /usr/local/go/src/regexp/syntax/prog.go:158
	}() {
//line /usr/local/go/src/regexp/syntax/prog.go:158
		_go_fuzz_dep_.CoverTab[64106]++
								buf.WriteRune(i.Rune[0])
								i = p.skipNop(i.Out)
//line /usr/local/go/src/regexp/syntax/prog.go:160
		// _ = "end of CoverTab[64106]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:161
	// _ = "end of CoverTab[64098]"
//line /usr/local/go/src/regexp/syntax/prog.go:161
	_go_fuzz_dep_.CoverTab[64099]++
							return buf.String(), i.Op == InstMatch
//line /usr/local/go/src/regexp/syntax/prog.go:162
	// _ = "end of CoverTab[64099]"
}

// StartCond returns the leading empty-width conditions that must
//line /usr/local/go/src/regexp/syntax/prog.go:165
// be true in any match. It returns ^EmptyOp(0) if no matches are possible.
//line /usr/local/go/src/regexp/syntax/prog.go:167
func (p *Prog) StartCond() EmptyOp {
//line /usr/local/go/src/regexp/syntax/prog.go:167
	_go_fuzz_dep_.CoverTab[64107]++
							var flag EmptyOp
							pc := uint32(p.Start)
							i := &p.Inst[pc]
Loop:
	for {
//line /usr/local/go/src/regexp/syntax/prog.go:172
		_go_fuzz_dep_.CoverTab[64109]++
								switch i.Op {
		case InstEmptyWidth:
//line /usr/local/go/src/regexp/syntax/prog.go:174
			_go_fuzz_dep_.CoverTab[64111]++
									flag |= EmptyOp(i.Arg)
//line /usr/local/go/src/regexp/syntax/prog.go:175
			// _ = "end of CoverTab[64111]"
		case InstFail:
//line /usr/local/go/src/regexp/syntax/prog.go:176
			_go_fuzz_dep_.CoverTab[64112]++
									return ^EmptyOp(0)
//line /usr/local/go/src/regexp/syntax/prog.go:177
			// _ = "end of CoverTab[64112]"
		case InstCapture, InstNop:
//line /usr/local/go/src/regexp/syntax/prog.go:178
			_go_fuzz_dep_.CoverTab[64113]++
//line /usr/local/go/src/regexp/syntax/prog.go:178
			// _ = "end of CoverTab[64113]"

		default:
//line /usr/local/go/src/regexp/syntax/prog.go:180
			_go_fuzz_dep_.CoverTab[64114]++
									break Loop
//line /usr/local/go/src/regexp/syntax/prog.go:181
			// _ = "end of CoverTab[64114]"
		}
//line /usr/local/go/src/regexp/syntax/prog.go:182
		// _ = "end of CoverTab[64109]"
//line /usr/local/go/src/regexp/syntax/prog.go:182
		_go_fuzz_dep_.CoverTab[64110]++
								pc = i.Out
								i = &p.Inst[pc]
//line /usr/local/go/src/regexp/syntax/prog.go:184
		// _ = "end of CoverTab[64110]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:185
	// _ = "end of CoverTab[64107]"
//line /usr/local/go/src/regexp/syntax/prog.go:185
	_go_fuzz_dep_.CoverTab[64108]++
							return flag
//line /usr/local/go/src/regexp/syntax/prog.go:186
	// _ = "end of CoverTab[64108]"
}

const noMatch = -1

// MatchRune reports whether the instruction matches (and consumes) r.
//line /usr/local/go/src/regexp/syntax/prog.go:191
// It should only be called when i.Op == InstRune.
//line /usr/local/go/src/regexp/syntax/prog.go:193
func (i *Inst) MatchRune(r rune) bool {
//line /usr/local/go/src/regexp/syntax/prog.go:193
	_go_fuzz_dep_.CoverTab[64115]++
							return i.MatchRunePos(r) != noMatch
//line /usr/local/go/src/regexp/syntax/prog.go:194
	// _ = "end of CoverTab[64115]"
}

// MatchRunePos checks whether the instruction matches (and consumes) r.
//line /usr/local/go/src/regexp/syntax/prog.go:197
// If so, MatchRunePos returns the index of the matching rune pair
//line /usr/local/go/src/regexp/syntax/prog.go:197
// (or, when len(i.Rune) == 1, rune singleton).
//line /usr/local/go/src/regexp/syntax/prog.go:197
// If not, MatchRunePos returns -1.
//line /usr/local/go/src/regexp/syntax/prog.go:197
// MatchRunePos should only be called when i.Op == InstRune.
//line /usr/local/go/src/regexp/syntax/prog.go:202
func (i *Inst) MatchRunePos(r rune) int {
//line /usr/local/go/src/regexp/syntax/prog.go:202
	_go_fuzz_dep_.CoverTab[64116]++
							rune := i.Rune

							switch len(rune) {
	case 0:
//line /usr/local/go/src/regexp/syntax/prog.go:206
		_go_fuzz_dep_.CoverTab[64119]++
								return noMatch
//line /usr/local/go/src/regexp/syntax/prog.go:207
		// _ = "end of CoverTab[64119]"

	case 1:
//line /usr/local/go/src/regexp/syntax/prog.go:209
		_go_fuzz_dep_.CoverTab[64120]++

								r0 := rune[0]
								if r == r0 {
//line /usr/local/go/src/regexp/syntax/prog.go:212
			_go_fuzz_dep_.CoverTab[64128]++
									return 0
//line /usr/local/go/src/regexp/syntax/prog.go:213
			// _ = "end of CoverTab[64128]"
		} else {
//line /usr/local/go/src/regexp/syntax/prog.go:214
			_go_fuzz_dep_.CoverTab[64129]++
//line /usr/local/go/src/regexp/syntax/prog.go:214
			// _ = "end of CoverTab[64129]"
//line /usr/local/go/src/regexp/syntax/prog.go:214
		}
//line /usr/local/go/src/regexp/syntax/prog.go:214
		// _ = "end of CoverTab[64120]"
//line /usr/local/go/src/regexp/syntax/prog.go:214
		_go_fuzz_dep_.CoverTab[64121]++
								if Flags(i.Arg)&FoldCase != 0 {
//line /usr/local/go/src/regexp/syntax/prog.go:215
			_go_fuzz_dep_.CoverTab[64130]++
									for r1 := unicode.SimpleFold(r0); r1 != r0; r1 = unicode.SimpleFold(r1) {
//line /usr/local/go/src/regexp/syntax/prog.go:216
				_go_fuzz_dep_.CoverTab[64131]++
										if r == r1 {
//line /usr/local/go/src/regexp/syntax/prog.go:217
					_go_fuzz_dep_.CoverTab[64132]++
											return 0
//line /usr/local/go/src/regexp/syntax/prog.go:218
					// _ = "end of CoverTab[64132]"
				} else {
//line /usr/local/go/src/regexp/syntax/prog.go:219
					_go_fuzz_dep_.CoverTab[64133]++
//line /usr/local/go/src/regexp/syntax/prog.go:219
					// _ = "end of CoverTab[64133]"
//line /usr/local/go/src/regexp/syntax/prog.go:219
				}
//line /usr/local/go/src/regexp/syntax/prog.go:219
				// _ = "end of CoverTab[64131]"
			}
//line /usr/local/go/src/regexp/syntax/prog.go:220
			// _ = "end of CoverTab[64130]"
		} else {
//line /usr/local/go/src/regexp/syntax/prog.go:221
			_go_fuzz_dep_.CoverTab[64134]++
//line /usr/local/go/src/regexp/syntax/prog.go:221
			// _ = "end of CoverTab[64134]"
//line /usr/local/go/src/regexp/syntax/prog.go:221
		}
//line /usr/local/go/src/regexp/syntax/prog.go:221
		// _ = "end of CoverTab[64121]"
//line /usr/local/go/src/regexp/syntax/prog.go:221
		_go_fuzz_dep_.CoverTab[64122]++
								return noMatch
//line /usr/local/go/src/regexp/syntax/prog.go:222
		// _ = "end of CoverTab[64122]"

	case 2:
//line /usr/local/go/src/regexp/syntax/prog.go:224
		_go_fuzz_dep_.CoverTab[64123]++
								if r >= rune[0] && func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:225
			_go_fuzz_dep_.CoverTab[64135]++
//line /usr/local/go/src/regexp/syntax/prog.go:225
			return r <= rune[1]
//line /usr/local/go/src/regexp/syntax/prog.go:225
			// _ = "end of CoverTab[64135]"
//line /usr/local/go/src/regexp/syntax/prog.go:225
		}() {
//line /usr/local/go/src/regexp/syntax/prog.go:225
			_go_fuzz_dep_.CoverTab[64136]++
									return 0
//line /usr/local/go/src/regexp/syntax/prog.go:226
			// _ = "end of CoverTab[64136]"
		} else {
//line /usr/local/go/src/regexp/syntax/prog.go:227
			_go_fuzz_dep_.CoverTab[64137]++
//line /usr/local/go/src/regexp/syntax/prog.go:227
			// _ = "end of CoverTab[64137]"
//line /usr/local/go/src/regexp/syntax/prog.go:227
		}
//line /usr/local/go/src/regexp/syntax/prog.go:227
		// _ = "end of CoverTab[64123]"
//line /usr/local/go/src/regexp/syntax/prog.go:227
		_go_fuzz_dep_.CoverTab[64124]++
								return noMatch
//line /usr/local/go/src/regexp/syntax/prog.go:228
		// _ = "end of CoverTab[64124]"

	case 4, 6, 8:
//line /usr/local/go/src/regexp/syntax/prog.go:230
		_go_fuzz_dep_.CoverTab[64125]++

//line /usr/local/go/src/regexp/syntax/prog.go:233
		for j := 0; j < len(rune); j += 2 {
//line /usr/local/go/src/regexp/syntax/prog.go:233
			_go_fuzz_dep_.CoverTab[64138]++
									if r < rune[j] {
//line /usr/local/go/src/regexp/syntax/prog.go:234
				_go_fuzz_dep_.CoverTab[64140]++
										return noMatch
//line /usr/local/go/src/regexp/syntax/prog.go:235
				// _ = "end of CoverTab[64140]"
			} else {
//line /usr/local/go/src/regexp/syntax/prog.go:236
				_go_fuzz_dep_.CoverTab[64141]++
//line /usr/local/go/src/regexp/syntax/prog.go:236
				// _ = "end of CoverTab[64141]"
//line /usr/local/go/src/regexp/syntax/prog.go:236
			}
//line /usr/local/go/src/regexp/syntax/prog.go:236
			// _ = "end of CoverTab[64138]"
//line /usr/local/go/src/regexp/syntax/prog.go:236
			_go_fuzz_dep_.CoverTab[64139]++
									if r <= rune[j+1] {
//line /usr/local/go/src/regexp/syntax/prog.go:237
				_go_fuzz_dep_.CoverTab[64142]++
										return j / 2
//line /usr/local/go/src/regexp/syntax/prog.go:238
				// _ = "end of CoverTab[64142]"
			} else {
//line /usr/local/go/src/regexp/syntax/prog.go:239
				_go_fuzz_dep_.CoverTab[64143]++
//line /usr/local/go/src/regexp/syntax/prog.go:239
				// _ = "end of CoverTab[64143]"
//line /usr/local/go/src/regexp/syntax/prog.go:239
			}
//line /usr/local/go/src/regexp/syntax/prog.go:239
			// _ = "end of CoverTab[64139]"
		}
//line /usr/local/go/src/regexp/syntax/prog.go:240
		// _ = "end of CoverTab[64125]"
//line /usr/local/go/src/regexp/syntax/prog.go:240
		_go_fuzz_dep_.CoverTab[64126]++
								return noMatch
//line /usr/local/go/src/regexp/syntax/prog.go:241
		// _ = "end of CoverTab[64126]"
//line /usr/local/go/src/regexp/syntax/prog.go:241
	default:
//line /usr/local/go/src/regexp/syntax/prog.go:241
		_go_fuzz_dep_.CoverTab[64127]++
//line /usr/local/go/src/regexp/syntax/prog.go:241
		// _ = "end of CoverTab[64127]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:242
	// _ = "end of CoverTab[64116]"
//line /usr/local/go/src/regexp/syntax/prog.go:242
	_go_fuzz_dep_.CoverTab[64117]++

//line /usr/local/go/src/regexp/syntax/prog.go:245
	lo := 0
	hi := len(rune) / 2
	for lo < hi {
//line /usr/local/go/src/regexp/syntax/prog.go:247
		_go_fuzz_dep_.CoverTab[64144]++
								m := lo + (hi-lo)/2
								if c := rune[2*m]; c <= r {
//line /usr/local/go/src/regexp/syntax/prog.go:249
			_go_fuzz_dep_.CoverTab[64145]++
									if r <= rune[2*m+1] {
//line /usr/local/go/src/regexp/syntax/prog.go:250
				_go_fuzz_dep_.CoverTab[64147]++
										return m
//line /usr/local/go/src/regexp/syntax/prog.go:251
				// _ = "end of CoverTab[64147]"
			} else {
//line /usr/local/go/src/regexp/syntax/prog.go:252
				_go_fuzz_dep_.CoverTab[64148]++
//line /usr/local/go/src/regexp/syntax/prog.go:252
				// _ = "end of CoverTab[64148]"
//line /usr/local/go/src/regexp/syntax/prog.go:252
			}
//line /usr/local/go/src/regexp/syntax/prog.go:252
			// _ = "end of CoverTab[64145]"
//line /usr/local/go/src/regexp/syntax/prog.go:252
			_go_fuzz_dep_.CoverTab[64146]++
									lo = m + 1
//line /usr/local/go/src/regexp/syntax/prog.go:253
			// _ = "end of CoverTab[64146]"
		} else {
//line /usr/local/go/src/regexp/syntax/prog.go:254
			_go_fuzz_dep_.CoverTab[64149]++
									hi = m
//line /usr/local/go/src/regexp/syntax/prog.go:255
			// _ = "end of CoverTab[64149]"
		}
//line /usr/local/go/src/regexp/syntax/prog.go:256
		// _ = "end of CoverTab[64144]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:257
	// _ = "end of CoverTab[64117]"
//line /usr/local/go/src/regexp/syntax/prog.go:257
	_go_fuzz_dep_.CoverTab[64118]++
							return noMatch
//line /usr/local/go/src/regexp/syntax/prog.go:258
	// _ = "end of CoverTab[64118]"
}

// MatchEmptyWidth reports whether the instruction matches
//line /usr/local/go/src/regexp/syntax/prog.go:261
// an empty string between the runes before and after.
//line /usr/local/go/src/regexp/syntax/prog.go:261
// It should only be called when i.Op == InstEmptyWidth.
//line /usr/local/go/src/regexp/syntax/prog.go:264
func (i *Inst) MatchEmptyWidth(before rune, after rune) bool {
//line /usr/local/go/src/regexp/syntax/prog.go:264
	_go_fuzz_dep_.CoverTab[64150]++
							switch EmptyOp(i.Arg) {
	case EmptyBeginLine:
//line /usr/local/go/src/regexp/syntax/prog.go:266
		_go_fuzz_dep_.CoverTab[64152]++
								return before == '\n' || func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:267
			_go_fuzz_dep_.CoverTab[64159]++
//line /usr/local/go/src/regexp/syntax/prog.go:267
			return before == -1
//line /usr/local/go/src/regexp/syntax/prog.go:267
			// _ = "end of CoverTab[64159]"
//line /usr/local/go/src/regexp/syntax/prog.go:267
		}()
//line /usr/local/go/src/regexp/syntax/prog.go:267
		// _ = "end of CoverTab[64152]"
	case EmptyEndLine:
//line /usr/local/go/src/regexp/syntax/prog.go:268
		_go_fuzz_dep_.CoverTab[64153]++
								return after == '\n' || func() bool {
//line /usr/local/go/src/regexp/syntax/prog.go:269
			_go_fuzz_dep_.CoverTab[64160]++
//line /usr/local/go/src/regexp/syntax/prog.go:269
			return after == -1
//line /usr/local/go/src/regexp/syntax/prog.go:269
			// _ = "end of CoverTab[64160]"
//line /usr/local/go/src/regexp/syntax/prog.go:269
		}()
//line /usr/local/go/src/regexp/syntax/prog.go:269
		// _ = "end of CoverTab[64153]"
	case EmptyBeginText:
//line /usr/local/go/src/regexp/syntax/prog.go:270
		_go_fuzz_dep_.CoverTab[64154]++
								return before == -1
//line /usr/local/go/src/regexp/syntax/prog.go:271
		// _ = "end of CoverTab[64154]"
	case EmptyEndText:
//line /usr/local/go/src/regexp/syntax/prog.go:272
		_go_fuzz_dep_.CoverTab[64155]++
								return after == -1
//line /usr/local/go/src/regexp/syntax/prog.go:273
		// _ = "end of CoverTab[64155]"
	case EmptyWordBoundary:
//line /usr/local/go/src/regexp/syntax/prog.go:274
		_go_fuzz_dep_.CoverTab[64156]++
								return IsWordChar(before) != IsWordChar(after)
//line /usr/local/go/src/regexp/syntax/prog.go:275
		// _ = "end of CoverTab[64156]"
	case EmptyNoWordBoundary:
//line /usr/local/go/src/regexp/syntax/prog.go:276
		_go_fuzz_dep_.CoverTab[64157]++
								return IsWordChar(before) == IsWordChar(after)
//line /usr/local/go/src/regexp/syntax/prog.go:277
		// _ = "end of CoverTab[64157]"
//line /usr/local/go/src/regexp/syntax/prog.go:277
	default:
//line /usr/local/go/src/regexp/syntax/prog.go:277
		_go_fuzz_dep_.CoverTab[64158]++
//line /usr/local/go/src/regexp/syntax/prog.go:277
		// _ = "end of CoverTab[64158]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:278
	// _ = "end of CoverTab[64150]"
//line /usr/local/go/src/regexp/syntax/prog.go:278
	_go_fuzz_dep_.CoverTab[64151]++
							panic("unknown empty width arg")
//line /usr/local/go/src/regexp/syntax/prog.go:279
	// _ = "end of CoverTab[64151]"
}

func (i *Inst) String() string {
//line /usr/local/go/src/regexp/syntax/prog.go:282
	_go_fuzz_dep_.CoverTab[64161]++
							var b strings.Builder
							dumpInst(&b, i)
							return b.String()
//line /usr/local/go/src/regexp/syntax/prog.go:285
	// _ = "end of CoverTab[64161]"
}

func bw(b *strings.Builder, args ...string) {
//line /usr/local/go/src/regexp/syntax/prog.go:288
	_go_fuzz_dep_.CoverTab[64162]++
							for _, s := range args {
//line /usr/local/go/src/regexp/syntax/prog.go:289
		_go_fuzz_dep_.CoverTab[64163]++
								b.WriteString(s)
//line /usr/local/go/src/regexp/syntax/prog.go:290
		// _ = "end of CoverTab[64163]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:291
	// _ = "end of CoverTab[64162]"
}

func dumpProg(b *strings.Builder, p *Prog) {
//line /usr/local/go/src/regexp/syntax/prog.go:294
	_go_fuzz_dep_.CoverTab[64164]++
							for j := range p.Inst {
//line /usr/local/go/src/regexp/syntax/prog.go:295
		_go_fuzz_dep_.CoverTab[64165]++
								i := &p.Inst[j]
								pc := strconv.Itoa(j)
								if len(pc) < 3 {
//line /usr/local/go/src/regexp/syntax/prog.go:298
			_go_fuzz_dep_.CoverTab[64168]++
									b.WriteString("   "[len(pc):])
//line /usr/local/go/src/regexp/syntax/prog.go:299
			// _ = "end of CoverTab[64168]"
		} else {
//line /usr/local/go/src/regexp/syntax/prog.go:300
			_go_fuzz_dep_.CoverTab[64169]++
//line /usr/local/go/src/regexp/syntax/prog.go:300
			// _ = "end of CoverTab[64169]"
//line /usr/local/go/src/regexp/syntax/prog.go:300
		}
//line /usr/local/go/src/regexp/syntax/prog.go:300
		// _ = "end of CoverTab[64165]"
//line /usr/local/go/src/regexp/syntax/prog.go:300
		_go_fuzz_dep_.CoverTab[64166]++
								if j == p.Start {
//line /usr/local/go/src/regexp/syntax/prog.go:301
			_go_fuzz_dep_.CoverTab[64170]++
									pc += "*"
//line /usr/local/go/src/regexp/syntax/prog.go:302
			// _ = "end of CoverTab[64170]"
		} else {
//line /usr/local/go/src/regexp/syntax/prog.go:303
			_go_fuzz_dep_.CoverTab[64171]++
//line /usr/local/go/src/regexp/syntax/prog.go:303
			// _ = "end of CoverTab[64171]"
//line /usr/local/go/src/regexp/syntax/prog.go:303
		}
//line /usr/local/go/src/regexp/syntax/prog.go:303
		// _ = "end of CoverTab[64166]"
//line /usr/local/go/src/regexp/syntax/prog.go:303
		_go_fuzz_dep_.CoverTab[64167]++
								bw(b, pc, "\t")
								dumpInst(b, i)
								bw(b, "\n")
//line /usr/local/go/src/regexp/syntax/prog.go:306
		// _ = "end of CoverTab[64167]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:307
	// _ = "end of CoverTab[64164]"
}

func u32(i uint32) string {
//line /usr/local/go/src/regexp/syntax/prog.go:310
	_go_fuzz_dep_.CoverTab[64172]++
							return strconv.FormatUint(uint64(i), 10)
//line /usr/local/go/src/regexp/syntax/prog.go:311
	// _ = "end of CoverTab[64172]"
}

func dumpInst(b *strings.Builder, i *Inst) {
//line /usr/local/go/src/regexp/syntax/prog.go:314
	_go_fuzz_dep_.CoverTab[64173]++
							switch i.Op {
	case InstAlt:
//line /usr/local/go/src/regexp/syntax/prog.go:316
		_go_fuzz_dep_.CoverTab[64174]++
								bw(b, "alt -> ", u32(i.Out), ", ", u32(i.Arg))
//line /usr/local/go/src/regexp/syntax/prog.go:317
		// _ = "end of CoverTab[64174]"
	case InstAltMatch:
//line /usr/local/go/src/regexp/syntax/prog.go:318
		_go_fuzz_dep_.CoverTab[64175]++
								bw(b, "altmatch -> ", u32(i.Out), ", ", u32(i.Arg))
//line /usr/local/go/src/regexp/syntax/prog.go:319
		// _ = "end of CoverTab[64175]"
	case InstCapture:
//line /usr/local/go/src/regexp/syntax/prog.go:320
		_go_fuzz_dep_.CoverTab[64176]++
								bw(b, "cap ", u32(i.Arg), " -> ", u32(i.Out))
//line /usr/local/go/src/regexp/syntax/prog.go:321
		// _ = "end of CoverTab[64176]"
	case InstEmptyWidth:
//line /usr/local/go/src/regexp/syntax/prog.go:322
		_go_fuzz_dep_.CoverTab[64177]++
								bw(b, "empty ", u32(i.Arg), " -> ", u32(i.Out))
//line /usr/local/go/src/regexp/syntax/prog.go:323
		// _ = "end of CoverTab[64177]"
	case InstMatch:
//line /usr/local/go/src/regexp/syntax/prog.go:324
		_go_fuzz_dep_.CoverTab[64178]++
								bw(b, "match")
//line /usr/local/go/src/regexp/syntax/prog.go:325
		// _ = "end of CoverTab[64178]"
	case InstFail:
//line /usr/local/go/src/regexp/syntax/prog.go:326
		_go_fuzz_dep_.CoverTab[64179]++
								bw(b, "fail")
//line /usr/local/go/src/regexp/syntax/prog.go:327
		// _ = "end of CoverTab[64179]"
	case InstNop:
//line /usr/local/go/src/regexp/syntax/prog.go:328
		_go_fuzz_dep_.CoverTab[64180]++
								bw(b, "nop -> ", u32(i.Out))
//line /usr/local/go/src/regexp/syntax/prog.go:329
		// _ = "end of CoverTab[64180]"
	case InstRune:
//line /usr/local/go/src/regexp/syntax/prog.go:330
		_go_fuzz_dep_.CoverTab[64181]++
								if i.Rune == nil {
//line /usr/local/go/src/regexp/syntax/prog.go:331
			_go_fuzz_dep_.CoverTab[64188]++

									bw(b, "rune <nil>")
//line /usr/local/go/src/regexp/syntax/prog.go:333
			// _ = "end of CoverTab[64188]"
		} else {
//line /usr/local/go/src/regexp/syntax/prog.go:334
			_go_fuzz_dep_.CoverTab[64189]++
//line /usr/local/go/src/regexp/syntax/prog.go:334
			// _ = "end of CoverTab[64189]"
//line /usr/local/go/src/regexp/syntax/prog.go:334
		}
//line /usr/local/go/src/regexp/syntax/prog.go:334
		// _ = "end of CoverTab[64181]"
//line /usr/local/go/src/regexp/syntax/prog.go:334
		_go_fuzz_dep_.CoverTab[64182]++
								bw(b, "rune ", strconv.QuoteToASCII(string(i.Rune)))
								if Flags(i.Arg)&FoldCase != 0 {
//line /usr/local/go/src/regexp/syntax/prog.go:336
			_go_fuzz_dep_.CoverTab[64190]++
									bw(b, "/i")
//line /usr/local/go/src/regexp/syntax/prog.go:337
			// _ = "end of CoverTab[64190]"
		} else {
//line /usr/local/go/src/regexp/syntax/prog.go:338
			_go_fuzz_dep_.CoverTab[64191]++
//line /usr/local/go/src/regexp/syntax/prog.go:338
			// _ = "end of CoverTab[64191]"
//line /usr/local/go/src/regexp/syntax/prog.go:338
		}
//line /usr/local/go/src/regexp/syntax/prog.go:338
		// _ = "end of CoverTab[64182]"
//line /usr/local/go/src/regexp/syntax/prog.go:338
		_go_fuzz_dep_.CoverTab[64183]++
								bw(b, " -> ", u32(i.Out))
//line /usr/local/go/src/regexp/syntax/prog.go:339
		// _ = "end of CoverTab[64183]"
	case InstRune1:
//line /usr/local/go/src/regexp/syntax/prog.go:340
		_go_fuzz_dep_.CoverTab[64184]++
								bw(b, "rune1 ", strconv.QuoteToASCII(string(i.Rune)), " -> ", u32(i.Out))
//line /usr/local/go/src/regexp/syntax/prog.go:341
		// _ = "end of CoverTab[64184]"
	case InstRuneAny:
//line /usr/local/go/src/regexp/syntax/prog.go:342
		_go_fuzz_dep_.CoverTab[64185]++
								bw(b, "any -> ", u32(i.Out))
//line /usr/local/go/src/regexp/syntax/prog.go:343
		// _ = "end of CoverTab[64185]"
	case InstRuneAnyNotNL:
//line /usr/local/go/src/regexp/syntax/prog.go:344
		_go_fuzz_dep_.CoverTab[64186]++
								bw(b, "anynotnl -> ", u32(i.Out))
//line /usr/local/go/src/regexp/syntax/prog.go:345
		// _ = "end of CoverTab[64186]"
//line /usr/local/go/src/regexp/syntax/prog.go:345
	default:
//line /usr/local/go/src/regexp/syntax/prog.go:345
		_go_fuzz_dep_.CoverTab[64187]++
//line /usr/local/go/src/regexp/syntax/prog.go:345
		// _ = "end of CoverTab[64187]"
	}
//line /usr/local/go/src/regexp/syntax/prog.go:346
	// _ = "end of CoverTab[64173]"
}

//line /usr/local/go/src/regexp/syntax/prog.go:347
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/regexp/syntax/prog.go:347
var _ = _go_fuzz_dep_.CoverTab
