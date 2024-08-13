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
	_go_fuzz_dep_.CoverTab[8693]++
							var sotype int
							switch net {
	case "unix":
//line /usr/local/go/src/net/unixsock_posix.go:19
		_go_fuzz_dep_.CoverTab[8697]++
								sotype = syscall.SOCK_STREAM
//line /usr/local/go/src/net/unixsock_posix.go:20
		// _ = "end of CoverTab[8697]"
	case "unixgram":
//line /usr/local/go/src/net/unixsock_posix.go:21
		_go_fuzz_dep_.CoverTab[8698]++
								sotype = syscall.SOCK_DGRAM
//line /usr/local/go/src/net/unixsock_posix.go:22
		// _ = "end of CoverTab[8698]"
	case "unixpacket":
//line /usr/local/go/src/net/unixsock_posix.go:23
		_go_fuzz_dep_.CoverTab[8699]++
								sotype = syscall.SOCK_SEQPACKET
//line /usr/local/go/src/net/unixsock_posix.go:24
		// _ = "end of CoverTab[8699]"
	default:
//line /usr/local/go/src/net/unixsock_posix.go:25
		_go_fuzz_dep_.CoverTab[8700]++
								return nil, UnknownNetworkError(net)
//line /usr/local/go/src/net/unixsock_posix.go:26
		// _ = "end of CoverTab[8700]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:27
	// _ = "end of CoverTab[8693]"
//line /usr/local/go/src/net/unixsock_posix.go:27
	_go_fuzz_dep_.CoverTab[8694]++

							switch mode {
	case "dial":
//line /usr/local/go/src/net/unixsock_posix.go:30
		_go_fuzz_dep_.CoverTab[8701]++
								if laddr != nil && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:31
			_go_fuzz_dep_.CoverTab[8706]++
//line /usr/local/go/src/net/unixsock_posix.go:31
			return laddr.isWildcard()
//line /usr/local/go/src/net/unixsock_posix.go:31
			// _ = "end of CoverTab[8706]"
//line /usr/local/go/src/net/unixsock_posix.go:31
		}() {
//line /usr/local/go/src/net/unixsock_posix.go:31
			_go_fuzz_dep_.CoverTab[8707]++
									laddr = nil
//line /usr/local/go/src/net/unixsock_posix.go:32
			// _ = "end of CoverTab[8707]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:33
			_go_fuzz_dep_.CoverTab[8708]++
//line /usr/local/go/src/net/unixsock_posix.go:33
			// _ = "end of CoverTab[8708]"
//line /usr/local/go/src/net/unixsock_posix.go:33
		}
//line /usr/local/go/src/net/unixsock_posix.go:33
		// _ = "end of CoverTab[8701]"
//line /usr/local/go/src/net/unixsock_posix.go:33
		_go_fuzz_dep_.CoverTab[8702]++
								if raddr != nil && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:34
			_go_fuzz_dep_.CoverTab[8709]++
//line /usr/local/go/src/net/unixsock_posix.go:34
			return raddr.isWildcard()
//line /usr/local/go/src/net/unixsock_posix.go:34
			// _ = "end of CoverTab[8709]"
//line /usr/local/go/src/net/unixsock_posix.go:34
		}() {
//line /usr/local/go/src/net/unixsock_posix.go:34
			_go_fuzz_dep_.CoverTab[8710]++
									raddr = nil
//line /usr/local/go/src/net/unixsock_posix.go:35
			// _ = "end of CoverTab[8710]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:36
			_go_fuzz_dep_.CoverTab[8711]++
//line /usr/local/go/src/net/unixsock_posix.go:36
			// _ = "end of CoverTab[8711]"
//line /usr/local/go/src/net/unixsock_posix.go:36
		}
//line /usr/local/go/src/net/unixsock_posix.go:36
		// _ = "end of CoverTab[8702]"
//line /usr/local/go/src/net/unixsock_posix.go:36
		_go_fuzz_dep_.CoverTab[8703]++
								if raddr == nil && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:37
			_go_fuzz_dep_.CoverTab[8712]++
//line /usr/local/go/src/net/unixsock_posix.go:37
			return (sotype != syscall.SOCK_DGRAM || func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:37
				_go_fuzz_dep_.CoverTab[8713]++
//line /usr/local/go/src/net/unixsock_posix.go:37
				return laddr == nil
//line /usr/local/go/src/net/unixsock_posix.go:37
				// _ = "end of CoverTab[8713]"
//line /usr/local/go/src/net/unixsock_posix.go:37
			}())
