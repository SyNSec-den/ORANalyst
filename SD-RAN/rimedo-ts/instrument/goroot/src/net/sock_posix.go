// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

//line /usr/local/go/src/net/sock_posix.go:7
package net

//line /usr/local/go/src/net/sock_posix.go:7
import (
//line /usr/local/go/src/net/sock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/sock_posix.go:7
)
//line /usr/local/go/src/net/sock_posix.go:7
import (
//line /usr/local/go/src/net/sock_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/sock_posix.go:7
)

import (
	"context"
	"internal/poll"
	"os"
	"syscall"
)

// socket returns a network file descriptor that is ready for
//line /usr/local/go/src/net/sock_posix.go:16
// asynchronous I/O using the network poller.
//line /usr/local/go/src/net/sock_posix.go:18
func socket(ctx context.Context, net string, family, sotype, proto int, ipv6only bool, laddr, raddr sockaddr, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) (fd *netFD, err error) {
//line /usr/local/go/src/net/sock_posix.go:18
	_go_fuzz_dep_.CoverTab[16138]++
						s, err := sysSocket(family, sotype, proto)
						if err != nil {
//line /usr/local/go/src/net/sock_posix.go:20
		_go_fuzz_dep_.CoverTab[16144]++
							return nil, err
//line /usr/local/go/src/net/sock_posix.go:21
		// _ = "end of CoverTab[16144]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:22
		_go_fuzz_dep_.CoverTab[16145]++
//line /usr/local/go/src/net/sock_posix.go:22
		// _ = "end of CoverTab[16145]"
//line /usr/local/go/src/net/sock_posix.go:22
	}
//line /usr/local/go/src/net/sock_posix.go:22
	// _ = "end of CoverTab[16138]"
//line /usr/local/go/src/net/sock_posix.go:22
	_go_fuzz_dep_.CoverTab[16139]++
						if err = setDefaultSockopts(s, family, sotype, ipv6only); err != nil {
//line /usr/local/go/src/net/sock_posix.go:23
		_go_fuzz_dep_.CoverTab[16146]++
							poll.CloseFunc(s)
							return nil, err
//line /usr/local/go/src/net/sock_posix.go:25
		// _ = "end of CoverTab[16146]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:26
		_go_fuzz_dep_.CoverTab[16147]++
//line /usr/local/go/src/net/sock_posix.go:26
		// _ = "end of CoverTab[16147]"
//line /usr/local/go/src/net/sock_posix.go:26
	}
//line /usr/local/go/src/net/sock_posix.go:26
	// _ = "end of CoverTab[16139]"
//line /usr/local/go/src/net/sock_posix.go:26
	_go_fuzz_dep_.CoverTab[16140]++
						if fd, err = newFD(s, family, sotype, net); err != nil {
//line /usr/local/go/src/net/sock_posix.go:27
		_go_fuzz_dep_.CoverTab[16148]++
							poll.CloseFunc(s)
							return nil, err
//line /usr/local/go/src/net/sock_posix.go:29
		// _ = "end of CoverTab[16148]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:30
		_go_fuzz_dep_.CoverTab[16149]++
//line /usr/local/go/src/net/sock_posix.go:30
		// _ = "end of CoverTab[16149]"
//line /usr/local/go/src/net/sock_posix.go:30
	}
//line /usr/local/go/src/net/sock_posix.go:30
	// _ = "end of CoverTab[16140]"
//line /usr/local/go/src/net/sock_posix.go:30
	_go_fuzz_dep_.CoverTab[16141]++

//line /usr/local/go/src/net/sock_posix.go:54
	if laddr != nil && func() bool {
//line /usr/local/go/src/net/sock_posix.go:54
		_go_fuzz_dep_.CoverTab[16150]++
//line /usr/local/go/src/net/sock_posix.go:54
		return raddr == nil
//line /usr/local/go/src/net/sock_posix.go:54
		// _ = "end of CoverTab[16150]"
//line /usr/local/go/src/net/sock_posix.go:54
	}() {
//line /usr/local/go/src/net/sock_posix.go:54
		_go_fuzz_dep_.CoverTab[16151]++
							switch sotype {
		case syscall.SOCK_STREAM, syscall.SOCK_SEQPACKET:
//line /usr/local/go/src/net/sock_posix.go:56
			_go_fuzz_dep_.CoverTab[16152]++
								if err := fd.listenStream(ctx, laddr, listenerBacklog(), ctrlCtxFn); err != nil {
//line /usr/local/go/src/net/sock_posix.go:57
				_go_fuzz_dep_.CoverTab[16157]++
									fd.Close()
									return nil, err
//line /usr/local/go/src/net/sock_posix.go:59
				// _ = "end of CoverTab[16157]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:60
				_go_fuzz_dep_.CoverTab[16158]++
//line /usr/local/go/src/net/sock_posix.go:60
				// _ = "end of CoverTab[16158]"
//line /usr/local/go/src/net/sock_posix.go:60
			}
//line /usr/local/go/src/net/sock_posix.go:60
			// _ = "end of CoverTab[16152]"
//line /usr/local/go/src/net/sock_posix.go:60
			_go_fuzz_dep_.CoverTab[16153]++
								return fd, nil
//line /usr/local/go/src/net/sock_posix.go:61
			// _ = "end of CoverTab[16153]"
		case syscall.SOCK_DGRAM:
//line /usr/local/go/src/net/sock_posix.go:62
			_go_fuzz_dep_.CoverTab[16154]++
								if err := fd.listenDatagram(ctx, laddr, ctrlCtxFn); err != nil {
//line /usr/local/go/src/net/sock_posix.go:63
				_go_fuzz_dep_.CoverTab[16159]++
									fd.Close()
									return nil, err
//line /usr/local/go/src/net/sock_posix.go:65
				// _ = "end of CoverTab[16159]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:66
				_go_fuzz_dep_.CoverTab[16160]++
//line /usr/local/go/src/net/sock_posix.go:66
				// _ = "end of CoverTab[16160]"
//line /usr/local/go/src/net/sock_posix.go:66
			}
//line /usr/local/go/src/net/sock_posix.go:66
			// _ = "end of CoverTab[16154]"
//line /usr/local/go/src/net/sock_posix.go:66
			_go_fuzz_dep_.CoverTab[16155]++
								return fd, nil
//line /usr/local/go/src/net/sock_posix.go:67
			// _ = "end of CoverTab[16155]"
//line /usr/local/go/src/net/sock_posix.go:67
		default:
//line /usr/local/go/src/net/sock_posix.go:67
			_go_fuzz_dep_.CoverTab[16156]++
//line /usr/local/go/src/net/sock_posix.go:67
			// _ = "end of CoverTab[16156]"
		}
//line /usr/local/go/src/net/sock_posix.go:68
		// _ = "end of CoverTab[16151]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:69
		_go_fuzz_dep_.CoverTab[16161]++
//line /usr/local/go/src/net/sock_posix.go:69
		// _ = "end of CoverTab[16161]"
//line /usr/local/go/src/net/sock_posix.go:69
	}
//line /usr/local/go/src/net/sock_posix.go:69
	// _ = "end of CoverTab[16141]"
//line /usr/local/go/src/net/sock_posix.go:69
	_go_fuzz_dep_.CoverTab[16142]++
						if err := fd.dial(ctx, laddr, raddr, ctrlCtxFn); err != nil {
//line /usr/local/go/src/net/sock_posix.go:70
		_go_fuzz_dep_.CoverTab[16162]++
							fd.Close()
							return nil, err
//line /usr/local/go/src/net/sock_posix.go:72
		// _ = "end of CoverTab[16162]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:73
		_go_fuzz_dep_.CoverTab[16163]++
//line /usr/local/go/src/net/sock_posix.go:73
		// _ = "end of CoverTab[16163]"
//line /usr/local/go/src/net/sock_posix.go:73
	}
//line /usr/local/go/src/net/sock_posix.go:73
	// _ = "end of CoverTab[16142]"
//line /usr/local/go/src/net/sock_posix.go:73
	_go_fuzz_dep_.CoverTab[16143]++
						return fd, nil
//line /usr/local/go/src/net/sock_posix.go:74
	// _ = "end of CoverTab[16143]"
}

