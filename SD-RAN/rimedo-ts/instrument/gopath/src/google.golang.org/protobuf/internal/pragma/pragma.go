// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:5
// Package pragma provides types that can be embedded into a struct to
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:5
// statically enforce or prevent certain language properties.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:7
package pragma

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:7
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:7
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:7
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:7
)

import "sync"

// NoUnkeyedLiterals can be embedded in a struct to prevent unkeyed literals.
type NoUnkeyedLiterals struct{}

// DoNotImplement can be embedded in an interface to prevent trivial
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:14
// implementations of the interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:14
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:14
// This is useful to prevent unauthorized implementations of an interface
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:14
// so that it can be extended in the future for any protobuf language changes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:19
type DoNotImplement interface{ ProtoInternal(DoNotImplement) }

// DoNotCompare can be embedded in a struct to prevent comparability.
type DoNotCompare [0]func()

// DoNotCopy can be embedded in a struct to help prevent shallow copies.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:24
// This does not rely on a Go language feature, but rather a special case
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:24
// within the vet checker.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:24
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:24
// See https://golang.org/issues/8005.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:29
type DoNotCopy [0]sync.Mutex

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/pragma/pragma.go:29
var _ = _go_fuzz_dep_.CoverTab
