// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

//line /snap/go/10455/src/net/unixsock_posix.go:7
package net

//line /snap/go/10455/src/net/unixsock_posix.go:7
import (
//line /snap/go/10455/src/net/unixsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/unixsock_posix.go:7
)
//line /snap/go/10455/src/net/unixsock_posix.go:7
import (
//line /snap/go/10455/src/net/unixsock_posix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/unixsock_posix.go:7
)

import (
	"context"
	"errors"
	"os"
	"syscall"
)

func unixSocket(ctx context.Context, net string, laddr, raddr sockaddr, mode string, ctxCtrlFn func(context.Context, string, string, syscall.RawConn) error) (*netFD, error) {
//line /snap/go/10455/src/net/unixsock_posix.go:16
	_go_fuzz_dep_.CoverTab[8993]++
							var sotype int
							switch net {
	case "unix":
//line /snap/go/10455/src/net/unixsock_posix.go:19
		_go_fuzz_dep_.CoverTab[530195]++
//line /snap/go/10455/src/net/unixsock_posix.go:19
		_go_fuzz_dep_.CoverTab[8997]++
								sotype = syscall.SOCK_STREAM
//line /snap/go/10455/src/net/unixsock_posix.go:20
		// _ = "end of CoverTab[8997]"
	case "unixgram":
//line /snap/go/10455/src/net/unixsock_posix.go:21
		_go_fuzz_dep_.CoverTab[530196]++
//line /snap/go/10455/src/net/unixsock_posix.go:21
		_go_fuzz_dep_.CoverTab[8998]++
								sotype = syscall.SOCK_DGRAM
//line /snap/go/10455/src/net/unixsock_posix.go:22
		// _ = "end of CoverTab[8998]"
	case "unixpacket":
//line /snap/go/10455/src/net/unixsock_posix.go:23
		_go_fuzz_dep_.CoverTab[530197]++
//line /snap/go/10455/src/net/unixsock_posix.go:23
		_go_fuzz_dep_.CoverTab[8999]++
								sotype = syscall.SOCK_SEQPACKET
//line /snap/go/10455/src/net/unixsock_posix.go:24
		// _ = "end of CoverTab[8999]"
	default:
//line /snap/go/10455/src/net/unixsock_posix.go:25
		_go_fuzz_dep_.CoverTab[530198]++
//line /snap/go/10455/src/net/unixsock_posix.go:25
		_go_fuzz_dep_.CoverTab[9000]++
								return nil, UnknownNetworkError(net)
//line /snap/go/10455/src/net/unixsock_posix.go:26
		// _ = "end of CoverTab[9000]"
	}
//line /snap/go/10455/src/net/unixsock_posix.go:27
	// _ = "end of CoverTab[8993]"
//line /snap/go/10455/src/net/unixsock_posix.go:27
	_go_fuzz_dep_.CoverTab[8994]++

							switch mode {
	case "dial":
//line /snap/go/10455/src/net/unixsock_posix.go:30
		_go_fuzz_dep_.CoverTab[530199]++
//line /snap/go/10455/src/net/unixsock_posix.go:30
		_go_fuzz_dep_.CoverTab[9001]++
								if laddr != nil && func() bool {
//line /snap/go/10455/src/net/unixsock_posix.go:31
			_go_fuzz_dep_.CoverTab[9006]++
//line /snap/go/10455/src/net/unixsock_posix.go:31
			return laddr.isWildcard()
//line /snap/go/10455/src/net/unixsock_posix.go:31
			// _ = "end of CoverTab[9006]"
//line /snap/go/10455/src/net/unixsock_posix.go:31
		}() {
//line /snap/go/10455/src/net/unixsock_posix.go:31
			_go_fuzz_dep_.CoverTab[530202]++
//line /snap/go/10455/src/net/unixsock_posix.go:31
			_go_fuzz_dep_.CoverTab[9007]++
									laddr = nil
//line /snap/go/10455/src/net/unixsock_posix.go:32
			// _ = "end of CoverTab[9007]"
		} else {
//line /snap/go/10455/src/net/unixsock_posix.go:33
			_go_fuzz_dep_.CoverTab[530203]++
//line /snap/go/10455/src/net/unixsock_posix.go:33
			_go_fuzz_dep_.CoverTab[9008]++
//line /snap/go/10455/src/net/unixsock_posix.go:33
			// _ = "end of CoverTab[9008]"
//line /snap/go/10455/src/net/unixsock_posix.go:33
		}
//line /snap/go/10455/src/net/unixsock_posix.go:33
		// _ = "end of CoverTab[9001]"
//line /snap/go/10455/src/net/unixsock_posix.go:33
		_go_fuzz_dep_.CoverTab[9002]++
								if raddr != nil && func() bool {
//line /snap/go/10455/src/net/unixsock_posix.go:34
			_go_fuzz_dep_.CoverTab[9009]++
//line /snap/go/10455/src/net/unixsock_posix.go:34
			return raddr.isWildcard()
//line /snap/go/10455/src/net/unixsock_posix.go:34
			// _ = "end of CoverTab[9009]"
//line /snap/go/10455/src/net/unixsock_posix.go:34
		}() {
//line /snap/go/10455/src/net/unixsock_posix.go:34
			_go_fuzz_dep_.CoverTab[530204]++
//line /snap/go/10455/src/net/unixsock_posix.go:34
			_go_fuzz_dep_.CoverTab[9010]++
									raddr = nil
//line /snap/go/10455/src/net/unixsock_posix.go:35
			// _ = "end of CoverTab[9010]"
		} else {
//line /snap/go/10455/src/net/unixsock_posix.go:36
			_go_fuzz_dep_.CoverTab[530205]++
//line /snap/go/10455/src/net/unixsock_posix.go:36
			_go_fuzz_dep_.CoverTab[9011]++
//line /snap/go/10455/src/net/unixsock_posix.go:36
			// _ = "end of CoverTab[9011]"
//line /snap/go/10455/src/net/unixsock_posix.go:36
		}
//line /snap/go/10455/src/net/unixsock_posix.go:36
		// _ = "end of CoverTab[9002]"
//line /snap/go/10455/src/net/unixsock_posix.go:36
		_go_fuzz_dep_.CoverTab[9003]++
								if raddr == nil && func() bool {
//line /snap/go/10455/src/net/unixsock_posix.go:37
			_go_fuzz_dep_.CoverTab[9012]++
//line /snap/go/10455/src/net/unixsock_posix.go:37
			return (sotype != syscall.SOCK_DGRAM || func() bool {
//line /snap/go/10455/src/net/unixsock_posix.go:37
				_go_fuzz_dep_.CoverTab[9013]++
//line /snap/go/10455/src/net/unixsock_posix.go:37
				return laddr == nil
//line /snap/go/10455/src/net/unixsock_posix.go:37
				// _ = "end of CoverTab[9013]"
//line /snap/go/10455/src/net/unixsock_posix.go:37
			}())
//line /snap/go/10455/src/net/unixsock_posix.go:37
			// _ = "end of CoverTab[9012]"
//line /snap/go/10455/src/net/unixsock_posix.go:37
		}() {
//line /snap/go/10455/src/net/unixsock_posix.go:37
			_go_fuzz_dep_.CoverTab[530206]++
//line /snap/go/10455/src/net/unixsock_posix.go:37
			_go_fuzz_dep_.CoverTab[9014]++
									return nil, errMissingAddress
//line /snap/go/10455/src/net/unixsock_posix.go:38
			// _ = "end of CoverTab[9014]"
		} else {
//line /snap/go/10455/src/net/unixsock_posix.go:39
			_go_fuzz_dep_.CoverTab[530207]++
//line /snap/go/10455/src/net/unixsock_posix.go:39
			_go_fuzz_dep_.CoverTab[9015]++
//line /snap/go/10455/src/net/unixsock_posix.go:39
			// _ = "end of CoverTab[9015]"
//line /snap/go/10455/src/net/unixsock_posix.go:39
		}
//line /snap/go/10455/src/net/unixsock_posix.go:39
		// _ = "end of CoverTab[9003]"
	case "listen":
//line /snap/go/10455/src/net/unixsock_posix.go:40
		_go_fuzz_dep_.CoverTab[530200]++
//line /snap/go/10455/src/net/unixsock_posix.go:40
		_go_fuzz_dep_.CoverTab[9004]++
//line /snap/go/10455/src/net/unixsock_posix.go:40
		// _ = "end of CoverTab[9004]"
	default:
//line /snap/go/10455/src/net/unixsock_posix.go:41
		_go_fuzz_dep_.CoverTab[530201]++
//line /snap/go/10455/src/net/unixsock_posix.go:41
		_go_fuzz_dep_.CoverTab[9005]++
								return nil, errors.New("unknown mode: " + mode)
//line /snap/go/10455/src/net/unixsock_posix.go:42
		// _ = "end of CoverTab[9005]"
	}
//line /snap/go/10455/src/net/unixsock_posix.go:43
	// _ = "end of CoverTab[8994]"
//line /snap/go/10455/src/net/unixsock_posix.go:43
	_go_fuzz_dep_.CoverTab[8995]++

							fd, err := socket(ctx, net, syscall.AF_UNIX, sotype, 0, false, laddr, raddr, ctxCtrlFn)
							if err != nil {
//line /snap/go/10455/src/net/unixsock_posix.go:46
		_go_fuzz_dep_.CoverTab[530208]++
//line /snap/go/10455/src/net/unixsock_posix.go:46
		_go_fuzz_dep_.CoverTab[9016]++
								return nil, err
//line /snap/go/10455/src/net/unixsock_posix.go:47
		// _ = "end of CoverTab[9016]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:48
		_go_fuzz_dep_.CoverTab[530209]++
//line /snap/go/10455/src/net/unixsock_posix.go:48
		_go_fuzz_dep_.CoverTab[9017]++
//line /snap/go/10455/src/net/unixsock_posix.go:48
		// _ = "end of CoverTab[9017]"
//line /snap/go/10455/src/net/unixsock_posix.go:48
	}
//line /snap/go/10455/src/net/unixsock_posix.go:48
	// _ = "end of CoverTab[8995]"
//line /snap/go/10455/src/net/unixsock_posix.go:48
	_go_fuzz_dep_.CoverTab[8996]++
							return fd, nil
//line /snap/go/10455/src/net/unixsock_posix.go:49
	// _ = "end of CoverTab[8996]"
}

