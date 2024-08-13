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
	_go_fuzz_dep_.CoverTab[5553]++
						s, call, err := poll.DupCloseOnExec(int(f.Fd()))
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:17
		_go_fuzz_dep_.CoverTab[5556]++
							if call != "" {
//line /usr/local/go/src/net/file_unix.go:18
			_go_fuzz_dep_.CoverTab[5558]++
								err = os.NewSyscallError(call, err)
//line /usr/local/go/src/net/file_unix.go:19
			// _ = "end of CoverTab[5558]"
		} else {
//line /usr/local/go/src/net/file_unix.go:20
			_go_fuzz_dep_.CoverTab[5559]++
//line /usr/local/go/src/net/file_unix.go:20
			// _ = "end of CoverTab[5559]"
//line /usr/local/go/src/net/file_unix.go:20
		}
//line /usr/local/go/src/net/file_unix.go:20
		// _ = "end of CoverTab[5556]"
//line /usr/local/go/src/net/file_unix.go:20
		_go_fuzz_dep_.CoverTab[5557]++
							return -1, err
//line /usr/local/go/src/net/file_unix.go:21
		// _ = "end of CoverTab[5557]"
	} else {
//line /usr/local/go/src/net/file_unix.go:22
		_go_fuzz_dep_.CoverTab[5560]++
//line /usr/local/go/src/net/file_unix.go:22
		// _ = "end of CoverTab[5560]"
//line /usr/local/go/src/net/file_unix.go:22
	}
//line /usr/local/go/src/net/file_unix.go:22
	// _ = "end of CoverTab[5553]"
//line /usr/local/go/src/net/file_unix.go:22
	_go_fuzz_dep_.CoverTab[5554]++
						if err := syscall.SetNonblock(s, true); err != nil {
//line /usr/local/go/src/net/file_unix.go:23
		_go_fuzz_dep_.CoverTab[5561]++
							poll.CloseFunc(s)
							return -1, os.NewSyscallError("setnonblock", err)
//line /usr/local/go/src/net/file_unix.go:25
		// _ = "end of CoverTab[5561]"
	} else {
//line /usr/local/go/src/net/file_unix.go:26
		_go_fuzz_dep_.CoverTab[5562]++
//line /usr/local/go/src/net/file_unix.go:26
		// _ = "end of CoverTab[5562]"
//line /usr/local/go/src/net/file_unix.go:26
	}
//line /usr/local/go/src/net/file_unix.go:26
	// _ = "end of CoverTab[5554]"
//line /usr/local/go/src/net/file_unix.go:26
	_go_fuzz_dep_.CoverTab[5555]++
						return s, nil
//line /usr/local/go/src/net/file_unix.go:27
	// _ = "end of CoverTab[5555]"
}

func newFileFD(f *os.File) (*netFD, error) {
//line /usr/local/go/src/net/file_unix.go:30
	_go_fuzz_dep_.CoverTab[5563]++
						s, err := dupSocket(f)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:32
		_go_fuzz_dep_.CoverTab[5569]++
							return nil, err
//line /usr/local/go/src/net/file_unix.go:33
		// _ = "end of CoverTab[5569]"
	} else {
//line /usr/local/go/src/net/file_unix.go:34
		_go_fuzz_dep_.CoverTab[5570]++
//line /usr/local/go/src/net/file_unix.go:34
		// _ = "end of CoverTab[5570]"
//line /usr/local/go/src/net/file_unix.go:34
	}
//line /usr/local/go/src/net/file_unix.go:34
	// _ = "end of CoverTab[5563]"
//line /usr/local/go/src/net/file_unix.go:34
	_go_fuzz_dep_.CoverTab[5564]++
						family := syscall.AF_UNSPEC
						sotype, err := syscall.GetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_TYPE)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:37
		_go_fuzz_dep_.CoverTab[5571]++
							poll.CloseFunc(s)
							return nil, os.NewSyscallError("getsockopt", err)
//line /usr/local/go/src/net/file_unix.go:39
		// _ = "end of CoverTab[5571]"
	} else {
//line /usr/local/go/src/net/file_unix.go:40
		_go_fuzz_dep_.CoverTab[5572]++
//line /usr/local/go/src/net/file_unix.go:40
		// _ = "end of CoverTab[5572]"
//line /usr/local/go/src/net/file_unix.go:40
	}
//line /usr/local/go/src/net/file_unix.go:40
	// _ = "end of CoverTab[5564]"
//line /usr/local/go/src/net/file_unix.go:40
	_go_fuzz_dep_.CoverTab[5565]++
						lsa, _ := syscall.Getsockname(s)
						rsa, _ := syscall.Getpeername(s)
						switch lsa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/file_unix.go:44
		_go_fuzz_dep_.CoverTab[5573]++
							family = syscall.AF_INET
