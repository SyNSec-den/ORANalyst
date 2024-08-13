// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/interface_linux.go:5
package net

//line /usr/local/go/src/net/interface_linux.go:5
import (
//line /usr/local/go/src/net/interface_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/interface_linux.go:5
)
//line /usr/local/go/src/net/interface_linux.go:5
import (
//line /usr/local/go/src/net/interface_linux.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/interface_linux.go:5
)

import (
	"os"
	"syscall"
	"unsafe"
)

// If the ifindex is zero, interfaceTable returns mappings of all
//line /usr/local/go/src/net/interface_linux.go:13
// network interfaces. Otherwise it returns a mapping of a specific
//line /usr/local/go/src/net/interface_linux.go:13
// interface.
//line /usr/local/go/src/net/interface_linux.go:16
func interfaceTable(ifindex int) ([]Interface, error) {
//line /usr/local/go/src/net/interface_linux.go:16
	_go_fuzz_dep_.CoverTab[14161]++
							tab, err := syscall.NetlinkRIB(syscall.RTM_GETLINK, syscall.AF_UNSPEC)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:18
		_go_fuzz_dep_.CoverTab[14165]++
								return nil, os.NewSyscallError("netlinkrib", err)
//line /usr/local/go/src/net/interface_linux.go:19
		// _ = "end of CoverTab[14165]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:20
		_go_fuzz_dep_.CoverTab[14166]++
//line /usr/local/go/src/net/interface_linux.go:20
		// _ = "end of CoverTab[14166]"
//line /usr/local/go/src/net/interface_linux.go:20
	}
//line /usr/local/go/src/net/interface_linux.go:20
	// _ = "end of CoverTab[14161]"
//line /usr/local/go/src/net/interface_linux.go:20
	_go_fuzz_dep_.CoverTab[14162]++
							msgs, err := syscall.ParseNetlinkMessage(tab)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:22
		_go_fuzz_dep_.CoverTab[14167]++
								return nil, os.NewSyscallError("parsenetlinkmessage", err)
//line /usr/local/go/src/net/interface_linux.go:23
		// _ = "end of CoverTab[14167]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:24
		_go_fuzz_dep_.CoverTab[14168]++
//line /usr/local/go/src/net/interface_linux.go:24
		// _ = "end of CoverTab[14168]"
//line /usr/local/go/src/net/interface_linux.go:24
	}
//line /usr/local/go/src/net/interface_linux.go:24
	// _ = "end of CoverTab[14162]"
//line /usr/local/go/src/net/interface_linux.go:24
	_go_fuzz_dep_.CoverTab[14163]++
							var ift []Interface
loop:
	for _, m := range msgs {
//line /usr/local/go/src/net/interface_linux.go:27
		_go_fuzz_dep_.CoverTab[14169]++
								switch m.Header.Type {
		case syscall.NLMSG_DONE:
//line /usr/local/go/src/net/interface_linux.go:29
			_go_fuzz_dep_.CoverTab[14170]++
									break loop
//line /usr/local/go/src/net/interface_linux.go:30
			// _ = "end of CoverTab[14170]"
		case syscall.RTM_NEWLINK:
//line /usr/local/go/src/net/interface_linux.go:31
			_go_fuzz_dep_.CoverTab[14171]++
									ifim := (*syscall.IfInfomsg)(unsafe.Pointer(&m.Data[0]))
									if ifindex == 0 || func() bool {
//line /usr/local/go/src/net/interface_linux.go:33
				_go_fuzz_dep_.CoverTab[14173]++
//line /usr/local/go/src/net/interface_linux.go:33
				return ifindex == int(ifim.Index)
//line /usr/local/go/src/net/interface_linux.go:33
				// _ = "end of CoverTab[14173]"
//line /usr/local/go/src/net/interface_linux.go:33
			}() {
//line /usr/local/go/src/net/interface_linux.go:33
				_go_fuzz_dep_.CoverTab[14174]++
										attrs, err := syscall.ParseNetlinkRouteAttr(&m)
										if err != nil {
//line /usr/local/go/src/net/interface_linux.go:35
					_go_fuzz_dep_.CoverTab[14176]++
											return nil, os.NewSyscallError("parsenetlinkrouteattr", err)
//line /usr/local/go/src/net/interface_linux.go:36
					// _ = "end of CoverTab[14176]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:37
					_go_fuzz_dep_.CoverTab[14177]++
//line /usr/local/go/src/net/interface_linux.go:37
					// _ = "end of CoverTab[14177]"
//line /usr/local/go/src/net/interface_linux.go:37
				}
//line /usr/local/go/src/net/interface_linux.go:37
				// _ = "end of CoverTab[14174]"
//line /usr/local/go/src/net/interface_linux.go:37
				_go_fuzz_dep_.CoverTab[14175]++
										ift = append(ift, *newLink(ifim, attrs))
										if ifindex == int(ifim.Index) {
//line /usr/local/go/src/net/interface_linux.go:39
					_go_fuzz_dep_.CoverTab[14178]++
											break loop
//line /usr/local/go/src/net/interface_linux.go:40
					// _ = "end of CoverTab[14178]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:41
					_go_fuzz_dep_.CoverTab[14179]++
//line /usr/local/go/src/net/interface_linux.go:41
					// _ = "end of CoverTab[14179]"
//line /usr/local/go/src/net/interface_linux.go:41
				}
//line /usr/local/go/src/net/interface_linux.go:41
				// _ = "end of CoverTab[14175]"
			} else {
//line /usr/local/go/src/net/interface_linux.go:42
				_go_fuzz_dep_.CoverTab[14180]++
//line /usr/local/go/src/net/interface_linux.go:42
				// _ = "end of CoverTab[14180]"
//line /usr/local/go/src/net/interface_linux.go:42
			}
//line /usr/local/go/src/net/interface_linux.go:42
			// _ = "end of CoverTab[14171]"
//line /usr/local/go/src/net/interface_linux.go:42
		default:
//line /usr/local/go/src/net/interface_linux.go:42
			_go_fuzz_dep_.CoverTab[14172]++
//line /usr/local/go/src/net/interface_linux.go:42
			// _ = "end of CoverTab[14172]"
		}
//line /usr/local/go/src/net/interface_linux.go:43
		// _ = "end of CoverTab[14169]"
	}
