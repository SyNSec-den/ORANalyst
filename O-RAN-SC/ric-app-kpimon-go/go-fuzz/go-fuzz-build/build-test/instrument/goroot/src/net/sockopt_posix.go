// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

//line /snap/go/10455/src/net/sockopt_posix.go:7
package net

//line /snap/go/10455/src/net/sockopt_posix.go:7
import (
//line /snap/go/10455/src/net/sockopt_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/sockopt_posix.go:7
)
//line /snap/go/10455/src/net/sockopt_posix.go:7
import (
//line /snap/go/10455/src/net/sockopt_posix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/sockopt_posix.go:7
)

import (
	"internal/bytealg"
	"runtime"
	"syscall"
)

// Boolean to int.
func boolint(b bool) int {
//line /snap/go/10455/src/net/sockopt_posix.go:16
	_go_fuzz_dep_.CoverTab[8197]++
							if b {
//line /snap/go/10455/src/net/sockopt_posix.go:17
		_go_fuzz_dep_.CoverTab[529733]++
//line /snap/go/10455/src/net/sockopt_posix.go:17
		_go_fuzz_dep_.CoverTab[8199]++
								return 1
//line /snap/go/10455/src/net/sockopt_posix.go:18
		// _ = "end of CoverTab[8199]"
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:19
		_go_fuzz_dep_.CoverTab[529734]++
//line /snap/go/10455/src/net/sockopt_posix.go:19
		_go_fuzz_dep_.CoverTab[8200]++
//line /snap/go/10455/src/net/sockopt_posix.go:19
		// _ = "end of CoverTab[8200]"
//line /snap/go/10455/src/net/sockopt_posix.go:19
	}
//line /snap/go/10455/src/net/sockopt_posix.go:19
	// _ = "end of CoverTab[8197]"
//line /snap/go/10455/src/net/sockopt_posix.go:19
	_go_fuzz_dep_.CoverTab[8198]++
							return 0
//line /snap/go/10455/src/net/sockopt_posix.go:20
	// _ = "end of CoverTab[8198]"
}

