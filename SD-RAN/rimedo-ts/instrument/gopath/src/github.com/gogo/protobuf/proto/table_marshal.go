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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:32
)

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"unicode/utf8"
)

// a sizer takes a pointer to a field and the size of its tag, computes the size of
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:47
// the encoded data.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:49
type sizer func(pointer, int) int

// a marshaler takes a byte slice, a pointer to a field, and its tag (in wire format),
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:51
// marshals the field to the end of the slice, returns the slice and error (if any).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:53
type marshaler func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error)

// marshalInfo is the information used for marshaling a message.
type marshalInfo struct {
	typ		reflect.Type
	fields		[]*marshalFieldInfo
	unrecognized	field				// offset of XXX_unrecognized
	extensions	field				// offset of XXX_InternalExtensions
	v1extensions	field				// offset of XXX_extensions
	sizecache	field				// offset of XXX_sizecache
	initialized	int32				// 0 -- only typ is set, 1 -- fully initialized
	messageset	bool				// uses message set wire format
	hasmarshaler	bool				// has custom marshaler
	sync.RWMutex					// protect extElems map, also for initialization
	extElems	map[int32]*marshalElemInfo	// info of extension elements

	hassizer	bool	// has custom sizer
	hasprotosizer	bool	// has custom protosizer

	bytesExtensions	field	// offset of XXX_extensions where the field type is []byte
}

// marshalFieldInfo is the information used for marshaling a field of a message.
type marshalFieldInfo struct {
	field		field
	wiretag		uint64	// tag in wire format
	tagsize		int	// size of tag in wire format
	sizer		sizer
	marshaler	marshaler
	isPointer	bool
	required	bool					// field is required
	name		string					// name of the field, for error reporting
	oneofElems	map[reflect.Type]*marshalElemInfo	// info of oneof elements
}

// marshalElemInfo is the information used for marshaling an extension or oneof element.
type marshalElemInfo struct {
	wiretag		uint64	// tag in wire format
	tagsize		int	// size of tag in wire format
	sizer		sizer
	marshaler	marshaler
	isptr		bool	// elem is pointer typed, thus interface of this type is a direct interface (extension only)
}

var (
	marshalInfoMap	= map[reflect.Type]*marshalInfo{}
	marshalInfoLock	sync.Mutex

	uint8SliceType	= reflect.TypeOf(([]uint8)(nil)).Kind()
)

// getMarshalInfo returns the information to marshal a given type of message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:104
// The info it returns may not necessarily initialized.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:104
// t is the type of the message (NOT the pointer to it).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:107
func getMarshalInfo(t reflect.Type) *marshalInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:107
	_go_fuzz_dep_.CoverTab[109159]++
												marshalInfoLock.Lock()
												u, ok := marshalInfoMap[t]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:110
		_go_fuzz_dep_.CoverTab[109161]++
													u = &marshalInfo{typ: t}
													marshalInfoMap[t] = u
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:112
		// _ = "end of CoverTab[109161]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:113
		_go_fuzz_dep_.CoverTab[109162]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:113
		// _ = "end of CoverTab[109162]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:113
	// _ = "end of CoverTab[109159]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:113
	_go_fuzz_dep_.CoverTab[109160]++
												marshalInfoLock.Unlock()
												return u
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:115
	// _ = "end of CoverTab[109160]"
}

// Size is the entry point from generated code,
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:118
// and should be ONLY called by generated code.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:118
// It computes the size of encoded data of msg.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:118
// a is a pointer to a place to store cached marshal info.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:122
func (a *InternalMessageInfo) Size(msg Message) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:122
	_go_fuzz_dep_.CoverTab[109163]++
												u := getMessageMarshalInfo(msg, a)
												ptr := toPointer(&msg)
												if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:125
		_go_fuzz_dep_.CoverTab[109165]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:129
		return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:129
		// _ = "end of CoverTab[109165]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:130
		_go_fuzz_dep_.CoverTab[109166]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:130
		// _ = "end of CoverTab[109166]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:130
	// _ = "end of CoverTab[109163]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:130
	_go_fuzz_dep_.CoverTab[109164]++
												return u.size(ptr)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:131
	// _ = "end of CoverTab[109164]"
}

// Marshal is the entry point from generated code,
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:134
// and should be ONLY called by generated code.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:134
// It marshals msg to the end of b.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:134
// a is a pointer to a place to store cached marshal info.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:138
func (a *InternalMessageInfo) Marshal(b []byte, msg Message, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:138
	_go_fuzz_dep_.CoverTab[109167]++
												u := getMessageMarshalInfo(msg, a)
												ptr := toPointer(&msg)
												if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:141
		_go_fuzz_dep_.CoverTab[109169]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:145
		return b, ErrNil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:145
		// _ = "end of CoverTab[109169]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:146
		_go_fuzz_dep_.CoverTab[109170]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:146
		// _ = "end of CoverTab[109170]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:146
	// _ = "end of CoverTab[109167]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:146
	_go_fuzz_dep_.CoverTab[109168]++
												return u.marshal(b, ptr, deterministic)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:147
	// _ = "end of CoverTab[109168]"
}

