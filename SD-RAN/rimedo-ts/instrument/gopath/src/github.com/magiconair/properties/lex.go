// Copyright 2018 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Parts of the lexer are from the template/text/parser package
// For these parts the following applies:
//
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file of the go 1.2
// distribution.

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:13
package properties

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:13
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:13
)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:13
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:13
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:13
)

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// item represents a token or text string returned from the scanner.
type item struct {
	typ	itemType	// The type of this item.
	pos	int		// The starting position, in bytes, of this item in the input string.
	val	string		// The value of this item.
}

func (i item) String() string {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:29
	_go_fuzz_dep_.CoverTab[115643]++
											switch {
	case i.typ == itemEOF:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:31
		_go_fuzz_dep_.CoverTab[115645]++
												return "EOF"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:32
		// _ = "end of CoverTab[115645]"
	case i.typ == itemError:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:33
		_go_fuzz_dep_.CoverTab[115646]++
												return i.val
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:34
		// _ = "end of CoverTab[115646]"
	case len(i.val) > 10:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:35
		_go_fuzz_dep_.CoverTab[115647]++
												return fmt.Sprintf("%.10q...", i.val)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:36
		// _ = "end of CoverTab[115647]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:36
	default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:36
		_go_fuzz_dep_.CoverTab[115648]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:36
		// _ = "end of CoverTab[115648]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:37
	// _ = "end of CoverTab[115643]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:37
	_go_fuzz_dep_.CoverTab[115644]++
											return fmt.Sprintf("%q", i.val)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:38
	// _ = "end of CoverTab[115644]"
}

// itemType identifies the type of lex items.
type itemType int

const (
	itemError	itemType	= iota	// error occurred; value is text of error
	itemEOF
	itemKey		// a key
	itemValue	// a value
	itemComment	// a comment
)

// defines a constant for EOF
const eof = -1

// permitted whitespace characters space, FF and TAB
const whitespace = " \f\t"

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	input	string		// the string being scanned
	state	stateFn		// the next lexing function to enter
	pos	int		// current position in the input
	start	int		// start position of this item
	width	int		// width of last rune read from input
	lastPos	int		// position of most recent item returned by nextItem
	runes	[]rune		// scanned runes for this item
	items	chan item	// channel of scanned items
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:74
	_go_fuzz_dep_.CoverTab[115649]++
											if l.pos >= len(l.input) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:75
		_go_fuzz_dep_.CoverTab[115651]++
												l.width = 0
												return eof
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:77
		// _ = "end of CoverTab[115651]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:78
		_go_fuzz_dep_.CoverTab[115652]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:78
		// _ = "end of CoverTab[115652]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:78
	// _ = "end of CoverTab[115649]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:78
	_go_fuzz_dep_.CoverTab[115650]++
											r, w := utf8.DecodeRuneInString(l.input[l.pos:])
											l.width = w
											l.pos += l.width
											return r
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:82
	// _ = "end of CoverTab[115650]"
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:86
	_go_fuzz_dep_.CoverTab[115653]++
											r := l.next()
											l.backup()
											return r
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:89
	// _ = "end of CoverTab[115653]"
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:93
	_go_fuzz_dep_.CoverTab[115654]++
											l.pos -= l.width
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:94
	// _ = "end of CoverTab[115654]"
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:98
	_go_fuzz_dep_.CoverTab[115655]++
											i := item{t, l.start, string(l.runes)}
											l.items <- i
											l.start = l.pos
											l.runes = l.runes[:0]
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:102
	// _ = "end of CoverTab[115655]"
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:106
	_go_fuzz_dep_.CoverTab[115656]++
											l.start = l.pos
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:107
	// _ = "end of CoverTab[115656]"
}

