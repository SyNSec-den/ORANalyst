// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/ipsock.go:5
package net

//line /snap/go/10455/src/net/ipsock.go:5
import (
//line /snap/go/10455/src/net/ipsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/ipsock.go:5
)
//line /snap/go/10455/src/net/ipsock.go:5
import (
//line /snap/go/10455/src/net/ipsock.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/ipsock.go:5
)

import (
	"context"
	"internal/bytealg"
	"runtime"
	"sync"
)

//line /snap/go/10455/src/net/ipsock.go:21
type ipStackCapabilities struct {
	sync.Once		// guards following
	ipv4Enabled		bool
	ipv6Enabled		bool
	ipv4MappedIPv6Enabled	bool
}

var ipStackCaps ipStackCapabilities

// supportsIPv4 reports whether the platform supports IPv4 networking
//line /snap/go/10455/src/net/ipsock.go:30
// functionality.
//line /snap/go/10455/src/net/ipsock.go:32
func supportsIPv4() bool {
//line /snap/go/10455/src/net/ipsock.go:32
	_go_fuzz_dep_.CoverTab[6717]++
						ipStackCaps.Once.Do(ipStackCaps.probe)
						return ipStackCaps.ipv4Enabled
//line /snap/go/10455/src/net/ipsock.go:34
	// _ = "end of CoverTab[6717]"
}

// supportsIPv6 reports whether the platform supports IPv6 networking
//line /snap/go/10455/src/net/ipsock.go:37
// functionality.
//line /snap/go/10455/src/net/ipsock.go:39
func supportsIPv6() bool {
//line /snap/go/10455/src/net/ipsock.go:39
	_go_fuzz_dep_.CoverTab[6718]++
						ipStackCaps.Once.Do(ipStackCaps.probe)
						return ipStackCaps.ipv6Enabled
//line /snap/go/10455/src/net/ipsock.go:41
	// _ = "end of CoverTab[6718]"
}

// supportsIPv4map reports whether the platform supports mapping an
//line /snap/go/10455/src/net/ipsock.go:44
// IPv4 address inside an IPv6 address at transport layer
//line /snap/go/10455/src/net/ipsock.go:44
// protocols. See RFC 4291, RFC 4038 and RFC 3493.
//line /snap/go/10455/src/net/ipsock.go:47
func supportsIPv4map() bool {
//line /snap/go/10455/src/net/ipsock.go:47
	_go_fuzz_dep_.CoverTab[6719]++

//line /snap/go/10455/src/net/ipsock.go:50
	switch runtime.GOOS {
	case "dragonfly", "openbsd":
//line /snap/go/10455/src/net/ipsock.go:51
		_go_fuzz_dep_.CoverTab[528797]++
//line /snap/go/10455/src/net/ipsock.go:51
		_go_fuzz_dep_.CoverTab[6721]++
							return false
//line /snap/go/10455/src/net/ipsock.go:52
		// _ = "end of CoverTab[6721]"
//line /snap/go/10455/src/net/ipsock.go:52
	default:
//line /snap/go/10455/src/net/ipsock.go:52
		_go_fuzz_dep_.CoverTab[528798]++
//line /snap/go/10455/src/net/ipsock.go:52
		_go_fuzz_dep_.CoverTab[6722]++
//line /snap/go/10455/src/net/ipsock.go:52
		// _ = "end of CoverTab[6722]"
	}
//line /snap/go/10455/src/net/ipsock.go:53
	// _ = "end of CoverTab[6719]"
//line /snap/go/10455/src/net/ipsock.go:53
	_go_fuzz_dep_.CoverTab[6720]++

						ipStackCaps.Once.Do(ipStackCaps.probe)
						return ipStackCaps.ipv4MappedIPv6Enabled
//line /snap/go/10455/src/net/ipsock.go:56
	// _ = "end of CoverTab[6720]"
}

// An addrList represents a list of network endpoint addresses.
type addrList []Addr

// isIPv4 reports whether addr contains an IPv4 address.
func isIPv4(addr Addr) bool {
//line /snap/go/10455/src/net/ipsock.go:63
	_go_fuzz_dep_.CoverTab[6723]++
						switch addr := addr.(type) {
	case *TCPAddr:
//line /snap/go/10455/src/net/ipsock.go:65
		_go_fuzz_dep_.CoverTab[528799]++
//line /snap/go/10455/src/net/ipsock.go:65
		_go_fuzz_dep_.CoverTab[6725]++
							return addr.IP.To4() != nil
//line /snap/go/10455/src/net/ipsock.go:66
		// _ = "end of CoverTab[6725]"
	case *UDPAddr:
//line /snap/go/10455/src/net/ipsock.go:67
		_go_fuzz_dep_.CoverTab[528800]++
//line /snap/go/10455/src/net/ipsock.go:67
		_go_fuzz_dep_.CoverTab[6726]++
							return addr.IP.To4() != nil
//line /snap/go/10455/src/net/ipsock.go:68
		// _ = "end of CoverTab[6726]"
	case *IPAddr:
//line /snap/go/10455/src/net/ipsock.go:69
		_go_fuzz_dep_.CoverTab[528801]++
//line /snap/go/10455/src/net/ipsock.go:69
		_go_fuzz_dep_.CoverTab[6727]++
							return addr.IP.To4() != nil
//line /snap/go/10455/src/net/ipsock.go:70
		// _ = "end of CoverTab[6727]"
	}
//line /snap/go/10455/src/net/ipsock.go:71
	// _ = "end of CoverTab[6723]"
//line /snap/go/10455/src/net/ipsock.go:71
	_go_fuzz_dep_.CoverTab[6724]++
						return false
//line /snap/go/10455/src/net/ipsock.go:72
	// _ = "end of CoverTab[6724]"
}

// isNotIPv4 reports whether addr does not contain an IPv4 address.
func isNotIPv4(addr Addr) bool {
//line /snap/go/10455/src/net/ipsock.go:76
	_go_fuzz_dep_.CoverTab[6728]++
//line /snap/go/10455/src/net/ipsock.go:76
	return !isIPv4(addr)
//line /snap/go/10455/src/net/ipsock.go:76
	// _ = "end of CoverTab[6728]"
//line /snap/go/10455/src/net/ipsock.go:76
}

