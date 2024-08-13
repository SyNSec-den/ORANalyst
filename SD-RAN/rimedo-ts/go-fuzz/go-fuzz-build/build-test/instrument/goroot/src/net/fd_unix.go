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
	_go_fuzz_dep_.CoverTab[5459]++
						ret := &netFD{
		pfd: poll.FD{
			Sysfd:		sysfd,
			IsStream:	sotype == syscall.SOCK_STREAM,
			ZeroReadIsEOF: sotype != syscall.SOCK_DGRAM && func() bool {
//line /usr/local/go/src/net/fd_unix.go:31
				_go_fuzz_dep_.CoverTab[5460]++
//line /usr/local/go/src/net/fd_unix.go:31
				return sotype != syscall.SOCK_RAW
//line /usr/local/go/src/net/fd_unix.go:31
				// _ = "end of CoverTab[5460]"
//line /usr/local/go/src/net/fd_unix.go:31
			}(),
		},
		family:	family,
		sotype:	sotype,
		net:	net,
	}
						return ret, nil
//line /usr/local/go/src/net/fd_unix.go:37
	// _ = "end of CoverTab[5459]"
}

func (fd *netFD) init() error {
	return fd.pfd.Init(fd.net, true)
}

func (fd *netFD) name() string {
//line /usr/local/go/src/net/fd_unix.go:44
	_go_fuzz_dep_.CoverTab[5461]++
						var ls, rs string
						if fd.laddr != nil {
//line /usr/local/go/src/net/fd_unix.go:46
		_go_fuzz_dep_.CoverTab[5464]++
							ls = fd.laddr.String()
//line /usr/local/go/src/net/fd_unix.go:47
		// _ = "end of CoverTab[5464]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:48
		_go_fuzz_dep_.CoverTab[5465]++
//line /usr/local/go/src/net/fd_unix.go:48
		// _ = "end of CoverTab[5465]"
//line /usr/local/go/src/net/fd_unix.go:48
	}
//line /usr/local/go/src/net/fd_unix.go:48
	// _ = "end of CoverTab[5461]"
//line /usr/local/go/src/net/fd_unix.go:48
	_go_fuzz_dep_.CoverTab[5462]++
						if fd.raddr != nil {
//line /usr/local/go/src/net/fd_unix.go:49
		_go_fuzz_dep_.CoverTab[5466]++
							rs = fd.raddr.String()
//line /usr/local/go/src/net/fd_unix.go:50
		// _ = "end of CoverTab[5466]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:51
		_go_fuzz_dep_.CoverTab[5467]++
//line /usr/local/go/src/net/fd_unix.go:51
		// _ = "end of CoverTab[5467]"
//line /usr/local/go/src/net/fd_unix.go:51
	}
//line /usr/local/go/src/net/fd_unix.go:51
	// _ = "end of CoverTab[5462]"
//line /usr/local/go/src/net/fd_unix.go:51
	_go_fuzz_dep_.CoverTab[5463]++
						return fd.net + ":" + ls + "->" + rs
//line /usr/local/go/src/net/fd_unix.go:52
	// _ = "end of CoverTab[5463]"
}