func getMessageMarshalInfo(msg interface{}, a *InternalMessageInfo) *marshalInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:150
	_go_fuzz_dep_.CoverTab[109171]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:153
	u := atomicLoadMarshalInfo(&a.marshal)
	if u == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:154
		_go_fuzz_dep_.CoverTab[109173]++

													t := reflect.ValueOf(msg).Type()
													if t.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:157
			_go_fuzz_dep_.CoverTab[109175]++
														panic(fmt.Sprintf("cannot handle non-pointer message type %v", t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:158
			// _ = "end of CoverTab[109175]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:159
			_go_fuzz_dep_.CoverTab[109176]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:159
			// _ = "end of CoverTab[109176]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:159
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:159
		// _ = "end of CoverTab[109173]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:159
		_go_fuzz_dep_.CoverTab[109174]++
													u = getMarshalInfo(t.Elem())

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:163
		atomicStoreMarshalInfo(&a.marshal, u)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:163
		// _ = "end of CoverTab[109174]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:164
		_go_fuzz_dep_.CoverTab[109177]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:164
		// _ = "end of CoverTab[109177]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:164
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:164
	// _ = "end of CoverTab[109171]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:164
	_go_fuzz_dep_.CoverTab[109172]++
												return u
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:165
	// _ = "end of CoverTab[109172]"
}

// size is the main function to compute the size of the encoded data of a message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:168
// ptr is the pointer to the message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:170
func (u *marshalInfo) size(ptr pointer) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:170
	_go_fuzz_dep_.CoverTab[109178]++
												if atomic.LoadInt32(&u.initialized) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:171
		_go_fuzz_dep_.CoverTab[109187]++
													u.computeMarshalInfo()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:172
		// _ = "end of CoverTab[109187]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:173
		_go_fuzz_dep_.CoverTab[109188]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:173
		// _ = "end of CoverTab[109188]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:173
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:173
	// _ = "end of CoverTab[109178]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:173
	_go_fuzz_dep_.CoverTab[109179]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:177
	if u.hasmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:177
		_go_fuzz_dep_.CoverTab[109189]++

													if u.hassizer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:179
			_go_fuzz_dep_.CoverTab[109192]++
														s := ptr.asPointerTo(u.typ).Interface().(Sizer)
														return s.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:181
			// _ = "end of CoverTab[109192]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:182
			_go_fuzz_dep_.CoverTab[109193]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:182
			// _ = "end of CoverTab[109193]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:182
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:182
		// _ = "end of CoverTab[109189]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:182
		_go_fuzz_dep_.CoverTab[109190]++

													if u.hasprotosizer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:184
			_go_fuzz_dep_.CoverTab[109194]++
														s := ptr.asPointerTo(u.typ).Interface().(ProtoSizer)
														return s.ProtoSize()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:186
			// _ = "end of CoverTab[109194]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:187
			_go_fuzz_dep_.CoverTab[109195]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:187
			// _ = "end of CoverTab[109195]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:187
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:187
		// _ = "end of CoverTab[109190]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:187
		_go_fuzz_dep_.CoverTab[109191]++

													m := ptr.asPointerTo(u.typ).Interface().(Marshaler)
													b, _ := m.Marshal()
													return len(b)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:191
		// _ = "end of CoverTab[109191]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:192
		_go_fuzz_dep_.CoverTab[109196]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:192
		// _ = "end of CoverTab[109196]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:192
	// _ = "end of CoverTab[109179]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:192
	_go_fuzz_dep_.CoverTab[109180]++

												n := 0
												for _, f := range u.fields {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:195
		_go_fuzz_dep_.CoverTab[109197]++
													if f.isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:196
			_go_fuzz_dep_.CoverTab[109199]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:196
			return ptr.offset(f.field).getPointer().isNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:196
			// _ = "end of CoverTab[109199]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:196
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:196
			_go_fuzz_dep_.CoverTab[109200]++

														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:198
			// _ = "end of CoverTab[109200]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:199
			_go_fuzz_dep_.CoverTab[109201]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:199
			// _ = "end of CoverTab[109201]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:199
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:199
		// _ = "end of CoverTab[109197]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:199
		_go_fuzz_dep_.CoverTab[109198]++
													n += f.sizer(ptr.offset(f.field), f.tagsize)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:200
		// _ = "end of CoverTab[109198]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:201
	// _ = "end of CoverTab[109180]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:201
	_go_fuzz_dep_.CoverTab[109181]++
												if u.extensions.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:202
		_go_fuzz_dep_.CoverTab[109202]++
													e := ptr.offset(u.extensions).toExtensions()
													if u.messageset {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:204
			_go_fuzz_dep_.CoverTab[109203]++
														n += u.sizeMessageSet(e)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:205
			// _ = "end of CoverTab[109203]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:206
			_go_fuzz_dep_.CoverTab[109204]++
														n += u.sizeExtensions(e)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:207
			// _ = "end of CoverTab[109204]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:208
		// _ = "end of CoverTab[109202]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:209
		_go_fuzz_dep_.CoverTab[109205]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:209
		// _ = "end of CoverTab[109205]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:209
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:209
	// _ = "end of CoverTab[109181]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:209
	_go_fuzz_dep_.CoverTab[109182]++
												if u.v1extensions.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:210
		_go_fuzz_dep_.CoverTab[109206]++
													m := *ptr.offset(u.v1extensions).toOldExtensions()
													n += u.sizeV1Extensions(m)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:212
		// _ = "end of CoverTab[109206]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:213
		_go_fuzz_dep_.CoverTab[109207]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:213
		// _ = "end of CoverTab[109207]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:213
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:213
	// _ = "end of CoverTab[109182]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:213
	_go_fuzz_dep_.CoverTab[109183]++
												if u.bytesExtensions.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:214
		_go_fuzz_dep_.CoverTab[109208]++
													s := *ptr.offset(u.bytesExtensions).toBytes()
													n += len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:216
		// _ = "end of CoverTab[109208]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:217
		_go_fuzz_dep_.CoverTab[109209]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:217
		// _ = "end of CoverTab[109209]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:217
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:217
	// _ = "end of CoverTab[109183]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:217
	_go_fuzz_dep_.CoverTab[109184]++
												if u.unrecognized.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:218
		_go_fuzz_dep_.CoverTab[109210]++
													s := *ptr.offset(u.unrecognized).toBytes()
													n += len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:220
		// _ = "end of CoverTab[109210]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:221
		_go_fuzz_dep_.CoverTab[109211]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:221
		// _ = "end of CoverTab[109211]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:221
	// _ = "end of CoverTab[109184]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:221
	_go_fuzz_dep_.CoverTab[109185]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:224
	if u.sizecache.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:224
		_go_fuzz_dep_.CoverTab[109212]++
													atomic.StoreInt32(ptr.offset(u.sizecache).toInt32(), int32(n))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:225
		// _ = "end of CoverTab[109212]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:226
		_go_fuzz_dep_.CoverTab[109213]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:226
		// _ = "end of CoverTab[109213]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:226
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:226
	// _ = "end of CoverTab[109185]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:226
	_go_fuzz_dep_.CoverTab[109186]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:227
	// _ = "end of CoverTab[109186]"
}

// cachedsize gets the size from cache. If there is no cache (i.e. message is not generated),
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:230
// fall back to compute the size.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:232
func (u *marshalInfo) cachedsize(ptr pointer) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:232
	_go_fuzz_dep_.CoverTab[109214]++
												if u.sizecache.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:233
		_go_fuzz_dep_.CoverTab[109216]++
													return int(atomic.LoadInt32(ptr.offset(u.sizecache).toInt32()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:234
		// _ = "end of CoverTab[109216]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:235
		_go_fuzz_dep_.CoverTab[109217]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:235
		// _ = "end of CoverTab[109217]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:235
	// _ = "end of CoverTab[109214]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:235
	_go_fuzz_dep_.CoverTab[109215]++
												return u.size(ptr)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:236
	// _ = "end of CoverTab[109215]"
}

// marshal is the main function to marshal a message. It takes a byte slice and appends
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:239
// the encoded data to the end of the slice, returns the slice and error (if any).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:239
// ptr is the pointer to the message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:239
// If deterministic is true, map is marshaled in deterministic order.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:243
func (u *marshalInfo) marshal(b []byte, ptr pointer, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:243
	_go_fuzz_dep_.CoverTab[109218]++
												if atomic.LoadInt32(&u.initialized) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:244
		_go_fuzz_dep_.CoverTab[109226]++
													u.computeMarshalInfo()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:245
		// _ = "end of CoverTab[109226]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:246
		_go_fuzz_dep_.CoverTab[109227]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:246
		// _ = "end of CoverTab[109227]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:246
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:246
	// _ = "end of CoverTab[109218]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:246
	_go_fuzz_dep_.CoverTab[109219]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:250
	if u.hasmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:250
		_go_fuzz_dep_.CoverTab[109228]++
													m := ptr.asPointerTo(u.typ).Interface().(Marshaler)
													b1, err := m.Marshal()
													b = append(b, b1...)
													return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:254
		// _ = "end of CoverTab[109228]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:255
		_go_fuzz_dep_.CoverTab[109229]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:255
		// _ = "end of CoverTab[109229]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:255
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:255
	// _ = "end of CoverTab[109219]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:255
	_go_fuzz_dep_.CoverTab[109220]++

												var err, errLater error

												if u.extensions.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:259
		_go_fuzz_dep_.CoverTab[109230]++
													e := ptr.offset(u.extensions).toExtensions()
													if u.messageset {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:261
			_go_fuzz_dep_.CoverTab[109232]++
														b, err = u.appendMessageSet(b, e, deterministic)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:262
			// _ = "end of CoverTab[109232]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:263
			_go_fuzz_dep_.CoverTab[109233]++
														b, err = u.appendExtensions(b, e, deterministic)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:264
			// _ = "end of CoverTab[109233]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:265
		// _ = "end of CoverTab[109230]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:265
		_go_fuzz_dep_.CoverTab[109231]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:266
			_go_fuzz_dep_.CoverTab[109234]++
														return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:267
			// _ = "end of CoverTab[109234]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:268
			_go_fuzz_dep_.CoverTab[109235]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:268
			// _ = "end of CoverTab[109235]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:268
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:268
		// _ = "end of CoverTab[109231]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:269
		_go_fuzz_dep_.CoverTab[109236]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:269
		// _ = "end of CoverTab[109236]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:269
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:269
	// _ = "end of CoverTab[109220]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:269
	_go_fuzz_dep_.CoverTab[109221]++
												if u.v1extensions.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:270
		_go_fuzz_dep_.CoverTab[109237]++
													m := *ptr.offset(u.v1extensions).toOldExtensions()
													b, err = u.appendV1Extensions(b, m, deterministic)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:273
			_go_fuzz_dep_.CoverTab[109238]++
														return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:274
			// _ = "end of CoverTab[109238]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:275
			_go_fuzz_dep_.CoverTab[109239]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:275
			// _ = "end of CoverTab[109239]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:275
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:275
		// _ = "end of CoverTab[109237]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:276
		_go_fuzz_dep_.CoverTab[109240]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:276
		// _ = "end of CoverTab[109240]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:276
	// _ = "end of CoverTab[109221]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:276
	_go_fuzz_dep_.CoverTab[109222]++
												if u.bytesExtensions.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:277
		_go_fuzz_dep_.CoverTab[109241]++
													s := *ptr.offset(u.bytesExtensions).toBytes()
													b = append(b, s...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:279
		// _ = "end of CoverTab[109241]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:280
		_go_fuzz_dep_.CoverTab[109242]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:280
		// _ = "end of CoverTab[109242]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:280
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:280
	// _ = "end of CoverTab[109222]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:280
	_go_fuzz_dep_.CoverTab[109223]++
												for _, f := range u.fields {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:281
		_go_fuzz_dep_.CoverTab[109243]++
													if f.required {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:282
			_go_fuzz_dep_.CoverTab[109246]++
														if f.isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:283
				_go_fuzz_dep_.CoverTab[109247]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:283
				return ptr.offset(f.field).getPointer().isNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:283
				// _ = "end of CoverTab[109247]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:283
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:283
				_go_fuzz_dep_.CoverTab[109248]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:286
				if errLater == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:286
					_go_fuzz_dep_.CoverTab[109250]++
																errLater = &RequiredNotSetError{f.name}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:287
					// _ = "end of CoverTab[109250]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:288
					_go_fuzz_dep_.CoverTab[109251]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:288
					// _ = "end of CoverTab[109251]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:288
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:288
				// _ = "end of CoverTab[109248]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:288
				_go_fuzz_dep_.CoverTab[109249]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:289
				// _ = "end of CoverTab[109249]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:290
				_go_fuzz_dep_.CoverTab[109252]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:290
				// _ = "end of CoverTab[109252]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:290
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:290
			// _ = "end of CoverTab[109246]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:291
			_go_fuzz_dep_.CoverTab[109253]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:291
			// _ = "end of CoverTab[109253]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:291
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:291
		// _ = "end of CoverTab[109243]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:291
		_go_fuzz_dep_.CoverTab[109244]++
													if f.isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:292
			_go_fuzz_dep_.CoverTab[109254]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:292
			return ptr.offset(f.field).getPointer().isNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:292
			// _ = "end of CoverTab[109254]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:292
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:292
			_go_fuzz_dep_.CoverTab[109255]++

														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:294
			// _ = "end of CoverTab[109255]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:295
			_go_fuzz_dep_.CoverTab[109256]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:295
			// _ = "end of CoverTab[109256]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:295
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:295
		// _ = "end of CoverTab[109244]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:295
		_go_fuzz_dep_.CoverTab[109245]++
													b, err = f.marshaler(b, ptr.offset(f.field), f.wiretag, deterministic)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:297
			_go_fuzz_dep_.CoverTab[109257]++
														if err1, ok := err.(*RequiredNotSetError); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:298
				_go_fuzz_dep_.CoverTab[109261]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:301
				if errLater == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:301
					_go_fuzz_dep_.CoverTab[109263]++
																errLater = &RequiredNotSetError{f.name + "." + err1.field}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:302
					// _ = "end of CoverTab[109263]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:303
					_go_fuzz_dep_.CoverTab[109264]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:303
					// _ = "end of CoverTab[109264]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:303
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:303
				// _ = "end of CoverTab[109261]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:303
				_go_fuzz_dep_.CoverTab[109262]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:304
				// _ = "end of CoverTab[109262]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:305
				_go_fuzz_dep_.CoverTab[109265]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:305
				// _ = "end of CoverTab[109265]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:305
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:305
			// _ = "end of CoverTab[109257]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:305
			_go_fuzz_dep_.CoverTab[109258]++
														if err == errRepeatedHasNil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:306
				_go_fuzz_dep_.CoverTab[109266]++
															err = errors.New("proto: repeated field " + f.name + " has nil element")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:307
				// _ = "end of CoverTab[109266]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:308
				_go_fuzz_dep_.CoverTab[109267]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:308
				// _ = "end of CoverTab[109267]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:308
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:308
			// _ = "end of CoverTab[109258]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:308
			_go_fuzz_dep_.CoverTab[109259]++
														if err == errInvalidUTF8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:309
				_go_fuzz_dep_.CoverTab[109268]++
															if errLater == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:310
					_go_fuzz_dep_.CoverTab[109270]++
																fullName := revProtoTypes[reflect.PtrTo(u.typ)] + "." + f.name
																errLater = &invalidUTF8Error{fullName}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:312
					// _ = "end of CoverTab[109270]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:313
					_go_fuzz_dep_.CoverTab[109271]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:313
					// _ = "end of CoverTab[109271]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:313
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:313
				// _ = "end of CoverTab[109268]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:313
				_go_fuzz_dep_.CoverTab[109269]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:314
				// _ = "end of CoverTab[109269]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:315
				_go_fuzz_dep_.CoverTab[109272]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:315
				// _ = "end of CoverTab[109272]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:315
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:315
			// _ = "end of CoverTab[109259]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:315
			_go_fuzz_dep_.CoverTab[109260]++
														return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:316
			// _ = "end of CoverTab[109260]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:317
			_go_fuzz_dep_.CoverTab[109273]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:317
			// _ = "end of CoverTab[109273]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:317
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:317
		// _ = "end of CoverTab[109245]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:318
	// _ = "end of CoverTab[109223]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:318
	_go_fuzz_dep_.CoverTab[109224]++
												if u.unrecognized.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:319
		_go_fuzz_dep_.CoverTab[109274]++
													s := *ptr.offset(u.unrecognized).toBytes()
													b = append(b, s...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:321
		// _ = "end of CoverTab[109274]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:322
		_go_fuzz_dep_.CoverTab[109275]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:322
		// _ = "end of CoverTab[109275]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:322
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:322
	// _ = "end of CoverTab[109224]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:322
	_go_fuzz_dep_.CoverTab[109225]++
												return b, errLater
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:323
	// _ = "end of CoverTab[109225]"
}

// computeMarshalInfo initializes the marshal info.
func (u *marshalInfo) computeMarshalInfo() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:327
	_go_fuzz_dep_.CoverTab[109276]++
												u.Lock()
												defer u.Unlock()
												if u.initialized != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:330
		_go_fuzz_dep_.CoverTab[109284]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:331
		// _ = "end of CoverTab[109284]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:332
		_go_fuzz_dep_.CoverTab[109285]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:332
		// _ = "end of CoverTab[109285]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:332
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:332
	// _ = "end of CoverTab[109276]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:332
	_go_fuzz_dep_.CoverTab[109277]++

												t := u.typ
												u.unrecognized = invalidField
												u.extensions = invalidField
												u.v1extensions = invalidField
												u.bytesExtensions = invalidField
												u.sizecache = invalidField
												isOneofMessage := false

												if reflect.PtrTo(t).Implements(sizerType) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:342
		_go_fuzz_dep_.CoverTab[109286]++
													u.hassizer = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:343
		// _ = "end of CoverTab[109286]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:344
		_go_fuzz_dep_.CoverTab[109287]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:344
		// _ = "end of CoverTab[109287]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:344
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:344
	// _ = "end of CoverTab[109277]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:344
	_go_fuzz_dep_.CoverTab[109278]++
												if reflect.PtrTo(t).Implements(protosizerType) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:345
		_go_fuzz_dep_.CoverTab[109288]++
													u.hasprotosizer = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:346
		// _ = "end of CoverTab[109288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:347
		_go_fuzz_dep_.CoverTab[109289]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:347
		// _ = "end of CoverTab[109289]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:347
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:347
	// _ = "end of CoverTab[109278]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:347
	_go_fuzz_dep_.CoverTab[109279]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:350
	if reflect.PtrTo(t).Implements(marshalerType) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:350
		_go_fuzz_dep_.CoverTab[109290]++
													u.hasmarshaler = true
													atomic.StoreInt32(&u.initialized, 1)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:353
		// _ = "end of CoverTab[109290]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:354
		_go_fuzz_dep_.CoverTab[109291]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:354
		// _ = "end of CoverTab[109291]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:354
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:354
	// _ = "end of CoverTab[109279]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:354
	_go_fuzz_dep_.CoverTab[109280]++

												n := t.NumField()

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:359
	for i := 0; i < t.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:359
		_go_fuzz_dep_.CoverTab[109292]++
													f := t.Field(i)
													if f.Tag.Get("protobuf_oneof") != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:361
			_go_fuzz_dep_.CoverTab[109296]++
														isOneofMessage = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:362
			// _ = "end of CoverTab[109296]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:363
			_go_fuzz_dep_.CoverTab[109297]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:363
			// _ = "end of CoverTab[109297]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:363
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:363
		// _ = "end of CoverTab[109292]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:363
		_go_fuzz_dep_.CoverTab[109293]++
													if !strings.HasPrefix(f.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:364
			_go_fuzz_dep_.CoverTab[109298]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:365
			// _ = "end of CoverTab[109298]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:366
			_go_fuzz_dep_.CoverTab[109299]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:366
			// _ = "end of CoverTab[109299]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:366
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:366
		// _ = "end of CoverTab[109293]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:366
		_go_fuzz_dep_.CoverTab[109294]++
													switch f.Name {
		case "XXX_sizecache":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:368
			_go_fuzz_dep_.CoverTab[109300]++
														u.sizecache = toField(&f)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:369
			// _ = "end of CoverTab[109300]"
		case "XXX_unrecognized":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:370
			_go_fuzz_dep_.CoverTab[109301]++
														u.unrecognized = toField(&f)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:371
			// _ = "end of CoverTab[109301]"
		case "XXX_InternalExtensions":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:372
			_go_fuzz_dep_.CoverTab[109302]++
														u.extensions = toField(&f)
														u.messageset = f.Tag.Get("protobuf_messageset") == "1"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:374
			// _ = "end of CoverTab[109302]"
		case "XXX_extensions":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:375
			_go_fuzz_dep_.CoverTab[109303]++
														if f.Type.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:376
				_go_fuzz_dep_.CoverTab[109306]++
															u.v1extensions = toField(&f)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:377
				// _ = "end of CoverTab[109306]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:378
				_go_fuzz_dep_.CoverTab[109307]++
															u.bytesExtensions = toField(&f)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:379
				// _ = "end of CoverTab[109307]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:380
			// _ = "end of CoverTab[109303]"
		case "XXX_NoUnkeyedLiteral":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:381
			_go_fuzz_dep_.CoverTab[109304]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:381
			// _ = "end of CoverTab[109304]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:383
			_go_fuzz_dep_.CoverTab[109305]++
														panic("unknown XXX field: " + f.Name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:384
			// _ = "end of CoverTab[109305]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:385
		// _ = "end of CoverTab[109294]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:385
		_go_fuzz_dep_.CoverTab[109295]++
													n--
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:386
		// _ = "end of CoverTab[109295]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:387
	// _ = "end of CoverTab[109280]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:387
	_go_fuzz_dep_.CoverTab[109281]++

	// get oneof implementers
	var oneofImplementers []interface{}

	if isOneofMessage {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:392
		_go_fuzz_dep_.CoverTab[109308]++
													switch m := reflect.Zero(reflect.PtrTo(t)).Interface().(type) {
		case oneofFuncsIface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:394
			_go_fuzz_dep_.CoverTab[109309]++
														_, _, _, oneofImplementers = m.XXX_OneofFuncs()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:395
			// _ = "end of CoverTab[109309]"
		case oneofWrappersIface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:396
			_go_fuzz_dep_.CoverTab[109310]++
														oneofImplementers = m.XXX_OneofWrappers()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:397
			// _ = "end of CoverTab[109310]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:398
		// _ = "end of CoverTab[109308]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:399
		_go_fuzz_dep_.CoverTab[109311]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:399
		// _ = "end of CoverTab[109311]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:399
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:399
	// _ = "end of CoverTab[109281]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:399
	_go_fuzz_dep_.CoverTab[109282]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:402
	fields := make([]marshalFieldInfo, n)
	u.fields = make([]*marshalFieldInfo, 0, n)
	for i, j := 0, 0; i < t.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:404
		_go_fuzz_dep_.CoverTab[109312]++
													f := t.Field(i)

													if strings.HasPrefix(f.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:407
			_go_fuzz_dep_.CoverTab[109316]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:408
			// _ = "end of CoverTab[109316]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:409
			_go_fuzz_dep_.CoverTab[109317]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:409
			// _ = "end of CoverTab[109317]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:409
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:409
		// _ = "end of CoverTab[109312]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:409
		_go_fuzz_dep_.CoverTab[109313]++
													field := &fields[j]
													j++
													field.name = f.Name
													u.fields = append(u.fields, field)
													if f.Tag.Get("protobuf_oneof") != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:414
			_go_fuzz_dep_.CoverTab[109318]++
														field.computeOneofFieldInfo(&f, oneofImplementers)
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:416
			// _ = "end of CoverTab[109318]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:417
			_go_fuzz_dep_.CoverTab[109319]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:417
			// _ = "end of CoverTab[109319]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:417
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:417
		// _ = "end of CoverTab[109313]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:417
		_go_fuzz_dep_.CoverTab[109314]++
													if f.Tag.Get("protobuf") == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:418
			_go_fuzz_dep_.CoverTab[109320]++

														u.fields = u.fields[:len(u.fields)-1]
														j--
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:422
			// _ = "end of CoverTab[109320]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:423
			_go_fuzz_dep_.CoverTab[109321]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:423
			// _ = "end of CoverTab[109321]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:423
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:423
		// _ = "end of CoverTab[109314]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:423
		_go_fuzz_dep_.CoverTab[109315]++
													field.computeMarshalFieldInfo(&f)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:424
		// _ = "end of CoverTab[109315]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:425
	// _ = "end of CoverTab[109282]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:425
	_go_fuzz_dep_.CoverTab[109283]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:428
	sort.Sort(byTag(u.fields))

												atomic.StoreInt32(&u.initialized, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:430
	// _ = "end of CoverTab[109283]"
}

// helper for sorting fields by tag
type byTag []*marshalFieldInfo

func (a byTag) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:436
	_go_fuzz_dep_.CoverTab[109322]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:436
	return len(a)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:436
	// _ = "end of CoverTab[109322]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:436
}
func (a byTag) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:437
	_go_fuzz_dep_.CoverTab[109323]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:437
	a[i], a[j] = a[j], a[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:437
	// _ = "end of CoverTab[109323]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:437
}
func (a byTag) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:438
	_go_fuzz_dep_.CoverTab[109324]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:438
	return a[i].wiretag < a[j].wiretag
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:438
	// _ = "end of CoverTab[109324]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:438
}

// getExtElemInfo returns the information to marshal an extension element.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:440
// The info it returns is initialized.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:442
func (u *marshalInfo) getExtElemInfo(desc *ExtensionDesc) *marshalElemInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:442
	_go_fuzz_dep_.CoverTab[109325]++

												u.RLock()
												e, ok := u.extElems[desc.Field]
												u.RUnlock()
												if ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:447
		_go_fuzz_dep_.CoverTab[109329]++
													return e
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:448
		// _ = "end of CoverTab[109329]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:449
		_go_fuzz_dep_.CoverTab[109330]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:449
		// _ = "end of CoverTab[109330]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:449
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:449
	// _ = "end of CoverTab[109325]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:449
	_go_fuzz_dep_.CoverTab[109326]++

												t := reflect.TypeOf(desc.ExtensionType)
												tags := strings.Split(desc.Tag, ",")
												tag, err := strconv.Atoi(tags[1])
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:454
		_go_fuzz_dep_.CoverTab[109331]++
													panic("tag is not an integer")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:455
		// _ = "end of CoverTab[109331]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:456
		_go_fuzz_dep_.CoverTab[109332]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:456
		// _ = "end of CoverTab[109332]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:456
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:456
	// _ = "end of CoverTab[109326]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:456
	_go_fuzz_dep_.CoverTab[109327]++
												wt := wiretype(tags[0])
												sizr, marshalr := typeMarshaler(t, tags, false, false)
												e = &marshalElemInfo{
		wiretag:	uint64(tag)<<3 | wt,
		tagsize:	SizeVarint(uint64(tag) << 3),
		sizer:		sizr,
		marshaler:	marshalr,
		isptr:		t.Kind() == reflect.Ptr,
	}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:468
	u.Lock()
	if u.extElems == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:469
		_go_fuzz_dep_.CoverTab[109333]++
													u.extElems = make(map[int32]*marshalElemInfo)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:470
		// _ = "end of CoverTab[109333]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:471
		_go_fuzz_dep_.CoverTab[109334]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:471
		// _ = "end of CoverTab[109334]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:471
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:471
	// _ = "end of CoverTab[109327]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:471
	_go_fuzz_dep_.CoverTab[109328]++
												u.extElems[desc.Field] = e
												u.Unlock()
												return e
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:474
	// _ = "end of CoverTab[109328]"
}

// computeMarshalFieldInfo fills up the information to marshal a field.
func (fi *marshalFieldInfo) computeMarshalFieldInfo(f *reflect.StructField) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:478
	_go_fuzz_dep_.CoverTab[109335]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:481
	tags := strings.Split(f.Tag.Get("protobuf"), ",")
	if tags[0] == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:482
		_go_fuzz_dep_.CoverTab[109339]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:483
		// _ = "end of CoverTab[109339]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:484
		_go_fuzz_dep_.CoverTab[109340]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:484
		// _ = "end of CoverTab[109340]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:484
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:484
	// _ = "end of CoverTab[109335]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:484
	_go_fuzz_dep_.CoverTab[109336]++
												tag, err := strconv.Atoi(tags[1])
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:486
		_go_fuzz_dep_.CoverTab[109341]++
													panic("tag is not an integer")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:487
		// _ = "end of CoverTab[109341]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:488
		_go_fuzz_dep_.CoverTab[109342]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:488
		// _ = "end of CoverTab[109342]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:488
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:488
	// _ = "end of CoverTab[109336]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:488
	_go_fuzz_dep_.CoverTab[109337]++
												wt := wiretype(tags[0])
												if tags[2] == "req" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:490
		_go_fuzz_dep_.CoverTab[109343]++
													fi.required = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:491
		// _ = "end of CoverTab[109343]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:492
		_go_fuzz_dep_.CoverTab[109344]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:492
		// _ = "end of CoverTab[109344]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:492
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:492
	// _ = "end of CoverTab[109337]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:492
	_go_fuzz_dep_.CoverTab[109338]++
												fi.setTag(f, tag, wt)
												fi.setMarshaler(f, tags)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:494
	// _ = "end of CoverTab[109338]"
}

func (fi *marshalFieldInfo) computeOneofFieldInfo(f *reflect.StructField, oneofImplementers []interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:497
	_go_fuzz_dep_.CoverTab[109345]++
												fi.field = toField(f)
												fi.wiretag = math.MaxInt32
												fi.isPointer = true
												fi.sizer, fi.marshaler = makeOneOfMarshaler(fi, f)
												fi.oneofElems = make(map[reflect.Type]*marshalElemInfo)

												ityp := f.Type
												for _, o := range oneofImplementers {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:505
		_go_fuzz_dep_.CoverTab[109346]++
													t := reflect.TypeOf(o)
													if !t.Implements(ityp) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:507
			_go_fuzz_dep_.CoverTab[109349]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:508
			// _ = "end of CoverTab[109349]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:509
			_go_fuzz_dep_.CoverTab[109350]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:509
			// _ = "end of CoverTab[109350]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:509
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:509
		// _ = "end of CoverTab[109346]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:509
		_go_fuzz_dep_.CoverTab[109347]++
													sf := t.Elem().Field(0)
													tags := strings.Split(sf.Tag.Get("protobuf"), ",")
													tag, err := strconv.Atoi(tags[1])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:513
			_go_fuzz_dep_.CoverTab[109351]++
														panic("tag is not an integer")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:514
			// _ = "end of CoverTab[109351]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:515
			_go_fuzz_dep_.CoverTab[109352]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:515
			// _ = "end of CoverTab[109352]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:515
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:515
		// _ = "end of CoverTab[109347]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:515
		_go_fuzz_dep_.CoverTab[109348]++
													wt := wiretype(tags[0])
													sizr, marshalr := typeMarshaler(sf.Type, tags, false, true)
													fi.oneofElems[t.Elem()] = &marshalElemInfo{
			wiretag:	uint64(tag)<<3 | wt,
			tagsize:	SizeVarint(uint64(tag) << 3),
			sizer:		sizr,
			marshaler:	marshalr,
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:523
		// _ = "end of CoverTab[109348]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:524
	// _ = "end of CoverTab[109345]"
}

// wiretype returns the wire encoding of the type.
func wiretype(encoding string) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:528
	_go_fuzz_dep_.CoverTab[109353]++
												switch encoding {
	case "fixed32":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:530
		_go_fuzz_dep_.CoverTab[109355]++
													return WireFixed32
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:531
		// _ = "end of CoverTab[109355]"
	case "fixed64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:532
		_go_fuzz_dep_.CoverTab[109356]++
													return WireFixed64
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:533
		// _ = "end of CoverTab[109356]"
	case "varint", "zigzag32", "zigzag64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:534
		_go_fuzz_dep_.CoverTab[109357]++
													return WireVarint
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:535
		// _ = "end of CoverTab[109357]"
	case "bytes":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:536
		_go_fuzz_dep_.CoverTab[109358]++
													return WireBytes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:537
		// _ = "end of CoverTab[109358]"
	case "group":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:538
		_go_fuzz_dep_.CoverTab[109359]++
													return WireStartGroup
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:539
		// _ = "end of CoverTab[109359]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:539
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:539
		_go_fuzz_dep_.CoverTab[109360]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:539
		// _ = "end of CoverTab[109360]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:540
	// _ = "end of CoverTab[109353]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:540
	_go_fuzz_dep_.CoverTab[109354]++
												panic("unknown wire type " + encoding)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:541
	// _ = "end of CoverTab[109354]"
}

// setTag fills up the tag (in wire format) and its size in the info of a field.
func (fi *marshalFieldInfo) setTag(f *reflect.StructField, tag int, wt uint64) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:545
	_go_fuzz_dep_.CoverTab[109361]++
												fi.field = toField(f)
												fi.wiretag = uint64(tag)<<3 | wt
												fi.tagsize = SizeVarint(uint64(tag) << 3)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:548
	// _ = "end of CoverTab[109361]"
}

// setMarshaler fills up the sizer and marshaler in the info of a field.
func (fi *marshalFieldInfo) setMarshaler(f *reflect.StructField, tags []string) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:552
	_go_fuzz_dep_.CoverTab[109362]++
												switch f.Type.Kind() {
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:554
		_go_fuzz_dep_.CoverTab[109364]++

													fi.isPointer = true
													fi.sizer, fi.marshaler = makeMapMarshaler(f)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:558
		// _ = "end of CoverTab[109364]"
	case reflect.Ptr, reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:559
		_go_fuzz_dep_.CoverTab[109365]++
													fi.isPointer = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:560
		// _ = "end of CoverTab[109365]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:560
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:560
		_go_fuzz_dep_.CoverTab[109366]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:560
		// _ = "end of CoverTab[109366]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:561
	// _ = "end of CoverTab[109362]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:561
	_go_fuzz_dep_.CoverTab[109363]++
												fi.sizer, fi.marshaler = typeMarshaler(f.Type, tags, true, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:562
	// _ = "end of CoverTab[109363]"
}

// typeMarshaler returns the sizer and marshaler of a given field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:565
// t is the type of the field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:565
// tags is the generated "protobuf" tag of the field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:565
// If nozero is true, zero value is not marshaled to the wire.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:565
// If oneof is true, it is a oneof field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:570
func typeMarshaler(t reflect.Type, tags []string, nozero, oneof bool) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:570
	_go_fuzz_dep_.CoverTab[109367]++
												encoding := tags[0]

												pointer := false
												slice := false
												if t.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:575
		_go_fuzz_dep_.CoverTab[109377]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:575
		return t.Elem().Kind() != reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:575
		// _ = "end of CoverTab[109377]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:575
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:575
		_go_fuzz_dep_.CoverTab[109378]++
													slice = true
													t = t.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:577
		// _ = "end of CoverTab[109378]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:578
		_go_fuzz_dep_.CoverTab[109379]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:578
		// _ = "end of CoverTab[109379]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:578
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:578
	// _ = "end of CoverTab[109367]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:578
	_go_fuzz_dep_.CoverTab[109368]++
												if t.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:579
		_go_fuzz_dep_.CoverTab[109380]++
													pointer = true
													t = t.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:581
		// _ = "end of CoverTab[109380]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:582
		_go_fuzz_dep_.CoverTab[109381]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:582
		// _ = "end of CoverTab[109381]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:582
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:582
	// _ = "end of CoverTab[109368]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:582
	_go_fuzz_dep_.CoverTab[109369]++

												packed := false
												proto3 := false
												ctype := false
												isTime := false
												isDuration := false
												isWktPointer := false
												validateUTF8 := true
												for i := 2; i < len(tags); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:591
		_go_fuzz_dep_.CoverTab[109382]++
													if tags[i] == "packed" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:592
			_go_fuzz_dep_.CoverTab[109388]++
														packed = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:593
			// _ = "end of CoverTab[109388]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:594
			_go_fuzz_dep_.CoverTab[109389]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:594
			// _ = "end of CoverTab[109389]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:594
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:594
		// _ = "end of CoverTab[109382]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:594
		_go_fuzz_dep_.CoverTab[109383]++
													if tags[i] == "proto3" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:595
			_go_fuzz_dep_.CoverTab[109390]++
														proto3 = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:596
			// _ = "end of CoverTab[109390]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:597
			_go_fuzz_dep_.CoverTab[109391]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:597
			// _ = "end of CoverTab[109391]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:597
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:597
		// _ = "end of CoverTab[109383]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:597
		_go_fuzz_dep_.CoverTab[109384]++
													if strings.HasPrefix(tags[i], "customtype=") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:598
			_go_fuzz_dep_.CoverTab[109392]++
														ctype = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:599
			// _ = "end of CoverTab[109392]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:600
			_go_fuzz_dep_.CoverTab[109393]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:600
			// _ = "end of CoverTab[109393]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:600
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:600
		// _ = "end of CoverTab[109384]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:600
		_go_fuzz_dep_.CoverTab[109385]++
													if tags[i] == "stdtime" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:601
			_go_fuzz_dep_.CoverTab[109394]++
														isTime = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:602
			// _ = "end of CoverTab[109394]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:603
			_go_fuzz_dep_.CoverTab[109395]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:603
			// _ = "end of CoverTab[109395]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:603
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:603
		// _ = "end of CoverTab[109385]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:603
		_go_fuzz_dep_.CoverTab[109386]++
													if tags[i] == "stdduration" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:604
			_go_fuzz_dep_.CoverTab[109396]++
														isDuration = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:605
			// _ = "end of CoverTab[109396]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:606
			_go_fuzz_dep_.CoverTab[109397]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:606
			// _ = "end of CoverTab[109397]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:606
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:606
		// _ = "end of CoverTab[109386]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:606
		_go_fuzz_dep_.CoverTab[109387]++
													if tags[i] == "wktptr" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:607
			_go_fuzz_dep_.CoverTab[109398]++
														isWktPointer = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:608
			// _ = "end of CoverTab[109398]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:609
			_go_fuzz_dep_.CoverTab[109399]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:609
			// _ = "end of CoverTab[109399]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:609
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:609
		// _ = "end of CoverTab[109387]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:610
	// _ = "end of CoverTab[109369]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:610
	_go_fuzz_dep_.CoverTab[109370]++
												validateUTF8 = validateUTF8 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:611
		_go_fuzz_dep_.CoverTab[109400]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:611
		return proto3
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:611
		// _ = "end of CoverTab[109400]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:611
	}()
												if !proto3 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:612
		_go_fuzz_dep_.CoverTab[109401]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:612
		return !pointer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:612
		// _ = "end of CoverTab[109401]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:612
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:612
		_go_fuzz_dep_.CoverTab[109402]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:612
		return !slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:612
		// _ = "end of CoverTab[109402]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:612
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:612
		_go_fuzz_dep_.CoverTab[109403]++
													nozero = false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:613
		// _ = "end of CoverTab[109403]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:614
		_go_fuzz_dep_.CoverTab[109404]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:614
		// _ = "end of CoverTab[109404]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:614
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:614
	// _ = "end of CoverTab[109370]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:614
	_go_fuzz_dep_.CoverTab[109371]++

												if ctype {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:616
		_go_fuzz_dep_.CoverTab[109405]++
													if reflect.PtrTo(t).Implements(customType) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:617
			_go_fuzz_dep_.CoverTab[109406]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:618
				_go_fuzz_dep_.CoverTab[109409]++
															return makeMessageRefSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:619
				// _ = "end of CoverTab[109409]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:620
				_go_fuzz_dep_.CoverTab[109410]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:620
				// _ = "end of CoverTab[109410]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:620
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:620
			// _ = "end of CoverTab[109406]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:620
			_go_fuzz_dep_.CoverTab[109407]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:621
				_go_fuzz_dep_.CoverTab[109411]++
															return makeCustomPtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:622
				// _ = "end of CoverTab[109411]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:623
				_go_fuzz_dep_.CoverTab[109412]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:623
				// _ = "end of CoverTab[109412]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:623
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:623
			// _ = "end of CoverTab[109407]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:623
			_go_fuzz_dep_.CoverTab[109408]++
														return makeCustomMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:624
			// _ = "end of CoverTab[109408]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:625
			_go_fuzz_dep_.CoverTab[109413]++
														panic(fmt.Sprintf("custom type: type: %v, does not implement the proto.custom interface", t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:626
			// _ = "end of CoverTab[109413]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:627
		// _ = "end of CoverTab[109405]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:628
		_go_fuzz_dep_.CoverTab[109414]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:628
		// _ = "end of CoverTab[109414]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:628
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:628
	// _ = "end of CoverTab[109371]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:628
	_go_fuzz_dep_.CoverTab[109372]++

												if isTime {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:630
		_go_fuzz_dep_.CoverTab[109415]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:631
			_go_fuzz_dep_.CoverTab[109418]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:632
				_go_fuzz_dep_.CoverTab[109420]++
															return makeTimePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:633
				// _ = "end of CoverTab[109420]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:634
				_go_fuzz_dep_.CoverTab[109421]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:634
				// _ = "end of CoverTab[109421]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:634
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:634
			// _ = "end of CoverTab[109418]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:634
			_go_fuzz_dep_.CoverTab[109419]++
														return makeTimePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:635
			// _ = "end of CoverTab[109419]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:636
			_go_fuzz_dep_.CoverTab[109422]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:636
			// _ = "end of CoverTab[109422]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:636
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:636
		// _ = "end of CoverTab[109415]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:636
		_go_fuzz_dep_.CoverTab[109416]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:637
			_go_fuzz_dep_.CoverTab[109423]++
														return makeTimeSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:638
			// _ = "end of CoverTab[109423]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:639
			_go_fuzz_dep_.CoverTab[109424]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:639
			// _ = "end of CoverTab[109424]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:639
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:639
		// _ = "end of CoverTab[109416]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:639
		_go_fuzz_dep_.CoverTab[109417]++
													return makeTimeMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:640
		// _ = "end of CoverTab[109417]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:641
		_go_fuzz_dep_.CoverTab[109425]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:641
		// _ = "end of CoverTab[109425]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:641
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:641
	// _ = "end of CoverTab[109372]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:641
	_go_fuzz_dep_.CoverTab[109373]++

												if isDuration {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:643
		_go_fuzz_dep_.CoverTab[109426]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:644
			_go_fuzz_dep_.CoverTab[109429]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:645
				_go_fuzz_dep_.CoverTab[109431]++
															return makeDurationPtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:646
				// _ = "end of CoverTab[109431]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:647
				_go_fuzz_dep_.CoverTab[109432]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:647
				// _ = "end of CoverTab[109432]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:647
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:647
			// _ = "end of CoverTab[109429]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:647
			_go_fuzz_dep_.CoverTab[109430]++
														return makeDurationPtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:648
			// _ = "end of CoverTab[109430]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:649
			_go_fuzz_dep_.CoverTab[109433]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:649
			// _ = "end of CoverTab[109433]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:649
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:649
		// _ = "end of CoverTab[109426]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:649
		_go_fuzz_dep_.CoverTab[109427]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:650
			_go_fuzz_dep_.CoverTab[109434]++
														return makeDurationSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:651
			// _ = "end of CoverTab[109434]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:652
			_go_fuzz_dep_.CoverTab[109435]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:652
			// _ = "end of CoverTab[109435]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:652
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:652
		// _ = "end of CoverTab[109427]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:652
		_go_fuzz_dep_.CoverTab[109428]++
													return makeDurationMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:653
		// _ = "end of CoverTab[109428]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:654
		_go_fuzz_dep_.CoverTab[109436]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:654
		// _ = "end of CoverTab[109436]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:654
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:654
	// _ = "end of CoverTab[109373]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:654
	_go_fuzz_dep_.CoverTab[109374]++

												if isWktPointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:656
		_go_fuzz_dep_.CoverTab[109437]++
													switch t.Kind() {
		case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:658
			_go_fuzz_dep_.CoverTab[109438]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:659
				_go_fuzz_dep_.CoverTab[109466]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:660
					_go_fuzz_dep_.CoverTab[109468]++
																return makeStdDoubleValuePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:661
					// _ = "end of CoverTab[109468]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:662
					_go_fuzz_dep_.CoverTab[109469]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:662
					// _ = "end of CoverTab[109469]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:662
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:662
				// _ = "end of CoverTab[109466]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:662
				_go_fuzz_dep_.CoverTab[109467]++
															return makeStdDoubleValuePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:663
				// _ = "end of CoverTab[109467]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:664
				_go_fuzz_dep_.CoverTab[109470]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:664
				// _ = "end of CoverTab[109470]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:664
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:664
			// _ = "end of CoverTab[109438]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:664
			_go_fuzz_dep_.CoverTab[109439]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:665
				_go_fuzz_dep_.CoverTab[109471]++
															return makeStdDoubleValueSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:666
				// _ = "end of CoverTab[109471]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:667
				_go_fuzz_dep_.CoverTab[109472]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:667
				// _ = "end of CoverTab[109472]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:667
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:667
			// _ = "end of CoverTab[109439]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:667
			_go_fuzz_dep_.CoverTab[109440]++
														return makeStdDoubleValueMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:668
			// _ = "end of CoverTab[109440]"
		case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:669
			_go_fuzz_dep_.CoverTab[109441]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:670
				_go_fuzz_dep_.CoverTab[109473]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:671
					_go_fuzz_dep_.CoverTab[109475]++
																return makeStdFloatValuePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:672
					// _ = "end of CoverTab[109475]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:673
					_go_fuzz_dep_.CoverTab[109476]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:673
					// _ = "end of CoverTab[109476]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:673
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:673
				// _ = "end of CoverTab[109473]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:673
				_go_fuzz_dep_.CoverTab[109474]++
															return makeStdFloatValuePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:674
				// _ = "end of CoverTab[109474]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:675
				_go_fuzz_dep_.CoverTab[109477]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:675
				// _ = "end of CoverTab[109477]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:675
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:675
			// _ = "end of CoverTab[109441]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:675
			_go_fuzz_dep_.CoverTab[109442]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:676
				_go_fuzz_dep_.CoverTab[109478]++
															return makeStdFloatValueSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:677
				// _ = "end of CoverTab[109478]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:678
				_go_fuzz_dep_.CoverTab[109479]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:678
				// _ = "end of CoverTab[109479]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:678
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:678
			// _ = "end of CoverTab[109442]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:678
			_go_fuzz_dep_.CoverTab[109443]++
														return makeStdFloatValueMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:679
			// _ = "end of CoverTab[109443]"
		case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:680
			_go_fuzz_dep_.CoverTab[109444]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:681
				_go_fuzz_dep_.CoverTab[109480]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:682
					_go_fuzz_dep_.CoverTab[109482]++
																return makeStdInt64ValuePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:683
					// _ = "end of CoverTab[109482]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:684
					_go_fuzz_dep_.CoverTab[109483]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:684
					// _ = "end of CoverTab[109483]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:684
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:684
				// _ = "end of CoverTab[109480]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:684
				_go_fuzz_dep_.CoverTab[109481]++
															return makeStdInt64ValuePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:685
				// _ = "end of CoverTab[109481]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:686
				_go_fuzz_dep_.CoverTab[109484]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:686
				// _ = "end of CoverTab[109484]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:686
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:686
			// _ = "end of CoverTab[109444]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:686
			_go_fuzz_dep_.CoverTab[109445]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:687
				_go_fuzz_dep_.CoverTab[109485]++
															return makeStdInt64ValueSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:688
				// _ = "end of CoverTab[109485]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:689
				_go_fuzz_dep_.CoverTab[109486]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:689
				// _ = "end of CoverTab[109486]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:689
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:689
			// _ = "end of CoverTab[109445]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:689
			_go_fuzz_dep_.CoverTab[109446]++
														return makeStdInt64ValueMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:690
			// _ = "end of CoverTab[109446]"
		case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:691
			_go_fuzz_dep_.CoverTab[109447]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:692
				_go_fuzz_dep_.CoverTab[109487]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:693
					_go_fuzz_dep_.CoverTab[109489]++
																return makeStdUInt64ValuePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:694
					// _ = "end of CoverTab[109489]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:695
					_go_fuzz_dep_.CoverTab[109490]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:695
					// _ = "end of CoverTab[109490]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:695
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:695
				// _ = "end of CoverTab[109487]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:695
				_go_fuzz_dep_.CoverTab[109488]++
															return makeStdUInt64ValuePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:696
				// _ = "end of CoverTab[109488]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:697
				_go_fuzz_dep_.CoverTab[109491]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:697
				// _ = "end of CoverTab[109491]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:697
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:697
			// _ = "end of CoverTab[109447]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:697
			_go_fuzz_dep_.CoverTab[109448]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:698
				_go_fuzz_dep_.CoverTab[109492]++
															return makeStdUInt64ValueSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:699
				// _ = "end of CoverTab[109492]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:700
				_go_fuzz_dep_.CoverTab[109493]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:700
				// _ = "end of CoverTab[109493]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:700
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:700
			// _ = "end of CoverTab[109448]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:700
			_go_fuzz_dep_.CoverTab[109449]++
														return makeStdUInt64ValueMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:701
			// _ = "end of CoverTab[109449]"
		case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:702
			_go_fuzz_dep_.CoverTab[109450]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:703
				_go_fuzz_dep_.CoverTab[109494]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:704
					_go_fuzz_dep_.CoverTab[109496]++
																return makeStdInt32ValuePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:705
					// _ = "end of CoverTab[109496]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:706
					_go_fuzz_dep_.CoverTab[109497]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:706
					// _ = "end of CoverTab[109497]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:706
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:706
				// _ = "end of CoverTab[109494]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:706
				_go_fuzz_dep_.CoverTab[109495]++
															return makeStdInt32ValuePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:707
				// _ = "end of CoverTab[109495]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:708
				_go_fuzz_dep_.CoverTab[109498]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:708
				// _ = "end of CoverTab[109498]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:708
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:708
			// _ = "end of CoverTab[109450]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:708
			_go_fuzz_dep_.CoverTab[109451]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:709
				_go_fuzz_dep_.CoverTab[109499]++
															return makeStdInt32ValueSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:710
				// _ = "end of CoverTab[109499]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:711
				_go_fuzz_dep_.CoverTab[109500]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:711
				// _ = "end of CoverTab[109500]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:711
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:711
			// _ = "end of CoverTab[109451]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:711
			_go_fuzz_dep_.CoverTab[109452]++
														return makeStdInt32ValueMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:712
			// _ = "end of CoverTab[109452]"
		case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:713
			_go_fuzz_dep_.CoverTab[109453]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:714
				_go_fuzz_dep_.CoverTab[109501]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:715
					_go_fuzz_dep_.CoverTab[109503]++
																return makeStdUInt32ValuePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:716
					// _ = "end of CoverTab[109503]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:717
					_go_fuzz_dep_.CoverTab[109504]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:717
					// _ = "end of CoverTab[109504]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:717
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:717
				// _ = "end of CoverTab[109501]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:717
				_go_fuzz_dep_.CoverTab[109502]++
															return makeStdUInt32ValuePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:718
				// _ = "end of CoverTab[109502]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:719
				_go_fuzz_dep_.CoverTab[109505]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:719
				// _ = "end of CoverTab[109505]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:719
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:719
			// _ = "end of CoverTab[109453]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:719
			_go_fuzz_dep_.CoverTab[109454]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:720
				_go_fuzz_dep_.CoverTab[109506]++
															return makeStdUInt32ValueSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:721
				// _ = "end of CoverTab[109506]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:722
				_go_fuzz_dep_.CoverTab[109507]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:722
				// _ = "end of CoverTab[109507]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:722
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:722
			// _ = "end of CoverTab[109454]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:722
			_go_fuzz_dep_.CoverTab[109455]++
														return makeStdUInt32ValueMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:723
			// _ = "end of CoverTab[109455]"
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:724
			_go_fuzz_dep_.CoverTab[109456]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:725
				_go_fuzz_dep_.CoverTab[109508]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:726
					_go_fuzz_dep_.CoverTab[109510]++
																return makeStdBoolValuePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:727
					// _ = "end of CoverTab[109510]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:728
					_go_fuzz_dep_.CoverTab[109511]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:728
					// _ = "end of CoverTab[109511]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:728
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:728
				// _ = "end of CoverTab[109508]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:728
				_go_fuzz_dep_.CoverTab[109509]++
															return makeStdBoolValuePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:729
				// _ = "end of CoverTab[109509]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:730
				_go_fuzz_dep_.CoverTab[109512]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:730
				// _ = "end of CoverTab[109512]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:730
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:730
			// _ = "end of CoverTab[109456]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:730
			_go_fuzz_dep_.CoverTab[109457]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:731
				_go_fuzz_dep_.CoverTab[109513]++
															return makeStdBoolValueSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:732
				// _ = "end of CoverTab[109513]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:733
				_go_fuzz_dep_.CoverTab[109514]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:733
				// _ = "end of CoverTab[109514]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:733
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:733
			// _ = "end of CoverTab[109457]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:733
			_go_fuzz_dep_.CoverTab[109458]++
														return makeStdBoolValueMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:734
			// _ = "end of CoverTab[109458]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:735
			_go_fuzz_dep_.CoverTab[109459]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:736
				_go_fuzz_dep_.CoverTab[109515]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:737
					_go_fuzz_dep_.CoverTab[109517]++
																return makeStdStringValuePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:738
					// _ = "end of CoverTab[109517]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:739
					_go_fuzz_dep_.CoverTab[109518]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:739
					// _ = "end of CoverTab[109518]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:739
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:739
				// _ = "end of CoverTab[109515]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:739
				_go_fuzz_dep_.CoverTab[109516]++
															return makeStdStringValuePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:740
				// _ = "end of CoverTab[109516]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:741
				_go_fuzz_dep_.CoverTab[109519]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:741
				// _ = "end of CoverTab[109519]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:741
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:741
			// _ = "end of CoverTab[109459]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:741
			_go_fuzz_dep_.CoverTab[109460]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:742
				_go_fuzz_dep_.CoverTab[109520]++
															return makeStdStringValueSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:743
				// _ = "end of CoverTab[109520]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:744
				_go_fuzz_dep_.CoverTab[109521]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:744
				// _ = "end of CoverTab[109521]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:744
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:744
			// _ = "end of CoverTab[109460]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:744
			_go_fuzz_dep_.CoverTab[109461]++
														return makeStdStringValueMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:745
			// _ = "end of CoverTab[109461]"
		case uint8SliceType:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:746
			_go_fuzz_dep_.CoverTab[109462]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:747
				_go_fuzz_dep_.CoverTab[109522]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:748
					_go_fuzz_dep_.CoverTab[109524]++
																return makeStdBytesValuePtrSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:749
					// _ = "end of CoverTab[109524]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:750
					_go_fuzz_dep_.CoverTab[109525]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:750
					// _ = "end of CoverTab[109525]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:750
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:750
				// _ = "end of CoverTab[109522]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:750
				_go_fuzz_dep_.CoverTab[109523]++
															return makeStdBytesValuePtrMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:751
				// _ = "end of CoverTab[109523]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:752
				_go_fuzz_dep_.CoverTab[109526]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:752
				// _ = "end of CoverTab[109526]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:752
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:752
			// _ = "end of CoverTab[109462]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:752
			_go_fuzz_dep_.CoverTab[109463]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:753
				_go_fuzz_dep_.CoverTab[109527]++
															return makeStdBytesValueSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:754
				// _ = "end of CoverTab[109527]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:755
				_go_fuzz_dep_.CoverTab[109528]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:755
				// _ = "end of CoverTab[109528]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:755
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:755
			// _ = "end of CoverTab[109463]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:755
			_go_fuzz_dep_.CoverTab[109464]++
														return makeStdBytesValueMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:756
			// _ = "end of CoverTab[109464]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:757
			_go_fuzz_dep_.CoverTab[109465]++
														panic(fmt.Sprintf("unknown wktpointer type %#v", t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:758
			// _ = "end of CoverTab[109465]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:759
		// _ = "end of CoverTab[109437]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:760
		_go_fuzz_dep_.CoverTab[109529]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:760
		// _ = "end of CoverTab[109529]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:760
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:760
	// _ = "end of CoverTab[109374]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:760
	_go_fuzz_dep_.CoverTab[109375]++

												switch t.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:763
		_go_fuzz_dep_.CoverTab[109530]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:764
			_go_fuzz_dep_.CoverTab[109557]++
														return sizeBoolPtr, appendBoolPtr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:765
			// _ = "end of CoverTab[109557]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:766
			_go_fuzz_dep_.CoverTab[109558]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:766
			// _ = "end of CoverTab[109558]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:766
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:766
		// _ = "end of CoverTab[109530]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:766
		_go_fuzz_dep_.CoverTab[109531]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:767
			_go_fuzz_dep_.CoverTab[109559]++
														if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:768
				_go_fuzz_dep_.CoverTab[109561]++
															return sizeBoolPackedSlice, appendBoolPackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:769
				// _ = "end of CoverTab[109561]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:770
				_go_fuzz_dep_.CoverTab[109562]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:770
				// _ = "end of CoverTab[109562]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:770
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:770
			// _ = "end of CoverTab[109559]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:770
			_go_fuzz_dep_.CoverTab[109560]++
														return sizeBoolSlice, appendBoolSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:771
			// _ = "end of CoverTab[109560]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:772
			_go_fuzz_dep_.CoverTab[109563]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:772
			// _ = "end of CoverTab[109563]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:772
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:772
		// _ = "end of CoverTab[109531]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:772
		_go_fuzz_dep_.CoverTab[109532]++
													if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:773
			_go_fuzz_dep_.CoverTab[109564]++
														return sizeBoolValueNoZero, appendBoolValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:774
			// _ = "end of CoverTab[109564]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:775
			_go_fuzz_dep_.CoverTab[109565]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:775
			// _ = "end of CoverTab[109565]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:775
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:775
		// _ = "end of CoverTab[109532]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:775
		_go_fuzz_dep_.CoverTab[109533]++
													return sizeBoolValue, appendBoolValue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:776
		// _ = "end of CoverTab[109533]"
	case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:777
		_go_fuzz_dep_.CoverTab[109534]++
													switch encoding {
		case "fixed32":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:779
			_go_fuzz_dep_.CoverTab[109566]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:780
				_go_fuzz_dep_.CoverTab[109575]++
															return sizeFixed32Ptr, appendFixed32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:781
				// _ = "end of CoverTab[109575]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:782
				_go_fuzz_dep_.CoverTab[109576]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:782
				// _ = "end of CoverTab[109576]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:782
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:782
			// _ = "end of CoverTab[109566]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:782
			_go_fuzz_dep_.CoverTab[109567]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:783
				_go_fuzz_dep_.CoverTab[109577]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:784
					_go_fuzz_dep_.CoverTab[109579]++
																return sizeFixed32PackedSlice, appendFixed32PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:785
					// _ = "end of CoverTab[109579]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:786
					_go_fuzz_dep_.CoverTab[109580]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:786
					// _ = "end of CoverTab[109580]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:786
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:786
				// _ = "end of CoverTab[109577]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:786
				_go_fuzz_dep_.CoverTab[109578]++
															return sizeFixed32Slice, appendFixed32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:787
				// _ = "end of CoverTab[109578]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:788
				_go_fuzz_dep_.CoverTab[109581]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:788
				// _ = "end of CoverTab[109581]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:788
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:788
			// _ = "end of CoverTab[109567]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:788
			_go_fuzz_dep_.CoverTab[109568]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:789
				_go_fuzz_dep_.CoverTab[109582]++
															return sizeFixed32ValueNoZero, appendFixed32ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:790
				// _ = "end of CoverTab[109582]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:791
				_go_fuzz_dep_.CoverTab[109583]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:791
				// _ = "end of CoverTab[109583]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:791
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:791
			// _ = "end of CoverTab[109568]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:791
			_go_fuzz_dep_.CoverTab[109569]++
														return sizeFixed32Value, appendFixed32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:792
			// _ = "end of CoverTab[109569]"
		case "varint":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:793
			_go_fuzz_dep_.CoverTab[109570]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:794
				_go_fuzz_dep_.CoverTab[109584]++
															return sizeVarint32Ptr, appendVarint32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:795
				// _ = "end of CoverTab[109584]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:796
				_go_fuzz_dep_.CoverTab[109585]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:796
				// _ = "end of CoverTab[109585]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:796
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:796
			// _ = "end of CoverTab[109570]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:796
			_go_fuzz_dep_.CoverTab[109571]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:797
				_go_fuzz_dep_.CoverTab[109586]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:798
					_go_fuzz_dep_.CoverTab[109588]++
																return sizeVarint32PackedSlice, appendVarint32PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:799
					// _ = "end of CoverTab[109588]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:800
					_go_fuzz_dep_.CoverTab[109589]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:800
					// _ = "end of CoverTab[109589]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:800
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:800
				// _ = "end of CoverTab[109586]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:800
				_go_fuzz_dep_.CoverTab[109587]++
															return sizeVarint32Slice, appendVarint32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:801
				// _ = "end of CoverTab[109587]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:802
				_go_fuzz_dep_.CoverTab[109590]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:802
				// _ = "end of CoverTab[109590]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:802
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:802
			// _ = "end of CoverTab[109571]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:802
			_go_fuzz_dep_.CoverTab[109572]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:803
				_go_fuzz_dep_.CoverTab[109591]++
															return sizeVarint32ValueNoZero, appendVarint32ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:804
				// _ = "end of CoverTab[109591]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:805
				_go_fuzz_dep_.CoverTab[109592]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:805
				// _ = "end of CoverTab[109592]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:805
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:805
			// _ = "end of CoverTab[109572]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:805
			_go_fuzz_dep_.CoverTab[109573]++
														return sizeVarint32Value, appendVarint32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:806
			// _ = "end of CoverTab[109573]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:806
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:806
			_go_fuzz_dep_.CoverTab[109574]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:806
			// _ = "end of CoverTab[109574]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:807
		// _ = "end of CoverTab[109534]"
	case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:808
		_go_fuzz_dep_.CoverTab[109535]++
													switch encoding {
		case "fixed32":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:810
			_go_fuzz_dep_.CoverTab[109593]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:811
				_go_fuzz_dep_.CoverTab[109606]++
															return sizeFixedS32Ptr, appendFixedS32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:812
				// _ = "end of CoverTab[109606]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:813
				_go_fuzz_dep_.CoverTab[109607]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:813
				// _ = "end of CoverTab[109607]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:813
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:813
			// _ = "end of CoverTab[109593]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:813
			_go_fuzz_dep_.CoverTab[109594]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:814
				_go_fuzz_dep_.CoverTab[109608]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:815
					_go_fuzz_dep_.CoverTab[109610]++
																return sizeFixedS32PackedSlice, appendFixedS32PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:816
					// _ = "end of CoverTab[109610]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:817
					_go_fuzz_dep_.CoverTab[109611]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:817
					// _ = "end of CoverTab[109611]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:817
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:817
				// _ = "end of CoverTab[109608]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:817
				_go_fuzz_dep_.CoverTab[109609]++
															return sizeFixedS32Slice, appendFixedS32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:818
				// _ = "end of CoverTab[109609]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:819
				_go_fuzz_dep_.CoverTab[109612]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:819
				// _ = "end of CoverTab[109612]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:819
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:819
			// _ = "end of CoverTab[109594]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:819
			_go_fuzz_dep_.CoverTab[109595]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:820
				_go_fuzz_dep_.CoverTab[109613]++
															return sizeFixedS32ValueNoZero, appendFixedS32ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:821
				// _ = "end of CoverTab[109613]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:822
				_go_fuzz_dep_.CoverTab[109614]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:822
				// _ = "end of CoverTab[109614]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:822
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:822
			// _ = "end of CoverTab[109595]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:822
			_go_fuzz_dep_.CoverTab[109596]++
														return sizeFixedS32Value, appendFixedS32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:823
			// _ = "end of CoverTab[109596]"
		case "varint":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:824
			_go_fuzz_dep_.CoverTab[109597]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:825
				_go_fuzz_dep_.CoverTab[109615]++
															return sizeVarintS32Ptr, appendVarintS32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:826
				// _ = "end of CoverTab[109615]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:827
				_go_fuzz_dep_.CoverTab[109616]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:827
				// _ = "end of CoverTab[109616]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:827
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:827
			// _ = "end of CoverTab[109597]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:827
			_go_fuzz_dep_.CoverTab[109598]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:828
				_go_fuzz_dep_.CoverTab[109617]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:829
					_go_fuzz_dep_.CoverTab[109619]++
																return sizeVarintS32PackedSlice, appendVarintS32PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:830
					// _ = "end of CoverTab[109619]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:831
					_go_fuzz_dep_.CoverTab[109620]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:831
					// _ = "end of CoverTab[109620]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:831
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:831
				// _ = "end of CoverTab[109617]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:831
				_go_fuzz_dep_.CoverTab[109618]++
															return sizeVarintS32Slice, appendVarintS32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:832
				// _ = "end of CoverTab[109618]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:833
				_go_fuzz_dep_.CoverTab[109621]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:833
				// _ = "end of CoverTab[109621]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:833
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:833
			// _ = "end of CoverTab[109598]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:833
			_go_fuzz_dep_.CoverTab[109599]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:834
				_go_fuzz_dep_.CoverTab[109622]++
															return sizeVarintS32ValueNoZero, appendVarintS32ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:835
				// _ = "end of CoverTab[109622]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:836
				_go_fuzz_dep_.CoverTab[109623]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:836
				// _ = "end of CoverTab[109623]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:836
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:836
			// _ = "end of CoverTab[109599]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:836
			_go_fuzz_dep_.CoverTab[109600]++
														return sizeVarintS32Value, appendVarintS32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:837
			// _ = "end of CoverTab[109600]"
		case "zigzag32":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:838
			_go_fuzz_dep_.CoverTab[109601]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:839
				_go_fuzz_dep_.CoverTab[109624]++
															return sizeZigzag32Ptr, appendZigzag32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:840
				// _ = "end of CoverTab[109624]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:841
				_go_fuzz_dep_.CoverTab[109625]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:841
				// _ = "end of CoverTab[109625]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:841
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:841
			// _ = "end of CoverTab[109601]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:841
			_go_fuzz_dep_.CoverTab[109602]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:842
				_go_fuzz_dep_.CoverTab[109626]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:843
					_go_fuzz_dep_.CoverTab[109628]++
																return sizeZigzag32PackedSlice, appendZigzag32PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:844
					// _ = "end of CoverTab[109628]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:845
					_go_fuzz_dep_.CoverTab[109629]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:845
					// _ = "end of CoverTab[109629]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:845
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:845
				// _ = "end of CoverTab[109626]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:845
				_go_fuzz_dep_.CoverTab[109627]++
															return sizeZigzag32Slice, appendZigzag32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:846
				// _ = "end of CoverTab[109627]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:847
				_go_fuzz_dep_.CoverTab[109630]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:847
				// _ = "end of CoverTab[109630]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:847
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:847
			// _ = "end of CoverTab[109602]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:847
			_go_fuzz_dep_.CoverTab[109603]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:848
				_go_fuzz_dep_.CoverTab[109631]++
															return sizeZigzag32ValueNoZero, appendZigzag32ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:849
				// _ = "end of CoverTab[109631]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:850
				_go_fuzz_dep_.CoverTab[109632]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:850
				// _ = "end of CoverTab[109632]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:850
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:850
			// _ = "end of CoverTab[109603]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:850
			_go_fuzz_dep_.CoverTab[109604]++
														return sizeZigzag32Value, appendZigzag32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:851
			// _ = "end of CoverTab[109604]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:851
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:851
			_go_fuzz_dep_.CoverTab[109605]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:851
			// _ = "end of CoverTab[109605]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:852
		// _ = "end of CoverTab[109535]"
	case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:853
		_go_fuzz_dep_.CoverTab[109536]++
													switch encoding {
		case "fixed64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:855
			_go_fuzz_dep_.CoverTab[109633]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:856
				_go_fuzz_dep_.CoverTab[109642]++
															return sizeFixed64Ptr, appendFixed64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:857
				// _ = "end of CoverTab[109642]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:858
				_go_fuzz_dep_.CoverTab[109643]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:858
				// _ = "end of CoverTab[109643]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:858
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:858
			// _ = "end of CoverTab[109633]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:858
			_go_fuzz_dep_.CoverTab[109634]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:859
				_go_fuzz_dep_.CoverTab[109644]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:860
					_go_fuzz_dep_.CoverTab[109646]++
																return sizeFixed64PackedSlice, appendFixed64PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:861
					// _ = "end of CoverTab[109646]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:862
					_go_fuzz_dep_.CoverTab[109647]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:862
					// _ = "end of CoverTab[109647]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:862
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:862
				// _ = "end of CoverTab[109644]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:862
				_go_fuzz_dep_.CoverTab[109645]++
															return sizeFixed64Slice, appendFixed64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:863
				// _ = "end of CoverTab[109645]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:864
				_go_fuzz_dep_.CoverTab[109648]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:864
				// _ = "end of CoverTab[109648]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:864
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:864
			// _ = "end of CoverTab[109634]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:864
			_go_fuzz_dep_.CoverTab[109635]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:865
				_go_fuzz_dep_.CoverTab[109649]++
															return sizeFixed64ValueNoZero, appendFixed64ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:866
				// _ = "end of CoverTab[109649]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:867
				_go_fuzz_dep_.CoverTab[109650]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:867
				// _ = "end of CoverTab[109650]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:867
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:867
			// _ = "end of CoverTab[109635]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:867
			_go_fuzz_dep_.CoverTab[109636]++
														return sizeFixed64Value, appendFixed64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:868
			// _ = "end of CoverTab[109636]"
		case "varint":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:869
			_go_fuzz_dep_.CoverTab[109637]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:870
				_go_fuzz_dep_.CoverTab[109651]++
															return sizeVarint64Ptr, appendVarint64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:871
				// _ = "end of CoverTab[109651]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:872
				_go_fuzz_dep_.CoverTab[109652]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:872
				// _ = "end of CoverTab[109652]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:872
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:872
			// _ = "end of CoverTab[109637]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:872
			_go_fuzz_dep_.CoverTab[109638]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:873
				_go_fuzz_dep_.CoverTab[109653]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:874
					_go_fuzz_dep_.CoverTab[109655]++
																return sizeVarint64PackedSlice, appendVarint64PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:875
					// _ = "end of CoverTab[109655]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:876
					_go_fuzz_dep_.CoverTab[109656]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:876
					// _ = "end of CoverTab[109656]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:876
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:876
				// _ = "end of CoverTab[109653]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:876
				_go_fuzz_dep_.CoverTab[109654]++
															return sizeVarint64Slice, appendVarint64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:877
				// _ = "end of CoverTab[109654]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:878
				_go_fuzz_dep_.CoverTab[109657]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:878
				// _ = "end of CoverTab[109657]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:878
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:878
			// _ = "end of CoverTab[109638]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:878
			_go_fuzz_dep_.CoverTab[109639]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:879
				_go_fuzz_dep_.CoverTab[109658]++
															return sizeVarint64ValueNoZero, appendVarint64ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:880
				// _ = "end of CoverTab[109658]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:881
				_go_fuzz_dep_.CoverTab[109659]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:881
				// _ = "end of CoverTab[109659]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:881
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:881
			// _ = "end of CoverTab[109639]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:881
			_go_fuzz_dep_.CoverTab[109640]++
														return sizeVarint64Value, appendVarint64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:882
			// _ = "end of CoverTab[109640]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:882
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:882
			_go_fuzz_dep_.CoverTab[109641]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:882
			// _ = "end of CoverTab[109641]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:883
		// _ = "end of CoverTab[109536]"
	case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:884
		_go_fuzz_dep_.CoverTab[109537]++
													switch encoding {
		case "fixed64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:886
			_go_fuzz_dep_.CoverTab[109660]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:887
				_go_fuzz_dep_.CoverTab[109673]++
															return sizeFixedS64Ptr, appendFixedS64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:888
				// _ = "end of CoverTab[109673]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:889
				_go_fuzz_dep_.CoverTab[109674]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:889
				// _ = "end of CoverTab[109674]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:889
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:889
			// _ = "end of CoverTab[109660]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:889
			_go_fuzz_dep_.CoverTab[109661]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:890
				_go_fuzz_dep_.CoverTab[109675]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:891
					_go_fuzz_dep_.CoverTab[109677]++
																return sizeFixedS64PackedSlice, appendFixedS64PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:892
					// _ = "end of CoverTab[109677]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:893
					_go_fuzz_dep_.CoverTab[109678]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:893
					// _ = "end of CoverTab[109678]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:893
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:893
				// _ = "end of CoverTab[109675]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:893
				_go_fuzz_dep_.CoverTab[109676]++
															return sizeFixedS64Slice, appendFixedS64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:894
				// _ = "end of CoverTab[109676]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:895
				_go_fuzz_dep_.CoverTab[109679]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:895
				// _ = "end of CoverTab[109679]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:895
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:895
			// _ = "end of CoverTab[109661]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:895
			_go_fuzz_dep_.CoverTab[109662]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:896
				_go_fuzz_dep_.CoverTab[109680]++
															return sizeFixedS64ValueNoZero, appendFixedS64ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:897
				// _ = "end of CoverTab[109680]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:898
				_go_fuzz_dep_.CoverTab[109681]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:898
				// _ = "end of CoverTab[109681]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:898
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:898
			// _ = "end of CoverTab[109662]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:898
			_go_fuzz_dep_.CoverTab[109663]++
														return sizeFixedS64Value, appendFixedS64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:899
			// _ = "end of CoverTab[109663]"
		case "varint":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:900
			_go_fuzz_dep_.CoverTab[109664]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:901
				_go_fuzz_dep_.CoverTab[109682]++
															return sizeVarintS64Ptr, appendVarintS64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:902
				// _ = "end of CoverTab[109682]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:903
				_go_fuzz_dep_.CoverTab[109683]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:903
				// _ = "end of CoverTab[109683]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:903
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:903
			// _ = "end of CoverTab[109664]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:903
			_go_fuzz_dep_.CoverTab[109665]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:904
				_go_fuzz_dep_.CoverTab[109684]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:905
					_go_fuzz_dep_.CoverTab[109686]++
																return sizeVarintS64PackedSlice, appendVarintS64PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:906
					// _ = "end of CoverTab[109686]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:907
					_go_fuzz_dep_.CoverTab[109687]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:907
					// _ = "end of CoverTab[109687]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:907
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:907
				// _ = "end of CoverTab[109684]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:907
				_go_fuzz_dep_.CoverTab[109685]++
															return sizeVarintS64Slice, appendVarintS64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:908
				// _ = "end of CoverTab[109685]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:909
				_go_fuzz_dep_.CoverTab[109688]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:909
				// _ = "end of CoverTab[109688]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:909
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:909
			// _ = "end of CoverTab[109665]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:909
			_go_fuzz_dep_.CoverTab[109666]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:910
				_go_fuzz_dep_.CoverTab[109689]++
															return sizeVarintS64ValueNoZero, appendVarintS64ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:911
				// _ = "end of CoverTab[109689]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:912
				_go_fuzz_dep_.CoverTab[109690]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:912
				// _ = "end of CoverTab[109690]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:912
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:912
			// _ = "end of CoverTab[109666]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:912
			_go_fuzz_dep_.CoverTab[109667]++
														return sizeVarintS64Value, appendVarintS64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:913
			// _ = "end of CoverTab[109667]"
		case "zigzag64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:914
			_go_fuzz_dep_.CoverTab[109668]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:915
				_go_fuzz_dep_.CoverTab[109691]++
															return sizeZigzag64Ptr, appendZigzag64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:916
				// _ = "end of CoverTab[109691]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:917
				_go_fuzz_dep_.CoverTab[109692]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:917
				// _ = "end of CoverTab[109692]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:917
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:917
			// _ = "end of CoverTab[109668]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:917
			_go_fuzz_dep_.CoverTab[109669]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:918
				_go_fuzz_dep_.CoverTab[109693]++
															if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:919
					_go_fuzz_dep_.CoverTab[109695]++
																return sizeZigzag64PackedSlice, appendZigzag64PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:920
					// _ = "end of CoverTab[109695]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:921
					_go_fuzz_dep_.CoverTab[109696]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:921
					// _ = "end of CoverTab[109696]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:921
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:921
				// _ = "end of CoverTab[109693]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:921
				_go_fuzz_dep_.CoverTab[109694]++
															return sizeZigzag64Slice, appendZigzag64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:922
				// _ = "end of CoverTab[109694]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:923
				_go_fuzz_dep_.CoverTab[109697]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:923
				// _ = "end of CoverTab[109697]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:923
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:923
			// _ = "end of CoverTab[109669]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:923
			_go_fuzz_dep_.CoverTab[109670]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:924
				_go_fuzz_dep_.CoverTab[109698]++
															return sizeZigzag64ValueNoZero, appendZigzag64ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:925
				// _ = "end of CoverTab[109698]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:926
				_go_fuzz_dep_.CoverTab[109699]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:926
				// _ = "end of CoverTab[109699]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:926
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:926
			// _ = "end of CoverTab[109670]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:926
			_go_fuzz_dep_.CoverTab[109671]++
														return sizeZigzag64Value, appendZigzag64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:927
			// _ = "end of CoverTab[109671]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:927
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:927
			_go_fuzz_dep_.CoverTab[109672]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:927
			// _ = "end of CoverTab[109672]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:928
		// _ = "end of CoverTab[109537]"
	case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:929
		_go_fuzz_dep_.CoverTab[109538]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:930
			_go_fuzz_dep_.CoverTab[109700]++
														return sizeFloat32Ptr, appendFloat32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:931
			// _ = "end of CoverTab[109700]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:932
			_go_fuzz_dep_.CoverTab[109701]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:932
			// _ = "end of CoverTab[109701]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:932
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:932
		// _ = "end of CoverTab[109538]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:932
		_go_fuzz_dep_.CoverTab[109539]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:933
			_go_fuzz_dep_.CoverTab[109702]++
														if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:934
				_go_fuzz_dep_.CoverTab[109704]++
															return sizeFloat32PackedSlice, appendFloat32PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:935
				// _ = "end of CoverTab[109704]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:936
				_go_fuzz_dep_.CoverTab[109705]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:936
				// _ = "end of CoverTab[109705]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:936
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:936
			// _ = "end of CoverTab[109702]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:936
			_go_fuzz_dep_.CoverTab[109703]++
														return sizeFloat32Slice, appendFloat32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:937
			// _ = "end of CoverTab[109703]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:938
			_go_fuzz_dep_.CoverTab[109706]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:938
			// _ = "end of CoverTab[109706]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:938
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:938
		// _ = "end of CoverTab[109539]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:938
		_go_fuzz_dep_.CoverTab[109540]++
													if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:939
			_go_fuzz_dep_.CoverTab[109707]++
														return sizeFloat32ValueNoZero, appendFloat32ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:940
			// _ = "end of CoverTab[109707]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:941
			_go_fuzz_dep_.CoverTab[109708]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:941
			// _ = "end of CoverTab[109708]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:941
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:941
		// _ = "end of CoverTab[109540]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:941
		_go_fuzz_dep_.CoverTab[109541]++
													return sizeFloat32Value, appendFloat32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:942
		// _ = "end of CoverTab[109541]"
	case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:943
		_go_fuzz_dep_.CoverTab[109542]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:944
			_go_fuzz_dep_.CoverTab[109709]++
														return sizeFloat64Ptr, appendFloat64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:945
			// _ = "end of CoverTab[109709]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:946
			_go_fuzz_dep_.CoverTab[109710]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:946
			// _ = "end of CoverTab[109710]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:946
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:946
		// _ = "end of CoverTab[109542]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:946
		_go_fuzz_dep_.CoverTab[109543]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:947
			_go_fuzz_dep_.CoverTab[109711]++
														if packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:948
				_go_fuzz_dep_.CoverTab[109713]++
															return sizeFloat64PackedSlice, appendFloat64PackedSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:949
				// _ = "end of CoverTab[109713]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:950
				_go_fuzz_dep_.CoverTab[109714]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:950
				// _ = "end of CoverTab[109714]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:950
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:950
			// _ = "end of CoverTab[109711]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:950
			_go_fuzz_dep_.CoverTab[109712]++
														return sizeFloat64Slice, appendFloat64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:951
			// _ = "end of CoverTab[109712]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:952
			_go_fuzz_dep_.CoverTab[109715]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:952
			// _ = "end of CoverTab[109715]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:952
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:952
		// _ = "end of CoverTab[109543]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:952
		_go_fuzz_dep_.CoverTab[109544]++
													if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:953
			_go_fuzz_dep_.CoverTab[109716]++
														return sizeFloat64ValueNoZero, appendFloat64ValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:954
			// _ = "end of CoverTab[109716]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:955
			_go_fuzz_dep_.CoverTab[109717]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:955
			// _ = "end of CoverTab[109717]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:955
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:955
		// _ = "end of CoverTab[109544]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:955
		_go_fuzz_dep_.CoverTab[109545]++
													return sizeFloat64Value, appendFloat64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:956
		// _ = "end of CoverTab[109545]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:957
		_go_fuzz_dep_.CoverTab[109546]++
													if validateUTF8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:958
			_go_fuzz_dep_.CoverTab[109718]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:959
				_go_fuzz_dep_.CoverTab[109722]++
															return sizeStringPtr, appendUTF8StringPtr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:960
				// _ = "end of CoverTab[109722]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:961
				_go_fuzz_dep_.CoverTab[109723]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:961
				// _ = "end of CoverTab[109723]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:961
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:961
			// _ = "end of CoverTab[109718]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:961
			_go_fuzz_dep_.CoverTab[109719]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:962
				_go_fuzz_dep_.CoverTab[109724]++
															return sizeStringSlice, appendUTF8StringSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:963
				// _ = "end of CoverTab[109724]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:964
				_go_fuzz_dep_.CoverTab[109725]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:964
				// _ = "end of CoverTab[109725]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:964
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:964
			// _ = "end of CoverTab[109719]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:964
			_go_fuzz_dep_.CoverTab[109720]++
														if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:965
				_go_fuzz_dep_.CoverTab[109726]++
															return sizeStringValueNoZero, appendUTF8StringValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:966
				// _ = "end of CoverTab[109726]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:967
				_go_fuzz_dep_.CoverTab[109727]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:967
				// _ = "end of CoverTab[109727]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:967
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:967
			// _ = "end of CoverTab[109720]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:967
			_go_fuzz_dep_.CoverTab[109721]++
														return sizeStringValue, appendUTF8StringValue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:968
			// _ = "end of CoverTab[109721]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:969
			_go_fuzz_dep_.CoverTab[109728]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:969
			// _ = "end of CoverTab[109728]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:969
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:969
		// _ = "end of CoverTab[109546]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:969
		_go_fuzz_dep_.CoverTab[109547]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:970
			_go_fuzz_dep_.CoverTab[109729]++
														return sizeStringPtr, appendStringPtr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:971
			// _ = "end of CoverTab[109729]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:972
			_go_fuzz_dep_.CoverTab[109730]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:972
			// _ = "end of CoverTab[109730]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:972
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:972
		// _ = "end of CoverTab[109547]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:972
		_go_fuzz_dep_.CoverTab[109548]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:973
			_go_fuzz_dep_.CoverTab[109731]++
														return sizeStringSlice, appendStringSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:974
			// _ = "end of CoverTab[109731]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:975
			_go_fuzz_dep_.CoverTab[109732]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:975
			// _ = "end of CoverTab[109732]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:975
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:975
		// _ = "end of CoverTab[109548]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:975
		_go_fuzz_dep_.CoverTab[109549]++
													if nozero {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:976
			_go_fuzz_dep_.CoverTab[109733]++
														return sizeStringValueNoZero, appendStringValueNoZero
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:977
			// _ = "end of CoverTab[109733]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:978
			_go_fuzz_dep_.CoverTab[109734]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:978
			// _ = "end of CoverTab[109734]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:978
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:978
		// _ = "end of CoverTab[109549]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:978
		_go_fuzz_dep_.CoverTab[109550]++
													return sizeStringValue, appendStringValue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:979
		// _ = "end of CoverTab[109550]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:980
		_go_fuzz_dep_.CoverTab[109551]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:981
			_go_fuzz_dep_.CoverTab[109735]++
														return sizeBytesSlice, appendBytesSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:982
			// _ = "end of CoverTab[109735]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:983
			_go_fuzz_dep_.CoverTab[109736]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:983
			// _ = "end of CoverTab[109736]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:983
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:983
		// _ = "end of CoverTab[109551]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:983
		_go_fuzz_dep_.CoverTab[109552]++
													if oneof {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:984
			_go_fuzz_dep_.CoverTab[109737]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:988
			return sizeBytesOneof, appendBytesOneof
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:988
			// _ = "end of CoverTab[109737]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:989
			_go_fuzz_dep_.CoverTab[109738]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:989
			// _ = "end of CoverTab[109738]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:989
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:989
		// _ = "end of CoverTab[109552]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:989
		_go_fuzz_dep_.CoverTab[109553]++
													if proto3 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:990
			_go_fuzz_dep_.CoverTab[109739]++
														return sizeBytes3, appendBytes3
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:991
			// _ = "end of CoverTab[109739]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:992
			_go_fuzz_dep_.CoverTab[109740]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:992
			// _ = "end of CoverTab[109740]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:992
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:992
		// _ = "end of CoverTab[109553]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:992
		_go_fuzz_dep_.CoverTab[109554]++
													return sizeBytes, appendBytes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:993
		// _ = "end of CoverTab[109554]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:994
		_go_fuzz_dep_.CoverTab[109555]++
													switch encoding {
		case "group":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:996
			_go_fuzz_dep_.CoverTab[109741]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:997
				_go_fuzz_dep_.CoverTab[109745]++
															return makeGroupSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:998
				// _ = "end of CoverTab[109745]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:999
				_go_fuzz_dep_.CoverTab[109746]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:999
				// _ = "end of CoverTab[109746]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:999
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:999
			// _ = "end of CoverTab[109741]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:999
			_go_fuzz_dep_.CoverTab[109742]++
														return makeGroupMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1000
			// _ = "end of CoverTab[109742]"
		case "bytes":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1001
			_go_fuzz_dep_.CoverTab[109743]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1002
				_go_fuzz_dep_.CoverTab[109747]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1003
					_go_fuzz_dep_.CoverTab[109749]++
																return makeMessageSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1004
					// _ = "end of CoverTab[109749]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1005
					_go_fuzz_dep_.CoverTab[109750]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1005
					// _ = "end of CoverTab[109750]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1005
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1005
				// _ = "end of CoverTab[109747]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1005
				_go_fuzz_dep_.CoverTab[109748]++
															return makeMessageMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1006
				// _ = "end of CoverTab[109748]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1007
				_go_fuzz_dep_.CoverTab[109751]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1008
					_go_fuzz_dep_.CoverTab[109753]++
																return makeMessageRefSliceMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1009
					// _ = "end of CoverTab[109753]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1010
					_go_fuzz_dep_.CoverTab[109754]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1010
					// _ = "end of CoverTab[109754]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1010
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1010
				// _ = "end of CoverTab[109751]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1010
				_go_fuzz_dep_.CoverTab[109752]++
															return makeMessageRefMarshaler(getMarshalInfo(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1011
				// _ = "end of CoverTab[109752]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1012
			// _ = "end of CoverTab[109743]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1012
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1012
			_go_fuzz_dep_.CoverTab[109744]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1012
			// _ = "end of CoverTab[109744]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1013
		// _ = "end of CoverTab[109555]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1013
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1013
		_go_fuzz_dep_.CoverTab[109556]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1013
		// _ = "end of CoverTab[109556]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1014
	// _ = "end of CoverTab[109375]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1014
	_go_fuzz_dep_.CoverTab[109376]++
												panic(fmt.Sprintf("unknown or mismatched type: type: %v, wire type: %v", t, encoding))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1015
	// _ = "end of CoverTab[109376]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1022
func sizeFixed32Value(_ pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1022
	_go_fuzz_dep_.CoverTab[109755]++
												return 4 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1023
	// _ = "end of CoverTab[109755]"
}
func sizeFixed32ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1025
	_go_fuzz_dep_.CoverTab[109756]++
												v := *ptr.toUint32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1027
		_go_fuzz_dep_.CoverTab[109758]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1028
		// _ = "end of CoverTab[109758]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1029
		_go_fuzz_dep_.CoverTab[109759]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1029
		// _ = "end of CoverTab[109759]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1029
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1029
	// _ = "end of CoverTab[109756]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1029
	_go_fuzz_dep_.CoverTab[109757]++
												return 4 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1030
	// _ = "end of CoverTab[109757]"
}
func sizeFixed32Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1032
	_go_fuzz_dep_.CoverTab[109760]++
												p := *ptr.toUint32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1034
		_go_fuzz_dep_.CoverTab[109762]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1035
		// _ = "end of CoverTab[109762]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1036
		_go_fuzz_dep_.CoverTab[109763]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1036
		// _ = "end of CoverTab[109763]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1036
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1036
	// _ = "end of CoverTab[109760]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1036
	_go_fuzz_dep_.CoverTab[109761]++
												return 4 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1037
	// _ = "end of CoverTab[109761]"
}
func sizeFixed32Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1039
	_go_fuzz_dep_.CoverTab[109764]++
												s := *ptr.toUint32Slice()
												return (4 + tagsize) * len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1041
	// _ = "end of CoverTab[109764]"
}
func sizeFixed32PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1043
	_go_fuzz_dep_.CoverTab[109765]++
												s := *ptr.toUint32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1045
		_go_fuzz_dep_.CoverTab[109767]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1046
		// _ = "end of CoverTab[109767]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1047
		_go_fuzz_dep_.CoverTab[109768]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1047
		// _ = "end of CoverTab[109768]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1047
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1047
	// _ = "end of CoverTab[109765]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1047
	_go_fuzz_dep_.CoverTab[109766]++
												return 4*len(s) + SizeVarint(uint64(4*len(s))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1048
	// _ = "end of CoverTab[109766]"
}
func sizeFixedS32Value(_ pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1050
	_go_fuzz_dep_.CoverTab[109769]++
												return 4 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1051
	// _ = "end of CoverTab[109769]"
}
func sizeFixedS32ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1053
	_go_fuzz_dep_.CoverTab[109770]++
												v := *ptr.toInt32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1055
		_go_fuzz_dep_.CoverTab[109772]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1056
		// _ = "end of CoverTab[109772]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1057
		_go_fuzz_dep_.CoverTab[109773]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1057
		// _ = "end of CoverTab[109773]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1057
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1057
	// _ = "end of CoverTab[109770]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1057
	_go_fuzz_dep_.CoverTab[109771]++
												return 4 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1058
	// _ = "end of CoverTab[109771]"
}
func sizeFixedS32Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1060
	_go_fuzz_dep_.CoverTab[109774]++
												p := ptr.getInt32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1062
		_go_fuzz_dep_.CoverTab[109776]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1063
		// _ = "end of CoverTab[109776]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1064
		_go_fuzz_dep_.CoverTab[109777]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1064
		// _ = "end of CoverTab[109777]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1064
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1064
	// _ = "end of CoverTab[109774]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1064
	_go_fuzz_dep_.CoverTab[109775]++
												return 4 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1065
	// _ = "end of CoverTab[109775]"
}
func sizeFixedS32Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1067
	_go_fuzz_dep_.CoverTab[109778]++
												s := ptr.getInt32Slice()
												return (4 + tagsize) * len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1069
	// _ = "end of CoverTab[109778]"
}
func sizeFixedS32PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1071
	_go_fuzz_dep_.CoverTab[109779]++
												s := ptr.getInt32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1073
		_go_fuzz_dep_.CoverTab[109781]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1074
		// _ = "end of CoverTab[109781]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1075
		_go_fuzz_dep_.CoverTab[109782]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1075
		// _ = "end of CoverTab[109782]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1075
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1075
	// _ = "end of CoverTab[109779]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1075
	_go_fuzz_dep_.CoverTab[109780]++
												return 4*len(s) + SizeVarint(uint64(4*len(s))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1076
	// _ = "end of CoverTab[109780]"
}
func sizeFloat32Value(_ pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1078
	_go_fuzz_dep_.CoverTab[109783]++
												return 4 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1079
	// _ = "end of CoverTab[109783]"
}
func sizeFloat32ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1081
	_go_fuzz_dep_.CoverTab[109784]++
												v := math.Float32bits(*ptr.toFloat32())
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1083
		_go_fuzz_dep_.CoverTab[109786]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1084
		// _ = "end of CoverTab[109786]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1085
		_go_fuzz_dep_.CoverTab[109787]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1085
		// _ = "end of CoverTab[109787]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1085
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1085
	// _ = "end of CoverTab[109784]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1085
	_go_fuzz_dep_.CoverTab[109785]++
												return 4 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1086
	// _ = "end of CoverTab[109785]"
}
func sizeFloat32Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1088
	_go_fuzz_dep_.CoverTab[109788]++
												p := *ptr.toFloat32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1090
		_go_fuzz_dep_.CoverTab[109790]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1091
		// _ = "end of CoverTab[109790]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1092
		_go_fuzz_dep_.CoverTab[109791]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1092
		// _ = "end of CoverTab[109791]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1092
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1092
	// _ = "end of CoverTab[109788]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1092
	_go_fuzz_dep_.CoverTab[109789]++
												return 4 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1093
	// _ = "end of CoverTab[109789]"
}
func sizeFloat32Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1095
	_go_fuzz_dep_.CoverTab[109792]++
												s := *ptr.toFloat32Slice()
												return (4 + tagsize) * len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1097
	// _ = "end of CoverTab[109792]"
}
func sizeFloat32PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1099
	_go_fuzz_dep_.CoverTab[109793]++
												s := *ptr.toFloat32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1101
		_go_fuzz_dep_.CoverTab[109795]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1102
		// _ = "end of CoverTab[109795]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1103
		_go_fuzz_dep_.CoverTab[109796]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1103
		// _ = "end of CoverTab[109796]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1103
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1103
	// _ = "end of CoverTab[109793]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1103
	_go_fuzz_dep_.CoverTab[109794]++
												return 4*len(s) + SizeVarint(uint64(4*len(s))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1104
	// _ = "end of CoverTab[109794]"
}
func sizeFixed64Value(_ pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1106
	_go_fuzz_dep_.CoverTab[109797]++
												return 8 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1107
	// _ = "end of CoverTab[109797]"
}
func sizeFixed64ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1109
	_go_fuzz_dep_.CoverTab[109798]++
												v := *ptr.toUint64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1111
		_go_fuzz_dep_.CoverTab[109800]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1112
		// _ = "end of CoverTab[109800]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1113
		_go_fuzz_dep_.CoverTab[109801]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1113
		// _ = "end of CoverTab[109801]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1113
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1113
	// _ = "end of CoverTab[109798]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1113
	_go_fuzz_dep_.CoverTab[109799]++
												return 8 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1114
	// _ = "end of CoverTab[109799]"
}
func sizeFixed64Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1116
	_go_fuzz_dep_.CoverTab[109802]++
												p := *ptr.toUint64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1118
		_go_fuzz_dep_.CoverTab[109804]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1119
		// _ = "end of CoverTab[109804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1120
		_go_fuzz_dep_.CoverTab[109805]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1120
		// _ = "end of CoverTab[109805]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1120
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1120
	// _ = "end of CoverTab[109802]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1120
	_go_fuzz_dep_.CoverTab[109803]++
												return 8 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1121
	// _ = "end of CoverTab[109803]"
}
func sizeFixed64Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1123
	_go_fuzz_dep_.CoverTab[109806]++
												s := *ptr.toUint64Slice()
												return (8 + tagsize) * len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1125
	// _ = "end of CoverTab[109806]"
}
func sizeFixed64PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1127
	_go_fuzz_dep_.CoverTab[109807]++
												s := *ptr.toUint64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1129
		_go_fuzz_dep_.CoverTab[109809]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1130
		// _ = "end of CoverTab[109809]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1131
		_go_fuzz_dep_.CoverTab[109810]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1131
		// _ = "end of CoverTab[109810]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1131
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1131
	// _ = "end of CoverTab[109807]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1131
	_go_fuzz_dep_.CoverTab[109808]++
												return 8*len(s) + SizeVarint(uint64(8*len(s))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1132
	// _ = "end of CoverTab[109808]"
}
func sizeFixedS64Value(_ pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1134
	_go_fuzz_dep_.CoverTab[109811]++
												return 8 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1135
	// _ = "end of CoverTab[109811]"
}
func sizeFixedS64ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1137
	_go_fuzz_dep_.CoverTab[109812]++
												v := *ptr.toInt64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1139
		_go_fuzz_dep_.CoverTab[109814]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1140
		// _ = "end of CoverTab[109814]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1141
		_go_fuzz_dep_.CoverTab[109815]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1141
		// _ = "end of CoverTab[109815]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1141
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1141
	// _ = "end of CoverTab[109812]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1141
	_go_fuzz_dep_.CoverTab[109813]++
												return 8 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1142
	// _ = "end of CoverTab[109813]"
}
func sizeFixedS64Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1144
	_go_fuzz_dep_.CoverTab[109816]++
												p := *ptr.toInt64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1146
		_go_fuzz_dep_.CoverTab[109818]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1147
		// _ = "end of CoverTab[109818]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1148
		_go_fuzz_dep_.CoverTab[109819]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1148
		// _ = "end of CoverTab[109819]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1148
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1148
	// _ = "end of CoverTab[109816]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1148
	_go_fuzz_dep_.CoverTab[109817]++
												return 8 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1149
	// _ = "end of CoverTab[109817]"
}
func sizeFixedS64Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1151
	_go_fuzz_dep_.CoverTab[109820]++
												s := *ptr.toInt64Slice()
												return (8 + tagsize) * len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1153
	// _ = "end of CoverTab[109820]"
}
func sizeFixedS64PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1155
	_go_fuzz_dep_.CoverTab[109821]++
												s := *ptr.toInt64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1157
		_go_fuzz_dep_.CoverTab[109823]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1158
		// _ = "end of CoverTab[109823]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1159
		_go_fuzz_dep_.CoverTab[109824]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1159
		// _ = "end of CoverTab[109824]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1159
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1159
	// _ = "end of CoverTab[109821]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1159
	_go_fuzz_dep_.CoverTab[109822]++
												return 8*len(s) + SizeVarint(uint64(8*len(s))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1160
	// _ = "end of CoverTab[109822]"
}
func sizeFloat64Value(_ pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1162
	_go_fuzz_dep_.CoverTab[109825]++
												return 8 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1163
	// _ = "end of CoverTab[109825]"
}
func sizeFloat64ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1165
	_go_fuzz_dep_.CoverTab[109826]++
												v := math.Float64bits(*ptr.toFloat64())
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1167
		_go_fuzz_dep_.CoverTab[109828]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1168
		// _ = "end of CoverTab[109828]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1169
		_go_fuzz_dep_.CoverTab[109829]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1169
		// _ = "end of CoverTab[109829]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1169
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1169
	// _ = "end of CoverTab[109826]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1169
	_go_fuzz_dep_.CoverTab[109827]++
												return 8 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1170
	// _ = "end of CoverTab[109827]"
}
func sizeFloat64Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1172
	_go_fuzz_dep_.CoverTab[109830]++
												p := *ptr.toFloat64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1174
		_go_fuzz_dep_.CoverTab[109832]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1175
		// _ = "end of CoverTab[109832]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1176
		_go_fuzz_dep_.CoverTab[109833]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1176
		// _ = "end of CoverTab[109833]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1176
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1176
	// _ = "end of CoverTab[109830]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1176
	_go_fuzz_dep_.CoverTab[109831]++
												return 8 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1177
	// _ = "end of CoverTab[109831]"
}
func sizeFloat64Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1179
	_go_fuzz_dep_.CoverTab[109834]++
												s := *ptr.toFloat64Slice()
												return (8 + tagsize) * len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1181
	// _ = "end of CoverTab[109834]"
}
func sizeFloat64PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1183
	_go_fuzz_dep_.CoverTab[109835]++
												s := *ptr.toFloat64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1185
		_go_fuzz_dep_.CoverTab[109837]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1186
		// _ = "end of CoverTab[109837]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1187
		_go_fuzz_dep_.CoverTab[109838]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1187
		// _ = "end of CoverTab[109838]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1187
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1187
	// _ = "end of CoverTab[109835]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1187
	_go_fuzz_dep_.CoverTab[109836]++
												return 8*len(s) + SizeVarint(uint64(8*len(s))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1188
	// _ = "end of CoverTab[109836]"
}
func sizeVarint32Value(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1190
	_go_fuzz_dep_.CoverTab[109839]++
												v := *ptr.toUint32()
												return SizeVarint(uint64(v)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1192
	// _ = "end of CoverTab[109839]"
}
func sizeVarint32ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1194
	_go_fuzz_dep_.CoverTab[109840]++
												v := *ptr.toUint32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1196
		_go_fuzz_dep_.CoverTab[109842]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1197
		// _ = "end of CoverTab[109842]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1198
		_go_fuzz_dep_.CoverTab[109843]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1198
		// _ = "end of CoverTab[109843]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1198
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1198
	// _ = "end of CoverTab[109840]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1198
	_go_fuzz_dep_.CoverTab[109841]++
												return SizeVarint(uint64(v)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1199
	// _ = "end of CoverTab[109841]"
}
func sizeVarint32Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1201
	_go_fuzz_dep_.CoverTab[109844]++
												p := *ptr.toUint32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1203
		_go_fuzz_dep_.CoverTab[109846]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1204
		// _ = "end of CoverTab[109846]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1205
		_go_fuzz_dep_.CoverTab[109847]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1205
		// _ = "end of CoverTab[109847]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1205
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1205
	// _ = "end of CoverTab[109844]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1205
	_go_fuzz_dep_.CoverTab[109845]++
												return SizeVarint(uint64(*p)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1206
	// _ = "end of CoverTab[109845]"
}
func sizeVarint32Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1208
	_go_fuzz_dep_.CoverTab[109848]++
												s := *ptr.toUint32Slice()
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1211
		_go_fuzz_dep_.CoverTab[109850]++
													n += SizeVarint(uint64(v)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1212
		// _ = "end of CoverTab[109850]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1213
	// _ = "end of CoverTab[109848]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1213
	_go_fuzz_dep_.CoverTab[109849]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1214
	// _ = "end of CoverTab[109849]"
}
func sizeVarint32PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1216
	_go_fuzz_dep_.CoverTab[109851]++
												s := *ptr.toUint32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1218
		_go_fuzz_dep_.CoverTab[109854]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1219
		// _ = "end of CoverTab[109854]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1220
		_go_fuzz_dep_.CoverTab[109855]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1220
		// _ = "end of CoverTab[109855]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1220
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1220
	// _ = "end of CoverTab[109851]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1220
	_go_fuzz_dep_.CoverTab[109852]++
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1222
		_go_fuzz_dep_.CoverTab[109856]++
													n += SizeVarint(uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1223
		// _ = "end of CoverTab[109856]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1224
	// _ = "end of CoverTab[109852]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1224
	_go_fuzz_dep_.CoverTab[109853]++
												return n + SizeVarint(uint64(n)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1225
	// _ = "end of CoverTab[109853]"
}
func sizeVarintS32Value(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1227
	_go_fuzz_dep_.CoverTab[109857]++
												v := *ptr.toInt32()
												return SizeVarint(uint64(v)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1229
	// _ = "end of CoverTab[109857]"
}
func sizeVarintS32ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1231
	_go_fuzz_dep_.CoverTab[109858]++
												v := *ptr.toInt32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1233
		_go_fuzz_dep_.CoverTab[109860]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1234
		// _ = "end of CoverTab[109860]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1235
		_go_fuzz_dep_.CoverTab[109861]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1235
		// _ = "end of CoverTab[109861]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1235
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1235
	// _ = "end of CoverTab[109858]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1235
	_go_fuzz_dep_.CoverTab[109859]++
												return SizeVarint(uint64(v)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1236
	// _ = "end of CoverTab[109859]"
}
func sizeVarintS32Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1238
	_go_fuzz_dep_.CoverTab[109862]++
												p := ptr.getInt32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1240
		_go_fuzz_dep_.CoverTab[109864]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1241
		// _ = "end of CoverTab[109864]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1242
		_go_fuzz_dep_.CoverTab[109865]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1242
		// _ = "end of CoverTab[109865]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1242
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1242
	// _ = "end of CoverTab[109862]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1242
	_go_fuzz_dep_.CoverTab[109863]++
												return SizeVarint(uint64(*p)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1243
	// _ = "end of CoverTab[109863]"
}
func sizeVarintS32Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1245
	_go_fuzz_dep_.CoverTab[109866]++
												s := ptr.getInt32Slice()
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1248
		_go_fuzz_dep_.CoverTab[109868]++
													n += SizeVarint(uint64(v)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1249
		// _ = "end of CoverTab[109868]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1250
	// _ = "end of CoverTab[109866]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1250
	_go_fuzz_dep_.CoverTab[109867]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1251
	// _ = "end of CoverTab[109867]"
}
func sizeVarintS32PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1253
	_go_fuzz_dep_.CoverTab[109869]++
												s := ptr.getInt32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1255
		_go_fuzz_dep_.CoverTab[109872]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1256
		// _ = "end of CoverTab[109872]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1257
		_go_fuzz_dep_.CoverTab[109873]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1257
		// _ = "end of CoverTab[109873]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1257
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1257
	// _ = "end of CoverTab[109869]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1257
	_go_fuzz_dep_.CoverTab[109870]++
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1259
		_go_fuzz_dep_.CoverTab[109874]++
													n += SizeVarint(uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1260
		// _ = "end of CoverTab[109874]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1261
	// _ = "end of CoverTab[109870]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1261
	_go_fuzz_dep_.CoverTab[109871]++
												return n + SizeVarint(uint64(n)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1262
	// _ = "end of CoverTab[109871]"
}
func sizeVarint64Value(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1264
	_go_fuzz_dep_.CoverTab[109875]++
												v := *ptr.toUint64()
												return SizeVarint(v) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1266
	// _ = "end of CoverTab[109875]"
}
func sizeVarint64ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1268
	_go_fuzz_dep_.CoverTab[109876]++
												v := *ptr.toUint64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1270
		_go_fuzz_dep_.CoverTab[109878]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1271
		// _ = "end of CoverTab[109878]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1272
		_go_fuzz_dep_.CoverTab[109879]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1272
		// _ = "end of CoverTab[109879]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1272
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1272
	// _ = "end of CoverTab[109876]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1272
	_go_fuzz_dep_.CoverTab[109877]++
												return SizeVarint(v) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1273
	// _ = "end of CoverTab[109877]"
}
func sizeVarint64Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1275
	_go_fuzz_dep_.CoverTab[109880]++
												p := *ptr.toUint64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1277
		_go_fuzz_dep_.CoverTab[109882]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1278
		// _ = "end of CoverTab[109882]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1279
		_go_fuzz_dep_.CoverTab[109883]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1279
		// _ = "end of CoverTab[109883]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1279
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1279
	// _ = "end of CoverTab[109880]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1279
	_go_fuzz_dep_.CoverTab[109881]++
												return SizeVarint(*p) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1280
	// _ = "end of CoverTab[109881]"
}
func sizeVarint64Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1282
	_go_fuzz_dep_.CoverTab[109884]++
												s := *ptr.toUint64Slice()
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1285
		_go_fuzz_dep_.CoverTab[109886]++
													n += SizeVarint(v) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1286
		// _ = "end of CoverTab[109886]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1287
	// _ = "end of CoverTab[109884]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1287
	_go_fuzz_dep_.CoverTab[109885]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1288
	// _ = "end of CoverTab[109885]"
}
func sizeVarint64PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1290
	_go_fuzz_dep_.CoverTab[109887]++
												s := *ptr.toUint64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1292
		_go_fuzz_dep_.CoverTab[109890]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1293
		// _ = "end of CoverTab[109890]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1294
		_go_fuzz_dep_.CoverTab[109891]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1294
		// _ = "end of CoverTab[109891]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1294
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1294
	// _ = "end of CoverTab[109887]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1294
	_go_fuzz_dep_.CoverTab[109888]++
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1296
		_go_fuzz_dep_.CoverTab[109892]++
													n += SizeVarint(v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1297
		// _ = "end of CoverTab[109892]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1298
	// _ = "end of CoverTab[109888]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1298
	_go_fuzz_dep_.CoverTab[109889]++
												return n + SizeVarint(uint64(n)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1299
	// _ = "end of CoverTab[109889]"
}
func sizeVarintS64Value(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1301
	_go_fuzz_dep_.CoverTab[109893]++
												v := *ptr.toInt64()
												return SizeVarint(uint64(v)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1303
	// _ = "end of CoverTab[109893]"
}
func sizeVarintS64ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1305
	_go_fuzz_dep_.CoverTab[109894]++
												v := *ptr.toInt64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1307
		_go_fuzz_dep_.CoverTab[109896]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1308
		// _ = "end of CoverTab[109896]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1309
		_go_fuzz_dep_.CoverTab[109897]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1309
		// _ = "end of CoverTab[109897]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1309
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1309
	// _ = "end of CoverTab[109894]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1309
	_go_fuzz_dep_.CoverTab[109895]++
												return SizeVarint(uint64(v)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1310
	// _ = "end of CoverTab[109895]"
}
func sizeVarintS64Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1312
	_go_fuzz_dep_.CoverTab[109898]++
												p := *ptr.toInt64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1314
		_go_fuzz_dep_.CoverTab[109900]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1315
		// _ = "end of CoverTab[109900]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1316
		_go_fuzz_dep_.CoverTab[109901]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1316
		// _ = "end of CoverTab[109901]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1316
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1316
	// _ = "end of CoverTab[109898]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1316
	_go_fuzz_dep_.CoverTab[109899]++
												return SizeVarint(uint64(*p)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1317
	// _ = "end of CoverTab[109899]"
}
func sizeVarintS64Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1319
	_go_fuzz_dep_.CoverTab[109902]++
												s := *ptr.toInt64Slice()
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1322
		_go_fuzz_dep_.CoverTab[109904]++
													n += SizeVarint(uint64(v)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1323
		// _ = "end of CoverTab[109904]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1324
	// _ = "end of CoverTab[109902]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1324
	_go_fuzz_dep_.CoverTab[109903]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1325
	// _ = "end of CoverTab[109903]"
}
func sizeVarintS64PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1327
	_go_fuzz_dep_.CoverTab[109905]++
												s := *ptr.toInt64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1329
		_go_fuzz_dep_.CoverTab[109908]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1330
		// _ = "end of CoverTab[109908]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1331
		_go_fuzz_dep_.CoverTab[109909]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1331
		// _ = "end of CoverTab[109909]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1331
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1331
	// _ = "end of CoverTab[109905]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1331
	_go_fuzz_dep_.CoverTab[109906]++
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1333
		_go_fuzz_dep_.CoverTab[109910]++
													n += SizeVarint(uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1334
		// _ = "end of CoverTab[109910]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1335
	// _ = "end of CoverTab[109906]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1335
	_go_fuzz_dep_.CoverTab[109907]++
												return n + SizeVarint(uint64(n)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1336
	// _ = "end of CoverTab[109907]"
}
func sizeZigzag32Value(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1338
	_go_fuzz_dep_.CoverTab[109911]++
												v := *ptr.toInt32()
												return SizeVarint(uint64((uint32(v)<<1)^uint32((int32(v)>>31)))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1340
	// _ = "end of CoverTab[109911]"
}
func sizeZigzag32ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1342
	_go_fuzz_dep_.CoverTab[109912]++
												v := *ptr.toInt32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1344
		_go_fuzz_dep_.CoverTab[109914]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1345
		// _ = "end of CoverTab[109914]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1346
		_go_fuzz_dep_.CoverTab[109915]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1346
		// _ = "end of CoverTab[109915]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1346
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1346
	// _ = "end of CoverTab[109912]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1346
	_go_fuzz_dep_.CoverTab[109913]++
												return SizeVarint(uint64((uint32(v)<<1)^uint32((int32(v)>>31)))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1347
	// _ = "end of CoverTab[109913]"
}
func sizeZigzag32Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1349
	_go_fuzz_dep_.CoverTab[109916]++
												p := ptr.getInt32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1351
		_go_fuzz_dep_.CoverTab[109918]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1352
		// _ = "end of CoverTab[109918]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1353
		_go_fuzz_dep_.CoverTab[109919]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1353
		// _ = "end of CoverTab[109919]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1353
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1353
	// _ = "end of CoverTab[109916]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1353
	_go_fuzz_dep_.CoverTab[109917]++
												v := *p
												return SizeVarint(uint64((uint32(v)<<1)^uint32((int32(v)>>31)))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1355
	// _ = "end of CoverTab[109917]"
}
func sizeZigzag32Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1357
	_go_fuzz_dep_.CoverTab[109920]++
												s := ptr.getInt32Slice()
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1360
		_go_fuzz_dep_.CoverTab[109922]++
													n += SizeVarint(uint64((uint32(v)<<1)^uint32((int32(v)>>31)))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1361
		// _ = "end of CoverTab[109922]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1362
	// _ = "end of CoverTab[109920]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1362
	_go_fuzz_dep_.CoverTab[109921]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1363
	// _ = "end of CoverTab[109921]"
}
func sizeZigzag32PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1365
	_go_fuzz_dep_.CoverTab[109923]++
												s := ptr.getInt32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1367
		_go_fuzz_dep_.CoverTab[109926]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1368
		// _ = "end of CoverTab[109926]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1369
		_go_fuzz_dep_.CoverTab[109927]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1369
		// _ = "end of CoverTab[109927]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1369
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1369
	// _ = "end of CoverTab[109923]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1369
	_go_fuzz_dep_.CoverTab[109924]++
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1371
		_go_fuzz_dep_.CoverTab[109928]++
													n += SizeVarint(uint64((uint32(v) << 1) ^ uint32((int32(v) >> 31))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1372
		// _ = "end of CoverTab[109928]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1373
	// _ = "end of CoverTab[109924]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1373
	_go_fuzz_dep_.CoverTab[109925]++
												return n + SizeVarint(uint64(n)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1374
	// _ = "end of CoverTab[109925]"
}
func sizeZigzag64Value(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1376
	_go_fuzz_dep_.CoverTab[109929]++
												v := *ptr.toInt64()
												return SizeVarint(uint64(v<<1)^uint64((int64(v)>>63))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1378
	// _ = "end of CoverTab[109929]"
}
func sizeZigzag64ValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1380
	_go_fuzz_dep_.CoverTab[109930]++
												v := *ptr.toInt64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1382
		_go_fuzz_dep_.CoverTab[109932]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1383
		// _ = "end of CoverTab[109932]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1384
		_go_fuzz_dep_.CoverTab[109933]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1384
		// _ = "end of CoverTab[109933]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1384
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1384
	// _ = "end of CoverTab[109930]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1384
	_go_fuzz_dep_.CoverTab[109931]++
												return SizeVarint(uint64(v<<1)^uint64((int64(v)>>63))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1385
	// _ = "end of CoverTab[109931]"
}
func sizeZigzag64Ptr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1387
	_go_fuzz_dep_.CoverTab[109934]++
												p := *ptr.toInt64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1389
		_go_fuzz_dep_.CoverTab[109936]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1390
		// _ = "end of CoverTab[109936]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1391
		_go_fuzz_dep_.CoverTab[109937]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1391
		// _ = "end of CoverTab[109937]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1391
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1391
	// _ = "end of CoverTab[109934]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1391
	_go_fuzz_dep_.CoverTab[109935]++
												v := *p
												return SizeVarint(uint64(v<<1)^uint64((int64(v)>>63))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1393
	// _ = "end of CoverTab[109935]"
}
func sizeZigzag64Slice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1395
	_go_fuzz_dep_.CoverTab[109938]++
												s := *ptr.toInt64Slice()
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1398
		_go_fuzz_dep_.CoverTab[109940]++
													n += SizeVarint(uint64(v<<1)^uint64((int64(v)>>63))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1399
		// _ = "end of CoverTab[109940]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1400
	// _ = "end of CoverTab[109938]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1400
	_go_fuzz_dep_.CoverTab[109939]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1401
	// _ = "end of CoverTab[109939]"
}
func sizeZigzag64PackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1403
	_go_fuzz_dep_.CoverTab[109941]++
												s := *ptr.toInt64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1405
		_go_fuzz_dep_.CoverTab[109944]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1406
		// _ = "end of CoverTab[109944]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1407
		_go_fuzz_dep_.CoverTab[109945]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1407
		// _ = "end of CoverTab[109945]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1407
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1407
	// _ = "end of CoverTab[109941]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1407
	_go_fuzz_dep_.CoverTab[109942]++
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1409
		_go_fuzz_dep_.CoverTab[109946]++
													n += SizeVarint(uint64(v<<1) ^ uint64((int64(v) >> 63)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1410
		// _ = "end of CoverTab[109946]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1411
	// _ = "end of CoverTab[109942]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1411
	_go_fuzz_dep_.CoverTab[109943]++
												return n + SizeVarint(uint64(n)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1412
	// _ = "end of CoverTab[109943]"
}
func sizeBoolValue(_ pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1414
	_go_fuzz_dep_.CoverTab[109947]++
												return 1 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1415
	// _ = "end of CoverTab[109947]"
}
func sizeBoolValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1417
	_go_fuzz_dep_.CoverTab[109948]++
												v := *ptr.toBool()
												if !v {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1419
		_go_fuzz_dep_.CoverTab[109950]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1420
		// _ = "end of CoverTab[109950]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1421
		_go_fuzz_dep_.CoverTab[109951]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1421
		// _ = "end of CoverTab[109951]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1421
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1421
	// _ = "end of CoverTab[109948]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1421
	_go_fuzz_dep_.CoverTab[109949]++
												return 1 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1422
	// _ = "end of CoverTab[109949]"
}
func sizeBoolPtr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1424
	_go_fuzz_dep_.CoverTab[109952]++
												p := *ptr.toBoolPtr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1426
		_go_fuzz_dep_.CoverTab[109954]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1427
		// _ = "end of CoverTab[109954]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1428
		_go_fuzz_dep_.CoverTab[109955]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1428
		// _ = "end of CoverTab[109955]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1428
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1428
	// _ = "end of CoverTab[109952]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1428
	_go_fuzz_dep_.CoverTab[109953]++
												return 1 + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1429
	// _ = "end of CoverTab[109953]"
}
func sizeBoolSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1431
	_go_fuzz_dep_.CoverTab[109956]++
												s := *ptr.toBoolSlice()
												return (1 + tagsize) * len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1433
	// _ = "end of CoverTab[109956]"
}
func sizeBoolPackedSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1435
	_go_fuzz_dep_.CoverTab[109957]++
												s := *ptr.toBoolSlice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1437
		_go_fuzz_dep_.CoverTab[109959]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1438
		// _ = "end of CoverTab[109959]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1439
		_go_fuzz_dep_.CoverTab[109960]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1439
		// _ = "end of CoverTab[109960]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1439
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1439
	// _ = "end of CoverTab[109957]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1439
	_go_fuzz_dep_.CoverTab[109958]++
												return len(s) + SizeVarint(uint64(len(s))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1440
	// _ = "end of CoverTab[109958]"
}
func sizeStringValue(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1442
	_go_fuzz_dep_.CoverTab[109961]++
												v := *ptr.toString()
												return len(v) + SizeVarint(uint64(len(v))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1444
	// _ = "end of CoverTab[109961]"
}
func sizeStringValueNoZero(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1446
	_go_fuzz_dep_.CoverTab[109962]++
												v := *ptr.toString()
												if v == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1448
		_go_fuzz_dep_.CoverTab[109964]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1449
		// _ = "end of CoverTab[109964]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1450
		_go_fuzz_dep_.CoverTab[109965]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1450
		// _ = "end of CoverTab[109965]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1450
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1450
	// _ = "end of CoverTab[109962]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1450
	_go_fuzz_dep_.CoverTab[109963]++
												return len(v) + SizeVarint(uint64(len(v))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1451
	// _ = "end of CoverTab[109963]"
}
func sizeStringPtr(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1453
	_go_fuzz_dep_.CoverTab[109966]++
												p := *ptr.toStringPtr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1455
		_go_fuzz_dep_.CoverTab[109968]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1456
		// _ = "end of CoverTab[109968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1457
		_go_fuzz_dep_.CoverTab[109969]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1457
		// _ = "end of CoverTab[109969]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1457
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1457
	// _ = "end of CoverTab[109966]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1457
	_go_fuzz_dep_.CoverTab[109967]++
												v := *p
												return len(v) + SizeVarint(uint64(len(v))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1459
	// _ = "end of CoverTab[109967]"
}
func sizeStringSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1461
	_go_fuzz_dep_.CoverTab[109970]++
												s := *ptr.toStringSlice()
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1464
		_go_fuzz_dep_.CoverTab[109972]++
													n += len(v) + SizeVarint(uint64(len(v))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1465
		// _ = "end of CoverTab[109972]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1466
	// _ = "end of CoverTab[109970]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1466
	_go_fuzz_dep_.CoverTab[109971]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1467
	// _ = "end of CoverTab[109971]"
}
func sizeBytes(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1469
	_go_fuzz_dep_.CoverTab[109973]++
												v := *ptr.toBytes()
												if v == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1471
		_go_fuzz_dep_.CoverTab[109975]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1472
		// _ = "end of CoverTab[109975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1473
		_go_fuzz_dep_.CoverTab[109976]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1473
		// _ = "end of CoverTab[109976]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1473
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1473
	// _ = "end of CoverTab[109973]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1473
	_go_fuzz_dep_.CoverTab[109974]++
												return len(v) + SizeVarint(uint64(len(v))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1474
	// _ = "end of CoverTab[109974]"
}
func sizeBytes3(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1476
	_go_fuzz_dep_.CoverTab[109977]++
												v := *ptr.toBytes()
												if len(v) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1478
		_go_fuzz_dep_.CoverTab[109979]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1479
		// _ = "end of CoverTab[109979]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1480
		_go_fuzz_dep_.CoverTab[109980]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1480
		// _ = "end of CoverTab[109980]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1480
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1480
	// _ = "end of CoverTab[109977]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1480
	_go_fuzz_dep_.CoverTab[109978]++
												return len(v) + SizeVarint(uint64(len(v))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1481
	// _ = "end of CoverTab[109978]"
}
func sizeBytesOneof(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1483
	_go_fuzz_dep_.CoverTab[109981]++
												v := *ptr.toBytes()
												return len(v) + SizeVarint(uint64(len(v))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1485
	// _ = "end of CoverTab[109981]"
}
func sizeBytesSlice(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1487
	_go_fuzz_dep_.CoverTab[109982]++
												s := *ptr.toBytesSlice()
												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1490
		_go_fuzz_dep_.CoverTab[109984]++
													n += len(v) + SizeVarint(uint64(len(v))) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1491
		// _ = "end of CoverTab[109984]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1492
	// _ = "end of CoverTab[109982]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1492
	_go_fuzz_dep_.CoverTab[109983]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1493
	// _ = "end of CoverTab[109983]"
}

// appendFixed32 appends an encoded fixed32 to b.
func appendFixed32(b []byte, v uint32) []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1497
	_go_fuzz_dep_.CoverTab[109985]++
												b = append(b,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24))
												return b
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1503
	// _ = "end of CoverTab[109985]"
}

// appendFixed64 appends an encoded fixed64 to b.
func appendFixed64(b []byte, v uint64) []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1507
	_go_fuzz_dep_.CoverTab[109986]++
												b = append(b,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
		byte(v>>32),
		byte(v>>40),
		byte(v>>48),
		byte(v>>56))
												return b
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1517
	// _ = "end of CoverTab[109986]"
}

// appendVarint appends an encoded varint to b.
func appendVarint(b []byte, v uint64) []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1521
	_go_fuzz_dep_.CoverTab[109987]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1524
	switch {
	case v < 1<<7:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1525
		_go_fuzz_dep_.CoverTab[109989]++
													b = append(b, byte(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1526
		// _ = "end of CoverTab[109989]"
	case v < 1<<14:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1527
		_go_fuzz_dep_.CoverTab[109990]++
													b = append(b,
			byte(v&0x7f|0x80),
			byte(v>>7))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1530
		// _ = "end of CoverTab[109990]"
	case v < 1<<21:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1531
		_go_fuzz_dep_.CoverTab[109991]++
													b = append(b,
			byte(v&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte(v>>14))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1535
		// _ = "end of CoverTab[109991]"
	case v < 1<<28:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1536
		_go_fuzz_dep_.CoverTab[109992]++
													b = append(b,
			byte(v&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte(v>>21))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1541
		// _ = "end of CoverTab[109992]"
	case v < 1<<35:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1542
		_go_fuzz_dep_.CoverTab[109993]++
													b = append(b,
			byte(v&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte(v>>28))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1548
		// _ = "end of CoverTab[109993]"
	case v < 1<<42:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1549
		_go_fuzz_dep_.CoverTab[109994]++
													b = append(b,
			byte(v&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte(v>>35))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1556
		// _ = "end of CoverTab[109994]"
	case v < 1<<49:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1557
		_go_fuzz_dep_.CoverTab[109995]++
													b = append(b,
			byte(v&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte((v>>35)&0x7f|0x80),
			byte(v>>42))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1565
		// _ = "end of CoverTab[109995]"
	case v < 1<<56:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1566
		_go_fuzz_dep_.CoverTab[109996]++
													b = append(b,
			byte(v&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte((v>>35)&0x7f|0x80),
			byte((v>>42)&0x7f|0x80),
			byte(v>>49))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1575
		// _ = "end of CoverTab[109996]"
	case v < 1<<63:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1576
		_go_fuzz_dep_.CoverTab[109997]++
													b = append(b,
			byte(v&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte((v>>35)&0x7f|0x80),
			byte((v>>42)&0x7f|0x80),
			byte((v>>49)&0x7f|0x80),
			byte(v>>56))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1586
		// _ = "end of CoverTab[109997]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1587
		_go_fuzz_dep_.CoverTab[109998]++
													b = append(b,
			byte(v&0x7f|0x80),
			byte((v>>7)&0x7f|0x80),
			byte((v>>14)&0x7f|0x80),
			byte((v>>21)&0x7f|0x80),
			byte((v>>28)&0x7f|0x80),
			byte((v>>35)&0x7f|0x80),
			byte((v>>42)&0x7f|0x80),
			byte((v>>49)&0x7f|0x80),
			byte((v>>56)&0x7f|0x80),
			1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1598
		// _ = "end of CoverTab[109998]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1599
	// _ = "end of CoverTab[109987]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1599
	_go_fuzz_dep_.CoverTab[109988]++
												return b
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1600
	// _ = "end of CoverTab[109988]"
}

func appendFixed32Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1603
	_go_fuzz_dep_.CoverTab[109999]++
												v := *ptr.toUint32()
												b = appendVarint(b, wiretag)
												b = appendFixed32(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1607
	// _ = "end of CoverTab[109999]"
}
func appendFixed32ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1609
	_go_fuzz_dep_.CoverTab[110000]++
												v := *ptr.toUint32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1611
		_go_fuzz_dep_.CoverTab[110002]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1612
		// _ = "end of CoverTab[110002]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1613
		_go_fuzz_dep_.CoverTab[110003]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1613
		// _ = "end of CoverTab[110003]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1613
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1613
	// _ = "end of CoverTab[110000]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1613
	_go_fuzz_dep_.CoverTab[110001]++
												b = appendVarint(b, wiretag)
												b = appendFixed32(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1616
	// _ = "end of CoverTab[110001]"
}
func appendFixed32Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1618
	_go_fuzz_dep_.CoverTab[110004]++
												p := *ptr.toUint32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1620
		_go_fuzz_dep_.CoverTab[110006]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1621
		// _ = "end of CoverTab[110006]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1622
		_go_fuzz_dep_.CoverTab[110007]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1622
		// _ = "end of CoverTab[110007]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1622
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1622
	// _ = "end of CoverTab[110004]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1622
	_go_fuzz_dep_.CoverTab[110005]++
												b = appendVarint(b, wiretag)
												b = appendFixed32(b, *p)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1625
	// _ = "end of CoverTab[110005]"
}
func appendFixed32Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1627
	_go_fuzz_dep_.CoverTab[110008]++
												s := *ptr.toUint32Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1629
		_go_fuzz_dep_.CoverTab[110010]++
													b = appendVarint(b, wiretag)
													b = appendFixed32(b, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1631
		// _ = "end of CoverTab[110010]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1632
	// _ = "end of CoverTab[110008]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1632
	_go_fuzz_dep_.CoverTab[110009]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1633
	// _ = "end of CoverTab[110009]"
}
func appendFixed32PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1635
	_go_fuzz_dep_.CoverTab[110011]++
												s := *ptr.toUint32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1637
		_go_fuzz_dep_.CoverTab[110014]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1638
		// _ = "end of CoverTab[110014]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1639
		_go_fuzz_dep_.CoverTab[110015]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1639
		// _ = "end of CoverTab[110015]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1639
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1639
	// _ = "end of CoverTab[110011]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1639
	_go_fuzz_dep_.CoverTab[110012]++
												b = appendVarint(b, wiretag&^7|WireBytes)
												b = appendVarint(b, uint64(4*len(s)))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1642
		_go_fuzz_dep_.CoverTab[110016]++
													b = appendFixed32(b, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1643
		// _ = "end of CoverTab[110016]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1644
	// _ = "end of CoverTab[110012]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1644
	_go_fuzz_dep_.CoverTab[110013]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1645
	// _ = "end of CoverTab[110013]"
}
func appendFixedS32Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1647
	_go_fuzz_dep_.CoverTab[110017]++
												v := *ptr.toInt32()
												b = appendVarint(b, wiretag)
												b = appendFixed32(b, uint32(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1651
	// _ = "end of CoverTab[110017]"
}
func appendFixedS32ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1653
	_go_fuzz_dep_.CoverTab[110018]++
												v := *ptr.toInt32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1655
		_go_fuzz_dep_.CoverTab[110020]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1656
		// _ = "end of CoverTab[110020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1657
		_go_fuzz_dep_.CoverTab[110021]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1657
		// _ = "end of CoverTab[110021]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1657
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1657
	// _ = "end of CoverTab[110018]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1657
	_go_fuzz_dep_.CoverTab[110019]++
												b = appendVarint(b, wiretag)
												b = appendFixed32(b, uint32(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1660
	// _ = "end of CoverTab[110019]"
}
func appendFixedS32Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1662
	_go_fuzz_dep_.CoverTab[110022]++
												p := ptr.getInt32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1664
		_go_fuzz_dep_.CoverTab[110024]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1665
		// _ = "end of CoverTab[110024]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1666
		_go_fuzz_dep_.CoverTab[110025]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1666
		// _ = "end of CoverTab[110025]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1666
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1666
	// _ = "end of CoverTab[110022]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1666
	_go_fuzz_dep_.CoverTab[110023]++
												b = appendVarint(b, wiretag)
												b = appendFixed32(b, uint32(*p))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1669
	// _ = "end of CoverTab[110023]"
}
func appendFixedS32Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1671
	_go_fuzz_dep_.CoverTab[110026]++
												s := ptr.getInt32Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1673
		_go_fuzz_dep_.CoverTab[110028]++
													b = appendVarint(b, wiretag)
													b = appendFixed32(b, uint32(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1675
		// _ = "end of CoverTab[110028]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1676
	// _ = "end of CoverTab[110026]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1676
	_go_fuzz_dep_.CoverTab[110027]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1677
	// _ = "end of CoverTab[110027]"
}
func appendFixedS32PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1679
	_go_fuzz_dep_.CoverTab[110029]++
												s := ptr.getInt32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1681
		_go_fuzz_dep_.CoverTab[110032]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1682
		// _ = "end of CoverTab[110032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1683
		_go_fuzz_dep_.CoverTab[110033]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1683
		// _ = "end of CoverTab[110033]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1683
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1683
	// _ = "end of CoverTab[110029]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1683
	_go_fuzz_dep_.CoverTab[110030]++
												b = appendVarint(b, wiretag&^7|WireBytes)
												b = appendVarint(b, uint64(4*len(s)))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1686
		_go_fuzz_dep_.CoverTab[110034]++
													b = appendFixed32(b, uint32(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1687
		// _ = "end of CoverTab[110034]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1688
	// _ = "end of CoverTab[110030]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1688
	_go_fuzz_dep_.CoverTab[110031]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1689
	// _ = "end of CoverTab[110031]"
}
func appendFloat32Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1691
	_go_fuzz_dep_.CoverTab[110035]++
												v := math.Float32bits(*ptr.toFloat32())
												b = appendVarint(b, wiretag)
												b = appendFixed32(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1695
	// _ = "end of CoverTab[110035]"
}
func appendFloat32ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1697
	_go_fuzz_dep_.CoverTab[110036]++
												v := math.Float32bits(*ptr.toFloat32())
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1699
		_go_fuzz_dep_.CoverTab[110038]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1700
		// _ = "end of CoverTab[110038]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1701
		_go_fuzz_dep_.CoverTab[110039]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1701
		// _ = "end of CoverTab[110039]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1701
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1701
	// _ = "end of CoverTab[110036]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1701
	_go_fuzz_dep_.CoverTab[110037]++
												b = appendVarint(b, wiretag)
												b = appendFixed32(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1704
	// _ = "end of CoverTab[110037]"
}
func appendFloat32Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1706
	_go_fuzz_dep_.CoverTab[110040]++
												p := *ptr.toFloat32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1708
		_go_fuzz_dep_.CoverTab[110042]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1709
		// _ = "end of CoverTab[110042]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1710
		_go_fuzz_dep_.CoverTab[110043]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1710
		// _ = "end of CoverTab[110043]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1710
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1710
	// _ = "end of CoverTab[110040]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1710
	_go_fuzz_dep_.CoverTab[110041]++
												b = appendVarint(b, wiretag)
												b = appendFixed32(b, math.Float32bits(*p))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1713
	// _ = "end of CoverTab[110041]"
}
func appendFloat32Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1715
	_go_fuzz_dep_.CoverTab[110044]++
												s := *ptr.toFloat32Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1717
		_go_fuzz_dep_.CoverTab[110046]++
													b = appendVarint(b, wiretag)
													b = appendFixed32(b, math.Float32bits(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1719
		// _ = "end of CoverTab[110046]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1720
	// _ = "end of CoverTab[110044]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1720
	_go_fuzz_dep_.CoverTab[110045]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1721
	// _ = "end of CoverTab[110045]"
}
func appendFloat32PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1723
	_go_fuzz_dep_.CoverTab[110047]++
												s := *ptr.toFloat32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1725
		_go_fuzz_dep_.CoverTab[110050]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1726
		// _ = "end of CoverTab[110050]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1727
		_go_fuzz_dep_.CoverTab[110051]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1727
		// _ = "end of CoverTab[110051]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1727
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1727
	// _ = "end of CoverTab[110047]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1727
	_go_fuzz_dep_.CoverTab[110048]++
												b = appendVarint(b, wiretag&^7|WireBytes)
												b = appendVarint(b, uint64(4*len(s)))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1730
		_go_fuzz_dep_.CoverTab[110052]++
													b = appendFixed32(b, math.Float32bits(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1731
		// _ = "end of CoverTab[110052]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1732
	// _ = "end of CoverTab[110048]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1732
	_go_fuzz_dep_.CoverTab[110049]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1733
	// _ = "end of CoverTab[110049]"
}
func appendFixed64Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1735
	_go_fuzz_dep_.CoverTab[110053]++
												v := *ptr.toUint64()
												b = appendVarint(b, wiretag)
												b = appendFixed64(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1739
	// _ = "end of CoverTab[110053]"
}
func appendFixed64ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1741
	_go_fuzz_dep_.CoverTab[110054]++
												v := *ptr.toUint64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1743
		_go_fuzz_dep_.CoverTab[110056]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1744
		// _ = "end of CoverTab[110056]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1745
		_go_fuzz_dep_.CoverTab[110057]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1745
		// _ = "end of CoverTab[110057]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1745
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1745
	// _ = "end of CoverTab[110054]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1745
	_go_fuzz_dep_.CoverTab[110055]++
												b = appendVarint(b, wiretag)
												b = appendFixed64(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1748
	// _ = "end of CoverTab[110055]"
}
func appendFixed64Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1750
	_go_fuzz_dep_.CoverTab[110058]++
												p := *ptr.toUint64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1752
		_go_fuzz_dep_.CoverTab[110060]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1753
		// _ = "end of CoverTab[110060]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1754
		_go_fuzz_dep_.CoverTab[110061]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1754
		// _ = "end of CoverTab[110061]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1754
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1754
	// _ = "end of CoverTab[110058]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1754
	_go_fuzz_dep_.CoverTab[110059]++
												b = appendVarint(b, wiretag)
												b = appendFixed64(b, *p)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1757
	// _ = "end of CoverTab[110059]"
}
func appendFixed64Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1759
	_go_fuzz_dep_.CoverTab[110062]++
												s := *ptr.toUint64Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1761
		_go_fuzz_dep_.CoverTab[110064]++
													b = appendVarint(b, wiretag)
													b = appendFixed64(b, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1763
		// _ = "end of CoverTab[110064]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1764
	// _ = "end of CoverTab[110062]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1764
	_go_fuzz_dep_.CoverTab[110063]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1765
	// _ = "end of CoverTab[110063]"
}
func appendFixed64PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1767
	_go_fuzz_dep_.CoverTab[110065]++
												s := *ptr.toUint64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1769
		_go_fuzz_dep_.CoverTab[110068]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1770
		// _ = "end of CoverTab[110068]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1771
		_go_fuzz_dep_.CoverTab[110069]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1771
		// _ = "end of CoverTab[110069]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1771
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1771
	// _ = "end of CoverTab[110065]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1771
	_go_fuzz_dep_.CoverTab[110066]++
												b = appendVarint(b, wiretag&^7|WireBytes)
												b = appendVarint(b, uint64(8*len(s)))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1774
		_go_fuzz_dep_.CoverTab[110070]++
													b = appendFixed64(b, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1775
		// _ = "end of CoverTab[110070]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1776
	// _ = "end of CoverTab[110066]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1776
	_go_fuzz_dep_.CoverTab[110067]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1777
	// _ = "end of CoverTab[110067]"
}
func appendFixedS64Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1779
	_go_fuzz_dep_.CoverTab[110071]++
												v := *ptr.toInt64()
												b = appendVarint(b, wiretag)
												b = appendFixed64(b, uint64(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1783
	// _ = "end of CoverTab[110071]"
}
func appendFixedS64ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1785
	_go_fuzz_dep_.CoverTab[110072]++
												v := *ptr.toInt64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1787
		_go_fuzz_dep_.CoverTab[110074]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1788
		// _ = "end of CoverTab[110074]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1789
		_go_fuzz_dep_.CoverTab[110075]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1789
		// _ = "end of CoverTab[110075]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1789
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1789
	// _ = "end of CoverTab[110072]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1789
	_go_fuzz_dep_.CoverTab[110073]++
												b = appendVarint(b, wiretag)
												b = appendFixed64(b, uint64(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1792
	// _ = "end of CoverTab[110073]"
}
func appendFixedS64Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1794
	_go_fuzz_dep_.CoverTab[110076]++
												p := *ptr.toInt64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1796
		_go_fuzz_dep_.CoverTab[110078]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1797
		// _ = "end of CoverTab[110078]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1798
		_go_fuzz_dep_.CoverTab[110079]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1798
		// _ = "end of CoverTab[110079]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1798
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1798
	// _ = "end of CoverTab[110076]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1798
	_go_fuzz_dep_.CoverTab[110077]++
												b = appendVarint(b, wiretag)
												b = appendFixed64(b, uint64(*p))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1801
	// _ = "end of CoverTab[110077]"
}
func appendFixedS64Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1803
	_go_fuzz_dep_.CoverTab[110080]++
												s := *ptr.toInt64Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1805
		_go_fuzz_dep_.CoverTab[110082]++
													b = appendVarint(b, wiretag)
													b = appendFixed64(b, uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1807
		// _ = "end of CoverTab[110082]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1808
	// _ = "end of CoverTab[110080]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1808
	_go_fuzz_dep_.CoverTab[110081]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1809
	// _ = "end of CoverTab[110081]"
}
func appendFixedS64PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1811
	_go_fuzz_dep_.CoverTab[110083]++
												s := *ptr.toInt64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1813
		_go_fuzz_dep_.CoverTab[110086]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1814
		// _ = "end of CoverTab[110086]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1815
		_go_fuzz_dep_.CoverTab[110087]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1815
		// _ = "end of CoverTab[110087]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1815
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1815
	// _ = "end of CoverTab[110083]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1815
	_go_fuzz_dep_.CoverTab[110084]++
												b = appendVarint(b, wiretag&^7|WireBytes)
												b = appendVarint(b, uint64(8*len(s)))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1818
		_go_fuzz_dep_.CoverTab[110088]++
													b = appendFixed64(b, uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1819
		// _ = "end of CoverTab[110088]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1820
	// _ = "end of CoverTab[110084]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1820
	_go_fuzz_dep_.CoverTab[110085]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1821
	// _ = "end of CoverTab[110085]"
}
func appendFloat64Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1823
	_go_fuzz_dep_.CoverTab[110089]++
												v := math.Float64bits(*ptr.toFloat64())
												b = appendVarint(b, wiretag)
												b = appendFixed64(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1827
	// _ = "end of CoverTab[110089]"
}
func appendFloat64ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1829
	_go_fuzz_dep_.CoverTab[110090]++
												v := math.Float64bits(*ptr.toFloat64())
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1831
		_go_fuzz_dep_.CoverTab[110092]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1832
		// _ = "end of CoverTab[110092]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1833
		_go_fuzz_dep_.CoverTab[110093]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1833
		// _ = "end of CoverTab[110093]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1833
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1833
	// _ = "end of CoverTab[110090]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1833
	_go_fuzz_dep_.CoverTab[110091]++
												b = appendVarint(b, wiretag)
												b = appendFixed64(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1836
	// _ = "end of CoverTab[110091]"
}
func appendFloat64Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1838
	_go_fuzz_dep_.CoverTab[110094]++
												p := *ptr.toFloat64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1840
		_go_fuzz_dep_.CoverTab[110096]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1841
		// _ = "end of CoverTab[110096]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1842
		_go_fuzz_dep_.CoverTab[110097]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1842
		// _ = "end of CoverTab[110097]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1842
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1842
	// _ = "end of CoverTab[110094]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1842
	_go_fuzz_dep_.CoverTab[110095]++
												b = appendVarint(b, wiretag)
												b = appendFixed64(b, math.Float64bits(*p))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1845
	// _ = "end of CoverTab[110095]"
}
func appendFloat64Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1847
	_go_fuzz_dep_.CoverTab[110098]++
												s := *ptr.toFloat64Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1849
		_go_fuzz_dep_.CoverTab[110100]++
													b = appendVarint(b, wiretag)
													b = appendFixed64(b, math.Float64bits(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1851
		// _ = "end of CoverTab[110100]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1852
	// _ = "end of CoverTab[110098]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1852
	_go_fuzz_dep_.CoverTab[110099]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1853
	// _ = "end of CoverTab[110099]"
}
func appendFloat64PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1855
	_go_fuzz_dep_.CoverTab[110101]++
												s := *ptr.toFloat64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1857
		_go_fuzz_dep_.CoverTab[110104]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1858
		// _ = "end of CoverTab[110104]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1859
		_go_fuzz_dep_.CoverTab[110105]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1859
		// _ = "end of CoverTab[110105]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1859
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1859
	// _ = "end of CoverTab[110101]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1859
	_go_fuzz_dep_.CoverTab[110102]++
												b = appendVarint(b, wiretag&^7|WireBytes)
												b = appendVarint(b, uint64(8*len(s)))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1862
		_go_fuzz_dep_.CoverTab[110106]++
													b = appendFixed64(b, math.Float64bits(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1863
		// _ = "end of CoverTab[110106]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1864
	// _ = "end of CoverTab[110102]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1864
	_go_fuzz_dep_.CoverTab[110103]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1865
	// _ = "end of CoverTab[110103]"
}
func appendVarint32Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1867
	_go_fuzz_dep_.CoverTab[110107]++
												v := *ptr.toUint32()
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1871
	// _ = "end of CoverTab[110107]"
}
func appendVarint32ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1873
	_go_fuzz_dep_.CoverTab[110108]++
												v := *ptr.toUint32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1875
		_go_fuzz_dep_.CoverTab[110110]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1876
		// _ = "end of CoverTab[110110]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1877
		_go_fuzz_dep_.CoverTab[110111]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1877
		// _ = "end of CoverTab[110111]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1877
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1877
	// _ = "end of CoverTab[110108]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1877
	_go_fuzz_dep_.CoverTab[110109]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1880
	// _ = "end of CoverTab[110109]"
}
func appendVarint32Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1882
	_go_fuzz_dep_.CoverTab[110112]++
												p := *ptr.toUint32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1884
		_go_fuzz_dep_.CoverTab[110114]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1885
		// _ = "end of CoverTab[110114]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1886
		_go_fuzz_dep_.CoverTab[110115]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1886
		// _ = "end of CoverTab[110115]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1886
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1886
	// _ = "end of CoverTab[110112]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1886
	_go_fuzz_dep_.CoverTab[110113]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(*p))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1889
	// _ = "end of CoverTab[110113]"
}
func appendVarint32Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1891
	_go_fuzz_dep_.CoverTab[110116]++
												s := *ptr.toUint32Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1893
		_go_fuzz_dep_.CoverTab[110118]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1895
		// _ = "end of CoverTab[110118]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1896
	// _ = "end of CoverTab[110116]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1896
	_go_fuzz_dep_.CoverTab[110117]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1897
	// _ = "end of CoverTab[110117]"
}
func appendVarint32PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1899
	_go_fuzz_dep_.CoverTab[110119]++
												s := *ptr.toUint32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1901
		_go_fuzz_dep_.CoverTab[110123]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1902
		// _ = "end of CoverTab[110123]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1903
		_go_fuzz_dep_.CoverTab[110124]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1903
		// _ = "end of CoverTab[110124]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1903
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1903
	// _ = "end of CoverTab[110119]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1903
	_go_fuzz_dep_.CoverTab[110120]++
												b = appendVarint(b, wiretag&^7|WireBytes)

												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1907
		_go_fuzz_dep_.CoverTab[110125]++
													n += SizeVarint(uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1908
		// _ = "end of CoverTab[110125]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1909
	// _ = "end of CoverTab[110120]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1909
	_go_fuzz_dep_.CoverTab[110121]++
												b = appendVarint(b, uint64(n))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1911
		_go_fuzz_dep_.CoverTab[110126]++
													b = appendVarint(b, uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1912
		// _ = "end of CoverTab[110126]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1913
	// _ = "end of CoverTab[110121]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1913
	_go_fuzz_dep_.CoverTab[110122]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1914
	// _ = "end of CoverTab[110122]"
}
func appendVarintS32Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1916
	_go_fuzz_dep_.CoverTab[110127]++
												v := *ptr.toInt32()
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1920
	// _ = "end of CoverTab[110127]"
}
func appendVarintS32ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1922
	_go_fuzz_dep_.CoverTab[110128]++
												v := *ptr.toInt32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1924
		_go_fuzz_dep_.CoverTab[110130]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1925
		// _ = "end of CoverTab[110130]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1926
		_go_fuzz_dep_.CoverTab[110131]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1926
		// _ = "end of CoverTab[110131]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1926
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1926
	// _ = "end of CoverTab[110128]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1926
	_go_fuzz_dep_.CoverTab[110129]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1929
	// _ = "end of CoverTab[110129]"
}
func appendVarintS32Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1931
	_go_fuzz_dep_.CoverTab[110132]++
												p := ptr.getInt32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1933
		_go_fuzz_dep_.CoverTab[110134]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1934
		// _ = "end of CoverTab[110134]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1935
		_go_fuzz_dep_.CoverTab[110135]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1935
		// _ = "end of CoverTab[110135]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1935
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1935
	// _ = "end of CoverTab[110132]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1935
	_go_fuzz_dep_.CoverTab[110133]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(*p))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1938
	// _ = "end of CoverTab[110133]"
}
func appendVarintS32Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1940
	_go_fuzz_dep_.CoverTab[110136]++
												s := ptr.getInt32Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1942
		_go_fuzz_dep_.CoverTab[110138]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1944
		// _ = "end of CoverTab[110138]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1945
	// _ = "end of CoverTab[110136]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1945
	_go_fuzz_dep_.CoverTab[110137]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1946
	// _ = "end of CoverTab[110137]"
}
func appendVarintS32PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1948
	_go_fuzz_dep_.CoverTab[110139]++
												s := ptr.getInt32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1950
		_go_fuzz_dep_.CoverTab[110143]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1951
		// _ = "end of CoverTab[110143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1952
		_go_fuzz_dep_.CoverTab[110144]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1952
		// _ = "end of CoverTab[110144]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1952
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1952
	// _ = "end of CoverTab[110139]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1952
	_go_fuzz_dep_.CoverTab[110140]++
												b = appendVarint(b, wiretag&^7|WireBytes)

												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1956
		_go_fuzz_dep_.CoverTab[110145]++
													n += SizeVarint(uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1957
		// _ = "end of CoverTab[110145]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1958
	// _ = "end of CoverTab[110140]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1958
	_go_fuzz_dep_.CoverTab[110141]++
												b = appendVarint(b, uint64(n))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1960
		_go_fuzz_dep_.CoverTab[110146]++
													b = appendVarint(b, uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1961
		// _ = "end of CoverTab[110146]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1962
	// _ = "end of CoverTab[110141]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1962
	_go_fuzz_dep_.CoverTab[110142]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1963
	// _ = "end of CoverTab[110142]"
}
func appendVarint64Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1965
	_go_fuzz_dep_.CoverTab[110147]++
												v := *ptr.toUint64()
												b = appendVarint(b, wiretag)
												b = appendVarint(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1969
	// _ = "end of CoverTab[110147]"
}
func appendVarint64ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1971
	_go_fuzz_dep_.CoverTab[110148]++
												v := *ptr.toUint64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1973
		_go_fuzz_dep_.CoverTab[110150]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1974
		// _ = "end of CoverTab[110150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1975
		_go_fuzz_dep_.CoverTab[110151]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1975
		// _ = "end of CoverTab[110151]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1975
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1975
	// _ = "end of CoverTab[110148]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1975
	_go_fuzz_dep_.CoverTab[110149]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1978
	// _ = "end of CoverTab[110149]"
}
func appendVarint64Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1980
	_go_fuzz_dep_.CoverTab[110152]++
												p := *ptr.toUint64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1982
		_go_fuzz_dep_.CoverTab[110154]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1983
		// _ = "end of CoverTab[110154]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1984
		_go_fuzz_dep_.CoverTab[110155]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1984
		// _ = "end of CoverTab[110155]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1984
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1984
	// _ = "end of CoverTab[110152]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1984
	_go_fuzz_dep_.CoverTab[110153]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, *p)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1987
	// _ = "end of CoverTab[110153]"
}
func appendVarint64Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1989
	_go_fuzz_dep_.CoverTab[110156]++
												s := *ptr.toUint64Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1991
		_go_fuzz_dep_.CoverTab[110158]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1993
		// _ = "end of CoverTab[110158]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1994
	// _ = "end of CoverTab[110156]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1994
	_go_fuzz_dep_.CoverTab[110157]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1995
	// _ = "end of CoverTab[110157]"
}
func appendVarint64PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1997
	_go_fuzz_dep_.CoverTab[110159]++
												s := *ptr.toUint64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:1999
		_go_fuzz_dep_.CoverTab[110163]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2000
		// _ = "end of CoverTab[110163]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2001
		_go_fuzz_dep_.CoverTab[110164]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2001
		// _ = "end of CoverTab[110164]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2001
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2001
	// _ = "end of CoverTab[110159]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2001
	_go_fuzz_dep_.CoverTab[110160]++
												b = appendVarint(b, wiretag&^7|WireBytes)

												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2005
		_go_fuzz_dep_.CoverTab[110165]++
													n += SizeVarint(v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2006
		// _ = "end of CoverTab[110165]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2007
	// _ = "end of CoverTab[110160]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2007
	_go_fuzz_dep_.CoverTab[110161]++
												b = appendVarint(b, uint64(n))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2009
		_go_fuzz_dep_.CoverTab[110166]++
													b = appendVarint(b, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2010
		// _ = "end of CoverTab[110166]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2011
	// _ = "end of CoverTab[110161]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2011
	_go_fuzz_dep_.CoverTab[110162]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2012
	// _ = "end of CoverTab[110162]"
}
func appendVarintS64Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2014
	_go_fuzz_dep_.CoverTab[110167]++
												v := *ptr.toInt64()
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2018
	// _ = "end of CoverTab[110167]"
}
func appendVarintS64ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2020
	_go_fuzz_dep_.CoverTab[110168]++
												v := *ptr.toInt64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2022
		_go_fuzz_dep_.CoverTab[110170]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2023
		// _ = "end of CoverTab[110170]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2024
		_go_fuzz_dep_.CoverTab[110171]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2024
		// _ = "end of CoverTab[110171]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2024
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2024
	// _ = "end of CoverTab[110168]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2024
	_go_fuzz_dep_.CoverTab[110169]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(v))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2027
	// _ = "end of CoverTab[110169]"
}
func appendVarintS64Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2029
	_go_fuzz_dep_.CoverTab[110172]++
												p := *ptr.toInt64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2031
		_go_fuzz_dep_.CoverTab[110174]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2032
		// _ = "end of CoverTab[110174]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2033
		_go_fuzz_dep_.CoverTab[110175]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2033
		// _ = "end of CoverTab[110175]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2033
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2033
	// _ = "end of CoverTab[110172]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2033
	_go_fuzz_dep_.CoverTab[110173]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(*p))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2036
	// _ = "end of CoverTab[110173]"
}
func appendVarintS64Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2038
	_go_fuzz_dep_.CoverTab[110176]++
												s := *ptr.toInt64Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2040
		_go_fuzz_dep_.CoverTab[110178]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2042
		// _ = "end of CoverTab[110178]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2043
	// _ = "end of CoverTab[110176]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2043
	_go_fuzz_dep_.CoverTab[110177]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2044
	// _ = "end of CoverTab[110177]"
}
func appendVarintS64PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2046
	_go_fuzz_dep_.CoverTab[110179]++
												s := *ptr.toInt64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2048
		_go_fuzz_dep_.CoverTab[110183]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2049
		// _ = "end of CoverTab[110183]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2050
		_go_fuzz_dep_.CoverTab[110184]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2050
		// _ = "end of CoverTab[110184]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2050
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2050
	// _ = "end of CoverTab[110179]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2050
	_go_fuzz_dep_.CoverTab[110180]++
												b = appendVarint(b, wiretag&^7|WireBytes)

												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2054
		_go_fuzz_dep_.CoverTab[110185]++
													n += SizeVarint(uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2055
		// _ = "end of CoverTab[110185]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2056
	// _ = "end of CoverTab[110180]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2056
	_go_fuzz_dep_.CoverTab[110181]++
												b = appendVarint(b, uint64(n))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2058
		_go_fuzz_dep_.CoverTab[110186]++
													b = appendVarint(b, uint64(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2059
		// _ = "end of CoverTab[110186]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2060
	// _ = "end of CoverTab[110181]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2060
	_go_fuzz_dep_.CoverTab[110182]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2061
	// _ = "end of CoverTab[110182]"
}
func appendZigzag32Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2063
	_go_fuzz_dep_.CoverTab[110187]++
												v := *ptr.toInt32()
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64((uint32(v)<<1)^uint32((int32(v)>>31))))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2067
	// _ = "end of CoverTab[110187]"
}
func appendZigzag32ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2069
	_go_fuzz_dep_.CoverTab[110188]++
												v := *ptr.toInt32()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2071
		_go_fuzz_dep_.CoverTab[110190]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2072
		// _ = "end of CoverTab[110190]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2073
		_go_fuzz_dep_.CoverTab[110191]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2073
		// _ = "end of CoverTab[110191]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2073
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2073
	// _ = "end of CoverTab[110188]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2073
	_go_fuzz_dep_.CoverTab[110189]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64((uint32(v)<<1)^uint32((int32(v)>>31))))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2076
	// _ = "end of CoverTab[110189]"
}
func appendZigzag32Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2078
	_go_fuzz_dep_.CoverTab[110192]++
												p := ptr.getInt32Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2080
		_go_fuzz_dep_.CoverTab[110194]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2081
		// _ = "end of CoverTab[110194]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2082
		_go_fuzz_dep_.CoverTab[110195]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2082
		// _ = "end of CoverTab[110195]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2082
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2082
	// _ = "end of CoverTab[110192]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2082
	_go_fuzz_dep_.CoverTab[110193]++
												b = appendVarint(b, wiretag)
												v := *p
												b = appendVarint(b, uint64((uint32(v)<<1)^uint32((int32(v)>>31))))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2086
	// _ = "end of CoverTab[110193]"
}
func appendZigzag32Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2088
	_go_fuzz_dep_.CoverTab[110196]++
												s := ptr.getInt32Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2090
		_go_fuzz_dep_.CoverTab[110198]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64((uint32(v)<<1)^uint32((int32(v)>>31))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2092
		// _ = "end of CoverTab[110198]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2093
	// _ = "end of CoverTab[110196]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2093
	_go_fuzz_dep_.CoverTab[110197]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2094
	// _ = "end of CoverTab[110197]"
}
func appendZigzag32PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2096
	_go_fuzz_dep_.CoverTab[110199]++
												s := ptr.getInt32Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2098
		_go_fuzz_dep_.CoverTab[110203]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2099
		// _ = "end of CoverTab[110203]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2100
		_go_fuzz_dep_.CoverTab[110204]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2100
		// _ = "end of CoverTab[110204]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2100
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2100
	// _ = "end of CoverTab[110199]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2100
	_go_fuzz_dep_.CoverTab[110200]++
												b = appendVarint(b, wiretag&^7|WireBytes)

												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2104
		_go_fuzz_dep_.CoverTab[110205]++
													n += SizeVarint(uint64((uint32(v) << 1) ^ uint32((int32(v) >> 31))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2105
		// _ = "end of CoverTab[110205]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2106
	// _ = "end of CoverTab[110200]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2106
	_go_fuzz_dep_.CoverTab[110201]++
												b = appendVarint(b, uint64(n))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2108
		_go_fuzz_dep_.CoverTab[110206]++
													b = appendVarint(b, uint64((uint32(v)<<1)^uint32((int32(v)>>31))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2109
		// _ = "end of CoverTab[110206]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2110
	// _ = "end of CoverTab[110201]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2110
	_go_fuzz_dep_.CoverTab[110202]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2111
	// _ = "end of CoverTab[110202]"
}
func appendZigzag64Value(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2113
	_go_fuzz_dep_.CoverTab[110207]++
												v := *ptr.toInt64()
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(v<<1)^uint64((int64(v)>>63)))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2117
	// _ = "end of CoverTab[110207]"
}
func appendZigzag64ValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2119
	_go_fuzz_dep_.CoverTab[110208]++
												v := *ptr.toInt64()
												if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2121
		_go_fuzz_dep_.CoverTab[110210]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2122
		// _ = "end of CoverTab[110210]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2123
		_go_fuzz_dep_.CoverTab[110211]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2123
		// _ = "end of CoverTab[110211]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2123
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2123
	// _ = "end of CoverTab[110208]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2123
	_go_fuzz_dep_.CoverTab[110209]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(v<<1)^uint64((int64(v)>>63)))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2126
	// _ = "end of CoverTab[110209]"
}
func appendZigzag64Ptr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2128
	_go_fuzz_dep_.CoverTab[110212]++
												p := *ptr.toInt64Ptr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2130
		_go_fuzz_dep_.CoverTab[110214]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2131
		// _ = "end of CoverTab[110214]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2132
		_go_fuzz_dep_.CoverTab[110215]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2132
		// _ = "end of CoverTab[110215]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2132
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2132
	// _ = "end of CoverTab[110212]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2132
	_go_fuzz_dep_.CoverTab[110213]++
												b = appendVarint(b, wiretag)
												v := *p
												b = appendVarint(b, uint64(v<<1)^uint64((int64(v)>>63)))
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2136
	// _ = "end of CoverTab[110213]"
}
func appendZigzag64Slice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2138
	_go_fuzz_dep_.CoverTab[110216]++
												s := *ptr.toInt64Slice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2140
		_go_fuzz_dep_.CoverTab[110218]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64(v<<1)^uint64((int64(v)>>63)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2142
		// _ = "end of CoverTab[110218]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2143
	// _ = "end of CoverTab[110216]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2143
	_go_fuzz_dep_.CoverTab[110217]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2144
	// _ = "end of CoverTab[110217]"
}
func appendZigzag64PackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2146
	_go_fuzz_dep_.CoverTab[110219]++
												s := *ptr.toInt64Slice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2148
		_go_fuzz_dep_.CoverTab[110223]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2149
		// _ = "end of CoverTab[110223]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2150
		_go_fuzz_dep_.CoverTab[110224]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2150
		// _ = "end of CoverTab[110224]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2150
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2150
	// _ = "end of CoverTab[110219]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2150
	_go_fuzz_dep_.CoverTab[110220]++
												b = appendVarint(b, wiretag&^7|WireBytes)

												n := 0
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2154
		_go_fuzz_dep_.CoverTab[110225]++
													n += SizeVarint(uint64(v<<1) ^ uint64((int64(v) >> 63)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2155
		// _ = "end of CoverTab[110225]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2156
	// _ = "end of CoverTab[110220]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2156
	_go_fuzz_dep_.CoverTab[110221]++
												b = appendVarint(b, uint64(n))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2158
		_go_fuzz_dep_.CoverTab[110226]++
													b = appendVarint(b, uint64(v<<1)^uint64((int64(v)>>63)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2159
		// _ = "end of CoverTab[110226]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2160
	// _ = "end of CoverTab[110221]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2160
	_go_fuzz_dep_.CoverTab[110222]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2161
	// _ = "end of CoverTab[110222]"
}
func appendBoolValue(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2163
	_go_fuzz_dep_.CoverTab[110227]++
												v := *ptr.toBool()
												b = appendVarint(b, wiretag)
												if v {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2166
		_go_fuzz_dep_.CoverTab[110229]++
													b = append(b, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2167
		// _ = "end of CoverTab[110229]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2168
		_go_fuzz_dep_.CoverTab[110230]++
													b = append(b, 0)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2169
		// _ = "end of CoverTab[110230]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2170
	// _ = "end of CoverTab[110227]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2170
	_go_fuzz_dep_.CoverTab[110228]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2171
	// _ = "end of CoverTab[110228]"
}
func appendBoolValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2173
	_go_fuzz_dep_.CoverTab[110231]++
												v := *ptr.toBool()
												if !v {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2175
		_go_fuzz_dep_.CoverTab[110233]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2176
		// _ = "end of CoverTab[110233]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2177
		_go_fuzz_dep_.CoverTab[110234]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2177
		// _ = "end of CoverTab[110234]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2177
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2177
	// _ = "end of CoverTab[110231]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2177
	_go_fuzz_dep_.CoverTab[110232]++
												b = appendVarint(b, wiretag)
												b = append(b, 1)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2180
	// _ = "end of CoverTab[110232]"
}

func appendBoolPtr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2183
	_go_fuzz_dep_.CoverTab[110235]++
												p := *ptr.toBoolPtr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2185
		_go_fuzz_dep_.CoverTab[110238]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2186
		// _ = "end of CoverTab[110238]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2187
		_go_fuzz_dep_.CoverTab[110239]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2187
		// _ = "end of CoverTab[110239]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2187
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2187
	// _ = "end of CoverTab[110235]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2187
	_go_fuzz_dep_.CoverTab[110236]++
												b = appendVarint(b, wiretag)
												if *p {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2189
		_go_fuzz_dep_.CoverTab[110240]++
													b = append(b, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2190
		// _ = "end of CoverTab[110240]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2191
		_go_fuzz_dep_.CoverTab[110241]++
													b = append(b, 0)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2192
		// _ = "end of CoverTab[110241]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2193
	// _ = "end of CoverTab[110236]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2193
	_go_fuzz_dep_.CoverTab[110237]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2194
	// _ = "end of CoverTab[110237]"
}
func appendBoolSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2196
	_go_fuzz_dep_.CoverTab[110242]++
												s := *ptr.toBoolSlice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2198
		_go_fuzz_dep_.CoverTab[110244]++
													b = appendVarint(b, wiretag)
													if v {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2200
			_go_fuzz_dep_.CoverTab[110245]++
														b = append(b, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2201
			// _ = "end of CoverTab[110245]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2202
			_go_fuzz_dep_.CoverTab[110246]++
														b = append(b, 0)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2203
			// _ = "end of CoverTab[110246]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2204
		// _ = "end of CoverTab[110244]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2205
	// _ = "end of CoverTab[110242]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2205
	_go_fuzz_dep_.CoverTab[110243]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2206
	// _ = "end of CoverTab[110243]"
}
func appendBoolPackedSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2208
	_go_fuzz_dep_.CoverTab[110247]++
												s := *ptr.toBoolSlice()
												if len(s) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2210
		_go_fuzz_dep_.CoverTab[110250]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2211
		// _ = "end of CoverTab[110250]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2212
		_go_fuzz_dep_.CoverTab[110251]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2212
		// _ = "end of CoverTab[110251]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2212
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2212
	// _ = "end of CoverTab[110247]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2212
	_go_fuzz_dep_.CoverTab[110248]++
												b = appendVarint(b, wiretag&^7|WireBytes)
												b = appendVarint(b, uint64(len(s)))
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2215
		_go_fuzz_dep_.CoverTab[110252]++
													if v {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2216
			_go_fuzz_dep_.CoverTab[110253]++
														b = append(b, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2217
			// _ = "end of CoverTab[110253]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2218
			_go_fuzz_dep_.CoverTab[110254]++
														b = append(b, 0)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2219
			// _ = "end of CoverTab[110254]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2220
		// _ = "end of CoverTab[110252]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2221
	// _ = "end of CoverTab[110248]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2221
	_go_fuzz_dep_.CoverTab[110249]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2222
	// _ = "end of CoverTab[110249]"
}
func appendStringValue(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2224
	_go_fuzz_dep_.CoverTab[110255]++
												v := *ptr.toString()
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(len(v)))
												b = append(b, v...)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2229
	// _ = "end of CoverTab[110255]"
}
func appendStringValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2231
	_go_fuzz_dep_.CoverTab[110256]++
												v := *ptr.toString()
												if v == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2233
		_go_fuzz_dep_.CoverTab[110258]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2234
		// _ = "end of CoverTab[110258]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2235
		_go_fuzz_dep_.CoverTab[110259]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2235
		// _ = "end of CoverTab[110259]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2235
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2235
	// _ = "end of CoverTab[110256]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2235
	_go_fuzz_dep_.CoverTab[110257]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(len(v)))
												b = append(b, v...)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2239
	// _ = "end of CoverTab[110257]"
}
func appendStringPtr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2241
	_go_fuzz_dep_.CoverTab[110260]++
												p := *ptr.toStringPtr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2243
		_go_fuzz_dep_.CoverTab[110262]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2244
		// _ = "end of CoverTab[110262]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2245
		_go_fuzz_dep_.CoverTab[110263]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2245
		// _ = "end of CoverTab[110263]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2245
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2245
	// _ = "end of CoverTab[110260]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2245
	_go_fuzz_dep_.CoverTab[110261]++
												v := *p
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(len(v)))
												b = append(b, v...)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2250
	// _ = "end of CoverTab[110261]"
}
func appendStringSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2252
	_go_fuzz_dep_.CoverTab[110264]++
												s := *ptr.toStringSlice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2254
		_go_fuzz_dep_.CoverTab[110266]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64(len(v)))
													b = append(b, v...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2257
		// _ = "end of CoverTab[110266]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2258
	// _ = "end of CoverTab[110264]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2258
	_go_fuzz_dep_.CoverTab[110265]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2259
	// _ = "end of CoverTab[110265]"
}
func appendUTF8StringValue(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2261
	_go_fuzz_dep_.CoverTab[110267]++
												var invalidUTF8 bool
												v := *ptr.toString()
												if !utf8.ValidString(v) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2264
		_go_fuzz_dep_.CoverTab[110270]++
													invalidUTF8 = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2265
		// _ = "end of CoverTab[110270]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2266
		_go_fuzz_dep_.CoverTab[110271]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2266
		// _ = "end of CoverTab[110271]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2266
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2266
	// _ = "end of CoverTab[110267]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2266
	_go_fuzz_dep_.CoverTab[110268]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(len(v)))
												b = append(b, v...)
												if invalidUTF8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2270
		_go_fuzz_dep_.CoverTab[110272]++
													return b, errInvalidUTF8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2271
		// _ = "end of CoverTab[110272]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2272
		_go_fuzz_dep_.CoverTab[110273]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2272
		// _ = "end of CoverTab[110273]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2272
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2272
	// _ = "end of CoverTab[110268]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2272
	_go_fuzz_dep_.CoverTab[110269]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2273
	// _ = "end of CoverTab[110269]"
}
func appendUTF8StringValueNoZero(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2275
	_go_fuzz_dep_.CoverTab[110274]++
												var invalidUTF8 bool
												v := *ptr.toString()
												if v == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2278
		_go_fuzz_dep_.CoverTab[110278]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2279
		// _ = "end of CoverTab[110278]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2280
		_go_fuzz_dep_.CoverTab[110279]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2280
		// _ = "end of CoverTab[110279]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2280
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2280
	// _ = "end of CoverTab[110274]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2280
	_go_fuzz_dep_.CoverTab[110275]++
												if !utf8.ValidString(v) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2281
		_go_fuzz_dep_.CoverTab[110280]++
													invalidUTF8 = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2282
		// _ = "end of CoverTab[110280]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2283
		_go_fuzz_dep_.CoverTab[110281]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2283
		// _ = "end of CoverTab[110281]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2283
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2283
	// _ = "end of CoverTab[110275]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2283
	_go_fuzz_dep_.CoverTab[110276]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(len(v)))
												b = append(b, v...)
												if invalidUTF8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2287
		_go_fuzz_dep_.CoverTab[110282]++
													return b, errInvalidUTF8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2288
		// _ = "end of CoverTab[110282]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2289
		_go_fuzz_dep_.CoverTab[110283]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2289
		// _ = "end of CoverTab[110283]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2289
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2289
	// _ = "end of CoverTab[110276]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2289
	_go_fuzz_dep_.CoverTab[110277]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2290
	// _ = "end of CoverTab[110277]"
}
func appendUTF8StringPtr(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2292
	_go_fuzz_dep_.CoverTab[110284]++
												var invalidUTF8 bool
												p := *ptr.toStringPtr()
												if p == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2295
		_go_fuzz_dep_.CoverTab[110288]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2296
		// _ = "end of CoverTab[110288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2297
		_go_fuzz_dep_.CoverTab[110289]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2297
		// _ = "end of CoverTab[110289]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2297
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2297
	// _ = "end of CoverTab[110284]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2297
	_go_fuzz_dep_.CoverTab[110285]++
												v := *p
												if !utf8.ValidString(v) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2299
		_go_fuzz_dep_.CoverTab[110290]++
													invalidUTF8 = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2300
		// _ = "end of CoverTab[110290]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2301
		_go_fuzz_dep_.CoverTab[110291]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2301
		// _ = "end of CoverTab[110291]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2301
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2301
	// _ = "end of CoverTab[110285]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2301
	_go_fuzz_dep_.CoverTab[110286]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(len(v)))
												b = append(b, v...)
												if invalidUTF8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2305
		_go_fuzz_dep_.CoverTab[110292]++
													return b, errInvalidUTF8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2306
		// _ = "end of CoverTab[110292]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2307
		_go_fuzz_dep_.CoverTab[110293]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2307
		// _ = "end of CoverTab[110293]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2307
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2307
	// _ = "end of CoverTab[110286]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2307
	_go_fuzz_dep_.CoverTab[110287]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2308
	// _ = "end of CoverTab[110287]"
}
func appendUTF8StringSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2310
	_go_fuzz_dep_.CoverTab[110294]++
												var invalidUTF8 bool
												s := *ptr.toStringSlice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2313
		_go_fuzz_dep_.CoverTab[110297]++
													if !utf8.ValidString(v) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2314
			_go_fuzz_dep_.CoverTab[110299]++
														invalidUTF8 = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2315
			// _ = "end of CoverTab[110299]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2316
			_go_fuzz_dep_.CoverTab[110300]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2316
			// _ = "end of CoverTab[110300]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2316
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2316
		// _ = "end of CoverTab[110297]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2316
		_go_fuzz_dep_.CoverTab[110298]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64(len(v)))
													b = append(b, v...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2319
		// _ = "end of CoverTab[110298]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2320
	// _ = "end of CoverTab[110294]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2320
	_go_fuzz_dep_.CoverTab[110295]++
												if invalidUTF8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2321
		_go_fuzz_dep_.CoverTab[110301]++
													return b, errInvalidUTF8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2322
		// _ = "end of CoverTab[110301]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2323
		_go_fuzz_dep_.CoverTab[110302]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2323
		// _ = "end of CoverTab[110302]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2323
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2323
	// _ = "end of CoverTab[110295]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2323
	_go_fuzz_dep_.CoverTab[110296]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2324
	// _ = "end of CoverTab[110296]"
}
func appendBytes(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2326
	_go_fuzz_dep_.CoverTab[110303]++
												v := *ptr.toBytes()
												if v == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2328
		_go_fuzz_dep_.CoverTab[110305]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2329
		// _ = "end of CoverTab[110305]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2330
		_go_fuzz_dep_.CoverTab[110306]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2330
		// _ = "end of CoverTab[110306]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2330
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2330
	// _ = "end of CoverTab[110303]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2330
	_go_fuzz_dep_.CoverTab[110304]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(len(v)))
												b = append(b, v...)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2334
	// _ = "end of CoverTab[110304]"
}
func appendBytes3(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2336
	_go_fuzz_dep_.CoverTab[110307]++
												v := *ptr.toBytes()
												if len(v) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2338
		_go_fuzz_dep_.CoverTab[110309]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2339
		// _ = "end of CoverTab[110309]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2340
		_go_fuzz_dep_.CoverTab[110310]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2340
		// _ = "end of CoverTab[110310]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2340
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2340
	// _ = "end of CoverTab[110307]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2340
	_go_fuzz_dep_.CoverTab[110308]++
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(len(v)))
												b = append(b, v...)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2344
	// _ = "end of CoverTab[110308]"
}
func appendBytesOneof(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2346
	_go_fuzz_dep_.CoverTab[110311]++
												v := *ptr.toBytes()
												b = appendVarint(b, wiretag)
												b = appendVarint(b, uint64(len(v)))
												b = append(b, v...)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2351
	// _ = "end of CoverTab[110311]"
}
func appendBytesSlice(b []byte, ptr pointer, wiretag uint64, _ bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2353
	_go_fuzz_dep_.CoverTab[110312]++
												s := *ptr.toBytesSlice()
												for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2355
		_go_fuzz_dep_.CoverTab[110314]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64(len(v)))
													b = append(b, v...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2358
		// _ = "end of CoverTab[110314]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2359
	// _ = "end of CoverTab[110312]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2359
	_go_fuzz_dep_.CoverTab[110313]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2360
	// _ = "end of CoverTab[110313]"
}

// makeGroupMarshaler returns the sizer and marshaler for a group.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2363
// u is the marshal info of the underlying message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2365
func makeGroupMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2365
	_go_fuzz_dep_.CoverTab[110315]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2366
			_go_fuzz_dep_.CoverTab[110316]++
														p := ptr.getPointer()
														if p.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2368
				_go_fuzz_dep_.CoverTab[110318]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2369
				// _ = "end of CoverTab[110318]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2370
				_go_fuzz_dep_.CoverTab[110319]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2370
				// _ = "end of CoverTab[110319]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2370
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2370
			// _ = "end of CoverTab[110316]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2370
			_go_fuzz_dep_.CoverTab[110317]++
														return u.size(p) + 2*tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2371
			// _ = "end of CoverTab[110317]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2373
			_go_fuzz_dep_.CoverTab[110320]++
														p := ptr.getPointer()
														if p.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2375
				_go_fuzz_dep_.CoverTab[110322]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2376
				// _ = "end of CoverTab[110322]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2377
				_go_fuzz_dep_.CoverTab[110323]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2377
				// _ = "end of CoverTab[110323]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2377
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2377
			// _ = "end of CoverTab[110320]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2377
			_go_fuzz_dep_.CoverTab[110321]++
														var err error
														b = appendVarint(b, wiretag)
														b, err = u.marshal(b, p, deterministic)
														b = appendVarint(b, wiretag+(WireEndGroup-WireStartGroup))
														return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2382
			// _ = "end of CoverTab[110321]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2383
	// _ = "end of CoverTab[110315]"
}

// makeGroupSliceMarshaler returns the sizer and marshaler for a group slice.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2386
// u is the marshal info of the underlying message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2388
func makeGroupSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2388
	_go_fuzz_dep_.CoverTab[110324]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2389
			_go_fuzz_dep_.CoverTab[110325]++
														s := ptr.getPointerSlice()
														n := 0
														for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2392
				_go_fuzz_dep_.CoverTab[110327]++
															if v.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2393
					_go_fuzz_dep_.CoverTab[110329]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2394
					// _ = "end of CoverTab[110329]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2395
					_go_fuzz_dep_.CoverTab[110330]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2395
					// _ = "end of CoverTab[110330]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2395
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2395
				// _ = "end of CoverTab[110327]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2395
				_go_fuzz_dep_.CoverTab[110328]++
															n += u.size(v) + 2*tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2396
				// _ = "end of CoverTab[110328]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2397
			// _ = "end of CoverTab[110325]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2397
			_go_fuzz_dep_.CoverTab[110326]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2398
			// _ = "end of CoverTab[110326]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2400
			_go_fuzz_dep_.CoverTab[110331]++
														s := ptr.getPointerSlice()
														var err error
														var nerr nonFatal
														for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2404
				_go_fuzz_dep_.CoverTab[110333]++
															if v.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2405
					_go_fuzz_dep_.CoverTab[110335]++
																return b, errRepeatedHasNil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2406
					// _ = "end of CoverTab[110335]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2407
					_go_fuzz_dep_.CoverTab[110336]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2407
					// _ = "end of CoverTab[110336]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2407
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2407
				// _ = "end of CoverTab[110333]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2407
				_go_fuzz_dep_.CoverTab[110334]++
															b = appendVarint(b, wiretag)
															b, err = u.marshal(b, v, deterministic)
															b = appendVarint(b, wiretag+(WireEndGroup-WireStartGroup))
															if !nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2411
					_go_fuzz_dep_.CoverTab[110337]++
																if err == ErrNil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2412
						_go_fuzz_dep_.CoverTab[110339]++
																	err = errRepeatedHasNil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2413
						// _ = "end of CoverTab[110339]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2414
						_go_fuzz_dep_.CoverTab[110340]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2414
						// _ = "end of CoverTab[110340]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2414
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2414
					// _ = "end of CoverTab[110337]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2414
					_go_fuzz_dep_.CoverTab[110338]++
																return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2415
					// _ = "end of CoverTab[110338]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2416
					_go_fuzz_dep_.CoverTab[110341]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2416
					// _ = "end of CoverTab[110341]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2416
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2416
				// _ = "end of CoverTab[110334]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2417
			// _ = "end of CoverTab[110331]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2417
			_go_fuzz_dep_.CoverTab[110332]++
														return b, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2418
			// _ = "end of CoverTab[110332]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2419
	// _ = "end of CoverTab[110324]"
}

// makeMessageMarshaler returns the sizer and marshaler for a message field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2422
// u is the marshal info of the message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2424
func makeMessageMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2424
	_go_fuzz_dep_.CoverTab[110342]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2425
			_go_fuzz_dep_.CoverTab[110343]++
														p := ptr.getPointer()
														if p.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2427
				_go_fuzz_dep_.CoverTab[110345]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2428
				// _ = "end of CoverTab[110345]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2429
				_go_fuzz_dep_.CoverTab[110346]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2429
				// _ = "end of CoverTab[110346]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2429
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2429
			// _ = "end of CoverTab[110343]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2429
			_go_fuzz_dep_.CoverTab[110344]++
														siz := u.size(p)
														return siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2431
			// _ = "end of CoverTab[110344]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2433
			_go_fuzz_dep_.CoverTab[110347]++
														p := ptr.getPointer()
														if p.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2435
				_go_fuzz_dep_.CoverTab[110349]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2436
				// _ = "end of CoverTab[110349]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2437
				_go_fuzz_dep_.CoverTab[110350]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2437
				// _ = "end of CoverTab[110350]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2437
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2437
			// _ = "end of CoverTab[110347]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2437
			_go_fuzz_dep_.CoverTab[110348]++
														b = appendVarint(b, wiretag)
														siz := u.cachedsize(p)
														b = appendVarint(b, uint64(siz))
														return u.marshal(b, p, deterministic)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2441
			// _ = "end of CoverTab[110348]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2442
	// _ = "end of CoverTab[110342]"
}

// makeMessageSliceMarshaler returns the sizer and marshaler for a message slice.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2445
// u is the marshal info of the message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2447
func makeMessageSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2447
	_go_fuzz_dep_.CoverTab[110351]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2448
			_go_fuzz_dep_.CoverTab[110352]++
														s := ptr.getPointerSlice()
														n := 0
														for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2451
				_go_fuzz_dep_.CoverTab[110354]++
															if v.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2452
					_go_fuzz_dep_.CoverTab[110356]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2453
					// _ = "end of CoverTab[110356]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2454
					_go_fuzz_dep_.CoverTab[110357]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2454
					// _ = "end of CoverTab[110357]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2454
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2454
				// _ = "end of CoverTab[110354]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2454
				_go_fuzz_dep_.CoverTab[110355]++
															siz := u.size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2456
				// _ = "end of CoverTab[110355]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2457
			// _ = "end of CoverTab[110352]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2457
			_go_fuzz_dep_.CoverTab[110353]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2458
			// _ = "end of CoverTab[110353]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2460
			_go_fuzz_dep_.CoverTab[110358]++
														s := ptr.getPointerSlice()
														var err error
														var nerr nonFatal
														for _, v := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2464
				_go_fuzz_dep_.CoverTab[110360]++
															if v.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2465
					_go_fuzz_dep_.CoverTab[110362]++
																return b, errRepeatedHasNil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2466
					// _ = "end of CoverTab[110362]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2467
					_go_fuzz_dep_.CoverTab[110363]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2467
					// _ = "end of CoverTab[110363]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2467
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2467
				// _ = "end of CoverTab[110360]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2467
				_go_fuzz_dep_.CoverTab[110361]++
															b = appendVarint(b, wiretag)
															siz := u.cachedsize(v)
															b = appendVarint(b, uint64(siz))
															b, err = u.marshal(b, v, deterministic)

															if !nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2473
					_go_fuzz_dep_.CoverTab[110364]++
																if err == ErrNil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2474
						_go_fuzz_dep_.CoverTab[110366]++
																	err = errRepeatedHasNil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2475
						// _ = "end of CoverTab[110366]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2476
						_go_fuzz_dep_.CoverTab[110367]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2476
						// _ = "end of CoverTab[110367]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2476
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2476
					// _ = "end of CoverTab[110364]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2476
					_go_fuzz_dep_.CoverTab[110365]++
																return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2477
					// _ = "end of CoverTab[110365]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2478
					_go_fuzz_dep_.CoverTab[110368]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2478
					// _ = "end of CoverTab[110368]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2478
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2478
				// _ = "end of CoverTab[110361]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2479
			// _ = "end of CoverTab[110358]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2479
			_go_fuzz_dep_.CoverTab[110359]++
														return b, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2480
			// _ = "end of CoverTab[110359]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2481
	// _ = "end of CoverTab[110351]"
}

// makeMapMarshaler returns the sizer and marshaler for a map field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2484
// f is the pointer to the reflect data structure of the field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2486
func makeMapMarshaler(f *reflect.StructField) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2486
	_go_fuzz_dep_.CoverTab[110369]++

												t := f.Type
												keyType := t.Key()
												valType := t.Elem()
												tags := strings.Split(f.Tag.Get("protobuf"), ",")
												keyTags := strings.Split(f.Tag.Get("protobuf_key"), ",")
												valTags := strings.Split(f.Tag.Get("protobuf_val"), ",")
												stdOptions := false
												for _, t := range tags {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2495
		_go_fuzz_dep_.CoverTab[110372]++
													if strings.HasPrefix(t, "customtype=") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2496
			_go_fuzz_dep_.CoverTab[110376]++
														valTags = append(valTags, t)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2497
			// _ = "end of CoverTab[110376]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2498
			_go_fuzz_dep_.CoverTab[110377]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2498
			// _ = "end of CoverTab[110377]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2498
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2498
		// _ = "end of CoverTab[110372]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2498
		_go_fuzz_dep_.CoverTab[110373]++
													if t == "stdtime" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2499
			_go_fuzz_dep_.CoverTab[110378]++
														valTags = append(valTags, t)
														stdOptions = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2501
			// _ = "end of CoverTab[110378]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2502
			_go_fuzz_dep_.CoverTab[110379]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2502
			// _ = "end of CoverTab[110379]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2502
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2502
		// _ = "end of CoverTab[110373]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2502
		_go_fuzz_dep_.CoverTab[110374]++
													if t == "stdduration" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2503
			_go_fuzz_dep_.CoverTab[110380]++
														valTags = append(valTags, t)
														stdOptions = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2505
			// _ = "end of CoverTab[110380]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2506
			_go_fuzz_dep_.CoverTab[110381]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2506
			// _ = "end of CoverTab[110381]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2506
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2506
		// _ = "end of CoverTab[110374]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2506
		_go_fuzz_dep_.CoverTab[110375]++
													if t == "wktptr" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2507
			_go_fuzz_dep_.CoverTab[110382]++
														valTags = append(valTags, t)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2508
			// _ = "end of CoverTab[110382]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2509
			_go_fuzz_dep_.CoverTab[110383]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2509
			// _ = "end of CoverTab[110383]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2509
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2509
		// _ = "end of CoverTab[110375]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2510
	// _ = "end of CoverTab[110369]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2510
	_go_fuzz_dep_.CoverTab[110370]++
												keySizer, keyMarshaler := typeMarshaler(keyType, keyTags, false, false)
												valSizer, valMarshaler := typeMarshaler(valType, valTags, false, false)
												keyWireTag := 1<<3 | wiretype(keyTags[0])
												valWireTag := 2<<3 | wiretype(valTags[0])

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2521
	valIsPtr := valType.Kind() == reflect.Ptr

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2528
	valCachedSizer := valSizer
	if valIsPtr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2529
		_go_fuzz_dep_.CoverTab[110384]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2529
		return !stdOptions
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2529
		// _ = "end of CoverTab[110384]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2529
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2529
		_go_fuzz_dep_.CoverTab[110385]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2529
		return valType.Elem().Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2529
		// _ = "end of CoverTab[110385]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2529
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2529
		_go_fuzz_dep_.CoverTab[110386]++
													u := getMarshalInfo(valType.Elem())
													valCachedSizer = func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2531
			_go_fuzz_dep_.CoverTab[110387]++

														p := ptr.getPointer()
														if p.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2534
				_go_fuzz_dep_.CoverTab[110389]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2535
				// _ = "end of CoverTab[110389]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2536
				_go_fuzz_dep_.CoverTab[110390]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2536
				// _ = "end of CoverTab[110390]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2536
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2536
			// _ = "end of CoverTab[110387]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2536
			_go_fuzz_dep_.CoverTab[110388]++
														siz := u.cachedsize(p)
														return siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2538
			// _ = "end of CoverTab[110388]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2539
		// _ = "end of CoverTab[110386]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2540
		_go_fuzz_dep_.CoverTab[110391]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2540
		// _ = "end of CoverTab[110391]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2540
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2540
	// _ = "end of CoverTab[110370]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2540
	_go_fuzz_dep_.CoverTab[110371]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2541
			_go_fuzz_dep_.CoverTab[110392]++
														m := ptr.asPointerTo(t).Elem()
														n := 0
														for _, k := range m.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2544
				_go_fuzz_dep_.CoverTab[110394]++
															ki := k.Interface()
															vi := m.MapIndex(k).Interface()
															kaddr := toAddrPointer(&ki, false)
															vaddr := toAddrPointer(&vi, valIsPtr)
															siz := keySizer(kaddr, 1) + valSizer(vaddr, 1)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2550
				// _ = "end of CoverTab[110394]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2551
			// _ = "end of CoverTab[110392]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2551
			_go_fuzz_dep_.CoverTab[110393]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2552
			// _ = "end of CoverTab[110393]"
		},
		func(b []byte, ptr pointer, tag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2554
			_go_fuzz_dep_.CoverTab[110395]++
														m := ptr.asPointerTo(t).Elem()
														var err error
														keys := m.MapKeys()
														if len(keys) > 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2558
				_go_fuzz_dep_.CoverTab[110398]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2558
				return deterministic
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2558
				// _ = "end of CoverTab[110398]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2558
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2558
				_go_fuzz_dep_.CoverTab[110399]++
															sort.Sort(mapKeys(keys))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2559
				// _ = "end of CoverTab[110399]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2560
				_go_fuzz_dep_.CoverTab[110400]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2560
				// _ = "end of CoverTab[110400]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2560
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2560
			// _ = "end of CoverTab[110395]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2560
			_go_fuzz_dep_.CoverTab[110396]++

														var nerr nonFatal
														for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2563
				_go_fuzz_dep_.CoverTab[110401]++
															ki := k.Interface()
															vi := m.MapIndex(k).Interface()
															kaddr := toAddrPointer(&ki, false)
															vaddr := toAddrPointer(&vi, valIsPtr)
															b = appendVarint(b, tag)
															siz := keySizer(kaddr, 1) + valCachedSizer(vaddr, 1)
															b = appendVarint(b, uint64(siz))
															b, err = keyMarshaler(b, kaddr, keyWireTag, deterministic)
															if !nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2572
					_go_fuzz_dep_.CoverTab[110403]++
																return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2573
					// _ = "end of CoverTab[110403]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2574
					_go_fuzz_dep_.CoverTab[110404]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2574
					// _ = "end of CoverTab[110404]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2574
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2574
				// _ = "end of CoverTab[110401]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2574
				_go_fuzz_dep_.CoverTab[110402]++
															b, err = valMarshaler(b, vaddr, valWireTag, deterministic)
															if err != ErrNil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2576
					_go_fuzz_dep_.CoverTab[110405]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2576
					return !nerr.Merge(err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2576
					// _ = "end of CoverTab[110405]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2576
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2576
					_go_fuzz_dep_.CoverTab[110406]++
																return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2577
					// _ = "end of CoverTab[110406]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2578
					_go_fuzz_dep_.CoverTab[110407]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2578
					// _ = "end of CoverTab[110407]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2578
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2578
				// _ = "end of CoverTab[110402]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2579
			// _ = "end of CoverTab[110396]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2579
			_go_fuzz_dep_.CoverTab[110397]++
														return b, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2580
			// _ = "end of CoverTab[110397]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2581
	// _ = "end of CoverTab[110371]"
}

// makeOneOfMarshaler returns the sizer and marshaler for a oneof field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2584
// fi is the marshal info of the field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2584
// f is the pointer to the reflect data structure of the field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2587
func makeOneOfMarshaler(fi *marshalFieldInfo, f *reflect.StructField) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2587
	_go_fuzz_dep_.CoverTab[110408]++

												t := f.Type
												return func(ptr pointer, _ int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2590
			_go_fuzz_dep_.CoverTab[110409]++
														p := ptr.getInterfacePointer()
														if p.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2592
				_go_fuzz_dep_.CoverTab[110411]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2593
				// _ = "end of CoverTab[110411]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2594
				_go_fuzz_dep_.CoverTab[110412]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2594
				// _ = "end of CoverTab[110412]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2594
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2594
			// _ = "end of CoverTab[110409]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2594
			_go_fuzz_dep_.CoverTab[110410]++
														v := ptr.asPointerTo(t).Elem().Elem().Elem()
														telem := v.Type()
														e := fi.oneofElems[telem]
														return e.sizer(p, e.tagsize)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2598
			// _ = "end of CoverTab[110410]"
		},
		func(b []byte, ptr pointer, _ uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2600
			_go_fuzz_dep_.CoverTab[110413]++
														p := ptr.getInterfacePointer()
														if p.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2602
				_go_fuzz_dep_.CoverTab[110416]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2603
				// _ = "end of CoverTab[110416]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2604
				_go_fuzz_dep_.CoverTab[110417]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2604
				// _ = "end of CoverTab[110417]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2604
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2604
			// _ = "end of CoverTab[110413]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2604
			_go_fuzz_dep_.CoverTab[110414]++
														v := ptr.asPointerTo(t).Elem().Elem().Elem()
														telem := v.Type()
														if telem.Field(0).Type.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2607
				_go_fuzz_dep_.CoverTab[110418]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2607
				return p.getPointer().isNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2607
				// _ = "end of CoverTab[110418]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2607
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2607
				_go_fuzz_dep_.CoverTab[110419]++
															return b, errOneofHasNil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2608
				// _ = "end of CoverTab[110419]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2609
				_go_fuzz_dep_.CoverTab[110420]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2609
				// _ = "end of CoverTab[110420]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2609
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2609
			// _ = "end of CoverTab[110414]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2609
			_go_fuzz_dep_.CoverTab[110415]++
														e := fi.oneofElems[telem]
														return e.marshaler(b, p, e.wiretag, deterministic)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2611
			// _ = "end of CoverTab[110415]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2612
	// _ = "end of CoverTab[110408]"
}

// sizeExtensions computes the size of encoded data for a XXX_InternalExtensions field.
func (u *marshalInfo) sizeExtensions(ext *XXX_InternalExtensions) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2616
	_go_fuzz_dep_.CoverTab[110421]++
												m, mu := ext.extensionsRead()
												if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2618
		_go_fuzz_dep_.CoverTab[110424]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2619
		// _ = "end of CoverTab[110424]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2620
		_go_fuzz_dep_.CoverTab[110425]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2620
		// _ = "end of CoverTab[110425]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2620
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2620
	// _ = "end of CoverTab[110421]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2620
	_go_fuzz_dep_.CoverTab[110422]++
												mu.Lock()

												n := 0
												for _, e := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2624
		_go_fuzz_dep_.CoverTab[110426]++
													if e.value == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2625
			_go_fuzz_dep_.CoverTab[110428]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2625
			return e.desc == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2625
			// _ = "end of CoverTab[110428]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2625
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2625
			_go_fuzz_dep_.CoverTab[110429]++

														n += len(e.enc)
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2628
			// _ = "end of CoverTab[110429]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2629
			_go_fuzz_dep_.CoverTab[110430]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2629
			// _ = "end of CoverTab[110430]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2629
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2629
		// _ = "end of CoverTab[110426]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2629
		_go_fuzz_dep_.CoverTab[110427]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2634
		ei := u.getExtElemInfo(e.desc)
													v := e.value
													p := toAddrPointer(&v, ei.isptr)
													n += ei.sizer(p, ei.tagsize)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2637
		// _ = "end of CoverTab[110427]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2638
	// _ = "end of CoverTab[110422]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2638
	_go_fuzz_dep_.CoverTab[110423]++
												mu.Unlock()
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2640
	// _ = "end of CoverTab[110423]"
}

// appendExtensions marshals a XXX_InternalExtensions field to the end of byte slice b.
func (u *marshalInfo) appendExtensions(b []byte, ext *XXX_InternalExtensions, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2644
	_go_fuzz_dep_.CoverTab[110431]++
												m, mu := ext.extensionsRead()
												if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2646
		_go_fuzz_dep_.CoverTab[110436]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2647
		// _ = "end of CoverTab[110436]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2648
		_go_fuzz_dep_.CoverTab[110437]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2648
		// _ = "end of CoverTab[110437]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2648
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2648
	// _ = "end of CoverTab[110431]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2648
	_go_fuzz_dep_.CoverTab[110432]++
												mu.Lock()
												defer mu.Unlock()

												var err error
												var nerr nonFatal

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2657
	if len(m) <= 1 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2657
		_go_fuzz_dep_.CoverTab[110438]++
													for _, e := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2658
			_go_fuzz_dep_.CoverTab[110440]++
														if e.value == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2659
				_go_fuzz_dep_.CoverTab[110442]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2659
				return e.desc == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2659
				// _ = "end of CoverTab[110442]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2659
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2659
				_go_fuzz_dep_.CoverTab[110443]++

															b = append(b, e.enc...)
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2662
				// _ = "end of CoverTab[110443]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2663
				_go_fuzz_dep_.CoverTab[110444]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2663
				// _ = "end of CoverTab[110444]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2663
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2663
			// _ = "end of CoverTab[110440]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2663
			_go_fuzz_dep_.CoverTab[110441]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2669
			ei := u.getExtElemInfo(e.desc)
			v := e.value
			p := toAddrPointer(&v, ei.isptr)
			b, err = ei.marshaler(b, p, ei.wiretag, deterministic)
			if !nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2673
				_go_fuzz_dep_.CoverTab[110445]++
															return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2674
				// _ = "end of CoverTab[110445]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2675
				_go_fuzz_dep_.CoverTab[110446]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2675
				// _ = "end of CoverTab[110446]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2675
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2675
			// _ = "end of CoverTab[110441]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2676
		// _ = "end of CoverTab[110438]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2676
		_go_fuzz_dep_.CoverTab[110439]++
													return b, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2677
		// _ = "end of CoverTab[110439]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2678
		_go_fuzz_dep_.CoverTab[110447]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2678
		// _ = "end of CoverTab[110447]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2678
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2678
	// _ = "end of CoverTab[110432]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2678
	_go_fuzz_dep_.CoverTab[110433]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2682
	keys := make([]int, 0, len(m))
	for k := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2683
		_go_fuzz_dep_.CoverTab[110448]++
													keys = append(keys, int(k))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2684
		// _ = "end of CoverTab[110448]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2685
	// _ = "end of CoverTab[110433]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2685
	_go_fuzz_dep_.CoverTab[110434]++
												sort.Ints(keys)

												for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2688
		_go_fuzz_dep_.CoverTab[110449]++
													e := m[int32(k)]
													if e.value == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2690
			_go_fuzz_dep_.CoverTab[110451]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2690
			return e.desc == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2690
			// _ = "end of CoverTab[110451]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2690
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2690
			_go_fuzz_dep_.CoverTab[110452]++

														b = append(b, e.enc...)
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2693
			// _ = "end of CoverTab[110452]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2694
			_go_fuzz_dep_.CoverTab[110453]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2694
			// _ = "end of CoverTab[110453]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2694
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2694
		// _ = "end of CoverTab[110449]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2694
		_go_fuzz_dep_.CoverTab[110450]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2700
		ei := u.getExtElemInfo(e.desc)
		v := e.value
		p := toAddrPointer(&v, ei.isptr)
		b, err = ei.marshaler(b, p, ei.wiretag, deterministic)
		if !nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2704
			_go_fuzz_dep_.CoverTab[110454]++
														return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2705
			// _ = "end of CoverTab[110454]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2706
			_go_fuzz_dep_.CoverTab[110455]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2706
			// _ = "end of CoverTab[110455]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2706
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2706
		// _ = "end of CoverTab[110450]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2707
	// _ = "end of CoverTab[110434]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2707
	_go_fuzz_dep_.CoverTab[110435]++
												return b, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2708
	// _ = "end of CoverTab[110435]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2719
// sizeMessageSet computes the size of encoded data for a XXX_InternalExtensions field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2719
// in message set format (above).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2721
func (u *marshalInfo) sizeMessageSet(ext *XXX_InternalExtensions) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2721
	_go_fuzz_dep_.CoverTab[110456]++
												m, mu := ext.extensionsRead()
												if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2723
		_go_fuzz_dep_.CoverTab[110459]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2724
		// _ = "end of CoverTab[110459]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2725
		_go_fuzz_dep_.CoverTab[110460]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2725
		// _ = "end of CoverTab[110460]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2725
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2725
	// _ = "end of CoverTab[110456]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2725
	_go_fuzz_dep_.CoverTab[110457]++
												mu.Lock()

												n := 0
												for id, e := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2729
		_go_fuzz_dep_.CoverTab[110461]++
													n += 2
													n += SizeVarint(uint64(id)) + 1

													if e.value == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2733
			_go_fuzz_dep_.CoverTab[110463]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2733
			return e.desc == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2733
			// _ = "end of CoverTab[110463]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2733
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2733
			_go_fuzz_dep_.CoverTab[110464]++

														msgWithLen := skipVarint(e.enc)
														siz := len(msgWithLen)
														n += siz + 1
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2738
			// _ = "end of CoverTab[110464]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2739
			_go_fuzz_dep_.CoverTab[110465]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2739
			// _ = "end of CoverTab[110465]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2739
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2739
		// _ = "end of CoverTab[110461]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2739
		_go_fuzz_dep_.CoverTab[110462]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2745
		ei := u.getExtElemInfo(e.desc)
													v := e.value
													p := toAddrPointer(&v, ei.isptr)
													n += ei.sizer(p, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2748
		// _ = "end of CoverTab[110462]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2749
	// _ = "end of CoverTab[110457]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2749
	_go_fuzz_dep_.CoverTab[110458]++
												mu.Unlock()
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2751
	// _ = "end of CoverTab[110458]"
}

// appendMessageSet marshals a XXX_InternalExtensions field in message set format (above)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2754
// to the end of byte slice b.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2756
func (u *marshalInfo) appendMessageSet(b []byte, ext *XXX_InternalExtensions, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2756
	_go_fuzz_dep_.CoverTab[110466]++
												m, mu := ext.extensionsRead()
												if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2758
		_go_fuzz_dep_.CoverTab[110471]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2759
		// _ = "end of CoverTab[110471]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2760
		_go_fuzz_dep_.CoverTab[110472]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2760
		// _ = "end of CoverTab[110472]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2760
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2760
	// _ = "end of CoverTab[110466]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2760
	_go_fuzz_dep_.CoverTab[110467]++
												mu.Lock()
												defer mu.Unlock()

												var err error
												var nerr nonFatal

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2769
	if len(m) <= 1 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2769
		_go_fuzz_dep_.CoverTab[110473]++
													for id, e := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2770
			_go_fuzz_dep_.CoverTab[110475]++
														b = append(b, 1<<3|WireStartGroup)
														b = append(b, 2<<3|WireVarint)
														b = appendVarint(b, uint64(id))

														if e.value == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2775
				_go_fuzz_dep_.CoverTab[110478]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2775
				return e.desc == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2775
				// _ = "end of CoverTab[110478]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2775
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2775
				_go_fuzz_dep_.CoverTab[110479]++

															msgWithLen := skipVarint(e.enc)
															b = append(b, 3<<3|WireBytes)
															b = append(b, msgWithLen...)
															b = append(b, 1<<3|WireEndGroup)
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2781
				// _ = "end of CoverTab[110479]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2782
				_go_fuzz_dep_.CoverTab[110480]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2782
				// _ = "end of CoverTab[110480]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2782
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2782
			// _ = "end of CoverTab[110475]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2782
			_go_fuzz_dep_.CoverTab[110476]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2788
			ei := u.getExtElemInfo(e.desc)
			v := e.value
			p := toAddrPointer(&v, ei.isptr)
			b, err = ei.marshaler(b, p, 3<<3|WireBytes, deterministic)
			if !nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2792
				_go_fuzz_dep_.CoverTab[110481]++
															return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2793
				// _ = "end of CoverTab[110481]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2794
				_go_fuzz_dep_.CoverTab[110482]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2794
				// _ = "end of CoverTab[110482]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2794
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2794
			// _ = "end of CoverTab[110476]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2794
			_go_fuzz_dep_.CoverTab[110477]++
														b = append(b, 1<<3|WireEndGroup)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2795
			// _ = "end of CoverTab[110477]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2796
		// _ = "end of CoverTab[110473]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2796
		_go_fuzz_dep_.CoverTab[110474]++
													return b, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2797
		// _ = "end of CoverTab[110474]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2798
		_go_fuzz_dep_.CoverTab[110483]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2798
		// _ = "end of CoverTab[110483]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2798
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2798
	// _ = "end of CoverTab[110467]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2798
	_go_fuzz_dep_.CoverTab[110468]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2801
	keys := make([]int, 0, len(m))
	for k := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2802
		_go_fuzz_dep_.CoverTab[110484]++
													keys = append(keys, int(k))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2803
		// _ = "end of CoverTab[110484]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2804
	// _ = "end of CoverTab[110468]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2804
	_go_fuzz_dep_.CoverTab[110469]++
												sort.Ints(keys)

												for _, id := range keys {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2807
		_go_fuzz_dep_.CoverTab[110485]++
													e := m[int32(id)]
													b = append(b, 1<<3|WireStartGroup)
													b = append(b, 2<<3|WireVarint)
													b = appendVarint(b, uint64(id))

													if e.value == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2813
			_go_fuzz_dep_.CoverTab[110487]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2813
			return e.desc == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2813
			// _ = "end of CoverTab[110487]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2813
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2813
			_go_fuzz_dep_.CoverTab[110488]++

														msgWithLen := skipVarint(e.enc)
														b = append(b, 3<<3|WireBytes)
														b = append(b, msgWithLen...)
														b = append(b, 1<<3|WireEndGroup)
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2819
			// _ = "end of CoverTab[110488]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2820
			_go_fuzz_dep_.CoverTab[110489]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2820
			// _ = "end of CoverTab[110489]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2820
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2820
		// _ = "end of CoverTab[110485]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2820
		_go_fuzz_dep_.CoverTab[110486]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2826
		ei := u.getExtElemInfo(e.desc)
		v := e.value
		p := toAddrPointer(&v, ei.isptr)
		b, err = ei.marshaler(b, p, 3<<3|WireBytes, deterministic)
		b = append(b, 1<<3|WireEndGroup)
		if !nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2831
			_go_fuzz_dep_.CoverTab[110490]++
														return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2832
			// _ = "end of CoverTab[110490]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2833
			_go_fuzz_dep_.CoverTab[110491]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2833
			// _ = "end of CoverTab[110491]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2833
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2833
		// _ = "end of CoverTab[110486]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2834
	// _ = "end of CoverTab[110469]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2834
	_go_fuzz_dep_.CoverTab[110470]++
												return b, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2835
	// _ = "end of CoverTab[110470]"
}

// sizeV1Extensions computes the size of encoded data for a V1-API extension field.
func (u *marshalInfo) sizeV1Extensions(m map[int32]Extension) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2839
	_go_fuzz_dep_.CoverTab[110492]++
												if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2840
		_go_fuzz_dep_.CoverTab[110495]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2841
		// _ = "end of CoverTab[110495]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2842
		_go_fuzz_dep_.CoverTab[110496]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2842
		// _ = "end of CoverTab[110496]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2842
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2842
	// _ = "end of CoverTab[110492]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2842
	_go_fuzz_dep_.CoverTab[110493]++

												n := 0
												for _, e := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2845
		_go_fuzz_dep_.CoverTab[110497]++
													if e.value == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2846
			_go_fuzz_dep_.CoverTab[110499]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2846
			return e.desc == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2846
			// _ = "end of CoverTab[110499]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2846
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2846
			_go_fuzz_dep_.CoverTab[110500]++

														n += len(e.enc)
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2849
			// _ = "end of CoverTab[110500]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2850
			_go_fuzz_dep_.CoverTab[110501]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2850
			// _ = "end of CoverTab[110501]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2850
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2850
		// _ = "end of CoverTab[110497]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2850
		_go_fuzz_dep_.CoverTab[110498]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2856
		ei := u.getExtElemInfo(e.desc)
													v := e.value
													p := toAddrPointer(&v, ei.isptr)
													n += ei.sizer(p, ei.tagsize)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2859
		// _ = "end of CoverTab[110498]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2860
	// _ = "end of CoverTab[110493]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2860
	_go_fuzz_dep_.CoverTab[110494]++
												return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2861
	// _ = "end of CoverTab[110494]"
}

// appendV1Extensions marshals a V1-API extension field to the end of byte slice b.
func (u *marshalInfo) appendV1Extensions(b []byte, m map[int32]Extension, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2865
	_go_fuzz_dep_.CoverTab[110502]++
												if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2866
		_go_fuzz_dep_.CoverTab[110506]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2867
		// _ = "end of CoverTab[110506]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2868
		_go_fuzz_dep_.CoverTab[110507]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2868
		// _ = "end of CoverTab[110507]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2868
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2868
	// _ = "end of CoverTab[110502]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2868
	_go_fuzz_dep_.CoverTab[110503]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2871
	keys := make([]int, 0, len(m))
	for k := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2872
		_go_fuzz_dep_.CoverTab[110508]++
													keys = append(keys, int(k))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2873
		// _ = "end of CoverTab[110508]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2874
	// _ = "end of CoverTab[110503]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2874
	_go_fuzz_dep_.CoverTab[110504]++
												sort.Ints(keys)

												var err error
												var nerr nonFatal
												for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2879
		_go_fuzz_dep_.CoverTab[110509]++
													e := m[int32(k)]
													if e.value == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2881
			_go_fuzz_dep_.CoverTab[110511]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2881
			return e.desc == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2881
			// _ = "end of CoverTab[110511]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2881
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2881
			_go_fuzz_dep_.CoverTab[110512]++

														b = append(b, e.enc...)
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2884
			// _ = "end of CoverTab[110512]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2885
			_go_fuzz_dep_.CoverTab[110513]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2885
			// _ = "end of CoverTab[110513]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2885
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2885
		// _ = "end of CoverTab[110509]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2885
		_go_fuzz_dep_.CoverTab[110510]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2891
		ei := u.getExtElemInfo(e.desc)
		v := e.value
		p := toAddrPointer(&v, ei.isptr)
		b, err = ei.marshaler(b, p, ei.wiretag, deterministic)
		if !nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2895
			_go_fuzz_dep_.CoverTab[110514]++
														return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2896
			// _ = "end of CoverTab[110514]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2897
			_go_fuzz_dep_.CoverTab[110515]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2897
			// _ = "end of CoverTab[110515]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2897
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2897
		// _ = "end of CoverTab[110510]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2898
	// _ = "end of CoverTab[110504]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2898
	_go_fuzz_dep_.CoverTab[110505]++
												return b, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2899
	// _ = "end of CoverTab[110505]"
}

// newMarshaler is the interface representing objects that can marshal themselves.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2902
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2902
// This exists to support protoc-gen-go generated messages.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2902
// The proto package will stop type-asserting to this interface in the future.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2902
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2902
// DO NOT DEPEND ON THIS.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2908
type newMarshaler interface {
	XXX_Size() int
	XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
}

// Size returns the encoded size of a protocol buffer message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2913
// This is the main entry point.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2915
func Size(pb Message) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2915
	_go_fuzz_dep_.CoverTab[110516]++
												if m, ok := pb.(newMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2916
		_go_fuzz_dep_.CoverTab[110520]++
													return m.XXX_Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2917
		// _ = "end of CoverTab[110520]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2918
		_go_fuzz_dep_.CoverTab[110521]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2918
		// _ = "end of CoverTab[110521]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2918
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2918
	// _ = "end of CoverTab[110516]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2918
	_go_fuzz_dep_.CoverTab[110517]++
												if m, ok := pb.(Marshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2919
		_go_fuzz_dep_.CoverTab[110522]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2922
		b, _ := m.Marshal()
													return len(b)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2923
		// _ = "end of CoverTab[110522]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2924
		_go_fuzz_dep_.CoverTab[110523]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2924
		// _ = "end of CoverTab[110523]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2924
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2924
	// _ = "end of CoverTab[110517]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2924
	_go_fuzz_dep_.CoverTab[110518]++

												if pb == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2926
		_go_fuzz_dep_.CoverTab[110524]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2927
		// _ = "end of CoverTab[110524]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2928
		_go_fuzz_dep_.CoverTab[110525]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2928
		// _ = "end of CoverTab[110525]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2928
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2928
	// _ = "end of CoverTab[110518]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2928
	_go_fuzz_dep_.CoverTab[110519]++
												var info InternalMessageInfo
												return info.Size(pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2930
	// _ = "end of CoverTab[110519]"
}

// Marshal takes a protocol buffer message
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2933
// and encodes it into the wire format, returning the data.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2933
// This is the main entry point.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2936
func Marshal(pb Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2936
	_go_fuzz_dep_.CoverTab[110526]++
												if m, ok := pb.(newMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2937
		_go_fuzz_dep_.CoverTab[110530]++
													siz := m.XXX_Size()
													b := make([]byte, 0, siz)
													return m.XXX_Marshal(b, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2940
		// _ = "end of CoverTab[110530]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2941
		_go_fuzz_dep_.CoverTab[110531]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2941
		// _ = "end of CoverTab[110531]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2941
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2941
	// _ = "end of CoverTab[110526]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2941
	_go_fuzz_dep_.CoverTab[110527]++
												if m, ok := pb.(Marshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2942
		_go_fuzz_dep_.CoverTab[110532]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2945
		return m.Marshal()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2945
		// _ = "end of CoverTab[110532]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2946
		_go_fuzz_dep_.CoverTab[110533]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2946
		// _ = "end of CoverTab[110533]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2946
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2946
	// _ = "end of CoverTab[110527]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2946
	_go_fuzz_dep_.CoverTab[110528]++

												if pb == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2948
		_go_fuzz_dep_.CoverTab[110534]++
													return nil, ErrNil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2949
		// _ = "end of CoverTab[110534]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2950
		_go_fuzz_dep_.CoverTab[110535]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2950
		// _ = "end of CoverTab[110535]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2950
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2950
	// _ = "end of CoverTab[110528]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2950
	_go_fuzz_dep_.CoverTab[110529]++
												var info InternalMessageInfo
												siz := info.Size(pb)
												b := make([]byte, 0, siz)
												return info.Marshal(b, pb, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2954
	// _ = "end of CoverTab[110529]"
}

// Marshal takes a protocol buffer message
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2957
// and encodes it into the wire format, writing the result to the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2957
// Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2957
// This is an alternative entry point. It is not necessary to use
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2957
// a Buffer for most applications.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2962
func (p *Buffer) Marshal(pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2962
	_go_fuzz_dep_.CoverTab[110536]++
												var err error
												if p.deterministic {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2964
		_go_fuzz_dep_.CoverTab[110541]++
													if _, ok := pb.(Marshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2965
			_go_fuzz_dep_.CoverTab[110542]++
														return fmt.Errorf("proto: deterministic not supported by the Marshal method of %T", pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2966
			// _ = "end of CoverTab[110542]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2967
			_go_fuzz_dep_.CoverTab[110543]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2967
			// _ = "end of CoverTab[110543]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2967
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2967
		// _ = "end of CoverTab[110541]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2968
		_go_fuzz_dep_.CoverTab[110544]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2968
		// _ = "end of CoverTab[110544]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2968
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2968
	// _ = "end of CoverTab[110536]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2968
	_go_fuzz_dep_.CoverTab[110537]++
												if m, ok := pb.(newMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2969
		_go_fuzz_dep_.CoverTab[110545]++
													siz := m.XXX_Size()
													p.grow(siz)
													pp := p.buf[len(p.buf) : len(p.buf) : len(p.buf)+siz]
													pp, err = m.XXX_Marshal(pp, p.deterministic)
													p.buf = append(p.buf, pp...)
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2975
		// _ = "end of CoverTab[110545]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2976
		_go_fuzz_dep_.CoverTab[110546]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2976
		// _ = "end of CoverTab[110546]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2976
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2976
	// _ = "end of CoverTab[110537]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2976
	_go_fuzz_dep_.CoverTab[110538]++
												if m, ok := pb.(Marshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2977
		_go_fuzz_dep_.CoverTab[110547]++
		// If the message can marshal itself, let it do it, for compatibility.
													// NOTE: This is not efficient.
													var b []byte
													b, err = m.Marshal()
													p.buf = append(p.buf, b...)
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2983
		// _ = "end of CoverTab[110547]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2984
		_go_fuzz_dep_.CoverTab[110548]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2984
		// _ = "end of CoverTab[110548]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2984
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2984
	// _ = "end of CoverTab[110538]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2984
	_go_fuzz_dep_.CoverTab[110539]++

												if pb == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2986
		_go_fuzz_dep_.CoverTab[110549]++
													return ErrNil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2987
		// _ = "end of CoverTab[110549]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2988
		_go_fuzz_dep_.CoverTab[110550]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2988
		// _ = "end of CoverTab[110550]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2988
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2988
	// _ = "end of CoverTab[110539]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2988
	_go_fuzz_dep_.CoverTab[110540]++
												var info InternalMessageInfo
												siz := info.Size(pb)
												p.grow(siz)
												p.buf, err = info.Marshal(p.buf, pb, p.deterministic)
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2993
	// _ = "end of CoverTab[110540]"
}

// grow grows the buffer's capacity, if necessary, to guarantee space for
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2996
// another n bytes. After grow(n), at least n bytes can be written to the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2996
// buffer without another allocation.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2999
func (p *Buffer) grow(n int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:2999
	_go_fuzz_dep_.CoverTab[110551]++
												need := len(p.buf) + n
												if need <= cap(p.buf) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3001
		_go_fuzz_dep_.CoverTab[110554]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3002
		// _ = "end of CoverTab[110554]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3003
		_go_fuzz_dep_.CoverTab[110555]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3003
		// _ = "end of CoverTab[110555]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3003
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3003
	// _ = "end of CoverTab[110551]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3003
	_go_fuzz_dep_.CoverTab[110552]++
												newCap := len(p.buf) * 2
												if newCap < need {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3005
		_go_fuzz_dep_.CoverTab[110556]++
													newCap = need
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3006
		// _ = "end of CoverTab[110556]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3007
		_go_fuzz_dep_.CoverTab[110557]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3007
		// _ = "end of CoverTab[110557]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3007
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3007
	// _ = "end of CoverTab[110552]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3007
	_go_fuzz_dep_.CoverTab[110553]++
												p.buf = append(make([]byte, 0, newCap), p.buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3008
	// _ = "end of CoverTab[110553]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3009
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal.go:3009
var _ = _go_fuzz_dep_.CoverTab
