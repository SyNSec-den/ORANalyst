// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/interface.go:5
package net

//line /snap/go/10455/src/net/interface.go:5
import (
//line /snap/go/10455/src/net/interface.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/interface.go:5
)
//line /snap/go/10455/src/net/interface.go:5
import (
//line /snap/go/10455/src/net/interface.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/interface.go:5
)

import (
	"errors"
	"internal/itoa"
	"sync"
	"time"
)

//line /snap/go/10455/src/net/interface.go:20
var (
	errInvalidInterface		= errors.New("invalid network interface")
	errInvalidInterfaceIndex	= errors.New("invalid network interface index")
	errInvalidInterfaceName		= errors.New("invalid network interface name")
	errNoSuchInterface		= errors.New("no such network interface")
	errNoSuchMulticastInterface	= errors.New("no such multicast network interface")
)

// Interface represents a mapping between network interface name
//line /snap/go/10455/src/net/interface.go:28
// and index. It also represents network interface facility
//line /snap/go/10455/src/net/interface.go:28
// information.
//line /snap/go/10455/src/net/interface.go:31
type Interface struct {
	Index		int		// positive integer that starts at one, zero is never used
	MTU		int		// maximum transmission unit
	Name		string		// e.g., "en0", "lo0", "eth0.100"
	HardwareAddr	HardwareAddr	// IEEE MAC-48, EUI-48 and EUI-64 form
	Flags		Flags		// e.g., FlagUp, FlagLoopback, FlagMulticast
}

type Flags uint

const (
	FlagUp			Flags	= 1 << iota	// interface is administratively up
	FlagBroadcast					// interface supports broadcast access capability
	FlagLoopback					// interface is a loopback interface
	FlagPointToPoint				// interface belongs to a point-to-point link
	FlagMulticast					// interface supports multicast access capability
	FlagRunning					// interface is in running state
)

var flagNames = []string{
	"up",
	"broadcast",
	"loopback",
	"pointtopoint",
	"multicast",
	"running",
}

func (f Flags) String() string {
//line /snap/go/10455/src/net/interface.go:59
	_go_fuzz_dep_.CoverTab[6045]++
						s := ""
//line /snap/go/10455/src/net/interface.go:60
	_go_fuzz_dep_.CoverTab[786677] = 0
						for i, name := range flagNames {
//line /snap/go/10455/src/net/interface.go:61
		if _go_fuzz_dep_.CoverTab[786677] == 0 {
//line /snap/go/10455/src/net/interface.go:61
			_go_fuzz_dep_.CoverTab[528428]++
//line /snap/go/10455/src/net/interface.go:61
		} else {
//line /snap/go/10455/src/net/interface.go:61
			_go_fuzz_dep_.CoverTab[528429]++
//line /snap/go/10455/src/net/interface.go:61
		}
//line /snap/go/10455/src/net/interface.go:61
		_go_fuzz_dep_.CoverTab[786677] = 1
//line /snap/go/10455/src/net/interface.go:61
		_go_fuzz_dep_.CoverTab[6048]++
							if f&(1<<uint(i)) != 0 {
//line /snap/go/10455/src/net/interface.go:62
			_go_fuzz_dep_.CoverTab[528372]++
//line /snap/go/10455/src/net/interface.go:62
			_go_fuzz_dep_.CoverTab[6049]++
								if s != "" {
//line /snap/go/10455/src/net/interface.go:63
				_go_fuzz_dep_.CoverTab[528374]++
//line /snap/go/10455/src/net/interface.go:63
				_go_fuzz_dep_.CoverTab[6051]++
									s += "|"
//line /snap/go/10455/src/net/interface.go:64
				// _ = "end of CoverTab[6051]"
			} else {
//line /snap/go/10455/src/net/interface.go:65
				_go_fuzz_dep_.CoverTab[528375]++
//line /snap/go/10455/src/net/interface.go:65
				_go_fuzz_dep_.CoverTab[6052]++
//line /snap/go/10455/src/net/interface.go:65
				// _ = "end of CoverTab[6052]"
//line /snap/go/10455/src/net/interface.go:65
			}
//line /snap/go/10455/src/net/interface.go:65
			// _ = "end of CoverTab[6049]"
//line /snap/go/10455/src/net/interface.go:65
			_go_fuzz_dep_.CoverTab[6050]++
								s += name
//line /snap/go/10455/src/net/interface.go:66
			// _ = "end of CoverTab[6050]"
		} else {
//line /snap/go/10455/src/net/interface.go:67
			_go_fuzz_dep_.CoverTab[528373]++
//line /snap/go/10455/src/net/interface.go:67
			_go_fuzz_dep_.CoverTab[6053]++
//line /snap/go/10455/src/net/interface.go:67
			// _ = "end of CoverTab[6053]"
//line /snap/go/10455/src/net/interface.go:67
		}
//line /snap/go/10455/src/net/interface.go:67
		// _ = "end of CoverTab[6048]"
	}
//line /snap/go/10455/src/net/interface.go:68
	if _go_fuzz_dep_.CoverTab[786677] == 0 {
//line /snap/go/10455/src/net/interface.go:68
		_go_fuzz_dep_.CoverTab[528430]++
//line /snap/go/10455/src/net/interface.go:68
	} else {
//line /snap/go/10455/src/net/interface.go:68
		_go_fuzz_dep_.CoverTab[528431]++
//line /snap/go/10455/src/net/interface.go:68
	}
//line /snap/go/10455/src/net/interface.go:68
	// _ = "end of CoverTab[6045]"
//line /snap/go/10455/src/net/interface.go:68
	_go_fuzz_dep_.CoverTab[6046]++
						if s == "" {
//line /snap/go/10455/src/net/interface.go:69
		_go_fuzz_dep_.CoverTab[528376]++
//line /snap/go/10455/src/net/interface.go:69
		_go_fuzz_dep_.CoverTab[6054]++
							s = "0"
//line /snap/go/10455/src/net/interface.go:70
		// _ = "end of CoverTab[6054]"
	} else {
//line /snap/go/10455/src/net/interface.go:71
		_go_fuzz_dep_.CoverTab[528377]++
//line /snap/go/10455/src/net/interface.go:71
		_go_fuzz_dep_.CoverTab[6055]++
//line /snap/go/10455/src/net/interface.go:71
		// _ = "end of CoverTab[6055]"
//line /snap/go/10455/src/net/interface.go:71
	}
//line /snap/go/10455/src/net/interface.go:71
	// _ = "end of CoverTab[6046]"
//line /snap/go/10455/src/net/interface.go:71
	_go_fuzz_dep_.CoverTab[6047]++
						return s
//line /snap/go/10455/src/net/interface.go:72
	// _ = "end of CoverTab[6047]"
}

