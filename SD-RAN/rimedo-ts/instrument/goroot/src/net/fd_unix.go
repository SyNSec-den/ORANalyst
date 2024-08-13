// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

//line /usr/local/go/src/net/fd_unix.go:7
package net

//line /usr/local/go/src/net/fd_unix.go:7
import (
//line /usr/local/go/src/net/fd_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/fd_unix.go:7
)
//line /usr/local/go/src/net/fd_unix.go:7
import (
//line /usr/local/go/src/net/fd_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/fd_unix.go:7
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
//line /usr/local/go/src/net/fd_unix.go:26
	_go_fuzz_dep_.CoverTab[13849]++
						ret := &netFD{
		pfd: poll.FD{
			Sysfd:		sysfd,
			IsStream:	sotype == syscall.SOCK_STREAM,
			ZeroReadIsEOF: sotype != syscall.SOCK_DGRAM && func() bool {
//line /usr/local/go/src/net/fd_unix.go:31
				_go_fuzz_dep_.CoverTab[13850]++
//line /usr/local/go/src/net/fd_unix.go:31
				return sotype != syscall.SOCK_RAW
//line /usr/local/go/src/net/fd_unix.go:31
				// _ = "end of CoverTab[13850]"
//line /usr/local/go/src/net/fd_unix.go:31
			}(),
		},
		family:	family,
		sotype:	sotype,
		net:	net,
	}
						return ret, nil
//line /usr/local/go/src/net/fd_unix.go:37
	// _ = "end of CoverTab[13849]"
}

func (fd *netFD) init() error {
	return fd.pfd.Init(fd.net, true)
}

func (fd *netFD) name() string {
//line /usr/local/go/src/net/fd_unix.go:44
	_go_fuzz_dep_.CoverTab[13851]++
						var ls, rs string
						if fd.laddr != nil {
//line /usr/local/go/src/net/fd_unix.go:46
		_go_fuzz_dep_.CoverTab[13854]++
							ls = fd.laddr.String()
//line /usr/local/go/src/net/fd_unix.go:47
		// _ = "end of CoverTab[13854]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:48
		_go_fuzz_dep_.CoverTab[13855]++
//line /usr/local/go/src/net/fd_unix.go:48
		// _ = "end of CoverTab[13855]"
//line /usr/local/go/src/net/fd_unix.go:48
	}
//line /usr/local/go/src/net/fd_unix.go:48
	// _ = "end of CoverTab[13851]"
//line /usr/local/go/src/net/fd_unix.go:48
	_go_fuzz_dep_.CoverTab[13852]++
						if fd.raddr != nil {
//line /usr/local/go/src/net/fd_unix.go:49
		_go_fuzz_dep_.CoverTab[13856]++
							rs = fd.raddr.String()
//line /usr/local/go/src/net/fd_unix.go:50
		// _ = "end of CoverTab[13856]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:51
		_go_fuzz_dep_.CoverTab[13857]++
//line /usr/local/go/src/net/fd_unix.go:51
		// _ = "end of CoverTab[13857]"
//line /usr/local/go/src/net/fd_unix.go:51
	}
//line /usr/local/go/src/net/fd_unix.go:51
	// _ = "end of CoverTab[13852]"
//line /usr/local/go/src/net/fd_unix.go:51
	_go_fuzz_dep_.CoverTab[13853]++
						return fd.net + ":" + ls + "->" + rs
//line /usr/local/go/src/net/fd_unix.go:52
	// _ = "end of CoverTab[13853]"
}

