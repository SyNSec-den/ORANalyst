// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Minimal RFC 6724 address selection.

//line /snap/go/10455/src/net/addrselect.go:7
package net

//line /snap/go/10455/src/net/addrselect.go:7
import (
//line /snap/go/10455/src/net/addrselect.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/addrselect.go:7
)
//line /snap/go/10455/src/net/addrselect.go:7
import (
//line /snap/go/10455/src/net/addrselect.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/addrselect.go:7
)

import (
	"net/netip"
	"sort"
)

func sortByRFC6724(addrs []IPAddr) {
//line /snap/go/10455/src/net/addrselect.go:14
	_go_fuzz_dep_.CoverTab[4484]++
						if len(addrs) < 2 {
//line /snap/go/10455/src/net/addrselect.go:15
		_go_fuzz_dep_.CoverTab[527369]++
//line /snap/go/10455/src/net/addrselect.go:15
		_go_fuzz_dep_.CoverTab[4486]++
							return
//line /snap/go/10455/src/net/addrselect.go:16
		// _ = "end of CoverTab[4486]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:17
		_go_fuzz_dep_.CoverTab[527370]++
//line /snap/go/10455/src/net/addrselect.go:17
		_go_fuzz_dep_.CoverTab[4487]++
//line /snap/go/10455/src/net/addrselect.go:17
		// _ = "end of CoverTab[4487]"
//line /snap/go/10455/src/net/addrselect.go:17
	}
//line /snap/go/10455/src/net/addrselect.go:17
	// _ = "end of CoverTab[4484]"
//line /snap/go/10455/src/net/addrselect.go:17
	_go_fuzz_dep_.CoverTab[4485]++
						sortByRFC6724withSrcs(addrs, srcAddrs(addrs))
//line /snap/go/10455/src/net/addrselect.go:18
	// _ = "end of CoverTab[4485]"
}

func sortByRFC6724withSrcs(addrs []IPAddr, srcs []netip.Addr) {
//line /snap/go/10455/src/net/addrselect.go:21
	_go_fuzz_dep_.CoverTab[4488]++
						if len(addrs) != len(srcs) {
//line /snap/go/10455/src/net/addrselect.go:22
		_go_fuzz_dep_.CoverTab[527371]++
//line /snap/go/10455/src/net/addrselect.go:22
		_go_fuzz_dep_.CoverTab[4491]++
							panic("internal error")
//line /snap/go/10455/src/net/addrselect.go:23
		// _ = "end of CoverTab[4491]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:24
		_go_fuzz_dep_.CoverTab[527372]++
//line /snap/go/10455/src/net/addrselect.go:24
		_go_fuzz_dep_.CoverTab[4492]++
//line /snap/go/10455/src/net/addrselect.go:24
		// _ = "end of CoverTab[4492]"
//line /snap/go/10455/src/net/addrselect.go:24
	}
//line /snap/go/10455/src/net/addrselect.go:24
	// _ = "end of CoverTab[4488]"
//line /snap/go/10455/src/net/addrselect.go:24
	_go_fuzz_dep_.CoverTab[4489]++
						addrAttr := make([]ipAttr, len(addrs))
						srcAttr := make([]ipAttr, len(srcs))
//line /snap/go/10455/src/net/addrselect.go:26
	_go_fuzz_dep_.CoverTab[786635] = 0
						for i, v := range addrs {
//line /snap/go/10455/src/net/addrselect.go:27
		if _go_fuzz_dep_.CoverTab[786635] == 0 {
//line /snap/go/10455/src/net/addrselect.go:27
			_go_fuzz_dep_.CoverTab[527427]++
//line /snap/go/10455/src/net/addrselect.go:27
		} else {
//line /snap/go/10455/src/net/addrselect.go:27
			_go_fuzz_dep_.CoverTab[527428]++
//line /snap/go/10455/src/net/addrselect.go:27
		}
//line /snap/go/10455/src/net/addrselect.go:27
		_go_fuzz_dep_.CoverTab[786635] = 1
//line /snap/go/10455/src/net/addrselect.go:27
		_go_fuzz_dep_.CoverTab[4493]++
							addrAttrIP, _ := netip.AddrFromSlice(v.IP)
							addrAttr[i] = ipAttrOf(addrAttrIP)
							srcAttr[i] = ipAttrOf(srcs[i])
//line /snap/go/10455/src/net/addrselect.go:30
		// _ = "end of CoverTab[4493]"
	}
//line /snap/go/10455/src/net/addrselect.go:31
	if _go_fuzz_dep_.CoverTab[786635] == 0 {
//line /snap/go/10455/src/net/addrselect.go:31
		_go_fuzz_dep_.CoverTab[527429]++
//line /snap/go/10455/src/net/addrselect.go:31
	} else {
//line /snap/go/10455/src/net/addrselect.go:31
		_go_fuzz_dep_.CoverTab[527430]++
//line /snap/go/10455/src/net/addrselect.go:31
	}
//line /snap/go/10455/src/net/addrselect.go:31
	// _ = "end of CoverTab[4489]"
//line /snap/go/10455/src/net/addrselect.go:31
	_go_fuzz_dep_.CoverTab[4490]++
						sort.Stable(&byRFC6724{
		addrs:		addrs,
		addrAttr:	addrAttr,
		srcs:		srcs,
		srcAttr:	srcAttr,
	})
//line /snap/go/10455/src/net/addrselect.go:37
	// _ = "end of CoverTab[4490]"
}

