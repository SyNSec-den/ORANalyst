// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/interface.go:5
package net

//line /usr/local/go/src/net/interface.go:5
import (
//line /usr/local/go/src/net/interface.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/interface.go:5
)
//line /usr/local/go/src/net/interface.go:5
import (
//line /usr/local/go/src/net/interface.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/interface.go:5
)

import (
	"errors"
	"internal/itoa"
	"sync"
	"time"
)

//line /usr/local/go/src/net/interface.go:20
var (
	errInvalidInterface		= errors.New("invalid network interface")
	errInvalidInterfaceIndex	= errors.New("invalid network interface index")
	errInvalidInterfaceName		= errors.New("invalid network interface name")
	errNoSuchInterface		= errors.New("no such network interface")
	errNoSuchMulticastInterface	= errors.New("no such multicast network interface")
)

// Interface represents a mapping between network interface name
//line /usr/local/go/src/net/interface.go:28
// and index. It also represents network interface facility
//line /usr/local/go/src/net/interface.go:28
// information.
//line /usr/local/go/src/net/interface.go:31
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
//line /usr/local/go/src/net/interface.go:59
	_go_fuzz_dep_.CoverTab[14060]++
						s := ""
						for i, name := range flagNames {
//line /usr/local/go/src/net/interface.go:61
		_go_fuzz_dep_.CoverTab[14063]++
							if f&(1<<uint(i)) != 0 {
//line /usr/local/go/src/net/interface.go:62
			_go_fuzz_dep_.CoverTab[14064]++
								if s != "" {
//line /usr/local/go/src/net/interface.go:63
				_go_fuzz_dep_.CoverTab[14066]++
									s += "|"
//line /usr/local/go/src/net/interface.go:64
				// _ = "end of CoverTab[14066]"
			} else {
//line /usr/local/go/src/net/interface.go:65
				_go_fuzz_dep_.CoverTab[14067]++
//line /usr/local/go/src/net/interface.go:65
				// _ = "end of CoverTab[14067]"
//line /usr/local/go/src/net/interface.go:65
			}
//line /usr/local/go/src/net/interface.go:65
			// _ = "end of CoverTab[14064]"
//line /usr/local/go/src/net/interface.go:65
			_go_fuzz_dep_.CoverTab[14065]++
								s += name
//line /usr/local/go/src/net/interface.go:66
			// _ = "end of CoverTab[14065]"
		} else {
//line /usr/local/go/src/net/interface.go:67
			_go_fuzz_dep_.CoverTab[14068]++
//line /usr/local/go/src/net/interface.go:67
			// _ = "end of CoverTab[14068]"
//line /usr/local/go/src/net/interface.go:67
		}
//line /usr/local/go/src/net/interface.go:67
		// _ = "end of CoverTab[14063]"
	}
//line /usr/local/go/src/net/interface.go:68
	// _ = "end of CoverTab[14060]"
//line /usr/local/go/src/net/interface.go:68
	_go_fuzz_dep_.CoverTab[14061]++
						if s == "" {
//line /usr/local/go/src/net/interface.go:69
		_go_fuzz_dep_.CoverTab[14069]++
							s = "0"
//line /usr/local/go/src/net/interface.go:70
		// _ = "end of CoverTab[14069]"
	} else {
//line /usr/local/go/src/net/interface.go:71
		_go_fuzz_dep_.CoverTab[14070]++
//line /usr/local/go/src/net/interface.go:71
		// _ = "end of CoverTab[14070]"
//line /usr/local/go/src/net/interface.go:71
	}
//line /usr/local/go/src/net/interface.go:71
	// _ = "end of CoverTab[14061]"
//line /usr/local/go/src/net/interface.go:71
	_go_fuzz_dep_.CoverTab[14062]++
						return s
//line /usr/local/go/src/net/interface.go:72
	// _ = "end of CoverTab[14062]"
}

// Addrs returns a list of unicast interface addresses for a specific
//line /usr/local/go/src/net/interface.go:75
// interface.
//line /usr/local/go/src/net/interface.go:77
func (ifi *Interface) Addrs() ([]Addr, error) {
//line /usr/local/go/src/net/interface.go:77
	_go_fuzz_dep_.CoverTab[14071]++
						if ifi == nil {
//line /usr/local/go/src/net/interface.go:78
		_go_fuzz_dep_.CoverTab[14074]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterface}
//line /usr/local/go/src/net/interface.go:79
		// _ = "end of CoverTab[14074]"
	} else {
//line /usr/local/go/src/net/interface.go:80
		_go_fuzz_dep_.CoverTab[14075]++
//line /usr/local/go/src/net/interface.go:80
		// _ = "end of CoverTab[14075]"
//line /usr/local/go/src/net/interface.go:80
	}
