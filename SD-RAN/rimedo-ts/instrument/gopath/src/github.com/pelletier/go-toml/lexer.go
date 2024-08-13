// TOML lexer.
//
// Written using the principles developed by Rob Pike in
// http://www.youtube.com/watch?v=HxaD_trXwRE

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:6
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:6
)

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define state functions
type tomlLexStateFn func() tomlLexStateFn

// Define lexer
type tomlLexer struct {
	inputIdx		int
	input			[]rune	// Textual source
	currentTokenStart	int
	currentTokenStop	int
	tokens			[]token
	brackets		[]rune
	line			int
	col			int
	endbufferLine		int
	endbufferCol		int
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:35
func (l *tomlLexer) read() rune {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:35
	_go_fuzz_dep_.CoverTab[122581]++
											r := l.peek()
											if r == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:37
		_go_fuzz_dep_.CoverTab[122583]++
												l.endbufferLine++
												l.endbufferCol = 1
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:39
		// _ = "end of CoverTab[122583]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:40
		_go_fuzz_dep_.CoverTab[122584]++
												l.endbufferCol++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:41
		// _ = "end of CoverTab[122584]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:42
	// _ = "end of CoverTab[122581]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:42
	_go_fuzz_dep_.CoverTab[122582]++
											l.inputIdx++
											return r
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:44
	// _ = "end of CoverTab[122582]"
}

func (l *tomlLexer) next() rune {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:47
	_go_fuzz_dep_.CoverTab[122585]++
											r := l.read()

											if r != eof {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:50
		_go_fuzz_dep_.CoverTab[122587]++
												l.currentTokenStop++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:51
		// _ = "end of CoverTab[122587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:52
		_go_fuzz_dep_.CoverTab[122588]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:52
		// _ = "end of CoverTab[122588]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:52
	// _ = "end of CoverTab[122585]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:52
	_go_fuzz_dep_.CoverTab[122586]++
											return r
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:53
	// _ = "end of CoverTab[122586]"
}

func (l *tomlLexer) ignore() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:56
	_go_fuzz_dep_.CoverTab[122589]++
											l.currentTokenStart = l.currentTokenStop
											l.line = l.endbufferLine
											l.col = l.endbufferCol
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:59
	// _ = "end of CoverTab[122589]"
}

func (l *tomlLexer) skip() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:62
	_go_fuzz_dep_.CoverTab[122590]++
											l.next()
											l.ignore()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:64
	// _ = "end of CoverTab[122590]"
}

func (l *tomlLexer) fastForward(n int) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:67
	_go_fuzz_dep_.CoverTab[122591]++
											for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:68
		_go_fuzz_dep_.CoverTab[122592]++
												l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:69
		// _ = "end of CoverTab[122592]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:70
	// _ = "end of CoverTab[122591]"
}

func (l *tomlLexer) emitWithValue(t tokenType, value string) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:73
	_go_fuzz_dep_.CoverTab[122593]++
											l.tokens = append(l.tokens, token{
		Position:	Position{l.line, l.col},
		typ:		t,
		val:		value,
	})
											l.ignore()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:79
	// _ = "end of CoverTab[122593]"
}

func (l *tomlLexer) emit(t tokenType) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:82
	_go_fuzz_dep_.CoverTab[122594]++
											l.emitWithValue(t, string(l.input[l.currentTokenStart:l.currentTokenStop]))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:83
	// _ = "end of CoverTab[122594]"
}

func (l *tomlLexer) peek() rune {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:86
	_go_fuzz_dep_.CoverTab[122595]++
											if l.inputIdx >= len(l.input) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:87
		_go_fuzz_dep_.CoverTab[122597]++
												return eof
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:88
		// _ = "end of CoverTab[122597]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:89
		_go_fuzz_dep_.CoverTab[122598]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:89
		// _ = "end of CoverTab[122598]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:89
	// _ = "end of CoverTab[122595]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:89
	_go_fuzz_dep_.CoverTab[122596]++
											return l.input[l.inputIdx]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:90
	// _ = "end of CoverTab[122596]"
}

func (l *tomlLexer) peekString(size int) string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:93
	_go_fuzz_dep_.CoverTab[122599]++
											maxIdx := len(l.input)
											upperIdx := l.inputIdx + size
											if upperIdx > maxIdx {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:96
		_go_fuzz_dep_.CoverTab[122601]++
												upperIdx = maxIdx
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:97
		// _ = "end of CoverTab[122601]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:98
		_go_fuzz_dep_.CoverTab[122602]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:98
		// _ = "end of CoverTab[122602]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:98
	// _ = "end of CoverTab[122599]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:98
	_go_fuzz_dep_.CoverTab[122600]++
											return string(l.input[l.inputIdx:upperIdx])
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:99
	// _ = "end of CoverTab[122600]"
}

func (l *tomlLexer) follow(next string) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:102
	_go_fuzz_dep_.CoverTab[122603]++
											return next == l.peekString(len(next))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:103
	// _ = "end of CoverTab[122603]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:108
func (l *tomlLexer) errorf(format string, args ...interface{}) tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:108
	_go_fuzz_dep_.CoverTab[122604]++
											l.tokens = append(l.tokens, token{
		Position:	Position{l.line, l.col},
		typ:		tokenError,
		val:		fmt.Sprintf(format, args...),
	})
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:114
	// _ = "end of CoverTab[122604]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:119
