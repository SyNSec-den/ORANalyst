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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:29
package descriptor

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:29
)

import (
	"strings"
)

func (msg *DescriptorProto) GetMapFields() (*FieldDescriptorProto, *FieldDescriptorProto) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:35
	_go_fuzz_dep_.CoverTab[133796]++
														if !msg.GetOptions().GetMapEntry() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:36
		_go_fuzz_dep_.CoverTab[133798]++
															return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:37
		// _ = "end of CoverTab[133798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:38
		_go_fuzz_dep_.CoverTab[133799]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:38
		// _ = "end of CoverTab[133799]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:38
	// _ = "end of CoverTab[133796]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:38
	_go_fuzz_dep_.CoverTab[133797]++
														return msg.GetField()[0], msg.GetField()[1]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:39
	// _ = "end of CoverTab[133797]"
}

func dotToUnderscore(r rune) rune {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:42
	_go_fuzz_dep_.CoverTab[133800]++
														if r == '.' {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:43
		_go_fuzz_dep_.CoverTab[133802]++
															return '_'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:44
		// _ = "end of CoverTab[133802]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:45
		_go_fuzz_dep_.CoverTab[133803]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:45
		// _ = "end of CoverTab[133803]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:45
	// _ = "end of CoverTab[133800]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:45
	_go_fuzz_dep_.CoverTab[133801]++
														return r
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:46
	// _ = "end of CoverTab[133801]"
}

func (field *FieldDescriptorProto) WireType() (wire int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:49
	_go_fuzz_dep_.CoverTab[133804]++
														switch *field.Type {
	case FieldDescriptorProto_TYPE_DOUBLE:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:51
		_go_fuzz_dep_.CoverTab[133806]++
															return 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:52
		// _ = "end of CoverTab[133806]"
	case FieldDescriptorProto_TYPE_FLOAT:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:53
		_go_fuzz_dep_.CoverTab[133807]++
															return 5
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:54
		// _ = "end of CoverTab[133807]"
	case FieldDescriptorProto_TYPE_INT64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:55
		_go_fuzz_dep_.CoverTab[133808]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:56
		// _ = "end of CoverTab[133808]"
	case FieldDescriptorProto_TYPE_UINT64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:57
		_go_fuzz_dep_.CoverTab[133809]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:58
		// _ = "end of CoverTab[133809]"
	case FieldDescriptorProto_TYPE_INT32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:59
		_go_fuzz_dep_.CoverTab[133810]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:60
		// _ = "end of CoverTab[133810]"
	case FieldDescriptorProto_TYPE_UINT32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:61
		_go_fuzz_dep_.CoverTab[133811]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:62
		// _ = "end of CoverTab[133811]"
	case FieldDescriptorProto_TYPE_FIXED64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:63
		_go_fuzz_dep_.CoverTab[133812]++
															return 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:64
		// _ = "end of CoverTab[133812]"
	case FieldDescriptorProto_TYPE_FIXED32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:65
		_go_fuzz_dep_.CoverTab[133813]++
															return 5
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:66
		// _ = "end of CoverTab[133813]"
	case FieldDescriptorProto_TYPE_BOOL:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:67
		_go_fuzz_dep_.CoverTab[133814]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:68
		// _ = "end of CoverTab[133814]"
	case FieldDescriptorProto_TYPE_STRING:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:69
		_go_fuzz_dep_.CoverTab[133815]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:70
		// _ = "end of CoverTab[133815]"
	case FieldDescriptorProto_TYPE_GROUP:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:71
		_go_fuzz_dep_.CoverTab[133816]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:72
		// _ = "end of CoverTab[133816]"
	case FieldDescriptorProto_TYPE_MESSAGE:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:73
		_go_fuzz_dep_.CoverTab[133817]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:74
		// _ = "end of CoverTab[133817]"
	case FieldDescriptorProto_TYPE_BYTES:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:75
		_go_fuzz_dep_.CoverTab[133818]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:76
		// _ = "end of CoverTab[133818]"
	case FieldDescriptorProto_TYPE_ENUM:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:77
		_go_fuzz_dep_.CoverTab[133819]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:78
		// _ = "end of CoverTab[133819]"
	case FieldDescriptorProto_TYPE_SFIXED32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:79
		_go_fuzz_dep_.CoverTab[133820]++
															return 5
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:80
		// _ = "end of CoverTab[133820]"
	case FieldDescriptorProto_TYPE_SFIXED64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:81
		_go_fuzz_dep_.CoverTab[133821]++
															return 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:82
		// _ = "end of CoverTab[133821]"
	case FieldDescriptorProto_TYPE_SINT32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:83
		_go_fuzz_dep_.CoverTab[133822]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:84
		// _ = "end of CoverTab[133822]"
	case FieldDescriptorProto_TYPE_SINT64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:85
		_go_fuzz_dep_.CoverTab[133823]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:86
		// _ = "end of CoverTab[133823]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:86
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:86
		_go_fuzz_dep_.CoverTab[133824]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:86
		// _ = "end of CoverTab[133824]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:87
	// _ = "end of CoverTab[133804]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:87
	_go_fuzz_dep_.CoverTab[133805]++
														panic("unreachable")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:88
	// _ = "end of CoverTab[133805]"
}

func (field *FieldDescriptorProto) GetKeyUint64() (x uint64) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:91
	_go_fuzz_dep_.CoverTab[133825]++
														packed := field.IsPacked()
														wireType := field.WireType()
														fieldNumber := field.GetNumber()
														if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:95
		_go_fuzz_dep_.CoverTab[133827]++
															wireType = 2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:96
		// _ = "end of CoverTab[133827]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:97
		_go_fuzz_dep_.CoverTab[133828]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:97
		// _ = "end of CoverTab[133828]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:97
	// _ = "end of CoverTab[133825]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:97
	_go_fuzz_dep_.CoverTab[133826]++
														x = uint64(uint32(fieldNumber)<<3 | uint32(wireType))
														return x
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:99
	// _ = "end of CoverTab[133826]"
}

func (field *FieldDescriptorProto) GetKey3Uint64() (x uint64) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:102
	_go_fuzz_dep_.CoverTab[133829]++
														packed := field.IsPacked3()
														wireType := field.WireType()
														fieldNumber := field.GetNumber()
														if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:106
		_go_fuzz_dep_.CoverTab[133831]++
															wireType = 2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:107
		// _ = "end of CoverTab[133831]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:108
		_go_fuzz_dep_.CoverTab[133832]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:108
		// _ = "end of CoverTab[133832]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:108
	// _ = "end of CoverTab[133829]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:108
	_go_fuzz_dep_.CoverTab[133830]++
														x = uint64(uint32(fieldNumber)<<3 | uint32(wireType))
														return x
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:110
	// _ = "end of CoverTab[133830]"
}

