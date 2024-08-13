// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || windows

//line /usr/local/go/src/net/unixsock_posix.go:7
package net

//line /usr/local/go/src/net/unixsock_posix.go:7
import (
//line /usr/local/go/src/net/unixsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/unixsock_posix.go:7
)
//line /usr/local/go/src/net/unixsock_posix.go:7
import (
//line /usr/local/go/src/net/unixsock_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/unixsock_posix.go:7
)

import (
	"context"
	"errors"
	"os"
	"syscall"
)

func unixSocket(ctx context.Context, net string, laddr, raddr sockaddr, mode string, ctxCtrlFn func(context.Context, string, string, syscall.RawConn) error) (*netFD, error) {
//line /usr/local/go/src/net/unixsock_posix.go:16
	_go_fuzz_dep_.CoverTab[17083]++
							var sotype int
							switch net {
	case "unix":
//line /usr/local/go/src/net/unixsock_posix.go:19
		_go_fuzz_dep_.CoverTab[17087]++
								sotype = syscall.SOCK_STREAM
//line /usr/local/go/src/net/unixsock_posix.go:20
		// _ = "end of CoverTab[17087]"
	case "unixgram":
//line /usr/local/go/src/net/unixsock_posix.go:21
		_go_fuzz_dep_.CoverTab[17088]++
								sotype = syscall.SOCK_DGRAM
//line /usr/local/go/src/net/unixsock_posix.go:22
		// _ = "end of CoverTab[17088]"
	case "unixpacket":
//line /usr/local/go/src/net/unixsock_posix.go:23
		_go_fuzz_dep_.CoverTab[17089]++
								sotype = syscall.SOCK_SEQPACKET
//line /usr/local/go/src/net/unixsock_posix.go:24
		// _ = "end of CoverTab[17089]"
	default:
//line /usr/local/go/src/net/unixsock_posix.go:25
		_go_fuzz_dep_.CoverTab[17090]++
								return nil, UnknownNetworkError(net)
//line /usr/local/go/src/net/unixsock_posix.go:26
		// _ = "end of CoverTab[17090]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:27
	// _ = "end of CoverTab[17083]"
//line /usr/local/go/src/net/unixsock_posix.go:27
	_go_fuzz_dep_.CoverTab[17084]++

							switch mode {
	case "dial":
//line /usr/local/go/src/net/unixsock_posix.go:30
		_go_fuzz_dep_.CoverTab[17091]++
								if laddr != nil && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:31
			_go_fuzz_dep_.CoverTab[17096]++
//line /usr/local/go/src/net/unixsock_posix.go:31
			return laddr.isWildcard()
//line /usr/local/go/src/net/unixsock_posix.go:31
			// _ = "end of CoverTab[17096]"
//line /usr/local/go/src/net/unixsock_posix.go:31
		}() {
//line /usr/local/go/src/net/unixsock_posix.go:31
			_go_fuzz_dep_.CoverTab[17097]++
									laddr = nil
//line /usr/local/go/src/net/unixsock_posix.go:32
			// _ = "end of CoverTab[17097]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:33
			_go_fuzz_dep_.CoverTab[17098]++
//line /usr/local/go/src/net/unixsock_posix.go:33
			// _ = "end of CoverTab[17098]"
//line /usr/local/go/src/net/unixsock_posix.go:33
		}
//line /usr/local/go/src/net/unixsock_posix.go:33
		// _ = "end of CoverTab[17091]"
//line /usr/local/go/src/net/unixsock_posix.go:33
		_go_fuzz_dep_.CoverTab[17092]++
								if raddr != nil && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:34
			_go_fuzz_dep_.CoverTab[17099]++
//line /usr/local/go/src/net/unixsock_posix.go:34
			return raddr.isWildcard()
//line /usr/local/go/src/net/unixsock_posix.go:34
			// _ = "end of CoverTab[17099]"
//line /usr/local/go/src/net/unixsock_posix.go:34
		}() {
//line /usr/local/go/src/net/unixsock_posix.go:34
			_go_fuzz_dep_.CoverTab[17100]++
									raddr = nil
//line /usr/local/go/src/net/unixsock_posix.go:35
			// _ = "end of CoverTab[17100]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:36
			_go_fuzz_dep_.CoverTab[17101]++
//line /usr/local/go/src/net/unixsock_posix.go:36
			// _ = "end of CoverTab[17101]"
//line /usr/local/go/src/net/unixsock_posix.go:36
		}
//line /usr/local/go/src/net/unixsock_posix.go:36
		// _ = "end of CoverTab[17092]"
//line /usr/local/go/src/net/unixsock_posix.go:36
		_go_fuzz_dep_.CoverTab[17093]++
								if raddr == nil && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:37
			_go_fuzz_dep_.CoverTab[17102]++
//line /usr/local/go/src/net/unixsock_posix.go:37
			return (sotype != syscall.SOCK_DGRAM || func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:37
				_go_fuzz_dep_.CoverTab[17103]++
//line /usr/local/go/src/net/unixsock_posix.go:37
				return laddr == nil
//line /usr/local/go/src/net/unixsock_posix.go:37
				// _ = "end of CoverTab[17103]"
//line /usr/local/go/src/net/unixsock_posix.go:37
			}())
