// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

//line /snap/go/10455/src/net/tcpsock_posix.go:7
package net

//line /snap/go/10455/src/net/tcpsock_posix.go:7
import (
//line /snap/go/10455/src/net/tcpsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/tcpsock_posix.go:7
)
//line /snap/go/10455/src/net/tcpsock_posix.go:7
import (
//line /snap/go/10455/src/net/tcpsock_posix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/tcpsock_posix.go:7
)

import (
	"context"
	"io"
	"os"
	"syscall"
)

func sockaddrToTCP(sa syscall.Sockaddr) Addr {
//line /snap/go/10455/src/net/tcpsock_posix.go:16
	_go_fuzz_dep_.CoverTab[8453]++
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /snap/go/10455/src/net/tcpsock_posix.go:18
		_go_fuzz_dep_.CoverTab[529894]++
//line /snap/go/10455/src/net/tcpsock_posix.go:18
		_go_fuzz_dep_.CoverTab[8455]++
								return &TCPAddr{IP: sa.Addr[0:], Port: sa.Port}
//line /snap/go/10455/src/net/tcpsock_posix.go:19
		// _ = "end of CoverTab[8455]"
	case *syscall.SockaddrInet6:
//line /snap/go/10455/src/net/tcpsock_posix.go:20
		_go_fuzz_dep_.CoverTab[529895]++
//line /snap/go/10455/src/net/tcpsock_posix.go:20
		_go_fuzz_dep_.CoverTab[8456]++
								return &TCPAddr{IP: sa.Addr[0:], Port: sa.Port, Zone: zoneCache.name(int(sa.ZoneId))}
//line /snap/go/10455/src/net/tcpsock_posix.go:21
		// _ = "end of CoverTab[8456]"
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:22
	// _ = "end of CoverTab[8453]"
//line /snap/go/10455/src/net/tcpsock_posix.go:22
	_go_fuzz_dep_.CoverTab[8454]++
							return nil
//line /snap/go/10455/src/net/tcpsock_posix.go:23
	// _ = "end of CoverTab[8454]"
}

func (a *TCPAddr) family() int {
//line /snap/go/10455/src/net/tcpsock_posix.go:26
	_go_fuzz_dep_.CoverTab[8457]++
							if a == nil || func() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:27
		_go_fuzz_dep_.CoverTab[8460]++
//line /snap/go/10455/src/net/tcpsock_posix.go:27
		return len(a.IP) <= IPv4len
//line /snap/go/10455/src/net/tcpsock_posix.go:27
		// _ = "end of CoverTab[8460]"
//line /snap/go/10455/src/net/tcpsock_posix.go:27
	}() {
//line /snap/go/10455/src/net/tcpsock_posix.go:27
		_go_fuzz_dep_.CoverTab[529896]++
//line /snap/go/10455/src/net/tcpsock_posix.go:27
		_go_fuzz_dep_.CoverTab[8461]++
								return syscall.AF_INET
//line /snap/go/10455/src/net/tcpsock_posix.go:28
		// _ = "end of CoverTab[8461]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:29
		_go_fuzz_dep_.CoverTab[529897]++
//line /snap/go/10455/src/net/tcpsock_posix.go:29
		_go_fuzz_dep_.CoverTab[8462]++
//line /snap/go/10455/src/net/tcpsock_posix.go:29
		// _ = "end of CoverTab[8462]"
//line /snap/go/10455/src/net/tcpsock_posix.go:29
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:29
	// _ = "end of CoverTab[8457]"
//line /snap/go/10455/src/net/tcpsock_posix.go:29
	_go_fuzz_dep_.CoverTab[8458]++
							if a.IP.To4() != nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:30
		_go_fuzz_dep_.CoverTab[529898]++
//line /snap/go/10455/src/net/tcpsock_posix.go:30
		_go_fuzz_dep_.CoverTab[8463]++
								return syscall.AF_INET
//line /snap/go/10455/src/net/tcpsock_posix.go:31
		// _ = "end of CoverTab[8463]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:32
		_go_fuzz_dep_.CoverTab[529899]++
//line /snap/go/10455/src/net/tcpsock_posix.go:32
		_go_fuzz_dep_.CoverTab[8464]++
//line /snap/go/10455/src/net/tcpsock_posix.go:32
		// _ = "end of CoverTab[8464]"
//line /snap/go/10455/src/net/tcpsock_posix.go:32
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:32
	// _ = "end of CoverTab[8458]"
//line /snap/go/10455/src/net/tcpsock_posix.go:32
	_go_fuzz_dep_.CoverTab[8459]++
							return syscall.AF_INET6
//line /snap/go/10455/src/net/tcpsock_posix.go:33
	// _ = "end of CoverTab[8459]"
}

