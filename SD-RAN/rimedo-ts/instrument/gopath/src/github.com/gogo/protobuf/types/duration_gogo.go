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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:29
package types

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:29
)

import (
	"fmt"
	"time"
)

func NewPopulatedDuration(r interface {
	Int63() int64
}, easy bool) *Duration {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:38
	_go_fuzz_dep_.CoverTab[136033]++
												this := &Duration{}
												maxSecs := time.Hour.Nanoseconds() / 1e9
												max := 2 * maxSecs
												s := int64(r.Int63()) % max
												s -= maxSecs
												neg := int64(1)
												if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:45
		_go_fuzz_dep_.CoverTab[136035]++
													neg = -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:46
		// _ = "end of CoverTab[136035]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:47
		_go_fuzz_dep_.CoverTab[136036]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:47
		// _ = "end of CoverTab[136036]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:47
	// _ = "end of CoverTab[136033]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:47
	_go_fuzz_dep_.CoverTab[136034]++
												this.Seconds = s
												this.Nanos = int32(neg * (r.Int63() % 1e9))
												return this
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:50
	// _ = "end of CoverTab[136034]"
}

func (d *Duration) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:53
	_go_fuzz_dep_.CoverTab[136037]++
												td, err := DurationFromProto(d)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:55
		_go_fuzz_dep_.CoverTab[136039]++
													return fmt.Sprintf("(%v)", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:56
		// _ = "end of CoverTab[136039]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:57
		_go_fuzz_dep_.CoverTab[136040]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:57
		// _ = "end of CoverTab[136040]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:57
	// _ = "end of CoverTab[136037]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:57
	_go_fuzz_dep_.CoverTab[136038]++
												return td.String()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:58
	// _ = "end of CoverTab[136038]"
}

func NewPopulatedStdDuration(r interface {
	Int63() int64
}, easy bool) *time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:63
	_go_fuzz_dep_.CoverTab[136041]++
												dur := NewPopulatedDuration(r, easy)
												d, err := DurationFromProto(dur)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:66
		_go_fuzz_dep_.CoverTab[136043]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:67
		// _ = "end of CoverTab[136043]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:68
		_go_fuzz_dep_.CoverTab[136044]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:68
		// _ = "end of CoverTab[136044]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:68
	// _ = "end of CoverTab[136041]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:68
	_go_fuzz_dep_.CoverTab[136042]++
												return &d
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:69
	// _ = "end of CoverTab[136042]"
}

func SizeOfStdDuration(d time.Duration) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:72
	_go_fuzz_dep_.CoverTab[136045]++
												dur := DurationProto(d)
												return dur.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:74
	// _ = "end of CoverTab[136045]"
}

func StdDurationMarshal(d time.Duration) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:77
	_go_fuzz_dep_.CoverTab[136046]++
												size := SizeOfStdDuration(d)
												buf := make([]byte, size)
												_, err := StdDurationMarshalTo(d, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:81
	// _ = "end of CoverTab[136046]"
}

func StdDurationMarshalTo(d time.Duration, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:84
	_go_fuzz_dep_.CoverTab[136047]++
												dur := DurationProto(d)
												return dur.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:86
	// _ = "end of CoverTab[136047]"
}

func StdDurationUnmarshal(d *time.Duration, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:89
	_go_fuzz_dep_.CoverTab[136048]++
												dur := &Duration{}
												if err := dur.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:91
		_go_fuzz_dep_.CoverTab[136051]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:92
		// _ = "end of CoverTab[136051]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:93
		_go_fuzz_dep_.CoverTab[136052]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:93
		// _ = "end of CoverTab[136052]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:93
	// _ = "end of CoverTab[136048]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:93
	_go_fuzz_dep_.CoverTab[136049]++
												dd, err := DurationFromProto(dur)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:95
		_go_fuzz_dep_.CoverTab[136053]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:96
		// _ = "end of CoverTab[136053]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:97
		_go_fuzz_dep_.CoverTab[136054]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:97
		// _ = "end of CoverTab[136054]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:97
	// _ = "end of CoverTab[136049]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:97
	_go_fuzz_dep_.CoverTab[136050]++
												*d = dd
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:99
	// _ = "end of CoverTab[136050]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:100
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/duration_gogo.go:100
var _ = _go_fuzz_dep_.CoverTab
