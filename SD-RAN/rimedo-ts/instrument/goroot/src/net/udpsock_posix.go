// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || windows

//line /usr/local/go/src/net/udpsock_posix.go:7
package net

//line /usr/local/go/src/net/udpsock_posix.go:7
import (
//line /usr/local/go/src/net/udpsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/udpsock_posix.go:7
)
//line /usr/local/go/src/net/udpsock_posix.go:7
import (
//line /usr/local/go/src/net/udpsock_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/udpsock_posix.go:7
)

import (
	"context"
	"net/netip"
	"syscall"
)

func sockaddrToUDP(sa syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/udpsock_posix.go:15
	_go_fuzz_dep_.CoverTab[16769]++
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/udpsock_posix.go:17
		_go_fuzz_dep_.CoverTab[16771]++
								return &UDPAddr{IP: sa.Addr[0:], Port: sa.Port}
//line /usr/local/go/src/net/udpsock_posix.go:18
		// _ = "end of CoverTab[16771]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/udpsock_posix.go:19
		_go_fuzz_dep_.CoverTab[16772]++
								return &UDPAddr{IP: sa.Addr[0:], Port: sa.Port, Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/udpsock_posix.go:20
		// _ = "end of CoverTab[16772]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:21
	// _ = "end of CoverTab[16769]"
//line /usr/local/go/src/net/udpsock_posix.go:21
	_go_fuzz_dep_.CoverTab[16770]++
							return nil
//line /usr/local/go/src/net/udpsock_posix.go:22
	// _ = "end of CoverTab[16770]"
}

func (a *UDPAddr) family() int {
//line /usr/local/go/src/net/udpsock_posix.go:25
	_go_fuzz_dep_.CoverTab[16773]++
							if a == nil || func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:26
		_go_fuzz_dep_.CoverTab[16776]++
//line /usr/local/go/src/net/udpsock_posix.go:26
		return len(a.IP) <= IPv4len
//line /usr/local/go/src/net/udpsock_posix.go:26
		// _ = "end of CoverTab[16776]"
//line /usr/local/go/src/net/udpsock_posix.go:26
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:26
		_go_fuzz_dep_.CoverTab[16777]++
								return syscall.AF_INET
//line /usr/local/go/src/net/udpsock_posix.go:27
		// _ = "end of CoverTab[16777]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:28
		_go_fuzz_dep_.CoverTab[16778]++
//line /usr/local/go/src/net/udpsock_posix.go:28
		// _ = "end of CoverTab[16778]"
//line /usr/local/go/src/net/udpsock_posix.go:28
	}
//line /usr/local/go/src/net/udpsock_posix.go:28
	// _ = "end of CoverTab[16773]"
//line /usr/local/go/src/net/udpsock_posix.go:28
	_go_fuzz_dep_.CoverTab[16774]++
							if a.IP.To4() != nil {
//line /usr/local/go/src/net/udpsock_posix.go:29
		_go_fuzz_dep_.CoverTab[16779]++
								return syscall.AF_INET
//line /usr/local/go/src/net/udpsock_posix.go:30
		// _ = "end of CoverTab[16779]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:31
		_go_fuzz_dep_.CoverTab[16780]++
//line /usr/local/go/src/net/udpsock_posix.go:31
		// _ = "end of CoverTab[16780]"
//line /usr/local/go/src/net/udpsock_posix.go:31
	}
//line /usr/local/go/src/net/udpsock_posix.go:31
	// _ = "end of CoverTab[16774]"
//line /usr/local/go/src/net/udpsock_posix.go:31
	_go_fuzz_dep_.CoverTab[16775]++
							return syscall.AF_INET6
//line /usr/local/go/src/net/udpsock_posix.go:32
	// _ = "end of CoverTab[16775]"
}

func (a *UDPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/udpsock_posix.go:35
	_go_fuzz_dep_.CoverTab[16781]++
							if a == nil {
//line /usr/local/go/src/net/udpsock_posix.go:36
		_go_fuzz_dep_.CoverTab[16783]++
								return nil, nil
//line /usr/local/go/src/net/udpsock_posix.go:37
		// _ = "end of CoverTab[16783]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:38
		_go_fuzz_dep_.CoverTab[16784]++
//line /usr/local/go/src/net/udpsock_posix.go:38
		// _ = "end of CoverTab[16784]"
//line /usr/local/go/src/net/udpsock_posix.go:38
	}
//line /usr/local/go/src/net/udpsock_posix.go:38
	// _ = "end of CoverTab[16781]"
//line /usr/local/go/src/net/udpsock_posix.go:38
	_go_fuzz_dep_.CoverTab[16782]++
							return ipToSockaddr(family, a.IP, a.Port, a.Zone)
//line /usr/local/go/src/net/udpsock_posix.go:39
	// _ = "end of CoverTab[16782]"
}

func (a *UDPAddr) toLocal(net string) sockaddr {
//line /usr/local/go/src/net/udpsock_posix.go:42
	_go_fuzz_dep_.CoverTab[16785]++
							return &UDPAddr{loopbackIP(net), a.Port, a.Zone}
//line /usr/local/go/src/net/udpsock_posix.go:43
	// _ = "end of CoverTab[16785]"
}