func (a *TCPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /snap/go/10455/src/net/tcpsock_posix.go:36
	_go_fuzz_dep_.CoverTab[8465]++
							if a == nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:37
		_go_fuzz_dep_.CoverTab[529900]++
//line /snap/go/10455/src/net/tcpsock_posix.go:37
		_go_fuzz_dep_.CoverTab[8467]++
								return nil, nil
//line /snap/go/10455/src/net/tcpsock_posix.go:38
		// _ = "end of CoverTab[8467]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:39
		_go_fuzz_dep_.CoverTab[529901]++
//line /snap/go/10455/src/net/tcpsock_posix.go:39
		_go_fuzz_dep_.CoverTab[8468]++
//line /snap/go/10455/src/net/tcpsock_posix.go:39
		// _ = "end of CoverTab[8468]"
//line /snap/go/10455/src/net/tcpsock_posix.go:39
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:39
	// _ = "end of CoverTab[8465]"
//line /snap/go/10455/src/net/tcpsock_posix.go:39
	_go_fuzz_dep_.CoverTab[8466]++
							return ipToSockaddr(family, a.IP, a.Port, a.Zone)
//line /snap/go/10455/src/net/tcpsock_posix.go:40
	// _ = "end of CoverTab[8466]"
}

func (a *TCPAddr) toLocal(net string) sockaddr {
//line /snap/go/10455/src/net/tcpsock_posix.go:43
	_go_fuzz_dep_.CoverTab[8469]++
							return &TCPAddr{loopbackIP(net), a.Port, a.Zone}
//line /snap/go/10455/src/net/tcpsock_posix.go:44
	// _ = "end of CoverTab[8469]"
}

