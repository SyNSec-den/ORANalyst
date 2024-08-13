// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || windows

//line /usr/local/go/src/net/iprawsock_posix.go:7
package net

//line /usr/local/go/src/net/iprawsock_posix.go:7
import (
//line /usr/local/go/src/net/iprawsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/iprawsock_posix.go:7
)
//line /usr/local/go/src/net/iprawsock_posix.go:7
import (
//line /usr/local/go/src/net/iprawsock_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/iprawsock_posix.go:7
)

import (
	"context"
	"syscall"
)

func sockaddrToIP(sa syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/iprawsock_posix.go:14
	_go_fuzz_dep_.CoverTab[6366]++
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/iprawsock_posix.go:16
		_go_fuzz_dep_.CoverTab[6368]++
								return &IPAddr{IP: sa.Addr[0:]}
//line /usr/local/go/src/net/iprawsock_posix.go:17
		// _ = "end of CoverTab[6368]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/iprawsock_posix.go:18
		_go_fuzz_dep_.CoverTab[6369]++
								return &IPAddr{IP: sa.Addr[0:], Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/iprawsock_posix.go:19
		// _ = "end of CoverTab[6369]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:20
	// _ = "end of CoverTab[6366]"
//line /usr/local/go/src/net/iprawsock_posix.go:20
	_go_fuzz_dep_.CoverTab[6367]++
							return nil
//line /usr/local/go/src/net/iprawsock_posix.go:21
	// _ = "end of CoverTab[6367]"
}

func (a *IPAddr) family() int {
//line /usr/local/go/src/net/iprawsock_posix.go:24
	_go_fuzz_dep_.CoverTab[6370]++
							if a == nil || func() bool {
//line /usr/local/go/src/net/iprawsock_posix.go:25
		_go_fuzz_dep_.CoverTab[6373]++
//line /usr/local/go/src/net/iprawsock_posix.go:25
		return len(a.IP) <= IPv4len
//line /usr/local/go/src/net/iprawsock_posix.go:25
		// _ = "end of CoverTab[6373]"
//line /usr/local/go/src/net/iprawsock_posix.go:25
	}() {
//line /usr/local/go/src/net/iprawsock_posix.go:25
		_go_fuzz_dep_.CoverTab[6374]++
								return syscall.AF_INET
//line /usr/local/go/src/net/iprawsock_posix.go:26
		// _ = "end of CoverTab[6374]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:27
		_go_fuzz_dep_.CoverTab[6375]++
//line /usr/local/go/src/net/iprawsock_posix.go:27
		// _ = "end of CoverTab[6375]"
//line /usr/local/go/src/net/iprawsock_posix.go:27
	}
//line /usr/local/go/src/net/iprawsock_posix.go:27
	// _ = "end of CoverTab[6370]"
//line /usr/local/go/src/net/iprawsock_posix.go:27
	_go_fuzz_dep_.CoverTab[6371]++
							if a.IP.To4() != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:28
		_go_fuzz_dep_.CoverTab[6376]++
								return syscall.AF_INET
//line /usr/local/go/src/net/iprawsock_posix.go:29
		// _ = "end of CoverTab[6376]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:30
		_go_fuzz_dep_.CoverTab[6377]++
//line /usr/local/go/src/net/iprawsock_posix.go:30
		// _ = "end of CoverTab[6377]"
//line /usr/local/go/src/net/iprawsock_posix.go:30
	}
//line /usr/local/go/src/net/iprawsock_posix.go:30
	// _ = "end of CoverTab[6371]"
//line /usr/local/go/src/net/iprawsock_posix.go:30
	_go_fuzz_dep_.CoverTab[6372]++
							return syscall.AF_INET6
//line /usr/local/go/src/net/iprawsock_posix.go:31
	// _ = "end of CoverTab[6372]"
}

