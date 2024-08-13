// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:5
)

import (
	protoV2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

// Size returns the size in bytes of the wire-format encoding of m.
func Size(m Message) int {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:13
	_go_fuzz_dep_.CoverTab[62386]++
											if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:14
		_go_fuzz_dep_.CoverTab[62388]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:15
		// _ = "end of CoverTab[62388]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:16
		_go_fuzz_dep_.CoverTab[62389]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:16
		// _ = "end of CoverTab[62389]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:16
	// _ = "end of CoverTab[62386]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:16
	_go_fuzz_dep_.CoverTab[62387]++
											mi := MessageV2(m)
											return protoV2.Size(mi)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:18
	// _ = "end of CoverTab[62387]"
}

// Marshal returns the wire-format encoding of m.
func Marshal(m Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:22
	_go_fuzz_dep_.CoverTab[62390]++
											b, err := marshalAppend(nil, m, false)
											if b == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:24
		_go_fuzz_dep_.CoverTab[62392]++
												b = zeroBytes
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:25
		// _ = "end of CoverTab[62392]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:26
		_go_fuzz_dep_.CoverTab[62393]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:26
		// _ = "end of CoverTab[62393]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:26
	// _ = "end of CoverTab[62390]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:26
	_go_fuzz_dep_.CoverTab[62391]++
											return b, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:27
	// _ = "end of CoverTab[62391]"
}

var zeroBytes = make([]byte, 0, 0)

func marshalAppend(buf []byte, m Message, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:32
	_go_fuzz_dep_.CoverTab[62394]++
											if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:33
		_go_fuzz_dep_.CoverTab[62398]++
												return nil, ErrNil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:34
		// _ = "end of CoverTab[62398]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:35
		_go_fuzz_dep_.CoverTab[62399]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:35
		// _ = "end of CoverTab[62399]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:35
	// _ = "end of CoverTab[62394]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:35
	_go_fuzz_dep_.CoverTab[62395]++
											mi := MessageV2(m)
											nbuf, err := protoV2.MarshalOptions{
		Deterministic:	deterministic,
		AllowPartial:	true,
	}.MarshalAppend(buf, mi)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:41
		_go_fuzz_dep_.CoverTab[62400]++
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:42
		// _ = "end of CoverTab[62400]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:43
		_go_fuzz_dep_.CoverTab[62401]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:43
		// _ = "end of CoverTab[62401]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:43
	// _ = "end of CoverTab[62395]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:43
	_go_fuzz_dep_.CoverTab[62396]++
											if len(buf) == len(nbuf) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:44
		_go_fuzz_dep_.CoverTab[62402]++
												if !mi.ProtoReflect().IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:45
			_go_fuzz_dep_.CoverTab[62403]++
													return buf, ErrNil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:46
			// _ = "end of CoverTab[62403]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:47
			_go_fuzz_dep_.CoverTab[62404]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:47
			// _ = "end of CoverTab[62404]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:47
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:47
		// _ = "end of CoverTab[62402]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:48
		_go_fuzz_dep_.CoverTab[62405]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:48
		// _ = "end of CoverTab[62405]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:48
	// _ = "end of CoverTab[62396]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:48
	_go_fuzz_dep_.CoverTab[62397]++
											return nbuf, checkRequiredNotSet(mi)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:49
	// _ = "end of CoverTab[62397]"
}

// Unmarshal parses a wire-format message in b and places the decoded results in m.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:52
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:52
// Unmarshal resets m before starting to unmarshal, so any existing data in m is always
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:52
// removed. Use UnmarshalMerge to preserve and append to existing data.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:56
func Unmarshal(b []byte, m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:56
	_go_fuzz_dep_.CoverTab[62406]++
											m.Reset()
											return UnmarshalMerge(b, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:58
	// _ = "end of CoverTab[62406]"
}

// UnmarshalMerge parses a wire-format message in b and places the decoded results in m.
func UnmarshalMerge(b []byte, m Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:62
	_go_fuzz_dep_.CoverTab[62407]++
											mi := MessageV2(m)
											out, err := protoV2.UnmarshalOptions{
		AllowPartial:	true,
		Merge:		true,
	}.UnmarshalState(protoiface.UnmarshalInput{
		Buf:		b,
		Message:	mi.ProtoReflect(),
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:71
		_go_fuzz_dep_.CoverTab[62410]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:72
		// _ = "end of CoverTab[62410]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:73
		_go_fuzz_dep_.CoverTab[62411]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:73
		// _ = "end of CoverTab[62411]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:73
	// _ = "end of CoverTab[62407]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:73
	_go_fuzz_dep_.CoverTab[62408]++
											if out.Flags&protoiface.UnmarshalInitialized > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:74
		_go_fuzz_dep_.CoverTab[62412]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:75
		// _ = "end of CoverTab[62412]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:76
		_go_fuzz_dep_.CoverTab[62413]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:76
		// _ = "end of CoverTab[62413]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:76
	// _ = "end of CoverTab[62408]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:76
	_go_fuzz_dep_.CoverTab[62409]++
											return checkRequiredNotSet(mi)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:77
	// _ = "end of CoverTab[62409]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:78
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/wire.go:78
var _ = _go_fuzz_dep_.CoverTab