func (fd *netFD) ctrlNetwork() string {
//line /usr/local/go/src/net/sock_posix.go:77
	_go_fuzz_dep_.CoverTab[16164]++
						switch fd.net {
	case "unix", "unixgram", "unixpacket":
//line /usr/local/go/src/net/sock_posix.go:79
		_go_fuzz_dep_.CoverTab[16168]++
							return fd.net
//line /usr/local/go/src/net/sock_posix.go:80
		// _ = "end of CoverTab[16168]"
//line /usr/local/go/src/net/sock_posix.go:80
	default:
//line /usr/local/go/src/net/sock_posix.go:80
		_go_fuzz_dep_.CoverTab[16169]++
//line /usr/local/go/src/net/sock_posix.go:80
		// _ = "end of CoverTab[16169]"
	}
//line /usr/local/go/src/net/sock_posix.go:81
	// _ = "end of CoverTab[16164]"
//line /usr/local/go/src/net/sock_posix.go:81
	_go_fuzz_dep_.CoverTab[16165]++
						switch fd.net[len(fd.net)-1] {
	case '4', '6':
//line /usr/local/go/src/net/sock_posix.go:83
		_go_fuzz_dep_.CoverTab[16170]++
							return fd.net
//line /usr/local/go/src/net/sock_posix.go:84
		// _ = "end of CoverTab[16170]"
//line /usr/local/go/src/net/sock_posix.go:84
	default:
//line /usr/local/go/src/net/sock_posix.go:84
		_go_fuzz_dep_.CoverTab[16171]++
//line /usr/local/go/src/net/sock_posix.go:84
		// _ = "end of CoverTab[16171]"
	}
//line /usr/local/go/src/net/sock_posix.go:85
	// _ = "end of CoverTab[16165]"
//line /usr/local/go/src/net/sock_posix.go:85
	_go_fuzz_dep_.CoverTab[16166]++
						if fd.family == syscall.AF_INET {
//line /usr/local/go/src/net/sock_posix.go:86
		_go_fuzz_dep_.CoverTab[16172]++
							return fd.net + "4"
//line /usr/local/go/src/net/sock_posix.go:87
		// _ = "end of CoverTab[16172]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:88
		_go_fuzz_dep_.CoverTab[16173]++
//line /usr/local/go/src/net/sock_posix.go:88
		// _ = "end of CoverTab[16173]"
//line /usr/local/go/src/net/sock_posix.go:88
	}
//line /usr/local/go/src/net/sock_posix.go:88
	// _ = "end of CoverTab[16166]"
//line /usr/local/go/src/net/sock_posix.go:88
	_go_fuzz_dep_.CoverTab[16167]++
						return fd.net + "6"
//line /usr/local/go/src/net/sock_posix.go:89
	// _ = "end of CoverTab[16167]"
}