func (a *IPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:34
	_go_fuzz_dep_.CoverTab[6378]++
							if a == nil {
//line /usr/local/go/src/net/iprawsock_posix.go:35
		_go_fuzz_dep_.CoverTab[6380]++
								return nil, nil
//line /usr/local/go/src/net/iprawsock_posix.go:36
		// _ = "end of CoverTab[6380]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:37
		_go_fuzz_dep_.CoverTab[6381]++
//line /usr/local/go/src/net/iprawsock_posix.go:37
		// _ = "end of CoverTab[6381]"
//line /usr/local/go/src/net/iprawsock_posix.go:37
	}
//line /usr/local/go/src/net/iprawsock_posix.go:37
	// _ = "end of CoverTab[6378]"
//line /usr/local/go/src/net/iprawsock_posix.go:37
	_go_fuzz_dep_.CoverTab[6379]++
							return ipToSockaddr(family, a.IP, 0, a.Zone)
//line /usr/local/go/src/net/iprawsock_posix.go:38
	// _ = "end of CoverTab[6379]"
}

func (a *IPAddr) toLocal(net string) sockaddr {
//line /usr/local/go/src/net/iprawsock_posix.go:41
	_go_fuzz_dep_.CoverTab[6382]++
							return &IPAddr{loopbackIP(net), a.Zone}
//line /usr/local/go/src/net/iprawsock_posix.go:42
	// _ = "end of CoverTab[6382]"
}