func (field *FieldDescriptorProto) GetKey() []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:113
	_go_fuzz_dep_.CoverTab[133833]++
														x := field.GetKeyUint64()
														i := 0
														keybuf := make([]byte, 0)
														for i = 0; x > 127; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:117
		_go_fuzz_dep_.CoverTab[133835]++
															keybuf = append(keybuf, 0x80|uint8(x&0x7F))
															x >>= 7
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:119
		// _ = "end of CoverTab[133835]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:120
	// _ = "end of CoverTab[133833]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:120
	_go_fuzz_dep_.CoverTab[133834]++
														keybuf = append(keybuf, uint8(x))
														return keybuf
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:122
	// _ = "end of CoverTab[133834]"
}

func (field *FieldDescriptorProto) GetKey3() []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:125
	_go_fuzz_dep_.CoverTab[133836]++
														x := field.GetKey3Uint64()
														i := 0
														keybuf := make([]byte, 0)
														for i = 0; x > 127; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:129
		_go_fuzz_dep_.CoverTab[133838]++
															keybuf = append(keybuf, 0x80|uint8(x&0x7F))
															x >>= 7
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:131
		// _ = "end of CoverTab[133838]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:132
	// _ = "end of CoverTab[133836]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:132
	_go_fuzz_dep_.CoverTab[133837]++
														keybuf = append(keybuf, uint8(x))
														return keybuf
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:134
	// _ = "end of CoverTab[133837]"
}

