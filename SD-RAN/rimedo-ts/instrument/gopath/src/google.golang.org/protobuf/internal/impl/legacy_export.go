// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:5
)

import (
	"encoding/binary"
	"encoding/json"
	"hash/crc32"
	"math"
	"reflect"

	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:22
// LegacyEnumName returns the name of enums used in legacy code.
func (Export) LegacyEnumName(ed protoreflect.EnumDescriptor) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:23
	_go_fuzz_dep_.CoverTab[57519]++
													return legacyEnumName(ed)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:24
	// _ = "end of CoverTab[57519]"
}

// LegacyMessageTypeOf returns the protoreflect.MessageType for m,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:27
// with name used as the message name if necessary.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:29
func (Export) LegacyMessageTypeOf(m protoiface.MessageV1, name protoreflect.FullName) protoreflect.MessageType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:29
	_go_fuzz_dep_.CoverTab[57520]++
													if mv := (Export{}).protoMessageV2Of(m); mv != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:30
		_go_fuzz_dep_.CoverTab[57522]++
														return mv.ProtoReflect().Type()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:31
		// _ = "end of CoverTab[57522]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:32
		_go_fuzz_dep_.CoverTab[57523]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:32
		// _ = "end of CoverTab[57523]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:32
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:32
	// _ = "end of CoverTab[57520]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:32
	_go_fuzz_dep_.CoverTab[57521]++
													return legacyLoadMessageType(reflect.TypeOf(m), name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:33
	// _ = "end of CoverTab[57521]"
}

// UnmarshalJSONEnum unmarshals an enum from a JSON-encoded input.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:36
// The input can either be a string representing the enum value by name,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:36
// or a number representing the enum number itself.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:39
func (Export) UnmarshalJSONEnum(ed protoreflect.EnumDescriptor, b []byte) (protoreflect.EnumNumber, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:39
	_go_fuzz_dep_.CoverTab[57524]++
													if b[0] == '"' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:40
		_go_fuzz_dep_.CoverTab[57525]++
														var name protoreflect.Name
														if err := json.Unmarshal(b, &name); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:42
			_go_fuzz_dep_.CoverTab[57528]++
															return 0, errors.New("invalid input for enum %v: %s", ed.FullName(), b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:43
			// _ = "end of CoverTab[57528]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:44
			_go_fuzz_dep_.CoverTab[57529]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:44
			// _ = "end of CoverTab[57529]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:44
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:44
		// _ = "end of CoverTab[57525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:44
		_go_fuzz_dep_.CoverTab[57526]++
														ev := ed.Values().ByName(name)
														if ev == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:46
			_go_fuzz_dep_.CoverTab[57530]++
															return 0, errors.New("invalid value for enum %v: %s", ed.FullName(), name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:47
			// _ = "end of CoverTab[57530]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:48
			_go_fuzz_dep_.CoverTab[57531]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:48
			// _ = "end of CoverTab[57531]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:48
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:48
		// _ = "end of CoverTab[57526]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:48
		_go_fuzz_dep_.CoverTab[57527]++
														return ev.Number(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:49
		// _ = "end of CoverTab[57527]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:50
		_go_fuzz_dep_.CoverTab[57532]++
														var num protoreflect.EnumNumber
														if err := json.Unmarshal(b, &num); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:52
			_go_fuzz_dep_.CoverTab[57534]++
															return 0, errors.New("invalid input for enum %v: %s", ed.FullName(), b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:53
			// _ = "end of CoverTab[57534]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:54
			_go_fuzz_dep_.CoverTab[57535]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:54
			// _ = "end of CoverTab[57535]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:54
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:54
		// _ = "end of CoverTab[57532]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:54
		_go_fuzz_dep_.CoverTab[57533]++
														return num, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:55
		// _ = "end of CoverTab[57533]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:56
	// _ = "end of CoverTab[57524]"
}

// CompressGZIP compresses the input as a GZIP-encoded file.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:59
// The current implementation does no compression.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:61
func (Export) CompressGZIP(in []byte) (out []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:61
	_go_fuzz_dep_.CoverTab[57536]++
	// RFC 1952, section 2.3.1.
	var gzipHeader = [10]byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff}

	// RFC 1951, section 3.2.4.
	var blockHeader [5]byte
	const maxBlockSize = math.MaxUint16
	numBlocks := 1 + len(in)/maxBlockSize

													// RFC 1952, section 2.3.1.
													var gzipFooter [8]byte
													binary.LittleEndian.PutUint32(gzipFooter[0:4], crc32.ChecksumIEEE(in))
													binary.LittleEndian.PutUint32(gzipFooter[4:8], uint32(len(in)))

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:76
	out = make([]byte, 0, len(gzipHeader)+len(blockHeader)*numBlocks+len(in)+len(gzipFooter))
	out = append(out, gzipHeader[:]...)
	for blockHeader[0] == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:78
		_go_fuzz_dep_.CoverTab[57538]++
														blockSize := maxBlockSize
														if blockSize > len(in) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:80
			_go_fuzz_dep_.CoverTab[57540]++
															blockHeader[0] = 0x01
															blockSize = len(in)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:82
			// _ = "end of CoverTab[57540]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:83
			_go_fuzz_dep_.CoverTab[57541]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:83
			// _ = "end of CoverTab[57541]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:83
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:83
		// _ = "end of CoverTab[57538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:83
		_go_fuzz_dep_.CoverTab[57539]++
														binary.LittleEndian.PutUint16(blockHeader[1:3], uint16(blockSize))
														binary.LittleEndian.PutUint16(blockHeader[3:5], ^uint16(blockSize))
														out = append(out, blockHeader[:]...)
														out = append(out, in[:blockSize]...)
														in = in[blockSize:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:88
		// _ = "end of CoverTab[57539]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:89
	// _ = "end of CoverTab[57536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:89
	_go_fuzz_dep_.CoverTab[57537]++
													out = append(out, gzipFooter[:]...)
													return out
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:91
	// _ = "end of CoverTab[57537]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:92
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_export.go:92
var _ = _go_fuzz_dep_.CoverTab