func (c *IPConn) readFrom(b []byte) (int, *IPAddr, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:45
	_go_fuzz_dep_.CoverTab[6383]++
	// TODO(cw,rsc): consider using readv if we know the family
	// type to avoid the header trim/copy
	var addr *IPAddr
	n, sa, err := c.fd.readFrom(b)
	switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/iprawsock_posix.go:51
		_go_fuzz_dep_.CoverTab[6385]++
								addr = &IPAddr{IP: sa.Addr[0:]}
								n = stripIPv4Header(n, b)
//line /usr/local/go/src/net/iprawsock_posix.go:53
		// _ = "end of CoverTab[6385]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/iprawsock_posix.go:54
		_go_fuzz_dep_.CoverTab[6386]++
								addr = &IPAddr{IP: sa.Addr[0:], Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/iprawsock_posix.go:55
		// _ = "end of CoverTab[6386]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:56
	// _ = "end of CoverTab[6383]"
//line /usr/local/go/src/net/iprawsock_posix.go:56
	_go_fuzz_dep_.CoverTab[6384]++
							return n, addr, err
//line /usr/local/go/src/net/iprawsock_posix.go:57
	// _ = "end of CoverTab[6384]"
}

func stripIPv4Header(n int, b []byte) int {
//line /usr/local/go/src/net/iprawsock_posix.go:60
	_go_fuzz_dep_.CoverTab[6387]++
							if len(b) < 20 {
//line /usr/local/go/src/net/iprawsock_posix.go:61
		_go_fuzz_dep_.CoverTab[6391]++
								return n
//line /usr/local/go/src/net/iprawsock_posix.go:62
		// _ = "end of CoverTab[6391]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:63
		_go_fuzz_dep_.CoverTab[6392]++
//line /usr/local/go/src/net/iprawsock_posix.go:63
		// _ = "end of CoverTab[6392]"
//line /usr/local/go/src/net/iprawsock_posix.go:63
	}
//line /usr/local/go/src/net/iprawsock_posix.go:63
	// _ = "end of CoverTab[6387]"
//line /usr/local/go/src/net/iprawsock_posix.go:63
	_go_fuzz_dep_.CoverTab[6388]++
							l := int(b[0]&0x0f) << 2
							if 20 > l || func() bool {
//line /usr/local/go/src/net/iprawsock_posix.go:65
		_go_fuzz_dep_.CoverTab[6393]++
//line /usr/local/go/src/net/iprawsock_posix.go:65
		return l > len(b)
//line /usr/local/go/src/net/iprawsock_posix.go:65
		// _ = "end of CoverTab[6393]"
//line /usr/local/go/src/net/iprawsock_posix.go:65
	}() {
//line /usr/local/go/src/net/iprawsock_posix.go:65
		_go_fuzz_dep_.CoverTab[6394]++
								return n
//line /usr/local/go/src/net/iprawsock_posix.go:66
		// _ = "end of CoverTab[6394]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:67
		_go_fuzz_dep_.CoverTab[6395]++
//line /usr/local/go/src/net/iprawsock_posix.go:67
		// _ = "end of CoverTab[6395]"
//line /usr/local/go/src/net/iprawsock_posix.go:67
	}
//line /usr/local/go/src/net/iprawsock_posix.go:67
	// _ = "end of CoverTab[6388]"
//line /usr/local/go/src/net/iprawsock_posix.go:67
	_go_fuzz_dep_.CoverTab[6389]++
							if b[0]>>4 != 4 {
//line /usr/local/go/src/net/iprawsock_posix.go:68
		_go_fuzz_dep_.CoverTab[6396]++
								return n
//line /usr/local/go/src/net/iprawsock_posix.go:69
		// _ = "end of CoverTab[6396]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:70
		_go_fuzz_dep_.CoverTab[6397]++
//line /usr/local/go/src/net/iprawsock_posix.go:70
		// _ = "end of CoverTab[6397]"
//line /usr/local/go/src/net/iprawsock_posix.go:70
	}
//line /usr/local/go/src/net/iprawsock_posix.go:70
	// _ = "end of CoverTab[6389]"
//line /usr/local/go/src/net/iprawsock_posix.go:70
	_go_fuzz_dep_.CoverTab[6390]++
							copy(b, b[l:])
							return n - l
//line /usr/local/go/src/net/iprawsock_posix.go:72
	// _ = "end of CoverTab[6390]"
}

func (c *IPConn) readMsg(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) {
//line /usr/local/go/src/net/iprawsock_posix.go:75
	_go_fuzz_dep_.CoverTab[6398]++
							var sa syscall.Sockaddr
							n, oobn, flags, sa, err = c.fd.readMsg(b, oob, 0)
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/iprawsock_posix.go:79
		_go_fuzz_dep_.CoverTab[6400]++
								addr = &IPAddr{IP: sa.Addr[0:]}
//line /usr/local/go/src/net/iprawsock_posix.go:80
		// _ = "end of CoverTab[6400]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/iprawsock_posix.go:81
		_go_fuzz_dep_.CoverTab[6401]++
								addr = &IPAddr{IP: sa.Addr[0:], Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/iprawsock_posix.go:82
		// _ = "end of CoverTab[6401]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:83
	// _ = "end of CoverTab[6398]"
//line /usr/local/go/src/net/iprawsock_posix.go:83
	_go_fuzz_dep_.CoverTab[6399]++
							return
//line /usr/local/go/src/net/iprawsock_posix.go:84
	// _ = "end of CoverTab[6399]"
}

func (c *IPConn) writeTo(b []byte, addr *IPAddr) (int, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:87
	_go_fuzz_dep_.CoverTab[6402]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/iprawsock_posix.go:88
		_go_fuzz_dep_.CoverTab[6406]++
								return 0, ErrWriteToConnected
//line /usr/local/go/src/net/iprawsock_posix.go:89
		// _ = "end of CoverTab[6406]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:90
		_go_fuzz_dep_.CoverTab[6407]++
//line /usr/local/go/src/net/iprawsock_posix.go:90
		// _ = "end of CoverTab[6407]"
//line /usr/local/go/src/net/iprawsock_posix.go:90
	}
//line /usr/local/go/src/net/iprawsock_posix.go:90
	// _ = "end of CoverTab[6402]"
//line /usr/local/go/src/net/iprawsock_posix.go:90
	_go_fuzz_dep_.CoverTab[6403]++
							if addr == nil {
//line /usr/local/go/src/net/iprawsock_posix.go:91
		_go_fuzz_dep_.CoverTab[6408]++
								return 0, errMissingAddress
//line /usr/local/go/src/net/iprawsock_posix.go:92
		// _ = "end of CoverTab[6408]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:93
		_go_fuzz_dep_.CoverTab[6409]++
//line /usr/local/go/src/net/iprawsock_posix.go:93
		// _ = "end of CoverTab[6409]"
//line /usr/local/go/src/net/iprawsock_posix.go:93
	}
//line /usr/local/go/src/net/iprawsock_posix.go:93
	// _ = "end of CoverTab[6403]"
//line /usr/local/go/src/net/iprawsock_posix.go:93
	_go_fuzz_dep_.CoverTab[6404]++
							sa, err := addr.sockaddr(c.fd.family)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:95
		_go_fuzz_dep_.CoverTab[6410]++
								return 0, err
//line /usr/local/go/src/net/iprawsock_posix.go:96
		// _ = "end of CoverTab[6410]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:97
		_go_fuzz_dep_.CoverTab[6411]++
//line /usr/local/go/src/net/iprawsock_posix.go:97
		// _ = "end of CoverTab[6411]"
//line /usr/local/go/src/net/iprawsock_posix.go:97
	}
//line /usr/local/go/src/net/iprawsock_posix.go:97
	// _ = "end of CoverTab[6404]"
//line /usr/local/go/src/net/iprawsock_posix.go:97
	_go_fuzz_dep_.CoverTab[6405]++
							return c.fd.writeTo(b, sa)
//line /usr/local/go/src/net/iprawsock_posix.go:98
	// _ = "end of CoverTab[6405]"
}

