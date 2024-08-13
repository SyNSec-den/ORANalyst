// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:5
// Package flags provides a set of flags controlled by build tags.
package flags

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:6
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:6
)

// ProtoLegacy specifies whether to enable support for legacy functionality
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:8
// such as MessageSets, weak fields, and various other obscure behavior
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:8
// that is necessary to maintain backwards compatibility with proto1 or
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:8
// the pre-release variants of proto2 and proto3.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:8
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:8
// This is disabled by default unless built with the "protolegacy" tag.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:8
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:8
// WARNING: The compatibility agreement covers nothing provided by this flag.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:8
// As such, functionality may suddenly be removed or changed at our discretion.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:17
const ProtoLegacy = protoLegacy

// LazyUnmarshalExtensions specifies whether to lazily unmarshal extensions.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:19
// Lazy extension unmarshaling validates the contents of message-valued
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:19
// extension fields at unmarshal time, but defers creating the message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:19
// structure until the extension is first accessed.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:24
const LazyUnmarshalExtensions = ProtoLegacy

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:24
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/flags/flags.go:24
var _ = _go_fuzz_dep_.CoverTab
