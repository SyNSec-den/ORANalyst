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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:32
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:38
import (
	"errors"
	"reflect"
)

var (
	// errRepeatedHasNil is the error returned if Marshal is called with
	// a struct with a repeated field containing a nil element.
	errRepeatedHasNil	= errors.New("proto: repeated field has nil element")

	// errOneofHasNil is the error returned if Marshal is called with
	// a struct with a oneof field containing a nil element.
	errOneofHasNil	= errors.New("proto: oneof field has nil value")

	// ErrNil is the error returned if Marshal is called with nil.
	ErrNil	= errors.New("proto: Marshal called with nil")

	// ErrTooLarge is the error returned if Marshal is called with a
	// message that encodes to >2GB.
	ErrTooLarge	= errors.New("proto: message encodes to over 2 GB")
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:64
const maxVarintBytes = 10	// maximum length of a varint

// EncodeVarint returns the varint encoding of x.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:66
// This is the format for the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:66
// int32, int64, uint32, uint64, bool, and enum
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:66
// protocol buffer types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:66
// Not used by the package itself, but helpful to clients
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:66
// wishing to use the same encoding.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:72
func EncodeVarint(x uint64) []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:72
	_go_fuzz_dep_.CoverTab[107992]++
											var buf [maxVarintBytes]byte
											var n int
											for n = 0; x > 127; n++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:75
		_go_fuzz_dep_.CoverTab[107994]++
												buf[n] = 0x80 | uint8(x&0x7F)
												x >>= 7
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:77
		// _ = "end of CoverTab[107994]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:78
	// _ = "end of CoverTab[107992]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:78
	_go_fuzz_dep_.CoverTab[107993]++
											buf[n] = uint8(x)
											n++
											return buf[0:n]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:81
	// _ = "end of CoverTab[107993]"
}

// EncodeVarint writes a varint-encoded integer to the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:84
// This is the format for the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:84
// int32, int64, uint32, uint64, bool, and enum
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:84
// protocol buffer types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:88
func (p *Buffer) EncodeVarint(x uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:88
	_go_fuzz_dep_.CoverTab[107995]++
											for x >= 1<<7 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:89
		_go_fuzz_dep_.CoverTab[107997]++
												p.buf = append(p.buf, uint8(x&0x7f|0x80))
												x >>= 7
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:91
		// _ = "end of CoverTab[107997]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:92
	// _ = "end of CoverTab[107995]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:92
	_go_fuzz_dep_.CoverTab[107996]++
											p.buf = append(p.buf, uint8(x))
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:94
	// _ = "end of CoverTab[107996]"
}

// SizeVarint returns the varint encoding size of an integer.
func SizeVarint(x uint64) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:98
	_go_fuzz_dep_.CoverTab[107998]++
											switch {
	case x < 1<<7:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:100
		_go_fuzz_dep_.CoverTab[108000]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:101
		// _ = "end of CoverTab[108000]"
	case x < 1<<14:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:102
		_go_fuzz_dep_.CoverTab[108001]++
												return 2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:103
		// _ = "end of CoverTab[108001]"
	case x < 1<<21:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:104
		_go_fuzz_dep_.CoverTab[108002]++
												return 3
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:105
		// _ = "end of CoverTab[108002]"
	case x < 1<<28:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:106
		_go_fuzz_dep_.CoverTab[108003]++
												return 4
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:107
		// _ = "end of CoverTab[108003]"
	case x < 1<<35:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:108
		_go_fuzz_dep_.CoverTab[108004]++
												return 5
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:109
		// _ = "end of CoverTab[108004]"
	case x < 1<<42:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:110
		_go_fuzz_dep_.CoverTab[108005]++
												return 6
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:111
		// _ = "end of CoverTab[108005]"
	case x < 1<<49:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:112
		_go_fuzz_dep_.CoverTab[108006]++
												return 7
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:113
		// _ = "end of CoverTab[108006]"
	case x < 1<<56:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:114
		_go_fuzz_dep_.CoverTab[108007]++
												return 8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:115
		// _ = "end of CoverTab[108007]"
	case x < 1<<63:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:116
		_go_fuzz_dep_.CoverTab[108008]++
												return 9
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:117
		// _ = "end of CoverTab[108008]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:117
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:117
		_go_fuzz_dep_.CoverTab[108009]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:117
		// _ = "end of CoverTab[108009]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:118
	// _ = "end of CoverTab[107998]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:118
	_go_fuzz_dep_.CoverTab[107999]++
											return 10
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:119
	// _ = "end of CoverTab[107999]"
}

