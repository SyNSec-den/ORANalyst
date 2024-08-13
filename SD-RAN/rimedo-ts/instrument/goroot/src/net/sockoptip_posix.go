// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

//line /usr/local/go/src/net/sockoptip_posix.go:7
package net

//line /usr/local/go/src/net/sockoptip_posix.go:7
import (
//line /usr/local/go/src/net/sockoptip_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/sockoptip_posix.go:7
)
//line /usr/local/go/src/net/sockoptip_posix.go:7
import (
//line /usr/local/go/src/net/sockoptip_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/sockoptip_posix.go:7
)

import (
	"runtime"
	"syscall"
)

func joinIPv4Group(fd *netFD, ifi *Interface, ip IP) error {
//line /usr/local/go/src/net/sockoptip_posix.go:14
	_go_fuzz_dep_.CoverTab[16361]++
							mreq := &syscall.IPMreq{Multiaddr: [4]byte{ip[0], ip[1], ip[2], ip[3]}}
							if err := setIPv4MreqToInterface(mreq, ifi); err != nil {
//line /usr/local/go/src/net/sockoptip_posix.go:16
		_go_fuzz_dep_.CoverTab[16363]++
								return err
//line /usr/local/go/src/net/sockoptip_posix.go:17
		// _ = "end of CoverTab[16363]"
	} else {
//line /usr/local/go/src/net/sockoptip_posix.go:18
		_go_fuzz_dep_.CoverTab[16364]++
//line /usr/local/go/src/net/sockoptip_posix.go:18
		// _ = "end of CoverTab[16364]"
//line /usr/local/go/src/net/sockoptip_posix.go:18
	}
//line /usr/local/go/src/net/sockoptip_posix.go:18
	// _ = "end of CoverTab[16361]"
//line /usr/local/go/src/net/sockoptip_posix.go:18
	_go_fuzz_dep_.CoverTab[16362]++
							err := fd.pfd.SetsockoptIPMreq(syscall.IPPROTO_IP, syscall.IP_ADD_MEMBERSHIP, mreq)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockoptip_posix.go:21
	// _ = "end of CoverTab[16362]"
}

func setIPv6MulticastInterface(fd *netFD, ifi *Interface) error {
//line /usr/local/go/src/net/sockoptip_posix.go:24
	_go_fuzz_dep_.CoverTab[16365]++
							var v int
							if ifi != nil {
//line /usr/local/go/src/net/sockoptip_posix.go:26
		_go_fuzz_dep_.CoverTab[16367]++
								v = ifi.Index
//line /usr/local/go/src/net/sockoptip_posix.go:27
		// _ = "end of CoverTab[16367]"
	} else {
//line /usr/local/go/src/net/sockoptip_posix.go:28
		_go_fuzz_dep_.CoverTab[16368]++
//line /usr/local/go/src/net/sockoptip_posix.go:28
		// _ = "end of CoverTab[16368]"
//line /usr/local/go/src/net/sockoptip_posix.go:28
	}
//line /usr/local/go/src/net/sockoptip_posix.go:28
	// _ = "end of CoverTab[16365]"
//line /usr/local/go/src/net/sockoptip_posix.go:28
	_go_fuzz_dep_.CoverTab[16366]++
							err := fd.pfd.SetsockoptInt(syscall.IPPROTO_IPV6, syscall.IPV6_MULTICAST_IF, v)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockoptip_posix.go:31
	// _ = "end of CoverTab[16366]"
}

func setIPv6MulticastLoopback(fd *netFD, v bool) error {
//line /usr/local/go/src/net/sockoptip_posix.go:34
	_go_fuzz_dep_.CoverTab[16369]++
							err := fd.pfd.SetsockoptInt(syscall.IPPROTO_IPV6, syscall.IPV6_MULTICAST_LOOP, boolint(v))
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockoptip_posix.go:37
	// _ = "end of CoverTab[16369]"
}

func joinIPv6Group(fd *netFD, ifi *Interface, ip IP) error {
//line /usr/local/go/src/net/sockoptip_posix.go:40
	_go_fuzz_dep_.CoverTab[16370]++
							mreq := &syscall.IPv6Mreq{}
							copy(mreq.Multiaddr[:], ip)
							if ifi != nil {
//line /usr/local/go/src/net/sockoptip_posix.go:43
		_go_fuzz_dep_.CoverTab[16372]++
								mreq.Interface = uint32(ifi.Index)
//line /usr/local/go/src/net/sockoptip_posix.go:44
		// _ = "end of CoverTab[16372]"
	} else {
//line /usr/local/go/src/net/sockoptip_posix.go:45
		_go_fuzz_dep_.CoverTab[16373]++
//line /usr/local/go/src/net/sockoptip_posix.go:45
		// _ = "end of CoverTab[16373]"
//line /usr/local/go/src/net/sockoptip_posix.go:45
	}
//line /usr/local/go/src/net/sockoptip_posix.go:45
	// _ = "end of CoverTab[16370]"
//line /usr/local/go/src/net/sockoptip_posix.go:45
	_go_fuzz_dep_.CoverTab[16371]++
							err := fd.pfd.SetsockoptIPv6Mreq(syscall.IPPROTO_IPV6, syscall.IPV6_JOIN_GROUP, mreq)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockoptip_posix.go:48
	// _ = "end of CoverTab[16371]"
}

//line /usr/local/go/src/net/sockoptip_posix.go:49
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sockoptip_posix.go:49
var _ = _go_fuzz_dep_.CoverTab
