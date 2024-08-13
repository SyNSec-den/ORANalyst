// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:5
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	protoV2 "google.golang.org/protobuf/proto"
)

var (
	// Deprecated: No longer returned.
	ErrNil	= errors.New("proto: Marshal called with nil")

	// Deprecated: No longer returned.
	ErrTooLarge	= errors.New("proto: message encodes to over 2 GB")

	// Deprecated: No longer returned.
	ErrInternalBadWireType	= errors.New("proto: internal error: bad wiretype for oneof")
)

// Deprecated: Do not use.
type Stats struct{ Emalloc, Dmalloc, Encode, Decode, Chit, Cmiss, Size uint64 }

// Deprecated: Do not use.
func GetStats() Stats {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:31
	_go_fuzz_dep_.CoverTab[61223]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:31
	return Stats{}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:31
	// _ = "end of CoverTab[61223]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:31
}

// Deprecated: Do not use.
func MarshalMessageSet(interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:34
	_go_fuzz_dep_.CoverTab[61224]++
												return nil, errors.New("proto: not implemented")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:35
	// _ = "end of CoverTab[61224]"
}

// Deprecated: Do not use.
func UnmarshalMessageSet([]byte, interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:39
	_go_fuzz_dep_.CoverTab[61225]++
												return errors.New("proto: not implemented")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:40
	// _ = "end of CoverTab[61225]"
}

// Deprecated: Do not use.
func MarshalMessageSetJSON(interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:44
	_go_fuzz_dep_.CoverTab[61226]++
												return nil, errors.New("proto: not implemented")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:45
	// _ = "end of CoverTab[61226]"
}

// Deprecated: Do not use.
func UnmarshalMessageSetJSON([]byte, interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:49
	_go_fuzz_dep_.CoverTab[61227]++
												return errors.New("proto: not implemented")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:50
	// _ = "end of CoverTab[61227]"
}

// Deprecated: Do not use.
func RegisterMessageSetType(Message, int32, string) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:54
	_go_fuzz_dep_.CoverTab[61228]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:54
	// _ = "end of CoverTab[61228]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:54
}

// Deprecated: Do not use.
func EnumName(m map[int32]string, v int32) string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:57
	_go_fuzz_dep_.CoverTab[61229]++
												s, ok := m[v]
												if ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:59
		_go_fuzz_dep_.CoverTab[61231]++
													return s
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:60
		// _ = "end of CoverTab[61231]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:61
		_go_fuzz_dep_.CoverTab[61232]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:61
		// _ = "end of CoverTab[61232]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:61
	// _ = "end of CoverTab[61229]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:61
	_go_fuzz_dep_.CoverTab[61230]++
												return strconv.Itoa(int(v))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:62
	// _ = "end of CoverTab[61230]"
}

// Deprecated: Do not use.
func UnmarshalJSONEnum(m map[string]int32, data []byte, enumName string) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:66
	_go_fuzz_dep_.CoverTab[61233]++
												if data[0] == '"' {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:67
		_go_fuzz_dep_.CoverTab[61236]++
		// New style: enums are strings.
		var repr string
		if err := json.Unmarshal(data, &repr); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:70
			_go_fuzz_dep_.CoverTab[61239]++
														return -1, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:71
			// _ = "end of CoverTab[61239]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:72
			_go_fuzz_dep_.CoverTab[61240]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:72
			// _ = "end of CoverTab[61240]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:72
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:72
		// _ = "end of CoverTab[61236]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:72
		_go_fuzz_dep_.CoverTab[61237]++
													val, ok := m[repr]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:74
			_go_fuzz_dep_.CoverTab[61241]++
														return 0, fmt.Errorf("unrecognized enum %s value %q", enumName, repr)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:75
			// _ = "end of CoverTab[61241]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:76
			_go_fuzz_dep_.CoverTab[61242]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:76
			// _ = "end of CoverTab[61242]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:76
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:76
		// _ = "end of CoverTab[61237]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:76
		_go_fuzz_dep_.CoverTab[61238]++
													return val, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:77
		// _ = "end of CoverTab[61238]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:78
		_go_fuzz_dep_.CoverTab[61243]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:78
		// _ = "end of CoverTab[61243]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:78
	// _ = "end of CoverTab[61233]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:78
	_go_fuzz_dep_.CoverTab[61234]++
	// Old style: enums are ints.
	var val int32
	if err := json.Unmarshal(data, &val); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:81
		_go_fuzz_dep_.CoverTab[61244]++
													return 0, fmt.Errorf("cannot unmarshal %#q into enum %s", data, enumName)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:82
		// _ = "end of CoverTab[61244]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:83
		_go_fuzz_dep_.CoverTab[61245]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:83
		// _ = "end of CoverTab[61245]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:83
	// _ = "end of CoverTab[61234]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:83
	_go_fuzz_dep_.CoverTab[61235]++
												return val, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:84
	// _ = "end of CoverTab[61235]"
}

// Deprecated: Do not use; this type existed for intenal-use only.
type InternalMessageInfo struct{}

// Deprecated: Do not use; this method existed for intenal-use only.
func (*InternalMessageInfo) DiscardUnknown(m Message) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:91
	_go_fuzz_dep_.CoverTab[61246]++
												DiscardUnknown(m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:92
	// _ = "end of CoverTab[61246]"
}

// Deprecated: Do not use; this method existed for intenal-use only.
func (*InternalMessageInfo) Marshal(b []byte, m Message, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:96
	_go_fuzz_dep_.CoverTab[61247]++
												return protoV2.MarshalOptions{Deterministic: deterministic}.MarshalAppend(b, MessageV2(m))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:97
	// _ = "end of CoverTab[61247]"
}

// Deprecated: Do not use; this method existed for intenal-use only.
func (*InternalMessageInfo) Merge(dst, src Message) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:101
	_go_fuzz_dep_.CoverTab[61248]++
												protoV2.Merge(MessageV2(dst), MessageV2(src))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:102
	// _ = "end of CoverTab[61248]"
}

// Deprecated: Do not use; this method existed for intenal-use only.
func (*InternalMessageInfo) Size(m Message) int {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:106
	_go_fuzz_dep_.CoverTab[61249]++
												return protoV2.Size(MessageV2(m))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:107
	// _ = "end of CoverTab[61249]"
}

// Deprecated: Do not use; this method existed for intenal-use only.
func (*InternalMessageInfo) Unmarshal(m Message, b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:111
	_go_fuzz_dep_.CoverTab[61250]++
												return protoV2.UnmarshalOptions{Merge: true}.Unmarshal(b, MessageV2(m))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:112
	// _ = "end of CoverTab[61250]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/deprecated.go:113
var _ = _go_fuzz_dep_.CoverTab