func (fd *netFD) connect(ctx context.Context, la, ra syscall.Sockaddr) (rsa syscall.Sockaddr, ret error) {
//line /usr/local/go/src/net/fd_unix.go:55
	_go_fuzz_dep_.CoverTab[5468]++

//line /usr/local/go/src/net/fd_unix.go:59
	switch err := connectFunc(fd.pfd.Sysfd, ra); err {
	case syscall.EINPROGRESS, syscall.EALREADY, syscall.EINTR:
//line /usr/local/go/src/net/fd_unix.go:60
		_go_fuzz_dep_.CoverTab[5473]++
//line /usr/local/go/src/net/fd_unix.go:60
		// _ = "end of CoverTab[5473]"
	case nil, syscall.EISCONN:
//line /usr/local/go/src/net/fd_unix.go:61
		_go_fuzz_dep_.CoverTab[5474]++
							select {
		case <-ctx.Done():
//line /usr/local/go/src/net/fd_unix.go:63
			_go_fuzz_dep_.CoverTab[5480]++
								return nil, mapErr(ctx.Err())
//line /usr/local/go/src/net/fd_unix.go:64
			// _ = "end of CoverTab[5480]"
		default:
//line /usr/local/go/src/net/fd_unix.go:65
			_go_fuzz_dep_.CoverTab[5481]++
//line /usr/local/go/src/net/fd_unix.go:65
			// _ = "end of CoverTab[5481]"
		}
//line /usr/local/go/src/net/fd_unix.go:66
		// _ = "end of CoverTab[5474]"
//line /usr/local/go/src/net/fd_unix.go:66
		_go_fuzz_dep_.CoverTab[5475]++
							if err := fd.pfd.Init(fd.net, true); err != nil {
//line /usr/local/go/src/net/fd_unix.go:67
			_go_fuzz_dep_.CoverTab[5482]++
								return nil, err
//line /usr/local/go/src/net/fd_unix.go:68
			// _ = "end of CoverTab[5482]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:69
			_go_fuzz_dep_.CoverTab[5483]++
//line /usr/local/go/src/net/fd_unix.go:69
			// _ = "end of CoverTab[5483]"
//line /usr/local/go/src/net/fd_unix.go:69
		}
//line /usr/local/go/src/net/fd_unix.go:69
		// _ = "end of CoverTab[5475]"
//line /usr/local/go/src/net/fd_unix.go:69
		_go_fuzz_dep_.CoverTab[5476]++
							runtime.KeepAlive(fd)
							return nil, nil
//line /usr/local/go/src/net/fd_unix.go:71
		// _ = "end of CoverTab[5476]"
	case syscall.EINVAL:
//line /usr/local/go/src/net/fd_unix.go:72
		_go_fuzz_dep_.CoverTab[5477]++

//line /usr/local/go/src/net/fd_unix.go:78
		if runtime.GOOS == "solaris" || func() bool {
//line /usr/local/go/src/net/fd_unix.go:78
			_go_fuzz_dep_.CoverTab[5484]++
//line /usr/local/go/src/net/fd_unix.go:78
			return runtime.GOOS == "illumos"
//line /usr/local/go/src/net/fd_unix.go:78
			// _ = "end of CoverTab[5484]"
//line /usr/local/go/src/net/fd_unix.go:78
		}() {
//line /usr/local/go/src/net/fd_unix.go:78
			_go_fuzz_dep_.CoverTab[5485]++
								return nil, nil
//line /usr/local/go/src/net/fd_unix.go:79
			// _ = "end of CoverTab[5485]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:80
			_go_fuzz_dep_.CoverTab[5486]++
//line /usr/local/go/src/net/fd_unix.go:80
			// _ = "end of CoverTab[5486]"
//line /usr/local/go/src/net/fd_unix.go:80
		}
//line /usr/local/go/src/net/fd_unix.go:80
		// _ = "end of CoverTab[5477]"
//line /usr/local/go/src/net/fd_unix.go:80
		_go_fuzz_dep_.CoverTab[5478]++
							fallthrough
//line /usr/local/go/src/net/fd_unix.go:81
		// _ = "end of CoverTab[5478]"
	default:
//line /usr/local/go/src/net/fd_unix.go:82
		_go_fuzz_dep_.CoverTab[5479]++
							return nil, os.NewSyscallError("connect", err)
//line /usr/local/go/src/net/fd_unix.go:83
		// _ = "end of CoverTab[5479]"
	}
//line /usr/local/go/src/net/fd_unix.go:84
	// _ = "end of CoverTab[5468]"
//line /usr/local/go/src/net/fd_unix.go:84
	_go_fuzz_dep_.CoverTab[5469]++
						if err := fd.pfd.Init(fd.net, true); err != nil {
//line /usr/local/go/src/net/fd_unix.go:85
		_go_fuzz_dep_.CoverTab[5487]++
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:86
		// _ = "end of CoverTab[5487]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:87
		_go_fuzz_dep_.CoverTab[5488]++
//line /usr/local/go/src/net/fd_unix.go:87
		// _ = "end of CoverTab[5488]"
//line /usr/local/go/src/net/fd_unix.go:87
	}
//line /usr/local/go/src/net/fd_unix.go:87
	// _ = "end of CoverTab[5469]"
//line /usr/local/go/src/net/fd_unix.go:87
	_go_fuzz_dep_.CoverTab[5470]++
						if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
//line /usr/local/go/src/net/fd_unix.go:88
		_go_fuzz_dep_.CoverTab[5489]++
							fd.pfd.SetWriteDeadline(deadline)
							defer fd.pfd.SetWriteDeadline(noDeadline)
//line /usr/local/go/src/net/fd_unix.go:90
		// _ = "end of CoverTab[5489]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:91
		_go_fuzz_dep_.CoverTab[5490]++
//line /usr/local/go/src/net/fd_unix.go:91
		// _ = "end of CoverTab[5490]"
//line /usr/local/go/src/net/fd_unix.go:91
	}