//line /usr/local/go/src/net/interface.go:80
	// _ = "end of CoverTab[14071]"
//line /usr/local/go/src/net/interface.go:80
	_go_fuzz_dep_.CoverTab[14072]++
						ifat, err := interfaceAddrTable(ifi)
						if err != nil {
//line /usr/local/go/src/net/interface.go:82
		_go_fuzz_dep_.CoverTab[14076]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:83
		// _ = "end of CoverTab[14076]"
	} else {
//line /usr/local/go/src/net/interface.go:84
		_go_fuzz_dep_.CoverTab[14077]++
//line /usr/local/go/src/net/interface.go:84
		// _ = "end of CoverTab[14077]"
//line /usr/local/go/src/net/interface.go:84
	}
//line /usr/local/go/src/net/interface.go:84
	// _ = "end of CoverTab[14072]"
//line /usr/local/go/src/net/interface.go:84
	_go_fuzz_dep_.CoverTab[14073]++
						return ifat, err
//line /usr/local/go/src/net/interface.go:85
	// _ = "end of CoverTab[14073]"
}

// MulticastAddrs returns a list of multicast, joined group addresses
//line /usr/local/go/src/net/interface.go:88
// for a specific interface.
//line /usr/local/go/src/net/interface.go:90
func (ifi *Interface) MulticastAddrs() ([]Addr, error) {
//line /usr/local/go/src/net/interface.go:90
	_go_fuzz_dep_.CoverTab[14078]++
						if ifi == nil {
//line /usr/local/go/src/net/interface.go:91
		_go_fuzz_dep_.CoverTab[14081]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterface}
//line /usr/local/go/src/net/interface.go:92
		// _ = "end of CoverTab[14081]"
	} else {
//line /usr/local/go/src/net/interface.go:93
		_go_fuzz_dep_.CoverTab[14082]++
//line /usr/local/go/src/net/interface.go:93
		// _ = "end of CoverTab[14082]"
//line /usr/local/go/src/net/interface.go:93
	}
//line /usr/local/go/src/net/interface.go:93
	// _ = "end of CoverTab[14078]"
//line /usr/local/go/src/net/interface.go:93
	_go_fuzz_dep_.CoverTab[14079]++
						ifat, err := interfaceMulticastAddrTable(ifi)
						if err != nil {
//line /usr/local/go/src/net/interface.go:95
		_go_fuzz_dep_.CoverTab[14083]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:96
		// _ = "end of CoverTab[14083]"
	} else {
//line /usr/local/go/src/net/interface.go:97
		_go_fuzz_dep_.CoverTab[14084]++
//line /usr/local/go/src/net/interface.go:97
		// _ = "end of CoverTab[14084]"
//line /usr/local/go/src/net/interface.go:97
	}
//line /usr/local/go/src/net/interface.go:97
	// _ = "end of CoverTab[14079]"
//line /usr/local/go/src/net/interface.go:97
	_go_fuzz_dep_.CoverTab[14080]++
						return ifat, err
//line /usr/local/go/src/net/interface.go:98
	// _ = "end of CoverTab[14080]"
}

