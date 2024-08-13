// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/interface_linux.go:5
package net

//line /snap/go/10455/src/net/interface_linux.go:5
import (
//line /snap/go/10455/src/net/interface_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/interface_linux.go:5
)
//line /snap/go/10455/src/net/interface_linux.go:5
import (
//line /snap/go/10455/src/net/interface_linux.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/interface_linux.go:5
)

import (
	"os"
	"syscall"
	"unsafe"
)

// If the ifindex is zero, interfaceTable returns mappings of all
//line /snap/go/10455/src/net/interface_linux.go:13
// network interfaces. Otherwise it returns a mapping of a specific
//line /snap/go/10455/src/net/interface_linux.go:13
// interface.
//line /snap/go/10455/src/net/interface_linux.go:16
func interfaceTable(ifindex int) ([]Interface, error) {
//line /snap/go/10455/src/net/interface_linux.go:16
	_go_fuzz_dep_.CoverTab[6146]++
							tab, err := syscall.NetlinkRIB(syscall.RTM_GETLINK, syscall.AF_UNSPEC)
							if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:18
		_go_fuzz_dep_.CoverTab[528444]++
//line /snap/go/10455/src/net/interface_linux.go:18
		_go_fuzz_dep_.CoverTab[6150]++
								return nil, os.NewSyscallError("netlinkrib", err)
//line /snap/go/10455/src/net/interface_linux.go:19
		// _ = "end of CoverTab[6150]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:20
		_go_fuzz_dep_.CoverTab[528445]++
//line /snap/go/10455/src/net/interface_linux.go:20
		_go_fuzz_dep_.CoverTab[6151]++
//line /snap/go/10455/src/net/interface_linux.go:20
		// _ = "end of CoverTab[6151]"
//line /snap/go/10455/src/net/interface_linux.go:20
	}
//line /snap/go/10455/src/net/interface_linux.go:20
	// _ = "end of CoverTab[6146]"
//line /snap/go/10455/src/net/interface_linux.go:20
	_go_fuzz_dep_.CoverTab[6147]++
							msgs, err := syscall.ParseNetlinkMessage(tab)
							if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:22
		_go_fuzz_dep_.CoverTab[528446]++
//line /snap/go/10455/src/net/interface_linux.go:22
		_go_fuzz_dep_.CoverTab[6152]++
								return nil, os.NewSyscallError("parsenetlinkmessage", err)
//line /snap/go/10455/src/net/interface_linux.go:23
		// _ = "end of CoverTab[6152]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:24
		_go_fuzz_dep_.CoverTab[528447]++
//line /snap/go/10455/src/net/interface_linux.go:24
		_go_fuzz_dep_.CoverTab[6153]++
//line /snap/go/10455/src/net/interface_linux.go:24
		// _ = "end of CoverTab[6153]"
//line /snap/go/10455/src/net/interface_linux.go:24
	}
//line /snap/go/10455/src/net/interface_linux.go:24
	// _ = "end of CoverTab[6147]"
//line /snap/go/10455/src/net/interface_linux.go:24
	_go_fuzz_dep_.CoverTab[6148]++
							var ift []Interface
loop:
	for _, m := range msgs {
//line /snap/go/10455/src/net/interface_linux.go:27
		_go_fuzz_dep_.CoverTab[6154]++
								switch m.Header.Type {
		case syscall.NLMSG_DONE:
//line /snap/go/10455/src/net/interface_linux.go:29
			_go_fuzz_dep_.CoverTab[528448]++
//line /snap/go/10455/src/net/interface_linux.go:29
			_go_fuzz_dep_.CoverTab[6155]++
									break loop
//line /snap/go/10455/src/net/interface_linux.go:30
			// _ = "end of CoverTab[6155]"
		case syscall.RTM_NEWLINK:
//line /snap/go/10455/src/net/interface_linux.go:31
			_go_fuzz_dep_.CoverTab[528449]++
//line /snap/go/10455/src/net/interface_linux.go:31
			_go_fuzz_dep_.CoverTab[6156]++
									ifim := (*syscall.IfInfomsg)(unsafe.Pointer(&m.Data[0]))
									if ifindex == 0 || func() bool {
//line /snap/go/10455/src/net/interface_linux.go:33
				_go_fuzz_dep_.CoverTab[6158]++
//line /snap/go/10455/src/net/interface_linux.go:33
				return ifindex == int(ifim.Index)
//line /snap/go/10455/src/net/interface_linux.go:33
				// _ = "end of CoverTab[6158]"
//line /snap/go/10455/src/net/interface_linux.go:33
			}() {
//line /snap/go/10455/src/net/interface_linux.go:33
				_go_fuzz_dep_.CoverTab[528451]++
//line /snap/go/10455/src/net/interface_linux.go:33
				_go_fuzz_dep_.CoverTab[6159]++
										attrs, err := syscall.ParseNetlinkRouteAttr(&m)
										if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:35
					_go_fuzz_dep_.CoverTab[528453]++
//line /snap/go/10455/src/net/interface_linux.go:35
					_go_fuzz_dep_.CoverTab[6161]++
											return nil, os.NewSyscallError("parsenetlinkrouteattr", err)
//line /snap/go/10455/src/net/interface_linux.go:36
					// _ = "end of CoverTab[6161]"
				} else {
//line /snap/go/10455/src/net/interface_linux.go:37
					_go_fuzz_dep_.CoverTab[528454]++
//line /snap/go/10455/src/net/interface_linux.go:37
					_go_fuzz_dep_.CoverTab[6162]++
//line /snap/go/10455/src/net/interface_linux.go:37
					// _ = "end of CoverTab[6162]"
//line /snap/go/10455/src/net/interface_linux.go:37
				}
//line /snap/go/10455/src/net/interface_linux.go:37
				// _ = "end of CoverTab[6159]"
//line /snap/go/10455/src/net/interface_linux.go:37
				_go_fuzz_dep_.CoverTab[6160]++
										ift = append(ift, *newLink(ifim, attrs))
										if ifindex == int(ifim.Index) {
//line /snap/go/10455/src/net/interface_linux.go:39
					_go_fuzz_dep_.CoverTab[528455]++
//line /snap/go/10455/src/net/interface_linux.go:39
					_go_fuzz_dep_.CoverTab[6163]++
											break loop
//line /snap/go/10455/src/net/interface_linux.go:40
					// _ = "end of CoverTab[6163]"
				} else {
//line /snap/go/10455/src/net/interface_linux.go:41
					_go_fuzz_dep_.CoverTab[528456]++
//line /snap/go/10455/src/net/interface_linux.go:41
					_go_fuzz_dep_.CoverTab[6164]++
//line /snap/go/10455/src/net/interface_linux.go:41
					// _ = "end of CoverTab[6164]"
//line /snap/go/10455/src/net/interface_linux.go:41
				}
//line /snap/go/10455/src/net/interface_linux.go:41
				// _ = "end of CoverTab[6160]"
			} else {
//line /snap/go/10455/src/net/interface_linux.go:42
				_go_fuzz_dep_.CoverTab[528452]++
//line /snap/go/10455/src/net/interface_linux.go:42
				_go_fuzz_dep_.CoverTab[6165]++
//line /snap/go/10455/src/net/interface_linux.go:42
				// _ = "end of CoverTab[6165]"
//line /snap/go/10455/src/net/interface_linux.go:42
			}
//line /snap/go/10455/src/net/interface_linux.go:42
			// _ = "end of CoverTab[6156]"
//line /snap/go/10455/src/net/interface_linux.go:42
		default:
//line /snap/go/10455/src/net/interface_linux.go:42
			_go_fuzz_dep_.CoverTab[528450]++
//line /snap/go/10455/src/net/interface_linux.go:42
			_go_fuzz_dep_.CoverTab[6157]++
//line /snap/go/10455/src/net/interface_linux.go:42
			// _ = "end of CoverTab[6157]"
		}
