// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/text/template/parse/lex.go:5
package parse

//line /usr/local/go/src/text/template/parse/lex.go:5
import (
//line /usr/local/go/src/text/template/parse/lex.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/text/template/parse/lex.go:5
)
//line /usr/local/go/src/text/template/parse/lex.go:5
import (
//line /usr/local/go/src/text/template/parse/lex.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/text/template/parse/lex.go:5
)

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// item represents a token or text string returned from the scanner.
type item struct {
	typ	itemType	// The type of this item.
	pos	Pos		// The starting position, in bytes, of this item in the input string.
	val	string		// The value of this item.
	line	int		// The line number at the start of this item.
}

func (i item) String() string {
//line /usr/local/go/src/text/template/parse/lex.go:22
	_go_fuzz_dep_.CoverTab[28914]++
							switch {
	case i.typ == itemEOF:
//line /usr/local/go/src/text/template/parse/lex.go:24
		_go_fuzz_dep_.CoverTab[28916]++
								return "EOF"
//line /usr/local/go/src/text/template/parse/lex.go:25
		// _ = "end of CoverTab[28916]"
	case i.typ == itemError:
//line /usr/local/go/src/text/template/parse/lex.go:26
		_go_fuzz_dep_.CoverTab[28917]++
								return i.val
//line /usr/local/go/src/text/template/parse/lex.go:27
		// _ = "end of CoverTab[28917]"
	case i.typ > itemKeyword:
//line /usr/local/go/src/text/template/parse/lex.go:28
		_go_fuzz_dep_.CoverTab[28918]++
								return fmt.Sprintf("<%s>", i.val)
//line /usr/local/go/src/text/template/parse/lex.go:29
		// _ = "end of CoverTab[28918]"
	case len(i.val) > 10:
//line /usr/local/go/src/text/template/parse/lex.go:30
		_go_fuzz_dep_.CoverTab[28919]++
								return fmt.Sprintf("%.10q...", i.val)
//line /usr/local/go/src/text/template/parse/lex.go:31
		// _ = "end of CoverTab[28919]"
//line /usr/local/go/src/text/template/parse/lex.go:31
	default:
//line /usr/local/go/src/text/template/parse/lex.go:31
		_go_fuzz_dep_.CoverTab[28920]++
//line /usr/local/go/src/text/template/parse/lex.go:31
		// _ = "end of CoverTab[28920]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:32
	// _ = "end of CoverTab[28914]"
//line /usr/local/go/src/text/template/parse/lex.go:32
	_go_fuzz_dep_.CoverTab[28915]++
							return fmt.Sprintf("%q", i.val)
//line /usr/local/go/src/text/template/parse/lex.go:33
	// _ = "end of CoverTab[28915]"
}

// itemType identifies the type of lex items.
type itemType int

const (
	itemError		itemType	= iota	// error occurred; value is text of error
	itemBool					// boolean constant
	itemChar					// printable ASCII character; grab bag for comma etc.
	itemCharConstant				// character constant
	itemComment					// comment text
	itemComplex					// complex constant (1+2i); imaginary is just a number
	itemAssign					// equals ('=') introducing an assignment
	itemDeclare					// colon-equals (':=') introducing a declaration
	itemEOF
	itemField	// alphanumeric identifier starting with '.'
	itemIdentifier	// alphanumeric identifier not starting with '.'
	itemLeftDelim	// left action delimiter
	itemLeftParen	// '(' inside action
	itemNumber	// simple number, including imaginary
	itemPipe	// pipe symbol
	itemRawString	// raw quoted string (includes quotes)
	itemRightDelim	// right action delimiter
	itemRightParen	// ')' inside action
	itemSpace	// run of spaces separating arguments
	itemString	// quoted string (includes quotes)
	itemText	// plain text
	itemVariable	// variable starting with '$', such as '$' or  '$1' or '$hello'
	// Keywords appear after all the rest.
	itemKeyword	// used only to delimit the keywords
	itemBlock	// block keyword
	itemBreak	// break keyword
	itemContinue	// continue keyword
	itemDot		// the cursor, spelled '.'
	itemDefine	// define keyword
	itemElse	// else keyword
	itemEnd		// end keyword
	itemIf		// if keyword
	itemNil		// the untyped nil constant, easiest to treat as a keyword
	itemRange	// range keyword
	itemTemplate	// template keyword
	itemWith	// with keyword
)

var key = map[string]itemType{
	".":		itemDot,
	"block":	itemBlock,
	"break":	itemBreak,
	"continue":	itemContinue,
	"define":	itemDefine,
	"else":		itemElse,
	"end":		itemEnd,
	"if":		itemIf,
	"range":	itemRange,
	"nil":		itemNil,
	"template":	itemTemplate,
	"with":		itemWith,
}

const eof = -1

// Trimming spaces.
//line /usr/local/go/src/text/template/parse/lex.go:95
// If the action begins "{{- " rather than "{{", then all space/tab/newlines
//line /usr/local/go/src/text/template/parse/lex.go:95
// preceding the action are trimmed; conversely if it ends " -}}" the
//line /usr/local/go/src/text/template/parse/lex.go:95
// leading spaces are trimmed. This is done entirely in the lexer; the
//line /usr/local/go/src/text/template/parse/lex.go:95
// parser never sees it happen. We require an ASCII space (' ', \t, \r, \n)
//line /usr/local/go/src/text/template/parse/lex.go:95
// to be present to avoid ambiguity with things like "{{-3}}". It reads
//line /usr/local/go/src/text/template/parse/lex.go:95
// better with the space present anyway. For simplicity, only ASCII
//line /usr/local/go/src/text/template/parse/lex.go:95
// does the job.
//line /usr/local/go/src/text/template/parse/lex.go:103
const (
	spaceChars	= " \t\r\n"	// These are the space characters defined by Go itself.
	trimMarker	= '-'		// Attached to left/right delimiter, trims trailing spaces from preceding/following text.
	trimMarkerLen	= Pos(1 + 1)	// marker plus space before or after
)

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	name		string	// the name of the input; used only for error reports
	input		string	// the string being scanned
	leftDelim	string	// start of action marker
	rightDelim	string	// end of action marker
	pos		Pos	// current position in the input
	start		Pos	// start position of this item
	atEOF		bool	// we have hit the end of input and returned eof
	parenDepth	int	// nesting depth of ( ) exprs
	line		int	// 1+number of newlines seen
	startLine	int	// start line of this item
	item		item	// item to return to parser
	insideAction	bool	// are we inside an action?
	options		lexOptions
}

// lexOptions control behavior of the lexer. All default to false.
type lexOptions struct {
	emitComment	bool	// emit itemComment tokens.
	breakOK		bool	// break keyword allowed
	continueOK	bool	// continue keyword allowed
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
//line /usr/local/go/src/text/template/parse/lex.go:137
	_go_fuzz_dep_.CoverTab[28921]++
								if int(l.pos) >= len(l.input) {
//line /usr/local/go/src/text/template/parse/lex.go:138
		_go_fuzz_dep_.CoverTab[28924]++
									l.atEOF = true
									return eof
//line /usr/local/go/src/text/template/parse/lex.go:140
		// _ = "end of CoverTab[28924]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:141
		_go_fuzz_dep_.CoverTab[28925]++
//line /usr/local/go/src/text/template/parse/lex.go:141
		// _ = "end of CoverTab[28925]"
//line /usr/local/go/src/text/template/parse/lex.go:141
	}
//line /usr/local/go/src/text/template/parse/lex.go:141
	// _ = "end of CoverTab[28921]"
//line /usr/local/go/src/text/template/parse/lex.go:141
	_go_fuzz_dep_.CoverTab[28922]++
								r, w := utf8.DecodeRuneInString(l.input[l.pos:])
								l.pos += Pos(w)
								if r == '\n' {
//line /usr/local/go/src/text/template/parse/lex.go:144
		_go_fuzz_dep_.CoverTab[28926]++
									l.line++
//line /usr/local/go/src/text/template/parse/lex.go:145
		// _ = "end of CoverTab[28926]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:146
		_go_fuzz_dep_.CoverTab[28927]++
//line /usr/local/go/src/text/template/parse/lex.go:146
		// _ = "end of CoverTab[28927]"
//line /usr/local/go/src/text/template/parse/lex.go:146
	}
//line /usr/local/go/src/text/template/parse/lex.go:146
	// _ = "end of CoverTab[28922]"
//line /usr/local/go/src/text/template/parse/lex.go:146
	_go_fuzz_dep_.CoverTab[28923]++
								return r
//line /usr/local/go/src/text/template/parse/lex.go:147
	// _ = "end of CoverTab[28923]"
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
//line /usr/local/go/src/text/template/parse/lex.go:151
	_go_fuzz_dep_.CoverTab[28928]++
								r := l.next()
								l.backup()
								return r
//line /usr/local/go/src/text/template/parse/lex.go:154
	// _ = "end of CoverTab[28928]"
}

// backup steps back one rune.
func (l *lexer) backup() {
//line /usr/local/go/src/text/template/parse/lex.go:158
	_go_fuzz_dep_.CoverTab[28929]++
								if !l.atEOF && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:159
		_go_fuzz_dep_.CoverTab[28930]++
//line /usr/local/go/src/text/template/parse/lex.go:159
		return l.pos > 0
//line /usr/local/go/src/text/template/parse/lex.go:159
		// _ = "end of CoverTab[28930]"
//line /usr/local/go/src/text/template/parse/lex.go:159
	}() {
//line /usr/local/go/src/text/template/parse/lex.go:159
		_go_fuzz_dep_.CoverTab[28931]++
									r, w := utf8.DecodeLastRuneInString(l.input[:l.pos])
									l.pos -= Pos(w)

									if r == '\n' {
//line /usr/local/go/src/text/template/parse/lex.go:163
			_go_fuzz_dep_.CoverTab[28932]++
										l.line--
//line /usr/local/go/src/text/template/parse/lex.go:164
			// _ = "end of CoverTab[28932]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:165
			_go_fuzz_dep_.CoverTab[28933]++
//line /usr/local/go/src/text/template/parse/lex.go:165
			// _ = "end of CoverTab[28933]"
//line /usr/local/go/src/text/template/parse/lex.go:165
		}
//line /usr/local/go/src/text/template/parse/lex.go:165
		// _ = "end of CoverTab[28931]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:166
		_go_fuzz_dep_.CoverTab[28934]++
//line /usr/local/go/src/text/template/parse/lex.go:166
		// _ = "end of CoverTab[28934]"
//line /usr/local/go/src/text/template/parse/lex.go:166
	}
//line /usr/local/go/src/text/template/parse/lex.go:166
	// _ = "end of CoverTab[28929]"
}