// Addrs returns a list of unicast interface addresses for a specific
//line /snap/go/10455/src/net/interface.go:75
// interface.
//line /snap/go/10455/src/net/interface.go:77
func (ifi *Interface) Addrs() ([]Addr, error) {
//line /snap/go/10455/src/net/interface.go:77
	_go_fuzz_dep_.CoverTab[6056]++
						if ifi == nil {
//line /snap/go/10455/src/net/interface.go:78
		_go_fuzz_dep_.CoverTab[528378]++
//line /snap/go/10455/src/net/interface.go:78
		_go_fuzz_dep_.CoverTab[6059]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterface}
//line /snap/go/10455/src/net/interface.go:79
		// _ = "end of CoverTab[6059]"
	} else {
//line /snap/go/10455/src/net/interface.go:80
		_go_fuzz_dep_.CoverTab[528379]++
//line /snap/go/10455/src/net/interface.go:80
		_go_fuzz_dep_.CoverTab[6060]++
//line /snap/go/10455/src/net/interface.go:80
		// _ = "end of CoverTab[6060]"
//line /snap/go/10455/src/net/interface.go:80
	}
//line /snap/go/10455/src/net/interface.go:80
	// _ = "end of CoverTab[6056]"
//line /snap/go/10455/src/net/interface.go:80
	_go_fuzz_dep_.CoverTab[6057]++
						ifat, err := interfaceAddrTable(ifi)
						if err != nil {
//line /snap/go/10455/src/net/interface.go:82
		_go_fuzz_dep_.CoverTab[528380]++
//line /snap/go/10455/src/net/interface.go:82
		_go_fuzz_dep_.CoverTab[6061]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/interface.go:83
		// _ = "end of CoverTab[6061]"
	} else {
//line /snap/go/10455/src/net/interface.go:84
		_go_fuzz_dep_.CoverTab[528381]++
//line /snap/go/10455/src/net/interface.go:84
		_go_fuzz_dep_.CoverTab[6062]++
//line /snap/go/10455/src/net/interface.go:84
		// _ = "end of CoverTab[6062]"
//line /snap/go/10455/src/net/interface.go:84
	}
//line /snap/go/10455/src/net/interface.go:84
	// _ = "end of CoverTab[6057]"
//line /snap/go/10455/src/net/interface.go:84
	_go_fuzz_dep_.CoverTab[6058]++
						return ifat, err
//line /snap/go/10455/src/net/interface.go:85
	// _ = "end of CoverTab[6058]"
}

// MulticastAddrs returns a list of multicast, joined group addresses
//line /snap/go/10455/src/net/interface.go:88
// for a specific interface.
//line /snap/go/10455/src/net/interface.go:90
func (ifi *Interface) MulticastAddrs() ([]Addr, error) {
//line /snap/go/10455/src/net/interface.go:90
	_go_fuzz_dep_.CoverTab[6063]++
						if ifi == nil {
//line /snap/go/10455/src/net/interface.go:91
		_go_fuzz_dep_.CoverTab[528382]++
//line /snap/go/10455/src/net/interface.go:91
		_go_fuzz_dep_.CoverTab[6066]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterface}
//line /snap/go/10455/src/net/interface.go:92
		// _ = "end of CoverTab[6066]"
	} else {
//line /snap/go/10455/src/net/interface.go:93
		_go_fuzz_dep_.CoverTab[528383]++
//line /snap/go/10455/src/net/interface.go:93
		_go_fuzz_dep_.CoverTab[6067]++
//line /snap/go/10455/src/net/interface.go:93
		// _ = "end of CoverTab[6067]"
//line /snap/go/10455/src/net/interface.go:93
	}
//line /snap/go/10455/src/net/interface.go:93
	// _ = "end of CoverTab[6063]"
//line /snap/go/10455/src/net/interface.go:93
	_go_fuzz_dep_.CoverTab[6064]++
						ifat, err := interfaceMulticastAddrTable(ifi)
						if err != nil {
//line /snap/go/10455/src/net/interface.go:95
		_go_fuzz_dep_.CoverTab[528384]++
//line /snap/go/10455/src/net/interface.go:95
		_go_fuzz_dep_.CoverTab[6068]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/interface.go:96
		// _ = "end of CoverTab[6068]"
	} else {
//line /snap/go/10455/src/net/interface.go:97
		_go_fuzz_dep_.CoverTab[528385]++
//line /snap/go/10455/src/net/interface.go:97
		_go_fuzz_dep_.CoverTab[6069]++
//line /snap/go/10455/src/net/interface.go:97
		// _ = "end of CoverTab[6069]"
//line /snap/go/10455/src/net/interface.go:97
	}
