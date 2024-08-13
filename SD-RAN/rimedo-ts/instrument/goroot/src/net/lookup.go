// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/lookup.go:5
package net

//line /usr/local/go/src/net/lookup.go:5
import (
//line /usr/local/go/src/net/lookup.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/lookup.go:5
)
//line /usr/local/go/src/net/lookup.go:5
import (
//line /usr/local/go/src/net/lookup.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/lookup.go:5
)

import (
	"context"
	"errors"
	"internal/nettrace"
	"internal/singleflight"
	"net/netip"
	"sync"

	"golang.org/x/net/dns/dnsmessage"
)

// protocols contains minimal mappings between internet protocol
//line /usr/local/go/src/net/lookup.go:18
// names and numbers for platforms that don't have a complete list of
//line /usr/local/go/src/net/lookup.go:18
// protocol numbers.
//line /usr/local/go/src/net/lookup.go:18
//
//line /usr/local/go/src/net/lookup.go:18
// See https://www.iana.org/assignments/protocol-numbers
//line /usr/local/go/src/net/lookup.go:18
//
//line /usr/local/go/src/net/lookup.go:18
// On Unix, this map is augmented by readProtocols via lookupProtocol.
//line /usr/local/go/src/net/lookup.go:25
var protocols = map[string]int{
	"icmp":		1,
	"igmp":		2,
	"tcp":		6,
	"udp":		17,
	"ipv6-icmp":	58,
}

// services contains minimal mappings between services names and port
//line /usr/local/go/src/net/lookup.go:33
// numbers for platforms that don't have a complete list of port numbers.
//line /usr/local/go/src/net/lookup.go:33
//
//line /usr/local/go/src/net/lookup.go:33
// See https://www.iana.org/assignments/service-names-port-numbers
//line /usr/local/go/src/net/lookup.go:33
//
//line /usr/local/go/src/net/lookup.go:33
// On Unix, this map is augmented by readServices via goLookupPort.
//line /usr/local/go/src/net/lookup.go:39
var services = map[string]map[string]int{
	"udp": {
		"domain": 53,
	},
	"tcp": {
		"ftp":		21,
		"ftps":		990,
		"gopher":	70,
		"http":		80,
		"https":	443,
		"imap2":	143,
		"imap3":	220,
		"imaps":	993,
		"pop3":		110,
		"pop3s":	995,
		"smtp":		25,
		"ssh":		22,
		"telnet":	23,
	},
}

// dnsWaitGroup can be used by tests to wait for all DNS goroutines to
//line /usr/local/go/src/net/lookup.go:60
// complete. This avoids races on the test hooks.
//line /usr/local/go/src/net/lookup.go:62
var dnsWaitGroup sync.WaitGroup

const maxProtoLength = len("RSVP-E2E-IGNORE") + 10	// with room to grow

func lookupProtocolMap(name string) (int, error) {
//line /usr/local/go/src/net/lookup.go:66
	_go_fuzz_dep_.CoverTab[15047]++
						var lowerProtocol [maxProtoLength]byte
						n := copy(lowerProtocol[:], name)
						lowerASCIIBytes(lowerProtocol[:n])
						proto, found := protocols[string(lowerProtocol[:n])]
						if !found || func() bool {
//line /usr/local/go/src/net/lookup.go:71
		_go_fuzz_dep_.CoverTab[15049]++
//line /usr/local/go/src/net/lookup.go:71
		return n != len(name)
//line /usr/local/go/src/net/lookup.go:71
		// _ = "end of CoverTab[15049]"
//line /usr/local/go/src/net/lookup.go:71
	}() {
//line /usr/local/go/src/net/lookup.go:71
		_go_fuzz_dep_.CoverTab[15050]++
							return 0, &AddrError{Err: "unknown IP protocol specified", Addr: name}
//line /usr/local/go/src/net/lookup.go:72
		// _ = "end of CoverTab[15050]"
	} else {
//line /usr/local/go/src/net/lookup.go:73
		_go_fuzz_dep_.CoverTab[15051]++
//line /usr/local/go/src/net/lookup.go:73
		// _ = "end of CoverTab[15051]"
//line /usr/local/go/src/net/lookup.go:73
	}
//line /usr/local/go/src/net/lookup.go:73
	// _ = "end of CoverTab[15047]"
//line /usr/local/go/src/net/lookup.go:73
	_go_fuzz_dep_.CoverTab[15048]++
						return proto, nil
//line /usr/local/go/src/net/lookup.go:74
	// _ = "end of CoverTab[15048]"
}

// maxPortBufSize is the longest reasonable name of a service
//line /usr/local/go/src/net/lookup.go:77
// (non-numeric port).
//line /usr/local/go/src/net/lookup.go:77
// Currently the longest known IANA-unregistered name is
//line /usr/local/go/src/net/lookup.go:77
// "mobility-header", so we use that length, plus some slop in case
//line /usr/local/go/src/net/lookup.go:77
// something longer is added in the future.
//line /usr/local/go/src/net/lookup.go:82
const maxPortBufSize = len("mobility-header") + 10

func lookupPortMap(network, service string) (port int, error error) {
//line /usr/local/go/src/net/lookup.go:84
	_go_fuzz_dep_.CoverTab[15052]++
						switch network {
	case "tcp4", "tcp6":
//line /usr/local/go/src/net/lookup.go:86
		_go_fuzz_dep_.CoverTab[15055]++
							network = "tcp"
//line /usr/local/go/src/net/lookup.go:87
		// _ = "end of CoverTab[15055]"
	case "udp4", "udp6":
//line /usr/local/go/src/net/lookup.go:88
		_go_fuzz_dep_.CoverTab[15056]++
							network = "udp"
//line /usr/local/go/src/net/lookup.go:89
		// _ = "end of CoverTab[15056]"
//line /usr/local/go/src/net/lookup.go:89
	default:
//line /usr/local/go/src/net/lookup.go:89
		_go_fuzz_dep_.CoverTab[15057]++
//line /usr/local/go/src/net/lookup.go:89
		// _ = "end of CoverTab[15057]"
	}
//line /usr/local/go/src/net/lookup.go:90
	// _ = "end of CoverTab[15052]"
//line /usr/local/go/src/net/lookup.go:90
	_go_fuzz_dep_.CoverTab[15053]++

						if m, ok := services[network]; ok {
//line /usr/local/go/src/net/lookup.go:92
		_go_fuzz_dep_.CoverTab[15058]++
							var lowerService [maxPortBufSize]byte
							n := copy(lowerService[:], service)
							lowerASCIIBytes(lowerService[:n])
							if port, ok := m[string(lowerService[:n])]; ok && func() bool {
//line /usr/local/go/src/net/lookup.go:96
			_go_fuzz_dep_.CoverTab[15059]++
//line /usr/local/go/src/net/lookup.go:96
			return n == len(service)
//line /usr/local/go/src/net/lookup.go:96
			// _ = "end of CoverTab[15059]"
//line /usr/local/go/src/net/lookup.go:96
		}() {
//line /usr/local/go/src/net/lookup.go:96
			_go_fuzz_dep_.CoverTab[15060]++
								return port, nil
//line /usr/local/go/src/net/lookup.go:97
			// _ = "end of CoverTab[15060]"
		} else {
//line /usr/local/go/src/net/lookup.go:98
			_go_fuzz_dep_.CoverTab[15061]++
//line /usr/local/go/src/net/lookup.go:98
			// _ = "end of CoverTab[15061]"
//line /usr/local/go/src/net/lookup.go:98
		}
//line /usr/local/go/src/net/lookup.go:98
		// _ = "end of CoverTab[15058]"
	} else {
//line /usr/local/go/src/net/lookup.go:99
		_go_fuzz_dep_.CoverTab[15062]++
//line /usr/local/go/src/net/lookup.go:99
		// _ = "end of CoverTab[15062]"
//line /usr/local/go/src/net/lookup.go:99
	}
//line /usr/local/go/src/net/lookup.go:99
	// _ = "end of CoverTab[15053]"
//line /usr/local/go/src/net/lookup.go:99
	_go_fuzz_dep_.CoverTab[15054]++
						return 0, &AddrError{Err: "unknown port", Addr: network + "/" + service}
//line /usr/local/go/src/net/lookup.go:100
	// _ = "end of CoverTab[15054]"
}

// ipVersion returns the provided network's IP version: '4', '6' or 0
//line /usr/local/go/src/net/lookup.go:103
// if network does not end in a '4' or '6' byte.
//line /usr/local/go/src/net/lookup.go:105
func ipVersion(network string) byte {
//line /usr/local/go/src/net/lookup.go:105
	_go_fuzz_dep_.CoverTab[15063]++
						if network == "" {
//line /usr/local/go/src/net/lookup.go:106
		_go_fuzz_dep_.CoverTab[15066]++
							return 0
//line /usr/local/go/src/net/lookup.go:107
		// _ = "end of CoverTab[15066]"
	} else {
//line /usr/local/go/src/net/lookup.go:108
		_go_fuzz_dep_.CoverTab[15067]++
//line /usr/local/go/src/net/lookup.go:108
		// _ = "end of CoverTab[15067]"
//line /usr/local/go/src/net/lookup.go:108
	}
//line /usr/local/go/src/net/lookup.go:108
	// _ = "end of CoverTab[15063]"
//line /usr/local/go/src/net/lookup.go:108
	_go_fuzz_dep_.CoverTab[15064]++
						n := network[len(network)-1]
						if n != '4' && func() bool {
//line /usr/local/go/src/net/lookup.go:110
		_go_fuzz_dep_.CoverTab[15068]++
//line /usr/local/go/src/net/lookup.go:110
		return n != '6'
//line /usr/local/go/src/net/lookup.go:110
		// _ = "end of CoverTab[15068]"
//line /usr/local/go/src/net/lookup.go:110
	}() {
//line /usr/local/go/src/net/lookup.go:110
		_go_fuzz_dep_.CoverTab[15069]++
							n = 0
//line /usr/local/go/src/net/lookup.go:111
		// _ = "end of CoverTab[15069]"
	} else {
//line /usr/local/go/src/net/lookup.go:112
		_go_fuzz_dep_.CoverTab[15070]++
//line /usr/local/go/src/net/lookup.go:112
		// _ = "end of CoverTab[15070]"
//line /usr/local/go/src/net/lookup.go:112
	}
//line /usr/local/go/src/net/lookup.go:112
	// _ = "end of CoverTab[15064]"
//line /usr/local/go/src/net/lookup.go:112
	_go_fuzz_dep_.CoverTab[15065]++
						return n
//line /usr/local/go/src/net/lookup.go:113
	// _ = "end of CoverTab[15065]"
}

// DefaultResolver is the resolver used by the package-level Lookup
//line /usr/local/go/src/net/lookup.go:116
// functions and by Dialers without a specified Resolver.
//line /usr/local/go/src/net/lookup.go:118
var DefaultResolver = &Resolver{}

// A Resolver looks up names and numbers.
//line /usr/local/go/src/net/lookup.go:120
//
//line /usr/local/go/src/net/lookup.go:120
// A nil *Resolver is equivalent to a zero Resolver.
//line /usr/local/go/src/net/lookup.go:123
type Resolver struct {
	// PreferGo controls whether Go's built-in DNS resolver is preferred
	// on platforms where it's available. It is equivalent to setting
	// GODEBUG=netdns=go, but scoped to just this resolver.
	PreferGo	bool

	// StrictErrors controls the behavior of temporary errors
	// (including timeout, socket errors, and SERVFAIL) when using
	// Go's built-in resolver. For a query composed of multiple
	// sub-queries (such as an A+AAAA address lookup, or walking the
	// DNS search list), this option causes such errors to abort the
	// whole query instead of returning a partial result. This is
	// not enabled by default because it may affect compatibility
	// with resolvers that process AAAA queries incorrectly.
	StrictErrors	bool

	// Dial optionally specifies an alternate dialer for use by
	// Go's built-in DNS resolver to make TCP and UDP connections
	// to DNS services. The host in the address parameter will
	// always be a literal IP address and not a host name, and the
	// port in the address parameter will be a literal port number
	// and not a service name.
	// If the Conn returned is also a PacketConn, sent and received DNS
	// messages must adhere to RFC 1035 section 4.2.1, "UDP usage".
	// Otherwise, DNS messages transmitted over Conn must adhere
	// to RFC 7766 section 5, "Transport Protocol Selection".
	// If nil, the default dialer is used.
	Dial	func(ctx context.Context, network, address string) (Conn, error)

	// lookupGroup merges LookupIPAddr calls together for lookups for the same
	// host. The lookupGroup key is the LookupIPAddr.host argument.
	// The return values are ([]IPAddr, error).
	lookupGroup	singleflight.Group
//line /usr/local/go/src/net/lookup.go:159
}

func (r *Resolver) preferGo() bool {
//line /usr/local/go/src/net/lookup.go:161
	_go_fuzz_dep_.CoverTab[15071]++
//line /usr/local/go/src/net/lookup.go:161
	return r != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:161
		_go_fuzz_dep_.CoverTab[15072]++
//line /usr/local/go/src/net/lookup.go:161
		return r.PreferGo
//line /usr/local/go/src/net/lookup.go:161
		// _ = "end of CoverTab[15072]"
//line /usr/local/go/src/net/lookup.go:161
	}()
//line /usr/local/go/src/net/lookup.go:161
	// _ = "end of CoverTab[15071]"
//line /usr/local/go/src/net/lookup.go:161
}
func (r *Resolver) strictErrors() bool {
//line /usr/local/go/src/net/lookup.go:162
	_go_fuzz_dep_.CoverTab[15073]++
//line /usr/local/go/src/net/lookup.go:162
	return r != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:162
		_go_fuzz_dep_.CoverTab[15074]++
//line /usr/local/go/src/net/lookup.go:162
		return r.StrictErrors
//line /usr/local/go/src/net/lookup.go:162
		// _ = "end of CoverTab[15074]"
//line /usr/local/go/src/net/lookup.go:162
	}()
//line /usr/local/go/src/net/lookup.go:162
	// _ = "end of CoverTab[15073]"