// Interfaces returns a list of the system's network interfaces.
func Interfaces() ([]Interface, error) {
//line /usr/local/go/src/net/interface.go:102
	_go_fuzz_dep_.CoverTab[14085]++
						ift, err := interfaceTable(0)
						if err != nil {
//line /usr/local/go/src/net/interface.go:104
		_go_fuzz_dep_.CoverTab[14088]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:105
		// _ = "end of CoverTab[14088]"
	} else {
//line /usr/local/go/src/net/interface.go:106
		_go_fuzz_dep_.CoverTab[14089]++
//line /usr/local/go/src/net/interface.go:106
		// _ = "end of CoverTab[14089]"
//line /usr/local/go/src/net/interface.go:106
	}
//line /usr/local/go/src/net/interface.go:106
	// _ = "end of CoverTab[14085]"
//line /usr/local/go/src/net/interface.go:106
	_go_fuzz_dep_.CoverTab[14086]++
						if len(ift) != 0 {
//line /usr/local/go/src/net/interface.go:107
		_go_fuzz_dep_.CoverTab[14090]++
							zoneCache.update(ift, false)
//line /usr/local/go/src/net/interface.go:108
		// _ = "end of CoverTab[14090]"
	} else {
//line /usr/local/go/src/net/interface.go:109
		_go_fuzz_dep_.CoverTab[14091]++
//line /usr/local/go/src/net/interface.go:109
		// _ = "end of CoverTab[14091]"
//line /usr/local/go/src/net/interface.go:109
	}
//line /usr/local/go/src/net/interface.go:109
	// _ = "end of CoverTab[14086]"
//line /usr/local/go/src/net/interface.go:109
	_go_fuzz_dep_.CoverTab[14087]++
						return ift, nil
//line /usr/local/go/src/net/interface.go:110
	// _ = "end of CoverTab[14087]"
}

// InterfaceAddrs returns a list of the system's unicast interface
//line /usr/local/go/src/net/interface.go:113
// addresses.
//line /usr/local/go/src/net/interface.go:113
//
//line /usr/local/go/src/net/interface.go:113
// The returned list does not identify the associated interface; use
//line /usr/local/go/src/net/interface.go:113
// Interfaces and Interface.Addrs for more detail.
//line /usr/local/go/src/net/interface.go:118
func InterfaceAddrs() ([]Addr, error) {
//line /usr/local/go/src/net/interface.go:118
	_go_fuzz_dep_.CoverTab[14092]++
						ifat, err := interfaceAddrTable(nil)
						if err != nil {
//line /usr/local/go/src/net/interface.go:120
		_go_fuzz_dep_.CoverTab[14094]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:121
		// _ = "end of CoverTab[14094]"
	} else {
//line /usr/local/go/src/net/interface.go:122
		_go_fuzz_dep_.CoverTab[14095]++
//line /usr/local/go/src/net/interface.go:122
		// _ = "end of CoverTab[14095]"
//line /usr/local/go/src/net/interface.go:122
	}
//line /usr/local/go/src/net/interface.go:122
	// _ = "end of CoverTab[14092]"
//line /usr/local/go/src/net/interface.go:122
	_go_fuzz_dep_.CoverTab[14093]++
						return ifat, err
//line /usr/local/go/src/net/interface.go:123
	// _ = "end of CoverTab[14093]"
}

// InterfaceByIndex returns the interface specified by index.
//line /usr/local/go/src/net/interface.go:126
//
//line /usr/local/go/src/net/interface.go:126
// On Solaris, it returns one of the logical network interfaces
//line /usr/local/go/src/net/interface.go:126
// sharing the logical data link; for more precision use
//line /usr/local/go/src/net/interface.go:126
// InterfaceByName.
//line /usr/local/go/src/net/interface.go:131
func InterfaceByIndex(index int) (*Interface, error) {
//line /usr/local/go/src/net/interface.go:131
	_go_fuzz_dep_.CoverTab[14096]++
						if index <= 0 {
//line /usr/local/go/src/net/interface.go:132
		_go_fuzz_dep_.CoverTab[14100]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterfaceIndex}
//line /usr/local/go/src/net/interface.go:133
		// _ = "end of CoverTab[14100]"
	} else {
//line /usr/local/go/src/net/interface.go:134
		_go_fuzz_dep_.CoverTab[14101]++
//line /usr/local/go/src/net/interface.go:134
		// _ = "end of CoverTab[14101]"
//line /usr/local/go/src/net/interface.go:134
	}
//line /usr/local/go/src/net/interface.go:134
	// _ = "end of CoverTab[14096]"
//line /usr/local/go/src/net/interface.go:134
	_go_fuzz_dep_.CoverTab[14097]++
						ift, err := interfaceTable(index)
						if err != nil {
//line /usr/local/go/src/net/interface.go:136
		_go_fuzz_dep_.CoverTab[14102]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:137
		// _ = "end of CoverTab[14102]"
	} else {
//line /usr/local/go/src/net/interface.go:138
		_go_fuzz_dep_.CoverTab[14103]++
//line /usr/local/go/src/net/interface.go:138
		// _ = "end of CoverTab[14103]"
//line /usr/local/go/src/net/interface.go:138
	}
//line /usr/local/go/src/net/interface.go:138
	// _ = "end of CoverTab[14097]"
