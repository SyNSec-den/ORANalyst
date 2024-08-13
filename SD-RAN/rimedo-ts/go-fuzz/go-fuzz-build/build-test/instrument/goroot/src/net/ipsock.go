// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/ipsock.go:5
package net

//line /usr/local/go/src/net/ipsock.go:5
import (
//line /usr/local/go/src/net/ipsock.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/ipsock.go:5
)
//line /usr/local/go/src/net/ipsock.go:5
import (
//line /usr/local/go/src/net/ipsock.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/ipsock.go:5
)

import (
	"context"
	"internal/bytealg"
	"runtime"
	"sync"
)

//line /usr/local/go/src/net/ipsock.go:21
type ipStackCapabilities struct {
	sync.Once		// guards following
	ipv4Enabled		bool
	ipv6Enabled		bool
	ipv4MappedIPv6Enabled	bool
}

var ipStackCaps ipStackCapabilities

// supportsIPv4 reports whether the platform supports IPv4 networking
//line /usr/local/go/src/net/ipsock.go:30
// functionality.
//line /usr/local/go/src/net/ipsock.go:32
func supportsIPv4() bool {
//line /usr/local/go/src/net/ipsock.go:32
	_go_fuzz_dep_.CoverTab[6451]++
						ipStackCaps.Once.Do(ipStackCaps.probe)
						return ipStackCaps.ipv4Enabled
//line /usr/local/go/src/net/ipsock.go:34
	// _ = "end of CoverTab[6451]"
}

// supportsIPv6 reports whether the platform supports IPv6 networking
//line /usr/local/go/src/net/ipsock.go:37
// functionality.
//line /usr/local/go/src/net/ipsock.go:39
func supportsIPv6() bool {
//line /usr/local/go/src/net/ipsock.go:39
	_go_fuzz_dep_.CoverTab[6452]++
						ipStackCaps.Once.Do(ipStackCaps.probe)
						return ipStackCaps.ipv6Enabled
//line /usr/local/go/src/net/ipsock.go:41
	// _ = "end of CoverTab[6452]"
}

// supportsIPv4map reports whether the platform supports mapping an
//line /usr/local/go/src/net/ipsock.go:44
// IPv4 address inside an IPv6 address at transport layer
//line /usr/local/go/src/net/ipsock.go:44
// protocols. See RFC 4291, RFC 4038 and RFC 3493.
//line /usr/local/go/src/net/ipsock.go:47
func supportsIPv4map() bool {
//line /usr/local/go/src/net/ipsock.go:47
	_go_fuzz_dep_.CoverTab[6453]++

//line /usr/local/go/src/net/ipsock.go:50
	switch runtime.GOOS {
	case "dragonfly", "openbsd":
//line /usr/local/go/src/net/ipsock.go:51
		_go_fuzz_dep_.CoverTab[6455]++
							return false
//line /usr/local/go/src/net/ipsock.go:52
		// _ = "end of CoverTab[6455]"
//line /usr/local/go/src/net/ipsock.go:52
	default:
//line /usr/local/go/src/net/ipsock.go:52
		_go_fuzz_dep_.CoverTab[6456]++
//line /usr/local/go/src/net/ipsock.go:52
		// _ = "end of CoverTab[6456]"
	}
//line /usr/local/go/src/net/ipsock.go:53
	// _ = "end of CoverTab[6453]"
//line /usr/local/go/src/net/ipsock.go:53
	_go_fuzz_dep_.CoverTab[6454]++

						ipStackCaps.Once.Do(ipStackCaps.probe)
						return ipStackCaps.ipv4MappedIPv6Enabled
//line /usr/local/go/src/net/ipsock.go:56
	// _ = "end of CoverTab[6454]"
}

// An addrList represents a list of network endpoint addresses.
type addrList []Addr

// isIPv4 reports whether addr contains an IPv4 address.
func isIPv4(addr Addr) bool {
//line /usr/local/go/src/net/ipsock.go:63
	_go_fuzz_dep_.CoverTab[6457]++
						switch addr := addr.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/ipsock.go:65
		_go_fuzz_dep_.CoverTab[6459]++
							return addr.IP.To4() != nil
//line /usr/local/go/src/net/ipsock.go:66
		// _ = "end of CoverTab[6459]"
	case *UDPAddr:
//line /usr/local/go/src/net/ipsock.go:67
		_go_fuzz_dep_.CoverTab[6460]++
							return addr.IP.To4() != nil
//line /usr/local/go/src/net/ipsock.go:68
		// _ = "end of CoverTab[6460]"
	case *IPAddr:
//line /usr/local/go/src/net/ipsock.go:69
		_go_fuzz_dep_.CoverTab[6461]++
							return addr.IP.To4() != nil
//line /usr/local/go/src/net/ipsock.go:70
		// _ = "end of CoverTab[6461]"
	}
//line /usr/local/go/src/net/ipsock.go:71
	// _ = "end of CoverTab[6457]"
//line /usr/local/go/src/net/ipsock.go:71
	_go_fuzz_dep_.CoverTab[6458]++
						return false
//line /usr/local/go/src/net/ipsock.go:72
	// _ = "end of CoverTab[6458]"
}

