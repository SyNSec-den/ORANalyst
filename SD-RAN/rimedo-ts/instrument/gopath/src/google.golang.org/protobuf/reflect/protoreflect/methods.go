// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:5
package protoreflect

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:5
)

import (
	"google.golang.org/protobuf/internal/pragma"
)

// The following types are used by the fast-path Message.ProtoMethods method.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:11
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:11
// To avoid polluting the public protoreflect API with types used only by
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:11
// low-level implementations, the canonical definitions of these types are
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:11
// in the runtime/protoiface package. The definitions here and in protoiface
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:11
// must be kept in sync.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:17
type (
	methods	= struct {
		pragma.NoUnkeyedLiterals
		Flags			supportFlags
		Size			func(sizeInput) sizeOutput
		Marshal			func(marshalInput) (marshalOutput, error)
		Unmarshal		func(unmarshalInput) (unmarshalOutput, error)
		Merge			func(mergeInput) mergeOutput
		CheckInitialized	func(checkInitializedInput) (checkInitializedOutput, error)
	}
	supportFlags	= uint64
	sizeInput	= struct {
		pragma.NoUnkeyedLiterals
		Message	Message
		Flags	uint8
	}
	sizeOutput	= struct {
		pragma.NoUnkeyedLiterals
		Size	int
	}
	marshalInput	= struct {
		pragma.NoUnkeyedLiterals
		Message	Message
		Buf	[]byte
		Flags	uint8
	}
	marshalOutput	= struct {
		pragma.NoUnkeyedLiterals
		Buf	[]byte
	}
	unmarshalInput	= struct {
		pragma.NoUnkeyedLiterals
		Message		Message
		Buf		[]byte
		Flags		uint8
		Resolver	interface {
			FindExtensionByName(field FullName) (ExtensionType, error)
			FindExtensionByNumber(message FullName, field FieldNumber) (ExtensionType, error)
		}
		Depth	int
	}
	unmarshalOutput	= struct {
		pragma.NoUnkeyedLiterals
		Flags	uint8
	}
	mergeInput	= struct {
		pragma.NoUnkeyedLiterals
		Source		Message
		Destination	Message
	}
	mergeOutput	= struct {
		pragma.NoUnkeyedLiterals
		Flags	uint8
	}
	checkInitializedInput	= struct {
		pragma.NoUnkeyedLiterals
		Message	Message
	}
	checkInitializedOutput	= struct {
		pragma.NoUnkeyedLiterals
	}
)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:78
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/methods.go:78
var _ = _go_fuzz_dep_.CoverTab