func (l *tomlLexer) lexVoid() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:119
	_go_fuzz_dep_.CoverTab[122605]++
											for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:120
		_go_fuzz_dep_.CoverTab[122607]++
												next := l.peek()
												switch next {
		case '}':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:123
			_go_fuzz_dep_.CoverTab[122611]++
													return l.lexRightCurlyBrace
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:124
			// _ = "end of CoverTab[122611]"
		case '[':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:125
			_go_fuzz_dep_.CoverTab[122612]++
													return l.lexTableKey
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:126
			// _ = "end of CoverTab[122612]"
		case '#':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:127
			_go_fuzz_dep_.CoverTab[122613]++
													return l.lexComment(l.lexVoid)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:128
			// _ = "end of CoverTab[122613]"
		case '=':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:129
			_go_fuzz_dep_.CoverTab[122614]++
													return l.lexEqual
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:130
			// _ = "end of CoverTab[122614]"
		case '\r':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:131
			_go_fuzz_dep_.CoverTab[122615]++
													fallthrough
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:132
			// _ = "end of CoverTab[122615]"
		case '\n':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:133
			_go_fuzz_dep_.CoverTab[122616]++
													l.skip()
													continue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:135
			// _ = "end of CoverTab[122616]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:135
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:135
			_go_fuzz_dep_.CoverTab[122617]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:135
			// _ = "end of CoverTab[122617]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:136
		// _ = "end of CoverTab[122607]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:136
		_go_fuzz_dep_.CoverTab[122608]++

												if isSpace(next) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:138
			_go_fuzz_dep_.CoverTab[122618]++
													l.skip()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:139
			// _ = "end of CoverTab[122618]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:140
			_go_fuzz_dep_.CoverTab[122619]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:140
			// _ = "end of CoverTab[122619]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:140
		// _ = "end of CoverTab[122608]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:140
		_go_fuzz_dep_.CoverTab[122609]++

												if isKeyStartChar(next) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:142
			_go_fuzz_dep_.CoverTab[122620]++
													return l.lexKey
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:143
			// _ = "end of CoverTab[122620]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:144
			_go_fuzz_dep_.CoverTab[122621]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:144
			// _ = "end of CoverTab[122621]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:144
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:144
		// _ = "end of CoverTab[122609]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:144
		_go_fuzz_dep_.CoverTab[122610]++

												if next == eof {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:146
			_go_fuzz_dep_.CoverTab[122622]++
													l.next()
													break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:148
			// _ = "end of CoverTab[122622]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:149
			_go_fuzz_dep_.CoverTab[122623]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:149
			// _ = "end of CoverTab[122623]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:149
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:149
		// _ = "end of CoverTab[122610]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:150
	// _ = "end of CoverTab[122605]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:150
	_go_fuzz_dep_.CoverTab[122606]++

											l.emit(tokenEOF)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:153
	// _ = "end of CoverTab[122606]"
}

func (l *tomlLexer) lexRvalue() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:156
	_go_fuzz_dep_.CoverTab[122624]++
											for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:157
		_go_fuzz_dep_.CoverTab[122626]++
												next := l.peek()
												switch next {
		case '.':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:160
			_go_fuzz_dep_.CoverTab[122636]++
													return l.errorf("cannot start float with a dot")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:161
			// _ = "end of CoverTab[122636]"
		case '=':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:162
			_go_fuzz_dep_.CoverTab[122637]++
													return l.lexEqual
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:163
			// _ = "end of CoverTab[122637]"
		case '[':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:164
			_go_fuzz_dep_.CoverTab[122638]++
													return l.lexLeftBracket
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:165
			// _ = "end of CoverTab[122638]"
		case ']':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:166
			_go_fuzz_dep_.CoverTab[122639]++
													return l.lexRightBracket
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:167
			// _ = "end of CoverTab[122639]"
		case '{':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:168
			_go_fuzz_dep_.CoverTab[122640]++
													return l.lexLeftCurlyBrace
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:169
			// _ = "end of CoverTab[122640]"
		case '}':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:170
			_go_fuzz_dep_.CoverTab[122641]++
													return l.lexRightCurlyBrace
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:171
			// _ = "end of CoverTab[122641]"
		case '#':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:172
			_go_fuzz_dep_.CoverTab[122642]++
													return l.lexComment(l.lexRvalue)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:173
			// _ = "end of CoverTab[122642]"
		case '"':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:174
			_go_fuzz_dep_.CoverTab[122643]++
													return l.lexString
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:175
			// _ = "end of CoverTab[122643]"
		case '\'':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:176
			_go_fuzz_dep_.CoverTab[122644]++
													return l.lexLiteralString
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:177
			// _ = "end of CoverTab[122644]"
		case ',':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:178
			_go_fuzz_dep_.CoverTab[122645]++
													return l.lexComma
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:179
			// _ = "end of CoverTab[122645]"
		case '\r':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:180
			_go_fuzz_dep_.CoverTab[122646]++
													fallthrough
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:181
			// _ = "end of CoverTab[122646]"
		case '\n':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:182
			_go_fuzz_dep_.CoverTab[122647]++
													l.skip()
													if len(l.brackets) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:184
				_go_fuzz_dep_.CoverTab[122650]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:184
				return l.brackets[len(l.brackets)-1] == '['
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:184
				// _ = "end of CoverTab[122650]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:184
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:184
				_go_fuzz_dep_.CoverTab[122651]++
														return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:185
				// _ = "end of CoverTab[122651]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:186
				_go_fuzz_dep_.CoverTab[122652]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:186
				// _ = "end of CoverTab[122652]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:186
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:186
			// _ = "end of CoverTab[122647]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:186
			_go_fuzz_dep_.CoverTab[122648]++
													return l.lexVoid
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:187
			// _ = "end of CoverTab[122648]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:187
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:187
			_go_fuzz_dep_.CoverTab[122649]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:187
			// _ = "end of CoverTab[122649]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:188
		// _ = "end of CoverTab[122626]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:188
		_go_fuzz_dep_.CoverTab[122627]++

												if l.follow("true") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:190
			_go_fuzz_dep_.CoverTab[122653]++
													return l.lexTrue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:191
			// _ = "end of CoverTab[122653]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:192
			_go_fuzz_dep_.CoverTab[122654]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:192
			// _ = "end of CoverTab[122654]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:192
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:192
		// _ = "end of CoverTab[122627]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:192
		_go_fuzz_dep_.CoverTab[122628]++

												if l.follow("false") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:194
			_go_fuzz_dep_.CoverTab[122655]++
													return l.lexFalse
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:195
			// _ = "end of CoverTab[122655]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:196
			_go_fuzz_dep_.CoverTab[122656]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:196
			// _ = "end of CoverTab[122656]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:196
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:196
		// _ = "end of CoverTab[122628]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:196
		_go_fuzz_dep_.CoverTab[122629]++

												if l.follow("inf") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:198
			_go_fuzz_dep_.CoverTab[122657]++
													return l.lexInf
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:199
			// _ = "end of CoverTab[122657]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:200
			_go_fuzz_dep_.CoverTab[122658]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:200
			// _ = "end of CoverTab[122658]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:200
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:200
		// _ = "end of CoverTab[122629]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:200
		_go_fuzz_dep_.CoverTab[122630]++

												if l.follow("nan") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:202
			_go_fuzz_dep_.CoverTab[122659]++
													return l.lexNan
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:203
			// _ = "end of CoverTab[122659]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:204
			_go_fuzz_dep_.CoverTab[122660]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:204
			// _ = "end of CoverTab[122660]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:204
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:204
		// _ = "end of CoverTab[122630]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:204
		_go_fuzz_dep_.CoverTab[122631]++

												if isSpace(next) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:206
			_go_fuzz_dep_.CoverTab[122661]++
													l.skip()
													continue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:208
			// _ = "end of CoverTab[122661]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:209
			_go_fuzz_dep_.CoverTab[122662]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:209
			// _ = "end of CoverTab[122662]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:209
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:209
		// _ = "end of CoverTab[122631]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:209
		_go_fuzz_dep_.CoverTab[122632]++

												if next == eof {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:211
			_go_fuzz_dep_.CoverTab[122663]++
													l.next()
													break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:213
			// _ = "end of CoverTab[122663]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:214
			_go_fuzz_dep_.CoverTab[122664]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:214
			// _ = "end of CoverTab[122664]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:214
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:214
		// _ = "end of CoverTab[122632]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:214
		_go_fuzz_dep_.CoverTab[122633]++

												if next == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:216
			_go_fuzz_dep_.CoverTab[122665]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:216
			return next == '-'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:216
			// _ = "end of CoverTab[122665]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:216
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:216
			_go_fuzz_dep_.CoverTab[122666]++
													return l.lexNumber
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:217
			// _ = "end of CoverTab[122666]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:218
			_go_fuzz_dep_.CoverTab[122667]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:218
			// _ = "end of CoverTab[122667]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:218
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:218
		// _ = "end of CoverTab[122633]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:218
		_go_fuzz_dep_.CoverTab[122634]++

												if isDigit(next) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:220
			_go_fuzz_dep_.CoverTab[122668]++
													return l.lexDateTimeOrNumber
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:221
			// _ = "end of CoverTab[122668]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:222
			_go_fuzz_dep_.CoverTab[122669]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:222
			// _ = "end of CoverTab[122669]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:222
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:222
		// _ = "end of CoverTab[122634]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:222
		_go_fuzz_dep_.CoverTab[122635]++

												return l.errorf("no value can start with %c", next)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:224
		// _ = "end of CoverTab[122635]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:225
	// _ = "end of CoverTab[122624]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:225
	_go_fuzz_dep_.CoverTab[122625]++

											l.emit(tokenEOF)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:228
	// _ = "end of CoverTab[122625]"
}

func (l *tomlLexer) lexDateTimeOrNumber() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:231
	_go_fuzz_dep_.CoverTab[122670]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:238
	lookAhead := l.peekString(5)
	if len(lookAhead) < 3 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:239
		_go_fuzz_dep_.CoverTab[122673]++
												return l.lexNumber()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:240
		// _ = "end of CoverTab[122673]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:241
		_go_fuzz_dep_.CoverTab[122674]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:241
		// _ = "end of CoverTab[122674]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:241
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:241
	// _ = "end of CoverTab[122670]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:241
	_go_fuzz_dep_.CoverTab[122671]++

											for idx, r := range lookAhead {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:243
		_go_fuzz_dep_.CoverTab[122675]++
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:244
			_go_fuzz_dep_.CoverTab[122676]++
													if idx == 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:245
				_go_fuzz_dep_.CoverTab[122679]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:245
				return r == ':'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:245
				// _ = "end of CoverTab[122679]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:245
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:245
				_go_fuzz_dep_.CoverTab[122680]++
														return l.lexDateTimeOrTime()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:246
				// _ = "end of CoverTab[122680]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:247
				_go_fuzz_dep_.CoverTab[122681]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:247
				// _ = "end of CoverTab[122681]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:247
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:247
			// _ = "end of CoverTab[122676]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:247
			_go_fuzz_dep_.CoverTab[122677]++
													if idx == 4 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:248
				_go_fuzz_dep_.CoverTab[122682]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:248
				return r == '-'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:248
				// _ = "end of CoverTab[122682]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:248
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:248
				_go_fuzz_dep_.CoverTab[122683]++
														return l.lexDateTimeOrTime()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:249
				// _ = "end of CoverTab[122683]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:250
				_go_fuzz_dep_.CoverTab[122684]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:250
				// _ = "end of CoverTab[122684]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:250
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:250
			// _ = "end of CoverTab[122677]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:250
			_go_fuzz_dep_.CoverTab[122678]++
													return l.lexNumber()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:251
			// _ = "end of CoverTab[122678]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:252
			_go_fuzz_dep_.CoverTab[122685]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:252
			// _ = "end of CoverTab[122685]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:252
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:252
		// _ = "end of CoverTab[122675]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:253
	// _ = "end of CoverTab[122671]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:253
	_go_fuzz_dep_.CoverTab[122672]++
											return l.lexNumber()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:254
	// _ = "end of CoverTab[122672]"
}

func (l *tomlLexer) lexLeftCurlyBrace() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:257
	_go_fuzz_dep_.CoverTab[122686]++
											l.next()
											l.emit(tokenLeftCurlyBrace)
											l.brackets = append(l.brackets, '{')
											return l.lexVoid
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:261
	// _ = "end of CoverTab[122686]"
}

func (l *tomlLexer) lexRightCurlyBrace() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:264
	_go_fuzz_dep_.CoverTab[122687]++
											l.next()
											l.emit(tokenRightCurlyBrace)
											if len(l.brackets) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:267
		_go_fuzz_dep_.CoverTab[122689]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:267
		return l.brackets[len(l.brackets)-1] != '{'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:267
		// _ = "end of CoverTab[122689]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:267
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:267
		_go_fuzz_dep_.CoverTab[122690]++
												return l.errorf("cannot have '}' here")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:268
		// _ = "end of CoverTab[122690]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:269
		_go_fuzz_dep_.CoverTab[122691]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:269
		// _ = "end of CoverTab[122691]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:269
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:269
	// _ = "end of CoverTab[122687]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:269
	_go_fuzz_dep_.CoverTab[122688]++
											l.brackets = l.brackets[:len(l.brackets)-1]
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:271
	// _ = "end of CoverTab[122688]"
}