func (c *IPConn) writeMsg(b, oob []byte, addr *IPAddr) (n, oobn int, err error) {
//line /usr/local/go/src/net/iprawsock_posix.go:101
	_go_fuzz_dep_.CoverTab[6412]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/iprawsock_posix.go:102
		_go_fuzz_dep_.CoverTab[6416]++
								return 0, 0, ErrWriteToConnected
//line /usr/local/go/src/net/iprawsock_posix.go:103
		// _ = "end of CoverTab[6416]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:104
		_go_fuzz_dep_.CoverTab[6417]++
//line /usr/local/go/src/net/iprawsock_posix.go:104
		// _ = "end of CoverTab[6417]"
//line /usr/local/go/src/net/iprawsock_posix.go:104
	}
//line /usr/local/go/src/net/iprawsock_posix.go:104
	// _ = "end of CoverTab[6412]"
//line /usr/local/go/src/net/iprawsock_posix.go:104
	_go_fuzz_dep_.CoverTab[6413]++
							if addr == nil {
//line /usr/local/go/src/net/iprawsock_posix.go:105
		_go_fuzz_dep_.CoverTab[6418]++
								return 0, 0, errMissingAddress
//line /usr/local/go/src/net/iprawsock_posix.go:106
		// _ = "end of CoverTab[6418]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:107
		_go_fuzz_dep_.CoverTab[6419]++
//line /usr/local/go/src/net/iprawsock_posix.go:107
		// _ = "end of CoverTab[6419]"
//line /usr/local/go/src/net/iprawsock_posix.go:107
	}
//line /usr/local/go/src/net/iprawsock_posix.go:107
	// _ = "end of CoverTab[6413]"
//line /usr/local/go/src/net/iprawsock_posix.go:107
	_go_fuzz_dep_.CoverTab[6414]++
							sa, err := addr.sockaddr(c.fd.family)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:109
		_go_fuzz_dep_.CoverTab[6420]++
								return 0, 0, err
//line /usr/local/go/src/net/iprawsock_posix.go:110
		// _ = "end of CoverTab[6420]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:111
		_go_fuzz_dep_.CoverTab[6421]++
//line /usr/local/go/src/net/iprawsock_posix.go:111
		// _ = "end of CoverTab[6421]"
//line /usr/local/go/src/net/iprawsock_posix.go:111
	}
//line /usr/local/go/src/net/iprawsock_posix.go:111
	// _ = "end of CoverTab[6414]"
