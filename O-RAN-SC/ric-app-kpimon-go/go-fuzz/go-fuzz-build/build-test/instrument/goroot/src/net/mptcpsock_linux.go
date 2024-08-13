// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/mptcpsock_linux.go:5
package net

//line /snap/go/10455/src/net/mptcpsock_linux.go:5
import (
//line /snap/go/10455/src/net/mptcpsock_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/mptcpsock_linux.go:5
)
//line /snap/go/10455/src/net/mptcpsock_linux.go:5
import (
//line /snap/go/10455/src/net/mptcpsock_linux.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/mptcpsock_linux.go:5
)

import (
	"context"
	"errors"
	"internal/poll"
	"internal/syscall/unix"
	"sync"
	"syscall"
)

var (
	mptcpOnce	sync.Once
	mptcpAvailable	bool
	hasSOLMPTCP	bool
)

// These constants aren't in the syscall package, which is frozen
const (
	_IPPROTO_MPTCP	= 0x106
	_SOL_MPTCP	= 0x11c
	_MPTCP_INFO	= 0x1
)

func supportsMultipathTCP() bool {
//line /snap/go/10455/src/net/mptcpsock_linux.go:29
	_go_fuzz_dep_.CoverTab[7365]++
							mptcpOnce.Do(initMPTCPavailable)
							return mptcpAvailable
//line /snap/go/10455/src/net/mptcpsock_linux.go:31
	// _ = "end of CoverTab[7365]"
}

// Check that MPTCP is supported by attempting to create an MPTCP socket and by
//line /snap/go/10455/src/net/mptcpsock_linux.go:34
// looking at the returned error if any.
//line /snap/go/10455/src/net/mptcpsock_linux.go:36
func initMPTCPavailable() {
//line /snap/go/10455/src/net/mptcpsock_linux.go:36
	_go_fuzz_dep_.CoverTab[7366]++
							s, err := sysSocket(syscall.AF_INET, syscall.SOCK_STREAM, _IPPROTO_MPTCP)
							switch {
	case errors.Is(err, syscall.EPROTONOSUPPORT):
//line /snap/go/10455/src/net/mptcpsock_linux.go:39
		_go_fuzz_dep_.CoverTab[529226]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:39
		_go_fuzz_dep_.CoverTab[7368]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:39
		// _ = "end of CoverTab[7368]"
	case errors.Is(err, syscall.EINVAL):
//line /snap/go/10455/src/net/mptcpsock_linux.go:40
		_go_fuzz_dep_.CoverTab[529227]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:40
		_go_fuzz_dep_.CoverTab[7369]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:40
		// _ = "end of CoverTab[7369]"
	case err == nil:
//line /snap/go/10455/src/net/mptcpsock_linux.go:41
		_go_fuzz_dep_.CoverTab[529228]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:41
		_go_fuzz_dep_.CoverTab[7370]++
								poll.CloseFunc(s)
								fallthrough
//line /snap/go/10455/src/net/mptcpsock_linux.go:43
		// _ = "end of CoverTab[7370]"
	default:
//line /snap/go/10455/src/net/mptcpsock_linux.go:44
		_go_fuzz_dep_.CoverTab[529229]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:44
		_go_fuzz_dep_.CoverTab[7371]++

								mptcpAvailable = true
//line /snap/go/10455/src/net/mptcpsock_linux.go:46
		// _ = "end of CoverTab[7371]"
	}
//line /snap/go/10455/src/net/mptcpsock_linux.go:47
	// _ = "end of CoverTab[7366]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:47
	_go_fuzz_dep_.CoverTab[7367]++

							major, minor := unix.KernelVersion()

							hasSOLMPTCP = major > 5 || func() bool {
//line /snap/go/10455/src/net/mptcpsock_linux.go:51
		_go_fuzz_dep_.CoverTab[7372]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:51
		return (major == 5 && func() bool {
//line /snap/go/10455/src/net/mptcpsock_linux.go:51
			_go_fuzz_dep_.CoverTab[7373]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:51
			return minor >= 16
//line /snap/go/10455/src/net/mptcpsock_linux.go:51
			// _ = "end of CoverTab[7373]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:51
		}())
//line /snap/go/10455/src/net/mptcpsock_linux.go:51
		// _ = "end of CoverTab[7372]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:51
	}()
//line /snap/go/10455/src/net/mptcpsock_linux.go:51
	// _ = "end of CoverTab[7367]"
}