func (c *UDPConn) readFrom(b []byte, addr *UDPAddr) (int, *UDPAddr, error) {
//line /usr/local/go/src/net/udpsock_posix.go:46
	_go_fuzz_dep_.CoverTab[16786]++
							var n int
							var err error
							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:50
		_go_fuzz_dep_.CoverTab[16789]++
								var from syscall.SockaddrInet4
								n, err = c.fd.readFromInet4(b, &from)
								if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:53
			_go_fuzz_dep_.CoverTab[16792]++
									ip := from.Addr
									*addr = UDPAddr{IP: ip[:], Port: from.Port}
//line /usr/local/go/src/net/udpsock_posix.go:55
			// _ = "end of CoverTab[16792]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:56
			_go_fuzz_dep_.CoverTab[16793]++
//line /usr/local/go/src/net/udpsock_posix.go:56
			// _ = "end of CoverTab[16793]"
//line /usr/local/go/src/net/udpsock_posix.go:56
		}
//line /usr/local/go/src/net/udpsock_posix.go:56
		// _ = "end of CoverTab[16789]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:57
		_go_fuzz_dep_.CoverTab[16790]++
								var from syscall.SockaddrInet6
								n, err = c.fd.readFromInet6(b, &from)
								if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:60
			_go_fuzz_dep_.CoverTab[16794]++
									ip := from.Addr
									*addr = UDPAddr{IP: ip[:], Port: from.Port, Zone: zoneCache.name(int(from.ZoneId))}
//line /usr/local/go/src/net/udpsock_posix.go:62
			// _ = "end of CoverTab[16794]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:63
			_go_fuzz_dep_.CoverTab[16795]++
//line /usr/local/go/src/net/udpsock_posix.go:63
			// _ = "end of CoverTab[16795]"
//line /usr/local/go/src/net/udpsock_posix.go:63
		}
//line /usr/local/go/src/net/udpsock_posix.go:63
		// _ = "end of CoverTab[16790]"
//line /usr/local/go/src/net/udpsock_posix.go:63
	default:
//line /usr/local/go/src/net/udpsock_posix.go:63
		_go_fuzz_dep_.CoverTab[16791]++
//line /usr/local/go/src/net/udpsock_posix.go:63
		// _ = "end of CoverTab[16791]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:64
	// _ = "end of CoverTab[16786]"
//line /usr/local/go/src/net/udpsock_posix.go:64
	_go_fuzz_dep_.CoverTab[16787]++
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:65
		_go_fuzz_dep_.CoverTab[16796]++

								addr = nil
//line /usr/local/go/src/net/udpsock_posix.go:67
		// _ = "end of CoverTab[16796]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:68
		_go_fuzz_dep_.CoverTab[16797]++
//line /usr/local/go/src/net/udpsock_posix.go:68
		// _ = "end of CoverTab[16797]"
//line /usr/local/go/src/net/udpsock_posix.go:68
	}
//line /usr/local/go/src/net/udpsock_posix.go:68
	// _ = "end of CoverTab[16787]"
//line /usr/local/go/src/net/udpsock_posix.go:68
	_go_fuzz_dep_.CoverTab[16788]++
							return n, addr, err
//line /usr/local/go/src/net/udpsock_posix.go:69
	// _ = "end of CoverTab[16788]"
}

func (c *UDPConn) readFromAddrPort(b []byte) (n int, addr netip.AddrPort, err error) {
//line /usr/local/go/src/net/udpsock_posix.go:72
	_go_fuzz_dep_.CoverTab[16798]++
							var ip netip.Addr
							var port int
							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:76
		_go_fuzz_dep_.CoverTab[16801]++
								var from syscall.SockaddrInet4
								n, err = c.fd.readFromInet4(b, &from)
								if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:79
			_go_fuzz_dep_.CoverTab[16804]++
									ip = netip.AddrFrom4(from.Addr)
									port = from.Port
//line /usr/local/go/src/net/udpsock_posix.go:81
			// _ = "end of CoverTab[16804]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:82
			_go_fuzz_dep_.CoverTab[16805]++
//line /usr/local/go/src/net/udpsock_posix.go:82
			// _ = "end of CoverTab[16805]"
//line /usr/local/go/src/net/udpsock_posix.go:82
		}
//line /usr/local/go/src/net/udpsock_posix.go:82
		// _ = "end of CoverTab[16801]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:83
		_go_fuzz_dep_.CoverTab[16802]++
								var from syscall.SockaddrInet6
								n, err = c.fd.readFromInet6(b, &from)
								if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:86
			_go_fuzz_dep_.CoverTab[16806]++
									ip = netip.AddrFrom16(from.Addr).WithZone(zoneCache.name(int(from.ZoneId)))
									port = from.Port
//line /usr/local/go/src/net/udpsock_posix.go:88
			// _ = "end of CoverTab[16806]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:89
			_go_fuzz_dep_.CoverTab[16807]++
//line /usr/local/go/src/net/udpsock_posix.go:89
			// _ = "end of CoverTab[16807]"
//line /usr/local/go/src/net/udpsock_posix.go:89
		}
//line /usr/local/go/src/net/udpsock_posix.go:89
		// _ = "end of CoverTab[16802]"
//line /usr/local/go/src/net/udpsock_posix.go:89
	default:
//line /usr/local/go/src/net/udpsock_posix.go:89
		_go_fuzz_dep_.CoverTab[16803]++
//line /usr/local/go/src/net/udpsock_posix.go:89
		// _ = "end of CoverTab[16803]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:90
	// _ = "end of CoverTab[16798]"
//line /usr/local/go/src/net/udpsock_posix.go:90
	_go_fuzz_dep_.CoverTab[16799]++
							if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:91
		_go_fuzz_dep_.CoverTab[16808]++
								addr = netip.AddrPortFrom(ip, uint16(port))
