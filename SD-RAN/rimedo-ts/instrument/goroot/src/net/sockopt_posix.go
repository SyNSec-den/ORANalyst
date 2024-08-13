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
	_go_fuzz_dep_.CoverTab[16293]++
							if b {
//line /usr/local/go/src/net/sockopt_posix.go:17
		_go_fuzz_dep_.CoverTab[16295]++
								return 1
//line /usr/local/go/src/net/sockopt_posix.go:18
		// _ = "end of CoverTab[16295]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:19
		_go_fuzz_dep_.CoverTab[16296]++
//line /usr/local/go/src/net/sockopt_posix.go:19
		// _ = "end of CoverTab[16296]"
//line /usr/local/go/src/net/sockopt_posix.go:19
	}
//line /usr/local/go/src/net/sockopt_posix.go:19
	// _ = "end of CoverTab[16293]"
//line /usr/local/go/src/net/sockopt_posix.go:19
	_go_fuzz_dep_.CoverTab[16294]++
							return 0
//line /usr/local/go/src/net/sockopt_posix.go:20
	// _ = "end of CoverTab[16294]"
}

func ipv4AddrToInterface(ip IP) (*Interface, error) {
//line /usr/local/go/src/net/sockopt_posix.go:23
	_go_fuzz_dep_.CoverTab[16297]++
							ift, err := Interfaces()
							if err != nil {
//line /usr/local/go/src/net/sockopt_posix.go:25
		_go_fuzz_dep_.CoverTab[16301]++
								return nil, err
//line /usr/local/go/src/net/sockopt_posix.go:26
		// _ = "end of CoverTab[16301]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:27
		_go_fuzz_dep_.CoverTab[16302]++
//line /usr/local/go/src/net/sockopt_posix.go:27
		// _ = "end of CoverTab[16302]"
//line /usr/local/go/src/net/sockopt_posix.go:27
	}
//line /usr/local/go/src/net/sockopt_posix.go:27
	// _ = "end of CoverTab[16297]"
//line /usr/local/go/src/net/sockopt_posix.go:27
	_go_fuzz_dep_.CoverTab[16298]++
							for _, ifi := range ift {
//line /usr/local/go/src/net/sockopt_posix.go:28
		_go_fuzz_dep_.CoverTab[16303]++
								ifat, err := ifi.Addrs()
								if err != nil {
//line /usr/local/go/src/net/sockopt_posix.go:30
			_go_fuzz_dep_.CoverTab[16305]++
									return nil, err
//line /usr/local/go/src/net/sockopt_posix.go:31
			// _ = "end of CoverTab[16305]"
		} else {
//line /usr/local/go/src/net/sockopt_posix.go:32
			_go_fuzz_dep_.CoverTab[16306]++
//line /usr/local/go/src/net/sockopt_posix.go:32
			// _ = "end of CoverTab[16306]"
//line /usr/local/go/src/net/sockopt_posix.go:32
		}
//line /usr/local/go/src/net/sockopt_posix.go:32
		// _ = "end of CoverTab[16303]"
//line /usr/local/go/src/net/sockopt_posix.go:32
		_go_fuzz_dep_.CoverTab[16304]++
								for _, ifa := range ifat {
//line /usr/local/go/src/net/sockopt_posix.go:33
			_go_fuzz_dep_.CoverTab[16307]++
									switch v := ifa.(type) {
			case *IPAddr:
//line /usr/local/go/src/net/sockopt_posix.go:35
				_go_fuzz_dep_.CoverTab[16308]++
										if ip.Equal(v.IP) {
//line /usr/local/go/src/net/sockopt_posix.go:36
					_go_fuzz_dep_.CoverTab[16310]++
											return &ifi, nil
//line /usr/local/go/src/net/sockopt_posix.go:37
					// _ = "end of CoverTab[16310]"
				} else {
//line /usr/local/go/src/net/sockopt_posix.go:38
					_go_fuzz_dep_.CoverTab[16311]++
//line /usr/local/go/src/net/sockopt_posix.go:38
					// _ = "end of CoverTab[16311]"
//line /usr/local/go/src/net/sockopt_posix.go:38
				}
//line /usr/local/go/src/net/sockopt_posix.go:38
				// _ = "end of CoverTab[16308]"
			case *IPNet:
//line /usr/local/go/src/net/sockopt_posix.go:39
				_go_fuzz_dep_.CoverTab[16309]++
										if ip.Equal(v.IP) {
//line /usr/local/go/src/net/sockopt_posix.go:40
					_go_fuzz_dep_.CoverTab[16312]++
											return &ifi, nil
//line /usr/local/go/src/net/sockopt_posix.go:41
					// _ = "end of CoverTab[16312]"
				} else {
//line /usr/local/go/src/net/sockopt_posix.go:42
					_go_fuzz_dep_.CoverTab[16313]++
//line /usr/local/go/src/net/sockopt_posix.go:42
					// _ = "end of CoverTab[16313]"
//line /usr/local/go/src/net/sockopt_posix.go:42
				}
//line /usr/local/go/src/net/sockopt_posix.go:42
				// _ = "end of CoverTab[16309]"
			}
//line /usr/local/go/src/net/sockopt_posix.go:43
			// _ = "end of CoverTab[16307]"
		}
//line /usr/local/go/src/net/sockopt_posix.go:44
		// _ = "end of CoverTab[16304]"
	}
//line /usr/local/go/src/net/sockopt_posix.go:45
	// _ = "end of CoverTab[16298]"
//line /usr/local/go/src/net/sockopt_posix.go:45
	_go_fuzz_dep_.CoverTab[16299]++
							if ip.Equal(IPv4zero) {
//line /usr/local/go/src/net/sockopt_posix.go:46
		_go_fuzz_dep_.CoverTab[16314]++
								return nil, nil
//line /usr/local/go/src/net/sockopt_posix.go:47
		// _ = "end of CoverTab[16314]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:48
		_go_fuzz_dep_.CoverTab[16315]++
//line /usr/local/go/src/net/sockopt_posix.go:48
		// _ = "end of CoverTab[16315]"
//line /usr/local/go/src/net/sockopt_posix.go:48
	}
//line /usr/local/go/src/net/sockopt_posix.go:48
	// _ = "end of CoverTab[16299]"
//line /usr/local/go/src/net/sockopt_posix.go:48
	_go_fuzz_dep_.CoverTab[16300]++
							return nil, errNoSuchInterface
//line /usr/local/go/src/net/sockopt_posix.go:49
	// _ = "end of CoverTab[16300]"
}

