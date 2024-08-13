// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

//line /usr/local/go/src/net/file_unix.go:7
package net

//line /usr/local/go/src/net/file_unix.go:7
import (
//line /usr/local/go/src/net/file_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/file_unix.go:7
)
//line /usr/local/go/src/net/file_unix.go:7
import (
//line /usr/local/go/src/net/file_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/file_unix.go:7
)

import (
	"internal/poll"
	"os"
	"syscall"
)

func dupSocket(f *os.File) (int, error) {
//line /usr/local/go/src/net/file_unix.go:15
	_go_fuzz_dep_.CoverTab[13943]++
						s, call, err := poll.DupCloseOnExec(int(f.Fd()))
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:17
		_go_fuzz_dep_.CoverTab[13946]++
							if call != "" {
//line /usr/local/go/src/net/file_unix.go:18
			_go_fuzz_dep_.CoverTab[13948]++
								err = os.NewSyscallError(call, err)
//line /usr/local/go/src/net/file_unix.go:19
			// _ = "end of CoverTab[13948]"
		} else {
//line /usr/local/go/src/net/file_unix.go:20
			_go_fuzz_dep_.CoverTab[13949]++
//line /usr/local/go/src/net/file_unix.go:20
			// _ = "end of CoverTab[13949]"
//line /usr/local/go/src/net/file_unix.go:20
		}
//line /usr/local/go/src/net/file_unix.go:20
		// _ = "end of CoverTab[13946]"
//line /usr/local/go/src/net/file_unix.go:20
		_go_fuzz_dep_.CoverTab[13947]++
							return -1, err
//line /usr/local/go/src/net/file_unix.go:21
		// _ = "end of CoverTab[13947]"
	} else {
//line /usr/local/go/src/net/file_unix.go:22
		_go_fuzz_dep_.CoverTab[13950]++
//line /usr/local/go/src/net/file_unix.go:22
		// _ = "end of CoverTab[13950]"
//line /usr/local/go/src/net/file_unix.go:22
	}
//line /usr/local/go/src/net/file_unix.go:22
	// _ = "end of CoverTab[13943]"
//line /usr/local/go/src/net/file_unix.go:22
	_go_fuzz_dep_.CoverTab[13944]++
						if err := syscall.SetNonblock(s, true); err != nil {
//line /usr/local/go/src/net/file_unix.go:23
		_go_fuzz_dep_.CoverTab[13951]++
							poll.CloseFunc(s)
							return -1, os.NewSyscallError("setnonblock", err)
//line /usr/local/go/src/net/file_unix.go:25
		// _ = "end of CoverTab[13951]"
	} else {
//line /usr/local/go/src/net/file_unix.go:26
		_go_fuzz_dep_.CoverTab[13952]++
//line /usr/local/go/src/net/file_unix.go:26
		// _ = "end of CoverTab[13952]"
//line /usr/local/go/src/net/file_unix.go:26
	}
//line /usr/local/go/src/net/file_unix.go:26
	// _ = "end of CoverTab[13944]"
//line /usr/local/go/src/net/file_unix.go:26
	_go_fuzz_dep_.CoverTab[13945]++
						return s, nil
//line /usr/local/go/src/net/file_unix.go:27
	// _ = "end of CoverTab[13945]"
}

func newFileFD(f *os.File) (*netFD, error) {
//line /usr/local/go/src/net/file_unix.go:30
	_go_fuzz_dep_.CoverTab[13953]++
						s, err := dupSocket(f)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:32
		_go_fuzz_dep_.CoverTab[13959]++
							return nil, err
//line /usr/local/go/src/net/file_unix.go:33
		// _ = "end of CoverTab[13959]"
	} else {
//line /usr/local/go/src/net/file_unix.go:34
		_go_fuzz_dep_.CoverTab[13960]++
//line /usr/local/go/src/net/file_unix.go:34
		// _ = "end of CoverTab[13960]"
//line /usr/local/go/src/net/file_unix.go:34
	}
//line /usr/local/go/src/net/file_unix.go:34
	// _ = "end of CoverTab[13953]"
//line /usr/local/go/src/net/file_unix.go:34
	_go_fuzz_dep_.CoverTab[13954]++
						family := syscall.AF_UNSPEC
						sotype, err := syscall.GetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_TYPE)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:37
		_go_fuzz_dep_.CoverTab[13961]++
							poll.CloseFunc(s)
							return nil, os.NewSyscallError("getsockopt", err)
//line /usr/local/go/src/net/file_unix.go:39
		// _ = "end of CoverTab[13961]"
	} else {
//line /usr/local/go/src/net/file_unix.go:40
		_go_fuzz_dep_.CoverTab[13962]++
//line /usr/local/go/src/net/file_unix.go:40
		// _ = "end of CoverTab[13962]"
//line /usr/local/go/src/net/file_unix.go:40
	}
//line /usr/local/go/src/net/file_unix.go:40
	// _ = "end of CoverTab[13954]"
