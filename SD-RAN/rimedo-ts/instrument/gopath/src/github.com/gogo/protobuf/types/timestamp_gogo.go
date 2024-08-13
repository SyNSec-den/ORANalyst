// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2016, The GoGo Authors. All rights reserved.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:29
package types

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:29
)

import (
	"time"
)

func NewPopulatedTimestamp(r interface {
	Int63() int64
}, easy bool) *Timestamp {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:37
	_go_fuzz_dep_.CoverTab[138335]++
												this := &Timestamp{}
												ns := int64(r.Int63())
												this.Seconds = ns / 1e9
												this.Nanos = int32(ns % 1e9)
												return this
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:42
	// _ = "end of CoverTab[138335]"
}

func (ts *Timestamp) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:45
	_go_fuzz_dep_.CoverTab[138336]++
												return TimestampString(ts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:46
	// _ = "end of CoverTab[138336]"
}

func NewPopulatedStdTime(r interface {
	Int63() int64
}, easy bool) *time.Time {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:51
	_go_fuzz_dep_.CoverTab[138337]++
												timestamp := NewPopulatedTimestamp(r, easy)
												t, err := TimestampFromProto(timestamp)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:54
		_go_fuzz_dep_.CoverTab[138339]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:55
		// _ = "end of CoverTab[138339]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:56
		_go_fuzz_dep_.CoverTab[138340]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:56
		// _ = "end of CoverTab[138340]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:56
	// _ = "end of CoverTab[138337]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:56
	_go_fuzz_dep_.CoverTab[138338]++
												return &t
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:57
	// _ = "end of CoverTab[138338]"
}

func SizeOfStdTime(t time.Time) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:60
	_go_fuzz_dep_.CoverTab[138341]++
												ts, err := TimestampProto(t)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:62
		_go_fuzz_dep_.CoverTab[138343]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:63
		// _ = "end of CoverTab[138343]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:64
		_go_fuzz_dep_.CoverTab[138344]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:64
		// _ = "end of CoverTab[138344]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:64
	// _ = "end of CoverTab[138341]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:64
	_go_fuzz_dep_.CoverTab[138342]++
												return ts.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:65
	// _ = "end of CoverTab[138342]"
}

func StdTimeMarshal(t time.Time) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:68
	_go_fuzz_dep_.CoverTab[138345]++
												size := SizeOfStdTime(t)
												buf := make([]byte, size)
												_, err := StdTimeMarshalTo(t, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:72
	// _ = "end of CoverTab[138345]"
}

func StdTimeMarshalTo(t time.Time, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:75
	_go_fuzz_dep_.CoverTab[138346]++
												ts, err := TimestampProto(t)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:77
		_go_fuzz_dep_.CoverTab[138348]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:78
		// _ = "end of CoverTab[138348]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:79
		_go_fuzz_dep_.CoverTab[138349]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:79
		// _ = "end of CoverTab[138349]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:79
	// _ = "end of CoverTab[138346]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:79
	_go_fuzz_dep_.CoverTab[138347]++
												return ts.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:80
	// _ = "end of CoverTab[138347]"
}

func StdTimeUnmarshal(t *time.Time, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:83
	_go_fuzz_dep_.CoverTab[138350]++
												ts := &Timestamp{}
												if err := ts.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:85
		_go_fuzz_dep_.CoverTab[138353]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:86
		// _ = "end of CoverTab[138353]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:87
		_go_fuzz_dep_.CoverTab[138354]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:87
		// _ = "end of CoverTab[138354]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:87
	// _ = "end of CoverTab[138350]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:87
	_go_fuzz_dep_.CoverTab[138351]++
												tt, err := TimestampFromProto(ts)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:89
		_go_fuzz_dep_.CoverTab[138355]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:90
		// _ = "end of CoverTab[138355]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:91
		_go_fuzz_dep_.CoverTab[138356]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:91
		// _ = "end of CoverTab[138356]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:91
	// _ = "end of CoverTab[138351]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:91
	_go_fuzz_dep_.CoverTab[138352]++
												*t = tt
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:93
	// _ = "end of CoverTab[138352]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:94
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp_gogo.go:94
var _ = _go_fuzz_dep_.CoverTab
