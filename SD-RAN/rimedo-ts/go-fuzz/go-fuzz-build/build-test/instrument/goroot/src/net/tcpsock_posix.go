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
	_go_fuzz_dep_.CoverTab[8155]++
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/tcpsock_posix.go:18
		_go_fuzz_dep_.CoverTab[8157]++
								return &TCPAddr{IP: sa.Addr[0:], Port: sa.Port}
//line /usr/local/go/src/net/tcpsock_posix.go:19
		// _ = "end of CoverTab[8157]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/tcpsock_posix.go:20
		_go_fuzz_dep_.CoverTab[8158]++
								return &TCPAddr{IP: sa.Addr[0:], Port: sa.Port, Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/tcpsock_posix.go:21
		// _ = "end of CoverTab[8158]"
	}
//line /usr/local/go/src/net/tcpsock_posix.go:22
	// _ = "end of CoverTab[8155]"
//line /usr/local/go/src/net/tcpsock_posix.go:22
	_go_fuzz_dep_.CoverTab[8156]++
							return nil
//line /usr/local/go/src/net/tcpsock_posix.go:23
	// _ = "end of CoverTab[8156]"
}

func (a *TCPAddr) family() int {
//line /usr/local/go/src/net/tcpsock_posix.go:26
	_go_fuzz_dep_.CoverTab[8159]++
							if a == nil || func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:27
		_go_fuzz_dep_.CoverTab[8162]++
//line /usr/local/go/src/net/tcpsock_posix.go:27
		return len(a.IP) <= IPv4len
//line /usr/local/go/src/net/tcpsock_posix.go:27
		// _ = "end of CoverTab[8162]"
//line /usr/local/go/src/net/tcpsock_posix.go:27
	}() {
//line /usr/local/go/src/net/tcpsock_posix.go:27
		_go_fuzz_dep_.CoverTab[8163]++
								return syscall.AF_INET
//line /usr/local/go/src/net/tcpsock_posix.go:28
		// _ = "end of CoverTab[8163]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:29
		_go_fuzz_dep_.CoverTab[8164]++
//line /usr/local/go/src/net/tcpsock_posix.go:29
		// _ = "end of CoverTab[8164]"
//line /usr/local/go/src/net/tcpsock_posix.go:29
	}
//line /usr/local/go/src/net/tcpsock_posix.go:29
	// _ = "end of CoverTab[8159]"
//line /usr/local/go/src/net/tcpsock_posix.go:29
	_go_fuzz_dep_.CoverTab[8160]++
							if a.IP.To4() != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:30
		_go_fuzz_dep_.CoverTab[8165]++
								return syscall.AF_INET
//line /usr/local/go/src/net/tcpsock_posix.go:31
		// _ = "end of CoverTab[8165]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:32
		_go_fuzz_dep_.CoverTab[8166]++
//line /usr/local/go/src/net/tcpsock_posix.go:32
		// _ = "end of CoverTab[8166]"
//line /usr/local/go/src/net/tcpsock_posix.go:32
	}
//line /usr/local/go/src/net/tcpsock_posix.go:32
	// _ = "end of CoverTab[8160]"
//line /usr/local/go/src/net/tcpsock_posix.go:32
	_go_fuzz_dep_.CoverTab[8161]++
							return syscall.AF_INET6
//line /usr/local/go/src/net/tcpsock_posix.go:33
	// _ = "end of CoverTab[8161]"
}

func (a *TCPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:36
	_go_fuzz_dep_.CoverTab[8167]++
							if a == nil {
//line /usr/local/go/src/net/tcpsock_posix.go:37
		_go_fuzz_dep_.CoverTab[8169]++
								return nil, nil
//line /usr/local/go/src/net/tcpsock_posix.go:38
		// _ = "end of CoverTab[8169]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:39
		_go_fuzz_dep_.CoverTab[8170]++
//line /usr/local/go/src/net/tcpsock_posix.go:39
		// _ = "end of CoverTab[8170]"
//line /usr/local/go/src/net/tcpsock_posix.go:39
	}
//line /usr/local/go/src/net/tcpsock_posix.go:39
	// _ = "end of CoverTab[8167]"
//line /usr/local/go/src/net/tcpsock_posix.go:39
	_go_fuzz_dep_.CoverTab[8168]++
							return ipToSockaddr(family, a.IP, a.Port, a.Zone)
//line /usr/local/go/src/net/tcpsock_posix.go:40
	// _ = "end of CoverTab[8168]"
}

