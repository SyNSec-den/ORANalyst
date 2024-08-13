// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

//line /snap/go/10455/src/net/file_unix.go:7
package net

//line /snap/go/10455/src/net/file_unix.go:7
import (
//line /snap/go/10455/src/net/file_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/file_unix.go:7
)
//line /snap/go/10455/src/net/file_unix.go:7
import (
//line /snap/go/10455/src/net/file_unix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/file_unix.go:7
)

import (
	"internal/poll"
	"os"
	"syscall"
)

func dupSocket(f *os.File) (int, error) {
//line /snap/go/10455/src/net/file_unix.go:15
	_go_fuzz_dep_.CoverTab[5929]++
						s, call, err := poll.DupCloseOnExec(int(f.Fd()))
						if err != nil {
//line /snap/go/10455/src/net/file_unix.go:17
		_go_fuzz_dep_.CoverTab[528297]++
//line /snap/go/10455/src/net/file_unix.go:17
		_go_fuzz_dep_.CoverTab[5932]++
							if call != "" {
//line /snap/go/10455/src/net/file_unix.go:18
			_go_fuzz_dep_.CoverTab[528299]++
//line /snap/go/10455/src/net/file_unix.go:18
			_go_fuzz_dep_.CoverTab[5934]++
								err = os.NewSyscallError(call, err)
//line /snap/go/10455/src/net/file_unix.go:19
			// _ = "end of CoverTab[5934]"
		} else {
//line /snap/go/10455/src/net/file_unix.go:20
			_go_fuzz_dep_.CoverTab[528300]++
//line /snap/go/10455/src/net/file_unix.go:20
			_go_fuzz_dep_.CoverTab[5935]++
//line /snap/go/10455/src/net/file_unix.go:20
			// _ = "end of CoverTab[5935]"
//line /snap/go/10455/src/net/file_unix.go:20
		}
//line /snap/go/10455/src/net/file_unix.go:20
		// _ = "end of CoverTab[5932]"
//line /snap/go/10455/src/net/file_unix.go:20
		_go_fuzz_dep_.CoverTab[5933]++
							return -1, err
//line /snap/go/10455/src/net/file_unix.go:21
		// _ = "end of CoverTab[5933]"
	} else {
//line /snap/go/10455/src/net/file_unix.go:22
		_go_fuzz_dep_.CoverTab[528298]++
//line /snap/go/10455/src/net/file_unix.go:22
		_go_fuzz_dep_.CoverTab[5936]++
//line /snap/go/10455/src/net/file_unix.go:22
		// _ = "end of CoverTab[5936]"
//line /snap/go/10455/src/net/file_unix.go:22
	}
//line /snap/go/10455/src/net/file_unix.go:22
	// _ = "end of CoverTab[5929]"
//line /snap/go/10455/src/net/file_unix.go:22
	_go_fuzz_dep_.CoverTab[5930]++
						if err := syscall.SetNonblock(s, true); err != nil {
//line /snap/go/10455/src/net/file_unix.go:23
		_go_fuzz_dep_.CoverTab[528301]++
//line /snap/go/10455/src/net/file_unix.go:23
		_go_fuzz_dep_.CoverTab[5937]++
							poll.CloseFunc(s)
							return -1, os.NewSyscallError("setnonblock", err)
//line /snap/go/10455/src/net/file_unix.go:25
		// _ = "end of CoverTab[5937]"
	} else {
//line /snap/go/10455/src/net/file_unix.go:26
		_go_fuzz_dep_.CoverTab[528302]++
//line /snap/go/10455/src/net/file_unix.go:26
		_go_fuzz_dep_.CoverTab[5938]++
//line /snap/go/10455/src/net/file_unix.go:26
		// _ = "end of CoverTab[5938]"
//line /snap/go/10455/src/net/file_unix.go:26
	}
//line /snap/go/10455/src/net/file_unix.go:26
	// _ = "end of CoverTab[5930]"
//line /snap/go/10455/src/net/file_unix.go:26
	_go_fuzz_dep_.CoverTab[5931]++
						return s, nil
//line /snap/go/10455/src/net/file_unix.go:27
	// _ = "end of CoverTab[5931]"
}

