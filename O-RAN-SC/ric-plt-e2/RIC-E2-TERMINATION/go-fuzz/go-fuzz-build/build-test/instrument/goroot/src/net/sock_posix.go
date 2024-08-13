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
	_go_fuzz_dep_.CoverTab[7748]++
						s, err := sysSocket(family, sotype, proto)
						if err != nil {
//line /usr/local/go/src/net/sock_posix.go:20
		_go_fuzz_dep_.CoverTab[7754]++
							return nil, err
//line /usr/local/go/src/net/sock_posix.go:21
		// _ = "end of CoverTab[7754]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:22
		_go_fuzz_dep_.CoverTab[7755]++
//line /usr/local/go/src/net/sock_posix.go:22
		// _ = "end of CoverTab[7755]"
//line /usr/local/go/src/net/sock_posix.go:22
	}
//line /usr/local/go/src/net/sock_posix.go:22
	// _ = "end of CoverTab[7748]"
//line /usr/local/go/src/net/sock_posix.go:22
	_go_fuzz_dep_.CoverTab[7749]++
						if err = setDefaultSockopts(s, family, sotype, ipv6only); err != nil {
//line /usr/local/go/src/net/sock_posix.go:23
		_go_fuzz_dep_.CoverTab[7756]++
							poll.CloseFunc(s)
							return nil, err
//line /usr/local/go/src/net/sock_posix.go:25
		// _ = "end of CoverTab[7756]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:26
		_go_fuzz_dep_.CoverTab[7757]++
//line /usr/local/go/src/net/sock_posix.go:26
		// _ = "end of CoverTab[7757]"
//line /usr/local/go/src/net/sock_posix.go:26
	}
//line /usr/local/go/src/net/sock_posix.go:26
	// _ = "end of CoverTab[7749]"
//line /usr/local/go/src/net/sock_posix.go:26
	_go_fuzz_dep_.CoverTab[7750]++
						if fd, err = newFD(s, family, sotype, net); err != nil {
//line /usr/local/go/src/net/sock_posix.go:27
		_go_fuzz_dep_.CoverTab[7758]++
							poll.CloseFunc(s)
							return nil, err
//line /usr/local/go/src/net/sock_posix.go:29
		// _ = "end of CoverTab[7758]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:30
		_go_fuzz_dep_.CoverTab[7759]++
//line /usr/local/go/src/net/sock_posix.go:30
		// _ = "end of CoverTab[7759]"
//line /usr/local/go/src/net/sock_posix.go:30
	}
//line /usr/local/go/src/net/sock_posix.go:30
	// _ = "end of CoverTab[7750]"
//line /usr/local/go/src/net/sock_posix.go:30
	_go_fuzz_dep_.CoverTab[7751]++

//line /usr/local/go/src/net/sock_posix.go:54
	if laddr != nil && func() bool {
//line /usr/local/go/src/net/sock_posix.go:54
		_go_fuzz_dep_.CoverTab[7760]++
//line /usr/local/go/src/net/sock_posix.go:54
		return raddr == nil
//line /usr/local/go/src/net/sock_posix.go:54
		// _ = "end of CoverTab[7760]"
//line /usr/local/go/src/net/sock_posix.go:54
	}() {
//line /usr/local/go/src/net/sock_posix.go:54
		_go_fuzz_dep_.CoverTab[7761]++
							switch sotype {
		case syscall.SOCK_STREAM, syscall.SOCK_SEQPACKET:
//line /usr/local/go/src/net/sock_posix.go:56
			_go_fuzz_dep_.CoverTab[7762]++
								if err := fd.listenStream(ctx, laddr, listenerBacklog(), ctrlCtxFn); err != nil {
//line /usr/local/go/src/net/sock_posix.go:57
				_go_fuzz_dep_.CoverTab[7767]++
									fd.Close()
									return nil, err
//line /usr/local/go/src/net/sock_posix.go:59
				// _ = "end of CoverTab[7767]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:60
				_go_fuzz_dep_.CoverTab[7768]++
//line /usr/local/go/src/net/sock_posix.go:60
				// _ = "end of CoverTab[7768]"
//line /usr/local/go/src/net/sock_posix.go:60
			}
//line /usr/local/go/src/net/sock_posix.go:60
			// _ = "end of CoverTab[7762]"
//line /usr/local/go/src/net/sock_posix.go:60
			_go_fuzz_dep_.CoverTab[7763]++
								return fd, nil
//line /usr/local/go/src/net/sock_posix.go:61
			// _ = "end of CoverTab[7763]"
		case syscall.SOCK_DGRAM:
//line /usr/local/go/src/net/sock_posix.go:62
			_go_fuzz_dep_.CoverTab[7764]++
								if err := fd.listenDatagram(ctx, laddr, ctrlCtxFn); err != nil {
//line /usr/local/go/src/net/sock_posix.go:63
				_go_fuzz_dep_.CoverTab[7769]++
									fd.Close()
									return nil, err
//line /usr/local/go/src/net/sock_posix.go:65
				// _ = "end of CoverTab[7769]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:66
				_go_fuzz_dep_.CoverTab[7770]++
//line /usr/local/go/src/net/sock_posix.go:66
				// _ = "end of CoverTab[7770]"
//line /usr/local/go/src/net/sock_posix.go:66
			}
//line /usr/local/go/src/net/sock_posix.go:66
			// _ = "end of CoverTab[7764]"
//line /usr/local/go/src/net/sock_posix.go:66
			_go_fuzz_dep_.CoverTab[7765]++
								return fd, nil
//line /usr/local/go/src/net/sock_posix.go:67
			// _ = "end of CoverTab[7765]"
//line /usr/local/go/src/net/sock_posix.go:67
		default:
//line /usr/local/go/src/net/sock_posix.go:67
			_go_fuzz_dep_.CoverTab[7766]++
//line /usr/local/go/src/net/sock_posix.go:67
			// _ = "end of CoverTab[7766]"
		}
//line /usr/local/go/src/net/sock_posix.go:68
		// _ = "end of CoverTab[7761]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:69
		_go_fuzz_dep_.CoverTab[7771]++
//line /usr/local/go/src/net/sock_posix.go:69
		// _ = "end of CoverTab[7771]"
//line /usr/local/go/src/net/sock_posix.go:69
	}