func (fd *netFD) addrFunc() func(syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/sock_posix.go:92
	_go_fuzz_dep_.CoverTab[16174]++
						switch fd.family {
	case syscall.AF_INET, syscall.AF_INET6:
//line /usr/local/go/src/net/sock_posix.go:94
		_go_fuzz_dep_.CoverTab[16176]++
							switch fd.sotype {
		case syscall.SOCK_STREAM:
//line /usr/local/go/src/net/sock_posix.go:96
			_go_fuzz_dep_.CoverTab[16179]++
								return sockaddrToTCP
//line /usr/local/go/src/net/sock_posix.go:97
			// _ = "end of CoverTab[16179]"
		case syscall.SOCK_DGRAM:
//line /usr/local/go/src/net/sock_posix.go:98
			_go_fuzz_dep_.CoverTab[16180]++
								return sockaddrToUDP
//line /usr/local/go/src/net/sock_posix.go:99
			// _ = "end of CoverTab[16180]"
		case syscall.SOCK_RAW:
//line /usr/local/go/src/net/sock_posix.go:100
			_go_fuzz_dep_.CoverTab[16181]++
								return sockaddrToIP
//line /usr/local/go/src/net/sock_posix.go:101
			// _ = "end of CoverTab[16181]"
//line /usr/local/go/src/net/sock_posix.go:101
		default:
//line /usr/local/go/src/net/sock_posix.go:101
			_go_fuzz_dep_.CoverTab[16182]++
//line /usr/local/go/src/net/sock_posix.go:101
			// _ = "end of CoverTab[16182]"
		}
//line /usr/local/go/src/net/sock_posix.go:102
		// _ = "end of CoverTab[16176]"
	case syscall.AF_UNIX:
//line /usr/local/go/src/net/sock_posix.go:103
		_go_fuzz_dep_.CoverTab[16177]++
							switch fd.sotype {
		case syscall.SOCK_STREAM:
//line /usr/local/go/src/net/sock_posix.go:105
			_go_fuzz_dep_.CoverTab[16183]++
								return sockaddrToUnix
//line /usr/local/go/src/net/sock_posix.go:106
			// _ = "end of CoverTab[16183]"
		case syscall.SOCK_DGRAM:
//line /usr/local/go/src/net/sock_posix.go:107
			_go_fuzz_dep_.CoverTab[16184]++
								return sockaddrToUnixgram
//line /usr/local/go/src/net/sock_posix.go:108
			// _ = "end of CoverTab[16184]"
		case syscall.SOCK_SEQPACKET:
//line /usr/local/go/src/net/sock_posix.go:109
			_go_fuzz_dep_.CoverTab[16185]++
								return sockaddrToUnixpacket
//line /usr/local/go/src/net/sock_posix.go:110
			// _ = "end of CoverTab[16185]"
//line /usr/local/go/src/net/sock_posix.go:110
		default:
//line /usr/local/go/src/net/sock_posix.go:110
			_go_fuzz_dep_.CoverTab[16186]++
//line /usr/local/go/src/net/sock_posix.go:110
			// _ = "end of CoverTab[16186]"
		}
//line /usr/local/go/src/net/sock_posix.go:111
		// _ = "end of CoverTab[16177]"
//line /usr/local/go/src/net/sock_posix.go:111
	default:
//line /usr/local/go/src/net/sock_posix.go:111
		_go_fuzz_dep_.CoverTab[16178]++
//line /usr/local/go/src/net/sock_posix.go:111
		// _ = "end of CoverTab[16178]"
	}
//line /usr/local/go/src/net/sock_posix.go:112
	// _ = "end of CoverTab[16174]"
//line /usr/local/go/src/net/sock_posix.go:112
	_go_fuzz_dep_.CoverTab[16175]++
						return func(syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/sock_posix.go:113
		_go_fuzz_dep_.CoverTab[16187]++
//line /usr/local/go/src/net/sock_posix.go:113
		return nil
//line /usr/local/go/src/net/sock_posix.go:113
		// _ = "end of CoverTab[16187]"
//line /usr/local/go/src/net/sock_posix.go:113
	}
//line /usr/local/go/src/net/sock_posix.go:113
	// _ = "end of CoverTab[16175]"
}