//line /snap/go/10455/src/net/interface.go:97
	// _ = "end of CoverTab[6064]"
//line /snap/go/10455/src/net/interface.go:97
	_go_fuzz_dep_.CoverTab[6065]++
						return ifat, err
//line /snap/go/10455/src/net/interface.go:98
	// _ = "end of CoverTab[6065]"
}

// Interfaces returns a list of the system's network interfaces.
func Interfaces() ([]Interface, error) {
//line /snap/go/10455/src/net/interface.go:102
	_go_fuzz_dep_.CoverTab[6070]++
						ift, err := interfaceTable(0)
						if err != nil {
//line /snap/go/10455/src/net/interface.go:104
		_go_fuzz_dep_.CoverTab[528386]++
//line /snap/go/10455/src/net/interface.go:104
		_go_fuzz_dep_.CoverTab[6073]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/interface.go:105
		// _ = "end of CoverTab[6073]"
	} else {
//line /snap/go/10455/src/net/interface.go:106
		_go_fuzz_dep_.CoverTab[528387]++
//line /snap/go/10455/src/net/interface.go:106
		_go_fuzz_dep_.CoverTab[6074]++
//line /snap/go/10455/src/net/interface.go:106
		// _ = "end of CoverTab[6074]"
//line /snap/go/10455/src/net/interface.go:106
	}
//line /snap/go/10455/src/net/interface.go:106
	// _ = "end of CoverTab[6070]"
//line /snap/go/10455/src/net/interface.go:106
	_go_fuzz_dep_.CoverTab[6071]++
						if len(ift) != 0 {
//line /snap/go/10455/src/net/interface.go:107
		_go_fuzz_dep_.CoverTab[528388]++
//line /snap/go/10455/src/net/interface.go:107
		_go_fuzz_dep_.CoverTab[6075]++
							zoneCache.update(ift, false)
//line /snap/go/10455/src/net/interface.go:108
		// _ = "end of CoverTab[6075]"
	} else {
//line /snap/go/10455/src/net/interface.go:109
		_go_fuzz_dep_.CoverTab[528389]++
//line /snap/go/10455/src/net/interface.go:109
		_go_fuzz_dep_.CoverTab[6076]++
//line /snap/go/10455/src/net/interface.go:109
		// _ = "end of CoverTab[6076]"
//line /snap/go/10455/src/net/interface.go:109
	}
//line /snap/go/10455/src/net/interface.go:109
	// _ = "end of CoverTab[6071]"
//line /snap/go/10455/src/net/interface.go:109
	_go_fuzz_dep_.CoverTab[6072]++
						return ift, nil
//line /snap/go/10455/src/net/interface.go:110
	// _ = "end of CoverTab[6072]"
}

// InterfaceAddrs returns a list of the system's unicast interface
//line /snap/go/10455/src/net/interface.go:113
// addresses.
//line /snap/go/10455/src/net/interface.go:113
//
//line /snap/go/10455/src/net/interface.go:113
// The returned list does not identify the associated interface; use
//line /snap/go/10455/src/net/interface.go:113
// Interfaces and Interface.Addrs for more detail.
//line /snap/go/10455/src/net/interface.go:118
func InterfaceAddrs() ([]Addr, error) {
//line /snap/go/10455/src/net/interface.go:118
	_go_fuzz_dep_.CoverTab[6077]++
						ifat, err := interfaceAddrTable(nil)
						if err != nil {
//line /snap/go/10455/src/net/interface.go:120
		_go_fuzz_dep_.CoverTab[528390]++
//line /snap/go/10455/src/net/interface.go:120
		_go_fuzz_dep_.CoverTab[6079]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/interface.go:121
		// _ = "end of CoverTab[6079]"
	} else {
//line /snap/go/10455/src/net/interface.go:122
		_go_fuzz_dep_.CoverTab[528391]++
//line /snap/go/10455/src/net/interface.go:122
		_go_fuzz_dep_.CoverTab[6080]++
//line /snap/go/10455/src/net/interface.go:122
		// _ = "end of CoverTab[6080]"
//line /snap/go/10455/src/net/interface.go:122
	}
//line /snap/go/10455/src/net/interface.go:122
	// _ = "end of CoverTab[6077]"
//line /snap/go/10455/src/net/interface.go:122
	_go_fuzz_dep_.CoverTab[6078]++
						return ifat, err
//line /snap/go/10455/src/net/interface.go:123
	// _ = "end of CoverTab[6078]"
}