//line /usr/local/go/src/net/interface_linux.go:44
	// _ = "end of CoverTab[14163]"
//line /usr/local/go/src/net/interface_linux.go:44
	_go_fuzz_dep_.CoverTab[14164]++
							return ift, nil
//line /usr/local/go/src/net/interface_linux.go:45
	// _ = "end of CoverTab[14164]"
}

const (
	// See linux/if_arp.h.
	// Note that Linux doesn't support IPv4 over IPv6 tunneling.
	sysARPHardwareIPv4IPv4	= 768	// IPv4 over IPv4 tunneling
	sysARPHardwareIPv6IPv6	= 769	// IPv6 over IPv6 tunneling
	sysARPHardwareIPv6IPv4	= 776	// IPv6 over IPv4 tunneling
	sysARPHardwareGREIPv4	= 778	// any over GRE over IPv4 tunneling
	sysARPHardwareGREIPv6	= 823	// any over GRE over IPv6 tunneling
)

func newLink(ifim *syscall.IfInfomsg, attrs []syscall.NetlinkRouteAttr) *Interface {
//line /usr/local/go/src/net/interface_linux.go:58
	_go_fuzz_dep_.CoverTab[14181]++
							ifi := &Interface{Index: int(ifim.Index), Flags: linkFlags(ifim.Flags)}
							for _, a := range attrs {
//line /usr/local/go/src/net/interface_linux.go:60
		_go_fuzz_dep_.CoverTab[14183]++
								switch a.Attr.Type {
		case syscall.IFLA_ADDRESS:
//line /usr/local/go/src/net/interface_linux.go:62
			_go_fuzz_dep_.CoverTab[14184]++

//line /usr/local/go/src/net/interface_linux.go:66
			switch len(a.Value) {
			case IPv4len:
//line /usr/local/go/src/net/interface_linux.go:67
				_go_fuzz_dep_.CoverTab[14190]++
										switch ifim.Type {
				case sysARPHardwareIPv4IPv4, sysARPHardwareGREIPv4, sysARPHardwareIPv6IPv4:
//line /usr/local/go/src/net/interface_linux.go:69
					_go_fuzz_dep_.CoverTab[14193]++
											continue
//line /usr/local/go/src/net/interface_linux.go:70
					// _ = "end of CoverTab[14193]"
//line /usr/local/go/src/net/interface_linux.go:70
				default:
//line /usr/local/go/src/net/interface_linux.go:70
					_go_fuzz_dep_.CoverTab[14194]++
//line /usr/local/go/src/net/interface_linux.go:70
					// _ = "end of CoverTab[14194]"
				}
//line /usr/local/go/src/net/interface_linux.go:71
				// _ = "end of CoverTab[14190]"
			case IPv6len:
//line /usr/local/go/src/net/interface_linux.go:72
				_go_fuzz_dep_.CoverTab[14191]++
										switch ifim.Type {
				case sysARPHardwareIPv6IPv6, sysARPHardwareGREIPv6:
//line /usr/local/go/src/net/interface_linux.go:74
					_go_fuzz_dep_.CoverTab[14195]++
											continue
//line /usr/local/go/src/net/interface_linux.go:75
					// _ = "end of CoverTab[14195]"
//line /usr/local/go/src/net/interface_linux.go:75
				default:
//line /usr/local/go/src/net/interface_linux.go:75
					_go_fuzz_dep_.CoverTab[14196]++
//line /usr/local/go/src/net/interface_linux.go:75
					// _ = "end of CoverTab[14196]"
				}
//line /usr/local/go/src/net/interface_linux.go:76
				// _ = "end of CoverTab[14191]"
//line /usr/local/go/src/net/interface_linux.go:76
			default:
//line /usr/local/go/src/net/interface_linux.go:76
				_go_fuzz_dep_.CoverTab[14192]++
//line /usr/local/go/src/net/interface_linux.go:76
				// _ = "end of CoverTab[14192]"
			}
//line /usr/local/go/src/net/interface_linux.go:77
			// _ = "end of CoverTab[14184]"
//line /usr/local/go/src/net/interface_linux.go:77
			_go_fuzz_dep_.CoverTab[14185]++
									var nonzero bool
									for _, b := range a.Value {
//line /usr/local/go/src/net/interface_linux.go:79
				_go_fuzz_dep_.CoverTab[14197]++
										if b != 0 {
//line /usr/local/go/src/net/interface_linux.go:80
					_go_fuzz_dep_.CoverTab[14198]++
											nonzero = true
											break
//line /usr/local/go/src/net/interface_linux.go:82
					// _ = "end of CoverTab[14198]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:83
					_go_fuzz_dep_.CoverTab[14199]++
//line /usr/local/go/src/net/interface_linux.go:83
					// _ = "end of CoverTab[14199]"
//line /usr/local/go/src/net/interface_linux.go:83
				}
//line /usr/local/go/src/net/interface_linux.go:83
				// _ = "end of CoverTab[14197]"
			}
//line /usr/local/go/src/net/interface_linux.go:84
			// _ = "end of CoverTab[14185]"
//line /usr/local/go/src/net/interface_linux.go:84
			_go_fuzz_dep_.CoverTab[14186]++
									if nonzero {
//line /usr/local/go/src/net/interface_linux.go:85
				_go_fuzz_dep_.CoverTab[14200]++
										ifi.HardwareAddr = a.Value[:]
//line /usr/local/go/src/net/interface_linux.go:86
				// _ = "end of CoverTab[14200]"
			} else {
//line /usr/local/go/src/net/interface_linux.go:87
				_go_fuzz_dep_.CoverTab[14201]++
//line /usr/local/go/src/net/interface_linux.go:87
				// _ = "end of CoverTab[14201]"
//line /usr/local/go/src/net/interface_linux.go:87
			}
//line /usr/local/go/src/net/interface_linux.go:87
			// _ = "end of CoverTab[14186]"
		case syscall.IFLA_IFNAME:
//line /usr/local/go/src/net/interface_linux.go:88
			_go_fuzz_dep_.CoverTab[14187]++
									ifi.Name = string(a.Value[:len(a.Value)-1])
//line /usr/local/go/src/net/interface_linux.go:89
			// _ = "end of CoverTab[14187]"
		case syscall.IFLA_MTU:
//line /usr/local/go/src/net/interface_linux.go:90
			_go_fuzz_dep_.CoverTab[14188]++
									ifi.MTU = int(*(*uint32)(unsafe.Pointer(&a.Value[:4][0])))
//line /usr/local/go/src/net/interface_linux.go:91
			// _ = "end of CoverTab[14188]"
//line /usr/local/go/src/net/interface_linux.go:91
		default:
//line /usr/local/go/src/net/interface_linux.go:91
			_go_fuzz_dep_.CoverTab[14189]++
//line /usr/local/go/src/net/interface_linux.go:91
			// _ = "end of CoverTab[14189]"
		}
//line /usr/local/go/src/net/interface_linux.go:92
		// _ = "end of CoverTab[14183]"
	}
//line /usr/local/go/src/net/interface_linux.go:93
	// _ = "end of CoverTab[14181]"
//line /usr/local/go/src/net/interface_linux.go:93
	_go_fuzz_dep_.CoverTab[14182]++
							return ifi
//line /usr/local/go/src/net/interface_linux.go:94
	// _ = "end of CoverTab[14182]"
}