func (c *TCPConn) readFrom(r io.Reader) (int64, error) {
//line /snap/go/10455/src/net/tcpsock_posix.go:47
	_go_fuzz_dep_.CoverTab[8470]++
							if n, err, handled := splice(c.fd, r); handled {
//line /snap/go/10455/src/net/tcpsock_posix.go:48
		_go_fuzz_dep_.CoverTab[529902]++
//line /snap/go/10455/src/net/tcpsock_posix.go:48
		_go_fuzz_dep_.CoverTab[8473]++
								return n, err
//line /snap/go/10455/src/net/tcpsock_posix.go:49
		// _ = "end of CoverTab[8473]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:50
		_go_fuzz_dep_.CoverTab[529903]++
//line /snap/go/10455/src/net/tcpsock_posix.go:50
		_go_fuzz_dep_.CoverTab[8474]++
//line /snap/go/10455/src/net/tcpsock_posix.go:50
		// _ = "end of CoverTab[8474]"
//line /snap/go/10455/src/net/tcpsock_posix.go:50
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:50
	// _ = "end of CoverTab[8470]"
//line /snap/go/10455/src/net/tcpsock_posix.go:50
	_go_fuzz_dep_.CoverTab[8471]++
							if n, err, handled := sendFile(c.fd, r); handled {
//line /snap/go/10455/src/net/tcpsock_posix.go:51
		_go_fuzz_dep_.CoverTab[529904]++
//line /snap/go/10455/src/net/tcpsock_posix.go:51
		_go_fuzz_dep_.CoverTab[8475]++
								return n, err
//line /snap/go/10455/src/net/tcpsock_posix.go:52
		// _ = "end of CoverTab[8475]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:53
		_go_fuzz_dep_.CoverTab[529905]++
//line /snap/go/10455/src/net/tcpsock_posix.go:53
		_go_fuzz_dep_.CoverTab[8476]++
//line /snap/go/10455/src/net/tcpsock_posix.go:53
		// _ = "end of CoverTab[8476]"
//line /snap/go/10455/src/net/tcpsock_posix.go:53
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:53
	// _ = "end of CoverTab[8471]"
//line /snap/go/10455/src/net/tcpsock_posix.go:53
	_go_fuzz_dep_.CoverTab[8472]++
							return genericReadFrom(c, r)
//line /snap/go/10455/src/net/tcpsock_posix.go:54
	// _ = "end of CoverTab[8472]"
}

func (sd *sysDialer) dialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
//line /snap/go/10455/src/net/tcpsock_posix.go:57
	_go_fuzz_dep_.CoverTab[8477]++
							if h := sd.testHookDialTCP; h != nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:58
		_go_fuzz_dep_.CoverTab[529906]++
//line /snap/go/10455/src/net/tcpsock_posix.go:58
		_go_fuzz_dep_.CoverTab[8480]++
								return h(ctx, sd.network, laddr, raddr)
//line /snap/go/10455/src/net/tcpsock_posix.go:59
		// _ = "end of CoverTab[8480]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:60
		_go_fuzz_dep_.CoverTab[529907]++
//line /snap/go/10455/src/net/tcpsock_posix.go:60
		_go_fuzz_dep_.CoverTab[8481]++
//line /snap/go/10455/src/net/tcpsock_posix.go:60
		// _ = "end of CoverTab[8481]"
//line /snap/go/10455/src/net/tcpsock_posix.go:60
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:60
	// _ = "end of CoverTab[8477]"
//line /snap/go/10455/src/net/tcpsock_posix.go:60
	_go_fuzz_dep_.CoverTab[8478]++
							if h := testHookDialTCP; h != nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:61
		_go_fuzz_dep_.CoverTab[529908]++
//line /snap/go/10455/src/net/tcpsock_posix.go:61
		_go_fuzz_dep_.CoverTab[8482]++
								return h(ctx, sd.network, laddr, raddr)
//line /snap/go/10455/src/net/tcpsock_posix.go:62
		// _ = "end of CoverTab[8482]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:63
		_go_fuzz_dep_.CoverTab[529909]++
//line /snap/go/10455/src/net/tcpsock_posix.go:63
		_go_fuzz_dep_.CoverTab[8483]++
//line /snap/go/10455/src/net/tcpsock_posix.go:63
		// _ = "end of CoverTab[8483]"
//line /snap/go/10455/src/net/tcpsock_posix.go:63
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:63
	// _ = "end of CoverTab[8478]"
//line /snap/go/10455/src/net/tcpsock_posix.go:63
	_go_fuzz_dep_.CoverTab[8479]++
							return sd.doDialTCP(ctx, laddr, raddr)
//line /snap/go/10455/src/net/tcpsock_posix.go:64
	// _ = "end of CoverTab[8479]"
}

func (sd *sysDialer) doDialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
//line /snap/go/10455/src/net/tcpsock_posix.go:67
	_go_fuzz_dep_.CoverTab[8484]++
							return sd.doDialTCPProto(ctx, laddr, raddr, 0)
//line /snap/go/10455/src/net/tcpsock_posix.go:68
	// _ = "end of CoverTab[8484]"
}

