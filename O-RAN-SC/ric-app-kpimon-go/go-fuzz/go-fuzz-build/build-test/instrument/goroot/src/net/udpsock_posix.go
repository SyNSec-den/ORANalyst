// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

//line /snap/go/10455/src/net/udpsock_posix.go:7
package net

//line /snap/go/10455/src/net/udpsock_posix.go:7
import (
//line /snap/go/10455/src/net/udpsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/udpsock_posix.go:7
)
//line /snap/go/10455/src/net/udpsock_posix.go:7
import (
//line /snap/go/10455/src/net/udpsock_posix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/udpsock_posix.go:7
)

import (
	"context"
	"net/netip"
	"syscall"
)

func sockaddrToUDP(sa syscall.Sockaddr) Addr {
//line /snap/go/10455/src/net/udpsock_posix.go:15
	_go_fuzz_dep_.CoverTab[8679]++
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /snap/go/10455/src/net/udpsock_posix.go:17
		_go_fuzz_dep_.CoverTab[530011]++
//line /snap/go/10455/src/net/udpsock_posix.go:17
		_go_fuzz_dep_.CoverTab[8681]++
								return &UDPAddr{IP: sa.Addr[0:], Port: sa.Port}
//line /snap/go/10455/src/net/udpsock_posix.go:18
		// _ = "end of CoverTab[8681]"
	case *syscall.SockaddrInet6:
//line /snap/go/10455/src/net/udpsock_posix.go:19
		_go_fuzz_dep_.CoverTab[530012]++
//line /snap/go/10455/src/net/udpsock_posix.go:19
		_go_fuzz_dep_.CoverTab[8682]++
								return &UDPAddr{IP: sa.Addr[0:], Port: sa.Port, Zone: zoneCache.name(int(sa.ZoneId))}
//line /snap/go/10455/src/net/udpsock_posix.go:20
		// _ = "end of CoverTab[8682]"
	}
//line /snap/go/10455/src/net/udpsock_posix.go:21
	// _ = "end of CoverTab[8679]"
//line /snap/go/10455/src/net/udpsock_posix.go:21
	_go_fuzz_dep_.CoverTab[8680]++
							return nil
//line /snap/go/10455/src/net/udpsock_posix.go:22
	// _ = "end of CoverTab[8680]"
}

func (a *UDPAddr) family() int {
//line /snap/go/10455/src/net/udpsock_posix.go:25
	_go_fuzz_dep_.CoverTab[8683]++
							if a == nil || func() bool {
//line /snap/go/10455/src/net/udpsock_posix.go:26
		_go_fuzz_dep_.CoverTab[8686]++
//line /snap/go/10455/src/net/udpsock_posix.go:26
		return len(a.IP) <= IPv4len
//line /snap/go/10455/src/net/udpsock_posix.go:26
		// _ = "end of CoverTab[8686]"
//line /snap/go/10455/src/net/udpsock_posix.go:26
	}() {
//line /snap/go/10455/src/net/udpsock_posix.go:26
		_go_fuzz_dep_.CoverTab[530013]++
//line /snap/go/10455/src/net/udpsock_posix.go:26
		_go_fuzz_dep_.CoverTab[8687]++
								return syscall.AF_INET
//line /snap/go/10455/src/net/udpsock_posix.go:27
		// _ = "end of CoverTab[8687]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:28
		_go_fuzz_dep_.CoverTab[530014]++
//line /snap/go/10455/src/net/udpsock_posix.go:28
		_go_fuzz_dep_.CoverTab[8688]++
//line /snap/go/10455/src/net/udpsock_posix.go:28
		// _ = "end of CoverTab[8688]"
//line /snap/go/10455/src/net/udpsock_posix.go:28
	}
//line /snap/go/10455/src/net/udpsock_posix.go:28
	// _ = "end of CoverTab[8683]"
//line /snap/go/10455/src/net/udpsock_posix.go:28
	_go_fuzz_dep_.CoverTab[8684]++
							if a.IP.To4() != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:29
		_go_fuzz_dep_.CoverTab[530015]++
//line /snap/go/10455/src/net/udpsock_posix.go:29
		_go_fuzz_dep_.CoverTab[8689]++
								return syscall.AF_INET
//line /snap/go/10455/src/net/udpsock_posix.go:30
		// _ = "end of CoverTab[8689]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:31
		_go_fuzz_dep_.CoverTab[530016]++
//line /snap/go/10455/src/net/udpsock_posix.go:31
		_go_fuzz_dep_.CoverTab[8690]++
//line /snap/go/10455/src/net/udpsock_posix.go:31
		// _ = "end of CoverTab[8690]"
//line /snap/go/10455/src/net/udpsock_posix.go:31
	}
//line /snap/go/10455/src/net/udpsock_posix.go:31
	// _ = "end of CoverTab[8684]"
//line /snap/go/10455/src/net/udpsock_posix.go:31
	_go_fuzz_dep_.CoverTab[8685]++
							return syscall.AF_INET6
//line /snap/go/10455/src/net/udpsock_posix.go:32
	// _ = "end of CoverTab[8685]"
}

func (a *UDPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /snap/go/10455/src/net/udpsock_posix.go:35
	_go_fuzz_dep_.CoverTab[8691]++
							if a == nil {
//line /snap/go/10455/src/net/udpsock_posix.go:36
		_go_fuzz_dep_.CoverTab[530017]++
//line /snap/go/10455/src/net/udpsock_posix.go:36
		_go_fuzz_dep_.CoverTab[8693]++
								return nil, nil
//line /snap/go/10455/src/net/udpsock_posix.go:37
		// _ = "end of CoverTab[8693]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:38
		_go_fuzz_dep_.CoverTab[530018]++
//line /snap/go/10455/src/net/udpsock_posix.go:38
		_go_fuzz_dep_.CoverTab[8694]++
//line /snap/go/10455/src/net/udpsock_posix.go:38
		// _ = "end of CoverTab[8694]"
//line /snap/go/10455/src/net/udpsock_posix.go:38
	}
//line /snap/go/10455/src/net/udpsock_posix.go:38
	// _ = "end of CoverTab[8691]"
//line /snap/go/10455/src/net/udpsock_posix.go:38
	_go_fuzz_dep_.CoverTab[8692]++
							return ipToSockaddr(family, a.IP, a.Port, a.Zone)
//line /snap/go/10455/src/net/udpsock_posix.go:39
	// _ = "end of CoverTab[8692]"
}

func (a *UDPAddr) toLocal(net string) sockaddr {
//line /snap/go/10455/src/net/udpsock_posix.go:42
	_go_fuzz_dep_.CoverTab[8695]++
							return &UDPAddr{loopbackIP(net), a.Port, a.Zone}
//line /snap/go/10455/src/net/udpsock_posix.go:43
	// _ = "end of CoverTab[8695]"
}