// thisItem returns the item at the current input point with the specified type
//line /usr/local/go/src/text/template/parse/lex.go:169
// and advances the input.
//line /usr/local/go/src/text/template/parse/lex.go:171
func (l *lexer) thisItem(t itemType) item {
//line /usr/local/go/src/text/template/parse/lex.go:171
	_go_fuzz_dep_.CoverTab[28935]++
								i := item{t, l.start, l.input[l.start:l.pos], l.startLine}
								l.start = l.pos
								l.startLine = l.line
								return i
//line /usr/local/go/src/text/template/parse/lex.go:175
	// _ = "end of CoverTab[28935]"
}

// emit passes the trailing text as an item back to the parser.
func (l *lexer) emit(t itemType) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:179
	_go_fuzz_dep_.CoverTab[28936]++
								return l.emitItem(l.thisItem(t))
//line /usr/local/go/src/text/template/parse/lex.go:180
	// _ = "end of CoverTab[28936]"
}

// emitItem passes the specified item to the parser.
func (l *lexer) emitItem(i item) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:184
	_go_fuzz_dep_.CoverTab[28937]++
								l.item = i
								return nil
//line /usr/local/go/src/text/template/parse/lex.go:186
	// _ = "end of CoverTab[28937]"
}

// ignore skips over the pending input before this point.
//line /usr/local/go/src/text/template/parse/lex.go:189
// It tracks newlines in the ignored text, so use it only
//line /usr/local/go/src/text/template/parse/lex.go:189
// for text that is skipped without calling l.next.
//line /usr/local/go/src/text/template/parse/lex.go:192
func (l *lexer) ignore() {
//line /usr/local/go/src/text/template/parse/lex.go:192
	_go_fuzz_dep_.CoverTab[28938]++
								l.line += strings.Count(l.input[l.start:l.pos], "\n")
								l.start = l.pos
								l.startLine = l.line
//line /usr/local/go/src/text/template/parse/lex.go:195
	// _ = "end of CoverTab[28938]"
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
//line /usr/local/go/src/text/template/parse/lex.go:199
	_go_fuzz_dep_.CoverTab[28939]++
								if strings.ContainsRune(valid, l.next()) {
//line /usr/local/go/src/text/template/parse/lex.go:200
		_go_fuzz_dep_.CoverTab[28941]++
									return true
//line /usr/local/go/src/text/template/parse/lex.go:201
		// _ = "end of CoverTab[28941]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:202
		_go_fuzz_dep_.CoverTab[28942]++
//line /usr/local/go/src/text/template/parse/lex.go:202
		// _ = "end of CoverTab[28942]"
//line /usr/local/go/src/text/template/parse/lex.go:202
	}
//line /usr/local/go/src/text/template/parse/lex.go:202
	// _ = "end of CoverTab[28939]"
//line /usr/local/go/src/text/template/parse/lex.go:202
	_go_fuzz_dep_.CoverTab[28940]++
								l.backup()
								return false
//line /usr/local/go/src/text/template/parse/lex.go:204
	// _ = "end of CoverTab[28940]"
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
//line /usr/local/go/src/text/template/parse/lex.go:208
	_go_fuzz_dep_.CoverTab[28943]++
								for strings.ContainsRune(valid, l.next()) {
//line /usr/local/go/src/text/template/parse/lex.go:209
		_go_fuzz_dep_.CoverTab[28945]++
//line /usr/local/go/src/text/template/parse/lex.go:209
		// _ = "end of CoverTab[28945]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:210
	// _ = "end of CoverTab[28943]"
//line /usr/local/go/src/text/template/parse/lex.go:210
	_go_fuzz_dep_.CoverTab[28944]++
								l.backup()
//line /usr/local/go/src/text/template/parse/lex.go:211
	// _ = "end of CoverTab[28944]"
}

// errorf returns an error token and terminates the scan by passing
//line /usr/local/go/src/text/template/parse/lex.go:214
// back a nil pointer that will be the next state, terminating l.nextItem.
//line /usr/local/go/src/text/template/parse/lex.go:216
func (l *lexer) errorf(format string, args ...any) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:216
	_go_fuzz_dep_.CoverTab[28946]++
								l.item = item{itemError, l.start, fmt.Sprintf(format, args...), l.startLine}
								l.start = 0
								l.pos = 0
								l.input = l.input[:0]
								return nil
//line /usr/local/go/src/text/template/parse/lex.go:221
	// _ = "end of CoverTab[28946]"
}

// nextItem returns the next item from the input.
//line /usr/local/go/src/text/template/parse/lex.go:224
// Called by the parser, not in the lexing goroutine.
//line /usr/local/go/src/text/template/parse/lex.go:226
func (l *lexer) nextItem() item {
//line /usr/local/go/src/text/template/parse/lex.go:226
	_go_fuzz_dep_.CoverTab[28947]++
								l.item = item{itemEOF, l.pos, "EOF", l.startLine}
								state := lexText
								if l.insideAction {
//line /usr/local/go/src/text/template/parse/lex.go:229
		_go_fuzz_dep_.CoverTab[28949]++
									state = lexInsideAction
//line /usr/local/go/src/text/template/parse/lex.go:230
		// _ = "end of CoverTab[28949]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:231
		_go_fuzz_dep_.CoverTab[28950]++
//line /usr/local/go/src/text/template/parse/lex.go:231
		// _ = "end of CoverTab[28950]"
//line /usr/local/go/src/text/template/parse/lex.go:231
	}
//line /usr/local/go/src/text/template/parse/lex.go:231
	// _ = "end of CoverTab[28947]"
//line /usr/local/go/src/text/template/parse/lex.go:231
	_go_fuzz_dep_.CoverTab[28948]++
								for {
//line /usr/local/go/src/text/template/parse/lex.go:232
		_go_fuzz_dep_.CoverTab[28951]++
									state = state(l)
									if state == nil {
//line /usr/local/go/src/text/template/parse/lex.go:234
			_go_fuzz_dep_.CoverTab[28952]++
										return l.item
//line /usr/local/go/src/text/template/parse/lex.go:235
			// _ = "end of CoverTab[28952]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:236
			_go_fuzz_dep_.CoverTab[28953]++
//line /usr/local/go/src/text/template/parse/lex.go:236
			// _ = "end of CoverTab[28953]"
//line /usr/local/go/src/text/template/parse/lex.go:236
		}
//line /usr/local/go/src/text/template/parse/lex.go:236
		// _ = "end of CoverTab[28951]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:237
	// _ = "end of CoverTab[28948]"
}

// lex creates a new scanner for the input string.
func lex(name, input, left, right string) *lexer {
//line /usr/local/go/src/text/template/parse/lex.go:241
	_go_fuzz_dep_.CoverTab[28954]++
								if left == "" {
//line /usr/local/go/src/text/template/parse/lex.go:242
		_go_fuzz_dep_.CoverTab[28957]++
									left = leftDelim
//line /usr/local/go/src/text/template/parse/lex.go:243
		// _ = "end of CoverTab[28957]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:244
		_go_fuzz_dep_.CoverTab[28958]++
//line /usr/local/go/src/text/template/parse/lex.go:244
		// _ = "end of CoverTab[28958]"
//line /usr/local/go/src/text/template/parse/lex.go:244
	}
//line /usr/local/go/src/text/template/parse/lex.go:244
	// _ = "end of CoverTab[28954]"
//line /usr/local/go/src/text/template/parse/lex.go:244
	_go_fuzz_dep_.CoverTab[28955]++
								if right == "" {
//line /usr/local/go/src/text/template/parse/lex.go:245
		_go_fuzz_dep_.CoverTab[28959]++
									right = rightDelim
//line /usr/local/go/src/text/template/parse/lex.go:246
		// _ = "end of CoverTab[28959]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:247
		_go_fuzz_dep_.CoverTab[28960]++
//line /usr/local/go/src/text/template/parse/lex.go:247
		// _ = "end of CoverTab[28960]"
//line /usr/local/go/src/text/template/parse/lex.go:247
	}
//line /usr/local/go/src/text/template/parse/lex.go:247
	// _ = "end of CoverTab[28955]"
//line /usr/local/go/src/text/template/parse/lex.go:247
	_go_fuzz_dep_.CoverTab[28956]++
								l := &lexer{
		name:		name,
		input:		input,
		leftDelim:	left,
		rightDelim:	right,
		line:		1,
		startLine:	1,
		insideAction:	false,
	}
								return l
//line /usr/local/go/src/text/template/parse/lex.go:257
	// _ = "end of CoverTab[28956]"
}

//line /usr/local/go/src/text/template/parse/lex.go:262
const (
	leftDelim	= "{{"
	rightDelim	= "}}"
	leftComment	= "/*"
	rightComment	= "*/"
)

// lexText scans until an opening action delimiter, "{{".
func lexText(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:270
	_go_fuzz_dep_.CoverTab[28961]++
								if x := strings.Index(l.input[l.pos:], l.leftDelim); x >= 0 {
//line /usr/local/go/src/text/template/parse/lex.go:271
		_go_fuzz_dep_.CoverTab[28964]++
									if x > 0 {
//line /usr/local/go/src/text/template/parse/lex.go:272
			_go_fuzz_dep_.CoverTab[28966]++
										l.pos += Pos(x)

										trimLength := Pos(0)
										delimEnd := l.pos + Pos(len(l.leftDelim))
										if hasLeftTrimMarker(l.input[delimEnd:]) {
//line /usr/local/go/src/text/template/parse/lex.go:277
				_go_fuzz_dep_.CoverTab[28968]++
											trimLength = rightTrimLength(l.input[l.start:l.pos])
//line /usr/local/go/src/text/template/parse/lex.go:278
				// _ = "end of CoverTab[28968]"
			} else {
//line /usr/local/go/src/text/template/parse/lex.go:279
				_go_fuzz_dep_.CoverTab[28969]++
//line /usr/local/go/src/text/template/parse/lex.go:279
				// _ = "end of CoverTab[28969]"
//line /usr/local/go/src/text/template/parse/lex.go:279
			}
//line /usr/local/go/src/text/template/parse/lex.go:279
			// _ = "end of CoverTab[28966]"
//line /usr/local/go/src/text/template/parse/lex.go:279
			_go_fuzz_dep_.CoverTab[28967]++
										l.pos -= trimLength
										l.line += strings.Count(l.input[l.start:l.pos], "\n")
										i := l.thisItem(itemText)
										l.pos += trimLength
										l.ignore()
										if len(i.val) > 0 {
//line /usr/local/go/src/text/template/parse/lex.go:285
				_go_fuzz_dep_.CoverTab[28970]++
											return l.emitItem(i)
//line /usr/local/go/src/text/template/parse/lex.go:286
				// _ = "end of CoverTab[28970]"
			} else {
//line /usr/local/go/src/text/template/parse/lex.go:287
				_go_fuzz_dep_.CoverTab[28971]++
//line /usr/local/go/src/text/template/parse/lex.go:287
				// _ = "end of CoverTab[28971]"
//line /usr/local/go/src/text/template/parse/lex.go:287
			}
//line /usr/local/go/src/text/template/parse/lex.go:287
			// _ = "end of CoverTab[28967]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:288
			_go_fuzz_dep_.CoverTab[28972]++
//line /usr/local/go/src/text/template/parse/lex.go:288
			// _ = "end of CoverTab[28972]"
//line /usr/local/go/src/text/template/parse/lex.go:288
		}
//line /usr/local/go/src/text/template/parse/lex.go:288
		// _ = "end of CoverTab[28964]"
//line /usr/local/go/src/text/template/parse/lex.go:288
		_go_fuzz_dep_.CoverTab[28965]++
									return lexLeftDelim
//line /usr/local/go/src/text/template/parse/lex.go:289
		// _ = "end of CoverTab[28965]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:290
		_go_fuzz_dep_.CoverTab[28973]++
//line /usr/local/go/src/text/template/parse/lex.go:290
		// _ = "end of CoverTab[28973]"
//line /usr/local/go/src/text/template/parse/lex.go:290
	}
//line /usr/local/go/src/text/template/parse/lex.go:290
	// _ = "end of CoverTab[28961]"
//line /usr/local/go/src/text/template/parse/lex.go:290
	_go_fuzz_dep_.CoverTab[28962]++
								l.pos = Pos(len(l.input))

								if l.pos > l.start {
//line /usr/local/go/src/text/template/parse/lex.go:293
		_go_fuzz_dep_.CoverTab[28974]++
									l.line += strings.Count(l.input[l.start:l.pos], "\n")
									return l.emit(itemText)
//line /usr/local/go/src/text/template/parse/lex.go:295
		// _ = "end of CoverTab[28974]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:296
		_go_fuzz_dep_.CoverTab[28975]++
//line /usr/local/go/src/text/template/parse/lex.go:296
		// _ = "end of CoverTab[28975]"
//line /usr/local/go/src/text/template/parse/lex.go:296
	}
//line /usr/local/go/src/text/template/parse/lex.go:296
	// _ = "end of CoverTab[28962]"
//line /usr/local/go/src/text/template/parse/lex.go:296
	_go_fuzz_dep_.CoverTab[28963]++
								return l.emit(itemEOF)
//line /usr/local/go/src/text/template/parse/lex.go:297
	// _ = "end of CoverTab[28963]"
}

// rightTrimLength returns the length of the spaces at the end of the string.
func rightTrimLength(s string) Pos {
//line /usr/local/go/src/text/template/parse/lex.go:301
	_go_fuzz_dep_.CoverTab[28976]++
								return Pos(len(s) - len(strings.TrimRight(s, spaceChars)))
//line /usr/local/go/src/text/template/parse/lex.go:302
	// _ = "end of CoverTab[28976]"
}

// atRightDelim reports whether the lexer is at a right delimiter, possibly preceded by a trim marker.
func (l *lexer) atRightDelim() (delim, trimSpaces bool) {
//line /usr/local/go/src/text/template/parse/lex.go:306
	_go_fuzz_dep_.CoverTab[28977]++
								if hasRightTrimMarker(l.input[l.pos:]) && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:307
		_go_fuzz_dep_.CoverTab[28980]++
//line /usr/local/go/src/text/template/parse/lex.go:307
		return strings.HasPrefix(l.input[l.pos+trimMarkerLen:], l.rightDelim)
//line /usr/local/go/src/text/template/parse/lex.go:307
		// _ = "end of CoverTab[28980]"
//line /usr/local/go/src/text/template/parse/lex.go:307
	}() {
//line /usr/local/go/src/text/template/parse/lex.go:307
		_go_fuzz_dep_.CoverTab[28981]++
									return true, true
//line /usr/local/go/src/text/template/parse/lex.go:308
		// _ = "end of CoverTab[28981]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:309
		_go_fuzz_dep_.CoverTab[28982]++
//line /usr/local/go/src/text/template/parse/lex.go:309
		// _ = "end of CoverTab[28982]"
//line /usr/local/go/src/text/template/parse/lex.go:309
	}
//line /usr/local/go/src/text/template/parse/lex.go:309
	// _ = "end of CoverTab[28977]"
//line /usr/local/go/src/text/template/parse/lex.go:309
	_go_fuzz_dep_.CoverTab[28978]++
								if strings.HasPrefix(l.input[l.pos:], l.rightDelim) {
//line /usr/local/go/src/text/template/parse/lex.go:310
		_go_fuzz_dep_.CoverTab[28983]++
									return true, false
//line /usr/local/go/src/text/template/parse/lex.go:311
		// _ = "end of CoverTab[28983]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:312
		_go_fuzz_dep_.CoverTab[28984]++
//line /usr/local/go/src/text/template/parse/lex.go:312
		// _ = "end of CoverTab[28984]"
//line /usr/local/go/src/text/template/parse/lex.go:312
	}
//line /usr/local/go/src/text/template/parse/lex.go:312
	// _ = "end of CoverTab[28978]"
//line /usr/local/go/src/text/template/parse/lex.go:312
	_go_fuzz_dep_.CoverTab[28979]++
								return false, false
//line /usr/local/go/src/text/template/parse/lex.go:313
	// _ = "end of CoverTab[28979]"
}

