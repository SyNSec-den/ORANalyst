// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:5
)

// Bool stores v in a new bool value and returns a pointer to it.
func Bool(v bool) *bool	{ _go_fuzz_dep_.CoverTab[62414]++; return &v; // _ = "end of CoverTab[62414]" }

// Int stores v in a new int32 value and returns a pointer to it.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:10
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:10
// Deprecated: Use Int32 instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:13
func Int(v int) *int32 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:13
	_go_fuzz_dep_.CoverTab[62415]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:13
	return Int32(int32(v))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:13
	// _ = "end of CoverTab[62415]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:13
}

// Int32 stores v in a new int32 value and returns a pointer to it.
func Int32(v int32) *int32 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:16
	_go_fuzz_dep_.CoverTab[62416]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:16
	return &v
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:16
	// _ = "end of CoverTab[62416]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:16
}

// Int64 stores v in a new int64 value and returns a pointer to it.
func Int64(v int64) *int64 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:19
	_go_fuzz_dep_.CoverTab[62417]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:19
	return &v
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:19
	// _ = "end of CoverTab[62417]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:19
}

// Uint32 stores v in a new uint32 value and returns a pointer to it.
func Uint32(v uint32) *uint32 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:22
	_go_fuzz_dep_.CoverTab[62418]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:22
	return &v
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:22
	// _ = "end of CoverTab[62418]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:22
}

// Uint64 stores v in a new uint64 value and returns a pointer to it.
func Uint64(v uint64) *uint64 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:25
	_go_fuzz_dep_.CoverTab[62419]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:25
	return &v
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:25
	// _ = "end of CoverTab[62419]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:25
}

// Float32 stores v in a new float32 value and returns a pointer to it.
func Float32(v float32) *float32 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:28
	_go_fuzz_dep_.CoverTab[62420]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:28
	return &v
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:28
	// _ = "end of CoverTab[62420]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:28
}

// Float64 stores v in a new float64 value and returns a pointer to it.
func Float64(v float64) *float64 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:31
	_go_fuzz_dep_.CoverTab[62421]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:31
	return &v
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:31
	// _ = "end of CoverTab[62421]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:31
}

// String stores v in a new string value and returns a pointer to it.
func String(v string) *string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:34
	_go_fuzz_dep_.CoverTab[62422]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:34
	return &v
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:34
	// _ = "end of CoverTab[62422]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:34
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wrappers.go:34
var _ = _go_fuzz_dep_.CoverTab