// srcAddrs tries to UDP-connect to each address to see if it has a
//line /snap/go/10455/src/net/addrselect.go:40
// route. (This doesn't send any packets). The destination port
//line /snap/go/10455/src/net/addrselect.go:40
// number is irrelevant.
//line /snap/go/10455/src/net/addrselect.go:43
func srcAddrs(addrs []IPAddr) []netip.Addr {
//line /snap/go/10455/src/net/addrselect.go:43
	_go_fuzz_dep_.CoverTab[4494]++
						srcs := make([]netip.Addr, len(addrs))
						dst := UDPAddr{Port: 9}
//line /snap/go/10455/src/net/addrselect.go:45
	_go_fuzz_dep_.CoverTab[786636] = 0
						for i := range addrs {
//line /snap/go/10455/src/net/addrselect.go:46
		if _go_fuzz_dep_.CoverTab[786636] == 0 {
//line /snap/go/10455/src/net/addrselect.go:46
			_go_fuzz_dep_.CoverTab[527431]++
//line /snap/go/10455/src/net/addrselect.go:46
		} else {
//line /snap/go/10455/src/net/addrselect.go:46
			_go_fuzz_dep_.CoverTab[527432]++
//line /snap/go/10455/src/net/addrselect.go:46
		}
//line /snap/go/10455/src/net/addrselect.go:46
		_go_fuzz_dep_.CoverTab[786636] = 1
//line /snap/go/10455/src/net/addrselect.go:46
		_go_fuzz_dep_.CoverTab[4496]++
							dst.IP = addrs[i].IP
							dst.Zone = addrs[i].Zone
							c, err := DialUDP("udp", nil, &dst)
							if err == nil {
//line /snap/go/10455/src/net/addrselect.go:50
			_go_fuzz_dep_.CoverTab[527373]++
//line /snap/go/10455/src/net/addrselect.go:50
			_go_fuzz_dep_.CoverTab[4497]++
								if src, ok := c.LocalAddr().(*UDPAddr); ok {
//line /snap/go/10455/src/net/addrselect.go:51
				_go_fuzz_dep_.CoverTab[527375]++
//line /snap/go/10455/src/net/addrselect.go:51
				_go_fuzz_dep_.CoverTab[4499]++
									srcs[i], _ = netip.AddrFromSlice(src.IP)
//line /snap/go/10455/src/net/addrselect.go:52
				// _ = "end of CoverTab[4499]"
			} else {
//line /snap/go/10455/src/net/addrselect.go:53
				_go_fuzz_dep_.CoverTab[527376]++
//line /snap/go/10455/src/net/addrselect.go:53
				_go_fuzz_dep_.CoverTab[4500]++
//line /snap/go/10455/src/net/addrselect.go:53
				// _ = "end of CoverTab[4500]"
//line /snap/go/10455/src/net/addrselect.go:53
			}
//line /snap/go/10455/src/net/addrselect.go:53
			// _ = "end of CoverTab[4497]"
//line /snap/go/10455/src/net/addrselect.go:53
			_go_fuzz_dep_.CoverTab[4498]++
								c.Close()
//line /snap/go/10455/src/net/addrselect.go:54
			// _ = "end of CoverTab[4498]"
		} else {
//line /snap/go/10455/src/net/addrselect.go:55
			_go_fuzz_dep_.CoverTab[527374]++
//line /snap/go/10455/src/net/addrselect.go:55
			_go_fuzz_dep_.CoverTab[4501]++
//line /snap/go/10455/src/net/addrselect.go:55
			// _ = "end of CoverTab[4501]"
//line /snap/go/10455/src/net/addrselect.go:55
		}
//line /snap/go/10455/src/net/addrselect.go:55
		// _ = "end of CoverTab[4496]"
	}
//line /snap/go/10455/src/net/addrselect.go:56
	if _go_fuzz_dep_.CoverTab[786636] == 0 {
//line /snap/go/10455/src/net/addrselect.go:56
		_go_fuzz_dep_.CoverTab[527433]++
//line /snap/go/10455/src/net/addrselect.go:56
	} else {
//line /snap/go/10455/src/net/addrselect.go:56
		_go_fuzz_dep_.CoverTab[527434]++
//line /snap/go/10455/src/net/addrselect.go:56
	}
//line /snap/go/10455/src/net/addrselect.go:56
	// _ = "end of CoverTab[4494]"
//line /snap/go/10455/src/net/addrselect.go:56
	_go_fuzz_dep_.CoverTab[4495]++
						return srcs
//line /snap/go/10455/src/net/addrselect.go:57
	// _ = "end of CoverTab[4495]"
}

type ipAttr struct {
	Scope		scope
	Precedence	uint8
	Label		uint8
}

func ipAttrOf(ip netip.Addr) ipAttr {
//line /snap/go/10455/src/net/addrselect.go:66
	_go_fuzz_dep_.CoverTab[4502]++
						if !ip.IsValid() {
//line /snap/go/10455/src/net/addrselect.go:67
		_go_fuzz_dep_.CoverTab[527377]++
//line /snap/go/10455/src/net/addrselect.go:67
		_go_fuzz_dep_.CoverTab[4504]++
							return ipAttr{}
//line /snap/go/10455/src/net/addrselect.go:68
		// _ = "end of CoverTab[4504]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:69
		_go_fuzz_dep_.CoverTab[527378]++
//line /snap/go/10455/src/net/addrselect.go:69
		_go_fuzz_dep_.CoverTab[4505]++
//line /snap/go/10455/src/net/addrselect.go:69
		// _ = "end of CoverTab[4505]"
//line /snap/go/10455/src/net/addrselect.go:69
	}
//line /snap/go/10455/src/net/addrselect.go:69
	// _ = "end of CoverTab[4502]"
//line /snap/go/10455/src/net/addrselect.go:69
	_go_fuzz_dep_.CoverTab[4503]++
						match := rfc6724policyTable.Classify(ip)
						return ipAttr{
		Scope:		classifyScope(ip),
		Precedence:	match.Precedence,
		Label:		match.Label,
	}
//line /snap/go/10455/src/net/addrselect.go:75
	// _ = "end of CoverTab[4503]"
}

type byRFC6724 struct {
	addrs		[]IPAddr	// addrs to sort
	addrAttr	[]ipAttr
	srcs		[]netip.Addr	// or not valid addr if unreachable
	srcAttr		[]ipAttr
}

func (s *byRFC6724) Len() int {
//line /snap/go/10455/src/net/addrselect.go:85
	_go_fuzz_dep_.CoverTab[4506]++
//line /snap/go/10455/src/net/addrselect.go:85
	return len(s.addrs)
//line /snap/go/10455/src/net/addrselect.go:85
	// _ = "end of CoverTab[4506]"
//line /snap/go/10455/src/net/addrselect.go:85
}

func (s *byRFC6724) Swap(i, j int) {
//line /snap/go/10455/src/net/addrselect.go:87
	_go_fuzz_dep_.CoverTab[4507]++
						s.addrs[i], s.addrs[j] = s.addrs[j], s.addrs[i]
						s.srcs[i], s.srcs[j] = s.srcs[j], s.srcs[i]
						s.addrAttr[i], s.addrAttr[j] = s.addrAttr[j], s.addrAttr[i]
						s.srcAttr[i], s.srcAttr[j] = s.srcAttr[j], s.srcAttr[i]
//line /snap/go/10455/src/net/addrselect.go:91
	// _ = "end of CoverTab[4507]"
}