//line /usr/local/go/src/net/unixsock_posix.go:37
			// _ = "end of CoverTab[17102]"
//line /usr/local/go/src/net/unixsock_posix.go:37
		}() {
//line /usr/local/go/src/net/unixsock_posix.go:37
			_go_fuzz_dep_.CoverTab[17104]++
									return nil, errMissingAddress
//line /usr/local/go/src/net/unixsock_posix.go:38
			// _ = "end of CoverTab[17104]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:39
			_go_fuzz_dep_.CoverTab[17105]++
//line /usr/local/go/src/net/unixsock_posix.go:39
			// _ = "end of CoverTab[17105]"
//line /usr/local/go/src/net/unixsock_posix.go:39
		}
//line /usr/local/go/src/net/unixsock_posix.go:39
		// _ = "end of CoverTab[17093]"
	case "listen":
//line /usr/local/go/src/net/unixsock_posix.go:40
		_go_fuzz_dep_.CoverTab[17094]++
//line /usr/local/go/src/net/unixsock_posix.go:40
		// _ = "end of CoverTab[17094]"
	default:
//line /usr/local/go/src/net/unixsock_posix.go:41
		_go_fuzz_dep_.CoverTab[17095]++
								return nil, errors.New("unknown mode: " + mode)
//line /usr/local/go/src/net/unixsock_posix.go:42
		// _ = "end of CoverTab[17095]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:43
	// _ = "end of CoverTab[17084]"
//line /usr/local/go/src/net/unixsock_posix.go:43
	_go_fuzz_dep_.CoverTab[17085]++

							fd, err := socket(ctx, net, syscall.AF_UNIX, sotype, 0, false, laddr, raddr, ctxCtrlFn)
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:46
		_go_fuzz_dep_.CoverTab[17106]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:47
		// _ = "end of CoverTab[17106]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:48
		_go_fuzz_dep_.CoverTab[17107]++
//line /usr/local/go/src/net/unixsock_posix.go:48
		// _ = "end of CoverTab[17107]"
//line /usr/local/go/src/net/unixsock_posix.go:48
	}
//line /usr/local/go/src/net/unixsock_posix.go:48
	// _ = "end of CoverTab[17085]"
//line /usr/local/go/src/net/unixsock_posix.go:48
	_go_fuzz_dep_.CoverTab[17086]++
							return fd, nil
//line /usr/local/go/src/net/unixsock_posix.go:49
	// _ = "end of CoverTab[17086]"
}

func sockaddrToUnix(sa syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/unixsock_posix.go:52
	_go_fuzz_dep_.CoverTab[17108]++
							if s, ok := sa.(*syscall.SockaddrUnix); ok {
//line /usr/local/go/src/net/unixsock_posix.go:53
		_go_fuzz_dep_.CoverTab[17110]++
								return &UnixAddr{Name: s.Name, Net: "unix"}
//line /usr/local/go/src/net/unixsock_posix.go:54
		// _ = "end of CoverTab[17110]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:55
		_go_fuzz_dep_.CoverTab[17111]++
//line /usr/local/go/src/net/unixsock_posix.go:55
		// _ = "end of CoverTab[17111]"
//line /usr/local/go/src/net/unixsock_posix.go:55
	}
//line /usr/local/go/src/net/unixsock_posix.go:55
	// _ = "end of CoverTab[17108]"
//line /usr/local/go/src/net/unixsock_posix.go:55
	_go_fuzz_dep_.CoverTab[17109]++
							return nil
//line /usr/local/go/src/net/unixsock_posix.go:56
	// _ = "end of CoverTab[17109]"
}

