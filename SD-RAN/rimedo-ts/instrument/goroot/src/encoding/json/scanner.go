// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/json/scanner.go:5
package json

//line /usr/local/go/src/encoding/json/scanner.go:5
import (
//line /usr/local/go/src/encoding/json/scanner.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/json/scanner.go:5
)
//line /usr/local/go/src/encoding/json/scanner.go:5
import (
//line /usr/local/go/src/encoding/json/scanner.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/json/scanner.go:5
)

//line /usr/local/go/src/encoding/json/scanner.go:16
import (
	"strconv"
	"sync"
)

// Valid reports whether data is a valid JSON encoding.
func Valid(data []byte) bool {
//line /usr/local/go/src/encoding/json/scanner.go:22
	_go_fuzz_dep_.CoverTab[28116]++
							scan := newScanner()
							defer freeScanner(scan)
							return checkValid(data, scan) == nil
//line /usr/local/go/src/encoding/json/scanner.go:25
	// _ = "end of CoverTab[28116]"
}

// checkValid verifies that data is valid JSON-encoded data.
//line /usr/local/go/src/encoding/json/scanner.go:28
// scan is passed in for use by checkValid to avoid an allocation.
//line /usr/local/go/src/encoding/json/scanner.go:28
// checkValid returns nil or a SyntaxError.
//line /usr/local/go/src/encoding/json/scanner.go:31
func checkValid(data []byte, scan *scanner) error {
//line /usr/local/go/src/encoding/json/scanner.go:31
	_go_fuzz_dep_.CoverTab[28117]++
							scan.reset()
							for _, c := range data {
//line /usr/local/go/src/encoding/json/scanner.go:33
		_go_fuzz_dep_.CoverTab[28120]++
								scan.bytes++
								if scan.step(scan, c) == scanError {
//line /usr/local/go/src/encoding/json/scanner.go:35
			_go_fuzz_dep_.CoverTab[28121]++
									return scan.err
//line /usr/local/go/src/encoding/json/scanner.go:36
			// _ = "end of CoverTab[28121]"
		} else {
//line /usr/local/go/src/encoding/json/scanner.go:37
			_go_fuzz_dep_.CoverTab[28122]++
//line /usr/local/go/src/encoding/json/scanner.go:37
			// _ = "end of CoverTab[28122]"
//line /usr/local/go/src/encoding/json/scanner.go:37
		}
//line /usr/local/go/src/encoding/json/scanner.go:37
		// _ = "end of CoverTab[28120]"
	}
//line /usr/local/go/src/encoding/json/scanner.go:38
	// _ = "end of CoverTab[28117]"
//line /usr/local/go/src/encoding/json/scanner.go:38
	_go_fuzz_dep_.CoverTab[28118]++
							if scan.eof() == scanError {
//line /usr/local/go/src/encoding/json/scanner.go:39
		_go_fuzz_dep_.CoverTab[28123]++
								return scan.err
//line /usr/local/go/src/encoding/json/scanner.go:40
		// _ = "end of CoverTab[28123]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:41
		_go_fuzz_dep_.CoverTab[28124]++
//line /usr/local/go/src/encoding/json/scanner.go:41
		// _ = "end of CoverTab[28124]"
//line /usr/local/go/src/encoding/json/scanner.go:41
	}
//line /usr/local/go/src/encoding/json/scanner.go:41
	// _ = "end of CoverTab[28118]"
//line /usr/local/go/src/encoding/json/scanner.go:41
	_go_fuzz_dep_.CoverTab[28119]++
							return nil
//line /usr/local/go/src/encoding/json/scanner.go:42
	// _ = "end of CoverTab[28119]"
}

// A SyntaxError is a description of a JSON syntax error.
//line /usr/local/go/src/encoding/json/scanner.go:45
// Unmarshal will return a SyntaxError if the JSON can't be parsed.
//line /usr/local/go/src/encoding/json/scanner.go:47
type SyntaxError struct {
	msg	string	// description of error
	Offset	int64	// error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string {
//line /usr/local/go/src/encoding/json/scanner.go:52
	_go_fuzz_dep_.CoverTab[28125]++
//line /usr/local/go/src/encoding/json/scanner.go:52
	return e.msg
//line /usr/local/go/src/encoding/json/scanner.go:52
	// _ = "end of CoverTab[28125]"
//line /usr/local/go/src/encoding/json/scanner.go:52
}

// A scanner is a JSON scanning state machine.
//line /usr/local/go/src/encoding/json/scanner.go:54
// Callers call scan.reset and then pass bytes in one at a time
//line /usr/local/go/src/encoding/json/scanner.go:54
// by calling scan.step(&scan, c) for each byte.
//line /usr/local/go/src/encoding/json/scanner.go:54
// The return value, referred to as an opcode, tells the
//line /usr/local/go/src/encoding/json/scanner.go:54
// caller about significant parsing events like beginning
//line /usr/local/go/src/encoding/json/scanner.go:54
// and ending literals, objects, and arrays, so that the
//line /usr/local/go/src/encoding/json/scanner.go:54
// caller can follow along if it wishes.
//line /usr/local/go/src/encoding/json/scanner.go:54
// The return value scanEnd indicates that a single top-level
//line /usr/local/go/src/encoding/json/scanner.go:54
// JSON value has been completed, *before* the byte that
//line /usr/local/go/src/encoding/json/scanner.go:54
// just got passed in.  (The indication must be delayed in order
//line /usr/local/go/src/encoding/json/scanner.go:54
// to recognize the end of numbers: is 123 a whole value or
//line /usr/local/go/src/encoding/json/scanner.go:54
// the beginning of 12345e+6?).
//line /usr/local/go/src/encoding/json/scanner.go:66
type scanner struct {
	// The step is a func to be called to execute the next transition.
	// Also tried using an integer constant and a single func
	// with a switch, but using the func directly was 10% faster
	// on a 64-bit Mac Mini, and it's nicer to read.
	step	func(*scanner, byte) int

	// Reached end of top-level value.
	endTop	bool

	// Stack of what we're in the middle of - array values, object keys, object values.
	parseState	[]int

	// Error that happened, if any.
	err	error

	// total bytes consumed, updated by decoder.Decode (and deliberately
	// not set to zero by scan.reset)
	bytes	int64
}

var scannerPool = sync.Pool{
	New: func() any {
//line /usr/local/go/src/encoding/json/scanner.go:88
		_go_fuzz_dep_.CoverTab[28126]++
								return &scanner{}
//line /usr/local/go/src/encoding/json/scanner.go:89
		// _ = "end of CoverTab[28126]"
	},
}

func newScanner() *scanner {
//line /usr/local/go/src/encoding/json/scanner.go:93
	_go_fuzz_dep_.CoverTab[28127]++
							scan := scannerPool.Get().(*scanner)

							scan.bytes = 0
							scan.reset()
							return scan
//line /usr/local/go/src/encoding/json/scanner.go:98
	// _ = "end of CoverTab[28127]"
}

func freeScanner(scan *scanner) {
//line /usr/local/go/src/encoding/json/scanner.go:101
	_go_fuzz_dep_.CoverTab[28128]++

							if len(scan.parseState) > 1024 {
//line /usr/local/go/src/encoding/json/scanner.go:103
		_go_fuzz_dep_.CoverTab[28130]++
								scan.parseState = nil
//line /usr/local/go/src/encoding/json/scanner.go:104
		// _ = "end of CoverTab[28130]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:105
		_go_fuzz_dep_.CoverTab[28131]++
//line /usr/local/go/src/encoding/json/scanner.go:105
		// _ = "end of CoverTab[28131]"
//line /usr/local/go/src/encoding/json/scanner.go:105
	}
//line /usr/local/go/src/encoding/json/scanner.go:105
	// _ = "end of CoverTab[28128]"
//line /usr/local/go/src/encoding/json/scanner.go:105
	_go_fuzz_dep_.CoverTab[28129]++
							scannerPool.Put(scan)
//line /usr/local/go/src/encoding/json/scanner.go:106
	// _ = "end of CoverTab[28129]"
}

// These values are returned by the state transition functions
//line /usr/local/go/src/encoding/json/scanner.go:109
// assigned to scanner.state and the method scanner.eof.
//line /usr/local/go/src/encoding/json/scanner.go:109
// They give details about the current state of the scan that
//line /usr/local/go/src/encoding/json/scanner.go:109
// callers might be interested to know about.
//line /usr/local/go/src/encoding/json/scanner.go:109
// It is okay to ignore the return value of any particular
//line /usr/local/go/src/encoding/json/scanner.go:109
// call to scanner.state: if one call returns scanError,
//line /usr/local/go/src/encoding/json/scanner.go:109
// every subsequent call will return scanError too.
//line /usr/local/go/src/encoding/json/scanner.go:116
const (
	// Continue.
	scanContinue		= iota	// uninteresting byte
	scanBeginLiteral		// end implied by next result != scanContinue
	scanBeginObject			// begin object
	scanObjectKey			// just finished object key (string)
	scanObjectValue			// just finished non-last object value
	scanEndObject			// end object (implies scanObjectValue if possible)
	scanBeginArray			// begin array
	scanArrayValue			// just finished array value
	scanEndArray			// end array (implies scanArrayValue if possible)
	scanSkipSpace			// space byte; can skip; known to be last "continue" result

	// Stop.
	scanEnd		// top-level value ended *before* this byte; known to be first "stop" result
	scanError	// hit an error, scanner.err.
)

// These values are stored in the parseState stack.
//line /usr/local/go/src/encoding/json/scanner.go:134
// They give the current state of a composite value
//line /usr/local/go/src/encoding/json/scanner.go:134
// being scanned. If the parser is inside a nested value
//line /usr/local/go/src/encoding/json/scanner.go:134
// the parseState describes the nested state, outermost at entry 0.
//line /usr/local/go/src/encoding/json/scanner.go:138
const (
	parseObjectKey		= iota	// parsing object key (before colon)
	parseObjectValue		// parsing object value (after colon)
	parseArrayValue			// parsing array value
)

