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
	_go_fuzz_dep_.CoverTab[5771]++
							tab, err := syscall.NetlinkRIB(syscall.RTM_GETLINK, syscall.AF_UNSPEC)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:18
		_go_fuzz_dep_.CoverTab[5775]++
								return nil, os.NewSyscallError("netlinkrib", err)
//line /usr/local/go/src/net/interface_linux.go:19
		// _ = "end of CoverTab[5775]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:20
		_go_fuzz_dep_.CoverTab[5776]++
//line /usr/local/go/src/net/interface_linux.go:20
		// _ = "end of CoverTab[5776]"
//line /usr/local/go/src/net/interface_linux.go:20
	}
//line /usr/local/go/src/net/interface_linux.go:20
	// _ = "end of CoverTab[5771]"
//line /usr/local/go/src/net/interface_linux.go:20
	_go_fuzz_dep_.CoverTab[5772]++
							msgs, err := syscall.ParseNetlinkMessage(tab)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:22
		_go_fuzz_dep_.CoverTab[5777]++
								return nil, os.NewSyscallError("parsenetlinkmessage", err)
//line /usr/local/go/src/net/interface_linux.go:23
		// _ = "end of CoverTab[5777]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:24
		_go_fuzz_dep_.CoverTab[5778]++
//line /usr/local/go/src/net/interface_linux.go:24
		// _ = "end of CoverTab[5778]"
//line /usr/local/go/src/net/interface_linux.go:24
	}
//line /usr/local/go/src/net/interface_linux.go:24
	// _ = "end of CoverTab[5772]"
//line /usr/local/go/src/net/interface_linux.go:24
	_go_fuzz_dep_.CoverTab[5773]++
							var ift []Interface
loop:
	for _, m := range msgs {
//line /usr/local/go/src/net/interface_linux.go:27
		_go_fuzz_dep_.CoverTab[5779]++
								switch m.Header.Type {
		case syscall.NLMSG_DONE:
//line /usr/local/go/src/net/interface_linux.go:29
			_go_fuzz_dep_.CoverTab[5780]++
									break loop
//line /usr/local/go/src/net/interface_linux.go:30
			// _ = "end of CoverTab[5780]"
		case syscall.RTM_NEWLINK:
//line /usr/local/go/src/net/interface_linux.go:31
			_go_fuzz_dep_.CoverTab[5781]++
									ifim := (*syscall.IfInfomsg)(unsafe.Pointer(&m.Data[0]))
									if ifindex == 0 || func() bool {
//line /usr/local/go/src/net/interface_linux.go:33
				_go_fuzz_dep_.CoverTab[5783]++
//line /usr/local/go/src/net/interface_linux.go:33
				return ifindex == int(ifim.Index)
//line /usr/local/go/src/net/interface_linux.go:33
				// _ = "end of CoverTab[5783]"
//line /usr/local/go/src/net/interface_linux.go:33
			}() {
//line /usr/local/go/src/net/interface_linux.go:33
				_go_fuzz_dep_.CoverTab[5784]++
										attrs, err := syscall.ParseNetlinkRouteAttr(&m)
										if err != nil {
//line /usr/local/go/src/net/interface_linux.go:35
					_go_fuzz_dep_.CoverTab[5786]++
											return nil, os.NewSyscallError("parsenetlinkrouteattr", err)
//line /usr/local/go/src/net/interface_linux.go:36
					// _ = "end of CoverTab[5786]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:37
					_go_fuzz_dep_.CoverTab[5787]++
//line /usr/local/go/src/net/interface_linux.go:37
					// _ = "end of CoverTab[5787]"
//line /usr/local/go/src/net/interface_linux.go:37
				}
//line /usr/local/go/src/net/interface_linux.go:37
				// _ = "end of CoverTab[5784]"
//line /usr/local/go/src/net/interface_linux.go:37
				_go_fuzz_dep_.CoverTab[5785]++
										ift = append(ift, *newLink(ifim, attrs))
										if ifindex == int(ifim.Index) {
//line /usr/local/go/src/net/interface_linux.go:39
					_go_fuzz_dep_.CoverTab[5788]++
											break loop
//line /usr/local/go/src/net/interface_linux.go:40
					// _ = "end of CoverTab[5788]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:41
					_go_fuzz_dep_.CoverTab[5789]++
//line /usr/local/go/src/net/interface_linux.go:41
					// _ = "end of CoverTab[5789]"
//line /usr/local/go/src/net/interface_linux.go:41
				}
//line /usr/local/go/src/net/interface_linux.go:41
				// _ = "end of CoverTab[5785]"
			} else {
//line /usr/local/go/src/net/interface_linux.go:42
				_go_fuzz_dep_.CoverTab[5790]++
//line /usr/local/go/src/net/interface_linux.go:42
				// _ = "end of CoverTab[5790]"
//line /usr/local/go/src/net/interface_linux.go:42
			}
//line /usr/local/go/src/net/interface_linux.go:42
			// _ = "end of CoverTab[5781]"
//line /usr/local/go/src/net/interface_linux.go:42
		default:
//line /usr/local/go/src/net/interface_linux.go:42
			_go_fuzz_dep_.CoverTab[5782]++
//line /usr/local/go/src/net/interface_linux.go:42
			// _ = "end of CoverTab[5782]"
		}
//line /usr/local/go/src/net/interface_linux.go:43
		// _ = "end of CoverTab[5779]"
	}