//line /usr/local/go/src/net/lookup.go:162
}

func (r *Resolver) getLookupGroup() *singleflight.Group {
//line /usr/local/go/src/net/lookup.go:164
	_go_fuzz_dep_.CoverTab[15075]++
						if r == nil {
//line /usr/local/go/src/net/lookup.go:165
		_go_fuzz_dep_.CoverTab[15077]++
							return &DefaultResolver.lookupGroup
//line /usr/local/go/src/net/lookup.go:166
		// _ = "end of CoverTab[15077]"
	} else {
//line /usr/local/go/src/net/lookup.go:167
		_go_fuzz_dep_.CoverTab[15078]++
//line /usr/local/go/src/net/lookup.go:167
		// _ = "end of CoverTab[15078]"
//line /usr/local/go/src/net/lookup.go:167
	}
//line /usr/local/go/src/net/lookup.go:167
	// _ = "end of CoverTab[15075]"
//line /usr/local/go/src/net/lookup.go:167
	_go_fuzz_dep_.CoverTab[15076]++
						return &r.lookupGroup
//line /usr/local/go/src/net/lookup.go:168
	// _ = "end of CoverTab[15076]"
}

// LookupHost looks up the given host using the local resolver.
//line /usr/local/go/src/net/lookup.go:171
// It returns a slice of that host's addresses.
//line /usr/local/go/src/net/lookup.go:171
//
//line /usr/local/go/src/net/lookup.go:171
// LookupHost uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/lookup.go:171
// Resolver.LookupHost.
//line /usr/local/go/src/net/lookup.go:176
func LookupHost(host string) (addrs []string, err error) {
//line /usr/local/go/src/net/lookup.go:176
	_go_fuzz_dep_.CoverTab[15079]++
						return DefaultResolver.LookupHost(context.Background(), host)
//line /usr/local/go/src/net/lookup.go:177
	// _ = "end of CoverTab[15079]"
}

// LookupHost looks up the given host using the local resolver.
//line /usr/local/go/src/net/lookup.go:180
// It returns a slice of that host's addresses.
//line /usr/local/go/src/net/lookup.go:182
func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error) {
//line /usr/local/go/src/net/lookup.go:182
	_go_fuzz_dep_.CoverTab[15080]++

//line /usr/local/go/src/net/lookup.go:185
	if host == "" {
//line /usr/local/go/src/net/lookup.go:185
		_go_fuzz_dep_.CoverTab[15083]++
							return nil, &DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
//line /usr/local/go/src/net/lookup.go:186
		// _ = "end of CoverTab[15083]"
	} else {
//line /usr/local/go/src/net/lookup.go:187
		_go_fuzz_dep_.CoverTab[15084]++
//line /usr/local/go/src/net/lookup.go:187
		// _ = "end of CoverTab[15084]"
//line /usr/local/go/src/net/lookup.go:187
	}
//line /usr/local/go/src/net/lookup.go:187
	// _ = "end of CoverTab[15080]"
//line /usr/local/go/src/net/lookup.go:187
	_go_fuzz_dep_.CoverTab[15081]++
						if ip, _ := parseIPZone(host); ip != nil {
//line /usr/local/go/src/net/lookup.go:188
		_go_fuzz_dep_.CoverTab[15085]++
							return []string{host}, nil
//line /usr/local/go/src/net/lookup.go:189
		// _ = "end of CoverTab[15085]"
	} else {
//line /usr/local/go/src/net/lookup.go:190
		_go_fuzz_dep_.CoverTab[15086]++
//line /usr/local/go/src/net/lookup.go:190
		// _ = "end of CoverTab[15086]"
//line /usr/local/go/src/net/lookup.go:190
	}
//line /usr/local/go/src/net/lookup.go:190
	// _ = "end of CoverTab[15081]"
//line /usr/local/go/src/net/lookup.go:190
	_go_fuzz_dep_.CoverTab[15082]++
						return r.lookupHost(ctx, host)
//line /usr/local/go/src/net/lookup.go:191
	// _ = "end of CoverTab[15082]"
}

// LookupIP looks up host using the local resolver.
//line /usr/local/go/src/net/lookup.go:194
// It returns a slice of that host's IPv4 and IPv6 addresses.
//line /usr/local/go/src/net/lookup.go:196
func LookupIP(host string) ([]IP, error) {
//line /usr/local/go/src/net/lookup.go:196
	_go_fuzz_dep_.CoverTab[15087]++
						addrs, err := DefaultResolver.LookupIPAddr(context.Background(), host)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:198
		_go_fuzz_dep_.CoverTab[15090]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:199
		// _ = "end of CoverTab[15090]"
	} else {
//line /usr/local/go/src/net/lookup.go:200
		_go_fuzz_dep_.CoverTab[15091]++
//line /usr/local/go/src/net/lookup.go:200
		// _ = "end of CoverTab[15091]"
//line /usr/local/go/src/net/lookup.go:200
	}
//line /usr/local/go/src/net/lookup.go:200
	// _ = "end of CoverTab[15087]"
//line /usr/local/go/src/net/lookup.go:200
	_go_fuzz_dep_.CoverTab[15088]++
						ips := make([]IP, len(addrs))
						for i, ia := range addrs {
//line /usr/local/go/src/net/lookup.go:202
		_go_fuzz_dep_.CoverTab[15092]++
							ips[i] = ia.IP
//line /usr/local/go/src/net/lookup.go:203
		// _ = "end of CoverTab[15092]"
	}
//line /usr/local/go/src/net/lookup.go:204
	// _ = "end of CoverTab[15088]"
//line /usr/local/go/src/net/lookup.go:204
	_go_fuzz_dep_.CoverTab[15089]++
						return ips, nil
//line /usr/local/go/src/net/lookup.go:205
	// _ = "end of CoverTab[15089]"
}

// LookupIPAddr looks up host using the local resolver.
//line /usr/local/go/src/net/lookup.go:208
// It returns a slice of that host's IPv4 and IPv6 addresses.
//line /usr/local/go/src/net/lookup.go:210
func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error) {
//line /usr/local/go/src/net/lookup.go:210
	_go_fuzz_dep_.CoverTab[15093]++
						return r.lookupIPAddr(ctx, "ip", host)
//line /usr/local/go/src/net/lookup.go:211
	// _ = "end of CoverTab[15093]"
}

// LookupIP looks up host for the given network using the local resolver.
//line /usr/local/go/src/net/lookup.go:214
// It returns a slice of that host's IP addresses of the type specified by
//line /usr/local/go/src/net/lookup.go:214
// network.
//line /usr/local/go/src/net/lookup.go:214
// network must be one of "ip", "ip4" or "ip6".
//line /usr/local/go/src/net/lookup.go:218
func (r *Resolver) LookupIP(ctx context.Context, network, host string) ([]IP, error) {
//line /usr/local/go/src/net/lookup.go:218
	_go_fuzz_dep_.CoverTab[15094]++
						afnet, _, err := parseNetwork(ctx, network, false)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:220
		_go_fuzz_dep_.CoverTab[15100]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:221
		// _ = "end of CoverTab[15100]"
	} else {
//line /usr/local/go/src/net/lookup.go:222
		_go_fuzz_dep_.CoverTab[15101]++
//line /usr/local/go/src/net/lookup.go:222
		// _ = "end of CoverTab[15101]"
//line /usr/local/go/src/net/lookup.go:222
	}
//line /usr/local/go/src/net/lookup.go:222
	// _ = "end of CoverTab[15094]"
//line /usr/local/go/src/net/lookup.go:222
	_go_fuzz_dep_.CoverTab[15095]++
						switch afnet {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/lookup.go:224
		_go_fuzz_dep_.CoverTab[15102]++
//line /usr/local/go/src/net/lookup.go:224
		// _ = "end of CoverTab[15102]"
	default:
//line /usr/local/go/src/net/lookup.go:225
		_go_fuzz_dep_.CoverTab[15103]++
							return nil, UnknownNetworkError(network)
//line /usr/local/go/src/net/lookup.go:226
		// _ = "end of CoverTab[15103]"
	}
//line /usr/local/go/src/net/lookup.go:227
	// _ = "end of CoverTab[15095]"
//line /usr/local/go/src/net/lookup.go:227
	_go_fuzz_dep_.CoverTab[15096]++

						if host == "" {
//line /usr/local/go/src/net/lookup.go:229
		_go_fuzz_dep_.CoverTab[15104]++
							return nil, &DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
//line /usr/local/go/src/net/lookup.go:230
		// _ = "end of CoverTab[15104]"
	} else {
//line /usr/local/go/src/net/lookup.go:231
		_go_fuzz_dep_.CoverTab[15105]++
//line /usr/local/go/src/net/lookup.go:231
		// _ = "end of CoverTab[15105]"
//line /usr/local/go/src/net/lookup.go:231
	}
//line /usr/local/go/src/net/lookup.go:231
	// _ = "end of CoverTab[15096]"
//line /usr/local/go/src/net/lookup.go:231
	_go_fuzz_dep_.CoverTab[15097]++
						addrs, err := r.internetAddrList(ctx, afnet, host)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:233
		_go_fuzz_dep_.CoverTab[15106]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:234
		// _ = "end of CoverTab[15106]"
	} else {
//line /usr/local/go/src/net/lookup.go:235
		_go_fuzz_dep_.CoverTab[15107]++
//line /usr/local/go/src/net/lookup.go:235
		// _ = "end of CoverTab[15107]"
//line /usr/local/go/src/net/lookup.go:235
	}
//line /usr/local/go/src/net/lookup.go:235
	// _ = "end of CoverTab[15097]"
//line /usr/local/go/src/net/lookup.go:235
	_go_fuzz_dep_.CoverTab[15098]++

						ips := make([]IP, 0, len(addrs))
						for _, addr := range addrs {
//line /usr/local/go/src/net/lookup.go:238
		_go_fuzz_dep_.CoverTab[15108]++
							ips = append(ips, addr.(*IPAddr).IP)
//line /usr/local/go/src/net/lookup.go:239
		// _ = "end of CoverTab[15108]"
	}
//line /usr/local/go/src/net/lookup.go:240
	// _ = "end of CoverTab[15098]"
//line /usr/local/go/src/net/lookup.go:240
	_go_fuzz_dep_.CoverTab[15099]++
						return ips, nil
//line /usr/local/go/src/net/lookup.go:241
	// _ = "end of CoverTab[15099]"
}

// LookupNetIP looks up host using the local resolver.
//line /usr/local/go/src/net/lookup.go:244
// It returns a slice of that host's IP addresses of the type specified by
//line /usr/local/go/src/net/lookup.go:244
// network.
//line /usr/local/go/src/net/lookup.go:244
// The network must be one of "ip", "ip4" or "ip6".
//line /usr/local/go/src/net/lookup.go:248
func (r *Resolver) LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error) {
//line /usr/local/go/src/net/lookup.go:248
	_go_fuzz_dep_.CoverTab[15109]++

//line /usr/local/go/src/net/lookup.go:253
	ips, err := r.LookupIP(ctx, network, host)
	if err != nil {
//line /usr/local/go/src/net/lookup.go:254
		_go_fuzz_dep_.CoverTab[15112]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:255
		// _ = "end of CoverTab[15112]"
	} else {
//line /usr/local/go/src/net/lookup.go:256
		_go_fuzz_dep_.CoverTab[15113]++
//line /usr/local/go/src/net/lookup.go:256
		// _ = "end of CoverTab[15113]"
//line /usr/local/go/src/net/lookup.go:256
	}
//line /usr/local/go/src/net/lookup.go:256
	// _ = "end of CoverTab[15109]"
//line /usr/local/go/src/net/lookup.go:256
	_go_fuzz_dep_.CoverTab[15110]++
						ret := make([]netip.Addr, 0, len(ips))
						for _, ip := range ips {
//line /usr/local/go/src/net/lookup.go:258
		_go_fuzz_dep_.CoverTab[15114]++
							if a, ok := netip.AddrFromSlice(ip); ok {
//line /usr/local/go/src/net/lookup.go:259
			_go_fuzz_dep_.CoverTab[15115]++
								ret = append(ret, a)
//line /usr/local/go/src/net/lookup.go:260
			// _ = "end of CoverTab[15115]"
		} else {
//line /usr/local/go/src/net/lookup.go:261
			_go_fuzz_dep_.CoverTab[15116]++
//line /usr/local/go/src/net/lookup.go:261
			// _ = "end of CoverTab[15116]"
//line /usr/local/go/src/net/lookup.go:261
		}
//line /usr/local/go/src/net/lookup.go:261
		// _ = "end of CoverTab[15114]"
	}
//line /usr/local/go/src/net/lookup.go:262
	// _ = "end of CoverTab[15110]"
//line /usr/local/go/src/net/lookup.go:262
	_go_fuzz_dep_.CoverTab[15111]++
						return ret, nil
//line /usr/local/go/src/net/lookup.go:263
	// _ = "end of CoverTab[15111]"
}

// onlyValuesCtx is a context that uses an underlying context
//line /usr/local/go/src/net/lookup.go:266
// for value lookup if the underlying context hasn't yet expired.
//line /usr/local/go/src/net/lookup.go:268
type onlyValuesCtx struct {
	context.Context
	lookupValues	context.Context
}

var _ context.Context = (*onlyValuesCtx)(nil)

// Value performs a lookup if the original context hasn't expired.
func (ovc *onlyValuesCtx) Value(key any) any {
//line /usr/local/go/src/net/lookup.go:276
	_go_fuzz_dep_.CoverTab[15117]++
						select {
	case <-ovc.lookupValues.Done():
//line /usr/local/go/src/net/lookup.go:278
		_go_fuzz_dep_.CoverTab[15118]++
							return nil
//line /usr/local/go/src/net/lookup.go:279
		// _ = "end of CoverTab[15118]"
	default:
//line /usr/local/go/src/net/lookup.go:280
		_go_fuzz_dep_.CoverTab[15119]++
							return ovc.lookupValues.Value(key)
//line /usr/local/go/src/net/lookup.go:281
		// _ = "end of CoverTab[15119]"
	}
//line /usr/local/go/src/net/lookup.go:282
	// _ = "end of CoverTab[15117]"
}