// isNotIPv4 reports whether addr does not contain an IPv4 address.
func isNotIPv4(addr Addr) bool {
//line /usr/local/go/src/net/ipsock.go:76
	_go_fuzz_dep_.CoverTab[6462]++
//line /usr/local/go/src/net/ipsock.go:76
	return !isIPv4(addr)
//line /usr/local/go/src/net/ipsock.go:76
	// _ = "end of CoverTab[6462]"
//line /usr/local/go/src/net/ipsock.go:76
}

// forResolve returns the most appropriate address in address for
//line /usr/local/go/src/net/ipsock.go:78
// a call to ResolveTCPAddr, ResolveUDPAddr, or ResolveIPAddr.
//line /usr/local/go/src/net/ipsock.go:78
// IPv4 is preferred, unless addr contains an IPv6 literal.
//line /usr/local/go/src/net/ipsock.go:81
func (addrs addrList) forResolve(network, addr string) Addr {
//line /usr/local/go/src/net/ipsock.go:81
	_go_fuzz_dep_.CoverTab[6463]++
						var want6 bool
						switch network {
	case "ip":
//line /usr/local/go/src/net/ipsock.go:84
		_go_fuzz_dep_.CoverTab[6466]++

							want6 = count(addr, ':') > 0
//line /usr/local/go/src/net/ipsock.go:86
		// _ = "end of CoverTab[6466]"
	case "tcp", "udp":
//line /usr/local/go/src/net/ipsock.go:87
		_go_fuzz_dep_.CoverTab[6467]++

							want6 = count(addr, '[') > 0
//line /usr/local/go/src/net/ipsock.go:89
		// _ = "end of CoverTab[6467]"
//line /usr/local/go/src/net/ipsock.go:89
	default:
//line /usr/local/go/src/net/ipsock.go:89
		_go_fuzz_dep_.CoverTab[6468]++
//line /usr/local/go/src/net/ipsock.go:89
		// _ = "end of CoverTab[6468]"
	}
//line /usr/local/go/src/net/ipsock.go:90
	// _ = "end of CoverTab[6463]"
//line /usr/local/go/src/net/ipsock.go:90
	_go_fuzz_dep_.CoverTab[6464]++
						if want6 {
//line /usr/local/go/src/net/ipsock.go:91
		_go_fuzz_dep_.CoverTab[6469]++
							return addrs.first(isNotIPv4)
//line /usr/local/go/src/net/ipsock.go:92
		// _ = "end of CoverTab[6469]"
	} else {
//line /usr/local/go/src/net/ipsock.go:93
		_go_fuzz_dep_.CoverTab[6470]++
//line /usr/local/go/src/net/ipsock.go:93
		// _ = "end of CoverTab[6470]"
//line /usr/local/go/src/net/ipsock.go:93
	}
//line /usr/local/go/src/net/ipsock.go:93
	// _ = "end of CoverTab[6464]"
//line /usr/local/go/src/net/ipsock.go:93
	_go_fuzz_dep_.CoverTab[6465]++
						return addrs.first(isIPv4)
//line /usr/local/go/src/net/ipsock.go:94
	// _ = "end of CoverTab[6465]"
}

// first returns the first address which satisfies strategy, or if
//line /usr/local/go/src/net/ipsock.go:97
// none do, then the first address of any kind.
//line /usr/local/go/src/net/ipsock.go:99
func (addrs addrList) first(strategy func(Addr) bool) Addr {
//line /usr/local/go/src/net/ipsock.go:99
	_go_fuzz_dep_.CoverTab[6471]++
						for _, addr := range addrs {
//line /usr/local/go/src/net/ipsock.go:100
		_go_fuzz_dep_.CoverTab[6473]++
							if strategy(addr) {
//line /usr/local/go/src/net/ipsock.go:101
			_go_fuzz_dep_.CoverTab[6474]++
								return addr
//line /usr/local/go/src/net/ipsock.go:102
			// _ = "end of CoverTab[6474]"
		} else {
//line /usr/local/go/src/net/ipsock.go:103
			_go_fuzz_dep_.CoverTab[6475]++
//line /usr/local/go/src/net/ipsock.go:103
			// _ = "end of CoverTab[6475]"
//line /usr/local/go/src/net/ipsock.go:103
		}
//line /usr/local/go/src/net/ipsock.go:103
		// _ = "end of CoverTab[6473]"
	}
//line /usr/local/go/src/net/ipsock.go:104
	// _ = "end of CoverTab[6471]"
//line /usr/local/go/src/net/ipsock.go:104
	_go_fuzz_dep_.CoverTab[6472]++
						return addrs[0]
//line /usr/local/go/src/net/ipsock.go:105
	// _ = "end of CoverTab[6472]"
}