func sockaddrToUnix(sa syscall.Sockaddr) Addr {
//line /snap/go/10455/src/net/unixsock_posix.go:52
	_go_fuzz_dep_.CoverTab[9018]++
							if s, ok := sa.(*syscall.SockaddrUnix); ok {
//line /snap/go/10455/src/net/unixsock_posix.go:53
		_go_fuzz_dep_.CoverTab[530210]++
//line /snap/go/10455/src/net/unixsock_posix.go:53
		_go_fuzz_dep_.CoverTab[9020]++
								return &UnixAddr{Name: s.Name, Net: "unix"}
//line /snap/go/10455/src/net/unixsock_posix.go:54
		// _ = "end of CoverTab[9020]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:55
		_go_fuzz_dep_.CoverTab[530211]++
//line /snap/go/10455/src/net/unixsock_posix.go:55
		_go_fuzz_dep_.CoverTab[9021]++
//line /snap/go/10455/src/net/unixsock_posix.go:55
		// _ = "end of CoverTab[9021]"
//line /snap/go/10455/src/net/unixsock_posix.go:55
	}
//line /snap/go/10455/src/net/unixsock_posix.go:55
	// _ = "end of CoverTab[9018]"
//line /snap/go/10455/src/net/unixsock_posix.go:55
	_go_fuzz_dep_.CoverTab[9019]++
							return nil
//line /snap/go/10455/src/net/unixsock_posix.go:56
	// _ = "end of CoverTab[9019]"
}

