// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:5
package protoimpl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:5
)

import (
	"google.golang.org/protobuf/internal/version"
)

const (
	// MaxVersion is the maximum supported version for generated .pb.go files.
	// It is always the current version of the module.
	MaxVersion	= version.Minor

	// GenVersion is the runtime version required by generated .pb.go files.
	// This is incremented when generated code relies on new functionality
	// in the runtime.
	GenVersion	= 20

	// MinVersion is the minimum supported version for generated .pb.go files.
	// This is incremented when the runtime drops support for old code.
	MinVersion	= 0
)

// EnforceVersion is used by code generated by protoc-gen-go
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
// to statically enforce minimum and maximum versions of this package.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
// A compilation failure implies either that:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//   - the runtime package is too old and needs to be updated OR
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//   - the generated code is too old and needs to be regenerated.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
// The runtime package can be upgraded by running:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//	go get google.golang.org/protobuf
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
// The generated code can be regenerated by running:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//	protoc --go_out=${PROTOC_GEN_GO_ARGS} ${PROTO_FILES}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
// Example usage by generated code:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//	const (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//		// Verify that this generated code is sufficiently up-to-date.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//		_ = protoimpl.EnforceVersion(genVersion - protoimpl.MinVersion)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//		// Verify that runtime/protoimpl is sufficiently up-to-date.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//		_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - genVersion)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//	)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
// The genVersion is the current minor version used to generated the code.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
// This compile-time check relies on negative integer overflow of a uint
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:26
// being a compilation failure (guaranteed by the Go specification).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:52
type EnforceVersion uint

// This enforces the following invariant:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:54
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:54
//	MinVersion ≤ GenVersion ≤ MaxVersion
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:57
const (
	_	= EnforceVersion(GenVersion - MinVersion)
	_	= EnforceVersion(MaxVersion - GenVersion)
)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:60
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoimpl/version.go:60
var _ = _go_fuzz_dep_.CoverTab