//line /usr/local/go/src/net/interface_linux.go:44
	// _ = "end of CoverTab[5773]"
//line /usr/local/go/src/net/interface_linux.go:44
	_go_fuzz_dep_.CoverTab[5774]++
							return ift, nil
//line /usr/local/go/src/net/interface_linux.go:45
	// _ = "end of CoverTab[5774]"
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
	_go_fuzz_dep_.CoverTab[5791]++
							ifi := &Interface{Index: int(ifim.Index), Flags: linkFlags(ifim.Flags)}
							for _, a := range attrs {
//line /usr/local/go/src/net/interface_linux.go:60
		_go_fuzz_dep_.CoverTab[5793]++
								switch a.Attr.Type {
		case syscall.IFLA_ADDRESS:
//line /usr/local/go/src/net/interface_linux.go:62
			_go_fuzz_dep_.CoverTab[5794]++

//line /usr/local/go/src/net/interface_linux.go:66
			switch len(a.Value) {
			case IPv4len:
//line /usr/local/go/src/net/interface_linux.go:67
				_go_fuzz_dep_.CoverTab[5800]++
										switch ifim.Type {
				case sysARPHardwareIPv4IPv4, sysARPHardwareGREIPv4, sysARPHardwareIPv6IPv4:
//line /usr/local/go/src/net/interface_linux.go:69
					_go_fuzz_dep_.CoverTab[5803]++
											continue
//line /usr/local/go/src/net/interface_linux.go:70
					// _ = "end of CoverTab[5803]"
//line /usr/local/go/src/net/interface_linux.go:70
				default:
//line /usr/local/go/src/net/interface_linux.go:70
					_go_fuzz_dep_.CoverTab[5804]++
//line /usr/local/go/src/net/interface_linux.go:70
					// _ = "end of CoverTab[5804]"
				}
//line /usr/local/go/src/net/interface_linux.go:71
				// _ = "end of CoverTab[5800]"
			case IPv6len:
//line /usr/local/go/src/net/interface_linux.go:72
				_go_fuzz_dep_.CoverTab[5801]++
										switch ifim.Type {
				case sysARPHardwareIPv6IPv6, sysARPHardwareGREIPv6:
//line /usr/local/go/src/net/interface_linux.go:74
					_go_fuzz_dep_.CoverTab[5805]++
											continue
//line /usr/local/go/src/net/interface_linux.go:75
					// _ = "end of CoverTab[5805]"
//line /usr/local/go/src/net/interface_linux.go:75
				default:
//line /usr/local/go/src/net/interface_linux.go:75
					_go_fuzz_dep_.CoverTab[5806]++
//line /usr/local/go/src/net/interface_linux.go:75
					// _ = "end of CoverTab[5806]"
				}
//line /usr/local/go/src/net/interface_linux.go:76
				// _ = "end of CoverTab[5801]"
//line /usr/local/go/src/net/interface_linux.go:76
			default:
//line /usr/local/go/src/net/interface_linux.go:76
				_go_fuzz_dep_.CoverTab[5802]++
//line /usr/local/go/src/net/interface_linux.go:76
				// _ = "end of CoverTab[5802]"
			}
//line /usr/local/go/src/net/interface_linux.go:77
			// _ = "end of CoverTab[5794]"
//line /usr/local/go/src/net/interface_linux.go:77
			_go_fuzz_dep_.CoverTab[5795]++
									var nonzero bool
									for _, b := range a.Value {
//line /usr/local/go/src/net/interface_linux.go:79
				_go_fuzz_dep_.CoverTab[5807]++
										if b != 0 {
//line /usr/local/go/src/net/interface_linux.go:80
					_go_fuzz_dep_.CoverTab[5808]++
											nonzero = true
											break
//line /usr/local/go/src/net/interface_linux.go:82
					// _ = "end of CoverTab[5808]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:83
					_go_fuzz_dep_.CoverTab[5809]++
//line /usr/local/go/src/net/interface_linux.go:83
					// _ = "end of CoverTab[5809]"
//line /usr/local/go/src/net/interface_linux.go:83
				}
//line /usr/local/go/src/net/interface_linux.go:83
				// _ = "end of CoverTab[5807]"
			}
//line /usr/local/go/src/net/interface_linux.go:84
			// _ = "end of CoverTab[5795]"
//line /usr/local/go/src/net/interface_linux.go:84
			_go_fuzz_dep_.CoverTab[5796]++
									if nonzero {
//line /usr/local/go/src/net/interface_linux.go:85
				_go_fuzz_dep_.CoverTab[5810]++
										ifi.HardwareAddr = a.Value[:]
//line /usr/local/go/src/net/interface_linux.go:86
				// _ = "end of CoverTab[5810]"
			} else {
//line /usr/local/go/src/net/interface_linux.go:87
				_go_fuzz_dep_.CoverTab[5811]++
//line /usr/local/go/src/net/interface_linux.go:87
				// _ = "end of CoverTab[5811]"
//line /usr/local/go/src/net/interface_linux.go:87
			}
//line /usr/local/go/src/net/interface_linux.go:87
			// _ = "end of CoverTab[5796]"
		case syscall.IFLA_IFNAME:
//line /usr/local/go/src/net/interface_linux.go:88
			_go_fuzz_dep_.CoverTab[5797]++
									ifi.Name = string(a.Value[:len(a.Value)-1])
//line /usr/local/go/src/net/interface_linux.go:89
			// _ = "end of CoverTab[5797]"
		case syscall.IFLA_MTU:
//line /usr/local/go/src/net/interface_linux.go:90
			_go_fuzz_dep_.CoverTab[5798]++
									ifi.MTU = int(*(*uint32)(unsafe.Pointer(&a.Value[:4][0])))
//line /usr/local/go/src/net/interface_linux.go:91
			// _ = "end of CoverTab[5798]"
//line /usr/local/go/src/net/interface_linux.go:91
		default:
//line /usr/local/go/src/net/interface_linux.go:91
			_go_fuzz_dep_.CoverTab[5799]++
//line /usr/local/go/src/net/interface_linux.go:91
			// _ = "end of CoverTab[5799]"
		}
//line /usr/local/go/src/net/interface_linux.go:92
		// _ = "end of CoverTab[5793]"
	}