func (c *UDPConn) readFrom(b []byte, addr *UDPAddr) (int, *UDPAddr, error) {
//line /snap/go/10455/src/net/udpsock_posix.go:46
	_go_fuzz_dep_.CoverTab[8696]++
							var n int
							var err error
							switch c.fd.family {
	case syscall.AF_INET:
//line /snap/go/10455/src/net/udpsock_posix.go:50
		_go_fuzz_dep_.CoverTab[530019]++
//line /snap/go/10455/src/net/udpsock_posix.go:50
		_go_fuzz_dep_.CoverTab[8699]++
								var from syscall.SockaddrInet4
								n, err = c.fd.readFromInet4(b, &from)
								if err == nil {
//line /snap/go/10455/src/net/udpsock_posix.go:53
			_go_fuzz_dep_.CoverTab[530022]++
//line /snap/go/10455/src/net/udpsock_posix.go:53
			_go_fuzz_dep_.CoverTab[8702]++
									ip := from.Addr
									*addr = UDPAddr{IP: ip[:], Port: from.Port}
//line /snap/go/10455/src/net/udpsock_posix.go:55
			// _ = "end of CoverTab[8702]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:56
			_go_fuzz_dep_.CoverTab[530023]++
//line /snap/go/10455/src/net/udpsock_posix.go:56
			_go_fuzz_dep_.CoverTab[8703]++
//line /snap/go/10455/src/net/udpsock_posix.go:56
			// _ = "end of CoverTab[8703]"
//line /snap/go/10455/src/net/udpsock_posix.go:56
		}
//line /snap/go/10455/src/net/udpsock_posix.go:56
		// _ = "end of CoverTab[8699]"
	case syscall.AF_INET6:
//line /snap/go/10455/src/net/udpsock_posix.go:57
		_go_fuzz_dep_.CoverTab[530020]++
//line /snap/go/10455/src/net/udpsock_posix.go:57
		_go_fuzz_dep_.CoverTab[8700]++
								var from syscall.SockaddrInet6
								n, err = c.fd.readFromInet6(b, &from)
								if err == nil {
//line /snap/go/10455/src/net/udpsock_posix.go:60
			_go_fuzz_dep_.CoverTab[530024]++
//line /snap/go/10455/src/net/udpsock_posix.go:60
			_go_fuzz_dep_.CoverTab[8704]++
									ip := from.Addr
									*addr = UDPAddr{IP: ip[:], Port: from.Port, Zone: zoneCache.name(int(from.ZoneId))}
//line /snap/go/10455/src/net/udpsock_posix.go:62
			// _ = "end of CoverTab[8704]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:63
			_go_fuzz_dep_.CoverTab[530025]++
//line /snap/go/10455/src/net/udpsock_posix.go:63
			_go_fuzz_dep_.CoverTab[8705]++
//line /snap/go/10455/src/net/udpsock_posix.go:63
			// _ = "end of CoverTab[8705]"
//line /snap/go/10455/src/net/udpsock_posix.go:63
		}
//line /snap/go/10455/src/net/udpsock_posix.go:63
		// _ = "end of CoverTab[8700]"
//line /snap/go/10455/src/net/udpsock_posix.go:63
	default:
//line /snap/go/10455/src/net/udpsock_posix.go:63
		_go_fuzz_dep_.CoverTab[530021]++
//line /snap/go/10455/src/net/udpsock_posix.go:63
		_go_fuzz_dep_.CoverTab[8701]++
//line /snap/go/10455/src/net/udpsock_posix.go:63
		// _ = "end of CoverTab[8701]"
	}
//line /snap/go/10455/src/net/udpsock_posix.go:64
	// _ = "end of CoverTab[8696]"
//line /snap/go/10455/src/net/udpsock_posix.go:64
	_go_fuzz_dep_.CoverTab[8697]++
							if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:65
		_go_fuzz_dep_.CoverTab[530026]++
//line /snap/go/10455/src/net/udpsock_posix.go:65
		_go_fuzz_dep_.CoverTab[8706]++

								addr = nil
//line /snap/go/10455/src/net/udpsock_posix.go:67
		// _ = "end of CoverTab[8706]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:68
		_go_fuzz_dep_.CoverTab[530027]++
//line /snap/go/10455/src/net/udpsock_posix.go:68
		_go_fuzz_dep_.CoverTab[8707]++
//line /snap/go/10455/src/net/udpsock_posix.go:68
		// _ = "end of CoverTab[8707]"
//line /snap/go/10455/src/net/udpsock_posix.go:68
	}
//line /snap/go/10455/src/net/udpsock_posix.go:68
	// _ = "end of CoverTab[8697]"
//line /snap/go/10455/src/net/udpsock_posix.go:68
	_go_fuzz_dep_.CoverTab[8698]++
							return n, addr, err
//line /snap/go/10455/src/net/udpsock_posix.go:69
	// _ = "end of CoverTab[8698]"
}

func (c *UDPConn) readFromAddrPort(b []byte) (n int, addr netip.AddrPort, err error) {
//line /snap/go/10455/src/net/udpsock_posix.go:72
	_go_fuzz_dep_.CoverTab[8708]++
							var ip netip.Addr
							var port int
							switch c.fd.family {
	case syscall.AF_INET:
//line /snap/go/10455/src/net/udpsock_posix.go:76
		_go_fuzz_dep_.CoverTab[530028]++
//line /snap/go/10455/src/net/udpsock_posix.go:76
		_go_fuzz_dep_.CoverTab[8711]++
								var from syscall.SockaddrInet4
								n, err = c.fd.readFromInet4(b, &from)
								if err == nil {
//line /snap/go/10455/src/net/udpsock_posix.go:79
			_go_fuzz_dep_.CoverTab[530031]++
//line /snap/go/10455/src/net/udpsock_posix.go:79
			_go_fuzz_dep_.CoverTab[8714]++
									ip = netip.AddrFrom4(from.Addr)
									port = from.Port
//line /snap/go/10455/src/net/udpsock_posix.go:81
			// _ = "end of CoverTab[8714]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:82
			_go_fuzz_dep_.CoverTab[530032]++
//line /snap/go/10455/src/net/udpsock_posix.go:82
			_go_fuzz_dep_.CoverTab[8715]++
//line /snap/go/10455/src/net/udpsock_posix.go:82
			// _ = "end of CoverTab[8715]"
//line /snap/go/10455/src/net/udpsock_posix.go:82
		}
//line /snap/go/10455/src/net/udpsock_posix.go:82
		// _ = "end of CoverTab[8711]"
	case syscall.AF_INET6:
//line /snap/go/10455/src/net/udpsock_posix.go:83
		_go_fuzz_dep_.CoverTab[530029]++
//line /snap/go/10455/src/net/udpsock_posix.go:83
		_go_fuzz_dep_.CoverTab[8712]++
								var from syscall.SockaddrInet6
								n, err = c.fd.readFromInet6(b, &from)
								if err == nil {
//line /snap/go/10455/src/net/udpsock_posix.go:86
			_go_fuzz_dep_.CoverTab[530033]++
//line /snap/go/10455/src/net/udpsock_posix.go:86
			_go_fuzz_dep_.CoverTab[8716]++
									ip = netip.AddrFrom16(from.Addr).WithZone(zoneCache.name(int(from.ZoneId)))
									port = from.Port
//line /snap/go/10455/src/net/udpsock_posix.go:88
			// _ = "end of CoverTab[8716]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:89
			_go_fuzz_dep_.CoverTab[530034]++
//line /snap/go/10455/src/net/udpsock_posix.go:89
			_go_fuzz_dep_.CoverTab[8717]++
//line /snap/go/10455/src/net/udpsock_posix.go:89
			// _ = "end of CoverTab[8717]"
//line /snap/go/10455/src/net/udpsock_posix.go:89
		}
//line /snap/go/10455/src/net/udpsock_posix.go:89
		// _ = "end of CoverTab[8712]"
//line /snap/go/10455/src/net/udpsock_posix.go:89
	default:
//line /snap/go/10455/src/net/udpsock_posix.go:89
		_go_fuzz_dep_.CoverTab[530030]++
//line /snap/go/10455/src/net/udpsock_posix.go:89
		_go_fuzz_dep_.CoverTab[8713]++
//line /snap/go/10455/src/net/udpsock_posix.go:89
		// _ = "end of CoverTab[8713]"
	}
//line /snap/go/10455/src/net/udpsock_posix.go:90
	// _ = "end of CoverTab[8708]"
//line /snap/go/10455/src/net/udpsock_posix.go:90
	_go_fuzz_dep_.CoverTab[8709]++
							if err == nil {
//line /snap/go/10455/src/net/udpsock_posix.go:91
		_go_fuzz_dep_.CoverTab[530035]++
//line /snap/go/10455/src/net/udpsock_posix.go:91
		_go_fuzz_dep_.CoverTab[8718]++
								addr = netip.AddrPortFrom(ip, uint16(port))
//line /snap/go/10455/src/net/udpsock_posix.go:92
		// _ = "end of CoverTab[8718]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:93
		_go_fuzz_dep_.CoverTab[530036]++
//line /snap/go/10455/src/net/udpsock_posix.go:93
		_go_fuzz_dep_.CoverTab[8719]++
//line /snap/go/10455/src/net/udpsock_posix.go:93
		// _ = "end of CoverTab[8719]"
//line /snap/go/10455/src/net/udpsock_posix.go:93
	}
//line /snap/go/10455/src/net/udpsock_posix.go:93
	// _ = "end of CoverTab[8709]"
//line /snap/go/10455/src/net/udpsock_posix.go:93
	_go_fuzz_dep_.CoverTab[8710]++
							return n, addr, err
//line /snap/go/10455/src/net/udpsock_posix.go:94
	// _ = "end of CoverTab[8710]"
}

