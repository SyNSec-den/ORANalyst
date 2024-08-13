// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:5
)

// Bool stores v in a new bool value and returns a pointer to it.
func Bool(v bool) *bool	{ _go_fuzz_dep_.CoverTab[51596]++; return &v; // _ = "end of CoverTab[51596]" }

// Int32 stores v in a new int32 value and returns a pointer to it.
func Int32(v int32) *int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:11
	_go_fuzz_dep_.CoverTab[51597]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:11
	return &v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:11
	// _ = "end of CoverTab[51597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:11
}

// Int64 stores v in a new int64 value and returns a pointer to it.
func Int64(v int64) *int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:14
	_go_fuzz_dep_.CoverTab[51598]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:14
	return &v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:14
	// _ = "end of CoverTab[51598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:14
}

// Float32 stores v in a new float32 value and returns a pointer to it.
func Float32(v float32) *float32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:17
	_go_fuzz_dep_.CoverTab[51599]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:17
	return &v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:17
	// _ = "end of CoverTab[51599]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:17
}

// Float64 stores v in a new float64 value and returns a pointer to it.
func Float64(v float64) *float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:20
	_go_fuzz_dep_.CoverTab[51600]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:20
	return &v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:20
	// _ = "end of CoverTab[51600]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:20
}

// Uint32 stores v in a new uint32 value and returns a pointer to it.
func Uint32(v uint32) *uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:23
	_go_fuzz_dep_.CoverTab[51601]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:23
	return &v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:23
	// _ = "end of CoverTab[51601]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:23
}

// Uint64 stores v in a new uint64 value and returns a pointer to it.
func Uint64(v uint64) *uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:26
	_go_fuzz_dep_.CoverTab[51602]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:26
	return &v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:26
	// _ = "end of CoverTab[51602]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:26
}

// String stores v in a new string value and returns a pointer to it.
func String(v string) *string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:29
	_go_fuzz_dep_.CoverTab[51603]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:29
	return &v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:29
	// _ = "end of CoverTab[51603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:29
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/wrappers.go:29
var _ = _go_fuzz_dep_.CoverTab
