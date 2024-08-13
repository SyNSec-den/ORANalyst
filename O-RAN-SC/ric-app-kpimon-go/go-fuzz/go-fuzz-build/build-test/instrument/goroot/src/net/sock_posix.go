// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

//line /snap/go/10455/src/net/sock_posix.go:7
package net

//line /snap/go/10455/src/net/sock_posix.go:7
import (
//line /snap/go/10455/src/net/sock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/sock_posix.go:7
)
//line /snap/go/10455/src/net/sock_posix.go:7
import (
//line /snap/go/10455/src/net/sock_posix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/sock_posix.go:7
)

import (
	"context"
	"internal/poll"
	"os"
	"syscall"
)

// socket returns a network file descriptor that is ready for
//line /snap/go/10455/src/net/sock_posix.go:16
// asynchronous I/O using the network poller.
//line /snap/go/10455/src/net/sock_posix.go:18
func socket(ctx context.Context, net string, family, sotype, proto int, ipv6only bool, laddr, raddr sockaddr, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) (fd *netFD, err error) {
//line /snap/go/10455/src/net/sock_posix.go:18
	_go_fuzz_dep_.CoverTab[8042]++
						s, err := sysSocket(family, sotype, proto)
						if err != nil {
//line /snap/go/10455/src/net/sock_posix.go:20
		_go_fuzz_dep_.CoverTab[529629]++
//line /snap/go/10455/src/net/sock_posix.go:20
		_go_fuzz_dep_.CoverTab[8048]++
							return nil, err
//line /snap/go/10455/src/net/sock_posix.go:21
		// _ = "end of CoverTab[8048]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:22
		_go_fuzz_dep_.CoverTab[529630]++
//line /snap/go/10455/src/net/sock_posix.go:22
		_go_fuzz_dep_.CoverTab[8049]++
//line /snap/go/10455/src/net/sock_posix.go:22
		// _ = "end of CoverTab[8049]"
//line /snap/go/10455/src/net/sock_posix.go:22
	}
//line /snap/go/10455/src/net/sock_posix.go:22
	// _ = "end of CoverTab[8042]"
//line /snap/go/10455/src/net/sock_posix.go:22
	_go_fuzz_dep_.CoverTab[8043]++
						if err = setDefaultSockopts(s, family, sotype, ipv6only); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:23
		_go_fuzz_dep_.CoverTab[529631]++
//line /snap/go/10455/src/net/sock_posix.go:23
		_go_fuzz_dep_.CoverTab[8050]++
							poll.CloseFunc(s)
							return nil, err
//line /snap/go/10455/src/net/sock_posix.go:25
		// _ = "end of CoverTab[8050]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:26
		_go_fuzz_dep_.CoverTab[529632]++
//line /snap/go/10455/src/net/sock_posix.go:26
		_go_fuzz_dep_.CoverTab[8051]++
//line /snap/go/10455/src/net/sock_posix.go:26
		// _ = "end of CoverTab[8051]"
//line /snap/go/10455/src/net/sock_posix.go:26
	}
//line /snap/go/10455/src/net/sock_posix.go:26
	// _ = "end of CoverTab[8043]"
//line /snap/go/10455/src/net/sock_posix.go:26
	_go_fuzz_dep_.CoverTab[8044]++
						if fd, err = newFD(s, family, sotype, net); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:27
		_go_fuzz_dep_.CoverTab[529633]++
//line /snap/go/10455/src/net/sock_posix.go:27
		_go_fuzz_dep_.CoverTab[8052]++
							poll.CloseFunc(s)
							return nil, err
//line /snap/go/10455/src/net/sock_posix.go:29
		// _ = "end of CoverTab[8052]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:30
		_go_fuzz_dep_.CoverTab[529634]++
//line /snap/go/10455/src/net/sock_posix.go:30
		_go_fuzz_dep_.CoverTab[8053]++
//line /snap/go/10455/src/net/sock_posix.go:30
		// _ = "end of CoverTab[8053]"
//line /snap/go/10455/src/net/sock_posix.go:30
	}
//line /snap/go/10455/src/net/sock_posix.go:30
	// _ = "end of CoverTab[8044]"
//line /snap/go/10455/src/net/sock_posix.go:30
	_go_fuzz_dep_.CoverTab[8045]++

//line /snap/go/10455/src/net/sock_posix.go:54
	if laddr != nil && func() bool {
//line /snap/go/10455/src/net/sock_posix.go:54
		_go_fuzz_dep_.CoverTab[8054]++
//line /snap/go/10455/src/net/sock_posix.go:54
		return raddr == nil
//line /snap/go/10455/src/net/sock_posix.go:54
		// _ = "end of CoverTab[8054]"
//line /snap/go/10455/src/net/sock_posix.go:54
	}() {
//line /snap/go/10455/src/net/sock_posix.go:54
		_go_fuzz_dep_.CoverTab[529635]++
//line /snap/go/10455/src/net/sock_posix.go:54
		_go_fuzz_dep_.CoverTab[8055]++
							switch sotype {
		case syscall.SOCK_STREAM, syscall.SOCK_SEQPACKET:
//line /snap/go/10455/src/net/sock_posix.go:56
			_go_fuzz_dep_.CoverTab[529637]++
//line /snap/go/10455/src/net/sock_posix.go:56
			_go_fuzz_dep_.CoverTab[8056]++
								if err := fd.listenStream(ctx, laddr, listenerBacklog(), ctrlCtxFn); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:57
				_go_fuzz_dep_.CoverTab[529640]++
//line /snap/go/10455/src/net/sock_posix.go:57
				_go_fuzz_dep_.CoverTab[8061]++
									fd.Close()
									return nil, err
//line /snap/go/10455/src/net/sock_posix.go:59
				// _ = "end of CoverTab[8061]"
			} else {
//line /snap/go/10455/src/net/sock_posix.go:60
				_go_fuzz_dep_.CoverTab[529641]++
//line /snap/go/10455/src/net/sock_posix.go:60
				_go_fuzz_dep_.CoverTab[8062]++
//line /snap/go/10455/src/net/sock_posix.go:60
				// _ = "end of CoverTab[8062]"
//line /snap/go/10455/src/net/sock_posix.go:60
			}
//line /snap/go/10455/src/net/sock_posix.go:60
			// _ = "end of CoverTab[8056]"
//line /snap/go/10455/src/net/sock_posix.go:60
			_go_fuzz_dep_.CoverTab[8057]++
								return fd, nil
//line /snap/go/10455/src/net/sock_posix.go:61
			// _ = "end of CoverTab[8057]"
		case syscall.SOCK_DGRAM:
//line /snap/go/10455/src/net/sock_posix.go:62
			_go_fuzz_dep_.CoverTab[529638]++
//line /snap/go/10455/src/net/sock_posix.go:62
			_go_fuzz_dep_.CoverTab[8058]++
								if err := fd.listenDatagram(ctx, laddr, ctrlCtxFn); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:63
				_go_fuzz_dep_.CoverTab[529642]++
//line /snap/go/10455/src/net/sock_posix.go:63
				_go_fuzz_dep_.CoverTab[8063]++
									fd.Close()
									return nil, err
//line /snap/go/10455/src/net/sock_posix.go:65
				// _ = "end of CoverTab[8063]"
			} else {
//line /snap/go/10455/src/net/sock_posix.go:66
				_go_fuzz_dep_.CoverTab[529643]++
//line /snap/go/10455/src/net/sock_posix.go:66
				_go_fuzz_dep_.CoverTab[8064]++
//line /snap/go/10455/src/net/sock_posix.go:66
				// _ = "end of CoverTab[8064]"
//line /snap/go/10455/src/net/sock_posix.go:66
			}