//line /usr/local/go/src/net/fd_unix.go:91
	// _ = "end of CoverTab[5470]"
//line /usr/local/go/src/net/fd_unix.go:91
	_go_fuzz_dep_.CoverTab[5471]++

//line /usr/local/go/src/net/fd_unix.go:98
	ctxDone := ctx.Done()
	if ctxDone != nil {
//line /usr/local/go/src/net/fd_unix.go:99
		_go_fuzz_dep_.CoverTab[5491]++

//line /usr/local/go/src/net/fd_unix.go:102
		done := make(chan struct{})
		interruptRes := make(chan error)
		defer func() {
//line /usr/local/go/src/net/fd_unix.go:104
			_go_fuzz_dep_.CoverTab[5493]++
								close(done)
								if ctxErr := <-interruptRes; ctxErr != nil && func() bool {
//line /usr/local/go/src/net/fd_unix.go:106
				_go_fuzz_dep_.CoverTab[5494]++
//line /usr/local/go/src/net/fd_unix.go:106
				return ret == nil
//line /usr/local/go/src/net/fd_unix.go:106
				// _ = "end of CoverTab[5494]"
//line /usr/local/go/src/net/fd_unix.go:106
			}() {
//line /usr/local/go/src/net/fd_unix.go:106
				_go_fuzz_dep_.CoverTab[5495]++

//line /usr/local/go/src/net/fd_unix.go:113
				ret = mapErr(ctxErr)
									fd.Close()
//line /usr/local/go/src/net/fd_unix.go:114
				// _ = "end of CoverTab[5495]"
			} else {
//line /usr/local/go/src/net/fd_unix.go:115
				_go_fuzz_dep_.CoverTab[5496]++
//line /usr/local/go/src/net/fd_unix.go:115
				// _ = "end of CoverTab[5496]"
//line /usr/local/go/src/net/fd_unix.go:115
			}
//line /usr/local/go/src/net/fd_unix.go:115
			// _ = "end of CoverTab[5493]"
		}()
//line /usr/local/go/src/net/fd_unix.go:116
		// _ = "end of CoverTab[5491]"
//line /usr/local/go/src/net/fd_unix.go:116
		_go_fuzz_dep_.CoverTab[5492]++
//line /usr/local/go/src/net/fd_unix.go:116
		_curRoutineNum9_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/fd_unix.go:116
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum9_)
							go func() {
//line /usr/local/go/src/net/fd_unix.go:117
			_go_fuzz_dep_.CoverTab[5497]++
//line /usr/local/go/src/net/fd_unix.go:117
			defer func() {
//line /usr/local/go/src/net/fd_unix.go:117
				_go_fuzz_dep_.CoverTab[5498]++
//line /usr/local/go/src/net/fd_unix.go:117
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum9_)
//line /usr/local/go/src/net/fd_unix.go:117
				// _ = "end of CoverTab[5498]"
//line /usr/local/go/src/net/fd_unix.go:117
			}()
								select {
			case <-ctxDone:
//line /usr/local/go/src/net/fd_unix.go:119
				_go_fuzz_dep_.CoverTab[5499]++

//line /usr/local/go/src/net/fd_unix.go:123
				fd.pfd.SetWriteDeadline(aLongTimeAgo)
									testHookCanceledDial()
									interruptRes <- ctx.Err()
//line /usr/local/go/src/net/fd_unix.go:125
				// _ = "end of CoverTab[5499]"
			case <-done:
//line /usr/local/go/src/net/fd_unix.go:126
				_go_fuzz_dep_.CoverTab[5500]++
									interruptRes <- nil
//line /usr/local/go/src/net/fd_unix.go:127
				// _ = "end of CoverTab[5500]"
			}