//line /usr/local/go/src/net/interface_linux.go:93
	// _ = "end of CoverTab[5791]"
//line /usr/local/go/src/net/interface_linux.go:93
	_go_fuzz_dep_.CoverTab[5792]++
							return ifi
//line /usr/local/go/src/net/interface_linux.go:94
	// _ = "end of CoverTab[5792]"
}

func linkFlags(rawFlags uint32) Flags {
//line /usr/local/go/src/net/interface_linux.go:97
	_go_fuzz_dep_.CoverTab[5812]++
							var f Flags
							if rawFlags&syscall.IFF_UP != 0 {
//line /usr/local/go/src/net/interface_linux.go:99
		_go_fuzz_dep_.CoverTab[5819]++
								f |= FlagUp
//line /usr/local/go/src/net/interface_linux.go:100
		// _ = "end of CoverTab[5819]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:101
		_go_fuzz_dep_.CoverTab[5820]++
//line /usr/local/go/src/net/interface_linux.go:101
		// _ = "end of CoverTab[5820]"
//line /usr/local/go/src/net/interface_linux.go:101
	}
//line /usr/local/go/src/net/interface_linux.go:101
	// _ = "end of CoverTab[5812]"
//line /usr/local/go/src/net/interface_linux.go:101
	_go_fuzz_dep_.CoverTab[5813]++
							if rawFlags&syscall.IFF_RUNNING != 0 {
//line /usr/local/go/src/net/interface_linux.go:102
		_go_fuzz_dep_.CoverTab[5821]++
								f |= FlagRunning
//line /usr/local/go/src/net/interface_linux.go:103
		// _ = "end of CoverTab[5821]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:104
		_go_fuzz_dep_.CoverTab[5822]++
//line /usr/local/go/src/net/interface_linux.go:104
		// _ = "end of CoverTab[5822]"
//line /usr/local/go/src/net/interface_linux.go:104
	}
//line /usr/local/go/src/net/interface_linux.go:104
	// _ = "end of CoverTab[5813]"
//line /usr/local/go/src/net/interface_linux.go:104
	_go_fuzz_dep_.CoverTab[5814]++
							if rawFlags&syscall.IFF_BROADCAST != 0 {
//line /usr/local/go/src/net/interface_linux.go:105
		_go_fuzz_dep_.CoverTab[5823]++
								f |= FlagBroadcast
//line /usr/local/go/src/net/interface_linux.go:106
		// _ = "end of CoverTab[5823]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:107
		_go_fuzz_dep_.CoverTab[5824]++
//line /usr/local/go/src/net/interface_linux.go:107
		// _ = "end of CoverTab[5824]"
//line /usr/local/go/src/net/interface_linux.go:107
	}
//line /usr/local/go/src/net/interface_linux.go:107
	// _ = "end of CoverTab[5814]"
//line /usr/local/go/src/net/interface_linux.go:107
	_go_fuzz_dep_.CoverTab[5815]++
							if rawFlags&syscall.IFF_LOOPBACK != 0 {
//line /usr/local/go/src/net/interface_linux.go:108
		_go_fuzz_dep_.CoverTab[5825]++
								f |= FlagLoopback
//line /usr/local/go/src/net/interface_linux.go:109
		// _ = "end of CoverTab[5825]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:110
		_go_fuzz_dep_.CoverTab[5826]++
//line /usr/local/go/src/net/interface_linux.go:110
		// _ = "end of CoverTab[5826]"
//line /usr/local/go/src/net/interface_linux.go:110
	}
//line /usr/local/go/src/net/interface_linux.go:110
	// _ = "end of CoverTab[5815]"
//line /usr/local/go/src/net/interface_linux.go:110
	_go_fuzz_dep_.CoverTab[5816]++
							if rawFlags&syscall.IFF_POINTOPOINT != 0 {
//line /usr/local/go/src/net/interface_linux.go:111
		_go_fuzz_dep_.CoverTab[5827]++
								f |= FlagPointToPoint
//line /usr/local/go/src/net/interface_linux.go:112
		// _ = "end of CoverTab[5827]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:113
		_go_fuzz_dep_.CoverTab[5828]++
//line /usr/local/go/src/net/interface_linux.go:113
		// _ = "end of CoverTab[5828]"
//line /usr/local/go/src/net/interface_linux.go:113
	}
//line /usr/local/go/src/net/interface_linux.go:113
	// _ = "end of CoverTab[5816]"
//line /usr/local/go/src/net/interface_linux.go:113
	_go_fuzz_dep_.CoverTab[5817]++
							if rawFlags&syscall.IFF_MULTICAST != 0 {
//line /usr/local/go/src/net/interface_linux.go:114
		_go_fuzz_dep_.CoverTab[5829]++
								f |= FlagMulticast
//line /usr/local/go/src/net/interface_linux.go:115
		// _ = "end of CoverTab[5829]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:116
		_go_fuzz_dep_.CoverTab[5830]++
//line /usr/local/go/src/net/interface_linux.go:116
		// _ = "end of CoverTab[5830]"
//line /usr/local/go/src/net/interface_linux.go:116
	}
//line /usr/local/go/src/net/interface_linux.go:116
	// _ = "end of CoverTab[5817]"
//line /usr/local/go/src/net/interface_linux.go:116
	_go_fuzz_dep_.CoverTab[5818]++
							return f
//line /usr/local/go/src/net/interface_linux.go:117
	// _ = "end of CoverTab[5818]"
}