// forResolve returns the most appropriate address in address for
//line /snap/go/10455/src/net/ipsock.go:78
// a call to ResolveTCPAddr, ResolveUDPAddr, or ResolveIPAddr.
//line /snap/go/10455/src/net/ipsock.go:78
// IPv4 is preferred, unless addr contains an IPv6 literal.
//line /snap/go/10455/src/net/ipsock.go:81
func (addrs addrList) forResolve(network, addr string) Addr {
//line /snap/go/10455/src/net/ipsock.go:81
	_go_fuzz_dep_.CoverTab[6729]++
						var want6 bool
						switch network {
	case "ip":
//line /snap/go/10455/src/net/ipsock.go:84
		_go_fuzz_dep_.CoverTab[528802]++
//line /snap/go/10455/src/net/ipsock.go:84
		_go_fuzz_dep_.CoverTab[6732]++

							want6 = count(addr, ':') > 0
//line /snap/go/10455/src/net/ipsock.go:86
		// _ = "end of CoverTab[6732]"
	case "tcp", "udp":
//line /snap/go/10455/src/net/ipsock.go:87
		_go_fuzz_dep_.CoverTab[528803]++
//line /snap/go/10455/src/net/ipsock.go:87
		_go_fuzz_dep_.CoverTab[6733]++

							want6 = count(addr, '[') > 0
//line /snap/go/10455/src/net/ipsock.go:89
		// _ = "end of CoverTab[6733]"
//line /snap/go/10455/src/net/ipsock.go:89
	default:
//line /snap/go/10455/src/net/ipsock.go:89
		_go_fuzz_dep_.CoverTab[528804]++
//line /snap/go/10455/src/net/ipsock.go:89
		_go_fuzz_dep_.CoverTab[6734]++
//line /snap/go/10455/src/net/ipsock.go:89
		// _ = "end of CoverTab[6734]"
	}
//line /snap/go/10455/src/net/ipsock.go:90
	// _ = "end of CoverTab[6729]"
//line /snap/go/10455/src/net/ipsock.go:90
	_go_fuzz_dep_.CoverTab[6730]++
						if want6 {
//line /snap/go/10455/src/net/ipsock.go:91
		_go_fuzz_dep_.CoverTab[528805]++
//line /snap/go/10455/src/net/ipsock.go:91
		_go_fuzz_dep_.CoverTab[6735]++
							return addrs.first(isNotIPv4)
//line /snap/go/10455/src/net/ipsock.go:92
		// _ = "end of CoverTab[6735]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:93
		_go_fuzz_dep_.CoverTab[528806]++
//line /snap/go/10455/src/net/ipsock.go:93
		_go_fuzz_dep_.CoverTab[6736]++
//line /snap/go/10455/src/net/ipsock.go:93
		// _ = "end of CoverTab[6736]"
//line /snap/go/10455/src/net/ipsock.go:93
	}
//line /snap/go/10455/src/net/ipsock.go:93
	// _ = "end of CoverTab[6730]"
//line /snap/go/10455/src/net/ipsock.go:93
	_go_fuzz_dep_.CoverTab[6731]++
						return addrs.first(isIPv4)
//line /snap/go/10455/src/net/ipsock.go:94
	// _ = "end of CoverTab[6731]"
}

// first returns the first address which satisfies strategy, or if
//line /snap/go/10455/src/net/ipsock.go:97
// none do, then the first address of any kind.
//line /snap/go/10455/src/net/ipsock.go:99
func (addrs addrList) first(strategy func(Addr) bool) Addr {
//line /snap/go/10455/src/net/ipsock.go:99
	_go_fuzz_dep_.CoverTab[6737]++
//line /snap/go/10455/src/net/ipsock.go:99
	_go_fuzz_dep_.CoverTab[786697] = 0
						for _, addr := range addrs {
//line /snap/go/10455/src/net/ipsock.go:100
		if _go_fuzz_dep_.CoverTab[786697] == 0 {
//line /snap/go/10455/src/net/ipsock.go:100
			_go_fuzz_dep_.CoverTab[528863]++
//line /snap/go/10455/src/net/ipsock.go:100
		} else {
//line /snap/go/10455/src/net/ipsock.go:100
			_go_fuzz_dep_.CoverTab[528864]++
//line /snap/go/10455/src/net/ipsock.go:100
		}
//line /snap/go/10455/src/net/ipsock.go:100
		_go_fuzz_dep_.CoverTab[786697] = 1
//line /snap/go/10455/src/net/ipsock.go:100
		_go_fuzz_dep_.CoverTab[6739]++
							if strategy(addr) {
//line /snap/go/10455/src/net/ipsock.go:101
			_go_fuzz_dep_.CoverTab[528807]++
//line /snap/go/10455/src/net/ipsock.go:101
			_go_fuzz_dep_.CoverTab[6740]++
								return addr
//line /snap/go/10455/src/net/ipsock.go:102
			// _ = "end of CoverTab[6740]"
		} else {
//line /snap/go/10455/src/net/ipsock.go:103
			_go_fuzz_dep_.CoverTab[528808]++
//line /snap/go/10455/src/net/ipsock.go:103
			_go_fuzz_dep_.CoverTab[6741]++
//line /snap/go/10455/src/net/ipsock.go:103
			// _ = "end of CoverTab[6741]"
//line /snap/go/10455/src/net/ipsock.go:103
		}
//line /snap/go/10455/src/net/ipsock.go:103
		// _ = "end of CoverTab[6739]"
	}
//line /snap/go/10455/src/net/ipsock.go:104
	if _go_fuzz_dep_.CoverTab[786697] == 0 {
//line /snap/go/10455/src/net/ipsock.go:104
		_go_fuzz_dep_.CoverTab[528865]++
//line /snap/go/10455/src/net/ipsock.go:104
	} else {
//line /snap/go/10455/src/net/ipsock.go:104
		_go_fuzz_dep_.CoverTab[528866]++
//line /snap/go/10455/src/net/ipsock.go:104
	}
//line /snap/go/10455/src/net/ipsock.go:104
	// _ = "end of CoverTab[6737]"
//line /snap/go/10455/src/net/ipsock.go:104
	_go_fuzz_dep_.CoverTab[6738]++
						return addrs[0]
//line /snap/go/10455/src/net/ipsock.go:105
	// _ = "end of CoverTab[6738]"
}