func sockaddrToUnixgram(sa syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/unixsock_posix.go:59
	_go_fuzz_dep_.CoverTab[17112]++
							if s, ok := sa.(*syscall.SockaddrUnix); ok {
//line /usr/local/go/src/net/unixsock_posix.go:60
		_go_fuzz_dep_.CoverTab[17114]++
								return &UnixAddr{Name: s.Name, Net: "unixgram"}
//line /usr/local/go/src/net/unixsock_posix.go:61
		// _ = "end of CoverTab[17114]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:62
		_go_fuzz_dep_.CoverTab[17115]++
//line /usr/local/go/src/net/unixsock_posix.go:62
		// _ = "end of CoverTab[17115]"
//line /usr/local/go/src/net/unixsock_posix.go:62
	}
//line /usr/local/go/src/net/unixsock_posix.go:62
	// _ = "end of CoverTab[17112]"
//line /usr/local/go/src/net/unixsock_posix.go:62
	_go_fuzz_dep_.CoverTab[17113]++
							return nil
//line /usr/local/go/src/net/unixsock_posix.go:63
	// _ = "end of CoverTab[17113]"
}

func sockaddrToUnixpacket(sa syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/unixsock_posix.go:66
	_go_fuzz_dep_.CoverTab[17116]++
							if s, ok := sa.(*syscall.SockaddrUnix); ok {
//line /usr/local/go/src/net/unixsock_posix.go:67
		_go_fuzz_dep_.CoverTab[17118]++
								return &UnixAddr{Name: s.Name, Net: "unixpacket"}
//line /usr/local/go/src/net/unixsock_posix.go:68
		// _ = "end of CoverTab[17118]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:69
		_go_fuzz_dep_.CoverTab[17119]++
//line /usr/local/go/src/net/unixsock_posix.go:69
		// _ = "end of CoverTab[17119]"
//line /usr/local/go/src/net/unixsock_posix.go:69
	}
//line /usr/local/go/src/net/unixsock_posix.go:69
	// _ = "end of CoverTab[17116]"
//line /usr/local/go/src/net/unixsock_posix.go:69
	_go_fuzz_dep_.CoverTab[17117]++
							return nil
//line /usr/local/go/src/net/unixsock_posix.go:70
	// _ = "end of CoverTab[17117]"
}