//line /snap/go/10455/src/net/sock_posix.go:66
			// _ = "end of CoverTab[8058]"
//line /snap/go/10455/src/net/sock_posix.go:66
			_go_fuzz_dep_.CoverTab[8059]++
								return fd, nil
//line /snap/go/10455/src/net/sock_posix.go:67
			// _ = "end of CoverTab[8059]"
//line /snap/go/10455/src/net/sock_posix.go:67
		default:
//line /snap/go/10455/src/net/sock_posix.go:67
			_go_fuzz_dep_.CoverTab[529639]++
//line /snap/go/10455/src/net/sock_posix.go:67
			_go_fuzz_dep_.CoverTab[8060]++
//line /snap/go/10455/src/net/sock_posix.go:67
			// _ = "end of CoverTab[8060]"
		}
//line /snap/go/10455/src/net/sock_posix.go:68
		// _ = "end of CoverTab[8055]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:69
		_go_fuzz_dep_.CoverTab[529636]++
//line /snap/go/10455/src/net/sock_posix.go:69
		_go_fuzz_dep_.CoverTab[8065]++
//line /snap/go/10455/src/net/sock_posix.go:69
		// _ = "end of CoverTab[8065]"
//line /snap/go/10455/src/net/sock_posix.go:69
	}
//line /snap/go/10455/src/net/sock_posix.go:69
	// _ = "end of CoverTab[8045]"
//line /snap/go/10455/src/net/sock_posix.go:69
	_go_fuzz_dep_.CoverTab[8046]++
						if err := fd.dial(ctx, laddr, raddr, ctrlCtxFn); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:70
		_go_fuzz_dep_.CoverTab[529644]++
//line /snap/go/10455/src/net/sock_posix.go:70
		_go_fuzz_dep_.CoverTab[8066]++
							fd.Close()
							return nil, err
//line /snap/go/10455/src/net/sock_posix.go:72
		// _ = "end of CoverTab[8066]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:73
		_go_fuzz_dep_.CoverTab[529645]++
//line /snap/go/10455/src/net/sock_posix.go:73
		_go_fuzz_dep_.CoverTab[8067]++
//line /snap/go/10455/src/net/sock_posix.go:73
		// _ = "end of CoverTab[8067]"
//line /snap/go/10455/src/net/sock_posix.go:73
	}
//line /snap/go/10455/src/net/sock_posix.go:73
	// _ = "end of CoverTab[8046]"
//line /snap/go/10455/src/net/sock_posix.go:73
	_go_fuzz_dep_.CoverTab[8047]++
						return fd, nil
//line /snap/go/10455/src/net/sock_posix.go:74
	// _ = "end of CoverTab[8047]"
}

func (fd *netFD) ctrlNetwork() string {
//line /snap/go/10455/src/net/sock_posix.go:77
	_go_fuzz_dep_.CoverTab[8068]++
						switch fd.net {
	case "unix", "unixgram", "unixpacket":
//line /snap/go/10455/src/net/sock_posix.go:79
		_go_fuzz_dep_.CoverTab[529646]++
//line /snap/go/10455/src/net/sock_posix.go:79
		_go_fuzz_dep_.CoverTab[8072]++
							return fd.net
//line /snap/go/10455/src/net/sock_posix.go:80
		// _ = "end of CoverTab[8072]"
//line /snap/go/10455/src/net/sock_posix.go:80
	default:
//line /snap/go/10455/src/net/sock_posix.go:80
		_go_fuzz_dep_.CoverTab[529647]++
//line /snap/go/10455/src/net/sock_posix.go:80
		_go_fuzz_dep_.CoverTab[8073]++
//line /snap/go/10455/src/net/sock_posix.go:80
		// _ = "end of CoverTab[8073]"
	}
//line /snap/go/10455/src/net/sock_posix.go:81
	// _ = "end of CoverTab[8068]"
//line /snap/go/10455/src/net/sock_posix.go:81
	_go_fuzz_dep_.CoverTab[8069]++
						switch fd.net[len(fd.net)-1] {
	case '4', '6':
//line /snap/go/10455/src/net/sock_posix.go:83
		_go_fuzz_dep_.CoverTab[529648]++
//line /snap/go/10455/src/net/sock_posix.go:83
		_go_fuzz_dep_.CoverTab[8074]++
							return fd.net
//line /snap/go/10455/src/net/sock_posix.go:84
		// _ = "end of CoverTab[8074]"
//line /snap/go/10455/src/net/sock_posix.go:84
	default:
//line /snap/go/10455/src/net/sock_posix.go:84
		_go_fuzz_dep_.CoverTab[529649]++
//line /snap/go/10455/src/net/sock_posix.go:84
		_go_fuzz_dep_.CoverTab[8075]++
//line /snap/go/10455/src/net/sock_posix.go:84
		// _ = "end of CoverTab[8075]"
	}
//line /snap/go/10455/src/net/sock_posix.go:85
	// _ = "end of CoverTab[8069]"
//line /snap/go/10455/src/net/sock_posix.go:85
	_go_fuzz_dep_.CoverTab[8070]++
						if fd.family == syscall.AF_INET {
//line /snap/go/10455/src/net/sock_posix.go:86
		_go_fuzz_dep_.CoverTab[529650]++
//line /snap/go/10455/src/net/sock_posix.go:86
		_go_fuzz_dep_.CoverTab[8076]++
							return fd.net + "4"
//line /snap/go/10455/src/net/sock_posix.go:87
		// _ = "end of CoverTab[8076]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:88
		_go_fuzz_dep_.CoverTab[529651]++
//line /snap/go/10455/src/net/sock_posix.go:88
		_go_fuzz_dep_.CoverTab[8077]++
//line /snap/go/10455/src/net/sock_posix.go:88
		// _ = "end of CoverTab[8077]"
//line /snap/go/10455/src/net/sock_posix.go:88
	}
//line /snap/go/10455/src/net/sock_posix.go:88
	// _ = "end of CoverTab[8070]"
//line /snap/go/10455/src/net/sock_posix.go:88
	_go_fuzz_dep_.CoverTab[8071]++
						return fd.net + "6"
//line /snap/go/10455/src/net/sock_posix.go:89
	// _ = "end of CoverTab[8071]"
}