func newFileFD(f *os.File) (*netFD, error) {
//line /snap/go/10455/src/net/file_unix.go:30
	_go_fuzz_dep_.CoverTab[5939]++
						s, err := dupSocket(f)
						if err != nil {
//line /snap/go/10455/src/net/file_unix.go:32
		_go_fuzz_dep_.CoverTab[528303]++
//line /snap/go/10455/src/net/file_unix.go:32
		_go_fuzz_dep_.CoverTab[5945]++
							return nil, err
//line /snap/go/10455/src/net/file_unix.go:33
		// _ = "end of CoverTab[5945]"
	} else {
//line /snap/go/10455/src/net/file_unix.go:34
		_go_fuzz_dep_.CoverTab[528304]++
//line /snap/go/10455/src/net/file_unix.go:34
		_go_fuzz_dep_.CoverTab[5946]++
//line /snap/go/10455/src/net/file_unix.go:34
		// _ = "end of CoverTab[5946]"
//line /snap/go/10455/src/net/file_unix.go:34
	}
//line /snap/go/10455/src/net/file_unix.go:34
	// _ = "end of CoverTab[5939]"
//line /snap/go/10455/src/net/file_unix.go:34
	_go_fuzz_dep_.CoverTab[5940]++
						family := syscall.AF_UNSPEC
						sotype, err := syscall.GetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_TYPE)
						if err != nil {
//line /snap/go/10455/src/net/file_unix.go:37
		_go_fuzz_dep_.CoverTab[528305]++
//line /snap/go/10455/src/net/file_unix.go:37
		_go_fuzz_dep_.CoverTab[5947]++
							poll.CloseFunc(s)
							return nil, os.NewSyscallError("getsockopt", err)
//line /snap/go/10455/src/net/file_unix.go:39
		// _ = "end of CoverTab[5947]"
	} else {
//line /snap/go/10455/src/net/file_unix.go:40
		_go_fuzz_dep_.CoverTab[528306]++
//line /snap/go/10455/src/net/file_unix.go:40
		_go_fuzz_dep_.CoverTab[5948]++
//line /snap/go/10455/src/net/file_unix.go:40
		// _ = "end of CoverTab[5948]"
//line /snap/go/10455/src/net/file_unix.go:40
	}
//line /snap/go/10455/src/net/file_unix.go:40
	// _ = "end of CoverTab[5940]"
//line /snap/go/10455/src/net/file_unix.go:40
	_go_fuzz_dep_.CoverTab[5941]++
						lsa, _ := syscall.Getsockname(s)
						rsa, _ := syscall.Getpeername(s)
						switch lsa.(type) {
	case *syscall.SockaddrInet4:
//line /snap/go/10455/src/net/file_unix.go:44
		_go_fuzz_dep_.CoverTab[528307]++
//line /snap/go/10455/src/net/file_unix.go:44
		_go_fuzz_dep_.CoverTab[5949]++
							family = syscall.AF_INET
//line /snap/go/10455/src/net/file_unix.go:45
		// _ = "end of CoverTab[5949]"
	case *syscall.SockaddrInet6:
//line /snap/go/10455/src/net/file_unix.go:46
		_go_fuzz_dep_.CoverTab[528308]++
//line /snap/go/10455/src/net/file_unix.go:46
		_go_fuzz_dep_.CoverTab[5950]++
							family = syscall.AF_INET6
//line /snap/go/10455/src/net/file_unix.go:47
		// _ = "end of CoverTab[5950]"
	case *syscall.SockaddrUnix:
//line /snap/go/10455/src/net/file_unix.go:48
		_go_fuzz_dep_.CoverTab[528309]++
//line /snap/go/10455/src/net/file_unix.go:48
		_go_fuzz_dep_.CoverTab[5951]++
							family = syscall.AF_UNIX
//line /snap/go/10455/src/net/file_unix.go:49
		// _ = "end of CoverTab[5951]"
	default:
//line /snap/go/10455/src/net/file_unix.go:50
		_go_fuzz_dep_.CoverTab[528310]++
//line /snap/go/10455/src/net/file_unix.go:50
		_go_fuzz_dep_.CoverTab[5952]++
							poll.CloseFunc(s)
							return nil, syscall.EPROTONOSUPPORT
//line /snap/go/10455/src/net/file_unix.go:52
		// _ = "end of CoverTab[5952]"
	}