//line /usr/local/go/src/net/unixsock_posix.go:37
			// _ = "end of CoverTab[8712]"
//line /usr/local/go/src/net/unixsock_posix.go:37
		}() {
//line /usr/local/go/src/net/unixsock_posix.go:37
			_go_fuzz_dep_.CoverTab[8714]++
									return nil, errMissingAddress
//line /usr/local/go/src/net/unixsock_posix.go:38
			// _ = "end of CoverTab[8714]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:39
			_go_fuzz_dep_.CoverTab[8715]++
//line /usr/local/go/src/net/unixsock_posix.go:39
			// _ = "end of CoverTab[8715]"
//line /usr/local/go/src/net/unixsock_posix.go:39
		}
//line /usr/local/go/src/net/unixsock_posix.go:39
		// _ = "end of CoverTab[8703]"
	case "listen":
//line /usr/local/go/src/net/unixsock_posix.go:40
		_go_fuzz_dep_.CoverTab[8704]++
//line /usr/local/go/src/net/unixsock_posix.go:40
		// _ = "end of CoverTab[8704]"
	default:
//line /usr/local/go/src/net/unixsock_posix.go:41
		_go_fuzz_dep_.CoverTab[8705]++
								return nil, errors.New("unknown mode: " + mode)
//line /usr/local/go/src/net/unixsock_posix.go:42
		// _ = "end of CoverTab[8705]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:43
	// _ = "end of CoverTab[8694]"
//line /usr/local/go/src/net/unixsock_posix.go:43
	_go_fuzz_dep_.CoverTab[8695]++

							fd, err := socket(ctx, net, syscall.AF_UNIX, sotype, 0, false, laddr, raddr, ctxCtrlFn)
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:46
		_go_fuzz_dep_.CoverTab[8716]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:47
		// _ = "end of CoverTab[8716]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:48
		_go_fuzz_dep_.CoverTab[8717]++
//line /usr/local/go/src/net/unixsock_posix.go:48
		// _ = "end of CoverTab[8717]"
//line /usr/local/go/src/net/unixsock_posix.go:48
	}
//line /usr/local/go/src/net/unixsock_posix.go:48
	// _ = "end of CoverTab[8695]"
//line /usr/local/go/src/net/unixsock_posix.go:48
	_go_fuzz_dep_.CoverTab[8696]++
							return fd, nil
//line /usr/local/go/src/net/unixsock_posix.go:49
	// _ = "end of CoverTab[8696]"
}

func sockaddrToUnix(sa syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/unixsock_posix.go:52
	_go_fuzz_dep_.CoverTab[8718]++
							if s, ok := sa.(*syscall.SockaddrUnix); ok {
//line /usr/local/go/src/net/unixsock_posix.go:53
		_go_fuzz_dep_.CoverTab[8720]++
								return &UnixAddr{Name: s.Name, Net: "unix"}
//line /usr/local/go/src/net/unixsock_posix.go:54
		// _ = "end of CoverTab[8720]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:55
		_go_fuzz_dep_.CoverTab[8721]++
//line /usr/local/go/src/net/unixsock_posix.go:55
		// _ = "end of CoverTab[8721]"
//line /usr/local/go/src/net/unixsock_posix.go:55
	}
//line /usr/local/go/src/net/unixsock_posix.go:55
	// _ = "end of CoverTab[8718]"
//line /usr/local/go/src/net/unixsock_posix.go:55
	_go_fuzz_dep_.CoverTab[8719]++
							return nil
//line /usr/local/go/src/net/unixsock_posix.go:56
	// _ = "end of CoverTab[8719]"
}

func sockaddrToUnixgram(sa syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/unixsock_posix.go:59
	_go_fuzz_dep_.CoverTab[8722]++
							if s, ok := sa.(*syscall.SockaddrUnix); ok {
//line /usr/local/go/src/net/unixsock_posix.go:60
		_go_fuzz_dep_.CoverTab[8724]++
								return &UnixAddr{Name: s.Name, Net: "unixgram"}
//line /usr/local/go/src/net/unixsock_posix.go:61
		// _ = "end of CoverTab[8724]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:62
		_go_fuzz_dep_.CoverTab[8725]++
//line /usr/local/go/src/net/unixsock_posix.go:62
		// _ = "end of CoverTab[8725]"
//line /usr/local/go/src/net/unixsock_posix.go:62
	}
//line /usr/local/go/src/net/unixsock_posix.go:62
	// _ = "end of CoverTab[8722]"
//line /usr/local/go/src/net/unixsock_posix.go:62
	_go_fuzz_dep_.CoverTab[8723]++
							return nil
//line /usr/local/go/src/net/unixsock_posix.go:63
	// _ = "end of CoverTab[8723]"
}