// Less reports whether i is a better destination address for this
//line /snap/go/10455/src/net/addrselect.go:94
// host than j.
//line /snap/go/10455/src/net/addrselect.go:94
//
//line /snap/go/10455/src/net/addrselect.go:94
// The algorithm and variable names comes from RFC 6724 section 6.
//line /snap/go/10455/src/net/addrselect.go:98
func (s *byRFC6724) Less(i, j int) bool {
//line /snap/go/10455/src/net/addrselect.go:98
		_go_fuzz_dep_.CoverTab[4508]++
							DA := s.addrs[i].IP
							DB := s.addrs[j].IP
							SourceDA := s.srcs[i]
							SourceDB := s.srcs[j]
							attrDA := &s.addrAttr[i]
							attrDB := &s.addrAttr[j]
							attrSourceDA := &s.srcAttr[i]
							attrSourceDB := &s.srcAttr[j]

							const preferDA = true
							const preferDB = false

//line /snap/go/10455/src/net/addrselect.go:115
	if !SourceDA.IsValid() && func() bool {
//line /snap/go/10455/src/net/addrselect.go:115
		_go_fuzz_dep_.CoverTab[4521]++
//line /snap/go/10455/src/net/addrselect.go:115
		return !SourceDB.IsValid()
//line /snap/go/10455/src/net/addrselect.go:115
		// _ = "end of CoverTab[4521]"
//line /snap/go/10455/src/net/addrselect.go:115
	}() {
//line /snap/go/10455/src/net/addrselect.go:115
		_go_fuzz_dep_.CoverTab[527379]++
//line /snap/go/10455/src/net/addrselect.go:115
		_go_fuzz_dep_.CoverTab[4522]++
								return false
//line /snap/go/10455/src/net/addrselect.go:116
		// _ = "end of CoverTab[4522]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:117
		_go_fuzz_dep_.CoverTab[527380]++
//line /snap/go/10455/src/net/addrselect.go:117
		_go_fuzz_dep_.CoverTab[4523]++
//line /snap/go/10455/src/net/addrselect.go:117
		// _ = "end of CoverTab[4523]"
//line /snap/go/10455/src/net/addrselect.go:117
	}
//line /snap/go/10455/src/net/addrselect.go:117
	// _ = "end of CoverTab[4508]"
//line /snap/go/10455/src/net/addrselect.go:117
	_go_fuzz_dep_.CoverTab[4509]++
							if !SourceDB.IsValid() {
//line /snap/go/10455/src/net/addrselect.go:118
		_go_fuzz_dep_.CoverTab[527381]++
//line /snap/go/10455/src/net/addrselect.go:118
		_go_fuzz_dep_.CoverTab[4524]++
								return preferDA
//line /snap/go/10455/src/net/addrselect.go:119
		// _ = "end of CoverTab[4524]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:120
		_go_fuzz_dep_.CoverTab[527382]++
//line /snap/go/10455/src/net/addrselect.go:120
		_go_fuzz_dep_.CoverTab[4525]++
//line /snap/go/10455/src/net/addrselect.go:120
		// _ = "end of CoverTab[4525]"
//line /snap/go/10455/src/net/addrselect.go:120
	}
//line /snap/go/10455/src/net/addrselect.go:120
	// _ = "end of CoverTab[4509]"
//line /snap/go/10455/src/net/addrselect.go:120
	_go_fuzz_dep_.CoverTab[4510]++
							if !SourceDA.IsValid() {
//line /snap/go/10455/src/net/addrselect.go:121
		_go_fuzz_dep_.CoverTab[527383]++
//line /snap/go/10455/src/net/addrselect.go:121
		_go_fuzz_dep_.CoverTab[4526]++
								return preferDB
//line /snap/go/10455/src/net/addrselect.go:122
		// _ = "end of CoverTab[4526]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:123
		_go_fuzz_dep_.CoverTab[527384]++
//line /snap/go/10455/src/net/addrselect.go:123
		_go_fuzz_dep_.CoverTab[4527]++
//line /snap/go/10455/src/net/addrselect.go:123
		// _ = "end of CoverTab[4527]"
//line /snap/go/10455/src/net/addrselect.go:123
	}
//line /snap/go/10455/src/net/addrselect.go:123
	// _ = "end of CoverTab[4510]"
//line /snap/go/10455/src/net/addrselect.go:123
	_go_fuzz_dep_.CoverTab[4511]++

//line /snap/go/10455/src/net/addrselect.go:129
	if attrDA.Scope == attrSourceDA.Scope && func() bool {
//line /snap/go/10455/src/net/addrselect.go:129
		_go_fuzz_dep_.CoverTab[4528]++
//line /snap/go/10455/src/net/addrselect.go:129
		return attrDB.Scope != attrSourceDB.Scope
//line /snap/go/10455/src/net/addrselect.go:129
		// _ = "end of CoverTab[4528]"
//line /snap/go/10455/src/net/addrselect.go:129
	}() {
//line /snap/go/10455/src/net/addrselect.go:129
		_go_fuzz_dep_.CoverTab[527385]++
//line /snap/go/10455/src/net/addrselect.go:129
		_go_fuzz_dep_.CoverTab[4529]++
								return preferDA
//line /snap/go/10455/src/net/addrselect.go:130
		// _ = "end of CoverTab[4529]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:131
		_go_fuzz_dep_.CoverTab[527386]++
//line /snap/go/10455/src/net/addrselect.go:131
		_go_fuzz_dep_.CoverTab[4530]++
//line /snap/go/10455/src/net/addrselect.go:131
		// _ = "end of CoverTab[4530]"
//line /snap/go/10455/src/net/addrselect.go:131
	}
//line /snap/go/10455/src/net/addrselect.go:131
	// _ = "end of CoverTab[4511]"
//line /snap/go/10455/src/net/addrselect.go:131
	_go_fuzz_dep_.CoverTab[4512]++
							if attrDA.Scope != attrSourceDA.Scope && func() bool {
//line /snap/go/10455/src/net/addrselect.go:132
		_go_fuzz_dep_.CoverTab[4531]++
//line /snap/go/10455/src/net/addrselect.go:132
		return attrDB.Scope == attrSourceDB.Scope
//line /snap/go/10455/src/net/addrselect.go:132
		// _ = "end of CoverTab[4531]"
//line /snap/go/10455/src/net/addrselect.go:132
	}() {
//line /snap/go/10455/src/net/addrselect.go:132
		_go_fuzz_dep_.CoverTab[527387]++
//line /snap/go/10455/src/net/addrselect.go:132
		_go_fuzz_dep_.CoverTab[4532]++
								return preferDB
//line /snap/go/10455/src/net/addrselect.go:133
		// _ = "end of CoverTab[4532]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:134
		_go_fuzz_dep_.CoverTab[527388]++
//line /snap/go/10455/src/net/addrselect.go:134
		_go_fuzz_dep_.CoverTab[4533]++
//line /snap/go/10455/src/net/addrselect.go:134
		// _ = "end of CoverTab[4533]"
//line /snap/go/10455/src/net/addrselect.go:134
	}
//line /snap/go/10455/src/net/addrselect.go:134
	// _ = "end of CoverTab[4512]"
//line /snap/go/10455/src/net/addrselect.go:134
	_go_fuzz_dep_.CoverTab[4513]++

//line /snap/go/10455/src/net/addrselect.go:155
	if attrSourceDA.Label == attrDA.Label && func() bool {
//line /snap/go/10455/src/net/addrselect.go:155
		_go_fuzz_dep_.CoverTab[4534]++
//line /snap/go/10455/src/net/addrselect.go:155
		return attrSourceDB.Label != attrDB.Label
								// _ = "end of CoverTab[4534]"
//line /snap/go/10455/src/net/addrselect.go:156
	}() {
//line /snap/go/10455/src/net/addrselect.go:156
		_go_fuzz_dep_.CoverTab[527389]++
//line /snap/go/10455/src/net/addrselect.go:156
		_go_fuzz_dep_.CoverTab[4535]++
								return preferDA
//line /snap/go/10455/src/net/addrselect.go:157
		// _ = "end of CoverTab[4535]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:158
		_go_fuzz_dep_.CoverTab[527390]++
//line /snap/go/10455/src/net/addrselect.go:158
		_go_fuzz_dep_.CoverTab[4536]++
//line /snap/go/10455/src/net/addrselect.go:158
		// _ = "end of CoverTab[4536]"
//line /snap/go/10455/src/net/addrselect.go:158
	}
//line /snap/go/10455/src/net/addrselect.go:158
	// _ = "end of CoverTab[4513]"
//line /snap/go/10455/src/net/addrselect.go:158
	_go_fuzz_dep_.CoverTab[4514]++
							if attrSourceDA.Label != attrDA.Label && func() bool {
//line /snap/go/10455/src/net/addrselect.go:159
		_go_fuzz_dep_.CoverTab[4537]++
//line /snap/go/10455/src/net/addrselect.go:159
		return attrSourceDB.Label == attrDB.Label
								// _ = "end of CoverTab[4537]"
//line /snap/go/10455/src/net/addrselect.go:160
	}() {
//line /snap/go/10455/src/net/addrselect.go:160
		_go_fuzz_dep_.CoverTab[527391]++
//line /snap/go/10455/src/net/addrselect.go:160
		_go_fuzz_dep_.CoverTab[4538]++
								return preferDB
//line /snap/go/10455/src/net/addrselect.go:161
		// _ = "end of CoverTab[4538]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:162
		_go_fuzz_dep_.CoverTab[527392]++
//line /snap/go/10455/src/net/addrselect.go:162
		_go_fuzz_dep_.CoverTab[4539]++
//line /snap/go/10455/src/net/addrselect.go:162
		// _ = "end of CoverTab[4539]"
//line /snap/go/10455/src/net/addrselect.go:162
	}
//line /snap/go/10455/src/net/addrselect.go:162
	// _ = "end of CoverTab[4514]"
//line /snap/go/10455/src/net/addrselect.go:162
	_go_fuzz_dep_.CoverTab[4515]++

//line /snap/go/10455/src/net/addrselect.go:167
	if attrDA.Precedence > attrDB.Precedence {
//line /snap/go/10455/src/net/addrselect.go:167
		_go_fuzz_dep_.CoverTab[527393]++
//line /snap/go/10455/src/net/addrselect.go:167
		_go_fuzz_dep_.CoverTab[4540]++
								return preferDA
//line /snap/go/10455/src/net/addrselect.go:168
		// _ = "end of CoverTab[4540]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:169
		_go_fuzz_dep_.CoverTab[527394]++
//line /snap/go/10455/src/net/addrselect.go:169
		_go_fuzz_dep_.CoverTab[4541]++
//line /snap/go/10455/src/net/addrselect.go:169
		// _ = "end of CoverTab[4541]"
//line /snap/go/10455/src/net/addrselect.go:169
	}
//line /snap/go/10455/src/net/addrselect.go:169
	// _ = "end of CoverTab[4515]"
//line /snap/go/10455/src/net/addrselect.go:169
	_go_fuzz_dep_.CoverTab[4516]++
							if attrDA.Precedence < attrDB.Precedence {
//line /snap/go/10455/src/net/addrselect.go:170
		_go_fuzz_dep_.CoverTab[527395]++
//line /snap/go/10455/src/net/addrselect.go:170
		_go_fuzz_dep_.CoverTab[4542]++
								return preferDB
//line /snap/go/10455/src/net/addrselect.go:171
		// _ = "end of CoverTab[4542]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:172
		_go_fuzz_dep_.CoverTab[527396]++
//line /snap/go/10455/src/net/addrselect.go:172
		_go_fuzz_dep_.CoverTab[4543]++
//line /snap/go/10455/src/net/addrselect.go:172
		// _ = "end of CoverTab[4543]"
//line /snap/go/10455/src/net/addrselect.go:172
	}
//line /snap/go/10455/src/net/addrselect.go:172
	// _ = "end of CoverTab[4516]"
//line /snap/go/10455/src/net/addrselect.go:172
	_go_fuzz_dep_.CoverTab[4517]++

//line /snap/go/10455/src/net/addrselect.go:184
	if attrDA.Scope < attrDB.Scope {
//line /snap/go/10455/src/net/addrselect.go:184
		_go_fuzz_dep_.CoverTab[527397]++
//line /snap/go/10455/src/net/addrselect.go:184
		_go_fuzz_dep_.CoverTab[4544]++
								return preferDA
//line /snap/go/10455/src/net/addrselect.go:185
		// _ = "end of CoverTab[4544]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:186
		_go_fuzz_dep_.CoverTab[527398]++
//line /snap/go/10455/src/net/addrselect.go:186
		_go_fuzz_dep_.CoverTab[4545]++
//line /snap/go/10455/src/net/addrselect.go:186
		// _ = "end of CoverTab[4545]"
//line /snap/go/10455/src/net/addrselect.go:186
	}
//line /snap/go/10455/src/net/addrselect.go:186
	// _ = "end of CoverTab[4517]"
//line /snap/go/10455/src/net/addrselect.go:186
	_go_fuzz_dep_.CoverTab[4518]++
							if attrDA.Scope > attrDB.Scope {
//line /snap/go/10455/src/net/addrselect.go:187
		_go_fuzz_dep_.CoverTab[527399]++
//line /snap/go/10455/src/net/addrselect.go:187
		_go_fuzz_dep_.CoverTab[4546]++
								return preferDB
//line /snap/go/10455/src/net/addrselect.go:188
		// _ = "end of CoverTab[4546]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:189
		_go_fuzz_dep_.CoverTab[527400]++
//line /snap/go/10455/src/net/addrselect.go:189
		_go_fuzz_dep_.CoverTab[4547]++
//line /snap/go/10455/src/net/addrselect.go:189
		// _ = "end of CoverTab[4547]"
//line /snap/go/10455/src/net/addrselect.go:189
	}
//line /snap/go/10455/src/net/addrselect.go:189
	// _ = "end of CoverTab[4518]"
//line /snap/go/10455/src/net/addrselect.go:189
	_go_fuzz_dep_.CoverTab[4519]++

//line /snap/go/10455/src/net/addrselect.go:200
	if DA.To4() == nil && func() bool {
//line /snap/go/10455/src/net/addrselect.go:200
		_go_fuzz_dep_.CoverTab[4548]++
//line /snap/go/10455/src/net/addrselect.go:200
		return DB.To4() == nil
//line /snap/go/10455/src/net/addrselect.go:200
		// _ = "end of CoverTab[4548]"
//line /snap/go/10455/src/net/addrselect.go:200
	}() {
//line /snap/go/10455/src/net/addrselect.go:200
		_go_fuzz_dep_.CoverTab[527401]++
//line /snap/go/10455/src/net/addrselect.go:200
		_go_fuzz_dep_.CoverTab[4549]++
								commonA := commonPrefixLen(SourceDA, DA)
								commonB := commonPrefixLen(SourceDB, DB)

								if commonA > commonB {
//line /snap/go/10455/src/net/addrselect.go:204
			_go_fuzz_dep_.CoverTab[527403]++
//line /snap/go/10455/src/net/addrselect.go:204
			_go_fuzz_dep_.CoverTab[4551]++
									return preferDA
//line /snap/go/10455/src/net/addrselect.go:205
			// _ = "end of CoverTab[4551]"
		} else {
//line /snap/go/10455/src/net/addrselect.go:206
			_go_fuzz_dep_.CoverTab[527404]++
//line /snap/go/10455/src/net/addrselect.go:206
			_go_fuzz_dep_.CoverTab[4552]++
//line /snap/go/10455/src/net/addrselect.go:206
			// _ = "end of CoverTab[4552]"
//line /snap/go/10455/src/net/addrselect.go:206
		}
//line /snap/go/10455/src/net/addrselect.go:206
		// _ = "end of CoverTab[4549]"
//line /snap/go/10455/src/net/addrselect.go:206
		_go_fuzz_dep_.CoverTab[4550]++
								if commonA < commonB {
//line /snap/go/10455/src/net/addrselect.go:207
			_go_fuzz_dep_.CoverTab[527405]++
//line /snap/go/10455/src/net/addrselect.go:207
			_go_fuzz_dep_.CoverTab[4553]++
									return preferDB
//line /snap/go/10455/src/net/addrselect.go:208
			// _ = "end of CoverTab[4553]"
		} else {
//line /snap/go/10455/src/net/addrselect.go:209
			_go_fuzz_dep_.CoverTab[527406]++
//line /snap/go/10455/src/net/addrselect.go:209
			_go_fuzz_dep_.CoverTab[4554]++
//line /snap/go/10455/src/net/addrselect.go:209
			// _ = "end of CoverTab[4554]"
//line /snap/go/10455/src/net/addrselect.go:209
		}
//line /snap/go/10455/src/net/addrselect.go:209
		// _ = "end of CoverTab[4550]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:210
		_go_fuzz_dep_.CoverTab[527402]++
//line /snap/go/10455/src/net/addrselect.go:210
		_go_fuzz_dep_.CoverTab[4555]++
//line /snap/go/10455/src/net/addrselect.go:210
		// _ = "end of CoverTab[4555]"
//line /snap/go/10455/src/net/addrselect.go:210
	}
//line /snap/go/10455/src/net/addrselect.go:210
	// _ = "end of CoverTab[4519]"
//line /snap/go/10455/src/net/addrselect.go:210
	_go_fuzz_dep_.CoverTab[4520]++

//line /snap/go/10455/src/net/addrselect.go:215
	return false
//line /snap/go/10455/src/net/addrselect.go:215
	// _ = "end of CoverTab[4520]"
}