// partition divides an address list into two categories, using a
//line /usr/local/go/src/net/ipsock.go:108
// strategy function to assign a boolean label to each address.
//line /usr/local/go/src/net/ipsock.go:108
// The first address, and any with a matching label, are returned as
//line /usr/local/go/src/net/ipsock.go:108
// primaries, while addresses with the opposite label are returned
//line /usr/local/go/src/net/ipsock.go:108
// as fallbacks. For non-empty inputs, primaries is guaranteed to be
//line /usr/local/go/src/net/ipsock.go:108
// non-empty.
//line /usr/local/go/src/net/ipsock.go:114
func (addrs addrList) partition(strategy func(Addr) bool) (primaries, fallbacks addrList) {
//line /usr/local/go/src/net/ipsock.go:114
	_go_fuzz_dep_.CoverTab[6476]++
						var primaryLabel bool
						for i, addr := range addrs {
//line /usr/local/go/src/net/ipsock.go:116
		_go_fuzz_dep_.CoverTab[6478]++
							label := strategy(addr)
							if i == 0 || func() bool {
//line /usr/local/go/src/net/ipsock.go:118
			_go_fuzz_dep_.CoverTab[6479]++
//line /usr/local/go/src/net/ipsock.go:118
			return label == primaryLabel
//line /usr/local/go/src/net/ipsock.go:118
			// _ = "end of CoverTab[6479]"
//line /usr/local/go/src/net/ipsock.go:118
		}() {
//line /usr/local/go/src/net/ipsock.go:118
			_go_fuzz_dep_.CoverTab[6480]++
								primaryLabel = label
								primaries = append(primaries, addr)
//line /usr/local/go/src/net/ipsock.go:120
			// _ = "end of CoverTab[6480]"
		} else {
//line /usr/local/go/src/net/ipsock.go:121
			_go_fuzz_dep_.CoverTab[6481]++
								fallbacks = append(fallbacks, addr)
//line /usr/local/go/src/net/ipsock.go:122
			// _ = "end of CoverTab[6481]"
		}
//line /usr/local/go/src/net/ipsock.go:123
		// _ = "end of CoverTab[6478]"
	}
//line /usr/local/go/src/net/ipsock.go:124
	// _ = "end of CoverTab[6476]"
//line /usr/local/go/src/net/ipsock.go:124
	_go_fuzz_dep_.CoverTab[6477]++
						return
//line /usr/local/go/src/net/ipsock.go:125
	// _ = "end of CoverTab[6477]"
}

// filterAddrList applies a filter to a list of IP addresses,
//line /usr/local/go/src/net/ipsock.go:128
// yielding a list of Addr objects. Known filters are nil, ipv4only,
//line /usr/local/go/src/net/ipsock.go:128
// and ipv6only. It returns every address when the filter is nil.
//line /usr/local/go/src/net/ipsock.go:128
// The result contains at least one address when error is nil.
//line /usr/local/go/src/net/ipsock.go:132
func filterAddrList(filter func(IPAddr) bool, ips []IPAddr, inetaddr func(IPAddr) Addr, originalAddr string) (addrList, error) {
//line /usr/local/go/src/net/ipsock.go:132
	_go_fuzz_dep_.CoverTab[6482]++
						var addrs addrList
						for _, ip := range ips {
//line /usr/local/go/src/net/ipsock.go:134
		_go_fuzz_dep_.CoverTab[6485]++
							if filter == nil || func() bool {
//line /usr/local/go/src/net/ipsock.go:135
			_go_fuzz_dep_.CoverTab[6486]++
//line /usr/local/go/src/net/ipsock.go:135
			return filter(ip)
//line /usr/local/go/src/net/ipsock.go:135
			// _ = "end of CoverTab[6486]"
//line /usr/local/go/src/net/ipsock.go:135
		}() {
//line /usr/local/go/src/net/ipsock.go:135
			_go_fuzz_dep_.CoverTab[6487]++
								addrs = append(addrs, inetaddr(ip))
//line /usr/local/go/src/net/ipsock.go:136
			// _ = "end of CoverTab[6487]"
		} else {
//line /usr/local/go/src/net/ipsock.go:137
			_go_fuzz_dep_.CoverTab[6488]++
//line /usr/local/go/src/net/ipsock.go:137
			// _ = "end of CoverTab[6488]"
//line /usr/local/go/src/net/ipsock.go:137
		}
//line /usr/local/go/src/net/ipsock.go:137
		// _ = "end of CoverTab[6485]"
	}
//line /usr/local/go/src/net/ipsock.go:138
	// _ = "end of CoverTab[6482]"
//line /usr/local/go/src/net/ipsock.go:138
	_go_fuzz_dep_.CoverTab[6483]++
						if len(addrs) == 0 {
//line /usr/local/go/src/net/ipsock.go:139
		_go_fuzz_dep_.CoverTab[6489]++
							return nil, &AddrError{Err: errNoSuitableAddress.Error(), Addr: originalAddr}
//line /usr/local/go/src/net/ipsock.go:140
		// _ = "end of CoverTab[6489]"
	} else {
//line /usr/local/go/src/net/ipsock.go:141
		_go_fuzz_dep_.CoverTab[6490]++
//line /usr/local/go/src/net/ipsock.go:141
		// _ = "end of CoverTab[6490]"
//line /usr/local/go/src/net/ipsock.go:141
	}
//line /usr/local/go/src/net/ipsock.go:141
	// _ = "end of CoverTab[6483]"
//line /usr/local/go/src/net/ipsock.go:141
	_go_fuzz_dep_.CoverTab[6484]++
						return addrs, nil
//line /usr/local/go/src/net/ipsock.go:142
	// _ = "end of CoverTab[6484]"
}

