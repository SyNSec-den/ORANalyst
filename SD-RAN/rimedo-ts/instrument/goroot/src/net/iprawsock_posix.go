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
	_go_fuzz_dep_.CoverTab[14756]++
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/iprawsock_posix.go:16
		_go_fuzz_dep_.CoverTab[14758]++
								return &IPAddr{IP: sa.Addr[0:]}
//line /usr/local/go/src/net/iprawsock_posix.go:17
		// _ = "end of CoverTab[14758]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/iprawsock_posix.go:18
		_go_fuzz_dep_.CoverTab[14759]++
								return &IPAddr{IP: sa.Addr[0:], Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/iprawsock_posix.go:19
		// _ = "end of CoverTab[14759]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:20
	// _ = "end of CoverTab[14756]"
//line /usr/local/go/src/net/iprawsock_posix.go:20
	_go_fuzz_dep_.CoverTab[14757]++
							return nil
//line /usr/local/go/src/net/iprawsock_posix.go:21
	// _ = "end of CoverTab[14757]"
}

func (a *IPAddr) family() int {
//line /usr/local/go/src/net/iprawsock_posix.go:24
	_go_fuzz_dep_.CoverTab[14760]++
							if a == nil || func() bool {
//line /usr/local/go/src/net/iprawsock_posix.go:25
		_go_fuzz_dep_.CoverTab[14763]++
//line /usr/local/go/src/net/iprawsock_posix.go:25
		return len(a.IP) <= IPv4len
//line /usr/local/go/src/net/iprawsock_posix.go:25
		// _ = "end of CoverTab[14763]"
//line /usr/local/go/src/net/iprawsock_posix.go:25
	}() {
//line /usr/local/go/src/net/iprawsock_posix.go:25
		_go_fuzz_dep_.CoverTab[14764]++
								return syscall.AF_INET
//line /usr/local/go/src/net/iprawsock_posix.go:26
		// _ = "end of CoverTab[14764]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:27
		_go_fuzz_dep_.CoverTab[14765]++
//line /usr/local/go/src/net/iprawsock_posix.go:27
		// _ = "end of CoverTab[14765]"
//line /usr/local/go/src/net/iprawsock_posix.go:27
	}
//line /usr/local/go/src/net/iprawsock_posix.go:27
	// _ = "end of CoverTab[14760]"
//line /usr/local/go/src/net/iprawsock_posix.go:27
	_go_fuzz_dep_.CoverTab[14761]++
							if a.IP.To4() != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:28
		_go_fuzz_dep_.CoverTab[14766]++
								return syscall.AF_INET
//line /usr/local/go/src/net/iprawsock_posix.go:29
		// _ = "end of CoverTab[14766]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:30
		_go_fuzz_dep_.CoverTab[14767]++
//line /usr/local/go/src/net/iprawsock_posix.go:30
		// _ = "end of CoverTab[14767]"
//line /usr/local/go/src/net/iprawsock_posix.go:30
	}
//line /usr/local/go/src/net/iprawsock_posix.go:30
	// _ = "end of CoverTab[14761]"
//line /usr/local/go/src/net/iprawsock_posix.go:30
	_go_fuzz_dep_.CoverTab[14762]++
							return syscall.AF_INET6
//line /usr/local/go/src/net/iprawsock_posix.go:31
	// _ = "end of CoverTab[14762]"
}

func (a *IPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:34
	_go_fuzz_dep_.CoverTab[14768]++
							if a == nil {
//line /usr/local/go/src/net/iprawsock_posix.go:35
		_go_fuzz_dep_.CoverTab[14770]++
								return nil, nil
//line /usr/local/go/src/net/iprawsock_posix.go:36
		// _ = "end of CoverTab[14770]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:37
		_go_fuzz_dep_.CoverTab[14771]++
//line /usr/local/go/src/net/iprawsock_posix.go:37
		// _ = "end of CoverTab[14771]"
//line /usr/local/go/src/net/iprawsock_posix.go:37
	}
//line /usr/local/go/src/net/iprawsock_posix.go:37
	// _ = "end of CoverTab[14768]"
//line /usr/local/go/src/net/iprawsock_posix.go:37
	_go_fuzz_dep_.CoverTab[14769]++
							return ipToSockaddr(family, a.IP, 0, a.Zone)
//line /usr/local/go/src/net/iprawsock_posix.go:38
	// _ = "end of CoverTab[14769]"
}