// leftTrimLength returns the length of the spaces at the beginning of the string.
func leftTrimLength(s string) Pos {
//line /usr/local/go/src/text/template/parse/lex.go:317
	_go_fuzz_dep_.CoverTab[28985]++
								return Pos(len(s) - len(strings.TrimLeft(s, spaceChars)))
//line /usr/local/go/src/text/template/parse/lex.go:318
	// _ = "end of CoverTab[28985]"
}

// lexLeftDelim scans the left delimiter, which is known to be present, possibly with a trim marker.
//line /usr/local/go/src/text/template/parse/lex.go:321
// (The text to be trimmed has already been emitted.)
//line /usr/local/go/src/text/template/parse/lex.go:323
func lexLeftDelim(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:323
	_go_fuzz_dep_.CoverTab[28986]++
								l.pos += Pos(len(l.leftDelim))
								trimSpace := hasLeftTrimMarker(l.input[l.pos:])
								afterMarker := Pos(0)
								if trimSpace {
//line /usr/local/go/src/text/template/parse/lex.go:327
		_go_fuzz_dep_.CoverTab[28989]++
									afterMarker = trimMarkerLen
//line /usr/local/go/src/text/template/parse/lex.go:328
		// _ = "end of CoverTab[28989]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:329
		_go_fuzz_dep_.CoverTab[28990]++
//line /usr/local/go/src/text/template/parse/lex.go:329
		// _ = "end of CoverTab[28990]"
//line /usr/local/go/src/text/template/parse/lex.go:329
	}
//line /usr/local/go/src/text/template/parse/lex.go:329
	// _ = "end of CoverTab[28986]"
//line /usr/local/go/src/text/template/parse/lex.go:329
	_go_fuzz_dep_.CoverTab[28987]++
								if strings.HasPrefix(l.input[l.pos+afterMarker:], leftComment) {
//line /usr/local/go/src/text/template/parse/lex.go:330
		_go_fuzz_dep_.CoverTab[28991]++
									l.pos += afterMarker
									l.ignore()
									return lexComment
//line /usr/local/go/src/text/template/parse/lex.go:333
		// _ = "end of CoverTab[28991]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:334
		_go_fuzz_dep_.CoverTab[28992]++
//line /usr/local/go/src/text/template/parse/lex.go:334
		// _ = "end of CoverTab[28992]"
//line /usr/local/go/src/text/template/parse/lex.go:334
	}
//line /usr/local/go/src/text/template/parse/lex.go:334
	// _ = "end of CoverTab[28987]"
//line /usr/local/go/src/text/template/parse/lex.go:334
	_go_fuzz_dep_.CoverTab[28988]++
								i := l.thisItem(itemLeftDelim)
								l.insideAction = true
								l.pos += afterMarker
								l.ignore()
								l.parenDepth = 0
								return l.emitItem(i)
//line /usr/local/go/src/text/template/parse/lex.go:340
	// _ = "end of CoverTab[28988]"
}

// lexComment scans a comment. The left comment marker is known to be present.
func lexComment(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:344
	_go_fuzz_dep_.CoverTab[28993]++
								l.pos += Pos(len(leftComment))
								x := strings.Index(l.input[l.pos:], rightComment)
								if x < 0 {
//line /usr/local/go/src/text/template/parse/lex.go:347
		_go_fuzz_dep_.CoverTab[28999]++
									return l.errorf("unclosed comment")
//line /usr/local/go/src/text/template/parse/lex.go:348
		// _ = "end of CoverTab[28999]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:349
		_go_fuzz_dep_.CoverTab[29000]++
//line /usr/local/go/src/text/template/parse/lex.go:349
		// _ = "end of CoverTab[29000]"
//line /usr/local/go/src/text/template/parse/lex.go:349
	}
//line /usr/local/go/src/text/template/parse/lex.go:349
	// _ = "end of CoverTab[28993]"
//line /usr/local/go/src/text/template/parse/lex.go:349
	_go_fuzz_dep_.CoverTab[28994]++
								l.pos += Pos(x + len(rightComment))
								delim, trimSpace := l.atRightDelim()
								if !delim {
//line /usr/local/go/src/text/template/parse/lex.go:352
		_go_fuzz_dep_.CoverTab[29001]++
									return l.errorf("comment ends before closing delimiter")
//line /usr/local/go/src/text/template/parse/lex.go:353
		// _ = "end of CoverTab[29001]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:354
		_go_fuzz_dep_.CoverTab[29002]++
//line /usr/local/go/src/text/template/parse/lex.go:354
		// _ = "end of CoverTab[29002]"
//line /usr/local/go/src/text/template/parse/lex.go:354
	}
//line /usr/local/go/src/text/template/parse/lex.go:354
	// _ = "end of CoverTab[28994]"
//line /usr/local/go/src/text/template/parse/lex.go:354
	_go_fuzz_dep_.CoverTab[28995]++
								i := l.thisItem(itemComment)
								if trimSpace {
//line /usr/local/go/src/text/template/parse/lex.go:356
		_go_fuzz_dep_.CoverTab[29003]++
									l.pos += trimMarkerLen
//line /usr/local/go/src/text/template/parse/lex.go:357
		// _ = "end of CoverTab[29003]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:358
		_go_fuzz_dep_.CoverTab[29004]++
//line /usr/local/go/src/text/template/parse/lex.go:358
		// _ = "end of CoverTab[29004]"
//line /usr/local/go/src/text/template/parse/lex.go:358
	}
//line /usr/local/go/src/text/template/parse/lex.go:358
	// _ = "end of CoverTab[28995]"
//line /usr/local/go/src/text/template/parse/lex.go:358
	_go_fuzz_dep_.CoverTab[28996]++
								l.pos += Pos(len(l.rightDelim))
								if trimSpace {
//line /usr/local/go/src/text/template/parse/lex.go:360
		_go_fuzz_dep_.CoverTab[29005]++
									l.pos += leftTrimLength(l.input[l.pos:])
//line /usr/local/go/src/text/template/parse/lex.go:361
		// _ = "end of CoverTab[29005]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:362
		_go_fuzz_dep_.CoverTab[29006]++
//line /usr/local/go/src/text/template/parse/lex.go:362
		// _ = "end of CoverTab[29006]"
//line /usr/local/go/src/text/template/parse/lex.go:362
	}
//line /usr/local/go/src/text/template/parse/lex.go:362
	// _ = "end of CoverTab[28996]"
//line /usr/local/go/src/text/template/parse/lex.go:362
	_go_fuzz_dep_.CoverTab[28997]++
								l.ignore()
								if l.options.emitComment {
//line /usr/local/go/src/text/template/parse/lex.go:364
		_go_fuzz_dep_.CoverTab[29007]++
									return l.emitItem(i)
//line /usr/local/go/src/text/template/parse/lex.go:365
		// _ = "end of CoverTab[29007]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:366
		_go_fuzz_dep_.CoverTab[29008]++
//line /usr/local/go/src/text/template/parse/lex.go:366
		// _ = "end of CoverTab[29008]"
//line /usr/local/go/src/text/template/parse/lex.go:366
	}
//line /usr/local/go/src/text/template/parse/lex.go:366
	// _ = "end of CoverTab[28997]"
//line /usr/local/go/src/text/template/parse/lex.go:366
	_go_fuzz_dep_.CoverTab[28998]++
								return lexText
//line /usr/local/go/src/text/template/parse/lex.go:367
	// _ = "end of CoverTab[28998]"
}

