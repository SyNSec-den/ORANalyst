// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/regexp/syntax/parse.go:5
package syntax

//line /usr/local/go/src/regexp/syntax/parse.go:5
import (
//line /usr/local/go/src/regexp/syntax/parse.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/regexp/syntax/parse.go:5
)
//line /usr/local/go/src/regexp/syntax/parse.go:5
import (
//line /usr/local/go/src/regexp/syntax/parse.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/regexp/syntax/parse.go:5
)

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// An Error describes a failure to parse a regular expression
//line /usr/local/go/src/regexp/syntax/parse.go:14
// and gives the offending expression.
//line /usr/local/go/src/regexp/syntax/parse.go:16
type Error struct {
	Code	ErrorCode
	Expr	string
}

func (e *Error) Error() string {
//line /usr/local/go/src/regexp/syntax/parse.go:21
	_go_fuzz_dep_.CoverTab[63070]++
							return "error parsing regexp: " + e.Code.String() + ": `" + e.Expr + "`"
//line /usr/local/go/src/regexp/syntax/parse.go:22
	// _ = "end of CoverTab[63070]"
}

// An ErrorCode describes a failure to parse a regular expression.
type ErrorCode string

const (
	// Unexpected error
	ErrInternalError	ErrorCode	= "regexp/syntax: internal error"

	// Parse errors
	ErrInvalidCharClass		ErrorCode	= "invalid character class"
	ErrInvalidCharRange		ErrorCode	= "invalid character class range"
	ErrInvalidEscape		ErrorCode	= "invalid escape sequence"
	ErrInvalidNamedCapture		ErrorCode	= "invalid named capture"
	ErrInvalidPerlOp		ErrorCode	= "invalid or unsupported Perl syntax"
	ErrInvalidRepeatOp		ErrorCode	= "invalid nested repetition operator"
	ErrInvalidRepeatSize		ErrorCode	= "invalid repeat count"
	ErrInvalidUTF8			ErrorCode	= "invalid UTF-8"
	ErrMissingBracket		ErrorCode	= "missing closing ]"
	ErrMissingParen			ErrorCode	= "missing closing )"
	ErrMissingRepeatArgument	ErrorCode	= "missing argument to repetition operator"
	ErrTrailingBackslash		ErrorCode	= "trailing backslash at end of expression"
	ErrUnexpectedParen		ErrorCode	= "unexpected )"
	ErrNestingDepth			ErrorCode	= "expression nests too deeply"
	ErrLarge			ErrorCode	= "expression too large"
)

func (e ErrorCode) String() string {
//line /usr/local/go/src/regexp/syntax/parse.go:50
	_go_fuzz_dep_.CoverTab[63071]++
							return string(e)
//line /usr/local/go/src/regexp/syntax/parse.go:51
	// _ = "end of CoverTab[63071]"
}

// Flags control the behavior of the parser and record information about regexp context.
type Flags uint16

const (
	FoldCase	Flags	= 1 << iota	// case-insensitive match
	Literal					// treat pattern as literal string
	ClassNL					// allow character classes like [^a-z] and [[:space:]] to match newline
	DotNL					// allow . to match newline
	OneLine					// treat ^ and $ as only matching at beginning and end of text
	NonGreedy				// make repetition operators default to non-greedy
	PerlX					// allow Perl extensions
	UnicodeGroups				// allow \p{Han}, \P{Han} for Unicode group and negation
	WasDollar				// regexp OpEndText was $, not \z
	Simple					// regexp contains no counted repetition

	MatchNL	= ClassNL | DotNL

	Perl		= ClassNL | OneLine | PerlX | UnicodeGroups	// as close to Perl as possible
	POSIX	Flags	= 0						// POSIX syntax
)

// Pseudo-ops for parsing stack.
const (
	opLeftParen	= opPseudo + iota
	opVerticalBar
)

// maxHeight is the maximum height of a regexp parse tree.
//line /usr/local/go/src/regexp/syntax/parse.go:81
// It is somewhat arbitrarily chosen, but the idea is to be large enough
//line /usr/local/go/src/regexp/syntax/parse.go:81
// that no one will actually hit in real use but at the same time small enough
//line /usr/local/go/src/regexp/syntax/parse.go:81
// that recursion on the Regexp tree will not hit the 1GB Go stack limit.
//line /usr/local/go/src/regexp/syntax/parse.go:81
// The maximum amount of stack for a single recursive frame is probably
//line /usr/local/go/src/regexp/syntax/parse.go:81
// closer to 1kB, so this could potentially be raised, but it seems unlikely
//line /usr/local/go/src/regexp/syntax/parse.go:81
// that people have regexps nested even this deeply.
//line /usr/local/go/src/regexp/syntax/parse.go:81
// We ran a test on Google's C++ code base and turned up only
//line /usr/local/go/src/regexp/syntax/parse.go:81
// a single use case with depth > 100; it had depth 128.
//line /usr/local/go/src/regexp/syntax/parse.go:81
// Using depth 1000 should be plenty of margin.
//line /usr/local/go/src/regexp/syntax/parse.go:81
// As an optimization, we don't even bother calculating heights
//line /usr/local/go/src/regexp/syntax/parse.go:81
// until we've allocated at least maxHeight Regexp structures.
//line /usr/local/go/src/regexp/syntax/parse.go:93
const maxHeight = 1000

// maxSize is the maximum size of a compiled regexp in Insts.
//line /usr/local/go/src/regexp/syntax/parse.go:95
// It too is somewhat arbitrarily chosen, but the idea is to be large enough
//line /usr/local/go/src/regexp/syntax/parse.go:95
// to allow significant regexps while at the same time small enough that
//line /usr/local/go/src/regexp/syntax/parse.go:95
// the compiled form will not take up too much memory.
//line /usr/local/go/src/regexp/syntax/parse.go:95
// 128 MB is enough for a 3.3 million Inst structures, which roughly
//line /usr/local/go/src/regexp/syntax/parse.go:95
// corresponds to a 3.3 MB regexp.
//line /usr/local/go/src/regexp/syntax/parse.go:101
const (
	maxSize		= 128 << 20 / instSize
	instSize	= 5 * 8	// byte, 2 uint32, slice is 5 64-bit words
)

// maxRunes is the maximum number of runes allowed in a regexp tree
//line /usr/local/go/src/regexp/syntax/parse.go:106
// counting the runes in all the nodes.
//line /usr/local/go/src/regexp/syntax/parse.go:106
// Ignoring character classes p.numRunes is always less than the length of the regexp.
//line /usr/local/go/src/regexp/syntax/parse.go:106
// Character classes can make it much larger: each \pL adds 1292 runes.
//line /usr/local/go/src/regexp/syntax/parse.go:106
// 128 MB is enough for 32M runes, which is over 26k \pL instances.
//line /usr/local/go/src/regexp/syntax/parse.go:106
// Note that repetitions do not make copies of the rune slices,
//line /usr/local/go/src/regexp/syntax/parse.go:106
// so \pL{1000} is only one rune slice, not 1000.
//line /usr/local/go/src/regexp/syntax/parse.go:106
// We could keep a cache of character classes we've seen,
//line /usr/local/go/src/regexp/syntax/parse.go:106
// so that all the \pL we see use the same rune list,
//line /usr/local/go/src/regexp/syntax/parse.go:106
// but that doesn't remove the problem entirely:
//line /usr/local/go/src/regexp/syntax/parse.go:106
// consider something like [\pL01234][\pL01235][\pL01236]...[\pL^&*()].
//line /usr/local/go/src/regexp/syntax/parse.go:106
// And because the Rune slice is exposed directly in the Regexp,
//line /usr/local/go/src/regexp/syntax/parse.go:106
// there is not an opportunity to change the representation to allow
//line /usr/local/go/src/regexp/syntax/parse.go:106
// partial sharing between different character classes.
//line /usr/local/go/src/regexp/syntax/parse.go:106
// So the limit is the best we can do.
//line /usr/local/go/src/regexp/syntax/parse.go:121
const (
	maxRunes	= 128 << 20 / runeSize
	runeSize	= 4	// rune is int32
)

type parser struct {
	flags		Flags		// parse mode flags
	stack		[]*Regexp	// stack of parsed expressions
	free		*Regexp
	numCap		int	// number of capturing groups seen
	wholeRegexp	string
	tmpClass	[]rune			// temporary char class work space
	numRegexp	int			// number of regexps allocated
	numRunes	int			// number of runes in char classes
	repeats		int64			// product of all repetitions seen
	height		map[*Regexp]int		// regexp height, for height limit check
	size		map[*Regexp]int64	// regexp compiled size, for size limit check
}

func (p *parser) newRegexp(op Op) *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:140
	_go_fuzz_dep_.CoverTab[63072]++
							re := p.free
							if re != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:142
		_go_fuzz_dep_.CoverTab[63074]++
								p.free = re.Sub0[0]
								*re = Regexp{}
//line /usr/local/go/src/regexp/syntax/parse.go:144
		// _ = "end of CoverTab[63074]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:145
		_go_fuzz_dep_.CoverTab[63075]++
								re = new(Regexp)
								p.numRegexp++
//line /usr/local/go/src/regexp/syntax/parse.go:147
		// _ = "end of CoverTab[63075]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:148
	// _ = "end of CoverTab[63072]"
//line /usr/local/go/src/regexp/syntax/parse.go:148
	_go_fuzz_dep_.CoverTab[63073]++
							re.Op = op
							return re
//line /usr/local/go/src/regexp/syntax/parse.go:150
	// _ = "end of CoverTab[63073]"
}

func (p *parser) reuse(re *Regexp) {
//line /usr/local/go/src/regexp/syntax/parse.go:153
	_go_fuzz_dep_.CoverTab[63076]++
							if p.height != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:154
		_go_fuzz_dep_.CoverTab[63078]++
								delete(p.height, re)
//line /usr/local/go/src/regexp/syntax/parse.go:155
		// _ = "end of CoverTab[63078]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:156
		_go_fuzz_dep_.CoverTab[63079]++
//line /usr/local/go/src/regexp/syntax/parse.go:156
		// _ = "end of CoverTab[63079]"
//line /usr/local/go/src/regexp/syntax/parse.go:156
	}
//line /usr/local/go/src/regexp/syntax/parse.go:156
	// _ = "end of CoverTab[63076]"
//line /usr/local/go/src/regexp/syntax/parse.go:156
	_go_fuzz_dep_.CoverTab[63077]++
							re.Sub0[0] = p.free
							p.free = re
//line /usr/local/go/src/regexp/syntax/parse.go:158
	// _ = "end of CoverTab[63077]"
}

func (p *parser) checkLimits(re *Regexp) {
//line /usr/local/go/src/regexp/syntax/parse.go:161
	_go_fuzz_dep_.CoverTab[63080]++
							if p.numRunes > maxRunes {
//line /usr/local/go/src/regexp/syntax/parse.go:162
		_go_fuzz_dep_.CoverTab[63082]++
								panic(ErrLarge)
//line /usr/local/go/src/regexp/syntax/parse.go:163
		// _ = "end of CoverTab[63082]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:164
		_go_fuzz_dep_.CoverTab[63083]++
//line /usr/local/go/src/regexp/syntax/parse.go:164
		// _ = "end of CoverTab[63083]"
//line /usr/local/go/src/regexp/syntax/parse.go:164
	}
//line /usr/local/go/src/regexp/syntax/parse.go:164
	// _ = "end of CoverTab[63080]"
//line /usr/local/go/src/regexp/syntax/parse.go:164
	_go_fuzz_dep_.CoverTab[63081]++
							p.checkSize(re)
							p.checkHeight(re)
//line /usr/local/go/src/regexp/syntax/parse.go:166
	// _ = "end of CoverTab[63081]"
}

func (p *parser) checkSize(re *Regexp) {
//line /usr/local/go/src/regexp/syntax/parse.go:169
	_go_fuzz_dep_.CoverTab[63084]++
							if p.size == nil {
//line /usr/local/go/src/regexp/syntax/parse.go:170
		_go_fuzz_dep_.CoverTab[63086]++

//line /usr/local/go/src/regexp/syntax/parse.go:176
		if p.repeats == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:176
			_go_fuzz_dep_.CoverTab[63090]++
									p.repeats = 1
//line /usr/local/go/src/regexp/syntax/parse.go:177
			// _ = "end of CoverTab[63090]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:178
			_go_fuzz_dep_.CoverTab[63091]++
//line /usr/local/go/src/regexp/syntax/parse.go:178
			// _ = "end of CoverTab[63091]"
//line /usr/local/go/src/regexp/syntax/parse.go:178
		}
//line /usr/local/go/src/regexp/syntax/parse.go:178
		// _ = "end of CoverTab[63086]"
//line /usr/local/go/src/regexp/syntax/parse.go:178
		_go_fuzz_dep_.CoverTab[63087]++
								if re.Op == OpRepeat {
//line /usr/local/go/src/regexp/syntax/parse.go:179
			_go_fuzz_dep_.CoverTab[63092]++
									n := re.Max
									if n == -1 {
//line /usr/local/go/src/regexp/syntax/parse.go:181
				_go_fuzz_dep_.CoverTab[63095]++
										n = re.Min
//line /usr/local/go/src/regexp/syntax/parse.go:182
				// _ = "end of CoverTab[63095]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:183
				_go_fuzz_dep_.CoverTab[63096]++
//line /usr/local/go/src/regexp/syntax/parse.go:183
				// _ = "end of CoverTab[63096]"
//line /usr/local/go/src/regexp/syntax/parse.go:183
			}
//line /usr/local/go/src/regexp/syntax/parse.go:183
			// _ = "end of CoverTab[63092]"
//line /usr/local/go/src/regexp/syntax/parse.go:183
			_go_fuzz_dep_.CoverTab[63093]++
									if n <= 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:184
				_go_fuzz_dep_.CoverTab[63097]++
										n = 1
//line /usr/local/go/src/regexp/syntax/parse.go:185
				// _ = "end of CoverTab[63097]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:186
				_go_fuzz_dep_.CoverTab[63098]++
//line /usr/local/go/src/regexp/syntax/parse.go:186
				// _ = "end of CoverTab[63098]"
//line /usr/local/go/src/regexp/syntax/parse.go:186
			}
//line /usr/local/go/src/regexp/syntax/parse.go:186
			// _ = "end of CoverTab[63093]"
//line /usr/local/go/src/regexp/syntax/parse.go:186
			_go_fuzz_dep_.CoverTab[63094]++
									if int64(n) > maxSize/p.repeats {
//line /usr/local/go/src/regexp/syntax/parse.go:187
				_go_fuzz_dep_.CoverTab[63099]++
										p.repeats = maxSize
//line /usr/local/go/src/regexp/syntax/parse.go:188
				// _ = "end of CoverTab[63099]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:189
				_go_fuzz_dep_.CoverTab[63100]++
										p.repeats *= int64(n)
//line /usr/local/go/src/regexp/syntax/parse.go:190
				// _ = "end of CoverTab[63100]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:191
			// _ = "end of CoverTab[63094]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:192
			_go_fuzz_dep_.CoverTab[63101]++
//line /usr/local/go/src/regexp/syntax/parse.go:192
			// _ = "end of CoverTab[63101]"
//line /usr/local/go/src/regexp/syntax/parse.go:192
		}
//line /usr/local/go/src/regexp/syntax/parse.go:192
		// _ = "end of CoverTab[63087]"
//line /usr/local/go/src/regexp/syntax/parse.go:192
		_go_fuzz_dep_.CoverTab[63088]++
								if int64(p.numRegexp) < maxSize/p.repeats {
//line /usr/local/go/src/regexp/syntax/parse.go:193
			_go_fuzz_dep_.CoverTab[63102]++
									return
//line /usr/local/go/src/regexp/syntax/parse.go:194
			// _ = "end of CoverTab[63102]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:195
			_go_fuzz_dep_.CoverTab[63103]++
//line /usr/local/go/src/regexp/syntax/parse.go:195
			// _ = "end of CoverTab[63103]"
//line /usr/local/go/src/regexp/syntax/parse.go:195
		}
//line /usr/local/go/src/regexp/syntax/parse.go:195
		// _ = "end of CoverTab[63088]"
//line /usr/local/go/src/regexp/syntax/parse.go:195
		_go_fuzz_dep_.CoverTab[63089]++

//line /usr/local/go/src/regexp/syntax/parse.go:200
		p.size = make(map[*Regexp]int64)
		for _, re := range p.stack {
//line /usr/local/go/src/regexp/syntax/parse.go:201
			_go_fuzz_dep_.CoverTab[63104]++
									p.checkSize(re)
//line /usr/local/go/src/regexp/syntax/parse.go:202
			// _ = "end of CoverTab[63104]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:203
		// _ = "end of CoverTab[63089]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:204
		_go_fuzz_dep_.CoverTab[63105]++
//line /usr/local/go/src/regexp/syntax/parse.go:204
		// _ = "end of CoverTab[63105]"
//line /usr/local/go/src/regexp/syntax/parse.go:204
	}
//line /usr/local/go/src/regexp/syntax/parse.go:204
	// _ = "end of CoverTab[63084]"
//line /usr/local/go/src/regexp/syntax/parse.go:204
	_go_fuzz_dep_.CoverTab[63085]++

							if p.calcSize(re, true) > maxSize {
//line /usr/local/go/src/regexp/syntax/parse.go:206
		_go_fuzz_dep_.CoverTab[63106]++
								panic(ErrLarge)
//line /usr/local/go/src/regexp/syntax/parse.go:207
		// _ = "end of CoverTab[63106]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:208
		_go_fuzz_dep_.CoverTab[63107]++
//line /usr/local/go/src/regexp/syntax/parse.go:208
		// _ = "end of CoverTab[63107]"
//line /usr/local/go/src/regexp/syntax/parse.go:208
	}
//line /usr/local/go/src/regexp/syntax/parse.go:208
	// _ = "end of CoverTab[63085]"
}

func (p *parser) calcSize(re *Regexp, force bool) int64 {
//line /usr/local/go/src/regexp/syntax/parse.go:211
	_go_fuzz_dep_.CoverTab[63108]++
							if !force {
//line /usr/local/go/src/regexp/syntax/parse.go:212
		_go_fuzz_dep_.CoverTab[63112]++
								if size, ok := p.size[re]; ok {
//line /usr/local/go/src/regexp/syntax/parse.go:213
			_go_fuzz_dep_.CoverTab[63113]++
									return size
//line /usr/local/go/src/regexp/syntax/parse.go:214
			// _ = "end of CoverTab[63113]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:215
			_go_fuzz_dep_.CoverTab[63114]++
//line /usr/local/go/src/regexp/syntax/parse.go:215
			// _ = "end of CoverTab[63114]"
//line /usr/local/go/src/regexp/syntax/parse.go:215
		}
//line /usr/local/go/src/regexp/syntax/parse.go:215
		// _ = "end of CoverTab[63112]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:216
		_go_fuzz_dep_.CoverTab[63115]++
//line /usr/local/go/src/regexp/syntax/parse.go:216
		// _ = "end of CoverTab[63115]"
//line /usr/local/go/src/regexp/syntax/parse.go:216
	}
//line /usr/local/go/src/regexp/syntax/parse.go:216
	// _ = "end of CoverTab[63108]"
//line /usr/local/go/src/regexp/syntax/parse.go:216
	_go_fuzz_dep_.CoverTab[63109]++

							var size int64
							switch re.Op {
	case OpLiteral:
//line /usr/local/go/src/regexp/syntax/parse.go:220
		_go_fuzz_dep_.CoverTab[63116]++
								size = int64(len(re.Rune))
//line /usr/local/go/src/regexp/syntax/parse.go:221
		// _ = "end of CoverTab[63116]"
	case OpCapture, OpStar:
//line /usr/local/go/src/regexp/syntax/parse.go:222
		_go_fuzz_dep_.CoverTab[63117]++

								size = 2 + p.calcSize(re.Sub[0], false)
//line /usr/local/go/src/regexp/syntax/parse.go:224
		// _ = "end of CoverTab[63117]"
	case OpPlus, OpQuest:
//line /usr/local/go/src/regexp/syntax/parse.go:225
		_go_fuzz_dep_.CoverTab[63118]++
								size = 1 + p.calcSize(re.Sub[0], false)
//line /usr/local/go/src/regexp/syntax/parse.go:226
		// _ = "end of CoverTab[63118]"
	case OpConcat:
//line /usr/local/go/src/regexp/syntax/parse.go:227
		_go_fuzz_dep_.CoverTab[63119]++
								for _, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/parse.go:228
			_go_fuzz_dep_.CoverTab[63125]++
									size += p.calcSize(sub, false)
//line /usr/local/go/src/regexp/syntax/parse.go:229
			// _ = "end of CoverTab[63125]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:230
		// _ = "end of CoverTab[63119]"
	case OpAlternate:
//line /usr/local/go/src/regexp/syntax/parse.go:231
		_go_fuzz_dep_.CoverTab[63120]++
								for _, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/parse.go:232
			_go_fuzz_dep_.CoverTab[63126]++
									size += p.calcSize(sub, false)
//line /usr/local/go/src/regexp/syntax/parse.go:233
			// _ = "end of CoverTab[63126]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:234
		// _ = "end of CoverTab[63120]"
//line /usr/local/go/src/regexp/syntax/parse.go:234
		_go_fuzz_dep_.CoverTab[63121]++
								if len(re.Sub) > 1 {
//line /usr/local/go/src/regexp/syntax/parse.go:235
			_go_fuzz_dep_.CoverTab[63127]++
									size += int64(len(re.Sub)) - 1
//line /usr/local/go/src/regexp/syntax/parse.go:236
			// _ = "end of CoverTab[63127]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:237
			_go_fuzz_dep_.CoverTab[63128]++
//line /usr/local/go/src/regexp/syntax/parse.go:237
			// _ = "end of CoverTab[63128]"
//line /usr/local/go/src/regexp/syntax/parse.go:237
		}
//line /usr/local/go/src/regexp/syntax/parse.go:237
		// _ = "end of CoverTab[63121]"
	case OpRepeat:
//line /usr/local/go/src/regexp/syntax/parse.go:238
		_go_fuzz_dep_.CoverTab[63122]++
								sub := p.calcSize(re.Sub[0], false)
								if re.Max == -1 {
//line /usr/local/go/src/regexp/syntax/parse.go:240
			_go_fuzz_dep_.CoverTab[63129]++
									if re.Min == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:241
				_go_fuzz_dep_.CoverTab[63131]++
										size = 2 + sub
//line /usr/local/go/src/regexp/syntax/parse.go:242
				// _ = "end of CoverTab[63131]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:243
				_go_fuzz_dep_.CoverTab[63132]++
										size = 1 + int64(re.Min)*sub
//line /usr/local/go/src/regexp/syntax/parse.go:244
				// _ = "end of CoverTab[63132]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:245
			// _ = "end of CoverTab[63129]"
//line /usr/local/go/src/regexp/syntax/parse.go:245
			_go_fuzz_dep_.CoverTab[63130]++
									break
//line /usr/local/go/src/regexp/syntax/parse.go:246
			// _ = "end of CoverTab[63130]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:247
			_go_fuzz_dep_.CoverTab[63133]++
//line /usr/local/go/src/regexp/syntax/parse.go:247
			// _ = "end of CoverTab[63133]"
//line /usr/local/go/src/regexp/syntax/parse.go:247
		}
//line /usr/local/go/src/regexp/syntax/parse.go:247
		// _ = "end of CoverTab[63122]"
//line /usr/local/go/src/regexp/syntax/parse.go:247
		_go_fuzz_dep_.CoverTab[63123]++

								size = int64(re.Max)*sub + int64(re.Max-re.Min)
//line /usr/local/go/src/regexp/syntax/parse.go:249
		// _ = "end of CoverTab[63123]"
//line /usr/local/go/src/regexp/syntax/parse.go:249
	default:
//line /usr/local/go/src/regexp/syntax/parse.go:249
		_go_fuzz_dep_.CoverTab[63124]++
//line /usr/local/go/src/regexp/syntax/parse.go:249
		// _ = "end of CoverTab[63124]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:250
	// _ = "end of CoverTab[63109]"
//line /usr/local/go/src/regexp/syntax/parse.go:250
	_go_fuzz_dep_.CoverTab[63110]++

							if size < 1 {
//line /usr/local/go/src/regexp/syntax/parse.go:252
		_go_fuzz_dep_.CoverTab[63134]++
								size = 1
//line /usr/local/go/src/regexp/syntax/parse.go:253
		// _ = "end of CoverTab[63134]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:254
		_go_fuzz_dep_.CoverTab[63135]++
//line /usr/local/go/src/regexp/syntax/parse.go:254
		// _ = "end of CoverTab[63135]"
//line /usr/local/go/src/regexp/syntax/parse.go:254
	}
//line /usr/local/go/src/regexp/syntax/parse.go:254
	// _ = "end of CoverTab[63110]"
//line /usr/local/go/src/regexp/syntax/parse.go:254
	_go_fuzz_dep_.CoverTab[63111]++
							p.size[re] = size
							return size
//line /usr/local/go/src/regexp/syntax/parse.go:256
	// _ = "end of CoverTab[63111]"
}

func (p *parser) checkHeight(re *Regexp) {
//line /usr/local/go/src/regexp/syntax/parse.go:259
	_go_fuzz_dep_.CoverTab[63136]++
							if p.numRegexp < maxHeight {
//line /usr/local/go/src/regexp/syntax/parse.go:260
		_go_fuzz_dep_.CoverTab[63139]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:261
		// _ = "end of CoverTab[63139]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:262
		_go_fuzz_dep_.CoverTab[63140]++
//line /usr/local/go/src/regexp/syntax/parse.go:262
		// _ = "end of CoverTab[63140]"
//line /usr/local/go/src/regexp/syntax/parse.go:262
	}
//line /usr/local/go/src/regexp/syntax/parse.go:262
	// _ = "end of CoverTab[63136]"
//line /usr/local/go/src/regexp/syntax/parse.go:262
	_go_fuzz_dep_.CoverTab[63137]++
							if p.height == nil {
//line /usr/local/go/src/regexp/syntax/parse.go:263
		_go_fuzz_dep_.CoverTab[63141]++
								p.height = make(map[*Regexp]int)
								for _, re := range p.stack {
//line /usr/local/go/src/regexp/syntax/parse.go:265
			_go_fuzz_dep_.CoverTab[63142]++
									p.checkHeight(re)
//line /usr/local/go/src/regexp/syntax/parse.go:266
			// _ = "end of CoverTab[63142]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:267
		// _ = "end of CoverTab[63141]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:268
		_go_fuzz_dep_.CoverTab[63143]++
//line /usr/local/go/src/regexp/syntax/parse.go:268
		// _ = "end of CoverTab[63143]"
//line /usr/local/go/src/regexp/syntax/parse.go:268
	}
//line /usr/local/go/src/regexp/syntax/parse.go:268
	// _ = "end of CoverTab[63137]"
//line /usr/local/go/src/regexp/syntax/parse.go:268
	_go_fuzz_dep_.CoverTab[63138]++
							if p.calcHeight(re, true) > maxHeight {
//line /usr/local/go/src/regexp/syntax/parse.go:269
		_go_fuzz_dep_.CoverTab[63144]++
								panic(ErrNestingDepth)
//line /usr/local/go/src/regexp/syntax/parse.go:270
		// _ = "end of CoverTab[63144]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:271
		_go_fuzz_dep_.CoverTab[63145]++
//line /usr/local/go/src/regexp/syntax/parse.go:271
		// _ = "end of CoverTab[63145]"
//line /usr/local/go/src/regexp/syntax/parse.go:271
	}
//line /usr/local/go/src/regexp/syntax/parse.go:271
	// _ = "end of CoverTab[63138]"
}

func (p *parser) calcHeight(re *Regexp, force bool) int {
//line /usr/local/go/src/regexp/syntax/parse.go:274
	_go_fuzz_dep_.CoverTab[63146]++
							if !force {
//line /usr/local/go/src/regexp/syntax/parse.go:275
		_go_fuzz_dep_.CoverTab[63149]++
								if h, ok := p.height[re]; ok {
//line /usr/local/go/src/regexp/syntax/parse.go:276
			_go_fuzz_dep_.CoverTab[63150]++
									return h
//line /usr/local/go/src/regexp/syntax/parse.go:277
			// _ = "end of CoverTab[63150]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:278
			_go_fuzz_dep_.CoverTab[63151]++
//line /usr/local/go/src/regexp/syntax/parse.go:278
			// _ = "end of CoverTab[63151]"
//line /usr/local/go/src/regexp/syntax/parse.go:278
		}
//line /usr/local/go/src/regexp/syntax/parse.go:278
		// _ = "end of CoverTab[63149]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:279
		_go_fuzz_dep_.CoverTab[63152]++
//line /usr/local/go/src/regexp/syntax/parse.go:279
		// _ = "end of CoverTab[63152]"
//line /usr/local/go/src/regexp/syntax/parse.go:279
	}
//line /usr/local/go/src/regexp/syntax/parse.go:279
	// _ = "end of CoverTab[63146]"
//line /usr/local/go/src/regexp/syntax/parse.go:279
	_go_fuzz_dep_.CoverTab[63147]++
							h := 1
							for _, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/parse.go:281
		_go_fuzz_dep_.CoverTab[63153]++
								hsub := p.calcHeight(sub, false)
								if h < 1+hsub {
//line /usr/local/go/src/regexp/syntax/parse.go:283
			_go_fuzz_dep_.CoverTab[63154]++
									h = 1 + hsub
//line /usr/local/go/src/regexp/syntax/parse.go:284
			// _ = "end of CoverTab[63154]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:285
			_go_fuzz_dep_.CoverTab[63155]++
//line /usr/local/go/src/regexp/syntax/parse.go:285
			// _ = "end of CoverTab[63155]"
//line /usr/local/go/src/regexp/syntax/parse.go:285
		}
//line /usr/local/go/src/regexp/syntax/parse.go:285
		// _ = "end of CoverTab[63153]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:286
	// _ = "end of CoverTab[63147]"
//line /usr/local/go/src/regexp/syntax/parse.go:286
	_go_fuzz_dep_.CoverTab[63148]++
							p.height[re] = h
							return h
//line /usr/local/go/src/regexp/syntax/parse.go:288
	// _ = "end of CoverTab[63148]"
}

//line /usr/local/go/src/regexp/syntax/parse.go:293
// push pushes the regexp re onto the parse stack and returns the regexp.
func (p *parser) push(re *Regexp) *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:294
	_go_fuzz_dep_.CoverTab[63156]++
							p.numRunes += len(re.Rune)
							if re.Op == OpCharClass && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:296
		_go_fuzz_dep_.CoverTab[63158]++
//line /usr/local/go/src/regexp/syntax/parse.go:296
		return len(re.Rune) == 2
//line /usr/local/go/src/regexp/syntax/parse.go:296
		// _ = "end of CoverTab[63158]"
//line /usr/local/go/src/regexp/syntax/parse.go:296
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:296
		_go_fuzz_dep_.CoverTab[63159]++
//line /usr/local/go/src/regexp/syntax/parse.go:296
		return re.Rune[0] == re.Rune[1]
//line /usr/local/go/src/regexp/syntax/parse.go:296
		// _ = "end of CoverTab[63159]"
//line /usr/local/go/src/regexp/syntax/parse.go:296
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:296
		_go_fuzz_dep_.CoverTab[63160]++

								if p.maybeConcat(re.Rune[0], p.flags&^FoldCase) {
//line /usr/local/go/src/regexp/syntax/parse.go:298
			_go_fuzz_dep_.CoverTab[63162]++
									return nil
//line /usr/local/go/src/regexp/syntax/parse.go:299
			// _ = "end of CoverTab[63162]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:300
			_go_fuzz_dep_.CoverTab[63163]++
//line /usr/local/go/src/regexp/syntax/parse.go:300
			// _ = "end of CoverTab[63163]"
//line /usr/local/go/src/regexp/syntax/parse.go:300
		}
//line /usr/local/go/src/regexp/syntax/parse.go:300
		// _ = "end of CoverTab[63160]"
//line /usr/local/go/src/regexp/syntax/parse.go:300
		_go_fuzz_dep_.CoverTab[63161]++
								re.Op = OpLiteral
								re.Rune = re.Rune[:1]
								re.Flags = p.flags &^ FoldCase
//line /usr/local/go/src/regexp/syntax/parse.go:303
		// _ = "end of CoverTab[63161]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:304
		_go_fuzz_dep_.CoverTab[63164]++
//line /usr/local/go/src/regexp/syntax/parse.go:304
		if re.Op == OpCharClass && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:304
			_go_fuzz_dep_.CoverTab[63165]++
//line /usr/local/go/src/regexp/syntax/parse.go:304
			return len(re.Rune) == 4
//line /usr/local/go/src/regexp/syntax/parse.go:304
			// _ = "end of CoverTab[63165]"
//line /usr/local/go/src/regexp/syntax/parse.go:304
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:304
			_go_fuzz_dep_.CoverTab[63166]++
//line /usr/local/go/src/regexp/syntax/parse.go:304
			return re.Rune[0] == re.Rune[1]
									// _ = "end of CoverTab[63166]"
//line /usr/local/go/src/regexp/syntax/parse.go:305
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:305
			_go_fuzz_dep_.CoverTab[63167]++
//line /usr/local/go/src/regexp/syntax/parse.go:305
			return re.Rune[2] == re.Rune[3]
//line /usr/local/go/src/regexp/syntax/parse.go:305
			// _ = "end of CoverTab[63167]"
//line /usr/local/go/src/regexp/syntax/parse.go:305
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:305
			_go_fuzz_dep_.CoverTab[63168]++
//line /usr/local/go/src/regexp/syntax/parse.go:305
			return unicode.SimpleFold(re.Rune[0]) == re.Rune[2]
									// _ = "end of CoverTab[63168]"
//line /usr/local/go/src/regexp/syntax/parse.go:306
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:306
			_go_fuzz_dep_.CoverTab[63169]++
//line /usr/local/go/src/regexp/syntax/parse.go:306
			return unicode.SimpleFold(re.Rune[2]) == re.Rune[0]
									// _ = "end of CoverTab[63169]"
//line /usr/local/go/src/regexp/syntax/parse.go:307
		}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:307
			_go_fuzz_dep_.CoverTab[63170]++
//line /usr/local/go/src/regexp/syntax/parse.go:307
			return re.Op == OpCharClass && func() bool {
										_go_fuzz_dep_.CoverTab[63171]++
//line /usr/local/go/src/regexp/syntax/parse.go:308
				return len(re.Rune) == 2
//line /usr/local/go/src/regexp/syntax/parse.go:308
				// _ = "end of CoverTab[63171]"
//line /usr/local/go/src/regexp/syntax/parse.go:308
			}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:308
				_go_fuzz_dep_.CoverTab[63172]++
//line /usr/local/go/src/regexp/syntax/parse.go:308
				return re.Rune[0]+1 == re.Rune[1]
										// _ = "end of CoverTab[63172]"
//line /usr/local/go/src/regexp/syntax/parse.go:309
			}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:309
				_go_fuzz_dep_.CoverTab[63173]++
//line /usr/local/go/src/regexp/syntax/parse.go:309
				return unicode.SimpleFold(re.Rune[0]) == re.Rune[1]
										// _ = "end of CoverTab[63173]"
//line /usr/local/go/src/regexp/syntax/parse.go:310
			}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:310
				_go_fuzz_dep_.CoverTab[63174]++
//line /usr/local/go/src/regexp/syntax/parse.go:310
				return unicode.SimpleFold(re.Rune[1]) == re.Rune[0]
										// _ = "end of CoverTab[63174]"
//line /usr/local/go/src/regexp/syntax/parse.go:311
			}()
//line /usr/local/go/src/regexp/syntax/parse.go:311
			// _ = "end of CoverTab[63170]"
