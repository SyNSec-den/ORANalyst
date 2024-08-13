// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:5
// Package filedesc provides functionality for constructing descriptors.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:5
// The types in this package implement interfaces in the protoreflect package
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:5
// related to protobuf descripriptors.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:9
package filedesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:9
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:9
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:9
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:9
)

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Builder construct a protoreflect.FileDescriptor from the raw descriptor.
type Builder struct {
	// GoPackagePath is the Go package path that is invoking this builder.
	GoPackagePath	string

	// RawDescriptor is the wire-encoded bytes of FileDescriptorProto
	// and must be populated.
	RawDescriptor	[]byte

	// NumEnums is the total number of enums declared in the file.
	NumEnums	int32
	// NumMessages is the total number of messages declared in the file.
	// It includes the implicit message declarations for map entries.
	NumMessages	int32
	// NumExtensions is the total number of extensions declared in the file.
	NumExtensions	int32
	// NumServices is the total number of services declared in the file.
	NumServices	int32

	// TypeResolver resolves extension field types for descriptor options.
	// If nil, it uses protoregistry.GlobalTypes.
	TypeResolver	interface {
		protoregistry.ExtensionTypeResolver
	}

	// FileRegistry is use to lookup file, enum, and message dependencies.
	// Once constructed, the file descriptor is registered here.
	// If nil, it uses protoregistry.GlobalFiles.
	FileRegistry	interface {
		FindFileByPath(string) (protoreflect.FileDescriptor, error)
		FindDescriptorByName(protoreflect.FullName) (protoreflect.Descriptor, error)
		RegisterFile(protoreflect.FileDescriptor) error
	}
}

// resolverByIndex is an interface Builder.FileRegistry may implement.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:53
// If so, it permits looking up an enum or message dependency based on the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:53
// sub-list and element index into filetype.Builder.DependencyIndexes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:56
type resolverByIndex interface {
	FindEnumByIndex(int32, int32, []Enum, []Message) protoreflect.EnumDescriptor
	FindMessageByIndex(int32, int32, []Enum, []Message) protoreflect.MessageDescriptor
}

// Indexes of each sub-list in filetype.Builder.DependencyIndexes.
const (
	listFieldDeps	int32	= iota
	listExtTargets
	listExtDeps
	listMethInDeps
	listMethOutDeps
)

// Out is the output of the Builder.
type Out struct {
	File	protoreflect.FileDescriptor

	// Enums is all enum descriptors in "flattened ordering".
	Enums	[]Enum
	// Messages is all message descriptors in "flattened ordering".
	// It includes the implicit message declarations for map entries.
	Messages	[]Message
	// Extensions is all extension descriptors in "flattened ordering".
	Extensions	[]Extension
	// Service is all service descriptors in "flattened ordering".
	Services	[]Service
}