func ipv4AddrToInterface(ip IP) (*Interface, error) {
//line /snap/go/10455/src/net/sockopt_posix.go:23
	_go_fuzz_dep_.CoverTab[8201]++
							ift, err := Interfaces()
							if err != nil {
//line /snap/go/10455/src/net/sockopt_posix.go:25
		_go_fuzz_dep_.CoverTab[529735]++
//line /snap/go/10455/src/net/sockopt_posix.go:25
		_go_fuzz_dep_.CoverTab[8205]++
								return nil, err
//line /snap/go/10455/src/net/sockopt_posix.go:26
		// _ = "end of CoverTab[8205]"
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:27
		_go_fuzz_dep_.CoverTab[529736]++
//line /snap/go/10455/src/net/sockopt_posix.go:27
		_go_fuzz_dep_.CoverTab[8206]++
//line /snap/go/10455/src/net/sockopt_posix.go:27
		// _ = "end of CoverTab[8206]"
//line /snap/go/10455/src/net/sockopt_posix.go:27
	}
//line /snap/go/10455/src/net/sockopt_posix.go:27
	// _ = "end of CoverTab[8201]"
//line /snap/go/10455/src/net/sockopt_posix.go:27
	_go_fuzz_dep_.CoverTab[8202]++
//line /snap/go/10455/src/net/sockopt_posix.go:27
	_go_fuzz_dep_.CoverTab[786744] = 0
							for _, ifi := range ift {
//line /snap/go/10455/src/net/sockopt_posix.go:28
		if _go_fuzz_dep_.CoverTab[786744] == 0 {
//line /snap/go/10455/src/net/sockopt_posix.go:28
			_go_fuzz_dep_.CoverTab[529771]++
//line /snap/go/10455/src/net/sockopt_posix.go:28
		} else {
//line /snap/go/10455/src/net/sockopt_posix.go:28
			_go_fuzz_dep_.CoverTab[529772]++
//line /snap/go/10455/src/net/sockopt_posix.go:28
		}
//line /snap/go/10455/src/net/sockopt_posix.go:28
		_go_fuzz_dep_.CoverTab[786744] = 1
//line /snap/go/10455/src/net/sockopt_posix.go:28
		_go_fuzz_dep_.CoverTab[8207]++
								ifat, err := ifi.Addrs()
								if err != nil {
//line /snap/go/10455/src/net/sockopt_posix.go:30
			_go_fuzz_dep_.CoverTab[529737]++
//line /snap/go/10455/src/net/sockopt_posix.go:30
			_go_fuzz_dep_.CoverTab[8209]++
									return nil, err
//line /snap/go/10455/src/net/sockopt_posix.go:31
			// _ = "end of CoverTab[8209]"
		} else {
//line /snap/go/10455/src/net/sockopt_posix.go:32
			_go_fuzz_dep_.CoverTab[529738]++
//line /snap/go/10455/src/net/sockopt_posix.go:32
			_go_fuzz_dep_.CoverTab[8210]++
//line /snap/go/10455/src/net/sockopt_posix.go:32
			// _ = "end of CoverTab[8210]"
//line /snap/go/10455/src/net/sockopt_posix.go:32
		}
//line /snap/go/10455/src/net/sockopt_posix.go:32
		// _ = "end of CoverTab[8207]"
//line /snap/go/10455/src/net/sockopt_posix.go:32
		_go_fuzz_dep_.CoverTab[8208]++
//line /snap/go/10455/src/net/sockopt_posix.go:32
		_go_fuzz_dep_.CoverTab[786745] = 0
								for _, ifa := range ifat {
//line /snap/go/10455/src/net/sockopt_posix.go:33
			if _go_fuzz_dep_.CoverTab[786745] == 0 {
//line /snap/go/10455/src/net/sockopt_posix.go:33
				_go_fuzz_dep_.CoverTab[529775]++
//line /snap/go/10455/src/net/sockopt_posix.go:33
			} else {
//line /snap/go/10455/src/net/sockopt_posix.go:33
				_go_fuzz_dep_.CoverTab[529776]++
//line /snap/go/10455/src/net/sockopt_posix.go:33
			}
//line /snap/go/10455/src/net/sockopt_posix.go:33
			_go_fuzz_dep_.CoverTab[786745] = 1
//line /snap/go/10455/src/net/sockopt_posix.go:33
			_go_fuzz_dep_.CoverTab[8211]++
									switch v := ifa.(type) {
			case *IPAddr:
//line /snap/go/10455/src/net/sockopt_posix.go:35
				_go_fuzz_dep_.CoverTab[529739]++
//line /snap/go/10455/src/net/sockopt_posix.go:35
				_go_fuzz_dep_.CoverTab[8212]++
										if ip.Equal(v.IP) {
//line /snap/go/10455/src/net/sockopt_posix.go:36
					_go_fuzz_dep_.CoverTab[529741]++
//line /snap/go/10455/src/net/sockopt_posix.go:36
					_go_fuzz_dep_.CoverTab[8214]++
											return &ifi, nil
//line /snap/go/10455/src/net/sockopt_posix.go:37
					// _ = "end of CoverTab[8214]"
				} else {
//line /snap/go/10455/src/net/sockopt_posix.go:38
					_go_fuzz_dep_.CoverTab[529742]++
//line /snap/go/10455/src/net/sockopt_posix.go:38
					_go_fuzz_dep_.CoverTab[8215]++
//line /snap/go/10455/src/net/sockopt_posix.go:38
					// _ = "end of CoverTab[8215]"
//line /snap/go/10455/src/net/sockopt_posix.go:38
				}
//line /snap/go/10455/src/net/sockopt_posix.go:38
				// _ = "end of CoverTab[8212]"
			case *IPNet:
//line /snap/go/10455/src/net/sockopt_posix.go:39
				_go_fuzz_dep_.CoverTab[529740]++
//line /snap/go/10455/src/net/sockopt_posix.go:39
				_go_fuzz_dep_.CoverTab[8213]++
										if ip.Equal(v.IP) {
//line /snap/go/10455/src/net/sockopt_posix.go:40
					_go_fuzz_dep_.CoverTab[529743]++
//line /snap/go/10455/src/net/sockopt_posix.go:40
					_go_fuzz_dep_.CoverTab[8216]++
											return &ifi, nil
//line /snap/go/10455/src/net/sockopt_posix.go:41
					// _ = "end of CoverTab[8216]"
				} else {
//line /snap/go/10455/src/net/sockopt_posix.go:42
					_go_fuzz_dep_.CoverTab[529744]++
//line /snap/go/10455/src/net/sockopt_posix.go:42
					_go_fuzz_dep_.CoverTab[8217]++
//line /snap/go/10455/src/net/sockopt_posix.go:42
					// _ = "end of CoverTab[8217]"
//line /snap/go/10455/src/net/sockopt_posix.go:42
				}
//line /snap/go/10455/src/net/sockopt_posix.go:42
				// _ = "end of CoverTab[8213]"
			}
//line /snap/go/10455/src/net/sockopt_posix.go:43
			// _ = "end of CoverTab[8211]"
		}
//line /snap/go/10455/src/net/sockopt_posix.go:44
		if _go_fuzz_dep_.CoverTab[786745] == 0 {
//line /snap/go/10455/src/net/sockopt_posix.go:44
			_go_fuzz_dep_.CoverTab[529777]++
//line /snap/go/10455/src/net/sockopt_posix.go:44
		} else {
//line /snap/go/10455/src/net/sockopt_posix.go:44
			_go_fuzz_dep_.CoverTab[529778]++
//line /snap/go/10455/src/net/sockopt_posix.go:44
		}
//line /snap/go/10455/src/net/sockopt_posix.go:44
		// _ = "end of CoverTab[8208]"
	}
//line /snap/go/10455/src/net/sockopt_posix.go:45
	if _go_fuzz_dep_.CoverTab[786744] == 0 {
//line /snap/go/10455/src/net/sockopt_posix.go:45
		_go_fuzz_dep_.CoverTab[529773]++
//line /snap/go/10455/src/net/sockopt_posix.go:45
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:45
		_go_fuzz_dep_.CoverTab[529774]++
//line /snap/go/10455/src/net/sockopt_posix.go:45
	}
//line /snap/go/10455/src/net/sockopt_posix.go:45
	// _ = "end of CoverTab[8202]"
//line /snap/go/10455/src/net/sockopt_posix.go:45
	_go_fuzz_dep_.CoverTab[8203]++
							if ip.Equal(IPv4zero) {
//line /snap/go/10455/src/net/sockopt_posix.go:46
		_go_fuzz_dep_.CoverTab[529745]++
//line /snap/go/10455/src/net/sockopt_posix.go:46
		_go_fuzz_dep_.CoverTab[8218]++
								return nil, nil
//line /snap/go/10455/src/net/sockopt_posix.go:47
		// _ = "end of CoverTab[8218]"
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:48
		_go_fuzz_dep_.CoverTab[529746]++
//line /snap/go/10455/src/net/sockopt_posix.go:48
		_go_fuzz_dep_.CoverTab[8219]++
//line /snap/go/10455/src/net/sockopt_posix.go:48
		// _ = "end of CoverTab[8219]"
//line /snap/go/10455/src/net/sockopt_posix.go:48
	}
//line /snap/go/10455/src/net/sockopt_posix.go:48
	// _ = "end of CoverTab[8203]"
//line /snap/go/10455/src/net/sockopt_posix.go:48
	_go_fuzz_dep_.CoverTab[8204]++
							return nil, errNoSuchInterface
//line /snap/go/10455/src/net/sockopt_posix.go:49
	// _ = "end of CoverTab[8204]"
}