// withUnexpiredValuesPreserved returns a context.Context that only uses lookupCtx
//line /usr/local/go/src/net/lookup.go:285
// for its values, otherwise it is never canceled and has no deadline.
//line /usr/local/go/src/net/lookup.go:285
// If the lookup context expires, any looked up values will return nil.
//line /usr/local/go/src/net/lookup.go:285
// See Issue 28600.
//line /usr/local/go/src/net/lookup.go:289
func withUnexpiredValuesPreserved(lookupCtx context.Context) context.Context {
//line /usr/local/go/src/net/lookup.go:289
	_go_fuzz_dep_.CoverTab[15120]++
						return &onlyValuesCtx{Context: context.Background(), lookupValues: lookupCtx}
//line /usr/local/go/src/net/lookup.go:290
	// _ = "end of CoverTab[15120]"
}

// lookupIPAddr looks up host using the local resolver and particular network.
//line /usr/local/go/src/net/lookup.go:293
// It returns a slice of that host's IPv4 and IPv6 addresses.
//line /usr/local/go/src/net/lookup.go:295
func (r *Resolver) lookupIPAddr(ctx context.Context, network, host string) ([]IPAddr, error) {
//line /usr/local/go/src/net/lookup.go:295
	_go_fuzz_dep_.CoverTab[15121]++

//line /usr/local/go/src/net/lookup.go:298
	if host == "" {
//line /usr/local/go/src/net/lookup.go:298
		_go_fuzz_dep_.CoverTab[15128]++
							return nil, &DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
//line /usr/local/go/src/net/lookup.go:299
		// _ = "end of CoverTab[15128]"
	} else {
//line /usr/local/go/src/net/lookup.go:300
		_go_fuzz_dep_.CoverTab[15129]++
//line /usr/local/go/src/net/lookup.go:300
		// _ = "end of CoverTab[15129]"
//line /usr/local/go/src/net/lookup.go:300
	}
//line /usr/local/go/src/net/lookup.go:300
	// _ = "end of CoverTab[15121]"
//line /usr/local/go/src/net/lookup.go:300
	_go_fuzz_dep_.CoverTab[15122]++
						if ip, zone := parseIPZone(host); ip != nil {
//line /usr/local/go/src/net/lookup.go:301
		_go_fuzz_dep_.CoverTab[15130]++
							return []IPAddr{{IP: ip, Zone: zone}}, nil
//line /usr/local/go/src/net/lookup.go:302
		// _ = "end of CoverTab[15130]"
	} else {
//line /usr/local/go/src/net/lookup.go:303
		_go_fuzz_dep_.CoverTab[15131]++
//line /usr/local/go/src/net/lookup.go:303
		// _ = "end of CoverTab[15131]"
//line /usr/local/go/src/net/lookup.go:303
	}
//line /usr/local/go/src/net/lookup.go:303
	// _ = "end of CoverTab[15122]"
//line /usr/local/go/src/net/lookup.go:303
	_go_fuzz_dep_.CoverTab[15123]++
						trace, _ := ctx.Value(nettrace.TraceKey{}).(*nettrace.Trace)
						if trace != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:305
		_go_fuzz_dep_.CoverTab[15132]++
//line /usr/local/go/src/net/lookup.go:305
		return trace.DNSStart != nil
//line /usr/local/go/src/net/lookup.go:305
		// _ = "end of CoverTab[15132]"
//line /usr/local/go/src/net/lookup.go:305
	}() {
//line /usr/local/go/src/net/lookup.go:305
		_go_fuzz_dep_.CoverTab[15133]++
							trace.DNSStart(host)
//line /usr/local/go/src/net/lookup.go:306
		// _ = "end of CoverTab[15133]"
	} else {
//line /usr/local/go/src/net/lookup.go:307
		_go_fuzz_dep_.CoverTab[15134]++
//line /usr/local/go/src/net/lookup.go:307
		// _ = "end of CoverTab[15134]"
//line /usr/local/go/src/net/lookup.go:307
	}
//line /usr/local/go/src/net/lookup.go:307
	// _ = "end of CoverTab[15123]"
//line /usr/local/go/src/net/lookup.go:307
	_go_fuzz_dep_.CoverTab[15124]++

//line /usr/local/go/src/net/lookup.go:311
	resolverFunc := r.lookupIP
	if alt, _ := ctx.Value(nettrace.LookupIPAltResolverKey{}).(func(context.Context, string, string) ([]IPAddr, error)); alt != nil {
//line /usr/local/go/src/net/lookup.go:312
		_go_fuzz_dep_.CoverTab[15135]++
							resolverFunc = alt
//line /usr/local/go/src/net/lookup.go:313
		// _ = "end of CoverTab[15135]"
	} else {
//line /usr/local/go/src/net/lookup.go:314
		_go_fuzz_dep_.CoverTab[15136]++
//line /usr/local/go/src/net/lookup.go:314
		// _ = "end of CoverTab[15136]"
//line /usr/local/go/src/net/lookup.go:314
	}
//line /usr/local/go/src/net/lookup.go:314
	// _ = "end of CoverTab[15124]"
//line /usr/local/go/src/net/lookup.go:314
	_go_fuzz_dep_.CoverTab[15125]++

//line /usr/local/go/src/net/lookup.go:321
	lookupGroupCtx, lookupGroupCancel := context.WithCancel(withUnexpiredValuesPreserved(ctx))

	lookupKey := network + "\000" + host
	dnsWaitGroup.Add(1)
	ch := r.getLookupGroup().DoChan(lookupKey, func() (any, error) {
//line /usr/local/go/src/net/lookup.go:325
		_go_fuzz_dep_.CoverTab[15137]++
							return testHookLookupIP(lookupGroupCtx, resolverFunc, network, host)
//line /usr/local/go/src/net/lookup.go:326
		// _ = "end of CoverTab[15137]"
	})
//line /usr/local/go/src/net/lookup.go:327
	// _ = "end of CoverTab[15125]"
//line /usr/local/go/src/net/lookup.go:327
	_go_fuzz_dep_.CoverTab[15126]++

						dnsWaitGroupDone := func(ch <-chan singleflight.Result, cancelFn context.CancelFunc) {
//line /usr/local/go/src/net/lookup.go:329
		_go_fuzz_dep_.CoverTab[15138]++
							<-ch
							dnsWaitGroup.Done()
							cancelFn()
//line /usr/local/go/src/net/lookup.go:332
		// _ = "end of CoverTab[15138]"
	}
//line /usr/local/go/src/net/lookup.go:333
	// _ = "end of CoverTab[15126]"
//line /usr/local/go/src/net/lookup.go:333
	_go_fuzz_dep_.CoverTab[15127]++
						select {
	case <-ctx.Done():
//line /usr/local/go/src/net/lookup.go:335
		_go_fuzz_dep_.CoverTab[15139]++

//line /usr/local/go/src/net/lookup.go:343
		if r.getLookupGroup().ForgetUnshared(lookupKey) {
//line /usr/local/go/src/net/lookup.go:343
			_go_fuzz_dep_.CoverTab[15145]++
								lookupGroupCancel()
//line /usr/local/go/src/net/lookup.go:344
			_curRoutineNum10_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/lookup.go:344
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum10_)
								go func() {
//line /usr/local/go/src/net/lookup.go:345
				_go_fuzz_dep_.CoverTab[15146]++
//line /usr/local/go/src/net/lookup.go:345
				defer func() {
//line /usr/local/go/src/net/lookup.go:345
					_go_fuzz_dep_.CoverTab[15147]++
//line /usr/local/go/src/net/lookup.go:345
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum10_)
//line /usr/local/go/src/net/lookup.go:345
					// _ = "end of CoverTab[15147]"
//line /usr/local/go/src/net/lookup.go:345
				}()
//line /usr/local/go/src/net/lookup.go:345
				dnsWaitGroupDone(ch, func() { _go_fuzz_dep_.CoverTab[15148]++; // _ = "end of CoverTab[15148]" })
//line /usr/local/go/src/net/lookup.go:345
				// _ = "end of CoverTab[15146]"
//line /usr/local/go/src/net/lookup.go:345
			}()
//line /usr/local/go/src/net/lookup.go:345
			// _ = "end of CoverTab[15145]"
		} else {
//line /usr/local/go/src/net/lookup.go:346
			_go_fuzz_dep_.CoverTab[15149]++
//line /usr/local/go/src/net/lookup.go:346
			_curRoutineNum11_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/lookup.go:346
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum11_)
								go func() {
//line /usr/local/go/src/net/lookup.go:347
				_go_fuzz_dep_.CoverTab[15150]++
//line /usr/local/go/src/net/lookup.go:347
				defer func() {
//line /usr/local/go/src/net/lookup.go:347
					_go_fuzz_dep_.CoverTab[15151]++
//line /usr/local/go/src/net/lookup.go:347
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum11_)
//line /usr/local/go/src/net/lookup.go:347
					// _ = "end of CoverTab[15151]"
//line /usr/local/go/src/net/lookup.go:347
				}()
//line /usr/local/go/src/net/lookup.go:347
				dnsWaitGroupDone(ch, lookupGroupCancel)
//line /usr/local/go/src/net/lookup.go:347
				// _ = "end of CoverTab[15150]"
//line /usr/local/go/src/net/lookup.go:347
			}()
//line /usr/local/go/src/net/lookup.go:347
			// _ = "end of CoverTab[15149]"
		}
//line /usr/local/go/src/net/lookup.go:348
		// _ = "end of CoverTab[15139]"
//line /usr/local/go/src/net/lookup.go:348
		_go_fuzz_dep_.CoverTab[15140]++
							ctxErr := ctx.Err()
							err := &DNSError{
			Err:		mapErr(ctxErr).Error(),
			Name:		host,
			IsTimeout:	ctxErr == context.DeadlineExceeded,
		}
		if trace != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:355
			_go_fuzz_dep_.CoverTab[15152]++
//line /usr/local/go/src/net/lookup.go:355
			return trace.DNSDone != nil
//line /usr/local/go/src/net/lookup.go:355
			// _ = "end of CoverTab[15152]"
//line /usr/local/go/src/net/lookup.go:355
		}() {
//line /usr/local/go/src/net/lookup.go:355
			_go_fuzz_dep_.CoverTab[15153]++
								trace.DNSDone(nil, false, err)
//line /usr/local/go/src/net/lookup.go:356
			// _ = "end of CoverTab[15153]"
		} else {
//line /usr/local/go/src/net/lookup.go:357
			_go_fuzz_dep_.CoverTab[15154]++
//line /usr/local/go/src/net/lookup.go:357
			// _ = "end of CoverTab[15154]"
//line /usr/local/go/src/net/lookup.go:357
		}
//line /usr/local/go/src/net/lookup.go:357
		// _ = "end of CoverTab[15140]"
//line /usr/local/go/src/net/lookup.go:357
		_go_fuzz_dep_.CoverTab[15141]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:358
		// _ = "end of CoverTab[15141]"
	case r := <-ch:
//line /usr/local/go/src/net/lookup.go:359
		_go_fuzz_dep_.CoverTab[15142]++
							dnsWaitGroup.Done()
							lookupGroupCancel()
							err := r.Err
							if err != nil {
//line /usr/local/go/src/net/lookup.go:363
			_go_fuzz_dep_.CoverTab[15155]++
								if _, ok := err.(*DNSError); !ok {
//line /usr/local/go/src/net/lookup.go:364
				_go_fuzz_dep_.CoverTab[15156]++
									isTimeout := false
									if err == context.DeadlineExceeded {
//line /usr/local/go/src/net/lookup.go:366
					_go_fuzz_dep_.CoverTab[15158]++
										isTimeout = true
//line /usr/local/go/src/net/lookup.go:367
					// _ = "end of CoverTab[15158]"
				} else {
//line /usr/local/go/src/net/lookup.go:368
					_go_fuzz_dep_.CoverTab[15159]++
//line /usr/local/go/src/net/lookup.go:368
					if terr, ok := err.(timeout); ok {
//line /usr/local/go/src/net/lookup.go:368
						_go_fuzz_dep_.CoverTab[15160]++
											isTimeout = terr.Timeout()
//line /usr/local/go/src/net/lookup.go:369
						// _ = "end of CoverTab[15160]"
					} else {
//line /usr/local/go/src/net/lookup.go:370
						_go_fuzz_dep_.CoverTab[15161]++
//line /usr/local/go/src/net/lookup.go:370
						// _ = "end of CoverTab[15161]"
//line /usr/local/go/src/net/lookup.go:370
					}
//line /usr/local/go/src/net/lookup.go:370
					// _ = "end of CoverTab[15159]"
//line /usr/local/go/src/net/lookup.go:370
				}
//line /usr/local/go/src/net/lookup.go:370
				// _ = "end of CoverTab[15156]"
//line /usr/local/go/src/net/lookup.go:370
				_go_fuzz_dep_.CoverTab[15157]++
									err = &DNSError{
					Err:		err.Error(),
					Name:		host,
					IsTimeout:	isTimeout,
				}
//line /usr/local/go/src/net/lookup.go:375
				// _ = "end of CoverTab[15157]"
			} else {
//line /usr/local/go/src/net/lookup.go:376
				_go_fuzz_dep_.CoverTab[15162]++
//line /usr/local/go/src/net/lookup.go:376
				// _ = "end of CoverTab[15162]"
//line /usr/local/go/src/net/lookup.go:376
			}
//line /usr/local/go/src/net/lookup.go:376
			// _ = "end of CoverTab[15155]"
		} else {
//line /usr/local/go/src/net/lookup.go:377
			_go_fuzz_dep_.CoverTab[15163]++
//line /usr/local/go/src/net/lookup.go:377
			// _ = "end of CoverTab[15163]"
//line /usr/local/go/src/net/lookup.go:377
		}
//line /usr/local/go/src/net/lookup.go:377
		// _ = "end of CoverTab[15142]"