// InterfaceByIndex returns the interface specified by index.
//line /snap/go/10455/src/net/interface.go:126
//
//line /snap/go/10455/src/net/interface.go:126
// On Solaris, it returns one of the logical network interfaces
//line /snap/go/10455/src/net/interface.go:126
// sharing the logical data link; for more precision use
//line /snap/go/10455/src/net/interface.go:126
// InterfaceByName.
//line /snap/go/10455/src/net/interface.go:131
func InterfaceByIndex(index int) (*Interface, error) {
//line /snap/go/10455/src/net/interface.go:131
	_go_fuzz_dep_.CoverTab[6081]++
						if index <= 0 {
//line /snap/go/10455/src/net/interface.go:132
		_go_fuzz_dep_.CoverTab[528392]++
//line /snap/go/10455/src/net/interface.go:132
		_go_fuzz_dep_.CoverTab[6085]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterfaceIndex}
//line /snap/go/10455/src/net/interface.go:133
		// _ = "end of CoverTab[6085]"
	} else {
//line /snap/go/10455/src/net/interface.go:134
		_go_fuzz_dep_.CoverTab[528393]++
//line /snap/go/10455/src/net/interface.go:134
		_go_fuzz_dep_.CoverTab[6086]++
//line /snap/go/10455/src/net/interface.go:134
		// _ = "end of CoverTab[6086]"
//line /snap/go/10455/src/net/interface.go:134
	}
//line /snap/go/10455/src/net/interface.go:134
	// _ = "end of CoverTab[6081]"
//line /snap/go/10455/src/net/interface.go:134
	_go_fuzz_dep_.CoverTab[6082]++
						ift, err := interfaceTable(index)
						if err != nil {
//line /snap/go/10455/src/net/interface.go:136
		_go_fuzz_dep_.CoverTab[528394]++
//line /snap/go/10455/src/net/interface.go:136
		_go_fuzz_dep_.CoverTab[6087]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/interface.go:137
		// _ = "end of CoverTab[6087]"
	} else {
//line /snap/go/10455/src/net/interface.go:138
		_go_fuzz_dep_.CoverTab[528395]++
//line /snap/go/10455/src/net/interface.go:138
		_go_fuzz_dep_.CoverTab[6088]++
//line /snap/go/10455/src/net/interface.go:138
		// _ = "end of CoverTab[6088]"
//line /snap/go/10455/src/net/interface.go:138
	}
//line /snap/go/10455/src/net/interface.go:138
	// _ = "end of CoverTab[6082]"
//line /snap/go/10455/src/net/interface.go:138
	_go_fuzz_dep_.CoverTab[6083]++
						ifi, err := interfaceByIndex(ift, index)
						if err != nil {
//line /snap/go/10455/src/net/interface.go:140
		_go_fuzz_dep_.CoverTab[528396]++
//line /snap/go/10455/src/net/interface.go:140
		_go_fuzz_dep_.CoverTab[6089]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/interface.go:141
		// _ = "end of CoverTab[6089]"
	} else {
//line /snap/go/10455/src/net/interface.go:142
		_go_fuzz_dep_.CoverTab[528397]++
//line /snap/go/10455/src/net/interface.go:142
		_go_fuzz_dep_.CoverTab[6090]++
//line /snap/go/10455/src/net/interface.go:142
		// _ = "end of CoverTab[6090]"
//line /snap/go/10455/src/net/interface.go:142
	}
//line /snap/go/10455/src/net/interface.go:142
	// _ = "end of CoverTab[6083]"
//line /snap/go/10455/src/net/interface.go:142
	_go_fuzz_dep_.CoverTab[6084]++
						return ifi, err
//line /snap/go/10455/src/net/interface.go:143
	// _ = "end of CoverTab[6084]"
}

func interfaceByIndex(ift []Interface, index int) (*Interface, error) {
//line /snap/go/10455/src/net/interface.go:146
	_go_fuzz_dep_.CoverTab[6091]++
//line /snap/go/10455/src/net/interface.go:146
	_go_fuzz_dep_.CoverTab[786678] = 0
						for _, ifi := range ift {
//line /snap/go/10455/src/net/interface.go:147
		if _go_fuzz_dep_.CoverTab[786678] == 0 {
//line /snap/go/10455/src/net/interface.go:147
			_go_fuzz_dep_.CoverTab[528432]++
//line /snap/go/10455/src/net/interface.go:147
		} else {
//line /snap/go/10455/src/net/interface.go:147
			_go_fuzz_dep_.CoverTab[528433]++
//line /snap/go/10455/src/net/interface.go:147
		}
//line /snap/go/10455/src/net/interface.go:147
		_go_fuzz_dep_.CoverTab[786678] = 1
//line /snap/go/10455/src/net/interface.go:147
		_go_fuzz_dep_.CoverTab[6093]++
							if index == ifi.Index {
//line /snap/go/10455/src/net/interface.go:148
			_go_fuzz_dep_.CoverTab[528398]++
//line /snap/go/10455/src/net/interface.go:148
			_go_fuzz_dep_.CoverTab[6094]++
								return &ifi, nil
//line /snap/go/10455/src/net/interface.go:149
			// _ = "end of CoverTab[6094]"
		} else {
//line /snap/go/10455/src/net/interface.go:150
			_go_fuzz_dep_.CoverTab[528399]++
//line /snap/go/10455/src/net/interface.go:150
			_go_fuzz_dep_.CoverTab[6095]++
//line /snap/go/10455/src/net/interface.go:150
			// _ = "end of CoverTab[6095]"
//line /snap/go/10455/src/net/interface.go:150
		}
//line /snap/go/10455/src/net/interface.go:150
		// _ = "end of CoverTab[6093]"
	}
//line /snap/go/10455/src/net/interface.go:151
	if _go_fuzz_dep_.CoverTab[786678] == 0 {
//line /snap/go/10455/src/net/interface.go:151
		_go_fuzz_dep_.CoverTab[528434]++
//line /snap/go/10455/src/net/interface.go:151
	} else {
//line /snap/go/10455/src/net/interface.go:151
		_go_fuzz_dep_.CoverTab[528435]++
//line /snap/go/10455/src/net/interface.go:151
	}
//line /snap/go/10455/src/net/interface.go:151
	// _ = "end of CoverTab[6091]"
//line /snap/go/10455/src/net/interface.go:151
	_go_fuzz_dep_.CoverTab[6092]++
						return nil, errNoSuchInterface
//line /snap/go/10455/src/net/interface.go:152
	// _ = "end of CoverTab[6092]"
}

