// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

//line /usr/local/go/src/net/sockopt_posix.go:7
package net

//line /usr/local/go/src/net/sockopt_posix.go:7
import (
//line /usr/local/go/src/net/sockopt_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/sockopt_posix.go:7
)
//line /usr/local/go/src/net/sockopt_posix.go:7
import (
//line /usr/local/go/src/net/sockopt_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/sockopt_posix.go:7
)

import (
	"internal/bytealg"
	"runtime"
	"syscall"
)

// Boolean to int.
func boolint(b bool) int {
//line /usr/local/go/src/net/sockopt_posix.go:16
	_go_fuzz_dep_.CoverTab[7903]++
							if b {
//line /usr/local/go/src/net/sockopt_posix.go:17
		_go_fuzz_dep_.CoverTab[7905]++
								return 1
//line /usr/local/go/src/net/sockopt_posix.go:18
		// _ = "end of CoverTab[7905]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:19
		_go_fuzz_dep_.CoverTab[7906]++
//line /usr/local/go/src/net/sockopt_posix.go:19
		// _ = "end of CoverTab[7906]"
//line /usr/local/go/src/net/sockopt_posix.go:19
	}
//line /usr/local/go/src/net/sockopt_posix.go:19
	// _ = "end of CoverTab[7903]"
//line /usr/local/go/src/net/sockopt_posix.go:19
	_go_fuzz_dep_.CoverTab[7904]++
							return 0
//line /usr/local/go/src/net/sockopt_posix.go:20
	// _ = "end of CoverTab[7904]"
}

func ipv4AddrToInterface(ip IP) (*Interface, error) {
//line /usr/local/go/src/net/sockopt_posix.go:23
	_go_fuzz_dep_.CoverTab[7907]++
							ift, err := Interfaces()
							if err != nil {
//line /usr/local/go/src/net/sockopt_posix.go:25
		_go_fuzz_dep_.CoverTab[7911]++
								return nil, err
//line /usr/local/go/src/net/sockopt_posix.go:26
		// _ = "end of CoverTab[7911]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:27
		_go_fuzz_dep_.CoverTab[7912]++
//line /usr/local/go/src/net/sockopt_posix.go:27
		// _ = "end of CoverTab[7912]"
//line /usr/local/go/src/net/sockopt_posix.go:27
	}
//line /usr/local/go/src/net/sockopt_posix.go:27
	// _ = "end of CoverTab[7907]"
//line /usr/local/go/src/net/sockopt_posix.go:27
	_go_fuzz_dep_.CoverTab[7908]++
							for _, ifi := range ift {
//line /usr/local/go/src/net/sockopt_posix.go:28
		_go_fuzz_dep_.CoverTab[7913]++
								ifat, err := ifi.Addrs()
								if err != nil {
//line /usr/local/go/src/net/sockopt_posix.go:30
			_go_fuzz_dep_.CoverTab[7915]++
									return nil, err
//line /usr/local/go/src/net/sockopt_posix.go:31
			// _ = "end of CoverTab[7915]"
		} else {
//line /usr/local/go/src/net/sockopt_posix.go:32
			_go_fuzz_dep_.CoverTab[7916]++
//line /usr/local/go/src/net/sockopt_posix.go:32
			// _ = "end of CoverTab[7916]"
//line /usr/local/go/src/net/sockopt_posix.go:32
		}
//line /usr/local/go/src/net/sockopt_posix.go:32
		// _ = "end of CoverTab[7913]"
//line /usr/local/go/src/net/sockopt_posix.go:32
		_go_fuzz_dep_.CoverTab[7914]++
								for _, ifa := range ifat {
//line /usr/local/go/src/net/sockopt_posix.go:33
			_go_fuzz_dep_.CoverTab[7917]++
									switch v := ifa.(type) {
			case *IPAddr:
//line /usr/local/go/src/net/sockopt_posix.go:35
				_go_fuzz_dep_.CoverTab[7918]++
										if ip.Equal(v.IP) {
//line /usr/local/go/src/net/sockopt_posix.go:36
					_go_fuzz_dep_.CoverTab[7920]++
											return &ifi, nil
//line /usr/local/go/src/net/sockopt_posix.go:37
					// _ = "end of CoverTab[7920]"
				} else {
//line /usr/local/go/src/net/sockopt_posix.go:38
					_go_fuzz_dep_.CoverTab[7921]++
//line /usr/local/go/src/net/sockopt_posix.go:38
					// _ = "end of CoverTab[7921]"
//line /usr/local/go/src/net/sockopt_posix.go:38
				}
//line /usr/local/go/src/net/sockopt_posix.go:38
				// _ = "end of CoverTab[7918]"
			case *IPNet:
//line /usr/local/go/src/net/sockopt_posix.go:39
				_go_fuzz_dep_.CoverTab[7919]++
										if ip.Equal(v.IP) {
//line /usr/local/go/src/net/sockopt_posix.go:40
					_go_fuzz_dep_.CoverTab[7922]++
											return &ifi, nil
//line /usr/local/go/src/net/sockopt_posix.go:41
					// _ = "end of CoverTab[7922]"
				} else {
//line /usr/local/go/src/net/sockopt_posix.go:42
					_go_fuzz_dep_.CoverTab[7923]++
//line /usr/local/go/src/net/sockopt_posix.go:42
					// _ = "end of CoverTab[7923]"
//line /usr/local/go/src/net/sockopt_posix.go:42
				}
//line /usr/local/go/src/net/sockopt_posix.go:42
				// _ = "end of CoverTab[7919]"
			}
//line /usr/local/go/src/net/sockopt_posix.go:43
			// _ = "end of CoverTab[7917]"
		}
//line /usr/local/go/src/net/sockopt_posix.go:44
		// _ = "end of CoverTab[7914]"
	}
//line /usr/local/go/src/net/sockopt_posix.go:45
	// _ = "end of CoverTab[7908]"
//line /usr/local/go/src/net/sockopt_posix.go:45
	_go_fuzz_dep_.CoverTab[7909]++
							if ip.Equal(IPv4zero) {
//line /usr/local/go/src/net/sockopt_posix.go:46
		_go_fuzz_dep_.CoverTab[7924]++
								return nil, nil
//line /usr/local/go/src/net/sockopt_posix.go:47
		// _ = "end of CoverTab[7924]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:48
		_go_fuzz_dep_.CoverTab[7925]++
//line /usr/local/go/src/net/sockopt_posix.go:48
		// _ = "end of CoverTab[7925]"
//line /usr/local/go/src/net/sockopt_posix.go:48
	}
//line /usr/local/go/src/net/sockopt_posix.go:48
	// _ = "end of CoverTab[7909]"
//line /usr/local/go/src/net/sockopt_posix.go:48
	_go_fuzz_dep_.CoverTab[7910]++
							return nil, errNoSuchInterface
//line /usr/local/go/src/net/sockopt_posix.go:49
	// _ = "end of CoverTab[7910]"
}

