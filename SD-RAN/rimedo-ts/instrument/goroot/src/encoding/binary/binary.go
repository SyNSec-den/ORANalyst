// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/binary/binary.go:5
// Package binary implements simple translation between numbers and byte
//line /usr/local/go/src/encoding/binary/binary.go:5
// sequences and encoding and decoding of varints.
//line /usr/local/go/src/encoding/binary/binary.go:5
//
//line /usr/local/go/src/encoding/binary/binary.go:5
// Numbers are translated by reading and writing fixed-size values.
//line /usr/local/go/src/encoding/binary/binary.go:5
// A fixed-size value is either a fixed-size arithmetic
//line /usr/local/go/src/encoding/binary/binary.go:5
// type (bool, int8, uint8, int16, float32, complex64, ...)
//line /usr/local/go/src/encoding/binary/binary.go:5
// or an array or struct containing only fixed-size values.
//line /usr/local/go/src/encoding/binary/binary.go:5
//
//line /usr/local/go/src/encoding/binary/binary.go:5
// The varint functions encode and decode single integer values using
//line /usr/local/go/src/encoding/binary/binary.go:5
// a variable-length encoding; smaller values require fewer bytes.
//line /usr/local/go/src/encoding/binary/binary.go:5
// For a specification, see
//line /usr/local/go/src/encoding/binary/binary.go:5
// https://developers.google.com/protocol-buffers/docs/encoding.
//line /usr/local/go/src/encoding/binary/binary.go:5
//
//line /usr/local/go/src/encoding/binary/binary.go:5
// This package favors simplicity over efficiency. Clients that require
//line /usr/local/go/src/encoding/binary/binary.go:5
// high-performance serialization, especially for large data structures,
//line /usr/local/go/src/encoding/binary/binary.go:5
// should look at more advanced solutions such as the encoding/gob
//line /usr/local/go/src/encoding/binary/binary.go:5
// package or protocol buffers.
//line /usr/local/go/src/encoding/binary/binary.go:22
package binary

//line /usr/local/go/src/encoding/binary/binary.go:22
import (
//line /usr/local/go/src/encoding/binary/binary.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/binary/binary.go:22
)
//line /usr/local/go/src/encoding/binary/binary.go:22
import (
//line /usr/local/go/src/encoding/binary/binary.go:22
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/binary/binary.go:22
)

import (
	"errors"
	"io"
	"math"
	"reflect"
	"sync"
)

// A ByteOrder specifies how to convert byte slices into
//line /usr/local/go/src/encoding/binary/binary.go:32
// 16-, 32-, or 64-bit unsigned integers.
//line /usr/local/go/src/encoding/binary/binary.go:34
type ByteOrder interface {
	Uint16([]byte) uint16
	Uint32([]byte) uint32
	Uint64([]byte) uint64
	PutUint16([]byte, uint16)
	PutUint32([]byte, uint32)
	PutUint64([]byte, uint64)
	String() string
}

// AppendByteOrder specifies how to append 16-, 32-, or 64-bit unsigned integers
//line /usr/local/go/src/encoding/binary/binary.go:44
// into a byte slice.
//line /usr/local/go/src/encoding/binary/binary.go:46
type AppendByteOrder interface {
	AppendUint16([]byte, uint16) []byte
	AppendUint32([]byte, uint32) []byte
	AppendUint64([]byte, uint64) []byte
	String() string
}

// LittleEndian is the little-endian implementation of ByteOrder and AppendByteOrder.
var LittleEndian littleEndian

// BigEndian is the big-endian implementation of ByteOrder and AppendByteOrder.
var BigEndian bigEndian

type littleEndian struct{}

func (littleEndian) Uint16(b []byte) uint16 {
//line /usr/local/go/src/encoding/binary/binary.go:61
	_go_fuzz_dep_.CoverTab[1170]++
							_ = b[1]
							return uint16(b[0]) | uint16(b[1])<<8
//line /usr/local/go/src/encoding/binary/binary.go:63
	// _ = "end of CoverTab[1170]"
}

func (littleEndian) PutUint16(b []byte, v uint16) {
//line /usr/local/go/src/encoding/binary/binary.go:66
	_go_fuzz_dep_.CoverTab[1171]++
							_ = b[1]
							b[0] = byte(v)
							b[1] = byte(v >> 8)
//line /usr/local/go/src/encoding/binary/binary.go:69
	// _ = "end of CoverTab[1171]"
}

func (littleEndian) AppendUint16(b []byte, v uint16) []byte {
//line /usr/local/go/src/encoding/binary/binary.go:72
	_go_fuzz_dep_.CoverTab[1172]++
							return append(b,
		byte(v),
		byte(v>>8),
	)
//line /usr/local/go/src/encoding/binary/binary.go:76
	// _ = "end of CoverTab[1172]"
}

func (littleEndian) Uint32(b []byte) uint32 {
//line /usr/local/go/src/encoding/binary/binary.go:79
	_go_fuzz_dep_.CoverTab[1173]++
							_ = b[3]
							return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
//line /usr/local/go/src/encoding/binary/binary.go:81
	// _ = "end of CoverTab[1173]"
}

func (littleEndian) PutUint32(b []byte, v uint32) {
//line /usr/local/go/src/encoding/binary/binary.go:84
	_go_fuzz_dep_.CoverTab[1174]++
							_ = b[3]
							b[0] = byte(v)
							b[1] = byte(v >> 8)
							b[2] = byte(v >> 16)
							b[3] = byte(v >> 24)
//line /usr/local/go/src/encoding/binary/binary.go:89
	// _ = "end of CoverTab[1174]"
}

func (littleEndian) AppendUint32(b []byte, v uint32) []byte {
//line /usr/local/go/src/encoding/binary/binary.go:92
	_go_fuzz_dep_.CoverTab[1175]++
							return append(b,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
	)
//line /usr/local/go/src/encoding/binary/binary.go:98
	// _ = "end of CoverTab[1175]"
}

func (littleEndian) Uint64(b []byte) uint64 {
//line /usr/local/go/src/encoding/binary/binary.go:101
	_go_fuzz_dep_.CoverTab[1176]++
							_ = b[7]
							return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
//line /usr/local/go/src/encoding/binary/binary.go:104
	// _ = "end of CoverTab[1176]"
}

func (littleEndian) PutUint64(b []byte, v uint64) {
//line /usr/local/go/src/encoding/binary/binary.go:107
	_go_fuzz_dep_.CoverTab[1177]++
							_ = b[7]
							b[0] = byte(v)
							b[1] = byte(v >> 8)
							b[2] = byte(v >> 16)
							b[3] = byte(v >> 24)
							b[4] = byte(v >> 32)
							b[5] = byte(v >> 40)
							b[6] = byte(v >> 48)
							b[7] = byte(v >> 56)
//line /usr/local/go/src/encoding/binary/binary.go:116
	// _ = "end of CoverTab[1177]"
}

func (littleEndian) AppendUint64(b []byte, v uint64) []byte {
//line /usr/local/go/src/encoding/binary/binary.go:119
	_go_fuzz_dep_.CoverTab[1178]++
							return append(b,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
		byte(v>>32),
		byte(v>>40),
		byte(v>>48),
		byte(v>>56),
	)
//line /usr/local/go/src/encoding/binary/binary.go:129
	// _ = "end of CoverTab[1178]"
}

func (littleEndian) String() string {
//line /usr/local/go/src/encoding/binary/binary.go:132
	_go_fuzz_dep_.CoverTab[1179]++
//line /usr/local/go/src/encoding/binary/binary.go:132
	return "LittleEndian"
//line /usr/local/go/src/encoding/binary/binary.go:132
	// _ = "end of CoverTab[1179]"
//line /usr/local/go/src/encoding/binary/binary.go:132
}

func (littleEndian) GoString() string {
//line /usr/local/go/src/encoding/binary/binary.go:134
	_go_fuzz_dep_.CoverTab[1180]++
//line /usr/local/go/src/encoding/binary/binary.go:134
	return "binary.LittleEndian"
//line /usr/local/go/src/encoding/binary/binary.go:134
	// _ = "end of CoverTab[1180]"
//line /usr/local/go/src/encoding/binary/binary.go:134
}

type bigEndian struct{}

func (bigEndian) Uint16(b []byte) uint16 {
//line /usr/local/go/src/encoding/binary/binary.go:138
	_go_fuzz_dep_.CoverTab[1181]++
							_ = b[1]
							return uint16(b[1]) | uint16(b[0])<<8
//line /usr/local/go/src/encoding/binary/binary.go:140
	// _ = "end of CoverTab[1181]"
}

func (bigEndian) PutUint16(b []byte, v uint16) {
//line /usr/local/go/src/encoding/binary/binary.go:143
	_go_fuzz_dep_.CoverTab[1182]++
							_ = b[1]
							b[0] = byte(v >> 8)
							b[1] = byte(v)
//line /usr/local/go/src/encoding/binary/binary.go:146
	// _ = "end of CoverTab[1182]"
}

func (bigEndian) AppendUint16(b []byte, v uint16) []byte {
//line /usr/local/go/src/encoding/binary/binary.go:149
	_go_fuzz_dep_.CoverTab[1183]++
							return append(b,
		byte(v>>8),
		byte(v),
	)
//line /usr/local/go/src/encoding/binary/binary.go:153
	// _ = "end of CoverTab[1183]"
}

func (bigEndian) Uint32(b []byte) uint32 {
//line /usr/local/go/src/encoding/binary/binary.go:156
	_go_fuzz_dep_.CoverTab[1184]++
							_ = b[3]
							return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
//line /usr/local/go/src/encoding/binary/binary.go:158
	// _ = "end of CoverTab[1184]"
}

func (bigEndian) PutUint32(b []byte, v uint32) {
//line /usr/local/go/src/encoding/binary/binary.go:161
	_go_fuzz_dep_.CoverTab[1185]++
							_ = b[3]
							b[0] = byte(v >> 24)
							b[1] = byte(v >> 16)
							b[2] = byte(v >> 8)
							b[3] = byte(v)
//line /usr/local/go/src/encoding/binary/binary.go:166
	// _ = "end of CoverTab[1185]"
}

func (bigEndian) AppendUint32(b []byte, v uint32) []byte {
//line /usr/local/go/src/encoding/binary/binary.go:169
	_go_fuzz_dep_.CoverTab[1186]++
							return append(b,
		byte(v>>24),
		byte(v>>16),
		byte(v>>8),
		byte(v),
	)
//line /usr/local/go/src/encoding/binary/binary.go:175
	// _ = "end of CoverTab[1186]"
}