// InterfaceByName returns the interface specified by name.
func InterfaceByName(name string) (*Interface, error) {
//line /snap/go/10455/src/net/interface.go:156
	_go_fuzz_dep_.CoverTab[6096]++
						if name == "" {
//line /snap/go/10455/src/net/interface.go:157
		_go_fuzz_dep_.CoverTab[528400]++
//line /snap/go/10455/src/net/interface.go:157
		_go_fuzz_dep_.CoverTab[6101]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterfaceName}
//line /snap/go/10455/src/net/interface.go:158
		// _ = "end of CoverTab[6101]"
	} else {
//line /snap/go/10455/src/net/interface.go:159
		_go_fuzz_dep_.CoverTab[528401]++
//line /snap/go/10455/src/net/interface.go:159
		_go_fuzz_dep_.CoverTab[6102]++
//line /snap/go/10455/src/net/interface.go:159
		// _ = "end of CoverTab[6102]"
//line /snap/go/10455/src/net/interface.go:159
	}
//line /snap/go/10455/src/net/interface.go:159
	// _ = "end of CoverTab[6096]"
//line /snap/go/10455/src/net/interface.go:159
	_go_fuzz_dep_.CoverTab[6097]++
						ift, err := interfaceTable(0)
						if err != nil {
//line /snap/go/10455/src/net/interface.go:161
		_go_fuzz_dep_.CoverTab[528402]++
//line /snap/go/10455/src/net/interface.go:161
		_go_fuzz_dep_.CoverTab[6103]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /snap/go/10455/src/net/interface.go:162
		// _ = "end of CoverTab[6103]"
	} else {
//line /snap/go/10455/src/net/interface.go:163
		_go_fuzz_dep_.CoverTab[528403]++
//line /snap/go/10455/src/net/interface.go:163
		_go_fuzz_dep_.CoverTab[6104]++
//line /snap/go/10455/src/net/interface.go:163
		// _ = "end of CoverTab[6104]"
//line /snap/go/10455/src/net/interface.go:163
	}
//line /snap/go/10455/src/net/interface.go:163
	// _ = "end of CoverTab[6097]"
//line /snap/go/10455/src/net/interface.go:163
	_go_fuzz_dep_.CoverTab[6098]++
						if len(ift) != 0 {
//line /snap/go/10455/src/net/interface.go:164
		_go_fuzz_dep_.CoverTab[528404]++
//line /snap/go/10455/src/net/interface.go:164
		_go_fuzz_dep_.CoverTab[6105]++
							zoneCache.update(ift, false)
//line /snap/go/10455/src/net/interface.go:165
		// _ = "end of CoverTab[6105]"
	} else {
//line /snap/go/10455/src/net/interface.go:166
		_go_fuzz_dep_.CoverTab[528405]++
//line /snap/go/10455/src/net/interface.go:166
		_go_fuzz_dep_.CoverTab[6106]++
//line /snap/go/10455/src/net/interface.go:166
		// _ = "end of CoverTab[6106]"
//line /snap/go/10455/src/net/interface.go:166
	}
//line /snap/go/10455/src/net/interface.go:166
	// _ = "end of CoverTab[6098]"
//line /snap/go/10455/src/net/interface.go:166
	_go_fuzz_dep_.CoverTab[6099]++
//line /snap/go/10455/src/net/interface.go:166
	_go_fuzz_dep_.CoverTab[786679] = 0
						for _, ifi := range ift {
//line /snap/go/10455/src/net/interface.go:167
		if _go_fuzz_dep_.CoverTab[786679] == 0 {
//line /snap/go/10455/src/net/interface.go:167
			_go_fuzz_dep_.CoverTab[528436]++
//line /snap/go/10455/src/net/interface.go:167
		} else {
//line /snap/go/10455/src/net/interface.go:167
			_go_fuzz_dep_.CoverTab[528437]++
//line /snap/go/10455/src/net/interface.go:167
		}
//line /snap/go/10455/src/net/interface.go:167
		_go_fuzz_dep_.CoverTab[786679] = 1
//line /snap/go/10455/src/net/interface.go:167
		_go_fuzz_dep_.CoverTab[6107]++
							if name == ifi.Name {
//line /snap/go/10455/src/net/interface.go:168
			_go_fuzz_dep_.CoverTab[528406]++
//line /snap/go/10455/src/net/interface.go:168
			_go_fuzz_dep_.CoverTab[6108]++
								return &ifi, nil
//line /snap/go/10455/src/net/interface.go:169
			// _ = "end of CoverTab[6108]"
		} else {
//line /snap/go/10455/src/net/interface.go:170
			_go_fuzz_dep_.CoverTab[528407]++
//line /snap/go/10455/src/net/interface.go:170
			_go_fuzz_dep_.CoverTab[6109]++
//line /snap/go/10455/src/net/interface.go:170
			// _ = "end of CoverTab[6109]"
//line /snap/go/10455/src/net/interface.go:170
		}
//line /snap/go/10455/src/net/interface.go:170
		// _ = "end of CoverTab[6107]"
	}