func (fd *netFD) addrFunc() func(syscall.Sockaddr) Addr {
//line /snap/go/10455/src/net/sock_posix.go:92
	_go_fuzz_dep_.CoverTab[8078]++
						switch fd.family {
	case syscall.AF_INET, syscall.AF_INET6:
//line /snap/go/10455/src/net/sock_posix.go:94
		_go_fuzz_dep_.CoverTab[529652]++
//line /snap/go/10455/src/net/sock_posix.go:94
		_go_fuzz_dep_.CoverTab[8080]++
							switch fd.sotype {
		case syscall.SOCK_STREAM:
//line /snap/go/10455/src/net/sock_posix.go:96
			_go_fuzz_dep_.CoverTab[529655]++
//line /snap/go/10455/src/net/sock_posix.go:96
			_go_fuzz_dep_.CoverTab[8083]++
								return sockaddrToTCP
//line /snap/go/10455/src/net/sock_posix.go:97
			// _ = "end of CoverTab[8083]"
		case syscall.SOCK_DGRAM:
//line /snap/go/10455/src/net/sock_posix.go:98
			_go_fuzz_dep_.CoverTab[529656]++
//line /snap/go/10455/src/net/sock_posix.go:98
			_go_fuzz_dep_.CoverTab[8084]++
								return sockaddrToUDP
//line /snap/go/10455/src/net/sock_posix.go:99
			// _ = "end of CoverTab[8084]"
		case syscall.SOCK_RAW:
//line /snap/go/10455/src/net/sock_posix.go:100
			_go_fuzz_dep_.CoverTab[529657]++
//line /snap/go/10455/src/net/sock_posix.go:100
			_go_fuzz_dep_.CoverTab[8085]++
									return sockaddrToIP
//line /snap/go/10455/src/net/sock_posix.go:101
			// _ = "end of CoverTab[8085]"
//line /snap/go/10455/src/net/sock_posix.go:101
		default:
//line /snap/go/10455/src/net/sock_posix.go:101
			_go_fuzz_dep_.CoverTab[529658]++
//line /snap/go/10455/src/net/sock_posix.go:101
			_go_fuzz_dep_.CoverTab[8086]++
//line /snap/go/10455/src/net/sock_posix.go:101
			// _ = "end of CoverTab[8086]"
		}
//line /snap/go/10455/src/net/sock_posix.go:102
		// _ = "end of CoverTab[8080]"
	case syscall.AF_UNIX:
//line /snap/go/10455/src/net/sock_posix.go:103
		_go_fuzz_dep_.CoverTab[529653]++
//line /snap/go/10455/src/net/sock_posix.go:103
		_go_fuzz_dep_.CoverTab[8081]++
								switch fd.sotype {
		case syscall.SOCK_STREAM:
//line /snap/go/10455/src/net/sock_posix.go:105
			_go_fuzz_dep_.CoverTab[529659]++
//line /snap/go/10455/src/net/sock_posix.go:105
			_go_fuzz_dep_.CoverTab[8087]++
									return sockaddrToUnix
//line /snap/go/10455/src/net/sock_posix.go:106
			// _ = "end of CoverTab[8087]"
		case syscall.SOCK_DGRAM:
//line /snap/go/10455/src/net/sock_posix.go:107
			_go_fuzz_dep_.CoverTab[529660]++
//line /snap/go/10455/src/net/sock_posix.go:107
			_go_fuzz_dep_.CoverTab[8088]++
									return sockaddrToUnixgram
//line /snap/go/10455/src/net/sock_posix.go:108
			// _ = "end of CoverTab[8088]"
		case syscall.SOCK_SEQPACKET:
//line /snap/go/10455/src/net/sock_posix.go:109
			_go_fuzz_dep_.CoverTab[529661]++
//line /snap/go/10455/src/net/sock_posix.go:109
			_go_fuzz_dep_.CoverTab[8089]++
									return sockaddrToUnixpacket
//line /snap/go/10455/src/net/sock_posix.go:110
			// _ = "end of CoverTab[8089]"
//line /snap/go/10455/src/net/sock_posix.go:110
		default:
//line /snap/go/10455/src/net/sock_posix.go:110
			_go_fuzz_dep_.CoverTab[529662]++
//line /snap/go/10455/src/net/sock_posix.go:110
			_go_fuzz_dep_.CoverTab[8090]++
//line /snap/go/10455/src/net/sock_posix.go:110
			// _ = "end of CoverTab[8090]"
		}
//line /snap/go/10455/src/net/sock_posix.go:111
		// _ = "end of CoverTab[8081]"
//line /snap/go/10455/src/net/sock_posix.go:111
	default:
//line /snap/go/10455/src/net/sock_posix.go:111
		_go_fuzz_dep_.CoverTab[529654]++
//line /snap/go/10455/src/net/sock_posix.go:111
		_go_fuzz_dep_.CoverTab[8082]++
//line /snap/go/10455/src/net/sock_posix.go:111
		// _ = "end of CoverTab[8082]"
	}
//line /snap/go/10455/src/net/sock_posix.go:112
	// _ = "end of CoverTab[8078]"
//line /snap/go/10455/src/net/sock_posix.go:112
	_go_fuzz_dep_.CoverTab[8079]++
							return func(syscall.Sockaddr) Addr {
//line /snap/go/10455/src/net/sock_posix.go:113
		_go_fuzz_dep_.CoverTab[8091]++
//line /snap/go/10455/src/net/sock_posix.go:113
		return nil
//line /snap/go/10455/src/net/sock_posix.go:113
		// _ = "end of CoverTab[8091]"
//line /snap/go/10455/src/net/sock_posix.go:113
	}
//line /snap/go/10455/src/net/sock_posix.go:113
	// _ = "end of CoverTab[8079]"
}