//line /usr/local/go/src/net/fd_unix.go:128
			// _ = "end of CoverTab[5497]"
		}()
//line /usr/local/go/src/net/fd_unix.go:129
		// _ = "end of CoverTab[5492]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:130
		_go_fuzz_dep_.CoverTab[5501]++
//line /usr/local/go/src/net/fd_unix.go:130
		// _ = "end of CoverTab[5501]"
//line /usr/local/go/src/net/fd_unix.go:130
	}
//line /usr/local/go/src/net/fd_unix.go:130
	// _ = "end of CoverTab[5471]"
//line /usr/local/go/src/net/fd_unix.go:130
	_go_fuzz_dep_.CoverTab[5472]++

						for {
//line /usr/local/go/src/net/fd_unix.go:132
		_go_fuzz_dep_.CoverTab[5502]++

//line /usr/local/go/src/net/fd_unix.go:141
		if err := fd.pfd.WaitWrite(); err != nil {
//line /usr/local/go/src/net/fd_unix.go:141
			_go_fuzz_dep_.CoverTab[5506]++
								select {
			case <-ctxDone:
//line /usr/local/go/src/net/fd_unix.go:143
				_go_fuzz_dep_.CoverTab[5508]++
									return nil, mapErr(ctx.Err())
//line /usr/local/go/src/net/fd_unix.go:144
				// _ = "end of CoverTab[5508]"
			default:
//line /usr/local/go/src/net/fd_unix.go:145
				_go_fuzz_dep_.CoverTab[5509]++
//line /usr/local/go/src/net/fd_unix.go:145
				// _ = "end of CoverTab[5509]"
			}
//line /usr/local/go/src/net/fd_unix.go:146
			// _ = "end of CoverTab[5506]"
//line /usr/local/go/src/net/fd_unix.go:146
			_go_fuzz_dep_.CoverTab[5507]++
								return nil, err
//line /usr/local/go/src/net/fd_unix.go:147
			// _ = "end of CoverTab[5507]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:148
			_go_fuzz_dep_.CoverTab[5510]++
//line /usr/local/go/src/net/fd_unix.go:148
			// _ = "end of CoverTab[5510]"
//line /usr/local/go/src/net/fd_unix.go:148
		}
//line /usr/local/go/src/net/fd_unix.go:148
		// _ = "end of CoverTab[5502]"
//line /usr/local/go/src/net/fd_unix.go:148
		_go_fuzz_dep_.CoverTab[5503]++
							nerr, err := getsockoptIntFunc(fd.pfd.Sysfd, syscall.SOL_SOCKET, syscall.SO_ERROR)
							if err != nil {
//line /usr/local/go/src/net/fd_unix.go:150
			_go_fuzz_dep_.CoverTab[5511]++
								return nil, os.NewSyscallError("getsockopt", err)
//line /usr/local/go/src/net/fd_unix.go:151
			// _ = "end of CoverTab[5511]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:152
			_go_fuzz_dep_.CoverTab[5512]++
//line /usr/local/go/src/net/fd_unix.go:152
			// _ = "end of CoverTab[5512]"
//line /usr/local/go/src/net/fd_unix.go:152
		}
//line /usr/local/go/src/net/fd_unix.go:152
		// _ = "end of CoverTab[5503]"
//line /usr/local/go/src/net/fd_unix.go:152
		_go_fuzz_dep_.CoverTab[5504]++
							switch err := syscall.Errno(nerr); err {
		case syscall.EINPROGRESS, syscall.EALREADY, syscall.EINTR:
//line /usr/local/go/src/net/fd_unix.go:154
			_go_fuzz_dep_.CoverTab[5513]++
//line /usr/local/go/src/net/fd_unix.go:154
			// _ = "end of CoverTab[5513]"
		case syscall.EISCONN:
//line /usr/local/go/src/net/fd_unix.go:155
			_go_fuzz_dep_.CoverTab[5514]++
								return nil, nil
//line /usr/local/go/src/net/fd_unix.go:156
			// _ = "end of CoverTab[5514]"
		case syscall.Errno(0):
//line /usr/local/go/src/net/fd_unix.go:157
			_go_fuzz_dep_.CoverTab[5515]++

//line /usr/local/go/src/net/fd_unix.go:161
			if rsa, err := syscall.Getpeername(fd.pfd.Sysfd); err == nil {
//line /usr/local/go/src/net/fd_unix.go:161
				_go_fuzz_dep_.CoverTab[5517]++
									return rsa, nil
//line /usr/local/go/src/net/fd_unix.go:162
				// _ = "end of CoverTab[5517]"
			} else {
//line /usr/local/go/src/net/fd_unix.go:163
				_go_fuzz_dep_.CoverTab[5518]++
//line /usr/local/go/src/net/fd_unix.go:163
				// _ = "end of CoverTab[5518]"
//line /usr/local/go/src/net/fd_unix.go:163
			}
//line /usr/local/go/src/net/fd_unix.go:163
			// _ = "end of CoverTab[5515]"
		default:
//line /usr/local/go/src/net/fd_unix.go:164
			_go_fuzz_dep_.CoverTab[5516]++
								return nil, os.NewSyscallError("connect", err)
//line /usr/local/go/src/net/fd_unix.go:165
			// _ = "end of CoverTab[5516]"
		}
//line /usr/local/go/src/net/fd_unix.go:166
		// _ = "end of CoverTab[5504]"
//line /usr/local/go/src/net/fd_unix.go:166
		_go_fuzz_dep_.CoverTab[5505]++
							runtime.KeepAlive(fd)
//line /usr/local/go/src/net/fd_unix.go:167
		// _ = "end of CoverTab[5505]"
	}