func sockaddrToUnixgram(sa syscall.Sockaddr) Addr {
//line /snap/go/10455/src/net/unixsock_posix.go:59
	_go_fuzz_dep_.CoverTab[9022]++
							if s, ok := sa.(*syscall.SockaddrUnix); ok {
//line /snap/go/10455/src/net/unixsock_posix.go:60
		_go_fuzz_dep_.CoverTab[530212]++
//line /snap/go/10455/src/net/unixsock_posix.go:60
		_go_fuzz_dep_.CoverTab[9024]++
								return &UnixAddr{Name: s.Name, Net: "unixgram"}
//line /snap/go/10455/src/net/unixsock_posix.go:61
		// _ = "end of CoverTab[9024]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:62
		_go_fuzz_dep_.CoverTab[530213]++
//line /snap/go/10455/src/net/unixsock_posix.go:62
		_go_fuzz_dep_.CoverTab[9025]++
//line /snap/go/10455/src/net/unixsock_posix.go:62
		// _ = "end of CoverTab[9025]"
//line /snap/go/10455/src/net/unixsock_posix.go:62
	}
//line /snap/go/10455/src/net/unixsock_posix.go:62
	// _ = "end of CoverTab[9022]"
//line /snap/go/10455/src/net/unixsock_posix.go:62
	_go_fuzz_dep_.CoverTab[9023]++
							return nil
//line /snap/go/10455/src/net/unixsock_posix.go:63
	// _ = "end of CoverTab[9023]"
}

func sockaddrToUnixpacket(sa syscall.Sockaddr) Addr {
//line /snap/go/10455/src/net/unixsock_posix.go:66
	_go_fuzz_dep_.CoverTab[9026]++
							if s, ok := sa.(*syscall.SockaddrUnix); ok {
//line /snap/go/10455/src/net/unixsock_posix.go:67
		_go_fuzz_dep_.CoverTab[530214]++
//line /snap/go/10455/src/net/unixsock_posix.go:67
		_go_fuzz_dep_.CoverTab[9028]++
								return &UnixAddr{Name: s.Name, Net: "unixpacket"}
//line /snap/go/10455/src/net/unixsock_posix.go:68
		// _ = "end of CoverTab[9028]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:69
		_go_fuzz_dep_.CoverTab[530215]++
//line /snap/go/10455/src/net/unixsock_posix.go:69
		_go_fuzz_dep_.CoverTab[9029]++
//line /snap/go/10455/src/net/unixsock_posix.go:69
		// _ = "end of CoverTab[9029]"
//line /snap/go/10455/src/net/unixsock_posix.go:69
	}
//line /snap/go/10455/src/net/unixsock_posix.go:69
	// _ = "end of CoverTab[9026]"
//line /snap/go/10455/src/net/unixsock_posix.go:69
	_go_fuzz_dep_.CoverTab[9027]++
							return nil
//line /snap/go/10455/src/net/unixsock_posix.go:70
	// _ = "end of CoverTab[9027]"
}

func sotypeToNet(sotype int) string {
//line /snap/go/10455/src/net/unixsock_posix.go:73
	_go_fuzz_dep_.CoverTab[9030]++
							switch sotype {
	case syscall.SOCK_STREAM:
//line /snap/go/10455/src/net/unixsock_posix.go:75
		_go_fuzz_dep_.CoverTab[530216]++
//line /snap/go/10455/src/net/unixsock_posix.go:75
		_go_fuzz_dep_.CoverTab[9031]++
								return "unix"
//line /snap/go/10455/src/net/unixsock_posix.go:76
		// _ = "end of CoverTab[9031]"
	case syscall.SOCK_DGRAM:
//line /snap/go/10455/src/net/unixsock_posix.go:77
		_go_fuzz_dep_.CoverTab[530217]++
//line /snap/go/10455/src/net/unixsock_posix.go:77
		_go_fuzz_dep_.CoverTab[9032]++
								return "unixgram"
//line /snap/go/10455/src/net/unixsock_posix.go:78
		// _ = "end of CoverTab[9032]"
	case syscall.SOCK_SEQPACKET:
//line /snap/go/10455/src/net/unixsock_posix.go:79
		_go_fuzz_dep_.CoverTab[530218]++
//line /snap/go/10455/src/net/unixsock_posix.go:79
		_go_fuzz_dep_.CoverTab[9033]++
								return "unixpacket"
//line /snap/go/10455/src/net/unixsock_posix.go:80
		// _ = "end of CoverTab[9033]"
	default:
//line /snap/go/10455/src/net/unixsock_posix.go:81
		_go_fuzz_dep_.CoverTab[530219]++
//line /snap/go/10455/src/net/unixsock_posix.go:81
		_go_fuzz_dep_.CoverTab[9034]++
								panic("sotypeToNet unknown socket type")
//line /snap/go/10455/src/net/unixsock_posix.go:82
		// _ = "end of CoverTab[9034]"
	}
//line /snap/go/10455/src/net/unixsock_posix.go:83
	// _ = "end of CoverTab[9030]"
}

func (a *UnixAddr) family() int {
//line /snap/go/10455/src/net/unixsock_posix.go:86
	_go_fuzz_dep_.CoverTab[9035]++
							return syscall.AF_UNIX
//line /snap/go/10455/src/net/unixsock_posix.go:87
	// _ = "end of CoverTab[9035]"
}

func (a *UnixAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /snap/go/10455/src/net/unixsock_posix.go:90
	_go_fuzz_dep_.CoverTab[9036]++
							if a == nil {
//line /snap/go/10455/src/net/unixsock_posix.go:91
		_go_fuzz_dep_.CoverTab[530220]++
//line /snap/go/10455/src/net/unixsock_posix.go:91
		_go_fuzz_dep_.CoverTab[9038]++
								return nil, nil
//line /snap/go/10455/src/net/unixsock_posix.go:92
		// _ = "end of CoverTab[9038]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:93
		_go_fuzz_dep_.CoverTab[530221]++
//line /snap/go/10455/src/net/unixsock_posix.go:93
		_go_fuzz_dep_.CoverTab[9039]++
//line /snap/go/10455/src/net/unixsock_posix.go:93
		// _ = "end of CoverTab[9039]"
//line /snap/go/10455/src/net/unixsock_posix.go:93
	}
//line /snap/go/10455/src/net/unixsock_posix.go:93
	// _ = "end of CoverTab[9036]"
//line /snap/go/10455/src/net/unixsock_posix.go:93
	_go_fuzz_dep_.CoverTab[9037]++
							return &syscall.SockaddrUnix{Name: a.Name}, nil
//line /snap/go/10455/src/net/unixsock_posix.go:94
	// _ = "end of CoverTab[9037]"
}