func (fd *netFD) dial(ctx context.Context, laddr, raddr sockaddr, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) error {
//line /usr/local/go/src/net/sock_posix.go:116
	_go_fuzz_dep_.CoverTab[16188]++
						var c *rawConn
						var err error
						if ctrlCtxFn != nil {
//line /usr/local/go/src/net/sock_posix.go:119
		_go_fuzz_dep_.CoverTab[16193]++
							c, err = newRawConn(fd)
							if err != nil {
//line /usr/local/go/src/net/sock_posix.go:121
			_go_fuzz_dep_.CoverTab[16196]++
								return err
//line /usr/local/go/src/net/sock_posix.go:122
			// _ = "end of CoverTab[16196]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:123
			_go_fuzz_dep_.CoverTab[16197]++
//line /usr/local/go/src/net/sock_posix.go:123
			// _ = "end of CoverTab[16197]"
//line /usr/local/go/src/net/sock_posix.go:123
		}
//line /usr/local/go/src/net/sock_posix.go:123
		// _ = "end of CoverTab[16193]"
//line /usr/local/go/src/net/sock_posix.go:123
		_go_fuzz_dep_.CoverTab[16194]++
							var ctrlAddr string
							if raddr != nil {
//line /usr/local/go/src/net/sock_posix.go:125
			_go_fuzz_dep_.CoverTab[16198]++
								ctrlAddr = raddr.String()
//line /usr/local/go/src/net/sock_posix.go:126
			// _ = "end of CoverTab[16198]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:127
			_go_fuzz_dep_.CoverTab[16199]++
//line /usr/local/go/src/net/sock_posix.go:127
			if laddr != nil {
//line /usr/local/go/src/net/sock_posix.go:127
				_go_fuzz_dep_.CoverTab[16200]++
									ctrlAddr = laddr.String()
//line /usr/local/go/src/net/sock_posix.go:128
				// _ = "end of CoverTab[16200]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:129
				_go_fuzz_dep_.CoverTab[16201]++
//line /usr/local/go/src/net/sock_posix.go:129
				// _ = "end of CoverTab[16201]"
//line /usr/local/go/src/net/sock_posix.go:129
			}
//line /usr/local/go/src/net/sock_posix.go:129
			// _ = "end of CoverTab[16199]"
//line /usr/local/go/src/net/sock_posix.go:129
		}
//line /usr/local/go/src/net/sock_posix.go:129
		// _ = "end of CoverTab[16194]"
//line /usr/local/go/src/net/sock_posix.go:129
		_go_fuzz_dep_.CoverTab[16195]++
							if err := ctrlCtxFn(ctx, fd.ctrlNetwork(), ctrlAddr, c); err != nil {
//line /usr/local/go/src/net/sock_posix.go:130
			_go_fuzz_dep_.CoverTab[16202]++
								return err
//line /usr/local/go/src/net/sock_posix.go:131
			// _ = "end of CoverTab[16202]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:132
			_go_fuzz_dep_.CoverTab[16203]++
//line /usr/local/go/src/net/sock_posix.go:132
			// _ = "end of CoverTab[16203]"
//line /usr/local/go/src/net/sock_posix.go:132
		}
//line /usr/local/go/src/net/sock_posix.go:132
		// _ = "end of CoverTab[16195]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:133
		_go_fuzz_dep_.CoverTab[16204]++
//line /usr/local/go/src/net/sock_posix.go:133
		// _ = "end of CoverTab[16204]"
//line /usr/local/go/src/net/sock_posix.go:133
	}
//line /usr/local/go/src/net/sock_posix.go:133
	// _ = "end of CoverTab[16188]"
//line /usr/local/go/src/net/sock_posix.go:133
	_go_fuzz_dep_.CoverTab[16189]++

						var lsa syscall.Sockaddr
						if laddr != nil {
//line /usr/local/go/src/net/sock_posix.go:136
		_go_fuzz_dep_.CoverTab[16205]++
							if lsa, err = laddr.sockaddr(fd.family); err != nil {
//line /usr/local/go/src/net/sock_posix.go:137
			_go_fuzz_dep_.CoverTab[16206]++
								return err
//line /usr/local/go/src/net/sock_posix.go:138
			// _ = "end of CoverTab[16206]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:139
			_go_fuzz_dep_.CoverTab[16207]++
//line /usr/local/go/src/net/sock_posix.go:139
			if lsa != nil {
//line /usr/local/go/src/net/sock_posix.go:139
				_go_fuzz_dep_.CoverTab[16208]++
									if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
//line /usr/local/go/src/net/sock_posix.go:140
					_go_fuzz_dep_.CoverTab[16209]++
										return os.NewSyscallError("bind", err)
//line /usr/local/go/src/net/sock_posix.go:141
					// _ = "end of CoverTab[16209]"
				} else {
//line /usr/local/go/src/net/sock_posix.go:142
					_go_fuzz_dep_.CoverTab[16210]++
//line /usr/local/go/src/net/sock_posix.go:142
					// _ = "end of CoverTab[16210]"
//line /usr/local/go/src/net/sock_posix.go:142
				}
//line /usr/local/go/src/net/sock_posix.go:142
				// _ = "end of CoverTab[16208]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:143
				_go_fuzz_dep_.CoverTab[16211]++
//line /usr/local/go/src/net/sock_posix.go:143
				// _ = "end of CoverTab[16211]"
//line /usr/local/go/src/net/sock_posix.go:143
			}
//line /usr/local/go/src/net/sock_posix.go:143
			// _ = "end of CoverTab[16207]"
//line /usr/local/go/src/net/sock_posix.go:143
		}
//line /usr/local/go/src/net/sock_posix.go:143
		// _ = "end of CoverTab[16205]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:144
		_go_fuzz_dep_.CoverTab[16212]++
//line /usr/local/go/src/net/sock_posix.go:144
		// _ = "end of CoverTab[16212]"
//line /usr/local/go/src/net/sock_posix.go:144
	}