func (fd *netFD) connect(ctx context.Context, la, ra syscall.Sockaddr) (rsa syscall.Sockaddr, ret error) {
//line /usr/local/go/src/net/fd_unix.go:55
	_go_fuzz_dep_.CoverTab[13858]++

//line /usr/local/go/src/net/fd_unix.go:59
	switch err := connectFunc(fd.pfd.Sysfd, ra); err {
	case syscall.EINPROGRESS, syscall.EALREADY, syscall.EINTR:
//line /usr/local/go/src/net/fd_unix.go:60
		_go_fuzz_dep_.CoverTab[13863]++
//line /usr/local/go/src/net/fd_unix.go:60
		// _ = "end of CoverTab[13863]"
	case nil, syscall.EISCONN:
//line /usr/local/go/src/net/fd_unix.go:61
		_go_fuzz_dep_.CoverTab[13864]++
							select {
		case <-ctx.Done():
//line /usr/local/go/src/net/fd_unix.go:63
			_go_fuzz_dep_.CoverTab[13870]++
								return nil, mapErr(ctx.Err())
//line /usr/local/go/src/net/fd_unix.go:64
			// _ = "end of CoverTab[13870]"
		default:
//line /usr/local/go/src/net/fd_unix.go:65
			_go_fuzz_dep_.CoverTab[13871]++
//line /usr/local/go/src/net/fd_unix.go:65
			// _ = "end of CoverTab[13871]"
		}
//line /usr/local/go/src/net/fd_unix.go:66
		// _ = "end of CoverTab[13864]"
//line /usr/local/go/src/net/fd_unix.go:66
		_go_fuzz_dep_.CoverTab[13865]++
							if err := fd.pfd.Init(fd.net, true); err != nil {
//line /usr/local/go/src/net/fd_unix.go:67
			_go_fuzz_dep_.CoverTab[13872]++
								return nil, err
//line /usr/local/go/src/net/fd_unix.go:68
			// _ = "end of CoverTab[13872]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:69
			_go_fuzz_dep_.CoverTab[13873]++
//line /usr/local/go/src/net/fd_unix.go:69
			// _ = "end of CoverTab[13873]"
//line /usr/local/go/src/net/fd_unix.go:69
		}
//line /usr/local/go/src/net/fd_unix.go:69
		// _ = "end of CoverTab[13865]"
//line /usr/local/go/src/net/fd_unix.go:69
		_go_fuzz_dep_.CoverTab[13866]++
							runtime.KeepAlive(fd)
							return nil, nil
//line /usr/local/go/src/net/fd_unix.go:71
		// _ = "end of CoverTab[13866]"
	case syscall.EINVAL:
//line /usr/local/go/src/net/fd_unix.go:72
		_go_fuzz_dep_.CoverTab[13867]++

//line /usr/local/go/src/net/fd_unix.go:78
		if runtime.GOOS == "solaris" || func() bool {
//line /usr/local/go/src/net/fd_unix.go:78
			_go_fuzz_dep_.CoverTab[13874]++
//line /usr/local/go/src/net/fd_unix.go:78
			return runtime.GOOS == "illumos"
//line /usr/local/go/src/net/fd_unix.go:78
			// _ = "end of CoverTab[13874]"
//line /usr/local/go/src/net/fd_unix.go:78
		}() {
//line /usr/local/go/src/net/fd_unix.go:78
			_go_fuzz_dep_.CoverTab[13875]++
								return nil, nil
//line /usr/local/go/src/net/fd_unix.go:79
			// _ = "end of CoverTab[13875]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:80
			_go_fuzz_dep_.CoverTab[13876]++
//line /usr/local/go/src/net/fd_unix.go:80
			// _ = "end of CoverTab[13876]"
//line /usr/local/go/src/net/fd_unix.go:80
		}
//line /usr/local/go/src/net/fd_unix.go:80
		// _ = "end of CoverTab[13867]"
//line /usr/local/go/src/net/fd_unix.go:80
		_go_fuzz_dep_.CoverTab[13868]++
							fallthrough
//line /usr/local/go/src/net/fd_unix.go:81
		// _ = "end of CoverTab[13868]"
	default:
//line /usr/local/go/src/net/fd_unix.go:82
		_go_fuzz_dep_.CoverTab[13869]++
							return nil, os.NewSyscallError("connect", err)
//line /usr/local/go/src/net/fd_unix.go:83
		// _ = "end of CoverTab[13869]"
	}
//line /usr/local/go/src/net/fd_unix.go:84
	// _ = "end of CoverTab[13858]"
//line /usr/local/go/src/net/fd_unix.go:84
	_go_fuzz_dep_.CoverTab[13859]++
						if err := fd.pfd.Init(fd.net, true); err != nil {
//line /usr/local/go/src/net/fd_unix.go:85
		_go_fuzz_dep_.CoverTab[13877]++
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:86
		// _ = "end of CoverTab[13877]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:87
		_go_fuzz_dep_.CoverTab[13878]++
//line /usr/local/go/src/net/fd_unix.go:87
		// _ = "end of CoverTab[13878]"
//line /usr/local/go/src/net/fd_unix.go:87
	}
//line /usr/local/go/src/net/fd_unix.go:87
	// _ = "end of CoverTab[13859]"
//line /usr/local/go/src/net/fd_unix.go:87
	_go_fuzz_dep_.CoverTab[13860]++
						if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
//line /usr/local/go/src/net/fd_unix.go:88
		_go_fuzz_dep_.CoverTab[13879]++
							fd.pfd.SetWriteDeadline(deadline)
							defer fd.pfd.SetWriteDeadline(noDeadline)
//line /usr/local/go/src/net/fd_unix.go:90
		// _ = "end of CoverTab[13879]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:91
		_go_fuzz_dep_.CoverTab[13880]++
//line /usr/local/go/src/net/fd_unix.go:91
		// _ = "end of CoverTab[13880]"
//line /usr/local/go/src/net/fd_unix.go:91
	}