//line /usr/local/go/src/net/sock_posix.go:69
	// _ = "end of CoverTab[7751]"
//line /usr/local/go/src/net/sock_posix.go:69
	_go_fuzz_dep_.CoverTab[7752]++
						if err := fd.dial(ctx, laddr, raddr, ctrlCtxFn); err != nil {
//line /usr/local/go/src/net/sock_posix.go:70
		_go_fuzz_dep_.CoverTab[7772]++
							fd.Close()
							return nil, err
//line /usr/local/go/src/net/sock_posix.go:72
		// _ = "end of CoverTab[7772]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:73
		_go_fuzz_dep_.CoverTab[7773]++
//line /usr/local/go/src/net/sock_posix.go:73
		// _ = "end of CoverTab[7773]"
//line /usr/local/go/src/net/sock_posix.go:73
	}
//line /usr/local/go/src/net/sock_posix.go:73
	// _ = "end of CoverTab[7752]"
//line /usr/local/go/src/net/sock_posix.go:73
	_go_fuzz_dep_.CoverTab[7753]++
						return fd, nil
//line /usr/local/go/src/net/sock_posix.go:74
	// _ = "end of CoverTab[7753]"
}

func (fd *netFD) ctrlNetwork() string {
//line /usr/local/go/src/net/sock_posix.go:77
	_go_fuzz_dep_.CoverTab[7774]++
						switch fd.net {
	case "unix", "unixgram", "unixpacket":
//line /usr/local/go/src/net/sock_posix.go:79
		_go_fuzz_dep_.CoverTab[7778]++
							return fd.net
//line /usr/local/go/src/net/sock_posix.go:80
		// _ = "end of CoverTab[7778]"
//line /usr/local/go/src/net/sock_posix.go:80
	default:
//line /usr/local/go/src/net/sock_posix.go:80
		_go_fuzz_dep_.CoverTab[7779]++
//line /usr/local/go/src/net/sock_posix.go:80
		// _ = "end of CoverTab[7779]"
	}
//line /usr/local/go/src/net/sock_posix.go:81
	// _ = "end of CoverTab[7774]"
//line /usr/local/go/src/net/sock_posix.go:81
	_go_fuzz_dep_.CoverTab[7775]++
						switch fd.net[len(fd.net)-1] {
	case '4', '6':
//line /usr/local/go/src/net/sock_posix.go:83
		_go_fuzz_dep_.CoverTab[7780]++
							return fd.net
//line /usr/local/go/src/net/sock_posix.go:84
		// _ = "end of CoverTab[7780]"
//line /usr/local/go/src/net/sock_posix.go:84
	default:
//line /usr/local/go/src/net/sock_posix.go:84
		_go_fuzz_dep_.CoverTab[7781]++
//line /usr/local/go/src/net/sock_posix.go:84
		// _ = "end of CoverTab[7781]"
	}
//line /usr/local/go/src/net/sock_posix.go:85
	// _ = "end of CoverTab[7775]"
//line /usr/local/go/src/net/sock_posix.go:85
	_go_fuzz_dep_.CoverTab[7776]++
						if fd.family == syscall.AF_INET {
//line /usr/local/go/src/net/sock_posix.go:86
		_go_fuzz_dep_.CoverTab[7782]++
							return fd.net + "4"
//line /usr/local/go/src/net/sock_posix.go:87
		// _ = "end of CoverTab[7782]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:88
		_go_fuzz_dep_.CoverTab[7783]++
//line /usr/local/go/src/net/sock_posix.go:88
		// _ = "end of CoverTab[7783]"
//line /usr/local/go/src/net/sock_posix.go:88
	}
//line /usr/local/go/src/net/sock_posix.go:88
	// _ = "end of CoverTab[7776]"
//line /usr/local/go/src/net/sock_posix.go:88
	_go_fuzz_dep_.CoverTab[7777]++
						return fd.net + "6"
//line /usr/local/go/src/net/sock_posix.go:89
	// _ = "end of CoverTab[7777]"
}

