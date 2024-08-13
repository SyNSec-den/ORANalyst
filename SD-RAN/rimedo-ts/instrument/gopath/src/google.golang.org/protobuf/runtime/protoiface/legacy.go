// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoiface/legacy.go:5
package protoiface

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoiface/legacy.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoiface/legacy.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoiface/legacy.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoiface/legacy.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoiface/legacy.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoiface/legacy.go:5
)

type MessageV1 interface {
	Reset()
	String() string
	ProtoMessage()
}

type ExtensionRangeV1 struct {
	Start, End int32	// both inclusive
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoiface/legacy.go:15
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/runtime/protoiface/legacy.go:15
var _ = _go_fuzz_dep_.CoverTab
