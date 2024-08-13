// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:5
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:5
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:5
)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:16
import "strconv"

// checkValid verifies that data is valid JSON-encoded data.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:18
// scan is passed in for use by checkValid to avoid an allocation.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:20
func checkValid(data []byte, scan *scanner) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:20
	_go_fuzz_dep_.CoverTab[185401]++
											scan.reset()
											for _, c := range data {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:22
		_go_fuzz_dep_.CoverTab[185404]++
												scan.bytes++
												if scan.step(scan, c) == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:24
			_go_fuzz_dep_.CoverTab[185405]++
													return scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:25
			// _ = "end of CoverTab[185405]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:26
			_go_fuzz_dep_.CoverTab[185406]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:26
			// _ = "end of CoverTab[185406]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:26
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:26
		// _ = "end of CoverTab[185404]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:27
	// _ = "end of CoverTab[185401]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:27
	_go_fuzz_dep_.CoverTab[185402]++
											if scan.eof() == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:28
		_go_fuzz_dep_.CoverTab[185407]++
												return scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:29
		// _ = "end of CoverTab[185407]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:30
		_go_fuzz_dep_.CoverTab[185408]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:30
		// _ = "end of CoverTab[185408]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:30
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:30
	// _ = "end of CoverTab[185402]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:30
	_go_fuzz_dep_.CoverTab[185403]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:31
	// _ = "end of CoverTab[185403]"
}