func (a *TCPAddr) toLocal(net string) sockaddr {
//line /usr/local/go/src/net/tcpsock_posix.go:43
	_go_fuzz_dep_.CoverTab[8171]++
							return &TCPAddr{loopbackIP(net), a.Port, a.Zone}
//line /usr/local/go/src/net/tcpsock_posix.go:44
	// _ = "end of CoverTab[8171]"
}

func (c *TCPConn) readFrom(r io.Reader) (int64, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:47
	_go_fuzz_dep_.CoverTab[8172]++
							if n, err, handled := splice(c.fd, r); handled {
//line /usr/local/go/src/net/tcpsock_posix.go:48
		_go_fuzz_dep_.CoverTab[8175]++
								return n, err
//line /usr/local/go/src/net/tcpsock_posix.go:49
		// _ = "end of CoverTab[8175]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:50
		_go_fuzz_dep_.CoverTab[8176]++
//line /usr/local/go/src/net/tcpsock_posix.go:50
		// _ = "end of CoverTab[8176]"
//line /usr/local/go/src/net/tcpsock_posix.go:50
	}
//line /usr/local/go/src/net/tcpsock_posix.go:50
	// _ = "end of CoverTab[8172]"
//line /usr/local/go/src/net/tcpsock_posix.go:50
	_go_fuzz_dep_.CoverTab[8173]++
							if n, err, handled := sendFile(c.fd, r); handled {
//line /usr/local/go/src/net/tcpsock_posix.go:51
		_go_fuzz_dep_.CoverTab[8177]++
								return n, err
//line /usr/local/go/src/net/tcpsock_posix.go:52
		// _ = "end of CoverTab[8177]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:53
		_go_fuzz_dep_.CoverTab[8178]++
//line /usr/local/go/src/net/tcpsock_posix.go:53
		// _ = "end of CoverTab[8178]"
//line /usr/local/go/src/net/tcpsock_posix.go:53
	}
//line /usr/local/go/src/net/tcpsock_posix.go:53
	// _ = "end of CoverTab[8173]"
//line /usr/local/go/src/net/tcpsock_posix.go:53
	_go_fuzz_dep_.CoverTab[8174]++
							return genericReadFrom(c, r)
//line /usr/local/go/src/net/tcpsock_posix.go:54
	// _ = "end of CoverTab[8174]"
}

func (sd *sysDialer) dialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:57
	_go_fuzz_dep_.CoverTab[8179]++
							if h := sd.testHookDialTCP; h != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:58
		_go_fuzz_dep_.CoverTab[8182]++
								return h(ctx, sd.network, laddr, raddr)
//line /usr/local/go/src/net/tcpsock_posix.go:59
		// _ = "end of CoverTab[8182]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:60
		_go_fuzz_dep_.CoverTab[8183]++
//line /usr/local/go/src/net/tcpsock_posix.go:60
		// _ = "end of CoverTab[8183]"
//line /usr/local/go/src/net/tcpsock_posix.go:60
	}
//line /usr/local/go/src/net/tcpsock_posix.go:60
	// _ = "end of CoverTab[8179]"
//line /usr/local/go/src/net/tcpsock_posix.go:60
	_go_fuzz_dep_.CoverTab[8180]++
							if h := testHookDialTCP; h != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:61
		_go_fuzz_dep_.CoverTab[8184]++
								return h(ctx, sd.network, laddr, raddr)
//line /usr/local/go/src/net/tcpsock_posix.go:62
		// _ = "end of CoverTab[8184]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:63
		_go_fuzz_dep_.CoverTab[8185]++
//line /usr/local/go/src/net/tcpsock_posix.go:63
		// _ = "end of CoverTab[8185]"