// ipv4only reports whether addr is an IPv4 address.
func ipv4only(addr IPAddr) bool {
//line /usr/local/go/src/net/ipsock.go:146
	_go_fuzz_dep_.CoverTab[6491]++
						return addr.IP.To4() != nil
//line /usr/local/go/src/net/ipsock.go:147
	// _ = "end of CoverTab[6491]"
}

// ipv6only reports whether addr is an IPv6 address except IPv4-mapped IPv6 address.
func ipv6only(addr IPAddr) bool {
//line /usr/local/go/src/net/ipsock.go:151
	_go_fuzz_dep_.CoverTab[6492]++
						return len(addr.IP) == IPv6len && func() bool {
//line /usr/local/go/src/net/ipsock.go:152
		_go_fuzz_dep_.CoverTab[6493]++
//line /usr/local/go/src/net/ipsock.go:152
		return addr.IP.To4() == nil
//line /usr/local/go/src/net/ipsock.go:152
		// _ = "end of CoverTab[6493]"
//line /usr/local/go/src/net/ipsock.go:152
	}()
//line /usr/local/go/src/net/ipsock.go:152
	// _ = "end of CoverTab[6492]"
}

// SplitHostPort splits a network address of the form "host:port",
//line /usr/local/go/src/net/ipsock.go:155
// "host%zone:port", "[host]:port" or "[host%zone]:port" into host or
//line /usr/local/go/src/net/ipsock.go:155
// host%zone and port.
//line /usr/local/go/src/net/ipsock.go:155
//
//line /usr/local/go/src/net/ipsock.go:155
// A literal IPv6 address in hostport must be enclosed in square
//line /usr/local/go/src/net/ipsock.go:155
// brackets, as in "[::1]:80", "[::1%lo0]:80".
//line /usr/local/go/src/net/ipsock.go:155
//
//line /usr/local/go/src/net/ipsock.go:155
// See func Dial for a description of the hostport parameter, and host
//line /usr/local/go/src/net/ipsock.go:155
// and port results.
//line /usr/local/go/src/net/ipsock.go:164
func SplitHostPort(hostport string) (host, port string, err error) {
//line /usr/local/go/src/net/ipsock.go:164
	_go_fuzz_dep_.CoverTab[6494]++
						const (
		missingPort	= "missing port in address"
		tooManyColons	= "too many colons in address"
	)
	addrErr := func(addr, why string) (host, port string, err error) {
//line /usr/local/go/src/net/ipsock.go:169
		_go_fuzz_dep_.CoverTab[6500]++
							return "", "", &AddrError{Err: why, Addr: addr}
//line /usr/local/go/src/net/ipsock.go:170
		// _ = "end of CoverTab[6500]"
	}
//line /usr/local/go/src/net/ipsock.go:171
	// _ = "end of CoverTab[6494]"
//line /usr/local/go/src/net/ipsock.go:171
	_go_fuzz_dep_.CoverTab[6495]++
						j, k := 0, 0

//line /usr/local/go/src/net/ipsock.go:175
	i := last(hostport, ':')
	if i < 0 {
//line /usr/local/go/src/net/ipsock.go:176
		_go_fuzz_dep_.CoverTab[6501]++
							return addrErr(hostport, missingPort)
//line /usr/local/go/src/net/ipsock.go:177
		// _ = "end of CoverTab[6501]"
	} else {
//line /usr/local/go/src/net/ipsock.go:178
		_go_fuzz_dep_.CoverTab[6502]++
//line /usr/local/go/src/net/ipsock.go:178
		// _ = "end of CoverTab[6502]"
//line /usr/local/go/src/net/ipsock.go:178
	}
//line /usr/local/go/src/net/ipsock.go:178
	// _ = "end of CoverTab[6495]"
//line /usr/local/go/src/net/ipsock.go:178
	_go_fuzz_dep_.CoverTab[6496]++

						if hostport[0] == '[' {
//line /usr/local/go/src/net/ipsock.go:180
		_go_fuzz_dep_.CoverTab[6503]++

							end := bytealg.IndexByteString(hostport, ']')
							if end < 0 {
//line /usr/local/go/src/net/ipsock.go:183
			_go_fuzz_dep_.CoverTab[6506]++
								return addrErr(hostport, "missing ']' in address")
//line /usr/local/go/src/net/ipsock.go:184
			// _ = "end of CoverTab[6506]"
		} else {
//line /usr/local/go/src/net/ipsock.go:185
			_go_fuzz_dep_.CoverTab[6507]++
//line /usr/local/go/src/net/ipsock.go:185
			// _ = "end of CoverTab[6507]"
//line /usr/local/go/src/net/ipsock.go:185
		}
//line /usr/local/go/src/net/ipsock.go:185
		// _ = "end of CoverTab[6503]"
//line /usr/local/go/src/net/ipsock.go:185
		_go_fuzz_dep_.CoverTab[6504]++
							switch end + 1 {
		case len(hostport):
//line /usr/local/go/src/net/ipsock.go:187
			_go_fuzz_dep_.CoverTab[6508]++

								return addrErr(hostport, missingPort)
//line /usr/local/go/src/net/ipsock.go:189
			// _ = "end of CoverTab[6508]"
		case i:
//line /usr/local/go/src/net/ipsock.go:190
			_go_fuzz_dep_.CoverTab[6509]++
//line /usr/local/go/src/net/ipsock.go:190
			// _ = "end of CoverTab[6509]"

		default:
//line /usr/local/go/src/net/ipsock.go:192
			_go_fuzz_dep_.CoverTab[6510]++

//line /usr/local/go/src/net/ipsock.go:195
			if hostport[end+1] == ':' {
//line /usr/local/go/src/net/ipsock.go:195
				_go_fuzz_dep_.CoverTab[6512]++
									return addrErr(hostport, tooManyColons)
//line /usr/local/go/src/net/ipsock.go:196
				// _ = "end of CoverTab[6512]"
			} else {
//line /usr/local/go/src/net/ipsock.go:197
				_go_fuzz_dep_.CoverTab[6513]++
//line /usr/local/go/src/net/ipsock.go:197
				// _ = "end of CoverTab[6513]"
//line /usr/local/go/src/net/ipsock.go:197
			}
//line /usr/local/go/src/net/ipsock.go:197
			// _ = "end of CoverTab[6510]"
//line /usr/local/go/src/net/ipsock.go:197
			_go_fuzz_dep_.CoverTab[6511]++
								return addrErr(hostport, missingPort)
//line /usr/local/go/src/net/ipsock.go:198
			// _ = "end of CoverTab[6511]"
		}
//line /usr/local/go/src/net/ipsock.go:199
		// _ = "end of CoverTab[6504]"
//line /usr/local/go/src/net/ipsock.go:199
		_go_fuzz_dep_.CoverTab[6505]++
							host = hostport[1:end]
							j, k = 1, end+1
//line /usr/local/go/src/net/ipsock.go:201
		// _ = "end of CoverTab[6505]"
	} else {
//line /usr/local/go/src/net/ipsock.go:202
		_go_fuzz_dep_.CoverTab[6514]++
							host = hostport[:i]
							if bytealg.IndexByteString(host, ':') >= 0 {
//line /usr/local/go/src/net/ipsock.go:204
			_go_fuzz_dep_.CoverTab[6515]++
								return addrErr(hostport, tooManyColons)
//line /usr/local/go/src/net/ipsock.go:205
			// _ = "end of CoverTab[6515]"
		} else {
//line /usr/local/go/src/net/ipsock.go:206
			_go_fuzz_dep_.CoverTab[6516]++
//line /usr/local/go/src/net/ipsock.go:206
			// _ = "end of CoverTab[6516]"
//line /usr/local/go/src/net/ipsock.go:206
		}
//line /usr/local/go/src/net/ipsock.go:206
		// _ = "end of CoverTab[6514]"
	}
//line /usr/local/go/src/net/ipsock.go:207
	// _ = "end of CoverTab[6496]"
//line /usr/local/go/src/net/ipsock.go:207
	_go_fuzz_dep_.CoverTab[6497]++
						if bytealg.IndexByteString(hostport[j:], '[') >= 0 {
//line /usr/local/go/src/net/ipsock.go:208
		_go_fuzz_dep_.CoverTab[6517]++
							return addrErr(hostport, "unexpected '[' in address")
//line /usr/local/go/src/net/ipsock.go:209
		// _ = "end of CoverTab[6517]"
	} else {
//line /usr/local/go/src/net/ipsock.go:210
		_go_fuzz_dep_.CoverTab[6518]++
//line /usr/local/go/src/net/ipsock.go:210
		// _ = "end of CoverTab[6518]"
//line /usr/local/go/src/net/ipsock.go:210
	}
//line /usr/local/go/src/net/ipsock.go:210
	// _ = "end of CoverTab[6497]"
//line /usr/local/go/src/net/ipsock.go:210
	_go_fuzz_dep_.CoverTab[6498]++
						if bytealg.IndexByteString(hostport[k:], ']') >= 0 {
//line /usr/local/go/src/net/ipsock.go:211
		_go_fuzz_dep_.CoverTab[6519]++
							return addrErr(hostport, "unexpected ']' in address")
//line /usr/local/go/src/net/ipsock.go:212
		// _ = "end of CoverTab[6519]"
	} else {
//line /usr/local/go/src/net/ipsock.go:213
		_go_fuzz_dep_.CoverTab[6520]++
//line /usr/local/go/src/net/ipsock.go:213
		// _ = "end of CoverTab[6520]"
//line /usr/local/go/src/net/ipsock.go:213
	}
//line /usr/local/go/src/net/ipsock.go:213
	// _ = "end of CoverTab[6498]"
//line /usr/local/go/src/net/ipsock.go:213
	_go_fuzz_dep_.CoverTab[6499]++

						port = hostport[i+1:]
						return host, port, nil
//line /usr/local/go/src/net/ipsock.go:216
	// _ = "end of CoverTab[6499]"
}