func (desc *FileDescriptorSet) GetField(packageName, messageName, fieldName string) *FieldDescriptorProto {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:137
	_go_fuzz_dep_.CoverTab[133839]++
														msg := desc.GetMessage(packageName, messageName)
														if msg == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:139
		_go_fuzz_dep_.CoverTab[133842]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:140
		// _ = "end of CoverTab[133842]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:141
		_go_fuzz_dep_.CoverTab[133843]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:141
		// _ = "end of CoverTab[133843]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:141
	// _ = "end of CoverTab[133839]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:141
	_go_fuzz_dep_.CoverTab[133840]++
														for _, field := range msg.GetField() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:142
		_go_fuzz_dep_.CoverTab[133844]++
															if field.GetName() == fieldName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:143
			_go_fuzz_dep_.CoverTab[133845]++
																return field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:144
			// _ = "end of CoverTab[133845]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:145
			_go_fuzz_dep_.CoverTab[133846]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:145
			// _ = "end of CoverTab[133846]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:145
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:145
		// _ = "end of CoverTab[133844]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:146
	// _ = "end of CoverTab[133840]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:146
	_go_fuzz_dep_.CoverTab[133841]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:147
	// _ = "end of CoverTab[133841]"
}

func (file *FileDescriptorProto) GetMessage(typeName string) *DescriptorProto {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:150
	_go_fuzz_dep_.CoverTab[133847]++
														for _, msg := range file.GetMessageType() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:151
		_go_fuzz_dep_.CoverTab[133849]++
															if msg.GetName() == typeName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:152
			_go_fuzz_dep_.CoverTab[133851]++
																return msg
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:153
			// _ = "end of CoverTab[133851]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:154
			_go_fuzz_dep_.CoverTab[133852]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:154
			// _ = "end of CoverTab[133852]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:154
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:154
		// _ = "end of CoverTab[133849]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:154
		_go_fuzz_dep_.CoverTab[133850]++
															nes := file.GetNestedMessage(msg, strings.TrimPrefix(typeName, msg.GetName()+"."))
															if nes != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:156
			_go_fuzz_dep_.CoverTab[133853]++
																return nes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:157
			// _ = "end of CoverTab[133853]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:158
			_go_fuzz_dep_.CoverTab[133854]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:158
			// _ = "end of CoverTab[133854]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:158
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:158
		// _ = "end of CoverTab[133850]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:159
	// _ = "end of CoverTab[133847]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:159
	_go_fuzz_dep_.CoverTab[133848]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:160
	// _ = "end of CoverTab[133848]"
}

func (file *FileDescriptorProto) GetNestedMessage(msg *DescriptorProto, typeName string) *DescriptorProto {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:163
	_go_fuzz_dep_.CoverTab[133855]++
														for _, nes := range msg.GetNestedType() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:164
		_go_fuzz_dep_.CoverTab[133857]++
															if nes.GetName() == typeName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:165
			_go_fuzz_dep_.CoverTab[133859]++
																return nes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:166
			// _ = "end of CoverTab[133859]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:167
			_go_fuzz_dep_.CoverTab[133860]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:167
			// _ = "end of CoverTab[133860]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:167
		// _ = "end of CoverTab[133857]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:167
		_go_fuzz_dep_.CoverTab[133858]++
															res := file.GetNestedMessage(nes, strings.TrimPrefix(typeName, nes.GetName()+"."))
															if res != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:169
			_go_fuzz_dep_.CoverTab[133861]++
																return res
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:170
			// _ = "end of CoverTab[133861]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:171
			_go_fuzz_dep_.CoverTab[133862]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:171
			// _ = "end of CoverTab[133862]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:171
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:171
		// _ = "end of CoverTab[133858]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:172
	// _ = "end of CoverTab[133855]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:172
	_go_fuzz_dep_.CoverTab[133856]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:173
	// _ = "end of CoverTab[133856]"
}