// lexRightDelim scans the right delimiter, which is known to be present, possibly with a trim marker.
func lexRightDelim(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:371
	_go_fuzz_dep_.CoverTab[29009]++
								_, trimSpace := l.atRightDelim()
								if trimSpace {
//line /usr/local/go/src/text/template/parse/lex.go:373
		_go_fuzz_dep_.CoverTab[29012]++
									l.pos += trimMarkerLen
									l.ignore()
//line /usr/local/go/src/text/template/parse/lex.go:375
		// _ = "end of CoverTab[29012]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:376
		_go_fuzz_dep_.CoverTab[29013]++
//line /usr/local/go/src/text/template/parse/lex.go:376
		// _ = "end of CoverTab[29013]"
//line /usr/local/go/src/text/template/parse/lex.go:376
	}
//line /usr/local/go/src/text/template/parse/lex.go:376
	// _ = "end of CoverTab[29009]"
//line /usr/local/go/src/text/template/parse/lex.go:376
	_go_fuzz_dep_.CoverTab[29010]++
								l.pos += Pos(len(l.rightDelim))
								i := l.thisItem(itemRightDelim)
								if trimSpace {
//line /usr/local/go/src/text/template/parse/lex.go:379
		_go_fuzz_dep_.CoverTab[29014]++
									l.pos += leftTrimLength(l.input[l.pos:])
									l.ignore()
//line /usr/local/go/src/text/template/parse/lex.go:381
		// _ = "end of CoverTab[29014]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:382
		_go_fuzz_dep_.CoverTab[29015]++
//line /usr/local/go/src/text/template/parse/lex.go:382
		// _ = "end of CoverTab[29015]"
//line /usr/local/go/src/text/template/parse/lex.go:382
	}
//line /usr/local/go/src/text/template/parse/lex.go:382
	// _ = "end of CoverTab[29010]"
//line /usr/local/go/src/text/template/parse/lex.go:382
	_go_fuzz_dep_.CoverTab[29011]++
								l.insideAction = false
								return l.emitItem(i)
//line /usr/local/go/src/text/template/parse/lex.go:384
	// _ = "end of CoverTab[29011]"
}

// lexInsideAction scans the elements inside action delimiters.
func lexInsideAction(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:388
	_go_fuzz_dep_.CoverTab[29016]++

//line /usr/local/go/src/text/template/parse/lex.go:392
	delim, _ := l.atRightDelim()
	if delim {
//line /usr/local/go/src/text/template/parse/lex.go:393
		_go_fuzz_dep_.CoverTab[29018]++
									if l.parenDepth == 0 {
//line /usr/local/go/src/text/template/parse/lex.go:394
			_go_fuzz_dep_.CoverTab[29020]++
										return lexRightDelim
//line /usr/local/go/src/text/template/parse/lex.go:395
			// _ = "end of CoverTab[29020]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:396
			_go_fuzz_dep_.CoverTab[29021]++
//line /usr/local/go/src/text/template/parse/lex.go:396
			// _ = "end of CoverTab[29021]"
//line /usr/local/go/src/text/template/parse/lex.go:396
		}
//line /usr/local/go/src/text/template/parse/lex.go:396
		// _ = "end of CoverTab[29018]"
//line /usr/local/go/src/text/template/parse/lex.go:396
		_go_fuzz_dep_.CoverTab[29019]++
									return l.errorf("unclosed left paren")
//line /usr/local/go/src/text/template/parse/lex.go:397
		// _ = "end of CoverTab[29019]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:398
		_go_fuzz_dep_.CoverTab[29022]++
//line /usr/local/go/src/text/template/parse/lex.go:398
		// _ = "end of CoverTab[29022]"
//line /usr/local/go/src/text/template/parse/lex.go:398
	}
//line /usr/local/go/src/text/template/parse/lex.go:398
	// _ = "end of CoverTab[29016]"
//line /usr/local/go/src/text/template/parse/lex.go:398
	_go_fuzz_dep_.CoverTab[29017]++
								switch r := l.next(); {
	case r == eof:
//line /usr/local/go/src/text/template/parse/lex.go:400
		_go_fuzz_dep_.CoverTab[29023]++
									return l.errorf("unclosed action")
//line /usr/local/go/src/text/template/parse/lex.go:401
		// _ = "end of CoverTab[29023]"
	case isSpace(r):
//line /usr/local/go/src/text/template/parse/lex.go:402
		_go_fuzz_dep_.CoverTab[29024]++
									l.backup()
									return lexSpace
//line /usr/local/go/src/text/template/parse/lex.go:404
		// _ = "end of CoverTab[29024]"
	case r == '=':
//line /usr/local/go/src/text/template/parse/lex.go:405
		_go_fuzz_dep_.CoverTab[29025]++
									return l.emit(itemAssign)
//line /usr/local/go/src/text/template/parse/lex.go:406
		// _ = "end of CoverTab[29025]"
	case r == ':':
//line /usr/local/go/src/text/template/parse/lex.go:407
		_go_fuzz_dep_.CoverTab[29026]++
									if l.next() != '=' {
//line /usr/local/go/src/text/template/parse/lex.go:408
			_go_fuzz_dep_.CoverTab[29042]++
										return l.errorf("expected :=")
//line /usr/local/go/src/text/template/parse/lex.go:409
			// _ = "end of CoverTab[29042]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:410
			_go_fuzz_dep_.CoverTab[29043]++
//line /usr/local/go/src/text/template/parse/lex.go:410
			// _ = "end of CoverTab[29043]"
//line /usr/local/go/src/text/template/parse/lex.go:410
		}
//line /usr/local/go/src/text/template/parse/lex.go:410
		// _ = "end of CoverTab[29026]"
//line /usr/local/go/src/text/template/parse/lex.go:410
		_go_fuzz_dep_.CoverTab[29027]++
									return l.emit(itemDeclare)
//line /usr/local/go/src/text/template/parse/lex.go:411
		// _ = "end of CoverTab[29027]"
	case r == '|':
//line /usr/local/go/src/text/template/parse/lex.go:412
		_go_fuzz_dep_.CoverTab[29028]++
									return l.emit(itemPipe)
//line /usr/local/go/src/text/template/parse/lex.go:413
		// _ = "end of CoverTab[29028]"
	case r == '"':
//line /usr/local/go/src/text/template/parse/lex.go:414
		_go_fuzz_dep_.CoverTab[29029]++
									return lexQuote
//line /usr/local/go/src/text/template/parse/lex.go:415
		// _ = "end of CoverTab[29029]"
	case r == '`':
//line /usr/local/go/src/text/template/parse/lex.go:416
		_go_fuzz_dep_.CoverTab[29030]++
									return lexRawQuote
//line /usr/local/go/src/text/template/parse/lex.go:417
		// _ = "end of CoverTab[29030]"
	case r == '$':
//line /usr/local/go/src/text/template/parse/lex.go:418
		_go_fuzz_dep_.CoverTab[29031]++
									return lexVariable
//line /usr/local/go/src/text/template/parse/lex.go:419
		// _ = "end of CoverTab[29031]"
	case r == '\'':
//line /usr/local/go/src/text/template/parse/lex.go:420
		_go_fuzz_dep_.CoverTab[29032]++
									return lexChar
//line /usr/local/go/src/text/template/parse/lex.go:421
		// _ = "end of CoverTab[29032]"
	case r == '.':
//line /usr/local/go/src/text/template/parse/lex.go:422
		_go_fuzz_dep_.CoverTab[29033]++

									if l.pos < Pos(len(l.input)) {
//line /usr/local/go/src/text/template/parse/lex.go:424
			_go_fuzz_dep_.CoverTab[29044]++
										r := l.input[l.pos]
										if r < '0' || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:426
				_go_fuzz_dep_.CoverTab[29045]++
//line /usr/local/go/src/text/template/parse/lex.go:426
				return '9' < r
//line /usr/local/go/src/text/template/parse/lex.go:426
				// _ = "end of CoverTab[29045]"
//line /usr/local/go/src/text/template/parse/lex.go:426
			}() {
//line /usr/local/go/src/text/template/parse/lex.go:426
				_go_fuzz_dep_.CoverTab[29046]++
											return lexField
//line /usr/local/go/src/text/template/parse/lex.go:427
				// _ = "end of CoverTab[29046]"
			} else {
//line /usr/local/go/src/text/template/parse/lex.go:428
				_go_fuzz_dep_.CoverTab[29047]++
//line /usr/local/go/src/text/template/parse/lex.go:428
				// _ = "end of CoverTab[29047]"
//line /usr/local/go/src/text/template/parse/lex.go:428
			}
//line /usr/local/go/src/text/template/parse/lex.go:428
			// _ = "end of CoverTab[29044]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:429
			_go_fuzz_dep_.CoverTab[29048]++
//line /usr/local/go/src/text/template/parse/lex.go:429
			// _ = "end of CoverTab[29048]"
//line /usr/local/go/src/text/template/parse/lex.go:429
		}
//line /usr/local/go/src/text/template/parse/lex.go:429
		// _ = "end of CoverTab[29033]"
//line /usr/local/go/src/text/template/parse/lex.go:429
		_go_fuzz_dep_.CoverTab[29034]++
									fallthrough
//line /usr/local/go/src/text/template/parse/lex.go:430
		// _ = "end of CoverTab[29034]"
	case r == '+' || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:431
		_go_fuzz_dep_.CoverTab[29049]++
//line /usr/local/go/src/text/template/parse/lex.go:431
		return r == '-'
//line /usr/local/go/src/text/template/parse/lex.go:431
		// _ = "end of CoverTab[29049]"
//line /usr/local/go/src/text/template/parse/lex.go:431
	}() || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:431
		_go_fuzz_dep_.CoverTab[29050]++