func linkFlags(rawFlags uint32) Flags {
//line /usr/local/go/src/net/interface_linux.go:97
	_go_fuzz_dep_.CoverTab[14202]++
							var f Flags
							if rawFlags&syscall.IFF_UP != 0 {
//line /usr/local/go/src/net/interface_linux.go:99
		_go_fuzz_dep_.CoverTab[14209]++
								f |= FlagUp
//line /usr/local/go/src/net/interface_linux.go:100
		// _ = "end of CoverTab[14209]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:101
		_go_fuzz_dep_.CoverTab[14210]++
//line /usr/local/go/src/net/interface_linux.go:101
		// _ = "end of CoverTab[14210]"
//line /usr/local/go/src/net/interface_linux.go:101
	}
//line /usr/local/go/src/net/interface_linux.go:101
	// _ = "end of CoverTab[14202]"
//line /usr/local/go/src/net/interface_linux.go:101
	_go_fuzz_dep_.CoverTab[14203]++
							if rawFlags&syscall.IFF_RUNNING != 0 {
//line /usr/local/go/src/net/interface_linux.go:102
		_go_fuzz_dep_.CoverTab[14211]++
								f |= FlagRunning
//line /usr/local/go/src/net/interface_linux.go:103
		// _ = "end of CoverTab[14211]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:104
		_go_fuzz_dep_.CoverTab[14212]++
//line /usr/local/go/src/net/interface_linux.go:104
		// _ = "end of CoverTab[14212]"
//line /usr/local/go/src/net/interface_linux.go:104
	}
//line /usr/local/go/src/net/interface_linux.go:104
	// _ = "end of CoverTab[14203]"
//line /usr/local/go/src/net/interface_linux.go:104
	_go_fuzz_dep_.CoverTab[14204]++
							if rawFlags&syscall.IFF_BROADCAST != 0 {
//line /usr/local/go/src/net/interface_linux.go:105
		_go_fuzz_dep_.CoverTab[14213]++
								f |= FlagBroadcast
//line /usr/local/go/src/net/interface_linux.go:106
		// _ = "end of CoverTab[14213]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:107
		_go_fuzz_dep_.CoverTab[14214]++
//line /usr/local/go/src/net/interface_linux.go:107
		// _ = "end of CoverTab[14214]"
//line /usr/local/go/src/net/interface_linux.go:107
	}
//line /usr/local/go/src/net/interface_linux.go:107
	// _ = "end of CoverTab[14204]"
//line /usr/local/go/src/net/interface_linux.go:107
	_go_fuzz_dep_.CoverTab[14205]++
							if rawFlags&syscall.IFF_LOOPBACK != 0 {
//line /usr/local/go/src/net/interface_linux.go:108
		_go_fuzz_dep_.CoverTab[14215]++
								f |= FlagLoopback
//line /usr/local/go/src/net/interface_linux.go:109
		// _ = "end of CoverTab[14215]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:110
		_go_fuzz_dep_.CoverTab[14216]++
//line /usr/local/go/src/net/interface_linux.go:110
		// _ = "end of CoverTab[14216]"
//line /usr/local/go/src/net/interface_linux.go:110
	}
//line /usr/local/go/src/net/interface_linux.go:110
	// _ = "end of CoverTab[14205]"
//line /usr/local/go/src/net/interface_linux.go:110
	_go_fuzz_dep_.CoverTab[14206]++
							if rawFlags&syscall.IFF_POINTOPOINT != 0 {
//line /usr/local/go/src/net/interface_linux.go:111
		_go_fuzz_dep_.CoverTab[14217]++
								f |= FlagPointToPoint
//line /usr/local/go/src/net/interface_linux.go:112
		// _ = "end of CoverTab[14217]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:113
		_go_fuzz_dep_.CoverTab[14218]++
//line /usr/local/go/src/net/interface_linux.go:113
		// _ = "end of CoverTab[14218]"
//line /usr/local/go/src/net/interface_linux.go:113
	}
//line /usr/local/go/src/net/interface_linux.go:113
	// _ = "end of CoverTab[14206]"
//line /usr/local/go/src/net/interface_linux.go:113
	_go_fuzz_dep_.CoverTab[14207]++
							if rawFlags&syscall.IFF_MULTICAST != 0 {
//line /usr/local/go/src/net/interface_linux.go:114
		_go_fuzz_dep_.CoverTab[14219]++
								f |= FlagMulticast
//line /usr/local/go/src/net/interface_linux.go:115
		// _ = "end of CoverTab[14219]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:116
		_go_fuzz_dep_.CoverTab[14220]++
//line /usr/local/go/src/net/interface_linux.go:116
		// _ = "end of CoverTab[14220]"
//line /usr/local/go/src/net/interface_linux.go:116
	}
//line /usr/local/go/src/net/interface_linux.go:116
	// _ = "end of CoverTab[14207]"
//line /usr/local/go/src/net/interface_linux.go:116
	_go_fuzz_dep_.CoverTab[14208]++
							return f
//line /usr/local/go/src/net/interface_linux.go:117
	// _ = "end of CoverTab[14208]"
}

