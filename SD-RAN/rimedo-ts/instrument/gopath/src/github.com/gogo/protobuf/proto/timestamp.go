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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:32
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:36
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
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:51
// A valid timestamp represents a time in the range
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:51
// [0001-01-01, 10000-01-01) and has a Nanos field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:51
// in the range [0, 1e9).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:51
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:51
// If the Timestamp is valid, validateTimestamp returns nil.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:51
// Otherwise, it returns an error that describes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:51
// the problem.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:51
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:51
// Every valid Timestamp can be represented by a time.Time, but the converse is not true.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:61
func validateTimestamp(ts *timestamp) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:61
	_go_fuzz_dep_.CoverTab[113531]++
												if ts == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:62
		_go_fuzz_dep_.CoverTab[113536]++
													return errors.New("timestamp: nil Timestamp")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:63
		// _ = "end of CoverTab[113536]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:64
		_go_fuzz_dep_.CoverTab[113537]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:64
		// _ = "end of CoverTab[113537]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:64
	// _ = "end of CoverTab[113531]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:64
	_go_fuzz_dep_.CoverTab[113532]++
												if ts.Seconds < minValidSeconds {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:65
		_go_fuzz_dep_.CoverTab[113538]++
													return fmt.Errorf("timestamp: %#v before 0001-01-01", ts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:66
		// _ = "end of CoverTab[113538]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:67
		_go_fuzz_dep_.CoverTab[113539]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:67
		// _ = "end of CoverTab[113539]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:67
	// _ = "end of CoverTab[113532]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:67
	_go_fuzz_dep_.CoverTab[113533]++
												if ts.Seconds >= maxValidSeconds {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:68
		_go_fuzz_dep_.CoverTab[113540]++
													return fmt.Errorf("timestamp: %#v after 10000-01-01", ts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:69
		// _ = "end of CoverTab[113540]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:70
		_go_fuzz_dep_.CoverTab[113541]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:70
		// _ = "end of CoverTab[113541]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:70
	// _ = "end of CoverTab[113533]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:70
	_go_fuzz_dep_.CoverTab[113534]++
												if ts.Nanos < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:71
		_go_fuzz_dep_.CoverTab[113542]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:71
		return ts.Nanos >= 1e9
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:71
		// _ = "end of CoverTab[113542]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:71
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:71
		_go_fuzz_dep_.CoverTab[113543]++
													return fmt.Errorf("timestamp: %#v: nanos not in range [0, 1e9)", ts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:72
		// _ = "end of CoverTab[113543]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:73
		_go_fuzz_dep_.CoverTab[113544]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:73
		// _ = "end of CoverTab[113544]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:73
	// _ = "end of CoverTab[113534]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:73
	_go_fuzz_dep_.CoverTab[113535]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:74
	// _ = "end of CoverTab[113535]"
}

// TimestampFromProto converts a google.protobuf.Timestamp proto to a time.Time.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
// It returns an error if the argument is invalid.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
// Unlike most Go functions, if Timestamp returns an error, the first return value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
// is not the zero time.Time. Instead, it is the value obtained from the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
// time.Unix function when passed the contents of the Timestamp, in the UTC
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
// locale. This may or may not be a meaningful time; many invalid Timestamps
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
// do map to valid time.Times.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
// A nil Timestamp returns an error. The first return value in that case is
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:77
// undefined.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:88
func timestampFromProto(ts *timestamp) (time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:88
	_go_fuzz_dep_.CoverTab[113545]++
	// Don't return the zero value on error, because corresponds to a valid
	// timestamp. Instead return whatever time.Unix gives us.
	var t time.Time
	if ts == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:92
		_go_fuzz_dep_.CoverTab[113547]++
													t = time.Unix(0, 0).UTC()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:93
		// _ = "end of CoverTab[113547]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:94
		_go_fuzz_dep_.CoverTab[113548]++
													t = time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:95
		// _ = "end of CoverTab[113548]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:96
	// _ = "end of CoverTab[113545]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:96
	_go_fuzz_dep_.CoverTab[113546]++
												return t, validateTimestamp(ts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:97
	// _ = "end of CoverTab[113546]"
}

// TimestampProto converts the time.Time to a google.protobuf.Timestamp proto.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:100
// It returns an error if the resulting Timestamp is invalid.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:102
func timestampProto(t time.Time) (*timestamp, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:102
	_go_fuzz_dep_.CoverTab[113549]++
												seconds := t.Unix()
												nanos := int32(t.Sub(time.Unix(seconds, 0)))
												ts := &timestamp{
		Seconds:	seconds,
		Nanos:		nanos,
	}
	if err := validateTimestamp(ts); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:109
		_go_fuzz_dep_.CoverTab[113551]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:110
		// _ = "end of CoverTab[113551]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:111
		_go_fuzz_dep_.CoverTab[113552]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:111
		// _ = "end of CoverTab[113552]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:111
	// _ = "end of CoverTab[113549]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:111
	_go_fuzz_dep_.CoverTab[113550]++
												return ts, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:112
	// _ = "end of CoverTab[113550]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/timestamp.go:113
var _ = _go_fuzz_dep_.CoverTab