func interfaceToIPv4Addr(ifi *Interface) (IP, error) {
//line /snap/go/10455/src/net/sockopt_posix.go:52
	_go_fuzz_dep_.CoverTab[8220]++
							if ifi == nil {
//line /snap/go/10455/src/net/sockopt_posix.go:53
		_go_fuzz_dep_.CoverTab[529747]++
//line /snap/go/10455/src/net/sockopt_posix.go:53
		_go_fuzz_dep_.CoverTab[8224]++
								return IPv4zero, nil
//line /snap/go/10455/src/net/sockopt_posix.go:54
		// _ = "end of CoverTab[8224]"
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:55
		_go_fuzz_dep_.CoverTab[529748]++
//line /snap/go/10455/src/net/sockopt_posix.go:55
		_go_fuzz_dep_.CoverTab[8225]++
//line /snap/go/10455/src/net/sockopt_posix.go:55
		// _ = "end of CoverTab[8225]"
//line /snap/go/10455/src/net/sockopt_posix.go:55
	}
//line /snap/go/10455/src/net/sockopt_posix.go:55
	// _ = "end of CoverTab[8220]"
//line /snap/go/10455/src/net/sockopt_posix.go:55
	_go_fuzz_dep_.CoverTab[8221]++
							ifat, err := ifi.Addrs()
							if err != nil {
//line /snap/go/10455/src/net/sockopt_posix.go:57
		_go_fuzz_dep_.CoverTab[529749]++
//line /snap/go/10455/src/net/sockopt_posix.go:57
		_go_fuzz_dep_.CoverTab[8226]++
								return nil, err
//line /snap/go/10455/src/net/sockopt_posix.go:58
		// _ = "end of CoverTab[8226]"
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:59
		_go_fuzz_dep_.CoverTab[529750]++
//line /snap/go/10455/src/net/sockopt_posix.go:59
		_go_fuzz_dep_.CoverTab[8227]++
//line /snap/go/10455/src/net/sockopt_posix.go:59
		// _ = "end of CoverTab[8227]"
//line /snap/go/10455/src/net/sockopt_posix.go:59
	}
//line /snap/go/10455/src/net/sockopt_posix.go:59
	// _ = "end of CoverTab[8221]"
//line /snap/go/10455/src/net/sockopt_posix.go:59
	_go_fuzz_dep_.CoverTab[8222]++
//line /snap/go/10455/src/net/sockopt_posix.go:59
	_go_fuzz_dep_.CoverTab[786746] = 0
							for _, ifa := range ifat {
//line /snap/go/10455/src/net/sockopt_posix.go:60
		if _go_fuzz_dep_.CoverTab[786746] == 0 {
//line /snap/go/10455/src/net/sockopt_posix.go:60
			_go_fuzz_dep_.CoverTab[529779]++
//line /snap/go/10455/src/net/sockopt_posix.go:60
		} else {
//line /snap/go/10455/src/net/sockopt_posix.go:60
			_go_fuzz_dep_.CoverTab[529780]++
//line /snap/go/10455/src/net/sockopt_posix.go:60
		}
//line /snap/go/10455/src/net/sockopt_posix.go:60
		_go_fuzz_dep_.CoverTab[786746] = 1
//line /snap/go/10455/src/net/sockopt_posix.go:60
		_go_fuzz_dep_.CoverTab[8228]++
								switch v := ifa.(type) {
		case *IPAddr:
//line /snap/go/10455/src/net/sockopt_posix.go:62
			_go_fuzz_dep_.CoverTab[529751]++
//line /snap/go/10455/src/net/sockopt_posix.go:62
			_go_fuzz_dep_.CoverTab[8229]++
									if v.IP.To4() != nil {
//line /snap/go/10455/src/net/sockopt_posix.go:63
				_go_fuzz_dep_.CoverTab[529753]++
//line /snap/go/10455/src/net/sockopt_posix.go:63
				_go_fuzz_dep_.CoverTab[8231]++
										return v.IP, nil
//line /snap/go/10455/src/net/sockopt_posix.go:64
				// _ = "end of CoverTab[8231]"
			} else {
//line /snap/go/10455/src/net/sockopt_posix.go:65
				_go_fuzz_dep_.CoverTab[529754]++
//line /snap/go/10455/src/net/sockopt_posix.go:65
				_go_fuzz_dep_.CoverTab[8232]++
//line /snap/go/10455/src/net/sockopt_posix.go:65
				// _ = "end of CoverTab[8232]"
//line /snap/go/10455/src/net/sockopt_posix.go:65
			}
//line /snap/go/10455/src/net/sockopt_posix.go:65
			// _ = "end of CoverTab[8229]"
		case *IPNet:
//line /snap/go/10455/src/net/sockopt_posix.go:66
			_go_fuzz_dep_.CoverTab[529752]++
//line /snap/go/10455/src/net/sockopt_posix.go:66
			_go_fuzz_dep_.CoverTab[8230]++
									if v.IP.To4() != nil {
//line /snap/go/10455/src/net/sockopt_posix.go:67
				_go_fuzz_dep_.CoverTab[529755]++
//line /snap/go/10455/src/net/sockopt_posix.go:67
				_go_fuzz_dep_.CoverTab[8233]++
										return v.IP, nil
//line /snap/go/10455/src/net/sockopt_posix.go:68
				// _ = "end of CoverTab[8233]"
			} else {
//line /snap/go/10455/src/net/sockopt_posix.go:69
				_go_fuzz_dep_.CoverTab[529756]++
//line /snap/go/10455/src/net/sockopt_posix.go:69
				_go_fuzz_dep_.CoverTab[8234]++
//line /snap/go/10455/src/net/sockopt_posix.go:69
				// _ = "end of CoverTab[8234]"
//line /snap/go/10455/src/net/sockopt_posix.go:69
			}
//line /snap/go/10455/src/net/sockopt_posix.go:69
			// _ = "end of CoverTab[8230]"
		}
//line /snap/go/10455/src/net/sockopt_posix.go:70
		// _ = "end of CoverTab[8228]"
	}
//line /snap/go/10455/src/net/sockopt_posix.go:71
	if _go_fuzz_dep_.CoverTab[786746] == 0 {
//line /snap/go/10455/src/net/sockopt_posix.go:71
		_go_fuzz_dep_.CoverTab[529781]++
//line /snap/go/10455/src/net/sockopt_posix.go:71
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:71
		_go_fuzz_dep_.CoverTab[529782]++
//line /snap/go/10455/src/net/sockopt_posix.go:71
	}
//line /snap/go/10455/src/net/sockopt_posix.go:71
	// _ = "end of CoverTab[8222]"
//line /snap/go/10455/src/net/sockopt_posix.go:71
	_go_fuzz_dep_.CoverTab[8223]++
							return nil, errNoSuchInterface
//line /snap/go/10455/src/net/sockopt_posix.go:72
	// _ = "end of CoverTab[8223]"
}