// partition divides an address list into two categories, using a
//line /snap/go/10455/src/net/ipsock.go:108
// strategy function to assign a boolean label to each address.
//line /snap/go/10455/src/net/ipsock.go:108
// The first address, and any with a matching label, are returned as
//line /snap/go/10455/src/net/ipsock.go:108
// primaries, while addresses with the opposite label are returned
//line /snap/go/10455/src/net/ipsock.go:108
// as fallbacks. For non-empty inputs, primaries is guaranteed to be
//line /snap/go/10455/src/net/ipsock.go:108
// non-empty.
//line /snap/go/10455/src/net/ipsock.go:114
func (addrs addrList) partition(strategy func(Addr) bool) (primaries, fallbacks addrList) {
//line /snap/go/10455/src/net/ipsock.go:114
	_go_fuzz_dep_.CoverTab[6742]++
						var primaryLabel bool
//line /snap/go/10455/src/net/ipsock.go:115
	_go_fuzz_dep_.CoverTab[786698] = 0
						for i, addr := range addrs {
//line /snap/go/10455/src/net/ipsock.go:116
		if _go_fuzz_dep_.CoverTab[786698] == 0 {
//line /snap/go/10455/src/net/ipsock.go:116
			_go_fuzz_dep_.CoverTab[528867]++
//line /snap/go/10455/src/net/ipsock.go:116
		} else {
//line /snap/go/10455/src/net/ipsock.go:116
			_go_fuzz_dep_.CoverTab[528868]++
//line /snap/go/10455/src/net/ipsock.go:116
		}
//line /snap/go/10455/src/net/ipsock.go:116
		_go_fuzz_dep_.CoverTab[786698] = 1
//line /snap/go/10455/src/net/ipsock.go:116
		_go_fuzz_dep_.CoverTab[6744]++
							label := strategy(addr)
							if i == 0 || func() bool {
//line /snap/go/10455/src/net/ipsock.go:118
			_go_fuzz_dep_.CoverTab[6745]++
//line /snap/go/10455/src/net/ipsock.go:118
			return label == primaryLabel
//line /snap/go/10455/src/net/ipsock.go:118
			// _ = "end of CoverTab[6745]"
//line /snap/go/10455/src/net/ipsock.go:118
		}() {
//line /snap/go/10455/src/net/ipsock.go:118
			_go_fuzz_dep_.CoverTab[528809]++
//line /snap/go/10455/src/net/ipsock.go:118
			_go_fuzz_dep_.CoverTab[6746]++
								primaryLabel = label
								primaries = append(primaries, addr)
//line /snap/go/10455/src/net/ipsock.go:120
			// _ = "end of CoverTab[6746]"
		} else {
//line /snap/go/10455/src/net/ipsock.go:121
			_go_fuzz_dep_.CoverTab[528810]++
//line /snap/go/10455/src/net/ipsock.go:121
			_go_fuzz_dep_.CoverTab[6747]++
								fallbacks = append(fallbacks, addr)
//line /snap/go/10455/src/net/ipsock.go:122
			// _ = "end of CoverTab[6747]"
		}
//line /snap/go/10455/src/net/ipsock.go:123
		// _ = "end of CoverTab[6744]"
	}
//line /snap/go/10455/src/net/ipsock.go:124
	if _go_fuzz_dep_.CoverTab[786698] == 0 {
//line /snap/go/10455/src/net/ipsock.go:124
		_go_fuzz_dep_.CoverTab[528869]++
//line /snap/go/10455/src/net/ipsock.go:124
	} else {
//line /snap/go/10455/src/net/ipsock.go:124
		_go_fuzz_dep_.CoverTab[528870]++
//line /snap/go/10455/src/net/ipsock.go:124
	}
//line /snap/go/10455/src/net/ipsock.go:124
	// _ = "end of CoverTab[6742]"
//line /snap/go/10455/src/net/ipsock.go:124
	_go_fuzz_dep_.CoverTab[6743]++
						return
//line /snap/go/10455/src/net/ipsock.go:125
	// _ = "end of CoverTab[6743]"
}

// filterAddrList applies a filter to a list of IP addresses,
//line /snap/go/10455/src/net/ipsock.go:128
// yielding a list of Addr objects. Known filters are nil, ipv4only,
//line /snap/go/10455/src/net/ipsock.go:128
// and ipv6only. It returns every address when the filter is nil.
//line /snap/go/10455/src/net/ipsock.go:128
// The result contains at least one address when error is nil.
//line /snap/go/10455/src/net/ipsock.go:132
func filterAddrList(filter func(IPAddr) bool, ips []IPAddr, inetaddr func(IPAddr) Addr, originalAddr string) (addrList, error) {
//line /snap/go/10455/src/net/ipsock.go:132
	_go_fuzz_dep_.CoverTab[6748]++
						var addrs addrList
//line /snap/go/10455/src/net/ipsock.go:133
	_go_fuzz_dep_.CoverTab[786699] = 0
						for _, ip := range ips {
//line /snap/go/10455/src/net/ipsock.go:134
		if _go_fuzz_dep_.CoverTab[786699] == 0 {
//line /snap/go/10455/src/net/ipsock.go:134
			_go_fuzz_dep_.CoverTab[528871]++
//line /snap/go/10455/src/net/ipsock.go:134
		} else {
//line /snap/go/10455/src/net/ipsock.go:134
			_go_fuzz_dep_.CoverTab[528872]++
//line /snap/go/10455/src/net/ipsock.go:134
		}
//line /snap/go/10455/src/net/ipsock.go:134
		_go_fuzz_dep_.CoverTab[786699] = 1
//line /snap/go/10455/src/net/ipsock.go:134
		_go_fuzz_dep_.CoverTab[6751]++
							if filter == nil || func() bool {
//line /snap/go/10455/src/net/ipsock.go:135
			_go_fuzz_dep_.CoverTab[6752]++
//line /snap/go/10455/src/net/ipsock.go:135
			return filter(ip)
//line /snap/go/10455/src/net/ipsock.go:135
			// _ = "end of CoverTab[6752]"
//line /snap/go/10455/src/net/ipsock.go:135
		}() {
//line /snap/go/10455/src/net/ipsock.go:135
			_go_fuzz_dep_.CoverTab[528811]++
//line /snap/go/10455/src/net/ipsock.go:135
			_go_fuzz_dep_.CoverTab[6753]++
								addrs = append(addrs, inetaddr(ip))
//line /snap/go/10455/src/net/ipsock.go:136
			// _ = "end of CoverTab[6753]"
		} else {
//line /snap/go/10455/src/net/ipsock.go:137
			_go_fuzz_dep_.CoverTab[528812]++
//line /snap/go/10455/src/net/ipsock.go:137
			_go_fuzz_dep_.CoverTab[6754]++
//line /snap/go/10455/src/net/ipsock.go:137
			// _ = "end of CoverTab[6754]"
//line /snap/go/10455/src/net/ipsock.go:137
		}
//line /snap/go/10455/src/net/ipsock.go:137
		// _ = "end of CoverTab[6751]"
	}
//line /snap/go/10455/src/net/ipsock.go:138
	if _go_fuzz_dep_.CoverTab[786699] == 0 {
//line /snap/go/10455/src/net/ipsock.go:138
		_go_fuzz_dep_.CoverTab[528873]++
//line /snap/go/10455/src/net/ipsock.go:138
	} else {
//line /snap/go/10455/src/net/ipsock.go:138
		_go_fuzz_dep_.CoverTab[528874]++
//line /snap/go/10455/src/net/ipsock.go:138
	}
//line /snap/go/10455/src/net/ipsock.go:138
	// _ = "end of CoverTab[6748]"
//line /snap/go/10455/src/net/ipsock.go:138
	_go_fuzz_dep_.CoverTab[6749]++
						if len(addrs) == 0 {
//line /snap/go/10455/src/net/ipsock.go:139
		_go_fuzz_dep_.CoverTab[528813]++
//line /snap/go/10455/src/net/ipsock.go:139
		_go_fuzz_dep_.CoverTab[6755]++
							return nil, &AddrError{Err: errNoSuitableAddress.Error(), Addr: originalAddr}
//line /snap/go/10455/src/net/ipsock.go:140
		// _ = "end of CoverTab[6755]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:141
		_go_fuzz_dep_.CoverTab[528814]++
//line /snap/go/10455/src/net/ipsock.go:141
		_go_fuzz_dep_.CoverTab[6756]++
//line /snap/go/10455/src/net/ipsock.go:141
		// _ = "end of CoverTab[6756]"
//line /snap/go/10455/src/net/ipsock.go:141
	}
//line /snap/go/10455/src/net/ipsock.go:141
	// _ = "end of CoverTab[6749]"
//line /snap/go/10455/src/net/ipsock.go:141
	_go_fuzz_dep_.CoverTab[6750]++
						return addrs, nil
//line /snap/go/10455/src/net/ipsock.go:142
	// _ = "end of CoverTab[6750]"
}

