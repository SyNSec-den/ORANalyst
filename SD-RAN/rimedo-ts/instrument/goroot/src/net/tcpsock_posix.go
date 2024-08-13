// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || windows

//line /usr/local/go/src/net/tcpsock_posix.go:7
package net

//line /usr/local/go/src/net/tcpsock_posix.go:7
import (
//line /usr/local/go/src/net/tcpsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/tcpsock_posix.go:7
)
//line /usr/local/go/src/net/tcpsock_posix.go:7
import (
//line /usr/local/go/src/net/tcpsock_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/tcpsock_posix.go:7
)

import (
	"context"
	"io"
	"os"
	"syscall"
)

func sockaddrToTCP(sa syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/tcpsock_posix.go:16
	_go_fuzz_dep_.CoverTab[16545]++
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/tcpsock_posix.go:18
		_go_fuzz_dep_.CoverTab[16547]++
								return &TCPAddr{IP: sa.Addr[0:], Port: sa.Port}
//line /usr/local/go/src/net/tcpsock_posix.go:19
		// _ = "end of CoverTab[16547]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/tcpsock_posix.go:20
		_go_fuzz_dep_.CoverTab[16548]++
								return &TCPAddr{IP: sa.Addr[0:], Port: sa.Port, Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/tcpsock_posix.go:21
		// _ = "end of CoverTab[16548]"
	}
//line /usr/local/go/src/net/tcpsock_posix.go:22
	// _ = "end of CoverTab[16545]"
//line /usr/local/go/src/net/tcpsock_posix.go:22
	_go_fuzz_dep_.CoverTab[16546]++
							return nil
//line /usr/local/go/src/net/tcpsock_posix.go:23
	// _ = "end of CoverTab[16546]"
}

func (a *TCPAddr) family() int {
//line /usr/local/go/src/net/tcpsock_posix.go:26
	_go_fuzz_dep_.CoverTab[16549]++
							if a == nil || func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:27
		_go_fuzz_dep_.CoverTab[16552]++
//line /usr/local/go/src/net/tcpsock_posix.go:27
		return len(a.IP) <= IPv4len
//line /usr/local/go/src/net/tcpsock_posix.go:27
		// _ = "end of CoverTab[16552]"
//line /usr/local/go/src/net/tcpsock_posix.go:27
	}() {
//line /usr/local/go/src/net/tcpsock_posix.go:27
		_go_fuzz_dep_.CoverTab[16553]++
								return syscall.AF_INET
//line /usr/local/go/src/net/tcpsock_posix.go:28
		// _ = "end of CoverTab[16553]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:29
		_go_fuzz_dep_.CoverTab[16554]++
//line /usr/local/go/src/net/tcpsock_posix.go:29
		// _ = "end of CoverTab[16554]"
//line /usr/local/go/src/net/tcpsock_posix.go:29
	}
//line /usr/local/go/src/net/tcpsock_posix.go:29
	// _ = "end of CoverTab[16549]"
//line /usr/local/go/src/net/tcpsock_posix.go:29
	_go_fuzz_dep_.CoverTab[16550]++
							if a.IP.To4() != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:30
		_go_fuzz_dep_.CoverTab[16555]++
								return syscall.AF_INET
//line /usr/local/go/src/net/tcpsock_posix.go:31
		// _ = "end of CoverTab[16555]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:32
		_go_fuzz_dep_.CoverTab[16556]++
//line /usr/local/go/src/net/tcpsock_posix.go:32
		// _ = "end of CoverTab[16556]"
//line /usr/local/go/src/net/tcpsock_posix.go:32
	}
//line /usr/local/go/src/net/tcpsock_posix.go:32
	// _ = "end of CoverTab[16550]"
//line /usr/local/go/src/net/tcpsock_posix.go:32
	_go_fuzz_dep_.CoverTab[16551]++
							return syscall.AF_INET6
//line /usr/local/go/src/net/tcpsock_posix.go:33
	// _ = "end of CoverTab[16551]"
}

