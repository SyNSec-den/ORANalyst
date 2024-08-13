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
	_go_fuzz_dep_.CoverTab[8379]++
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /usr/local/go/src/net/udpsock_posix.go:17
		_go_fuzz_dep_.CoverTab[8381]++
								return &UDPAddr{IP: sa.Addr[0:], Port: sa.Port}
//line /usr/local/go/src/net/udpsock_posix.go:18
		// _ = "end of CoverTab[8381]"
	case *syscall.SockaddrInet6:
//line /usr/local/go/src/net/udpsock_posix.go:19
		_go_fuzz_dep_.CoverTab[8382]++
								return &UDPAddr{IP: sa.Addr[0:], Port: sa.Port, Zone: zoneCache.name(int(sa.ZoneId))}
//line /usr/local/go/src/net/udpsock_posix.go:20
		// _ = "end of CoverTab[8382]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:21
	// _ = "end of CoverTab[8379]"
//line /usr/local/go/src/net/udpsock_posix.go:21
	_go_fuzz_dep_.CoverTab[8380]++
							return nil
//line /usr/local/go/src/net/udpsock_posix.go:22
	// _ = "end of CoverTab[8380]"
}

func (a *UDPAddr) family() int {
//line /usr/local/go/src/net/udpsock_posix.go:25
	_go_fuzz_dep_.CoverTab[8383]++
							if a == nil || func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:26
		_go_fuzz_dep_.CoverTab[8386]++
//line /usr/local/go/src/net/udpsock_posix.go:26
		return len(a.IP) <= IPv4len
//line /usr/local/go/src/net/udpsock_posix.go:26
		// _ = "end of CoverTab[8386]"
//line /usr/local/go/src/net/udpsock_posix.go:26
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:26
		_go_fuzz_dep_.CoverTab[8387]++
								return syscall.AF_INET
//line /usr/local/go/src/net/udpsock_posix.go:27
		// _ = "end of CoverTab[8387]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:28
		_go_fuzz_dep_.CoverTab[8388]++
//line /usr/local/go/src/net/udpsock_posix.go:28
		// _ = "end of CoverTab[8388]"
//line /usr/local/go/src/net/udpsock_posix.go:28
	}
//line /usr/local/go/src/net/udpsock_posix.go:28
	// _ = "end of CoverTab[8383]"
//line /usr/local/go/src/net/udpsock_posix.go:28
	_go_fuzz_dep_.CoverTab[8384]++
							if a.IP.To4() != nil {
//line /usr/local/go/src/net/udpsock_posix.go:29
		_go_fuzz_dep_.CoverTab[8389]++
								return syscall.AF_INET
//line /usr/local/go/src/net/udpsock_posix.go:30
		// _ = "end of CoverTab[8389]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:31
		_go_fuzz_dep_.CoverTab[8390]++
//line /usr/local/go/src/net/udpsock_posix.go:31
		// _ = "end of CoverTab[8390]"
//line /usr/local/go/src/net/udpsock_posix.go:31
	}
//line /usr/local/go/src/net/udpsock_posix.go:31
	// _ = "end of CoverTab[8384]"
//line /usr/local/go/src/net/udpsock_posix.go:31
	_go_fuzz_dep_.CoverTab[8385]++
							return syscall.AF_INET6
//line /usr/local/go/src/net/udpsock_posix.go:32
	// _ = "end of CoverTab[8385]"
}

func (a *UDPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/udpsock_posix.go:35
	_go_fuzz_dep_.CoverTab[8391]++
							if a == nil {
//line /usr/local/go/src/net/udpsock_posix.go:36
		_go_fuzz_dep_.CoverTab[8393]++
								return nil, nil
//line /usr/local/go/src/net/udpsock_posix.go:37
		// _ = "end of CoverTab[8393]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:38
		_go_fuzz_dep_.CoverTab[8394]++
//line /usr/local/go/src/net/udpsock_posix.go:38
		// _ = "end of CoverTab[8394]"
//line /usr/local/go/src/net/udpsock_posix.go:38
	}
//line /usr/local/go/src/net/udpsock_posix.go:38
	// _ = "end of CoverTab[8391]"
//line /usr/local/go/src/net/udpsock_posix.go:38
	_go_fuzz_dep_.CoverTab[8392]++
							return ipToSockaddr(family, a.IP, a.Port, a.Zone)
//line /usr/local/go/src/net/udpsock_posix.go:39
	// _ = "end of CoverTab[8392]"
}

func (a *UDPAddr) toLocal(net string) sockaddr {
//line /usr/local/go/src/net/udpsock_posix.go:42
	_go_fuzz_dep_.CoverTab[8395]++
							return &UDPAddr{loopbackIP(net), a.Port, a.Zone}
//line /usr/local/go/src/net/udpsock_posix.go:43
	// _ = "end of CoverTab[8395]"
}

func (c *UDPConn) readFrom(b []byte, addr *UDPAddr) (int, *UDPAddr, error) {
//line /usr/local/go/src/net/udpsock_posix.go:46
	_go_fuzz_dep_.CoverTab[8396]++
							var n int
							var err error
							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:50
		_go_fuzz_dep_.CoverTab[8399]++
								var from syscall.SockaddrInet4
								n, err = c.fd.readFromInet4(b, &from)
								if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:53
			_go_fuzz_dep_.CoverTab[8402]++
									ip := from.Addr
									*addr = UDPAddr{IP: ip[:], Port: from.Port}
//line /usr/local/go/src/net/udpsock_posix.go:55
			// _ = "end of CoverTab[8402]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:56
			_go_fuzz_dep_.CoverTab[8403]++
//line /usr/local/go/src/net/udpsock_posix.go:56
			// _ = "end of CoverTab[8403]"
//line /usr/local/go/src/net/udpsock_posix.go:56
		}
