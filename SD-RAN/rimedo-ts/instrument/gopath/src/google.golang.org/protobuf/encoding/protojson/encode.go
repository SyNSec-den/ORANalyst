// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:5
package protojson

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:5
)

import (
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/internal/encoding/json"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/order"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const defaultIndent = "  "

// Format formats the message as a multiline string.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:26
// This function is only intended for human consumption and ignores errors.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:26
// Do not depend on the output being stable. It may change over time across
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:26
// different versions of the program.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:30
func Format(m proto.Message) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:30
	_go_fuzz_dep_.CoverTab[66234]++
													return MarshalOptions{Multiline: true}.Format(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:31
	// _ = "end of CoverTab[66234]"
}

// Marshal writes the given proto.Message in JSON format using default options.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:34
// Do not depend on the output being stable. It may change over time across
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:34
// different versions of the program.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:37
func Marshal(m proto.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:37
	_go_fuzz_dep_.CoverTab[66235]++
													return MarshalOptions{}.Marshal(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:38
	// _ = "end of CoverTab[66235]"
}

// MarshalOptions is a configurable JSON format marshaler.
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

	// AllowPartial allows messages that have missing required fields to marshal
	// without returning an error. If AllowPartial is false (the default),
	// Marshal will return error if there are any missing required fields.
	AllowPartial	bool

	// UseProtoNames uses proto field name instead of lowerCamelCase name in JSON
	// field names.
	UseProtoNames	bool

	// UseEnumNumbers emits enum values as numbers.
	UseEnumNumbers	bool

	// EmitUnpopulated specifies whether to emit unpopulated fields. It does not
	// emit unpopulated oneof fields or unpopulated extension fields.
	// The JSON value emitted for unpopulated fields are as follows:
	//  ╔═══════╤════════════════════════════╗
	//  ║ JSON  │ Protobuf field             ║
	//  ╠═══════╪════════════════════════════╣
	//  ║ false │ proto3 boolean fields      ║
	//  ║ 0     │ proto3 numeric fields      ║
	//  ║ ""    │ proto3 string/bytes fields ║
	//  ║ null  │ proto2 scalar fields       ║
	//  ║ null  │ message fields             ║
	//  ║ []    │ list fields                ║
	//  ║ {}    │ map fields                 ║
	//  ╚═══════╧════════════════════════════╝
	EmitUnpopulated	bool

	// Resolver is used for looking up types when expanding google.protobuf.Any
	// messages. If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver	interface {
		protoregistry.ExtensionTypeResolver
		protoregistry.MessageTypeResolver
	}
}

// Format formats the message as a string.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:92
// This method is only intended for human consumption and ignores errors.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:92
// Do not depend on the output being stable. It may change over time across
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:92
// different versions of the program.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:96
func (o MarshalOptions) Format(m proto.Message) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:96
	_go_fuzz_dep_.CoverTab[66236]++
													if m == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:97
		_go_fuzz_dep_.CoverTab[66238]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:97
		return !m.ProtoReflect().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:97
		// _ = "end of CoverTab[66238]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:97
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:97
		_go_fuzz_dep_.CoverTab[66239]++
														return "<nil>"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:98
		// _ = "end of CoverTab[66239]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:99
		_go_fuzz_dep_.CoverTab[66240]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:99
		// _ = "end of CoverTab[66240]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:99
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:99
	// _ = "end of CoverTab[66236]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:99
	_go_fuzz_dep_.CoverTab[66237]++
													o.AllowPartial = true
													b, _ := o.Marshal(m)
													return string(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:102
	// _ = "end of CoverTab[66237]"
}