func (a *TCPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:36
	_go_fuzz_dep_.CoverTab[16557]++
							if a == nil {
//line /usr/local/go/src/net/tcpsock_posix.go:37
		_go_fuzz_dep_.CoverTab[16559]++
								return nil, nil
//line /usr/local/go/src/net/tcpsock_posix.go:38
		// _ = "end of CoverTab[16559]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:39
		_go_fuzz_dep_.CoverTab[16560]++
//line /usr/local/go/src/net/tcpsock_posix.go:39
		// _ = "end of CoverTab[16560]"
//line /usr/local/go/src/net/tcpsock_posix.go:39
	}
//line /usr/local/go/src/net/tcpsock_posix.go:39
	// _ = "end of CoverTab[16557]"
//line /usr/local/go/src/net/tcpsock_posix.go:39
	_go_fuzz_dep_.CoverTab[16558]++
							return ipToSockaddr(family, a.IP, a.Port, a.Zone)
//line /usr/local/go/src/net/tcpsock_posix.go:40
	// _ = "end of CoverTab[16558]"
}

func (a *TCPAddr) toLocal(net string) sockaddr {
//line /usr/local/go/src/net/tcpsock_posix.go:43
	_go_fuzz_dep_.CoverTab[16561]++
							return &TCPAddr{loopbackIP(net), a.Port, a.Zone}
//line /usr/local/go/src/net/tcpsock_posix.go:44
	// _ = "end of CoverTab[16561]"
}

func (c *TCPConn) readFrom(r io.Reader) (int64, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:47
	_go_fuzz_dep_.CoverTab[16562]++
							if n, err, handled := splice(c.fd, r); handled {
//line /usr/local/go/src/net/tcpsock_posix.go:48
		_go_fuzz_dep_.CoverTab[16565]++
								return n, err
//line /usr/local/go/src/net/tcpsock_posix.go:49
		// _ = "end of CoverTab[16565]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:50
		_go_fuzz_dep_.CoverTab[16566]++
//line /usr/local/go/src/net/tcpsock_posix.go:50
		// _ = "end of CoverTab[16566]"
//line /usr/local/go/src/net/tcpsock_posix.go:50
	}
//line /usr/local/go/src/net/tcpsock_posix.go:50
	// _ = "end of CoverTab[16562]"
//line /usr/local/go/src/net/tcpsock_posix.go:50
	_go_fuzz_dep_.CoverTab[16563]++
							if n, err, handled := sendFile(c.fd, r); handled {
//line /usr/local/go/src/net/tcpsock_posix.go:51
		_go_fuzz_dep_.CoverTab[16567]++
								return n, err
//line /usr/local/go/src/net/tcpsock_posix.go:52
		// _ = "end of CoverTab[16567]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:53
		_go_fuzz_dep_.CoverTab[16568]++
//line /usr/local/go/src/net/tcpsock_posix.go:53
		// _ = "end of CoverTab[16568]"
//line /usr/local/go/src/net/tcpsock_posix.go:53
	}
//line /usr/local/go/src/net/tcpsock_posix.go:53
	// _ = "end of CoverTab[16563]"
//line /usr/local/go/src/net/tcpsock_posix.go:53
	_go_fuzz_dep_.CoverTab[16564]++
							return genericReadFrom(c, r)
//line /usr/local/go/src/net/tcpsock_posix.go:54
	// _ = "end of CoverTab[16564]"
}

func (sd *sysDialer) dialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:57
	_go_fuzz_dep_.CoverTab[16569]++
							if h := sd.testHookDialTCP; h != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:58
		_go_fuzz_dep_.CoverTab[16572]++
								return h(ctx, sd.network, laddr, raddr)
//line /usr/local/go/src/net/tcpsock_posix.go:59
		// _ = "end of CoverTab[16572]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:60
		_go_fuzz_dep_.CoverTab[16573]++
//line /usr/local/go/src/net/tcpsock_posix.go:60
		// _ = "end of CoverTab[16573]"
//line /usr/local/go/src/net/tcpsock_posix.go:60
	}
//line /usr/local/go/src/net/tcpsock_posix.go:60
	// _ = "end of CoverTab[16569]"