//line /usr/local/go/src/net/interface.go:138
	_go_fuzz_dep_.CoverTab[14098]++
						ifi, err := interfaceByIndex(ift, index)
						if err != nil {
//line /usr/local/go/src/net/interface.go:140
		_go_fuzz_dep_.CoverTab[14104]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:141
		// _ = "end of CoverTab[14104]"
	} else {
//line /usr/local/go/src/net/interface.go:142
		_go_fuzz_dep_.CoverTab[14105]++
//line /usr/local/go/src/net/interface.go:142
		// _ = "end of CoverTab[14105]"
//line /usr/local/go/src/net/interface.go:142
	}
//line /usr/local/go/src/net/interface.go:142
	// _ = "end of CoverTab[14098]"
//line /usr/local/go/src/net/interface.go:142
	_go_fuzz_dep_.CoverTab[14099]++
						return ifi, err
//line /usr/local/go/src/net/interface.go:143
	// _ = "end of CoverTab[14099]"
}

func interfaceByIndex(ift []Interface, index int) (*Interface, error) {
//line /usr/local/go/src/net/interface.go:146
	_go_fuzz_dep_.CoverTab[14106]++
						for _, ifi := range ift {
//line /usr/local/go/src/net/interface.go:147
		_go_fuzz_dep_.CoverTab[14108]++
							if index == ifi.Index {
//line /usr/local/go/src/net/interface.go:148
			_go_fuzz_dep_.CoverTab[14109]++
								return &ifi, nil
//line /usr/local/go/src/net/interface.go:149
			// _ = "end of CoverTab[14109]"
		} else {
//line /usr/local/go/src/net/interface.go:150
			_go_fuzz_dep_.CoverTab[14110]++
//line /usr/local/go/src/net/interface.go:150
			// _ = "end of CoverTab[14110]"
//line /usr/local/go/src/net/interface.go:150
		}
//line /usr/local/go/src/net/interface.go:150
		// _ = "end of CoverTab[14108]"
	}
//line /usr/local/go/src/net/interface.go:151
	// _ = "end of CoverTab[14106]"
//line /usr/local/go/src/net/interface.go:151
	_go_fuzz_dep_.CoverTab[14107]++
						return nil, errNoSuchInterface
//line /usr/local/go/src/net/interface.go:152
	// _ = "end of CoverTab[14107]"
}

// InterfaceByName returns the interface specified by name.
func InterfaceByName(name string) (*Interface, error) {
//line /usr/local/go/src/net/interface.go:156
	_go_fuzz_dep_.CoverTab[14111]++
						if name == "" {
//line /usr/local/go/src/net/interface.go:157
		_go_fuzz_dep_.CoverTab[14116]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterfaceName}
//line /usr/local/go/src/net/interface.go:158
		// _ = "end of CoverTab[14116]"
	} else {
//line /usr/local/go/src/net/interface.go:159
		_go_fuzz_dep_.CoverTab[14117]++
//line /usr/local/go/src/net/interface.go:159
		// _ = "end of CoverTab[14117]"
//line /usr/local/go/src/net/interface.go:159
	}
//line /usr/local/go/src/net/interface.go:159
	// _ = "end of CoverTab[14111]"
//line /usr/local/go/src/net/interface.go:159
	_go_fuzz_dep_.CoverTab[14112]++
						ift, err := interfaceTable(0)
						if err != nil {
//line /usr/local/go/src/net/interface.go:161
		_go_fuzz_dep_.CoverTab[14118]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:162
		// _ = "end of CoverTab[14118]"
	} else {
//line /usr/local/go/src/net/interface.go:163
		_go_fuzz_dep_.CoverTab[14119]++
//line /usr/local/go/src/net/interface.go:163
		// _ = "end of CoverTab[14119]"
//line /usr/local/go/src/net/interface.go:163
	}
//line /usr/local/go/src/net/interface.go:163
	// _ = "end of CoverTab[14112]"
//line /usr/local/go/src/net/interface.go:163
	_go_fuzz_dep_.CoverTab[14113]++
						if len(ift) != 0 {
//line /usr/local/go/src/net/interface.go:164
		_go_fuzz_dep_.CoverTab[14120]++
							zoneCache.update(ift, false)
//line /usr/local/go/src/net/interface.go:165
		// _ = "end of CoverTab[14120]"
	} else {
//line /usr/local/go/src/net/interface.go:166
		_go_fuzz_dep_.CoverTab[14121]++
//line /usr/local/go/src/net/interface.go:166
		// _ = "end of CoverTab[14121]"
//line /usr/local/go/src/net/interface.go:166
	}
