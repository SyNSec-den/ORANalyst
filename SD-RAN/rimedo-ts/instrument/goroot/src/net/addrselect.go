// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Minimal RFC 6724 address selection.

//line /usr/local/go/src/net/addrselect.go:7
package net

//line /usr/local/go/src/net/addrselect.go:7
import (
//line /usr/local/go/src/net/addrselect.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/addrselect.go:7
)
//line /usr/local/go/src/net/addrselect.go:7
import (
//line /usr/local/go/src/net/addrselect.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/addrselect.go:7
)

import (
	"net/netip"
	"sort"
)

func sortByRFC6724(addrs []IPAddr) {
//line /usr/local/go/src/net/addrselect.go:14
	_go_fuzz_dep_.CoverTab[12542]++
						if len(addrs) < 2 {
//line /usr/local/go/src/net/addrselect.go:15
		_go_fuzz_dep_.CoverTab[12544]++
							return
//line /usr/local/go/src/net/addrselect.go:16
		// _ = "end of CoverTab[12544]"
	} else {
//line /usr/local/go/src/net/addrselect.go:17
		_go_fuzz_dep_.CoverTab[12545]++
//line /usr/local/go/src/net/addrselect.go:17
		// _ = "end of CoverTab[12545]"
//line /usr/local/go/src/net/addrselect.go:17
	}
//line /usr/local/go/src/net/addrselect.go:17
	// _ = "end of CoverTab[12542]"
//line /usr/local/go/src/net/addrselect.go:17
	_go_fuzz_dep_.CoverTab[12543]++
						sortByRFC6724withSrcs(addrs, srcAddrs(addrs))
//line /usr/local/go/src/net/addrselect.go:18
	// _ = "end of CoverTab[12543]"
}

func sortByRFC6724withSrcs(addrs []IPAddr, srcs []netip.Addr) {
//line /usr/local/go/src/net/addrselect.go:21
	_go_fuzz_dep_.CoverTab[12546]++
						if len(addrs) != len(srcs) {
//line /usr/local/go/src/net/addrselect.go:22
		_go_fuzz_dep_.CoverTab[12549]++
							panic("internal error")
//line /usr/local/go/src/net/addrselect.go:23
		// _ = "end of CoverTab[12549]"
	} else {
//line /usr/local/go/src/net/addrselect.go:24
		_go_fuzz_dep_.CoverTab[12550]++
//line /usr/local/go/src/net/addrselect.go:24
		// _ = "end of CoverTab[12550]"
//line /usr/local/go/src/net/addrselect.go:24
	}
//line /usr/local/go/src/net/addrselect.go:24
	// _ = "end of CoverTab[12546]"
//line /usr/local/go/src/net/addrselect.go:24
	_go_fuzz_dep_.CoverTab[12547]++
						addrAttr := make([]ipAttr, len(addrs))
						srcAttr := make([]ipAttr, len(srcs))
						for i, v := range addrs {
//line /usr/local/go/src/net/addrselect.go:27
		_go_fuzz_dep_.CoverTab[12551]++
							addrAttrIP, _ := netip.AddrFromSlice(v.IP)
							addrAttr[i] = ipAttrOf(addrAttrIP)
							srcAttr[i] = ipAttrOf(srcs[i])
//line /usr/local/go/src/net/addrselect.go:30
		// _ = "end of CoverTab[12551]"
	}
//line /usr/local/go/src/net/addrselect.go:31
	// _ = "end of CoverTab[12547]"
//line /usr/local/go/src/net/addrselect.go:31
	_go_fuzz_dep_.CoverTab[12548]++
						sort.Stable(&byRFC6724{
		addrs:		addrs,
		addrAttr:	addrAttr,
		srcs:		srcs,
		srcAttr:	srcAttr,
	})
//line /usr/local/go/src/net/addrselect.go:37
	// _ = "end of CoverTab[12548]"
}

