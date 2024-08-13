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
	_go_fuzz_dep_.CoverTab[5670]++
						s := ""
						for i, name := range flagNames {
//line /usr/local/go/src/net/interface.go:61
		_go_fuzz_dep_.CoverTab[5673]++
							if f&(1<<uint(i)) != 0 {
//line /usr/local/go/src/net/interface.go:62
			_go_fuzz_dep_.CoverTab[5674]++
								if s != "" {
//line /usr/local/go/src/net/interface.go:63
				_go_fuzz_dep_.CoverTab[5676]++
									s += "|"
//line /usr/local/go/src/net/interface.go:64
				// _ = "end of CoverTab[5676]"
			} else {
//line /usr/local/go/src/net/interface.go:65
				_go_fuzz_dep_.CoverTab[5677]++
//line /usr/local/go/src/net/interface.go:65
				// _ = "end of CoverTab[5677]"
//line /usr/local/go/src/net/interface.go:65
			}
//line /usr/local/go/src/net/interface.go:65
			// _ = "end of CoverTab[5674]"
//line /usr/local/go/src/net/interface.go:65
			_go_fuzz_dep_.CoverTab[5675]++
								s += name
//line /usr/local/go/src/net/interface.go:66
			// _ = "end of CoverTab[5675]"
		} else {
//line /usr/local/go/src/net/interface.go:67
			_go_fuzz_dep_.CoverTab[5678]++
//line /usr/local/go/src/net/interface.go:67
			// _ = "end of CoverTab[5678]"
//line /usr/local/go/src/net/interface.go:67
		}
//line /usr/local/go/src/net/interface.go:67
		// _ = "end of CoverTab[5673]"
	}
//line /usr/local/go/src/net/interface.go:68
	// _ = "end of CoverTab[5670]"
//line /usr/local/go/src/net/interface.go:68
	_go_fuzz_dep_.CoverTab[5671]++
						if s == "" {
//line /usr/local/go/src/net/interface.go:69
		_go_fuzz_dep_.CoverTab[5679]++
							s = "0"
//line /usr/local/go/src/net/interface.go:70
		// _ = "end of CoverTab[5679]"
	} else {
//line /usr/local/go/src/net/interface.go:71
		_go_fuzz_dep_.CoverTab[5680]++
//line /usr/local/go/src/net/interface.go:71
		// _ = "end of CoverTab[5680]"
//line /usr/local/go/src/net/interface.go:71
	}
//line /usr/local/go/src/net/interface.go:71
	// _ = "end of CoverTab[5671]"
//line /usr/local/go/src/net/interface.go:71
	_go_fuzz_dep_.CoverTab[5672]++
						return s
//line /usr/local/go/src/net/interface.go:72
	// _ = "end of CoverTab[5672]"
}

// Addrs returns a list of unicast interface addresses for a specific
//line /usr/local/go/src/net/interface.go:75
// interface.
//line /usr/local/go/src/net/interface.go:77
func (ifi *Interface) Addrs() ([]Addr, error) {
//line /usr/local/go/src/net/interface.go:77
	_go_fuzz_dep_.CoverTab[5681]++
						if ifi == nil {
//line /usr/local/go/src/net/interface.go:78
		_go_fuzz_dep_.CoverTab[5684]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterface}
//line /usr/local/go/src/net/interface.go:79
		// _ = "end of CoverTab[5684]"
	} else {
//line /usr/local/go/src/net/interface.go:80
		_go_fuzz_dep_.CoverTab[5685]++
//line /usr/local/go/src/net/interface.go:80
		// _ = "end of CoverTab[5685]"
//line /usr/local/go/src/net/interface.go:80
	}
//line /usr/local/go/src/net/interface.go:80
	// _ = "end of CoverTab[5681]"