//line /usr/local/go/src/regexp/syntax/parse.go:311
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:311
			_go_fuzz_dep_.CoverTab[63175]++

									if p.maybeConcat(re.Rune[0], p.flags|FoldCase) {
//line /usr/local/go/src/regexp/syntax/parse.go:313
				_go_fuzz_dep_.CoverTab[63177]++
										return nil
//line /usr/local/go/src/regexp/syntax/parse.go:314
				// _ = "end of CoverTab[63177]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:315
				_go_fuzz_dep_.CoverTab[63178]++
//line /usr/local/go/src/regexp/syntax/parse.go:315
				// _ = "end of CoverTab[63178]"
//line /usr/local/go/src/regexp/syntax/parse.go:315
			}
//line /usr/local/go/src/regexp/syntax/parse.go:315
			// _ = "end of CoverTab[63175]"
//line /usr/local/go/src/regexp/syntax/parse.go:315
			_go_fuzz_dep_.CoverTab[63176]++

//line /usr/local/go/src/regexp/syntax/parse.go:318
			re.Op = OpLiteral
									re.Rune = re.Rune[:1]
									re.Flags = p.flags | FoldCase
//line /usr/local/go/src/regexp/syntax/parse.go:320
			// _ = "end of CoverTab[63176]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:321
			_go_fuzz_dep_.CoverTab[63179]++

									p.maybeConcat(-1, 0)
//line /usr/local/go/src/regexp/syntax/parse.go:323
			// _ = "end of CoverTab[63179]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:324
		// _ = "end of CoverTab[63164]"
//line /usr/local/go/src/regexp/syntax/parse.go:324
	}
//line /usr/local/go/src/regexp/syntax/parse.go:324
	// _ = "end of CoverTab[63156]"
//line /usr/local/go/src/regexp/syntax/parse.go:324
	_go_fuzz_dep_.CoverTab[63157]++

							p.stack = append(p.stack, re)
							p.checkLimits(re)
							return re
//line /usr/local/go/src/regexp/syntax/parse.go:328
	// _ = "end of CoverTab[63157]"
}

// maybeConcat implements incremental concatenation
//line /usr/local/go/src/regexp/syntax/parse.go:331
// of literal runes into string nodes. The parser calls this
//line /usr/local/go/src/regexp/syntax/parse.go:331
// before each push, so only the top fragment of the stack
//line /usr/local/go/src/regexp/syntax/parse.go:331
// might need processing. Since this is called before a push,
//line /usr/local/go/src/regexp/syntax/parse.go:331
// the topmost literal is no longer subject to operators like *
//line /usr/local/go/src/regexp/syntax/parse.go:331
// (Otherwise ab* would turn into (ab)*.)
//line /usr/local/go/src/regexp/syntax/parse.go:331
// If r >= 0 and there's a node left over, maybeConcat uses it
//line /usr/local/go/src/regexp/syntax/parse.go:331
// to push r with the given flags.
//line /usr/local/go/src/regexp/syntax/parse.go:331
// maybeConcat reports whether r was pushed.
//line /usr/local/go/src/regexp/syntax/parse.go:340
func (p *parser) maybeConcat(r rune, flags Flags) bool {
//line /usr/local/go/src/regexp/syntax/parse.go:340
	_go_fuzz_dep_.CoverTab[63180]++
							n := len(p.stack)
							if n < 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:342
		_go_fuzz_dep_.CoverTab[63184]++
								return false
//line /usr/local/go/src/regexp/syntax/parse.go:343
		// _ = "end of CoverTab[63184]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:344
		_go_fuzz_dep_.CoverTab[63185]++
//line /usr/local/go/src/regexp/syntax/parse.go:344
		// _ = "end of CoverTab[63185]"
//line /usr/local/go/src/regexp/syntax/parse.go:344
	}
//line /usr/local/go/src/regexp/syntax/parse.go:344
	// _ = "end of CoverTab[63180]"
//line /usr/local/go/src/regexp/syntax/parse.go:344
	_go_fuzz_dep_.CoverTab[63181]++

							re1 := p.stack[n-1]
							re2 := p.stack[n-2]
							if re1.Op != OpLiteral || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:348
		_go_fuzz_dep_.CoverTab[63186]++
//line /usr/local/go/src/regexp/syntax/parse.go:348
		return re2.Op != OpLiteral
//line /usr/local/go/src/regexp/syntax/parse.go:348
		// _ = "end of CoverTab[63186]"
//line /usr/local/go/src/regexp/syntax/parse.go:348
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:348
		_go_fuzz_dep_.CoverTab[63187]++
//line /usr/local/go/src/regexp/syntax/parse.go:348
		return re1.Flags&FoldCase != re2.Flags&FoldCase
//line /usr/local/go/src/regexp/syntax/parse.go:348
		// _ = "end of CoverTab[63187]"
//line /usr/local/go/src/regexp/syntax/parse.go:348
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:348
		_go_fuzz_dep_.CoverTab[63188]++
								return false
//line /usr/local/go/src/regexp/syntax/parse.go:349
		// _ = "end of CoverTab[63188]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:350
		_go_fuzz_dep_.CoverTab[63189]++
//line /usr/local/go/src/regexp/syntax/parse.go:350
		// _ = "end of CoverTab[63189]"
//line /usr/local/go/src/regexp/syntax/parse.go:350
	}
//line /usr/local/go/src/regexp/syntax/parse.go:350
	// _ = "end of CoverTab[63181]"
//line /usr/local/go/src/regexp/syntax/parse.go:350
	_go_fuzz_dep_.CoverTab[63182]++

//line /usr/local/go/src/regexp/syntax/parse.go:353
	re2.Rune = append(re2.Rune, re1.Rune...)

//line /usr/local/go/src/regexp/syntax/parse.go:356
	if r >= 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:356
		_go_fuzz_dep_.CoverTab[63190]++
								re1.Rune = re1.Rune0[:1]
								re1.Rune[0] = r
								re1.Flags = flags
								return true
//line /usr/local/go/src/regexp/syntax/parse.go:360
		// _ = "end of CoverTab[63190]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:361
		_go_fuzz_dep_.CoverTab[63191]++
//line /usr/local/go/src/regexp/syntax/parse.go:361
		// _ = "end of CoverTab[63191]"
//line /usr/local/go/src/regexp/syntax/parse.go:361
	}
//line /usr/local/go/src/regexp/syntax/parse.go:361
	// _ = "end of CoverTab[63182]"
//line /usr/local/go/src/regexp/syntax/parse.go:361
	_go_fuzz_dep_.CoverTab[63183]++

							p.stack = p.stack[:n-1]
							p.reuse(re1)
							return false
//line /usr/local/go/src/regexp/syntax/parse.go:365
	// _ = "end of CoverTab[63183]"
}

// literal pushes a literal regexp for the rune r on the stack.
func (p *parser) literal(r rune) {
//line /usr/local/go/src/regexp/syntax/parse.go:369
	_go_fuzz_dep_.CoverTab[63192]++
							re := p.newRegexp(OpLiteral)
							re.Flags = p.flags
							if p.flags&FoldCase != 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:372
		_go_fuzz_dep_.CoverTab[63194]++
								r = minFoldRune(r)
//line /usr/local/go/src/regexp/syntax/parse.go:373
		// _ = "end of CoverTab[63194]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:374
		_go_fuzz_dep_.CoverTab[63195]++
//line /usr/local/go/src/regexp/syntax/parse.go:374
		// _ = "end of CoverTab[63195]"
//line /usr/local/go/src/regexp/syntax/parse.go:374
	}
//line /usr/local/go/src/regexp/syntax/parse.go:374
	// _ = "end of CoverTab[63192]"
//line /usr/local/go/src/regexp/syntax/parse.go:374
	_go_fuzz_dep_.CoverTab[63193]++
							re.Rune0[0] = r
							re.Rune = re.Rune0[:1]
							p.push(re)
//line /usr/local/go/src/regexp/syntax/parse.go:377
	// _ = "end of CoverTab[63193]"
}

// minFoldRune returns the minimum rune fold-equivalent to r.
func minFoldRune(r rune) rune {
//line /usr/local/go/src/regexp/syntax/parse.go:381
	_go_fuzz_dep_.CoverTab[63196]++
							if r < minFold || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:382
		_go_fuzz_dep_.CoverTab[63199]++
//line /usr/local/go/src/regexp/syntax/parse.go:382
		return r > maxFold
//line /usr/local/go/src/regexp/syntax/parse.go:382
		// _ = "end of CoverTab[63199]"
//line /usr/local/go/src/regexp/syntax/parse.go:382
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:382
		_go_fuzz_dep_.CoverTab[63200]++
								return r
//line /usr/local/go/src/regexp/syntax/parse.go:383
		// _ = "end of CoverTab[63200]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:384
		_go_fuzz_dep_.CoverTab[63201]++
//line /usr/local/go/src/regexp/syntax/parse.go:384
		// _ = "end of CoverTab[63201]"
//line /usr/local/go/src/regexp/syntax/parse.go:384
	}
//line /usr/local/go/src/regexp/syntax/parse.go:384
	// _ = "end of CoverTab[63196]"
//line /usr/local/go/src/regexp/syntax/parse.go:384
	_go_fuzz_dep_.CoverTab[63197]++
							min := r
							r0 := r
							for r = unicode.SimpleFold(r); r != r0; r = unicode.SimpleFold(r) {
//line /usr/local/go/src/regexp/syntax/parse.go:387
		_go_fuzz_dep_.CoverTab[63202]++
								if min > r {
//line /usr/local/go/src/regexp/syntax/parse.go:388
			_go_fuzz_dep_.CoverTab[63203]++
									min = r
//line /usr/local/go/src/regexp/syntax/parse.go:389
			// _ = "end of CoverTab[63203]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:390
			_go_fuzz_dep_.CoverTab[63204]++
//line /usr/local/go/src/regexp/syntax/parse.go:390
			// _ = "end of CoverTab[63204]"
//line /usr/local/go/src/regexp/syntax/parse.go:390
		}
//line /usr/local/go/src/regexp/syntax/parse.go:390
		// _ = "end of CoverTab[63202]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:391
	// _ = "end of CoverTab[63197]"
//line /usr/local/go/src/regexp/syntax/parse.go:391
	_go_fuzz_dep_.CoverTab[63198]++
							return min
//line /usr/local/go/src/regexp/syntax/parse.go:392
	// _ = "end of CoverTab[63198]"
}

// op pushes a regexp with the given op onto the stack
//line /usr/local/go/src/regexp/syntax/parse.go:395
// and returns that regexp.
//line /usr/local/go/src/regexp/syntax/parse.go:397
func (p *parser) op(op Op) *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:397
	_go_fuzz_dep_.CoverTab[63205]++
							re := p.newRegexp(op)
							re.Flags = p.flags
							return p.push(re)
//line /usr/local/go/src/regexp/syntax/parse.go:400
	// _ = "end of CoverTab[63205]"
}

// repeat replaces the top stack element with itself repeated according to op, min, max.
//line /usr/local/go/src/regexp/syntax/parse.go:403
// before is the regexp suffix starting at the repetition operator.
//line /usr/local/go/src/regexp/syntax/parse.go:403
// after is the regexp suffix following after the repetition operator.
//line /usr/local/go/src/regexp/syntax/parse.go:403
// repeat returns an updated 'after' and an error, if any.
//line /usr/local/go/src/regexp/syntax/parse.go:407
func (p *parser) repeat(op Op, min, max int, before, after, lastRepeat string) (string, error) {
//line /usr/local/go/src/regexp/syntax/parse.go:407
	_go_fuzz_dep_.CoverTab[63206]++
							flags := p.flags
							if p.flags&PerlX != 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:409
		_go_fuzz_dep_.CoverTab[63211]++
								if len(after) > 0 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:410
			_go_fuzz_dep_.CoverTab[63213]++
//line /usr/local/go/src/regexp/syntax/parse.go:410
			return after[0] == '?'
//line /usr/local/go/src/regexp/syntax/parse.go:410
			// _ = "end of CoverTab[63213]"
//line /usr/local/go/src/regexp/syntax/parse.go:410
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:410
			_go_fuzz_dep_.CoverTab[63214]++
									after = after[1:]
									flags ^= NonGreedy
//line /usr/local/go/src/regexp/syntax/parse.go:412
			// _ = "end of CoverTab[63214]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:413
			_go_fuzz_dep_.CoverTab[63215]++
//line /usr/local/go/src/regexp/syntax/parse.go:413
			// _ = "end of CoverTab[63215]"
//line /usr/local/go/src/regexp/syntax/parse.go:413
		}
//line /usr/local/go/src/regexp/syntax/parse.go:413
		// _ = "end of CoverTab[63211]"
//line /usr/local/go/src/regexp/syntax/parse.go:413
		_go_fuzz_dep_.CoverTab[63212]++
								if lastRepeat != "" {
//line /usr/local/go/src/regexp/syntax/parse.go:414
			_go_fuzz_dep_.CoverTab[63216]++

//line /usr/local/go/src/regexp/syntax/parse.go:418
			return "", &Error{ErrInvalidRepeatOp, lastRepeat[:len(lastRepeat)-len(after)]}
//line /usr/local/go/src/regexp/syntax/parse.go:418
			// _ = "end of CoverTab[63216]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:419
			_go_fuzz_dep_.CoverTab[63217]++
//line /usr/local/go/src/regexp/syntax/parse.go:419
			// _ = "end of CoverTab[63217]"
//line /usr/local/go/src/regexp/syntax/parse.go:419
		}
//line /usr/local/go/src/regexp/syntax/parse.go:419
		// _ = "end of CoverTab[63212]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:420
		_go_fuzz_dep_.CoverTab[63218]++
//line /usr/local/go/src/regexp/syntax/parse.go:420
		// _ = "end of CoverTab[63218]"
//line /usr/local/go/src/regexp/syntax/parse.go:420
	}
//line /usr/local/go/src/regexp/syntax/parse.go:420
	// _ = "end of CoverTab[63206]"
//line /usr/local/go/src/regexp/syntax/parse.go:420
	_go_fuzz_dep_.CoverTab[63207]++
							n := len(p.stack)
							if n == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:422
		_go_fuzz_dep_.CoverTab[63219]++
								return "", &Error{ErrMissingRepeatArgument, before[:len(before)-len(after)]}
//line /usr/local/go/src/regexp/syntax/parse.go:423
		// _ = "end of CoverTab[63219]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:424
		_go_fuzz_dep_.CoverTab[63220]++
//line /usr/local/go/src/regexp/syntax/parse.go:424
		// _ = "end of CoverTab[63220]"
//line /usr/local/go/src/regexp/syntax/parse.go:424
	}
//line /usr/local/go/src/regexp/syntax/parse.go:424
	// _ = "end of CoverTab[63207]"
//line /usr/local/go/src/regexp/syntax/parse.go:424
	_go_fuzz_dep_.CoverTab[63208]++
							sub := p.stack[n-1]
							if sub.Op >= opPseudo {
//line /usr/local/go/src/regexp/syntax/parse.go:426
		_go_fuzz_dep_.CoverTab[63221]++
								return "", &Error{ErrMissingRepeatArgument, before[:len(before)-len(after)]}
//line /usr/local/go/src/regexp/syntax/parse.go:427
		// _ = "end of CoverTab[63221]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:428
		_go_fuzz_dep_.CoverTab[63222]++
//line /usr/local/go/src/regexp/syntax/parse.go:428
		// _ = "end of CoverTab[63222]"
//line /usr/local/go/src/regexp/syntax/parse.go:428
	}
//line /usr/local/go/src/regexp/syntax/parse.go:428
	// _ = "end of CoverTab[63208]"
//line /usr/local/go/src/regexp/syntax/parse.go:428
	_go_fuzz_dep_.CoverTab[63209]++

							re := p.newRegexp(op)
							re.Min = min
							re.Max = max
							re.Flags = flags
							re.Sub = re.Sub0[:1]
							re.Sub[0] = sub
							p.stack[n-1] = re
							p.checkLimits(re)

							if op == OpRepeat && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:439
		_go_fuzz_dep_.CoverTab[63223]++
//line /usr/local/go/src/regexp/syntax/parse.go:439
		return (min >= 2 || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:439
			_go_fuzz_dep_.CoverTab[63224]++
//line /usr/local/go/src/regexp/syntax/parse.go:439
			return max >= 2
//line /usr/local/go/src/regexp/syntax/parse.go:439
			// _ = "end of CoverTab[63224]"
//line /usr/local/go/src/regexp/syntax/parse.go:439
		}())
//line /usr/local/go/src/regexp/syntax/parse.go:439
		// _ = "end of CoverTab[63223]"
//line /usr/local/go/src/regexp/syntax/parse.go:439
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:439
		_go_fuzz_dep_.CoverTab[63225]++
//line /usr/local/go/src/regexp/syntax/parse.go:439
		return !repeatIsValid(re, 1000)
//line /usr/local/go/src/regexp/syntax/parse.go:439
		// _ = "end of CoverTab[63225]"
//line /usr/local/go/src/regexp/syntax/parse.go:439
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:439
		_go_fuzz_dep_.CoverTab[63226]++
								return "", &Error{ErrInvalidRepeatSize, before[:len(before)-len(after)]}
//line /usr/local/go/src/regexp/syntax/parse.go:440
		// _ = "end of CoverTab[63226]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:441
		_go_fuzz_dep_.CoverTab[63227]++
//line /usr/local/go/src/regexp/syntax/parse.go:441
		// _ = "end of CoverTab[63227]"
//line /usr/local/go/src/regexp/syntax/parse.go:441
	}
//line /usr/local/go/src/regexp/syntax/parse.go:441
	// _ = "end of CoverTab[63209]"
//line /usr/local/go/src/regexp/syntax/parse.go:441
	_go_fuzz_dep_.CoverTab[63210]++

							return after, nil
//line /usr/local/go/src/regexp/syntax/parse.go:443
	// _ = "end of CoverTab[63210]"
}

// repeatIsValid reports whether the repetition re is valid.
//line /usr/local/go/src/regexp/syntax/parse.go:446
// Valid means that the combination of the top-level repetition
//line /usr/local/go/src/regexp/syntax/parse.go:446
// and any inner repetitions does not exceed n copies of the
//line /usr/local/go/src/regexp/syntax/parse.go:446
// innermost thing.
//line /usr/local/go/src/regexp/syntax/parse.go:446
// This function rewalks the regexp tree and is called for every repetition,
//line /usr/local/go/src/regexp/syntax/parse.go:446
// so we have to worry about inducing quadratic behavior in the parser.
//line /usr/local/go/src/regexp/syntax/parse.go:446
// We avoid this by only calling repeatIsValid when min or max >= 2.
//line /usr/local/go/src/regexp/syntax/parse.go:446
// In that case the depth of any >= 2 nesting can only get to 9 without
//line /usr/local/go/src/regexp/syntax/parse.go:446
// triggering a parse error, so each subtree can only be rewalked 9 times.
//line /usr/local/go/src/regexp/syntax/parse.go:455
func repeatIsValid(re *Regexp, n int) bool {
//line /usr/local/go/src/regexp/syntax/parse.go:455
	_go_fuzz_dep_.CoverTab[63228]++
							if re.Op == OpRepeat {
//line /usr/local/go/src/regexp/syntax/parse.go:456
		_go_fuzz_dep_.CoverTab[63231]++
								m := re.Max
								if m == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:458
			_go_fuzz_dep_.CoverTab[63235]++
									return true
//line /usr/local/go/src/regexp/syntax/parse.go:459
			// _ = "end of CoverTab[63235]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:460
			_go_fuzz_dep_.CoverTab[63236]++
//line /usr/local/go/src/regexp/syntax/parse.go:460
			// _ = "end of CoverTab[63236]"
//line /usr/local/go/src/regexp/syntax/parse.go:460
		}
//line /usr/local/go/src/regexp/syntax/parse.go:460
		// _ = "end of CoverTab[63231]"
//line /usr/local/go/src/regexp/syntax/parse.go:460
		_go_fuzz_dep_.CoverTab[63232]++
								if m < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:461
			_go_fuzz_dep_.CoverTab[63237]++
									m = re.Min
//line /usr/local/go/src/regexp/syntax/parse.go:462
			// _ = "end of CoverTab[63237]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:463
			_go_fuzz_dep_.CoverTab[63238]++
//line /usr/local/go/src/regexp/syntax/parse.go:463
			// _ = "end of CoverTab[63238]"
//line /usr/local/go/src/regexp/syntax/parse.go:463
		}
//line /usr/local/go/src/regexp/syntax/parse.go:463
		// _ = "end of CoverTab[63232]"
//line /usr/local/go/src/regexp/syntax/parse.go:463
		_go_fuzz_dep_.CoverTab[63233]++
								if m > n {
//line /usr/local/go/src/regexp/syntax/parse.go:464
			_go_fuzz_dep_.CoverTab[63239]++
									return false
//line /usr/local/go/src/regexp/syntax/parse.go:465
			// _ = "end of CoverTab[63239]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:466
			_go_fuzz_dep_.CoverTab[63240]++
//line /usr/local/go/src/regexp/syntax/parse.go:466
			// _ = "end of CoverTab[63240]"
//line /usr/local/go/src/regexp/syntax/parse.go:466
		}
//line /usr/local/go/src/regexp/syntax/parse.go:466
		// _ = "end of CoverTab[63233]"
//line /usr/local/go/src/regexp/syntax/parse.go:466
		_go_fuzz_dep_.CoverTab[63234]++
								if m > 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:467
			_go_fuzz_dep_.CoverTab[63241]++
									n /= m
//line /usr/local/go/src/regexp/syntax/parse.go:468
			// _ = "end of CoverTab[63241]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:469
			_go_fuzz_dep_.CoverTab[63242]++
//line /usr/local/go/src/regexp/syntax/parse.go:469
			// _ = "end of CoverTab[63242]"
//line /usr/local/go/src/regexp/syntax/parse.go:469
		}
//line /usr/local/go/src/regexp/syntax/parse.go:469
		// _ = "end of CoverTab[63234]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:470
		_go_fuzz_dep_.CoverTab[63243]++
//line /usr/local/go/src/regexp/syntax/parse.go:470
		// _ = "end of CoverTab[63243]"
//line /usr/local/go/src/regexp/syntax/parse.go:470
	}
//line /usr/local/go/src/regexp/syntax/parse.go:470
	// _ = "end of CoverTab[63228]"
//line /usr/local/go/src/regexp/syntax/parse.go:470
	_go_fuzz_dep_.CoverTab[63229]++
							for _, sub := range re.Sub {
//line /usr/local/go/src/regexp/syntax/parse.go:471
		_go_fuzz_dep_.CoverTab[63244]++
								if !repeatIsValid(sub, n) {
//line /usr/local/go/src/regexp/syntax/parse.go:472
			_go_fuzz_dep_.CoverTab[63245]++
									return false
//line /usr/local/go/src/regexp/syntax/parse.go:473
			// _ = "end of CoverTab[63245]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:474
			_go_fuzz_dep_.CoverTab[63246]++
//line /usr/local/go/src/regexp/syntax/parse.go:474
			// _ = "end of CoverTab[63246]"
//line /usr/local/go/src/regexp/syntax/parse.go:474
		}
//line /usr/local/go/src/regexp/syntax/parse.go:474
		// _ = "end of CoverTab[63244]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:475
	// _ = "end of CoverTab[63229]"
//line /usr/local/go/src/regexp/syntax/parse.go:475
	_go_fuzz_dep_.CoverTab[63230]++
							return true
//line /usr/local/go/src/regexp/syntax/parse.go:476
	// _ = "end of CoverTab[63230]"
}

// concat replaces the top of the stack (above the topmost '|' or '(') with its concatenation.
func (p *parser) concat() *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:480
	_go_fuzz_dep_.CoverTab[63247]++
							p.maybeConcat(-1, 0)

//line /usr/local/go/src/regexp/syntax/parse.go:484
	i := len(p.stack)
	for i > 0 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:485
		_go_fuzz_dep_.CoverTab[63250]++
//line /usr/local/go/src/regexp/syntax/parse.go:485
		return p.stack[i-1].Op < opPseudo
//line /usr/local/go/src/regexp/syntax/parse.go:485
		// _ = "end of CoverTab[63250]"
//line /usr/local/go/src/regexp/syntax/parse.go:485
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:485
		_go_fuzz_dep_.CoverTab[63251]++
								i--
//line /usr/local/go/src/regexp/syntax/parse.go:486
		// _ = "end of CoverTab[63251]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:487
	// _ = "end of CoverTab[63247]"
//line /usr/local/go/src/regexp/syntax/parse.go:487
	_go_fuzz_dep_.CoverTab[63248]++
							subs := p.stack[i:]
							p.stack = p.stack[:i]

//line /usr/local/go/src/regexp/syntax/parse.go:492
	if len(subs) == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:492
		_go_fuzz_dep_.CoverTab[63252]++
								return p.push(p.newRegexp(OpEmptyMatch))
//line /usr/local/go/src/regexp/syntax/parse.go:493
		// _ = "end of CoverTab[63252]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:494
		_go_fuzz_dep_.CoverTab[63253]++
//line /usr/local/go/src/regexp/syntax/parse.go:494
		// _ = "end of CoverTab[63253]"
//line /usr/local/go/src/regexp/syntax/parse.go:494
	}
//line /usr/local/go/src/regexp/syntax/parse.go:494
	// _ = "end of CoverTab[63248]"
//line /usr/local/go/src/regexp/syntax/parse.go:494
	_go_fuzz_dep_.CoverTab[63249]++

							return p.push(p.collapse(subs, OpConcat))
//line /usr/local/go/src/regexp/syntax/parse.go:496
	// _ = "end of CoverTab[63249]"
}

// alternate replaces the top of the stack (above the topmost '(') with its alternation.
func (p *parser) alternate() *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:500
	_go_fuzz_dep_.CoverTab[63254]++

//line /usr/local/go/src/regexp/syntax/parse.go:503
	i := len(p.stack)
	for i > 0 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:504
		_go_fuzz_dep_.CoverTab[63258]++
//line /usr/local/go/src/regexp/syntax/parse.go:504
		return p.stack[i-1].Op < opPseudo
//line /usr/local/go/src/regexp/syntax/parse.go:504
		// _ = "end of CoverTab[63258]"
//line /usr/local/go/src/regexp/syntax/parse.go:504
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:504
		_go_fuzz_dep_.CoverTab[63259]++
								i--
//line /usr/local/go/src/regexp/syntax/parse.go:505
		// _ = "end of CoverTab[63259]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:506
	// _ = "end of CoverTab[63254]"
//line /usr/local/go/src/regexp/syntax/parse.go:506
	_go_fuzz_dep_.CoverTab[63255]++
							subs := p.stack[i:]
							p.stack = p.stack[:i]

//line /usr/local/go/src/regexp/syntax/parse.go:512
	if len(subs) > 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:512
		_go_fuzz_dep_.CoverTab[63260]++
								cleanAlt(subs[len(subs)-1])
//line /usr/local/go/src/regexp/syntax/parse.go:513
		// _ = "end of CoverTab[63260]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:514
		_go_fuzz_dep_.CoverTab[63261]++
//line /usr/local/go/src/regexp/syntax/parse.go:514
		// _ = "end of CoverTab[63261]"
//line /usr/local/go/src/regexp/syntax/parse.go:514
	}
//line /usr/local/go/src/regexp/syntax/parse.go:514
	// _ = "end of CoverTab[63255]"
//line /usr/local/go/src/regexp/syntax/parse.go:514
	_go_fuzz_dep_.CoverTab[63256]++

//line /usr/local/go/src/regexp/syntax/parse.go:518
	if len(subs) == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:518
		_go_fuzz_dep_.CoverTab[63262]++
								return p.push(p.newRegexp(OpNoMatch))
//line /usr/local/go/src/regexp/syntax/parse.go:519
		// _ = "end of CoverTab[63262]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:520
		_go_fuzz_dep_.CoverTab[63263]++
//line /usr/local/go/src/regexp/syntax/parse.go:520
		// _ = "end of CoverTab[63263]"
//line /usr/local/go/src/regexp/syntax/parse.go:520
	}
//line /usr/local/go/src/regexp/syntax/parse.go:520
	// _ = "end of CoverTab[63256]"
//line /usr/local/go/src/regexp/syntax/parse.go:520
	_go_fuzz_dep_.CoverTab[63257]++

							return p.push(p.collapse(subs, OpAlternate))
//line /usr/local/go/src/regexp/syntax/parse.go:522
	// _ = "end of CoverTab[63257]"
}

// cleanAlt cleans re for eventual inclusion in an alternation.
func cleanAlt(re *Regexp) {
//line /usr/local/go/src/regexp/syntax/parse.go:526
	_go_fuzz_dep_.CoverTab[63264]++
							switch re.Op {
	case OpCharClass:
//line /usr/local/go/src/regexp/syntax/parse.go:528
		_go_fuzz_dep_.CoverTab[63265]++
								re.Rune = cleanClass(&re.Rune)
								if len(re.Rune) == 2 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:530
			_go_fuzz_dep_.CoverTab[63269]++
//line /usr/local/go/src/regexp/syntax/parse.go:530
			return re.Rune[0] == 0
//line /usr/local/go/src/regexp/syntax/parse.go:530
			// _ = "end of CoverTab[63269]"
//line /usr/local/go/src/regexp/syntax/parse.go:530
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:530
			_go_fuzz_dep_.CoverTab[63270]++
//line /usr/local/go/src/regexp/syntax/parse.go:530
			return re.Rune[1] == unicode.MaxRune
//line /usr/local/go/src/regexp/syntax/parse.go:530
			// _ = "end of CoverTab[63270]"
//line /usr/local/go/src/regexp/syntax/parse.go:530
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:530
			_go_fuzz_dep_.CoverTab[63271]++
									re.Rune = nil
									re.Op = OpAnyChar
									return
//line /usr/local/go/src/regexp/syntax/parse.go:533
			// _ = "end of CoverTab[63271]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:534
			_go_fuzz_dep_.CoverTab[63272]++
//line /usr/local/go/src/regexp/syntax/parse.go:534
			// _ = "end of CoverTab[63272]"
//line /usr/local/go/src/regexp/syntax/parse.go:534
		}
//line /usr/local/go/src/regexp/syntax/parse.go:534
		// _ = "end of CoverTab[63265]"
//line /usr/local/go/src/regexp/syntax/parse.go:534
		_go_fuzz_dep_.CoverTab[63266]++
								if len(re.Rune) == 4 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:535
			_go_fuzz_dep_.CoverTab[63273]++
//line /usr/local/go/src/regexp/syntax/parse.go:535
			return re.Rune[0] == 0
//line /usr/local/go/src/regexp/syntax/parse.go:535
			// _ = "end of CoverTab[63273]"
//line /usr/local/go/src/regexp/syntax/parse.go:535
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:535
			_go_fuzz_dep_.CoverTab[63274]++
//line /usr/local/go/src/regexp/syntax/parse.go:535
			return re.Rune[1] == '\n'-1
//line /usr/local/go/src/regexp/syntax/parse.go:535
			// _ = "end of CoverTab[63274]"
//line /usr/local/go/src/regexp/syntax/parse.go:535
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:535
			_go_fuzz_dep_.CoverTab[63275]++
//line /usr/local/go/src/regexp/syntax/parse.go:535
			return re.Rune[2] == '\n'+1
//line /usr/local/go/src/regexp/syntax/parse.go:535
			// _ = "end of CoverTab[63275]"
//line /usr/local/go/src/regexp/syntax/parse.go:535
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:535
			_go_fuzz_dep_.CoverTab[63276]++
//line /usr/local/go/src/regexp/syntax/parse.go:535
			return re.Rune[3] == unicode.MaxRune
//line /usr/local/go/src/regexp/syntax/parse.go:535
			// _ = "end of CoverTab[63276]"
//line /usr/local/go/src/regexp/syntax/parse.go:535
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:535
			_go_fuzz_dep_.CoverTab[63277]++
									re.Rune = nil
									re.Op = OpAnyCharNotNL
									return
//line /usr/local/go/src/regexp/syntax/parse.go:538
			// _ = "end of CoverTab[63277]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:539
			_go_fuzz_dep_.CoverTab[63278]++
//line /usr/local/go/src/regexp/syntax/parse.go:539
			// _ = "end of CoverTab[63278]"
//line /usr/local/go/src/regexp/syntax/parse.go:539
		}
//line /usr/local/go/src/regexp/syntax/parse.go:539
		// _ = "end of CoverTab[63266]"
//line /usr/local/go/src/regexp/syntax/parse.go:539
		_go_fuzz_dep_.CoverTab[63267]++
								if cap(re.Rune)-len(re.Rune) > 100 {
//line /usr/local/go/src/regexp/syntax/parse.go:540
			_go_fuzz_dep_.CoverTab[63279]++

//line /usr/local/go/src/regexp/syntax/parse.go:543
			re.Rune = append(re.Rune0[:0], re.Rune...)
//line /usr/local/go/src/regexp/syntax/parse.go:543
			// _ = "end of CoverTab[63279]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:544
			_go_fuzz_dep_.CoverTab[63280]++
//line /usr/local/go/src/regexp/syntax/parse.go:544
			// _ = "end of CoverTab[63280]"
//line /usr/local/go/src/regexp/syntax/parse.go:544
		}
//line /usr/local/go/src/regexp/syntax/parse.go:544
		// _ = "end of CoverTab[63267]"
//line /usr/local/go/src/regexp/syntax/parse.go:544
	default:
//line /usr/local/go/src/regexp/syntax/parse.go:544
		_go_fuzz_dep_.CoverTab[63268]++
//line /usr/local/go/src/regexp/syntax/parse.go:544
		// _ = "end of CoverTab[63268]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:545
	// _ = "end of CoverTab[63264]"
}

// collapse returns the result of applying op to sub.
//line /usr/local/go/src/regexp/syntax/parse.go:548
// If sub contains op nodes, they all get hoisted up
//line /usr/local/go/src/regexp/syntax/parse.go:548
// so that there is never a concat of a concat or an
//line /usr/local/go/src/regexp/syntax/parse.go:548
// alternate of an alternate.
//line /usr/local/go/src/regexp/syntax/parse.go:552
func (p *parser) collapse(subs []*Regexp, op Op) *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:552
	_go_fuzz_dep_.CoverTab[63281]++
							if len(subs) == 1 {
//line /usr/local/go/src/regexp/syntax/parse.go:553
		_go_fuzz_dep_.CoverTab[63285]++
								return subs[0]
//line /usr/local/go/src/regexp/syntax/parse.go:554
		// _ = "end of CoverTab[63285]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:555
		_go_fuzz_dep_.CoverTab[63286]++
//line /usr/local/go/src/regexp/syntax/parse.go:555
		// _ = "end of CoverTab[63286]"
//line /usr/local/go/src/regexp/syntax/parse.go:555
	}
//line /usr/local/go/src/regexp/syntax/parse.go:555
	// _ = "end of CoverTab[63281]"