//line /usr/local/go/src/net/sock_posix.go:144
	// _ = "end of CoverTab[16189]"
//line /usr/local/go/src/net/sock_posix.go:144
	_go_fuzz_dep_.CoverTab[16190]++
						var rsa syscall.Sockaddr	// remote address from the user
						var crsa syscall.Sockaddr	// remote address we actually connected to
						if raddr != nil {
//line /usr/local/go/src/net/sock_posix.go:147
		_go_fuzz_dep_.CoverTab[16213]++
							if rsa, err = raddr.sockaddr(fd.family); err != nil {
//line /usr/local/go/src/net/sock_posix.go:148
			_go_fuzz_dep_.CoverTab[16216]++
								return err
//line /usr/local/go/src/net/sock_posix.go:149
			// _ = "end of CoverTab[16216]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:150
			_go_fuzz_dep_.CoverTab[16217]++
//line /usr/local/go/src/net/sock_posix.go:150
			// _ = "end of CoverTab[16217]"
//line /usr/local/go/src/net/sock_posix.go:150
		}
//line /usr/local/go/src/net/sock_posix.go:150
		// _ = "end of CoverTab[16213]"
//line /usr/local/go/src/net/sock_posix.go:150
		_go_fuzz_dep_.CoverTab[16214]++
							if crsa, err = fd.connect(ctx, lsa, rsa); err != nil {
//line /usr/local/go/src/net/sock_posix.go:151
			_go_fuzz_dep_.CoverTab[16218]++
								return err
//line /usr/local/go/src/net/sock_posix.go:152
			// _ = "end of CoverTab[16218]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:153
			_go_fuzz_dep_.CoverTab[16219]++
//line /usr/local/go/src/net/sock_posix.go:153
			// _ = "end of CoverTab[16219]"
//line /usr/local/go/src/net/sock_posix.go:153
		}
//line /usr/local/go/src/net/sock_posix.go:153
		// _ = "end of CoverTab[16214]"
//line /usr/local/go/src/net/sock_posix.go:153
		_go_fuzz_dep_.CoverTab[16215]++
							fd.isConnected = true
//line /usr/local/go/src/net/sock_posix.go:154
		// _ = "end of CoverTab[16215]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:155
		_go_fuzz_dep_.CoverTab[16220]++
							if err := fd.init(); err != nil {
//line /usr/local/go/src/net/sock_posix.go:156
			_go_fuzz_dep_.CoverTab[16221]++
								return err
//line /usr/local/go/src/net/sock_posix.go:157
			// _ = "end of CoverTab[16221]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:158
			_go_fuzz_dep_.CoverTab[16222]++
//line /usr/local/go/src/net/sock_posix.go:158
			// _ = "end of CoverTab[16222]"
//line /usr/local/go/src/net/sock_posix.go:158
		}
//line /usr/local/go/src/net/sock_posix.go:158
		// _ = "end of CoverTab[16220]"
	}
//line /usr/local/go/src/net/sock_posix.go:159
	// _ = "end of CoverTab[16190]"
//line /usr/local/go/src/net/sock_posix.go:159
	_go_fuzz_dep_.CoverTab[16191]++

//line /usr/local/go/src/net/sock_posix.go:166
	lsa, _ = syscall.Getsockname(fd.pfd.Sysfd)
	if crsa != nil {
//line /usr/local/go/src/net/sock_posix.go:167
		_go_fuzz_dep_.CoverTab[16223]++
							fd.setAddr(fd.addrFunc()(lsa), fd.addrFunc()(crsa))
//line /usr/local/go/src/net/sock_posix.go:168
		// _ = "end of CoverTab[16223]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:169
		_go_fuzz_dep_.CoverTab[16224]++
//line /usr/local/go/src/net/sock_posix.go:169
		if rsa, _ = syscall.Getpeername(fd.pfd.Sysfd); rsa != nil {
//line /usr/local/go/src/net/sock_posix.go:169
			_go_fuzz_dep_.CoverTab[16225]++
								fd.setAddr(fd.addrFunc()(lsa), fd.addrFunc()(rsa))
//line /usr/local/go/src/net/sock_posix.go:170
			// _ = "end of CoverTab[16225]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:171
			_go_fuzz_dep_.CoverTab[16226]++
								fd.setAddr(fd.addrFunc()(lsa), raddr)
//line /usr/local/go/src/net/sock_posix.go:172
			// _ = "end of CoverTab[16226]"
		}
//line /usr/local/go/src/net/sock_posix.go:173
		// _ = "end of CoverTab[16224]"
//line /usr/local/go/src/net/sock_posix.go:173
	}