// If the ifi is nil, interfaceAddrTable returns addresses for all
//line /usr/local/go/src/net/interface_linux.go:120
// network interfaces. Otherwise it returns addresses for a specific
//line /usr/local/go/src/net/interface_linux.go:120
// interface.
//line /usr/local/go/src/net/interface_linux.go:123
func interfaceAddrTable(ifi *Interface) ([]Addr, error) {
//line /usr/local/go/src/net/interface_linux.go:123
	_go_fuzz_dep_.CoverTab[14221]++
							tab, err := syscall.NetlinkRIB(syscall.RTM_GETADDR, syscall.AF_UNSPEC)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:125
		_go_fuzz_dep_.CoverTab[14226]++
								return nil, os.NewSyscallError("netlinkrib", err)
//line /usr/local/go/src/net/interface_linux.go:126
		// _ = "end of CoverTab[14226]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:127
		_go_fuzz_dep_.CoverTab[14227]++
//line /usr/local/go/src/net/interface_linux.go:127
		// _ = "end of CoverTab[14227]"
//line /usr/local/go/src/net/interface_linux.go:127
	}
//line /usr/local/go/src/net/interface_linux.go:127
	// _ = "end of CoverTab[14221]"
//line /usr/local/go/src/net/interface_linux.go:127
	_go_fuzz_dep_.CoverTab[14222]++
							msgs, err := syscall.ParseNetlinkMessage(tab)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:129
		_go_fuzz_dep_.CoverTab[14228]++
								return nil, os.NewSyscallError("parsenetlinkmessage", err)
//line /usr/local/go/src/net/interface_linux.go:130
		// _ = "end of CoverTab[14228]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:131
		_go_fuzz_dep_.CoverTab[14229]++
//line /usr/local/go/src/net/interface_linux.go:131
		// _ = "end of CoverTab[14229]"
//line /usr/local/go/src/net/interface_linux.go:131
	}
//line /usr/local/go/src/net/interface_linux.go:131
	// _ = "end of CoverTab[14222]"
//line /usr/local/go/src/net/interface_linux.go:131
	_go_fuzz_dep_.CoverTab[14223]++
							var ift []Interface
							if ifi == nil {
//line /usr/local/go/src/net/interface_linux.go:133
		_go_fuzz_dep_.CoverTab[14230]++
								var err error
								ift, err = interfaceTable(0)
								if err != nil {
//line /usr/local/go/src/net/interface_linux.go:136
			_go_fuzz_dep_.CoverTab[14231]++
									return nil, err
//line /usr/local/go/src/net/interface_linux.go:137
			// _ = "end of CoverTab[14231]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:138
			_go_fuzz_dep_.CoverTab[14232]++
//line /usr/local/go/src/net/interface_linux.go:138
			// _ = "end of CoverTab[14232]"
//line /usr/local/go/src/net/interface_linux.go:138
		}
//line /usr/local/go/src/net/interface_linux.go:138
		// _ = "end of CoverTab[14230]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:139
		_go_fuzz_dep_.CoverTab[14233]++
//line /usr/local/go/src/net/interface_linux.go:139
		// _ = "end of CoverTab[14233]"
//line /usr/local/go/src/net/interface_linux.go:139
	}
