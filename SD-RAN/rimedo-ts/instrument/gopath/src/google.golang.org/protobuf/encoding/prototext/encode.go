// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:5
package prototext

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:5
)

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/encoding/text"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/order"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const defaultIndent = "  "

// Format formats the message as a multiline string.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:28
// This function is only intended for human consumption and ignores errors.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:28
// Do not depend on the output being stable. It may change over time across
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:28
// different versions of the program.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:32
func Format(m proto.Message) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:32
	_go_fuzz_dep_.CoverTab[51994]++
													return MarshalOptions{Multiline: true}.Format(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:33
	// _ = "end of CoverTab[51994]"
}

// Marshal writes the given proto.Message in textproto format using default
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:36
// options. Do not depend on the output being stable. It may change over time
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:36
// across different versions of the program.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:39
func Marshal(m proto.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:39
	_go_fuzz_dep_.CoverTab[51995]++
													return MarshalOptions{}.Marshal(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:40
	// _ = "end of CoverTab[51995]"
}

// MarshalOptions is a configurable text format marshaler.
type MarshalOptions struct {
	pragma.NoUnkeyedLiterals

	// Multiline specifies whether the marshaler should format the output in
	// indented-form with every textual element on a new line.
	// If Indent is an empty string, then an arbitrary indent is chosen.
	Multiline	bool

	// Indent specifies the set of indentation characters to use in a multiline
	// formatted output such that every entry is preceded by Indent and
	// terminated by a newline. If non-empty, then Multiline is treated as true.
	// Indent can only be composed of space or tab characters.
	Indent	string

	// EmitASCII specifies whether to format strings and bytes as ASCII only
	// as opposed to using UTF-8 encoding when possible.
	EmitASCII	bool

	// allowInvalidUTF8 specifies whether to permit the encoding of strings
	// with invalid UTF-8. This is unexported as it is intended to only
	// be specified by the Format method.
	allowInvalidUTF8	bool

	// AllowPartial allows messages that have missing required fields to marshal
	// without returning an error. If AllowPartial is false (the default),
	// Marshal will return error if there are any missing required fields.
	AllowPartial	bool

	// EmitUnknown specifies whether to emit unknown fields in the output.
	// If specified, the unmarshaler may be unable to parse the output.
	// The default is to exclude unknown fields.
	EmitUnknown	bool

	// Resolver is used for looking up types when expanding google.protobuf.Any
	// messages. If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver	interface {
		protoregistry.ExtensionTypeResolver
		protoregistry.MessageTypeResolver
	}
}

// Format formats the message as a string.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:85
// This method is only intended for human consumption and ignores errors.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:85
// Do not depend on the output being stable. It may change over time across
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:85
// different versions of the program.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:89
func (o MarshalOptions) Format(m proto.Message) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:89
	_go_fuzz_dep_.CoverTab[51996]++
													if m == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:90
		_go_fuzz_dep_.CoverTab[51998]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:90
		return !m.ProtoReflect().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:90
		// _ = "end of CoverTab[51998]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:90
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:90
		_go_fuzz_dep_.CoverTab[51999]++
														return "<nil>"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:91
		// _ = "end of CoverTab[51999]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:92
		_go_fuzz_dep_.CoverTab[52000]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:92
		// _ = "end of CoverTab[52000]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:92
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:92
	// _ = "end of CoverTab[51996]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:92
	_go_fuzz_dep_.CoverTab[51997]++
													o.allowInvalidUTF8 = true
													o.AllowPartial = true
													o.EmitUnknown = true
													b, _ := o.Marshal(m)
													return string(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:97
	// _ = "end of CoverTab[51997]"
}