//line /snap/go/10455/src/net/interface_linux.go:43
		// _ = "end of CoverTab[6154]"
	}
//line /snap/go/10455/src/net/interface_linux.go:44
	// _ = "end of CoverTab[6148]"
//line /snap/go/10455/src/net/interface_linux.go:44
	_go_fuzz_dep_.CoverTab[6149]++
							return ift, nil
//line /snap/go/10455/src/net/interface_linux.go:45
	// _ = "end of CoverTab[6149]"
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
//line /snap/go/10455/src/net/interface_linux.go:58
	_go_fuzz_dep_.CoverTab[6166]++
							ifi := &Interface{Index: int(ifim.Index), Flags: linkFlags(ifim.Flags)}
//line /snap/go/10455/src/net/interface_linux.go:59
	_go_fuzz_dep_.CoverTab[786681] = 0
							for _, a := range attrs {
//line /snap/go/10455/src/net/interface_linux.go:60
		if _go_fuzz_dep_.CoverTab[786681] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:60
			_go_fuzz_dep_.CoverTab[528529]++
//line /snap/go/10455/src/net/interface_linux.go:60
		} else {
//line /snap/go/10455/src/net/interface_linux.go:60
			_go_fuzz_dep_.CoverTab[528530]++
//line /snap/go/10455/src/net/interface_linux.go:60
		}
//line /snap/go/10455/src/net/interface_linux.go:60
		_go_fuzz_dep_.CoverTab[786681] = 1
//line /snap/go/10455/src/net/interface_linux.go:60
		_go_fuzz_dep_.CoverTab[6168]++
								switch a.Attr.Type {
		case syscall.IFLA_ADDRESS:
//line /snap/go/10455/src/net/interface_linux.go:62
			_go_fuzz_dep_.CoverTab[528457]++
//line /snap/go/10455/src/net/interface_linux.go:62
			_go_fuzz_dep_.CoverTab[6169]++

//line /snap/go/10455/src/net/interface_linux.go:66
			switch len(a.Value) {
			case IPv4len:
//line /snap/go/10455/src/net/interface_linux.go:67
				_go_fuzz_dep_.CoverTab[528461]++
//line /snap/go/10455/src/net/interface_linux.go:67
				_go_fuzz_dep_.CoverTab[6175]++
										switch ifim.Type {
				case sysARPHardwareIPv4IPv4, sysARPHardwareGREIPv4, sysARPHardwareIPv6IPv4:
//line /snap/go/10455/src/net/interface_linux.go:69
					_go_fuzz_dep_.CoverTab[528464]++
//line /snap/go/10455/src/net/interface_linux.go:69
					_go_fuzz_dep_.CoverTab[6178]++
											continue
//line /snap/go/10455/src/net/interface_linux.go:70
					// _ = "end of CoverTab[6178]"
//line /snap/go/10455/src/net/interface_linux.go:70
				default:
//line /snap/go/10455/src/net/interface_linux.go:70
					_go_fuzz_dep_.CoverTab[528465]++
//line /snap/go/10455/src/net/interface_linux.go:70
					_go_fuzz_dep_.CoverTab[6179]++
//line /snap/go/10455/src/net/interface_linux.go:70
					// _ = "end of CoverTab[6179]"
				}
//line /snap/go/10455/src/net/interface_linux.go:71
				// _ = "end of CoverTab[6175]"
			case IPv6len:
//line /snap/go/10455/src/net/interface_linux.go:72
				_go_fuzz_dep_.CoverTab[528462]++
//line /snap/go/10455/src/net/interface_linux.go:72
				_go_fuzz_dep_.CoverTab[6176]++
										switch ifim.Type {
				case sysARPHardwareIPv6IPv6, sysARPHardwareGREIPv6:
//line /snap/go/10455/src/net/interface_linux.go:74
					_go_fuzz_dep_.CoverTab[528466]++
//line /snap/go/10455/src/net/interface_linux.go:74
					_go_fuzz_dep_.CoverTab[6180]++
											continue
//line /snap/go/10455/src/net/interface_linux.go:75
					// _ = "end of CoverTab[6180]"
//line /snap/go/10455/src/net/interface_linux.go:75
				default:
//line /snap/go/10455/src/net/interface_linux.go:75
					_go_fuzz_dep_.CoverTab[528467]++
//line /snap/go/10455/src/net/interface_linux.go:75
					_go_fuzz_dep_.CoverTab[6181]++
//line /snap/go/10455/src/net/interface_linux.go:75
					// _ = "end of CoverTab[6181]"
				}
//line /snap/go/10455/src/net/interface_linux.go:76
				// _ = "end of CoverTab[6176]"
//line /snap/go/10455/src/net/interface_linux.go:76
			default:
//line /snap/go/10455/src/net/interface_linux.go:76
				_go_fuzz_dep_.CoverTab[528463]++
//line /snap/go/10455/src/net/interface_linux.go:76
				_go_fuzz_dep_.CoverTab[6177]++
//line /snap/go/10455/src/net/interface_linux.go:76
				// _ = "end of CoverTab[6177]"
			}
//line /snap/go/10455/src/net/interface_linux.go:77
			// _ = "end of CoverTab[6169]"
//line /snap/go/10455/src/net/interface_linux.go:77
			_go_fuzz_dep_.CoverTab[6170]++
									var nonzero bool
									for _, b := range a.Value {
//line /snap/go/10455/src/net/interface_linux.go:79
				_go_fuzz_dep_.CoverTab[6182]++
										if b != 0 {
//line /snap/go/10455/src/net/interface_linux.go:80
					_go_fuzz_dep_.CoverTab[528468]++
//line /snap/go/10455/src/net/interface_linux.go:80
					_go_fuzz_dep_.CoverTab[6183]++
											nonzero = true
											break
//line /snap/go/10455/src/net/interface_linux.go:82
					// _ = "end of CoverTab[6183]"
				} else {
//line /snap/go/10455/src/net/interface_linux.go:83
					_go_fuzz_dep_.CoverTab[528469]++
//line /snap/go/10455/src/net/interface_linux.go:83
					_go_fuzz_dep_.CoverTab[6184]++
//line /snap/go/10455/src/net/interface_linux.go:83
					// _ = "end of CoverTab[6184]"
//line /snap/go/10455/src/net/interface_linux.go:83
				}
//line /snap/go/10455/src/net/interface_linux.go:83
				// _ = "end of CoverTab[6182]"
			}
//line /snap/go/10455/src/net/interface_linux.go:84
			// _ = "end of CoverTab[6170]"
//line /snap/go/10455/src/net/interface_linux.go:84
			_go_fuzz_dep_.CoverTab[6171]++
									if nonzero {
//line /snap/go/10455/src/net/interface_linux.go:85
				_go_fuzz_dep_.CoverTab[528470]++
//line /snap/go/10455/src/net/interface_linux.go:85
				_go_fuzz_dep_.CoverTab[6185]++
										ifi.HardwareAddr = a.Value[:]
//line /snap/go/10455/src/net/interface_linux.go:86
				// _ = "end of CoverTab[6185]"
			} else {
//line /snap/go/10455/src/net/interface_linux.go:87
				_go_fuzz_dep_.CoverTab[528471]++
//line /snap/go/10455/src/net/interface_linux.go:87
				_go_fuzz_dep_.CoverTab[6186]++
//line /snap/go/10455/src/net/interface_linux.go:87
				// _ = "end of CoverTab[6186]"
//line /snap/go/10455/src/net/interface_linux.go:87
			}
//line /snap/go/10455/src/net/interface_linux.go:87
			// _ = "end of CoverTab[6171]"
		case syscall.IFLA_IFNAME:
//line /snap/go/10455/src/net/interface_linux.go:88
			_go_fuzz_dep_.CoverTab[528458]++
//line /snap/go/10455/src/net/interface_linux.go:88
			_go_fuzz_dep_.CoverTab[6172]++
									ifi.Name = string(a.Value[:len(a.Value)-1])
//line /snap/go/10455/src/net/interface_linux.go:89
			// _ = "end of CoverTab[6172]"
		case syscall.IFLA_MTU:
//line /snap/go/10455/src/net/interface_linux.go:90
			_go_fuzz_dep_.CoverTab[528459]++
//line /snap/go/10455/src/net/interface_linux.go:90
			_go_fuzz_dep_.CoverTab[6173]++
									ifi.MTU = int(*(*uint32)(unsafe.Pointer(&a.Value[:4][0])))
//line /snap/go/10455/src/net/interface_linux.go:91
			// _ = "end of CoverTab[6173]"
//line /snap/go/10455/src/net/interface_linux.go:91
		default:
//line /snap/go/10455/src/net/interface_linux.go:91
			_go_fuzz_dep_.CoverTab[528460]++
//line /snap/go/10455/src/net/interface_linux.go:91
			_go_fuzz_dep_.CoverTab[6174]++
//line /snap/go/10455/src/net/interface_linux.go:91
			// _ = "end of CoverTab[6174]"
		}