func sotypeToNet(sotype int) string {
//line /usr/local/go/src/net/unixsock_posix.go:73
	_go_fuzz_dep_.CoverTab[17120]++
							switch sotype {
	case syscall.SOCK_STREAM:
//line /usr/local/go/src/net/unixsock_posix.go:75
		_go_fuzz_dep_.CoverTab[17121]++
								return "unix"
//line /usr/local/go/src/net/unixsock_posix.go:76
		// _ = "end of CoverTab[17121]"
	case syscall.SOCK_DGRAM:
//line /usr/local/go/src/net/unixsock_posix.go:77
		_go_fuzz_dep_.CoverTab[17122]++
								return "unixgram"
//line /usr/local/go/src/net/unixsock_posix.go:78
		// _ = "end of CoverTab[17122]"
	case syscall.SOCK_SEQPACKET:
//line /usr/local/go/src/net/unixsock_posix.go:79
		_go_fuzz_dep_.CoverTab[17123]++
								return "unixpacket"
//line /usr/local/go/src/net/unixsock_posix.go:80
		// _ = "end of CoverTab[17123]"
	default:
//line /usr/local/go/src/net/unixsock_posix.go:81
		_go_fuzz_dep_.CoverTab[17124]++
								panic("sotypeToNet unknown socket type")
//line /usr/local/go/src/net/unixsock_posix.go:82
		// _ = "end of CoverTab[17124]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:83
	// _ = "end of CoverTab[17120]"
}

func (a *UnixAddr) family() int {
//line /usr/local/go/src/net/unixsock_posix.go:86
	_go_fuzz_dep_.CoverTab[17125]++
							return syscall.AF_UNIX
//line /usr/local/go/src/net/unixsock_posix.go:87
	// _ = "end of CoverTab[17125]"
}

func (a *UnixAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/unixsock_posix.go:90
	_go_fuzz_dep_.CoverTab[17126]++
							if a == nil {
//line /usr/local/go/src/net/unixsock_posix.go:91
		_go_fuzz_dep_.CoverTab[17128]++
								return nil, nil
//line /usr/local/go/src/net/unixsock_posix.go:92
		// _ = "end of CoverTab[17128]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:93
		_go_fuzz_dep_.CoverTab[17129]++
//line /usr/local/go/src/net/unixsock_posix.go:93
		// _ = "end of CoverTab[17129]"
//line /usr/local/go/src/net/unixsock_posix.go:93
	}
//line /usr/local/go/src/net/unixsock_posix.go:93
	// _ = "end of CoverTab[17126]"
//line /usr/local/go/src/net/unixsock_posix.go:93
	_go_fuzz_dep_.CoverTab[17127]++
							return &syscall.SockaddrUnix{Name: a.Name}, nil
//line /usr/local/go/src/net/unixsock_posix.go:94
	// _ = "end of CoverTab[17127]"
}

func (a *UnixAddr) toLocal(net string) sockaddr {
//line /usr/local/go/src/net/unixsock_posix.go:97
	_go_fuzz_dep_.CoverTab[17130]++
							return a
//line /usr/local/go/src/net/unixsock_posix.go:98
	// _ = "end of CoverTab[17130]"
}

func (c *UnixConn) readFrom(b []byte) (int, *UnixAddr, error) {
//line /usr/local/go/src/net/unixsock_posix.go:101
	_go_fuzz_dep_.CoverTab[17131]++
							var addr *UnixAddr
							n, sa, err := c.fd.readFrom(b)
							switch sa := sa.(type) {
	case *syscall.SockaddrUnix:
//line /usr/local/go/src/net/unixsock_posix.go:105
		_go_fuzz_dep_.CoverTab[17133]++
								if sa.Name != "" {
//line /usr/local/go/src/net/unixsock_posix.go:106
			_go_fuzz_dep_.CoverTab[17134]++
									addr = &UnixAddr{Name: sa.Name, Net: sotypeToNet(c.fd.sotype)}
//line /usr/local/go/src/net/unixsock_posix.go:107
			// _ = "end of CoverTab[17134]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:108
			_go_fuzz_dep_.CoverTab[17135]++
//line /usr/local/go/src/net/unixsock_posix.go:108
			// _ = "end of CoverTab[17135]"
//line /usr/local/go/src/net/unixsock_posix.go:108
		}
//line /usr/local/go/src/net/unixsock_posix.go:108
		// _ = "end of CoverTab[17133]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:109
	// _ = "end of CoverTab[17131]"
//line /usr/local/go/src/net/unixsock_posix.go:109
	_go_fuzz_dep_.CoverTab[17132]++
							return n, addr, err
//line /usr/local/go/src/net/unixsock_posix.go:110
	// _ = "end of CoverTab[17132]"
}

func (c *UnixConn) readMsg(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error) {
//line /usr/local/go/src/net/unixsock_posix.go:113
	_go_fuzz_dep_.CoverTab[17136]++
							var sa syscall.Sockaddr
							n, oobn, flags, sa, err = c.fd.readMsg(b, oob, readMsgFlags)
							if readMsgFlags == 0 && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[17139]++
//line /usr/local/go/src/net/unixsock_posix.go:116
		return err == nil
//line /usr/local/go/src/net/unixsock_posix.go:116
		// _ = "end of CoverTab[17139]"
//line /usr/local/go/src/net/unixsock_posix.go:116
	}() && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[17140]++
//line /usr/local/go/src/net/unixsock_posix.go:116
		return oobn > 0
//line /usr/local/go/src/net/unixsock_posix.go:116
		// _ = "end of CoverTab[17140]"
//line /usr/local/go/src/net/unixsock_posix.go:116
	}() {
//line /usr/local/go/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[17141]++
								setReadMsgCloseOnExec(oob[:oobn])
//line /usr/local/go/src/net/unixsock_posix.go:117
		// _ = "end of CoverTab[17141]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:118
		_go_fuzz_dep_.CoverTab[17142]++
//line /usr/local/go/src/net/unixsock_posix.go:118
		// _ = "end of CoverTab[17142]"
//line /usr/local/go/src/net/unixsock_posix.go:118
	}
//line /usr/local/go/src/net/unixsock_posix.go:118
	// _ = "end of CoverTab[17136]"
//line /usr/local/go/src/net/unixsock_posix.go:118
	_go_fuzz_dep_.CoverTab[17137]++

							switch sa := sa.(type) {
	case *syscall.SockaddrUnix:
//line /usr/local/go/src/net/unixsock_posix.go:121
		_go_fuzz_dep_.CoverTab[17143]++
								if sa.Name != "" {
//line /usr/local/go/src/net/unixsock_posix.go:122
			_go_fuzz_dep_.CoverTab[17144]++
									addr = &UnixAddr{Name: sa.Name, Net: sotypeToNet(c.fd.sotype)}
//line /usr/local/go/src/net/unixsock_posix.go:123
			// _ = "end of CoverTab[17144]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:124
			_go_fuzz_dep_.CoverTab[17145]++
//line /usr/local/go/src/net/unixsock_posix.go:124
			// _ = "end of CoverTab[17145]"
//line /usr/local/go/src/net/unixsock_posix.go:124
		}
//line /usr/local/go/src/net/unixsock_posix.go:124
		// _ = "end of CoverTab[17143]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:125
	// _ = "end of CoverTab[17137]"
//line /usr/local/go/src/net/unixsock_posix.go:125
	_go_fuzz_dep_.CoverTab[17138]++
							return
//line /usr/local/go/src/net/unixsock_posix.go:126
	// _ = "end of CoverTab[17138]"
}

func (c *UnixConn) writeTo(b []byte, addr *UnixAddr) (int, error) {
//line /usr/local/go/src/net/unixsock_posix.go:129
	_go_fuzz_dep_.CoverTab[17146]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/unixsock_posix.go:130
		_go_fuzz_dep_.CoverTab[17150]++
								return 0, ErrWriteToConnected
//line /usr/local/go/src/net/unixsock_posix.go:131
		// _ = "end of CoverTab[17150]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:132
		_go_fuzz_dep_.CoverTab[17151]++
//line /usr/local/go/src/net/unixsock_posix.go:132
		// _ = "end of CoverTab[17151]"
//line /usr/local/go/src/net/unixsock_posix.go:132
	}
//line /usr/local/go/src/net/unixsock_posix.go:132
	// _ = "end of CoverTab[17146]"
//line /usr/local/go/src/net/unixsock_posix.go:132
	_go_fuzz_dep_.CoverTab[17147]++
							if addr == nil {
//line /usr/local/go/src/net/unixsock_posix.go:133
		_go_fuzz_dep_.CoverTab[17152]++
								return 0, errMissingAddress
//line /usr/local/go/src/net/unixsock_posix.go:134
		// _ = "end of CoverTab[17152]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:135
		_go_fuzz_dep_.CoverTab[17153]++
//line /usr/local/go/src/net/unixsock_posix.go:135
		// _ = "end of CoverTab[17153]"
//line /usr/local/go/src/net/unixsock_posix.go:135
	}
//line /usr/local/go/src/net/unixsock_posix.go:135
	// _ = "end of CoverTab[17147]"
//line /usr/local/go/src/net/unixsock_posix.go:135
	_go_fuzz_dep_.CoverTab[17148]++
							if addr.Net != sotypeToNet(c.fd.sotype) {
//line /usr/local/go/src/net/unixsock_posix.go:136
		_go_fuzz_dep_.CoverTab[17154]++
								return 0, syscall.EAFNOSUPPORT
//line /usr/local/go/src/net/unixsock_posix.go:137
		// _ = "end of CoverTab[17154]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:138
		_go_fuzz_dep_.CoverTab[17155]++
//line /usr/local/go/src/net/unixsock_posix.go:138
		// _ = "end of CoverTab[17155]"
//line /usr/local/go/src/net/unixsock_posix.go:138
	}
//line /usr/local/go/src/net/unixsock_posix.go:138
	// _ = "end of CoverTab[17148]"
//line /usr/local/go/src/net/unixsock_posix.go:138
	_go_fuzz_dep_.CoverTab[17149]++
							sa := &syscall.SockaddrUnix{Name: addr.Name}
							return c.fd.writeTo(b, sa)
//line /usr/local/go/src/net/unixsock_posix.go:140
	// _ = "end of CoverTab[17149]"
}