func (desc *FileDescriptorSet) GetMessage(packageName string, typeName string) *DescriptorProto {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:176
	_go_fuzz_dep_.CoverTab[133863]++
														for _, file := range desc.GetFile() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:177
		_go_fuzz_dep_.CoverTab[133865]++
															if strings.Map(dotToUnderscore, file.GetPackage()) != strings.Map(dotToUnderscore, packageName) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:178
			_go_fuzz_dep_.CoverTab[133868]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:179
			// _ = "end of CoverTab[133868]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:180
			_go_fuzz_dep_.CoverTab[133869]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:180
			// _ = "end of CoverTab[133869]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:180
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:180
		// _ = "end of CoverTab[133865]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:180
		_go_fuzz_dep_.CoverTab[133866]++
															for _, msg := range file.GetMessageType() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:181
			_go_fuzz_dep_.CoverTab[133870]++
																if msg.GetName() == typeName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:182
				_go_fuzz_dep_.CoverTab[133871]++
																	return msg
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:183
				// _ = "end of CoverTab[133871]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:184
				_go_fuzz_dep_.CoverTab[133872]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:184
				// _ = "end of CoverTab[133872]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:184
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:184
			// _ = "end of CoverTab[133870]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:185
		// _ = "end of CoverTab[133866]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:185
		_go_fuzz_dep_.CoverTab[133867]++
															for _, msg := range file.GetMessageType() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:186
			_go_fuzz_dep_.CoverTab[133873]++
																for _, nes := range msg.GetNestedType() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:187
				_go_fuzz_dep_.CoverTab[133874]++
																	if nes.GetName() == typeName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:188
					_go_fuzz_dep_.CoverTab[133876]++
																		return nes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:189
					// _ = "end of CoverTab[133876]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:190
					_go_fuzz_dep_.CoverTab[133877]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:190
					// _ = "end of CoverTab[133877]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:190
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:190
				// _ = "end of CoverTab[133874]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:190
				_go_fuzz_dep_.CoverTab[133875]++
																	if msg.GetName()+"."+nes.GetName() == typeName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:191
					_go_fuzz_dep_.CoverTab[133878]++
																		return nes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:192
					// _ = "end of CoverTab[133878]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:193
					_go_fuzz_dep_.CoverTab[133879]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:193
					// _ = "end of CoverTab[133879]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:193
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:193
				// _ = "end of CoverTab[133875]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:194
			// _ = "end of CoverTab[133873]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:195
		// _ = "end of CoverTab[133867]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:196
	// _ = "end of CoverTab[133863]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:196
	_go_fuzz_dep_.CoverTab[133864]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:197
	// _ = "end of CoverTab[133864]"
}

func (desc *FileDescriptorSet) IsProto3(packageName string, typeName string) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:200
	_go_fuzz_dep_.CoverTab[133880]++
														for _, file := range desc.GetFile() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:201
		_go_fuzz_dep_.CoverTab[133882]++
															if strings.Map(dotToUnderscore, file.GetPackage()) != strings.Map(dotToUnderscore, packageName) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:202
			_go_fuzz_dep_.CoverTab[133885]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:203
			// _ = "end of CoverTab[133885]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:204
			_go_fuzz_dep_.CoverTab[133886]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:204
			// _ = "end of CoverTab[133886]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:204
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:204
		// _ = "end of CoverTab[133882]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:204
		_go_fuzz_dep_.CoverTab[133883]++
															for _, msg := range file.GetMessageType() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:205
			_go_fuzz_dep_.CoverTab[133887]++
																if msg.GetName() == typeName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:206
				_go_fuzz_dep_.CoverTab[133888]++
																	return file.GetSyntax() == "proto3"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:207
				// _ = "end of CoverTab[133888]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:208
				_go_fuzz_dep_.CoverTab[133889]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:208
				// _ = "end of CoverTab[133889]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:208
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:208
			// _ = "end of CoverTab[133887]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:209
		// _ = "end of CoverTab[133883]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:209
		_go_fuzz_dep_.CoverTab[133884]++
															for _, msg := range file.GetMessageType() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:210
			_go_fuzz_dep_.CoverTab[133890]++
																for _, nes := range msg.GetNestedType() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:211
				_go_fuzz_dep_.CoverTab[133891]++
																	if nes.GetName() == typeName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:212
					_go_fuzz_dep_.CoverTab[133893]++
																		return file.GetSyntax() == "proto3"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:213
					// _ = "end of CoverTab[133893]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:214
					_go_fuzz_dep_.CoverTab[133894]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:214
					// _ = "end of CoverTab[133894]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:214
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:214
				// _ = "end of CoverTab[133891]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:214
				_go_fuzz_dep_.CoverTab[133892]++
																	if msg.GetName()+"."+nes.GetName() == typeName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:215
					_go_fuzz_dep_.CoverTab[133895]++
																		return file.GetSyntax() == "proto3"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:216
					// _ = "end of CoverTab[133895]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:217
					_go_fuzz_dep_.CoverTab[133896]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:217
					// _ = "end of CoverTab[133896]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:217
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:217
				// _ = "end of CoverTab[133892]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:218
			// _ = "end of CoverTab[133890]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:219
		// _ = "end of CoverTab[133884]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:220
	// _ = "end of CoverTab[133880]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:220
	_go_fuzz_dep_.CoverTab[133881]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:221
	// _ = "end of CoverTab[133881]"
}

