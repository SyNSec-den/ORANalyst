// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

//line /snap/go/10455/src/net/fd_unix.go:7
package net

//line /snap/go/10455/src/net/fd_unix.go:7
import (
//line /snap/go/10455/src/net/fd_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/fd_unix.go:7
)
//line /snap/go/10455/src/net/fd_unix.go:7
import (
//line /snap/go/10455/src/net/fd_unix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/fd_unix.go:7
)

import (
	"context"
	"internal/poll"
	"os"
	"runtime"
	"syscall"
)

const (
	readSyscallName		= "read"
	readFromSyscallName	= "recvfrom"
	readMsgSyscallName	= "recvmsg"
	writeSyscallName	= "write"
	writeToSyscallName	= "sendto"
	writeMsgSyscallName	= "sendmsg"
)

func newFD(sysfd, family, sotype int, net string) (*netFD, error) {
//line /snap/go/10455/src/net/fd_unix.go:26
	_go_fuzz_dep_.CoverTab[5835]++
						ret := &netFD{
		pfd: poll.FD{
			Sysfd:		sysfd,
			IsStream:	sotype == syscall.SOCK_STREAM,
			ZeroReadIsEOF: sotype != syscall.SOCK_DGRAM && func() bool {
//line /snap/go/10455/src/net/fd_unix.go:31
				_go_fuzz_dep_.CoverTab[5836]++
//line /snap/go/10455/src/net/fd_unix.go:31
				return sotype != syscall.SOCK_RAW
//line /snap/go/10455/src/net/fd_unix.go:31
				// _ = "end of CoverTab[5836]"
//line /snap/go/10455/src/net/fd_unix.go:31
			}(),
		},
		family:	family,
		sotype:	sotype,
		net:	net,
	}
						return ret, nil
//line /snap/go/10455/src/net/fd_unix.go:37
	// _ = "end of CoverTab[5835]"
}

func (fd *netFD) init() error {
	return fd.pfd.Init(fd.net, true)
}

func (fd *netFD) name() string {
//line /snap/go/10455/src/net/fd_unix.go:44
	_go_fuzz_dep_.CoverTab[5837]++
						var ls, rs string
						if fd.laddr != nil {
//line /snap/go/10455/src/net/fd_unix.go:46
		_go_fuzz_dep_.CoverTab[528245]++
//line /snap/go/10455/src/net/fd_unix.go:46
		_go_fuzz_dep_.CoverTab[5840]++
							ls = fd.laddr.String()
//line /snap/go/10455/src/net/fd_unix.go:47
		// _ = "end of CoverTab[5840]"
	} else {
//line /snap/go/10455/src/net/fd_unix.go:48
		_go_fuzz_dep_.CoverTab[528246]++
//line /snap/go/10455/src/net/fd_unix.go:48
		_go_fuzz_dep_.CoverTab[5841]++
//line /snap/go/10455/src/net/fd_unix.go:48
		// _ = "end of CoverTab[5841]"
//line /snap/go/10455/src/net/fd_unix.go:48
	}
//line /snap/go/10455/src/net/fd_unix.go:48
	// _ = "end of CoverTab[5837]"
//line /snap/go/10455/src/net/fd_unix.go:48
	_go_fuzz_dep_.CoverTab[5838]++
						if fd.raddr != nil {
//line /snap/go/10455/src/net/fd_unix.go:49
		_go_fuzz_dep_.CoverTab[528247]++
//line /snap/go/10455/src/net/fd_unix.go:49
		_go_fuzz_dep_.CoverTab[5842]++
							rs = fd.raddr.String()
//line /snap/go/10455/src/net/fd_unix.go:50
		// _ = "end of CoverTab[5842]"
	} else {
//line /snap/go/10455/src/net/fd_unix.go:51
		_go_fuzz_dep_.CoverTab[528248]++
//line /snap/go/10455/src/net/fd_unix.go:51
		_go_fuzz_dep_.CoverTab[5843]++
//line /snap/go/10455/src/net/fd_unix.go:51
		// _ = "end of CoverTab[5843]"
//line /snap/go/10455/src/net/fd_unix.go:51
	}
//line /snap/go/10455/src/net/fd_unix.go:51
	// _ = "end of CoverTab[5838]"
//line /snap/go/10455/src/net/fd_unix.go:51
	_go_fuzz_dep_.CoverTab[5839]++
						return fd.net + ":" + ls + "->" + rs
//line /snap/go/10455/src/net/fd_unix.go:52
	// _ = "end of CoverTab[5839]"
}