//line /usr/local/go/src/net/fd_unix.go:91
	// _ = "end of CoverTab[13860]"
//line /usr/local/go/src/net/fd_unix.go:91
	_go_fuzz_dep_.CoverTab[13861]++

//line /usr/local/go/src/net/fd_unix.go:98
	ctxDone := ctx.Done()
	if ctxDone != nil {
//line /usr/local/go/src/net/fd_unix.go:99
		_go_fuzz_dep_.CoverTab[13881]++

//line /usr/local/go/src/net/fd_unix.go:102
		done := make(chan struct{})
		interruptRes := make(chan error)
		defer func() {
//line /usr/local/go/src/net/fd_unix.go:104
			_go_fuzz_dep_.CoverTab[13883]++
								close(done)
								if ctxErr := <-interruptRes; ctxErr != nil && func() bool {
//line /usr/local/go/src/net/fd_unix.go:106
				_go_fuzz_dep_.CoverTab[13884]++
//line /usr/local/go/src/net/fd_unix.go:106
				return ret == nil
//line /usr/local/go/src/net/fd_unix.go:106
				// _ = "end of CoverTab[13884]"
//line /usr/local/go/src/net/fd_unix.go:106
			}() {
//line /usr/local/go/src/net/fd_unix.go:106
				_go_fuzz_dep_.CoverTab[13885]++

//line /usr/local/go/src/net/fd_unix.go:113
				ret = mapErr(ctxErr)
									fd.Close()
//line /usr/local/go/src/net/fd_unix.go:114
				// _ = "end of CoverTab[13885]"
			} else {
//line /usr/local/go/src/net/fd_unix.go:115
				_go_fuzz_dep_.CoverTab[13886]++
//line /usr/local/go/src/net/fd_unix.go:115
				// _ = "end of CoverTab[13886]"
//line /usr/local/go/src/net/fd_unix.go:115
			}
//line /usr/local/go/src/net/fd_unix.go:115
			// _ = "end of CoverTab[13883]"
		}()
//line /usr/local/go/src/net/fd_unix.go:116
		// _ = "end of CoverTab[13881]"
//line /usr/local/go/src/net/fd_unix.go:116
		_go_fuzz_dep_.CoverTab[13882]++
//line /usr/local/go/src/net/fd_unix.go:116
		_curRoutineNum9_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/fd_unix.go:116
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum9_)
							go func() {
//line /usr/local/go/src/net/fd_unix.go:117
			_go_fuzz_dep_.CoverTab[13887]++
//line /usr/local/go/src/net/fd_unix.go:117
			defer func() {
//line /usr/local/go/src/net/fd_unix.go:117
				_go_fuzz_dep_.CoverTab[13888]++
//line /usr/local/go/src/net/fd_unix.go:117
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum9_)
//line /usr/local/go/src/net/fd_unix.go:117
				// _ = "end of CoverTab[13888]"
//line /usr/local/go/src/net/fd_unix.go:117
			}()
								select {
			case <-ctxDone:
//line /usr/local/go/src/net/fd_unix.go:119
				_go_fuzz_dep_.CoverTab[13889]++

//line /usr/local/go/src/net/fd_unix.go:123
				fd.pfd.SetWriteDeadline(aLongTimeAgo)
									testHookCanceledDial()
									interruptRes <- ctx.Err()
//line /usr/local/go/src/net/fd_unix.go:125
				// _ = "end of CoverTab[13889]"
			case <-done:
//line /usr/local/go/src/net/fd_unix.go:126
				_go_fuzz_dep_.CoverTab[13890]++
									interruptRes <- nil
//line /usr/local/go/src/net/fd_unix.go:127
				// _ = "end of CoverTab[13890]"
			}
