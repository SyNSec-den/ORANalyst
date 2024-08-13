// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:5
package ptypes

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:5
)

import (
	"errors"
	"fmt"
	"time"

	durationpb "github.com/golang/protobuf/ptypes/duration"
)

// Range of google.protobuf.Duration as specified in duration.proto.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:15
// This is about 10,000 years in seconds.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:17
const (
	maxSeconds	= int64(10000 * 365.25 * 24 * 60 * 60)
	minSeconds	= -maxSeconds
)

// Duration converts a durationpb.Duration to a time.Duration.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:22
// Duration returns an error if dur is invalid or overflows a time.Duration.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:22
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:22
// Deprecated: Call the dur.AsDuration and dur.CheckValid methods instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:26
func Duration(dur *durationpb.Duration) (time.Duration, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:26
	_go_fuzz_dep_.CoverTab[68086]++
												if err := validateDuration(dur); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:27
		_go_fuzz_dep_.CoverTab[68090]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:28
		// _ = "end of CoverTab[68090]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:29
		_go_fuzz_dep_.CoverTab[68091]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:29
		// _ = "end of CoverTab[68091]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:29
	// _ = "end of CoverTab[68086]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:29
	_go_fuzz_dep_.CoverTab[68087]++
												d := time.Duration(dur.Seconds) * time.Second
												if int64(d/time.Second) != dur.Seconds {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:31
		_go_fuzz_dep_.CoverTab[68092]++
													return 0, fmt.Errorf("duration: %v is out of range for time.Duration", dur)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:32
		// _ = "end of CoverTab[68092]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:33
		_go_fuzz_dep_.CoverTab[68093]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:33
		// _ = "end of CoverTab[68093]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:33
	// _ = "end of CoverTab[68087]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:33
	_go_fuzz_dep_.CoverTab[68088]++
												if dur.Nanos != 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:34
		_go_fuzz_dep_.CoverTab[68094]++
													d += time.Duration(dur.Nanos) * time.Nanosecond
													if (d < 0) != (dur.Nanos < 0) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:36
			_go_fuzz_dep_.CoverTab[68095]++
														return 0, fmt.Errorf("duration: %v is out of range for time.Duration", dur)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:37
			// _ = "end of CoverTab[68095]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:38
			_go_fuzz_dep_.CoverTab[68096]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:38
			// _ = "end of CoverTab[68096]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:38
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:38
		// _ = "end of CoverTab[68094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:39
		_go_fuzz_dep_.CoverTab[68097]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:39
		// _ = "end of CoverTab[68097]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:39
	// _ = "end of CoverTab[68088]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:39
	_go_fuzz_dep_.CoverTab[68089]++
												return d, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:40
	// _ = "end of CoverTab[68089]"
}

// DurationProto converts a time.Duration to a durationpb.Duration.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:43
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:43
// Deprecated: Call the durationpb.New function instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:46
func DurationProto(d time.Duration) *durationpb.Duration {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:46
	_go_fuzz_dep_.CoverTab[68098]++
												nanos := d.Nanoseconds()
												secs := nanos / 1e9
												nanos -= secs * 1e9
												return &durationpb.Duration{
		Seconds:	int64(secs),
		Nanos:		int32(nanos),
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:53
	// _ = "end of CoverTab[68098]"
}

// validateDuration determines whether the durationpb.Duration is valid
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:56
// according to the definition in google/protobuf/duration.proto.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:56
// A valid durpb.Duration may still be too large to fit into a time.Duration
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:56
// Note that the range of durationpb.Duration is about 10,000 years,
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:56
// while the range of time.Duration is about 290 years.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:61
func validateDuration(dur *durationpb.Duration) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:61
	_go_fuzz_dep_.CoverTab[68099]++
												if dur == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:62
		_go_fuzz_dep_.CoverTab[68104]++
													return errors.New("duration: nil Duration")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:63
		// _ = "end of CoverTab[68104]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:64
		_go_fuzz_dep_.CoverTab[68105]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:64
		// _ = "end of CoverTab[68105]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:64
	// _ = "end of CoverTab[68099]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:64
	_go_fuzz_dep_.CoverTab[68100]++
												if dur.Seconds < minSeconds || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:65
		_go_fuzz_dep_.CoverTab[68106]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:65
		return dur.Seconds > maxSeconds
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:65
		// _ = "end of CoverTab[68106]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:65
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:65
		_go_fuzz_dep_.CoverTab[68107]++
													return fmt.Errorf("duration: %v: seconds out of range", dur)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:66
		// _ = "end of CoverTab[68107]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:67
		_go_fuzz_dep_.CoverTab[68108]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:67
		// _ = "end of CoverTab[68108]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:67
	// _ = "end of CoverTab[68100]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:67
	_go_fuzz_dep_.CoverTab[68101]++
												if dur.Nanos <= -1e9 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:68
		_go_fuzz_dep_.CoverTab[68109]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:68
		return dur.Nanos >= 1e9
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:68
		// _ = "end of CoverTab[68109]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:68
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:68
		_go_fuzz_dep_.CoverTab[68110]++
													return fmt.Errorf("duration: %v: nanos out of range", dur)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:69
		// _ = "end of CoverTab[68110]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:70
		_go_fuzz_dep_.CoverTab[68111]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:70
		// _ = "end of CoverTab[68111]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:70
	// _ = "end of CoverTab[68101]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:70
	_go_fuzz_dep_.CoverTab[68102]++

												if (dur.Seconds < 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
		_go_fuzz_dep_.CoverTab[68112]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
		return dur.Nanos > 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
		// _ = "end of CoverTab[68112]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
		_go_fuzz_dep_.CoverTab[68113]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
		return (dur.Seconds > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
			_go_fuzz_dep_.CoverTab[68114]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
			return dur.Nanos < 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
			// _ = "end of CoverTab[68114]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
		}())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
		// _ = "end of CoverTab[68113]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:72
		_go_fuzz_dep_.CoverTab[68115]++
													return fmt.Errorf("duration: %v: seconds and nanos have different signs", dur)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:73
		// _ = "end of CoverTab[68115]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:74
		_go_fuzz_dep_.CoverTab[68116]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:74
		// _ = "end of CoverTab[68116]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:74
	// _ = "end of CoverTab[68102]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:74
	_go_fuzz_dep_.CoverTab[68103]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:75
	// _ = "end of CoverTab[68103]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:76
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/duration.go:76
var _ = _go_fuzz_dep_.CoverTab