func (fd *netFD) connect(ctx context.Context, la, ra syscall.Sockaddr) (rsa syscall.Sockaddr, ret error) {
//line /snap/go/10455/src/net/fd_unix.go:55
	_go_fuzz_dep_.CoverTab[5844]++

//line /snap/go/10455/src/net/fd_unix.go:59
	switch err := connectFunc(fd.pfd.Sysfd, ra); err {
	case syscall.EINPROGRESS, syscall.EALREADY, syscall.EINTR:
//line /snap/go/10455/src/net/fd_unix.go:60
		_go_fuzz_dep_.CoverTab[528249]++
//line /snap/go/10455/src/net/fd_unix.go:60
		_go_fuzz_dep_.CoverTab[5849]++
//line /snap/go/10455/src/net/fd_unix.go:60
		// _ = "end of CoverTab[5849]"
	case nil, syscall.EISCONN:
//line /snap/go/10455/src/net/fd_unix.go:61
		_go_fuzz_dep_.CoverTab[528250]++
//line /snap/go/10455/src/net/fd_unix.go:61
		_go_fuzz_dep_.CoverTab[5850]++
							select {
		case <-ctx.Done():
//line /snap/go/10455/src/net/fd_unix.go:63
			_go_fuzz_dep_.CoverTab[5856]++
								return nil, mapErr(ctx.Err())
//line /snap/go/10455/src/net/fd_unix.go:64
			// _ = "end of CoverTab[5856]"
		default:
//line /snap/go/10455/src/net/fd_unix.go:65
			_go_fuzz_dep_.CoverTab[5857]++
//line /snap/go/10455/src/net/fd_unix.go:65
			// _ = "end of CoverTab[5857]"
		}
//line /snap/go/10455/src/net/fd_unix.go:66
		// _ = "end of CoverTab[5850]"
//line /snap/go/10455/src/net/fd_unix.go:66
		_go_fuzz_dep_.CoverTab[5851]++
							if err := fd.pfd.Init(fd.net, true); err != nil {
//line /snap/go/10455/src/net/fd_unix.go:67
			_go_fuzz_dep_.CoverTab[528253]++
//line /snap/go/10455/src/net/fd_unix.go:67
			_go_fuzz_dep_.CoverTab[5858]++
								return nil, err
//line /snap/go/10455/src/net/fd_unix.go:68
			// _ = "end of CoverTab[5858]"
		} else {
//line /snap/go/10455/src/net/fd_unix.go:69
			_go_fuzz_dep_.CoverTab[528254]++
//line /snap/go/10455/src/net/fd_unix.go:69
			_go_fuzz_dep_.CoverTab[5859]++
//line /snap/go/10455/src/net/fd_unix.go:69
			// _ = "end of CoverTab[5859]"
//line /snap/go/10455/src/net/fd_unix.go:69
		}
//line /snap/go/10455/src/net/fd_unix.go:69
		// _ = "end of CoverTab[5851]"
//line /snap/go/10455/src/net/fd_unix.go:69
		_go_fuzz_dep_.CoverTab[5852]++
							runtime.KeepAlive(fd)
							return nil, nil
//line /snap/go/10455/src/net/fd_unix.go:71
		// _ = "end of CoverTab[5852]"
	case syscall.EINVAL:
//line /snap/go/10455/src/net/fd_unix.go:72
		_go_fuzz_dep_.CoverTab[528251]++
//line /snap/go/10455/src/net/fd_unix.go:72
		_go_fuzz_dep_.CoverTab[5853]++

//line /snap/go/10455/src/net/fd_unix.go:78
		if runtime.GOOS == "solaris" || func() bool {
//line /snap/go/10455/src/net/fd_unix.go:78
			_go_fuzz_dep_.CoverTab[5860]++
//line /snap/go/10455/src/net/fd_unix.go:78
			return runtime.GOOS == "illumos"
//line /snap/go/10455/src/net/fd_unix.go:78
			// _ = "end of CoverTab[5860]"
//line /snap/go/10455/src/net/fd_unix.go:78
		}() {
//line /snap/go/10455/src/net/fd_unix.go:78
			_go_fuzz_dep_.CoverTab[528255]++
//line /snap/go/10455/src/net/fd_unix.go:78
			_go_fuzz_dep_.CoverTab[5861]++
								return nil, nil
//line /snap/go/10455/src/net/fd_unix.go:79
			// _ = "end of CoverTab[5861]"
		} else {
//line /snap/go/10455/src/net/fd_unix.go:80
			_go_fuzz_dep_.CoverTab[528256]++
//line /snap/go/10455/src/net/fd_unix.go:80
			_go_fuzz_dep_.CoverTab[5862]++
//line /snap/go/10455/src/net/fd_unix.go:80
			// _ = "end of CoverTab[5862]"
//line /snap/go/10455/src/net/fd_unix.go:80
		}
//line /snap/go/10455/src/net/fd_unix.go:80
		// _ = "end of CoverTab[5853]"
//line /snap/go/10455/src/net/fd_unix.go:80
		_go_fuzz_dep_.CoverTab[5854]++
							fallthrough
//line /snap/go/10455/src/net/fd_unix.go:81
		// _ = "end of CoverTab[5854]"
	default:
//line /snap/go/10455/src/net/fd_unix.go:82
		_go_fuzz_dep_.CoverTab[528252]++
//line /snap/go/10455/src/net/fd_unix.go:82
		_go_fuzz_dep_.CoverTab[5855]++
							return nil, os.NewSyscallError("connect", err)
//line /snap/go/10455/src/net/fd_unix.go:83
		// _ = "end of CoverTab[5855]"
	}
//line /snap/go/10455/src/net/fd_unix.go:84
	// _ = "end of CoverTab[5844]"
//line /snap/go/10455/src/net/fd_unix.go:84
	_go_fuzz_dep_.CoverTab[5845]++
						if err := fd.pfd.Init(fd.net, true); err != nil {
//line /snap/go/10455/src/net/fd_unix.go:85
		_go_fuzz_dep_.CoverTab[528257]++
//line /snap/go/10455/src/net/fd_unix.go:85
		_go_fuzz_dep_.CoverTab[5863]++
							return nil, err
//line /snap/go/10455/src/net/fd_unix.go:86
		// _ = "end of CoverTab[5863]"
	} else {
//line /snap/go/10455/src/net/fd_unix.go:87
		_go_fuzz_dep_.CoverTab[528258]++
//line /snap/go/10455/src/net/fd_unix.go:87
		_go_fuzz_dep_.CoverTab[5864]++
//line /snap/go/10455/src/net/fd_unix.go:87
		// _ = "end of CoverTab[5864]"
//line /snap/go/10455/src/net/fd_unix.go:87
	}
//line /snap/go/10455/src/net/fd_unix.go:87
	// _ = "end of CoverTab[5845]"
//line /snap/go/10455/src/net/fd_unix.go:87
	_go_fuzz_dep_.CoverTab[5846]++
						if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
//line /snap/go/10455/src/net/fd_unix.go:88
		_go_fuzz_dep_.CoverTab[528259]++
//line /snap/go/10455/src/net/fd_unix.go:88
		_go_fuzz_dep_.CoverTab[5865]++
							fd.pfd.SetWriteDeadline(deadline)
							defer fd.pfd.SetWriteDeadline(noDeadline)
//line /snap/go/10455/src/net/fd_unix.go:90
		// _ = "end of CoverTab[5865]"
	} else {
//line /snap/go/10455/src/net/fd_unix.go:91
		_go_fuzz_dep_.CoverTab[528260]++
//line /snap/go/10455/src/net/fd_unix.go:91
		_go_fuzz_dep_.CoverTab[5866]++
//line /snap/go/10455/src/net/fd_unix.go:91
		// _ = "end of CoverTab[5866]"
//line /snap/go/10455/src/net/fd_unix.go:91
	}