//line /usr/local/go/src/net/fd_unix.go:128
			// _ = "end of CoverTab[13887]"
		}()
//line /usr/local/go/src/net/fd_unix.go:129
		// _ = "end of CoverTab[13882]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:130
		_go_fuzz_dep_.CoverTab[13891]++
//line /usr/local/go/src/net/fd_unix.go:130
		// _ = "end of CoverTab[13891]"
//line /usr/local/go/src/net/fd_unix.go:130
	}
//line /usr/local/go/src/net/fd_unix.go:130
	// _ = "end of CoverTab[13861]"
//line /usr/local/go/src/net/fd_unix.go:130
	_go_fuzz_dep_.CoverTab[13862]++

						for {
//line /usr/local/go/src/net/fd_unix.go:132
		_go_fuzz_dep_.CoverTab[13892]++

//line /usr/local/go/src/net/fd_unix.go:141
		if err := fd.pfd.WaitWrite(); err != nil {
//line /usr/local/go/src/net/fd_unix.go:141
			_go_fuzz_dep_.CoverTab[13896]++
								select {
			case <-ctxDone:
//line /usr/local/go/src/net/fd_unix.go:143
				_go_fuzz_dep_.CoverTab[13898]++
									return nil, mapErr(ctx.Err())
//line /usr/local/go/src/net/fd_unix.go:144
				// _ = "end of CoverTab[13898]"
			default:
//line /usr/local/go/src/net/fd_unix.go:145
				_go_fuzz_dep_.CoverTab[13899]++
//line /usr/local/go/src/net/fd_unix.go:145
				// _ = "end of CoverTab[13899]"
			}
//line /usr/local/go/src/net/fd_unix.go:146
			// _ = "end of CoverTab[13896]"
//line /usr/local/go/src/net/fd_unix.go:146
			_go_fuzz_dep_.CoverTab[13897]++
								return nil, err
//line /usr/local/go/src/net/fd_unix.go:147
			// _ = "end of CoverTab[13897]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:148
			_go_fuzz_dep_.CoverTab[13900]++
//line /usr/local/go/src/net/fd_unix.go:148
			// _ = "end of CoverTab[13900]"
//line /usr/local/go/src/net/fd_unix.go:148
		}
//line /usr/local/go/src/net/fd_unix.go:148
		// _ = "end of CoverTab[13892]"
//line /usr/local/go/src/net/fd_unix.go:148
		_go_fuzz_dep_.CoverTab[13893]++
							nerr, err := getsockoptIntFunc(fd.pfd.Sysfd, syscall.SOL_SOCKET, syscall.SO_ERROR)
							if err != nil {
//line /usr/local/go/src/net/fd_unix.go:150
			_go_fuzz_dep_.CoverTab[13901]++
								return nil, os.NewSyscallError("getsockopt", err)
//line /usr/local/go/src/net/fd_unix.go:151
			// _ = "end of CoverTab[13901]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:152
			_go_fuzz_dep_.CoverTab[13902]++
//line /usr/local/go/src/net/fd_unix.go:152
			// _ = "end of CoverTab[13902]"
//line /usr/local/go/src/net/fd_unix.go:152
		}
//line /usr/local/go/src/net/fd_unix.go:152
		// _ = "end of CoverTab[13893]"
//line /usr/local/go/src/net/fd_unix.go:152
		_go_fuzz_dep_.CoverTab[13894]++
							switch err := syscall.Errno(nerr); err {
		case syscall.EINPROGRESS, syscall.EALREADY, syscall.EINTR:
//line /usr/local/go/src/net/fd_unix.go:154
			_go_fuzz_dep_.CoverTab[13903]++
//line /usr/local/go/src/net/fd_unix.go:154
			// _ = "end of CoverTab[13903]"
		case syscall.EISCONN:
//line /usr/local/go/src/net/fd_unix.go:155
			_go_fuzz_dep_.CoverTab[13904]++
								return nil, nil
//line /usr/local/go/src/net/fd_unix.go:156
			// _ = "end of CoverTab[13904]"
		case syscall.Errno(0):
//line /usr/local/go/src/net/fd_unix.go:157
			_go_fuzz_dep_.CoverTab[13905]++

//line /usr/local/go/src/net/fd_unix.go:161
			if rsa, err := syscall.Getpeername(fd.pfd.Sysfd); err == nil {
//line /usr/local/go/src/net/fd_unix.go:161
				_go_fuzz_dep_.CoverTab[13907]++
									return rsa, nil
//line /usr/local/go/src/net/fd_unix.go:162
				// _ = "end of CoverTab[13907]"
			} else {
//line /usr/local/go/src/net/fd_unix.go:163
				_go_fuzz_dep_.CoverTab[13908]++
//line /usr/local/go/src/net/fd_unix.go:163
				// _ = "end of CoverTab[13908]"
//line /usr/local/go/src/net/fd_unix.go:163
			}
//line /usr/local/go/src/net/fd_unix.go:163
			// _ = "end of CoverTab[13905]"
		default:
//line /usr/local/go/src/net/fd_unix.go:164
			_go_fuzz_dep_.CoverTab[13906]++
								return nil, os.NewSyscallError("connect", err)
//line /usr/local/go/src/net/fd_unix.go:165
			// _ = "end of CoverTab[13906]"
		}
//line /usr/local/go/src/net/fd_unix.go:166
		// _ = "end of CoverTab[13894]"
//line /usr/local/go/src/net/fd_unix.go:166
		_go_fuzz_dep_.CoverTab[13895]++
							runtime.KeepAlive(fd)
//line /usr/local/go/src/net/fd_unix.go:167
		// _ = "end of CoverTab[13895]"
	}
