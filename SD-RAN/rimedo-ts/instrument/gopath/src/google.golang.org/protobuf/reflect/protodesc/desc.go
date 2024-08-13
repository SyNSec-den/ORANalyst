// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:5
// Package protodesc provides functionality for converting
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:5
// FileDescriptorProto messages to/from protoreflect.FileDescriptor values.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:5
// The google.protobuf.FileDescriptorProto is a protobuf message that describes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:5
// the type information for a .proto file in a form that is easily serializable.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:5
// The protoreflect.FileDescriptor is a more structured representation of
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:5
// the FileDescriptorProto message where references and remote dependencies
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:5
// can be directly followed.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:13
package protodesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:13
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:13
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:13
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:13
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:13
)

import (
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	"google.golang.org/protobuf/types/descriptorpb"
)

// Resolver is the resolver used by NewFile to resolve dependencies.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:27
// The enums and messages provided must belong to some parent file,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:27
// which is also registered.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:27
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:27
// It is implemented by protoregistry.Files.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:32
type Resolver interface {
	FindFileByPath(string) (protoreflect.FileDescriptor, error)
	FindDescriptorByName(protoreflect.FullName) (protoreflect.Descriptor, error)
}

// FileOptions configures the construction of file descriptors.
type FileOptions struct {
	pragma.NoUnkeyedLiterals

	// AllowUnresolvable configures New to permissively allow unresolvable
	// file, enum, or message dependencies. Unresolved dependencies are replaced
	// by placeholder equivalents.
	//
	// The following dependencies may be left unresolved:
	//	• Resolving an imported file.
	//	• Resolving the type for a message field or extension field.
	//	If the kind of the field is unknown, then a placeholder is used for both
	//	the Enum and Message accessors on the protoreflect.FieldDescriptor.
	//	• Resolving an enum value set as the default for an optional enum field.
	//	If unresolvable, the protoreflect.FieldDescriptor.Default is set to the
	//	first value in the associated enum (or zero if the also enum dependency
	//	is also unresolvable). The protoreflect.FieldDescriptor.DefaultEnumValue
	//	is populated with a placeholder.
	//	• Resolving the extended message type for an extension field.
	//	• Resolving the input or output message type for a service method.
	//
	// If the unresolved dependency uses a relative name,
	// then the placeholder will contain an invalid FullName with a "*." prefix,
	// indicating that the starting prefix of the full name is unknown.
	AllowUnresolvable	bool
}

// NewFile creates a new protoreflect.FileDescriptor from the provided
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:64
// file descriptor message. See FileOptions.New for more information.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:66
func NewFile(fd *descriptorpb.FileDescriptorProto, r Resolver) (protoreflect.FileDescriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:66
	_go_fuzz_dep_.CoverTab[60292]++
													return FileOptions{}.New(fd, r)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:67
	// _ = "end of CoverTab[60292]"
}