func (sd *sysDialer) doDialTCPProto(ctx context.Context, laddr, raddr *TCPAddr, proto int) (*TCPConn, error) {
//line /snap/go/10455/src/net/tcpsock_posix.go:71
	_go_fuzz_dep_.CoverTab[8485]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:73
		_go_fuzz_dep_.CoverTab[8489]++
//line /snap/go/10455/src/net/tcpsock_posix.go:73
		return sd.Dialer.Control != nil
//line /snap/go/10455/src/net/tcpsock_posix.go:73
		// _ = "end of CoverTab[8489]"
//line /snap/go/10455/src/net/tcpsock_posix.go:73
	}() {
//line /snap/go/10455/src/net/tcpsock_posix.go:73
		_go_fuzz_dep_.CoverTab[529910]++
//line /snap/go/10455/src/net/tcpsock_posix.go:73
		_go_fuzz_dep_.CoverTab[8490]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/tcpsock_posix.go:74
			_go_fuzz_dep_.CoverTab[8491]++
									return sd.Dialer.Control(network, address, c)
//line /snap/go/10455/src/net/tcpsock_posix.go:75
			// _ = "end of CoverTab[8491]"
		}
//line /snap/go/10455/src/net/tcpsock_posix.go:76
		// _ = "end of CoverTab[8490]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:77
		_go_fuzz_dep_.CoverTab[529911]++
//line /snap/go/10455/src/net/tcpsock_posix.go:77
		_go_fuzz_dep_.CoverTab[8492]++
//line /snap/go/10455/src/net/tcpsock_posix.go:77
		// _ = "end of CoverTab[8492]"
//line /snap/go/10455/src/net/tcpsock_posix.go:77
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:77
	// _ = "end of CoverTab[8485]"
//line /snap/go/10455/src/net/tcpsock_posix.go:77
	_go_fuzz_dep_.CoverTab[8486]++
							fd, err := internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_STREAM, proto, "dial", ctrlCtxFn)
//line /snap/go/10455/src/net/tcpsock_posix.go:78
	_go_fuzz_dep_.CoverTab[786748] = 0

//line /snap/go/10455/src/net/tcpsock_posix.go:104
	for i := 0; i < 2 && func() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		_go_fuzz_dep_.CoverTab[8493]++
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		return (laddr == nil || func() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:104
			_go_fuzz_dep_.CoverTab[8494]++
//line /snap/go/10455/src/net/tcpsock_posix.go:104
			return laddr.Port == 0
//line /snap/go/10455/src/net/tcpsock_posix.go:104
			// _ = "end of CoverTab[8494]"
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		}())
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		// _ = "end of CoverTab[8493]"
//line /snap/go/10455/src/net/tcpsock_posix.go:104
	}() && func() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		_go_fuzz_dep_.CoverTab[8495]++
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		return (selfConnect(fd, err) || func() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:104
			_go_fuzz_dep_.CoverTab[8496]++
//line /snap/go/10455/src/net/tcpsock_posix.go:104
			return spuriousENOTAVAIL(err)
//line /snap/go/10455/src/net/tcpsock_posix.go:104
			// _ = "end of CoverTab[8496]"
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		}())
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		// _ = "end of CoverTab[8495]"
//line /snap/go/10455/src/net/tcpsock_posix.go:104
	}(); i++ {
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		if _go_fuzz_dep_.CoverTab[786748] == 0 {
//line /snap/go/10455/src/net/tcpsock_posix.go:104
			_go_fuzz_dep_.CoverTab[529932]++
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:104
			_go_fuzz_dep_.CoverTab[529933]++
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		}
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		_go_fuzz_dep_.CoverTab[786748] = 1
//line /snap/go/10455/src/net/tcpsock_posix.go:104
		_go_fuzz_dep_.CoverTab[8497]++
								if err == nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:105
			_go_fuzz_dep_.CoverTab[529912]++
//line /snap/go/10455/src/net/tcpsock_posix.go:105
			_go_fuzz_dep_.CoverTab[8499]++
									fd.Close()
//line /snap/go/10455/src/net/tcpsock_posix.go:106
			// _ = "end of CoverTab[8499]"
		} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:107
			_go_fuzz_dep_.CoverTab[529913]++