// appends the rune to the current value
func (l *lexer) appendRune(r rune) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:111
	_go_fuzz_dep_.CoverTab[115657]++
											l.runes = append(l.runes, r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:112
	// _ = "end of CoverTab[115657]"
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:116
	_go_fuzz_dep_.CoverTab[115658]++
											if strings.ContainsRune(valid, l.next()) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:117
		_go_fuzz_dep_.CoverTab[115660]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:118
		// _ = "end of CoverTab[115660]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:119
		_go_fuzz_dep_.CoverTab[115661]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:119
		// _ = "end of CoverTab[115661]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:119
	// _ = "end of CoverTab[115658]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:119
	_go_fuzz_dep_.CoverTab[115659]++
											l.backup()
											return false
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:121
	// _ = "end of CoverTab[115659]"
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:125
	_go_fuzz_dep_.CoverTab[115662]++
											for strings.ContainsRune(valid, l.next()) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:126
		_go_fuzz_dep_.CoverTab[115664]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:126
		// _ = "end of CoverTab[115664]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:127
	// _ = "end of CoverTab[115662]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:127
	_go_fuzz_dep_.CoverTab[115663]++
											l.backup()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:128
	// _ = "end of CoverTab[115663]"
}

// acceptRunUntil consumes a run of runes up to a terminator.
func (l *lexer) acceptRunUntil(term rune) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:132
	_go_fuzz_dep_.CoverTab[115665]++
											for term != l.next() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:133
		_go_fuzz_dep_.CoverTab[115667]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:133
		// _ = "end of CoverTab[115667]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:134
	// _ = "end of CoverTab[115665]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:134
	_go_fuzz_dep_.CoverTab[115666]++
											l.backup()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:135
	// _ = "end of CoverTab[115666]"
}

// hasText returns true if the current parsed text is not empty.
func (l *lexer) isNotEmpty() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:139
	_go_fuzz_dep_.CoverTab[115668]++
											return l.pos > l.start
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:140
	// _ = "end of CoverTab[115668]"
}

// lineNumber reports which line we're on, based on the position of
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:143
// the previous item returned by nextItem. Doing it this way
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:143
// means we don't have to worry about peek double counting.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:146
func (l *lexer) lineNumber() int {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:146
	_go_fuzz_dep_.CoverTab[115669]++
											return 1 + strings.Count(l.input[:l.lastPos], "\n")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:147
	// _ = "end of CoverTab[115669]"
}

// errorf returns an error token and terminates the scan by passing
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:150
// back a nil pointer that will be the next state, terminating l.nextItem.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:152
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:152
	_go_fuzz_dep_.CoverTab[115670]++
											l.items <- item{itemError, l.start, fmt.Sprintf(format, args...)}
											return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:154
	// _ = "end of CoverTab[115670]"
}

// nextItem returns the next item from the input.
func (l *lexer) nextItem() item {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:158
	_go_fuzz_dep_.CoverTab[115671]++
											i := <-l.items
											l.lastPos = i.pos
											return i
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:161
	// _ = "end of CoverTab[115671]"
}