//line /usr/local/go/src/net/udpsock_posix.go:56
		// _ = "end of CoverTab[8399]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:57
		_go_fuzz_dep_.CoverTab[8400]++
								var from syscall.SockaddrInet6
								n, err = c.fd.readFromInet6(b, &from)
								if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:60
			_go_fuzz_dep_.CoverTab[8404]++
									ip := from.Addr
									*addr = UDPAddr{IP: ip[:], Port: from.Port, Zone: zoneCache.name(int(from.ZoneId))}
//line /usr/local/go/src/net/udpsock_posix.go:62
			// _ = "end of CoverTab[8404]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:63
			_go_fuzz_dep_.CoverTab[8405]++
//line /usr/local/go/src/net/udpsock_posix.go:63
			// _ = "end of CoverTab[8405]"
//line /usr/local/go/src/net/udpsock_posix.go:63
		}
//line /usr/local/go/src/net/udpsock_posix.go:63
		// _ = "end of CoverTab[8400]"
//line /usr/local/go/src/net/udpsock_posix.go:63
	default:
//line /usr/local/go/src/net/udpsock_posix.go:63
		_go_fuzz_dep_.CoverTab[8401]++
//line /usr/local/go/src/net/udpsock_posix.go:63
		// _ = "end of CoverTab[8401]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:64
	// _ = "end of CoverTab[8396]"
//line /usr/local/go/src/net/udpsock_posix.go:64
	_go_fuzz_dep_.CoverTab[8397]++
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:65
		_go_fuzz_dep_.CoverTab[8406]++

								addr = nil
//line /usr/local/go/src/net/udpsock_posix.go:67
		// _ = "end of CoverTab[8406]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:68
		_go_fuzz_dep_.CoverTab[8407]++
//line /usr/local/go/src/net/udpsock_posix.go:68
		// _ = "end of CoverTab[8407]"
//line /usr/local/go/src/net/udpsock_posix.go:68
	}
//line /usr/local/go/src/net/udpsock_posix.go:68
	// _ = "end of CoverTab[8397]"
//line /usr/local/go/src/net/udpsock_posix.go:68
	_go_fuzz_dep_.CoverTab[8398]++
							return n, addr, err
//line /usr/local/go/src/net/udpsock_posix.go:69
	// _ = "end of CoverTab[8398]"
}

func (c *UDPConn) readFromAddrPort(b []byte) (n int, addr netip.AddrPort, err error) {
//line /usr/local/go/src/net/udpsock_posix.go:72
	_go_fuzz_dep_.CoverTab[8408]++
							var ip netip.Addr
							var port int
							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:76
		_go_fuzz_dep_.CoverTab[8411]++
								var from syscall.SockaddrInet4
								n, err = c.fd.readFromInet4(b, &from)
								if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:79
			_go_fuzz_dep_.CoverTab[8414]++
									ip = netip.AddrFrom4(from.Addr)
									port = from.Port
//line /usr/local/go/src/net/udpsock_posix.go:81
			// _ = "end of CoverTab[8414]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:82
			_go_fuzz_dep_.CoverTab[8415]++
//line /usr/local/go/src/net/udpsock_posix.go:82
			// _ = "end of CoverTab[8415]"
//line /usr/local/go/src/net/udpsock_posix.go:82
		}
//line /usr/local/go/src/net/udpsock_posix.go:82
		// _ = "end of CoverTab[8411]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:83
		_go_fuzz_dep_.CoverTab[8412]++
								var from syscall.SockaddrInet6
								n, err = c.fd.readFromInet6(b, &from)
								if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:86
			_go_fuzz_dep_.CoverTab[8416]++
									ip = netip.AddrFrom16(from.Addr).WithZone(zoneCache.name(int(from.ZoneId)))
									port = from.Port
//line /usr/local/go/src/net/udpsock_posix.go:88
			// _ = "end of CoverTab[8416]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:89
			_go_fuzz_dep_.CoverTab[8417]++
//line /usr/local/go/src/net/udpsock_posix.go:89
			// _ = "end of CoverTab[8417]"
//line /usr/local/go/src/net/udpsock_posix.go:89
		}
//line /usr/local/go/src/net/udpsock_posix.go:89
		// _ = "end of CoverTab[8412]"
//line /usr/local/go/src/net/udpsock_posix.go:89
	default:
//line /usr/local/go/src/net/udpsock_posix.go:89
		_go_fuzz_dep_.CoverTab[8413]++
//line /usr/local/go/src/net/udpsock_posix.go:89
		// _ = "end of CoverTab[8413]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:90
	// _ = "end of CoverTab[8408]"
//line /usr/local/go/src/net/udpsock_posix.go:90
	_go_fuzz_dep_.CoverTab[8409]++
							if err == nil {
//line /usr/local/go/src/net/udpsock_posix.go:91
		_go_fuzz_dep_.CoverTab[8418]++
								addr = netip.AddrPortFrom(ip, uint16(port))
//line /usr/local/go/src/net/udpsock_posix.go:92
		// _ = "end of CoverTab[8418]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:93
		_go_fuzz_dep_.CoverTab[8419]++
//line /usr/local/go/src/net/udpsock_posix.go:93
		// _ = "end of CoverTab[8419]"
//line /usr/local/go/src/net/udpsock_posix.go:93
	}