//line /usr/local/go/src/net/udpsock_posix.go:92
		// _ = "end of CoverTab[16808]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:93
		_go_fuzz_dep_.CoverTab[16809]++
//line /usr/local/go/src/net/udpsock_posix.go:93
		// _ = "end of CoverTab[16809]"
//line /usr/local/go/src/net/udpsock_posix.go:93
	}
//line /usr/local/go/src/net/udpsock_posix.go:93
	// _ = "end of CoverTab[16799]"
//line /usr/local/go/src/net/udpsock_posix.go:93
	_go_fuzz_dep_.CoverTab[16800]++
							return n, addr, err
//line /usr/local/go/src/net/udpsock_posix.go:94
	// _ = "end of CoverTab[16800]"
}

func (c *UDPConn) readMsg(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error) {
//line /usr/local/go/src/net/udpsock_posix.go:97
	_go_fuzz_dep_.CoverTab[16810]++
							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:99
		_go_fuzz_dep_.CoverTab[16812]++
								var sa syscall.SockaddrInet4
								n, oobn, flags, err = c.fd.readMsgInet4(b, oob, 0, &sa)
								ip := netip.AddrFrom4(sa.Addr)
								addr = netip.AddrPortFrom(ip, uint16(sa.Port))
//line /usr/local/go/src/net/udpsock_posix.go:103
		// _ = "end of CoverTab[16812]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:104
		_go_fuzz_dep_.CoverTab[16813]++
								var sa syscall.SockaddrInet6
								n, oobn, flags, err = c.fd.readMsgInet6(b, oob, 0, &sa)
								ip := netip.AddrFrom16(sa.Addr).WithZone(zoneCache.name(int(sa.ZoneId)))
								addr = netip.AddrPortFrom(ip, uint16(sa.Port))
//line /usr/local/go/src/net/udpsock_posix.go:108
		// _ = "end of CoverTab[16813]"
//line /usr/local/go/src/net/udpsock_posix.go:108
	default:
//line /usr/local/go/src/net/udpsock_posix.go:108
		_go_fuzz_dep_.CoverTab[16814]++
//line /usr/local/go/src/net/udpsock_posix.go:108
		// _ = "end of CoverTab[16814]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:109
	// _ = "end of CoverTab[16810]"
//line /usr/local/go/src/net/udpsock_posix.go:109
	_go_fuzz_dep_.CoverTab[16811]++
							return
//line /usr/local/go/src/net/udpsock_posix.go:110
	// _ = "end of CoverTab[16811]"
}

func (c *UDPConn) writeTo(b []byte, addr *UDPAddr) (int, error) {
//line /usr/local/go/src/net/udpsock_posix.go:113
	_go_fuzz_dep_.CoverTab[16815]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/udpsock_posix.go:114
		_go_fuzz_dep_.CoverTab[16818]++
								return 0, ErrWriteToConnected
//line /usr/local/go/src/net/udpsock_posix.go:115
		// _ = "end of CoverTab[16818]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:116
		_go_fuzz_dep_.CoverTab[16819]++
//line /usr/local/go/src/net/udpsock_posix.go:116
		// _ = "end of CoverTab[16819]"
//line /usr/local/go/src/net/udpsock_posix.go:116
	}
//line /usr/local/go/src/net/udpsock_posix.go:116
	// _ = "end of CoverTab[16815]"
//line /usr/local/go/src/net/udpsock_posix.go:116
	_go_fuzz_dep_.CoverTab[16816]++
							if addr == nil {
//line /usr/local/go/src/net/udpsock_posix.go:117
		_go_fuzz_dep_.CoverTab[16820]++
								return 0, errMissingAddress
//line /usr/local/go/src/net/udpsock_posix.go:118
		// _ = "end of CoverTab[16820]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:119
		_go_fuzz_dep_.CoverTab[16821]++
//line /usr/local/go/src/net/udpsock_posix.go:119
		// _ = "end of CoverTab[16821]"
//line /usr/local/go/src/net/udpsock_posix.go:119
	}
//line /usr/local/go/src/net/udpsock_posix.go:119
	// _ = "end of CoverTab[16816]"
//line /usr/local/go/src/net/udpsock_posix.go:119
	_go_fuzz_dep_.CoverTab[16817]++

							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:122
		_go_fuzz_dep_.CoverTab[16822]++
								sa, err := ipToSockaddrInet4(addr.IP, addr.Port)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:124
			_go_fuzz_dep_.CoverTab[16827]++
									return 0, err
//line /usr/local/go/src/net/udpsock_posix.go:125
			// _ = "end of CoverTab[16827]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:126
			_go_fuzz_dep_.CoverTab[16828]++
//line /usr/local/go/src/net/udpsock_posix.go:126
			// _ = "end of CoverTab[16828]"
//line /usr/local/go/src/net/udpsock_posix.go:126
		}
//line /usr/local/go/src/net/udpsock_posix.go:126
		// _ = "end of CoverTab[16822]"
//line /usr/local/go/src/net/udpsock_posix.go:126
		_go_fuzz_dep_.CoverTab[16823]++
								return c.fd.writeToInet4(b, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:127
		// _ = "end of CoverTab[16823]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:128
		_go_fuzz_dep_.CoverTab[16824]++
								sa, err := ipToSockaddrInet6(addr.IP, addr.Port, addr.Zone)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:130
			_go_fuzz_dep_.CoverTab[16829]++
									return 0, err
//line /usr/local/go/src/net/udpsock_posix.go:131
			// _ = "end of CoverTab[16829]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:132
			_go_fuzz_dep_.CoverTab[16830]++
//line /usr/local/go/src/net/udpsock_posix.go:132
			// _ = "end of CoverTab[16830]"
//line /usr/local/go/src/net/udpsock_posix.go:132
		}