//line /usr/local/go/src/net/tcpsock_posix.go:63
	}
//line /usr/local/go/src/net/tcpsock_posix.go:63
	// _ = "end of CoverTab[8180]"
//line /usr/local/go/src/net/tcpsock_posix.go:63
	_go_fuzz_dep_.CoverTab[8181]++
							return sd.doDialTCP(ctx, laddr, raddr)
//line /usr/local/go/src/net/tcpsock_posix.go:64
	// _ = "end of CoverTab[8181]"
}

func (sd *sysDialer) doDialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:67
	_go_fuzz_dep_.CoverTab[8186]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:69
		_go_fuzz_dep_.CoverTab[8190]++
//line /usr/local/go/src/net/tcpsock_posix.go:69
		return sd.Dialer.Control != nil
//line /usr/local/go/src/net/tcpsock_posix.go:69
		// _ = "end of CoverTab[8190]"
//line /usr/local/go/src/net/tcpsock_posix.go:69
	}() {
//line /usr/local/go/src/net/tcpsock_posix.go:69
		_go_fuzz_dep_.CoverTab[8191]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/tcpsock_posix.go:70
			_go_fuzz_dep_.CoverTab[8192]++
									return sd.Dialer.Control(network, address, c)
//line /usr/local/go/src/net/tcpsock_posix.go:71
			// _ = "end of CoverTab[8192]"
		}
//line /usr/local/go/src/net/tcpsock_posix.go:72
		// _ = "end of CoverTab[8191]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:73
		_go_fuzz_dep_.CoverTab[8193]++
//line /usr/local/go/src/net/tcpsock_posix.go:73
		// _ = "end of CoverTab[8193]"
//line /usr/local/go/src/net/tcpsock_posix.go:73
	}
//line /usr/local/go/src/net/tcpsock_posix.go:73
	// _ = "end of CoverTab[8186]"
//line /usr/local/go/src/net/tcpsock_posix.go:73
	_go_fuzz_dep_.CoverTab[8187]++
							fd, err := internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_STREAM, 0, "dial", ctrlCtxFn)

//line /usr/local/go/src/net/tcpsock_posix.go:100
	for i := 0; i < 2 && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:100
		_go_fuzz_dep_.CoverTab[8194]++
//line /usr/local/go/src/net/tcpsock_posix.go:100
		return (laddr == nil || func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:100
			_go_fuzz_dep_.CoverTab[8195]++
//line /usr/local/go/src/net/tcpsock_posix.go:100
			return laddr.Port == 0
//line /usr/local/go/src/net/tcpsock_posix.go:100
			// _ = "end of CoverTab[8195]"
//line /usr/local/go/src/net/tcpsock_posix.go:100
		}())
//line /usr/local/go/src/net/tcpsock_posix.go:100
		// _ = "end of CoverTab[8194]"
//line /usr/local/go/src/net/tcpsock_posix.go:100
	}() && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:100
		_go_fuzz_dep_.CoverTab[8196]++
//line /usr/local/go/src/net/tcpsock_posix.go:100
		return (selfConnect(fd, err) || func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:100
			_go_fuzz_dep_.CoverTab[8197]++
//line /usr/local/go/src/net/tcpsock_posix.go:100
			return spuriousENOTAVAIL(err)
//line /usr/local/go/src/net/tcpsock_posix.go:100
			// _ = "end of CoverTab[8197]"
//line /usr/local/go/src/net/tcpsock_posix.go:100
		}())
//line /usr/local/go/src/net/tcpsock_posix.go:100
		// _ = "end of CoverTab[8196]"