//line /usr/local/go/src/net/interface.go:80
	_go_fuzz_dep_.CoverTab[5682]++
						ifat, err := interfaceAddrTable(ifi)
						if err != nil {
//line /usr/local/go/src/net/interface.go:82
		_go_fuzz_dep_.CoverTab[5686]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:83
		// _ = "end of CoverTab[5686]"
	} else {
//line /usr/local/go/src/net/interface.go:84
		_go_fuzz_dep_.CoverTab[5687]++
//line /usr/local/go/src/net/interface.go:84
		// _ = "end of CoverTab[5687]"
//line /usr/local/go/src/net/interface.go:84
	}
//line /usr/local/go/src/net/interface.go:84
	// _ = "end of CoverTab[5682]"
//line /usr/local/go/src/net/interface.go:84
	_go_fuzz_dep_.CoverTab[5683]++
						return ifat, err
//line /usr/local/go/src/net/interface.go:85
	// _ = "end of CoverTab[5683]"
}

// MulticastAddrs returns a list of multicast, joined group addresses
//line /usr/local/go/src/net/interface.go:88
// for a specific interface.
//line /usr/local/go/src/net/interface.go:90
func (ifi *Interface) MulticastAddrs() ([]Addr, error) {
//line /usr/local/go/src/net/interface.go:90
	_go_fuzz_dep_.CoverTab[5688]++
						if ifi == nil {
//line /usr/local/go/src/net/interface.go:91
		_go_fuzz_dep_.CoverTab[5691]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterface}
//line /usr/local/go/src/net/interface.go:92
		// _ = "end of CoverTab[5691]"
	} else {
//line /usr/local/go/src/net/interface.go:93
		_go_fuzz_dep_.CoverTab[5692]++
//line /usr/local/go/src/net/interface.go:93
		// _ = "end of CoverTab[5692]"
//line /usr/local/go/src/net/interface.go:93
	}
//line /usr/local/go/src/net/interface.go:93
	// _ = "end of CoverTab[5688]"
//line /usr/local/go/src/net/interface.go:93
	_go_fuzz_dep_.CoverTab[5689]++
						ifat, err := interfaceMulticastAddrTable(ifi)
						if err != nil {
//line /usr/local/go/src/net/interface.go:95
		_go_fuzz_dep_.CoverTab[5693]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:96
		// _ = "end of CoverTab[5693]"
	} else {
//line /usr/local/go/src/net/interface.go:97
		_go_fuzz_dep_.CoverTab[5694]++
//line /usr/local/go/src/net/interface.go:97
		// _ = "end of CoverTab[5694]"
//line /usr/local/go/src/net/interface.go:97
	}
//line /usr/local/go/src/net/interface.go:97
	// _ = "end of CoverTab[5689]"
//line /usr/local/go/src/net/interface.go:97
	_go_fuzz_dep_.CoverTab[5690]++
						return ifat, err
//line /usr/local/go/src/net/interface.go:98
	// _ = "end of CoverTab[5690]"
}