//line /usr/local/go/src/net/udpsock_posix.go:132
		// _ = "end of CoverTab[16824]"
//line /usr/local/go/src/net/udpsock_posix.go:132
		_go_fuzz_dep_.CoverTab[16825]++
								return c.fd.writeToInet6(b, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:133
		// _ = "end of CoverTab[16825]"
	default:
//line /usr/local/go/src/net/udpsock_posix.go:134
		_go_fuzz_dep_.CoverTab[16826]++
								return 0, &AddrError{Err: "invalid address family", Addr: addr.IP.String()}
//line /usr/local/go/src/net/udpsock_posix.go:135
		// _ = "end of CoverTab[16826]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:136
	// _ = "end of CoverTab[16817]"
}

func (c *UDPConn) writeToAddrPort(b []byte, addr netip.AddrPort) (int, error) {
//line /usr/local/go/src/net/udpsock_posix.go:139
	_go_fuzz_dep_.CoverTab[16831]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/udpsock_posix.go:140
		_go_fuzz_dep_.CoverTab[16834]++
								return 0, ErrWriteToConnected
//line /usr/local/go/src/net/udpsock_posix.go:141
		// _ = "end of CoverTab[16834]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:142
		_go_fuzz_dep_.CoverTab[16835]++
//line /usr/local/go/src/net/udpsock_posix.go:142
		// _ = "end of CoverTab[16835]"
//line /usr/local/go/src/net/udpsock_posix.go:142
	}
//line /usr/local/go/src/net/udpsock_posix.go:142
	// _ = "end of CoverTab[16831]"
//line /usr/local/go/src/net/udpsock_posix.go:142
	_go_fuzz_dep_.CoverTab[16832]++
							if !addr.IsValid() {
//line /usr/local/go/src/net/udpsock_posix.go:143
		_go_fuzz_dep_.CoverTab[16836]++
								return 0, errMissingAddress
//line /usr/local/go/src/net/udpsock_posix.go:144
		// _ = "end of CoverTab[16836]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:145
		_go_fuzz_dep_.CoverTab[16837]++
//line /usr/local/go/src/net/udpsock_posix.go:145
		// _ = "end of CoverTab[16837]"
//line /usr/local/go/src/net/udpsock_posix.go:145
	}
//line /usr/local/go/src/net/udpsock_posix.go:145
	// _ = "end of CoverTab[16832]"
//line /usr/local/go/src/net/udpsock_posix.go:145
	_go_fuzz_dep_.CoverTab[16833]++

							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:148
		_go_fuzz_dep_.CoverTab[16838]++
								sa, err := addrPortToSockaddrInet4(addr)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:150
			_go_fuzz_dep_.CoverTab[16843]++
									return 0, err
//line /usr/local/go/src/net/udpsock_posix.go:151
			// _ = "end of CoverTab[16843]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:152
			_go_fuzz_dep_.CoverTab[16844]++
//line /usr/local/go/src/net/udpsock_posix.go:152
			// _ = "end of CoverTab[16844]"
//line /usr/local/go/src/net/udpsock_posix.go:152
		}
//line /usr/local/go/src/net/udpsock_posix.go:152
		// _ = "end of CoverTab[16838]"
//line /usr/local/go/src/net/udpsock_posix.go:152
		_go_fuzz_dep_.CoverTab[16839]++
								return c.fd.writeToInet4(b, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:153
		// _ = "end of CoverTab[16839]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:154
		_go_fuzz_dep_.CoverTab[16840]++
								sa, err := addrPortToSockaddrInet6(addr)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:156
			_go_fuzz_dep_.CoverTab[16845]++
									return 0, err
//line /usr/local/go/src/net/udpsock_posix.go:157
			// _ = "end of CoverTab[16845]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:158
			_go_fuzz_dep_.CoverTab[16846]++
//line /usr/local/go/src/net/udpsock_posix.go:158
			// _ = "end of CoverTab[16846]"
//line /usr/local/go/src/net/udpsock_posix.go:158
		}
//line /usr/local/go/src/net/udpsock_posix.go:158
		// _ = "end of CoverTab[16840]"
//line /usr/local/go/src/net/udpsock_posix.go:158
		_go_fuzz_dep_.CoverTab[16841]++
								return c.fd.writeToInet6(b, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:159
		// _ = "end of CoverTab[16841]"
	default:
//line /usr/local/go/src/net/udpsock_posix.go:160
		_go_fuzz_dep_.CoverTab[16842]++
								return 0, &AddrError{Err: "invalid address family", Addr: addr.Addr().String()}
//line /usr/local/go/src/net/udpsock_posix.go:161
		// _ = "end of CoverTab[16842]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:162
	// _ = "end of CoverTab[16833]"
}

