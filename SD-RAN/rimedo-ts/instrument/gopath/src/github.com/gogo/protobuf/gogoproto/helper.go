// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:29
package gogoproto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:29
)

import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import proto "github.com/gogo/protobuf/proto"

func IsEmbed(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:34
	_go_fuzz_dep_.CoverTab[133999]++
												return proto.GetBoolExtension(field.Options, E_Embed, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:35
	// _ = "end of CoverTab[133999]"
}

func IsNullable(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:38
	_go_fuzz_dep_.CoverTab[134000]++
												return proto.GetBoolExtension(field.Options, E_Nullable, true)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:39
	// _ = "end of CoverTab[134000]"
}

func IsStdTime(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:42
	_go_fuzz_dep_.CoverTab[134001]++
												return proto.GetBoolExtension(field.Options, E_Stdtime, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:43
	// _ = "end of CoverTab[134001]"
}

func IsStdDuration(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:46
	_go_fuzz_dep_.CoverTab[134002]++
												return proto.GetBoolExtension(field.Options, E_Stdduration, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:47
	// _ = "end of CoverTab[134002]"
}

func IsStdDouble(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:50
	_go_fuzz_dep_.CoverTab[134003]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:51
		_go_fuzz_dep_.CoverTab[134004]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:51
		return *field.TypeName == ".google.protobuf.DoubleValue"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:51
		// _ = "end of CoverTab[134004]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:51
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:51
	// _ = "end of CoverTab[134003]"
}

func IsStdFloat(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:54
	_go_fuzz_dep_.CoverTab[134005]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:55
		_go_fuzz_dep_.CoverTab[134006]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:55
		return *field.TypeName == ".google.protobuf.FloatValue"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:55
		// _ = "end of CoverTab[134006]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:55
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:55
	// _ = "end of CoverTab[134005]"
}

func IsStdInt64(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:58
	_go_fuzz_dep_.CoverTab[134007]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:59
		_go_fuzz_dep_.CoverTab[134008]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:59
		return *field.TypeName == ".google.protobuf.Int64Value"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:59
		// _ = "end of CoverTab[134008]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:59
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:59
	// _ = "end of CoverTab[134007]"
}

func IsStdUInt64(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:62
	_go_fuzz_dep_.CoverTab[134009]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:63
		_go_fuzz_dep_.CoverTab[134010]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:63
		return *field.TypeName == ".google.protobuf.UInt64Value"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:63
		// _ = "end of CoverTab[134010]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:63
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:63
	// _ = "end of CoverTab[134009]"
}

func IsStdInt32(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:66
	_go_fuzz_dep_.CoverTab[134011]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:67
		_go_fuzz_dep_.CoverTab[134012]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:67
		return *field.TypeName == ".google.protobuf.Int32Value"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:67
		// _ = "end of CoverTab[134012]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:67
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:67
	// _ = "end of CoverTab[134011]"
}

func IsStdUInt32(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:70
	_go_fuzz_dep_.CoverTab[134013]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:71
		_go_fuzz_dep_.CoverTab[134014]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:71
		return *field.TypeName == ".google.protobuf.UInt32Value"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:71
		// _ = "end of CoverTab[134014]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:71
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:71
	// _ = "end of CoverTab[134013]"
}

func IsStdBool(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:74
	_go_fuzz_dep_.CoverTab[134015]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:75
		_go_fuzz_dep_.CoverTab[134016]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:75
		return *field.TypeName == ".google.protobuf.BoolValue"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:75
		// _ = "end of CoverTab[134016]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:75
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:75
	// _ = "end of CoverTab[134015]"
}

func IsStdString(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:78
	_go_fuzz_dep_.CoverTab[134017]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:79
		_go_fuzz_dep_.CoverTab[134018]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:79
		return *field.TypeName == ".google.protobuf.StringValue"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:79
		// _ = "end of CoverTab[134018]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:79
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:79
	// _ = "end of CoverTab[134017]"
}

func IsStdBytes(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:82
	_go_fuzz_dep_.CoverTab[134019]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:83
		_go_fuzz_dep_.CoverTab[134020]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:83
		return *field.TypeName == ".google.protobuf.BytesValue"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:83
		// _ = "end of CoverTab[134020]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:83
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:83
	// _ = "end of CoverTab[134019]"
}

func IsStdType(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:86
	_go_fuzz_dep_.CoverTab[134021]++
												return (IsStdTime(field) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:87
		_go_fuzz_dep_.CoverTab[134022]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:87
		return IsStdDuration(field)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:87
		// _ = "end of CoverTab[134022]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:87
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:87
		_go_fuzz_dep_.CoverTab[134023]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:87
		return IsStdDouble(field)
													// _ = "end of CoverTab[134023]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:88
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:88
		_go_fuzz_dep_.CoverTab[134024]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:88
		return IsStdFloat(field)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:88
		// _ = "end of CoverTab[134024]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:88
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:88
		_go_fuzz_dep_.CoverTab[134025]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:88
		return IsStdInt64(field)
													// _ = "end of CoverTab[134025]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:89
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:89
		_go_fuzz_dep_.CoverTab[134026]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:89
		return IsStdUInt64(field)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:89
		// _ = "end of CoverTab[134026]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:89
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:89
		_go_fuzz_dep_.CoverTab[134027]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:89
		return IsStdInt32(field)
													// _ = "end of CoverTab[134027]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:90
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:90
		_go_fuzz_dep_.CoverTab[134028]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:90
		return IsStdUInt32(field)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:90
		// _ = "end of CoverTab[134028]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:90
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:90
		_go_fuzz_dep_.CoverTab[134029]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:90
		return IsStdBool(field)
													// _ = "end of CoverTab[134029]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:91
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:91
		_go_fuzz_dep_.CoverTab[134030]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:91
		return IsStdString(field)
													// _ = "end of CoverTab[134030]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:92
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:92
		_go_fuzz_dep_.CoverTab[134031]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:92
		return IsStdBytes(field)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:92
		// _ = "end of CoverTab[134031]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:92
	}())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:92
	// _ = "end of CoverTab[134021]"
}

func IsWktPtr(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:95
	_go_fuzz_dep_.CoverTab[134032]++
												return proto.GetBoolExtension(field.Options, E_Wktpointer, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:96
	// _ = "end of CoverTab[134032]"
}

func NeedsNilCheck(proto3 bool, field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:99
	_go_fuzz_dep_.CoverTab[134033]++
												nullable := IsNullable(field)
												if field.IsMessage() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:101
		_go_fuzz_dep_.CoverTab[134036]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:101
		return IsCustomType(field)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:101
		// _ = "end of CoverTab[134036]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:101
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:101
		_go_fuzz_dep_.CoverTab[134037]++
													return nullable
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:102
		// _ = "end of CoverTab[134037]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:103
		_go_fuzz_dep_.CoverTab[134038]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:103
		// _ = "end of CoverTab[134038]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:103
	// _ = "end of CoverTab[134033]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:103
	_go_fuzz_dep_.CoverTab[134034]++
												if proto3 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:104
		_go_fuzz_dep_.CoverTab[134039]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:105
		// _ = "end of CoverTab[134039]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:106
		_go_fuzz_dep_.CoverTab[134040]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:106
		// _ = "end of CoverTab[134040]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:106
	// _ = "end of CoverTab[134034]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:106
	_go_fuzz_dep_.CoverTab[134035]++
												return nullable || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:107
		_go_fuzz_dep_.CoverTab[134041]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:107
		return *field.Type == google_protobuf.FieldDescriptorProto_TYPE_BYTES
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:107
		// _ = "end of CoverTab[134041]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:107
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:107
	// _ = "end of CoverTab[134035]"
}

func IsCustomType(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:110
	_go_fuzz_dep_.CoverTab[134042]++
												typ := GetCustomType(field)
												if len(typ) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:112
		_go_fuzz_dep_.CoverTab[134044]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:113
		// _ = "end of CoverTab[134044]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:114
		_go_fuzz_dep_.CoverTab[134045]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:114
		// _ = "end of CoverTab[134045]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:114
	// _ = "end of CoverTab[134042]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:114
	_go_fuzz_dep_.CoverTab[134043]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:115
	// _ = "end of CoverTab[134043]"
}

func IsCastType(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:118
	_go_fuzz_dep_.CoverTab[134046]++
												typ := GetCastType(field)
												if len(typ) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:120
		_go_fuzz_dep_.CoverTab[134048]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:121
		// _ = "end of CoverTab[134048]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:122
		_go_fuzz_dep_.CoverTab[134049]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:122
		// _ = "end of CoverTab[134049]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:122
	// _ = "end of CoverTab[134046]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:122
	_go_fuzz_dep_.CoverTab[134047]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:123
	// _ = "end of CoverTab[134047]"
}

func IsCastKey(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:126
	_go_fuzz_dep_.CoverTab[134050]++
												typ := GetCastKey(field)
												if len(typ) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:128
		_go_fuzz_dep_.CoverTab[134052]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:129
		// _ = "end of CoverTab[134052]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:130
		_go_fuzz_dep_.CoverTab[134053]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:130
		// _ = "end of CoverTab[134053]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:130
	// _ = "end of CoverTab[134050]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:130
	_go_fuzz_dep_.CoverTab[134051]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:131
	// _ = "end of CoverTab[134051]"
}

func IsCastValue(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:134
	_go_fuzz_dep_.CoverTab[134054]++
												typ := GetCastValue(field)
												if len(typ) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:136
		_go_fuzz_dep_.CoverTab[134056]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:137
		// _ = "end of CoverTab[134056]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:138
		_go_fuzz_dep_.CoverTab[134057]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:138
		// _ = "end of CoverTab[134057]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:138
	// _ = "end of CoverTab[134054]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:138
	_go_fuzz_dep_.CoverTab[134055]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:139
	// _ = "end of CoverTab[134055]"
}

func HasEnumDecl(file *google_protobuf.FileDescriptorProto, enum *google_protobuf.EnumDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:142
	_go_fuzz_dep_.CoverTab[134058]++
												return proto.GetBoolExtension(enum.Options, E_Enumdecl, proto.GetBoolExtension(file.Options, E_EnumdeclAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:143
	// _ = "end of CoverTab[134058]"
}

func HasTypeDecl(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:146
	_go_fuzz_dep_.CoverTab[134059]++
												return proto.GetBoolExtension(message.Options, E_Typedecl, proto.GetBoolExtension(file.Options, E_TypedeclAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:147
	// _ = "end of CoverTab[134059]"
}

func GetCustomType(field *google_protobuf.FieldDescriptorProto) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:150
	_go_fuzz_dep_.CoverTab[134060]++
												if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:151
		_go_fuzz_dep_.CoverTab[134063]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:152
		// _ = "end of CoverTab[134063]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:153
		_go_fuzz_dep_.CoverTab[134064]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:153
		// _ = "end of CoverTab[134064]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:153
	// _ = "end of CoverTab[134060]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:153
	_go_fuzz_dep_.CoverTab[134061]++
												if field.Options != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:154
		_go_fuzz_dep_.CoverTab[134065]++
													v, err := proto.GetExtension(field.Options, E_Customtype)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:156
			_go_fuzz_dep_.CoverTab[134066]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:156
			return v.(*string) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:156
			// _ = "end of CoverTab[134066]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:156
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:156
			_go_fuzz_dep_.CoverTab[134067]++
														return *(v.(*string))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:157
			// _ = "end of CoverTab[134067]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:158
			_go_fuzz_dep_.CoverTab[134068]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:158
			// _ = "end of CoverTab[134068]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:158
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:158
		// _ = "end of CoverTab[134065]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:159
		_go_fuzz_dep_.CoverTab[134069]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:159
		// _ = "end of CoverTab[134069]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:159
	// _ = "end of CoverTab[134061]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:159
	_go_fuzz_dep_.CoverTab[134062]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:160
	// _ = "end of CoverTab[134062]"
}

func GetCastType(field *google_protobuf.FieldDescriptorProto) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:163
	_go_fuzz_dep_.CoverTab[134070]++
												if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:164
		_go_fuzz_dep_.CoverTab[134073]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:165
		// _ = "end of CoverTab[134073]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:166
		_go_fuzz_dep_.CoverTab[134074]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:166
		// _ = "end of CoverTab[134074]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:166
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:166
	// _ = "end of CoverTab[134070]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:166
	_go_fuzz_dep_.CoverTab[134071]++
												if field.Options != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:167
		_go_fuzz_dep_.CoverTab[134075]++
													v, err := proto.GetExtension(field.Options, E_Casttype)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:169
			_go_fuzz_dep_.CoverTab[134076]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:169
			return v.(*string) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:169
			// _ = "end of CoverTab[134076]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:169
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:169
			_go_fuzz_dep_.CoverTab[134077]++
														return *(v.(*string))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:170
			// _ = "end of CoverTab[134077]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:171
			_go_fuzz_dep_.CoverTab[134078]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:171
			// _ = "end of CoverTab[134078]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:171
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:171
		// _ = "end of CoverTab[134075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:172
		_go_fuzz_dep_.CoverTab[134079]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:172
		// _ = "end of CoverTab[134079]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:172
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:172
	// _ = "end of CoverTab[134071]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:172
	_go_fuzz_dep_.CoverTab[134072]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:173
	// _ = "end of CoverTab[134072]"
}

func GetCastKey(field *google_protobuf.FieldDescriptorProto) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:176
	_go_fuzz_dep_.CoverTab[134080]++
												if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:177
		_go_fuzz_dep_.CoverTab[134083]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:178
		// _ = "end of CoverTab[134083]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:179
		_go_fuzz_dep_.CoverTab[134084]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:179
		// _ = "end of CoverTab[134084]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:179
	// _ = "end of CoverTab[134080]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:179
	_go_fuzz_dep_.CoverTab[134081]++
												if field.Options != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:180
		_go_fuzz_dep_.CoverTab[134085]++
													v, err := proto.GetExtension(field.Options, E_Castkey)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:182
			_go_fuzz_dep_.CoverTab[134086]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:182
			return v.(*string) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:182
			// _ = "end of CoverTab[134086]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:182
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:182
			_go_fuzz_dep_.CoverTab[134087]++
														return *(v.(*string))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:183
			// _ = "end of CoverTab[134087]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:184
			_go_fuzz_dep_.CoverTab[134088]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:184
			// _ = "end of CoverTab[134088]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:184
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:184
		// _ = "end of CoverTab[134085]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:185
		_go_fuzz_dep_.CoverTab[134089]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:185
		// _ = "end of CoverTab[134089]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:185
	// _ = "end of CoverTab[134081]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:185
	_go_fuzz_dep_.CoverTab[134082]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:186
	// _ = "end of CoverTab[134082]"
}

func GetCastValue(field *google_protobuf.FieldDescriptorProto) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:189
	_go_fuzz_dep_.CoverTab[134090]++
												if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:190
		_go_fuzz_dep_.CoverTab[134093]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:191
		// _ = "end of CoverTab[134093]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:192
		_go_fuzz_dep_.CoverTab[134094]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:192
		// _ = "end of CoverTab[134094]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:192
	// _ = "end of CoverTab[134090]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:192
	_go_fuzz_dep_.CoverTab[134091]++
												if field.Options != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:193
		_go_fuzz_dep_.CoverTab[134095]++
													v, err := proto.GetExtension(field.Options, E_Castvalue)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:195
			_go_fuzz_dep_.CoverTab[134096]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:195
			return v.(*string) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:195
			// _ = "end of CoverTab[134096]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:195
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:195
			_go_fuzz_dep_.CoverTab[134097]++
														return *(v.(*string))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:196
			// _ = "end of CoverTab[134097]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:197
			_go_fuzz_dep_.CoverTab[134098]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:197
			// _ = "end of CoverTab[134098]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:197
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:197
		// _ = "end of CoverTab[134095]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:198
		_go_fuzz_dep_.CoverTab[134099]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:198
		// _ = "end of CoverTab[134099]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:198
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:198
	// _ = "end of CoverTab[134091]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:198
	_go_fuzz_dep_.CoverTab[134092]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:199
	// _ = "end of CoverTab[134092]"
}

func IsCustomName(field *google_protobuf.FieldDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:202
	_go_fuzz_dep_.CoverTab[134100]++
												name := GetCustomName(field)
												if len(name) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:204
		_go_fuzz_dep_.CoverTab[134102]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:205
		// _ = "end of CoverTab[134102]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:206
		_go_fuzz_dep_.CoverTab[134103]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:206
		// _ = "end of CoverTab[134103]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:206
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:206
	// _ = "end of CoverTab[134100]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:206
	_go_fuzz_dep_.CoverTab[134101]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:207
	// _ = "end of CoverTab[134101]"
}

func IsEnumCustomName(field *google_protobuf.EnumDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:210
	_go_fuzz_dep_.CoverTab[134104]++
												name := GetEnumCustomName(field)
												if len(name) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:212
		_go_fuzz_dep_.CoverTab[134106]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:213
		// _ = "end of CoverTab[134106]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:214
		_go_fuzz_dep_.CoverTab[134107]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:214
		// _ = "end of CoverTab[134107]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:214
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:214
	// _ = "end of CoverTab[134104]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:214
	_go_fuzz_dep_.CoverTab[134105]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:215
	// _ = "end of CoverTab[134105]"
}

func IsEnumValueCustomName(field *google_protobuf.EnumValueDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:218
	_go_fuzz_dep_.CoverTab[134108]++
												name := GetEnumValueCustomName(field)
												if len(name) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:220
		_go_fuzz_dep_.CoverTab[134110]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:221
		// _ = "end of CoverTab[134110]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:222
		_go_fuzz_dep_.CoverTab[134111]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:222
		// _ = "end of CoverTab[134111]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:222
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:222
	// _ = "end of CoverTab[134108]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:222
	_go_fuzz_dep_.CoverTab[134109]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:223
	// _ = "end of CoverTab[134109]"
}

func GetCustomName(field *google_protobuf.FieldDescriptorProto) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:226
	_go_fuzz_dep_.CoverTab[134112]++
												if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:227
		_go_fuzz_dep_.CoverTab[134115]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:228
		// _ = "end of CoverTab[134115]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:229
		_go_fuzz_dep_.CoverTab[134116]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:229
		// _ = "end of CoverTab[134116]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:229
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:229
	// _ = "end of CoverTab[134112]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:229
	_go_fuzz_dep_.CoverTab[134113]++
												if field.Options != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:230
		_go_fuzz_dep_.CoverTab[134117]++
													v, err := proto.GetExtension(field.Options, E_Customname)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:232
			_go_fuzz_dep_.CoverTab[134118]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:232
			return v.(*string) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:232
			// _ = "end of CoverTab[134118]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:232
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:232
			_go_fuzz_dep_.CoverTab[134119]++
														return *(v.(*string))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:233
			// _ = "end of CoverTab[134119]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:234
			_go_fuzz_dep_.CoverTab[134120]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:234
			// _ = "end of CoverTab[134120]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:234
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:234
		// _ = "end of CoverTab[134117]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:235
		_go_fuzz_dep_.CoverTab[134121]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:235
		// _ = "end of CoverTab[134121]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:235
	// _ = "end of CoverTab[134113]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:235
	_go_fuzz_dep_.CoverTab[134114]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:236
	// _ = "end of CoverTab[134114]"
}

func GetEnumCustomName(field *google_protobuf.EnumDescriptorProto) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:239
	_go_fuzz_dep_.CoverTab[134122]++
												if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:240
		_go_fuzz_dep_.CoverTab[134125]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:241
		// _ = "end of CoverTab[134125]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:242
		_go_fuzz_dep_.CoverTab[134126]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:242
		// _ = "end of CoverTab[134126]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:242
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:242
	// _ = "end of CoverTab[134122]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:242
	_go_fuzz_dep_.CoverTab[134123]++
												if field.Options != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:243
		_go_fuzz_dep_.CoverTab[134127]++
													v, err := proto.GetExtension(field.Options, E_EnumCustomname)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:245
			_go_fuzz_dep_.CoverTab[134128]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:245
			return v.(*string) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:245
			// _ = "end of CoverTab[134128]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:245
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:245
			_go_fuzz_dep_.CoverTab[134129]++
														return *(v.(*string))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:246
			// _ = "end of CoverTab[134129]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:247
			_go_fuzz_dep_.CoverTab[134130]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:247
			// _ = "end of CoverTab[134130]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:247
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:247
		// _ = "end of CoverTab[134127]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:248
		_go_fuzz_dep_.CoverTab[134131]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:248
		// _ = "end of CoverTab[134131]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:248
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:248
	// _ = "end of CoverTab[134123]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:248
	_go_fuzz_dep_.CoverTab[134124]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:249
	// _ = "end of CoverTab[134124]"
}

func GetEnumValueCustomName(field *google_protobuf.EnumValueDescriptorProto) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:252
	_go_fuzz_dep_.CoverTab[134132]++
												if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:253
		_go_fuzz_dep_.CoverTab[134135]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:254
		// _ = "end of CoverTab[134135]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:255
		_go_fuzz_dep_.CoverTab[134136]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:255
		// _ = "end of CoverTab[134136]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:255
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:255
	// _ = "end of CoverTab[134132]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:255
	_go_fuzz_dep_.CoverTab[134133]++
												if field.Options != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:256
		_go_fuzz_dep_.CoverTab[134137]++
													v, err := proto.GetExtension(field.Options, E_EnumvalueCustomname)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:258
			_go_fuzz_dep_.CoverTab[134138]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:258
			return v.(*string) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:258
			// _ = "end of CoverTab[134138]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:258
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:258
			_go_fuzz_dep_.CoverTab[134139]++
														return *(v.(*string))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:259
			// _ = "end of CoverTab[134139]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:260
			_go_fuzz_dep_.CoverTab[134140]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:260
			// _ = "end of CoverTab[134140]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:260
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:260
		// _ = "end of CoverTab[134137]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:261
		_go_fuzz_dep_.CoverTab[134141]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:261
		// _ = "end of CoverTab[134141]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:261
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:261
	// _ = "end of CoverTab[134133]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:261
	_go_fuzz_dep_.CoverTab[134134]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:262
	// _ = "end of CoverTab[134134]"
}