// Interfaces returns a list of the system's network interfaces.
func Interfaces() ([]Interface, error) {
//line /usr/local/go/src/net/interface.go:102
	_go_fuzz_dep_.CoverTab[5695]++
						ift, err := interfaceTable(0)
						if err != nil {
//line /usr/local/go/src/net/interface.go:104
		_go_fuzz_dep_.CoverTab[5698]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:105
		// _ = "end of CoverTab[5698]"
	} else {
//line /usr/local/go/src/net/interface.go:106
		_go_fuzz_dep_.CoverTab[5699]++
//line /usr/local/go/src/net/interface.go:106
		// _ = "end of CoverTab[5699]"
//line /usr/local/go/src/net/interface.go:106
	}
//line /usr/local/go/src/net/interface.go:106
	// _ = "end of CoverTab[5695]"
//line /usr/local/go/src/net/interface.go:106
	_go_fuzz_dep_.CoverTab[5696]++
						if len(ift) != 0 {
//line /usr/local/go/src/net/interface.go:107
		_go_fuzz_dep_.CoverTab[5700]++
							zoneCache.update(ift, false)
//line /usr/local/go/src/net/interface.go:108
		// _ = "end of CoverTab[5700]"
	} else {
//line /usr/local/go/src/net/interface.go:109
		_go_fuzz_dep_.CoverTab[5701]++
//line /usr/local/go/src/net/interface.go:109
		// _ = "end of CoverTab[5701]"
//line /usr/local/go/src/net/interface.go:109
	}
//line /usr/local/go/src/net/interface.go:109
	// _ = "end of CoverTab[5696]"
//line /usr/local/go/src/net/interface.go:109
	_go_fuzz_dep_.CoverTab[5697]++
						return ift, nil
//line /usr/local/go/src/net/interface.go:110
	// _ = "end of CoverTab[5697]"
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
	_go_fuzz_dep_.CoverTab[5702]++
						ifat, err := interfaceAddrTable(nil)
						if err != nil {
//line /usr/local/go/src/net/interface.go:120
		_go_fuzz_dep_.CoverTab[5704]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:121
		// _ = "end of CoverTab[5704]"
	} else {
//line /usr/local/go/src/net/interface.go:122
		_go_fuzz_dep_.CoverTab[5705]++
//line /usr/local/go/src/net/interface.go:122
		// _ = "end of CoverTab[5705]"
//line /usr/local/go/src/net/interface.go:122
	}
//line /usr/local/go/src/net/interface.go:122
	// _ = "end of CoverTab[5702]"
//line /usr/local/go/src/net/interface.go:122
	_go_fuzz_dep_.CoverTab[5703]++
						return ifat, err
//line /usr/local/go/src/net/interface.go:123
	// _ = "end of CoverTab[5703]"
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
	_go_fuzz_dep_.CoverTab[5706]++
						if index <= 0 {
//line /usr/local/go/src/net/interface.go:132
		_go_fuzz_dep_.CoverTab[5710]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterfaceIndex}
//line /usr/local/go/src/net/interface.go:133
		// _ = "end of CoverTab[5710]"
	} else {
//line /usr/local/go/src/net/interface.go:134
		_go_fuzz_dep_.CoverTab[5711]++
//line /usr/local/go/src/net/interface.go:134
		// _ = "end of CoverTab[5711]"
//line /usr/local/go/src/net/interface.go:134
	}
//line /usr/local/go/src/net/interface.go:134
	// _ = "end of CoverTab[5706]"
//line /usr/local/go/src/net/interface.go:134
	_go_fuzz_dep_.CoverTab[5707]++
						ift, err := interfaceTable(index)
						if err != nil {
//line /usr/local/go/src/net/interface.go:136
		_go_fuzz_dep_.CoverTab[5712]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:137
		// _ = "end of CoverTab[5712]"
	} else {
//line /usr/local/go/src/net/interface.go:138
		_go_fuzz_dep_.CoverTab[5713]++
//line /usr/local/go/src/net/interface.go:138
		// _ = "end of CoverTab[5713]"
//line /usr/local/go/src/net/interface.go:138
	}
//line /usr/local/go/src/net/interface.go:138
	// _ = "end of CoverTab[5707]"
//line /usr/local/go/src/net/interface.go:138
	_go_fuzz_dep_.CoverTab[5708]++
						ifi, err := interfaceByIndex(ift, index)
						if err != nil {
//line /usr/local/go/src/net/interface.go:140
		_go_fuzz_dep_.CoverTab[5714]++
							err = &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:141
		// _ = "end of CoverTab[5714]"
	} else {
//line /usr/local/go/src/net/interface.go:142
		_go_fuzz_dep_.CoverTab[5715]++
//line /usr/local/go/src/net/interface.go:142
		// _ = "end of CoverTab[5715]"
//line /usr/local/go/src/net/interface.go:142
	}
//line /usr/local/go/src/net/interface.go:142
	// _ = "end of CoverTab[5708]"
//line /usr/local/go/src/net/interface.go:142
	_go_fuzz_dep_.CoverTab[5709]++
						return ifi, err
