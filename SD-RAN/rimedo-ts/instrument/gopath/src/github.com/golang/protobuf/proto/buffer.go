// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:5
)

import (
	"errors"
	"fmt"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	WireVarint	= 0
	WireFixed32	= 5
	WireFixed64	= 1
	WireBytes	= 2
	WireStartGroup	= 3
	WireEndGroup	= 4
)

// EncodeVarint returns the varint encoded bytes of v.
func EncodeVarint(v uint64) []byte {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:26
	_go_fuzz_dep_.CoverTab[61111]++
											return protowire.AppendVarint(nil, v)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:27
	// _ = "end of CoverTab[61111]"
}

// SizeVarint returns the length of the varint encoded bytes of v.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:30
// This is equal to len(EncodeVarint(v)).
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:32
func SizeVarint(v uint64) int {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:32
	_go_fuzz_dep_.CoverTab[61112]++
											return protowire.SizeVarint(v)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:33
	// _ = "end of CoverTab[61112]"
}

// DecodeVarint parses a varint encoded integer from b,
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:36
// returning the integer value and the length of the varint.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:36
// It returns (0, 0) if there is a parse error.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:39
func DecodeVarint(b []byte) (uint64, int) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:39
	_go_fuzz_dep_.CoverTab[61113]++
											v, n := protowire.ConsumeVarint(b)
											if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:41
		_go_fuzz_dep_.CoverTab[61115]++
												return 0, 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:42
		// _ = "end of CoverTab[61115]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:43
		_go_fuzz_dep_.CoverTab[61116]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:43
		// _ = "end of CoverTab[61116]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:43
	// _ = "end of CoverTab[61113]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:43
	_go_fuzz_dep_.CoverTab[61114]++
											return v, n
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:44
	// _ = "end of CoverTab[61114]"
}

// Buffer is a buffer for encoding and decoding the protobuf wire format.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:47
// It may be reused between invocations to reduce memory usage.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:49
type Buffer struct {
	buf		[]byte
	idx		int
	deterministic	bool
}

// NewBuffer allocates a new Buffer initialized with buf,
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:55
// where the contents of buf are considered the unread portion of the buffer.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:57
func NewBuffer(buf []byte) *Buffer {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:57
	_go_fuzz_dep_.CoverTab[61117]++
											return &Buffer{buf: buf}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:58
	// _ = "end of CoverTab[61117]"
}

// SetDeterministic specifies whether to use deterministic serialization.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// Deterministic serialization guarantees that for a given binary, equal
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// messages will always be serialized to the same bytes. This implies:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
//   - Repeated serialization of a message will return the same bytes.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
//   - Different processes of the same binary (which may be executing on
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
//     different machines) will serialize equal messages to the same bytes.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// Note that the deterministic serialization is NOT canonical across
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// languages. It is not guaranteed to remain stable over time. It is unstable
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// across different builds with schema changes due to unknown fields.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// Users who need canonical serialization (e.g., persistent storage in a
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// canonical form, fingerprinting, etc.) should define their own
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// canonicalization specification and implement their own serializer rather
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// than relying on this API.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// If deterministic serialization is requested, map entries will be sorted
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// by keys in lexographical order. This is an implementation detail and
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:61
// subject to change.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:81
func (b *Buffer) SetDeterministic(deterministic bool) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:81
	_go_fuzz_dep_.CoverTab[61118]++
											b.deterministic = deterministic
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:82
	// _ = "end of CoverTab[61118]"
}

// SetBuf sets buf as the internal buffer,
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:85
// where the contents of buf are considered the unread portion of the buffer.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:87
func (b *Buffer) SetBuf(buf []byte) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:87
	_go_fuzz_dep_.CoverTab[61119]++
											b.buf = buf
											b.idx = 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:89
	// _ = "end of CoverTab[61119]"
}

// Reset clears the internal buffer of all written and unread data.
func (b *Buffer) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:93
	_go_fuzz_dep_.CoverTab[61120]++
											b.buf = b.buf[:0]
											b.idx = 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:95
	// _ = "end of CoverTab[61120]"
}