func interfaceToIPv4Addr(ifi *Interface) (IP, error) {
//line /usr/local/go/src/net/sockopt_posix.go:52
	_go_fuzz_dep_.CoverTab[16316]++
							if ifi == nil {
//line /usr/local/go/src/net/sockopt_posix.go:53
		_go_fuzz_dep_.CoverTab[16320]++
								return IPv4zero, nil
//line /usr/local/go/src/net/sockopt_posix.go:54
		// _ = "end of CoverTab[16320]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:55
		_go_fuzz_dep_.CoverTab[16321]++
//line /usr/local/go/src/net/sockopt_posix.go:55
		// _ = "end of CoverTab[16321]"
//line /usr/local/go/src/net/sockopt_posix.go:55
	}
//line /usr/local/go/src/net/sockopt_posix.go:55
	// _ = "end of CoverTab[16316]"
//line /usr/local/go/src/net/sockopt_posix.go:55
	_go_fuzz_dep_.CoverTab[16317]++
							ifat, err := ifi.Addrs()
							if err != nil {
//line /usr/local/go/src/net/sockopt_posix.go:57
		_go_fuzz_dep_.CoverTab[16322]++
								return nil, err
//line /usr/local/go/src/net/sockopt_posix.go:58
		// _ = "end of CoverTab[16322]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:59
		_go_fuzz_dep_.CoverTab[16323]++
//line /usr/local/go/src/net/sockopt_posix.go:59
		// _ = "end of CoverTab[16323]"
//line /usr/local/go/src/net/sockopt_posix.go:59
	}
//line /usr/local/go/src/net/sockopt_posix.go:59
	// _ = "end of CoverTab[16317]"
//line /usr/local/go/src/net/sockopt_posix.go:59
	_go_fuzz_dep_.CoverTab[16318]++
							for _, ifa := range ifat {
//line /usr/local/go/src/net/sockopt_posix.go:60
		_go_fuzz_dep_.CoverTab[16324]++
								switch v := ifa.(type) {
		case *IPAddr:
//line /usr/local/go/src/net/sockopt_posix.go:62
			_go_fuzz_dep_.CoverTab[16325]++
									if v.IP.To4() != nil {
//line /usr/local/go/src/net/sockopt_posix.go:63
				_go_fuzz_dep_.CoverTab[16327]++
										return v.IP, nil
//line /usr/local/go/src/net/sockopt_posix.go:64
				// _ = "end of CoverTab[16327]"
			} else {
//line /usr/local/go/src/net/sockopt_posix.go:65
				_go_fuzz_dep_.CoverTab[16328]++
//line /usr/local/go/src/net/sockopt_posix.go:65
				// _ = "end of CoverTab[16328]"
//line /usr/local/go/src/net/sockopt_posix.go:65
			}
//line /usr/local/go/src/net/sockopt_posix.go:65
			// _ = "end of CoverTab[16325]"
		case *IPNet:
//line /usr/local/go/src/net/sockopt_posix.go:66
			_go_fuzz_dep_.CoverTab[16326]++
									if v.IP.To4() != nil {
//line /usr/local/go/src/net/sockopt_posix.go:67
				_go_fuzz_dep_.CoverTab[16329]++
										return v.IP, nil
//line /usr/local/go/src/net/sockopt_posix.go:68
				// _ = "end of CoverTab[16329]"
			} else {
//line /usr/local/go/src/net/sockopt_posix.go:69
				_go_fuzz_dep_.CoverTab[16330]++
//line /usr/local/go/src/net/sockopt_posix.go:69
				// _ = "end of CoverTab[16330]"
//line /usr/local/go/src/net/sockopt_posix.go:69
			}
//line /usr/local/go/src/net/sockopt_posix.go:69
			// _ = "end of CoverTab[16326]"
		}
//line /usr/local/go/src/net/sockopt_posix.go:70
		// _ = "end of CoverTab[16324]"
	}
//line /usr/local/go/src/net/sockopt_posix.go:71
	// _ = "end of CoverTab[16318]"
//line /usr/local/go/src/net/sockopt_posix.go:71
	_go_fuzz_dep_.CoverTab[16319]++
							return nil, errNoSuchInterface
//line /usr/local/go/src/net/sockopt_posix.go:72
	// _ = "end of CoverTab[16319]"
}