//line /usr/local/go/src/net/tcpsock_posix.go:60
	_go_fuzz_dep_.CoverTab[16570]++
							if h := testHookDialTCP; h != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:61
		_go_fuzz_dep_.CoverTab[16574]++
								return h(ctx, sd.network, laddr, raddr)
//line /usr/local/go/src/net/tcpsock_posix.go:62
		// _ = "end of CoverTab[16574]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:63
		_go_fuzz_dep_.CoverTab[16575]++
//line /usr/local/go/src/net/tcpsock_posix.go:63
		// _ = "end of CoverTab[16575]"
//line /usr/local/go/src/net/tcpsock_posix.go:63
	}
//line /usr/local/go/src/net/tcpsock_posix.go:63
	// _ = "end of CoverTab[16570]"
//line /usr/local/go/src/net/tcpsock_posix.go:63
	_go_fuzz_dep_.CoverTab[16571]++
							return sd.doDialTCP(ctx, laddr, raddr)
//line /usr/local/go/src/net/tcpsock_posix.go:64
	// _ = "end of CoverTab[16571]"
}

func (sd *sysDialer) doDialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:67
	_go_fuzz_dep_.CoverTab[16576]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:69
		_go_fuzz_dep_.CoverTab[16580]++
//line /usr/local/go/src/net/tcpsock_posix.go:69
		return sd.Dialer.Control != nil
//line /usr/local/go/src/net/tcpsock_posix.go:69
		// _ = "end of CoverTab[16580]"
//line /usr/local/go/src/net/tcpsock_posix.go:69
	}() {
//line /usr/local/go/src/net/tcpsock_posix.go:69
		_go_fuzz_dep_.CoverTab[16581]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/tcpsock_posix.go:70
			_go_fuzz_dep_.CoverTab[16582]++
									return sd.Dialer.Control(network, address, c)
//line /usr/local/go/src/net/tcpsock_posix.go:71
			// _ = "end of CoverTab[16582]"
		}
//line /usr/local/go/src/net/tcpsock_posix.go:72
		// _ = "end of CoverTab[16581]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:73
		_go_fuzz_dep_.CoverTab[16583]++
//line /usr/local/go/src/net/tcpsock_posix.go:73
		// _ = "end of CoverTab[16583]"
//line /usr/local/go/src/net/tcpsock_posix.go:73
	}
//line /usr/local/go/src/net/tcpsock_posix.go:73
	// _ = "end of CoverTab[16576]"
//line /usr/local/go/src/net/tcpsock_posix.go:73
	_go_fuzz_dep_.CoverTab[16577]++
							fd, err := internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_STREAM, 0, "dial", ctrlCtxFn)

//line /usr/local/go/src/net/tcpsock_posix.go:100
	for i := 0; i < 2 && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:100
		_go_fuzz_dep_.CoverTab[16584]++
//line /usr/local/go/src/net/tcpsock_posix.go:100
		return (laddr == nil || func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:100
			_go_fuzz_dep_.CoverTab[16585]++
//line /usr/local/go/src/net/tcpsock_posix.go:100
			return laddr.Port == 0
//line /usr/local/go/src/net/tcpsock_posix.go:100
			// _ = "end of CoverTab[16585]"
//line /usr/local/go/src/net/tcpsock_posix.go:100
		}())
//line /usr/local/go/src/net/tcpsock_posix.go:100
		// _ = "end of CoverTab[16584]"
//line /usr/local/go/src/net/tcpsock_posix.go:100
	}() && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:100
		_go_fuzz_dep_.CoverTab[16586]++
//line /usr/local/go/src/net/tcpsock_posix.go:100
		return (selfConnect(fd, err) || func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:100
			_go_fuzz_dep_.CoverTab[16587]++
//line /usr/local/go/src/net/tcpsock_posix.go:100
			return spuriousENOTAVAIL(err)
//line /usr/local/go/src/net/tcpsock_posix.go:100
			// _ = "end of CoverTab[16587]"
//line /usr/local/go/src/net/tcpsock_posix.go:100
		}())
//line /usr/local/go/src/net/tcpsock_posix.go:100
		// _ = "end of CoverTab[16586]"