func (c *UDPConn) readMsg(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error) {
//line /snap/go/10455/src/net/udpsock_posix.go:97
	_go_fuzz_dep_.CoverTab[8720]++
							switch c.fd.family {
	case syscall.AF_INET:
//line /snap/go/10455/src/net/udpsock_posix.go:99
		_go_fuzz_dep_.CoverTab[530037]++
//line /snap/go/10455/src/net/udpsock_posix.go:99
		_go_fuzz_dep_.CoverTab[8722]++
								var sa syscall.SockaddrInet4
								n, oobn, flags, err = c.fd.readMsgInet4(b, oob, 0, &sa)
								ip := netip.AddrFrom4(sa.Addr)
								addr = netip.AddrPortFrom(ip, uint16(sa.Port))
//line /snap/go/10455/src/net/udpsock_posix.go:103
		// _ = "end of CoverTab[8722]"
	case syscall.AF_INET6:
//line /snap/go/10455/src/net/udpsock_posix.go:104
		_go_fuzz_dep_.CoverTab[530038]++
//line /snap/go/10455/src/net/udpsock_posix.go:104
		_go_fuzz_dep_.CoverTab[8723]++
								var sa syscall.SockaddrInet6
								n, oobn, flags, err = c.fd.readMsgInet6(b, oob, 0, &sa)
								ip := netip.AddrFrom16(sa.Addr).WithZone(zoneCache.name(int(sa.ZoneId)))
								addr = netip.AddrPortFrom(ip, uint16(sa.Port))
//line /snap/go/10455/src/net/udpsock_posix.go:108
		// _ = "end of CoverTab[8723]"
//line /snap/go/10455/src/net/udpsock_posix.go:108
	default:
//line /snap/go/10455/src/net/udpsock_posix.go:108
		_go_fuzz_dep_.CoverTab[530039]++
//line /snap/go/10455/src/net/udpsock_posix.go:108
		_go_fuzz_dep_.CoverTab[8724]++
//line /snap/go/10455/src/net/udpsock_posix.go:108
		// _ = "end of CoverTab[8724]"
	}
//line /snap/go/10455/src/net/udpsock_posix.go:109
	// _ = "end of CoverTab[8720]"
//line /snap/go/10455/src/net/udpsock_posix.go:109
	_go_fuzz_dep_.CoverTab[8721]++
							return
//line /snap/go/10455/src/net/udpsock_posix.go:110
	// _ = "end of CoverTab[8721]"
}

func (c *UDPConn) writeTo(b []byte, addr *UDPAddr) (int, error) {
//line /snap/go/10455/src/net/udpsock_posix.go:113
	_go_fuzz_dep_.CoverTab[8725]++
							if c.fd.isConnected {
//line /snap/go/10455/src/net/udpsock_posix.go:114
		_go_fuzz_dep_.CoverTab[530040]++
//line /snap/go/10455/src/net/udpsock_posix.go:114
		_go_fuzz_dep_.CoverTab[8728]++
								return 0, ErrWriteToConnected
//line /snap/go/10455/src/net/udpsock_posix.go:115
		// _ = "end of CoverTab[8728]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:116
		_go_fuzz_dep_.CoverTab[530041]++
//line /snap/go/10455/src/net/udpsock_posix.go:116
		_go_fuzz_dep_.CoverTab[8729]++
//line /snap/go/10455/src/net/udpsock_posix.go:116
		// _ = "end of CoverTab[8729]"
//line /snap/go/10455/src/net/udpsock_posix.go:116
	}
//line /snap/go/10455/src/net/udpsock_posix.go:116
	// _ = "end of CoverTab[8725]"
//line /snap/go/10455/src/net/udpsock_posix.go:116
	_go_fuzz_dep_.CoverTab[8726]++
							if addr == nil {
//line /snap/go/10455/src/net/udpsock_posix.go:117
		_go_fuzz_dep_.CoverTab[530042]++
//line /snap/go/10455/src/net/udpsock_posix.go:117
		_go_fuzz_dep_.CoverTab[8730]++
								return 0, errMissingAddress
//line /snap/go/10455/src/net/udpsock_posix.go:118
		// _ = "end of CoverTab[8730]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:119
		_go_fuzz_dep_.CoverTab[530043]++
//line /snap/go/10455/src/net/udpsock_posix.go:119
		_go_fuzz_dep_.CoverTab[8731]++
//line /snap/go/10455/src/net/udpsock_posix.go:119
		// _ = "end of CoverTab[8731]"
//line /snap/go/10455/src/net/udpsock_posix.go:119
	}
//line /snap/go/10455/src/net/udpsock_posix.go:119
	// _ = "end of CoverTab[8726]"
//line /snap/go/10455/src/net/udpsock_posix.go:119
	_go_fuzz_dep_.CoverTab[8727]++

							switch c.fd.family {
	case syscall.AF_INET:
//line /snap/go/10455/src/net/udpsock_posix.go:122
		_go_fuzz_dep_.CoverTab[530044]++
//line /snap/go/10455/src/net/udpsock_posix.go:122
		_go_fuzz_dep_.CoverTab[8732]++
								sa, err := ipToSockaddrInet4(addr.IP, addr.Port)
								if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:124
			_go_fuzz_dep_.CoverTab[530047]++
//line /snap/go/10455/src/net/udpsock_posix.go:124
			_go_fuzz_dep_.CoverTab[8737]++
									return 0, err
//line /snap/go/10455/src/net/udpsock_posix.go:125
			// _ = "end of CoverTab[8737]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:126
			_go_fuzz_dep_.CoverTab[530048]++
//line /snap/go/10455/src/net/udpsock_posix.go:126
			_go_fuzz_dep_.CoverTab[8738]++
//line /snap/go/10455/src/net/udpsock_posix.go:126
			// _ = "end of CoverTab[8738]"
//line /snap/go/10455/src/net/udpsock_posix.go:126
		}
//line /snap/go/10455/src/net/udpsock_posix.go:126
		// _ = "end of CoverTab[8732]"
//line /snap/go/10455/src/net/udpsock_posix.go:126
		_go_fuzz_dep_.CoverTab[8733]++
								return c.fd.writeToInet4(b, &sa)
//line /snap/go/10455/src/net/udpsock_posix.go:127
		// _ = "end of CoverTab[8733]"
	case syscall.AF_INET6:
//line /snap/go/10455/src/net/udpsock_posix.go:128
		_go_fuzz_dep_.CoverTab[530045]++
//line /snap/go/10455/src/net/udpsock_posix.go:128
		_go_fuzz_dep_.CoverTab[8734]++
								sa, err := ipToSockaddrInet6(addr.IP, addr.Port, addr.Zone)
								if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:130
			_go_fuzz_dep_.CoverTab[530049]++
//line /snap/go/10455/src/net/udpsock_posix.go:130
			_go_fuzz_dep_.CoverTab[8739]++
									return 0, err
//line /snap/go/10455/src/net/udpsock_posix.go:131
			// _ = "end of CoverTab[8739]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:132
			_go_fuzz_dep_.CoverTab[530050]++
//line /snap/go/10455/src/net/udpsock_posix.go:132
			_go_fuzz_dep_.CoverTab[8740]++
//line /snap/go/10455/src/net/udpsock_posix.go:132
			// _ = "end of CoverTab[8740]"
//line /snap/go/10455/src/net/udpsock_posix.go:132
		}