func interfaceToIPv4Addr(ifi *Interface) (IP, error) {
//line /usr/local/go/src/net/sockopt_posix.go:52
	_go_fuzz_dep_.CoverTab[7926]++
							if ifi == nil {
//line /usr/local/go/src/net/sockopt_posix.go:53
		_go_fuzz_dep_.CoverTab[7930]++
								return IPv4zero, nil
//line /usr/local/go/src/net/sockopt_posix.go:54
		// _ = "end of CoverTab[7930]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:55
		_go_fuzz_dep_.CoverTab[7931]++
//line /usr/local/go/src/net/sockopt_posix.go:55
		// _ = "end of CoverTab[7931]"
//line /usr/local/go/src/net/sockopt_posix.go:55
	}
//line /usr/local/go/src/net/sockopt_posix.go:55
	// _ = "end of CoverTab[7926]"
//line /usr/local/go/src/net/sockopt_posix.go:55
	_go_fuzz_dep_.CoverTab[7927]++
							ifat, err := ifi.Addrs()
							if err != nil {
//line /usr/local/go/src/net/sockopt_posix.go:57
		_go_fuzz_dep_.CoverTab[7932]++
								return nil, err
//line /usr/local/go/src/net/sockopt_posix.go:58
		// _ = "end of CoverTab[7932]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:59
		_go_fuzz_dep_.CoverTab[7933]++
//line /usr/local/go/src/net/sockopt_posix.go:59
		// _ = "end of CoverTab[7933]"
//line /usr/local/go/src/net/sockopt_posix.go:59
	}
//line /usr/local/go/src/net/sockopt_posix.go:59
	// _ = "end of CoverTab[7927]"
//line /usr/local/go/src/net/sockopt_posix.go:59
	_go_fuzz_dep_.CoverTab[7928]++
							for _, ifa := range ifat {
//line /usr/local/go/src/net/sockopt_posix.go:60
		_go_fuzz_dep_.CoverTab[7934]++
								switch v := ifa.(type) {
		case *IPAddr:
//line /usr/local/go/src/net/sockopt_posix.go:62
			_go_fuzz_dep_.CoverTab[7935]++
									if v.IP.To4() != nil {
//line /usr/local/go/src/net/sockopt_posix.go:63
				_go_fuzz_dep_.CoverTab[7937]++
										return v.IP, nil
//line /usr/local/go/src/net/sockopt_posix.go:64
				// _ = "end of CoverTab[7937]"
			} else {
//line /usr/local/go/src/net/sockopt_posix.go:65
				_go_fuzz_dep_.CoverTab[7938]++
//line /usr/local/go/src/net/sockopt_posix.go:65
				// _ = "end of CoverTab[7938]"
//line /usr/local/go/src/net/sockopt_posix.go:65
			}
//line /usr/local/go/src/net/sockopt_posix.go:65
			// _ = "end of CoverTab[7935]"
		case *IPNet:
//line /usr/local/go/src/net/sockopt_posix.go:66
			_go_fuzz_dep_.CoverTab[7936]++
									if v.IP.To4() != nil {
//line /usr/local/go/src/net/sockopt_posix.go:67
				_go_fuzz_dep_.CoverTab[7939]++
										return v.IP, nil
//line /usr/local/go/src/net/sockopt_posix.go:68
				// _ = "end of CoverTab[7939]"
			} else {
//line /usr/local/go/src/net/sockopt_posix.go:69
				_go_fuzz_dep_.CoverTab[7940]++
//line /usr/local/go/src/net/sockopt_posix.go:69
				// _ = "end of CoverTab[7940]"
//line /usr/local/go/src/net/sockopt_posix.go:69
			}
//line /usr/local/go/src/net/sockopt_posix.go:69
			// _ = "end of CoverTab[7936]"
		}
//line /usr/local/go/src/net/sockopt_posix.go:70
		// _ = "end of CoverTab[7934]"
	}
//line /usr/local/go/src/net/sockopt_posix.go:71
	// _ = "end of CoverTab[7928]"
//line /usr/local/go/src/net/sockopt_posix.go:71
	_go_fuzz_dep_.CoverTab[7929]++
							return nil, errNoSuchInterface
//line /usr/local/go/src/net/sockopt_posix.go:72
	// _ = "end of CoverTab[7929]"
}