func (c *UDPConn) writeMsg(b, oob []byte, addr *UDPAddr) (n, oobn int, err error) {
//line /usr/local/go/src/net/udpsock_posix.go:165
	_go_fuzz_dep_.CoverTab[16847]++
							if c.fd.isConnected && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:166
		_go_fuzz_dep_.CoverTab[16851]++
//line /usr/local/go/src/net/udpsock_posix.go:166
		return addr != nil
//line /usr/local/go/src/net/udpsock_posix.go:166
		// _ = "end of CoverTab[16851]"
//line /usr/local/go/src/net/udpsock_posix.go:166
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:166
		_go_fuzz_dep_.CoverTab[16852]++
								return 0, 0, ErrWriteToConnected
//line /usr/local/go/src/net/udpsock_posix.go:167
		// _ = "end of CoverTab[16852]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:168
		_go_fuzz_dep_.CoverTab[16853]++
//line /usr/local/go/src/net/udpsock_posix.go:168
		// _ = "end of CoverTab[16853]"
//line /usr/local/go/src/net/udpsock_posix.go:168
	}
//line /usr/local/go/src/net/udpsock_posix.go:168
	// _ = "end of CoverTab[16847]"
//line /usr/local/go/src/net/udpsock_posix.go:168
	_go_fuzz_dep_.CoverTab[16848]++
							if !c.fd.isConnected && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:169
		_go_fuzz_dep_.CoverTab[16854]++
//line /usr/local/go/src/net/udpsock_posix.go:169
		return addr == nil
//line /usr/local/go/src/net/udpsock_posix.go:169
		// _ = "end of CoverTab[16854]"
//line /usr/local/go/src/net/udpsock_posix.go:169
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:169
		_go_fuzz_dep_.CoverTab[16855]++
								return 0, 0, errMissingAddress
//line /usr/local/go/src/net/udpsock_posix.go:170
		// _ = "end of CoverTab[16855]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:171
		_go_fuzz_dep_.CoverTab[16856]++
//line /usr/local/go/src/net/udpsock_posix.go:171
		// _ = "end of CoverTab[16856]"
//line /usr/local/go/src/net/udpsock_posix.go:171
	}
//line /usr/local/go/src/net/udpsock_posix.go:171
	// _ = "end of CoverTab[16848]"
//line /usr/local/go/src/net/udpsock_posix.go:171
	_go_fuzz_dep_.CoverTab[16849]++
							sa, err := addr.sockaddr(c.fd.family)
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:173
		_go_fuzz_dep_.CoverTab[16857]++
								return 0, 0, err
//line /usr/local/go/src/net/udpsock_posix.go:174
		// _ = "end of CoverTab[16857]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:175
		_go_fuzz_dep_.CoverTab[16858]++
//line /usr/local/go/src/net/udpsock_posix.go:175
		// _ = "end of CoverTab[16858]"
//line /usr/local/go/src/net/udpsock_posix.go:175
	}
//line /usr/local/go/src/net/udpsock_posix.go:175
	// _ = "end of CoverTab[16849]"
//line /usr/local/go/src/net/udpsock_posix.go:175
	_go_fuzz_dep_.CoverTab[16850]++
							return c.fd.writeMsg(b, oob, sa)
//line /usr/local/go/src/net/udpsock_posix.go:176
	// _ = "end of CoverTab[16850]"
}

func (c *UDPConn) writeMsgAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error) {
//line /usr/local/go/src/net/udpsock_posix.go:179
	_go_fuzz_dep_.CoverTab[16859]++
							if c.fd.isConnected && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:180
		_go_fuzz_dep_.CoverTab[16862]++
//line /usr/local/go/src/net/udpsock_posix.go:180
		return addr.IsValid()
//line /usr/local/go/src/net/udpsock_posix.go:180
		// _ = "end of CoverTab[16862]"
//line /usr/local/go/src/net/udpsock_posix.go:180
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:180
		_go_fuzz_dep_.CoverTab[16863]++
								return 0, 0, ErrWriteToConnected
//line /usr/local/go/src/net/udpsock_posix.go:181
		// _ = "end of CoverTab[16863]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:182
		_go_fuzz_dep_.CoverTab[16864]++
//line /usr/local/go/src/net/udpsock_posix.go:182
		// _ = "end of CoverTab[16864]"
//line /usr/local/go/src/net/udpsock_posix.go:182
	}
//line /usr/local/go/src/net/udpsock_posix.go:182
	// _ = "end of CoverTab[16859]"
//line /usr/local/go/src/net/udpsock_posix.go:182
	_go_fuzz_dep_.CoverTab[16860]++
							if !c.fd.isConnected && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:183
		_go_fuzz_dep_.CoverTab[16865]++
//line /usr/local/go/src/net/udpsock_posix.go:183
		return !addr.IsValid()
//line /usr/local/go/src/net/udpsock_posix.go:183
		// _ = "end of CoverTab[16865]"
//line /usr/local/go/src/net/udpsock_posix.go:183
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:183
		_go_fuzz_dep_.CoverTab[16866]++
								return 0, 0, errMissingAddress
//line /usr/local/go/src/net/udpsock_posix.go:184
		// _ = "end of CoverTab[16866]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:185
		_go_fuzz_dep_.CoverTab[16867]++
//line /usr/local/go/src/net/udpsock_posix.go:185
		// _ = "end of CoverTab[16867]"
//line /usr/local/go/src/net/udpsock_posix.go:185
	}
//line /usr/local/go/src/net/udpsock_posix.go:185
	// _ = "end of CoverTab[16860]"
//line /usr/local/go/src/net/udpsock_posix.go:185
	_go_fuzz_dep_.CoverTab[16861]++

							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:188
		_go_fuzz_dep_.CoverTab[16868]++
								sa, err := addrPortToSockaddrInet4(addr)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:190
			_go_fuzz_dep_.CoverTab[16873]++
									return 0, 0, err