//line /usr/local/go/src/net/sock_posix.go:173
	// _ = "end of CoverTab[16191]"
//line /usr/local/go/src/net/sock_posix.go:173
	_go_fuzz_dep_.CoverTab[16192]++
						return nil
//line /usr/local/go/src/net/sock_posix.go:174
	// _ = "end of CoverTab[16192]"
}

func (fd *netFD) listenStream(ctx context.Context, laddr sockaddr, backlog int, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) error {
//line /usr/local/go/src/net/sock_posix.go:177
	_go_fuzz_dep_.CoverTab[16227]++
						var err error
						if err = setDefaultListenerSockopts(fd.pfd.Sysfd); err != nil {
//line /usr/local/go/src/net/sock_posix.go:179
		_go_fuzz_dep_.CoverTab[16234]++
							return err
//line /usr/local/go/src/net/sock_posix.go:180
		// _ = "end of CoverTab[16234]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:181
		_go_fuzz_dep_.CoverTab[16235]++
//line /usr/local/go/src/net/sock_posix.go:181
		// _ = "end of CoverTab[16235]"
//line /usr/local/go/src/net/sock_posix.go:181
	}
//line /usr/local/go/src/net/sock_posix.go:181
	// _ = "end of CoverTab[16227]"
//line /usr/local/go/src/net/sock_posix.go:181
	_go_fuzz_dep_.CoverTab[16228]++
						var lsa syscall.Sockaddr
						if lsa, err = laddr.sockaddr(fd.family); err != nil {
//line /usr/local/go/src/net/sock_posix.go:183
		_go_fuzz_dep_.CoverTab[16236]++
							return err
//line /usr/local/go/src/net/sock_posix.go:184
		// _ = "end of CoverTab[16236]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:185
		_go_fuzz_dep_.CoverTab[16237]++
//line /usr/local/go/src/net/sock_posix.go:185
		// _ = "end of CoverTab[16237]"
//line /usr/local/go/src/net/sock_posix.go:185
	}
//line /usr/local/go/src/net/sock_posix.go:185
	// _ = "end of CoverTab[16228]"
//line /usr/local/go/src/net/sock_posix.go:185
	_go_fuzz_dep_.CoverTab[16229]++

						if ctrlCtxFn != nil {
//line /usr/local/go/src/net/sock_posix.go:187
		_go_fuzz_dep_.CoverTab[16238]++
							c, err := newRawConn(fd)
							if err != nil {
//line /usr/local/go/src/net/sock_posix.go:189
			_go_fuzz_dep_.CoverTab[16240]++
								return err
//line /usr/local/go/src/net/sock_posix.go:190
			// _ = "end of CoverTab[16240]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:191
			_go_fuzz_dep_.CoverTab[16241]++
//line /usr/local/go/src/net/sock_posix.go:191
			// _ = "end of CoverTab[16241]"
//line /usr/local/go/src/net/sock_posix.go:191
		}
//line /usr/local/go/src/net/sock_posix.go:191
		// _ = "end of CoverTab[16238]"
//line /usr/local/go/src/net/sock_posix.go:191
		_go_fuzz_dep_.CoverTab[16239]++
							if err := ctrlCtxFn(ctx, fd.ctrlNetwork(), laddr.String(), c); err != nil {
//line /usr/local/go/src/net/sock_posix.go:192
			_go_fuzz_dep_.CoverTab[16242]++
								return err
//line /usr/local/go/src/net/sock_posix.go:193
			// _ = "end of CoverTab[16242]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:194
			_go_fuzz_dep_.CoverTab[16243]++
//line /usr/local/go/src/net/sock_posix.go:194
			// _ = "end of CoverTab[16243]"
//line /usr/local/go/src/net/sock_posix.go:194
		}
//line /usr/local/go/src/net/sock_posix.go:194
		// _ = "end of CoverTab[16239]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:195
		_go_fuzz_dep_.CoverTab[16244]++
//line /usr/local/go/src/net/sock_posix.go:195
		// _ = "end of CoverTab[16244]"
//line /usr/local/go/src/net/sock_posix.go:195
	}
//line /usr/local/go/src/net/sock_posix.go:195
	// _ = "end of CoverTab[16229]"
//line /usr/local/go/src/net/sock_posix.go:195
	_go_fuzz_dep_.CoverTab[16230]++

						if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
//line /usr/local/go/src/net/sock_posix.go:197
		_go_fuzz_dep_.CoverTab[16245]++
							return os.NewSyscallError("bind", err)
//line /usr/local/go/src/net/sock_posix.go:198
		// _ = "end of CoverTab[16245]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:199
		_go_fuzz_dep_.CoverTab[16246]++
//line /usr/local/go/src/net/sock_posix.go:199
		// _ = "end of CoverTab[16246]"
//line /usr/local/go/src/net/sock_posix.go:199
	}
//line /usr/local/go/src/net/sock_posix.go:199
	// _ = "end of CoverTab[16230]"