func (msg *DescriptorProto) IsExtendable() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:224
	_go_fuzz_dep_.CoverTab[133897]++
														return len(msg.GetExtensionRange()) > 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:225
	// _ = "end of CoverTab[133897]"
}

func (desc *FileDescriptorSet) FindExtension(packageName string, typeName string, fieldName string) (extPackageName string, field *FieldDescriptorProto) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:228
	_go_fuzz_dep_.CoverTab[133898]++
														parent := desc.GetMessage(packageName, typeName)
														if parent == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:230
		_go_fuzz_dep_.CoverTab[133902]++
															return "", nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:231
		// _ = "end of CoverTab[133902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:232
		_go_fuzz_dep_.CoverTab[133903]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:232
		// _ = "end of CoverTab[133903]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:232
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:232
	// _ = "end of CoverTab[133898]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:232
	_go_fuzz_dep_.CoverTab[133899]++
														if !parent.IsExtendable() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:233
		_go_fuzz_dep_.CoverTab[133904]++
															return "", nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:234
		// _ = "end of CoverTab[133904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:235
		_go_fuzz_dep_.CoverTab[133905]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:235
		// _ = "end of CoverTab[133905]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:235
	// _ = "end of CoverTab[133899]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:235
	_go_fuzz_dep_.CoverTab[133900]++
														extendee := "." + packageName + "." + typeName
														for _, file := range desc.GetFile() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:237
		_go_fuzz_dep_.CoverTab[133906]++
															for _, ext := range file.GetExtension() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:238
			_go_fuzz_dep_.CoverTab[133907]++
																if strings.Map(dotToUnderscore, file.GetPackage()) == strings.Map(dotToUnderscore, packageName) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:239
				_go_fuzz_dep_.CoverTab[133909]++
																	if !(ext.GetExtendee() == typeName || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:240
					_go_fuzz_dep_.CoverTab[133910]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:240
					return ext.GetExtendee() == extendee
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:240
					// _ = "end of CoverTab[133910]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:240
				}()) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:240
					_go_fuzz_dep_.CoverTab[133911]++
																		continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:241
					// _ = "end of CoverTab[133911]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:242
					_go_fuzz_dep_.CoverTab[133912]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:242
					// _ = "end of CoverTab[133912]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:242
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:242
				// _ = "end of CoverTab[133909]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:243
				_go_fuzz_dep_.CoverTab[133913]++
																	if ext.GetExtendee() != extendee {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:244
					_go_fuzz_dep_.CoverTab[133914]++
																		continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:245
					// _ = "end of CoverTab[133914]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:246
					_go_fuzz_dep_.CoverTab[133915]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:246
					// _ = "end of CoverTab[133915]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:246
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:246
				// _ = "end of CoverTab[133913]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:247
			// _ = "end of CoverTab[133907]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:247
			_go_fuzz_dep_.CoverTab[133908]++
																if ext.GetName() == fieldName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:248
				_go_fuzz_dep_.CoverTab[133916]++
																	return file.GetPackage(), ext
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:249
				// _ = "end of CoverTab[133916]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:250
				_go_fuzz_dep_.CoverTab[133917]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:250
				// _ = "end of CoverTab[133917]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:250
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:250
			// _ = "end of CoverTab[133908]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:251
		// _ = "end of CoverTab[133906]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:252
	// _ = "end of CoverTab[133900]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:252
	_go_fuzz_dep_.CoverTab[133901]++
														return "", nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:253
	// _ = "end of CoverTab[133901]"
}

