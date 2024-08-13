// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:5
// Package descopts contains the nil pointers to concrete descriptor options.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:5
// This package exists as a form of reverse dependency injection so that certain
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:5
// packages (e.g., internal/filedesc and internal/filetype can avoid a direct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:5
// dependency on the descriptor proto package).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:10
package descopts

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:10
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:10
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:10
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:10
)

import pref "google.golang.org/protobuf/reflect/protoreflect"

// These variables are set by the init function in descriptor.pb.go via logic
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:14
// in internal/filetype. In other words, so long as the descriptor proto package
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:14
// is linked in, these variables will be populated.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:14
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:14
// Each variable is populated with a nil pointer to the options struct.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:19
var (
	File		pref.ProtoMessage
	Enum		pref.ProtoMessage
	EnumValue	pref.ProtoMessage
	Message		pref.ProtoMessage
	Field		pref.ProtoMessage
	Oneof		pref.ProtoMessage
	ExtensionRange	pref.ProtoMessage
	Service		pref.ProtoMessage
	Method		pref.ProtoMessage
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descopts/options.go:29
var _ = _go_fuzz_dep_.CoverTab