//line /usr/local/go/src/text/template/parse/lex.go:431
		return ('0' <= r && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:431
			_go_fuzz_dep_.CoverTab[29051]++
//line /usr/local/go/src/text/template/parse/lex.go:431
			return r <= '9'
//line /usr/local/go/src/text/template/parse/lex.go:431
			// _ = "end of CoverTab[29051]"
//line /usr/local/go/src/text/template/parse/lex.go:431
		}())
//line /usr/local/go/src/text/template/parse/lex.go:431
		// _ = "end of CoverTab[29050]"
//line /usr/local/go/src/text/template/parse/lex.go:431
	}():
//line /usr/local/go/src/text/template/parse/lex.go:431
		_go_fuzz_dep_.CoverTab[29035]++
									l.backup()
									return lexNumber
//line /usr/local/go/src/text/template/parse/lex.go:433
		// _ = "end of CoverTab[29035]"
	case isAlphaNumeric(r):
//line /usr/local/go/src/text/template/parse/lex.go:434
		_go_fuzz_dep_.CoverTab[29036]++
									l.backup()
									return lexIdentifier
//line /usr/local/go/src/text/template/parse/lex.go:436
		// _ = "end of CoverTab[29036]"
	case r == '(':
//line /usr/local/go/src/text/template/parse/lex.go:437
		_go_fuzz_dep_.CoverTab[29037]++
									l.parenDepth++
									return l.emit(itemLeftParen)
//line /usr/local/go/src/text/template/parse/lex.go:439
		// _ = "end of CoverTab[29037]"
	case r == ')':
//line /usr/local/go/src/text/template/parse/lex.go:440
		_go_fuzz_dep_.CoverTab[29038]++
									l.parenDepth--
									if l.parenDepth < 0 {
//line /usr/local/go/src/text/template/parse/lex.go:442
			_go_fuzz_dep_.CoverTab[29052]++
										return l.errorf("unexpected right paren")
//line /usr/local/go/src/text/template/parse/lex.go:443
			// _ = "end of CoverTab[29052]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:444
			_go_fuzz_dep_.CoverTab[29053]++
//line /usr/local/go/src/text/template/parse/lex.go:444
			// _ = "end of CoverTab[29053]"
//line /usr/local/go/src/text/template/parse/lex.go:444
		}
//line /usr/local/go/src/text/template/parse/lex.go:444
		// _ = "end of CoverTab[29038]"
//line /usr/local/go/src/text/template/parse/lex.go:444
		_go_fuzz_dep_.CoverTab[29039]++
									return l.emit(itemRightParen)
//line /usr/local/go/src/text/template/parse/lex.go:445
		// _ = "end of CoverTab[29039]"
	case r <= unicode.MaxASCII && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:446
		_go_fuzz_dep_.CoverTab[29054]++
//line /usr/local/go/src/text/template/parse/lex.go:446
		return unicode.IsPrint(r)
//line /usr/local/go/src/text/template/parse/lex.go:446
		// _ = "end of CoverTab[29054]"
//line /usr/local/go/src/text/template/parse/lex.go:446
	}():
//line /usr/local/go/src/text/template/parse/lex.go:446
		_go_fuzz_dep_.CoverTab[29040]++
									return l.emit(itemChar)
//line /usr/local/go/src/text/template/parse/lex.go:447
		// _ = "end of CoverTab[29040]"
	default:
//line /usr/local/go/src/text/template/parse/lex.go:448
		_go_fuzz_dep_.CoverTab[29041]++
									return l.errorf("unrecognized character in action: %#U", r)
//line /usr/local/go/src/text/template/parse/lex.go:449
		// _ = "end of CoverTab[29041]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:450
	// _ = "end of CoverTab[29017]"
}

// lexSpace scans a run of space characters.
//line /usr/local/go/src/text/template/parse/lex.go:453
// We have not consumed the first space, which is known to be present.
//line /usr/local/go/src/text/template/parse/lex.go:453
// Take care if there is a trim-marked right delimiter, which starts with a space.
//line /usr/local/go/src/text/template/parse/lex.go:456
func lexSpace(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:456
	_go_fuzz_dep_.CoverTab[29055]++
								var r rune
								var numSpaces int
								for {
//line /usr/local/go/src/text/template/parse/lex.go:459
		_go_fuzz_dep_.CoverTab[29058]++
									r = l.peek()
									if !isSpace(r) {
//line /usr/local/go/src/text/template/parse/lex.go:461
			_go_fuzz_dep_.CoverTab[29060]++
										break
//line /usr/local/go/src/text/template/parse/lex.go:462
			// _ = "end of CoverTab[29060]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:463
			_go_fuzz_dep_.CoverTab[29061]++
//line /usr/local/go/src/text/template/parse/lex.go:463
			// _ = "end of CoverTab[29061]"
//line /usr/local/go/src/text/template/parse/lex.go:463
		}
//line /usr/local/go/src/text/template/parse/lex.go:463
		// _ = "end of CoverTab[29058]"
//line /usr/local/go/src/text/template/parse/lex.go:463
		_go_fuzz_dep_.CoverTab[29059]++
									l.next()
									numSpaces++
//line /usr/local/go/src/text/template/parse/lex.go:465
		// _ = "end of CoverTab[29059]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:466
	// _ = "end of CoverTab[29055]"
//line /usr/local/go/src/text/template/parse/lex.go:466
	_go_fuzz_dep_.CoverTab[29056]++

//line /usr/local/go/src/text/template/parse/lex.go:469
	if hasRightTrimMarker(l.input[l.pos-1:]) && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:469
		_go_fuzz_dep_.CoverTab[29062]++
//line /usr/local/go/src/text/template/parse/lex.go:469
		return strings.HasPrefix(l.input[l.pos-1+trimMarkerLen:], l.rightDelim)
//line /usr/local/go/src/text/template/parse/lex.go:469
		// _ = "end of CoverTab[29062]"
//line /usr/local/go/src/text/template/parse/lex.go:469
	}() {
//line /usr/local/go/src/text/template/parse/lex.go:469
		_go_fuzz_dep_.CoverTab[29063]++
									l.backup()
									if numSpaces == 1 {
//line /usr/local/go/src/text/template/parse/lex.go:471
			_go_fuzz_dep_.CoverTab[29064]++
										return lexRightDelim
//line /usr/local/go/src/text/template/parse/lex.go:472
			// _ = "end of CoverTab[29064]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:473
			_go_fuzz_dep_.CoverTab[29065]++
//line /usr/local/go/src/text/template/parse/lex.go:473
			// _ = "end of CoverTab[29065]"
//line /usr/local/go/src/text/template/parse/lex.go:473
		}
//line /usr/local/go/src/text/template/parse/lex.go:473
		// _ = "end of CoverTab[29063]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:474
		_go_fuzz_dep_.CoverTab[29066]++
//line /usr/local/go/src/text/template/parse/lex.go:474
		// _ = "end of CoverTab[29066]"
//line /usr/local/go/src/text/template/parse/lex.go:474
	}
//line /usr/local/go/src/text/template/parse/lex.go:474
	// _ = "end of CoverTab[29056]"
//line /usr/local/go/src/text/template/parse/lex.go:474
	_go_fuzz_dep_.CoverTab[29057]++
								return l.emit(itemSpace)
//line /usr/local/go/src/text/template/parse/lex.go:475
	// _ = "end of CoverTab[29057]"
}

// lexIdentifier scans an alphanumeric.
func lexIdentifier(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:479
	_go_fuzz_dep_.CoverTab[29067]++
								for {
//line /usr/local/go/src/text/template/parse/lex.go:480
		_go_fuzz_dep_.CoverTab[29068]++
									switch r := l.next(); {
		case isAlphaNumeric(r):
//line /usr/local/go/src/text/template/parse/lex.go:482
			_go_fuzz_dep_.CoverTab[29069]++
//line /usr/local/go/src/text/template/parse/lex.go:482
			// _ = "end of CoverTab[29069]"

		default:
//line /usr/local/go/src/text/template/parse/lex.go:484
			_go_fuzz_dep_.CoverTab[29070]++
										l.backup()
										word := l.input[l.start:l.pos]
										if !l.atTerminator() {
//line /usr/local/go/src/text/template/parse/lex.go:487
				_go_fuzz_dep_.CoverTab[29072]++
											return l.errorf("bad character %#U", r)
//line /usr/local/go/src/text/template/parse/lex.go:488
				// _ = "end of CoverTab[29072]"
			} else {
//line /usr/local/go/src/text/template/parse/lex.go:489
				_go_fuzz_dep_.CoverTab[29073]++
//line /usr/local/go/src/text/template/parse/lex.go:489
				// _ = "end of CoverTab[29073]"
//line /usr/local/go/src/text/template/parse/lex.go:489
			}
//line /usr/local/go/src/text/template/parse/lex.go:489
			// _ = "end of CoverTab[29070]"
//line /usr/local/go/src/text/template/parse/lex.go:489
			_go_fuzz_dep_.CoverTab[29071]++
										switch {
			case key[word] > itemKeyword:
//line /usr/local/go/src/text/template/parse/lex.go:491
				_go_fuzz_dep_.CoverTab[29074]++
											item := key[word]
											if item == itemBreak && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:493
					_go_fuzz_dep_.CoverTab[29079]++
//line /usr/local/go/src/text/template/parse/lex.go:493
					return !l.options.breakOK
//line /usr/local/go/src/text/template/parse/lex.go:493
					// _ = "end of CoverTab[29079]"
//line /usr/local/go/src/text/template/parse/lex.go:493
				}() || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:493
					_go_fuzz_dep_.CoverTab[29080]++
//line /usr/local/go/src/text/template/parse/lex.go:493
					return item == itemContinue && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:493
						_go_fuzz_dep_.CoverTab[29081]++
//line /usr/local/go/src/text/template/parse/lex.go:493
						return !l.options.continueOK
//line /usr/local/go/src/text/template/parse/lex.go:493
						// _ = "end of CoverTab[29081]"
//line /usr/local/go/src/text/template/parse/lex.go:493
					}()
//line /usr/local/go/src/text/template/parse/lex.go:493
					// _ = "end of CoverTab[29080]"
//line /usr/local/go/src/text/template/parse/lex.go:493
				}() {
//line /usr/local/go/src/text/template/parse/lex.go:493
					_go_fuzz_dep_.CoverTab[29082]++
												return l.emit(itemIdentifier)
//line /usr/local/go/src/text/template/parse/lex.go:494
					// _ = "end of CoverTab[29082]"
				} else {
//line /usr/local/go/src/text/template/parse/lex.go:495
					_go_fuzz_dep_.CoverTab[29083]++
//line /usr/local/go/src/text/template/parse/lex.go:495
					// _ = "end of CoverTab[29083]"
//line /usr/local/go/src/text/template/parse/lex.go:495
				}
//line /usr/local/go/src/text/template/parse/lex.go:495
				// _ = "end of CoverTab[29074]"
//line /usr/local/go/src/text/template/parse/lex.go:495
				_go_fuzz_dep_.CoverTab[29075]++
											return l.emit(item)
//line /usr/local/go/src/text/template/parse/lex.go:496
				// _ = "end of CoverTab[29075]"
			case word[0] == '.':
//line /usr/local/go/src/text/template/parse/lex.go:497
				_go_fuzz_dep_.CoverTab[29076]++
											return l.emit(itemField)
//line /usr/local/go/src/text/template/parse/lex.go:498
				// _ = "end of CoverTab[29076]"
			case word == "true", word == "false":
//line /usr/local/go/src/text/template/parse/lex.go:499
				_go_fuzz_dep_.CoverTab[29077]++
											return l.emit(itemBool)
//line /usr/local/go/src/text/template/parse/lex.go:500
				// _ = "end of CoverTab[29077]"
			default:
//line /usr/local/go/src/text/template/parse/lex.go:501
				_go_fuzz_dep_.CoverTab[29078]++
											return l.emit(itemIdentifier)
//line /usr/local/go/src/text/template/parse/lex.go:502
				// _ = "end of CoverTab[29078]"
			}