//line /usr/local/go/src/regexp/syntax/parse.go:555
	_go_fuzz_dep_.CoverTab[63282]++
							re := p.newRegexp(op)
							re.Sub = re.Sub0[:0]
							for _, sub := range subs {
//line /usr/local/go/src/regexp/syntax/parse.go:558
		_go_fuzz_dep_.CoverTab[63287]++
								if sub.Op == op {
//line /usr/local/go/src/regexp/syntax/parse.go:559
			_go_fuzz_dep_.CoverTab[63288]++
									re.Sub = append(re.Sub, sub.Sub...)
									p.reuse(sub)
//line /usr/local/go/src/regexp/syntax/parse.go:561
			// _ = "end of CoverTab[63288]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:562
			_go_fuzz_dep_.CoverTab[63289]++
									re.Sub = append(re.Sub, sub)
//line /usr/local/go/src/regexp/syntax/parse.go:563
			// _ = "end of CoverTab[63289]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:564
		// _ = "end of CoverTab[63287]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:565
	// _ = "end of CoverTab[63282]"
//line /usr/local/go/src/regexp/syntax/parse.go:565
	_go_fuzz_dep_.CoverTab[63283]++
							if op == OpAlternate {
//line /usr/local/go/src/regexp/syntax/parse.go:566
		_go_fuzz_dep_.CoverTab[63290]++
								re.Sub = p.factor(re.Sub)
								if len(re.Sub) == 1 {
//line /usr/local/go/src/regexp/syntax/parse.go:568
			_go_fuzz_dep_.CoverTab[63291]++
									old := re
									re = re.Sub[0]
									p.reuse(old)
//line /usr/local/go/src/regexp/syntax/parse.go:571
			// _ = "end of CoverTab[63291]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:572
			_go_fuzz_dep_.CoverTab[63292]++
//line /usr/local/go/src/regexp/syntax/parse.go:572
			// _ = "end of CoverTab[63292]"
//line /usr/local/go/src/regexp/syntax/parse.go:572
		}
//line /usr/local/go/src/regexp/syntax/parse.go:572
		// _ = "end of CoverTab[63290]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:573
		_go_fuzz_dep_.CoverTab[63293]++
//line /usr/local/go/src/regexp/syntax/parse.go:573
		// _ = "end of CoverTab[63293]"
//line /usr/local/go/src/regexp/syntax/parse.go:573
	}
//line /usr/local/go/src/regexp/syntax/parse.go:573
	// _ = "end of CoverTab[63283]"
//line /usr/local/go/src/regexp/syntax/parse.go:573
	_go_fuzz_dep_.CoverTab[63284]++
							return re
//line /usr/local/go/src/regexp/syntax/parse.go:574
	// _ = "end of CoverTab[63284]"
}

// factor factors common prefixes from the alternation list sub.
//line /usr/local/go/src/regexp/syntax/parse.go:577
// It returns a replacement list that reuses the same storage and
//line /usr/local/go/src/regexp/syntax/parse.go:577
// frees (passes to p.reuse) any removed *Regexps.
//line /usr/local/go/src/regexp/syntax/parse.go:577
//
//line /usr/local/go/src/regexp/syntax/parse.go:577
// For example,
//line /usr/local/go/src/regexp/syntax/parse.go:577
//
//line /usr/local/go/src/regexp/syntax/parse.go:577
//	ABC|ABD|AEF|BCX|BCY
//line /usr/local/go/src/regexp/syntax/parse.go:577
//
//line /usr/local/go/src/regexp/syntax/parse.go:577
// simplifies by literal prefix extraction to
//line /usr/local/go/src/regexp/syntax/parse.go:577
//
//line /usr/local/go/src/regexp/syntax/parse.go:577
//	A(B(C|D)|EF)|BC(X|Y)
//line /usr/local/go/src/regexp/syntax/parse.go:577
//
//line /usr/local/go/src/regexp/syntax/parse.go:577
// which simplifies by character class introduction to
//line /usr/local/go/src/regexp/syntax/parse.go:577
//
//line /usr/local/go/src/regexp/syntax/parse.go:577
//	A(B[CD]|EF)|BC[XY]
//line /usr/local/go/src/regexp/syntax/parse.go:592
func (p *parser) factor(sub []*Regexp) []*Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:592
	_go_fuzz_dep_.CoverTab[63294]++
							if len(sub) < 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:593
		_go_fuzz_dep_.CoverTab[63300]++
								return sub
//line /usr/local/go/src/regexp/syntax/parse.go:594
		// _ = "end of CoverTab[63300]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:595
		_go_fuzz_dep_.CoverTab[63301]++
//line /usr/local/go/src/regexp/syntax/parse.go:595
		// _ = "end of CoverTab[63301]"
//line /usr/local/go/src/regexp/syntax/parse.go:595
	}
//line /usr/local/go/src/regexp/syntax/parse.go:595
	// _ = "end of CoverTab[63294]"
//line /usr/local/go/src/regexp/syntax/parse.go:595
	_go_fuzz_dep_.CoverTab[63295]++

	// Round 1: Factor out common literal prefixes.
	var str []rune
	var strflags Flags
	start := 0
	out := sub[:0]
	for i := 0; i <= len(sub); i++ {
//line /usr/local/go/src/regexp/syntax/parse.go:602
		_go_fuzz_dep_.CoverTab[63302]++
		// Invariant: the Regexps that were in sub[0:start] have been
		// used or marked for reuse, and the slice space has been reused
		// for out (len(out) <= start).
		//
		// Invariant: sub[start:i] consists of regexps that all begin
		// with str as modified by strflags.
		var istr []rune
		var iflags Flags
		if i < len(sub) {
//line /usr/local/go/src/regexp/syntax/parse.go:611
			_go_fuzz_dep_.CoverTab[63305]++
									istr, iflags = p.leadingString(sub[i])
									if iflags == strflags {
//line /usr/local/go/src/regexp/syntax/parse.go:613
				_go_fuzz_dep_.CoverTab[63306]++
										same := 0
										for same < len(str) && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:615
					_go_fuzz_dep_.CoverTab[63308]++
//line /usr/local/go/src/regexp/syntax/parse.go:615
					return same < len(istr)
//line /usr/local/go/src/regexp/syntax/parse.go:615
					// _ = "end of CoverTab[63308]"
//line /usr/local/go/src/regexp/syntax/parse.go:615
				}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:615
					_go_fuzz_dep_.CoverTab[63309]++
//line /usr/local/go/src/regexp/syntax/parse.go:615
					return str[same] == istr[same]
//line /usr/local/go/src/regexp/syntax/parse.go:615
					// _ = "end of CoverTab[63309]"
//line /usr/local/go/src/regexp/syntax/parse.go:615
				}() {
//line /usr/local/go/src/regexp/syntax/parse.go:615
					_go_fuzz_dep_.CoverTab[63310]++
											same++
//line /usr/local/go/src/regexp/syntax/parse.go:616
					// _ = "end of CoverTab[63310]"
				}
//line /usr/local/go/src/regexp/syntax/parse.go:617
				// _ = "end of CoverTab[63306]"
//line /usr/local/go/src/regexp/syntax/parse.go:617
				_go_fuzz_dep_.CoverTab[63307]++
										if same > 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:618
					_go_fuzz_dep_.CoverTab[63311]++

//line /usr/local/go/src/regexp/syntax/parse.go:621
					str = str[:same]
											continue
//line /usr/local/go/src/regexp/syntax/parse.go:622
					// _ = "end of CoverTab[63311]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:623
					_go_fuzz_dep_.CoverTab[63312]++
//line /usr/local/go/src/regexp/syntax/parse.go:623
					// _ = "end of CoverTab[63312]"
//line /usr/local/go/src/regexp/syntax/parse.go:623
				}
//line /usr/local/go/src/regexp/syntax/parse.go:623
				// _ = "end of CoverTab[63307]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:624
				_go_fuzz_dep_.CoverTab[63313]++
//line /usr/local/go/src/regexp/syntax/parse.go:624
				// _ = "end of CoverTab[63313]"
//line /usr/local/go/src/regexp/syntax/parse.go:624
			}
//line /usr/local/go/src/regexp/syntax/parse.go:624
			// _ = "end of CoverTab[63305]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:625
			_go_fuzz_dep_.CoverTab[63314]++
//line /usr/local/go/src/regexp/syntax/parse.go:625
			// _ = "end of CoverTab[63314]"
//line /usr/local/go/src/regexp/syntax/parse.go:625
		}
//line /usr/local/go/src/regexp/syntax/parse.go:625
		// _ = "end of CoverTab[63302]"
//line /usr/local/go/src/regexp/syntax/parse.go:625
		_go_fuzz_dep_.CoverTab[63303]++

//line /usr/local/go/src/regexp/syntax/parse.go:632
		if i == start {
//line /usr/local/go/src/regexp/syntax/parse.go:632
			_go_fuzz_dep_.CoverTab[63315]++
//line /usr/local/go/src/regexp/syntax/parse.go:632
			// _ = "end of CoverTab[63315]"

		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:634
			_go_fuzz_dep_.CoverTab[63316]++
//line /usr/local/go/src/regexp/syntax/parse.go:634
			if i == start+1 {
//line /usr/local/go/src/regexp/syntax/parse.go:634
				_go_fuzz_dep_.CoverTab[63317]++

										out = append(out, sub[start])
//line /usr/local/go/src/regexp/syntax/parse.go:636
				// _ = "end of CoverTab[63317]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:637
				_go_fuzz_dep_.CoverTab[63318]++

										prefix := p.newRegexp(OpLiteral)
										prefix.Flags = strflags
										prefix.Rune = append(prefix.Rune[:0], str...)

										for j := start; j < i; j++ {
//line /usr/local/go/src/regexp/syntax/parse.go:643
					_go_fuzz_dep_.CoverTab[63320]++
											sub[j] = p.removeLeadingString(sub[j], len(str))
											p.checkLimits(sub[j])
//line /usr/local/go/src/regexp/syntax/parse.go:645
					// _ = "end of CoverTab[63320]"
				}
//line /usr/local/go/src/regexp/syntax/parse.go:646
				// _ = "end of CoverTab[63318]"
//line /usr/local/go/src/regexp/syntax/parse.go:646
				_go_fuzz_dep_.CoverTab[63319]++
										suffix := p.collapse(sub[start:i], OpAlternate)

										re := p.newRegexp(OpConcat)
										re.Sub = append(re.Sub[:0], prefix, suffix)
										out = append(out, re)
//line /usr/local/go/src/regexp/syntax/parse.go:651
				// _ = "end of CoverTab[63319]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:652
			// _ = "end of CoverTab[63316]"
//line /usr/local/go/src/regexp/syntax/parse.go:652
		}
//line /usr/local/go/src/regexp/syntax/parse.go:652
		// _ = "end of CoverTab[63303]"
//line /usr/local/go/src/regexp/syntax/parse.go:652
		_go_fuzz_dep_.CoverTab[63304]++

//line /usr/local/go/src/regexp/syntax/parse.go:655
		start = i
								str = istr
								strflags = iflags
//line /usr/local/go/src/regexp/syntax/parse.go:657
		// _ = "end of CoverTab[63304]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:658
	// _ = "end of CoverTab[63295]"
//line /usr/local/go/src/regexp/syntax/parse.go:658
	_go_fuzz_dep_.CoverTab[63296]++
							sub = out

//line /usr/local/go/src/regexp/syntax/parse.go:669
	start = 0
	out = sub[:0]
	var first *Regexp
	for i := 0; i <= len(sub); i++ {
//line /usr/local/go/src/regexp/syntax/parse.go:672
		_go_fuzz_dep_.CoverTab[63321]++
		// Invariant: the Regexps that were in sub[0:start] have been
		// used or marked for reuse, and the slice space has been reused
		// for out (len(out) <= start).
		//
		// Invariant: sub[start:i] consists of regexps that all begin with ifirst.
		var ifirst *Regexp
		if i < len(sub) {
//line /usr/local/go/src/regexp/syntax/parse.go:679
			_go_fuzz_dep_.CoverTab[63324]++
									ifirst = p.leadingRegexp(sub[i])
									if first != nil && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:681
				_go_fuzz_dep_.CoverTab[63325]++
//line /usr/local/go/src/regexp/syntax/parse.go:681
				return first.Equal(ifirst)
//line /usr/local/go/src/regexp/syntax/parse.go:681
				// _ = "end of CoverTab[63325]"
//line /usr/local/go/src/regexp/syntax/parse.go:681
			}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:681
				_go_fuzz_dep_.CoverTab[63326]++
//line /usr/local/go/src/regexp/syntax/parse.go:681
				return (isCharClass(first) || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:683
					_go_fuzz_dep_.CoverTab[63327]++
//line /usr/local/go/src/regexp/syntax/parse.go:683
					return (first.Op == OpRepeat && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:683
						_go_fuzz_dep_.CoverTab[63328]++
//line /usr/local/go/src/regexp/syntax/parse.go:683
						return first.Min == first.Max
//line /usr/local/go/src/regexp/syntax/parse.go:683
						// _ = "end of CoverTab[63328]"
//line /usr/local/go/src/regexp/syntax/parse.go:683
					}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:683
						_go_fuzz_dep_.CoverTab[63329]++
//line /usr/local/go/src/regexp/syntax/parse.go:683
						return isCharClass(first.Sub[0])
//line /usr/local/go/src/regexp/syntax/parse.go:683
						// _ = "end of CoverTab[63329]"
//line /usr/local/go/src/regexp/syntax/parse.go:683
					}())
//line /usr/local/go/src/regexp/syntax/parse.go:683
					// _ = "end of CoverTab[63327]"
//line /usr/local/go/src/regexp/syntax/parse.go:683
				}())
//line /usr/local/go/src/regexp/syntax/parse.go:683
				// _ = "end of CoverTab[63326]"
//line /usr/local/go/src/regexp/syntax/parse.go:683
			}() {
//line /usr/local/go/src/regexp/syntax/parse.go:683
				_go_fuzz_dep_.CoverTab[63330]++
										continue
//line /usr/local/go/src/regexp/syntax/parse.go:684
				// _ = "end of CoverTab[63330]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:685
				_go_fuzz_dep_.CoverTab[63331]++
//line /usr/local/go/src/regexp/syntax/parse.go:685
				// _ = "end of CoverTab[63331]"
//line /usr/local/go/src/regexp/syntax/parse.go:685
			}
//line /usr/local/go/src/regexp/syntax/parse.go:685
			// _ = "end of CoverTab[63324]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:686
			_go_fuzz_dep_.CoverTab[63332]++
//line /usr/local/go/src/regexp/syntax/parse.go:686
			// _ = "end of CoverTab[63332]"
//line /usr/local/go/src/regexp/syntax/parse.go:686
		}
//line /usr/local/go/src/regexp/syntax/parse.go:686
		// _ = "end of CoverTab[63321]"
//line /usr/local/go/src/regexp/syntax/parse.go:686
		_go_fuzz_dep_.CoverTab[63322]++

//line /usr/local/go/src/regexp/syntax/parse.go:692
		if i == start {
//line /usr/local/go/src/regexp/syntax/parse.go:692
			_go_fuzz_dep_.CoverTab[63333]++
//line /usr/local/go/src/regexp/syntax/parse.go:692
			// _ = "end of CoverTab[63333]"

		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:694
			_go_fuzz_dep_.CoverTab[63334]++
//line /usr/local/go/src/regexp/syntax/parse.go:694
			if i == start+1 {
//line /usr/local/go/src/regexp/syntax/parse.go:694
				_go_fuzz_dep_.CoverTab[63335]++

										out = append(out, sub[start])
//line /usr/local/go/src/regexp/syntax/parse.go:696
				// _ = "end of CoverTab[63335]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:697
				_go_fuzz_dep_.CoverTab[63336]++

										prefix := first
										for j := start; j < i; j++ {
//line /usr/local/go/src/regexp/syntax/parse.go:700
					_go_fuzz_dep_.CoverTab[63338]++
											reuse := j != start
											sub[j] = p.removeLeadingRegexp(sub[j], reuse)
											p.checkLimits(sub[j])
//line /usr/local/go/src/regexp/syntax/parse.go:703
					// _ = "end of CoverTab[63338]"
				}
//line /usr/local/go/src/regexp/syntax/parse.go:704
				// _ = "end of CoverTab[63336]"
//line /usr/local/go/src/regexp/syntax/parse.go:704
				_go_fuzz_dep_.CoverTab[63337]++
										suffix := p.collapse(sub[start:i], OpAlternate)

										re := p.newRegexp(OpConcat)
										re.Sub = append(re.Sub[:0], prefix, suffix)
										out = append(out, re)
//line /usr/local/go/src/regexp/syntax/parse.go:709
				// _ = "end of CoverTab[63337]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:710
			// _ = "end of CoverTab[63334]"
//line /usr/local/go/src/regexp/syntax/parse.go:710
		}
//line /usr/local/go/src/regexp/syntax/parse.go:710
		// _ = "end of CoverTab[63322]"
//line /usr/local/go/src/regexp/syntax/parse.go:710
		_go_fuzz_dep_.CoverTab[63323]++

//line /usr/local/go/src/regexp/syntax/parse.go:713
		start = i
								first = ifirst
//line /usr/local/go/src/regexp/syntax/parse.go:714
		// _ = "end of CoverTab[63323]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:715
	// _ = "end of CoverTab[63296]"
//line /usr/local/go/src/regexp/syntax/parse.go:715
	_go_fuzz_dep_.CoverTab[63297]++
							sub = out

//line /usr/local/go/src/regexp/syntax/parse.go:719
	start = 0
	out = sub[:0]
	for i := 0; i <= len(sub); i++ {
//line /usr/local/go/src/regexp/syntax/parse.go:721
		_go_fuzz_dep_.CoverTab[63339]++

//line /usr/local/go/src/regexp/syntax/parse.go:728
		if i < len(sub) && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:728
			_go_fuzz_dep_.CoverTab[63343]++
//line /usr/local/go/src/regexp/syntax/parse.go:728
			return isCharClass(sub[i])
//line /usr/local/go/src/regexp/syntax/parse.go:728
			// _ = "end of CoverTab[63343]"
//line /usr/local/go/src/regexp/syntax/parse.go:728
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:728
			_go_fuzz_dep_.CoverTab[63344]++
									continue
//line /usr/local/go/src/regexp/syntax/parse.go:729
			// _ = "end of CoverTab[63344]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:730
			_go_fuzz_dep_.CoverTab[63345]++
//line /usr/local/go/src/regexp/syntax/parse.go:730
			// _ = "end of CoverTab[63345]"
//line /usr/local/go/src/regexp/syntax/parse.go:730
		}
//line /usr/local/go/src/regexp/syntax/parse.go:730
		// _ = "end of CoverTab[63339]"
//line /usr/local/go/src/regexp/syntax/parse.go:730
		_go_fuzz_dep_.CoverTab[63340]++

//line /usr/local/go/src/regexp/syntax/parse.go:734
		if i == start {
//line /usr/local/go/src/regexp/syntax/parse.go:734
			_go_fuzz_dep_.CoverTab[63346]++
//line /usr/local/go/src/regexp/syntax/parse.go:734
			// _ = "end of CoverTab[63346]"

		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:736
			_go_fuzz_dep_.CoverTab[63347]++
//line /usr/local/go/src/regexp/syntax/parse.go:736
			if i == start+1 {
//line /usr/local/go/src/regexp/syntax/parse.go:736
				_go_fuzz_dep_.CoverTab[63348]++
										out = append(out, sub[start])
//line /usr/local/go/src/regexp/syntax/parse.go:737
				// _ = "end of CoverTab[63348]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:738
				_go_fuzz_dep_.CoverTab[63349]++

//line /usr/local/go/src/regexp/syntax/parse.go:741
				max := start
				for j := start + 1; j < i; j++ {
//line /usr/local/go/src/regexp/syntax/parse.go:742
					_go_fuzz_dep_.CoverTab[63352]++
											if sub[max].Op < sub[j].Op || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:743
						_go_fuzz_dep_.CoverTab[63353]++
//line /usr/local/go/src/regexp/syntax/parse.go:743
						return sub[max].Op == sub[j].Op && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:743
							_go_fuzz_dep_.CoverTab[63354]++
//line /usr/local/go/src/regexp/syntax/parse.go:743
							return len(sub[max].Rune) < len(sub[j].Rune)
//line /usr/local/go/src/regexp/syntax/parse.go:743
							// _ = "end of CoverTab[63354]"
//line /usr/local/go/src/regexp/syntax/parse.go:743
						}()
//line /usr/local/go/src/regexp/syntax/parse.go:743
						// _ = "end of CoverTab[63353]"
//line /usr/local/go/src/regexp/syntax/parse.go:743
					}() {
//line /usr/local/go/src/regexp/syntax/parse.go:743
						_go_fuzz_dep_.CoverTab[63355]++
												max = j
//line /usr/local/go/src/regexp/syntax/parse.go:744
						// _ = "end of CoverTab[63355]"
					} else {
//line /usr/local/go/src/regexp/syntax/parse.go:745
						_go_fuzz_dep_.CoverTab[63356]++
//line /usr/local/go/src/regexp/syntax/parse.go:745
						// _ = "end of CoverTab[63356]"
//line /usr/local/go/src/regexp/syntax/parse.go:745
					}
//line /usr/local/go/src/regexp/syntax/parse.go:745
					// _ = "end of CoverTab[63352]"
				}
//line /usr/local/go/src/regexp/syntax/parse.go:746
				// _ = "end of CoverTab[63349]"
//line /usr/local/go/src/regexp/syntax/parse.go:746
				_go_fuzz_dep_.CoverTab[63350]++
										sub[start], sub[max] = sub[max], sub[start]

										for j := start + 1; j < i; j++ {
//line /usr/local/go/src/regexp/syntax/parse.go:749
					_go_fuzz_dep_.CoverTab[63357]++
											mergeCharClass(sub[start], sub[j])
											p.reuse(sub[j])
//line /usr/local/go/src/regexp/syntax/parse.go:751
					// _ = "end of CoverTab[63357]"
				}
//line /usr/local/go/src/regexp/syntax/parse.go:752
				// _ = "end of CoverTab[63350]"
//line /usr/local/go/src/regexp/syntax/parse.go:752
				_go_fuzz_dep_.CoverTab[63351]++
										cleanAlt(sub[start])
										out = append(out, sub[start])
//line /usr/local/go/src/regexp/syntax/parse.go:754
				// _ = "end of CoverTab[63351]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:755
			// _ = "end of CoverTab[63347]"
//line /usr/local/go/src/regexp/syntax/parse.go:755
		}
//line /usr/local/go/src/regexp/syntax/parse.go:755
		// _ = "end of CoverTab[63340]"
//line /usr/local/go/src/regexp/syntax/parse.go:755
		_go_fuzz_dep_.CoverTab[63341]++

//line /usr/local/go/src/regexp/syntax/parse.go:758
		if i < len(sub) {
//line /usr/local/go/src/regexp/syntax/parse.go:758
			_go_fuzz_dep_.CoverTab[63358]++
									out = append(out, sub[i])
//line /usr/local/go/src/regexp/syntax/parse.go:759
			// _ = "end of CoverTab[63358]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:760
			_go_fuzz_dep_.CoverTab[63359]++
//line /usr/local/go/src/regexp/syntax/parse.go:760
			// _ = "end of CoverTab[63359]"
//line /usr/local/go/src/regexp/syntax/parse.go:760
		}
//line /usr/local/go/src/regexp/syntax/parse.go:760
		// _ = "end of CoverTab[63341]"
//line /usr/local/go/src/regexp/syntax/parse.go:760
		_go_fuzz_dep_.CoverTab[63342]++
								start = i + 1
//line /usr/local/go/src/regexp/syntax/parse.go:761
		// _ = "end of CoverTab[63342]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:762
	// _ = "end of CoverTab[63297]"
//line /usr/local/go/src/regexp/syntax/parse.go:762
	_go_fuzz_dep_.CoverTab[63298]++
							sub = out

//line /usr/local/go/src/regexp/syntax/parse.go:766
	start = 0
	out = sub[:0]
	for i := range sub {
//line /usr/local/go/src/regexp/syntax/parse.go:768
		_go_fuzz_dep_.CoverTab[63360]++
								if i+1 < len(sub) && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:769
			_go_fuzz_dep_.CoverTab[63362]++
//line /usr/local/go/src/regexp/syntax/parse.go:769
			return sub[i].Op == OpEmptyMatch
//line /usr/local/go/src/regexp/syntax/parse.go:769
			// _ = "end of CoverTab[63362]"
//line /usr/local/go/src/regexp/syntax/parse.go:769
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:769
			_go_fuzz_dep_.CoverTab[63363]++
//line /usr/local/go/src/regexp/syntax/parse.go:769
			return sub[i+1].Op == OpEmptyMatch
//line /usr/local/go/src/regexp/syntax/parse.go:769
			// _ = "end of CoverTab[63363]"
//line /usr/local/go/src/regexp/syntax/parse.go:769
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:769
			_go_fuzz_dep_.CoverTab[63364]++
									continue
//line /usr/local/go/src/regexp/syntax/parse.go:770
			// _ = "end of CoverTab[63364]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:771
			_go_fuzz_dep_.CoverTab[63365]++
//line /usr/local/go/src/regexp/syntax/parse.go:771
			// _ = "end of CoverTab[63365]"
//line /usr/local/go/src/regexp/syntax/parse.go:771
		}
//line /usr/local/go/src/regexp/syntax/parse.go:771
		// _ = "end of CoverTab[63360]"
//line /usr/local/go/src/regexp/syntax/parse.go:771
		_go_fuzz_dep_.CoverTab[63361]++
								out = append(out, sub[i])
//line /usr/local/go/src/regexp/syntax/parse.go:772
		// _ = "end of CoverTab[63361]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:773
	// _ = "end of CoverTab[63298]"
//line /usr/local/go/src/regexp/syntax/parse.go:773
	_go_fuzz_dep_.CoverTab[63299]++
							sub = out

							return sub
//line /usr/local/go/src/regexp/syntax/parse.go:776
	// _ = "end of CoverTab[63299]"
}

// leadingString returns the leading literal string that re begins with.
//line /usr/local/go/src/regexp/syntax/parse.go:779
// The string refers to storage in re or its children.
//line /usr/local/go/src/regexp/syntax/parse.go:781
func (p *parser) leadingString(re *Regexp) ([]rune, Flags) {
//line /usr/local/go/src/regexp/syntax/parse.go:781
	_go_fuzz_dep_.CoverTab[63366]++
							if re.Op == OpConcat && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:782
		_go_fuzz_dep_.CoverTab[63369]++
//line /usr/local/go/src/regexp/syntax/parse.go:782
		return len(re.Sub) > 0
//line /usr/local/go/src/regexp/syntax/parse.go:782
		// _ = "end of CoverTab[63369]"
//line /usr/local/go/src/regexp/syntax/parse.go:782
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:782
		_go_fuzz_dep_.CoverTab[63370]++
								re = re.Sub[0]
//line /usr/local/go/src/regexp/syntax/parse.go:783
		// _ = "end of CoverTab[63370]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:784
		_go_fuzz_dep_.CoverTab[63371]++
//line /usr/local/go/src/regexp/syntax/parse.go:784
		// _ = "end of CoverTab[63371]"
//line /usr/local/go/src/regexp/syntax/parse.go:784
	}
//line /usr/local/go/src/regexp/syntax/parse.go:784
	// _ = "end of CoverTab[63366]"
//line /usr/local/go/src/regexp/syntax/parse.go:784
	_go_fuzz_dep_.CoverTab[63367]++
							if re.Op != OpLiteral {
//line /usr/local/go/src/regexp/syntax/parse.go:785
		_go_fuzz_dep_.CoverTab[63372]++
								return nil, 0
//line /usr/local/go/src/regexp/syntax/parse.go:786
		// _ = "end of CoverTab[63372]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:787
		_go_fuzz_dep_.CoverTab[63373]++
//line /usr/local/go/src/regexp/syntax/parse.go:787
		// _ = "end of CoverTab[63373]"
//line /usr/local/go/src/regexp/syntax/parse.go:787
	}
//line /usr/local/go/src/regexp/syntax/parse.go:787
	// _ = "end of CoverTab[63367]"
//line /usr/local/go/src/regexp/syntax/parse.go:787
	_go_fuzz_dep_.CoverTab[63368]++
							return re.Rune, re.Flags & FoldCase
//line /usr/local/go/src/regexp/syntax/parse.go:788
	// _ = "end of CoverTab[63368]"
}

// removeLeadingString removes the first n leading runes
//line /usr/local/go/src/regexp/syntax/parse.go:791
// from the beginning of re. It returns the replacement for re.
//line /usr/local/go/src/regexp/syntax/parse.go:793
func (p *parser) removeLeadingString(re *Regexp, n int) *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:793
	_go_fuzz_dep_.CoverTab[63374]++
							if re.Op == OpConcat && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:794
		_go_fuzz_dep_.CoverTab[63377]++
//line /usr/local/go/src/regexp/syntax/parse.go:794
		return len(re.Sub) > 0
//line /usr/local/go/src/regexp/syntax/parse.go:794
		// _ = "end of CoverTab[63377]"
//line /usr/local/go/src/regexp/syntax/parse.go:794
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:794
		_go_fuzz_dep_.CoverTab[63378]++

//line /usr/local/go/src/regexp/syntax/parse.go:797
		sub := re.Sub[0]
		sub = p.removeLeadingString(sub, n)
		re.Sub[0] = sub
		if sub.Op == OpEmptyMatch {
//line /usr/local/go/src/regexp/syntax/parse.go:800
			_go_fuzz_dep_.CoverTab[63380]++
									p.reuse(sub)
									switch len(re.Sub) {
			case 0, 1:
//line /usr/local/go/src/regexp/syntax/parse.go:803
				_go_fuzz_dep_.CoverTab[63381]++

										re.Op = OpEmptyMatch
										re.Sub = nil
//line /usr/local/go/src/regexp/syntax/parse.go:806
				// _ = "end of CoverTab[63381]"
			case 2:
//line /usr/local/go/src/regexp/syntax/parse.go:807
				_go_fuzz_dep_.CoverTab[63382]++
										old := re
										re = re.Sub[1]
										p.reuse(old)
//line /usr/local/go/src/regexp/syntax/parse.go:810
				// _ = "end of CoverTab[63382]"
			default:
//line /usr/local/go/src/regexp/syntax/parse.go:811
				_go_fuzz_dep_.CoverTab[63383]++
										copy(re.Sub, re.Sub[1:])
										re.Sub = re.Sub[:len(re.Sub)-1]
//line /usr/local/go/src/regexp/syntax/parse.go:813
				// _ = "end of CoverTab[63383]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:814
			// _ = "end of CoverTab[63380]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:815
			_go_fuzz_dep_.CoverTab[63384]++
//line /usr/local/go/src/regexp/syntax/parse.go:815
			// _ = "end of CoverTab[63384]"
//line /usr/local/go/src/regexp/syntax/parse.go:815
		}
//line /usr/local/go/src/regexp/syntax/parse.go:815
		// _ = "end of CoverTab[63378]"
//line /usr/local/go/src/regexp/syntax/parse.go:815
		_go_fuzz_dep_.CoverTab[63379]++
								return re
//line /usr/local/go/src/regexp/syntax/parse.go:816
		// _ = "end of CoverTab[63379]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:817
		_go_fuzz_dep_.CoverTab[63385]++
//line /usr/local/go/src/regexp/syntax/parse.go:817
		// _ = "end of CoverTab[63385]"
//line /usr/local/go/src/regexp/syntax/parse.go:817
	}
//line /usr/local/go/src/regexp/syntax/parse.go:817
	// _ = "end of CoverTab[63374]"
//line /usr/local/go/src/regexp/syntax/parse.go:817
	_go_fuzz_dep_.CoverTab[63375]++

							if re.Op == OpLiteral {
//line /usr/local/go/src/regexp/syntax/parse.go:819
		_go_fuzz_dep_.CoverTab[63386]++
								re.Rune = re.Rune[:copy(re.Rune, re.Rune[n:])]
								if len(re.Rune) == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:821
			_go_fuzz_dep_.CoverTab[63387]++
									re.Op = OpEmptyMatch
//line /usr/local/go/src/regexp/syntax/parse.go:822
			// _ = "end of CoverTab[63387]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:823
			_go_fuzz_dep_.CoverTab[63388]++
//line /usr/local/go/src/regexp/syntax/parse.go:823
			// _ = "end of CoverTab[63388]"
//line /usr/local/go/src/regexp/syntax/parse.go:823
		}
//line /usr/local/go/src/regexp/syntax/parse.go:823
		// _ = "end of CoverTab[63386]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:824
		_go_fuzz_dep_.CoverTab[63389]++
//line /usr/local/go/src/regexp/syntax/parse.go:824
		// _ = "end of CoverTab[63389]"
//line /usr/local/go/src/regexp/syntax/parse.go:824
	}
//line /usr/local/go/src/regexp/syntax/parse.go:824
	// _ = "end of CoverTab[63375]"
//line /usr/local/go/src/regexp/syntax/parse.go:824
	_go_fuzz_dep_.CoverTab[63376]++
							return re
//line /usr/local/go/src/regexp/syntax/parse.go:825
	// _ = "end of CoverTab[63376]"
}

// leadingRegexp returns the leading regexp that re begins with.
//line /usr/local/go/src/regexp/syntax/parse.go:828
// The regexp refers to storage in re or its children.
//line /usr/local/go/src/regexp/syntax/parse.go:830
func (p *parser) leadingRegexp(re *Regexp) *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:830
	_go_fuzz_dep_.CoverTab[63390]++
							if re.Op == OpEmptyMatch {
//line /usr/local/go/src/regexp/syntax/parse.go:831
		_go_fuzz_dep_.CoverTab[63393]++
								return nil
//line /usr/local/go/src/regexp/syntax/parse.go:832
		// _ = "end of CoverTab[63393]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:833
		_go_fuzz_dep_.CoverTab[63394]++
//line /usr/local/go/src/regexp/syntax/parse.go:833
		// _ = "end of CoverTab[63394]"
//line /usr/local/go/src/regexp/syntax/parse.go:833
	}
//line /usr/local/go/src/regexp/syntax/parse.go:833
	// _ = "end of CoverTab[63390]"
//line /usr/local/go/src/regexp/syntax/parse.go:833
	_go_fuzz_dep_.CoverTab[63391]++
							if re.Op == OpConcat && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:834
		_go_fuzz_dep_.CoverTab[63395]++
//line /usr/local/go/src/regexp/syntax/parse.go:834
		return len(re.Sub) > 0
//line /usr/local/go/src/regexp/syntax/parse.go:834
		// _ = "end of CoverTab[63395]"
//line /usr/local/go/src/regexp/syntax/parse.go:834
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:834
		_go_fuzz_dep_.CoverTab[63396]++
								sub := re.Sub[0]
								if sub.Op == OpEmptyMatch {
//line /usr/local/go/src/regexp/syntax/parse.go:836
			_go_fuzz_dep_.CoverTab[63398]++
									return nil
//line /usr/local/go/src/regexp/syntax/parse.go:837
			// _ = "end of CoverTab[63398]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:838
			_go_fuzz_dep_.CoverTab[63399]++
//line /usr/local/go/src/regexp/syntax/parse.go:838
			// _ = "end of CoverTab[63399]"
//line /usr/local/go/src/regexp/syntax/parse.go:838
		}
//line /usr/local/go/src/regexp/syntax/parse.go:838
		// _ = "end of CoverTab[63396]"
//line /usr/local/go/src/regexp/syntax/parse.go:838
		_go_fuzz_dep_.CoverTab[63397]++
								return sub
//line /usr/local/go/src/regexp/syntax/parse.go:839
		// _ = "end of CoverTab[63397]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:840
		_go_fuzz_dep_.CoverTab[63400]++
//line /usr/local/go/src/regexp/syntax/parse.go:840
		// _ = "end of CoverTab[63400]"
//line /usr/local/go/src/regexp/syntax/parse.go:840
	}
//line /usr/local/go/src/regexp/syntax/parse.go:840
	// _ = "end of CoverTab[63391]"
//line /usr/local/go/src/regexp/syntax/parse.go:840
	_go_fuzz_dep_.CoverTab[63392]++
							return re
//line /usr/local/go/src/regexp/syntax/parse.go:841
	// _ = "end of CoverTab[63392]"
}