func (a *IPAddr) toLocal(net string) sockaddr {
//line /usr/local/go/src/net/iprawsock_posix.go:41
	_go_fuzz_dep_.CoverTab[14772]++
							return &IPAddr{loopbackIP(net), a.Zone}
//line /usr/local/go/src/net/iprawsock_posix.go:42
	// _ = "end of CoverTab[14772]"
}

func (c *IPConn) readFrom(b []byte) (int, *IPAddr, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:45
	_go_fuzz_dep_.CoverTab[14773]++
	// TODO(cw,rsc): consider using readv if we know the family
	// type to avoid the header trim/copy
	var addr *IPAddr
	n, sa, err := c.fd.readFrom(b)
	switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/iprawsock_posix.go:51
		_go_fuzz_dep_.CoverTab[14775]++
								addr = &IPAddr{IP: sa.Addr[0:]}
								n = stripIPv4Header(n, b)
//line /usr/local/go/src/net/iprawsock_posix.go:53
		// _ = "end of CoverTab[14775]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/iprawsock_posix.go:54
		_go_fuzz_dep_.CoverTab[14776]++
								addr = &IPAddr{IP: sa.Addr[0:], Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/iprawsock_posix.go:55
		// _ = "end of CoverTab[14776]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:56
	// _ = "end of CoverTab[14773]"
//line /usr/local/go/src/net/iprawsock_posix.go:56
	_go_fuzz_dep_.CoverTab[14774]++
							return n, addr, err
//line /usr/local/go/src/net/iprawsock_posix.go:57
	// _ = "end of CoverTab[14774]"
}

func stripIPv4Header(n int, b []byte) int {
//line /usr/local/go/src/net/iprawsock_posix.go:60
	_go_fuzz_dep_.CoverTab[14777]++
							if len(b) < 20 {
//line /usr/local/go/src/net/iprawsock_posix.go:61
		_go_fuzz_dep_.CoverTab[14781]++
								return n
//line /usr/local/go/src/net/iprawsock_posix.go:62
		// _ = "end of CoverTab[14781]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:63
		_go_fuzz_dep_.CoverTab[14782]++
//line /usr/local/go/src/net/iprawsock_posix.go:63
		// _ = "end of CoverTab[14782]"
//line /usr/local/go/src/net/iprawsock_posix.go:63
	}
//line /usr/local/go/src/net/iprawsock_posix.go:63
	// _ = "end of CoverTab[14777]"
//line /usr/local/go/src/net/iprawsock_posix.go:63
	_go_fuzz_dep_.CoverTab[14778]++
							l := int(b[0]&0x0f) << 2
							if 20 > l || func() bool {
//line /usr/local/go/src/net/iprawsock_posix.go:65
		_go_fuzz_dep_.CoverTab[14783]++
//line /usr/local/go/src/net/iprawsock_posix.go:65
		return l > len(b)
//line /usr/local/go/src/net/iprawsock_posix.go:65
		// _ = "end of CoverTab[14783]"
//line /usr/local/go/src/net/iprawsock_posix.go:65
	}() {
//line /usr/local/go/src/net/iprawsock_posix.go:65
		_go_fuzz_dep_.CoverTab[14784]++
								return n
//line /usr/local/go/src/net/iprawsock_posix.go:66
		// _ = "end of CoverTab[14784]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:67
		_go_fuzz_dep_.CoverTab[14785]++
//line /usr/local/go/src/net/iprawsock_posix.go:67
		// _ = "end of CoverTab[14785]"
//line /usr/local/go/src/net/iprawsock_posix.go:67
	}
//line /usr/local/go/src/net/iprawsock_posix.go:67
	// _ = "end of CoverTab[14778]"
//line /usr/local/go/src/net/iprawsock_posix.go:67
	_go_fuzz_dep_.CoverTab[14779]++
							if b[0]>>4 != 4 {
//line /usr/local/go/src/net/iprawsock_posix.go:68
		_go_fuzz_dep_.CoverTab[14786]++
								return n
//line /usr/local/go/src/net/iprawsock_posix.go:69
		// _ = "end of CoverTab[14786]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:70
		_go_fuzz_dep_.CoverTab[14787]++
//line /usr/local/go/src/net/iprawsock_posix.go:70
		// _ = "end of CoverTab[14787]"
//line /usr/local/go/src/net/iprawsock_posix.go:70
	}
//line /usr/local/go/src/net/iprawsock_posix.go:70
	// _ = "end of CoverTab[14779]"
//line /usr/local/go/src/net/iprawsock_posix.go:70
	_go_fuzz_dep_.CoverTab[14780]++
							copy(b, b[l:])
							return n - l
//line /usr/local/go/src/net/iprawsock_posix.go:72
	// _ = "end of CoverTab[14780]"
}

func (c *IPConn) readMsg(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) {
//line /usr/local/go/src/net/iprawsock_posix.go:75
	_go_fuzz_dep_.CoverTab[14788]++
							var sa syscall.Sockaddr
							n, oobn, flags, sa, err = c.fd.readMsg(b, oob, 0)
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/iprawsock_posix.go:79
		_go_fuzz_dep_.CoverTab[14790]++
								addr = &IPAddr{IP: sa.Addr[0:]}
//line /usr/local/go/src/net/iprawsock_posix.go:80
		// _ = "end of CoverTab[14790]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/iprawsock_posix.go:81
		_go_fuzz_dep_.CoverTab[14791]++
								addr = &IPAddr{IP: sa.Addr[0:], Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/iprawsock_posix.go:82
		// _ = "end of CoverTab[14791]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:83
	// _ = "end of CoverTab[14788]"
//line /usr/local/go/src/net/iprawsock_posix.go:83
	_go_fuzz_dep_.CoverTab[14789]++
							return
//line /usr/local/go/src/net/iprawsock_posix.go:84
	// _ = "end of CoverTab[14789]"
}

func (c *IPConn) writeTo(b []byte, addr *IPAddr) (int, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:87
	_go_fuzz_dep_.CoverTab[14792]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/iprawsock_posix.go:88
		_go_fuzz_dep_.CoverTab[14796]++
								return 0, ErrWriteToConnected
//line /usr/local/go/src/net/iprawsock_posix.go:89
		// _ = "end of CoverTab[14796]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:90
		_go_fuzz_dep_.CoverTab[14797]++
//line /usr/local/go/src/net/iprawsock_posix.go:90
		// _ = "end of CoverTab[14797]"
//line /usr/local/go/src/net/iprawsock_posix.go:90
	}
//line /usr/local/go/src/net/iprawsock_posix.go:90
	// _ = "end of CoverTab[14792]"
//line /usr/local/go/src/net/iprawsock_posix.go:90
	_go_fuzz_dep_.CoverTab[14793]++
							if addr == nil {
//line /usr/local/go/src/net/iprawsock_posix.go:91
		_go_fuzz_dep_.CoverTab[14798]++
								return 0, errMissingAddress
//line /usr/local/go/src/net/iprawsock_posix.go:92
		// _ = "end of CoverTab[14798]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:93
		_go_fuzz_dep_.CoverTab[14799]++
//line /usr/local/go/src/net/iprawsock_posix.go:93
		// _ = "end of CoverTab[14799]"
//line /usr/local/go/src/net/iprawsock_posix.go:93
	}
//line /usr/local/go/src/net/iprawsock_posix.go:93
	// _ = "end of CoverTab[14793]"
//line /usr/local/go/src/net/iprawsock_posix.go:93
	_go_fuzz_dep_.CoverTab[14794]++
							sa, err := addr.sockaddr(c.fd.family)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:95
		_go_fuzz_dep_.CoverTab[14800]++
								return 0, err
//line /usr/local/go/src/net/iprawsock_posix.go:96
		// _ = "end of CoverTab[14800]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:97
		_go_fuzz_dep_.CoverTab[14801]++
//line /usr/local/go/src/net/iprawsock_posix.go:97
		// _ = "end of CoverTab[14801]"
//line /usr/local/go/src/net/iprawsock_posix.go:97
	}
//line /usr/local/go/src/net/iprawsock_posix.go:97
	// _ = "end of CoverTab[14794]"
//line /usr/local/go/src/net/iprawsock_posix.go:97
	_go_fuzz_dep_.CoverTab[14795]++
							return c.fd.writeTo(b, sa)
//line /usr/local/go/src/net/iprawsock_posix.go:98
	// _ = "end of CoverTab[14795]"
}