func (fd *netFD) addrFunc() func(syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/sock_posix.go:92
	_go_fuzz_dep_.CoverTab[7784]++
						switch fd.family {
	case syscall.AF_INET, syscall.AF_INET6:
//line /usr/local/go/src/net/sock_posix.go:94
		_go_fuzz_dep_.CoverTab[7786]++
							switch fd.sotype {
		case syscall.SOCK_STREAM:
//line /usr/local/go/src/net/sock_posix.go:96
			_go_fuzz_dep_.CoverTab[7789]++
								return sockaddrToTCP
//line /usr/local/go/src/net/sock_posix.go:97
			// _ = "end of CoverTab[7789]"
		case syscall.SOCK_DGRAM:
//line /usr/local/go/src/net/sock_posix.go:98
			_go_fuzz_dep_.CoverTab[7790]++
								return sockaddrToUDP
//line /usr/local/go/src/net/sock_posix.go:99
			// _ = "end of CoverTab[7790]"
		case syscall.SOCK_RAW:
//line /usr/local/go/src/net/sock_posix.go:100
			_go_fuzz_dep_.CoverTab[7791]++
								return sockaddrToIP
//line /usr/local/go/src/net/sock_posix.go:101
			// _ = "end of CoverTab[7791]"
//line /usr/local/go/src/net/sock_posix.go:101
		default:
//line /usr/local/go/src/net/sock_posix.go:101
			_go_fuzz_dep_.CoverTab[7792]++
//line /usr/local/go/src/net/sock_posix.go:101
			// _ = "end of CoverTab[7792]"
		}
//line /usr/local/go/src/net/sock_posix.go:102
		// _ = "end of CoverTab[7786]"
	case syscall.AF_UNIX:
//line /usr/local/go/src/net/sock_posix.go:103
		_go_fuzz_dep_.CoverTab[7787]++
							switch fd.sotype {
		case syscall.SOCK_STREAM:
//line /usr/local/go/src/net/sock_posix.go:105
			_go_fuzz_dep_.CoverTab[7793]++
								return sockaddrToUnix
//line /usr/local/go/src/net/sock_posix.go:106
			// _ = "end of CoverTab[7793]"
		case syscall.SOCK_DGRAM:
//line /usr/local/go/src/net/sock_posix.go:107
			_go_fuzz_dep_.CoverTab[7794]++
								return sockaddrToUnixgram
//line /usr/local/go/src/net/sock_posix.go:108
			// _ = "end of CoverTab[7794]"
		case syscall.SOCK_SEQPACKET:
//line /usr/local/go/src/net/sock_posix.go:109
			_go_fuzz_dep_.CoverTab[7795]++
								return sockaddrToUnixpacket
//line /usr/local/go/src/net/sock_posix.go:110
			// _ = "end of CoverTab[7795]"
//line /usr/local/go/src/net/sock_posix.go:110
		default:
//line /usr/local/go/src/net/sock_posix.go:110
			_go_fuzz_dep_.CoverTab[7796]++
//line /usr/local/go/src/net/sock_posix.go:110
			// _ = "end of CoverTab[7796]"
		}
//line /usr/local/go/src/net/sock_posix.go:111
		// _ = "end of CoverTab[7787]"
//line /usr/local/go/src/net/sock_posix.go:111
	default:
//line /usr/local/go/src/net/sock_posix.go:111
		_go_fuzz_dep_.CoverTab[7788]++
//line /usr/local/go/src/net/sock_posix.go:111
		// _ = "end of CoverTab[7788]"
	}
//line /usr/local/go/src/net/sock_posix.go:112
	// _ = "end of CoverTab[7784]"
//line /usr/local/go/src/net/sock_posix.go:112
	_go_fuzz_dep_.CoverTab[7785]++
						return func(syscall.Sockaddr) Addr {
//line /usr/local/go/src/net/sock_posix.go:113
		_go_fuzz_dep_.CoverTab[7797]++
//line /usr/local/go/src/net/sock_posix.go:113
		return nil
//line /usr/local/go/src/net/sock_posix.go:113
		// _ = "end of CoverTab[7797]"
//line /usr/local/go/src/net/sock_posix.go:113
	}
//line /usr/local/go/src/net/sock_posix.go:113
	// _ = "end of CoverTab[7785]"
}