// Bytes returns the internal buffer.
func (b *Buffer) Bytes() []byte {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:99
		_go_fuzz_dep_.CoverTab[61121]++
												return b.buf
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:100
	// _ = "end of CoverTab[61121]"
}

// Unread returns the unread portion of the buffer.
func (b *Buffer) Unread() []byte {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:104
	_go_fuzz_dep_.CoverTab[61122]++
												return b.buf[b.idx:]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:105
	// _ = "end of CoverTab[61122]"
}

// Marshal appends the wire-format encoding of m to the buffer.
func (b *Buffer) Marshal(m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:109
	_go_fuzz_dep_.CoverTab[61123]++
												var err error
												b.buf, err = marshalAppend(b.buf, m, b.deterministic)
												return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:112
	// _ = "end of CoverTab[61123]"
}

// Unmarshal parses the wire-format message in the buffer and
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:115
// places the decoded results in m.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:115
// It does not reset m before unmarshaling.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:118
func (b *Buffer) Unmarshal(m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:118
	_go_fuzz_dep_.CoverTab[61124]++
												err := UnmarshalMerge(b.Unread(), m)
												b.idx = len(b.buf)
												return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:121
	// _ = "end of CoverTab[61124]"
}

type unknownFields struct{ XXX_unrecognized protoimpl.UnknownFields }

func (m *unknownFields) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:126
	_go_fuzz_dep_.CoverTab[61125]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:126
	panic("not implemented")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:126
	// _ = "end of CoverTab[61125]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:126
}
func (m *unknownFields) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:127
	_go_fuzz_dep_.CoverTab[61126]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:127
	panic("not implemented")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:127
	// _ = "end of CoverTab[61126]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:127
}
func (m *unknownFields) ProtoMessage() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:128
	_go_fuzz_dep_.CoverTab[61127]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:128
	panic("not implemented")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:128
	// _ = "end of CoverTab[61127]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:128
}

// DebugPrint dumps the encoded bytes of b with a header and footer including s
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:130
// to stdout. This is only intended for debugging.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:132
func (*Buffer) DebugPrint(s string, b []byte) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:132
	_go_fuzz_dep_.CoverTab[61128]++
												m := MessageReflect(new(unknownFields))
												m.SetUnknown(b)
												b, _ = prototext.MarshalOptions{AllowPartial: true, Indent: "\t"}.Marshal(m.Interface())
												fmt.Printf("==== %s ====\n%s==== %s ====\n", s, b, s)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:136
	// _ = "end of CoverTab[61128]"
}

// EncodeVarint appends an unsigned varint encoding to the buffer.
func (b *Buffer) EncodeVarint(v uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:140
	_go_fuzz_dep_.CoverTab[61129]++
												b.buf = protowire.AppendVarint(b.buf, v)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:142
	// _ = "end of CoverTab[61129]"
}

// EncodeZigzag32 appends a 32-bit zig-zag varint encoding to the buffer.
func (b *Buffer) EncodeZigzag32(v uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:146
	_go_fuzz_dep_.CoverTab[61130]++
												return b.EncodeVarint(uint64((uint32(v) << 1) ^ uint32((int32(v) >> 31))))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:147
	// _ = "end of CoverTab[61130]"
}

// EncodeZigzag64 appends a 64-bit zig-zag varint encoding to the buffer.
func (b *Buffer) EncodeZigzag64(v uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:151
	_go_fuzz_dep_.CoverTab[61131]++
												return b.EncodeVarint(uint64((uint64(v) << 1) ^ uint64((int64(v) >> 63))))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:152
	// _ = "end of CoverTab[61131]"
}

// EncodeFixed32 appends a 32-bit little-endian integer to the buffer.
func (b *Buffer) EncodeFixed32(v uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:156
	_go_fuzz_dep_.CoverTab[61132]++
												b.buf = protowire.AppendFixed32(b.buf, uint32(v))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:158
	// _ = "end of CoverTab[61132]"
}

// EncodeFixed64 appends a 64-bit little-endian integer to the buffer.
func (b *Buffer) EncodeFixed64(v uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:162
	_go_fuzz_dep_.CoverTab[61133]++
												b.buf = protowire.AppendFixed64(b.buf, uint64(v))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:164
	// _ = "end of CoverTab[61133]"
}