func (c *IPConn) writeMsg(b, oob []byte, addr *IPAddr) (n, oobn int, err error) {
//line /usr/local/go/src/net/iprawsock_posix.go:101
	_go_fuzz_dep_.CoverTab[14802]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/iprawsock_posix.go:102
		_go_fuzz_dep_.CoverTab[14806]++
								return 0, 0, ErrWriteToConnected
//line /usr/local/go/src/net/iprawsock_posix.go:103
		// _ = "end of CoverTab[14806]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:104
		_go_fuzz_dep_.CoverTab[14807]++
//line /usr/local/go/src/net/iprawsock_posix.go:104
		// _ = "end of CoverTab[14807]"
//line /usr/local/go/src/net/iprawsock_posix.go:104
	}
//line /usr/local/go/src/net/iprawsock_posix.go:104
	// _ = "end of CoverTab[14802]"
//line /usr/local/go/src/net/iprawsock_posix.go:104
	_go_fuzz_dep_.CoverTab[14803]++
							if addr == nil {
//line /usr/local/go/src/net/iprawsock_posix.go:105
		_go_fuzz_dep_.CoverTab[14808]++
								return 0, 0, errMissingAddress
//line /usr/local/go/src/net/iprawsock_posix.go:106
		// _ = "end of CoverTab[14808]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:107
		_go_fuzz_dep_.CoverTab[14809]++
//line /usr/local/go/src/net/iprawsock_posix.go:107
		// _ = "end of CoverTab[14809]"
//line /usr/local/go/src/net/iprawsock_posix.go:107
	}