func setIPv4MreqToInterface(mreq *syscall.IPMreq, ifi *Interface) error {
//line /usr/local/go/src/net/sockopt_posix.go:75
	_go_fuzz_dep_.CoverTab[16331]++
							if ifi == nil {
//line /usr/local/go/src/net/sockopt_posix.go:76
		_go_fuzz_dep_.CoverTab[16336]++
								return nil
//line /usr/local/go/src/net/sockopt_posix.go:77
		// _ = "end of CoverTab[16336]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:78
		_go_fuzz_dep_.CoverTab[16337]++
//line /usr/local/go/src/net/sockopt_posix.go:78
		// _ = "end of CoverTab[16337]"
//line /usr/local/go/src/net/sockopt_posix.go:78
	}
//line /usr/local/go/src/net/sockopt_posix.go:78
	// _ = "end of CoverTab[16331]"
//line /usr/local/go/src/net/sockopt_posix.go:78
	_go_fuzz_dep_.CoverTab[16332]++
							ifat, err := ifi.Addrs()
							if err != nil {
//line /usr/local/go/src/net/sockopt_posix.go:80
		_go_fuzz_dep_.CoverTab[16338]++
								return err
//line /usr/local/go/src/net/sockopt_posix.go:81
		// _ = "end of CoverTab[16338]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:82
		_go_fuzz_dep_.CoverTab[16339]++
//line /usr/local/go/src/net/sockopt_posix.go:82
		// _ = "end of CoverTab[16339]"
//line /usr/local/go/src/net/sockopt_posix.go:82
	}
//line /usr/local/go/src/net/sockopt_posix.go:82
	// _ = "end of CoverTab[16332]"
//line /usr/local/go/src/net/sockopt_posix.go:82
	_go_fuzz_dep_.CoverTab[16333]++
							for _, ifa := range ifat {
//line /usr/local/go/src/net/sockopt_posix.go:83
		_go_fuzz_dep_.CoverTab[16340]++
								switch v := ifa.(type) {
		case *IPAddr:
//line /usr/local/go/src/net/sockopt_posix.go:85
			_go_fuzz_dep_.CoverTab[16341]++
									if a := v.IP.To4(); a != nil {
//line /usr/local/go/src/net/sockopt_posix.go:86
				_go_fuzz_dep_.CoverTab[16343]++
										copy(mreq.Interface[:], a)
										goto done
//line /usr/local/go/src/net/sockopt_posix.go:88
				// _ = "end of CoverTab[16343]"
			} else {
//line /usr/local/go/src/net/sockopt_posix.go:89
				_go_fuzz_dep_.CoverTab[16344]++
//line /usr/local/go/src/net/sockopt_posix.go:89
				// _ = "end of CoverTab[16344]"
//line /usr/local/go/src/net/sockopt_posix.go:89
			}
//line /usr/local/go/src/net/sockopt_posix.go:89
			// _ = "end of CoverTab[16341]"
		case *IPNet:
//line /usr/local/go/src/net/sockopt_posix.go:90
			_go_fuzz_dep_.CoverTab[16342]++
									if a := v.IP.To4(); a != nil {
//line /usr/local/go/src/net/sockopt_posix.go:91
				_go_fuzz_dep_.CoverTab[16345]++
										copy(mreq.Interface[:], a)
										goto done
//line /usr/local/go/src/net/sockopt_posix.go:93
				// _ = "end of CoverTab[16345]"
			} else {
//line /usr/local/go/src/net/sockopt_posix.go:94
				_go_fuzz_dep_.CoverTab[16346]++
//line /usr/local/go/src/net/sockopt_posix.go:94
				// _ = "end of CoverTab[16346]"
//line /usr/local/go/src/net/sockopt_posix.go:94
			}
//line /usr/local/go/src/net/sockopt_posix.go:94
			// _ = "end of CoverTab[16342]"
		}
//line /usr/local/go/src/net/sockopt_posix.go:95
		// _ = "end of CoverTab[16340]"
	}
//line /usr/local/go/src/net/sockopt_posix.go:96
	// _ = "end of CoverTab[16333]"
//line /usr/local/go/src/net/sockopt_posix.go:96
	_go_fuzz_dep_.CoverTab[16334]++
done:
	if bytealg.Equal(mreq.Multiaddr[:], IPv4zero.To4()) {
//line /usr/local/go/src/net/sockopt_posix.go:98
		_go_fuzz_dep_.CoverTab[16347]++
								return errNoSuchMulticastInterface
//line /usr/local/go/src/net/sockopt_posix.go:99
		// _ = "end of CoverTab[16347]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:100
		_go_fuzz_dep_.CoverTab[16348]++
//line /usr/local/go/src/net/sockopt_posix.go:100
		// _ = "end of CoverTab[16348]"
//line /usr/local/go/src/net/sockopt_posix.go:100
	}
//line /usr/local/go/src/net/sockopt_posix.go:100
	// _ = "end of CoverTab[16334]"
//line /usr/local/go/src/net/sockopt_posix.go:100
	_go_fuzz_dep_.CoverTab[16335]++
							return nil
//line /usr/local/go/src/net/sockopt_posix.go:101
	// _ = "end of CoverTab[16335]"
}