// removeLeadingRegexp removes the leading regexp in re.
//line /usr/local/go/src/regexp/syntax/parse.go:844
// It returns the replacement for re.
//line /usr/local/go/src/regexp/syntax/parse.go:844
// If reuse is true, it passes the removed regexp (if no longer needed) to p.reuse.
//line /usr/local/go/src/regexp/syntax/parse.go:847
func (p *parser) removeLeadingRegexp(re *Regexp, reuse bool) *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:847
	_go_fuzz_dep_.CoverTab[63401]++
							if re.Op == OpConcat && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:848
		_go_fuzz_dep_.CoverTab[63404]++
//line /usr/local/go/src/regexp/syntax/parse.go:848
		return len(re.Sub) > 0
//line /usr/local/go/src/regexp/syntax/parse.go:848
		// _ = "end of CoverTab[63404]"
//line /usr/local/go/src/regexp/syntax/parse.go:848
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:848
		_go_fuzz_dep_.CoverTab[63405]++
								if reuse {
//line /usr/local/go/src/regexp/syntax/parse.go:849
			_go_fuzz_dep_.CoverTab[63408]++
									p.reuse(re.Sub[0])
//line /usr/local/go/src/regexp/syntax/parse.go:850
			// _ = "end of CoverTab[63408]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:851
			_go_fuzz_dep_.CoverTab[63409]++
//line /usr/local/go/src/regexp/syntax/parse.go:851
			// _ = "end of CoverTab[63409]"
//line /usr/local/go/src/regexp/syntax/parse.go:851
		}
//line /usr/local/go/src/regexp/syntax/parse.go:851
		// _ = "end of CoverTab[63405]"
//line /usr/local/go/src/regexp/syntax/parse.go:851
		_go_fuzz_dep_.CoverTab[63406]++
								re.Sub = re.Sub[:copy(re.Sub, re.Sub[1:])]
								switch len(re.Sub) {
		case 0:
//line /usr/local/go/src/regexp/syntax/parse.go:854
			_go_fuzz_dep_.CoverTab[63410]++
									re.Op = OpEmptyMatch
									re.Sub = nil
//line /usr/local/go/src/regexp/syntax/parse.go:856
			// _ = "end of CoverTab[63410]"
		case 1:
//line /usr/local/go/src/regexp/syntax/parse.go:857
			_go_fuzz_dep_.CoverTab[63411]++
									old := re
									re = re.Sub[0]
									p.reuse(old)
//line /usr/local/go/src/regexp/syntax/parse.go:860
			// _ = "end of CoverTab[63411]"
//line /usr/local/go/src/regexp/syntax/parse.go:860
		default:
//line /usr/local/go/src/regexp/syntax/parse.go:860
			_go_fuzz_dep_.CoverTab[63412]++
//line /usr/local/go/src/regexp/syntax/parse.go:860
			// _ = "end of CoverTab[63412]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:861
		// _ = "end of CoverTab[63406]"
//line /usr/local/go/src/regexp/syntax/parse.go:861
		_go_fuzz_dep_.CoverTab[63407]++
								return re
//line /usr/local/go/src/regexp/syntax/parse.go:862
		// _ = "end of CoverTab[63407]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:863
		_go_fuzz_dep_.CoverTab[63413]++
//line /usr/local/go/src/regexp/syntax/parse.go:863
		// _ = "end of CoverTab[63413]"
//line /usr/local/go/src/regexp/syntax/parse.go:863
	}
//line /usr/local/go/src/regexp/syntax/parse.go:863
	// _ = "end of CoverTab[63401]"
//line /usr/local/go/src/regexp/syntax/parse.go:863
	_go_fuzz_dep_.CoverTab[63402]++
							if reuse {
//line /usr/local/go/src/regexp/syntax/parse.go:864
		_go_fuzz_dep_.CoverTab[63414]++
								p.reuse(re)
//line /usr/local/go/src/regexp/syntax/parse.go:865
		// _ = "end of CoverTab[63414]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:866
		_go_fuzz_dep_.CoverTab[63415]++
//line /usr/local/go/src/regexp/syntax/parse.go:866
		// _ = "end of CoverTab[63415]"
//line /usr/local/go/src/regexp/syntax/parse.go:866
	}
//line /usr/local/go/src/regexp/syntax/parse.go:866
	// _ = "end of CoverTab[63402]"
//line /usr/local/go/src/regexp/syntax/parse.go:866
	_go_fuzz_dep_.CoverTab[63403]++
							return p.newRegexp(OpEmptyMatch)
//line /usr/local/go/src/regexp/syntax/parse.go:867
	// _ = "end of CoverTab[63403]"
}

func literalRegexp(s string, flags Flags) *Regexp {
//line /usr/local/go/src/regexp/syntax/parse.go:870
	_go_fuzz_dep_.CoverTab[63416]++
							re := &Regexp{Op: OpLiteral}
							re.Flags = flags
							re.Rune = re.Rune0[:0]
							for _, c := range s {
//line /usr/local/go/src/regexp/syntax/parse.go:874
		_go_fuzz_dep_.CoverTab[63418]++
								if len(re.Rune) >= cap(re.Rune) {
//line /usr/local/go/src/regexp/syntax/parse.go:875
			_go_fuzz_dep_.CoverTab[63420]++

									re.Rune = []rune(s)
									break
//line /usr/local/go/src/regexp/syntax/parse.go:878
			// _ = "end of CoverTab[63420]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:879
			_go_fuzz_dep_.CoverTab[63421]++
//line /usr/local/go/src/regexp/syntax/parse.go:879
			// _ = "end of CoverTab[63421]"
//line /usr/local/go/src/regexp/syntax/parse.go:879
		}
//line /usr/local/go/src/regexp/syntax/parse.go:879
		// _ = "end of CoverTab[63418]"
//line /usr/local/go/src/regexp/syntax/parse.go:879
		_go_fuzz_dep_.CoverTab[63419]++
								re.Rune = append(re.Rune, c)
//line /usr/local/go/src/regexp/syntax/parse.go:880
		// _ = "end of CoverTab[63419]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:881
	// _ = "end of CoverTab[63416]"
//line /usr/local/go/src/regexp/syntax/parse.go:881
	_go_fuzz_dep_.CoverTab[63417]++
							return re
//line /usr/local/go/src/regexp/syntax/parse.go:882
	// _ = "end of CoverTab[63417]"
}

//line /usr/local/go/src/regexp/syntax/parse.go:887
// Parse parses a regular expression string s, controlled by the specified
//line /usr/local/go/src/regexp/syntax/parse.go:887
// Flags, and returns a regular expression parse tree. The syntax is
//line /usr/local/go/src/regexp/syntax/parse.go:887
// described in the top-level comment.
//line /usr/local/go/src/regexp/syntax/parse.go:890
func Parse(s string, flags Flags) (*Regexp, error) {
//line /usr/local/go/src/regexp/syntax/parse.go:890
	_go_fuzz_dep_.CoverTab[63422]++
							return parse(s, flags)
//line /usr/local/go/src/regexp/syntax/parse.go:891
	// _ = "end of CoverTab[63422]"
}

func parse(s string, flags Flags) (_ *Regexp, err error) {
//line /usr/local/go/src/regexp/syntax/parse.go:894
	_go_fuzz_dep_.CoverTab[63423]++
							defer func() {
//line /usr/local/go/src/regexp/syntax/parse.go:895
		_go_fuzz_dep_.CoverTab[63429]++
								switch r := recover(); r {
		default:
//line /usr/local/go/src/regexp/syntax/parse.go:897
			_go_fuzz_dep_.CoverTab[63430]++
									panic(r)
//line /usr/local/go/src/regexp/syntax/parse.go:898
			// _ = "end of CoverTab[63430]"
		case nil:
//line /usr/local/go/src/regexp/syntax/parse.go:899
			_go_fuzz_dep_.CoverTab[63431]++
//line /usr/local/go/src/regexp/syntax/parse.go:899
			// _ = "end of CoverTab[63431]"

		case ErrLarge:
//line /usr/local/go/src/regexp/syntax/parse.go:901
			_go_fuzz_dep_.CoverTab[63432]++
									err = &Error{Code: ErrLarge, Expr: s}
//line /usr/local/go/src/regexp/syntax/parse.go:902
			// _ = "end of CoverTab[63432]"
		case ErrNestingDepth:
//line /usr/local/go/src/regexp/syntax/parse.go:903
			_go_fuzz_dep_.CoverTab[63433]++
									err = &Error{Code: ErrNestingDepth, Expr: s}
//line /usr/local/go/src/regexp/syntax/parse.go:904
			// _ = "end of CoverTab[63433]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:905
		// _ = "end of CoverTab[63429]"
	}()
//line /usr/local/go/src/regexp/syntax/parse.go:906
	// _ = "end of CoverTab[63423]"
//line /usr/local/go/src/regexp/syntax/parse.go:906
	_go_fuzz_dep_.CoverTab[63424]++

							if flags&Literal != 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:908
		_go_fuzz_dep_.CoverTab[63434]++

								if err := checkUTF8(s); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:910
			_go_fuzz_dep_.CoverTab[63436]++
									return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:911
			// _ = "end of CoverTab[63436]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:912
			_go_fuzz_dep_.CoverTab[63437]++
//line /usr/local/go/src/regexp/syntax/parse.go:912
			// _ = "end of CoverTab[63437]"
//line /usr/local/go/src/regexp/syntax/parse.go:912
		}
//line /usr/local/go/src/regexp/syntax/parse.go:912
		// _ = "end of CoverTab[63434]"
//line /usr/local/go/src/regexp/syntax/parse.go:912
		_go_fuzz_dep_.CoverTab[63435]++
								return literalRegexp(s, flags), nil
//line /usr/local/go/src/regexp/syntax/parse.go:913
		// _ = "end of CoverTab[63435]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:914
		_go_fuzz_dep_.CoverTab[63438]++
//line /usr/local/go/src/regexp/syntax/parse.go:914
		// _ = "end of CoverTab[63438]"
//line /usr/local/go/src/regexp/syntax/parse.go:914
	}
//line /usr/local/go/src/regexp/syntax/parse.go:914
	// _ = "end of CoverTab[63424]"
//line /usr/local/go/src/regexp/syntax/parse.go:914
	_go_fuzz_dep_.CoverTab[63425]++

	// Otherwise, must do real work.
	var (
		p		parser
		c		rune
		op		Op
		lastRepeat	string
	)
	p.flags = flags
	p.wholeRegexp = s
	t := s
	for t != "" {
//line /usr/local/go/src/regexp/syntax/parse.go:926
		_go_fuzz_dep_.CoverTab[63439]++
								repeat := ""
	BigSwitch:
		switch t[0] {
		default:
//line /usr/local/go/src/regexp/syntax/parse.go:930
			_go_fuzz_dep_.CoverTab[63441]++
									if c, t, err = nextRune(t); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:931
				_go_fuzz_dep_.CoverTab[63468]++
										return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:932
				// _ = "end of CoverTab[63468]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:933
				_go_fuzz_dep_.CoverTab[63469]++
//line /usr/local/go/src/regexp/syntax/parse.go:933
				// _ = "end of CoverTab[63469]"
//line /usr/local/go/src/regexp/syntax/parse.go:933
			}
//line /usr/local/go/src/regexp/syntax/parse.go:933
			// _ = "end of CoverTab[63441]"
//line /usr/local/go/src/regexp/syntax/parse.go:933
			_go_fuzz_dep_.CoverTab[63442]++
									p.literal(c)
//line /usr/local/go/src/regexp/syntax/parse.go:934
			// _ = "end of CoverTab[63442]"

		case '(':
//line /usr/local/go/src/regexp/syntax/parse.go:936
			_go_fuzz_dep_.CoverTab[63443]++
									if p.flags&PerlX != 0 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:937
				_go_fuzz_dep_.CoverTab[63470]++
//line /usr/local/go/src/regexp/syntax/parse.go:937
				return len(t) >= 2
//line /usr/local/go/src/regexp/syntax/parse.go:937
				// _ = "end of CoverTab[63470]"
//line /usr/local/go/src/regexp/syntax/parse.go:937
			}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:937
				_go_fuzz_dep_.CoverTab[63471]++
//line /usr/local/go/src/regexp/syntax/parse.go:937
				return t[1] == '?'
//line /usr/local/go/src/regexp/syntax/parse.go:937
				// _ = "end of CoverTab[63471]"
//line /usr/local/go/src/regexp/syntax/parse.go:937
			}() {
//line /usr/local/go/src/regexp/syntax/parse.go:937
				_go_fuzz_dep_.CoverTab[63472]++

										if t, err = p.parsePerlFlags(t); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:939
					_go_fuzz_dep_.CoverTab[63474]++
											return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:940
					// _ = "end of CoverTab[63474]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:941
					_go_fuzz_dep_.CoverTab[63475]++
//line /usr/local/go/src/regexp/syntax/parse.go:941
					// _ = "end of CoverTab[63475]"
//line /usr/local/go/src/regexp/syntax/parse.go:941
				}
//line /usr/local/go/src/regexp/syntax/parse.go:941
				// _ = "end of CoverTab[63472]"
//line /usr/local/go/src/regexp/syntax/parse.go:941
				_go_fuzz_dep_.CoverTab[63473]++
										break
//line /usr/local/go/src/regexp/syntax/parse.go:942
				// _ = "end of CoverTab[63473]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:943
				_go_fuzz_dep_.CoverTab[63476]++
//line /usr/local/go/src/regexp/syntax/parse.go:943
				// _ = "end of CoverTab[63476]"
//line /usr/local/go/src/regexp/syntax/parse.go:943
			}
//line /usr/local/go/src/regexp/syntax/parse.go:943
			// _ = "end of CoverTab[63443]"
//line /usr/local/go/src/regexp/syntax/parse.go:943
			_go_fuzz_dep_.CoverTab[63444]++
									p.numCap++
									p.op(opLeftParen).Cap = p.numCap
									t = t[1:]
//line /usr/local/go/src/regexp/syntax/parse.go:946
			// _ = "end of CoverTab[63444]"
		case '|':
//line /usr/local/go/src/regexp/syntax/parse.go:947
			_go_fuzz_dep_.CoverTab[63445]++
									if err = p.parseVerticalBar(); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:948
				_go_fuzz_dep_.CoverTab[63477]++
										return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:949
				// _ = "end of CoverTab[63477]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:950
				_go_fuzz_dep_.CoverTab[63478]++
//line /usr/local/go/src/regexp/syntax/parse.go:950
				// _ = "end of CoverTab[63478]"
//line /usr/local/go/src/regexp/syntax/parse.go:950
			}
//line /usr/local/go/src/regexp/syntax/parse.go:950
			// _ = "end of CoverTab[63445]"
//line /usr/local/go/src/regexp/syntax/parse.go:950
			_go_fuzz_dep_.CoverTab[63446]++
									t = t[1:]
//line /usr/local/go/src/regexp/syntax/parse.go:951
			// _ = "end of CoverTab[63446]"
		case ')':
//line /usr/local/go/src/regexp/syntax/parse.go:952
			_go_fuzz_dep_.CoverTab[63447]++
									if err = p.parseRightParen(); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:953
				_go_fuzz_dep_.CoverTab[63479]++
										return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:954
				// _ = "end of CoverTab[63479]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:955
				_go_fuzz_dep_.CoverTab[63480]++
//line /usr/local/go/src/regexp/syntax/parse.go:955
				// _ = "end of CoverTab[63480]"
//line /usr/local/go/src/regexp/syntax/parse.go:955
			}
//line /usr/local/go/src/regexp/syntax/parse.go:955
			// _ = "end of CoverTab[63447]"
//line /usr/local/go/src/regexp/syntax/parse.go:955
			_go_fuzz_dep_.CoverTab[63448]++
									t = t[1:]
//line /usr/local/go/src/regexp/syntax/parse.go:956
			// _ = "end of CoverTab[63448]"
		case '^':
//line /usr/local/go/src/regexp/syntax/parse.go:957
			_go_fuzz_dep_.CoverTab[63449]++
									if p.flags&OneLine != 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:958
				_go_fuzz_dep_.CoverTab[63481]++
										p.op(OpBeginText)
//line /usr/local/go/src/regexp/syntax/parse.go:959
				// _ = "end of CoverTab[63481]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:960
				_go_fuzz_dep_.CoverTab[63482]++
										p.op(OpBeginLine)
//line /usr/local/go/src/regexp/syntax/parse.go:961
				// _ = "end of CoverTab[63482]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:962
			// _ = "end of CoverTab[63449]"
//line /usr/local/go/src/regexp/syntax/parse.go:962
			_go_fuzz_dep_.CoverTab[63450]++
									t = t[1:]
//line /usr/local/go/src/regexp/syntax/parse.go:963
			// _ = "end of CoverTab[63450]"
		case '$':
//line /usr/local/go/src/regexp/syntax/parse.go:964
			_go_fuzz_dep_.CoverTab[63451]++
									if p.flags&OneLine != 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:965
				_go_fuzz_dep_.CoverTab[63483]++
										p.op(OpEndText).Flags |= WasDollar
//line /usr/local/go/src/regexp/syntax/parse.go:966
				// _ = "end of CoverTab[63483]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:967
				_go_fuzz_dep_.CoverTab[63484]++
										p.op(OpEndLine)
//line /usr/local/go/src/regexp/syntax/parse.go:968
				// _ = "end of CoverTab[63484]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:969
			// _ = "end of CoverTab[63451]"
//line /usr/local/go/src/regexp/syntax/parse.go:969
			_go_fuzz_dep_.CoverTab[63452]++
									t = t[1:]
//line /usr/local/go/src/regexp/syntax/parse.go:970
			// _ = "end of CoverTab[63452]"
		case '.':
//line /usr/local/go/src/regexp/syntax/parse.go:971
			_go_fuzz_dep_.CoverTab[63453]++
									if p.flags&DotNL != 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:972
				_go_fuzz_dep_.CoverTab[63485]++
										p.op(OpAnyChar)
//line /usr/local/go/src/regexp/syntax/parse.go:973
				// _ = "end of CoverTab[63485]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:974
				_go_fuzz_dep_.CoverTab[63486]++
										p.op(OpAnyCharNotNL)
//line /usr/local/go/src/regexp/syntax/parse.go:975
				// _ = "end of CoverTab[63486]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:976
			// _ = "end of CoverTab[63453]"
//line /usr/local/go/src/regexp/syntax/parse.go:976
			_go_fuzz_dep_.CoverTab[63454]++
									t = t[1:]
//line /usr/local/go/src/regexp/syntax/parse.go:977
			// _ = "end of CoverTab[63454]"
		case '[':
//line /usr/local/go/src/regexp/syntax/parse.go:978
			_go_fuzz_dep_.CoverTab[63455]++
									if t, err = p.parseClass(t); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:979
				_go_fuzz_dep_.CoverTab[63487]++
										return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:980
				// _ = "end of CoverTab[63487]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:981
				_go_fuzz_dep_.CoverTab[63488]++
//line /usr/local/go/src/regexp/syntax/parse.go:981
				// _ = "end of CoverTab[63488]"
//line /usr/local/go/src/regexp/syntax/parse.go:981
			}
//line /usr/local/go/src/regexp/syntax/parse.go:981
			// _ = "end of CoverTab[63455]"
		case '*', '+', '?':
//line /usr/local/go/src/regexp/syntax/parse.go:982
			_go_fuzz_dep_.CoverTab[63456]++
									before := t
									switch t[0] {
			case '*':
//line /usr/local/go/src/regexp/syntax/parse.go:985
				_go_fuzz_dep_.CoverTab[63489]++
										op = OpStar
//line /usr/local/go/src/regexp/syntax/parse.go:986
				// _ = "end of CoverTab[63489]"
			case '+':
//line /usr/local/go/src/regexp/syntax/parse.go:987
				_go_fuzz_dep_.CoverTab[63490]++
										op = OpPlus
//line /usr/local/go/src/regexp/syntax/parse.go:988
				// _ = "end of CoverTab[63490]"
			case '?':
//line /usr/local/go/src/regexp/syntax/parse.go:989
				_go_fuzz_dep_.CoverTab[63491]++
										op = OpQuest
//line /usr/local/go/src/regexp/syntax/parse.go:990
				// _ = "end of CoverTab[63491]"
//line /usr/local/go/src/regexp/syntax/parse.go:990
			default:
//line /usr/local/go/src/regexp/syntax/parse.go:990
				_go_fuzz_dep_.CoverTab[63492]++
//line /usr/local/go/src/regexp/syntax/parse.go:990
				// _ = "end of CoverTab[63492]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:991
			// _ = "end of CoverTab[63456]"
//line /usr/local/go/src/regexp/syntax/parse.go:991
			_go_fuzz_dep_.CoverTab[63457]++
									after := t[1:]
									if after, err = p.repeat(op, 0, 0, before, after, lastRepeat); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:993
				_go_fuzz_dep_.CoverTab[63493]++
										return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:994
				// _ = "end of CoverTab[63493]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:995
				_go_fuzz_dep_.CoverTab[63494]++
//line /usr/local/go/src/regexp/syntax/parse.go:995
				// _ = "end of CoverTab[63494]"
//line /usr/local/go/src/regexp/syntax/parse.go:995
			}
//line /usr/local/go/src/regexp/syntax/parse.go:995
			// _ = "end of CoverTab[63457]"
//line /usr/local/go/src/regexp/syntax/parse.go:995
			_go_fuzz_dep_.CoverTab[63458]++
									repeat = before
									t = after
//line /usr/local/go/src/regexp/syntax/parse.go:997
			// _ = "end of CoverTab[63458]"
		case '{':
//line /usr/local/go/src/regexp/syntax/parse.go:998
			_go_fuzz_dep_.CoverTab[63459]++
									op = OpRepeat
									before := t
									min, max, after, ok := p.parseRepeat(t)
									if !ok {
//line /usr/local/go/src/regexp/syntax/parse.go:1002
				_go_fuzz_dep_.CoverTab[63495]++

										p.literal('{')
										t = t[1:]
										break
//line /usr/local/go/src/regexp/syntax/parse.go:1006
				// _ = "end of CoverTab[63495]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1007
				_go_fuzz_dep_.CoverTab[63496]++
//line /usr/local/go/src/regexp/syntax/parse.go:1007
				// _ = "end of CoverTab[63496]"
//line /usr/local/go/src/regexp/syntax/parse.go:1007
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1007
			// _ = "end of CoverTab[63459]"
//line /usr/local/go/src/regexp/syntax/parse.go:1007
			_go_fuzz_dep_.CoverTab[63460]++
									if min < 0 || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				_go_fuzz_dep_.CoverTab[63497]++
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				return min > 1000
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				// _ = "end of CoverTab[63497]"
//line /usr/local/go/src/regexp/syntax/parse.go:1008
			}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				_go_fuzz_dep_.CoverTab[63498]++
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				return max > 1000
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				// _ = "end of CoverTab[63498]"
//line /usr/local/go/src/regexp/syntax/parse.go:1008
			}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				_go_fuzz_dep_.CoverTab[63499]++
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				return max >= 0 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1008
					_go_fuzz_dep_.CoverTab[63500]++
//line /usr/local/go/src/regexp/syntax/parse.go:1008
					return min > max
//line /usr/local/go/src/regexp/syntax/parse.go:1008
					// _ = "end of CoverTab[63500]"
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				}()
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				// _ = "end of CoverTab[63499]"
//line /usr/local/go/src/regexp/syntax/parse.go:1008
			}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1008
				_go_fuzz_dep_.CoverTab[63501]++

										return nil, &Error{ErrInvalidRepeatSize, before[:len(before)-len(after)]}
//line /usr/local/go/src/regexp/syntax/parse.go:1010
				// _ = "end of CoverTab[63501]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1011
				_go_fuzz_dep_.CoverTab[63502]++
//line /usr/local/go/src/regexp/syntax/parse.go:1011
				// _ = "end of CoverTab[63502]"
//line /usr/local/go/src/regexp/syntax/parse.go:1011
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1011
			// _ = "end of CoverTab[63460]"
//line /usr/local/go/src/regexp/syntax/parse.go:1011
			_go_fuzz_dep_.CoverTab[63461]++
									if after, err = p.repeat(op, min, max, before, after, lastRepeat); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1012
				_go_fuzz_dep_.CoverTab[63503]++
										return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:1013
				// _ = "end of CoverTab[63503]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1014
				_go_fuzz_dep_.CoverTab[63504]++
//line /usr/local/go/src/regexp/syntax/parse.go:1014
				// _ = "end of CoverTab[63504]"
//line /usr/local/go/src/regexp/syntax/parse.go:1014
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1014
			// _ = "end of CoverTab[63461]"
//line /usr/local/go/src/regexp/syntax/parse.go:1014
			_go_fuzz_dep_.CoverTab[63462]++
									repeat = before
									t = after
//line /usr/local/go/src/regexp/syntax/parse.go:1016
			// _ = "end of CoverTab[63462]"
		case '\\':
//line /usr/local/go/src/regexp/syntax/parse.go:1017
			_go_fuzz_dep_.CoverTab[63463]++
									if p.flags&PerlX != 0 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1018
				_go_fuzz_dep_.CoverTab[63505]++
//line /usr/local/go/src/regexp/syntax/parse.go:1018
				return len(t) >= 2
//line /usr/local/go/src/regexp/syntax/parse.go:1018
				// _ = "end of CoverTab[63505]"
//line /usr/local/go/src/regexp/syntax/parse.go:1018
			}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1018
				_go_fuzz_dep_.CoverTab[63506]++
										switch t[1] {
				case 'A':
//line /usr/local/go/src/regexp/syntax/parse.go:1020
					_go_fuzz_dep_.CoverTab[63507]++
											p.op(OpBeginText)
											t = t[2:]
											break BigSwitch
//line /usr/local/go/src/regexp/syntax/parse.go:1023
					// _ = "end of CoverTab[63507]"
				case 'b':
//line /usr/local/go/src/regexp/syntax/parse.go:1024
					_go_fuzz_dep_.CoverTab[63508]++
											p.op(OpWordBoundary)
											t = t[2:]
											break BigSwitch
//line /usr/local/go/src/regexp/syntax/parse.go:1027
					// _ = "end of CoverTab[63508]"
				case 'B':
//line /usr/local/go/src/regexp/syntax/parse.go:1028
					_go_fuzz_dep_.CoverTab[63509]++
											p.op(OpNoWordBoundary)
											t = t[2:]
											break BigSwitch
//line /usr/local/go/src/regexp/syntax/parse.go:1031
					// _ = "end of CoverTab[63509]"
				case 'C':
//line /usr/local/go/src/regexp/syntax/parse.go:1032
					_go_fuzz_dep_.CoverTab[63510]++

											return nil, &Error{ErrInvalidEscape, t[:2]}
//line /usr/local/go/src/regexp/syntax/parse.go:1034
					// _ = "end of CoverTab[63510]"
				case 'Q':
//line /usr/local/go/src/regexp/syntax/parse.go:1035
					_go_fuzz_dep_.CoverTab[63511]++
					// \Q ... \E: the ... is always literals
					var lit string
					lit, t, _ = strings.Cut(t[2:], `\E`)
					for lit != "" {
//line /usr/local/go/src/regexp/syntax/parse.go:1039
						_go_fuzz_dep_.CoverTab[63515]++
												c, rest, err := nextRune(lit)
												if err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1041
							_go_fuzz_dep_.CoverTab[63517]++
													return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:1042
							// _ = "end of CoverTab[63517]"
						} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1043
							_go_fuzz_dep_.CoverTab[63518]++
//line /usr/local/go/src/regexp/syntax/parse.go:1043
							// _ = "end of CoverTab[63518]"
//line /usr/local/go/src/regexp/syntax/parse.go:1043
						}
//line /usr/local/go/src/regexp/syntax/parse.go:1043
						// _ = "end of CoverTab[63515]"
//line /usr/local/go/src/regexp/syntax/parse.go:1043
						_go_fuzz_dep_.CoverTab[63516]++
												p.literal(c)
												lit = rest
//line /usr/local/go/src/regexp/syntax/parse.go:1045
						// _ = "end of CoverTab[63516]"
					}
//line /usr/local/go/src/regexp/syntax/parse.go:1046
					// _ = "end of CoverTab[63511]"
//line /usr/local/go/src/regexp/syntax/parse.go:1046
					_go_fuzz_dep_.CoverTab[63512]++
											break BigSwitch
//line /usr/local/go/src/regexp/syntax/parse.go:1047
					// _ = "end of CoverTab[63512]"
				case 'z':
//line /usr/local/go/src/regexp/syntax/parse.go:1048
					_go_fuzz_dep_.CoverTab[63513]++
											p.op(OpEndText)
											t = t[2:]
											break BigSwitch
//line /usr/local/go/src/regexp/syntax/parse.go:1051
					// _ = "end of CoverTab[63513]"
//line /usr/local/go/src/regexp/syntax/parse.go:1051
				default:
//line /usr/local/go/src/regexp/syntax/parse.go:1051
					_go_fuzz_dep_.CoverTab[63514]++
//line /usr/local/go/src/regexp/syntax/parse.go:1051
					// _ = "end of CoverTab[63514]"
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1052
				// _ = "end of CoverTab[63506]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1053
				_go_fuzz_dep_.CoverTab[63519]++
//line /usr/local/go/src/regexp/syntax/parse.go:1053
				// _ = "end of CoverTab[63519]"
//line /usr/local/go/src/regexp/syntax/parse.go:1053
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1053
			// _ = "end of CoverTab[63463]"
//line /usr/local/go/src/regexp/syntax/parse.go:1053
			_go_fuzz_dep_.CoverTab[63464]++

									re := p.newRegexp(OpCharClass)
									re.Flags = p.flags

//line /usr/local/go/src/regexp/syntax/parse.go:1059
			if len(t) >= 2 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1059
				_go_fuzz_dep_.CoverTab[63520]++
//line /usr/local/go/src/regexp/syntax/parse.go:1059
				return (t[1] == 'p' || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1059
					_go_fuzz_dep_.CoverTab[63521]++
//line /usr/local/go/src/regexp/syntax/parse.go:1059
					return t[1] == 'P'
//line /usr/local/go/src/regexp/syntax/parse.go:1059
					// _ = "end of CoverTab[63521]"
//line /usr/local/go/src/regexp/syntax/parse.go:1059
				}())
//line /usr/local/go/src/regexp/syntax/parse.go:1059
				// _ = "end of CoverTab[63520]"
//line /usr/local/go/src/regexp/syntax/parse.go:1059
			}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1059
				_go_fuzz_dep_.CoverTab[63522]++
										r, rest, err := p.parseUnicodeClass(t, re.Rune0[:0])
										if err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1061
					_go_fuzz_dep_.CoverTab[63524]++
											return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:1062
					// _ = "end of CoverTab[63524]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1063
					_go_fuzz_dep_.CoverTab[63525]++
//line /usr/local/go/src/regexp/syntax/parse.go:1063
					// _ = "end of CoverTab[63525]"
//line /usr/local/go/src/regexp/syntax/parse.go:1063
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1063
				// _ = "end of CoverTab[63522]"
//line /usr/local/go/src/regexp/syntax/parse.go:1063
				_go_fuzz_dep_.CoverTab[63523]++
										if r != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1064
					_go_fuzz_dep_.CoverTab[63526]++
											re.Rune = r
											t = rest
											p.push(re)
											break BigSwitch
//line /usr/local/go/src/regexp/syntax/parse.go:1068
					// _ = "end of CoverTab[63526]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1069
					_go_fuzz_dep_.CoverTab[63527]++
//line /usr/local/go/src/regexp/syntax/parse.go:1069
					// _ = "end of CoverTab[63527]"
//line /usr/local/go/src/regexp/syntax/parse.go:1069
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1069
				// _ = "end of CoverTab[63523]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1070
				_go_fuzz_dep_.CoverTab[63528]++
//line /usr/local/go/src/regexp/syntax/parse.go:1070
				// _ = "end of CoverTab[63528]"
//line /usr/local/go/src/regexp/syntax/parse.go:1070
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1070
			// _ = "end of CoverTab[63464]"
//line /usr/local/go/src/regexp/syntax/parse.go:1070
			_go_fuzz_dep_.CoverTab[63465]++

//line /usr/local/go/src/regexp/syntax/parse.go:1073
			if r, rest := p.parsePerlClassEscape(t, re.Rune0[:0]); r != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1073
				_go_fuzz_dep_.CoverTab[63529]++
										re.Rune = r
										t = rest
										p.push(re)
										break BigSwitch
//line /usr/local/go/src/regexp/syntax/parse.go:1077
				// _ = "end of CoverTab[63529]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1078
				_go_fuzz_dep_.CoverTab[63530]++
//line /usr/local/go/src/regexp/syntax/parse.go:1078
				// _ = "end of CoverTab[63530]"
//line /usr/local/go/src/regexp/syntax/parse.go:1078
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1078
			// _ = "end of CoverTab[63465]"
//line /usr/local/go/src/regexp/syntax/parse.go:1078
			_go_fuzz_dep_.CoverTab[63466]++
									p.reuse(re)

//line /usr/local/go/src/regexp/syntax/parse.go:1082
			if c, t, err = p.parseEscape(t); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1082
				_go_fuzz_dep_.CoverTab[63531]++
										return nil, err
//line /usr/local/go/src/regexp/syntax/parse.go:1083
				// _ = "end of CoverTab[63531]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1084
				_go_fuzz_dep_.CoverTab[63532]++
//line /usr/local/go/src/regexp/syntax/parse.go:1084
				// _ = "end of CoverTab[63532]"
//line /usr/local/go/src/regexp/syntax/parse.go:1084
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1084
			// _ = "end of CoverTab[63466]"
//line /usr/local/go/src/regexp/syntax/parse.go:1084
			_go_fuzz_dep_.CoverTab[63467]++
									p.literal(c)
//line /usr/local/go/src/regexp/syntax/parse.go:1085
			// _ = "end of CoverTab[63467]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1086
		// _ = "end of CoverTab[63439]"
//line /usr/local/go/src/regexp/syntax/parse.go:1086
		_go_fuzz_dep_.CoverTab[63440]++
								lastRepeat = repeat
//line /usr/local/go/src/regexp/syntax/parse.go:1087
		// _ = "end of CoverTab[63440]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1088
	// _ = "end of CoverTab[63425]"