//line /usr/local/go/src/text/template/parse/lex.go:503
			// _ = "end of CoverTab[29071]"
		}
//line /usr/local/go/src/text/template/parse/lex.go:504
		// _ = "end of CoverTab[29068]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:505
	// _ = "end of CoverTab[29067]"
}

// lexField scans a field: .Alphanumeric.
//line /usr/local/go/src/text/template/parse/lex.go:508
// The . has been scanned.
//line /usr/local/go/src/text/template/parse/lex.go:510
func lexField(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:510
	_go_fuzz_dep_.CoverTab[29084]++
								return lexFieldOrVariable(l, itemField)
//line /usr/local/go/src/text/template/parse/lex.go:511
	// _ = "end of CoverTab[29084]"
}

// lexVariable scans a Variable: $Alphanumeric.
//line /usr/local/go/src/text/template/parse/lex.go:514
// The $ has been scanned.
//line /usr/local/go/src/text/template/parse/lex.go:516
func lexVariable(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:516
	_go_fuzz_dep_.CoverTab[29085]++
								if l.atTerminator() {
//line /usr/local/go/src/text/template/parse/lex.go:517
		_go_fuzz_dep_.CoverTab[29087]++
									return l.emit(itemVariable)
//line /usr/local/go/src/text/template/parse/lex.go:518
		// _ = "end of CoverTab[29087]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:519
		_go_fuzz_dep_.CoverTab[29088]++
//line /usr/local/go/src/text/template/parse/lex.go:519
		// _ = "end of CoverTab[29088]"
//line /usr/local/go/src/text/template/parse/lex.go:519
	}
//line /usr/local/go/src/text/template/parse/lex.go:519
	// _ = "end of CoverTab[29085]"
//line /usr/local/go/src/text/template/parse/lex.go:519
	_go_fuzz_dep_.CoverTab[29086]++
								return lexFieldOrVariable(l, itemVariable)
//line /usr/local/go/src/text/template/parse/lex.go:520
	// _ = "end of CoverTab[29086]"
}

// lexFieldOrVariable scans a field or variable: [.$]Alphanumeric.
//line /usr/local/go/src/text/template/parse/lex.go:523
// The . or $ has been scanned.
//line /usr/local/go/src/text/template/parse/lex.go:525
func lexFieldOrVariable(l *lexer, typ itemType) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:525
	_go_fuzz_dep_.CoverTab[29089]++
								if l.atTerminator() {
//line /usr/local/go/src/text/template/parse/lex.go:526
		_go_fuzz_dep_.CoverTab[29093]++
									if typ == itemVariable {
//line /usr/local/go/src/text/template/parse/lex.go:527
			_go_fuzz_dep_.CoverTab[29095]++
										return l.emit(itemVariable)
//line /usr/local/go/src/text/template/parse/lex.go:528
			// _ = "end of CoverTab[29095]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:529
			_go_fuzz_dep_.CoverTab[29096]++
//line /usr/local/go/src/text/template/parse/lex.go:529
			// _ = "end of CoverTab[29096]"
//line /usr/local/go/src/text/template/parse/lex.go:529
		}
//line /usr/local/go/src/text/template/parse/lex.go:529
		// _ = "end of CoverTab[29093]"
//line /usr/local/go/src/text/template/parse/lex.go:529
		_go_fuzz_dep_.CoverTab[29094]++
									return l.emit(itemDot)
//line /usr/local/go/src/text/template/parse/lex.go:530
		// _ = "end of CoverTab[29094]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:531
		_go_fuzz_dep_.CoverTab[29097]++
//line /usr/local/go/src/text/template/parse/lex.go:531
		// _ = "end of CoverTab[29097]"
//line /usr/local/go/src/text/template/parse/lex.go:531
	}
//line /usr/local/go/src/text/template/parse/lex.go:531
	// _ = "end of CoverTab[29089]"
//line /usr/local/go/src/text/template/parse/lex.go:531
	_go_fuzz_dep_.CoverTab[29090]++
								var r rune
								for {
//line /usr/local/go/src/text/template/parse/lex.go:533
		_go_fuzz_dep_.CoverTab[29098]++
									r = l.next()
									if !isAlphaNumeric(r) {
//line /usr/local/go/src/text/template/parse/lex.go:535
			_go_fuzz_dep_.CoverTab[29099]++
										l.backup()
										break
//line /usr/local/go/src/text/template/parse/lex.go:537
			// _ = "end of CoverTab[29099]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:538
			_go_fuzz_dep_.CoverTab[29100]++
//line /usr/local/go/src/text/template/parse/lex.go:538
			// _ = "end of CoverTab[29100]"
//line /usr/local/go/src/text/template/parse/lex.go:538
		}
//line /usr/local/go/src/text/template/parse/lex.go:538
		// _ = "end of CoverTab[29098]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:539
	// _ = "end of CoverTab[29090]"
//line /usr/local/go/src/text/template/parse/lex.go:539
	_go_fuzz_dep_.CoverTab[29091]++
								if !l.atTerminator() {
//line /usr/local/go/src/text/template/parse/lex.go:540
		_go_fuzz_dep_.CoverTab[29101]++
									return l.errorf("bad character %#U", r)
//line /usr/local/go/src/text/template/parse/lex.go:541
		// _ = "end of CoverTab[29101]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:542
		_go_fuzz_dep_.CoverTab[29102]++
//line /usr/local/go/src/text/template/parse/lex.go:542
		// _ = "end of CoverTab[29102]"
//line /usr/local/go/src/text/template/parse/lex.go:542
	}
//line /usr/local/go/src/text/template/parse/lex.go:542
	// _ = "end of CoverTab[29091]"
//line /usr/local/go/src/text/template/parse/lex.go:542
	_go_fuzz_dep_.CoverTab[29092]++
								return l.emit(typ)
//line /usr/local/go/src/text/template/parse/lex.go:543
	// _ = "end of CoverTab[29092]"
}