// ipv4only reports whether addr is an IPv4 address.
func ipv4only(addr IPAddr) bool {
//line /snap/go/10455/src/net/ipsock.go:146
	_go_fuzz_dep_.CoverTab[6757]++
						return addr.IP.To4() != nil
//line /snap/go/10455/src/net/ipsock.go:147
	// _ = "end of CoverTab[6757]"
}

// ipv6only reports whether addr is an IPv6 address except IPv4-mapped IPv6 address.
func ipv6only(addr IPAddr) bool {
//line /snap/go/10455/src/net/ipsock.go:151
	_go_fuzz_dep_.CoverTab[6758]++
						return len(addr.IP) == IPv6len && func() bool {
//line /snap/go/10455/src/net/ipsock.go:152
		_go_fuzz_dep_.CoverTab[6759]++
//line /snap/go/10455/src/net/ipsock.go:152
		return addr.IP.To4() == nil
//line /snap/go/10455/src/net/ipsock.go:152
		// _ = "end of CoverTab[6759]"
//line /snap/go/10455/src/net/ipsock.go:152
	}()
//line /snap/go/10455/src/net/ipsock.go:152
	// _ = "end of CoverTab[6758]"
}

// SplitHostPort splits a network address of the form "host:port",
//line /snap/go/10455/src/net/ipsock.go:155
// "host%zone:port", "[host]:port" or "[host%zone]:port" into host or
//line /snap/go/10455/src/net/ipsock.go:155
// host%zone and port.
//line /snap/go/10455/src/net/ipsock.go:155
//
//line /snap/go/10455/src/net/ipsock.go:155
// A literal IPv6 address in hostport must be enclosed in square
//line /snap/go/10455/src/net/ipsock.go:155
// brackets, as in "[::1]:80", "[::1%lo0]:80".
//line /snap/go/10455/src/net/ipsock.go:155
//
//line /snap/go/10455/src/net/ipsock.go:155
// See func Dial for a description of the hostport parameter, and host
//line /snap/go/10455/src/net/ipsock.go:155
// and port results.
//line /snap/go/10455/src/net/ipsock.go:164
func SplitHostPort(hostport string) (host, port string, err error) {
//line /snap/go/10455/src/net/ipsock.go:164
	_go_fuzz_dep_.CoverTab[6760]++
						const (
		missingPort	= "missing port in address"
		tooManyColons	= "too many colons in address"
	)
	addrErr := func(addr, why string) (host, port string, err error) {
//line /snap/go/10455/src/net/ipsock.go:169
		_go_fuzz_dep_.CoverTab[6766]++
							return "", "", &AddrError{Err: why, Addr: addr}
//line /snap/go/10455/src/net/ipsock.go:170
		// _ = "end of CoverTab[6766]"
	}
//line /snap/go/10455/src/net/ipsock.go:171
	// _ = "end of CoverTab[6760]"
//line /snap/go/10455/src/net/ipsock.go:171
	_go_fuzz_dep_.CoverTab[6761]++
						j, k := 0, 0

//line /snap/go/10455/src/net/ipsock.go:175
	i := last(hostport, ':')
	if i < 0 {
//line /snap/go/10455/src/net/ipsock.go:176
		_go_fuzz_dep_.CoverTab[528815]++
//line /snap/go/10455/src/net/ipsock.go:176
		_go_fuzz_dep_.CoverTab[6767]++
							return addrErr(hostport, missingPort)
//line /snap/go/10455/src/net/ipsock.go:177
		// _ = "end of CoverTab[6767]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:178
		_go_fuzz_dep_.CoverTab[528816]++
//line /snap/go/10455/src/net/ipsock.go:178
		_go_fuzz_dep_.CoverTab[6768]++
//line /snap/go/10455/src/net/ipsock.go:178
		// _ = "end of CoverTab[6768]"
//line /snap/go/10455/src/net/ipsock.go:178
	}
//line /snap/go/10455/src/net/ipsock.go:178
	// _ = "end of CoverTab[6761]"
//line /snap/go/10455/src/net/ipsock.go:178
	_go_fuzz_dep_.CoverTab[6762]++

						if hostport[0] == '[' {
//line /snap/go/10455/src/net/ipsock.go:180
		_go_fuzz_dep_.CoverTab[528817]++
//line /snap/go/10455/src/net/ipsock.go:180
		_go_fuzz_dep_.CoverTab[6769]++

							end := bytealg.IndexByteString(hostport, ']')
							if end < 0 {
//line /snap/go/10455/src/net/ipsock.go:183
			_go_fuzz_dep_.CoverTab[528819]++
//line /snap/go/10455/src/net/ipsock.go:183
			_go_fuzz_dep_.CoverTab[6772]++
								return addrErr(hostport, "missing ']' in address")
//line /snap/go/10455/src/net/ipsock.go:184
			// _ = "end of CoverTab[6772]"
		} else {
//line /snap/go/10455/src/net/ipsock.go:185
			_go_fuzz_dep_.CoverTab[528820]++
//line /snap/go/10455/src/net/ipsock.go:185
			_go_fuzz_dep_.CoverTab[6773]++
//line /snap/go/10455/src/net/ipsock.go:185
			// _ = "end of CoverTab[6773]"
//line /snap/go/10455/src/net/ipsock.go:185
		}
//line /snap/go/10455/src/net/ipsock.go:185
		// _ = "end of CoverTab[6769]"
//line /snap/go/10455/src/net/ipsock.go:185
		_go_fuzz_dep_.CoverTab[6770]++
							switch end + 1 {
		case len(hostport):
//line /snap/go/10455/src/net/ipsock.go:187
			_go_fuzz_dep_.CoverTab[528821]++
//line /snap/go/10455/src/net/ipsock.go:187
			_go_fuzz_dep_.CoverTab[6774]++

								return addrErr(hostport, missingPort)
//line /snap/go/10455/src/net/ipsock.go:189
			// _ = "end of CoverTab[6774]"
		case i:
//line /snap/go/10455/src/net/ipsock.go:190
			_go_fuzz_dep_.CoverTab[528822]++
//line /snap/go/10455/src/net/ipsock.go:190
			_go_fuzz_dep_.CoverTab[6775]++
//line /snap/go/10455/src/net/ipsock.go:190
			// _ = "end of CoverTab[6775]"

		default:
//line /snap/go/10455/src/net/ipsock.go:192
			_go_fuzz_dep_.CoverTab[528823]++
//line /snap/go/10455/src/net/ipsock.go:192
			_go_fuzz_dep_.CoverTab[6776]++

//line /snap/go/10455/src/net/ipsock.go:195
			if hostport[end+1] == ':' {
//line /snap/go/10455/src/net/ipsock.go:195
				_go_fuzz_dep_.CoverTab[528824]++
//line /snap/go/10455/src/net/ipsock.go:195
				_go_fuzz_dep_.CoverTab[6778]++
									return addrErr(hostport, tooManyColons)
//line /snap/go/10455/src/net/ipsock.go:196
				// _ = "end of CoverTab[6778]"
			} else {
//line /snap/go/10455/src/net/ipsock.go:197
				_go_fuzz_dep_.CoverTab[528825]++
//line /snap/go/10455/src/net/ipsock.go:197
				_go_fuzz_dep_.CoverTab[6779]++
//line /snap/go/10455/src/net/ipsock.go:197
				// _ = "end of CoverTab[6779]"
//line /snap/go/10455/src/net/ipsock.go:197
			}
//line /snap/go/10455/src/net/ipsock.go:197
			// _ = "end of CoverTab[6776]"
//line /snap/go/10455/src/net/ipsock.go:197
			_go_fuzz_dep_.CoverTab[6777]++
								return addrErr(hostport, missingPort)
//line /snap/go/10455/src/net/ipsock.go:198
			// _ = "end of CoverTab[6777]"
		}
//line /snap/go/10455/src/net/ipsock.go:199
		// _ = "end of CoverTab[6770]"
//line /snap/go/10455/src/net/ipsock.go:199
		_go_fuzz_dep_.CoverTab[6771]++
							host = hostport[1:end]
							j, k = 1, end+1
//line /snap/go/10455/src/net/ipsock.go:201
		// _ = "end of CoverTab[6771]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:202
		_go_fuzz_dep_.CoverTab[528818]++
//line /snap/go/10455/src/net/ipsock.go:202
		_go_fuzz_dep_.CoverTab[6780]++
							host = hostport[:i]
							if bytealg.IndexByteString(host, ':') >= 0 {
//line /snap/go/10455/src/net/ipsock.go:204
			_go_fuzz_dep_.CoverTab[528826]++
//line /snap/go/10455/src/net/ipsock.go:204
			_go_fuzz_dep_.CoverTab[6781]++
								return addrErr(hostport, tooManyColons)
//line /snap/go/10455/src/net/ipsock.go:205
			// _ = "end of CoverTab[6781]"
		} else {
//line /snap/go/10455/src/net/ipsock.go:206
			_go_fuzz_dep_.CoverTab[528827]++
//line /snap/go/10455/src/net/ipsock.go:206
			_go_fuzz_dep_.CoverTab[6782]++
//line /snap/go/10455/src/net/ipsock.go:206
			// _ = "end of CoverTab[6782]"
//line /snap/go/10455/src/net/ipsock.go:206
		}
//line /snap/go/10455/src/net/ipsock.go:206
		// _ = "end of CoverTab[6780]"
	}
//line /snap/go/10455/src/net/ipsock.go:207
	// _ = "end of CoverTab[6762]"
//line /snap/go/10455/src/net/ipsock.go:207
	_go_fuzz_dep_.CoverTab[6763]++
						if bytealg.IndexByteString(hostport[j:], '[') >= 0 {
//line /snap/go/10455/src/net/ipsock.go:208
		_go_fuzz_dep_.CoverTab[528828]++
//line /snap/go/10455/src/net/ipsock.go:208
		_go_fuzz_dep_.CoverTab[6783]++
							return addrErr(hostport, "unexpected '[' in address")
//line /snap/go/10455/src/net/ipsock.go:209
		// _ = "end of CoverTab[6783]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:210
		_go_fuzz_dep_.CoverTab[528829]++
//line /snap/go/10455/src/net/ipsock.go:210
		_go_fuzz_dep_.CoverTab[6784]++
//line /snap/go/10455/src/net/ipsock.go:210
		// _ = "end of CoverTab[6784]"
//line /snap/go/10455/src/net/ipsock.go:210
	}
//line /snap/go/10455/src/net/ipsock.go:210
	// _ = "end of CoverTab[6763]"
//line /snap/go/10455/src/net/ipsock.go:210
	_go_fuzz_dep_.CoverTab[6764]++
						if bytealg.IndexByteString(hostport[k:], ']') >= 0 {
//line /snap/go/10455/src/net/ipsock.go:211
		_go_fuzz_dep_.CoverTab[528830]++
//line /snap/go/10455/src/net/ipsock.go:211
		_go_fuzz_dep_.CoverTab[6785]++
							return addrErr(hostport, "unexpected ']' in address")
//line /snap/go/10455/src/net/ipsock.go:212
		// _ = "end of CoverTab[6785]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:213
		_go_fuzz_dep_.CoverTab[528831]++
//line /snap/go/10455/src/net/ipsock.go:213
		_go_fuzz_dep_.CoverTab[6786]++
//line /snap/go/10455/src/net/ipsock.go:213
		// _ = "end of CoverTab[6786]"
//line /snap/go/10455/src/net/ipsock.go:213
	}
//line /snap/go/10455/src/net/ipsock.go:213
	// _ = "end of CoverTab[6764]"
//line /snap/go/10455/src/net/ipsock.go:213
	_go_fuzz_dep_.CoverTab[6765]++

						port = hostport[i+1:]
						return host, port, nil
//line /snap/go/10455/src/net/ipsock.go:216
	// _ = "end of CoverTab[6765]"
}