type policyTableEntry struct {
	Prefix		netip.Prefix
	Precedence	uint8
	Label		uint8
}

type policyTable []policyTableEntry

// RFC 6724 section 2.1.
//line /snap/go/10455/src/net/addrselect.go:226
// Items are sorted by the size of their Prefix.Mask.Size,
//line /snap/go/10455/src/net/addrselect.go:228
var rfc6724policyTable = policyTable{
	{

		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}), 128),
		Precedence:	50,
		Label:		0,
	},
	{

//line /snap/go/10455/src/net/addrselect.go:238
		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}), 96),
								Precedence:	35,
								Label:		4,
	},
	{

		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{}), 96),
		Precedence:	1,
		Label:		3,
	},
	{

//line /snap/go/10455/src/net/addrselect.go:251
		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{0x20, 0x01}), 32),
								Precedence:	5,
								Label:		5,
	},
	{

//line /snap/go/10455/src/net/addrselect.go:258
		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{0x20, 0x02}), 16),
								Precedence:	30,
								Label:		2,
	},
	{

		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{0x3f, 0xfe}), 16),
		Precedence:	1,
		Label:		12,
	},
	{

		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{0xfe, 0xc0}), 10),
		Precedence:	1,
		Label:		11,
	},
	{

		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{0xfc}), 7),
		Precedence:	3,
		Label:		13,
	},
	{

		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{}), 0),
		Precedence:	40,
		Label:		1,
	},
}