//line /usr/local/go/src/net/udpsock_posix.go:93
	// _ = "end of CoverTab[8409]"
//line /usr/local/go/src/net/udpsock_posix.go:93
	_go_fuzz_dep_.CoverTab[8410]++
							return n, addr, err
//line /usr/local/go/src/net/udpsock_posix.go:94
	// _ = "end of CoverTab[8410]"
}

func (c *UDPConn) readMsg(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error) {
//line /usr/local/go/src/net/udpsock_posix.go:97
	_go_fuzz_dep_.CoverTab[8420]++
							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:99
		_go_fuzz_dep_.CoverTab[8422]++
								var sa syscall.SockaddrInet4
								n, oobn, flags, err = c.fd.readMsgInet4(b, oob, 0, &sa)
								ip := netip.AddrFrom4(sa.Addr)
								addr = netip.AddrPortFrom(ip, uint16(sa.Port))
//line /usr/local/go/src/net/udpsock_posix.go:103
		// _ = "end of CoverTab[8422]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:104
		_go_fuzz_dep_.CoverTab[8423]++
								var sa syscall.SockaddrInet6
								n, oobn, flags, err = c.fd.readMsgInet6(b, oob, 0, &sa)
								ip := netip.AddrFrom16(sa.Addr).WithZone(zoneCache.name(int(sa.ZoneId)))
								addr = netip.AddrPortFrom(ip, uint16(sa.Port))
//line /usr/local/go/src/net/udpsock_posix.go:108
		// _ = "end of CoverTab[8423]"
//line /usr/local/go/src/net/udpsock_posix.go:108
	default:
//line /usr/local/go/src/net/udpsock_posix.go:108
		_go_fuzz_dep_.CoverTab[8424]++
//line /usr/local/go/src/net/udpsock_posix.go:108
		// _ = "end of CoverTab[8424]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:109
	// _ = "end of CoverTab[8420]"
//line /usr/local/go/src/net/udpsock_posix.go:109
	_go_fuzz_dep_.CoverTab[8421]++
							return
//line /usr/local/go/src/net/udpsock_posix.go:110
	// _ = "end of CoverTab[8421]"
}

func (c *UDPConn) writeTo(b []byte, addr *UDPAddr) (int, error) {
//line /usr/local/go/src/net/udpsock_posix.go:113
	_go_fuzz_dep_.CoverTab[8425]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/udpsock_posix.go:114
		_go_fuzz_dep_.CoverTab[8428]++
								return 0, ErrWriteToConnected
//line /usr/local/go/src/net/udpsock_posix.go:115
		// _ = "end of CoverTab[8428]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:116
		_go_fuzz_dep_.CoverTab[8429]++
//line /usr/local/go/src/net/udpsock_posix.go:116
		// _ = "end of CoverTab[8429]"
//line /usr/local/go/src/net/udpsock_posix.go:116
	}
//line /usr/local/go/src/net/udpsock_posix.go:116
	// _ = "end of CoverTab[8425]"
//line /usr/local/go/src/net/udpsock_posix.go:116
	_go_fuzz_dep_.CoverTab[8426]++
							if addr == nil {
//line /usr/local/go/src/net/udpsock_posix.go:117
		_go_fuzz_dep_.CoverTab[8430]++
								return 0, errMissingAddress
//line /usr/local/go/src/net/udpsock_posix.go:118
		// _ = "end of CoverTab[8430]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:119
		_go_fuzz_dep_.CoverTab[8431]++
//line /usr/local/go/src/net/udpsock_posix.go:119
		// _ = "end of CoverTab[8431]"
//line /usr/local/go/src/net/udpsock_posix.go:119
	}
//line /usr/local/go/src/net/udpsock_posix.go:119
	// _ = "end of CoverTab[8426]"
//line /usr/local/go/src/net/udpsock_posix.go:119
	_go_fuzz_dep_.CoverTab[8427]++

							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:122
		_go_fuzz_dep_.CoverTab[8432]++
								sa, err := ipToSockaddrInet4(addr.IP, addr.Port)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:124
			_go_fuzz_dep_.CoverTab[8437]++
									return 0, err
//line /usr/local/go/src/net/udpsock_posix.go:125
			// _ = "end of CoverTab[8437]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:126
			_go_fuzz_dep_.CoverTab[8438]++
//line /usr/local/go/src/net/udpsock_posix.go:126
			// _ = "end of CoverTab[8438]"
//line /usr/local/go/src/net/udpsock_posix.go:126
		}
//line /usr/local/go/src/net/udpsock_posix.go:126
		// _ = "end of CoverTab[8432]"
//line /usr/local/go/src/net/udpsock_posix.go:126
		_go_fuzz_dep_.CoverTab[8433]++
								return c.fd.writeToInet4(b, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:127
		// _ = "end of CoverTab[8433]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:128
		_go_fuzz_dep_.CoverTab[8434]++
								sa, err := ipToSockaddrInet6(addr.IP, addr.Port, addr.Zone)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:130
			_go_fuzz_dep_.CoverTab[8439]++
									return 0, err
//line /usr/local/go/src/net/udpsock_posix.go:131
			// _ = "end of CoverTab[8439]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:132
			_go_fuzz_dep_.CoverTab[8440]++
//line /usr/local/go/src/net/udpsock_posix.go:132
			// _ = "end of CoverTab[8440]"
//line /usr/local/go/src/net/udpsock_posix.go:132
		}
//line /usr/local/go/src/net/udpsock_posix.go:132
		// _ = "end of CoverTab[8434]"