func setIPv4MreqToInterface(mreq *syscall.IPMreq, ifi *Interface) error {
//line /usr/local/go/src/net/sockopt_posix.go:75
	_go_fuzz_dep_.CoverTab[7941]++
							if ifi == nil {
//line /usr/local/go/src/net/sockopt_posix.go:76
		_go_fuzz_dep_.CoverTab[7946]++
								return nil
//line /usr/local/go/src/net/sockopt_posix.go:77
		// _ = "end of CoverTab[7946]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:78
		_go_fuzz_dep_.CoverTab[7947]++
//line /usr/local/go/src/net/sockopt_posix.go:78
		// _ = "end of CoverTab[7947]"
//line /usr/local/go/src/net/sockopt_posix.go:78
	}
//line /usr/local/go/src/net/sockopt_posix.go:78
	// _ = "end of CoverTab[7941]"
//line /usr/local/go/src/net/sockopt_posix.go:78
	_go_fuzz_dep_.CoverTab[7942]++
							ifat, err := ifi.Addrs()
							if err != nil {
//line /usr/local/go/src/net/sockopt_posix.go:80
		_go_fuzz_dep_.CoverTab[7948]++
								return err
//line /usr/local/go/src/net/sockopt_posix.go:81
		// _ = "end of CoverTab[7948]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:82
		_go_fuzz_dep_.CoverTab[7949]++
//line /usr/local/go/src/net/sockopt_posix.go:82
		// _ = "end of CoverTab[7949]"
//line /usr/local/go/src/net/sockopt_posix.go:82
	}
//line /usr/local/go/src/net/sockopt_posix.go:82
	// _ = "end of CoverTab[7942]"
//line /usr/local/go/src/net/sockopt_posix.go:82
	_go_fuzz_dep_.CoverTab[7943]++
							for _, ifa := range ifat {
//line /usr/local/go/src/net/sockopt_posix.go:83
		_go_fuzz_dep_.CoverTab[7950]++
								switch v := ifa.(type) {
		case *IPAddr:
//line /usr/local/go/src/net/sockopt_posix.go:85
			_go_fuzz_dep_.CoverTab[7951]++
									if a := v.IP.To4(); a != nil {
//line /usr/local/go/src/net/sockopt_posix.go:86
				_go_fuzz_dep_.CoverTab[7953]++
										copy(mreq.Interface[:], a)
										goto done
//line /usr/local/go/src/net/sockopt_posix.go:88
				// _ = "end of CoverTab[7953]"
			} else {
//line /usr/local/go/src/net/sockopt_posix.go:89
				_go_fuzz_dep_.CoverTab[7954]++
//line /usr/local/go/src/net/sockopt_posix.go:89
				// _ = "end of CoverTab[7954]"
//line /usr/local/go/src/net/sockopt_posix.go:89
			}
//line /usr/local/go/src/net/sockopt_posix.go:89
			// _ = "end of CoverTab[7951]"
		case *IPNet:
//line /usr/local/go/src/net/sockopt_posix.go:90
			_go_fuzz_dep_.CoverTab[7952]++
									if a := v.IP.To4(); a != nil {
//line /usr/local/go/src/net/sockopt_posix.go:91
				_go_fuzz_dep_.CoverTab[7955]++
										copy(mreq.Interface[:], a)
										goto done
//line /usr/local/go/src/net/sockopt_posix.go:93
				// _ = "end of CoverTab[7955]"
			} else {
//line /usr/local/go/src/net/sockopt_posix.go:94
				_go_fuzz_dep_.CoverTab[7956]++
//line /usr/local/go/src/net/sockopt_posix.go:94
				// _ = "end of CoverTab[7956]"
//line /usr/local/go/src/net/sockopt_posix.go:94
			}
//line /usr/local/go/src/net/sockopt_posix.go:94
			// _ = "end of CoverTab[7952]"
		}
//line /usr/local/go/src/net/sockopt_posix.go:95
		// _ = "end of CoverTab[7950]"
	}
//line /usr/local/go/src/net/sockopt_posix.go:96
	// _ = "end of CoverTab[7943]"
//line /usr/local/go/src/net/sockopt_posix.go:96
	_go_fuzz_dep_.CoverTab[7944]++
done:
	if bytealg.Equal(mreq.Multiaddr[:], IPv4zero.To4()) {
//line /usr/local/go/src/net/sockopt_posix.go:98
		_go_fuzz_dep_.CoverTab[7957]++
								return errNoSuchMulticastInterface
//line /usr/local/go/src/net/sockopt_posix.go:99
		// _ = "end of CoverTab[7957]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:100
		_go_fuzz_dep_.CoverTab[7958]++
//line /usr/local/go/src/net/sockopt_posix.go:100
		// _ = "end of CoverTab[7958]"
//line /usr/local/go/src/net/sockopt_posix.go:100
	}
//line /usr/local/go/src/net/sockopt_posix.go:100
	// _ = "end of CoverTab[7944]"
//line /usr/local/go/src/net/sockopt_posix.go:100
	_go_fuzz_dep_.CoverTab[7945]++
							return nil
//line /usr/local/go/src/net/sockopt_posix.go:101
	// _ = "end of CoverTab[7945]"
}