// Classify returns the policyTableEntry of the entry with the longest
//line /snap/go/10455/src/net/addrselect.go:288
// matching prefix that contains ip.
//line /snap/go/10455/src/net/addrselect.go:288
// The table t must be sorted from largest mask size to smallest.
//line /snap/go/10455/src/net/addrselect.go:291
func (t policyTable) Classify(ip netip.Addr) policyTableEntry {
//line /snap/go/10455/src/net/addrselect.go:291
	_go_fuzz_dep_.CoverTab[4556]++

							if ip.Is4() {
//line /snap/go/10455/src/net/addrselect.go:293
		_go_fuzz_dep_.CoverTab[527407]++
//line /snap/go/10455/src/net/addrselect.go:293
		_go_fuzz_dep_.CoverTab[4559]++
								ip = netip.AddrFrom16(ip.As16())
//line /snap/go/10455/src/net/addrselect.go:294
		// _ = "end of CoverTab[4559]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:295
		_go_fuzz_dep_.CoverTab[527408]++
//line /snap/go/10455/src/net/addrselect.go:295
		_go_fuzz_dep_.CoverTab[4560]++
//line /snap/go/10455/src/net/addrselect.go:295
		// _ = "end of CoverTab[4560]"
//line /snap/go/10455/src/net/addrselect.go:295
	}
//line /snap/go/10455/src/net/addrselect.go:295
	// _ = "end of CoverTab[4556]"
//line /snap/go/10455/src/net/addrselect.go:295
	_go_fuzz_dep_.CoverTab[4557]++
//line /snap/go/10455/src/net/addrselect.go:295
	_go_fuzz_dep_.CoverTab[786637] = 0
							for _, ent := range t {
//line /snap/go/10455/src/net/addrselect.go:296
		if _go_fuzz_dep_.CoverTab[786637] == 0 {
//line /snap/go/10455/src/net/addrselect.go:296
			_go_fuzz_dep_.CoverTab[527435]++
//line /snap/go/10455/src/net/addrselect.go:296
		} else {
//line /snap/go/10455/src/net/addrselect.go:296
			_go_fuzz_dep_.CoverTab[527436]++
//line /snap/go/10455/src/net/addrselect.go:296
		}
//line /snap/go/10455/src/net/addrselect.go:296
		_go_fuzz_dep_.CoverTab[786637] = 1
//line /snap/go/10455/src/net/addrselect.go:296
		_go_fuzz_dep_.CoverTab[4561]++
								if ent.Prefix.Contains(ip) {
//line /snap/go/10455/src/net/addrselect.go:297
			_go_fuzz_dep_.CoverTab[527409]++
//line /snap/go/10455/src/net/addrselect.go:297
			_go_fuzz_dep_.CoverTab[4562]++
									return ent
//line /snap/go/10455/src/net/addrselect.go:298
			// _ = "end of CoverTab[4562]"
		} else {
//line /snap/go/10455/src/net/addrselect.go:299
			_go_fuzz_dep_.CoverTab[527410]++
//line /snap/go/10455/src/net/addrselect.go:299
			_go_fuzz_dep_.CoverTab[4563]++
//line /snap/go/10455/src/net/addrselect.go:299
			// _ = "end of CoverTab[4563]"
//line /snap/go/10455/src/net/addrselect.go:299
		}
//line /snap/go/10455/src/net/addrselect.go:299
		// _ = "end of CoverTab[4561]"
	}
//line /snap/go/10455/src/net/addrselect.go:300
	if _go_fuzz_dep_.CoverTab[786637] == 0 {
//line /snap/go/10455/src/net/addrselect.go:300
		_go_fuzz_dep_.CoverTab[527437]++
//line /snap/go/10455/src/net/addrselect.go:300
	} else {
//line /snap/go/10455/src/net/addrselect.go:300
		_go_fuzz_dep_.CoverTab[527438]++
//line /snap/go/10455/src/net/addrselect.go:300
	}
//line /snap/go/10455/src/net/addrselect.go:300
	// _ = "end of CoverTab[4557]"
//line /snap/go/10455/src/net/addrselect.go:300
	_go_fuzz_dep_.CoverTab[4558]++
							return policyTableEntry{}
//line /snap/go/10455/src/net/addrselect.go:301
	// _ = "end of CoverTab[4558]"
}