func (fd *netFD) dial(ctx context.Context, laddr, raddr sockaddr, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) error {
//line /usr/local/go/src/net/sock_posix.go:116
	_go_fuzz_dep_.CoverTab[7798]++
						var c *rawConn
						var err error
						if ctrlCtxFn != nil {
//line /usr/local/go/src/net/sock_posix.go:119
		_go_fuzz_dep_.CoverTab[7803]++
							c, err = newRawConn(fd)
							if err != nil {
//line /usr/local/go/src/net/sock_posix.go:121
			_go_fuzz_dep_.CoverTab[7806]++
								return err
//line /usr/local/go/src/net/sock_posix.go:122
			// _ = "end of CoverTab[7806]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:123
			_go_fuzz_dep_.CoverTab[7807]++
//line /usr/local/go/src/net/sock_posix.go:123
			// _ = "end of CoverTab[7807]"
//line /usr/local/go/src/net/sock_posix.go:123
		}
//line /usr/local/go/src/net/sock_posix.go:123
		// _ = "end of CoverTab[7803]"
//line /usr/local/go/src/net/sock_posix.go:123
		_go_fuzz_dep_.CoverTab[7804]++
							var ctrlAddr string
							if raddr != nil {
//line /usr/local/go/src/net/sock_posix.go:125
			_go_fuzz_dep_.CoverTab[7808]++
								ctrlAddr = raddr.String()
//line /usr/local/go/src/net/sock_posix.go:126
			// _ = "end of CoverTab[7808]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:127
			_go_fuzz_dep_.CoverTab[7809]++
//line /usr/local/go/src/net/sock_posix.go:127
			if laddr != nil {
//line /usr/local/go/src/net/sock_posix.go:127
				_go_fuzz_dep_.CoverTab[7810]++
									ctrlAddr = laddr.String()
//line /usr/local/go/src/net/sock_posix.go:128
				// _ = "end of CoverTab[7810]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:129
				_go_fuzz_dep_.CoverTab[7811]++
//line /usr/local/go/src/net/sock_posix.go:129
				// _ = "end of CoverTab[7811]"
//line /usr/local/go/src/net/sock_posix.go:129
			}
//line /usr/local/go/src/net/sock_posix.go:129
			// _ = "end of CoverTab[7809]"
//line /usr/local/go/src/net/sock_posix.go:129
		}
//line /usr/local/go/src/net/sock_posix.go:129
		// _ = "end of CoverTab[7804]"
//line /usr/local/go/src/net/sock_posix.go:129
		_go_fuzz_dep_.CoverTab[7805]++
							if err := ctrlCtxFn(ctx, fd.ctrlNetwork(), ctrlAddr, c); err != nil {
//line /usr/local/go/src/net/sock_posix.go:130
			_go_fuzz_dep_.CoverTab[7812]++
								return err
//line /usr/local/go/src/net/sock_posix.go:131
			// _ = "end of CoverTab[7812]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:132
			_go_fuzz_dep_.CoverTab[7813]++
//line /usr/local/go/src/net/sock_posix.go:132
			// _ = "end of CoverTab[7813]"
//line /usr/local/go/src/net/sock_posix.go:132
		}
//line /usr/local/go/src/net/sock_posix.go:132
		// _ = "end of CoverTab[7805]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:133
		_go_fuzz_dep_.CoverTab[7814]++
//line /usr/local/go/src/net/sock_posix.go:133
		// _ = "end of CoverTab[7814]"
//line /usr/local/go/src/net/sock_posix.go:133
	}
//line /usr/local/go/src/net/sock_posix.go:133
	// _ = "end of CoverTab[7798]"
//line /usr/local/go/src/net/sock_posix.go:133
	_go_fuzz_dep_.CoverTab[7799]++

						var lsa syscall.Sockaddr
						if laddr != nil {
//line /usr/local/go/src/net/sock_posix.go:136
		_go_fuzz_dep_.CoverTab[7815]++
							if lsa, err = laddr.sockaddr(fd.family); err != nil {
//line /usr/local/go/src/net/sock_posix.go:137
			_go_fuzz_dep_.CoverTab[7816]++
								return err
//line /usr/local/go/src/net/sock_posix.go:138
			// _ = "end of CoverTab[7816]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:139
			_go_fuzz_dep_.CoverTab[7817]++
//line /usr/local/go/src/net/sock_posix.go:139
			if lsa != nil {
//line /usr/local/go/src/net/sock_posix.go:139
				_go_fuzz_dep_.CoverTab[7818]++
									if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
//line /usr/local/go/src/net/sock_posix.go:140
					_go_fuzz_dep_.CoverTab[7819]++
										return os.NewSyscallError("bind", err)
//line /usr/local/go/src/net/sock_posix.go:141
					// _ = "end of CoverTab[7819]"
				} else {
//line /usr/local/go/src/net/sock_posix.go:142
					_go_fuzz_dep_.CoverTab[7820]++
//line /usr/local/go/src/net/sock_posix.go:142
					// _ = "end of CoverTab[7820]"
//line /usr/local/go/src/net/sock_posix.go:142
				}
//line /usr/local/go/src/net/sock_posix.go:142
				// _ = "end of CoverTab[7818]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:143
				_go_fuzz_dep_.CoverTab[7821]++
//line /usr/local/go/src/net/sock_posix.go:143
				// _ = "end of CoverTab[7821]"
//line /usr/local/go/src/net/sock_posix.go:143
			}
//line /usr/local/go/src/net/sock_posix.go:143
			// _ = "end of CoverTab[7817]"
//line /usr/local/go/src/net/sock_posix.go:143
		}
//line /usr/local/go/src/net/sock_posix.go:143
		// _ = "end of CoverTab[7815]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:144
		_go_fuzz_dep_.CoverTab[7822]++
//line /usr/local/go/src/net/sock_posix.go:144
		// _ = "end of CoverTab[7822]"
//line /usr/local/go/src/net/sock_posix.go:144
	}
