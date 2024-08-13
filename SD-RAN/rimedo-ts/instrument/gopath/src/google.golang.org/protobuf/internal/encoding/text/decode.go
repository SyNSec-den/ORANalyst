// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:5
package text

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:5
)

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/errors"
)

// Decoder is a token-based textproto decoder.
type Decoder struct {
	// lastCall is last method called, either readCall or peekCall.
	// Initial value is readCall.
	lastCall	call

	// lastToken contains the last read token.
	lastToken	Token

	// lastErr contains the last read error.
	lastErr	error

	// openStack is a stack containing the byte characters for MessageOpen and
	// ListOpen kinds. The top of stack represents the message or the list that
	// the current token is nested in. An empty stack means the current token is
	// at the top level message. The characters '{' and '<' both represent the
	// MessageOpen kind.
	openStack	[]byte

	// orig is used in reporting line and column.
	orig	[]byte
	// in contains the unconsumed input.
	in	[]byte
}

// NewDecoder returns a Decoder to read the given []byte.
func NewDecoder(b []byte) *Decoder {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:43
	_go_fuzz_dep_.CoverTab[49403]++
														return &Decoder{orig: b, in: b}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:44
	// _ = "end of CoverTab[49403]"
}

// ErrUnexpectedEOF means that EOF was encountered in the middle of the input.
var ErrUnexpectedEOF = errors.New("%v", io.ErrUnexpectedEOF)

// call specifies which Decoder method was invoked.
type call uint8

const (
	readCall	call	= iota
	peekCall
)