//line /usr/local/go/src/net/fd_unix.go:168
	// _ = "end of CoverTab[13862]"
}

func (fd *netFD) accept() (netfd *netFD, err error) {
//line /usr/local/go/src/net/fd_unix.go:171
	_go_fuzz_dep_.CoverTab[13909]++
						d, rsa, errcall, err := fd.pfd.Accept()
						if err != nil {
//line /usr/local/go/src/net/fd_unix.go:173
		_go_fuzz_dep_.CoverTab[13913]++
							if errcall != "" {
//line /usr/local/go/src/net/fd_unix.go:174
			_go_fuzz_dep_.CoverTab[13915]++
								err = wrapSyscallError(errcall, err)
//line /usr/local/go/src/net/fd_unix.go:175
			// _ = "end of CoverTab[13915]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:176
			_go_fuzz_dep_.CoverTab[13916]++
//line /usr/local/go/src/net/fd_unix.go:176
			// _ = "end of CoverTab[13916]"
//line /usr/local/go/src/net/fd_unix.go:176
		}
//line /usr/local/go/src/net/fd_unix.go:176
		// _ = "end of CoverTab[13913]"
//line /usr/local/go/src/net/fd_unix.go:176
		_go_fuzz_dep_.CoverTab[13914]++
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:177
		// _ = "end of CoverTab[13914]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:178
		_go_fuzz_dep_.CoverTab[13917]++
//line /usr/local/go/src/net/fd_unix.go:178
		// _ = "end of CoverTab[13917]"
//line /usr/local/go/src/net/fd_unix.go:178
	}
//line /usr/local/go/src/net/fd_unix.go:178
	// _ = "end of CoverTab[13909]"
//line /usr/local/go/src/net/fd_unix.go:178
	_go_fuzz_dep_.CoverTab[13910]++

						if netfd, err = newFD(d, fd.family, fd.sotype, fd.net); err != nil {
//line /usr/local/go/src/net/fd_unix.go:180
		_go_fuzz_dep_.CoverTab[13918]++
							poll.CloseFunc(d)
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:182
		// _ = "end of CoverTab[13918]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:183
		_go_fuzz_dep_.CoverTab[13919]++
//line /usr/local/go/src/net/fd_unix.go:183
		// _ = "end of CoverTab[13919]"
//line /usr/local/go/src/net/fd_unix.go:183
	}
//line /usr/local/go/src/net/fd_unix.go:183
	// _ = "end of CoverTab[13910]"
//line /usr/local/go/src/net/fd_unix.go:183
	_go_fuzz_dep_.CoverTab[13911]++
						if err = netfd.init(); err != nil {
//line /usr/local/go/src/net/fd_unix.go:184
		_go_fuzz_dep_.CoverTab[13920]++
							netfd.Close()
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:186
		// _ = "end of CoverTab[13920]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:187
		_go_fuzz_dep_.CoverTab[13921]++
//line /usr/local/go/src/net/fd_unix.go:187
		// _ = "end of CoverTab[13921]"
//line /usr/local/go/src/net/fd_unix.go:187
	}
//line /usr/local/go/src/net/fd_unix.go:187
	// _ = "end of CoverTab[13911]"
//line /usr/local/go/src/net/fd_unix.go:187
	_go_fuzz_dep_.CoverTab[13912]++
						lsa, _ := syscall.Getsockname(netfd.pfd.Sysfd)
						netfd.setAddr(netfd.addrFunc()(lsa), netfd.addrFunc()(rsa))
						return netfd, nil
//line /usr/local/go/src/net/fd_unix.go:190
	// _ = "end of CoverTab[13912]"
}