//line /usr/local/go/src/net/tcpsock_posix.go:100
	}(); i++ {
//line /usr/local/go/src/net/tcpsock_posix.go:100
		_go_fuzz_dep_.CoverTab[16588]++
								if err == nil {
//line /usr/local/go/src/net/tcpsock_posix.go:101
			_go_fuzz_dep_.CoverTab[16590]++
									fd.Close()
//line /usr/local/go/src/net/tcpsock_posix.go:102
			// _ = "end of CoverTab[16590]"
		} else {
//line /usr/local/go/src/net/tcpsock_posix.go:103
			_go_fuzz_dep_.CoverTab[16591]++
//line /usr/local/go/src/net/tcpsock_posix.go:103
			// _ = "end of CoverTab[16591]"
//line /usr/local/go/src/net/tcpsock_posix.go:103
		}
//line /usr/local/go/src/net/tcpsock_posix.go:103
		// _ = "end of CoverTab[16588]"
//line /usr/local/go/src/net/tcpsock_posix.go:103
		_go_fuzz_dep_.CoverTab[16589]++
								fd, err = internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_STREAM, 0, "dial", ctrlCtxFn)
//line /usr/local/go/src/net/tcpsock_posix.go:104
		// _ = "end of CoverTab[16589]"
	}
//line /usr/local/go/src/net/tcpsock_posix.go:105
	// _ = "end of CoverTab[16577]"
//line /usr/local/go/src/net/tcpsock_posix.go:105
	_go_fuzz_dep_.CoverTab[16578]++

							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:107
		_go_fuzz_dep_.CoverTab[16592]++
								return nil, err
//line /usr/local/go/src/net/tcpsock_posix.go:108
		// _ = "end of CoverTab[16592]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:109
		_go_fuzz_dep_.CoverTab[16593]++
//line /usr/local/go/src/net/tcpsock_posix.go:109
		// _ = "end of CoverTab[16593]"
//line /usr/local/go/src/net/tcpsock_posix.go:109
	}
//line /usr/local/go/src/net/tcpsock_posix.go:109
	// _ = "end of CoverTab[16578]"
//line /usr/local/go/src/net/tcpsock_posix.go:109
	_go_fuzz_dep_.CoverTab[16579]++
							return newTCPConn(fd, sd.Dialer.KeepAlive, testHookSetKeepAlive), nil
//line /usr/local/go/src/net/tcpsock_posix.go:110
	// _ = "end of CoverTab[16579]"
}

func selfConnect(fd *netFD, err error) bool {
//line /usr/local/go/src/net/tcpsock_posix.go:113
	_go_fuzz_dep_.CoverTab[16594]++

							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:115
		_go_fuzz_dep_.CoverTab[16597]++
								return false
//line /usr/local/go/src/net/tcpsock_posix.go:116
		// _ = "end of CoverTab[16597]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:117
		_go_fuzz_dep_.CoverTab[16598]++
//line /usr/local/go/src/net/tcpsock_posix.go:117
		// _ = "end of CoverTab[16598]"
//line /usr/local/go/src/net/tcpsock_posix.go:117
	}
//line /usr/local/go/src/net/tcpsock_posix.go:117
	// _ = "end of CoverTab[16594]"
//line /usr/local/go/src/net/tcpsock_posix.go:117
	_go_fuzz_dep_.CoverTab[16595]++

//line /usr/local/go/src/net/tcpsock_posix.go:127
	if fd.laddr == nil || func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:127
		_go_fuzz_dep_.CoverTab[16599]++
//line /usr/local/go/src/net/tcpsock_posix.go:127
		return fd.raddr == nil
//line /usr/local/go/src/net/tcpsock_posix.go:127
		// _ = "end of CoverTab[16599]"
//line /usr/local/go/src/net/tcpsock_posix.go:127
	}() {
//line /usr/local/go/src/net/tcpsock_posix.go:127
		_go_fuzz_dep_.CoverTab[16600]++
								return true
//line /usr/local/go/src/net/tcpsock_posix.go:128
		// _ = "end of CoverTab[16600]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:129
		_go_fuzz_dep_.CoverTab[16601]++
//line /usr/local/go/src/net/tcpsock_posix.go:129
		// _ = "end of CoverTab[16601]"
//line /usr/local/go/src/net/tcpsock_posix.go:129
	}