//line /snap/go/10455/src/net/tcpsock_posix.go:107
			_go_fuzz_dep_.CoverTab[8500]++
//line /snap/go/10455/src/net/tcpsock_posix.go:107
			// _ = "end of CoverTab[8500]"
//line /snap/go/10455/src/net/tcpsock_posix.go:107
		}
//line /snap/go/10455/src/net/tcpsock_posix.go:107
		// _ = "end of CoverTab[8497]"
//line /snap/go/10455/src/net/tcpsock_posix.go:107
		_go_fuzz_dep_.CoverTab[8498]++
								fd, err = internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_STREAM, proto, "dial", ctrlCtxFn)
//line /snap/go/10455/src/net/tcpsock_posix.go:108
		// _ = "end of CoverTab[8498]"
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:109
	if _go_fuzz_dep_.CoverTab[786748] == 0 {
//line /snap/go/10455/src/net/tcpsock_posix.go:109
		_go_fuzz_dep_.CoverTab[529934]++
//line /snap/go/10455/src/net/tcpsock_posix.go:109
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:109
		_go_fuzz_dep_.CoverTab[529935]++
//line /snap/go/10455/src/net/tcpsock_posix.go:109
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:109
	// _ = "end of CoverTab[8486]"
//line /snap/go/10455/src/net/tcpsock_posix.go:109
	_go_fuzz_dep_.CoverTab[8487]++

							if err != nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:111
		_go_fuzz_dep_.CoverTab[529914]++
//line /snap/go/10455/src/net/tcpsock_posix.go:111
		_go_fuzz_dep_.CoverTab[8501]++
								return nil, err
//line /snap/go/10455/src/net/tcpsock_posix.go:112
		// _ = "end of CoverTab[8501]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:113
		_go_fuzz_dep_.CoverTab[529915]++
//line /snap/go/10455/src/net/tcpsock_posix.go:113
		_go_fuzz_dep_.CoverTab[8502]++
//line /snap/go/10455/src/net/tcpsock_posix.go:113
		// _ = "end of CoverTab[8502]"
//line /snap/go/10455/src/net/tcpsock_posix.go:113
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:113
	// _ = "end of CoverTab[8487]"
//line /snap/go/10455/src/net/tcpsock_posix.go:113
	_go_fuzz_dep_.CoverTab[8488]++
							return newTCPConn(fd, sd.Dialer.KeepAlive, testHookSetKeepAlive), nil
//line /snap/go/10455/src/net/tcpsock_posix.go:114
	// _ = "end of CoverTab[8488]"
}

func selfConnect(fd *netFD, err error) bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:117
	_go_fuzz_dep_.CoverTab[8503]++

							if err != nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:119
		_go_fuzz_dep_.CoverTab[529916]++
//line /snap/go/10455/src/net/tcpsock_posix.go:119
		_go_fuzz_dep_.CoverTab[8506]++
								return false
//line /snap/go/10455/src/net/tcpsock_posix.go:120
		// _ = "end of CoverTab[8506]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:121
		_go_fuzz_dep_.CoverTab[529917]++
//line /snap/go/10455/src/net/tcpsock_posix.go:121
		_go_fuzz_dep_.CoverTab[8507]++
//line /snap/go/10455/src/net/tcpsock_posix.go:121
		// _ = "end of CoverTab[8507]"
//line /snap/go/10455/src/net/tcpsock_posix.go:121
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:121
	// _ = "end of CoverTab[8503]"
//line /snap/go/10455/src/net/tcpsock_posix.go:121
	_go_fuzz_dep_.CoverTab[8504]++

//line /snap/go/10455/src/net/tcpsock_posix.go:131
	if fd.laddr == nil || func() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:131
		_go_fuzz_dep_.CoverTab[8508]++