//line /snap/go/10455/src/net/interface_linux.go:92
		// _ = "end of CoverTab[6168]"
	}
//line /snap/go/10455/src/net/interface_linux.go:93
	if _go_fuzz_dep_.CoverTab[786681] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:93
		_go_fuzz_dep_.CoverTab[528531]++
//line /snap/go/10455/src/net/interface_linux.go:93
	} else {
//line /snap/go/10455/src/net/interface_linux.go:93
		_go_fuzz_dep_.CoverTab[528532]++
//line /snap/go/10455/src/net/interface_linux.go:93
	}
//line /snap/go/10455/src/net/interface_linux.go:93
	// _ = "end of CoverTab[6166]"
//line /snap/go/10455/src/net/interface_linux.go:93
	_go_fuzz_dep_.CoverTab[6167]++
							return ifi
//line /snap/go/10455/src/net/interface_linux.go:94
	// _ = "end of CoverTab[6167]"
}

func linkFlags(rawFlags uint32) Flags {
//line /snap/go/10455/src/net/interface_linux.go:97
	_go_fuzz_dep_.CoverTab[6187]++
							var f Flags
							if rawFlags&syscall.IFF_UP != 0 {
//line /snap/go/10455/src/net/interface_linux.go:99
		_go_fuzz_dep_.CoverTab[528472]++
//line /snap/go/10455/src/net/interface_linux.go:99
		_go_fuzz_dep_.CoverTab[6194]++
								f |= FlagUp
//line /snap/go/10455/src/net/interface_linux.go:100
		// _ = "end of CoverTab[6194]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:101
		_go_fuzz_dep_.CoverTab[528473]++
//line /snap/go/10455/src/net/interface_linux.go:101
		_go_fuzz_dep_.CoverTab[6195]++
//line /snap/go/10455/src/net/interface_linux.go:101
		// _ = "end of CoverTab[6195]"
//line /snap/go/10455/src/net/interface_linux.go:101
	}
//line /snap/go/10455/src/net/interface_linux.go:101
	// _ = "end of CoverTab[6187]"
//line /snap/go/10455/src/net/interface_linux.go:101
	_go_fuzz_dep_.CoverTab[6188]++
							if rawFlags&syscall.IFF_RUNNING != 0 {
//line /snap/go/10455/src/net/interface_linux.go:102
		_go_fuzz_dep_.CoverTab[528474]++
//line /snap/go/10455/src/net/interface_linux.go:102
		_go_fuzz_dep_.CoverTab[6196]++
								f |= FlagRunning
//line /snap/go/10455/src/net/interface_linux.go:103
		// _ = "end of CoverTab[6196]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:104
		_go_fuzz_dep_.CoverTab[528475]++
//line /snap/go/10455/src/net/interface_linux.go:104
		_go_fuzz_dep_.CoverTab[6197]++
//line /snap/go/10455/src/net/interface_linux.go:104
		// _ = "end of CoverTab[6197]"
//line /snap/go/10455/src/net/interface_linux.go:104
	}
//line /snap/go/10455/src/net/interface_linux.go:104
	// _ = "end of CoverTab[6188]"
//line /snap/go/10455/src/net/interface_linux.go:104
	_go_fuzz_dep_.CoverTab[6189]++
							if rawFlags&syscall.IFF_BROADCAST != 0 {
//line /snap/go/10455/src/net/interface_linux.go:105
		_go_fuzz_dep_.CoverTab[528476]++
//line /snap/go/10455/src/net/interface_linux.go:105
		_go_fuzz_dep_.CoverTab[6198]++
								f |= FlagBroadcast
//line /snap/go/10455/src/net/interface_linux.go:106
		// _ = "end of CoverTab[6198]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:107
		_go_fuzz_dep_.CoverTab[528477]++
//line /snap/go/10455/src/net/interface_linux.go:107
		_go_fuzz_dep_.CoverTab[6199]++
//line /snap/go/10455/src/net/interface_linux.go:107
		// _ = "end of CoverTab[6199]"
//line /snap/go/10455/src/net/interface_linux.go:107
	}
//line /snap/go/10455/src/net/interface_linux.go:107
	// _ = "end of CoverTab[6189]"
//line /snap/go/10455/src/net/interface_linux.go:107
	_go_fuzz_dep_.CoverTab[6190]++
							if rawFlags&syscall.IFF_LOOPBACK != 0 {
//line /snap/go/10455/src/net/interface_linux.go:108
		_go_fuzz_dep_.CoverTab[528478]++
//line /snap/go/10455/src/net/interface_linux.go:108
		_go_fuzz_dep_.CoverTab[6200]++
								f |= FlagLoopback
//line /snap/go/10455/src/net/interface_linux.go:109
		// _ = "end of CoverTab[6200]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:110
		_go_fuzz_dep_.CoverTab[528479]++
//line /snap/go/10455/src/net/interface_linux.go:110
		_go_fuzz_dep_.CoverTab[6201]++
//line /snap/go/10455/src/net/interface_linux.go:110
		// _ = "end of CoverTab[6201]"
//line /snap/go/10455/src/net/interface_linux.go:110
	}
//line /snap/go/10455/src/net/interface_linux.go:110
	// _ = "end of CoverTab[6190]"
//line /snap/go/10455/src/net/interface_linux.go:110
	_go_fuzz_dep_.CoverTab[6191]++
							if rawFlags&syscall.IFF_POINTOPOINT != 0 {
//line /snap/go/10455/src/net/interface_linux.go:111
		_go_fuzz_dep_.CoverTab[528480]++
//line /snap/go/10455/src/net/interface_linux.go:111
		_go_fuzz_dep_.CoverTab[6202]++
								f |= FlagPointToPoint
//line /snap/go/10455/src/net/interface_linux.go:112
		// _ = "end of CoverTab[6202]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:113
		_go_fuzz_dep_.CoverTab[528481]++
//line /snap/go/10455/src/net/interface_linux.go:113
		_go_fuzz_dep_.CoverTab[6203]++
//line /snap/go/10455/src/net/interface_linux.go:113
		// _ = "end of CoverTab[6203]"
//line /snap/go/10455/src/net/interface_linux.go:113
	}
//line /snap/go/10455/src/net/interface_linux.go:113
	// _ = "end of CoverTab[6191]"
//line /snap/go/10455/src/net/interface_linux.go:113
	_go_fuzz_dep_.CoverTab[6192]++
							if rawFlags&syscall.IFF_MULTICAST != 0 {
//line /snap/go/10455/src/net/interface_linux.go:114
		_go_fuzz_dep_.CoverTab[528482]++
//line /snap/go/10455/src/net/interface_linux.go:114
		_go_fuzz_dep_.CoverTab[6204]++
								f |= FlagMulticast
//line /snap/go/10455/src/net/interface_linux.go:115
		// _ = "end of CoverTab[6204]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:116
		_go_fuzz_dep_.CoverTab[528483]++
//line /snap/go/10455/src/net/interface_linux.go:116
		_go_fuzz_dep_.CoverTab[6205]++
//line /snap/go/10455/src/net/interface_linux.go:116
		// _ = "end of CoverTab[6205]"
//line /snap/go/10455/src/net/interface_linux.go:116
	}
//line /snap/go/10455/src/net/interface_linux.go:116
	// _ = "end of CoverTab[6192]"
//line /snap/go/10455/src/net/interface_linux.go:116
	_go_fuzz_dep_.CoverTab[6193]++
							return f
//line /snap/go/10455/src/net/interface_linux.go:117
	// _ = "end of CoverTab[6193]"
}