// EncodeRawBytes appends a length-prefixed raw bytes to the buffer.
func (b *Buffer) EncodeRawBytes(v []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:168
	_go_fuzz_dep_.CoverTab[61134]++
												b.buf = protowire.AppendBytes(b.buf, v)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:170
	// _ = "end of CoverTab[61134]"
}

// EncodeStringBytes appends a length-prefixed raw bytes to the buffer.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:173
// It does not validate whether v contains valid UTF-8.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:175
func (b *Buffer) EncodeStringBytes(v string) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:175
	_go_fuzz_dep_.CoverTab[61135]++
												b.buf = protowire.AppendString(b.buf, v)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:177
	// _ = "end of CoverTab[61135]"
}

// EncodeMessage appends a length-prefixed encoded message to the buffer.
func (b *Buffer) EncodeMessage(m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:181
	_go_fuzz_dep_.CoverTab[61136]++
												var err error
												b.buf = protowire.AppendVarint(b.buf, uint64(Size(m)))
												b.buf, err = marshalAppend(b.buf, m, b.deterministic)
												return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:185
	// _ = "end of CoverTab[61136]"
}

// DecodeVarint consumes an encoded unsigned varint from the buffer.
func (b *Buffer) DecodeVarint() (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:189
	_go_fuzz_dep_.CoverTab[61137]++
												v, n := protowire.ConsumeVarint(b.buf[b.idx:])
												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:191
		_go_fuzz_dep_.CoverTab[61139]++
													return 0, protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:192
		// _ = "end of CoverTab[61139]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:193
		_go_fuzz_dep_.CoverTab[61140]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:193
		// _ = "end of CoverTab[61140]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:193
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:193
	// _ = "end of CoverTab[61137]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:193
	_go_fuzz_dep_.CoverTab[61138]++
												b.idx += n
												return uint64(v), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:195
	// _ = "end of CoverTab[61138]"
}