//line /usr/local/go/src/net/tcpsock_posix.go:100
	}(); i++ {
//line /usr/local/go/src/net/tcpsock_posix.go:100
		_go_fuzz_dep_.CoverTab[8198]++
								if err == nil {
//line /usr/local/go/src/net/tcpsock_posix.go:101
			_go_fuzz_dep_.CoverTab[8200]++
									fd.Close()
//line /usr/local/go/src/net/tcpsock_posix.go:102
			// _ = "end of CoverTab[8200]"
		} else {
//line /usr/local/go/src/net/tcpsock_posix.go:103
			_go_fuzz_dep_.CoverTab[8201]++
//line /usr/local/go/src/net/tcpsock_posix.go:103
			// _ = "end of CoverTab[8201]"
//line /usr/local/go/src/net/tcpsock_posix.go:103
		}
//line /usr/local/go/src/net/tcpsock_posix.go:103
		// _ = "end of CoverTab[8198]"
//line /usr/local/go/src/net/tcpsock_posix.go:103
		_go_fuzz_dep_.CoverTab[8199]++
								fd, err = internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_STREAM, 0, "dial", ctrlCtxFn)
//line /usr/local/go/src/net/tcpsock_posix.go:104
		// _ = "end of CoverTab[8199]"
	}
//line /usr/local/go/src/net/tcpsock_posix.go:105
	// _ = "end of CoverTab[8187]"
//line /usr/local/go/src/net/tcpsock_posix.go:105
	_go_fuzz_dep_.CoverTab[8188]++

							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:107
		_go_fuzz_dep_.CoverTab[8202]++
								return nil, err
//line /usr/local/go/src/net/tcpsock_posix.go:108
		// _ = "end of CoverTab[8202]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:109
		_go_fuzz_dep_.CoverTab[8203]++
//line /usr/local/go/src/net/tcpsock_posix.go:109
		// _ = "end of CoverTab[8203]"
//line /usr/local/go/src/net/tcpsock_posix.go:109
	}
//line /usr/local/go/src/net/tcpsock_posix.go:109
	// _ = "end of CoverTab[8188]"
//line /usr/local/go/src/net/tcpsock_posix.go:109
	_go_fuzz_dep_.CoverTab[8189]++
							return newTCPConn(fd, sd.Dialer.KeepAlive, testHookSetKeepAlive), nil
//line /usr/local/go/src/net/tcpsock_posix.go:110
	// _ = "end of CoverTab[8189]"
}

func selfConnect(fd *netFD, err error) bool {
//line /usr/local/go/src/net/tcpsock_posix.go:113
	_go_fuzz_dep_.CoverTab[8204]++

							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:115
		_go_fuzz_dep_.CoverTab[8207]++
								return false
//line /usr/local/go/src/net/tcpsock_posix.go:116
		// _ = "end of CoverTab[8207]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:117
		_go_fuzz_dep_.CoverTab[8208]++
//line /usr/local/go/src/net/tcpsock_posix.go:117
		// _ = "end of CoverTab[8208]"
//line /usr/local/go/src/net/tcpsock_posix.go:117
	}
//line /usr/local/go/src/net/tcpsock_posix.go:117
	// _ = "end of CoverTab[8204]"
//line /usr/local/go/src/net/tcpsock_posix.go:117
	_go_fuzz_dep_.CoverTab[8205]++

//line /usr/local/go/src/net/tcpsock_posix.go:127
	if fd.laddr == nil || func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:127
		_go_fuzz_dep_.CoverTab[8209]++
//line /usr/local/go/src/net/tcpsock_posix.go:127
		return fd.raddr == nil
//line /usr/local/go/src/net/tcpsock_posix.go:127
		// _ = "end of CoverTab[8209]"
//line /usr/local/go/src/net/tcpsock_posix.go:127
	}() {
//line /usr/local/go/src/net/tcpsock_posix.go:127
		_go_fuzz_dep_.CoverTab[8210]++
								return true
//line /usr/local/go/src/net/tcpsock_posix.go:128
		// _ = "end of CoverTab[8210]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:129
		_go_fuzz_dep_.CoverTab[8211]++
//line /usr/local/go/src/net/tcpsock_posix.go:129
		// _ = "end of CoverTab[8211]"
//line /usr/local/go/src/net/tcpsock_posix.go:129
	}
//line /usr/local/go/src/net/tcpsock_posix.go:129
	// _ = "end of CoverTab[8205]"
//line /usr/local/go/src/net/tcpsock_posix.go:129
	_go_fuzz_dep_.CoverTab[8206]++
							l := fd.laddr.(*TCPAddr)
							r := fd.raddr.(*TCPAddr)
							return l.Port == r.Port && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:132
		_go_fuzz_dep_.CoverTab[8212]++