//line /usr/local/go/src/net/file_unix.go:45
		// _ = "end of CoverTab[5573]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/file_unix.go:46
		_go_fuzz_dep_.CoverTab[5574]++
							family = syscall.AF_INET6
//line /usr/local/go/src/net/file_unix.go:47
		// _ = "end of CoverTab[5574]"
	case *syscall.SockaddrUnix:
//line /usr/local/go/src/net/file_unix.go:48
		_go_fuzz_dep_.CoverTab[5575]++
							family = syscall.AF_UNIX
//line /usr/local/go/src/net/file_unix.go:49
		// _ = "end of CoverTab[5575]"
	default:
//line /usr/local/go/src/net/file_unix.go:50
		_go_fuzz_dep_.CoverTab[5576]++
							poll.CloseFunc(s)
							return nil, syscall.EPROTONOSUPPORT
//line /usr/local/go/src/net/file_unix.go:52
		// _ = "end of CoverTab[5576]"
	}
//line /usr/local/go/src/net/file_unix.go:53
	// _ = "end of CoverTab[5565]"
//line /usr/local/go/src/net/file_unix.go:53
	_go_fuzz_dep_.CoverTab[5566]++
						fd, err := newFD(s, family, sotype, "")
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:55
		_go_fuzz_dep_.CoverTab[5577]++
							poll.CloseFunc(s)
							return nil, err
//line /usr/local/go/src/net/file_unix.go:57
		// _ = "end of CoverTab[5577]"
	} else {
//line /usr/local/go/src/net/file_unix.go:58
		_go_fuzz_dep_.CoverTab[5578]++
//line /usr/local/go/src/net/file_unix.go:58
		// _ = "end of CoverTab[5578]"
//line /usr/local/go/src/net/file_unix.go:58
	}
//line /usr/local/go/src/net/file_unix.go:58
	// _ = "end of CoverTab[5566]"
//line /usr/local/go/src/net/file_unix.go:58
	_go_fuzz_dep_.CoverTab[5567]++
						laddr := fd.addrFunc()(lsa)
						raddr := fd.addrFunc()(rsa)
						fd.net = laddr.Network()
						if err := fd.init(); err != nil {
//line /usr/local/go/src/net/file_unix.go:62
		_go_fuzz_dep_.CoverTab[5579]++
							fd.Close()
							return nil, err
//line /usr/local/go/src/net/file_unix.go:64
		// _ = "end of CoverTab[5579]"
	} else {
//line /usr/local/go/src/net/file_unix.go:65
		_go_fuzz_dep_.CoverTab[5580]++
//line /usr/local/go/src/net/file_unix.go:65
		// _ = "end of CoverTab[5580]"
//line /usr/local/go/src/net/file_unix.go:65
	}
//line /usr/local/go/src/net/file_unix.go:65
	// _ = "end of CoverTab[5567]"
//line /usr/local/go/src/net/file_unix.go:65
	_go_fuzz_dep_.CoverTab[5568]++
						fd.setAddr(laddr, raddr)
						return fd, nil
//line /usr/local/go/src/net/file_unix.go:67
	// _ = "end of CoverTab[5568]"
}

func fileConn(f *os.File) (Conn, error) {
//line /usr/local/go/src/net/file_unix.go:70
	_go_fuzz_dep_.CoverTab[5581]++
						fd, err := newFileFD(f)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:72
		_go_fuzz_dep_.CoverTab[5584]++
							return nil, err
//line /usr/local/go/src/net/file_unix.go:73
		// _ = "end of CoverTab[5584]"
	} else {
//line /usr/local/go/src/net/file_unix.go:74
		_go_fuzz_dep_.CoverTab[5585]++
//line /usr/local/go/src/net/file_unix.go:74
		// _ = "end of CoverTab[5585]"
//line /usr/local/go/src/net/file_unix.go:74
	}
//line /usr/local/go/src/net/file_unix.go:74
	// _ = "end of CoverTab[5581]"
//line /usr/local/go/src/net/file_unix.go:74
	_go_fuzz_dep_.CoverTab[5582]++
						switch fd.laddr.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/file_unix.go:76
		_go_fuzz_dep_.CoverTab[5586]++
							return newTCPConn(fd, defaultTCPKeepAlive, testHookSetKeepAlive), nil
//line /usr/local/go/src/net/file_unix.go:77
		// _ = "end of CoverTab[5586]"
	case *UDPAddr:
//line /usr/local/go/src/net/file_unix.go:78
		_go_fuzz_dep_.CoverTab[5587]++
							return newUDPConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:79
		// _ = "end of CoverTab[5587]"
	case *IPAddr:
//line /usr/local/go/src/net/file_unix.go:80
		_go_fuzz_dep_.CoverTab[5588]++
							return newIPConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:81
		// _ = "end of CoverTab[5588]"
	case *UnixAddr:
//line /usr/local/go/src/net/file_unix.go:82
		_go_fuzz_dep_.CoverTab[5589]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:83
		// _ = "end of CoverTab[5589]"
	}
//line /usr/local/go/src/net/file_unix.go:84
	// _ = "end of CoverTab[5582]"
//line /usr/local/go/src/net/file_unix.go:84
	_go_fuzz_dep_.CoverTab[5583]++
						fd.Close()
						return nil, syscall.EINVAL
//line /usr/local/go/src/net/file_unix.go:86
	// _ = "end of CoverTab[5583]"
}

func fileListener(f *os.File) (Listener, error) {
//line /usr/local/go/src/net/file_unix.go:89
	_go_fuzz_dep_.CoverTab[5590]++
						fd, err := newFileFD(f)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:91
		_go_fuzz_dep_.CoverTab[5593]++
							return nil, err
//line /usr/local/go/src/net/file_unix.go:92
		// _ = "end of CoverTab[5593]"
	} else {
//line /usr/local/go/src/net/file_unix.go:93
		_go_fuzz_dep_.CoverTab[5594]++
//line /usr/local/go/src/net/file_unix.go:93
		// _ = "end of CoverTab[5594]"
//line /usr/local/go/src/net/file_unix.go:93
	}
//line /usr/local/go/src/net/file_unix.go:93
	// _ = "end of CoverTab[5590]"
//line /usr/local/go/src/net/file_unix.go:93
	_go_fuzz_dep_.CoverTab[5591]++
						switch laddr := fd.laddr.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/file_unix.go:95
		_go_fuzz_dep_.CoverTab[5595]++
							return &TCPListener{fd: fd}, nil
//line /usr/local/go/src/net/file_unix.go:96
		// _ = "end of CoverTab[5595]"
	case *UnixAddr:
//line /usr/local/go/src/net/file_unix.go:97
		_go_fuzz_dep_.CoverTab[5596]++
							return &UnixListener{fd: fd, path: laddr.Name, unlink: false}, nil
//line /usr/local/go/src/net/file_unix.go:98
		// _ = "end of CoverTab[5596]"
	}
//line /usr/local/go/src/net/file_unix.go:99
	// _ = "end of CoverTab[5591]"
//line /usr/local/go/src/net/file_unix.go:99
	_go_fuzz_dep_.CoverTab[5592]++
						fd.Close()
						return nil, syscall.EINVAL
//line /usr/local/go/src/net/file_unix.go:101
	// _ = "end of CoverTab[5592]"
}

func filePacketConn(f *os.File) (PacketConn, error) {
//line /usr/local/go/src/net/file_unix.go:104
	_go_fuzz_dep_.CoverTab[5597]++
						fd, err := newFileFD(f)
						if err != nil {
//line /usr/local/go/src/net/file_unix.go:106
		_go_fuzz_dep_.CoverTab[5600]++
							return nil, err
//line /usr/local/go/src/net/file_unix.go:107
		// _ = "end of CoverTab[5600]"
	} else {
//line /usr/local/go/src/net/file_unix.go:108
		_go_fuzz_dep_.CoverTab[5601]++
//line /usr/local/go/src/net/file_unix.go:108
		// _ = "end of CoverTab[5601]"
//line /usr/local/go/src/net/file_unix.go:108
	}
//line /usr/local/go/src/net/file_unix.go:108
	// _ = "end of CoverTab[5597]"
//line /usr/local/go/src/net/file_unix.go:108
	_go_fuzz_dep_.CoverTab[5598]++
						switch fd.laddr.(type) {
	case *UDPAddr:
//line /usr/local/go/src/net/file_unix.go:110
		_go_fuzz_dep_.CoverTab[5602]++
							return newUDPConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:111
		// _ = "end of CoverTab[5602]"
	case *IPAddr:
//line /usr/local/go/src/net/file_unix.go:112
		_go_fuzz_dep_.CoverTab[5603]++
							return newIPConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:113
		// _ = "end of CoverTab[5603]"
	case *UnixAddr:
//line /usr/local/go/src/net/file_unix.go:114
		_go_fuzz_dep_.CoverTab[5604]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/file_unix.go:115
		// _ = "end of CoverTab[5604]"
	}
//line /usr/local/go/src/net/file_unix.go:116
	// _ = "end of CoverTab[5598]"
//line /usr/local/go/src/net/file_unix.go:116
	_go_fuzz_dep_.CoverTab[5599]++
						fd.Close()
						return nil, syscall.EINVAL
//line /usr/local/go/src/net/file_unix.go:118
	// _ = "end of CoverTab[5599]"
}

//line /usr/local/go/src/net/file_unix.go:119
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/file_unix.go:119
var _ = _go_fuzz_dep_.CoverTab