//line /usr/local/go/src/net/iprawsock_posix.go:107
	// _ = "end of CoverTab[14803]"
//line /usr/local/go/src/net/iprawsock_posix.go:107
	_go_fuzz_dep_.CoverTab[14804]++
							sa, err := addr.sockaddr(c.fd.family)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:109
		_go_fuzz_dep_.CoverTab[14810]++
								return 0, 0, err
//line /usr/local/go/src/net/iprawsock_posix.go:110
		// _ = "end of CoverTab[14810]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:111
		_go_fuzz_dep_.CoverTab[14811]++
//line /usr/local/go/src/net/iprawsock_posix.go:111
		// _ = "end of CoverTab[14811]"
//line /usr/local/go/src/net/iprawsock_posix.go:111
	}
//line /usr/local/go/src/net/iprawsock_posix.go:111
	// _ = "end of CoverTab[14804]"
//line /usr/local/go/src/net/iprawsock_posix.go:111
	_go_fuzz_dep_.CoverTab[14805]++
							return c.fd.writeMsg(b, oob, sa)
//line /usr/local/go/src/net/iprawsock_posix.go:112
	// _ = "end of CoverTab[14805]"
}

func (sd *sysDialer) dialIP(ctx context.Context, laddr, raddr *IPAddr) (*IPConn, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:115
	_go_fuzz_dep_.CoverTab[14812]++
							network, proto, err := parseNetwork(ctx, sd.network, true)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:117
		_go_fuzz_dep_.CoverTab[14817]++
								return nil, err
//line /usr/local/go/src/net/iprawsock_posix.go:118
		// _ = "end of CoverTab[14817]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:119
		_go_fuzz_dep_.CoverTab[14818]++
//line /usr/local/go/src/net/iprawsock_posix.go:119
		// _ = "end of CoverTab[14818]"
//line /usr/local/go/src/net/iprawsock_posix.go:119
	}