//line /usr/local/go/src/net/iprawsock_posix.go:111
	_go_fuzz_dep_.CoverTab[6415]++
							return c.fd.writeMsg(b, oob, sa)
//line /usr/local/go/src/net/iprawsock_posix.go:112
	// _ = "end of CoverTab[6415]"
}

func (sd *sysDialer) dialIP(ctx context.Context, laddr, raddr *IPAddr) (*IPConn, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:115
	_go_fuzz_dep_.CoverTab[6422]++
							network, proto, err := parseNetwork(ctx, sd.network, true)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:117
		_go_fuzz_dep_.CoverTab[6427]++
								return nil, err
//line /usr/local/go/src/net/iprawsock_posix.go:118
		// _ = "end of CoverTab[6427]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:119
		_go_fuzz_dep_.CoverTab[6428]++
//line /usr/local/go/src/net/iprawsock_posix.go:119
		// _ = "end of CoverTab[6428]"
//line /usr/local/go/src/net/iprawsock_posix.go:119
	}
//line /usr/local/go/src/net/iprawsock_posix.go:119
	// _ = "end of CoverTab[6422]"
//line /usr/local/go/src/net/iprawsock_posix.go:119
	_go_fuzz_dep_.CoverTab[6423]++
							switch network {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/iprawsock_posix.go:121
		_go_fuzz_dep_.CoverTab[6429]++
//line /usr/local/go/src/net/iprawsock_posix.go:121
		// _ = "end of CoverTab[6429]"
	default:
//line /usr/local/go/src/net/iprawsock_posix.go:122
		_go_fuzz_dep_.CoverTab[6430]++
								return nil, UnknownNetworkError(sd.network)
//line /usr/local/go/src/net/iprawsock_posix.go:123
		// _ = "end of CoverTab[6430]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:124
	// _ = "end of CoverTab[6423]"
//line /usr/local/go/src/net/iprawsock_posix.go:124
	_go_fuzz_dep_.CoverTab[6424]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /usr/local/go/src/net/iprawsock_posix.go:126
		_go_fuzz_dep_.CoverTab[6431]++
//line /usr/local/go/src/net/iprawsock_posix.go:126
		return sd.Dialer.Control != nil
//line /usr/local/go/src/net/iprawsock_posix.go:126
		// _ = "end of CoverTab[6431]"
//line /usr/local/go/src/net/iprawsock_posix.go:126
	}() {
//line /usr/local/go/src/net/iprawsock_posix.go:126
		_go_fuzz_dep_.CoverTab[6432]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/iprawsock_posix.go:127
			_go_fuzz_dep_.CoverTab[6433]++
									return sd.Dialer.Control(network, address, c)
//line /usr/local/go/src/net/iprawsock_posix.go:128
			// _ = "end of CoverTab[6433]"
		}
//line /usr/local/go/src/net/iprawsock_posix.go:129
		// _ = "end of CoverTab[6432]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:130
		_go_fuzz_dep_.CoverTab[6434]++
//line /usr/local/go/src/net/iprawsock_posix.go:130
		// _ = "end of CoverTab[6434]"
//line /usr/local/go/src/net/iprawsock_posix.go:130
	}
//line /usr/local/go/src/net/iprawsock_posix.go:130
	// _ = "end of CoverTab[6424]"
//line /usr/local/go/src/net/iprawsock_posix.go:130
	_go_fuzz_dep_.CoverTab[6425]++
							fd, err := internetSocket(ctx, network, laddr, raddr, syscall.SOCK_RAW, proto, "dial", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:132
		_go_fuzz_dep_.CoverTab[6435]++
								return nil, err
//line /usr/local/go/src/net/iprawsock_posix.go:133
		// _ = "end of CoverTab[6435]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:134
		_go_fuzz_dep_.CoverTab[6436]++
//line /usr/local/go/src/net/iprawsock_posix.go:134
		// _ = "end of CoverTab[6436]"
//line /usr/local/go/src/net/iprawsock_posix.go:134
	}
//line /usr/local/go/src/net/iprawsock_posix.go:134
	// _ = "end of CoverTab[6425]"