// This limits the max nesting depth to prevent stack overflow.
//line /usr/local/go/src/encoding/json/scanner.go:144
// This is permitted by https://tools.ietf.org/html/rfc7159#section-9
//line /usr/local/go/src/encoding/json/scanner.go:146
const maxNestingDepth = 10000

// reset prepares the scanner for use.
//line /usr/local/go/src/encoding/json/scanner.go:148
// It must be called before calling s.step.
//line /usr/local/go/src/encoding/json/scanner.go:150
func (s *scanner) reset() {
//line /usr/local/go/src/encoding/json/scanner.go:150
	_go_fuzz_dep_.CoverTab[28132]++
							s.step = stateBeginValue
							s.parseState = s.parseState[0:0]
							s.err = nil
							s.endTop = false
//line /usr/local/go/src/encoding/json/scanner.go:154
	// _ = "end of CoverTab[28132]"
}

// eof tells the scanner that the end of input has been reached.
//line /usr/local/go/src/encoding/json/scanner.go:157
// It returns a scan status just as s.step does.
//line /usr/local/go/src/encoding/json/scanner.go:159
func (s *scanner) eof() int {
//line /usr/local/go/src/encoding/json/scanner.go:159
	_go_fuzz_dep_.CoverTab[28133]++
							if s.err != nil {
//line /usr/local/go/src/encoding/json/scanner.go:160
		_go_fuzz_dep_.CoverTab[28138]++
								return scanError
//line /usr/local/go/src/encoding/json/scanner.go:161
		// _ = "end of CoverTab[28138]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:162
		_go_fuzz_dep_.CoverTab[28139]++
//line /usr/local/go/src/encoding/json/scanner.go:162
		// _ = "end of CoverTab[28139]"
//line /usr/local/go/src/encoding/json/scanner.go:162
	}
//line /usr/local/go/src/encoding/json/scanner.go:162
	// _ = "end of CoverTab[28133]"
//line /usr/local/go/src/encoding/json/scanner.go:162
	_go_fuzz_dep_.CoverTab[28134]++
							if s.endTop {
//line /usr/local/go/src/encoding/json/scanner.go:163
		_go_fuzz_dep_.CoverTab[28140]++
								return scanEnd
//line /usr/local/go/src/encoding/json/scanner.go:164
		// _ = "end of CoverTab[28140]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:165
		_go_fuzz_dep_.CoverTab[28141]++
//line /usr/local/go/src/encoding/json/scanner.go:165
		// _ = "end of CoverTab[28141]"
//line /usr/local/go/src/encoding/json/scanner.go:165
	}
//line /usr/local/go/src/encoding/json/scanner.go:165
	// _ = "end of CoverTab[28134]"
//line /usr/local/go/src/encoding/json/scanner.go:165
	_go_fuzz_dep_.CoverTab[28135]++
							s.step(s, ' ')
							if s.endTop {
//line /usr/local/go/src/encoding/json/scanner.go:167
		_go_fuzz_dep_.CoverTab[28142]++
								return scanEnd
//line /usr/local/go/src/encoding/json/scanner.go:168
		// _ = "end of CoverTab[28142]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:169
		_go_fuzz_dep_.CoverTab[28143]++
//line /usr/local/go/src/encoding/json/scanner.go:169
		// _ = "end of CoverTab[28143]"
//line /usr/local/go/src/encoding/json/scanner.go:169
	}
//line /usr/local/go/src/encoding/json/scanner.go:169
	// _ = "end of CoverTab[28135]"
//line /usr/local/go/src/encoding/json/scanner.go:169
	_go_fuzz_dep_.CoverTab[28136]++
							if s.err == nil {
//line /usr/local/go/src/encoding/json/scanner.go:170
		_go_fuzz_dep_.CoverTab[28144]++
								s.err = &SyntaxError{"unexpected end of JSON input", s.bytes}
//line /usr/local/go/src/encoding/json/scanner.go:171
		// _ = "end of CoverTab[28144]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:172
		_go_fuzz_dep_.CoverTab[28145]++
//line /usr/local/go/src/encoding/json/scanner.go:172
		// _ = "end of CoverTab[28145]"
//line /usr/local/go/src/encoding/json/scanner.go:172
	}
//line /usr/local/go/src/encoding/json/scanner.go:172
	// _ = "end of CoverTab[28136]"
//line /usr/local/go/src/encoding/json/scanner.go:172
	_go_fuzz_dep_.CoverTab[28137]++
							return scanError
//line /usr/local/go/src/encoding/json/scanner.go:173
	// _ = "end of CoverTab[28137]"
}

// pushParseState pushes a new parse state p onto the parse stack.
//line /usr/local/go/src/encoding/json/scanner.go:176
// an error state is returned if maxNestingDepth was exceeded, otherwise successState is returned.
//line /usr/local/go/src/encoding/json/scanner.go:178
func (s *scanner) pushParseState(c byte, newParseState int, successState int) int {
//line /usr/local/go/src/encoding/json/scanner.go:178
	_go_fuzz_dep_.CoverTab[28146]++
							s.parseState = append(s.parseState, newParseState)
							if len(s.parseState) <= maxNestingDepth {
//line /usr/local/go/src/encoding/json/scanner.go:180
		_go_fuzz_dep_.CoverTab[28148]++
								return successState
//line /usr/local/go/src/encoding/json/scanner.go:181
		// _ = "end of CoverTab[28148]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:182
		_go_fuzz_dep_.CoverTab[28149]++
//line /usr/local/go/src/encoding/json/scanner.go:182
		// _ = "end of CoverTab[28149]"
//line /usr/local/go/src/encoding/json/scanner.go:182
	}
//line /usr/local/go/src/encoding/json/scanner.go:182
	// _ = "end of CoverTab[28146]"
//line /usr/local/go/src/encoding/json/scanner.go:182
	_go_fuzz_dep_.CoverTab[28147]++
							return s.error(c, "exceeded max depth")
//line /usr/local/go/src/encoding/json/scanner.go:183
	// _ = "end of CoverTab[28147]"
}

// popParseState pops a parse state (already obtained) off the stack
//line /usr/local/go/src/encoding/json/scanner.go:186
// and updates s.step accordingly.
//line /usr/local/go/src/encoding/json/scanner.go:188
func (s *scanner) popParseState() {
//line /usr/local/go/src/encoding/json/scanner.go:188
	_go_fuzz_dep_.CoverTab[28150]++
							n := len(s.parseState) - 1
							s.parseState = s.parseState[0:n]
							if n == 0 {
//line /usr/local/go/src/encoding/json/scanner.go:191
		_go_fuzz_dep_.CoverTab[28151]++
								s.step = stateEndTop
								s.endTop = true
//line /usr/local/go/src/encoding/json/scanner.go:193
		// _ = "end of CoverTab[28151]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:194
		_go_fuzz_dep_.CoverTab[28152]++
								s.step = stateEndValue
//line /usr/local/go/src/encoding/json/scanner.go:195
		// _ = "end of CoverTab[28152]"
	}
//line /usr/local/go/src/encoding/json/scanner.go:196
	// _ = "end of CoverTab[28150]"
}

func isSpace(c byte) bool {
//line /usr/local/go/src/encoding/json/scanner.go:199
	_go_fuzz_dep_.CoverTab[28153]++
							return c <= ' ' && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:200
		_go_fuzz_dep_.CoverTab[28154]++
//line /usr/local/go/src/encoding/json/scanner.go:200
		return (c == ' ' || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:200
			_go_fuzz_dep_.CoverTab[28155]++
//line /usr/local/go/src/encoding/json/scanner.go:200
			return c == '\t'
//line /usr/local/go/src/encoding/json/scanner.go:200
			// _ = "end of CoverTab[28155]"
//line /usr/local/go/src/encoding/json/scanner.go:200
		}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:200
			_go_fuzz_dep_.CoverTab[28156]++
//line /usr/local/go/src/encoding/json/scanner.go:200
			return c == '\r'
//line /usr/local/go/src/encoding/json/scanner.go:200
			// _ = "end of CoverTab[28156]"
//line /usr/local/go/src/encoding/json/scanner.go:200
		}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:200
			_go_fuzz_dep_.CoverTab[28157]++
//line /usr/local/go/src/encoding/json/scanner.go:200
			return c == '\n'
//line /usr/local/go/src/encoding/json/scanner.go:200
			// _ = "end of CoverTab[28157]"
//line /usr/local/go/src/encoding/json/scanner.go:200
		}())
//line /usr/local/go/src/encoding/json/scanner.go:200
		// _ = "end of CoverTab[28154]"
//line /usr/local/go/src/encoding/json/scanner.go:200
	}()
//line /usr/local/go/src/encoding/json/scanner.go:200
	// _ = "end of CoverTab[28153]"
}