//line /usr/local/go/src/net/tcpsock_posix.go:129
	// _ = "end of CoverTab[16595]"
//line /usr/local/go/src/net/tcpsock_posix.go:129
	_go_fuzz_dep_.CoverTab[16596]++
							l := fd.laddr.(*TCPAddr)
							r := fd.raddr.(*TCPAddr)
							return l.Port == r.Port && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:132
		_go_fuzz_dep_.CoverTab[16602]++
//line /usr/local/go/src/net/tcpsock_posix.go:132
		return l.IP.Equal(r.IP)
//line /usr/local/go/src/net/tcpsock_posix.go:132
		// _ = "end of CoverTab[16602]"
//line /usr/local/go/src/net/tcpsock_posix.go:132
	}()
//line /usr/local/go/src/net/tcpsock_posix.go:132
	// _ = "end of CoverTab[16596]"
}

func spuriousENOTAVAIL(err error) bool {
//line /usr/local/go/src/net/tcpsock_posix.go:135
	_go_fuzz_dep_.CoverTab[16603]++
							if op, ok := err.(*OpError); ok {
//line /usr/local/go/src/net/tcpsock_posix.go:136
		_go_fuzz_dep_.CoverTab[16606]++
								err = op.Err
//line /usr/local/go/src/net/tcpsock_posix.go:137
		// _ = "end of CoverTab[16606]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:138
		_go_fuzz_dep_.CoverTab[16607]++
//line /usr/local/go/src/net/tcpsock_posix.go:138
		// _ = "end of CoverTab[16607]"
//line /usr/local/go/src/net/tcpsock_posix.go:138
	}
//line /usr/local/go/src/net/tcpsock_posix.go:138
	// _ = "end of CoverTab[16603]"
//line /usr/local/go/src/net/tcpsock_posix.go:138
	_go_fuzz_dep_.CoverTab[16604]++
							if sys, ok := err.(*os.SyscallError); ok {
//line /usr/local/go/src/net/tcpsock_posix.go:139
		_go_fuzz_dep_.CoverTab[16608]++
								err = sys.Err
//line /usr/local/go/src/net/tcpsock_posix.go:140
		// _ = "end of CoverTab[16608]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:141
		_go_fuzz_dep_.CoverTab[16609]++
//line /usr/local/go/src/net/tcpsock_posix.go:141
		// _ = "end of CoverTab[16609]"
//line /usr/local/go/src/net/tcpsock_posix.go:141
	}
//line /usr/local/go/src/net/tcpsock_posix.go:141
	// _ = "end of CoverTab[16604]"
//line /usr/local/go/src/net/tcpsock_posix.go:141
	_go_fuzz_dep_.CoverTab[16605]++
							return err == syscall.EADDRNOTAVAIL
//line /usr/local/go/src/net/tcpsock_posix.go:142
	// _ = "end of CoverTab[16605]"
}

func (ln *TCPListener) ok() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:145
	_go_fuzz_dep_.CoverTab[16610]++
//line /usr/local/go/src/net/tcpsock_posix.go:145
	return ln != nil && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:145
		_go_fuzz_dep_.CoverTab[16611]++
//line /usr/local/go/src/net/tcpsock_posix.go:145
		return ln.fd != nil
//line /usr/local/go/src/net/tcpsock_posix.go:145
		// _ = "end of CoverTab[16611]"
//line /usr/local/go/src/net/tcpsock_posix.go:145
	}()
//line /usr/local/go/src/net/tcpsock_posix.go:145
	// _ = "end of CoverTab[16610]"
//line /usr/local/go/src/net/tcpsock_posix.go:145
}

func (ln *TCPListener) accept() (*TCPConn, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:147
	_go_fuzz_dep_.CoverTab[16612]++
							fd, err := ln.fd.accept()
							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:149
		_go_fuzz_dep_.CoverTab[16614]++
								return nil, err
//line /usr/local/go/src/net/tcpsock_posix.go:150
		// _ = "end of CoverTab[16614]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:151
		_go_fuzz_dep_.CoverTab[16615]++
//line /usr/local/go/src/net/tcpsock_posix.go:151
		// _ = "end of CoverTab[16615]"
//line /usr/local/go/src/net/tcpsock_posix.go:151
	}
