// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

//line /snap/go/10455/src/net/iprawsock_posix.go:7
package net

//line /snap/go/10455/src/net/iprawsock_posix.go:7
import (
//line /snap/go/10455/src/net/iprawsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/iprawsock_posix.go:7
)
//line /snap/go/10455/src/net/iprawsock_posix.go:7
import (
//line /snap/go/10455/src/net/iprawsock_posix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/iprawsock_posix.go:7
)

import (
	"context"
	"syscall"
)

func sockaddrToIP(sa syscall.Sockaddr) Addr {
//line /snap/go/10455/src/net/iprawsock_posix.go:14
	_go_fuzz_dep_.CoverTab[6632]++
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /snap/go/10455/src/net/iprawsock_posix.go:16
		_go_fuzz_dep_.CoverTab[528751]++
//line /snap/go/10455/src/net/iprawsock_posix.go:16
		_go_fuzz_dep_.CoverTab[6634]++
								return &IPAddr{IP: sa.Addr[0:]}
//line /snap/go/10455/src/net/iprawsock_posix.go:17
		// _ = "end of CoverTab[6634]"
	case *syscall.SockaddrInet6:
//line /snap/go/10455/src/net/iprawsock_posix.go:18
		_go_fuzz_dep_.CoverTab[528752]++
//line /snap/go/10455/src/net/iprawsock_posix.go:18
		_go_fuzz_dep_.CoverTab[6635]++
								return &IPAddr{IP: sa.Addr[0:], Zone: zoneCache.name(int(sa.ZoneId))}
//line /snap/go/10455/src/net/iprawsock_posix.go:19
		// _ = "end of CoverTab[6635]"
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:20
	// _ = "end of CoverTab[6632]"
//line /snap/go/10455/src/net/iprawsock_posix.go:20
	_go_fuzz_dep_.CoverTab[6633]++
							return nil
//line /snap/go/10455/src/net/iprawsock_posix.go:21
	// _ = "end of CoverTab[6633]"
}