// Marshal writes the given proto.Message in textproto format using options in
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:100
// MarshalOptions object. Do not depend on the output being stable. It may
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:100
// change over time across different versions of the program.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:103
func (o MarshalOptions) Marshal(m proto.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:103
	_go_fuzz_dep_.CoverTab[52001]++
													return o.marshal(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:104
	// _ = "end of CoverTab[52001]"
}

// marshal is a centralized function that all marshal operations go through.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:107
// For profiling purposes, avoid changing the name of this function or
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:107
// introducing other code paths for marshal that do not go through this.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:110
func (o MarshalOptions) marshal(m proto.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:110
	_go_fuzz_dep_.CoverTab[52002]++
													var delims = [2]byte{'{', '}'}

													if o.Multiline && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:113
		_go_fuzz_dep_.CoverTab[52010]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:113
		return o.Indent == ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:113
		// _ = "end of CoverTab[52010]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:113
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:113
		_go_fuzz_dep_.CoverTab[52011]++
														o.Indent = defaultIndent
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:114
		// _ = "end of CoverTab[52011]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:115
		_go_fuzz_dep_.CoverTab[52012]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:115
		// _ = "end of CoverTab[52012]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:115
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:115
	// _ = "end of CoverTab[52002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:115
	_go_fuzz_dep_.CoverTab[52003]++
													if o.Resolver == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:116
		_go_fuzz_dep_.CoverTab[52013]++
														o.Resolver = protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:117
		// _ = "end of CoverTab[52013]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:118
		_go_fuzz_dep_.CoverTab[52014]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:118
		// _ = "end of CoverTab[52014]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:118
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:118
	// _ = "end of CoverTab[52003]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:118
	_go_fuzz_dep_.CoverTab[52004]++

													internalEnc, err := text.NewEncoder(o.Indent, delims, o.EmitASCII)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:121
		_go_fuzz_dep_.CoverTab[52015]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:122
		// _ = "end of CoverTab[52015]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:123
		_go_fuzz_dep_.CoverTab[52016]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:123
		// _ = "end of CoverTab[52016]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:123
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:123
	// _ = "end of CoverTab[52004]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:123
	_go_fuzz_dep_.CoverTab[52005]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:127
	if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:127
		_go_fuzz_dep_.CoverTab[52017]++
														return []byte{}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:128
		// _ = "end of CoverTab[52017]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:129
		_go_fuzz_dep_.CoverTab[52018]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:129
		// _ = "end of CoverTab[52018]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:129
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:129
	// _ = "end of CoverTab[52005]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:129
	_go_fuzz_dep_.CoverTab[52006]++

													enc := encoder{internalEnc, o}
													err = enc.marshalMessage(m.ProtoReflect(), false)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:133
		_go_fuzz_dep_.CoverTab[52019]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:134
		// _ = "end of CoverTab[52019]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:135
		_go_fuzz_dep_.CoverTab[52020]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:135
		// _ = "end of CoverTab[52020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:135
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:135
	// _ = "end of CoverTab[52006]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:135
	_go_fuzz_dep_.CoverTab[52007]++
													out := enc.Bytes()
													if len(o.Indent) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:137
		_go_fuzz_dep_.CoverTab[52021]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:137
		return len(out) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:137
		// _ = "end of CoverTab[52021]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:137
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:137
		_go_fuzz_dep_.CoverTab[52022]++
														out = append(out, '\n')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:138
		// _ = "end of CoverTab[52022]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:139
		_go_fuzz_dep_.CoverTab[52023]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:139
		// _ = "end of CoverTab[52023]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:139
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:139
	// _ = "end of CoverTab[52007]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:139
	_go_fuzz_dep_.CoverTab[52008]++
													if o.AllowPartial {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:140
		_go_fuzz_dep_.CoverTab[52024]++
														return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:141
		// _ = "end of CoverTab[52024]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:142
		_go_fuzz_dep_.CoverTab[52025]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:142
		// _ = "end of CoverTab[52025]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:142
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:142
	// _ = "end of CoverTab[52008]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:142
	_go_fuzz_dep_.CoverTab[52009]++
													return out, proto.CheckInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:143
	// _ = "end of CoverTab[52009]"
}

type encoder struct {
	*text.Encoder
	opts	MarshalOptions
}

// marshalMessage marshals the given protoreflect.Message.
func (e encoder) marshalMessage(m protoreflect.Message, inclDelims bool) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:152
	_go_fuzz_dep_.CoverTab[52026]++
													messageDesc := m.Descriptor()
													if !flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:154
		_go_fuzz_dep_.CoverTab[52033]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:154
		return messageset.IsMessageSet(messageDesc)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:154
		// _ = "end of CoverTab[52033]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:154
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:154
		_go_fuzz_dep_.CoverTab[52034]++
														return errors.New("no support for proto1 MessageSets")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:155
		// _ = "end of CoverTab[52034]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:156
		_go_fuzz_dep_.CoverTab[52035]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:156
		// _ = "end of CoverTab[52035]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:156
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:156
	// _ = "end of CoverTab[52026]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:156
	_go_fuzz_dep_.CoverTab[52027]++

													if inclDelims {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:158
		_go_fuzz_dep_.CoverTab[52036]++
														e.StartMessage()
														defer e.EndMessage()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:160
		// _ = "end of CoverTab[52036]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:161
		_go_fuzz_dep_.CoverTab[52037]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:161
		// _ = "end of CoverTab[52037]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:161
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:161
	// _ = "end of CoverTab[52027]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:161
	_go_fuzz_dep_.CoverTab[52028]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:164
	if messageDesc.FullName() == genid.Any_message_fullname {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:164
		_go_fuzz_dep_.CoverTab[52038]++
														if e.marshalAny(m) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:165
			_go_fuzz_dep_.CoverTab[52039]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:166
			// _ = "end of CoverTab[52039]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:167
			_go_fuzz_dep_.CoverTab[52040]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:167
			// _ = "end of CoverTab[52040]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:167
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:167
		// _ = "end of CoverTab[52038]"

	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:169
		_go_fuzz_dep_.CoverTab[52041]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:169
		// _ = "end of CoverTab[52041]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:169
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:169
	// _ = "end of CoverTab[52028]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:169
	_go_fuzz_dep_.CoverTab[52029]++

	// Marshal fields.
	var err error
	order.RangeFields(m, order.IndexNameFieldOrder, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:173
		_go_fuzz_dep_.CoverTab[52042]++
														if err = e.marshalField(fd.TextName(), v, fd); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:174
			_go_fuzz_dep_.CoverTab[52044]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:175
			// _ = "end of CoverTab[52044]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:176
			_go_fuzz_dep_.CoverTab[52045]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:176
			// _ = "end of CoverTab[52045]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:176
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:176
		// _ = "end of CoverTab[52042]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:176
		_go_fuzz_dep_.CoverTab[52043]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:177
		// _ = "end of CoverTab[52043]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:178
	// _ = "end of CoverTab[52029]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:178
	_go_fuzz_dep_.CoverTab[52030]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:179
		_go_fuzz_dep_.CoverTab[52046]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:180
		// _ = "end of CoverTab[52046]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:181
		_go_fuzz_dep_.CoverTab[52047]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:181
		// _ = "end of CoverTab[52047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:181
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:181
	// _ = "end of CoverTab[52030]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:181
	_go_fuzz_dep_.CoverTab[52031]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:184
	if e.opts.EmitUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:184
		_go_fuzz_dep_.CoverTab[52048]++
														e.marshalUnknown(m.GetUnknown())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:185
		// _ = "end of CoverTab[52048]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:186
		_go_fuzz_dep_.CoverTab[52049]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:186
		// _ = "end of CoverTab[52049]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:186
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:186
	// _ = "end of CoverTab[52031]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:186
	_go_fuzz_dep_.CoverTab[52032]++

													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:188
	// _ = "end of CoverTab[52032]"
}

// marshalField marshals the given field with protoreflect.Value.
func (e encoder) marshalField(name string, val protoreflect.Value, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:192
	_go_fuzz_dep_.CoverTab[52050]++
													switch {
	case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:194
		_go_fuzz_dep_.CoverTab[52051]++
														return e.marshalList(name, val.List(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:195
		// _ = "end of CoverTab[52051]"
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:196
		_go_fuzz_dep_.CoverTab[52052]++
														return e.marshalMap(name, val.Map(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:197
		// _ = "end of CoverTab[52052]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:198
		_go_fuzz_dep_.CoverTab[52053]++
														e.WriteName(name)
														return e.marshalSingular(val, fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:200
		// _ = "end of CoverTab[52053]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:201
	// _ = "end of CoverTab[52050]"
}

// marshalSingular marshals the given non-repeated field value. This includes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:204
// all scalar types, enums, messages, and groups.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:206
func (e encoder) marshalSingular(val protoreflect.Value, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:206
	_go_fuzz_dep_.CoverTab[52054]++
													kind := fd.Kind()
													switch kind {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:209
		_go_fuzz_dep_.CoverTab[52056]++
														e.WriteBool(val.Bool())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:210
		// _ = "end of CoverTab[52056]"

	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:212
		_go_fuzz_dep_.CoverTab[52057]++
														s := val.String()
														if !e.opts.allowInvalidUTF8 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:214
			_go_fuzz_dep_.CoverTab[52067]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:214
			return strs.EnforceUTF8(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:214
			// _ = "end of CoverTab[52067]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:214
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:214
			_go_fuzz_dep_.CoverTab[52068]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:214
			return !utf8.ValidString(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:214
			// _ = "end of CoverTab[52068]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:214
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:214
			_go_fuzz_dep_.CoverTab[52069]++
															return errors.InvalidUTF8(string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:215
			// _ = "end of CoverTab[52069]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:216
			_go_fuzz_dep_.CoverTab[52070]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:216
			// _ = "end of CoverTab[52070]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:216
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:216
		// _ = "end of CoverTab[52057]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:216
		_go_fuzz_dep_.CoverTab[52058]++
														e.WriteString(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:217
		// _ = "end of CoverTab[52058]"

	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:221
		_go_fuzz_dep_.CoverTab[52059]++
														e.WriteInt(val.Int())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:222
		// _ = "end of CoverTab[52059]"

	case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:225
		_go_fuzz_dep_.CoverTab[52060]++
														e.WriteUint(val.Uint())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:226
		// _ = "end of CoverTab[52060]"

	case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:228
		_go_fuzz_dep_.CoverTab[52061]++

														e.WriteFloat(val.Float(), 32)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:230
		// _ = "end of CoverTab[52061]"

	case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:232
		_go_fuzz_dep_.CoverTab[52062]++

														e.WriteFloat(val.Float(), 64)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:234
		// _ = "end of CoverTab[52062]"

	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:236
		_go_fuzz_dep_.CoverTab[52063]++
														e.WriteString(string(val.Bytes()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:237
		// _ = "end of CoverTab[52063]"

	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:239
		_go_fuzz_dep_.CoverTab[52064]++
														num := val.Enum()
														if desc := fd.Enum().Values().ByNumber(num); desc != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:241
			_go_fuzz_dep_.CoverTab[52071]++
															e.WriteLiteral(string(desc.Name()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:242
			// _ = "end of CoverTab[52071]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:243
			_go_fuzz_dep_.CoverTab[52072]++

															e.WriteInt(int64(num))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:245
			// _ = "end of CoverTab[52072]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:246
		// _ = "end of CoverTab[52064]"

	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:248
		_go_fuzz_dep_.CoverTab[52065]++
														return e.marshalMessage(val.Message(), true)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:249
		// _ = "end of CoverTab[52065]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:251
		_go_fuzz_dep_.CoverTab[52066]++
														panic(fmt.Sprintf("%v has unknown kind: %v", fd.FullName(), kind))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:252
		// _ = "end of CoverTab[52066]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:253
	// _ = "end of CoverTab[52054]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:253
	_go_fuzz_dep_.CoverTab[52055]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:254
	// _ = "end of CoverTab[52055]"
}

// marshalList marshals the given protoreflect.List as multiple name-value fields.
func (e encoder) marshalList(name string, list protoreflect.List, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:258
	_go_fuzz_dep_.CoverTab[52073]++
													size := list.Len()
													for i := 0; i < size; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:260
		_go_fuzz_dep_.CoverTab[52075]++
														e.WriteName(name)
														if err := e.marshalSingular(list.Get(i), fd); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:262
			_go_fuzz_dep_.CoverTab[52076]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:263
			// _ = "end of CoverTab[52076]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:264
			_go_fuzz_dep_.CoverTab[52077]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:264
			// _ = "end of CoverTab[52077]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:264
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:264
		// _ = "end of CoverTab[52075]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:265
	// _ = "end of CoverTab[52073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:265
	_go_fuzz_dep_.CoverTab[52074]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:266
	// _ = "end of CoverTab[52074]"
}

// marshalMap marshals the given protoreflect.Map as multiple name-value fields.
func (e encoder) marshalMap(name string, mmap protoreflect.Map, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:270
	_go_fuzz_dep_.CoverTab[52078]++
													var err error
													order.RangeEntries(mmap, order.GenericKeyOrder, func(key protoreflect.MapKey, val protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:272
		_go_fuzz_dep_.CoverTab[52080]++
														e.WriteName(name)
														e.StartMessage()
														defer e.EndMessage()

														e.WriteName(string(genid.MapEntry_Key_field_name))
														err = e.marshalSingular(key.Value(), fd.MapKey())
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:279
			_go_fuzz_dep_.CoverTab[52083]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:280
			// _ = "end of CoverTab[52083]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:281
			_go_fuzz_dep_.CoverTab[52084]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:281
			// _ = "end of CoverTab[52084]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:281
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:281
		// _ = "end of CoverTab[52080]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:281
		_go_fuzz_dep_.CoverTab[52081]++

														e.WriteName(string(genid.MapEntry_Value_field_name))
														err = e.marshalSingular(val, fd.MapValue())
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:285
			_go_fuzz_dep_.CoverTab[52085]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:286
			// _ = "end of CoverTab[52085]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:287
			_go_fuzz_dep_.CoverTab[52086]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:287
			// _ = "end of CoverTab[52086]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:287
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:287
		// _ = "end of CoverTab[52081]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:287
		_go_fuzz_dep_.CoverTab[52082]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:288
		// _ = "end of CoverTab[52082]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:289
	// _ = "end of CoverTab[52078]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:289
	_go_fuzz_dep_.CoverTab[52079]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:290
	// _ = "end of CoverTab[52079]"
}

// marshalUnknown parses the given []byte and marshals fields out.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:293
// This function assumes proper encoding in the given []byte.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:295
func (e encoder) marshalUnknown(b []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:295
	_go_fuzz_dep_.CoverTab[52087]++
													const dec = 10
													const hex = 16
													for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:298
		_go_fuzz_dep_.CoverTab[52088]++
														num, wtype, n := protowire.ConsumeTag(b)
														b = b[n:]
														e.WriteName(strconv.FormatInt(int64(num), dec))

														switch wtype {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:304
			_go_fuzz_dep_.CoverTab[52090]++
															var v uint64
															v, n = protowire.ConsumeVarint(b)
															e.WriteUint(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:307
			// _ = "end of CoverTab[52090]"
		case protowire.Fixed32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:308
			_go_fuzz_dep_.CoverTab[52091]++
															var v uint32
															v, n = protowire.ConsumeFixed32(b)
															e.WriteLiteral("0x" + strconv.FormatUint(uint64(v), hex))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:311
			// _ = "end of CoverTab[52091]"
		case protowire.Fixed64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:312
			_go_fuzz_dep_.CoverTab[52092]++
															var v uint64
															v, n = protowire.ConsumeFixed64(b)
															e.WriteLiteral("0x" + strconv.FormatUint(v, hex))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:315
			// _ = "end of CoverTab[52092]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:316
			_go_fuzz_dep_.CoverTab[52093]++
															var v []byte
															v, n = protowire.ConsumeBytes(b)
															e.WriteString(string(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:319
			// _ = "end of CoverTab[52093]"
		case protowire.StartGroupType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:320
			_go_fuzz_dep_.CoverTab[52094]++
															e.StartMessage()
															var v []byte
															v, n = protowire.ConsumeGroup(num, b)
															e.marshalUnknown(v)
															e.EndMessage()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:325
			// _ = "end of CoverTab[52094]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:326
			_go_fuzz_dep_.CoverTab[52095]++
															panic(fmt.Sprintf("prototext: error parsing unknown field wire type: %v", wtype))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:327
			// _ = "end of CoverTab[52095]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:328
		// _ = "end of CoverTab[52088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:328
		_go_fuzz_dep_.CoverTab[52089]++

														b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:330
		// _ = "end of CoverTab[52089]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:331
	// _ = "end of CoverTab[52087]"
}

// marshalAny marshals the given google.protobuf.Any message in expanded form.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:334
// It returns true if it was able to marshal, else false.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:336
func (e encoder) marshalAny(any protoreflect.Message) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:336
	_go_fuzz_dep_.CoverTab[52096]++

													fds := any.Descriptor().Fields()
													fdType := fds.ByNumber(genid.Any_TypeUrl_field_number)
													typeURL := any.Get(fdType).String()
													mt, err := e.opts.Resolver.FindMessageByURL(typeURL)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:342
		_go_fuzz_dep_.CoverTab[52100]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:343
		// _ = "end of CoverTab[52100]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:344
		_go_fuzz_dep_.CoverTab[52101]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:344
		// _ = "end of CoverTab[52101]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:344
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:344
	// _ = "end of CoverTab[52096]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:344
	_go_fuzz_dep_.CoverTab[52097]++
													m := mt.New().Interface()

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:348
	fdValue := fds.ByNumber(genid.Any_Value_field_number)
	value := any.Get(fdValue)
	err = proto.UnmarshalOptions{
		AllowPartial:	true,
		Resolver:	e.opts.Resolver,
	}.Unmarshal(value.Bytes(), m)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:354
		_go_fuzz_dep_.CoverTab[52102]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:355
		// _ = "end of CoverTab[52102]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:356
		_go_fuzz_dep_.CoverTab[52103]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:356
		// _ = "end of CoverTab[52103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:356
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:356
	// _ = "end of CoverTab[52097]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:356
	_go_fuzz_dep_.CoverTab[52098]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:360
	pos := e.Snapshot()

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:363
	e.WriteName("[" + typeURL + "]")
	err = e.marshalMessage(m.ProtoReflect(), true)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:365
		_go_fuzz_dep_.CoverTab[52104]++
														e.Reset(pos)
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:367
		// _ = "end of CoverTab[52104]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:368
		_go_fuzz_dep_.CoverTab[52105]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:368
		// _ = "end of CoverTab[52105]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:368
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:368
	// _ = "end of CoverTab[52098]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:368
	_go_fuzz_dep_.CoverTab[52099]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:369
	// _ = "end of CoverTab[52099]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:370
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/prototext/encode.go:370
var _ = _go_fuzz_dep_.CoverTab
