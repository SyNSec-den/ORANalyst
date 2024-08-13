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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:32
)

import (
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"unicode/utf8"
)

// Unmarshal is the entry point from the generated .pb.go files.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:47
// This function is not intended to be used by non-generated code.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:47
// This function is not subject to any compatibility guarantee.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:47
// msg contains a pointer to a protocol buffer struct.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:47
// b is the data to be unmarshaled into the protocol buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:47
// a is a pointer to a place to store cached unmarshal information.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:53
func (a *InternalMessageInfo) Unmarshal(msg Message, b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:53
	_go_fuzz_dep_.CoverTab[111005]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:56
	u := atomicLoadUnmarshalInfo(&a.unmarshal)
	if u == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:57
		_go_fuzz_dep_.CoverTab[111007]++

													u = getUnmarshalInfo(reflect.TypeOf(msg).Elem())
													atomicStoreUnmarshalInfo(&a.unmarshal, u)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:60
		// _ = "end of CoverTab[111007]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:61
		_go_fuzz_dep_.CoverTab[111008]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:61
		// _ = "end of CoverTab[111008]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:61
	// _ = "end of CoverTab[111005]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:61
	_go_fuzz_dep_.CoverTab[111006]++

												err := u.unmarshal(toPointer(&msg), b)
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:64
	// _ = "end of CoverTab[111006]"
}

type unmarshalInfo struct {
	typ	reflect.Type	// type of the protobuf struct

	// 0 = only typ field is initialized
	// 1 = completely initialized
	initialized	int32
	lock		sync.Mutex			// prevents double initialization
	dense		[]unmarshalFieldInfo		// fields indexed by tag #
	sparse		map[uint64]unmarshalFieldInfo	// fields indexed by tag #
	reqFields	[]string			// names of required fields
	reqMask		uint64				// 1<<len(reqFields)-1
	unrecognized	field				// offset of []byte to put unrecognized data (or invalidField if we should throw it away)
	extensions	field				// offset of extensions field (of type proto.XXX_InternalExtensions), or invalidField if it does not exist
	oldExtensions	field				// offset of old-form extensions field (of type map[int]Extension)
	extensionRanges	[]ExtensionRange		// if non-nil, implies extensions field is valid
	isMessageSet	bool				// if true, implies extensions field is valid

	bytesExtensions	field	// offset of XXX_extensions with type []byte
}

// An unmarshaler takes a stream of bytes and a pointer to a field of a message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:87
// It decodes the field, stores it at f, and returns the unused bytes.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:87
// w is the wire encoding.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:87
// b is the data after the tag and wire encoding have been read.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:91
type unmarshaler func(b []byte, f pointer, w int) ([]byte, error)

type unmarshalFieldInfo struct {
	// location of the field in the proto message structure.
	field	field

	// function to unmarshal the data for the field.
	unmarshal	unmarshaler

	// if a required field, contains a single set bit at this field's index in the required field list.
	reqMask	uint64

	name	string	// name of the field, for error reporting
}

var (
	unmarshalInfoMap	= map[reflect.Type]*unmarshalInfo{}
	unmarshalInfoLock	sync.Mutex
)

// getUnmarshalInfo returns the data structure which can be
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:111
// subsequently used to unmarshal a message of the given type.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:111
// t is the type of the message (note: not pointer to message).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:114
func getUnmarshalInfo(t reflect.Type) *unmarshalInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:114
	_go_fuzz_dep_.CoverTab[111009]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:119
	unmarshalInfoLock.Lock()
	defer unmarshalInfoLock.Unlock()
	u := unmarshalInfoMap[t]
	if u == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:122
		_go_fuzz_dep_.CoverTab[111011]++
													u = &unmarshalInfo{typ: t}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:126
		unmarshalInfoMap[t] = u
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:126
		// _ = "end of CoverTab[111011]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:127
		_go_fuzz_dep_.CoverTab[111012]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:127
		// _ = "end of CoverTab[111012]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:127
	// _ = "end of CoverTab[111009]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:127
	_go_fuzz_dep_.CoverTab[111010]++
												return u
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:128
	// _ = "end of CoverTab[111010]"
}