// If the ifi is nil, interfaceAddrTable returns addresses for all
//line /usr/local/go/src/net/interface_linux.go:120
// network interfaces. Otherwise it returns addresses for a specific
//line /usr/local/go/src/net/interface_linux.go:120
// interface.
//line /usr/local/go/src/net/interface_linux.go:123
func interfaceAddrTable(ifi *Interface) ([]Addr, error) {
//line /usr/local/go/src/net/interface_linux.go:123
	_go_fuzz_dep_.CoverTab[5831]++
							tab, err := syscall.NetlinkRIB(syscall.RTM_GETADDR, syscall.AF_UNSPEC)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:125
		_go_fuzz_dep_.CoverTab[5836]++
								return nil, os.NewSyscallError("netlinkrib", err)
//line /usr/local/go/src/net/interface_linux.go:126
		// _ = "end of CoverTab[5836]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:127
		_go_fuzz_dep_.CoverTab[5837]++
//line /usr/local/go/src/net/interface_linux.go:127
		// _ = "end of CoverTab[5837]"
//line /usr/local/go/src/net/interface_linux.go:127
	}
//line /usr/local/go/src/net/interface_linux.go:127
	// _ = "end of CoverTab[5831]"
//line /usr/local/go/src/net/interface_linux.go:127
	_go_fuzz_dep_.CoverTab[5832]++
							msgs, err := syscall.ParseNetlinkMessage(tab)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:129
		_go_fuzz_dep_.CoverTab[5838]++
								return nil, os.NewSyscallError("parsenetlinkmessage", err)
//line /usr/local/go/src/net/interface_linux.go:130
		// _ = "end of CoverTab[5838]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:131
		_go_fuzz_dep_.CoverTab[5839]++
//line /usr/local/go/src/net/interface_linux.go:131
		// _ = "end of CoverTab[5839]"
//line /usr/local/go/src/net/interface_linux.go:131
	}
//line /usr/local/go/src/net/interface_linux.go:131
	// _ = "end of CoverTab[5832]"
//line /usr/local/go/src/net/interface_linux.go:131
	_go_fuzz_dep_.CoverTab[5833]++
							var ift []Interface
							if ifi == nil {
//line /usr/local/go/src/net/interface_linux.go:133
		_go_fuzz_dep_.CoverTab[5840]++
								var err error
								ift, err = interfaceTable(0)
								if err != nil {
//line /usr/local/go/src/net/interface_linux.go:136
			_go_fuzz_dep_.CoverTab[5841]++
									return nil, err
//line /usr/local/go/src/net/interface_linux.go:137
			// _ = "end of CoverTab[5841]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:138
			_go_fuzz_dep_.CoverTab[5842]++
//line /usr/local/go/src/net/interface_linux.go:138
			// _ = "end of CoverTab[5842]"
//line /usr/local/go/src/net/interface_linux.go:138
		}
//line /usr/local/go/src/net/interface_linux.go:138
		// _ = "end of CoverTab[5840]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:139
		_go_fuzz_dep_.CoverTab[5843]++
//line /usr/local/go/src/net/interface_linux.go:139
		// _ = "end of CoverTab[5843]"
//line /usr/local/go/src/net/interface_linux.go:139
	}
//line /usr/local/go/src/net/interface_linux.go:139
	// _ = "end of CoverTab[5833]"
//line /usr/local/go/src/net/interface_linux.go:139
	_go_fuzz_dep_.CoverTab[5834]++
							ifat, err := addrTable(ift, ifi, msgs)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:141
		_go_fuzz_dep_.CoverTab[5844]++
								return nil, err
//line /usr/local/go/src/net/interface_linux.go:142
		// _ = "end of CoverTab[5844]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:143
		_go_fuzz_dep_.CoverTab[5845]++
//line /usr/local/go/src/net/interface_linux.go:143
		// _ = "end of CoverTab[5845]"
//line /usr/local/go/src/net/interface_linux.go:143
	}
//line /usr/local/go/src/net/interface_linux.go:143
	// _ = "end of CoverTab[5834]"
//line /usr/local/go/src/net/interface_linux.go:143
	_go_fuzz_dep_.CoverTab[5835]++
							return ifat, nil