// srcAddrs tries to UDP-connect to each address to see if it has a
//line /usr/local/go/src/net/addrselect.go:40
// route. (This doesn't send any packets). The destination port
//line /usr/local/go/src/net/addrselect.go:40
// number is irrelevant.
//line /usr/local/go/src/net/addrselect.go:43
func srcAddrs(addrs []IPAddr) []netip.Addr {
//line /usr/local/go/src/net/addrselect.go:43
	_go_fuzz_dep_.CoverTab[12552]++
						srcs := make([]netip.Addr, len(addrs))
						dst := UDPAddr{Port: 9}
						for i := range addrs {
//line /usr/local/go/src/net/addrselect.go:46
		_go_fuzz_dep_.CoverTab[12554]++
							dst.IP = addrs[i].IP
							dst.Zone = addrs[i].Zone
							c, err := DialUDP("udp", nil, &dst)
							if err == nil {
//line /usr/local/go/src/net/addrselect.go:50
			_go_fuzz_dep_.CoverTab[12555]++
								if src, ok := c.LocalAddr().(*UDPAddr); ok {
//line /usr/local/go/src/net/addrselect.go:51
				_go_fuzz_dep_.CoverTab[12557]++
									srcs[i], _ = netip.AddrFromSlice(src.IP)
//line /usr/local/go/src/net/addrselect.go:52
				// _ = "end of CoverTab[12557]"
			} else {
//line /usr/local/go/src/net/addrselect.go:53
				_go_fuzz_dep_.CoverTab[12558]++
//line /usr/local/go/src/net/addrselect.go:53
				// _ = "end of CoverTab[12558]"
//line /usr/local/go/src/net/addrselect.go:53
			}
//line /usr/local/go/src/net/addrselect.go:53
			// _ = "end of CoverTab[12555]"
//line /usr/local/go/src/net/addrselect.go:53
			_go_fuzz_dep_.CoverTab[12556]++
								c.Close()
//line /usr/local/go/src/net/addrselect.go:54
			// _ = "end of CoverTab[12556]"
		} else {
//line /usr/local/go/src/net/addrselect.go:55
			_go_fuzz_dep_.CoverTab[12559]++
//line /usr/local/go/src/net/addrselect.go:55
			// _ = "end of CoverTab[12559]"
//line /usr/local/go/src/net/addrselect.go:55
		}
//line /usr/local/go/src/net/addrselect.go:55
		// _ = "end of CoverTab[12554]"
	}
//line /usr/local/go/src/net/addrselect.go:56
	// _ = "end of CoverTab[12552]"
//line /usr/local/go/src/net/addrselect.go:56
	_go_fuzz_dep_.CoverTab[12553]++
						return srcs
//line /usr/local/go/src/net/addrselect.go:57
	// _ = "end of CoverTab[12553]"
}

type ipAttr struct {
	Scope		scope
	Precedence	uint8
	Label		uint8
}

func ipAttrOf(ip netip.Addr) ipAttr {
//line /usr/local/go/src/net/addrselect.go:66
	_go_fuzz_dep_.CoverTab[12560]++
						if !ip.IsValid() {
//line /usr/local/go/src/net/addrselect.go:67
		_go_fuzz_dep_.CoverTab[12562]++
							return ipAttr{}
//line /usr/local/go/src/net/addrselect.go:68
		// _ = "end of CoverTab[12562]"
	} else {
//line /usr/local/go/src/net/addrselect.go:69
		_go_fuzz_dep_.CoverTab[12563]++
//line /usr/local/go/src/net/addrselect.go:69
		// _ = "end of CoverTab[12563]"
//line /usr/local/go/src/net/addrselect.go:69
	}
//line /usr/local/go/src/net/addrselect.go:69
	// _ = "end of CoverTab[12560]"
//line /usr/local/go/src/net/addrselect.go:69
	_go_fuzz_dep_.CoverTab[12561]++
						match := rfc6724policyTable.Classify(ip)
						return ipAttr{
		Scope:		classifyScope(ip),
		Precedence:	match.Precedence,
		Label:		match.Label,
	}
//line /usr/local/go/src/net/addrselect.go:75
	// _ = "end of CoverTab[12561]"
}

type byRFC6724 struct {
	addrs		[]IPAddr	// addrs to sort
	addrAttr	[]ipAttr
	srcs		[]netip.Addr	// or not valid addr if unreachable
	srcAttr		[]ipAttr
}

func (s *byRFC6724) Len() int {
//line /usr/local/go/src/net/addrselect.go:85
	_go_fuzz_dep_.CoverTab[12564]++
//line /usr/local/go/src/net/addrselect.go:85
	return len(s.addrs)
//line /usr/local/go/src/net/addrselect.go:85
	// _ = "end of CoverTab[12564]"
//line /usr/local/go/src/net/addrselect.go:85
}

func (s *byRFC6724) Swap(i, j int) {
//line /usr/local/go/src/net/addrselect.go:87
	_go_fuzz_dep_.CoverTab[12565]++
						s.addrs[i], s.addrs[j] = s.addrs[j], s.addrs[i]
						s.srcs[i], s.srcs[j] = s.srcs[j], s.srcs[i]
						s.addrAttr[i], s.addrAttr[j] = s.addrAttr[j], s.addrAttr[i]
						s.srcAttr[i], s.srcAttr[j] = s.srcAttr[j], s.srcAttr[i]
//line /usr/local/go/src/net/addrselect.go:91
	// _ = "end of CoverTab[12565]"
}

