// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/rpc/status.proto

//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:21
package status

//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:21
)

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_	= protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_	= protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The `Status` type defines a logical error model that is suitable for
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:39
// different programming environments, including REST APIs and RPC APIs. It is
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:39
// used by [gRPC](https://github.com/grpc). Each `Status` message contains
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:39
// three pieces of data: error code, error message, and error details.
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:39
//
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:39
// You can find out more about this error model and how to work with it in the
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:39
// [API Design Guide](https://cloud.google.com/apis/design/errors).
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:46
type Status struct {
	state		protoimpl.MessageState
	sizeCache	protoimpl.SizeCache
	unknownFields	protoimpl.UnknownFields

	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code	int32	`protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	// by the client.
	Message	string	`protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A list of messages that carry the error details.  There is a common set of
	// message types for APIs to use.
	Details	[]*anypb.Any	`protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Status) Reset() {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:64
	_go_fuzz_dep_.CoverTab[68454]++
																	*x = Status{}
																	if protoimpl.UnsafeEnabled {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:66
		_go_fuzz_dep_.CoverTab[68455]++
																		mi := &file_google_rpc_status_proto_msgTypes[0]
																		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
																		ms.StoreMessageInfo(mi)
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:69
		// _ = "end of CoverTab[68455]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:70
		_go_fuzz_dep_.CoverTab[68456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:70
		// _ = "end of CoverTab[68456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:70
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:70
	// _ = "end of CoverTab[68454]"
}

func (x *Status) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:73
	_go_fuzz_dep_.CoverTab[68457]++
																	return protoimpl.X.MessageStringOf(x)
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:74
	// _ = "end of CoverTab[68457]"
}

func (*Status) ProtoMessage()	{ _go_fuzz_dep_.CoverTab[68458]++; // _ = "end of CoverTab[68458]" }

func (x *Status) ProtoReflect() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:79
	_go_fuzz_dep_.CoverTab[68459]++
																	mi := &file_google_rpc_status_proto_msgTypes[0]
																	if protoimpl.UnsafeEnabled && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:81
		_go_fuzz_dep_.CoverTab[68461]++
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:81
		return x != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:81
		// _ = "end of CoverTab[68461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:81
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:81
		_go_fuzz_dep_.CoverTab[68462]++
																		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
																		if ms.LoadMessageInfo() == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:83
			_go_fuzz_dep_.CoverTab[68464]++
																			ms.StoreMessageInfo(mi)
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:84
			// _ = "end of CoverTab[68464]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:85
			_go_fuzz_dep_.CoverTab[68465]++
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:85
			// _ = "end of CoverTab[68465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:85
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:85
		// _ = "end of CoverTab[68462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:85
		_go_fuzz_dep_.CoverTab[68463]++
																		return ms
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:86
		// _ = "end of CoverTab[68463]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:87
		_go_fuzz_dep_.CoverTab[68466]++
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:87
		// _ = "end of CoverTab[68466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:87
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:87
	// _ = "end of CoverTab[68459]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:87
	_go_fuzz_dep_.CoverTab[68460]++
																	return mi.MessageOf(x)
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:88
	// _ = "end of CoverTab[68460]"
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:92
	_go_fuzz_dep_.CoverTab[68467]++
																	return file_google_rpc_status_proto_rawDescGZIP(), []int{0}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:93
	// _ = "end of CoverTab[68467]"
}

func (x *Status) GetCode() int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:96
	_go_fuzz_dep_.CoverTab[68468]++
																	if x != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:97
		_go_fuzz_dep_.CoverTab[68470]++
																		return x.Code
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:98
		// _ = "end of CoverTab[68470]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:99
		_go_fuzz_dep_.CoverTab[68471]++
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:99
		// _ = "end of CoverTab[68471]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:99
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:99
	// _ = "end of CoverTab[68468]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:99
	_go_fuzz_dep_.CoverTab[68469]++
																	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:100
	// _ = "end of CoverTab[68469]"
}

func (x *Status) GetMessage() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:103
	_go_fuzz_dep_.CoverTab[68472]++
																	if x != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:104
		_go_fuzz_dep_.CoverTab[68474]++
																		return x.Message
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:105
		// _ = "end of CoverTab[68474]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:106
		_go_fuzz_dep_.CoverTab[68475]++
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:106
		// _ = "end of CoverTab[68475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:106
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:106
	// _ = "end of CoverTab[68472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:106
	_go_fuzz_dep_.CoverTab[68473]++
																	return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:107
	// _ = "end of CoverTab[68473]"
}

func (x *Status) GetDetails() []*anypb.Any {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:110
	_go_fuzz_dep_.CoverTab[68476]++
																	if x != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:111
		_go_fuzz_dep_.CoverTab[68478]++
																		return x.Details
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:112
		// _ = "end of CoverTab[68478]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:113
		_go_fuzz_dep_.CoverTab[68479]++
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:113
		// _ = "end of CoverTab[68479]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:113
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:113
	// _ = "end of CoverTab[68476]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:113
	_go_fuzz_dep_.CoverTab[68477]++
																	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:114
	// _ = "end of CoverTab[68477]"
}

var File_google_rpc_status_proto protoreflect.FileDescriptor

var file_google_rpc_status_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x66, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x61, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_rpc_status_proto_rawDescOnce	sync.Once
	file_google_rpc_status_proto_rawDescData	= file_google_rpc_status_proto_rawDesc
)

func file_google_rpc_status_proto_rawDescGZIP() []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:145
	_go_fuzz_dep_.CoverTab[68480]++
																	file_google_rpc_status_proto_rawDescOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:146
		_go_fuzz_dep_.CoverTab[68482]++
																		file_google_rpc_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_rpc_status_proto_rawDescData)
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:147
		// _ = "end of CoverTab[68482]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:148
	// _ = "end of CoverTab[68480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:148
	_go_fuzz_dep_.CoverTab[68481]++
																	return file_google_rpc_status_proto_rawDescData
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:149
	// _ = "end of CoverTab[68481]"
}

var file_google_rpc_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_rpc_status_proto_goTypes = []interface{}{
	(*Status)(nil),
	(*anypb.Any)(nil),
}
var file_google_rpc_status_proto_depIdxs = []int32{
	1,
	1,
	1,
	1,
	1,
	0,
}

func init()	{ file_google_rpc_status_proto_init() }
func file_google_rpc_status_proto_init() {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:167
	_go_fuzz_dep_.CoverTab[68483]++
																	if File_google_rpc_status_proto != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:168
		_go_fuzz_dep_.CoverTab[68486]++
																		return
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:169
		// _ = "end of CoverTab[68486]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:170
		_go_fuzz_dep_.CoverTab[68487]++
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:170
		// _ = "end of CoverTab[68487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:170
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:170
	// _ = "end of CoverTab[68483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:170
	_go_fuzz_dep_.CoverTab[68484]++
																	if !protoimpl.UnsafeEnabled {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:171
		_go_fuzz_dep_.CoverTab[68488]++
																		file_google_rpc_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:172
			_go_fuzz_dep_.CoverTab[68489]++
																			switch v := v.(*Status); i {
			case 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:174
				_go_fuzz_dep_.CoverTab[68490]++
																				return &v.state
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:175
				// _ = "end of CoverTab[68490]"
			case 1:
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:176
				_go_fuzz_dep_.CoverTab[68491]++
																				return &v.sizeCache
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:177
				// _ = "end of CoverTab[68491]"
			case 2:
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:178
				_go_fuzz_dep_.CoverTab[68492]++
																				return &v.unknownFields
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:179
				// _ = "end of CoverTab[68492]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:180
				_go_fuzz_dep_.CoverTab[68493]++
																				return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:181
				// _ = "end of CoverTab[68493]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:182
			// _ = "end of CoverTab[68489]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:183
		// _ = "end of CoverTab[68488]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:184
		_go_fuzz_dep_.CoverTab[68494]++
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:184
		// _ = "end of CoverTab[68494]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:184
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:184
	// _ = "end of CoverTab[68484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:184
	_go_fuzz_dep_.CoverTab[68485]++
																	type x struct{}
																	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath:	reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor:	file_google_rpc_status_proto_rawDesc,
			NumEnums:	0,
			NumMessages:	1,
			NumExtensions:	0,
			NumServices:	0,
		},
		GoTypes:		file_google_rpc_status_proto_goTypes,
		DependencyIndexes:	file_google_rpc_status_proto_depIdxs,
		MessageInfos:		file_google_rpc_status_proto_msgTypes,
	}.Build()
																	File_google_rpc_status_proto = out.File
																	file_google_rpc_status_proto_rawDesc = nil
																	file_google_rpc_status_proto_goTypes = nil
																	file_google_rpc_status_proto_depIdxs = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:202
	// _ = "end of CoverTab[68485]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:203
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/genproto@v0.0.0-20230110181048-76db0878b65f/googleapis/rpc/status/status.pb.go:203
var _ = _go_fuzz_dep_.CoverTab