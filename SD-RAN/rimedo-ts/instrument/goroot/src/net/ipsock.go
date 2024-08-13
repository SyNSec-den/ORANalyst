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
	_go_fuzz_dep_.CoverTab[14841]++
						ipStackCaps.Once.Do(ipStackCaps.probe)
						return ipStackCaps.ipv4Enabled
//line /usr/local/go/src/net/ipsock.go:34
	// _ = "end of CoverTab[14841]"
}

// supportsIPv6 reports whether the platform supports IPv6 networking
//line /usr/local/go/src/net/ipsock.go:37
// functionality.
//line /usr/local/go/src/net/ipsock.go:39
func supportsIPv6() bool {
//line /usr/local/go/src/net/ipsock.go:39
	_go_fuzz_dep_.CoverTab[14842]++
						ipStackCaps.Once.Do(ipStackCaps.probe)
						return ipStackCaps.ipv6Enabled
//line /usr/local/go/src/net/ipsock.go:41
	// _ = "end of CoverTab[14842]"
}

// supportsIPv4map reports whether the platform supports mapping an
//line /usr/local/go/src/net/ipsock.go:44
// IPv4 address inside an IPv6 address at transport layer
//line /usr/local/go/src/net/ipsock.go:44
// protocols. See RFC 4291, RFC 4038 and RFC 3493.
//line /usr/local/go/src/net/ipsock.go:47
func supportsIPv4map() bool {
//line /usr/local/go/src/net/ipsock.go:47
	_go_fuzz_dep_.CoverTab[14843]++

//line /usr/local/go/src/net/ipsock.go:50
	switch runtime.GOOS {
	case "dragonfly", "openbsd":
//line /usr/local/go/src/net/ipsock.go:51
		_go_fuzz_dep_.CoverTab[14845]++
							return false
//line /usr/local/go/src/net/ipsock.go:52
		// _ = "end of CoverTab[14845]"
//line /usr/local/go/src/net/ipsock.go:52
	default:
//line /usr/local/go/src/net/ipsock.go:52
		_go_fuzz_dep_.CoverTab[14846]++
//line /usr/local/go/src/net/ipsock.go:52
		// _ = "end of CoverTab[14846]"
	}
//line /usr/local/go/src/net/ipsock.go:53
	// _ = "end of CoverTab[14843]"
//line /usr/local/go/src/net/ipsock.go:53
	_go_fuzz_dep_.CoverTab[14844]++

						ipStackCaps.Once.Do(ipStackCaps.probe)
						return ipStackCaps.ipv4MappedIPv6Enabled
//line /usr/local/go/src/net/ipsock.go:56
	// _ = "end of CoverTab[14844]"
}

// An addrList represents a list of network endpoint addresses.
type addrList []Addr

// isIPv4 reports whether addr contains an IPv4 address.
func isIPv4(addr Addr) bool {
//line /usr/local/go/src/net/ipsock.go:63
	_go_fuzz_dep_.CoverTab[14847]++
						switch addr := addr.(type) {
	case *TCPAddr:
//line /usr/local/go/src/net/ipsock.go:65
		_go_fuzz_dep_.CoverTab[14849]++
							return addr.IP.To4() != nil
//line /usr/local/go/src/net/ipsock.go:66
		// _ = "end of CoverTab[14849]"
	case *UDPAddr:
//line /usr/local/go/src/net/ipsock.go:67
		_go_fuzz_dep_.CoverTab[14850]++
							return addr.IP.To4() != nil
//line /usr/local/go/src/net/ipsock.go:68
		// _ = "end of CoverTab[14850]"
	case *IPAddr:
//line /usr/local/go/src/net/ipsock.go:69
		_go_fuzz_dep_.CoverTab[14851]++
							return addr.IP.To4() != nil
//line /usr/local/go/src/net/ipsock.go:70
		// _ = "end of CoverTab[14851]"
	}
//line /usr/local/go/src/net/ipsock.go:71
	// _ = "end of CoverTab[14847]"
//line /usr/local/go/src/net/ipsock.go:71
	_go_fuzz_dep_.CoverTab[14848]++
						return false
//line /usr/local/go/src/net/ipsock.go:72
	// _ = "end of CoverTab[14848]"
}