// If the ifi is nil, interfaceAddrTable returns addresses for all
//line /snap/go/10455/src/net/interface_linux.go:120
// network interfaces. Otherwise it returns addresses for a specific
//line /snap/go/10455/src/net/interface_linux.go:120
// interface.
//line /snap/go/10455/src/net/interface_linux.go:123
func interfaceAddrTable(ifi *Interface) ([]Addr, error) {
//line /snap/go/10455/src/net/interface_linux.go:123
	_go_fuzz_dep_.CoverTab[6206]++
							tab, err := syscall.NetlinkRIB(syscall.RTM_GETADDR, syscall.AF_UNSPEC)
							if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:125
		_go_fuzz_dep_.CoverTab[528484]++
//line /snap/go/10455/src/net/interface_linux.go:125
		_go_fuzz_dep_.CoverTab[6211]++
								return nil, os.NewSyscallError("netlinkrib", err)
//line /snap/go/10455/src/net/interface_linux.go:126
		// _ = "end of CoverTab[6211]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:127
		_go_fuzz_dep_.CoverTab[528485]++
//line /snap/go/10455/src/net/interface_linux.go:127
		_go_fuzz_dep_.CoverTab[6212]++
//line /snap/go/10455/src/net/interface_linux.go:127
		// _ = "end of CoverTab[6212]"
//line /snap/go/10455/src/net/interface_linux.go:127
	}
//line /snap/go/10455/src/net/interface_linux.go:127
	// _ = "end of CoverTab[6206]"
//line /snap/go/10455/src/net/interface_linux.go:127
	_go_fuzz_dep_.CoverTab[6207]++
							msgs, err := syscall.ParseNetlinkMessage(tab)
							if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:129
		_go_fuzz_dep_.CoverTab[528486]++
//line /snap/go/10455/src/net/interface_linux.go:129
		_go_fuzz_dep_.CoverTab[6213]++
								return nil, os.NewSyscallError("parsenetlinkmessage", err)
//line /snap/go/10455/src/net/interface_linux.go:130
		// _ = "end of CoverTab[6213]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:131
		_go_fuzz_dep_.CoverTab[528487]++
//line /snap/go/10455/src/net/interface_linux.go:131
		_go_fuzz_dep_.CoverTab[6214]++
//line /snap/go/10455/src/net/interface_linux.go:131
		// _ = "end of CoverTab[6214]"
//line /snap/go/10455/src/net/interface_linux.go:131
	}
//line /snap/go/10455/src/net/interface_linux.go:131
	// _ = "end of CoverTab[6207]"
//line /snap/go/10455/src/net/interface_linux.go:131
	_go_fuzz_dep_.CoverTab[6208]++
							var ift []Interface
							if ifi == nil {
//line /snap/go/10455/src/net/interface_linux.go:133
		_go_fuzz_dep_.CoverTab[528488]++
//line /snap/go/10455/src/net/interface_linux.go:133
		_go_fuzz_dep_.CoverTab[6215]++
								var err error
								ift, err = interfaceTable(0)
								if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:136
			_go_fuzz_dep_.CoverTab[528490]++
//line /snap/go/10455/src/net/interface_linux.go:136
			_go_fuzz_dep_.CoverTab[6216]++
									return nil, err
//line /snap/go/10455/src/net/interface_linux.go:137
			// _ = "end of CoverTab[6216]"
		} else {
//line /snap/go/10455/src/net/interface_linux.go:138
			_go_fuzz_dep_.CoverTab[528491]++
//line /snap/go/10455/src/net/interface_linux.go:138
			_go_fuzz_dep_.CoverTab[6217]++
//line /snap/go/10455/src/net/interface_linux.go:138
			// _ = "end of CoverTab[6217]"
//line /snap/go/10455/src/net/interface_linux.go:138
		}
//line /snap/go/10455/src/net/interface_linux.go:138
		// _ = "end of CoverTab[6215]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:139
		_go_fuzz_dep_.CoverTab[528489]++
//line /snap/go/10455/src/net/interface_linux.go:139
		_go_fuzz_dep_.CoverTab[6218]++
//line /snap/go/10455/src/net/interface_linux.go:139
		// _ = "end of CoverTab[6218]"
//line /snap/go/10455/src/net/interface_linux.go:139
	}
//line /snap/go/10455/src/net/interface_linux.go:139
	// _ = "end of CoverTab[6208]"
//line /snap/go/10455/src/net/interface_linux.go:139
	_go_fuzz_dep_.CoverTab[6209]++
							ifat, err := addrTable(ift, ifi, msgs)
							if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:141
		_go_fuzz_dep_.CoverTab[528492]++
//line /snap/go/10455/src/net/interface_linux.go:141
		_go_fuzz_dep_.CoverTab[6219]++
								return nil, err
//line /snap/go/10455/src/net/interface_linux.go:142
		// _ = "end of CoverTab[6219]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:143
		_go_fuzz_dep_.CoverTab[528493]++
//line /snap/go/10455/src/net/interface_linux.go:143
		_go_fuzz_dep_.CoverTab[6220]++
//line /snap/go/10455/src/net/interface_linux.go:143
		// _ = "end of CoverTab[6220]"
//line /snap/go/10455/src/net/interface_linux.go:143
	}
//line /snap/go/10455/src/net/interface_linux.go:143
	// _ = "end of CoverTab[6209]"
//line /snap/go/10455/src/net/interface_linux.go:143
	_go_fuzz_dep_.CoverTab[6210]++
							return ifat, nil
//line /snap/go/10455/src/net/interface_linux.go:144
	// _ = "end of CoverTab[6210]"
}