//line /usr/local/go/src/net/interface.go:143
	// _ = "end of CoverTab[5709]"
}

func interfaceByIndex(ift []Interface, index int) (*Interface, error) {
//line /usr/local/go/src/net/interface.go:146
	_go_fuzz_dep_.CoverTab[5716]++
						for _, ifi := range ift {
//line /usr/local/go/src/net/interface.go:147
		_go_fuzz_dep_.CoverTab[5718]++
							if index == ifi.Index {
//line /usr/local/go/src/net/interface.go:148
			_go_fuzz_dep_.CoverTab[5719]++
								return &ifi, nil
//line /usr/local/go/src/net/interface.go:149
			// _ = "end of CoverTab[5719]"
		} else {
//line /usr/local/go/src/net/interface.go:150
			_go_fuzz_dep_.CoverTab[5720]++
//line /usr/local/go/src/net/interface.go:150
			// _ = "end of CoverTab[5720]"
//line /usr/local/go/src/net/interface.go:150
		}
//line /usr/local/go/src/net/interface.go:150
		// _ = "end of CoverTab[5718]"
	}
//line /usr/local/go/src/net/interface.go:151
	// _ = "end of CoverTab[5716]"
//line /usr/local/go/src/net/interface.go:151
	_go_fuzz_dep_.CoverTab[5717]++
						return nil, errNoSuchInterface
//line /usr/local/go/src/net/interface.go:152
	// _ = "end of CoverTab[5717]"
}

// InterfaceByName returns the interface specified by name.
func InterfaceByName(name string) (*Interface, error) {
//line /usr/local/go/src/net/interface.go:156
	_go_fuzz_dep_.CoverTab[5721]++
						if name == "" {
//line /usr/local/go/src/net/interface.go:157
		_go_fuzz_dep_.CoverTab[5726]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterfaceName}
//line /usr/local/go/src/net/interface.go:158
		// _ = "end of CoverTab[5726]"
	} else {
//line /usr/local/go/src/net/interface.go:159
		_go_fuzz_dep_.CoverTab[5727]++
//line /usr/local/go/src/net/interface.go:159
		// _ = "end of CoverTab[5727]"
//line /usr/local/go/src/net/interface.go:159
	}
//line /usr/local/go/src/net/interface.go:159
	// _ = "end of CoverTab[5721]"
//line /usr/local/go/src/net/interface.go:159
	_go_fuzz_dep_.CoverTab[5722]++
						ift, err := interfaceTable(0)
						if err != nil {
//line /usr/local/go/src/net/interface.go:161
		_go_fuzz_dep_.CoverTab[5728]++
							return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
//line /usr/local/go/src/net/interface.go:162
		// _ = "end of CoverTab[5728]"
	} else {
//line /usr/local/go/src/net/interface.go:163
		_go_fuzz_dep_.CoverTab[5729]++
//line /usr/local/go/src/net/interface.go:163
		// _ = "end of CoverTab[5729]"
//line /usr/local/go/src/net/interface.go:163
	}
//line /usr/local/go/src/net/interface.go:163
	// _ = "end of CoverTab[5722]"
//line /usr/local/go/src/net/interface.go:163
	_go_fuzz_dep_.CoverTab[5723]++
						if len(ift) != 0 {
//line /usr/local/go/src/net/interface.go:164
		_go_fuzz_dep_.CoverTab[5730]++
							zoneCache.update(ift, false)
//line /usr/local/go/src/net/interface.go:165
		// _ = "end of CoverTab[5730]"
	} else {
//line /usr/local/go/src/net/interface.go:166
		_go_fuzz_dep_.CoverTab[5731]++
//line /usr/local/go/src/net/interface.go:166
		// _ = "end of CoverTab[5731]"
//line /usr/local/go/src/net/interface.go:166
	}
//line /usr/local/go/src/net/interface.go:166
	// _ = "end of CoverTab[5723]"