func (c *UnixConn) writeMsg(b, oob []byte, addr *UnixAddr) (n, oobn int, err error) {
//line /usr/local/go/src/net/unixsock_posix.go:143
	_go_fuzz_dep_.CoverTab[17156]++
							if c.fd.sotype == syscall.SOCK_DGRAM && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:144
		_go_fuzz_dep_.CoverTab[17159]++
//line /usr/local/go/src/net/unixsock_posix.go:144
		return c.fd.isConnected
//line /usr/local/go/src/net/unixsock_posix.go:144
		// _ = "end of CoverTab[17159]"
//line /usr/local/go/src/net/unixsock_posix.go:144
	}() {
//line /usr/local/go/src/net/unixsock_posix.go:144
		_go_fuzz_dep_.CoverTab[17160]++
								return 0, 0, ErrWriteToConnected
//line /usr/local/go/src/net/unixsock_posix.go:145
		// _ = "end of CoverTab[17160]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:146
		_go_fuzz_dep_.CoverTab[17161]++
//line /usr/local/go/src/net/unixsock_posix.go:146
		// _ = "end of CoverTab[17161]"
//line /usr/local/go/src/net/unixsock_posix.go:146
	}
//line /usr/local/go/src/net/unixsock_posix.go:146
	// _ = "end of CoverTab[17156]"