func (fd *netFD) dial(ctx context.Context, laddr, raddr sockaddr, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) error {
//line /snap/go/10455/src/net/sock_posix.go:116
	_go_fuzz_dep_.CoverTab[8092]++
							var c *rawConn
							var err error
							if ctrlCtxFn != nil {
//line /snap/go/10455/src/net/sock_posix.go:119
		_go_fuzz_dep_.CoverTab[529663]++
//line /snap/go/10455/src/net/sock_posix.go:119
		_go_fuzz_dep_.CoverTab[8097]++
								c, err = newRawConn(fd)
								if err != nil {
//line /snap/go/10455/src/net/sock_posix.go:121
			_go_fuzz_dep_.CoverTab[529665]++
//line /snap/go/10455/src/net/sock_posix.go:121
			_go_fuzz_dep_.CoverTab[8100]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:122
			// _ = "end of CoverTab[8100]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:123
			_go_fuzz_dep_.CoverTab[529666]++
//line /snap/go/10455/src/net/sock_posix.go:123
			_go_fuzz_dep_.CoverTab[8101]++
//line /snap/go/10455/src/net/sock_posix.go:123
			// _ = "end of CoverTab[8101]"
//line /snap/go/10455/src/net/sock_posix.go:123
		}
//line /snap/go/10455/src/net/sock_posix.go:123
		// _ = "end of CoverTab[8097]"
//line /snap/go/10455/src/net/sock_posix.go:123
		_go_fuzz_dep_.CoverTab[8098]++
								var ctrlAddr string
								if raddr != nil {
//line /snap/go/10455/src/net/sock_posix.go:125
			_go_fuzz_dep_.CoverTab[529667]++
//line /snap/go/10455/src/net/sock_posix.go:125
			_go_fuzz_dep_.CoverTab[8102]++
									ctrlAddr = raddr.String()
//line /snap/go/10455/src/net/sock_posix.go:126
			// _ = "end of CoverTab[8102]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:127
			_go_fuzz_dep_.CoverTab[529668]++
//line /snap/go/10455/src/net/sock_posix.go:127
			_go_fuzz_dep_.CoverTab[8103]++
//line /snap/go/10455/src/net/sock_posix.go:127
			if laddr != nil {
//line /snap/go/10455/src/net/sock_posix.go:127
				_go_fuzz_dep_.CoverTab[529669]++
//line /snap/go/10455/src/net/sock_posix.go:127
				_go_fuzz_dep_.CoverTab[8104]++
										ctrlAddr = laddr.String()
//line /snap/go/10455/src/net/sock_posix.go:128
				// _ = "end of CoverTab[8104]"
			} else {
//line /snap/go/10455/src/net/sock_posix.go:129
				_go_fuzz_dep_.CoverTab[529670]++
//line /snap/go/10455/src/net/sock_posix.go:129
				_go_fuzz_dep_.CoverTab[8105]++
//line /snap/go/10455/src/net/sock_posix.go:129
				// _ = "end of CoverTab[8105]"
//line /snap/go/10455/src/net/sock_posix.go:129
			}
//line /snap/go/10455/src/net/sock_posix.go:129
			// _ = "end of CoverTab[8103]"
//line /snap/go/10455/src/net/sock_posix.go:129
		}
//line /snap/go/10455/src/net/sock_posix.go:129
		// _ = "end of CoverTab[8098]"
//line /snap/go/10455/src/net/sock_posix.go:129
		_go_fuzz_dep_.CoverTab[8099]++
								if err := ctrlCtxFn(ctx, fd.ctrlNetwork(), ctrlAddr, c); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:130
			_go_fuzz_dep_.CoverTab[529671]++
//line /snap/go/10455/src/net/sock_posix.go:130
			_go_fuzz_dep_.CoverTab[8106]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:131
			// _ = "end of CoverTab[8106]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:132
			_go_fuzz_dep_.CoverTab[529672]++
//line /snap/go/10455/src/net/sock_posix.go:132
			_go_fuzz_dep_.CoverTab[8107]++
//line /snap/go/10455/src/net/sock_posix.go:132
			// _ = "end of CoverTab[8107]"
//line /snap/go/10455/src/net/sock_posix.go:132
		}
//line /snap/go/10455/src/net/sock_posix.go:132
		// _ = "end of CoverTab[8099]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:133
		_go_fuzz_dep_.CoverTab[529664]++
//line /snap/go/10455/src/net/sock_posix.go:133
		_go_fuzz_dep_.CoverTab[8108]++
//line /snap/go/10455/src/net/sock_posix.go:133
		// _ = "end of CoverTab[8108]"
//line /snap/go/10455/src/net/sock_posix.go:133
	}
//line /snap/go/10455/src/net/sock_posix.go:133
	// _ = "end of CoverTab[8092]"
//line /snap/go/10455/src/net/sock_posix.go:133
	_go_fuzz_dep_.CoverTab[8093]++

							var lsa syscall.Sockaddr
							if laddr != nil {
//line /snap/go/10455/src/net/sock_posix.go:136
		_go_fuzz_dep_.CoverTab[529673]++
//line /snap/go/10455/src/net/sock_posix.go:136
		_go_fuzz_dep_.CoverTab[8109]++
								if lsa, err = laddr.sockaddr(fd.family); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:137
			_go_fuzz_dep_.CoverTab[529675]++
//line /snap/go/10455/src/net/sock_posix.go:137
			_go_fuzz_dep_.CoverTab[8110]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:138
			// _ = "end of CoverTab[8110]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:139
			_go_fuzz_dep_.CoverTab[529676]++
//line /snap/go/10455/src/net/sock_posix.go:139
			_go_fuzz_dep_.CoverTab[8111]++
//line /snap/go/10455/src/net/sock_posix.go:139
			if lsa != nil {
//line /snap/go/10455/src/net/sock_posix.go:139
				_go_fuzz_dep_.CoverTab[529677]++
//line /snap/go/10455/src/net/sock_posix.go:139
				_go_fuzz_dep_.CoverTab[8112]++
										if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:140
					_go_fuzz_dep_.CoverTab[529679]++
//line /snap/go/10455/src/net/sock_posix.go:140
					_go_fuzz_dep_.CoverTab[8113]++
											return os.NewSyscallError("bind", err)
//line /snap/go/10455/src/net/sock_posix.go:141
					// _ = "end of CoverTab[8113]"
				} else {
//line /snap/go/10455/src/net/sock_posix.go:142
					_go_fuzz_dep_.CoverTab[529680]++
//line /snap/go/10455/src/net/sock_posix.go:142
					_go_fuzz_dep_.CoverTab[8114]++
//line /snap/go/10455/src/net/sock_posix.go:142
					// _ = "end of CoverTab[8114]"
//line /snap/go/10455/src/net/sock_posix.go:142
				}
//line /snap/go/10455/src/net/sock_posix.go:142
				// _ = "end of CoverTab[8112]"
			} else {
//line /snap/go/10455/src/net/sock_posix.go:143
				_go_fuzz_dep_.CoverTab[529678]++
//line /snap/go/10455/src/net/sock_posix.go:143
				_go_fuzz_dep_.CoverTab[8115]++
//line /snap/go/10455/src/net/sock_posix.go:143
				// _ = "end of CoverTab[8115]"
//line /snap/go/10455/src/net/sock_posix.go:143
			}