//line /usr/local/go/src/net/udpsock_posix.go:132
		_go_fuzz_dep_.CoverTab[8435]++
								return c.fd.writeToInet6(b, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:133
		// _ = "end of CoverTab[8435]"
	default:
//line /usr/local/go/src/net/udpsock_posix.go:134
		_go_fuzz_dep_.CoverTab[8436]++
								return 0, &AddrError{Err: "invalid address family", Addr: addr.IP.String()}
//line /usr/local/go/src/net/udpsock_posix.go:135
		// _ = "end of CoverTab[8436]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:136
	// _ = "end of CoverTab[8427]"
}

func (c *UDPConn) writeToAddrPort(b []byte, addr netip.AddrPort) (int, error) {
//line /usr/local/go/src/net/udpsock_posix.go:139
	_go_fuzz_dep_.CoverTab[8441]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/udpsock_posix.go:140
		_go_fuzz_dep_.CoverTab[8444]++
								return 0, ErrWriteToConnected
//line /usr/local/go/src/net/udpsock_posix.go:141
		// _ = "end of CoverTab[8444]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:142
		_go_fuzz_dep_.CoverTab[8445]++
//line /usr/local/go/src/net/udpsock_posix.go:142
		// _ = "end of CoverTab[8445]"
//line /usr/local/go/src/net/udpsock_posix.go:142
	}
//line /usr/local/go/src/net/udpsock_posix.go:142
	// _ = "end of CoverTab[8441]"
//line /usr/local/go/src/net/udpsock_posix.go:142
	_go_fuzz_dep_.CoverTab[8442]++
							if !addr.IsValid() {
//line /usr/local/go/src/net/udpsock_posix.go:143
		_go_fuzz_dep_.CoverTab[8446]++
								return 0, errMissingAddress
//line /usr/local/go/src/net/udpsock_posix.go:144
		// _ = "end of CoverTab[8446]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:145
		_go_fuzz_dep_.CoverTab[8447]++
//line /usr/local/go/src/net/udpsock_posix.go:145
		// _ = "end of CoverTab[8447]"
//line /usr/local/go/src/net/udpsock_posix.go:145
	}
//line /usr/local/go/src/net/udpsock_posix.go:145
	// _ = "end of CoverTab[8442]"
//line /usr/local/go/src/net/udpsock_posix.go:145
	_go_fuzz_dep_.CoverTab[8443]++

							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:148
		_go_fuzz_dep_.CoverTab[8448]++
								sa, err := addrPortToSockaddrInet4(addr)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:150
			_go_fuzz_dep_.CoverTab[8453]++
									return 0, err
//line /usr/local/go/src/net/udpsock_posix.go:151
			// _ = "end of CoverTab[8453]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:152
			_go_fuzz_dep_.CoverTab[8454]++
//line /usr/local/go/src/net/udpsock_posix.go:152
			// _ = "end of CoverTab[8454]"
//line /usr/local/go/src/net/udpsock_posix.go:152
		}
//line /usr/local/go/src/net/udpsock_posix.go:152
		// _ = "end of CoverTab[8448]"
//line /usr/local/go/src/net/udpsock_posix.go:152
		_go_fuzz_dep_.CoverTab[8449]++
								return c.fd.writeToInet4(b, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:153
		// _ = "end of CoverTab[8449]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:154
		_go_fuzz_dep_.CoverTab[8450]++
								sa, err := addrPortToSockaddrInet6(addr)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:156
			_go_fuzz_dep_.CoverTab[8455]++
									return 0, err
//line /usr/local/go/src/net/udpsock_posix.go:157
			// _ = "end of CoverTab[8455]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:158
			_go_fuzz_dep_.CoverTab[8456]++
//line /usr/local/go/src/net/udpsock_posix.go:158
			// _ = "end of CoverTab[8456]"
//line /usr/local/go/src/net/udpsock_posix.go:158
		}
//line /usr/local/go/src/net/udpsock_posix.go:158
		// _ = "end of CoverTab[8450]"
//line /usr/local/go/src/net/udpsock_posix.go:158
		_go_fuzz_dep_.CoverTab[8451]++
								return c.fd.writeToInet6(b, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:159
		// _ = "end of CoverTab[8451]"
	default:
//line /usr/local/go/src/net/udpsock_posix.go:160
		_go_fuzz_dep_.CoverTab[8452]++
								return 0, &AddrError{Err: "invalid address family", Addr: addr.Addr().String()}
//line /usr/local/go/src/net/udpsock_posix.go:161
		// _ = "end of CoverTab[8452]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:162
	// _ = "end of CoverTab[8443]"
}

func (c *UDPConn) writeMsg(b, oob []byte, addr *UDPAddr) (n, oobn int, err error) {
//line /usr/local/go/src/net/udpsock_posix.go:165
	_go_fuzz_dep_.CoverTab[8457]++
							if c.fd.isConnected && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:166
		_go_fuzz_dep_.CoverTab[8461]++
//line /usr/local/go/src/net/udpsock_posix.go:166
		return addr != nil
//line /usr/local/go/src/net/udpsock_posix.go:166
		// _ = "end of CoverTab[8461]"
//line /usr/local/go/src/net/udpsock_posix.go:166
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:166
		_go_fuzz_dep_.CoverTab[8462]++
								return 0, 0, ErrWriteToConnected
//line /usr/local/go/src/net/udpsock_posix.go:167
		// _ = "end of CoverTab[8462]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:168
		_go_fuzz_dep_.CoverTab[8463]++
//line /usr/local/go/src/net/udpsock_posix.go:168
		// _ = "end of CoverTab[8463]"
//line /usr/local/go/src/net/udpsock_posix.go:168
	}