func (a *UnixAddr) toLocal(net string) sockaddr {
//line /snap/go/10455/src/net/unixsock_posix.go:97
	_go_fuzz_dep_.CoverTab[9040]++
							return a
//line /snap/go/10455/src/net/unixsock_posix.go:98
	// _ = "end of CoverTab[9040]"
}

func (c *UnixConn) readFrom(b []byte) (int, *UnixAddr, error) {
//line /snap/go/10455/src/net/unixsock_posix.go:101
	_go_fuzz_dep_.CoverTab[9041]++
							var addr *UnixAddr
							n, sa, err := c.fd.readFrom(b)
							switch sa := sa.(type) {
	case *syscall.SockaddrUnix:
//line /snap/go/10455/src/net/unixsock_posix.go:105
		_go_fuzz_dep_.CoverTab[530222]++
//line /snap/go/10455/src/net/unixsock_posix.go:105
		_go_fuzz_dep_.CoverTab[9043]++
								if sa.Name != "" {
//line /snap/go/10455/src/net/unixsock_posix.go:106
			_go_fuzz_dep_.CoverTab[530223]++
//line /snap/go/10455/src/net/unixsock_posix.go:106
			_go_fuzz_dep_.CoverTab[9044]++
									addr = &UnixAddr{Name: sa.Name, Net: sotypeToNet(c.fd.sotype)}
//line /snap/go/10455/src/net/unixsock_posix.go:107
			// _ = "end of CoverTab[9044]"
		} else {
//line /snap/go/10455/src/net/unixsock_posix.go:108
			_go_fuzz_dep_.CoverTab[530224]++
//line /snap/go/10455/src/net/unixsock_posix.go:108
			_go_fuzz_dep_.CoverTab[9045]++
//line /snap/go/10455/src/net/unixsock_posix.go:108
			// _ = "end of CoverTab[9045]"
//line /snap/go/10455/src/net/unixsock_posix.go:108
		}
//line /snap/go/10455/src/net/unixsock_posix.go:108
		// _ = "end of CoverTab[9043]"
	}
//line /snap/go/10455/src/net/unixsock_posix.go:109
	// _ = "end of CoverTab[9041]"
//line /snap/go/10455/src/net/unixsock_posix.go:109
	_go_fuzz_dep_.CoverTab[9042]++
							return n, addr, err
//line /snap/go/10455/src/net/unixsock_posix.go:110
	// _ = "end of CoverTab[9042]"
}

func (c *UnixConn) readMsg(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error) {
//line /snap/go/10455/src/net/unixsock_posix.go:113
	_go_fuzz_dep_.CoverTab[9046]++
							var sa syscall.Sockaddr
							n, oobn, flags, sa, err = c.fd.readMsg(b, oob, readMsgFlags)
							if readMsgFlags == 0 && func() bool {
//line /snap/go/10455/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[9049]++
//line /snap/go/10455/src/net/unixsock_posix.go:116
		return err == nil
//line /snap/go/10455/src/net/unixsock_posix.go:116
		// _ = "end of CoverTab[9049]"
//line /snap/go/10455/src/net/unixsock_posix.go:116
	}() && func() bool {
//line /snap/go/10455/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[9050]++
//line /snap/go/10455/src/net/unixsock_posix.go:116
		return oobn > 0
//line /snap/go/10455/src/net/unixsock_posix.go:116
		// _ = "end of CoverTab[9050]"
//line /snap/go/10455/src/net/unixsock_posix.go:116
	}() {
//line /snap/go/10455/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[530225]++
//line /snap/go/10455/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[9051]++
								setReadMsgCloseOnExec(oob[:oobn])
//line /snap/go/10455/src/net/unixsock_posix.go:117
		// _ = "end of CoverTab[9051]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:118
		_go_fuzz_dep_.CoverTab[530226]++
//line /snap/go/10455/src/net/unixsock_posix.go:118
		_go_fuzz_dep_.CoverTab[9052]++
//line /snap/go/10455/src/net/unixsock_posix.go:118
		// _ = "end of CoverTab[9052]"
//line /snap/go/10455/src/net/unixsock_posix.go:118
	}
//line /snap/go/10455/src/net/unixsock_posix.go:118
	// _ = "end of CoverTab[9046]"
//line /snap/go/10455/src/net/unixsock_posix.go:118
	_go_fuzz_dep_.CoverTab[9047]++

							switch sa := sa.(type) {
	case *syscall.SockaddrUnix:
//line /snap/go/10455/src/net/unixsock_posix.go:121
		_go_fuzz_dep_.CoverTab[530227]++
//line /snap/go/10455/src/net/unixsock_posix.go:121
		_go_fuzz_dep_.CoverTab[9053]++
								if sa.Name != "" {
//line /snap/go/10455/src/net/unixsock_posix.go:122
			_go_fuzz_dep_.CoverTab[530228]++
//line /snap/go/10455/src/net/unixsock_posix.go:122
			_go_fuzz_dep_.CoverTab[9054]++
									addr = &UnixAddr{Name: sa.Name, Net: sotypeToNet(c.fd.sotype)}
//line /snap/go/10455/src/net/unixsock_posix.go:123
			// _ = "end of CoverTab[9054]"
		} else {
//line /snap/go/10455/src/net/unixsock_posix.go:124
			_go_fuzz_dep_.CoverTab[530229]++
//line /snap/go/10455/src/net/unixsock_posix.go:124
			_go_fuzz_dep_.CoverTab[9055]++
//line /snap/go/10455/src/net/unixsock_posix.go:124
			// _ = "end of CoverTab[9055]"
//line /snap/go/10455/src/net/unixsock_posix.go:124
		}
//line /snap/go/10455/src/net/unixsock_posix.go:124
		// _ = "end of CoverTab[9053]"
	}
//line /snap/go/10455/src/net/unixsock_posix.go:125
	// _ = "end of CoverTab[9047]"
//line /snap/go/10455/src/net/unixsock_posix.go:125
	_go_fuzz_dep_.CoverTab[9048]++
							return
//line /snap/go/10455/src/net/unixsock_posix.go:126
	// _ = "end of CoverTab[9048]"
}

