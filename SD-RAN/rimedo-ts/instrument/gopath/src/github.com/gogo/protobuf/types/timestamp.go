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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:32
package types

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:32
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:36
import (
	"errors"
	"fmt"
	"time"
)

const (
	// Seconds field of the earliest valid Timestamp.
	// This is time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	minValidSeconds	= -62135596800
	// Seconds field just after the latest valid Timestamp.
	// This is time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	maxValidSeconds	= 253402300800
)

// validateTimestamp determines whether a Timestamp is valid.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:51
// A valid timestamp represents a time in the range
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:51
// [0001-01-01, 10000-01-01) and has a Nanos field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:51
// in the range [0, 1e9).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:51
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:51
// If the Timestamp is valid, validateTimestamp returns nil.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:51
// Otherwise, it returns an error that describes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:51
// the problem.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:51
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:51
// Every valid Timestamp can be represented by a time.Time, but the converse is not true.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:61
func validateTimestamp(ts *Timestamp) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:61
	_go_fuzz_dep_.CoverTab[138070]++
												if ts == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:62
		_go_fuzz_dep_.CoverTab[138075]++
													return errors.New("timestamp: nil Timestamp")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:63
		// _ = "end of CoverTab[138075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:64
		_go_fuzz_dep_.CoverTab[138076]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:64
		// _ = "end of CoverTab[138076]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:64
	// _ = "end of CoverTab[138070]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:64
	_go_fuzz_dep_.CoverTab[138071]++
												if ts.Seconds < minValidSeconds {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:65
		_go_fuzz_dep_.CoverTab[138077]++
													return fmt.Errorf("timestamp: %#v before 0001-01-01", ts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:66
		// _ = "end of CoverTab[138077]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:67
		_go_fuzz_dep_.CoverTab[138078]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:67
		// _ = "end of CoverTab[138078]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:67
	// _ = "end of CoverTab[138071]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:67
	_go_fuzz_dep_.CoverTab[138072]++
												if ts.Seconds >= maxValidSeconds {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:68
		_go_fuzz_dep_.CoverTab[138079]++
													return fmt.Errorf("timestamp: %#v after 10000-01-01", ts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:69
		// _ = "end of CoverTab[138079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:70
		_go_fuzz_dep_.CoverTab[138080]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:70
		// _ = "end of CoverTab[138080]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:70
	// _ = "end of CoverTab[138072]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:70
	_go_fuzz_dep_.CoverTab[138073]++
												if ts.Nanos < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:71
		_go_fuzz_dep_.CoverTab[138081]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:71
		return ts.Nanos >= 1e9
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:71
		// _ = "end of CoverTab[138081]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:71
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:71
		_go_fuzz_dep_.CoverTab[138082]++
													return fmt.Errorf("timestamp: %#v: nanos not in range [0, 1e9)", ts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:72
		// _ = "end of CoverTab[138082]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:73
		_go_fuzz_dep_.CoverTab[138083]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:73
		// _ = "end of CoverTab[138083]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:73
	// _ = "end of CoverTab[138073]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:73
	_go_fuzz_dep_.CoverTab[138074]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:74
	// _ = "end of CoverTab[138074]"
}

// TimestampFromProto converts a google.protobuf.Timestamp proto to a time.Time.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
// It returns an error if the argument is invalid.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
// Unlike most Go functions, if Timestamp returns an error, the first return value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
// is not the zero time.Time. Instead, it is the value obtained from the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
// time.Unix function when passed the contents of the Timestamp, in the UTC
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
// locale. This may or may not be a meaningful time; many invalid Timestamps
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
// do map to valid time.Times.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
// A nil Timestamp returns an error. The first return value in that case is
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:77
// undefined.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:88
func TimestampFromProto(ts *Timestamp) (time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:88
	_go_fuzz_dep_.CoverTab[138084]++
	// Don't return the zero value on error, because corresponds to a valid
	// timestamp. Instead return whatever time.Unix gives us.
	var t time.Time
	if ts == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:92
		_go_fuzz_dep_.CoverTab[138086]++
													t = time.Unix(0, 0).UTC()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:93
		// _ = "end of CoverTab[138086]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:94
		_go_fuzz_dep_.CoverTab[138087]++
													t = time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:95
		// _ = "end of CoverTab[138087]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:96
	// _ = "end of CoverTab[138084]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:96
	_go_fuzz_dep_.CoverTab[138085]++
												return t, validateTimestamp(ts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:97
	// _ = "end of CoverTab[138085]"
}

// TimestampNow returns a google.protobuf.Timestamp for the current time.
func TimestampNow() *Timestamp {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:101
	_go_fuzz_dep_.CoverTab[138088]++
												ts, err := TimestampProto(time.Now())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:103
		_go_fuzz_dep_.CoverTab[138090]++
													panic("ptypes: time.Now() out of Timestamp range")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:104
		// _ = "end of CoverTab[138090]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:105
		_go_fuzz_dep_.CoverTab[138091]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:105
		// _ = "end of CoverTab[138091]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:105
	// _ = "end of CoverTab[138088]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:105
	_go_fuzz_dep_.CoverTab[138089]++
												return ts
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:106
	// _ = "end of CoverTab[138089]"
}

// TimestampProto converts the time.Time to a google.protobuf.Timestamp proto.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:109
// It returns an error if the resulting Timestamp is invalid.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:111
func TimestampProto(t time.Time) (*Timestamp, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:111
	_go_fuzz_dep_.CoverTab[138092]++
												ts := &Timestamp{
		Seconds:	t.Unix(),
		Nanos:		int32(t.Nanosecond()),
	}
	if err := validateTimestamp(ts); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:116
		_go_fuzz_dep_.CoverTab[138094]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:117
		// _ = "end of CoverTab[138094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:118
		_go_fuzz_dep_.CoverTab[138095]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:118
		// _ = "end of CoverTab[138095]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:118
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:118
	// _ = "end of CoverTab[138092]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:118
	_go_fuzz_dep_.CoverTab[138093]++
												return ts, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:119
	// _ = "end of CoverTab[138093]"
}

// TimestampString returns the RFC 3339 string for valid Timestamps. For invalid
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:122
// Timestamps, it returns an error message in parentheses.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:124
func TimestampString(ts *Timestamp) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:124
	_go_fuzz_dep_.CoverTab[138096]++
												t, err := TimestampFromProto(ts)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:126
		_go_fuzz_dep_.CoverTab[138098]++
													return fmt.Sprintf("(%v)", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:127
		// _ = "end of CoverTab[138098]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:128
		_go_fuzz_dep_.CoverTab[138099]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:128
		// _ = "end of CoverTab[138099]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:128
	// _ = "end of CoverTab[138096]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:128
	_go_fuzz_dep_.CoverTab[138097]++
												return t.Format(time.RFC3339Nano)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:129
	// _ = "end of CoverTab[138097]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:130
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/timestamp.go:130
var _ = _go_fuzz_dep_.CoverTab