//line /usr/local/go/src/net/interface.go:166
	_go_fuzz_dep_.CoverTab[5724]++
						for _, ifi := range ift {
//line /usr/local/go/src/net/interface.go:167
		_go_fuzz_dep_.CoverTab[5732]++
							if name == ifi.Name {
//line /usr/local/go/src/net/interface.go:168
			_go_fuzz_dep_.CoverTab[5733]++
								return &ifi, nil
//line /usr/local/go/src/net/interface.go:169
			// _ = "end of CoverTab[5733]"
		} else {
//line /usr/local/go/src/net/interface.go:170
			_go_fuzz_dep_.CoverTab[5734]++
//line /usr/local/go/src/net/interface.go:170
			// _ = "end of CoverTab[5734]"
//line /usr/local/go/src/net/interface.go:170
		}
//line /usr/local/go/src/net/interface.go:170
		// _ = "end of CoverTab[5732]"
	}
//line /usr/local/go/src/net/interface.go:171
	// _ = "end of CoverTab[5724]"
//line /usr/local/go/src/net/interface.go:171
	_go_fuzz_dep_.CoverTab[5725]++
						return nil, &OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errNoSuchInterface}
//line /usr/local/go/src/net/interface.go:172
	// _ = "end of CoverTab[5725]"
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
	_go_fuzz_dep_.CoverTab[5735]++
						zc.Lock()
						defer zc.Unlock()
						now := time.Now()
						if !force && func() bool {
//line /usr/local/go/src/net/interface.go:200
		_go_fuzz_dep_.CoverTab[5739]++
//line /usr/local/go/src/net/interface.go:200
		return zc.lastFetched.After(now.Add(-60 * time.Second))
//line /usr/local/go/src/net/interface.go:200
		// _ = "end of CoverTab[5739]"
//line /usr/local/go/src/net/interface.go:200
	}() {
//line /usr/local/go/src/net/interface.go:200
		_go_fuzz_dep_.CoverTab[5740]++
							return false
//line /usr/local/go/src/net/interface.go:201
		// _ = "end of CoverTab[5740]"
	} else {
//line /usr/local/go/src/net/interface.go:202
		_go_fuzz_dep_.CoverTab[5741]++
//line /usr/local/go/src/net/interface.go:202
		// _ = "end of CoverTab[5741]"
//line /usr/local/go/src/net/interface.go:202
	}
//line /usr/local/go/src/net/interface.go:202
	// _ = "end of CoverTab[5735]"
//line /usr/local/go/src/net/interface.go:202
	_go_fuzz_dep_.CoverTab[5736]++
						zc.lastFetched = now
						if len(ift) == 0 {
//line /usr/local/go/src/net/interface.go:204
		_go_fuzz_dep_.CoverTab[5742]++
							var err error
							if ift, err = interfaceTable(0); err != nil {
//line /usr/local/go/src/net/interface.go:206
			_go_fuzz_dep_.CoverTab[5743]++
								return false
//line /usr/local/go/src/net/interface.go:207
			// _ = "end of CoverTab[5743]"
		} else {
//line /usr/local/go/src/net/interface.go:208
			_go_fuzz_dep_.CoverTab[5744]++
//line /usr/local/go/src/net/interface.go:208
			// _ = "end of CoverTab[5744]"
//line /usr/local/go/src/net/interface.go:208
		}
//line /usr/local/go/src/net/interface.go:208
		// _ = "end of CoverTab[5742]"
	} else {
//line /usr/local/go/src/net/interface.go:209
		_go_fuzz_dep_.CoverTab[5745]++
//line /usr/local/go/src/net/interface.go:209
		// _ = "end of CoverTab[5745]"
//line /usr/local/go/src/net/interface.go:209
	}
//line /usr/local/go/src/net/interface.go:209
	// _ = "end of CoverTab[5736]"