//line /usr/local/go/src/net/interface_linux.go:144
	// _ = "end of CoverTab[5835]"
}

func addrTable(ift []Interface, ifi *Interface, msgs []syscall.NetlinkMessage) ([]Addr, error) {
//line /usr/local/go/src/net/interface_linux.go:147
	_go_fuzz_dep_.CoverTab[5846]++
							var ifat []Addr
loop:
	for _, m := range msgs {
//line /usr/local/go/src/net/interface_linux.go:150
		_go_fuzz_dep_.CoverTab[5848]++
								switch m.Header.Type {
		case syscall.NLMSG_DONE:
//line /usr/local/go/src/net/interface_linux.go:152
			_go_fuzz_dep_.CoverTab[5849]++
									break loop
//line /usr/local/go/src/net/interface_linux.go:153
			// _ = "end of CoverTab[5849]"
		case syscall.RTM_NEWADDR:
//line /usr/local/go/src/net/interface_linux.go:154
			_go_fuzz_dep_.CoverTab[5850]++
									ifam := (*syscall.IfAddrmsg)(unsafe.Pointer(&m.Data[0]))
									if len(ift) != 0 || func() bool {
//line /usr/local/go/src/net/interface_linux.go:156
				_go_fuzz_dep_.CoverTab[5852]++
//line /usr/local/go/src/net/interface_linux.go:156
				return ifi.Index == int(ifam.Index)
//line /usr/local/go/src/net/interface_linux.go:156
				// _ = "end of CoverTab[5852]"
//line /usr/local/go/src/net/interface_linux.go:156
			}() {
//line /usr/local/go/src/net/interface_linux.go:156
				_go_fuzz_dep_.CoverTab[5853]++
										if len(ift) != 0 {
//line /usr/local/go/src/net/interface_linux.go:157
					_go_fuzz_dep_.CoverTab[5856]++
											var err error
											ifi, err = interfaceByIndex(ift, int(ifam.Index))
											if err != nil {
//line /usr/local/go/src/net/interface_linux.go:160
						_go_fuzz_dep_.CoverTab[5857]++
												return nil, err
//line /usr/local/go/src/net/interface_linux.go:161
						// _ = "end of CoverTab[5857]"
					} else {
//line /usr/local/go/src/net/interface_linux.go:162
						_go_fuzz_dep_.CoverTab[5858]++
//line /usr/local/go/src/net/interface_linux.go:162
						// _ = "end of CoverTab[5858]"
//line /usr/local/go/src/net/interface_linux.go:162
					}
//line /usr/local/go/src/net/interface_linux.go:162
					// _ = "end of CoverTab[5856]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:163
					_go_fuzz_dep_.CoverTab[5859]++
//line /usr/local/go/src/net/interface_linux.go:163
					// _ = "end of CoverTab[5859]"
//line /usr/local/go/src/net/interface_linux.go:163
				}
//line /usr/local/go/src/net/interface_linux.go:163
				// _ = "end of CoverTab[5853]"
//line /usr/local/go/src/net/interface_linux.go:163
				_go_fuzz_dep_.CoverTab[5854]++
										attrs, err := syscall.ParseNetlinkRouteAttr(&m)
										if err != nil {
//line /usr/local/go/src/net/interface_linux.go:165
					_go_fuzz_dep_.CoverTab[5860]++
											return nil, os.NewSyscallError("parsenetlinkrouteattr", err)
//line /usr/local/go/src/net/interface_linux.go:166
					// _ = "end of CoverTab[5860]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:167
					_go_fuzz_dep_.CoverTab[5861]++
//line /usr/local/go/src/net/interface_linux.go:167
					// _ = "end of CoverTab[5861]"
//line /usr/local/go/src/net/interface_linux.go:167
				}
//line /usr/local/go/src/net/interface_linux.go:167
				// _ = "end of CoverTab[5854]"
//line /usr/local/go/src/net/interface_linux.go:167
				_go_fuzz_dep_.CoverTab[5855]++
										ifa := newAddr(ifam, attrs)
										if ifa != nil {
//line /usr/local/go/src/net/interface_linux.go:169
					_go_fuzz_dep_.CoverTab[5862]++
											ifat = append(ifat, ifa)
//line /usr/local/go/src/net/interface_linux.go:170
					// _ = "end of CoverTab[5862]"
				} else {
//line /usr/local/go/src/net/interface_linux.go:171
					_go_fuzz_dep_.CoverTab[5863]++
//line /usr/local/go/src/net/interface_linux.go:171
					// _ = "end of CoverTab[5863]"
//line /usr/local/go/src/net/interface_linux.go:171
				}
//line /usr/local/go/src/net/interface_linux.go:171
				// _ = "end of CoverTab[5855]"
			} else {
//line /usr/local/go/src/net/interface_linux.go:172
				_go_fuzz_dep_.CoverTab[5864]++
//line /usr/local/go/src/net/interface_linux.go:172
				// _ = "end of CoverTab[5864]"
//line /usr/local/go/src/net/interface_linux.go:172
			}
//line /usr/local/go/src/net/interface_linux.go:172
			// _ = "end of CoverTab[5850]"
//line /usr/local/go/src/net/interface_linux.go:172
		default:
//line /usr/local/go/src/net/interface_linux.go:172
			_go_fuzz_dep_.CoverTab[5851]++
//line /usr/local/go/src/net/interface_linux.go:172
			// _ = "end of CoverTab[5851]"
		}
//line /usr/local/go/src/net/interface_linux.go:173
		// _ = "end of CoverTab[5848]"
	}
//line /usr/local/go/src/net/interface_linux.go:174
	// _ = "end of CoverTab[5846]"
//line /usr/local/go/src/net/interface_linux.go:174
	_go_fuzz_dep_.CoverTab[5847]++
							return ifat, nil
//line /usr/local/go/src/net/interface_linux.go:175
	// _ = "end of CoverTab[5847]"
}