func (l *tomlLexer) lexDateTimeOrTime() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:274
	_go_fuzz_dep_.CoverTab[122692]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:291
	l.next()
											l.next()

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:296
	r := l.next()
	if r == ':' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:297
		_go_fuzz_dep_.CoverTab[122694]++
												return l.lexTime()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:298
		// _ = "end of CoverTab[122694]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:299
		_go_fuzz_dep_.CoverTab[122695]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:299
		// _ = "end of CoverTab[122695]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:299
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:299
	// _ = "end of CoverTab[122692]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:299
	_go_fuzz_dep_.CoverTab[122693]++

											return l.lexDateTime()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:301
	// _ = "end of CoverTab[122693]"
}

func (l *tomlLexer) lexDateTime() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:304
	_go_fuzz_dep_.CoverTab[122696]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:323
	l.next()
	l.next()

	for i := 0; i < 2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:326
		_go_fuzz_dep_.CoverTab[122709]++
												r := l.next()
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:328
			_go_fuzz_dep_.CoverTab[122710]++
													return l.errorf("invalid month digit in date: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:329
			// _ = "end of CoverTab[122710]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:330
			_go_fuzz_dep_.CoverTab[122711]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:330
			// _ = "end of CoverTab[122711]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:330
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:330
		// _ = "end of CoverTab[122709]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:331
	// _ = "end of CoverTab[122696]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:331
	_go_fuzz_dep_.CoverTab[122697]++

											r := l.next()
											if r != '-' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:334
		_go_fuzz_dep_.CoverTab[122712]++
												return l.errorf("expected - to separate month of a date, not %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:335
		// _ = "end of CoverTab[122712]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:336
		_go_fuzz_dep_.CoverTab[122713]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:336
		// _ = "end of CoverTab[122713]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:336
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:336
	// _ = "end of CoverTab[122697]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:336
	_go_fuzz_dep_.CoverTab[122698]++

											for i := 0; i < 2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:338
		_go_fuzz_dep_.CoverTab[122714]++
												r := l.next()
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:340
			_go_fuzz_dep_.CoverTab[122715]++
													return l.errorf("invalid day digit in date: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:341
			// _ = "end of CoverTab[122715]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:342
			_go_fuzz_dep_.CoverTab[122716]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:342
			// _ = "end of CoverTab[122716]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:342
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:342
		// _ = "end of CoverTab[122714]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:343
	// _ = "end of CoverTab[122698]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:343
	_go_fuzz_dep_.CoverTab[122699]++

											l.emit(tokenLocalDate)

											r = l.peek()

											if r == eof {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:349
		_go_fuzz_dep_.CoverTab[122717]++

												return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:351
		// _ = "end of CoverTab[122717]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:352
		_go_fuzz_dep_.CoverTab[122718]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:352
		// _ = "end of CoverTab[122718]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:352
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:352
	// _ = "end of CoverTab[122699]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:352
	_go_fuzz_dep_.CoverTab[122700]++

											if r != ' ' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:354
		_go_fuzz_dep_.CoverTab[122719]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:354
		return r != 'T'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:354
		// _ = "end of CoverTab[122719]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:354
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:354
		_go_fuzz_dep_.CoverTab[122720]++
												return l.errorf("incorrect date/time separation character: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:355
		// _ = "end of CoverTab[122720]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:356
		_go_fuzz_dep_.CoverTab[122721]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:356
		// _ = "end of CoverTab[122721]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:356
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:356
	// _ = "end of CoverTab[122700]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:356
	_go_fuzz_dep_.CoverTab[122701]++

											if r == ' ' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:358
		_go_fuzz_dep_.CoverTab[122722]++
												lookAhead := l.peekString(3)[1:]
												if len(lookAhead) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:360
			_go_fuzz_dep_.CoverTab[122724]++
													return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:361
			// _ = "end of CoverTab[122724]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:362
			_go_fuzz_dep_.CoverTab[122725]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:362
			// _ = "end of CoverTab[122725]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:362
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:362
		// _ = "end of CoverTab[122722]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:362
		_go_fuzz_dep_.CoverTab[122723]++
												for _, r := range lookAhead {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:363
			_go_fuzz_dep_.CoverTab[122726]++
													if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:364
				_go_fuzz_dep_.CoverTab[122727]++
														return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:365
				// _ = "end of CoverTab[122727]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:366
				_go_fuzz_dep_.CoverTab[122728]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:366
				// _ = "end of CoverTab[122728]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:366
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:366
			// _ = "end of CoverTab[122726]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:367
		// _ = "end of CoverTab[122723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:368
		_go_fuzz_dep_.CoverTab[122729]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:368
		// _ = "end of CoverTab[122729]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:368
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:368
	// _ = "end of CoverTab[122701]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:368
	_go_fuzz_dep_.CoverTab[122702]++

											l.skip()

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:374
	for i := 0; i < 2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:374
		_go_fuzz_dep_.CoverTab[122730]++
												r := l.next()
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:376
			_go_fuzz_dep_.CoverTab[122731]++
													return l.errorf("invalid hour digit in time: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:377
			// _ = "end of CoverTab[122731]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:378
			_go_fuzz_dep_.CoverTab[122732]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:378
			// _ = "end of CoverTab[122732]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:378
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:378
		// _ = "end of CoverTab[122730]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:379
	// _ = "end of CoverTab[122702]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:379
	_go_fuzz_dep_.CoverTab[122703]++

											r = l.next()
											if r != ':' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:382
		_go_fuzz_dep_.CoverTab[122733]++
												return l.errorf("time hour/minute separator should be :, not %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:383
		// _ = "end of CoverTab[122733]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:384
		_go_fuzz_dep_.CoverTab[122734]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:384
		// _ = "end of CoverTab[122734]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:384
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:384
	// _ = "end of CoverTab[122703]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:384
	_go_fuzz_dep_.CoverTab[122704]++

											for i := 0; i < 2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:386
		_go_fuzz_dep_.CoverTab[122735]++
												r := l.next()
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:388
			_go_fuzz_dep_.CoverTab[122736]++
													return l.errorf("invalid minute digit in time: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:389
			// _ = "end of CoverTab[122736]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:390
			_go_fuzz_dep_.CoverTab[122737]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:390
			// _ = "end of CoverTab[122737]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:390
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:390
		// _ = "end of CoverTab[122735]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:391
	// _ = "end of CoverTab[122704]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:391
	_go_fuzz_dep_.CoverTab[122705]++

											r = l.next()
											if r != ':' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:394
		_go_fuzz_dep_.CoverTab[122738]++
												return l.errorf("time minute/second separator should be :, not %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:395
		// _ = "end of CoverTab[122738]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:396
		_go_fuzz_dep_.CoverTab[122739]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:396
		// _ = "end of CoverTab[122739]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:396
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:396
	// _ = "end of CoverTab[122705]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:396
	_go_fuzz_dep_.CoverTab[122706]++

											for i := 0; i < 2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:398
		_go_fuzz_dep_.CoverTab[122740]++
												r := l.next()
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:400
			_go_fuzz_dep_.CoverTab[122741]++
													return l.errorf("invalid second digit in time: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:401
			// _ = "end of CoverTab[122741]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:402
			_go_fuzz_dep_.CoverTab[122742]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:402
			// _ = "end of CoverTab[122742]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:402
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:402
		// _ = "end of CoverTab[122740]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:403
	// _ = "end of CoverTab[122706]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:403
	_go_fuzz_dep_.CoverTab[122707]++

											r = l.peek()
											if r == '.' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:406
		_go_fuzz_dep_.CoverTab[122743]++
												l.next()
												r := l.next()
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:409
			_go_fuzz_dep_.CoverTab[122745]++
													return l.errorf("expected at least one digit in time's fraction, not %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:410
			// _ = "end of CoverTab[122745]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:411
			_go_fuzz_dep_.CoverTab[122746]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:411
			// _ = "end of CoverTab[122746]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:411
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:411
		// _ = "end of CoverTab[122743]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:411
		_go_fuzz_dep_.CoverTab[122744]++

												for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:413
			_go_fuzz_dep_.CoverTab[122747]++
													r := l.peek()
													if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:415
				_go_fuzz_dep_.CoverTab[122749]++
														break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:416
				// _ = "end of CoverTab[122749]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:417
				_go_fuzz_dep_.CoverTab[122750]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:417
				// _ = "end of CoverTab[122750]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:417
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:417
			// _ = "end of CoverTab[122747]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:417
			_go_fuzz_dep_.CoverTab[122748]++
													l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:418
			// _ = "end of CoverTab[122748]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:419
		// _ = "end of CoverTab[122744]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:420
		_go_fuzz_dep_.CoverTab[122751]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:420
		// _ = "end of CoverTab[122751]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:420
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:420
	// _ = "end of CoverTab[122707]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:420
	_go_fuzz_dep_.CoverTab[122708]++

											l.emit(tokenLocalTime)

											return l.lexTimeOffset
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:424
	// _ = "end of CoverTab[122708]"

}

func (l *tomlLexer) lexTimeOffset() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:428
	_go_fuzz_dep_.CoverTab[122752]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:436
	r := l.peek()

	if r == 'Z' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:438
		_go_fuzz_dep_.CoverTab[122754]++
												l.next()
												l.emit(tokenTimeOffset)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:440
		// _ = "end of CoverTab[122754]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:441
		_go_fuzz_dep_.CoverTab[122755]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:441
		if r == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:441
			_go_fuzz_dep_.CoverTab[122756]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:441
			return r == '-'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:441
			// _ = "end of CoverTab[122756]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:441
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:441
			_go_fuzz_dep_.CoverTab[122757]++
													l.next()

													for i := 0; i < 2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:444
				_go_fuzz_dep_.CoverTab[122761]++
														r := l.next()
														if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:446
					_go_fuzz_dep_.CoverTab[122762]++
															return l.errorf("invalid hour digit in time offset: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:447
					// _ = "end of CoverTab[122762]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:448
					_go_fuzz_dep_.CoverTab[122763]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:448
					// _ = "end of CoverTab[122763]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:448
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:448
				// _ = "end of CoverTab[122761]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:449
			// _ = "end of CoverTab[122757]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:449
			_go_fuzz_dep_.CoverTab[122758]++

													r = l.next()
													if r != ':' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:452
				_go_fuzz_dep_.CoverTab[122764]++
														return l.errorf("time offset hour/minute separator should be :, not %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:453
				// _ = "end of CoverTab[122764]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:454
				_go_fuzz_dep_.CoverTab[122765]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:454
				// _ = "end of CoverTab[122765]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:454
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:454
			// _ = "end of CoverTab[122758]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:454
			_go_fuzz_dep_.CoverTab[122759]++

													for i := 0; i < 2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:456
				_go_fuzz_dep_.CoverTab[122766]++
														r := l.next()
														if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:458
					_go_fuzz_dep_.CoverTab[122767]++
															return l.errorf("invalid minute digit in time offset: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:459
					// _ = "end of CoverTab[122767]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:460
					_go_fuzz_dep_.CoverTab[122768]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:460
					// _ = "end of CoverTab[122768]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:460
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:460
				// _ = "end of CoverTab[122766]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:461
			// _ = "end of CoverTab[122759]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:461
			_go_fuzz_dep_.CoverTab[122760]++

													l.emit(tokenTimeOffset)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:463
			// _ = "end of CoverTab[122760]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:464
			_go_fuzz_dep_.CoverTab[122769]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:464
			// _ = "end of CoverTab[122769]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:464
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:464
		// _ = "end of CoverTab[122755]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:464
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:464
	// _ = "end of CoverTab[122752]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:464
	_go_fuzz_dep_.CoverTab[122753]++

											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:466
	// _ = "end of CoverTab[122753]"
}

func (l *tomlLexer) lexTime() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:469
	_go_fuzz_dep_.CoverTab[122770]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:474
	for i := 0; i < 2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:474
		_go_fuzz_dep_.CoverTab[122775]++
												r := l.next()
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:476
			_go_fuzz_dep_.CoverTab[122776]++
													return l.errorf("invalid minute digit in time: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:477
			// _ = "end of CoverTab[122776]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:478
			_go_fuzz_dep_.CoverTab[122777]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:478
			// _ = "end of CoverTab[122777]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:478
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:478
		// _ = "end of CoverTab[122775]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:479
	// _ = "end of CoverTab[122770]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:479
	_go_fuzz_dep_.CoverTab[122771]++

											r := l.next()
											if r != ':' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:482
		_go_fuzz_dep_.CoverTab[122778]++
												return l.errorf("time minute/second separator should be :, not %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:483
		// _ = "end of CoverTab[122778]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:484
		_go_fuzz_dep_.CoverTab[122779]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:484
		// _ = "end of CoverTab[122779]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:484
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:484
	// _ = "end of CoverTab[122771]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:484
	_go_fuzz_dep_.CoverTab[122772]++

											for i := 0; i < 2; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:486
		_go_fuzz_dep_.CoverTab[122780]++
												r := l.next()
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:488
			_go_fuzz_dep_.CoverTab[122781]++
													return l.errorf("invalid second digit in time: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:489
			// _ = "end of CoverTab[122781]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:490
			_go_fuzz_dep_.CoverTab[122782]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:490
			// _ = "end of CoverTab[122782]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:490
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:490
		// _ = "end of CoverTab[122780]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:491
	// _ = "end of CoverTab[122772]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:491
	_go_fuzz_dep_.CoverTab[122773]++

											r = l.peek()
											if r == '.' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:494
		_go_fuzz_dep_.CoverTab[122783]++
												l.next()
												r := l.next()
												if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:497
			_go_fuzz_dep_.CoverTab[122785]++
													return l.errorf("expected at least one digit in time's fraction, not %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:498
			// _ = "end of CoverTab[122785]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:499
			_go_fuzz_dep_.CoverTab[122786]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:499
			// _ = "end of CoverTab[122786]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:499
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:499
		// _ = "end of CoverTab[122783]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:499
		_go_fuzz_dep_.CoverTab[122784]++

												for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:501
			_go_fuzz_dep_.CoverTab[122787]++
													r := l.peek()
													if !isDigit(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:503
				_go_fuzz_dep_.CoverTab[122789]++
														break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:504
				// _ = "end of CoverTab[122789]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:505
				_go_fuzz_dep_.CoverTab[122790]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:505
				// _ = "end of CoverTab[122790]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:505
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:505
			// _ = "end of CoverTab[122787]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:505
			_go_fuzz_dep_.CoverTab[122788]++
													l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:506
			// _ = "end of CoverTab[122788]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:507
		// _ = "end of CoverTab[122784]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:508
		_go_fuzz_dep_.CoverTab[122791]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:508
		// _ = "end of CoverTab[122791]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:508
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:508
	// _ = "end of CoverTab[122773]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:508
	_go_fuzz_dep_.CoverTab[122774]++

											l.emit(tokenLocalTime)
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:511
	// _ = "end of CoverTab[122774]"

}

func (l *tomlLexer) lexTrue() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:515
	_go_fuzz_dep_.CoverTab[122792]++
											l.fastForward(4)
											l.emit(tokenTrue)
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:518
	// _ = "end of CoverTab[122792]"
}

func (l *tomlLexer) lexFalse() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:521
	_go_fuzz_dep_.CoverTab[122793]++
											l.fastForward(5)
											l.emit(tokenFalse)
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:524
	// _ = "end of CoverTab[122793]"
}

func (l *tomlLexer) lexInf() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:527
	_go_fuzz_dep_.CoverTab[122794]++
											l.fastForward(3)
											l.emit(tokenInf)
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:530
	// _ = "end of CoverTab[122794]"
}

func (l *tomlLexer) lexNan() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:533
	_go_fuzz_dep_.CoverTab[122795]++
											l.fastForward(3)
											l.emit(tokenNan)
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:536
	// _ = "end of CoverTab[122795]"
}

func (l *tomlLexer) lexEqual() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:539
	_go_fuzz_dep_.CoverTab[122796]++
											l.next()
											l.emit(tokenEqual)
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:542
	// _ = "end of CoverTab[122796]"
}

func (l *tomlLexer) lexComma() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:545
	_go_fuzz_dep_.CoverTab[122797]++
											l.next()
											l.emit(tokenComma)
											if len(l.brackets) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:548
		_go_fuzz_dep_.CoverTab[122799]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:548
		return l.brackets[len(l.brackets)-1] == '{'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:548
		// _ = "end of CoverTab[122799]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:548
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:548
		_go_fuzz_dep_.CoverTab[122800]++
												return l.lexVoid
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:549
		// _ = "end of CoverTab[122800]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:550
		_go_fuzz_dep_.CoverTab[122801]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:550
		// _ = "end of CoverTab[122801]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:550
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:550
	// _ = "end of CoverTab[122797]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:550
	_go_fuzz_dep_.CoverTab[122798]++
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:551
	// _ = "end of CoverTab[122798]"
}

// Parse the key and emits its value without escape sequences.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:554
// bare keys, basic string keys and literal string keys are supported.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:556
func (l *tomlLexer) lexKey() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:556
	_go_fuzz_dep_.CoverTab[122802]++
											var sb strings.Builder

											for r := l.peek(); isKeyChar(r) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:559
		_go_fuzz_dep_.CoverTab[122804]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:559
		return r == '\n'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:559
		// _ = "end of CoverTab[122804]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:559
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:559
		_go_fuzz_dep_.CoverTab[122805]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:559
		return r == '\r'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:559
		// _ = "end of CoverTab[122805]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:559
	}(); r = l.peek() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:559
		_go_fuzz_dep_.CoverTab[122806]++
												if r == '"' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:560
			_go_fuzz_dep_.CoverTab[122808]++
													l.next()
													str, err := l.lexStringAsString(`"`, false, true)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:563
				_go_fuzz_dep_.CoverTab[122810]++
														return l.errorf(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:564
				// _ = "end of CoverTab[122810]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:565
				_go_fuzz_dep_.CoverTab[122811]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:565
				// _ = "end of CoverTab[122811]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:565
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:565
			// _ = "end of CoverTab[122808]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:565
			_go_fuzz_dep_.CoverTab[122809]++
													sb.WriteString("\"")
													sb.WriteString(str)
													sb.WriteString("\"")
													l.next()
													continue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:570
			// _ = "end of CoverTab[122809]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:571
			_go_fuzz_dep_.CoverTab[122812]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:571
			if r == '\'' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:571
				_go_fuzz_dep_.CoverTab[122813]++
														l.next()
														str, err := l.lexLiteralStringAsString(`'`, false)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:574
					_go_fuzz_dep_.CoverTab[122815]++
															return l.errorf(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:575
					// _ = "end of CoverTab[122815]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:576
					_go_fuzz_dep_.CoverTab[122816]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:576
					// _ = "end of CoverTab[122816]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:576
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:576
				// _ = "end of CoverTab[122813]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:576
				_go_fuzz_dep_.CoverTab[122814]++
														sb.WriteString("'")
														sb.WriteString(str)
														sb.WriteString("'")
														l.next()
														continue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:581
				// _ = "end of CoverTab[122814]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:582
				_go_fuzz_dep_.CoverTab[122817]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:582
				if r == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:582
					_go_fuzz_dep_.CoverTab[122818]++
															return l.errorf("keys cannot contain new lines")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:583
					// _ = "end of CoverTab[122818]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:584
					_go_fuzz_dep_.CoverTab[122819]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:584
					if isSpace(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:584
						_go_fuzz_dep_.CoverTab[122820]++
																var str strings.Builder
																str.WriteString(" ")

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:589
						l.next()
						for r = l.peek(); isSpace(r); r = l.peek() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:590
							_go_fuzz_dep_.CoverTab[122824]++
																	str.WriteRune(r)
																	l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:592
							// _ = "end of CoverTab[122824]"
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:593
						// _ = "end of CoverTab[122820]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:593
						_go_fuzz_dep_.CoverTab[122821]++

																if r != '.' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:595
							_go_fuzz_dep_.CoverTab[122825]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:596
							// _ = "end of CoverTab[122825]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:597
							_go_fuzz_dep_.CoverTab[122826]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:597
							// _ = "end of CoverTab[122826]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:597
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:597
						// _ = "end of CoverTab[122821]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:597
						_go_fuzz_dep_.CoverTab[122822]++
																str.WriteString(".")

																l.next()
																for r = l.peek(); isSpace(r); r = l.peek() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:601
							_go_fuzz_dep_.CoverTab[122827]++
																	str.WriteRune(r)
																	l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:603
							// _ = "end of CoverTab[122827]"
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:604
						// _ = "end of CoverTab[122822]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:604
						_go_fuzz_dep_.CoverTab[122823]++
																sb.WriteString(str.String())
																continue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:606
						// _ = "end of CoverTab[122823]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:607
						_go_fuzz_dep_.CoverTab[122828]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:607
						if r == '.' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:607
							_go_fuzz_dep_.CoverTab[122829]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:607
							// _ = "end of CoverTab[122829]"

						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:609
							_go_fuzz_dep_.CoverTab[122830]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:609
							if !isValidBareChar(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:609
								_go_fuzz_dep_.CoverTab[122831]++
																		return l.errorf("keys cannot contain %c character", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:610
								// _ = "end of CoverTab[122831]"
							} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
								_go_fuzz_dep_.CoverTab[122832]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
								// _ = "end of CoverTab[122832]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
							}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
							// _ = "end of CoverTab[122830]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
						// _ = "end of CoverTab[122828]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
					// _ = "end of CoverTab[122819]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
				// _ = "end of CoverTab[122817]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
			// _ = "end of CoverTab[122812]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
		// _ = "end of CoverTab[122806]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:611
		_go_fuzz_dep_.CoverTab[122807]++
												sb.WriteRune(r)
												l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:613
		// _ = "end of CoverTab[122807]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:614
	// _ = "end of CoverTab[122802]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:614
	_go_fuzz_dep_.CoverTab[122803]++
											l.emitWithValue(tokenKey, sb.String())
											return l.lexVoid
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:616
	// _ = "end of CoverTab[122803]"
}

func (l *tomlLexer) lexComment(previousState tomlLexStateFn) tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:619
	_go_fuzz_dep_.CoverTab[122833]++
											return func() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:620
		_go_fuzz_dep_.CoverTab[122834]++
												for next := l.peek(); next != '\n' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:621
			_go_fuzz_dep_.CoverTab[122836]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:621
			return next != eof
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:621
			// _ = "end of CoverTab[122836]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:621
		}(); next = l.peek() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:621
			_go_fuzz_dep_.CoverTab[122837]++
													if next == '\r' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:622
				_go_fuzz_dep_.CoverTab[122839]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:622
				return l.follow("\r\n")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:622
				// _ = "end of CoverTab[122839]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:622
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:622
				_go_fuzz_dep_.CoverTab[122840]++
														break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:623
				// _ = "end of CoverTab[122840]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:624
				_go_fuzz_dep_.CoverTab[122841]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:624
				// _ = "end of CoverTab[122841]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:624
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:624
			// _ = "end of CoverTab[122837]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:624
			_go_fuzz_dep_.CoverTab[122838]++
													l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:625
			// _ = "end of CoverTab[122838]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:626
		// _ = "end of CoverTab[122834]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:626
		_go_fuzz_dep_.CoverTab[122835]++
												l.ignore()
												return previousState
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:628
		// _ = "end of CoverTab[122835]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:629
	// _ = "end of CoverTab[122833]"
}

func (l *tomlLexer) lexLeftBracket() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:632
	_go_fuzz_dep_.CoverTab[122842]++
											l.next()
											l.emit(tokenLeftBracket)
											l.brackets = append(l.brackets, '[')
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:636
	// _ = "end of CoverTab[122842]"
}

func (l *tomlLexer) lexLiteralStringAsString(terminator string, discardLeadingNewLine bool) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:639
	_go_fuzz_dep_.CoverTab[122843]++
											var sb strings.Builder

											if discardLeadingNewLine {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:642
		_go_fuzz_dep_.CoverTab[122846]++
												if l.follow("\r\n") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:643
			_go_fuzz_dep_.CoverTab[122847]++
													l.skip()
													l.skip()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:645
			// _ = "end of CoverTab[122847]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:646
			_go_fuzz_dep_.CoverTab[122848]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:646
			if l.peek() == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:646
				_go_fuzz_dep_.CoverTab[122849]++
														l.skip()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:647
				// _ = "end of CoverTab[122849]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:648
				_go_fuzz_dep_.CoverTab[122850]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:648
				// _ = "end of CoverTab[122850]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:648
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:648
			// _ = "end of CoverTab[122848]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:648
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:648
		// _ = "end of CoverTab[122846]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:649
		_go_fuzz_dep_.CoverTab[122851]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:649
		// _ = "end of CoverTab[122851]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:649
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:649
	// _ = "end of CoverTab[122843]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:649
	_go_fuzz_dep_.CoverTab[122844]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:652
	for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:652
		_go_fuzz_dep_.CoverTab[122852]++
												if l.follow(terminator) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:653
			_go_fuzz_dep_.CoverTab[122855]++
													return sb.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:654
			// _ = "end of CoverTab[122855]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:655
			_go_fuzz_dep_.CoverTab[122856]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:655
			// _ = "end of CoverTab[122856]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:655
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:655
		// _ = "end of CoverTab[122852]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:655
		_go_fuzz_dep_.CoverTab[122853]++

												next := l.peek()
												if next == eof {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:658
			_go_fuzz_dep_.CoverTab[122857]++
													break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:659
			// _ = "end of CoverTab[122857]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:660
			_go_fuzz_dep_.CoverTab[122858]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:660
			// _ = "end of CoverTab[122858]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:660
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:660
		// _ = "end of CoverTab[122853]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:660
		_go_fuzz_dep_.CoverTab[122854]++
												sb.WriteRune(l.next())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:661
		// _ = "end of CoverTab[122854]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:662
	// _ = "end of CoverTab[122844]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:662
	_go_fuzz_dep_.CoverTab[122845]++

											return "", errors.New("unclosed string")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:664
	// _ = "end of CoverTab[122845]"
}

func (l *tomlLexer) lexLiteralString() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:667
	_go_fuzz_dep_.CoverTab[122859]++
											l.skip()

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:671
	terminator := "'"
	discardLeadingNewLine := false
	if l.follow("''") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:673
		_go_fuzz_dep_.CoverTab[122862]++
												l.skip()
												l.skip()
												terminator = "'''"
												discardLeadingNewLine = true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:677
		// _ = "end of CoverTab[122862]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:678
		_go_fuzz_dep_.CoverTab[122863]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:678
		// _ = "end of CoverTab[122863]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:678
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:678
	// _ = "end of CoverTab[122859]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:678
	_go_fuzz_dep_.CoverTab[122860]++

											str, err := l.lexLiteralStringAsString(terminator, discardLeadingNewLine)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:681
		_go_fuzz_dep_.CoverTab[122864]++
												return l.errorf(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:682
		// _ = "end of CoverTab[122864]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:683
		_go_fuzz_dep_.CoverTab[122865]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:683
		// _ = "end of CoverTab[122865]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:683
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:683
	// _ = "end of CoverTab[122860]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:683
	_go_fuzz_dep_.CoverTab[122861]++

											l.emitWithValue(tokenString, str)
											l.fastForward(len(terminator))
											l.ignore()
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:688
	// _ = "end of CoverTab[122861]"
}

// Lex a string and return the results as a string.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:691
// Terminator is the substring indicating the end of the token.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:691
// The resulting string does not include the terminator.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:694
func (l *tomlLexer) lexStringAsString(terminator string, discardLeadingNewLine, acceptNewLines bool) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:694
	_go_fuzz_dep_.CoverTab[122866]++
											var sb strings.Builder

											if discardLeadingNewLine {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:697
		_go_fuzz_dep_.CoverTab[122869]++
												if l.follow("\r\n") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:698
			_go_fuzz_dep_.CoverTab[122870]++
													l.skip()
													l.skip()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:700
			// _ = "end of CoverTab[122870]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:701
			_go_fuzz_dep_.CoverTab[122871]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:701
			if l.peek() == '\n' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:701
				_go_fuzz_dep_.CoverTab[122872]++
														l.skip()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:702
				// _ = "end of CoverTab[122872]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:703
				_go_fuzz_dep_.CoverTab[122873]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:703
				// _ = "end of CoverTab[122873]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:703
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:703
			// _ = "end of CoverTab[122871]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:703
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:703
		// _ = "end of CoverTab[122869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:704
		_go_fuzz_dep_.CoverTab[122874]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:704
		// _ = "end of CoverTab[122874]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:704
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:704
	// _ = "end of CoverTab[122866]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:704
	_go_fuzz_dep_.CoverTab[122867]++

											for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:706
		_go_fuzz_dep_.CoverTab[122875]++
												if l.follow(terminator) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:707
			_go_fuzz_dep_.CoverTab[122878]++
													return sb.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:708
			// _ = "end of CoverTab[122878]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:709
			_go_fuzz_dep_.CoverTab[122879]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:709
			// _ = "end of CoverTab[122879]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:709
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:709
		// _ = "end of CoverTab[122875]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:709
		_go_fuzz_dep_.CoverTab[122876]++

												if l.follow("\\") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:711
			_go_fuzz_dep_.CoverTab[122880]++
													l.next()
													switch l.peek() {
			case '\r':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:714
				_go_fuzz_dep_.CoverTab[122881]++
														fallthrough
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:715
				// _ = "end of CoverTab[122881]"
			case '\n':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:716
				_go_fuzz_dep_.CoverTab[122882]++
														fallthrough
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:717
				// _ = "end of CoverTab[122882]"
			case '\t':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:718
				_go_fuzz_dep_.CoverTab[122883]++
														fallthrough
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:719
				// _ = "end of CoverTab[122883]"
			case ' ':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:720
				_go_fuzz_dep_.CoverTab[122884]++

														for strings.ContainsRune("\r\n\t ", l.peek()) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:722
					_go_fuzz_dep_.CoverTab[122900]++
															l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:723
					// _ = "end of CoverTab[122900]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:724
				// _ = "end of CoverTab[122884]"
			case '"':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:725
				_go_fuzz_dep_.CoverTab[122885]++
														sb.WriteString("\"")
														l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:727
				// _ = "end of CoverTab[122885]"
			case 'n':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:728
				_go_fuzz_dep_.CoverTab[122886]++
														sb.WriteString("\n")
														l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:730
				// _ = "end of CoverTab[122886]"
			case 'b':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:731
				_go_fuzz_dep_.CoverTab[122887]++
														sb.WriteString("\b")
														l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:733
				// _ = "end of CoverTab[122887]"
			case 'f':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:734
				_go_fuzz_dep_.CoverTab[122888]++
														sb.WriteString("\f")
														l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:736
				// _ = "end of CoverTab[122888]"
			case '/':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:737
				_go_fuzz_dep_.CoverTab[122889]++
														sb.WriteString("/")
														l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:739
				// _ = "end of CoverTab[122889]"
			case 't':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:740
				_go_fuzz_dep_.CoverTab[122890]++
														sb.WriteString("\t")
														l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:742
				// _ = "end of CoverTab[122890]"
			case 'r':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:743
				_go_fuzz_dep_.CoverTab[122891]++
														sb.WriteString("\r")
														l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:745
				// _ = "end of CoverTab[122891]"
			case '\\':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:746
				_go_fuzz_dep_.CoverTab[122892]++
														sb.WriteString("\\")
														l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:748
				// _ = "end of CoverTab[122892]"
			case 'u':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:749
				_go_fuzz_dep_.CoverTab[122893]++
														l.next()
														var code strings.Builder
														for i := 0; i < 4; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:752
					_go_fuzz_dep_.CoverTab[122901]++
															c := l.peek()
															if !isHexDigit(c) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:754
						_go_fuzz_dep_.CoverTab[122903]++
																return "", errors.New("unfinished unicode escape")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:755
						// _ = "end of CoverTab[122903]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:756
						_go_fuzz_dep_.CoverTab[122904]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:756
						// _ = "end of CoverTab[122904]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:756
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:756
					// _ = "end of CoverTab[122901]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:756
					_go_fuzz_dep_.CoverTab[122902]++
															l.next()
															code.WriteRune(c)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:758
					// _ = "end of CoverTab[122902]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:759
				// _ = "end of CoverTab[122893]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:759
				_go_fuzz_dep_.CoverTab[122894]++
														intcode, err := strconv.ParseInt(code.String(), 16, 32)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:761
					_go_fuzz_dep_.CoverTab[122905]++
															return "", errors.New("invalid unicode escape: \\u" + code.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:762
					// _ = "end of CoverTab[122905]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:763
					_go_fuzz_dep_.CoverTab[122906]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:763
					// _ = "end of CoverTab[122906]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:763
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:763
				// _ = "end of CoverTab[122894]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:763
				_go_fuzz_dep_.CoverTab[122895]++
														sb.WriteRune(rune(intcode))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:764
				// _ = "end of CoverTab[122895]"
			case 'U':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:765
				_go_fuzz_dep_.CoverTab[122896]++
														l.next()
														var code strings.Builder
														for i := 0; i < 8; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:768
					_go_fuzz_dep_.CoverTab[122907]++
															c := l.peek()
															if !isHexDigit(c) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:770
						_go_fuzz_dep_.CoverTab[122909]++
																return "", errors.New("unfinished unicode escape")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:771
						// _ = "end of CoverTab[122909]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:772
						_go_fuzz_dep_.CoverTab[122910]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:772
						// _ = "end of CoverTab[122910]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:772
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:772
					// _ = "end of CoverTab[122907]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:772
					_go_fuzz_dep_.CoverTab[122908]++
															l.next()
															code.WriteRune(c)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:774
					// _ = "end of CoverTab[122908]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:775
				// _ = "end of CoverTab[122896]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:775
				_go_fuzz_dep_.CoverTab[122897]++
														intcode, err := strconv.ParseInt(code.String(), 16, 64)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:777
					_go_fuzz_dep_.CoverTab[122911]++
															return "", errors.New("invalid unicode escape: \\U" + code.String())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:778
					// _ = "end of CoverTab[122911]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:779
					_go_fuzz_dep_.CoverTab[122912]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:779
					// _ = "end of CoverTab[122912]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:779
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:779
				// _ = "end of CoverTab[122897]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:779
				_go_fuzz_dep_.CoverTab[122898]++
														sb.WriteRune(rune(intcode))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:780
				// _ = "end of CoverTab[122898]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:781
				_go_fuzz_dep_.CoverTab[122899]++
														return "", errors.New("invalid escape sequence: \\" + string(l.peek()))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:782
				// _ = "end of CoverTab[122899]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:783
			// _ = "end of CoverTab[122880]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:784
			_go_fuzz_dep_.CoverTab[122913]++
													r := l.peek()

													if 0x00 <= r && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				_go_fuzz_dep_.CoverTab[122915]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				return r <= 0x1F
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				// _ = "end of CoverTab[122915]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				_go_fuzz_dep_.CoverTab[122916]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				return r != '\t'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				// _ = "end of CoverTab[122916]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				_go_fuzz_dep_.CoverTab[122917]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				return !(acceptNewLines && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
					_go_fuzz_dep_.CoverTab[122918]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
					return (r == '\n' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
						_go_fuzz_dep_.CoverTab[122919]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
						return r == '\r'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
						// _ = "end of CoverTab[122919]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
					}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
					// _ = "end of CoverTab[122918]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				}())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				// _ = "end of CoverTab[122917]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:787
				_go_fuzz_dep_.CoverTab[122920]++
														return "", fmt.Errorf("unescaped control character %U", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:788
				// _ = "end of CoverTab[122920]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:789
				_go_fuzz_dep_.CoverTab[122921]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:789
				// _ = "end of CoverTab[122921]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:789
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:789
			// _ = "end of CoverTab[122913]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:789
			_go_fuzz_dep_.CoverTab[122914]++
													l.next()
													sb.WriteRune(r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:791
			// _ = "end of CoverTab[122914]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:792
		// _ = "end of CoverTab[122876]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:792
		_go_fuzz_dep_.CoverTab[122877]++

												if l.peek() == eof {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:794
			_go_fuzz_dep_.CoverTab[122922]++
													break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:795
			// _ = "end of CoverTab[122922]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:796
			_go_fuzz_dep_.CoverTab[122923]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:796
			// _ = "end of CoverTab[122923]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:796
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:796
		// _ = "end of CoverTab[122877]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:797
	// _ = "end of CoverTab[122867]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:797
	_go_fuzz_dep_.CoverTab[122868]++

											return "", errors.New("unclosed string")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:799
	// _ = "end of CoverTab[122868]"
}

func (l *tomlLexer) lexString() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:802
	_go_fuzz_dep_.CoverTab[122924]++
											l.skip()

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:806
	terminator := `"`
	discardLeadingNewLine := false
	acceptNewLines := false
	if l.follow(`""`) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:809
		_go_fuzz_dep_.CoverTab[122927]++
												l.skip()
												l.skip()
												terminator = `"""`
												discardLeadingNewLine = true
												acceptNewLines = true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:814
		// _ = "end of CoverTab[122927]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:815
		_go_fuzz_dep_.CoverTab[122928]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:815
		// _ = "end of CoverTab[122928]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:815
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:815
	// _ = "end of CoverTab[122924]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:815
	_go_fuzz_dep_.CoverTab[122925]++

											str, err := l.lexStringAsString(terminator, discardLeadingNewLine, acceptNewLines)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:818
		_go_fuzz_dep_.CoverTab[122929]++
												return l.errorf(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:819
		// _ = "end of CoverTab[122929]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:820
		_go_fuzz_dep_.CoverTab[122930]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:820
		// _ = "end of CoverTab[122930]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:820
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:820
	// _ = "end of CoverTab[122925]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:820
	_go_fuzz_dep_.CoverTab[122926]++

											l.emitWithValue(tokenString, str)
											l.fastForward(len(terminator))
											l.ignore()
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:825
	// _ = "end of CoverTab[122926]"
}

func (l *tomlLexer) lexTableKey() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:828
	_go_fuzz_dep_.CoverTab[122931]++
											l.next()

											if l.peek() == '[' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:831
		_go_fuzz_dep_.CoverTab[122933]++

												l.next()
												l.emit(tokenDoubleLeftBracket)
												return l.lexInsideTableArrayKey
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:835
		// _ = "end of CoverTab[122933]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:836
		_go_fuzz_dep_.CoverTab[122934]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:836
		// _ = "end of CoverTab[122934]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:836
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:836
	// _ = "end of CoverTab[122931]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:836
	_go_fuzz_dep_.CoverTab[122932]++

											l.emit(tokenLeftBracket)
											return l.lexInsideTableKey
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:839
	// _ = "end of CoverTab[122932]"
}

// Parse the key till "]]", but only bare keys are supported
func (l *tomlLexer) lexInsideTableArrayKey() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:843
	_go_fuzz_dep_.CoverTab[122935]++
											for r := l.peek(); r != eof; r = l.peek() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:844
		_go_fuzz_dep_.CoverTab[122937]++
												switch r {
		case ']':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:846
			_go_fuzz_dep_.CoverTab[122938]++
													if l.currentTokenStop > l.currentTokenStart {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:847
				_go_fuzz_dep_.CoverTab[122943]++
														l.emit(tokenKeyGroupArray)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:848
				// _ = "end of CoverTab[122943]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:849
				_go_fuzz_dep_.CoverTab[122944]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:849
				// _ = "end of CoverTab[122944]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:849
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:849
			// _ = "end of CoverTab[122938]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:849
			_go_fuzz_dep_.CoverTab[122939]++
													l.next()
													if l.peek() != ']' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:851
				_go_fuzz_dep_.CoverTab[122945]++
														break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:852
				// _ = "end of CoverTab[122945]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:853
				_go_fuzz_dep_.CoverTab[122946]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:853
				// _ = "end of CoverTab[122946]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:853
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:853
			// _ = "end of CoverTab[122939]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:853
			_go_fuzz_dep_.CoverTab[122940]++
													l.next()
													l.emit(tokenDoubleRightBracket)
													return l.lexVoid
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:856
			// _ = "end of CoverTab[122940]"
		case '[':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:857
			_go_fuzz_dep_.CoverTab[122941]++
													return l.errorf("table array key cannot contain ']'")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:858
			// _ = "end of CoverTab[122941]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:859
			_go_fuzz_dep_.CoverTab[122942]++
													l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:860
			// _ = "end of CoverTab[122942]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:861
		// _ = "end of CoverTab[122937]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:862
	// _ = "end of CoverTab[122935]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:862
	_go_fuzz_dep_.CoverTab[122936]++
											return l.errorf("unclosed table array key")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:863
	// _ = "end of CoverTab[122936]"
}

// Parse the key till "]" but only bare keys are supported
func (l *tomlLexer) lexInsideTableKey() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:867
	_go_fuzz_dep_.CoverTab[122947]++
											for r := l.peek(); r != eof; r = l.peek() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:868
		_go_fuzz_dep_.CoverTab[122949]++
												switch r {
		case ']':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:870
			_go_fuzz_dep_.CoverTab[122950]++
													if l.currentTokenStop > l.currentTokenStart {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:871
				_go_fuzz_dep_.CoverTab[122954]++
														l.emit(tokenKeyGroup)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:872
				// _ = "end of CoverTab[122954]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:873
				_go_fuzz_dep_.CoverTab[122955]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:873
				// _ = "end of CoverTab[122955]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:873
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:873
			// _ = "end of CoverTab[122950]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:873
			_go_fuzz_dep_.CoverTab[122951]++
													l.next()
													l.emit(tokenRightBracket)
													return l.lexVoid
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:876
			// _ = "end of CoverTab[122951]"
		case '[':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:877
			_go_fuzz_dep_.CoverTab[122952]++
													return l.errorf("table key cannot contain ']'")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:878
			// _ = "end of CoverTab[122952]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:879
			_go_fuzz_dep_.CoverTab[122953]++
													l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:880
			// _ = "end of CoverTab[122953]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:881
		// _ = "end of CoverTab[122949]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:882
	// _ = "end of CoverTab[122947]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:882
	_go_fuzz_dep_.CoverTab[122948]++
											return l.errorf("unclosed table key")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:883
	// _ = "end of CoverTab[122948]"
}

func (l *tomlLexer) lexRightBracket() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:886
	_go_fuzz_dep_.CoverTab[122956]++
											l.next()
											l.emit(tokenRightBracket)
											if len(l.brackets) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:889
		_go_fuzz_dep_.CoverTab[122958]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:889
		return l.brackets[len(l.brackets)-1] != '['
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:889
		// _ = "end of CoverTab[122958]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:889
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:889
		_go_fuzz_dep_.CoverTab[122959]++
												return l.errorf("cannot have ']' here")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:890
		// _ = "end of CoverTab[122959]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:891
		_go_fuzz_dep_.CoverTab[122960]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:891
		// _ = "end of CoverTab[122960]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:891
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:891
	// _ = "end of CoverTab[122956]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:891
	_go_fuzz_dep_.CoverTab[122957]++
											l.brackets = l.brackets[:len(l.brackets)-1]
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:893
	// _ = "end of CoverTab[122957]"
}

type validRuneFn func(r rune) bool

func isValidHexRune(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:898
	_go_fuzz_dep_.CoverTab[122961]++
											return r >= 'a' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:899
		_go_fuzz_dep_.CoverTab[122962]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:899
		return r <= 'f'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:899
		// _ = "end of CoverTab[122962]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:899
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:899
		_go_fuzz_dep_.CoverTab[122963]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:899
		return r >= 'A' && func() bool {
													_go_fuzz_dep_.CoverTab[122964]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:900
			return r <= 'F'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:900
			// _ = "end of CoverTab[122964]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:900
		}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:900
		// _ = "end of CoverTab[122963]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:900
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:900
		_go_fuzz_dep_.CoverTab[122965]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:900
		return r >= '0' && func() bool {
													_go_fuzz_dep_.CoverTab[122966]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:901
			return r <= '9'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:901
			// _ = "end of CoverTab[122966]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:901
		}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:901
		// _ = "end of CoverTab[122965]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:901
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:901
		_go_fuzz_dep_.CoverTab[122967]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:901
		return r == '_'
												// _ = "end of CoverTab[122967]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:902
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:902
	// _ = "end of CoverTab[122961]"
}

func isValidOctalRune(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:905
	_go_fuzz_dep_.CoverTab[122968]++
											return r >= '0' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:906
		_go_fuzz_dep_.CoverTab[122969]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:906
		return r <= '7'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:906
		// _ = "end of CoverTab[122969]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:906
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:906
		_go_fuzz_dep_.CoverTab[122970]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:906
		return r == '_'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:906
		// _ = "end of CoverTab[122970]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:906
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:906
	// _ = "end of CoverTab[122968]"
}

func isValidBinaryRune(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:909
	_go_fuzz_dep_.CoverTab[122971]++
											return r == '0' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:910
		_go_fuzz_dep_.CoverTab[122972]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:910
		return r == '1'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:910
		// _ = "end of CoverTab[122972]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:910
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:910
		_go_fuzz_dep_.CoverTab[122973]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:910
		return r == '_'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:910
		// _ = "end of CoverTab[122973]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:910
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:910
	// _ = "end of CoverTab[122971]"
}

func (l *tomlLexer) lexNumber() tomlLexStateFn {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:913
	_go_fuzz_dep_.CoverTab[122974]++
											r := l.peek()

											if r == '0' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:916
		_go_fuzz_dep_.CoverTab[122980]++
												follow := l.peekString(2)
												if len(follow) == 2 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:918
			_go_fuzz_dep_.CoverTab[122981]++
													var isValidRune validRuneFn
													switch follow[1] {
			case 'x':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:921
				_go_fuzz_dep_.CoverTab[122983]++
														isValidRune = isValidHexRune
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:922
				// _ = "end of CoverTab[122983]"
			case 'o':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:923
				_go_fuzz_dep_.CoverTab[122984]++
														isValidRune = isValidOctalRune
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:924
				// _ = "end of CoverTab[122984]"
			case 'b':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:925
				_go_fuzz_dep_.CoverTab[122985]++
														isValidRune = isValidBinaryRune
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:926
				// _ = "end of CoverTab[122985]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:927
				_go_fuzz_dep_.CoverTab[122986]++
														if follow[1] >= 'a' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
					_go_fuzz_dep_.CoverTab[122987]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
					return follow[1] <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
					// _ = "end of CoverTab[122987]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
					_go_fuzz_dep_.CoverTab[122988]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
					return follow[1] >= 'A' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
						_go_fuzz_dep_.CoverTab[122989]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
						return follow[1] <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
						// _ = "end of CoverTab[122989]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
					}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
					// _ = "end of CoverTab[122988]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
				}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:928
					_go_fuzz_dep_.CoverTab[122990]++
															return l.errorf("unknown number base: %s. possible options are x (hex) o (octal) b (binary)", string(follow[1]))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:929
					// _ = "end of CoverTab[122990]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:930
					_go_fuzz_dep_.CoverTab[122991]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:930
					// _ = "end of CoverTab[122991]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:930
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:930
				// _ = "end of CoverTab[122986]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:931
			// _ = "end of CoverTab[122981]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:931
			_go_fuzz_dep_.CoverTab[122982]++

													if isValidRune != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:933
				_go_fuzz_dep_.CoverTab[122992]++
														l.next()
														l.next()
														digitSeen := false
														for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:937
					_go_fuzz_dep_.CoverTab[122995]++
															next := l.peek()
															if !isValidRune(next) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:939
						_go_fuzz_dep_.CoverTab[122997]++
																break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:940
						// _ = "end of CoverTab[122997]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:941
						_go_fuzz_dep_.CoverTab[122998]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:941
						// _ = "end of CoverTab[122998]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:941
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:941
					// _ = "end of CoverTab[122995]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:941
					_go_fuzz_dep_.CoverTab[122996]++
															digitSeen = true
															l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:943
					// _ = "end of CoverTab[122996]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:944
				// _ = "end of CoverTab[122992]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:944
				_go_fuzz_dep_.CoverTab[122993]++

														if !digitSeen {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:946
					_go_fuzz_dep_.CoverTab[122999]++
															return l.errorf("number needs at least one digit")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:947
					// _ = "end of CoverTab[122999]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:948
					_go_fuzz_dep_.CoverTab[123000]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:948
					// _ = "end of CoverTab[123000]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:948
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:948
				// _ = "end of CoverTab[122993]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:948
				_go_fuzz_dep_.CoverTab[122994]++

														l.emit(tokenInteger)

														return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:952
				// _ = "end of CoverTab[122994]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:953
				_go_fuzz_dep_.CoverTab[123001]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:953
				// _ = "end of CoverTab[123001]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:953
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:953
			// _ = "end of CoverTab[122982]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:954
			_go_fuzz_dep_.CoverTab[123002]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:954
			// _ = "end of CoverTab[123002]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:954
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:954
		// _ = "end of CoverTab[122980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:955
		_go_fuzz_dep_.CoverTab[123003]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:955
		// _ = "end of CoverTab[123003]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:955
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:955
	// _ = "end of CoverTab[122974]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:955
	_go_fuzz_dep_.CoverTab[122975]++

											if r == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:957
		_go_fuzz_dep_.CoverTab[123004]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:957
		return r == '-'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:957
		// _ = "end of CoverTab[123004]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:957
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:957
		_go_fuzz_dep_.CoverTab[123005]++
												l.next()
												if l.follow("inf") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:959
			_go_fuzz_dep_.CoverTab[123007]++
													return l.lexInf
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:960
			// _ = "end of CoverTab[123007]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:961
			_go_fuzz_dep_.CoverTab[123008]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:961
			// _ = "end of CoverTab[123008]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:961
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:961
		// _ = "end of CoverTab[123005]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:961
		_go_fuzz_dep_.CoverTab[123006]++
												if l.follow("nan") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:962
			_go_fuzz_dep_.CoverTab[123009]++
													return l.lexNan
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:963
			// _ = "end of CoverTab[123009]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:964
			_go_fuzz_dep_.CoverTab[123010]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:964
			// _ = "end of CoverTab[123010]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:964
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:964
		// _ = "end of CoverTab[123006]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:965
		_go_fuzz_dep_.CoverTab[123011]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:965
		// _ = "end of CoverTab[123011]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:965
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:965
	// _ = "end of CoverTab[122975]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:965
	_go_fuzz_dep_.CoverTab[122976]++

											pointSeen := false
											expSeen := false
											digitSeen := false
											for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:970
		_go_fuzz_dep_.CoverTab[123012]++
												next := l.peek()
												if next == '.' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:972
			_go_fuzz_dep_.CoverTab[123014]++
													if pointSeen {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:973
				_go_fuzz_dep_.CoverTab[123017]++
														return l.errorf("cannot have two dots in one float")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:974
				// _ = "end of CoverTab[123017]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:975
				_go_fuzz_dep_.CoverTab[123018]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:975
				// _ = "end of CoverTab[123018]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:975
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:975
			// _ = "end of CoverTab[123014]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:975
			_go_fuzz_dep_.CoverTab[123015]++
													l.next()
													if !isDigit(l.peek()) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:977
				_go_fuzz_dep_.CoverTab[123019]++
														return l.errorf("float cannot end with a dot")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:978
				// _ = "end of CoverTab[123019]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:979
				_go_fuzz_dep_.CoverTab[123020]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:979
				// _ = "end of CoverTab[123020]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:979
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:979
			// _ = "end of CoverTab[123015]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:979
			_go_fuzz_dep_.CoverTab[123016]++
													pointSeen = true
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:980
			// _ = "end of CoverTab[123016]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:981
			_go_fuzz_dep_.CoverTab[123021]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:981
			if next == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:981
				_go_fuzz_dep_.CoverTab[123022]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:981
				return next == 'E'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:981
				// _ = "end of CoverTab[123022]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:981
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:981
				_go_fuzz_dep_.CoverTab[123023]++
														expSeen = true
														l.next()
														r := l.peek()
														if r == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:985
					_go_fuzz_dep_.CoverTab[123024]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:985
					return r == '-'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:985
					// _ = "end of CoverTab[123024]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:985
				}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:985
					_go_fuzz_dep_.CoverTab[123025]++
															l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:986
					// _ = "end of CoverTab[123025]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:987
					_go_fuzz_dep_.CoverTab[123026]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:987
					// _ = "end of CoverTab[123026]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:987
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:987
				// _ = "end of CoverTab[123023]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:988
				_go_fuzz_dep_.CoverTab[123027]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:988
				if isDigit(next) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:988
					_go_fuzz_dep_.CoverTab[123028]++
															digitSeen = true
															l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:990
					// _ = "end of CoverTab[123028]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:991
					_go_fuzz_dep_.CoverTab[123029]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:991
					if next == '_' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:991
						_go_fuzz_dep_.CoverTab[123030]++
																l.next()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:992
						// _ = "end of CoverTab[123030]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:993
						_go_fuzz_dep_.CoverTab[123031]++
																break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:994
						// _ = "end of CoverTab[123031]"
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:995
					// _ = "end of CoverTab[123029]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:995
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:995
				// _ = "end of CoverTab[123027]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:995
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:995
			// _ = "end of CoverTab[123021]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:995
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:995
		// _ = "end of CoverTab[123012]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:995
		_go_fuzz_dep_.CoverTab[123013]++
												if pointSeen && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:996
			_go_fuzz_dep_.CoverTab[123032]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:996
			return !digitSeen
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:996
			// _ = "end of CoverTab[123032]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:996
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:996
			_go_fuzz_dep_.CoverTab[123033]++
													return l.errorf("cannot start float with a dot")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:997
			// _ = "end of CoverTab[123033]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:998
			_go_fuzz_dep_.CoverTab[123034]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:998
			// _ = "end of CoverTab[123034]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:998
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:998
		// _ = "end of CoverTab[123013]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:999
	// _ = "end of CoverTab[122976]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:999
	_go_fuzz_dep_.CoverTab[122977]++

											if !digitSeen {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1001
		_go_fuzz_dep_.CoverTab[123035]++
												return l.errorf("no digit in that number")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1002
		// _ = "end of CoverTab[123035]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1003
		_go_fuzz_dep_.CoverTab[123036]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1003
		// _ = "end of CoverTab[123036]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1003
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1003
	// _ = "end of CoverTab[122977]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1003
	_go_fuzz_dep_.CoverTab[122978]++
											if pointSeen || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1004
		_go_fuzz_dep_.CoverTab[123037]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1004
		return expSeen
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1004
		// _ = "end of CoverTab[123037]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1004
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1004
		_go_fuzz_dep_.CoverTab[123038]++
												l.emit(tokenFloat)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1005
		// _ = "end of CoverTab[123038]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1006
		_go_fuzz_dep_.CoverTab[123039]++
												l.emit(tokenInteger)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1007
		// _ = "end of CoverTab[123039]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1008
	// _ = "end of CoverTab[122978]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1008
	_go_fuzz_dep_.CoverTab[122979]++
											return l.lexRvalue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1009
	// _ = "end of CoverTab[122979]"
}

func (l *tomlLexer) run() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1012
	_go_fuzz_dep_.CoverTab[123040]++
											for state := l.lexVoid; state != nil; {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1013
		_go_fuzz_dep_.CoverTab[123041]++
												state = state()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1014
		// _ = "end of CoverTab[123041]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1015
	// _ = "end of CoverTab[123040]"
}

// Entry point
func lexToml(inputBytes []byte) []token {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1019
	_go_fuzz_dep_.CoverTab[123042]++
											runes := bytes.Runes(inputBytes)
											l := &tomlLexer{
		input:		runes,
		tokens:		make([]token, 0, 256),
		line:		1,
		col:		1,
		endbufferLine:	1,
		endbufferCol:	1,
	}
											l.run()
											return l.tokens
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1030
	// _ = "end of CoverTab[123042]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1031
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/lexer.go:1031
var _ = _go_fuzz_dep_.CoverTab