//line /usr/local/go/src/net/interface_linux.go:139
	// _ = "end of CoverTab[14223]"
//line /usr/local/go/src/net/interface_linux.go:139
	_go_fuzz_dep_.CoverTab[14224]++
							ifat, err := addrTable(ift, ifi, msgs)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:141
		_go_fuzz_dep_.CoverTab[14234]++
								return nil, err
//line /usr/local/go/src/net/interface_linux.go:142
		// _ = "end of CoverTab[14234]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:143
		_go_fuzz_dep_.CoverTab[14235]++
//line /usr/local/go/src/net/interface_linux.go:143
		// _ = "end of CoverTab[14235]"
//line /usr/local/go/src/net/interface_linux.go:143
	}
//line /usr/local/go/src/net/interface_linux.go:143
	// _ = "end of CoverTab[14224]"
//line /usr/local/go/src/net/interface_linux.go:143
	_go_fuzz_dep_.CoverTab[14225]++
							return ifat, nil
//line /usr/local/go/src/net/interface_linux.go:144
	// _ = "end of CoverTab[14225]"
}

func addrTable(ift []Interface, ifi *Interface, msgs []syscall.NetlinkMessage) ([]Addr, error) {
//line /usr/local/go/src/net/interface_linux.go:147
	_go_fuzz_dep_.CoverTab[14236]++
							var ifat []Addr
loop:
	for _, m := range msgs {
//line /usr/local/go/src/net/interface_linux.go:150
		_go_fuzz_dep_.CoverTab[14238]++
								switch m.Header.Type {
		case syscall.NLMSG_DONE:
//line /usr/local/go/src/net/interface_linux.go:152
			_go_fuzz_dep_.CoverTab[14239]++
									break loop
//line /usr/local/go/src/net/interface_linux.go:153
			// _ = "end of CoverTab[14239]"
		case syscall.RTM_NEWADDR:
//line /usr/local/go/src/net/interface_linux.go:154
			_go_fuzz_dep_.CoverTab[14240]++
									ifam := (*syscall.IfAddrmsg)(unsafe.Pointer(&m.Data[0]))
									if len(ift) != 0 || func() bool {
//line /usr/local/go/src/net/interface_linux.go:156
				_go_fuzz_dep_.CoverTab[14242]++
//line /usr/local/go/src/net/interface_linux.go:156
				return ifi.Index == int(ifam.Index)
//line /usr/local/go/src/net/interface_linux.go:156
				// _ = "end of CoverTab[14242]"
//line /usr/local/go/src/net/interface_linux.go:156
			}() {
//line /usr/local/go/src/net/interface_linux.go:156
				_go_fuzz_dep_.CoverTab[14243]++
										if len(ift) != 0 {
//line /usr/local/go/src/net/interface_linux.go:157
					_go_fuzz_dep_.CoverTab[14246]++
											var err error
											ifi, err = interfaceByIndex(ift, int(ifam.Index))
											if err != nil {
//line /usr/local/go/src/net/interface_linux.go:160
						_go_fuzz_dep_.CoverTab[14247]++
												return nil, err
//line /usr/local/go/src/net/interface_linux.go:161
						// _ = "end of CoverTab[14247]"
					} else {
//line /usr/local/go/src/net/interface_linux.go:162
						_go_fuzz_dep_.CoverTab[14248]++
//line /usr/local/go/src/net/interface_linux.go:162
						// _ = "end of CoverTab[14248]"
//line /usr/local/go/src/net/interface_linux.go:162
					}
//line /usr/local/go/src/net/interface_linux.go:162
					// _ = "end of CoverTab[14246]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:163
					_go_fuzz_dep_.CoverTab[14249]++
//line /usr/local/go/src/net/interface_linux.go:163
					// _ = "end of CoverTab[14249]"
//line /usr/local/go/src/net/interface_linux.go:163
				}
//line /usr/local/go/src/net/interface_linux.go:163
				// _ = "end of CoverTab[14243]"
//line /usr/local/go/src/net/interface_linux.go:163
				_go_fuzz_dep_.CoverTab[14244]++
										attrs, err := syscall.ParseNetlinkRouteAttr(&m)
										if err != nil {
//line /usr/local/go/src/net/interface_linux.go:165
					_go_fuzz_dep_.CoverTab[14250]++
											return nil, os.NewSyscallError("parsenetlinkrouteattr", err)
//line /usr/local/go/src/net/interface_linux.go:166
					// _ = "end of CoverTab[14250]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:167
					_go_fuzz_dep_.CoverTab[14251]++
//line /usr/local/go/src/net/interface_linux.go:167
					// _ = "end of CoverTab[14251]"
//line /usr/local/go/src/net/interface_linux.go:167
				}
//line /usr/local/go/src/net/interface_linux.go:167
				// _ = "end of CoverTab[14244]"
//line /usr/local/go/src/net/interface_linux.go:167
				_go_fuzz_dep_.CoverTab[14245]++
										ifa := newAddr(ifam, attrs)
										if ifa != nil {
//line /usr/local/go/src/net/interface_linux.go:169
					_go_fuzz_dep_.CoverTab[14252]++
											ifat = append(ifat, ifa)
//line /usr/local/go/src/net/interface_linux.go:170
					// _ = "end of CoverTab[14252]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:171
					_go_fuzz_dep_.CoverTab[14253]++
//line /usr/local/go/src/net/interface_linux.go:171
					// _ = "end of CoverTab[14253]"
//line /usr/local/go/src/net/interface_linux.go:171
				}
//line /usr/local/go/src/net/interface_linux.go:171
				// _ = "end of CoverTab[14245]"
			} else {
//line /usr/local/go/src/net/interface_linux.go:172
				_go_fuzz_dep_.CoverTab[14254]++
//line /usr/local/go/src/net/interface_linux.go:172
				// _ = "end of CoverTab[14254]"
//line /usr/local/go/src/net/interface_linux.go:172
			}
//line /usr/local/go/src/net/interface_linux.go:172
			// _ = "end of CoverTab[14240]"
//line /usr/local/go/src/net/interface_linux.go:172
		default:
//line /usr/local/go/src/net/interface_linux.go:172
			_go_fuzz_dep_.CoverTab[14241]++
//line /usr/local/go/src/net/interface_linux.go:172
			// _ = "end of CoverTab[14241]"
		}
//line /usr/local/go/src/net/interface_linux.go:173
		// _ = "end of CoverTab[14238]"
	}
//line /usr/local/go/src/net/interface_linux.go:174
	// _ = "end of CoverTab[14236]"
//line /usr/local/go/src/net/interface_linux.go:174
	_go_fuzz_dep_.CoverTab[14237]++
							return ifat, nil
//line /usr/local/go/src/net/interface_linux.go:175
	// _ = "end of CoverTab[14237]"
}