func setIPv4MreqToInterface(mreq *syscall.IPMreq, ifi *Interface) error {
//line /snap/go/10455/src/net/sockopt_posix.go:75
	_go_fuzz_dep_.CoverTab[8235]++
							if ifi == nil {
//line /snap/go/10455/src/net/sockopt_posix.go:76
		_go_fuzz_dep_.CoverTab[529757]++
//line /snap/go/10455/src/net/sockopt_posix.go:76
		_go_fuzz_dep_.CoverTab[8240]++
								return nil
//line /snap/go/10455/src/net/sockopt_posix.go:77
		// _ = "end of CoverTab[8240]"
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:78
		_go_fuzz_dep_.CoverTab[529758]++
//line /snap/go/10455/src/net/sockopt_posix.go:78
		_go_fuzz_dep_.CoverTab[8241]++
//line /snap/go/10455/src/net/sockopt_posix.go:78
		// _ = "end of CoverTab[8241]"
//line /snap/go/10455/src/net/sockopt_posix.go:78
	}
//line /snap/go/10455/src/net/sockopt_posix.go:78
	// _ = "end of CoverTab[8235]"
//line /snap/go/10455/src/net/sockopt_posix.go:78
	_go_fuzz_dep_.CoverTab[8236]++
							ifat, err := ifi.Addrs()
							if err != nil {
//line /snap/go/10455/src/net/sockopt_posix.go:80
		_go_fuzz_dep_.CoverTab[529759]++
//line /snap/go/10455/src/net/sockopt_posix.go:80
		_go_fuzz_dep_.CoverTab[8242]++
								return err
//line /snap/go/10455/src/net/sockopt_posix.go:81
		// _ = "end of CoverTab[8242]"
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:82
		_go_fuzz_dep_.CoverTab[529760]++
//line /snap/go/10455/src/net/sockopt_posix.go:82
		_go_fuzz_dep_.CoverTab[8243]++
//line /snap/go/10455/src/net/sockopt_posix.go:82
		// _ = "end of CoverTab[8243]"
//line /snap/go/10455/src/net/sockopt_posix.go:82
	}
//line /snap/go/10455/src/net/sockopt_posix.go:82
	// _ = "end of CoverTab[8236]"
//line /snap/go/10455/src/net/sockopt_posix.go:82
	_go_fuzz_dep_.CoverTab[8237]++
//line /snap/go/10455/src/net/sockopt_posix.go:82
	_go_fuzz_dep_.CoverTab[786747] = 0
							for _, ifa := range ifat {
//line /snap/go/10455/src/net/sockopt_posix.go:83
		if _go_fuzz_dep_.CoverTab[786747] == 0 {
//line /snap/go/10455/src/net/sockopt_posix.go:83
			_go_fuzz_dep_.CoverTab[529783]++
//line /snap/go/10455/src/net/sockopt_posix.go:83
		} else {
//line /snap/go/10455/src/net/sockopt_posix.go:83
			_go_fuzz_dep_.CoverTab[529784]++
//line /snap/go/10455/src/net/sockopt_posix.go:83
		}
//line /snap/go/10455/src/net/sockopt_posix.go:83
		_go_fuzz_dep_.CoverTab[786747] = 1
//line /snap/go/10455/src/net/sockopt_posix.go:83
		_go_fuzz_dep_.CoverTab[8244]++
								switch v := ifa.(type) {
		case *IPAddr:
//line /snap/go/10455/src/net/sockopt_posix.go:85
			_go_fuzz_dep_.CoverTab[529761]++
//line /snap/go/10455/src/net/sockopt_posix.go:85
			_go_fuzz_dep_.CoverTab[8245]++
									if a := v.IP.To4(); a != nil {
//line /snap/go/10455/src/net/sockopt_posix.go:86
				_go_fuzz_dep_.CoverTab[529763]++
//line /snap/go/10455/src/net/sockopt_posix.go:86
				_go_fuzz_dep_.CoverTab[8247]++
										copy(mreq.Interface[:], a)
										goto done
//line /snap/go/10455/src/net/sockopt_posix.go:88
				// _ = "end of CoverTab[8247]"
			} else {
//line /snap/go/10455/src/net/sockopt_posix.go:89
				_go_fuzz_dep_.CoverTab[529764]++
//line /snap/go/10455/src/net/sockopt_posix.go:89
				_go_fuzz_dep_.CoverTab[8248]++
//line /snap/go/10455/src/net/sockopt_posix.go:89
				// _ = "end of CoverTab[8248]"
//line /snap/go/10455/src/net/sockopt_posix.go:89
			}
//line /snap/go/10455/src/net/sockopt_posix.go:89
			// _ = "end of CoverTab[8245]"
		case *IPNet:
//line /snap/go/10455/src/net/sockopt_posix.go:90
			_go_fuzz_dep_.CoverTab[529762]++
//line /snap/go/10455/src/net/sockopt_posix.go:90
			_go_fuzz_dep_.CoverTab[8246]++
									if a := v.IP.To4(); a != nil {
//line /snap/go/10455/src/net/sockopt_posix.go:91
				_go_fuzz_dep_.CoverTab[529765]++
//line /snap/go/10455/src/net/sockopt_posix.go:91
				_go_fuzz_dep_.CoverTab[8249]++
										copy(mreq.Interface[:], a)
										goto done
//line /snap/go/10455/src/net/sockopt_posix.go:93
				// _ = "end of CoverTab[8249]"
			} else {
//line /snap/go/10455/src/net/sockopt_posix.go:94
				_go_fuzz_dep_.CoverTab[529766]++
//line /snap/go/10455/src/net/sockopt_posix.go:94
				_go_fuzz_dep_.CoverTab[8250]++
//line /snap/go/10455/src/net/sockopt_posix.go:94
				// _ = "end of CoverTab[8250]"
//line /snap/go/10455/src/net/sockopt_posix.go:94
			}
//line /snap/go/10455/src/net/sockopt_posix.go:94
			// _ = "end of CoverTab[8246]"
		}
//line /snap/go/10455/src/net/sockopt_posix.go:95
		// _ = "end of CoverTab[8244]"
	}
//line /snap/go/10455/src/net/sockopt_posix.go:96
	if _go_fuzz_dep_.CoverTab[786747] == 0 {
//line /snap/go/10455/src/net/sockopt_posix.go:96
		_go_fuzz_dep_.CoverTab[529785]++
//line /snap/go/10455/src/net/sockopt_posix.go:96
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:96
		_go_fuzz_dep_.CoverTab[529786]++
//line /snap/go/10455/src/net/sockopt_posix.go:96
	}
//line /snap/go/10455/src/net/sockopt_posix.go:96
	// _ = "end of CoverTab[8237]"
//line /snap/go/10455/src/net/sockopt_posix.go:96
	_go_fuzz_dep_.CoverTab[8238]++
done:
	if bytealg.Equal(mreq.Multiaddr[:], IPv4zero.To4()) {
//line /snap/go/10455/src/net/sockopt_posix.go:98
		_go_fuzz_dep_.CoverTab[529767]++
//line /snap/go/10455/src/net/sockopt_posix.go:98
		_go_fuzz_dep_.CoverTab[8251]++
								return errNoSuchMulticastInterface
//line /snap/go/10455/src/net/sockopt_posix.go:99
		// _ = "end of CoverTab[8251]"
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:100
		_go_fuzz_dep_.CoverTab[529768]++
//line /snap/go/10455/src/net/sockopt_posix.go:100
		_go_fuzz_dep_.CoverTab[8252]++
//line /snap/go/10455/src/net/sockopt_posix.go:100
		// _ = "end of CoverTab[8252]"
//line /snap/go/10455/src/net/sockopt_posix.go:100
	}
//line /snap/go/10455/src/net/sockopt_posix.go:100
	// _ = "end of CoverTab[8238]"
//line /snap/go/10455/src/net/sockopt_posix.go:100
	_go_fuzz_dep_.CoverTab[8239]++
							return nil
//line /snap/go/10455/src/net/sockopt_posix.go:101
	// _ = "end of CoverTab[8239]"
}