//line /snap/go/10455/src/net/tcpsock_posix.go:131
		return fd.raddr == nil
//line /snap/go/10455/src/net/tcpsock_posix.go:131
		// _ = "end of CoverTab[8508]"
//line /snap/go/10455/src/net/tcpsock_posix.go:131
	}() {
//line /snap/go/10455/src/net/tcpsock_posix.go:131
		_go_fuzz_dep_.CoverTab[529918]++
//line /snap/go/10455/src/net/tcpsock_posix.go:131
		_go_fuzz_dep_.CoverTab[8509]++
								return true
//line /snap/go/10455/src/net/tcpsock_posix.go:132
		// _ = "end of CoverTab[8509]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:133
		_go_fuzz_dep_.CoverTab[529919]++
//line /snap/go/10455/src/net/tcpsock_posix.go:133
		_go_fuzz_dep_.CoverTab[8510]++
//line /snap/go/10455/src/net/tcpsock_posix.go:133
		// _ = "end of CoverTab[8510]"
//line /snap/go/10455/src/net/tcpsock_posix.go:133
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:133
	// _ = "end of CoverTab[8504]"
//line /snap/go/10455/src/net/tcpsock_posix.go:133
	_go_fuzz_dep_.CoverTab[8505]++
							l := fd.laddr.(*TCPAddr)
							r := fd.raddr.(*TCPAddr)
							return l.Port == r.Port && func() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:136
		_go_fuzz_dep_.CoverTab[8511]++
//line /snap/go/10455/src/net/tcpsock_posix.go:136
		return l.IP.Equal(r.IP)
//line /snap/go/10455/src/net/tcpsock_posix.go:136
		// _ = "end of CoverTab[8511]"
//line /snap/go/10455/src/net/tcpsock_posix.go:136
	}()
//line /snap/go/10455/src/net/tcpsock_posix.go:136
	// _ = "end of CoverTab[8505]"
}

func spuriousENOTAVAIL(err error) bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:139
	_go_fuzz_dep_.CoverTab[8512]++
							if op, ok := err.(*OpError); ok {
//line /snap/go/10455/src/net/tcpsock_posix.go:140
		_go_fuzz_dep_.CoverTab[529920]++
//line /snap/go/10455/src/net/tcpsock_posix.go:140
		_go_fuzz_dep_.CoverTab[8515]++
								err = op.Err
//line /snap/go/10455/src/net/tcpsock_posix.go:141
		// _ = "end of CoverTab[8515]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:142
		_go_fuzz_dep_.CoverTab[529921]++
//line /snap/go/10455/src/net/tcpsock_posix.go:142
		_go_fuzz_dep_.CoverTab[8516]++
//line /snap/go/10455/src/net/tcpsock_posix.go:142
		// _ = "end of CoverTab[8516]"
//line /snap/go/10455/src/net/tcpsock_posix.go:142
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:142
	// _ = "end of CoverTab[8512]"
//line /snap/go/10455/src/net/tcpsock_posix.go:142
	_go_fuzz_dep_.CoverTab[8513]++
							if sys, ok := err.(*os.SyscallError); ok {
//line /snap/go/10455/src/net/tcpsock_posix.go:143
		_go_fuzz_dep_.CoverTab[529922]++
//line /snap/go/10455/src/net/tcpsock_posix.go:143
		_go_fuzz_dep_.CoverTab[8517]++
								err = sys.Err
//line /snap/go/10455/src/net/tcpsock_posix.go:144
		// _ = "end of CoverTab[8517]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:145
		_go_fuzz_dep_.CoverTab[529923]++
//line /snap/go/10455/src/net/tcpsock_posix.go:145
		_go_fuzz_dep_.CoverTab[8518]++
//line /snap/go/10455/src/net/tcpsock_posix.go:145
		// _ = "end of CoverTab[8518]"
//line /snap/go/10455/src/net/tcpsock_posix.go:145
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:145
	// _ = "end of CoverTab[8513]"
//line /snap/go/10455/src/net/tcpsock_posix.go:145
	_go_fuzz_dep_.CoverTab[8514]++
							return err == syscall.EADDRNOTAVAIL
//line /snap/go/10455/src/net/tcpsock_posix.go:146
	// _ = "end of CoverTab[8514]"
}

func (ln *TCPListener) ok() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:149
	_go_fuzz_dep_.CoverTab[8519]++
//line /snap/go/10455/src/net/tcpsock_posix.go:149
	return ln != nil && func() bool {
//line /snap/go/10455/src/net/tcpsock_posix.go:149
		_go_fuzz_dep_.CoverTab[8520]++
//line /snap/go/10455/src/net/tcpsock_posix.go:149
		return ln.fd != nil
//line /snap/go/10455/src/net/tcpsock_posix.go:149
		// _ = "end of CoverTab[8520]"
//line /snap/go/10455/src/net/tcpsock_posix.go:149
	}()
//line /snap/go/10455/src/net/tcpsock_posix.go:149
	// _ = "end of CoverTab[8519]"
//line /snap/go/10455/src/net/tcpsock_posix.go:149
}