// isNotIPv4 reports whether addr does not contain an IPv4 address.
func isNotIPv4(addr Addr) bool {
//line /usr/local/go/src/net/ipsock.go:76
	_go_fuzz_dep_.CoverTab[14852]++
//line /usr/local/go/src/net/ipsock.go:76
	return !isIPv4(addr)
//line /usr/local/go/src/net/ipsock.go:76
	// _ = "end of CoverTab[14852]"
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
	_go_fuzz_dep_.CoverTab[14853]++
						var want6 bool
						switch network {
	case "ip":
//line /usr/local/go/src/net/ipsock.go:84
		_go_fuzz_dep_.CoverTab[14856]++

							want6 = count(addr, ':') > 0
//line /usr/local/go/src/net/ipsock.go:86
		// _ = "end of CoverTab[14856]"
	case "tcp", "udp":
//line /usr/local/go/src/net/ipsock.go:87
		_go_fuzz_dep_.CoverTab[14857]++

							want6 = count(addr, '[') > 0
//line /usr/local/go/src/net/ipsock.go:89
		// _ = "end of CoverTab[14857]"
//line /usr/local/go/src/net/ipsock.go:89
	default:
//line /usr/local/go/src/net/ipsock.go:89
		_go_fuzz_dep_.CoverTab[14858]++
//line /usr/local/go/src/net/ipsock.go:89
		// _ = "end of CoverTab[14858]"
	}
//line /usr/local/go/src/net/ipsock.go:90
	// _ = "end of CoverTab[14853]"
//line /usr/local/go/src/net/ipsock.go:90
	_go_fuzz_dep_.CoverTab[14854]++
						if want6 {
//line /usr/local/go/src/net/ipsock.go:91
		_go_fuzz_dep_.CoverTab[14859]++
							return addrs.first(isNotIPv4)
//line /usr/local/go/src/net/ipsock.go:92
		// _ = "end of CoverTab[14859]"
	} else {
//line /usr/local/go/src/net/ipsock.go:93
		_go_fuzz_dep_.CoverTab[14860]++
//line /usr/local/go/src/net/ipsock.go:93
		// _ = "end of CoverTab[14860]"
//line /usr/local/go/src/net/ipsock.go:93
	}
//line /usr/local/go/src/net/ipsock.go:93
	// _ = "end of CoverTab[14854]"
//line /usr/local/go/src/net/ipsock.go:93
	_go_fuzz_dep_.CoverTab[14855]++
						return addrs.first(isIPv4)
//line /usr/local/go/src/net/ipsock.go:94
	// _ = "end of CoverTab[14855]"
}