//line /usr/local/go/src/net/iprawsock_posix.go:119
	// _ = "end of CoverTab[14812]"
//line /usr/local/go/src/net/iprawsock_posix.go:119
	_go_fuzz_dep_.CoverTab[14813]++
							switch network {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/iprawsock_posix.go:121
		_go_fuzz_dep_.CoverTab[14819]++
//line /usr/local/go/src/net/iprawsock_posix.go:121
		// _ = "end of CoverTab[14819]"
	default:
//line /usr/local/go/src/net/iprawsock_posix.go:122
		_go_fuzz_dep_.CoverTab[14820]++
								return nil, UnknownNetworkError(sd.network)
//line /usr/local/go/src/net/iprawsock_posix.go:123
		// _ = "end of CoverTab[14820]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:124
	// _ = "end of CoverTab[14813]"
//line /usr/local/go/src/net/iprawsock_posix.go:124
	_go_fuzz_dep_.CoverTab[14814]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /usr/local/go/src/net/iprawsock_posix.go:126
		_go_fuzz_dep_.CoverTab[14821]++
//line /usr/local/go/src/net/iprawsock_posix.go:126
		return sd.Dialer.Control != nil
//line /usr/local/go/src/net/iprawsock_posix.go:126
		// _ = "end of CoverTab[14821]"
//line /usr/local/go/src/net/iprawsock_posix.go:126
	}() {
//line /usr/local/go/src/net/iprawsock_posix.go:126
		_go_fuzz_dep_.CoverTab[14822]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/iprawsock_posix.go:127
			_go_fuzz_dep_.CoverTab[14823]++
									return sd.Dialer.Control(network, address, c)
//line /usr/local/go/src/net/iprawsock_posix.go:128
			// _ = "end of CoverTab[14823]"
		}
//line /usr/local/go/src/net/iprawsock_posix.go:129
		// _ = "end of CoverTab[14822]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:130
		_go_fuzz_dep_.CoverTab[14824]++
//line /usr/local/go/src/net/iprawsock_posix.go:130
		// _ = "end of CoverTab[14824]"
//line /usr/local/go/src/net/iprawsock_posix.go:130
	}
//line /usr/local/go/src/net/iprawsock_posix.go:130
	// _ = "end of CoverTab[14814]"
//line /usr/local/go/src/net/iprawsock_posix.go:130
	_go_fuzz_dep_.CoverTab[14815]++
							fd, err := internetSocket(ctx, network, laddr, raddr, syscall.SOCK_RAW, proto, "dial", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:132
		_go_fuzz_dep_.CoverTab[14825]++
								return nil, err
//line /usr/local/go/src/net/iprawsock_posix.go:133
		// _ = "end of CoverTab[14825]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:134
		_go_fuzz_dep_.CoverTab[14826]++
//line /usr/local/go/src/net/iprawsock_posix.go:134
		// _ = "end of CoverTab[14826]"
//line /usr/local/go/src/net/iprawsock_posix.go:134
	}
//line /usr/local/go/src/net/iprawsock_posix.go:134
	// _ = "end of CoverTab[14815]"
//line /usr/local/go/src/net/iprawsock_posix.go:134
	_go_fuzz_dep_.CoverTab[14816]++
							return newIPConn(fd), nil
//line /usr/local/go/src/net/iprawsock_posix.go:135
	// _ = "end of CoverTab[14816]"
}