func (c *UnixConn) writeTo(b []byte, addr *UnixAddr) (int, error) {
//line /snap/go/10455/src/net/unixsock_posix.go:129
	_go_fuzz_dep_.CoverTab[9056]++
							if c.fd.isConnected {
//line /snap/go/10455/src/net/unixsock_posix.go:130
		_go_fuzz_dep_.CoverTab[530230]++
//line /snap/go/10455/src/net/unixsock_posix.go:130
		_go_fuzz_dep_.CoverTab[9060]++
								return 0, ErrWriteToConnected
//line /snap/go/10455/src/net/unixsock_posix.go:131
		// _ = "end of CoverTab[9060]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:132
		_go_fuzz_dep_.CoverTab[530231]++
//line /snap/go/10455/src/net/unixsock_posix.go:132
		_go_fuzz_dep_.CoverTab[9061]++
//line /snap/go/10455/src/net/unixsock_posix.go:132
		// _ = "end of CoverTab[9061]"
//line /snap/go/10455/src/net/unixsock_posix.go:132
	}
//line /snap/go/10455/src/net/unixsock_posix.go:132
	// _ = "end of CoverTab[9056]"
//line /snap/go/10455/src/net/unixsock_posix.go:132
	_go_fuzz_dep_.CoverTab[9057]++
							if addr == nil {
//line /snap/go/10455/src/net/unixsock_posix.go:133
		_go_fuzz_dep_.CoverTab[530232]++
//line /snap/go/10455/src/net/unixsock_posix.go:133
		_go_fuzz_dep_.CoverTab[9062]++
								return 0, errMissingAddress
//line /snap/go/10455/src/net/unixsock_posix.go:134
		// _ = "end of CoverTab[9062]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:135
		_go_fuzz_dep_.CoverTab[530233]++
//line /snap/go/10455/src/net/unixsock_posix.go:135
		_go_fuzz_dep_.CoverTab[9063]++
//line /snap/go/10455/src/net/unixsock_posix.go:135
		// _ = "end of CoverTab[9063]"
//line /snap/go/10455/src/net/unixsock_posix.go:135
	}
//line /snap/go/10455/src/net/unixsock_posix.go:135
	// _ = "end of CoverTab[9057]"
//line /snap/go/10455/src/net/unixsock_posix.go:135
	_go_fuzz_dep_.CoverTab[9058]++
							if addr.Net != sotypeToNet(c.fd.sotype) {
//line /snap/go/10455/src/net/unixsock_posix.go:136
		_go_fuzz_dep_.CoverTab[530234]++
//line /snap/go/10455/src/net/unixsock_posix.go:136
		_go_fuzz_dep_.CoverTab[9064]++
								return 0, syscall.EAFNOSUPPORT
//line /snap/go/10455/src/net/unixsock_posix.go:137
		// _ = "end of CoverTab[9064]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:138
		_go_fuzz_dep_.CoverTab[530235]++
//line /snap/go/10455/src/net/unixsock_posix.go:138
		_go_fuzz_dep_.CoverTab[9065]++
//line /snap/go/10455/src/net/unixsock_posix.go:138
		// _ = "end of CoverTab[9065]"
//line /snap/go/10455/src/net/unixsock_posix.go:138
	}
//line /snap/go/10455/src/net/unixsock_posix.go:138
	// _ = "end of CoverTab[9058]"
//line /snap/go/10455/src/net/unixsock_posix.go:138
	_go_fuzz_dep_.CoverTab[9059]++
							sa := &syscall.SockaddrUnix{Name: addr.Name}
							return c.fd.writeTo(b, sa)
//line /snap/go/10455/src/net/unixsock_posix.go:140
	// _ = "end of CoverTab[9059]"
}