// Peek looks ahead and returns the next token and error without advancing a read.
func (d *Decoder) Peek() (Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:59
	_go_fuzz_dep_.CoverTab[49404]++
														defer func() { _go_fuzz_dep_.CoverTab[49407]++; d.lastCall = peekCall; // _ = "end of CoverTab[49407]" }()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:60
	// _ = "end of CoverTab[49404]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:60
	_go_fuzz_dep_.CoverTab[49405]++
														if d.lastCall == readCall {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:61
		_go_fuzz_dep_.CoverTab[49408]++
															d.lastToken, d.lastErr = d.Read()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:62
		// _ = "end of CoverTab[49408]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:63
		_go_fuzz_dep_.CoverTab[49409]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:63
		// _ = "end of CoverTab[49409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:63
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:63
	// _ = "end of CoverTab[49405]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:63
	_go_fuzz_dep_.CoverTab[49406]++
														return d.lastToken, d.lastErr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:64
	// _ = "end of CoverTab[49406]"
}

// Read returns the next token.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:67
// It will return an error if there is no valid token.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:69
func (d *Decoder) Read() (Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:69
	_go_fuzz_dep_.CoverTab[49410]++
														defer func() { _go_fuzz_dep_.CoverTab[49415]++; d.lastCall = readCall; // _ = "end of CoverTab[49415]" }()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:70
	// _ = "end of CoverTab[49410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:70
	_go_fuzz_dep_.CoverTab[49411]++
														if d.lastCall == peekCall {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:71
		_go_fuzz_dep_.CoverTab[49416]++
															return d.lastToken, d.lastErr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:72
		// _ = "end of CoverTab[49416]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:73
		_go_fuzz_dep_.CoverTab[49417]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:73
		// _ = "end of CoverTab[49417]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:73
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:73
	// _ = "end of CoverTab[49411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:73
	_go_fuzz_dep_.CoverTab[49412]++

														tok, err := d.parseNext(d.lastToken.kind)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:76
		_go_fuzz_dep_.CoverTab[49418]++
															return Token{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:77
		// _ = "end of CoverTab[49418]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:78
		_go_fuzz_dep_.CoverTab[49419]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:78
		// _ = "end of CoverTab[49419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:78
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:78
	// _ = "end of CoverTab[49412]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:78
	_go_fuzz_dep_.CoverTab[49413]++

														switch tok.kind {
	case comma, semicolon:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:81
		_go_fuzz_dep_.CoverTab[49420]++
															tok, err = d.parseNext(tok.kind)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:83
			_go_fuzz_dep_.CoverTab[49422]++
																return Token{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:84
			// _ = "end of CoverTab[49422]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:85
			_go_fuzz_dep_.CoverTab[49423]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:85
			// _ = "end of CoverTab[49423]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:85
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:85
		// _ = "end of CoverTab[49420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:85
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:85
		_go_fuzz_dep_.CoverTab[49421]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:85
		// _ = "end of CoverTab[49421]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:86
	// _ = "end of CoverTab[49413]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:86
	_go_fuzz_dep_.CoverTab[49414]++
														d.lastToken = tok
														return tok, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:88
	// _ = "end of CoverTab[49414]"
}

const (
	mismatchedFmt	= "mismatched close character %q"
	unexpectedFmt	= "unexpected character %q"
)

// parseNext parses the next Token based on given last kind.
func (d *Decoder) parseNext(lastKind Kind) (Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:97
	_go_fuzz_dep_.CoverTab[49424]++

														d.consume(0)
														isEOF := false
														if len(d.in) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:101
		_go_fuzz_dep_.CoverTab[49427]++
															isEOF = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:102
		// _ = "end of CoverTab[49427]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:103
		_go_fuzz_dep_.CoverTab[49428]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:103
		// _ = "end of CoverTab[49428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:103
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:103
	// _ = "end of CoverTab[49424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:103
	_go_fuzz_dep_.CoverTab[49425]++

														switch lastKind {
	case EOF:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:106
		_go_fuzz_dep_.CoverTab[49429]++
															return d.consumeToken(EOF, 0, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:107
		// _ = "end of CoverTab[49429]"

	case bof:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:109
		_go_fuzz_dep_.CoverTab[49430]++

															if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:111
			_go_fuzz_dep_.CoverTab[49443]++
																return d.consumeToken(EOF, 0, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:112
			// _ = "end of CoverTab[49443]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:113
			_go_fuzz_dep_.CoverTab[49444]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:113
			// _ = "end of CoverTab[49444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:113
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:113
		// _ = "end of CoverTab[49430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:113
		_go_fuzz_dep_.CoverTab[49431]++
															return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:114
		// _ = "end of CoverTab[49431]"

	case Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:116
		_go_fuzz_dep_.CoverTab[49432]++

															if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:118
			_go_fuzz_dep_.CoverTab[49445]++
																return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:119
			// _ = "end of CoverTab[49445]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:120
			_go_fuzz_dep_.CoverTab[49446]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:120
			// _ = "end of CoverTab[49446]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:120
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:120
		// _ = "end of CoverTab[49432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:120
		_go_fuzz_dep_.CoverTab[49433]++
															switch ch := d.in[0]; ch {
		case '{', '<':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:122
			_go_fuzz_dep_.CoverTab[49447]++
																d.pushOpenStack(ch)
																return d.consumeToken(MessageOpen, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:124
			// _ = "end of CoverTab[49447]"
		case '[':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:125
			_go_fuzz_dep_.CoverTab[49448]++
																d.pushOpenStack(ch)
																return d.consumeToken(ListOpen, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:127
			// _ = "end of CoverTab[49448]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:128
			_go_fuzz_dep_.CoverTab[49449]++
																return d.parseScalar()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:129
			// _ = "end of CoverTab[49449]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:130
		// _ = "end of CoverTab[49433]"

	case Scalar:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:132
		_go_fuzz_dep_.CoverTab[49434]++
															openKind, closeCh := d.currentOpenKind()
															switch openKind {
		case bof:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:135
			_go_fuzz_dep_.CoverTab[49450]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:138
			if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:138
				_go_fuzz_dep_.CoverTab[49457]++
																	return d.consumeToken(EOF, 0, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:139
				// _ = "end of CoverTab[49457]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:140
				_go_fuzz_dep_.CoverTab[49458]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:140
				// _ = "end of CoverTab[49458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:140
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:140
			// _ = "end of CoverTab[49450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:140
			_go_fuzz_dep_.CoverTab[49451]++
																switch d.in[0] {
			case ',':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:142
				_go_fuzz_dep_.CoverTab[49459]++
																	return d.consumeToken(comma, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:143
				// _ = "end of CoverTab[49459]"
			case ';':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:144
				_go_fuzz_dep_.CoverTab[49460]++
																	return d.consumeToken(semicolon, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:145
				// _ = "end of CoverTab[49460]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:146
				_go_fuzz_dep_.CoverTab[49461]++
																	return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:147
				// _ = "end of CoverTab[49461]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:148
			// _ = "end of CoverTab[49451]"

		case MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:150
			_go_fuzz_dep_.CoverTab[49452]++

																if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:152
				_go_fuzz_dep_.CoverTab[49462]++
																	return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:153
				// _ = "end of CoverTab[49462]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:154
				_go_fuzz_dep_.CoverTab[49463]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:154
				// _ = "end of CoverTab[49463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:154
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:154
			// _ = "end of CoverTab[49452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:154
			_go_fuzz_dep_.CoverTab[49453]++
																switch ch := d.in[0]; ch {
			case closeCh:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:156
				_go_fuzz_dep_.CoverTab[49464]++
																	d.popOpenStack()
																	return d.consumeToken(MessageClose, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:158
				// _ = "end of CoverTab[49464]"
			case otherCloseChar[closeCh]:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:159
				_go_fuzz_dep_.CoverTab[49465]++
																	return Token{}, d.newSyntaxError(mismatchedFmt, ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:160
				// _ = "end of CoverTab[49465]"
			case ',':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:161
				_go_fuzz_dep_.CoverTab[49466]++
																	return d.consumeToken(comma, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:162
				// _ = "end of CoverTab[49466]"
			case ';':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:163
				_go_fuzz_dep_.CoverTab[49467]++
																	return d.consumeToken(semicolon, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:164
				// _ = "end of CoverTab[49467]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:165
				_go_fuzz_dep_.CoverTab[49468]++
																	return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:166
				// _ = "end of CoverTab[49468]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:167
			// _ = "end of CoverTab[49453]"

		case ListOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:169
			_go_fuzz_dep_.CoverTab[49454]++

																if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:171
				_go_fuzz_dep_.CoverTab[49469]++
																	return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:172
				// _ = "end of CoverTab[49469]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:173
				_go_fuzz_dep_.CoverTab[49470]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:173
				// _ = "end of CoverTab[49470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:173
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:173
			// _ = "end of CoverTab[49454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:173
			_go_fuzz_dep_.CoverTab[49455]++
																switch ch := d.in[0]; ch {
			case ']':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:175
				_go_fuzz_dep_.CoverTab[49471]++
																	d.popOpenStack()
																	return d.consumeToken(ListClose, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:177
				// _ = "end of CoverTab[49471]"
			case ',':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:178
				_go_fuzz_dep_.CoverTab[49472]++
																	return d.consumeToken(comma, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:179
				// _ = "end of CoverTab[49472]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:180
				_go_fuzz_dep_.CoverTab[49473]++
																	return Token{}, d.newSyntaxError(unexpectedFmt, ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:181
				// _ = "end of CoverTab[49473]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:182
			// _ = "end of CoverTab[49455]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:182
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:182
			_go_fuzz_dep_.CoverTab[49456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:182
			// _ = "end of CoverTab[49456]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:183
		// _ = "end of CoverTab[49434]"

	case MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:185
		_go_fuzz_dep_.CoverTab[49435]++

															if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:187
			_go_fuzz_dep_.CoverTab[49474]++
																return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:188
			// _ = "end of CoverTab[49474]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:189
			_go_fuzz_dep_.CoverTab[49475]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:189
			// _ = "end of CoverTab[49475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:189
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:189
		// _ = "end of CoverTab[49435]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:189
		_go_fuzz_dep_.CoverTab[49436]++
															_, closeCh := d.currentOpenKind()
															switch ch := d.in[0]; ch {
		case closeCh:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:192
			_go_fuzz_dep_.CoverTab[49476]++
																d.popOpenStack()
																return d.consumeToken(MessageClose, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:194
			// _ = "end of CoverTab[49476]"
		case otherCloseChar[closeCh]:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:195
			_go_fuzz_dep_.CoverTab[49477]++
																return Token{}, d.newSyntaxError(mismatchedFmt, ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:196
			// _ = "end of CoverTab[49477]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:197
			_go_fuzz_dep_.CoverTab[49478]++
																return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:198
			// _ = "end of CoverTab[49478]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:199
		// _ = "end of CoverTab[49436]"

	case MessageClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:201
		_go_fuzz_dep_.CoverTab[49437]++
															openKind, closeCh := d.currentOpenKind()
															switch openKind {
		case bof:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:204
			_go_fuzz_dep_.CoverTab[49479]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:207
			if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:207
				_go_fuzz_dep_.CoverTab[49486]++
																	return d.consumeToken(EOF, 0, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:208
				// _ = "end of CoverTab[49486]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:209
				_go_fuzz_dep_.CoverTab[49487]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:209
				// _ = "end of CoverTab[49487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:209
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:209
			// _ = "end of CoverTab[49479]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:209
			_go_fuzz_dep_.CoverTab[49480]++
																switch ch := d.in[0]; ch {
			case ',':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:211
				_go_fuzz_dep_.CoverTab[49488]++
																	return d.consumeToken(comma, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:212
				// _ = "end of CoverTab[49488]"
			case ';':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:213
				_go_fuzz_dep_.CoverTab[49489]++
																	return d.consumeToken(semicolon, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:214
				// _ = "end of CoverTab[49489]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:215
				_go_fuzz_dep_.CoverTab[49490]++
																	return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:216
				// _ = "end of CoverTab[49490]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:217
			// _ = "end of CoverTab[49480]"

		case MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:219
			_go_fuzz_dep_.CoverTab[49481]++

																if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:221
				_go_fuzz_dep_.CoverTab[49491]++
																	return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:222
				// _ = "end of CoverTab[49491]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:223
				_go_fuzz_dep_.CoverTab[49492]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:223
				// _ = "end of CoverTab[49492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:223
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:223
			// _ = "end of CoverTab[49481]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:223
			_go_fuzz_dep_.CoverTab[49482]++
																switch ch := d.in[0]; ch {
			case closeCh:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:225
				_go_fuzz_dep_.CoverTab[49493]++
																	d.popOpenStack()
																	return d.consumeToken(MessageClose, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:227
				// _ = "end of CoverTab[49493]"
			case otherCloseChar[closeCh]:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:228
				_go_fuzz_dep_.CoverTab[49494]++
																	return Token{}, d.newSyntaxError(mismatchedFmt, ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:229
				// _ = "end of CoverTab[49494]"
			case ',':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:230
				_go_fuzz_dep_.CoverTab[49495]++
																	return d.consumeToken(comma, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:231
				// _ = "end of CoverTab[49495]"
			case ';':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:232
				_go_fuzz_dep_.CoverTab[49496]++
																	return d.consumeToken(semicolon, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:233
				// _ = "end of CoverTab[49496]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:234
				_go_fuzz_dep_.CoverTab[49497]++
																	return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:235
				// _ = "end of CoverTab[49497]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:236
			// _ = "end of CoverTab[49482]"

		case ListOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:238
			_go_fuzz_dep_.CoverTab[49483]++

																if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:240
				_go_fuzz_dep_.CoverTab[49498]++
																	return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:241
				// _ = "end of CoverTab[49498]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:242
				_go_fuzz_dep_.CoverTab[49499]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:242
				// _ = "end of CoverTab[49499]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:242
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:242
			// _ = "end of CoverTab[49483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:242
			_go_fuzz_dep_.CoverTab[49484]++
																switch ch := d.in[0]; ch {
			case closeCh:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:244
				_go_fuzz_dep_.CoverTab[49500]++
																	d.popOpenStack()
																	return d.consumeToken(ListClose, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:246
				// _ = "end of CoverTab[49500]"
			case ',':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:247
				_go_fuzz_dep_.CoverTab[49501]++
																	return d.consumeToken(comma, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:248
				// _ = "end of CoverTab[49501]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:249
				_go_fuzz_dep_.CoverTab[49502]++
																	return Token{}, d.newSyntaxError(unexpectedFmt, ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:250
				// _ = "end of CoverTab[49502]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:251
			// _ = "end of CoverTab[49484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:251
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:251
			_go_fuzz_dep_.CoverTab[49485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:251
			// _ = "end of CoverTab[49485]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:252
		// _ = "end of CoverTab[49437]"

	case ListOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:254
		_go_fuzz_dep_.CoverTab[49438]++

															if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:256
			_go_fuzz_dep_.CoverTab[49503]++
																return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:257
			// _ = "end of CoverTab[49503]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:258
			_go_fuzz_dep_.CoverTab[49504]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:258
			// _ = "end of CoverTab[49504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:258
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:258
		// _ = "end of CoverTab[49438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:258
		_go_fuzz_dep_.CoverTab[49439]++
															switch ch := d.in[0]; ch {
		case ']':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:260
			_go_fuzz_dep_.CoverTab[49505]++
																d.popOpenStack()
																return d.consumeToken(ListClose, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:262
			// _ = "end of CoverTab[49505]"
		case '{', '<':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:263
			_go_fuzz_dep_.CoverTab[49506]++
																d.pushOpenStack(ch)
																return d.consumeToken(MessageOpen, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:265
			// _ = "end of CoverTab[49506]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:266
			_go_fuzz_dep_.CoverTab[49507]++
																return d.parseScalar()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:267
			// _ = "end of CoverTab[49507]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:268
		// _ = "end of CoverTab[49439]"

	case ListClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:270
		_go_fuzz_dep_.CoverTab[49440]++
															openKind, closeCh := d.currentOpenKind()
															switch openKind {
		case bof:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:273
			_go_fuzz_dep_.CoverTab[49508]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:276
			if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:276
				_go_fuzz_dep_.CoverTab[49513]++
																	return d.consumeToken(EOF, 0, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:277
				// _ = "end of CoverTab[49513]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:278
				_go_fuzz_dep_.CoverTab[49514]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:278
				// _ = "end of CoverTab[49514]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:278
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:278
			// _ = "end of CoverTab[49508]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:278
			_go_fuzz_dep_.CoverTab[49509]++
																switch ch := d.in[0]; ch {
			case ',':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:280
				_go_fuzz_dep_.CoverTab[49515]++
																	return d.consumeToken(comma, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:281
				// _ = "end of CoverTab[49515]"
			case ';':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:282
				_go_fuzz_dep_.CoverTab[49516]++
																	return d.consumeToken(semicolon, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:283
				// _ = "end of CoverTab[49516]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:284
				_go_fuzz_dep_.CoverTab[49517]++
																	return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:285
				// _ = "end of CoverTab[49517]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:286
			// _ = "end of CoverTab[49509]"

		case MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:288
			_go_fuzz_dep_.CoverTab[49510]++

																if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:290
				_go_fuzz_dep_.CoverTab[49518]++
																	return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:291
				// _ = "end of CoverTab[49518]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:292
				_go_fuzz_dep_.CoverTab[49519]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:292
				// _ = "end of CoverTab[49519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:292
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:292
			// _ = "end of CoverTab[49510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:292
			_go_fuzz_dep_.CoverTab[49511]++
																switch ch := d.in[0]; ch {
			case closeCh:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:294
				_go_fuzz_dep_.CoverTab[49520]++
																	d.popOpenStack()
																	return d.consumeToken(MessageClose, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:296
				// _ = "end of CoverTab[49520]"
			case otherCloseChar[closeCh]:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:297
				_go_fuzz_dep_.CoverTab[49521]++
																	return Token{}, d.newSyntaxError(mismatchedFmt, ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:298
				// _ = "end of CoverTab[49521]"
			case ',':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:299
				_go_fuzz_dep_.CoverTab[49522]++
																	return d.consumeToken(comma, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:300
				// _ = "end of CoverTab[49522]"
			case ';':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:301
				_go_fuzz_dep_.CoverTab[49523]++
																	return d.consumeToken(semicolon, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:302
				// _ = "end of CoverTab[49523]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:303
				_go_fuzz_dep_.CoverTab[49524]++
																	return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:304
				// _ = "end of CoverTab[49524]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:305
			// _ = "end of CoverTab[49511]"

		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:307
			_go_fuzz_dep_.CoverTab[49512]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:307
			// _ = "end of CoverTab[49512]"

		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:309
		// _ = "end of CoverTab[49440]"

	case comma, semicolon:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:311
		_go_fuzz_dep_.CoverTab[49441]++
															openKind, closeCh := d.currentOpenKind()
															switch openKind {
		case bof:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:314
			_go_fuzz_dep_.CoverTab[49525]++

																if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:316
				_go_fuzz_dep_.CoverTab[49533]++
																	return d.consumeToken(EOF, 0, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:317
				// _ = "end of CoverTab[49533]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:318
				_go_fuzz_dep_.CoverTab[49534]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:318
				// _ = "end of CoverTab[49534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:318
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:318
			// _ = "end of CoverTab[49525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:318
			_go_fuzz_dep_.CoverTab[49526]++
																return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:319
			// _ = "end of CoverTab[49526]"

		case MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:321
			_go_fuzz_dep_.CoverTab[49527]++

																if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:323
				_go_fuzz_dep_.CoverTab[49535]++
																	return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:324
				// _ = "end of CoverTab[49535]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:325
				_go_fuzz_dep_.CoverTab[49536]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:325
				// _ = "end of CoverTab[49536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:325
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:325
			// _ = "end of CoverTab[49527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:325
			_go_fuzz_dep_.CoverTab[49528]++
																switch ch := d.in[0]; ch {
			case closeCh:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:327
				_go_fuzz_dep_.CoverTab[49537]++
																	d.popOpenStack()
																	return d.consumeToken(MessageClose, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:329
				// _ = "end of CoverTab[49537]"
			case otherCloseChar[closeCh]:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:330
				_go_fuzz_dep_.CoverTab[49538]++
																	return Token{}, d.newSyntaxError(mismatchedFmt, ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:331
				// _ = "end of CoverTab[49538]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:332
				_go_fuzz_dep_.CoverTab[49539]++
																	return d.parseFieldName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:333
				// _ = "end of CoverTab[49539]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:334
			// _ = "end of CoverTab[49528]"

		case ListOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:336
			_go_fuzz_dep_.CoverTab[49529]++
																if lastKind == semicolon {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:337
				_go_fuzz_dep_.CoverTab[49540]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:341
				break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:341
				// _ = "end of CoverTab[49540]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:342
				_go_fuzz_dep_.CoverTab[49541]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:342
				// _ = "end of CoverTab[49541]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:342
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:342
			// _ = "end of CoverTab[49529]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:342
			_go_fuzz_dep_.CoverTab[49530]++

																if isEOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:344
				_go_fuzz_dep_.CoverTab[49542]++
																	return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:345
				// _ = "end of CoverTab[49542]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:346
				_go_fuzz_dep_.CoverTab[49543]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:346
				// _ = "end of CoverTab[49543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:346
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:346
			// _ = "end of CoverTab[49530]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:346
			_go_fuzz_dep_.CoverTab[49531]++
																switch ch := d.in[0]; ch {
			case '{', '<':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:348
				_go_fuzz_dep_.CoverTab[49544]++
																	d.pushOpenStack(ch)
																	return d.consumeToken(MessageOpen, 1, 0), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:350
				// _ = "end of CoverTab[49544]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:351
				_go_fuzz_dep_.CoverTab[49545]++
																	return d.parseScalar()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:352
				// _ = "end of CoverTab[49545]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:353
			// _ = "end of CoverTab[49531]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:353
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:353
			_go_fuzz_dep_.CoverTab[49532]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:353
			// _ = "end of CoverTab[49532]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:354
		// _ = "end of CoverTab[49441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:354
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:354
		_go_fuzz_dep_.CoverTab[49442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:354
		// _ = "end of CoverTab[49442]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:355
	// _ = "end of CoverTab[49425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:355
	_go_fuzz_dep_.CoverTab[49426]++

														line, column := d.Position(len(d.orig) - len(d.in))
														panic(fmt.Sprintf("Decoder.parseNext: bug at handling line %d:%d with lastKind=%v", line, column, lastKind))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:358
	// _ = "end of CoverTab[49426]"
}

var otherCloseChar = map[byte]byte{
	'}':	'>',
	'>':	'}',
}

// currentOpenKind indicates whether current position is inside a message, list
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:366
// or top-level message by returning MessageOpen, ListOpen or bof respectively.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:366
// If the returned kind is either a MessageOpen or ListOpen, it also returns the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:366
// corresponding closing character.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:370
func (d *Decoder) currentOpenKind() (Kind, byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:370
	_go_fuzz_dep_.CoverTab[49546]++
														if len(d.openStack) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:371
		_go_fuzz_dep_.CoverTab[49549]++
															return bof, 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:372
		// _ = "end of CoverTab[49549]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:373
		_go_fuzz_dep_.CoverTab[49550]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:373
		// _ = "end of CoverTab[49550]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:373
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:373
	// _ = "end of CoverTab[49546]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:373
	_go_fuzz_dep_.CoverTab[49547]++
														openCh := d.openStack[len(d.openStack)-1]
														switch openCh {
	case '{':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:376
		_go_fuzz_dep_.CoverTab[49551]++
															return MessageOpen, '}'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:377
		// _ = "end of CoverTab[49551]"
	case '<':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:378
		_go_fuzz_dep_.CoverTab[49552]++
															return MessageOpen, '>'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:379
		// _ = "end of CoverTab[49552]"
	case '[':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:380
		_go_fuzz_dep_.CoverTab[49553]++
															return ListOpen, ']'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:381
		// _ = "end of CoverTab[49553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:381
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:381
		_go_fuzz_dep_.CoverTab[49554]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:381
		// _ = "end of CoverTab[49554]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:382
	// _ = "end of CoverTab[49547]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:382
	_go_fuzz_dep_.CoverTab[49548]++
														panic(fmt.Sprintf("Decoder: openStack contains invalid byte %c", openCh))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:383
	// _ = "end of CoverTab[49548]"
}

func (d *Decoder) pushOpenStack(ch byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:386
	_go_fuzz_dep_.CoverTab[49555]++
														d.openStack = append(d.openStack, ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:387
	// _ = "end of CoverTab[49555]"
}

func (d *Decoder) popOpenStack() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:390
	_go_fuzz_dep_.CoverTab[49556]++
														d.openStack = d.openStack[:len(d.openStack)-1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:391
	// _ = "end of CoverTab[49556]"
}

// parseFieldName parses field name and separator.
func (d *Decoder) parseFieldName() (tok Token, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:395
	_go_fuzz_dep_.CoverTab[49557]++
														defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:396
		_go_fuzz_dep_.CoverTab[49562]++
															if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:397
			_go_fuzz_dep_.CoverTab[49563]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:397
			return d.tryConsumeChar(':')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:397
			// _ = "end of CoverTab[49563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:397
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:397
			_go_fuzz_dep_.CoverTab[49564]++
																tok.attrs |= hasSeparator
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:398
			// _ = "end of CoverTab[49564]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:399
			_go_fuzz_dep_.CoverTab[49565]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:399
			// _ = "end of CoverTab[49565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:399
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:399
		// _ = "end of CoverTab[49562]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:400
	// _ = "end of CoverTab[49557]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:400
	_go_fuzz_dep_.CoverTab[49558]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:403
	if d.in[0] == '[' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:403
		_go_fuzz_dep_.CoverTab[49566]++
															return d.parseTypeName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:404
		// _ = "end of CoverTab[49566]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:405
		_go_fuzz_dep_.CoverTab[49567]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:405
		// _ = "end of CoverTab[49567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:405
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:405
	// _ = "end of CoverTab[49558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:405
	_go_fuzz_dep_.CoverTab[49559]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:408
	if size := parseIdent(d.in, false); size > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:408
		_go_fuzz_dep_.CoverTab[49568]++
															return d.consumeToken(Name, size, uint8(IdentName)), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:409
		// _ = "end of CoverTab[49568]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:410
		_go_fuzz_dep_.CoverTab[49569]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:410
		// _ = "end of CoverTab[49569]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:410
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:410
	// _ = "end of CoverTab[49559]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:410
	_go_fuzz_dep_.CoverTab[49560]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:414
	if num := parseNumber(d.in); num.size > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:414
		_go_fuzz_dep_.CoverTab[49570]++
															if !num.neg && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:415
			_go_fuzz_dep_.CoverTab[49572]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:415
			return num.kind == numDec
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:415
			// _ = "end of CoverTab[49572]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:415
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:415
			_go_fuzz_dep_.CoverTab[49573]++
																if _, err := strconv.ParseInt(string(d.in[:num.size]), 10, 32); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:416
				_go_fuzz_dep_.CoverTab[49574]++
																	return d.consumeToken(Name, num.size, uint8(FieldNumber)), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:417
				// _ = "end of CoverTab[49574]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:418
				_go_fuzz_dep_.CoverTab[49575]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:418
				// _ = "end of CoverTab[49575]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:418
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:418
			// _ = "end of CoverTab[49573]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:419
			_go_fuzz_dep_.CoverTab[49576]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:419
			// _ = "end of CoverTab[49576]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:419
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:419
		// _ = "end of CoverTab[49570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:419
		_go_fuzz_dep_.CoverTab[49571]++
															return Token{}, d.newSyntaxError("invalid field number: %s", d.in[:num.size])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:420
		// _ = "end of CoverTab[49571]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:421
		_go_fuzz_dep_.CoverTab[49577]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:421
		// _ = "end of CoverTab[49577]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:421
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:421
	// _ = "end of CoverTab[49560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:421
	_go_fuzz_dep_.CoverTab[49561]++

														return Token{}, d.newSyntaxError("invalid field name: %s", errId(d.in))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:423
	// _ = "end of CoverTab[49561]"
}

// parseTypeName parses Any type URL or extension field name. The name is
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:426
// enclosed in [ and ] characters. The C++ parser does not handle many legal URL
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:426
// strings. This implementation is more liberal and allows for the pattern
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:426
// ^[-_a-zA-Z0-9]+([./][-_a-zA-Z0-9]+)*`). Whitespaces and comments are allowed
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:426
// in between [ ], '.', '/' and the sub names.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:431
func (d *Decoder) parseTypeName() (Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:431
	_go_fuzz_dep_.CoverTab[49578]++
														startPos := len(d.orig) - len(d.in)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:435
	s := consume(d.in[1:], 0)
	if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:436
		_go_fuzz_dep_.CoverTab[49584]++
															return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:437
		// _ = "end of CoverTab[49584]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:438
		_go_fuzz_dep_.CoverTab[49585]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:438
		// _ = "end of CoverTab[49585]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:438
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:438
	// _ = "end of CoverTab[49578]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:438
	_go_fuzz_dep_.CoverTab[49579]++

														var name []byte
														for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:441
		_go_fuzz_dep_.CoverTab[49586]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:441
		return isTypeNameChar(s[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:441
		// _ = "end of CoverTab[49586]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:441
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:441
		_go_fuzz_dep_.CoverTab[49587]++
															name = append(name, s[0])
															s = s[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:443
		// _ = "end of CoverTab[49587]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:444
	// _ = "end of CoverTab[49579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:444
	_go_fuzz_dep_.CoverTab[49580]++
														s = consume(s, 0)

														var closed bool
														for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:448
		_go_fuzz_dep_.CoverTab[49588]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:448
		return !closed
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:448
		// _ = "end of CoverTab[49588]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:448
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:448
		_go_fuzz_dep_.CoverTab[49589]++
															switch {
		case s[0] == ']':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:450
			_go_fuzz_dep_.CoverTab[49590]++
																s = s[1:]
																closed = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:452
			// _ = "end of CoverTab[49590]"

		case s[0] == '/', s[0] == '.':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:454
			_go_fuzz_dep_.CoverTab[49591]++
																if len(name) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:455
				_go_fuzz_dep_.CoverTab[49595]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:455
				return (name[len(name)-1] == '/' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:455
					_go_fuzz_dep_.CoverTab[49596]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:455
					return name[len(name)-1] == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:455
					// _ = "end of CoverTab[49596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:455
				}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:455
				// _ = "end of CoverTab[49595]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:455
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:455
				_go_fuzz_dep_.CoverTab[49597]++
																	return Token{}, d.newSyntaxError("invalid type URL/extension field name: %s",
					d.orig[startPos:len(d.orig)-len(s)+1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:457
				// _ = "end of CoverTab[49597]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:458
				_go_fuzz_dep_.CoverTab[49598]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:458
				// _ = "end of CoverTab[49598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:458
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:458
			// _ = "end of CoverTab[49591]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:458
			_go_fuzz_dep_.CoverTab[49592]++
																name = append(name, s[0])
																s = s[1:]
																s = consume(s, 0)
																for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:462
				_go_fuzz_dep_.CoverTab[49599]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:462
				return isTypeNameChar(s[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:462
				// _ = "end of CoverTab[49599]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:462
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:462
				_go_fuzz_dep_.CoverTab[49600]++
																	name = append(name, s[0])
																	s = s[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:464
				// _ = "end of CoverTab[49600]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:465
			// _ = "end of CoverTab[49592]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:465
			_go_fuzz_dep_.CoverTab[49593]++
																s = consume(s, 0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:466
			// _ = "end of CoverTab[49593]"

		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:468
			_go_fuzz_dep_.CoverTab[49594]++
																return Token{}, d.newSyntaxError(
				"invalid type URL/extension field name: %s", d.orig[startPos:len(d.orig)-len(s)+1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:470
			// _ = "end of CoverTab[49594]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:471
		// _ = "end of CoverTab[49589]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:472
	// _ = "end of CoverTab[49580]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:472
	_go_fuzz_dep_.CoverTab[49581]++

														if !closed {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:474
		_go_fuzz_dep_.CoverTab[49601]++
															return Token{}, ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:475
		// _ = "end of CoverTab[49601]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:476
		_go_fuzz_dep_.CoverTab[49602]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:476
		// _ = "end of CoverTab[49602]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:476
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:476
	// _ = "end of CoverTab[49581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:476
	_go_fuzz_dep_.CoverTab[49582]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:479
	size := len(name)
	if size == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		_go_fuzz_dep_.CoverTab[49603]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		return name[0] == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		// _ = "end of CoverTab[49603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		_go_fuzz_dep_.CoverTab[49604]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		return name[size-1] == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		// _ = "end of CoverTab[49604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		_go_fuzz_dep_.CoverTab[49605]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		return name[size-1] == '/'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		// _ = "end of CoverTab[49605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:480
		_go_fuzz_dep_.CoverTab[49606]++
															return Token{}, d.newSyntaxError("invalid type URL/extension field name: %s",
			d.orig[startPos:len(d.orig)-len(s)])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:482
		// _ = "end of CoverTab[49606]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:483
		_go_fuzz_dep_.CoverTab[49607]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:483
		// _ = "end of CoverTab[49607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:483
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:483
	// _ = "end of CoverTab[49582]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:483
	_go_fuzz_dep_.CoverTab[49583]++

														d.in = s
														endPos := len(d.orig) - len(d.in)
														d.consume(0)

														return Token{
		kind:	Name,
		attrs:	uint8(TypeName),
		pos:	startPos,
		raw:	d.orig[startPos:endPos],
		str:	string(name),
	}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:495
	// _ = "end of CoverTab[49583]"
}

func isTypeNameChar(b byte) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:498
	_go_fuzz_dep_.CoverTab[49608]++
														return (b == '-' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:499
		_go_fuzz_dep_.CoverTab[49609]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:499
		return b == '_'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:499
		// _ = "end of CoverTab[49609]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:499
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:499
		_go_fuzz_dep_.CoverTab[49610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:499
		return ('0' <= b && func() bool {
																_go_fuzz_dep_.CoverTab[49611]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:500
			return b <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:500
			// _ = "end of CoverTab[49611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:500
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:500
		// _ = "end of CoverTab[49610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:500
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:500
		_go_fuzz_dep_.CoverTab[49612]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:500
		return ('a' <= b && func() bool {
																_go_fuzz_dep_.CoverTab[49613]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:501
			return b <= 'z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:501
			// _ = "end of CoverTab[49613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:501
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:501
		// _ = "end of CoverTab[49612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:501
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:501
		_go_fuzz_dep_.CoverTab[49614]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:501
		return ('A' <= b && func() bool {
																_go_fuzz_dep_.CoverTab[49615]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:502
			return b <= 'Z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:502
			// _ = "end of CoverTab[49615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:502
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:502
		// _ = "end of CoverTab[49614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:502
	}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:502
	// _ = "end of CoverTab[49608]"
}

func isWhiteSpace(b byte) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:505
	_go_fuzz_dep_.CoverTab[49616]++
														switch b {
	case ' ', '\n', '\r', '\t':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:507
		_go_fuzz_dep_.CoverTab[49617]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:508
		// _ = "end of CoverTab[49617]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:509
		_go_fuzz_dep_.CoverTab[49618]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:510
		// _ = "end of CoverTab[49618]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:511
	// _ = "end of CoverTab[49616]"
}

// parseIdent parses an unquoted proto identifier and returns size.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:514
// If allowNeg is true, it allows '-' to be the first character in the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:514
// identifier. This is used when parsing literal values like -infinity, etc.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:514
// Regular expression matches an identifier: `^[_a-zA-Z][_a-zA-Z0-9]*`
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:518
func parseIdent(input []byte, allowNeg bool) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:518
	_go_fuzz_dep_.CoverTab[49619]++
														var size int

														s := input
														if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:522
		_go_fuzz_dep_.CoverTab[49625]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:523
		// _ = "end of CoverTab[49625]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:524
		_go_fuzz_dep_.CoverTab[49626]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:524
		// _ = "end of CoverTab[49626]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:524
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:524
	// _ = "end of CoverTab[49619]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:524
	_go_fuzz_dep_.CoverTab[49620]++

														if allowNeg && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:526
		_go_fuzz_dep_.CoverTab[49627]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:526
		return s[0] == '-'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:526
		// _ = "end of CoverTab[49627]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:526
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:526
		_go_fuzz_dep_.CoverTab[49628]++
															s = s[1:]
															size++
															if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:529
			_go_fuzz_dep_.CoverTab[49629]++
																return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:530
			// _ = "end of CoverTab[49629]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:531
			_go_fuzz_dep_.CoverTab[49630]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:531
			// _ = "end of CoverTab[49630]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:531
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:531
		// _ = "end of CoverTab[49628]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:532
		_go_fuzz_dep_.CoverTab[49631]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:532
		// _ = "end of CoverTab[49631]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:532
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:532
	// _ = "end of CoverTab[49620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:532
	_go_fuzz_dep_.CoverTab[49621]++

														switch {
	case s[0] == '_',
		'a' <= s[0] && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:536
			_go_fuzz_dep_.CoverTab[49634]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:536
			return s[0] <= 'z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:536
			// _ = "end of CoverTab[49634]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:536
		}(),
		'A' <= s[0] && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:537
			_go_fuzz_dep_.CoverTab[49635]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:537
			return s[0] <= 'Z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:537
			// _ = "end of CoverTab[49635]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:537
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:537
		_go_fuzz_dep_.CoverTab[49632]++
															s = s[1:]
															size++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:539
		// _ = "end of CoverTab[49632]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:540
		_go_fuzz_dep_.CoverTab[49633]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:541
		// _ = "end of CoverTab[49633]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:542
	// _ = "end of CoverTab[49621]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:542
	_go_fuzz_dep_.CoverTab[49622]++

														for len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:544
		_go_fuzz_dep_.CoverTab[49636]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:544
		return (s[0] == '_' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:544
			_go_fuzz_dep_.CoverTab[49637]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:544
			return 'a' <= s[0] && func() bool {
																	_go_fuzz_dep_.CoverTab[49638]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:545
				return s[0] <= 'z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:545
				// _ = "end of CoverTab[49638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:545
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:545
			// _ = "end of CoverTab[49637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:545
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:545
			_go_fuzz_dep_.CoverTab[49639]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:545
			return 'A' <= s[0] && func() bool {
																	_go_fuzz_dep_.CoverTab[49640]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:546
				return s[0] <= 'Z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:546
				// _ = "end of CoverTab[49640]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:546
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:546
			// _ = "end of CoverTab[49639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:546
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:546
			_go_fuzz_dep_.CoverTab[49641]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:546
			return '0' <= s[0] && func() bool {
																	_go_fuzz_dep_.CoverTab[49642]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:547
				return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:547
				// _ = "end of CoverTab[49642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:547
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:547
			// _ = "end of CoverTab[49641]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:547
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:547
		// _ = "end of CoverTab[49636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:547
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:547
		_go_fuzz_dep_.CoverTab[49643]++
															s = s[1:]
															size++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:549
		// _ = "end of CoverTab[49643]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:550
	// _ = "end of CoverTab[49622]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:550
	_go_fuzz_dep_.CoverTab[49623]++

														if len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:552
		_go_fuzz_dep_.CoverTab[49644]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:552
		return !isDelim(s[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:552
		// _ = "end of CoverTab[49644]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:552
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:552
		_go_fuzz_dep_.CoverTab[49645]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:553
		// _ = "end of CoverTab[49645]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:554
		_go_fuzz_dep_.CoverTab[49646]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:554
		// _ = "end of CoverTab[49646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:554
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:554
	// _ = "end of CoverTab[49623]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:554
	_go_fuzz_dep_.CoverTab[49624]++

														return size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:556
	// _ = "end of CoverTab[49624]"
}

// parseScalar parses for a string, literal or number value.
func (d *Decoder) parseScalar() (Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:560
	_go_fuzz_dep_.CoverTab[49647]++
														if d.in[0] == '"' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:561
		_go_fuzz_dep_.CoverTab[49651]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:561
		return d.in[0] == '\''
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:561
		// _ = "end of CoverTab[49651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:561
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:561
		_go_fuzz_dep_.CoverTab[49652]++
															return d.parseStringValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:562
		// _ = "end of CoverTab[49652]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:563
		_go_fuzz_dep_.CoverTab[49653]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:563
		// _ = "end of CoverTab[49653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:563
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:563
	// _ = "end of CoverTab[49647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:563
	_go_fuzz_dep_.CoverTab[49648]++

														if tok, ok := d.parseLiteralValue(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:565
		_go_fuzz_dep_.CoverTab[49654]++
															return tok, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:566
		// _ = "end of CoverTab[49654]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:567
		_go_fuzz_dep_.CoverTab[49655]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:567
		// _ = "end of CoverTab[49655]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:567
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:567
	// _ = "end of CoverTab[49648]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:567
	_go_fuzz_dep_.CoverTab[49649]++

														if tok, ok := d.parseNumberValue(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:569
		_go_fuzz_dep_.CoverTab[49656]++
															return tok, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:570
		// _ = "end of CoverTab[49656]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:571
		_go_fuzz_dep_.CoverTab[49657]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:571
		// _ = "end of CoverTab[49657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:571
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:571
	// _ = "end of CoverTab[49649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:571
	_go_fuzz_dep_.CoverTab[49650]++

														return Token{}, d.newSyntaxError("invalid scalar value: %s", errId(d.in))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:573
	// _ = "end of CoverTab[49650]"
}

// parseLiteralValue parses a literal value. A literal value is used for
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:576
// bools, special floats and enums. This function simply identifies that the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:576
// field value is a literal.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:579
func (d *Decoder) parseLiteralValue() (Token, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:579
	_go_fuzz_dep_.CoverTab[49658]++
														size := parseIdent(d.in, true)
														if size == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:581
		_go_fuzz_dep_.CoverTab[49660]++
															return Token{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:582
		// _ = "end of CoverTab[49660]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:583
		_go_fuzz_dep_.CoverTab[49661]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:583
		// _ = "end of CoverTab[49661]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:583
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:583
	// _ = "end of CoverTab[49658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:583
	_go_fuzz_dep_.CoverTab[49659]++
														return d.consumeToken(Scalar, size, literalValue), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:584
	// _ = "end of CoverTab[49659]"
}

// consumeToken constructs a Token for given Kind from d.in and consumes given
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:587
// size-length from it.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:589
func (d *Decoder) consumeToken(kind Kind, size int, attrs uint8) Token {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:589
	_go_fuzz_dep_.CoverTab[49662]++

														tok := Token{
		kind:	kind,
		attrs:	attrs,
		pos:	len(d.orig) - len(d.in),
		raw:	d.in[:size],
	}
														d.consume(size)
														return tok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:598
	// _ = "end of CoverTab[49662]"
}

// newSyntaxError returns a syntax error with line and column information for
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:601
// current position.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:603
func (d *Decoder) newSyntaxError(f string, x ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:603
	_go_fuzz_dep_.CoverTab[49663]++
														e := errors.New(f, x...)
														line, column := d.Position(len(d.orig) - len(d.in))
														return errors.New("syntax error (line %d:%d): %v", line, column, e)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:606
	// _ = "end of CoverTab[49663]"
}

// Position returns line and column number of given index of the original input.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:609
// It will panic if index is out of range.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:611
func (d *Decoder) Position(idx int) (line int, column int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:611
	_go_fuzz_dep_.CoverTab[49664]++
														b := d.orig[:idx]
														line = bytes.Count(b, []byte("\n")) + 1
														if i := bytes.LastIndexByte(b, '\n'); i >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:614
		_go_fuzz_dep_.CoverTab[49666]++
															b = b[i+1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:615
		// _ = "end of CoverTab[49666]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:616
		_go_fuzz_dep_.CoverTab[49667]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:616
		// _ = "end of CoverTab[49667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:616
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:616
	// _ = "end of CoverTab[49664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:616
	_go_fuzz_dep_.CoverTab[49665]++
														column = utf8.RuneCount(b) + 1
														return line, column
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:618
	// _ = "end of CoverTab[49665]"
}

func (d *Decoder) tryConsumeChar(c byte) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:621
	_go_fuzz_dep_.CoverTab[49668]++
														if len(d.in) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:622
		_go_fuzz_dep_.CoverTab[49670]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:622
		return d.in[0] == c
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:622
		// _ = "end of CoverTab[49670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:622
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:622
		_go_fuzz_dep_.CoverTab[49671]++
															d.consume(1)
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:624
		// _ = "end of CoverTab[49671]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:625
		_go_fuzz_dep_.CoverTab[49672]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:625
		// _ = "end of CoverTab[49672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:625
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:625
	// _ = "end of CoverTab[49668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:625
	_go_fuzz_dep_.CoverTab[49669]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:626
	// _ = "end of CoverTab[49669]"
}

// consume consumes n bytes of input and any subsequent whitespace or comments.
func (d *Decoder) consume(n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:630
	_go_fuzz_dep_.CoverTab[49673]++
														d.in = consume(d.in, n)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:632
	// _ = "end of CoverTab[49673]"
}

// consume consumes n bytes of input and any subsequent whitespace or comments.
func consume(b []byte, n int) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:636
	_go_fuzz_dep_.CoverTab[49674]++
														b = b[n:]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:638
		_go_fuzz_dep_.CoverTab[49676]++
															switch b[0] {
		case ' ', '\n', '\r', '\t':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:640
			_go_fuzz_dep_.CoverTab[49677]++
																b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:641
			// _ = "end of CoverTab[49677]"
		case '#':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:642
			_go_fuzz_dep_.CoverTab[49678]++
																if i := bytes.IndexByte(b, '\n'); i >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:643
				_go_fuzz_dep_.CoverTab[49680]++
																	b = b[i+len("\n"):]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:644
				// _ = "end of CoverTab[49680]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:645
				_go_fuzz_dep_.CoverTab[49681]++
																	b = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:646
				// _ = "end of CoverTab[49681]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:647
			// _ = "end of CoverTab[49678]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:648
			_go_fuzz_dep_.CoverTab[49679]++
																return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:649
			// _ = "end of CoverTab[49679]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:650
		// _ = "end of CoverTab[49676]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:651
	// _ = "end of CoverTab[49674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:651
	_go_fuzz_dep_.CoverTab[49675]++
														return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:652
	// _ = "end of CoverTab[49675]"
}

// errId extracts a byte sequence that looks like an invalid ID
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:655
// (for the purposes of error reporting).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:657
func errId(seq []byte) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:657
	_go_fuzz_dep_.CoverTab[49682]++
														const maxLen = 32
														for i := 0; i < len(seq); {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:659
		_go_fuzz_dep_.CoverTab[49684]++
															if i > maxLen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:660
			_go_fuzz_dep_.CoverTab[49687]++
																return append(seq[:i:i], "…"...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:661
			// _ = "end of CoverTab[49687]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:662
			_go_fuzz_dep_.CoverTab[49688]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:662
			// _ = "end of CoverTab[49688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:662
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:662
		// _ = "end of CoverTab[49684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:662
		_go_fuzz_dep_.CoverTab[49685]++
															r, size := utf8.DecodeRune(seq[i:])
															if r > utf8.RuneSelf || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:664
			_go_fuzz_dep_.CoverTab[49689]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:664
			return (r != '/' && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:664
				_go_fuzz_dep_.CoverTab[49690]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:664
				return isDelim(byte(r))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:664
				// _ = "end of CoverTab[49690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:664
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:664
			// _ = "end of CoverTab[49689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:664
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:664
			_go_fuzz_dep_.CoverTab[49691]++
																if i == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:665
				_go_fuzz_dep_.CoverTab[49693]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:669
				i = size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:669
				// _ = "end of CoverTab[49693]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:670
				_go_fuzz_dep_.CoverTab[49694]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:670
				// _ = "end of CoverTab[49694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:670
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:670
			// _ = "end of CoverTab[49691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:670
			_go_fuzz_dep_.CoverTab[49692]++
																return seq[:i:i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:671
			// _ = "end of CoverTab[49692]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:672
			_go_fuzz_dep_.CoverTab[49695]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:672
			// _ = "end of CoverTab[49695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:672
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:672
		// _ = "end of CoverTab[49685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:672
		_go_fuzz_dep_.CoverTab[49686]++
															i += size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:673
		// _ = "end of CoverTab[49686]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:674
	// _ = "end of CoverTab[49682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:674
	_go_fuzz_dep_.CoverTab[49683]++

														return seq
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:676
	// _ = "end of CoverTab[49683]"
}

// isDelim returns true if given byte is a delimiter character.
func isDelim(c byte) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:680
	_go_fuzz_dep_.CoverTab[49696]++
														return !(c == '-' || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		_go_fuzz_dep_.CoverTab[49697]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		return c == '+'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		// _ = "end of CoverTab[49697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		_go_fuzz_dep_.CoverTab[49698]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		return c == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		// _ = "end of CoverTab[49698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		_go_fuzz_dep_.CoverTab[49699]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		return c == '_'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		// _ = "end of CoverTab[49699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		_go_fuzz_dep_.CoverTab[49700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:681
		return ('a' <= c && func() bool {
																_go_fuzz_dep_.CoverTab[49701]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:682
			return c <= 'z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:682
			// _ = "end of CoverTab[49701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:682
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:682
		// _ = "end of CoverTab[49700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:682
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:682
		_go_fuzz_dep_.CoverTab[49702]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:682
		return ('A' <= c && func() bool {
																_go_fuzz_dep_.CoverTab[49703]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:683
			return c <= 'Z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:683
			// _ = "end of CoverTab[49703]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:683
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:683
		// _ = "end of CoverTab[49702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:683
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:683
		_go_fuzz_dep_.CoverTab[49704]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:683
		return ('0' <= c && func() bool {
																_go_fuzz_dep_.CoverTab[49705]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:684
			return c <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:684
			// _ = "end of CoverTab[49705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:684
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:684
		// _ = "end of CoverTab[49704]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:684
	}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:684
	// _ = "end of CoverTab[49696]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:685
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/text/decode.go:685
var _ = _go_fuzz_dep_.CoverTab