//line /snap/go/10455/src/net/sock_posix.go:143
			// _ = "end of CoverTab[8111]"
//line /snap/go/10455/src/net/sock_posix.go:143
		}
//line /snap/go/10455/src/net/sock_posix.go:143
		// _ = "end of CoverTab[8109]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:144
		_go_fuzz_dep_.CoverTab[529674]++
//line /snap/go/10455/src/net/sock_posix.go:144
		_go_fuzz_dep_.CoverTab[8116]++
//line /snap/go/10455/src/net/sock_posix.go:144
		// _ = "end of CoverTab[8116]"
//line /snap/go/10455/src/net/sock_posix.go:144
	}
//line /snap/go/10455/src/net/sock_posix.go:144
	// _ = "end of CoverTab[8093]"
//line /snap/go/10455/src/net/sock_posix.go:144
	_go_fuzz_dep_.CoverTab[8094]++
							var rsa syscall.Sockaddr	// remote address from the user
							var crsa syscall.Sockaddr	// remote address we actually connected to
							if raddr != nil {
//line /snap/go/10455/src/net/sock_posix.go:147
		_go_fuzz_dep_.CoverTab[529681]++
//line /snap/go/10455/src/net/sock_posix.go:147
		_go_fuzz_dep_.CoverTab[8117]++
								if rsa, err = raddr.sockaddr(fd.family); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:148
			_go_fuzz_dep_.CoverTab[529683]++
//line /snap/go/10455/src/net/sock_posix.go:148
			_go_fuzz_dep_.CoverTab[8120]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:149
			// _ = "end of CoverTab[8120]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:150
			_go_fuzz_dep_.CoverTab[529684]++
//line /snap/go/10455/src/net/sock_posix.go:150
			_go_fuzz_dep_.CoverTab[8121]++
//line /snap/go/10455/src/net/sock_posix.go:150
			// _ = "end of CoverTab[8121]"
//line /snap/go/10455/src/net/sock_posix.go:150
		}
//line /snap/go/10455/src/net/sock_posix.go:150
		// _ = "end of CoverTab[8117]"
//line /snap/go/10455/src/net/sock_posix.go:150
		_go_fuzz_dep_.CoverTab[8118]++
								if crsa, err = fd.connect(ctx, lsa, rsa); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:151
			_go_fuzz_dep_.CoverTab[529685]++
//line /snap/go/10455/src/net/sock_posix.go:151
			_go_fuzz_dep_.CoverTab[8122]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:152
			// _ = "end of CoverTab[8122]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:153
			_go_fuzz_dep_.CoverTab[529686]++
//line /snap/go/10455/src/net/sock_posix.go:153
			_go_fuzz_dep_.CoverTab[8123]++
//line /snap/go/10455/src/net/sock_posix.go:153
			// _ = "end of CoverTab[8123]"
//line /snap/go/10455/src/net/sock_posix.go:153
		}
//line /snap/go/10455/src/net/sock_posix.go:153
		// _ = "end of CoverTab[8118]"
//line /snap/go/10455/src/net/sock_posix.go:153
		_go_fuzz_dep_.CoverTab[8119]++
								fd.isConnected = true
//line /snap/go/10455/src/net/sock_posix.go:154
		// _ = "end of CoverTab[8119]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:155
		_go_fuzz_dep_.CoverTab[529682]++
//line /snap/go/10455/src/net/sock_posix.go:155
		_go_fuzz_dep_.CoverTab[8124]++
								if err := fd.init(); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:156
			_go_fuzz_dep_.CoverTab[529687]++
//line /snap/go/10455/src/net/sock_posix.go:156
			_go_fuzz_dep_.CoverTab[8125]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:157
			// _ = "end of CoverTab[8125]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:158
			_go_fuzz_dep_.CoverTab[529688]++
//line /snap/go/10455/src/net/sock_posix.go:158
			_go_fuzz_dep_.CoverTab[8126]++
//line /snap/go/10455/src/net/sock_posix.go:158
			// _ = "end of CoverTab[8126]"
//line /snap/go/10455/src/net/sock_posix.go:158
		}
//line /snap/go/10455/src/net/sock_posix.go:158
		// _ = "end of CoverTab[8124]"
	}
//line /snap/go/10455/src/net/sock_posix.go:159
	// _ = "end of CoverTab[8094]"
//line /snap/go/10455/src/net/sock_posix.go:159
	_go_fuzz_dep_.CoverTab[8095]++

//line /snap/go/10455/src/net/sock_posix.go:166
	lsa, _ = syscall.Getsockname(fd.pfd.Sysfd)
	if crsa != nil {
//line /snap/go/10455/src/net/sock_posix.go:167
		_go_fuzz_dep_.CoverTab[529689]++
//line /snap/go/10455/src/net/sock_posix.go:167
		_go_fuzz_dep_.CoverTab[8127]++
								fd.setAddr(fd.addrFunc()(lsa), fd.addrFunc()(crsa))
//line /snap/go/10455/src/net/sock_posix.go:168
		// _ = "end of CoverTab[8127]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:169
		_go_fuzz_dep_.CoverTab[529690]++
//line /snap/go/10455/src/net/sock_posix.go:169
		_go_fuzz_dep_.CoverTab[8128]++
//line /snap/go/10455/src/net/sock_posix.go:169
		if rsa, _ = syscall.Getpeername(fd.pfd.Sysfd); rsa != nil {
//line /snap/go/10455/src/net/sock_posix.go:169
			_go_fuzz_dep_.CoverTab[529691]++
//line /snap/go/10455/src/net/sock_posix.go:169
			_go_fuzz_dep_.CoverTab[8129]++
									fd.setAddr(fd.addrFunc()(lsa), fd.addrFunc()(rsa))
//line /snap/go/10455/src/net/sock_posix.go:170
			// _ = "end of CoverTab[8129]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:171
			_go_fuzz_dep_.CoverTab[529692]++
//line /snap/go/10455/src/net/sock_posix.go:171
			_go_fuzz_dep_.CoverTab[8130]++
									fd.setAddr(fd.addrFunc()(lsa), raddr)
//line /snap/go/10455/src/net/sock_posix.go:172
			// _ = "end of CoverTab[8130]"
		}
//line /snap/go/10455/src/net/sock_posix.go:173
		// _ = "end of CoverTab[8128]"
//line /snap/go/10455/src/net/sock_posix.go:173
	}