//line /usr/local/go/src/net/udpsock_posix.go:168
	// _ = "end of CoverTab[8457]"
//line /usr/local/go/src/net/udpsock_posix.go:168
	_go_fuzz_dep_.CoverTab[8458]++
							if !c.fd.isConnected && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:169
		_go_fuzz_dep_.CoverTab[8464]++
//line /usr/local/go/src/net/udpsock_posix.go:169
		return addr == nil
//line /usr/local/go/src/net/udpsock_posix.go:169
		// _ = "end of CoverTab[8464]"
//line /usr/local/go/src/net/udpsock_posix.go:169
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:169
		_go_fuzz_dep_.CoverTab[8465]++
								return 0, 0, errMissingAddress
//line /usr/local/go/src/net/udpsock_posix.go:170
		// _ = "end of CoverTab[8465]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:171
		_go_fuzz_dep_.CoverTab[8466]++
//line /usr/local/go/src/net/udpsock_posix.go:171
		// _ = "end of CoverTab[8466]"
//line /usr/local/go/src/net/udpsock_posix.go:171
	}
//line /usr/local/go/src/net/udpsock_posix.go:171
	// _ = "end of CoverTab[8458]"
//line /usr/local/go/src/net/udpsock_posix.go:171
	_go_fuzz_dep_.CoverTab[8459]++
							sa, err := addr.sockaddr(c.fd.family)
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:173
		_go_fuzz_dep_.CoverTab[8467]++
								return 0, 0, err
//line /usr/local/go/src/net/udpsock_posix.go:174
		// _ = "end of CoverTab[8467]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:175
		_go_fuzz_dep_.CoverTab[8468]++
//line /usr/local/go/src/net/udpsock_posix.go:175
		// _ = "end of CoverTab[8468]"
//line /usr/local/go/src/net/udpsock_posix.go:175
	}
//line /usr/local/go/src/net/udpsock_posix.go:175
	// _ = "end of CoverTab[8459]"
//line /usr/local/go/src/net/udpsock_posix.go:175
	_go_fuzz_dep_.CoverTab[8460]++
							return c.fd.writeMsg(b, oob, sa)
//line /usr/local/go/src/net/udpsock_posix.go:176
	// _ = "end of CoverTab[8460]"
}

func (c *UDPConn) writeMsgAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error) {
//line /usr/local/go/src/net/udpsock_posix.go:179
	_go_fuzz_dep_.CoverTab[8469]++
							if c.fd.isConnected && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:180
		_go_fuzz_dep_.CoverTab[8472]++
//line /usr/local/go/src/net/udpsock_posix.go:180
		return addr.IsValid()
//line /usr/local/go/src/net/udpsock_posix.go:180
		// _ = "end of CoverTab[8472]"
//line /usr/local/go/src/net/udpsock_posix.go:180
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:180
		_go_fuzz_dep_.CoverTab[8473]++
								return 0, 0, ErrWriteToConnected
//line /usr/local/go/src/net/udpsock_posix.go:181
		// _ = "end of CoverTab[8473]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:182
		_go_fuzz_dep_.CoverTab[8474]++
//line /usr/local/go/src/net/udpsock_posix.go:182
		// _ = "end of CoverTab[8474]"
//line /usr/local/go/src/net/udpsock_posix.go:182
	}
//line /usr/local/go/src/net/udpsock_posix.go:182
	// _ = "end of CoverTab[8469]"
//line /usr/local/go/src/net/udpsock_posix.go:182
	_go_fuzz_dep_.CoverTab[8470]++
							if !c.fd.isConnected && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:183
		_go_fuzz_dep_.CoverTab[8475]++
//line /usr/local/go/src/net/udpsock_posix.go:183
		return !addr.IsValid()
//line /usr/local/go/src/net/udpsock_posix.go:183
		// _ = "end of CoverTab[8475]"
//line /usr/local/go/src/net/udpsock_posix.go:183
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:183
		_go_fuzz_dep_.CoverTab[8476]++
								return 0, 0, errMissingAddress
//line /usr/local/go/src/net/udpsock_posix.go:184
		// _ = "end of CoverTab[8476]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:185
		_go_fuzz_dep_.CoverTab[8477]++
//line /usr/local/go/src/net/udpsock_posix.go:185
		// _ = "end of CoverTab[8477]"
//line /usr/local/go/src/net/udpsock_posix.go:185
	}
//line /usr/local/go/src/net/udpsock_posix.go:185
	// _ = "end of CoverTab[8470]"
//line /usr/local/go/src/net/udpsock_posix.go:185
	_go_fuzz_dep_.CoverTab[8471]++

							switch c.fd.family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/udpsock_posix.go:188
		_go_fuzz_dep_.CoverTab[8478]++
								sa, err := addrPortToSockaddrInet4(addr)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:190
			_go_fuzz_dep_.CoverTab[8483]++
									return 0, 0, err
//line /usr/local/go/src/net/udpsock_posix.go:191
			// _ = "end of CoverTab[8483]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:192
			_go_fuzz_dep_.CoverTab[8484]++
//line /usr/local/go/src/net/udpsock_posix.go:192
			// _ = "end of CoverTab[8484]"
//line /usr/local/go/src/net/udpsock_posix.go:192
		}
//line /usr/local/go/src/net/udpsock_posix.go:192
		// _ = "end of CoverTab[8478]"