func splitHostZone(s string) (host, zone string) {
//line /usr/local/go/src/net/ipsock.go:219
	_go_fuzz_dep_.CoverTab[6521]++

//line /usr/local/go/src/net/ipsock.go:222
	if i := last(s, '%'); i > 0 {
//line /usr/local/go/src/net/ipsock.go:222
		_go_fuzz_dep_.CoverTab[6523]++
							host, zone = s[:i], s[i+1:]
//line /usr/local/go/src/net/ipsock.go:223
		// _ = "end of CoverTab[6523]"
	} else {
//line /usr/local/go/src/net/ipsock.go:224
		_go_fuzz_dep_.CoverTab[6524]++
							host = s
//line /usr/local/go/src/net/ipsock.go:225
		// _ = "end of CoverTab[6524]"
	}
//line /usr/local/go/src/net/ipsock.go:226
	// _ = "end of CoverTab[6521]"
//line /usr/local/go/src/net/ipsock.go:226
	_go_fuzz_dep_.CoverTab[6522]++
						return
//line /usr/local/go/src/net/ipsock.go:227
	// _ = "end of CoverTab[6522]"
}

// JoinHostPort combines host and port into a network address of the
//line /usr/local/go/src/net/ipsock.go:230
// form "host:port". If host contains a colon, as found in literal
//line /usr/local/go/src/net/ipsock.go:230
// IPv6 addresses, then JoinHostPort returns "[host]:port".
//line /usr/local/go/src/net/ipsock.go:230
//
//line /usr/local/go/src/net/ipsock.go:230
// See func Dial for a description of the host and port parameters.
//line /usr/local/go/src/net/ipsock.go:235
func JoinHostPort(host, port string) string {
//line /usr/local/go/src/net/ipsock.go:235
	_go_fuzz_dep_.CoverTab[6525]++

//line /usr/local/go/src/net/ipsock.go:238
	if bytealg.IndexByteString(host, ':') >= 0 {
//line /usr/local/go/src/net/ipsock.go:238
		_go_fuzz_dep_.CoverTab[6527]++
							return "[" + host + "]:" + port
//line /usr/local/go/src/net/ipsock.go:239
		// _ = "end of CoverTab[6527]"
	} else {
//line /usr/local/go/src/net/ipsock.go:240
		_go_fuzz_dep_.CoverTab[6528]++
//line /usr/local/go/src/net/ipsock.go:240
		// _ = "end of CoverTab[6528]"
//line /usr/local/go/src/net/ipsock.go:240
	}
//line /usr/local/go/src/net/ipsock.go:240
	// _ = "end of CoverTab[6525]"
//line /usr/local/go/src/net/ipsock.go:240
	_go_fuzz_dep_.CoverTab[6526]++
						return host + ":" + port
//line /usr/local/go/src/net/ipsock.go:241
	// _ = "end of CoverTab[6526]"
}