// Marshal marshals the given proto.Message in the JSON format using options in
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:105
// MarshalOptions. Do not depend on the output being stable. It may change over
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:105
// time across different versions of the program.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:108
func (o MarshalOptions) Marshal(m proto.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:108
	_go_fuzz_dep_.CoverTab[66241]++
													return o.marshal(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:109
	// _ = "end of CoverTab[66241]"
}

// marshal is a centralized function that all marshal operations go through.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:112
// For profiling purposes, avoid changing the name of this function or
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:112
// introducing other code paths for marshal that do not go through this.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:115
func (o MarshalOptions) marshal(m proto.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:115
	_go_fuzz_dep_.CoverTab[66242]++
													if o.Multiline && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:116
		_go_fuzz_dep_.CoverTab[66249]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:116
		return o.Indent == ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:116
		// _ = "end of CoverTab[66249]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:116
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:116
		_go_fuzz_dep_.CoverTab[66250]++
														o.Indent = defaultIndent
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:117
		// _ = "end of CoverTab[66250]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:118
		_go_fuzz_dep_.CoverTab[66251]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:118
		// _ = "end of CoverTab[66251]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:118
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:118
	// _ = "end of CoverTab[66242]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:118
	_go_fuzz_dep_.CoverTab[66243]++
													if o.Resolver == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:119
		_go_fuzz_dep_.CoverTab[66252]++
														o.Resolver = protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:120
		// _ = "end of CoverTab[66252]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:121
		_go_fuzz_dep_.CoverTab[66253]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:121
		// _ = "end of CoverTab[66253]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:121
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:121
	// _ = "end of CoverTab[66243]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:121
	_go_fuzz_dep_.CoverTab[66244]++

													internalEnc, err := json.NewEncoder(o.Indent)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:124
		_go_fuzz_dep_.CoverTab[66254]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:125
		// _ = "end of CoverTab[66254]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:126
		_go_fuzz_dep_.CoverTab[66255]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:126
		// _ = "end of CoverTab[66255]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:126
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:126
	// _ = "end of CoverTab[66244]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:126
	_go_fuzz_dep_.CoverTab[66245]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:130
	if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:130
		_go_fuzz_dep_.CoverTab[66256]++
														return []byte("{}"), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:131
		// _ = "end of CoverTab[66256]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:132
		_go_fuzz_dep_.CoverTab[66257]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:132
		// _ = "end of CoverTab[66257]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:132
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:132
	// _ = "end of CoverTab[66245]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:132
	_go_fuzz_dep_.CoverTab[66246]++

													enc := encoder{internalEnc, o}
													if err := enc.marshalMessage(m.ProtoReflect(), ""); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:135
		_go_fuzz_dep_.CoverTab[66258]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:136
		// _ = "end of CoverTab[66258]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:137
		_go_fuzz_dep_.CoverTab[66259]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:137
		// _ = "end of CoverTab[66259]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:137
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:137
	// _ = "end of CoverTab[66246]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:137
	_go_fuzz_dep_.CoverTab[66247]++
													if o.AllowPartial {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:138
		_go_fuzz_dep_.CoverTab[66260]++
														return enc.Bytes(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:139
		// _ = "end of CoverTab[66260]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:140
		_go_fuzz_dep_.CoverTab[66261]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:140
		// _ = "end of CoverTab[66261]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:140
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:140
	// _ = "end of CoverTab[66247]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:140
	_go_fuzz_dep_.CoverTab[66248]++
													return enc.Bytes(), proto.CheckInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:141
	// _ = "end of CoverTab[66248]"
}

type encoder struct {
	*json.Encoder
	opts	MarshalOptions
}

// typeFieldDesc is a synthetic field descriptor used for the "@type" field.
var typeFieldDesc = func() protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:150
	_go_fuzz_dep_.CoverTab[66262]++
													var fd filedesc.Field
													fd.L0.FullName = "@type"
													fd.L0.Index = -1
													fd.L1.Cardinality = protoreflect.Optional
													fd.L1.Kind = protoreflect.StringKind
													return &fd
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:156
	// _ = "end of CoverTab[66262]"
}()

// typeURLFieldRanger wraps a protoreflect.Message and modifies its Range method
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:159
// to additionally iterate over a synthetic field for the type URL.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:161
type typeURLFieldRanger struct {
	order.FieldRanger
	typeURL	string
}

func (m typeURLFieldRanger) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:166
	_go_fuzz_dep_.CoverTab[66263]++
													if !f(typeFieldDesc, protoreflect.ValueOfString(m.typeURL)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:167
		_go_fuzz_dep_.CoverTab[66265]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:168
		// _ = "end of CoverTab[66265]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:169
		_go_fuzz_dep_.CoverTab[66266]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:169
		// _ = "end of CoverTab[66266]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:169
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:169
	// _ = "end of CoverTab[66263]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:169
	_go_fuzz_dep_.CoverTab[66264]++
													m.FieldRanger.Range(f)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:170
	// _ = "end of CoverTab[66264]"
}

// unpopulatedFieldRanger wraps a protoreflect.Message and modifies its Range
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:173
// method to additionally iterate over unpopulated fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:175
type unpopulatedFieldRanger struct{ protoreflect.Message }

func (m unpopulatedFieldRanger) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:177
	_go_fuzz_dep_.CoverTab[66267]++
													fds := m.Descriptor().Fields()
													for i := 0; i < fds.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:179
		_go_fuzz_dep_.CoverTab[66269]++
														fd := fds.Get(i)
														if m.Has(fd) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:181
			_go_fuzz_dep_.CoverTab[66272]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:181
			return fd.ContainingOneof() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:181
			// _ = "end of CoverTab[66272]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:181
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:181
			_go_fuzz_dep_.CoverTab[66273]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:182
			// _ = "end of CoverTab[66273]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:183
			_go_fuzz_dep_.CoverTab[66274]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:183
			// _ = "end of CoverTab[66274]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:183
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:183
		// _ = "end of CoverTab[66269]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:183
		_go_fuzz_dep_.CoverTab[66270]++

														v := m.Get(fd)
														isProto2Scalar := fd.Syntax() == protoreflect.Proto2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:186
			_go_fuzz_dep_.CoverTab[66275]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:186
			return fd.Default().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:186
			// _ = "end of CoverTab[66275]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:186
		}()
														isSingularMessage := fd.Cardinality() != protoreflect.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:187
			_go_fuzz_dep_.CoverTab[66276]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:187
			return fd.Message() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:187
			// _ = "end of CoverTab[66276]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:187
		}()
														if isProto2Scalar || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:188
			_go_fuzz_dep_.CoverTab[66277]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:188
			return isSingularMessage
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:188
			// _ = "end of CoverTab[66277]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:188
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:188
			_go_fuzz_dep_.CoverTab[66278]++
															v = protoreflect.Value{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:189
			// _ = "end of CoverTab[66278]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:190
			_go_fuzz_dep_.CoverTab[66279]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:190
			// _ = "end of CoverTab[66279]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:190
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:190
		// _ = "end of CoverTab[66270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:190
		_go_fuzz_dep_.CoverTab[66271]++
														if !f(fd, v) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:191
			_go_fuzz_dep_.CoverTab[66280]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:192
			// _ = "end of CoverTab[66280]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:193
			_go_fuzz_dep_.CoverTab[66281]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:193
			// _ = "end of CoverTab[66281]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:193
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:193
		// _ = "end of CoverTab[66271]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:194
	// _ = "end of CoverTab[66267]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:194
	_go_fuzz_dep_.CoverTab[66268]++
													m.Message.Range(f)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:195
	// _ = "end of CoverTab[66268]"
}