func (c *UnixConn) writeMsg(b, oob []byte, addr *UnixAddr) (n, oobn int, err error) {
//line /snap/go/10455/src/net/unixsock_posix.go:143
	_go_fuzz_dep_.CoverTab[9066]++
							if c.fd.sotype == syscall.SOCK_DGRAM && func() bool {
//line /snap/go/10455/src/net/unixsock_posix.go:144
		_go_fuzz_dep_.CoverTab[9069]++
//line /snap/go/10455/src/net/unixsock_posix.go:144
		return c.fd.isConnected
//line /snap/go/10455/src/net/unixsock_posix.go:144
		// _ = "end of CoverTab[9069]"
//line /snap/go/10455/src/net/unixsock_posix.go:144
	}() {
//line /snap/go/10455/src/net/unixsock_posix.go:144
		_go_fuzz_dep_.CoverTab[530236]++
//line /snap/go/10455/src/net/unixsock_posix.go:144
		_go_fuzz_dep_.CoverTab[9070]++
								return 0, 0, ErrWriteToConnected
//line /snap/go/10455/src/net/unixsock_posix.go:145
		// _ = "end of CoverTab[9070]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:146
		_go_fuzz_dep_.CoverTab[530237]++
//line /snap/go/10455/src/net/unixsock_posix.go:146
		_go_fuzz_dep_.CoverTab[9071]++
//line /snap/go/10455/src/net/unixsock_posix.go:146
		// _ = "end of CoverTab[9071]"
//line /snap/go/10455/src/net/unixsock_posix.go:146
	}
//line /snap/go/10455/src/net/unixsock_posix.go:146
	// _ = "end of CoverTab[9066]"
//line /snap/go/10455/src/net/unixsock_posix.go:146
	_go_fuzz_dep_.CoverTab[9067]++
							var sa syscall.Sockaddr
							if addr != nil {
//line /snap/go/10455/src/net/unixsock_posix.go:148
		_go_fuzz_dep_.CoverTab[530238]++
//line /snap/go/10455/src/net/unixsock_posix.go:148
		_go_fuzz_dep_.CoverTab[9072]++
								if addr.Net != sotypeToNet(c.fd.sotype) {
//line /snap/go/10455/src/net/unixsock_posix.go:149
			_go_fuzz_dep_.CoverTab[530240]++
//line /snap/go/10455/src/net/unixsock_posix.go:149
			_go_fuzz_dep_.CoverTab[9074]++
									return 0, 0, syscall.EAFNOSUPPORT
//line /snap/go/10455/src/net/unixsock_posix.go:150
			// _ = "end of CoverTab[9074]"
		} else {
//line /snap/go/10455/src/net/unixsock_posix.go:151
			_go_fuzz_dep_.CoverTab[530241]++
//line /snap/go/10455/src/net/unixsock_posix.go:151
			_go_fuzz_dep_.CoverTab[9075]++
//line /snap/go/10455/src/net/unixsock_posix.go:151
			// _ = "end of CoverTab[9075]"
//line /snap/go/10455/src/net/unixsock_posix.go:151
		}
//line /snap/go/10455/src/net/unixsock_posix.go:151
		// _ = "end of CoverTab[9072]"
//line /snap/go/10455/src/net/unixsock_posix.go:151
		_go_fuzz_dep_.CoverTab[9073]++
								sa = &syscall.SockaddrUnix{Name: addr.Name}
//line /snap/go/10455/src/net/unixsock_posix.go:152
		// _ = "end of CoverTab[9073]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:153
		_go_fuzz_dep_.CoverTab[530239]++
//line /snap/go/10455/src/net/unixsock_posix.go:153
		_go_fuzz_dep_.CoverTab[9076]++
//line /snap/go/10455/src/net/unixsock_posix.go:153
		// _ = "end of CoverTab[9076]"
//line /snap/go/10455/src/net/unixsock_posix.go:153
	}
//line /snap/go/10455/src/net/unixsock_posix.go:153
	// _ = "end of CoverTab[9067]"
//line /snap/go/10455/src/net/unixsock_posix.go:153
	_go_fuzz_dep_.CoverTab[9068]++
							return c.fd.writeMsg(b, oob, sa)
//line /snap/go/10455/src/net/unixsock_posix.go:154
	// _ = "end of CoverTab[9068]"
}

func (sd *sysDialer) dialUnix(ctx context.Context, laddr, raddr *UnixAddr) (*UnixConn, error) {
//line /snap/go/10455/src/net/unixsock_posix.go:157
	_go_fuzz_dep_.CoverTab[9077]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /snap/go/10455/src/net/unixsock_posix.go:159
		_go_fuzz_dep_.CoverTab[9080]++
//line /snap/go/10455/src/net/unixsock_posix.go:159
		return sd.Dialer.Control != nil
//line /snap/go/10455/src/net/unixsock_posix.go:159
		// _ = "end of CoverTab[9080]"
//line /snap/go/10455/src/net/unixsock_posix.go:159
	}() {
//line /snap/go/10455/src/net/unixsock_posix.go:159
		_go_fuzz_dep_.CoverTab[530242]++
//line /snap/go/10455/src/net/unixsock_posix.go:159
		_go_fuzz_dep_.CoverTab[9081]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/unixsock_posix.go:160
			_go_fuzz_dep_.CoverTab[9082]++
									return sd.Dialer.Control(network, address, c)
//line /snap/go/10455/src/net/unixsock_posix.go:161
			// _ = "end of CoverTab[9082]"
		}
//line /snap/go/10455/src/net/unixsock_posix.go:162
		// _ = "end of CoverTab[9081]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:163
		_go_fuzz_dep_.CoverTab[530243]++
//line /snap/go/10455/src/net/unixsock_posix.go:163
		_go_fuzz_dep_.CoverTab[9083]++
//line /snap/go/10455/src/net/unixsock_posix.go:163
		// _ = "end of CoverTab[9083]"
//line /snap/go/10455/src/net/unixsock_posix.go:163
	}
//line /snap/go/10455/src/net/unixsock_posix.go:163
	// _ = "end of CoverTab[9077]"
//line /snap/go/10455/src/net/unixsock_posix.go:163
	_go_fuzz_dep_.CoverTab[9078]++
							fd, err := unixSocket(ctx, sd.network, laddr, raddr, "dial", ctrlCtxFn)
							if err != nil {
//line /snap/go/10455/src/net/unixsock_posix.go:165
		_go_fuzz_dep_.CoverTab[530244]++
//line /snap/go/10455/src/net/unixsock_posix.go:165
		_go_fuzz_dep_.CoverTab[9084]++
								return nil, err
//line /snap/go/10455/src/net/unixsock_posix.go:166
		// _ = "end of CoverTab[9084]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:167
		_go_fuzz_dep_.CoverTab[530245]++
//line /snap/go/10455/src/net/unixsock_posix.go:167
		_go_fuzz_dep_.CoverTab[9085]++
//line /snap/go/10455/src/net/unixsock_posix.go:167
		// _ = "end of CoverTab[9085]"
//line /snap/go/10455/src/net/unixsock_posix.go:167
	}
//line /snap/go/10455/src/net/unixsock_posix.go:167
	// _ = "end of CoverTab[9078]"
//line /snap/go/10455/src/net/unixsock_posix.go:167
	_go_fuzz_dep_.CoverTab[9079]++
							return newUnixConn(fd), nil
//line /snap/go/10455/src/net/unixsock_posix.go:168
	// _ = "end of CoverTab[9079]"
}