// Defined in os package.
func newUnixFile(fd uintptr, name string) *os.File

func (fd *netFD) dup() (f *os.File, err error) {
//line /usr/local/go/src/net/fd_unix.go:196
	_go_fuzz_dep_.CoverTab[13922]++
						ns, call, err := fd.pfd.Dup()
						if err != nil {
//line /usr/local/go/src/net/fd_unix.go:198
		_go_fuzz_dep_.CoverTab[13924]++
							if call != "" {
//line /usr/local/go/src/net/fd_unix.go:199
			_go_fuzz_dep_.CoverTab[13926]++
								err = os.NewSyscallError(call, err)
//line /usr/local/go/src/net/fd_unix.go:200
			// _ = "end of CoverTab[13926]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:201
			_go_fuzz_dep_.CoverTab[13927]++
//line /usr/local/go/src/net/fd_unix.go:201
			// _ = "end of CoverTab[13927]"
//line /usr/local/go/src/net/fd_unix.go:201
		}
//line /usr/local/go/src/net/fd_unix.go:201
		// _ = "end of CoverTab[13924]"
//line /usr/local/go/src/net/fd_unix.go:201
		_go_fuzz_dep_.CoverTab[13925]++
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:202
		// _ = "end of CoverTab[13925]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:203
		_go_fuzz_dep_.CoverTab[13928]++
//line /usr/local/go/src/net/fd_unix.go:203
		// _ = "end of CoverTab[13928]"
//line /usr/local/go/src/net/fd_unix.go:203
	}
//line /usr/local/go/src/net/fd_unix.go:203
	// _ = "end of CoverTab[13922]"
//line /usr/local/go/src/net/fd_unix.go:203
	_go_fuzz_dep_.CoverTab[13923]++

						return newUnixFile(uintptr(ns), fd.name()), nil
//line /usr/local/go/src/net/fd_unix.go:205
	// _ = "end of CoverTab[13923]"
}

//line /usr/local/go/src/net/fd_unix.go:206
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/fd_unix.go:206
var _ = _go_fuzz_dep_.CoverTab