// first returns the first address which satisfies strategy, or if
//line /usr/local/go/src/net/ipsock.go:97
// none do, then the first address of any kind.
//line /usr/local/go/src/net/ipsock.go:99
func (addrs addrList) first(strategy func(Addr) bool) Addr {
//line /usr/local/go/src/net/ipsock.go:99
	_go_fuzz_dep_.CoverTab[14861]++
						for _, addr := range addrs {
//line /usr/local/go/src/net/ipsock.go:100
		_go_fuzz_dep_.CoverTab[14863]++
							if strategy(addr) {
//line /usr/local/go/src/net/ipsock.go:101
			_go_fuzz_dep_.CoverTab[14864]++
								return addr
//line /usr/local/go/src/net/ipsock.go:102
			// _ = "end of CoverTab[14864]"
		} else {
//line /usr/local/go/src/net/ipsock.go:103
			_go_fuzz_dep_.CoverTab[14865]++
//line /usr/local/go/src/net/ipsock.go:103
			// _ = "end of CoverTab[14865]"
//line /usr/local/go/src/net/ipsock.go:103
		}
//line /usr/local/go/src/net/ipsock.go:103
		// _ = "end of CoverTab[14863]"
	}
//line /usr/local/go/src/net/ipsock.go:104
	// _ = "end of CoverTab[14861]"
//line /usr/local/go/src/net/ipsock.go:104
	_go_fuzz_dep_.CoverTab[14862]++
						return addrs[0]
//line /usr/local/go/src/net/ipsock.go:105
	// _ = "end of CoverTab[14862]"
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
	_go_fuzz_dep_.CoverTab[14866]++
						var primaryLabel bool
						for i, addr := range addrs {
//line /usr/local/go/src/net/ipsock.go:116
		_go_fuzz_dep_.CoverTab[14868]++
							label := strategy(addr)
							if i == 0 || func() bool {
//line /usr/local/go/src/net/ipsock.go:118
			_go_fuzz_dep_.CoverTab[14869]++
//line /usr/local/go/src/net/ipsock.go:118
			return label == primaryLabel
//line /usr/local/go/src/net/ipsock.go:118
			// _ = "end of CoverTab[14869]"
//line /usr/local/go/src/net/ipsock.go:118
		}() {
//line /usr/local/go/src/net/ipsock.go:118
			_go_fuzz_dep_.CoverTab[14870]++
								primaryLabel = label
								primaries = append(primaries, addr)
//line /usr/local/go/src/net/ipsock.go:120
			// _ = "end of CoverTab[14870]"
		} else {
//line /usr/local/go/src/net/ipsock.go:121
			_go_fuzz_dep_.CoverTab[14871]++
								fallbacks = append(fallbacks, addr)
//line /usr/local/go/src/net/ipsock.go:122
			// _ = "end of CoverTab[14871]"
		}
//line /usr/local/go/src/net/ipsock.go:123
		// _ = "end of CoverTab[14868]"
	}
//line /usr/local/go/src/net/ipsock.go:124
	// _ = "end of CoverTab[14866]"
//line /usr/local/go/src/net/ipsock.go:124
	_go_fuzz_dep_.CoverTab[14867]++
						return
//line /usr/local/go/src/net/ipsock.go:125
	// _ = "end of CoverTab[14867]"
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
	_go_fuzz_dep_.CoverTab[14872]++
						var addrs addrList
						for _, ip := range ips {
//line /usr/local/go/src/net/ipsock.go:134
		_go_fuzz_dep_.CoverTab[14875]++
							if filter == nil || func() bool {
//line /usr/local/go/src/net/ipsock.go:135
			_go_fuzz_dep_.CoverTab[14876]++
//line /usr/local/go/src/net/ipsock.go:135
			return filter(ip)
//line /usr/local/go/src/net/ipsock.go:135
			// _ = "end of CoverTab[14876]"
//line /usr/local/go/src/net/ipsock.go:135
		}() {
//line /usr/local/go/src/net/ipsock.go:135
			_go_fuzz_dep_.CoverTab[14877]++
								addrs = append(addrs, inetaddr(ip))
//line /usr/local/go/src/net/ipsock.go:136
			// _ = "end of CoverTab[14877]"
		} else {
//line /usr/local/go/src/net/ipsock.go:137
			_go_fuzz_dep_.CoverTab[14878]++
//line /usr/local/go/src/net/ipsock.go:137
			// _ = "end of CoverTab[14878]"
//line /usr/local/go/src/net/ipsock.go:137
		}
//line /usr/local/go/src/net/ipsock.go:137
		// _ = "end of CoverTab[14875]"
	}
//line /usr/local/go/src/net/ipsock.go:138
	// _ = "end of CoverTab[14872]"
//line /usr/local/go/src/net/ipsock.go:138
	_go_fuzz_dep_.CoverTab[14873]++
						if len(addrs) == 0 {
//line /usr/local/go/src/net/ipsock.go:139
		_go_fuzz_dep_.CoverTab[14879]++
							return nil, &AddrError{Err: errNoSuitableAddress.Error(), Addr: originalAddr}
//line /usr/local/go/src/net/ipsock.go:140
		// _ = "end of CoverTab[14879]"
	} else {
//line /usr/local/go/src/net/ipsock.go:141
		_go_fuzz_dep_.CoverTab[14880]++
//line /usr/local/go/src/net/ipsock.go:141
		// _ = "end of CoverTab[14880]"
//line /usr/local/go/src/net/ipsock.go:141
	}
//line /usr/local/go/src/net/ipsock.go:141
	// _ = "end of CoverTab[14873]"
//line /usr/local/go/src/net/ipsock.go:141
	_go_fuzz_dep_.CoverTab[14874]++
						return addrs, nil
//line /usr/local/go/src/net/ipsock.go:142
	// _ = "end of CoverTab[14874]"
}

// ipv4only reports whether addr is an IPv4 address.
func ipv4only(addr IPAddr) bool {
//line /usr/local/go/src/net/ipsock.go:146
	_go_fuzz_dep_.CoverTab[14881]++
						return addr.IP.To4() != nil
//line /usr/local/go/src/net/ipsock.go:147
	// _ = "end of CoverTab[14881]"
}