//line /usr/local/go/src/net/sock_posix.go:144
	// _ = "end of CoverTab[7799]"
//line /usr/local/go/src/net/sock_posix.go:144
	_go_fuzz_dep_.CoverTab[7800]++
						var rsa syscall.Sockaddr	// remote address from the user
						var crsa syscall.Sockaddr	// remote address we actually connected to
						if raddr != nil {
//line /usr/local/go/src/net/sock_posix.go:147
		_go_fuzz_dep_.CoverTab[7823]++
							if rsa, err = raddr.sockaddr(fd.family); err != nil {
//line /usr/local/go/src/net/sock_posix.go:148
			_go_fuzz_dep_.CoverTab[7826]++
								return err
//line /usr/local/go/src/net/sock_posix.go:149
			// _ = "end of CoverTab[7826]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:150
			_go_fuzz_dep_.CoverTab[7827]++
//line /usr/local/go/src/net/sock_posix.go:150
			// _ = "end of CoverTab[7827]"
//line /usr/local/go/src/net/sock_posix.go:150
		}
//line /usr/local/go/src/net/sock_posix.go:150
		// _ = "end of CoverTab[7823]"
//line /usr/local/go/src/net/sock_posix.go:150
		_go_fuzz_dep_.CoverTab[7824]++
							if crsa, err = fd.connect(ctx, lsa, rsa); err != nil {
//line /usr/local/go/src/net/sock_posix.go:151
			_go_fuzz_dep_.CoverTab[7828]++
								return err
//line /usr/local/go/src/net/sock_posix.go:152
			// _ = "end of CoverTab[7828]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:153
			_go_fuzz_dep_.CoverTab[7829]++
//line /usr/local/go/src/net/sock_posix.go:153
			// _ = "end of CoverTab[7829]"
//line /usr/local/go/src/net/sock_posix.go:153
		}
//line /usr/local/go/src/net/sock_posix.go:153
		// _ = "end of CoverTab[7824]"
//line /usr/local/go/src/net/sock_posix.go:153
		_go_fuzz_dep_.CoverTab[7825]++
							fd.isConnected = true
//line /usr/local/go/src/net/sock_posix.go:154
		// _ = "end of CoverTab[7825]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:155
		_go_fuzz_dep_.CoverTab[7830]++
							if err := fd.init(); err != nil {
//line /usr/local/go/src/net/sock_posix.go:156
			_go_fuzz_dep_.CoverTab[7831]++
								return err
//line /usr/local/go/src/net/sock_posix.go:157
			// _ = "end of CoverTab[7831]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:158
			_go_fuzz_dep_.CoverTab[7832]++
//line /usr/local/go/src/net/sock_posix.go:158
			// _ = "end of CoverTab[7832]"
//line /usr/local/go/src/net/sock_posix.go:158
		}
//line /usr/local/go/src/net/sock_posix.go:158
		// _ = "end of CoverTab[7830]"
	}
//line /usr/local/go/src/net/sock_posix.go:159
	// _ = "end of CoverTab[7800]"
//line /usr/local/go/src/net/sock_posix.go:159
	_go_fuzz_dep_.CoverTab[7801]++

//line /usr/local/go/src/net/sock_posix.go:166
	lsa, _ = syscall.Getsockname(fd.pfd.Sysfd)
	if crsa != nil {
//line /usr/local/go/src/net/sock_posix.go:167
		_go_fuzz_dep_.CoverTab[7833]++
							fd.setAddr(fd.addrFunc()(lsa), fd.addrFunc()(crsa))
//line /usr/local/go/src/net/sock_posix.go:168
		// _ = "end of CoverTab[7833]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:169
		_go_fuzz_dep_.CoverTab[7834]++
//line /usr/local/go/src/net/sock_posix.go:169
		if rsa, _ = syscall.Getpeername(fd.pfd.Sysfd); rsa != nil {
//line /usr/local/go/src/net/sock_posix.go:169
			_go_fuzz_dep_.CoverTab[7835]++
								fd.setAddr(fd.addrFunc()(lsa), fd.addrFunc()(rsa))
//line /usr/local/go/src/net/sock_posix.go:170
			// _ = "end of CoverTab[7835]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:171
			_go_fuzz_dep_.CoverTab[7836]++
								fd.setAddr(fd.addrFunc()(lsa), raddr)
//line /usr/local/go/src/net/sock_posix.go:172
			// _ = "end of CoverTab[7836]"
		}
//line /usr/local/go/src/net/sock_posix.go:173
		// _ = "end of CoverTab[7834]"
//line /usr/local/go/src/net/sock_posix.go:173
	}
//line /usr/local/go/src/net/sock_posix.go:173
	// _ = "end of CoverTab[7801]"
//line /usr/local/go/src/net/sock_posix.go:173
	_go_fuzz_dep_.CoverTab[7802]++
						return nil
//line /usr/local/go/src/net/sock_posix.go:174
	// _ = "end of CoverTab[7802]"
}