func GetJsonTag(field *google_protobuf.FieldDescriptorProto) *string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:265
	_go_fuzz_dep_.CoverTab[134142]++
												if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:266
		_go_fuzz_dep_.CoverTab[134145]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:267
		// _ = "end of CoverTab[134145]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:268
		_go_fuzz_dep_.CoverTab[134146]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:268
		// _ = "end of CoverTab[134146]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:268
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:268
	// _ = "end of CoverTab[134142]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:268
	_go_fuzz_dep_.CoverTab[134143]++
												if field.Options != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:269
		_go_fuzz_dep_.CoverTab[134147]++
													v, err := proto.GetExtension(field.Options, E_Jsontag)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:271
			_go_fuzz_dep_.CoverTab[134148]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:271
			return v.(*string) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:271
			// _ = "end of CoverTab[134148]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:271
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:271
			_go_fuzz_dep_.CoverTab[134149]++
														return (v.(*string))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:272
			// _ = "end of CoverTab[134149]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:273
			_go_fuzz_dep_.CoverTab[134150]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:273
			// _ = "end of CoverTab[134150]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:273
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:273
		// _ = "end of CoverTab[134147]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:274
		_go_fuzz_dep_.CoverTab[134151]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:274
		// _ = "end of CoverTab[134151]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:274
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:274
	// _ = "end of CoverTab[134143]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:274
	_go_fuzz_dep_.CoverTab[134144]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:275
	// _ = "end of CoverTab[134144]"
}