//line /snap/go/10455/src/net/fd_unix.go:91
	// _ = "end of CoverTab[5846]"
//line /snap/go/10455/src/net/fd_unix.go:91
	_go_fuzz_dep_.CoverTab[5847]++

//line /snap/go/10455/src/net/fd_unix.go:98
	ctxDone := ctx.Done()
	if ctxDone != nil {
//line /snap/go/10455/src/net/fd_unix.go:99
		_go_fuzz_dep_.CoverTab[528261]++
//line /snap/go/10455/src/net/fd_unix.go:99
		_go_fuzz_dep_.CoverTab[5867]++

//line /snap/go/10455/src/net/fd_unix.go:102
		done := make(chan struct{})
		interruptRes := make(chan error)
		defer func() {
//line /snap/go/10455/src/net/fd_unix.go:104
			_go_fuzz_dep_.CoverTab[5869]++
								close(done)
								if ctxErr := <-interruptRes; ctxErr != nil && func() bool {
//line /snap/go/10455/src/net/fd_unix.go:106
				_go_fuzz_dep_.CoverTab[5870]++
//line /snap/go/10455/src/net/fd_unix.go:106
				return ret == nil
//line /snap/go/10455/src/net/fd_unix.go:106
				// _ = "end of CoverTab[5870]"
//line /snap/go/10455/src/net/fd_unix.go:106
			}() {
//line /snap/go/10455/src/net/fd_unix.go:106
				_go_fuzz_dep_.CoverTab[528263]++
//line /snap/go/10455/src/net/fd_unix.go:106
				_go_fuzz_dep_.CoverTab[5871]++

//line /snap/go/10455/src/net/fd_unix.go:113
				ret = mapErr(ctxErr)
									fd.Close()
//line /snap/go/10455/src/net/fd_unix.go:114
				// _ = "end of CoverTab[5871]"
			} else {
//line /snap/go/10455/src/net/fd_unix.go:115
				_go_fuzz_dep_.CoverTab[528264]++
//line /snap/go/10455/src/net/fd_unix.go:115
				_go_fuzz_dep_.CoverTab[5872]++
//line /snap/go/10455/src/net/fd_unix.go:115
				// _ = "end of CoverTab[5872]"
//line /snap/go/10455/src/net/fd_unix.go:115
			}
//line /snap/go/10455/src/net/fd_unix.go:115
			// _ = "end of CoverTab[5869]"
		}()
//line /snap/go/10455/src/net/fd_unix.go:116
		// _ = "end of CoverTab[5867]"
//line /snap/go/10455/src/net/fd_unix.go:116
		_go_fuzz_dep_.CoverTab[5868]++
//line /snap/go/10455/src/net/fd_unix.go:116
		_curRoutineNum7_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/net/fd_unix.go:116
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum7_)
							go func() {
//line /snap/go/10455/src/net/fd_unix.go:117
			_go_fuzz_dep_.CoverTab[5873]++
//line /snap/go/10455/src/net/fd_unix.go:117
			defer func() {
//line /snap/go/10455/src/net/fd_unix.go:117
				_go_fuzz_dep_.CoverTab[5874]++
//line /snap/go/10455/src/net/fd_unix.go:117
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum7_)
//line /snap/go/10455/src/net/fd_unix.go:117
				// _ = "end of CoverTab[5874]"
//line /snap/go/10455/src/net/fd_unix.go:117
			}()
								select {
			case <-ctxDone:
//line /snap/go/10455/src/net/fd_unix.go:119
				_go_fuzz_dep_.CoverTab[5875]++

//line /snap/go/10455/src/net/fd_unix.go:123
				fd.pfd.SetWriteDeadline(aLongTimeAgo)
									testHookCanceledDial()
									interruptRes <- ctx.Err()
//line /snap/go/10455/src/net/fd_unix.go:125
				// _ = "end of CoverTab[5875]"
			case <-done:
//line /snap/go/10455/src/net/fd_unix.go:126
				_go_fuzz_dep_.CoverTab[5876]++
									interruptRes <- nil
//line /snap/go/10455/src/net/fd_unix.go:127
				// _ = "end of CoverTab[5876]"
			}