func (desc *FileDescriptorSet) FindExtensionByFieldNumber(packageName string, typeName string, fieldNum int32) (extPackageName string, field *FieldDescriptorProto) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:256
	_go_fuzz_dep_.CoverTab[133918]++
														parent := desc.GetMessage(packageName, typeName)
														if parent == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:258
		_go_fuzz_dep_.CoverTab[133922]++
															return "", nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:259
		// _ = "end of CoverTab[133922]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:260
		_go_fuzz_dep_.CoverTab[133923]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:260
		// _ = "end of CoverTab[133923]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:260
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:260
	// _ = "end of CoverTab[133918]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:260
	_go_fuzz_dep_.CoverTab[133919]++
														if !parent.IsExtendable() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:261
		_go_fuzz_dep_.CoverTab[133924]++
															return "", nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:262
		// _ = "end of CoverTab[133924]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:263
		_go_fuzz_dep_.CoverTab[133925]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:263
		// _ = "end of CoverTab[133925]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:263
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:263
	// _ = "end of CoverTab[133919]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:263
	_go_fuzz_dep_.CoverTab[133920]++
														extendee := "." + packageName + "." + typeName
														for _, file := range desc.GetFile() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:265
		_go_fuzz_dep_.CoverTab[133926]++
															for _, ext := range file.GetExtension() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:266
			_go_fuzz_dep_.CoverTab[133927]++
																if strings.Map(dotToUnderscore, file.GetPackage()) == strings.Map(dotToUnderscore, packageName) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:267
				_go_fuzz_dep_.CoverTab[133929]++
																	if !(ext.GetExtendee() == typeName || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:268
					_go_fuzz_dep_.CoverTab[133930]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:268
					return ext.GetExtendee() == extendee
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:268
					// _ = "end of CoverTab[133930]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:268
				}()) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:268
					_go_fuzz_dep_.CoverTab[133931]++
																		continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:269
					// _ = "end of CoverTab[133931]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:270
					_go_fuzz_dep_.CoverTab[133932]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:270
					// _ = "end of CoverTab[133932]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:270
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:270
				// _ = "end of CoverTab[133929]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:271
				_go_fuzz_dep_.CoverTab[133933]++
																	if ext.GetExtendee() != extendee {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:272
					_go_fuzz_dep_.CoverTab[133934]++
																		continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:273
					// _ = "end of CoverTab[133934]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:274
					_go_fuzz_dep_.CoverTab[133935]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:274
					// _ = "end of CoverTab[133935]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:274
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:274
				// _ = "end of CoverTab[133933]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:275
			// _ = "end of CoverTab[133927]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:275
			_go_fuzz_dep_.CoverTab[133928]++
																if ext.GetNumber() == fieldNum {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:276
				_go_fuzz_dep_.CoverTab[133936]++
																	return file.GetPackage(), ext
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:277
				// _ = "end of CoverTab[133936]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:278
				_go_fuzz_dep_.CoverTab[133937]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:278
				// _ = "end of CoverTab[133937]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:278
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:278
			// _ = "end of CoverTab[133928]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:279
		// _ = "end of CoverTab[133926]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:280
	// _ = "end of CoverTab[133920]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:280
	_go_fuzz_dep_.CoverTab[133921]++
														return "", nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:281
	// _ = "end of CoverTab[133921]"
}

func (desc *FileDescriptorSet) FindMessage(packageName string, typeName string, fieldName string) (msgPackageName string, msgName string) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:284
	_go_fuzz_dep_.CoverTab[133938]++
														parent := desc.GetMessage(packageName, typeName)
														if parent == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:286
		_go_fuzz_dep_.CoverTab[133943]++
															return "", ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:287
		// _ = "end of CoverTab[133943]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:288
		_go_fuzz_dep_.CoverTab[133944]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:288
		// _ = "end of CoverTab[133944]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:288
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:288
	// _ = "end of CoverTab[133938]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:288
	_go_fuzz_dep_.CoverTab[133939]++
														field := parent.GetFieldDescriptor(fieldName)
														if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:290
		_go_fuzz_dep_.CoverTab[133945]++
															var extPackageName string
															extPackageName, field = desc.FindExtension(packageName, typeName, fieldName)
															if field == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:293
			_go_fuzz_dep_.CoverTab[133947]++
																return "", ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:294
			// _ = "end of CoverTab[133947]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:295
			_go_fuzz_dep_.CoverTab[133948]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:295
			// _ = "end of CoverTab[133948]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:295
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:295
		// _ = "end of CoverTab[133945]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:295
		_go_fuzz_dep_.CoverTab[133946]++
															packageName = extPackageName
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:296
		// _ = "end of CoverTab[133946]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:297
		_go_fuzz_dep_.CoverTab[133949]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:297
		// _ = "end of CoverTab[133949]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:297
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:297
	// _ = "end of CoverTab[133939]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:297
	_go_fuzz_dep_.CoverTab[133940]++
														typeNames := strings.Split(field.GetTypeName(), ".")
														if len(typeNames) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:299
		_go_fuzz_dep_.CoverTab[133950]++
															msg := desc.GetMessage(packageName, typeName)
															if msg == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:301
			_go_fuzz_dep_.CoverTab[133952]++
																return "", ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:302
			// _ = "end of CoverTab[133952]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:303
			_go_fuzz_dep_.CoverTab[133953]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:303
			// _ = "end of CoverTab[133953]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:303
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:303
		// _ = "end of CoverTab[133950]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:303
		_go_fuzz_dep_.CoverTab[133951]++
															return packageName, msg.GetName()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:304
		// _ = "end of CoverTab[133951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:305
		_go_fuzz_dep_.CoverTab[133954]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:305
		// _ = "end of CoverTab[133954]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:305
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:305
	// _ = "end of CoverTab[133940]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:305
	_go_fuzz_dep_.CoverTab[133941]++
														if len(typeNames) > 2 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:306
		_go_fuzz_dep_.CoverTab[133955]++
															for i := 1; i < len(typeNames)-1; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:307
			_go_fuzz_dep_.CoverTab[133956]++
																packageName = strings.Join(typeNames[1:len(typeNames)-i], ".")
																typeName = strings.Join(typeNames[len(typeNames)-i:], ".")
																msg := desc.GetMessage(packageName, typeName)
																if msg != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:311
				_go_fuzz_dep_.CoverTab[133957]++
																	typeNames := strings.Split(msg.GetName(), ".")
																	if len(typeNames) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:313
					_go_fuzz_dep_.CoverTab[133959]++
																		return packageName, msg.GetName()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:314
					// _ = "end of CoverTab[133959]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:315
					_go_fuzz_dep_.CoverTab[133960]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:315
					// _ = "end of CoverTab[133960]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:315
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:315
				// _ = "end of CoverTab[133957]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:315
				_go_fuzz_dep_.CoverTab[133958]++
																	return strings.Join(typeNames[1:len(typeNames)-1], "."), typeNames[len(typeNames)-1]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:316
				// _ = "end of CoverTab[133958]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:317
				_go_fuzz_dep_.CoverTab[133961]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:317
				// _ = "end of CoverTab[133961]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:317
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:317
			// _ = "end of CoverTab[133956]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:318
		// _ = "end of CoverTab[133955]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:319
		_go_fuzz_dep_.CoverTab[133962]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:319
		// _ = "end of CoverTab[133962]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:319
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:319
	// _ = "end of CoverTab[133941]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:319
	_go_fuzz_dep_.CoverTab[133942]++
														return "", ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:320
	// _ = "end of CoverTab[133942]"
}