func sockaddrToUnixpacket(sa syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/unixsock_posix.go:66
	_go_fuzz_dep_.CoverTab[8726]++
							if s, ok := sa.(*syscall.SockaddrUnix); ok {
//line /usr/local/go/src/net/unixsock_posix.go:67
		_go_fuzz_dep_.CoverTab[8728]++
								return &UnixAddr{Name: s.Name, Net: "unixpacket"}
//line /usr/local/go/src/net/unixsock_posix.go:68
		// _ = "end of CoverTab[8728]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:69
		_go_fuzz_dep_.CoverTab[8729]++
//line /usr/local/go/src/net/unixsock_posix.go:69
		// _ = "end of CoverTab[8729]"
//line /usr/local/go/src/net/unixsock_posix.go:69
	}
//line /usr/local/go/src/net/unixsock_posix.go:69
	// _ = "end of CoverTab[8726]"
//line /usr/local/go/src/net/unixsock_posix.go:69
	_go_fuzz_dep_.CoverTab[8727]++
							return nil
//line /usr/local/go/src/net/unixsock_posix.go:70
	// _ = "end of CoverTab[8727]"
}

func sotypeToNet(sotype int) string {
//line /usr/local/go/src/net/unixsock_posix.go:73
	_go_fuzz_dep_.CoverTab[8730]++
							switch sotype {
	case syscall.SOCK_STREAM:
//line /usr/local/go/src/net/unixsock_posix.go:75
		_go_fuzz_dep_.CoverTab[8731]++
								return "unix"
//line /usr/local/go/src/net/unixsock_posix.go:76
		// _ = "end of CoverTab[8731]"
	case syscall.SOCK_DGRAM:
//line /usr/local/go/src/net/unixsock_posix.go:77
		_go_fuzz_dep_.CoverTab[8732]++
								return "unixgram"
//line /usr/local/go/src/net/unixsock_posix.go:78
		// _ = "end of CoverTab[8732]"
	case syscall.SOCK_SEQPACKET:
//line /usr/local/go/src/net/unixsock_posix.go:79
		_go_fuzz_dep_.CoverTab[8733]++
								return "unixpacket"
//line /usr/local/go/src/net/unixsock_posix.go:80
		// _ = "end of CoverTab[8733]"
	default:
//line /usr/local/go/src/net/unixsock_posix.go:81
		_go_fuzz_dep_.CoverTab[8734]++
								panic("sotypeToNet unknown socket type")
//line /usr/local/go/src/net/unixsock_posix.go:82
		// _ = "end of CoverTab[8734]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:83
	// _ = "end of CoverTab[8730]"
}

func (a *UnixAddr) family() int {
//line /usr/local/go/src/net/unixsock_posix.go:86
	_go_fuzz_dep_.CoverTab[8735]++
							return syscall.AF_UNIX
//line /usr/local/go/src/net/unixsock_posix.go:87
	// _ = "end of CoverTab[8735]"
}

func (a *UnixAddr) sockaddr(family int) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/unixsock_posix.go:90
	_go_fuzz_dep_.CoverTab[8736]++
							if a == nil {
//line /usr/local/go/src/net/unixsock_posix.go:91
		_go_fuzz_dep_.CoverTab[8738]++
								return nil, nil
//line /usr/local/go/src/net/unixsock_posix.go:92
		// _ = "end of CoverTab[8738]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:93
		_go_fuzz_dep_.CoverTab[8739]++
//line /usr/local/go/src/net/unixsock_posix.go:93
		// _ = "end of CoverTab[8739]"
//line /usr/local/go/src/net/unixsock_posix.go:93
	}
//line /usr/local/go/src/net/unixsock_posix.go:93
	// _ = "end of CoverTab[8736]"
//line /usr/local/go/src/net/unixsock_posix.go:93
	_go_fuzz_dep_.CoverTab[8737]++
							return &syscall.SockaddrUnix{Name: a.Name}, nil
//line /usr/local/go/src/net/unixsock_posix.go:94
	// _ = "end of CoverTab[8737]"
}

