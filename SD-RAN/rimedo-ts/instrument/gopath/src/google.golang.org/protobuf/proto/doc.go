// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// Package proto provides functions operating on protocol buffer messages.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// For documentation on protocol buffers in general, see:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//	https://developers.google.com/protocol-buffers
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// For a tutorial on using protocol buffers with Go, see:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//	https://developers.google.com/protocol-buffers/docs/gotutorial
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// For a guide to generated Go protocol buffer code, see:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//	https://developers.google.com/protocol-buffers/docs/reference/go-generated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// # Binary serialization
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// This package contains functions to convert to and from the wire format,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// an efficient binary serialization of protocol buffers.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Size reports the size of a message in the wire format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Marshal converts a message to the wire format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// The MarshalOptions type provides more control over wire marshaling.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Unmarshal converts a message from the wire format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// The UnmarshalOptions type provides more control over wire unmarshaling.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// # Basic message operations
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Clone makes a deep copy of a message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Merge merges the content of a message into another.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Equal compares two messages. For more control over comparisons
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// and detailed reporting of differences, see package
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// "google.golang.org/protobuf/testing/protocmp".
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Reset clears the content of a message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • CheckInitialized reports whether all required fields in a message are set.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// # Optional scalar constructors
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// The API for some generated messages represents optional scalar fields
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// as pointers to a value. For example, an optional string field has the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// Go type *string.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Bool, Int32, Int64, Uint32, Uint64, Float32, Float64, and String
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// take a value and return a pointer to a new instance of it,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// to simplify construction of optional field values.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// Generated enum types usually have an Enum method which performs the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// same operation.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// Optional scalar fields are only supported in proto2.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// # Extension accessors
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • HasExtension, GetExtension, SetExtension, and ClearExtension
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// access extension field values in a protocol buffer message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// Extension fields are only supported in proto2.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// # Related packages
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Package "google.golang.org/protobuf/encoding/protojson" converts messages to
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// and from JSON.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Package "google.golang.org/protobuf/encoding/prototext" converts messages to
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// and from the text format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Package "google.golang.org/protobuf/reflect/protoreflect" provides a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// reflection interface for protocol buffer data types.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Package "google.golang.org/protobuf/testing/protocmp" provides features
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// to compare protocol buffer messages with the "github.com/google/go-cmp/cmp"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// package.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// • Package "google.golang.org/protobuf/types/dynamicpb" provides a dynamic
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// message type, suitable for working with messages where the protocol buffer
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// type is only known at runtime.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// This module contains additional packages for more specialized use cases.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:5
// Consult the individual package documentation for details.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:89
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:89
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:89
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:89
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:89
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:89
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:89
)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/doc.go:89
var _ = _go_fuzz_dep_.CoverTab