//line /usr/local/go/src/net/lookup.go:377
		_go_fuzz_dep_.CoverTab[15143]++
							if trace != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:378
			_go_fuzz_dep_.CoverTab[15164]++
//line /usr/local/go/src/net/lookup.go:378
			return trace.DNSDone != nil
//line /usr/local/go/src/net/lookup.go:378
			// _ = "end of CoverTab[15164]"
//line /usr/local/go/src/net/lookup.go:378
		}() {
//line /usr/local/go/src/net/lookup.go:378
			_go_fuzz_dep_.CoverTab[15165]++
								addrs, _ := r.Val.([]IPAddr)
								trace.DNSDone(ipAddrsEface(addrs), r.Shared, err)
//line /usr/local/go/src/net/lookup.go:380
			// _ = "end of CoverTab[15165]"
		} else {
//line /usr/local/go/src/net/lookup.go:381
			_go_fuzz_dep_.CoverTab[15166]++
//line /usr/local/go/src/net/lookup.go:381
			// _ = "end of CoverTab[15166]"
//line /usr/local/go/src/net/lookup.go:381
		}
//line /usr/local/go/src/net/lookup.go:381
		// _ = "end of CoverTab[15143]"
//line /usr/local/go/src/net/lookup.go:381
		_go_fuzz_dep_.CoverTab[15144]++
							return lookupIPReturn(r.Val, err, r.Shared)
//line /usr/local/go/src/net/lookup.go:382
		// _ = "end of CoverTab[15144]"
	}
//line /usr/local/go/src/net/lookup.go:383
	// _ = "end of CoverTab[15127]"
}

// lookupIPReturn turns the return values from singleflight.Do into
//line /usr/local/go/src/net/lookup.go:386
// the return values from LookupIP.
//line /usr/local/go/src/net/lookup.go:388
func lookupIPReturn(addrsi any, err error, shared bool) ([]IPAddr, error) {
//line /usr/local/go/src/net/lookup.go:388
	_go_fuzz_dep_.CoverTab[15167]++
						if err != nil {
//line /usr/local/go/src/net/lookup.go:389
		_go_fuzz_dep_.CoverTab[15170]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:390
		// _ = "end of CoverTab[15170]"
	} else {
//line /usr/local/go/src/net/lookup.go:391
		_go_fuzz_dep_.CoverTab[15171]++
//line /usr/local/go/src/net/lookup.go:391
		// _ = "end of CoverTab[15171]"
//line /usr/local/go/src/net/lookup.go:391
	}
//line /usr/local/go/src/net/lookup.go:391
	// _ = "end of CoverTab[15167]"
//line /usr/local/go/src/net/lookup.go:391
	_go_fuzz_dep_.CoverTab[15168]++
						addrs := addrsi.([]IPAddr)
						if shared {
//line /usr/local/go/src/net/lookup.go:393
		_go_fuzz_dep_.CoverTab[15172]++
							clone := make([]IPAddr, len(addrs))
							copy(clone, addrs)
							addrs = clone
//line /usr/local/go/src/net/lookup.go:396
		// _ = "end of CoverTab[15172]"
	} else {
//line /usr/local/go/src/net/lookup.go:397
		_go_fuzz_dep_.CoverTab[15173]++
//line /usr/local/go/src/net/lookup.go:397
		// _ = "end of CoverTab[15173]"
//line /usr/local/go/src/net/lookup.go:397
	}
//line /usr/local/go/src/net/lookup.go:397
	// _ = "end of CoverTab[15168]"
//line /usr/local/go/src/net/lookup.go:397
	_go_fuzz_dep_.CoverTab[15169]++
						return addrs, nil
//line /usr/local/go/src/net/lookup.go:398
	// _ = "end of CoverTab[15169]"
}

// ipAddrsEface returns an empty interface slice of addrs.
func ipAddrsEface(addrs []IPAddr) []any {
//line /usr/local/go/src/net/lookup.go:402
	_go_fuzz_dep_.CoverTab[15174]++
						s := make([]any, len(addrs))
						for i, v := range addrs {
//line /usr/local/go/src/net/lookup.go:404
		_go_fuzz_dep_.CoverTab[15176]++
							s[i] = v
//line /usr/local/go/src/net/lookup.go:405
		// _ = "end of CoverTab[15176]"
	}
//line /usr/local/go/src/net/lookup.go:406
	// _ = "end of CoverTab[15174]"
//line /usr/local/go/src/net/lookup.go:406
	_go_fuzz_dep_.CoverTab[15175]++
						return s
//line /usr/local/go/src/net/lookup.go:407
	// _ = "end of CoverTab[15175]"
}

// LookupPort looks up the port for the given network and service.
//line /usr/local/go/src/net/lookup.go:410
//
//line /usr/local/go/src/net/lookup.go:410
// LookupPort uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/lookup.go:410
// Resolver.LookupPort.
//line /usr/local/go/src/net/lookup.go:414
func LookupPort(network, service string) (port int, err error) {
//line /usr/local/go/src/net/lookup.go:414
	_go_fuzz_dep_.CoverTab[15177]++
						return DefaultResolver.LookupPort(context.Background(), network, service)
//line /usr/local/go/src/net/lookup.go:415
	// _ = "end of CoverTab[15177]"
}

// LookupPort looks up the port for the given network and service.
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error) {
//line /usr/local/go/src/net/lookup.go:419
	_go_fuzz_dep_.CoverTab[15178]++
						port, needsLookup := parsePort(service)
						if needsLookup {
//line /usr/local/go/src/net/lookup.go:421
		_go_fuzz_dep_.CoverTab[15181]++
							switch network {
		case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
//line /usr/local/go/src/net/lookup.go:423
			_go_fuzz_dep_.CoverTab[15183]++
//line /usr/local/go/src/net/lookup.go:423
			// _ = "end of CoverTab[15183]"
		case "":
//line /usr/local/go/src/net/lookup.go:424
			_go_fuzz_dep_.CoverTab[15184]++
								network = "ip"
//line /usr/local/go/src/net/lookup.go:425
			// _ = "end of CoverTab[15184]"
		default:
//line /usr/local/go/src/net/lookup.go:426
			_go_fuzz_dep_.CoverTab[15185]++
								return 0, &AddrError{Err: "unknown network", Addr: network}
//line /usr/local/go/src/net/lookup.go:427
			// _ = "end of CoverTab[15185]"
		}
//line /usr/local/go/src/net/lookup.go:428
		// _ = "end of CoverTab[15181]"
//line /usr/local/go/src/net/lookup.go:428
		_go_fuzz_dep_.CoverTab[15182]++
							port, err = r.lookupPort(ctx, network, service)
							if err != nil {
//line /usr/local/go/src/net/lookup.go:430
			_go_fuzz_dep_.CoverTab[15186]++
								return 0, err
//line /usr/local/go/src/net/lookup.go:431
			// _ = "end of CoverTab[15186]"
		} else {
//line /usr/local/go/src/net/lookup.go:432
			_go_fuzz_dep_.CoverTab[15187]++
//line /usr/local/go/src/net/lookup.go:432
			// _ = "end of CoverTab[15187]"
//line /usr/local/go/src/net/lookup.go:432
		}
//line /usr/local/go/src/net/lookup.go:432
		// _ = "end of CoverTab[15182]"
	} else {
//line /usr/local/go/src/net/lookup.go:433
		_go_fuzz_dep_.CoverTab[15188]++
//line /usr/local/go/src/net/lookup.go:433
		// _ = "end of CoverTab[15188]"
//line /usr/local/go/src/net/lookup.go:433
	}
//line /usr/local/go/src/net/lookup.go:433
	// _ = "end of CoverTab[15178]"
//line /usr/local/go/src/net/lookup.go:433
	_go_fuzz_dep_.CoverTab[15179]++
						if 0 > port || func() bool {
//line /usr/local/go/src/net/lookup.go:434
		_go_fuzz_dep_.CoverTab[15189]++
//line /usr/local/go/src/net/lookup.go:434
		return port > 65535
//line /usr/local/go/src/net/lookup.go:434
		// _ = "end of CoverTab[15189]"
//line /usr/local/go/src/net/lookup.go:434
	}() {
//line /usr/local/go/src/net/lookup.go:434
		_go_fuzz_dep_.CoverTab[15190]++
							return 0, &AddrError{Err: "invalid port", Addr: service}
//line /usr/local/go/src/net/lookup.go:435
		// _ = "end of CoverTab[15190]"
	} else {
//line /usr/local/go/src/net/lookup.go:436
		_go_fuzz_dep_.CoverTab[15191]++
//line /usr/local/go/src/net/lookup.go:436
		// _ = "end of CoverTab[15191]"
//line /usr/local/go/src/net/lookup.go:436
	}
//line /usr/local/go/src/net/lookup.go:436
	// _ = "end of CoverTab[15179]"
//line /usr/local/go/src/net/lookup.go:436
	_go_fuzz_dep_.CoverTab[15180]++
						return port, nil
//line /usr/local/go/src/net/lookup.go:437
	// _ = "end of CoverTab[15180]"
}

// LookupCNAME returns the canonical name for the given host.
//line /usr/local/go/src/net/lookup.go:440
// Callers that do not care about the canonical name can call
//line /usr/local/go/src/net/lookup.go:440
// LookupHost or LookupIP directly; both take care of resolving
//line /usr/local/go/src/net/lookup.go:440
// the canonical name as part of the lookup.
//line /usr/local/go/src/net/lookup.go:440
//
//line /usr/local/go/src/net/lookup.go:440
// A canonical name is the final name after following zero
//line /usr/local/go/src/net/lookup.go:440
// or more CNAME records.
//line /usr/local/go/src/net/lookup.go:440
// LookupCNAME does not return an error if host does not
//line /usr/local/go/src/net/lookup.go:440
// contain DNS "CNAME" records, as long as host resolves to
//line /usr/local/go/src/net/lookup.go:440
// address records.
//line /usr/local/go/src/net/lookup.go:440
//
//line /usr/local/go/src/net/lookup.go:440
// The returned canonical name is validated to be a properly
//line /usr/local/go/src/net/lookup.go:440
// formatted presentation-format domain name.
//line /usr/local/go/src/net/lookup.go:440
//
//line /usr/local/go/src/net/lookup.go:440
// LookupCNAME uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/lookup.go:440
// Resolver.LookupCNAME.
//line /usr/local/go/src/net/lookup.go:456
func LookupCNAME(host string) (cname string, err error) {
//line /usr/local/go/src/net/lookup.go:456
	_go_fuzz_dep_.CoverTab[15192]++
						return DefaultResolver.LookupCNAME(context.Background(), host)
//line /usr/local/go/src/net/lookup.go:457
	// _ = "end of CoverTab[15192]"
}

// LookupCNAME returns the canonical name for the given host.
//line /usr/local/go/src/net/lookup.go:460
// Callers that do not care about the canonical name can call
//line /usr/local/go/src/net/lookup.go:460
// LookupHost or LookupIP directly; both take care of resolving
//line /usr/local/go/src/net/lookup.go:460
// the canonical name as part of the lookup.
//line /usr/local/go/src/net/lookup.go:460
//
//line /usr/local/go/src/net/lookup.go:460
// A canonical name is the final name after following zero
//line /usr/local/go/src/net/lookup.go:460
// or more CNAME records.
//line /usr/local/go/src/net/lookup.go:460
// LookupCNAME does not return an error if host does not
//line /usr/local/go/src/net/lookup.go:460
// contain DNS "CNAME" records, as long as host resolves to
//line /usr/local/go/src/net/lookup.go:460
// address records.
//line /usr/local/go/src/net/lookup.go:460
//
//line /usr/local/go/src/net/lookup.go:460
// The returned canonical name is validated to be a properly
//line /usr/local/go/src/net/lookup.go:460
// formatted presentation-format domain name.
//line /usr/local/go/src/net/lookup.go:473
func (r *Resolver) LookupCNAME(ctx context.Context, host string) (string, error) {
//line /usr/local/go/src/net/lookup.go:473
	_go_fuzz_dep_.CoverTab[15193]++
						cname, err := r.lookupCNAME(ctx, host)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:475
		_go_fuzz_dep_.CoverTab[15196]++
							return "", err
//line /usr/local/go/src/net/lookup.go:476
		// _ = "end of CoverTab[15196]"
	} else {
//line /usr/local/go/src/net/lookup.go:477
		_go_fuzz_dep_.CoverTab[15197]++
//line /usr/local/go/src/net/lookup.go:477
		// _ = "end of CoverTab[15197]"
//line /usr/local/go/src/net/lookup.go:477
	}
//line /usr/local/go/src/net/lookup.go:477
	// _ = "end of CoverTab[15193]"
//line /usr/local/go/src/net/lookup.go:477
	_go_fuzz_dep_.CoverTab[15194]++
						if !isDomainName(cname) {
//line /usr/local/go/src/net/lookup.go:478
		_go_fuzz_dep_.CoverTab[15198]++
							return "", &DNSError{Err: errMalformedDNSRecordsDetail, Name: host}
//line /usr/local/go/src/net/lookup.go:479
		// _ = "end of CoverTab[15198]"
	} else {
//line /usr/local/go/src/net/lookup.go:480
		_go_fuzz_dep_.CoverTab[15199]++
//line /usr/local/go/src/net/lookup.go:480
		// _ = "end of CoverTab[15199]"
//line /usr/local/go/src/net/lookup.go:480
	}
//line /usr/local/go/src/net/lookup.go:480
	// _ = "end of CoverTab[15194]"
//line /usr/local/go/src/net/lookup.go:480
	_go_fuzz_dep_.CoverTab[15195]++
						return cname, nil
//line /usr/local/go/src/net/lookup.go:481
	// _ = "end of CoverTab[15195]"
}