//line /snap/go/10455/src/net/udpsock_posix.go:132
		// _ = "end of CoverTab[8734]"
//line /snap/go/10455/src/net/udpsock_posix.go:132
		_go_fuzz_dep_.CoverTab[8735]++
								return c.fd.writeToInet6(b, &sa)
//line /snap/go/10455/src/net/udpsock_posix.go:133
		// _ = "end of CoverTab[8735]"
	default:
//line /snap/go/10455/src/net/udpsock_posix.go:134
		_go_fuzz_dep_.CoverTab[530046]++
//line /snap/go/10455/src/net/udpsock_posix.go:134
		_go_fuzz_dep_.CoverTab[8736]++
								return 0, &AddrError{Err: "invalid address family", Addr: addr.IP.String()}
//line /snap/go/10455/src/net/udpsock_posix.go:135
		// _ = "end of CoverTab[8736]"
	}
//line /snap/go/10455/src/net/udpsock_posix.go:136
	// _ = "end of CoverTab[8727]"
}

func (c *UDPConn) writeToAddrPort(b []byte, addr netip.AddrPort) (int, error) {
//line /snap/go/10455/src/net/udpsock_posix.go:139
	_go_fuzz_dep_.CoverTab[8741]++
							if c.fd.isConnected {
//line /snap/go/10455/src/net/udpsock_posix.go:140
		_go_fuzz_dep_.CoverTab[530051]++
//line /snap/go/10455/src/net/udpsock_posix.go:140
		_go_fuzz_dep_.CoverTab[8744]++
								return 0, ErrWriteToConnected
//line /snap/go/10455/src/net/udpsock_posix.go:141
		// _ = "end of CoverTab[8744]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:142
		_go_fuzz_dep_.CoverTab[530052]++
//line /snap/go/10455/src/net/udpsock_posix.go:142
		_go_fuzz_dep_.CoverTab[8745]++
//line /snap/go/10455/src/net/udpsock_posix.go:142
		// _ = "end of CoverTab[8745]"
//line /snap/go/10455/src/net/udpsock_posix.go:142
	}
//line /snap/go/10455/src/net/udpsock_posix.go:142
	// _ = "end of CoverTab[8741]"
//line /snap/go/10455/src/net/udpsock_posix.go:142
	_go_fuzz_dep_.CoverTab[8742]++
							if !addr.IsValid() {
//line /snap/go/10455/src/net/udpsock_posix.go:143
		_go_fuzz_dep_.CoverTab[530053]++
//line /snap/go/10455/src/net/udpsock_posix.go:143
		_go_fuzz_dep_.CoverTab[8746]++
								return 0, errMissingAddress
//line /snap/go/10455/src/net/udpsock_posix.go:144
		// _ = "end of CoverTab[8746]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:145
		_go_fuzz_dep_.CoverTab[530054]++
//line /snap/go/10455/src/net/udpsock_posix.go:145
		_go_fuzz_dep_.CoverTab[8747]++
//line /snap/go/10455/src/net/udpsock_posix.go:145
		// _ = "end of CoverTab[8747]"
//line /snap/go/10455/src/net/udpsock_posix.go:145
	}
//line /snap/go/10455/src/net/udpsock_posix.go:145
	// _ = "end of CoverTab[8742]"
//line /snap/go/10455/src/net/udpsock_posix.go:145
	_go_fuzz_dep_.CoverTab[8743]++

							switch c.fd.family {
	case syscall.AF_INET:
//line /snap/go/10455/src/net/udpsock_posix.go:148
		_go_fuzz_dep_.CoverTab[530055]++
//line /snap/go/10455/src/net/udpsock_posix.go:148
		_go_fuzz_dep_.CoverTab[8748]++
								sa, err := addrPortToSockaddrInet4(addr)
								if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:150
			_go_fuzz_dep_.CoverTab[530058]++
//line /snap/go/10455/src/net/udpsock_posix.go:150
			_go_fuzz_dep_.CoverTab[8753]++
									return 0, err
//line /snap/go/10455/src/net/udpsock_posix.go:151
			// _ = "end of CoverTab[8753]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:152
			_go_fuzz_dep_.CoverTab[530059]++
//line /snap/go/10455/src/net/udpsock_posix.go:152
			_go_fuzz_dep_.CoverTab[8754]++
//line /snap/go/10455/src/net/udpsock_posix.go:152
			// _ = "end of CoverTab[8754]"
//line /snap/go/10455/src/net/udpsock_posix.go:152
		}
//line /snap/go/10455/src/net/udpsock_posix.go:152
		// _ = "end of CoverTab[8748]"
//line /snap/go/10455/src/net/udpsock_posix.go:152
		_go_fuzz_dep_.CoverTab[8749]++
								return c.fd.writeToInet4(b, &sa)
//line /snap/go/10455/src/net/udpsock_posix.go:153
		// _ = "end of CoverTab[8749]"
	case syscall.AF_INET6:
//line /snap/go/10455/src/net/udpsock_posix.go:154
		_go_fuzz_dep_.CoverTab[530056]++
//line /snap/go/10455/src/net/udpsock_posix.go:154
		_go_fuzz_dep_.CoverTab[8750]++
								sa, err := addrPortToSockaddrInet6(addr)
								if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:156
			_go_fuzz_dep_.CoverTab[530060]++
//line /snap/go/10455/src/net/udpsock_posix.go:156
			_go_fuzz_dep_.CoverTab[8755]++
									return 0, err
//line /snap/go/10455/src/net/udpsock_posix.go:157
			// _ = "end of CoverTab[8755]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:158
			_go_fuzz_dep_.CoverTab[530061]++
//line /snap/go/10455/src/net/udpsock_posix.go:158
			_go_fuzz_dep_.CoverTab[8756]++
//line /snap/go/10455/src/net/udpsock_posix.go:158
			// _ = "end of CoverTab[8756]"
//line /snap/go/10455/src/net/udpsock_posix.go:158
		}
//line /snap/go/10455/src/net/udpsock_posix.go:158
		// _ = "end of CoverTab[8750]"
//line /snap/go/10455/src/net/udpsock_posix.go:158
		_go_fuzz_dep_.CoverTab[8751]++
								return c.fd.writeToInet6(b, &sa)
//line /snap/go/10455/src/net/udpsock_posix.go:159
		// _ = "end of CoverTab[8751]"
	default:
//line /snap/go/10455/src/net/udpsock_posix.go:160
		_go_fuzz_dep_.CoverTab[530057]++
//line /snap/go/10455/src/net/udpsock_posix.go:160
		_go_fuzz_dep_.CoverTab[8752]++
								return 0, &AddrError{Err: "invalid address family", Addr: addr.Addr().String()}
//line /snap/go/10455/src/net/udpsock_posix.go:161
		// _ = "end of CoverTab[8752]"
	}
//line /snap/go/10455/src/net/udpsock_posix.go:162
	// _ = "end of CoverTab[8743]"
}

