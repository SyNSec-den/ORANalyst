// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:5
package ptypes

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:5
)

import (
	"errors"
	"fmt"
	"time"

	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
)

// Range of google.protobuf.Duration as specified in timestamp.proto.
const (
	// Seconds field of the earliest valid Timestamp.
	// This is time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	minValidSeconds	= -62135596800
	// Seconds field just after the latest valid Timestamp.
	// This is time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	maxValidSeconds	= 253402300800
)

// Timestamp converts a timestamppb.Timestamp to a time.Time.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
// It returns an error if the argument is invalid.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
// Unlike most Go functions, if Timestamp returns an error, the first return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
// value is not the zero time.Time. Instead, it is the value obtained from the
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
// time.Unix function when passed the contents of the Timestamp, in the UTC
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
// locale. This may or may not be a meaningful time; many invalid Timestamps
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
// do map to valid time.Times.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
// A nil Timestamp returns an error. The first return value in that case is
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
// undefined.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:25
// Deprecated: Call the ts.AsTime and ts.CheckValid methods instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:38
func Timestamp(ts *timestamppb.Timestamp) (time.Time, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:38
	_go_fuzz_dep_.CoverTab[68117]++
	// Don't return the zero value on error, because corresponds to a valid
	// timestamp. Instead return whatever time.Unix gives us.
	var t time.Time
	if ts == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:42
		_go_fuzz_dep_.CoverTab[68119]++
													t = time.Unix(0, 0).UTC()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:43
		// _ = "end of CoverTab[68119]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:44
		_go_fuzz_dep_.CoverTab[68120]++
													t = time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:45
		// _ = "end of CoverTab[68120]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:46
	// _ = "end of CoverTab[68117]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:46
	_go_fuzz_dep_.CoverTab[68118]++
												return t, validateTimestamp(ts)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:47
	// _ = "end of CoverTab[68118]"
}

// TimestampNow returns a google.protobuf.Timestamp for the current time.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:50
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:50
// Deprecated: Call the timestamppb.Now function instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:53
func TimestampNow() *timestamppb.Timestamp {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:53
	_go_fuzz_dep_.CoverTab[68121]++
												ts, err := TimestampProto(time.Now())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:55
		_go_fuzz_dep_.CoverTab[68123]++
													panic("ptypes: time.Now() out of Timestamp range")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:56
		// _ = "end of CoverTab[68123]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:57
		_go_fuzz_dep_.CoverTab[68124]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:57
		// _ = "end of CoverTab[68124]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:57
	// _ = "end of CoverTab[68121]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:57
	_go_fuzz_dep_.CoverTab[68122]++
												return ts
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:58
	// _ = "end of CoverTab[68122]"
}

// TimestampProto converts the time.Time to a google.protobuf.Timestamp proto.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:61
// It returns an error if the resulting Timestamp is invalid.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:61
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:61
// Deprecated: Call the timestamppb.New function instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:65
func TimestampProto(t time.Time) (*timestamppb.Timestamp, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:65
	_go_fuzz_dep_.CoverTab[68125]++
												ts := &timestamppb.Timestamp{
		Seconds:	t.Unix(),
		Nanos:		int32(t.Nanosecond()),
	}
	if err := validateTimestamp(ts); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:70
		_go_fuzz_dep_.CoverTab[68127]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:71
		// _ = "end of CoverTab[68127]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:72
		_go_fuzz_dep_.CoverTab[68128]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:72
		// _ = "end of CoverTab[68128]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:72
	// _ = "end of CoverTab[68125]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:72
	_go_fuzz_dep_.CoverTab[68126]++
												return ts, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:73
	// _ = "end of CoverTab[68126]"
}

// TimestampString returns the RFC 3339 string for valid Timestamps.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:76
// For invalid Timestamps, it returns an error message in parentheses.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:76
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:76
// Deprecated: Call the ts.AsTime method instead,
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:76
// followed by a call to the Format method on the time.Time value.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:81
func TimestampString(ts *timestamppb.Timestamp) string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:81
	_go_fuzz_dep_.CoverTab[68129]++
												t, err := Timestamp(ts)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:83
		_go_fuzz_dep_.CoverTab[68131]++
													return fmt.Sprintf("(%v)", err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:84
		// _ = "end of CoverTab[68131]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:85
		_go_fuzz_dep_.CoverTab[68132]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:85
		// _ = "end of CoverTab[68132]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:85
	// _ = "end of CoverTab[68129]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:85
	_go_fuzz_dep_.CoverTab[68130]++
												return t.Format(time.RFC3339Nano)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:86
	// _ = "end of CoverTab[68130]"
}

// validateTimestamp determines whether a Timestamp is valid.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:89
// A valid timestamp represents a time in the range [0001-01-01, 10000-01-01)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:89
// and has a Nanos field in the range [0, 1e9).
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:89
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:89
// If the Timestamp is valid, validateTimestamp returns nil.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:89
// Otherwise, it returns an error that describes the problem.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:89
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:89
// Every valid Timestamp can be represented by a time.Time,
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:89
// but the converse is not true.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:98
func validateTimestamp(ts *timestamppb.Timestamp) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:98
	_go_fuzz_dep_.CoverTab[68133]++
												if ts == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:99
		_go_fuzz_dep_.CoverTab[68138]++
													return errors.New("timestamp: nil Timestamp")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:100
		// _ = "end of CoverTab[68138]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:101
		_go_fuzz_dep_.CoverTab[68139]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:101
		// _ = "end of CoverTab[68139]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:101
	// _ = "end of CoverTab[68133]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:101
	_go_fuzz_dep_.CoverTab[68134]++
												if ts.Seconds < minValidSeconds {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:102
		_go_fuzz_dep_.CoverTab[68140]++
													return fmt.Errorf("timestamp: %v before 0001-01-01", ts)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:103
		// _ = "end of CoverTab[68140]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:104
		_go_fuzz_dep_.CoverTab[68141]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:104
		// _ = "end of CoverTab[68141]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:104
	// _ = "end of CoverTab[68134]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:104
	_go_fuzz_dep_.CoverTab[68135]++
												if ts.Seconds >= maxValidSeconds {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:105
		_go_fuzz_dep_.CoverTab[68142]++
													return fmt.Errorf("timestamp: %v after 10000-01-01", ts)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:106
		// _ = "end of CoverTab[68142]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:107
		_go_fuzz_dep_.CoverTab[68143]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:107
		// _ = "end of CoverTab[68143]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:107
	// _ = "end of CoverTab[68135]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:107
	_go_fuzz_dep_.CoverTab[68136]++
												if ts.Nanos < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:108
		_go_fuzz_dep_.CoverTab[68144]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:108
		return ts.Nanos >= 1e9
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:108
		// _ = "end of CoverTab[68144]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:108
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:108
		_go_fuzz_dep_.CoverTab[68145]++
													return fmt.Errorf("timestamp: %v: nanos not in range [0, 1e9)", ts)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:109
		// _ = "end of CoverTab[68145]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:110
		_go_fuzz_dep_.CoverTab[68146]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:110
		// _ = "end of CoverTab[68146]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:110
	// _ = "end of CoverTab[68136]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:110
	_go_fuzz_dep_.CoverTab[68137]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:111
	// _ = "end of CoverTab[68137]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:112
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/timestamp.go:112
var _ = _go_fuzz_dep_.CoverTab