// stateBeginValueOrEmpty is the state after reading `[`.
func stateBeginValueOrEmpty(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:204
	_go_fuzz_dep_.CoverTab[28158]++
							if isSpace(c) {
//line /usr/local/go/src/encoding/json/scanner.go:205
		_go_fuzz_dep_.CoverTab[28161]++
								return scanSkipSpace
//line /usr/local/go/src/encoding/json/scanner.go:206
		// _ = "end of CoverTab[28161]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:207
		_go_fuzz_dep_.CoverTab[28162]++
//line /usr/local/go/src/encoding/json/scanner.go:207
		// _ = "end of CoverTab[28162]"
//line /usr/local/go/src/encoding/json/scanner.go:207
	}
//line /usr/local/go/src/encoding/json/scanner.go:207
	// _ = "end of CoverTab[28158]"
//line /usr/local/go/src/encoding/json/scanner.go:207
	_go_fuzz_dep_.CoverTab[28159]++
							if c == ']' {
//line /usr/local/go/src/encoding/json/scanner.go:208
		_go_fuzz_dep_.CoverTab[28163]++
								return stateEndValue(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:209
		// _ = "end of CoverTab[28163]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:210
		_go_fuzz_dep_.CoverTab[28164]++
//line /usr/local/go/src/encoding/json/scanner.go:210
		// _ = "end of CoverTab[28164]"
//line /usr/local/go/src/encoding/json/scanner.go:210
	}
//line /usr/local/go/src/encoding/json/scanner.go:210
	// _ = "end of CoverTab[28159]"
//line /usr/local/go/src/encoding/json/scanner.go:210
	_go_fuzz_dep_.CoverTab[28160]++
							return stateBeginValue(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:211
	// _ = "end of CoverTab[28160]"
}

// stateBeginValue is the state at the beginning of the input.
func stateBeginValue(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:215
	_go_fuzz_dep_.CoverTab[28165]++
							if isSpace(c) {
//line /usr/local/go/src/encoding/json/scanner.go:216
		_go_fuzz_dep_.CoverTab[28169]++
								return scanSkipSpace
//line /usr/local/go/src/encoding/json/scanner.go:217
		// _ = "end of CoverTab[28169]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:218
		_go_fuzz_dep_.CoverTab[28170]++
//line /usr/local/go/src/encoding/json/scanner.go:218
		// _ = "end of CoverTab[28170]"
//line /usr/local/go/src/encoding/json/scanner.go:218
	}
//line /usr/local/go/src/encoding/json/scanner.go:218
	// _ = "end of CoverTab[28165]"
//line /usr/local/go/src/encoding/json/scanner.go:218
	_go_fuzz_dep_.CoverTab[28166]++
							switch c {
	case '{':
//line /usr/local/go/src/encoding/json/scanner.go:220
		_go_fuzz_dep_.CoverTab[28171]++
								s.step = stateBeginStringOrEmpty
								return s.pushParseState(c, parseObjectKey, scanBeginObject)
//line /usr/local/go/src/encoding/json/scanner.go:222
		// _ = "end of CoverTab[28171]"
	case '[':
//line /usr/local/go/src/encoding/json/scanner.go:223
		_go_fuzz_dep_.CoverTab[28172]++
								s.step = stateBeginValueOrEmpty
								return s.pushParseState(c, parseArrayValue, scanBeginArray)
//line /usr/local/go/src/encoding/json/scanner.go:225
		// _ = "end of CoverTab[28172]"
	case '"':
//line /usr/local/go/src/encoding/json/scanner.go:226
		_go_fuzz_dep_.CoverTab[28173]++
								s.step = stateInString
								return scanBeginLiteral
//line /usr/local/go/src/encoding/json/scanner.go:228
		// _ = "end of CoverTab[28173]"
	case '-':
//line /usr/local/go/src/encoding/json/scanner.go:229
		_go_fuzz_dep_.CoverTab[28174]++
								s.step = stateNeg
								return scanBeginLiteral
//line /usr/local/go/src/encoding/json/scanner.go:231
		// _ = "end of CoverTab[28174]"
	case '0':
//line /usr/local/go/src/encoding/json/scanner.go:232
		_go_fuzz_dep_.CoverTab[28175]++
								s.step = state0
								return scanBeginLiteral
//line /usr/local/go/src/encoding/json/scanner.go:234
		// _ = "end of CoverTab[28175]"
	case 't':
//line /usr/local/go/src/encoding/json/scanner.go:235
		_go_fuzz_dep_.CoverTab[28176]++
								s.step = stateT
								return scanBeginLiteral
//line /usr/local/go/src/encoding/json/scanner.go:237
		// _ = "end of CoverTab[28176]"
	case 'f':
//line /usr/local/go/src/encoding/json/scanner.go:238
		_go_fuzz_dep_.CoverTab[28177]++
								s.step = stateF
								return scanBeginLiteral
//line /usr/local/go/src/encoding/json/scanner.go:240
		// _ = "end of CoverTab[28177]"
	case 'n':
//line /usr/local/go/src/encoding/json/scanner.go:241
		_go_fuzz_dep_.CoverTab[28178]++
								s.step = stateN
								return scanBeginLiteral
//line /usr/local/go/src/encoding/json/scanner.go:243
		// _ = "end of CoverTab[28178]"
//line /usr/local/go/src/encoding/json/scanner.go:243
	default:
//line /usr/local/go/src/encoding/json/scanner.go:243
		_go_fuzz_dep_.CoverTab[28179]++
//line /usr/local/go/src/encoding/json/scanner.go:243
		// _ = "end of CoverTab[28179]"
	}
//line /usr/local/go/src/encoding/json/scanner.go:244
	// _ = "end of CoverTab[28166]"
//line /usr/local/go/src/encoding/json/scanner.go:244
	_go_fuzz_dep_.CoverTab[28167]++
							if '1' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:245
		_go_fuzz_dep_.CoverTab[28180]++
//line /usr/local/go/src/encoding/json/scanner.go:245
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:245
		// _ = "end of CoverTab[28180]"
//line /usr/local/go/src/encoding/json/scanner.go:245
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:245
		_go_fuzz_dep_.CoverTab[28181]++
								s.step = state1
								return scanBeginLiteral
//line /usr/local/go/src/encoding/json/scanner.go:247
		// _ = "end of CoverTab[28181]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:248
		_go_fuzz_dep_.CoverTab[28182]++
//line /usr/local/go/src/encoding/json/scanner.go:248
		// _ = "end of CoverTab[28182]"
//line /usr/local/go/src/encoding/json/scanner.go:248
	}
//line /usr/local/go/src/encoding/json/scanner.go:248
	// _ = "end of CoverTab[28167]"
//line /usr/local/go/src/encoding/json/scanner.go:248
	_go_fuzz_dep_.CoverTab[28168]++
							return s.error(c, "looking for beginning of value")
//line /usr/local/go/src/encoding/json/scanner.go:249
	// _ = "end of CoverTab[28168]"
}

// stateBeginStringOrEmpty is the state after reading `{`.
func stateBeginStringOrEmpty(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:253
	_go_fuzz_dep_.CoverTab[28183]++
							if isSpace(c) {
//line /usr/local/go/src/encoding/json/scanner.go:254
		_go_fuzz_dep_.CoverTab[28186]++
								return scanSkipSpace
//line /usr/local/go/src/encoding/json/scanner.go:255
		// _ = "end of CoverTab[28186]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:256
		_go_fuzz_dep_.CoverTab[28187]++
//line /usr/local/go/src/encoding/json/scanner.go:256
		// _ = "end of CoverTab[28187]"
//line /usr/local/go/src/encoding/json/scanner.go:256
	}
//line /usr/local/go/src/encoding/json/scanner.go:256
	// _ = "end of CoverTab[28183]"
//line /usr/local/go/src/encoding/json/scanner.go:256
	_go_fuzz_dep_.CoverTab[28184]++
							if c == '}' {
//line /usr/local/go/src/encoding/json/scanner.go:257
		_go_fuzz_dep_.CoverTab[28188]++
								n := len(s.parseState)
								s.parseState[n-1] = parseObjectValue
								return stateEndValue(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:260
		// _ = "end of CoverTab[28188]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:261
		_go_fuzz_dep_.CoverTab[28189]++
//line /usr/local/go/src/encoding/json/scanner.go:261
		// _ = "end of CoverTab[28189]"
//line /usr/local/go/src/encoding/json/scanner.go:261
	}
//line /usr/local/go/src/encoding/json/scanner.go:261
	// _ = "end of CoverTab[28184]"
//line /usr/local/go/src/encoding/json/scanner.go:261
	_go_fuzz_dep_.CoverTab[28185]++
							return stateBeginString(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:262
	// _ = "end of CoverTab[28185]"
}

// stateBeginString is the state after reading `{"key": value,`.
func stateBeginString(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:266
	_go_fuzz_dep_.CoverTab[28190]++
							if isSpace(c) {
//line /usr/local/go/src/encoding/json/scanner.go:267
		_go_fuzz_dep_.CoverTab[28193]++
								return scanSkipSpace
//line /usr/local/go/src/encoding/json/scanner.go:268
		// _ = "end of CoverTab[28193]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:269
		_go_fuzz_dep_.CoverTab[28194]++
//line /usr/local/go/src/encoding/json/scanner.go:269
		// _ = "end of CoverTab[28194]"
//line /usr/local/go/src/encoding/json/scanner.go:269
	}
//line /usr/local/go/src/encoding/json/scanner.go:269
	// _ = "end of CoverTab[28190]"
//line /usr/local/go/src/encoding/json/scanner.go:269
	_go_fuzz_dep_.CoverTab[28191]++
							if c == '"' {
//line /usr/local/go/src/encoding/json/scanner.go:270
		_go_fuzz_dep_.CoverTab[28195]++
								s.step = stateInString
								return scanBeginLiteral
//line /usr/local/go/src/encoding/json/scanner.go:272
		// _ = "end of CoverTab[28195]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:273
		_go_fuzz_dep_.CoverTab[28196]++
//line /usr/local/go/src/encoding/json/scanner.go:273
		// _ = "end of CoverTab[28196]"
//line /usr/local/go/src/encoding/json/scanner.go:273
	}
//line /usr/local/go/src/encoding/json/scanner.go:273
	// _ = "end of CoverTab[28191]"
//line /usr/local/go/src/encoding/json/scanner.go:273
	_go_fuzz_dep_.CoverTab[28192]++
							return s.error(c, "looking for beginning of object key string")
//line /usr/local/go/src/encoding/json/scanner.go:274
	// _ = "end of CoverTab[28192]"
}

// stateEndValue is the state after completing a value,
//line /usr/local/go/src/encoding/json/scanner.go:277
// such as after reading `{}` or `true` or `["x"`.
//line /usr/local/go/src/encoding/json/scanner.go:279
func stateEndValue(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:279
	_go_fuzz_dep_.CoverTab[28197]++
							n := len(s.parseState)
							if n == 0 {
//line /usr/local/go/src/encoding/json/scanner.go:281
		_go_fuzz_dep_.CoverTab[28201]++

								s.step = stateEndTop
								s.endTop = true
								return stateEndTop(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:285
		// _ = "end of CoverTab[28201]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:286
		_go_fuzz_dep_.CoverTab[28202]++
//line /usr/local/go/src/encoding/json/scanner.go:286
		// _ = "end of CoverTab[28202]"
//line /usr/local/go/src/encoding/json/scanner.go:286
	}
//line /usr/local/go/src/encoding/json/scanner.go:286
	// _ = "end of CoverTab[28197]"
//line /usr/local/go/src/encoding/json/scanner.go:286
	_go_fuzz_dep_.CoverTab[28198]++
							if isSpace(c) {
//line /usr/local/go/src/encoding/json/scanner.go:287
		_go_fuzz_dep_.CoverTab[28203]++
								s.step = stateEndValue
								return scanSkipSpace
//line /usr/local/go/src/encoding/json/scanner.go:289
		// _ = "end of CoverTab[28203]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:290
		_go_fuzz_dep_.CoverTab[28204]++
//line /usr/local/go/src/encoding/json/scanner.go:290
		// _ = "end of CoverTab[28204]"
//line /usr/local/go/src/encoding/json/scanner.go:290
	}
//line /usr/local/go/src/encoding/json/scanner.go:290
	// _ = "end of CoverTab[28198]"
//line /usr/local/go/src/encoding/json/scanner.go:290
	_go_fuzz_dep_.CoverTab[28199]++
							ps := s.parseState[n-1]
							switch ps {
	case parseObjectKey:
//line /usr/local/go/src/encoding/json/scanner.go:293
		_go_fuzz_dep_.CoverTab[28205]++
								if c == ':' {
//line /usr/local/go/src/encoding/json/scanner.go:294
			_go_fuzz_dep_.CoverTab[28214]++
									s.parseState[n-1] = parseObjectValue
									s.step = stateBeginValue
									return scanObjectKey
//line /usr/local/go/src/encoding/json/scanner.go:297
			// _ = "end of CoverTab[28214]"
		} else {
//line /usr/local/go/src/encoding/json/scanner.go:298
			_go_fuzz_dep_.CoverTab[28215]++
//line /usr/local/go/src/encoding/json/scanner.go:298
			// _ = "end of CoverTab[28215]"
//line /usr/local/go/src/encoding/json/scanner.go:298
		}
//line /usr/local/go/src/encoding/json/scanner.go:298
		// _ = "end of CoverTab[28205]"
//line /usr/local/go/src/encoding/json/scanner.go:298
		_go_fuzz_dep_.CoverTab[28206]++
								return s.error(c, "after object key")
//line /usr/local/go/src/encoding/json/scanner.go:299
		// _ = "end of CoverTab[28206]"
	case parseObjectValue:
//line /usr/local/go/src/encoding/json/scanner.go:300
		_go_fuzz_dep_.CoverTab[28207]++
								if c == ',' {
//line /usr/local/go/src/encoding/json/scanner.go:301
			_go_fuzz_dep_.CoverTab[28216]++
									s.parseState[n-1] = parseObjectKey
									s.step = stateBeginString
									return scanObjectValue
//line /usr/local/go/src/encoding/json/scanner.go:304
			// _ = "end of CoverTab[28216]"
		} else {
//line /usr/local/go/src/encoding/json/scanner.go:305
			_go_fuzz_dep_.CoverTab[28217]++
//line /usr/local/go/src/encoding/json/scanner.go:305
			// _ = "end of CoverTab[28217]"
//line /usr/local/go/src/encoding/json/scanner.go:305
		}
//line /usr/local/go/src/encoding/json/scanner.go:305
		// _ = "end of CoverTab[28207]"
//line /usr/local/go/src/encoding/json/scanner.go:305
		_go_fuzz_dep_.CoverTab[28208]++
								if c == '}' {
//line /usr/local/go/src/encoding/json/scanner.go:306
			_go_fuzz_dep_.CoverTab[28218]++
									s.popParseState()
									return scanEndObject
//line /usr/local/go/src/encoding/json/scanner.go:308
			// _ = "end of CoverTab[28218]"
		} else {
//line /usr/local/go/src/encoding/json/scanner.go:309
			_go_fuzz_dep_.CoverTab[28219]++
//line /usr/local/go/src/encoding/json/scanner.go:309
			// _ = "end of CoverTab[28219]"
//line /usr/local/go/src/encoding/json/scanner.go:309
		}
//line /usr/local/go/src/encoding/json/scanner.go:309
		// _ = "end of CoverTab[28208]"
//line /usr/local/go/src/encoding/json/scanner.go:309
		_go_fuzz_dep_.CoverTab[28209]++
								return s.error(c, "after object key:value pair")
//line /usr/local/go/src/encoding/json/scanner.go:310
		// _ = "end of CoverTab[28209]"
	case parseArrayValue:
//line /usr/local/go/src/encoding/json/scanner.go:311
		_go_fuzz_dep_.CoverTab[28210]++
								if c == ',' {
//line /usr/local/go/src/encoding/json/scanner.go:312
			_go_fuzz_dep_.CoverTab[28220]++
									s.step = stateBeginValue
									return scanArrayValue
//line /usr/local/go/src/encoding/json/scanner.go:314
			// _ = "end of CoverTab[28220]"
		} else {
//line /usr/local/go/src/encoding/json/scanner.go:315
			_go_fuzz_dep_.CoverTab[28221]++
//line /usr/local/go/src/encoding/json/scanner.go:315
			// _ = "end of CoverTab[28221]"
//line /usr/local/go/src/encoding/json/scanner.go:315
		}
//line /usr/local/go/src/encoding/json/scanner.go:315
		// _ = "end of CoverTab[28210]"
//line /usr/local/go/src/encoding/json/scanner.go:315
		_go_fuzz_dep_.CoverTab[28211]++
								if c == ']' {
//line /usr/local/go/src/encoding/json/scanner.go:316
			_go_fuzz_dep_.CoverTab[28222]++
									s.popParseState()
									return scanEndArray
//line /usr/local/go/src/encoding/json/scanner.go:318
			// _ = "end of CoverTab[28222]"
		} else {
//line /usr/local/go/src/encoding/json/scanner.go:319
			_go_fuzz_dep_.CoverTab[28223]++
//line /usr/local/go/src/encoding/json/scanner.go:319
			// _ = "end of CoverTab[28223]"
//line /usr/local/go/src/encoding/json/scanner.go:319
		}
//line /usr/local/go/src/encoding/json/scanner.go:319
		// _ = "end of CoverTab[28211]"
//line /usr/local/go/src/encoding/json/scanner.go:319
		_go_fuzz_dep_.CoverTab[28212]++
								return s.error(c, "after array element")
//line /usr/local/go/src/encoding/json/scanner.go:320
		// _ = "end of CoverTab[28212]"
//line /usr/local/go/src/encoding/json/scanner.go:320
	default:
//line /usr/local/go/src/encoding/json/scanner.go:320
		_go_fuzz_dep_.CoverTab[28213]++
//line /usr/local/go/src/encoding/json/scanner.go:320
		// _ = "end of CoverTab[28213]"
	}
//line /usr/local/go/src/encoding/json/scanner.go:321
	// _ = "end of CoverTab[28199]"
//line /usr/local/go/src/encoding/json/scanner.go:321
	_go_fuzz_dep_.CoverTab[28200]++
							return s.error(c, "")
//line /usr/local/go/src/encoding/json/scanner.go:322
	// _ = "end of CoverTab[28200]"
}

// stateEndTop is the state after finishing the top-level value,
//line /usr/local/go/src/encoding/json/scanner.go:325
// such as after reading `{}` or `[1,2,3]`.
//line /usr/local/go/src/encoding/json/scanner.go:325
// Only space characters should be seen now.
//line /usr/local/go/src/encoding/json/scanner.go:328
func stateEndTop(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:328
	_go_fuzz_dep_.CoverTab[28224]++
							if !isSpace(c) {
//line /usr/local/go/src/encoding/json/scanner.go:329
		_go_fuzz_dep_.CoverTab[28226]++

								s.error(c, "after top-level value")
//line /usr/local/go/src/encoding/json/scanner.go:331
		// _ = "end of CoverTab[28226]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:332
		_go_fuzz_dep_.CoverTab[28227]++
//line /usr/local/go/src/encoding/json/scanner.go:332
		// _ = "end of CoverTab[28227]"
//line /usr/local/go/src/encoding/json/scanner.go:332
	}
//line /usr/local/go/src/encoding/json/scanner.go:332
	// _ = "end of CoverTab[28224]"
//line /usr/local/go/src/encoding/json/scanner.go:332
	_go_fuzz_dep_.CoverTab[28225]++
							return scanEnd
//line /usr/local/go/src/encoding/json/scanner.go:333
	// _ = "end of CoverTab[28225]"
}

// stateInString is the state after reading `"`.
func stateInString(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:337
	_go_fuzz_dep_.CoverTab[28228]++
							if c == '"' {
//line /usr/local/go/src/encoding/json/scanner.go:338
		_go_fuzz_dep_.CoverTab[28232]++
								s.step = stateEndValue
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:340
		// _ = "end of CoverTab[28232]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:341
		_go_fuzz_dep_.CoverTab[28233]++
//line /usr/local/go/src/encoding/json/scanner.go:341
		// _ = "end of CoverTab[28233]"
//line /usr/local/go/src/encoding/json/scanner.go:341
	}
//line /usr/local/go/src/encoding/json/scanner.go:341
	// _ = "end of CoverTab[28228]"
//line /usr/local/go/src/encoding/json/scanner.go:341
	_go_fuzz_dep_.CoverTab[28229]++
							if c == '\\' {
//line /usr/local/go/src/encoding/json/scanner.go:342
		_go_fuzz_dep_.CoverTab[28234]++
								s.step = stateInStringEsc
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:344
		// _ = "end of CoverTab[28234]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:345
		_go_fuzz_dep_.CoverTab[28235]++
//line /usr/local/go/src/encoding/json/scanner.go:345
		// _ = "end of CoverTab[28235]"
//line /usr/local/go/src/encoding/json/scanner.go:345
	}
//line /usr/local/go/src/encoding/json/scanner.go:345
	// _ = "end of CoverTab[28229]"
//line /usr/local/go/src/encoding/json/scanner.go:345
	_go_fuzz_dep_.CoverTab[28230]++
							if c < 0x20 {
//line /usr/local/go/src/encoding/json/scanner.go:346
		_go_fuzz_dep_.CoverTab[28236]++
								return s.error(c, "in string literal")
//line /usr/local/go/src/encoding/json/scanner.go:347
		// _ = "end of CoverTab[28236]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:348
		_go_fuzz_dep_.CoverTab[28237]++
//line /usr/local/go/src/encoding/json/scanner.go:348
		// _ = "end of CoverTab[28237]"
//line /usr/local/go/src/encoding/json/scanner.go:348
	}
//line /usr/local/go/src/encoding/json/scanner.go:348
	// _ = "end of CoverTab[28230]"
//line /usr/local/go/src/encoding/json/scanner.go:348
	_go_fuzz_dep_.CoverTab[28231]++
							return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:349
	// _ = "end of CoverTab[28231]"
}

// stateInStringEsc is the state after reading `"\` during a quoted string.
func stateInStringEsc(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:353
	_go_fuzz_dep_.CoverTab[28238]++
							switch c {
	case 'b', 'f', 'n', 'r', 't', '\\', '/', '"':
//line /usr/local/go/src/encoding/json/scanner.go:355
		_go_fuzz_dep_.CoverTab[28240]++
								s.step = stateInString
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:357
		// _ = "end of CoverTab[28240]"
	case 'u':
//line /usr/local/go/src/encoding/json/scanner.go:358
		_go_fuzz_dep_.CoverTab[28241]++
								s.step = stateInStringEscU
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:360
		// _ = "end of CoverTab[28241]"
//line /usr/local/go/src/encoding/json/scanner.go:360
	default:
//line /usr/local/go/src/encoding/json/scanner.go:360
		_go_fuzz_dep_.CoverTab[28242]++
//line /usr/local/go/src/encoding/json/scanner.go:360
		// _ = "end of CoverTab[28242]"
	}
//line /usr/local/go/src/encoding/json/scanner.go:361
	// _ = "end of CoverTab[28238]"
//line /usr/local/go/src/encoding/json/scanner.go:361
	_go_fuzz_dep_.CoverTab[28239]++
							return s.error(c, "in string escape code")
//line /usr/local/go/src/encoding/json/scanner.go:362
	// _ = "end of CoverTab[28239]"
}

// stateInStringEscU is the state after reading `"\u` during a quoted string.
func stateInStringEscU(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:366
	_go_fuzz_dep_.CoverTab[28243]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:367
		_go_fuzz_dep_.CoverTab[28245]++
//line /usr/local/go/src/encoding/json/scanner.go:367
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:367
		// _ = "end of CoverTab[28245]"
//line /usr/local/go/src/encoding/json/scanner.go:367
	}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:367
		_go_fuzz_dep_.CoverTab[28246]++
//line /usr/local/go/src/encoding/json/scanner.go:367
		return 'a' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:367
			_go_fuzz_dep_.CoverTab[28247]++
//line /usr/local/go/src/encoding/json/scanner.go:367
			return c <= 'f'
//line /usr/local/go/src/encoding/json/scanner.go:367
			// _ = "end of CoverTab[28247]"
//line /usr/local/go/src/encoding/json/scanner.go:367
		}()
//line /usr/local/go/src/encoding/json/scanner.go:367
		// _ = "end of CoverTab[28246]"
//line /usr/local/go/src/encoding/json/scanner.go:367
	}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:367
		_go_fuzz_dep_.CoverTab[28248]++
//line /usr/local/go/src/encoding/json/scanner.go:367
		return 'A' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:367
			_go_fuzz_dep_.CoverTab[28249]++
//line /usr/local/go/src/encoding/json/scanner.go:367
			return c <= 'F'
//line /usr/local/go/src/encoding/json/scanner.go:367
			// _ = "end of CoverTab[28249]"
//line /usr/local/go/src/encoding/json/scanner.go:367
		}()
//line /usr/local/go/src/encoding/json/scanner.go:367
		// _ = "end of CoverTab[28248]"
//line /usr/local/go/src/encoding/json/scanner.go:367
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:367
		_go_fuzz_dep_.CoverTab[28250]++
								s.step = stateInStringEscU1
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:369
		// _ = "end of CoverTab[28250]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:370
		_go_fuzz_dep_.CoverTab[28251]++
//line /usr/local/go/src/encoding/json/scanner.go:370
		// _ = "end of CoverTab[28251]"
//line /usr/local/go/src/encoding/json/scanner.go:370
	}