//line /usr/local/go/src/net/fd_unix.go:168
	// _ = "end of CoverTab[5472]"
}

func (fd *netFD) accept() (netfd *netFD, err error) {
//line /usr/local/go/src/net/fd_unix.go:171
	_go_fuzz_dep_.CoverTab[5519]++
						d, rsa, errcall, err := fd.pfd.Accept()
						if err != nil {
//line /usr/local/go/src/net/fd_unix.go:173
		_go_fuzz_dep_.CoverTab[5523]++
							if errcall != "" {
//line /usr/local/go/src/net/fd_unix.go:174
			_go_fuzz_dep_.CoverTab[5525]++
								err = wrapSyscallError(errcall, err)
//line /usr/local/go/src/net/fd_unix.go:175
			// _ = "end of CoverTab[5525]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:176
			_go_fuzz_dep_.CoverTab[5526]++
//line /usr/local/go/src/net/fd_unix.go:176
			// _ = "end of CoverTab[5526]"
//line /usr/local/go/src/net/fd_unix.go:176
		}
//line /usr/local/go/src/net/fd_unix.go:176
		// _ = "end of CoverTab[5523]"
//line /usr/local/go/src/net/fd_unix.go:176
		_go_fuzz_dep_.CoverTab[5524]++
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:177
		// _ = "end of CoverTab[5524]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:178
		_go_fuzz_dep_.CoverTab[5527]++
//line /usr/local/go/src/net/fd_unix.go:178
		// _ = "end of CoverTab[5527]"
//line /usr/local/go/src/net/fd_unix.go:178
	}
//line /usr/local/go/src/net/fd_unix.go:178
	// _ = "end of CoverTab[5519]"
//line /usr/local/go/src/net/fd_unix.go:178
	_go_fuzz_dep_.CoverTab[5520]++

						if netfd, err = newFD(d, fd.family, fd.sotype, fd.net); err != nil {
//line /usr/local/go/src/net/fd_unix.go:180
		_go_fuzz_dep_.CoverTab[5528]++
							poll.CloseFunc(d)
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:182
		// _ = "end of CoverTab[5528]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:183
		_go_fuzz_dep_.CoverTab[5529]++
//line /usr/local/go/src/net/fd_unix.go:183
		// _ = "end of CoverTab[5529]"
//line /usr/local/go/src/net/fd_unix.go:183
	}
//line /usr/local/go/src/net/fd_unix.go:183
	// _ = "end of CoverTab[5520]"
//line /usr/local/go/src/net/fd_unix.go:183
	_go_fuzz_dep_.CoverTab[5521]++
						if err = netfd.init(); err != nil {
//line /usr/local/go/src/net/fd_unix.go:184
		_go_fuzz_dep_.CoverTab[5530]++
							netfd.Close()
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:186
		// _ = "end of CoverTab[5530]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:187
		_go_fuzz_dep_.CoverTab[5531]++
//line /usr/local/go/src/net/fd_unix.go:187
		// _ = "end of CoverTab[5531]"
//line /usr/local/go/src/net/fd_unix.go:187
	}
//line /usr/local/go/src/net/fd_unix.go:187
	// _ = "end of CoverTab[5521]"
//line /usr/local/go/src/net/fd_unix.go:187
	_go_fuzz_dep_.CoverTab[5522]++
						lsa, _ := syscall.Getsockname(netfd.pfd.Sysfd)
						netfd.setAddr(netfd.addrFunc()(lsa), netfd.addrFunc()(rsa))
						return netfd, nil
//line /usr/local/go/src/net/fd_unix.go:190
	// _ = "end of CoverTab[5522]"
}