// EncodeFixed64 writes a 64-bit integer to the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:122
// This is the format for the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:122
// fixed64, sfixed64, and double protocol buffer types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:125
func (p *Buffer) EncodeFixed64(x uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:125
	_go_fuzz_dep_.CoverTab[108010]++
											p.buf = append(p.buf,
		uint8(x),
		uint8(x>>8),
		uint8(x>>16),
		uint8(x>>24),
		uint8(x>>32),
		uint8(x>>40),
		uint8(x>>48),
		uint8(x>>56))
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:135
	// _ = "end of CoverTab[108010]"
}

// EncodeFixed32 writes a 32-bit integer to the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:138
// This is the format for the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:138
// fixed32, sfixed32, and float protocol buffer types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:141
func (p *Buffer) EncodeFixed32(x uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:141
	_go_fuzz_dep_.CoverTab[108011]++
											p.buf = append(p.buf,
		uint8(x),
		uint8(x>>8),
		uint8(x>>16),
		uint8(x>>24))
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:147
	// _ = "end of CoverTab[108011]"
}

// EncodeZigzag64 writes a zigzag-encoded 64-bit integer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:150
// to the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:150
// This is the format used for the sint64 protocol buffer type.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:153
func (p *Buffer) EncodeZigzag64(x uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:153
	_go_fuzz_dep_.CoverTab[108012]++

											return p.EncodeVarint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:155
	// _ = "end of CoverTab[108012]"
}

// EncodeZigzag32 writes a zigzag-encoded 32-bit integer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:158
// to the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:158
// This is the format used for the sint32 protocol buffer type.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:161
func (p *Buffer) EncodeZigzag32(x uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:161
	_go_fuzz_dep_.CoverTab[108013]++

											return p.EncodeVarint(uint64((uint32(x) << 1) ^ uint32((int32(x) >> 31))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:163
	// _ = "end of CoverTab[108013]"
}

// EncodeRawBytes writes a count-delimited byte buffer to the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:166
// This is the format used for the bytes protocol buffer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:166
// type and for embedded messages.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:169
func (p *Buffer) EncodeRawBytes(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:169
	_go_fuzz_dep_.CoverTab[108014]++
											p.EncodeVarint(uint64(len(b)))
											p.buf = append(p.buf, b...)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:172
	// _ = "end of CoverTab[108014]"
}

// EncodeStringBytes writes an encoded string to the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:175
// This is the format used for the proto2 string type.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:177
func (p *Buffer) EncodeStringBytes(s string) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:177
	_go_fuzz_dep_.CoverTab[108015]++
											p.EncodeVarint(uint64(len(s)))
											p.buf = append(p.buf, s...)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:180
	// _ = "end of CoverTab[108015]"
}

// Marshaler is the interface representing objects that can marshal themselves.
type Marshaler interface {
	Marshal() ([]byte, error)
}

// EncodeMessage writes the protocol buffer to the Buffer,
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:188
// prefixed by a varint-encoded length.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:190
func (p *Buffer) EncodeMessage(pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:190
	_go_fuzz_dep_.CoverTab[108016]++
											siz := Size(pb)
											sizVar := SizeVarint(uint64(siz))
											p.grow(siz + sizVar)
											p.EncodeVarint(uint64(siz))
											return p.Marshal(pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:195
	// _ = "end of CoverTab[108016]"
}

// All protocol buffer fields are nillable, but be careful.
func isNil(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:199
	_go_fuzz_dep_.CoverTab[108017]++
											switch v.Kind() {
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:201
		_go_fuzz_dep_.CoverTab[108019]++
												return v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:202
		// _ = "end of CoverTab[108019]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:202
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:202
		_go_fuzz_dep_.CoverTab[108020]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:202
		// _ = "end of CoverTab[108020]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:203
	// _ = "end of CoverTab[108017]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:203
	_go_fuzz_dep_.CoverTab[108018]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:204
	// _ = "end of CoverTab[108018]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:205
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/encode.go:205
var _ = _go_fuzz_dep_.CoverTab