// ipv6only reports whether addr is an IPv6 address except IPv4-mapped IPv6 address.
func ipv6only(addr IPAddr) bool {
//line /usr/local/go/src/net/ipsock.go:151
	_go_fuzz_dep_.CoverTab[14882]++
						return len(addr.IP) == IPv6len && func() bool {
//line /usr/local/go/src/net/ipsock.go:152
		_go_fuzz_dep_.CoverTab[14883]++
//line /usr/local/go/src/net/ipsock.go:152
		return addr.IP.To4() == nil
//line /usr/local/go/src/net/ipsock.go:152
		// _ = "end of CoverTab[14883]"
//line /usr/local/go/src/net/ipsock.go:152
	}()
//line /usr/local/go/src/net/ipsock.go:152
	// _ = "end of CoverTab[14882]"
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
	_go_fuzz_dep_.CoverTab[14884]++
						const (
		missingPort	= "missing port in address"
		tooManyColons	= "too many colons in address"
	)
	addrErr := func(addr, why string) (host, port string, err error) {
//line /usr/local/go/src/net/ipsock.go:169
		_go_fuzz_dep_.CoverTab[14890]++
							return "", "", &AddrError{Err: why, Addr: addr}
//line /usr/local/go/src/net/ipsock.go:170
		// _ = "end of CoverTab[14890]"
	}
//line /usr/local/go/src/net/ipsock.go:171
	// _ = "end of CoverTab[14884]"
//line /usr/local/go/src/net/ipsock.go:171
	_go_fuzz_dep_.CoverTab[14885]++
						j, k := 0, 0

//line /usr/local/go/src/net/ipsock.go:175
	i := last(hostport, ':')
	if i < 0 {
//line /usr/local/go/src/net/ipsock.go:176
		_go_fuzz_dep_.CoverTab[14891]++
							return addrErr(hostport, missingPort)
//line /usr/local/go/src/net/ipsock.go:177
		// _ = "end of CoverTab[14891]"
	} else {
//line /usr/local/go/src/net/ipsock.go:178
		_go_fuzz_dep_.CoverTab[14892]++
//line /usr/local/go/src/net/ipsock.go:178
		// _ = "end of CoverTab[14892]"
//line /usr/local/go/src/net/ipsock.go:178
	}
//line /usr/local/go/src/net/ipsock.go:178
	// _ = "end of CoverTab[14885]"
//line /usr/local/go/src/net/ipsock.go:178
	_go_fuzz_dep_.CoverTab[14886]++

						if hostport[0] == '[' {
//line /usr/local/go/src/net/ipsock.go:180
		_go_fuzz_dep_.CoverTab[14893]++

							end := bytealg.IndexByteString(hostport, ']')
							if end < 0 {
//line /usr/local/go/src/net/ipsock.go:183
			_go_fuzz_dep_.CoverTab[14896]++
								return addrErr(hostport, "missing ']' in address")
//line /usr/local/go/src/net/ipsock.go:184
			// _ = "end of CoverTab[14896]"
		} else {
//line /usr/local/go/src/net/ipsock.go:185
			_go_fuzz_dep_.CoverTab[14897]++
//line /usr/local/go/src/net/ipsock.go:185
			// _ = "end of CoverTab[14897]"
//line /usr/local/go/src/net/ipsock.go:185
		}
//line /usr/local/go/src/net/ipsock.go:185
		// _ = "end of CoverTab[14893]"
//line /usr/local/go/src/net/ipsock.go:185
		_go_fuzz_dep_.CoverTab[14894]++
							switch end + 1 {
		case len(hostport):
//line /usr/local/go/src/net/ipsock.go:187
			_go_fuzz_dep_.CoverTab[14898]++

								return addrErr(hostport, missingPort)
//line /usr/local/go/src/net/ipsock.go:189
			// _ = "end of CoverTab[14898]"
		case i:
//line /usr/local/go/src/net/ipsock.go:190
			_go_fuzz_dep_.CoverTab[14899]++
//line /usr/local/go/src/net/ipsock.go:190
			// _ = "end of CoverTab[14899]"

		default:
//line /usr/local/go/src/net/ipsock.go:192
			_go_fuzz_dep_.CoverTab[14900]++

//line /usr/local/go/src/net/ipsock.go:195
			if hostport[end+1] == ':' {
//line /usr/local/go/src/net/ipsock.go:195
				_go_fuzz_dep_.CoverTab[14902]++
									return addrErr(hostport, tooManyColons)
//line /usr/local/go/src/net/ipsock.go:196
				// _ = "end of CoverTab[14902]"
			} else {
//line /usr/local/go/src/net/ipsock.go:197
				_go_fuzz_dep_.CoverTab[14903]++
//line /usr/local/go/src/net/ipsock.go:197
				// _ = "end of CoverTab[14903]"
//line /usr/local/go/src/net/ipsock.go:197
			}
//line /usr/local/go/src/net/ipsock.go:197
			// _ = "end of CoverTab[14900]"
//line /usr/local/go/src/net/ipsock.go:197
			_go_fuzz_dep_.CoverTab[14901]++
								return addrErr(hostport, missingPort)
//line /usr/local/go/src/net/ipsock.go:198
			// _ = "end of CoverTab[14901]"
		}
//line /usr/local/go/src/net/ipsock.go:199
		// _ = "end of CoverTab[14894]"
//line /usr/local/go/src/net/ipsock.go:199
		_go_fuzz_dep_.CoverTab[14895]++
							host = hostport[1:end]
							j, k = 1, end+1
//line /usr/local/go/src/net/ipsock.go:201
		// _ = "end of CoverTab[14895]"
	} else {
//line /usr/local/go/src/net/ipsock.go:202
		_go_fuzz_dep_.CoverTab[14904]++
							host = hostport[:i]
							if bytealg.IndexByteString(host, ':') >= 0 {
//line /usr/local/go/src/net/ipsock.go:204
			_go_fuzz_dep_.CoverTab[14905]++
								return addrErr(hostport, tooManyColons)
//line /usr/local/go/src/net/ipsock.go:205
			// _ = "end of CoverTab[14905]"
		} else {
//line /usr/local/go/src/net/ipsock.go:206
			_go_fuzz_dep_.CoverTab[14906]++
//line /usr/local/go/src/net/ipsock.go:206
			// _ = "end of CoverTab[14906]"
//line /usr/local/go/src/net/ipsock.go:206
		}
//line /usr/local/go/src/net/ipsock.go:206
		// _ = "end of CoverTab[14904]"
	}
//line /usr/local/go/src/net/ipsock.go:207
	// _ = "end of CoverTab[14886]"
//line /usr/local/go/src/net/ipsock.go:207
	_go_fuzz_dep_.CoverTab[14887]++
						if bytealg.IndexByteString(hostport[j:], '[') >= 0 {
//line /usr/local/go/src/net/ipsock.go:208
		_go_fuzz_dep_.CoverTab[14907]++
							return addrErr(hostport, "unexpected '[' in address")
//line /usr/local/go/src/net/ipsock.go:209
		// _ = "end of CoverTab[14907]"
	} else {
//line /usr/local/go/src/net/ipsock.go:210
		_go_fuzz_dep_.CoverTab[14908]++
//line /usr/local/go/src/net/ipsock.go:210
		// _ = "end of CoverTab[14908]"
//line /usr/local/go/src/net/ipsock.go:210
	}
//line /usr/local/go/src/net/ipsock.go:210
	// _ = "end of CoverTab[14887]"
//line /usr/local/go/src/net/ipsock.go:210
	_go_fuzz_dep_.CoverTab[14888]++
						if bytealg.IndexByteString(hostport[k:], ']') >= 0 {
//line /usr/local/go/src/net/ipsock.go:211
		_go_fuzz_dep_.CoverTab[14909]++
							return addrErr(hostport, "unexpected ']' in address")
//line /usr/local/go/src/net/ipsock.go:212
		// _ = "end of CoverTab[14909]"
	} else {
//line /usr/local/go/src/net/ipsock.go:213
		_go_fuzz_dep_.CoverTab[14910]++
//line /usr/local/go/src/net/ipsock.go:213
		// _ = "end of CoverTab[14910]"
//line /usr/local/go/src/net/ipsock.go:213
	}
//line /usr/local/go/src/net/ipsock.go:213
	// _ = "end of CoverTab[14888]"
//line /usr/local/go/src/net/ipsock.go:213
	_go_fuzz_dep_.CoverTab[14889]++

						port = hostport[i+1:]
						return host, port, nil
//line /usr/local/go/src/net/ipsock.go:216
	// _ = "end of CoverTab[14889]"
}