//line /snap/go/10455/src/net/interface.go:171
	if _go_fuzz_dep_.CoverTab[786679] == 0 {
//line /snap/go/10455/src/net/interface.go:171
		_go_fuzz_dep_.CoverTab[528438]++
//line /snap/go/10455/src/net/interface.go:171
	} else {
//line /snap/go/10455/src/net/interface.go:171
		_go_fuzz_dep_.CoverTab[528439]++
//line /snap/go/10455/src/net/interface.go:171
	}
//line /snap/go/10455/src/net/interface.go:171
	// _ = "end of CoverTab[6099]"
//line /snap/go/10455/src/net/interface.go:171
	_go_fuzz_dep_.CoverTab[6100]++
						return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errNoSuchInterface}
//line /snap/go/10455/src/net/interface.go:172
	// _ = "end of CoverTab[6100]"
}

// An ipv6ZoneCache represents a cache holding partial network
//line /snap/go/10455/src/net/interface.go:175
// interface information. It is used for reducing the cost of IPv6
//line /snap/go/10455/src/net/interface.go:175
// addressing scope zone resolution.
//line /snap/go/10455/src/net/interface.go:175
//
//line /snap/go/10455/src/net/interface.go:175
// Multiple names sharing the index are managed by first-come
//line /snap/go/10455/src/net/interface.go:175
// first-served basis for consistency.
//line /snap/go/10455/src/net/interface.go:181
type ipv6ZoneCache struct {
	sync.RWMutex			// guard the following
	lastFetched	time.Time	// last time routing information was fetched
	toIndex		map[string]int	// interface name to its index
	toName		map[int]string	// interface index to its name
}

var zoneCache = ipv6ZoneCache{
	toIndex:	make(map[string]int),
	toName:		make(map[int]string),
}

// update refreshes the network interface information if the cache was last
//line /snap/go/10455/src/net/interface.go:193
// updated more than 1 minute ago, or if force is set. It reports whether the
//line /snap/go/10455/src/net/interface.go:193
// cache was updated.
//line /snap/go/10455/src/net/interface.go:196
func (zc *ipv6ZoneCache) update(ift []Interface, force bool) (updated bool) {
//line /snap/go/10455/src/net/interface.go:196
	_go_fuzz_dep_.CoverTab[6110]++
						zc.Lock()
						defer zc.Unlock()
						now := time.Now()
						if !force && func() bool {
//line /snap/go/10455/src/net/interface.go:200
		_go_fuzz_dep_.CoverTab[6114]++
//line /snap/go/10455/src/net/interface.go:200
		return zc.lastFetched.After(now.Add(-60 * time.Second))
//line /snap/go/10455/src/net/interface.go:200
		// _ = "end of CoverTab[6114]"
//line /snap/go/10455/src/net/interface.go:200
	}() {
//line /snap/go/10455/src/net/interface.go:200
		_go_fuzz_dep_.CoverTab[528408]++
//line /snap/go/10455/src/net/interface.go:200
		_go_fuzz_dep_.CoverTab[6115]++
							return false
//line /snap/go/10455/src/net/interface.go:201
		// _ = "end of CoverTab[6115]"
	} else {
//line /snap/go/10455/src/net/interface.go:202
		_go_fuzz_dep_.CoverTab[528409]++
//line /snap/go/10455/src/net/interface.go:202
		_go_fuzz_dep_.CoverTab[6116]++
//line /snap/go/10455/src/net/interface.go:202
		// _ = "end of CoverTab[6116]"
//line /snap/go/10455/src/net/interface.go:202
	}
//line /snap/go/10455/src/net/interface.go:202
	// _ = "end of CoverTab[6110]"
//line /snap/go/10455/src/net/interface.go:202
	_go_fuzz_dep_.CoverTab[6111]++
						zc.lastFetched = now
						if len(ift) == 0 {
//line /snap/go/10455/src/net/interface.go:204
		_go_fuzz_dep_.CoverTab[528410]++
//line /snap/go/10455/src/net/interface.go:204
		_go_fuzz_dep_.CoverTab[6117]++
							var err error
							if ift, err = interfaceTable(0); err != nil {
//line /snap/go/10455/src/net/interface.go:206
			_go_fuzz_dep_.CoverTab[528412]++
//line /snap/go/10455/src/net/interface.go:206
			_go_fuzz_dep_.CoverTab[6118]++
								return false
//line /snap/go/10455/src/net/interface.go:207
			// _ = "end of CoverTab[6118]"
		} else {
//line /snap/go/10455/src/net/interface.go:208
			_go_fuzz_dep_.CoverTab[528413]++
//line /snap/go/10455/src/net/interface.go:208
			_go_fuzz_dep_.CoverTab[6119]++
//line /snap/go/10455/src/net/interface.go:208
			// _ = "end of CoverTab[6119]"
//line /snap/go/10455/src/net/interface.go:208
		}
//line /snap/go/10455/src/net/interface.go:208
		// _ = "end of CoverTab[6117]"
	} else {
//line /snap/go/10455/src/net/interface.go:209
		_go_fuzz_dep_.CoverTab[528411]++
//line /snap/go/10455/src/net/interface.go:209
		_go_fuzz_dep_.CoverTab[6120]++
//line /snap/go/10455/src/net/interface.go:209
		// _ = "end of CoverTab[6120]"
//line /snap/go/10455/src/net/interface.go:209
	}
//line /snap/go/10455/src/net/interface.go:209
	// _ = "end of CoverTab[6111]"
//line /snap/go/10455/src/net/interface.go:209
	_go_fuzz_dep_.CoverTab[6112]++
						zc.toIndex = make(map[string]int, len(ift))
						zc.toName = make(map[int]string, len(ift))
//line /snap/go/10455/src/net/interface.go:211
	_go_fuzz_dep_.CoverTab[786680] = 0
						for _, ifi := range ift {
//line /snap/go/10455/src/net/interface.go:212
		if _go_fuzz_dep_.CoverTab[786680] == 0 {
//line /snap/go/10455/src/net/interface.go:212
			_go_fuzz_dep_.CoverTab[528440]++
//line /snap/go/10455/src/net/interface.go:212
		} else {
//line /snap/go/10455/src/net/interface.go:212
			_go_fuzz_dep_.CoverTab[528441]++
//line /snap/go/10455/src/net/interface.go:212
		}
//line /snap/go/10455/src/net/interface.go:212
		_go_fuzz_dep_.CoverTab[786680] = 1
//line /snap/go/10455/src/net/interface.go:212
		_go_fuzz_dep_.CoverTab[6121]++
							zc.toIndex[ifi.Name] = ifi.Index
							if _, ok := zc.toName[ifi.Index]; !ok {
//line /snap/go/10455/src/net/interface.go:214
			_go_fuzz_dep_.CoverTab[528414]++
//line /snap/go/10455/src/net/interface.go:214
			_go_fuzz_dep_.CoverTab[6122]++
								zc.toName[ifi.Index] = ifi.Name
//line /snap/go/10455/src/net/interface.go:215
			// _ = "end of CoverTab[6122]"
		} else {
//line /snap/go/10455/src/net/interface.go:216
			_go_fuzz_dep_.CoverTab[528415]++
//line /snap/go/10455/src/net/interface.go:216
			_go_fuzz_dep_.CoverTab[6123]++
//line /snap/go/10455/src/net/interface.go:216
			// _ = "end of CoverTab[6123]"
//line /snap/go/10455/src/net/interface.go:216
		}
//line /snap/go/10455/src/net/interface.go:216
		// _ = "end of CoverTab[6121]"
	}
//line /snap/go/10455/src/net/interface.go:217
	if _go_fuzz_dep_.CoverTab[786680] == 0 {
//line /snap/go/10455/src/net/interface.go:217
		_go_fuzz_dep_.CoverTab[528442]++
//line /snap/go/10455/src/net/interface.go:217
	} else {
//line /snap/go/10455/src/net/interface.go:217
		_go_fuzz_dep_.CoverTab[528443]++
//line /snap/go/10455/src/net/interface.go:217
	}
//line /snap/go/10455/src/net/interface.go:217
	// _ = "end of CoverTab[6112]"
//line /snap/go/10455/src/net/interface.go:217
	_go_fuzz_dep_.CoverTab[6113]++
						return true
//line /snap/go/10455/src/net/interface.go:218
	// _ = "end of CoverTab[6113]"
}