func setReadBuffer(fd *netFD, bytes int) error {
//line /usr/local/go/src/net/sockopt_posix.go:104
	_go_fuzz_dep_.CoverTab[7959]++
							err := fd.pfd.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_RCVBUF, bytes)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockopt_posix.go:107
	// _ = "end of CoverTab[7959]"
}

func setWriteBuffer(fd *netFD, bytes int) error {
//line /usr/local/go/src/net/sockopt_posix.go:110
	_go_fuzz_dep_.CoverTab[7960]++
							err := fd.pfd.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_SNDBUF, bytes)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockopt_posix.go:113
	// _ = "end of CoverTab[7960]"
}

func setKeepAlive(fd *netFD, keepalive bool) error {
//line /usr/local/go/src/net/sockopt_posix.go:116
	_go_fuzz_dep_.CoverTab[7961]++
							err := fd.pfd.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_KEEPALIVE, boolint(keepalive))
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockopt_posix.go:119
	// _ = "end of CoverTab[7961]"
}

func setLinger(fd *netFD, sec int) error {
//line /usr/local/go/src/net/sockopt_posix.go:122
	_go_fuzz_dep_.CoverTab[7962]++
							var l syscall.Linger
							if sec >= 0 {
//line /usr/local/go/src/net/sockopt_posix.go:124
		_go_fuzz_dep_.CoverTab[7964]++
								l.Onoff = 1
								l.Linger = int32(sec)
//line /usr/local/go/src/net/sockopt_posix.go:126
		// _ = "end of CoverTab[7964]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:127
		_go_fuzz_dep_.CoverTab[7965]++
								l.Onoff = 0
								l.Linger = 0
//line /usr/local/go/src/net/sockopt_posix.go:129
		// _ = "end of CoverTab[7965]"
	}
//line /usr/local/go/src/net/sockopt_posix.go:130
	// _ = "end of CoverTab[7962]"
//line /usr/local/go/src/net/sockopt_posix.go:130
	_go_fuzz_dep_.CoverTab[7963]++
							err := fd.pfd.SetsockoptLinger(syscall.SOL_SOCKET, syscall.SO_LINGER, &l)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockopt_posix.go:133
	// _ = "end of CoverTab[7963]"
}

//line /usr/local/go/src/net/sockopt_posix.go:134
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sockopt_posix.go:134
var _ = _go_fuzz_dep_.CoverTab