// NewFiles creates a new protoregistry.Files from the provided
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:70
// FileDescriptorSet message. See FileOptions.NewFiles for more information.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:72
func NewFiles(fd *descriptorpb.FileDescriptorSet) (*protoregistry.Files, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:72
	_go_fuzz_dep_.CoverTab[60293]++
													return FileOptions{}.NewFiles(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:73
	// _ = "end of CoverTab[60293]"
}

// New creates a new protoreflect.FileDescriptor from the provided
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:76
// file descriptor message. The file must represent a valid proto file according
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:76
// to protobuf semantics. The returned descriptor is a deep copy of the input.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:76
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:76
// Any imported files, enum types, or message types referenced in the file are
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:76
// resolved using the provided registry. When looking up an import file path,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:76
// the path must be unique. The newly created file descriptor is not registered
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:76
// back into the provided file registry.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:84
func (o FileOptions) New(fd *descriptorpb.FileDescriptorProto, r Resolver) (protoreflect.FileDescriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:84
	_go_fuzz_dep_.CoverTab[60294]++
													if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:85
		_go_fuzz_dep_.CoverTab[60315]++
														r = (*protoregistry.Files)(nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:86
		// _ = "end of CoverTab[60315]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:87
		_go_fuzz_dep_.CoverTab[60316]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:87
		// _ = "end of CoverTab[60316]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:87
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:87
	// _ = "end of CoverTab[60294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:87
	_go_fuzz_dep_.CoverTab[60295]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:90
	f := &filedesc.File{L2: &filedesc.FileL2{}}
	switch fd.GetSyntax() {
	case "proto2", "":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:92
		_go_fuzz_dep_.CoverTab[60317]++
														f.L1.Syntax = protoreflect.Proto2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:93
		// _ = "end of CoverTab[60317]"
	case "proto3":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:94
		_go_fuzz_dep_.CoverTab[60318]++
														f.L1.Syntax = protoreflect.Proto3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:95
		// _ = "end of CoverTab[60318]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:96
		_go_fuzz_dep_.CoverTab[60319]++
														return nil, errors.New("invalid syntax: %q", fd.GetSyntax())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:97
		// _ = "end of CoverTab[60319]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:98
	// _ = "end of CoverTab[60295]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:98
	_go_fuzz_dep_.CoverTab[60296]++
													f.L1.Path = fd.GetName()
													if f.L1.Path == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:100
		_go_fuzz_dep_.CoverTab[60320]++
														return nil, errors.New("file path must be populated")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:101
		// _ = "end of CoverTab[60320]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:102
		_go_fuzz_dep_.CoverTab[60321]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:102
		// _ = "end of CoverTab[60321]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:102
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:102
	// _ = "end of CoverTab[60296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:102
	_go_fuzz_dep_.CoverTab[60297]++
													f.L1.Package = protoreflect.FullName(fd.GetPackage())
													if !f.L1.Package.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:104
		_go_fuzz_dep_.CoverTab[60322]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:104
		return f.L1.Package != ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:104
		// _ = "end of CoverTab[60322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:104
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:104
		_go_fuzz_dep_.CoverTab[60323]++
														return nil, errors.New("invalid package: %q", f.L1.Package)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:105
		// _ = "end of CoverTab[60323]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:106
		_go_fuzz_dep_.CoverTab[60324]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:106
		// _ = "end of CoverTab[60324]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:106
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:106
	// _ = "end of CoverTab[60297]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:106
	_go_fuzz_dep_.CoverTab[60298]++
													if opts := fd.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:107
		_go_fuzz_dep_.CoverTab[60325]++
														opts = proto.Clone(opts).(*descriptorpb.FileOptions)
														f.L2.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:109
			_go_fuzz_dep_.CoverTab[60326]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:109
			return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:109
			// _ = "end of CoverTab[60326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:109
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:109
		// _ = "end of CoverTab[60325]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:110
		_go_fuzz_dep_.CoverTab[60327]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:110
		// _ = "end of CoverTab[60327]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:110
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:110
	// _ = "end of CoverTab[60298]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:110
	_go_fuzz_dep_.CoverTab[60299]++

													f.L2.Imports = make(filedesc.FileImports, len(fd.GetDependency()))
													for _, i := range fd.GetPublicDependency() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:113
		_go_fuzz_dep_.CoverTab[60328]++
														if !(0 <= i && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:114
			_go_fuzz_dep_.CoverTab[60330]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:114
			return int(i) < len(f.L2.Imports)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:114
			// _ = "end of CoverTab[60330]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:114
		}()) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:114
			_go_fuzz_dep_.CoverTab[60331]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:114
			return f.L2.Imports[i].IsPublic
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:114
			// _ = "end of CoverTab[60331]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:114
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:114
			_go_fuzz_dep_.CoverTab[60332]++
															return nil, errors.New("invalid or duplicate public import index: %d", i)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:115
			// _ = "end of CoverTab[60332]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:116
			_go_fuzz_dep_.CoverTab[60333]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:116
			// _ = "end of CoverTab[60333]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:116
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:116
		// _ = "end of CoverTab[60328]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:116
		_go_fuzz_dep_.CoverTab[60329]++
														f.L2.Imports[i].IsPublic = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:117
		// _ = "end of CoverTab[60329]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:118
	// _ = "end of CoverTab[60299]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:118
	_go_fuzz_dep_.CoverTab[60300]++
													for _, i := range fd.GetWeakDependency() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:119
		_go_fuzz_dep_.CoverTab[60334]++
														if !(0 <= i && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:120
			_go_fuzz_dep_.CoverTab[60336]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:120
			return int(i) < len(f.L2.Imports)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:120
			// _ = "end of CoverTab[60336]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:120
		}()) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:120
			_go_fuzz_dep_.CoverTab[60337]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:120
			return f.L2.Imports[i].IsWeak
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:120
			// _ = "end of CoverTab[60337]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:120
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:120
			_go_fuzz_dep_.CoverTab[60338]++
															return nil, errors.New("invalid or duplicate weak import index: %d", i)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:121
			// _ = "end of CoverTab[60338]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:122
			_go_fuzz_dep_.CoverTab[60339]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:122
			// _ = "end of CoverTab[60339]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:122
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:122
		// _ = "end of CoverTab[60334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:122
		_go_fuzz_dep_.CoverTab[60335]++
														f.L2.Imports[i].IsWeak = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:123
		// _ = "end of CoverTab[60335]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:124
	// _ = "end of CoverTab[60300]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:124
	_go_fuzz_dep_.CoverTab[60301]++
													imps := importSet{f.Path(): true}
													for i, path := range fd.GetDependency() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:126
		_go_fuzz_dep_.CoverTab[60340]++
														imp := &f.L2.Imports[i]
														f, err := r.FindFileByPath(path)
														if err == protoregistry.NotFound && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:129
			_go_fuzz_dep_.CoverTab[60343]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:129
			return (o.AllowUnresolvable || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:129
				_go_fuzz_dep_.CoverTab[60344]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:129
				return imp.IsWeak
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:129
				// _ = "end of CoverTab[60344]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:129
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:129
			// _ = "end of CoverTab[60343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:129
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:129
			_go_fuzz_dep_.CoverTab[60345]++
															f = filedesc.PlaceholderFile(path)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:130
			// _ = "end of CoverTab[60345]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:131
			_go_fuzz_dep_.CoverTab[60346]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:131
			if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:131
				_go_fuzz_dep_.CoverTab[60347]++
																return nil, errors.New("could not resolve import %q: %v", path, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:132
				// _ = "end of CoverTab[60347]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:133
				_go_fuzz_dep_.CoverTab[60348]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:133
				// _ = "end of CoverTab[60348]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:133
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:133
			// _ = "end of CoverTab[60346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:133
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:133
		// _ = "end of CoverTab[60340]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:133
		_go_fuzz_dep_.CoverTab[60341]++
														imp.FileDescriptor = f

														if imps[imp.Path()] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:136
			_go_fuzz_dep_.CoverTab[60349]++
															return nil, errors.New("already imported %q", path)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:137
			// _ = "end of CoverTab[60349]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:138
			_go_fuzz_dep_.CoverTab[60350]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:138
			// _ = "end of CoverTab[60350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:138
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:138
		// _ = "end of CoverTab[60341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:138
		_go_fuzz_dep_.CoverTab[60342]++
														imps[imp.Path()] = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:139
		// _ = "end of CoverTab[60342]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:140
	// _ = "end of CoverTab[60301]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:140
	_go_fuzz_dep_.CoverTab[60302]++
													for i := range fd.GetDependency() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:141
		_go_fuzz_dep_.CoverTab[60351]++
														imp := &f.L2.Imports[i]
														imps.importPublic(imp.Imports())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:143
		// _ = "end of CoverTab[60351]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:144
	// _ = "end of CoverTab[60302]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:144
	_go_fuzz_dep_.CoverTab[60303]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:147
	f.L2.Locations.File = f
	for _, loc := range fd.GetSourceCodeInfo().GetLocation() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:148
		_go_fuzz_dep_.CoverTab[60352]++
														var l protoreflect.SourceLocation

														l.Path = protoreflect.SourcePath(loc.GetPath())
														s := loc.GetSpan()
														switch len(s) {
		case 3:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:154
			_go_fuzz_dep_.CoverTab[60355]++
															l.StartLine, l.StartColumn, l.EndLine, l.EndColumn = int(s[0]), int(s[1]), int(s[0]), int(s[2])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:155
			// _ = "end of CoverTab[60355]"
		case 4:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:156
			_go_fuzz_dep_.CoverTab[60356]++
															l.StartLine, l.StartColumn, l.EndLine, l.EndColumn = int(s[0]), int(s[1]), int(s[2]), int(s[3])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:157
			// _ = "end of CoverTab[60356]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:158
			_go_fuzz_dep_.CoverTab[60357]++
															return nil, errors.New("invalid span: %v", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:159
			// _ = "end of CoverTab[60357]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:160
		// _ = "end of CoverTab[60352]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:160
		_go_fuzz_dep_.CoverTab[60353]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
		if false && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
			_go_fuzz_dep_.CoverTab[60358]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
			return (l.EndLine < l.StartLine || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				_go_fuzz_dep_.CoverTab[60359]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				return l.StartLine < 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				// _ = "end of CoverTab[60359]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				_go_fuzz_dep_.CoverTab[60360]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				return l.StartColumn < 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				// _ = "end of CoverTab[60360]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				_go_fuzz_dep_.CoverTab[60361]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				return l.EndColumn < 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				// _ = "end of CoverTab[60361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				_go_fuzz_dep_.CoverTab[60362]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:163
				return (l.StartLine == l.EndLine && func() bool {
																	_go_fuzz_dep_.CoverTab[60363]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:164
					return l.EndColumn <= l.StartColumn
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:164
					// _ = "end of CoverTab[60363]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:164
				}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:164
				// _ = "end of CoverTab[60362]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:164
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:164
			// _ = "end of CoverTab[60358]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:164
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:164
			_go_fuzz_dep_.CoverTab[60364]++
															return nil, errors.New("invalid span: %v", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:165
			// _ = "end of CoverTab[60364]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:166
			_go_fuzz_dep_.CoverTab[60365]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:166
			// _ = "end of CoverTab[60365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:166
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:166
		// _ = "end of CoverTab[60353]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:166
		_go_fuzz_dep_.CoverTab[60354]++
														l.LeadingDetachedComments = loc.GetLeadingDetachedComments()
														l.LeadingComments = loc.GetLeadingComments()
														l.TrailingComments = loc.GetTrailingComments()
														f.L2.Locations.List = append(f.L2.Locations.List, l)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:170
		// _ = "end of CoverTab[60354]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:171
	// _ = "end of CoverTab[60303]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:171
	_go_fuzz_dep_.CoverTab[60304]++

	// Step 1: Allocate and derive the names for all declarations.
	// This copies all fields from the descriptor proto except:
	//	google.protobuf.FieldDescriptorProto.type_name
	//	google.protobuf.FieldDescriptorProto.default_value
	//	google.protobuf.FieldDescriptorProto.oneof_index
	//	google.protobuf.FieldDescriptorProto.extendee
	//	google.protobuf.MethodDescriptorProto.input
	//	google.protobuf.MethodDescriptorProto.output
	var err error
	sb := new(strs.Builder)
	r1 := make(descsByName)
	if f.L1.Enums.List, err = r1.initEnumDeclarations(fd.GetEnumType(), f, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:184
		_go_fuzz_dep_.CoverTab[60366]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:185
		// _ = "end of CoverTab[60366]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:186
		_go_fuzz_dep_.CoverTab[60367]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:186
		// _ = "end of CoverTab[60367]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:186
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:186
	// _ = "end of CoverTab[60304]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:186
	_go_fuzz_dep_.CoverTab[60305]++
													if f.L1.Messages.List, err = r1.initMessagesDeclarations(fd.GetMessageType(), f, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:187
		_go_fuzz_dep_.CoverTab[60368]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:188
		// _ = "end of CoverTab[60368]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:189
		_go_fuzz_dep_.CoverTab[60369]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:189
		// _ = "end of CoverTab[60369]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:189
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:189
	// _ = "end of CoverTab[60305]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:189
	_go_fuzz_dep_.CoverTab[60306]++
													if f.L1.Extensions.List, err = r1.initExtensionDeclarations(fd.GetExtension(), f, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:190
		_go_fuzz_dep_.CoverTab[60370]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:191
		// _ = "end of CoverTab[60370]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:192
		_go_fuzz_dep_.CoverTab[60371]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:192
		// _ = "end of CoverTab[60371]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:192
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:192
	// _ = "end of CoverTab[60306]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:192
	_go_fuzz_dep_.CoverTab[60307]++
													if f.L1.Services.List, err = r1.initServiceDeclarations(fd.GetService(), f, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:193
		_go_fuzz_dep_.CoverTab[60372]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:194
		// _ = "end of CoverTab[60372]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:195
		_go_fuzz_dep_.CoverTab[60373]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:195
		// _ = "end of CoverTab[60373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:195
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:195
	// _ = "end of CoverTab[60307]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:195
	_go_fuzz_dep_.CoverTab[60308]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:198
	r2 := &resolver{local: r1, remote: r, imports: imps, allowUnresolvable: o.AllowUnresolvable}
	if err := r2.resolveMessageDependencies(f.L1.Messages.List, fd.GetMessageType()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:199
		_go_fuzz_dep_.CoverTab[60374]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:200
		// _ = "end of CoverTab[60374]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:201
		_go_fuzz_dep_.CoverTab[60375]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:201
		// _ = "end of CoverTab[60375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:201
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:201
	// _ = "end of CoverTab[60308]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:201
	_go_fuzz_dep_.CoverTab[60309]++
													if err := r2.resolveExtensionDependencies(f.L1.Extensions.List, fd.GetExtension()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:202
		_go_fuzz_dep_.CoverTab[60376]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:203
		// _ = "end of CoverTab[60376]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:204
		_go_fuzz_dep_.CoverTab[60377]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:204
		// _ = "end of CoverTab[60377]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:204
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:204
	// _ = "end of CoverTab[60309]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:204
	_go_fuzz_dep_.CoverTab[60310]++
													if err := r2.resolveServiceDependencies(f.L1.Services.List, fd.GetService()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:205
		_go_fuzz_dep_.CoverTab[60378]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:206
		// _ = "end of CoverTab[60378]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:207
		_go_fuzz_dep_.CoverTab[60379]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:207
		// _ = "end of CoverTab[60379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:207
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:207
	// _ = "end of CoverTab[60310]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:207
	_go_fuzz_dep_.CoverTab[60311]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:210
	if err := validateEnumDeclarations(f.L1.Enums.List, fd.GetEnumType()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:210
		_go_fuzz_dep_.CoverTab[60380]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:211
		// _ = "end of CoverTab[60380]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:212
		_go_fuzz_dep_.CoverTab[60381]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:212
		// _ = "end of CoverTab[60381]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:212
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:212
	// _ = "end of CoverTab[60311]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:212
	_go_fuzz_dep_.CoverTab[60312]++
													if err := validateMessageDeclarations(f.L1.Messages.List, fd.GetMessageType()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:213
		_go_fuzz_dep_.CoverTab[60382]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:214
		// _ = "end of CoverTab[60382]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:215
		_go_fuzz_dep_.CoverTab[60383]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:215
		// _ = "end of CoverTab[60383]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:215
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:215
	// _ = "end of CoverTab[60312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:215
	_go_fuzz_dep_.CoverTab[60313]++
													if err := validateExtensionDeclarations(f.L1.Extensions.List, fd.GetExtension()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:216
		_go_fuzz_dep_.CoverTab[60384]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:217
		// _ = "end of CoverTab[60384]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:218
		_go_fuzz_dep_.CoverTab[60385]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:218
		// _ = "end of CoverTab[60385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:218
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:218
	// _ = "end of CoverTab[60313]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:218
	_go_fuzz_dep_.CoverTab[60314]++

													return f, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:220
	// _ = "end of CoverTab[60314]"
}

type importSet map[string]bool

func (is importSet) importPublic(imps protoreflect.FileImports) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:225
	_go_fuzz_dep_.CoverTab[60386]++
													for i := 0; i < imps.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:226
		_go_fuzz_dep_.CoverTab[60387]++
														if imp := imps.Get(i); imp.IsPublic {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:227
			_go_fuzz_dep_.CoverTab[60388]++
															is[imp.Path()] = true
															is.importPublic(imp.Imports())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:229
			// _ = "end of CoverTab[60388]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:230
			_go_fuzz_dep_.CoverTab[60389]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:230
			// _ = "end of CoverTab[60389]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:230
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:230
		// _ = "end of CoverTab[60387]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:231
	// _ = "end of CoverTab[60386]"
}

// NewFiles creates a new protoregistry.Files from the provided
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:234
// FileDescriptorSet message. The descriptor set must include only
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:234
// valid files according to protobuf semantics. The returned descriptors
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:234
// are a deep copy of the input.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:238
func (o FileOptions) NewFiles(fds *descriptorpb.FileDescriptorSet) (*protoregistry.Files, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:238
	_go_fuzz_dep_.CoverTab[60390]++
													files := make(map[string]*descriptorpb.FileDescriptorProto)
													for _, fd := range fds.File {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:240
		_go_fuzz_dep_.CoverTab[60393]++
														if _, ok := files[fd.GetName()]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:241
			_go_fuzz_dep_.CoverTab[60395]++
															return nil, errors.New("file appears multiple times: %q", fd.GetName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:242
			// _ = "end of CoverTab[60395]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:243
			_go_fuzz_dep_.CoverTab[60396]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:243
			// _ = "end of CoverTab[60396]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:243
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:243
		// _ = "end of CoverTab[60393]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:243
		_go_fuzz_dep_.CoverTab[60394]++
														files[fd.GetName()] = fd
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:244
		// _ = "end of CoverTab[60394]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:245
	// _ = "end of CoverTab[60390]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:245
	_go_fuzz_dep_.CoverTab[60391]++
													r := &protoregistry.Files{}
													for _, fd := range files {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:247
		_go_fuzz_dep_.CoverTab[60397]++
														if err := o.addFileDeps(r, fd, files); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:248
			_go_fuzz_dep_.CoverTab[60398]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:249
			// _ = "end of CoverTab[60398]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:250
			_go_fuzz_dep_.CoverTab[60399]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:250
			// _ = "end of CoverTab[60399]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:250
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:250
		// _ = "end of CoverTab[60397]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:251
	// _ = "end of CoverTab[60391]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:251
	_go_fuzz_dep_.CoverTab[60392]++
													return r, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:252
	// _ = "end of CoverTab[60392]"
}
func (o FileOptions) addFileDeps(r *protoregistry.Files, fd *descriptorpb.FileDescriptorProto, files map[string]*descriptorpb.FileDescriptorProto) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:254
	_go_fuzz_dep_.CoverTab[60400]++

													files[fd.GetName()] = nil
													for _, dep := range fd.Dependency {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:257
		_go_fuzz_dep_.CoverTab[60403]++
														depfd, ok := files[dep]
														if depfd == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:259
			_go_fuzz_dep_.CoverTab[60405]++
															if ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:260
				_go_fuzz_dep_.CoverTab[60407]++
																return errors.New("import cycle in file: %q", dep)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:261
				// _ = "end of CoverTab[60407]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:262
				_go_fuzz_dep_.CoverTab[60408]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:262
				// _ = "end of CoverTab[60408]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:262
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:262
			// _ = "end of CoverTab[60405]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:262
			_go_fuzz_dep_.CoverTab[60406]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:263
			// _ = "end of CoverTab[60406]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:264
			_go_fuzz_dep_.CoverTab[60409]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:264
			// _ = "end of CoverTab[60409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:264
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:264
		// _ = "end of CoverTab[60403]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:264
		_go_fuzz_dep_.CoverTab[60404]++
														if err := o.addFileDeps(r, depfd, files); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:265
			_go_fuzz_dep_.CoverTab[60410]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:266
			// _ = "end of CoverTab[60410]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:267
			_go_fuzz_dep_.CoverTab[60411]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:267
			// _ = "end of CoverTab[60411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:267
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:267
		// _ = "end of CoverTab[60404]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:268
	// _ = "end of CoverTab[60400]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:268
	_go_fuzz_dep_.CoverTab[60401]++

													delete(files, fd.GetName())
													f, err := o.New(fd, r)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:272
		_go_fuzz_dep_.CoverTab[60412]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:273
		// _ = "end of CoverTab[60412]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:274
		_go_fuzz_dep_.CoverTab[60413]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:274
		// _ = "end of CoverTab[60413]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:274
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:274
	// _ = "end of CoverTab[60401]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:274
	_go_fuzz_dep_.CoverTab[60402]++
													return r.RegisterFile(f)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:275
	// _ = "end of CoverTab[60402]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:276
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc.go:276
var _ = _go_fuzz_dep_.CoverTab