//line /usr/local/go/src/encoding/json/scanner.go:370
	// _ = "end of CoverTab[28243]"
//line /usr/local/go/src/encoding/json/scanner.go:370
	_go_fuzz_dep_.CoverTab[28244]++

							return s.error(c, "in \\u hexadecimal character escape")
//line /usr/local/go/src/encoding/json/scanner.go:372
	// _ = "end of CoverTab[28244]"
}

// stateInStringEscU1 is the state after reading `"\u1` during a quoted string.
func stateInStringEscU1(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:376
	_go_fuzz_dep_.CoverTab[28252]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:377
		_go_fuzz_dep_.CoverTab[28254]++
//line /usr/local/go/src/encoding/json/scanner.go:377
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:377
		// _ = "end of CoverTab[28254]"
//line /usr/local/go/src/encoding/json/scanner.go:377
	}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:377
		_go_fuzz_dep_.CoverTab[28255]++
//line /usr/local/go/src/encoding/json/scanner.go:377
		return 'a' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:377
			_go_fuzz_dep_.CoverTab[28256]++
//line /usr/local/go/src/encoding/json/scanner.go:377
			return c <= 'f'
//line /usr/local/go/src/encoding/json/scanner.go:377
			// _ = "end of CoverTab[28256]"
//line /usr/local/go/src/encoding/json/scanner.go:377
		}()
