// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-protos. DO NOT EDIT.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/any_gen.go:7
package genid

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/any_gen.go:7
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/any_gen.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/any_gen.go:7
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/any_gen.go:7
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/any_gen.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/any_gen.go:7
)

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

const File_google_protobuf_any_proto = "google/protobuf/any.proto"

// Names for google.protobuf.Any.
const (
	Any_message_name	protoreflect.Name	= "Any"
	Any_message_fullname	protoreflect.FullName	= "google.protobuf.Any"
)

// Field names for google.protobuf.Any.
const (
	Any_TypeUrl_field_name	protoreflect.Name	= "type_url"
	Any_Value_field_name	protoreflect.Name	= "value"

	Any_TypeUrl_field_fullname	protoreflect.FullName	= "google.protobuf.Any.type_url"
	Any_Value_field_fullname	protoreflect.FullName	= "google.protobuf.Any.value"
)

// Field numbers for google.protobuf.Any.
const (
	Any_TypeUrl_field_number	protoreflect.FieldNumber	= 1
	Any_Value_field_number		protoreflect.FieldNumber	= 2
)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/any_gen.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/any_gen.go:34
var _ = _go_fuzz_dep_.CoverTab