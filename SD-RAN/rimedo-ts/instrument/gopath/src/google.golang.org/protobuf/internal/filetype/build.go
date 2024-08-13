// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:5
// Package filetype provides functionality for wrapping descriptors
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:5
// with Go type information.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:7
package filetype

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:7
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:7
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:7
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:7
)

import (
	"reflect"

	"google.golang.org/protobuf/internal/descopts"
	"google.golang.org/protobuf/internal/filedesc"
	pimpl "google.golang.org/protobuf/internal/impl"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Builder constructs type descriptors from a raw file descriptor
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// and associated Go types for each enum and message declaration.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// # Flattened Ordering
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// The protobuf type system represents declarations as a tree. Certain nodes in
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// the tree require us to either associate it with a concrete Go type or to
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// resolve a dependency, which is information that must be provided separately
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// since it cannot be derived from the file descriptor alone.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// However, representing a tree as Go literals is difficult to simply do in a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// space and time efficient way. Thus, we store them as a flattened list of
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// objects where the serialization order from the tree-based form is important.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// The "flattened ordering" is defined as a tree traversal of all enum, message,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// extension, and service declarations using the following algorithm:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//	def VisitFileDecls(fd):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//		for e in fd.Enums:      yield e
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//		for m in fd.Messages:   yield m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//		for x in fd.Extensions: yield x
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//		for s in fd.Services:   yield s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//		for m in fd.Messages:   yield from VisitMessageDecls(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//	def VisitMessageDecls(md):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//		for e in md.Enums:      yield e
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//		for m in md.Messages:   yield m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//		for x in md.Extensions: yield x
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//		for m in md.Messages:   yield from VisitMessageDecls(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// The traversal starts at the root file descriptor and yields each direct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// declaration within each node before traversing into sub-declarations
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:19
// that children themselves may have.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:52
type Builder struct {
	// File is the underlying file descriptor builder.
	File	filedesc.Builder

	// GoTypes is a unique set of the Go types for all declarations and
	// dependencies. Each type is represented as a zero value of the Go type.
	//
	// Declarations are Go types generated for enums and messages directly
	// declared (not publicly imported) in the proto source file.
	// Messages for map entries are accounted for, but represented by nil.
	// Enum declarations in "flattened ordering" come first, followed by
	// message declarations in "flattened ordering".
	//
	// Dependencies are Go types for enums or messages referenced by
	// message fields (excluding weak fields), for parent extended messages of
	// extension fields, for enums or messages referenced by extension fields,
	// and for input and output messages referenced by service methods.
	// Dependencies must come after declarations, but the ordering of
	// dependencies themselves is unspecified.
	GoTypes	[]interface{}

	// DependencyIndexes is an ordered list of indexes into GoTypes for the
	// dependencies of messages, extensions, or services.
	//
	// There are 5 sub-lists in "flattened ordering" concatenated back-to-back:
	//	0. Message field dependencies: list of the enum or message type
	//	referred to by every message field.
	//	1. Extension field targets: list of the extended parent message of
	//	every extension.
	//	2. Extension field dependencies: list of the enum or message type
	//	referred to by every extension field.
	//	3. Service method inputs: list of the input message type
	//	referred to by every service method.
	//	4. Service method outputs: list of the output message type
	//	referred to by every service method.
	//
	// The offset into DependencyIndexes for the start of each sub-list
	// is appended to the end in reverse order.
	DependencyIndexes	[]int32

	// EnumInfos is a list of enum infos in "flattened ordering".
	EnumInfos	[]pimpl.EnumInfo

	// MessageInfos is a list of message infos in "flattened ordering".
	// If provided, the GoType and PBType for each element is populated.
	//
	// Requirement: len(MessageInfos) == len(Build.Messages)
	MessageInfos	[]pimpl.MessageInfo

	// ExtensionInfos is a list of extension infos in "flattened ordering".
	// Each element is initialized and registered with the protoregistry package.
	//
	// Requirement: len(LegacyExtensions) == len(Build.Extensions)
	ExtensionInfos	[]pimpl.ExtensionInfo

	// TypeRegistry is the registry to register each type descriptor.
	// If nil, it uses protoregistry.GlobalTypes.
	TypeRegistry	interface {
		RegisterMessage(protoreflect.MessageType) error
		RegisterEnum(protoreflect.EnumType) error
		RegisterExtension(protoreflect.ExtensionType) error
	}
}

// Out is the output of the builder.
type Out struct {
	File protoreflect.FileDescriptor
}

func (tb Builder) Build() (out Out) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:121
	_go_fuzz_dep_.CoverTab[59025]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:124
	if tb.File.FileRegistry == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:124
		_go_fuzz_dep_.CoverTab[59034]++
														tb.File.FileRegistry = protoregistry.GlobalFiles
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:125
		// _ = "end of CoverTab[59034]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:126
		_go_fuzz_dep_.CoverTab[59035]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:126
		// _ = "end of CoverTab[59035]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:126
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:126
	// _ = "end of CoverTab[59025]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:126
	_go_fuzz_dep_.CoverTab[59026]++
													tb.File.FileRegistry = &resolverByIndex{
		goTypes:	tb.GoTypes,
		depIdxs:	tb.DependencyIndexes,
		fileRegistry:	tb.File.FileRegistry,
	}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:134
	if tb.TypeRegistry == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:134
		_go_fuzz_dep_.CoverTab[59036]++
														tb.TypeRegistry = protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:135
		// _ = "end of CoverTab[59036]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:136
		_go_fuzz_dep_.CoverTab[59037]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:136
		// _ = "end of CoverTab[59037]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:136
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:136
	// _ = "end of CoverTab[59026]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:136
	_go_fuzz_dep_.CoverTab[59027]++

													fbOut := tb.File.Build()
													out.File = fbOut.File

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:142
	enumGoTypes := tb.GoTypes[:len(fbOut.Enums)]
	if len(tb.EnumInfos) != len(fbOut.Enums) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:143
		_go_fuzz_dep_.CoverTab[59038]++
														panic("mismatching enum lengths")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:144
		// _ = "end of CoverTab[59038]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:145
		_go_fuzz_dep_.CoverTab[59039]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:145
		// _ = "end of CoverTab[59039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:145
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:145
	// _ = "end of CoverTab[59027]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:145
	_go_fuzz_dep_.CoverTab[59028]++
													if len(fbOut.Enums) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:146
		_go_fuzz_dep_.CoverTab[59040]++
														for i := range fbOut.Enums {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:147
			_go_fuzz_dep_.CoverTab[59041]++
															tb.EnumInfos[i] = pimpl.EnumInfo{
				GoReflectType:	reflect.TypeOf(enumGoTypes[i]),
				Desc:		&fbOut.Enums[i],
			}

			if err := tb.TypeRegistry.RegisterEnum(&tb.EnumInfos[i]); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:153
				_go_fuzz_dep_.CoverTab[59042]++
																panic(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:154
				// _ = "end of CoverTab[59042]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:155
				_go_fuzz_dep_.CoverTab[59043]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:155
				// _ = "end of CoverTab[59043]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:155
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:155
			// _ = "end of CoverTab[59041]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:156
		// _ = "end of CoverTab[59040]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:157
		_go_fuzz_dep_.CoverTab[59044]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:157
		// _ = "end of CoverTab[59044]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:157
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:157
	// _ = "end of CoverTab[59028]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:157
	_go_fuzz_dep_.CoverTab[59029]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:160
	messageGoTypes := tb.GoTypes[len(fbOut.Enums):][:len(fbOut.Messages)]
	if len(tb.MessageInfos) != len(fbOut.Messages) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:161
		_go_fuzz_dep_.CoverTab[59045]++
														panic("mismatching message lengths")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:162
		// _ = "end of CoverTab[59045]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:163
		_go_fuzz_dep_.CoverTab[59046]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:163
		// _ = "end of CoverTab[59046]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:163
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:163
	// _ = "end of CoverTab[59029]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:163
	_go_fuzz_dep_.CoverTab[59030]++
													if len(fbOut.Messages) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:164
		_go_fuzz_dep_.CoverTab[59047]++
														for i := range fbOut.Messages {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:165
			_go_fuzz_dep_.CoverTab[59049]++
															if messageGoTypes[i] == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:166
				_go_fuzz_dep_.CoverTab[59051]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:167
				// _ = "end of CoverTab[59051]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:168
				_go_fuzz_dep_.CoverTab[59052]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:168
				// _ = "end of CoverTab[59052]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:168
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:168
			// _ = "end of CoverTab[59049]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:168
			_go_fuzz_dep_.CoverTab[59050]++

															tb.MessageInfos[i].GoReflectType = reflect.TypeOf(messageGoTypes[i])
															tb.MessageInfos[i].Desc = &fbOut.Messages[i]

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:174
			if err := tb.TypeRegistry.RegisterMessage(&tb.MessageInfos[i]); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:174
				_go_fuzz_dep_.CoverTab[59053]++
																panic(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:175
				// _ = "end of CoverTab[59053]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:176
				_go_fuzz_dep_.CoverTab[59054]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:176
				// _ = "end of CoverTab[59054]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:176
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:176
			// _ = "end of CoverTab[59050]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:177
		// _ = "end of CoverTab[59047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:177
		_go_fuzz_dep_.CoverTab[59048]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:181
		if out.File.Path() == "google/protobuf/descriptor.proto" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:181
			_go_fuzz_dep_.CoverTab[59055]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:181
			return out.File.Package() == "google.protobuf"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:181
			// _ = "end of CoverTab[59055]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:181
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:181
			_go_fuzz_dep_.CoverTab[59056]++
															for i := range fbOut.Messages {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:182
				_go_fuzz_dep_.CoverTab[59057]++
																switch fbOut.Messages[i].Name() {
				case "FileOptions":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:184
					_go_fuzz_dep_.CoverTab[59058]++
																	descopts.File = messageGoTypes[i].(protoreflect.ProtoMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:185
					// _ = "end of CoverTab[59058]"
				case "EnumOptions":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:186
					_go_fuzz_dep_.CoverTab[59059]++
																	descopts.Enum = messageGoTypes[i].(protoreflect.ProtoMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:187
					// _ = "end of CoverTab[59059]"
				case "EnumValueOptions":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:188
					_go_fuzz_dep_.CoverTab[59060]++
																	descopts.EnumValue = messageGoTypes[i].(protoreflect.ProtoMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:189
					// _ = "end of CoverTab[59060]"
				case "MessageOptions":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:190
					_go_fuzz_dep_.CoverTab[59061]++
																	descopts.Message = messageGoTypes[i].(protoreflect.ProtoMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:191
					// _ = "end of CoverTab[59061]"
				case "FieldOptions":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:192
					_go_fuzz_dep_.CoverTab[59062]++
																	descopts.Field = messageGoTypes[i].(protoreflect.ProtoMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:193
					// _ = "end of CoverTab[59062]"
				case "OneofOptions":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:194
					_go_fuzz_dep_.CoverTab[59063]++
																	descopts.Oneof = messageGoTypes[i].(protoreflect.ProtoMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:195
					// _ = "end of CoverTab[59063]"
				case "ExtensionRangeOptions":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:196
					_go_fuzz_dep_.CoverTab[59064]++
																	descopts.ExtensionRange = messageGoTypes[i].(protoreflect.ProtoMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:197
					// _ = "end of CoverTab[59064]"
				case "ServiceOptions":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:198
					_go_fuzz_dep_.CoverTab[59065]++
																	descopts.Service = messageGoTypes[i].(protoreflect.ProtoMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:199
					// _ = "end of CoverTab[59065]"
				case "MethodOptions":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:200
					_go_fuzz_dep_.CoverTab[59066]++
																	descopts.Method = messageGoTypes[i].(protoreflect.ProtoMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:201
					// _ = "end of CoverTab[59066]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:201
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:201
					_go_fuzz_dep_.CoverTab[59067]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:201
					// _ = "end of CoverTab[59067]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:202
				// _ = "end of CoverTab[59057]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:203
			// _ = "end of CoverTab[59056]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:204
			_go_fuzz_dep_.CoverTab[59068]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:204
			// _ = "end of CoverTab[59068]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:204
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:204
		// _ = "end of CoverTab[59048]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:205
		_go_fuzz_dep_.CoverTab[59069]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:205
		// _ = "end of CoverTab[59069]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:205
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:205
	// _ = "end of CoverTab[59030]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:205
	_go_fuzz_dep_.CoverTab[59031]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:208
	if len(tb.ExtensionInfos) != len(fbOut.Extensions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:208
		_go_fuzz_dep_.CoverTab[59070]++
														panic("mismatching extension lengths")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:209
		// _ = "end of CoverTab[59070]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:210
		_go_fuzz_dep_.CoverTab[59071]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:210
		// _ = "end of CoverTab[59071]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:210
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:210
	// _ = "end of CoverTab[59031]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:210
	_go_fuzz_dep_.CoverTab[59032]++
													var depIdx int32
													for i := range fbOut.Extensions {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:212
		_go_fuzz_dep_.CoverTab[59072]++
		// For enum and message kinds, determine the referent Go type so
		// that we can construct their constructors.
		const listExtDeps = 2
		var goType reflect.Type
		switch fbOut.Extensions[i].L1.Kind {
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:218
			_go_fuzz_dep_.CoverTab[59075]++
															j := depIdxs.Get(tb.DependencyIndexes, listExtDeps, depIdx)
															goType = reflect.TypeOf(tb.GoTypes[j])
															depIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:221
			// _ = "end of CoverTab[59075]"
		case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:222
			_go_fuzz_dep_.CoverTab[59076]++
															j := depIdxs.Get(tb.DependencyIndexes, listExtDeps, depIdx)
															goType = reflect.TypeOf(tb.GoTypes[j])
															depIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:225
			// _ = "end of CoverTab[59076]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:226
			_go_fuzz_dep_.CoverTab[59077]++
															goType = goTypeForPBKind[fbOut.Extensions[i].L1.Kind]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:227
			// _ = "end of CoverTab[59077]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:228
		// _ = "end of CoverTab[59072]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:228
		_go_fuzz_dep_.CoverTab[59073]++
														if fbOut.Extensions[i].IsList() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:229
			_go_fuzz_dep_.CoverTab[59078]++
															goType = reflect.SliceOf(goType)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:230
			// _ = "end of CoverTab[59078]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:231
			_go_fuzz_dep_.CoverTab[59079]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:231
			// _ = "end of CoverTab[59079]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:231
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:231
		// _ = "end of CoverTab[59073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:231
		_go_fuzz_dep_.CoverTab[59074]++

														pimpl.InitExtensionInfo(&tb.ExtensionInfos[i], &fbOut.Extensions[i], goType)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:236
		if err := tb.TypeRegistry.RegisterExtension(&tb.ExtensionInfos[i]); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:236
			_go_fuzz_dep_.CoverTab[59080]++
															panic(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:237
			// _ = "end of CoverTab[59080]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:238
			_go_fuzz_dep_.CoverTab[59081]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:238
			// _ = "end of CoverTab[59081]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:238
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:238
		// _ = "end of CoverTab[59074]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:239
	// _ = "end of CoverTab[59032]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:239
	_go_fuzz_dep_.CoverTab[59033]++

													return out
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:241
	// _ = "end of CoverTab[59033]"
}

var goTypeForPBKind = map[protoreflect.Kind]reflect.Type{
	protoreflect.BoolKind:		reflect.TypeOf(bool(false)),
	protoreflect.Int32Kind:		reflect.TypeOf(int32(0)),
	protoreflect.Sint32Kind:	reflect.TypeOf(int32(0)),
	protoreflect.Sfixed32Kind:	reflect.TypeOf(int32(0)),
	protoreflect.Int64Kind:		reflect.TypeOf(int64(0)),
	protoreflect.Sint64Kind:	reflect.TypeOf(int64(0)),
	protoreflect.Sfixed64Kind:	reflect.TypeOf(int64(0)),
	protoreflect.Uint32Kind:	reflect.TypeOf(uint32(0)),
	protoreflect.Fixed32Kind:	reflect.TypeOf(uint32(0)),
	protoreflect.Uint64Kind:	reflect.TypeOf(uint64(0)),
	protoreflect.Fixed64Kind:	reflect.TypeOf(uint64(0)),
	protoreflect.FloatKind:		reflect.TypeOf(float32(0)),
	protoreflect.DoubleKind:	reflect.TypeOf(float64(0)),
	protoreflect.StringKind:	reflect.TypeOf(string("")),
	protoreflect.BytesKind:		reflect.TypeOf([]byte(nil)),
}

type depIdxs []int32

// Get retrieves the jth element of the ith sub-list.
func (x depIdxs) Get(i, j int32) int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:265
	_go_fuzz_dep_.CoverTab[59082]++
													return x[x[int32(len(x))-i-1]+j]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:266
	// _ = "end of CoverTab[59082]"
}

type (
	resolverByIndex	struct {
		goTypes	[]interface{}
		depIdxs	depIdxs
		fileRegistry
	}
	fileRegistry	interface {
		FindFileByPath(string) (protoreflect.FileDescriptor, error)
		FindDescriptorByName(protoreflect.FullName) (protoreflect.Descriptor, error)
		RegisterFile(protoreflect.FileDescriptor) error
	}
)

func (r *resolverByIndex) FindEnumByIndex(i, j int32, es []filedesc.Enum, ms []filedesc.Message) protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:282
	_go_fuzz_dep_.CoverTab[59083]++
													if depIdx := int(r.depIdxs.Get(i, j)); int(depIdx) < len(es)+len(ms) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:283
		_go_fuzz_dep_.CoverTab[59084]++
														return &es[depIdx]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:284
		// _ = "end of CoverTab[59084]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:285
		_go_fuzz_dep_.CoverTab[59085]++
														return pimpl.Export{}.EnumDescriptorOf(r.goTypes[depIdx])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:286
		// _ = "end of CoverTab[59085]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:287
	// _ = "end of CoverTab[59083]"
}

func (r *resolverByIndex) FindMessageByIndex(i, j int32, es []filedesc.Enum, ms []filedesc.Message) protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:290
	_go_fuzz_dep_.CoverTab[59086]++
													if depIdx := int(r.depIdxs.Get(i, j)); depIdx < len(es)+len(ms) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:291
		_go_fuzz_dep_.CoverTab[59087]++
														return &ms[depIdx-len(es)]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:292
		// _ = "end of CoverTab[59087]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:293
		_go_fuzz_dep_.CoverTab[59088]++
														return pimpl.Export{}.MessageDescriptorOf(r.goTypes[depIdx])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:294
		// _ = "end of CoverTab[59088]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:295
	// _ = "end of CoverTab[59086]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:296
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filetype/build.go:296
var _ = _go_fuzz_dep_.CoverTab