//line /usr/local/go/src/encoding/json/scanner.go:377
		// _ = "end of CoverTab[28255]"
//line /usr/local/go/src/encoding/json/scanner.go:377
	}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:377
		_go_fuzz_dep_.CoverTab[28257]++
//line /usr/local/go/src/encoding/json/scanner.go:377
		return 'A' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:377
			_go_fuzz_dep_.CoverTab[28258]++
//line /usr/local/go/src/encoding/json/scanner.go:377
			return c <= 'F'
//line /usr/local/go/src/encoding/json/scanner.go:377
			// _ = "end of CoverTab[28258]"
//line /usr/local/go/src/encoding/json/scanner.go:377
		}()
//line /usr/local/go/src/encoding/json/scanner.go:377
		// _ = "end of CoverTab[28257]"
//line /usr/local/go/src/encoding/json/scanner.go:377
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:377
		_go_fuzz_dep_.CoverTab[28259]++
								s.step = stateInStringEscU12
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:379
		// _ = "end of CoverTab[28259]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:380
		_go_fuzz_dep_.CoverTab[28260]++
//line /usr/local/go/src/encoding/json/scanner.go:380
		// _ = "end of CoverTab[28260]"
//line /usr/local/go/src/encoding/json/scanner.go:380
	}
//line /usr/local/go/src/encoding/json/scanner.go:380
	// _ = "end of CoverTab[28252]"
//line /usr/local/go/src/encoding/json/scanner.go:380
	_go_fuzz_dep_.CoverTab[28253]++

							return s.error(c, "in \\u hexadecimal character escape")
//line /usr/local/go/src/encoding/json/scanner.go:382
	// _ = "end of CoverTab[28253]"
}

// stateInStringEscU12 is the state after reading `"\u12` during a quoted string.
func stateInStringEscU12(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:386
	_go_fuzz_dep_.CoverTab[28261]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:387
		_go_fuzz_dep_.CoverTab[28263]++
//line /usr/local/go/src/encoding/json/scanner.go:387
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:387
		// _ = "end of CoverTab[28263]"
//line /usr/local/go/src/encoding/json/scanner.go:387
	}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:387
		_go_fuzz_dep_.CoverTab[28264]++
//line /usr/local/go/src/encoding/json/scanner.go:387
		return 'a' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:387
			_go_fuzz_dep_.CoverTab[28265]++
//line /usr/local/go/src/encoding/json/scanner.go:387
			return c <= 'f'
//line /usr/local/go/src/encoding/json/scanner.go:387
			// _ = "end of CoverTab[28265]"
//line /usr/local/go/src/encoding/json/scanner.go:387
		}()
//line /usr/local/go/src/encoding/json/scanner.go:387
		// _ = "end of CoverTab[28264]"
//line /usr/local/go/src/encoding/json/scanner.go:387
	}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:387
		_go_fuzz_dep_.CoverTab[28266]++
//line /usr/local/go/src/encoding/json/scanner.go:387
		return 'A' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:387
			_go_fuzz_dep_.CoverTab[28267]++
//line /usr/local/go/src/encoding/json/scanner.go:387
			return c <= 'F'
//line /usr/local/go/src/encoding/json/scanner.go:387
			// _ = "end of CoverTab[28267]"
//line /usr/local/go/src/encoding/json/scanner.go:387
		}()
//line /usr/local/go/src/encoding/json/scanner.go:387
		// _ = "end of CoverTab[28266]"
//line /usr/local/go/src/encoding/json/scanner.go:387
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:387
		_go_fuzz_dep_.CoverTab[28268]++
								s.step = stateInStringEscU123
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:389
		// _ = "end of CoverTab[28268]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:390
		_go_fuzz_dep_.CoverTab[28269]++