func (msg *DescriptorProto) GetFieldDescriptor(fieldName string) *FieldDescriptorProto {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:323
	_go_fuzz_dep_.CoverTab[133963]++
														for _, field := range msg.GetField() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:324
		_go_fuzz_dep_.CoverTab[133965]++
															if field.GetName() == fieldName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:325
			_go_fuzz_dep_.CoverTab[133966]++
																return field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:326
			// _ = "end of CoverTab[133966]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:327
			_go_fuzz_dep_.CoverTab[133967]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:327
			// _ = "end of CoverTab[133967]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:327
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:327
		// _ = "end of CoverTab[133965]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:328
	// _ = "end of CoverTab[133963]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:328
	_go_fuzz_dep_.CoverTab[133964]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:329
	// _ = "end of CoverTab[133964]"
}

func (desc *FileDescriptorSet) GetEnum(packageName string, typeName string) *EnumDescriptorProto {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:332
	_go_fuzz_dep_.CoverTab[133968]++
														for _, file := range desc.GetFile() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:333
		_go_fuzz_dep_.CoverTab[133970]++
															if strings.Map(dotToUnderscore, file.GetPackage()) != strings.Map(dotToUnderscore, packageName) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:334
			_go_fuzz_dep_.CoverTab[133972]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:335
			// _ = "end of CoverTab[133972]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:336
			_go_fuzz_dep_.CoverTab[133973]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:336
			// _ = "end of CoverTab[133973]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:336
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:336
		// _ = "end of CoverTab[133970]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:336
		_go_fuzz_dep_.CoverTab[133971]++
															for _, enum := range file.GetEnumType() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:337
			_go_fuzz_dep_.CoverTab[133974]++
																if enum.GetName() == typeName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:338
				_go_fuzz_dep_.CoverTab[133975]++
																	return enum
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:339
				// _ = "end of CoverTab[133975]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:340
				_go_fuzz_dep_.CoverTab[133976]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:340
				// _ = "end of CoverTab[133976]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:340
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:340
			// _ = "end of CoverTab[133974]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:341
		// _ = "end of CoverTab[133971]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:342
	// _ = "end of CoverTab[133968]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:342
	_go_fuzz_dep_.CoverTab[133969]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:343
	// _ = "end of CoverTab[133969]"
}

