// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.13
// +build go1.13

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:8
package errors

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:8
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:8
)

import "errors"

// Is is errors.Is.
func Is(err, target error) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:13
	_go_fuzz_dep_.CoverTab[48326]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:13
	return errors.Is(err, target)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:13
	// _ = "end of CoverTab[48326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:13
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:13
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/errors/is_go113.go:13
var _ = _go_fuzz_dep_.CoverTab