func (a *UnixAddr) toLocal(net string) sockaddr {
//line /usr/local/go/src/net/unixsock_posix.go:97
	_go_fuzz_dep_.CoverTab[8740]++
							return a
//line /usr/local/go/src/net/unixsock_posix.go:98
	// _ = "end of CoverTab[8740]"
}

func (c *UnixConn) readFrom(b []byte) (int, *UnixAddr, error) {
//line /usr/local/go/src/net/unixsock_posix.go:101
	_go_fuzz_dep_.CoverTab[8741]++
							var addr *UnixAddr
							n, sa, err := c.fd.readFrom(b)
							switch sa := sa.(type) {
	case *syscall.SockaddrUnix:
//line /usr/local/go/src/net/unixsock_posix.go:105
		_go_fuzz_dep_.CoverTab[8743]++
								if sa.Name != "" {
//line /usr/local/go/src/net/unixsock_posix.go:106
			_go_fuzz_dep_.CoverTab[8744]++
									addr = &UnixAddr{Name: sa.Name, Net: sotypeToNet(c.fd.sotype)}
//line /usr/local/go/src/net/unixsock_posix.go:107
			// _ = "end of CoverTab[8744]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:108
			_go_fuzz_dep_.CoverTab[8745]++
//line /usr/local/go/src/net/unixsock_posix.go:108
			// _ = "end of CoverTab[8745]"
//line /usr/local/go/src/net/unixsock_posix.go:108
		}
//line /usr/local/go/src/net/unixsock_posix.go:108
		// _ = "end of CoverTab[8743]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:109
	// _ = "end of CoverTab[8741]"
//line /usr/local/go/src/net/unixsock_posix.go:109
	_go_fuzz_dep_.CoverTab[8742]++
							return n, addr, err
//line /usr/local/go/src/net/unixsock_posix.go:110
	// _ = "end of CoverTab[8742]"
}

func (c *UnixConn) readMsg(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error) {
//line /usr/local/go/src/net/unixsock_posix.go:113
	_go_fuzz_dep_.CoverTab[8746]++
							var sa syscall.Sockaddr
							n, oobn, flags, sa, err = c.fd.readMsg(b, oob, readMsgFlags)
							if readMsgFlags == 0 && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[8749]++
//line /usr/local/go/src/net/unixsock_posix.go:116
		return err == nil
//line /usr/local/go/src/net/unixsock_posix.go:116
		// _ = "end of CoverTab[8749]"
//line /usr/local/go/src/net/unixsock_posix.go:116
	}() && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[8750]++
//line /usr/local/go/src/net/unixsock_posix.go:116
		return oobn > 0
//line /usr/local/go/src/net/unixsock_posix.go:116
		// _ = "end of CoverTab[8750]"
//line /usr/local/go/src/net/unixsock_posix.go:116
	}() {
//line /usr/local/go/src/net/unixsock_posix.go:116
		_go_fuzz_dep_.CoverTab[8751]++
								setReadMsgCloseOnExec(oob[:oobn])
//line /usr/local/go/src/net/unixsock_posix.go:117
		// _ = "end of CoverTab[8751]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:118
		_go_fuzz_dep_.CoverTab[8752]++
//line /usr/local/go/src/net/unixsock_posix.go:118
		// _ = "end of CoverTab[8752]"
//line /usr/local/go/src/net/unixsock_posix.go:118
	}
//line /usr/local/go/src/net/unixsock_posix.go:118
	// _ = "end of CoverTab[8746]"
//line /usr/local/go/src/net/unixsock_posix.go:118
	_go_fuzz_dep_.CoverTab[8747]++

							switch sa := sa.(type) {
	case *syscall.SockaddrUnix:
//line /usr/local/go/src/net/unixsock_posix.go:121
		_go_fuzz_dep_.CoverTab[8753]++
								if sa.Name != "" {
//line /usr/local/go/src/net/unixsock_posix.go:122
			_go_fuzz_dep_.CoverTab[8754]++
									addr = &UnixAddr{Name: sa.Name, Net: sotypeToNet(c.fd.sotype)}
//line /usr/local/go/src/net/unixsock_posix.go:123
			// _ = "end of CoverTab[8754]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:124
			_go_fuzz_dep_.CoverTab[8755]++
//line /usr/local/go/src/net/unixsock_posix.go:124
			// _ = "end of CoverTab[8755]"
//line /usr/local/go/src/net/unixsock_posix.go:124
		}
//line /usr/local/go/src/net/unixsock_posix.go:124
		// _ = "end of CoverTab[8753]"
	}