//line /usr/local/go/src/net/tcpsock_posix.go:132
		return l.IP.Equal(r.IP)
//line /usr/local/go/src/net/tcpsock_posix.go:132
		// _ = "end of CoverTab[8212]"
//line /usr/local/go/src/net/tcpsock_posix.go:132
	}()
//line /usr/local/go/src/net/tcpsock_posix.go:132
	// _ = "end of CoverTab[8206]"
}

func spuriousENOTAVAIL(err error) bool {
//line /usr/local/go/src/net/tcpsock_posix.go:135
	_go_fuzz_dep_.CoverTab[8213]++
							if op, ok := err.(*OpError); ok {
//line /usr/local/go/src/net/tcpsock_posix.go:136
		_go_fuzz_dep_.CoverTab[8216]++
								err = op.Err
//line /usr/local/go/src/net/tcpsock_posix.go:137
		// _ = "end of CoverTab[8216]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:138
		_go_fuzz_dep_.CoverTab[8217]++
//line /usr/local/go/src/net/tcpsock_posix.go:138
		// _ = "end of CoverTab[8217]"
//line /usr/local/go/src/net/tcpsock_posix.go:138
	}
//line /usr/local/go/src/net/tcpsock_posix.go:138
	// _ = "end of CoverTab[8213]"
//line /usr/local/go/src/net/tcpsock_posix.go:138
	_go_fuzz_dep_.CoverTab[8214]++
							if sys, ok := err.(*os.SyscallError); ok {
//line /usr/local/go/src/net/tcpsock_posix.go:139
		_go_fuzz_dep_.CoverTab[8218]++
								err = sys.Err
//line /usr/local/go/src/net/tcpsock_posix.go:140
		// _ = "end of CoverTab[8218]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:141
		_go_fuzz_dep_.CoverTab[8219]++
//line /usr/local/go/src/net/tcpsock_posix.go:141
		// _ = "end of CoverTab[8219]"
//line /usr/local/go/src/net/tcpsock_posix.go:141
	}
//line /usr/local/go/src/net/tcpsock_posix.go:141
	// _ = "end of CoverTab[8214]"
//line /usr/local/go/src/net/tcpsock_posix.go:141
	_go_fuzz_dep_.CoverTab[8215]++
							return err == syscall.EADDRNOTAVAIL
//line /usr/local/go/src/net/tcpsock_posix.go:142
	// _ = "end of CoverTab[8215]"
}

func (ln *TCPListener) ok() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:145
	_go_fuzz_dep_.CoverTab[8220]++
//line /usr/local/go/src/net/tcpsock_posix.go:145
	return ln != nil && func() bool {
//line /usr/local/go/src/net/tcpsock_posix.go:145
		_go_fuzz_dep_.CoverTab[8221]++
//line /usr/local/go/src/net/tcpsock_posix.go:145
		return ln.fd != nil
//line /usr/local/go/src/net/tcpsock_posix.go:145
		// _ = "end of CoverTab[8221]"
//line /usr/local/go/src/net/tcpsock_posix.go:145
	}()
//line /usr/local/go/src/net/tcpsock_posix.go:145
	// _ = "end of CoverTab[8220]"
//line /usr/local/go/src/net/tcpsock_posix.go:145
}

func (ln *TCPListener) accept() (*TCPConn, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:147
	_go_fuzz_dep_.CoverTab[8222]++
							fd, err := ln.fd.accept()
							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:149
		_go_fuzz_dep_.CoverTab[8224]++
								return nil, err
//line /usr/local/go/src/net/tcpsock_posix.go:150
		// _ = "end of CoverTab[8224]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:151
		_go_fuzz_dep_.CoverTab[8225]++
//line /usr/local/go/src/net/tcpsock_posix.go:151
		// _ = "end of CoverTab[8225]"
//line /usr/local/go/src/net/tcpsock_posix.go:151
	}
//line /usr/local/go/src/net/tcpsock_posix.go:151
	// _ = "end of CoverTab[8222]"