//line /usr/local/go/src/net/sock_posix.go:199
	_go_fuzz_dep_.CoverTab[16231]++
						if err = listenFunc(fd.pfd.Sysfd, backlog); err != nil {
//line /usr/local/go/src/net/sock_posix.go:200
		_go_fuzz_dep_.CoverTab[16247]++
							return os.NewSyscallError("listen", err)
//line /usr/local/go/src/net/sock_posix.go:201
		// _ = "end of CoverTab[16247]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:202
		_go_fuzz_dep_.CoverTab[16248]++
//line /usr/local/go/src/net/sock_posix.go:202
		// _ = "end of CoverTab[16248]"
//line /usr/local/go/src/net/sock_posix.go:202
	}
//line /usr/local/go/src/net/sock_posix.go:202
	// _ = "end of CoverTab[16231]"
//line /usr/local/go/src/net/sock_posix.go:202
	_go_fuzz_dep_.CoverTab[16232]++
						if err = fd.init(); err != nil {
//line /usr/local/go/src/net/sock_posix.go:203
		_go_fuzz_dep_.CoverTab[16249]++
							return err
//line /usr/local/go/src/net/sock_posix.go:204
		// _ = "end of CoverTab[16249]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:205
		_go_fuzz_dep_.CoverTab[16250]++
//line /usr/local/go/src/net/sock_posix.go:205
		// _ = "end of CoverTab[16250]"
//line /usr/local/go/src/net/sock_posix.go:205
	}
//line /usr/local/go/src/net/sock_posix.go:205
	// _ = "end of CoverTab[16232]"
//line /usr/local/go/src/net/sock_posix.go:205
	_go_fuzz_dep_.CoverTab[16233]++
						lsa, _ = syscall.Getsockname(fd.pfd.Sysfd)
						fd.setAddr(fd.addrFunc()(lsa), nil)
						return nil
//line /usr/local/go/src/net/sock_posix.go:208
	// _ = "end of CoverTab[16233]"
}