//line /snap/go/10455/src/net/sock_posix.go:173
	// _ = "end of CoverTab[8095]"
//line /snap/go/10455/src/net/sock_posix.go:173
	_go_fuzz_dep_.CoverTab[8096]++
							return nil
//line /snap/go/10455/src/net/sock_posix.go:174
	// _ = "end of CoverTab[8096]"
}

func (fd *netFD) listenStream(ctx context.Context, laddr sockaddr, backlog int, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) error {
//line /snap/go/10455/src/net/sock_posix.go:177
	_go_fuzz_dep_.CoverTab[8131]++
							var err error
							if err = setDefaultListenerSockopts(fd.pfd.Sysfd); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:179
		_go_fuzz_dep_.CoverTab[529693]++
//line /snap/go/10455/src/net/sock_posix.go:179
		_go_fuzz_dep_.CoverTab[8138]++
								return err
//line /snap/go/10455/src/net/sock_posix.go:180
		// _ = "end of CoverTab[8138]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:181
		_go_fuzz_dep_.CoverTab[529694]++
//line /snap/go/10455/src/net/sock_posix.go:181
		_go_fuzz_dep_.CoverTab[8139]++
//line /snap/go/10455/src/net/sock_posix.go:181
		// _ = "end of CoverTab[8139]"
//line /snap/go/10455/src/net/sock_posix.go:181
	}
//line /snap/go/10455/src/net/sock_posix.go:181
	// _ = "end of CoverTab[8131]"
//line /snap/go/10455/src/net/sock_posix.go:181
	_go_fuzz_dep_.CoverTab[8132]++
							var lsa syscall.Sockaddr
							if lsa, err = laddr.sockaddr(fd.family); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:183
		_go_fuzz_dep_.CoverTab[529695]++
//line /snap/go/10455/src/net/sock_posix.go:183
		_go_fuzz_dep_.CoverTab[8140]++
								return err
//line /snap/go/10455/src/net/sock_posix.go:184
		// _ = "end of CoverTab[8140]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:185
		_go_fuzz_dep_.CoverTab[529696]++
//line /snap/go/10455/src/net/sock_posix.go:185
		_go_fuzz_dep_.CoverTab[8141]++
//line /snap/go/10455/src/net/sock_posix.go:185
		// _ = "end of CoverTab[8141]"
//line /snap/go/10455/src/net/sock_posix.go:185
	}
//line /snap/go/10455/src/net/sock_posix.go:185
	// _ = "end of CoverTab[8132]"
//line /snap/go/10455/src/net/sock_posix.go:185
	_go_fuzz_dep_.CoverTab[8133]++

							if ctrlCtxFn != nil {
//line /snap/go/10455/src/net/sock_posix.go:187
		_go_fuzz_dep_.CoverTab[529697]++
//line /snap/go/10455/src/net/sock_posix.go:187
		_go_fuzz_dep_.CoverTab[8142]++
								c, err := newRawConn(fd)
								if err != nil {
//line /snap/go/10455/src/net/sock_posix.go:189
			_go_fuzz_dep_.CoverTab[529699]++
//line /snap/go/10455/src/net/sock_posix.go:189
			_go_fuzz_dep_.CoverTab[8144]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:190
			// _ = "end of CoverTab[8144]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:191
			_go_fuzz_dep_.CoverTab[529700]++
//line /snap/go/10455/src/net/sock_posix.go:191
			_go_fuzz_dep_.CoverTab[8145]++
//line /snap/go/10455/src/net/sock_posix.go:191
			// _ = "end of CoverTab[8145]"
//line /snap/go/10455/src/net/sock_posix.go:191
		}
//line /snap/go/10455/src/net/sock_posix.go:191
		// _ = "end of CoverTab[8142]"
//line /snap/go/10455/src/net/sock_posix.go:191
		_go_fuzz_dep_.CoverTab[8143]++
								if err := ctrlCtxFn(ctx, fd.ctrlNetwork(), laddr.String(), c); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:192
			_go_fuzz_dep_.CoverTab[529701]++
//line /snap/go/10455/src/net/sock_posix.go:192
			_go_fuzz_dep_.CoverTab[8146]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:193
			// _ = "end of CoverTab[8146]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:194
			_go_fuzz_dep_.CoverTab[529702]++
//line /snap/go/10455/src/net/sock_posix.go:194
			_go_fuzz_dep_.CoverTab[8147]++
//line /snap/go/10455/src/net/sock_posix.go:194
			// _ = "end of CoverTab[8147]"
//line /snap/go/10455/src/net/sock_posix.go:194
		}
//line /snap/go/10455/src/net/sock_posix.go:194
		// _ = "end of CoverTab[8143]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:195
		_go_fuzz_dep_.CoverTab[529698]++
//line /snap/go/10455/src/net/sock_posix.go:195
		_go_fuzz_dep_.CoverTab[8148]++
//line /snap/go/10455/src/net/sock_posix.go:195
		// _ = "end of CoverTab[8148]"
//line /snap/go/10455/src/net/sock_posix.go:195
	}
//line /snap/go/10455/src/net/sock_posix.go:195
	// _ = "end of CoverTab[8133]"
//line /snap/go/10455/src/net/sock_posix.go:195
	_go_fuzz_dep_.CoverTab[8134]++

							if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:197
		_go_fuzz_dep_.CoverTab[529703]++
//line /snap/go/10455/src/net/sock_posix.go:197
		_go_fuzz_dep_.CoverTab[8149]++
								return os.NewSyscallError("bind", err)
//line /snap/go/10455/src/net/sock_posix.go:198
		// _ = "end of CoverTab[8149]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:199
		_go_fuzz_dep_.CoverTab[529704]++
//line /snap/go/10455/src/net/sock_posix.go:199
		_go_fuzz_dep_.CoverTab[8150]++
//line /snap/go/10455/src/net/sock_posix.go:199
		// _ = "end of CoverTab[8150]"
//line /snap/go/10455/src/net/sock_posix.go:199
	}
//line /snap/go/10455/src/net/sock_posix.go:199
	// _ = "end of CoverTab[8134]"
//line /snap/go/10455/src/net/sock_posix.go:199
	_go_fuzz_dep_.CoverTab[8135]++
							if err = listenFunc(fd.pfd.Sysfd, backlog); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:200
		_go_fuzz_dep_.CoverTab[529705]++