//line /usr/local/go/src/net/unixsock_posix.go:146
	_go_fuzz_dep_.CoverTab[17157]++
							var sa syscall.Sockaddr
							if addr != nil {
//line /usr/local/go/src/net/unixsock_posix.go:148
		_go_fuzz_dep_.CoverTab[17162]++
								if addr.Net != sotypeToNet(c.fd.sotype) {
//line /usr/local/go/src/net/unixsock_posix.go:149
			_go_fuzz_dep_.CoverTab[17164]++
									return 0, 0, syscall.EAFNOSUPPORT
//line /usr/local/go/src/net/unixsock_posix.go:150
			// _ = "end of CoverTab[17164]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:151
			_go_fuzz_dep_.CoverTab[17165]++
//line /usr/local/go/src/net/unixsock_posix.go:151
			// _ = "end of CoverTab[17165]"
//line /usr/local/go/src/net/unixsock_posix.go:151
		}
//line /usr/local/go/src/net/unixsock_posix.go:151
		// _ = "end of CoverTab[17162]"
//line /usr/local/go/src/net/unixsock_posix.go:151
		_go_fuzz_dep_.CoverTab[17163]++
								sa = &syscall.SockaddrUnix{Name: addr.Name}
//line /usr/local/go/src/net/unixsock_posix.go:152
		// _ = "end of CoverTab[17163]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:153
		_go_fuzz_dep_.CoverTab[17166]++
//line /usr/local/go/src/net/unixsock_posix.go:153
		// _ = "end of CoverTab[17166]"
//line /usr/local/go/src/net/unixsock_posix.go:153
	}
//line /usr/local/go/src/net/unixsock_posix.go:153
	// _ = "end of CoverTab[17157]"
//line /usr/local/go/src/net/unixsock_posix.go:153
	_go_fuzz_dep_.CoverTab[17158]++
							return c.fd.writeMsg(b, oob, sa)
//line /usr/local/go/src/net/unixsock_posix.go:154
	// _ = "end of CoverTab[17158]"
}

func (sd *sysDialer) dialUnix(ctx context.Context, laddr, raddr *UnixAddr) (*UnixConn, error) {
//line /usr/local/go/src/net/unixsock_posix.go:157
	_go_fuzz_dep_.CoverTab[17167]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:159
		_go_fuzz_dep_.CoverTab[17170]++
//line /usr/local/go/src/net/unixsock_posix.go:159
		return sd.Dialer.Control != nil
//line /usr/local/go/src/net/unixsock_posix.go:159
		// _ = "end of CoverTab[17170]"
//line /usr/local/go/src/net/unixsock_posix.go:159
	}() {
//line /usr/local/go/src/net/unixsock_posix.go:159
		_go_fuzz_dep_.CoverTab[17171]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/unixsock_posix.go:160
			_go_fuzz_dep_.CoverTab[17172]++
									return sd.Dialer.Control(network, address, c)
//line /usr/local/go/src/net/unixsock_posix.go:161
			// _ = "end of CoverTab[17172]"
		}
//line /usr/local/go/src/net/unixsock_posix.go:162
		// _ = "end of CoverTab[17171]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:163
		_go_fuzz_dep_.CoverTab[17173]++
//line /usr/local/go/src/net/unixsock_posix.go:163
		// _ = "end of CoverTab[17173]"
//line /usr/local/go/src/net/unixsock_posix.go:163
	}
//line /usr/local/go/src/net/unixsock_posix.go:163
	// _ = "end of CoverTab[17167]"
//line /usr/local/go/src/net/unixsock_posix.go:163
	_go_fuzz_dep_.CoverTab[17168]++
							fd, err := unixSocket(ctx, sd.network, laddr, raddr, "dial", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:165
		_go_fuzz_dep_.CoverTab[17174]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:166
		// _ = "end of CoverTab[17174]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:167
		_go_fuzz_dep_.CoverTab[17175]++
//line /usr/local/go/src/net/unixsock_posix.go:167
		// _ = "end of CoverTab[17175]"
//line /usr/local/go/src/net/unixsock_posix.go:167
	}
//line /usr/local/go/src/net/unixsock_posix.go:167
	// _ = "end of CoverTab[17168]"
//line /usr/local/go/src/net/unixsock_posix.go:167
	_go_fuzz_dep_.CoverTab[17169]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/unixsock_posix.go:168
	// _ = "end of CoverTab[17169]"
}