//line /snap/go/10455/src/net/file_unix.go:53
	// _ = "end of CoverTab[5941]"
//line /snap/go/10455/src/net/file_unix.go:53
	_go_fuzz_dep_.CoverTab[5942]++
						fd, err := newFD(s, family, sotype, "")
						if err != nil {
//line /snap/go/10455/src/net/file_unix.go:55
		_go_fuzz_dep_.CoverTab[528311]++
//line /snap/go/10455/src/net/file_unix.go:55
		_go_fuzz_dep_.CoverTab[5953]++
							poll.CloseFunc(s)
							return nil, err
//line /snap/go/10455/src/net/file_unix.go:57
		// _ = "end of CoverTab[5953]"
	} else {
//line /snap/go/10455/src/net/file_unix.go:58
		_go_fuzz_dep_.CoverTab[528312]++
//line /snap/go/10455/src/net/file_unix.go:58
		_go_fuzz_dep_.CoverTab[5954]++
//line /snap/go/10455/src/net/file_unix.go:58
		// _ = "end of CoverTab[5954]"
//line /snap/go/10455/src/net/file_unix.go:58
	}
//line /snap/go/10455/src/net/file_unix.go:58
	// _ = "end of CoverTab[5942]"
//line /snap/go/10455/src/net/file_unix.go:58
	_go_fuzz_dep_.CoverTab[5943]++
						laddr := fd.addrFunc()(lsa)
						raddr := fd.addrFunc()(rsa)
						fd.net = laddr.Network()
						if err := fd.init(); err != nil {
//line /snap/go/10455/src/net/file_unix.go:62
		_go_fuzz_dep_.CoverTab[528313]++
//line /snap/go/10455/src/net/file_unix.go:62
		_go_fuzz_dep_.CoverTab[5955]++
							fd.Close()
							return nil, err
//line /snap/go/10455/src/net/file_unix.go:64
		// _ = "end of CoverTab[5955]"
	} else {
//line /snap/go/10455/src/net/file_unix.go:65
		_go_fuzz_dep_.CoverTab[528314]++
//line /snap/go/10455/src/net/file_unix.go:65
		_go_fuzz_dep_.CoverTab[5956]++
//line /snap/go/10455/src/net/file_unix.go:65
		// _ = "end of CoverTab[5956]"
//line /snap/go/10455/src/net/file_unix.go:65
	}
//line /snap/go/10455/src/net/file_unix.go:65
	// _ = "end of CoverTab[5943]"
//line /snap/go/10455/src/net/file_unix.go:65
	_go_fuzz_dep_.CoverTab[5944]++
						fd.setAddr(laddr, raddr)
						return fd, nil
//line /snap/go/10455/src/net/file_unix.go:67
	// _ = "end of CoverTab[5944]"
}