func setReadBuffer(fd *netFD, bytes int) error {
//line /snap/go/10455/src/net/sockopt_posix.go:104
	_go_fuzz_dep_.CoverTab[8253]++
							err := fd.pfd.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_RCVBUF, bytes)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockopt_posix.go:107
	// _ = "end of CoverTab[8253]"
}

func setWriteBuffer(fd *netFD, bytes int) error {
//line /snap/go/10455/src/net/sockopt_posix.go:110
	_go_fuzz_dep_.CoverTab[8254]++
							err := fd.pfd.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_SNDBUF, bytes)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockopt_posix.go:113
	// _ = "end of CoverTab[8254]"
}

func setKeepAlive(fd *netFD, keepalive bool) error {
//line /snap/go/10455/src/net/sockopt_posix.go:116
	_go_fuzz_dep_.CoverTab[8255]++
							err := fd.pfd.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_KEEPALIVE, boolint(keepalive))
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockopt_posix.go:119
	// _ = "end of CoverTab[8255]"
}

func setLinger(fd *netFD, sec int) error {
//line /snap/go/10455/src/net/sockopt_posix.go:122
	_go_fuzz_dep_.CoverTab[8256]++
							var l syscall.Linger
							if sec >= 0 {
//line /snap/go/10455/src/net/sockopt_posix.go:124
		_go_fuzz_dep_.CoverTab[529769]++
//line /snap/go/10455/src/net/sockopt_posix.go:124
		_go_fuzz_dep_.CoverTab[8258]++
								l.Onoff = 1
								l.Linger = int32(sec)
//line /snap/go/10455/src/net/sockopt_posix.go:126
		// _ = "end of CoverTab[8258]"
	} else {
//line /snap/go/10455/src/net/sockopt_posix.go:127
		_go_fuzz_dep_.CoverTab[529770]++
//line /snap/go/10455/src/net/sockopt_posix.go:127
		_go_fuzz_dep_.CoverTab[8259]++
								l.Onoff = 0
								l.Linger = 0
//line /snap/go/10455/src/net/sockopt_posix.go:129
		// _ = "end of CoverTab[8259]"
	}
//line /snap/go/10455/src/net/sockopt_posix.go:130
	// _ = "end of CoverTab[8256]"
//line /snap/go/10455/src/net/sockopt_posix.go:130
	_go_fuzz_dep_.CoverTab[8257]++
							err := fd.pfd.SetsockoptLinger(syscall.SOL_SOCKET, syscall.SO_LINGER, &l)
							runtime.KeepAlive(fd)
							return wrapSyscallError("setsockopt", err)
//line /snap/go/10455/src/net/sockopt_posix.go:133
	// _ = "end of CoverTab[8257]"
}

//line /snap/go/10455/src/net/sockopt_posix.go:134
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/sockopt_posix.go:134
var _ = _go_fuzz_dep_.CoverTab