func splitHostZone(s string) (host, zone string) {
//line /usr/local/go/src/net/ipsock.go:219
	_go_fuzz_dep_.CoverTab[14911]++

//line /usr/local/go/src/net/ipsock.go:222
	if i := last(s, '%'); i > 0 {
//line /usr/local/go/src/net/ipsock.go:222
		_go_fuzz_dep_.CoverTab[14913]++
							host, zone = s[:i], s[i+1:]
//line /usr/local/go/src/net/ipsock.go:223
		// _ = "end of CoverTab[14913]"
	} else {
//line /usr/local/go/src/net/ipsock.go:224
		_go_fuzz_dep_.CoverTab[14914]++
							host = s
//line /usr/local/go/src/net/ipsock.go:225
		// _ = "end of CoverTab[14914]"
	}
//line /usr/local/go/src/net/ipsock.go:226
	// _ = "end of CoverTab[14911]"
//line /usr/local/go/src/net/ipsock.go:226
	_go_fuzz_dep_.CoverTab[14912]++
						return
//line /usr/local/go/src/net/ipsock.go:227
	// _ = "end of CoverTab[14912]"
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
	_go_fuzz_dep_.CoverTab[14915]++

//line /usr/local/go/src/net/ipsock.go:238
	if bytealg.IndexByteString(host, ':') >= 0 {
//line /usr/local/go/src/net/ipsock.go:238
		_go_fuzz_dep_.CoverTab[14917]++
							return "[" + host + "]:" + port
//line /usr/local/go/src/net/ipsock.go:239
		// _ = "end of CoverTab[14917]"
	} else {
//line /usr/local/go/src/net/ipsock.go:240
		_go_fuzz_dep_.CoverTab[14918]++
//line /usr/local/go/src/net/ipsock.go:240
		// _ = "end of CoverTab[14918]"
//line /usr/local/go/src/net/ipsock.go:240
	}
//line /usr/local/go/src/net/ipsock.go:240
	// _ = "end of CoverTab[14915]"
//line /usr/local/go/src/net/ipsock.go:240
	_go_fuzz_dep_.CoverTab[14916]++
						return host + ":" + port
//line /usr/local/go/src/net/ipsock.go:241
	// _ = "end of CoverTab[14916]"
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
	_go_fuzz_dep_.CoverTab[14919]++
						var (
		err		error
		host, port	string
		portnum		int
	)
	switch net {
	case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
//line /usr/local/go/src/net/ipsock.go:255
		_go_fuzz_dep_.CoverTab[14927]++
							if addr != "" {
//line /usr/local/go/src/net/ipsock.go:256
			_go_fuzz_dep_.CoverTab[14930]++
								if host, port, err = SplitHostPort(addr); err != nil {
//line /usr/local/go/src/net/ipsock.go:257
				_go_fuzz_dep_.CoverTab[14932]++
									return nil, err
//line /usr/local/go/src/net/ipsock.go:258
				// _ = "end of CoverTab[14932]"
			} else {
//line /usr/local/go/src/net/ipsock.go:259
				_go_fuzz_dep_.CoverTab[14933]++
//line /usr/local/go/src/net/ipsock.go:259
				// _ = "end of CoverTab[14933]"
//line /usr/local/go/src/net/ipsock.go:259
			}
//line /usr/local/go/src/net/ipsock.go:259
			// _ = "end of CoverTab[14930]"
//line /usr/local/go/src/net/ipsock.go:259
			_go_fuzz_dep_.CoverTab[14931]++
								if portnum, err = r.LookupPort(ctx, net, port); err != nil {
//line /usr/local/go/src/net/ipsock.go:260
				_go_fuzz_dep_.CoverTab[14934]++
									return nil, err
//line /usr/local/go/src/net/ipsock.go:261
				// _ = "end of CoverTab[14934]"
			} else {
//line /usr/local/go/src/net/ipsock.go:262
				_go_fuzz_dep_.CoverTab[14935]++
//line /usr/local/go/src/net/ipsock.go:262
				// _ = "end of CoverTab[14935]"
//line /usr/local/go/src/net/ipsock.go:262
			}
//line /usr/local/go/src/net/ipsock.go:262
			// _ = "end of CoverTab[14931]"
		} else {
//line /usr/local/go/src/net/ipsock.go:263
			_go_fuzz_dep_.CoverTab[14936]++
//line /usr/local/go/src/net/ipsock.go:263
			// _ = "end of CoverTab[14936]"
//line /usr/local/go/src/net/ipsock.go:263
		}
//line /usr/local/go/src/net/ipsock.go:263
		// _ = "end of CoverTab[14927]"
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/ipsock.go:264
		_go_fuzz_dep_.CoverTab[14928]++
							if addr != "" {
//line /usr/local/go/src/net/ipsock.go:265
			_go_fuzz_dep_.CoverTab[14937]++
								host = addr
//line /usr/local/go/src/net/ipsock.go:266
			// _ = "end of CoverTab[14937]"
		} else {
//line /usr/local/go/src/net/ipsock.go:267
			_go_fuzz_dep_.CoverTab[14938]++
//line /usr/local/go/src/net/ipsock.go:267
			// _ = "end of CoverTab[14938]"
//line /usr/local/go/src/net/ipsock.go:267
		}
//line /usr/local/go/src/net/ipsock.go:267
		// _ = "end of CoverTab[14928]"
	default:
//line /usr/local/go/src/net/ipsock.go:268
		_go_fuzz_dep_.CoverTab[14929]++
							return nil, UnknownNetworkError(net)
//line /usr/local/go/src/net/ipsock.go:269
		// _ = "end of CoverTab[14929]"
	}
//line /usr/local/go/src/net/ipsock.go:270
	// _ = "end of CoverTab[14919]"
//line /usr/local/go/src/net/ipsock.go:270
	_go_fuzz_dep_.CoverTab[14920]++
						inetaddr := func(ip IPAddr) Addr {
//line /usr/local/go/src/net/ipsock.go:271
		_go_fuzz_dep_.CoverTab[14939]++
							switch net {
		case "tcp", "tcp4", "tcp6":
//line /usr/local/go/src/net/ipsock.go:273
			_go_fuzz_dep_.CoverTab[14940]++
								return &TCPAddr{IP: ip.IP, Port: portnum, Zone: ip.Zone}
//line /usr/local/go/src/net/ipsock.go:274
			// _ = "end of CoverTab[14940]"
		case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/ipsock.go:275
			_go_fuzz_dep_.CoverTab[14941]++
								return &UDPAddr{IP: ip.IP, Port: portnum, Zone: ip.Zone}
//line /usr/local/go/src/net/ipsock.go:276
			// _ = "end of CoverTab[14941]"
		case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/ipsock.go:277
			_go_fuzz_dep_.CoverTab[14942]++
								return &IPAddr{IP: ip.IP, Zone: ip.Zone}
//line /usr/local/go/src/net/ipsock.go:278
			// _ = "end of CoverTab[14942]"
		default:
//line /usr/local/go/src/net/ipsock.go:279
			_go_fuzz_dep_.CoverTab[14943]++
								panic("unexpected network: " + net)
//line /usr/local/go/src/net/ipsock.go:280
			// _ = "end of CoverTab[14943]"
		}
//line /usr/local/go/src/net/ipsock.go:281
		// _ = "end of CoverTab[14939]"
	}
//line /usr/local/go/src/net/ipsock.go:282
	// _ = "end of CoverTab[14920]"
//line /usr/local/go/src/net/ipsock.go:282
	_go_fuzz_dep_.CoverTab[14921]++
						if host == "" {
//line /usr/local/go/src/net/ipsock.go:283
		_go_fuzz_dep_.CoverTab[14944]++
							return addrList{inetaddr(IPAddr{})}, nil
//line /usr/local/go/src/net/ipsock.go:284
		// _ = "end of CoverTab[14944]"
	} else {
//line /usr/local/go/src/net/ipsock.go:285
		_go_fuzz_dep_.CoverTab[14945]++
//line /usr/local/go/src/net/ipsock.go:285
		// _ = "end of CoverTab[14945]"
//line /usr/local/go/src/net/ipsock.go:285
	}
//line /usr/local/go/src/net/ipsock.go:285
	// _ = "end of CoverTab[14921]"
//line /usr/local/go/src/net/ipsock.go:285
	_go_fuzz_dep_.CoverTab[14922]++

//line /usr/local/go/src/net/ipsock.go:288
	ips, err := r.lookupIPAddr(ctx, net, host)
	if err != nil {
//line /usr/local/go/src/net/ipsock.go:289
		_go_fuzz_dep_.CoverTab[14946]++
							return nil, err
//line /usr/local/go/src/net/ipsock.go:290
		// _ = "end of CoverTab[14946]"
	} else {
//line /usr/local/go/src/net/ipsock.go:291
		_go_fuzz_dep_.CoverTab[14947]++
//line /usr/local/go/src/net/ipsock.go:291
		// _ = "end of CoverTab[14947]"
//line /usr/local/go/src/net/ipsock.go:291
	}
//line /usr/local/go/src/net/ipsock.go:291
	// _ = "end of CoverTab[14922]"
//line /usr/local/go/src/net/ipsock.go:291
	_go_fuzz_dep_.CoverTab[14923]++

//line /usr/local/go/src/net/ipsock.go:296
	if len(ips) == 1 && func() bool {
//line /usr/local/go/src/net/ipsock.go:296
		_go_fuzz_dep_.CoverTab[14948]++
//line /usr/local/go/src/net/ipsock.go:296
		return ips[0].IP.Equal(IPv6unspecified)
//line /usr/local/go/src/net/ipsock.go:296
		// _ = "end of CoverTab[14948]"
//line /usr/local/go/src/net/ipsock.go:296
	}() {
//line /usr/local/go/src/net/ipsock.go:296
		_go_fuzz_dep_.CoverTab[14949]++
							ips = append(ips, IPAddr{IP: IPv4zero})
//line /usr/local/go/src/net/ipsock.go:297
		// _ = "end of CoverTab[14949]"
	} else {
//line /usr/local/go/src/net/ipsock.go:298
		_go_fuzz_dep_.CoverTab[14950]++
//line /usr/local/go/src/net/ipsock.go:298
		// _ = "end of CoverTab[14950]"
//line /usr/local/go/src/net/ipsock.go:298
	}
//line /usr/local/go/src/net/ipsock.go:298
	// _ = "end of CoverTab[14923]"
//line /usr/local/go/src/net/ipsock.go:298
	_go_fuzz_dep_.CoverTab[14924]++

						var filter func(IPAddr) bool
						if net != "" && func() bool {
//line /usr/local/go/src/net/ipsock.go:301
		_go_fuzz_dep_.CoverTab[14951]++
//line /usr/local/go/src/net/ipsock.go:301
		return net[len(net)-1] == '4'
//line /usr/local/go/src/net/ipsock.go:301
		// _ = "end of CoverTab[14951]"
//line /usr/local/go/src/net/ipsock.go:301
	}() {
//line /usr/local/go/src/net/ipsock.go:301
		_go_fuzz_dep_.CoverTab[14952]++
							filter = ipv4only
//line /usr/local/go/src/net/ipsock.go:302
		// _ = "end of CoverTab[14952]"
	} else {
//line /usr/local/go/src/net/ipsock.go:303
		_go_fuzz_dep_.CoverTab[14953]++
//line /usr/local/go/src/net/ipsock.go:303
		// _ = "end of CoverTab[14953]"
//line /usr/local/go/src/net/ipsock.go:303
	}
//line /usr/local/go/src/net/ipsock.go:303
	// _ = "end of CoverTab[14924]"
//line /usr/local/go/src/net/ipsock.go:303
	_go_fuzz_dep_.CoverTab[14925]++
						if net != "" && func() bool {
//line /usr/local/go/src/net/ipsock.go:304
		_go_fuzz_dep_.CoverTab[14954]++
//line /usr/local/go/src/net/ipsock.go:304
		return net[len(net)-1] == '6'
//line /usr/local/go/src/net/ipsock.go:304
		// _ = "end of CoverTab[14954]"
//line /usr/local/go/src/net/ipsock.go:304
	}() {
//line /usr/local/go/src/net/ipsock.go:304
		_go_fuzz_dep_.CoverTab[14955]++
							filter = ipv6only
//line /usr/local/go/src/net/ipsock.go:305
		// _ = "end of CoverTab[14955]"
	} else {
//line /usr/local/go/src/net/ipsock.go:306
		_go_fuzz_dep_.CoverTab[14956]++
//line /usr/local/go/src/net/ipsock.go:306
		// _ = "end of CoverTab[14956]"
//line /usr/local/go/src/net/ipsock.go:306
	}
//line /usr/local/go/src/net/ipsock.go:306
	// _ = "end of CoverTab[14925]"
//line /usr/local/go/src/net/ipsock.go:306
	_go_fuzz_dep_.CoverTab[14926]++
						return filterAddrList(filter, ips, inetaddr, host)
//line /usr/local/go/src/net/ipsock.go:307
	// _ = "end of CoverTab[14926]"
}