//line /usr/local/go/src/net/unixsock_posix.go:125
	// _ = "end of CoverTab[8747]"
//line /usr/local/go/src/net/unixsock_posix.go:125
	_go_fuzz_dep_.CoverTab[8748]++
							return
//line /usr/local/go/src/net/unixsock_posix.go:126
	// _ = "end of CoverTab[8748]"
}

func (c *UnixConn) writeTo(b []byte, addr *UnixAddr) (int, error) {
//line /usr/local/go/src/net/unixsock_posix.go:129
	_go_fuzz_dep_.CoverTab[8756]++
							if c.fd.isConnected {
//line /usr/local/go/src/net/unixsock_posix.go:130
		_go_fuzz_dep_.CoverTab[8760]++
								return 0, ErrWriteToConnected
//line /usr/local/go/src/net/unixsock_posix.go:131
		// _ = "end of CoverTab[8760]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:132
		_go_fuzz_dep_.CoverTab[8761]++
//line /usr/local/go/src/net/unixsock_posix.go:132
		// _ = "end of CoverTab[8761]"
//line /usr/local/go/src/net/unixsock_posix.go:132
	}
//line /usr/local/go/src/net/unixsock_posix.go:132
	// _ = "end of CoverTab[8756]"
//line /usr/local/go/src/net/unixsock_posix.go:132
	_go_fuzz_dep_.CoverTab[8757]++
							if addr == nil {
//line /usr/local/go/src/net/unixsock_posix.go:133
		_go_fuzz_dep_.CoverTab[8762]++
								return 0, errMissingAddress
//line /usr/local/go/src/net/unixsock_posix.go:134
		// _ = "end of CoverTab[8762]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:135
		_go_fuzz_dep_.CoverTab[8763]++
//line /usr/local/go/src/net/unixsock_posix.go:135
		// _ = "end of CoverTab[8763]"
//line /usr/local/go/src/net/unixsock_posix.go:135
	}
//line /usr/local/go/src/net/unixsock_posix.go:135
	// _ = "end of CoverTab[8757]"
//line /usr/local/go/src/net/unixsock_posix.go:135
	_go_fuzz_dep_.CoverTab[8758]++
							if addr.Net != sotypeToNet(c.fd.sotype) {
//line /usr/local/go/src/net/unixsock_posix.go:136
		_go_fuzz_dep_.CoverTab[8764]++
								return 0, syscall.EAFNOSUPPORT
//line /usr/local/go/src/net/unixsock_posix.go:137
		// _ = "end of CoverTab[8764]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:138
		_go_fuzz_dep_.CoverTab[8765]++
//line /usr/local/go/src/net/unixsock_posix.go:138
		// _ = "end of CoverTab[8765]"
//line /usr/local/go/src/net/unixsock_posix.go:138
	}
//line /usr/local/go/src/net/unixsock_posix.go:138
	// _ = "end of CoverTab[8758]"
//line /usr/local/go/src/net/unixsock_posix.go:138
	_go_fuzz_dep_.CoverTab[8759]++
							sa := &syscall.SockaddrUnix{Name: addr.Name}
							return c.fd.writeTo(b, sa)
//line /usr/local/go/src/net/unixsock_posix.go:140
	// _ = "end of CoverTab[8759]"
}

func (c *UnixConn) writeMsg(b, oob []byte, addr *UnixAddr) (n, oobn int, err error) {
//line /usr/local/go/src/net/unixsock_posix.go:143
	_go_fuzz_dep_.CoverTab[8766]++
							if c.fd.sotype == syscall.SOCK_DGRAM && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:144
		_go_fuzz_dep_.CoverTab[8769]++
//line /usr/local/go/src/net/unixsock_posix.go:144
		return c.fd.isConnected
//line /usr/local/go/src/net/unixsock_posix.go:144
		// _ = "end of CoverTab[8769]"
//line /usr/local/go/src/net/unixsock_posix.go:144
	}() {
//line /usr/local/go/src/net/unixsock_posix.go:144
		_go_fuzz_dep_.CoverTab[8770]++
								return 0, 0, ErrWriteToConnected
//line /usr/local/go/src/net/unixsock_posix.go:145
		// _ = "end of CoverTab[8770]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:146
		_go_fuzz_dep_.CoverTab[8771]++
//line /usr/local/go/src/net/unixsock_posix.go:146
		// _ = "end of CoverTab[8771]"
//line /usr/local/go/src/net/unixsock_posix.go:146
	}
//line /usr/local/go/src/net/unixsock_posix.go:146
	// _ = "end of CoverTab[8766]"