//line /usr/local/go/src/net/udpsock_posix.go:191
			// _ = "end of CoverTab[16873]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:192
			_go_fuzz_dep_.CoverTab[16874]++
//line /usr/local/go/src/net/udpsock_posix.go:192
			// _ = "end of CoverTab[16874]"
//line /usr/local/go/src/net/udpsock_posix.go:192
		}
//line /usr/local/go/src/net/udpsock_posix.go:192
		// _ = "end of CoverTab[16868]"
//line /usr/local/go/src/net/udpsock_posix.go:192
		_go_fuzz_dep_.CoverTab[16869]++
								return c.fd.writeMsgInet4(b, oob, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:193
		// _ = "end of CoverTab[16869]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:194
		_go_fuzz_dep_.CoverTab[16870]++
								sa, err := addrPortToSockaddrInet6(addr)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:196
			_go_fuzz_dep_.CoverTab[16875]++
									return 0, 0, err
//line /usr/local/go/src/net/udpsock_posix.go:197
			// _ = "end of CoverTab[16875]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:198
			_go_fuzz_dep_.CoverTab[16876]++
//line /usr/local/go/src/net/udpsock_posix.go:198
			// _ = "end of CoverTab[16876]"
//line /usr/local/go/src/net/udpsock_posix.go:198
		}
//line /usr/local/go/src/net/udpsock_posix.go:198
		// _ = "end of CoverTab[16870]"
//line /usr/local/go/src/net/udpsock_posix.go:198
		_go_fuzz_dep_.CoverTab[16871]++
								return c.fd.writeMsgInet6(b, oob, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:199
		// _ = "end of CoverTab[16871]"
	default:
//line /usr/local/go/src/net/udpsock_posix.go:200
		_go_fuzz_dep_.CoverTab[16872]++
								return 0, 0, &AddrError{Err: "invalid address family", Addr: addr.Addr().String()}
//line /usr/local/go/src/net/udpsock_posix.go:201
		// _ = "end of CoverTab[16872]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:202
	// _ = "end of CoverTab[16861]"
}

func (sd *sysDialer) dialUDP(ctx context.Context, laddr, raddr *UDPAddr) (*UDPConn, error) {
//line /usr/local/go/src/net/udpsock_posix.go:205
	_go_fuzz_dep_.CoverTab[16877]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:207
		_go_fuzz_dep_.CoverTab[16880]++
//line /usr/local/go/src/net/udpsock_posix.go:207
		return sd.Dialer.Control != nil
//line /usr/local/go/src/net/udpsock_posix.go:207
		// _ = "end of CoverTab[16880]"
//line /usr/local/go/src/net/udpsock_posix.go:207
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:207
		_go_fuzz_dep_.CoverTab[16881]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/udpsock_posix.go:208
			_go_fuzz_dep_.CoverTab[16882]++
									return sd.Dialer.Control(network, address, c)
//line /usr/local/go/src/net/udpsock_posix.go:209
			// _ = "end of CoverTab[16882]"
		}
//line /usr/local/go/src/net/udpsock_posix.go:210
		// _ = "end of CoverTab[16881]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:211
		_go_fuzz_dep_.CoverTab[16883]++
//line /usr/local/go/src/net/udpsock_posix.go:211
		// _ = "end of CoverTab[16883]"
//line /usr/local/go/src/net/udpsock_posix.go:211
	}
//line /usr/local/go/src/net/udpsock_posix.go:211
	// _ = "end of CoverTab[16877]"
//line /usr/local/go/src/net/udpsock_posix.go:211
	_go_fuzz_dep_.CoverTab[16878]++
							fd, err := internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_DGRAM, 0, "dial", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:213
		_go_fuzz_dep_.CoverTab[16884]++
								return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:214
		// _ = "end of CoverTab[16884]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:215
		_go_fuzz_dep_.CoverTab[16885]++
//line /usr/local/go/src/net/udpsock_posix.go:215
		// _ = "end of CoverTab[16885]"
//line /usr/local/go/src/net/udpsock_posix.go:215
	}
//line /usr/local/go/src/net/udpsock_posix.go:215
	// _ = "end of CoverTab[16878]"
//line /usr/local/go/src/net/udpsock_posix.go:215
	_go_fuzz_dep_.CoverTab[16879]++
							return newUDPConn(fd), nil
//line /usr/local/go/src/net/udpsock_posix.go:216
	// _ = "end of CoverTab[16879]"
}

func (sl *sysListener) listenUDP(ctx context.Context, laddr *UDPAddr) (*UDPConn, error) {
//line /usr/local/go/src/net/udpsock_posix.go:219
	_go_fuzz_dep_.CoverTab[16886]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/udpsock_posix.go:221
		_go_fuzz_dep_.CoverTab[16889]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/udpsock_posix.go:222
			_go_fuzz_dep_.CoverTab[16890]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/udpsock_posix.go:223
			// _ = "end of CoverTab[16890]"
		}
//line /usr/local/go/src/net/udpsock_posix.go:224
		// _ = "end of CoverTab[16889]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:225
		_go_fuzz_dep_.CoverTab[16891]++
//line /usr/local/go/src/net/udpsock_posix.go:225
		// _ = "end of CoverTab[16891]"
//line /usr/local/go/src/net/udpsock_posix.go:225
	}
//line /usr/local/go/src/net/udpsock_posix.go:225
	// _ = "end of CoverTab[16886]"