func (zc *ipv6ZoneCache) name(index int) string {
//line /snap/go/10455/src/net/interface.go:221
	_go_fuzz_dep_.CoverTab[6124]++
						if index == 0 {
//line /snap/go/10455/src/net/interface.go:222
		_go_fuzz_dep_.CoverTab[528416]++
//line /snap/go/10455/src/net/interface.go:222
		_go_fuzz_dep_.CoverTab[6128]++
							return ""
//line /snap/go/10455/src/net/interface.go:223
		// _ = "end of CoverTab[6128]"
	} else {
//line /snap/go/10455/src/net/interface.go:224
		_go_fuzz_dep_.CoverTab[528417]++
//line /snap/go/10455/src/net/interface.go:224
		_go_fuzz_dep_.CoverTab[6129]++
//line /snap/go/10455/src/net/interface.go:224
		// _ = "end of CoverTab[6129]"
//line /snap/go/10455/src/net/interface.go:224
	}
//line /snap/go/10455/src/net/interface.go:224
	// _ = "end of CoverTab[6124]"
//line /snap/go/10455/src/net/interface.go:224
	_go_fuzz_dep_.CoverTab[6125]++
						updated := zoneCache.update(nil, false)
						zoneCache.RLock()
						name, ok := zoneCache.toName[index]
						zoneCache.RUnlock()
						if !ok && func() bool {
//line /snap/go/10455/src/net/interface.go:229
		_go_fuzz_dep_.CoverTab[6130]++
//line /snap/go/10455/src/net/interface.go:229
		return !updated
//line /snap/go/10455/src/net/interface.go:229
		// _ = "end of CoverTab[6130]"
//line /snap/go/10455/src/net/interface.go:229
	}() {
//line /snap/go/10455/src/net/interface.go:229
		_go_fuzz_dep_.CoverTab[528418]++
//line /snap/go/10455/src/net/interface.go:229
		_go_fuzz_dep_.CoverTab[6131]++
							zoneCache.update(nil, true)
							zoneCache.RLock()
							name, ok = zoneCache.toName[index]
							zoneCache.RUnlock()
//line /snap/go/10455/src/net/interface.go:233
		// _ = "end of CoverTab[6131]"
	} else {
//line /snap/go/10455/src/net/interface.go:234
		_go_fuzz_dep_.CoverTab[528419]++
//line /snap/go/10455/src/net/interface.go:234
		_go_fuzz_dep_.CoverTab[6132]++
//line /snap/go/10455/src/net/interface.go:234
		// _ = "end of CoverTab[6132]"
//line /snap/go/10455/src/net/interface.go:234
	}
//line /snap/go/10455/src/net/interface.go:234
	// _ = "end of CoverTab[6125]"
//line /snap/go/10455/src/net/interface.go:234
	_go_fuzz_dep_.CoverTab[6126]++
						if !ok {
//line /snap/go/10455/src/net/interface.go:235
		_go_fuzz_dep_.CoverTab[528420]++
//line /snap/go/10455/src/net/interface.go:235
		_go_fuzz_dep_.CoverTab[6133]++
							name = itoa.Uitoa(uint(index))
//line /snap/go/10455/src/net/interface.go:236
		// _ = "end of CoverTab[6133]"
	} else {
//line /snap/go/10455/src/net/interface.go:237
		_go_fuzz_dep_.CoverTab[528421]++
//line /snap/go/10455/src/net/interface.go:237
		_go_fuzz_dep_.CoverTab[6134]++
//line /snap/go/10455/src/net/interface.go:237
		// _ = "end of CoverTab[6134]"
//line /snap/go/10455/src/net/interface.go:237
	}
//line /snap/go/10455/src/net/interface.go:237
	// _ = "end of CoverTab[6126]"
//line /snap/go/10455/src/net/interface.go:237
	_go_fuzz_dep_.CoverTab[6127]++
						return name
//line /snap/go/10455/src/net/interface.go:238
	// _ = "end of CoverTab[6127]"
}