//line /usr/local/go/src/net/unixsock_posix.go:146
	_go_fuzz_dep_.CoverTab[8767]++
							var sa syscall.Sockaddr
							if addr != nil {
//line /usr/local/go/src/net/unixsock_posix.go:148
		_go_fuzz_dep_.CoverTab[8772]++
								if addr.Net != sotypeToNet(c.fd.sotype) {
//line /usr/local/go/src/net/unixsock_posix.go:149
			_go_fuzz_dep_.CoverTab[8774]++
									return 0, 0, syscall.EAFNOSUPPORT
//line /usr/local/go/src/net/unixsock_posix.go:150
			// _ = "end of CoverTab[8774]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:151
			_go_fuzz_dep_.CoverTab[8775]++
//line /usr/local/go/src/net/unixsock_posix.go:151
			// _ = "end of CoverTab[8775]"
//line /usr/local/go/src/net/unixsock_posix.go:151
		}
//line /usr/local/go/src/net/unixsock_posix.go:151
		// _ = "end of CoverTab[8772]"
//line /usr/local/go/src/net/unixsock_posix.go:151
		_go_fuzz_dep_.CoverTab[8773]++
								sa = &syscall.SockaddrUnix{Name: addr.Name}
//line /usr/local/go/src/net/unixsock_posix.go:152
		// _ = "end of CoverTab[8773]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:153
		_go_fuzz_dep_.CoverTab[8776]++
//line /usr/local/go/src/net/unixsock_posix.go:153
		// _ = "end of CoverTab[8776]"
//line /usr/local/go/src/net/unixsock_posix.go:153
	}
//line /usr/local/go/src/net/unixsock_posix.go:153
	// _ = "end of CoverTab[8767]"
//line /usr/local/go/src/net/unixsock_posix.go:153
	_go_fuzz_dep_.CoverTab[8768]++
							return c.fd.writeMsg(b, oob, sa)
//line /usr/local/go/src/net/unixsock_posix.go:154
	// _ = "end of CoverTab[8768]"
}

func (sd *sysDialer) dialUnix(ctx context.Context, laddr, raddr *UnixAddr) (*UnixConn, error) {
//line /usr/local/go/src/net/unixsock_posix.go:157
	_go_fuzz_dep_.CoverTab[8777]++
							ctrlCtxFn := sd.Dialer.ControlContext
							if ctrlCtxFn == nil && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:159
		_go_fuzz_dep_.CoverTab[8780]++
//line /usr/local/go/src/net/unixsock_posix.go:159
		return sd.Dialer.Control != nil
//line /usr/local/go/src/net/unixsock_posix.go:159
		// _ = "end of CoverTab[8780]"
//line /usr/local/go/src/net/unixsock_posix.go:159
	}() {
//line /usr/local/go/src/net/unixsock_posix.go:159
		_go_fuzz_dep_.CoverTab[8781]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/unixsock_posix.go:160
			_go_fuzz_dep_.CoverTab[8782]++
									return sd.Dialer.Control(network, address, c)
//line /usr/local/go/src/net/unixsock_posix.go:161
			// _ = "end of CoverTab[8782]"
		}
//line /usr/local/go/src/net/unixsock_posix.go:162
		// _ = "end of CoverTab[8781]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:163
		_go_fuzz_dep_.CoverTab[8783]++
//line /usr/local/go/src/net/unixsock_posix.go:163
		// _ = "end of CoverTab[8783]"
//line /usr/local/go/src/net/unixsock_posix.go:163
	}
//line /usr/local/go/src/net/unixsock_posix.go:163
	// _ = "end of CoverTab[8777]"
//line /usr/local/go/src/net/unixsock_posix.go:163
	_go_fuzz_dep_.CoverTab[8778]++
							fd, err := unixSocket(ctx, sd.network, laddr, raddr, "dial", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:165
		_go_fuzz_dep_.CoverTab[8784]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:166
		// _ = "end of CoverTab[8784]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:167
		_go_fuzz_dep_.CoverTab[8785]++
//line /usr/local/go/src/net/unixsock_posix.go:167
		// _ = "end of CoverTab[8785]"
//line /usr/local/go/src/net/unixsock_posix.go:167
	}
//line /usr/local/go/src/net/unixsock_posix.go:167
	// _ = "end of CoverTab[8778]"
//line /usr/local/go/src/net/unixsock_posix.go:167
	_go_fuzz_dep_.CoverTab[8779]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/unixsock_posix.go:168
	// _ = "end of CoverTab[8779]"
}

