// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 The Go Authors.  All rights reserved.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:32
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:38
import (
	"errors"
)

// errNoMessageTypeID occurs when a protocol buffer does not have a message type ID.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:42
// A message type ID is required for storing a protocol buffer in a message set.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:44
var errNoMessageTypeID = errors.New("proto does not have a message type ID")

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:57
type _MessageSet_Item struct {
	TypeId	*int32	`protobuf:"varint,2,req,name=type_id"`
	Message	[]byte	`protobuf:"bytes,3,req,name=message"`
}

type messageSet struct {
	Item			[]*_MessageSet_Item	`protobuf:"group,1,rep"`
	XXX_unrecognized	[]byte
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:66
}

// Make sure messageSet is a Message.
var _ Message = (*messageSet)(nil)

// messageTypeIder is an interface satisfied by a protocol buffer type
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:71
// that may be stored in a MessageSet.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:73
type messageTypeIder interface {
	MessageTypeId() int32
}

func (ms *messageSet) find(pb Message) *_MessageSet_Item {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:77
	_go_fuzz_dep_.CoverTab[108806]++
												mti, ok := pb.(messageTypeIder)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:79
		_go_fuzz_dep_.CoverTab[108809]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:80
		// _ = "end of CoverTab[108809]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:81
		_go_fuzz_dep_.CoverTab[108810]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:81
		// _ = "end of CoverTab[108810]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:81
	// _ = "end of CoverTab[108806]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:81
	_go_fuzz_dep_.CoverTab[108807]++
												id := mti.MessageTypeId()
												for _, item := range ms.Item {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:83
		_go_fuzz_dep_.CoverTab[108811]++
													if *item.TypeId == id {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:84
			_go_fuzz_dep_.CoverTab[108812]++
														return item
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:85
			// _ = "end of CoverTab[108812]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:86
			_go_fuzz_dep_.CoverTab[108813]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:86
			// _ = "end of CoverTab[108813]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:86
		// _ = "end of CoverTab[108811]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:87
	// _ = "end of CoverTab[108807]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:87
	_go_fuzz_dep_.CoverTab[108808]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:88
	// _ = "end of CoverTab[108808]"
}

func (ms *messageSet) Has(pb Message) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:91
	_go_fuzz_dep_.CoverTab[108814]++
												return ms.find(pb) != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:92
	// _ = "end of CoverTab[108814]"
}

func (ms *messageSet) Unmarshal(pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:95
	_go_fuzz_dep_.CoverTab[108815]++
												if item := ms.find(pb); item != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:96
		_go_fuzz_dep_.CoverTab[108818]++
													return Unmarshal(item.Message, pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:97
		// _ = "end of CoverTab[108818]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:98
		_go_fuzz_dep_.CoverTab[108819]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:98
		// _ = "end of CoverTab[108819]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:98
	// _ = "end of CoverTab[108815]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:98
	_go_fuzz_dep_.CoverTab[108816]++
												if _, ok := pb.(messageTypeIder); !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:99
		_go_fuzz_dep_.CoverTab[108820]++
													return errNoMessageTypeID
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:100
		// _ = "end of CoverTab[108820]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:101
		_go_fuzz_dep_.CoverTab[108821]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:101
		// _ = "end of CoverTab[108821]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:101
	// _ = "end of CoverTab[108816]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:101
	_go_fuzz_dep_.CoverTab[108817]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:102
	// _ = "end of CoverTab[108817]"
}

func (ms *messageSet) Marshal(pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:105
	_go_fuzz_dep_.CoverTab[108822]++
												msg, err := Marshal(pb)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:107
		_go_fuzz_dep_.CoverTab[108826]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:108
		// _ = "end of CoverTab[108826]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:109
		_go_fuzz_dep_.CoverTab[108827]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:109
		// _ = "end of CoverTab[108827]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:109
	// _ = "end of CoverTab[108822]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:109
	_go_fuzz_dep_.CoverTab[108823]++
												if item := ms.find(pb); item != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:110
		_go_fuzz_dep_.CoverTab[108828]++

													item.Message = msg
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:113
		// _ = "end of CoverTab[108828]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:114
		_go_fuzz_dep_.CoverTab[108829]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:114
		// _ = "end of CoverTab[108829]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:114
	// _ = "end of CoverTab[108823]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:114
	_go_fuzz_dep_.CoverTab[108824]++

												mti, ok := pb.(messageTypeIder)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:117
		_go_fuzz_dep_.CoverTab[108830]++
													return errNoMessageTypeID
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:118
		// _ = "end of CoverTab[108830]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:119
		_go_fuzz_dep_.CoverTab[108831]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:119
		// _ = "end of CoverTab[108831]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:119
	// _ = "end of CoverTab[108824]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:119
	_go_fuzz_dep_.CoverTab[108825]++

												mtid := mti.MessageTypeId()
												ms.Item = append(ms.Item, &_MessageSet_Item{
		TypeId:		&mtid,
		Message:	msg,
	})
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:126
	// _ = "end of CoverTab[108825]"
}

func (ms *messageSet) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:129
	_go_fuzz_dep_.CoverTab[108832]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:129
	*ms = messageSet{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:129
	// _ = "end of CoverTab[108832]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:129
}
func (ms *messageSet) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:130
	_go_fuzz_dep_.CoverTab[108833]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:130
	return CompactTextString(ms)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:130
	// _ = "end of CoverTab[108833]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:130
}
func (*messageSet) ProtoMessage()	{ _go_fuzz_dep_.CoverTab[108834]++; // _ = "end of CoverTab[108834]" }

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:135
func skipVarint(buf []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:135
	_go_fuzz_dep_.CoverTab[108835]++
												i := 0
												for ; buf[i]&0x80 != 0; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:137
		_go_fuzz_dep_.CoverTab[108837]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:137
		// _ = "end of CoverTab[108837]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:138
	// _ = "end of CoverTab[108835]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:138
	_go_fuzz_dep_.CoverTab[108836]++
												return buf[i+1:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:139
	// _ = "end of CoverTab[108836]"
}

// unmarshalMessageSet decodes the extension map encoded in buf in the message set wire format.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:142
// It is called by Unmarshal methods on protocol buffer messages with the message_set_wire_format option.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:144
func unmarshalMessageSet(buf []byte, exts interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:144
	_go_fuzz_dep_.CoverTab[108838]++
												var m map[int32]Extension
												switch exts := exts.(type) {
	case *XXX_InternalExtensions:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:147
		_go_fuzz_dep_.CoverTab[108842]++
													m = exts.extensionsWrite()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:148
		// _ = "end of CoverTab[108842]"
	case map[int32]Extension:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:149
		_go_fuzz_dep_.CoverTab[108843]++
													m = exts
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:150
		// _ = "end of CoverTab[108843]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:151
		_go_fuzz_dep_.CoverTab[108844]++
													return errors.New("proto: not an extension map")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:152
		// _ = "end of CoverTab[108844]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:153
	// _ = "end of CoverTab[108838]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:153
	_go_fuzz_dep_.CoverTab[108839]++

												ms := new(messageSet)
												if err := Unmarshal(buf, ms); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:156
		_go_fuzz_dep_.CoverTab[108845]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:157
		// _ = "end of CoverTab[108845]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:158
		_go_fuzz_dep_.CoverTab[108846]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:158
		// _ = "end of CoverTab[108846]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:158
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:158
	// _ = "end of CoverTab[108839]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:158
	_go_fuzz_dep_.CoverTab[108840]++
												for _, item := range ms.Item {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:159
		_go_fuzz_dep_.CoverTab[108847]++
													id := *item.TypeId
													msg := item.Message

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:165
		b := EncodeVarint(uint64(id)<<3 | WireBytes)
		if ext, ok := m[id]; ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:166
			_go_fuzz_dep_.CoverTab[108849]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:170
			o := ext.enc[len(b):]
														_, n := DecodeVarint(o)
														o = o[n:]
														msg = append(o, msg...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:173
			// _ = "end of CoverTab[108849]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:174
			_go_fuzz_dep_.CoverTab[108850]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:174
			// _ = "end of CoverTab[108850]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:174
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:174
		// _ = "end of CoverTab[108847]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:174
		_go_fuzz_dep_.CoverTab[108848]++
													b = append(b, EncodeVarint(uint64(len(msg)))...)
													b = append(b, msg...)

													m[id] = Extension{enc: b}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:178
		// _ = "end of CoverTab[108848]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:179
	// _ = "end of CoverTab[108840]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:179
	_go_fuzz_dep_.CoverTab[108841]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:180
	// _ = "end of CoverTab[108841]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:181
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/message_set.go:181
var _ = _go_fuzz_dep_.CoverTab