func (ln *UnixListener) accept() (*UnixConn, error) {
//line /usr/local/go/src/net/unixsock_posix.go:171
	_go_fuzz_dep_.CoverTab[17176]++
							fd, err := ln.fd.accept()
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:173
		_go_fuzz_dep_.CoverTab[17178]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:174
		// _ = "end of CoverTab[17178]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:175
		_go_fuzz_dep_.CoverTab[17179]++
//line /usr/local/go/src/net/unixsock_posix.go:175
		// _ = "end of CoverTab[17179]"
//line /usr/local/go/src/net/unixsock_posix.go:175
	}
//line /usr/local/go/src/net/unixsock_posix.go:175
	// _ = "end of CoverTab[17176]"
//line /usr/local/go/src/net/unixsock_posix.go:175
	_go_fuzz_dep_.CoverTab[17177]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/unixsock_posix.go:176
	// _ = "end of CoverTab[17177]"
}

func (ln *UnixListener) close() error {
//line /usr/local/go/src/net/unixsock_posix.go:179
	_go_fuzz_dep_.CoverTab[17180]++

//line /usr/local/go/src/net/unixsock_posix.go:191
	ln.unlinkOnce.Do(func() {
//line /usr/local/go/src/net/unixsock_posix.go:191
		_go_fuzz_dep_.CoverTab[17182]++
								if ln.path[0] != '@' && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:192
			_go_fuzz_dep_.CoverTab[17183]++
//line /usr/local/go/src/net/unixsock_posix.go:192
			return ln.unlink
//line /usr/local/go/src/net/unixsock_posix.go:192
			// _ = "end of CoverTab[17183]"
//line /usr/local/go/src/net/unixsock_posix.go:192
		}() {
//line /usr/local/go/src/net/unixsock_posix.go:192
			_go_fuzz_dep_.CoverTab[17184]++
									syscall.Unlink(ln.path)
//line /usr/local/go/src/net/unixsock_posix.go:193
			// _ = "end of CoverTab[17184]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:194
			_go_fuzz_dep_.CoverTab[17185]++
//line /usr/local/go/src/net/unixsock_posix.go:194
			// _ = "end of CoverTab[17185]"
//line /usr/local/go/src/net/unixsock_posix.go:194
		}
//line /usr/local/go/src/net/unixsock_posix.go:194
		// _ = "end of CoverTab[17182]"
	})
//line /usr/local/go/src/net/unixsock_posix.go:195
	// _ = "end of CoverTab[17180]"
//line /usr/local/go/src/net/unixsock_posix.go:195
	_go_fuzz_dep_.CoverTab[17181]++
							return ln.fd.Close()
//line /usr/local/go/src/net/unixsock_posix.go:196
	// _ = "end of CoverTab[17181]"
}

func (ln *UnixListener) file() (*os.File, error) {
//line /usr/local/go/src/net/unixsock_posix.go:199
	_go_fuzz_dep_.CoverTab[17186]++
							f, err := ln.fd.dup()
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:201
		_go_fuzz_dep_.CoverTab[17188]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:202
		// _ = "end of CoverTab[17188]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:203
		_go_fuzz_dep_.CoverTab[17189]++
//line /usr/local/go/src/net/unixsock_posix.go:203
		// _ = "end of CoverTab[17189]"
//line /usr/local/go/src/net/unixsock_posix.go:203
	}
//line /usr/local/go/src/net/unixsock_posix.go:203
	// _ = "end of CoverTab[17186]"
//line /usr/local/go/src/net/unixsock_posix.go:203
	_go_fuzz_dep_.CoverTab[17187]++
							return f, nil
//line /usr/local/go/src/net/unixsock_posix.go:204
	// _ = "end of CoverTab[17187]"
}

// SetUnlinkOnClose sets whether the underlying socket file should be removed
//line /usr/local/go/src/net/unixsock_posix.go:207
// from the file system when the listener is closed.
//line /usr/local/go/src/net/unixsock_posix.go:207
//
//line /usr/local/go/src/net/unixsock_posix.go:207
// The default behavior is to unlink the socket file only when package net created it.
//line /usr/local/go/src/net/unixsock_posix.go:207
// That is, when the listener and the underlying socket file were created by a call to
//line /usr/local/go/src/net/unixsock_posix.go:207
// Listen or ListenUnix, then by default closing the listener will remove the socket file.
//line /usr/local/go/src/net/unixsock_posix.go:207
// but if the listener was created by a call to FileListener to use an already existing
//line /usr/local/go/src/net/unixsock_posix.go:207
// socket file, then by default closing the listener will not remove the socket file.
//line /usr/local/go/src/net/unixsock_posix.go:215
func (l *UnixListener) SetUnlinkOnClose(unlink bool) {
//line /usr/local/go/src/net/unixsock_posix.go:215
	_go_fuzz_dep_.CoverTab[17190]++
							l.unlink = unlink
//line /usr/local/go/src/net/unixsock_posix.go:216
	// _ = "end of CoverTab[17190]"
}