func newAddr(ifam *syscall.IfAddrmsg, attrs []syscall.NetlinkRouteAttr) Addr {
//line /usr/local/go/src/net/interface_linux.go:178
	_go_fuzz_dep_.CoverTab[14255]++
							var ipPointToPoint bool

//line /usr/local/go/src/net/interface_linux.go:183
	for _, a := range attrs {
//line /usr/local/go/src/net/interface_linux.go:183
		_go_fuzz_dep_.CoverTab[14258]++
								if a.Attr.Type == syscall.IFA_LOCAL {
//line /usr/local/go/src/net/interface_linux.go:184
			_go_fuzz_dep_.CoverTab[14259]++
									ipPointToPoint = true
									break
//line /usr/local/go/src/net/interface_linux.go:186
			// _ = "end of CoverTab[14259]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:187
			_go_fuzz_dep_.CoverTab[14260]++
//line /usr/local/go/src/net/interface_linux.go:187
			// _ = "end of CoverTab[14260]"
//line /usr/local/go/src/net/interface_linux.go:187
		}
//line /usr/local/go/src/net/interface_linux.go:187
		// _ = "end of CoverTab[14258]"
	}
//line /usr/local/go/src/net/interface_linux.go:188
	// _ = "end of CoverTab[14255]"
//line /usr/local/go/src/net/interface_linux.go:188
	_go_fuzz_dep_.CoverTab[14256]++
							for _, a := range attrs {
//line /usr/local/go/src/net/interface_linux.go:189
		_go_fuzz_dep_.CoverTab[14261]++
								if ipPointToPoint && func() bool {
//line /usr/local/go/src/net/interface_linux.go:190
			_go_fuzz_dep_.CoverTab[14263]++
//line /usr/local/go/src/net/interface_linux.go:190
			return a.Attr.Type == syscall.IFA_ADDRESS
//line /usr/local/go/src/net/interface_linux.go:190
			// _ = "end of CoverTab[14263]"
//line /usr/local/go/src/net/interface_linux.go:190
		}() {
//line /usr/local/go/src/net/interface_linux.go:190
			_go_fuzz_dep_.CoverTab[14264]++
									continue
//line /usr/local/go/src/net/interface_linux.go:191
			// _ = "end of CoverTab[14264]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:192
			_go_fuzz_dep_.CoverTab[14265]++
//line /usr/local/go/src/net/interface_linux.go:192
			// _ = "end of CoverTab[14265]"
//line /usr/local/go/src/net/interface_linux.go:192
		}
//line /usr/local/go/src/net/interface_linux.go:192
		// _ = "end of CoverTab[14261]"
//line /usr/local/go/src/net/interface_linux.go:192
		_go_fuzz_dep_.CoverTab[14262]++
								switch ifam.Family {
		case syscall.AF_INET:
//line /usr/local/go/src/net/interface_linux.go:194
			_go_fuzz_dep_.CoverTab[14266]++
									return &IPNet{IP: IPv4(a.Value[0], a.Value[1], a.Value[2], a.Value[3]), Mask: CIDRMask(int(ifam.Prefixlen), 8*IPv4len)}
//line /usr/local/go/src/net/interface_linux.go:195
			// _ = "end of CoverTab[14266]"
		case syscall.AF_INET6:
//line /usr/local/go/src/net/interface_linux.go:196
			_go_fuzz_dep_.CoverTab[14267]++
									ifa := &IPNet{IP: make(IP, IPv6len), Mask: CIDRMask(int(ifam.Prefixlen), 8*IPv6len)}
									copy(ifa.IP, a.Value[:])
									return ifa
//line /usr/local/go/src/net/interface_linux.go:199
			// _ = "end of CoverTab[14267]"
//line /usr/local/go/src/net/interface_linux.go:199
		default:
//line /usr/local/go/src/net/interface_linux.go:199
			_go_fuzz_dep_.CoverTab[14268]++
//line /usr/local/go/src/net/interface_linux.go:199
			// _ = "end of CoverTab[14268]"
		}
//line /usr/local/go/src/net/interface_linux.go:200
		// _ = "end of CoverTab[14262]"
	}
//line /usr/local/go/src/net/interface_linux.go:201
	// _ = "end of CoverTab[14256]"
//line /usr/local/go/src/net/interface_linux.go:201
	_go_fuzz_dep_.CoverTab[14257]++
							return nil
//line /usr/local/go/src/net/interface_linux.go:202
	// _ = "end of CoverTab[14257]"
}