func (f *FieldDescriptorProto) IsEnum() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:346
	_go_fuzz_dep_.CoverTab[133977]++
														return *f.Type == FieldDescriptorProto_TYPE_ENUM
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:347
	// _ = "end of CoverTab[133977]"
}

func (f *FieldDescriptorProto) IsMessage() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:350
	_go_fuzz_dep_.CoverTab[133978]++
														return *f.Type == FieldDescriptorProto_TYPE_MESSAGE
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:351
	// _ = "end of CoverTab[133978]"
}

func (f *FieldDescriptorProto) IsBytes() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:354
	_go_fuzz_dep_.CoverTab[133979]++
														return *f.Type == FieldDescriptorProto_TYPE_BYTES
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:355
	// _ = "end of CoverTab[133979]"
}

func (f *FieldDescriptorProto) IsRepeated() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:358
	_go_fuzz_dep_.CoverTab[133980]++
														return f.Label != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:359
		_go_fuzz_dep_.CoverTab[133981]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:359
		return *f.Label == FieldDescriptorProto_LABEL_REPEATED
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:359
		// _ = "end of CoverTab[133981]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:359
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:359
	// _ = "end of CoverTab[133980]"
}

func (f *FieldDescriptorProto) IsString() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:362
	_go_fuzz_dep_.CoverTab[133982]++
														return *f.Type == FieldDescriptorProto_TYPE_STRING
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:363
	// _ = "end of CoverTab[133982]"
}

func (f *FieldDescriptorProto) IsBool() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:366
	_go_fuzz_dep_.CoverTab[133983]++
														return *f.Type == FieldDescriptorProto_TYPE_BOOL
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:367
	// _ = "end of CoverTab[133983]"
}

func (f *FieldDescriptorProto) IsRequired() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:370
	_go_fuzz_dep_.CoverTab[133984]++
														return f.Label != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:371
		_go_fuzz_dep_.CoverTab[133985]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:371
		return *f.Label == FieldDescriptorProto_LABEL_REQUIRED
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:371
		// _ = "end of CoverTab[133985]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:371
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:371
	// _ = "end of CoverTab[133984]"
}

func (f *FieldDescriptorProto) IsPacked() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:374
	_go_fuzz_dep_.CoverTab[133986]++
														return f.Options != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:375
		_go_fuzz_dep_.CoverTab[133987]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:375
		return f.GetOptions().GetPacked()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:375
		// _ = "end of CoverTab[133987]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:375
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:375
	// _ = "end of CoverTab[133986]"
}

func (f *FieldDescriptorProto) IsPacked3() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:378
	_go_fuzz_dep_.CoverTab[133988]++
														if f.IsRepeated() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:379
		_go_fuzz_dep_.CoverTab[133990]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:379
		return f.IsScalar()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:379
		// _ = "end of CoverTab[133990]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:379
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:379
		_go_fuzz_dep_.CoverTab[133991]++
															if f.Options == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:380
			_go_fuzz_dep_.CoverTab[133993]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:380
			return f.GetOptions().Packed == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:380
			// _ = "end of CoverTab[133993]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:380
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:380
			_go_fuzz_dep_.CoverTab[133994]++
																return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:381
			// _ = "end of CoverTab[133994]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:382
			_go_fuzz_dep_.CoverTab[133995]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:382
			// _ = "end of CoverTab[133995]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:382
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:382
		// _ = "end of CoverTab[133991]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:382
		_go_fuzz_dep_.CoverTab[133992]++
															return f.Options != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:383
			_go_fuzz_dep_.CoverTab[133996]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:383
			return f.GetOptions().GetPacked()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:383
			// _ = "end of CoverTab[133996]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:383
		}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:383
		// _ = "end of CoverTab[133992]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:384
		_go_fuzz_dep_.CoverTab[133997]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:384
		// _ = "end of CoverTab[133997]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:384
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:384
	// _ = "end of CoverTab[133988]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:384
	_go_fuzz_dep_.CoverTab[133989]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:385
	// _ = "end of CoverTab[133989]"
}

func (m *DescriptorProto) HasExtension() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:388
	_go_fuzz_dep_.CoverTab[133998]++
														return len(m.ExtensionRange) > 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:389
	// _ = "end of CoverTab[133998]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:390
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor/helper.go:390
var _ = _go_fuzz_dep_.CoverTab