// lex creates a new scanner for the input string.
func lex(input string) *lexer {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:165
	_go_fuzz_dep_.CoverTab[115672]++
											l := &lexer{
		input:	input,
		items:	make(chan item),
		runes:	make([]rune, 0, 32),
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:170
	_curRoutineNum153_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:170
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum153_)
											go l.run()
											return l
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:172
	// _ = "end of CoverTab[115672]"
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:176
	_go_fuzz_dep_.CoverTab[115673]++
											for l.state = lexBeforeKey(l); l.state != nil; {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:177
		_go_fuzz_dep_.CoverTab[115674]++
												l.state = l.state(l)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:178
		// _ = "end of CoverTab[115674]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:179
	// _ = "end of CoverTab[115673]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:184
// lexBeforeKey scans until a key begins.
func lexBeforeKey(l *lexer) stateFn {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:185
	_go_fuzz_dep_.CoverTab[115675]++
											switch r := l.next(); {
	case isEOF(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:187
		_go_fuzz_dep_.CoverTab[115676]++
												l.emit(itemEOF)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:189
		// _ = "end of CoverTab[115676]"

	case isEOL(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:191
		_go_fuzz_dep_.CoverTab[115677]++
												l.ignore()
												return lexBeforeKey
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:193
		// _ = "end of CoverTab[115677]"

	case isComment(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:195
		_go_fuzz_dep_.CoverTab[115678]++
												return lexComment
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:196
		// _ = "end of CoverTab[115678]"

	case isWhitespace(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:198
		_go_fuzz_dep_.CoverTab[115679]++
												l.ignore()
												return lexBeforeKey
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:200
		// _ = "end of CoverTab[115679]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:202
		_go_fuzz_dep_.CoverTab[115680]++
												l.backup()
												return lexKey
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:204
		// _ = "end of CoverTab[115680]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:205
	// _ = "end of CoverTab[115675]"
}

// lexComment scans a comment line. The comment character has already been scanned.
func lexComment(l *lexer) stateFn {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:209
	_go_fuzz_dep_.CoverTab[115681]++
											l.acceptRun(whitespace)
											l.ignore()
											for {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:212
		_go_fuzz_dep_.CoverTab[115682]++
												switch r := l.next(); {
		case isEOF(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:214
			_go_fuzz_dep_.CoverTab[115683]++
													l.ignore()
													l.emit(itemEOF)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:217
			// _ = "end of CoverTab[115683]"
		case isEOL(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:218
			_go_fuzz_dep_.CoverTab[115684]++
													l.emit(itemComment)
													return lexBeforeKey
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:220
			// _ = "end of CoverTab[115684]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:221
			_go_fuzz_dep_.CoverTab[115685]++
													l.appendRune(r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:222
			// _ = "end of CoverTab[115685]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:223
		// _ = "end of CoverTab[115682]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:224
	// _ = "end of CoverTab[115681]"
}

// lexKey scans the key up to a delimiter
func lexKey(l *lexer) stateFn {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:228
	_go_fuzz_dep_.CoverTab[115686]++
											var r rune

Loop:
	for {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:232
		_go_fuzz_dep_.CoverTab[115690]++
												switch r = l.next(); {

		case isEscape(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:235
			_go_fuzz_dep_.CoverTab[115691]++
													err := l.scanEscapeSequence()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:237
				_go_fuzz_dep_.CoverTab[115695]++
														return l.errorf(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:238
				// _ = "end of CoverTab[115695]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:239
				_go_fuzz_dep_.CoverTab[115696]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:239
				// _ = "end of CoverTab[115696]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:239
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:239
			// _ = "end of CoverTab[115691]"

		case isEndOfKey(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:241
			_go_fuzz_dep_.CoverTab[115692]++
													l.backup()
													break Loop
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:243
			// _ = "end of CoverTab[115692]"

		case isEOF(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:245
			_go_fuzz_dep_.CoverTab[115693]++
													break Loop
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:246
			// _ = "end of CoverTab[115693]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:248
			_go_fuzz_dep_.CoverTab[115694]++
													l.appendRune(r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:249
			// _ = "end of CoverTab[115694]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:250
		// _ = "end of CoverTab[115690]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:251
	// _ = "end of CoverTab[115686]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:251
	_go_fuzz_dep_.CoverTab[115687]++

											if len(l.runes) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:253
		_go_fuzz_dep_.CoverTab[115697]++
												l.emit(itemKey)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:254
		// _ = "end of CoverTab[115697]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:255
		_go_fuzz_dep_.CoverTab[115698]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:255
		// _ = "end of CoverTab[115698]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:255
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:255
	// _ = "end of CoverTab[115687]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:255
	_go_fuzz_dep_.CoverTab[115688]++

											if isEOF(r) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:257
		_go_fuzz_dep_.CoverTab[115699]++
												l.emit(itemEOF)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:259
		// _ = "end of CoverTab[115699]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:260
		_go_fuzz_dep_.CoverTab[115700]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:260
		// _ = "end of CoverTab[115700]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:260
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:260
	// _ = "end of CoverTab[115688]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:260
	_go_fuzz_dep_.CoverTab[115689]++

											return lexBeforeValue
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:262
	// _ = "end of CoverTab[115689]"
}

// lexBeforeValue scans the delimiter between key and value.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:265
// Leading and trailing whitespace is ignored.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:265
// We expect to be just after the key.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:268
func lexBeforeValue(l *lexer) stateFn {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:268
	_go_fuzz_dep_.CoverTab[115701]++
											l.acceptRun(whitespace)
											l.accept(":=")
											l.acceptRun(whitespace)
											l.ignore()
											return lexValue
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:273
	// _ = "end of CoverTab[115701]"
}

// lexValue scans text until the end of the line. We expect to be just after the delimiter.
func lexValue(l *lexer) stateFn {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:277
	_go_fuzz_dep_.CoverTab[115702]++
											for {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:278
		_go_fuzz_dep_.CoverTab[115703]++
												switch r := l.next(); {
		case isEscape(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:280
			_go_fuzz_dep_.CoverTab[115704]++
													if isEOL(l.peek()) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:281
				_go_fuzz_dep_.CoverTab[115708]++
														l.next()
														l.acceptRun(whitespace)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:283
				// _ = "end of CoverTab[115708]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:284
				_go_fuzz_dep_.CoverTab[115709]++
														err := l.scanEscapeSequence()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:286
					_go_fuzz_dep_.CoverTab[115710]++
															return l.errorf(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:287
					// _ = "end of CoverTab[115710]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:288
					_go_fuzz_dep_.CoverTab[115711]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:288
					// _ = "end of CoverTab[115711]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:288
				}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:288
				// _ = "end of CoverTab[115709]"
			}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:289
			// _ = "end of CoverTab[115704]"

		case isEOL(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:291
			_go_fuzz_dep_.CoverTab[115705]++
													l.emit(itemValue)
													l.ignore()
													return lexBeforeKey
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:294
			// _ = "end of CoverTab[115705]"

		case isEOF(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:296
			_go_fuzz_dep_.CoverTab[115706]++
													l.emit(itemValue)
													l.emit(itemEOF)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:299
			// _ = "end of CoverTab[115706]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:301
			_go_fuzz_dep_.CoverTab[115707]++
													l.appendRune(r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:302
			// _ = "end of CoverTab[115707]"
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:303
		// _ = "end of CoverTab[115703]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:304
	// _ = "end of CoverTab[115702]"
}

// scanEscapeSequence scans either one of the escaped characters
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:307
// or a unicode literal. We expect to be after the escape character.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:309
func (l *lexer) scanEscapeSequence() error {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:309
	_go_fuzz_dep_.CoverTab[115712]++
											switch r := l.next(); {

	case isEscapedCharacter(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:312
		_go_fuzz_dep_.CoverTab[115713]++
												l.appendRune(decodeEscapedCharacter(r))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:314
		// _ = "end of CoverTab[115713]"

	case atUnicodeLiteral(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:316
		_go_fuzz_dep_.CoverTab[115714]++
												return l.scanUnicodeLiteral()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:317
		// _ = "end of CoverTab[115714]"

	case isEOF(r):
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:319
		_go_fuzz_dep_.CoverTab[115715]++
												return fmt.Errorf("premature EOF")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:320
		// _ = "end of CoverTab[115715]"

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:323
	default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:323
		_go_fuzz_dep_.CoverTab[115716]++
												l.appendRune(r)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:325
		// _ = "end of CoverTab[115716]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:326
	// _ = "end of CoverTab[115712]"
}

// scans a unicode literal in the form \uXXXX. We expect to be after the \u.
func (l *lexer) scanUnicodeLiteral() error {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:330
	_go_fuzz_dep_.CoverTab[115717]++

											d := make([]rune, 4)
											for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:333
		_go_fuzz_dep_.CoverTab[115720]++
												d[i] = l.next()
												if d[i] == eof || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:335
			_go_fuzz_dep_.CoverTab[115721]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:335
			return !strings.ContainsRune("0123456789abcdefABCDEF", d[i])
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:335
			// _ = "end of CoverTab[115721]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:335
		}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:335
			_go_fuzz_dep_.CoverTab[115722]++
													return fmt.Errorf("invalid unicode literal")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:336
			// _ = "end of CoverTab[115722]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:337
			_go_fuzz_dep_.CoverTab[115723]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:337
			// _ = "end of CoverTab[115723]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:337
		}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:337
		// _ = "end of CoverTab[115720]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:338
	// _ = "end of CoverTab[115717]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:338
	_go_fuzz_dep_.CoverTab[115718]++

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:341
	r, err := strconv.ParseInt(string(d), 16, 0)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:342
		_go_fuzz_dep_.CoverTab[115724]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:343
		// _ = "end of CoverTab[115724]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:344
		_go_fuzz_dep_.CoverTab[115725]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:344
		// _ = "end of CoverTab[115725]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:344
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:344
	// _ = "end of CoverTab[115718]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:344
	_go_fuzz_dep_.CoverTab[115719]++

											l.appendRune(rune(r))
											return nil
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:347
	// _ = "end of CoverTab[115719]"
}

// decodeEscapedCharacter returns the unescaped rune. We expect to be after the escape character.
func decodeEscapedCharacter(r rune) rune {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:351
	_go_fuzz_dep_.CoverTab[115726]++
											switch r {
	case 'f':
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:353
		_go_fuzz_dep_.CoverTab[115727]++
												return '\f'
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:354
		// _ = "end of CoverTab[115727]"
	case 'n':
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:355
		_go_fuzz_dep_.CoverTab[115728]++
												return '\n'
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:356
		// _ = "end of CoverTab[115728]"
	case 'r':
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:357
		_go_fuzz_dep_.CoverTab[115729]++
												return '\r'
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:358
		// _ = "end of CoverTab[115729]"
	case 't':
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:359
		_go_fuzz_dep_.CoverTab[115730]++
												return '\t'
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:360
		// _ = "end of CoverTab[115730]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:361
		_go_fuzz_dep_.CoverTab[115731]++
												return r
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:362
		// _ = "end of CoverTab[115731]"
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:363
	// _ = "end of CoverTab[115726]"
}

// atUnicodeLiteral reports whether we are at a unicode literal.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:366
// The escape character has already been consumed.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:368
func atUnicodeLiteral(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:368
	_go_fuzz_dep_.CoverTab[115732]++
											return r == 'u'
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:369
	// _ = "end of CoverTab[115732]"
}

// isComment reports whether we are at the start of a comment.
func isComment(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:373
	_go_fuzz_dep_.CoverTab[115733]++
											return r == '#' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:374
		_go_fuzz_dep_.CoverTab[115734]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:374
		return r == '!'
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:374
		// _ = "end of CoverTab[115734]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:374
	}()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:374
	// _ = "end of CoverTab[115733]"
}

// isEndOfKey reports whether the rune terminates the current key.
func isEndOfKey(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:378
	_go_fuzz_dep_.CoverTab[115735]++
											return strings.ContainsRune(" \f\t\r\n:=", r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:379
	// _ = "end of CoverTab[115735]"
}

// isEOF reports whether we are at EOF.
func isEOF(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:383
	_go_fuzz_dep_.CoverTab[115736]++
											return r == eof
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:384
	// _ = "end of CoverTab[115736]"
}

// isEOL reports whether we are at a new line character.
func isEOL(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:388
	_go_fuzz_dep_.CoverTab[115737]++
											return r == '\n' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:389
		_go_fuzz_dep_.CoverTab[115738]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:389
		return r == '\r'
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:389
		// _ = "end of CoverTab[115738]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:389
	}()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:389
	// _ = "end of CoverTab[115737]"
}

// isEscape reports whether the rune is the escape character which
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:392
// prefixes unicode literals and other escaped characters.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:394
func isEscape(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:394
	_go_fuzz_dep_.CoverTab[115739]++
											return r == '\\'
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:395
	// _ = "end of CoverTab[115739]"
}

// isEscapedCharacter reports whether we are at one of the characters that need escaping.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:398
// The escape character has already been consumed.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:400
func isEscapedCharacter(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:400
	_go_fuzz_dep_.CoverTab[115740]++
											return strings.ContainsRune(" :=fnrt", r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:401
	// _ = "end of CoverTab[115740]"
}

// isWhitespace reports whether the rune is a whitespace character.
func isWhitespace(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:405
	_go_fuzz_dep_.CoverTab[115741]++
											return strings.ContainsRune(whitespace, r)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:406
	// _ = "end of CoverTab[115741]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:407
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/lex.go:407
var _ = _go_fuzz_dep_.CoverTab