//line /usr/local/go/src/net/file_unix.go:40
	_go_fuzz_dep_.CoverTab[13955]++
						lsa, _ := syscall.Getsockname(s)
						rsa, _ := syscall.Getpeername(s)
						switch lsa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/file_unix.go:44
		_go_fuzz_dep_.CoverTab[13963]++
							family = syscall.AF_INET
//line /usr/local/go/src/net/file_unix.go:45
		// _ = "end of CoverTab[13963]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/file_unix.go:46
		_go_fuzz_dep_.CoverTab[13964]++
							family = syscall.AF_INET6
//line /usr/local/go/src/net/file_unix.go:47
		// _ = "end of CoverTab[13964]"
	case *syscall.SockaddrUnix:
//line /usr/local/go/src/net/file_unix.go:48
		_go_fuzz_dep_.CoverTab[13965]++
							family = syscall.AF_UNIX
//line /usr/local/go/src/net/file_unix.go:49
		// _ = "end of CoverTab[13965]"
	default:
//line /usr/local/go/src/net/file_unix.go:50
		_go_fuzz_dep_.CoverTab[13966]++
							poll.CloseFunc(s)
							return nil, syscall.EPROTONOSUPPORT
//line /usr/local/go/src/net/file_unix.go:52
		// _ = "end of CoverTab[13966]"
	}
//line /usr/local/go/src/net/file_unix.go:53
	// _ = "end of CoverTab[13955]"
//line /usr/local/go/src/net/file_unix.go:53
	_go_fuzz_dep_.CoverTab[13956]++
						fd, err := newFD(s, family, sotype, "")
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:55
		_go_fuzz_dep_.CoverTab[13967]++
							poll.CloseFunc(s)
							return nil, err
//line /usr/local/go/src/net/file_unix.go:57
		// _ = "end of CoverTab[13967]"
	} else {
//line /usr/local/go/src/net/file_unix.go:58
		_go_fuzz_dep_.CoverTab[13968]++
//line /usr/local/go/src/net/file_unix.go:58
		// _ = "end of CoverTab[13968]"
//line /usr/local/go/src/net/file_unix.go:58
	}
//line /usr/local/go/src/net/file_unix.go:58
	// _ = "end of CoverTab[13956]"
//line /usr/local/go/src/net/file_unix.go:58
	_go_fuzz_dep_.CoverTab[13957]++
						laddr := fd.addrFunc()(lsa)
						raddr := fd.addrFunc()(rsa)
						fd.net = laddr.Network()
						if err := fd.init(); err != nil {
//line /usr/local/go/src/net/file_unix.go:62
		_go_fuzz_dep_.CoverTab[13969]++
							fd.Close()
							return nil, err
//line /usr/local/go/src/net/file_unix.go:64
		// _ = "end of CoverTab[13969]"
	} else {
//line /usr/local/go/src/net/file_unix.go:65
		_go_fuzz_dep_.CoverTab[13970]++
//line /usr/local/go/src/net/file_unix.go:65
		// _ = "end of CoverTab[13970]"
//line /usr/local/go/src/net/file_unix.go:65
	}
//line /usr/local/go/src/net/file_unix.go:65
	// _ = "end of CoverTab[13957]"
//line /usr/local/go/src/net/file_unix.go:65
	_go_fuzz_dep_.CoverTab[13958]++
						fd.setAddr(laddr, raddr)
						return fd, nil
//line /usr/local/go/src/net/file_unix.go:67
	// _ = "end of CoverTab[13958]"
}

func fileConn(f *os.File) (Conn, error) {
//line /usr/local/go/src/net/file_unix.go:70
	_go_fuzz_dep_.CoverTab[13971]++
						fd, err := newFileFD(f)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:72
		_go_fuzz_dep_.CoverTab[13974]++
							return nil, err
//line /usr/local/go/src/net/file_unix.go:73
		// _ = "end of CoverTab[13974]"
	} else {
//line /usr/local/go/src/net/file_unix.go:74
		_go_fuzz_dep_.CoverTab[13975]++
//line /usr/local/go/src/net/file_unix.go:74
		// _ = "end of CoverTab[13975]"
//line /usr/local/go/src/net/file_unix.go:74
	}
//line /usr/local/go/src/net/file_unix.go:74
	// _ = "end of CoverTab[13971]"
//line /usr/local/go/src/net/file_unix.go:74
	_go_fuzz_dep_.CoverTab[13972]++
						switch fd.laddr.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/file_unix.go:76
		_go_fuzz_dep_.CoverTab[13976]++
							return newTCPConn(fd, defaultTCPKeepAlive, testHookSetKeepAlive), nil
//line /usr/local/go/src/net/file_unix.go:77
		// _ = "end of CoverTab[13976]"
	case *UDPAddr:
//line /usr/local/go/src/net/file_unix.go:78
		_go_fuzz_dep_.CoverTab[13977]++
							return newUDPConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:79
		// _ = "end of CoverTab[13977]"
	case *IPAddr:
