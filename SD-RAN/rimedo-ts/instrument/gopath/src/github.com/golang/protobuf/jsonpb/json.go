// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:5
// Package jsonpb provides functionality to marshal and unmarshal between a
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:5
// protocol buffer message and JSON. It follows the specification at
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:5
// https://developers.google.com/protocol-buffers/docs/proto3#json.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:5
// Do not rely on the default behavior of the standard encoding/json package
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:5
// when called on generated message types as it does not operate correctly.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:5
// Deprecated: Use the "google.golang.org/protobuf/encoding/protojson"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:5
// package instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:14
package jsonpb

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:14
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:14
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:14
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:14
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:14
)

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// AnyResolver takes a type URL, present in an Any message,
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:23
// and resolves it into an instance of the associated message.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:25
type AnyResolver interface {
	Resolve(typeURL string) (proto.Message, error)
}

type anyResolver struct{ AnyResolver }

func (r anyResolver) FindMessageByName(message protoreflect.FullName) (protoreflect.MessageType, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:31
	_go_fuzz_dep_.CoverTab[67245]++
											return r.FindMessageByURL(string(message))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:32
	// _ = "end of CoverTab[67245]"
}

func (r anyResolver) FindMessageByURL(url string) (protoreflect.MessageType, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:35
	_go_fuzz_dep_.CoverTab[67246]++
											m, err := r.Resolve(url)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:37
		_go_fuzz_dep_.CoverTab[67248]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:38
		// _ = "end of CoverTab[67248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:39
		_go_fuzz_dep_.CoverTab[67249]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:39
		// _ = "end of CoverTab[67249]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:39
	// _ = "end of CoverTab[67246]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:39
	_go_fuzz_dep_.CoverTab[67247]++
											return protoimpl.X.MessageTypeOf(m), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:40
	// _ = "end of CoverTab[67247]"
}

func (r anyResolver) FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:43
	_go_fuzz_dep_.CoverTab[67250]++
											return protoregistry.GlobalTypes.FindExtensionByName(field)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:44
	// _ = "end of CoverTab[67250]"
}

func (r anyResolver) FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:47
	_go_fuzz_dep_.CoverTab[67251]++
											return protoregistry.GlobalTypes.FindExtensionByNumber(message, field)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:48
	// _ = "end of CoverTab[67251]"
}

func wellKnownType(s protoreflect.FullName) string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:51
	_go_fuzz_dep_.CoverTab[67252]++
											if s.Parent() == "google.protobuf" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:52
		_go_fuzz_dep_.CoverTab[67254]++
												switch s.Name() {
		case "Empty", "Any",
			"BoolValue", "BytesValue", "StringValue",
			"Int32Value", "UInt32Value", "FloatValue",
			"Int64Value", "UInt64Value", "DoubleValue",
			"Duration", "Timestamp",
			"NullValue", "Struct", "Value", "ListValue":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:59
			_go_fuzz_dep_.CoverTab[67255]++
													return string(s.Name())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:60
			// _ = "end of CoverTab[67255]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:60
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:60
			_go_fuzz_dep_.CoverTab[67256]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:60
			// _ = "end of CoverTab[67256]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:61
		// _ = "end of CoverTab[67254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:62
		_go_fuzz_dep_.CoverTab[67257]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:62
		// _ = "end of CoverTab[67257]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:62
	// _ = "end of CoverTab[67252]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:62
	_go_fuzz_dep_.CoverTab[67253]++
											return ""
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:63
	// _ = "end of CoverTab[67253]"
}

func isMessageSet(md protoreflect.MessageDescriptor) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:66
	_go_fuzz_dep_.CoverTab[67258]++
											ms, ok := md.(interface{ IsMessageSet() bool })
											return ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:68
		_go_fuzz_dep_.CoverTab[67259]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:68
		return ms.IsMessageSet()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:68
		// _ = "end of CoverTab[67259]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:68
	}()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:68
	// _ = "end of CoverTab[67258]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:69
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/json.go:69
var _ = _go_fuzz_dep_.CoverTab