//line /usr/local/go/src/net/interface.go:166
	// _ = "end of CoverTab[14113]"
//line /usr/local/go/src/net/interface.go:166
	_go_fuzz_dep_.CoverTab[14114]++
						for _, ifi := range ift {
//line /usr/local/go/src/net/interface.go:167
		_go_fuzz_dep_.CoverTab[14122]++
							if name == ifi.Name {
//line /usr/local/go/src/net/interface.go:168
			_go_fuzz_dep_.CoverTab[14123]++
								return &ifi, nil
//line /usr/local/go/src/net/interface.go:169
			// _ = "end of CoverTab[14123]"
		} else {
//line /usr/local/go/src/net/interface.go:170
			_go_fuzz_dep_.CoverTab[14124]++
//line /usr/local/go/src/net/interface.go:170
			// _ = "end of CoverTab[14124]"
//line /usr/local/go/src/net/interface.go:170
		}
//line /usr/local/go/src/net/interface.go:170
		// _ = "end of CoverTab[14122]"
	}
//line /usr/local/go/src/net/interface.go:171
	// _ = "end of CoverTab[14114]"
//line /usr/local/go/src/net/interface.go:171
	_go_fuzz_dep_.CoverTab[14115]++
						return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errNoSuchInterface}
//line /usr/local/go/src/net/interface.go:172
	// _ = "end of CoverTab[14115]"
}

// An ipv6ZoneCache represents a cache holding partial network
//line /usr/local/go/src/net/interface.go:175
// interface information. It is used for reducing the cost of IPv6
//line /usr/local/go/src/net/interface.go:175
// addressing scope zone resolution.
//line /usr/local/go/src/net/interface.go:175
//
//line /usr/local/go/src/net/interface.go:175
// Multiple names sharing the index are managed by first-come
//line /usr/local/go/src/net/interface.go:175
// first-served basis for consistency.
//line /usr/local/go/src/net/interface.go:181
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
//line /usr/local/go/src/net/interface.go:193
// updated more than 1 minute ago, or if force is set. It reports whether the
//line /usr/local/go/src/net/interface.go:193
// cache was updated.
//line /usr/local/go/src/net/interface.go:196
func (zc *ipv6ZoneCache) update(ift []Interface, force bool) (updated bool) {
//line /usr/local/go/src/net/interface.go:196
	_go_fuzz_dep_.CoverTab[14125]++
						zc.Lock()
						defer zc.Unlock()
						now := time.Now()
						if !force && func() bool {
//line /usr/local/go/src/net/interface.go:200
		_go_fuzz_dep_.CoverTab[14129]++
//line /usr/local/go/src/net/interface.go:200
		return zc.lastFetched.After(now.Add(-60 * time.Second))
//line /usr/local/go/src/net/interface.go:200
		// _ = "end of CoverTab[14129]"
//line /usr/local/go/src/net/interface.go:200
	}() {
//line /usr/local/go/src/net/interface.go:200
		_go_fuzz_dep_.CoverTab[14130]++
							return false
//line /usr/local/go/src/net/interface.go:201
		// _ = "end of CoverTab[14130]"
	} else {
//line /usr/local/go/src/net/interface.go:202
		_go_fuzz_dep_.CoverTab[14131]++
//line /usr/local/go/src/net/interface.go:202
		// _ = "end of CoverTab[14131]"
//line /usr/local/go/src/net/interface.go:202
	}
//line /usr/local/go/src/net/interface.go:202
	// _ = "end of CoverTab[14125]"
//line /usr/local/go/src/net/interface.go:202
	_go_fuzz_dep_.CoverTab[14126]++
						zc.lastFetched = now
						if len(ift) == 0 {
//line /usr/local/go/src/net/interface.go:204
		_go_fuzz_dep_.CoverTab[14132]++
							var err error
							if ift, err = interfaceTable(0); err != nil {
//line /usr/local/go/src/net/interface.go:206
			_go_fuzz_dep_.CoverTab[14133]++
								return false
//line /usr/local/go/src/net/interface.go:207
			// _ = "end of CoverTab[14133]"
		} else {
//line /usr/local/go/src/net/interface.go:208
			_go_fuzz_dep_.CoverTab[14134]++
//line /usr/local/go/src/net/interface.go:208
			// _ = "end of CoverTab[14134]"
//line /usr/local/go/src/net/interface.go:208
		}
//line /usr/local/go/src/net/interface.go:208
		// _ = "end of CoverTab[14132]"
	} else {
//line /usr/local/go/src/net/interface.go:209
		_go_fuzz_dep_.CoverTab[14135]++
//line /usr/local/go/src/net/interface.go:209
		// _ = "end of CoverTab[14135]"
//line /usr/local/go/src/net/interface.go:209
	}
//line /usr/local/go/src/net/interface.go:209
	// _ = "end of CoverTab[14126]"
//line /usr/local/go/src/net/interface.go:209
	_go_fuzz_dep_.CoverTab[14127]++
						zc.toIndex = make(map[string]int, len(ift))
						zc.toName = make(map[int]string, len(ift))
						for _, ifi := range ift {
//line /usr/local/go/src/net/interface.go:212
		_go_fuzz_dep_.CoverTab[14136]++
							zc.toIndex[ifi.Name] = ifi.Index
							if _, ok := zc.toName[ifi.Index]; !ok {
//line /usr/local/go/src/net/interface.go:214
			_go_fuzz_dep_.CoverTab[14137]++
								zc.toName[ifi.Index] = ifi.Name
//line /usr/local/go/src/net/interface.go:215
			// _ = "end of CoverTab[14137]"
		} else {
//line /usr/local/go/src/net/interface.go:216
			_go_fuzz_dep_.CoverTab[14138]++
//line /usr/local/go/src/net/interface.go:216
			// _ = "end of CoverTab[14138]"
//line /usr/local/go/src/net/interface.go:216
		}
//line /usr/local/go/src/net/interface.go:216
		// _ = "end of CoverTab[14136]"
	}
//line /usr/local/go/src/net/interface.go:217
	// _ = "end of CoverTab[14127]"
//line /usr/local/go/src/net/interface.go:217
	_go_fuzz_dep_.CoverTab[14128]++
						return true
//line /usr/local/go/src/net/interface.go:218
	// _ = "end of CoverTab[14128]"
}