func (fd *netFD) listenStream(ctx context.Context, laddr sockaddr, backlog int, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) error {
//line /usr/local/go/src/net/sock_posix.go:177
	_go_fuzz_dep_.CoverTab[7837]++
						var err error
						if err = setDefaultListenerSockopts(fd.pfd.Sysfd); err != nil {
//line /usr/local/go/src/net/sock_posix.go:179
		_go_fuzz_dep_.CoverTab[7844]++
							return err
//line /usr/local/go/src/net/sock_posix.go:180
		// _ = "end of CoverTab[7844]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:181
		_go_fuzz_dep_.CoverTab[7845]++
//line /usr/local/go/src/net/sock_posix.go:181
		// _ = "end of CoverTab[7845]"
//line /usr/local/go/src/net/sock_posix.go:181
	}
//line /usr/local/go/src/net/sock_posix.go:181
	// _ = "end of CoverTab[7837]"
//line /usr/local/go/src/net/sock_posix.go:181
	_go_fuzz_dep_.CoverTab[7838]++
						var lsa syscall.Sockaddr
						if lsa, err = laddr.sockaddr(fd.family); err != nil {
//line /usr/local/go/src/net/sock_posix.go:183
		_go_fuzz_dep_.CoverTab[7846]++
							return err
//line /usr/local/go/src/net/sock_posix.go:184
		// _ = "end of CoverTab[7846]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:185
		_go_fuzz_dep_.CoverTab[7847]++
//line /usr/local/go/src/net/sock_posix.go:185
		// _ = "end of CoverTab[7847]"
//line /usr/local/go/src/net/sock_posix.go:185
	}
//line /usr/local/go/src/net/sock_posix.go:185
	// _ = "end of CoverTab[7838]"
//line /usr/local/go/src/net/sock_posix.go:185
	_go_fuzz_dep_.CoverTab[7839]++

						if ctrlCtxFn != nil {
//line /usr/local/go/src/net/sock_posix.go:187
		_go_fuzz_dep_.CoverTab[7848]++
							c, err := newRawConn(fd)
							if err != nil {
//line /usr/local/go/src/net/sock_posix.go:189
			_go_fuzz_dep_.CoverTab[7850]++
								return err
//line /usr/local/go/src/net/sock_posix.go:190
			// _ = "end of CoverTab[7850]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:191
			_go_fuzz_dep_.CoverTab[7851]++
//line /usr/local/go/src/net/sock_posix.go:191
			// _ = "end of CoverTab[7851]"
//line /usr/local/go/src/net/sock_posix.go:191
		}
//line /usr/local/go/src/net/sock_posix.go:191
		// _ = "end of CoverTab[7848]"
//line /usr/local/go/src/net/sock_posix.go:191
		_go_fuzz_dep_.CoverTab[7849]++
							if err := ctrlCtxFn(ctx, fd.ctrlNetwork(), laddr.String(), c); err != nil {
//line /usr/local/go/src/net/sock_posix.go:192
			_go_fuzz_dep_.CoverTab[7852]++
								return err
//line /usr/local/go/src/net/sock_posix.go:193
			// _ = "end of CoverTab[7852]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:194
			_go_fuzz_dep_.CoverTab[7853]++
//line /usr/local/go/src/net/sock_posix.go:194
			// _ = "end of CoverTab[7853]"
//line /usr/local/go/src/net/sock_posix.go:194
		}
//line /usr/local/go/src/net/sock_posix.go:194
		// _ = "end of CoverTab[7849]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:195
		_go_fuzz_dep_.CoverTab[7854]++
//line /usr/local/go/src/net/sock_posix.go:195
		// _ = "end of CoverTab[7854]"
//line /usr/local/go/src/net/sock_posix.go:195
	}
//line /usr/local/go/src/net/sock_posix.go:195
	// _ = "end of CoverTab[7839]"
//line /usr/local/go/src/net/sock_posix.go:195
	_go_fuzz_dep_.CoverTab[7840]++

						if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
//line /usr/local/go/src/net/sock_posix.go:197
		_go_fuzz_dep_.CoverTab[7855]++
							return os.NewSyscallError("bind", err)
//line /usr/local/go/src/net/sock_posix.go:198
		// _ = "end of CoverTab[7855]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:199
		_go_fuzz_dep_.CoverTab[7856]++
//line /usr/local/go/src/net/sock_posix.go:199
		// _ = "end of CoverTab[7856]"
//line /usr/local/go/src/net/sock_posix.go:199
	}
//line /usr/local/go/src/net/sock_posix.go:199
	// _ = "end of CoverTab[7840]"
//line /usr/local/go/src/net/sock_posix.go:199
	_go_fuzz_dep_.CoverTab[7841]++
						if err = listenFunc(fd.pfd.Sysfd, backlog); err != nil {
//line /usr/local/go/src/net/sock_posix.go:200
		_go_fuzz_dep_.CoverTab[7857]++
							return os.NewSyscallError("listen", err)
//line /usr/local/go/src/net/sock_posix.go:201
		// _ = "end of CoverTab[7857]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:202
		_go_fuzz_dep_.CoverTab[7858]++
//line /usr/local/go/src/net/sock_posix.go:202
		// _ = "end of CoverTab[7858]"
//line /usr/local/go/src/net/sock_posix.go:202
	}