// Defined in os package.
func newUnixFile(fd uintptr, name string) *os.File

func (fd *netFD) dup() (f *os.File, err error) {
//line /usr/local/go/src/net/fd_unix.go:196
	_go_fuzz_dep_.CoverTab[5532]++
						ns, call, err := fd.pfd.Dup()
						if err != nil {
//line /usr/local/go/src/net/fd_unix.go:198
		_go_fuzz_dep_.CoverTab[5534]++
							if call != "" {
//line /usr/local/go/src/net/fd_unix.go:199
			_go_fuzz_dep_.CoverTab[5536]++
								err = os.NewSyscallError(call, err)
//line /usr/local/go/src/net/fd_unix.go:200
			// _ = "end of CoverTab[5536]"
		} else {
//line /usr/local/go/src/net/fd_unix.go:201
			_go_fuzz_dep_.CoverTab[5537]++
//line /usr/local/go/src/net/fd_unix.go:201
			// _ = "end of CoverTab[5537]"
//line /usr/local/go/src/net/fd_unix.go:201
		}
//line /usr/local/go/src/net/fd_unix.go:201
		// _ = "end of CoverTab[5534]"
//line /usr/local/go/src/net/fd_unix.go:201
		_go_fuzz_dep_.CoverTab[5535]++
							return nil, err
//line /usr/local/go/src/net/fd_unix.go:202
		// _ = "end of CoverTab[5535]"
	} else {
//line /usr/local/go/src/net/fd_unix.go:203
		_go_fuzz_dep_.CoverTab[5538]++
//line /usr/local/go/src/net/fd_unix.go:203
		// _ = "end of CoverTab[5538]"
//line /usr/local/go/src/net/fd_unix.go:203
	}
//line /usr/local/go/src/net/fd_unix.go:203
	// _ = "end of CoverTab[5532]"
//line /usr/local/go/src/net/fd_unix.go:203
	_go_fuzz_dep_.CoverTab[5533]++

						return newUnixFile(uintptr(ns), fd.name()), nil
//line /usr/local/go/src/net/fd_unix.go:205
	// _ = "end of CoverTab[5533]"
}

//line /usr/local/go/src/net/fd_unix.go:206
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/fd_unix.go:206
var _ = _go_fuzz_dep_.CoverTab