// interfaceMulticastAddrTable returns addresses for a specific
//line /usr/local/go/src/net/interface_linux.go:205
// interface.
//line /usr/local/go/src/net/interface_linux.go:207
func interfaceMulticastAddrTable(ifi *Interface) ([]Addr, error) {
//line /usr/local/go/src/net/interface_linux.go:207
	_go_fuzz_dep_.CoverTab[14269]++
							ifmat4 := parseProcNetIGMP("/proc/net/igmp", ifi)
							ifmat6 := parseProcNetIGMP6("/proc/net/igmp6", ifi)
							return append(ifmat4, ifmat6...), nil
//line /usr/local/go/src/net/interface_linux.go:210
	// _ = "end of CoverTab[14269]"
}

func parseProcNetIGMP(path string, ifi *Interface) []Addr {
//line /usr/local/go/src/net/interface_linux.go:213
	_go_fuzz_dep_.CoverTab[14270]++
							fd, err := open(path)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:215
		_go_fuzz_dep_.CoverTab[14273]++
								return nil
//line /usr/local/go/src/net/interface_linux.go:216
		// _ = "end of CoverTab[14273]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:217
		_go_fuzz_dep_.CoverTab[14274]++
//line /usr/local/go/src/net/interface_linux.go:217
		// _ = "end of CoverTab[14274]"
//line /usr/local/go/src/net/interface_linux.go:217
	}
//line /usr/local/go/src/net/interface_linux.go:217
	// _ = "end of CoverTab[14270]"
//line /usr/local/go/src/net/interface_linux.go:217
	_go_fuzz_dep_.CoverTab[14271]++
							defer fd.close()
							var (
		ifmat	[]Addr
		name	string
	)
	fd.readLine()
	b := make([]byte, IPv4len)
	for l, ok := fd.readLine(); ok; l, ok = fd.readLine() {
//line /usr/local/go/src/net/interface_linux.go:225
		_go_fuzz_dep_.CoverTab[14275]++
								f := splitAtBytes(l, " :\r\t\n")
								if len(f) < 4 {
//line /usr/local/go/src/net/interface_linux.go:227
			_go_fuzz_dep_.CoverTab[14277]++
									continue
//line /usr/local/go/src/net/interface_linux.go:228
			// _ = "end of CoverTab[14277]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:229
			_go_fuzz_dep_.CoverTab[14278]++
//line /usr/local/go/src/net/interface_linux.go:229
			// _ = "end of CoverTab[14278]"
//line /usr/local/go/src/net/interface_linux.go:229
		}
//line /usr/local/go/src/net/interface_linux.go:229
		// _ = "end of CoverTab[14275]"
//line /usr/local/go/src/net/interface_linux.go:229
		_go_fuzz_dep_.CoverTab[14276]++
								switch {
		case l[0] != ' ' && func() bool {
//line /usr/local/go/src/net/interface_linux.go:231
			_go_fuzz_dep_.CoverTab[14282]++
//line /usr/local/go/src/net/interface_linux.go:231
			return l[0] != '\t'
//line /usr/local/go/src/net/interface_linux.go:231
			// _ = "end of CoverTab[14282]"
//line /usr/local/go/src/net/interface_linux.go:231
		}():
//line /usr/local/go/src/net/interface_linux.go:231
			_go_fuzz_dep_.CoverTab[14279]++
									name = f[1]
//line /usr/local/go/src/net/interface_linux.go:232
			// _ = "end of CoverTab[14279]"
		case len(f[0]) == 8:
//line /usr/local/go/src/net/interface_linux.go:233
			_go_fuzz_dep_.CoverTab[14280]++
									if ifi == nil || func() bool {
//line /usr/local/go/src/net/interface_linux.go:234
				_go_fuzz_dep_.CoverTab[14283]++
//line /usr/local/go/src/net/interface_linux.go:234
				return name == ifi.Name
//line /usr/local/go/src/net/interface_linux.go:234
				// _ = "end of CoverTab[14283]"
//line /usr/local/go/src/net/interface_linux.go:234
			}() {
//line /usr/local/go/src/net/interface_linux.go:234
				_go_fuzz_dep_.CoverTab[14284]++

//line /usr/local/go/src/net/interface_linux.go:238
				for i := 0; i+1 < len(f[0]); i += 2 {
//line /usr/local/go/src/net/interface_linux.go:238
					_go_fuzz_dep_.CoverTab[14286]++
											b[i/2], _ = xtoi2(f[0][i:i+2], 0)
//line /usr/local/go/src/net/interface_linux.go:239
					// _ = "end of CoverTab[14286]"
				}
//line /usr/local/go/src/net/interface_linux.go:240
				// _ = "end of CoverTab[14284]"
//line /usr/local/go/src/net/interface_linux.go:240
				_go_fuzz_dep_.CoverTab[14285]++
										i := *(*uint32)(unsafe.Pointer(&b[:4][0]))
										ifma := &IPAddr{IP: IPv4(byte(i>>24), byte(i>>16), byte(i>>8), byte(i))}
										ifmat = append(ifmat, ifma)
//line /usr/local/go/src/net/interface_linux.go:243
				// _ = "end of CoverTab[14285]"
			} else {
//line /usr/local/go/src/net/interface_linux.go:244
				_go_fuzz_dep_.CoverTab[14287]++
//line /usr/local/go/src/net/interface_linux.go:244
				// _ = "end of CoverTab[14287]"
//line /usr/local/go/src/net/interface_linux.go:244
			}
//line /usr/local/go/src/net/interface_linux.go:244
			// _ = "end of CoverTab[14280]"
//line /usr/local/go/src/net/interface_linux.go:244
		default:
//line /usr/local/go/src/net/interface_linux.go:244
			_go_fuzz_dep_.CoverTab[14281]++
//line /usr/local/go/src/net/interface_linux.go:244
			// _ = "end of CoverTab[14281]"
		}
//line /usr/local/go/src/net/interface_linux.go:245
		// _ = "end of CoverTab[14276]"
	}