func (zc *ipv6ZoneCache) name(index int) string {
//line /usr/local/go/src/net/interface.go:221
	_go_fuzz_dep_.CoverTab[14139]++
						if index == 0 {
//line /usr/local/go/src/net/interface.go:222
		_go_fuzz_dep_.CoverTab[14143]++
							return ""
//line /usr/local/go/src/net/interface.go:223
		// _ = "end of CoverTab[14143]"
	} else {
//line /usr/local/go/src/net/interface.go:224
		_go_fuzz_dep_.CoverTab[14144]++
//line /usr/local/go/src/net/interface.go:224
		// _ = "end of CoverTab[14144]"
//line /usr/local/go/src/net/interface.go:224
	}
//line /usr/local/go/src/net/interface.go:224
	// _ = "end of CoverTab[14139]"
//line /usr/local/go/src/net/interface.go:224
	_go_fuzz_dep_.CoverTab[14140]++
						updated := zoneCache.update(nil, false)
						zoneCache.RLock()
						name, ok := zoneCache.toName[index]
						zoneCache.RUnlock()
						if !ok && func() bool {
//line /usr/local/go/src/net/interface.go:229
		_go_fuzz_dep_.CoverTab[14145]++
//line /usr/local/go/src/net/interface.go:229
		return !updated
//line /usr/local/go/src/net/interface.go:229
		// _ = "end of CoverTab[14145]"
//line /usr/local/go/src/net/interface.go:229
	}() {
//line /usr/local/go/src/net/interface.go:229
		_go_fuzz_dep_.CoverTab[14146]++
							zoneCache.update(nil, true)
							zoneCache.RLock()
							name, ok = zoneCache.toName[index]
							zoneCache.RUnlock()
//line /usr/local/go/src/net/interface.go:233
		// _ = "end of CoverTab[14146]"
	} else {
//line /usr/local/go/src/net/interface.go:234
		_go_fuzz_dep_.CoverTab[14147]++
//line /usr/local/go/src/net/interface.go:234
		// _ = "end of CoverTab[14147]"
//line /usr/local/go/src/net/interface.go:234
	}
//line /usr/local/go/src/net/interface.go:234
	// _ = "end of CoverTab[14140]"
//line /usr/local/go/src/net/interface.go:234
	_go_fuzz_dep_.CoverTab[14141]++
						if !ok {
//line /usr/local/go/src/net/interface.go:235
		_go_fuzz_dep_.CoverTab[14148]++
							name = itoa.Uitoa(uint(index))
//line /usr/local/go/src/net/interface.go:236
		// _ = "end of CoverTab[14148]"
	} else {
//line /usr/local/go/src/net/interface.go:237
		_go_fuzz_dep_.CoverTab[14149]++
//line /usr/local/go/src/net/interface.go:237
		// _ = "end of CoverTab[14149]"
//line /usr/local/go/src/net/interface.go:237
	}
//line /usr/local/go/src/net/interface.go:237
	// _ = "end of CoverTab[14141]"
//line /usr/local/go/src/net/interface.go:237
	_go_fuzz_dep_.CoverTab[14142]++
						return name
//line /usr/local/go/src/net/interface.go:238
	// _ = "end of CoverTab[14142]"
}