//line /usr/local/go/src/net/udpsock_posix.go:192
		_go_fuzz_dep_.CoverTab[8479]++
								return c.fd.writeMsgInet4(b, oob, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:193
		// _ = "end of CoverTab[8479]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/udpsock_posix.go:194
		_go_fuzz_dep_.CoverTab[8480]++
								sa, err := addrPortToSockaddrInet6(addr)
								if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:196
			_go_fuzz_dep_.CoverTab[8485]++
									return 0, 0, err
//line /usr/local/go/src/net/udpsock_posix.go:197
			// _ = "end of CoverTab[8485]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:198
			_go_fuzz_dep_.CoverTab[8486]++
//line /usr/local/go/src/net/udpsock_posix.go:198
			// _ = "end of CoverTab[8486]"
//line /usr/local/go/src/net/udpsock_posix.go:198
		}
//line /usr/local/go/src/net/udpsock_posix.go:198
		// _ = "end of CoverTab[8480]"
//line /usr/local/go/src/net/udpsock_posix.go:198
		_go_fuzz_dep_.CoverTab[8481]++
								return c.fd.writeMsgInet6(b, oob, &sa)
//line /usr/local/go/src/net/udpsock_posix.go:199
		// _ = "end of CoverTab[8481]"
	default:
//line /usr/local/go/src/net/udpsock_posix.go:200
		_go_fuzz_dep_.CoverTab[8482]++
								return 0, 0, &AddrError{Err: "invalid address family", Addr: addr.Addr().String()}
//line /usr/local/go/src/net/udpsock_posix.go:201
		// _ = "end of CoverTab[8482]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:202
	// _ = "end of CoverTab[8471]"
}

func (sd *sysDialer) dialUDP(ctx context.Context, laddr, raddr *UDPAddr) (*UDPConn, error) {
//line /usr/local/go/src/net/udpsock_posix.go:205
	_go_fuzz_dep_.CoverTab[8487]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /usr/local/go/src/net/udpsock_posix.go:207
		_go_fuzz_dep_.CoverTab[8490]++
//line /usr/local/go/src/net/udpsock_posix.go:207
		return sd.Dialer.Control != nil
//line /usr/local/go/src/net/udpsock_posix.go:207
		// _ = "end of CoverTab[8490]"
//line /usr/local/go/src/net/udpsock_posix.go:207
	}() {
//line /usr/local/go/src/net/udpsock_posix.go:207
		_go_fuzz_dep_.CoverTab[8491]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/udpsock_posix.go:208
			_go_fuzz_dep_.CoverTab[8492]++
									return sd.Dialer.Control(network, address, c)
//line /usr/local/go/src/net/udpsock_posix.go:209
			// _ = "end of CoverTab[8492]"
		}
//line /usr/local/go/src/net/udpsock_posix.go:210
		// _ = "end of CoverTab[8491]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:211
		_go_fuzz_dep_.CoverTab[8493]++
//line /usr/local/go/src/net/udpsock_posix.go:211
		// _ = "end of CoverTab[8493]"
//line /usr/local/go/src/net/udpsock_posix.go:211
	}
//line /usr/local/go/src/net/udpsock_posix.go:211
	// _ = "end of CoverTab[8487]"
//line /usr/local/go/src/net/udpsock_posix.go:211
	_go_fuzz_dep_.CoverTab[8488]++
							fd, err := internetSocket(ctx, sd.network, laddr, raddr, syscall.SOCK_DGRAM, 0, "dial", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:213
		_go_fuzz_dep_.CoverTab[8494]++
								return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:214
		// _ = "end of CoverTab[8494]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:215
		_go_fuzz_dep_.CoverTab[8495]++
//line /usr/local/go/src/net/udpsock_posix.go:215
		// _ = "end of CoverTab[8495]"
//line /usr/local/go/src/net/udpsock_posix.go:215
	}
//line /usr/local/go/src/net/udpsock_posix.go:215
	// _ = "end of CoverTab[8488]"
//line /usr/local/go/src/net/udpsock_posix.go:215
	_go_fuzz_dep_.CoverTab[8489]++
							return newUDPConn(fd), nil
//line /usr/local/go/src/net/udpsock_posix.go:216
	// _ = "end of CoverTab[8489]"
}

func (sl *sysListener) listenUDP(ctx context.Context, laddr *UDPAddr) (*UDPConn, error) {
//line /usr/local/go/src/net/udpsock_posix.go:219
	_go_fuzz_dep_.CoverTab[8496]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/udpsock_posix.go:221
		_go_fuzz_dep_.CoverTab[8499]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/udpsock_posix.go:222
			_go_fuzz_dep_.CoverTab[8500]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/udpsock_posix.go:223
			// _ = "end of CoverTab[8500]"
		}
//line /usr/local/go/src/net/udpsock_posix.go:224
		// _ = "end of CoverTab[8499]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:225
		_go_fuzz_dep_.CoverTab[8501]++
//line /usr/local/go/src/net/udpsock_posix.go:225
		// _ = "end of CoverTab[8501]"
//line /usr/local/go/src/net/udpsock_posix.go:225
	}
//line /usr/local/go/src/net/udpsock_posix.go:225
	// _ = "end of CoverTab[8496]"