func newAddr(ifam *syscall.IfAddrmsg, attrs []syscall.NetlinkRouteAttr) Addr {
//line /usr/local/go/src/net/interface_linux.go:178
	_go_fuzz_dep_.CoverTab[5865]++
							var ipPointToPoint bool

//line /usr/local/go/src/net/interface_linux.go:183
	for _, a := range attrs {
//line /usr/local/go/src/net/interface_linux.go:183
		_go_fuzz_dep_.CoverTab[5868]++
								if a.Attr.Type == syscall.IFA_LOCAL {
//line /usr/local/go/src/net/interface_linux.go:184
			_go_fuzz_dep_.CoverTab[5869]++
									ipPointToPoint = true
									break
//line /usr/local/go/src/net/interface_linux.go:186
			// _ = "end of CoverTab[5869]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:187
			_go_fuzz_dep_.CoverTab[5870]++
//line /usr/local/go/src/net/interface_linux.go:187
			// _ = "end of CoverTab[5870]"
//line /usr/local/go/src/net/interface_linux.go:187
		}
//line /usr/local/go/src/net/interface_linux.go:187
		// _ = "end of CoverTab[5868]"
	}
//line /usr/local/go/src/net/interface_linux.go:188
	// _ = "end of CoverTab[5865]"
//line /usr/local/go/src/net/interface_linux.go:188
	_go_fuzz_dep_.CoverTab[5866]++
							for _, a := range attrs {
//line /usr/local/go/src/net/interface_linux.go:189
		_go_fuzz_dep_.CoverTab[5871]++
								if ipPointToPoint && func() bool {
//line /usr/local/go/src/net/interface_linux.go:190
			_go_fuzz_dep_.CoverTab[5873]++
//line /usr/local/go/src/net/interface_linux.go:190
			return a.Attr.Type == syscall.IFA_ADDRESS
//line /usr/local/go/src/net/interface_linux.go:190
			// _ = "end of CoverTab[5873]"
//line /usr/local/go/src/net/interface_linux.go:190
		}() {
//line /usr/local/go/src/net/interface_linux.go:190
			_go_fuzz_dep_.CoverTab[5874]++
									continue
//line /usr/local/go/src/net/interface_linux.go:191
			// _ = "end of CoverTab[5874]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:192
			_go_fuzz_dep_.CoverTab[5875]++
//line /usr/local/go/src/net/interface_linux.go:192
			// _ = "end of CoverTab[5875]"
//line /usr/local/go/src/net/interface_linux.go:192
		}
//line /usr/local/go/src/net/interface_linux.go:192
		// _ = "end of CoverTab[5871]"
//line /usr/local/go/src/net/interface_linux.go:192
		_go_fuzz_dep_.CoverTab[5872]++
								switch ifam.Family {
		case syscall.AF_INET:
//line /usr/local/go/src/net/interface_linux.go:194
			_go_fuzz_dep_.CoverTab[5876]++
									return &IPNet{IP: IPv4(a.Value[0], a.Value[1], a.Value[2], a.Value[3]), Mask: CIDRMask(int(ifam.Prefixlen), 8*IPv4len)}
//line /usr/local/go/src/net/interface_linux.go:195
			// _ = "end of CoverTab[5876]"
		case syscall.AF_INET6:
//line /usr/local/go/src/net/interface_linux.go:196
			_go_fuzz_dep_.CoverTab[5877]++
									ifa := &IPNet{IP: make(IP, IPv6len), Mask: CIDRMask(int(ifam.Prefixlen), 8*IPv6len)}
									copy(ifa.IP, a.Value[:])
									return ifa
//line /usr/local/go/src/net/interface_linux.go:199
			// _ = "end of CoverTab[5877]"
//line /usr/local/go/src/net/interface_linux.go:199
		default:
//line /usr/local/go/src/net/interface_linux.go:199
			_go_fuzz_dep_.CoverTab[5878]++
//line /usr/local/go/src/net/interface_linux.go:199
			// _ = "end of CoverTab[5878]"
		}
//line /usr/local/go/src/net/interface_linux.go:200
		// _ = "end of CoverTab[5872]"
	}
//line /usr/local/go/src/net/interface_linux.go:201
	// _ = "end of CoverTab[5866]"
//line /usr/local/go/src/net/interface_linux.go:201
	_go_fuzz_dep_.CoverTab[5867]++
							return nil
//line /usr/local/go/src/net/interface_linux.go:202
	// _ = "end of CoverTab[5867]"
}