// unmarshal does the main work of unmarshaling a message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:131
// u provides type information used to unmarshal the message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:131
// m is a pointer to a protocol buffer message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:131
// b is a byte stream to unmarshal into m.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:131
// This is top routine used when recursively unmarshaling submessages.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:136
func (u *unmarshalInfo) unmarshal(m pointer, b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:136
	_go_fuzz_dep_.CoverTab[111013]++
												if atomic.LoadInt32(&u.initialized) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:137
		_go_fuzz_dep_.CoverTab[111018]++
													u.computeUnmarshalInfo()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:138
		// _ = "end of CoverTab[111018]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:139
		_go_fuzz_dep_.CoverTab[111019]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:139
		// _ = "end of CoverTab[111019]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:139
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:139
	// _ = "end of CoverTab[111013]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:139
	_go_fuzz_dep_.CoverTab[111014]++
												if u.isMessageSet {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:140
		_go_fuzz_dep_.CoverTab[111020]++
													return unmarshalMessageSet(b, m.offset(u.extensions).toExtensions())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:141
		// _ = "end of CoverTab[111020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:142
		_go_fuzz_dep_.CoverTab[111021]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:142
		// _ = "end of CoverTab[111021]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:142
	// _ = "end of CoverTab[111014]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:142
	_go_fuzz_dep_.CoverTab[111015]++
												var reqMask uint64	// bitmask of required fields we've seen.
												var errLater error
												for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:145
		_go_fuzz_dep_.CoverTab[111022]++
		// Read tag and wire type.
		// Special case 1 and 2 byte varints.
		var x uint64
		if b[0] < 128 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:149
			_go_fuzz_dep_.CoverTab[111029]++
														x = uint64(b[0])
														b = b[1:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:151
			// _ = "end of CoverTab[111029]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:152
			_go_fuzz_dep_.CoverTab[111030]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:152
			if len(b) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:152
				_go_fuzz_dep_.CoverTab[111031]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:152
				return b[1] < 128
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:152
				// _ = "end of CoverTab[111031]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:152
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:152
				_go_fuzz_dep_.CoverTab[111032]++
															x = uint64(b[0]&0x7f) + uint64(b[1])<<7
															b = b[2:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:154
				// _ = "end of CoverTab[111032]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:155
				_go_fuzz_dep_.CoverTab[111033]++
															var n int
															x, n = decodeVarint(b)
															if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:158
					_go_fuzz_dep_.CoverTab[111035]++
																return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:159
					// _ = "end of CoverTab[111035]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:160
					_go_fuzz_dep_.CoverTab[111036]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:160
					// _ = "end of CoverTab[111036]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:160
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:160
				// _ = "end of CoverTab[111033]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:160
				_go_fuzz_dep_.CoverTab[111034]++
															b = b[n:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:161
				// _ = "end of CoverTab[111034]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:162
			// _ = "end of CoverTab[111030]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:162
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:162
		// _ = "end of CoverTab[111022]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:162
		_go_fuzz_dep_.CoverTab[111023]++
													tag := x >> 3
													wire := int(x) & 7

		// Dispatch on the tag to one of the unmarshal* functions below.
		var f unmarshalFieldInfo
		if tag < uint64(len(u.dense)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:168
			_go_fuzz_dep_.CoverTab[111037]++
														f = u.dense[tag]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:169
			// _ = "end of CoverTab[111037]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:170
			_go_fuzz_dep_.CoverTab[111038]++
														f = u.sparse[tag]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:171
			// _ = "end of CoverTab[111038]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:172
		// _ = "end of CoverTab[111023]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:172
		_go_fuzz_dep_.CoverTab[111024]++
													if fn := f.unmarshal; fn != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:173
			_go_fuzz_dep_.CoverTab[111039]++
														var err error
														b, err = fn(b, m.offset(f.field), wire)
														if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:176
				_go_fuzz_dep_.CoverTab[111042]++
															reqMask |= f.reqMask
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:178
				// _ = "end of CoverTab[111042]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:179
				_go_fuzz_dep_.CoverTab[111043]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:179
				// _ = "end of CoverTab[111043]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:179
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:179
			// _ = "end of CoverTab[111039]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:179
			_go_fuzz_dep_.CoverTab[111040]++
														if r, ok := err.(*RequiredNotSetError); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:180
				_go_fuzz_dep_.CoverTab[111044]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:183
				if errLater == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:183
					_go_fuzz_dep_.CoverTab[111046]++
																errLater = r
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:184
					// _ = "end of CoverTab[111046]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:185
					_go_fuzz_dep_.CoverTab[111047]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:185
					// _ = "end of CoverTab[111047]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:185
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:185
				// _ = "end of CoverTab[111044]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:185
				_go_fuzz_dep_.CoverTab[111045]++
															reqMask |= f.reqMask
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:187
				// _ = "end of CoverTab[111045]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:188
				_go_fuzz_dep_.CoverTab[111048]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:188
				// _ = "end of CoverTab[111048]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:188
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:188
			// _ = "end of CoverTab[111040]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:188
			_go_fuzz_dep_.CoverTab[111041]++
														if err != errInternalBadWireType {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:189
				_go_fuzz_dep_.CoverTab[111049]++
															if err == errInvalidUTF8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:190
					_go_fuzz_dep_.CoverTab[111051]++
																if errLater == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:191
						_go_fuzz_dep_.CoverTab[111053]++
																	fullName := revProtoTypes[reflect.PtrTo(u.typ)] + "." + f.name
																	errLater = &invalidUTF8Error{fullName}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:193
						// _ = "end of CoverTab[111053]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:194
						_go_fuzz_dep_.CoverTab[111054]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:194
						// _ = "end of CoverTab[111054]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:194
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:194
					// _ = "end of CoverTab[111051]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:194
					_go_fuzz_dep_.CoverTab[111052]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:195
					// _ = "end of CoverTab[111052]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:196
					_go_fuzz_dep_.CoverTab[111055]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:196
					// _ = "end of CoverTab[111055]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:196
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:196
				// _ = "end of CoverTab[111049]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:196
				_go_fuzz_dep_.CoverTab[111050]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:197
				// _ = "end of CoverTab[111050]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:198
				_go_fuzz_dep_.CoverTab[111056]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:198
				// _ = "end of CoverTab[111056]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:198
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:198
			// _ = "end of CoverTab[111041]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:200
			_go_fuzz_dep_.CoverTab[111057]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:200
			// _ = "end of CoverTab[111057]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:200
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:200
		// _ = "end of CoverTab[111024]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:200
		_go_fuzz_dep_.CoverTab[111025]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:203
		if !u.unrecognized.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:203
			_go_fuzz_dep_.CoverTab[111058]++
			// Don't keep unrecognized data; just skip it.
			var err error
			b, err = skipField(b, wire)
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:207
				_go_fuzz_dep_.CoverTab[111060]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:208
				// _ = "end of CoverTab[111060]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:209
				_go_fuzz_dep_.CoverTab[111061]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:209
				// _ = "end of CoverTab[111061]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:209
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:209
			// _ = "end of CoverTab[111058]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:209
			_go_fuzz_dep_.CoverTab[111059]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:210
			// _ = "end of CoverTab[111059]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:211
			_go_fuzz_dep_.CoverTab[111062]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:211
			// _ = "end of CoverTab[111062]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:211
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:211
		// _ = "end of CoverTab[111025]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:211
		_go_fuzz_dep_.CoverTab[111026]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:214
		z := m.offset(u.unrecognized).toBytes()
		var emap map[int32]Extension
		var e Extension
		for _, r := range u.extensionRanges {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:217
			_go_fuzz_dep_.CoverTab[111063]++
														if uint64(r.Start) <= tag && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:218
				_go_fuzz_dep_.CoverTab[111064]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:218
				return tag <= uint64(r.End)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:218
				// _ = "end of CoverTab[111064]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:218
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:218
				_go_fuzz_dep_.CoverTab[111065]++
															if u.extensions.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:219
					_go_fuzz_dep_.CoverTab[111069]++
																mp := m.offset(u.extensions).toExtensions()
																emap = mp.extensionsWrite()
																e = emap[int32(tag)]
																z = &e.enc
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:224
					// _ = "end of CoverTab[111069]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:225
					_go_fuzz_dep_.CoverTab[111070]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:225
					// _ = "end of CoverTab[111070]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:225
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:225
				// _ = "end of CoverTab[111065]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:225
				_go_fuzz_dep_.CoverTab[111066]++
															if u.oldExtensions.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:226
					_go_fuzz_dep_.CoverTab[111071]++
																p := m.offset(u.oldExtensions).toOldExtensions()
																emap = *p
																if emap == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:229
						_go_fuzz_dep_.CoverTab[111073]++
																	emap = map[int32]Extension{}
																	*p = emap
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:231
						// _ = "end of CoverTab[111073]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:232
						_go_fuzz_dep_.CoverTab[111074]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:232
						// _ = "end of CoverTab[111074]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:232
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:232
					// _ = "end of CoverTab[111071]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:232
					_go_fuzz_dep_.CoverTab[111072]++
																e = emap[int32(tag)]
																z = &e.enc
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:235
					// _ = "end of CoverTab[111072]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:236
					_go_fuzz_dep_.CoverTab[111075]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:236
					// _ = "end of CoverTab[111075]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:236
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:236
				// _ = "end of CoverTab[111066]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:236
				_go_fuzz_dep_.CoverTab[111067]++
															if u.bytesExtensions.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:237
					_go_fuzz_dep_.CoverTab[111076]++
																z = m.offset(u.bytesExtensions).toBytes()
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:239
					// _ = "end of CoverTab[111076]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:240
					_go_fuzz_dep_.CoverTab[111077]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:240
					// _ = "end of CoverTab[111077]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:240
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:240
				// _ = "end of CoverTab[111067]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:240
				_go_fuzz_dep_.CoverTab[111068]++
															panic("no extensions field available")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:241
				// _ = "end of CoverTab[111068]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:242
				_go_fuzz_dep_.CoverTab[111078]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:242
				// _ = "end of CoverTab[111078]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:242
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:242
			// _ = "end of CoverTab[111063]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:243
		// _ = "end of CoverTab[111026]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:243
		_go_fuzz_dep_.CoverTab[111027]++
		// Use wire type to skip data.
		var err error
		b0 := b
		b, err = skipField(b, wire)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:248
			_go_fuzz_dep_.CoverTab[111079]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:249
			// _ = "end of CoverTab[111079]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:250
			_go_fuzz_dep_.CoverTab[111080]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:250
			// _ = "end of CoverTab[111080]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:250
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:250
		// _ = "end of CoverTab[111027]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:250
		_go_fuzz_dep_.CoverTab[111028]++
													*z = encodeVarint(*z, tag<<3|uint64(wire))
													*z = append(*z, b0[:len(b0)-len(b)]...)

													if emap != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:254
			_go_fuzz_dep_.CoverTab[111081]++
														emap[int32(tag)] = e
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:255
			// _ = "end of CoverTab[111081]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:256
			_go_fuzz_dep_.CoverTab[111082]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:256
			// _ = "end of CoverTab[111082]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:256
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:256
		// _ = "end of CoverTab[111028]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:257
	// _ = "end of CoverTab[111015]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:257
	_go_fuzz_dep_.CoverTab[111016]++
												if reqMask != u.reqMask && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:258
		_go_fuzz_dep_.CoverTab[111083]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:258
		return errLater == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:258
		// _ = "end of CoverTab[111083]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:258
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:258
		_go_fuzz_dep_.CoverTab[111084]++

													for _, n := range u.reqFields {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:260
			_go_fuzz_dep_.CoverTab[111085]++
														if reqMask&1 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:261
				_go_fuzz_dep_.CoverTab[111087]++
															errLater = &RequiredNotSetError{n}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:262
				// _ = "end of CoverTab[111087]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:263
				_go_fuzz_dep_.CoverTab[111088]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:263
				// _ = "end of CoverTab[111088]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:263
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:263
			// _ = "end of CoverTab[111085]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:263
			_go_fuzz_dep_.CoverTab[111086]++
														reqMask >>= 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:264
			// _ = "end of CoverTab[111086]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:265
		// _ = "end of CoverTab[111084]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:266
		_go_fuzz_dep_.CoverTab[111089]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:266
		// _ = "end of CoverTab[111089]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:266
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:266
	// _ = "end of CoverTab[111016]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:266
	_go_fuzz_dep_.CoverTab[111017]++
												return errLater
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:267
	// _ = "end of CoverTab[111017]"
}

// computeUnmarshalInfo fills in u with information for use
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:270
// in unmarshaling protocol buffers of type u.typ.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:272
func (u *unmarshalInfo) computeUnmarshalInfo() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:272
	_go_fuzz_dep_.CoverTab[111090]++
												u.lock.Lock()
												defer u.lock.Unlock()
												if u.initialized != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:275
		_go_fuzz_dep_.CoverTab[111096]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:276
		// _ = "end of CoverTab[111096]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:277
		_go_fuzz_dep_.CoverTab[111097]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:277
		// _ = "end of CoverTab[111097]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:277
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:277
	// _ = "end of CoverTab[111090]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:277
	_go_fuzz_dep_.CoverTab[111091]++
												t := u.typ
												n := t.NumField()

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:283
	u.unrecognized = invalidField
	u.extensions = invalidField
	u.oldExtensions = invalidField
	u.bytesExtensions = invalidField

	// List of the generated type and offset for each oneof field.
	type oneofField struct {
		ityp	reflect.Type	// interface type of oneof field
		field	field		// offset in containing message
	}
	var oneofFields []oneofField

	for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:295
		_go_fuzz_dep_.CoverTab[111098]++
													f := t.Field(i)
													if f.Name == "XXX_unrecognized" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:297
			_go_fuzz_dep_.CoverTab[111108]++

														if f.Type != reflect.TypeOf(([]byte)(nil)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:299
				_go_fuzz_dep_.CoverTab[111110]++
															panic("bad type for XXX_unrecognized field: " + f.Type.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:300
				// _ = "end of CoverTab[111110]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:301
				_go_fuzz_dep_.CoverTab[111111]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:301
				// _ = "end of CoverTab[111111]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:301
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:301
			// _ = "end of CoverTab[111108]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:301
			_go_fuzz_dep_.CoverTab[111109]++
														u.unrecognized = toField(&f)
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:303
			// _ = "end of CoverTab[111109]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:304
			_go_fuzz_dep_.CoverTab[111112]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:304
			// _ = "end of CoverTab[111112]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:304
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:304
		// _ = "end of CoverTab[111098]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:304
		_go_fuzz_dep_.CoverTab[111099]++
													if f.Name == "XXX_InternalExtensions" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:305
			_go_fuzz_dep_.CoverTab[111113]++

														if f.Type != reflect.TypeOf(XXX_InternalExtensions{}) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:307
				_go_fuzz_dep_.CoverTab[111116]++
															panic("bad type for XXX_InternalExtensions field: " + f.Type.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:308
				// _ = "end of CoverTab[111116]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:309
				_go_fuzz_dep_.CoverTab[111117]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:309
				// _ = "end of CoverTab[111117]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:309
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:309
			// _ = "end of CoverTab[111113]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:309
			_go_fuzz_dep_.CoverTab[111114]++
														u.extensions = toField(&f)
														if f.Tag.Get("protobuf_messageset") == "1" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:311
				_go_fuzz_dep_.CoverTab[111118]++
															u.isMessageSet = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:312
				// _ = "end of CoverTab[111118]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:313
				_go_fuzz_dep_.CoverTab[111119]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:313
				// _ = "end of CoverTab[111119]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:313
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:313
			// _ = "end of CoverTab[111114]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:313
			_go_fuzz_dep_.CoverTab[111115]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:314
			// _ = "end of CoverTab[111115]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:315
			_go_fuzz_dep_.CoverTab[111120]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:315
			// _ = "end of CoverTab[111120]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:315
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:315
		// _ = "end of CoverTab[111099]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:315
		_go_fuzz_dep_.CoverTab[111100]++
													if f.Name == "XXX_extensions" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:316
			_go_fuzz_dep_.CoverTab[111121]++

														if f.Type == reflect.TypeOf((map[int32]Extension)(nil)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:318
				_go_fuzz_dep_.CoverTab[111123]++
															u.oldExtensions = toField(&f)
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:320
				// _ = "end of CoverTab[111123]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:321
				_go_fuzz_dep_.CoverTab[111124]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:321
				if f.Type == reflect.TypeOf(([]byte)(nil)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:321
					_go_fuzz_dep_.CoverTab[111125]++
																u.bytesExtensions = toField(&f)
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:323
					// _ = "end of CoverTab[111125]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:324
					_go_fuzz_dep_.CoverTab[111126]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:324
					// _ = "end of CoverTab[111126]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:324
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:324
				// _ = "end of CoverTab[111124]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:324
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:324
			// _ = "end of CoverTab[111121]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:324
			_go_fuzz_dep_.CoverTab[111122]++
														panic("bad type for XXX_extensions field: " + f.Type.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:325
			// _ = "end of CoverTab[111122]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:326
			_go_fuzz_dep_.CoverTab[111127]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:326
			// _ = "end of CoverTab[111127]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:326
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:326
		// _ = "end of CoverTab[111100]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:326
		_go_fuzz_dep_.CoverTab[111101]++
													if f.Name == "XXX_NoUnkeyedLiteral" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:327
			_go_fuzz_dep_.CoverTab[111128]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:327
			return f.Name == "XXX_sizecache"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:327
			// _ = "end of CoverTab[111128]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:327
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:327
			_go_fuzz_dep_.CoverTab[111129]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:328
			// _ = "end of CoverTab[111129]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:329
			_go_fuzz_dep_.CoverTab[111130]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:329
			// _ = "end of CoverTab[111130]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:329
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:329
		// _ = "end of CoverTab[111101]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:329
		_go_fuzz_dep_.CoverTab[111102]++

													oneof := f.Tag.Get("protobuf_oneof")
													if oneof != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:332
			_go_fuzz_dep_.CoverTab[111131]++
														oneofFields = append(oneofFields, oneofField{f.Type, toField(&f)})

														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:335
			// _ = "end of CoverTab[111131]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:336
			_go_fuzz_dep_.CoverTab[111132]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:336
			// _ = "end of CoverTab[111132]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:336
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:336
		// _ = "end of CoverTab[111102]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:336
		_go_fuzz_dep_.CoverTab[111103]++

													tags := f.Tag.Get("protobuf")
													tagArray := strings.Split(tags, ",")
													if len(tagArray) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:340
			_go_fuzz_dep_.CoverTab[111133]++
														panic("protobuf tag not enough fields in " + t.Name() + "." + f.Name + ": " + tags)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:341
			// _ = "end of CoverTab[111133]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:342
			_go_fuzz_dep_.CoverTab[111134]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:342
			// _ = "end of CoverTab[111134]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:342
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:342
		// _ = "end of CoverTab[111103]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:342
		_go_fuzz_dep_.CoverTab[111104]++
													tag, err := strconv.Atoi(tagArray[1])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:344
			_go_fuzz_dep_.CoverTab[111135]++
														panic("protobuf tag field not an integer: " + tagArray[1])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:345
			// _ = "end of CoverTab[111135]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:346
			_go_fuzz_dep_.CoverTab[111136]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:346
			// _ = "end of CoverTab[111136]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:346
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:346
		// _ = "end of CoverTab[111104]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:346
		_go_fuzz_dep_.CoverTab[111105]++

													name := ""
													for _, tag := range tagArray[3:] {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:349
			_go_fuzz_dep_.CoverTab[111137]++
														if strings.HasPrefix(tag, "name=") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:350
				_go_fuzz_dep_.CoverTab[111138]++
															name = tag[5:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:351
				// _ = "end of CoverTab[111138]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:352
				_go_fuzz_dep_.CoverTab[111139]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:352
				// _ = "end of CoverTab[111139]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:352
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:352
			// _ = "end of CoverTab[111137]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:353
		// _ = "end of CoverTab[111105]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:353
		_go_fuzz_dep_.CoverTab[111106]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:356
		unmarshal := fieldUnmarshaler(&f)

		// Required field?
		var reqMask uint64
		if tagArray[2] == "req" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:360
			_go_fuzz_dep_.CoverTab[111140]++
														bit := len(u.reqFields)
														u.reqFields = append(u.reqFields, name)
														reqMask = uint64(1) << uint(bit)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:363
			// _ = "end of CoverTab[111140]"

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:367
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:367
			_go_fuzz_dep_.CoverTab[111141]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:367
			// _ = "end of CoverTab[111141]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:367
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:367
		// _ = "end of CoverTab[111106]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:367
		_go_fuzz_dep_.CoverTab[111107]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:370
		u.setTag(tag, toField(&f), unmarshal, reqMask, name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:370
		// _ = "end of CoverTab[111107]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:371
	// _ = "end of CoverTab[111091]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:371
	_go_fuzz_dep_.CoverTab[111092]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:375
	if len(oneofFields) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:375
		_go_fuzz_dep_.CoverTab[111142]++
													var oneofImplementers []interface{}
													switch m := reflect.Zero(reflect.PtrTo(t)).Interface().(type) {
		case oneofFuncsIface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:378
			_go_fuzz_dep_.CoverTab[111144]++
														_, _, _, oneofImplementers = m.XXX_OneofFuncs()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:379
			// _ = "end of CoverTab[111144]"
		case oneofWrappersIface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:380
			_go_fuzz_dep_.CoverTab[111145]++
														oneofImplementers = m.XXX_OneofWrappers()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:381
			// _ = "end of CoverTab[111145]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:382
		// _ = "end of CoverTab[111142]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:382
		_go_fuzz_dep_.CoverTab[111143]++
													for _, v := range oneofImplementers {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:383
			_go_fuzz_dep_.CoverTab[111146]++
														tptr := reflect.TypeOf(v)
														typ := tptr.Elem()

														f := typ.Field(0)
														baseUnmarshal := fieldUnmarshaler(&f)
														tags := strings.Split(f.Tag.Get("protobuf"), ",")
														fieldNum, err := strconv.Atoi(tags[1])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:391
				_go_fuzz_dep_.CoverTab[111149]++
															panic("protobuf tag field not an integer: " + tags[1])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:392
				// _ = "end of CoverTab[111149]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:393
				_go_fuzz_dep_.CoverTab[111150]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:393
				// _ = "end of CoverTab[111150]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:393
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:393
			// _ = "end of CoverTab[111146]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:393
			_go_fuzz_dep_.CoverTab[111147]++
														var name string
														for _, tag := range tags {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:395
				_go_fuzz_dep_.CoverTab[111151]++
															if strings.HasPrefix(tag, "name=") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:396
					_go_fuzz_dep_.CoverTab[111152]++
																name = strings.TrimPrefix(tag, "name=")
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:398
					// _ = "end of CoverTab[111152]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:399
					_go_fuzz_dep_.CoverTab[111153]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:399
					// _ = "end of CoverTab[111153]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:399
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:399
				// _ = "end of CoverTab[111151]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:400
			// _ = "end of CoverTab[111147]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:400
			_go_fuzz_dep_.CoverTab[111148]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:404
			for _, of := range oneofFields {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:404
				_go_fuzz_dep_.CoverTab[111154]++
															if tptr.Implements(of.ityp) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:405
					_go_fuzz_dep_.CoverTab[111155]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:409
					unmarshal := makeUnmarshalOneof(typ, of.ityp, baseUnmarshal)
																u.setTag(fieldNum, of.field, unmarshal, 0, name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:410
					// _ = "end of CoverTab[111155]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:411
					_go_fuzz_dep_.CoverTab[111156]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:411
					// _ = "end of CoverTab[111156]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:411
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:411
				// _ = "end of CoverTab[111154]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:412
			// _ = "end of CoverTab[111148]"

		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:414
		// _ = "end of CoverTab[111143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:415
		_go_fuzz_dep_.CoverTab[111157]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:415
		// _ = "end of CoverTab[111157]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:415
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:415
	// _ = "end of CoverTab[111092]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:415
	_go_fuzz_dep_.CoverTab[111093]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:418
	fn := reflect.Zero(reflect.PtrTo(t)).MethodByName("ExtensionRangeArray")
	if fn.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:419
		_go_fuzz_dep_.CoverTab[111158]++
													if !u.extensions.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:420
			_go_fuzz_dep_.CoverTab[111160]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:420
			return !u.oldExtensions.IsValid()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:420
			// _ = "end of CoverTab[111160]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:420
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:420
			_go_fuzz_dep_.CoverTab[111161]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:420
			return !u.bytesExtensions.IsValid()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:420
			// _ = "end of CoverTab[111161]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:420
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:420
			_go_fuzz_dep_.CoverTab[111162]++
														panic("a message with extensions, but no extensions field in " + t.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:421
			// _ = "end of CoverTab[111162]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:422
			_go_fuzz_dep_.CoverTab[111163]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:422
			// _ = "end of CoverTab[111163]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:422
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:422
		// _ = "end of CoverTab[111158]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:422
		_go_fuzz_dep_.CoverTab[111159]++
													u.extensionRanges = fn.Call(nil)[0].Interface().([]ExtensionRange)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:423
		// _ = "end of CoverTab[111159]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:424
		_go_fuzz_dep_.CoverTab[111164]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:424
		// _ = "end of CoverTab[111164]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:424
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:424
	// _ = "end of CoverTab[111093]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:424
	_go_fuzz_dep_.CoverTab[111094]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:430
	u.setTag(0, zeroField, func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:430
		_go_fuzz_dep_.CoverTab[111165]++
													return nil, fmt.Errorf("proto: %s: illegal tag 0 (wire type %d)", t, w)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:431
		// _ = "end of CoverTab[111165]"
	}, 0, "")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:432
	// _ = "end of CoverTab[111094]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:432
	_go_fuzz_dep_.CoverTab[111095]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:435
	u.reqMask = uint64(1)<<uint(len(u.reqFields)) - 1

												atomic.StoreInt32(&u.initialized, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:437
	// _ = "end of CoverTab[111095]"
}

// setTag stores the unmarshal information for the given tag.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:440
// tag = tag # for field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:440
// field/unmarshal = unmarshal info for that field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:440
// reqMask = if required, bitmask for field position in required field list. 0 otherwise.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:440
// name = short name of the field.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:445
func (u *unmarshalInfo) setTag(tag int, field field, unmarshal unmarshaler, reqMask uint64, name string) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:445
	_go_fuzz_dep_.CoverTab[111166]++
												i := unmarshalFieldInfo{field: field, unmarshal: unmarshal, reqMask: reqMask, name: name}
												n := u.typ.NumField()
												if tag >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:448
		_go_fuzz_dep_.CoverTab[111169]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:448
		return (tag < 16 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:448
			_go_fuzz_dep_.CoverTab[111170]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:448
			return tag < 2*n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:448
			// _ = "end of CoverTab[111170]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:448
		}())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:448
		// _ = "end of CoverTab[111169]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:448
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:448
		_go_fuzz_dep_.CoverTab[111171]++
													for len(u.dense) <= tag {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:449
			_go_fuzz_dep_.CoverTab[111173]++
														u.dense = append(u.dense, unmarshalFieldInfo{})
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:450
			// _ = "end of CoverTab[111173]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:451
		// _ = "end of CoverTab[111171]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:451
		_go_fuzz_dep_.CoverTab[111172]++
													u.dense[tag] = i
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:453
		// _ = "end of CoverTab[111172]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:454
		_go_fuzz_dep_.CoverTab[111174]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:454
		// _ = "end of CoverTab[111174]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:454
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:454
	// _ = "end of CoverTab[111166]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:454
	_go_fuzz_dep_.CoverTab[111167]++
												if u.sparse == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:455
		_go_fuzz_dep_.CoverTab[111175]++
													u.sparse = map[uint64]unmarshalFieldInfo{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:456
		// _ = "end of CoverTab[111175]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:457
		_go_fuzz_dep_.CoverTab[111176]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:457
		// _ = "end of CoverTab[111176]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:457
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:457
	// _ = "end of CoverTab[111167]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:457
	_go_fuzz_dep_.CoverTab[111168]++
												u.sparse[uint64(tag)] = i
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:458
	// _ = "end of CoverTab[111168]"
}

// fieldUnmarshaler returns an unmarshaler for the given field.
func fieldUnmarshaler(f *reflect.StructField) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:462
	_go_fuzz_dep_.CoverTab[111177]++
												if f.Type.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:463
		_go_fuzz_dep_.CoverTab[111179]++
													return makeUnmarshalMap(f)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:464
		// _ = "end of CoverTab[111179]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:465
		_go_fuzz_dep_.CoverTab[111180]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:465
		// _ = "end of CoverTab[111180]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:465
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:465
	// _ = "end of CoverTab[111177]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:465
	_go_fuzz_dep_.CoverTab[111178]++
												return typeUnmarshaler(f.Type, f.Tag.Get("protobuf"))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:466
	// _ = "end of CoverTab[111178]"
}

// typeUnmarshaler returns an unmarshaler for the given field type / field tag pair.
func typeUnmarshaler(t reflect.Type, tags string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:470
	_go_fuzz_dep_.CoverTab[111181]++
												tagArray := strings.Split(tags, ",")
												encoding := tagArray[0]
												name := "unknown"
												ctype := false
												isTime := false
												isDuration := false
												isWktPointer := false
												proto3 := false
												validateUTF8 := true
												for _, tag := range tagArray[3:] {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:480
		_go_fuzz_dep_.CoverTab[111191]++
													if strings.HasPrefix(tag, "name=") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:481
			_go_fuzz_dep_.CoverTab[111197]++
														name = tag[5:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:482
			// _ = "end of CoverTab[111197]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:483
			_go_fuzz_dep_.CoverTab[111198]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:483
			// _ = "end of CoverTab[111198]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:483
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:483
		// _ = "end of CoverTab[111191]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:483
		_go_fuzz_dep_.CoverTab[111192]++
													if tag == "proto3" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:484
			_go_fuzz_dep_.CoverTab[111199]++
														proto3 = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:485
			// _ = "end of CoverTab[111199]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:486
			_go_fuzz_dep_.CoverTab[111200]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:486
			// _ = "end of CoverTab[111200]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:486
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:486
		// _ = "end of CoverTab[111192]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:486
		_go_fuzz_dep_.CoverTab[111193]++
													if strings.HasPrefix(tag, "customtype=") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:487
			_go_fuzz_dep_.CoverTab[111201]++
														ctype = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:488
			// _ = "end of CoverTab[111201]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:489
			_go_fuzz_dep_.CoverTab[111202]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:489
			// _ = "end of CoverTab[111202]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:489
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:489
		// _ = "end of CoverTab[111193]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:489
		_go_fuzz_dep_.CoverTab[111194]++
													if tag == "stdtime" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:490
			_go_fuzz_dep_.CoverTab[111203]++
														isTime = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:491
			// _ = "end of CoverTab[111203]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:492
			_go_fuzz_dep_.CoverTab[111204]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:492
			// _ = "end of CoverTab[111204]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:492
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:492
		// _ = "end of CoverTab[111194]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:492
		_go_fuzz_dep_.CoverTab[111195]++
													if tag == "stdduration" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:493
			_go_fuzz_dep_.CoverTab[111205]++
														isDuration = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:494
			// _ = "end of CoverTab[111205]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:495
			_go_fuzz_dep_.CoverTab[111206]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:495
			// _ = "end of CoverTab[111206]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:495
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:495
		// _ = "end of CoverTab[111195]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:495
		_go_fuzz_dep_.CoverTab[111196]++
													if tag == "wktptr" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:496
			_go_fuzz_dep_.CoverTab[111207]++
														isWktPointer = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:497
			// _ = "end of CoverTab[111207]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:498
			_go_fuzz_dep_.CoverTab[111208]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:498
			// _ = "end of CoverTab[111208]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:498
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:498
		// _ = "end of CoverTab[111196]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:499
	// _ = "end of CoverTab[111181]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:499
	_go_fuzz_dep_.CoverTab[111182]++
												validateUTF8 = validateUTF8 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:500
		_go_fuzz_dep_.CoverTab[111209]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:500
		return proto3
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:500
		// _ = "end of CoverTab[111209]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:500
	}()

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:503
	slice := false
	pointer := false
	if t.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:505
		_go_fuzz_dep_.CoverTab[111210]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:505
		return t.Elem().Kind() != reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:505
		// _ = "end of CoverTab[111210]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:505
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:505
		_go_fuzz_dep_.CoverTab[111211]++
													slice = true
													t = t.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:507
		// _ = "end of CoverTab[111211]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:508
		_go_fuzz_dep_.CoverTab[111212]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:508
		// _ = "end of CoverTab[111212]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:508
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:508
	// _ = "end of CoverTab[111182]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:508
	_go_fuzz_dep_.CoverTab[111183]++
												if t.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:509
		_go_fuzz_dep_.CoverTab[111213]++
													pointer = true
													t = t.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:511
		// _ = "end of CoverTab[111213]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:512
		_go_fuzz_dep_.CoverTab[111214]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:512
		// _ = "end of CoverTab[111214]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:512
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:512
	// _ = "end of CoverTab[111183]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:512
	_go_fuzz_dep_.CoverTab[111184]++

												if ctype {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:514
		_go_fuzz_dep_.CoverTab[111215]++
													if reflect.PtrTo(t).Implements(customType) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:515
			_go_fuzz_dep_.CoverTab[111216]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:516
				_go_fuzz_dep_.CoverTab[111219]++
															return makeUnmarshalCustomSlice(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:517
				// _ = "end of CoverTab[111219]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:518
				_go_fuzz_dep_.CoverTab[111220]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:518
				// _ = "end of CoverTab[111220]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:518
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:518
			// _ = "end of CoverTab[111216]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:518
			_go_fuzz_dep_.CoverTab[111217]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:519
				_go_fuzz_dep_.CoverTab[111221]++
															return makeUnmarshalCustomPtr(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:520
				// _ = "end of CoverTab[111221]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:521
				_go_fuzz_dep_.CoverTab[111222]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:521
				// _ = "end of CoverTab[111222]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:521
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:521
			// _ = "end of CoverTab[111217]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:521
			_go_fuzz_dep_.CoverTab[111218]++
														return makeUnmarshalCustom(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:522
			// _ = "end of CoverTab[111218]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:523
			_go_fuzz_dep_.CoverTab[111223]++
														panic(fmt.Sprintf("custom type: type: %v, does not implement the proto.custom interface", t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:524
			// _ = "end of CoverTab[111223]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:525
		// _ = "end of CoverTab[111215]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:526
		_go_fuzz_dep_.CoverTab[111224]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:526
		// _ = "end of CoverTab[111224]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:526
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:526
	// _ = "end of CoverTab[111184]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:526
	_go_fuzz_dep_.CoverTab[111185]++

												if isTime {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:528
		_go_fuzz_dep_.CoverTab[111225]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:529
			_go_fuzz_dep_.CoverTab[111228]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:530
				_go_fuzz_dep_.CoverTab[111230]++
															return makeUnmarshalTimePtrSlice(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:531
				// _ = "end of CoverTab[111230]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:532
				_go_fuzz_dep_.CoverTab[111231]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:532
				// _ = "end of CoverTab[111231]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:532
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:532
			// _ = "end of CoverTab[111228]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:532
			_go_fuzz_dep_.CoverTab[111229]++
														return makeUnmarshalTimePtr(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:533
			// _ = "end of CoverTab[111229]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:534
			_go_fuzz_dep_.CoverTab[111232]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:534
			// _ = "end of CoverTab[111232]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:534
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:534
		// _ = "end of CoverTab[111225]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:534
		_go_fuzz_dep_.CoverTab[111226]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:535
			_go_fuzz_dep_.CoverTab[111233]++
														return makeUnmarshalTimeSlice(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:536
			// _ = "end of CoverTab[111233]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:537
			_go_fuzz_dep_.CoverTab[111234]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:537
			// _ = "end of CoverTab[111234]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:537
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:537
		// _ = "end of CoverTab[111226]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:537
		_go_fuzz_dep_.CoverTab[111227]++
													return makeUnmarshalTime(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:538
		// _ = "end of CoverTab[111227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:539
		_go_fuzz_dep_.CoverTab[111235]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:539
		// _ = "end of CoverTab[111235]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:539
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:539
	// _ = "end of CoverTab[111185]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:539
	_go_fuzz_dep_.CoverTab[111186]++

												if isDuration {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:541
		_go_fuzz_dep_.CoverTab[111236]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:542
			_go_fuzz_dep_.CoverTab[111239]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:543
				_go_fuzz_dep_.CoverTab[111241]++
															return makeUnmarshalDurationPtrSlice(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:544
				// _ = "end of CoverTab[111241]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:545
				_go_fuzz_dep_.CoverTab[111242]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:545
				// _ = "end of CoverTab[111242]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:545
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:545
			// _ = "end of CoverTab[111239]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:545
			_go_fuzz_dep_.CoverTab[111240]++
														return makeUnmarshalDurationPtr(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:546
			// _ = "end of CoverTab[111240]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:547
			_go_fuzz_dep_.CoverTab[111243]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:547
			// _ = "end of CoverTab[111243]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:547
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:547
		// _ = "end of CoverTab[111236]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:547
		_go_fuzz_dep_.CoverTab[111237]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:548
			_go_fuzz_dep_.CoverTab[111244]++
														return makeUnmarshalDurationSlice(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:549
			// _ = "end of CoverTab[111244]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:550
			_go_fuzz_dep_.CoverTab[111245]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:550
			// _ = "end of CoverTab[111245]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:550
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:550
		// _ = "end of CoverTab[111237]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:550
		_go_fuzz_dep_.CoverTab[111238]++
													return makeUnmarshalDuration(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:551
		// _ = "end of CoverTab[111238]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:552
		_go_fuzz_dep_.CoverTab[111246]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:552
		// _ = "end of CoverTab[111246]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:552
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:552
	// _ = "end of CoverTab[111186]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:552
	_go_fuzz_dep_.CoverTab[111187]++

												if isWktPointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:554
		_go_fuzz_dep_.CoverTab[111247]++
													switch t.Kind() {
		case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:556
			_go_fuzz_dep_.CoverTab[111248]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:557
				_go_fuzz_dep_.CoverTab[111276]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:558
					_go_fuzz_dep_.CoverTab[111278]++
																return makeStdDoubleValuePtrSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:559
					// _ = "end of CoverTab[111278]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:560
					_go_fuzz_dep_.CoverTab[111279]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:560
					// _ = "end of CoverTab[111279]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:560
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:560
				// _ = "end of CoverTab[111276]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:560
				_go_fuzz_dep_.CoverTab[111277]++
															return makeStdDoubleValuePtrUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:561
				// _ = "end of CoverTab[111277]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:562
				_go_fuzz_dep_.CoverTab[111280]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:562
				// _ = "end of CoverTab[111280]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:562
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:562
			// _ = "end of CoverTab[111248]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:562
			_go_fuzz_dep_.CoverTab[111249]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:563
				_go_fuzz_dep_.CoverTab[111281]++
															return makeStdDoubleValueSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:564
				// _ = "end of CoverTab[111281]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:565
				_go_fuzz_dep_.CoverTab[111282]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:565
				// _ = "end of CoverTab[111282]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:565
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:565
			// _ = "end of CoverTab[111249]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:565
			_go_fuzz_dep_.CoverTab[111250]++
														return makeStdDoubleValueUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:566
			// _ = "end of CoverTab[111250]"
		case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:567
			_go_fuzz_dep_.CoverTab[111251]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:568
				_go_fuzz_dep_.CoverTab[111283]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:569
					_go_fuzz_dep_.CoverTab[111285]++
																return makeStdFloatValuePtrSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:570
					// _ = "end of CoverTab[111285]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:571
					_go_fuzz_dep_.CoverTab[111286]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:571
					// _ = "end of CoverTab[111286]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:571
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:571
				// _ = "end of CoverTab[111283]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:571
				_go_fuzz_dep_.CoverTab[111284]++
															return makeStdFloatValuePtrUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:572
				// _ = "end of CoverTab[111284]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:573
				_go_fuzz_dep_.CoverTab[111287]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:573
				// _ = "end of CoverTab[111287]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:573
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:573
			// _ = "end of CoverTab[111251]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:573
			_go_fuzz_dep_.CoverTab[111252]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:574
				_go_fuzz_dep_.CoverTab[111288]++
															return makeStdFloatValueSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:575
				// _ = "end of CoverTab[111288]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:576
				_go_fuzz_dep_.CoverTab[111289]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:576
				// _ = "end of CoverTab[111289]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:576
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:576
			// _ = "end of CoverTab[111252]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:576
			_go_fuzz_dep_.CoverTab[111253]++
														return makeStdFloatValueUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:577
			// _ = "end of CoverTab[111253]"
		case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:578
			_go_fuzz_dep_.CoverTab[111254]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:579
				_go_fuzz_dep_.CoverTab[111290]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:580
					_go_fuzz_dep_.CoverTab[111292]++
																return makeStdInt64ValuePtrSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:581
					// _ = "end of CoverTab[111292]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:582
					_go_fuzz_dep_.CoverTab[111293]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:582
					// _ = "end of CoverTab[111293]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:582
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:582
				// _ = "end of CoverTab[111290]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:582
				_go_fuzz_dep_.CoverTab[111291]++
															return makeStdInt64ValuePtrUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:583
				// _ = "end of CoverTab[111291]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:584
				_go_fuzz_dep_.CoverTab[111294]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:584
				// _ = "end of CoverTab[111294]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:584
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:584
			// _ = "end of CoverTab[111254]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:584
			_go_fuzz_dep_.CoverTab[111255]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:585
				_go_fuzz_dep_.CoverTab[111295]++
															return makeStdInt64ValueSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:586
				// _ = "end of CoverTab[111295]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:587
				_go_fuzz_dep_.CoverTab[111296]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:587
				// _ = "end of CoverTab[111296]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:587
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:587
			// _ = "end of CoverTab[111255]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:587
			_go_fuzz_dep_.CoverTab[111256]++
														return makeStdInt64ValueUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:588
			// _ = "end of CoverTab[111256]"
		case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:589
			_go_fuzz_dep_.CoverTab[111257]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:590
				_go_fuzz_dep_.CoverTab[111297]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:591
					_go_fuzz_dep_.CoverTab[111299]++
																return makeStdUInt64ValuePtrSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:592
					// _ = "end of CoverTab[111299]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:593
					_go_fuzz_dep_.CoverTab[111300]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:593
					// _ = "end of CoverTab[111300]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:593
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:593
				// _ = "end of CoverTab[111297]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:593
				_go_fuzz_dep_.CoverTab[111298]++
															return makeStdUInt64ValuePtrUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:594
				// _ = "end of CoverTab[111298]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:595
				_go_fuzz_dep_.CoverTab[111301]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:595
				// _ = "end of CoverTab[111301]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:595
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:595
			// _ = "end of CoverTab[111257]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:595
			_go_fuzz_dep_.CoverTab[111258]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:596
				_go_fuzz_dep_.CoverTab[111302]++
															return makeStdUInt64ValueSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:597
				// _ = "end of CoverTab[111302]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:598
				_go_fuzz_dep_.CoverTab[111303]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:598
				// _ = "end of CoverTab[111303]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:598
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:598
			// _ = "end of CoverTab[111258]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:598
			_go_fuzz_dep_.CoverTab[111259]++
														return makeStdUInt64ValueUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:599
			// _ = "end of CoverTab[111259]"
		case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:600
			_go_fuzz_dep_.CoverTab[111260]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:601
				_go_fuzz_dep_.CoverTab[111304]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:602
					_go_fuzz_dep_.CoverTab[111306]++
																return makeStdInt32ValuePtrSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:603
					// _ = "end of CoverTab[111306]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:604
					_go_fuzz_dep_.CoverTab[111307]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:604
					// _ = "end of CoverTab[111307]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:604
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:604
				// _ = "end of CoverTab[111304]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:604
				_go_fuzz_dep_.CoverTab[111305]++
															return makeStdInt32ValuePtrUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:605
				// _ = "end of CoverTab[111305]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:606
				_go_fuzz_dep_.CoverTab[111308]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:606
				// _ = "end of CoverTab[111308]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:606
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:606
			// _ = "end of CoverTab[111260]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:606
			_go_fuzz_dep_.CoverTab[111261]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:607
				_go_fuzz_dep_.CoverTab[111309]++
															return makeStdInt32ValueSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:608
				// _ = "end of CoverTab[111309]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:609
				_go_fuzz_dep_.CoverTab[111310]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:609
				// _ = "end of CoverTab[111310]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:609
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:609
			// _ = "end of CoverTab[111261]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:609
			_go_fuzz_dep_.CoverTab[111262]++
														return makeStdInt32ValueUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:610
			// _ = "end of CoverTab[111262]"
		case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:611
			_go_fuzz_dep_.CoverTab[111263]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:612
				_go_fuzz_dep_.CoverTab[111311]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:613
					_go_fuzz_dep_.CoverTab[111313]++
																return makeStdUInt32ValuePtrSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:614
					// _ = "end of CoverTab[111313]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:615
					_go_fuzz_dep_.CoverTab[111314]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:615
					// _ = "end of CoverTab[111314]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:615
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:615
				// _ = "end of CoverTab[111311]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:615
				_go_fuzz_dep_.CoverTab[111312]++
															return makeStdUInt32ValuePtrUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:616
				// _ = "end of CoverTab[111312]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:617
				_go_fuzz_dep_.CoverTab[111315]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:617
				// _ = "end of CoverTab[111315]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:617
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:617
			// _ = "end of CoverTab[111263]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:617
			_go_fuzz_dep_.CoverTab[111264]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:618
				_go_fuzz_dep_.CoverTab[111316]++
															return makeStdUInt32ValueSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:619
				// _ = "end of CoverTab[111316]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:620
				_go_fuzz_dep_.CoverTab[111317]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:620
				// _ = "end of CoverTab[111317]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:620
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:620
			// _ = "end of CoverTab[111264]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:620
			_go_fuzz_dep_.CoverTab[111265]++
														return makeStdUInt32ValueUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:621
			// _ = "end of CoverTab[111265]"
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:622
			_go_fuzz_dep_.CoverTab[111266]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:623
				_go_fuzz_dep_.CoverTab[111318]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:624
					_go_fuzz_dep_.CoverTab[111320]++
																return makeStdBoolValuePtrSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:625
					// _ = "end of CoverTab[111320]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:626
					_go_fuzz_dep_.CoverTab[111321]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:626
					// _ = "end of CoverTab[111321]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:626
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:626
				// _ = "end of CoverTab[111318]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:626
				_go_fuzz_dep_.CoverTab[111319]++
															return makeStdBoolValuePtrUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:627
				// _ = "end of CoverTab[111319]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:628
				_go_fuzz_dep_.CoverTab[111322]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:628
				// _ = "end of CoverTab[111322]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:628
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:628
			// _ = "end of CoverTab[111266]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:628
			_go_fuzz_dep_.CoverTab[111267]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:629
				_go_fuzz_dep_.CoverTab[111323]++
															return makeStdBoolValueSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:630
				// _ = "end of CoverTab[111323]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:631
				_go_fuzz_dep_.CoverTab[111324]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:631
				// _ = "end of CoverTab[111324]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:631
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:631
			// _ = "end of CoverTab[111267]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:631
			_go_fuzz_dep_.CoverTab[111268]++
														return makeStdBoolValueUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:632
			// _ = "end of CoverTab[111268]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:633
			_go_fuzz_dep_.CoverTab[111269]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:634
				_go_fuzz_dep_.CoverTab[111325]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:635
					_go_fuzz_dep_.CoverTab[111327]++
																return makeStdStringValuePtrSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:636
					// _ = "end of CoverTab[111327]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:637
					_go_fuzz_dep_.CoverTab[111328]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:637
					// _ = "end of CoverTab[111328]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:637
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:637
				// _ = "end of CoverTab[111325]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:637
				_go_fuzz_dep_.CoverTab[111326]++
															return makeStdStringValuePtrUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:638
				// _ = "end of CoverTab[111326]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:639
				_go_fuzz_dep_.CoverTab[111329]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:639
				// _ = "end of CoverTab[111329]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:639
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:639
			// _ = "end of CoverTab[111269]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:639
			_go_fuzz_dep_.CoverTab[111270]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:640
				_go_fuzz_dep_.CoverTab[111330]++
															return makeStdStringValueSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:641
				// _ = "end of CoverTab[111330]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:642
				_go_fuzz_dep_.CoverTab[111331]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:642
				// _ = "end of CoverTab[111331]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:642
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:642
			// _ = "end of CoverTab[111270]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:642
			_go_fuzz_dep_.CoverTab[111271]++
														return makeStdStringValueUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:643
			// _ = "end of CoverTab[111271]"
		case uint8SliceType:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:644
			_go_fuzz_dep_.CoverTab[111272]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:645
				_go_fuzz_dep_.CoverTab[111332]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:646
					_go_fuzz_dep_.CoverTab[111334]++
																return makeStdBytesValuePtrSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:647
					// _ = "end of CoverTab[111334]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:648
					_go_fuzz_dep_.CoverTab[111335]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:648
					// _ = "end of CoverTab[111335]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:648
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:648
				// _ = "end of CoverTab[111332]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:648
				_go_fuzz_dep_.CoverTab[111333]++
															return makeStdBytesValuePtrUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:649
				// _ = "end of CoverTab[111333]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:650
				_go_fuzz_dep_.CoverTab[111336]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:650
				// _ = "end of CoverTab[111336]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:650
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:650
			// _ = "end of CoverTab[111272]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:650
			_go_fuzz_dep_.CoverTab[111273]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:651
				_go_fuzz_dep_.CoverTab[111337]++
															return makeStdBytesValueSliceUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:652
				// _ = "end of CoverTab[111337]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:653
				_go_fuzz_dep_.CoverTab[111338]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:653
				// _ = "end of CoverTab[111338]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:653
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:653
			// _ = "end of CoverTab[111273]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:653
			_go_fuzz_dep_.CoverTab[111274]++
														return makeStdBytesValueUnmarshaler(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:654
			// _ = "end of CoverTab[111274]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:655
			_go_fuzz_dep_.CoverTab[111275]++
														panic(fmt.Sprintf("unknown wktpointer type %#v", t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:656
			// _ = "end of CoverTab[111275]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:657
		// _ = "end of CoverTab[111247]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:658
		_go_fuzz_dep_.CoverTab[111339]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:658
		// _ = "end of CoverTab[111339]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:658
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:658
	// _ = "end of CoverTab[111187]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:658
	_go_fuzz_dep_.CoverTab[111188]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
	if pointer && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
		_go_fuzz_dep_.CoverTab[111340]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
		return slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
		// _ = "end of CoverTab[111340]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
		_go_fuzz_dep_.CoverTab[111341]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
		return t.Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
		// _ = "end of CoverTab[111341]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:661
		_go_fuzz_dep_.CoverTab[111342]++
													panic("both pointer and slice for basic type in " + t.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:662
		// _ = "end of CoverTab[111342]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:663
		_go_fuzz_dep_.CoverTab[111343]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:663
		// _ = "end of CoverTab[111343]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:663
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:663
	// _ = "end of CoverTab[111188]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:663
	_go_fuzz_dep_.CoverTab[111189]++

												switch t.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:666
		_go_fuzz_dep_.CoverTab[111344]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:667
			_go_fuzz_dep_.CoverTab[111368]++
														return unmarshalBoolPtr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:668
			// _ = "end of CoverTab[111368]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:669
			_go_fuzz_dep_.CoverTab[111369]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:669
			// _ = "end of CoverTab[111369]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:669
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:669
		// _ = "end of CoverTab[111344]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:669
		_go_fuzz_dep_.CoverTab[111345]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:670
			_go_fuzz_dep_.CoverTab[111370]++
														return unmarshalBoolSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:671
			// _ = "end of CoverTab[111370]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:672
			_go_fuzz_dep_.CoverTab[111371]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:672
			// _ = "end of CoverTab[111371]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:672
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:672
		// _ = "end of CoverTab[111345]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:672
		_go_fuzz_dep_.CoverTab[111346]++
													return unmarshalBoolValue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:673
		// _ = "end of CoverTab[111346]"
	case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:674
		_go_fuzz_dep_.CoverTab[111347]++
													switch encoding {
		case "fixed32":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:676
			_go_fuzz_dep_.CoverTab[111372]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:677
				_go_fuzz_dep_.CoverTab[111382]++
															return unmarshalFixedS32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:678
				// _ = "end of CoverTab[111382]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:679
				_go_fuzz_dep_.CoverTab[111383]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:679
				// _ = "end of CoverTab[111383]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:679
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:679
			// _ = "end of CoverTab[111372]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:679
			_go_fuzz_dep_.CoverTab[111373]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:680
				_go_fuzz_dep_.CoverTab[111384]++
															return unmarshalFixedS32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:681
				// _ = "end of CoverTab[111384]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:682
				_go_fuzz_dep_.CoverTab[111385]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:682
				// _ = "end of CoverTab[111385]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:682
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:682
			// _ = "end of CoverTab[111373]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:682
			_go_fuzz_dep_.CoverTab[111374]++
														return unmarshalFixedS32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:683
			// _ = "end of CoverTab[111374]"
		case "varint":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:684
			_go_fuzz_dep_.CoverTab[111375]++

														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:686
				_go_fuzz_dep_.CoverTab[111386]++
															return unmarshalInt32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:687
				// _ = "end of CoverTab[111386]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:688
				_go_fuzz_dep_.CoverTab[111387]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:688
				// _ = "end of CoverTab[111387]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:688
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:688
			// _ = "end of CoverTab[111375]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:688
			_go_fuzz_dep_.CoverTab[111376]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:689
				_go_fuzz_dep_.CoverTab[111388]++
															return unmarshalInt32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:690
				// _ = "end of CoverTab[111388]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:691
				_go_fuzz_dep_.CoverTab[111389]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:691
				// _ = "end of CoverTab[111389]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:691
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:691
			// _ = "end of CoverTab[111376]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:691
			_go_fuzz_dep_.CoverTab[111377]++
														return unmarshalInt32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:692
			// _ = "end of CoverTab[111377]"
		case "zigzag32":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:693
			_go_fuzz_dep_.CoverTab[111378]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:694
				_go_fuzz_dep_.CoverTab[111390]++
															return unmarshalSint32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:695
				// _ = "end of CoverTab[111390]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:696
				_go_fuzz_dep_.CoverTab[111391]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:696
				// _ = "end of CoverTab[111391]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:696
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:696
			// _ = "end of CoverTab[111378]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:696
			_go_fuzz_dep_.CoverTab[111379]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:697
				_go_fuzz_dep_.CoverTab[111392]++
															return unmarshalSint32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:698
				// _ = "end of CoverTab[111392]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:699
				_go_fuzz_dep_.CoverTab[111393]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:699
				// _ = "end of CoverTab[111393]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:699
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:699
			// _ = "end of CoverTab[111379]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:699
			_go_fuzz_dep_.CoverTab[111380]++
														return unmarshalSint32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:700
			// _ = "end of CoverTab[111380]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:700
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:700
			_go_fuzz_dep_.CoverTab[111381]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:700
			// _ = "end of CoverTab[111381]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:701
		// _ = "end of CoverTab[111347]"
	case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:702
		_go_fuzz_dep_.CoverTab[111348]++
													switch encoding {
		case "fixed64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:704
			_go_fuzz_dep_.CoverTab[111394]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:705
				_go_fuzz_dep_.CoverTab[111404]++
															return unmarshalFixedS64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:706
				// _ = "end of CoverTab[111404]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:707
				_go_fuzz_dep_.CoverTab[111405]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:707
				// _ = "end of CoverTab[111405]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:707
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:707
			// _ = "end of CoverTab[111394]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:707
			_go_fuzz_dep_.CoverTab[111395]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:708
				_go_fuzz_dep_.CoverTab[111406]++
															return unmarshalFixedS64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:709
				// _ = "end of CoverTab[111406]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:710
				_go_fuzz_dep_.CoverTab[111407]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:710
				// _ = "end of CoverTab[111407]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:710
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:710
			// _ = "end of CoverTab[111395]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:710
			_go_fuzz_dep_.CoverTab[111396]++
														return unmarshalFixedS64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:711
			// _ = "end of CoverTab[111396]"
		case "varint":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:712
			_go_fuzz_dep_.CoverTab[111397]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:713
				_go_fuzz_dep_.CoverTab[111408]++
															return unmarshalInt64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:714
				// _ = "end of CoverTab[111408]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:715
				_go_fuzz_dep_.CoverTab[111409]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:715
				// _ = "end of CoverTab[111409]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:715
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:715
			// _ = "end of CoverTab[111397]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:715
			_go_fuzz_dep_.CoverTab[111398]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:716
				_go_fuzz_dep_.CoverTab[111410]++
															return unmarshalInt64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:717
				// _ = "end of CoverTab[111410]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:718
				_go_fuzz_dep_.CoverTab[111411]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:718
				// _ = "end of CoverTab[111411]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:718
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:718
			// _ = "end of CoverTab[111398]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:718
			_go_fuzz_dep_.CoverTab[111399]++
														return unmarshalInt64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:719
			// _ = "end of CoverTab[111399]"
		case "zigzag64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:720
			_go_fuzz_dep_.CoverTab[111400]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:721
				_go_fuzz_dep_.CoverTab[111412]++
															return unmarshalSint64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:722
				// _ = "end of CoverTab[111412]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:723
				_go_fuzz_dep_.CoverTab[111413]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:723
				// _ = "end of CoverTab[111413]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:723
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:723
			// _ = "end of CoverTab[111400]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:723
			_go_fuzz_dep_.CoverTab[111401]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:724
				_go_fuzz_dep_.CoverTab[111414]++
															return unmarshalSint64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:725
				// _ = "end of CoverTab[111414]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:726
				_go_fuzz_dep_.CoverTab[111415]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:726
				// _ = "end of CoverTab[111415]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:726
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:726
			// _ = "end of CoverTab[111401]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:726
			_go_fuzz_dep_.CoverTab[111402]++
														return unmarshalSint64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:727
			// _ = "end of CoverTab[111402]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:727
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:727
			_go_fuzz_dep_.CoverTab[111403]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:727
			// _ = "end of CoverTab[111403]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:728
		// _ = "end of CoverTab[111348]"
	case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:729
		_go_fuzz_dep_.CoverTab[111349]++
													switch encoding {
		case "fixed32":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:731
			_go_fuzz_dep_.CoverTab[111416]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:732
				_go_fuzz_dep_.CoverTab[111423]++
															return unmarshalFixed32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:733
				// _ = "end of CoverTab[111423]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:734
				_go_fuzz_dep_.CoverTab[111424]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:734
				// _ = "end of CoverTab[111424]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:734
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:734
			// _ = "end of CoverTab[111416]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:734
			_go_fuzz_dep_.CoverTab[111417]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:735
				_go_fuzz_dep_.CoverTab[111425]++
															return unmarshalFixed32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:736
				// _ = "end of CoverTab[111425]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:737
				_go_fuzz_dep_.CoverTab[111426]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:737
				// _ = "end of CoverTab[111426]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:737
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:737
			// _ = "end of CoverTab[111417]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:737
			_go_fuzz_dep_.CoverTab[111418]++
														return unmarshalFixed32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:738
			// _ = "end of CoverTab[111418]"
		case "varint":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:739
			_go_fuzz_dep_.CoverTab[111419]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:740
				_go_fuzz_dep_.CoverTab[111427]++
															return unmarshalUint32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:741
				// _ = "end of CoverTab[111427]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:742
				_go_fuzz_dep_.CoverTab[111428]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:742
				// _ = "end of CoverTab[111428]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:742
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:742
			// _ = "end of CoverTab[111419]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:742
			_go_fuzz_dep_.CoverTab[111420]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:743
				_go_fuzz_dep_.CoverTab[111429]++
															return unmarshalUint32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:744
				// _ = "end of CoverTab[111429]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:745
				_go_fuzz_dep_.CoverTab[111430]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:745
				// _ = "end of CoverTab[111430]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:745
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:745
			// _ = "end of CoverTab[111420]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:745
			_go_fuzz_dep_.CoverTab[111421]++
														return unmarshalUint32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:746
			// _ = "end of CoverTab[111421]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:746
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:746
			_go_fuzz_dep_.CoverTab[111422]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:746
			// _ = "end of CoverTab[111422]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:747
		// _ = "end of CoverTab[111349]"
	case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:748
		_go_fuzz_dep_.CoverTab[111350]++
													switch encoding {
		case "fixed64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:750
			_go_fuzz_dep_.CoverTab[111431]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:751
				_go_fuzz_dep_.CoverTab[111438]++
															return unmarshalFixed64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:752
				// _ = "end of CoverTab[111438]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:753
				_go_fuzz_dep_.CoverTab[111439]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:753
				// _ = "end of CoverTab[111439]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:753
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:753
			// _ = "end of CoverTab[111431]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:753
			_go_fuzz_dep_.CoverTab[111432]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:754
				_go_fuzz_dep_.CoverTab[111440]++
															return unmarshalFixed64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:755
				// _ = "end of CoverTab[111440]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:756
				_go_fuzz_dep_.CoverTab[111441]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:756
				// _ = "end of CoverTab[111441]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:756
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:756
			// _ = "end of CoverTab[111432]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:756
			_go_fuzz_dep_.CoverTab[111433]++
														return unmarshalFixed64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:757
			// _ = "end of CoverTab[111433]"
		case "varint":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:758
			_go_fuzz_dep_.CoverTab[111434]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:759
				_go_fuzz_dep_.CoverTab[111442]++
															return unmarshalUint64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:760
				// _ = "end of CoverTab[111442]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:761
				_go_fuzz_dep_.CoverTab[111443]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:761
				// _ = "end of CoverTab[111443]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:761
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:761
			// _ = "end of CoverTab[111434]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:761
			_go_fuzz_dep_.CoverTab[111435]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:762
				_go_fuzz_dep_.CoverTab[111444]++
															return unmarshalUint64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:763
				// _ = "end of CoverTab[111444]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:764
				_go_fuzz_dep_.CoverTab[111445]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:764
				// _ = "end of CoverTab[111445]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:764
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:764
			// _ = "end of CoverTab[111435]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:764
			_go_fuzz_dep_.CoverTab[111436]++
														return unmarshalUint64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:765
			// _ = "end of CoverTab[111436]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:765
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:765
			_go_fuzz_dep_.CoverTab[111437]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:765
			// _ = "end of CoverTab[111437]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:766
		// _ = "end of CoverTab[111350]"
	case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:767
		_go_fuzz_dep_.CoverTab[111351]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:768
			_go_fuzz_dep_.CoverTab[111446]++
														return unmarshalFloat32Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:769
			// _ = "end of CoverTab[111446]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:770
			_go_fuzz_dep_.CoverTab[111447]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:770
			// _ = "end of CoverTab[111447]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:770
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:770
		// _ = "end of CoverTab[111351]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:770
		_go_fuzz_dep_.CoverTab[111352]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:771
			_go_fuzz_dep_.CoverTab[111448]++
														return unmarshalFloat32Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:772
			// _ = "end of CoverTab[111448]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:773
			_go_fuzz_dep_.CoverTab[111449]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:773
			// _ = "end of CoverTab[111449]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:773
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:773
		// _ = "end of CoverTab[111352]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:773
		_go_fuzz_dep_.CoverTab[111353]++
													return unmarshalFloat32Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:774
		// _ = "end of CoverTab[111353]"
	case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:775
		_go_fuzz_dep_.CoverTab[111354]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:776
			_go_fuzz_dep_.CoverTab[111450]++
														return unmarshalFloat64Ptr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:777
			// _ = "end of CoverTab[111450]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:778
			_go_fuzz_dep_.CoverTab[111451]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:778
			// _ = "end of CoverTab[111451]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:778
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:778
		// _ = "end of CoverTab[111354]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:778
		_go_fuzz_dep_.CoverTab[111355]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:779
			_go_fuzz_dep_.CoverTab[111452]++
														return unmarshalFloat64Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:780
			// _ = "end of CoverTab[111452]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:781
			_go_fuzz_dep_.CoverTab[111453]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:781
			// _ = "end of CoverTab[111453]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:781
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:781
		// _ = "end of CoverTab[111355]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:781
		_go_fuzz_dep_.CoverTab[111356]++
													return unmarshalFloat64Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:782
		// _ = "end of CoverTab[111356]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:783
		_go_fuzz_dep_.CoverTab[111357]++
													panic("map type in typeUnmarshaler in " + t.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:784
		// _ = "end of CoverTab[111357]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:785
		_go_fuzz_dep_.CoverTab[111358]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:786
			_go_fuzz_dep_.CoverTab[111454]++
														panic("bad pointer in slice case in " + t.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:787
			// _ = "end of CoverTab[111454]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:788
			_go_fuzz_dep_.CoverTab[111455]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:788
			// _ = "end of CoverTab[111455]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:788
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:788
		// _ = "end of CoverTab[111358]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:788
		_go_fuzz_dep_.CoverTab[111359]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:789
			_go_fuzz_dep_.CoverTab[111456]++
														return unmarshalBytesSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:790
			// _ = "end of CoverTab[111456]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:791
			_go_fuzz_dep_.CoverTab[111457]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:791
			// _ = "end of CoverTab[111457]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:791
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:791
		// _ = "end of CoverTab[111359]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:791
		_go_fuzz_dep_.CoverTab[111360]++
													return unmarshalBytesValue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:792
		// _ = "end of CoverTab[111360]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:793
		_go_fuzz_dep_.CoverTab[111361]++
													if validateUTF8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:794
			_go_fuzz_dep_.CoverTab[111458]++
														if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:795
				_go_fuzz_dep_.CoverTab[111461]++
															return unmarshalUTF8StringPtr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:796
				// _ = "end of CoverTab[111461]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:797
				_go_fuzz_dep_.CoverTab[111462]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:797
				// _ = "end of CoverTab[111462]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:797
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:797
			// _ = "end of CoverTab[111458]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:797
			_go_fuzz_dep_.CoverTab[111459]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:798
				_go_fuzz_dep_.CoverTab[111463]++
															return unmarshalUTF8StringSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:799
				// _ = "end of CoverTab[111463]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:800
				_go_fuzz_dep_.CoverTab[111464]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:800
				// _ = "end of CoverTab[111464]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:800
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:800
			// _ = "end of CoverTab[111459]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:800
			_go_fuzz_dep_.CoverTab[111460]++
														return unmarshalUTF8StringValue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:801
			// _ = "end of CoverTab[111460]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:802
			_go_fuzz_dep_.CoverTab[111465]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:802
			// _ = "end of CoverTab[111465]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:802
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:802
		// _ = "end of CoverTab[111361]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:802
		_go_fuzz_dep_.CoverTab[111362]++
													if pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:803
			_go_fuzz_dep_.CoverTab[111466]++
														return unmarshalStringPtr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:804
			// _ = "end of CoverTab[111466]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:805
			_go_fuzz_dep_.CoverTab[111467]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:805
			// _ = "end of CoverTab[111467]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:805
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:805
		// _ = "end of CoverTab[111362]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:805
		_go_fuzz_dep_.CoverTab[111363]++
													if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:806
			_go_fuzz_dep_.CoverTab[111468]++
														return unmarshalStringSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:807
			// _ = "end of CoverTab[111468]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:808
			_go_fuzz_dep_.CoverTab[111469]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:808
			// _ = "end of CoverTab[111469]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:808
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:808
		// _ = "end of CoverTab[111363]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:808
		_go_fuzz_dep_.CoverTab[111364]++
													return unmarshalStringValue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:809
		// _ = "end of CoverTab[111364]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:810
		_go_fuzz_dep_.CoverTab[111365]++

													if !pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:812
			_go_fuzz_dep_.CoverTab[111470]++
														switch encoding {
			case "bytes":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:814
				_go_fuzz_dep_.CoverTab[111471]++
															if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:815
					_go_fuzz_dep_.CoverTab[111474]++
																return makeUnmarshalMessageSlice(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:816
					// _ = "end of CoverTab[111474]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:817
					_go_fuzz_dep_.CoverTab[111475]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:817
					// _ = "end of CoverTab[111475]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:817
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:817
				// _ = "end of CoverTab[111471]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:817
				_go_fuzz_dep_.CoverTab[111472]++
															return makeUnmarshalMessage(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:818
				// _ = "end of CoverTab[111472]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:818
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:818
				_go_fuzz_dep_.CoverTab[111473]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:818
				// _ = "end of CoverTab[111473]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:819
			// _ = "end of CoverTab[111470]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:820
			_go_fuzz_dep_.CoverTab[111476]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:820
			// _ = "end of CoverTab[111476]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:820
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:820
		// _ = "end of CoverTab[111365]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:820
		_go_fuzz_dep_.CoverTab[111366]++
													switch encoding {
		case "bytes":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:822
			_go_fuzz_dep_.CoverTab[111477]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:823
				_go_fuzz_dep_.CoverTab[111482]++
															return makeUnmarshalMessageSlicePtr(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:824
				// _ = "end of CoverTab[111482]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:825
				_go_fuzz_dep_.CoverTab[111483]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:825
				// _ = "end of CoverTab[111483]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:825
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:825
			// _ = "end of CoverTab[111477]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:825
			_go_fuzz_dep_.CoverTab[111478]++
														return makeUnmarshalMessagePtr(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:826
			// _ = "end of CoverTab[111478]"
		case "group":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:827
			_go_fuzz_dep_.CoverTab[111479]++
														if slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:828
				_go_fuzz_dep_.CoverTab[111484]++
															return makeUnmarshalGroupSlicePtr(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:829
				// _ = "end of CoverTab[111484]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:830
				_go_fuzz_dep_.CoverTab[111485]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:830
				// _ = "end of CoverTab[111485]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:830
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:830
			// _ = "end of CoverTab[111479]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:830
			_go_fuzz_dep_.CoverTab[111480]++
														return makeUnmarshalGroupPtr(getUnmarshalInfo(t), name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:831
			// _ = "end of CoverTab[111480]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:831
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:831
			_go_fuzz_dep_.CoverTab[111481]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:831
			// _ = "end of CoverTab[111481]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:832
		// _ = "end of CoverTab[111366]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:832
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:832
		_go_fuzz_dep_.CoverTab[111367]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:832
		// _ = "end of CoverTab[111367]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:833
	// _ = "end of CoverTab[111189]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:833
	_go_fuzz_dep_.CoverTab[111190]++
												panic(fmt.Sprintf("unmarshaler not found type:%s encoding:%s", t, encoding))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:834
	// _ = "end of CoverTab[111190]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:839
func unmarshalInt64Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:839
	_go_fuzz_dep_.CoverTab[111486]++
												if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:840
		_go_fuzz_dep_.CoverTab[111489]++
													return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:841
		// _ = "end of CoverTab[111489]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:842
		_go_fuzz_dep_.CoverTab[111490]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:842
		// _ = "end of CoverTab[111490]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:842
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:842
	// _ = "end of CoverTab[111486]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:842
	_go_fuzz_dep_.CoverTab[111487]++
												x, n := decodeVarint(b)
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:844
		_go_fuzz_dep_.CoverTab[111491]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:845
		// _ = "end of CoverTab[111491]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:846
		_go_fuzz_dep_.CoverTab[111492]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:846
		// _ = "end of CoverTab[111492]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:846
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:846
	// _ = "end of CoverTab[111487]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:846
	_go_fuzz_dep_.CoverTab[111488]++
												b = b[n:]
												v := int64(x)
												*f.toInt64() = v
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:850
	// _ = "end of CoverTab[111488]"
}

func unmarshalInt64Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:853
	_go_fuzz_dep_.CoverTab[111493]++
												if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:854
		_go_fuzz_dep_.CoverTab[111496]++
													return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:855
		// _ = "end of CoverTab[111496]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:856
		_go_fuzz_dep_.CoverTab[111497]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:856
		// _ = "end of CoverTab[111497]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:856
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:856
	// _ = "end of CoverTab[111493]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:856
	_go_fuzz_dep_.CoverTab[111494]++
												x, n := decodeVarint(b)
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:858
		_go_fuzz_dep_.CoverTab[111498]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:859
		// _ = "end of CoverTab[111498]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:860
		_go_fuzz_dep_.CoverTab[111499]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:860
		// _ = "end of CoverTab[111499]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:860
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:860
	// _ = "end of CoverTab[111494]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:860
	_go_fuzz_dep_.CoverTab[111495]++
												b = b[n:]
												v := int64(x)
												*f.toInt64Ptr() = &v
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:864
	// _ = "end of CoverTab[111495]"
}

func unmarshalInt64Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:867
	_go_fuzz_dep_.CoverTab[111500]++
												if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:868
		_go_fuzz_dep_.CoverTab[111504]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:870
			_go_fuzz_dep_.CoverTab[111508]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:871
			// _ = "end of CoverTab[111508]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:872
			_go_fuzz_dep_.CoverTab[111509]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:872
			// _ = "end of CoverTab[111509]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:872
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:872
		// _ = "end of CoverTab[111504]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:872
		_go_fuzz_dep_.CoverTab[111505]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:874
			_go_fuzz_dep_.CoverTab[111510]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:875
			// _ = "end of CoverTab[111510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:876
			_go_fuzz_dep_.CoverTab[111511]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:876
			// _ = "end of CoverTab[111511]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:876
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:876
		// _ = "end of CoverTab[111505]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:876
		_go_fuzz_dep_.CoverTab[111506]++
													res := b[x:]
													b = b[:x]
													for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:879
			_go_fuzz_dep_.CoverTab[111512]++
														x, n = decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:881
				_go_fuzz_dep_.CoverTab[111514]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:882
				// _ = "end of CoverTab[111514]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:883
				_go_fuzz_dep_.CoverTab[111515]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:883
				// _ = "end of CoverTab[111515]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:883
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:883
			// _ = "end of CoverTab[111512]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:883
			_go_fuzz_dep_.CoverTab[111513]++
														b = b[n:]
														v := int64(x)
														s := f.toInt64Slice()
														*s = append(*s, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:887
			// _ = "end of CoverTab[111513]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:888
		// _ = "end of CoverTab[111506]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:888
		_go_fuzz_dep_.CoverTab[111507]++
													return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:889
		// _ = "end of CoverTab[111507]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:890
		_go_fuzz_dep_.CoverTab[111516]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:890
		// _ = "end of CoverTab[111516]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:890
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:890
	// _ = "end of CoverTab[111500]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:890
	_go_fuzz_dep_.CoverTab[111501]++
												if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:891
		_go_fuzz_dep_.CoverTab[111517]++
													return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:892
		// _ = "end of CoverTab[111517]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:893
		_go_fuzz_dep_.CoverTab[111518]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:893
		// _ = "end of CoverTab[111518]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:893
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:893
	// _ = "end of CoverTab[111501]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:893
	_go_fuzz_dep_.CoverTab[111502]++
												x, n := decodeVarint(b)
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:895
		_go_fuzz_dep_.CoverTab[111519]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:896
		// _ = "end of CoverTab[111519]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:897
		_go_fuzz_dep_.CoverTab[111520]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:897
		// _ = "end of CoverTab[111520]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:897
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:897
	// _ = "end of CoverTab[111502]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:897
	_go_fuzz_dep_.CoverTab[111503]++
												b = b[n:]
												v := int64(x)
												s := f.toInt64Slice()
												*s = append(*s, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:902
	// _ = "end of CoverTab[111503]"
}

func unmarshalSint64Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:905
	_go_fuzz_dep_.CoverTab[111521]++
												if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:906
		_go_fuzz_dep_.CoverTab[111524]++
													return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:907
		// _ = "end of CoverTab[111524]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:908
		_go_fuzz_dep_.CoverTab[111525]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:908
		// _ = "end of CoverTab[111525]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:908
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:908
	// _ = "end of CoverTab[111521]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:908
	_go_fuzz_dep_.CoverTab[111522]++
												x, n := decodeVarint(b)
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:910
		_go_fuzz_dep_.CoverTab[111526]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:911
		// _ = "end of CoverTab[111526]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:912
		_go_fuzz_dep_.CoverTab[111527]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:912
		// _ = "end of CoverTab[111527]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:912
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:912
	// _ = "end of CoverTab[111522]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:912
	_go_fuzz_dep_.CoverTab[111523]++
												b = b[n:]
												v := int64(x>>1) ^ int64(x)<<63>>63
												*f.toInt64() = v
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:916
	// _ = "end of CoverTab[111523]"
}

func unmarshalSint64Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:919
	_go_fuzz_dep_.CoverTab[111528]++
												if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:920
		_go_fuzz_dep_.CoverTab[111531]++
													return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:921
		// _ = "end of CoverTab[111531]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:922
		_go_fuzz_dep_.CoverTab[111532]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:922
		// _ = "end of CoverTab[111532]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:922
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:922
	// _ = "end of CoverTab[111528]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:922
	_go_fuzz_dep_.CoverTab[111529]++
												x, n := decodeVarint(b)
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:924
		_go_fuzz_dep_.CoverTab[111533]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:925
		// _ = "end of CoverTab[111533]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:926
		_go_fuzz_dep_.CoverTab[111534]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:926
		// _ = "end of CoverTab[111534]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:926
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:926
	// _ = "end of CoverTab[111529]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:926
	_go_fuzz_dep_.CoverTab[111530]++
												b = b[n:]
												v := int64(x>>1) ^ int64(x)<<63>>63
												*f.toInt64Ptr() = &v
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:930
	// _ = "end of CoverTab[111530]"
}

func unmarshalSint64Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:933
	_go_fuzz_dep_.CoverTab[111535]++
												if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:934
		_go_fuzz_dep_.CoverTab[111539]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:936
			_go_fuzz_dep_.CoverTab[111543]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:937
			// _ = "end of CoverTab[111543]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:938
			_go_fuzz_dep_.CoverTab[111544]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:938
			// _ = "end of CoverTab[111544]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:938
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:938
		// _ = "end of CoverTab[111539]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:938
		_go_fuzz_dep_.CoverTab[111540]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:940
			_go_fuzz_dep_.CoverTab[111545]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:941
			// _ = "end of CoverTab[111545]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:942
			_go_fuzz_dep_.CoverTab[111546]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:942
			// _ = "end of CoverTab[111546]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:942
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:942
		// _ = "end of CoverTab[111540]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:942
		_go_fuzz_dep_.CoverTab[111541]++
													res := b[x:]
													b = b[:x]
													for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:945
			_go_fuzz_dep_.CoverTab[111547]++
														x, n = decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:947
				_go_fuzz_dep_.CoverTab[111549]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:948
				// _ = "end of CoverTab[111549]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:949
				_go_fuzz_dep_.CoverTab[111550]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:949
				// _ = "end of CoverTab[111550]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:949
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:949
			// _ = "end of CoverTab[111547]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:949
			_go_fuzz_dep_.CoverTab[111548]++
														b = b[n:]
														v := int64(x>>1) ^ int64(x)<<63>>63
														s := f.toInt64Slice()
														*s = append(*s, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:953
			// _ = "end of CoverTab[111548]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:954
		// _ = "end of CoverTab[111541]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:954
		_go_fuzz_dep_.CoverTab[111542]++
													return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:955
		// _ = "end of CoverTab[111542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:956
		_go_fuzz_dep_.CoverTab[111551]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:956
		// _ = "end of CoverTab[111551]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:956
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:956
	// _ = "end of CoverTab[111535]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:956
	_go_fuzz_dep_.CoverTab[111536]++
												if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:957
		_go_fuzz_dep_.CoverTab[111552]++
													return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:958
		// _ = "end of CoverTab[111552]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:959
		_go_fuzz_dep_.CoverTab[111553]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:959
		// _ = "end of CoverTab[111553]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:959
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:959
	// _ = "end of CoverTab[111536]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:959
	_go_fuzz_dep_.CoverTab[111537]++
												x, n := decodeVarint(b)
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:961
		_go_fuzz_dep_.CoverTab[111554]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:962
		// _ = "end of CoverTab[111554]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:963
		_go_fuzz_dep_.CoverTab[111555]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:963
		// _ = "end of CoverTab[111555]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:963
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:963
	// _ = "end of CoverTab[111537]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:963
	_go_fuzz_dep_.CoverTab[111538]++
												b = b[n:]
												v := int64(x>>1) ^ int64(x)<<63>>63
												s := f.toInt64Slice()
												*s = append(*s, v)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:968
	// _ = "end of CoverTab[111538]"
}

func unmarshalUint64Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:971
	_go_fuzz_dep_.CoverTab[111556]++
												if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:972
		_go_fuzz_dep_.CoverTab[111559]++
													return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:973
		// _ = "end of CoverTab[111559]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:974
		_go_fuzz_dep_.CoverTab[111560]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:974
		// _ = "end of CoverTab[111560]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:974
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:974
	// _ = "end of CoverTab[111556]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:974
	_go_fuzz_dep_.CoverTab[111557]++
												x, n := decodeVarint(b)
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:976
		_go_fuzz_dep_.CoverTab[111561]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:977
		// _ = "end of CoverTab[111561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:978
		_go_fuzz_dep_.CoverTab[111562]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:978
		// _ = "end of CoverTab[111562]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:978
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:978
	// _ = "end of CoverTab[111557]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:978
	_go_fuzz_dep_.CoverTab[111558]++
												b = b[n:]
												v := uint64(x)
												*f.toUint64() = v
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:982
	// _ = "end of CoverTab[111558]"
}

func unmarshalUint64Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:985
	_go_fuzz_dep_.CoverTab[111563]++
												if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:986
		_go_fuzz_dep_.CoverTab[111566]++
													return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:987
		// _ = "end of CoverTab[111566]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:988
		_go_fuzz_dep_.CoverTab[111567]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:988
		// _ = "end of CoverTab[111567]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:988
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:988
	// _ = "end of CoverTab[111563]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:988
	_go_fuzz_dep_.CoverTab[111564]++
												x, n := decodeVarint(b)
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:990
		_go_fuzz_dep_.CoverTab[111568]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:991
		// _ = "end of CoverTab[111568]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:992
		_go_fuzz_dep_.CoverTab[111569]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:992
		// _ = "end of CoverTab[111569]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:992
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:992
	// _ = "end of CoverTab[111564]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:992
	_go_fuzz_dep_.CoverTab[111565]++
												b = b[n:]
												v := uint64(x)
												*f.toUint64Ptr() = &v
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:996
	// _ = "end of CoverTab[111565]"
}

func unmarshalUint64Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:999
	_go_fuzz_dep_.CoverTab[111570]++
												if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1000
		_go_fuzz_dep_.CoverTab[111574]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1002
			_go_fuzz_dep_.CoverTab[111578]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1003
			// _ = "end of CoverTab[111578]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1004
			_go_fuzz_dep_.CoverTab[111579]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1004
			// _ = "end of CoverTab[111579]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1004
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1004
		// _ = "end of CoverTab[111574]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1004
		_go_fuzz_dep_.CoverTab[111575]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1006
			_go_fuzz_dep_.CoverTab[111580]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1007
			// _ = "end of CoverTab[111580]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1008
			_go_fuzz_dep_.CoverTab[111581]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1008
			// _ = "end of CoverTab[111581]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1008
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1008
		// _ = "end of CoverTab[111575]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1008
		_go_fuzz_dep_.CoverTab[111576]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1011
			_go_fuzz_dep_.CoverTab[111582]++
															x, n = decodeVarint(b)
															if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1013
				_go_fuzz_dep_.CoverTab[111584]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1014
				// _ = "end of CoverTab[111584]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1015
				_go_fuzz_dep_.CoverTab[111585]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1015
				// _ = "end of CoverTab[111585]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1015
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1015
			// _ = "end of CoverTab[111582]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1015
			_go_fuzz_dep_.CoverTab[111583]++
															b = b[n:]
															v := uint64(x)
															s := f.toUint64Slice()
															*s = append(*s, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1019
			// _ = "end of CoverTab[111583]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1020
		// _ = "end of CoverTab[111576]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1020
		_go_fuzz_dep_.CoverTab[111577]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1021
		// _ = "end of CoverTab[111577]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1022
		_go_fuzz_dep_.CoverTab[111586]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1022
		// _ = "end of CoverTab[111586]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1022
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1022
	// _ = "end of CoverTab[111570]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1022
	_go_fuzz_dep_.CoverTab[111571]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1023
		_go_fuzz_dep_.CoverTab[111587]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1024
		// _ = "end of CoverTab[111587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1025
		_go_fuzz_dep_.CoverTab[111588]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1025
		// _ = "end of CoverTab[111588]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1025
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1025
	// _ = "end of CoverTab[111571]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1025
	_go_fuzz_dep_.CoverTab[111572]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1027
		_go_fuzz_dep_.CoverTab[111589]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1028
		// _ = "end of CoverTab[111589]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1029
		_go_fuzz_dep_.CoverTab[111590]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1029
		// _ = "end of CoverTab[111590]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1029
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1029
	// _ = "end of CoverTab[111572]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1029
	_go_fuzz_dep_.CoverTab[111573]++
													b = b[n:]
													v := uint64(x)
													s := f.toUint64Slice()
													*s = append(*s, v)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1034
	// _ = "end of CoverTab[111573]"
}

func unmarshalInt32Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1037
	_go_fuzz_dep_.CoverTab[111591]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1038
		_go_fuzz_dep_.CoverTab[111594]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1039
		// _ = "end of CoverTab[111594]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1040
		_go_fuzz_dep_.CoverTab[111595]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1040
		// _ = "end of CoverTab[111595]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1040
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1040
	// _ = "end of CoverTab[111591]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1040
	_go_fuzz_dep_.CoverTab[111592]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1042
		_go_fuzz_dep_.CoverTab[111596]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1043
		// _ = "end of CoverTab[111596]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1044
		_go_fuzz_dep_.CoverTab[111597]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1044
		// _ = "end of CoverTab[111597]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1044
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1044
	// _ = "end of CoverTab[111592]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1044
	_go_fuzz_dep_.CoverTab[111593]++
													b = b[n:]
													v := int32(x)
													*f.toInt32() = v
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1048
	// _ = "end of CoverTab[111593]"
}

func unmarshalInt32Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1051
	_go_fuzz_dep_.CoverTab[111598]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1052
		_go_fuzz_dep_.CoverTab[111601]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1053
		// _ = "end of CoverTab[111601]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1054
		_go_fuzz_dep_.CoverTab[111602]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1054
		// _ = "end of CoverTab[111602]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1054
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1054
	// _ = "end of CoverTab[111598]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1054
	_go_fuzz_dep_.CoverTab[111599]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1056
		_go_fuzz_dep_.CoverTab[111603]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1057
		// _ = "end of CoverTab[111603]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1058
		_go_fuzz_dep_.CoverTab[111604]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1058
		// _ = "end of CoverTab[111604]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1058
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1058
	// _ = "end of CoverTab[111599]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1058
	_go_fuzz_dep_.CoverTab[111600]++
													b = b[n:]
													v := int32(x)
													f.setInt32Ptr(v)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1062
	// _ = "end of CoverTab[111600]"
}

func unmarshalInt32Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1065
	_go_fuzz_dep_.CoverTab[111605]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1066
		_go_fuzz_dep_.CoverTab[111609]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1068
			_go_fuzz_dep_.CoverTab[111613]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1069
			// _ = "end of CoverTab[111613]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1070
			_go_fuzz_dep_.CoverTab[111614]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1070
			// _ = "end of CoverTab[111614]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1070
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1070
		// _ = "end of CoverTab[111609]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1070
		_go_fuzz_dep_.CoverTab[111610]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1072
			_go_fuzz_dep_.CoverTab[111615]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1073
			// _ = "end of CoverTab[111615]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1074
			_go_fuzz_dep_.CoverTab[111616]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1074
			// _ = "end of CoverTab[111616]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1074
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1074
		// _ = "end of CoverTab[111610]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1074
		_go_fuzz_dep_.CoverTab[111611]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1077
			_go_fuzz_dep_.CoverTab[111617]++
															x, n = decodeVarint(b)
															if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1079
				_go_fuzz_dep_.CoverTab[111619]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1080
				// _ = "end of CoverTab[111619]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1081
				_go_fuzz_dep_.CoverTab[111620]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1081
				// _ = "end of CoverTab[111620]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1081
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1081
			// _ = "end of CoverTab[111617]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1081
			_go_fuzz_dep_.CoverTab[111618]++
															b = b[n:]
															v := int32(x)
															f.appendInt32Slice(v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1084
			// _ = "end of CoverTab[111618]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1085
		// _ = "end of CoverTab[111611]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1085
		_go_fuzz_dep_.CoverTab[111612]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1086
		// _ = "end of CoverTab[111612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1087
		_go_fuzz_dep_.CoverTab[111621]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1087
		// _ = "end of CoverTab[111621]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1087
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1087
	// _ = "end of CoverTab[111605]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1087
	_go_fuzz_dep_.CoverTab[111606]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1088
		_go_fuzz_dep_.CoverTab[111622]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1089
		// _ = "end of CoverTab[111622]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1090
		_go_fuzz_dep_.CoverTab[111623]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1090
		// _ = "end of CoverTab[111623]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1090
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1090
	// _ = "end of CoverTab[111606]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1090
	_go_fuzz_dep_.CoverTab[111607]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1092
		_go_fuzz_dep_.CoverTab[111624]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1093
		// _ = "end of CoverTab[111624]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1094
		_go_fuzz_dep_.CoverTab[111625]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1094
		// _ = "end of CoverTab[111625]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1094
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1094
	// _ = "end of CoverTab[111607]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1094
	_go_fuzz_dep_.CoverTab[111608]++
													b = b[n:]
													v := int32(x)
													f.appendInt32Slice(v)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1098
	// _ = "end of CoverTab[111608]"
}

func unmarshalSint32Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1101
	_go_fuzz_dep_.CoverTab[111626]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1102
		_go_fuzz_dep_.CoverTab[111629]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1103
		// _ = "end of CoverTab[111629]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1104
		_go_fuzz_dep_.CoverTab[111630]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1104
		// _ = "end of CoverTab[111630]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1104
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1104
	// _ = "end of CoverTab[111626]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1104
	_go_fuzz_dep_.CoverTab[111627]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1106
		_go_fuzz_dep_.CoverTab[111631]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1107
		// _ = "end of CoverTab[111631]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1108
		_go_fuzz_dep_.CoverTab[111632]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1108
		// _ = "end of CoverTab[111632]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1108
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1108
	// _ = "end of CoverTab[111627]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1108
	_go_fuzz_dep_.CoverTab[111628]++
													b = b[n:]
													v := int32(x>>1) ^ int32(x)<<31>>31
													*f.toInt32() = v
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1112
	// _ = "end of CoverTab[111628]"
}

func unmarshalSint32Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1115
	_go_fuzz_dep_.CoverTab[111633]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1116
		_go_fuzz_dep_.CoverTab[111636]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1117
		// _ = "end of CoverTab[111636]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1118
		_go_fuzz_dep_.CoverTab[111637]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1118
		// _ = "end of CoverTab[111637]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1118
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1118
	// _ = "end of CoverTab[111633]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1118
	_go_fuzz_dep_.CoverTab[111634]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1120
		_go_fuzz_dep_.CoverTab[111638]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1121
		// _ = "end of CoverTab[111638]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1122
		_go_fuzz_dep_.CoverTab[111639]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1122
		// _ = "end of CoverTab[111639]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1122
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1122
	// _ = "end of CoverTab[111634]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1122
	_go_fuzz_dep_.CoverTab[111635]++
													b = b[n:]
													v := int32(x>>1) ^ int32(x)<<31>>31
													f.setInt32Ptr(v)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1126
	// _ = "end of CoverTab[111635]"
}

func unmarshalSint32Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1129
	_go_fuzz_dep_.CoverTab[111640]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1130
		_go_fuzz_dep_.CoverTab[111644]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1132
			_go_fuzz_dep_.CoverTab[111648]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1133
			// _ = "end of CoverTab[111648]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1134
			_go_fuzz_dep_.CoverTab[111649]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1134
			// _ = "end of CoverTab[111649]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1134
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1134
		// _ = "end of CoverTab[111644]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1134
		_go_fuzz_dep_.CoverTab[111645]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1136
			_go_fuzz_dep_.CoverTab[111650]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1137
			// _ = "end of CoverTab[111650]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1138
			_go_fuzz_dep_.CoverTab[111651]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1138
			// _ = "end of CoverTab[111651]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1138
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1138
		// _ = "end of CoverTab[111645]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1138
		_go_fuzz_dep_.CoverTab[111646]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1141
			_go_fuzz_dep_.CoverTab[111652]++
															x, n = decodeVarint(b)
															if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1143
				_go_fuzz_dep_.CoverTab[111654]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1144
				// _ = "end of CoverTab[111654]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1145
				_go_fuzz_dep_.CoverTab[111655]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1145
				// _ = "end of CoverTab[111655]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1145
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1145
			// _ = "end of CoverTab[111652]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1145
			_go_fuzz_dep_.CoverTab[111653]++
															b = b[n:]
															v := int32(x>>1) ^ int32(x)<<31>>31
															f.appendInt32Slice(v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1148
			// _ = "end of CoverTab[111653]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1149
		// _ = "end of CoverTab[111646]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1149
		_go_fuzz_dep_.CoverTab[111647]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1150
		// _ = "end of CoverTab[111647]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1151
		_go_fuzz_dep_.CoverTab[111656]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1151
		// _ = "end of CoverTab[111656]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1151
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1151
	// _ = "end of CoverTab[111640]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1151
	_go_fuzz_dep_.CoverTab[111641]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1152
		_go_fuzz_dep_.CoverTab[111657]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1153
		// _ = "end of CoverTab[111657]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1154
		_go_fuzz_dep_.CoverTab[111658]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1154
		// _ = "end of CoverTab[111658]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1154
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1154
	// _ = "end of CoverTab[111641]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1154
	_go_fuzz_dep_.CoverTab[111642]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1156
		_go_fuzz_dep_.CoverTab[111659]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1157
		// _ = "end of CoverTab[111659]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1158
		_go_fuzz_dep_.CoverTab[111660]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1158
		// _ = "end of CoverTab[111660]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1158
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1158
	// _ = "end of CoverTab[111642]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1158
	_go_fuzz_dep_.CoverTab[111643]++
													b = b[n:]
													v := int32(x>>1) ^ int32(x)<<31>>31
													f.appendInt32Slice(v)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1162
	// _ = "end of CoverTab[111643]"
}

func unmarshalUint32Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1165
	_go_fuzz_dep_.CoverTab[111661]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1166
		_go_fuzz_dep_.CoverTab[111664]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1167
		// _ = "end of CoverTab[111664]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1168
		_go_fuzz_dep_.CoverTab[111665]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1168
		// _ = "end of CoverTab[111665]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1168
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1168
	// _ = "end of CoverTab[111661]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1168
	_go_fuzz_dep_.CoverTab[111662]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1170
		_go_fuzz_dep_.CoverTab[111666]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1171
		// _ = "end of CoverTab[111666]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1172
		_go_fuzz_dep_.CoverTab[111667]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1172
		// _ = "end of CoverTab[111667]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1172
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1172
	// _ = "end of CoverTab[111662]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1172
	_go_fuzz_dep_.CoverTab[111663]++
													b = b[n:]
													v := uint32(x)
													*f.toUint32() = v
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1176
	// _ = "end of CoverTab[111663]"
}

func unmarshalUint32Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1179
	_go_fuzz_dep_.CoverTab[111668]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1180
		_go_fuzz_dep_.CoverTab[111671]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1181
		// _ = "end of CoverTab[111671]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1182
		_go_fuzz_dep_.CoverTab[111672]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1182
		// _ = "end of CoverTab[111672]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1182
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1182
	// _ = "end of CoverTab[111668]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1182
	_go_fuzz_dep_.CoverTab[111669]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1184
		_go_fuzz_dep_.CoverTab[111673]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1185
		// _ = "end of CoverTab[111673]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1186
		_go_fuzz_dep_.CoverTab[111674]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1186
		// _ = "end of CoverTab[111674]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1186
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1186
	// _ = "end of CoverTab[111669]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1186
	_go_fuzz_dep_.CoverTab[111670]++
													b = b[n:]
													v := uint32(x)
													*f.toUint32Ptr() = &v
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1190
	// _ = "end of CoverTab[111670]"
}

func unmarshalUint32Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1193
	_go_fuzz_dep_.CoverTab[111675]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1194
		_go_fuzz_dep_.CoverTab[111679]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1196
			_go_fuzz_dep_.CoverTab[111683]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1197
			// _ = "end of CoverTab[111683]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1198
			_go_fuzz_dep_.CoverTab[111684]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1198
			// _ = "end of CoverTab[111684]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1198
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1198
		// _ = "end of CoverTab[111679]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1198
		_go_fuzz_dep_.CoverTab[111680]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1200
			_go_fuzz_dep_.CoverTab[111685]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1201
			// _ = "end of CoverTab[111685]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1202
			_go_fuzz_dep_.CoverTab[111686]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1202
			// _ = "end of CoverTab[111686]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1202
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1202
		// _ = "end of CoverTab[111680]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1202
		_go_fuzz_dep_.CoverTab[111681]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1205
			_go_fuzz_dep_.CoverTab[111687]++
															x, n = decodeVarint(b)
															if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1207
				_go_fuzz_dep_.CoverTab[111689]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1208
				// _ = "end of CoverTab[111689]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1209
				_go_fuzz_dep_.CoverTab[111690]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1209
				// _ = "end of CoverTab[111690]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1209
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1209
			// _ = "end of CoverTab[111687]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1209
			_go_fuzz_dep_.CoverTab[111688]++
															b = b[n:]
															v := uint32(x)
															s := f.toUint32Slice()
															*s = append(*s, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1213
			// _ = "end of CoverTab[111688]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1214
		// _ = "end of CoverTab[111681]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1214
		_go_fuzz_dep_.CoverTab[111682]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1215
		// _ = "end of CoverTab[111682]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1216
		_go_fuzz_dep_.CoverTab[111691]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1216
		// _ = "end of CoverTab[111691]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1216
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1216
	// _ = "end of CoverTab[111675]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1216
	_go_fuzz_dep_.CoverTab[111676]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1217
		_go_fuzz_dep_.CoverTab[111692]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1218
		// _ = "end of CoverTab[111692]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1219
		_go_fuzz_dep_.CoverTab[111693]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1219
		// _ = "end of CoverTab[111693]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1219
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1219
	// _ = "end of CoverTab[111676]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1219
	_go_fuzz_dep_.CoverTab[111677]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1221
		_go_fuzz_dep_.CoverTab[111694]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1222
		// _ = "end of CoverTab[111694]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1223
		_go_fuzz_dep_.CoverTab[111695]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1223
		// _ = "end of CoverTab[111695]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1223
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1223
	// _ = "end of CoverTab[111677]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1223
	_go_fuzz_dep_.CoverTab[111678]++
													b = b[n:]
													v := uint32(x)
													s := f.toUint32Slice()
													*s = append(*s, v)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1228
	// _ = "end of CoverTab[111678]"
}

func unmarshalFixed64Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1231
	_go_fuzz_dep_.CoverTab[111696]++
													if w != WireFixed64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1232
		_go_fuzz_dep_.CoverTab[111699]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1233
		// _ = "end of CoverTab[111699]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1234
		_go_fuzz_dep_.CoverTab[111700]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1234
		// _ = "end of CoverTab[111700]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1234
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1234
	// _ = "end of CoverTab[111696]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1234
	_go_fuzz_dep_.CoverTab[111697]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1235
		_go_fuzz_dep_.CoverTab[111701]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1236
		// _ = "end of CoverTab[111701]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1237
		_go_fuzz_dep_.CoverTab[111702]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1237
		// _ = "end of CoverTab[111702]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1237
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1237
	// _ = "end of CoverTab[111697]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1237
	_go_fuzz_dep_.CoverTab[111698]++
													v := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
													*f.toUint64() = v
													return b[8:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1240
	// _ = "end of CoverTab[111698]"
}

func unmarshalFixed64Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1243
	_go_fuzz_dep_.CoverTab[111703]++
													if w != WireFixed64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1244
		_go_fuzz_dep_.CoverTab[111706]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1245
		// _ = "end of CoverTab[111706]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1246
		_go_fuzz_dep_.CoverTab[111707]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1246
		// _ = "end of CoverTab[111707]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1246
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1246
	// _ = "end of CoverTab[111703]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1246
	_go_fuzz_dep_.CoverTab[111704]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1247
		_go_fuzz_dep_.CoverTab[111708]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1248
		// _ = "end of CoverTab[111708]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1249
		_go_fuzz_dep_.CoverTab[111709]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1249
		// _ = "end of CoverTab[111709]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1249
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1249
	// _ = "end of CoverTab[111704]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1249
	_go_fuzz_dep_.CoverTab[111705]++
													v := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
													*f.toUint64Ptr() = &v
													return b[8:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1252
	// _ = "end of CoverTab[111705]"
}

func unmarshalFixed64Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1255
	_go_fuzz_dep_.CoverTab[111710]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1256
		_go_fuzz_dep_.CoverTab[111714]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1258
			_go_fuzz_dep_.CoverTab[111718]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1259
			// _ = "end of CoverTab[111718]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1260
			_go_fuzz_dep_.CoverTab[111719]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1260
			// _ = "end of CoverTab[111719]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1260
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1260
		// _ = "end of CoverTab[111714]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1260
		_go_fuzz_dep_.CoverTab[111715]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1262
			_go_fuzz_dep_.CoverTab[111720]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1263
			// _ = "end of CoverTab[111720]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1264
			_go_fuzz_dep_.CoverTab[111721]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1264
			// _ = "end of CoverTab[111721]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1264
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1264
		// _ = "end of CoverTab[111715]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1264
		_go_fuzz_dep_.CoverTab[111716]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1267
			_go_fuzz_dep_.CoverTab[111722]++
															if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1268
				_go_fuzz_dep_.CoverTab[111724]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1269
				// _ = "end of CoverTab[111724]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1270
				_go_fuzz_dep_.CoverTab[111725]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1270
				// _ = "end of CoverTab[111725]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1270
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1270
			// _ = "end of CoverTab[111722]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1270
			_go_fuzz_dep_.CoverTab[111723]++
															v := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
															s := f.toUint64Slice()
															*s = append(*s, v)
															b = b[8:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1274
			// _ = "end of CoverTab[111723]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1275
		// _ = "end of CoverTab[111716]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1275
		_go_fuzz_dep_.CoverTab[111717]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1276
		// _ = "end of CoverTab[111717]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1277
		_go_fuzz_dep_.CoverTab[111726]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1277
		// _ = "end of CoverTab[111726]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1277
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1277
	// _ = "end of CoverTab[111710]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1277
	_go_fuzz_dep_.CoverTab[111711]++
													if w != WireFixed64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1278
		_go_fuzz_dep_.CoverTab[111727]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1279
		// _ = "end of CoverTab[111727]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1280
		_go_fuzz_dep_.CoverTab[111728]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1280
		// _ = "end of CoverTab[111728]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1280
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1280
	// _ = "end of CoverTab[111711]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1280
	_go_fuzz_dep_.CoverTab[111712]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1281
		_go_fuzz_dep_.CoverTab[111729]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1282
		// _ = "end of CoverTab[111729]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1283
		_go_fuzz_dep_.CoverTab[111730]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1283
		// _ = "end of CoverTab[111730]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1283
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1283
	// _ = "end of CoverTab[111712]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1283
	_go_fuzz_dep_.CoverTab[111713]++
													v := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
													s := f.toUint64Slice()
													*s = append(*s, v)
													return b[8:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1287
	// _ = "end of CoverTab[111713]"
}

func unmarshalFixedS64Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1290
	_go_fuzz_dep_.CoverTab[111731]++
													if w != WireFixed64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1291
		_go_fuzz_dep_.CoverTab[111734]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1292
		// _ = "end of CoverTab[111734]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1293
		_go_fuzz_dep_.CoverTab[111735]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1293
		// _ = "end of CoverTab[111735]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1293
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1293
	// _ = "end of CoverTab[111731]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1293
	_go_fuzz_dep_.CoverTab[111732]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1294
		_go_fuzz_dep_.CoverTab[111736]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1295
		// _ = "end of CoverTab[111736]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1296
		_go_fuzz_dep_.CoverTab[111737]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1296
		// _ = "end of CoverTab[111737]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1296
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1296
	// _ = "end of CoverTab[111732]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1296
	_go_fuzz_dep_.CoverTab[111733]++
													v := int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
													*f.toInt64() = v
													return b[8:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1299
	// _ = "end of CoverTab[111733]"
}

func unmarshalFixedS64Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1302
	_go_fuzz_dep_.CoverTab[111738]++
													if w != WireFixed64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1303
		_go_fuzz_dep_.CoverTab[111741]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1304
		// _ = "end of CoverTab[111741]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1305
		_go_fuzz_dep_.CoverTab[111742]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1305
		// _ = "end of CoverTab[111742]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1305
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1305
	// _ = "end of CoverTab[111738]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1305
	_go_fuzz_dep_.CoverTab[111739]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1306
		_go_fuzz_dep_.CoverTab[111743]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1307
		// _ = "end of CoverTab[111743]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1308
		_go_fuzz_dep_.CoverTab[111744]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1308
		// _ = "end of CoverTab[111744]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1308
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1308
	// _ = "end of CoverTab[111739]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1308
	_go_fuzz_dep_.CoverTab[111740]++
													v := int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
													*f.toInt64Ptr() = &v
													return b[8:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1311
	// _ = "end of CoverTab[111740]"
}

func unmarshalFixedS64Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1314
	_go_fuzz_dep_.CoverTab[111745]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1315
		_go_fuzz_dep_.CoverTab[111749]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1317
			_go_fuzz_dep_.CoverTab[111753]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1318
			// _ = "end of CoverTab[111753]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1319
			_go_fuzz_dep_.CoverTab[111754]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1319
			// _ = "end of CoverTab[111754]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1319
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1319
		// _ = "end of CoverTab[111749]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1319
		_go_fuzz_dep_.CoverTab[111750]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1321
			_go_fuzz_dep_.CoverTab[111755]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1322
			// _ = "end of CoverTab[111755]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1323
			_go_fuzz_dep_.CoverTab[111756]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1323
			// _ = "end of CoverTab[111756]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1323
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1323
		// _ = "end of CoverTab[111750]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1323
		_go_fuzz_dep_.CoverTab[111751]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1326
			_go_fuzz_dep_.CoverTab[111757]++
															if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1327
				_go_fuzz_dep_.CoverTab[111759]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1328
				// _ = "end of CoverTab[111759]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1329
				_go_fuzz_dep_.CoverTab[111760]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1329
				// _ = "end of CoverTab[111760]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1329
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1329
			// _ = "end of CoverTab[111757]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1329
			_go_fuzz_dep_.CoverTab[111758]++
															v := int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
															s := f.toInt64Slice()
															*s = append(*s, v)
															b = b[8:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1333
			// _ = "end of CoverTab[111758]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1334
		// _ = "end of CoverTab[111751]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1334
		_go_fuzz_dep_.CoverTab[111752]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1335
		// _ = "end of CoverTab[111752]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1336
		_go_fuzz_dep_.CoverTab[111761]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1336
		// _ = "end of CoverTab[111761]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1336
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1336
	// _ = "end of CoverTab[111745]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1336
	_go_fuzz_dep_.CoverTab[111746]++
													if w != WireFixed64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1337
		_go_fuzz_dep_.CoverTab[111762]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1338
		// _ = "end of CoverTab[111762]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1339
		_go_fuzz_dep_.CoverTab[111763]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1339
		// _ = "end of CoverTab[111763]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1339
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1339
	// _ = "end of CoverTab[111746]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1339
	_go_fuzz_dep_.CoverTab[111747]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1340
		_go_fuzz_dep_.CoverTab[111764]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1341
		// _ = "end of CoverTab[111764]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1342
		_go_fuzz_dep_.CoverTab[111765]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1342
		// _ = "end of CoverTab[111765]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1342
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1342
	// _ = "end of CoverTab[111747]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1342
	_go_fuzz_dep_.CoverTab[111748]++
													v := int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
													s := f.toInt64Slice()
													*s = append(*s, v)
													return b[8:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1346
	// _ = "end of CoverTab[111748]"
}

func unmarshalFixed32Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1349
	_go_fuzz_dep_.CoverTab[111766]++
													if w != WireFixed32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1350
		_go_fuzz_dep_.CoverTab[111769]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1351
		// _ = "end of CoverTab[111769]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1352
		_go_fuzz_dep_.CoverTab[111770]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1352
		// _ = "end of CoverTab[111770]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1352
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1352
	// _ = "end of CoverTab[111766]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1352
	_go_fuzz_dep_.CoverTab[111767]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1353
		_go_fuzz_dep_.CoverTab[111771]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1354
		// _ = "end of CoverTab[111771]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1355
		_go_fuzz_dep_.CoverTab[111772]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1355
		// _ = "end of CoverTab[111772]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1355
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1355
	// _ = "end of CoverTab[111767]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1355
	_go_fuzz_dep_.CoverTab[111768]++
													v := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
													*f.toUint32() = v
													return b[4:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1358
	// _ = "end of CoverTab[111768]"
}

func unmarshalFixed32Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1361
	_go_fuzz_dep_.CoverTab[111773]++
													if w != WireFixed32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1362
		_go_fuzz_dep_.CoverTab[111776]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1363
		// _ = "end of CoverTab[111776]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1364
		_go_fuzz_dep_.CoverTab[111777]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1364
		// _ = "end of CoverTab[111777]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1364
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1364
	// _ = "end of CoverTab[111773]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1364
	_go_fuzz_dep_.CoverTab[111774]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1365
		_go_fuzz_dep_.CoverTab[111778]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1366
		// _ = "end of CoverTab[111778]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1367
		_go_fuzz_dep_.CoverTab[111779]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1367
		// _ = "end of CoverTab[111779]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1367
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1367
	// _ = "end of CoverTab[111774]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1367
	_go_fuzz_dep_.CoverTab[111775]++
													v := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
													*f.toUint32Ptr() = &v
													return b[4:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1370
	// _ = "end of CoverTab[111775]"
}

func unmarshalFixed32Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1373
	_go_fuzz_dep_.CoverTab[111780]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1374
		_go_fuzz_dep_.CoverTab[111784]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1376
			_go_fuzz_dep_.CoverTab[111788]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1377
			// _ = "end of CoverTab[111788]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1378
			_go_fuzz_dep_.CoverTab[111789]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1378
			// _ = "end of CoverTab[111789]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1378
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1378
		// _ = "end of CoverTab[111784]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1378
		_go_fuzz_dep_.CoverTab[111785]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1380
			_go_fuzz_dep_.CoverTab[111790]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1381
			// _ = "end of CoverTab[111790]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1382
			_go_fuzz_dep_.CoverTab[111791]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1382
			// _ = "end of CoverTab[111791]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1382
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1382
		// _ = "end of CoverTab[111785]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1382
		_go_fuzz_dep_.CoverTab[111786]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1385
			_go_fuzz_dep_.CoverTab[111792]++
															if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1386
				_go_fuzz_dep_.CoverTab[111794]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1387
				// _ = "end of CoverTab[111794]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1388
				_go_fuzz_dep_.CoverTab[111795]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1388
				// _ = "end of CoverTab[111795]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1388
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1388
			// _ = "end of CoverTab[111792]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1388
			_go_fuzz_dep_.CoverTab[111793]++
															v := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
															s := f.toUint32Slice()
															*s = append(*s, v)
															b = b[4:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1392
			// _ = "end of CoverTab[111793]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1393
		// _ = "end of CoverTab[111786]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1393
		_go_fuzz_dep_.CoverTab[111787]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1394
		// _ = "end of CoverTab[111787]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1395
		_go_fuzz_dep_.CoverTab[111796]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1395
		// _ = "end of CoverTab[111796]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1395
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1395
	// _ = "end of CoverTab[111780]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1395
	_go_fuzz_dep_.CoverTab[111781]++
													if w != WireFixed32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1396
		_go_fuzz_dep_.CoverTab[111797]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1397
		// _ = "end of CoverTab[111797]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1398
		_go_fuzz_dep_.CoverTab[111798]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1398
		// _ = "end of CoverTab[111798]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1398
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1398
	// _ = "end of CoverTab[111781]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1398
	_go_fuzz_dep_.CoverTab[111782]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1399
		_go_fuzz_dep_.CoverTab[111799]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1400
		// _ = "end of CoverTab[111799]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1401
		_go_fuzz_dep_.CoverTab[111800]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1401
		// _ = "end of CoverTab[111800]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1401
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1401
	// _ = "end of CoverTab[111782]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1401
	_go_fuzz_dep_.CoverTab[111783]++
													v := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
													s := f.toUint32Slice()
													*s = append(*s, v)
													return b[4:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1405
	// _ = "end of CoverTab[111783]"
}

func unmarshalFixedS32Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1408
	_go_fuzz_dep_.CoverTab[111801]++
													if w != WireFixed32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1409
		_go_fuzz_dep_.CoverTab[111804]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1410
		// _ = "end of CoverTab[111804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1411
		_go_fuzz_dep_.CoverTab[111805]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1411
		// _ = "end of CoverTab[111805]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1411
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1411
	// _ = "end of CoverTab[111801]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1411
	_go_fuzz_dep_.CoverTab[111802]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1412
		_go_fuzz_dep_.CoverTab[111806]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1413
		// _ = "end of CoverTab[111806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1414
		_go_fuzz_dep_.CoverTab[111807]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1414
		// _ = "end of CoverTab[111807]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1414
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1414
	// _ = "end of CoverTab[111802]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1414
	_go_fuzz_dep_.CoverTab[111803]++
													v := int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
													*f.toInt32() = v
													return b[4:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1417
	// _ = "end of CoverTab[111803]"
}

func unmarshalFixedS32Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1420
	_go_fuzz_dep_.CoverTab[111808]++
													if w != WireFixed32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1421
		_go_fuzz_dep_.CoverTab[111811]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1422
		// _ = "end of CoverTab[111811]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1423
		_go_fuzz_dep_.CoverTab[111812]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1423
		// _ = "end of CoverTab[111812]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1423
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1423
	// _ = "end of CoverTab[111808]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1423
	_go_fuzz_dep_.CoverTab[111809]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1424
		_go_fuzz_dep_.CoverTab[111813]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1425
		// _ = "end of CoverTab[111813]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1426
		_go_fuzz_dep_.CoverTab[111814]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1426
		// _ = "end of CoverTab[111814]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1426
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1426
	// _ = "end of CoverTab[111809]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1426
	_go_fuzz_dep_.CoverTab[111810]++
													v := int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
													f.setInt32Ptr(v)
													return b[4:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1429
	// _ = "end of CoverTab[111810]"
}

func unmarshalFixedS32Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1432
	_go_fuzz_dep_.CoverTab[111815]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1433
		_go_fuzz_dep_.CoverTab[111819]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1435
			_go_fuzz_dep_.CoverTab[111823]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1436
			// _ = "end of CoverTab[111823]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1437
			_go_fuzz_dep_.CoverTab[111824]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1437
			// _ = "end of CoverTab[111824]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1437
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1437
		// _ = "end of CoverTab[111819]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1437
		_go_fuzz_dep_.CoverTab[111820]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1439
			_go_fuzz_dep_.CoverTab[111825]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1440
			// _ = "end of CoverTab[111825]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1441
			_go_fuzz_dep_.CoverTab[111826]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1441
			// _ = "end of CoverTab[111826]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1441
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1441
		// _ = "end of CoverTab[111820]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1441
		_go_fuzz_dep_.CoverTab[111821]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1444
			_go_fuzz_dep_.CoverTab[111827]++
															if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1445
				_go_fuzz_dep_.CoverTab[111829]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1446
				// _ = "end of CoverTab[111829]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1447
				_go_fuzz_dep_.CoverTab[111830]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1447
				// _ = "end of CoverTab[111830]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1447
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1447
			// _ = "end of CoverTab[111827]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1447
			_go_fuzz_dep_.CoverTab[111828]++
															v := int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
															f.appendInt32Slice(v)
															b = b[4:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1450
			// _ = "end of CoverTab[111828]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1451
		// _ = "end of CoverTab[111821]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1451
		_go_fuzz_dep_.CoverTab[111822]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1452
		// _ = "end of CoverTab[111822]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1453
		_go_fuzz_dep_.CoverTab[111831]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1453
		// _ = "end of CoverTab[111831]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1453
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1453
	// _ = "end of CoverTab[111815]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1453
	_go_fuzz_dep_.CoverTab[111816]++
													if w != WireFixed32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1454
		_go_fuzz_dep_.CoverTab[111832]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1455
		// _ = "end of CoverTab[111832]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1456
		_go_fuzz_dep_.CoverTab[111833]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1456
		// _ = "end of CoverTab[111833]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1456
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1456
	// _ = "end of CoverTab[111816]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1456
	_go_fuzz_dep_.CoverTab[111817]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1457
		_go_fuzz_dep_.CoverTab[111834]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1458
		// _ = "end of CoverTab[111834]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1459
		_go_fuzz_dep_.CoverTab[111835]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1459
		// _ = "end of CoverTab[111835]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1459
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1459
	// _ = "end of CoverTab[111817]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1459
	_go_fuzz_dep_.CoverTab[111818]++
													v := int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
													f.appendInt32Slice(v)
													return b[4:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1462
	// _ = "end of CoverTab[111818]"
}

func unmarshalBoolValue(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1465
	_go_fuzz_dep_.CoverTab[111836]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1466
		_go_fuzz_dep_.CoverTab[111839]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1467
		// _ = "end of CoverTab[111839]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1468
		_go_fuzz_dep_.CoverTab[111840]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1468
		// _ = "end of CoverTab[111840]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1468
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1468
	// _ = "end of CoverTab[111836]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1468
	_go_fuzz_dep_.CoverTab[111837]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1472
	x, n := decodeVarint(b)
	if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1473
		_go_fuzz_dep_.CoverTab[111841]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1474
		// _ = "end of CoverTab[111841]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1475
		_go_fuzz_dep_.CoverTab[111842]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1475
		// _ = "end of CoverTab[111842]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1475
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1475
	// _ = "end of CoverTab[111837]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1475
	_go_fuzz_dep_.CoverTab[111838]++

													v := x != 0
													*f.toBool() = v
													return b[n:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1479
	// _ = "end of CoverTab[111838]"
}

func unmarshalBoolPtr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1482
	_go_fuzz_dep_.CoverTab[111843]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1483
		_go_fuzz_dep_.CoverTab[111846]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1484
		// _ = "end of CoverTab[111846]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1485
		_go_fuzz_dep_.CoverTab[111847]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1485
		// _ = "end of CoverTab[111847]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1485
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1485
	// _ = "end of CoverTab[111843]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1485
	_go_fuzz_dep_.CoverTab[111844]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1487
		_go_fuzz_dep_.CoverTab[111848]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1488
		// _ = "end of CoverTab[111848]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1489
		_go_fuzz_dep_.CoverTab[111849]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1489
		// _ = "end of CoverTab[111849]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1489
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1489
	// _ = "end of CoverTab[111844]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1489
	_go_fuzz_dep_.CoverTab[111845]++
													v := x != 0
													*f.toBoolPtr() = &v
													return b[n:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1492
	// _ = "end of CoverTab[111845]"
}

func unmarshalBoolSlice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1495
	_go_fuzz_dep_.CoverTab[111850]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1496
		_go_fuzz_dep_.CoverTab[111854]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1498
			_go_fuzz_dep_.CoverTab[111858]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1499
			// _ = "end of CoverTab[111858]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1500
			_go_fuzz_dep_.CoverTab[111859]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1500
			// _ = "end of CoverTab[111859]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1500
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1500
		// _ = "end of CoverTab[111854]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1500
		_go_fuzz_dep_.CoverTab[111855]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1502
			_go_fuzz_dep_.CoverTab[111860]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1503
			// _ = "end of CoverTab[111860]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1504
			_go_fuzz_dep_.CoverTab[111861]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1504
			// _ = "end of CoverTab[111861]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1504
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1504
		// _ = "end of CoverTab[111855]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1504
		_go_fuzz_dep_.CoverTab[111856]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1507
			_go_fuzz_dep_.CoverTab[111862]++
															x, n = decodeVarint(b)
															if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1509
				_go_fuzz_dep_.CoverTab[111864]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1510
				// _ = "end of CoverTab[111864]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1511
				_go_fuzz_dep_.CoverTab[111865]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1511
				// _ = "end of CoverTab[111865]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1511
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1511
			// _ = "end of CoverTab[111862]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1511
			_go_fuzz_dep_.CoverTab[111863]++
															v := x != 0
															s := f.toBoolSlice()
															*s = append(*s, v)
															b = b[n:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1515
			// _ = "end of CoverTab[111863]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1516
		// _ = "end of CoverTab[111856]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1516
		_go_fuzz_dep_.CoverTab[111857]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1517
		// _ = "end of CoverTab[111857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1518
		_go_fuzz_dep_.CoverTab[111866]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1518
		// _ = "end of CoverTab[111866]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1518
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1518
	// _ = "end of CoverTab[111850]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1518
	_go_fuzz_dep_.CoverTab[111851]++
													if w != WireVarint {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1519
		_go_fuzz_dep_.CoverTab[111867]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1520
		// _ = "end of CoverTab[111867]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1521
		_go_fuzz_dep_.CoverTab[111868]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1521
		// _ = "end of CoverTab[111868]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1521
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1521
	// _ = "end of CoverTab[111851]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1521
	_go_fuzz_dep_.CoverTab[111852]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1523
		_go_fuzz_dep_.CoverTab[111869]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1524
		// _ = "end of CoverTab[111869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1525
		_go_fuzz_dep_.CoverTab[111870]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1525
		// _ = "end of CoverTab[111870]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1525
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1525
	// _ = "end of CoverTab[111852]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1525
	_go_fuzz_dep_.CoverTab[111853]++
													v := x != 0
													s := f.toBoolSlice()
													*s = append(*s, v)
													return b[n:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1529
	// _ = "end of CoverTab[111853]"
}

func unmarshalFloat64Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1532
	_go_fuzz_dep_.CoverTab[111871]++
													if w != WireFixed64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1533
		_go_fuzz_dep_.CoverTab[111874]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1534
		// _ = "end of CoverTab[111874]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1535
		_go_fuzz_dep_.CoverTab[111875]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1535
		// _ = "end of CoverTab[111875]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1535
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1535
	// _ = "end of CoverTab[111871]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1535
	_go_fuzz_dep_.CoverTab[111872]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1536
		_go_fuzz_dep_.CoverTab[111876]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1537
		// _ = "end of CoverTab[111876]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1538
		_go_fuzz_dep_.CoverTab[111877]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1538
		// _ = "end of CoverTab[111877]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1538
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1538
	// _ = "end of CoverTab[111872]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1538
	_go_fuzz_dep_.CoverTab[111873]++
													v := math.Float64frombits(uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56)
													*f.toFloat64() = v
													return b[8:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1541
	// _ = "end of CoverTab[111873]"
}

func unmarshalFloat64Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1544
	_go_fuzz_dep_.CoverTab[111878]++
													if w != WireFixed64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1545
		_go_fuzz_dep_.CoverTab[111881]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1546
		// _ = "end of CoverTab[111881]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1547
		_go_fuzz_dep_.CoverTab[111882]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1547
		// _ = "end of CoverTab[111882]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1547
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1547
	// _ = "end of CoverTab[111878]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1547
	_go_fuzz_dep_.CoverTab[111879]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1548
		_go_fuzz_dep_.CoverTab[111883]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1549
		// _ = "end of CoverTab[111883]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1550
		_go_fuzz_dep_.CoverTab[111884]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1550
		// _ = "end of CoverTab[111884]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1550
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1550
	// _ = "end of CoverTab[111879]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1550
	_go_fuzz_dep_.CoverTab[111880]++
													v := math.Float64frombits(uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56)
													*f.toFloat64Ptr() = &v
													return b[8:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1553
	// _ = "end of CoverTab[111880]"
}

func unmarshalFloat64Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1556
	_go_fuzz_dep_.CoverTab[111885]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1557
		_go_fuzz_dep_.CoverTab[111889]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1559
			_go_fuzz_dep_.CoverTab[111893]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1560
			// _ = "end of CoverTab[111893]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1561
			_go_fuzz_dep_.CoverTab[111894]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1561
			// _ = "end of CoverTab[111894]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1561
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1561
		// _ = "end of CoverTab[111889]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1561
		_go_fuzz_dep_.CoverTab[111890]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1563
			_go_fuzz_dep_.CoverTab[111895]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1564
			// _ = "end of CoverTab[111895]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1565
			_go_fuzz_dep_.CoverTab[111896]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1565
			// _ = "end of CoverTab[111896]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1565
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1565
		// _ = "end of CoverTab[111890]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1565
		_go_fuzz_dep_.CoverTab[111891]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1568
			_go_fuzz_dep_.CoverTab[111897]++
															if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1569
				_go_fuzz_dep_.CoverTab[111899]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1570
				// _ = "end of CoverTab[111899]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1571
				_go_fuzz_dep_.CoverTab[111900]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1571
				// _ = "end of CoverTab[111900]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1571
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1571
			// _ = "end of CoverTab[111897]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1571
			_go_fuzz_dep_.CoverTab[111898]++
															v := math.Float64frombits(uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56)
															s := f.toFloat64Slice()
															*s = append(*s, v)
															b = b[8:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1575
			// _ = "end of CoverTab[111898]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1576
		// _ = "end of CoverTab[111891]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1576
		_go_fuzz_dep_.CoverTab[111892]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1577
		// _ = "end of CoverTab[111892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1578
		_go_fuzz_dep_.CoverTab[111901]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1578
		// _ = "end of CoverTab[111901]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1578
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1578
	// _ = "end of CoverTab[111885]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1578
	_go_fuzz_dep_.CoverTab[111886]++
													if w != WireFixed64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1579
		_go_fuzz_dep_.CoverTab[111902]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1580
		// _ = "end of CoverTab[111902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1581
		_go_fuzz_dep_.CoverTab[111903]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1581
		// _ = "end of CoverTab[111903]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1581
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1581
	// _ = "end of CoverTab[111886]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1581
	_go_fuzz_dep_.CoverTab[111887]++
													if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1582
		_go_fuzz_dep_.CoverTab[111904]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1583
		// _ = "end of CoverTab[111904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1584
		_go_fuzz_dep_.CoverTab[111905]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1584
		// _ = "end of CoverTab[111905]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1584
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1584
	// _ = "end of CoverTab[111887]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1584
	_go_fuzz_dep_.CoverTab[111888]++
													v := math.Float64frombits(uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56)
													s := f.toFloat64Slice()
													*s = append(*s, v)
													return b[8:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1588
	// _ = "end of CoverTab[111888]"
}

func unmarshalFloat32Value(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1591
	_go_fuzz_dep_.CoverTab[111906]++
													if w != WireFixed32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1592
		_go_fuzz_dep_.CoverTab[111909]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1593
		// _ = "end of CoverTab[111909]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1594
		_go_fuzz_dep_.CoverTab[111910]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1594
		// _ = "end of CoverTab[111910]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1594
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1594
	// _ = "end of CoverTab[111906]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1594
	_go_fuzz_dep_.CoverTab[111907]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1595
		_go_fuzz_dep_.CoverTab[111911]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1596
		// _ = "end of CoverTab[111911]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1597
		_go_fuzz_dep_.CoverTab[111912]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1597
		// _ = "end of CoverTab[111912]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1597
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1597
	// _ = "end of CoverTab[111907]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1597
	_go_fuzz_dep_.CoverTab[111908]++
													v := math.Float32frombits(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
													*f.toFloat32() = v
													return b[4:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1600
	// _ = "end of CoverTab[111908]"
}

func unmarshalFloat32Ptr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1603
	_go_fuzz_dep_.CoverTab[111913]++
													if w != WireFixed32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1604
		_go_fuzz_dep_.CoverTab[111916]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1605
		// _ = "end of CoverTab[111916]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1606
		_go_fuzz_dep_.CoverTab[111917]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1606
		// _ = "end of CoverTab[111917]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1606
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1606
	// _ = "end of CoverTab[111913]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1606
	_go_fuzz_dep_.CoverTab[111914]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1607
		_go_fuzz_dep_.CoverTab[111918]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1608
		// _ = "end of CoverTab[111918]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1609
		_go_fuzz_dep_.CoverTab[111919]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1609
		// _ = "end of CoverTab[111919]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1609
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1609
	// _ = "end of CoverTab[111914]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1609
	_go_fuzz_dep_.CoverTab[111915]++
													v := math.Float32frombits(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
													*f.toFloat32Ptr() = &v
													return b[4:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1612
	// _ = "end of CoverTab[111915]"
}

func unmarshalFloat32Slice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1615
	_go_fuzz_dep_.CoverTab[111920]++
													if w == WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1616
		_go_fuzz_dep_.CoverTab[111924]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1618
			_go_fuzz_dep_.CoverTab[111928]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1619
			// _ = "end of CoverTab[111928]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1620
			_go_fuzz_dep_.CoverTab[111929]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1620
			// _ = "end of CoverTab[111929]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1620
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1620
		// _ = "end of CoverTab[111924]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1620
		_go_fuzz_dep_.CoverTab[111925]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1622
			_go_fuzz_dep_.CoverTab[111930]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1623
			// _ = "end of CoverTab[111930]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1624
			_go_fuzz_dep_.CoverTab[111931]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1624
			// _ = "end of CoverTab[111931]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1624
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1624
		// _ = "end of CoverTab[111925]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1624
		_go_fuzz_dep_.CoverTab[111926]++
														res := b[x:]
														b = b[:x]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1627
			_go_fuzz_dep_.CoverTab[111932]++
															if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1628
				_go_fuzz_dep_.CoverTab[111934]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1629
				// _ = "end of CoverTab[111934]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1630
				_go_fuzz_dep_.CoverTab[111935]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1630
				// _ = "end of CoverTab[111935]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1630
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1630
			// _ = "end of CoverTab[111932]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1630
			_go_fuzz_dep_.CoverTab[111933]++
															v := math.Float32frombits(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
															s := f.toFloat32Slice()
															*s = append(*s, v)
															b = b[4:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1634
			// _ = "end of CoverTab[111933]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1635
		// _ = "end of CoverTab[111926]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1635
		_go_fuzz_dep_.CoverTab[111927]++
														return res, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1636
		// _ = "end of CoverTab[111927]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1637
		_go_fuzz_dep_.CoverTab[111936]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1637
		// _ = "end of CoverTab[111936]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1637
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1637
	// _ = "end of CoverTab[111920]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1637
	_go_fuzz_dep_.CoverTab[111921]++
													if w != WireFixed32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1638
		_go_fuzz_dep_.CoverTab[111937]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1639
		// _ = "end of CoverTab[111937]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1640
		_go_fuzz_dep_.CoverTab[111938]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1640
		// _ = "end of CoverTab[111938]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1640
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1640
	// _ = "end of CoverTab[111921]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1640
	_go_fuzz_dep_.CoverTab[111922]++
													if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1641
		_go_fuzz_dep_.CoverTab[111939]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1642
		// _ = "end of CoverTab[111939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1643
		_go_fuzz_dep_.CoverTab[111940]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1643
		// _ = "end of CoverTab[111940]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1643
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1643
	// _ = "end of CoverTab[111922]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1643
	_go_fuzz_dep_.CoverTab[111923]++
													v := math.Float32frombits(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
													s := f.toFloat32Slice()
													*s = append(*s, v)
													return b[4:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1647
	// _ = "end of CoverTab[111923]"
}

func unmarshalStringValue(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1650
	_go_fuzz_dep_.CoverTab[111941]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1651
		_go_fuzz_dep_.CoverTab[111945]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1652
		// _ = "end of CoverTab[111945]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1653
		_go_fuzz_dep_.CoverTab[111946]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1653
		// _ = "end of CoverTab[111946]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1653
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1653
	// _ = "end of CoverTab[111941]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1653
	_go_fuzz_dep_.CoverTab[111942]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1655
		_go_fuzz_dep_.CoverTab[111947]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1656
		// _ = "end of CoverTab[111947]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1657
		_go_fuzz_dep_.CoverTab[111948]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1657
		// _ = "end of CoverTab[111948]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1657
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1657
	// _ = "end of CoverTab[111942]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1657
	_go_fuzz_dep_.CoverTab[111943]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1659
		_go_fuzz_dep_.CoverTab[111949]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1660
		// _ = "end of CoverTab[111949]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1661
		_go_fuzz_dep_.CoverTab[111950]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1661
		// _ = "end of CoverTab[111950]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1661
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1661
	// _ = "end of CoverTab[111943]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1661
	_go_fuzz_dep_.CoverTab[111944]++
													v := string(b[:x])
													*f.toString() = v
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1664
	// _ = "end of CoverTab[111944]"
}

func unmarshalStringPtr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1667
	_go_fuzz_dep_.CoverTab[111951]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1668
		_go_fuzz_dep_.CoverTab[111955]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1669
		// _ = "end of CoverTab[111955]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1670
		_go_fuzz_dep_.CoverTab[111956]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1670
		// _ = "end of CoverTab[111956]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1670
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1670
	// _ = "end of CoverTab[111951]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1670
	_go_fuzz_dep_.CoverTab[111952]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1672
		_go_fuzz_dep_.CoverTab[111957]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1673
		// _ = "end of CoverTab[111957]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1674
		_go_fuzz_dep_.CoverTab[111958]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1674
		// _ = "end of CoverTab[111958]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1674
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1674
	// _ = "end of CoverTab[111952]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1674
	_go_fuzz_dep_.CoverTab[111953]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1676
		_go_fuzz_dep_.CoverTab[111959]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1677
		// _ = "end of CoverTab[111959]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1678
		_go_fuzz_dep_.CoverTab[111960]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1678
		// _ = "end of CoverTab[111960]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1678
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1678
	// _ = "end of CoverTab[111953]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1678
	_go_fuzz_dep_.CoverTab[111954]++
													v := string(b[:x])
													*f.toStringPtr() = &v
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1681
	// _ = "end of CoverTab[111954]"
}

func unmarshalStringSlice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1684
	_go_fuzz_dep_.CoverTab[111961]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1685
		_go_fuzz_dep_.CoverTab[111965]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1686
		// _ = "end of CoverTab[111965]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1687
		_go_fuzz_dep_.CoverTab[111966]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1687
		// _ = "end of CoverTab[111966]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1687
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1687
	// _ = "end of CoverTab[111961]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1687
	_go_fuzz_dep_.CoverTab[111962]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1689
		_go_fuzz_dep_.CoverTab[111967]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1690
		// _ = "end of CoverTab[111967]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1691
		_go_fuzz_dep_.CoverTab[111968]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1691
		// _ = "end of CoverTab[111968]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1691
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1691
	// _ = "end of CoverTab[111962]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1691
	_go_fuzz_dep_.CoverTab[111963]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1693
		_go_fuzz_dep_.CoverTab[111969]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1694
		// _ = "end of CoverTab[111969]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1695
		_go_fuzz_dep_.CoverTab[111970]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1695
		// _ = "end of CoverTab[111970]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1695
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1695
	// _ = "end of CoverTab[111963]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1695
	_go_fuzz_dep_.CoverTab[111964]++
													v := string(b[:x])
													s := f.toStringSlice()
													*s = append(*s, v)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1699
	// _ = "end of CoverTab[111964]"
}

func unmarshalUTF8StringValue(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1702
	_go_fuzz_dep_.CoverTab[111971]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1703
		_go_fuzz_dep_.CoverTab[111976]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1704
		// _ = "end of CoverTab[111976]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1705
		_go_fuzz_dep_.CoverTab[111977]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1705
		// _ = "end of CoverTab[111977]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1705
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1705
	// _ = "end of CoverTab[111971]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1705
	_go_fuzz_dep_.CoverTab[111972]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1707
		_go_fuzz_dep_.CoverTab[111978]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1708
		// _ = "end of CoverTab[111978]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1709
		_go_fuzz_dep_.CoverTab[111979]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1709
		// _ = "end of CoverTab[111979]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1709
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1709
	// _ = "end of CoverTab[111972]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1709
	_go_fuzz_dep_.CoverTab[111973]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1711
		_go_fuzz_dep_.CoverTab[111980]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1712
		// _ = "end of CoverTab[111980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1713
		_go_fuzz_dep_.CoverTab[111981]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1713
		// _ = "end of CoverTab[111981]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1713
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1713
	// _ = "end of CoverTab[111973]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1713
	_go_fuzz_dep_.CoverTab[111974]++
													v := string(b[:x])
													*f.toString() = v
													if !utf8.ValidString(v) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1716
		_go_fuzz_dep_.CoverTab[111982]++
														return b[x:], errInvalidUTF8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1717
		// _ = "end of CoverTab[111982]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1718
		_go_fuzz_dep_.CoverTab[111983]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1718
		// _ = "end of CoverTab[111983]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1718
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1718
	// _ = "end of CoverTab[111974]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1718
	_go_fuzz_dep_.CoverTab[111975]++
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1719
	// _ = "end of CoverTab[111975]"
}

func unmarshalUTF8StringPtr(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1722
	_go_fuzz_dep_.CoverTab[111984]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1723
		_go_fuzz_dep_.CoverTab[111989]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1724
		// _ = "end of CoverTab[111989]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1725
		_go_fuzz_dep_.CoverTab[111990]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1725
		// _ = "end of CoverTab[111990]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1725
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1725
	// _ = "end of CoverTab[111984]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1725
	_go_fuzz_dep_.CoverTab[111985]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1727
		_go_fuzz_dep_.CoverTab[111991]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1728
		// _ = "end of CoverTab[111991]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1729
		_go_fuzz_dep_.CoverTab[111992]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1729
		// _ = "end of CoverTab[111992]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1729
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1729
	// _ = "end of CoverTab[111985]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1729
	_go_fuzz_dep_.CoverTab[111986]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1731
		_go_fuzz_dep_.CoverTab[111993]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1732
		// _ = "end of CoverTab[111993]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1733
		_go_fuzz_dep_.CoverTab[111994]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1733
		// _ = "end of CoverTab[111994]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1733
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1733
	// _ = "end of CoverTab[111986]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1733
	_go_fuzz_dep_.CoverTab[111987]++
													v := string(b[:x])
													*f.toStringPtr() = &v
													if !utf8.ValidString(v) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1736
		_go_fuzz_dep_.CoverTab[111995]++
														return b[x:], errInvalidUTF8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1737
		// _ = "end of CoverTab[111995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1738
		_go_fuzz_dep_.CoverTab[111996]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1738
		// _ = "end of CoverTab[111996]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1738
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1738
	// _ = "end of CoverTab[111987]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1738
	_go_fuzz_dep_.CoverTab[111988]++
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1739
	// _ = "end of CoverTab[111988]"
}

func unmarshalUTF8StringSlice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1742
	_go_fuzz_dep_.CoverTab[111997]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1743
		_go_fuzz_dep_.CoverTab[112002]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1744
		// _ = "end of CoverTab[112002]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1745
		_go_fuzz_dep_.CoverTab[112003]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1745
		// _ = "end of CoverTab[112003]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1745
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1745
	// _ = "end of CoverTab[111997]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1745
	_go_fuzz_dep_.CoverTab[111998]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1747
		_go_fuzz_dep_.CoverTab[112004]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1748
		// _ = "end of CoverTab[112004]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1749
		_go_fuzz_dep_.CoverTab[112005]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1749
		// _ = "end of CoverTab[112005]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1749
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1749
	// _ = "end of CoverTab[111998]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1749
	_go_fuzz_dep_.CoverTab[111999]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1751
		_go_fuzz_dep_.CoverTab[112006]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1752
		// _ = "end of CoverTab[112006]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1753
		_go_fuzz_dep_.CoverTab[112007]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1753
		// _ = "end of CoverTab[112007]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1753
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1753
	// _ = "end of CoverTab[111999]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1753
	_go_fuzz_dep_.CoverTab[112000]++
													v := string(b[:x])
													s := f.toStringSlice()
													*s = append(*s, v)
													if !utf8.ValidString(v) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1757
		_go_fuzz_dep_.CoverTab[112008]++
														return b[x:], errInvalidUTF8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1758
		// _ = "end of CoverTab[112008]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1759
		_go_fuzz_dep_.CoverTab[112009]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1759
		// _ = "end of CoverTab[112009]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1759
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1759
	// _ = "end of CoverTab[112000]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1759
	_go_fuzz_dep_.CoverTab[112001]++
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1760
	// _ = "end of CoverTab[112001]"
}

var emptyBuf [0]byte

func unmarshalBytesValue(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1765
	_go_fuzz_dep_.CoverTab[112010]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1766
		_go_fuzz_dep_.CoverTab[112014]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1767
		// _ = "end of CoverTab[112014]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1768
		_go_fuzz_dep_.CoverTab[112015]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1768
		// _ = "end of CoverTab[112015]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1768
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1768
	// _ = "end of CoverTab[112010]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1768
	_go_fuzz_dep_.CoverTab[112011]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1770
		_go_fuzz_dep_.CoverTab[112016]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1771
		// _ = "end of CoverTab[112016]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1772
		_go_fuzz_dep_.CoverTab[112017]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1772
		// _ = "end of CoverTab[112017]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1772
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1772
	// _ = "end of CoverTab[112011]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1772
	_go_fuzz_dep_.CoverTab[112012]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1774
		_go_fuzz_dep_.CoverTab[112018]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1775
		// _ = "end of CoverTab[112018]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1776
		_go_fuzz_dep_.CoverTab[112019]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1776
		// _ = "end of CoverTab[112019]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1776
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1776
	// _ = "end of CoverTab[112012]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1776
	_go_fuzz_dep_.CoverTab[112013]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1781
	v := append(emptyBuf[:], b[:x]...)
													*f.toBytes() = v
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1783
	// _ = "end of CoverTab[112013]"
}

func unmarshalBytesSlice(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1786
	_go_fuzz_dep_.CoverTab[112020]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1787
		_go_fuzz_dep_.CoverTab[112024]++
														return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1788
		// _ = "end of CoverTab[112024]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1789
		_go_fuzz_dep_.CoverTab[112025]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1789
		// _ = "end of CoverTab[112025]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1789
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1789
	// _ = "end of CoverTab[112020]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1789
	_go_fuzz_dep_.CoverTab[112021]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1791
		_go_fuzz_dep_.CoverTab[112026]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1792
		// _ = "end of CoverTab[112026]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1793
		_go_fuzz_dep_.CoverTab[112027]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1793
		// _ = "end of CoverTab[112027]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1793
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1793
	// _ = "end of CoverTab[112021]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1793
	_go_fuzz_dep_.CoverTab[112022]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1795
		_go_fuzz_dep_.CoverTab[112028]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1796
		// _ = "end of CoverTab[112028]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1797
		_go_fuzz_dep_.CoverTab[112029]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1797
		// _ = "end of CoverTab[112029]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1797
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1797
	// _ = "end of CoverTab[112022]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1797
	_go_fuzz_dep_.CoverTab[112023]++
													v := append(emptyBuf[:], b[:x]...)
													s := f.toBytesSlice()
													*s = append(*s, v)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1801
	// _ = "end of CoverTab[112023]"
}

func makeUnmarshalMessagePtr(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1804
	_go_fuzz_dep_.CoverTab[112030]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1805
		_go_fuzz_dep_.CoverTab[112031]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1806
			_go_fuzz_dep_.CoverTab[112037]++
															return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1807
			// _ = "end of CoverTab[112037]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1808
			_go_fuzz_dep_.CoverTab[112038]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1808
			// _ = "end of CoverTab[112038]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1808
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1808
		// _ = "end of CoverTab[112031]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1808
		_go_fuzz_dep_.CoverTab[112032]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1810
			_go_fuzz_dep_.CoverTab[112039]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1811
			// _ = "end of CoverTab[112039]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1812
			_go_fuzz_dep_.CoverTab[112040]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1812
			// _ = "end of CoverTab[112040]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1812
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1812
		// _ = "end of CoverTab[112032]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1812
		_go_fuzz_dep_.CoverTab[112033]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1814
			_go_fuzz_dep_.CoverTab[112041]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1815
			// _ = "end of CoverTab[112041]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1816
			_go_fuzz_dep_.CoverTab[112042]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1816
			// _ = "end of CoverTab[112042]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1816
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1816
		// _ = "end of CoverTab[112033]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1816
		_go_fuzz_dep_.CoverTab[112034]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1821
		v := f.getPointer()
		if v.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1822
			_go_fuzz_dep_.CoverTab[112043]++
															v = valToPointer(reflect.New(sub.typ))
															f.setPointer(v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1824
			// _ = "end of CoverTab[112043]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1825
			_go_fuzz_dep_.CoverTab[112044]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1825
			// _ = "end of CoverTab[112044]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1825
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1825
		// _ = "end of CoverTab[112034]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1825
		_go_fuzz_dep_.CoverTab[112035]++
														err := sub.unmarshal(v, b[:x])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1827
			_go_fuzz_dep_.CoverTab[112045]++
															if r, ok := err.(*RequiredNotSetError); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1828
				_go_fuzz_dep_.CoverTab[112046]++
																r.field = name + "." + r.field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1829
				// _ = "end of CoverTab[112046]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1830
				_go_fuzz_dep_.CoverTab[112047]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1831
				// _ = "end of CoverTab[112047]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1832
			// _ = "end of CoverTab[112045]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1833
			_go_fuzz_dep_.CoverTab[112048]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1833
			// _ = "end of CoverTab[112048]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1833
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1833
		// _ = "end of CoverTab[112035]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1833
		_go_fuzz_dep_.CoverTab[112036]++
														return b[x:], err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1834
		// _ = "end of CoverTab[112036]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1835
	// _ = "end of CoverTab[112030]"
}

func makeUnmarshalMessageSlicePtr(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1838
	_go_fuzz_dep_.CoverTab[112049]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1839
		_go_fuzz_dep_.CoverTab[112050]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1840
			_go_fuzz_dep_.CoverTab[112055]++
															return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1841
			// _ = "end of CoverTab[112055]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1842
			_go_fuzz_dep_.CoverTab[112056]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1842
			// _ = "end of CoverTab[112056]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1842
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1842
		// _ = "end of CoverTab[112050]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1842
		_go_fuzz_dep_.CoverTab[112051]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1844
			_go_fuzz_dep_.CoverTab[112057]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1845
			// _ = "end of CoverTab[112057]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1846
			_go_fuzz_dep_.CoverTab[112058]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1846
			// _ = "end of CoverTab[112058]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1846
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1846
		// _ = "end of CoverTab[112051]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1846
		_go_fuzz_dep_.CoverTab[112052]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1848
			_go_fuzz_dep_.CoverTab[112059]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1849
			// _ = "end of CoverTab[112059]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1850
			_go_fuzz_dep_.CoverTab[112060]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1850
			// _ = "end of CoverTab[112060]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1850
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1850
		// _ = "end of CoverTab[112052]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1850
		_go_fuzz_dep_.CoverTab[112053]++
														v := valToPointer(reflect.New(sub.typ))
														err := sub.unmarshal(v, b[:x])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1853
			_go_fuzz_dep_.CoverTab[112061]++
															if r, ok := err.(*RequiredNotSetError); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1854
				_go_fuzz_dep_.CoverTab[112062]++
																r.field = name + "." + r.field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1855
				// _ = "end of CoverTab[112062]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1856
				_go_fuzz_dep_.CoverTab[112063]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1857
				// _ = "end of CoverTab[112063]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1858
			// _ = "end of CoverTab[112061]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1859
			_go_fuzz_dep_.CoverTab[112064]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1859
			// _ = "end of CoverTab[112064]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1859
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1859
		// _ = "end of CoverTab[112053]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1859
		_go_fuzz_dep_.CoverTab[112054]++
														f.appendPointer(v)
														return b[x:], err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1861
		// _ = "end of CoverTab[112054]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1862
	// _ = "end of CoverTab[112049]"
}

func makeUnmarshalGroupPtr(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1865
	_go_fuzz_dep_.CoverTab[112065]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1866
		_go_fuzz_dep_.CoverTab[112066]++
														if w != WireStartGroup {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1867
			_go_fuzz_dep_.CoverTab[112071]++
															return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1868
			// _ = "end of CoverTab[112071]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1869
			_go_fuzz_dep_.CoverTab[112072]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1869
			// _ = "end of CoverTab[112072]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1869
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1869
		// _ = "end of CoverTab[112066]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1869
		_go_fuzz_dep_.CoverTab[112067]++
														x, y := findEndGroup(b)
														if x < 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1871
			_go_fuzz_dep_.CoverTab[112073]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1872
			// _ = "end of CoverTab[112073]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1873
			_go_fuzz_dep_.CoverTab[112074]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1873
			// _ = "end of CoverTab[112074]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1873
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1873
		// _ = "end of CoverTab[112067]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1873
		_go_fuzz_dep_.CoverTab[112068]++
														v := f.getPointer()
														if v.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1875
			_go_fuzz_dep_.CoverTab[112075]++
															v = valToPointer(reflect.New(sub.typ))
															f.setPointer(v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1877
			// _ = "end of CoverTab[112075]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1878
			_go_fuzz_dep_.CoverTab[112076]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1878
			// _ = "end of CoverTab[112076]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1878
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1878
		// _ = "end of CoverTab[112068]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1878
		_go_fuzz_dep_.CoverTab[112069]++
														err := sub.unmarshal(v, b[:x])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1880
			_go_fuzz_dep_.CoverTab[112077]++
															if r, ok := err.(*RequiredNotSetError); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1881
				_go_fuzz_dep_.CoverTab[112078]++
																r.field = name + "." + r.field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1882
				// _ = "end of CoverTab[112078]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1883
				_go_fuzz_dep_.CoverTab[112079]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1884
				// _ = "end of CoverTab[112079]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1885
			// _ = "end of CoverTab[112077]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1886
			_go_fuzz_dep_.CoverTab[112080]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1886
			// _ = "end of CoverTab[112080]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1886
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1886
		// _ = "end of CoverTab[112069]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1886
		_go_fuzz_dep_.CoverTab[112070]++
														return b[y:], err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1887
		// _ = "end of CoverTab[112070]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1888
	// _ = "end of CoverTab[112065]"
}

func makeUnmarshalGroupSlicePtr(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1891
	_go_fuzz_dep_.CoverTab[112081]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1892
		_go_fuzz_dep_.CoverTab[112082]++
														if w != WireStartGroup {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1893
			_go_fuzz_dep_.CoverTab[112086]++
															return b, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1894
			// _ = "end of CoverTab[112086]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1895
			_go_fuzz_dep_.CoverTab[112087]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1895
			// _ = "end of CoverTab[112087]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1895
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1895
		// _ = "end of CoverTab[112082]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1895
		_go_fuzz_dep_.CoverTab[112083]++
														x, y := findEndGroup(b)
														if x < 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1897
			_go_fuzz_dep_.CoverTab[112088]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1898
			// _ = "end of CoverTab[112088]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1899
			_go_fuzz_dep_.CoverTab[112089]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1899
			// _ = "end of CoverTab[112089]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1899
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1899
		// _ = "end of CoverTab[112083]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1899
		_go_fuzz_dep_.CoverTab[112084]++
														v := valToPointer(reflect.New(sub.typ))
														err := sub.unmarshal(v, b[:x])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1902
			_go_fuzz_dep_.CoverTab[112090]++
															if r, ok := err.(*RequiredNotSetError); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1903
				_go_fuzz_dep_.CoverTab[112091]++
																r.field = name + "." + r.field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1904
				// _ = "end of CoverTab[112091]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1905
				_go_fuzz_dep_.CoverTab[112092]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1906
				// _ = "end of CoverTab[112092]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1907
			// _ = "end of CoverTab[112090]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1908
			_go_fuzz_dep_.CoverTab[112093]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1908
			// _ = "end of CoverTab[112093]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1908
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1908
		// _ = "end of CoverTab[112084]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1908
		_go_fuzz_dep_.CoverTab[112085]++
														f.appendPointer(v)
														return b[y:], err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1910
		// _ = "end of CoverTab[112085]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1911
	// _ = "end of CoverTab[112081]"
}

func makeUnmarshalMap(f *reflect.StructField) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1914
	_go_fuzz_dep_.CoverTab[112094]++
													t := f.Type
													kt := t.Key()
													vt := t.Elem()
													tagArray := strings.Split(f.Tag.Get("protobuf"), ",")
													valTags := strings.Split(f.Tag.Get("protobuf_val"), ",")
													for _, t := range tagArray {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1920
		_go_fuzz_dep_.CoverTab[112096]++
														if strings.HasPrefix(t, "customtype=") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1921
			_go_fuzz_dep_.CoverTab[112100]++
															valTags = append(valTags, t)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1922
			// _ = "end of CoverTab[112100]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1923
			_go_fuzz_dep_.CoverTab[112101]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1923
			// _ = "end of CoverTab[112101]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1923
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1923
		// _ = "end of CoverTab[112096]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1923
		_go_fuzz_dep_.CoverTab[112097]++
														if t == "stdtime" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1924
			_go_fuzz_dep_.CoverTab[112102]++
															valTags = append(valTags, t)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1925
			// _ = "end of CoverTab[112102]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1926
			_go_fuzz_dep_.CoverTab[112103]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1926
			// _ = "end of CoverTab[112103]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1926
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1926
		// _ = "end of CoverTab[112097]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1926
		_go_fuzz_dep_.CoverTab[112098]++
														if t == "stdduration" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1927
			_go_fuzz_dep_.CoverTab[112104]++
															valTags = append(valTags, t)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1928
			// _ = "end of CoverTab[112104]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1929
			_go_fuzz_dep_.CoverTab[112105]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1929
			// _ = "end of CoverTab[112105]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1929
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1929
		// _ = "end of CoverTab[112098]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1929
		_go_fuzz_dep_.CoverTab[112099]++
														if t == "wktptr" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1930
			_go_fuzz_dep_.CoverTab[112106]++
															valTags = append(valTags, t)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1931
			// _ = "end of CoverTab[112106]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1932
			_go_fuzz_dep_.CoverTab[112107]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1932
			// _ = "end of CoverTab[112107]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1932
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1932
		// _ = "end of CoverTab[112099]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1933
	// _ = "end of CoverTab[112094]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1933
	_go_fuzz_dep_.CoverTab[112095]++
													unmarshalKey := typeUnmarshaler(kt, f.Tag.Get("protobuf_key"))
													unmarshalVal := typeUnmarshaler(vt, strings.Join(valTags, ","))
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1936
		_go_fuzz_dep_.CoverTab[112108]++

														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1938
			_go_fuzz_dep_.CoverTab[112114]++
															return nil, fmt.Errorf("proto: bad wiretype for map field: got %d want %d", w, WireBytes)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1939
			// _ = "end of CoverTab[112114]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1940
			_go_fuzz_dep_.CoverTab[112115]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1940
			// _ = "end of CoverTab[112115]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1940
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1940
		// _ = "end of CoverTab[112108]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1940
		_go_fuzz_dep_.CoverTab[112109]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1942
			_go_fuzz_dep_.CoverTab[112116]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1943
			// _ = "end of CoverTab[112116]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1944
			_go_fuzz_dep_.CoverTab[112117]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1944
			// _ = "end of CoverTab[112117]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1944
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1944
		// _ = "end of CoverTab[112109]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1944
		_go_fuzz_dep_.CoverTab[112110]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1946
			_go_fuzz_dep_.CoverTab[112118]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1947
			// _ = "end of CoverTab[112118]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1948
			_go_fuzz_dep_.CoverTab[112119]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1948
			// _ = "end of CoverTab[112119]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1948
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1948
		// _ = "end of CoverTab[112110]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1948
		_go_fuzz_dep_.CoverTab[112111]++
														r := b[x:]
														b = b[:x]

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1956
		// Read key and value from data.
														var nerr nonFatal
														k := reflect.New(kt)
														v := reflect.New(vt)
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1960
			_go_fuzz_dep_.CoverTab[112120]++
															x, n := decodeVarint(b)
															if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1962
				_go_fuzz_dep_.CoverTab[112125]++
																return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1963
				// _ = "end of CoverTab[112125]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1964
				_go_fuzz_dep_.CoverTab[112126]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1964
				// _ = "end of CoverTab[112126]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1964
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1964
			// _ = "end of CoverTab[112120]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1964
			_go_fuzz_dep_.CoverTab[112121]++
															wire := int(x) & 7
															b = b[n:]

															var err error
															switch x >> 3 {
			case 1:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1970
				_go_fuzz_dep_.CoverTab[112127]++
																b, err = unmarshalKey(b, valToPointer(k), wire)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1971
				// _ = "end of CoverTab[112127]"
			case 2:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1972
				_go_fuzz_dep_.CoverTab[112128]++
																b, err = unmarshalVal(b, valToPointer(v), wire)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1973
				// _ = "end of CoverTab[112128]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1974
				_go_fuzz_dep_.CoverTab[112129]++
																err = errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1975
				// _ = "end of CoverTab[112129]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1976
			// _ = "end of CoverTab[112121]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1976
			_go_fuzz_dep_.CoverTab[112122]++

															if nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1978
				_go_fuzz_dep_.CoverTab[112130]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1979
				// _ = "end of CoverTab[112130]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1980
				_go_fuzz_dep_.CoverTab[112131]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1980
				// _ = "end of CoverTab[112131]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1980
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1980
			// _ = "end of CoverTab[112122]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1980
			_go_fuzz_dep_.CoverTab[112123]++
															if err != errInternalBadWireType {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1981
				_go_fuzz_dep_.CoverTab[112132]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1982
				// _ = "end of CoverTab[112132]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1983
				_go_fuzz_dep_.CoverTab[112133]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1983
				// _ = "end of CoverTab[112133]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1983
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1983
			// _ = "end of CoverTab[112123]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1983
			_go_fuzz_dep_.CoverTab[112124]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1986
			b, err = skipField(b, wire)
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1987
				_go_fuzz_dep_.CoverTab[112134]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1988
				// _ = "end of CoverTab[112134]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1989
				_go_fuzz_dep_.CoverTab[112135]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1989
				// _ = "end of CoverTab[112135]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1989
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1989
			// _ = "end of CoverTab[112124]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1990
		// _ = "end of CoverTab[112111]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1990
		_go_fuzz_dep_.CoverTab[112112]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1993
		m := f.asPointerTo(t).Elem()
		if m.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1994
			_go_fuzz_dep_.CoverTab[112136]++
															m.Set(reflect.MakeMap(t))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1995
			// _ = "end of CoverTab[112136]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1996
			_go_fuzz_dep_.CoverTab[112137]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1996
			// _ = "end of CoverTab[112137]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1996
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1996
		// _ = "end of CoverTab[112112]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1996
		_go_fuzz_dep_.CoverTab[112113]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:1999
		m.SetMapIndex(k.Elem(), v.Elem())

														return r, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2001
		// _ = "end of CoverTab[112113]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2002
	// _ = "end of CoverTab[112095]"
}

// makeUnmarshalOneof makes an unmarshaler for oneof fields.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
// for:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
//	message Msg {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
//	  oneof F {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
//	    int64 X = 1;
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
//	    float64 Y = 2;
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
//	  }
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
//	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
// typ is the type of the concrete entry for a oneof case (e.g. Msg_X).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
// ityp is the interface type of the oneof field (e.g. isMsg_F).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
// unmarshal is the unmarshaler for the base type of the oneof case (e.g. int64).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2005
// Note that this function will be called once for each case in the oneof.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2017
func makeUnmarshalOneof(typ, ityp reflect.Type, unmarshal unmarshaler) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2017
	_go_fuzz_dep_.CoverTab[112138]++
													sf := typ.Field(0)
													field0 := toField(&sf)
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2020
		_go_fuzz_dep_.CoverTab[112139]++

														v := reflect.New(typ)

		// Unmarshal data into holder.
		// We unmarshal into the first field of the holder object.
		var err error
		var nerr nonFatal
		b, err = unmarshal(b, valToPointer(v).offset(field0), w)
		if !nerr.Merge(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2029
			_go_fuzz_dep_.CoverTab[112141]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2030
			// _ = "end of CoverTab[112141]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2031
			_go_fuzz_dep_.CoverTab[112142]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2031
			// _ = "end of CoverTab[112142]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2031
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2031
		// _ = "end of CoverTab[112139]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2031
		_go_fuzz_dep_.CoverTab[112140]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2034
		f.asPointerTo(ityp).Elem().Set(v)

														return b, nerr.E
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2036
		// _ = "end of CoverTab[112140]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2037
	// _ = "end of CoverTab[112138]"
}

// Error used by decode internally.
var errInternalBadWireType = errors.New("proto: internal error: bad wiretype")

// skipField skips past a field of type wire and returns the remaining bytes.
func skipField(b []byte, wire int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2044
	_go_fuzz_dep_.CoverTab[112143]++
													switch wire {
	case WireVarint:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2046
		_go_fuzz_dep_.CoverTab[112145]++
														_, k := decodeVarint(b)
														if k == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2048
			_go_fuzz_dep_.CoverTab[112156]++
															return b, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2049
			// _ = "end of CoverTab[112156]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2050
			_go_fuzz_dep_.CoverTab[112157]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2050
			// _ = "end of CoverTab[112157]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2050
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2050
		// _ = "end of CoverTab[112145]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2050
		_go_fuzz_dep_.CoverTab[112146]++
														b = b[k:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2051
		// _ = "end of CoverTab[112146]"
	case WireFixed32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2052
		_go_fuzz_dep_.CoverTab[112147]++
														if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2053
			_go_fuzz_dep_.CoverTab[112158]++
															return b, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2054
			// _ = "end of CoverTab[112158]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2055
			_go_fuzz_dep_.CoverTab[112159]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2055
			// _ = "end of CoverTab[112159]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2055
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2055
		// _ = "end of CoverTab[112147]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2055
		_go_fuzz_dep_.CoverTab[112148]++
														b = b[4:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2056
		// _ = "end of CoverTab[112148]"
	case WireFixed64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2057
		_go_fuzz_dep_.CoverTab[112149]++
														if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2058
			_go_fuzz_dep_.CoverTab[112160]++
															return b, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2059
			// _ = "end of CoverTab[112160]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2060
			_go_fuzz_dep_.CoverTab[112161]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2060
			// _ = "end of CoverTab[112161]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2060
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2060
		// _ = "end of CoverTab[112149]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2060
		_go_fuzz_dep_.CoverTab[112150]++
														b = b[8:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2061
		// _ = "end of CoverTab[112150]"
	case WireBytes:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2062
		_go_fuzz_dep_.CoverTab[112151]++
														m, k := decodeVarint(b)
														if k == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2064
			_go_fuzz_dep_.CoverTab[112162]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2064
			return uint64(len(b)-k) < m
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2064
			// _ = "end of CoverTab[112162]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2064
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2064
			_go_fuzz_dep_.CoverTab[112163]++
															return b, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2065
			// _ = "end of CoverTab[112163]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2066
			_go_fuzz_dep_.CoverTab[112164]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2066
			// _ = "end of CoverTab[112164]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2066
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2066
		// _ = "end of CoverTab[112151]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2066
		_go_fuzz_dep_.CoverTab[112152]++
														b = b[uint64(k)+m:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2067
		// _ = "end of CoverTab[112152]"
	case WireStartGroup:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2068
		_go_fuzz_dep_.CoverTab[112153]++
														_, i := findEndGroup(b)
														if i == -1 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2070
			_go_fuzz_dep_.CoverTab[112165]++
															return b, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2071
			// _ = "end of CoverTab[112165]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2072
			_go_fuzz_dep_.CoverTab[112166]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2072
			// _ = "end of CoverTab[112166]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2072
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2072
		// _ = "end of CoverTab[112153]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2072
		_go_fuzz_dep_.CoverTab[112154]++
														b = b[i:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2073
		// _ = "end of CoverTab[112154]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2074
		_go_fuzz_dep_.CoverTab[112155]++
														return b, fmt.Errorf("proto: can't skip unknown wire type %d", wire)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2075
		// _ = "end of CoverTab[112155]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2076
	// _ = "end of CoverTab[112143]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2076
	_go_fuzz_dep_.CoverTab[112144]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2077
	// _ = "end of CoverTab[112144]"
}

// findEndGroup finds the index of the next EndGroup tag.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2080
// Groups may be nested, so the "next" EndGroup tag is the first
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2080
// unpaired EndGroup.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2080
// findEndGroup returns the indexes of the start and end of the EndGroup tag.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2080
// Returns (-1,-1) if it can't find one.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2085
func findEndGroup(b []byte) (int, int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2085
	_go_fuzz_dep_.CoverTab[112167]++
													depth := 1
													i := 0
													for {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2088
		_go_fuzz_dep_.CoverTab[112168]++
														x, n := decodeVarint(b[i:])
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2090
			_go_fuzz_dep_.CoverTab[112170]++
															return -1, -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2091
			// _ = "end of CoverTab[112170]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2092
			_go_fuzz_dep_.CoverTab[112171]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2092
			// _ = "end of CoverTab[112171]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2092
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2092
		// _ = "end of CoverTab[112168]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2092
		_go_fuzz_dep_.CoverTab[112169]++
														j := i
														i += n
														switch x & 7 {
		case WireVarint:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2096
			_go_fuzz_dep_.CoverTab[112172]++
															_, k := decodeVarint(b[i:])
															if k == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2098
				_go_fuzz_dep_.CoverTab[112184]++
																return -1, -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2099
				// _ = "end of CoverTab[112184]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2100
				_go_fuzz_dep_.CoverTab[112185]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2100
				// _ = "end of CoverTab[112185]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2100
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2100
			// _ = "end of CoverTab[112172]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2100
			_go_fuzz_dep_.CoverTab[112173]++
															i += k
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2101
			// _ = "end of CoverTab[112173]"
		case WireFixed32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2102
			_go_fuzz_dep_.CoverTab[112174]++
															if len(b)-4 < i {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2103
				_go_fuzz_dep_.CoverTab[112186]++
																return -1, -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2104
				// _ = "end of CoverTab[112186]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2105
				_go_fuzz_dep_.CoverTab[112187]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2105
				// _ = "end of CoverTab[112187]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2105
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2105
			// _ = "end of CoverTab[112174]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2105
			_go_fuzz_dep_.CoverTab[112175]++
															i += 4
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2106
			// _ = "end of CoverTab[112175]"
		case WireFixed64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2107
			_go_fuzz_dep_.CoverTab[112176]++
															if len(b)-8 < i {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2108
				_go_fuzz_dep_.CoverTab[112188]++
																return -1, -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2109
				// _ = "end of CoverTab[112188]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2110
				_go_fuzz_dep_.CoverTab[112189]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2110
				// _ = "end of CoverTab[112189]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2110
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2110
			// _ = "end of CoverTab[112176]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2110
			_go_fuzz_dep_.CoverTab[112177]++
															i += 8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2111
			// _ = "end of CoverTab[112177]"
		case WireBytes:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2112
			_go_fuzz_dep_.CoverTab[112178]++
															m, k := decodeVarint(b[i:])
															if k == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2114
				_go_fuzz_dep_.CoverTab[112190]++
																return -1, -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2115
				// _ = "end of CoverTab[112190]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2116
				_go_fuzz_dep_.CoverTab[112191]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2116
				// _ = "end of CoverTab[112191]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2116
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2116
			// _ = "end of CoverTab[112178]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2116
			_go_fuzz_dep_.CoverTab[112179]++
															i += k
															if uint64(len(b)-i) < m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2118
				_go_fuzz_dep_.CoverTab[112192]++
																return -1, -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2119
				// _ = "end of CoverTab[112192]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2120
				_go_fuzz_dep_.CoverTab[112193]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2120
				// _ = "end of CoverTab[112193]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2120
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2120
			// _ = "end of CoverTab[112179]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2120
			_go_fuzz_dep_.CoverTab[112180]++
															i += int(m)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2121
			// _ = "end of CoverTab[112180]"
		case WireStartGroup:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2122
			_go_fuzz_dep_.CoverTab[112181]++
															depth++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2123
			// _ = "end of CoverTab[112181]"
		case WireEndGroup:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2124
			_go_fuzz_dep_.CoverTab[112182]++
															depth--
															if depth == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2126
				_go_fuzz_dep_.CoverTab[112194]++
																return j, i
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2127
				// _ = "end of CoverTab[112194]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2128
				_go_fuzz_dep_.CoverTab[112195]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2128
				// _ = "end of CoverTab[112195]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2128
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2128
			// _ = "end of CoverTab[112182]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2129
			_go_fuzz_dep_.CoverTab[112183]++
															return -1, -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2130
			// _ = "end of CoverTab[112183]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2131
		// _ = "end of CoverTab[112169]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2132
	// _ = "end of CoverTab[112167]"
}

// encodeVarint appends a varint-encoded integer to b and returns the result.
func encodeVarint(b []byte, x uint64) []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2136
	_go_fuzz_dep_.CoverTab[112196]++
													for x >= 1<<7 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2137
		_go_fuzz_dep_.CoverTab[112198]++
														b = append(b, byte(x&0x7f|0x80))
														x >>= 7
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2139
		// _ = "end of CoverTab[112198]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2140
	// _ = "end of CoverTab[112196]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2140
	_go_fuzz_dep_.CoverTab[112197]++
													return append(b, byte(x))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2141
	// _ = "end of CoverTab[112197]"
}

// decodeVarint reads a varint-encoded integer from b.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2144
// Returns the decoded integer and the number of bytes read.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2144
// If there is an error, it returns 0,0.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2147
func decodeVarint(b []byte) (uint64, int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2147
	_go_fuzz_dep_.CoverTab[112199]++
													var x, y uint64
													if len(b) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2149
		_go_fuzz_dep_.CoverTab[112220]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2150
		// _ = "end of CoverTab[112220]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2151
		_go_fuzz_dep_.CoverTab[112221]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2151
		// _ = "end of CoverTab[112221]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2151
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2151
	// _ = "end of CoverTab[112199]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2151
	_go_fuzz_dep_.CoverTab[112200]++
													x = uint64(b[0])
													if x < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2153
		_go_fuzz_dep_.CoverTab[112222]++
														return x, 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2154
		// _ = "end of CoverTab[112222]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2155
		_go_fuzz_dep_.CoverTab[112223]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2155
		// _ = "end of CoverTab[112223]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2155
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2155
	// _ = "end of CoverTab[112200]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2155
	_go_fuzz_dep_.CoverTab[112201]++
													x -= 0x80

													if len(b) <= 1 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2158
		_go_fuzz_dep_.CoverTab[112224]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2159
		// _ = "end of CoverTab[112224]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2160
		_go_fuzz_dep_.CoverTab[112225]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2160
		// _ = "end of CoverTab[112225]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2160
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2160
	// _ = "end of CoverTab[112201]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2160
	_go_fuzz_dep_.CoverTab[112202]++
													y = uint64(b[1])
													x += y << 7
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2163
		_go_fuzz_dep_.CoverTab[112226]++
														return x, 2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2164
		// _ = "end of CoverTab[112226]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2165
		_go_fuzz_dep_.CoverTab[112227]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2165
		// _ = "end of CoverTab[112227]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2165
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2165
	// _ = "end of CoverTab[112202]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2165
	_go_fuzz_dep_.CoverTab[112203]++
													x -= 0x80 << 7

													if len(b) <= 2 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2168
		_go_fuzz_dep_.CoverTab[112228]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2169
		// _ = "end of CoverTab[112228]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2170
		_go_fuzz_dep_.CoverTab[112229]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2170
		// _ = "end of CoverTab[112229]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2170
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2170
	// _ = "end of CoverTab[112203]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2170
	_go_fuzz_dep_.CoverTab[112204]++
													y = uint64(b[2])
													x += y << 14
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2173
		_go_fuzz_dep_.CoverTab[112230]++
														return x, 3
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2174
		// _ = "end of CoverTab[112230]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2175
		_go_fuzz_dep_.CoverTab[112231]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2175
		// _ = "end of CoverTab[112231]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2175
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2175
	// _ = "end of CoverTab[112204]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2175
	_go_fuzz_dep_.CoverTab[112205]++
													x -= 0x80 << 14

													if len(b) <= 3 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2178
		_go_fuzz_dep_.CoverTab[112232]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2179
		// _ = "end of CoverTab[112232]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2180
		_go_fuzz_dep_.CoverTab[112233]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2180
		// _ = "end of CoverTab[112233]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2180
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2180
	// _ = "end of CoverTab[112205]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2180
	_go_fuzz_dep_.CoverTab[112206]++
													y = uint64(b[3])
													x += y << 21
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2183
		_go_fuzz_dep_.CoverTab[112234]++
														return x, 4
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2184
		// _ = "end of CoverTab[112234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2185
		_go_fuzz_dep_.CoverTab[112235]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2185
		// _ = "end of CoverTab[112235]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2185
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2185
	// _ = "end of CoverTab[112206]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2185
	_go_fuzz_dep_.CoverTab[112207]++
													x -= 0x80 << 21

													if len(b) <= 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2188
		_go_fuzz_dep_.CoverTab[112236]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2189
		// _ = "end of CoverTab[112236]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2190
		_go_fuzz_dep_.CoverTab[112237]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2190
		// _ = "end of CoverTab[112237]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2190
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2190
	// _ = "end of CoverTab[112207]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2190
	_go_fuzz_dep_.CoverTab[112208]++
													y = uint64(b[4])
													x += y << 28
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2193
		_go_fuzz_dep_.CoverTab[112238]++
														return x, 5
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2194
		// _ = "end of CoverTab[112238]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2195
		_go_fuzz_dep_.CoverTab[112239]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2195
		// _ = "end of CoverTab[112239]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2195
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2195
	// _ = "end of CoverTab[112208]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2195
	_go_fuzz_dep_.CoverTab[112209]++
													x -= 0x80 << 28

													if len(b) <= 5 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2198
		_go_fuzz_dep_.CoverTab[112240]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2199
		// _ = "end of CoverTab[112240]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2200
		_go_fuzz_dep_.CoverTab[112241]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2200
		// _ = "end of CoverTab[112241]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2200
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2200
	// _ = "end of CoverTab[112209]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2200
	_go_fuzz_dep_.CoverTab[112210]++
													y = uint64(b[5])
													x += y << 35
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2203
		_go_fuzz_dep_.CoverTab[112242]++
														return x, 6
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2204
		// _ = "end of CoverTab[112242]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2205
		_go_fuzz_dep_.CoverTab[112243]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2205
		// _ = "end of CoverTab[112243]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2205
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2205
	// _ = "end of CoverTab[112210]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2205
	_go_fuzz_dep_.CoverTab[112211]++
													x -= 0x80 << 35

													if len(b) <= 6 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2208
		_go_fuzz_dep_.CoverTab[112244]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2209
		// _ = "end of CoverTab[112244]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2210
		_go_fuzz_dep_.CoverTab[112245]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2210
		// _ = "end of CoverTab[112245]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2210
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2210
	// _ = "end of CoverTab[112211]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2210
	_go_fuzz_dep_.CoverTab[112212]++
													y = uint64(b[6])
													x += y << 42
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2213
		_go_fuzz_dep_.CoverTab[112246]++
														return x, 7
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2214
		// _ = "end of CoverTab[112246]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2215
		_go_fuzz_dep_.CoverTab[112247]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2215
		// _ = "end of CoverTab[112247]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2215
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2215
	// _ = "end of CoverTab[112212]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2215
	_go_fuzz_dep_.CoverTab[112213]++
													x -= 0x80 << 42

													if len(b) <= 7 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2218
		_go_fuzz_dep_.CoverTab[112248]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2219
		// _ = "end of CoverTab[112248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2220
		_go_fuzz_dep_.CoverTab[112249]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2220
		// _ = "end of CoverTab[112249]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2220
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2220
	// _ = "end of CoverTab[112213]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2220
	_go_fuzz_dep_.CoverTab[112214]++
													y = uint64(b[7])
													x += y << 49
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2223
		_go_fuzz_dep_.CoverTab[112250]++
														return x, 8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2224
		// _ = "end of CoverTab[112250]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2225
		_go_fuzz_dep_.CoverTab[112251]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2225
		// _ = "end of CoverTab[112251]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2225
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2225
	// _ = "end of CoverTab[112214]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2225
	_go_fuzz_dep_.CoverTab[112215]++
													x -= 0x80 << 49

													if len(b) <= 8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2228
		_go_fuzz_dep_.CoverTab[112252]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2229
		// _ = "end of CoverTab[112252]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2230
		_go_fuzz_dep_.CoverTab[112253]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2230
		// _ = "end of CoverTab[112253]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2230
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2230
	// _ = "end of CoverTab[112215]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2230
	_go_fuzz_dep_.CoverTab[112216]++
													y = uint64(b[8])
													x += y << 56
													if y < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2233
		_go_fuzz_dep_.CoverTab[112254]++
														return x, 9
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2234
		// _ = "end of CoverTab[112254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2235
		_go_fuzz_dep_.CoverTab[112255]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2235
		// _ = "end of CoverTab[112255]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2235
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2235
	// _ = "end of CoverTab[112216]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2235
	_go_fuzz_dep_.CoverTab[112217]++
													x -= 0x80 << 56

													if len(b) <= 9 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2238
		_go_fuzz_dep_.CoverTab[112256]++
														goto bad
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2239
		// _ = "end of CoverTab[112256]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2240
		_go_fuzz_dep_.CoverTab[112257]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2240
		// _ = "end of CoverTab[112257]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2240
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2240
	// _ = "end of CoverTab[112217]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2240
	_go_fuzz_dep_.CoverTab[112218]++
													y = uint64(b[9])
													x += y << 63
													if y < 2 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2243
		_go_fuzz_dep_.CoverTab[112258]++
														return x, 10
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2244
		// _ = "end of CoverTab[112258]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2245
		_go_fuzz_dep_.CoverTab[112259]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2245
		// _ = "end of CoverTab[112259]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2245
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2245
	// _ = "end of CoverTab[112218]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2245
	_go_fuzz_dep_.CoverTab[112219]++

bad:
													return 0, 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2248
	// _ = "end of CoverTab[112219]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2249
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal.go:2249
var _ = _go_fuzz_dep_.CoverTab