//line /snap/go/10455/src/net/fd_unix.go:128
			// _ = "end of CoverTab[5873]"
		}()
//line /snap/go/10455/src/net/fd_unix.go:129
		// _ = "end of CoverTab[5868]"
	} else {
//line /snap/go/10455/src/net/fd_unix.go:130
		_go_fuzz_dep_.CoverTab[528262]++
//line /snap/go/10455/src/net/fd_unix.go:130
		_go_fuzz_dep_.CoverTab[5877]++
//line /snap/go/10455/src/net/fd_unix.go:130
		// _ = "end of CoverTab[5877]"
//line /snap/go/10455/src/net/fd_unix.go:130
	}
//line /snap/go/10455/src/net/fd_unix.go:130
	// _ = "end of CoverTab[5847]"
//line /snap/go/10455/src/net/fd_unix.go:130
	_go_fuzz_dep_.CoverTab[5848]++
//line /snap/go/10455/src/net/fd_unix.go:130
	_go_fuzz_dep_.CoverTab[786674] = 0

						for {
//line /snap/go/10455/src/net/fd_unix.go:132
		if _go_fuzz_dep_.CoverTab[786674] == 0 {
//line /snap/go/10455/src/net/fd_unix.go:132
			_go_fuzz_dep_.CoverTab[528287]++
//line /snap/go/10455/src/net/fd_unix.go:132
		} else {
//line /snap/go/10455/src/net/fd_unix.go:132
			_go_fuzz_dep_.CoverTab[528288]++
//line /snap/go/10455/src/net/fd_unix.go:132
		}
//line /snap/go/10455/src/net/fd_unix.go:132
		_go_fuzz_dep_.CoverTab[786674] = 1
//line /snap/go/10455/src/net/fd_unix.go:132
		_go_fuzz_dep_.CoverTab[5878]++

//line /snap/go/10455/src/net/fd_unix.go:141
		if err := fd.pfd.WaitWrite(); err != nil {
//line /snap/go/10455/src/net/fd_unix.go:141
			_go_fuzz_dep_.CoverTab[528265]++
//line /snap/go/10455/src/net/fd_unix.go:141
			_go_fuzz_dep_.CoverTab[5882]++
								select {
			case <-ctxDone:
//line /snap/go/10455/src/net/fd_unix.go:143
				_go_fuzz_dep_.CoverTab[5884]++
									return nil, mapErr(ctx.Err())
//line /snap/go/10455/src/net/fd_unix.go:144
				// _ = "end of CoverTab[5884]"
			default:
//line /snap/go/10455/src/net/fd_unix.go:145
				_go_fuzz_dep_.CoverTab[5885]++
//line /snap/go/10455/src/net/fd_unix.go:145
				// _ = "end of CoverTab[5885]"
			}
//line /snap/go/10455/src/net/fd_unix.go:146
			// _ = "end of CoverTab[5882]"
//line /snap/go/10455/src/net/fd_unix.go:146
			_go_fuzz_dep_.CoverTab[5883]++
								return nil, err
//line /snap/go/10455/src/net/fd_unix.go:147
			// _ = "end of CoverTab[5883]"
		} else {
//line /snap/go/10455/src/net/fd_unix.go:148
			_go_fuzz_dep_.CoverTab[528266]++
//line /snap/go/10455/src/net/fd_unix.go:148
			_go_fuzz_dep_.CoverTab[5886]++
//line /snap/go/10455/src/net/fd_unix.go:148
			// _ = "end of CoverTab[5886]"
//line /snap/go/10455/src/net/fd_unix.go:148
		}
//line /snap/go/10455/src/net/fd_unix.go:148
		// _ = "end of CoverTab[5878]"
//line /snap/go/10455/src/net/fd_unix.go:148
		_go_fuzz_dep_.CoverTab[5879]++
							nerr, err := getsockoptIntFunc(fd.pfd.Sysfd, syscall.SOL_SOCKET, syscall.SO_ERROR)
							if err != nil {
//line /snap/go/10455/src/net/fd_unix.go:150
			_go_fuzz_dep_.CoverTab[528267]++
//line /snap/go/10455/src/net/fd_unix.go:150
			_go_fuzz_dep_.CoverTab[5887]++
								return nil, os.NewSyscallError("getsockopt", err)
//line /snap/go/10455/src/net/fd_unix.go:151
			// _ = "end of CoverTab[5887]"
		} else {
//line /snap/go/10455/src/net/fd_unix.go:152
			_go_fuzz_dep_.CoverTab[528268]++
//line /snap/go/10455/src/net/fd_unix.go:152
			_go_fuzz_dep_.CoverTab[5888]++
//line /snap/go/10455/src/net/fd_unix.go:152
			// _ = "end of CoverTab[5888]"
//line /snap/go/10455/src/net/fd_unix.go:152
		}
//line /snap/go/10455/src/net/fd_unix.go:152
		// _ = "end of CoverTab[5879]"
//line /snap/go/10455/src/net/fd_unix.go:152
		_go_fuzz_dep_.CoverTab[5880]++
							switch err := syscall.Errno(nerr); err {
		case syscall.EINPROGRESS, syscall.EALREADY, syscall.EINTR:
//line /snap/go/10455/src/net/fd_unix.go:154
			_go_fuzz_dep_.CoverTab[528269]++
//line /snap/go/10455/src/net/fd_unix.go:154
			_go_fuzz_dep_.CoverTab[5889]++
//line /snap/go/10455/src/net/fd_unix.go:154
			// _ = "end of CoverTab[5889]"
		case syscall.EISCONN:
//line /snap/go/10455/src/net/fd_unix.go:155
			_go_fuzz_dep_.CoverTab[528270]++
//line /snap/go/10455/src/net/fd_unix.go:155
			_go_fuzz_dep_.CoverTab[5890]++
								return nil, nil
//line /snap/go/10455/src/net/fd_unix.go:156
			// _ = "end of CoverTab[5890]"
		case syscall.Errno(0):
//line /snap/go/10455/src/net/fd_unix.go:157
			_go_fuzz_dep_.CoverTab[528271]++
//line /snap/go/10455/src/net/fd_unix.go:157
			_go_fuzz_dep_.CoverTab[5891]++

//line /snap/go/10455/src/net/fd_unix.go:161
			if rsa, err := syscall.Getpeername(fd.pfd.Sysfd); err == nil {
//line /snap/go/10455/src/net/fd_unix.go:161
				_go_fuzz_dep_.CoverTab[528273]++
//line /snap/go/10455/src/net/fd_unix.go:161
				_go_fuzz_dep_.CoverTab[5893]++
									return rsa, nil
//line /snap/go/10455/src/net/fd_unix.go:162
				// _ = "end of CoverTab[5893]"
			} else {
//line /snap/go/10455/src/net/fd_unix.go:163
				_go_fuzz_dep_.CoverTab[528274]++
//line /snap/go/10455/src/net/fd_unix.go:163
				_go_fuzz_dep_.CoverTab[5894]++
//line /snap/go/10455/src/net/fd_unix.go:163
				// _ = "end of CoverTab[5894]"
//line /snap/go/10455/src/net/fd_unix.go:163
			}
//line /snap/go/10455/src/net/fd_unix.go:163
			// _ = "end of CoverTab[5891]"
		default:
//line /snap/go/10455/src/net/fd_unix.go:164
			_go_fuzz_dep_.CoverTab[528272]++
//line /snap/go/10455/src/net/fd_unix.go:164
			_go_fuzz_dep_.CoverTab[5892]++
								return nil, os.NewSyscallError("connect", err)
//line /snap/go/10455/src/net/fd_unix.go:165
			// _ = "end of CoverTab[5892]"
		}
//line /snap/go/10455/src/net/fd_unix.go:166
		// _ = "end of CoverTab[5880]"
//line /snap/go/10455/src/net/fd_unix.go:166
		_go_fuzz_dep_.CoverTab[5881]++
							runtime.KeepAlive(fd)
//line /snap/go/10455/src/net/fd_unix.go:167
		// _ = "end of CoverTab[5881]"
	}