//line /usr/local/go/src/net/udpsock_posix.go:225
	_go_fuzz_dep_.CoverTab[8497]++
							fd, err := internetSocket(ctx, sl.network, laddr, nil, syscall.SOCK_DGRAM, 0, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:227
		_go_fuzz_dep_.CoverTab[8502]++
								return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:228
		// _ = "end of CoverTab[8502]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:229
		_go_fuzz_dep_.CoverTab[8503]++
//line /usr/local/go/src/net/udpsock_posix.go:229
		// _ = "end of CoverTab[8503]"
//line /usr/local/go/src/net/udpsock_posix.go:229
	}
//line /usr/local/go/src/net/udpsock_posix.go:229
	// _ = "end of CoverTab[8497]"
//line /usr/local/go/src/net/udpsock_posix.go:229
	_go_fuzz_dep_.CoverTab[8498]++
							return newUDPConn(fd), nil
//line /usr/local/go/src/net/udpsock_posix.go:230
	// _ = "end of CoverTab[8498]"
}

func (sl *sysListener) listenMulticastUDP(ctx context.Context, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error) {
//line /usr/local/go/src/net/udpsock_posix.go:233
	_go_fuzz_dep_.CoverTab[8504]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/udpsock_posix.go:235
		_go_fuzz_dep_.CoverTab[8508]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/udpsock_posix.go:236
			_go_fuzz_dep_.CoverTab[8509]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/udpsock_posix.go:237
			// _ = "end of CoverTab[8509]"
		}
//line /usr/local/go/src/net/udpsock_posix.go:238
		// _ = "end of CoverTab[8508]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:239
		_go_fuzz_dep_.CoverTab[8510]++
//line /usr/local/go/src/net/udpsock_posix.go:239
		// _ = "end of CoverTab[8510]"
//line /usr/local/go/src/net/udpsock_posix.go:239
	}
//line /usr/local/go/src/net/udpsock_posix.go:239
	// _ = "end of CoverTab[8504]"
//line /usr/local/go/src/net/udpsock_posix.go:239
	_go_fuzz_dep_.CoverTab[8505]++
							fd, err := internetSocket(ctx, sl.network, gaddr, nil, syscall.SOCK_DGRAM, 0, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:241
		_go_fuzz_dep_.CoverTab[8511]++
								return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:242
		// _ = "end of CoverTab[8511]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:243
		_go_fuzz_dep_.CoverTab[8512]++
//line /usr/local/go/src/net/udpsock_posix.go:243
		// _ = "end of CoverTab[8512]"
//line /usr/local/go/src/net/udpsock_posix.go:243
	}
//line /usr/local/go/src/net/udpsock_posix.go:243
	// _ = "end of CoverTab[8505]"
//line /usr/local/go/src/net/udpsock_posix.go:243
	_go_fuzz_dep_.CoverTab[8506]++
							c := newUDPConn(fd)
							if ip4 := gaddr.IP.To4(); ip4 != nil {
//line /usr/local/go/src/net/udpsock_posix.go:245
		_go_fuzz_dep_.CoverTab[8513]++
								if err := listenIPv4MulticastUDP(c, ifi, ip4); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:246
			_go_fuzz_dep_.CoverTab[8514]++
									c.Close()
									return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:248
			// _ = "end of CoverTab[8514]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:249
			_go_fuzz_dep_.CoverTab[8515]++
//line /usr/local/go/src/net/udpsock_posix.go:249
			// _ = "end of CoverTab[8515]"
//line /usr/local/go/src/net/udpsock_posix.go:249
		}
//line /usr/local/go/src/net/udpsock_posix.go:249
		// _ = "end of CoverTab[8513]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:250
		_go_fuzz_dep_.CoverTab[8516]++
								if err := listenIPv6MulticastUDP(c, ifi, gaddr.IP); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:251
			_go_fuzz_dep_.CoverTab[8517]++
									c.Close()
									return nil, err
//line /usr/local/go/src/net/udpsock_posix.go:253
			// _ = "end of CoverTab[8517]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:254
			_go_fuzz_dep_.CoverTab[8518]++
//line /usr/local/go/src/net/udpsock_posix.go:254
			// _ = "end of CoverTab[8518]"
//line /usr/local/go/src/net/udpsock_posix.go:254
		}
//line /usr/local/go/src/net/udpsock_posix.go:254
		// _ = "end of CoverTab[8516]"
	}
//line /usr/local/go/src/net/udpsock_posix.go:255
	// _ = "end of CoverTab[8506]"
//line /usr/local/go/src/net/udpsock_posix.go:255
	_go_fuzz_dep_.CoverTab[8507]++
							return c, nil
//line /usr/local/go/src/net/udpsock_posix.go:256
	// _ = "end of CoverTab[8507]"
}

func listenIPv4MulticastUDP(c *UDPConn, ifi *Interface, ip IP) error {
//line /usr/local/go/src/net/udpsock_posix.go:259
	_go_fuzz_dep_.CoverTab[8519]++
							if ifi != nil {
//line /usr/local/go/src/net/udpsock_posix.go:260
		_go_fuzz_dep_.CoverTab[8523]++
								if err := setIPv4MulticastInterface(c.fd, ifi); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:261
			_go_fuzz_dep_.CoverTab[8524]++
									return err
//line /usr/local/go/src/net/udpsock_posix.go:262
			// _ = "end of CoverTab[8524]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:263
			_go_fuzz_dep_.CoverTab[8525]++
//line /usr/local/go/src/net/udpsock_posix.go:263
			// _ = "end of CoverTab[8525]"
//line /usr/local/go/src/net/udpsock_posix.go:263
		}
//line /usr/local/go/src/net/udpsock_posix.go:263
		// _ = "end of CoverTab[8523]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:264
		_go_fuzz_dep_.CoverTab[8526]++
//line /usr/local/go/src/net/udpsock_posix.go:264
		// _ = "end of CoverTab[8526]"
//line /usr/local/go/src/net/udpsock_posix.go:264
	}