//line /usr/local/go/src/net/iprawsock_posix.go:134
	_go_fuzz_dep_.CoverTab[6426]++
							return newIPConn(fd), nil
//line /usr/local/go/src/net/iprawsock_posix.go:135
	// _ = "end of CoverTab[6426]"
}

func (sl *sysListener) listenIP(ctx context.Context, laddr *IPAddr) (*IPConn, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:138
	_go_fuzz_dep_.CoverTab[6437]++
							network, proto, err := parseNetwork(ctx, sl.network, true)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:140
		_go_fuzz_dep_.CoverTab[6442]++
								return nil, err
//line /usr/local/go/src/net/iprawsock_posix.go:141
		// _ = "end of CoverTab[6442]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:142
		_go_fuzz_dep_.CoverTab[6443]++
//line /usr/local/go/src/net/iprawsock_posix.go:142
		// _ = "end of CoverTab[6443]"
//line /usr/local/go/src/net/iprawsock_posix.go:142
	}
//line /usr/local/go/src/net/iprawsock_posix.go:142
	// _ = "end of CoverTab[6437]"
//line /usr/local/go/src/net/iprawsock_posix.go:142
	_go_fuzz_dep_.CoverTab[6438]++
							switch network {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/iprawsock_posix.go:144
		_go_fuzz_dep_.CoverTab[6444]++
//line /usr/local/go/src/net/iprawsock_posix.go:144
		// _ = "end of CoverTab[6444]"
	default:
//line /usr/local/go/src/net/iprawsock_posix.go:145
		_go_fuzz_dep_.CoverTab[6445]++
								return nil, UnknownNetworkError(sl.network)
//line /usr/local/go/src/net/iprawsock_posix.go:146
		// _ = "end of CoverTab[6445]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:147
	// _ = "end of CoverTab[6438]"
//line /usr/local/go/src/net/iprawsock_posix.go:147
	_go_fuzz_dep_.CoverTab[6439]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:149
		_go_fuzz_dep_.CoverTab[6446]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/iprawsock_posix.go:150
			_go_fuzz_dep_.CoverTab[6447]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/iprawsock_posix.go:151
			// _ = "end of CoverTab[6447]"
		}
//line /usr/local/go/src/net/iprawsock_posix.go:152
		// _ = "end of CoverTab[6446]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:153
		_go_fuzz_dep_.CoverTab[6448]++
//line /usr/local/go/src/net/iprawsock_posix.go:153
		// _ = "end of CoverTab[6448]"
//line /usr/local/go/src/net/iprawsock_posix.go:153
	}
//line /usr/local/go/src/net/iprawsock_posix.go:153
	// _ = "end of CoverTab[6439]"
//line /usr/local/go/src/net/iprawsock_posix.go:153
	_go_fuzz_dep_.CoverTab[6440]++
							fd, err := internetSocket(ctx, network, laddr, nil, syscall.SOCK_RAW, proto, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:155
		_go_fuzz_dep_.CoverTab[6449]++
								return nil, err
//line /usr/local/go/src/net/iprawsock_posix.go:156
		// _ = "end of CoverTab[6449]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:157
		_go_fuzz_dep_.CoverTab[6450]++
//line /usr/local/go/src/net/iprawsock_posix.go:157
		// _ = "end of CoverTab[6450]"
//line /usr/local/go/src/net/iprawsock_posix.go:157
	}
//line /usr/local/go/src/net/iprawsock_posix.go:157
	// _ = "end of CoverTab[6440]"
//line /usr/local/go/src/net/iprawsock_posix.go:157
	_go_fuzz_dep_.CoverTab[6441]++
							return newIPConn(fd), nil
//line /usr/local/go/src/net/iprawsock_posix.go:158
	// _ = "end of CoverTab[6441]"
}

//line /usr/local/go/src/net/iprawsock_posix.go:159
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/iprawsock_posix.go:159
var _ = _go_fuzz_dep_.CoverTab