// RFC 6724 section 3.1.
type scope uint8

const (
	scopeInterfaceLocal	scope	= 0x1
	scopeLinkLocal		scope	= 0x2
	scopeAdminLocal		scope	= 0x4
	scopeSiteLocal		scope	= 0x5
	scopeOrgLocal		scope	= 0x8
	scopeGlobal		scope	= 0xe
)

func classifyScope(ip netip.Addr) scope {
//line /snap/go/10455/src/net/addrselect.go:316
	_go_fuzz_dep_.CoverTab[4564]++
							if ip.IsLoopback() || func() bool {
//line /snap/go/10455/src/net/addrselect.go:317
		_go_fuzz_dep_.CoverTab[4568]++
//line /snap/go/10455/src/net/addrselect.go:317
		return ip.IsLinkLocalUnicast()
//line /snap/go/10455/src/net/addrselect.go:317
		// _ = "end of CoverTab[4568]"
//line /snap/go/10455/src/net/addrselect.go:317
	}() {
//line /snap/go/10455/src/net/addrselect.go:317
		_go_fuzz_dep_.CoverTab[527411]++
//line /snap/go/10455/src/net/addrselect.go:317
		_go_fuzz_dep_.CoverTab[4569]++
								return scopeLinkLocal
//line /snap/go/10455/src/net/addrselect.go:318
		// _ = "end of CoverTab[4569]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:319
		_go_fuzz_dep_.CoverTab[527412]++
//line /snap/go/10455/src/net/addrselect.go:319
		_go_fuzz_dep_.CoverTab[4570]++
//line /snap/go/10455/src/net/addrselect.go:319
		// _ = "end of CoverTab[4570]"
//line /snap/go/10455/src/net/addrselect.go:319
	}
//line /snap/go/10455/src/net/addrselect.go:319
	// _ = "end of CoverTab[4564]"
//line /snap/go/10455/src/net/addrselect.go:319
	_go_fuzz_dep_.CoverTab[4565]++
							ipv6 := ip.Is6() && func() bool {
//line /snap/go/10455/src/net/addrselect.go:320
		_go_fuzz_dep_.CoverTab[4571]++
//line /snap/go/10455/src/net/addrselect.go:320
		return !ip.Is4In6()
//line /snap/go/10455/src/net/addrselect.go:320
		// _ = "end of CoverTab[4571]"
//line /snap/go/10455/src/net/addrselect.go:320
	}()
							ipv6AsBytes := ip.As16()
							if ipv6 && func() bool {
//line /snap/go/10455/src/net/addrselect.go:322
		_go_fuzz_dep_.CoverTab[4572]++
//line /snap/go/10455/src/net/addrselect.go:322
		return ip.IsMulticast()
//line /snap/go/10455/src/net/addrselect.go:322
		// _ = "end of CoverTab[4572]"
//line /snap/go/10455/src/net/addrselect.go:322
	}() {
//line /snap/go/10455/src/net/addrselect.go:322
		_go_fuzz_dep_.CoverTab[527413]++
//line /snap/go/10455/src/net/addrselect.go:322
		_go_fuzz_dep_.CoverTab[4573]++
								return scope(ipv6AsBytes[1] & 0xf)
//line /snap/go/10455/src/net/addrselect.go:323
		// _ = "end of CoverTab[4573]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:324
		_go_fuzz_dep_.CoverTab[527414]++
//line /snap/go/10455/src/net/addrselect.go:324
		_go_fuzz_dep_.CoverTab[4574]++
//line /snap/go/10455/src/net/addrselect.go:324
		// _ = "end of CoverTab[4574]"
//line /snap/go/10455/src/net/addrselect.go:324
	}
//line /snap/go/10455/src/net/addrselect.go:324
	// _ = "end of CoverTab[4565]"
//line /snap/go/10455/src/net/addrselect.go:324
	_go_fuzz_dep_.CoverTab[4566]++

//line /snap/go/10455/src/net/addrselect.go:327
	if ipv6 && func() bool {
//line /snap/go/10455/src/net/addrselect.go:327
		_go_fuzz_dep_.CoverTab[4575]++
//line /snap/go/10455/src/net/addrselect.go:327
		return ipv6AsBytes[0] == 0xfe
//line /snap/go/10455/src/net/addrselect.go:327
		// _ = "end of CoverTab[4575]"
//line /snap/go/10455/src/net/addrselect.go:327
	}() && func() bool {
//line /snap/go/10455/src/net/addrselect.go:327
		_go_fuzz_dep_.CoverTab[4576]++
//line /snap/go/10455/src/net/addrselect.go:327
		return ipv6AsBytes[1]&0xc0 == 0xc0
//line /snap/go/10455/src/net/addrselect.go:327
		// _ = "end of CoverTab[4576]"
//line /snap/go/10455/src/net/addrselect.go:327
	}() {
//line /snap/go/10455/src/net/addrselect.go:327
		_go_fuzz_dep_.CoverTab[527415]++
//line /snap/go/10455/src/net/addrselect.go:327
		_go_fuzz_dep_.CoverTab[4577]++
								return scopeSiteLocal
//line /snap/go/10455/src/net/addrselect.go:328
		// _ = "end of CoverTab[4577]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:329
		_go_fuzz_dep_.CoverTab[527416]++
//line /snap/go/10455/src/net/addrselect.go:329
		_go_fuzz_dep_.CoverTab[4578]++
//line /snap/go/10455/src/net/addrselect.go:329
		// _ = "end of CoverTab[4578]"
//line /snap/go/10455/src/net/addrselect.go:329
	}
//line /snap/go/10455/src/net/addrselect.go:329
	// _ = "end of CoverTab[4566]"
//line /snap/go/10455/src/net/addrselect.go:329
	_go_fuzz_dep_.CoverTab[4567]++
							return scopeGlobal
//line /snap/go/10455/src/net/addrselect.go:330
	// _ = "end of CoverTab[4567]"
}