// LookupSRV tries to resolve an SRV query of the given service,
//line /usr/local/go/src/net/lookup.go:484
// protocol, and domain name. The proto is "tcp" or "udp".
//line /usr/local/go/src/net/lookup.go:484
// The returned records are sorted by priority and randomized
//line /usr/local/go/src/net/lookup.go:484
// by weight within a priority.
//line /usr/local/go/src/net/lookup.go:484
//
//line /usr/local/go/src/net/lookup.go:484
// LookupSRV constructs the DNS name to look up following RFC 2782.
//line /usr/local/go/src/net/lookup.go:484
// That is, it looks up _service._proto.name. To accommodate services
//line /usr/local/go/src/net/lookup.go:484
// publishing SRV records under non-standard names, if both service
//line /usr/local/go/src/net/lookup.go:484
// and proto are empty strings, LookupSRV looks up name directly.
//line /usr/local/go/src/net/lookup.go:484
//
//line /usr/local/go/src/net/lookup.go:484
// The returned service names are validated to be properly
//line /usr/local/go/src/net/lookup.go:484
// formatted presentation-format domain names. If the response contains
//line /usr/local/go/src/net/lookup.go:484
// invalid names, those records are filtered out and an error
//line /usr/local/go/src/net/lookup.go:484
// will be returned alongside the remaining results, if any.
//line /usr/local/go/src/net/lookup.go:498
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error) {
//line /usr/local/go/src/net/lookup.go:498
	_go_fuzz_dep_.CoverTab[15200]++
						return DefaultResolver.LookupSRV(context.Background(), service, proto, name)
//line /usr/local/go/src/net/lookup.go:499
	// _ = "end of CoverTab[15200]"
}

// LookupSRV tries to resolve an SRV query of the given service,
//line /usr/local/go/src/net/lookup.go:502
// protocol, and domain name. The proto is "tcp" or "udp".
//line /usr/local/go/src/net/lookup.go:502
// The returned records are sorted by priority and randomized
//line /usr/local/go/src/net/lookup.go:502
// by weight within a priority.
//line /usr/local/go/src/net/lookup.go:502
//
//line /usr/local/go/src/net/lookup.go:502
// LookupSRV constructs the DNS name to look up following RFC 2782.
//line /usr/local/go/src/net/lookup.go:502
// That is, it looks up _service._proto.name. To accommodate services
//line /usr/local/go/src/net/lookup.go:502
// publishing SRV records under non-standard names, if both service
//line /usr/local/go/src/net/lookup.go:502
// and proto are empty strings, LookupSRV looks up name directly.
//line /usr/local/go/src/net/lookup.go:502
//
//line /usr/local/go/src/net/lookup.go:502
// The returned service names are validated to be properly
//line /usr/local/go/src/net/lookup.go:502
// formatted presentation-format domain names. If the response contains
//line /usr/local/go/src/net/lookup.go:502
// invalid names, those records are filtered out and an error
//line /usr/local/go/src/net/lookup.go:502
// will be returned alongside the remaining results, if any.
//line /usr/local/go/src/net/lookup.go:516
func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error) {
//line /usr/local/go/src/net/lookup.go:516
	_go_fuzz_dep_.CoverTab[15201]++
						cname, addrs, err := r.lookupSRV(ctx, service, proto, name)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:518
		_go_fuzz_dep_.CoverTab[15206]++
							return "", nil, err
//line /usr/local/go/src/net/lookup.go:519
		// _ = "end of CoverTab[15206]"
	} else {
//line /usr/local/go/src/net/lookup.go:520
		_go_fuzz_dep_.CoverTab[15207]++
//line /usr/local/go/src/net/lookup.go:520
		// _ = "end of CoverTab[15207]"
//line /usr/local/go/src/net/lookup.go:520
	}
//line /usr/local/go/src/net/lookup.go:520
	// _ = "end of CoverTab[15201]"
//line /usr/local/go/src/net/lookup.go:520
	_go_fuzz_dep_.CoverTab[15202]++
						if cname != "" && func() bool {
//line /usr/local/go/src/net/lookup.go:521
		_go_fuzz_dep_.CoverTab[15208]++
//line /usr/local/go/src/net/lookup.go:521
		return !isDomainName(cname)
//line /usr/local/go/src/net/lookup.go:521
		// _ = "end of CoverTab[15208]"
//line /usr/local/go/src/net/lookup.go:521
	}() {
//line /usr/local/go/src/net/lookup.go:521
		_go_fuzz_dep_.CoverTab[15209]++
							return "", nil, &DNSError{Err: "SRV header name is invalid", Name: name}
//line /usr/local/go/src/net/lookup.go:522
		// _ = "end of CoverTab[15209]"
	} else {
//line /usr/local/go/src/net/lookup.go:523
		_go_fuzz_dep_.CoverTab[15210]++
//line /usr/local/go/src/net/lookup.go:523
		// _ = "end of CoverTab[15210]"
//line /usr/local/go/src/net/lookup.go:523
	}
//line /usr/local/go/src/net/lookup.go:523
	// _ = "end of CoverTab[15202]"
//line /usr/local/go/src/net/lookup.go:523
	_go_fuzz_dep_.CoverTab[15203]++
						filteredAddrs := make([]*SRV, 0, len(addrs))
						for _, addr := range addrs {
//line /usr/local/go/src/net/lookup.go:525
		_go_fuzz_dep_.CoverTab[15211]++
							if addr == nil {
//line /usr/local/go/src/net/lookup.go:526
			_go_fuzz_dep_.CoverTab[15214]++
								continue
//line /usr/local/go/src/net/lookup.go:527
			// _ = "end of CoverTab[15214]"
		} else {
//line /usr/local/go/src/net/lookup.go:528
			_go_fuzz_dep_.CoverTab[15215]++
//line /usr/local/go/src/net/lookup.go:528
			// _ = "end of CoverTab[15215]"
//line /usr/local/go/src/net/lookup.go:528
		}
//line /usr/local/go/src/net/lookup.go:528
		// _ = "end of CoverTab[15211]"
//line /usr/local/go/src/net/lookup.go:528
		_go_fuzz_dep_.CoverTab[15212]++
							if !isDomainName(addr.Target) {
//line /usr/local/go/src/net/lookup.go:529
			_go_fuzz_dep_.CoverTab[15216]++
								continue
//line /usr/local/go/src/net/lookup.go:530
			// _ = "end of CoverTab[15216]"
		} else {
//line /usr/local/go/src/net/lookup.go:531
			_go_fuzz_dep_.CoverTab[15217]++
//line /usr/local/go/src/net/lookup.go:531
			// _ = "end of CoverTab[15217]"
//line /usr/local/go/src/net/lookup.go:531
		}
//line /usr/local/go/src/net/lookup.go:531
		// _ = "end of CoverTab[15212]"
//line /usr/local/go/src/net/lookup.go:531
		_go_fuzz_dep_.CoverTab[15213]++
							filteredAddrs = append(filteredAddrs, addr)
//line /usr/local/go/src/net/lookup.go:532
		// _ = "end of CoverTab[15213]"
	}
//line /usr/local/go/src/net/lookup.go:533
	// _ = "end of CoverTab[15203]"
//line /usr/local/go/src/net/lookup.go:533
	_go_fuzz_dep_.CoverTab[15204]++
						if len(addrs) != len(filteredAddrs) {
//line /usr/local/go/src/net/lookup.go:534
		_go_fuzz_dep_.CoverTab[15218]++
							return cname, filteredAddrs, &DNSError{Err: errMalformedDNSRecordsDetail, Name: name}
//line /usr/local/go/src/net/lookup.go:535
		// _ = "end of CoverTab[15218]"
	} else {
//line /usr/local/go/src/net/lookup.go:536
		_go_fuzz_dep_.CoverTab[15219]++
//line /usr/local/go/src/net/lookup.go:536
		// _ = "end of CoverTab[15219]"
//line /usr/local/go/src/net/lookup.go:536
	}
//line /usr/local/go/src/net/lookup.go:536
	// _ = "end of CoverTab[15204]"
//line /usr/local/go/src/net/lookup.go:536
	_go_fuzz_dep_.CoverTab[15205]++
						return cname, filteredAddrs, nil
//line /usr/local/go/src/net/lookup.go:537
	// _ = "end of CoverTab[15205]"
}

// LookupMX returns the DNS MX records for the given domain name sorted by preference.
//line /usr/local/go/src/net/lookup.go:540
//
//line /usr/local/go/src/net/lookup.go:540
// The returned mail server names are validated to be properly
//line /usr/local/go/src/net/lookup.go:540
// formatted presentation-format domain names. If the response contains
//line /usr/local/go/src/net/lookup.go:540
// invalid names, those records are filtered out and an error
//line /usr/local/go/src/net/lookup.go:540
// will be returned alongside the remaining results, if any.
//line /usr/local/go/src/net/lookup.go:540
//
//line /usr/local/go/src/net/lookup.go:540
// LookupMX uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/lookup.go:540
// Resolver.LookupMX.
//line /usr/local/go/src/net/lookup.go:549
func LookupMX(name string) ([]*MX, error) {
//line /usr/local/go/src/net/lookup.go:549
	_go_fuzz_dep_.CoverTab[15220]++
						return DefaultResolver.LookupMX(context.Background(), name)
//line /usr/local/go/src/net/lookup.go:550
	// _ = "end of CoverTab[15220]"
}

// LookupMX returns the DNS MX records for the given domain name sorted by preference.
//line /usr/local/go/src/net/lookup.go:553
//
//line /usr/local/go/src/net/lookup.go:553
// The returned mail server names are validated to be properly
//line /usr/local/go/src/net/lookup.go:553
// formatted presentation-format domain names. If the response contains
//line /usr/local/go/src/net/lookup.go:553
// invalid names, those records are filtered out and an error
//line /usr/local/go/src/net/lookup.go:553
// will be returned alongside the remaining results, if any.
//line /usr/local/go/src/net/lookup.go:559
func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error) {
//line /usr/local/go/src/net/lookup.go:559
	_go_fuzz_dep_.CoverTab[15221]++
						records, err := r.lookupMX(ctx, name)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:561
		_go_fuzz_dep_.CoverTab[15225]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:562
		// _ = "end of CoverTab[15225]"
	} else {
//line /usr/local/go/src/net/lookup.go:563
		_go_fuzz_dep_.CoverTab[15226]++
//line /usr/local/go/src/net/lookup.go:563
		// _ = "end of CoverTab[15226]"
//line /usr/local/go/src/net/lookup.go:563
	}
//line /usr/local/go/src/net/lookup.go:563
	// _ = "end of CoverTab[15221]"
//line /usr/local/go/src/net/lookup.go:563
	_go_fuzz_dep_.CoverTab[15222]++
						filteredMX := make([]*MX, 0, len(records))
						for _, mx := range records {
//line /usr/local/go/src/net/lookup.go:565
		_go_fuzz_dep_.CoverTab[15227]++
							if mx == nil {
//line /usr/local/go/src/net/lookup.go:566
			_go_fuzz_dep_.CoverTab[15230]++
								continue
//line /usr/local/go/src/net/lookup.go:567
			// _ = "end of CoverTab[15230]"
		} else {
//line /usr/local/go/src/net/lookup.go:568
			_go_fuzz_dep_.CoverTab[15231]++
//line /usr/local/go/src/net/lookup.go:568
			// _ = "end of CoverTab[15231]"
//line /usr/local/go/src/net/lookup.go:568
		}
//line /usr/local/go/src/net/lookup.go:568
		// _ = "end of CoverTab[15227]"
//line /usr/local/go/src/net/lookup.go:568
		_go_fuzz_dep_.CoverTab[15228]++
							if !isDomainName(mx.Host) {
//line /usr/local/go/src/net/lookup.go:569
			_go_fuzz_dep_.CoverTab[15232]++
								continue
//line /usr/local/go/src/net/lookup.go:570
			// _ = "end of CoverTab[15232]"
		} else {
//line /usr/local/go/src/net/lookup.go:571
			_go_fuzz_dep_.CoverTab[15233]++
//line /usr/local/go/src/net/lookup.go:571
			// _ = "end of CoverTab[15233]"
//line /usr/local/go/src/net/lookup.go:571
		}
//line /usr/local/go/src/net/lookup.go:571
		// _ = "end of CoverTab[15228]"
//line /usr/local/go/src/net/lookup.go:571
		_go_fuzz_dep_.CoverTab[15229]++
							filteredMX = append(filteredMX, mx)
//line /usr/local/go/src/net/lookup.go:572
		// _ = "end of CoverTab[15229]"
	}
//line /usr/local/go/src/net/lookup.go:573
	// _ = "end of CoverTab[15222]"
//line /usr/local/go/src/net/lookup.go:573
	_go_fuzz_dep_.CoverTab[15223]++
						if len(records) != len(filteredMX) {
//line /usr/local/go/src/net/lookup.go:574
		_go_fuzz_dep_.CoverTab[15234]++
							return filteredMX, &DNSError{Err: errMalformedDNSRecordsDetail, Name: name}
//line /usr/local/go/src/net/lookup.go:575
		// _ = "end of CoverTab[15234]"
	} else {
//line /usr/local/go/src/net/lookup.go:576
		_go_fuzz_dep_.CoverTab[15235]++
//line /usr/local/go/src/net/lookup.go:576
		// _ = "end of CoverTab[15235]"
//line /usr/local/go/src/net/lookup.go:576
	}
//line /usr/local/go/src/net/lookup.go:576
	// _ = "end of CoverTab[15223]"
//line /usr/local/go/src/net/lookup.go:576
	_go_fuzz_dep_.CoverTab[15224]++
						return filteredMX, nil
//line /usr/local/go/src/net/lookup.go:577
	// _ = "end of CoverTab[15224]"
}

// LookupNS returns the DNS NS records for the given domain name.
//line /usr/local/go/src/net/lookup.go:580
//
//line /usr/local/go/src/net/lookup.go:580
// The returned name server names are validated to be properly
//line /usr/local/go/src/net/lookup.go:580
// formatted presentation-format domain names. If the response contains
//line /usr/local/go/src/net/lookup.go:580
// invalid names, those records are filtered out and an error
//line /usr/local/go/src/net/lookup.go:580
// will be returned alongside the remaining results, if any.
//line /usr/local/go/src/net/lookup.go:580
//
//line /usr/local/go/src/net/lookup.go:580
// LookupNS uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/lookup.go:580
// Resolver.LookupNS.
//line /usr/local/go/src/net/lookup.go:589
func LookupNS(name string) ([]*NS, error) {
//line /usr/local/go/src/net/lookup.go:589
	_go_fuzz_dep_.CoverTab[15236]++
						return DefaultResolver.LookupNS(context.Background(), name)
//line /usr/local/go/src/net/lookup.go:590
	// _ = "end of CoverTab[15236]"
}