func addrTable(ift []Interface, ifi *Interface, msgs []syscall.NetlinkMessage) ([]Addr, error) {
//line /snap/go/10455/src/net/interface_linux.go:147
	_go_fuzz_dep_.CoverTab[6221]++
							var ifat []Addr
loop:
	for _, m := range msgs {
//line /snap/go/10455/src/net/interface_linux.go:150
		_go_fuzz_dep_.CoverTab[6223]++
								switch m.Header.Type {
		case syscall.NLMSG_DONE:
//line /snap/go/10455/src/net/interface_linux.go:152
			_go_fuzz_dep_.CoverTab[528494]++
//line /snap/go/10455/src/net/interface_linux.go:152
			_go_fuzz_dep_.CoverTab[6224]++
									break loop
//line /snap/go/10455/src/net/interface_linux.go:153
			// _ = "end of CoverTab[6224]"
		case syscall.RTM_NEWADDR:
//line /snap/go/10455/src/net/interface_linux.go:154
			_go_fuzz_dep_.CoverTab[528495]++
//line /snap/go/10455/src/net/interface_linux.go:154
			_go_fuzz_dep_.CoverTab[6225]++
									ifam := (*syscall.IfAddrmsg)(unsafe.Pointer(&m.Data[0]))
									if len(ift) != 0 || func() bool {
//line /snap/go/10455/src/net/interface_linux.go:156
				_go_fuzz_dep_.CoverTab[6227]++
//line /snap/go/10455/src/net/interface_linux.go:156
				return ifi.Index == int(ifam.Index)
//line /snap/go/10455/src/net/interface_linux.go:156
				// _ = "end of CoverTab[6227]"
//line /snap/go/10455/src/net/interface_linux.go:156
			}() {
//line /snap/go/10455/src/net/interface_linux.go:156
				_go_fuzz_dep_.CoverTab[528497]++
//line /snap/go/10455/src/net/interface_linux.go:156
				_go_fuzz_dep_.CoverTab[6228]++
										if len(ift) != 0 {
//line /snap/go/10455/src/net/interface_linux.go:157
					_go_fuzz_dep_.CoverTab[528499]++
//line /snap/go/10455/src/net/interface_linux.go:157
					_go_fuzz_dep_.CoverTab[6231]++
											var err error
											ifi, err = interfaceByIndex(ift, int(ifam.Index))
											if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:160
						_go_fuzz_dep_.CoverTab[528501]++
//line /snap/go/10455/src/net/interface_linux.go:160
						_go_fuzz_dep_.CoverTab[6232]++
												return nil, err
//line /snap/go/10455/src/net/interface_linux.go:161
						// _ = "end of CoverTab[6232]"
					} else {
//line /snap/go/10455/src/net/interface_linux.go:162
						_go_fuzz_dep_.CoverTab[528502]++
//line /snap/go/10455/src/net/interface_linux.go:162
						_go_fuzz_dep_.CoverTab[6233]++
//line /snap/go/10455/src/net/interface_linux.go:162
						// _ = "end of CoverTab[6233]"
//line /snap/go/10455/src/net/interface_linux.go:162
					}
//line /snap/go/10455/src/net/interface_linux.go:162
					// _ = "end of CoverTab[6231]"
				} else {
//line /snap/go/10455/src/net/interface_linux.go:163
					_go_fuzz_dep_.CoverTab[528500]++
//line /snap/go/10455/src/net/interface_linux.go:163
					_go_fuzz_dep_.CoverTab[6234]++
//line /snap/go/10455/src/net/interface_linux.go:163
					// _ = "end of CoverTab[6234]"
//line /snap/go/10455/src/net/interface_linux.go:163
				}
//line /snap/go/10455/src/net/interface_linux.go:163
				// _ = "end of CoverTab[6228]"
//line /snap/go/10455/src/net/interface_linux.go:163
				_go_fuzz_dep_.CoverTab[6229]++
										attrs, err := syscall.ParseNetlinkRouteAttr(&m)
										if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:165
					_go_fuzz_dep_.CoverTab[528503]++
//line /snap/go/10455/src/net/interface_linux.go:165
					_go_fuzz_dep_.CoverTab[6235]++
											return nil, os.NewSyscallError("parsenetlinkrouteattr", err)
//line /snap/go/10455/src/net/interface_linux.go:166
					// _ = "end of CoverTab[6235]"
				} else {
//line /snap/go/10455/src/net/interface_linux.go:167
					_go_fuzz_dep_.CoverTab[528504]++
//line /snap/go/10455/src/net/interface_linux.go:167
					_go_fuzz_dep_.CoverTab[6236]++
//line /snap/go/10455/src/net/interface_linux.go:167
					// _ = "end of CoverTab[6236]"
//line /snap/go/10455/src/net/interface_linux.go:167
				}
//line /snap/go/10455/src/net/interface_linux.go:167
				// _ = "end of CoverTab[6229]"
//line /snap/go/10455/src/net/interface_linux.go:167
				_go_fuzz_dep_.CoverTab[6230]++
										ifa := newAddr(ifam, attrs)
										if ifa != nil {
//line /snap/go/10455/src/net/interface_linux.go:169
					_go_fuzz_dep_.CoverTab[528505]++
//line /snap/go/10455/src/net/interface_linux.go:169
					_go_fuzz_dep_.CoverTab[6237]++
											ifat = append(ifat, ifa)
//line /snap/go/10455/src/net/interface_linux.go:170
					// _ = "end of CoverTab[6237]"
				} else {
//line /snap/go/10455/src/net/interface_linux.go:171
					_go_fuzz_dep_.CoverTab[528506]++
//line /snap/go/10455/src/net/interface_linux.go:171
					_go_fuzz_dep_.CoverTab[6238]++
//line /snap/go/10455/src/net/interface_linux.go:171
					// _ = "end of CoverTab[6238]"
//line /snap/go/10455/src/net/interface_linux.go:171
				}
//line /snap/go/10455/src/net/interface_linux.go:171
				// _ = "end of CoverTab[6230]"
			} else {
//line /snap/go/10455/src/net/interface_linux.go:172
				_go_fuzz_dep_.CoverTab[528498]++
//line /snap/go/10455/src/net/interface_linux.go:172
				_go_fuzz_dep_.CoverTab[6239]++
//line /snap/go/10455/src/net/interface_linux.go:172
				// _ = "end of CoverTab[6239]"
//line /snap/go/10455/src/net/interface_linux.go:172
			}
//line /snap/go/10455/src/net/interface_linux.go:172
			// _ = "end of CoverTab[6225]"
//line /snap/go/10455/src/net/interface_linux.go:172
		default:
//line /snap/go/10455/src/net/interface_linux.go:172
			_go_fuzz_dep_.CoverTab[528496]++
//line /snap/go/10455/src/net/interface_linux.go:172
			_go_fuzz_dep_.CoverTab[6226]++
//line /snap/go/10455/src/net/interface_linux.go:172
			// _ = "end of CoverTab[6226]"
		}
//line /snap/go/10455/src/net/interface_linux.go:173
		// _ = "end of CoverTab[6223]"
	}
//line /snap/go/10455/src/net/interface_linux.go:174
	// _ = "end of CoverTab[6221]"
//line /snap/go/10455/src/net/interface_linux.go:174
	_go_fuzz_dep_.CoverTab[6222]++
							return ifat, nil
//line /snap/go/10455/src/net/interface_linux.go:175
	// _ = "end of CoverTab[6222]"
}

func newAddr(ifam *syscall.IfAddrmsg, attrs []syscall.NetlinkRouteAttr) Addr {
//line /snap/go/10455/src/net/interface_linux.go:178
	_go_fuzz_dep_.CoverTab[6240]++
							var ipPointToPoint bool
//line /snap/go/10455/src/net/interface_linux.go:179
	_go_fuzz_dep_.CoverTab[786682] = 0

//line /snap/go/10455/src/net/interface_linux.go:183
	for _, a := range attrs {
//line /snap/go/10455/src/net/interface_linux.go:183
		if _go_fuzz_dep_.CoverTab[786682] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:183
			_go_fuzz_dep_.CoverTab[528533]++
//line /snap/go/10455/src/net/interface_linux.go:183
		} else {
//line /snap/go/10455/src/net/interface_linux.go:183
			_go_fuzz_dep_.CoverTab[528534]++
//line /snap/go/10455/src/net/interface_linux.go:183
		}
//line /snap/go/10455/src/net/interface_linux.go:183
		_go_fuzz_dep_.CoverTab[786682] = 1
//line /snap/go/10455/src/net/interface_linux.go:183
		_go_fuzz_dep_.CoverTab[6243]++
								if a.Attr.Type == syscall.IFA_LOCAL {
//line /snap/go/10455/src/net/interface_linux.go:184
			_go_fuzz_dep_.CoverTab[528507]++
//line /snap/go/10455/src/net/interface_linux.go:184
			_go_fuzz_dep_.CoverTab[6244]++
									ipPointToPoint = true
									break
//line /snap/go/10455/src/net/interface_linux.go:186
			// _ = "end of CoverTab[6244]"
		} else {
//line /snap/go/10455/src/net/interface_linux.go:187
			_go_fuzz_dep_.CoverTab[528508]++
//line /snap/go/10455/src/net/interface_linux.go:187
			_go_fuzz_dep_.CoverTab[6245]++
//line /snap/go/10455/src/net/interface_linux.go:187
			// _ = "end of CoverTab[6245]"
//line /snap/go/10455/src/net/interface_linux.go:187
		}
//line /snap/go/10455/src/net/interface_linux.go:187
		// _ = "end of CoverTab[6243]"
	}