func (ln *UnixListener) accept() (*UnixConn, error) {
//line /usr/local/go/src/net/unixsock_posix.go:171
	_go_fuzz_dep_.CoverTab[8786]++
							fd, err := ln.fd.accept()
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:173
		_go_fuzz_dep_.CoverTab[8788]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:174
		// _ = "end of CoverTab[8788]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:175
		_go_fuzz_dep_.CoverTab[8789]++
//line /usr/local/go/src/net/unixsock_posix.go:175
		// _ = "end of CoverTab[8789]"
//line /usr/local/go/src/net/unixsock_posix.go:175
	}
//line /usr/local/go/src/net/unixsock_posix.go:175
	// _ = "end of CoverTab[8786]"
//line /usr/local/go/src/net/unixsock_posix.go:175
	_go_fuzz_dep_.CoverTab[8787]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/unixsock_posix.go:176
	// _ = "end of CoverTab[8787]"
}

func (ln *UnixListener) close() error {
//line /usr/local/go/src/net/unixsock_posix.go:179
	_go_fuzz_dep_.CoverTab[8790]++

//line /usr/local/go/src/net/unixsock_posix.go:191
	ln.unlinkOnce.Do(func() {
//line /usr/local/go/src/net/unixsock_posix.go:191
		_go_fuzz_dep_.CoverTab[8792]++
								if ln.path[0] != '@' && func() bool {
//line /usr/local/go/src/net/unixsock_posix.go:192
			_go_fuzz_dep_.CoverTab[8793]++
//line /usr/local/go/src/net/unixsock_posix.go:192
			return ln.unlink
//line /usr/local/go/src/net/unixsock_posix.go:192
			// _ = "end of CoverTab[8793]"
//line /usr/local/go/src/net/unixsock_posix.go:192
		}() {
//line /usr/local/go/src/net/unixsock_posix.go:192
			_go_fuzz_dep_.CoverTab[8794]++
									syscall.Unlink(ln.path)
//line /usr/local/go/src/net/unixsock_posix.go:193
			// _ = "end of CoverTab[8794]"
		} else {
//line /usr/local/go/src/net/unixsock_posix.go:194
			_go_fuzz_dep_.CoverTab[8795]++
//line /usr/local/go/src/net/unixsock_posix.go:194
			// _ = "end of CoverTab[8795]"
//line /usr/local/go/src/net/unixsock_posix.go:194
		}
//line /usr/local/go/src/net/unixsock_posix.go:194
		// _ = "end of CoverTab[8792]"
	})
//line /usr/local/go/src/net/unixsock_posix.go:195
	// _ = "end of CoverTab[8790]"
//line /usr/local/go/src/net/unixsock_posix.go:195
	_go_fuzz_dep_.CoverTab[8791]++
							return ln.fd.Close()
//line /usr/local/go/src/net/unixsock_posix.go:196
	// _ = "end of CoverTab[8791]"
}

func (ln *UnixListener) file() (*os.File, error) {
//line /usr/local/go/src/net/unixsock_posix.go:199
	_go_fuzz_dep_.CoverTab[8796]++
							f, err := ln.fd.dup()
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:201
		_go_fuzz_dep_.CoverTab[8798]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:202
		// _ = "end of CoverTab[8798]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:203
		_go_fuzz_dep_.CoverTab[8799]++
//line /usr/local/go/src/net/unixsock_posix.go:203
		// _ = "end of CoverTab[8799]"
//line /usr/local/go/src/net/unixsock_posix.go:203
	}
//line /usr/local/go/src/net/unixsock_posix.go:203
	// _ = "end of CoverTab[8796]"
//line /usr/local/go/src/net/unixsock_posix.go:203
	_go_fuzz_dep_.CoverTab[8797]++
							return f, nil
//line /usr/local/go/src/net/unixsock_posix.go:204
	// _ = "end of CoverTab[8797]"
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
	_go_fuzz_dep_.CoverTab[8800]++
							l.unlink = unlink
//line /usr/local/go/src/net/unixsock_posix.go:216
	// _ = "end of CoverTab[8800]"
}

func (sl *sysListener) listenUnix(ctx context.Context, laddr *UnixAddr) (*UnixListener, error) {
//line /usr/local/go/src/net/unixsock_posix.go:219
	_go_fuzz_dep_.CoverTab[8801]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/unixsock_posix.go:221
		_go_fuzz_dep_.CoverTab[8804]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/unixsock_posix.go:222
			_go_fuzz_dep_.CoverTab[8805]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/unixsock_posix.go:223
			// _ = "end of CoverTab[8805]"
		}