//line /usr/local/go/src/net/sock_posix.go:202
	// _ = "end of CoverTab[7841]"
//line /usr/local/go/src/net/sock_posix.go:202
	_go_fuzz_dep_.CoverTab[7842]++
						if err = fd.init(); err != nil {
//line /usr/local/go/src/net/sock_posix.go:203
		_go_fuzz_dep_.CoverTab[7859]++
							return err
//line /usr/local/go/src/net/sock_posix.go:204
		// _ = "end of CoverTab[7859]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:205
		_go_fuzz_dep_.CoverTab[7860]++
//line /usr/local/go/src/net/sock_posix.go:205
		// _ = "end of CoverTab[7860]"
//line /usr/local/go/src/net/sock_posix.go:205
	}
//line /usr/local/go/src/net/sock_posix.go:205
	// _ = "end of CoverTab[7842]"
//line /usr/local/go/src/net/sock_posix.go:205
	_go_fuzz_dep_.CoverTab[7843]++
						lsa, _ = syscall.Getsockname(fd.pfd.Sysfd)
						fd.setAddr(fd.addrFunc()(lsa), nil)
						return nil
//line /usr/local/go/src/net/sock_posix.go:208
	// _ = "end of CoverTab[7843]"
}

func (fd *netFD) listenDatagram(ctx context.Context, laddr sockaddr, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) error {
//line /usr/local/go/src/net/sock_posix.go:211
	_go_fuzz_dep_.CoverTab[7861]++
						switch addr := laddr.(type) {
	case *UDPAddr:
//line /usr/local/go/src/net/sock_posix.go:213
		_go_fuzz_dep_.CoverTab[7867]++

//line /usr/local/go/src/net/sock_posix.go:221
		if addr.IP != nil && func() bool {
//line /usr/local/go/src/net/sock_posix.go:221
			_go_fuzz_dep_.CoverTab[7868]++
//line /usr/local/go/src/net/sock_posix.go:221
			return addr.IP.IsMulticast()
//line /usr/local/go/src/net/sock_posix.go:221
			// _ = "end of CoverTab[7868]"
//line /usr/local/go/src/net/sock_posix.go:221
		}() {
//line /usr/local/go/src/net/sock_posix.go:221
			_go_fuzz_dep_.CoverTab[7869]++
								if err := setDefaultMulticastSockopts(fd.pfd.Sysfd); err != nil {
//line /usr/local/go/src/net/sock_posix.go:222
				_go_fuzz_dep_.CoverTab[7872]++
									return err
//line /usr/local/go/src/net/sock_posix.go:223
				// _ = "end of CoverTab[7872]"
			} else {
//line /usr/local/go/src/net/sock_posix.go:224
				_go_fuzz_dep_.CoverTab[7873]++
//line /usr/local/go/src/net/sock_posix.go:224
				// _ = "end of CoverTab[7873]"
//line /usr/local/go/src/net/sock_posix.go:224
			}
//line /usr/local/go/src/net/sock_posix.go:224
			// _ = "end of CoverTab[7869]"
//line /usr/local/go/src/net/sock_posix.go:224
			_go_fuzz_dep_.CoverTab[7870]++
								addr := *addr
								switch fd.family {
			case syscall.AF_INET:
//line /usr/local/go/src/net/sock_posix.go:227
				_go_fuzz_dep_.CoverTab[7874]++
									addr.IP = IPv4zero
//line /usr/local/go/src/net/sock_posix.go:228
				// _ = "end of CoverTab[7874]"
			case syscall.AF_INET6:
//line /usr/local/go/src/net/sock_posix.go:229
				_go_fuzz_dep_.CoverTab[7875]++
									addr.IP = IPv6unspecified
//line /usr/local/go/src/net/sock_posix.go:230
				// _ = "end of CoverTab[7875]"
//line /usr/local/go/src/net/sock_posix.go:230
			default:
//line /usr/local/go/src/net/sock_posix.go:230
				_go_fuzz_dep_.CoverTab[7876]++
//line /usr/local/go/src/net/sock_posix.go:230
				// _ = "end of CoverTab[7876]"
			}
//line /usr/local/go/src/net/sock_posix.go:231
			// _ = "end of CoverTab[7870]"
//line /usr/local/go/src/net/sock_posix.go:231
			_go_fuzz_dep_.CoverTab[7871]++
								laddr = &addr
//line /usr/local/go/src/net/sock_posix.go:232
			// _ = "end of CoverTab[7871]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:233
			_go_fuzz_dep_.CoverTab[7877]++
//line /usr/local/go/src/net/sock_posix.go:233
			// _ = "end of CoverTab[7877]"
//line /usr/local/go/src/net/sock_posix.go:233
		}
//line /usr/local/go/src/net/sock_posix.go:233
		// _ = "end of CoverTab[7867]"
	}
//line /usr/local/go/src/net/sock_posix.go:234
	// _ = "end of CoverTab[7861]"
//line /usr/local/go/src/net/sock_posix.go:234
	_go_fuzz_dep_.CoverTab[7862]++
						var err error
						var lsa syscall.Sockaddr
						if lsa, err = laddr.sockaddr(fd.family); err != nil {
//line /usr/local/go/src/net/sock_posix.go:237
		_go_fuzz_dep_.CoverTab[7878]++
							return err
//line /usr/local/go/src/net/sock_posix.go:238
		// _ = "end of CoverTab[7878]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:239
		_go_fuzz_dep_.CoverTab[7879]++
//line /usr/local/go/src/net/sock_posix.go:239
		// _ = "end of CoverTab[7879]"
//line /usr/local/go/src/net/sock_posix.go:239
	}
//line /usr/local/go/src/net/sock_posix.go:239
	// _ = "end of CoverTab[7862]"
//line /usr/local/go/src/net/sock_posix.go:239
	_go_fuzz_dep_.CoverTab[7863]++

						if ctrlCtxFn != nil {
//line /usr/local/go/src/net/sock_posix.go:241
		_go_fuzz_dep_.CoverTab[7880]++
							c, err := newRawConn(fd)
							if err != nil {
//line /usr/local/go/src/net/sock_posix.go:243
			_go_fuzz_dep_.CoverTab[7882]++
								return err
//line /usr/local/go/src/net/sock_posix.go:244
			// _ = "end of CoverTab[7882]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:245
			_go_fuzz_dep_.CoverTab[7883]++
//line /usr/local/go/src/net/sock_posix.go:245
			// _ = "end of CoverTab[7883]"
//line /usr/local/go/src/net/sock_posix.go:245
		}
//line /usr/local/go/src/net/sock_posix.go:245
		// _ = "end of CoverTab[7880]"
//line /usr/local/go/src/net/sock_posix.go:245
		_go_fuzz_dep_.CoverTab[7881]++
							if err := ctrlCtxFn(ctx, fd.ctrlNetwork(), laddr.String(), c); err != nil {
//line /usr/local/go/src/net/sock_posix.go:246
			_go_fuzz_dep_.CoverTab[7884]++
								return err
//line /usr/local/go/src/net/sock_posix.go:247
			// _ = "end of CoverTab[7884]"
		} else {
//line /usr/local/go/src/net/sock_posix.go:248
			_go_fuzz_dep_.CoverTab[7885]++
//line /usr/local/go/src/net/sock_posix.go:248
			// _ = "end of CoverTab[7885]"
//line /usr/local/go/src/net/sock_posix.go:248
		}
//line /usr/local/go/src/net/sock_posix.go:248
		// _ = "end of CoverTab[7881]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:249
		_go_fuzz_dep_.CoverTab[7886]++
//line /usr/local/go/src/net/sock_posix.go:249
		// _ = "end of CoverTab[7886]"
//line /usr/local/go/src/net/sock_posix.go:249
	}
//line /usr/local/go/src/net/sock_posix.go:249
	// _ = "end of CoverTab[7863]"
//line /usr/local/go/src/net/sock_posix.go:249
	_go_fuzz_dep_.CoverTab[7864]++
						if err = syscall.Bind(fd.pfd.Sysfd, lsa); err != nil {
//line /usr/local/go/src/net/sock_posix.go:250
		_go_fuzz_dep_.CoverTab[7887]++
							return os.NewSyscallError("bind", err)
//line /usr/local/go/src/net/sock_posix.go:251
		// _ = "end of CoverTab[7887]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:252
		_go_fuzz_dep_.CoverTab[7888]++
//line /usr/local/go/src/net/sock_posix.go:252
		// _ = "end of CoverTab[7888]"
//line /usr/local/go/src/net/sock_posix.go:252
	}
//line /usr/local/go/src/net/sock_posix.go:252
	// _ = "end of CoverTab[7864]"
//line /usr/local/go/src/net/sock_posix.go:252
	_go_fuzz_dep_.CoverTab[7865]++
						if err = fd.init(); err != nil {
//line /usr/local/go/src/net/sock_posix.go:253
		_go_fuzz_dep_.CoverTab[7889]++
							return err
//line /usr/local/go/src/net/sock_posix.go:254
		// _ = "end of CoverTab[7889]"
	} else {
//line /usr/local/go/src/net/sock_posix.go:255
		_go_fuzz_dep_.CoverTab[7890]++
//line /usr/local/go/src/net/sock_posix.go:255
		// _ = "end of CoverTab[7890]"
//line /usr/local/go/src/net/sock_posix.go:255
	}
//line /usr/local/go/src/net/sock_posix.go:255
	// _ = "end of CoverTab[7865]"
//line /usr/local/go/src/net/sock_posix.go:255
	_go_fuzz_dep_.CoverTab[7866]++
						lsa, _ = syscall.Getsockname(fd.pfd.Sysfd)
						fd.setAddr(fd.addrFunc()(lsa), nil)
						return nil
//line /usr/local/go/src/net/sock_posix.go:258
	// _ = "end of CoverTab[7866]"
}

//line /usr/local/go/src/net/sock_posix.go:259
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sock_posix.go:259
var _ = _go_fuzz_dep_.CoverTab