//line /snap/go/10455/src/net/fd_unix.go:168
	// _ = "end of CoverTab[5848]"
}

func (fd *netFD) accept() (netfd *netFD, err error) {
//line /snap/go/10455/src/net/fd_unix.go:171
	_go_fuzz_dep_.CoverTab[5895]++
						d, rsa, errcall, err := fd.pfd.Accept()
						if err != nil {
//line /snap/go/10455/src/net/fd_unix.go:173
		_go_fuzz_dep_.CoverTab[528275]++
//line /snap/go/10455/src/net/fd_unix.go:173
		_go_fuzz_dep_.CoverTab[5899]++
							if errcall != "" {
//line /snap/go/10455/src/net/fd_unix.go:174
			_go_fuzz_dep_.CoverTab[528277]++
//line /snap/go/10455/src/net/fd_unix.go:174
			_go_fuzz_dep_.CoverTab[5901]++
								err = wrapSyscallError(errcall, err)
//line /snap/go/10455/src/net/fd_unix.go:175
			// _ = "end of CoverTab[5901]"
		} else {
//line /snap/go/10455/src/net/fd_unix.go:176
			_go_fuzz_dep_.CoverTab[528278]++
//line /snap/go/10455/src/net/fd_unix.go:176
			_go_fuzz_dep_.CoverTab[5902]++
//line /snap/go/10455/src/net/fd_unix.go:176
			// _ = "end of CoverTab[5902]"
//line /snap/go/10455/src/net/fd_unix.go:176
		}
//line /snap/go/10455/src/net/fd_unix.go:176
		// _ = "end of CoverTab[5899]"
//line /snap/go/10455/src/net/fd_unix.go:176
		_go_fuzz_dep_.CoverTab[5900]++
							return nil, err
//line /snap/go/10455/src/net/fd_unix.go:177
		// _ = "end of CoverTab[5900]"
	} else {
//line /snap/go/10455/src/net/fd_unix.go:178
		_go_fuzz_dep_.CoverTab[528276]++
//line /snap/go/10455/src/net/fd_unix.go:178
		_go_fuzz_dep_.CoverTab[5903]++
//line /snap/go/10455/src/net/fd_unix.go:178
		// _ = "end of CoverTab[5903]"
//line /snap/go/10455/src/net/fd_unix.go:178
	}
//line /snap/go/10455/src/net/fd_unix.go:178
	// _ = "end of CoverTab[5895]"
//line /snap/go/10455/src/net/fd_unix.go:178
	_go_fuzz_dep_.CoverTab[5896]++

						if netfd, err = newFD(d, fd.family, fd.sotype, fd.net); err != nil {
//line /snap/go/10455/src/net/fd_unix.go:180
		_go_fuzz_dep_.CoverTab[528279]++
//line /snap/go/10455/src/net/fd_unix.go:180
		_go_fuzz_dep_.CoverTab[5904]++
							poll.CloseFunc(d)
							return nil, err
//line /snap/go/10455/src/net/fd_unix.go:182
		// _ = "end of CoverTab[5904]"
	} else {
//line /snap/go/10455/src/net/fd_unix.go:183
		_go_fuzz_dep_.CoverTab[528280]++
//line /snap/go/10455/src/net/fd_unix.go:183
		_go_fuzz_dep_.CoverTab[5905]++
//line /snap/go/10455/src/net/fd_unix.go:183
		// _ = "end of CoverTab[5905]"
//line /snap/go/10455/src/net/fd_unix.go:183
	}
//line /snap/go/10455/src/net/fd_unix.go:183
	// _ = "end of CoverTab[5896]"
//line /snap/go/10455/src/net/fd_unix.go:183
	_go_fuzz_dep_.CoverTab[5897]++
						if err = netfd.init(); err != nil {
//line /snap/go/10455/src/net/fd_unix.go:184
		_go_fuzz_dep_.CoverTab[528281]++
//line /snap/go/10455/src/net/fd_unix.go:184
		_go_fuzz_dep_.CoverTab[5906]++
							netfd.Close()
							return nil, err
//line /snap/go/10455/src/net/fd_unix.go:186
		// _ = "end of CoverTab[5906]"
	} else {
//line /snap/go/10455/src/net/fd_unix.go:187
		_go_fuzz_dep_.CoverTab[528282]++
//line /snap/go/10455/src/net/fd_unix.go:187
		_go_fuzz_dep_.CoverTab[5907]++
//line /snap/go/10455/src/net/fd_unix.go:187
		// _ = "end of CoverTab[5907]"
//line /snap/go/10455/src/net/fd_unix.go:187
	}
//line /snap/go/10455/src/net/fd_unix.go:187
	// _ = "end of CoverTab[5897]"
//line /snap/go/10455/src/net/fd_unix.go:187
	_go_fuzz_dep_.CoverTab[5898]++
						lsa, _ := syscall.Getsockname(netfd.pfd.Sysfd)
						netfd.setAddr(netfd.addrFunc()(lsa), netfd.addrFunc()(rsa))
						return netfd, nil
//line /snap/go/10455/src/net/fd_unix.go:190
	// _ = "end of CoverTab[5898]"
}