//line /usr/local/go/src/encoding/json/scanner.go:390
		// _ = "end of CoverTab[28269]"
//line /usr/local/go/src/encoding/json/scanner.go:390
	}
//line /usr/local/go/src/encoding/json/scanner.go:390
	// _ = "end of CoverTab[28261]"
//line /usr/local/go/src/encoding/json/scanner.go:390
	_go_fuzz_dep_.CoverTab[28262]++

							return s.error(c, "in \\u hexadecimal character escape")
//line /usr/local/go/src/encoding/json/scanner.go:392
	// _ = "end of CoverTab[28262]"
}

// stateInStringEscU123 is the state after reading `"\u123` during a quoted string.
func stateInStringEscU123(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:396
	_go_fuzz_dep_.CoverTab[28270]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:397
		_go_fuzz_dep_.CoverTab[28272]++
//line /usr/local/go/src/encoding/json/scanner.go:397
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:397
		// _ = "end of CoverTab[28272]"
//line /usr/local/go/src/encoding/json/scanner.go:397
	}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:397
		_go_fuzz_dep_.CoverTab[28273]++
//line /usr/local/go/src/encoding/json/scanner.go:397
		return 'a' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:397
			_go_fuzz_dep_.CoverTab[28274]++
//line /usr/local/go/src/encoding/json/scanner.go:397
			return c <= 'f'
//line /usr/local/go/src/encoding/json/scanner.go:397
			// _ = "end of CoverTab[28274]"
//line /usr/local/go/src/encoding/json/scanner.go:397
		}()
//line /usr/local/go/src/encoding/json/scanner.go:397
		// _ = "end of CoverTab[28273]"
//line /usr/local/go/src/encoding/json/scanner.go:397
	}() || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:397
		_go_fuzz_dep_.CoverTab[28275]++
//line /usr/local/go/src/encoding/json/scanner.go:397
		return 'A' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:397
			_go_fuzz_dep_.CoverTab[28276]++
//line /usr/local/go/src/encoding/json/scanner.go:397
			return c <= 'F'
//line /usr/local/go/src/encoding/json/scanner.go:397
			// _ = "end of CoverTab[28276]"
//line /usr/local/go/src/encoding/json/scanner.go:397
		}()
//line /usr/local/go/src/encoding/json/scanner.go:397
		// _ = "end of CoverTab[28275]"
//line /usr/local/go/src/encoding/json/scanner.go:397
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:397
		_go_fuzz_dep_.CoverTab[28277]++
								s.step = stateInString
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:399
		// _ = "end of CoverTab[28277]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:400
		_go_fuzz_dep_.CoverTab[28278]++
//line /usr/local/go/src/encoding/json/scanner.go:400
		// _ = "end of CoverTab[28278]"
//line /usr/local/go/src/encoding/json/scanner.go:400
	}
//line /usr/local/go/src/encoding/json/scanner.go:400
	// _ = "end of CoverTab[28270]"
//line /usr/local/go/src/encoding/json/scanner.go:400
	_go_fuzz_dep_.CoverTab[28271]++

							return s.error(c, "in \\u hexadecimal character escape")
//line /usr/local/go/src/encoding/json/scanner.go:402
	// _ = "end of CoverTab[28271]"
}

// stateNeg is the state after reading `-` during a number.
func stateNeg(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:406
	_go_fuzz_dep_.CoverTab[28279]++
							if c == '0' {
//line /usr/local/go/src/encoding/json/scanner.go:407
		_go_fuzz_dep_.CoverTab[28282]++
								s.step = state0
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:409
		// _ = "end of CoverTab[28282]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:410
		_go_fuzz_dep_.CoverTab[28283]++
//line /usr/local/go/src/encoding/json/scanner.go:410
		// _ = "end of CoverTab[28283]"
//line /usr/local/go/src/encoding/json/scanner.go:410
	}
//line /usr/local/go/src/encoding/json/scanner.go:410
	// _ = "end of CoverTab[28279]"
//line /usr/local/go/src/encoding/json/scanner.go:410
	_go_fuzz_dep_.CoverTab[28280]++
							if '1' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:411
		_go_fuzz_dep_.CoverTab[28284]++
//line /usr/local/go/src/encoding/json/scanner.go:411
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:411
		// _ = "end of CoverTab[28284]"
//line /usr/local/go/src/encoding/json/scanner.go:411
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:411
		_go_fuzz_dep_.CoverTab[28285]++
								s.step = state1
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:413
		// _ = "end of CoverTab[28285]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:414
		_go_fuzz_dep_.CoverTab[28286]++
//line /usr/local/go/src/encoding/json/scanner.go:414
		// _ = "end of CoverTab[28286]"
//line /usr/local/go/src/encoding/json/scanner.go:414
	}
//line /usr/local/go/src/encoding/json/scanner.go:414
	// _ = "end of CoverTab[28280]"
//line /usr/local/go/src/encoding/json/scanner.go:414
	_go_fuzz_dep_.CoverTab[28281]++
							return s.error(c, "in numeric literal")
//line /usr/local/go/src/encoding/json/scanner.go:415
	// _ = "end of CoverTab[28281]"
}

// state1 is the state after reading a non-zero integer during a number,
//line /usr/local/go/src/encoding/json/scanner.go:418
// such as after reading `1` or `100` but not `0`.
//line /usr/local/go/src/encoding/json/scanner.go:420
func state1(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:420
	_go_fuzz_dep_.CoverTab[28287]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:421
		_go_fuzz_dep_.CoverTab[28289]++
//line /usr/local/go/src/encoding/json/scanner.go:421
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:421
		// _ = "end of CoverTab[28289]"
//line /usr/local/go/src/encoding/json/scanner.go:421
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:421
		_go_fuzz_dep_.CoverTab[28290]++
								s.step = state1
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:423
		// _ = "end of CoverTab[28290]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:424
		_go_fuzz_dep_.CoverTab[28291]++
//line /usr/local/go/src/encoding/json/scanner.go:424
		// _ = "end of CoverTab[28291]"
//line /usr/local/go/src/encoding/json/scanner.go:424
	}
//line /usr/local/go/src/encoding/json/scanner.go:424
	// _ = "end of CoverTab[28287]"
//line /usr/local/go/src/encoding/json/scanner.go:424
	_go_fuzz_dep_.CoverTab[28288]++
							return state0(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:425
	// _ = "end of CoverTab[28288]"
}

// state0 is the state after reading `0` during a number.
func state0(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:429
	_go_fuzz_dep_.CoverTab[28292]++
							if c == '.' {
//line /usr/local/go/src/encoding/json/scanner.go:430
		_go_fuzz_dep_.CoverTab[28295]++
								s.step = stateDot
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:432
		// _ = "end of CoverTab[28295]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:433
		_go_fuzz_dep_.CoverTab[28296]++
//line /usr/local/go/src/encoding/json/scanner.go:433
		// _ = "end of CoverTab[28296]"
//line /usr/local/go/src/encoding/json/scanner.go:433
	}
//line /usr/local/go/src/encoding/json/scanner.go:433
	// _ = "end of CoverTab[28292]"
//line /usr/local/go/src/encoding/json/scanner.go:433
	_go_fuzz_dep_.CoverTab[28293]++
							if c == 'e' || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:434
		_go_fuzz_dep_.CoverTab[28297]++
//line /usr/local/go/src/encoding/json/scanner.go:434
		return c == 'E'
//line /usr/local/go/src/encoding/json/scanner.go:434
		// _ = "end of CoverTab[28297]"
//line /usr/local/go/src/encoding/json/scanner.go:434
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:434
		_go_fuzz_dep_.CoverTab[28298]++
								s.step = stateE
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:436
		// _ = "end of CoverTab[28298]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:437
		_go_fuzz_dep_.CoverTab[28299]++
//line /usr/local/go/src/encoding/json/scanner.go:437
		// _ = "end of CoverTab[28299]"
//line /usr/local/go/src/encoding/json/scanner.go:437
	}
//line /usr/local/go/src/encoding/json/scanner.go:437
	// _ = "end of CoverTab[28293]"
//line /usr/local/go/src/encoding/json/scanner.go:437
	_go_fuzz_dep_.CoverTab[28294]++
							return stateEndValue(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:438
	// _ = "end of CoverTab[28294]"
}

// stateDot is the state after reading the integer and decimal point in a number,
//line /usr/local/go/src/encoding/json/scanner.go:441
// such as after reading `1.`.
//line /usr/local/go/src/encoding/json/scanner.go:443
func stateDot(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:443
	_go_fuzz_dep_.CoverTab[28300]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:444
		_go_fuzz_dep_.CoverTab[28302]++
//line /usr/local/go/src/encoding/json/scanner.go:444
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:444
		// _ = "end of CoverTab[28302]"
//line /usr/local/go/src/encoding/json/scanner.go:444
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:444
		_go_fuzz_dep_.CoverTab[28303]++
								s.step = stateDot0
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:446
		// _ = "end of CoverTab[28303]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:447
		_go_fuzz_dep_.CoverTab[28304]++
//line /usr/local/go/src/encoding/json/scanner.go:447
		// _ = "end of CoverTab[28304]"
//line /usr/local/go/src/encoding/json/scanner.go:447
	}
//line /usr/local/go/src/encoding/json/scanner.go:447
	// _ = "end of CoverTab[28300]"
//line /usr/local/go/src/encoding/json/scanner.go:447
	_go_fuzz_dep_.CoverTab[28301]++
							return s.error(c, "after decimal point in numeric literal")
//line /usr/local/go/src/encoding/json/scanner.go:448
	// _ = "end of CoverTab[28301]"
}