func splitHostZone(s string) (host, zone string) {
//line /snap/go/10455/src/net/ipsock.go:219
	_go_fuzz_dep_.CoverTab[6787]++

//line /snap/go/10455/src/net/ipsock.go:222
	if i := last(s, '%'); i > 0 {
//line /snap/go/10455/src/net/ipsock.go:222
		_go_fuzz_dep_.CoverTab[528832]++
//line /snap/go/10455/src/net/ipsock.go:222
		_go_fuzz_dep_.CoverTab[6789]++
							host, zone = s[:i], s[i+1:]
//line /snap/go/10455/src/net/ipsock.go:223
		// _ = "end of CoverTab[6789]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:224
		_go_fuzz_dep_.CoverTab[528833]++
//line /snap/go/10455/src/net/ipsock.go:224
		_go_fuzz_dep_.CoverTab[6790]++
							host = s
//line /snap/go/10455/src/net/ipsock.go:225
		// _ = "end of CoverTab[6790]"
	}
//line /snap/go/10455/src/net/ipsock.go:226
	// _ = "end of CoverTab[6787]"
//line /snap/go/10455/src/net/ipsock.go:226
	_go_fuzz_dep_.CoverTab[6788]++
						return
//line /snap/go/10455/src/net/ipsock.go:227
	// _ = "end of CoverTab[6788]"
}