// marshalMessage marshals the fields in the given protoreflect.Message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:198
// If the typeURL is non-empty, then a synthetic "@type" field is injected
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:198
// containing the URL as the value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:201
func (e encoder) marshalMessage(m protoreflect.Message, typeURL string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:201
	_go_fuzz_dep_.CoverTab[66282]++
													if !flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:202
		_go_fuzz_dep_.CoverTab[66288]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:202
		return messageset.IsMessageSet(m.Descriptor())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:202
		// _ = "end of CoverTab[66288]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:202
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:202
		_go_fuzz_dep_.CoverTab[66289]++
														return errors.New("no support for proto1 MessageSets")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:203
		// _ = "end of CoverTab[66289]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:204
		_go_fuzz_dep_.CoverTab[66290]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:204
		// _ = "end of CoverTab[66290]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:204
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:204
	// _ = "end of CoverTab[66282]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:204
	_go_fuzz_dep_.CoverTab[66283]++

													if marshal := wellKnownTypeMarshaler(m.Descriptor().FullName()); marshal != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:206
		_go_fuzz_dep_.CoverTab[66291]++
														return marshal(e, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:207
		// _ = "end of CoverTab[66291]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:208
		_go_fuzz_dep_.CoverTab[66292]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:208
		// _ = "end of CoverTab[66292]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:208
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:208
	// _ = "end of CoverTab[66283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:208
	_go_fuzz_dep_.CoverTab[66284]++

													e.StartObject()
													defer e.EndObject()

													var fields order.FieldRanger = m
													if e.opts.EmitUnpopulated {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:214
		_go_fuzz_dep_.CoverTab[66293]++
														fields = unpopulatedFieldRanger{m}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:215
		// _ = "end of CoverTab[66293]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:216
		_go_fuzz_dep_.CoverTab[66294]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:216
		// _ = "end of CoverTab[66294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:216
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:216
	// _ = "end of CoverTab[66284]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:216
	_go_fuzz_dep_.CoverTab[66285]++
													if typeURL != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:217
		_go_fuzz_dep_.CoverTab[66295]++
														fields = typeURLFieldRanger{fields, typeURL}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:218
		// _ = "end of CoverTab[66295]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:219
		_go_fuzz_dep_.CoverTab[66296]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:219
		// _ = "end of CoverTab[66296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:219
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:219
	// _ = "end of CoverTab[66285]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:219
	_go_fuzz_dep_.CoverTab[66286]++

													var err error
													order.RangeFields(fields, order.IndexNameFieldOrder, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:222
		_go_fuzz_dep_.CoverTab[66297]++
														name := fd.JSONName()
														if e.opts.UseProtoNames {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:224
			_go_fuzz_dep_.CoverTab[66301]++
															name = fd.TextName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:225
			// _ = "end of CoverTab[66301]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:226
			_go_fuzz_dep_.CoverTab[66302]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:226
			// _ = "end of CoverTab[66302]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:226
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:226
		// _ = "end of CoverTab[66297]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:226
		_go_fuzz_dep_.CoverTab[66298]++

														if err = e.WriteName(name); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:228
			_go_fuzz_dep_.CoverTab[66303]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:229
			// _ = "end of CoverTab[66303]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:230
			_go_fuzz_dep_.CoverTab[66304]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:230
			// _ = "end of CoverTab[66304]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:230
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:230
		// _ = "end of CoverTab[66298]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:230
		_go_fuzz_dep_.CoverTab[66299]++
														if err = e.marshalValue(v, fd); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:231
			_go_fuzz_dep_.CoverTab[66305]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:232
			// _ = "end of CoverTab[66305]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:233
			_go_fuzz_dep_.CoverTab[66306]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:233
			// _ = "end of CoverTab[66306]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:233
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:233
		// _ = "end of CoverTab[66299]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:233
		_go_fuzz_dep_.CoverTab[66300]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:234
		// _ = "end of CoverTab[66300]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:235
	// _ = "end of CoverTab[66286]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:235
	_go_fuzz_dep_.CoverTab[66287]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:236
	// _ = "end of CoverTab[66287]"
}

// marshalValue marshals the given protoreflect.Value.
func (e encoder) marshalValue(val protoreflect.Value, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:240
	_go_fuzz_dep_.CoverTab[66307]++
													switch {
	case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:242
		_go_fuzz_dep_.CoverTab[66308]++
														return e.marshalList(val.List(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:243
		// _ = "end of CoverTab[66308]"
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:244
		_go_fuzz_dep_.CoverTab[66309]++
														return e.marshalMap(val.Map(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:245
		// _ = "end of CoverTab[66309]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:246
		_go_fuzz_dep_.CoverTab[66310]++
														return e.marshalSingular(val, fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:247
		// _ = "end of CoverTab[66310]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:248
	// _ = "end of CoverTab[66307]"
}

// marshalSingular marshals the given non-repeated field value. This includes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:251
// all scalar types, enums, messages, and groups.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:253
func (e encoder) marshalSingular(val protoreflect.Value, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:253
	_go_fuzz_dep_.CoverTab[66311]++
													if !val.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:254
		_go_fuzz_dep_.CoverTab[66314]++
														e.WriteNull()
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:256
		// _ = "end of CoverTab[66314]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:257
		_go_fuzz_dep_.CoverTab[66315]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:257
		// _ = "end of CoverTab[66315]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:257
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:257
	// _ = "end of CoverTab[66311]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:257
	_go_fuzz_dep_.CoverTab[66312]++

													switch kind := fd.Kind(); kind {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:260
		_go_fuzz_dep_.CoverTab[66316]++
														e.WriteBool(val.Bool())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:261
		// _ = "end of CoverTab[66316]"

	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:263
		_go_fuzz_dep_.CoverTab[66317]++
														if e.WriteString(val.String()) != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:264
			_go_fuzz_dep_.CoverTab[66327]++
															return errors.InvalidUTF8(string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:265
			// _ = "end of CoverTab[66327]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:266
			_go_fuzz_dep_.CoverTab[66328]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:266
			// _ = "end of CoverTab[66328]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:266
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:266
		// _ = "end of CoverTab[66317]"

	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:268
		_go_fuzz_dep_.CoverTab[66318]++
														e.WriteInt(val.Int())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:269
		// _ = "end of CoverTab[66318]"

	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:271
		_go_fuzz_dep_.CoverTab[66319]++
														e.WriteUint(val.Uint())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:272
		// _ = "end of CoverTab[66319]"

	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind,
		protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:275
		_go_fuzz_dep_.CoverTab[66320]++

														e.WriteString(val.String())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:277
		// _ = "end of CoverTab[66320]"

	case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:279
		_go_fuzz_dep_.CoverTab[66321]++

														e.WriteFloat(val.Float(), 32)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:281
		// _ = "end of CoverTab[66321]"

	case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:283
		_go_fuzz_dep_.CoverTab[66322]++

														e.WriteFloat(val.Float(), 64)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:285
		// _ = "end of CoverTab[66322]"

	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:287
		_go_fuzz_dep_.CoverTab[66323]++
														e.WriteString(base64.StdEncoding.EncodeToString(val.Bytes()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:288
		// _ = "end of CoverTab[66323]"

	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:290
		_go_fuzz_dep_.CoverTab[66324]++
														if fd.Enum().FullName() == genid.NullValue_enum_fullname {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:291
			_go_fuzz_dep_.CoverTab[66329]++
															e.WriteNull()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:292
			// _ = "end of CoverTab[66329]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:293
			_go_fuzz_dep_.CoverTab[66330]++
															desc := fd.Enum().Values().ByNumber(val.Enum())
															if e.opts.UseEnumNumbers || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:295
				_go_fuzz_dep_.CoverTab[66331]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:295
				return desc == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:295
				// _ = "end of CoverTab[66331]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:295
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:295
				_go_fuzz_dep_.CoverTab[66332]++
																e.WriteInt(int64(val.Enum()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:296
				// _ = "end of CoverTab[66332]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:297
				_go_fuzz_dep_.CoverTab[66333]++
																e.WriteString(string(desc.Name()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:298
				// _ = "end of CoverTab[66333]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:299
			// _ = "end of CoverTab[66330]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:300
		// _ = "end of CoverTab[66324]"

	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:302
		_go_fuzz_dep_.CoverTab[66325]++
														if err := e.marshalMessage(val.Message(), ""); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:303
			_go_fuzz_dep_.CoverTab[66334]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:304
			// _ = "end of CoverTab[66334]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:305
			_go_fuzz_dep_.CoverTab[66335]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:305
			// _ = "end of CoverTab[66335]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:305
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:305
		// _ = "end of CoverTab[66325]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:307
		_go_fuzz_dep_.CoverTab[66326]++
														panic(fmt.Sprintf("%v has unknown kind: %v", fd.FullName(), kind))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:308
		// _ = "end of CoverTab[66326]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:309
	// _ = "end of CoverTab[66312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:309
	_go_fuzz_dep_.CoverTab[66313]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:310
	// _ = "end of CoverTab[66313]"
}

// marshalList marshals the given protoreflect.List.
func (e encoder) marshalList(list protoreflect.List, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:314
	_go_fuzz_dep_.CoverTab[66336]++
													e.StartArray()
													defer e.EndArray()

													for i := 0; i < list.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:318
		_go_fuzz_dep_.CoverTab[66338]++
														item := list.Get(i)
														if err := e.marshalSingular(item, fd); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:320
			_go_fuzz_dep_.CoverTab[66339]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:321
			// _ = "end of CoverTab[66339]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:322
			_go_fuzz_dep_.CoverTab[66340]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:322
			// _ = "end of CoverTab[66340]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:322
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:322
		// _ = "end of CoverTab[66338]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:323
	// _ = "end of CoverTab[66336]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:323
	_go_fuzz_dep_.CoverTab[66337]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:324
	// _ = "end of CoverTab[66337]"
}

// marshalMap marshals given protoreflect.Map.
func (e encoder) marshalMap(mmap protoreflect.Map, fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:328
	_go_fuzz_dep_.CoverTab[66341]++
													e.StartObject()
													defer e.EndObject()

													var err error
													order.RangeEntries(mmap, order.GenericKeyOrder, func(k protoreflect.MapKey, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:333
		_go_fuzz_dep_.CoverTab[66343]++
														if err = e.WriteName(k.String()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:334
			_go_fuzz_dep_.CoverTab[66346]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:335
			// _ = "end of CoverTab[66346]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:336
			_go_fuzz_dep_.CoverTab[66347]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:336
			// _ = "end of CoverTab[66347]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:336
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:336
		// _ = "end of CoverTab[66343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:336
		_go_fuzz_dep_.CoverTab[66344]++
														if err = e.marshalSingular(v, fd.MapValue()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:337
			_go_fuzz_dep_.CoverTab[66348]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:338
			// _ = "end of CoverTab[66348]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:339
			_go_fuzz_dep_.CoverTab[66349]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:339
			// _ = "end of CoverTab[66349]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:339
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:339
		// _ = "end of CoverTab[66344]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:339
		_go_fuzz_dep_.CoverTab[66345]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:340
		// _ = "end of CoverTab[66345]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:341
	// _ = "end of CoverTab[66341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:341
	_go_fuzz_dep_.CoverTab[66342]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:342
	// _ = "end of CoverTab[66342]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:343
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/encode.go:343
var _ = _go_fuzz_dep_.CoverTab