func GetMoreTags(field *google_protobuf.FieldDescriptorProto) *string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:278
	_go_fuzz_dep_.CoverTab[134152]++
												if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:279
		_go_fuzz_dep_.CoverTab[134155]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:280
		// _ = "end of CoverTab[134155]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:281
		_go_fuzz_dep_.CoverTab[134156]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:281
		// _ = "end of CoverTab[134156]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:281
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:281
	// _ = "end of CoverTab[134152]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:281
	_go_fuzz_dep_.CoverTab[134153]++
												if field.Options != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:282
		_go_fuzz_dep_.CoverTab[134157]++
													v, err := proto.GetExtension(field.Options, E_Moretags)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:284
			_go_fuzz_dep_.CoverTab[134158]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:284
			return v.(*string) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:284
			// _ = "end of CoverTab[134158]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:284
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:284
			_go_fuzz_dep_.CoverTab[134159]++
														return (v.(*string))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:285
			// _ = "end of CoverTab[134159]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:286
			_go_fuzz_dep_.CoverTab[134160]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:286
			// _ = "end of CoverTab[134160]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:286
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:286
		// _ = "end of CoverTab[134157]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:287
		_go_fuzz_dep_.CoverTab[134161]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:287
		// _ = "end of CoverTab[134161]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:287
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:287
	// _ = "end of CoverTab[134153]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:287
	_go_fuzz_dep_.CoverTab[134154]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:288
	// _ = "end of CoverTab[134154]"
}