//line /usr/local/go/src/net/interface_linux.go:246
	// _ = "end of CoverTab[14271]"
//line /usr/local/go/src/net/interface_linux.go:246
	_go_fuzz_dep_.CoverTab[14272]++
							return ifmat
//line /usr/local/go/src/net/interface_linux.go:247
	// _ = "end of CoverTab[14272]"
}

func parseProcNetIGMP6(path string, ifi *Interface) []Addr {
//line /usr/local/go/src/net/interface_linux.go:250
	_go_fuzz_dep_.CoverTab[14288]++
							fd, err := open(path)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:252
		_go_fuzz_dep_.CoverTab[14291]++
								return nil
//line /usr/local/go/src/net/interface_linux.go:253
		// _ = "end of CoverTab[14291]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:254
		_go_fuzz_dep_.CoverTab[14292]++
//line /usr/local/go/src/net/interface_linux.go:254
		// _ = "end of CoverTab[14292]"
//line /usr/local/go/src/net/interface_linux.go:254
	}
//line /usr/local/go/src/net/interface_linux.go:254
	// _ = "end of CoverTab[14288]"
//line /usr/local/go/src/net/interface_linux.go:254
	_go_fuzz_dep_.CoverTab[14289]++
							defer fd.close()
							var ifmat []Addr
							b := make([]byte, IPv6len)
							for l, ok := fd.readLine(); ok; l, ok = fd.readLine() {
//line /usr/local/go/src/net/interface_linux.go:258
		_go_fuzz_dep_.CoverTab[14293]++
								f := splitAtBytes(l, " \r\t\n")
								if len(f) < 6 {
//line /usr/local/go/src/net/interface_linux.go:260
			_go_fuzz_dep_.CoverTab[14295]++
									continue
//line /usr/local/go/src/net/interface_linux.go:261
			// _ = "end of CoverTab[14295]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:262
			_go_fuzz_dep_.CoverTab[14296]++
//line /usr/local/go/src/net/interface_linux.go:262
			// _ = "end of CoverTab[14296]"
//line /usr/local/go/src/net/interface_linux.go:262
		}
//line /usr/local/go/src/net/interface_linux.go:262
		// _ = "end of CoverTab[14293]"
//line /usr/local/go/src/net/interface_linux.go:262
		_go_fuzz_dep_.CoverTab[14294]++
								if ifi == nil || func() bool {
//line /usr/local/go/src/net/interface_linux.go:263
			_go_fuzz_dep_.CoverTab[14297]++
//line /usr/local/go/src/net/interface_linux.go:263
			return f[1] == ifi.Name
//line /usr/local/go/src/net/interface_linux.go:263
			// _ = "end of CoverTab[14297]"
//line /usr/local/go/src/net/interface_linux.go:263
		}() {
//line /usr/local/go/src/net/interface_linux.go:263
			_go_fuzz_dep_.CoverTab[14298]++
									for i := 0; i+1 < len(f[2]); i += 2 {
//line /usr/local/go/src/net/interface_linux.go:264
				_go_fuzz_dep_.CoverTab[14300]++
										b[i/2], _ = xtoi2(f[2][i:i+2], 0)
//line /usr/local/go/src/net/interface_linux.go:265
				// _ = "end of CoverTab[14300]"
			}
//line /usr/local/go/src/net/interface_linux.go:266
			// _ = "end of CoverTab[14298]"
//line /usr/local/go/src/net/interface_linux.go:266
			_go_fuzz_dep_.CoverTab[14299]++
									ifma := &IPAddr{IP: IP{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15]}}
									ifmat = append(ifmat, ifma)
//line /usr/local/go/src/net/interface_linux.go:268
			// _ = "end of CoverTab[14299]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:269
			_go_fuzz_dep_.CoverTab[14301]++
//line /usr/local/go/src/net/interface_linux.go:269
			// _ = "end of CoverTab[14301]"
//line /usr/local/go/src/net/interface_linux.go:269
		}
//line /usr/local/go/src/net/interface_linux.go:269
		// _ = "end of CoverTab[14294]"
	}
//line /usr/local/go/src/net/interface_linux.go:270
	// _ = "end of CoverTab[14289]"
//line /usr/local/go/src/net/interface_linux.go:270
	_go_fuzz_dep_.CoverTab[14290]++
							return ifmat
//line /usr/local/go/src/net/interface_linux.go:271
	// _ = "end of CoverTab[14290]"
}

//line /usr/local/go/src/net/interface_linux.go:272
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/interface_linux.go:272
var _ = _go_fuzz_dep_.CoverTab
