// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:5
package protojson

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:5
)

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"strings"

	"google.golang.org/protobuf/internal/encoding/json"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/internal/set"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Unmarshal reads the given []byte into the given proto.Message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:26
// The provided message must be mutable (e.g., a non-nil pointer to a message).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:28
func Unmarshal(b []byte, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:28
	_go_fuzz_dep_.CoverTab[65910]++
													return UnmarshalOptions{}.Unmarshal(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:29
	// _ = "end of CoverTab[65910]"
}

// UnmarshalOptions is a configurable JSON format parser.
type UnmarshalOptions struct {
	pragma.NoUnkeyedLiterals

	// If AllowPartial is set, input for messages that will result in missing
	// required fields will not return an error.
	AllowPartial	bool

	// If DiscardUnknown is set, unknown fields are ignored.
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
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:52
// using options in the UnmarshalOptions object.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:52
// It will clear the message first before setting the fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:52
// If it returns an error, the given message may be partially set.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:52
// The provided message must be mutable (e.g., a non-nil pointer to a message).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:57
func (o UnmarshalOptions) Unmarshal(b []byte, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:57
	_go_fuzz_dep_.CoverTab[65911]++
													return o.unmarshal(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:58
	// _ = "end of CoverTab[65911]"
}

// unmarshal is a centralized function that all unmarshal operations go through.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:61
// For profiling purposes, avoid changing the name of this function or
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:61
// introducing other code paths for unmarshal that do not go through this.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:64
func (o UnmarshalOptions) unmarshal(b []byte, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:64
	_go_fuzz_dep_.CoverTab[65912]++
													proto.Reset(m)

													if o.Resolver == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:67
		_go_fuzz_dep_.CoverTab[65918]++
														o.Resolver = protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:68
		// _ = "end of CoverTab[65918]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:69
		_go_fuzz_dep_.CoverTab[65919]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:69
		// _ = "end of CoverTab[65919]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:69
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:69
	// _ = "end of CoverTab[65912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:69
	_go_fuzz_dep_.CoverTab[65913]++

													dec := decoder{json.NewDecoder(b), o}
													if err := dec.unmarshalMessage(m.ProtoReflect(), false); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:72
		_go_fuzz_dep_.CoverTab[65920]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:73
		// _ = "end of CoverTab[65920]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:74
		_go_fuzz_dep_.CoverTab[65921]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:74
		// _ = "end of CoverTab[65921]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:74
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:74
	// _ = "end of CoverTab[65913]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:74
	_go_fuzz_dep_.CoverTab[65914]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:77
	tok, err := dec.Read()
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:78
		_go_fuzz_dep_.CoverTab[65922]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:79
		// _ = "end of CoverTab[65922]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:80
		_go_fuzz_dep_.CoverTab[65923]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:80
		// _ = "end of CoverTab[65923]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:80
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:80
	// _ = "end of CoverTab[65914]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:80
	_go_fuzz_dep_.CoverTab[65915]++
													if tok.Kind() != json.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:81
		_go_fuzz_dep_.CoverTab[65924]++
														return dec.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:82
		// _ = "end of CoverTab[65924]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:83
		_go_fuzz_dep_.CoverTab[65925]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:83
		// _ = "end of CoverTab[65925]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:83
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:83
	// _ = "end of CoverTab[65915]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:83
	_go_fuzz_dep_.CoverTab[65916]++

													if o.AllowPartial {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:85
		_go_fuzz_dep_.CoverTab[65926]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:86
		// _ = "end of CoverTab[65926]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:87
		_go_fuzz_dep_.CoverTab[65927]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:87
		// _ = "end of CoverTab[65927]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:87
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:87
	// _ = "end of CoverTab[65916]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:87
	_go_fuzz_dep_.CoverTab[65917]++
													return proto.CheckInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:88
	// _ = "end of CoverTab[65917]"
}

type decoder struct {
	*json.Decoder
	opts	UnmarshalOptions
}

// newError returns an error object with position info.
func (d decoder) newError(pos int, f string, x ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:97
	_go_fuzz_dep_.CoverTab[65928]++
													line, column := d.Position(pos)
													head := fmt.Sprintf("(line %d:%d): ", line, column)
													return errors.New(head+f, x...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:100
	// _ = "end of CoverTab[65928]"
}

// unexpectedTokenError returns a syntax error for the given unexpected token.
func (d decoder) unexpectedTokenError(tok json.Token) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:104
	_go_fuzz_dep_.CoverTab[65929]++
													return d.syntaxError(tok.Pos(), "unexpected token %s", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:105
	// _ = "end of CoverTab[65929]"
}

// syntaxError returns a syntax error for given position.
func (d decoder) syntaxError(pos int, f string, x ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:109
	_go_fuzz_dep_.CoverTab[65930]++
													line, column := d.Position(pos)
													head := fmt.Sprintf("syntax error (line %d:%d): ", line, column)
													return errors.New(head+f, x...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:112
	// _ = "end of CoverTab[65930]"
}

// unmarshalMessage unmarshals a message into the given protoreflect.Message.
func (d decoder) unmarshalMessage(m protoreflect.Message, skipTypeURL bool) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:116
	_go_fuzz_dep_.CoverTab[65931]++
													if unmarshal := wellKnownTypeUnmarshaler(m.Descriptor().FullName()); unmarshal != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:117
		_go_fuzz_dep_.CoverTab[65936]++
														return unmarshal(d, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:118
		// _ = "end of CoverTab[65936]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:119
		_go_fuzz_dep_.CoverTab[65937]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:119
		// _ = "end of CoverTab[65937]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:119
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:119
	// _ = "end of CoverTab[65931]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:119
	_go_fuzz_dep_.CoverTab[65932]++

													tok, err := d.Read()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:122
		_go_fuzz_dep_.CoverTab[65938]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:123
		// _ = "end of CoverTab[65938]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:124
		_go_fuzz_dep_.CoverTab[65939]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:124
		// _ = "end of CoverTab[65939]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:124
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:124
	// _ = "end of CoverTab[65932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:124
	_go_fuzz_dep_.CoverTab[65933]++
													if tok.Kind() != json.ObjectOpen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:125
		_go_fuzz_dep_.CoverTab[65940]++
														return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:126
		// _ = "end of CoverTab[65940]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:127
		_go_fuzz_dep_.CoverTab[65941]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:127
		// _ = "end of CoverTab[65941]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:127
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:127
	// _ = "end of CoverTab[65933]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:127
	_go_fuzz_dep_.CoverTab[65934]++

													messageDesc := m.Descriptor()
													if !flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:130
		_go_fuzz_dep_.CoverTab[65942]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:130
		return messageset.IsMessageSet(messageDesc)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:130
		// _ = "end of CoverTab[65942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:130
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:130
		_go_fuzz_dep_.CoverTab[65943]++
														return errors.New("no support for proto1 MessageSets")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:131
		// _ = "end of CoverTab[65943]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:132
		_go_fuzz_dep_.CoverTab[65944]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:132
		// _ = "end of CoverTab[65944]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:132
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:132
	// _ = "end of CoverTab[65934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:132
	_go_fuzz_dep_.CoverTab[65935]++

													var seenNums set.Ints
													var seenOneofs set.Ints
													fieldDescs := messageDesc.Fields()
													for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:137
		_go_fuzz_dep_.CoverTab[65945]++

														tok, err := d.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:140
			_go_fuzz_dep_.CoverTab[65954]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:141
			// _ = "end of CoverTab[65954]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:142
			_go_fuzz_dep_.CoverTab[65955]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:142
			// _ = "end of CoverTab[65955]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:142
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:142
		// _ = "end of CoverTab[65945]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:142
		_go_fuzz_dep_.CoverTab[65946]++
														switch tok.Kind() {
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:144
			_go_fuzz_dep_.CoverTab[65956]++
															return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:145
			// _ = "end of CoverTab[65956]"
		case json.ObjectClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:146
			_go_fuzz_dep_.CoverTab[65957]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:147
			// _ = "end of CoverTab[65957]"
		case json.Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:148
			_go_fuzz_dep_.CoverTab[65958]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:148
			// _ = "end of CoverTab[65958]"

		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:150
		// _ = "end of CoverTab[65946]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:150
		_go_fuzz_dep_.CoverTab[65947]++

														name := tok.Name()

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:156
		if skipTypeURL && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:156
			_go_fuzz_dep_.CoverTab[65959]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:156
			return name == "@type"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:156
			// _ = "end of CoverTab[65959]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:156
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:156
			_go_fuzz_dep_.CoverTab[65960]++
															d.Read()
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:158
			// _ = "end of CoverTab[65960]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:159
			_go_fuzz_dep_.CoverTab[65961]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:159
			// _ = "end of CoverTab[65961]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:159
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:159
		// _ = "end of CoverTab[65947]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:159
		_go_fuzz_dep_.CoverTab[65948]++

		// Get the FieldDescriptor.
		var fd protoreflect.FieldDescriptor
		if strings.HasPrefix(name, "[") && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:163
			_go_fuzz_dep_.CoverTab[65962]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:163
			return strings.HasSuffix(name, "]")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:163
			// _ = "end of CoverTab[65962]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:163
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:163
			_go_fuzz_dep_.CoverTab[65963]++

															extName := protoreflect.FullName(name[1 : len(name)-1])
															extType, err := d.opts.Resolver.FindExtensionByName(extName)
															if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:167
				_go_fuzz_dep_.CoverTab[65965]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:167
				return err != protoregistry.NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:167
				// _ = "end of CoverTab[65965]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:167
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:167
				_go_fuzz_dep_.CoverTab[65966]++
																return d.newError(tok.Pos(), "unable to resolve %s: %v", tok.RawString(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:168
				// _ = "end of CoverTab[65966]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:169
				_go_fuzz_dep_.CoverTab[65967]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:169
				// _ = "end of CoverTab[65967]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:169
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:169
			// _ = "end of CoverTab[65963]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:169
			_go_fuzz_dep_.CoverTab[65964]++
															if extType != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:170
				_go_fuzz_dep_.CoverTab[65968]++
																fd = extType.TypeDescriptor()
																if !messageDesc.ExtensionRanges().Has(fd.Number()) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:172
					_go_fuzz_dep_.CoverTab[65969]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:172
					return fd.ContainingMessage().FullName() != messageDesc.FullName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:172
					// _ = "end of CoverTab[65969]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:172
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:172
					_go_fuzz_dep_.CoverTab[65970]++
																	return d.newError(tok.Pos(), "message %v cannot be extended by %v", messageDesc.FullName(), fd.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:173
					// _ = "end of CoverTab[65970]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:174
					_go_fuzz_dep_.CoverTab[65971]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:174
					// _ = "end of CoverTab[65971]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:174
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:174
				// _ = "end of CoverTab[65968]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:175
				_go_fuzz_dep_.CoverTab[65972]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:175
				// _ = "end of CoverTab[65972]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:175
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:175
			// _ = "end of CoverTab[65964]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:176
			_go_fuzz_dep_.CoverTab[65973]++

															fd = fieldDescs.ByJSONName(name)
															if fd == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:179
				_go_fuzz_dep_.CoverTab[65974]++
																fd = fieldDescs.ByTextName(name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:180
				// _ = "end of CoverTab[65974]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:181
				_go_fuzz_dep_.CoverTab[65975]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:181
				// _ = "end of CoverTab[65975]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:181
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:181
			// _ = "end of CoverTab[65973]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:182
		// _ = "end of CoverTab[65948]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:182
		_go_fuzz_dep_.CoverTab[65949]++
														if flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:183
			_go_fuzz_dep_.CoverTab[65976]++
															if fd != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:184
				_go_fuzz_dep_.CoverTab[65977]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:184
				return fd.IsWeak()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:184
				// _ = "end of CoverTab[65977]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:184
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:184
				_go_fuzz_dep_.CoverTab[65978]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:184
				return fd.Message().IsPlaceholder()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:184
				// _ = "end of CoverTab[65978]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:184
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:184
				_go_fuzz_dep_.CoverTab[65979]++
																fd = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:185
				// _ = "end of CoverTab[65979]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:186
				_go_fuzz_dep_.CoverTab[65980]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:186
				// _ = "end of CoverTab[65980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:186
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:186
			// _ = "end of CoverTab[65976]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:187
			_go_fuzz_dep_.CoverTab[65981]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:187
			// _ = "end of CoverTab[65981]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:187
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:187
		// _ = "end of CoverTab[65949]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:187
		_go_fuzz_dep_.CoverTab[65950]++

														if fd == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:189
			_go_fuzz_dep_.CoverTab[65982]++

															if d.opts.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:191
				_go_fuzz_dep_.CoverTab[65984]++
																if err := d.skipJSONValue(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:192
					_go_fuzz_dep_.CoverTab[65986]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:193
					// _ = "end of CoverTab[65986]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:194
					_go_fuzz_dep_.CoverTab[65987]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:194
					// _ = "end of CoverTab[65987]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:194
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:194
				// _ = "end of CoverTab[65984]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:194
				_go_fuzz_dep_.CoverTab[65985]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:195
				// _ = "end of CoverTab[65985]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:196
				_go_fuzz_dep_.CoverTab[65988]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:196
				// _ = "end of CoverTab[65988]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:196
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:196
			// _ = "end of CoverTab[65982]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:196
			_go_fuzz_dep_.CoverTab[65983]++
															return d.newError(tok.Pos(), "unknown field %v", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:197
			// _ = "end of CoverTab[65983]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:198
			_go_fuzz_dep_.CoverTab[65989]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:198
			// _ = "end of CoverTab[65989]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:198
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:198
		// _ = "end of CoverTab[65950]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:198
		_go_fuzz_dep_.CoverTab[65951]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:201
		num := uint64(fd.Number())
		if seenNums.Has(num) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:202
			_go_fuzz_dep_.CoverTab[65990]++
															return d.newError(tok.Pos(), "duplicate field %v", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:203
			// _ = "end of CoverTab[65990]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:204
			_go_fuzz_dep_.CoverTab[65991]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:204
			// _ = "end of CoverTab[65991]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:204
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:204
		// _ = "end of CoverTab[65951]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:204
		_go_fuzz_dep_.CoverTab[65952]++
														seenNums.Set(num)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
		if tok, _ := d.Peek(); tok.Kind() == json.Null && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
			_go_fuzz_dep_.CoverTab[65992]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
			return !isKnownValue(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
			// _ = "end of CoverTab[65992]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
			_go_fuzz_dep_.CoverTab[65993]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
			return !isNullValue(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
			// _ = "end of CoverTab[65993]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:209
			_go_fuzz_dep_.CoverTab[65994]++
															d.Read()
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:211
			// _ = "end of CoverTab[65994]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:212
			_go_fuzz_dep_.CoverTab[65995]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:212
			// _ = "end of CoverTab[65995]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:212
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:212
		// _ = "end of CoverTab[65952]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:212
		_go_fuzz_dep_.CoverTab[65953]++

														switch {
		case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:215
			_go_fuzz_dep_.CoverTab[65996]++
															list := m.Mutable(fd).List()
															if err := d.unmarshalList(list, fd); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:217
				_go_fuzz_dep_.CoverTab[66000]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:218
				// _ = "end of CoverTab[66000]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:219
				_go_fuzz_dep_.CoverTab[66001]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:219
				// _ = "end of CoverTab[66001]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:219
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:219
			// _ = "end of CoverTab[65996]"
		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:220
			_go_fuzz_dep_.CoverTab[65997]++
															mmap := m.Mutable(fd).Map()
															if err := d.unmarshalMap(mmap, fd); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:222
				_go_fuzz_dep_.CoverTab[66002]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:223
				// _ = "end of CoverTab[66002]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:224
				_go_fuzz_dep_.CoverTab[66003]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:224
				// _ = "end of CoverTab[66003]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:224
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:224
			// _ = "end of CoverTab[65997]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:225
			_go_fuzz_dep_.CoverTab[65998]++

															if od := fd.ContainingOneof(); od != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:227
				_go_fuzz_dep_.CoverTab[66004]++
																idx := uint64(od.Index())
																if seenOneofs.Has(idx) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:229
					_go_fuzz_dep_.CoverTab[66006]++
																	return d.newError(tok.Pos(), "error parsing %s, oneof %v is already set", tok.RawString(), od.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:230
					// _ = "end of CoverTab[66006]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:231
					_go_fuzz_dep_.CoverTab[66007]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:231
					// _ = "end of CoverTab[66007]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:231
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:231
				// _ = "end of CoverTab[66004]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:231
				_go_fuzz_dep_.CoverTab[66005]++
																seenOneofs.Set(idx)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:232
				// _ = "end of CoverTab[66005]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:233
				_go_fuzz_dep_.CoverTab[66008]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:233
				// _ = "end of CoverTab[66008]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:233
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:233
			// _ = "end of CoverTab[65998]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:233
			_go_fuzz_dep_.CoverTab[65999]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:236
			if err := d.unmarshalSingular(m, fd); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:236
				_go_fuzz_dep_.CoverTab[66009]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:237
				// _ = "end of CoverTab[66009]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:238
				_go_fuzz_dep_.CoverTab[66010]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:238
				// _ = "end of CoverTab[66010]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:238
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:238
			// _ = "end of CoverTab[65999]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:239
		// _ = "end of CoverTab[65953]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:240
	// _ = "end of CoverTab[65935]"
}

func isKnownValue(fd protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:243
	_go_fuzz_dep_.CoverTab[66011]++
													md := fd.Message()
													return md != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:245
		_go_fuzz_dep_.CoverTab[66012]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:245
		return md.FullName() == genid.Value_message_fullname
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:245
		// _ = "end of CoverTab[66012]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:245
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:245
	// _ = "end of CoverTab[66011]"
}

func isNullValue(fd protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:248
	_go_fuzz_dep_.CoverTab[66013]++
													ed := fd.Enum()
													return ed != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:250
		_go_fuzz_dep_.CoverTab[66014]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:250
		return ed.FullName() == genid.NullValue_enum_fullname
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:250
		// _ = "end of CoverTab[66014]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:250
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:250
	// _ = "end of CoverTab[66013]"
}

// unmarshalSingular unmarshals to the non-repeated field specified
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:253
// by the given FieldDescriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:255
func (d decoder) unmarshalSingular(m protoreflect.Message, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:255
	_go_fuzz_dep_.CoverTab[66015]++
													var val protoreflect.Value
													var err error
													switch fd.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:259
		_go_fuzz_dep_.CoverTab[66018]++
														val = m.NewField(fd)
														err = d.unmarshalMessage(val.Message(), false)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:261
		// _ = "end of CoverTab[66018]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:262
		_go_fuzz_dep_.CoverTab[66019]++
														val, err = d.unmarshalScalar(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:263
		// _ = "end of CoverTab[66019]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:264
	// _ = "end of CoverTab[66015]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:264
	_go_fuzz_dep_.CoverTab[66016]++

													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:266
		_go_fuzz_dep_.CoverTab[66020]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:267
		// _ = "end of CoverTab[66020]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:268
		_go_fuzz_dep_.CoverTab[66021]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:268
		// _ = "end of CoverTab[66021]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:268
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:268
	// _ = "end of CoverTab[66016]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:268
	_go_fuzz_dep_.CoverTab[66017]++
													m.Set(fd, val)
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:270
	// _ = "end of CoverTab[66017]"
}

// unmarshalScalar unmarshals to a scalar/enum protoreflect.Value specified by
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:273
// the given FieldDescriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:275
func (d decoder) unmarshalScalar(fd protoreflect.FieldDescriptor) (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:275
	_go_fuzz_dep_.CoverTab[66022]++
													const b32 int = 32
													const b64 int = 64

													tok, err := d.Read()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:280
		_go_fuzz_dep_.CoverTab[66025]++
														return protoreflect.Value{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:281
		// _ = "end of CoverTab[66025]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:282
		_go_fuzz_dep_.CoverTab[66026]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:282
		// _ = "end of CoverTab[66026]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:282
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:282
	// _ = "end of CoverTab[66022]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:282
	_go_fuzz_dep_.CoverTab[66023]++

													kind := fd.Kind()
													switch kind {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:286
		_go_fuzz_dep_.CoverTab[66027]++
														if tok.Kind() == json.Bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:287
			_go_fuzz_dep_.CoverTab[66038]++
															return protoreflect.ValueOfBool(tok.Bool()), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:288
			// _ = "end of CoverTab[66038]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:289
			_go_fuzz_dep_.CoverTab[66039]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:289
			// _ = "end of CoverTab[66039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:289
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:289
		// _ = "end of CoverTab[66027]"

	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:291
		_go_fuzz_dep_.CoverTab[66028]++
														if v, ok := unmarshalInt(tok, b32); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:292
			_go_fuzz_dep_.CoverTab[66040]++
															return v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:293
			// _ = "end of CoverTab[66040]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:294
			_go_fuzz_dep_.CoverTab[66041]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:294
			// _ = "end of CoverTab[66041]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:294
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:294
		// _ = "end of CoverTab[66028]"

	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:296
		_go_fuzz_dep_.CoverTab[66029]++
														if v, ok := unmarshalInt(tok, b64); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:297
			_go_fuzz_dep_.CoverTab[66042]++
															return v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:298
			// _ = "end of CoverTab[66042]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:299
			_go_fuzz_dep_.CoverTab[66043]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:299
			// _ = "end of CoverTab[66043]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:299
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:299
		// _ = "end of CoverTab[66029]"

	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:301
		_go_fuzz_dep_.CoverTab[66030]++
														if v, ok := unmarshalUint(tok, b32); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:302
			_go_fuzz_dep_.CoverTab[66044]++
															return v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:303
			// _ = "end of CoverTab[66044]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:304
			_go_fuzz_dep_.CoverTab[66045]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:304
			// _ = "end of CoverTab[66045]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:304
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:304
		// _ = "end of CoverTab[66030]"

	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:306
		_go_fuzz_dep_.CoverTab[66031]++
														if v, ok := unmarshalUint(tok, b64); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:307
			_go_fuzz_dep_.CoverTab[66046]++
															return v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:308
			// _ = "end of CoverTab[66046]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:309
			_go_fuzz_dep_.CoverTab[66047]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:309
			// _ = "end of CoverTab[66047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:309
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:309
		// _ = "end of CoverTab[66031]"

	case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:311
		_go_fuzz_dep_.CoverTab[66032]++
														if v, ok := unmarshalFloat(tok, b32); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:312
			_go_fuzz_dep_.CoverTab[66048]++
															return v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:313
			// _ = "end of CoverTab[66048]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:314
			_go_fuzz_dep_.CoverTab[66049]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:314
			// _ = "end of CoverTab[66049]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:314
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:314
		// _ = "end of CoverTab[66032]"

	case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:316
		_go_fuzz_dep_.CoverTab[66033]++
														if v, ok := unmarshalFloat(tok, b64); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:317
			_go_fuzz_dep_.CoverTab[66050]++
															return v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:318
			// _ = "end of CoverTab[66050]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:319
			_go_fuzz_dep_.CoverTab[66051]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:319
			// _ = "end of CoverTab[66051]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:319
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:319
		// _ = "end of CoverTab[66033]"

	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:321
		_go_fuzz_dep_.CoverTab[66034]++
														if tok.Kind() == json.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:322
			_go_fuzz_dep_.CoverTab[66052]++
															return protoreflect.ValueOfString(tok.ParsedString()), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:323
			// _ = "end of CoverTab[66052]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:324
			_go_fuzz_dep_.CoverTab[66053]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:324
			// _ = "end of CoverTab[66053]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:324
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:324
		// _ = "end of CoverTab[66034]"

	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:326
		_go_fuzz_dep_.CoverTab[66035]++
														if v, ok := unmarshalBytes(tok); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:327
			_go_fuzz_dep_.CoverTab[66054]++
															return v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:328
			// _ = "end of CoverTab[66054]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:329
			_go_fuzz_dep_.CoverTab[66055]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:329
			// _ = "end of CoverTab[66055]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:329
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:329
		// _ = "end of CoverTab[66035]"

	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:331
		_go_fuzz_dep_.CoverTab[66036]++
														if v, ok := unmarshalEnum(tok, fd); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:332
			_go_fuzz_dep_.CoverTab[66056]++
															return v, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:333
			// _ = "end of CoverTab[66056]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:334
			_go_fuzz_dep_.CoverTab[66057]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:334
			// _ = "end of CoverTab[66057]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:334
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:334
		// _ = "end of CoverTab[66036]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:336
		_go_fuzz_dep_.CoverTab[66037]++
														panic(fmt.Sprintf("unmarshalScalar: invalid scalar kind %v", kind))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:337
		// _ = "end of CoverTab[66037]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:338
	// _ = "end of CoverTab[66023]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:338
	_go_fuzz_dep_.CoverTab[66024]++

													return protoreflect.Value{}, d.newError(tok.Pos(), "invalid value for %v type: %v", kind, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:340
	// _ = "end of CoverTab[66024]"
}

func unmarshalInt(tok json.Token, bitSize int) (protoreflect.Value, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:343
	_go_fuzz_dep_.CoverTab[66058]++
													switch tok.Kind() {
	case json.Number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:345
		_go_fuzz_dep_.CoverTab[66060]++
														return getInt(tok, bitSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:346
		// _ = "end of CoverTab[66060]"

	case json.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:348
		_go_fuzz_dep_.CoverTab[66061]++

														s := strings.TrimSpace(tok.ParsedString())
														if len(s) != len(tok.ParsedString()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:351
			_go_fuzz_dep_.CoverTab[66065]++
															return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:352
			// _ = "end of CoverTab[66065]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:353
			_go_fuzz_dep_.CoverTab[66066]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:353
			// _ = "end of CoverTab[66066]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:353
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:353
		// _ = "end of CoverTab[66061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:353
		_go_fuzz_dep_.CoverTab[66062]++
														dec := json.NewDecoder([]byte(s))
														tok, err := dec.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:356
			_go_fuzz_dep_.CoverTab[66067]++
															return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:357
			// _ = "end of CoverTab[66067]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:358
			_go_fuzz_dep_.CoverTab[66068]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:358
			// _ = "end of CoverTab[66068]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:358
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:358
		// _ = "end of CoverTab[66062]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:358
		_go_fuzz_dep_.CoverTab[66063]++
														return getInt(tok, bitSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:359
		// _ = "end of CoverTab[66063]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:359
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:359
		_go_fuzz_dep_.CoverTab[66064]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:359
		// _ = "end of CoverTab[66064]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:360
	// _ = "end of CoverTab[66058]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:360
	_go_fuzz_dep_.CoverTab[66059]++
													return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:361
	// _ = "end of CoverTab[66059]"
}

func getInt(tok json.Token, bitSize int) (protoreflect.Value, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:364
	_go_fuzz_dep_.CoverTab[66069]++
													n, ok := tok.Int(bitSize)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:366
		_go_fuzz_dep_.CoverTab[66072]++
														return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:367
		// _ = "end of CoverTab[66072]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:368
		_go_fuzz_dep_.CoverTab[66073]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:368
		// _ = "end of CoverTab[66073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:368
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:368
	// _ = "end of CoverTab[66069]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:368
	_go_fuzz_dep_.CoverTab[66070]++
													if bitSize == 32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:369
		_go_fuzz_dep_.CoverTab[66074]++
														return protoreflect.ValueOfInt32(int32(n)), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:370
		// _ = "end of CoverTab[66074]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:371
		_go_fuzz_dep_.CoverTab[66075]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:371
		// _ = "end of CoverTab[66075]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:371
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:371
	// _ = "end of CoverTab[66070]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:371
	_go_fuzz_dep_.CoverTab[66071]++
													return protoreflect.ValueOfInt64(n), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:372
	// _ = "end of CoverTab[66071]"
}

func unmarshalUint(tok json.Token, bitSize int) (protoreflect.Value, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:375
	_go_fuzz_dep_.CoverTab[66076]++
													switch tok.Kind() {
	case json.Number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:377
		_go_fuzz_dep_.CoverTab[66078]++
														return getUint(tok, bitSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:378
		// _ = "end of CoverTab[66078]"

	case json.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:380
		_go_fuzz_dep_.CoverTab[66079]++

														s := strings.TrimSpace(tok.ParsedString())
														if len(s) != len(tok.ParsedString()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:383
			_go_fuzz_dep_.CoverTab[66083]++
															return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:384
			// _ = "end of CoverTab[66083]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:385
			_go_fuzz_dep_.CoverTab[66084]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:385
			// _ = "end of CoverTab[66084]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:385
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:385
		// _ = "end of CoverTab[66079]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:385
		_go_fuzz_dep_.CoverTab[66080]++
														dec := json.NewDecoder([]byte(s))
														tok, err := dec.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:388
			_go_fuzz_dep_.CoverTab[66085]++
															return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:389
			// _ = "end of CoverTab[66085]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:390
			_go_fuzz_dep_.CoverTab[66086]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:390
			// _ = "end of CoverTab[66086]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:390
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:390
		// _ = "end of CoverTab[66080]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:390
		_go_fuzz_dep_.CoverTab[66081]++
														return getUint(tok, bitSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:391
		// _ = "end of CoverTab[66081]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:391
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:391
		_go_fuzz_dep_.CoverTab[66082]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:391
		// _ = "end of CoverTab[66082]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:392
	// _ = "end of CoverTab[66076]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:392
	_go_fuzz_dep_.CoverTab[66077]++
													return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:393
	// _ = "end of CoverTab[66077]"
}

func getUint(tok json.Token, bitSize int) (protoreflect.Value, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:396
	_go_fuzz_dep_.CoverTab[66087]++
													n, ok := tok.Uint(bitSize)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:398
		_go_fuzz_dep_.CoverTab[66090]++
														return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:399
		// _ = "end of CoverTab[66090]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:400
		_go_fuzz_dep_.CoverTab[66091]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:400
		// _ = "end of CoverTab[66091]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:400
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:400
	// _ = "end of CoverTab[66087]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:400
	_go_fuzz_dep_.CoverTab[66088]++
													if bitSize == 32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:401
		_go_fuzz_dep_.CoverTab[66092]++
														return protoreflect.ValueOfUint32(uint32(n)), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:402
		// _ = "end of CoverTab[66092]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:403
		_go_fuzz_dep_.CoverTab[66093]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:403
		// _ = "end of CoverTab[66093]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:403
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:403
	// _ = "end of CoverTab[66088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:403
	_go_fuzz_dep_.CoverTab[66089]++
													return protoreflect.ValueOfUint64(n), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:404
	// _ = "end of CoverTab[66089]"
}

func unmarshalFloat(tok json.Token, bitSize int) (protoreflect.Value, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:407
	_go_fuzz_dep_.CoverTab[66094]++
													switch tok.Kind() {
	case json.Number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:409
		_go_fuzz_dep_.CoverTab[66096]++
														return getFloat(tok, bitSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:410
		// _ = "end of CoverTab[66096]"

	case json.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:412
		_go_fuzz_dep_.CoverTab[66097]++
														s := tok.ParsedString()
														switch s {
		case "NaN":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:415
			_go_fuzz_dep_.CoverTab[66102]++
															if bitSize == 32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:416
				_go_fuzz_dep_.CoverTab[66109]++
																return protoreflect.ValueOfFloat32(float32(math.NaN())), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:417
				// _ = "end of CoverTab[66109]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:418
				_go_fuzz_dep_.CoverTab[66110]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:418
				// _ = "end of CoverTab[66110]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:418
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:418
			// _ = "end of CoverTab[66102]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:418
			_go_fuzz_dep_.CoverTab[66103]++
															return protoreflect.ValueOfFloat64(math.NaN()), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:419
			// _ = "end of CoverTab[66103]"
		case "Infinity":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:420
			_go_fuzz_dep_.CoverTab[66104]++
															if bitSize == 32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:421
				_go_fuzz_dep_.CoverTab[66111]++
																return protoreflect.ValueOfFloat32(float32(math.Inf(+1))), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:422
				// _ = "end of CoverTab[66111]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:423
				_go_fuzz_dep_.CoverTab[66112]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:423
				// _ = "end of CoverTab[66112]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:423
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:423
			// _ = "end of CoverTab[66104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:423
			_go_fuzz_dep_.CoverTab[66105]++
															return protoreflect.ValueOfFloat64(math.Inf(+1)), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:424
			// _ = "end of CoverTab[66105]"
		case "-Infinity":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:425
			_go_fuzz_dep_.CoverTab[66106]++
															if bitSize == 32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:426
				_go_fuzz_dep_.CoverTab[66113]++
																return protoreflect.ValueOfFloat32(float32(math.Inf(-1))), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:427
				// _ = "end of CoverTab[66113]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:428
				_go_fuzz_dep_.CoverTab[66114]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:428
				// _ = "end of CoverTab[66114]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:428
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:428
			// _ = "end of CoverTab[66106]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:428
			_go_fuzz_dep_.CoverTab[66107]++
															return protoreflect.ValueOfFloat64(math.Inf(-1)), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:429
			// _ = "end of CoverTab[66107]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:429
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:429
			_go_fuzz_dep_.CoverTab[66108]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:429
			// _ = "end of CoverTab[66108]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:430
		// _ = "end of CoverTab[66097]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:430
		_go_fuzz_dep_.CoverTab[66098]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:433
		if len(s) != len(strings.TrimSpace(s)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:433
			_go_fuzz_dep_.CoverTab[66115]++
															return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:434
			// _ = "end of CoverTab[66115]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:435
			_go_fuzz_dep_.CoverTab[66116]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:435
			// _ = "end of CoverTab[66116]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:435
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:435
		// _ = "end of CoverTab[66098]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:435
		_go_fuzz_dep_.CoverTab[66099]++
														dec := json.NewDecoder([]byte(s))
														tok, err := dec.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:438
			_go_fuzz_dep_.CoverTab[66117]++
															return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:439
			// _ = "end of CoverTab[66117]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:440
			_go_fuzz_dep_.CoverTab[66118]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:440
			// _ = "end of CoverTab[66118]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:440
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:440
		// _ = "end of CoverTab[66099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:440
		_go_fuzz_dep_.CoverTab[66100]++
														return getFloat(tok, bitSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:441
		// _ = "end of CoverTab[66100]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:441
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:441
		_go_fuzz_dep_.CoverTab[66101]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:441
		// _ = "end of CoverTab[66101]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:442
	// _ = "end of CoverTab[66094]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:442
	_go_fuzz_dep_.CoverTab[66095]++
													return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:443
	// _ = "end of CoverTab[66095]"
}

func getFloat(tok json.Token, bitSize int) (protoreflect.Value, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:446
	_go_fuzz_dep_.CoverTab[66119]++
													n, ok := tok.Float(bitSize)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:448
		_go_fuzz_dep_.CoverTab[66122]++
														return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:449
		// _ = "end of CoverTab[66122]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:450
		_go_fuzz_dep_.CoverTab[66123]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:450
		// _ = "end of CoverTab[66123]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:450
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:450
	// _ = "end of CoverTab[66119]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:450
	_go_fuzz_dep_.CoverTab[66120]++
													if bitSize == 32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:451
		_go_fuzz_dep_.CoverTab[66124]++
														return protoreflect.ValueOfFloat32(float32(n)), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:452
		// _ = "end of CoverTab[66124]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:453
		_go_fuzz_dep_.CoverTab[66125]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:453
		// _ = "end of CoverTab[66125]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:453
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:453
	// _ = "end of CoverTab[66120]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:453
	_go_fuzz_dep_.CoverTab[66121]++
													return protoreflect.ValueOfFloat64(n), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:454
	// _ = "end of CoverTab[66121]"
}

func unmarshalBytes(tok json.Token) (protoreflect.Value, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:457
	_go_fuzz_dep_.CoverTab[66126]++
													if tok.Kind() != json.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:458
		_go_fuzz_dep_.CoverTab[66131]++
														return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:459
		// _ = "end of CoverTab[66131]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:460
		_go_fuzz_dep_.CoverTab[66132]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:460
		// _ = "end of CoverTab[66132]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:460
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:460
	// _ = "end of CoverTab[66126]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:460
	_go_fuzz_dep_.CoverTab[66127]++

													s := tok.ParsedString()
													enc := base64.StdEncoding
													if strings.ContainsAny(s, "-_") {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:464
		_go_fuzz_dep_.CoverTab[66133]++
														enc = base64.URLEncoding
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:465
		// _ = "end of CoverTab[66133]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:466
		_go_fuzz_dep_.CoverTab[66134]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:466
		// _ = "end of CoverTab[66134]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:466
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:466
	// _ = "end of CoverTab[66127]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:466
	_go_fuzz_dep_.CoverTab[66128]++
													if len(s)%4 != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:467
		_go_fuzz_dep_.CoverTab[66135]++
														enc = enc.WithPadding(base64.NoPadding)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:468
		// _ = "end of CoverTab[66135]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:469
		_go_fuzz_dep_.CoverTab[66136]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:469
		// _ = "end of CoverTab[66136]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:469
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:469
	// _ = "end of CoverTab[66128]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:469
	_go_fuzz_dep_.CoverTab[66129]++
													b, err := enc.DecodeString(s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:471
		_go_fuzz_dep_.CoverTab[66137]++
														return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:472
		// _ = "end of CoverTab[66137]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:473
		_go_fuzz_dep_.CoverTab[66138]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:473
		// _ = "end of CoverTab[66138]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:473
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:473
	// _ = "end of CoverTab[66129]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:473
	_go_fuzz_dep_.CoverTab[66130]++
													return protoreflect.ValueOfBytes(b), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:474
	// _ = "end of CoverTab[66130]"
}

func unmarshalEnum(tok json.Token, fd protoreflect.FieldDescriptor) (protoreflect.Value, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:477
	_go_fuzz_dep_.CoverTab[66139]++
													switch tok.Kind() {
	case json.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:479
		_go_fuzz_dep_.CoverTab[66141]++

														s := tok.ParsedString()
														if enumVal := fd.Enum().Values().ByName(protoreflect.Name(s)); enumVal != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:482
			_go_fuzz_dep_.CoverTab[66145]++
															return protoreflect.ValueOfEnum(enumVal.Number()), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:483
			// _ = "end of CoverTab[66145]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:484
			_go_fuzz_dep_.CoverTab[66146]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:484
			// _ = "end of CoverTab[66146]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:484
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:484
		// _ = "end of CoverTab[66141]"

	case json.Number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:486
		_go_fuzz_dep_.CoverTab[66142]++
														if n, ok := tok.Int(32); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:487
			_go_fuzz_dep_.CoverTab[66147]++
															return protoreflect.ValueOfEnum(protoreflect.EnumNumber(n)), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:488
			// _ = "end of CoverTab[66147]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:489
			_go_fuzz_dep_.CoverTab[66148]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:489
			// _ = "end of CoverTab[66148]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:489
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:489
		// _ = "end of CoverTab[66142]"

	case json.Null:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:491
		_go_fuzz_dep_.CoverTab[66143]++

														if isNullValue(fd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:493
			_go_fuzz_dep_.CoverTab[66149]++
															return protoreflect.ValueOfEnum(0), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:494
			// _ = "end of CoverTab[66149]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:495
			_go_fuzz_dep_.CoverTab[66150]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:495
			// _ = "end of CoverTab[66150]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:495
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:495
		// _ = "end of CoverTab[66143]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:495
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:495
		_go_fuzz_dep_.CoverTab[66144]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:495
		// _ = "end of CoverTab[66144]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:496
	// _ = "end of CoverTab[66139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:496
	_go_fuzz_dep_.CoverTab[66140]++

													return protoreflect.Value{}, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:498
	// _ = "end of CoverTab[66140]"
}

func (d decoder) unmarshalList(list protoreflect.List, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:501
	_go_fuzz_dep_.CoverTab[66151]++
													tok, err := d.Read()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:503
		_go_fuzz_dep_.CoverTab[66155]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:504
		// _ = "end of CoverTab[66155]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:505
		_go_fuzz_dep_.CoverTab[66156]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:505
		// _ = "end of CoverTab[66156]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:505
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:505
	// _ = "end of CoverTab[66151]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:505
	_go_fuzz_dep_.CoverTab[66152]++
													if tok.Kind() != json.ArrayOpen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:506
		_go_fuzz_dep_.CoverTab[66157]++
														return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:507
		// _ = "end of CoverTab[66157]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:508
		_go_fuzz_dep_.CoverTab[66158]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:508
		// _ = "end of CoverTab[66158]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:508
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:508
	// _ = "end of CoverTab[66152]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:508
	_go_fuzz_dep_.CoverTab[66153]++

													switch fd.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:511
		_go_fuzz_dep_.CoverTab[66159]++
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:512
			_go_fuzz_dep_.CoverTab[66161]++
															tok, err := d.Peek()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:514
				_go_fuzz_dep_.CoverTab[66165]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:515
				// _ = "end of CoverTab[66165]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:516
				_go_fuzz_dep_.CoverTab[66166]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:516
				// _ = "end of CoverTab[66166]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:516
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:516
			// _ = "end of CoverTab[66161]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:516
			_go_fuzz_dep_.CoverTab[66162]++

															if tok.Kind() == json.ArrayClose {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:518
				_go_fuzz_dep_.CoverTab[66167]++
																d.Read()
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:520
				// _ = "end of CoverTab[66167]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:521
				_go_fuzz_dep_.CoverTab[66168]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:521
				// _ = "end of CoverTab[66168]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:521
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:521
			// _ = "end of CoverTab[66162]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:521
			_go_fuzz_dep_.CoverTab[66163]++

															val := list.NewElement()
															if err := d.unmarshalMessage(val.Message(), false); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:524
				_go_fuzz_dep_.CoverTab[66169]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:525
				// _ = "end of CoverTab[66169]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:526
				_go_fuzz_dep_.CoverTab[66170]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:526
				// _ = "end of CoverTab[66170]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:526
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:526
			// _ = "end of CoverTab[66163]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:526
			_go_fuzz_dep_.CoverTab[66164]++
															list.Append(val)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:527
			// _ = "end of CoverTab[66164]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:528
		// _ = "end of CoverTab[66159]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:529
		_go_fuzz_dep_.CoverTab[66160]++
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:530
			_go_fuzz_dep_.CoverTab[66171]++
															tok, err := d.Peek()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:532
				_go_fuzz_dep_.CoverTab[66175]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:533
				// _ = "end of CoverTab[66175]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:534
				_go_fuzz_dep_.CoverTab[66176]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:534
				// _ = "end of CoverTab[66176]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:534
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:534
			// _ = "end of CoverTab[66171]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:534
			_go_fuzz_dep_.CoverTab[66172]++

															if tok.Kind() == json.ArrayClose {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:536
				_go_fuzz_dep_.CoverTab[66177]++
																d.Read()
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:538
				// _ = "end of CoverTab[66177]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:539
				_go_fuzz_dep_.CoverTab[66178]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:539
				// _ = "end of CoverTab[66178]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:539
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:539
			// _ = "end of CoverTab[66172]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:539
			_go_fuzz_dep_.CoverTab[66173]++

															val, err := d.unmarshalScalar(fd)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:542
				_go_fuzz_dep_.CoverTab[66179]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:543
				// _ = "end of CoverTab[66179]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:544
				_go_fuzz_dep_.CoverTab[66180]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:544
				// _ = "end of CoverTab[66180]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:544
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:544
			// _ = "end of CoverTab[66173]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:544
			_go_fuzz_dep_.CoverTab[66174]++
															list.Append(val)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:545
			// _ = "end of CoverTab[66174]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:546
		// _ = "end of CoverTab[66160]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:547
	// _ = "end of CoverTab[66153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:547
	_go_fuzz_dep_.CoverTab[66154]++

													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:549
	// _ = "end of CoverTab[66154]"
}

func (d decoder) unmarshalMap(mmap protoreflect.Map, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:552
	_go_fuzz_dep_.CoverTab[66181]++
													tok, err := d.Read()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:554
		_go_fuzz_dep_.CoverTab[66186]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:555
		// _ = "end of CoverTab[66186]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:556
		_go_fuzz_dep_.CoverTab[66187]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:556
		// _ = "end of CoverTab[66187]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:556
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:556
	// _ = "end of CoverTab[66181]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:556
	_go_fuzz_dep_.CoverTab[66182]++
													if tok.Kind() != json.ObjectOpen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:557
		_go_fuzz_dep_.CoverTab[66188]++
														return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:558
		// _ = "end of CoverTab[66188]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:559
		_go_fuzz_dep_.CoverTab[66189]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:559
		// _ = "end of CoverTab[66189]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:559
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:559
	// _ = "end of CoverTab[66182]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:559
	_go_fuzz_dep_.CoverTab[66183]++

	// Determine ahead whether map entry is a scalar type or a message type in
	// order to call the appropriate unmarshalMapValue func inside the for loop
	// below.
	var unmarshalMapValue func() (protoreflect.Value, error)
	switch fd.MapValue().Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:566
		_go_fuzz_dep_.CoverTab[66190]++
														unmarshalMapValue = func() (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:567
			_go_fuzz_dep_.CoverTab[66192]++
															val := mmap.NewValue()
															if err := d.unmarshalMessage(val.Message(), false); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:569
				_go_fuzz_dep_.CoverTab[66194]++
																return protoreflect.Value{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:570
				// _ = "end of CoverTab[66194]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:571
				_go_fuzz_dep_.CoverTab[66195]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:571
				// _ = "end of CoverTab[66195]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:571
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:571
			// _ = "end of CoverTab[66192]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:571
			_go_fuzz_dep_.CoverTab[66193]++
															return val, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:572
			// _ = "end of CoverTab[66193]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:573
		// _ = "end of CoverTab[66190]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:574
		_go_fuzz_dep_.CoverTab[66191]++
														unmarshalMapValue = func() (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:575
			_go_fuzz_dep_.CoverTab[66196]++
															return d.unmarshalScalar(fd.MapValue())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:576
			// _ = "end of CoverTab[66196]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:577
		// _ = "end of CoverTab[66191]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:578
	// _ = "end of CoverTab[66183]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:578
	_go_fuzz_dep_.CoverTab[66184]++

Loop:
	for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:581
		_go_fuzz_dep_.CoverTab[66197]++

														tok, err := d.Read()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:584
			_go_fuzz_dep_.CoverTab[66203]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:585
			// _ = "end of CoverTab[66203]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:586
			_go_fuzz_dep_.CoverTab[66204]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:586
			// _ = "end of CoverTab[66204]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:586
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:586
		// _ = "end of CoverTab[66197]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:586
		_go_fuzz_dep_.CoverTab[66198]++
														switch tok.Kind() {
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:588
			_go_fuzz_dep_.CoverTab[66205]++
															return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:589
			// _ = "end of CoverTab[66205]"
		case json.ObjectClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:590
			_go_fuzz_dep_.CoverTab[66206]++
															break Loop
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:591
			// _ = "end of CoverTab[66206]"
		case json.Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:592
			_go_fuzz_dep_.CoverTab[66207]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:592
			// _ = "end of CoverTab[66207]"

		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:594
		// _ = "end of CoverTab[66198]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:594
		_go_fuzz_dep_.CoverTab[66199]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:597
		pkey, err := d.unmarshalMapKey(tok, fd.MapKey())
		if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:598
			_go_fuzz_dep_.CoverTab[66208]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:599
			// _ = "end of CoverTab[66208]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:600
			_go_fuzz_dep_.CoverTab[66209]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:600
			// _ = "end of CoverTab[66209]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:600
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:600
		// _ = "end of CoverTab[66199]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:600
		_go_fuzz_dep_.CoverTab[66200]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:603
		if mmap.Has(pkey) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:603
			_go_fuzz_dep_.CoverTab[66210]++
															return d.newError(tok.Pos(), "duplicate map key %v", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:604
			// _ = "end of CoverTab[66210]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:605
			_go_fuzz_dep_.CoverTab[66211]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:605
			// _ = "end of CoverTab[66211]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:605
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:605
		// _ = "end of CoverTab[66200]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:605
		_go_fuzz_dep_.CoverTab[66201]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:608
		pval, err := unmarshalMapValue()
		if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:609
			_go_fuzz_dep_.CoverTab[66212]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:610
			// _ = "end of CoverTab[66212]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:611
			_go_fuzz_dep_.CoverTab[66213]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:611
			// _ = "end of CoverTab[66213]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:611
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:611
		// _ = "end of CoverTab[66201]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:611
		_go_fuzz_dep_.CoverTab[66202]++

														mmap.Set(pkey, pval)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:613
		// _ = "end of CoverTab[66202]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:614
	// _ = "end of CoverTab[66184]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:614
	_go_fuzz_dep_.CoverTab[66185]++

													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:616
	// _ = "end of CoverTab[66185]"
}

// unmarshalMapKey converts given token of Name kind into a protoreflect.MapKey.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:619
// A map key type is any integral or string type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:621
func (d decoder) unmarshalMapKey(tok json.Token, fd protoreflect.FieldDescriptor) (protoreflect.MapKey, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:621
	_go_fuzz_dep_.CoverTab[66214]++
													const b32 = 32
													const b64 = 64
													const base10 = 10

													name := tok.Name()
													kind := fd.Kind()
													switch kind {
	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:629
		_go_fuzz_dep_.CoverTab[66216]++
														return protoreflect.ValueOfString(name).MapKey(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:630
		// _ = "end of CoverTab[66216]"

	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:632
		_go_fuzz_dep_.CoverTab[66217]++
														switch name {
		case "true":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:634
			_go_fuzz_dep_.CoverTab[66223]++
															return protoreflect.ValueOfBool(true).MapKey(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:635
			// _ = "end of CoverTab[66223]"
		case "false":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:636
			_go_fuzz_dep_.CoverTab[66224]++
															return protoreflect.ValueOfBool(false).MapKey(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:637
			// _ = "end of CoverTab[66224]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:637
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:637
			_go_fuzz_dep_.CoverTab[66225]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:637
			// _ = "end of CoverTab[66225]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:638
		// _ = "end of CoverTab[66217]"

	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:640
		_go_fuzz_dep_.CoverTab[66218]++
														if n, err := strconv.ParseInt(name, base10, b32); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:641
			_go_fuzz_dep_.CoverTab[66226]++
															return protoreflect.ValueOfInt32(int32(n)).MapKey(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:642
			// _ = "end of CoverTab[66226]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:643
			_go_fuzz_dep_.CoverTab[66227]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:643
			// _ = "end of CoverTab[66227]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:643
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:643
		// _ = "end of CoverTab[66218]"

	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:645
		_go_fuzz_dep_.CoverTab[66219]++
														if n, err := strconv.ParseInt(name, base10, b64); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:646
			_go_fuzz_dep_.CoverTab[66228]++
															return protoreflect.ValueOfInt64(int64(n)).MapKey(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:647
			// _ = "end of CoverTab[66228]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:648
			_go_fuzz_dep_.CoverTab[66229]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:648
			// _ = "end of CoverTab[66229]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:648
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:648
		// _ = "end of CoverTab[66219]"

	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:650
		_go_fuzz_dep_.CoverTab[66220]++
														if n, err := strconv.ParseUint(name, base10, b32); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:651
			_go_fuzz_dep_.CoverTab[66230]++
															return protoreflect.ValueOfUint32(uint32(n)).MapKey(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:652
			// _ = "end of CoverTab[66230]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:653
			_go_fuzz_dep_.CoverTab[66231]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:653
			// _ = "end of CoverTab[66231]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:653
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:653
		// _ = "end of CoverTab[66220]"

	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:655
		_go_fuzz_dep_.CoverTab[66221]++
														if n, err := strconv.ParseUint(name, base10, b64); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:656
			_go_fuzz_dep_.CoverTab[66232]++
															return protoreflect.ValueOfUint64(uint64(n)).MapKey(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:657
			// _ = "end of CoverTab[66232]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:658
			_go_fuzz_dep_.CoverTab[66233]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:658
			// _ = "end of CoverTab[66233]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:658
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:658
		// _ = "end of CoverTab[66221]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:660
		_go_fuzz_dep_.CoverTab[66222]++
														panic(fmt.Sprintf("invalid kind for map key: %v", kind))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:661
		// _ = "end of CoverTab[66222]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:662
	// _ = "end of CoverTab[66214]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:662
	_go_fuzz_dep_.CoverTab[66215]++

													return protoreflect.MapKey{}, d.newError(tok.Pos(), "invalid value for %v key: %s", kind, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:664
	// _ = "end of CoverTab[66215]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:665
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/decode.go:665
var _ = _go_fuzz_dep_.CoverTab