// stateDot0 is the state after reading the integer, decimal point, and subsequent
//line /usr/local/go/src/encoding/json/scanner.go:451
// digits of a number, such as after reading `3.14`.
//line /usr/local/go/src/encoding/json/scanner.go:453
func stateDot0(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:453
	_go_fuzz_dep_.CoverTab[28305]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:454
		_go_fuzz_dep_.CoverTab[28308]++
//line /usr/local/go/src/encoding/json/scanner.go:454
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:454
		// _ = "end of CoverTab[28308]"
//line /usr/local/go/src/encoding/json/scanner.go:454
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:454
		_go_fuzz_dep_.CoverTab[28309]++
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:455
		// _ = "end of CoverTab[28309]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:456
		_go_fuzz_dep_.CoverTab[28310]++
//line /usr/local/go/src/encoding/json/scanner.go:456
		// _ = "end of CoverTab[28310]"
//line /usr/local/go/src/encoding/json/scanner.go:456
	}
//line /usr/local/go/src/encoding/json/scanner.go:456
	// _ = "end of CoverTab[28305]"
//line /usr/local/go/src/encoding/json/scanner.go:456
	_go_fuzz_dep_.CoverTab[28306]++
							if c == 'e' || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:457
		_go_fuzz_dep_.CoverTab[28311]++
//line /usr/local/go/src/encoding/json/scanner.go:457
		return c == 'E'
//line /usr/local/go/src/encoding/json/scanner.go:457
		// _ = "end of CoverTab[28311]"
//line /usr/local/go/src/encoding/json/scanner.go:457
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:457
		_go_fuzz_dep_.CoverTab[28312]++
								s.step = stateE
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:459
		// _ = "end of CoverTab[28312]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:460
		_go_fuzz_dep_.CoverTab[28313]++
//line /usr/local/go/src/encoding/json/scanner.go:460
		// _ = "end of CoverTab[28313]"
//line /usr/local/go/src/encoding/json/scanner.go:460
	}
//line /usr/local/go/src/encoding/json/scanner.go:460
	// _ = "end of CoverTab[28306]"
//line /usr/local/go/src/encoding/json/scanner.go:460
	_go_fuzz_dep_.CoverTab[28307]++
							return stateEndValue(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:461
	// _ = "end of CoverTab[28307]"
}

// stateE is the state after reading the mantissa and e in a number,
//line /usr/local/go/src/encoding/json/scanner.go:464
// such as after reading `314e` or `0.314e`.
//line /usr/local/go/src/encoding/json/scanner.go:466
func stateE(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:466
	_go_fuzz_dep_.CoverTab[28314]++
							if c == '+' || func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:467
		_go_fuzz_dep_.CoverTab[28316]++
//line /usr/local/go/src/encoding/json/scanner.go:467
		return c == '-'
//line /usr/local/go/src/encoding/json/scanner.go:467
		// _ = "end of CoverTab[28316]"
//line /usr/local/go/src/encoding/json/scanner.go:467
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:467
		_go_fuzz_dep_.CoverTab[28317]++
								s.step = stateESign
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:469
		// _ = "end of CoverTab[28317]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:470
		_go_fuzz_dep_.CoverTab[28318]++
//line /usr/local/go/src/encoding/json/scanner.go:470
		// _ = "end of CoverTab[28318]"
//line /usr/local/go/src/encoding/json/scanner.go:470
	}
//line /usr/local/go/src/encoding/json/scanner.go:470
	// _ = "end of CoverTab[28314]"
//line /usr/local/go/src/encoding/json/scanner.go:470
	_go_fuzz_dep_.CoverTab[28315]++
							return stateESign(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:471
	// _ = "end of CoverTab[28315]"
}

// stateESign is the state after reading the mantissa, e, and sign in a number,
//line /usr/local/go/src/encoding/json/scanner.go:474
// such as after reading `314e-` or `0.314e+`.
//line /usr/local/go/src/encoding/json/scanner.go:476
func stateESign(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:476
	_go_fuzz_dep_.CoverTab[28319]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:477
		_go_fuzz_dep_.CoverTab[28321]++
//line /usr/local/go/src/encoding/json/scanner.go:477
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:477
		// _ = "end of CoverTab[28321]"
//line /usr/local/go/src/encoding/json/scanner.go:477
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:477
		_go_fuzz_dep_.CoverTab[28322]++
								s.step = stateE0
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:479
		// _ = "end of CoverTab[28322]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:480
		_go_fuzz_dep_.CoverTab[28323]++
//line /usr/local/go/src/encoding/json/scanner.go:480
		// _ = "end of CoverTab[28323]"
//line /usr/local/go/src/encoding/json/scanner.go:480
	}
//line /usr/local/go/src/encoding/json/scanner.go:480
	// _ = "end of CoverTab[28319]"
//line /usr/local/go/src/encoding/json/scanner.go:480
	_go_fuzz_dep_.CoverTab[28320]++
							return s.error(c, "in exponent of numeric literal")
//line /usr/local/go/src/encoding/json/scanner.go:481
	// _ = "end of CoverTab[28320]"
}

// stateE0 is the state after reading the mantissa, e, optional sign,
//line /usr/local/go/src/encoding/json/scanner.go:484
// and at least one digit of the exponent in a number,
//line /usr/local/go/src/encoding/json/scanner.go:484
// such as after reading `314e-2` or `0.314e+1` or `3.14e0`.
//line /usr/local/go/src/encoding/json/scanner.go:487
func stateE0(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:487
	_go_fuzz_dep_.CoverTab[28324]++
							if '0' <= c && func() bool {
//line /usr/local/go/src/encoding/json/scanner.go:488
		_go_fuzz_dep_.CoverTab[28326]++
//line /usr/local/go/src/encoding/json/scanner.go:488
		return c <= '9'
//line /usr/local/go/src/encoding/json/scanner.go:488
		// _ = "end of CoverTab[28326]"
//line /usr/local/go/src/encoding/json/scanner.go:488
	}() {
//line /usr/local/go/src/encoding/json/scanner.go:488
		_go_fuzz_dep_.CoverTab[28327]++
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:489
		// _ = "end of CoverTab[28327]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:490
		_go_fuzz_dep_.CoverTab[28328]++
//line /usr/local/go/src/encoding/json/scanner.go:490
		// _ = "end of CoverTab[28328]"
//line /usr/local/go/src/encoding/json/scanner.go:490
	}
//line /usr/local/go/src/encoding/json/scanner.go:490
	// _ = "end of CoverTab[28324]"
//line /usr/local/go/src/encoding/json/scanner.go:490
	_go_fuzz_dep_.CoverTab[28325]++
							return stateEndValue(s, c)
//line /usr/local/go/src/encoding/json/scanner.go:491
	// _ = "end of CoverTab[28325]"
}

// stateT is the state after reading `t`.
func stateT(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:495
	_go_fuzz_dep_.CoverTab[28329]++
							if c == 'r' {
//line /usr/local/go/src/encoding/json/scanner.go:496
		_go_fuzz_dep_.CoverTab[28331]++
								s.step = stateTr
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:498
		// _ = "end of CoverTab[28331]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:499
		_go_fuzz_dep_.CoverTab[28332]++
//line /usr/local/go/src/encoding/json/scanner.go:499
		// _ = "end of CoverTab[28332]"
//line /usr/local/go/src/encoding/json/scanner.go:499
	}
//line /usr/local/go/src/encoding/json/scanner.go:499
	// _ = "end of CoverTab[28329]"
//line /usr/local/go/src/encoding/json/scanner.go:499
	_go_fuzz_dep_.CoverTab[28330]++
							return s.error(c, "in literal true (expecting 'r')")
//line /usr/local/go/src/encoding/json/scanner.go:500
	// _ = "end of CoverTab[28330]"
}

// stateTr is the state after reading `tr`.
func stateTr(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:504
	_go_fuzz_dep_.CoverTab[28333]++
							if c == 'u' {
//line /usr/local/go/src/encoding/json/scanner.go:505
		_go_fuzz_dep_.CoverTab[28335]++
								s.step = stateTru
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:507
		// _ = "end of CoverTab[28335]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:508
		_go_fuzz_dep_.CoverTab[28336]++
//line /usr/local/go/src/encoding/json/scanner.go:508
		// _ = "end of CoverTab[28336]"
//line /usr/local/go/src/encoding/json/scanner.go:508
	}
//line /usr/local/go/src/encoding/json/scanner.go:508
	// _ = "end of CoverTab[28333]"
//line /usr/local/go/src/encoding/json/scanner.go:508
	_go_fuzz_dep_.CoverTab[28334]++
							return s.error(c, "in literal true (expecting 'u')")
//line /usr/local/go/src/encoding/json/scanner.go:509
	// _ = "end of CoverTab[28334]"
}

// stateTru is the state after reading `tru`.
func stateTru(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:513
	_go_fuzz_dep_.CoverTab[28337]++
							if c == 'e' {
//line /usr/local/go/src/encoding/json/scanner.go:514
		_go_fuzz_dep_.CoverTab[28339]++
								s.step = stateEndValue
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:516
		// _ = "end of CoverTab[28339]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:517
		_go_fuzz_dep_.CoverTab[28340]++
//line /usr/local/go/src/encoding/json/scanner.go:517
		// _ = "end of CoverTab[28340]"
//line /usr/local/go/src/encoding/json/scanner.go:517
	}
//line /usr/local/go/src/encoding/json/scanner.go:517
	// _ = "end of CoverTab[28337]"
//line /usr/local/go/src/encoding/json/scanner.go:517
	_go_fuzz_dep_.CoverTab[28338]++
							return s.error(c, "in literal true (expecting 'e')")
//line /usr/local/go/src/encoding/json/scanner.go:518
	// _ = "end of CoverTab[28338]"
}