func (ln *TCPListener) accept() (*TCPConn, error) {
//line /snap/go/10455/src/net/tcpsock_posix.go:151
	_go_fuzz_dep_.CoverTab[8521]++
							fd, err := ln.fd.accept()
							if err != nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:153
		_go_fuzz_dep_.CoverTab[529924]++
//line /snap/go/10455/src/net/tcpsock_posix.go:153
		_go_fuzz_dep_.CoverTab[8523]++
								return nil, err
//line /snap/go/10455/src/net/tcpsock_posix.go:154
		// _ = "end of CoverTab[8523]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:155
		_go_fuzz_dep_.CoverTab[529925]++
//line /snap/go/10455/src/net/tcpsock_posix.go:155
		_go_fuzz_dep_.CoverTab[8524]++
//line /snap/go/10455/src/net/tcpsock_posix.go:155
		// _ = "end of CoverTab[8524]"
//line /snap/go/10455/src/net/tcpsock_posix.go:155
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:155
	// _ = "end of CoverTab[8521]"
//line /snap/go/10455/src/net/tcpsock_posix.go:155
	_go_fuzz_dep_.CoverTab[8522]++
							return newTCPConn(fd, ln.lc.KeepAlive, nil), nil
//line /snap/go/10455/src/net/tcpsock_posix.go:156
	// _ = "end of CoverTab[8522]"
}

func (ln *TCPListener) close() error {
//line /snap/go/10455/src/net/tcpsock_posix.go:159
	_go_fuzz_dep_.CoverTab[8525]++
							return ln.fd.Close()
//line /snap/go/10455/src/net/tcpsock_posix.go:160
	// _ = "end of CoverTab[8525]"
}

func (ln *TCPListener) file() (*os.File, error) {
//line /snap/go/10455/src/net/tcpsock_posix.go:163
	_go_fuzz_dep_.CoverTab[8526]++
							f, err := ln.fd.dup()
							if err != nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:165
		_go_fuzz_dep_.CoverTab[529926]++
//line /snap/go/10455/src/net/tcpsock_posix.go:165
		_go_fuzz_dep_.CoverTab[8528]++
								return nil, err
//line /snap/go/10455/src/net/tcpsock_posix.go:166
		// _ = "end of CoverTab[8528]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:167
		_go_fuzz_dep_.CoverTab[529927]++
//line /snap/go/10455/src/net/tcpsock_posix.go:167
		_go_fuzz_dep_.CoverTab[8529]++
//line /snap/go/10455/src/net/tcpsock_posix.go:167
		// _ = "end of CoverTab[8529]"
//line /snap/go/10455/src/net/tcpsock_posix.go:167
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:167
	// _ = "end of CoverTab[8526]"
//line /snap/go/10455/src/net/tcpsock_posix.go:167
	_go_fuzz_dep_.CoverTab[8527]++
							return f, nil
//line /snap/go/10455/src/net/tcpsock_posix.go:168
	// _ = "end of CoverTab[8527]"
}

