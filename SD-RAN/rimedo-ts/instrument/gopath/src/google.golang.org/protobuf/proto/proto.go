// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:5
)

import (
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Message is the top-level interface that all messages must implement.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
// It provides access to a reflective view of a message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
// Any implementation of this interface may be used with all functions in the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
// protobuf module that accept a Message, except where otherwise specified.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
// This is the v2 interface definition for protobuf messages.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
// The v1 interface definition is "github.com/golang/protobuf/proto".Message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
// To convert a v1 message to a v2 message,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
// use "github.com/golang/protobuf/proto".MessageV2.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
// To convert a v2 message to a v1 message,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:12
// use "github.com/golang/protobuf/proto".MessageV1.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:24
type Message = protoreflect.ProtoMessage

// Error matches all errors produced by packages in the protobuf module.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:26
// That is, errors.Is(err, Error) reports whether an error is produced
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:26
// by this module.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:30
var Error error

func init() {
	Error = errors.Error
}

// MessageName returns the full name of m.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:36
// If m is nil, it returns an empty string.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:38
func MessageName(m Message) protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:38
	_go_fuzz_dep_.CoverTab[51522]++
											if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:39
		_go_fuzz_dep_.CoverTab[51524]++
												return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:40
		// _ = "end of CoverTab[51524]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:41
		_go_fuzz_dep_.CoverTab[51525]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:41
		// _ = "end of CoverTab[51525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:41
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:41
	// _ = "end of CoverTab[51522]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:41
	_go_fuzz_dep_.CoverTab[51523]++
											return m.ProtoReflect().Descriptor().FullName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:42
	// _ = "end of CoverTab[51523]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:43
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/proto.go:43
var _ = _go_fuzz_dep_.CoverTab