func (a *IPAddr) family() int {
//line /snap/go/10455/src/net/iprawsock_posix.go:24
	_go_fuzz_dep_.CoverTab[6636]++
							if a == nil || func() bool {
//line /snap/go/10455/src/net/iprawsock_posix.go:25
		_go_fuzz_dep_.CoverTab[6639]++
//line /snap/go/10455/src/net/iprawsock_posix.go:25
		return len(a.IP) <= IPv4len
//line /snap/go/10455/src/net/iprawsock_posix.go:25
		// _ = "end of CoverTab[6639]"
//line /snap/go/10455/src/net/iprawsock_posix.go:25
	}() {
//line /snap/go/10455/src/net/iprawsock_posix.go:25
		_go_fuzz_dep_.CoverTab[528753]++
//line /snap/go/10455/src/net/iprawsock_posix.go:25
		_go_fuzz_dep_.CoverTab[6640]++
								return syscall.AF_INET
//line /snap/go/10455/src/net/iprawsock_posix.go:26
		// _ = "end of CoverTab[6640]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:27
		_go_fuzz_dep_.CoverTab[528754]++
//line /snap/go/10455/src/net/iprawsock_posix.go:27
		_go_fuzz_dep_.CoverTab[6641]++
//line /snap/go/10455/src/net/iprawsock_posix.go:27
		// _ = "end of CoverTab[6641]"
//line /snap/go/10455/src/net/iprawsock_posix.go:27
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:27
	// _ = "end of CoverTab[6636]"
//line /snap/go/10455/src/net/iprawsock_posix.go:27
	_go_fuzz_dep_.CoverTab[6637]++
							if a.IP.To4() != nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:28
		_go_fuzz_dep_.CoverTab[528755]++
//line /snap/go/10455/src/net/iprawsock_posix.go:28
		_go_fuzz_dep_.CoverTab[6642]++
								return syscall.AF_INET
//line /snap/go/10455/src/net/iprawsock_posix.go:29
		// _ = "end of CoverTab[6642]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:30
		_go_fuzz_dep_.CoverTab[528756]++
//line /snap/go/10455/src/net/iprawsock_posix.go:30
		_go_fuzz_dep_.CoverTab[6643]++
//line /snap/go/10455/src/net/iprawsock_posix.go:30
		// _ = "end of CoverTab[6643]"
//line /snap/go/10455/src/net/iprawsock_posix.go:30
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:30
	// _ = "end of CoverTab[6637]"
//line /snap/go/10455/src/net/iprawsock_posix.go:30
	_go_fuzz_dep_.CoverTab[6638]++
							return syscall.AF_INET6
//line /snap/go/10455/src/net/iprawsock_posix.go:31
	// _ = "end of CoverTab[6638]"
}

func (a *IPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /snap/go/10455/src/net/iprawsock_posix.go:34
	_go_fuzz_dep_.CoverTab[6644]++
							if a == nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:35
		_go_fuzz_dep_.CoverTab[528757]++
//line /snap/go/10455/src/net/iprawsock_posix.go:35
		_go_fuzz_dep_.CoverTab[6646]++
								return nil, nil
//line /snap/go/10455/src/net/iprawsock_posix.go:36
		// _ = "end of CoverTab[6646]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:37
		_go_fuzz_dep_.CoverTab[528758]++
//line /snap/go/10455/src/net/iprawsock_posix.go:37
		_go_fuzz_dep_.CoverTab[6647]++
//line /snap/go/10455/src/net/iprawsock_posix.go:37
		// _ = "end of CoverTab[6647]"
//line /snap/go/10455/src/net/iprawsock_posix.go:37
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:37
	// _ = "end of CoverTab[6644]"
//line /snap/go/10455/src/net/iprawsock_posix.go:37
	_go_fuzz_dep_.CoverTab[6645]++
							return ipToSockaddr(family, a.IP, 0, a.Zone)
//line /snap/go/10455/src/net/iprawsock_posix.go:38
	// _ = "end of CoverTab[6645]"
}

func (a *IPAddr) toLocal(net string) sockaddr {
//line /snap/go/10455/src/net/iprawsock_posix.go:41
	_go_fuzz_dep_.CoverTab[6648]++
							return &IPAddr{loopbackIP(net), a.Zone}
//line /snap/go/10455/src/net/iprawsock_posix.go:42
	// _ = "end of CoverTab[6648]"
}

func (c *IPConn) readFrom(b []byte) (int, *IPAddr, error) {
//line /snap/go/10455/src/net/iprawsock_posix.go:45
	_go_fuzz_dep_.CoverTab[6649]++
	// TODO(cw,rsc): consider using readv if we know the family
	// type to avoid the header trim/copy
	var addr *IPAddr
	n, sa, err := c.fd.readFrom(b)
	switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /snap/go/10455/src/net/iprawsock_posix.go:51
		_go_fuzz_dep_.CoverTab[528759]++
//line /snap/go/10455/src/net/iprawsock_posix.go:51
		_go_fuzz_dep_.CoverTab[6651]++
								addr = &IPAddr{IP: sa.Addr[0:]}
								n = stripIPv4Header(n, b)
//line /snap/go/10455/src/net/iprawsock_posix.go:53
		// _ = "end of CoverTab[6651]"
	case *syscall.SockaddrInet6:
//line /snap/go/10455/src/net/iprawsock_posix.go:54
		_go_fuzz_dep_.CoverTab[528760]++
//line /snap/go/10455/src/net/iprawsock_posix.go:54
		_go_fuzz_dep_.CoverTab[6652]++
								addr = &IPAddr{IP: sa.Addr[0:], Zone: zoneCache.name(int(sa.ZoneId))}
//line /snap/go/10455/src/net/iprawsock_posix.go:55
		// _ = "end of CoverTab[6652]"
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:56
	// _ = "end of CoverTab[6649]"
//line /snap/go/10455/src/net/iprawsock_posix.go:56
	_go_fuzz_dep_.CoverTab[6650]++
							return n, addr, err
//line /snap/go/10455/src/net/iprawsock_posix.go:57
	// _ = "end of CoverTab[6650]"
}

func stripIPv4Header(n int, b []byte) int {
//line /snap/go/10455/src/net/iprawsock_posix.go:60
	_go_fuzz_dep_.CoverTab[6653]++
							if len(b) < 20 {
//line /snap/go/10455/src/net/iprawsock_posix.go:61
		_go_fuzz_dep_.CoverTab[528761]++
//line /snap/go/10455/src/net/iprawsock_posix.go:61
		_go_fuzz_dep_.CoverTab[6657]++
								return n
//line /snap/go/10455/src/net/iprawsock_posix.go:62
		// _ = "end of CoverTab[6657]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:63
		_go_fuzz_dep_.CoverTab[528762]++
//line /snap/go/10455/src/net/iprawsock_posix.go:63
		_go_fuzz_dep_.CoverTab[6658]++
//line /snap/go/10455/src/net/iprawsock_posix.go:63
		// _ = "end of CoverTab[6658]"
//line /snap/go/10455/src/net/iprawsock_posix.go:63
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:63
	// _ = "end of CoverTab[6653]"
//line /snap/go/10455/src/net/iprawsock_posix.go:63
	_go_fuzz_dep_.CoverTab[6654]++
							l := int(b[0]&0x0f) << 2
							if 20 > l || func() bool {
//line /snap/go/10455/src/net/iprawsock_posix.go:65
		_go_fuzz_dep_.CoverTab[6659]++
//line /snap/go/10455/src/net/iprawsock_posix.go:65
		return l > len(b)
//line /snap/go/10455/src/net/iprawsock_posix.go:65
		// _ = "end of CoverTab[6659]"
//line /snap/go/10455/src/net/iprawsock_posix.go:65
	}() {
//line /snap/go/10455/src/net/iprawsock_posix.go:65
		_go_fuzz_dep_.CoverTab[528763]++
//line /snap/go/10455/src/net/iprawsock_posix.go:65
		_go_fuzz_dep_.CoverTab[6660]++
								return n
//line /snap/go/10455/src/net/iprawsock_posix.go:66
		// _ = "end of CoverTab[6660]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:67
		_go_fuzz_dep_.CoverTab[528764]++
//line /snap/go/10455/src/net/iprawsock_posix.go:67
		_go_fuzz_dep_.CoverTab[6661]++
//line /snap/go/10455/src/net/iprawsock_posix.go:67
		// _ = "end of CoverTab[6661]"
//line /snap/go/10455/src/net/iprawsock_posix.go:67
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:67
	// _ = "end of CoverTab[6654]"
//line /snap/go/10455/src/net/iprawsock_posix.go:67
	_go_fuzz_dep_.CoverTab[6655]++
							if b[0]>>4 != 4 {
//line /snap/go/10455/src/net/iprawsock_posix.go:68
		_go_fuzz_dep_.CoverTab[528765]++
//line /snap/go/10455/src/net/iprawsock_posix.go:68
		_go_fuzz_dep_.CoverTab[6662]++
								return n
//line /snap/go/10455/src/net/iprawsock_posix.go:69
		// _ = "end of CoverTab[6662]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:70
		_go_fuzz_dep_.CoverTab[528766]++
//line /snap/go/10455/src/net/iprawsock_posix.go:70
		_go_fuzz_dep_.CoverTab[6663]++
//line /snap/go/10455/src/net/iprawsock_posix.go:70
		// _ = "end of CoverTab[6663]"
//line /snap/go/10455/src/net/iprawsock_posix.go:70
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:70
	// _ = "end of CoverTab[6655]"
//line /snap/go/10455/src/net/iprawsock_posix.go:70
	_go_fuzz_dep_.CoverTab[6656]++
							copy(b, b[l:])
							return n - l
//line /snap/go/10455/src/net/iprawsock_posix.go:72
	// _ = "end of CoverTab[6656]"
}

func (c *IPConn) readMsg(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) {
//line /snap/go/10455/src/net/iprawsock_posix.go:75
	_go_fuzz_dep_.CoverTab[6664]++
							var sa syscall.Sockaddr
							n, oobn, flags, sa, err = c.fd.readMsg(b, oob, 0)
							switch sa := sa.(type) {
	case *syscall.SockaddrInet4:
//line /snap/go/10455/src/net/iprawsock_posix.go:79
		_go_fuzz_dep_.CoverTab[528767]++
//line /snap/go/10455/src/net/iprawsock_posix.go:79
		_go_fuzz_dep_.CoverTab[6666]++
								addr = &IPAddr{IP: sa.Addr[0:]}
//line /snap/go/10455/src/net/iprawsock_posix.go:80
		// _ = "end of CoverTab[6666]"
	case *syscall.SockaddrInet6:
//line /snap/go/10455/src/net/iprawsock_posix.go:81
		_go_fuzz_dep_.CoverTab[528768]++
//line /snap/go/10455/src/net/iprawsock_posix.go:81
		_go_fuzz_dep_.CoverTab[6667]++
								addr = &IPAddr{IP: sa.Addr[0:], Zone: zoneCache.name(int(sa.ZoneId))}
//line /snap/go/10455/src/net/iprawsock_posix.go:82
		// _ = "end of CoverTab[6667]"
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:83
	// _ = "end of CoverTab[6664]"
//line /snap/go/10455/src/net/iprawsock_posix.go:83
	_go_fuzz_dep_.CoverTab[6665]++
							return
//line /snap/go/10455/src/net/iprawsock_posix.go:84
	// _ = "end of CoverTab[6665]"
}

func (c *IPConn) writeTo(b []byte, addr *IPAddr) (int, error) {
//line /snap/go/10455/src/net/iprawsock_posix.go:87
	_go_fuzz_dep_.CoverTab[6668]++
							if c.fd.isConnected {
//line /snap/go/10455/src/net/iprawsock_posix.go:88
		_go_fuzz_dep_.CoverTab[528769]++
//line /snap/go/10455/src/net/iprawsock_posix.go:88
		_go_fuzz_dep_.CoverTab[6672]++
								return 0, ErrWriteToConnected
//line /snap/go/10455/src/net/iprawsock_posix.go:89
		// _ = "end of CoverTab[6672]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:90
		_go_fuzz_dep_.CoverTab[528770]++
//line /snap/go/10455/src/net/iprawsock_posix.go:90
		_go_fuzz_dep_.CoverTab[6673]++
//line /snap/go/10455/src/net/iprawsock_posix.go:90
		// _ = "end of CoverTab[6673]"
//line /snap/go/10455/src/net/iprawsock_posix.go:90
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:90
	// _ = "end of CoverTab[6668]"
//line /snap/go/10455/src/net/iprawsock_posix.go:90
	_go_fuzz_dep_.CoverTab[6669]++
							if addr == nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:91
		_go_fuzz_dep_.CoverTab[528771]++
//line /snap/go/10455/src/net/iprawsock_posix.go:91
		_go_fuzz_dep_.CoverTab[6674]++
								return 0, errMissingAddress
//line /snap/go/10455/src/net/iprawsock_posix.go:92
		// _ = "end of CoverTab[6674]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:93
		_go_fuzz_dep_.CoverTab[528772]++
//line /snap/go/10455/src/net/iprawsock_posix.go:93
		_go_fuzz_dep_.CoverTab[6675]++
//line /snap/go/10455/src/net/iprawsock_posix.go:93
		// _ = "end of CoverTab[6675]"
//line /snap/go/10455/src/net/iprawsock_posix.go:93
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:93
	// _ = "end of CoverTab[6669]"
//line /snap/go/10455/src/net/iprawsock_posix.go:93
	_go_fuzz_dep_.CoverTab[6670]++
							sa, err := addr.sockaddr(c.fd.family)
							if err != nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:95
		_go_fuzz_dep_.CoverTab[528773]++
//line /snap/go/10455/src/net/iprawsock_posix.go:95
		_go_fuzz_dep_.CoverTab[6676]++
								return 0, err
//line /snap/go/10455/src/net/iprawsock_posix.go:96
		// _ = "end of CoverTab[6676]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:97
		_go_fuzz_dep_.CoverTab[528774]++
//line /snap/go/10455/src/net/iprawsock_posix.go:97
		_go_fuzz_dep_.CoverTab[6677]++
//line /snap/go/10455/src/net/iprawsock_posix.go:97
		// _ = "end of CoverTab[6677]"
//line /snap/go/10455/src/net/iprawsock_posix.go:97
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:97
	// _ = "end of CoverTab[6670]"
//line /snap/go/10455/src/net/iprawsock_posix.go:97
	_go_fuzz_dep_.CoverTab[6671]++
							return c.fd.writeTo(b, sa)
//line /snap/go/10455/src/net/iprawsock_posix.go:98
	// _ = "end of CoverTab[6671]"
}

func (c *IPConn) writeMsg(b, oob []byte, addr *IPAddr) (n, oobn int, err error) {
//line /snap/go/10455/src/net/iprawsock_posix.go:101
	_go_fuzz_dep_.CoverTab[6678]++
							if c.fd.isConnected {
//line /snap/go/10455/src/net/iprawsock_posix.go:102
		_go_fuzz_dep_.CoverTab[528775]++
//line /snap/go/10455/src/net/iprawsock_posix.go:102
		_go_fuzz_dep_.CoverTab[6682]++
								return 0, 0, ErrWriteToConnected
//line /snap/go/10455/src/net/iprawsock_posix.go:103
		// _ = "end of CoverTab[6682]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:104
		_go_fuzz_dep_.CoverTab[528776]++
//line /snap/go/10455/src/net/iprawsock_posix.go:104
		_go_fuzz_dep_.CoverTab[6683]++
//line /snap/go/10455/src/net/iprawsock_posix.go:104
		// _ = "end of CoverTab[6683]"
//line /snap/go/10455/src/net/iprawsock_posix.go:104
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:104
	// _ = "end of CoverTab[6678]"
//line /snap/go/10455/src/net/iprawsock_posix.go:104
	_go_fuzz_dep_.CoverTab[6679]++
							if addr == nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:105
		_go_fuzz_dep_.CoverTab[528777]++
//line /snap/go/10455/src/net/iprawsock_posix.go:105
		_go_fuzz_dep_.CoverTab[6684]++
								return 0, 0, errMissingAddress
//line /snap/go/10455/src/net/iprawsock_posix.go:106
		// _ = "end of CoverTab[6684]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:107
		_go_fuzz_dep_.CoverTab[528778]++
//line /snap/go/10455/src/net/iprawsock_posix.go:107
		_go_fuzz_dep_.CoverTab[6685]++
//line /snap/go/10455/src/net/iprawsock_posix.go:107
		// _ = "end of CoverTab[6685]"
//line /snap/go/10455/src/net/iprawsock_posix.go:107
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:107
	// _ = "end of CoverTab[6679]"
//line /snap/go/10455/src/net/iprawsock_posix.go:107
	_go_fuzz_dep_.CoverTab[6680]++
							sa, err := addr.sockaddr(c.fd.family)
							if err != nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:109
		_go_fuzz_dep_.CoverTab[528779]++
//line /snap/go/10455/src/net/iprawsock_posix.go:109
		_go_fuzz_dep_.CoverTab[6686]++
								return 0, 0, err
//line /snap/go/10455/src/net/iprawsock_posix.go:110
		// _ = "end of CoverTab[6686]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:111
		_go_fuzz_dep_.CoverTab[528780]++
//line /snap/go/10455/src/net/iprawsock_posix.go:111
		_go_fuzz_dep_.CoverTab[6687]++
//line /snap/go/10455/src/net/iprawsock_posix.go:111
		// _ = "end of CoverTab[6687]"
//line /snap/go/10455/src/net/iprawsock_posix.go:111
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:111
	// _ = "end of CoverTab[6680]"
//line /snap/go/10455/src/net/iprawsock_posix.go:111
	_go_fuzz_dep_.CoverTab[6681]++
							return c.fd.writeMsg(b, oob, sa)
//line /snap/go/10455/src/net/iprawsock_posix.go:112
	// _ = "end of CoverTab[6681]"
}

func (sd *sysDialer) dialIP(ctx context.Context, laddr, raddr *IPAddr) (*IPConn, error) {
//line /snap/go/10455/src/net/iprawsock_posix.go:115
	_go_fuzz_dep_.CoverTab[6688]++
							network, proto, err := parseNetwork(ctx, sd.network, true)
							if err != nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:117
		_go_fuzz_dep_.CoverTab[528781]++
//line /snap/go/10455/src/net/iprawsock_posix.go:117
		_go_fuzz_dep_.CoverTab[6693]++
								return nil, err
//line /snap/go/10455/src/net/iprawsock_posix.go:118
		// _ = "end of CoverTab[6693]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:119
		_go_fuzz_dep_.CoverTab[528782]++
//line /snap/go/10455/src/net/iprawsock_posix.go:119
		_go_fuzz_dep_.CoverTab[6694]++
//line /snap/go/10455/src/net/iprawsock_posix.go:119
		// _ = "end of CoverTab[6694]"
//line /snap/go/10455/src/net/iprawsock_posix.go:119
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:119
	// _ = "end of CoverTab[6688]"
//line /snap/go/10455/src/net/iprawsock_posix.go:119
	_go_fuzz_dep_.CoverTab[6689]++
							switch network {
	case "ip", "ip4", "ip6":
//line /snap/go/10455/src/net/iprawsock_posix.go:121
		_go_fuzz_dep_.CoverTab[528783]++
//line /snap/go/10455/src/net/iprawsock_posix.go:121
		_go_fuzz_dep_.CoverTab[6695]++
//line /snap/go/10455/src/net/iprawsock_posix.go:121
		// _ = "end of CoverTab[6695]"
	default:
//line /snap/go/10455/src/net/iprawsock_posix.go:122
		_go_fuzz_dep_.CoverTab[528784]++
//line /snap/go/10455/src/net/iprawsock_posix.go:122
		_go_fuzz_dep_.CoverTab[6696]++
								return nil, UnknownNetworkError(sd.network)
//line /snap/go/10455/src/net/iprawsock_posix.go:123
		// _ = "end of CoverTab[6696]"
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:124
	// _ = "end of CoverTab[6689]"
//line /snap/go/10455/src/net/iprawsock_posix.go:124
	_go_fuzz_dep_.CoverTab[6690]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /snap/go/10455/src/net/iprawsock_posix.go:126
		_go_fuzz_dep_.CoverTab[6697]++
//line /snap/go/10455/src/net/iprawsock_posix.go:126
		return sd.Dialer.Control != nil
//line /snap/go/10455/src/net/iprawsock_posix.go:126
		// _ = "end of CoverTab[6697]"
//line /snap/go/10455/src/net/iprawsock_posix.go:126
	}() {
//line /snap/go/10455/src/net/iprawsock_posix.go:126
		_go_fuzz_dep_.CoverTab[528785]++
//line /snap/go/10455/src/net/iprawsock_posix.go:126
		_go_fuzz_dep_.CoverTab[6698]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/iprawsock_posix.go:127
			_go_fuzz_dep_.CoverTab[6699]++
									return sd.Dialer.Control(network, address, c)
//line /snap/go/10455/src/net/iprawsock_posix.go:128
			// _ = "end of CoverTab[6699]"
		}
//line /snap/go/10455/src/net/iprawsock_posix.go:129
		// _ = "end of CoverTab[6698]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:130
		_go_fuzz_dep_.CoverTab[528786]++
//line /snap/go/10455/src/net/iprawsock_posix.go:130
		_go_fuzz_dep_.CoverTab[6700]++
//line /snap/go/10455/src/net/iprawsock_posix.go:130
		// _ = "end of CoverTab[6700]"
//line /snap/go/10455/src/net/iprawsock_posix.go:130
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:130
	// _ = "end of CoverTab[6690]"
//line /snap/go/10455/src/net/iprawsock_posix.go:130
	_go_fuzz_dep_.CoverTab[6691]++
							fd, err := internetSocket(ctx, network, laddr, raddr, syscall.SOCK_RAW, proto, "dial", ctrlCtxFn)
							if err != nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:132
		_go_fuzz_dep_.CoverTab[528787]++
//line /snap/go/10455/src/net/iprawsock_posix.go:132
		_go_fuzz_dep_.CoverTab[6701]++
								return nil, err
//line /snap/go/10455/src/net/iprawsock_posix.go:133
		// _ = "end of CoverTab[6701]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:134
		_go_fuzz_dep_.CoverTab[528788]++
//line /snap/go/10455/src/net/iprawsock_posix.go:134
		_go_fuzz_dep_.CoverTab[6702]++
//line /snap/go/10455/src/net/iprawsock_posix.go:134
		// _ = "end of CoverTab[6702]"
//line /snap/go/10455/src/net/iprawsock_posix.go:134
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:134
	// _ = "end of CoverTab[6691]"
//line /snap/go/10455/src/net/iprawsock_posix.go:134
	_go_fuzz_dep_.CoverTab[6692]++
							return newIPConn(fd), nil
//line /snap/go/10455/src/net/iprawsock_posix.go:135
	// _ = "end of CoverTab[6692]"
}

func (sl *sysListener) listenIP(ctx context.Context, laddr *IPAddr) (*IPConn, error) {
//line /snap/go/10455/src/net/iprawsock_posix.go:138
	_go_fuzz_dep_.CoverTab[6703]++
							network, proto, err := parseNetwork(ctx, sl.network, true)
							if err != nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:140
		_go_fuzz_dep_.CoverTab[528789]++
//line /snap/go/10455/src/net/iprawsock_posix.go:140
		_go_fuzz_dep_.CoverTab[6708]++
								return nil, err
//line /snap/go/10455/src/net/iprawsock_posix.go:141
		// _ = "end of CoverTab[6708]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:142
		_go_fuzz_dep_.CoverTab[528790]++
//line /snap/go/10455/src/net/iprawsock_posix.go:142
		_go_fuzz_dep_.CoverTab[6709]++
//line /snap/go/10455/src/net/iprawsock_posix.go:142
		// _ = "end of CoverTab[6709]"
//line /snap/go/10455/src/net/iprawsock_posix.go:142
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:142
	// _ = "end of CoverTab[6703]"
//line /snap/go/10455/src/net/iprawsock_posix.go:142
	_go_fuzz_dep_.CoverTab[6704]++
							switch network {
	case "ip", "ip4", "ip6":
//line /snap/go/10455/src/net/iprawsock_posix.go:144
		_go_fuzz_dep_.CoverTab[528791]++
//line /snap/go/10455/src/net/iprawsock_posix.go:144
		_go_fuzz_dep_.CoverTab[6710]++
//line /snap/go/10455/src/net/iprawsock_posix.go:144
		// _ = "end of CoverTab[6710]"
	default:
//line /snap/go/10455/src/net/iprawsock_posix.go:145
		_go_fuzz_dep_.CoverTab[528792]++
//line /snap/go/10455/src/net/iprawsock_posix.go:145
		_go_fuzz_dep_.CoverTab[6711]++
								return nil, UnknownNetworkError(sl.network)
//line /snap/go/10455/src/net/iprawsock_posix.go:146
		// _ = "end of CoverTab[6711]"
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:147
	// _ = "end of CoverTab[6704]"
//line /snap/go/10455/src/net/iprawsock_posix.go:147
	_go_fuzz_dep_.CoverTab[6705]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:149
		_go_fuzz_dep_.CoverTab[528793]++
//line /snap/go/10455/src/net/iprawsock_posix.go:149
		_go_fuzz_dep_.CoverTab[6712]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/iprawsock_posix.go:150
			_go_fuzz_dep_.CoverTab[6713]++
									return sl.ListenConfig.Control(network, address, c)
//line /snap/go/10455/src/net/iprawsock_posix.go:151
			// _ = "end of CoverTab[6713]"
		}
//line /snap/go/10455/src/net/iprawsock_posix.go:152
		// _ = "end of CoverTab[6712]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:153
		_go_fuzz_dep_.CoverTab[528794]++
//line /snap/go/10455/src/net/iprawsock_posix.go:153
		_go_fuzz_dep_.CoverTab[6714]++
//line /snap/go/10455/src/net/iprawsock_posix.go:153
		// _ = "end of CoverTab[6714]"
//line /snap/go/10455/src/net/iprawsock_posix.go:153
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:153
	// _ = "end of CoverTab[6705]"
//line /snap/go/10455/src/net/iprawsock_posix.go:153
	_go_fuzz_dep_.CoverTab[6706]++
							fd, err := internetSocket(ctx, network, laddr, nil, syscall.SOCK_RAW, proto, "listen", ctrlCtxFn)
							if err != nil {
//line /snap/go/10455/src/net/iprawsock_posix.go:155
		_go_fuzz_dep_.CoverTab[528795]++
//line /snap/go/10455/src/net/iprawsock_posix.go:155
		_go_fuzz_dep_.CoverTab[6715]++
								return nil, err
//line /snap/go/10455/src/net/iprawsock_posix.go:156
		// _ = "end of CoverTab[6715]"
	} else {
//line /snap/go/10455/src/net/iprawsock_posix.go:157
		_go_fuzz_dep_.CoverTab[528796]++
//line /snap/go/10455/src/net/iprawsock_posix.go:157
		_go_fuzz_dep_.CoverTab[6716]++
//line /snap/go/10455/src/net/iprawsock_posix.go:157
		// _ = "end of CoverTab[6716]"
//line /snap/go/10455/src/net/iprawsock_posix.go:157
	}
//line /snap/go/10455/src/net/iprawsock_posix.go:157
	// _ = "end of CoverTab[6706]"
//line /snap/go/10455/src/net/iprawsock_posix.go:157
	_go_fuzz_dep_.CoverTab[6707]++
							return newIPConn(fd), nil
//line /snap/go/10455/src/net/iprawsock_posix.go:158
	// _ = "end of CoverTab[6707]"
}

//line /snap/go/10455/src/net/iprawsock_posix.go:159
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/iprawsock_posix.go:159
var _ = _go_fuzz_dep_.CoverTab