//line /usr/local/go/src/net/unixsock_posix.go:224
		// _ = "end of CoverTab[8804]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:225
		_go_fuzz_dep_.CoverTab[8806]++
//line /usr/local/go/src/net/unixsock_posix.go:225
		// _ = "end of CoverTab[8806]"
//line /usr/local/go/src/net/unixsock_posix.go:225
	}
//line /usr/local/go/src/net/unixsock_posix.go:225
	// _ = "end of CoverTab[8801]"
//line /usr/local/go/src/net/unixsock_posix.go:225
	_go_fuzz_dep_.CoverTab[8802]++
							fd, err := unixSocket(ctx, sl.network, laddr, nil, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:227
		_go_fuzz_dep_.CoverTab[8807]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:228
		// _ = "end of CoverTab[8807]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:229
		_go_fuzz_dep_.CoverTab[8808]++
//line /usr/local/go/src/net/unixsock_posix.go:229
		// _ = "end of CoverTab[8808]"
//line /usr/local/go/src/net/unixsock_posix.go:229
	}
//line /usr/local/go/src/net/unixsock_posix.go:229
	// _ = "end of CoverTab[8802]"
//line /usr/local/go/src/net/unixsock_posix.go:229
	_go_fuzz_dep_.CoverTab[8803]++
							return &UnixListener{fd: fd, path: fd.laddr.String(), unlink: true}, nil
//line /usr/local/go/src/net/unixsock_posix.go:230
	// _ = "end of CoverTab[8803]"
}

func (sl *sysListener) listenUnixgram(ctx context.Context, laddr *UnixAddr) (*UnixConn, error) {
//line /usr/local/go/src/net/unixsock_posix.go:233
	_go_fuzz_dep_.CoverTab[8809]++
							var ctrlCtxFn func(cxt context.Context, network, address string, c syscall.RawConn) error
							if sl.ListenConfig.Control != nil {
//line /usr/local/go/src/net/unixsock_posix.go:235
		_go_fuzz_dep_.CoverTab[8812]++
								ctrlCtxFn = func(cxt context.Context, network, address string, c syscall.RawConn) error {
//line /usr/local/go/src/net/unixsock_posix.go:236
			_go_fuzz_dep_.CoverTab[8813]++
									return sl.ListenConfig.Control(network, address, c)
//line /usr/local/go/src/net/unixsock_posix.go:237
			// _ = "end of CoverTab[8813]"
		}
//line /usr/local/go/src/net/unixsock_posix.go:238
		// _ = "end of CoverTab[8812]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:239
		_go_fuzz_dep_.CoverTab[8814]++
//line /usr/local/go/src/net/unixsock_posix.go:239
		// _ = "end of CoverTab[8814]"
//line /usr/local/go/src/net/unixsock_posix.go:239
	}
//line /usr/local/go/src/net/unixsock_posix.go:239
	// _ = "end of CoverTab[8809]"
//line /usr/local/go/src/net/unixsock_posix.go:239
	_go_fuzz_dep_.CoverTab[8810]++
							fd, err := unixSocket(ctx, sl.network, laddr, nil, "listen", ctrlCtxFn)
							if err != nil {
//line /usr/local/go/src/net/unixsock_posix.go:241
		_go_fuzz_dep_.CoverTab[8815]++
								return nil, err
//line /usr/local/go/src/net/unixsock_posix.go:242
		// _ = "end of CoverTab[8815]"
	} else {
//line /usr/local/go/src/net/unixsock_posix.go:243
		_go_fuzz_dep_.CoverTab[8816]++
//line /usr/local/go/src/net/unixsock_posix.go:243
		// _ = "end of CoverTab[8816]"
//line /usr/local/go/src/net/unixsock_posix.go:243
	}
//line /usr/local/go/src/net/unixsock_posix.go:243
	// _ = "end of CoverTab[8810]"
//line /usr/local/go/src/net/unixsock_posix.go:243
	_go_fuzz_dep_.CoverTab[8811]++
							return newUnixConn(fd), nil
//line /usr/local/go/src/net/unixsock_posix.go:244
	// _ = "end of CoverTab[8811]"
}

//line /usr/local/go/src/net/unixsock_posix.go:245
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/unixsock_posix.go:245
var _ = _go_fuzz_dep_.CoverTab