// Defined in os package.
func newUnixFile(fd int, name string) *os.File

func (fd *netFD) dup() (f *os.File, err error) {
//line /snap/go/10455/src/net/fd_unix.go:196
	_go_fuzz_dep_.CoverTab[5908]++
						ns, call, err := fd.pfd.Dup()
						if err != nil {
//line /snap/go/10455/src/net/fd_unix.go:198
		_go_fuzz_dep_.CoverTab[528283]++
//line /snap/go/10455/src/net/fd_unix.go:198
		_go_fuzz_dep_.CoverTab[5910]++
							if call != "" {
//line /snap/go/10455/src/net/fd_unix.go:199
			_go_fuzz_dep_.CoverTab[528285]++
//line /snap/go/10455/src/net/fd_unix.go:199
			_go_fuzz_dep_.CoverTab[5912]++
								err = os.NewSyscallError(call, err)
//line /snap/go/10455/src/net/fd_unix.go:200
			// _ = "end of CoverTab[5912]"
		} else {
//line /snap/go/10455/src/net/fd_unix.go:201
			_go_fuzz_dep_.CoverTab[528286]++
//line /snap/go/10455/src/net/fd_unix.go:201
			_go_fuzz_dep_.CoverTab[5913]++
//line /snap/go/10455/src/net/fd_unix.go:201
			// _ = "end of CoverTab[5913]"
//line /snap/go/10455/src/net/fd_unix.go:201
		}
//line /snap/go/10455/src/net/fd_unix.go:201
		// _ = "end of CoverTab[5910]"
//line /snap/go/10455/src/net/fd_unix.go:201
		_go_fuzz_dep_.CoverTab[5911]++
							return nil, err
//line /snap/go/10455/src/net/fd_unix.go:202
		// _ = "end of CoverTab[5911]"
	} else {
//line /snap/go/10455/src/net/fd_unix.go:203
		_go_fuzz_dep_.CoverTab[528284]++
//line /snap/go/10455/src/net/fd_unix.go:203
		_go_fuzz_dep_.CoverTab[5914]++
//line /snap/go/10455/src/net/fd_unix.go:203
		// _ = "end of CoverTab[5914]"
//line /snap/go/10455/src/net/fd_unix.go:203
	}
//line /snap/go/10455/src/net/fd_unix.go:203
	// _ = "end of CoverTab[5908]"
//line /snap/go/10455/src/net/fd_unix.go:203
	_go_fuzz_dep_.CoverTab[5909]++

						return newUnixFile(ns, fd.name()), nil
//line /snap/go/10455/src/net/fd_unix.go:205
	// _ = "end of CoverTab[5909]"
}

//line /snap/go/10455/src/net/fd_unix.go:206
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/fd_unix.go:206
var _ = _go_fuzz_dep_.CoverTab