func setReadBuffer(fd *netFD, bytes int) error {
//line /usr/local/go/src/net/sockopt_posix.go:104
	_go_fuzz_dep_.CoverTab[16349]++
							err := fd.pfd.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_RCVBUF, bytes)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockopt_posix.go:107
	// _ = "end of CoverTab[16349]"
}

func setWriteBuffer(fd *netFD, bytes int) error {
//line /usr/local/go/src/net/sockopt_posix.go:110
	_go_fuzz_dep_.CoverTab[16350]++
							err := fd.pfd.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_SNDBUF, bytes)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockopt_posix.go:113
	// _ = "end of CoverTab[16350]"
}

func setKeepAlive(fd *netFD, keepalive bool) error {
//line /usr/local/go/src/net/sockopt_posix.go:116
	_go_fuzz_dep_.CoverTab[16351]++
							err := fd.pfd.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_KEEPALIVE, boolint(keepalive))
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockopt_posix.go:119
	// _ = "end of CoverTab[16351]"
}

func setLinger(fd *netFD, sec int) error {
//line /usr/local/go/src/net/sockopt_posix.go:122
	_go_fuzz_dep_.CoverTab[16352]++
							var l syscall.Linger
							if sec >= 0 {
//line /usr/local/go/src/net/sockopt_posix.go:124
		_go_fuzz_dep_.CoverTab[16354]++
								l.Onoff = 1
								l.Linger = int32(sec)
//line /usr/local/go/src/net/sockopt_posix.go:126
		// _ = "end of CoverTab[16354]"
	} else {
//line /usr/local/go/src/net/sockopt_posix.go:127
		_go_fuzz_dep_.CoverTab[16355]++
								l.Onoff = 0
								l.Linger = 0
//line /usr/local/go/src/net/sockopt_posix.go:129
		// _ = "end of CoverTab[16355]"
	}
//line /usr/local/go/src/net/sockopt_posix.go:130
	// _ = "end of CoverTab[16352]"
//line /usr/local/go/src/net/sockopt_posix.go:130
	_go_fuzz_dep_.CoverTab[16353]++
							err := fd.pfd.SetsockoptLinger(syscall.SOL_SOCKET, syscall.SO_LINGER, &l)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /usr/local/go/src/net/sockopt_posix.go:133
	// _ = "end of CoverTab[16353]"
}

//line /usr/local/go/src/net/sockopt_posix.go:134
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/sockopt_posix.go:134
var _ = _go_fuzz_dep_.CoverTab
