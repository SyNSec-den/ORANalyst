// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2016 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:32
// Package descriptor provides functions for obtaining protocol buffer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:32
// descriptors for generated Go types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:32
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:32
// These functions cannot go in package proto because they depend on the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:32
// generated protobuf descriptor messages, which themselves depend on proto.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:37
package descriptor

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:37
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:37
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:37
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:37
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:37
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:37
)

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
)

// extractFile extracts a FileDescriptorProto from a gzip'd buffer.
func extractFile(gz []byte) (*FileDescriptorProto, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:49
	_go_fuzz_dep_.CoverTab[132334]++
														r, err := gzip.NewReader(bytes.NewReader(gz))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:51
		_go_fuzz_dep_.CoverTab[132338]++
															return nil, fmt.Errorf("failed to open gzip reader: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:52
		// _ = "end of CoverTab[132338]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:53
		_go_fuzz_dep_.CoverTab[132339]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:53
		// _ = "end of CoverTab[132339]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:53
	// _ = "end of CoverTab[132334]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:53
	_go_fuzz_dep_.CoverTab[132335]++
														defer r.Close()

														b, err := ioutil.ReadAll(r)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:57
		_go_fuzz_dep_.CoverTab[132340]++
															return nil, fmt.Errorf("failed to uncompress descriptor: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:58
		// _ = "end of CoverTab[132340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:59
		_go_fuzz_dep_.CoverTab[132341]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:59
		// _ = "end of CoverTab[132341]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:59
	// _ = "end of CoverTab[132335]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:59
	_go_fuzz_dep_.CoverTab[132336]++

														fd := new(FileDescriptorProto)
														if err := proto.Unmarshal(b, fd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:62
		_go_fuzz_dep_.CoverTab[132342]++
															return nil, fmt.Errorf("malformed FileDescriptorProto: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:63
		// _ = "end of CoverTab[132342]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:64
		_go_fuzz_dep_.CoverTab[132343]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:64
		// _ = "end of CoverTab[132343]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:64
	// _ = "end of CoverTab[132336]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:64
	_go_fuzz_dep_.CoverTab[132337]++

														return fd, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:66
	// _ = "end of CoverTab[132337]"
}

// Message is a proto.Message with a method to return its descriptor.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:69
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:69
// Message types generated by the protocol compiler always satisfy
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:69
// the Message interface.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:73
type Message interface {
	proto.Message
	Descriptor() ([]byte, []int)
}

// ForMessage returns a FileDescriptorProto and a DescriptorProto from within it
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:78
// describing the given message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:80
func ForMessage(msg Message) (fd *FileDescriptorProto, md *DescriptorProto) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:80
	_go_fuzz_dep_.CoverTab[132344]++
														gz, path := msg.Descriptor()
														fd, err := extractFile(gz)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:83
		_go_fuzz_dep_.CoverTab[132347]++
															panic(fmt.Sprintf("invalid FileDescriptorProto for %T: %v", msg, err))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:84
		// _ = "end of CoverTab[132347]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:85
		_go_fuzz_dep_.CoverTab[132348]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:85
		// _ = "end of CoverTab[132348]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:85
	// _ = "end of CoverTab[132344]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:85
	_go_fuzz_dep_.CoverTab[132345]++

														md = fd.MessageType[path[0]]
														for _, i := range path[1:] {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:88
		_go_fuzz_dep_.CoverTab[132349]++
															md = md.NestedType[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:89
		// _ = "end of CoverTab[132349]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:90
	// _ = "end of CoverTab[132345]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:90
	_go_fuzz_dep_.CoverTab[132346]++
														return fd, md
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:91
	// _ = "end of CoverTab[132346]"
}

// Is this field a scalar numeric type?
func (field *FieldDescriptorProto) IsScalar() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:95
	_go_fuzz_dep_.CoverTab[132350]++
														if field.Type == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:96
		_go_fuzz_dep_.CoverTab[132352]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:97
		// _ = "end of CoverTab[132352]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:98
		_go_fuzz_dep_.CoverTab[132353]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:98
		// _ = "end of CoverTab[132353]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:98
	// _ = "end of CoverTab[132350]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:98
	_go_fuzz_dep_.CoverTab[132351]++
														switch *field.Type {
	case FieldDescriptorProto_TYPE_DOUBLE,
		FieldDescriptorProto_TYPE_FLOAT,
		FieldDescriptorProto_TYPE_INT64,
		FieldDescriptorProto_TYPE_UINT64,
		FieldDescriptorProto_TYPE_INT32,
		FieldDescriptorProto_TYPE_FIXED64,
		FieldDescriptorProto_TYPE_FIXED32,
		FieldDescriptorProto_TYPE_BOOL,
		FieldDescriptorProto_TYPE_UINT32,
		FieldDescriptorProto_TYPE_ENUM,
		FieldDescriptorProto_TYPE_SFIXED32,
		FieldDescriptorProto_TYPE_SFIXED64,
		FieldDescriptorProto_TYPE_SINT32,
		FieldDescriptorProto_TYPE_SINT64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:113
		_go_fuzz_dep_.CoverTab[132354]++
															return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:114
		// _ = "end of CoverTab[132354]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:115
		_go_fuzz_dep_.CoverTab[132355]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:116
		// _ = "end of CoverTab[132355]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:117
	// _ = "end of CoverTab[132351]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:118
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/descriptor.go:118
var _ = _go_fuzz_dep_.CoverTab