func (c *UDPConn) writeMsg(b, oob []byte, addr *UDPAddr) (n, oobn int, err error) {
//line /snap/go/10455/src/net/udpsock_posix.go:165
	_go_fuzz_dep_.CoverTab[8757]++
							if c.fd.isConnected && func() bool {
//line /snap/go/10455/src/net/udpsock_posix.go:166
		_go_fuzz_dep_.CoverTab[8761]++
//line /snap/go/10455/src/net/udpsock_posix.go:166
		return addr != nil
//line /snap/go/10455/src/net/udpsock_posix.go:166
		// _ = "end of CoverTab[8761]"
//line /snap/go/10455/src/net/udpsock_posix.go:166
	}() {
//line /snap/go/10455/src/net/udpsock_posix.go:166
		_go_fuzz_dep_.CoverTab[530062]++
//line /snap/go/10455/src/net/udpsock_posix.go:166
		_go_fuzz_dep_.CoverTab[8762]++
								return 0, 0, ErrWriteToConnected
//line /snap/go/10455/src/net/udpsock_posix.go:167
		// _ = "end of CoverTab[8762]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:168
		_go_fuzz_dep_.CoverTab[530063]++
//line /snap/go/10455/src/net/udpsock_posix.go:168
		_go_fuzz_dep_.CoverTab[8763]++
//line /snap/go/10455/src/net/udpsock_posix.go:168
		// _ = "end of CoverTab[8763]"
//line /snap/go/10455/src/net/udpsock_posix.go:168
	}
//line /snap/go/10455/src/net/udpsock_posix.go:168
	// _ = "end of CoverTab[8757]"
//line /snap/go/10455/src/net/udpsock_posix.go:168
	_go_fuzz_dep_.CoverTab[8758]++
							if !c.fd.isConnected && func() bool {
//line /snap/go/10455/src/net/udpsock_posix.go:169
		_go_fuzz_dep_.CoverTab[8764]++
//line /snap/go/10455/src/net/udpsock_posix.go:169
		return addr == nil
//line /snap/go/10455/src/net/udpsock_posix.go:169
		// _ = "end of CoverTab[8764]"
//line /snap/go/10455/src/net/udpsock_posix.go:169
	}() {
//line /snap/go/10455/src/net/udpsock_posix.go:169
		_go_fuzz_dep_.CoverTab[530064]++
//line /snap/go/10455/src/net/udpsock_posix.go:169
		_go_fuzz_dep_.CoverTab[8765]++
								return 0, 0, errMissingAddress
//line /snap/go/10455/src/net/udpsock_posix.go:170
		// _ = "end of CoverTab[8765]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:171
		_go_fuzz_dep_.CoverTab[530065]++
//line /snap/go/10455/src/net/udpsock_posix.go:171
		_go_fuzz_dep_.CoverTab[8766]++
//line /snap/go/10455/src/net/udpsock_posix.go:171
		// _ = "end of CoverTab[8766]"
//line /snap/go/10455/src/net/udpsock_posix.go:171
	}
//line /snap/go/10455/src/net/udpsock_posix.go:171
	// _ = "end of CoverTab[8758]"
//line /snap/go/10455/src/net/udpsock_posix.go:171
	_go_fuzz_dep_.CoverTab[8759]++
							sa, err := addr.sockaddr(c.fd.family)
							if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:173
		_go_fuzz_dep_.CoverTab[530066]++
//line /snap/go/10455/src/net/udpsock_posix.go:173
		_go_fuzz_dep_.CoverTab[8767]++
								return 0, 0, err
//line /snap/go/10455/src/net/udpsock_posix.go:174
		// _ = "end of CoverTab[8767]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:175
		_go_fuzz_dep_.CoverTab[530067]++
//line /snap/go/10455/src/net/udpsock_posix.go:175
		_go_fuzz_dep_.CoverTab[8768]++
//line /snap/go/10455/src/net/udpsock_posix.go:175
		// _ = "end of CoverTab[8768]"
//line /snap/go/10455/src/net/udpsock_posix.go:175
	}
//line /snap/go/10455/src/net/udpsock_posix.go:175
	// _ = "end of CoverTab[8759]"
//line /snap/go/10455/src/net/udpsock_posix.go:175
	_go_fuzz_dep_.CoverTab[8760]++
							return c.fd.writeMsg(b, oob, sa)
//line /snap/go/10455/src/net/udpsock_posix.go:176
	// _ = "end of CoverTab[8760]"
}

func (c *UDPConn) writeMsgAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error) {
//line /snap/go/10455/src/net/udpsock_posix.go:179
	_go_fuzz_dep_.CoverTab[8769]++
							if c.fd.isConnected && func() bool {
//line /snap/go/10455/src/net/udpsock_posix.go:180
		_go_fuzz_dep_.CoverTab[8772]++
//line /snap/go/10455/src/net/udpsock_posix.go:180
		return addr.IsValid()
//line /snap/go/10455/src/net/udpsock_posix.go:180
		// _ = "end of CoverTab[8772]"
//line /snap/go/10455/src/net/udpsock_posix.go:180
	}() {
//line /snap/go/10455/src/net/udpsock_posix.go:180
		_go_fuzz_dep_.CoverTab[530068]++
//line /snap/go/10455/src/net/udpsock_posix.go:180
		_go_fuzz_dep_.CoverTab[8773]++
								return 0, 0, ErrWriteToConnected
//line /snap/go/10455/src/net/udpsock_posix.go:181
		// _ = "end of CoverTab[8773]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:182
		_go_fuzz_dep_.CoverTab[530069]++
//line /snap/go/10455/src/net/udpsock_posix.go:182
		_go_fuzz_dep_.CoverTab[8774]++
//line /snap/go/10455/src/net/udpsock_posix.go:182
		// _ = "end of CoverTab[8774]"
//line /snap/go/10455/src/net/udpsock_posix.go:182
	}
//line /snap/go/10455/src/net/udpsock_posix.go:182
	// _ = "end of CoverTab[8769]"
//line /snap/go/10455/src/net/udpsock_posix.go:182
	_go_fuzz_dep_.CoverTab[8770]++
							if !c.fd.isConnected && func() bool {
//line /snap/go/10455/src/net/udpsock_posix.go:183
		_go_fuzz_dep_.CoverTab[8775]++
//line /snap/go/10455/src/net/udpsock_posix.go:183
		return !addr.IsValid()
//line /snap/go/10455/src/net/udpsock_posix.go:183
		// _ = "end of CoverTab[8775]"
//line /snap/go/10455/src/net/udpsock_posix.go:183
	}() {
//line /snap/go/10455/src/net/udpsock_posix.go:183
		_go_fuzz_dep_.CoverTab[530070]++
//line /snap/go/10455/src/net/udpsock_posix.go:183
		_go_fuzz_dep_.CoverTab[8776]++
								return 0, 0, errMissingAddress
//line /snap/go/10455/src/net/udpsock_posix.go:184
		// _ = "end of CoverTab[8776]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:185
		_go_fuzz_dep_.CoverTab[530071]++
//line /snap/go/10455/src/net/udpsock_posix.go:185
		_go_fuzz_dep_.CoverTab[8777]++
//line /snap/go/10455/src/net/udpsock_posix.go:185
		// _ = "end of CoverTab[8777]"
//line /snap/go/10455/src/net/udpsock_posix.go:185
	}
//line /snap/go/10455/src/net/udpsock_posix.go:185
	// _ = "end of CoverTab[8770]"
//line /snap/go/10455/src/net/udpsock_posix.go:185
	_go_fuzz_dep_.CoverTab[8771]++

							switch c.fd.family {
	case syscall.AF_INET:
//line /snap/go/10455/src/net/udpsock_posix.go:188
		_go_fuzz_dep_.CoverTab[530072]++
//line /snap/go/10455/src/net/udpsock_posix.go:188
		_go_fuzz_dep_.CoverTab[8778]++
								sa, err := addrPortToSockaddrInet4(addr)
								if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:190
			_go_fuzz_dep_.CoverTab[530075]++
//line /snap/go/10455/src/net/udpsock_posix.go:190
			_go_fuzz_dep_.CoverTab[8783]++
									return 0, 0, err
//line /snap/go/10455/src/net/udpsock_posix.go:191
			// _ = "end of CoverTab[8783]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:192
			_go_fuzz_dep_.CoverTab[530076]++
//line /snap/go/10455/src/net/udpsock_posix.go:192
			_go_fuzz_dep_.CoverTab[8784]++
//line /snap/go/10455/src/net/udpsock_posix.go:192
			// _ = "end of CoverTab[8784]"
//line /snap/go/10455/src/net/udpsock_posix.go:192
		}
//line /snap/go/10455/src/net/udpsock_posix.go:192
		// _ = "end of CoverTab[8778]"
//line /snap/go/10455/src/net/udpsock_posix.go:192
		_go_fuzz_dep_.CoverTab[8779]++
								return c.fd.writeMsgInet4(b, oob, &sa)
//line /snap/go/10455/src/net/udpsock_posix.go:193
		// _ = "end of CoverTab[8779]"
	case syscall.AF_INET6:
//line /snap/go/10455/src/net/udpsock_posix.go:194
		_go_fuzz_dep_.CoverTab[530073]++
//line /snap/go/10455/src/net/udpsock_posix.go:194
		_go_fuzz_dep_.CoverTab[8780]++
								sa, err := addrPortToSockaddrInet6(addr)
								if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:196
			_go_fuzz_dep_.CoverTab[530077]++
//line /snap/go/10455/src/net/udpsock_posix.go:196
			_go_fuzz_dep_.CoverTab[8785]++
									return 0, 0, err
//line /snap/go/10455/src/net/udpsock_posix.go:197
			// _ = "end of CoverTab[8785]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:198
			_go_fuzz_dep_.CoverTab[530078]++
//line /snap/go/10455/src/net/udpsock_posix.go:198
			_go_fuzz_dep_.CoverTab[8786]++
//line /snap/go/10455/src/net/udpsock_posix.go:198
			// _ = "end of CoverTab[8786]"
//line /snap/go/10455/src/net/udpsock_posix.go:198
		}