//line /usr/local/go/src/net/udpsock_posix.go:225
	_go_fuzz_dep_.CoverTab[16887]++
							fd, err := internetSocket(ctx, sl.network, laddr, nil, syscall.SOCK_DGRAM, 0, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:227
		_go_fuzz_dep_.CoverTab[16892]++
								return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:228
		// _ = "end of CoverTab[16892]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:229
		_go_fuzz_dep_.CoverTab[16893]++
//line /usr/local/go/src/net/udpsock_posix.go:229
		// _ = "end of CoverTab[16893]"
//line /usr/local/go/src/net/udpsock_posix.go:229
	}
//line /usr/local/go/src/net/udpsock_posix.go:229
	// _ = "end of CoverTab[16887]"
//line /usr/local/go/src/net/udpsock_posix.go:229
	_go_fuzz_dep_.CoverTab[16888]++
							return newUDPConn(fd), nil
//line /usr/local/go/src/net/udpsock_posix.go:230
	// _ = "end of CoverTab[16888]"
}

func (sl *sysListener) listenMulticastUDP(ctx context.Context, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error) {
//line /usr/local/go/src/net/udpsock_posix.go:233
	_go_fuzz_dep_.CoverTab[16894]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/udpsock_posix.go:235
		_go_fuzz_dep_.CoverTab[16898]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/udpsock_posix.go:236
			_go_fuzz_dep_.CoverTab[16899]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/udpsock_posix.go:237
			// _ = "end of CoverTab[16899]"
		}
//line /usr/local/go/src/net/udpsock_posix.go:238
		// _ = "end of CoverTab[16898]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:239
		_go_fuzz_dep_.CoverTab[16900]++
//line /usr/local/go/src/net/udpsock_posix.go:239
		// _ = "end of CoverTab[16900]"
//line /usr/local/go/src/net/udpsock_posix.go:239
	}
//line /usr/local/go/src/net/udpsock_posix.go:239
	// _ = "end of CoverTab[16894]"
//line /usr/local/go/src/net/udpsock_posix.go:239
	_go_fuzz_dep_.CoverTab[16895]++
							fd, err := internetSocket(ctx, sl.network, gaddr, nil, syscall.SOCK_DGRAM, 0, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:241
		_go_fuzz_dep_.CoverTab[16901]++
								return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:242
		// _ = "end of CoverTab[16901]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:243
		_go_fuzz_dep_.CoverTab[16902]++
//line /usr/local/go/src/net/udpsock_posix.go:243
		// _ = "end of CoverTab[16902]"
//line /usr/local/go/src/net/udpsock_posix.go:243
	}
//line /usr/local/go/src/net/udpsock_posix.go:243
	// _ = "end of CoverTab[16895]"
//line /usr/local/go/src/net/udpsock_posix.go:243
	_go_fuzz_dep_.CoverTab[16896]++
							c := newUDPConn(fd)
							if ip4 := gaddr.IP.To4(); ip4 != nil {
//line /usr/local/go/src/net/udpsock_posix.go:245
		_go_fuzz_dep_.CoverTab[16903]++
								if err := listenIPv4MulticastUDP(c, ifi, ip4); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:246
			_go_fuzz_dep_.CoverTab[16904]++
									c.Close()
									return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:248
			// _ = "end of CoverTab[16904]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:249
			_go_fuzz_dep_.CoverTab[16905]++
//line /usr/local/go/src/net/udpsock_posix.go:249
			// _ = "end of CoverTab[16905]"
//line /usr/local/go/src/net/udpsock_posix.go:249
		}
//line /usr/local/go/src/net/udpsock_posix.go:249
		// _ = "end of CoverTab[16903]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:250
		_go_fuzz_dep_.CoverTab[16906]++
								if err := listenIPv6MulticastUDP(c, ifi, gaddr.IP); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:251
			_go_fuzz_dep_.CoverTab[16907]++
									c.Close()
									return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:253
			// _ = "end of CoverTab[16907]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:254
			_go_fuzz_dep_.CoverTab[16908]++
//line /usr/local/go/src/net/udpsock_posix.go:254
			// _ = "end of CoverTab[16908]"
//line /usr/local/go/src/net/udpsock_posix.go:254
		}
//line /usr/local/go/src/net/udpsock_posix.go:254
		// _ = "end of CoverTab[16906]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:255
	// _ = "end of CoverTab[16896]"
//line /usr/local/go/src/net/udpsock_posix.go:255
	_go_fuzz_dep_.CoverTab[16897]++
							return c, nil
//line /usr/local/go/src/net/udpsock_posix.go:256
	// _ = "end of CoverTab[16897]"
}

func listenIPv4MulticastUDP(c *UDPConn, ifi *Interface, ip IP) error {
//line /usr/local/go/src/net/udpsock_posix.go:259
	_go_fuzz_dep_.CoverTab[16909]++
							if ifi != nil {
//line /usr/local/go/src/net/udpsock_posix.go:260
		_go_fuzz_dep_.CoverTab[16913]++
								if err := setIPv4MulticastInterface(c.fd, ifi); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:261
			_go_fuzz_dep_.CoverTab[16914]++
									return err
//line /usr/local/go/src/net/udpsock_posix.go:262
			// _ = "end of CoverTab[16914]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:263
			_go_fuzz_dep_.CoverTab[16915]++
//line /usr/local/go/src/net/udpsock_posix.go:263
			// _ = "end of CoverTab[16915]"
//line /usr/local/go/src/net/udpsock_posix.go:263
		}
//line /usr/local/go/src/net/udpsock_posix.go:263
		// _ = "end of CoverTab[16913]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:264
		_go_fuzz_dep_.CoverTab[16916]++
//line /usr/local/go/src/net/udpsock_posix.go:264
		// _ = "end of CoverTab[16916]"
//line /usr/local/go/src/net/udpsock_posix.go:264
	}