// Less reports whether i is a better destination address for this
//line /usr/local/go/src/net/addrselect.go:94
// host than j.
//line /usr/local/go/src/net/addrselect.go:94
//
//line /usr/local/go/src/net/addrselect.go:94
// The algorithm and variable names comes from RFC 6724 section 6.
//line /usr/local/go/src/net/addrselect.go:98
func (s *byRFC6724) Less(i, j int) bool {
//line /usr/local/go/src/net/addrselect.go:98
	_go_fuzz_dep_.CoverTab[12566]++
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

//line /usr/local/go/src/net/addrselect.go:115
	if !SourceDA.IsValid() && func() bool {
//line /usr/local/go/src/net/addrselect.go:115
		_go_fuzz_dep_.CoverTab[12579]++
//line /usr/local/go/src/net/addrselect.go:115
		return !SourceDB.IsValid()
//line /usr/local/go/src/net/addrselect.go:115
		// _ = "end of CoverTab[12579]"
//line /usr/local/go/src/net/addrselect.go:115
	}() {
//line /usr/local/go/src/net/addrselect.go:115
		_go_fuzz_dep_.CoverTab[12580]++
							return false
//line /usr/local/go/src/net/addrselect.go:116
		// _ = "end of CoverTab[12580]"
	} else {
//line /usr/local/go/src/net/addrselect.go:117
		_go_fuzz_dep_.CoverTab[12581]++
//line /usr/local/go/src/net/addrselect.go:117
		// _ = "end of CoverTab[12581]"
//line /usr/local/go/src/net/addrselect.go:117
	}
//line /usr/local/go/src/net/addrselect.go:117
	// _ = "end of CoverTab[12566]"
//line /usr/local/go/src/net/addrselect.go:117
	_go_fuzz_dep_.CoverTab[12567]++
						if !SourceDB.IsValid() {
//line /usr/local/go/src/net/addrselect.go:118
		_go_fuzz_dep_.CoverTab[12582]++
							return preferDA
//line /usr/local/go/src/net/addrselect.go:119
		// _ = "end of CoverTab[12582]"
	} else {
//line /usr/local/go/src/net/addrselect.go:120
		_go_fuzz_dep_.CoverTab[12583]++
//line /usr/local/go/src/net/addrselect.go:120
		// _ = "end of CoverTab[12583]"
//line /usr/local/go/src/net/addrselect.go:120
	}
//line /usr/local/go/src/net/addrselect.go:120
	// _ = "end of CoverTab[12567]"
//line /usr/local/go/src/net/addrselect.go:120
	_go_fuzz_dep_.CoverTab[12568]++
						if !SourceDA.IsValid() {
//line /usr/local/go/src/net/addrselect.go:121
		_go_fuzz_dep_.CoverTab[12584]++
							return preferDB
//line /usr/local/go/src/net/addrselect.go:122
		// _ = "end of CoverTab[12584]"
	} else {
//line /usr/local/go/src/net/addrselect.go:123
		_go_fuzz_dep_.CoverTab[12585]++
//line /usr/local/go/src/net/addrselect.go:123
		// _ = "end of CoverTab[12585]"
//line /usr/local/go/src/net/addrselect.go:123
	}
//line /usr/local/go/src/net/addrselect.go:123
	// _ = "end of CoverTab[12568]"
//line /usr/local/go/src/net/addrselect.go:123
	_go_fuzz_dep_.CoverTab[12569]++

//line /usr/local/go/src/net/addrselect.go:129
	if attrDA.Scope == attrSourceDA.Scope && func() bool {
//line /usr/local/go/src/net/addrselect.go:129
		_go_fuzz_dep_.CoverTab[12586]++
//line /usr/local/go/src/net/addrselect.go:129
		return attrDB.Scope != attrSourceDB.Scope
//line /usr/local/go/src/net/addrselect.go:129
		// _ = "end of CoverTab[12586]"
//line /usr/local/go/src/net/addrselect.go:129
	}() {
//line /usr/local/go/src/net/addrselect.go:129
		_go_fuzz_dep_.CoverTab[12587]++
							return preferDA
//line /usr/local/go/src/net/addrselect.go:130
		// _ = "end of CoverTab[12587]"
	} else {
//line /usr/local/go/src/net/addrselect.go:131
		_go_fuzz_dep_.CoverTab[12588]++
//line /usr/local/go/src/net/addrselect.go:131
		// _ = "end of CoverTab[12588]"
//line /usr/local/go/src/net/addrselect.go:131
	}
//line /usr/local/go/src/net/addrselect.go:131
	// _ = "end of CoverTab[12569]"
//line /usr/local/go/src/net/addrselect.go:131
	_go_fuzz_dep_.CoverTab[12570]++
						if attrDA.Scope != attrSourceDA.Scope && func() bool {
//line /usr/local/go/src/net/addrselect.go:132
		_go_fuzz_dep_.CoverTab[12589]++
//line /usr/local/go/src/net/addrselect.go:132
		return attrDB.Scope == attrSourceDB.Scope
//line /usr/local/go/src/net/addrselect.go:132
		// _ = "end of CoverTab[12589]"
//line /usr/local/go/src/net/addrselect.go:132
	}() {
//line /usr/local/go/src/net/addrselect.go:132
		_go_fuzz_dep_.CoverTab[12590]++
							return preferDB
//line /usr/local/go/src/net/addrselect.go:133
		// _ = "end of CoverTab[12590]"
	} else {
//line /usr/local/go/src/net/addrselect.go:134
		_go_fuzz_dep_.CoverTab[12591]++
//line /usr/local/go/src/net/addrselect.go:134
		// _ = "end of CoverTab[12591]"
//line /usr/local/go/src/net/addrselect.go:134
	}
//line /usr/local/go/src/net/addrselect.go:134
	// _ = "end of CoverTab[12570]"
//line /usr/local/go/src/net/addrselect.go:134
	_go_fuzz_dep_.CoverTab[12571]++

//line /usr/local/go/src/net/addrselect.go:155
	if attrSourceDA.Label == attrDA.Label && func() bool {
//line /usr/local/go/src/net/addrselect.go:155
		_go_fuzz_dep_.CoverTab[12592]++
//line /usr/local/go/src/net/addrselect.go:155
		return attrSourceDB.Label != attrDB.Label
							// _ = "end of CoverTab[12592]"
//line /usr/local/go/src/net/addrselect.go:156
	}() {
//line /usr/local/go/src/net/addrselect.go:156
		_go_fuzz_dep_.CoverTab[12593]++
							return preferDA
//line /usr/local/go/src/net/addrselect.go:157
		// _ = "end of CoverTab[12593]"
	} else {
//line /usr/local/go/src/net/addrselect.go:158
		_go_fuzz_dep_.CoverTab[12594]++
//line /usr/local/go/src/net/addrselect.go:158
		// _ = "end of CoverTab[12594]"
//line /usr/local/go/src/net/addrselect.go:158
	}
//line /usr/local/go/src/net/addrselect.go:158
	// _ = "end of CoverTab[12571]"
//line /usr/local/go/src/net/addrselect.go:158
	_go_fuzz_dep_.CoverTab[12572]++
						if attrSourceDA.Label != attrDA.Label && func() bool {
//line /usr/local/go/src/net/addrselect.go:159
		_go_fuzz_dep_.CoverTab[12595]++
//line /usr/local/go/src/net/addrselect.go:159
		return attrSourceDB.Label == attrDB.Label
							// _ = "end of CoverTab[12595]"
//line /usr/local/go/src/net/addrselect.go:160
	}() {
//line /usr/local/go/src/net/addrselect.go:160
		_go_fuzz_dep_.CoverTab[12596]++
							return preferDB
//line /usr/local/go/src/net/addrselect.go:161
		// _ = "end of CoverTab[12596]"
	} else {
//line /usr/local/go/src/net/addrselect.go:162
		_go_fuzz_dep_.CoverTab[12597]++
//line /usr/local/go/src/net/addrselect.go:162
		// _ = "end of CoverTab[12597]"
//line /usr/local/go/src/net/addrselect.go:162
	}
//line /usr/local/go/src/net/addrselect.go:162
	// _ = "end of CoverTab[12572]"
//line /usr/local/go/src/net/addrselect.go:162
	_go_fuzz_dep_.CoverTab[12573]++

//line /usr/local/go/src/net/addrselect.go:167
	if attrDA.Precedence > attrDB.Precedence {
//line /usr/local/go/src/net/addrselect.go:167
		_go_fuzz_dep_.CoverTab[12598]++
							return preferDA
//line /usr/local/go/src/net/addrselect.go:168
		// _ = "end of CoverTab[12598]"
	} else {
//line /usr/local/go/src/net/addrselect.go:169
		_go_fuzz_dep_.CoverTab[12599]++
//line /usr/local/go/src/net/addrselect.go:169
		// _ = "end of CoverTab[12599]"
//line /usr/local/go/src/net/addrselect.go:169
	}
//line /usr/local/go/src/net/addrselect.go:169
	// _ = "end of CoverTab[12573]"
//line /usr/local/go/src/net/addrselect.go:169
	_go_fuzz_dep_.CoverTab[12574]++
						if attrDA.Precedence < attrDB.Precedence {
//line /usr/local/go/src/net/addrselect.go:170
		_go_fuzz_dep_.CoverTab[12600]++
							return preferDB
//line /usr/local/go/src/net/addrselect.go:171
		// _ = "end of CoverTab[12600]"
	} else {
//line /usr/local/go/src/net/addrselect.go:172
		_go_fuzz_dep_.CoverTab[12601]++
//line /usr/local/go/src/net/addrselect.go:172
		// _ = "end of CoverTab[12601]"
//line /usr/local/go/src/net/addrselect.go:172
	}
//line /usr/local/go/src/net/addrselect.go:172
	// _ = "end of CoverTab[12574]"
//line /usr/local/go/src/net/addrselect.go:172
	_go_fuzz_dep_.CoverTab[12575]++

//line /usr/local/go/src/net/addrselect.go:184
	if attrDA.Scope < attrDB.Scope {
//line /usr/local/go/src/net/addrselect.go:184
		_go_fuzz_dep_.CoverTab[12602]++
							return preferDA
//line /usr/local/go/src/net/addrselect.go:185
		// _ = "end of CoverTab[12602]"
	} else {
//line /usr/local/go/src/net/addrselect.go:186
		_go_fuzz_dep_.CoverTab[12603]++
//line /usr/local/go/src/net/addrselect.go:186
		// _ = "end of CoverTab[12603]"
//line /usr/local/go/src/net/addrselect.go:186
	}
//line /usr/local/go/src/net/addrselect.go:186
	// _ = "end of CoverTab[12575]"
//line /usr/local/go/src/net/addrselect.go:186
	_go_fuzz_dep_.CoverTab[12576]++
						if attrDA.Scope > attrDB.Scope {
//line /usr/local/go/src/net/addrselect.go:187
		_go_fuzz_dep_.CoverTab[12604]++
							return preferDB
//line /usr/local/go/src/net/addrselect.go:188
		// _ = "end of CoverTab[12604]"
	} else {
//line /usr/local/go/src/net/addrselect.go:189
		_go_fuzz_dep_.CoverTab[12605]++
//line /usr/local/go/src/net/addrselect.go:189
		// _ = "end of CoverTab[12605]"
//line /usr/local/go/src/net/addrselect.go:189
	}
//line /usr/local/go/src/net/addrselect.go:189
	// _ = "end of CoverTab[12576]"
//line /usr/local/go/src/net/addrselect.go:189
	_go_fuzz_dep_.CoverTab[12577]++

//line /usr/local/go/src/net/addrselect.go:200
	if DA.To4() == nil && func() bool {
//line /usr/local/go/src/net/addrselect.go:200
		_go_fuzz_dep_.CoverTab[12606]++
//line /usr/local/go/src/net/addrselect.go:200
		return DB.To4() == nil
//line /usr/local/go/src/net/addrselect.go:200
		// _ = "end of CoverTab[12606]"
//line /usr/local/go/src/net/addrselect.go:200
	}() {
//line /usr/local/go/src/net/addrselect.go:200
		_go_fuzz_dep_.CoverTab[12607]++
							commonA := commonPrefixLen(SourceDA, DA)
							commonB := commonPrefixLen(SourceDB, DB)

							if commonA > commonB {
//line /usr/local/go/src/net/addrselect.go:204
			_go_fuzz_dep_.CoverTab[12609]++
								return preferDA
//line /usr/local/go/src/net/addrselect.go:205
			// _ = "end of CoverTab[12609]"
		} else {
//line /usr/local/go/src/net/addrselect.go:206
			_go_fuzz_dep_.CoverTab[12610]++
//line /usr/local/go/src/net/addrselect.go:206
			// _ = "end of CoverTab[12610]"
//line /usr/local/go/src/net/addrselect.go:206
		}
//line /usr/local/go/src/net/addrselect.go:206
		// _ = "end of CoverTab[12607]"
//line /usr/local/go/src/net/addrselect.go:206
		_go_fuzz_dep_.CoverTab[12608]++
							if commonA < commonB {
//line /usr/local/go/src/net/addrselect.go:207
			_go_fuzz_dep_.CoverTab[12611]++
								return preferDB
//line /usr/local/go/src/net/addrselect.go:208
			// _ = "end of CoverTab[12611]"
		} else {
//line /usr/local/go/src/net/addrselect.go:209
			_go_fuzz_dep_.CoverTab[12612]++
//line /usr/local/go/src/net/addrselect.go:209
			// _ = "end of CoverTab[12612]"
//line /usr/local/go/src/net/addrselect.go:209
		}
//line /usr/local/go/src/net/addrselect.go:209
		// _ = "end of CoverTab[12608]"
	} else {
//line /usr/local/go/src/net/addrselect.go:210
		_go_fuzz_dep_.CoverTab[12613]++
//line /usr/local/go/src/net/addrselect.go:210
		// _ = "end of CoverTab[12613]"
//line /usr/local/go/src/net/addrselect.go:210
	}
//line /usr/local/go/src/net/addrselect.go:210
	// _ = "end of CoverTab[12577]"
//line /usr/local/go/src/net/addrselect.go:210
	_go_fuzz_dep_.CoverTab[12578]++

//line /usr/local/go/src/net/addrselect.go:215
	return false
//line /usr/local/go/src/net/addrselect.go:215
	// _ = "end of CoverTab[12578]"
}