// JoinHostPort combines host and port into a network address of the
//line /snap/go/10455/src/net/ipsock.go:230
// form "host:port". If host contains a colon, as found in literal
//line /snap/go/10455/src/net/ipsock.go:230
// IPv6 addresses, then JoinHostPort returns "[host]:port".
//line /snap/go/10455/src/net/ipsock.go:230
//
//line /snap/go/10455/src/net/ipsock.go:230
// See func Dial for a description of the host and port parameters.
//line /snap/go/10455/src/net/ipsock.go:235
func JoinHostPort(host, port string) string {
//line /snap/go/10455/src/net/ipsock.go:235
	_go_fuzz_dep_.CoverTab[6791]++

//line /snap/go/10455/src/net/ipsock.go:238
	if bytealg.IndexByteString(host, ':') >= 0 {
//line /snap/go/10455/src/net/ipsock.go:238
		_go_fuzz_dep_.CoverTab[528834]++
//line /snap/go/10455/src/net/ipsock.go:238
		_go_fuzz_dep_.CoverTab[6793]++
							return "[" + host + "]:" + port
//line /snap/go/10455/src/net/ipsock.go:239
		// _ = "end of CoverTab[6793]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:240
		_go_fuzz_dep_.CoverTab[528835]++
//line /snap/go/10455/src/net/ipsock.go:240
		_go_fuzz_dep_.CoverTab[6794]++
//line /snap/go/10455/src/net/ipsock.go:240
		// _ = "end of CoverTab[6794]"
//line /snap/go/10455/src/net/ipsock.go:240
	}
//line /snap/go/10455/src/net/ipsock.go:240
	// _ = "end of CoverTab[6791]"
//line /snap/go/10455/src/net/ipsock.go:240
	_go_fuzz_dep_.CoverTab[6792]++
						return host + ":" + port
//line /snap/go/10455/src/net/ipsock.go:241
	// _ = "end of CoverTab[6792]"
}

