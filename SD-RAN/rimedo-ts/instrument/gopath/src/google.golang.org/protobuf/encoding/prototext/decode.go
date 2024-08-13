// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:5
package prototext

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:5
)

import (
	"fmt"
	"unicode/utf8"

	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/encoding/text"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/internal/set"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Unmarshal reads the given []byte into the given proto.Message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:24
// The provided message must be mutable (e.g., a non-nil pointer to a message).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:26
func Unmarshal(b []byte, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:26
	_go_fuzz_dep_.CoverTab[51604]++
													return UnmarshalOptions{}.Unmarshal(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:27
	// _ = "end of CoverTab[51604]"
}

// UnmarshalOptions is a configurable textproto format unmarshaler.
type UnmarshalOptions struct {
	pragma.NoUnkeyedLiterals

	// AllowPartial accepts input for messages that will result in missing
	// required fields. If AllowPartial is false (the default), Unmarshal will
	// return error if there are any missing required fields.
	AllowPartial	bool

	// DiscardUnknown specifies whether to ignore unknown fields when parsing.
	// An unknown field is any field whose field name or field number does not
	// resolve to any known or extension field in the message.
	// By default, unmarshal rejects unknown fields as an error.
	DiscardUnknown	bool

	// Resolver is used for looking up types when unmarshaling
	// google.protobuf.Any messages or extension fields.
	// If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver	interface {
		protoregistry.MessageTypeResolver
		protoregistry.ExtensionTypeResolver
	}
}

// Unmarshal reads the given []byte and populates the given proto.Message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:54
// using options in the UnmarshalOptions object.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:54
// The provided message must be mutable (e.g., a non-nil pointer to a message).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:57
func (o UnmarshalOptions) Unmarshal(b []byte, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:57
	_go_fuzz_dep_.CoverTab[51605]++
													return o.unmarshal(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:58
	// _ = "end of CoverTab[51605]"
}

// unmarshal is a centralized function that all unmarshal operations go through.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:61
// For profiling purposes, avoid changing the name of this function or
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:61
// introducing other code paths for unmarshal that do not go through this.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:64
func (o UnmarshalOptions) unmarshal(b []byte, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:64
	_go_fuzz_dep_.CoverTab[51606]++
													proto.Reset(m)

													if o.Resolver == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:67
		_go_fuzz_dep_.CoverTab[51610]++
														o.Resolver = protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:68
		// _ = "end of CoverTab[51610]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:69
		_go_fuzz_dep_.CoverTab[51611]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:69
		// _ = "end of CoverTab[51611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:69
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:69
	// _ = "end of CoverTab[51606]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:69
	_go_fuzz_dep_.CoverTab[51607]++

													dec := decoder{text.NewDecoder(b), o}
													if err := dec.unmarshalMessage(m.ProtoReflect(), false); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:72
		_go_fuzz_dep_.CoverTab[51612]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:73
		// _ = "end of CoverTab[51612]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:74
		_go_fuzz_dep_.CoverTab[51613]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:74
		// _ = "end of CoverTab[51613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:74
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:74
	// _ = "end of CoverTab[51607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:74
	_go_fuzz_dep_.CoverTab[51608]++
													if o.AllowPartial {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:75
		_go_fuzz_dep_.CoverTab[51614]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:76
		// _ = "end of CoverTab[51614]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:77
		_go_fuzz_dep_.CoverTab[51615]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:77
		// _ = "end of CoverTab[51615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:77
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:77
	// _ = "end of CoverTab[51608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:77
	_go_fuzz_dep_.CoverTab[51609]++
													return proto.CheckInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:78
	// _ = "end of CoverTab[51609]"
}

type decoder struct {
	*text.Decoder
	opts	UnmarshalOptions
}

// newError returns an error object with position info.
func (d decoder) newError(pos int, f string, x ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:87
	_go_fuzz_dep_.CoverTab[51616]++
													line, column := d.Position(pos)
													head := fmt.Sprintf("(line %d:%d): ", line, column)
													return errors.New(head+f, x...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:90
	// _ = "end of CoverTab[51616]"
}

// unexpectedTokenError returns a syntax error for the given unexpected token.
func (d decoder) unexpectedTokenError(tok text.Token) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:94
	_go_fuzz_dep_.CoverTab[51617]++
													return d.syntaxError(tok.Pos(), "unexpected token: %s", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:95
	// _ = "end of CoverTab[51617]"
}

// syntaxError returns a syntax error for given position.
func (d decoder) syntaxError(pos int, f string, x ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:99
	_go_fuzz_dep_.CoverTab[51618]++
													line, column := d.Position(pos)
													head := fmt.Sprintf("syntax error (line %d:%d): ", line, column)
													return errors.New(head+f, x...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:102
	// _ = "end of CoverTab[51618]"
}

// unmarshalMessage unmarshals into the given protoreflect.Message.
func (d decoder) unmarshalMessage(m protoreflect.Message, checkDelims bool) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:106
	_go_fuzz_dep_.CoverTab[51619]++
													messageDesc := m.Descriptor()
													if !flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:108
		_go_fuzz_dep_.CoverTab[51624]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:108
		return messageset.IsMessageSet(messageDesc)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:108
		// _ = "end of CoverTab[51624]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:108
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:108
		_go_fuzz_dep_.CoverTab[51625]++
														return errors.New("no support for proto1 MessageSets")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:109
		// _ = "end of CoverTab[51625]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:110
		_go_fuzz_dep_.CoverTab[51626]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:110
		// _ = "end of CoverTab[51626]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:110
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:110
	// _ = "end of CoverTab[51619]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:110
	_go_fuzz_dep_.CoverTab[51620]++

													if messageDesc.FullName() == genid.Any_message_fullname {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:112
		_go_fuzz_dep_.CoverTab[51627]++
														return d.unmarshalAny(m, checkDelims)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:113
		// _ = "end of CoverTab[51627]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:114
		_go_fuzz_dep_.CoverTab[51628]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:114
		// _ = "end of CoverTab[51628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:114
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:114
	// _ = "end of CoverTab[51620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:114
	_go_fuzz_dep_.CoverTab[51621]++

													if checkDelims {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:116
		_go_fuzz_dep_.CoverTab[51629]++
														tok, err := d.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:118
			_go_fuzz_dep_.CoverTab[51631]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:119
			// _ = "end of CoverTab[51631]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:120
			_go_fuzz_dep_.CoverTab[51632]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:120
			// _ = "end of CoverTab[51632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:120
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:120
		// _ = "end of CoverTab[51629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:120
		_go_fuzz_dep_.CoverTab[51630]++

														if tok.Kind() != text.MessageOpen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:122
			_go_fuzz_dep_.CoverTab[51633]++
															return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:123
			// _ = "end of CoverTab[51633]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:124
			_go_fuzz_dep_.CoverTab[51634]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:124
			// _ = "end of CoverTab[51634]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:124
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:124
		// _ = "end of CoverTab[51630]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:125
		_go_fuzz_dep_.CoverTab[51635]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:125
		// _ = "end of CoverTab[51635]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:125
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:125
	// _ = "end of CoverTab[51621]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:125
	_go_fuzz_dep_.CoverTab[51622]++

													var seenNums set.Ints
													var seenOneofs set.Ints
													fieldDescs := messageDesc.Fields()

													for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:131
		_go_fuzz_dep_.CoverTab[51636]++

														tok, err := d.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:134
			_go_fuzz_dep_.CoverTab[51644]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:135
			// _ = "end of CoverTab[51644]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:136
			_go_fuzz_dep_.CoverTab[51645]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:136
			// _ = "end of CoverTab[51645]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:136
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:136
		// _ = "end of CoverTab[51636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:136
		_go_fuzz_dep_.CoverTab[51637]++
														switch typ := tok.Kind(); typ {
		case text.Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:138
			_go_fuzz_dep_.CoverTab[51646]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:138
			// _ = "end of CoverTab[51646]"

		case text.EOF:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:140
			_go_fuzz_dep_.CoverTab[51647]++
															if checkDelims {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:141
				_go_fuzz_dep_.CoverTab[51651]++
																return text.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:142
				// _ = "end of CoverTab[51651]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:143
				_go_fuzz_dep_.CoverTab[51652]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:143
				// _ = "end of CoverTab[51652]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:143
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:143
			// _ = "end of CoverTab[51647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:143
			_go_fuzz_dep_.CoverTab[51648]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:144
			// _ = "end of CoverTab[51648]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:145
			_go_fuzz_dep_.CoverTab[51649]++
															if checkDelims && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:146
				_go_fuzz_dep_.CoverTab[51653]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:146
				return typ == text.MessageClose
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:146
				// _ = "end of CoverTab[51653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:146
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:146
				_go_fuzz_dep_.CoverTab[51654]++
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:147
				// _ = "end of CoverTab[51654]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:148
				_go_fuzz_dep_.CoverTab[51655]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:148
				// _ = "end of CoverTab[51655]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:148
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:148
			// _ = "end of CoverTab[51649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:148
			_go_fuzz_dep_.CoverTab[51650]++
															return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:149
			// _ = "end of CoverTab[51650]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:150
		// _ = "end of CoverTab[51637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:150
		_go_fuzz_dep_.CoverTab[51638]++

		// Resolve the field descriptor.
		var name protoreflect.Name
		var fd protoreflect.FieldDescriptor
		var xt protoreflect.ExtensionType
		var xtErr error
		var isFieldNumberName bool

		switch tok.NameKind() {
		case text.IdentName:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:160
			_go_fuzz_dep_.CoverTab[51656]++
															name = protoreflect.Name(tok.IdentName())
															fd = fieldDescs.ByTextName(string(name))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:162
			// _ = "end of CoverTab[51656]"

		case text.TypeName:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:164
			_go_fuzz_dep_.CoverTab[51657]++

															xt, xtErr = d.opts.Resolver.FindExtensionByName(protoreflect.FullName(tok.TypeName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:166
			// _ = "end of CoverTab[51657]"

		case text.FieldNumber:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:168
			_go_fuzz_dep_.CoverTab[51658]++
															isFieldNumberName = true
															num := protoreflect.FieldNumber(tok.FieldNumber())
															if !num.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:171
				_go_fuzz_dep_.CoverTab[51661]++
																return d.newError(tok.Pos(), "invalid field number: %d", num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:172
				// _ = "end of CoverTab[51661]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:173
				_go_fuzz_dep_.CoverTab[51662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:173
				// _ = "end of CoverTab[51662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:173
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:173
			// _ = "end of CoverTab[51658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:173
			_go_fuzz_dep_.CoverTab[51659]++
															fd = fieldDescs.ByNumber(num)
															if fd == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:175
				_go_fuzz_dep_.CoverTab[51663]++
																xt, xtErr = d.opts.Resolver.FindExtensionByNumber(messageDesc.FullName(), num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:176
				// _ = "end of CoverTab[51663]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:177
				_go_fuzz_dep_.CoverTab[51664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:177
				// _ = "end of CoverTab[51664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:177
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:177
			// _ = "end of CoverTab[51659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:177
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:177
			_go_fuzz_dep_.CoverTab[51660]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:177
			// _ = "end of CoverTab[51660]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:178
		// _ = "end of CoverTab[51638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:178
		_go_fuzz_dep_.CoverTab[51639]++

														if xt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:180
			_go_fuzz_dep_.CoverTab[51665]++
															fd = xt.TypeDescriptor()
															if !messageDesc.ExtensionRanges().Has(fd.Number()) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:182
				_go_fuzz_dep_.CoverTab[51666]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:182
				return fd.ContainingMessage().FullName() != messageDesc.FullName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:182
				// _ = "end of CoverTab[51666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:182
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:182
				_go_fuzz_dep_.CoverTab[51667]++
																return d.newError(tok.Pos(), "message %v cannot be extended by %v", messageDesc.FullName(), fd.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:183
				// _ = "end of CoverTab[51667]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:184
				_go_fuzz_dep_.CoverTab[51668]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:184
				// _ = "end of CoverTab[51668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:184
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:184
			// _ = "end of CoverTab[51665]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:185
			_go_fuzz_dep_.CoverTab[51669]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:185
			if xtErr != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:185
				_go_fuzz_dep_.CoverTab[51670]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:185
				return xtErr != protoregistry.NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:185
				// _ = "end of CoverTab[51670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:185
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:185
				_go_fuzz_dep_.CoverTab[51671]++
																return d.newError(tok.Pos(), "unable to resolve [%s]: %v", tok.RawString(), xtErr)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:186
				// _ = "end of CoverTab[51671]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:187
				_go_fuzz_dep_.CoverTab[51672]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:187
				// _ = "end of CoverTab[51672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:187
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:187
			// _ = "end of CoverTab[51669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:187
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:187
		// _ = "end of CoverTab[51639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:187
		_go_fuzz_dep_.CoverTab[51640]++
														if flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:188
			_go_fuzz_dep_.CoverTab[51673]++
															if fd != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:189
				_go_fuzz_dep_.CoverTab[51674]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:189
				return fd.IsWeak()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:189
				// _ = "end of CoverTab[51674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:189
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:189
				_go_fuzz_dep_.CoverTab[51675]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:189
				return fd.Message().IsPlaceholder()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:189
				// _ = "end of CoverTab[51675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:189
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:189
				_go_fuzz_dep_.CoverTab[51676]++
																fd = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:190
				// _ = "end of CoverTab[51676]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:191
				_go_fuzz_dep_.CoverTab[51677]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:191
				// _ = "end of CoverTab[51677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:191
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:191
			// _ = "end of CoverTab[51673]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:192
			_go_fuzz_dep_.CoverTab[51678]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:192
			// _ = "end of CoverTab[51678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:192
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:192
		// _ = "end of CoverTab[51640]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:192
		_go_fuzz_dep_.CoverTab[51641]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:195
		if fd == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:195
			_go_fuzz_dep_.CoverTab[51679]++
															if d.opts.DiscardUnknown || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:196
				_go_fuzz_dep_.CoverTab[51681]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:196
				return messageDesc.ReservedNames().Has(name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:196
				// _ = "end of CoverTab[51681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:196
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:196
				_go_fuzz_dep_.CoverTab[51682]++
																d.skipValue()
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:198
				// _ = "end of CoverTab[51682]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:199
				_go_fuzz_dep_.CoverTab[51683]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:199
				// _ = "end of CoverTab[51683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:199
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:199
			// _ = "end of CoverTab[51679]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:199
			_go_fuzz_dep_.CoverTab[51680]++
															return d.newError(tok.Pos(), "unknown field: %v", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:200
			// _ = "end of CoverTab[51680]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:201
			_go_fuzz_dep_.CoverTab[51684]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:201
			// _ = "end of CoverTab[51684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:201
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:201
		// _ = "end of CoverTab[51641]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:201
		_go_fuzz_dep_.CoverTab[51642]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:204
		if isFieldNumberName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:204
			_go_fuzz_dep_.CoverTab[51685]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:212
			return d.newError(tok.Pos(), "cannot specify field by number: %v", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:212
			// _ = "end of CoverTab[51685]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:213
			_go_fuzz_dep_.CoverTab[51686]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:213
			// _ = "end of CoverTab[51686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:213
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:213
		// _ = "end of CoverTab[51642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:213
		_go_fuzz_dep_.CoverTab[51643]++

														switch {
		case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:216
			_go_fuzz_dep_.CoverTab[51687]++
															kind := fd.Kind()
															if kind != protoreflect.MessageKind && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:218
				_go_fuzz_dep_.CoverTab[51695]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:218
				return kind != protoreflect.GroupKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:218
				// _ = "end of CoverTab[51695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:218
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:218
				_go_fuzz_dep_.CoverTab[51696]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:218
				return !tok.HasSeparator()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:218
				// _ = "end of CoverTab[51696]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:218
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:218
				_go_fuzz_dep_.CoverTab[51697]++
																return d.syntaxError(tok.Pos(), "missing field separator :")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:219
				// _ = "end of CoverTab[51697]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:220
				_go_fuzz_dep_.CoverTab[51698]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:220
				// _ = "end of CoverTab[51698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:220
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:220
			// _ = "end of CoverTab[51687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:220
			_go_fuzz_dep_.CoverTab[51688]++

															list := m.Mutable(fd).List()
															if err := d.unmarshalList(fd, list); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:223
				_go_fuzz_dep_.CoverTab[51699]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:224
				// _ = "end of CoverTab[51699]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:225
				_go_fuzz_dep_.CoverTab[51700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:225
				// _ = "end of CoverTab[51700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:225
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:225
			// _ = "end of CoverTab[51688]"

		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:227
			_go_fuzz_dep_.CoverTab[51689]++
															mmap := m.Mutable(fd).Map()
															if err := d.unmarshalMap(fd, mmap); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:229
				_go_fuzz_dep_.CoverTab[51701]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:230
				// _ = "end of CoverTab[51701]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:231
				_go_fuzz_dep_.CoverTab[51702]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:231
				// _ = "end of CoverTab[51702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:231
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:231
			// _ = "end of CoverTab[51689]"

		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:233
			_go_fuzz_dep_.CoverTab[51690]++
															kind := fd.Kind()
															if kind != protoreflect.MessageKind && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:235
				_go_fuzz_dep_.CoverTab[51703]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:235
				return kind != protoreflect.GroupKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:235
				// _ = "end of CoverTab[51703]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:235
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:235
				_go_fuzz_dep_.CoverTab[51704]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:235
				return !tok.HasSeparator()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:235
				// _ = "end of CoverTab[51704]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:235
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:235
				_go_fuzz_dep_.CoverTab[51705]++
																return d.syntaxError(tok.Pos(), "missing field separator :")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:236
				// _ = "end of CoverTab[51705]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:237
				_go_fuzz_dep_.CoverTab[51706]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:237
				// _ = "end of CoverTab[51706]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:237
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:237
			// _ = "end of CoverTab[51690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:237
			_go_fuzz_dep_.CoverTab[51691]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:240
			if od := fd.ContainingOneof(); od != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:240
				_go_fuzz_dep_.CoverTab[51707]++
																idx := uint64(od.Index())
																if seenOneofs.Has(idx) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:242
					_go_fuzz_dep_.CoverTab[51709]++
																	return d.newError(tok.Pos(), "error parsing %q, oneof %v is already set", tok.RawString(), od.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:243
					// _ = "end of CoverTab[51709]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:244
					_go_fuzz_dep_.CoverTab[51710]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:244
					// _ = "end of CoverTab[51710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:244
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:244
				// _ = "end of CoverTab[51707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:244
				_go_fuzz_dep_.CoverTab[51708]++
																seenOneofs.Set(idx)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:245
				// _ = "end of CoverTab[51708]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:246
				_go_fuzz_dep_.CoverTab[51711]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:246
				// _ = "end of CoverTab[51711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:246
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:246
			// _ = "end of CoverTab[51691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:246
			_go_fuzz_dep_.CoverTab[51692]++

															num := uint64(fd.Number())
															if seenNums.Has(num) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:249
				_go_fuzz_dep_.CoverTab[51712]++
																return d.newError(tok.Pos(), "non-repeated field %q is repeated", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:250
				// _ = "end of CoverTab[51712]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:251
				_go_fuzz_dep_.CoverTab[51713]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:251
				// _ = "end of CoverTab[51713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:251
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:251
			// _ = "end of CoverTab[51692]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:251
			_go_fuzz_dep_.CoverTab[51693]++

															if err := d.unmarshalSingular(fd, m); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:253
				_go_fuzz_dep_.CoverTab[51714]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:254
				// _ = "end of CoverTab[51714]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:255
				_go_fuzz_dep_.CoverTab[51715]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:255
				// _ = "end of CoverTab[51715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:255
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:255
			// _ = "end of CoverTab[51693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:255
			_go_fuzz_dep_.CoverTab[51694]++
															seenNums.Set(num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:256
			// _ = "end of CoverTab[51694]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:257
		// _ = "end of CoverTab[51643]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:258
	// _ = "end of CoverTab[51622]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:258
	_go_fuzz_dep_.CoverTab[51623]++

													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:260
	// _ = "end of CoverTab[51623]"
}

// unmarshalSingular unmarshals a non-repeated field value specified by the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:263
// given FieldDescriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:265
func (d decoder) unmarshalSingular(fd protoreflect.FieldDescriptor, m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:265
	_go_fuzz_dep_.CoverTab[51716]++
													var val protoreflect.Value
													var err error
													switch fd.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:269
		_go_fuzz_dep_.CoverTab[51719]++
														val = m.NewField(fd)
														err = d.unmarshalMessage(val.Message(), true)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:271
		// _ = "end of CoverTab[51719]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:272
		_go_fuzz_dep_.CoverTab[51720]++
														val, err = d.unmarshalScalar(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:273
		// _ = "end of CoverTab[51720]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:274
	// _ = "end of CoverTab[51716]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:274
	_go_fuzz_dep_.CoverTab[51717]++
													if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:275
		_go_fuzz_dep_.CoverTab[51721]++
														m.Set(fd, val)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:276
		// _ = "end of CoverTab[51721]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:277
		_go_fuzz_dep_.CoverTab[51722]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:277
		// _ = "end of CoverTab[51722]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:277
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:277
	// _ = "end of CoverTab[51717]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:277
	_go_fuzz_dep_.CoverTab[51718]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:278
	// _ = "end of CoverTab[51718]"
}

// unmarshalScalar unmarshals a scalar/enum protoreflect.Value specified by the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:281
// given FieldDescriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:283
func (d decoder) unmarshalScalar(fd protoreflect.FieldDescriptor) (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:283
	_go_fuzz_dep_.CoverTab[51723]++
													tok, err := d.Read()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:285
		_go_fuzz_dep_.CoverTab[51727]++
														return protoreflect.Value{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:286
		// _ = "end of CoverTab[51727]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:287
		_go_fuzz_dep_.CoverTab[51728]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:287
		// _ = "end of CoverTab[51728]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:287
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:287
	// _ = "end of CoverTab[51723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:287
	_go_fuzz_dep_.CoverTab[51724]++

													if tok.Kind() != text.Scalar {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:289
		_go_fuzz_dep_.CoverTab[51729]++
														return protoreflect.Value{}, d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:290
		// _ = "end of CoverTab[51729]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:291
		_go_fuzz_dep_.CoverTab[51730]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:291
		// _ = "end of CoverTab[51730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:291
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:291
	// _ = "end of CoverTab[51724]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:291
	_go_fuzz_dep_.CoverTab[51725]++

													kind := fd.Kind()
													switch kind {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:295
		_go_fuzz_dep_.CoverTab[51731]++
														if b, ok := tok.Bool(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:296
			_go_fuzz_dep_.CoverTab[51743]++
															return protoreflect.ValueOfBool(b), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:297
			// _ = "end of CoverTab[51743]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:298
			_go_fuzz_dep_.CoverTab[51744]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:298
			// _ = "end of CoverTab[51744]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:298
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:298
		// _ = "end of CoverTab[51731]"

	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:300
		_go_fuzz_dep_.CoverTab[51732]++
														if n, ok := tok.Int32(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:301
			_go_fuzz_dep_.CoverTab[51745]++
															return protoreflect.ValueOfInt32(n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:302
			// _ = "end of CoverTab[51745]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:303
			_go_fuzz_dep_.CoverTab[51746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:303
			// _ = "end of CoverTab[51746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:303
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:303
		// _ = "end of CoverTab[51732]"

	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:305
		_go_fuzz_dep_.CoverTab[51733]++
														if n, ok := tok.Int64(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:306
			_go_fuzz_dep_.CoverTab[51747]++
															return protoreflect.ValueOfInt64(n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:307
			// _ = "end of CoverTab[51747]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:308
			_go_fuzz_dep_.CoverTab[51748]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:308
			// _ = "end of CoverTab[51748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:308
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:308
		// _ = "end of CoverTab[51733]"

	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:310
		_go_fuzz_dep_.CoverTab[51734]++
														if n, ok := tok.Uint32(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:311
			_go_fuzz_dep_.CoverTab[51749]++
															return protoreflect.ValueOfUint32(n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:312
			// _ = "end of CoverTab[51749]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:313
			_go_fuzz_dep_.CoverTab[51750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:313
			// _ = "end of CoverTab[51750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:313
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:313
		// _ = "end of CoverTab[51734]"

	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:315
		_go_fuzz_dep_.CoverTab[51735]++
														if n, ok := tok.Uint64(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:316
			_go_fuzz_dep_.CoverTab[51751]++
															return protoreflect.ValueOfUint64(n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:317
			// _ = "end of CoverTab[51751]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:318
			_go_fuzz_dep_.CoverTab[51752]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:318
			// _ = "end of CoverTab[51752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:318
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:318
		// _ = "end of CoverTab[51735]"

	case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:320
		_go_fuzz_dep_.CoverTab[51736]++
														if n, ok := tok.Float32(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:321
			_go_fuzz_dep_.CoverTab[51753]++
															return protoreflect.ValueOfFloat32(n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:322
			// _ = "end of CoverTab[51753]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:323
			_go_fuzz_dep_.CoverTab[51754]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:323
			// _ = "end of CoverTab[51754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:323
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:323
		// _ = "end of CoverTab[51736]"

	case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:325
		_go_fuzz_dep_.CoverTab[51737]++
														if n, ok := tok.Float64(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:326
			_go_fuzz_dep_.CoverTab[51755]++
															return protoreflect.ValueOfFloat64(n), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:327
			// _ = "end of CoverTab[51755]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:328
			_go_fuzz_dep_.CoverTab[51756]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:328
			// _ = "end of CoverTab[51756]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:328
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:328
		// _ = "end of CoverTab[51737]"

	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:330
		_go_fuzz_dep_.CoverTab[51738]++
														if s, ok := tok.String(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:331
			_go_fuzz_dep_.CoverTab[51757]++
															if strs.EnforceUTF8(fd) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:332
				_go_fuzz_dep_.CoverTab[51759]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:332
				return !utf8.ValidString(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:332
				// _ = "end of CoverTab[51759]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:332
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:332
				_go_fuzz_dep_.CoverTab[51760]++
																return protoreflect.Value{}, d.newError(tok.Pos(), "contains invalid UTF-8")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:333
				// _ = "end of CoverTab[51760]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:334
				_go_fuzz_dep_.CoverTab[51761]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:334
				// _ = "end of CoverTab[51761]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:334
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:334
			// _ = "end of CoverTab[51757]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:334
			_go_fuzz_dep_.CoverTab[51758]++
															return protoreflect.ValueOfString(s), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:335
			// _ = "end of CoverTab[51758]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:336
			_go_fuzz_dep_.CoverTab[51762]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:336
			// _ = "end of CoverTab[51762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:336
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:336
		// _ = "end of CoverTab[51738]"

	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:338
		_go_fuzz_dep_.CoverTab[51739]++
														if b, ok := tok.String(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:339
			_go_fuzz_dep_.CoverTab[51763]++
															return protoreflect.ValueOfBytes([]byte(b)), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:340
			// _ = "end of CoverTab[51763]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:341
			_go_fuzz_dep_.CoverTab[51764]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:341
			// _ = "end of CoverTab[51764]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:341
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:341
		// _ = "end of CoverTab[51739]"

	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:343
		_go_fuzz_dep_.CoverTab[51740]++
														if lit, ok := tok.Enum(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:344
			_go_fuzz_dep_.CoverTab[51765]++

															if enumVal := fd.Enum().Values().ByName(protoreflect.Name(lit)); enumVal != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:346
				_go_fuzz_dep_.CoverTab[51766]++
																return protoreflect.ValueOfEnum(enumVal.Number()), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:347
				// _ = "end of CoverTab[51766]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:348
				_go_fuzz_dep_.CoverTab[51767]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:348
				// _ = "end of CoverTab[51767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:348
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:348
			// _ = "end of CoverTab[51765]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:349
			_go_fuzz_dep_.CoverTab[51768]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:349
			// _ = "end of CoverTab[51768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:349
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:349
		// _ = "end of CoverTab[51740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:349
		_go_fuzz_dep_.CoverTab[51741]++
														if num, ok := tok.Int32(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:350
			_go_fuzz_dep_.CoverTab[51769]++
															return protoreflect.ValueOfEnum(protoreflect.EnumNumber(num)), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:351
			// _ = "end of CoverTab[51769]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:352
			_go_fuzz_dep_.CoverTab[51770]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:352
			// _ = "end of CoverTab[51770]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:352
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:352
		// _ = "end of CoverTab[51741]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:354
		_go_fuzz_dep_.CoverTab[51742]++
														panic(fmt.Sprintf("invalid scalar kind %v", kind))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:355
		// _ = "end of CoverTab[51742]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:356
	// _ = "end of CoverTab[51725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:356
	_go_fuzz_dep_.CoverTab[51726]++

													return protoreflect.Value{}, d.newError(tok.Pos(), "invalid value for %v type: %v", kind, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:358
	// _ = "end of CoverTab[51726]"
}

// unmarshalList unmarshals into given protoreflect.List. A list value can
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:361
// either be in [] syntax or simply just a single scalar/message value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:363
func (d decoder) unmarshalList(fd protoreflect.FieldDescriptor, list protoreflect.List) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:363
	_go_fuzz_dep_.CoverTab[51771]++
													tok, err := d.Peek()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:365
		_go_fuzz_dep_.CoverTab[51774]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:366
		// _ = "end of CoverTab[51774]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:367
		_go_fuzz_dep_.CoverTab[51775]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:367
		// _ = "end of CoverTab[51775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:367
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:367
	// _ = "end of CoverTab[51771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:367
	_go_fuzz_dep_.CoverTab[51772]++

													switch fd.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:370
		_go_fuzz_dep_.CoverTab[51776]++
														switch tok.Kind() {
		case text.ListOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:372
			_go_fuzz_dep_.CoverTab[51778]++
															d.Read()
															for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:374
				_go_fuzz_dep_.CoverTab[51782]++
																tok, err := d.Peek()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:376
					_go_fuzz_dep_.CoverTab[51784]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:377
					// _ = "end of CoverTab[51784]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:378
					_go_fuzz_dep_.CoverTab[51785]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:378
					// _ = "end of CoverTab[51785]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:378
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:378
				// _ = "end of CoverTab[51782]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:378
				_go_fuzz_dep_.CoverTab[51783]++

																switch tok.Kind() {
				case text.ListClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:381
					_go_fuzz_dep_.CoverTab[51786]++
																	d.Read()
																	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:383
					// _ = "end of CoverTab[51786]"
				case text.MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:384
					_go_fuzz_dep_.CoverTab[51787]++
																	pval := list.NewElement()
																	if err := d.unmarshalMessage(pval.Message(), true); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:386
						_go_fuzz_dep_.CoverTab[51790]++
																		return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:387
						// _ = "end of CoverTab[51790]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:388
						_go_fuzz_dep_.CoverTab[51791]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:388
						// _ = "end of CoverTab[51791]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:388
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:388
					// _ = "end of CoverTab[51787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:388
					_go_fuzz_dep_.CoverTab[51788]++
																	list.Append(pval)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:389
					// _ = "end of CoverTab[51788]"
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:390
					_go_fuzz_dep_.CoverTab[51789]++
																	return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:391
					// _ = "end of CoverTab[51789]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:392
				// _ = "end of CoverTab[51783]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:393
			// _ = "end of CoverTab[51778]"

		case text.MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:395
			_go_fuzz_dep_.CoverTab[51779]++
															pval := list.NewElement()
															if err := d.unmarshalMessage(pval.Message(), true); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:397
				_go_fuzz_dep_.CoverTab[51792]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:398
				// _ = "end of CoverTab[51792]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:399
				_go_fuzz_dep_.CoverTab[51793]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:399
				// _ = "end of CoverTab[51793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:399
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:399
			// _ = "end of CoverTab[51779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:399
			_go_fuzz_dep_.CoverTab[51780]++
															list.Append(pval)
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:401
			// _ = "end of CoverTab[51780]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:401
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:401
			_go_fuzz_dep_.CoverTab[51781]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:401
			// _ = "end of CoverTab[51781]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:402
		// _ = "end of CoverTab[51776]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:404
		_go_fuzz_dep_.CoverTab[51777]++
														switch tok.Kind() {
		case text.ListOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:406
			_go_fuzz_dep_.CoverTab[51794]++
															d.Read()
															for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:408
				_go_fuzz_dep_.CoverTab[51798]++
																tok, err := d.Peek()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:410
					_go_fuzz_dep_.CoverTab[51800]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:411
					// _ = "end of CoverTab[51800]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:412
					_go_fuzz_dep_.CoverTab[51801]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:412
					// _ = "end of CoverTab[51801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:412
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:412
				// _ = "end of CoverTab[51798]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:412
				_go_fuzz_dep_.CoverTab[51799]++

																switch tok.Kind() {
				case text.ListClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:415
					_go_fuzz_dep_.CoverTab[51802]++
																	d.Read()
																	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:417
					// _ = "end of CoverTab[51802]"
				case text.Scalar:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:418
					_go_fuzz_dep_.CoverTab[51803]++
																	pval, err := d.unmarshalScalar(fd)
																	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:420
						_go_fuzz_dep_.CoverTab[51806]++
																		return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:421
						// _ = "end of CoverTab[51806]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:422
						_go_fuzz_dep_.CoverTab[51807]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:422
						// _ = "end of CoverTab[51807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:422
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:422
					// _ = "end of CoverTab[51803]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:422
					_go_fuzz_dep_.CoverTab[51804]++
																	list.Append(pval)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:423
					// _ = "end of CoverTab[51804]"
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:424
					_go_fuzz_dep_.CoverTab[51805]++
																	return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:425
					// _ = "end of CoverTab[51805]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:426
				// _ = "end of CoverTab[51799]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:427
			// _ = "end of CoverTab[51794]"

		case text.Scalar:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:429
			_go_fuzz_dep_.CoverTab[51795]++
															pval, err := d.unmarshalScalar(fd)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:431
				_go_fuzz_dep_.CoverTab[51808]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:432
				// _ = "end of CoverTab[51808]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:433
				_go_fuzz_dep_.CoverTab[51809]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:433
				// _ = "end of CoverTab[51809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:433
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:433
			// _ = "end of CoverTab[51795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:433
			_go_fuzz_dep_.CoverTab[51796]++
															list.Append(pval)
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:435
			// _ = "end of CoverTab[51796]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:435
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:435
			_go_fuzz_dep_.CoverTab[51797]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:435
			// _ = "end of CoverTab[51797]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:436
		// _ = "end of CoverTab[51777]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:437
	// _ = "end of CoverTab[51772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:437
	_go_fuzz_dep_.CoverTab[51773]++

													return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:439
	// _ = "end of CoverTab[51773]"
}

// unmarshalMap unmarshals into given protoreflect.Map. A map value is a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:442
// textproto message containing {key: <kvalue>, value: <mvalue>}.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:444
func (d decoder) unmarshalMap(fd protoreflect.FieldDescriptor, mmap protoreflect.Map) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:444
	_go_fuzz_dep_.CoverTab[51810]++
	// Determine ahead whether map entry is a scalar type or a message type in
	// order to call the appropriate unmarshalMapValue func inside
	// unmarshalMapEntry.
	var unmarshalMapValue func() (protoreflect.Value, error)
	switch fd.MapValue().Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:450
		_go_fuzz_dep_.CoverTab[51813]++
														unmarshalMapValue = func() (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:451
			_go_fuzz_dep_.CoverTab[51815]++
															pval := mmap.NewValue()
															if err := d.unmarshalMessage(pval.Message(), true); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:453
				_go_fuzz_dep_.CoverTab[51817]++
																return protoreflect.Value{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:454
				// _ = "end of CoverTab[51817]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:455
				_go_fuzz_dep_.CoverTab[51818]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:455
				// _ = "end of CoverTab[51818]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:455
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:455
			// _ = "end of CoverTab[51815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:455
			_go_fuzz_dep_.CoverTab[51816]++
															return pval, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:456
			// _ = "end of CoverTab[51816]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:457
		// _ = "end of CoverTab[51813]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:458
		_go_fuzz_dep_.CoverTab[51814]++
														unmarshalMapValue = func() (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:459
			_go_fuzz_dep_.CoverTab[51819]++
															return d.unmarshalScalar(fd.MapValue())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:460
			// _ = "end of CoverTab[51819]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:461
		// _ = "end of CoverTab[51814]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:462
	// _ = "end of CoverTab[51810]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:462
	_go_fuzz_dep_.CoverTab[51811]++

													tok, err := d.Read()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:465
		_go_fuzz_dep_.CoverTab[51820]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:466
		// _ = "end of CoverTab[51820]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:467
		_go_fuzz_dep_.CoverTab[51821]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:467
		// _ = "end of CoverTab[51821]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:467
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:467
	// _ = "end of CoverTab[51811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:467
	_go_fuzz_dep_.CoverTab[51812]++
													switch tok.Kind() {
	case text.MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:469
		_go_fuzz_dep_.CoverTab[51822]++
														return d.unmarshalMapEntry(fd, mmap, unmarshalMapValue)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:470
		// _ = "end of CoverTab[51822]"

	case text.ListOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:472
		_go_fuzz_dep_.CoverTab[51823]++
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:473
			_go_fuzz_dep_.CoverTab[51825]++
															tok, err := d.Read()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:475
				_go_fuzz_dep_.CoverTab[51827]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:476
				// _ = "end of CoverTab[51827]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:477
				_go_fuzz_dep_.CoverTab[51828]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:477
				// _ = "end of CoverTab[51828]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:477
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:477
			// _ = "end of CoverTab[51825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:477
			_go_fuzz_dep_.CoverTab[51826]++
															switch tok.Kind() {
			case text.ListClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:479
				_go_fuzz_dep_.CoverTab[51829]++
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:480
				// _ = "end of CoverTab[51829]"
			case text.MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:481
				_go_fuzz_dep_.CoverTab[51830]++
																if err := d.unmarshalMapEntry(fd, mmap, unmarshalMapValue); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:482
					_go_fuzz_dep_.CoverTab[51832]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:483
					// _ = "end of CoverTab[51832]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:484
					_go_fuzz_dep_.CoverTab[51833]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:484
					// _ = "end of CoverTab[51833]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:484
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:484
				// _ = "end of CoverTab[51830]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:485
				_go_fuzz_dep_.CoverTab[51831]++
																return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:486
				// _ = "end of CoverTab[51831]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:487
			// _ = "end of CoverTab[51826]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:488
		// _ = "end of CoverTab[51823]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:490
		_go_fuzz_dep_.CoverTab[51824]++
														return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:491
		// _ = "end of CoverTab[51824]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:492
	// _ = "end of CoverTab[51812]"
}

// unmarshalMap unmarshals into given protoreflect.Map. A map value is a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:495
// textproto message containing {key: <kvalue>, value: <mvalue>}.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:497
func (d decoder) unmarshalMapEntry(fd protoreflect.FieldDescriptor, mmap protoreflect.Map, unmarshalMapValue func() (protoreflect.Value, error)) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:497
	_go_fuzz_dep_.CoverTab[51834]++
													var key protoreflect.MapKey
													var pval protoreflect.Value
Loop:
	for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:501
		_go_fuzz_dep_.CoverTab[51838]++

														tok, err := d.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:504
			_go_fuzz_dep_.CoverTab[51841]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:505
			// _ = "end of CoverTab[51841]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:506
			_go_fuzz_dep_.CoverTab[51842]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:506
			// _ = "end of CoverTab[51842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:506
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:506
		// _ = "end of CoverTab[51838]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:506
		_go_fuzz_dep_.CoverTab[51839]++
														switch tok.Kind() {
		case text.Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:508
			_go_fuzz_dep_.CoverTab[51843]++
															if tok.NameKind() != text.IdentName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:509
				_go_fuzz_dep_.CoverTab[51846]++
																if !d.opts.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:510
					_go_fuzz_dep_.CoverTab[51848]++
																	return d.newError(tok.Pos(), "unknown map entry field %q", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:511
					// _ = "end of CoverTab[51848]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:512
					_go_fuzz_dep_.CoverTab[51849]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:512
					// _ = "end of CoverTab[51849]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:512
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:512
				// _ = "end of CoverTab[51846]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:512
				_go_fuzz_dep_.CoverTab[51847]++
																d.skipValue()
																continue Loop
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:514
				// _ = "end of CoverTab[51847]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:515
				_go_fuzz_dep_.CoverTab[51850]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:515
				// _ = "end of CoverTab[51850]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:515
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:515
			// _ = "end of CoverTab[51843]"

		case text.MessageClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:517
			_go_fuzz_dep_.CoverTab[51844]++
															break Loop
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:518
			// _ = "end of CoverTab[51844]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:519
			_go_fuzz_dep_.CoverTab[51845]++
															return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:520
			// _ = "end of CoverTab[51845]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:521
		// _ = "end of CoverTab[51839]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:521
		_go_fuzz_dep_.CoverTab[51840]++

														switch name := protoreflect.Name(tok.IdentName()); name {
		case genid.MapEntry_Key_field_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:524
			_go_fuzz_dep_.CoverTab[51851]++
															if !tok.HasSeparator() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:525
				_go_fuzz_dep_.CoverTab[51860]++
																return d.syntaxError(tok.Pos(), "missing field separator :")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:526
				// _ = "end of CoverTab[51860]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:527
				_go_fuzz_dep_.CoverTab[51861]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:527
				// _ = "end of CoverTab[51861]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:527
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:527
			// _ = "end of CoverTab[51851]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:527
			_go_fuzz_dep_.CoverTab[51852]++
															if key.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:528
				_go_fuzz_dep_.CoverTab[51862]++
																return d.newError(tok.Pos(), "map entry %q cannot be repeated", name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:529
				// _ = "end of CoverTab[51862]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:530
				_go_fuzz_dep_.CoverTab[51863]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:530
				// _ = "end of CoverTab[51863]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:530
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:530
			// _ = "end of CoverTab[51852]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:530
			_go_fuzz_dep_.CoverTab[51853]++
															val, err := d.unmarshalScalar(fd.MapKey())
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:532
				_go_fuzz_dep_.CoverTab[51864]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:533
				// _ = "end of CoverTab[51864]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:534
				_go_fuzz_dep_.CoverTab[51865]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:534
				// _ = "end of CoverTab[51865]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:534
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:534
			// _ = "end of CoverTab[51853]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:534
			_go_fuzz_dep_.CoverTab[51854]++
															key = val.MapKey()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:535
			// _ = "end of CoverTab[51854]"

		case genid.MapEntry_Value_field_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:537
			_go_fuzz_dep_.CoverTab[51855]++
															if kind := fd.MapValue().Kind(); (kind != protoreflect.MessageKind) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:538
				_go_fuzz_dep_.CoverTab[51866]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:538
				return (kind != protoreflect.GroupKind)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:538
				// _ = "end of CoverTab[51866]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:538
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:538
				_go_fuzz_dep_.CoverTab[51867]++
																if !tok.HasSeparator() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:539
					_go_fuzz_dep_.CoverTab[51868]++
																	return d.syntaxError(tok.Pos(), "missing field separator :")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:540
					// _ = "end of CoverTab[51868]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:541
					_go_fuzz_dep_.CoverTab[51869]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:541
					// _ = "end of CoverTab[51869]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:541
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:541
				// _ = "end of CoverTab[51867]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:542
				_go_fuzz_dep_.CoverTab[51870]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:542
				// _ = "end of CoverTab[51870]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:542
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:542
			// _ = "end of CoverTab[51855]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:542
			_go_fuzz_dep_.CoverTab[51856]++
															if pval.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:543
				_go_fuzz_dep_.CoverTab[51871]++
																return d.newError(tok.Pos(), "map entry %q cannot be repeated", name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:544
				// _ = "end of CoverTab[51871]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:545
				_go_fuzz_dep_.CoverTab[51872]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:545
				// _ = "end of CoverTab[51872]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:545
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:545
			// _ = "end of CoverTab[51856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:545
			_go_fuzz_dep_.CoverTab[51857]++
															pval, err = unmarshalMapValue()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:547
				_go_fuzz_dep_.CoverTab[51873]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:548
				// _ = "end of CoverTab[51873]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:549
				_go_fuzz_dep_.CoverTab[51874]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:549
				// _ = "end of CoverTab[51874]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:549
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:549
			// _ = "end of CoverTab[51857]"

		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:551
			_go_fuzz_dep_.CoverTab[51858]++
															if !d.opts.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:552
				_go_fuzz_dep_.CoverTab[51875]++
																return d.newError(tok.Pos(), "unknown map entry field %q", name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:553
				// _ = "end of CoverTab[51875]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:554
				_go_fuzz_dep_.CoverTab[51876]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:554
				// _ = "end of CoverTab[51876]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:554
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:554
			// _ = "end of CoverTab[51858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:554
			_go_fuzz_dep_.CoverTab[51859]++
															d.skipValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:555
			// _ = "end of CoverTab[51859]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:556
		// _ = "end of CoverTab[51840]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:557
	// _ = "end of CoverTab[51834]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:557
	_go_fuzz_dep_.CoverTab[51835]++

													if !key.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:559
		_go_fuzz_dep_.CoverTab[51877]++
														key = fd.MapKey().Default().MapKey()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:560
		// _ = "end of CoverTab[51877]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:561
		_go_fuzz_dep_.CoverTab[51878]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:561
		// _ = "end of CoverTab[51878]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:561
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:561
	// _ = "end of CoverTab[51835]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:561
	_go_fuzz_dep_.CoverTab[51836]++
													if !pval.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:562
		_go_fuzz_dep_.CoverTab[51879]++
														switch fd.MapValue().Kind() {
		case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:564
			_go_fuzz_dep_.CoverTab[51880]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:567
			pval = mmap.NewValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:567
			// _ = "end of CoverTab[51880]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:568
			_go_fuzz_dep_.CoverTab[51881]++
															pval = fd.MapValue().Default()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:569
			// _ = "end of CoverTab[51881]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:570
		// _ = "end of CoverTab[51879]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:571
		_go_fuzz_dep_.CoverTab[51882]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:571
		// _ = "end of CoverTab[51882]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:571
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:571
	// _ = "end of CoverTab[51836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:571
	_go_fuzz_dep_.CoverTab[51837]++
													mmap.Set(key, pval)
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:573
	// _ = "end of CoverTab[51837]"
}

// unmarshalAny unmarshals an Any textproto. It can either be in expanded form
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:576
// or non-expanded form.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:578
func (d decoder) unmarshalAny(m protoreflect.Message, checkDelims bool) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:578
	_go_fuzz_dep_.CoverTab[51883]++
													var typeURL string
													var bValue []byte
													var seenTypeUrl bool
													var seenValue bool
													var isExpanded bool

													if checkDelims {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:585
		_go_fuzz_dep_.CoverTab[51888]++
														tok, err := d.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:587
			_go_fuzz_dep_.CoverTab[51890]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:588
			// _ = "end of CoverTab[51890]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:589
			_go_fuzz_dep_.CoverTab[51891]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:589
			// _ = "end of CoverTab[51891]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:589
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:589
		// _ = "end of CoverTab[51888]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:589
		_go_fuzz_dep_.CoverTab[51889]++

														if tok.Kind() != text.MessageOpen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:591
			_go_fuzz_dep_.CoverTab[51892]++
															return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:592
			// _ = "end of CoverTab[51892]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:593
			_go_fuzz_dep_.CoverTab[51893]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:593
			// _ = "end of CoverTab[51893]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:593
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:593
		// _ = "end of CoverTab[51889]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:594
		_go_fuzz_dep_.CoverTab[51894]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:594
		// _ = "end of CoverTab[51894]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:594
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:594
	// _ = "end of CoverTab[51883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:594
	_go_fuzz_dep_.CoverTab[51884]++

Loop:
	for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:597
		_go_fuzz_dep_.CoverTab[51895]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:600
		tok, err := d.Read()
		if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:601
			_go_fuzz_dep_.CoverTab[51898]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:602
			// _ = "end of CoverTab[51898]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:603
			_go_fuzz_dep_.CoverTab[51899]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:603
			// _ = "end of CoverTab[51899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:603
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:603
		// _ = "end of CoverTab[51895]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:603
		_go_fuzz_dep_.CoverTab[51896]++
														if typ := tok.Kind(); typ != text.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:604
			_go_fuzz_dep_.CoverTab[51900]++
															if checkDelims {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:605
				_go_fuzz_dep_.CoverTab[51902]++
																if typ == text.MessageClose {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:606
					_go_fuzz_dep_.CoverTab[51903]++
																	break Loop
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:607
					// _ = "end of CoverTab[51903]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:608
					_go_fuzz_dep_.CoverTab[51904]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:608
					// _ = "end of CoverTab[51904]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:608
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:608
				// _ = "end of CoverTab[51902]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:609
				_go_fuzz_dep_.CoverTab[51905]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:609
				if typ == text.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:609
					_go_fuzz_dep_.CoverTab[51906]++
																	break Loop
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:610
					// _ = "end of CoverTab[51906]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:611
					_go_fuzz_dep_.CoverTab[51907]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:611
					// _ = "end of CoverTab[51907]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:611
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:611
				// _ = "end of CoverTab[51905]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:611
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:611
			// _ = "end of CoverTab[51900]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:611
			_go_fuzz_dep_.CoverTab[51901]++
															return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:612
			// _ = "end of CoverTab[51901]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:613
			_go_fuzz_dep_.CoverTab[51908]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:613
			// _ = "end of CoverTab[51908]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:613
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:613
		// _ = "end of CoverTab[51896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:613
		_go_fuzz_dep_.CoverTab[51897]++

														switch tok.NameKind() {
		case text.IdentName:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:616
			_go_fuzz_dep_.CoverTab[51909]++

															if !tok.HasSeparator() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:618
				_go_fuzz_dep_.CoverTab[51916]++
																return d.syntaxError(tok.Pos(), "missing field separator :")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:619
				// _ = "end of CoverTab[51916]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:620
				_go_fuzz_dep_.CoverTab[51917]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:620
				// _ = "end of CoverTab[51917]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:620
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:620
			// _ = "end of CoverTab[51909]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:620
			_go_fuzz_dep_.CoverTab[51910]++

															switch name := protoreflect.Name(tok.IdentName()); name {
			case genid.Any_TypeUrl_field_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:623
				_go_fuzz_dep_.CoverTab[51918]++
																if seenTypeUrl {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:624
					_go_fuzz_dep_.CoverTab[51929]++
																	return d.newError(tok.Pos(), "duplicate %v field", genid.Any_TypeUrl_field_fullname)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:625
					// _ = "end of CoverTab[51929]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:626
					_go_fuzz_dep_.CoverTab[51930]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:626
					// _ = "end of CoverTab[51930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:626
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:626
				// _ = "end of CoverTab[51918]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:626
				_go_fuzz_dep_.CoverTab[51919]++
																if isExpanded {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:627
					_go_fuzz_dep_.CoverTab[51931]++
																	return d.newError(tok.Pos(), "conflict with [%s] field", typeURL)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:628
					// _ = "end of CoverTab[51931]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:629
					_go_fuzz_dep_.CoverTab[51932]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:629
					// _ = "end of CoverTab[51932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:629
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:629
				// _ = "end of CoverTab[51919]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:629
				_go_fuzz_dep_.CoverTab[51920]++
																tok, err := d.Read()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:631
					_go_fuzz_dep_.CoverTab[51933]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:632
					// _ = "end of CoverTab[51933]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:633
					_go_fuzz_dep_.CoverTab[51934]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:633
					// _ = "end of CoverTab[51934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:633
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:633
				// _ = "end of CoverTab[51920]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:633
				_go_fuzz_dep_.CoverTab[51921]++
																var ok bool
																typeURL, ok = tok.String()
																if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:636
					_go_fuzz_dep_.CoverTab[51935]++
																	return d.newError(tok.Pos(), "invalid %v field value: %v", genid.Any_TypeUrl_field_fullname, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:637
					// _ = "end of CoverTab[51935]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:638
					_go_fuzz_dep_.CoverTab[51936]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:638
					// _ = "end of CoverTab[51936]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:638
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:638
				// _ = "end of CoverTab[51921]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:638
				_go_fuzz_dep_.CoverTab[51922]++
																seenTypeUrl = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:639
				// _ = "end of CoverTab[51922]"

			case genid.Any_Value_field_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:641
				_go_fuzz_dep_.CoverTab[51923]++
																if seenValue {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:642
					_go_fuzz_dep_.CoverTab[51937]++
																	return d.newError(tok.Pos(), "duplicate %v field", genid.Any_Value_field_fullname)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:643
					// _ = "end of CoverTab[51937]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:644
					_go_fuzz_dep_.CoverTab[51938]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:644
					// _ = "end of CoverTab[51938]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:644
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:644
				// _ = "end of CoverTab[51923]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:644
				_go_fuzz_dep_.CoverTab[51924]++
																if isExpanded {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:645
					_go_fuzz_dep_.CoverTab[51939]++
																	return d.newError(tok.Pos(), "conflict with [%s] field", typeURL)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:646
					// _ = "end of CoverTab[51939]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:647
					_go_fuzz_dep_.CoverTab[51940]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:647
					// _ = "end of CoverTab[51940]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:647
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:647
				// _ = "end of CoverTab[51924]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:647
				_go_fuzz_dep_.CoverTab[51925]++
																tok, err := d.Read()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:649
					_go_fuzz_dep_.CoverTab[51941]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:650
					// _ = "end of CoverTab[51941]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:651
					_go_fuzz_dep_.CoverTab[51942]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:651
					// _ = "end of CoverTab[51942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:651
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:651
				// _ = "end of CoverTab[51925]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:651
				_go_fuzz_dep_.CoverTab[51926]++
																s, ok := tok.String()
																if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:653
					_go_fuzz_dep_.CoverTab[51943]++
																	return d.newError(tok.Pos(), "invalid %v field value: %v", genid.Any_Value_field_fullname, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:654
					// _ = "end of CoverTab[51943]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:655
					_go_fuzz_dep_.CoverTab[51944]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:655
					// _ = "end of CoverTab[51944]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:655
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:655
				// _ = "end of CoverTab[51926]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:655
				_go_fuzz_dep_.CoverTab[51927]++
																bValue = []byte(s)
																seenValue = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:657
				// _ = "end of CoverTab[51927]"

			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:659
				_go_fuzz_dep_.CoverTab[51928]++
																if !d.opts.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:660
					_go_fuzz_dep_.CoverTab[51945]++
																	return d.newError(tok.Pos(), "invalid field name %q in %v message", tok.RawString(), genid.Any_message_fullname)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:661
					// _ = "end of CoverTab[51945]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:662
					_go_fuzz_dep_.CoverTab[51946]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:662
					// _ = "end of CoverTab[51946]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:662
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:662
				// _ = "end of CoverTab[51928]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:663
			// _ = "end of CoverTab[51910]"

		case text.TypeName:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:665
			_go_fuzz_dep_.CoverTab[51911]++
															if isExpanded {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:666
				_go_fuzz_dep_.CoverTab[51947]++
																return d.newError(tok.Pos(), "cannot have more than one type")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:667
				// _ = "end of CoverTab[51947]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:668
				_go_fuzz_dep_.CoverTab[51948]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:668
				// _ = "end of CoverTab[51948]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:668
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:668
			// _ = "end of CoverTab[51911]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:668
			_go_fuzz_dep_.CoverTab[51912]++
															if seenTypeUrl {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:669
				_go_fuzz_dep_.CoverTab[51949]++
																return d.newError(tok.Pos(), "conflict with type_url field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:670
				// _ = "end of CoverTab[51949]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:671
				_go_fuzz_dep_.CoverTab[51950]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:671
				// _ = "end of CoverTab[51950]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:671
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:671
			// _ = "end of CoverTab[51912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:671
			_go_fuzz_dep_.CoverTab[51913]++
															typeURL = tok.TypeName()
															var err error
															bValue, err = d.unmarshalExpandedAny(typeURL, tok.Pos())
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:675
				_go_fuzz_dep_.CoverTab[51951]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:676
				// _ = "end of CoverTab[51951]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:677
				_go_fuzz_dep_.CoverTab[51952]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:677
				// _ = "end of CoverTab[51952]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:677
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:677
			// _ = "end of CoverTab[51913]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:677
			_go_fuzz_dep_.CoverTab[51914]++
															isExpanded = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:678
			// _ = "end of CoverTab[51914]"

		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:680
			_go_fuzz_dep_.CoverTab[51915]++
															if !d.opts.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:681
				_go_fuzz_dep_.CoverTab[51953]++
																return d.newError(tok.Pos(), "invalid field name %q in %v message", tok.RawString(), genid.Any_message_fullname)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:682
				// _ = "end of CoverTab[51953]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:683
				_go_fuzz_dep_.CoverTab[51954]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:683
				// _ = "end of CoverTab[51954]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:683
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:683
			// _ = "end of CoverTab[51915]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:684
		// _ = "end of CoverTab[51897]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:685
	// _ = "end of CoverTab[51884]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:685
	_go_fuzz_dep_.CoverTab[51885]++

													fds := m.Descriptor().Fields()
													if len(typeURL) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:688
		_go_fuzz_dep_.CoverTab[51955]++
														m.Set(fds.ByNumber(genid.Any_TypeUrl_field_number), protoreflect.ValueOfString(typeURL))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:689
		// _ = "end of CoverTab[51955]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:690
		_go_fuzz_dep_.CoverTab[51956]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:690
		// _ = "end of CoverTab[51956]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:690
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:690
	// _ = "end of CoverTab[51885]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:690
	_go_fuzz_dep_.CoverTab[51886]++
													if len(bValue) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:691
		_go_fuzz_dep_.CoverTab[51957]++
														m.Set(fds.ByNumber(genid.Any_Value_field_number), protoreflect.ValueOfBytes(bValue))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:692
		// _ = "end of CoverTab[51957]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:693
		_go_fuzz_dep_.CoverTab[51958]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:693
		// _ = "end of CoverTab[51958]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:693
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:693
	// _ = "end of CoverTab[51886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:693
	_go_fuzz_dep_.CoverTab[51887]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:694
	// _ = "end of CoverTab[51887]"
}

func (d decoder) unmarshalExpandedAny(typeURL string, pos int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:697
	_go_fuzz_dep_.CoverTab[51959]++
													mt, err := d.opts.Resolver.FindMessageByURL(typeURL)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:699
		_go_fuzz_dep_.CoverTab[51963]++
														return nil, d.newError(pos, "unable to resolve message [%v]: %v", typeURL, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:700
		// _ = "end of CoverTab[51963]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:701
		_go_fuzz_dep_.CoverTab[51964]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:701
		// _ = "end of CoverTab[51964]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:701
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:701
	// _ = "end of CoverTab[51959]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:701
	_go_fuzz_dep_.CoverTab[51960]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:704
	m := mt.New()
	if err := d.unmarshalMessage(m, true); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:705
		_go_fuzz_dep_.CoverTab[51965]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:706
		// _ = "end of CoverTab[51965]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:707
		_go_fuzz_dep_.CoverTab[51966]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:707
		// _ = "end of CoverTab[51966]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:707
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:707
	// _ = "end of CoverTab[51960]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:707
	_go_fuzz_dep_.CoverTab[51961]++

													b, err := proto.MarshalOptions{
		AllowPartial:	true,
		Deterministic:	true,
	}.Marshal(m.Interface())
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:713
		_go_fuzz_dep_.CoverTab[51967]++
														return nil, d.newError(pos, "error in marshaling message into Any.value: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:714
		// _ = "end of CoverTab[51967]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:715
		_go_fuzz_dep_.CoverTab[51968]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:715
		// _ = "end of CoverTab[51968]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:715
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:715
	// _ = "end of CoverTab[51961]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:715
	_go_fuzz_dep_.CoverTab[51962]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:716
	// _ = "end of CoverTab[51962]"
}

// skipValue makes the decoder parse a field value in order to advance the read
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:719
// to the next field. It relies on Read returning an error if the types are not
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:719
// in valid sequence.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:722
func (d decoder) skipValue() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:722
	_go_fuzz_dep_.CoverTab[51969]++
													tok, err := d.Read()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:724
		_go_fuzz_dep_.CoverTab[51972]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:725
		// _ = "end of CoverTab[51972]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:726
		_go_fuzz_dep_.CoverTab[51973]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:726
		// _ = "end of CoverTab[51973]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:726
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:726
	// _ = "end of CoverTab[51969]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:726
	_go_fuzz_dep_.CoverTab[51970]++

													switch tok.Kind() {
	case text.MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:729
		_go_fuzz_dep_.CoverTab[51974]++
														return d.skipMessageValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:730
		// _ = "end of CoverTab[51974]"

	case text.ListOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:732
		_go_fuzz_dep_.CoverTab[51975]++
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:733
			_go_fuzz_dep_.CoverTab[51977]++
															tok, err := d.Read()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:735
				_go_fuzz_dep_.CoverTab[51979]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:736
				// _ = "end of CoverTab[51979]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:737
				_go_fuzz_dep_.CoverTab[51980]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:737
				// _ = "end of CoverTab[51980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:737
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:737
			// _ = "end of CoverTab[51977]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:737
			_go_fuzz_dep_.CoverTab[51978]++
															switch tok.Kind() {
			case text.ListClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:739
				_go_fuzz_dep_.CoverTab[51981]++
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:740
				// _ = "end of CoverTab[51981]"
			case text.MessageOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:741
				_go_fuzz_dep_.CoverTab[51982]++
																return d.skipMessageValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:742
				// _ = "end of CoverTab[51982]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:743
				_go_fuzz_dep_.CoverTab[51983]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:743
				// _ = "end of CoverTab[51983]"

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:747
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:747
			// _ = "end of CoverTab[51978]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:748
		// _ = "end of CoverTab[51975]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:748
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:748
		_go_fuzz_dep_.CoverTab[51976]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:748
		// _ = "end of CoverTab[51976]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:749
	// _ = "end of CoverTab[51970]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:749
	_go_fuzz_dep_.CoverTab[51971]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:750
	// _ = "end of CoverTab[51971]"
}

// skipMessageValue makes the decoder parse and skip over all fields in a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:753
// message. It assumes that the previous read type is MessageOpen.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:755
func (d decoder) skipMessageValue() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:755
	_go_fuzz_dep_.CoverTab[51984]++
													for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:756
		_go_fuzz_dep_.CoverTab[51985]++
														tok, err := d.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:758
			_go_fuzz_dep_.CoverTab[51987]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:759
			// _ = "end of CoverTab[51987]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:760
			_go_fuzz_dep_.CoverTab[51988]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:760
			// _ = "end of CoverTab[51988]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:760
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:760
		// _ = "end of CoverTab[51985]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:760
		_go_fuzz_dep_.CoverTab[51986]++
														switch tok.Kind() {
		case text.MessageClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:762
			_go_fuzz_dep_.CoverTab[51989]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:763
			// _ = "end of CoverTab[51989]"
		case text.Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:764
			_go_fuzz_dep_.CoverTab[51990]++
															if err := d.skipValue(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:765
				_go_fuzz_dep_.CoverTab[51992]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:766
				// _ = "end of CoverTab[51992]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:767
				_go_fuzz_dep_.CoverTab[51993]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:767
				// _ = "end of CoverTab[51993]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:767
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:767
			// _ = "end of CoverTab[51990]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:767
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:767
			_go_fuzz_dep_.CoverTab[51991]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:767
			// _ = "end of CoverTab[51991]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:768
		// _ = "end of CoverTab[51986]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:769
	// _ = "end of CoverTab[51984]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:770
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/decode.go:770
var _ = _go_fuzz_dep_.CoverTab