func (bigEndian) Uint64(b []byte) uint64 {
//line /usr/local/go/src/encoding/binary/binary.go:178
	_go_fuzz_dep_.CoverTab[1187]++
							_ = b[7]
							return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
//line /usr/local/go/src/encoding/binary/binary.go:181
	// _ = "end of CoverTab[1187]"
}

func (bigEndian) PutUint64(b []byte, v uint64) {
//line /usr/local/go/src/encoding/binary/binary.go:184
	_go_fuzz_dep_.CoverTab[1188]++
							_ = b[7]
							b[0] = byte(v >> 56)
							b[1] = byte(v >> 48)
							b[2] = byte(v >> 40)
							b[3] = byte(v >> 32)
							b[4] = byte(v >> 24)
							b[5] = byte(v >> 16)
							b[6] = byte(v >> 8)
							b[7] = byte(v)
//line /usr/local/go/src/encoding/binary/binary.go:193
	// _ = "end of CoverTab[1188]"
}

func (bigEndian) AppendUint64(b []byte, v uint64) []byte {
//line /usr/local/go/src/encoding/binary/binary.go:196
	_go_fuzz_dep_.CoverTab[1189]++
							return append(b,
		byte(v>>56),
		byte(v>>48),
		byte(v>>40),
		byte(v>>32),
		byte(v>>24),
		byte(v>>16),
		byte(v>>8),
		byte(v),
	)
//line /usr/local/go/src/encoding/binary/binary.go:206
	// _ = "end of CoverTab[1189]"
}

func (bigEndian) String() string {
//line /usr/local/go/src/encoding/binary/binary.go:209
	_go_fuzz_dep_.CoverTab[1190]++
//line /usr/local/go/src/encoding/binary/binary.go:209
	return "BigEndian"
//line /usr/local/go/src/encoding/binary/binary.go:209
	// _ = "end of CoverTab[1190]"
//line /usr/local/go/src/encoding/binary/binary.go:209
}

func (bigEndian) GoString() string {
//line /usr/local/go/src/encoding/binary/binary.go:211
	_go_fuzz_dep_.CoverTab[1191]++
//line /usr/local/go/src/encoding/binary/binary.go:211
	return "binary.BigEndian"
//line /usr/local/go/src/encoding/binary/binary.go:211
	// _ = "end of CoverTab[1191]"
//line /usr/local/go/src/encoding/binary/binary.go:211
}