func (sl *sysListener) listenUnix(ctx context.Context, laddr *UnixAddr) (*UnixListener, error) {
//line /usr/local/go/src/net/unixsock_posix.go:219
	_go_fuzz_dep_.CoverTab[17191]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/unixsock_posix.go:221
		_go_fuzz_dep_.CoverTab[17194]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/unixsock_posix.go:222
			_go_fuzz_dep_.CoverTab[17195]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/unixsock_posix.go:223
			// _ = "end of CoverTab[17195]"
		}
//line /usr/local/go/src/net/unixsock_posix.go:224
		// _ = "end of CoverTab[17194]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:225
		_go_fuzz_dep_.CoverTab[17196]++
//line /usr/local/go/src/net/unixsock_posix.go:225
		// _ = "end of CoverTab[17196]"
//line /usr/local/go/src/net/unixsock_posix.go:225
	}
//line /usr/local/go/src/net/unixsock_posix.go:225
	// _ = "end of CoverTab[17191]"
//line /usr/local/go/src/net/unixsock_posix.go:225
	_go_fuzz_dep_.CoverTab[17192]++
							fd, err := unixSocket(ctx, sl.network, laddr, nil, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:227
		_go_fuzz_dep_.CoverTab[17197]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:228
		// _ = "end of CoverTab[17197]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:229
		_go_fuzz_dep_.CoverTab[17198]++
//line /usr/local/go/src/net/unixsock_posix.go:229
		// _ = "end of CoverTab[17198]"
//line /usr/local/go/src/net/unixsock_posix.go:229
	}
//line /usr/local/go/src/net/unixsock_posix.go:229
	// _ = "end of CoverTab[17192]"
//line /usr/local/go/src/net/unixsock_posix.go:229
	_go_fuzz_dep_.CoverTab[17193]++
							return &UnixListener{fd: fd, path: fd.laddr.String(), unlink: true}, nil
//line /usr/local/go/src/net/unixsock_posix.go:230
	// _ = "end of CoverTab[17193]"
}

func (sl *sysListener) listenUnixgram(ctx context.Context, laddr *UnixAddr) (*UnixConn, error) {
//line /usr/local/go/src/net/unixsock_posix.go:233
	_go_fuzz_dep_.CoverTab[17199]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/unixsock_posix.go:235
		_go_fuzz_dep_.CoverTab[17202]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/unixsock_posix.go:236
			_go_fuzz_dep_.CoverTab[17203]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/unixsock_posix.go:237
			// _ = "end of CoverTab[17203]"
		}
//line /usr/local/go/src/net/unixsock_posix.go:238
		// _ = "end of CoverTab[17202]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:239
		_go_fuzz_dep_.CoverTab[17204]++
//line /usr/local/go/src/net/unixsock_posix.go:239
		// _ = "end of CoverTab[17204]"
//line /usr/local/go/src/net/unixsock_posix.go:239
	}
//line /usr/local/go/src/net/unixsock_posix.go:239
	// _ = "end of CoverTab[17199]"
//line /usr/local/go/src/net/unixsock_posix.go:239
	_go_fuzz_dep_.CoverTab[17200]++
							fd, err := unixSocket(ctx, sl.network, laddr, nil, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:241
		_go_fuzz_dep_.CoverTab[17205]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:242
		// _ = "end of CoverTab[17205]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:243
		_go_fuzz_dep_.CoverTab[17206]++
//line /usr/local/go/src/net/unixsock_posix.go:243
		// _ = "end of CoverTab[17206]"
//line /usr/local/go/src/net/unixsock_posix.go:243
	}
//line /usr/local/go/src/net/unixsock_posix.go:243
	// _ = "end of CoverTab[17200]"
//line /usr/local/go/src/net/unixsock_posix.go:243
	_go_fuzz_dep_.CoverTab[17201]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/unixsock_posix.go:244
	// _ = "end of CoverTab[17201]"
}

//line /usr/local/go/src/net/unixsock_posix.go:245
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/unixsock_posix.go:245
var _ = _go_fuzz_dep_.CoverTab