//line /usr/local/go/src/net/udpsock_posix.go:264
	// _ = "end of CoverTab[16909]"
//line /usr/local/go/src/net/udpsock_posix.go:264
	_go_fuzz_dep_.CoverTab[16910]++
							if err := setIPv4MulticastLoopback(c.fd, false); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:265
		_go_fuzz_dep_.CoverTab[16917]++
								return err
//line /usr/local/go/src/net/udpsock_posix.go:266
		// _ = "end of CoverTab[16917]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:267
		_go_fuzz_dep_.CoverTab[16918]++
//line /usr/local/go/src/net/udpsock_posix.go:267
		// _ = "end of CoverTab[16918]"
//line /usr/local/go/src/net/udpsock_posix.go:267
	}
//line /usr/local/go/src/net/udpsock_posix.go:267
	// _ = "end of CoverTab[16910]"
//line /usr/local/go/src/net/udpsock_posix.go:267
	_go_fuzz_dep_.CoverTab[16911]++
							if err := joinIPv4Group(c.fd, ifi, ip); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:268
		_go_fuzz_dep_.CoverTab[16919]++
								return err
//line /usr/local/go/src/net/udpsock_posix.go:269
		// _ = "end of CoverTab[16919]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:270
		_go_fuzz_dep_.CoverTab[16920]++
//line /usr/local/go/src/net/udpsock_posix.go:270
		// _ = "end of CoverTab[16920]"
//line /usr/local/go/src/net/udpsock_posix.go:270
	}
//line /usr/local/go/src/net/udpsock_posix.go:270
	// _ = "end of CoverTab[16911]"
//line /usr/local/go/src/net/udpsock_posix.go:270
	_go_fuzz_dep_.CoverTab[16912]++
							return nil
//line /usr/local/go/src/net/udpsock_posix.go:271
	// _ = "end of CoverTab[16912]"
}

func listenIPv6MulticastUDP(c *UDPConn, ifi *Interface, ip IP) error {
//line /usr/local/go/src/net/udpsock_posix.go:274
	_go_fuzz_dep_.CoverTab[16921]++
							if ifi != nil {
//line /usr/local/go/src/net/udpsock_posix.go:275
		_go_fuzz_dep_.CoverTab[16925]++
								if err := setIPv6MulticastInterface(c.fd, ifi); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:276
			_go_fuzz_dep_.CoverTab[16926]++
									return err
//line /usr/local/go/src/net/udpsock_posix.go:277
			// _ = "end of CoverTab[16926]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:278
			_go_fuzz_dep_.CoverTab[16927]++
//line /usr/local/go/src/net/udpsock_posix.go:278
			// _ = "end of CoverTab[16927]"
//line /usr/local/go/src/net/udpsock_posix.go:278
		}
//line /usr/local/go/src/net/udpsock_posix.go:278
		// _ = "end of CoverTab[16925]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:279
		_go_fuzz_dep_.CoverTab[16928]++
//line /usr/local/go/src/net/udpsock_posix.go:279
		// _ = "end of CoverTab[16928]"
//line /usr/local/go/src/net/udpsock_posix.go:279
	}
//line /usr/local/go/src/net/udpsock_posix.go:279
	// _ = "end of CoverTab[16921]"
//line /usr/local/go/src/net/udpsock_posix.go:279
	_go_fuzz_dep_.CoverTab[16922]++
							if err := setIPv6MulticastLoopback(c.fd, false); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:280
		_go_fuzz_dep_.CoverTab[16929]++
								return err
//line /usr/local/go/src/net/udpsock_posix.go:281
		// _ = "end of CoverTab[16929]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:282
		_go_fuzz_dep_.CoverTab[16930]++
//line /usr/local/go/src/net/udpsock_posix.go:282
		// _ = "end of CoverTab[16930]"
//line /usr/local/go/src/net/udpsock_posix.go:282
	}
//line /usr/local/go/src/net/udpsock_posix.go:282
	// _ = "end of CoverTab[16922]"
//line /usr/local/go/src/net/udpsock_posix.go:282
	_go_fuzz_dep_.CoverTab[16923]++
							if err := joinIPv6Group(c.fd, ifi, ip); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:283
		_go_fuzz_dep_.CoverTab[16931]++
								return err
//line /usr/local/go/src/net/udpsock_posix.go:284
		// _ = "end of CoverTab[16931]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:285
		_go_fuzz_dep_.CoverTab[16932]++
//line /usr/local/go/src/net/udpsock_posix.go:285
		// _ = "end of CoverTab[16932]"
//line /usr/local/go/src/net/udpsock_posix.go:285
	}
//line /usr/local/go/src/net/udpsock_posix.go:285
	// _ = "end of CoverTab[16923]"
//line /usr/local/go/src/net/udpsock_posix.go:285
	_go_fuzz_dep_.CoverTab[16924]++
							return nil
//line /usr/local/go/src/net/udpsock_posix.go:286
	// _ = "end of CoverTab[16924]"
}

//line /usr/local/go/src/net/udpsock_posix.go:287
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/udpsock_posix.go:287
var _ = _go_fuzz_dep_.CoverTab