//line /snap/go/10455/src/net/interface_linux.go:188
	if _go_fuzz_dep_.CoverTab[786682] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:188
		_go_fuzz_dep_.CoverTab[528535]++
//line /snap/go/10455/src/net/interface_linux.go:188
	} else {
//line /snap/go/10455/src/net/interface_linux.go:188
		_go_fuzz_dep_.CoverTab[528536]++
//line /snap/go/10455/src/net/interface_linux.go:188
	}
//line /snap/go/10455/src/net/interface_linux.go:188
	// _ = "end of CoverTab[6240]"
//line /snap/go/10455/src/net/interface_linux.go:188
	_go_fuzz_dep_.CoverTab[6241]++
//line /snap/go/10455/src/net/interface_linux.go:188
	_go_fuzz_dep_.CoverTab[786683] = 0
							for _, a := range attrs {
//line /snap/go/10455/src/net/interface_linux.go:189
		if _go_fuzz_dep_.CoverTab[786683] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:189
			_go_fuzz_dep_.CoverTab[528537]++
//line /snap/go/10455/src/net/interface_linux.go:189
		} else {
//line /snap/go/10455/src/net/interface_linux.go:189
			_go_fuzz_dep_.CoverTab[528538]++
//line /snap/go/10455/src/net/interface_linux.go:189
		}
//line /snap/go/10455/src/net/interface_linux.go:189
		_go_fuzz_dep_.CoverTab[786683] = 1
//line /snap/go/10455/src/net/interface_linux.go:189
		_go_fuzz_dep_.CoverTab[6246]++
								if ipPointToPoint && func() bool {
//line /snap/go/10455/src/net/interface_linux.go:190
			_go_fuzz_dep_.CoverTab[6248]++
//line /snap/go/10455/src/net/interface_linux.go:190
			return a.Attr.Type == syscall.IFA_ADDRESS
//line /snap/go/10455/src/net/interface_linux.go:190
			// _ = "end of CoverTab[6248]"
//line /snap/go/10455/src/net/interface_linux.go:190
		}() {
//line /snap/go/10455/src/net/interface_linux.go:190
			_go_fuzz_dep_.CoverTab[528509]++
//line /snap/go/10455/src/net/interface_linux.go:190
			_go_fuzz_dep_.CoverTab[6249]++
									continue
//line /snap/go/10455/src/net/interface_linux.go:191
			// _ = "end of CoverTab[6249]"
		} else {
//line /snap/go/10455/src/net/interface_linux.go:192
			_go_fuzz_dep_.CoverTab[528510]++
//line /snap/go/10455/src/net/interface_linux.go:192
			_go_fuzz_dep_.CoverTab[6250]++
//line /snap/go/10455/src/net/interface_linux.go:192
			// _ = "end of CoverTab[6250]"
//line /snap/go/10455/src/net/interface_linux.go:192
		}
//line /snap/go/10455/src/net/interface_linux.go:192
		// _ = "end of CoverTab[6246]"
//line /snap/go/10455/src/net/interface_linux.go:192
		_go_fuzz_dep_.CoverTab[6247]++
								switch ifam.Family {
		case syscall.AF_INET:
//line /snap/go/10455/src/net/interface_linux.go:194
			_go_fuzz_dep_.CoverTab[528511]++
//line /snap/go/10455/src/net/interface_linux.go:194
			_go_fuzz_dep_.CoverTab[6251]++
									return &IPNet{IP: IPv4(a.Value[0], a.Value[1], a.Value[2], a.Value[3]), Mask: CIDRMask(int(ifam.Prefixlen), 8*IPv4len)}
//line /snap/go/10455/src/net/interface_linux.go:195
			// _ = "end of CoverTab[6251]"
		case syscall.AF_INET6:
//line /snap/go/10455/src/net/interface_linux.go:196
			_go_fuzz_dep_.CoverTab[528512]++
//line /snap/go/10455/src/net/interface_linux.go:196
			_go_fuzz_dep_.CoverTab[6252]++
									ifa := &IPNet{IP: make(IP, IPv6len), Mask: CIDRMask(int(ifam.Prefixlen), 8*IPv6len)}
									copy(ifa.IP, a.Value[:])
									return ifa
//line /snap/go/10455/src/net/interface_linux.go:199
			// _ = "end of CoverTab[6252]"
//line /snap/go/10455/src/net/interface_linux.go:199
		default:
//line /snap/go/10455/src/net/interface_linux.go:199
			_go_fuzz_dep_.CoverTab[528513]++
//line /snap/go/10455/src/net/interface_linux.go:199
			_go_fuzz_dep_.CoverTab[6253]++
//line /snap/go/10455/src/net/interface_linux.go:199
			// _ = "end of CoverTab[6253]"
		}
//line /snap/go/10455/src/net/interface_linux.go:200
		// _ = "end of CoverTab[6247]"
	}
//line /snap/go/10455/src/net/interface_linux.go:201
	if _go_fuzz_dep_.CoverTab[786683] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:201
		_go_fuzz_dep_.CoverTab[528539]++
//line /snap/go/10455/src/net/interface_linux.go:201
	} else {
//line /snap/go/10455/src/net/interface_linux.go:201
		_go_fuzz_dep_.CoverTab[528540]++
//line /snap/go/10455/src/net/interface_linux.go:201
	}
//line /snap/go/10455/src/net/interface_linux.go:201
	// _ = "end of CoverTab[6241]"
//line /snap/go/10455/src/net/interface_linux.go:201
	_go_fuzz_dep_.CoverTab[6242]++
							return nil
//line /snap/go/10455/src/net/interface_linux.go:202
	// _ = "end of CoverTab[6242]"
}

// interfaceMulticastAddrTable returns addresses for a specific
//line /snap/go/10455/src/net/interface_linux.go:205
// interface.
//line /snap/go/10455/src/net/interface_linux.go:207
func interfaceMulticastAddrTable(ifi *Interface) ([]Addr, error) {
//line /snap/go/10455/src/net/interface_linux.go:207
	_go_fuzz_dep_.CoverTab[6254]++
							ifmat4 := parseProcNetIGMP("/proc/net/igmp", ifi)
							ifmat6 := parseProcNetIGMP6("/proc/net/igmp6", ifi)
							return append(ifmat4, ifmat6...), nil
//line /snap/go/10455/src/net/interface_linux.go:210
	// _ = "end of CoverTab[6254]"
}

func parseProcNetIGMP(path string, ifi *Interface) []Addr {
//line /snap/go/10455/src/net/interface_linux.go:213
	_go_fuzz_dep_.CoverTab[6255]++
							fd, err := open(path)
							if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:215
		_go_fuzz_dep_.CoverTab[528514]++
//line /snap/go/10455/src/net/interface_linux.go:215
		_go_fuzz_dep_.CoverTab[6258]++
								return nil
//line /snap/go/10455/src/net/interface_linux.go:216
		// _ = "end of CoverTab[6258]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:217
		_go_fuzz_dep_.CoverTab[528515]++
//line /snap/go/10455/src/net/interface_linux.go:217
		_go_fuzz_dep_.CoverTab[6259]++
//line /snap/go/10455/src/net/interface_linux.go:217
		// _ = "end of CoverTab[6259]"
//line /snap/go/10455/src/net/interface_linux.go:217
	}
//line /snap/go/10455/src/net/interface_linux.go:217
	// _ = "end of CoverTab[6255]"