//line /snap/go/10455/src/net/sock_posix.go:200
		_go_fuzz_dep_.CoverTab[8151]++
								return os.NewSyscallError("listen", err)
//line /snap/go/10455/src/net/sock_posix.go:201
		// _ = "end of CoverTab[8151]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:202
		_go_fuzz_dep_.CoverTab[529706]++
//line /snap/go/10455/src/net/sock_posix.go:202
		_go_fuzz_dep_.CoverTab[8152]++
//line /snap/go/10455/src/net/sock_posix.go:202
		// _ = "end of CoverTab[8152]"
//line /snap/go/10455/src/net/sock_posix.go:202
	}
//line /snap/go/10455/src/net/sock_posix.go:202
	// _ = "end of CoverTab[8135]"
//line /snap/go/10455/src/net/sock_posix.go:202
	_go_fuzz_dep_.CoverTab[8136]++
							if err = fd.init(); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:203
		_go_fuzz_dep_.CoverTab[529707]++
//line /snap/go/10455/src/net/sock_posix.go:203
		_go_fuzz_dep_.CoverTab[8153]++
								return err
//line /snap/go/10455/src/net/sock_posix.go:204
		// _ = "end of CoverTab[8153]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:205
		_go_fuzz_dep_.CoverTab[529708]++
//line /snap/go/10455/src/net/sock_posix.go:205
		_go_fuzz_dep_.CoverTab[8154]++
//line /snap/go/10455/src/net/sock_posix.go:205
		// _ = "end of CoverTab[8154]"
//line /snap/go/10455/src/net/sock_posix.go:205
	}
//line /snap/go/10455/src/net/sock_posix.go:205
	// _ = "end of CoverTab[8136]"
//line /snap/go/10455/src/net/sock_posix.go:205
	_go_fuzz_dep_.CoverTab[8137]++
							lsa, _ = syscall.Getsockname(fd.pfd.Sysfd)
							fd.setAddr(fd.addrFunc()(lsa), nil)
							return nil
//line /snap/go/10455/src/net/sock_posix.go:208
	// _ = "end of CoverTab[8137]"
}

func (fd *netFD) listenDatagram(ctx context.Context, laddr sockaddr, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) error {
//line /snap/go/10455/src/net/sock_posix.go:211
	_go_fuzz_dep_.CoverTab[8155]++
							switch addr := laddr.(type) {
	case *UDPAddr:
//line /snap/go/10455/src/net/sock_posix.go:213
		_go_fuzz_dep_.CoverTab[529709]++
//line /snap/go/10455/src/net/sock_posix.go:213
		_go_fuzz_dep_.CoverTab[8161]++

//line /snap/go/10455/src/net/sock_posix.go:221
		if addr.IP != nil && func() bool {
//line /snap/go/10455/src/net/sock_posix.go:221
			_go_fuzz_dep_.CoverTab[8162]++
//line /snap/go/10455/src/net/sock_posix.go:221
			return addr.IP.IsMulticast()
//line /snap/go/10455/src/net/sock_posix.go:221
			// _ = "end of CoverTab[8162]"
//line /snap/go/10455/src/net/sock_posix.go:221
		}() {
//line /snap/go/10455/src/net/sock_posix.go:221
			_go_fuzz_dep_.CoverTab[529710]++
//line /snap/go/10455/src/net/sock_posix.go:221
			_go_fuzz_dep_.CoverTab[8163]++
									if err := setDefaultMulticastSockopts(fd.pfd.Sysfd); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:222
				_go_fuzz_dep_.CoverTab[529712]++
//line /snap/go/10455/src/net/sock_posix.go:222
				_go_fuzz_dep_.CoverTab[8166]++
										return err
//line /snap/go/10455/src/net/sock_posix.go:223
				// _ = "end of CoverTab[8166]"
			} else {
//line /snap/go/10455/src/net/sock_posix.go:224
				_go_fuzz_dep_.CoverTab[529713]++
//line /snap/go/10455/src/net/sock_posix.go:224
				_go_fuzz_dep_.CoverTab[8167]++
//line /snap/go/10455/src/net/sock_posix.go:224
				// _ = "end of CoverTab[8167]"
//line /snap/go/10455/src/net/sock_posix.go:224
			}
//line /snap/go/10455/src/net/sock_posix.go:224
			// _ = "end of CoverTab[8163]"
//line /snap/go/10455/src/net/sock_posix.go:224
			_go_fuzz_dep_.CoverTab[8164]++
									addr := *addr
									switch fd.family {
			case syscall.AF_INET:
//line /snap/go/10455/src/net/sock_posix.go:227
				_go_fuzz_dep_.CoverTab[529714]++
//line /snap/go/10455/src/net/sock_posix.go:227
				_go_fuzz_dep_.CoverTab[8168]++
										addr.IP = IPv4zero
//line /snap/go/10455/src/net/sock_posix.go:228
				// _ = "end of CoverTab[8168]"
			case syscall.AF_INET6:
//line /snap/go/10455/src/net/sock_posix.go:229
				_go_fuzz_dep_.CoverTab[529715]++
//line /snap/go/10455/src/net/sock_posix.go:229
				_go_fuzz_dep_.CoverTab[8169]++
										addr.IP = IPv6unspecified
//line /snap/go/10455/src/net/sock_posix.go:230
				// _ = "end of CoverTab[8169]"
//line /snap/go/10455/src/net/sock_posix.go:230
			default:
//line /snap/go/10455/src/net/sock_posix.go:230
				_go_fuzz_dep_.CoverTab[529716]++
//line /snap/go/10455/src/net/sock_posix.go:230
				_go_fuzz_dep_.CoverTab[8170]++
//line /snap/go/10455/src/net/sock_posix.go:230
				// _ = "end of CoverTab[8170]"
			}
//line /snap/go/10455/src/net/sock_posix.go:231
			// _ = "end of CoverTab[8164]"
//line /snap/go/10455/src/net/sock_posix.go:231
			_go_fuzz_dep_.CoverTab[8165]++
									laddr = &addr
//line /snap/go/10455/src/net/sock_posix.go:232
			// _ = "end of CoverTab[8165]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:233
			_go_fuzz_dep_.CoverTab[529711]++
//line /snap/go/10455/src/net/sock_posix.go:233
			_go_fuzz_dep_.CoverTab[8171]++
//line /snap/go/10455/src/net/sock_posix.go:233
			// _ = "end of CoverTab[8171]"
//line /snap/go/10455/src/net/sock_posix.go:233
		}
//line /snap/go/10455/src/net/sock_posix.go:233
		// _ = "end of CoverTab[8161]"
	}