// interfaceMulticastAddrTable returns addresses for a specific
//line /usr/local/go/src/net/interface_linux.go:205
// interface.
//line /usr/local/go/src/net/interface_linux.go:207
func interfaceMulticastAddrTable(ifi *Interface) ([]Addr, error) {
//line /usr/local/go/src/net/interface_linux.go:207
	_go_fuzz_dep_.CoverTab[5879]++
							ifmat4 := parseProcNetIGMP("/proc/net/igmp", ifi)
							ifmat6 := parseProcNetIGMP6("/proc/net/igmp6", ifi)
							return append(ifmat4, ifmat6...), nil
//line /usr/local/go/src/net/interface_linux.go:210
	// _ = "end of CoverTab[5879]"
}

func parseProcNetIGMP(path string, ifi *Interface) []Addr {
//line /usr/local/go/src/net/interface_linux.go:213
	_go_fuzz_dep_.CoverTab[5880]++
							fd, err := open(path)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:215
		_go_fuzz_dep_.CoverTab[5883]++
								return nil
//line /usr/local/go/src/net/interface_linux.go:216
		// _ = "end of CoverTab[5883]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:217
		_go_fuzz_dep_.CoverTab[5884]++
//line /usr/local/go/src/net/interface_linux.go:217
		// _ = "end of CoverTab[5884]"
//line /usr/local/go/src/net/interface_linux.go:217
	}
//line /usr/local/go/src/net/interface_linux.go:217
	// _ = "end of CoverTab[5880]"
//line /usr/local/go/src/net/interface_linux.go:217
	_go_fuzz_dep_.CoverTab[5881]++
							defer fd.close()
							var (
		ifmat	[]Addr
		name	string
	)
	fd.readLine()
	b := make([]byte, IPv4len)
	for l, ok := fd.readLine(); ok; l, ok = fd.readLine() {
//line /usr/local/go/src/net/interface_linux.go:225
		_go_fuzz_dep_.CoverTab[5885]++
								f := splitAtBytes(l, " :\r\t\n")
								if len(f) < 4 {
//line /usr/local/go/src/net/interface_linux.go:227
			_go_fuzz_dep_.CoverTab[5887]++
									continue
//line /usr/local/go/src/net/interface_linux.go:228
			// _ = "end of CoverTab[5887]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:229
			_go_fuzz_dep_.CoverTab[5888]++
//line /usr/local/go/src/net/interface_linux.go:229
			// _ = "end of CoverTab[5888]"
//line /usr/local/go/src/net/interface_linux.go:229
		}
//line /usr/local/go/src/net/interface_linux.go:229
		// _ = "end of CoverTab[5885]"
//line /usr/local/go/src/net/interface_linux.go:229
		_go_fuzz_dep_.CoverTab[5886]++
								switch {
		case l[0] != ' ' && func() bool {
//line /usr/local/go/src/net/interface_linux.go:231
			_go_fuzz_dep_.CoverTab[5892]++
//line /usr/local/go/src/net/interface_linux.go:231
			return l[0] != '\t'
//line /usr/local/go/src/net/interface_linux.go:231
			// _ = "end of CoverTab[5892]"
//line /usr/local/go/src/net/interface_linux.go:231
		}():
//line /usr/local/go/src/net/interface_linux.go:231
			_go_fuzz_dep_.CoverTab[5889]++
									name = f[1]
//line /usr/local/go/src/net/interface_linux.go:232
			// _ = "end of CoverTab[5889]"
		case len(f[0]) == 8:
//line /usr/local/go/src/net/interface_linux.go:233
			_go_fuzz_dep_.CoverTab[5890]++
									if ifi == nil || func() bool {
//line /usr/local/go/src/net/interface_linux.go:234
				_go_fuzz_dep_.CoverTab[5893]++
//line /usr/local/go/src/net/interface_linux.go:234
				return name == ifi.Name
//line /usr/local/go/src/net/interface_linux.go:234
				// _ = "end of CoverTab[5893]"
//line /usr/local/go/src/net/interface_linux.go:234
			}() {
//line /usr/local/go/src/net/interface_linux.go:234
				_go_fuzz_dep_.CoverTab[5894]++

//line /usr/local/go/src/net/interface_linux.go:238
				for i := 0; i+1 < len(f[0]); i += 2 {
//line /usr/local/go/src/net/interface_linux.go:238
					_go_fuzz_dep_.CoverTab[5896]++
											b[i/2], _ = xtoi2(f[0][i:i+2], 0)
//line /usr/local/go/src/net/interface_linux.go:239
					// _ = "end of CoverTab[5896]"
				}
//line /usr/local/go/src/net/interface_linux.go:240
				// _ = "end of CoverTab[5894]"
//line /usr/local/go/src/net/interface_linux.go:240
				_go_fuzz_dep_.CoverTab[5895]++
										i := *(*uint32)(unsafe.Pointer(&b[:4][0]))
										ifma := &IPAddr{IP: IPv4(byte(i>>24), byte(i>>16), byte(i>>8), byte(i))}
										ifmat = append(ifmat, ifma)
//line /usr/local/go/src/net/interface_linux.go:243
				// _ = "end of CoverTab[5895]"
			} else {
//line /usr/local/go/src/net/interface_linux.go:244
				_go_fuzz_dep_.CoverTab[5897]++
//line /usr/local/go/src/net/interface_linux.go:244
				// _ = "end of CoverTab[5897]"
//line /usr/local/go/src/net/interface_linux.go:244
			}
//line /usr/local/go/src/net/interface_linux.go:244
			// _ = "end of CoverTab[5890]"
//line /usr/local/go/src/net/interface_linux.go:244
		default:
//line /usr/local/go/src/net/interface_linux.go:244
			_go_fuzz_dep_.CoverTab[5891]++
//line /usr/local/go/src/net/interface_linux.go:244
			// _ = "end of CoverTab[5891]"
		}
//line /usr/local/go/src/net/interface_linux.go:245
		// _ = "end of CoverTab[5886]"
	}