// DecodeZigzag32 consumes an encoded 32-bit zig-zag varint from the buffer.
func (b *Buffer) DecodeZigzag32() (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:199
	_go_fuzz_dep_.CoverTab[61141]++
												v, err := b.DecodeVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:201
		_go_fuzz_dep_.CoverTab[61143]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:202
		// _ = "end of CoverTab[61143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:203
		_go_fuzz_dep_.CoverTab[61144]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:203
		// _ = "end of CoverTab[61144]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:203
	// _ = "end of CoverTab[61141]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:203
	_go_fuzz_dep_.CoverTab[61142]++
												return uint64((uint32(v) >> 1) ^ uint32((int32(v&1)<<31)>>31)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:204
	// _ = "end of CoverTab[61142]"
}

// DecodeZigzag64 consumes an encoded 64-bit zig-zag varint from the buffer.
func (b *Buffer) DecodeZigzag64() (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:208
	_go_fuzz_dep_.CoverTab[61145]++
												v, err := b.DecodeVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:210
		_go_fuzz_dep_.CoverTab[61147]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:211
		// _ = "end of CoverTab[61147]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:212
		_go_fuzz_dep_.CoverTab[61148]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:212
		// _ = "end of CoverTab[61148]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:212
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:212
	// _ = "end of CoverTab[61145]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:212
	_go_fuzz_dep_.CoverTab[61146]++
												return uint64((uint64(v) >> 1) ^ uint64((int64(v&1)<<63)>>63)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:213
	// _ = "end of CoverTab[61146]"
}

// DecodeFixed32 consumes a 32-bit little-endian integer from the buffer.
func (b *Buffer) DecodeFixed32() (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:217
	_go_fuzz_dep_.CoverTab[61149]++
												v, n := protowire.ConsumeFixed32(b.buf[b.idx:])
												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:219
		_go_fuzz_dep_.CoverTab[61151]++
													return 0, protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:220
		// _ = "end of CoverTab[61151]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:221
		_go_fuzz_dep_.CoverTab[61152]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:221
		// _ = "end of CoverTab[61152]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:221
	// _ = "end of CoverTab[61149]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:221
	_go_fuzz_dep_.CoverTab[61150]++
												b.idx += n
												return uint64(v), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:223
	// _ = "end of CoverTab[61150]"
}

// DecodeFixed64 consumes a 64-bit little-endian integer from the buffer.
func (b *Buffer) DecodeFixed64() (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:227
	_go_fuzz_dep_.CoverTab[61153]++
												v, n := protowire.ConsumeFixed64(b.buf[b.idx:])
												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:229
		_go_fuzz_dep_.CoverTab[61155]++
													return 0, protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:230
		// _ = "end of CoverTab[61155]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:231
		_go_fuzz_dep_.CoverTab[61156]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:231
		// _ = "end of CoverTab[61156]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:231
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:231
	// _ = "end of CoverTab[61153]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:231
	_go_fuzz_dep_.CoverTab[61154]++
												b.idx += n
												return uint64(v), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:233
	// _ = "end of CoverTab[61154]"
}

// DecodeRawBytes consumes a length-prefixed raw bytes from the buffer.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:236
// If alloc is specified, it returns a copy the raw bytes
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:236
// rather than a sub-slice of the buffer.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:239
func (b *Buffer) DecodeRawBytes(alloc bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:239
	_go_fuzz_dep_.CoverTab[61157]++
												v, n := protowire.ConsumeBytes(b.buf[b.idx:])
												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:241
		_go_fuzz_dep_.CoverTab[61160]++
													return nil, protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:242
		// _ = "end of CoverTab[61160]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:243
		_go_fuzz_dep_.CoverTab[61161]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:243
		// _ = "end of CoverTab[61161]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:243
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:243
	// _ = "end of CoverTab[61157]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:243
	_go_fuzz_dep_.CoverTab[61158]++
												b.idx += n
												if alloc {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:245
		_go_fuzz_dep_.CoverTab[61162]++
													v = append([]byte(nil), v...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:246
		// _ = "end of CoverTab[61162]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:247
		_go_fuzz_dep_.CoverTab[61163]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:247
		// _ = "end of CoverTab[61163]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:247
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:247
	// _ = "end of CoverTab[61158]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:247
	_go_fuzz_dep_.CoverTab[61159]++
												return v, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:248
	// _ = "end of CoverTab[61159]"
}

// DecodeStringBytes consumes a length-prefixed raw bytes from the buffer.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:251
// It does not validate whether the raw bytes contain valid UTF-8.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:253
func (b *Buffer) DecodeStringBytes() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:253
	_go_fuzz_dep_.CoverTab[61164]++
												v, n := protowire.ConsumeString(b.buf[b.idx:])
												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:255
		_go_fuzz_dep_.CoverTab[61166]++
													return "", protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:256
		// _ = "end of CoverTab[61166]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:257
		_go_fuzz_dep_.CoverTab[61167]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:257
		// _ = "end of CoverTab[61167]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:257
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:257
	// _ = "end of CoverTab[61164]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:257
	_go_fuzz_dep_.CoverTab[61165]++
												b.idx += n
												return v, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:259
	// _ = "end of CoverTab[61165]"
}

// DecodeMessage consumes a length-prefixed message from the buffer.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:262
// It does not reset m before unmarshaling.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:264
func (b *Buffer) DecodeMessage(m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:264
	_go_fuzz_dep_.CoverTab[61168]++
												v, err := b.DecodeRawBytes(false)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:266
		_go_fuzz_dep_.CoverTab[61170]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:267
		// _ = "end of CoverTab[61170]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:268
		_go_fuzz_dep_.CoverTab[61171]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:268
		// _ = "end of CoverTab[61171]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:268
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:268
	// _ = "end of CoverTab[61168]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:268
	_go_fuzz_dep_.CoverTab[61169]++
												return UnmarshalMerge(v, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:269
	// _ = "end of CoverTab[61169]"
}

// DecodeGroup consumes a message group from the buffer.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:272
// It assumes that the start group marker has already been consumed and
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:272
// consumes all bytes until (and including the end group marker).
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:272
// It does not reset m before unmarshaling.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:276
func (b *Buffer) DecodeGroup(m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:276
	_go_fuzz_dep_.CoverTab[61172]++
												v, n, err := consumeGroup(b.buf[b.idx:])
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:278
		_go_fuzz_dep_.CoverTab[61174]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:279
		// _ = "end of CoverTab[61174]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:280
		_go_fuzz_dep_.CoverTab[61175]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:280
		// _ = "end of CoverTab[61175]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:280
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:280
	// _ = "end of CoverTab[61172]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:280
	_go_fuzz_dep_.CoverTab[61173]++
												b.idx += n
												return UnmarshalMerge(v, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:282
	// _ = "end of CoverTab[61173]"
}

// consumeGroup parses b until it finds an end group marker, returning
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:285
// the raw bytes of the message (excluding the end group marker) and the
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:285
// the total length of the message (including the end group marker).
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:288
func consumeGroup(b []byte) ([]byte, int, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:288
	_go_fuzz_dep_.CoverTab[61176]++
												b0 := b
												depth := 1
												for {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:291
		_go_fuzz_dep_.CoverTab[61177]++
													_, wtyp, tagLen := protowire.ConsumeTag(b)
													if tagLen < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:293
			_go_fuzz_dep_.CoverTab[61181]++
														return nil, 0, protowire.ParseError(tagLen)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:294
			// _ = "end of CoverTab[61181]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:295
			_go_fuzz_dep_.CoverTab[61182]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:295
			// _ = "end of CoverTab[61182]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:295
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:295
		// _ = "end of CoverTab[61177]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:295
		_go_fuzz_dep_.CoverTab[61178]++
													b = b[tagLen:]

													var valLen int
													switch wtyp {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:300
			_go_fuzz_dep_.CoverTab[61183]++
														_, valLen = protowire.ConsumeVarint(b)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:301
			// _ = "end of CoverTab[61183]"
		case protowire.Fixed32Type:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:302
			_go_fuzz_dep_.CoverTab[61184]++
														_, valLen = protowire.ConsumeFixed32(b)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:303
			// _ = "end of CoverTab[61184]"
		case protowire.Fixed64Type:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:304
			_go_fuzz_dep_.CoverTab[61185]++
														_, valLen = protowire.ConsumeFixed64(b)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:305
			// _ = "end of CoverTab[61185]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:306
			_go_fuzz_dep_.CoverTab[61186]++
														_, valLen = protowire.ConsumeBytes(b)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:307
			// _ = "end of CoverTab[61186]"
		case protowire.StartGroupType:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:308
			_go_fuzz_dep_.CoverTab[61187]++
														depth++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:309
			// _ = "end of CoverTab[61187]"
		case protowire.EndGroupType:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:310
			_go_fuzz_dep_.CoverTab[61188]++
														depth--
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:311
			// _ = "end of CoverTab[61188]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:312
			_go_fuzz_dep_.CoverTab[61189]++
														return nil, 0, errors.New("proto: cannot parse reserved wire type")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:313
			// _ = "end of CoverTab[61189]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:314
		// _ = "end of CoverTab[61178]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:314
		_go_fuzz_dep_.CoverTab[61179]++
													if valLen < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:315
			_go_fuzz_dep_.CoverTab[61190]++
														return nil, 0, protowire.ParseError(valLen)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:316
			// _ = "end of CoverTab[61190]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:317
			_go_fuzz_dep_.CoverTab[61191]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:317
			// _ = "end of CoverTab[61191]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:317
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:317
		// _ = "end of CoverTab[61179]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:317
		_go_fuzz_dep_.CoverTab[61180]++
													b = b[valLen:]

													if depth == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:320
			_go_fuzz_dep_.CoverTab[61192]++
														return b0[:len(b0)-len(b)-tagLen], len(b0) - len(b), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:321
			// _ = "end of CoverTab[61192]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:322
			_go_fuzz_dep_.CoverTab[61193]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:322
			// _ = "end of CoverTab[61193]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:322
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:322
		// _ = "end of CoverTab[61180]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:323
	// _ = "end of CoverTab[61176]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:324
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/buffer.go:324
var _ = _go_fuzz_dep_.CoverTab