// Read reads structured binary data from r into data.
//line /usr/local/go/src/encoding/binary/binary.go:213
// Data must be a pointer to a fixed-size value or a slice
//line /usr/local/go/src/encoding/binary/binary.go:213
// of fixed-size values.
//line /usr/local/go/src/encoding/binary/binary.go:213
// Bytes read from r are decoded using the specified byte order
//line /usr/local/go/src/encoding/binary/binary.go:213
// and written to successive fields of the data.
//line /usr/local/go/src/encoding/binary/binary.go:213
// When decoding boolean values, a zero byte is decoded as false, and
//line /usr/local/go/src/encoding/binary/binary.go:213
// any other non-zero byte is decoded as true.
//line /usr/local/go/src/encoding/binary/binary.go:213
// When reading into structs, the field data for fields with
//line /usr/local/go/src/encoding/binary/binary.go:213
// blank (_) field names is skipped; i.e., blank field names
//line /usr/local/go/src/encoding/binary/binary.go:213
// may be used for padding.
//line /usr/local/go/src/encoding/binary/binary.go:213
// When reading into a struct, all non-blank fields must be exported
//line /usr/local/go/src/encoding/binary/binary.go:213
// or Read may panic.
//line /usr/local/go/src/encoding/binary/binary.go:213
//
//line /usr/local/go/src/encoding/binary/binary.go:213
// The error is EOF only if no bytes were read.
//line /usr/local/go/src/encoding/binary/binary.go:213
// If an EOF happens after reading some but not all the bytes,
//line /usr/local/go/src/encoding/binary/binary.go:213
// Read returns ErrUnexpectedEOF.
//line /usr/local/go/src/encoding/binary/binary.go:229
func Read(r io.Reader, order ByteOrder, data any) error {
//line /usr/local/go/src/encoding/binary/binary.go:229
	_go_fuzz_dep_.CoverTab[1192]++

							if n := intDataSize(data); n != 0 {
//line /usr/local/go/src/encoding/binary/binary.go:231
		_go_fuzz_dep_.CoverTab[1197]++
								bs := make([]byte, n)
								if _, err := io.ReadFull(r, bs); err != nil {
//line /usr/local/go/src/encoding/binary/binary.go:233
			_go_fuzz_dep_.CoverTab[1200]++
									return err
//line /usr/local/go/src/encoding/binary/binary.go:234
			// _ = "end of CoverTab[1200]"
		} else {
//line /usr/local/go/src/encoding/binary/binary.go:235
			_go_fuzz_dep_.CoverTab[1201]++
//line /usr/local/go/src/encoding/binary/binary.go:235
			// _ = "end of CoverTab[1201]"
//line /usr/local/go/src/encoding/binary/binary.go:235
		}
//line /usr/local/go/src/encoding/binary/binary.go:235
		// _ = "end of CoverTab[1197]"
//line /usr/local/go/src/encoding/binary/binary.go:235
		_go_fuzz_dep_.CoverTab[1198]++
								switch data := data.(type) {
		case *bool:
//line /usr/local/go/src/encoding/binary/binary.go:237
			_go_fuzz_dep_.CoverTab[1202]++
									*data = bs[0] != 0
//line /usr/local/go/src/encoding/binary/binary.go:238
			// _ = "end of CoverTab[1202]"
		case *int8:
//line /usr/local/go/src/encoding/binary/binary.go:239
			_go_fuzz_dep_.CoverTab[1203]++
									*data = int8(bs[0])
//line /usr/local/go/src/encoding/binary/binary.go:240
			// _ = "end of CoverTab[1203]"
		case *uint8:
//line /usr/local/go/src/encoding/binary/binary.go:241
			_go_fuzz_dep_.CoverTab[1204]++
									*data = bs[0]
//line /usr/local/go/src/encoding/binary/binary.go:242
			// _ = "end of CoverTab[1204]"
		case *int16:
//line /usr/local/go/src/encoding/binary/binary.go:243
			_go_fuzz_dep_.CoverTab[1205]++
									*data = int16(order.Uint16(bs))
//line /usr/local/go/src/encoding/binary/binary.go:244
			// _ = "end of CoverTab[1205]"
		case *uint16:
//line /usr/local/go/src/encoding/binary/binary.go:245
			_go_fuzz_dep_.CoverTab[1206]++
									*data = order.Uint16(bs)
//line /usr/local/go/src/encoding/binary/binary.go:246
			// _ = "end of CoverTab[1206]"
		case *int32:
//line /usr/local/go/src/encoding/binary/binary.go:247
			_go_fuzz_dep_.CoverTab[1207]++
									*data = int32(order.Uint32(bs))
//line /usr/local/go/src/encoding/binary/binary.go:248
			// _ = "end of CoverTab[1207]"
		case *uint32:
//line /usr/local/go/src/encoding/binary/binary.go:249
			_go_fuzz_dep_.CoverTab[1208]++
									*data = order.Uint32(bs)
//line /usr/local/go/src/encoding/binary/binary.go:250
			// _ = "end of CoverTab[1208]"
		case *int64:
//line /usr/local/go/src/encoding/binary/binary.go:251
			_go_fuzz_dep_.CoverTab[1209]++
									*data = int64(order.Uint64(bs))
//line /usr/local/go/src/encoding/binary/binary.go:252
			// _ = "end of CoverTab[1209]"
		case *uint64:
//line /usr/local/go/src/encoding/binary/binary.go:253
			_go_fuzz_dep_.CoverTab[1210]++
									*data = order.Uint64(bs)
//line /usr/local/go/src/encoding/binary/binary.go:254
			// _ = "end of CoverTab[1210]"
		case *float32:
//line /usr/local/go/src/encoding/binary/binary.go:255
			_go_fuzz_dep_.CoverTab[1211]++
									*data = math.Float32frombits(order.Uint32(bs))
//line /usr/local/go/src/encoding/binary/binary.go:256
			// _ = "end of CoverTab[1211]"
		case *float64:
//line /usr/local/go/src/encoding/binary/binary.go:257
			_go_fuzz_dep_.CoverTab[1212]++
									*data = math.Float64frombits(order.Uint64(bs))
//line /usr/local/go/src/encoding/binary/binary.go:258
			// _ = "end of CoverTab[1212]"
		case []bool:
//line /usr/local/go/src/encoding/binary/binary.go:259
			_go_fuzz_dep_.CoverTab[1213]++
									for i, x := range bs {
//line /usr/local/go/src/encoding/binary/binary.go:260
				_go_fuzz_dep_.CoverTab[1225]++
										data[i] = x != 0
//line /usr/local/go/src/encoding/binary/binary.go:261
				// _ = "end of CoverTab[1225]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:262
			// _ = "end of CoverTab[1213]"
		case []int8:
//line /usr/local/go/src/encoding/binary/binary.go:263
			_go_fuzz_dep_.CoverTab[1214]++
									for i, x := range bs {
//line /usr/local/go/src/encoding/binary/binary.go:264
				_go_fuzz_dep_.CoverTab[1226]++
										data[i] = int8(x)
//line /usr/local/go/src/encoding/binary/binary.go:265
				// _ = "end of CoverTab[1226]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:266
			// _ = "end of CoverTab[1214]"
		case []uint8:
//line /usr/local/go/src/encoding/binary/binary.go:267
			_go_fuzz_dep_.CoverTab[1215]++
									copy(data, bs)
//line /usr/local/go/src/encoding/binary/binary.go:268
			// _ = "end of CoverTab[1215]"
		case []int16:
//line /usr/local/go/src/encoding/binary/binary.go:269
			_go_fuzz_dep_.CoverTab[1216]++
									for i := range data {
//line /usr/local/go/src/encoding/binary/binary.go:270
				_go_fuzz_dep_.CoverTab[1227]++
										data[i] = int16(order.Uint16(bs[2*i:]))
//line /usr/local/go/src/encoding/binary/binary.go:271
				// _ = "end of CoverTab[1227]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:272
			// _ = "end of CoverTab[1216]"
		case []uint16:
//line /usr/local/go/src/encoding/binary/binary.go:273
			_go_fuzz_dep_.CoverTab[1217]++
									for i := range data {
//line /usr/local/go/src/encoding/binary/binary.go:274
				_go_fuzz_dep_.CoverTab[1228]++
										data[i] = order.Uint16(bs[2*i:])
//line /usr/local/go/src/encoding/binary/binary.go:275
				// _ = "end of CoverTab[1228]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:276
			// _ = "end of CoverTab[1217]"
		case []int32:
//line /usr/local/go/src/encoding/binary/binary.go:277
			_go_fuzz_dep_.CoverTab[1218]++
									for i := range data {
//line /usr/local/go/src/encoding/binary/binary.go:278
				_go_fuzz_dep_.CoverTab[1229]++
										data[i] = int32(order.Uint32(bs[4*i:]))
//line /usr/local/go/src/encoding/binary/binary.go:279
				// _ = "end of CoverTab[1229]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:280
			// _ = "end of CoverTab[1218]"
		case []uint32:
//line /usr/local/go/src/encoding/binary/binary.go:281
			_go_fuzz_dep_.CoverTab[1219]++
									for i := range data {
//line /usr/local/go/src/encoding/binary/binary.go:282
				_go_fuzz_dep_.CoverTab[1230]++
										data[i] = order.Uint32(bs[4*i:])
//line /usr/local/go/src/encoding/binary/binary.go:283
				// _ = "end of CoverTab[1230]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:284
			// _ = "end of CoverTab[1219]"
		case []int64:
//line /usr/local/go/src/encoding/binary/binary.go:285
			_go_fuzz_dep_.CoverTab[1220]++
									for i := range data {
//line /usr/local/go/src/encoding/binary/binary.go:286
				_go_fuzz_dep_.CoverTab[1231]++
										data[i] = int64(order.Uint64(bs[8*i:]))
//line /usr/local/go/src/encoding/binary/binary.go:287
				// _ = "end of CoverTab[1231]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:288
			// _ = "end of CoverTab[1220]"
		case []uint64:
//line /usr/local/go/src/encoding/binary/binary.go:289
			_go_fuzz_dep_.CoverTab[1221]++
									for i := range data {
//line /usr/local/go/src/encoding/binary/binary.go:290
				_go_fuzz_dep_.CoverTab[1232]++
										data[i] = order.Uint64(bs[8*i:])
//line /usr/local/go/src/encoding/binary/binary.go:291
				// _ = "end of CoverTab[1232]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:292
			// _ = "end of CoverTab[1221]"
		case []float32:
//line /usr/local/go/src/encoding/binary/binary.go:293
			_go_fuzz_dep_.CoverTab[1222]++
									for i := range data {
//line /usr/local/go/src/encoding/binary/binary.go:294
				_go_fuzz_dep_.CoverTab[1233]++
										data[i] = math.Float32frombits(order.Uint32(bs[4*i:]))
//line /usr/local/go/src/encoding/binary/binary.go:295
				// _ = "end of CoverTab[1233]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:296
			// _ = "end of CoverTab[1222]"
		case []float64:
//line /usr/local/go/src/encoding/binary/binary.go:297
			_go_fuzz_dep_.CoverTab[1223]++
									for i := range data {
//line /usr/local/go/src/encoding/binary/binary.go:298
				_go_fuzz_dep_.CoverTab[1234]++
										data[i] = math.Float64frombits(order.Uint64(bs[8*i:]))
//line /usr/local/go/src/encoding/binary/binary.go:299
				// _ = "end of CoverTab[1234]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:300
			// _ = "end of CoverTab[1223]"
		default:
//line /usr/local/go/src/encoding/binary/binary.go:301
			_go_fuzz_dep_.CoverTab[1224]++
									n = 0
//line /usr/local/go/src/encoding/binary/binary.go:302
			// _ = "end of CoverTab[1224]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:303
		// _ = "end of CoverTab[1198]"
//line /usr/local/go/src/encoding/binary/binary.go:303
		_go_fuzz_dep_.CoverTab[1199]++
								if n != 0 {
//line /usr/local/go/src/encoding/binary/binary.go:304
			_go_fuzz_dep_.CoverTab[1235]++
									return nil
//line /usr/local/go/src/encoding/binary/binary.go:305
			// _ = "end of CoverTab[1235]"
		} else {
//line /usr/local/go/src/encoding/binary/binary.go:306
			_go_fuzz_dep_.CoverTab[1236]++
//line /usr/local/go/src/encoding/binary/binary.go:306
			// _ = "end of CoverTab[1236]"
//line /usr/local/go/src/encoding/binary/binary.go:306
		}
//line /usr/local/go/src/encoding/binary/binary.go:306
		// _ = "end of CoverTab[1199]"
	} else {
//line /usr/local/go/src/encoding/binary/binary.go:307
		_go_fuzz_dep_.CoverTab[1237]++
//line /usr/local/go/src/encoding/binary/binary.go:307
		// _ = "end of CoverTab[1237]"
//line /usr/local/go/src/encoding/binary/binary.go:307
	}
//line /usr/local/go/src/encoding/binary/binary.go:307
	// _ = "end of CoverTab[1192]"
//line /usr/local/go/src/encoding/binary/binary.go:307
	_go_fuzz_dep_.CoverTab[1193]++

//line /usr/local/go/src/encoding/binary/binary.go:310
	v := reflect.ValueOf(data)
	size := -1
	switch v.Kind() {
	case reflect.Pointer:
//line /usr/local/go/src/encoding/binary/binary.go:313
		_go_fuzz_dep_.CoverTab[1238]++
								v = v.Elem()
								size = dataSize(v)
//line /usr/local/go/src/encoding/binary/binary.go:315
		// _ = "end of CoverTab[1238]"
	case reflect.Slice:
//line /usr/local/go/src/encoding/binary/binary.go:316
		_go_fuzz_dep_.CoverTab[1239]++
								size = dataSize(v)
//line /usr/local/go/src/encoding/binary/binary.go:317
		// _ = "end of CoverTab[1239]"
//line /usr/local/go/src/encoding/binary/binary.go:317
	default:
//line /usr/local/go/src/encoding/binary/binary.go:317
		_go_fuzz_dep_.CoverTab[1240]++
//line /usr/local/go/src/encoding/binary/binary.go:317
		// _ = "end of CoverTab[1240]"
	}
//line /usr/local/go/src/encoding/binary/binary.go:318
	// _ = "end of CoverTab[1193]"
//line /usr/local/go/src/encoding/binary/binary.go:318
	_go_fuzz_dep_.CoverTab[1194]++
							if size < 0 {
//line /usr/local/go/src/encoding/binary/binary.go:319
		_go_fuzz_dep_.CoverTab[1241]++
								return errors.New("binary.Read: invalid type " + reflect.TypeOf(data).String())
//line /usr/local/go/src/encoding/binary/binary.go:320
		// _ = "end of CoverTab[1241]"
	} else {
//line /usr/local/go/src/encoding/binary/binary.go:321
		_go_fuzz_dep_.CoverTab[1242]++
//line /usr/local/go/src/encoding/binary/binary.go:321
		// _ = "end of CoverTab[1242]"
//line /usr/local/go/src/encoding/binary/binary.go:321
	}
//line /usr/local/go/src/encoding/binary/binary.go:321
	// _ = "end of CoverTab[1194]"
//line /usr/local/go/src/encoding/binary/binary.go:321
	_go_fuzz_dep_.CoverTab[1195]++
							d := &decoder{order: order, buf: make([]byte, size)}
							if _, err := io.ReadFull(r, d.buf); err != nil {
//line /usr/local/go/src/encoding/binary/binary.go:323
		_go_fuzz_dep_.CoverTab[1243]++
								return err
//line /usr/local/go/src/encoding/binary/binary.go:324
		// _ = "end of CoverTab[1243]"
	} else {
//line /usr/local/go/src/encoding/binary/binary.go:325
		_go_fuzz_dep_.CoverTab[1244]++
//line /usr/local/go/src/encoding/binary/binary.go:325
		// _ = "end of CoverTab[1244]"
//line /usr/local/go/src/encoding/binary/binary.go:325
	}
//line /usr/local/go/src/encoding/binary/binary.go:325
	// _ = "end of CoverTab[1195]"
//line /usr/local/go/src/encoding/binary/binary.go:325
	_go_fuzz_dep_.CoverTab[1196]++
							d.value(v)
							return nil
//line /usr/local/go/src/encoding/binary/binary.go:327
	// _ = "end of CoverTab[1196]"
}

// Write writes the binary representation of data into w.
//line /usr/local/go/src/encoding/binary/binary.go:330
// Data must be a fixed-size value or a slice of fixed-size
//line /usr/local/go/src/encoding/binary/binary.go:330
// values, or a pointer to such data.
//line /usr/local/go/src/encoding/binary/binary.go:330
// Boolean values encode as one byte: 1 for true, and 0 for false.
//line /usr/local/go/src/encoding/binary/binary.go:330
// Bytes written to w are encoded using the specified byte order
//line /usr/local/go/src/encoding/binary/binary.go:330
// and read from successive fields of the data.
//line /usr/local/go/src/encoding/binary/binary.go:330
// When writing structs, zero values are written for fields
//line /usr/local/go/src/encoding/binary/binary.go:330
// with blank (_) field names.
//line /usr/local/go/src/encoding/binary/binary.go:338
func Write(w io.Writer, order ByteOrder, data any) error {
//line /usr/local/go/src/encoding/binary/binary.go:338
	_go_fuzz_dep_.CoverTab[1245]++

							if n := intDataSize(data); n != 0 {
//line /usr/local/go/src/encoding/binary/binary.go:340
		_go_fuzz_dep_.CoverTab[1248]++
								bs := make([]byte, n)
								switch v := data.(type) {
		case *bool:
//line /usr/local/go/src/encoding/binary/binary.go:343
			_go_fuzz_dep_.CoverTab[1250]++
									if *v {
//line /usr/local/go/src/encoding/binary/binary.go:344
				_go_fuzz_dep_.CoverTab[1283]++
										bs[0] = 1
//line /usr/local/go/src/encoding/binary/binary.go:345
				// _ = "end of CoverTab[1283]"
			} else {
//line /usr/local/go/src/encoding/binary/binary.go:346
				_go_fuzz_dep_.CoverTab[1284]++
										bs[0] = 0
//line /usr/local/go/src/encoding/binary/binary.go:347
				// _ = "end of CoverTab[1284]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:348
			// _ = "end of CoverTab[1250]"
		case bool:
//line /usr/local/go/src/encoding/binary/binary.go:349
			_go_fuzz_dep_.CoverTab[1251]++
									if v {
//line /usr/local/go/src/encoding/binary/binary.go:350
				_go_fuzz_dep_.CoverTab[1285]++
										bs[0] = 1
//line /usr/local/go/src/encoding/binary/binary.go:351
				// _ = "end of CoverTab[1285]"
			} else {
//line /usr/local/go/src/encoding/binary/binary.go:352
				_go_fuzz_dep_.CoverTab[1286]++
										bs[0] = 0
//line /usr/local/go/src/encoding/binary/binary.go:353
				// _ = "end of CoverTab[1286]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:354
			// _ = "end of CoverTab[1251]"
		case []bool:
//line /usr/local/go/src/encoding/binary/binary.go:355
			_go_fuzz_dep_.CoverTab[1252]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:356
				_go_fuzz_dep_.CoverTab[1287]++
										if x {
//line /usr/local/go/src/encoding/binary/binary.go:357
					_go_fuzz_dep_.CoverTab[1288]++
											bs[i] = 1
//line /usr/local/go/src/encoding/binary/binary.go:358
					// _ = "end of CoverTab[1288]"
				} else {
//line /usr/local/go/src/encoding/binary/binary.go:359
					_go_fuzz_dep_.CoverTab[1289]++
											bs[i] = 0
//line /usr/local/go/src/encoding/binary/binary.go:360
					// _ = "end of CoverTab[1289]"
				}
//line /usr/local/go/src/encoding/binary/binary.go:361
				// _ = "end of CoverTab[1287]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:362
			// _ = "end of CoverTab[1252]"
		case *int8:
//line /usr/local/go/src/encoding/binary/binary.go:363
			_go_fuzz_dep_.CoverTab[1253]++
									bs[0] = byte(*v)
//line /usr/local/go/src/encoding/binary/binary.go:364
			// _ = "end of CoverTab[1253]"
		case int8:
//line /usr/local/go/src/encoding/binary/binary.go:365
			_go_fuzz_dep_.CoverTab[1254]++
									bs[0] = byte(v)
//line /usr/local/go/src/encoding/binary/binary.go:366
			// _ = "end of CoverTab[1254]"
		case []int8:
//line /usr/local/go/src/encoding/binary/binary.go:367
			_go_fuzz_dep_.CoverTab[1255]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:368
				_go_fuzz_dep_.CoverTab[1290]++
										bs[i] = byte(x)
//line /usr/local/go/src/encoding/binary/binary.go:369
				// _ = "end of CoverTab[1290]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:370
			// _ = "end of CoverTab[1255]"
		case *uint8:
//line /usr/local/go/src/encoding/binary/binary.go:371
			_go_fuzz_dep_.CoverTab[1256]++
									bs[0] = *v
//line /usr/local/go/src/encoding/binary/binary.go:372
			// _ = "end of CoverTab[1256]"
		case uint8:
//line /usr/local/go/src/encoding/binary/binary.go:373
			_go_fuzz_dep_.CoverTab[1257]++
									bs[0] = v
//line /usr/local/go/src/encoding/binary/binary.go:374
			// _ = "end of CoverTab[1257]"
		case []uint8:
//line /usr/local/go/src/encoding/binary/binary.go:375
			_go_fuzz_dep_.CoverTab[1258]++
									bs = v
//line /usr/local/go/src/encoding/binary/binary.go:376
			// _ = "end of CoverTab[1258]"
		case *int16:
//line /usr/local/go/src/encoding/binary/binary.go:377
			_go_fuzz_dep_.CoverTab[1259]++
									order.PutUint16(bs, uint16(*v))
//line /usr/local/go/src/encoding/binary/binary.go:378
			// _ = "end of CoverTab[1259]"
		case int16:
//line /usr/local/go/src/encoding/binary/binary.go:379
			_go_fuzz_dep_.CoverTab[1260]++
									order.PutUint16(bs, uint16(v))
//line /usr/local/go/src/encoding/binary/binary.go:380
			// _ = "end of CoverTab[1260]"
		case []int16:
//line /usr/local/go/src/encoding/binary/binary.go:381
			_go_fuzz_dep_.CoverTab[1261]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:382
				_go_fuzz_dep_.CoverTab[1291]++
										order.PutUint16(bs[2*i:], uint16(x))
//line /usr/local/go/src/encoding/binary/binary.go:383
				// _ = "end of CoverTab[1291]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:384
			// _ = "end of CoverTab[1261]"
		case *uint16:
//line /usr/local/go/src/encoding/binary/binary.go:385
			_go_fuzz_dep_.CoverTab[1262]++
									order.PutUint16(bs, *v)
//line /usr/local/go/src/encoding/binary/binary.go:386
			// _ = "end of CoverTab[1262]"
		case uint16:
//line /usr/local/go/src/encoding/binary/binary.go:387
			_go_fuzz_dep_.CoverTab[1263]++
									order.PutUint16(bs, v)
//line /usr/local/go/src/encoding/binary/binary.go:388
			// _ = "end of CoverTab[1263]"
		case []uint16:
//line /usr/local/go/src/encoding/binary/binary.go:389
			_go_fuzz_dep_.CoverTab[1264]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:390
				_go_fuzz_dep_.CoverTab[1292]++
										order.PutUint16(bs[2*i:], x)
//line /usr/local/go/src/encoding/binary/binary.go:391
				// _ = "end of CoverTab[1292]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:392
			// _ = "end of CoverTab[1264]"
		case *int32:
//line /usr/local/go/src/encoding/binary/binary.go:393
			_go_fuzz_dep_.CoverTab[1265]++
									order.PutUint32(bs, uint32(*v))
//line /usr/local/go/src/encoding/binary/binary.go:394
			// _ = "end of CoverTab[1265]"
		case int32:
//line /usr/local/go/src/encoding/binary/binary.go:395
			_go_fuzz_dep_.CoverTab[1266]++
									order.PutUint32(bs, uint32(v))
//line /usr/local/go/src/encoding/binary/binary.go:396
			// _ = "end of CoverTab[1266]"
		case []int32:
//line /usr/local/go/src/encoding/binary/binary.go:397
			_go_fuzz_dep_.CoverTab[1267]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:398
				_go_fuzz_dep_.CoverTab[1293]++
										order.PutUint32(bs[4*i:], uint32(x))
//line /usr/local/go/src/encoding/binary/binary.go:399
				// _ = "end of CoverTab[1293]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:400
			// _ = "end of CoverTab[1267]"
		case *uint32:
//line /usr/local/go/src/encoding/binary/binary.go:401
			_go_fuzz_dep_.CoverTab[1268]++
									order.PutUint32(bs, *v)
//line /usr/local/go/src/encoding/binary/binary.go:402
			// _ = "end of CoverTab[1268]"
		case uint32:
//line /usr/local/go/src/encoding/binary/binary.go:403
			_go_fuzz_dep_.CoverTab[1269]++
									order.PutUint32(bs, v)
//line /usr/local/go/src/encoding/binary/binary.go:404
			// _ = "end of CoverTab[1269]"
		case []uint32:
//line /usr/local/go/src/encoding/binary/binary.go:405
			_go_fuzz_dep_.CoverTab[1270]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:406
				_go_fuzz_dep_.CoverTab[1294]++
										order.PutUint32(bs[4*i:], x)
//line /usr/local/go/src/encoding/binary/binary.go:407
				// _ = "end of CoverTab[1294]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:408
			// _ = "end of CoverTab[1270]"
		case *int64:
//line /usr/local/go/src/encoding/binary/binary.go:409
			_go_fuzz_dep_.CoverTab[1271]++
									order.PutUint64(bs, uint64(*v))
//line /usr/local/go/src/encoding/binary/binary.go:410
			// _ = "end of CoverTab[1271]"
		case int64:
//line /usr/local/go/src/encoding/binary/binary.go:411
			_go_fuzz_dep_.CoverTab[1272]++
									order.PutUint64(bs, uint64(v))
//line /usr/local/go/src/encoding/binary/binary.go:412
			// _ = "end of CoverTab[1272]"
		case []int64:
//line /usr/local/go/src/encoding/binary/binary.go:413
			_go_fuzz_dep_.CoverTab[1273]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:414
				_go_fuzz_dep_.CoverTab[1295]++
										order.PutUint64(bs[8*i:], uint64(x))
//line /usr/local/go/src/encoding/binary/binary.go:415
				// _ = "end of CoverTab[1295]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:416
			// _ = "end of CoverTab[1273]"
		case *uint64:
//line /usr/local/go/src/encoding/binary/binary.go:417
			_go_fuzz_dep_.CoverTab[1274]++
									order.PutUint64(bs, *v)
//line /usr/local/go/src/encoding/binary/binary.go:418
			// _ = "end of CoverTab[1274]"
		case uint64:
//line /usr/local/go/src/encoding/binary/binary.go:419
			_go_fuzz_dep_.CoverTab[1275]++
									order.PutUint64(bs, v)
//line /usr/local/go/src/encoding/binary/binary.go:420
			// _ = "end of CoverTab[1275]"
		case []uint64:
//line /usr/local/go/src/encoding/binary/binary.go:421
			_go_fuzz_dep_.CoverTab[1276]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:422
				_go_fuzz_dep_.CoverTab[1296]++
										order.PutUint64(bs[8*i:], x)
//line /usr/local/go/src/encoding/binary/binary.go:423
				// _ = "end of CoverTab[1296]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:424
			// _ = "end of CoverTab[1276]"
		case *float32:
//line /usr/local/go/src/encoding/binary/binary.go:425
			_go_fuzz_dep_.CoverTab[1277]++
									order.PutUint32(bs, math.Float32bits(*v))
//line /usr/local/go/src/encoding/binary/binary.go:426
			// _ = "end of CoverTab[1277]"
		case float32:
//line /usr/local/go/src/encoding/binary/binary.go:427
			_go_fuzz_dep_.CoverTab[1278]++
									order.PutUint32(bs, math.Float32bits(v))
//line /usr/local/go/src/encoding/binary/binary.go:428
			// _ = "end of CoverTab[1278]"
		case []float32:
//line /usr/local/go/src/encoding/binary/binary.go:429
			_go_fuzz_dep_.CoverTab[1279]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:430
				_go_fuzz_dep_.CoverTab[1297]++
										order.PutUint32(bs[4*i:], math.Float32bits(x))
//line /usr/local/go/src/encoding/binary/binary.go:431
				// _ = "end of CoverTab[1297]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:432
			// _ = "end of CoverTab[1279]"
		case *float64:
//line /usr/local/go/src/encoding/binary/binary.go:433
			_go_fuzz_dep_.CoverTab[1280]++
									order.PutUint64(bs, math.Float64bits(*v))
//line /usr/local/go/src/encoding/binary/binary.go:434
			// _ = "end of CoverTab[1280]"
		case float64:
//line /usr/local/go/src/encoding/binary/binary.go:435
			_go_fuzz_dep_.CoverTab[1281]++
									order.PutUint64(bs, math.Float64bits(v))
//line /usr/local/go/src/encoding/binary/binary.go:436
			// _ = "end of CoverTab[1281]"
		case []float64:
//line /usr/local/go/src/encoding/binary/binary.go:437
			_go_fuzz_dep_.CoverTab[1282]++
									for i, x := range v {
//line /usr/local/go/src/encoding/binary/binary.go:438
				_go_fuzz_dep_.CoverTab[1298]++
										order.PutUint64(bs[8*i:], math.Float64bits(x))
//line /usr/local/go/src/encoding/binary/binary.go:439
				// _ = "end of CoverTab[1298]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:440
			// _ = "end of CoverTab[1282]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:441
		// _ = "end of CoverTab[1248]"
//line /usr/local/go/src/encoding/binary/binary.go:441
		_go_fuzz_dep_.CoverTab[1249]++
								_, err := w.Write(bs)
								return err
//line /usr/local/go/src/encoding/binary/binary.go:443
		// _ = "end of CoverTab[1249]"
	} else {
//line /usr/local/go/src/encoding/binary/binary.go:444
		_go_fuzz_dep_.CoverTab[1299]++
//line /usr/local/go/src/encoding/binary/binary.go:444
		// _ = "end of CoverTab[1299]"
//line /usr/local/go/src/encoding/binary/binary.go:444
	}
//line /usr/local/go/src/encoding/binary/binary.go:444
	// _ = "end of CoverTab[1245]"
//line /usr/local/go/src/encoding/binary/binary.go:444
	_go_fuzz_dep_.CoverTab[1246]++

//line /usr/local/go/src/encoding/binary/binary.go:447
	v := reflect.Indirect(reflect.ValueOf(data))
	size := dataSize(v)
	if size < 0 {
//line /usr/local/go/src/encoding/binary/binary.go:449
		_go_fuzz_dep_.CoverTab[1300]++
								return errors.New("binary.Write: invalid type " + reflect.TypeOf(data).String())
//line /usr/local/go/src/encoding/binary/binary.go:450
		// _ = "end of CoverTab[1300]"
	} else {
//line /usr/local/go/src/encoding/binary/binary.go:451
		_go_fuzz_dep_.CoverTab[1301]++
//line /usr/local/go/src/encoding/binary/binary.go:451
		// _ = "end of CoverTab[1301]"
//line /usr/local/go/src/encoding/binary/binary.go:451
	}
//line /usr/local/go/src/encoding/binary/binary.go:451
	// _ = "end of CoverTab[1246]"
//line /usr/local/go/src/encoding/binary/binary.go:451
	_go_fuzz_dep_.CoverTab[1247]++
							buf := make([]byte, size)
							e := &encoder{order: order, buf: buf}
							e.value(v)
							_, err := w.Write(buf)
							return err
//line /usr/local/go/src/encoding/binary/binary.go:456
	// _ = "end of CoverTab[1247]"
}

// Size returns how many bytes Write would generate to encode the value v, which
//line /usr/local/go/src/encoding/binary/binary.go:459
// must be a fixed-size value or a slice of fixed-size values, or a pointer to such data.
//line /usr/local/go/src/encoding/binary/binary.go:459
// If v is neither of these, Size returns -1.
//line /usr/local/go/src/encoding/binary/binary.go:462
func Size(v any) int {
//line /usr/local/go/src/encoding/binary/binary.go:462
	_go_fuzz_dep_.CoverTab[1302]++
							return dataSize(reflect.Indirect(reflect.ValueOf(v)))
//line /usr/local/go/src/encoding/binary/binary.go:463
	// _ = "end of CoverTab[1302]"
}

var structSize sync.Map	// map[reflect.Type]int

// dataSize returns the number of bytes the actual data represented by v occupies in memory.
//line /usr/local/go/src/encoding/binary/binary.go:468
// For compound structures, it sums the sizes of the elements. Thus, for instance, for a slice
//line /usr/local/go/src/encoding/binary/binary.go:468
// it returns the length of the slice times the element size and does not count the memory
//line /usr/local/go/src/encoding/binary/binary.go:468
// occupied by the header. If the type of v is not acceptable, dataSize returns -1.
//line /usr/local/go/src/encoding/binary/binary.go:472
func dataSize(v reflect.Value) int {
//line /usr/local/go/src/encoding/binary/binary.go:472
	_go_fuzz_dep_.CoverTab[1303]++
							switch v.Kind() {
	case reflect.Slice:
//line /usr/local/go/src/encoding/binary/binary.go:474
		_go_fuzz_dep_.CoverTab[1304]++
								if s := sizeof(v.Type().Elem()); s >= 0 {
//line /usr/local/go/src/encoding/binary/binary.go:475
			_go_fuzz_dep_.CoverTab[1309]++
									return s * v.Len()
//line /usr/local/go/src/encoding/binary/binary.go:476
			// _ = "end of CoverTab[1309]"
		} else {
//line /usr/local/go/src/encoding/binary/binary.go:477
			_go_fuzz_dep_.CoverTab[1310]++
//line /usr/local/go/src/encoding/binary/binary.go:477
			// _ = "end of CoverTab[1310]"
//line /usr/local/go/src/encoding/binary/binary.go:477
		}
//line /usr/local/go/src/encoding/binary/binary.go:477
		// _ = "end of CoverTab[1304]"
//line /usr/local/go/src/encoding/binary/binary.go:477
		_go_fuzz_dep_.CoverTab[1305]++
								return -1
//line /usr/local/go/src/encoding/binary/binary.go:478
		// _ = "end of CoverTab[1305]"

	case reflect.Struct:
//line /usr/local/go/src/encoding/binary/binary.go:480
		_go_fuzz_dep_.CoverTab[1306]++
								t := v.Type()
								if size, ok := structSize.Load(t); ok {
//line /usr/local/go/src/encoding/binary/binary.go:482
			_go_fuzz_dep_.CoverTab[1311]++
									return size.(int)
//line /usr/local/go/src/encoding/binary/binary.go:483
			// _ = "end of CoverTab[1311]"
		} else {
//line /usr/local/go/src/encoding/binary/binary.go:484
			_go_fuzz_dep_.CoverTab[1312]++
//line /usr/local/go/src/encoding/binary/binary.go:484
			// _ = "end of CoverTab[1312]"
//line /usr/local/go/src/encoding/binary/binary.go:484
		}
//line /usr/local/go/src/encoding/binary/binary.go:484
		// _ = "end of CoverTab[1306]"
//line /usr/local/go/src/encoding/binary/binary.go:484
		_go_fuzz_dep_.CoverTab[1307]++
								size := sizeof(t)
								structSize.Store(t, size)
								return size
//line /usr/local/go/src/encoding/binary/binary.go:487
		// _ = "end of CoverTab[1307]"

	default:
//line /usr/local/go/src/encoding/binary/binary.go:489
		_go_fuzz_dep_.CoverTab[1308]++
								return sizeof(v.Type())
//line /usr/local/go/src/encoding/binary/binary.go:490
		// _ = "end of CoverTab[1308]"
	}
//line /usr/local/go/src/encoding/binary/binary.go:491
	// _ = "end of CoverTab[1303]"
}

// sizeof returns the size >= 0 of variables for the given type or -1 if the type is not acceptable.
func sizeof(t reflect.Type) int {
//line /usr/local/go/src/encoding/binary/binary.go:495
	_go_fuzz_dep_.CoverTab[1313]++
							switch t.Kind() {
	case reflect.Array:
//line /usr/local/go/src/encoding/binary/binary.go:497
		_go_fuzz_dep_.CoverTab[1315]++
								if s := sizeof(t.Elem()); s >= 0 {
//line /usr/local/go/src/encoding/binary/binary.go:498
			_go_fuzz_dep_.CoverTab[1320]++
									return s * t.Len()
//line /usr/local/go/src/encoding/binary/binary.go:499
			// _ = "end of CoverTab[1320]"
		} else {
//line /usr/local/go/src/encoding/binary/binary.go:500
			_go_fuzz_dep_.CoverTab[1321]++
//line /usr/local/go/src/encoding/binary/binary.go:500
			// _ = "end of CoverTab[1321]"
//line /usr/local/go/src/encoding/binary/binary.go:500
		}
//line /usr/local/go/src/encoding/binary/binary.go:500
		// _ = "end of CoverTab[1315]"

	case reflect.Struct:
//line /usr/local/go/src/encoding/binary/binary.go:502
		_go_fuzz_dep_.CoverTab[1316]++
								sum := 0
								for i, n := 0, t.NumField(); i < n; i++ {
//line /usr/local/go/src/encoding/binary/binary.go:504
			_go_fuzz_dep_.CoverTab[1322]++
									s := sizeof(t.Field(i).Type)
									if s < 0 {
//line /usr/local/go/src/encoding/binary/binary.go:506
				_go_fuzz_dep_.CoverTab[1324]++
										return -1
//line /usr/local/go/src/encoding/binary/binary.go:507
				// _ = "end of CoverTab[1324]"
			} else {
//line /usr/local/go/src/encoding/binary/binary.go:508
				_go_fuzz_dep_.CoverTab[1325]++
//line /usr/local/go/src/encoding/binary/binary.go:508
				// _ = "end of CoverTab[1325]"
//line /usr/local/go/src/encoding/binary/binary.go:508
			}
//line /usr/local/go/src/encoding/binary/binary.go:508
			// _ = "end of CoverTab[1322]"
//line /usr/local/go/src/encoding/binary/binary.go:508
			_go_fuzz_dep_.CoverTab[1323]++
									sum += s
//line /usr/local/go/src/encoding/binary/binary.go:509
			// _ = "end of CoverTab[1323]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:510
		// _ = "end of CoverTab[1316]"
//line /usr/local/go/src/encoding/binary/binary.go:510
		_go_fuzz_dep_.CoverTab[1317]++
								return sum
//line /usr/local/go/src/encoding/binary/binary.go:511
		// _ = "end of CoverTab[1317]"

	case reflect.Bool,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
//line /usr/local/go/src/encoding/binary/binary.go:516
		_go_fuzz_dep_.CoverTab[1318]++
								return int(t.Size())
//line /usr/local/go/src/encoding/binary/binary.go:517
		// _ = "end of CoverTab[1318]"
//line /usr/local/go/src/encoding/binary/binary.go:517
	default:
//line /usr/local/go/src/encoding/binary/binary.go:517
		_go_fuzz_dep_.CoverTab[1319]++
//line /usr/local/go/src/encoding/binary/binary.go:517
		// _ = "end of CoverTab[1319]"
	}
//line /usr/local/go/src/encoding/binary/binary.go:518
	// _ = "end of CoverTab[1313]"
//line /usr/local/go/src/encoding/binary/binary.go:518
	_go_fuzz_dep_.CoverTab[1314]++

							return -1
//line /usr/local/go/src/encoding/binary/binary.go:520
	// _ = "end of CoverTab[1314]"
}

type coder struct {
	order	ByteOrder
	buf	[]byte
	offset	int
}

type decoder coder
type encoder coder

func (d *decoder) bool() bool {
//line /usr/local/go/src/encoding/binary/binary.go:532
	_go_fuzz_dep_.CoverTab[1326]++
							x := d.buf[d.offset]
							d.offset++
							return x != 0
//line /usr/local/go/src/encoding/binary/binary.go:535
	// _ = "end of CoverTab[1326]"
}

func (e *encoder) bool(x bool) {
//line /usr/local/go/src/encoding/binary/binary.go:538
	_go_fuzz_dep_.CoverTab[1327]++
							if x {
//line /usr/local/go/src/encoding/binary/binary.go:539
		_go_fuzz_dep_.CoverTab[1329]++
								e.buf[e.offset] = 1
//line /usr/local/go/src/encoding/binary/binary.go:540
		// _ = "end of CoverTab[1329]"
	} else {
//line /usr/local/go/src/encoding/binary/binary.go:541
		_go_fuzz_dep_.CoverTab[1330]++
								e.buf[e.offset] = 0
//line /usr/local/go/src/encoding/binary/binary.go:542
		// _ = "end of CoverTab[1330]"
	}
//line /usr/local/go/src/encoding/binary/binary.go:543
	// _ = "end of CoverTab[1327]"
//line /usr/local/go/src/encoding/binary/binary.go:543
	_go_fuzz_dep_.CoverTab[1328]++
							e.offset++
//line /usr/local/go/src/encoding/binary/binary.go:544
	// _ = "end of CoverTab[1328]"
}

func (d *decoder) uint8() uint8 {
//line /usr/local/go/src/encoding/binary/binary.go:547
	_go_fuzz_dep_.CoverTab[1331]++
							x := d.buf[d.offset]
							d.offset++
							return x
//line /usr/local/go/src/encoding/binary/binary.go:550
	// _ = "end of CoverTab[1331]"
}

func (e *encoder) uint8(x uint8) {
//line /usr/local/go/src/encoding/binary/binary.go:553
	_go_fuzz_dep_.CoverTab[1332]++
							e.buf[e.offset] = x
							e.offset++
//line /usr/local/go/src/encoding/binary/binary.go:555
	// _ = "end of CoverTab[1332]"
}

func (d *decoder) uint16() uint16 {
//line /usr/local/go/src/encoding/binary/binary.go:558
	_go_fuzz_dep_.CoverTab[1333]++
							x := d.order.Uint16(d.buf[d.offset : d.offset+2])
							d.offset += 2
							return x
//line /usr/local/go/src/encoding/binary/binary.go:561
	// _ = "end of CoverTab[1333]"
}

func (e *encoder) uint16(x uint16) {
//line /usr/local/go/src/encoding/binary/binary.go:564
	_go_fuzz_dep_.CoverTab[1334]++
							e.order.PutUint16(e.buf[e.offset:e.offset+2], x)
							e.offset += 2
//line /usr/local/go/src/encoding/binary/binary.go:566
	// _ = "end of CoverTab[1334]"
}

func (d *decoder) uint32() uint32 {
//line /usr/local/go/src/encoding/binary/binary.go:569
	_go_fuzz_dep_.CoverTab[1335]++
							x := d.order.Uint32(d.buf[d.offset : d.offset+4])
							d.offset += 4
							return x
//line /usr/local/go/src/encoding/binary/binary.go:572
	// _ = "end of CoverTab[1335]"
}

func (e *encoder) uint32(x uint32) {
//line /usr/local/go/src/encoding/binary/binary.go:575
	_go_fuzz_dep_.CoverTab[1336]++
							e.order.PutUint32(e.buf[e.offset:e.offset+4], x)
							e.offset += 4
//line /usr/local/go/src/encoding/binary/binary.go:577
	// _ = "end of CoverTab[1336]"
}

func (d *decoder) uint64() uint64 {
//line /usr/local/go/src/encoding/binary/binary.go:580
	_go_fuzz_dep_.CoverTab[1337]++
							x := d.order.Uint64(d.buf[d.offset : d.offset+8])
							d.offset += 8
							return x
//line /usr/local/go/src/encoding/binary/binary.go:583
	// _ = "end of CoverTab[1337]"
}

func (e *encoder) uint64(x uint64) {
//line /usr/local/go/src/encoding/binary/binary.go:586
	_go_fuzz_dep_.CoverTab[1338]++
							e.order.PutUint64(e.buf[e.offset:e.offset+8], x)
							e.offset += 8
//line /usr/local/go/src/encoding/binary/binary.go:588
	// _ = "end of CoverTab[1338]"
}

func (d *decoder) int8() int8 {
//line /usr/local/go/src/encoding/binary/binary.go:591
	_go_fuzz_dep_.CoverTab[1339]++
//line /usr/local/go/src/encoding/binary/binary.go:591
	return int8(d.uint8())
//line /usr/local/go/src/encoding/binary/binary.go:591
	// _ = "end of CoverTab[1339]"
//line /usr/local/go/src/encoding/binary/binary.go:591
}

func (e *encoder) int8(x int8) {
//line /usr/local/go/src/encoding/binary/binary.go:593
	_go_fuzz_dep_.CoverTab[1340]++
//line /usr/local/go/src/encoding/binary/binary.go:593
	e.uint8(uint8(x))
//line /usr/local/go/src/encoding/binary/binary.go:593
	// _ = "end of CoverTab[1340]"
//line /usr/local/go/src/encoding/binary/binary.go:593
}

func (d *decoder) int16() int16 {
//line /usr/local/go/src/encoding/binary/binary.go:595
	_go_fuzz_dep_.CoverTab[1341]++
//line /usr/local/go/src/encoding/binary/binary.go:595
	return int16(d.uint16())
//line /usr/local/go/src/encoding/binary/binary.go:595
	// _ = "end of CoverTab[1341]"
//line /usr/local/go/src/encoding/binary/binary.go:595
}

func (e *encoder) int16(x int16) {
//line /usr/local/go/src/encoding/binary/binary.go:597
	_go_fuzz_dep_.CoverTab[1342]++
//line /usr/local/go/src/encoding/binary/binary.go:597
	e.uint16(uint16(x))
//line /usr/local/go/src/encoding/binary/binary.go:597
	// _ = "end of CoverTab[1342]"
//line /usr/local/go/src/encoding/binary/binary.go:597
}

func (d *decoder) int32() int32 {
//line /usr/local/go/src/encoding/binary/binary.go:599
	_go_fuzz_dep_.CoverTab[1343]++
//line /usr/local/go/src/encoding/binary/binary.go:599
	return int32(d.uint32())
//line /usr/local/go/src/encoding/binary/binary.go:599
	// _ = "end of CoverTab[1343]"
//line /usr/local/go/src/encoding/binary/binary.go:599
}

func (e *encoder) int32(x int32) {
//line /usr/local/go/src/encoding/binary/binary.go:601
	_go_fuzz_dep_.CoverTab[1344]++
//line /usr/local/go/src/encoding/binary/binary.go:601
	e.uint32(uint32(x))
//line /usr/local/go/src/encoding/binary/binary.go:601
	// _ = "end of CoverTab[1344]"
//line /usr/local/go/src/encoding/binary/binary.go:601
}

func (d *decoder) int64() int64 {
//line /usr/local/go/src/encoding/binary/binary.go:603
	_go_fuzz_dep_.CoverTab[1345]++
//line /usr/local/go/src/encoding/binary/binary.go:603
	return int64(d.uint64())
//line /usr/local/go/src/encoding/binary/binary.go:603
	// _ = "end of CoverTab[1345]"
//line /usr/local/go/src/encoding/binary/binary.go:603
}

func (e *encoder) int64(x int64) {
//line /usr/local/go/src/encoding/binary/binary.go:605
	_go_fuzz_dep_.CoverTab[1346]++
//line /usr/local/go/src/encoding/binary/binary.go:605
	e.uint64(uint64(x))
//line /usr/local/go/src/encoding/binary/binary.go:605
	// _ = "end of CoverTab[1346]"
//line /usr/local/go/src/encoding/binary/binary.go:605
}

func (d *decoder) value(v reflect.Value) {
//line /usr/local/go/src/encoding/binary/binary.go:607
	_go_fuzz_dep_.CoverTab[1347]++
							switch v.Kind() {
	case reflect.Array:
//line /usr/local/go/src/encoding/binary/binary.go:609
		_go_fuzz_dep_.CoverTab[1348]++
								l := v.Len()
								for i := 0; i < l; i++ {
//line /usr/local/go/src/encoding/binary/binary.go:611
			_go_fuzz_dep_.CoverTab[1365]++
									d.value(v.Index(i))
//line /usr/local/go/src/encoding/binary/binary.go:612
			// _ = "end of CoverTab[1365]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:613
		// _ = "end of CoverTab[1348]"

	case reflect.Struct:
//line /usr/local/go/src/encoding/binary/binary.go:615
		_go_fuzz_dep_.CoverTab[1349]++
								t := v.Type()
								l := v.NumField()
								for i := 0; i < l; i++ {
//line /usr/local/go/src/encoding/binary/binary.go:618
			_go_fuzz_dep_.CoverTab[1366]++

//line /usr/local/go/src/encoding/binary/binary.go:624
			if v := v.Field(i); v.CanSet() || func() bool {
//line /usr/local/go/src/encoding/binary/binary.go:624
				_go_fuzz_dep_.CoverTab[1367]++
//line /usr/local/go/src/encoding/binary/binary.go:624
				return t.Field(i).Name != "_"
//line /usr/local/go/src/encoding/binary/binary.go:624
				// _ = "end of CoverTab[1367]"
//line /usr/local/go/src/encoding/binary/binary.go:624
			}() {
//line /usr/local/go/src/encoding/binary/binary.go:624
				_go_fuzz_dep_.CoverTab[1368]++
										d.value(v)
//line /usr/local/go/src/encoding/binary/binary.go:625
				// _ = "end of CoverTab[1368]"
			} else {
//line /usr/local/go/src/encoding/binary/binary.go:626
				_go_fuzz_dep_.CoverTab[1369]++
										d.skip(v)
//line /usr/local/go/src/encoding/binary/binary.go:627
				// _ = "end of CoverTab[1369]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:628
			// _ = "end of CoverTab[1366]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:629
		// _ = "end of CoverTab[1349]"

	case reflect.Slice:
//line /usr/local/go/src/encoding/binary/binary.go:631
		_go_fuzz_dep_.CoverTab[1350]++
								l := v.Len()
								for i := 0; i < l; i++ {
//line /usr/local/go/src/encoding/binary/binary.go:633
			_go_fuzz_dep_.CoverTab[1370]++
									d.value(v.Index(i))
//line /usr/local/go/src/encoding/binary/binary.go:634
			// _ = "end of CoverTab[1370]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:635
		// _ = "end of CoverTab[1350]"

	case reflect.Bool:
//line /usr/local/go/src/encoding/binary/binary.go:637
		_go_fuzz_dep_.CoverTab[1351]++
								v.SetBool(d.bool())
//line /usr/local/go/src/encoding/binary/binary.go:638
		// _ = "end of CoverTab[1351]"

	case reflect.Int8:
//line /usr/local/go/src/encoding/binary/binary.go:640
		_go_fuzz_dep_.CoverTab[1352]++
								v.SetInt(int64(d.int8()))
//line /usr/local/go/src/encoding/binary/binary.go:641
		// _ = "end of CoverTab[1352]"
	case reflect.Int16:
//line /usr/local/go/src/encoding/binary/binary.go:642
		_go_fuzz_dep_.CoverTab[1353]++
								v.SetInt(int64(d.int16()))
//line /usr/local/go/src/encoding/binary/binary.go:643
		// _ = "end of CoverTab[1353]"
	case reflect.Int32:
//line /usr/local/go/src/encoding/binary/binary.go:644
		_go_fuzz_dep_.CoverTab[1354]++
								v.SetInt(int64(d.int32()))
//line /usr/local/go/src/encoding/binary/binary.go:645
		// _ = "end of CoverTab[1354]"
	case reflect.Int64:
//line /usr/local/go/src/encoding/binary/binary.go:646
		_go_fuzz_dep_.CoverTab[1355]++
								v.SetInt(d.int64())
//line /usr/local/go/src/encoding/binary/binary.go:647
		// _ = "end of CoverTab[1355]"

	case reflect.Uint8:
//line /usr/local/go/src/encoding/binary/binary.go:649
		_go_fuzz_dep_.CoverTab[1356]++
								v.SetUint(uint64(d.uint8()))
//line /usr/local/go/src/encoding/binary/binary.go:650
		// _ = "end of CoverTab[1356]"
	case reflect.Uint16:
//line /usr/local/go/src/encoding/binary/binary.go:651
		_go_fuzz_dep_.CoverTab[1357]++
								v.SetUint(uint64(d.uint16()))
//line /usr/local/go/src/encoding/binary/binary.go:652
		// _ = "end of CoverTab[1357]"
	case reflect.Uint32:
//line /usr/local/go/src/encoding/binary/binary.go:653
		_go_fuzz_dep_.CoverTab[1358]++
								v.SetUint(uint64(d.uint32()))
//line /usr/local/go/src/encoding/binary/binary.go:654
		// _ = "end of CoverTab[1358]"
	case reflect.Uint64:
//line /usr/local/go/src/encoding/binary/binary.go:655
		_go_fuzz_dep_.CoverTab[1359]++
								v.SetUint(d.uint64())
//line /usr/local/go/src/encoding/binary/binary.go:656
		// _ = "end of CoverTab[1359]"

	case reflect.Float32:
//line /usr/local/go/src/encoding/binary/binary.go:658
		_go_fuzz_dep_.CoverTab[1360]++
								v.SetFloat(float64(math.Float32frombits(d.uint32())))
//line /usr/local/go/src/encoding/binary/binary.go:659
		// _ = "end of CoverTab[1360]"
	case reflect.Float64:
//line /usr/local/go/src/encoding/binary/binary.go:660
		_go_fuzz_dep_.CoverTab[1361]++
								v.SetFloat(math.Float64frombits(d.uint64()))
//line /usr/local/go/src/encoding/binary/binary.go:661
		// _ = "end of CoverTab[1361]"

	case reflect.Complex64:
//line /usr/local/go/src/encoding/binary/binary.go:663
		_go_fuzz_dep_.CoverTab[1362]++
								v.SetComplex(complex(
			float64(math.Float32frombits(d.uint32())),
			float64(math.Float32frombits(d.uint32())),
		))
//line /usr/local/go/src/encoding/binary/binary.go:667
		// _ = "end of CoverTab[1362]"
	case reflect.Complex128:
//line /usr/local/go/src/encoding/binary/binary.go:668
		_go_fuzz_dep_.CoverTab[1363]++
								v.SetComplex(complex(
			math.Float64frombits(d.uint64()),
			math.Float64frombits(d.uint64()),
		))
//line /usr/local/go/src/encoding/binary/binary.go:672
		// _ = "end of CoverTab[1363]"
//line /usr/local/go/src/encoding/binary/binary.go:672
	default:
//line /usr/local/go/src/encoding/binary/binary.go:672
		_go_fuzz_dep_.CoverTab[1364]++
//line /usr/local/go/src/encoding/binary/binary.go:672
		// _ = "end of CoverTab[1364]"
	}
//line /usr/local/go/src/encoding/binary/binary.go:673
	// _ = "end of CoverTab[1347]"
}

func (e *encoder) value(v reflect.Value) {
//line /usr/local/go/src/encoding/binary/binary.go:676
	_go_fuzz_dep_.CoverTab[1371]++
							switch v.Kind() {
	case reflect.Array:
//line /usr/local/go/src/encoding/binary/binary.go:678
		_go_fuzz_dep_.CoverTab[1372]++
								l := v.Len()
								for i := 0; i < l; i++ {
//line /usr/local/go/src/encoding/binary/binary.go:680
			_go_fuzz_dep_.CoverTab[1381]++
									e.value(v.Index(i))
//line /usr/local/go/src/encoding/binary/binary.go:681
			// _ = "end of CoverTab[1381]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:682
		// _ = "end of CoverTab[1372]"

	case reflect.Struct:
//line /usr/local/go/src/encoding/binary/binary.go:684
		_go_fuzz_dep_.CoverTab[1373]++
								t := v.Type()
								l := v.NumField()
								for i := 0; i < l; i++ {
//line /usr/local/go/src/encoding/binary/binary.go:687
			_go_fuzz_dep_.CoverTab[1382]++

									if v := v.Field(i); v.CanSet() || func() bool {
//line /usr/local/go/src/encoding/binary/binary.go:689
				_go_fuzz_dep_.CoverTab[1383]++
//line /usr/local/go/src/encoding/binary/binary.go:689
				return t.Field(i).Name != "_"
//line /usr/local/go/src/encoding/binary/binary.go:689
				// _ = "end of CoverTab[1383]"
//line /usr/local/go/src/encoding/binary/binary.go:689
			}() {
//line /usr/local/go/src/encoding/binary/binary.go:689
				_go_fuzz_dep_.CoverTab[1384]++
										e.value(v)
//line /usr/local/go/src/encoding/binary/binary.go:690
				// _ = "end of CoverTab[1384]"
			} else {
//line /usr/local/go/src/encoding/binary/binary.go:691
				_go_fuzz_dep_.CoverTab[1385]++
										e.skip(v)
//line /usr/local/go/src/encoding/binary/binary.go:692
				// _ = "end of CoverTab[1385]"
			}
//line /usr/local/go/src/encoding/binary/binary.go:693
			// _ = "end of CoverTab[1382]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:694
		// _ = "end of CoverTab[1373]"

	case reflect.Slice:
//line /usr/local/go/src/encoding/binary/binary.go:696
		_go_fuzz_dep_.CoverTab[1374]++
								l := v.Len()
								for i := 0; i < l; i++ {
//line /usr/local/go/src/encoding/binary/binary.go:698
			_go_fuzz_dep_.CoverTab[1386]++
									e.value(v.Index(i))
//line /usr/local/go/src/encoding/binary/binary.go:699
			// _ = "end of CoverTab[1386]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:700
		// _ = "end of CoverTab[1374]"

	case reflect.Bool:
//line /usr/local/go/src/encoding/binary/binary.go:702
		_go_fuzz_dep_.CoverTab[1375]++
								e.bool(v.Bool())
//line /usr/local/go/src/encoding/binary/binary.go:703
		// _ = "end of CoverTab[1375]"

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/binary/binary.go:705
		_go_fuzz_dep_.CoverTab[1376]++
								switch v.Type().Kind() {
		case reflect.Int8:
//line /usr/local/go/src/encoding/binary/binary.go:707
			_go_fuzz_dep_.CoverTab[1387]++
									e.int8(int8(v.Int()))
//line /usr/local/go/src/encoding/binary/binary.go:708
			// _ = "end of CoverTab[1387]"
		case reflect.Int16:
//line /usr/local/go/src/encoding/binary/binary.go:709
			_go_fuzz_dep_.CoverTab[1388]++
									e.int16(int16(v.Int()))
//line /usr/local/go/src/encoding/binary/binary.go:710
			// _ = "end of CoverTab[1388]"
		case reflect.Int32:
//line /usr/local/go/src/encoding/binary/binary.go:711
			_go_fuzz_dep_.CoverTab[1389]++
									e.int32(int32(v.Int()))
//line /usr/local/go/src/encoding/binary/binary.go:712
			// _ = "end of CoverTab[1389]"
		case reflect.Int64:
//line /usr/local/go/src/encoding/binary/binary.go:713
			_go_fuzz_dep_.CoverTab[1390]++
									e.int64(v.Int())
//line /usr/local/go/src/encoding/binary/binary.go:714
			// _ = "end of CoverTab[1390]"
//line /usr/local/go/src/encoding/binary/binary.go:714
		default:
//line /usr/local/go/src/encoding/binary/binary.go:714
			_go_fuzz_dep_.CoverTab[1391]++
//line /usr/local/go/src/encoding/binary/binary.go:714
			// _ = "end of CoverTab[1391]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:715
		// _ = "end of CoverTab[1376]"

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/binary/binary.go:717
		_go_fuzz_dep_.CoverTab[1377]++
								switch v.Type().Kind() {
		case reflect.Uint8:
//line /usr/local/go/src/encoding/binary/binary.go:719
			_go_fuzz_dep_.CoverTab[1392]++
									e.uint8(uint8(v.Uint()))
//line /usr/local/go/src/encoding/binary/binary.go:720
			// _ = "end of CoverTab[1392]"
		case reflect.Uint16:
//line /usr/local/go/src/encoding/binary/binary.go:721
			_go_fuzz_dep_.CoverTab[1393]++
									e.uint16(uint16(v.Uint()))
//line /usr/local/go/src/encoding/binary/binary.go:722
			// _ = "end of CoverTab[1393]"
		case reflect.Uint32:
//line /usr/local/go/src/encoding/binary/binary.go:723
			_go_fuzz_dep_.CoverTab[1394]++
									e.uint32(uint32(v.Uint()))
//line /usr/local/go/src/encoding/binary/binary.go:724
			// _ = "end of CoverTab[1394]"
		case reflect.Uint64:
//line /usr/local/go/src/encoding/binary/binary.go:725
			_go_fuzz_dep_.CoverTab[1395]++
									e.uint64(v.Uint())
//line /usr/local/go/src/encoding/binary/binary.go:726
			// _ = "end of CoverTab[1395]"
//line /usr/local/go/src/encoding/binary/binary.go:726
		default:
//line /usr/local/go/src/encoding/binary/binary.go:726
			_go_fuzz_dep_.CoverTab[1396]++
//line /usr/local/go/src/encoding/binary/binary.go:726
			// _ = "end of CoverTab[1396]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:727
		// _ = "end of CoverTab[1377]"

	case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/encoding/binary/binary.go:729
		_go_fuzz_dep_.CoverTab[1378]++
								switch v.Type().Kind() {
		case reflect.Float32:
//line /usr/local/go/src/encoding/binary/binary.go:731
			_go_fuzz_dep_.CoverTab[1397]++
									e.uint32(math.Float32bits(float32(v.Float())))
//line /usr/local/go/src/encoding/binary/binary.go:732
			// _ = "end of CoverTab[1397]"
		case reflect.Float64:
//line /usr/local/go/src/encoding/binary/binary.go:733
			_go_fuzz_dep_.CoverTab[1398]++
									e.uint64(math.Float64bits(v.Float()))
//line /usr/local/go/src/encoding/binary/binary.go:734
			// _ = "end of CoverTab[1398]"
//line /usr/local/go/src/encoding/binary/binary.go:734
		default:
//line /usr/local/go/src/encoding/binary/binary.go:734
			_go_fuzz_dep_.CoverTab[1399]++
//line /usr/local/go/src/encoding/binary/binary.go:734
			// _ = "end of CoverTab[1399]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:735
		// _ = "end of CoverTab[1378]"

	case reflect.Complex64, reflect.Complex128:
//line /usr/local/go/src/encoding/binary/binary.go:737
		_go_fuzz_dep_.CoverTab[1379]++
								switch v.Type().Kind() {
		case reflect.Complex64:
//line /usr/local/go/src/encoding/binary/binary.go:739
			_go_fuzz_dep_.CoverTab[1400]++
									x := v.Complex()
									e.uint32(math.Float32bits(float32(real(x))))
									e.uint32(math.Float32bits(float32(imag(x))))
//line /usr/local/go/src/encoding/binary/binary.go:742
			// _ = "end of CoverTab[1400]"
		case reflect.Complex128:
//line /usr/local/go/src/encoding/binary/binary.go:743
			_go_fuzz_dep_.CoverTab[1401]++
									x := v.Complex()
									e.uint64(math.Float64bits(real(x)))
									e.uint64(math.Float64bits(imag(x)))
//line /usr/local/go/src/encoding/binary/binary.go:746
			// _ = "end of CoverTab[1401]"
//line /usr/local/go/src/encoding/binary/binary.go:746
		default:
//line /usr/local/go/src/encoding/binary/binary.go:746
			_go_fuzz_dep_.CoverTab[1402]++
//line /usr/local/go/src/encoding/binary/binary.go:746
			// _ = "end of CoverTab[1402]"
		}
//line /usr/local/go/src/encoding/binary/binary.go:747
		// _ = "end of CoverTab[1379]"
//line /usr/local/go/src/encoding/binary/binary.go:747
	default:
//line /usr/local/go/src/encoding/binary/binary.go:747
		_go_fuzz_dep_.CoverTab[1380]++
//line /usr/local/go/src/encoding/binary/binary.go:747
		// _ = "end of CoverTab[1380]"
	}
//line /usr/local/go/src/encoding/binary/binary.go:748
	// _ = "end of CoverTab[1371]"
}

func (d *decoder) skip(v reflect.Value) {
//line /usr/local/go/src/encoding/binary/binary.go:751
	_go_fuzz_dep_.CoverTab[1403]++
							d.offset += dataSize(v)
//line /usr/local/go/src/encoding/binary/binary.go:752
	// _ = "end of CoverTab[1403]"
}

func (e *encoder) skip(v reflect.Value) {
//line /usr/local/go/src/encoding/binary/binary.go:755
	_go_fuzz_dep_.CoverTab[1404]++
							n := dataSize(v)
							zero := e.buf[e.offset : e.offset+n]
							for i := range zero {
//line /usr/local/go/src/encoding/binary/binary.go:758
		_go_fuzz_dep_.CoverTab[1406]++
								zero[i] = 0
//line /usr/local/go/src/encoding/binary/binary.go:759
		// _ = "end of CoverTab[1406]"
	}
//line /usr/local/go/src/encoding/binary/binary.go:760
	// _ = "end of CoverTab[1404]"
//line /usr/local/go/src/encoding/binary/binary.go:760
	_go_fuzz_dep_.CoverTab[1405]++
							e.offset += n
//line /usr/local/go/src/encoding/binary/binary.go:761
	// _ = "end of CoverTab[1405]"
}

// intDataSize returns the size of the data required to represent the data when encoded.
//line /usr/local/go/src/encoding/binary/binary.go:764
// It returns zero if the type cannot be implemented by the fast path in Read or Write.
//line /usr/local/go/src/encoding/binary/binary.go:766
func intDataSize(data any) int {
//line /usr/local/go/src/encoding/binary/binary.go:766
	_go_fuzz_dep_.CoverTab[1407]++
							switch data := data.(type) {
	case bool, int8, uint8, *bool, *int8, *uint8:
//line /usr/local/go/src/encoding/binary/binary.go:768
		_go_fuzz_dep_.CoverTab[1409]++
								return 1
//line /usr/local/go/src/encoding/binary/binary.go:769
		// _ = "end of CoverTab[1409]"
	case []bool:
//line /usr/local/go/src/encoding/binary/binary.go:770
		_go_fuzz_dep_.CoverTab[1410]++
								return len(data)
//line /usr/local/go/src/encoding/binary/binary.go:771
		// _ = "end of CoverTab[1410]"
	case []int8:
//line /usr/local/go/src/encoding/binary/binary.go:772
		_go_fuzz_dep_.CoverTab[1411]++
								return len(data)
//line /usr/local/go/src/encoding/binary/binary.go:773
		// _ = "end of CoverTab[1411]"
	case []uint8:
//line /usr/local/go/src/encoding/binary/binary.go:774
		_go_fuzz_dep_.CoverTab[1412]++
								return len(data)
//line /usr/local/go/src/encoding/binary/binary.go:775
		// _ = "end of CoverTab[1412]"
	case int16, uint16, *int16, *uint16:
//line /usr/local/go/src/encoding/binary/binary.go:776
		_go_fuzz_dep_.CoverTab[1413]++
								return 2
//line /usr/local/go/src/encoding/binary/binary.go:777
		// _ = "end of CoverTab[1413]"
	case []int16:
//line /usr/local/go/src/encoding/binary/binary.go:778
		_go_fuzz_dep_.CoverTab[1414]++
								return 2 * len(data)
//line /usr/local/go/src/encoding/binary/binary.go:779
		// _ = "end of CoverTab[1414]"
	case []uint16:
//line /usr/local/go/src/encoding/binary/binary.go:780
		_go_fuzz_dep_.CoverTab[1415]++
								return 2 * len(data)
//line /usr/local/go/src/encoding/binary/binary.go:781
		// _ = "end of CoverTab[1415]"
	case int32, uint32, *int32, *uint32:
//line /usr/local/go/src/encoding/binary/binary.go:782
		_go_fuzz_dep_.CoverTab[1416]++
								return 4
//line /usr/local/go/src/encoding/binary/binary.go:783
		// _ = "end of CoverTab[1416]"
	case []int32:
//line /usr/local/go/src/encoding/binary/binary.go:784
		_go_fuzz_dep_.CoverTab[1417]++
								return 4 * len(data)
//line /usr/local/go/src/encoding/binary/binary.go:785
		// _ = "end of CoverTab[1417]"
	case []uint32:
//line /usr/local/go/src/encoding/binary/binary.go:786
		_go_fuzz_dep_.CoverTab[1418]++
								return 4 * len(data)
//line /usr/local/go/src/encoding/binary/binary.go:787
		// _ = "end of CoverTab[1418]"
	case int64, uint64, *int64, *uint64:
//line /usr/local/go/src/encoding/binary/binary.go:788
		_go_fuzz_dep_.CoverTab[1419]++
								return 8
//line /usr/local/go/src/encoding/binary/binary.go:789
		// _ = "end of CoverTab[1419]"
	case []int64:
//line /usr/local/go/src/encoding/binary/binary.go:790
		_go_fuzz_dep_.CoverTab[1420]++
								return 8 * len(data)
//line /usr/local/go/src/encoding/binary/binary.go:791
		// _ = "end of CoverTab[1420]"
	case []uint64:
//line /usr/local/go/src/encoding/binary/binary.go:792
		_go_fuzz_dep_.CoverTab[1421]++
								return 8 * len(data)
//line /usr/local/go/src/encoding/binary/binary.go:793
		// _ = "end of CoverTab[1421]"
	case float32, *float32:
//line /usr/local/go/src/encoding/binary/binary.go:794
		_go_fuzz_dep_.CoverTab[1422]++
								return 4
//line /usr/local/go/src/encoding/binary/binary.go:795
		// _ = "end of CoverTab[1422]"
	case float64, *float64:
//line /usr/local/go/src/encoding/binary/binary.go:796
		_go_fuzz_dep_.CoverTab[1423]++
								return 8
//line /usr/local/go/src/encoding/binary/binary.go:797
		// _ = "end of CoverTab[1423]"
	case []float32:
//line /usr/local/go/src/encoding/binary/binary.go:798
		_go_fuzz_dep_.CoverTab[1424]++
								return 4 * len(data)
//line /usr/local/go/src/encoding/binary/binary.go:799
		// _ = "end of CoverTab[1424]"
	case []float64:
//line /usr/local/go/src/encoding/binary/binary.go:800
		_go_fuzz_dep_.CoverTab[1425]++
								return 8 * len(data)
//line /usr/local/go/src/encoding/binary/binary.go:801
		// _ = "end of CoverTab[1425]"
	}
//line /usr/local/go/src/encoding/binary/binary.go:802
	// _ = "end of CoverTab[1407]"
//line /usr/local/go/src/encoding/binary/binary.go:802
	_go_fuzz_dep_.CoverTab[1408]++
							return 0
//line /usr/local/go/src/encoding/binary/binary.go:803
	// _ = "end of CoverTab[1408]"
}

//line /usr/local/go/src/encoding/binary/binary.go:804
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/binary/binary.go:804
var _ = _go_fuzz_dep_.CoverTab