//line /usr/local/go/src/net/interface_linux.go:246
	// _ = "end of CoverTab[5881]"
//line /usr/local/go/src/net/interface_linux.go:246
	_go_fuzz_dep_.CoverTab[5882]++
							return ifmat
//line /usr/local/go/src/net/interface_linux.go:247
	// _ = "end of CoverTab[5882]"
}

func parseProcNetIGMP6(path string, ifi *Interface) []Addr {
//line /usr/local/go/src/net/interface_linux.go:250
	_go_fuzz_dep_.CoverTab[5898]++
							fd, err := open(path)
							if err != nil {
//line /usr/local/go/src/net/interface_linux.go:252
		_go_fuzz_dep_.CoverTab[5901]++
								return nil
//line /usr/local/go/src/net/interface_linux.go:253
		// _ = "end of CoverTab[5901]"
	} else {
//line /usr/local/go/src/net/interface_linux.go:254
		_go_fuzz_dep_.CoverTab[5902]++
//line /usr/local/go/src/net/interface_linux.go:254
		// _ = "end of CoverTab[5902]"
//line /usr/local/go/src/net/interface_linux.go:254
	}
//line /usr/local/go/src/net/interface_linux.go:254
	// _ = "end of CoverTab[5898]"
//line /usr/local/go/src/net/interface_linux.go:254
	_go_fuzz_dep_.CoverTab[5899]++
							defer fd.close()
							var ifmat []Addr
							b := make([]byte, IPv6len)
							for l, ok := fd.readLine(); ok; l, ok = fd.readLine() {
//line /usr/local/go/src/net/interface_linux.go:258
		_go_fuzz_dep_.CoverTab[5903]++
								f := splitAtBytes(l, " \r\t\n")
								if len(f) < 6 {
//line /usr/local/go/src/net/interface_linux.go:260
			_go_fuzz_dep_.CoverTab[5905]++
									continue
//line /usr/local/go/src/net/interface_linux.go:261
			// _ = "end of CoverTab[5905]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:262
			_go_fuzz_dep_.CoverTab[5906]++
//line /usr/local/go/src/net/interface_linux.go:262
			// _ = "end of CoverTab[5906]"
//line /usr/local/go/src/net/interface_linux.go:262
		}
//line /usr/local/go/src/net/interface_linux.go:262
		// _ = "end of CoverTab[5903]"
//line /usr/local/go/src/net/interface_linux.go:262
		_go_fuzz_dep_.CoverTab[5904]++
								if ifi == nil || func() bool {
//line /usr/local/go/src/net/interface_linux.go:263
			_go_fuzz_dep_.CoverTab[5907]++
//line /usr/local/go/src/net/interface_linux.go:263
			return f[1] == ifi.Name
//line /usr/local/go/src/net/interface_linux.go:263
			// _ = "end of CoverTab[5907]"
//line /usr/local/go/src/net/interface_linux.go:263
		}() {
//line /usr/local/go/src/net/interface_linux.go:263
			_go_fuzz_dep_.CoverTab[5908]++
									for i := 0; i+1 < len(f[2]); i += 2 {
//line /usr/local/go/src/net/interface_linux.go:264
				_go_fuzz_dep_.CoverTab[5910]++
										b[i/2], _ = xtoi2(f[2][i:i+2], 0)
//line /usr/local/go/src/net/interface_linux.go:265
				// _ = "end of CoverTab[5910]"
			}
//line /usr/local/go/src/net/interface_linux.go:266
			// _ = "end of CoverTab[5908]"
//line /usr/local/go/src/net/interface_linux.go:266
			_go_fuzz_dep_.CoverTab[5909]++
									ifma := &IPAddr{IP: IP{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15]}}
									ifmat = append(ifmat, ifma)
//line /usr/local/go/src/net/interface_linux.go:268
			// _ = "end of CoverTab[5909]"
		} else {
//line /usr/local/go/src/net/interface_linux.go:269
			_go_fuzz_dep_.CoverTab[5911]++
//line /usr/local/go/src/net/interface_linux.go:269
			// _ = "end of CoverTab[5911]"
//line /usr/local/go/src/net/interface_linux.go:269
		}
//line /usr/local/go/src/net/interface_linux.go:269
		// _ = "end of CoverTab[5904]"
	}
//line /usr/local/go/src/net/interface_linux.go:270
	// _ = "end of CoverTab[5899]"
//line /usr/local/go/src/net/interface_linux.go:270
	_go_fuzz_dep_.CoverTab[5900]++
							return ifmat
//line /usr/local/go/src/net/interface_linux.go:271
	// _ = "end of CoverTab[5900]"
}

//line /usr/local/go/src/net/interface_linux.go:272
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/interface_linux.go:272
var _ = _go_fuzz_dep_.CoverTab