func fileConn(f *os.File) (Conn, error) {
//line /snap/go/10455/src/net/file_unix.go:70
	_go_fuzz_dep_.CoverTab[5957]++
						fd, err := newFileFD(f)
						if err != nil {
//line /snap/go/10455/src/net/file_unix.go:72
		_go_fuzz_dep_.CoverTab[528315]++
//line /snap/go/10455/src/net/file_unix.go:72
		_go_fuzz_dep_.CoverTab[5960]++
							return nil, err
//line /snap/go/10455/src/net/file_unix.go:73
		// _ = "end of CoverTab[5960]"
	} else {
//line /snap/go/10455/src/net/file_unix.go:74
		_go_fuzz_dep_.CoverTab[528316]++
//line /snap/go/10455/src/net/file_unix.go:74
		_go_fuzz_dep_.CoverTab[5961]++
//line /snap/go/10455/src/net/file_unix.go:74
		// _ = "end of CoverTab[5961]"
//line /snap/go/10455/src/net/file_unix.go:74
	}
//line /snap/go/10455/src/net/file_unix.go:74
	// _ = "end of CoverTab[5957]"
//line /snap/go/10455/src/net/file_unix.go:74
	_go_fuzz_dep_.CoverTab[5958]++
						switch fd.laddr.(type) {
	case *TCPAddr:
//line /snap/go/10455/src/net/file_unix.go:76
		_go_fuzz_dep_.CoverTab[528317]++
//line /snap/go/10455/src/net/file_unix.go:76
		_go_fuzz_dep_.CoverTab[5962]++
							return newTCPConn(fd, defaultTCPKeepAlive, testHookSetKeepAlive), nil
//line /snap/go/10455/src/net/file_unix.go:77
		// _ = "end of CoverTab[5962]"
	case *UDPAddr:
//line /snap/go/10455/src/net/file_unix.go:78
		_go_fuzz_dep_.CoverTab[528318]++
//line /snap/go/10455/src/net/file_unix.go:78
		_go_fuzz_dep_.CoverTab[5963]++
							return newUDPConn(fd), nil
//line /snap/go/10455/src/net/file_unix.go:79
		// _ = "end of CoverTab[5963]"
	case *IPAddr:
//line /snap/go/10455/src/net/file_unix.go:80
		_go_fuzz_dep_.CoverTab[528319]++
//line /snap/go/10455/src/net/file_unix.go:80
		_go_fuzz_dep_.CoverTab[5964]++
							return newIPConn(fd), nil
//line /snap/go/10455/src/net/file_unix.go:81
		// _ = "end of CoverTab[5964]"
	case *UnixAddr:
//line /snap/go/10455/src/net/file_unix.go:82
		_go_fuzz_dep_.CoverTab[528320]++
//line /snap/go/10455/src/net/file_unix.go:82
		_go_fuzz_dep_.CoverTab[5965]++
							return newUnixConn(fd), nil
//line /snap/go/10455/src/net/file_unix.go:83
		// _ = "end of CoverTab[5965]"
	}
//line /snap/go/10455/src/net/file_unix.go:84
	// _ = "end of CoverTab[5958]"
//line /snap/go/10455/src/net/file_unix.go:84
	_go_fuzz_dep_.CoverTab[5959]++
						fd.Close()
						return nil, syscall.EINVAL
//line /snap/go/10455/src/net/file_unix.go:86
	// _ = "end of CoverTab[5959]"
}

func fileListener(f *os.File) (Listener, error) {
//line /snap/go/10455/src/net/file_unix.go:89
	_go_fuzz_dep_.CoverTab[5966]++
						fd, err := newFileFD(f)
						if err != nil {
//line /snap/go/10455/src/net/file_unix.go:91
		_go_fuzz_dep_.CoverTab[528321]++
//line /snap/go/10455/src/net/file_unix.go:91
		_go_fuzz_dep_.CoverTab[5969]++
							return nil, err
//line /snap/go/10455/src/net/file_unix.go:92
		// _ = "end of CoverTab[5969]"
	} else {
//line /snap/go/10455/src/net/file_unix.go:93
		_go_fuzz_dep_.CoverTab[528322]++
//line /snap/go/10455/src/net/file_unix.go:93
		_go_fuzz_dep_.CoverTab[5970]++
//line /snap/go/10455/src/net/file_unix.go:93
		// _ = "end of CoverTab[5970]"
//line /snap/go/10455/src/net/file_unix.go:93
	}
//line /snap/go/10455/src/net/file_unix.go:93
	// _ = "end of CoverTab[5966]"
//line /snap/go/10455/src/net/file_unix.go:93
	_go_fuzz_dep_.CoverTab[5967]++
						switch laddr := fd.laddr.(type) {
	case *TCPAddr:
//line /snap/go/10455/src/net/file_unix.go:95
		_go_fuzz_dep_.CoverTab[528323]++
//line /snap/go/10455/src/net/file_unix.go:95
		_go_fuzz_dep_.CoverTab[5971]++
							return &TCPListener{fd: fd}, nil
//line /snap/go/10455/src/net/file_unix.go:96
		// _ = "end of CoverTab[5971]"
	case *UnixAddr:
//line /snap/go/10455/src/net/file_unix.go:97
		_go_fuzz_dep_.CoverTab[528324]++
//line /snap/go/10455/src/net/file_unix.go:97
		_go_fuzz_dep_.CoverTab[5972]++
							return &UnixListener{fd: fd, path: laddr.Name, unlink: false}, nil
//line /snap/go/10455/src/net/file_unix.go:98
		// _ = "end of CoverTab[5972]"
	}
//line /snap/go/10455/src/net/file_unix.go:99
	// _ = "end of CoverTab[5967]"
//line /snap/go/10455/src/net/file_unix.go:99
	_go_fuzz_dep_.CoverTab[5968]++
						fd.Close()
						return nil, syscall.EINVAL
//line /snap/go/10455/src/net/file_unix.go:101
	// _ = "end of CoverTab[5968]"
}