// internetAddrList resolves addr, which may be a literal IP
//line /usr/local/go/src/net/ipsock.go:244
// address or a DNS name, and returns a list of internet protocol
//line /usr/local/go/src/net/ipsock.go:244
// family addresses. The result contains at least one address when
//line /usr/local/go/src/net/ipsock.go:244
// error is nil.
//line /usr/local/go/src/net/ipsock.go:248
func (r *Resolver) internetAddrList(ctx context.Context, net, addr string) (addrList, error) {
//line /usr/local/go/src/net/ipsock.go:248
	_go_fuzz_dep_.CoverTab[6529]++
						var (
		err		error
		host, port	string
		portnum		int
	)
	switch net {
	case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
//line /usr/local/go/src/net/ipsock.go:255
		_go_fuzz_dep_.CoverTab[6537]++
							if addr != "" {
//line /usr/local/go/src/net/ipsock.go:256
			_go_fuzz_dep_.CoverTab[6540]++
								if host, port, err = SplitHostPort(addr); err != nil {
//line /usr/local/go/src/net/ipsock.go:257
				_go_fuzz_dep_.CoverTab[6542]++
									return nil, err
//line /usr/local/go/src/net/ipsock.go:258
				// _ = "end of CoverTab[6542]"
			} else {
//line /usr/local/go/src/net/ipsock.go:259
				_go_fuzz_dep_.CoverTab[6543]++
//line /usr/local/go/src/net/ipsock.go:259
				// _ = "end of CoverTab[6543]"
//line /usr/local/go/src/net/ipsock.go:259
			}
//line /usr/local/go/src/net/ipsock.go:259
			// _ = "end of CoverTab[6540]"
//line /usr/local/go/src/net/ipsock.go:259
			_go_fuzz_dep_.CoverTab[6541]++
								if portnum, err = r.LookupPort(ctx, net, port); err != nil {
//line /usr/local/go/src/net/ipsock.go:260
				_go_fuzz_dep_.CoverTab[6544]++
									return nil, err
//line /usr/local/go/src/net/ipsock.go:261
				// _ = "end of CoverTab[6544]"
			} else {
//line /usr/local/go/src/net/ipsock.go:262
				_go_fuzz_dep_.CoverTab[6545]++
//line /usr/local/go/src/net/ipsock.go:262
				// _ = "end of CoverTab[6545]"
//line /usr/local/go/src/net/ipsock.go:262
			}
//line /usr/local/go/src/net/ipsock.go:262
			// _ = "end of CoverTab[6541]"
		} else {
//line /usr/local/go/src/net/ipsock.go:263
			_go_fuzz_dep_.CoverTab[6546]++
//line /usr/local/go/src/net/ipsock.go:263
			// _ = "end of CoverTab[6546]"
//line /usr/local/go/src/net/ipsock.go:263
		}
//line /usr/local/go/src/net/ipsock.go:263
		// _ = "end of CoverTab[6537]"
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/ipsock.go:264
		_go_fuzz_dep_.CoverTab[6538]++
							if addr != "" {
//line /usr/local/go/src/net/ipsock.go:265
			_go_fuzz_dep_.CoverTab[6547]++
								host = addr
//line /usr/local/go/src/net/ipsock.go:266
			// _ = "end of CoverTab[6547]"
		} else {
//line /usr/local/go/src/net/ipsock.go:267
			_go_fuzz_dep_.CoverTab[6548]++
//line /usr/local/go/src/net/ipsock.go:267
			// _ = "end of CoverTab[6548]"
//line /usr/local/go/src/net/ipsock.go:267
		}
//line /usr/local/go/src/net/ipsock.go:267
		// _ = "end of CoverTab[6538]"
	default:
//line /usr/local/go/src/net/ipsock.go:268
		_go_fuzz_dep_.CoverTab[6539]++
							return nil, UnknownNetworkError(net)
//line /usr/local/go/src/net/ipsock.go:269
		// _ = "end of CoverTab[6539]"
	}
//line /usr/local/go/src/net/ipsock.go:270
	// _ = "end of CoverTab[6529]"
//line /usr/local/go/src/net/ipsock.go:270
	_go_fuzz_dep_.CoverTab[6530]++
						inetaddr := func(ip IPAddr) Addr {
//line /usr/local/go/src/net/ipsock.go:271
		_go_fuzz_dep_.CoverTab[6549]++
							switch net {
		case "tcp", "tcp4", "tcp6":
//line /usr/local/go/src/net/ipsock.go:273
			_go_fuzz_dep_.CoverTab[6550]++
								return &TCPAddr{IP: ip.IP, Port: portnum, Zone: ip.Zone}
//line /usr/local/go/src/net/ipsock.go:274
			// _ = "end of CoverTab[6550]"
		case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/ipsock.go:275
			_go_fuzz_dep_.CoverTab[6551]++
								return &UDPAddr{IP: ip.IP, Port: portnum, Zone: ip.Zone}
//line /usr/local/go/src/net/ipsock.go:276
			// _ = "end of CoverTab[6551]"
		case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/ipsock.go:277
			_go_fuzz_dep_.CoverTab[6552]++
								return &IPAddr{IP: ip.IP, Zone: ip.Zone}
//line /usr/local/go/src/net/ipsock.go:278
			// _ = "end of CoverTab[6552]"
		default:
//line /usr/local/go/src/net/ipsock.go:279
			_go_fuzz_dep_.CoverTab[6553]++
								panic("unexpected network: " + net)
//line /usr/local/go/src/net/ipsock.go:280
			// _ = "end of CoverTab[6553]"
		}
//line /usr/local/go/src/net/ipsock.go:281
		// _ = "end of CoverTab[6549]"
	}
//line /usr/local/go/src/net/ipsock.go:282
	// _ = "end of CoverTab[6530]"
//line /usr/local/go/src/net/ipsock.go:282
	_go_fuzz_dep_.CoverTab[6531]++
						if host == "" {
//line /usr/local/go/src/net/ipsock.go:283
		_go_fuzz_dep_.CoverTab[6554]++
							return addrList{inetaddr(IPAddr{})}, nil
//line /usr/local/go/src/net/ipsock.go:284
		// _ = "end of CoverTab[6554]"
	} else {
//line /usr/local/go/src/net/ipsock.go:285
		_go_fuzz_dep_.CoverTab[6555]++
//line /usr/local/go/src/net/ipsock.go:285
		// _ = "end of CoverTab[6555]"
//line /usr/local/go/src/net/ipsock.go:285
	}
//line /usr/local/go/src/net/ipsock.go:285
	// _ = "end of CoverTab[6531]"
//line /usr/local/go/src/net/ipsock.go:285
	_go_fuzz_dep_.CoverTab[6532]++

//line /usr/local/go/src/net/ipsock.go:288
	ips, err := r.lookupIPAddr(ctx, net, host)
	if err != nil {
//line /usr/local/go/src/net/ipsock.go:289
		_go_fuzz_dep_.CoverTab[6556]++
							return nil, err
//line /usr/local/go/src/net/ipsock.go:290
		// _ = "end of CoverTab[6556]"
	} else {
//line /usr/local/go/src/net/ipsock.go:291
		_go_fuzz_dep_.CoverTab[6557]++
//line /usr/local/go/src/net/ipsock.go:291
		// _ = "end of CoverTab[6557]"
//line /usr/local/go/src/net/ipsock.go:291
	}
//line /usr/local/go/src/net/ipsock.go:291
	// _ = "end of CoverTab[6532]"
//line /usr/local/go/src/net/ipsock.go:291
	_go_fuzz_dep_.CoverTab[6533]++

//line /usr/local/go/src/net/ipsock.go:296
	if len(ips) == 1 && func() bool {
//line /usr/local/go/src/net/ipsock.go:296
		_go_fuzz_dep_.CoverTab[6558]++
//line /usr/local/go/src/net/ipsock.go:296
		return ips[0].IP.Equal(IPv6unspecified)
//line /usr/local/go/src/net/ipsock.go:296
		// _ = "end of CoverTab[6558]"
//line /usr/local/go/src/net/ipsock.go:296
	}() {
//line /usr/local/go/src/net/ipsock.go:296
		_go_fuzz_dep_.CoverTab[6559]++
							ips = append(ips, IPAddr{IP: IPv4zero})
//line /usr/local/go/src/net/ipsock.go:297
		// _ = "end of CoverTab[6559]"
	} else {
//line /usr/local/go/src/net/ipsock.go:298
		_go_fuzz_dep_.CoverTab[6560]++
//line /usr/local/go/src/net/ipsock.go:298
		// _ = "end of CoverTab[6560]"
//line /usr/local/go/src/net/ipsock.go:298
	}
//line /usr/local/go/src/net/ipsock.go:298
	// _ = "end of CoverTab[6533]"
//line /usr/local/go/src/net/ipsock.go:298
	_go_fuzz_dep_.CoverTab[6534]++

						var filter func(IPAddr) bool
						if net != "" && func() bool {
//line /usr/local/go/src/net/ipsock.go:301
		_go_fuzz_dep_.CoverTab[6561]++
//line /usr/local/go/src/net/ipsock.go:301
		return net[len(net)-1] == '4'
//line /usr/local/go/src/net/ipsock.go:301
		// _ = "end of CoverTab[6561]"
//line /usr/local/go/src/net/ipsock.go:301
	}() {
//line /usr/local/go/src/net/ipsock.go:301
		_go_fuzz_dep_.CoverTab[6562]++
							filter = ipv4only
//line /usr/local/go/src/net/ipsock.go:302
		// _ = "end of CoverTab[6562]"
	} else {
//line /usr/local/go/src/net/ipsock.go:303
		_go_fuzz_dep_.CoverTab[6563]++
//line /usr/local/go/src/net/ipsock.go:303
		// _ = "end of CoverTab[6563]"
//line /usr/local/go/src/net/ipsock.go:303
	}
//line /usr/local/go/src/net/ipsock.go:303
	// _ = "end of CoverTab[6534]"
//line /usr/local/go/src/net/ipsock.go:303
	_go_fuzz_dep_.CoverTab[6535]++
						if net != "" && func() bool {
//line /usr/local/go/src/net/ipsock.go:304
		_go_fuzz_dep_.CoverTab[6564]++
//line /usr/local/go/src/net/ipsock.go:304
		return net[len(net)-1] == '6'
//line /usr/local/go/src/net/ipsock.go:304
		// _ = "end of CoverTab[6564]"
//line /usr/local/go/src/net/ipsock.go:304
	}() {
//line /usr/local/go/src/net/ipsock.go:304
		_go_fuzz_dep_.CoverTab[6565]++
							filter = ipv6only
//line /usr/local/go/src/net/ipsock.go:305
		// _ = "end of CoverTab[6565]"
	} else {
//line /usr/local/go/src/net/ipsock.go:306
		_go_fuzz_dep_.CoverTab[6566]++
//line /usr/local/go/src/net/ipsock.go:306
		// _ = "end of CoverTab[6566]"
//line /usr/local/go/src/net/ipsock.go:306
	}
//line /usr/local/go/src/net/ipsock.go:306
	// _ = "end of CoverTab[6535]"
//line /usr/local/go/src/net/ipsock.go:306
	_go_fuzz_dep_.CoverTab[6536]++
						return filterAddrList(filter, ips, inetaddr, host)
//line /usr/local/go/src/net/ipsock.go:307
	// _ = "end of CoverTab[6536]"
}