// LookupNS returns the DNS NS records for the given domain name.
//line /usr/local/go/src/net/lookup.go:593
//
//line /usr/local/go/src/net/lookup.go:593
// The returned name server names are validated to be properly
//line /usr/local/go/src/net/lookup.go:593
// formatted presentation-format domain names. If the response contains
//line /usr/local/go/src/net/lookup.go:593
// invalid names, those records are filtered out and an error
//line /usr/local/go/src/net/lookup.go:593
// will be returned alongside the remaining results, if any.
//line /usr/local/go/src/net/lookup.go:599
func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error) {
//line /usr/local/go/src/net/lookup.go:599
	_go_fuzz_dep_.CoverTab[15237]++
						records, err := r.lookupNS(ctx, name)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:601
		_go_fuzz_dep_.CoverTab[15241]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:602
		// _ = "end of CoverTab[15241]"
	} else {
//line /usr/local/go/src/net/lookup.go:603
		_go_fuzz_dep_.CoverTab[15242]++
//line /usr/local/go/src/net/lookup.go:603
		// _ = "end of CoverTab[15242]"
//line /usr/local/go/src/net/lookup.go:603
	}
//line /usr/local/go/src/net/lookup.go:603
	// _ = "end of CoverTab[15237]"
//line /usr/local/go/src/net/lookup.go:603
	_go_fuzz_dep_.CoverTab[15238]++
						filteredNS := make([]*NS, 0, len(records))
						for _, ns := range records {
//line /usr/local/go/src/net/lookup.go:605
		_go_fuzz_dep_.CoverTab[15243]++
							if ns == nil {
//line /usr/local/go/src/net/lookup.go:606
			_go_fuzz_dep_.CoverTab[15246]++
								continue
//line /usr/local/go/src/net/lookup.go:607
			// _ = "end of CoverTab[15246]"
		} else {
//line /usr/local/go/src/net/lookup.go:608
			_go_fuzz_dep_.CoverTab[15247]++
//line /usr/local/go/src/net/lookup.go:608
			// _ = "end of CoverTab[15247]"
//line /usr/local/go/src/net/lookup.go:608
		}
//line /usr/local/go/src/net/lookup.go:608
		// _ = "end of CoverTab[15243]"
//line /usr/local/go/src/net/lookup.go:608
		_go_fuzz_dep_.CoverTab[15244]++
							if !isDomainName(ns.Host) {
//line /usr/local/go/src/net/lookup.go:609
			_go_fuzz_dep_.CoverTab[15248]++
								continue
//line /usr/local/go/src/net/lookup.go:610
			// _ = "end of CoverTab[15248]"
		} else {
//line /usr/local/go/src/net/lookup.go:611
			_go_fuzz_dep_.CoverTab[15249]++
//line /usr/local/go/src/net/lookup.go:611
			// _ = "end of CoverTab[15249]"
//line /usr/local/go/src/net/lookup.go:611
		}
//line /usr/local/go/src/net/lookup.go:611
		// _ = "end of CoverTab[15244]"
//line /usr/local/go/src/net/lookup.go:611
		_go_fuzz_dep_.CoverTab[15245]++
							filteredNS = append(filteredNS, ns)
//line /usr/local/go/src/net/lookup.go:612
		// _ = "end of CoverTab[15245]"
	}
//line /usr/local/go/src/net/lookup.go:613
	// _ = "end of CoverTab[15238]"
//line /usr/local/go/src/net/lookup.go:613
	_go_fuzz_dep_.CoverTab[15239]++
						if len(records) != len(filteredNS) {
//line /usr/local/go/src/net/lookup.go:614
		_go_fuzz_dep_.CoverTab[15250]++
							return filteredNS, &DNSError{Err: errMalformedDNSRecordsDetail, Name: name}
//line /usr/local/go/src/net/lookup.go:615
		// _ = "end of CoverTab[15250]"
	} else {
//line /usr/local/go/src/net/lookup.go:616
		_go_fuzz_dep_.CoverTab[15251]++
//line /usr/local/go/src/net/lookup.go:616
		// _ = "end of CoverTab[15251]"
//line /usr/local/go/src/net/lookup.go:616
	}
//line /usr/local/go/src/net/lookup.go:616
	// _ = "end of CoverTab[15239]"
//line /usr/local/go/src/net/lookup.go:616
	_go_fuzz_dep_.CoverTab[15240]++
						return filteredNS, nil
//line /usr/local/go/src/net/lookup.go:617
	// _ = "end of CoverTab[15240]"
}

// LookupTXT returns the DNS TXT records for the given domain name.
//line /usr/local/go/src/net/lookup.go:620
//
//line /usr/local/go/src/net/lookup.go:620
// LookupTXT uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/lookup.go:620
// Resolver.LookupTXT.
//line /usr/local/go/src/net/lookup.go:624
func LookupTXT(name string) ([]string, error) {
//line /usr/local/go/src/net/lookup.go:624
	_go_fuzz_dep_.CoverTab[15252]++
						return DefaultResolver.lookupTXT(context.Background(), name)
//line /usr/local/go/src/net/lookup.go:625
	// _ = "end of CoverTab[15252]"
}

// LookupTXT returns the DNS TXT records for the given domain name.
func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error) {
//line /usr/local/go/src/net/lookup.go:629
	_go_fuzz_dep_.CoverTab[15253]++
						return r.lookupTXT(ctx, name)
//line /usr/local/go/src/net/lookup.go:630
	// _ = "end of CoverTab[15253]"
}

// LookupAddr performs a reverse lookup for the given address, returning a list
//line /usr/local/go/src/net/lookup.go:633
// of names mapping to that address.
//line /usr/local/go/src/net/lookup.go:633
//
//line /usr/local/go/src/net/lookup.go:633
// The returned names are validated to be properly formatted presentation-format
//line /usr/local/go/src/net/lookup.go:633
// domain names. If the response contains invalid names, those records are filtered
//line /usr/local/go/src/net/lookup.go:633
// out and an error will be returned alongside the remaining results, if any.
//line /usr/local/go/src/net/lookup.go:633
//
//line /usr/local/go/src/net/lookup.go:633
// When using the host C library resolver, at most one result will be
//line /usr/local/go/src/net/lookup.go:633
// returned. To bypass the host resolver, use a custom Resolver.
//line /usr/local/go/src/net/lookup.go:633
//
//line /usr/local/go/src/net/lookup.go:633
// LookupAddr uses context.Background internally; to specify the context, use
//line /usr/local/go/src/net/lookup.go:633
// Resolver.LookupAddr.
//line /usr/local/go/src/net/lookup.go:645
func LookupAddr(addr string) (names []string, err error) {
//line /usr/local/go/src/net/lookup.go:645
	_go_fuzz_dep_.CoverTab[15254]++
						return DefaultResolver.LookupAddr(context.Background(), addr)
//line /usr/local/go/src/net/lookup.go:646
	// _ = "end of CoverTab[15254]"
}

// LookupAddr performs a reverse lookup for the given address, returning a list
//line /usr/local/go/src/net/lookup.go:649
// of names mapping to that address.
//line /usr/local/go/src/net/lookup.go:649
//
//line /usr/local/go/src/net/lookup.go:649
// The returned names are validated to be properly formatted presentation-format
//line /usr/local/go/src/net/lookup.go:649
// domain names. If the response contains invalid names, those records are filtered
//line /usr/local/go/src/net/lookup.go:649
// out and an error will be returned alongside the remaining results, if any.
//line /usr/local/go/src/net/lookup.go:655
func (r *Resolver) LookupAddr(ctx context.Context, addr string) ([]string, error) {
//line /usr/local/go/src/net/lookup.go:655
	_go_fuzz_dep_.CoverTab[15255]++
						names, err := r.lookupAddr(ctx, addr)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:657
		_go_fuzz_dep_.CoverTab[15259]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:658
		// _ = "end of CoverTab[15259]"
	} else {
//line /usr/local/go/src/net/lookup.go:659
		_go_fuzz_dep_.CoverTab[15260]++
//line /usr/local/go/src/net/lookup.go:659
		// _ = "end of CoverTab[15260]"
//line /usr/local/go/src/net/lookup.go:659
	}
//line /usr/local/go/src/net/lookup.go:659
	// _ = "end of CoverTab[15255]"
//line /usr/local/go/src/net/lookup.go:659
	_go_fuzz_dep_.CoverTab[15256]++
						filteredNames := make([]string, 0, len(names))
						for _, name := range names {
//line /usr/local/go/src/net/lookup.go:661
		_go_fuzz_dep_.CoverTab[15261]++
							if isDomainName(name) {
//line /usr/local/go/src/net/lookup.go:662
			_go_fuzz_dep_.CoverTab[15262]++
								filteredNames = append(filteredNames, name)
//line /usr/local/go/src/net/lookup.go:663
			// _ = "end of CoverTab[15262]"
		} else {
//line /usr/local/go/src/net/lookup.go:664
			_go_fuzz_dep_.CoverTab[15263]++
//line /usr/local/go/src/net/lookup.go:664
			// _ = "end of CoverTab[15263]"
//line /usr/local/go/src/net/lookup.go:664
		}
//line /usr/local/go/src/net/lookup.go:664
		// _ = "end of CoverTab[15261]"
	}
//line /usr/local/go/src/net/lookup.go:665
	// _ = "end of CoverTab[15256]"
//line /usr/local/go/src/net/lookup.go:665
	_go_fuzz_dep_.CoverTab[15257]++
						if len(names) != len(filteredNames) {
//line /usr/local/go/src/net/lookup.go:666
		_go_fuzz_dep_.CoverTab[15264]++
							return filteredNames, &DNSError{Err: errMalformedDNSRecordsDetail, Name: addr}
//line /usr/local/go/src/net/lookup.go:667
		// _ = "end of CoverTab[15264]"
	} else {
//line /usr/local/go/src/net/lookup.go:668
		_go_fuzz_dep_.CoverTab[15265]++
//line /usr/local/go/src/net/lookup.go:668
		// _ = "end of CoverTab[15265]"
//line /usr/local/go/src/net/lookup.go:668
	}
//line /usr/local/go/src/net/lookup.go:668
	// _ = "end of CoverTab[15257]"
//line /usr/local/go/src/net/lookup.go:668
	_go_fuzz_dep_.CoverTab[15258]++
						return filteredNames, nil
//line /usr/local/go/src/net/lookup.go:669
	// _ = "end of CoverTab[15258]"
}

// errMalformedDNSRecordsDetail is the DNSError detail which is returned when a Resolver.Lookup...
//line /usr/local/go/src/net/lookup.go:672
// method receives DNS records which contain invalid DNS names. This may be returned alongside
//line /usr/local/go/src/net/lookup.go:672
// results which have had the malformed records filtered out.
//line /usr/local/go/src/net/lookup.go:675
var errMalformedDNSRecordsDetail = "DNS response contained records which contain invalid names"

// dial makes a new connection to the provided server (which must be
//line /usr/local/go/src/net/lookup.go:677
// an IP address) with the provided network type, using either r.Dial
//line /usr/local/go/src/net/lookup.go:677
// (if both r and r.Dial are non-nil) or else Dialer.DialContext.
//line /usr/local/go/src/net/lookup.go:680
func (r *Resolver) dial(ctx context.Context, network, server string) (Conn, error) {
//line /usr/local/go/src/net/lookup.go:680
	_go_fuzz_dep_.CoverTab[15266]++
	// Calling Dial here is scary -- we have to be sure not to
	// dial a name that will require a DNS lookup, or Dial will
	// call back here to translate it. The DNS config parser has
	// already checked that all the cfg.servers are IP
	// addresses, which Dial will use without a DNS lookup.
	var c Conn
	var err error
	if r != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:688
		_go_fuzz_dep_.CoverTab[15269]++
//line /usr/local/go/src/net/lookup.go:688
		return r.Dial != nil
//line /usr/local/go/src/net/lookup.go:688
		// _ = "end of CoverTab[15269]"
//line /usr/local/go/src/net/lookup.go:688
	}() {
//line /usr/local/go/src/net/lookup.go:688
		_go_fuzz_dep_.CoverTab[15270]++
							c, err = r.Dial(ctx, network, server)
//line /usr/local/go/src/net/lookup.go:689
		// _ = "end of CoverTab[15270]"
	} else {
//line /usr/local/go/src/net/lookup.go:690
		_go_fuzz_dep_.CoverTab[15271]++
							var d Dialer
							c, err = d.DialContext(ctx, network, server)
//line /usr/local/go/src/net/lookup.go:692
		// _ = "end of CoverTab[15271]"
	}
//line /usr/local/go/src/net/lookup.go:693
	// _ = "end of CoverTab[15266]"
//line /usr/local/go/src/net/lookup.go:693
	_go_fuzz_dep_.CoverTab[15267]++
						if err != nil {
//line /usr/local/go/src/net/lookup.go:694
		_go_fuzz_dep_.CoverTab[15272]++
							return nil, mapErr(err)
//line /usr/local/go/src/net/lookup.go:695
		// _ = "end of CoverTab[15272]"
	} else {
//line /usr/local/go/src/net/lookup.go:696
		_go_fuzz_dep_.CoverTab[15273]++
//line /usr/local/go/src/net/lookup.go:696
		// _ = "end of CoverTab[15273]"
//line /usr/local/go/src/net/lookup.go:696
	}
//line /usr/local/go/src/net/lookup.go:696
	// _ = "end of CoverTab[15267]"
//line /usr/local/go/src/net/lookup.go:696
	_go_fuzz_dep_.CoverTab[15268]++
						return c, nil
//line /usr/local/go/src/net/lookup.go:697
	// _ = "end of CoverTab[15268]"
}