func (ln *UnixListener) accept() (*UnixConn, error) {
//line /snap/go/10455/src/net/unixsock_posix.go:171
	_go_fuzz_dep_.CoverTab[9086]++
							fd, err := ln.fd.accept()
							if err != nil {
//line /snap/go/10455/src/net/unixsock_posix.go:173
		_go_fuzz_dep_.CoverTab[530246]++
//line /snap/go/10455/src/net/unixsock_posix.go:173
		_go_fuzz_dep_.CoverTab[9088]++
								return nil, err
//line /snap/go/10455/src/net/unixsock_posix.go:174
		// _ = "end of CoverTab[9088]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:175
		_go_fuzz_dep_.CoverTab[530247]++
//line /snap/go/10455/src/net/unixsock_posix.go:175
		_go_fuzz_dep_.CoverTab[9089]++
//line /snap/go/10455/src/net/unixsock_posix.go:175
		// _ = "end of CoverTab[9089]"
//line /snap/go/10455/src/net/unixsock_posix.go:175
	}
//line /snap/go/10455/src/net/unixsock_posix.go:175
	// _ = "end of CoverTab[9086]"
//line /snap/go/10455/src/net/unixsock_posix.go:175
	_go_fuzz_dep_.CoverTab[9087]++
							return newUnixConn(fd), nil
//line /snap/go/10455/src/net/unixsock_posix.go:176
	// _ = "end of CoverTab[9087]"
}

func (ln *UnixListener) close() error {
//line /snap/go/10455/src/net/unixsock_posix.go:179
	_go_fuzz_dep_.CoverTab[9090]++

//line /snap/go/10455/src/net/unixsock_posix.go:191
	ln.unlinkOnce.Do(func() {
//line /snap/go/10455/src/net/unixsock_posix.go:191
		_go_fuzz_dep_.CoverTab[9092]++
								if ln.path[0] != '@' && func() bool {
//line /snap/go/10455/src/net/unixsock_posix.go:192
			_go_fuzz_dep_.CoverTab[9093]++
//line /snap/go/10455/src/net/unixsock_posix.go:192
			return ln.unlink
//line /snap/go/10455/src/net/unixsock_posix.go:192
			// _ = "end of CoverTab[9093]"
//line /snap/go/10455/src/net/unixsock_posix.go:192
		}() {
//line /snap/go/10455/src/net/unixsock_posix.go:192
			_go_fuzz_dep_.CoverTab[530248]++
//line /snap/go/10455/src/net/unixsock_posix.go:192
			_go_fuzz_dep_.CoverTab[9094]++
									syscall.Unlink(ln.path)
//line /snap/go/10455/src/net/unixsock_posix.go:193
			// _ = "end of CoverTab[9094]"
		} else {
//line /snap/go/10455/src/net/unixsock_posix.go:194
			_go_fuzz_dep_.CoverTab[530249]++
//line /snap/go/10455/src/net/unixsock_posix.go:194
			_go_fuzz_dep_.CoverTab[9095]++
//line /snap/go/10455/src/net/unixsock_posix.go:194
			// _ = "end of CoverTab[9095]"
//line /snap/go/10455/src/net/unixsock_posix.go:194
		}
//line /snap/go/10455/src/net/unixsock_posix.go:194
		// _ = "end of CoverTab[9092]"
	})
//line /snap/go/10455/src/net/unixsock_posix.go:195
	// _ = "end of CoverTab[9090]"
//line /snap/go/10455/src/net/unixsock_posix.go:195
	_go_fuzz_dep_.CoverTab[9091]++
							return ln.fd.Close()
//line /snap/go/10455/src/net/unixsock_posix.go:196
	// _ = "end of CoverTab[9091]"
}

func (ln *UnixListener) file() (*os.File, error) {
//line /snap/go/10455/src/net/unixsock_posix.go:199
	_go_fuzz_dep_.CoverTab[9096]++
							f, err := ln.fd.dup()
							if err != nil {
//line /snap/go/10455/src/net/unixsock_posix.go:201
		_go_fuzz_dep_.CoverTab[530250]++
//line /snap/go/10455/src/net/unixsock_posix.go:201
		_go_fuzz_dep_.CoverTab[9098]++
								return nil, err
//line /snap/go/10455/src/net/unixsock_posix.go:202
		// _ = "end of CoverTab[9098]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:203
		_go_fuzz_dep_.CoverTab[530251]++
//line /snap/go/10455/src/net/unixsock_posix.go:203
		_go_fuzz_dep_.CoverTab[9099]++
//line /snap/go/10455/src/net/unixsock_posix.go:203
		// _ = "end of CoverTab[9099]"
//line /snap/go/10455/src/net/unixsock_posix.go:203
	}
//line /snap/go/10455/src/net/unixsock_posix.go:203
	// _ = "end of CoverTab[9096]"
//line /snap/go/10455/src/net/unixsock_posix.go:203
	_go_fuzz_dep_.CoverTab[9097]++
							return f, nil
//line /snap/go/10455/src/net/unixsock_posix.go:204
	// _ = "end of CoverTab[9097]"
}

// SetUnlinkOnClose sets whether the underlying socket file should be removed
//line /snap/go/10455/src/net/unixsock_posix.go:207
// from the file system when the listener is closed.
//line /snap/go/10455/src/net/unixsock_posix.go:207
//
//line /snap/go/10455/src/net/unixsock_posix.go:207
// The default behavior is to unlink the socket file only when package net created it.
//line /snap/go/10455/src/net/unixsock_posix.go:207
// That is, when the listener and the underlying socket file were created by a call to
//line /snap/go/10455/src/net/unixsock_posix.go:207
// Listen or ListenUnix, then by default closing the listener will remove the socket file.
//line /snap/go/10455/src/net/unixsock_posix.go:207
// but if the listener was created by a call to FileListener to use an already existing
//line /snap/go/10455/src/net/unixsock_posix.go:207
// socket file, then by default closing the listener will not remove the socket file.
//line /snap/go/10455/src/net/unixsock_posix.go:215
func (l *UnixListener) SetUnlinkOnClose(unlink bool) {
//line /snap/go/10455/src/net/unixsock_posix.go:215
	_go_fuzz_dep_.CoverTab[9100]++
							l.unlink = unlink
//line /snap/go/10455/src/net/unixsock_posix.go:216
	// _ = "end of CoverTab[9100]"
}