func filePacketConn(f *os.File) (PacketConn, error) {
//line /snap/go/10455/src/net/file_unix.go:104
	_go_fuzz_dep_.CoverTab[5973]++
						fd, err := newFileFD(f)
						if err != nil {
//line /snap/go/10455/src/net/file_unix.go:106
		_go_fuzz_dep_.CoverTab[528325]++
//line /snap/go/10455/src/net/file_unix.go:106
		_go_fuzz_dep_.CoverTab[5976]++
							return nil, err
//line /snap/go/10455/src/net/file_unix.go:107
		// _ = "end of CoverTab[5976]"
	} else {
//line /snap/go/10455/src/net/file_unix.go:108
		_go_fuzz_dep_.CoverTab[528326]++
//line /snap/go/10455/src/net/file_unix.go:108
		_go_fuzz_dep_.CoverTab[5977]++
//line /snap/go/10455/src/net/file_unix.go:108
		// _ = "end of CoverTab[5977]"
//line /snap/go/10455/src/net/file_unix.go:108
	}
//line /snap/go/10455/src/net/file_unix.go:108
	// _ = "end of CoverTab[5973]"
//line /snap/go/10455/src/net/file_unix.go:108
	_go_fuzz_dep_.CoverTab[5974]++
						switch fd.laddr.(type) {
	case *UDPAddr:
//line /snap/go/10455/src/net/file_unix.go:110
		_go_fuzz_dep_.CoverTab[528327]++
//line /snap/go/10455/src/net/file_unix.go:110
		_go_fuzz_dep_.CoverTab[5978]++
							return newUDPConn(fd), nil
//line /snap/go/10455/src/net/file_unix.go:111
		// _ = "end of CoverTab[5978]"
	case *IPAddr:
//line /snap/go/10455/src/net/file_unix.go:112
		_go_fuzz_dep_.CoverTab[528328]++
//line /snap/go/10455/src/net/file_unix.go:112
		_go_fuzz_dep_.CoverTab[5979]++
							return newIPConn(fd), nil
//line /snap/go/10455/src/net/file_unix.go:113
		// _ = "end of CoverTab[5979]"
	case *UnixAddr:
//line /snap/go/10455/src/net/file_unix.go:114
		_go_fuzz_dep_.CoverTab[528329]++
//line /snap/go/10455/src/net/file_unix.go:114
		_go_fuzz_dep_.CoverTab[5980]++
							return newUnixConn(fd), nil
//line /snap/go/10455/src/net/file_unix.go:115
		// _ = "end of CoverTab[5980]"
	}
//line /snap/go/10455/src/net/file_unix.go:116
	// _ = "end of CoverTab[5974]"
//line /snap/go/10455/src/net/file_unix.go:116
	_go_fuzz_dep_.CoverTab[5975]++
						fd.Close()
						return nil, syscall.EINVAL
//line /snap/go/10455/src/net/file_unix.go:118
	// _ = "end of CoverTab[5975]"
}

//line /snap/go/10455/src/net/file_unix.go:119
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/file_unix.go:119
var _ = _go_fuzz_dep_.CoverTab