func (sd *sysDialer) dialMPTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
//line /snap/go/10455/src/net/mptcpsock_linux.go:54
	_go_fuzz_dep_.CoverTab[7374]++
							if supportsMultipathTCP() {
//line /snap/go/10455/src/net/mptcpsock_linux.go:55
		_go_fuzz_dep_.CoverTab[529230]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:55
		_go_fuzz_dep_.CoverTab[7376]++
								if conn, err := sd.doDialTCPProto(ctx, laddr, raddr, _IPPROTO_MPTCP); err == nil {
//line /snap/go/10455/src/net/mptcpsock_linux.go:56
			_go_fuzz_dep_.CoverTab[529232]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:56
			_go_fuzz_dep_.CoverTab[7377]++
									return conn, nil
//line /snap/go/10455/src/net/mptcpsock_linux.go:57
			// _ = "end of CoverTab[7377]"
		} else {
//line /snap/go/10455/src/net/mptcpsock_linux.go:58
			_go_fuzz_dep_.CoverTab[529233]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:58
			_go_fuzz_dep_.CoverTab[7378]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:58
			// _ = "end of CoverTab[7378]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:58
		}
//line /snap/go/10455/src/net/mptcpsock_linux.go:58
		// _ = "end of CoverTab[7376]"
	} else {
//line /snap/go/10455/src/net/mptcpsock_linux.go:59
		_go_fuzz_dep_.CoverTab[529231]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:59
		_go_fuzz_dep_.CoverTab[7379]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:59
		// _ = "end of CoverTab[7379]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:59
	}
//line /snap/go/10455/src/net/mptcpsock_linux.go:59
	// _ = "end of CoverTab[7374]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:59
	_go_fuzz_dep_.CoverTab[7375]++

//line /snap/go/10455/src/net/mptcpsock_linux.go:67
	return sd.dialTCP(ctx, laddr, raddr)
//line /snap/go/10455/src/net/mptcpsock_linux.go:67
	// _ = "end of CoverTab[7375]"
}

func (sl *sysListener) listenMPTCP(ctx context.Context, laddr *TCPAddr) (*TCPListener, error) {
//line /snap/go/10455/src/net/mptcpsock_linux.go:70
	_go_fuzz_dep_.CoverTab[7380]++
							if supportsMultipathTCP() {
//line /snap/go/10455/src/net/mptcpsock_linux.go:71
		_go_fuzz_dep_.CoverTab[529234]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:71
		_go_fuzz_dep_.CoverTab[7382]++
								if dial, err := sl.listenTCPProto(ctx, laddr, _IPPROTO_MPTCP); err == nil {
//line /snap/go/10455/src/net/mptcpsock_linux.go:72
			_go_fuzz_dep_.CoverTab[529236]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:72
			_go_fuzz_dep_.CoverTab[7383]++
									return dial, nil
//line /snap/go/10455/src/net/mptcpsock_linux.go:73
			// _ = "end of CoverTab[7383]"
		} else {
//line /snap/go/10455/src/net/mptcpsock_linux.go:74
			_go_fuzz_dep_.CoverTab[529237]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:74
			_go_fuzz_dep_.CoverTab[7384]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:74
			// _ = "end of CoverTab[7384]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:74
		}
//line /snap/go/10455/src/net/mptcpsock_linux.go:74
		// _ = "end of CoverTab[7382]"
	} else {
//line /snap/go/10455/src/net/mptcpsock_linux.go:75
		_go_fuzz_dep_.CoverTab[529235]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:75
		_go_fuzz_dep_.CoverTab[7385]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:75
		// _ = "end of CoverTab[7385]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:75
	}
//line /snap/go/10455/src/net/mptcpsock_linux.go:75
	// _ = "end of CoverTab[7380]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:75
	_go_fuzz_dep_.CoverTab[7381]++

//line /snap/go/10455/src/net/mptcpsock_linux.go:83
	return sl.listenTCP(ctx, laddr)
//line /snap/go/10455/src/net/mptcpsock_linux.go:83
	// _ = "end of CoverTab[7381]"
}