type policyTableEntry struct {
	Prefix		netip.Prefix
	Precedence	uint8
	Label		uint8
}

type policyTable []policyTableEntry

// RFC 6724 section 2.1.
//line /usr/local/go/src/net/addrselect.go:226
// Items are sorted by the size of their Prefix.Mask.Size,
//line /usr/local/go/src/net/addrselect.go:228
var rfc6724policyTable = policyTable{
	{

		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}), 128),
		Precedence:	50,
		Label:		0,
	},
	{

//line /usr/local/go/src/net/addrselect.go:238
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

//line /usr/local/go/src/net/addrselect.go:251
		Prefix:		netip.PrefixFrom(netip.AddrFrom16([16]byte{0x20, 0x01}), 32),
							Precedence:	5,
							Label:		5,
	},
	{

//line /usr/local/go/src/net/addrselect.go:258
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
//line /usr/local/go/src/net/addrselect.go:288
// matching prefix that contains ip.
//line /usr/local/go/src/net/addrselect.go:288
// The table t must be sorted from largest mask size to smallest.
//line /usr/local/go/src/net/addrselect.go:291
func (t policyTable) Classify(ip netip.Addr) policyTableEntry {
//line /usr/local/go/src/net/addrselect.go:291
	_go_fuzz_dep_.CoverTab[12614]++

						if ip.Is4() {
//line /usr/local/go/src/net/addrselect.go:293
		_go_fuzz_dep_.CoverTab[12617]++
							ip = netip.AddrFrom16(ip.As16())
//line /usr/local/go/src/net/addrselect.go:294
		// _ = "end of CoverTab[12617]"
	} else {
//line /usr/local/go/src/net/addrselect.go:295
		_go_fuzz_dep_.CoverTab[12618]++
//line /usr/local/go/src/net/addrselect.go:295
		// _ = "end of CoverTab[12618]"
//line /usr/local/go/src/net/addrselect.go:295
	}
//line /usr/local/go/src/net/addrselect.go:295
	// _ = "end of CoverTab[12614]"
//line /usr/local/go/src/net/addrselect.go:295
	_go_fuzz_dep_.CoverTab[12615]++
						for _, ent := range t {
//line /usr/local/go/src/net/addrselect.go:296
		_go_fuzz_dep_.CoverTab[12619]++
							if ent.Prefix.Contains(ip) {
//line /usr/local/go/src/net/addrselect.go:297
			_go_fuzz_dep_.CoverTab[12620]++
								return ent
//line /usr/local/go/src/net/addrselect.go:298
			// _ = "end of CoverTab[12620]"
		} else {
//line /usr/local/go/src/net/addrselect.go:299
			_go_fuzz_dep_.CoverTab[12621]++
//line /usr/local/go/src/net/addrselect.go:299
			// _ = "end of CoverTab[12621]"
//line /usr/local/go/src/net/addrselect.go:299
		}
//line /usr/local/go/src/net/addrselect.go:299
		// _ = "end of CoverTab[12619]"
	}
//line /usr/local/go/src/net/addrselect.go:300
	// _ = "end of CoverTab[12615]"
//line /usr/local/go/src/net/addrselect.go:300
	_go_fuzz_dep_.CoverTab[12616]++
						return policyTableEntry{}
//line /usr/local/go/src/net/addrselect.go:301
	// _ = "end of CoverTab[12616]"
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
//line /usr/local/go/src/net/addrselect.go:316
	_go_fuzz_dep_.CoverTab[12622]++
						if ip.IsLoopback() || func() bool {
//line /usr/local/go/src/net/addrselect.go:317
		_go_fuzz_dep_.CoverTab[12626]++
//line /usr/local/go/src/net/addrselect.go:317
		return ip.IsLinkLocalUnicast()
//line /usr/local/go/src/net/addrselect.go:317
		// _ = "end of CoverTab[12626]"
//line /usr/local/go/src/net/addrselect.go:317
	}() {
//line /usr/local/go/src/net/addrselect.go:317
		_go_fuzz_dep_.CoverTab[12627]++
							return scopeLinkLocal
//line /usr/local/go/src/net/addrselect.go:318
		// _ = "end of CoverTab[12627]"
	} else {
//line /usr/local/go/src/net/addrselect.go:319
		_go_fuzz_dep_.CoverTab[12628]++
//line /usr/local/go/src/net/addrselect.go:319
		// _ = "end of CoverTab[12628]"
//line /usr/local/go/src/net/addrselect.go:319
	}
//line /usr/local/go/src/net/addrselect.go:319
	// _ = "end of CoverTab[12622]"
//line /usr/local/go/src/net/addrselect.go:319
	_go_fuzz_dep_.CoverTab[12623]++
						ipv6 := ip.Is6() && func() bool {
//line /usr/local/go/src/net/addrselect.go:320
		_go_fuzz_dep_.CoverTab[12629]++
//line /usr/local/go/src/net/addrselect.go:320
		return !ip.Is4In6()
//line /usr/local/go/src/net/addrselect.go:320
		// _ = "end of CoverTab[12629]"
//line /usr/local/go/src/net/addrselect.go:320
	}()
						ipv6AsBytes := ip.As16()
						if ipv6 && func() bool {
//line /usr/local/go/src/net/addrselect.go:322
		_go_fuzz_dep_.CoverTab[12630]++
//line /usr/local/go/src/net/addrselect.go:322
		return ip.IsMulticast()
//line /usr/local/go/src/net/addrselect.go:322
		// _ = "end of CoverTab[12630]"
//line /usr/local/go/src/net/addrselect.go:322
	}() {
//line /usr/local/go/src/net/addrselect.go:322
		_go_fuzz_dep_.CoverTab[12631]++
							return scope(ipv6AsBytes[1] & 0xf)
//line /usr/local/go/src/net/addrselect.go:323
		// _ = "end of CoverTab[12631]"
	} else {
//line /usr/local/go/src/net/addrselect.go:324
		_go_fuzz_dep_.CoverTab[12632]++
//line /usr/local/go/src/net/addrselect.go:324
		// _ = "end of CoverTab[12632]"
//line /usr/local/go/src/net/addrselect.go:324
	}
//line /usr/local/go/src/net/addrselect.go:324
	// _ = "end of CoverTab[12623]"
//line /usr/local/go/src/net/addrselect.go:324
	_go_fuzz_dep_.CoverTab[12624]++

//line /usr/local/go/src/net/addrselect.go:327
	if ipv6 && func() bool {
//line /usr/local/go/src/net/addrselect.go:327
		_go_fuzz_dep_.CoverTab[12633]++
//line /usr/local/go/src/net/addrselect.go:327
		return ipv6AsBytes[0] == 0xfe
//line /usr/local/go/src/net/addrselect.go:327
		// _ = "end of CoverTab[12633]"
//line /usr/local/go/src/net/addrselect.go:327
	}() && func() bool {
//line /usr/local/go/src/net/addrselect.go:327
		_go_fuzz_dep_.CoverTab[12634]++
//line /usr/local/go/src/net/addrselect.go:327
		return ipv6AsBytes[1]&0xc0 == 0xc0
//line /usr/local/go/src/net/addrselect.go:327
		// _ = "end of CoverTab[12634]"
//line /usr/local/go/src/net/addrselect.go:327
	}() {
//line /usr/local/go/src/net/addrselect.go:327
		_go_fuzz_dep_.CoverTab[12635]++
							return scopeSiteLocal
//line /usr/local/go/src/net/addrselect.go:328
		// _ = "end of CoverTab[12635]"
	} else {
//line /usr/local/go/src/net/addrselect.go:329
		_go_fuzz_dep_.CoverTab[12636]++
//line /usr/local/go/src/net/addrselect.go:329
		// _ = "end of CoverTab[12636]"
//line /usr/local/go/src/net/addrselect.go:329
	}
//line /usr/local/go/src/net/addrselect.go:329
	// _ = "end of CoverTab[12624]"
//line /usr/local/go/src/net/addrselect.go:329
	_go_fuzz_dep_.CoverTab[12625]++
						return scopeGlobal
//line /usr/local/go/src/net/addrselect.go:330
	// _ = "end of CoverTab[12625]"
}

// commonPrefixLen reports the length of the longest prefix (looking
//line /usr/local/go/src/net/addrselect.go:333
// at the most significant, or leftmost, bits) that the
//line /usr/local/go/src/net/addrselect.go:333
// two addresses have in common, up to the length of a's prefix (i.e.,
//line /usr/local/go/src/net/addrselect.go:333
// the portion of the address not including the interface ID).
//line /usr/local/go/src/net/addrselect.go:333
//
//line /usr/local/go/src/net/addrselect.go:333
// If a or b is an IPv4 address as an IPv6 address, the IPv4 addresses
//line /usr/local/go/src/net/addrselect.go:333
// are compared (with max common prefix length of 32).
//line /usr/local/go/src/net/addrselect.go:333
// If a and b are different IP versions, 0 is returned.
//line /usr/local/go/src/net/addrselect.go:333
//
//line /usr/local/go/src/net/addrselect.go:333
// See https://tools.ietf.org/html/rfc6724#section-2.2
//line /usr/local/go/src/net/addrselect.go:343
func commonPrefixLen(a netip.Addr, b IP) (cpl int) {
//line /usr/local/go/src/net/addrselect.go:343
	_go_fuzz_dep_.CoverTab[12637]++
						if b4 := b.To4(); b4 != nil {
//line /usr/local/go/src/net/addrselect.go:344
		_go_fuzz_dep_.CoverTab[12642]++
							b = b4
//line /usr/local/go/src/net/addrselect.go:345
		// _ = "end of CoverTab[12642]"
	} else {
//line /usr/local/go/src/net/addrselect.go:346
		_go_fuzz_dep_.CoverTab[12643]++
//line /usr/local/go/src/net/addrselect.go:346
		// _ = "end of CoverTab[12643]"
//line /usr/local/go/src/net/addrselect.go:346
	}
//line /usr/local/go/src/net/addrselect.go:346
	// _ = "end of CoverTab[12637]"
//line /usr/local/go/src/net/addrselect.go:346
	_go_fuzz_dep_.CoverTab[12638]++
						aAsSlice := a.AsSlice()
						if len(aAsSlice) != len(b) {
//line /usr/local/go/src/net/addrselect.go:348
		_go_fuzz_dep_.CoverTab[12644]++
							return 0
//line /usr/local/go/src/net/addrselect.go:349
		// _ = "end of CoverTab[12644]"
	} else {
//line /usr/local/go/src/net/addrselect.go:350
		_go_fuzz_dep_.CoverTab[12645]++
//line /usr/local/go/src/net/addrselect.go:350
		// _ = "end of CoverTab[12645]"
//line /usr/local/go/src/net/addrselect.go:350
	}
//line /usr/local/go/src/net/addrselect.go:350
	// _ = "end of CoverTab[12638]"
//line /usr/local/go/src/net/addrselect.go:350
	_go_fuzz_dep_.CoverTab[12639]++

						if len(aAsSlice) > 8 {
//line /usr/local/go/src/net/addrselect.go:352
		_go_fuzz_dep_.CoverTab[12646]++
							aAsSlice = aAsSlice[:8]
							b = b[:8]
//line /usr/local/go/src/net/addrselect.go:354
		// _ = "end of CoverTab[12646]"
	} else {
//line /usr/local/go/src/net/addrselect.go:355
		_go_fuzz_dep_.CoverTab[12647]++
//line /usr/local/go/src/net/addrselect.go:355
		// _ = "end of CoverTab[12647]"
//line /usr/local/go/src/net/addrselect.go:355
	}
//line /usr/local/go/src/net/addrselect.go:355
	// _ = "end of CoverTab[12639]"
//line /usr/local/go/src/net/addrselect.go:355
	_go_fuzz_dep_.CoverTab[12640]++
						for len(aAsSlice) > 0 {
//line /usr/local/go/src/net/addrselect.go:356
		_go_fuzz_dep_.CoverTab[12648]++
							if aAsSlice[0] == b[0] {
//line /usr/local/go/src/net/addrselect.go:357
			_go_fuzz_dep_.CoverTab[12650]++
								cpl += 8
								aAsSlice = aAsSlice[1:]
								b = b[1:]
								continue
//line /usr/local/go/src/net/addrselect.go:361
			// _ = "end of CoverTab[12650]"
		} else {
//line /usr/local/go/src/net/addrselect.go:362
			_go_fuzz_dep_.CoverTab[12651]++
//line /usr/local/go/src/net/addrselect.go:362
			// _ = "end of CoverTab[12651]"
//line /usr/local/go/src/net/addrselect.go:362
		}
//line /usr/local/go/src/net/addrselect.go:362
		// _ = "end of CoverTab[12648]"
//line /usr/local/go/src/net/addrselect.go:362
		_go_fuzz_dep_.CoverTab[12649]++
							bits := 8
							ab, bb := aAsSlice[0], b[0]
							for {
//line /usr/local/go/src/net/addrselect.go:365
			_go_fuzz_dep_.CoverTab[12652]++
								ab >>= 1
								bb >>= 1
								bits--
								if ab == bb {
//line /usr/local/go/src/net/addrselect.go:369
				_go_fuzz_dep_.CoverTab[12653]++
									cpl += bits
									return
//line /usr/local/go/src/net/addrselect.go:371
				// _ = "end of CoverTab[12653]"
			} else {
//line /usr/local/go/src/net/addrselect.go:372
				_go_fuzz_dep_.CoverTab[12654]++
//line /usr/local/go/src/net/addrselect.go:372
				// _ = "end of CoverTab[12654]"
//line /usr/local/go/src/net/addrselect.go:372
			}
//line /usr/local/go/src/net/addrselect.go:372
			// _ = "end of CoverTab[12652]"
		}
//line /usr/local/go/src/net/addrselect.go:373
		// _ = "end of CoverTab[12649]"
	}
//line /usr/local/go/src/net/addrselect.go:374
	// _ = "end of CoverTab[12640]"
//line /usr/local/go/src/net/addrselect.go:374
	_go_fuzz_dep_.CoverTab[12641]++
						return
//line /usr/local/go/src/net/addrselect.go:375
	// _ = "end of CoverTab[12641]"
}

//line /usr/local/go/src/net/addrselect.go:376
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/addrselect.go:376
var _ = _go_fuzz_dep_.CoverTab