//line /snap/go/10455/src/net/interface_linux.go:217
	_go_fuzz_dep_.CoverTab[6256]++
							defer fd.close()
							var (
		ifmat	[]Addr
		name	string
	)
							fd.readLine()
							b := make([]byte, IPv4len)
//line /snap/go/10455/src/net/interface_linux.go:224
	_go_fuzz_dep_.CoverTab[786684] = 0
							for l, ok := fd.readLine(); ok; l, ok = fd.readLine() {
//line /snap/go/10455/src/net/interface_linux.go:225
		if _go_fuzz_dep_.CoverTab[786684] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:225
			_go_fuzz_dep_.CoverTab[528541]++
//line /snap/go/10455/src/net/interface_linux.go:225
		} else {
//line /snap/go/10455/src/net/interface_linux.go:225
			_go_fuzz_dep_.CoverTab[528542]++
//line /snap/go/10455/src/net/interface_linux.go:225
		}
//line /snap/go/10455/src/net/interface_linux.go:225
		_go_fuzz_dep_.CoverTab[786684] = 1
//line /snap/go/10455/src/net/interface_linux.go:225
		_go_fuzz_dep_.CoverTab[6260]++
								f := splitAtBytes(l, " :\r\t\n")
								if len(f) < 4 {
//line /snap/go/10455/src/net/interface_linux.go:227
			_go_fuzz_dep_.CoverTab[528516]++
//line /snap/go/10455/src/net/interface_linux.go:227
			_go_fuzz_dep_.CoverTab[6262]++
									continue
//line /snap/go/10455/src/net/interface_linux.go:228
			// _ = "end of CoverTab[6262]"
		} else {
//line /snap/go/10455/src/net/interface_linux.go:229
			_go_fuzz_dep_.CoverTab[528517]++
//line /snap/go/10455/src/net/interface_linux.go:229
			_go_fuzz_dep_.CoverTab[6263]++
//line /snap/go/10455/src/net/interface_linux.go:229
			// _ = "end of CoverTab[6263]"
//line /snap/go/10455/src/net/interface_linux.go:229
		}
//line /snap/go/10455/src/net/interface_linux.go:229
		// _ = "end of CoverTab[6260]"
//line /snap/go/10455/src/net/interface_linux.go:229
		_go_fuzz_dep_.CoverTab[6261]++
								switch {
		case l[0] != ' ' && func() bool {
//line /snap/go/10455/src/net/interface_linux.go:231
			_go_fuzz_dep_.CoverTab[6267]++
//line /snap/go/10455/src/net/interface_linux.go:231
			return l[0] != '\t'
//line /snap/go/10455/src/net/interface_linux.go:231
			// _ = "end of CoverTab[6267]"
//line /snap/go/10455/src/net/interface_linux.go:231
		}():
//line /snap/go/10455/src/net/interface_linux.go:231
			_go_fuzz_dep_.CoverTab[528518]++
//line /snap/go/10455/src/net/interface_linux.go:231
			_go_fuzz_dep_.CoverTab[6264]++
									name = f[1]
//line /snap/go/10455/src/net/interface_linux.go:232
			// _ = "end of CoverTab[6264]"
		case len(f[0]) == 8:
//line /snap/go/10455/src/net/interface_linux.go:233
			_go_fuzz_dep_.CoverTab[528519]++
//line /snap/go/10455/src/net/interface_linux.go:233
			_go_fuzz_dep_.CoverTab[6265]++
									if ifi == nil || func() bool {
//line /snap/go/10455/src/net/interface_linux.go:234
				_go_fuzz_dep_.CoverTab[6268]++
//line /snap/go/10455/src/net/interface_linux.go:234
				return name == ifi.Name
//line /snap/go/10455/src/net/interface_linux.go:234
				// _ = "end of CoverTab[6268]"
//line /snap/go/10455/src/net/interface_linux.go:234
			}() {
//line /snap/go/10455/src/net/interface_linux.go:234
				_go_fuzz_dep_.CoverTab[528521]++
//line /snap/go/10455/src/net/interface_linux.go:234
				_go_fuzz_dep_.CoverTab[6269]++
//line /snap/go/10455/src/net/interface_linux.go:234
				_go_fuzz_dep_.CoverTab[786685] = 0

//line /snap/go/10455/src/net/interface_linux.go:238
				for i := 0; i+1 < len(f[0]); i += 2 {
//line /snap/go/10455/src/net/interface_linux.go:238
					if _go_fuzz_dep_.CoverTab[786685] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:238
						_go_fuzz_dep_.CoverTab[528545]++
//line /snap/go/10455/src/net/interface_linux.go:238
					} else {
//line /snap/go/10455/src/net/interface_linux.go:238
						_go_fuzz_dep_.CoverTab[528546]++
//line /snap/go/10455/src/net/interface_linux.go:238
					}
//line /snap/go/10455/src/net/interface_linux.go:238
					_go_fuzz_dep_.CoverTab[786685] = 1
//line /snap/go/10455/src/net/interface_linux.go:238
					_go_fuzz_dep_.CoverTab[6271]++
											b[i/2], _ = xtoi2(f[0][i:i+2], 0)
//line /snap/go/10455/src/net/interface_linux.go:239
					// _ = "end of CoverTab[6271]"
				}
//line /snap/go/10455/src/net/interface_linux.go:240
				if _go_fuzz_dep_.CoverTab[786685] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:240
					_go_fuzz_dep_.CoverTab[528547]++
//line /snap/go/10455/src/net/interface_linux.go:240
				} else {
//line /snap/go/10455/src/net/interface_linux.go:240
					_go_fuzz_dep_.CoverTab[528548]++
//line /snap/go/10455/src/net/interface_linux.go:240
				}
//line /snap/go/10455/src/net/interface_linux.go:240
				// _ = "end of CoverTab[6269]"
//line /snap/go/10455/src/net/interface_linux.go:240
				_go_fuzz_dep_.CoverTab[6270]++
										i := *(*uint32)(unsafe.Pointer(&b[:4][0]))
										ifma := &IPAddr{IP: IPv4(byte(i>>24), byte(i>>16), byte(i>>8), byte(i))}
										ifmat = append(ifmat, ifma)
//line /snap/go/10455/src/net/interface_linux.go:243
				// _ = "end of CoverTab[6270]"
			} else {
//line /snap/go/10455/src/net/interface_linux.go:244
				_go_fuzz_dep_.CoverTab[528522]++
//line /snap/go/10455/src/net/interface_linux.go:244
				_go_fuzz_dep_.CoverTab[6272]++
//line /snap/go/10455/src/net/interface_linux.go:244
				// _ = "end of CoverTab[6272]"
//line /snap/go/10455/src/net/interface_linux.go:244
			}
//line /snap/go/10455/src/net/interface_linux.go:244
			// _ = "end of CoverTab[6265]"
//line /snap/go/10455/src/net/interface_linux.go:244
		default:
//line /snap/go/10455/src/net/interface_linux.go:244
			_go_fuzz_dep_.CoverTab[528520]++
//line /snap/go/10455/src/net/interface_linux.go:244
			_go_fuzz_dep_.CoverTab[6266]++
//line /snap/go/10455/src/net/interface_linux.go:244
			// _ = "end of CoverTab[6266]"
		}
//line /snap/go/10455/src/net/interface_linux.go:245
		// _ = "end of CoverTab[6261]"
	}
//line /snap/go/10455/src/net/interface_linux.go:246
	if _go_fuzz_dep_.CoverTab[786684] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:246
		_go_fuzz_dep_.CoverTab[528543]++
//line /snap/go/10455/src/net/interface_linux.go:246
	} else {
//line /snap/go/10455/src/net/interface_linux.go:246
		_go_fuzz_dep_.CoverTab[528544]++
//line /snap/go/10455/src/net/interface_linux.go:246
	}
//line /snap/go/10455/src/net/interface_linux.go:246
	// _ = "end of CoverTab[6256]"
