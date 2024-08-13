// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 && linux
// +build amd64,linux

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:8
)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:27
func Lstat(path string, stat *Stat_t) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:27
	_go_fuzz_dep_.CoverTab[46610]++
												return Fstatat(AT_FDCWD, path, stat, AT_SYMLINK_NOFOLLOW)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:28
	// _ = "end of CoverTab[46610]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:38
func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:38
	_go_fuzz_dep_.CoverTab[46611]++
												var ts *Timespec
												if timeout != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:40
		_go_fuzz_dep_.CoverTab[46613]++
													ts = &Timespec{Sec: timeout.Sec, Nsec: timeout.Usec * 1000}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:41
		// _ = "end of CoverTab[46613]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:42
		_go_fuzz_dep_.CoverTab[46614]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:42
		// _ = "end of CoverTab[46614]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:42
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:42
	// _ = "end of CoverTab[46611]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:42
	_go_fuzz_dep_.CoverTab[46612]++
												return Pselect(nfd, r, w, e, ts, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:43
	// _ = "end of CoverTab[46612]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:52
func Stat(path string, stat *Stat_t) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:52
	_go_fuzz_dep_.CoverTab[46615]++

												return Fstatat(AT_FDCWD, path, stat, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:54
	// _ = "end of CoverTab[46615]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:80
func Gettimeofday(tv *Timeval) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:80
	_go_fuzz_dep_.CoverTab[46616]++
												errno := gettimeofday(tv)
												if errno != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:82
		_go_fuzz_dep_.CoverTab[46618]++
													return errno
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:83
		// _ = "end of CoverTab[46618]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:84
		_go_fuzz_dep_.CoverTab[46619]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:84
		// _ = "end of CoverTab[46619]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:84
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:84
	// _ = "end of CoverTab[46616]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:84
	_go_fuzz_dep_.CoverTab[46617]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:85
	// _ = "end of CoverTab[46617]"
}

func Time(t *Time_t) (tt Time_t, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:88
	_go_fuzz_dep_.CoverTab[46620]++
												var tv Timeval
												errno := gettimeofday(&tv)
												if errno != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:91
		_go_fuzz_dep_.CoverTab[46623]++
													return 0, errno
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:92
		// _ = "end of CoverTab[46623]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:93
		_go_fuzz_dep_.CoverTab[46624]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:93
		// _ = "end of CoverTab[46624]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:93
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:93
	// _ = "end of CoverTab[46620]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:93
	_go_fuzz_dep_.CoverTab[46621]++
												if t != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:94
		_go_fuzz_dep_.CoverTab[46625]++
													*t = Time_t(tv.Sec)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:95
		// _ = "end of CoverTab[46625]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:96
		_go_fuzz_dep_.CoverTab[46626]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:96
		// _ = "end of CoverTab[46626]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:96
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:96
	// _ = "end of CoverTab[46621]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:96
	_go_fuzz_dep_.CoverTab[46622]++
												return Time_t(tv.Sec), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:97
	// _ = "end of CoverTab[46622]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:103
func setTimespec(sec, nsec int64) Timespec {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:103
	_go_fuzz_dep_.CoverTab[46627]++
												return Timespec{Sec: sec, Nsec: nsec}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:104
	// _ = "end of CoverTab[46627]"
}

func setTimeval(sec, usec int64) Timeval {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:107
	_go_fuzz_dep_.CoverTab[46628]++
												return Timeval{Sec: sec, Usec: usec}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:108
	// _ = "end of CoverTab[46628]"
}

func (r *PtraceRegs) PC() uint64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:111
	_go_fuzz_dep_.CoverTab[46629]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:111
	return r.Rip
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:111
	// _ = "end of CoverTab[46629]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:111
}

func (r *PtraceRegs) SetPC(pc uint64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:113
	_go_fuzz_dep_.CoverTab[46630]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:113
	r.Rip = pc
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:113
	// _ = "end of CoverTab[46630]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:113
}

func (iov *Iovec) SetLen(length int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:115
	_go_fuzz_dep_.CoverTab[46631]++
												iov.Len = uint64(length)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:116
	// _ = "end of CoverTab[46631]"
}

func (msghdr *Msghdr) SetControllen(length int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:119
	_go_fuzz_dep_.CoverTab[46632]++
												msghdr.Controllen = uint64(length)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:120
	// _ = "end of CoverTab[46632]"
}

func (msghdr *Msghdr) SetIovlen(length int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:123
	_go_fuzz_dep_.CoverTab[46633]++
												msghdr.Iovlen = uint64(length)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:124
	// _ = "end of CoverTab[46633]"
}

func (cmsg *Cmsghdr) SetLen(length int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:127
	_go_fuzz_dep_.CoverTab[46634]++
												cmsg.Len = uint64(length)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:128
	// _ = "end of CoverTab[46634]"
}

func (rsa *RawSockaddrNFCLLCP) SetServiceNameLen(length int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:131
	_go_fuzz_dep_.CoverTab[46635]++
												rsa.Service_name_len = uint64(length)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:132
	// _ = "end of CoverTab[46635]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:137
func KexecFileLoad(kernelFd int, initrdFd int, cmdline string, flags int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:137
	_go_fuzz_dep_.CoverTab[46636]++
												cmdlineLen := len(cmdline)
												if cmdlineLen > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:139
		_go_fuzz_dep_.CoverTab[46638]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:143
		cmdlineLen++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:143
		// _ = "end of CoverTab[46638]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:144
		_go_fuzz_dep_.CoverTab[46639]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:144
		// _ = "end of CoverTab[46639]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:144
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:144
	// _ = "end of CoverTab[46636]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:144
	_go_fuzz_dep_.CoverTab[46637]++
												return kexecFileLoad(kernelFd, initrdFd, cmdlineLen, cmdline, flags)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:145
	// _ = "end of CoverTab[46637]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:146
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_linux_amd64.go:146
var _ = _go_fuzz_dep_.CoverTab