//line /usr/local/go/src/net/udpsock_posix.go:264
	// _ = "end of CoverTab[8519]"
//line /usr/local/go/src/net/udpsock_posix.go:264
	_go_fuzz_dep_.CoverTab[8520]++
							if err := setIPv4MulticastLoopback(c.fd, false); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:265
		_go_fuzz_dep_.CoverTab[8527]++
								return err
//line /usr/local/go/src/net/udpsock_posix.go:266
		// _ = "end of CoverTab[8527]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:267
		_go_fuzz_dep_.CoverTab[8528]++
//line /usr/local/go/src/net/udpsock_posix.go:267
		// _ = "end of CoverTab[8528]"
//line /usr/local/go/src/net/udpsock_posix.go:267
	}
//line /usr/local/go/src/net/udpsock_posix.go:267
	// _ = "end of CoverTab[8520]"
//line /usr/local/go/src/net/udpsock_posix.go:267
	_go_fuzz_dep_.CoverTab[8521]++
							if err := joinIPv4Group(c.fd, ifi, ip); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:268
		_go_fuzz_dep_.CoverTab[8529]++
								return err
//line /usr/local/go/src/net/udpsock_posix.go:269
		// _ = "end of CoverTab[8529]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:270
		_go_fuzz_dep_.CoverTab[8530]++
//line /usr/local/go/src/net/udpsock_posix.go:270
		// _ = "end of CoverTab[8530]"
//line /usr/local/go/src/net/udpsock_posix.go:270
	}
//line /usr/local/go/src/net/udpsock_posix.go:270
	// _ = "end of CoverTab[8521]"
//line /usr/local/go/src/net/udpsock_posix.go:270
	_go_fuzz_dep_.CoverTab[8522]++
							return nil
//line /usr/local/go/src/net/udpsock_posix.go:271
	// _ = "end of CoverTab[8522]"
}

func listenIPv6MulticastUDP(c *UDPConn, ifi *Interface, ip IP) error {
//line /usr/local/go/src/net/udpsock_posix.go:274
	_go_fuzz_dep_.CoverTab[8531]++
							if ifi != nil {
//line /usr/local/go/src/net/udpsock_posix.go:275
		_go_fuzz_dep_.CoverTab[8535]++
								if err := setIPv6MulticastInterface(c.fd, ifi); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:276
			_go_fuzz_dep_.CoverTab[8536]++
									return err
//line /usr/local/go/src/net/udpsock_posix.go:277
			// _ = "end of CoverTab[8536]"
		} else {
//line /usr/local/go/src/net/udpsock_posix.go:278
			_go_fuzz_dep_.CoverTab[8537]++
//line /usr/local/go/src/net/udpsock_posix.go:278
			// _ = "end of CoverTab[8537]"
//line /usr/local/go/src/net/udpsock_posix.go:278
		}
//line /usr/local/go/src/net/udpsock_posix.go:278
		// _ = "end of CoverTab[8535]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:279
		_go_fuzz_dep_.CoverTab[8538]++
//line /usr/local/go/src/net/udpsock_posix.go:279
		// _ = "end of CoverTab[8538]"
//line /usr/local/go/src/net/udpsock_posix.go:279
	}
//line /usr/local/go/src/net/udpsock_posix.go:279
	// _ = "end of CoverTab[8531]"
//line /usr/local/go/src/net/udpsock_posix.go:279
	_go_fuzz_dep_.CoverTab[8532]++
							if err := setIPv6MulticastLoopback(c.fd, false); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:280
		_go_fuzz_dep_.CoverTab[8539]++
								return err
//line /usr/local/go/src/net/udpsock_posix.go:281
		// _ = "end of CoverTab[8539]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:282
		_go_fuzz_dep_.CoverTab[8540]++
//line /usr/local/go/src/net/udpsock_posix.go:282
		// _ = "end of CoverTab[8540]"
//line /usr/local/go/src/net/udpsock_posix.go:282
	}
//line /usr/local/go/src/net/udpsock_posix.go:282
	// _ = "end of CoverTab[8532]"
//line /usr/local/go/src/net/udpsock_posix.go:282
	_go_fuzz_dep_.CoverTab[8533]++
							if err := joinIPv6Group(c.fd, ifi, ip); err != nil {
//line /usr/local/go/src/net/udpsock_posix.go:283
		_go_fuzz_dep_.CoverTab[8541]++
								return err
//line /usr/local/go/src/net/udpsock_posix.go:284
		// _ = "end of CoverTab[8541]"
	} else {
//line /usr/local/go/src/net/udpsock_posix.go:285
		_go_fuzz_dep_.CoverTab[8542]++
//line /usr/local/go/src/net/udpsock_posix.go:285
		// _ = "end of CoverTab[8542]"
//line /usr/local/go/src/net/udpsock_posix.go:285
	}
//line /usr/local/go/src/net/udpsock_posix.go:285
	// _ = "end of CoverTab[8533]"
//line /usr/local/go/src/net/udpsock_posix.go:285
	_go_fuzz_dep_.CoverTab[8534]++
							return nil
//line /usr/local/go/src/net/udpsock_posix.go:286
	// _ = "end of CoverTab[8534]"
}

//line /usr/local/go/src/net/udpsock_posix.go:287
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/udpsock_posix.go:287
var _ = _go_fuzz_dep_.CoverTab