// internetAddrList resolves addr, which may be a literal IP
//line /snap/go/10455/src/net/ipsock.go:244
// address or a DNS name, and returns a list of internet protocol
//line /snap/go/10455/src/net/ipsock.go:244
// family addresses. The result contains at least one address when
//line /snap/go/10455/src/net/ipsock.go:244
// error is nil.
//line /snap/go/10455/src/net/ipsock.go:248
func (r *Resolver) internetAddrList(ctx context.Context, net, addr string) (addrList, error) {
//line /snap/go/10455/src/net/ipsock.go:248
	_go_fuzz_dep_.CoverTab[6795]++
						var (
		err		error
		host, port	string
		portnum		int
	)
	switch net {
	case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
//line /snap/go/10455/src/net/ipsock.go:255
		_go_fuzz_dep_.CoverTab[528836]++
//line /snap/go/10455/src/net/ipsock.go:255
		_go_fuzz_dep_.CoverTab[6803]++
							if addr != "" {
//line /snap/go/10455/src/net/ipsock.go:256
			_go_fuzz_dep_.CoverTab[528839]++
//line /snap/go/10455/src/net/ipsock.go:256
			_go_fuzz_dep_.CoverTab[6806]++
								if host, port, err = SplitHostPort(addr); err != nil {
//line /snap/go/10455/src/net/ipsock.go:257
				_go_fuzz_dep_.CoverTab[528841]++
//line /snap/go/10455/src/net/ipsock.go:257
				_go_fuzz_dep_.CoverTab[6808]++
									return nil, err
//line /snap/go/10455/src/net/ipsock.go:258
				// _ = "end of CoverTab[6808]"
			} else {
//line /snap/go/10455/src/net/ipsock.go:259
				_go_fuzz_dep_.CoverTab[528842]++
//line /snap/go/10455/src/net/ipsock.go:259
				_go_fuzz_dep_.CoverTab[6809]++
//line /snap/go/10455/src/net/ipsock.go:259
				// _ = "end of CoverTab[6809]"
//line /snap/go/10455/src/net/ipsock.go:259
			}
//line /snap/go/10455/src/net/ipsock.go:259
			// _ = "end of CoverTab[6806]"
//line /snap/go/10455/src/net/ipsock.go:259
			_go_fuzz_dep_.CoverTab[6807]++
								if portnum, err = r.LookupPort(ctx, net, port); err != nil {
//line /snap/go/10455/src/net/ipsock.go:260
				_go_fuzz_dep_.CoverTab[528843]++
//line /snap/go/10455/src/net/ipsock.go:260
				_go_fuzz_dep_.CoverTab[6810]++
									return nil, err
//line /snap/go/10455/src/net/ipsock.go:261
				// _ = "end of CoverTab[6810]"
			} else {
//line /snap/go/10455/src/net/ipsock.go:262
				_go_fuzz_dep_.CoverTab[528844]++
//line /snap/go/10455/src/net/ipsock.go:262
				_go_fuzz_dep_.CoverTab[6811]++
//line /snap/go/10455/src/net/ipsock.go:262
				// _ = "end of CoverTab[6811]"
//line /snap/go/10455/src/net/ipsock.go:262
			}
//line /snap/go/10455/src/net/ipsock.go:262
			// _ = "end of CoverTab[6807]"
		} else {
//line /snap/go/10455/src/net/ipsock.go:263
			_go_fuzz_dep_.CoverTab[528840]++
//line /snap/go/10455/src/net/ipsock.go:263
			_go_fuzz_dep_.CoverTab[6812]++
//line /snap/go/10455/src/net/ipsock.go:263
			// _ = "end of CoverTab[6812]"
//line /snap/go/10455/src/net/ipsock.go:263
		}
//line /snap/go/10455/src/net/ipsock.go:263
		// _ = "end of CoverTab[6803]"
	case "ip", "ip4", "ip6":
//line /snap/go/10455/src/net/ipsock.go:264
		_go_fuzz_dep_.CoverTab[528837]++
//line /snap/go/10455/src/net/ipsock.go:264
		_go_fuzz_dep_.CoverTab[6804]++
							if addr != "" {
//line /snap/go/10455/src/net/ipsock.go:265
			_go_fuzz_dep_.CoverTab[528845]++
//line /snap/go/10455/src/net/ipsock.go:265
			_go_fuzz_dep_.CoverTab[6813]++
								host = addr
//line /snap/go/10455/src/net/ipsock.go:266
			// _ = "end of CoverTab[6813]"
		} else {
//line /snap/go/10455/src/net/ipsock.go:267
			_go_fuzz_dep_.CoverTab[528846]++
//line /snap/go/10455/src/net/ipsock.go:267
			_go_fuzz_dep_.CoverTab[6814]++
//line /snap/go/10455/src/net/ipsock.go:267
			// _ = "end of CoverTab[6814]"
//line /snap/go/10455/src/net/ipsock.go:267
		}
//line /snap/go/10455/src/net/ipsock.go:267
		// _ = "end of CoverTab[6804]"
	default:
//line /snap/go/10455/src/net/ipsock.go:268
		_go_fuzz_dep_.CoverTab[528838]++
//line /snap/go/10455/src/net/ipsock.go:268
		_go_fuzz_dep_.CoverTab[6805]++
							return nil, UnknownNetworkError(net)
//line /snap/go/10455/src/net/ipsock.go:269
		// _ = "end of CoverTab[6805]"
	}
//line /snap/go/10455/src/net/ipsock.go:270
	// _ = "end of CoverTab[6795]"
//line /snap/go/10455/src/net/ipsock.go:270
	_go_fuzz_dep_.CoverTab[6796]++
						inetaddr := func(ip IPAddr) Addr {
//line /snap/go/10455/src/net/ipsock.go:271
		_go_fuzz_dep_.CoverTab[6815]++
							switch net {
		case "tcp", "tcp4", "tcp6":
//line /snap/go/10455/src/net/ipsock.go:273
			_go_fuzz_dep_.CoverTab[528847]++
//line /snap/go/10455/src/net/ipsock.go:273
			_go_fuzz_dep_.CoverTab[6816]++
								return &TCPAddr{IP: ip.IP, Port: portnum, Zone: ip.Zone}
//line /snap/go/10455/src/net/ipsock.go:274
			// _ = "end of CoverTab[6816]"
		case "udp", "udp4", "udp6":
//line /snap/go/10455/src/net/ipsock.go:275
			_go_fuzz_dep_.CoverTab[528848]++
//line /snap/go/10455/src/net/ipsock.go:275
			_go_fuzz_dep_.CoverTab[6817]++
								return &UDPAddr{IP: ip.IP, Port: portnum, Zone: ip.Zone}
//line /snap/go/10455/src/net/ipsock.go:276
			// _ = "end of CoverTab[6817]"
		case "ip", "ip4", "ip6":
//line /snap/go/10455/src/net/ipsock.go:277
			_go_fuzz_dep_.CoverTab[528849]++
//line /snap/go/10455/src/net/ipsock.go:277
			_go_fuzz_dep_.CoverTab[6818]++
								return &IPAddr{IP: ip.IP, Zone: ip.Zone}
//line /snap/go/10455/src/net/ipsock.go:278
			// _ = "end of CoverTab[6818]"
		default:
//line /snap/go/10455/src/net/ipsock.go:279
			_go_fuzz_dep_.CoverTab[528850]++
//line /snap/go/10455/src/net/ipsock.go:279
			_go_fuzz_dep_.CoverTab[6819]++
								panic("unexpected network: " + net)
//line /snap/go/10455/src/net/ipsock.go:280
			// _ = "end of CoverTab[6819]"
		}
//line /snap/go/10455/src/net/ipsock.go:281
		// _ = "end of CoverTab[6815]"
	}
//line /snap/go/10455/src/net/ipsock.go:282
	// _ = "end of CoverTab[6796]"
//line /snap/go/10455/src/net/ipsock.go:282
	_go_fuzz_dep_.CoverTab[6797]++
						if host == "" {
//line /snap/go/10455/src/net/ipsock.go:283
		_go_fuzz_dep_.CoverTab[528851]++
//line /snap/go/10455/src/net/ipsock.go:283
		_go_fuzz_dep_.CoverTab[6820]++
							return addrList{inetaddr(IPAddr{})}, nil
//line /snap/go/10455/src/net/ipsock.go:284
		// _ = "end of CoverTab[6820]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:285
		_go_fuzz_dep_.CoverTab[528852]++
//line /snap/go/10455/src/net/ipsock.go:285
		_go_fuzz_dep_.CoverTab[6821]++
//line /snap/go/10455/src/net/ipsock.go:285
		// _ = "end of CoverTab[6821]"
//line /snap/go/10455/src/net/ipsock.go:285
	}
//line /snap/go/10455/src/net/ipsock.go:285
	// _ = "end of CoverTab[6797]"
//line /snap/go/10455/src/net/ipsock.go:285
	_go_fuzz_dep_.CoverTab[6798]++

//line /snap/go/10455/src/net/ipsock.go:288
	ips, err := r.lookupIPAddr(ctx, net, host)
	if err != nil {
//line /snap/go/10455/src/net/ipsock.go:289
		_go_fuzz_dep_.CoverTab[528853]++
//line /snap/go/10455/src/net/ipsock.go:289
		_go_fuzz_dep_.CoverTab[6822]++
							return nil, err
//line /snap/go/10455/src/net/ipsock.go:290
		// _ = "end of CoverTab[6822]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:291
		_go_fuzz_dep_.CoverTab[528854]++
//line /snap/go/10455/src/net/ipsock.go:291
		_go_fuzz_dep_.CoverTab[6823]++
//line /snap/go/10455/src/net/ipsock.go:291
		// _ = "end of CoverTab[6823]"
//line /snap/go/10455/src/net/ipsock.go:291
	}
//line /snap/go/10455/src/net/ipsock.go:291
	// _ = "end of CoverTab[6798]"
//line /snap/go/10455/src/net/ipsock.go:291
	_go_fuzz_dep_.CoverTab[6799]++

//line /snap/go/10455/src/net/ipsock.go:296
	if len(ips) == 1 && func() bool {
//line /snap/go/10455/src/net/ipsock.go:296
		_go_fuzz_dep_.CoverTab[6824]++
//line /snap/go/10455/src/net/ipsock.go:296
		return ips[0].IP.Equal(IPv6unspecified)
//line /snap/go/10455/src/net/ipsock.go:296
		// _ = "end of CoverTab[6824]"
//line /snap/go/10455/src/net/ipsock.go:296
	}() {
//line /snap/go/10455/src/net/ipsock.go:296
		_go_fuzz_dep_.CoverTab[528855]++
//line /snap/go/10455/src/net/ipsock.go:296
		_go_fuzz_dep_.CoverTab[6825]++
							ips = append(ips, IPAddr{IP: IPv4zero})
//line /snap/go/10455/src/net/ipsock.go:297
		// _ = "end of CoverTab[6825]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:298
		_go_fuzz_dep_.CoverTab[528856]++
//line /snap/go/10455/src/net/ipsock.go:298
		_go_fuzz_dep_.CoverTab[6826]++
//line /snap/go/10455/src/net/ipsock.go:298
		// _ = "end of CoverTab[6826]"
//line /snap/go/10455/src/net/ipsock.go:298
	}
//line /snap/go/10455/src/net/ipsock.go:298
	// _ = "end of CoverTab[6799]"
//line /snap/go/10455/src/net/ipsock.go:298
	_go_fuzz_dep_.CoverTab[6800]++

						var filter func(IPAddr) bool
						if net != "" && func() bool {
//line /snap/go/10455/src/net/ipsock.go:301
		_go_fuzz_dep_.CoverTab[6827]++
//line /snap/go/10455/src/net/ipsock.go:301
		return net[len(net)-1] == '4'
//line /snap/go/10455/src/net/ipsock.go:301
		// _ = "end of CoverTab[6827]"
//line /snap/go/10455/src/net/ipsock.go:301
	}() {
//line /snap/go/10455/src/net/ipsock.go:301
		_go_fuzz_dep_.CoverTab[528857]++
//line /snap/go/10455/src/net/ipsock.go:301
		_go_fuzz_dep_.CoverTab[6828]++
							filter = ipv4only
//line /snap/go/10455/src/net/ipsock.go:302
		// _ = "end of CoverTab[6828]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:303
		_go_fuzz_dep_.CoverTab[528858]++
//line /snap/go/10455/src/net/ipsock.go:303
		_go_fuzz_dep_.CoverTab[6829]++
//line /snap/go/10455/src/net/ipsock.go:303
		// _ = "end of CoverTab[6829]"
//line /snap/go/10455/src/net/ipsock.go:303
	}
//line /snap/go/10455/src/net/ipsock.go:303
	// _ = "end of CoverTab[6800]"
//line /snap/go/10455/src/net/ipsock.go:303
	_go_fuzz_dep_.CoverTab[6801]++
						if net != "" && func() bool {
//line /snap/go/10455/src/net/ipsock.go:304
		_go_fuzz_dep_.CoverTab[6830]++
//line /snap/go/10455/src/net/ipsock.go:304
		return net[len(net)-1] == '6'
//line /snap/go/10455/src/net/ipsock.go:304
		// _ = "end of CoverTab[6830]"
//line /snap/go/10455/src/net/ipsock.go:304
	}() {
//line /snap/go/10455/src/net/ipsock.go:304
		_go_fuzz_dep_.CoverTab[528859]++
//line /snap/go/10455/src/net/ipsock.go:304
		_go_fuzz_dep_.CoverTab[6831]++
							filter = ipv6only
//line /snap/go/10455/src/net/ipsock.go:305
		// _ = "end of CoverTab[6831]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:306
		_go_fuzz_dep_.CoverTab[528860]++
//line /snap/go/10455/src/net/ipsock.go:306
		_go_fuzz_dep_.CoverTab[6832]++
//line /snap/go/10455/src/net/ipsock.go:306
		// _ = "end of CoverTab[6832]"
//line /snap/go/10455/src/net/ipsock.go:306
	}
//line /snap/go/10455/src/net/ipsock.go:306
	// _ = "end of CoverTab[6801]"
//line /snap/go/10455/src/net/ipsock.go:306
	_go_fuzz_dep_.CoverTab[6802]++
						return filterAddrList(filter, ips, inetaddr, host)
//line /snap/go/10455/src/net/ipsock.go:307
	// _ = "end of CoverTab[6802]"
}