func (sl *sysListener) listenIP(ctx context.Context, laddr *IPAddr) (*IPConn, error) {
//line /usr/local/go/src/net/iprawsock_posix.go:138
	_go_fuzz_dep_.CoverTab[14827]++
							network, proto, err := parseNetwork(ctx, sl.network, true)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:140
		_go_fuzz_dep_.CoverTab[14832]++
								return nil, err
//line /usr/local/go/src/net/iprawsock_posix.go:141
		// _ = "end of CoverTab[14832]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:142
		_go_fuzz_dep_.CoverTab[14833]++
//line /usr/local/go/src/net/iprawsock_posix.go:142
		// _ = "end of CoverTab[14833]"
//line /usr/local/go/src/net/iprawsock_posix.go:142
	}
//line /usr/local/go/src/net/iprawsock_posix.go:142
	// _ = "end of CoverTab[14827]"
//line /usr/local/go/src/net/iprawsock_posix.go:142
	_go_fuzz_dep_.CoverTab[14828]++
							switch network {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/iprawsock_posix.go:144
		_go_fuzz_dep_.CoverTab[14834]++
//line /usr/local/go/src/net/iprawsock_posix.go:144
		// _ = "end of CoverTab[14834]"
	default:
//line /usr/local/go/src/net/iprawsock_posix.go:145
		_go_fuzz_dep_.CoverTab[14835]++
								return nil, UnknownNetworkError(sl.network)
//line /usr/local/go/src/net/iprawsock_posix.go:146
		// _ = "end of CoverTab[14835]"
	}
//line /usr/local/go/src/net/iprawsock_posix.go:147
	// _ = "end of CoverTab[14828]"
//line /usr/local/go/src/net/iprawsock_posix.go:147
	_go_fuzz_dep_.CoverTab[14829]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:149
		_go_fuzz_dep_.CoverTab[14836]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/iprawsock_posix.go:150
			_go_fuzz_dep_.CoverTab[14837]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/iprawsock_posix.go:151
			// _ = "end of CoverTab[14837]"
		}
//line /usr/local/go/src/net/iprawsock_posix.go:152
		// _ = "end of CoverTab[14836]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:153
		_go_fuzz_dep_.CoverTab[14838]++
//line /usr/local/go/src/net/iprawsock_posix.go:153
		// _ = "end of CoverTab[14838]"
//line /usr/local/go/src/net/iprawsock_posix.go:153
	}
//line /usr/local/go/src/net/iprawsock_posix.go:153
	// _ = "end of CoverTab[14829]"
//line /usr/local/go/src/net/iprawsock_posix.go:153
	_go_fuzz_dep_.CoverTab[14830]++
							fd, err := internetSocket(ctx, network, laddr, nil, syscall.SOCK_RAW, proto, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/iprawsock_posix.go:155
		_go_fuzz_dep_.CoverTab[14839]++
								return nil, err
//line /usr/local/go/src/net/iprawsock_posix.go:156
		// _ = "end of CoverTab[14839]"
	} else {
//line /usr/local/go/src/net/iprawsock_posix.go:157
		_go_fuzz_dep_.CoverTab[14840]++
//line /usr/local/go/src/net/iprawsock_posix.go:157
		// _ = "end of CoverTab[14840]"
//line /usr/local/go/src/net/iprawsock_posix.go:157
	}
//line /usr/local/go/src/net/iprawsock_posix.go:157
	// _ = "end of CoverTab[14830]"
//line /usr/local/go/src/net/iprawsock_posix.go:157
	_go_fuzz_dep_.CoverTab[14831]++
							return newIPConn(fd), nil
//line /usr/local/go/src/net/iprawsock_posix.go:158
	// _ = "end of CoverTab[14831]"
}

//line /usr/local/go/src/net/iprawsock_posix.go:159
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/iprawsock_posix.go:159
var _ = _go_fuzz_dep_.CoverTab
