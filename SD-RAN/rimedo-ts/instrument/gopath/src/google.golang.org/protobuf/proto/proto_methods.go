// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoreflect build tag disables use of fast-path methods.
//go:build !protoreflect
// +build !protoreflect

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:9
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:9
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:9
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:9
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:9
)

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

const hasProtoMethods = true

func protoMethods(m protoreflect.Message) *protoiface.Methods {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:18
	_go_fuzz_dep_.CoverTab[51526]++
												return m.ProtoMethods()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:19
	// _ = "end of CoverTab[51526]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:20
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto_methods.go:20
var _ = _go_fuzz_dep_.CoverTab