//line /usr/local/go/src/net/interface.go:209
	_go_fuzz_dep_.CoverTab[5737]++
						zc.toIndex = make(map[string]int, len(ift))
						zc.toName = make(map[int]string, len(ift))
						for _, ifi := range ift {
//line /usr/local/go/src/net/interface.go:212
		_go_fuzz_dep_.CoverTab[5746]++
							zc.toIndex[ifi.Name] = ifi.Index
							if _, ok := zc.toName[ifi.Index]; !ok {
//line /usr/local/go/src/net/interface.go:214
			_go_fuzz_dep_.CoverTab[5747]++
								zc.toName[ifi.Index] = ifi.Name
//line /usr/local/go/src/net/interface.go:215
			// _ = "end of CoverTab[5747]"
		} else {
//line /usr/local/go/src/net/interface.go:216
			_go_fuzz_dep_.CoverTab[5748]++
//line /usr/local/go/src/net/interface.go:216
			// _ = "end of CoverTab[5748]"
//line /usr/local/go/src/net/interface.go:216
		}
//line /usr/local/go/src/net/interface.go:216
		// _ = "end of CoverTab[5746]"
	}
//line /usr/local/go/src/net/interface.go:217
	// _ = "end of CoverTab[5737]"
//line /usr/local/go/src/net/interface.go:217
	_go_fuzz_dep_.CoverTab[5738]++
						return true
//line /usr/local/go/src/net/interface.go:218
	// _ = "end of CoverTab[5738]"
}

func (zc *ipv6ZoneCache) name(index int) string {
//line /usr/local/go/src/net/interface.go:221
	_go_fuzz_dep_.CoverTab[5749]++
						if index == 0 {
//line /usr/local/go/src/net/interface.go:222
		_go_fuzz_dep_.CoverTab[5753]++
							return ""
//line /usr/local/go/src/net/interface.go:223
		// _ = "end of CoverTab[5753]"
	} else {
//line /usr/local/go/src/net/interface.go:224
		_go_fuzz_dep_.CoverTab[5754]++
//line /usr/local/go/src/net/interface.go:224
		// _ = "end of CoverTab[5754]"
//line /usr/local/go/src/net/interface.go:224
	}
//line /usr/local/go/src/net/interface.go:224
	// _ = "end of CoverTab[5749]"
//line /usr/local/go/src/net/interface.go:224
	_go_fuzz_dep_.CoverTab[5750]++
						updated := zoneCache.update(nil, false)
						zoneCache.RLock()
						name, ok := zoneCache.toName[index]
						zoneCache.RUnlock()
						if !ok && func() bool {
//line /usr/local/go/src/net/interface.go:229
		_go_fuzz_dep_.CoverTab[5755]++
//line /usr/local/go/src/net/interface.go:229
		return !updated
//line /usr/local/go/src/net/interface.go:229
		// _ = "end of CoverTab[5755]"
//line /usr/local/go/src/net/interface.go:229
	}() {
//line /usr/local/go/src/net/interface.go:229
		_go_fuzz_dep_.CoverTab[5756]++
							zoneCache.update(nil, true)
							zoneCache.RLock()
							name, ok = zoneCache.toName[index]
							zoneCache.RUnlock()
//line /usr/local/go/src/net/interface.go:233
		// _ = "end of CoverTab[5756]"
	} else {
//line /usr/local/go/src/net/interface.go:234
		_go_fuzz_dep_.CoverTab[5757]++
//line /usr/local/go/src/net/interface.go:234
		// _ = "end of CoverTab[5757]"
//line /usr/local/go/src/net/interface.go:234
	}
//line /usr/local/go/src/net/interface.go:234
	// _ = "end of CoverTab[5750]"
//line /usr/local/go/src/net/interface.go:234
	_go_fuzz_dep_.CoverTab[5751]++
						if !ok {
//line /usr/local/go/src/net/interface.go:235
		_go_fuzz_dep_.CoverTab[5758]++
							name = itoa.Uitoa(uint(index))
//line /usr/local/go/src/net/interface.go:236
		// _ = "end of CoverTab[5758]"
	} else {
//line /usr/local/go/src/net/interface.go:237
		_go_fuzz_dep_.CoverTab[5759]++
//line /usr/local/go/src/net/interface.go:237
		// _ = "end of CoverTab[5759]"
//line /usr/local/go/src/net/interface.go:237
	}
//line /usr/local/go/src/net/interface.go:237
	// _ = "end of CoverTab[5751]"
//line /usr/local/go/src/net/interface.go:237
	_go_fuzz_dep_.CoverTab[5752]++
						return name
//line /usr/local/go/src/net/interface.go:238
	// _ = "end of CoverTab[5752]"
}