//line /usr/local/go/src/net/tcpsock_posix.go:151
	// _ = "end of CoverTab[16612]"
//line /usr/local/go/src/net/tcpsock_posix.go:151
	_go_fuzz_dep_.CoverTab[16613]++
							return newTCPConn(fd, ln.lc.KeepAlive, nil), nil
//line /usr/local/go/src/net/tcpsock_posix.go:152
	// _ = "end of CoverTab[16613]"
}

func (ln *TCPListener) close() error {
//line /usr/local/go/src/net/tcpsock_posix.go:155
	_go_fuzz_dep_.CoverTab[16616]++
							return ln.fd.Close()
//line /usr/local/go/src/net/tcpsock_posix.go:156
	// _ = "end of CoverTab[16616]"
}

func (ln *TCPListener) file() (*os.File, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:159
	_go_fuzz_dep_.CoverTab[16617]++
							f, err := ln.fd.dup()
							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:161
		_go_fuzz_dep_.CoverTab[16619]++
								return nil, err
//line /usr/local/go/src/net/tcpsock_posix.go:162
		// _ = "end of CoverTab[16619]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:163
		_go_fuzz_dep_.CoverTab[16620]++
//line /usr/local/go/src/net/tcpsock_posix.go:163
		// _ = "end of CoverTab[16620]"
//line /usr/local/go/src/net/tcpsock_posix.go:163
	}
//line /usr/local/go/src/net/tcpsock_posix.go:163
	// _ = "end of CoverTab[16617]"
//line /usr/local/go/src/net/tcpsock_posix.go:163
	_go_fuzz_dep_.CoverTab[16618]++
							return f, nil
//line /usr/local/go/src/net/tcpsock_posix.go:164
	// _ = "end of CoverTab[16618]"
}

func (sl *sysListener) listenTCP(ctx context.Context, laddr *TCPAddr) (*TCPListener, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:167
	_go_fuzz_dep_.CoverTab[16621]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:169
		_go_fuzz_dep_.CoverTab[16624]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/tcpsock_posix.go:170
			_go_fuzz_dep_.CoverTab[16625]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/tcpsock_posix.go:171
			// _ = "end of CoverTab[16625]"
		}
//line /usr/local/go/src/net/tcpsock_posix.go:172
		// _ = "end of CoverTab[16624]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:173
		_go_fuzz_dep_.CoverTab[16626]++
//line /usr/local/go/src/net/tcpsock_posix.go:173
		// _ = "end of CoverTab[16626]"
//line /usr/local/go/src/net/tcpsock_posix.go:173
	}
//line /usr/local/go/src/net/tcpsock_posix.go:173
	// _ = "end of CoverTab[16621]"
//line /usr/local/go/src/net/tcpsock_posix.go:173
	_go_fuzz_dep_.CoverTab[16622]++
							fd, err := internetSocket(ctx, sl.network, laddr, nil, syscall.SOCK_STREAM, 0, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:175
		_go_fuzz_dep_.CoverTab[16627]++
								return nil, err
//line /usr/local/go/src/net/tcpsock_posix.go:176
		// _ = "end of CoverTab[16627]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:177
		_go_fuzz_dep_.CoverTab[16628]++
//line /usr/local/go/src/net/tcpsock_posix.go:177
		// _ = "end of CoverTab[16628]"
//line /usr/local/go/src/net/tcpsock_posix.go:177
	}
//line /usr/local/go/src/net/tcpsock_posix.go:177
	// _ = "end of CoverTab[16622]"
//line /usr/local/go/src/net/tcpsock_posix.go:177
	_go_fuzz_dep_.CoverTab[16623]++
							return &TCPListener{fd: fd, lc: sl.ListenConfig}, nil
//line /usr/local/go/src/net/tcpsock_posix.go:178
	// _ = "end of CoverTab[16623]"
}

//line /usr/local/go/src/net/tcpsock_posix.go:179
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/tcpsock_posix.go:179
var _ = _go_fuzz_dep_.CoverTab