//line /snap/go/10455/src/net/udpsock_posix.go:198
		// _ = "end of CoverTab[8780]"
//line /snap/go/10455/src/net/udpsock_posix.go:198
		_go_fuzz_dep_.CoverTab[8781]++
								return c.fd.writeMsgInet6(b, oob, &sa)
//line /snap/go/10455/src/net/udpsock_posix.go:199
		// _ = "end of CoverTab[8781]"
	default:
//line /snap/go/10455/src/net/udpsock_posix.go:200
		_go_fuzz_dep_.CoverTab[530074]++
//line /snap/go/10455/src/net/udpsock_posix.go:200
		_go_fuzz_dep_.CoverTab[8782]++
								return 0, 0, &AddrError{Err: "invalid address family", Addr: addr.Addr().String()}
//line /snap/go/10455/src/net/udpsock_posix.go:201
		// _ = "end of CoverTab[8782]"
	}
//line /snap/go/10455/src/net/udpsock_posix.go:202
	// _ = "end of CoverTab[8771]"
}

func (sd *sysDialer) dialUDP(ctx context.Context, laddr, raddr *UDPAddr) (*UDPConn, error) {
//line /snap/go/10455/src/net/udpsock_posix.go:205
	_go_fuzz_dep_.CoverTab[8787]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /snap/go/10455/src/net/udpsock_posix.go:207
		_go_fuzz_dep_.CoverTab[8790]++
//line /snap/go/10455/src/net/udpsock_posix.go:207
		return sd.Dialer.Control != nil
//line /snap/go/10455/src/net/udpsock_posix.go:207
		// _ = "end of CoverTab[8790]"
//line /snap/go/10455/src/net/udpsock_posix.go:207
	}() {
//line /snap/go/10455/src/net/udpsock_posix.go:207
		_go_fuzz_dep_.CoverTab[530079]++
//line /snap/go/10455/src/net/udpsock_posix.go:207
		_go_fuzz_dep_.CoverTab[8791]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/udpsock_posix.go:208
			_go_fuzz_dep_.CoverTab[8792]++
									return sd.Dialer.Control(network, address, c)
//line /snap/go/10455/src/net/udpsock_posix.go:209
			// _ = "end of CoverTab[8792]"
		}
//line /snap/go/10455/src/net/udpsock_posix.go:210
		// _ = "end of CoverTab[8791]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:211
		_go_fuzz_dep_.CoverTab[530080]++
//line /snap/go/10455/src/net/udpsock_posix.go:211
		_go_fuzz_dep_.CoverTab[8793]++
//line /snap/go/10455/src/net/udpsock_posix.go:211
		// _ = "end of CoverTab[8793]"
//line /snap/go/10455/src/net/udpsock_posix.go:211
	}
//line /snap/go/10455/src/net/udpsock_posix.go:211
	// _ = "end of CoverTab[8787]"
//line /snap/go/10455/src/net/udpsock_posix.go:211
	_go_fuzz_dep_.CoverTab[8788]++
							fd, err := internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_DGRAM, 0, "dial", ctrlCtxFn)
							if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:213
		_go_fuzz_dep_.CoverTab[530081]++
//line /snap/go/10455/src/net/udpsock_posix.go:213
		_go_fuzz_dep_.CoverTab[8794]++
								return nil, err
//line /snap/go/10455/src/net/udpsock_posix.go:214
		// _ = "end of CoverTab[8794]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:215
		_go_fuzz_dep_.CoverTab[530082]++
//line /snap/go/10455/src/net/udpsock_posix.go:215
		_go_fuzz_dep_.CoverTab[8795]++
//line /snap/go/10455/src/net/udpsock_posix.go:215
		// _ = "end of CoverTab[8795]"
//line /snap/go/10455/src/net/udpsock_posix.go:215
	}
//line /snap/go/10455/src/net/udpsock_posix.go:215
	// _ = "end of CoverTab[8788]"
//line /snap/go/10455/src/net/udpsock_posix.go:215
	_go_fuzz_dep_.CoverTab[8789]++
							return newUDPConn(fd), nil
//line /snap/go/10455/src/net/udpsock_posix.go:216
	// _ = "end of CoverTab[8789]"
}

func (sl *sysListener) listenUDP(ctx context.Context, laddr *UDPAddr) (*UDPConn, error) {
//line /snap/go/10455/src/net/udpsock_posix.go:219
	_go_fuzz_dep_.CoverTab[8796]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:221
		_go_fuzz_dep_.CoverTab[530083]++
//line /snap/go/10455/src/net/udpsock_posix.go:221
		_go_fuzz_dep_.CoverTab[8799]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/udpsock_posix.go:222
			_go_fuzz_dep_.CoverTab[8800]++
									return sl.ListenConfig.Control(network, address, c)
//line /snap/go/10455/src/net/udpsock_posix.go:223
			// _ = "end of CoverTab[8800]"
		}
//line /snap/go/10455/src/net/udpsock_posix.go:224
		// _ = "end of CoverTab[8799]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:225
		_go_fuzz_dep_.CoverTab[530084]++
//line /snap/go/10455/src/net/udpsock_posix.go:225
		_go_fuzz_dep_.CoverTab[8801]++
