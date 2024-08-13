// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

//line /snap/go/10455/src/net/sockoptip_posix.go:7
package net

//line /snap/go/10455/src/net/sockoptip_posix.go:7
import (
//line /snap/go/10455/src/net/sockoptip_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/sockoptip_posix.go:7
)
//line /snap/go/10455/src/net/sockoptip_posix.go:7
import (
//line /snap/go/10455/src/net/sockoptip_posix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/sockoptip_posix.go:7
)

import (
	"runtime"
	"syscall"
)

func joinIPv4Group(fd *netFD, ifi *Interface, ip IP) error {
//line /snap/go/10455/src/net/sockoptip_posix.go:14
	_go_fuzz_dep_.CoverTab[8265]++
							mreq := &syscall.IPMreq{Multiaddr: [4]byte{ip[0], ip[1], ip[2], ip[3]}}
							if err := setIPv4MreqToInterface(mreq, ifi); err != nil {
//line /snap/go/10455/src/net/sockoptip_posix.go:16
		_go_fuzz_dep_.CoverTab[529789]++
//line /snap/go/10455/src/net/sockoptip_posix.go:16
		_go_fuzz_dep_.CoverTab[8267]++
								return err
//line /snap/go/10455/src/net/sockoptip_posix.go:17
		// _ = "end of CoverTab[8267]"
	} else {
//line /snap/go/10455/src/net/sockoptip_posix.go:18
		_go_fuzz_dep_.CoverTab[529790]++
//line /snap/go/10455/src/net/sockoptip_posix.go:18
		_go_fuzz_dep_.CoverTab[8268]++
//line /snap/go/10455/src/net/sockoptip_posix.go:18
		// _ = "end of CoverTab[8268]"
//line /snap/go/10455/src/net/sockoptip_posix.go:18
	}
//line /snap/go/10455/src/net/sockoptip_posix.go:18
	// _ = "end of CoverTab[8265]"
//line /snap/go/10455/src/net/sockoptip_posix.go:18
	_go_fuzz_dep_.CoverTab[8266]++
							err := fd.pfd.SetsockoptIPMreq(syscall.IPPROTO_IP, syscall.IP_ADD_MEMBERSHIP, mreq)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockoptip_posix.go:21
	// _ = "end of CoverTab[8266]"
}

func setIPv6MulticastInterface(fd *netFD, ifi *Interface) error {
//line /snap/go/10455/src/net/sockoptip_posix.go:24
	_go_fuzz_dep_.CoverTab[8269]++
							var v int
							if ifi != nil {
//line /snap/go/10455/src/net/sockoptip_posix.go:26
		_go_fuzz_dep_.CoverTab[529791]++
//line /snap/go/10455/src/net/sockoptip_posix.go:26
		_go_fuzz_dep_.CoverTab[8271]++
								v = ifi.Index
//line /snap/go/10455/src/net/sockoptip_posix.go:27
		// _ = "end of CoverTab[8271]"
	} else {
//line /snap/go/10455/src/net/sockoptip_posix.go:28
		_go_fuzz_dep_.CoverTab[529792]++
//line /snap/go/10455/src/net/sockoptip_posix.go:28
		_go_fuzz_dep_.CoverTab[8272]++
//line /snap/go/10455/src/net/sockoptip_posix.go:28
		// _ = "end of CoverTab[8272]"
//line /snap/go/10455/src/net/sockoptip_posix.go:28
	}
//line /snap/go/10455/src/net/sockoptip_posix.go:28
	// _ = "end of CoverTab[8269]"
//line /snap/go/10455/src/net/sockoptip_posix.go:28
	_go_fuzz_dep_.CoverTab[8270]++
							err := fd.pfd.SetsockoptInt(syscall.IPPROTO_IPV6, syscall.IPV6_MULTICAST_IF, v)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockoptip_posix.go:31
	// _ = "end of CoverTab[8270]"
}

func setIPv6MulticastLoopback(fd *netFD, v bool) error {
//line /snap/go/10455/src/net/sockoptip_posix.go:34
	_go_fuzz_dep_.CoverTab[8273]++
							err := fd.pfd.SetsockoptInt(syscall.IPPROTO_IPV6, syscall.IPV6_MULTICAST_LOOP, boolint(v))
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockoptip_posix.go:37
	// _ = "end of CoverTab[8273]"
}

func joinIPv6Group(fd *netFD, ifi *Interface, ip IP) error {
//line /snap/go/10455/src/net/sockoptip_posix.go:40
	_go_fuzz_dep_.CoverTab[8274]++
							mreq := &syscall.IPv6Mreq{}
							copy(mreq.Multiaddr[:], ip)
							if ifi != nil {
//line /snap/go/10455/src/net/sockoptip_posix.go:43
		_go_fuzz_dep_.CoverTab[529793]++
//line /snap/go/10455/src/net/sockoptip_posix.go:43
		_go_fuzz_dep_.CoverTab[8276]++
								mreq.Interface = uint32(ifi.Index)
//line /snap/go/10455/src/net/sockoptip_posix.go:44
		// _ = "end of CoverTab[8276]"
	} else {
//line /snap/go/10455/src/net/sockoptip_posix.go:45
		_go_fuzz_dep_.CoverTab[529794]++
//line /snap/go/10455/src/net/sockoptip_posix.go:45
		_go_fuzz_dep_.CoverTab[8277]++
//line /snap/go/10455/src/net/sockoptip_posix.go:45
		// _ = "end of CoverTab[8277]"
//line /snap/go/10455/src/net/sockoptip_posix.go:45
	}
//line /snap/go/10455/src/net/sockoptip_posix.go:45
	// _ = "end of CoverTab[8274]"
//line /snap/go/10455/src/net/sockoptip_posix.go:45
	_go_fuzz_dep_.CoverTab[8275]++
							err := fd.pfd.SetsockoptIPv6Mreq(syscall.IPPROTO_IPV6, syscall.IPV6_JOIN_GROUP, mreq)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockoptip_posix.go:48
	// _ = "end of CoverTab[8275]"
}

//line /snap/go/10455/src/net/sockoptip_posix.go:49
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/sockoptip_posix.go:49
var _ = _go_fuzz_dep_.CoverTab