// hasFallenBack reports whether the MPTCP connection has fallen back to "plain"
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
// TCP.
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
//
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
// A connection can fallback to TCP for different reasons, e.g. the other peer
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
// doesn't support it, a middle box "accidentally" drops the option, etc.
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
//
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
// If the MPTCP protocol has not been requested when creating the socket, this
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
// method will return true: MPTCP is not being used.
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
//
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
// Kernel >= 5.16 returns EOPNOTSUPP/ENOPROTOOPT in case of fallback.
//line /snap/go/10455/src/net/mptcpsock_linux.go:86
// Older kernels will always return them even if MPTCP is used: not usable.
//line /snap/go/10455/src/net/mptcpsock_linux.go:97
func hasFallenBack(fd *netFD) bool {
//line /snap/go/10455/src/net/mptcpsock_linux.go:97
	_go_fuzz_dep_.CoverTab[7386]++
							_, err := fd.pfd.GetsockoptInt(_SOL_MPTCP, _MPTCP_INFO)

//line /snap/go/10455/src/net/mptcpsock_linux.go:103
	return err == syscall.EOPNOTSUPP || func() bool {
//line /snap/go/10455/src/net/mptcpsock_linux.go:103
		_go_fuzz_dep_.CoverTab[7387]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:103
		return err == syscall.ENOPROTOOPT
//line /snap/go/10455/src/net/mptcpsock_linux.go:103
		// _ = "end of CoverTab[7387]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:103
	}()
//line /snap/go/10455/src/net/mptcpsock_linux.go:103
	// _ = "end of CoverTab[7386]"
}

// isUsingMPTCPProto reports whether the socket protocol is MPTCP.
//line /snap/go/10455/src/net/mptcpsock_linux.go:106
//
//line /snap/go/10455/src/net/mptcpsock_linux.go:106
// Compared to hasFallenBack method, here only the socket protocol being used is
//line /snap/go/10455/src/net/mptcpsock_linux.go:106
// checked: it can be MPTCP but it doesn't mean MPTCP is used on the wire, maybe
//line /snap/go/10455/src/net/mptcpsock_linux.go:106
// a fallback to TCP has been done.
//line /snap/go/10455/src/net/mptcpsock_linux.go:111
func isUsingMPTCPProto(fd *netFD) bool {
//line /snap/go/10455/src/net/mptcpsock_linux.go:111
	_go_fuzz_dep_.CoverTab[7388]++
							proto, _ := fd.pfd.GetsockoptInt(syscall.SOL_SOCKET, syscall.SO_PROTOCOL)

							return proto == _IPPROTO_MPTCP
//line /snap/go/10455/src/net/mptcpsock_linux.go:114
	// _ = "end of CoverTab[7388]"
}

// isUsingMultipathTCP reports whether MPTCP is still being used.
//line /snap/go/10455/src/net/mptcpsock_linux.go:117
//
//line /snap/go/10455/src/net/mptcpsock_linux.go:117
// Please look at the description of hasFallenBack (kernel >=5.16) and
//line /snap/go/10455/src/net/mptcpsock_linux.go:117
// isUsingMPTCPProto methods for more details about what is being checked here.
//line /snap/go/10455/src/net/mptcpsock_linux.go:121
func isUsingMultipathTCP(fd *netFD) bool {
//line /snap/go/10455/src/net/mptcpsock_linux.go:121
	_go_fuzz_dep_.CoverTab[7389]++
							if hasSOLMPTCP {
//line /snap/go/10455/src/net/mptcpsock_linux.go:122
		_go_fuzz_dep_.CoverTab[529238]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:122
		_go_fuzz_dep_.CoverTab[7391]++
								return !hasFallenBack(fd)
//line /snap/go/10455/src/net/mptcpsock_linux.go:123
		// _ = "end of CoverTab[7391]"
	} else {
//line /snap/go/10455/src/net/mptcpsock_linux.go:124
		_go_fuzz_dep_.CoverTab[529239]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:124
		_go_fuzz_dep_.CoverTab[7392]++
//line /snap/go/10455/src/net/mptcpsock_linux.go:124
		// _ = "end of CoverTab[7392]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:124
	}
//line /snap/go/10455/src/net/mptcpsock_linux.go:124
	// _ = "end of CoverTab[7389]"
//line /snap/go/10455/src/net/mptcpsock_linux.go:124
	_go_fuzz_dep_.CoverTab[7390]++

							return isUsingMPTCPProto(fd)
//line /snap/go/10455/src/net/mptcpsock_linux.go:126
	// _ = "end of CoverTab[7390]"
}

//line /snap/go/10455/src/net/mptcpsock_linux.go:127
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/mptcpsock_linux.go:127
var _ = _go_fuzz_dep_.CoverTab