//line /snap/go/10455/src/net/udpsock_posix.go:225
		// _ = "end of CoverTab[8801]"
//line /snap/go/10455/src/net/udpsock_posix.go:225
	}
//line /snap/go/10455/src/net/udpsock_posix.go:225
	// _ = "end of CoverTab[8796]"
//line /snap/go/10455/src/net/udpsock_posix.go:225
	_go_fuzz_dep_.CoverTab[8797]++
							fd, err := internetSocket(ctx, sl.network, laddr, nil, syscall.SOCK_DGRAM, 0, "listen", ctrlCtxFn)
							if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:227
		_go_fuzz_dep_.CoverTab[530085]++
//line /snap/go/10455/src/net/udpsock_posix.go:227
		_go_fuzz_dep_.CoverTab[8802]++
								return nil, err
//line /snap/go/10455/src/net/udpsock_posix.go:228
		// _ = "end of CoverTab[8802]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:229
		_go_fuzz_dep_.CoverTab[530086]++
//line /snap/go/10455/src/net/udpsock_posix.go:229
		_go_fuzz_dep_.CoverTab[8803]++
//line /snap/go/10455/src/net/udpsock_posix.go:229
		// _ = "end of CoverTab[8803]"
//line /snap/go/10455/src/net/udpsock_posix.go:229
	}
//line /snap/go/10455/src/net/udpsock_posix.go:229
	// _ = "end of CoverTab[8797]"
//line /snap/go/10455/src/net/udpsock_posix.go:229
	_go_fuzz_dep_.CoverTab[8798]++
							return newUDPConn(fd), nil
//line /snap/go/10455/src/net/udpsock_posix.go:230
	// _ = "end of CoverTab[8798]"
}

func (sl *sysListener) listenMulticastUDP(ctx context.Context, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error) {
//line /snap/go/10455/src/net/udpsock_posix.go:233
	_go_fuzz_dep_.CoverTab[8804]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:235
		_go_fuzz_dep_.CoverTab[530087]++
//line /snap/go/10455/src/net/udpsock_posix.go:235
		_go_fuzz_dep_.CoverTab[8808]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/udpsock_posix.go:236
			_go_fuzz_dep_.CoverTab[8809]++
									return sl.ListenConfig.Control(network, address, c)
//line /snap/go/10455/src/net/udpsock_posix.go:237
			// _ = "end of CoverTab[8809]"
		}
//line /snap/go/10455/src/net/udpsock_posix.go:238
		// _ = "end of CoverTab[8808]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:239
		_go_fuzz_dep_.CoverTab[530088]++
//line /snap/go/10455/src/net/udpsock_posix.go:239
		_go_fuzz_dep_.CoverTab[8810]++
//line /snap/go/10455/src/net/udpsock_posix.go:239
		// _ = "end of CoverTab[8810]"
//line /snap/go/10455/src/net/udpsock_posix.go:239
	}
//line /snap/go/10455/src/net/udpsock_posix.go:239
	// _ = "end of CoverTab[8804]"
//line /snap/go/10455/src/net/udpsock_posix.go:239
	_go_fuzz_dep_.CoverTab[8805]++
							fd, err := internetSocket(ctx, sl.network, gaddr, nil, syscall.SOCK_DGRAM, 0, "listen", ctrlCtxFn)
							if err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:241
		_go_fuzz_dep_.CoverTab[530089]++
//line /snap/go/10455/src/net/udpsock_posix.go:241
		_go_fuzz_dep_.CoverTab[8811]++
								return nil, err
//line /snap/go/10455/src/net/udpsock_posix.go:242
		// _ = "end of CoverTab[8811]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:243
		_go_fuzz_dep_.CoverTab[530090]++
//line /snap/go/10455/src/net/udpsock_posix.go:243
		_go_fuzz_dep_.CoverTab[8812]++
//line /snap/go/10455/src/net/udpsock_posix.go:243
		// _ = "end of CoverTab[8812]"
//line /snap/go/10455/src/net/udpsock_posix.go:243
	}
//line /snap/go/10455/src/net/udpsock_posix.go:243
	// _ = "end of CoverTab[8805]"
//line /snap/go/10455/src/net/udpsock_posix.go:243
	_go_fuzz_dep_.CoverTab[8806]++
							c := newUDPConn(fd)
							if ip4 := gaddr.IP.To4(); ip4 != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:245
		_go_fuzz_dep_.CoverTab[530091]++
//line /snap/go/10455/src/net/udpsock_posix.go:245
		_go_fuzz_dep_.CoverTab[8813]++
								if err := listenIPv4MulticastUDP(c, ifi, ip4); err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:246
			_go_fuzz_dep_.CoverTab[530093]++
//line /snap/go/10455/src/net/udpsock_posix.go:246
			_go_fuzz_dep_.CoverTab[8814]++
									c.Close()
									return nil, err
//line /snap/go/10455/src/net/udpsock_posix.go:248
			// _ = "end of CoverTab[8814]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:249
			_go_fuzz_dep_.CoverTab[530094]++
//line /snap/go/10455/src/net/udpsock_posix.go:249
			_go_fuzz_dep_.CoverTab[8815]++
//line /snap/go/10455/src/net/udpsock_posix.go:249
			// _ = "end of CoverTab[8815]"
//line /snap/go/10455/src/net/udpsock_posix.go:249
		}
//line /snap/go/10455/src/net/udpsock_posix.go:249
		// _ = "end of CoverTab[8813]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:250
		_go_fuzz_dep_.CoverTab[530092]++
//line /snap/go/10455/src/net/udpsock_posix.go:250
		_go_fuzz_dep_.CoverTab[8816]++
								if err := listenIPv6MulticastUDP(c, ifi, gaddr.IP); err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:251
			_go_fuzz_dep_.CoverTab[530095]++
//line /snap/go/10455/src/net/udpsock_posix.go:251
			_go_fuzz_dep_.CoverTab[8817]++
									c.Close()
									return nil, err
//line /snap/go/10455/src/net/udpsock_posix.go:253
			// _ = "end of CoverTab[8817]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:254
			_go_fuzz_dep_.CoverTab[530096]++
//line /snap/go/10455/src/net/udpsock_posix.go:254
			_go_fuzz_dep_.CoverTab[8818]++
//line /snap/go/10455/src/net/udpsock_posix.go:254
			// _ = "end of CoverTab[8818]"
//line /snap/go/10455/src/net/udpsock_posix.go:254
		}
//line /snap/go/10455/src/net/udpsock_posix.go:254
		// _ = "end of CoverTab[8816]"
	}
//line /snap/go/10455/src/net/udpsock_posix.go:255
	// _ = "end of CoverTab[8806]"
//line /snap/go/10455/src/net/udpsock_posix.go:255
	_go_fuzz_dep_.CoverTab[8807]++
							return c, nil
//line /snap/go/10455/src/net/udpsock_posix.go:256
	// _ = "end of CoverTab[8807]"
}

func listenIPv4MulticastUDP(c *UDPConn, ifi *Interface, ip IP) error {
//line /snap/go/10455/src/net/udpsock_posix.go:259
	_go_fuzz_dep_.CoverTab[8819]++
							if ifi != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:260
		_go_fuzz_dep_.CoverTab[530097]++
//line /snap/go/10455/src/net/udpsock_posix.go:260
		_go_fuzz_dep_.CoverTab[8823]++
								if err := setIPv4MulticastInterface(c.fd, ifi); err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:261
			_go_fuzz_dep_.CoverTab[530099]++
//line /snap/go/10455/src/net/udpsock_posix.go:261
			_go_fuzz_dep_.CoverTab[8824]++
									return err
//line /snap/go/10455/src/net/udpsock_posix.go:262
			// _ = "end of CoverTab[8824]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:263
			_go_fuzz_dep_.CoverTab[530100]++
//line /snap/go/10455/src/net/udpsock_posix.go:263
			_go_fuzz_dep_.CoverTab[8825]++
//line /snap/go/10455/src/net/udpsock_posix.go:263
			// _ = "end of CoverTab[8825]"
//line /snap/go/10455/src/net/udpsock_posix.go:263
		}
