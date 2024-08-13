// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:5
package json

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:5
)

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/errors"
)

// call specifies which Decoder method was invoked.
type call uint8

const (
	readCall	call	= iota
	peekCall
)

const unexpectedFmt = "unexpected token %s"

// ErrUnexpectedEOF means that EOF was encountered in the middle of the input.
var ErrUnexpectedEOF = errors.New("%v", io.ErrUnexpectedEOF)

// Decoder is a token-based JSON decoder.
type Decoder struct {
	// lastCall is last method called, either readCall or peekCall.
	// Initial value is readCall.
	lastCall	call

	// lastToken contains the last read token.
	lastToken	Token

	// lastErr contains the last read error.
	lastErr	error

	// openStack is a stack containing ObjectOpen and ArrayOpen values. The
	// top of stack represents the object or the array the current value is
	// directly located in.
	openStack	[]Kind

	// orig is used in reporting line and column.
	orig	[]byte
	// in contains the unconsumed input.
	in	[]byte
}

// NewDecoder returns a Decoder to read the given []byte.
func NewDecoder(b []byte) *Decoder {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:54
	_go_fuzz_dep_.CoverTab[65454]++
														return &Decoder{orig: b, in: b}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:55
	// _ = "end of CoverTab[65454]"
}

// Peek looks ahead and returns the next token kind without advancing a read.
func (d *Decoder) Peek() (Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:59
	_go_fuzz_dep_.CoverTab[65455]++
														defer func() { _go_fuzz_dep_.CoverTab[65458]++; d.lastCall = peekCall; // _ = "end of CoverTab[65458]" }()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:60
	// _ = "end of CoverTab[65455]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:60
	_go_fuzz_dep_.CoverTab[65456]++
														if d.lastCall == readCall {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:61
		_go_fuzz_dep_.CoverTab[65459]++
															d.lastToken, d.lastErr = d.Read()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:62
		// _ = "end of CoverTab[65459]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:63
		_go_fuzz_dep_.CoverTab[65460]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:63
		// _ = "end of CoverTab[65460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:63
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:63
	// _ = "end of CoverTab[65456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:63
	_go_fuzz_dep_.CoverTab[65457]++
														return d.lastToken, d.lastErr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:64
	// _ = "end of CoverTab[65457]"
}

// Read returns the next JSON token.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:67
// It will return an error if there is no valid token.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:69
func (d *Decoder) Read() (Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:69
	_go_fuzz_dep_.CoverTab[65461]++
														const scalar = Null | Bool | Number | String

														defer func() { _go_fuzz_dep_.CoverTab[65467]++; d.lastCall = readCall; // _ = "end of CoverTab[65467]" }()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:72
	// _ = "end of CoverTab[65461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:72
	_go_fuzz_dep_.CoverTab[65462]++
														if d.lastCall == peekCall {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:73
		_go_fuzz_dep_.CoverTab[65468]++
															return d.lastToken, d.lastErr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:74
		// _ = "end of CoverTab[65468]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:75
		_go_fuzz_dep_.CoverTab[65469]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:75
		// _ = "end of CoverTab[65469]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:75
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:75
	// _ = "end of CoverTab[65462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:75
	_go_fuzz_dep_.CoverTab[65463]++

														tok, err := d.parseNext()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:78
		_go_fuzz_dep_.CoverTab[65470]++
															return Token{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:79
		// _ = "end of CoverTab[65470]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:80
		_go_fuzz_dep_.CoverTab[65471]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:80
		// _ = "end of CoverTab[65471]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:80
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:80
	// _ = "end of CoverTab[65463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:80
	_go_fuzz_dep_.CoverTab[65464]++

														switch tok.kind {
	case EOF:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:83
		_go_fuzz_dep_.CoverTab[65472]++
															if len(d.openStack) != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:84
			_go_fuzz_dep_.CoverTab[65488]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:84
			return d.lastToken.kind&scalar|ObjectClose|ArrayClose == 0
																// _ = "end of CoverTab[65488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:85
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:85
			_go_fuzz_dep_.CoverTab[65489]++
																return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:86
			// _ = "end of CoverTab[65489]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:87
			_go_fuzz_dep_.CoverTab[65490]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:87
			// _ = "end of CoverTab[65490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:87
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:87
		// _ = "end of CoverTab[65472]"

	case Null:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:89
		_go_fuzz_dep_.CoverTab[65473]++
															if !d.isValueNext() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:90
			_go_fuzz_dep_.CoverTab[65491]++
																return Token{}, d.newSyntaxError(tok.pos, unexpectedFmt, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:91
			// _ = "end of CoverTab[65491]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:92
			_go_fuzz_dep_.CoverTab[65492]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:92
			// _ = "end of CoverTab[65492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:92
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:92
		// _ = "end of CoverTab[65473]"

	case Bool, Number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:94
		_go_fuzz_dep_.CoverTab[65474]++
															if !d.isValueNext() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:95
			_go_fuzz_dep_.CoverTab[65493]++
																return Token{}, d.newSyntaxError(tok.pos, unexpectedFmt, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:96
			// _ = "end of CoverTab[65493]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:97
			_go_fuzz_dep_.CoverTab[65494]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:97
			// _ = "end of CoverTab[65494]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:97
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:97
		// _ = "end of CoverTab[65474]"

	case String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:99
		_go_fuzz_dep_.CoverTab[65475]++
															if d.isValueNext() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:100
			_go_fuzz_dep_.CoverTab[65495]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:101
			// _ = "end of CoverTab[65495]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:102
			_go_fuzz_dep_.CoverTab[65496]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:102
			// _ = "end of CoverTab[65496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:102
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:102
		// _ = "end of CoverTab[65475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:102
		_go_fuzz_dep_.CoverTab[65476]++

															if d.lastToken.kind&(ObjectOpen|comma) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:104
			_go_fuzz_dep_.CoverTab[65497]++
																return Token{}, d.newSyntaxError(tok.pos, unexpectedFmt, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:105
			// _ = "end of CoverTab[65497]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:106
			_go_fuzz_dep_.CoverTab[65498]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:106
			// _ = "end of CoverTab[65498]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:106
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:106
		// _ = "end of CoverTab[65476]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:106
		_go_fuzz_dep_.CoverTab[65477]++
															if len(d.in) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:107
			_go_fuzz_dep_.CoverTab[65499]++
																return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:108
			// _ = "end of CoverTab[65499]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:109
			_go_fuzz_dep_.CoverTab[65500]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:109
			// _ = "end of CoverTab[65500]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:109
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:109
		// _ = "end of CoverTab[65477]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:109
		_go_fuzz_dep_.CoverTab[65478]++
															if c := d.in[0]; c != ':' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:110
			_go_fuzz_dep_.CoverTab[65501]++
																return Token{}, d.newSyntaxError(d.currPos(), `unexpected character %s, missing ":" after field name`, string(c))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:111
			// _ = "end of CoverTab[65501]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:112
			_go_fuzz_dep_.CoverTab[65502]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:112
			// _ = "end of CoverTab[65502]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:112
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:112
		// _ = "end of CoverTab[65478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:112
		_go_fuzz_dep_.CoverTab[65479]++
															tok.kind = Name
															d.consume(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:114
		// _ = "end of CoverTab[65479]"

	case ObjectOpen, ArrayOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:116
		_go_fuzz_dep_.CoverTab[65480]++
															if !d.isValueNext() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:117
			_go_fuzz_dep_.CoverTab[65503]++
																return Token{}, d.newSyntaxError(tok.pos, unexpectedFmt, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:118
			// _ = "end of CoverTab[65503]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:119
			_go_fuzz_dep_.CoverTab[65504]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:119
			// _ = "end of CoverTab[65504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:119
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:119
		// _ = "end of CoverTab[65480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:119
		_go_fuzz_dep_.CoverTab[65481]++
															d.openStack = append(d.openStack, tok.kind)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:120
		// _ = "end of CoverTab[65481]"

	case ObjectClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:122
		_go_fuzz_dep_.CoverTab[65482]++
															if len(d.openStack) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:123
			_go_fuzz_dep_.CoverTab[65505]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:123
			return d.lastToken.kind == comma
																// _ = "end of CoverTab[65505]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:124
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:124
			_go_fuzz_dep_.CoverTab[65506]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:124
			return d.openStack[len(d.openStack)-1] != ObjectOpen
																// _ = "end of CoverTab[65506]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:125
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:125
			_go_fuzz_dep_.CoverTab[65507]++
																return Token{}, d.newSyntaxError(tok.pos, unexpectedFmt, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:126
			// _ = "end of CoverTab[65507]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:127
			_go_fuzz_dep_.CoverTab[65508]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:127
			// _ = "end of CoverTab[65508]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:127
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:127
		// _ = "end of CoverTab[65482]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:127
		_go_fuzz_dep_.CoverTab[65483]++
															d.openStack = d.openStack[:len(d.openStack)-1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:128
		// _ = "end of CoverTab[65483]"

	case ArrayClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:130
		_go_fuzz_dep_.CoverTab[65484]++
															if len(d.openStack) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:131
			_go_fuzz_dep_.CoverTab[65509]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:131
			return d.lastToken.kind == comma
																// _ = "end of CoverTab[65509]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:132
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:132
			_go_fuzz_dep_.CoverTab[65510]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:132
			return d.openStack[len(d.openStack)-1] != ArrayOpen
																// _ = "end of CoverTab[65510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:133
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:133
			_go_fuzz_dep_.CoverTab[65511]++
																return Token{}, d.newSyntaxError(tok.pos, unexpectedFmt, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:134
			// _ = "end of CoverTab[65511]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:135
			_go_fuzz_dep_.CoverTab[65512]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:135
			// _ = "end of CoverTab[65512]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:135
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:135
		// _ = "end of CoverTab[65484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:135
		_go_fuzz_dep_.CoverTab[65485]++
															d.openStack = d.openStack[:len(d.openStack)-1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:136
		// _ = "end of CoverTab[65485]"

	case comma:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:138
		_go_fuzz_dep_.CoverTab[65486]++
															if len(d.openStack) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:139
			_go_fuzz_dep_.CoverTab[65513]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:139
			return d.lastToken.kind&(scalar|ObjectClose|ArrayClose) == 0
																// _ = "end of CoverTab[65513]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:140
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:140
			_go_fuzz_dep_.CoverTab[65514]++
																return Token{}, d.newSyntaxError(tok.pos, unexpectedFmt, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:141
			// _ = "end of CoverTab[65514]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:142
			_go_fuzz_dep_.CoverTab[65515]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:142
			// _ = "end of CoverTab[65515]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:142
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:142
		// _ = "end of CoverTab[65486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:142
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:142
		_go_fuzz_dep_.CoverTab[65487]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:142
		// _ = "end of CoverTab[65487]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:143
	// _ = "end of CoverTab[65464]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:143
	_go_fuzz_dep_.CoverTab[65465]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:146
	d.lastToken = tok

	if d.lastToken.kind == comma {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:148
		_go_fuzz_dep_.CoverTab[65516]++
															return d.Read()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:149
		// _ = "end of CoverTab[65516]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:150
		_go_fuzz_dep_.CoverTab[65517]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:150
		// _ = "end of CoverTab[65517]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:150
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:150
	// _ = "end of CoverTab[65465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:150
	_go_fuzz_dep_.CoverTab[65466]++
														return tok, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:151
	// _ = "end of CoverTab[65466]"
}

// Any sequence that looks like a non-delimiter (for error reporting).
var errRegexp = regexp.MustCompile(`^([-+._a-zA-Z0-9]{1,32}|.)`)

// parseNext parses for the next JSON token. It returns a Token object for
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:157
// different types, except for Name. It does not handle whether the next token
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:157
// is in a valid sequence or not.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:160
func (d *Decoder) parseNext() (Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:160
	_go_fuzz_dep_.CoverTab[65518]++

														d.consume(0)

														in := d.in
														if len(in) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:165
		_go_fuzz_dep_.CoverTab[65521]++
															return d.consumeToken(EOF, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:166
		// _ = "end of CoverTab[65521]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:167
		_go_fuzz_dep_.CoverTab[65522]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:167
		// _ = "end of CoverTab[65522]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:167
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:167
	// _ = "end of CoverTab[65518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:167
	_go_fuzz_dep_.CoverTab[65519]++

														switch in[0] {
	case 'n':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:170
		_go_fuzz_dep_.CoverTab[65523]++
															if n := matchWithDelim("null", in); n != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:171
			_go_fuzz_dep_.CoverTab[65535]++
																return d.consumeToken(Null, n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:172
			// _ = "end of CoverTab[65535]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:173
			_go_fuzz_dep_.CoverTab[65536]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:173
			// _ = "end of CoverTab[65536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:173
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:173
		// _ = "end of CoverTab[65523]"

	case 't':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:175
		_go_fuzz_dep_.CoverTab[65524]++
															if n := matchWithDelim("true", in); n != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:176
			_go_fuzz_dep_.CoverTab[65537]++
																return d.consumeBoolToken(true, n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:177
			// _ = "end of CoverTab[65537]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:178
			_go_fuzz_dep_.CoverTab[65538]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:178
			// _ = "end of CoverTab[65538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:178
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:178
		// _ = "end of CoverTab[65524]"

	case 'f':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:180
		_go_fuzz_dep_.CoverTab[65525]++
															if n := matchWithDelim("false", in); n != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:181
			_go_fuzz_dep_.CoverTab[65539]++
																return d.consumeBoolToken(false, n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:182
			// _ = "end of CoverTab[65539]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:183
			_go_fuzz_dep_.CoverTab[65540]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:183
			// _ = "end of CoverTab[65540]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:183
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:183
		// _ = "end of CoverTab[65525]"

	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:185
		_go_fuzz_dep_.CoverTab[65526]++
															if n, ok := parseNumber(in); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:186
			_go_fuzz_dep_.CoverTab[65541]++
																return d.consumeToken(Number, n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:187
			// _ = "end of CoverTab[65541]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:188
			_go_fuzz_dep_.CoverTab[65542]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:188
			// _ = "end of CoverTab[65542]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:188
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:188
		// _ = "end of CoverTab[65526]"

	case '"':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:190
		_go_fuzz_dep_.CoverTab[65527]++
															s, n, err := d.parseString(in)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:192
			_go_fuzz_dep_.CoverTab[65543]++
																return Token{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:193
			// _ = "end of CoverTab[65543]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:194
			_go_fuzz_dep_.CoverTab[65544]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:194
			// _ = "end of CoverTab[65544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:194
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:194
		// _ = "end of CoverTab[65527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:194
		_go_fuzz_dep_.CoverTab[65528]++
															return d.consumeStringToken(s, n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:195
		// _ = "end of CoverTab[65528]"

	case '{':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:197
		_go_fuzz_dep_.CoverTab[65529]++
															return d.consumeToken(ObjectOpen, 1), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:198
		// _ = "end of CoverTab[65529]"

	case '}':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:200
		_go_fuzz_dep_.CoverTab[65530]++
															return d.consumeToken(ObjectClose, 1), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:201
		// _ = "end of CoverTab[65530]"

	case '[':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:203
		_go_fuzz_dep_.CoverTab[65531]++
															return d.consumeToken(ArrayOpen, 1), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:204
		// _ = "end of CoverTab[65531]"

	case ']':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:206
		_go_fuzz_dep_.CoverTab[65532]++
															return d.consumeToken(ArrayClose, 1), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:207
		// _ = "end of CoverTab[65532]"

	case ',':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:209
		_go_fuzz_dep_.CoverTab[65533]++
															return d.consumeToken(comma, 1), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:210
		// _ = "end of CoverTab[65533]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:210
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:210
		_go_fuzz_dep_.CoverTab[65534]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:210
		// _ = "end of CoverTab[65534]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:211
	// _ = "end of CoverTab[65519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:211
	_go_fuzz_dep_.CoverTab[65520]++
														return Token{}, d.newSyntaxError(d.currPos(), "invalid value %s", errRegexp.Find(in))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:212
	// _ = "end of CoverTab[65520]"
}

// newSyntaxError returns an error with line and column information useful for
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:215
// syntax errors.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:217
func (d *Decoder) newSyntaxError(pos int, f string, x ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:217
	_go_fuzz_dep_.CoverTab[65545]++
														e := errors.New(f, x...)
														line, column := d.Position(pos)
														return errors.New("syntax error (line %d:%d): %v", line, column, e)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:220
	// _ = "end of CoverTab[65545]"
}

// Position returns line and column number of given index of the original input.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:223
// It will panic if index is out of range.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:225
func (d *Decoder) Position(idx int) (line int, column int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:225
	_go_fuzz_dep_.CoverTab[65546]++
														b := d.orig[:idx]
														line = bytes.Count(b, []byte("\n")) + 1
														if i := bytes.LastIndexByte(b, '\n'); i >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:228
		_go_fuzz_dep_.CoverTab[65548]++
															b = b[i+1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:229
		// _ = "end of CoverTab[65548]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:230
		_go_fuzz_dep_.CoverTab[65549]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:230
		// _ = "end of CoverTab[65549]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:230
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:230
	// _ = "end of CoverTab[65546]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:230
	_go_fuzz_dep_.CoverTab[65547]++
														column = utf8.RuneCount(b) + 1
														return line, column
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:232
	// _ = "end of CoverTab[65547]"
}

// currPos returns the current index position of d.in from d.orig.
func (d *Decoder) currPos() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:236
	_go_fuzz_dep_.CoverTab[65550]++
														return len(d.orig) - len(d.in)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:237
	// _ = "end of CoverTab[65550]"
}

// matchWithDelim matches s with the input b and verifies that the match
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:240
// terminates with a delimiter of some form (e.g., r"[^-+_.a-zA-Z0-9]").
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:240
// As a special case, EOF is considered a delimiter. It returns the length of s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:240
// if there is a match, else 0.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:244
func matchWithDelim(s string, b []byte) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:244
	_go_fuzz_dep_.CoverTab[65551]++
														if !bytes.HasPrefix(b, []byte(s)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:245
		_go_fuzz_dep_.CoverTab[65554]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:246
		// _ = "end of CoverTab[65554]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:247
		_go_fuzz_dep_.CoverTab[65555]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:247
		// _ = "end of CoverTab[65555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:247
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:247
	// _ = "end of CoverTab[65551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:247
	_go_fuzz_dep_.CoverTab[65552]++

														n := len(s)
														if n < len(b) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:250
		_go_fuzz_dep_.CoverTab[65556]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:250
		return isNotDelim(b[n])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:250
		// _ = "end of CoverTab[65556]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:250
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:250
		_go_fuzz_dep_.CoverTab[65557]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:251
		// _ = "end of CoverTab[65557]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:252
		_go_fuzz_dep_.CoverTab[65558]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:252
		// _ = "end of CoverTab[65558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:252
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:252
	// _ = "end of CoverTab[65552]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:252
	_go_fuzz_dep_.CoverTab[65553]++
														return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:253
	// _ = "end of CoverTab[65553]"
}

// isNotDelim returns true if given byte is a not delimiter character.
func isNotDelim(c byte) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:257
	_go_fuzz_dep_.CoverTab[65559]++
														return (c == '-' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		_go_fuzz_dep_.CoverTab[65560]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		return c == '+'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		// _ = "end of CoverTab[65560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		_go_fuzz_dep_.CoverTab[65561]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		return c == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		// _ = "end of CoverTab[65561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		_go_fuzz_dep_.CoverTab[65562]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		return c == '_'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		// _ = "end of CoverTab[65562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		_go_fuzz_dep_.CoverTab[65563]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:258
		return ('a' <= c && func() bool {
																_go_fuzz_dep_.CoverTab[65564]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:259
			return c <= 'z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:259
			// _ = "end of CoverTab[65564]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:259
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:259
		// _ = "end of CoverTab[65563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:259
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:259
		_go_fuzz_dep_.CoverTab[65565]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:259
		return ('A' <= c && func() bool {
																_go_fuzz_dep_.CoverTab[65566]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:260
			return c <= 'Z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:260
			// _ = "end of CoverTab[65566]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:260
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:260
		// _ = "end of CoverTab[65565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:260
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:260
		_go_fuzz_dep_.CoverTab[65567]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:260
		return ('0' <= c && func() bool {
																_go_fuzz_dep_.CoverTab[65568]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:261
			return c <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:261
			// _ = "end of CoverTab[65568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:261
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:261
		// _ = "end of CoverTab[65567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:261
	}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:261
	// _ = "end of CoverTab[65559]"
}

// consume consumes n bytes of input and any subsequent whitespace.
func (d *Decoder) consume(n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:265
	_go_fuzz_dep_.CoverTab[65569]++
														d.in = d.in[n:]
														for len(d.in) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:267
		_go_fuzz_dep_.CoverTab[65570]++
															switch d.in[0] {
		case ' ', '\n', '\r', '\t':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:269
			_go_fuzz_dep_.CoverTab[65571]++
																d.in = d.in[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:270
			// _ = "end of CoverTab[65571]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:271
			_go_fuzz_dep_.CoverTab[65572]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:272
			// _ = "end of CoverTab[65572]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:273
		// _ = "end of CoverTab[65570]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:274
	// _ = "end of CoverTab[65569]"
}

// isValueNext returns true if next type should be a JSON value: Null,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:277
// Number, String or Bool.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:279
func (d *Decoder) isValueNext() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:279
	_go_fuzz_dep_.CoverTab[65573]++
														if len(d.openStack) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:280
		_go_fuzz_dep_.CoverTab[65576]++
															return d.lastToken.kind == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:281
		// _ = "end of CoverTab[65576]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:282
		_go_fuzz_dep_.CoverTab[65577]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:282
		// _ = "end of CoverTab[65577]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:282
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:282
	// _ = "end of CoverTab[65573]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:282
	_go_fuzz_dep_.CoverTab[65574]++

														start := d.openStack[len(d.openStack)-1]
														switch start {
	case ObjectOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:286
		_go_fuzz_dep_.CoverTab[65578]++
															return d.lastToken.kind&Name != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:287
		// _ = "end of CoverTab[65578]"
	case ArrayOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:288
		_go_fuzz_dep_.CoverTab[65579]++
															return d.lastToken.kind&(ArrayOpen|comma) != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:289
		// _ = "end of CoverTab[65579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:289
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:289
		_go_fuzz_dep_.CoverTab[65580]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:289
		// _ = "end of CoverTab[65580]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:290
	// _ = "end of CoverTab[65574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:290
	_go_fuzz_dep_.CoverTab[65575]++
														panic(fmt.Sprintf(
		"unreachable logic in Decoder.isValueNext, lastToken.kind: %v, openStack: %v",
		d.lastToken.kind, start))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:293
	// _ = "end of CoverTab[65575]"
}

// consumeToken constructs a Token for given Kind with raw value derived from
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:296
// current d.in and given size, and consumes the given size-lenght of it.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:298
func (d *Decoder) consumeToken(kind Kind, size int) Token {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:298
	_go_fuzz_dep_.CoverTab[65581]++
														tok := Token{
		kind:	kind,
		raw:	d.in[:size],
		pos:	len(d.orig) - len(d.in),
	}
														d.consume(size)
														return tok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:305
	// _ = "end of CoverTab[65581]"
}

// consumeBoolToken constructs a Token for a Bool kind with raw value derived from
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:308
// current d.in and given size.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:310
func (d *Decoder) consumeBoolToken(b bool, size int) Token {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:310
	_go_fuzz_dep_.CoverTab[65582]++
														tok := Token{
		kind:	Bool,
		raw:	d.in[:size],
		pos:	len(d.orig) - len(d.in),
		boo:	b,
	}
														d.consume(size)
														return tok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:318
	// _ = "end of CoverTab[65582]"
}

// consumeStringToken constructs a Token for a String kind with raw value derived
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:321
// from current d.in and given size.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:323
func (d *Decoder) consumeStringToken(s string, size int) Token {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:323
	_go_fuzz_dep_.CoverTab[65583]++
														tok := Token{
		kind:	String,
		raw:	d.in[:size],
		pos:	len(d.orig) - len(d.in),
		str:	s,
	}
														d.consume(size)
														return tok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:331
	// _ = "end of CoverTab[65583]"
}

// Clone returns a copy of the Decoder for use in reading ahead the next JSON
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:334
// object, array or other values without affecting current Decoder.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:336
func (d *Decoder) Clone() *Decoder {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:336
	_go_fuzz_dep_.CoverTab[65584]++
														ret := *d
														ret.openStack = append([]Kind(nil), ret.openStack...)
														return &ret
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:339
	// _ = "end of CoverTab[65584]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:340
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/json/decode.go:340
var _ = _go_fuzz_dep_.CoverTab