func (fd *netFD) listenDatagram(ctx context.Context, laddr sockaddr, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) error {
//line /usr/local/go/src/net/sock_posix.go:211
	_go_fuzz_dep_.CoverTab[16251]++
						switch addr := laddr.(type) {
	case *UDPAddr:
//line /usr/local/go/src/net/sock_posix.go:213
		_go_fuzz_dep_.CoverTab[16257]++

//line /usr/local/go/src/net/sock_posix.go:221
		if addr.IP != nil && func() bool {
//line /usr/local/go/src/net/sock_posix.go:221
			_go_fuzz_dep_.CoverTab[16258]++
//line /usr/local/go/src/net/sock_posix.go:221
			return addr.IP.IsMulticast()
//line /usr/local/go/src/net/sock_posix.go:221
			// _ = "end of CoverTab[16258]"
//line /usr/local/go/src/net/sock_posix.go:221
		}() {
//line /usr/local/go/src/net/sock_posix.go:221
			_go_fuzz_dep_.CoverTab[16259]++
								if err := setDefaultMulticastSockopts(fd.pfd.Sysfd); err != nil {
//line /usr/local/go/src/net/sock_posix.go:222
				_go_fuzz_dep_.CoverTab[16262]++
									return err
//line /usr/local/go/src/net/sock_posix.go:223
				// _ = "end of CoverTab[16262]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:224
				_go_fuzz_dep_.CoverTab[16263]++
//line /usr/local/go/src/net/sock_posix.go:224
				// _ = "end of CoverTab[16263]"
//line /usr/local/go/src/net/sock_posix.go:224
			}
//line /usr/local/go/src/net/sock_posix.go:224
			// _ = "end of CoverTab[16259]"
//line /usr/local/go/src/net/sock_posix.go:224
			_go_fuzz_dep_.CoverTab[16260]++
								addr := *addr
								switch fd.family {
			case syscall.AF_INET:
//line /usr/local/go/src/net/sock_posix.go:227
				_go_fuzz_dep_.CoverTab[16264]++
									addr.IP = IPv4zero
//line /usr/local/go/src/net/sock_posix.go:228
				// _ = "end of CoverTab[16264]"
			case syscall.AF_INET6:
//line /usr/local/go/src/net/sock_posix.go:229
				_go_fuzz_dep_.CoverTab[16265]++
									addr.IP = IPv6unspecified
//line /usr/local/go/src/net/sock_posix.go:230
				// _ = "end of CoverTab[16265]"
//line /usr/local/go/src/net/sock_posix.go:230
			default:
//line /usr/local/go/src/net/sock_posix.go:230
				_go_fuzz_dep_.CoverTab[16266]++
//line /usr/local/go/src/net/sock_posix.go:230
				// _ = "end of CoverTab[16266]"
			}
//line /usr/local/go/src/net/sock_posix.go:231
			// _ = "end of CoverTab[16260]"
//line /usr/local/go/src/net/sock_posix.go:231
			_go_fuzz_dep_.CoverTab[16261]++
								laddr = &addr
//line /usr/local/go/src/net/sock_posix.go:232
			// _ = "end of CoverTab[16261]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:233
			_go_fuzz_dep_.CoverTab[16267]++
//line /usr/local/go/src/net/sock_posix.go:233
			// _ = "end of CoverTab[16267]"
//line /usr/local/go/src/net/sock_posix.go:233
		}
//line /usr/local/go/src/net/sock_posix.go:233
		// _ = "end of CoverTab[16257]"
	}
//line /usr/local/go/src/net/sock_posix.go:234
	// _ = "end of CoverTab[16251]"
//line /usr/local/go/src/net/sock_posix.go:234
	_go_fuzz_dep_.CoverTab[16252]++
						var err error
						var lsa syscall.Sockaddr
						if lsa, err = laddr.sockaddr(fd.family); err != nil {
//line /usr/local/go/src/net/sock_posix.go:237
		_go_fuzz_dep_.CoverTab[16268]++
							return err
//line /usr/local/go/src/net/sock_posix.go:238
		// _ = "end of CoverTab[16268]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:239
		_go_fuzz_dep_.CoverTab[16269]++
//line /usr/local/go/src/net/sock_posix.go:239
		// _ = "end of CoverTab[16269]"
//line /usr/local/go/src/net/sock_posix.go:239
	}
//line /usr/local/go/src/net/sock_posix.go:239
	// _ = "end of CoverTab[16252]"
//line /usr/local/go/src/net/sock_posix.go:239
	_go_fuzz_dep_.CoverTab[16253]++

						if ctrlCtxFn != nil {
//line /usr/local/go/src/net/sock_posix.go:241
		_go_fuzz_dep_.CoverTab[16270]++
							c, err := newRawConn(fd)
							if err != nil {
//line /usr/local/go/src/net/sock_posix.go:243
			_go_fuzz_dep_.CoverTab[16272]++
								return err
//line /usr/local/go/src/net/sock_posix.go:244
			// _ = "end of CoverTab[16272]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:245
			_go_fuzz_dep_.CoverTab[16273]++
//line /usr/local/go/src/net/sock_posix.go:245
			// _ = "end of CoverTab[16273]"
//line /usr/local/go/src/net/sock_posix.go:245
		}
//line /usr/local/go/src/net/sock_posix.go:245
		// _ = "end of CoverTab[16270]"
//line /usr/local/go/src/net/sock_posix.go:245
		_go_fuzz_dep_.CoverTab[16271]++
							if err := ctrlCtxFn(ctx, fd.ctrlNetwork(), laddr.String(), c); err != nil {
//line /usr/local/go/src/net/sock_posix.go:246
			_go_fuzz_dep_.CoverTab[16274]++
								return err
//line /usr/local/go/src/net/sock_posix.go:247
			// _ = "end of CoverTab[16274]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:248
			_go_fuzz_dep_.CoverTab[16275]++
//line /usr/local/go/src/net/sock_posix.go:248
			// _ = "end of CoverTab[16275]"
//line /usr/local/go/src/net/sock_posix.go:248
		}
//line /usr/local/go/src/net/sock_posix.go:248
		// _ = "end of CoverTab[16271]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:249
		_go_fuzz_dep_.CoverTab[16276]++
//line /usr/local/go/src/net/sock_posix.go:249
		// _ = "end of CoverTab[16276]"
//line /usr/local/go/src/net/sock_posix.go:249
	}
//line /usr/local/go/src/net/sock_posix.go:249
	// _ = "end of CoverTab[16253]"
//line /usr/local/go/src/net/sock_posix.go:249
	_go_fuzz_dep_.CoverTab[16254]++
						if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
//line /usr/local/go/src/net/sock_posix.go:250
		_go_fuzz_dep_.CoverTab[16277]++
							return os.NewSyscallError("bind", err)
//line /usr/local/go/src/net/sock_posix.go:251
		// _ = "end of CoverTab[16277]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:252
		_go_fuzz_dep_.CoverTab[16278]++
//line /usr/local/go/src/net/sock_posix.go:252
		// _ = "end of CoverTab[16278]"
//line /usr/local/go/src/net/sock_posix.go:252
	}
//line /usr/local/go/src/net/sock_posix.go:252
	// _ = "end of CoverTab[16254]"
//line /usr/local/go/src/net/sock_posix.go:252
	_go_fuzz_dep_.CoverTab[16255]++
						if err = fd.init(); err != nil {
//line /usr/local/go/src/net/sock_posix.go:253
		_go_fuzz_dep_.CoverTab[16279]++
							return err
//line /usr/local/go/src/net/sock_posix.go:254
		// _ = "end of CoverTab[16279]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:255
		_go_fuzz_dep_.CoverTab[16280]++
//line /usr/local/go/src/net/sock_posix.go:255
		// _ = "end of CoverTab[16280]"
//line /usr/local/go/src/net/sock_posix.go:255
	}
//line /usr/local/go/src/net/sock_posix.go:255
	// _ = "end of CoverTab[16255]"
//line /usr/local/go/src/net/sock_posix.go:255
	_go_fuzz_dep_.CoverTab[16256]++
						lsa, _ = syscall.Getsockname(fd.pfd.Sysfd)
						fd.setAddr(fd.addrFunc()(lsa), nil)
						return nil
//line /usr/local/go/src/net/sock_posix.go:258
	// _ = "end of CoverTab[16256]"
}

//line /usr/local/go/src/net/sock_posix.go:259
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sock_posix.go:259
var _ = _go_fuzz_dep_.CoverTab