// stateF is the state after reading `f`.
func stateF(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:522
	_go_fuzz_dep_.CoverTab[28341]++
							if c == 'a' {
//line /usr/local/go/src/encoding/json/scanner.go:523
		_go_fuzz_dep_.CoverTab[28343]++
								s.step = stateFa
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:525
		// _ = "end of CoverTab[28343]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:526
		_go_fuzz_dep_.CoverTab[28344]++
//line /usr/local/go/src/encoding/json/scanner.go:526
		// _ = "end of CoverTab[28344]"
//line /usr/local/go/src/encoding/json/scanner.go:526
	}
//line /usr/local/go/src/encoding/json/scanner.go:526
	// _ = "end of CoverTab[28341]"
//line /usr/local/go/src/encoding/json/scanner.go:526
	_go_fuzz_dep_.CoverTab[28342]++
							return s.error(c, "in literal false (expecting 'a')")
//line /usr/local/go/src/encoding/json/scanner.go:527
	// _ = "end of CoverTab[28342]"
}

// stateFa is the state after reading `fa`.
func stateFa(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:531
	_go_fuzz_dep_.CoverTab[28345]++
							if c == 'l' {
//line /usr/local/go/src/encoding/json/scanner.go:532
		_go_fuzz_dep_.CoverTab[28347]++
								s.step = stateFal
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:534
		// _ = "end of CoverTab[28347]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:535
		_go_fuzz_dep_.CoverTab[28348]++
//line /usr/local/go/src/encoding/json/scanner.go:535
		// _ = "end of CoverTab[28348]"
//line /usr/local/go/src/encoding/json/scanner.go:535
	}
//line /usr/local/go/src/encoding/json/scanner.go:535
	// _ = "end of CoverTab[28345]"
//line /usr/local/go/src/encoding/json/scanner.go:535
	_go_fuzz_dep_.CoverTab[28346]++
							return s.error(c, "in literal false (expecting 'l')")
//line /usr/local/go/src/encoding/json/scanner.go:536
	// _ = "end of CoverTab[28346]"
}

// stateFal is the state after reading `fal`.
func stateFal(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:540
	_go_fuzz_dep_.CoverTab[28349]++
							if c == 's' {
//line /usr/local/go/src/encoding/json/scanner.go:541
		_go_fuzz_dep_.CoverTab[28351]++
								s.step = stateFals
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:543
		// _ = "end of CoverTab[28351]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:544
		_go_fuzz_dep_.CoverTab[28352]++
//line /usr/local/go/src/encoding/json/scanner.go:544
		// _ = "end of CoverTab[28352]"
//line /usr/local/go/src/encoding/json/scanner.go:544
	}
//line /usr/local/go/src/encoding/json/scanner.go:544
	// _ = "end of CoverTab[28349]"
//line /usr/local/go/src/encoding/json/scanner.go:544
	_go_fuzz_dep_.CoverTab[28350]++
							return s.error(c, "in literal false (expecting 's')")
//line /usr/local/go/src/encoding/json/scanner.go:545
	// _ = "end of CoverTab[28350]"
}

// stateFals is the state after reading `fals`.
func stateFals(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:549
	_go_fuzz_dep_.CoverTab[28353]++
							if c == 'e' {
//line /usr/local/go/src/encoding/json/scanner.go:550
		_go_fuzz_dep_.CoverTab[28355]++
								s.step = stateEndValue
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:552
		// _ = "end of CoverTab[28355]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:553
		_go_fuzz_dep_.CoverTab[28356]++
//line /usr/local/go/src/encoding/json/scanner.go:553
		// _ = "end of CoverTab[28356]"
//line /usr/local/go/src/encoding/json/scanner.go:553
	}
//line /usr/local/go/src/encoding/json/scanner.go:553
	// _ = "end of CoverTab[28353]"
//line /usr/local/go/src/encoding/json/scanner.go:553
	_go_fuzz_dep_.CoverTab[28354]++
							return s.error(c, "in literal false (expecting 'e')")
//line /usr/local/go/src/encoding/json/scanner.go:554
	// _ = "end of CoverTab[28354]"
}

// stateN is the state after reading `n`.
func stateN(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:558
	_go_fuzz_dep_.CoverTab[28357]++
							if c == 'u' {
//line /usr/local/go/src/encoding/json/scanner.go:559
		_go_fuzz_dep_.CoverTab[28359]++
								s.step = stateNu
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:561
		// _ = "end of CoverTab[28359]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:562
		_go_fuzz_dep_.CoverTab[28360]++
//line /usr/local/go/src/encoding/json/scanner.go:562
		// _ = "end of CoverTab[28360]"
//line /usr/local/go/src/encoding/json/scanner.go:562
	}
//line /usr/local/go/src/encoding/json/scanner.go:562
	// _ = "end of CoverTab[28357]"
//line /usr/local/go/src/encoding/json/scanner.go:562
	_go_fuzz_dep_.CoverTab[28358]++
							return s.error(c, "in literal null (expecting 'u')")
//line /usr/local/go/src/encoding/json/scanner.go:563
	// _ = "end of CoverTab[28358]"
}

// stateNu is the state after reading `nu`.
func stateNu(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:567
	_go_fuzz_dep_.CoverTab[28361]++
							if c == 'l' {
//line /usr/local/go/src/encoding/json/scanner.go:568
		_go_fuzz_dep_.CoverTab[28363]++
								s.step = stateNul
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:570
		// _ = "end of CoverTab[28363]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:571
		_go_fuzz_dep_.CoverTab[28364]++
//line /usr/local/go/src/encoding/json/scanner.go:571
		// _ = "end of CoverTab[28364]"
//line /usr/local/go/src/encoding/json/scanner.go:571
	}
//line /usr/local/go/src/encoding/json/scanner.go:571
	// _ = "end of CoverTab[28361]"
//line /usr/local/go/src/encoding/json/scanner.go:571
	_go_fuzz_dep_.CoverTab[28362]++
							return s.error(c, "in literal null (expecting 'l')")
//line /usr/local/go/src/encoding/json/scanner.go:572
	// _ = "end of CoverTab[28362]"
}

// stateNul is the state after reading `nul`.
func stateNul(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:576
	_go_fuzz_dep_.CoverTab[28365]++
							if c == 'l' {
//line /usr/local/go/src/encoding/json/scanner.go:577
		_go_fuzz_dep_.CoverTab[28367]++
								s.step = stateEndValue
								return scanContinue
//line /usr/local/go/src/encoding/json/scanner.go:579
		// _ = "end of CoverTab[28367]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:580
		_go_fuzz_dep_.CoverTab[28368]++
//line /usr/local/go/src/encoding/json/scanner.go:580
		// _ = "end of CoverTab[28368]"
//line /usr/local/go/src/encoding/json/scanner.go:580
	}
//line /usr/local/go/src/encoding/json/scanner.go:580
	// _ = "end of CoverTab[28365]"
//line /usr/local/go/src/encoding/json/scanner.go:580
	_go_fuzz_dep_.CoverTab[28366]++
							return s.error(c, "in literal null (expecting 'l')")
//line /usr/local/go/src/encoding/json/scanner.go:581
	// _ = "end of CoverTab[28366]"
}

// stateError is the state after reaching a syntax error,
//line /usr/local/go/src/encoding/json/scanner.go:584
// such as after reading `[1}` or `5.1.2`.
//line /usr/local/go/src/encoding/json/scanner.go:586
func stateError(s *scanner, c byte) int {
//line /usr/local/go/src/encoding/json/scanner.go:586
	_go_fuzz_dep_.CoverTab[28369]++
							return scanError
//line /usr/local/go/src/encoding/json/scanner.go:587
	// _ = "end of CoverTab[28369]"
}

// error records an error and switches to the error state.
func (s *scanner) error(c byte, context string) int {
//line /usr/local/go/src/encoding/json/scanner.go:591
	_go_fuzz_dep_.CoverTab[28370]++
							s.step = stateError
							s.err = &SyntaxError{"invalid character " + quoteChar(c) + " " + context, s.bytes}
							return scanError
//line /usr/local/go/src/encoding/json/scanner.go:594
	// _ = "end of CoverTab[28370]"
}

// quoteChar formats c as a quoted character literal.
func quoteChar(c byte) string {
//line /usr/local/go/src/encoding/json/scanner.go:598
	_go_fuzz_dep_.CoverTab[28371]++

							if c == '\'' {
//line /usr/local/go/src/encoding/json/scanner.go:600
		_go_fuzz_dep_.CoverTab[28374]++
								return `'\''`
//line /usr/local/go/src/encoding/json/scanner.go:601
		// _ = "end of CoverTab[28374]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:602
		_go_fuzz_dep_.CoverTab[28375]++
//line /usr/local/go/src/encoding/json/scanner.go:602
		// _ = "end of CoverTab[28375]"
//line /usr/local/go/src/encoding/json/scanner.go:602
	}
//line /usr/local/go/src/encoding/json/scanner.go:602
	// _ = "end of CoverTab[28371]"
//line /usr/local/go/src/encoding/json/scanner.go:602
	_go_fuzz_dep_.CoverTab[28372]++
							if c == '"' {
//line /usr/local/go/src/encoding/json/scanner.go:603
		_go_fuzz_dep_.CoverTab[28376]++
								return `'"'`
//line /usr/local/go/src/encoding/json/scanner.go:604
		// _ = "end of CoverTab[28376]"
	} else {
//line /usr/local/go/src/encoding/json/scanner.go:605
		_go_fuzz_dep_.CoverTab[28377]++
//line /usr/local/go/src/encoding/json/scanner.go:605
		// _ = "end of CoverTab[28377]"
//line /usr/local/go/src/encoding/json/scanner.go:605
	}
//line /usr/local/go/src/encoding/json/scanner.go:605
	// _ = "end of CoverTab[28372]"
//line /usr/local/go/src/encoding/json/scanner.go:605
	_go_fuzz_dep_.CoverTab[28373]++

//line /usr/local/go/src/encoding/json/scanner.go:608
	s := strconv.Quote(string(c))
							return "'" + s[1:len(s)-1] + "'"
//line /usr/local/go/src/encoding/json/scanner.go:609
	// _ = "end of CoverTab[28373]"
}

//line /usr/local/go/src/encoding/json/scanner.go:610
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/json/scanner.go:610
var _ = _go_fuzz_dep_.CoverTab