// atTerminator reports whether the input is at valid termination character to
//line /usr/local/go/src/text/template/parse/lex.go:546
// appear after an identifier. Breaks .X.Y into two pieces. Also catches cases
//line /usr/local/go/src/text/template/parse/lex.go:546
// like "$x+2" not being acceptable without a space, in case we decide one
//line /usr/local/go/src/text/template/parse/lex.go:546
// day to implement arithmetic.
//line /usr/local/go/src/text/template/parse/lex.go:550
func (l *lexer) atTerminator() bool {
//line /usr/local/go/src/text/template/parse/lex.go:550
	_go_fuzz_dep_.CoverTab[29103]++
								r := l.peek()
								if isSpace(r) {
//line /usr/local/go/src/text/template/parse/lex.go:552
		_go_fuzz_dep_.CoverTab[29106]++
									return true
//line /usr/local/go/src/text/template/parse/lex.go:553
		// _ = "end of CoverTab[29106]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:554
		_go_fuzz_dep_.CoverTab[29107]++
//line /usr/local/go/src/text/template/parse/lex.go:554
		// _ = "end of CoverTab[29107]"
//line /usr/local/go/src/text/template/parse/lex.go:554
	}
//line /usr/local/go/src/text/template/parse/lex.go:554
	// _ = "end of CoverTab[29103]"
//line /usr/local/go/src/text/template/parse/lex.go:554
	_go_fuzz_dep_.CoverTab[29104]++
								switch r {
	case eof, '.', ',', '|', ':', ')', '(':
//line /usr/local/go/src/text/template/parse/lex.go:556
		_go_fuzz_dep_.CoverTab[29108]++
									return true
//line /usr/local/go/src/text/template/parse/lex.go:557
		// _ = "end of CoverTab[29108]"
//line /usr/local/go/src/text/template/parse/lex.go:557
	default:
//line /usr/local/go/src/text/template/parse/lex.go:557
		_go_fuzz_dep_.CoverTab[29109]++
//line /usr/local/go/src/text/template/parse/lex.go:557
		// _ = "end of CoverTab[29109]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:558
	// _ = "end of CoverTab[29104]"
//line /usr/local/go/src/text/template/parse/lex.go:558
	_go_fuzz_dep_.CoverTab[29105]++
								return strings.HasPrefix(l.input[l.pos:], l.rightDelim)
//line /usr/local/go/src/text/template/parse/lex.go:559
	// _ = "end of CoverTab[29105]"
}

// lexChar scans a character constant. The initial quote is already
//line /usr/local/go/src/text/template/parse/lex.go:562
// scanned. Syntax checking is done by the parser.
//line /usr/local/go/src/text/template/parse/lex.go:564
func lexChar(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:564
	_go_fuzz_dep_.CoverTab[29110]++
Loop:
	for {
//line /usr/local/go/src/text/template/parse/lex.go:566
		_go_fuzz_dep_.CoverTab[29112]++
									switch l.next() {
		case '\\':
//line /usr/local/go/src/text/template/parse/lex.go:568
			_go_fuzz_dep_.CoverTab[29113]++
										if r := l.next(); r != eof && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:569
				_go_fuzz_dep_.CoverTab[29118]++
//line /usr/local/go/src/text/template/parse/lex.go:569
				return r != '\n'
//line /usr/local/go/src/text/template/parse/lex.go:569
				// _ = "end of CoverTab[29118]"
//line /usr/local/go/src/text/template/parse/lex.go:569
			}() {
//line /usr/local/go/src/text/template/parse/lex.go:569
				_go_fuzz_dep_.CoverTab[29119]++
											break
//line /usr/local/go/src/text/template/parse/lex.go:570
				// _ = "end of CoverTab[29119]"
			} else {
//line /usr/local/go/src/text/template/parse/lex.go:571
				_go_fuzz_dep_.CoverTab[29120]++
//line /usr/local/go/src/text/template/parse/lex.go:571
				// _ = "end of CoverTab[29120]"
//line /usr/local/go/src/text/template/parse/lex.go:571
			}
//line /usr/local/go/src/text/template/parse/lex.go:571
			// _ = "end of CoverTab[29113]"
//line /usr/local/go/src/text/template/parse/lex.go:571
			_go_fuzz_dep_.CoverTab[29114]++
										fallthrough
//line /usr/local/go/src/text/template/parse/lex.go:572
			// _ = "end of CoverTab[29114]"
		case eof, '\n':
//line /usr/local/go/src/text/template/parse/lex.go:573
			_go_fuzz_dep_.CoverTab[29115]++
										return l.errorf("unterminated character constant")
//line /usr/local/go/src/text/template/parse/lex.go:574
			// _ = "end of CoverTab[29115]"
		case '\'':
//line /usr/local/go/src/text/template/parse/lex.go:575
			_go_fuzz_dep_.CoverTab[29116]++
										break Loop
//line /usr/local/go/src/text/template/parse/lex.go:576
			// _ = "end of CoverTab[29116]"
//line /usr/local/go/src/text/template/parse/lex.go:576
		default:
//line /usr/local/go/src/text/template/parse/lex.go:576
			_go_fuzz_dep_.CoverTab[29117]++
//line /usr/local/go/src/text/template/parse/lex.go:576
			// _ = "end of CoverTab[29117]"
		}
//line /usr/local/go/src/text/template/parse/lex.go:577
		// _ = "end of CoverTab[29112]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:578
	// _ = "end of CoverTab[29110]"
//line /usr/local/go/src/text/template/parse/lex.go:578
	_go_fuzz_dep_.CoverTab[29111]++
								return l.emit(itemCharConstant)
//line /usr/local/go/src/text/template/parse/lex.go:579
	// _ = "end of CoverTab[29111]"
}

// lexNumber scans a number: decimal, octal, hex, float, or imaginary. This
//line /usr/local/go/src/text/template/parse/lex.go:582
// isn't a perfect number scanner - for instance it accepts "." and "0x0.2"
//line /usr/local/go/src/text/template/parse/lex.go:582
// and "089" - but when it's wrong the input is invalid and the parser (via
//line /usr/local/go/src/text/template/parse/lex.go:582
// strconv) will notice.
//line /usr/local/go/src/text/template/parse/lex.go:586
func lexNumber(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:586
	_go_fuzz_dep_.CoverTab[29121]++
								if !l.scanNumber() {
//line /usr/local/go/src/text/template/parse/lex.go:587
		_go_fuzz_dep_.CoverTab[29124]++
									return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
//line /usr/local/go/src/text/template/parse/lex.go:588
		// _ = "end of CoverTab[29124]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:589
		_go_fuzz_dep_.CoverTab[29125]++
//line /usr/local/go/src/text/template/parse/lex.go:589
		// _ = "end of CoverTab[29125]"
//line /usr/local/go/src/text/template/parse/lex.go:589
	}
//line /usr/local/go/src/text/template/parse/lex.go:589
	// _ = "end of CoverTab[29121]"
//line /usr/local/go/src/text/template/parse/lex.go:589
	_go_fuzz_dep_.CoverTab[29122]++
								if sign := l.peek(); sign == '+' || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:590
		_go_fuzz_dep_.CoverTab[29126]++
//line /usr/local/go/src/text/template/parse/lex.go:590
		return sign == '-'
//line /usr/local/go/src/text/template/parse/lex.go:590
		// _ = "end of CoverTab[29126]"
//line /usr/local/go/src/text/template/parse/lex.go:590
	}() {
//line /usr/local/go/src/text/template/parse/lex.go:590
		_go_fuzz_dep_.CoverTab[29127]++

									if !l.scanNumber() || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:592
			_go_fuzz_dep_.CoverTab[29129]++
//line /usr/local/go/src/text/template/parse/lex.go:592
			return l.input[l.pos-1] != 'i'
//line /usr/local/go/src/text/template/parse/lex.go:592
			// _ = "end of CoverTab[29129]"
//line /usr/local/go/src/text/template/parse/lex.go:592
		}() {
//line /usr/local/go/src/text/template/parse/lex.go:592
			_go_fuzz_dep_.CoverTab[29130]++
										return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
//line /usr/local/go/src/text/template/parse/lex.go:593
			// _ = "end of CoverTab[29130]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:594
			_go_fuzz_dep_.CoverTab[29131]++
//line /usr/local/go/src/text/template/parse/lex.go:594
			// _ = "end of CoverTab[29131]"
//line /usr/local/go/src/text/template/parse/lex.go:594
		}
//line /usr/local/go/src/text/template/parse/lex.go:594
		// _ = "end of CoverTab[29127]"
//line /usr/local/go/src/text/template/parse/lex.go:594
		_go_fuzz_dep_.CoverTab[29128]++
									return l.emit(itemComplex)
//line /usr/local/go/src/text/template/parse/lex.go:595
		// _ = "end of CoverTab[29128]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:596
		_go_fuzz_dep_.CoverTab[29132]++
//line /usr/local/go/src/text/template/parse/lex.go:596
		// _ = "end of CoverTab[29132]"
//line /usr/local/go/src/text/template/parse/lex.go:596
	}
//line /usr/local/go/src/text/template/parse/lex.go:596
	// _ = "end of CoverTab[29122]"
//line /usr/local/go/src/text/template/parse/lex.go:596
	_go_fuzz_dep_.CoverTab[29123]++
								return l.emit(itemNumber)
//line /usr/local/go/src/text/template/parse/lex.go:597
	// _ = "end of CoverTab[29123]"
}

func (l *lexer) scanNumber() bool {
//line /usr/local/go/src/text/template/parse/lex.go:600
	_go_fuzz_dep_.CoverTab[29133]++

								l.accept("+-")

								digits := "0123456789_"
								if l.accept("0") {
//line /usr/local/go/src/text/template/parse/lex.go:605
		_go_fuzz_dep_.CoverTab[29139]++

									if l.accept("xX") {
//line /usr/local/go/src/text/template/parse/lex.go:607
			_go_fuzz_dep_.CoverTab[29140]++
										digits = "0123456789abcdefABCDEF_"
//line /usr/local/go/src/text/template/parse/lex.go:608
			// _ = "end of CoverTab[29140]"
		} else {
//line /usr/local/go/src/text/template/parse/lex.go:609
			_go_fuzz_dep_.CoverTab[29141]++
//line /usr/local/go/src/text/template/parse/lex.go:609
			if l.accept("oO") {
//line /usr/local/go/src/text/template/parse/lex.go:609
				_go_fuzz_dep_.CoverTab[29142]++
											digits = "01234567_"
//line /usr/local/go/src/text/template/parse/lex.go:610
				// _ = "end of CoverTab[29142]"
			} else {
//line /usr/local/go/src/text/template/parse/lex.go:611
				_go_fuzz_dep_.CoverTab[29143]++
//line /usr/local/go/src/text/template/parse/lex.go:611
				if l.accept("bB") {
//line /usr/local/go/src/text/template/parse/lex.go:611
					_go_fuzz_dep_.CoverTab[29144]++
												digits = "01_"
//line /usr/local/go/src/text/template/parse/lex.go:612
					// _ = "end of CoverTab[29144]"
				} else {
//line /usr/local/go/src/text/template/parse/lex.go:613
					_go_fuzz_dep_.CoverTab[29145]++
//line /usr/local/go/src/text/template/parse/lex.go:613
					// _ = "end of CoverTab[29145]"
//line /usr/local/go/src/text/template/parse/lex.go:613
				}
//line /usr/local/go/src/text/template/parse/lex.go:613
				// _ = "end of CoverTab[29143]"
//line /usr/local/go/src/text/template/parse/lex.go:613
			}
//line /usr/local/go/src/text/template/parse/lex.go:613
			// _ = "end of CoverTab[29141]"
//line /usr/local/go/src/text/template/parse/lex.go:613
		}
//line /usr/local/go/src/text/template/parse/lex.go:613
		// _ = "end of CoverTab[29139]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:614
		_go_fuzz_dep_.CoverTab[29146]++
//line /usr/local/go/src/text/template/parse/lex.go:614
		// _ = "end of CoverTab[29146]"
//line /usr/local/go/src/text/template/parse/lex.go:614
	}
//line /usr/local/go/src/text/template/parse/lex.go:614
	// _ = "end of CoverTab[29133]"
//line /usr/local/go/src/text/template/parse/lex.go:614
	_go_fuzz_dep_.CoverTab[29134]++
								l.acceptRun(digits)
								if l.accept(".") {
//line /usr/local/go/src/text/template/parse/lex.go:616
		_go_fuzz_dep_.CoverTab[29147]++
									l.acceptRun(digits)
//line /usr/local/go/src/text/template/parse/lex.go:617
		// _ = "end of CoverTab[29147]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:618
		_go_fuzz_dep_.CoverTab[29148]++
//line /usr/local/go/src/text/template/parse/lex.go:618
		// _ = "end of CoverTab[29148]"
//line /usr/local/go/src/text/template/parse/lex.go:618
	}
//line /usr/local/go/src/text/template/parse/lex.go:618
	// _ = "end of CoverTab[29134]"
//line /usr/local/go/src/text/template/parse/lex.go:618
	_go_fuzz_dep_.CoverTab[29135]++
								if len(digits) == 10+1 && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:619
		_go_fuzz_dep_.CoverTab[29149]++
//line /usr/local/go/src/text/template/parse/lex.go:619
		return l.accept("eE")
//line /usr/local/go/src/text/template/parse/lex.go:619
		// _ = "end of CoverTab[29149]"
//line /usr/local/go/src/text/template/parse/lex.go:619
	}() {
//line /usr/local/go/src/text/template/parse/lex.go:619
		_go_fuzz_dep_.CoverTab[29150]++
									l.accept("+-")
									l.acceptRun("0123456789_")
//line /usr/local/go/src/text/template/parse/lex.go:621
		// _ = "end of CoverTab[29150]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:622
		_go_fuzz_dep_.CoverTab[29151]++
//line /usr/local/go/src/text/template/parse/lex.go:622
		// _ = "end of CoverTab[29151]"
//line /usr/local/go/src/text/template/parse/lex.go:622
	}
//line /usr/local/go/src/text/template/parse/lex.go:622
	// _ = "end of CoverTab[29135]"
//line /usr/local/go/src/text/template/parse/lex.go:622
	_go_fuzz_dep_.CoverTab[29136]++
								if len(digits) == 16+6+1 && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:623
		_go_fuzz_dep_.CoverTab[29152]++
//line /usr/local/go/src/text/template/parse/lex.go:623
		return l.accept("pP")
//line /usr/local/go/src/text/template/parse/lex.go:623
		// _ = "end of CoverTab[29152]"