//line /snap/go/10455/src/net/interface_linux.go:246
	_go_fuzz_dep_.CoverTab[6257]++
							return ifmat
//line /snap/go/10455/src/net/interface_linux.go:247
	// _ = "end of CoverTab[6257]"
}

func parseProcNetIGMP6(path string, ifi *Interface) []Addr {
//line /snap/go/10455/src/net/interface_linux.go:250
	_go_fuzz_dep_.CoverTab[6273]++
							fd, err := open(path)
							if err != nil {
//line /snap/go/10455/src/net/interface_linux.go:252
		_go_fuzz_dep_.CoverTab[528523]++
//line /snap/go/10455/src/net/interface_linux.go:252
		_go_fuzz_dep_.CoverTab[6276]++
								return nil
//line /snap/go/10455/src/net/interface_linux.go:253
		// _ = "end of CoverTab[6276]"
	} else {
//line /snap/go/10455/src/net/interface_linux.go:254
		_go_fuzz_dep_.CoverTab[528524]++
//line /snap/go/10455/src/net/interface_linux.go:254
		_go_fuzz_dep_.CoverTab[6277]++
//line /snap/go/10455/src/net/interface_linux.go:254
		// _ = "end of CoverTab[6277]"
//line /snap/go/10455/src/net/interface_linux.go:254
	}
//line /snap/go/10455/src/net/interface_linux.go:254
	// _ = "end of CoverTab[6273]"
//line /snap/go/10455/src/net/interface_linux.go:254
	_go_fuzz_dep_.CoverTab[6274]++
							defer fd.close()
							var ifmat []Addr
							b := make([]byte, IPv6len)
//line /snap/go/10455/src/net/interface_linux.go:257
	_go_fuzz_dep_.CoverTab[786686] = 0
							for l, ok := fd.readLine(); ok; l, ok = fd.readLine() {
//line /snap/go/10455/src/net/interface_linux.go:258
		if _go_fuzz_dep_.CoverTab[786686] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:258
			_go_fuzz_dep_.CoverTab[528549]++
//line /snap/go/10455/src/net/interface_linux.go:258
		} else {
//line /snap/go/10455/src/net/interface_linux.go:258
			_go_fuzz_dep_.CoverTab[528550]++
//line /snap/go/10455/src/net/interface_linux.go:258
		}
//line /snap/go/10455/src/net/interface_linux.go:258
		_go_fuzz_dep_.CoverTab[786686] = 1
//line /snap/go/10455/src/net/interface_linux.go:258
		_go_fuzz_dep_.CoverTab[6278]++
								f := splitAtBytes(l, " \r\t\n")
								if len(f) < 6 {
//line /snap/go/10455/src/net/interface_linux.go:260
			_go_fuzz_dep_.CoverTab[528525]++
//line /snap/go/10455/src/net/interface_linux.go:260
			_go_fuzz_dep_.CoverTab[6280]++
									continue
//line /snap/go/10455/src/net/interface_linux.go:261
			// _ = "end of CoverTab[6280]"
		} else {
//line /snap/go/10455/src/net/interface_linux.go:262
			_go_fuzz_dep_.CoverTab[528526]++
//line /snap/go/10455/src/net/interface_linux.go:262
			_go_fuzz_dep_.CoverTab[6281]++
//line /snap/go/10455/src/net/interface_linux.go:262
			// _ = "end of CoverTab[6281]"
//line /snap/go/10455/src/net/interface_linux.go:262
		}
//line /snap/go/10455/src/net/interface_linux.go:262
		// _ = "end of CoverTab[6278]"
//line /snap/go/10455/src/net/interface_linux.go:262
		_go_fuzz_dep_.CoverTab[6279]++
								if ifi == nil || func() bool {
//line /snap/go/10455/src/net/interface_linux.go:263
			_go_fuzz_dep_.CoverTab[6282]++
//line /snap/go/10455/src/net/interface_linux.go:263
			return f[1] == ifi.Name
//line /snap/go/10455/src/net/interface_linux.go:263
			// _ = "end of CoverTab[6282]"
//line /snap/go/10455/src/net/interface_linux.go:263
		}() {
//line /snap/go/10455/src/net/interface_linux.go:263
			_go_fuzz_dep_.CoverTab[528527]++
//line /snap/go/10455/src/net/interface_linux.go:263
			_go_fuzz_dep_.CoverTab[6283]++
//line /snap/go/10455/src/net/interface_linux.go:263
			_go_fuzz_dep_.CoverTab[786687] = 0
									for i := 0; i+1 < len(f[2]); i += 2 {
//line /snap/go/10455/src/net/interface_linux.go:264
				if _go_fuzz_dep_.CoverTab[786687] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:264
					_go_fuzz_dep_.CoverTab[528553]++
//line /snap/go/10455/src/net/interface_linux.go:264
				} else {
//line /snap/go/10455/src/net/interface_linux.go:264
					_go_fuzz_dep_.CoverTab[528554]++
//line /snap/go/10455/src/net/interface_linux.go:264
				}
//line /snap/go/10455/src/net/interface_linux.go:264
				_go_fuzz_dep_.CoverTab[786687] = 1
//line /snap/go/10455/src/net/interface_linux.go:264
				_go_fuzz_dep_.CoverTab[6285]++
										b[i/2], _ = xtoi2(f[2][i:i+2], 0)
//line /snap/go/10455/src/net/interface_linux.go:265
				// _ = "end of CoverTab[6285]"
			}
//line /snap/go/10455/src/net/interface_linux.go:266
			if _go_fuzz_dep_.CoverTab[786687] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:266
				_go_fuzz_dep_.CoverTab[528555]++
//line /snap/go/10455/src/net/interface_linux.go:266
			} else {
//line /snap/go/10455/src/net/interface_linux.go:266
				_go_fuzz_dep_.CoverTab[528556]++
//line /snap/go/10455/src/net/interface_linux.go:266
			}
//line /snap/go/10455/src/net/interface_linux.go:266
			// _ = "end of CoverTab[6283]"
//line /snap/go/10455/src/net/interface_linux.go:266
			_go_fuzz_dep_.CoverTab[6284]++
									ifma := &IPAddr{IP: IP{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15]}}
									ifmat = append(ifmat, ifma)
//line /snap/go/10455/src/net/interface_linux.go:268
			// _ = "end of CoverTab[6284]"
		} else {
//line /snap/go/10455/src/net/interface_linux.go:269
			_go_fuzz_dep_.CoverTab[528528]++
//line /snap/go/10455/src/net/interface_linux.go:269
			_go_fuzz_dep_.CoverTab[6286]++
//line /snap/go/10455/src/net/interface_linux.go:269
			// _ = "end of CoverTab[6286]"
//line /snap/go/10455/src/net/interface_linux.go:269
		}
//line /snap/go/10455/src/net/interface_linux.go:269
		// _ = "end of CoverTab[6279]"
	}
//line /snap/go/10455/src/net/interface_linux.go:270
	if _go_fuzz_dep_.CoverTab[786686] == 0 {
//line /snap/go/10455/src/net/interface_linux.go:270
		_go_fuzz_dep_.CoverTab[528551]++
//line /snap/go/10455/src/net/interface_linux.go:270
	} else {
//line /snap/go/10455/src/net/interface_linux.go:270
		_go_fuzz_dep_.CoverTab[528552]++
//line /snap/go/10455/src/net/interface_linux.go:270
	}
//line /snap/go/10455/src/net/interface_linux.go:270
	// _ = "end of CoverTab[6274]"
//line /snap/go/10455/src/net/interface_linux.go:270
	_go_fuzz_dep_.CoverTab[6275]++
							return ifmat
//line /snap/go/10455/src/net/interface_linux.go:271
	// _ = "end of CoverTab[6275]"
}

//line /snap/go/10455/src/net/interface_linux.go:272
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/interface_linux.go:272
var _ = _go_fuzz_dep_.CoverTab