func loopbackIP(net string) IP {
//line /usr/local/go/src/net/ipsock.go:310
	_go_fuzz_dep_.CoverTab[14957]++
						if net != "" && func() bool {
//line /usr/local/go/src/net/ipsock.go:311
		_go_fuzz_dep_.CoverTab[14959]++
//line /usr/local/go/src/net/ipsock.go:311
		return net[len(net)-1] == '6'
//line /usr/local/go/src/net/ipsock.go:311
		// _ = "end of CoverTab[14959]"
//line /usr/local/go/src/net/ipsock.go:311
	}() {
//line /usr/local/go/src/net/ipsock.go:311
		_go_fuzz_dep_.CoverTab[14960]++
							return IPv6loopback
//line /usr/local/go/src/net/ipsock.go:312
		// _ = "end of CoverTab[14960]"
	} else {
//line /usr/local/go/src/net/ipsock.go:313
		_go_fuzz_dep_.CoverTab[14961]++
//line /usr/local/go/src/net/ipsock.go:313
		// _ = "end of CoverTab[14961]"
//line /usr/local/go/src/net/ipsock.go:313
	}
//line /usr/local/go/src/net/ipsock.go:313
	// _ = "end of CoverTab[14957]"
//line /usr/local/go/src/net/ipsock.go:313
	_go_fuzz_dep_.CoverTab[14958]++
						return IP{127, 0, 0, 1}
//line /usr/local/go/src/net/ipsock.go:314
	// _ = "end of CoverTab[14958]"
}

//line /usr/local/go/src/net/ipsock.go:315
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/ipsock.go:315
var _ = _go_fuzz_dep_.CoverTab