//line /usr/local/go/src/regexp/syntax/parse.go:1088
	_go_fuzz_dep_.CoverTab[63426]++

							p.concat()
							if p.swapVerticalBar() {
//line /usr/local/go/src/regexp/syntax/parse.go:1091
		_go_fuzz_dep_.CoverTab[63533]++

								p.stack = p.stack[:len(p.stack)-1]
//line /usr/local/go/src/regexp/syntax/parse.go:1093
		// _ = "end of CoverTab[63533]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1094
		_go_fuzz_dep_.CoverTab[63534]++
//line /usr/local/go/src/regexp/syntax/parse.go:1094
		// _ = "end of CoverTab[63534]"
//line /usr/local/go/src/regexp/syntax/parse.go:1094
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1094
	// _ = "end of CoverTab[63426]"
//line /usr/local/go/src/regexp/syntax/parse.go:1094
	_go_fuzz_dep_.CoverTab[63427]++
							p.alternate()

							n := len(p.stack)
							if n != 1 {
//line /usr/local/go/src/regexp/syntax/parse.go:1098
		_go_fuzz_dep_.CoverTab[63535]++
								return nil, &Error{ErrMissingParen, s}
//line /usr/local/go/src/regexp/syntax/parse.go:1099
		// _ = "end of CoverTab[63535]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1100
		_go_fuzz_dep_.CoverTab[63536]++
//line /usr/local/go/src/regexp/syntax/parse.go:1100
		// _ = "end of CoverTab[63536]"
//line /usr/local/go/src/regexp/syntax/parse.go:1100
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1100
	// _ = "end of CoverTab[63427]"
//line /usr/local/go/src/regexp/syntax/parse.go:1100
	_go_fuzz_dep_.CoverTab[63428]++
							return p.stack[0], nil
//line /usr/local/go/src/regexp/syntax/parse.go:1101
	// _ = "end of CoverTab[63428]"
}

// parseRepeat parses {min} (max=min) or {min,} (max=-1) or {min,max}.
//line /usr/local/go/src/regexp/syntax/parse.go:1104
// If s is not of that form, it returns ok == false.
//line /usr/local/go/src/regexp/syntax/parse.go:1104
// If s has the right form but the values are too big, it returns min == -1, ok == true.
//line /usr/local/go/src/regexp/syntax/parse.go:1107
func (p *parser) parseRepeat(s string) (min, max int, rest string, ok bool) {
//line /usr/local/go/src/regexp/syntax/parse.go:1107
	_go_fuzz_dep_.CoverTab[63537]++
							if s == "" || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1108
		_go_fuzz_dep_.CoverTab[63543]++
//line /usr/local/go/src/regexp/syntax/parse.go:1108
		return s[0] != '{'
//line /usr/local/go/src/regexp/syntax/parse.go:1108
		// _ = "end of CoverTab[63543]"
//line /usr/local/go/src/regexp/syntax/parse.go:1108
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1108
		_go_fuzz_dep_.CoverTab[63544]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1109
		// _ = "end of CoverTab[63544]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1110
		_go_fuzz_dep_.CoverTab[63545]++
//line /usr/local/go/src/regexp/syntax/parse.go:1110
		// _ = "end of CoverTab[63545]"
//line /usr/local/go/src/regexp/syntax/parse.go:1110
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1110
	// _ = "end of CoverTab[63537]"
//line /usr/local/go/src/regexp/syntax/parse.go:1110
	_go_fuzz_dep_.CoverTab[63538]++
							s = s[1:]
							var ok1 bool
							if min, s, ok1 = p.parseInt(s); !ok1 {
//line /usr/local/go/src/regexp/syntax/parse.go:1113
		_go_fuzz_dep_.CoverTab[63546]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1114
		// _ = "end of CoverTab[63546]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1115
		_go_fuzz_dep_.CoverTab[63547]++
//line /usr/local/go/src/regexp/syntax/parse.go:1115
		// _ = "end of CoverTab[63547]"
//line /usr/local/go/src/regexp/syntax/parse.go:1115
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1115
	// _ = "end of CoverTab[63538]"
//line /usr/local/go/src/regexp/syntax/parse.go:1115
	_go_fuzz_dep_.CoverTab[63539]++
							if s == "" {
//line /usr/local/go/src/regexp/syntax/parse.go:1116
		_go_fuzz_dep_.CoverTab[63548]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1117
		// _ = "end of CoverTab[63548]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1118
		_go_fuzz_dep_.CoverTab[63549]++
//line /usr/local/go/src/regexp/syntax/parse.go:1118
		// _ = "end of CoverTab[63549]"
//line /usr/local/go/src/regexp/syntax/parse.go:1118
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1118
	// _ = "end of CoverTab[63539]"
//line /usr/local/go/src/regexp/syntax/parse.go:1118
	_go_fuzz_dep_.CoverTab[63540]++
							if s[0] != ',' {
//line /usr/local/go/src/regexp/syntax/parse.go:1119
		_go_fuzz_dep_.CoverTab[63550]++
								max = min
//line /usr/local/go/src/regexp/syntax/parse.go:1120
		// _ = "end of CoverTab[63550]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1121
		_go_fuzz_dep_.CoverTab[63551]++
								s = s[1:]
								if s == "" {
//line /usr/local/go/src/regexp/syntax/parse.go:1123
			_go_fuzz_dep_.CoverTab[63553]++
									return
//line /usr/local/go/src/regexp/syntax/parse.go:1124
			// _ = "end of CoverTab[63553]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1125
			_go_fuzz_dep_.CoverTab[63554]++
//line /usr/local/go/src/regexp/syntax/parse.go:1125
			// _ = "end of CoverTab[63554]"
//line /usr/local/go/src/regexp/syntax/parse.go:1125
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1125
		// _ = "end of CoverTab[63551]"
//line /usr/local/go/src/regexp/syntax/parse.go:1125
		_go_fuzz_dep_.CoverTab[63552]++
								if s[0] == '}' {
//line /usr/local/go/src/regexp/syntax/parse.go:1126
			_go_fuzz_dep_.CoverTab[63555]++
									max = -1
//line /usr/local/go/src/regexp/syntax/parse.go:1127
			// _ = "end of CoverTab[63555]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1128
			_go_fuzz_dep_.CoverTab[63556]++
//line /usr/local/go/src/regexp/syntax/parse.go:1128
			if max, s, ok1 = p.parseInt(s); !ok1 {
//line /usr/local/go/src/regexp/syntax/parse.go:1128
				_go_fuzz_dep_.CoverTab[63557]++
										return
//line /usr/local/go/src/regexp/syntax/parse.go:1129
				// _ = "end of CoverTab[63557]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1130
				_go_fuzz_dep_.CoverTab[63558]++
//line /usr/local/go/src/regexp/syntax/parse.go:1130
				if max < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1130
					_go_fuzz_dep_.CoverTab[63559]++

											min = -1
//line /usr/local/go/src/regexp/syntax/parse.go:1132
					// _ = "end of CoverTab[63559]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1133
					_go_fuzz_dep_.CoverTab[63560]++
//line /usr/local/go/src/regexp/syntax/parse.go:1133
					// _ = "end of CoverTab[63560]"
//line /usr/local/go/src/regexp/syntax/parse.go:1133
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1133
				// _ = "end of CoverTab[63558]"
//line /usr/local/go/src/regexp/syntax/parse.go:1133
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1133
			// _ = "end of CoverTab[63556]"
//line /usr/local/go/src/regexp/syntax/parse.go:1133
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1133
		// _ = "end of CoverTab[63552]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1134
	// _ = "end of CoverTab[63540]"
//line /usr/local/go/src/regexp/syntax/parse.go:1134
	_go_fuzz_dep_.CoverTab[63541]++
							if s == "" || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1135
		_go_fuzz_dep_.CoverTab[63561]++
//line /usr/local/go/src/regexp/syntax/parse.go:1135
		return s[0] != '}'
//line /usr/local/go/src/regexp/syntax/parse.go:1135
		// _ = "end of CoverTab[63561]"
//line /usr/local/go/src/regexp/syntax/parse.go:1135
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1135
		_go_fuzz_dep_.CoverTab[63562]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1136
		// _ = "end of CoverTab[63562]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1137
		_go_fuzz_dep_.CoverTab[63563]++
//line /usr/local/go/src/regexp/syntax/parse.go:1137
		// _ = "end of CoverTab[63563]"
//line /usr/local/go/src/regexp/syntax/parse.go:1137
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1137
	// _ = "end of CoverTab[63541]"
//line /usr/local/go/src/regexp/syntax/parse.go:1137
	_go_fuzz_dep_.CoverTab[63542]++
							rest = s[1:]
							ok = true
							return
//line /usr/local/go/src/regexp/syntax/parse.go:1140
	// _ = "end of CoverTab[63542]"
}

// parsePerlFlags parses a Perl flag setting or non-capturing group or both,
//line /usr/local/go/src/regexp/syntax/parse.go:1143
// like (?i) or (?: or (?i:.  It removes the prefix from s and updates the parse state.
//line /usr/local/go/src/regexp/syntax/parse.go:1143
// The caller must have ensured that s begins with "(?".
//line /usr/local/go/src/regexp/syntax/parse.go:1146
func (p *parser) parsePerlFlags(s string) (rest string, err error) {
//line /usr/local/go/src/regexp/syntax/parse.go:1146
	_go_fuzz_dep_.CoverTab[63564]++
							t := s

//line /usr/local/go/src/regexp/syntax/parse.go:1164
	if len(t) > 4 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1164
		_go_fuzz_dep_.CoverTab[63567]++
//line /usr/local/go/src/regexp/syntax/parse.go:1164
		return t[2] == 'P'
//line /usr/local/go/src/regexp/syntax/parse.go:1164
		// _ = "end of CoverTab[63567]"
//line /usr/local/go/src/regexp/syntax/parse.go:1164
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1164
		_go_fuzz_dep_.CoverTab[63568]++
//line /usr/local/go/src/regexp/syntax/parse.go:1164
		return t[3] == '<'
//line /usr/local/go/src/regexp/syntax/parse.go:1164
		// _ = "end of CoverTab[63568]"
//line /usr/local/go/src/regexp/syntax/parse.go:1164
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1164
		_go_fuzz_dep_.CoverTab[63569]++

								end := strings.IndexRune(t, '>')
								if end < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1167
			_go_fuzz_dep_.CoverTab[63573]++
									if err = checkUTF8(t); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1168
				_go_fuzz_dep_.CoverTab[63575]++
										return "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1169
				// _ = "end of CoverTab[63575]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1170
				_go_fuzz_dep_.CoverTab[63576]++
//line /usr/local/go/src/regexp/syntax/parse.go:1170
				// _ = "end of CoverTab[63576]"
//line /usr/local/go/src/regexp/syntax/parse.go:1170
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1170
			// _ = "end of CoverTab[63573]"
//line /usr/local/go/src/regexp/syntax/parse.go:1170
			_go_fuzz_dep_.CoverTab[63574]++
									return "", &Error{ErrInvalidNamedCapture, s}
//line /usr/local/go/src/regexp/syntax/parse.go:1171
			// _ = "end of CoverTab[63574]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1172
			_go_fuzz_dep_.CoverTab[63577]++
//line /usr/local/go/src/regexp/syntax/parse.go:1172
			// _ = "end of CoverTab[63577]"
//line /usr/local/go/src/regexp/syntax/parse.go:1172
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1172
		// _ = "end of CoverTab[63569]"
//line /usr/local/go/src/regexp/syntax/parse.go:1172
		_go_fuzz_dep_.CoverTab[63570]++

								capture := t[:end+1]
								name := t[4:end]
								if err = checkUTF8(name); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1176
			_go_fuzz_dep_.CoverTab[63578]++
									return "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1177
			// _ = "end of CoverTab[63578]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1178
			_go_fuzz_dep_.CoverTab[63579]++
//line /usr/local/go/src/regexp/syntax/parse.go:1178
			// _ = "end of CoverTab[63579]"
//line /usr/local/go/src/regexp/syntax/parse.go:1178
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1178
		// _ = "end of CoverTab[63570]"
//line /usr/local/go/src/regexp/syntax/parse.go:1178
		_go_fuzz_dep_.CoverTab[63571]++
								if !isValidCaptureName(name) {
//line /usr/local/go/src/regexp/syntax/parse.go:1179
			_go_fuzz_dep_.CoverTab[63580]++
									return "", &Error{ErrInvalidNamedCapture, capture}
//line /usr/local/go/src/regexp/syntax/parse.go:1180
			// _ = "end of CoverTab[63580]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1181
			_go_fuzz_dep_.CoverTab[63581]++
//line /usr/local/go/src/regexp/syntax/parse.go:1181
			// _ = "end of CoverTab[63581]"
//line /usr/local/go/src/regexp/syntax/parse.go:1181
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1181
		// _ = "end of CoverTab[63571]"
//line /usr/local/go/src/regexp/syntax/parse.go:1181
		_go_fuzz_dep_.CoverTab[63572]++

//line /usr/local/go/src/regexp/syntax/parse.go:1184
		p.numCap++
								re := p.op(opLeftParen)
								re.Cap = p.numCap
								re.Name = name
								return t[end+1:], nil
//line /usr/local/go/src/regexp/syntax/parse.go:1188
		// _ = "end of CoverTab[63572]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1189
		_go_fuzz_dep_.CoverTab[63582]++
//line /usr/local/go/src/regexp/syntax/parse.go:1189
		// _ = "end of CoverTab[63582]"
//line /usr/local/go/src/regexp/syntax/parse.go:1189
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1189
	// _ = "end of CoverTab[63564]"
//line /usr/local/go/src/regexp/syntax/parse.go:1189
	_go_fuzz_dep_.CoverTab[63565]++

	// Non-capturing group. Might also twiddle Perl flags.
	var c rune
	t = t[2:]
	flags := p.flags
	sign := +1
	sawFlag := false
Loop:
	for t != "" {
//line /usr/local/go/src/regexp/syntax/parse.go:1198
		_go_fuzz_dep_.CoverTab[63583]++
								if c, t, err = nextRune(t); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1199
			_go_fuzz_dep_.CoverTab[63585]++
									return "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1200
			// _ = "end of CoverTab[63585]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1201
			_go_fuzz_dep_.CoverTab[63586]++
//line /usr/local/go/src/regexp/syntax/parse.go:1201
			// _ = "end of CoverTab[63586]"
//line /usr/local/go/src/regexp/syntax/parse.go:1201
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1201
		// _ = "end of CoverTab[63583]"
//line /usr/local/go/src/regexp/syntax/parse.go:1201
		_go_fuzz_dep_.CoverTab[63584]++
								switch c {
		default:
//line /usr/local/go/src/regexp/syntax/parse.go:1203
			_go_fuzz_dep_.CoverTab[63587]++
									break Loop
//line /usr/local/go/src/regexp/syntax/parse.go:1204
			// _ = "end of CoverTab[63587]"

//line /usr/local/go/src/regexp/syntax/parse.go:1207
		case 'i':
//line /usr/local/go/src/regexp/syntax/parse.go:1207
			_go_fuzz_dep_.CoverTab[63588]++
									flags |= FoldCase
									sawFlag = true
//line /usr/local/go/src/regexp/syntax/parse.go:1209
			// _ = "end of CoverTab[63588]"
		case 'm':
//line /usr/local/go/src/regexp/syntax/parse.go:1210
			_go_fuzz_dep_.CoverTab[63589]++
									flags &^= OneLine
									sawFlag = true
//line /usr/local/go/src/regexp/syntax/parse.go:1212
			// _ = "end of CoverTab[63589]"
		case 's':
//line /usr/local/go/src/regexp/syntax/parse.go:1213
			_go_fuzz_dep_.CoverTab[63590]++
									flags |= DotNL
									sawFlag = true
//line /usr/local/go/src/regexp/syntax/parse.go:1215
			// _ = "end of CoverTab[63590]"
		case 'U':
//line /usr/local/go/src/regexp/syntax/parse.go:1216
			_go_fuzz_dep_.CoverTab[63591]++
									flags |= NonGreedy
									sawFlag = true
//line /usr/local/go/src/regexp/syntax/parse.go:1218
			// _ = "end of CoverTab[63591]"

//line /usr/local/go/src/regexp/syntax/parse.go:1221
		case '-':
//line /usr/local/go/src/regexp/syntax/parse.go:1221
			_go_fuzz_dep_.CoverTab[63592]++
									if sign < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1222
				_go_fuzz_dep_.CoverTab[63597]++
										break Loop
//line /usr/local/go/src/regexp/syntax/parse.go:1223
				// _ = "end of CoverTab[63597]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1224
				_go_fuzz_dep_.CoverTab[63598]++
//line /usr/local/go/src/regexp/syntax/parse.go:1224
				// _ = "end of CoverTab[63598]"
//line /usr/local/go/src/regexp/syntax/parse.go:1224
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1224
			// _ = "end of CoverTab[63592]"
//line /usr/local/go/src/regexp/syntax/parse.go:1224
			_go_fuzz_dep_.CoverTab[63593]++
									sign = -1

//line /usr/local/go/src/regexp/syntax/parse.go:1228
			flags = ^flags
									sawFlag = false
//line /usr/local/go/src/regexp/syntax/parse.go:1229
			// _ = "end of CoverTab[63593]"

//line /usr/local/go/src/regexp/syntax/parse.go:1232
		case ':', ')':
//line /usr/local/go/src/regexp/syntax/parse.go:1232
			_go_fuzz_dep_.CoverTab[63594]++
									if sign < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1233
				_go_fuzz_dep_.CoverTab[63599]++
										if !sawFlag {
//line /usr/local/go/src/regexp/syntax/parse.go:1234
					_go_fuzz_dep_.CoverTab[63601]++
											break Loop
//line /usr/local/go/src/regexp/syntax/parse.go:1235
					// _ = "end of CoverTab[63601]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1236
					_go_fuzz_dep_.CoverTab[63602]++
//line /usr/local/go/src/regexp/syntax/parse.go:1236
					// _ = "end of CoverTab[63602]"
//line /usr/local/go/src/regexp/syntax/parse.go:1236
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1236
				// _ = "end of CoverTab[63599]"
//line /usr/local/go/src/regexp/syntax/parse.go:1236
				_go_fuzz_dep_.CoverTab[63600]++
										flags = ^flags
//line /usr/local/go/src/regexp/syntax/parse.go:1237
				// _ = "end of CoverTab[63600]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1238
				_go_fuzz_dep_.CoverTab[63603]++
//line /usr/local/go/src/regexp/syntax/parse.go:1238
				// _ = "end of CoverTab[63603]"
//line /usr/local/go/src/regexp/syntax/parse.go:1238
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1238
			// _ = "end of CoverTab[63594]"
//line /usr/local/go/src/regexp/syntax/parse.go:1238
			_go_fuzz_dep_.CoverTab[63595]++
									if c == ':' {
//line /usr/local/go/src/regexp/syntax/parse.go:1239
				_go_fuzz_dep_.CoverTab[63604]++

										p.op(opLeftParen)
//line /usr/local/go/src/regexp/syntax/parse.go:1241
				// _ = "end of CoverTab[63604]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1242
				_go_fuzz_dep_.CoverTab[63605]++
//line /usr/local/go/src/regexp/syntax/parse.go:1242
				// _ = "end of CoverTab[63605]"
//line /usr/local/go/src/regexp/syntax/parse.go:1242
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1242
			// _ = "end of CoverTab[63595]"
//line /usr/local/go/src/regexp/syntax/parse.go:1242
			_go_fuzz_dep_.CoverTab[63596]++
									p.flags = flags
									return t, nil
//line /usr/local/go/src/regexp/syntax/parse.go:1244
			// _ = "end of CoverTab[63596]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1245
		// _ = "end of CoverTab[63584]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1246
	// _ = "end of CoverTab[63565]"
//line /usr/local/go/src/regexp/syntax/parse.go:1246
	_go_fuzz_dep_.CoverTab[63566]++

							return "", &Error{ErrInvalidPerlOp, s[:len(s)-len(t)]}
//line /usr/local/go/src/regexp/syntax/parse.go:1248
	// _ = "end of CoverTab[63566]"
}

// isValidCaptureName reports whether name
//line /usr/local/go/src/regexp/syntax/parse.go:1251
// is a valid capture name: [A-Za-z0-9_]+.
//line /usr/local/go/src/regexp/syntax/parse.go:1251
// PCRE limits names to 32 bytes.
//line /usr/local/go/src/regexp/syntax/parse.go:1251
// Python rejects names starting with digits.
//line /usr/local/go/src/regexp/syntax/parse.go:1251
// We don't enforce either of those.
//line /usr/local/go/src/regexp/syntax/parse.go:1256
func isValidCaptureName(name string) bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1256
	_go_fuzz_dep_.CoverTab[63606]++
							if name == "" {
//line /usr/local/go/src/regexp/syntax/parse.go:1257
		_go_fuzz_dep_.CoverTab[63609]++
								return false
//line /usr/local/go/src/regexp/syntax/parse.go:1258
		// _ = "end of CoverTab[63609]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1259
		_go_fuzz_dep_.CoverTab[63610]++
//line /usr/local/go/src/regexp/syntax/parse.go:1259
		// _ = "end of CoverTab[63610]"
//line /usr/local/go/src/regexp/syntax/parse.go:1259
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1259
	// _ = "end of CoverTab[63606]"
//line /usr/local/go/src/regexp/syntax/parse.go:1259
	_go_fuzz_dep_.CoverTab[63607]++
							for _, c := range name {
//line /usr/local/go/src/regexp/syntax/parse.go:1260
		_go_fuzz_dep_.CoverTab[63611]++
								if c != '_' && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1261
			_go_fuzz_dep_.CoverTab[63612]++
//line /usr/local/go/src/regexp/syntax/parse.go:1261
			return !isalnum(c)
//line /usr/local/go/src/regexp/syntax/parse.go:1261
			// _ = "end of CoverTab[63612]"
//line /usr/local/go/src/regexp/syntax/parse.go:1261
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1261
			_go_fuzz_dep_.CoverTab[63613]++
									return false
//line /usr/local/go/src/regexp/syntax/parse.go:1262
			// _ = "end of CoverTab[63613]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1263
			_go_fuzz_dep_.CoverTab[63614]++
//line /usr/local/go/src/regexp/syntax/parse.go:1263
			// _ = "end of CoverTab[63614]"
//line /usr/local/go/src/regexp/syntax/parse.go:1263
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1263
		// _ = "end of CoverTab[63611]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1264
	// _ = "end of CoverTab[63607]"
//line /usr/local/go/src/regexp/syntax/parse.go:1264
	_go_fuzz_dep_.CoverTab[63608]++
							return true
//line /usr/local/go/src/regexp/syntax/parse.go:1265
	// _ = "end of CoverTab[63608]"
}

// parseInt parses a decimal integer.
func (p *parser) parseInt(s string) (n int, rest string, ok bool) {
//line /usr/local/go/src/regexp/syntax/parse.go:1269
	_go_fuzz_dep_.CoverTab[63615]++
							if s == "" || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1270
		_go_fuzz_dep_.CoverTab[63620]++
//line /usr/local/go/src/regexp/syntax/parse.go:1270
		return s[0] < '0'
//line /usr/local/go/src/regexp/syntax/parse.go:1270
		// _ = "end of CoverTab[63620]"
//line /usr/local/go/src/regexp/syntax/parse.go:1270
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1270
		_go_fuzz_dep_.CoverTab[63621]++
//line /usr/local/go/src/regexp/syntax/parse.go:1270
		return '9' < s[0]
//line /usr/local/go/src/regexp/syntax/parse.go:1270
		// _ = "end of CoverTab[63621]"
//line /usr/local/go/src/regexp/syntax/parse.go:1270
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1270
		_go_fuzz_dep_.CoverTab[63622]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1271
		// _ = "end of CoverTab[63622]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1272
		_go_fuzz_dep_.CoverTab[63623]++
//line /usr/local/go/src/regexp/syntax/parse.go:1272
		// _ = "end of CoverTab[63623]"
//line /usr/local/go/src/regexp/syntax/parse.go:1272
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1272
	// _ = "end of CoverTab[63615]"
//line /usr/local/go/src/regexp/syntax/parse.go:1272
	_go_fuzz_dep_.CoverTab[63616]++

							if len(s) >= 2 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		_go_fuzz_dep_.CoverTab[63624]++
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		return s[0] == '0'
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		// _ = "end of CoverTab[63624]"
//line /usr/local/go/src/regexp/syntax/parse.go:1274
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		_go_fuzz_dep_.CoverTab[63625]++
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		return '0' <= s[1]
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		// _ = "end of CoverTab[63625]"
//line /usr/local/go/src/regexp/syntax/parse.go:1274
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		_go_fuzz_dep_.CoverTab[63626]++
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		return s[1] <= '9'
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		// _ = "end of CoverTab[63626]"
//line /usr/local/go/src/regexp/syntax/parse.go:1274
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1274
		_go_fuzz_dep_.CoverTab[63627]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1275
		// _ = "end of CoverTab[63627]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1276
		_go_fuzz_dep_.CoverTab[63628]++
//line /usr/local/go/src/regexp/syntax/parse.go:1276
		// _ = "end of CoverTab[63628]"
//line /usr/local/go/src/regexp/syntax/parse.go:1276
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1276
	// _ = "end of CoverTab[63616]"
//line /usr/local/go/src/regexp/syntax/parse.go:1276
	_go_fuzz_dep_.CoverTab[63617]++
							t := s
							for s != "" && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1278
		_go_fuzz_dep_.CoverTab[63629]++
//line /usr/local/go/src/regexp/syntax/parse.go:1278
		return '0' <= s[0]
//line /usr/local/go/src/regexp/syntax/parse.go:1278
		// _ = "end of CoverTab[63629]"
//line /usr/local/go/src/regexp/syntax/parse.go:1278
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1278
		_go_fuzz_dep_.CoverTab[63630]++
//line /usr/local/go/src/regexp/syntax/parse.go:1278
		return s[0] <= '9'
//line /usr/local/go/src/regexp/syntax/parse.go:1278
		// _ = "end of CoverTab[63630]"
//line /usr/local/go/src/regexp/syntax/parse.go:1278
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1278
		_go_fuzz_dep_.CoverTab[63631]++
								s = s[1:]
//line /usr/local/go/src/regexp/syntax/parse.go:1279
		// _ = "end of CoverTab[63631]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1280
	// _ = "end of CoverTab[63617]"
//line /usr/local/go/src/regexp/syntax/parse.go:1280
	_go_fuzz_dep_.CoverTab[63618]++
							rest = s
							ok = true

							t = t[:len(t)-len(s)]
							for i := 0; i < len(t); i++ {
//line /usr/local/go/src/regexp/syntax/parse.go:1285
		_go_fuzz_dep_.CoverTab[63632]++

								if n >= 1e8 {
//line /usr/local/go/src/regexp/syntax/parse.go:1287
			_go_fuzz_dep_.CoverTab[63634]++
									n = -1
									break
//line /usr/local/go/src/regexp/syntax/parse.go:1289
			// _ = "end of CoverTab[63634]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1290
			_go_fuzz_dep_.CoverTab[63635]++
//line /usr/local/go/src/regexp/syntax/parse.go:1290
			// _ = "end of CoverTab[63635]"
//line /usr/local/go/src/regexp/syntax/parse.go:1290
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1290
		// _ = "end of CoverTab[63632]"
//line /usr/local/go/src/regexp/syntax/parse.go:1290
		_go_fuzz_dep_.CoverTab[63633]++
								n = n*10 + int(t[i]) - '0'
//line /usr/local/go/src/regexp/syntax/parse.go:1291
		// _ = "end of CoverTab[63633]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1292
	// _ = "end of CoverTab[63618]"
//line /usr/local/go/src/regexp/syntax/parse.go:1292
	_go_fuzz_dep_.CoverTab[63619]++
							return
//line /usr/local/go/src/regexp/syntax/parse.go:1293
	// _ = "end of CoverTab[63619]"
}

// can this be represented as a character class?
//line /usr/local/go/src/regexp/syntax/parse.go:1296
// single-rune literal string, char class, ., and .|\n.
//line /usr/local/go/src/regexp/syntax/parse.go:1298
func isCharClass(re *Regexp) bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1298
	_go_fuzz_dep_.CoverTab[63636]++
							return re.Op == OpLiteral && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1299
		_go_fuzz_dep_.CoverTab[63637]++
//line /usr/local/go/src/regexp/syntax/parse.go:1299
		return len(re.Rune) == 1
//line /usr/local/go/src/regexp/syntax/parse.go:1299
		// _ = "end of CoverTab[63637]"
//line /usr/local/go/src/regexp/syntax/parse.go:1299
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1299
		_go_fuzz_dep_.CoverTab[63638]++
//line /usr/local/go/src/regexp/syntax/parse.go:1299
		return re.Op == OpCharClass
								// _ = "end of CoverTab[63638]"
//line /usr/local/go/src/regexp/syntax/parse.go:1300
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1300
		_go_fuzz_dep_.CoverTab[63639]++
//line /usr/local/go/src/regexp/syntax/parse.go:1300
		return re.Op == OpAnyCharNotNL
								// _ = "end of CoverTab[63639]"
//line /usr/local/go/src/regexp/syntax/parse.go:1301
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1301
		_go_fuzz_dep_.CoverTab[63640]++
//line /usr/local/go/src/regexp/syntax/parse.go:1301
		return re.Op == OpAnyChar
								// _ = "end of CoverTab[63640]"
//line /usr/local/go/src/regexp/syntax/parse.go:1302
	}()
//line /usr/local/go/src/regexp/syntax/parse.go:1302
	// _ = "end of CoverTab[63636]"
}

// does re match r?
func matchRune(re *Regexp, r rune) bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1306
	_go_fuzz_dep_.CoverTab[63641]++
							switch re.Op {
	case OpLiteral:
//line /usr/local/go/src/regexp/syntax/parse.go:1308
		_go_fuzz_dep_.CoverTab[63643]++
								return len(re.Rune) == 1 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1309
			_go_fuzz_dep_.CoverTab[63649]++
//line /usr/local/go/src/regexp/syntax/parse.go:1309
			return re.Rune[0] == r
//line /usr/local/go/src/regexp/syntax/parse.go:1309
			// _ = "end of CoverTab[63649]"
//line /usr/local/go/src/regexp/syntax/parse.go:1309
		}()
//line /usr/local/go/src/regexp/syntax/parse.go:1309
		// _ = "end of CoverTab[63643]"
	case OpCharClass:
//line /usr/local/go/src/regexp/syntax/parse.go:1310
		_go_fuzz_dep_.CoverTab[63644]++
								for i := 0; i < len(re.Rune); i += 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:1311
			_go_fuzz_dep_.CoverTab[63650]++
									if re.Rune[i] <= r && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1312
				_go_fuzz_dep_.CoverTab[63651]++
//line /usr/local/go/src/regexp/syntax/parse.go:1312
				return r <= re.Rune[i+1]
//line /usr/local/go/src/regexp/syntax/parse.go:1312
				// _ = "end of CoverTab[63651]"
//line /usr/local/go/src/regexp/syntax/parse.go:1312
			}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1312
				_go_fuzz_dep_.CoverTab[63652]++
										return true
//line /usr/local/go/src/regexp/syntax/parse.go:1313
				// _ = "end of CoverTab[63652]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1314
				_go_fuzz_dep_.CoverTab[63653]++
//line /usr/local/go/src/regexp/syntax/parse.go:1314
				// _ = "end of CoverTab[63653]"
//line /usr/local/go/src/regexp/syntax/parse.go:1314
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1314
			// _ = "end of CoverTab[63650]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1315
		// _ = "end of CoverTab[63644]"
//line /usr/local/go/src/regexp/syntax/parse.go:1315
		_go_fuzz_dep_.CoverTab[63645]++
								return false
//line /usr/local/go/src/regexp/syntax/parse.go:1316
		// _ = "end of CoverTab[63645]"
	case OpAnyCharNotNL:
//line /usr/local/go/src/regexp/syntax/parse.go:1317
		_go_fuzz_dep_.CoverTab[63646]++
								return r != '\n'
//line /usr/local/go/src/regexp/syntax/parse.go:1318
		// _ = "end of CoverTab[63646]"
	case OpAnyChar:
//line /usr/local/go/src/regexp/syntax/parse.go:1319
		_go_fuzz_dep_.CoverTab[63647]++
								return true
//line /usr/local/go/src/regexp/syntax/parse.go:1320
		// _ = "end of CoverTab[63647]"
//line /usr/local/go/src/regexp/syntax/parse.go:1320
	default:
//line /usr/local/go/src/regexp/syntax/parse.go:1320
		_go_fuzz_dep_.CoverTab[63648]++
//line /usr/local/go/src/regexp/syntax/parse.go:1320
		// _ = "end of CoverTab[63648]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1321
	// _ = "end of CoverTab[63641]"
//line /usr/local/go/src/regexp/syntax/parse.go:1321
	_go_fuzz_dep_.CoverTab[63642]++
							return false
//line /usr/local/go/src/regexp/syntax/parse.go:1322
	// _ = "end of CoverTab[63642]"
}

// parseVerticalBar handles a | in the input.
func (p *parser) parseVerticalBar() error {
//line /usr/local/go/src/regexp/syntax/parse.go:1326
	_go_fuzz_dep_.CoverTab[63654]++
							p.concat()

//line /usr/local/go/src/regexp/syntax/parse.go:1333
	if !p.swapVerticalBar() {
//line /usr/local/go/src/regexp/syntax/parse.go:1333
		_go_fuzz_dep_.CoverTab[63656]++
								p.op(opVerticalBar)
//line /usr/local/go/src/regexp/syntax/parse.go:1334
		// _ = "end of CoverTab[63656]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1335
		_go_fuzz_dep_.CoverTab[63657]++
//line /usr/local/go/src/regexp/syntax/parse.go:1335
		// _ = "end of CoverTab[63657]"
//line /usr/local/go/src/regexp/syntax/parse.go:1335
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1335
	// _ = "end of CoverTab[63654]"
//line /usr/local/go/src/regexp/syntax/parse.go:1335
	_go_fuzz_dep_.CoverTab[63655]++

							return nil
//line /usr/local/go/src/regexp/syntax/parse.go:1337
	// _ = "end of CoverTab[63655]"
}

// mergeCharClass makes dst = dst|src.
//line /usr/local/go/src/regexp/syntax/parse.go:1340
// The caller must ensure that dst.Op >= src.Op,
//line /usr/local/go/src/regexp/syntax/parse.go:1340
// to reduce the amount of copying.
//line /usr/local/go/src/regexp/syntax/parse.go:1343
func mergeCharClass(dst, src *Regexp) {
//line /usr/local/go/src/regexp/syntax/parse.go:1343
	_go_fuzz_dep_.CoverTab[63658]++
							switch dst.Op {
	case OpAnyChar:
//line /usr/local/go/src/regexp/syntax/parse.go:1345
		_go_fuzz_dep_.CoverTab[63659]++
//line /usr/local/go/src/regexp/syntax/parse.go:1345
		// _ = "end of CoverTab[63659]"

	case OpAnyCharNotNL:
//line /usr/local/go/src/regexp/syntax/parse.go:1347
		_go_fuzz_dep_.CoverTab[63660]++

								if matchRune(src, '\n') {
//line /usr/local/go/src/regexp/syntax/parse.go:1349
			_go_fuzz_dep_.CoverTab[63665]++
									dst.Op = OpAnyChar
//line /usr/local/go/src/regexp/syntax/parse.go:1350
			// _ = "end of CoverTab[63665]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1351
			_go_fuzz_dep_.CoverTab[63666]++
//line /usr/local/go/src/regexp/syntax/parse.go:1351
			// _ = "end of CoverTab[63666]"
//line /usr/local/go/src/regexp/syntax/parse.go:1351
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1351
		// _ = "end of CoverTab[63660]"
	case OpCharClass:
//line /usr/local/go/src/regexp/syntax/parse.go:1352
		_go_fuzz_dep_.CoverTab[63661]++

								if src.Op == OpLiteral {
//line /usr/local/go/src/regexp/syntax/parse.go:1354
			_go_fuzz_dep_.CoverTab[63667]++
									dst.Rune = appendLiteral(dst.Rune, src.Rune[0], src.Flags)
//line /usr/local/go/src/regexp/syntax/parse.go:1355
			// _ = "end of CoverTab[63667]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1356
			_go_fuzz_dep_.CoverTab[63668]++
									dst.Rune = appendClass(dst.Rune, src.Rune)
//line /usr/local/go/src/regexp/syntax/parse.go:1357
			// _ = "end of CoverTab[63668]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1358
		// _ = "end of CoverTab[63661]"
	case OpLiteral:
//line /usr/local/go/src/regexp/syntax/parse.go:1359
		_go_fuzz_dep_.CoverTab[63662]++

								if src.Rune[0] == dst.Rune[0] && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1361
			_go_fuzz_dep_.CoverTab[63669]++
//line /usr/local/go/src/regexp/syntax/parse.go:1361
			return src.Flags == dst.Flags
//line /usr/local/go/src/regexp/syntax/parse.go:1361
			// _ = "end of CoverTab[63669]"
//line /usr/local/go/src/regexp/syntax/parse.go:1361
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1361
			_go_fuzz_dep_.CoverTab[63670]++
									break
//line /usr/local/go/src/regexp/syntax/parse.go:1362
			// _ = "end of CoverTab[63670]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1363
			_go_fuzz_dep_.CoverTab[63671]++
//line /usr/local/go/src/regexp/syntax/parse.go:1363
			// _ = "end of CoverTab[63671]"
//line /usr/local/go/src/regexp/syntax/parse.go:1363
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1363
		// _ = "end of CoverTab[63662]"
//line /usr/local/go/src/regexp/syntax/parse.go:1363
		_go_fuzz_dep_.CoverTab[63663]++
								dst.Op = OpCharClass
								dst.Rune = appendLiteral(dst.Rune[:0], dst.Rune[0], dst.Flags)
								dst.Rune = appendLiteral(dst.Rune, src.Rune[0], src.Flags)
//line /usr/local/go/src/regexp/syntax/parse.go:1366
		// _ = "end of CoverTab[63663]"
//line /usr/local/go/src/regexp/syntax/parse.go:1366
	default:
//line /usr/local/go/src/regexp/syntax/parse.go:1366
		_go_fuzz_dep_.CoverTab[63664]++
//line /usr/local/go/src/regexp/syntax/parse.go:1366
		// _ = "end of CoverTab[63664]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1367
	// _ = "end of CoverTab[63658]"
}

// If the top of the stack is an element followed by an opVerticalBar
//line /usr/local/go/src/regexp/syntax/parse.go:1370
// swapVerticalBar swaps the two and returns true.
//line /usr/local/go/src/regexp/syntax/parse.go:1370
// Otherwise it returns false.
//line /usr/local/go/src/regexp/syntax/parse.go:1373
func (p *parser) swapVerticalBar() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1373
	_go_fuzz_dep_.CoverTab[63672]++

//line /usr/local/go/src/regexp/syntax/parse.go:1376
	n := len(p.stack)
	if n >= 3 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		_go_fuzz_dep_.CoverTab[63675]++
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		return p.stack[n-2].Op == opVerticalBar
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		// _ = "end of CoverTab[63675]"
//line /usr/local/go/src/regexp/syntax/parse.go:1377
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		_go_fuzz_dep_.CoverTab[63676]++
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		return isCharClass(p.stack[n-1])
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		// _ = "end of CoverTab[63676]"
//line /usr/local/go/src/regexp/syntax/parse.go:1377
	}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		_go_fuzz_dep_.CoverTab[63677]++
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		return isCharClass(p.stack[n-3])
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		// _ = "end of CoverTab[63677]"
//line /usr/local/go/src/regexp/syntax/parse.go:1377
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1377
		_go_fuzz_dep_.CoverTab[63678]++
								re1 := p.stack[n-1]
								re3 := p.stack[n-3]

								if re1.Op > re3.Op {
//line /usr/local/go/src/regexp/syntax/parse.go:1381
			_go_fuzz_dep_.CoverTab[63680]++
									re1, re3 = re3, re1
									p.stack[n-3] = re3
//line /usr/local/go/src/regexp/syntax/parse.go:1383
			// _ = "end of CoverTab[63680]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1384
			_go_fuzz_dep_.CoverTab[63681]++
//line /usr/local/go/src/regexp/syntax/parse.go:1384
			// _ = "end of CoverTab[63681]"
//line /usr/local/go/src/regexp/syntax/parse.go:1384
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1384
		// _ = "end of CoverTab[63678]"
//line /usr/local/go/src/regexp/syntax/parse.go:1384
		_go_fuzz_dep_.CoverTab[63679]++
								mergeCharClass(re3, re1)
								p.reuse(re1)
								p.stack = p.stack[:n-1]
								return true
//line /usr/local/go/src/regexp/syntax/parse.go:1388
		// _ = "end of CoverTab[63679]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1389
		_go_fuzz_dep_.CoverTab[63682]++
//line /usr/local/go/src/regexp/syntax/parse.go:1389
		// _ = "end of CoverTab[63682]"
//line /usr/local/go/src/regexp/syntax/parse.go:1389
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1389
	// _ = "end of CoverTab[63672]"
//line /usr/local/go/src/regexp/syntax/parse.go:1389
	_go_fuzz_dep_.CoverTab[63673]++

							if n >= 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:1391
		_go_fuzz_dep_.CoverTab[63683]++
								re1 := p.stack[n-1]
								re2 := p.stack[n-2]
								if re2.Op == opVerticalBar {
//line /usr/local/go/src/regexp/syntax/parse.go:1394
			_go_fuzz_dep_.CoverTab[63684]++
									if n >= 3 {
//line /usr/local/go/src/regexp/syntax/parse.go:1395
				_go_fuzz_dep_.CoverTab[63686]++

//line /usr/local/go/src/regexp/syntax/parse.go:1398
				cleanAlt(p.stack[n-3])
//line /usr/local/go/src/regexp/syntax/parse.go:1398
				// _ = "end of CoverTab[63686]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1399
				_go_fuzz_dep_.CoverTab[63687]++
//line /usr/local/go/src/regexp/syntax/parse.go:1399
				// _ = "end of CoverTab[63687]"
//line /usr/local/go/src/regexp/syntax/parse.go:1399
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1399
			// _ = "end of CoverTab[63684]"
//line /usr/local/go/src/regexp/syntax/parse.go:1399
			_go_fuzz_dep_.CoverTab[63685]++
									p.stack[n-2] = re1
									p.stack[n-1] = re2
									return true
//line /usr/local/go/src/regexp/syntax/parse.go:1402
			// _ = "end of CoverTab[63685]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1403
			_go_fuzz_dep_.CoverTab[63688]++
//line /usr/local/go/src/regexp/syntax/parse.go:1403
			// _ = "end of CoverTab[63688]"
//line /usr/local/go/src/regexp/syntax/parse.go:1403
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1403
		// _ = "end of CoverTab[63683]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1404
		_go_fuzz_dep_.CoverTab[63689]++
//line /usr/local/go/src/regexp/syntax/parse.go:1404
		// _ = "end of CoverTab[63689]"
//line /usr/local/go/src/regexp/syntax/parse.go:1404
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1404
	// _ = "end of CoverTab[63673]"
//line /usr/local/go/src/regexp/syntax/parse.go:1404
	_go_fuzz_dep_.CoverTab[63674]++
							return false
//line /usr/local/go/src/regexp/syntax/parse.go:1405
	// _ = "end of CoverTab[63674]"
}

// parseRightParen handles a ) in the input.
func (p *parser) parseRightParen() error {
//line /usr/local/go/src/regexp/syntax/parse.go:1409
	_go_fuzz_dep_.CoverTab[63690]++
							p.concat()
							if p.swapVerticalBar() {
//line /usr/local/go/src/regexp/syntax/parse.go:1411
		_go_fuzz_dep_.CoverTab[63695]++

								p.stack = p.stack[:len(p.stack)-1]
//line /usr/local/go/src/regexp/syntax/parse.go:1413
		// _ = "end of CoverTab[63695]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1414
		_go_fuzz_dep_.CoverTab[63696]++
//line /usr/local/go/src/regexp/syntax/parse.go:1414
		// _ = "end of CoverTab[63696]"
//line /usr/local/go/src/regexp/syntax/parse.go:1414
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1414
	// _ = "end of CoverTab[63690]"
//line /usr/local/go/src/regexp/syntax/parse.go:1414
	_go_fuzz_dep_.CoverTab[63691]++
							p.alternate()

							n := len(p.stack)
							if n < 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:1418
		_go_fuzz_dep_.CoverTab[63697]++
								return &Error{ErrUnexpectedParen, p.wholeRegexp}
//line /usr/local/go/src/regexp/syntax/parse.go:1419
		// _ = "end of CoverTab[63697]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1420
		_go_fuzz_dep_.CoverTab[63698]++
//line /usr/local/go/src/regexp/syntax/parse.go:1420
		// _ = "end of CoverTab[63698]"
//line /usr/local/go/src/regexp/syntax/parse.go:1420
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1420
	// _ = "end of CoverTab[63691]"
//line /usr/local/go/src/regexp/syntax/parse.go:1420
	_go_fuzz_dep_.CoverTab[63692]++
							re1 := p.stack[n-1]
							re2 := p.stack[n-2]
							p.stack = p.stack[:n-2]
							if re2.Op != opLeftParen {
//line /usr/local/go/src/regexp/syntax/parse.go:1424
		_go_fuzz_dep_.CoverTab[63699]++
								return &Error{ErrUnexpectedParen, p.wholeRegexp}
//line /usr/local/go/src/regexp/syntax/parse.go:1425
		// _ = "end of CoverTab[63699]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1426
		_go_fuzz_dep_.CoverTab[63700]++
//line /usr/local/go/src/regexp/syntax/parse.go:1426
		// _ = "end of CoverTab[63700]"
//line /usr/local/go/src/regexp/syntax/parse.go:1426
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1426
	// _ = "end of CoverTab[63692]"
//line /usr/local/go/src/regexp/syntax/parse.go:1426
	_go_fuzz_dep_.CoverTab[63693]++

							p.flags = re2.Flags
							if re2.Cap == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1429
		_go_fuzz_dep_.CoverTab[63701]++

								p.push(re1)
//line /usr/local/go/src/regexp/syntax/parse.go:1431
		// _ = "end of CoverTab[63701]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1432
		_go_fuzz_dep_.CoverTab[63702]++
								re2.Op = OpCapture
								re2.Sub = re2.Sub0[:1]
								re2.Sub[0] = re1
								p.push(re2)
//line /usr/local/go/src/regexp/syntax/parse.go:1436
		// _ = "end of CoverTab[63702]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1437
	// _ = "end of CoverTab[63693]"
//line /usr/local/go/src/regexp/syntax/parse.go:1437
	_go_fuzz_dep_.CoverTab[63694]++
							return nil
//line /usr/local/go/src/regexp/syntax/parse.go:1438
	// _ = "end of CoverTab[63694]"
}

// parseEscape parses an escape sequence at the beginning of s
//line /usr/local/go/src/regexp/syntax/parse.go:1441
// and returns the rune.
//line /usr/local/go/src/regexp/syntax/parse.go:1443
func (p *parser) parseEscape(s string) (r rune, rest string, err error) {
//line /usr/local/go/src/regexp/syntax/parse.go:1443
	_go_fuzz_dep_.CoverTab[63703]++
							t := s[1:]
							if t == "" {
//line /usr/local/go/src/regexp/syntax/parse.go:1445
		_go_fuzz_dep_.CoverTab[63707]++
								return 0, "", &Error{ErrTrailingBackslash, ""}
//line /usr/local/go/src/regexp/syntax/parse.go:1446
		// _ = "end of CoverTab[63707]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1447
		_go_fuzz_dep_.CoverTab[63708]++
//line /usr/local/go/src/regexp/syntax/parse.go:1447
		// _ = "end of CoverTab[63708]"
//line /usr/local/go/src/regexp/syntax/parse.go:1447
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1447
	// _ = "end of CoverTab[63703]"
//line /usr/local/go/src/regexp/syntax/parse.go:1447
	_go_fuzz_dep_.CoverTab[63704]++
							c, t, err := nextRune(t)
							if err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1449
		_go_fuzz_dep_.CoverTab[63709]++
								return 0, "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1450
		// _ = "end of CoverTab[63709]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1451
		_go_fuzz_dep_.CoverTab[63710]++
//line /usr/local/go/src/regexp/syntax/parse.go:1451
		// _ = "end of CoverTab[63710]"
//line /usr/local/go/src/regexp/syntax/parse.go:1451
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1451
	// _ = "end of CoverTab[63704]"
//line /usr/local/go/src/regexp/syntax/parse.go:1451
	_go_fuzz_dep_.CoverTab[63705]++

Switch:
	switch c {
	default:
//line /usr/local/go/src/regexp/syntax/parse.go:1455
		_go_fuzz_dep_.CoverTab[63711]++
								if c < utf8.RuneSelf && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1456
			_go_fuzz_dep_.CoverTab[63728]++
//line /usr/local/go/src/regexp/syntax/parse.go:1456
			return !isalnum(c)
//line /usr/local/go/src/regexp/syntax/parse.go:1456
			// _ = "end of CoverTab[63728]"
//line /usr/local/go/src/regexp/syntax/parse.go:1456
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1456
			_go_fuzz_dep_.CoverTab[63729]++

//line /usr/local/go/src/regexp/syntax/parse.go:1461
			return c, t, nil
//line /usr/local/go/src/regexp/syntax/parse.go:1461
			// _ = "end of CoverTab[63729]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1462
			_go_fuzz_dep_.CoverTab[63730]++
//line /usr/local/go/src/regexp/syntax/parse.go:1462
			// _ = "end of CoverTab[63730]"
//line /usr/local/go/src/regexp/syntax/parse.go:1462
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1462
		// _ = "end of CoverTab[63711]"

//line /usr/local/go/src/regexp/syntax/parse.go:1465
	case '1', '2', '3', '4', '5', '6', '7':
//line /usr/local/go/src/regexp/syntax/parse.go:1465
		_go_fuzz_dep_.CoverTab[63712]++

								if t == "" || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1467
			_go_fuzz_dep_.CoverTab[63731]++
//line /usr/local/go/src/regexp/syntax/parse.go:1467
			return t[0] < '0'
//line /usr/local/go/src/regexp/syntax/parse.go:1467
			// _ = "end of CoverTab[63731]"
//line /usr/local/go/src/regexp/syntax/parse.go:1467
		}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1467
			_go_fuzz_dep_.CoverTab[63732]++
//line /usr/local/go/src/regexp/syntax/parse.go:1467
			return t[0] > '7'
//line /usr/local/go/src/regexp/syntax/parse.go:1467
			// _ = "end of CoverTab[63732]"
//line /usr/local/go/src/regexp/syntax/parse.go:1467
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1467
			_go_fuzz_dep_.CoverTab[63733]++
									break
//line /usr/local/go/src/regexp/syntax/parse.go:1468
			// _ = "end of CoverTab[63733]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1469
			_go_fuzz_dep_.CoverTab[63734]++
//line /usr/local/go/src/regexp/syntax/parse.go:1469
			// _ = "end of CoverTab[63734]"
//line /usr/local/go/src/regexp/syntax/parse.go:1469
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1469
		// _ = "end of CoverTab[63712]"
//line /usr/local/go/src/regexp/syntax/parse.go:1469
		_go_fuzz_dep_.CoverTab[63713]++
								fallthrough
//line /usr/local/go/src/regexp/syntax/parse.go:1470
		// _ = "end of CoverTab[63713]"
	case '0':
//line /usr/local/go/src/regexp/syntax/parse.go:1471
		_go_fuzz_dep_.CoverTab[63714]++

								r = c - '0'
								for i := 1; i < 3; i++ {
//line /usr/local/go/src/regexp/syntax/parse.go:1474
			_go_fuzz_dep_.CoverTab[63735]++
									if t == "" || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1475
				_go_fuzz_dep_.CoverTab[63737]++
//line /usr/local/go/src/regexp/syntax/parse.go:1475
				return t[0] < '0'
//line /usr/local/go/src/regexp/syntax/parse.go:1475
				// _ = "end of CoverTab[63737]"
//line /usr/local/go/src/regexp/syntax/parse.go:1475
			}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1475
				_go_fuzz_dep_.CoverTab[63738]++
//line /usr/local/go/src/regexp/syntax/parse.go:1475
				return t[0] > '7'
//line /usr/local/go/src/regexp/syntax/parse.go:1475
				// _ = "end of CoverTab[63738]"
//line /usr/local/go/src/regexp/syntax/parse.go:1475
			}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1475
				_go_fuzz_dep_.CoverTab[63739]++
										break
//line /usr/local/go/src/regexp/syntax/parse.go:1476
				// _ = "end of CoverTab[63739]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1477
				_go_fuzz_dep_.CoverTab[63740]++
//line /usr/local/go/src/regexp/syntax/parse.go:1477
				// _ = "end of CoverTab[63740]"
//line /usr/local/go/src/regexp/syntax/parse.go:1477
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1477
			// _ = "end of CoverTab[63735]"
//line /usr/local/go/src/regexp/syntax/parse.go:1477
			_go_fuzz_dep_.CoverTab[63736]++
									r = r*8 + rune(t[0]) - '0'
									t = t[1:]
//line /usr/local/go/src/regexp/syntax/parse.go:1479
			// _ = "end of CoverTab[63736]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1480
		// _ = "end of CoverTab[63714]"
//line /usr/local/go/src/regexp/syntax/parse.go:1480
		_go_fuzz_dep_.CoverTab[63715]++
								return r, t, nil
//line /usr/local/go/src/regexp/syntax/parse.go:1481
		// _ = "end of CoverTab[63715]"

//line /usr/local/go/src/regexp/syntax/parse.go:1484
	case 'x':
//line /usr/local/go/src/regexp/syntax/parse.go:1484
		_go_fuzz_dep_.CoverTab[63716]++
								if t == "" {
//line /usr/local/go/src/regexp/syntax/parse.go:1485
			_go_fuzz_dep_.CoverTab[63741]++
									break
//line /usr/local/go/src/regexp/syntax/parse.go:1486
			// _ = "end of CoverTab[63741]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1487
			_go_fuzz_dep_.CoverTab[63742]++
//line /usr/local/go/src/regexp/syntax/parse.go:1487
			// _ = "end of CoverTab[63742]"
//line /usr/local/go/src/regexp/syntax/parse.go:1487
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1487
		// _ = "end of CoverTab[63716]"
//line /usr/local/go/src/regexp/syntax/parse.go:1487
		_go_fuzz_dep_.CoverTab[63717]++
								if c, t, err = nextRune(t); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1488
			_go_fuzz_dep_.CoverTab[63743]++
									return 0, "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1489
			// _ = "end of CoverTab[63743]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1490
			_go_fuzz_dep_.CoverTab[63744]++
//line /usr/local/go/src/regexp/syntax/parse.go:1490
			// _ = "end of CoverTab[63744]"
//line /usr/local/go/src/regexp/syntax/parse.go:1490
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1490
		// _ = "end of CoverTab[63717]"
//line /usr/local/go/src/regexp/syntax/parse.go:1490
		_go_fuzz_dep_.CoverTab[63718]++
								if c == '{' {
//line /usr/local/go/src/regexp/syntax/parse.go:1491
			_go_fuzz_dep_.CoverTab[63745]++

//line /usr/local/go/src/regexp/syntax/parse.go:1496
			nhex := 0
			r = 0
			for {
//line /usr/local/go/src/regexp/syntax/parse.go:1498
				_go_fuzz_dep_.CoverTab[63748]++
										if t == "" {
//line /usr/local/go/src/regexp/syntax/parse.go:1499
					_go_fuzz_dep_.CoverTab[63754]++
											break Switch
//line /usr/local/go/src/regexp/syntax/parse.go:1500
					// _ = "end of CoverTab[63754]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1501
					_go_fuzz_dep_.CoverTab[63755]++
//line /usr/local/go/src/regexp/syntax/parse.go:1501
					// _ = "end of CoverTab[63755]"
//line /usr/local/go/src/regexp/syntax/parse.go:1501
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1501
				// _ = "end of CoverTab[63748]"
//line /usr/local/go/src/regexp/syntax/parse.go:1501
				_go_fuzz_dep_.CoverTab[63749]++
										if c, t, err = nextRune(t); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1502
					_go_fuzz_dep_.CoverTab[63756]++
											return 0, "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1503
					// _ = "end of CoverTab[63756]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1504
					_go_fuzz_dep_.CoverTab[63757]++
//line /usr/local/go/src/regexp/syntax/parse.go:1504
					// _ = "end of CoverTab[63757]"
//line /usr/local/go/src/regexp/syntax/parse.go:1504
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1504
				// _ = "end of CoverTab[63749]"
//line /usr/local/go/src/regexp/syntax/parse.go:1504
				_go_fuzz_dep_.CoverTab[63750]++
										if c == '}' {
//line /usr/local/go/src/regexp/syntax/parse.go:1505
					_go_fuzz_dep_.CoverTab[63758]++
											break
//line /usr/local/go/src/regexp/syntax/parse.go:1506
					// _ = "end of CoverTab[63758]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1507
					_go_fuzz_dep_.CoverTab[63759]++
//line /usr/local/go/src/regexp/syntax/parse.go:1507
					// _ = "end of CoverTab[63759]"
//line /usr/local/go/src/regexp/syntax/parse.go:1507
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1507
				// _ = "end of CoverTab[63750]"
//line /usr/local/go/src/regexp/syntax/parse.go:1507
				_go_fuzz_dep_.CoverTab[63751]++
										v := unhex(c)
										if v < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1509
					_go_fuzz_dep_.CoverTab[63760]++
											break Switch
//line /usr/local/go/src/regexp/syntax/parse.go:1510
					// _ = "end of CoverTab[63760]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1511
					_go_fuzz_dep_.CoverTab[63761]++
//line /usr/local/go/src/regexp/syntax/parse.go:1511
					// _ = "end of CoverTab[63761]"
//line /usr/local/go/src/regexp/syntax/parse.go:1511
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1511
				// _ = "end of CoverTab[63751]"
//line /usr/local/go/src/regexp/syntax/parse.go:1511
				_go_fuzz_dep_.CoverTab[63752]++
										r = r*16 + v
										if r > unicode.MaxRune {
//line /usr/local/go/src/regexp/syntax/parse.go:1513
					_go_fuzz_dep_.CoverTab[63762]++
											break Switch
//line /usr/local/go/src/regexp/syntax/parse.go:1514
					// _ = "end of CoverTab[63762]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1515
					_go_fuzz_dep_.CoverTab[63763]++
//line /usr/local/go/src/regexp/syntax/parse.go:1515
					// _ = "end of CoverTab[63763]"
//line /usr/local/go/src/regexp/syntax/parse.go:1515
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1515
				// _ = "end of CoverTab[63752]"
//line /usr/local/go/src/regexp/syntax/parse.go:1515
				_go_fuzz_dep_.CoverTab[63753]++
										nhex++
//line /usr/local/go/src/regexp/syntax/parse.go:1516
				// _ = "end of CoverTab[63753]"
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1517
			// _ = "end of CoverTab[63745]"
//line /usr/local/go/src/regexp/syntax/parse.go:1517
			_go_fuzz_dep_.CoverTab[63746]++
									if nhex == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1518
				_go_fuzz_dep_.CoverTab[63764]++
										break Switch
//line /usr/local/go/src/regexp/syntax/parse.go:1519
				// _ = "end of CoverTab[63764]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1520
				_go_fuzz_dep_.CoverTab[63765]++
//line /usr/local/go/src/regexp/syntax/parse.go:1520
				// _ = "end of CoverTab[63765]"
//line /usr/local/go/src/regexp/syntax/parse.go:1520
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1520
			// _ = "end of CoverTab[63746]"
//line /usr/local/go/src/regexp/syntax/parse.go:1520
			_go_fuzz_dep_.CoverTab[63747]++
									return r, t, nil
//line /usr/local/go/src/regexp/syntax/parse.go:1521
			// _ = "end of CoverTab[63747]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1522
			_go_fuzz_dep_.CoverTab[63766]++
//line /usr/local/go/src/regexp/syntax/parse.go:1522
			// _ = "end of CoverTab[63766]"
//line /usr/local/go/src/regexp/syntax/parse.go:1522
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1522
		// _ = "end of CoverTab[63718]"
//line /usr/local/go/src/regexp/syntax/parse.go:1522
		_go_fuzz_dep_.CoverTab[63719]++

//line /usr/local/go/src/regexp/syntax/parse.go:1525
		x := unhex(c)
		if c, t, err = nextRune(t); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1526
			_go_fuzz_dep_.CoverTab[63767]++
									return 0, "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1527
			// _ = "end of CoverTab[63767]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1528
			_go_fuzz_dep_.CoverTab[63768]++
//line /usr/local/go/src/regexp/syntax/parse.go:1528
			// _ = "end of CoverTab[63768]"
//line /usr/local/go/src/regexp/syntax/parse.go:1528
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1528
		// _ = "end of CoverTab[63719]"
//line /usr/local/go/src/regexp/syntax/parse.go:1528
		_go_fuzz_dep_.CoverTab[63720]++
								y := unhex(c)
								if x < 0 || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1530
			_go_fuzz_dep_.CoverTab[63769]++
//line /usr/local/go/src/regexp/syntax/parse.go:1530
			return y < 0
//line /usr/local/go/src/regexp/syntax/parse.go:1530
			// _ = "end of CoverTab[63769]"
//line /usr/local/go/src/regexp/syntax/parse.go:1530
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1530
			_go_fuzz_dep_.CoverTab[63770]++
									break
//line /usr/local/go/src/regexp/syntax/parse.go:1531
			// _ = "end of CoverTab[63770]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1532
			_go_fuzz_dep_.CoverTab[63771]++
//line /usr/local/go/src/regexp/syntax/parse.go:1532
			// _ = "end of CoverTab[63771]"
//line /usr/local/go/src/regexp/syntax/parse.go:1532
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1532
		// _ = "end of CoverTab[63720]"
//line /usr/local/go/src/regexp/syntax/parse.go:1532
		_go_fuzz_dep_.CoverTab[63721]++
								return x*16 + y, t, nil
//line /usr/local/go/src/regexp/syntax/parse.go:1533
		// _ = "end of CoverTab[63721]"

//line /usr/local/go/src/regexp/syntax/parse.go:1541
	case 'a':
//line /usr/local/go/src/regexp/syntax/parse.go:1541
		_go_fuzz_dep_.CoverTab[63722]++
								return '\a', t, err
//line /usr/local/go/src/regexp/syntax/parse.go:1542
		// _ = "end of CoverTab[63722]"
	case 'f':
//line /usr/local/go/src/regexp/syntax/parse.go:1543
		_go_fuzz_dep_.CoverTab[63723]++
								return '\f', t, err
//line /usr/local/go/src/regexp/syntax/parse.go:1544
		// _ = "end of CoverTab[63723]"
	case 'n':
//line /usr/local/go/src/regexp/syntax/parse.go:1545
		_go_fuzz_dep_.CoverTab[63724]++
								return '\n', t, err
//line /usr/local/go/src/regexp/syntax/parse.go:1546
		// _ = "end of CoverTab[63724]"
	case 'r':
//line /usr/local/go/src/regexp/syntax/parse.go:1547
		_go_fuzz_dep_.CoverTab[63725]++
								return '\r', t, err
//line /usr/local/go/src/regexp/syntax/parse.go:1548
		// _ = "end of CoverTab[63725]"
	case 't':
//line /usr/local/go/src/regexp/syntax/parse.go:1549
		_go_fuzz_dep_.CoverTab[63726]++
								return '\t', t, err
//line /usr/local/go/src/regexp/syntax/parse.go:1550
		// _ = "end of CoverTab[63726]"
	case 'v':
//line /usr/local/go/src/regexp/syntax/parse.go:1551
		_go_fuzz_dep_.CoverTab[63727]++
								return '\v', t, err
//line /usr/local/go/src/regexp/syntax/parse.go:1552
		// _ = "end of CoverTab[63727]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1553
	// _ = "end of CoverTab[63705]"
//line /usr/local/go/src/regexp/syntax/parse.go:1553
	_go_fuzz_dep_.CoverTab[63706]++
							return 0, "", &Error{ErrInvalidEscape, s[:len(s)-len(t)]}
//line /usr/local/go/src/regexp/syntax/parse.go:1554
	// _ = "end of CoverTab[63706]"
}

// parseClassChar parses a character class character at the beginning of s
//line /usr/local/go/src/regexp/syntax/parse.go:1557
// and returns it.
//line /usr/local/go/src/regexp/syntax/parse.go:1559
func (p *parser) parseClassChar(s, wholeClass string) (r rune, rest string, err error) {
//line /usr/local/go/src/regexp/syntax/parse.go:1559
	_go_fuzz_dep_.CoverTab[63772]++
							if s == "" {
//line /usr/local/go/src/regexp/syntax/parse.go:1560
		_go_fuzz_dep_.CoverTab[63775]++
								return 0, "", &Error{Code: ErrMissingBracket, Expr: wholeClass}
//line /usr/local/go/src/regexp/syntax/parse.go:1561
		// _ = "end of CoverTab[63775]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1562
		_go_fuzz_dep_.CoverTab[63776]++
//line /usr/local/go/src/regexp/syntax/parse.go:1562
		// _ = "end of CoverTab[63776]"
//line /usr/local/go/src/regexp/syntax/parse.go:1562
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1562
	// _ = "end of CoverTab[63772]"
//line /usr/local/go/src/regexp/syntax/parse.go:1562
	_go_fuzz_dep_.CoverTab[63773]++

//line /usr/local/go/src/regexp/syntax/parse.go:1566
	if s[0] == '\\' {
//line /usr/local/go/src/regexp/syntax/parse.go:1566
		_go_fuzz_dep_.CoverTab[63777]++
								return p.parseEscape(s)
//line /usr/local/go/src/regexp/syntax/parse.go:1567
		// _ = "end of CoverTab[63777]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1568
		_go_fuzz_dep_.CoverTab[63778]++
//line /usr/local/go/src/regexp/syntax/parse.go:1568
		// _ = "end of CoverTab[63778]"
//line /usr/local/go/src/regexp/syntax/parse.go:1568
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1568
	// _ = "end of CoverTab[63773]"
//line /usr/local/go/src/regexp/syntax/parse.go:1568
	_go_fuzz_dep_.CoverTab[63774]++

							return nextRune(s)
//line /usr/local/go/src/regexp/syntax/parse.go:1570
	// _ = "end of CoverTab[63774]"
}

type charGroup struct {
	sign	int
	class	[]rune
}

// parsePerlClassEscape parses a leading Perl character class escape like \d
//line /usr/local/go/src/regexp/syntax/parse.go:1578
// from the beginning of s. If one is present, it appends the characters to r
//line /usr/local/go/src/regexp/syntax/parse.go:1578
// and returns the new slice r and the remainder of the string.
//line /usr/local/go/src/regexp/syntax/parse.go:1581
func (p *parser) parsePerlClassEscape(s string, r []rune) (out []rune, rest string) {
//line /usr/local/go/src/regexp/syntax/parse.go:1581
	_go_fuzz_dep_.CoverTab[63779]++
							if p.flags&PerlX == 0 || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1582
		_go_fuzz_dep_.CoverTab[63782]++
//line /usr/local/go/src/regexp/syntax/parse.go:1582
		return len(s) < 2
//line /usr/local/go/src/regexp/syntax/parse.go:1582
		// _ = "end of CoverTab[63782]"
//line /usr/local/go/src/regexp/syntax/parse.go:1582
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1582
		_go_fuzz_dep_.CoverTab[63783]++
//line /usr/local/go/src/regexp/syntax/parse.go:1582
		return s[0] != '\\'
//line /usr/local/go/src/regexp/syntax/parse.go:1582
		// _ = "end of CoverTab[63783]"
//line /usr/local/go/src/regexp/syntax/parse.go:1582
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1582
		_go_fuzz_dep_.CoverTab[63784]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1583
		// _ = "end of CoverTab[63784]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1584
		_go_fuzz_dep_.CoverTab[63785]++
//line /usr/local/go/src/regexp/syntax/parse.go:1584
		// _ = "end of CoverTab[63785]"
//line /usr/local/go/src/regexp/syntax/parse.go:1584
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1584
	// _ = "end of CoverTab[63779]"
//line /usr/local/go/src/regexp/syntax/parse.go:1584
	_go_fuzz_dep_.CoverTab[63780]++
							g := perlGroup[s[0:2]]
							if g.sign == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1586
		_go_fuzz_dep_.CoverTab[63786]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1587
		// _ = "end of CoverTab[63786]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1588
		_go_fuzz_dep_.CoverTab[63787]++
//line /usr/local/go/src/regexp/syntax/parse.go:1588
		// _ = "end of CoverTab[63787]"
//line /usr/local/go/src/regexp/syntax/parse.go:1588
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1588
	// _ = "end of CoverTab[63780]"
//line /usr/local/go/src/regexp/syntax/parse.go:1588
	_go_fuzz_dep_.CoverTab[63781]++
							return p.appendGroup(r, g), s[2:]
//line /usr/local/go/src/regexp/syntax/parse.go:1589
	// _ = "end of CoverTab[63781]"
}

// parseNamedClass parses a leading POSIX named character class like [:alnum:]
//line /usr/local/go/src/regexp/syntax/parse.go:1592
// from the beginning of s. If one is present, it appends the characters to r
//line /usr/local/go/src/regexp/syntax/parse.go:1592
// and returns the new slice r and the remainder of the string.
//line /usr/local/go/src/regexp/syntax/parse.go:1595
func (p *parser) parseNamedClass(s string, r []rune) (out []rune, rest string, err error) {
//line /usr/local/go/src/regexp/syntax/parse.go:1595
	_go_fuzz_dep_.CoverTab[63788]++
							if len(s) < 2 || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1596
		_go_fuzz_dep_.CoverTab[63792]++
//line /usr/local/go/src/regexp/syntax/parse.go:1596
		return s[0] != '['
//line /usr/local/go/src/regexp/syntax/parse.go:1596
		// _ = "end of CoverTab[63792]"
//line /usr/local/go/src/regexp/syntax/parse.go:1596
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1596
		_go_fuzz_dep_.CoverTab[63793]++
//line /usr/local/go/src/regexp/syntax/parse.go:1596
		return s[1] != ':'
//line /usr/local/go/src/regexp/syntax/parse.go:1596
		// _ = "end of CoverTab[63793]"
//line /usr/local/go/src/regexp/syntax/parse.go:1596
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1596
		_go_fuzz_dep_.CoverTab[63794]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1597
		// _ = "end of CoverTab[63794]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1598
		_go_fuzz_dep_.CoverTab[63795]++
//line /usr/local/go/src/regexp/syntax/parse.go:1598
		// _ = "end of CoverTab[63795]"
//line /usr/local/go/src/regexp/syntax/parse.go:1598
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1598
	// _ = "end of CoverTab[63788]"
//line /usr/local/go/src/regexp/syntax/parse.go:1598
	_go_fuzz_dep_.CoverTab[63789]++

							i := strings.Index(s[2:], ":]")
							if i < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1601
		_go_fuzz_dep_.CoverTab[63796]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1602
		// _ = "end of CoverTab[63796]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1603
		_go_fuzz_dep_.CoverTab[63797]++
//line /usr/local/go/src/regexp/syntax/parse.go:1603
		// _ = "end of CoverTab[63797]"
//line /usr/local/go/src/regexp/syntax/parse.go:1603
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1603
	// _ = "end of CoverTab[63789]"
//line /usr/local/go/src/regexp/syntax/parse.go:1603
	_go_fuzz_dep_.CoverTab[63790]++
							i += 2
							name, s := s[0:i+2], s[i+2:]
							g := posixGroup[name]
							if g.sign == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1607
		_go_fuzz_dep_.CoverTab[63798]++
								return nil, "", &Error{ErrInvalidCharRange, name}
//line /usr/local/go/src/regexp/syntax/parse.go:1608
		// _ = "end of CoverTab[63798]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1609
		_go_fuzz_dep_.CoverTab[63799]++
//line /usr/local/go/src/regexp/syntax/parse.go:1609
		// _ = "end of CoverTab[63799]"
//line /usr/local/go/src/regexp/syntax/parse.go:1609
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1609
	// _ = "end of CoverTab[63790]"
//line /usr/local/go/src/regexp/syntax/parse.go:1609
	_go_fuzz_dep_.CoverTab[63791]++
							return p.appendGroup(r, g), s, nil
//line /usr/local/go/src/regexp/syntax/parse.go:1610
	// _ = "end of CoverTab[63791]"
}

func (p *parser) appendGroup(r []rune, g charGroup) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1613
	_go_fuzz_dep_.CoverTab[63800]++
							if p.flags&FoldCase == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1614
		_go_fuzz_dep_.CoverTab[63802]++
								if g.sign < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1615
			_go_fuzz_dep_.CoverTab[63803]++
									r = appendNegatedClass(r, g.class)
//line /usr/local/go/src/regexp/syntax/parse.go:1616
			// _ = "end of CoverTab[63803]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1617
			_go_fuzz_dep_.CoverTab[63804]++
									r = appendClass(r, g.class)
//line /usr/local/go/src/regexp/syntax/parse.go:1618
			// _ = "end of CoverTab[63804]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1619
		// _ = "end of CoverTab[63802]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1620
		_go_fuzz_dep_.CoverTab[63805]++
								tmp := p.tmpClass[:0]
								tmp = appendFoldedClass(tmp, g.class)
								p.tmpClass = tmp
								tmp = cleanClass(&p.tmpClass)
								if g.sign < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1625
			_go_fuzz_dep_.CoverTab[63806]++
									r = appendNegatedClass(r, tmp)
//line /usr/local/go/src/regexp/syntax/parse.go:1626
			// _ = "end of CoverTab[63806]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1627
			_go_fuzz_dep_.CoverTab[63807]++
									r = appendClass(r, tmp)
//line /usr/local/go/src/regexp/syntax/parse.go:1628
			// _ = "end of CoverTab[63807]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1629
		// _ = "end of CoverTab[63805]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1630
	// _ = "end of CoverTab[63800]"
//line /usr/local/go/src/regexp/syntax/parse.go:1630
	_go_fuzz_dep_.CoverTab[63801]++
							return r
//line /usr/local/go/src/regexp/syntax/parse.go:1631
	// _ = "end of CoverTab[63801]"
}

var anyTable = &unicode.RangeTable{
	R16:	[]unicode.Range16{{Lo: 0, Hi: 1<<16 - 1, Stride: 1}},
	R32:	[]unicode.Range32{{Lo: 1 << 16, Hi: unicode.MaxRune, Stride: 1}},
}

// unicodeTable returns the unicode.RangeTable identified by name
//line /usr/local/go/src/regexp/syntax/parse.go:1639
// and the table of additional fold-equivalent code points.
//line /usr/local/go/src/regexp/syntax/parse.go:1641
func unicodeTable(name string) (*unicode.RangeTable, *unicode.RangeTable) {
//line /usr/local/go/src/regexp/syntax/parse.go:1641
	_go_fuzz_dep_.CoverTab[63808]++

							if name == "Any" {
//line /usr/local/go/src/regexp/syntax/parse.go:1643
		_go_fuzz_dep_.CoverTab[63812]++
								return anyTable, anyTable
//line /usr/local/go/src/regexp/syntax/parse.go:1644
		// _ = "end of CoverTab[63812]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1645
		_go_fuzz_dep_.CoverTab[63813]++
//line /usr/local/go/src/regexp/syntax/parse.go:1645
		// _ = "end of CoverTab[63813]"
//line /usr/local/go/src/regexp/syntax/parse.go:1645
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1645
	// _ = "end of CoverTab[63808]"
//line /usr/local/go/src/regexp/syntax/parse.go:1645
	_go_fuzz_dep_.CoverTab[63809]++
							if t := unicode.Categories[name]; t != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1646
		_go_fuzz_dep_.CoverTab[63814]++
								return t, unicode.FoldCategory[name]
//line /usr/local/go/src/regexp/syntax/parse.go:1647
		// _ = "end of CoverTab[63814]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1648
		_go_fuzz_dep_.CoverTab[63815]++
//line /usr/local/go/src/regexp/syntax/parse.go:1648
		// _ = "end of CoverTab[63815]"
//line /usr/local/go/src/regexp/syntax/parse.go:1648
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1648
	// _ = "end of CoverTab[63809]"
//line /usr/local/go/src/regexp/syntax/parse.go:1648
	_go_fuzz_dep_.CoverTab[63810]++
							if t := unicode.Scripts[name]; t != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1649
		_go_fuzz_dep_.CoverTab[63816]++
								return t, unicode.FoldScript[name]
//line /usr/local/go/src/regexp/syntax/parse.go:1650
		// _ = "end of CoverTab[63816]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1651
		_go_fuzz_dep_.CoverTab[63817]++
//line /usr/local/go/src/regexp/syntax/parse.go:1651
		// _ = "end of CoverTab[63817]"
//line /usr/local/go/src/regexp/syntax/parse.go:1651
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1651
	// _ = "end of CoverTab[63810]"
//line /usr/local/go/src/regexp/syntax/parse.go:1651
	_go_fuzz_dep_.CoverTab[63811]++
							return nil, nil
//line /usr/local/go/src/regexp/syntax/parse.go:1652
	// _ = "end of CoverTab[63811]"
}

// parseUnicodeClass parses a leading Unicode character class like \p{Han}
//line /usr/local/go/src/regexp/syntax/parse.go:1655
// from the beginning of s. If one is present, it appends the characters to r
//line /usr/local/go/src/regexp/syntax/parse.go:1655
// and returns the new slice r and the remainder of the string.
//line /usr/local/go/src/regexp/syntax/parse.go:1658
func (p *parser) parseUnicodeClass(s string, r []rune) (out []rune, rest string, err error) {
//line /usr/local/go/src/regexp/syntax/parse.go:1658
	_go_fuzz_dep_.CoverTab[63818]++
							if p.flags&UnicodeGroups == 0 || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		_go_fuzz_dep_.CoverTab[63826]++
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		return len(s) < 2
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		// _ = "end of CoverTab[63826]"
//line /usr/local/go/src/regexp/syntax/parse.go:1659
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		_go_fuzz_dep_.CoverTab[63827]++
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		return s[0] != '\\'
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		// _ = "end of CoverTab[63827]"
//line /usr/local/go/src/regexp/syntax/parse.go:1659
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		_go_fuzz_dep_.CoverTab[63828]++
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		return s[1] != 'p' && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1659
			_go_fuzz_dep_.CoverTab[63829]++
//line /usr/local/go/src/regexp/syntax/parse.go:1659
			return s[1] != 'P'
//line /usr/local/go/src/regexp/syntax/parse.go:1659
			// _ = "end of CoverTab[63829]"
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		}()
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		// _ = "end of CoverTab[63828]"
//line /usr/local/go/src/regexp/syntax/parse.go:1659
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1659
		_go_fuzz_dep_.CoverTab[63830]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1660
		// _ = "end of CoverTab[63830]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1661
		_go_fuzz_dep_.CoverTab[63831]++
//line /usr/local/go/src/regexp/syntax/parse.go:1661
		// _ = "end of CoverTab[63831]"
//line /usr/local/go/src/regexp/syntax/parse.go:1661
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1661
	// _ = "end of CoverTab[63818]"
//line /usr/local/go/src/regexp/syntax/parse.go:1661
	_go_fuzz_dep_.CoverTab[63819]++

//line /usr/local/go/src/regexp/syntax/parse.go:1664
	sign := +1
	if s[1] == 'P' {
//line /usr/local/go/src/regexp/syntax/parse.go:1665
		_go_fuzz_dep_.CoverTab[63832]++
								sign = -1
//line /usr/local/go/src/regexp/syntax/parse.go:1666
		// _ = "end of CoverTab[63832]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1667
		_go_fuzz_dep_.CoverTab[63833]++
//line /usr/local/go/src/regexp/syntax/parse.go:1667
		// _ = "end of CoverTab[63833]"
//line /usr/local/go/src/regexp/syntax/parse.go:1667
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1667
	// _ = "end of CoverTab[63819]"
//line /usr/local/go/src/regexp/syntax/parse.go:1667
	_go_fuzz_dep_.CoverTab[63820]++
							t := s[2:]
							c, t, err := nextRune(t)
							if err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1670
		_go_fuzz_dep_.CoverTab[63834]++
								return
//line /usr/local/go/src/regexp/syntax/parse.go:1671
		// _ = "end of CoverTab[63834]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1672
		_go_fuzz_dep_.CoverTab[63835]++
//line /usr/local/go/src/regexp/syntax/parse.go:1672
		// _ = "end of CoverTab[63835]"
//line /usr/local/go/src/regexp/syntax/parse.go:1672
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1672
	// _ = "end of CoverTab[63820]"
//line /usr/local/go/src/regexp/syntax/parse.go:1672
	_go_fuzz_dep_.CoverTab[63821]++
							var seq, name string
							if c != '{' {
//line /usr/local/go/src/regexp/syntax/parse.go:1674
		_go_fuzz_dep_.CoverTab[63836]++

								seq = s[:len(s)-len(t)]
								name = seq[2:]
//line /usr/local/go/src/regexp/syntax/parse.go:1677
		// _ = "end of CoverTab[63836]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1678
		_go_fuzz_dep_.CoverTab[63837]++

								end := strings.IndexRune(s, '}')
								if end < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1681
			_go_fuzz_dep_.CoverTab[63839]++
									if err = checkUTF8(s); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1682
				_go_fuzz_dep_.CoverTab[63841]++
										return
//line /usr/local/go/src/regexp/syntax/parse.go:1683
				// _ = "end of CoverTab[63841]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1684
				_go_fuzz_dep_.CoverTab[63842]++
//line /usr/local/go/src/regexp/syntax/parse.go:1684
				// _ = "end of CoverTab[63842]"
//line /usr/local/go/src/regexp/syntax/parse.go:1684
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1684
			// _ = "end of CoverTab[63839]"
//line /usr/local/go/src/regexp/syntax/parse.go:1684
			_go_fuzz_dep_.CoverTab[63840]++
									return nil, "", &Error{ErrInvalidCharRange, s}
//line /usr/local/go/src/regexp/syntax/parse.go:1685
			// _ = "end of CoverTab[63840]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1686
			_go_fuzz_dep_.CoverTab[63843]++
//line /usr/local/go/src/regexp/syntax/parse.go:1686
			// _ = "end of CoverTab[63843]"
//line /usr/local/go/src/regexp/syntax/parse.go:1686
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1686
		// _ = "end of CoverTab[63837]"
//line /usr/local/go/src/regexp/syntax/parse.go:1686
		_go_fuzz_dep_.CoverTab[63838]++
								seq, t = s[:end+1], s[end+1:]
								name = s[3:end]
								if err = checkUTF8(name); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1689
			_go_fuzz_dep_.CoverTab[63844]++
									return
//line /usr/local/go/src/regexp/syntax/parse.go:1690
			// _ = "end of CoverTab[63844]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1691
			_go_fuzz_dep_.CoverTab[63845]++
//line /usr/local/go/src/regexp/syntax/parse.go:1691
			// _ = "end of CoverTab[63845]"
//line /usr/local/go/src/regexp/syntax/parse.go:1691
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1691
		// _ = "end of CoverTab[63838]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1692
	// _ = "end of CoverTab[63821]"
//line /usr/local/go/src/regexp/syntax/parse.go:1692
	_go_fuzz_dep_.CoverTab[63822]++

//line /usr/local/go/src/regexp/syntax/parse.go:1695
	if name != "" && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1695
		_go_fuzz_dep_.CoverTab[63846]++
//line /usr/local/go/src/regexp/syntax/parse.go:1695
		return name[0] == '^'
//line /usr/local/go/src/regexp/syntax/parse.go:1695
		// _ = "end of CoverTab[63846]"
//line /usr/local/go/src/regexp/syntax/parse.go:1695
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1695
		_go_fuzz_dep_.CoverTab[63847]++
								sign = -sign
								name = name[1:]
//line /usr/local/go/src/regexp/syntax/parse.go:1697
		// _ = "end of CoverTab[63847]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1698
		_go_fuzz_dep_.CoverTab[63848]++
//line /usr/local/go/src/regexp/syntax/parse.go:1698
		// _ = "end of CoverTab[63848]"
//line /usr/local/go/src/regexp/syntax/parse.go:1698
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1698
	// _ = "end of CoverTab[63822]"
//line /usr/local/go/src/regexp/syntax/parse.go:1698
	_go_fuzz_dep_.CoverTab[63823]++

							tab, fold := unicodeTable(name)
							if tab == nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1701
		_go_fuzz_dep_.CoverTab[63849]++
								return nil, "", &Error{ErrInvalidCharRange, seq}
//line /usr/local/go/src/regexp/syntax/parse.go:1702
		// _ = "end of CoverTab[63849]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1703
		_go_fuzz_dep_.CoverTab[63850]++
//line /usr/local/go/src/regexp/syntax/parse.go:1703
		// _ = "end of CoverTab[63850]"
//line /usr/local/go/src/regexp/syntax/parse.go:1703
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1703
	// _ = "end of CoverTab[63823]"
//line /usr/local/go/src/regexp/syntax/parse.go:1703
	_go_fuzz_dep_.CoverTab[63824]++

							if p.flags&FoldCase == 0 || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1705
		_go_fuzz_dep_.CoverTab[63851]++
//line /usr/local/go/src/regexp/syntax/parse.go:1705
		return fold == nil
//line /usr/local/go/src/regexp/syntax/parse.go:1705
		// _ = "end of CoverTab[63851]"
//line /usr/local/go/src/regexp/syntax/parse.go:1705
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1705
		_go_fuzz_dep_.CoverTab[63852]++
								if sign > 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1706
			_go_fuzz_dep_.CoverTab[63853]++
									r = appendTable(r, tab)
//line /usr/local/go/src/regexp/syntax/parse.go:1707
			// _ = "end of CoverTab[63853]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1708
			_go_fuzz_dep_.CoverTab[63854]++
									r = appendNegatedTable(r, tab)
//line /usr/local/go/src/regexp/syntax/parse.go:1709
			// _ = "end of CoverTab[63854]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1710
		// _ = "end of CoverTab[63852]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1711
		_go_fuzz_dep_.CoverTab[63855]++

//line /usr/local/go/src/regexp/syntax/parse.go:1715
		tmp := p.tmpClass[:0]
		tmp = appendTable(tmp, tab)
		tmp = appendTable(tmp, fold)
		p.tmpClass = tmp
		tmp = cleanClass(&p.tmpClass)
		if sign > 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1720
			_go_fuzz_dep_.CoverTab[63856]++
									r = appendClass(r, tmp)
//line /usr/local/go/src/regexp/syntax/parse.go:1721
			// _ = "end of CoverTab[63856]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1722
			_go_fuzz_dep_.CoverTab[63857]++
									r = appendNegatedClass(r, tmp)
//line /usr/local/go/src/regexp/syntax/parse.go:1723
			// _ = "end of CoverTab[63857]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1724
		// _ = "end of CoverTab[63855]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1725
	// _ = "end of CoverTab[63824]"
//line /usr/local/go/src/regexp/syntax/parse.go:1725
	_go_fuzz_dep_.CoverTab[63825]++
							return r, t, nil
//line /usr/local/go/src/regexp/syntax/parse.go:1726
	// _ = "end of CoverTab[63825]"
}

// parseClass parses a character class at the beginning of s
//line /usr/local/go/src/regexp/syntax/parse.go:1729
// and pushes it onto the parse stack.
//line /usr/local/go/src/regexp/syntax/parse.go:1731
func (p *parser) parseClass(s string) (rest string, err error) {
//line /usr/local/go/src/regexp/syntax/parse.go:1731
	_go_fuzz_dep_.CoverTab[63858]++
							t := s[1:]
							re := p.newRegexp(OpCharClass)
							re.Flags = p.flags
							re.Rune = re.Rune0[:0]

							sign := +1
							if t != "" && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1738
		_go_fuzz_dep_.CoverTab[63862]++
//line /usr/local/go/src/regexp/syntax/parse.go:1738
		return t[0] == '^'
//line /usr/local/go/src/regexp/syntax/parse.go:1738
		// _ = "end of CoverTab[63862]"
//line /usr/local/go/src/regexp/syntax/parse.go:1738
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1738
		_go_fuzz_dep_.CoverTab[63863]++
								sign = -1
								t = t[1:]

//line /usr/local/go/src/regexp/syntax/parse.go:1744
		if p.flags&ClassNL == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1744
			_go_fuzz_dep_.CoverTab[63864]++
									re.Rune = append(re.Rune, '\n', '\n')
//line /usr/local/go/src/regexp/syntax/parse.go:1745
			// _ = "end of CoverTab[63864]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1746
			_go_fuzz_dep_.CoverTab[63865]++
//line /usr/local/go/src/regexp/syntax/parse.go:1746
			// _ = "end of CoverTab[63865]"
//line /usr/local/go/src/regexp/syntax/parse.go:1746
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1746
		// _ = "end of CoverTab[63863]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1747
		_go_fuzz_dep_.CoverTab[63866]++
//line /usr/local/go/src/regexp/syntax/parse.go:1747
		// _ = "end of CoverTab[63866]"
//line /usr/local/go/src/regexp/syntax/parse.go:1747
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1747
	// _ = "end of CoverTab[63858]"
//line /usr/local/go/src/regexp/syntax/parse.go:1747
	_go_fuzz_dep_.CoverTab[63859]++

							class := re.Rune
							first := true
							for t == "" || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1751
		_go_fuzz_dep_.CoverTab[63867]++
//line /usr/local/go/src/regexp/syntax/parse.go:1751
		return t[0] != ']'
//line /usr/local/go/src/regexp/syntax/parse.go:1751
		// _ = "end of CoverTab[63867]"
//line /usr/local/go/src/regexp/syntax/parse.go:1751
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1751
		_go_fuzz_dep_.CoverTab[63868]++
//line /usr/local/go/src/regexp/syntax/parse.go:1751
		return first
//line /usr/local/go/src/regexp/syntax/parse.go:1751
		// _ = "end of CoverTab[63868]"
//line /usr/local/go/src/regexp/syntax/parse.go:1751
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1751
		_go_fuzz_dep_.CoverTab[63869]++

//line /usr/local/go/src/regexp/syntax/parse.go:1754
		if t != "" && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			_go_fuzz_dep_.CoverTab[63877]++
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			return t[0] == '-'
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			// _ = "end of CoverTab[63877]"
//line /usr/local/go/src/regexp/syntax/parse.go:1754
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			_go_fuzz_dep_.CoverTab[63878]++
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			return p.flags&PerlX == 0
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			// _ = "end of CoverTab[63878]"
//line /usr/local/go/src/regexp/syntax/parse.go:1754
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			_go_fuzz_dep_.CoverTab[63879]++
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			return !first
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			// _ = "end of CoverTab[63879]"
//line /usr/local/go/src/regexp/syntax/parse.go:1754
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			_go_fuzz_dep_.CoverTab[63880]++
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			return (len(t) == 1 || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1754
				_go_fuzz_dep_.CoverTab[63881]++
//line /usr/local/go/src/regexp/syntax/parse.go:1754
				return t[1] != ']'
//line /usr/local/go/src/regexp/syntax/parse.go:1754
				// _ = "end of CoverTab[63881]"
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			}())
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			// _ = "end of CoverTab[63880]"
//line /usr/local/go/src/regexp/syntax/parse.go:1754
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1754
			_go_fuzz_dep_.CoverTab[63882]++
									_, size := utf8.DecodeRuneInString(t[1:])
									return "", &Error{Code: ErrInvalidCharRange, Expr: t[:1+size]}
//line /usr/local/go/src/regexp/syntax/parse.go:1756
			// _ = "end of CoverTab[63882]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1757
			_go_fuzz_dep_.CoverTab[63883]++
//line /usr/local/go/src/regexp/syntax/parse.go:1757
			// _ = "end of CoverTab[63883]"
//line /usr/local/go/src/regexp/syntax/parse.go:1757
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1757
		// _ = "end of CoverTab[63869]"
//line /usr/local/go/src/regexp/syntax/parse.go:1757
		_go_fuzz_dep_.CoverTab[63870]++
								first = false

//line /usr/local/go/src/regexp/syntax/parse.go:1761
		if len(t) > 2 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1761
			_go_fuzz_dep_.CoverTab[63884]++
//line /usr/local/go/src/regexp/syntax/parse.go:1761
			return t[0] == '['
//line /usr/local/go/src/regexp/syntax/parse.go:1761
			// _ = "end of CoverTab[63884]"
//line /usr/local/go/src/regexp/syntax/parse.go:1761
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1761
			_go_fuzz_dep_.CoverTab[63885]++
//line /usr/local/go/src/regexp/syntax/parse.go:1761
			return t[1] == ':'
//line /usr/local/go/src/regexp/syntax/parse.go:1761
			// _ = "end of CoverTab[63885]"
//line /usr/local/go/src/regexp/syntax/parse.go:1761
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1761
			_go_fuzz_dep_.CoverTab[63886]++
									nclass, nt, err := p.parseNamedClass(t, class)
									if err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1763
				_go_fuzz_dep_.CoverTab[63888]++
										return "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1764
				// _ = "end of CoverTab[63888]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1765
				_go_fuzz_dep_.CoverTab[63889]++
//line /usr/local/go/src/regexp/syntax/parse.go:1765
				// _ = "end of CoverTab[63889]"
//line /usr/local/go/src/regexp/syntax/parse.go:1765
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1765
			// _ = "end of CoverTab[63886]"
//line /usr/local/go/src/regexp/syntax/parse.go:1765
			_go_fuzz_dep_.CoverTab[63887]++
									if nclass != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1766
				_go_fuzz_dep_.CoverTab[63890]++
										class, t = nclass, nt
										continue
//line /usr/local/go/src/regexp/syntax/parse.go:1768
				// _ = "end of CoverTab[63890]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1769
				_go_fuzz_dep_.CoverTab[63891]++
//line /usr/local/go/src/regexp/syntax/parse.go:1769
				// _ = "end of CoverTab[63891]"
//line /usr/local/go/src/regexp/syntax/parse.go:1769
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1769
			// _ = "end of CoverTab[63887]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1770
			_go_fuzz_dep_.CoverTab[63892]++
//line /usr/local/go/src/regexp/syntax/parse.go:1770
			// _ = "end of CoverTab[63892]"
//line /usr/local/go/src/regexp/syntax/parse.go:1770
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1770
		// _ = "end of CoverTab[63870]"
//line /usr/local/go/src/regexp/syntax/parse.go:1770
		_go_fuzz_dep_.CoverTab[63871]++

//line /usr/local/go/src/regexp/syntax/parse.go:1773
		nclass, nt, err := p.parseUnicodeClass(t, class)
		if err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1774
			_go_fuzz_dep_.CoverTab[63893]++
									return "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1775
			// _ = "end of CoverTab[63893]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1776
			_go_fuzz_dep_.CoverTab[63894]++
//line /usr/local/go/src/regexp/syntax/parse.go:1776
			// _ = "end of CoverTab[63894]"
//line /usr/local/go/src/regexp/syntax/parse.go:1776
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1776
		// _ = "end of CoverTab[63871]"
//line /usr/local/go/src/regexp/syntax/parse.go:1776
		_go_fuzz_dep_.CoverTab[63872]++
								if nclass != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1777
			_go_fuzz_dep_.CoverTab[63895]++
									class, t = nclass, nt
									continue
//line /usr/local/go/src/regexp/syntax/parse.go:1779
			// _ = "end of CoverTab[63895]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1780
			_go_fuzz_dep_.CoverTab[63896]++
//line /usr/local/go/src/regexp/syntax/parse.go:1780
			// _ = "end of CoverTab[63896]"
//line /usr/local/go/src/regexp/syntax/parse.go:1780
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1780
		// _ = "end of CoverTab[63872]"
//line /usr/local/go/src/regexp/syntax/parse.go:1780
		_go_fuzz_dep_.CoverTab[63873]++

//line /usr/local/go/src/regexp/syntax/parse.go:1783
		if nclass, nt := p.parsePerlClassEscape(t, class); nclass != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1783
			_go_fuzz_dep_.CoverTab[63897]++
									class, t = nclass, nt
									continue
//line /usr/local/go/src/regexp/syntax/parse.go:1785
			// _ = "end of CoverTab[63897]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1786
			_go_fuzz_dep_.CoverTab[63898]++
//line /usr/local/go/src/regexp/syntax/parse.go:1786
			// _ = "end of CoverTab[63898]"
//line /usr/local/go/src/regexp/syntax/parse.go:1786
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1786
		// _ = "end of CoverTab[63873]"
//line /usr/local/go/src/regexp/syntax/parse.go:1786
		_go_fuzz_dep_.CoverTab[63874]++

//line /usr/local/go/src/regexp/syntax/parse.go:1789
		rng := t
		var lo, hi rune
		if lo, t, err = p.parseClassChar(t, s); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1791
			_go_fuzz_dep_.CoverTab[63899]++
									return "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1792
			// _ = "end of CoverTab[63899]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1793
			_go_fuzz_dep_.CoverTab[63900]++
//line /usr/local/go/src/regexp/syntax/parse.go:1793
			// _ = "end of CoverTab[63900]"
//line /usr/local/go/src/regexp/syntax/parse.go:1793
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1793
		// _ = "end of CoverTab[63874]"
//line /usr/local/go/src/regexp/syntax/parse.go:1793
		_go_fuzz_dep_.CoverTab[63875]++
								hi = lo

								if len(t) >= 2 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1796
			_go_fuzz_dep_.CoverTab[63901]++
//line /usr/local/go/src/regexp/syntax/parse.go:1796
			return t[0] == '-'
//line /usr/local/go/src/regexp/syntax/parse.go:1796
			// _ = "end of CoverTab[63901]"
//line /usr/local/go/src/regexp/syntax/parse.go:1796
		}() && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1796
			_go_fuzz_dep_.CoverTab[63902]++
//line /usr/local/go/src/regexp/syntax/parse.go:1796
			return t[1] != ']'
//line /usr/local/go/src/regexp/syntax/parse.go:1796
			// _ = "end of CoverTab[63902]"
//line /usr/local/go/src/regexp/syntax/parse.go:1796
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1796
			_go_fuzz_dep_.CoverTab[63903]++
									t = t[1:]
									if hi, t, err = p.parseClassChar(t, s); err != nil {
//line /usr/local/go/src/regexp/syntax/parse.go:1798
				_go_fuzz_dep_.CoverTab[63905]++
										return "", err
//line /usr/local/go/src/regexp/syntax/parse.go:1799
				// _ = "end of CoverTab[63905]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1800
				_go_fuzz_dep_.CoverTab[63906]++
//line /usr/local/go/src/regexp/syntax/parse.go:1800
				// _ = "end of CoverTab[63906]"
//line /usr/local/go/src/regexp/syntax/parse.go:1800
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1800
			// _ = "end of CoverTab[63903]"
//line /usr/local/go/src/regexp/syntax/parse.go:1800
			_go_fuzz_dep_.CoverTab[63904]++
									if hi < lo {
//line /usr/local/go/src/regexp/syntax/parse.go:1801
				_go_fuzz_dep_.CoverTab[63907]++
										rng = rng[:len(rng)-len(t)]
										return "", &Error{Code: ErrInvalidCharRange, Expr: rng}
//line /usr/local/go/src/regexp/syntax/parse.go:1803
				// _ = "end of CoverTab[63907]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1804
				_go_fuzz_dep_.CoverTab[63908]++
//line /usr/local/go/src/regexp/syntax/parse.go:1804
				// _ = "end of CoverTab[63908]"
//line /usr/local/go/src/regexp/syntax/parse.go:1804
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1804
			// _ = "end of CoverTab[63904]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1805
			_go_fuzz_dep_.CoverTab[63909]++
//line /usr/local/go/src/regexp/syntax/parse.go:1805
			// _ = "end of CoverTab[63909]"
//line /usr/local/go/src/regexp/syntax/parse.go:1805
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1805
		// _ = "end of CoverTab[63875]"
//line /usr/local/go/src/regexp/syntax/parse.go:1805
		_go_fuzz_dep_.CoverTab[63876]++
								if p.flags&FoldCase == 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1806
			_go_fuzz_dep_.CoverTab[63910]++
									class = appendRange(class, lo, hi)
//line /usr/local/go/src/regexp/syntax/parse.go:1807
			// _ = "end of CoverTab[63910]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1808
			_go_fuzz_dep_.CoverTab[63911]++
									class = appendFoldedRange(class, lo, hi)
//line /usr/local/go/src/regexp/syntax/parse.go:1809
			// _ = "end of CoverTab[63911]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1810
		// _ = "end of CoverTab[63876]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1811
	// _ = "end of CoverTab[63859]"
//line /usr/local/go/src/regexp/syntax/parse.go:1811
	_go_fuzz_dep_.CoverTab[63860]++
							t = t[1:]

//line /usr/local/go/src/regexp/syntax/parse.go:1815
	re.Rune = class
	class = cleanClass(&re.Rune)
	if sign < 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1817
		_go_fuzz_dep_.CoverTab[63912]++
								class = negateClass(class)
//line /usr/local/go/src/regexp/syntax/parse.go:1818
		// _ = "end of CoverTab[63912]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1819
		_go_fuzz_dep_.CoverTab[63913]++
//line /usr/local/go/src/regexp/syntax/parse.go:1819
		// _ = "end of CoverTab[63913]"
//line /usr/local/go/src/regexp/syntax/parse.go:1819
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1819
	// _ = "end of CoverTab[63860]"
//line /usr/local/go/src/regexp/syntax/parse.go:1819
	_go_fuzz_dep_.CoverTab[63861]++
							re.Rune = class
							p.push(re)
							return t, nil
//line /usr/local/go/src/regexp/syntax/parse.go:1822
	// _ = "end of CoverTab[63861]"
}

// cleanClass sorts the ranges (pairs of elements of r),
//line /usr/local/go/src/regexp/syntax/parse.go:1825
// merges them, and eliminates duplicates.
//line /usr/local/go/src/regexp/syntax/parse.go:1827
func cleanClass(rp *[]rune) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1827
	_go_fuzz_dep_.CoverTab[63914]++

//line /usr/local/go/src/regexp/syntax/parse.go:1830
	sort.Sort(ranges{rp})

	r := *rp
	if len(r) < 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:1833
		_go_fuzz_dep_.CoverTab[63917]++
								return r
//line /usr/local/go/src/regexp/syntax/parse.go:1834
		// _ = "end of CoverTab[63917]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1835
		_go_fuzz_dep_.CoverTab[63918]++
//line /usr/local/go/src/regexp/syntax/parse.go:1835
		// _ = "end of CoverTab[63918]"
//line /usr/local/go/src/regexp/syntax/parse.go:1835
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1835
	// _ = "end of CoverTab[63914]"
//line /usr/local/go/src/regexp/syntax/parse.go:1835
	_go_fuzz_dep_.CoverTab[63915]++

//line /usr/local/go/src/regexp/syntax/parse.go:1838
	w := 2
	for i := 2; i < len(r); i += 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:1839
		_go_fuzz_dep_.CoverTab[63919]++
								lo, hi := r[i], r[i+1]
								if lo <= r[w-1]+1 {
//line /usr/local/go/src/regexp/syntax/parse.go:1841
			_go_fuzz_dep_.CoverTab[63921]++

									if hi > r[w-1] {
//line /usr/local/go/src/regexp/syntax/parse.go:1843
				_go_fuzz_dep_.CoverTab[63923]++
										r[w-1] = hi
//line /usr/local/go/src/regexp/syntax/parse.go:1844
				// _ = "end of CoverTab[63923]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1845
				_go_fuzz_dep_.CoverTab[63924]++
//line /usr/local/go/src/regexp/syntax/parse.go:1845
				// _ = "end of CoverTab[63924]"
//line /usr/local/go/src/regexp/syntax/parse.go:1845
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1845
			// _ = "end of CoverTab[63921]"
//line /usr/local/go/src/regexp/syntax/parse.go:1845
			_go_fuzz_dep_.CoverTab[63922]++
									continue
//line /usr/local/go/src/regexp/syntax/parse.go:1846
			// _ = "end of CoverTab[63922]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1847
			_go_fuzz_dep_.CoverTab[63925]++
//line /usr/local/go/src/regexp/syntax/parse.go:1847
			// _ = "end of CoverTab[63925]"
//line /usr/local/go/src/regexp/syntax/parse.go:1847
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1847
		// _ = "end of CoverTab[63919]"
//line /usr/local/go/src/regexp/syntax/parse.go:1847
		_go_fuzz_dep_.CoverTab[63920]++

								r[w] = lo
								r[w+1] = hi
								w += 2
//line /usr/local/go/src/regexp/syntax/parse.go:1851
		// _ = "end of CoverTab[63920]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1852
	// _ = "end of CoverTab[63915]"
//line /usr/local/go/src/regexp/syntax/parse.go:1852
	_go_fuzz_dep_.CoverTab[63916]++

							return r[:w]
//line /usr/local/go/src/regexp/syntax/parse.go:1854
	// _ = "end of CoverTab[63916]"
}

// appendLiteral returns the result of appending the literal x to the class r.
func appendLiteral(r []rune, x rune, flags Flags) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1858
	_go_fuzz_dep_.CoverTab[63926]++
							if flags&FoldCase != 0 {
//line /usr/local/go/src/regexp/syntax/parse.go:1859
		_go_fuzz_dep_.CoverTab[63928]++
								return appendFoldedRange(r, x, x)
//line /usr/local/go/src/regexp/syntax/parse.go:1860
		// _ = "end of CoverTab[63928]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1861
		_go_fuzz_dep_.CoverTab[63929]++
//line /usr/local/go/src/regexp/syntax/parse.go:1861
		// _ = "end of CoverTab[63929]"
//line /usr/local/go/src/regexp/syntax/parse.go:1861
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1861
	// _ = "end of CoverTab[63926]"
//line /usr/local/go/src/regexp/syntax/parse.go:1861
	_go_fuzz_dep_.CoverTab[63927]++
							return appendRange(r, x, x)
//line /usr/local/go/src/regexp/syntax/parse.go:1862
	// _ = "end of CoverTab[63927]"
}

// appendRange returns the result of appending the range lo-hi to the class r.
func appendRange(r []rune, lo, hi rune) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1866
	_go_fuzz_dep_.CoverTab[63930]++

//line /usr/local/go/src/regexp/syntax/parse.go:1871
	n := len(r)
	for i := 2; i <= 4; i += 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:1872
		_go_fuzz_dep_.CoverTab[63932]++
								if n >= i {
//line /usr/local/go/src/regexp/syntax/parse.go:1873
			_go_fuzz_dep_.CoverTab[63933]++
									rlo, rhi := r[n-i], r[n-i+1]
									if lo <= rhi+1 && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1875
				_go_fuzz_dep_.CoverTab[63934]++
//line /usr/local/go/src/regexp/syntax/parse.go:1875
				return rlo <= hi+1
//line /usr/local/go/src/regexp/syntax/parse.go:1875
				// _ = "end of CoverTab[63934]"
//line /usr/local/go/src/regexp/syntax/parse.go:1875
			}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1875
				_go_fuzz_dep_.CoverTab[63935]++
										if lo < rlo {
//line /usr/local/go/src/regexp/syntax/parse.go:1876
					_go_fuzz_dep_.CoverTab[63938]++
											r[n-i] = lo
//line /usr/local/go/src/regexp/syntax/parse.go:1877
					// _ = "end of CoverTab[63938]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1878
					_go_fuzz_dep_.CoverTab[63939]++
//line /usr/local/go/src/regexp/syntax/parse.go:1878
					// _ = "end of CoverTab[63939]"
//line /usr/local/go/src/regexp/syntax/parse.go:1878
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1878
				// _ = "end of CoverTab[63935]"
//line /usr/local/go/src/regexp/syntax/parse.go:1878
				_go_fuzz_dep_.CoverTab[63936]++
										if hi > rhi {
//line /usr/local/go/src/regexp/syntax/parse.go:1879
					_go_fuzz_dep_.CoverTab[63940]++
											r[n-i+1] = hi
//line /usr/local/go/src/regexp/syntax/parse.go:1880
					// _ = "end of CoverTab[63940]"
				} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1881
					_go_fuzz_dep_.CoverTab[63941]++
//line /usr/local/go/src/regexp/syntax/parse.go:1881
					// _ = "end of CoverTab[63941]"
//line /usr/local/go/src/regexp/syntax/parse.go:1881
				}
//line /usr/local/go/src/regexp/syntax/parse.go:1881
				// _ = "end of CoverTab[63936]"
//line /usr/local/go/src/regexp/syntax/parse.go:1881
				_go_fuzz_dep_.CoverTab[63937]++
										return r
//line /usr/local/go/src/regexp/syntax/parse.go:1882
				// _ = "end of CoverTab[63937]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1883
				_go_fuzz_dep_.CoverTab[63942]++
//line /usr/local/go/src/regexp/syntax/parse.go:1883
				// _ = "end of CoverTab[63942]"
//line /usr/local/go/src/regexp/syntax/parse.go:1883
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1883
			// _ = "end of CoverTab[63933]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1884
			_go_fuzz_dep_.CoverTab[63943]++
//line /usr/local/go/src/regexp/syntax/parse.go:1884
			// _ = "end of CoverTab[63943]"
//line /usr/local/go/src/regexp/syntax/parse.go:1884
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1884
		// _ = "end of CoverTab[63932]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1885
	// _ = "end of CoverTab[63930]"
//line /usr/local/go/src/regexp/syntax/parse.go:1885
	_go_fuzz_dep_.CoverTab[63931]++

							return append(r, lo, hi)
//line /usr/local/go/src/regexp/syntax/parse.go:1887
	// _ = "end of CoverTab[63931]"
}

const (
	// minimum and maximum runes involved in folding.
	// checked during test.
	minFold	= 0x0041
	maxFold	= 0x1e943
)

// appendFoldedRange returns the result of appending the range lo-hi
//line /usr/local/go/src/regexp/syntax/parse.go:1897
// and its case folding-equivalent runes to the class r.
//line /usr/local/go/src/regexp/syntax/parse.go:1899
func appendFoldedRange(r []rune, lo, hi rune) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1899
	_go_fuzz_dep_.CoverTab[63944]++

							if lo <= minFold && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1901
		_go_fuzz_dep_.CoverTab[63950]++
//line /usr/local/go/src/regexp/syntax/parse.go:1901
		return hi >= maxFold
//line /usr/local/go/src/regexp/syntax/parse.go:1901
		// _ = "end of CoverTab[63950]"
//line /usr/local/go/src/regexp/syntax/parse.go:1901
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1901
		_go_fuzz_dep_.CoverTab[63951]++

								return appendRange(r, lo, hi)
//line /usr/local/go/src/regexp/syntax/parse.go:1903
		// _ = "end of CoverTab[63951]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1904
		_go_fuzz_dep_.CoverTab[63952]++
//line /usr/local/go/src/regexp/syntax/parse.go:1904
		// _ = "end of CoverTab[63952]"
//line /usr/local/go/src/regexp/syntax/parse.go:1904
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1904
	// _ = "end of CoverTab[63944]"
//line /usr/local/go/src/regexp/syntax/parse.go:1904
	_go_fuzz_dep_.CoverTab[63945]++
							if hi < minFold || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:1905
		_go_fuzz_dep_.CoverTab[63953]++
//line /usr/local/go/src/regexp/syntax/parse.go:1905
		return lo > maxFold
//line /usr/local/go/src/regexp/syntax/parse.go:1905
		// _ = "end of CoverTab[63953]"
//line /usr/local/go/src/regexp/syntax/parse.go:1905
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:1905
		_go_fuzz_dep_.CoverTab[63954]++

								return appendRange(r, lo, hi)
//line /usr/local/go/src/regexp/syntax/parse.go:1907
		// _ = "end of CoverTab[63954]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1908
		_go_fuzz_dep_.CoverTab[63955]++
//line /usr/local/go/src/regexp/syntax/parse.go:1908
		// _ = "end of CoverTab[63955]"
//line /usr/local/go/src/regexp/syntax/parse.go:1908
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1908
	// _ = "end of CoverTab[63945]"
//line /usr/local/go/src/regexp/syntax/parse.go:1908
	_go_fuzz_dep_.CoverTab[63946]++
							if lo < minFold {
//line /usr/local/go/src/regexp/syntax/parse.go:1909
		_go_fuzz_dep_.CoverTab[63956]++

								r = appendRange(r, lo, minFold-1)
								lo = minFold
//line /usr/local/go/src/regexp/syntax/parse.go:1912
		// _ = "end of CoverTab[63956]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1913
		_go_fuzz_dep_.CoverTab[63957]++
//line /usr/local/go/src/regexp/syntax/parse.go:1913
		// _ = "end of CoverTab[63957]"
//line /usr/local/go/src/regexp/syntax/parse.go:1913
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1913
	// _ = "end of CoverTab[63946]"
//line /usr/local/go/src/regexp/syntax/parse.go:1913
	_go_fuzz_dep_.CoverTab[63947]++
							if hi > maxFold {
//line /usr/local/go/src/regexp/syntax/parse.go:1914
		_go_fuzz_dep_.CoverTab[63958]++

								r = appendRange(r, maxFold+1, hi)
								hi = maxFold
//line /usr/local/go/src/regexp/syntax/parse.go:1917
		// _ = "end of CoverTab[63958]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1918
		_go_fuzz_dep_.CoverTab[63959]++
//line /usr/local/go/src/regexp/syntax/parse.go:1918
		// _ = "end of CoverTab[63959]"
//line /usr/local/go/src/regexp/syntax/parse.go:1918
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1918
	// _ = "end of CoverTab[63947]"
//line /usr/local/go/src/regexp/syntax/parse.go:1918
	_go_fuzz_dep_.CoverTab[63948]++

//line /usr/local/go/src/regexp/syntax/parse.go:1921
	for c := lo; c <= hi; c++ {
//line /usr/local/go/src/regexp/syntax/parse.go:1921
		_go_fuzz_dep_.CoverTab[63960]++
								r = appendRange(r, c, c)
								f := unicode.SimpleFold(c)
								for f != c {
//line /usr/local/go/src/regexp/syntax/parse.go:1924
			_go_fuzz_dep_.CoverTab[63961]++
									r = appendRange(r, f, f)
									f = unicode.SimpleFold(f)
//line /usr/local/go/src/regexp/syntax/parse.go:1926
			// _ = "end of CoverTab[63961]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1927
		// _ = "end of CoverTab[63960]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1928
	// _ = "end of CoverTab[63948]"
//line /usr/local/go/src/regexp/syntax/parse.go:1928
	_go_fuzz_dep_.CoverTab[63949]++
							return r
//line /usr/local/go/src/regexp/syntax/parse.go:1929
	// _ = "end of CoverTab[63949]"
}

// appendClass returns the result of appending the class x to the class r.
//line /usr/local/go/src/regexp/syntax/parse.go:1932
// It assume x is clean.
//line /usr/local/go/src/regexp/syntax/parse.go:1934
func appendClass(r []rune, x []rune) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1934
	_go_fuzz_dep_.CoverTab[63962]++
							for i := 0; i < len(x); i += 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:1935
		_go_fuzz_dep_.CoverTab[63964]++
								r = appendRange(r, x[i], x[i+1])
//line /usr/local/go/src/regexp/syntax/parse.go:1936
		// _ = "end of CoverTab[63964]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1937
	// _ = "end of CoverTab[63962]"
//line /usr/local/go/src/regexp/syntax/parse.go:1937
	_go_fuzz_dep_.CoverTab[63963]++
							return r
//line /usr/local/go/src/regexp/syntax/parse.go:1938
	// _ = "end of CoverTab[63963]"
}

// appendFoldedClass returns the result of appending the case folding of the class x to the class r.
func appendFoldedClass(r []rune, x []rune) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1942
	_go_fuzz_dep_.CoverTab[63965]++
							for i := 0; i < len(x); i += 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:1943
		_go_fuzz_dep_.CoverTab[63967]++
								r = appendFoldedRange(r, x[i], x[i+1])
//line /usr/local/go/src/regexp/syntax/parse.go:1944
		// _ = "end of CoverTab[63967]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1945
	// _ = "end of CoverTab[63965]"
//line /usr/local/go/src/regexp/syntax/parse.go:1945
	_go_fuzz_dep_.CoverTab[63966]++
							return r
//line /usr/local/go/src/regexp/syntax/parse.go:1946
	// _ = "end of CoverTab[63966]"
}

// appendNegatedClass returns the result of appending the negation of the class x to the class r.
//line /usr/local/go/src/regexp/syntax/parse.go:1949
// It assumes x is clean.
//line /usr/local/go/src/regexp/syntax/parse.go:1951
func appendNegatedClass(r []rune, x []rune) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1951
	_go_fuzz_dep_.CoverTab[63968]++
							nextLo := '\u0000'
							for i := 0; i < len(x); i += 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:1953
		_go_fuzz_dep_.CoverTab[63971]++
								lo, hi := x[i], x[i+1]
								if nextLo <= lo-1 {
//line /usr/local/go/src/regexp/syntax/parse.go:1955
			_go_fuzz_dep_.CoverTab[63973]++
									r = appendRange(r, nextLo, lo-1)
//line /usr/local/go/src/regexp/syntax/parse.go:1956
			// _ = "end of CoverTab[63973]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1957
			_go_fuzz_dep_.CoverTab[63974]++
//line /usr/local/go/src/regexp/syntax/parse.go:1957
			// _ = "end of CoverTab[63974]"
//line /usr/local/go/src/regexp/syntax/parse.go:1957
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1957
		// _ = "end of CoverTab[63971]"
//line /usr/local/go/src/regexp/syntax/parse.go:1957
		_go_fuzz_dep_.CoverTab[63972]++
								nextLo = hi + 1
//line /usr/local/go/src/regexp/syntax/parse.go:1958
		// _ = "end of CoverTab[63972]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1959
	// _ = "end of CoverTab[63968]"
//line /usr/local/go/src/regexp/syntax/parse.go:1959
	_go_fuzz_dep_.CoverTab[63969]++
							if nextLo <= unicode.MaxRune {
//line /usr/local/go/src/regexp/syntax/parse.go:1960
		_go_fuzz_dep_.CoverTab[63975]++
								r = appendRange(r, nextLo, unicode.MaxRune)
//line /usr/local/go/src/regexp/syntax/parse.go:1961
		// _ = "end of CoverTab[63975]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1962
		_go_fuzz_dep_.CoverTab[63976]++
//line /usr/local/go/src/regexp/syntax/parse.go:1962
		// _ = "end of CoverTab[63976]"
//line /usr/local/go/src/regexp/syntax/parse.go:1962
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1962
	// _ = "end of CoverTab[63969]"
//line /usr/local/go/src/regexp/syntax/parse.go:1962
	_go_fuzz_dep_.CoverTab[63970]++
							return r
//line /usr/local/go/src/regexp/syntax/parse.go:1963
	// _ = "end of CoverTab[63970]"
}

// appendTable returns the result of appending x to the class r.
func appendTable(r []rune, x *unicode.RangeTable) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1967
	_go_fuzz_dep_.CoverTab[63977]++
							for _, xr := range x.R16 {
//line /usr/local/go/src/regexp/syntax/parse.go:1968
		_go_fuzz_dep_.CoverTab[63980]++
								lo, hi, stride := rune(xr.Lo), rune(xr.Hi), rune(xr.Stride)
								if stride == 1 {
//line /usr/local/go/src/regexp/syntax/parse.go:1970
			_go_fuzz_dep_.CoverTab[63982]++
									r = appendRange(r, lo, hi)
									continue
//line /usr/local/go/src/regexp/syntax/parse.go:1972
			// _ = "end of CoverTab[63982]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1973
			_go_fuzz_dep_.CoverTab[63983]++
//line /usr/local/go/src/regexp/syntax/parse.go:1973
			// _ = "end of CoverTab[63983]"
//line /usr/local/go/src/regexp/syntax/parse.go:1973
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1973
		// _ = "end of CoverTab[63980]"
//line /usr/local/go/src/regexp/syntax/parse.go:1973
		_go_fuzz_dep_.CoverTab[63981]++
								for c := lo; c <= hi; c += stride {
//line /usr/local/go/src/regexp/syntax/parse.go:1974
			_go_fuzz_dep_.CoverTab[63984]++
									r = appendRange(r, c, c)
//line /usr/local/go/src/regexp/syntax/parse.go:1975
			// _ = "end of CoverTab[63984]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1976
		// _ = "end of CoverTab[63981]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1977
	// _ = "end of CoverTab[63977]"
//line /usr/local/go/src/regexp/syntax/parse.go:1977
	_go_fuzz_dep_.CoverTab[63978]++
							for _, xr := range x.R32 {
//line /usr/local/go/src/regexp/syntax/parse.go:1978
		_go_fuzz_dep_.CoverTab[63985]++
								lo, hi, stride := rune(xr.Lo), rune(xr.Hi), rune(xr.Stride)
								if stride == 1 {
//line /usr/local/go/src/regexp/syntax/parse.go:1980
			_go_fuzz_dep_.CoverTab[63987]++
									r = appendRange(r, lo, hi)
									continue
//line /usr/local/go/src/regexp/syntax/parse.go:1982
			// _ = "end of CoverTab[63987]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1983
			_go_fuzz_dep_.CoverTab[63988]++
//line /usr/local/go/src/regexp/syntax/parse.go:1983
			// _ = "end of CoverTab[63988]"
//line /usr/local/go/src/regexp/syntax/parse.go:1983
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1983
		// _ = "end of CoverTab[63985]"
//line /usr/local/go/src/regexp/syntax/parse.go:1983
		_go_fuzz_dep_.CoverTab[63986]++
								for c := lo; c <= hi; c += stride {
//line /usr/local/go/src/regexp/syntax/parse.go:1984
			_go_fuzz_dep_.CoverTab[63989]++
									r = appendRange(r, c, c)
//line /usr/local/go/src/regexp/syntax/parse.go:1985
			// _ = "end of CoverTab[63989]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:1986
		// _ = "end of CoverTab[63986]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:1987
	// _ = "end of CoverTab[63978]"
//line /usr/local/go/src/regexp/syntax/parse.go:1987
	_go_fuzz_dep_.CoverTab[63979]++
							return r
//line /usr/local/go/src/regexp/syntax/parse.go:1988
	// _ = "end of CoverTab[63979]"
}

// appendNegatedTable returns the result of appending the negation of x to the class r.
func appendNegatedTable(r []rune, x *unicode.RangeTable) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:1992
	_go_fuzz_dep_.CoverTab[63990]++
							nextLo := '\u0000'
							for _, xr := range x.R16 {
//line /usr/local/go/src/regexp/syntax/parse.go:1994
		_go_fuzz_dep_.CoverTab[63994]++
								lo, hi, stride := rune(xr.Lo), rune(xr.Hi), rune(xr.Stride)
								if stride == 1 {
//line /usr/local/go/src/regexp/syntax/parse.go:1996
			_go_fuzz_dep_.CoverTab[63996]++
									if nextLo <= lo-1 {
//line /usr/local/go/src/regexp/syntax/parse.go:1997
				_go_fuzz_dep_.CoverTab[63998]++
										r = appendRange(r, nextLo, lo-1)
//line /usr/local/go/src/regexp/syntax/parse.go:1998
				// _ = "end of CoverTab[63998]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:1999
				_go_fuzz_dep_.CoverTab[63999]++
//line /usr/local/go/src/regexp/syntax/parse.go:1999
				// _ = "end of CoverTab[63999]"
//line /usr/local/go/src/regexp/syntax/parse.go:1999
			}
//line /usr/local/go/src/regexp/syntax/parse.go:1999
			// _ = "end of CoverTab[63996]"
//line /usr/local/go/src/regexp/syntax/parse.go:1999
			_go_fuzz_dep_.CoverTab[63997]++
									nextLo = hi + 1
									continue
//line /usr/local/go/src/regexp/syntax/parse.go:2001
			// _ = "end of CoverTab[63997]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2002
			_go_fuzz_dep_.CoverTab[64000]++
//line /usr/local/go/src/regexp/syntax/parse.go:2002
			// _ = "end of CoverTab[64000]"
//line /usr/local/go/src/regexp/syntax/parse.go:2002
		}
//line /usr/local/go/src/regexp/syntax/parse.go:2002
		// _ = "end of CoverTab[63994]"
//line /usr/local/go/src/regexp/syntax/parse.go:2002
		_go_fuzz_dep_.CoverTab[63995]++
								for c := lo; c <= hi; c += stride {
//line /usr/local/go/src/regexp/syntax/parse.go:2003
			_go_fuzz_dep_.CoverTab[64001]++
									if nextLo <= c-1 {
//line /usr/local/go/src/regexp/syntax/parse.go:2004
				_go_fuzz_dep_.CoverTab[64003]++
										r = appendRange(r, nextLo, c-1)
//line /usr/local/go/src/regexp/syntax/parse.go:2005
				// _ = "end of CoverTab[64003]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2006
				_go_fuzz_dep_.CoverTab[64004]++
//line /usr/local/go/src/regexp/syntax/parse.go:2006
				// _ = "end of CoverTab[64004]"
//line /usr/local/go/src/regexp/syntax/parse.go:2006
			}
//line /usr/local/go/src/regexp/syntax/parse.go:2006
			// _ = "end of CoverTab[64001]"
//line /usr/local/go/src/regexp/syntax/parse.go:2006
			_go_fuzz_dep_.CoverTab[64002]++
									nextLo = c + 1
//line /usr/local/go/src/regexp/syntax/parse.go:2007
			// _ = "end of CoverTab[64002]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:2008
		// _ = "end of CoverTab[63995]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2009
	// _ = "end of CoverTab[63990]"
//line /usr/local/go/src/regexp/syntax/parse.go:2009
	_go_fuzz_dep_.CoverTab[63991]++
							for _, xr := range x.R32 {
//line /usr/local/go/src/regexp/syntax/parse.go:2010
		_go_fuzz_dep_.CoverTab[64005]++
								lo, hi, stride := rune(xr.Lo), rune(xr.Hi), rune(xr.Stride)
								if stride == 1 {
//line /usr/local/go/src/regexp/syntax/parse.go:2012
			_go_fuzz_dep_.CoverTab[64007]++
									if nextLo <= lo-1 {
//line /usr/local/go/src/regexp/syntax/parse.go:2013
				_go_fuzz_dep_.CoverTab[64009]++
										r = appendRange(r, nextLo, lo-1)
//line /usr/local/go/src/regexp/syntax/parse.go:2014
				// _ = "end of CoverTab[64009]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2015
				_go_fuzz_dep_.CoverTab[64010]++
//line /usr/local/go/src/regexp/syntax/parse.go:2015
				// _ = "end of CoverTab[64010]"
//line /usr/local/go/src/regexp/syntax/parse.go:2015
			}
//line /usr/local/go/src/regexp/syntax/parse.go:2015
			// _ = "end of CoverTab[64007]"
//line /usr/local/go/src/regexp/syntax/parse.go:2015
			_go_fuzz_dep_.CoverTab[64008]++
									nextLo = hi + 1
									continue
//line /usr/local/go/src/regexp/syntax/parse.go:2017
			// _ = "end of CoverTab[64008]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2018
			_go_fuzz_dep_.CoverTab[64011]++
//line /usr/local/go/src/regexp/syntax/parse.go:2018
			// _ = "end of CoverTab[64011]"
//line /usr/local/go/src/regexp/syntax/parse.go:2018
		}
//line /usr/local/go/src/regexp/syntax/parse.go:2018
		// _ = "end of CoverTab[64005]"
//line /usr/local/go/src/regexp/syntax/parse.go:2018
		_go_fuzz_dep_.CoverTab[64006]++
								for c := lo; c <= hi; c += stride {
//line /usr/local/go/src/regexp/syntax/parse.go:2019
			_go_fuzz_dep_.CoverTab[64012]++
									if nextLo <= c-1 {
//line /usr/local/go/src/regexp/syntax/parse.go:2020
				_go_fuzz_dep_.CoverTab[64014]++
										r = appendRange(r, nextLo, c-1)
//line /usr/local/go/src/regexp/syntax/parse.go:2021
				// _ = "end of CoverTab[64014]"
			} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2022
				_go_fuzz_dep_.CoverTab[64015]++
//line /usr/local/go/src/regexp/syntax/parse.go:2022
				// _ = "end of CoverTab[64015]"
//line /usr/local/go/src/regexp/syntax/parse.go:2022
			}
//line /usr/local/go/src/regexp/syntax/parse.go:2022
			// _ = "end of CoverTab[64012]"
//line /usr/local/go/src/regexp/syntax/parse.go:2022
			_go_fuzz_dep_.CoverTab[64013]++
									nextLo = c + 1
//line /usr/local/go/src/regexp/syntax/parse.go:2023
			// _ = "end of CoverTab[64013]"
		}
//line /usr/local/go/src/regexp/syntax/parse.go:2024
		// _ = "end of CoverTab[64006]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2025
	// _ = "end of CoverTab[63991]"
//line /usr/local/go/src/regexp/syntax/parse.go:2025
	_go_fuzz_dep_.CoverTab[63992]++
							if nextLo <= unicode.MaxRune {
//line /usr/local/go/src/regexp/syntax/parse.go:2026
		_go_fuzz_dep_.CoverTab[64016]++
								r = appendRange(r, nextLo, unicode.MaxRune)
//line /usr/local/go/src/regexp/syntax/parse.go:2027
		// _ = "end of CoverTab[64016]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2028
		_go_fuzz_dep_.CoverTab[64017]++
//line /usr/local/go/src/regexp/syntax/parse.go:2028
		// _ = "end of CoverTab[64017]"
//line /usr/local/go/src/regexp/syntax/parse.go:2028
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2028
	// _ = "end of CoverTab[63992]"
//line /usr/local/go/src/regexp/syntax/parse.go:2028
	_go_fuzz_dep_.CoverTab[63993]++
							return r
//line /usr/local/go/src/regexp/syntax/parse.go:2029
	// _ = "end of CoverTab[63993]"
}

// negateClass overwrites r and returns r's negation.
//line /usr/local/go/src/regexp/syntax/parse.go:2032
// It assumes the class r is already clean.
//line /usr/local/go/src/regexp/syntax/parse.go:2034
func negateClass(r []rune) []rune {
//line /usr/local/go/src/regexp/syntax/parse.go:2034
	_go_fuzz_dep_.CoverTab[64018]++
							nextLo := '\u0000'
							w := 0
							for i := 0; i < len(r); i += 2 {
//line /usr/local/go/src/regexp/syntax/parse.go:2037
		_go_fuzz_dep_.CoverTab[64021]++
								lo, hi := r[i], r[i+1]
								if nextLo <= lo-1 {
//line /usr/local/go/src/regexp/syntax/parse.go:2039
			_go_fuzz_dep_.CoverTab[64023]++
									r[w] = nextLo
									r[w+1] = lo - 1
									w += 2
//line /usr/local/go/src/regexp/syntax/parse.go:2042
			// _ = "end of CoverTab[64023]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2043
			_go_fuzz_dep_.CoverTab[64024]++
//line /usr/local/go/src/regexp/syntax/parse.go:2043
			// _ = "end of CoverTab[64024]"
//line /usr/local/go/src/regexp/syntax/parse.go:2043
		}
//line /usr/local/go/src/regexp/syntax/parse.go:2043
		// _ = "end of CoverTab[64021]"
//line /usr/local/go/src/regexp/syntax/parse.go:2043
		_go_fuzz_dep_.CoverTab[64022]++
								nextLo = hi + 1
//line /usr/local/go/src/regexp/syntax/parse.go:2044
		// _ = "end of CoverTab[64022]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2045
	// _ = "end of CoverTab[64018]"
//line /usr/local/go/src/regexp/syntax/parse.go:2045
	_go_fuzz_dep_.CoverTab[64019]++
							r = r[:w]
							if nextLo <= unicode.MaxRune {
//line /usr/local/go/src/regexp/syntax/parse.go:2047
		_go_fuzz_dep_.CoverTab[64025]++

//line /usr/local/go/src/regexp/syntax/parse.go:2050
		r = append(r, nextLo, unicode.MaxRune)
//line /usr/local/go/src/regexp/syntax/parse.go:2050
		// _ = "end of CoverTab[64025]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2051
		_go_fuzz_dep_.CoverTab[64026]++
//line /usr/local/go/src/regexp/syntax/parse.go:2051
		// _ = "end of CoverTab[64026]"
//line /usr/local/go/src/regexp/syntax/parse.go:2051
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2051
	// _ = "end of CoverTab[64019]"
//line /usr/local/go/src/regexp/syntax/parse.go:2051
	_go_fuzz_dep_.CoverTab[64020]++
							return r
//line /usr/local/go/src/regexp/syntax/parse.go:2052
	// _ = "end of CoverTab[64020]"
}

// ranges implements sort.Interface on a []rune.
//line /usr/local/go/src/regexp/syntax/parse.go:2055
// The choice of receiver type definition is strange
//line /usr/local/go/src/regexp/syntax/parse.go:2055
// but avoids an allocation since we already have
//line /usr/local/go/src/regexp/syntax/parse.go:2055
// a *[]rune.
//line /usr/local/go/src/regexp/syntax/parse.go:2059
type ranges struct {
	p *[]rune
}

func (ra ranges) Less(i, j int) bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2063
	_go_fuzz_dep_.CoverTab[64027]++
							p := *ra.p
							i *= 2
							j *= 2
							return p[i] < p[j] || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2067
		_go_fuzz_dep_.CoverTab[64028]++
//line /usr/local/go/src/regexp/syntax/parse.go:2067
		return p[i] == p[j] && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2067
			_go_fuzz_dep_.CoverTab[64029]++
//line /usr/local/go/src/regexp/syntax/parse.go:2067
			return p[i+1] > p[j+1]
//line /usr/local/go/src/regexp/syntax/parse.go:2067
			// _ = "end of CoverTab[64029]"
//line /usr/local/go/src/regexp/syntax/parse.go:2067
		}()
//line /usr/local/go/src/regexp/syntax/parse.go:2067
		// _ = "end of CoverTab[64028]"
//line /usr/local/go/src/regexp/syntax/parse.go:2067
	}()
//line /usr/local/go/src/regexp/syntax/parse.go:2067
	// _ = "end of CoverTab[64027]"
}

func (ra ranges) Len() int {
//line /usr/local/go/src/regexp/syntax/parse.go:2070
	_go_fuzz_dep_.CoverTab[64030]++
							return len(*ra.p) / 2
//line /usr/local/go/src/regexp/syntax/parse.go:2071
	// _ = "end of CoverTab[64030]"
}

func (ra ranges) Swap(i, j int) {
//line /usr/local/go/src/regexp/syntax/parse.go:2074
	_go_fuzz_dep_.CoverTab[64031]++
							p := *ra.p
							i *= 2
							j *= 2
							p[i], p[i+1], p[j], p[j+1] = p[j], p[j+1], p[i], p[i+1]
//line /usr/local/go/src/regexp/syntax/parse.go:2078
	// _ = "end of CoverTab[64031]"
}

func checkUTF8(s string) error {
//line /usr/local/go/src/regexp/syntax/parse.go:2081
	_go_fuzz_dep_.CoverTab[64032]++
							for s != "" {
//line /usr/local/go/src/regexp/syntax/parse.go:2082
		_go_fuzz_dep_.CoverTab[64034]++
								rune, size := utf8.DecodeRuneInString(s)
								if rune == utf8.RuneError && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2084
			_go_fuzz_dep_.CoverTab[64036]++
//line /usr/local/go/src/regexp/syntax/parse.go:2084
			return size == 1
//line /usr/local/go/src/regexp/syntax/parse.go:2084
			// _ = "end of CoverTab[64036]"
//line /usr/local/go/src/regexp/syntax/parse.go:2084
		}() {
//line /usr/local/go/src/regexp/syntax/parse.go:2084
			_go_fuzz_dep_.CoverTab[64037]++
									return &Error{Code: ErrInvalidUTF8, Expr: s}
//line /usr/local/go/src/regexp/syntax/parse.go:2085
			// _ = "end of CoverTab[64037]"
		} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2086
			_go_fuzz_dep_.CoverTab[64038]++
//line /usr/local/go/src/regexp/syntax/parse.go:2086
			// _ = "end of CoverTab[64038]"
//line /usr/local/go/src/regexp/syntax/parse.go:2086
		}
//line /usr/local/go/src/regexp/syntax/parse.go:2086
		// _ = "end of CoverTab[64034]"
//line /usr/local/go/src/regexp/syntax/parse.go:2086
		_go_fuzz_dep_.CoverTab[64035]++
								s = s[size:]
//line /usr/local/go/src/regexp/syntax/parse.go:2087
		// _ = "end of CoverTab[64035]"
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2088
	// _ = "end of CoverTab[64032]"
//line /usr/local/go/src/regexp/syntax/parse.go:2088
	_go_fuzz_dep_.CoverTab[64033]++
							return nil
//line /usr/local/go/src/regexp/syntax/parse.go:2089
	// _ = "end of CoverTab[64033]"
}

func nextRune(s string) (c rune, t string, err error) {
//line /usr/local/go/src/regexp/syntax/parse.go:2092
	_go_fuzz_dep_.CoverTab[64039]++
							c, size := utf8.DecodeRuneInString(s)
							if c == utf8.RuneError && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2094
		_go_fuzz_dep_.CoverTab[64041]++
//line /usr/local/go/src/regexp/syntax/parse.go:2094
		return size == 1
//line /usr/local/go/src/regexp/syntax/parse.go:2094
		// _ = "end of CoverTab[64041]"
//line /usr/local/go/src/regexp/syntax/parse.go:2094
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:2094
		_go_fuzz_dep_.CoverTab[64042]++
								return 0, "", &Error{Code: ErrInvalidUTF8, Expr: s}
//line /usr/local/go/src/regexp/syntax/parse.go:2095
		// _ = "end of CoverTab[64042]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2096
		_go_fuzz_dep_.CoverTab[64043]++
//line /usr/local/go/src/regexp/syntax/parse.go:2096
		// _ = "end of CoverTab[64043]"
//line /usr/local/go/src/regexp/syntax/parse.go:2096
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2096
	// _ = "end of CoverTab[64039]"
//line /usr/local/go/src/regexp/syntax/parse.go:2096
	_go_fuzz_dep_.CoverTab[64040]++
							return c, s[size:], nil
//line /usr/local/go/src/regexp/syntax/parse.go:2097
	// _ = "end of CoverTab[64040]"
}

func isalnum(c rune) bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2100
	_go_fuzz_dep_.CoverTab[64044]++
							return '0' <= c && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		_go_fuzz_dep_.CoverTab[64045]++
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		return c <= '9'
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		// _ = "end of CoverTab[64045]"
//line /usr/local/go/src/regexp/syntax/parse.go:2101
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		_go_fuzz_dep_.CoverTab[64046]++
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		return 'A' <= c && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2101
			_go_fuzz_dep_.CoverTab[64047]++
//line /usr/local/go/src/regexp/syntax/parse.go:2101
			return c <= 'Z'
//line /usr/local/go/src/regexp/syntax/parse.go:2101
			// _ = "end of CoverTab[64047]"
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		}()
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		// _ = "end of CoverTab[64046]"
//line /usr/local/go/src/regexp/syntax/parse.go:2101
	}() || func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		_go_fuzz_dep_.CoverTab[64048]++
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		return 'a' <= c && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2101
			_go_fuzz_dep_.CoverTab[64049]++
//line /usr/local/go/src/regexp/syntax/parse.go:2101
			return c <= 'z'
//line /usr/local/go/src/regexp/syntax/parse.go:2101
			// _ = "end of CoverTab[64049]"
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		}()
//line /usr/local/go/src/regexp/syntax/parse.go:2101
		// _ = "end of CoverTab[64048]"
//line /usr/local/go/src/regexp/syntax/parse.go:2101
	}()
//line /usr/local/go/src/regexp/syntax/parse.go:2101
	// _ = "end of CoverTab[64044]"
}

func unhex(c rune) rune {
//line /usr/local/go/src/regexp/syntax/parse.go:2104
	_go_fuzz_dep_.CoverTab[64050]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2105
		_go_fuzz_dep_.CoverTab[64054]++
//line /usr/local/go/src/regexp/syntax/parse.go:2105
		return c <= '9'
//line /usr/local/go/src/regexp/syntax/parse.go:2105
		// _ = "end of CoverTab[64054]"
//line /usr/local/go/src/regexp/syntax/parse.go:2105
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:2105
		_go_fuzz_dep_.CoverTab[64055]++
								return c - '0'
//line /usr/local/go/src/regexp/syntax/parse.go:2106
		// _ = "end of CoverTab[64055]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2107
		_go_fuzz_dep_.CoverTab[64056]++
//line /usr/local/go/src/regexp/syntax/parse.go:2107
		// _ = "end of CoverTab[64056]"
//line /usr/local/go/src/regexp/syntax/parse.go:2107
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2107
	// _ = "end of CoverTab[64050]"
//line /usr/local/go/src/regexp/syntax/parse.go:2107
	_go_fuzz_dep_.CoverTab[64051]++
							if 'a' <= c && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2108
		_go_fuzz_dep_.CoverTab[64057]++
//line /usr/local/go/src/regexp/syntax/parse.go:2108
		return c <= 'f'
//line /usr/local/go/src/regexp/syntax/parse.go:2108
		// _ = "end of CoverTab[64057]"
//line /usr/local/go/src/regexp/syntax/parse.go:2108
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:2108
		_go_fuzz_dep_.CoverTab[64058]++
								return c - 'a' + 10
//line /usr/local/go/src/regexp/syntax/parse.go:2109
		// _ = "end of CoverTab[64058]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2110
		_go_fuzz_dep_.CoverTab[64059]++
//line /usr/local/go/src/regexp/syntax/parse.go:2110
		// _ = "end of CoverTab[64059]"
//line /usr/local/go/src/regexp/syntax/parse.go:2110
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2110
	// _ = "end of CoverTab[64051]"
//line /usr/local/go/src/regexp/syntax/parse.go:2110
	_go_fuzz_dep_.CoverTab[64052]++
							if 'A' <= c && func() bool {
//line /usr/local/go/src/regexp/syntax/parse.go:2111
		_go_fuzz_dep_.CoverTab[64060]++
//line /usr/local/go/src/regexp/syntax/parse.go:2111
		return c <= 'F'
//line /usr/local/go/src/regexp/syntax/parse.go:2111
		// _ = "end of CoverTab[64060]"
//line /usr/local/go/src/regexp/syntax/parse.go:2111
	}() {
//line /usr/local/go/src/regexp/syntax/parse.go:2111
		_go_fuzz_dep_.CoverTab[64061]++
								return c - 'A' + 10
//line /usr/local/go/src/regexp/syntax/parse.go:2112
		// _ = "end of CoverTab[64061]"
	} else {
//line /usr/local/go/src/regexp/syntax/parse.go:2113
		_go_fuzz_dep_.CoverTab[64062]++
//line /usr/local/go/src/regexp/syntax/parse.go:2113
		// _ = "end of CoverTab[64062]"
//line /usr/local/go/src/regexp/syntax/parse.go:2113
	}
//line /usr/local/go/src/regexp/syntax/parse.go:2113
	// _ = "end of CoverTab[64052]"
//line /usr/local/go/src/regexp/syntax/parse.go:2113
	_go_fuzz_dep_.CoverTab[64053]++
							return -1
//line /usr/local/go/src/regexp/syntax/parse.go:2114
	// _ = "end of CoverTab[64053]"
}

//line /usr/local/go/src/regexp/syntax/parse.go:2115
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/regexp/syntax/parse.go:2115
var _ = _go_fuzz_dep_.CoverTab