func loopbackIP(net string) IP {
//line /usr/local/go/src/net/ipsock.go:310
	_go_fuzz_dep_.CoverTab[6567]++
						if net != "" && func() bool {
//line /usr/local/go/src/net/ipsock.go:311
		_go_fuzz_dep_.CoverTab[6569]++
//line /usr/local/go/src/net/ipsock.go:311
		return net[len(net)-1] == '6'
//line /usr/local/go/src/net/ipsock.go:311
		// _ = "end of CoverTab[6569]"
//line /usr/local/go/src/net/ipsock.go:311
	}() {
//line /usr/local/go/src/net/ipsock.go:311
		_go_fuzz_dep_.CoverTab[6570]++
							return IPv6loopback
//line /usr/local/go/src/net/ipsock.go:312
		// _ = "end of CoverTab[6570]"
	} else {
//line /usr/local/go/src/net/ipsock.go:313
		_go_fuzz_dep_.CoverTab[6571]++
//line /usr/local/go/src/net/ipsock.go:313
		// _ = "end of CoverTab[6571]"
//line /usr/local/go/src/net/ipsock.go:313
	}
//line /usr/local/go/src/net/ipsock.go:313
	// _ = "end of CoverTab[6567]"
//line /usr/local/go/src/net/ipsock.go:313
	_go_fuzz_dep_.CoverTab[6568]++
						return IP{127, 0, 0, 1}
//line /usr/local/go/src/net/ipsock.go:314
	// _ = "end of CoverTab[6568]"
}

//line /usr/local/go/src/net/ipsock.go:315
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/ipsock.go:315
var _ = _go_fuzz_dep_.CoverTab
