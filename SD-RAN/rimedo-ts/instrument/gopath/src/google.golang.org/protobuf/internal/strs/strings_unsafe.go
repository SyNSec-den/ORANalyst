// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego && !appengine
// +build !purego,!appengine

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:8
package strs

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:8
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:8
)

import (
	"unsafe"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	stringHeader	struct {
		Data	unsafe.Pointer
		Len	int
	}
	sliceHeader	struct {
		Data	unsafe.Pointer
		Len	int
		Cap	int
	}
)

// UnsafeString returns an unsafe string reference of b.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:28
// The caller must treat the input slice as immutable.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:28
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:28
// WARNING: Use carefully. The returned result must not leak to the end user
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:28
// unless the input slice is provably immutable.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:33
func UnsafeString(b []byte) (s string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:33
	_go_fuzz_dep_.CoverTab[49391]++
														src := (*sliceHeader)(unsafe.Pointer(&b))
														dst := (*stringHeader)(unsafe.Pointer(&s))
														dst.Data = src.Data
														dst.Len = src.Len
														return s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:38
	// _ = "end of CoverTab[49391]"
}

// UnsafeBytes returns an unsafe bytes slice reference of s.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:41
// The caller must treat returned slice as immutable.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:41
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:41
// WARNING: Use carefully. The returned result must not leak to the end user.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:45
func UnsafeBytes(s string) (b []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:45
	_go_fuzz_dep_.CoverTab[49392]++
														src := (*stringHeader)(unsafe.Pointer(&s))
														dst := (*sliceHeader)(unsafe.Pointer(&b))
														dst.Data = src.Data
														dst.Len = src.Len
														dst.Cap = src.Len
														return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:51
	// _ = "end of CoverTab[49392]"
}

// Builder builds a set of strings with shared lifetime.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:54
// This differs from strings.Builder, which is for building a single string.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:56
type Builder struct {
	buf []byte
}

// AppendFullName is equivalent to protoreflect.FullName.Append,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:60
// but optimized for large batches where each name has a shared lifetime.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:62
func (sb *Builder) AppendFullName(prefix protoreflect.FullName, name protoreflect.Name) protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:62
	_go_fuzz_dep_.CoverTab[49393]++
														n := len(prefix) + len(".") + len(name)
														if len(prefix) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:64
		_go_fuzz_dep_.CoverTab[49395]++
															n -= len(".")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:65
		// _ = "end of CoverTab[49395]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:66
		_go_fuzz_dep_.CoverTab[49396]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:66
		// _ = "end of CoverTab[49396]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:66
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:66
	// _ = "end of CoverTab[49393]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:66
	_go_fuzz_dep_.CoverTab[49394]++
														sb.grow(n)
														sb.buf = append(sb.buf, prefix...)
														sb.buf = append(sb.buf, '.')
														sb.buf = append(sb.buf, name...)
														return protoreflect.FullName(sb.last(n))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:71
	// _ = "end of CoverTab[49394]"
}

// MakeString is equivalent to string(b), but optimized for large batches
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:74
// with a shared lifetime.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:76
func (sb *Builder) MakeString(b []byte) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:76
	_go_fuzz_dep_.CoverTab[49397]++
														sb.grow(len(b))
														sb.buf = append(sb.buf, b...)
														return sb.last(len(b))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:79
	// _ = "end of CoverTab[49397]"
}

func (sb *Builder) grow(n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:82
	_go_fuzz_dep_.CoverTab[49398]++
														if cap(sb.buf)-len(sb.buf) >= n {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:83
		_go_fuzz_dep_.CoverTab[49400]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:84
		// _ = "end of CoverTab[49400]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:85
		_go_fuzz_dep_.CoverTab[49401]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:85
		// _ = "end of CoverTab[49401]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:85
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:85
	// _ = "end of CoverTab[49398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:85
	_go_fuzz_dep_.CoverTab[49399]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:90
	sb.buf = make([]byte, 2*(cap(sb.buf)+n))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:90
	// _ = "end of CoverTab[49399]"
}

func (sb *Builder) last(n int) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:93
	_go_fuzz_dep_.CoverTab[49402]++
														return UnsafeString(sb.buf[len(sb.buf)-n:])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:94
	// _ = "end of CoverTab[49402]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:95
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/strs/strings_unsafe.go:95
var _ = _go_fuzz_dep_.CoverTab