func (zc *ipv6ZoneCache) index(name string) int {
//line /usr/local/go/src/net/interface.go:241
	_go_fuzz_dep_.CoverTab[14150]++
						if name == "" {
//line /usr/local/go/src/net/interface.go:242
		_go_fuzz_dep_.CoverTab[14154]++
							return 0
//line /usr/local/go/src/net/interface.go:243
		// _ = "end of CoverTab[14154]"
	} else {
//line /usr/local/go/src/net/interface.go:244
		_go_fuzz_dep_.CoverTab[14155]++
//line /usr/local/go/src/net/interface.go:244
		// _ = "end of CoverTab[14155]"
//line /usr/local/go/src/net/interface.go:244
	}
//line /usr/local/go/src/net/interface.go:244
	// _ = "end of CoverTab[14150]"
//line /usr/local/go/src/net/interface.go:244
	_go_fuzz_dep_.CoverTab[14151]++
						updated := zoneCache.update(nil, false)
						zoneCache.RLock()
						index, ok := zoneCache.toIndex[name]
						zoneCache.RUnlock()
						if !ok && func() bool {
//line /usr/local/go/src/net/interface.go:249
		_go_fuzz_dep_.CoverTab[14156]++
//line /usr/local/go/src/net/interface.go:249
		return !updated
//line /usr/local/go/src/net/interface.go:249
		// _ = "end of CoverTab[14156]"
//line /usr/local/go/src/net/interface.go:249
	}() {
//line /usr/local/go/src/net/interface.go:249
		_go_fuzz_dep_.CoverTab[14157]++
							zoneCache.update(nil, true)
							zoneCache.RLock()
							index, ok = zoneCache.toIndex[name]
							zoneCache.RUnlock()
//line /usr/local/go/src/net/interface.go:253
		// _ = "end of CoverTab[14157]"
	} else {
//line /usr/local/go/src/net/interface.go:254
		_go_fuzz_dep_.CoverTab[14158]++
//line /usr/local/go/src/net/interface.go:254
		// _ = "end of CoverTab[14158]"
//line /usr/local/go/src/net/interface.go:254
	}
//line /usr/local/go/src/net/interface.go:254
	// _ = "end of CoverTab[14151]"
//line /usr/local/go/src/net/interface.go:254
	_go_fuzz_dep_.CoverTab[14152]++
						if !ok {
//line /usr/local/go/src/net/interface.go:255
		_go_fuzz_dep_.CoverTab[14159]++
							index, _, _ = dtoi(name)
//line /usr/local/go/src/net/interface.go:256
		// _ = "end of CoverTab[14159]"
	} else {
//line /usr/local/go/src/net/interface.go:257
		_go_fuzz_dep_.CoverTab[14160]++
//line /usr/local/go/src/net/interface.go:257
		// _ = "end of CoverTab[14160]"
//line /usr/local/go/src/net/interface.go:257
	}
//line /usr/local/go/src/net/interface.go:257
	// _ = "end of CoverTab[14152]"
//line /usr/local/go/src/net/interface.go:257
	_go_fuzz_dep_.CoverTab[14153]++
						return index
//line /usr/local/go/src/net/interface.go:258
	// _ = "end of CoverTab[14153]"
}

//line /usr/local/go/src/net/interface.go:259
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/interface.go:259
var _ = _go_fuzz_dep_.CoverTab