// goLookupSRV returns the SRV records for a target name, built either
//line /usr/local/go/src/net/lookup.go:700
// from its component service ("sip"), protocol ("tcp"), and name
//line /usr/local/go/src/net/lookup.go:700
// ("example.com."), or from name directly (if service and proto are
//line /usr/local/go/src/net/lookup.go:700
// both empty).
//line /usr/local/go/src/net/lookup.go:700
//
//line /usr/local/go/src/net/lookup.go:700
// In either case, the returned target name ("_sip._tcp.example.com.")
//line /usr/local/go/src/net/lookup.go:700
// is also returned on success.
//line /usr/local/go/src/net/lookup.go:700
//
//line /usr/local/go/src/net/lookup.go:700
// The records are sorted by weight.
//line /usr/local/go/src/net/lookup.go:709
func (r *Resolver) goLookupSRV(ctx context.Context, service, proto, name string) (target string, srvs []*SRV, err error) {
//line /usr/local/go/src/net/lookup.go:709
	_go_fuzz_dep_.CoverTab[15274]++
						if service == "" && func() bool {
//line /usr/local/go/src/net/lookup.go:710
		_go_fuzz_dep_.CoverTab[15278]++
//line /usr/local/go/src/net/lookup.go:710
		return proto == ""
//line /usr/local/go/src/net/lookup.go:710
		// _ = "end of CoverTab[15278]"
//line /usr/local/go/src/net/lookup.go:710
	}() {
//line /usr/local/go/src/net/lookup.go:710
		_go_fuzz_dep_.CoverTab[15279]++
							target = name
//line /usr/local/go/src/net/lookup.go:711
		// _ = "end of CoverTab[15279]"
	} else {
//line /usr/local/go/src/net/lookup.go:712
		_go_fuzz_dep_.CoverTab[15280]++
							target = "_" + service + "._" + proto + "." + name
//line /usr/local/go/src/net/lookup.go:713
		// _ = "end of CoverTab[15280]"
	}
//line /usr/local/go/src/net/lookup.go:714
	// _ = "end of CoverTab[15274]"
//line /usr/local/go/src/net/lookup.go:714
	_go_fuzz_dep_.CoverTab[15275]++
						p, server, err := r.lookup(ctx, target, dnsmessage.TypeSRV, nil)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:716
		_go_fuzz_dep_.CoverTab[15281]++
							return "", nil, err
//line /usr/local/go/src/net/lookup.go:717
		// _ = "end of CoverTab[15281]"
	} else {
//line /usr/local/go/src/net/lookup.go:718
		_go_fuzz_dep_.CoverTab[15282]++
//line /usr/local/go/src/net/lookup.go:718
		// _ = "end of CoverTab[15282]"
//line /usr/local/go/src/net/lookup.go:718
	}
//line /usr/local/go/src/net/lookup.go:718
	// _ = "end of CoverTab[15275]"
//line /usr/local/go/src/net/lookup.go:718
	_go_fuzz_dep_.CoverTab[15276]++
						var cname dnsmessage.Name
						for {
//line /usr/local/go/src/net/lookup.go:720
		_go_fuzz_dep_.CoverTab[15283]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/lookup.go:722
			_go_fuzz_dep_.CoverTab[15289]++
								break
//line /usr/local/go/src/net/lookup.go:723
			// _ = "end of CoverTab[15289]"
		} else {
//line /usr/local/go/src/net/lookup.go:724
			_go_fuzz_dep_.CoverTab[15290]++
//line /usr/local/go/src/net/lookup.go:724
			// _ = "end of CoverTab[15290]"
//line /usr/local/go/src/net/lookup.go:724
		}
//line /usr/local/go/src/net/lookup.go:724
		// _ = "end of CoverTab[15283]"
//line /usr/local/go/src/net/lookup.go:724
		_go_fuzz_dep_.CoverTab[15284]++
							if err != nil {
//line /usr/local/go/src/net/lookup.go:725
			_go_fuzz_dep_.CoverTab[15291]++
								return "", nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:730
			// _ = "end of CoverTab[15291]"
		} else {
//line /usr/local/go/src/net/lookup.go:731
			_go_fuzz_dep_.CoverTab[15292]++
//line /usr/local/go/src/net/lookup.go:731
			// _ = "end of CoverTab[15292]"
//line /usr/local/go/src/net/lookup.go:731
		}
//line /usr/local/go/src/net/lookup.go:731
		// _ = "end of CoverTab[15284]"
//line /usr/local/go/src/net/lookup.go:731
		_go_fuzz_dep_.CoverTab[15285]++
							if h.Type != dnsmessage.TypeSRV {
//line /usr/local/go/src/net/lookup.go:732
			_go_fuzz_dep_.CoverTab[15293]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/lookup.go:733
				_go_fuzz_dep_.CoverTab[15295]++
									return "", nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /usr/local/go/src/net/lookup.go:738
				// _ = "end of CoverTab[15295]"
			} else {
//line /usr/local/go/src/net/lookup.go:739
				_go_fuzz_dep_.CoverTab[15296]++
//line /usr/local/go/src/net/lookup.go:739
				// _ = "end of CoverTab[15296]"
//line /usr/local/go/src/net/lookup.go:739
			}
//line /usr/local/go/src/net/lookup.go:739
			// _ = "end of CoverTab[15293]"
//line /usr/local/go/src/net/lookup.go:739
			_go_fuzz_dep_.CoverTab[15294]++
								continue
//line /usr/local/go/src/net/lookup.go:740
			// _ = "end of CoverTab[15294]"
		} else {
//line /usr/local/go/src/net/lookup.go:741
			_go_fuzz_dep_.CoverTab[15297]++
//line /usr/local/go/src/net/lookup.go:741
			// _ = "end of CoverTab[15297]"
//line /usr/local/go/src/net/lookup.go:741
		}
//line /usr/local/go/src/net/lookup.go:741
		// _ = "end of CoverTab[15285]"
//line /usr/local/go/src/net/lookup.go:741
		_go_fuzz_dep_.CoverTab[15286]++
							if cname.Length == 0 && func() bool {
//line /usr/local/go/src/net/lookup.go:742
			_go_fuzz_dep_.CoverTab[15298]++
//line /usr/local/go/src/net/lookup.go:742
			return h.Name.Length != 0
//line /usr/local/go/src/net/lookup.go:742
			// _ = "end of CoverTab[15298]"
//line /usr/local/go/src/net/lookup.go:742
		}() {
//line /usr/local/go/src/net/lookup.go:742
			_go_fuzz_dep_.CoverTab[15299]++
								cname = h.Name
//line /usr/local/go/src/net/lookup.go:743
			// _ = "end of CoverTab[15299]"
		} else {
//line /usr/local/go/src/net/lookup.go:744
			_go_fuzz_dep_.CoverTab[15300]++
//line /usr/local/go/src/net/lookup.go:744
			// _ = "end of CoverTab[15300]"
//line /usr/local/go/src/net/lookup.go:744
		}
//line /usr/local/go/src/net/lookup.go:744
		// _ = "end of CoverTab[15286]"
//line /usr/local/go/src/net/lookup.go:744
		_go_fuzz_dep_.CoverTab[15287]++
							srv, err := p.SRVResource()
							if err != nil {
//line /usr/local/go/src/net/lookup.go:746
			_go_fuzz_dep_.CoverTab[15301]++
								return "", nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:751
			// _ = "end of CoverTab[15301]"
		} else {
//line /usr/local/go/src/net/lookup.go:752
			_go_fuzz_dep_.CoverTab[15302]++
//line /usr/local/go/src/net/lookup.go:752
			// _ = "end of CoverTab[15302]"
//line /usr/local/go/src/net/lookup.go:752
		}
//line /usr/local/go/src/net/lookup.go:752
		// _ = "end of CoverTab[15287]"
//line /usr/local/go/src/net/lookup.go:752
		_go_fuzz_dep_.CoverTab[15288]++
							srvs = append(srvs, &SRV{Target: srv.Target.String(), Port: srv.Port, Priority: srv.Priority, Weight: srv.Weight})
//line /usr/local/go/src/net/lookup.go:753
		// _ = "end of CoverTab[15288]"
	}
//line /usr/local/go/src/net/lookup.go:754
	// _ = "end of CoverTab[15276]"
//line /usr/local/go/src/net/lookup.go:754
	_go_fuzz_dep_.CoverTab[15277]++
						byPriorityWeight(srvs).sort()
						return cname.String(), srvs, nil
//line /usr/local/go/src/net/lookup.go:756
	// _ = "end of CoverTab[15277]"
}

// goLookupMX returns the MX records for name.
func (r *Resolver) goLookupMX(ctx context.Context, name string) ([]*MX, error) {
//line /usr/local/go/src/net/lookup.go:760
	_go_fuzz_dep_.CoverTab[15303]++
						p, server, err := r.lookup(ctx, name, dnsmessage.TypeMX, nil)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:762
		_go_fuzz_dep_.CoverTab[15306]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:763
		// _ = "end of CoverTab[15306]"
	} else {
//line /usr/local/go/src/net/lookup.go:764
		_go_fuzz_dep_.CoverTab[15307]++
//line /usr/local/go/src/net/lookup.go:764
		// _ = "end of CoverTab[15307]"
//line /usr/local/go/src/net/lookup.go:764
	}
//line /usr/local/go/src/net/lookup.go:764
	// _ = "end of CoverTab[15303]"
//line /usr/local/go/src/net/lookup.go:764
	_go_fuzz_dep_.CoverTab[15304]++
						var mxs []*MX
						for {
//line /usr/local/go/src/net/lookup.go:766
		_go_fuzz_dep_.CoverTab[15308]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/lookup.go:768
			_go_fuzz_dep_.CoverTab[15313]++
								break
//line /usr/local/go/src/net/lookup.go:769
			// _ = "end of CoverTab[15313]"
		} else {
//line /usr/local/go/src/net/lookup.go:770
			_go_fuzz_dep_.CoverTab[15314]++
//line /usr/local/go/src/net/lookup.go:770
			// _ = "end of CoverTab[15314]"
//line /usr/local/go/src/net/lookup.go:770
		}
//line /usr/local/go/src/net/lookup.go:770
		// _ = "end of CoverTab[15308]"
//line /usr/local/go/src/net/lookup.go:770
		_go_fuzz_dep_.CoverTab[15309]++
							if err != nil {
//line /usr/local/go/src/net/lookup.go:771
			_go_fuzz_dep_.CoverTab[15315]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:776
			// _ = "end of CoverTab[15315]"
		} else {
//line /usr/local/go/src/net/lookup.go:777
			_go_fuzz_dep_.CoverTab[15316]++
//line /usr/local/go/src/net/lookup.go:777
			// _ = "end of CoverTab[15316]"
//line /usr/local/go/src/net/lookup.go:777
		}
//line /usr/local/go/src/net/lookup.go:777
		// _ = "end of CoverTab[15309]"
//line /usr/local/go/src/net/lookup.go:777
		_go_fuzz_dep_.CoverTab[15310]++
							if h.Type != dnsmessage.TypeMX {
//line /usr/local/go/src/net/lookup.go:778
			_go_fuzz_dep_.CoverTab[15317]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/lookup.go:779
				_go_fuzz_dep_.CoverTab[15319]++
									return nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /usr/local/go/src/net/lookup.go:784
				// _ = "end of CoverTab[15319]"
			} else {
//line /usr/local/go/src/net/lookup.go:785
				_go_fuzz_dep_.CoverTab[15320]++
//line /usr/local/go/src/net/lookup.go:785
				// _ = "end of CoverTab[15320]"
//line /usr/local/go/src/net/lookup.go:785
			}
//line /usr/local/go/src/net/lookup.go:785
			// _ = "end of CoverTab[15317]"
//line /usr/local/go/src/net/lookup.go:785
			_go_fuzz_dep_.CoverTab[15318]++
								continue
//line /usr/local/go/src/net/lookup.go:786
			// _ = "end of CoverTab[15318]"
		} else {
//line /usr/local/go/src/net/lookup.go:787
			_go_fuzz_dep_.CoverTab[15321]++
//line /usr/local/go/src/net/lookup.go:787
			// _ = "end of CoverTab[15321]"
//line /usr/local/go/src/net/lookup.go:787
		}
//line /usr/local/go/src/net/lookup.go:787
		// _ = "end of CoverTab[15310]"
//line /usr/local/go/src/net/lookup.go:787
		_go_fuzz_dep_.CoverTab[15311]++
							mx, err := p.MXResource()
							if err != nil {
//line /usr/local/go/src/net/lookup.go:789
			_go_fuzz_dep_.CoverTab[15322]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:794
			// _ = "end of CoverTab[15322]"
		} else {
//line /usr/local/go/src/net/lookup.go:795
			_go_fuzz_dep_.CoverTab[15323]++
//line /usr/local/go/src/net/lookup.go:795
			// _ = "end of CoverTab[15323]"
//line /usr/local/go/src/net/lookup.go:795
		}
//line /usr/local/go/src/net/lookup.go:795
		// _ = "end of CoverTab[15311]"
//line /usr/local/go/src/net/lookup.go:795
		_go_fuzz_dep_.CoverTab[15312]++
							mxs = append(mxs, &MX{Host: mx.MX.String(), Pref: mx.Pref})
//line /usr/local/go/src/net/lookup.go:796
		// _ = "end of CoverTab[15312]"

	}
//line /usr/local/go/src/net/lookup.go:798
	// _ = "end of CoverTab[15304]"
//line /usr/local/go/src/net/lookup.go:798
	_go_fuzz_dep_.CoverTab[15305]++
						byPref(mxs).sort()
						return mxs, nil
//line /usr/local/go/src/net/lookup.go:800
	// _ = "end of CoverTab[15305]"
}

