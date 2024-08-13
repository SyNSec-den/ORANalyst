// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:8
)

import "time"

// TimespecToNsec returns the time stored in ts as nanoseconds.
func TimespecToNsec(ts Timespec) int64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:13
	_go_fuzz_dep_.CoverTab[46856]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:13
	return ts.Nano()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:13
	// _ = "end of CoverTab[46856]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:13
}

// NsecToTimespec converts a number of nanoseconds into a Timespec.
func NsecToTimespec(nsec int64) Timespec {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:16
	_go_fuzz_dep_.CoverTab[46857]++
											sec := nsec / 1e9
											nsec = nsec % 1e9
											if nsec < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:19
		_go_fuzz_dep_.CoverTab[46859]++
												nsec += 1e9
												sec--
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:21
		// _ = "end of CoverTab[46859]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:22
		_go_fuzz_dep_.CoverTab[46860]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:22
		// _ = "end of CoverTab[46860]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:22
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:22
	// _ = "end of CoverTab[46857]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:22
	_go_fuzz_dep_.CoverTab[46858]++
											return setTimespec(sec, nsec)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:23
	// _ = "end of CoverTab[46858]"
}

// TimeToTimespec converts t into a Timespec.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:26
// On some 32-bit systems the range of valid Timespec values are smaller
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:26
// than that of time.Time values.  So if t is out of the valid range of
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:26
// Timespec, it returns a zero Timespec and ERANGE.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:30
func TimeToTimespec(t time.Time) (Timespec, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:30
	_go_fuzz_dep_.CoverTab[46861]++
											sec := t.Unix()
											nsec := int64(t.Nanosecond())
											ts := setTimespec(sec, nsec)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:38
	if int64(ts.Sec) != sec {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:38
		_go_fuzz_dep_.CoverTab[46863]++
												return Timespec{}, ERANGE
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:39
		// _ = "end of CoverTab[46863]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:40
		_go_fuzz_dep_.CoverTab[46864]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:40
		// _ = "end of CoverTab[46864]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:40
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:40
	// _ = "end of CoverTab[46861]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:40
	_go_fuzz_dep_.CoverTab[46862]++
											return ts, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:41
	// _ = "end of CoverTab[46862]"
}

// TimevalToNsec returns the time stored in tv as nanoseconds.
func TimevalToNsec(tv Timeval) int64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:45
	_go_fuzz_dep_.CoverTab[46865]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:45
	return tv.Nano()
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:45
	// _ = "end of CoverTab[46865]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:45
}

// NsecToTimeval converts a number of nanoseconds into a Timeval.
func NsecToTimeval(nsec int64) Timeval {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:48
	_go_fuzz_dep_.CoverTab[46866]++
											nsec += 999
											usec := nsec % 1e9 / 1e3
											sec := nsec / 1e9
											if usec < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:52
		_go_fuzz_dep_.CoverTab[46868]++
												usec += 1e6
												sec--
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:54
		// _ = "end of CoverTab[46868]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:55
		_go_fuzz_dep_.CoverTab[46869]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:55
		// _ = "end of CoverTab[46869]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:55
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:55
	// _ = "end of CoverTab[46866]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:55
	_go_fuzz_dep_.CoverTab[46867]++
											return setTimeval(sec, usec)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:56
	// _ = "end of CoverTab[46867]"
}

// Unix returns the time stored in ts as seconds plus nanoseconds.
func (ts *Timespec) Unix() (sec int64, nsec int64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:60
	_go_fuzz_dep_.CoverTab[46870]++
											return int64(ts.Sec), int64(ts.Nsec)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:61
	// _ = "end of CoverTab[46870]"
}

// Unix returns the time stored in tv as seconds plus nanoseconds.
func (tv *Timeval) Unix() (sec int64, nsec int64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:65
	_go_fuzz_dep_.CoverTab[46871]++
											return int64(tv.Sec), int64(tv.Usec) * 1000
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:66
	// _ = "end of CoverTab[46871]"
}

// Nano returns the time stored in ts as nanoseconds.
func (ts *Timespec) Nano() int64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:70
	_go_fuzz_dep_.CoverTab[46872]++
											return int64(ts.Sec)*1e9 + int64(ts.Nsec)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:71
	// _ = "end of CoverTab[46872]"
}

// Nano returns the time stored in tv as nanoseconds.
func (tv *Timeval) Nano() int64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:75
	_go_fuzz_dep_.CoverTab[46873]++
											return int64(tv.Sec)*1e9 + int64(tv.Usec)*1000
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:76
	// _ = "end of CoverTab[46873]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:77
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/timestruct.go:77
var _ = _go_fuzz_dep_.CoverTab