//line /usr/local/go/src/text/template/parse/lex.go:623
	}() {
//line /usr/local/go/src/text/template/parse/lex.go:623
		_go_fuzz_dep_.CoverTab[29153]++
									l.accept("+-")
									l.acceptRun("0123456789_")
//line /usr/local/go/src/text/template/parse/lex.go:625
		// _ = "end of CoverTab[29153]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:626
		_go_fuzz_dep_.CoverTab[29154]++
//line /usr/local/go/src/text/template/parse/lex.go:626
		// _ = "end of CoverTab[29154]"
//line /usr/local/go/src/text/template/parse/lex.go:626
	}
//line /usr/local/go/src/text/template/parse/lex.go:626
	// _ = "end of CoverTab[29136]"
//line /usr/local/go/src/text/template/parse/lex.go:626
	_go_fuzz_dep_.CoverTab[29137]++

								l.accept("i")

								if isAlphaNumeric(l.peek()) {
//line /usr/local/go/src/text/template/parse/lex.go:630
		_go_fuzz_dep_.CoverTab[29155]++
									l.next()
									return false
//line /usr/local/go/src/text/template/parse/lex.go:632
		// _ = "end of CoverTab[29155]"
	} else {
//line /usr/local/go/src/text/template/parse/lex.go:633
		_go_fuzz_dep_.CoverTab[29156]++
//line /usr/local/go/src/text/template/parse/lex.go:633
		// _ = "end of CoverTab[29156]"
//line /usr/local/go/src/text/template/parse/lex.go:633
	}
//line /usr/local/go/src/text/template/parse/lex.go:633
	// _ = "end of CoverTab[29137]"
//line /usr/local/go/src/text/template/parse/lex.go:633
	_go_fuzz_dep_.CoverTab[29138]++
								return true
//line /usr/local/go/src/text/template/parse/lex.go:634
	// _ = "end of CoverTab[29138]"
}

// lexQuote scans a quoted string.
func lexQuote(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:638
	_go_fuzz_dep_.CoverTab[29157]++
Loop:
	for {
//line /usr/local/go/src/text/template/parse/lex.go:640
		_go_fuzz_dep_.CoverTab[29159]++
									switch l.next() {
		case '\\':
//line /usr/local/go/src/text/template/parse/lex.go:642
			_go_fuzz_dep_.CoverTab[29160]++
										if r := l.next(); r != eof && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:643
				_go_fuzz_dep_.CoverTab[29165]++
//line /usr/local/go/src/text/template/parse/lex.go:643
				return r != '\n'
//line /usr/local/go/src/text/template/parse/lex.go:643
				// _ = "end of CoverTab[29165]"
//line /usr/local/go/src/text/template/parse/lex.go:643
			}() {
//line /usr/local/go/src/text/template/parse/lex.go:643
				_go_fuzz_dep_.CoverTab[29166]++
											break
//line /usr/local/go/src/text/template/parse/lex.go:644
				// _ = "end of CoverTab[29166]"
			} else {
//line /usr/local/go/src/text/template/parse/lex.go:645
				_go_fuzz_dep_.CoverTab[29167]++
//line /usr/local/go/src/text/template/parse/lex.go:645
				// _ = "end of CoverTab[29167]"
//line /usr/local/go/src/text/template/parse/lex.go:645
			}
//line /usr/local/go/src/text/template/parse/lex.go:645
			// _ = "end of CoverTab[29160]"
//line /usr/local/go/src/text/template/parse/lex.go:645
			_go_fuzz_dep_.CoverTab[29161]++
										fallthrough
//line /usr/local/go/src/text/template/parse/lex.go:646
			// _ = "end of CoverTab[29161]"
		case eof, '\n':
//line /usr/local/go/src/text/template/parse/lex.go:647
			_go_fuzz_dep_.CoverTab[29162]++
										return l.errorf("unterminated quoted string")
//line /usr/local/go/src/text/template/parse/lex.go:648
			// _ = "end of CoverTab[29162]"
		case '"':
//line /usr/local/go/src/text/template/parse/lex.go:649
			_go_fuzz_dep_.CoverTab[29163]++
										break Loop
//line /usr/local/go/src/text/template/parse/lex.go:650
			// _ = "end of CoverTab[29163]"
//line /usr/local/go/src/text/template/parse/lex.go:650
		default:
//line /usr/local/go/src/text/template/parse/lex.go:650
			_go_fuzz_dep_.CoverTab[29164]++
//line /usr/local/go/src/text/template/parse/lex.go:650
			// _ = "end of CoverTab[29164]"
		}
//line /usr/local/go/src/text/template/parse/lex.go:651
		// _ = "end of CoverTab[29159]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:652
	// _ = "end of CoverTab[29157]"
//line /usr/local/go/src/text/template/parse/lex.go:652
	_go_fuzz_dep_.CoverTab[29158]++
								return l.emit(itemString)
//line /usr/local/go/src/text/template/parse/lex.go:653
	// _ = "end of CoverTab[29158]"
}

// lexRawQuote scans a raw quoted string.
func lexRawQuote(l *lexer) stateFn {
//line /usr/local/go/src/text/template/parse/lex.go:657
	_go_fuzz_dep_.CoverTab[29168]++
Loop:
	for {
//line /usr/local/go/src/text/template/parse/lex.go:659
		_go_fuzz_dep_.CoverTab[29170]++
									switch l.next() {
		case eof:
//line /usr/local/go/src/text/template/parse/lex.go:661
			_go_fuzz_dep_.CoverTab[29171]++
										return l.errorf("unterminated raw quoted string")
//line /usr/local/go/src/text/template/parse/lex.go:662
			// _ = "end of CoverTab[29171]"
		case '`':
//line /usr/local/go/src/text/template/parse/lex.go:663
			_go_fuzz_dep_.CoverTab[29172]++
										break Loop
//line /usr/local/go/src/text/template/parse/lex.go:664
			// _ = "end of CoverTab[29172]"
//line /usr/local/go/src/text/template/parse/lex.go:664
		default:
//line /usr/local/go/src/text/template/parse/lex.go:664
			_go_fuzz_dep_.CoverTab[29173]++
//line /usr/local/go/src/text/template/parse/lex.go:664
			// _ = "end of CoverTab[29173]"
		}
//line /usr/local/go/src/text/template/parse/lex.go:665
		// _ = "end of CoverTab[29170]"
	}
//line /usr/local/go/src/text/template/parse/lex.go:666
	// _ = "end of CoverTab[29168]"
//line /usr/local/go/src/text/template/parse/lex.go:666
	_go_fuzz_dep_.CoverTab[29169]++
								return l.emit(itemRawString)
//line /usr/local/go/src/text/template/parse/lex.go:667
	// _ = "end of CoverTab[29169]"
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
//line /usr/local/go/src/text/template/parse/lex.go:671
	_go_fuzz_dep_.CoverTab[29174]++
								return r == ' ' || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:672
		_go_fuzz_dep_.CoverTab[29175]++
//line /usr/local/go/src/text/template/parse/lex.go:672
		return r == '\t'
//line /usr/local/go/src/text/template/parse/lex.go:672
		// _ = "end of CoverTab[29175]"
//line /usr/local/go/src/text/template/parse/lex.go:672
	}() || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:672
		_go_fuzz_dep_.CoverTab[29176]++
//line /usr/local/go/src/text/template/parse/lex.go:672
		return r == '\r'
//line /usr/local/go/src/text/template/parse/lex.go:672
		// _ = "end of CoverTab[29176]"
//line /usr/local/go/src/text/template/parse/lex.go:672
	}() || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:672
		_go_fuzz_dep_.CoverTab[29177]++
//line /usr/local/go/src/text/template/parse/lex.go:672
		return r == '\n'
//line /usr/local/go/src/text/template/parse/lex.go:672
		// _ = "end of CoverTab[29177]"
//line /usr/local/go/src/text/template/parse/lex.go:672
	}()
//line /usr/local/go/src/text/template/parse/lex.go:672
	// _ = "end of CoverTab[29174]"
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
//line /usr/local/go/src/text/template/parse/lex.go:676
	_go_fuzz_dep_.CoverTab[29178]++
								return r == '_' || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:677
		_go_fuzz_dep_.CoverTab[29179]++
//line /usr/local/go/src/text/template/parse/lex.go:677
		return unicode.IsLetter(r)
//line /usr/local/go/src/text/template/parse/lex.go:677
		// _ = "end of CoverTab[29179]"
//line /usr/local/go/src/text/template/parse/lex.go:677
	}() || func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:677
		_go_fuzz_dep_.CoverTab[29180]++
//line /usr/local/go/src/text/template/parse/lex.go:677
		return unicode.IsDigit(r)
//line /usr/local/go/src/text/template/parse/lex.go:677
		// _ = "end of CoverTab[29180]"
//line /usr/local/go/src/text/template/parse/lex.go:677
	}()
//line /usr/local/go/src/text/template/parse/lex.go:677
	// _ = "end of CoverTab[29178]"
}

func hasLeftTrimMarker(s string) bool {
//line /usr/local/go/src/text/template/parse/lex.go:680
	_go_fuzz_dep_.CoverTab[29181]++
								return len(s) >= 2 && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:681
		_go_fuzz_dep_.CoverTab[29182]++
//line /usr/local/go/src/text/template/parse/lex.go:681
		return s[0] == trimMarker
//line /usr/local/go/src/text/template/parse/lex.go:681
		// _ = "end of CoverTab[29182]"
//line /usr/local/go/src/text/template/parse/lex.go:681
	}() && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:681
		_go_fuzz_dep_.CoverTab[29183]++
//line /usr/local/go/src/text/template/parse/lex.go:681
		return isSpace(rune(s[1]))
//line /usr/local/go/src/text/template/parse/lex.go:681
		// _ = "end of CoverTab[29183]"
//line /usr/local/go/src/text/template/parse/lex.go:681
	}()
//line /usr/local/go/src/text/template/parse/lex.go:681
	// _ = "end of CoverTab[29181]"
}

func hasRightTrimMarker(s string) bool {
//line /usr/local/go/src/text/template/parse/lex.go:684
	_go_fuzz_dep_.CoverTab[29184]++
								return len(s) >= 2 && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:685
		_go_fuzz_dep_.CoverTab[29185]++
//line /usr/local/go/src/text/template/parse/lex.go:685
		return isSpace(rune(s[0]))
//line /usr/local/go/src/text/template/parse/lex.go:685
		// _ = "end of CoverTab[29185]"
//line /usr/local/go/src/text/template/parse/lex.go:685
	}() && func() bool {
//line /usr/local/go/src/text/template/parse/lex.go:685
		_go_fuzz_dep_.CoverTab[29186]++
//line /usr/local/go/src/text/template/parse/lex.go:685
		return s[1] == trimMarker
//line /usr/local/go/src/text/template/parse/lex.go:685
		// _ = "end of CoverTab[29186]"
//line /usr/local/go/src/text/template/parse/lex.go:685
	}()
//line /usr/local/go/src/text/template/parse/lex.go:685
	// _ = "end of CoverTab[29184]"
}

//line /usr/local/go/src/text/template/parse/lex.go:686
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/text/template/parse/lex.go:686
var _ = _go_fuzz_dep_.CoverTab