func (zc *ipv6ZoneCache) index(name string) int {
//line /usr/local/go/src/net/interface.go:241
	_go_fuzz_dep_.CoverTab[5760]++
						if name == "" {
//line /usr/local/go/src/net/interface.go:242
		_go_fuzz_dep_.CoverTab[5764]++
							return 0
//line /usr/local/go/src/net/interface.go:243
		// _ = "end of CoverTab[5764]"
	} else {
//line /usr/local/go/src/net/interface.go:244
		_go_fuzz_dep_.CoverTab[5765]++
//line /usr/local/go/src/net/interface.go:244
		// _ = "end of CoverTab[5765]"
//line /usr/local/go/src/net/interface.go:244
	}
//line /usr/local/go/src/net/interface.go:244
	// _ = "end of CoverTab[5760]"
//line /usr/local/go/src/net/interface.go:244
	_go_fuzz_dep_.CoverTab[5761]++
						updated := zoneCache.update(nil, false)
						zoneCache.RLock()
						index, ok := zoneCache.toIndex[name]
						zoneCache.RUnlock()
						if !ok && func() bool {
//line /usr/local/go/src/net/interface.go:249
		_go_fuzz_dep_.CoverTab[5766]++
//line /usr/local/go/src/net/interface.go:249
		return !updated
//line /usr/local/go/src/net/interface.go:249
		// _ = "end of CoverTab[5766]"
//line /usr/local/go/src/net/interface.go:249
	}() {
//line /usr/local/go/src/net/interface.go:249
		_go_fuzz_dep_.CoverTab[5767]++
							zoneCache.update(nil, true)
							zoneCache.RLock()
							index, ok = zoneCache.toIndex[name]
							zoneCache.RUnlock()
//line /usr/local/go/src/net/interface.go:253
		// _ = "end of CoverTab[5767]"
	} else {
//line /usr/local/go/src/net/interface.go:254
		_go_fuzz_dep_.CoverTab[5768]++
//line /usr/local/go/src/net/interface.go:254
		// _ = "end of CoverTab[5768]"
//line /usr/local/go/src/net/interface.go:254
	}
//line /usr/local/go/src/net/interface.go:254
	// _ = "end of CoverTab[5761]"
//line /usr/local/go/src/net/interface.go:254
	_go_fuzz_dep_.CoverTab[5762]++
						if !ok {
//line /usr/local/go/src/net/interface.go:255
		_go_fuzz_dep_.CoverTab[5769]++
							index, _, _ = dtoi(name)
//line /usr/local/go/src/net/interface.go:256
		// _ = "end of CoverTab[5769]"
	} else {
//line /usr/local/go/src/net/interface.go:257
		_go_fuzz_dep_.CoverTab[5770]++
//line /usr/local/go/src/net/interface.go:257
		// _ = "end of CoverTab[5770]"
//line /usr/local/go/src/net/interface.go:257
	}
//line /usr/local/go/src/net/interface.go:257
	// _ = "end of CoverTab[5762]"
//line /usr/local/go/src/net/interface.go:257
	_go_fuzz_dep_.CoverTab[5763]++
						return index
//line /usr/local/go/src/net/interface.go:258
	// _ = "end of CoverTab[5763]"
}

//line /usr/local/go/src/net/interface.go:259
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/interface.go:259
var _ = _go_fuzz_dep_.CoverTab