func loopbackIP(net string) IP {
//line /snap/go/10455/src/net/ipsock.go:310
	_go_fuzz_dep_.CoverTab[6833]++
						if net != "" && func() bool {
//line /snap/go/10455/src/net/ipsock.go:311
		_go_fuzz_dep_.CoverTab[6835]++
//line /snap/go/10455/src/net/ipsock.go:311
		return net[len(net)-1] == '6'
//line /snap/go/10455/src/net/ipsock.go:311
		// _ = "end of CoverTab[6835]"
//line /snap/go/10455/src/net/ipsock.go:311
	}() {
//line /snap/go/10455/src/net/ipsock.go:311
		_go_fuzz_dep_.CoverTab[528861]++
//line /snap/go/10455/src/net/ipsock.go:311
		_go_fuzz_dep_.CoverTab[6836]++
							return IPv6loopback
//line /snap/go/10455/src/net/ipsock.go:312
		// _ = "end of CoverTab[6836]"
	} else {
//line /snap/go/10455/src/net/ipsock.go:313
		_go_fuzz_dep_.CoverTab[528862]++
//line /snap/go/10455/src/net/ipsock.go:313
		_go_fuzz_dep_.CoverTab[6837]++
//line /snap/go/10455/src/net/ipsock.go:313
		// _ = "end of CoverTab[6837]"
//line /snap/go/10455/src/net/ipsock.go:313
	}
//line /snap/go/10455/src/net/ipsock.go:313
	// _ = "end of CoverTab[6833]"
//line /snap/go/10455/src/net/ipsock.go:313
	_go_fuzz_dep_.CoverTab[6834]++
						return IP{127, 0, 0, 1}
//line /snap/go/10455/src/net/ipsock.go:314
	// _ = "end of CoverTab[6834]"
}

//line /snap/go/10455/src/net/ipsock.go:315
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/ipsock.go:315
var _ = _go_fuzz_dep_.CoverTab