//line /snap/go/10455/src/net/sock_posix.go:234
	// _ = "end of CoverTab[8155]"
//line /snap/go/10455/src/net/sock_posix.go:234
	_go_fuzz_dep_.CoverTab[8156]++
							var err error
							var lsa syscall.Sockaddr
							if lsa, err = laddr.sockaddr(fd.family); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:237
		_go_fuzz_dep_.CoverTab[529717]++
//line /snap/go/10455/src/net/sock_posix.go:237
		_go_fuzz_dep_.CoverTab[8172]++
								return err
//line /snap/go/10455/src/net/sock_posix.go:238
		// _ = "end of CoverTab[8172]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:239
		_go_fuzz_dep_.CoverTab[529718]++
//line /snap/go/10455/src/net/sock_posix.go:239
		_go_fuzz_dep_.CoverTab[8173]++
//line /snap/go/10455/src/net/sock_posix.go:239
		// _ = "end of CoverTab[8173]"
//line /snap/go/10455/src/net/sock_posix.go:239
	}
//line /snap/go/10455/src/net/sock_posix.go:239
	// _ = "end of CoverTab[8156]"
//line /snap/go/10455/src/net/sock_posix.go:239
	_go_fuzz_dep_.CoverTab[8157]++

							if ctrlCtxFn != nil {
//line /snap/go/10455/src/net/sock_posix.go:241
		_go_fuzz_dep_.CoverTab[529719]++
//line /snap/go/10455/src/net/sock_posix.go:241
		_go_fuzz_dep_.CoverTab[8174]++
								c, err := newRawConn(fd)
								if err != nil {
//line /snap/go/10455/src/net/sock_posix.go:243
			_go_fuzz_dep_.CoverTab[529721]++
//line /snap/go/10455/src/net/sock_posix.go:243
			_go_fuzz_dep_.CoverTab[8176]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:244
			// _ = "end of CoverTab[8176]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:245
			_go_fuzz_dep_.CoverTab[529722]++
//line /snap/go/10455/src/net/sock_posix.go:245
			_go_fuzz_dep_.CoverTab[8177]++
//line /snap/go/10455/src/net/sock_posix.go:245
			// _ = "end of CoverTab[8177]"
//line /snap/go/10455/src/net/sock_posix.go:245
		}
//line /snap/go/10455/src/net/sock_posix.go:245
		// _ = "end of CoverTab[8174]"
//line /snap/go/10455/src/net/sock_posix.go:245
		_go_fuzz_dep_.CoverTab[8175]++
								if err := ctrlCtxFn(ctx, fd.ctrlNetwork(), laddr.String(), c); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:246
			_go_fuzz_dep_.CoverTab[529723]++
//line /snap/go/10455/src/net/sock_posix.go:246
			_go_fuzz_dep_.CoverTab[8178]++
									return err
//line /snap/go/10455/src/net/sock_posix.go:247
			// _ = "end of CoverTab[8178]"
		} else {
//line /snap/go/10455/src/net/sock_posix.go:248
			_go_fuzz_dep_.CoverTab[529724]++
//line /snap/go/10455/src/net/sock_posix.go:248
			_go_fuzz_dep_.CoverTab[8179]++
//line /snap/go/10455/src/net/sock_posix.go:248
			// _ = "end of CoverTab[8179]"
//line /snap/go/10455/src/net/sock_posix.go:248
		}
//line /snap/go/10455/src/net/sock_posix.go:248
		// _ = "end of CoverTab[8175]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:249
		_go_fuzz_dep_.CoverTab[529720]++
//line /snap/go/10455/src/net/sock_posix.go:249
		_go_fuzz_dep_.CoverTab[8180]++
//line /snap/go/10455/src/net/sock_posix.go:249
		// _ = "end of CoverTab[8180]"
//line /snap/go/10455/src/net/sock_posix.go:249
	}
//line /snap/go/10455/src/net/sock_posix.go:249
	// _ = "end of CoverTab[8157]"
//line /snap/go/10455/src/net/sock_posix.go:249
	_go_fuzz_dep_.CoverTab[8158]++
							if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:250
		_go_fuzz_dep_.CoverTab[529725]++
//line /snap/go/10455/src/net/sock_posix.go:250
		_go_fuzz_dep_.CoverTab[8181]++
								return os.NewSyscallError("bind", err)
//line /snap/go/10455/src/net/sock_posix.go:251
		// _ = "end of CoverTab[8181]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:252
		_go_fuzz_dep_.CoverTab[529726]++
//line /snap/go/10455/src/net/sock_posix.go:252
		_go_fuzz_dep_.CoverTab[8182]++
//line /snap/go/10455/src/net/sock_posix.go:252
		// _ = "end of CoverTab[8182]"
//line /snap/go/10455/src/net/sock_posix.go:252
	}
//line /snap/go/10455/src/net/sock_posix.go:252
	// _ = "end of CoverTab[8158]"
//line /snap/go/10455/src/net/sock_posix.go:252
	_go_fuzz_dep_.CoverTab[8159]++
							if err = fd.init(); err != nil {
//line /snap/go/10455/src/net/sock_posix.go:253
		_go_fuzz_dep_.CoverTab[529727]++
//line /snap/go/10455/src/net/sock_posix.go:253
		_go_fuzz_dep_.CoverTab[8183]++
								return err
//line /snap/go/10455/src/net/sock_posix.go:254
		// _ = "end of CoverTab[8183]"
	} else {
//line /snap/go/10455/src/net/sock_posix.go:255
		_go_fuzz_dep_.CoverTab[529728]++
//line /snap/go/10455/src/net/sock_posix.go:255
		_go_fuzz_dep_.CoverTab[8184]++
//line /snap/go/10455/src/net/sock_posix.go:255
		// _ = "end of CoverTab[8184]"
//line /snap/go/10455/src/net/sock_posix.go:255
	}
//line /snap/go/10455/src/net/sock_posix.go:255
	// _ = "end of CoverTab[8159]"
//line /snap/go/10455/src/net/sock_posix.go:255
	_go_fuzz_dep_.CoverTab[8160]++
							lsa, _ = syscall.Getsockname(fd.pfd.Sysfd)
							fd.setAddr(fd.addrFunc()(lsa), nil)
							return nil
//line /snap/go/10455/src/net/sock_posix.go:258
	// _ = "end of CoverTab[8160]"
}

//line /snap/go/10455/src/net/sock_posix.go:259
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/sock_posix.go:259
var _ = _go_fuzz_dep_.CoverTab
