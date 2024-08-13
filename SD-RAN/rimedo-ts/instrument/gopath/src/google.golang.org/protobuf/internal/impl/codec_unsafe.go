// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego && !appengine
// +build !purego,!appengine

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:8
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:8
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:8
)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:12
var (
	coderEnumNoZero		= coderInt32NoZero
	coderEnum		= coderInt32
	coderEnumPtr		= coderInt32Ptr
	coderEnumSlice		= coderInt32Slice
	coderEnumPackedSlice	= coderInt32PackedSlice
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:18
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_unsafe.go:18
var _ = _go_fuzz_dep_.CoverTab