func (sl *sysListener) listenUnix(ctx context.Context, laddr *UnixAddr) (*UnixListener, error) {
//line /snap/go/10455/src/net/unixsock_posix.go:219
	_go_fuzz_dep_.CoverTab[9101]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /snap/go/10455/src/net/unixsock_posix.go:221
		_go_fuzz_dep_.CoverTab[530252]++
//line /snap/go/10455/src/net/unixsock_posix.go:221
		_go_fuzz_dep_.CoverTab[9104]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/unixsock_posix.go:222
			_go_fuzz_dep_.CoverTab[9105]++
									return sl.ListenConfig.Control(network, address, c)
//line /snap/go/10455/src/net/unixsock_posix.go:223
			// _ = "end of CoverTab[9105]"
		}
//line /snap/go/10455/src/net/unixsock_posix.go:224
		// _ = "end of CoverTab[9104]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:225
		_go_fuzz_dep_.CoverTab[530253]++
//line /snap/go/10455/src/net/unixsock_posix.go:225
		_go_fuzz_dep_.CoverTab[9106]++
//line /snap/go/10455/src/net/unixsock_posix.go:225
		// _ = "end of CoverTab[9106]"
//line /snap/go/10455/src/net/unixsock_posix.go:225
	}
//line /snap/go/10455/src/net/unixsock_posix.go:225
	// _ = "end of CoverTab[9101]"
//line /snap/go/10455/src/net/unixsock_posix.go:225
	_go_fuzz_dep_.CoverTab[9102]++
							fd, err := unixSocket(ctx, sl.network, laddr, nil, "listen", ctrlCtxFn)
							if err != nil {
//line /snap/go/10455/src/net/unixsock_posix.go:227
		_go_fuzz_dep_.CoverTab[530254]++
//line /snap/go/10455/src/net/unixsock_posix.go:227
		_go_fuzz_dep_.CoverTab[9107]++
								return nil, err
//line /snap/go/10455/src/net/unixsock_posix.go:228
		// _ = "end of CoverTab[9107]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:229
		_go_fuzz_dep_.CoverTab[530255]++
//line /snap/go/10455/src/net/unixsock_posix.go:229
		_go_fuzz_dep_.CoverTab[9108]++
//line /snap/go/10455/src/net/unixsock_posix.go:229
		// _ = "end of CoverTab[9108]"
//line /snap/go/10455/src/net/unixsock_posix.go:229
	}
//line /snap/go/10455/src/net/unixsock_posix.go:229
	// _ = "end of CoverTab[9102]"
//line /snap/go/10455/src/net/unixsock_posix.go:229
	_go_fuzz_dep_.CoverTab[9103]++
							return &UnixListener{fd: fd, path: fd.laddr.String(), unlink: true}, nil
//line /snap/go/10455/src/net/unixsock_posix.go:230
	// _ = "end of CoverTab[9103]"
}

func (sl *sysListener) listenUnixgram(ctx context.Context, laddr *UnixAddr) (*UnixConn, error) {
//line /snap/go/10455/src/net/unixsock_posix.go:233
	_go_fuzz_dep_.CoverTab[9109]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /snap/go/10455/src/net/unixsock_posix.go:235
		_go_fuzz_dep_.CoverTab[530256]++
//line /snap/go/10455/src/net/unixsock_posix.go:235
		_go_fuzz_dep_.CoverTab[9112]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /snap/go/10455/src/net/unixsock_posix.go:236
			_go_fuzz_dep_.CoverTab[9113]++
									return sl.ListenConfig.Control(network, address, c)
//line /snap/go/10455/src/net/unixsock_posix.go:237
			// _ = "end of CoverTab[9113]"
		}
//line /snap/go/10455/src/net/unixsock_posix.go:238
		// _ = "end of CoverTab[9112]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:239
		_go_fuzz_dep_.CoverTab[530257]++
//line /snap/go/10455/src/net/unixsock_posix.go:239
		_go_fuzz_dep_.CoverTab[9114]++
//line /snap/go/10455/src/net/unixsock_posix.go:239
		// _ = "end of CoverTab[9114]"
//line /snap/go/10455/src/net/unixsock_posix.go:239
	}
//line /snap/go/10455/src/net/unixsock_posix.go:239
	// _ = "end of CoverTab[9109]"
//line /snap/go/10455/src/net/unixsock_posix.go:239
	_go_fuzz_dep_.CoverTab[9110]++
							fd, err := unixSocket(ctx, sl.network, laddr, nil, "listen", ctrlCtxFn)
							if err != nil {
//line /snap/go/10455/src/net/unixsock_posix.go:241
		_go_fuzz_dep_.CoverTab[530258]++
//line /snap/go/10455/src/net/unixsock_posix.go:241
		_go_fuzz_dep_.CoverTab[9115]++
								return nil, err
//line /snap/go/10455/src/net/unixsock_posix.go:242
		// _ = "end of CoverTab[9115]"
	} else {
//line /snap/go/10455/src/net/unixsock_posix.go:243
		_go_fuzz_dep_.CoverTab[530259]++
//line /snap/go/10455/src/net/unixsock_posix.go:243
		_go_fuzz_dep_.CoverTab[9116]++
//line /snap/go/10455/src/net/unixsock_posix.go:243
		// _ = "end of CoverTab[9116]"
//line /snap/go/10455/src/net/unixsock_posix.go:243
	}
//line /snap/go/10455/src/net/unixsock_posix.go:243
	// _ = "end of CoverTab[9110]"
//line /snap/go/10455/src/net/unixsock_posix.go:243
	_go_fuzz_dep_.CoverTab[9111]++
							return newUnixConn(fd), nil
//line /snap/go/10455/src/net/unixsock_posix.go:244
	// _ = "end of CoverTab[9111]"
}

//line /snap/go/10455/src/net/unixsock_posix.go:245
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/unixsock_posix.go:245
var _ = _go_fuzz_dep_.CoverTab
