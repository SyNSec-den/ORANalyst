// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2018, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:29
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:29
)

type float64Value struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *float64Value) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:35
	_go_fuzz_dep_.CoverTab[114402]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:35
	*m = float64Value{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:35
	// _ = "end of CoverTab[114402]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:35
}
func (*float64Value) ProtoMessage() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:36
	_go_fuzz_dep_.CoverTab[114403]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:36
	// _ = "end of CoverTab[114403]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:36
}
func (*float64Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:37
	_go_fuzz_dep_.CoverTab[114404]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:37
	return "float64<string>"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:37
	// _ = "end of CoverTab[114404]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:37
}

type float32Value struct {
	Value float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *float32Value) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:43
	_go_fuzz_dep_.CoverTab[114405]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:43
	*m = float32Value{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:43
	// _ = "end of CoverTab[114405]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:43
}
func (*float32Value) ProtoMessage() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:44
	_go_fuzz_dep_.CoverTab[114406]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:44
	// _ = "end of CoverTab[114406]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:44
}
func (*float32Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:45
	_go_fuzz_dep_.CoverTab[114407]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:45
	return "float32<string>"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:45
	// _ = "end of CoverTab[114407]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:45
}

type int64Value struct {
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *int64Value) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:51
	_go_fuzz_dep_.CoverTab[114408]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:51
	*m = int64Value{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:51
	// _ = "end of CoverTab[114408]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:51
}
func (*int64Value) ProtoMessage()	{ _go_fuzz_dep_.CoverTab[114409]++; // _ = "end of CoverTab[114409]" }
func (*int64Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:53
	_go_fuzz_dep_.CoverTab[114410]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:53
	return "int64<string>"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:53
	// _ = "end of CoverTab[114410]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:53
}

type uint64Value struct {
	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *uint64Value) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:59
	_go_fuzz_dep_.CoverTab[114411]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:59
	*m = uint64Value{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:59
	// _ = "end of CoverTab[114411]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:59
}
func (*uint64Value) ProtoMessage() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:60
	_go_fuzz_dep_.CoverTab[114412]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:60
	// _ = "end of CoverTab[114412]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:60
}
func (*uint64Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:61
	_go_fuzz_dep_.CoverTab[114413]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:61
	return "uint64<string>"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:61
	// _ = "end of CoverTab[114413]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:61
}

type int32Value struct {
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *int32Value) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:67
	_go_fuzz_dep_.CoverTab[114414]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:67
	*m = int32Value{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:67
	// _ = "end of CoverTab[114414]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:67
}
func (*int32Value) ProtoMessage()	{ _go_fuzz_dep_.CoverTab[114415]++; // _ = "end of CoverTab[114415]" }
func (*int32Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:69
	_go_fuzz_dep_.CoverTab[114416]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:69
	return "int32<string>"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:69
	// _ = "end of CoverTab[114416]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:69
}

type uint32Value struct {
	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *uint32Value) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:75
	_go_fuzz_dep_.CoverTab[114417]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:75
	*m = uint32Value{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:75
	// _ = "end of CoverTab[114417]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:75
}
func (*uint32Value) ProtoMessage() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:76
	_go_fuzz_dep_.CoverTab[114418]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:76
	// _ = "end of CoverTab[114418]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:76
}
func (*uint32Value) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:77
	_go_fuzz_dep_.CoverTab[114419]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:77
	return "uint32<string>"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:77
	// _ = "end of CoverTab[114419]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:77
}

type boolValue struct {
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *boolValue) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:83
	_go_fuzz_dep_.CoverTab[114420]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:83
	*m = boolValue{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:83
	// _ = "end of CoverTab[114420]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:83
}
func (*boolValue) ProtoMessage()	{ _go_fuzz_dep_.CoverTab[114421]++; // _ = "end of CoverTab[114421]" }
func (*boolValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:85
	_go_fuzz_dep_.CoverTab[114422]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:85
	return "bool<string>"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:85
	// _ = "end of CoverTab[114422]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:85
}

type stringValue struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *stringValue) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:91
	_go_fuzz_dep_.CoverTab[114423]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:91
	*m = stringValue{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:91
	// _ = "end of CoverTab[114423]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:91
}
func (*stringValue) ProtoMessage() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:92
	_go_fuzz_dep_.CoverTab[114424]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:92
	// _ = "end of CoverTab[114424]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:92
}
func (*stringValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:93
	_go_fuzz_dep_.CoverTab[114425]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:93
	return "string<string>"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:93
	// _ = "end of CoverTab[114425]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:93
}

type bytesValue struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *bytesValue) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:99
	_go_fuzz_dep_.CoverTab[114426]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:99
	*m = bytesValue{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:99
	// _ = "end of CoverTab[114426]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:99
}
func (*bytesValue) ProtoMessage()	{ _go_fuzz_dep_.CoverTab[114427]++; // _ = "end of CoverTab[114427]" }
func (*bytesValue) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:101
	_go_fuzz_dep_.CoverTab[114428]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:101
	return "[]byte<string>"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:101
	// _ = "end of CoverTab[114428]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:101
}

func init() {
	RegisterType((*float64Value)(nil), "gogo.protobuf.proto.DoubleValue")
	RegisterType((*float32Value)(nil), "gogo.protobuf.proto.FloatValue")
	RegisterType((*int64Value)(nil), "gogo.protobuf.proto.Int64Value")
	RegisterType((*uint64Value)(nil), "gogo.protobuf.proto.UInt64Value")
	RegisterType((*int32Value)(nil), "gogo.protobuf.proto.Int32Value")
	RegisterType((*uint32Value)(nil), "gogo.protobuf.proto.UInt32Value")
	RegisterType((*boolValue)(nil), "gogo.protobuf.proto.BoolValue")
	RegisterType((*stringValue)(nil), "gogo.protobuf.proto.StringValue")
	RegisterType((*bytesValue)(nil), "gogo.protobuf.proto.BytesValue")
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers_gogo.go:113
var _ = _go_fuzz_dep_.CoverTab