type EnableFunc func(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool

func EnabledGoEnumPrefix(file *google_protobuf.FileDescriptorProto, enum *google_protobuf.EnumDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:293
	_go_fuzz_dep_.CoverTab[134162]++
												return proto.GetBoolExtension(enum.Options, E_GoprotoEnumPrefix, proto.GetBoolExtension(file.Options, E_GoprotoEnumPrefixAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:294
	// _ = "end of CoverTab[134162]"
}

func EnabledGoStringer(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:297
	_go_fuzz_dep_.CoverTab[134163]++
												return proto.GetBoolExtension(message.Options, E_GoprotoStringer, proto.GetBoolExtension(file.Options, E_GoprotoStringerAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:298
	// _ = "end of CoverTab[134163]"
}

func HasGoGetters(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:301
	_go_fuzz_dep_.CoverTab[134164]++
												return proto.GetBoolExtension(message.Options, E_GoprotoGetters, proto.GetBoolExtension(file.Options, E_GoprotoGettersAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:302
	// _ = "end of CoverTab[134164]"
}

func IsUnion(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:305
	_go_fuzz_dep_.CoverTab[134165]++
												return proto.GetBoolExtension(message.Options, E_Onlyone, proto.GetBoolExtension(file.Options, E_OnlyoneAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:306
	// _ = "end of CoverTab[134165]"
}

func HasGoString(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:309
	_go_fuzz_dep_.CoverTab[134166]++
												return proto.GetBoolExtension(message.Options, E_Gostring, proto.GetBoolExtension(file.Options, E_GostringAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:310
	// _ = "end of CoverTab[134166]"
}

func HasEqual(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:313
	_go_fuzz_dep_.CoverTab[134167]++
												return proto.GetBoolExtension(message.Options, E_Equal, proto.GetBoolExtension(file.Options, E_EqualAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:314
	// _ = "end of CoverTab[134167]"
}

func HasVerboseEqual(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:317
	_go_fuzz_dep_.CoverTab[134168]++
												return proto.GetBoolExtension(message.Options, E_VerboseEqual, proto.GetBoolExtension(file.Options, E_VerboseEqualAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:318
	// _ = "end of CoverTab[134168]"
}

func IsStringer(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:321
	_go_fuzz_dep_.CoverTab[134169]++
												return proto.GetBoolExtension(message.Options, E_Stringer, proto.GetBoolExtension(file.Options, E_StringerAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:322
	// _ = "end of CoverTab[134169]"
}

func IsFace(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:325
	_go_fuzz_dep_.CoverTab[134170]++
												return proto.GetBoolExtension(message.Options, E_Face, proto.GetBoolExtension(file.Options, E_FaceAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:326
	// _ = "end of CoverTab[134170]"
}

func HasDescription(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:329
	_go_fuzz_dep_.CoverTab[134171]++
												return proto.GetBoolExtension(message.Options, E_Description, proto.GetBoolExtension(file.Options, E_DescriptionAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:330
	// _ = "end of CoverTab[134171]"
}

func HasPopulate(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:333
	_go_fuzz_dep_.CoverTab[134172]++
												return proto.GetBoolExtension(message.Options, E_Populate, proto.GetBoolExtension(file.Options, E_PopulateAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:334
	// _ = "end of CoverTab[134172]"
}

func HasTestGen(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:337
	_go_fuzz_dep_.CoverTab[134173]++
												return proto.GetBoolExtension(message.Options, E_Testgen, proto.GetBoolExtension(file.Options, E_TestgenAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:338
	// _ = "end of CoverTab[134173]"
}

func HasBenchGen(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:341
	_go_fuzz_dep_.CoverTab[134174]++
												return proto.GetBoolExtension(message.Options, E_Benchgen, proto.GetBoolExtension(file.Options, E_BenchgenAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:342
	// _ = "end of CoverTab[134174]"
}

func IsMarshaler(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:345
	_go_fuzz_dep_.CoverTab[134175]++
												return proto.GetBoolExtension(message.Options, E_Marshaler, proto.GetBoolExtension(file.Options, E_MarshalerAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:346
	// _ = "end of CoverTab[134175]"
}

func IsUnmarshaler(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:349
	_go_fuzz_dep_.CoverTab[134176]++
												return proto.GetBoolExtension(message.Options, E_Unmarshaler, proto.GetBoolExtension(file.Options, E_UnmarshalerAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:350
	// _ = "end of CoverTab[134176]"
}

func IsStableMarshaler(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:353
	_go_fuzz_dep_.CoverTab[134177]++
												return proto.GetBoolExtension(message.Options, E_StableMarshaler, proto.GetBoolExtension(file.Options, E_StableMarshalerAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:354
	// _ = "end of CoverTab[134177]"
}

func IsSizer(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:357
	_go_fuzz_dep_.CoverTab[134178]++
												return proto.GetBoolExtension(message.Options, E_Sizer, proto.GetBoolExtension(file.Options, E_SizerAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:358
	// _ = "end of CoverTab[134178]"
}

func IsProtoSizer(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:361
	_go_fuzz_dep_.CoverTab[134179]++
												return proto.GetBoolExtension(message.Options, E_Protosizer, proto.GetBoolExtension(file.Options, E_ProtosizerAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:362
	// _ = "end of CoverTab[134179]"
}

func IsGoEnumStringer(file *google_protobuf.FileDescriptorProto, enum *google_protobuf.EnumDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:365
	_go_fuzz_dep_.CoverTab[134180]++
												return proto.GetBoolExtension(enum.Options, E_GoprotoEnumStringer, proto.GetBoolExtension(file.Options, E_GoprotoEnumStringerAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:366
	// _ = "end of CoverTab[134180]"
}

func IsEnumStringer(file *google_protobuf.FileDescriptorProto, enum *google_protobuf.EnumDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:369
	_go_fuzz_dep_.CoverTab[134181]++
												return proto.GetBoolExtension(enum.Options, E_EnumStringer, proto.GetBoolExtension(file.Options, E_EnumStringerAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:370
	// _ = "end of CoverTab[134181]"
}

func IsUnsafeMarshaler(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:373
	_go_fuzz_dep_.CoverTab[134182]++
												return proto.GetBoolExtension(message.Options, E_UnsafeMarshaler, proto.GetBoolExtension(file.Options, E_UnsafeMarshalerAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:374
	// _ = "end of CoverTab[134182]"
}

func IsUnsafeUnmarshaler(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:377
	_go_fuzz_dep_.CoverTab[134183]++
												return proto.GetBoolExtension(message.Options, E_UnsafeUnmarshaler, proto.GetBoolExtension(file.Options, E_UnsafeUnmarshalerAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:378
	// _ = "end of CoverTab[134183]"
}

func HasExtensionsMap(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:381
	_go_fuzz_dep_.CoverTab[134184]++
												return proto.GetBoolExtension(message.Options, E_GoprotoExtensionsMap, proto.GetBoolExtension(file.Options, E_GoprotoExtensionsMapAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:382
	// _ = "end of CoverTab[134184]"
}

func HasUnrecognized(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:385
	_go_fuzz_dep_.CoverTab[134185]++
												return proto.GetBoolExtension(message.Options, E_GoprotoUnrecognized, proto.GetBoolExtension(file.Options, E_GoprotoUnrecognizedAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:386
	// _ = "end of CoverTab[134185]"
}

func IsProto3(file *google_protobuf.FileDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:389
	_go_fuzz_dep_.CoverTab[134186]++
												return file.GetSyntax() == "proto3"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:390
	// _ = "end of CoverTab[134186]"
}

func ImportsGoGoProto(file *google_protobuf.FileDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:393
	_go_fuzz_dep_.CoverTab[134187]++
												return proto.GetBoolExtension(file.Options, E_GogoprotoImport, true)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:394
	// _ = "end of CoverTab[134187]"
}

func HasCompare(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:397
	_go_fuzz_dep_.CoverTab[134188]++
												return proto.GetBoolExtension(message.Options, E_Compare, proto.GetBoolExtension(file.Options, E_CompareAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:398
	// _ = "end of CoverTab[134188]"
}

func RegistersGolangProto(file *google_protobuf.FileDescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:401
	_go_fuzz_dep_.CoverTab[134189]++
												return proto.GetBoolExtension(file.Options, E_GoprotoRegistration, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:402
	// _ = "end of CoverTab[134189]"
}

func HasMessageName(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:405
	_go_fuzz_dep_.CoverTab[134190]++
												return proto.GetBoolExtension(message.Options, E_Messagename, proto.GetBoolExtension(file.Options, E_MessagenameAll, false))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:406
	// _ = "end of CoverTab[134190]"
}

func HasSizecache(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:409
	_go_fuzz_dep_.CoverTab[134191]++
												return proto.GetBoolExtension(message.Options, E_GoprotoSizecache, proto.GetBoolExtension(file.Options, E_GoprotoSizecacheAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:410
	// _ = "end of CoverTab[134191]"
}

func HasUnkeyed(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:413
	_go_fuzz_dep_.CoverTab[134192]++
												return proto.GetBoolExtension(message.Options, E_GoprotoUnkeyed, proto.GetBoolExtension(file.Options, E_GoprotoUnkeyedAll, true))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:414
	// _ = "end of CoverTab[134192]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:415
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto/helper.go:415
var _ = _go_fuzz_dep_.CoverTab