func (zc *ipv6ZoneCache) index(name string) int {
//line /snap/go/10455/src/net/interface.go:241
	_go_fuzz_dep_.CoverTab[6135]++
						if name == "" {
//line /snap/go/10455/src/net/interface.go:242
		_go_fuzz_dep_.CoverTab[528422]++
//line /snap/go/10455/src/net/interface.go:242
		_go_fuzz_dep_.CoverTab[6139]++
							return 0
//line /snap/go/10455/src/net/interface.go:243
		// _ = "end of CoverTab[6139]"
	} else {
//line /snap/go/10455/src/net/interface.go:244
		_go_fuzz_dep_.CoverTab[528423]++
//line /snap/go/10455/src/net/interface.go:244
		_go_fuzz_dep_.CoverTab[6140]++
//line /snap/go/10455/src/net/interface.go:244
		// _ = "end of CoverTab[6140]"
//line /snap/go/10455/src/net/interface.go:244
	}
//line /snap/go/10455/src/net/interface.go:244
	// _ = "end of CoverTab[6135]"
//line /snap/go/10455/src/net/interface.go:244
	_go_fuzz_dep_.CoverTab[6136]++
						updated := zoneCache.update(nil, false)
						zoneCache.RLock()
						index, ok := zoneCache.toIndex[name]
						zoneCache.RUnlock()
						if !ok && func() bool {
//line /snap/go/10455/src/net/interface.go:249
		_go_fuzz_dep_.CoverTab[6141]++
//line /snap/go/10455/src/net/interface.go:249
		return !updated
//line /snap/go/10455/src/net/interface.go:249
		// _ = "end of CoverTab[6141]"
//line /snap/go/10455/src/net/interface.go:249
	}() {
//line /snap/go/10455/src/net/interface.go:249
		_go_fuzz_dep_.CoverTab[528424]++
//line /snap/go/10455/src/net/interface.go:249
		_go_fuzz_dep_.CoverTab[6142]++
							zoneCache.update(nil, true)
							zoneCache.RLock()
							index, ok = zoneCache.toIndex[name]
							zoneCache.RUnlock()
//line /snap/go/10455/src/net/interface.go:253
		// _ = "end of CoverTab[6142]"
	} else {
//line /snap/go/10455/src/net/interface.go:254
		_go_fuzz_dep_.CoverTab[528425]++
//line /snap/go/10455/src/net/interface.go:254
		_go_fuzz_dep_.CoverTab[6143]++
//line /snap/go/10455/src/net/interface.go:254
		// _ = "end of CoverTab[6143]"
//line /snap/go/10455/src/net/interface.go:254
	}
//line /snap/go/10455/src/net/interface.go:254
	// _ = "end of CoverTab[6136]"
//line /snap/go/10455/src/net/interface.go:254
	_go_fuzz_dep_.CoverTab[6137]++
						if !ok {
//line /snap/go/10455/src/net/interface.go:255
		_go_fuzz_dep_.CoverTab[528426]++
//line /snap/go/10455/src/net/interface.go:255
		_go_fuzz_dep_.CoverTab[6144]++
							index, _, _ = dtoi(name)
//line /snap/go/10455/src/net/interface.go:256
		// _ = "end of CoverTab[6144]"
	} else {
//line /snap/go/10455/src/net/interface.go:257
		_go_fuzz_dep_.CoverTab[528427]++
//line /snap/go/10455/src/net/interface.go:257
		_go_fuzz_dep_.CoverTab[6145]++
//line /snap/go/10455/src/net/interface.go:257
		// _ = "end of CoverTab[6145]"
//line /snap/go/10455/src/net/interface.go:257
	}
//line /snap/go/10455/src/net/interface.go:257
	// _ = "end of CoverTab[6137]"
//line /snap/go/10455/src/net/interface.go:257
	_go_fuzz_dep_.CoverTab[6138]++
						return index
//line /snap/go/10455/src/net/interface.go:258
	// _ = "end of CoverTab[6138]"
}

//line /snap/go/10455/src/net/interface.go:259
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/interface.go:259
var _ = _go_fuzz_dep_.CoverTab