//line /snap/go/10455/src/net/udpsock_posix.go:263
		// _ = "end of CoverTab[8823]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:264
		_go_fuzz_dep_.CoverTab[530098]++
//line /snap/go/10455/src/net/udpsock_posix.go:264
		_go_fuzz_dep_.CoverTab[8826]++
//line /snap/go/10455/src/net/udpsock_posix.go:264
		// _ = "end of CoverTab[8826]"
//line /snap/go/10455/src/net/udpsock_posix.go:264
	}
//line /snap/go/10455/src/net/udpsock_posix.go:264
	// _ = "end of CoverTab[8819]"
//line /snap/go/10455/src/net/udpsock_posix.go:264
	_go_fuzz_dep_.CoverTab[8820]++
							if err := setIPv4MulticastLoopback(c.fd, false); err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:265
		_go_fuzz_dep_.CoverTab[530101]++
//line /snap/go/10455/src/net/udpsock_posix.go:265
		_go_fuzz_dep_.CoverTab[8827]++
								return err
//line /snap/go/10455/src/net/udpsock_posix.go:266
		// _ = "end of CoverTab[8827]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:267
		_go_fuzz_dep_.CoverTab[530102]++
//line /snap/go/10455/src/net/udpsock_posix.go:267
		_go_fuzz_dep_.CoverTab[8828]++
//line /snap/go/10455/src/net/udpsock_posix.go:267
		// _ = "end of CoverTab[8828]"
//line /snap/go/10455/src/net/udpsock_posix.go:267
	}
//line /snap/go/10455/src/net/udpsock_posix.go:267
	// _ = "end of CoverTab[8820]"
//line /snap/go/10455/src/net/udpsock_posix.go:267
	_go_fuzz_dep_.CoverTab[8821]++
							if err := joinIPv4Group(c.fd, ifi, ip); err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:268
		_go_fuzz_dep_.CoverTab[530103]++
//line /snap/go/10455/src/net/udpsock_posix.go:268
		_go_fuzz_dep_.CoverTab[8829]++
								return err
//line /snap/go/10455/src/net/udpsock_posix.go:269
		// _ = "end of CoverTab[8829]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:270
		_go_fuzz_dep_.CoverTab[530104]++
//line /snap/go/10455/src/net/udpsock_posix.go:270
		_go_fuzz_dep_.CoverTab[8830]++
//line /snap/go/10455/src/net/udpsock_posix.go:270
		// _ = "end of CoverTab[8830]"
//line /snap/go/10455/src/net/udpsock_posix.go:270
	}
//line /snap/go/10455/src/net/udpsock_posix.go:270
	// _ = "end of CoverTab[8821]"
//line /snap/go/10455/src/net/udpsock_posix.go:270
	_go_fuzz_dep_.CoverTab[8822]++
							return nil
//line /snap/go/10455/src/net/udpsock_posix.go:271
	// _ = "end of CoverTab[8822]"
}

func listenIPv6MulticastUDP(c *UDPConn, ifi *Interface, ip IP) error {
//line /snap/go/10455/src/net/udpsock_posix.go:274
	_go_fuzz_dep_.CoverTab[8831]++
							if ifi != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:275
		_go_fuzz_dep_.CoverTab[530105]++
//line /snap/go/10455/src/net/udpsock_posix.go:275
		_go_fuzz_dep_.CoverTab[8835]++
								if err := setIPv6MulticastInterface(c.fd, ifi); err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:276
			_go_fuzz_dep_.CoverTab[530107]++
//line /snap/go/10455/src/net/udpsock_posix.go:276
			_go_fuzz_dep_.CoverTab[8836]++
									return err
//line /snap/go/10455/src/net/udpsock_posix.go:277
			// _ = "end of CoverTab[8836]"
		} else {
//line /snap/go/10455/src/net/udpsock_posix.go:278
			_go_fuzz_dep_.CoverTab[530108]++
//line /snap/go/10455/src/net/udpsock_posix.go:278
			_go_fuzz_dep_.CoverTab[8837]++
//line /snap/go/10455/src/net/udpsock_posix.go:278
			// _ = "end of CoverTab[8837]"
//line /snap/go/10455/src/net/udpsock_posix.go:278
		}
//line /snap/go/10455/src/net/udpsock_posix.go:278
		// _ = "end of CoverTab[8835]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:279
		_go_fuzz_dep_.CoverTab[530106]++
//line /snap/go/10455/src/net/udpsock_posix.go:279
		_go_fuzz_dep_.CoverTab[8838]++
//line /snap/go/10455/src/net/udpsock_posix.go:279
		// _ = "end of CoverTab[8838]"
//line /snap/go/10455/src/net/udpsock_posix.go:279
	}
//line /snap/go/10455/src/net/udpsock_posix.go:279
	// _ = "end of CoverTab[8831]"
//line /snap/go/10455/src/net/udpsock_posix.go:279
	_go_fuzz_dep_.CoverTab[8832]++
							if err := setIPv6MulticastLoopback(c.fd, false); err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:280
		_go_fuzz_dep_.CoverTab[530109]++
//line /snap/go/10455/src/net/udpsock_posix.go:280
		_go_fuzz_dep_.CoverTab[8839]++
								return err
//line /snap/go/10455/src/net/udpsock_posix.go:281
		// _ = "end of CoverTab[8839]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:282
		_go_fuzz_dep_.CoverTab[530110]++
//line /snap/go/10455/src/net/udpsock_posix.go:282
		_go_fuzz_dep_.CoverTab[8840]++
//line /snap/go/10455/src/net/udpsock_posix.go:282
		// _ = "end of CoverTab[8840]"
//line /snap/go/10455/src/net/udpsock_posix.go:282
	}
//line /snap/go/10455/src/net/udpsock_posix.go:282
	// _ = "end of CoverTab[8832]"
//line /snap/go/10455/src/net/udpsock_posix.go:282
	_go_fuzz_dep_.CoverTab[8833]++
							if err := joinIPv6Group(c.fd, ifi, ip); err != nil {
//line /snap/go/10455/src/net/udpsock_posix.go:283
		_go_fuzz_dep_.CoverTab[530111]++
//line /snap/go/10455/src/net/udpsock_posix.go:283
		_go_fuzz_dep_.CoverTab[8841]++
								return err
//line /snap/go/10455/src/net/udpsock_posix.go:284
		// _ = "end of CoverTab[8841]"
	} else {
//line /snap/go/10455/src/net/udpsock_posix.go:285
		_go_fuzz_dep_.CoverTab[530112]++
//line /snap/go/10455/src/net/udpsock_posix.go:285
		_go_fuzz_dep_.CoverTab[8842]++
//line /snap/go/10455/src/net/udpsock_posix.go:285
		// _ = "end of CoverTab[8842]"
//line /snap/go/10455/src/net/udpsock_posix.go:285
	}
//line /snap/go/10455/src/net/udpsock_posix.go:285
	// _ = "end of CoverTab[8833]"
//line /snap/go/10455/src/net/udpsock_posix.go:285
	_go_fuzz_dep_.CoverTab[8834]++
							return nil
//line /snap/go/10455/src/net/udpsock_posix.go:286
	// _ = "end of CoverTab[8834]"
}

//line /snap/go/10455/src/net/udpsock_posix.go:287
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/udpsock_posix.go:287
var _ = _go_fuzz_dep_.CoverTab
