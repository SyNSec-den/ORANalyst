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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:32
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:37
import (
	"errors"
	"fmt"
	"time"
)

const (
	// Range of a Duration in seconds, as specified in
	// google/protobuf/duration.proto. This is about 10,000 years in seconds.
	maxSeconds	= int64(10000 * 365.25 * 24 * 60 * 60)
	minSeconds	= -maxSeconds
)

// validateDuration determines whether the Duration is valid according to the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:50
// definition in google/protobuf/duration.proto. A valid Duration
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:50
// may still be too large to fit into a time.Duration (the range of Duration
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:50
// is about 10,000 years, and the range of time.Duration is about 290).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:54
func validateDuration(d *duration) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:54
	_go_fuzz_dep_.CoverTab[107958]++
											if d == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:55
		_go_fuzz_dep_.CoverTab[107963]++
												return errors.New("duration: nil Duration")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:56
		// _ = "end of CoverTab[107963]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:57
		_go_fuzz_dep_.CoverTab[107964]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:57
		// _ = "end of CoverTab[107964]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:57
	// _ = "end of CoverTab[107958]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:57
	_go_fuzz_dep_.CoverTab[107959]++
											if d.Seconds < minSeconds || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:58
		_go_fuzz_dep_.CoverTab[107965]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:58
		return d.Seconds > maxSeconds
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:58
		// _ = "end of CoverTab[107965]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:58
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:58
		_go_fuzz_dep_.CoverTab[107966]++
												return fmt.Errorf("duration: %#v: seconds out of range", d)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:59
		// _ = "end of CoverTab[107966]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:60
		_go_fuzz_dep_.CoverTab[107967]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:60
		// _ = "end of CoverTab[107967]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:60
	// _ = "end of CoverTab[107959]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:60
	_go_fuzz_dep_.CoverTab[107960]++
											if d.Nanos <= -1e9 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:61
		_go_fuzz_dep_.CoverTab[107968]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:61
		return d.Nanos >= 1e9
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:61
		// _ = "end of CoverTab[107968]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:61
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:61
		_go_fuzz_dep_.CoverTab[107969]++
												return fmt.Errorf("duration: %#v: nanos out of range", d)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:62
		// _ = "end of CoverTab[107969]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:63
		_go_fuzz_dep_.CoverTab[107970]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:63
		// _ = "end of CoverTab[107970]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:63
	// _ = "end of CoverTab[107960]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:63
	_go_fuzz_dep_.CoverTab[107961]++

											if (d.Seconds < 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
		_go_fuzz_dep_.CoverTab[107971]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
		return d.Nanos > 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
		// _ = "end of CoverTab[107971]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
		_go_fuzz_dep_.CoverTab[107972]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
		return (d.Seconds > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
			_go_fuzz_dep_.CoverTab[107973]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
			return d.Nanos < 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
			// _ = "end of CoverTab[107973]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
		}())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
		// _ = "end of CoverTab[107972]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:65
		_go_fuzz_dep_.CoverTab[107974]++
												return fmt.Errorf("duration: %#v: seconds and nanos have different signs", d)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:66
		// _ = "end of CoverTab[107974]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:67
		_go_fuzz_dep_.CoverTab[107975]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:67
		// _ = "end of CoverTab[107975]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:67
	// _ = "end of CoverTab[107961]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:67
	_go_fuzz_dep_.CoverTab[107962]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:68
	// _ = "end of CoverTab[107962]"
}

// DurationFromProto converts a Duration to a time.Duration. DurationFromProto
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:71
// returns an error if the Duration is invalid or is too large to be
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:71
// represented in a time.Duration.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:74
func durationFromProto(p *duration) (time.Duration, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:74
	_go_fuzz_dep_.CoverTab[107976]++
											if err := validateDuration(p); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:75
		_go_fuzz_dep_.CoverTab[107980]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:76
		// _ = "end of CoverTab[107980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:77
		_go_fuzz_dep_.CoverTab[107981]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:77
		// _ = "end of CoverTab[107981]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:77
	// _ = "end of CoverTab[107976]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:77
	_go_fuzz_dep_.CoverTab[107977]++
											d := time.Duration(p.Seconds) * time.Second
											if int64(d/time.Second) != p.Seconds {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:79
		_go_fuzz_dep_.CoverTab[107982]++
												return 0, fmt.Errorf("duration: %#v is out of range for time.Duration", p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:80
		// _ = "end of CoverTab[107982]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:81
		_go_fuzz_dep_.CoverTab[107983]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:81
		// _ = "end of CoverTab[107983]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:81
	// _ = "end of CoverTab[107977]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:81
	_go_fuzz_dep_.CoverTab[107978]++
											if p.Nanos != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:82
		_go_fuzz_dep_.CoverTab[107984]++
												d += time.Duration(p.Nanos)
												if (d < 0) != (p.Nanos < 0) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:84
			_go_fuzz_dep_.CoverTab[107985]++
													return 0, fmt.Errorf("duration: %#v is out of range for time.Duration", p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:85
			// _ = "end of CoverTab[107985]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:86
			_go_fuzz_dep_.CoverTab[107986]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:86
			// _ = "end of CoverTab[107986]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:86
		// _ = "end of CoverTab[107984]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:87
		_go_fuzz_dep_.CoverTab[107987]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:87
		// _ = "end of CoverTab[107987]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:87
	// _ = "end of CoverTab[107978]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:87
	_go_fuzz_dep_.CoverTab[107979]++
											return d, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:88
	// _ = "end of CoverTab[107979]"
}

// DurationProto converts a time.Duration to a Duration.
func durationProto(d time.Duration) *duration {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:92
	_go_fuzz_dep_.CoverTab[107988]++
											nanos := d.Nanoseconds()
											secs := nanos / 1e9
											nanos -= secs * 1e9
											return &duration{
		Seconds:	secs,
		Nanos:		int32(nanos),
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:99
	// _ = "end of CoverTab[107988]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:100
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/duration.go:100
var _ = _go_fuzz_dep_.CoverTab