func (sl *sysListener) listenTCP(ctx context.Context, laddr *TCPAddr) (*TCPListener, error) {
//line /snap/go/10455/src/net/tcpsock_posix.go:171
	_go_fuzz_dep_.CoverTab[8530]++
							return sl.listenTCPProto(ctx, laddr, 0)
//line /snap/go/10455/src/net/tcpsock_posix.go:172
	// _ = "end of CoverTab[8530]"
}

func (sl *sysListener) listenTCPProto(ctx context.Context, laddr *TCPAddr, proto int) (*TCPListener, error) {
//line /snap/go/10455/src/net/tcpsock_posix.go:175
	_go_fuzz_dep_.CoverTab[8531]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:177
		_go_fuzz_dep_.CoverTab[529928]++
//line /snap/go/10455/src/net/tcpsock_posix.go:177
		_go_fuzz_dep_.CoverTab[8534]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/tcpsock_posix.go:178
			_go_fuzz_dep_.CoverTab[8535]++
									return sl.ListenConfig.Control(network, address, c)
//line /snap/go/10455/src/net/tcpsock_posix.go:179
			// _ = "end of CoverTab[8535]"
		}
//line /snap/go/10455/src/net/tcpsock_posix.go:180
		// _ = "end of CoverTab[8534]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:181
		_go_fuzz_dep_.CoverTab[529929]++
//line /snap/go/10455/src/net/tcpsock_posix.go:181
		_go_fuzz_dep_.CoverTab[8536]++
//line /snap/go/10455/src/net/tcpsock_posix.go:181
		// _ = "end of CoverTab[8536]"
//line /snap/go/10455/src/net/tcpsock_posix.go:181
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:181
	// _ = "end of CoverTab[8531]"
//line /snap/go/10455/src/net/tcpsock_posix.go:181
	_go_fuzz_dep_.CoverTab[8532]++
							fd, err := internetSocket(ctx, sl.network, laddr, nil, syscall.SOCK_STREAM, proto, "listen", ctrlCtxFn)
							if err != nil {
//line /snap/go/10455/src/net/tcpsock_posix.go:183
		_go_fuzz_dep_.CoverTab[529930]++
//line /snap/go/10455/src/net/tcpsock_posix.go:183
		_go_fuzz_dep_.CoverTab[8537]++
								return nil, err
//line /snap/go/10455/src/net/tcpsock_posix.go:184
		// _ = "end of CoverTab[8537]"
	} else {
//line /snap/go/10455/src/net/tcpsock_posix.go:185
		_go_fuzz_dep_.CoverTab[529931]++
//line /snap/go/10455/src/net/tcpsock_posix.go:185
		_go_fuzz_dep_.CoverTab[8538]++
//line /snap/go/10455/src/net/tcpsock_posix.go:185
		// _ = "end of CoverTab[8538]"
//line /snap/go/10455/src/net/tcpsock_posix.go:185
	}
//line /snap/go/10455/src/net/tcpsock_posix.go:185
	// _ = "end of CoverTab[8532]"
//line /snap/go/10455/src/net/tcpsock_posix.go:185
	_go_fuzz_dep_.CoverTab[8533]++
							return &TCPListener{fd: fd, lc: sl.ListenConfig}, nil
//line /snap/go/10455/src/net/tcpsock_posix.go:186
	// _ = "end of CoverTab[8533]"
}

//line /snap/go/10455/src/net/tcpsock_posix.go:187
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/tcpsock_posix.go:187
var _ = _go_fuzz_dep_.CoverTab