// nextValue splits data after the next whole JSON value,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:34
// returning that value and the bytes that follow it as separate slices.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:34
// scan is passed in for use by nextValue to avoid an allocation.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:37
func nextValue(data []byte, scan *scanner) (value, rest []byte, err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:37
	_go_fuzz_dep_.CoverTab[185409]++
											scan.reset()
											for i, c := range data {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:39
		_go_fuzz_dep_.CoverTab[185412]++
												v := scan.step(scan, c)
												if v >= scanEndObject {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:41
			_go_fuzz_dep_.CoverTab[185413]++
													switch v {

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:46
			case scanEndObject, scanEndArray:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:46
				_go_fuzz_dep_.CoverTab[185414]++
														if scan.step(scan, ' ') == scanEnd {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:47
					_go_fuzz_dep_.CoverTab[185418]++
															return data[:i+1], data[i+1:], nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:48
					// _ = "end of CoverTab[185418]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:49
					_go_fuzz_dep_.CoverTab[185419]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:49
					// _ = "end of CoverTab[185419]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:49
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:49
				// _ = "end of CoverTab[185414]"
			case scanError:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:50
				_go_fuzz_dep_.CoverTab[185415]++
														return nil, nil, scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:51
				// _ = "end of CoverTab[185415]"
			case scanEnd:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:52
				_go_fuzz_dep_.CoverTab[185416]++
														return data[:i], data[i:], nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:53
				// _ = "end of CoverTab[185416]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:53
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:53
				_go_fuzz_dep_.CoverTab[185417]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:53
				// _ = "end of CoverTab[185417]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:54
			// _ = "end of CoverTab[185413]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:55
			_go_fuzz_dep_.CoverTab[185420]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:55
			// _ = "end of CoverTab[185420]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:55
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:55
		// _ = "end of CoverTab[185412]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:56
	// _ = "end of CoverTab[185409]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:56
	_go_fuzz_dep_.CoverTab[185410]++
											if scan.eof() == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:57
		_go_fuzz_dep_.CoverTab[185421]++
												return nil, nil, scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:58
		// _ = "end of CoverTab[185421]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:59
		_go_fuzz_dep_.CoverTab[185422]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:59
		// _ = "end of CoverTab[185422]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:59
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:59
	// _ = "end of CoverTab[185410]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:59
	_go_fuzz_dep_.CoverTab[185411]++
											return data, nil, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:60
	// _ = "end of CoverTab[185411]"
}

// A SyntaxError is a description of a JSON syntax error.
type SyntaxError struct {
	msg	string	// description of error
	Offset	int64	// error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:69
	_go_fuzz_dep_.CoverTab[185423]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:69
	return e.msg
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:69
	// _ = "end of CoverTab[185423]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:69
}

// A scanner is a JSON scanning state machine.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// Callers call scan.reset() and then pass bytes in one at a time
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// by calling scan.step(&scan, c) for each byte.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// The return value, referred to as an opcode, tells the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// caller about significant parsing events like beginning
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// and ending literals, objects, and arrays, so that the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// caller can follow along if it wishes.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// The return value scanEnd indicates that a single top-level
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// JSON value has been completed, *before* the byte that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// just got passed in.  (The indication must be delayed in order
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// to recognize the end of numbers: is 123 a whole value or
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:71
// the beginning of 12345e+6?).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:83
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

	// 1-byte redo (see undo method)
	redo		bool
	redoCode	int
	redoState	func(*scanner, byte) int

	// total bytes consumed, updated by decoder.Decode
	bytes	int64
}

// These values are returned by the state transition functions
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:108
// assigned to scanner.state and the method scanner.eof.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:108
// They give details about the current state of the scan that
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:108
// callers might be interested to know about.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:108
// It is okay to ignore the return value of any particular
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:108
// call to scanner.state: if one call returns scanError,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:108
// every subsequent call will return scanError too.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:115
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:133
// They give the current state of a composite value
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:133
// being scanned.  If the parser is inside a nested value
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:133
// the parseState describes the nested state, outermost at entry 0.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:137
const (
	parseObjectKey		= iota	// parsing object key (before colon)
	parseObjectValue		// parsing object value (after colon)
	parseArrayValue			// parsing array value
)

// reset prepares the scanner for use.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:143
// It must be called before calling s.step.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:145
func (s *scanner) reset() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:145
	_go_fuzz_dep_.CoverTab[185424]++
												s.step = stateBeginValue
												s.parseState = s.parseState[0:0]
												s.err = nil
												s.redo = false
												s.endTop = false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:150
	// _ = "end of CoverTab[185424]"
}

// eof tells the scanner that the end of input has been reached.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:153
// It returns a scan status just as s.step does.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:155
func (s *scanner) eof() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:155
	_go_fuzz_dep_.CoverTab[185425]++
												if s.err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:156
		_go_fuzz_dep_.CoverTab[185430]++
													return scanError
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:157
		// _ = "end of CoverTab[185430]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:158
		_go_fuzz_dep_.CoverTab[185431]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:158
		// _ = "end of CoverTab[185431]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:158
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:158
	// _ = "end of CoverTab[185425]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:158
	_go_fuzz_dep_.CoverTab[185426]++
												if s.endTop {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:159
		_go_fuzz_dep_.CoverTab[185432]++
													return scanEnd
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:160
		// _ = "end of CoverTab[185432]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:161
		_go_fuzz_dep_.CoverTab[185433]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:161
		// _ = "end of CoverTab[185433]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:161
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:161
	// _ = "end of CoverTab[185426]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:161
	_go_fuzz_dep_.CoverTab[185427]++
												s.step(s, ' ')
												if s.endTop {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:163
		_go_fuzz_dep_.CoverTab[185434]++
													return scanEnd
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:164
		// _ = "end of CoverTab[185434]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:165
		_go_fuzz_dep_.CoverTab[185435]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:165
		// _ = "end of CoverTab[185435]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:165
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:165
	// _ = "end of CoverTab[185427]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:165
	_go_fuzz_dep_.CoverTab[185428]++
												if s.err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:166
		_go_fuzz_dep_.CoverTab[185436]++
													s.err = &SyntaxError{"unexpected end of JSON input", s.bytes}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:167
		// _ = "end of CoverTab[185436]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:168
		_go_fuzz_dep_.CoverTab[185437]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:168
		// _ = "end of CoverTab[185437]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:168
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:168
	// _ = "end of CoverTab[185428]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:168
	_go_fuzz_dep_.CoverTab[185429]++
												return scanError
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:169
	// _ = "end of CoverTab[185429]"
}

// pushParseState pushes a new parse state p onto the parse stack.
func (s *scanner) pushParseState(p int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:173
	_go_fuzz_dep_.CoverTab[185438]++
												s.parseState = append(s.parseState, p)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:174
	// _ = "end of CoverTab[185438]"
}

// popParseState pops a parse state (already obtained) off the stack
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:177
// and updates s.step accordingly.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:179
func (s *scanner) popParseState() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:179
	_go_fuzz_dep_.CoverTab[185439]++
												n := len(s.parseState) - 1
												s.parseState = s.parseState[0:n]
												s.redo = false
												if n == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:183
		_go_fuzz_dep_.CoverTab[185440]++
													s.step = stateEndTop
													s.endTop = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:185
		// _ = "end of CoverTab[185440]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:186
		_go_fuzz_dep_.CoverTab[185441]++
													s.step = stateEndValue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:187
		// _ = "end of CoverTab[185441]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:188
	// _ = "end of CoverTab[185439]"
}

func isSpace(c byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:191
	_go_fuzz_dep_.CoverTab[185442]++
												return c == ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
		_go_fuzz_dep_.CoverTab[185443]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
		return c == '\t'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
		// _ = "end of CoverTab[185443]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
		_go_fuzz_dep_.CoverTab[185444]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
		return c == '\r'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
		// _ = "end of CoverTab[185444]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
		_go_fuzz_dep_.CoverTab[185445]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
		return c == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
		// _ = "end of CoverTab[185445]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:192
	// _ = "end of CoverTab[185442]"
}

// stateBeginValueOrEmpty is the state after reading `[`.
func stateBeginValueOrEmpty(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:196
	_go_fuzz_dep_.CoverTab[185446]++
												if c <= ' ' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:197
		_go_fuzz_dep_.CoverTab[185449]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:197
		return isSpace(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:197
		// _ = "end of CoverTab[185449]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:197
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:197
		_go_fuzz_dep_.CoverTab[185450]++
													return scanSkipSpace
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:198
		// _ = "end of CoverTab[185450]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:199
		_go_fuzz_dep_.CoverTab[185451]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:199
		// _ = "end of CoverTab[185451]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:199
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:199
	// _ = "end of CoverTab[185446]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:199
	_go_fuzz_dep_.CoverTab[185447]++
												if c == ']' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:200
		_go_fuzz_dep_.CoverTab[185452]++
													return stateEndValue(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:201
		// _ = "end of CoverTab[185452]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:202
		_go_fuzz_dep_.CoverTab[185453]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:202
		// _ = "end of CoverTab[185453]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:202
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:202
	// _ = "end of CoverTab[185447]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:202
	_go_fuzz_dep_.CoverTab[185448]++
												return stateBeginValue(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:203
	// _ = "end of CoverTab[185448]"
}

// stateBeginValue is the state at the beginning of the input.
func stateBeginValue(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:207
	_go_fuzz_dep_.CoverTab[185454]++
												if c <= ' ' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:208
		_go_fuzz_dep_.CoverTab[185458]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:208
		return isSpace(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:208
		// _ = "end of CoverTab[185458]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:208
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:208
		_go_fuzz_dep_.CoverTab[185459]++
													return scanSkipSpace
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:209
		// _ = "end of CoverTab[185459]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:210
		_go_fuzz_dep_.CoverTab[185460]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:210
		// _ = "end of CoverTab[185460]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:210
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:210
	// _ = "end of CoverTab[185454]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:210
	_go_fuzz_dep_.CoverTab[185455]++
												switch c {
	case '{':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:212
		_go_fuzz_dep_.CoverTab[185461]++
													s.step = stateBeginStringOrEmpty
													s.pushParseState(parseObjectKey)
													return scanBeginObject
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:215
		// _ = "end of CoverTab[185461]"
	case '[':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:216
		_go_fuzz_dep_.CoverTab[185462]++
													s.step = stateBeginValueOrEmpty
													s.pushParseState(parseArrayValue)
													return scanBeginArray
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:219
		// _ = "end of CoverTab[185462]"
	case '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:220
		_go_fuzz_dep_.CoverTab[185463]++
													s.step = stateInString
													return scanBeginLiteral
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:222
		// _ = "end of CoverTab[185463]"
	case '-':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:223
		_go_fuzz_dep_.CoverTab[185464]++
													s.step = stateNeg
													return scanBeginLiteral
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:225
		// _ = "end of CoverTab[185464]"
	case '0':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:226
		_go_fuzz_dep_.CoverTab[185465]++
													s.step = state0
													return scanBeginLiteral
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:228
		// _ = "end of CoverTab[185465]"
	case 't':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:229
		_go_fuzz_dep_.CoverTab[185466]++
													s.step = stateT
													return scanBeginLiteral
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:231
		// _ = "end of CoverTab[185466]"
	case 'f':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:232
		_go_fuzz_dep_.CoverTab[185467]++
													s.step = stateF
													return scanBeginLiteral
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:234
		// _ = "end of CoverTab[185467]"
	case 'n':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:235
		_go_fuzz_dep_.CoverTab[185468]++
													s.step = stateN
													return scanBeginLiteral
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:237
		// _ = "end of CoverTab[185468]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:237
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:237
		_go_fuzz_dep_.CoverTab[185469]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:237
		// _ = "end of CoverTab[185469]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:238
	// _ = "end of CoverTab[185455]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:238
	_go_fuzz_dep_.CoverTab[185456]++
												if '1' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:239
		_go_fuzz_dep_.CoverTab[185470]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:239
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:239
		// _ = "end of CoverTab[185470]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:239
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:239
		_go_fuzz_dep_.CoverTab[185471]++
													s.step = state1
													return scanBeginLiteral
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:241
		// _ = "end of CoverTab[185471]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:242
		_go_fuzz_dep_.CoverTab[185472]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:242
		// _ = "end of CoverTab[185472]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:242
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:242
	// _ = "end of CoverTab[185456]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:242
	_go_fuzz_dep_.CoverTab[185457]++
												return s.error(c, "looking for beginning of value")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:243
	// _ = "end of CoverTab[185457]"
}

// stateBeginStringOrEmpty is the state after reading `{`.
func stateBeginStringOrEmpty(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:247
	_go_fuzz_dep_.CoverTab[185473]++
												if c <= ' ' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:248
		_go_fuzz_dep_.CoverTab[185476]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:248
		return isSpace(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:248
		// _ = "end of CoverTab[185476]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:248
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:248
		_go_fuzz_dep_.CoverTab[185477]++
													return scanSkipSpace
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:249
		// _ = "end of CoverTab[185477]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:250
		_go_fuzz_dep_.CoverTab[185478]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:250
		// _ = "end of CoverTab[185478]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:250
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:250
	// _ = "end of CoverTab[185473]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:250
	_go_fuzz_dep_.CoverTab[185474]++
												if c == '}' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:251
		_go_fuzz_dep_.CoverTab[185479]++
													n := len(s.parseState)
													s.parseState[n-1] = parseObjectValue
													return stateEndValue(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:254
		// _ = "end of CoverTab[185479]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:255
		_go_fuzz_dep_.CoverTab[185480]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:255
		// _ = "end of CoverTab[185480]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:255
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:255
	// _ = "end of CoverTab[185474]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:255
	_go_fuzz_dep_.CoverTab[185475]++
												return stateBeginString(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:256
	// _ = "end of CoverTab[185475]"
}

// stateBeginString is the state after reading `{"key": value,`.
func stateBeginString(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:260
	_go_fuzz_dep_.CoverTab[185481]++
												if c <= ' ' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:261
		_go_fuzz_dep_.CoverTab[185484]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:261
		return isSpace(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:261
		// _ = "end of CoverTab[185484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:261
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:261
		_go_fuzz_dep_.CoverTab[185485]++
													return scanSkipSpace
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:262
		// _ = "end of CoverTab[185485]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:263
		_go_fuzz_dep_.CoverTab[185486]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:263
		// _ = "end of CoverTab[185486]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:263
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:263
	// _ = "end of CoverTab[185481]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:263
	_go_fuzz_dep_.CoverTab[185482]++
												if c == '"' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:264
		_go_fuzz_dep_.CoverTab[185487]++
													s.step = stateInString
													return scanBeginLiteral
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:266
		// _ = "end of CoverTab[185487]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:267
		_go_fuzz_dep_.CoverTab[185488]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:267
		// _ = "end of CoverTab[185488]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:267
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:267
	// _ = "end of CoverTab[185482]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:267
	_go_fuzz_dep_.CoverTab[185483]++
												return s.error(c, "looking for beginning of object key string")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:268
	// _ = "end of CoverTab[185483]"
}

// stateEndValue is the state after completing a value,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:271
// such as after reading `{}` or `true` or `["x"`.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:273
func stateEndValue(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:273
	_go_fuzz_dep_.CoverTab[185489]++
												n := len(s.parseState)
												if n == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:275
		_go_fuzz_dep_.CoverTab[185493]++

													s.step = stateEndTop
													s.endTop = true
													return stateEndTop(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:279
		// _ = "end of CoverTab[185493]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:280
		_go_fuzz_dep_.CoverTab[185494]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:280
		// _ = "end of CoverTab[185494]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:280
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:280
	// _ = "end of CoverTab[185489]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:280
	_go_fuzz_dep_.CoverTab[185490]++
												if c <= ' ' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:281
		_go_fuzz_dep_.CoverTab[185495]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:281
		return isSpace(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:281
		// _ = "end of CoverTab[185495]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:281
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:281
		_go_fuzz_dep_.CoverTab[185496]++
													s.step = stateEndValue
													return scanSkipSpace
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:283
		// _ = "end of CoverTab[185496]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:284
		_go_fuzz_dep_.CoverTab[185497]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:284
		// _ = "end of CoverTab[185497]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:284
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:284
	// _ = "end of CoverTab[185490]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:284
	_go_fuzz_dep_.CoverTab[185491]++
												ps := s.parseState[n-1]
												switch ps {
	case parseObjectKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:287
		_go_fuzz_dep_.CoverTab[185498]++
													if c == ':' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:288
			_go_fuzz_dep_.CoverTab[185507]++
														s.parseState[n-1] = parseObjectValue
														s.step = stateBeginValue
														return scanObjectKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:291
			// _ = "end of CoverTab[185507]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:292
			_go_fuzz_dep_.CoverTab[185508]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:292
			// _ = "end of CoverTab[185508]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:292
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:292
		// _ = "end of CoverTab[185498]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:292
		_go_fuzz_dep_.CoverTab[185499]++
													return s.error(c, "after object key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:293
		// _ = "end of CoverTab[185499]"
	case parseObjectValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:294
		_go_fuzz_dep_.CoverTab[185500]++
													if c == ',' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:295
			_go_fuzz_dep_.CoverTab[185509]++
														s.parseState[n-1] = parseObjectKey
														s.step = stateBeginString
														return scanObjectValue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:298
			// _ = "end of CoverTab[185509]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:299
			_go_fuzz_dep_.CoverTab[185510]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:299
			// _ = "end of CoverTab[185510]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:299
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:299
		// _ = "end of CoverTab[185500]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:299
		_go_fuzz_dep_.CoverTab[185501]++
													if c == '}' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:300
			_go_fuzz_dep_.CoverTab[185511]++
														s.popParseState()
														return scanEndObject
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:302
			// _ = "end of CoverTab[185511]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:303
			_go_fuzz_dep_.CoverTab[185512]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:303
			// _ = "end of CoverTab[185512]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:303
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:303
		// _ = "end of CoverTab[185501]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:303
		_go_fuzz_dep_.CoverTab[185502]++
													return s.error(c, "after object key:value pair")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:304
		// _ = "end of CoverTab[185502]"
	case parseArrayValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:305
		_go_fuzz_dep_.CoverTab[185503]++
													if c == ',' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:306
			_go_fuzz_dep_.CoverTab[185513]++
														s.step = stateBeginValue
														return scanArrayValue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:308
			// _ = "end of CoverTab[185513]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:309
			_go_fuzz_dep_.CoverTab[185514]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:309
			// _ = "end of CoverTab[185514]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:309
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:309
		// _ = "end of CoverTab[185503]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:309
		_go_fuzz_dep_.CoverTab[185504]++
													if c == ']' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:310
			_go_fuzz_dep_.CoverTab[185515]++
														s.popParseState()
														return scanEndArray
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:312
			// _ = "end of CoverTab[185515]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:313
			_go_fuzz_dep_.CoverTab[185516]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:313
			// _ = "end of CoverTab[185516]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:313
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:313
		// _ = "end of CoverTab[185504]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:313
		_go_fuzz_dep_.CoverTab[185505]++
													return s.error(c, "after array element")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:314
		// _ = "end of CoverTab[185505]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:314
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:314
		_go_fuzz_dep_.CoverTab[185506]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:314
		// _ = "end of CoverTab[185506]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:315
	// _ = "end of CoverTab[185491]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:315
	_go_fuzz_dep_.CoverTab[185492]++
												return s.error(c, "")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:316
	// _ = "end of CoverTab[185492]"
}

// stateEndTop is the state after finishing the top-level value,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:319
// such as after reading `{}` or `[1,2,3]`.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:319
// Only space characters should be seen now.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:322
func stateEndTop(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:322
	_go_fuzz_dep_.CoverTab[185517]++
												if c != ' ' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		_go_fuzz_dep_.CoverTab[185519]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		return c != '\t'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		// _ = "end of CoverTab[185519]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		_go_fuzz_dep_.CoverTab[185520]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		return c != '\r'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		// _ = "end of CoverTab[185520]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		_go_fuzz_dep_.CoverTab[185521]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		return c != '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		// _ = "end of CoverTab[185521]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:323
		_go_fuzz_dep_.CoverTab[185522]++

													s.error(c, "after top-level value")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:325
		// _ = "end of CoverTab[185522]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:326
		_go_fuzz_dep_.CoverTab[185523]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:326
		// _ = "end of CoverTab[185523]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:326
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:326
	// _ = "end of CoverTab[185517]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:326
	_go_fuzz_dep_.CoverTab[185518]++
												return scanEnd
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:327
	// _ = "end of CoverTab[185518]"
}

// stateInString is the state after reading `"`.
func stateInString(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:331
	_go_fuzz_dep_.CoverTab[185524]++
												if c == '"' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:332
		_go_fuzz_dep_.CoverTab[185528]++
													s.step = stateEndValue
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:334
		// _ = "end of CoverTab[185528]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:335
		_go_fuzz_dep_.CoverTab[185529]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:335
		// _ = "end of CoverTab[185529]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:335
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:335
	// _ = "end of CoverTab[185524]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:335
	_go_fuzz_dep_.CoverTab[185525]++
												if c == '\\' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:336
		_go_fuzz_dep_.CoverTab[185530]++
													s.step = stateInStringEsc
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:338
		// _ = "end of CoverTab[185530]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:339
		_go_fuzz_dep_.CoverTab[185531]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:339
		// _ = "end of CoverTab[185531]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:339
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:339
	// _ = "end of CoverTab[185525]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:339
	_go_fuzz_dep_.CoverTab[185526]++
												if c < 0x20 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:340
		_go_fuzz_dep_.CoverTab[185532]++
													return s.error(c, "in string literal")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:341
		// _ = "end of CoverTab[185532]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:342
		_go_fuzz_dep_.CoverTab[185533]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:342
		// _ = "end of CoverTab[185533]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:342
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:342
	// _ = "end of CoverTab[185526]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:342
	_go_fuzz_dep_.CoverTab[185527]++
												return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:343
	// _ = "end of CoverTab[185527]"
}

// stateInStringEsc is the state after reading `"\` during a quoted string.
func stateInStringEsc(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:347
	_go_fuzz_dep_.CoverTab[185534]++
												switch c {
	case 'b', 'f', 'n', 'r', 't', '\\', '/', '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:349
		_go_fuzz_dep_.CoverTab[185536]++
													s.step = stateInString
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:351
		// _ = "end of CoverTab[185536]"
	case 'u':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:352
		_go_fuzz_dep_.CoverTab[185537]++
													s.step = stateInStringEscU
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:354
		// _ = "end of CoverTab[185537]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:354
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:354
		_go_fuzz_dep_.CoverTab[185538]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:354
		// _ = "end of CoverTab[185538]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:355
	// _ = "end of CoverTab[185534]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:355
	_go_fuzz_dep_.CoverTab[185535]++
												return s.error(c, "in string escape code")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:356
	// _ = "end of CoverTab[185535]"
}

// stateInStringEscU is the state after reading `"\u` during a quoted string.
func stateInStringEscU(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:360
	_go_fuzz_dep_.CoverTab[185539]++
												if '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		_go_fuzz_dep_.CoverTab[185541]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		// _ = "end of CoverTab[185541]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		_go_fuzz_dep_.CoverTab[185542]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		return 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
			_go_fuzz_dep_.CoverTab[185543]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
			return c <= 'f'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
			// _ = "end of CoverTab[185543]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		// _ = "end of CoverTab[185542]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		_go_fuzz_dep_.CoverTab[185544]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		return 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
			_go_fuzz_dep_.CoverTab[185545]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
			return c <= 'F'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
			// _ = "end of CoverTab[185545]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		// _ = "end of CoverTab[185544]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:361
		_go_fuzz_dep_.CoverTab[185546]++
													s.step = stateInStringEscU1
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:363
		// _ = "end of CoverTab[185546]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:364
		_go_fuzz_dep_.CoverTab[185547]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:364
		// _ = "end of CoverTab[185547]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:364
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:364
	// _ = "end of CoverTab[185539]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:364
	_go_fuzz_dep_.CoverTab[185540]++

												return s.error(c, "in \\u hexadecimal character escape")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:366
	// _ = "end of CoverTab[185540]"
}

// stateInStringEscU1 is the state after reading `"\u1` during a quoted string.
func stateInStringEscU1(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:370
	_go_fuzz_dep_.CoverTab[185548]++
												if '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		_go_fuzz_dep_.CoverTab[185550]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		// _ = "end of CoverTab[185550]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		_go_fuzz_dep_.CoverTab[185551]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		return 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
			_go_fuzz_dep_.CoverTab[185552]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
			return c <= 'f'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
			// _ = "end of CoverTab[185552]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		// _ = "end of CoverTab[185551]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		_go_fuzz_dep_.CoverTab[185553]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		return 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
			_go_fuzz_dep_.CoverTab[185554]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
			return c <= 'F'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
			// _ = "end of CoverTab[185554]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		// _ = "end of CoverTab[185553]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:371
		_go_fuzz_dep_.CoverTab[185555]++
													s.step = stateInStringEscU12
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:373
		// _ = "end of CoverTab[185555]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:374
		_go_fuzz_dep_.CoverTab[185556]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:374
		// _ = "end of CoverTab[185556]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:374
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:374
	// _ = "end of CoverTab[185548]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:374
	_go_fuzz_dep_.CoverTab[185549]++

												return s.error(c, "in \\u hexadecimal character escape")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:376
	// _ = "end of CoverTab[185549]"
}

// stateInStringEscU12 is the state after reading `"\u12` during a quoted string.
func stateInStringEscU12(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:380
	_go_fuzz_dep_.CoverTab[185557]++
												if '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		_go_fuzz_dep_.CoverTab[185559]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		// _ = "end of CoverTab[185559]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		_go_fuzz_dep_.CoverTab[185560]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		return 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
			_go_fuzz_dep_.CoverTab[185561]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
			return c <= 'f'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
			// _ = "end of CoverTab[185561]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		// _ = "end of CoverTab[185560]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		_go_fuzz_dep_.CoverTab[185562]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		return 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
			_go_fuzz_dep_.CoverTab[185563]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
			return c <= 'F'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
			// _ = "end of CoverTab[185563]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		// _ = "end of CoverTab[185562]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:381
		_go_fuzz_dep_.CoverTab[185564]++
													s.step = stateInStringEscU123
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:383
		// _ = "end of CoverTab[185564]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:384
		_go_fuzz_dep_.CoverTab[185565]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:384
		// _ = "end of CoverTab[185565]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:384
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:384
	// _ = "end of CoverTab[185557]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:384
	_go_fuzz_dep_.CoverTab[185558]++

												return s.error(c, "in \\u hexadecimal character escape")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:386
	// _ = "end of CoverTab[185558]"
}

// stateInStringEscU123 is the state after reading `"\u123` during a quoted string.
func stateInStringEscU123(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:390
	_go_fuzz_dep_.CoverTab[185566]++
												if '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		_go_fuzz_dep_.CoverTab[185568]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		// _ = "end of CoverTab[185568]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		_go_fuzz_dep_.CoverTab[185569]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		return 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
			_go_fuzz_dep_.CoverTab[185570]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
			return c <= 'f'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
			// _ = "end of CoverTab[185570]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		// _ = "end of CoverTab[185569]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		_go_fuzz_dep_.CoverTab[185571]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		return 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
			_go_fuzz_dep_.CoverTab[185572]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
			return c <= 'F'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
			// _ = "end of CoverTab[185572]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		// _ = "end of CoverTab[185571]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:391
		_go_fuzz_dep_.CoverTab[185573]++
													s.step = stateInString
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:393
		// _ = "end of CoverTab[185573]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:394
		_go_fuzz_dep_.CoverTab[185574]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:394
		// _ = "end of CoverTab[185574]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:394
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:394
	// _ = "end of CoverTab[185566]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:394
	_go_fuzz_dep_.CoverTab[185567]++

												return s.error(c, "in \\u hexadecimal character escape")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:396
	// _ = "end of CoverTab[185567]"
}

// stateNeg is the state after reading `-` during a number.
func stateNeg(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:400
	_go_fuzz_dep_.CoverTab[185575]++
												if c == '0' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:401
		_go_fuzz_dep_.CoverTab[185578]++
													s.step = state0
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:403
		// _ = "end of CoverTab[185578]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:404
		_go_fuzz_dep_.CoverTab[185579]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:404
		// _ = "end of CoverTab[185579]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:404
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:404
	// _ = "end of CoverTab[185575]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:404
	_go_fuzz_dep_.CoverTab[185576]++
												if '1' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:405
		_go_fuzz_dep_.CoverTab[185580]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:405
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:405
		// _ = "end of CoverTab[185580]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:405
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:405
		_go_fuzz_dep_.CoverTab[185581]++
													s.step = state1
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:407
		// _ = "end of CoverTab[185581]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:408
		_go_fuzz_dep_.CoverTab[185582]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:408
		// _ = "end of CoverTab[185582]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:408
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:408
	// _ = "end of CoverTab[185576]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:408
	_go_fuzz_dep_.CoverTab[185577]++
												return s.error(c, "in numeric literal")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:409
	// _ = "end of CoverTab[185577]"
}

// state1 is the state after reading a non-zero integer during a number,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:412
// such as after reading `1` or `100` but not `0`.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:414
func state1(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:414
	_go_fuzz_dep_.CoverTab[185583]++
												if '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:415
		_go_fuzz_dep_.CoverTab[185585]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:415
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:415
		// _ = "end of CoverTab[185585]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:415
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:415
		_go_fuzz_dep_.CoverTab[185586]++
													s.step = state1
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:417
		// _ = "end of CoverTab[185586]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:418
		_go_fuzz_dep_.CoverTab[185587]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:418
		// _ = "end of CoverTab[185587]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:418
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:418
	// _ = "end of CoverTab[185583]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:418
	_go_fuzz_dep_.CoverTab[185584]++
												return state0(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:419
	// _ = "end of CoverTab[185584]"
}

// state0 is the state after reading `0` during a number.
func state0(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:423
	_go_fuzz_dep_.CoverTab[185588]++
												if c == '.' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:424
		_go_fuzz_dep_.CoverTab[185591]++
													s.step = stateDot
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:426
		// _ = "end of CoverTab[185591]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:427
		_go_fuzz_dep_.CoverTab[185592]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:427
		// _ = "end of CoverTab[185592]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:427
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:427
	// _ = "end of CoverTab[185588]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:427
	_go_fuzz_dep_.CoverTab[185589]++
												if c == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:428
		_go_fuzz_dep_.CoverTab[185593]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:428
		return c == 'E'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:428
		// _ = "end of CoverTab[185593]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:428
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:428
		_go_fuzz_dep_.CoverTab[185594]++
													s.step = stateE
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:430
		// _ = "end of CoverTab[185594]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:431
		_go_fuzz_dep_.CoverTab[185595]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:431
		// _ = "end of CoverTab[185595]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:431
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:431
	// _ = "end of CoverTab[185589]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:431
	_go_fuzz_dep_.CoverTab[185590]++
												return stateEndValue(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:432
	// _ = "end of CoverTab[185590]"
}

// stateDot is the state after reading the integer and decimal point in a number,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:435
// such as after reading `1.`.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:437
func stateDot(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:437
	_go_fuzz_dep_.CoverTab[185596]++
												if '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:438
		_go_fuzz_dep_.CoverTab[185598]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:438
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:438
		// _ = "end of CoverTab[185598]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:438
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:438
		_go_fuzz_dep_.CoverTab[185599]++
													s.step = stateDot0
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:440
		// _ = "end of CoverTab[185599]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:441
		_go_fuzz_dep_.CoverTab[185600]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:441
		// _ = "end of CoverTab[185600]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:441
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:441
	// _ = "end of CoverTab[185596]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:441
	_go_fuzz_dep_.CoverTab[185597]++
												return s.error(c, "after decimal point in numeric literal")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:442
	// _ = "end of CoverTab[185597]"
}

// stateDot0 is the state after reading the integer, decimal point, and subsequent
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:445
// digits of a number, such as after reading `3.14`.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:447
func stateDot0(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:447
	_go_fuzz_dep_.CoverTab[185601]++
												if '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:448
		_go_fuzz_dep_.CoverTab[185604]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:448
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:448
		// _ = "end of CoverTab[185604]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:448
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:448
		_go_fuzz_dep_.CoverTab[185605]++
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:449
		// _ = "end of CoverTab[185605]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:450
		_go_fuzz_dep_.CoverTab[185606]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:450
		// _ = "end of CoverTab[185606]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:450
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:450
	// _ = "end of CoverTab[185601]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:450
	_go_fuzz_dep_.CoverTab[185602]++
												if c == 'e' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:451
		_go_fuzz_dep_.CoverTab[185607]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:451
		return c == 'E'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:451
		// _ = "end of CoverTab[185607]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:451
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:451
		_go_fuzz_dep_.CoverTab[185608]++
													s.step = stateE
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:453
		// _ = "end of CoverTab[185608]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:454
		_go_fuzz_dep_.CoverTab[185609]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:454
		// _ = "end of CoverTab[185609]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:454
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:454
	// _ = "end of CoverTab[185602]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:454
	_go_fuzz_dep_.CoverTab[185603]++
												return stateEndValue(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:455
	// _ = "end of CoverTab[185603]"
}

// stateE is the state after reading the mantissa and e in a number,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:458
// such as after reading `314e` or `0.314e`.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:460
func stateE(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:460
	_go_fuzz_dep_.CoverTab[185610]++
												if c == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:461
		_go_fuzz_dep_.CoverTab[185612]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:461
		return c == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:461
		// _ = "end of CoverTab[185612]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:461
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:461
		_go_fuzz_dep_.CoverTab[185613]++
													s.step = stateESign
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:463
		// _ = "end of CoverTab[185613]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:464
		_go_fuzz_dep_.CoverTab[185614]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:464
		// _ = "end of CoverTab[185614]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:464
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:464
	// _ = "end of CoverTab[185610]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:464
	_go_fuzz_dep_.CoverTab[185611]++
												return stateESign(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:465
	// _ = "end of CoverTab[185611]"
}

// stateESign is the state after reading the mantissa, e, and sign in a number,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:468
// such as after reading `314e-` or `0.314e+`.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:470
func stateESign(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:470
	_go_fuzz_dep_.CoverTab[185615]++
												if '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:471
		_go_fuzz_dep_.CoverTab[185617]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:471
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:471
		// _ = "end of CoverTab[185617]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:471
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:471
		_go_fuzz_dep_.CoverTab[185618]++
													s.step = stateE0
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:473
		// _ = "end of CoverTab[185618]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:474
		_go_fuzz_dep_.CoverTab[185619]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:474
		// _ = "end of CoverTab[185619]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:474
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:474
	// _ = "end of CoverTab[185615]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:474
	_go_fuzz_dep_.CoverTab[185616]++
												return s.error(c, "in exponent of numeric literal")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:475
	// _ = "end of CoverTab[185616]"
}

// stateE0 is the state after reading the mantissa, e, optional sign,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:478
// and at least one digit of the exponent in a number,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:478
// such as after reading `314e-2` or `0.314e+1` or `3.14e0`.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:481
func stateE0(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:481
	_go_fuzz_dep_.CoverTab[185620]++
												if '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:482
		_go_fuzz_dep_.CoverTab[185622]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:482
		return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:482
		// _ = "end of CoverTab[185622]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:482
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:482
		_go_fuzz_dep_.CoverTab[185623]++
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:483
		// _ = "end of CoverTab[185623]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:484
		_go_fuzz_dep_.CoverTab[185624]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:484
		// _ = "end of CoverTab[185624]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:484
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:484
	// _ = "end of CoverTab[185620]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:484
	_go_fuzz_dep_.CoverTab[185621]++
												return stateEndValue(s, c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:485
	// _ = "end of CoverTab[185621]"
}

// stateT is the state after reading `t`.
func stateT(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:489
	_go_fuzz_dep_.CoverTab[185625]++
												if c == 'r' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:490
		_go_fuzz_dep_.CoverTab[185627]++
													s.step = stateTr
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:492
		// _ = "end of CoverTab[185627]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:493
		_go_fuzz_dep_.CoverTab[185628]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:493
		// _ = "end of CoverTab[185628]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:493
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:493
	// _ = "end of CoverTab[185625]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:493
	_go_fuzz_dep_.CoverTab[185626]++
												return s.error(c, "in literal true (expecting 'r')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:494
	// _ = "end of CoverTab[185626]"
}

// stateTr is the state after reading `tr`.
func stateTr(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:498
	_go_fuzz_dep_.CoverTab[185629]++
												if c == 'u' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:499
		_go_fuzz_dep_.CoverTab[185631]++
													s.step = stateTru
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:501
		// _ = "end of CoverTab[185631]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:502
		_go_fuzz_dep_.CoverTab[185632]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:502
		// _ = "end of CoverTab[185632]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:502
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:502
	// _ = "end of CoverTab[185629]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:502
	_go_fuzz_dep_.CoverTab[185630]++
												return s.error(c, "in literal true (expecting 'u')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:503
	// _ = "end of CoverTab[185630]"
}

// stateTru is the state after reading `tru`.
func stateTru(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:507
	_go_fuzz_dep_.CoverTab[185633]++
												if c == 'e' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:508
		_go_fuzz_dep_.CoverTab[185635]++
													s.step = stateEndValue
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:510
		// _ = "end of CoverTab[185635]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:511
		_go_fuzz_dep_.CoverTab[185636]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:511
		// _ = "end of CoverTab[185636]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:511
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:511
	// _ = "end of CoverTab[185633]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:511
	_go_fuzz_dep_.CoverTab[185634]++
												return s.error(c, "in literal true (expecting 'e')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:512
	// _ = "end of CoverTab[185634]"
}

// stateF is the state after reading `f`.
func stateF(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:516
	_go_fuzz_dep_.CoverTab[185637]++
												if c == 'a' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:517
		_go_fuzz_dep_.CoverTab[185639]++
													s.step = stateFa
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:519
		// _ = "end of CoverTab[185639]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:520
		_go_fuzz_dep_.CoverTab[185640]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:520
		// _ = "end of CoverTab[185640]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:520
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:520
	// _ = "end of CoverTab[185637]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:520
	_go_fuzz_dep_.CoverTab[185638]++
												return s.error(c, "in literal false (expecting 'a')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:521
	// _ = "end of CoverTab[185638]"
}

// stateFa is the state after reading `fa`.
func stateFa(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:525
	_go_fuzz_dep_.CoverTab[185641]++
												if c == 'l' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:526
		_go_fuzz_dep_.CoverTab[185643]++
													s.step = stateFal
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:528
		// _ = "end of CoverTab[185643]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:529
		_go_fuzz_dep_.CoverTab[185644]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:529
		// _ = "end of CoverTab[185644]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:529
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:529
	// _ = "end of CoverTab[185641]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:529
	_go_fuzz_dep_.CoverTab[185642]++
												return s.error(c, "in literal false (expecting 'l')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:530
	// _ = "end of CoverTab[185642]"
}

// stateFal is the state after reading `fal`.
func stateFal(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:534
	_go_fuzz_dep_.CoverTab[185645]++
												if c == 's' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:535
		_go_fuzz_dep_.CoverTab[185647]++
													s.step = stateFals
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:537
		// _ = "end of CoverTab[185647]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:538
		_go_fuzz_dep_.CoverTab[185648]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:538
		// _ = "end of CoverTab[185648]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:538
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:538
	// _ = "end of CoverTab[185645]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:538
	_go_fuzz_dep_.CoverTab[185646]++
												return s.error(c, "in literal false (expecting 's')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:539
	// _ = "end of CoverTab[185646]"
}

// stateFals is the state after reading `fals`.
func stateFals(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:543
	_go_fuzz_dep_.CoverTab[185649]++
												if c == 'e' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:544
		_go_fuzz_dep_.CoverTab[185651]++
													s.step = stateEndValue
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:546
		// _ = "end of CoverTab[185651]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:547
		_go_fuzz_dep_.CoverTab[185652]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:547
		// _ = "end of CoverTab[185652]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:547
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:547
	// _ = "end of CoverTab[185649]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:547
	_go_fuzz_dep_.CoverTab[185650]++
												return s.error(c, "in literal false (expecting 'e')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:548
	// _ = "end of CoverTab[185650]"
}

// stateN is the state after reading `n`.
func stateN(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:552
	_go_fuzz_dep_.CoverTab[185653]++
												if c == 'u' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:553
		_go_fuzz_dep_.CoverTab[185655]++
													s.step = stateNu
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:555
		// _ = "end of CoverTab[185655]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:556
		_go_fuzz_dep_.CoverTab[185656]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:556
		// _ = "end of CoverTab[185656]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:556
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:556
	// _ = "end of CoverTab[185653]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:556
	_go_fuzz_dep_.CoverTab[185654]++
												return s.error(c, "in literal null (expecting 'u')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:557
	// _ = "end of CoverTab[185654]"
}

// stateNu is the state after reading `nu`.
func stateNu(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:561
	_go_fuzz_dep_.CoverTab[185657]++
												if c == 'l' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:562
		_go_fuzz_dep_.CoverTab[185659]++
													s.step = stateNul
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:564
		// _ = "end of CoverTab[185659]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:565
		_go_fuzz_dep_.CoverTab[185660]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:565
		// _ = "end of CoverTab[185660]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:565
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:565
	// _ = "end of CoverTab[185657]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:565
	_go_fuzz_dep_.CoverTab[185658]++
												return s.error(c, "in literal null (expecting 'l')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:566
	// _ = "end of CoverTab[185658]"
}

// stateNul is the state after reading `nul`.
func stateNul(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:570
	_go_fuzz_dep_.CoverTab[185661]++
												if c == 'l' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:571
		_go_fuzz_dep_.CoverTab[185663]++
													s.step = stateEndValue
													return scanContinue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:573
		// _ = "end of CoverTab[185663]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:574
		_go_fuzz_dep_.CoverTab[185664]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:574
		// _ = "end of CoverTab[185664]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:574
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:574
	// _ = "end of CoverTab[185661]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:574
	_go_fuzz_dep_.CoverTab[185662]++
												return s.error(c, "in literal null (expecting 'l')")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:575
	// _ = "end of CoverTab[185662]"
}

// stateError is the state after reaching a syntax error,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:578
// such as after reading `[1}` or `5.1.2`.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:580
func stateError(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:580
	_go_fuzz_dep_.CoverTab[185665]++
												return scanError
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:581
	// _ = "end of CoverTab[185665]"
}

// error records an error and switches to the error state.
func (s *scanner) error(c byte, context string) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:585
	_go_fuzz_dep_.CoverTab[185666]++
												s.step = stateError
												s.err = &SyntaxError{"invalid character " + quoteChar(c) + " " + context, s.bytes}
												return scanError
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:588
	// _ = "end of CoverTab[185666]"
}

// quoteChar formats c as a quoted character literal
func quoteChar(c byte) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:592
	_go_fuzz_dep_.CoverTab[185667]++

												if c == '\'' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:594
		_go_fuzz_dep_.CoverTab[185670]++
													return `'\''`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:595
		// _ = "end of CoverTab[185670]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:596
		_go_fuzz_dep_.CoverTab[185671]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:596
		// _ = "end of CoverTab[185671]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:596
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:596
	// _ = "end of CoverTab[185667]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:596
	_go_fuzz_dep_.CoverTab[185668]++
												if c == '"' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:597
		_go_fuzz_dep_.CoverTab[185672]++
													return `'"'`
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:598
		// _ = "end of CoverTab[185672]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:599
		_go_fuzz_dep_.CoverTab[185673]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:599
		// _ = "end of CoverTab[185673]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:599
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:599
	// _ = "end of CoverTab[185668]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:599
	_go_fuzz_dep_.CoverTab[185669]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:602
	s := strconv.Quote(string(c))
												return "'" + s[1:len(s)-1] + "'"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:603
	// _ = "end of CoverTab[185669]"
}

// undo causes the scanner to return scanCode from the next state transition.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:606
// This gives callers a simple 1-byte undo mechanism.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:608
func (s *scanner) undo(scanCode int) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:608
	_go_fuzz_dep_.CoverTab[185674]++
												if s.redo {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:609
		_go_fuzz_dep_.CoverTab[185676]++
													panic("json: invalid use of scanner")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:610
		// _ = "end of CoverTab[185676]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:611
		_go_fuzz_dep_.CoverTab[185677]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:611
		// _ = "end of CoverTab[185677]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:611
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:611
	// _ = "end of CoverTab[185674]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:611
	_go_fuzz_dep_.CoverTab[185675]++
												s.redoCode = scanCode
												s.redoState = s.step
												s.step = stateRedo
												s.redo = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:615
	// _ = "end of CoverTab[185675]"
}

// stateRedo helps implement the scanner's 1-byte undo.
func stateRedo(s *scanner, c byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:619
	_go_fuzz_dep_.CoverTab[185678]++
												s.redo = false
												s.step = s.redoState
												return s.redoCode
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:622
	// _ = "end of CoverTab[185678]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:623
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/scanner.go:623
var _ = _go_fuzz_dep_.CoverTab