// Build constructs a FileDescriptor given the parameters set in Builder.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:85
// It assumes that the inputs are well-formed and panics if any inconsistencies
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:85
// are encountered.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:85
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:85
// If NumEnums+NumMessages+NumExtensions+NumServices is zero,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:85
// then Build automatically derives them from the raw descriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:91
func (db Builder) Build() (out Out) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:91
	_go_fuzz_dep_.CoverTab[52344]++

													if db.NumEnums+db.NumMessages+db.NumExtensions+db.NumServices == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:93
		_go_fuzz_dep_.CoverTab[52349]++
														db.unmarshalCounts(db.RawDescriptor, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:94
		// _ = "end of CoverTab[52349]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:95
		_go_fuzz_dep_.CoverTab[52350]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:95
		// _ = "end of CoverTab[52350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:95
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:95
	// _ = "end of CoverTab[52344]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:95
	_go_fuzz_dep_.CoverTab[52345]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:98
	if db.TypeResolver == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:98
		_go_fuzz_dep_.CoverTab[52351]++
														db.TypeResolver = protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:99
		// _ = "end of CoverTab[52351]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:100
		_go_fuzz_dep_.CoverTab[52352]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:100
		// _ = "end of CoverTab[52352]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:100
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:100
	// _ = "end of CoverTab[52345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:100
	_go_fuzz_dep_.CoverTab[52346]++
													if db.FileRegistry == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:101
		_go_fuzz_dep_.CoverTab[52353]++
														db.FileRegistry = protoregistry.GlobalFiles
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:102
		// _ = "end of CoverTab[52353]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:103
		_go_fuzz_dep_.CoverTab[52354]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:103
		// _ = "end of CoverTab[52354]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:103
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:103
	// _ = "end of CoverTab[52346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:103
	_go_fuzz_dep_.CoverTab[52347]++

													fd := newRawFile(db)
													out.File = fd
													out.Enums = fd.allEnums
													out.Messages = fd.allMessages
													out.Extensions = fd.allExtensions
													out.Services = fd.allServices

													if err := db.FileRegistry.RegisterFile(fd); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:112
		_go_fuzz_dep_.CoverTab[52355]++
														panic(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:113
		// _ = "end of CoverTab[52355]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:114
		_go_fuzz_dep_.CoverTab[52356]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:114
		// _ = "end of CoverTab[52356]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:114
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:114
	// _ = "end of CoverTab[52347]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:114
	_go_fuzz_dep_.CoverTab[52348]++
													return out
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:115
	// _ = "end of CoverTab[52348]"
}

// unmarshalCounts counts the number of enum, message, extension, and service
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:118
// declarations in the raw message, which is either a FileDescriptorProto
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:118
// or a MessageDescriptorProto depending on whether isFile is set.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:121
func (db *Builder) unmarshalCounts(b []byte, isFile bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:121
	_go_fuzz_dep_.CoverTab[52357]++
													for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:122
		_go_fuzz_dep_.CoverTab[52358]++
														num, typ, n := protowire.ConsumeTag(b)
														b = b[n:]
														switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:126
			_go_fuzz_dep_.CoverTab[52359]++
															v, m := protowire.ConsumeBytes(b)
															b = b[m:]
															if isFile {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:129
				_go_fuzz_dep_.CoverTab[52361]++
																switch num {
				case genid.FileDescriptorProto_EnumType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:131
					_go_fuzz_dep_.CoverTab[52362]++
																	db.NumEnums++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:132
					// _ = "end of CoverTab[52362]"
				case genid.FileDescriptorProto_MessageType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:133
					_go_fuzz_dep_.CoverTab[52363]++
																	db.unmarshalCounts(v, false)
																	db.NumMessages++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:135
					// _ = "end of CoverTab[52363]"
				case genid.FileDescriptorProto_Extension_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:136
					_go_fuzz_dep_.CoverTab[52364]++
																	db.NumExtensions++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:137
					// _ = "end of CoverTab[52364]"
				case genid.FileDescriptorProto_Service_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:138
					_go_fuzz_dep_.CoverTab[52365]++
																	db.NumServices++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:139
					// _ = "end of CoverTab[52365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:139
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:139
					_go_fuzz_dep_.CoverTab[52366]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:139
					// _ = "end of CoverTab[52366]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:140
				// _ = "end of CoverTab[52361]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:141
				_go_fuzz_dep_.CoverTab[52367]++
																switch num {
				case genid.DescriptorProto_EnumType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:143
					_go_fuzz_dep_.CoverTab[52368]++
																	db.NumEnums++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:144
					// _ = "end of CoverTab[52368]"
				case genid.DescriptorProto_NestedType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:145
					_go_fuzz_dep_.CoverTab[52369]++
																	db.unmarshalCounts(v, false)
																	db.NumMessages++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:147
					// _ = "end of CoverTab[52369]"
				case genid.DescriptorProto_Extension_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:148
					_go_fuzz_dep_.CoverTab[52370]++
																	db.NumExtensions++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:149
					// _ = "end of CoverTab[52370]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:149
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:149
					_go_fuzz_dep_.CoverTab[52371]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:149
					// _ = "end of CoverTab[52371]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:150
				// _ = "end of CoverTab[52367]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:151
			// _ = "end of CoverTab[52359]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:152
			_go_fuzz_dep_.CoverTab[52360]++
															m := protowire.ConsumeFieldValue(num, typ, b)
															b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:154
			// _ = "end of CoverTab[52360]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:155
		// _ = "end of CoverTab[52358]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:156
	// _ = "end of CoverTab[52357]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:157
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/build.go:157
var _ = _go_fuzz_dep_.CoverTab