//line /usr/local/go/src/net/tcpsock_posix.go:151
	_go_fuzz_dep_.CoverTab[8223]++
							return newTCPConn(fd, ln.lc.KeepAlive, nil), nil
//line /usr/local/go/src/net/tcpsock_posix.go:152
	// _ = "end of CoverTab[8223]"
}

func (ln *TCPListener) close() error {
//line /usr/local/go/src/net/tcpsock_posix.go:155
	_go_fuzz_dep_.CoverTab[8226]++
							return ln.fd.Close()
//line /usr/local/go/src/net/tcpsock_posix.go:156
	// _ = "end of CoverTab[8226]"
}

func (ln *TCPListener) file() (*os.File, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:159
	_go_fuzz_dep_.CoverTab[8227]++
							f, err := ln.fd.dup()
							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:161
		_go_fuzz_dep_.CoverTab[8229]++
								return nil, err
//line /usr/local/go/src/net/tcpsock_posix.go:162
		// _ = "end of CoverTab[8229]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:163
		_go_fuzz_dep_.CoverTab[8230]++
//line /usr/local/go/src/net/tcpsock_posix.go:163
		// _ = "end of CoverTab[8230]"
//line /usr/local/go/src/net/tcpsock_posix.go:163
	}
//line /usr/local/go/src/net/tcpsock_posix.go:163
	// _ = "end of CoverTab[8227]"
//line /usr/local/go/src/net/tcpsock_posix.go:163
	_go_fuzz_dep_.CoverTab[8228]++
							return f, nil
//line /usr/local/go/src/net/tcpsock_posix.go:164
	// _ = "end of CoverTab[8228]"
}

func (sl *sysListener) listenTCP(ctx context.Context, laddr *TCPAddr) (*TCPListener, error) {
//line /usr/local/go/src/net/tcpsock_posix.go:167
	_go_fuzz_dep_.CoverTab[8231]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:169
		_go_fuzz_dep_.CoverTab[8234]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/tcpsock_posix.go:170
			_go_fuzz_dep_.CoverTab[8235]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/tcpsock_posix.go:171
			// _ = "end of CoverTab[8235]"
		}
//line /usr/local/go/src/net/tcpsock_posix.go:172
		// _ = "end of CoverTab[8234]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:173
		_go_fuzz_dep_.CoverTab[8236]++
//line /usr/local/go/src/net/tcpsock_posix.go:173
		// _ = "end of CoverTab[8236]"
//line /usr/local/go/src/net/tcpsock_posix.go:173
	}
//line /usr/local/go/src/net/tcpsock_posix.go:173
	// _ = "end of CoverTab[8231]"
//line /usr/local/go/src/net/tcpsock_posix.go:173
	_go_fuzz_dep_.CoverTab[8232]++
							fd, err := internetSocket(ctx, sl.network, laddr, nil, syscall.SOCK_STREAM, 0, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/tcpsock_posix.go:175
		_go_fuzz_dep_.CoverTab[8237]++
								return nil, err
//line /usr/local/go/src/net/tcpsock_posix.go:176
		// _ = "end of CoverTab[8237]"
	} else {
//line /usr/local/go/src/net/tcpsock_posix.go:177
		_go_fuzz_dep_.CoverTab[8238]++
//line /usr/local/go/src/net/tcpsock_posix.go:177
		// _ = "end of CoverTab[8238]"
//line /usr/local/go/src/net/tcpsock_posix.go:177
	}
//line /usr/local/go/src/net/tcpsock_posix.go:177
	// _ = "end of CoverTab[8232]"
//line /usr/local/go/src/net/tcpsock_posix.go:177
	_go_fuzz_dep_.CoverTab[8233]++
							return &TCPListener{fd: fd, lc: sl.ListenConfig}, nil
//line /usr/local/go/src/net/tcpsock_posix.go:178
	// _ = "end of CoverTab[8233]"
}

//line /usr/local/go/src/net/tcpsock_posix.go:179
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/tcpsock_posix.go:179
var _ = _go_fuzz_dep_.CoverTab