// goLookupNS returns the NS records for name.
func (r *Resolver) goLookupNS(ctx context.Context, name string) ([]*NS, error) {
//line /usr/local/go/src/net/lookup.go:804
	_go_fuzz_dep_.CoverTab[15324]++
						p, server, err := r.lookup(ctx, name, dnsmessage.TypeNS, nil)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:806
		_go_fuzz_dep_.CoverTab[15327]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:807
		// _ = "end of CoverTab[15327]"
	} else {
//line /usr/local/go/src/net/lookup.go:808
		_go_fuzz_dep_.CoverTab[15328]++
//line /usr/local/go/src/net/lookup.go:808
		// _ = "end of CoverTab[15328]"
//line /usr/local/go/src/net/lookup.go:808
	}
//line /usr/local/go/src/net/lookup.go:808
	// _ = "end of CoverTab[15324]"
//line /usr/local/go/src/net/lookup.go:808
	_go_fuzz_dep_.CoverTab[15325]++
						var nss []*NS
						for {
//line /usr/local/go/src/net/lookup.go:810
		_go_fuzz_dep_.CoverTab[15329]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/lookup.go:812
			_go_fuzz_dep_.CoverTab[15334]++
								break
//line /usr/local/go/src/net/lookup.go:813
			// _ = "end of CoverTab[15334]"
		} else {
//line /usr/local/go/src/net/lookup.go:814
			_go_fuzz_dep_.CoverTab[15335]++
//line /usr/local/go/src/net/lookup.go:814
			// _ = "end of CoverTab[15335]"
//line /usr/local/go/src/net/lookup.go:814
		}
//line /usr/local/go/src/net/lookup.go:814
		// _ = "end of CoverTab[15329]"
//line /usr/local/go/src/net/lookup.go:814
		_go_fuzz_dep_.CoverTab[15330]++
							if err != nil {
//line /usr/local/go/src/net/lookup.go:815
			_go_fuzz_dep_.CoverTab[15336]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:820
			// _ = "end of CoverTab[15336]"
		} else {
//line /usr/local/go/src/net/lookup.go:821
			_go_fuzz_dep_.CoverTab[15337]++
//line /usr/local/go/src/net/lookup.go:821
			// _ = "end of CoverTab[15337]"
//line /usr/local/go/src/net/lookup.go:821
		}
//line /usr/local/go/src/net/lookup.go:821
		// _ = "end of CoverTab[15330]"
//line /usr/local/go/src/net/lookup.go:821
		_go_fuzz_dep_.CoverTab[15331]++
							if h.Type != dnsmessage.TypeNS {
//line /usr/local/go/src/net/lookup.go:822
			_go_fuzz_dep_.CoverTab[15338]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/lookup.go:823
				_go_fuzz_dep_.CoverTab[15340]++
									return nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /usr/local/go/src/net/lookup.go:828
				// _ = "end of CoverTab[15340]"
			} else {
//line /usr/local/go/src/net/lookup.go:829
				_go_fuzz_dep_.CoverTab[15341]++
//line /usr/local/go/src/net/lookup.go:829
				// _ = "end of CoverTab[15341]"
//line /usr/local/go/src/net/lookup.go:829
			}
//line /usr/local/go/src/net/lookup.go:829
			// _ = "end of CoverTab[15338]"
//line /usr/local/go/src/net/lookup.go:829
			_go_fuzz_dep_.CoverTab[15339]++
								continue
//line /usr/local/go/src/net/lookup.go:830
			// _ = "end of CoverTab[15339]"
		} else {
//line /usr/local/go/src/net/lookup.go:831
			_go_fuzz_dep_.CoverTab[15342]++
//line /usr/local/go/src/net/lookup.go:831
			// _ = "end of CoverTab[15342]"
//line /usr/local/go/src/net/lookup.go:831
		}
//line /usr/local/go/src/net/lookup.go:831
		// _ = "end of CoverTab[15331]"
//line /usr/local/go/src/net/lookup.go:831
		_go_fuzz_dep_.CoverTab[15332]++
							ns, err := p.NSResource()
							if err != nil {
//line /usr/local/go/src/net/lookup.go:833
			_go_fuzz_dep_.CoverTab[15343]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:838
			// _ = "end of CoverTab[15343]"
		} else {
//line /usr/local/go/src/net/lookup.go:839
			_go_fuzz_dep_.CoverTab[15344]++
//line /usr/local/go/src/net/lookup.go:839
			// _ = "end of CoverTab[15344]"
//line /usr/local/go/src/net/lookup.go:839
		}
//line /usr/local/go/src/net/lookup.go:839
		// _ = "end of CoverTab[15332]"
//line /usr/local/go/src/net/lookup.go:839
		_go_fuzz_dep_.CoverTab[15333]++
							nss = append(nss, &NS{Host: ns.NS.String()})
//line /usr/local/go/src/net/lookup.go:840
		// _ = "end of CoverTab[15333]"
	}
//line /usr/local/go/src/net/lookup.go:841
	// _ = "end of CoverTab[15325]"
//line /usr/local/go/src/net/lookup.go:841
	_go_fuzz_dep_.CoverTab[15326]++
						return nss, nil
//line /usr/local/go/src/net/lookup.go:842
	// _ = "end of CoverTab[15326]"
}

// goLookupTXT returns the TXT records from name.
func (r *Resolver) goLookupTXT(ctx context.Context, name string) ([]string, error) {
//line /usr/local/go/src/net/lookup.go:846
	_go_fuzz_dep_.CoverTab[15345]++
						p, server, err := r.lookup(ctx, name, dnsmessage.TypeTXT, nil)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:848
		_go_fuzz_dep_.CoverTab[15348]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:849
		// _ = "end of CoverTab[15348]"
	} else {
//line /usr/local/go/src/net/lookup.go:850
		_go_fuzz_dep_.CoverTab[15349]++
//line /usr/local/go/src/net/lookup.go:850
		// _ = "end of CoverTab[15349]"
//line /usr/local/go/src/net/lookup.go:850
	}
//line /usr/local/go/src/net/lookup.go:850
	// _ = "end of CoverTab[15345]"
//line /usr/local/go/src/net/lookup.go:850
	_go_fuzz_dep_.CoverTab[15346]++
						var txts []string
						for {
//line /usr/local/go/src/net/lookup.go:852
		_go_fuzz_dep_.CoverTab[15350]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/lookup.go:854
			_go_fuzz_dep_.CoverTab[15358]++
								break
//line /usr/local/go/src/net/lookup.go:855
			// _ = "end of CoverTab[15358]"
		} else {
//line /usr/local/go/src/net/lookup.go:856
			_go_fuzz_dep_.CoverTab[15359]++
//line /usr/local/go/src/net/lookup.go:856
			// _ = "end of CoverTab[15359]"
//line /usr/local/go/src/net/lookup.go:856
		}
//line /usr/local/go/src/net/lookup.go:856
		// _ = "end of CoverTab[15350]"
//line /usr/local/go/src/net/lookup.go:856
		_go_fuzz_dep_.CoverTab[15351]++
							if err != nil {
//line /usr/local/go/src/net/lookup.go:857
			_go_fuzz_dep_.CoverTab[15360]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:862
			// _ = "end of CoverTab[15360]"
		} else {
//line /usr/local/go/src/net/lookup.go:863
			_go_fuzz_dep_.CoverTab[15361]++
//line /usr/local/go/src/net/lookup.go:863
			// _ = "end of CoverTab[15361]"
//line /usr/local/go/src/net/lookup.go:863
		}
//line /usr/local/go/src/net/lookup.go:863
		// _ = "end of CoverTab[15351]"
//line /usr/local/go/src/net/lookup.go:863
		_go_fuzz_dep_.CoverTab[15352]++
							if h.Type != dnsmessage.TypeTXT {
//line /usr/local/go/src/net/lookup.go:864
			_go_fuzz_dep_.CoverTab[15362]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/lookup.go:865
				_go_fuzz_dep_.CoverTab[15364]++
									return nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /usr/local/go/src/net/lookup.go:870
				// _ = "end of CoverTab[15364]"
			} else {
//line /usr/local/go/src/net/lookup.go:871
				_go_fuzz_dep_.CoverTab[15365]++
//line /usr/local/go/src/net/lookup.go:871
				// _ = "end of CoverTab[15365]"
//line /usr/local/go/src/net/lookup.go:871
			}
//line /usr/local/go/src/net/lookup.go:871
			// _ = "end of CoverTab[15362]"
//line /usr/local/go/src/net/lookup.go:871
			_go_fuzz_dep_.CoverTab[15363]++
								continue
//line /usr/local/go/src/net/lookup.go:872
			// _ = "end of CoverTab[15363]"
		} else {
//line /usr/local/go/src/net/lookup.go:873
			_go_fuzz_dep_.CoverTab[15366]++
//line /usr/local/go/src/net/lookup.go:873
			// _ = "end of CoverTab[15366]"
//line /usr/local/go/src/net/lookup.go:873
		}
//line /usr/local/go/src/net/lookup.go:873
		// _ = "end of CoverTab[15352]"
//line /usr/local/go/src/net/lookup.go:873
		_go_fuzz_dep_.CoverTab[15353]++
							txt, err := p.TXTResource()
							if err != nil {
//line /usr/local/go/src/net/lookup.go:875
			_go_fuzz_dep_.CoverTab[15367]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:880
			// _ = "end of CoverTab[15367]"
		} else {
//line /usr/local/go/src/net/lookup.go:881
			_go_fuzz_dep_.CoverTab[15368]++
//line /usr/local/go/src/net/lookup.go:881
			// _ = "end of CoverTab[15368]"
//line /usr/local/go/src/net/lookup.go:881
		}
//line /usr/local/go/src/net/lookup.go:881
		// _ = "end of CoverTab[15353]"
//line /usr/local/go/src/net/lookup.go:881
		_go_fuzz_dep_.CoverTab[15354]++

//line /usr/local/go/src/net/lookup.go:885
		n := 0
		for _, s := range txt.TXT {
//line /usr/local/go/src/net/lookup.go:886
			_go_fuzz_dep_.CoverTab[15369]++
								n += len(s)
//line /usr/local/go/src/net/lookup.go:887
			// _ = "end of CoverTab[15369]"
		}
//line /usr/local/go/src/net/lookup.go:888
		// _ = "end of CoverTab[15354]"
//line /usr/local/go/src/net/lookup.go:888
		_go_fuzz_dep_.CoverTab[15355]++
							txtJoin := make([]byte, 0, n)
							for _, s := range txt.TXT {
//line /usr/local/go/src/net/lookup.go:890
			_go_fuzz_dep_.CoverTab[15370]++
								txtJoin = append(txtJoin, s...)
//line /usr/local/go/src/net/lookup.go:891
			// _ = "end of CoverTab[15370]"
		}
//line /usr/local/go/src/net/lookup.go:892
		// _ = "end of CoverTab[15355]"
//line /usr/local/go/src/net/lookup.go:892
		_go_fuzz_dep_.CoverTab[15356]++
							if len(txts) == 0 {
//line /usr/local/go/src/net/lookup.go:893
			_go_fuzz_dep_.CoverTab[15371]++
								txts = make([]string, 0, 1)
//line /usr/local/go/src/net/lookup.go:894
			// _ = "end of CoverTab[15371]"
		} else {
//line /usr/local/go/src/net/lookup.go:895
			_go_fuzz_dep_.CoverTab[15372]++
//line /usr/local/go/src/net/lookup.go:895
			// _ = "end of CoverTab[15372]"
//line /usr/local/go/src/net/lookup.go:895
		}
//line /usr/local/go/src/net/lookup.go:895
		// _ = "end of CoverTab[15356]"
//line /usr/local/go/src/net/lookup.go:895
		_go_fuzz_dep_.CoverTab[15357]++
							txts = append(txts, string(txtJoin))
//line /usr/local/go/src/net/lookup.go:896
		// _ = "end of CoverTab[15357]"
	}
//line /usr/local/go/src/net/lookup.go:897
	// _ = "end of CoverTab[15346]"
//line /usr/local/go/src/net/lookup.go:897
	_go_fuzz_dep_.CoverTab[15347]++
						return txts, nil
//line /usr/local/go/src/net/lookup.go:898
	// _ = "end of CoverTab[15347]"
}

func parseCNAMEFromResources(resources []dnsmessage.Resource) (string, error) {
//line /usr/local/go/src/net/lookup.go:901
	_go_fuzz_dep_.CoverTab[15373]++
						if len(resources) == 0 {
//line /usr/local/go/src/net/lookup.go:902
		_go_fuzz_dep_.CoverTab[15376]++
							return "", errors.New("no CNAME record received")
//line /usr/local/go/src/net/lookup.go:903
		// _ = "end of CoverTab[15376]"
	} else {
//line /usr/local/go/src/net/lookup.go:904
		_go_fuzz_dep_.CoverTab[15377]++
//line /usr/local/go/src/net/lookup.go:904
		// _ = "end of CoverTab[15377]"
//line /usr/local/go/src/net/lookup.go:904
	}
//line /usr/local/go/src/net/lookup.go:904
	// _ = "end of CoverTab[15373]"
//line /usr/local/go/src/net/lookup.go:904
	_go_fuzz_dep_.CoverTab[15374]++
						c, ok := resources[0].Body.(*dnsmessage.CNAMEResource)
						if !ok {
//line /usr/local/go/src/net/lookup.go:906
		_go_fuzz_dep_.CoverTab[15378]++
							return "", errors.New("could not parse CNAME record")
//line /usr/local/go/src/net/lookup.go:907
		// _ = "end of CoverTab[15378]"
	} else {
//line /usr/local/go/src/net/lookup.go:908
		_go_fuzz_dep_.CoverTab[15379]++
//line /usr/local/go/src/net/lookup.go:908
		// _ = "end of CoverTab[15379]"
//line /usr/local/go/src/net/lookup.go:908
	}
//line /usr/local/go/src/net/lookup.go:908
	// _ = "end of CoverTab[15374]"
//line /usr/local/go/src/net/lookup.go:908
	_go_fuzz_dep_.CoverTab[15375]++
						return c.CNAME.String(), nil
//line /usr/local/go/src/net/lookup.go:909
	// _ = "end of CoverTab[15375]"
}

//line /usr/local/go/src/net/lookup.go:910
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/lookup.go:910
var _ = _go_fuzz_dep_.CoverTab