//line /usr/local/go/src/net/file_unix.go:80
		_go_fuzz_dep_.CoverTab[13978]++
							return newIPConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:81
		// _ = "end of CoverTab[13978]"
	case *UnixAddr:
//line /usr/local/go/src/net/file_unix.go:82
		_go_fuzz_dep_.CoverTab[13979]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:83
		// _ = "end of CoverTab[13979]"
	}
//line /usr/local/go/src/net/file_unix.go:84
	// _ = "end of CoverTab[13972]"
//line /usr/local/go/src/net/file_unix.go:84
	_go_fuzz_dep_.CoverTab[13973]++
						fd.Close()
						return nil, syscall.EINVAL
//line /usr/local/go/src/net/file_unix.go:86
	// _ = "end of CoverTab[13973]"
}

func fileListener(f *os.File) (Listener, error) {
//line /usr/local/go/src/net/file_unix.go:89
	_go_fuzz_dep_.CoverTab[13980]++
						fd, err := newFileFD(f)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:91
		_go_fuzz_dep_.CoverTab[13983]++
							return nil, err
//line /usr/local/go/src/net/file_unix.go:92
		// _ = "end of CoverTab[13983]"
	} else {
//line /usr/local/go/src/net/file_unix.go:93
		_go_fuzz_dep_.CoverTab[13984]++
//line /usr/local/go/src/net/file_unix.go:93
		// _ = "end of CoverTab[13984]"
//line /usr/local/go/src/net/file_unix.go:93
	}
//line /usr/local/go/src/net/file_unix.go:93
	// _ = "end of CoverTab[13980]"
//line /usr/local/go/src/net/file_unix.go:93
	_go_fuzz_dep_.CoverTab[13981]++
						switch laddr := fd.laddr.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/file_unix.go:95
		_go_fuzz_dep_.CoverTab[13985]++
							return &TCPListener{fd: fd}, nil
//line /usr/local/go/src/net/file_unix.go:96
		// _ = "end of CoverTab[13985]"
	case *UnixAddr:
//line /usr/local/go/src/net/file_unix.go:97
		_go_fuzz_dep_.CoverTab[13986]++
							return &UnixListener{fd: fd, path: laddr.Name, unlink: false}, nil
//line /usr/local/go/src/net/file_unix.go:98
		// _ = "end of CoverTab[13986]"
	}
//line /usr/local/go/src/net/file_unix.go:99
	// _ = "end of CoverTab[13981]"
//line /usr/local/go/src/net/file_unix.go:99
	_go_fuzz_dep_.CoverTab[13982]++
						fd.Close()
						return nil, syscall.EINVAL
//line /usr/local/go/src/net/file_unix.go:101
	// _ = "end of CoverTab[13982]"
}

func filePacketConn(f *os.File) (PacketConn, error) {
//line /usr/local/go/src/net/file_unix.go:104
	_go_fuzz_dep_.CoverTab[13987]++
						fd, err := newFileFD(f)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:106
		_go_fuzz_dep_.CoverTab[13990]++
							return nil, err
//line /usr/local/go/src/net/file_unix.go:107
		// _ = "end of CoverTab[13990]"
	} else {
//line /usr/local/go/src/net/file_unix.go:108
		_go_fuzz_dep_.CoverTab[13991]++
//line /usr/local/go/src/net/file_unix.go:108
		// _ = "end of CoverTab[13991]"
//line /usr/local/go/src/net/file_unix.go:108
	}
//line /usr/local/go/src/net/file_unix.go:108
	// _ = "end of CoverTab[13987]"
//line /usr/local/go/src/net/file_unix.go:108
	_go_fuzz_dep_.CoverTab[13988]++
						switch fd.laddr.(type) {
	case *UDPAddr:
//line /usr/local/go/src/net/file_unix.go:110
		_go_fuzz_dep_.CoverTab[13992]++
							return newUDPConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:111
		// _ = "end of CoverTab[13992]"
	case *IPAddr:
//line /usr/local/go/src/net/file_unix.go:112
		_go_fuzz_dep_.CoverTab[13993]++
							return newIPConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:113
		// _ = "end of CoverTab[13993]"
	case *UnixAddr:
//line /usr/local/go/src/net/file_unix.go:114
		_go_fuzz_dep_.CoverTab[13994]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:115
		// _ = "end of CoverTab[13994]"
	}
//line /usr/local/go/src/net/file_unix.go:116
	// _ = "end of CoverTab[13988]"
//line /usr/local/go/src/net/file_unix.go:116
	_go_fuzz_dep_.CoverTab[13989]++
						fd.Close()
						return nil, syscall.EINVAL
//line /usr/local/go/src/net/file_unix.go:118
	// _ = "end of CoverTab[13989]"
}

//line /usr/local/go/src/net/file_unix.go:119
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/file_unix.go:119
var _ = _go_fuzz_dep_.CoverTab