// commonPrefixLen reports the length of the longest prefix (looking
//line /snap/go/10455/src/net/addrselect.go:333
// at the most significant, or leftmost, bits) that the
//line /snap/go/10455/src/net/addrselect.go:333
// two addresses have in common, up to the length of a's prefix (i.e.,
//line /snap/go/10455/src/net/addrselect.go:333
// the portion of the address not including the interface ID).
//line /snap/go/10455/src/net/addrselect.go:333
//
//line /snap/go/10455/src/net/addrselect.go:333
// If a or b is an IPv4 address as an IPv6 address, the IPv4 addresses
//line /snap/go/10455/src/net/addrselect.go:333
// are compared (with max common prefix length of 32).
//line /snap/go/10455/src/net/addrselect.go:333
// If a and b are different IP versions, 0 is returned.
//line /snap/go/10455/src/net/addrselect.go:333
//
//line /snap/go/10455/src/net/addrselect.go:333
// See https://tools.ietf.org/html/rfc6724#section-2.2
//line /snap/go/10455/src/net/addrselect.go:343
func commonPrefixLen(a netip.Addr, b IP) (cpl int) {
//line /snap/go/10455/src/net/addrselect.go:343
	_go_fuzz_dep_.CoverTab[4579]++
							if b4 := b.To4(); b4 != nil {
//line /snap/go/10455/src/net/addrselect.go:344
		_go_fuzz_dep_.CoverTab[527417]++
//line /snap/go/10455/src/net/addrselect.go:344
		_go_fuzz_dep_.CoverTab[4584]++
								b = b4
//line /snap/go/10455/src/net/addrselect.go:345
		// _ = "end of CoverTab[4584]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:346
		_go_fuzz_dep_.CoverTab[527418]++
//line /snap/go/10455/src/net/addrselect.go:346
		_go_fuzz_dep_.CoverTab[4585]++
//line /snap/go/10455/src/net/addrselect.go:346
		// _ = "end of CoverTab[4585]"
//line /snap/go/10455/src/net/addrselect.go:346
	}
//line /snap/go/10455/src/net/addrselect.go:346
	// _ = "end of CoverTab[4579]"
//line /snap/go/10455/src/net/addrselect.go:346
	_go_fuzz_dep_.CoverTab[4580]++
							aAsSlice := a.AsSlice()
							if len(aAsSlice) != len(b) {
//line /snap/go/10455/src/net/addrselect.go:348
		_go_fuzz_dep_.CoverTab[527419]++
//line /snap/go/10455/src/net/addrselect.go:348
		_go_fuzz_dep_.CoverTab[4586]++
								return 0
//line /snap/go/10455/src/net/addrselect.go:349
		// _ = "end of CoverTab[4586]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:350
		_go_fuzz_dep_.CoverTab[527420]++
//line /snap/go/10455/src/net/addrselect.go:350
		_go_fuzz_dep_.CoverTab[4587]++
//line /snap/go/10455/src/net/addrselect.go:350
		// _ = "end of CoverTab[4587]"
//line /snap/go/10455/src/net/addrselect.go:350
	}
//line /snap/go/10455/src/net/addrselect.go:350
	// _ = "end of CoverTab[4580]"
//line /snap/go/10455/src/net/addrselect.go:350
	_go_fuzz_dep_.CoverTab[4581]++

							if len(aAsSlice) > 8 {
//line /snap/go/10455/src/net/addrselect.go:352
		_go_fuzz_dep_.CoverTab[527421]++
//line /snap/go/10455/src/net/addrselect.go:352
		_go_fuzz_dep_.CoverTab[4588]++
								aAsSlice = aAsSlice[:8]
								b = b[:8]
//line /snap/go/10455/src/net/addrselect.go:354
		// _ = "end of CoverTab[4588]"
	} else {
//line /snap/go/10455/src/net/addrselect.go:355
		_go_fuzz_dep_.CoverTab[527422]++
//line /snap/go/10455/src/net/addrselect.go:355
		_go_fuzz_dep_.CoverTab[4589]++
//line /snap/go/10455/src/net/addrselect.go:355
		// _ = "end of CoverTab[4589]"
//line /snap/go/10455/src/net/addrselect.go:355
	}
//line /snap/go/10455/src/net/addrselect.go:355
	// _ = "end of CoverTab[4581]"
//line /snap/go/10455/src/net/addrselect.go:355
	_go_fuzz_dep_.CoverTab[4582]++
//line /snap/go/10455/src/net/addrselect.go:355
	_go_fuzz_dep_.CoverTab[786638] = 0
							for len(aAsSlice) > 0 {
//line /snap/go/10455/src/net/addrselect.go:356
		if _go_fuzz_dep_.CoverTab[786638] == 0 {
//line /snap/go/10455/src/net/addrselect.go:356
			_go_fuzz_dep_.CoverTab[527439]++
//line /snap/go/10455/src/net/addrselect.go:356
		} else {
//line /snap/go/10455/src/net/addrselect.go:356
			_go_fuzz_dep_.CoverTab[527440]++
//line /snap/go/10455/src/net/addrselect.go:356
		}
//line /snap/go/10455/src/net/addrselect.go:356
		_go_fuzz_dep_.CoverTab[786638] = 1
//line /snap/go/10455/src/net/addrselect.go:356
		_go_fuzz_dep_.CoverTab[4590]++
								if aAsSlice[0] == b[0] {
//line /snap/go/10455/src/net/addrselect.go:357
			_go_fuzz_dep_.CoverTab[527423]++
//line /snap/go/10455/src/net/addrselect.go:357
			_go_fuzz_dep_.CoverTab[4592]++
									cpl += 8
									aAsSlice = aAsSlice[1:]
									b = b[1:]
									continue
//line /snap/go/10455/src/net/addrselect.go:361
			// _ = "end of CoverTab[4592]"
		} else {
//line /snap/go/10455/src/net/addrselect.go:362
			_go_fuzz_dep_.CoverTab[527424]++
//line /snap/go/10455/src/net/addrselect.go:362
			_go_fuzz_dep_.CoverTab[4593]++
//line /snap/go/10455/src/net/addrselect.go:362
			// _ = "end of CoverTab[4593]"
//line /snap/go/10455/src/net/addrselect.go:362
		}
//line /snap/go/10455/src/net/addrselect.go:362
		// _ = "end of CoverTab[4590]"
//line /snap/go/10455/src/net/addrselect.go:362
		_go_fuzz_dep_.CoverTab[4591]++
								bits := 8
								ab, bb := aAsSlice[0], b[0]
//line /snap/go/10455/src/net/addrselect.go:364
		_go_fuzz_dep_.CoverTab[786639] = 0
								for {
//line /snap/go/10455/src/net/addrselect.go:365
			if _go_fuzz_dep_.CoverTab[786639] == 0 {
//line /snap/go/10455/src/net/addrselect.go:365
				_go_fuzz_dep_.CoverTab[527443]++
//line /snap/go/10455/src/net/addrselect.go:365
			} else {
//line /snap/go/10455/src/net/addrselect.go:365
				_go_fuzz_dep_.CoverTab[527444]++
//line /snap/go/10455/src/net/addrselect.go:365
			}
//line /snap/go/10455/src/net/addrselect.go:365
			_go_fuzz_dep_.CoverTab[786639] = 1
//line /snap/go/10455/src/net/addrselect.go:365
			_go_fuzz_dep_.CoverTab[4594]++
									ab >>= 1
									bb >>= 1
									bits--
									if ab == bb {
//line /snap/go/10455/src/net/addrselect.go:369
				_go_fuzz_dep_.CoverTab[527425]++
//line /snap/go/10455/src/net/addrselect.go:369
				_go_fuzz_dep_.CoverTab[4595]++
										cpl += bits
										return
//line /snap/go/10455/src/net/addrselect.go:371
				// _ = "end of CoverTab[4595]"
			} else {
//line /snap/go/10455/src/net/addrselect.go:372
				_go_fuzz_dep_.CoverTab[527426]++
//line /snap/go/10455/src/net/addrselect.go:372
				_go_fuzz_dep_.CoverTab[4596]++
//line /snap/go/10455/src/net/addrselect.go:372
				// _ = "end of CoverTab[4596]"
//line /snap/go/10455/src/net/addrselect.go:372
			}
//line /snap/go/10455/src/net/addrselect.go:372
			// _ = "end of CoverTab[4594]"
		}
//line /snap/go/10455/src/net/addrselect.go:373
		// _ = "end of CoverTab[4591]"
	}
//line /snap/go/10455/src/net/addrselect.go:374
	if _go_fuzz_dep_.CoverTab[786638] == 0 {
//line /snap/go/10455/src/net/addrselect.go:374
		_go_fuzz_dep_.CoverTab[527441]++
//line /snap/go/10455/src/net/addrselect.go:374
	} else {
//line /snap/go/10455/src/net/addrselect.go:374
		_go_fuzz_dep_.CoverTab[527442]++
//line /snap/go/10455/src/net/addrselect.go:374
	}
//line /snap/go/10455/src/net/addrselect.go:374
	// _ = "end of CoverTab[4582]"
//line /snap/go/10455/src/net/addrselect.go:374
	_go_fuzz_dep_.CoverTab[4583]++
							return
//line /snap/go/10455/src/net/addrselect.go:375
	// _ = "end of CoverTab[4583]"
}

//line /snap/go/10455/src/net/addrselect.go:376
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/addrselect.go:376
var _ = _go_fuzz_dep_.CoverTab
