// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:5
)

import (
	"encoding/binary"
	"sync"
	"time"
)

// A Time represents a time as the number of 100's of nanoseconds since 15 Oct
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:13
// 1582.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:15
type Time int64

const (
	lillian		= 2299160		// Julian day of 15 Oct 1582
	unix		= 2440587		// Julian day of 1 Jan 1970
	epoch		= unix - lillian	// Days between epochs
	g1582		= epoch * 86400		// seconds between epochs
	g1582ns100	= g1582 * 10000000	// 100s of a nanoseconds between epochs
)

var (
	timeMu		sync.Mutex
	lasttime	uint64	// last time we returned
	clockSeq	uint16	// clock sequence for this run

	timeNow	= time.Now	// for testing
)

// UnixTime converts t the number of seconds and nanoseconds using the Unix
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:33
// epoch of 1 Jan 1970.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:35
func (t Time) UnixTime() (sec, nsec int64) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:35
	_go_fuzz_dep_.CoverTab[179413]++
										sec = int64(t - g1582ns100)
										nsec = (sec % 10000000) * 100
										sec /= 10000000
										return sec, nsec
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:39
	// _ = "end of CoverTab[179413]"
}

// GetTime returns the current Time (100s of nanoseconds since 15 Oct 1582) and
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:42
// clock sequence as well as adjusting the clock sequence as needed.  An error
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:42
// is returned if the current time cannot be determined.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:45
func GetTime() (Time, uint16, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:45
	_go_fuzz_dep_.CoverTab[179414]++
										defer timeMu.Unlock()
										timeMu.Lock()
										return getTime()
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:48
	// _ = "end of CoverTab[179414]"
}

func getTime() (Time, uint16, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:51
	_go_fuzz_dep_.CoverTab[179415]++
										t := timeNow()

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:55
	if clockSeq == 0 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:55
		_go_fuzz_dep_.CoverTab[179418]++
											setClockSequence(-1)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:56
		// _ = "end of CoverTab[179418]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:57
		_go_fuzz_dep_.CoverTab[179419]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:57
		// _ = "end of CoverTab[179419]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:57
	// _ = "end of CoverTab[179415]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:57
	_go_fuzz_dep_.CoverTab[179416]++
										now := uint64(t.UnixNano()/100) + g1582ns100

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:62
	if now <= lasttime {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:62
		_go_fuzz_dep_.CoverTab[179420]++
											clockSeq = ((clockSeq + 1) & 0x3fff) | 0x8000
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:63
		// _ = "end of CoverTab[179420]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:64
		_go_fuzz_dep_.CoverTab[179421]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:64
		// _ = "end of CoverTab[179421]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:64
	// _ = "end of CoverTab[179416]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:64
	_go_fuzz_dep_.CoverTab[179417]++
										lasttime = now
										return Time(now), clockSeq, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:66
	// _ = "end of CoverTab[179417]"
}

// ClockSequence returns the current clock sequence, generating one if not
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:69
// already set.  The clock sequence is only used for Version 1 UUIDs.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:69
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:69
// The uuid package does not use global static storage for the clock sequence or
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:69
// the last time a UUID was generated.  Unless SetClockSequence is used, a new
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:69
// random clock sequence is generated the first time a clock sequence is
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:69
// requested by ClockSequence, GetTime, or NewUUID.  (section 4.2.1.1)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:76
func ClockSequence() int {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:76
	_go_fuzz_dep_.CoverTab[179422]++
										defer timeMu.Unlock()
										timeMu.Lock()
										return clockSequence()
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:79
	// _ = "end of CoverTab[179422]"
}

func clockSequence() int {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:82
	_go_fuzz_dep_.CoverTab[179423]++
										if clockSeq == 0 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:83
		_go_fuzz_dep_.CoverTab[179425]++
											setClockSequence(-1)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:84
		// _ = "end of CoverTab[179425]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:85
		_go_fuzz_dep_.CoverTab[179426]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:85
		// _ = "end of CoverTab[179426]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:85
	// _ = "end of CoverTab[179423]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:85
	_go_fuzz_dep_.CoverTab[179424]++
										return int(clockSeq & 0x3fff)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:86
	// _ = "end of CoverTab[179424]"
}

// SetClockSequence sets the clock sequence to the lower 14 bits of seq.  Setting to
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:89
// -1 causes a new sequence to be generated.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:91
func SetClockSequence(seq int) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:91
	_go_fuzz_dep_.CoverTab[179427]++
										defer timeMu.Unlock()
										timeMu.Lock()
										setClockSequence(seq)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:94
	// _ = "end of CoverTab[179427]"
}

func setClockSequence(seq int) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:97
	_go_fuzz_dep_.CoverTab[179428]++
										if seq == -1 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:98
		_go_fuzz_dep_.CoverTab[179430]++
											var b [2]byte
											randomBits(b[:])
											seq = int(b[0])<<8 | int(b[1])
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:101
		// _ = "end of CoverTab[179430]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:102
		_go_fuzz_dep_.CoverTab[179431]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:102
		// _ = "end of CoverTab[179431]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:102
	// _ = "end of CoverTab[179428]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:102
	_go_fuzz_dep_.CoverTab[179429]++
										oldSeq := clockSeq
										clockSeq = uint16(seq&0x3fff) | 0x8000
										if oldSeq != clockSeq {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:105
		_go_fuzz_dep_.CoverTab[179432]++
											lasttime = 0
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:106
		// _ = "end of CoverTab[179432]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:107
		_go_fuzz_dep_.CoverTab[179433]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:107
		// _ = "end of CoverTab[179433]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:107
	// _ = "end of CoverTab[179429]"
}

// Time returns the time in 100s of nanoseconds since 15 Oct 1582 encoded in
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:110
// uuid.  The time is only defined for version 1 and 2 UUIDs.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:112
func (uuid UUID) Time() Time {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:112
	_go_fuzz_dep_.CoverTab[179434]++
										time := int64(binary.BigEndian.Uint32(uuid[0:4]))
										time |= int64(binary.BigEndian.Uint16(uuid[4:6])) << 32
										time |= int64(binary.BigEndian.Uint16(uuid[6:8])&0xfff) << 48
										return Time(time)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:116
	// _ = "end of CoverTab[179434]"
}

// ClockSequence returns the clock sequence encoded in uuid.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:119
// The clock sequence is only well defined for version 1 and 2 UUIDs.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:121
func (uuid UUID) ClockSequence() int {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:121
	_go_fuzz_dep_.CoverTab[179435]++
										return int(binary.BigEndian.Uint16(uuid[8:10])) & 0x3fff
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:122
	// _ = "end of CoverTab[179435]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:123
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/time.go:123
var _ = _go_fuzz_dep_.CoverTab
