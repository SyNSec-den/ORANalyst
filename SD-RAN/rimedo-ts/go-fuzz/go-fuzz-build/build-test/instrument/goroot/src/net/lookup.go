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
	_go_fuzz_dep_.CoverTab[6657]++
						var lowerProtocol [maxProtoLength]byte
						n := copy(lowerProtocol[:], name)
						lowerASCIIBytes(lowerProtocol[:n])
						proto, found := protocols[string(lowerProtocol[:n])]
						if !found || func() bool {
//line /usr/local/go/src/net/lookup.go:71
		_go_fuzz_dep_.CoverTab[6659]++
//line /usr/local/go/src/net/lookup.go:71
		return n != len(name)
//line /usr/local/go/src/net/lookup.go:71
		// _ = "end of CoverTab[6659]"
//line /usr/local/go/src/net/lookup.go:71
	}() {
//line /usr/local/go/src/net/lookup.go:71
		_go_fuzz_dep_.CoverTab[6660]++
							return 0, &AddrError{Err: "unknown IP protocol specified", Addr: name}
//line /usr/local/go/src/net/lookup.go:72
		// _ = "end of CoverTab[6660]"
	} else {
//line /usr/local/go/src/net/lookup.go:73
		_go_fuzz_dep_.CoverTab[6661]++
//line /usr/local/go/src/net/lookup.go:73
		// _ = "end of CoverTab[6661]"
//line /usr/local/go/src/net/lookup.go:73
	}
//line /usr/local/go/src/net/lookup.go:73
	// _ = "end of CoverTab[6657]"
//line /usr/local/go/src/net/lookup.go:73
	_go_fuzz_dep_.CoverTab[6658]++
						return proto, nil
//line /usr/local/go/src/net/lookup.go:74
	// _ = "end of CoverTab[6658]"
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
	_go_fuzz_dep_.CoverTab[6662]++
						switch network {
	case "tcp4", "tcp6":
//line /usr/local/go/src/net/lookup.go:86
		_go_fuzz_dep_.CoverTab[6665]++
							network = "tcp"
//line /usr/local/go/src/net/lookup.go:87
		// _ = "end of CoverTab[6665]"
	case "udp4", "udp6":
//line /usr/local/go/src/net/lookup.go:88
		_go_fuzz_dep_.CoverTab[6666]++
							network = "udp"
//line /usr/local/go/src/net/lookup.go:89
		// _ = "end of CoverTab[6666]"
//line /usr/local/go/src/net/lookup.go:89
	default:
//line /usr/local/go/src/net/lookup.go:89
		_go_fuzz_dep_.CoverTab[6667]++
//line /usr/local/go/src/net/lookup.go:89
		// _ = "end of CoverTab[6667]"
	}
//line /usr/local/go/src/net/lookup.go:90
	// _ = "end of CoverTab[6662]"
//line /usr/local/go/src/net/lookup.go:90
	_go_fuzz_dep_.CoverTab[6663]++

						if m, ok := services[network]; ok {
//line /usr/local/go/src/net/lookup.go:92
		_go_fuzz_dep_.CoverTab[6668]++
							var lowerService [maxPortBufSize]byte
							n := copy(lowerService[:], service)
							lowerASCIIBytes(lowerService[:n])
							if port, ok := m[string(lowerService[:n])]; ok && func() bool {
//line /usr/local/go/src/net/lookup.go:96
			_go_fuzz_dep_.CoverTab[6669]++
//line /usr/local/go/src/net/lookup.go:96
			return n == len(service)
//line /usr/local/go/src/net/lookup.go:96
			// _ = "end of CoverTab[6669]"
//line /usr/local/go/src/net/lookup.go:96
		}() {
//line /usr/local/go/src/net/lookup.go:96
			_go_fuzz_dep_.CoverTab[6670]++
								return port, nil
//line /usr/local/go/src/net/lookup.go:97
			// _ = "end of CoverTab[6670]"
		} else {
//line /usr/local/go/src/net/lookup.go:98
			_go_fuzz_dep_.CoverTab[6671]++
//line /usr/local/go/src/net/lookup.go:98
			// _ = "end of CoverTab[6671]"
//line /usr/local/go/src/net/lookup.go:98
		}
//line /usr/local/go/src/net/lookup.go:98
		// _ = "end of CoverTab[6668]"
	} else {
//line /usr/local/go/src/net/lookup.go:99
		_go_fuzz_dep_.CoverTab[6672]++
//line /usr/local/go/src/net/lookup.go:99
		// _ = "end of CoverTab[6672]"
//line /usr/local/go/src/net/lookup.go:99
	}
//line /usr/local/go/src/net/lookup.go:99
	// _ = "end of CoverTab[6663]"
//line /usr/local/go/src/net/lookup.go:99
	_go_fuzz_dep_.CoverTab[6664]++
						return 0, &AddrError{Err: "unknown port", Addr: network + "/" + service}
//line /usr/local/go/src/net/lookup.go:100
	// _ = "end of CoverTab[6664]"
}

// ipVersion returns the provided network's IP version: '4', '6' or 0
//line /usr/local/go/src/net/lookup.go:103
// if network does not end in a '4' or '6' byte.
//line /usr/local/go/src/net/lookup.go:105
func ipVersion(network string) byte {
//line /usr/local/go/src/net/lookup.go:105
	_go_fuzz_dep_.CoverTab[6673]++
						if network == "" {
//line /usr/local/go/src/net/lookup.go:106
		_go_fuzz_dep_.CoverTab[6676]++
							return 0
//line /usr/local/go/src/net/lookup.go:107
		// _ = "end of CoverTab[6676]"
	} else {
//line /usr/local/go/src/net/lookup.go:108
		_go_fuzz_dep_.CoverTab[6677]++
//line /usr/local/go/src/net/lookup.go:108
		// _ = "end of CoverTab[6677]"
//line /usr/local/go/src/net/lookup.go:108
	}
//line /usr/local/go/src/net/lookup.go:108
	// _ = "end of CoverTab[6673]"
//line /usr/local/go/src/net/lookup.go:108
	_go_fuzz_dep_.CoverTab[6674]++
						n := network[len(network)-1]
						if n != '4' && func() bool {
//line /usr/local/go/src/net/lookup.go:110
		_go_fuzz_dep_.CoverTab[6678]++
//line /usr/local/go/src/net/lookup.go:110
		return n != '6'
//line /usr/local/go/src/net/lookup.go:110
		// _ = "end of CoverTab[6678]"
//line /usr/local/go/src/net/lookup.go:110
	}() {
//line /usr/local/go/src/net/lookup.go:110
		_go_fuzz_dep_.CoverTab[6679]++
							n = 0
//line /usr/local/go/src/net/lookup.go:111
		// _ = "end of CoverTab[6679]"
	} else {
//line /usr/local/go/src/net/lookup.go:112
		_go_fuzz_dep_.CoverTab[6680]++
//line /usr/local/go/src/net/lookup.go:112
		// _ = "end of CoverTab[6680]"
//line /usr/local/go/src/net/lookup.go:112
	}
//line /usr/local/go/src/net/lookup.go:112
	// _ = "end of CoverTab[6674]"
//line /usr/local/go/src/net/lookup.go:112
	_go_fuzz_dep_.CoverTab[6675]++
						return n
//line /usr/local/go/src/net/lookup.go:113
	// _ = "end of CoverTab[6675]"
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
	_go_fuzz_dep_.CoverTab[6681]++
//line /usr/local/go/src/net/lookup.go:161
	return r != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:161
		_go_fuzz_dep_.CoverTab[6682]++
//line /usr/local/go/src/net/lookup.go:161
		return r.PreferGo
//line /usr/local/go/src/net/lookup.go:161
		// _ = "end of CoverTab[6682]"
//line /usr/local/go/src/net/lookup.go:161
	}()
//line /usr/local/go/src/net/lookup.go:161
	// _ = "end of CoverTab[6681]"
//line /usr/local/go/src/net/lookup.go:161
}
func (r *Resolver) strictErrors() bool {
//line /usr/local/go/src/net/lookup.go:162
	_go_fuzz_dep_.CoverTab[6683]++
//line /usr/local/go/src/net/lookup.go:162
	return r != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:162
		_go_fuzz_dep_.CoverTab[6684]++
//line /usr/local/go/src/net/lookup.go:162
		return r.StrictErrors
//line /usr/local/go/src/net/lookup.go:162
		// _ = "end of CoverTab[6684]"
//line /usr/local/go/src/net/lookup.go:162
	}()
//line /usr/local/go/src/net/lookup.go:162
	// _ = "end of CoverTab[6683]"
//line /usr/local/go/src/net/lookup.go:162
}

func (r *Resolver) getLookupGroup() *singleflight.Group {
//line /usr/local/go/src/net/lookup.go:164
	_go_fuzz_dep_.CoverTab[6685]++
						if r == nil {
//line /usr/local/go/src/net/lookup.go:165
		_go_fuzz_dep_.CoverTab[6687]++
							return &DefaultResolver.lookupGroup
//line /usr/local/go/src/net/lookup.go:166
		// _ = "end of CoverTab[6687]"
	} else {
//line /usr/local/go/src/net/lookup.go:167
		_go_fuzz_dep_.CoverTab[6688]++
//line /usr/local/go/src/net/lookup.go:167
		// _ = "end of CoverTab[6688]"
//line /usr/local/go/src/net/lookup.go:167
	}
//line /usr/local/go/src/net/lookup.go:167
	// _ = "end of CoverTab[6685]"
//line /usr/local/go/src/net/lookup.go:167
	_go_fuzz_dep_.CoverTab[6686]++
						return &r.lookupGroup
//line /usr/local/go/src/net/lookup.go:168
	// _ = "end of CoverTab[6686]"
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
	_go_fuzz_dep_.CoverTab[6689]++
						return DefaultResolver.LookupHost(context.Background(), host)
//line /usr/local/go/src/net/lookup.go:177
	// _ = "end of CoverTab[6689]"
}

// LookupHost looks up the given host using the local resolver.
//line /usr/local/go/src/net/lookup.go:180
// It returns a slice of that host's addresses.
//line /usr/local/go/src/net/lookup.go:182
func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error) {
//line /usr/local/go/src/net/lookup.go:182
	_go_fuzz_dep_.CoverTab[6690]++

//line /usr/local/go/src/net/lookup.go:185
	if host == "" {
//line /usr/local/go/src/net/lookup.go:185
		_go_fuzz_dep_.CoverTab[6693]++
							return nil, &DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
//line /usr/local/go/src/net/lookup.go:186
		// _ = "end of CoverTab[6693]"
	} else {
//line /usr/local/go/src/net/lookup.go:187
		_go_fuzz_dep_.CoverTab[6694]++
//line /usr/local/go/src/net/lookup.go:187
		// _ = "end of CoverTab[6694]"
//line /usr/local/go/src/net/lookup.go:187
	}
//line /usr/local/go/src/net/lookup.go:187
	// _ = "end of CoverTab[6690]"
//line /usr/local/go/src/net/lookup.go:187
	_go_fuzz_dep_.CoverTab[6691]++
						if ip, _ := parseIPZone(host); ip != nil {
//line /usr/local/go/src/net/lookup.go:188
		_go_fuzz_dep_.CoverTab[6695]++
							return []string{host}, nil
//line /usr/local/go/src/net/lookup.go:189
		// _ = "end of CoverTab[6695]"
	} else {
//line /usr/local/go/src/net/lookup.go:190
		_go_fuzz_dep_.CoverTab[6696]++
//line /usr/local/go/src/net/lookup.go:190
		// _ = "end of CoverTab[6696]"
//line /usr/local/go/src/net/lookup.go:190
	}
//line /usr/local/go/src/net/lookup.go:190
	// _ = "end of CoverTab[6691]"
//line /usr/local/go/src/net/lookup.go:190
	_go_fuzz_dep_.CoverTab[6692]++
						return r.lookupHost(ctx, host)
//line /usr/local/go/src/net/lookup.go:191
	// _ = "end of CoverTab[6692]"
}

// LookupIP looks up host using the local resolver.
//line /usr/local/go/src/net/lookup.go:194
// It returns a slice of that host's IPv4 and IPv6 addresses.
//line /usr/local/go/src/net/lookup.go:196
func LookupIP(host string) ([]IP, error) {
//line /usr/local/go/src/net/lookup.go:196
	_go_fuzz_dep_.CoverTab[6697]++
						addrs, err := DefaultResolver.LookupIPAddr(context.Background(), host)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:198
		_go_fuzz_dep_.CoverTab[6700]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:199
		// _ = "end of CoverTab[6700]"
	} else {
//line /usr/local/go/src/net/lookup.go:200
		_go_fuzz_dep_.CoverTab[6701]++
//line /usr/local/go/src/net/lookup.go:200
		// _ = "end of CoverTab[6701]"
//line /usr/local/go/src/net/lookup.go:200
	}
//line /usr/local/go/src/net/lookup.go:200
	// _ = "end of CoverTab[6697]"
//line /usr/local/go/src/net/lookup.go:200
	_go_fuzz_dep_.CoverTab[6698]++
						ips := make([]IP, len(addrs))
						for i, ia := range addrs {
//line /usr/local/go/src/net/lookup.go:202
		_go_fuzz_dep_.CoverTab[6702]++
							ips[i] = ia.IP
//line /usr/local/go/src/net/lookup.go:203
		// _ = "end of CoverTab[6702]"
	}
//line /usr/local/go/src/net/lookup.go:204
	// _ = "end of CoverTab[6698]"
//line /usr/local/go/src/net/lookup.go:204
	_go_fuzz_dep_.CoverTab[6699]++
						return ips, nil
//line /usr/local/go/src/net/lookup.go:205
	// _ = "end of CoverTab[6699]"
}

// LookupIPAddr looks up host using the local resolver.
//line /usr/local/go/src/net/lookup.go:208
// It returns a slice of that host's IPv4 and IPv6 addresses.
//line /usr/local/go/src/net/lookup.go:210
func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error) {
//line /usr/local/go/src/net/lookup.go:210
	_go_fuzz_dep_.CoverTab[6703]++
						return r.lookupIPAddr(ctx, "ip", host)
//line /usr/local/go/src/net/lookup.go:211
	// _ = "end of CoverTab[6703]"
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
	_go_fuzz_dep_.CoverTab[6704]++
						afnet, _, err := parseNetwork(ctx, network, false)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:220
		_go_fuzz_dep_.CoverTab[6710]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:221
		// _ = "end of CoverTab[6710]"
	} else {
//line /usr/local/go/src/net/lookup.go:222
		_go_fuzz_dep_.CoverTab[6711]++
//line /usr/local/go/src/net/lookup.go:222
		// _ = "end of CoverTab[6711]"
//line /usr/local/go/src/net/lookup.go:222
	}
//line /usr/local/go/src/net/lookup.go:222
	// _ = "end of CoverTab[6704]"
//line /usr/local/go/src/net/lookup.go:222
	_go_fuzz_dep_.CoverTab[6705]++
						switch afnet {
	case "ip", "ip4", "ip6":
//line /usr/local/go/src/net/lookup.go:224
		_go_fuzz_dep_.CoverTab[6712]++
//line /usr/local/go/src/net/lookup.go:224
		// _ = "end of CoverTab[6712]"
	default:
//line /usr/local/go/src/net/lookup.go:225
		_go_fuzz_dep_.CoverTab[6713]++
							return nil, UnknownNetworkError(network)
//line /usr/local/go/src/net/lookup.go:226
		// _ = "end of CoverTab[6713]"
	}
//line /usr/local/go/src/net/lookup.go:227
	// _ = "end of CoverTab[6705]"
//line /usr/local/go/src/net/lookup.go:227
	_go_fuzz_dep_.CoverTab[6706]++

						if host == "" {
//line /usr/local/go/src/net/lookup.go:229
		_go_fuzz_dep_.CoverTab[6714]++
							return nil, &DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
//line /usr/local/go/src/net/lookup.go:230
		// _ = "end of CoverTab[6714]"
	} else {
//line /usr/local/go/src/net/lookup.go:231
		_go_fuzz_dep_.CoverTab[6715]++
//line /usr/local/go/src/net/lookup.go:231
		// _ = "end of CoverTab[6715]"
//line /usr/local/go/src/net/lookup.go:231
	}
//line /usr/local/go/src/net/lookup.go:231
	// _ = "end of CoverTab[6706]"
//line /usr/local/go/src/net/lookup.go:231
	_go_fuzz_dep_.CoverTab[6707]++
						addrs, err := r.internetAddrList(ctx, afnet, host)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:233
		_go_fuzz_dep_.CoverTab[6716]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:234
		// _ = "end of CoverTab[6716]"
	} else {
//line /usr/local/go/src/net/lookup.go:235
		_go_fuzz_dep_.CoverTab[6717]++
//line /usr/local/go/src/net/lookup.go:235
		// _ = "end of CoverTab[6717]"
//line /usr/local/go/src/net/lookup.go:235
	}
//line /usr/local/go/src/net/lookup.go:235
	// _ = "end of CoverTab[6707]"
//line /usr/local/go/src/net/lookup.go:235
	_go_fuzz_dep_.CoverTab[6708]++

						ips := make([]IP, 0, len(addrs))
						for _, addr := range addrs {
//line /usr/local/go/src/net/lookup.go:238
		_go_fuzz_dep_.CoverTab[6718]++
							ips = append(ips, addr.(*IPAddr).IP)
//line /usr/local/go/src/net/lookup.go:239
		// _ = "end of CoverTab[6718]"
	}
//line /usr/local/go/src/net/lookup.go:240
	// _ = "end of CoverTab[6708]"
//line /usr/local/go/src/net/lookup.go:240
	_go_fuzz_dep_.CoverTab[6709]++
						return ips, nil
//line /usr/local/go/src/net/lookup.go:241
	// _ = "end of CoverTab[6709]"
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
	_go_fuzz_dep_.CoverTab[6719]++

//line /usr/local/go/src/net/lookup.go:253
	ips, err := r.LookupIP(ctx, network, host)
	if err != nil {
//line /usr/local/go/src/net/lookup.go:254
		_go_fuzz_dep_.CoverTab[6722]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:255
		// _ = "end of CoverTab[6722]"
	} else {
//line /usr/local/go/src/net/lookup.go:256
		_go_fuzz_dep_.CoverTab[6723]++
//line /usr/local/go/src/net/lookup.go:256
		// _ = "end of CoverTab[6723]"
//line /usr/local/go/src/net/lookup.go:256
	}
//line /usr/local/go/src/net/lookup.go:256
	// _ = "end of CoverTab[6719]"
//line /usr/local/go/src/net/lookup.go:256
	_go_fuzz_dep_.CoverTab[6720]++
						ret := make([]netip.Addr, 0, len(ips))
						for _, ip := range ips {
//line /usr/local/go/src/net/lookup.go:258
		_go_fuzz_dep_.CoverTab[6724]++
							if a, ok := netip.AddrFromSlice(ip); ok {
//line /usr/local/go/src/net/lookup.go:259
			_go_fuzz_dep_.CoverTab[6725]++
								ret = append(ret, a)
//line /usr/local/go/src/net/lookup.go:260
			// _ = "end of CoverTab[6725]"
		} else {
//line /usr/local/go/src/net/lookup.go:261
			_go_fuzz_dep_.CoverTab[6726]++
//line /usr/local/go/src/net/lookup.go:261
			// _ = "end of CoverTab[6726]"
//line /usr/local/go/src/net/lookup.go:261
		}
//line /usr/local/go/src/net/lookup.go:261
		// _ = "end of CoverTab[6724]"
	}
//line /usr/local/go/src/net/lookup.go:262
	// _ = "end of CoverTab[6720]"
//line /usr/local/go/src/net/lookup.go:262
	_go_fuzz_dep_.CoverTab[6721]++
						return ret, nil
//line /usr/local/go/src/net/lookup.go:263
	// _ = "end of CoverTab[6721]"
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
	_go_fuzz_dep_.CoverTab[6727]++
						select {
	case <-ovc.lookupValues.Done():
//line /usr/local/go/src/net/lookup.go:278
		_go_fuzz_dep_.CoverTab[6728]++
							return nil
//line /usr/local/go/src/net/lookup.go:279
		// _ = "end of CoverTab[6728]"
	default:
//line /usr/local/go/src/net/lookup.go:280
		_go_fuzz_dep_.CoverTab[6729]++
							return ovc.lookupValues.Value(key)
//line /usr/local/go/src/net/lookup.go:281
		// _ = "end of CoverTab[6729]"
	}
//line /usr/local/go/src/net/lookup.go:282
	// _ = "end of CoverTab[6727]"
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
	_go_fuzz_dep_.CoverTab[6730]++
						return &onlyValuesCtx{Context: context.Background(), lookupValues: lookupCtx}
//line /usr/local/go/src/net/lookup.go:290
	// _ = "end of CoverTab[6730]"
}

// lookupIPAddr looks up host using the local resolver and particular network.
//line /usr/local/go/src/net/lookup.go:293
// It returns a slice of that host's IPv4 and IPv6 addresses.
//line /usr/local/go/src/net/lookup.go:295
func (r *Resolver) lookupIPAddr(ctx context.Context, network, host string) ([]IPAddr, error) {
//line /usr/local/go/src/net/lookup.go:295
	_go_fuzz_dep_.CoverTab[6731]++

//line /usr/local/go/src/net/lookup.go:298
	if host == "" {
//line /usr/local/go/src/net/lookup.go:298
		_go_fuzz_dep_.CoverTab[6738]++
							return nil, &DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
//line /usr/local/go/src/net/lookup.go:299
		// _ = "end of CoverTab[6738]"
	} else {
//line /usr/local/go/src/net/lookup.go:300
		_go_fuzz_dep_.CoverTab[6739]++
//line /usr/local/go/src/net/lookup.go:300
		// _ = "end of CoverTab[6739]"
//line /usr/local/go/src/net/lookup.go:300
	}
//line /usr/local/go/src/net/lookup.go:300
	// _ = "end of CoverTab[6731]"
//line /usr/local/go/src/net/lookup.go:300
	_go_fuzz_dep_.CoverTab[6732]++
						if ip, zone := parseIPZone(host); ip != nil {
//line /usr/local/go/src/net/lookup.go:301
		_go_fuzz_dep_.CoverTab[6740]++
							return []IPAddr{{IP: ip, Zone: zone}}, nil
//line /usr/local/go/src/net/lookup.go:302
		// _ = "end of CoverTab[6740]"
	} else {
//line /usr/local/go/src/net/lookup.go:303
		_go_fuzz_dep_.CoverTab[6741]++
//line /usr/local/go/src/net/lookup.go:303
		// _ = "end of CoverTab[6741]"
//line /usr/local/go/src/net/lookup.go:303
	}
//line /usr/local/go/src/net/lookup.go:303
	// _ = "end of CoverTab[6732]"
//line /usr/local/go/src/net/lookup.go:303
	_go_fuzz_dep_.CoverTab[6733]++
						trace, _ := ctx.Value(nettrace.TraceKey{}).(*nettrace.Trace)
						if trace != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:305
		_go_fuzz_dep_.CoverTab[6742]++
//line /usr/local/go/src/net/lookup.go:305
		return trace.DNSStart != nil
//line /usr/local/go/src/net/lookup.go:305
		// _ = "end of CoverTab[6742]"
//line /usr/local/go/src/net/lookup.go:305
	}() {
//line /usr/local/go/src/net/lookup.go:305
		_go_fuzz_dep_.CoverTab[6743]++
							trace.DNSStart(host)
//line /usr/local/go/src/net/lookup.go:306
		// _ = "end of CoverTab[6743]"
	} else {
//line /usr/local/go/src/net/lookup.go:307
		_go_fuzz_dep_.CoverTab[6744]++
//line /usr/local/go/src/net/lookup.go:307
		// _ = "end of CoverTab[6744]"
//line /usr/local/go/src/net/lookup.go:307
	}
//line /usr/local/go/src/net/lookup.go:307
	// _ = "end of CoverTab[6733]"
//line /usr/local/go/src/net/lookup.go:307
	_go_fuzz_dep_.CoverTab[6734]++

//line /usr/local/go/src/net/lookup.go:311
	resolverFunc := r.lookupIP
	if alt, _ := ctx.Value(nettrace.LookupIPAltResolverKey{}).(func(context.Context, string, string) ([]IPAddr, error)); alt != nil {
//line /usr/local/go/src/net/lookup.go:312
		_go_fuzz_dep_.CoverTab[6745]++
							resolverFunc = alt
//line /usr/local/go/src/net/lookup.go:313
		// _ = "end of CoverTab[6745]"
	} else {
//line /usr/local/go/src/net/lookup.go:314
		_go_fuzz_dep_.CoverTab[6746]++
//line /usr/local/go/src/net/lookup.go:314
		// _ = "end of CoverTab[6746]"
//line /usr/local/go/src/net/lookup.go:314
	}
//line /usr/local/go/src/net/lookup.go:314
	// _ = "end of CoverTab[6734]"
//line /usr/local/go/src/net/lookup.go:314
	_go_fuzz_dep_.CoverTab[6735]++

//line /usr/local/go/src/net/lookup.go:321
	lookupGroupCtx, lookupGroupCancel := context.WithCancel(withUnexpiredValuesPreserved(ctx))

	lookupKey := network + "\000" + host
	dnsWaitGroup.Add(1)
	ch := r.getLookupGroup().DoChan(lookupKey, func() (any, error) {
//line /usr/local/go/src/net/lookup.go:325
		_go_fuzz_dep_.CoverTab[6747]++
							return testHookLookupIP(lookupGroupCtx, resolverFunc, network, host)
//line /usr/local/go/src/net/lookup.go:326
		// _ = "end of CoverTab[6747]"
	})
//line /usr/local/go/src/net/lookup.go:327
	// _ = "end of CoverTab[6735]"
//line /usr/local/go/src/net/lookup.go:327
	_go_fuzz_dep_.CoverTab[6736]++

						dnsWaitGroupDone := func(ch <-chan singleflight.Result, cancelFn context.CancelFunc) {
//line /usr/local/go/src/net/lookup.go:329
		_go_fuzz_dep_.CoverTab[6748]++
							<-ch
							dnsWaitGroup.Done()
							cancelFn()
//line /usr/local/go/src/net/lookup.go:332
		// _ = "end of CoverTab[6748]"
	}
//line /usr/local/go/src/net/lookup.go:333
	// _ = "end of CoverTab[6736]"
//line /usr/local/go/src/net/lookup.go:333
	_go_fuzz_dep_.CoverTab[6737]++
						select {
	case <-ctx.Done():
//line /usr/local/go/src/net/lookup.go:335
		_go_fuzz_dep_.CoverTab[6749]++

//line /usr/local/go/src/net/lookup.go:343
		if r.getLookupGroup().ForgetUnshared(lookupKey) {
//line /usr/local/go/src/net/lookup.go:343
			_go_fuzz_dep_.CoverTab[6755]++
								lookupGroupCancel()
//line /usr/local/go/src/net/lookup.go:344
			_curRoutineNum10_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/lookup.go:344
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum10_)
								go func() {
//line /usr/local/go/src/net/lookup.go:345
				_go_fuzz_dep_.CoverTab[6756]++
//line /usr/local/go/src/net/lookup.go:345
				defer func() {
//line /usr/local/go/src/net/lookup.go:345
					_go_fuzz_dep_.CoverTab[6757]++
//line /usr/local/go/src/net/lookup.go:345
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum10_)
//line /usr/local/go/src/net/lookup.go:345
					// _ = "end of CoverTab[6757]"
//line /usr/local/go/src/net/lookup.go:345
				}()
//line /usr/local/go/src/net/lookup.go:345
				dnsWaitGroupDone(ch, func() { _go_fuzz_dep_.CoverTab[6758]++; // _ = "end of CoverTab[6758]" })
//line /usr/local/go/src/net/lookup.go:345
				// _ = "end of CoverTab[6756]"
//line /usr/local/go/src/net/lookup.go:345
			}()
//line /usr/local/go/src/net/lookup.go:345
			// _ = "end of CoverTab[6755]"
		} else {
//line /usr/local/go/src/net/lookup.go:346
			_go_fuzz_dep_.CoverTab[6759]++
//line /usr/local/go/src/net/lookup.go:346
			_curRoutineNum11_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/lookup.go:346
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum11_)
								go func() {
//line /usr/local/go/src/net/lookup.go:347
				_go_fuzz_dep_.CoverTab[6760]++
//line /usr/local/go/src/net/lookup.go:347
				defer func() {
//line /usr/local/go/src/net/lookup.go:347
					_go_fuzz_dep_.CoverTab[6761]++
//line /usr/local/go/src/net/lookup.go:347
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum11_)
//line /usr/local/go/src/net/lookup.go:347
					// _ = "end of CoverTab[6761]"
//line /usr/local/go/src/net/lookup.go:347
				}()
//line /usr/local/go/src/net/lookup.go:347
				dnsWaitGroupDone(ch, lookupGroupCancel)
//line /usr/local/go/src/net/lookup.go:347
				// _ = "end of CoverTab[6760]"
//line /usr/local/go/src/net/lookup.go:347
			}()
//line /usr/local/go/src/net/lookup.go:347
			// _ = "end of CoverTab[6759]"
		}
//line /usr/local/go/src/net/lookup.go:348
		// _ = "end of CoverTab[6749]"
//line /usr/local/go/src/net/lookup.go:348
		_go_fuzz_dep_.CoverTab[6750]++
							ctxErr := ctx.Err()
							err := &DNSError{
			Err:		mapErr(ctxErr).Error(),
			Name:		host,
			IsTimeout:	ctxErr == context.DeadlineExceeded,
		}
		if trace != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:355
			_go_fuzz_dep_.CoverTab[6762]++
//line /usr/local/go/src/net/lookup.go:355
			return trace.DNSDone != nil
//line /usr/local/go/src/net/lookup.go:355
			// _ = "end of CoverTab[6762]"
//line /usr/local/go/src/net/lookup.go:355
		}() {
//line /usr/local/go/src/net/lookup.go:355
			_go_fuzz_dep_.CoverTab[6763]++
								trace.DNSDone(nil, false, err)
//line /usr/local/go/src/net/lookup.go:356
			// _ = "end of CoverTab[6763]"
		} else {
//line /usr/local/go/src/net/lookup.go:357
			_go_fuzz_dep_.CoverTab[6764]++
//line /usr/local/go/src/net/lookup.go:357
			// _ = "end of CoverTab[6764]"
//line /usr/local/go/src/net/lookup.go:357
		}
//line /usr/local/go/src/net/lookup.go:357
		// _ = "end of CoverTab[6750]"
//line /usr/local/go/src/net/lookup.go:357
		_go_fuzz_dep_.CoverTab[6751]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:358
		// _ = "end of CoverTab[6751]"
	case r := <-ch:
//line /usr/local/go/src/net/lookup.go:359
		_go_fuzz_dep_.CoverTab[6752]++
							dnsWaitGroup.Done()
							lookupGroupCancel()
							err := r.Err
							if err != nil {
//line /usr/local/go/src/net/lookup.go:363
			_go_fuzz_dep_.CoverTab[6765]++
								if _, ok := err.(*DNSError); !ok {
//line /usr/local/go/src/net/lookup.go:364
				_go_fuzz_dep_.CoverTab[6766]++
									isTimeout := false
									if err == context.DeadlineExceeded {
//line /usr/local/go/src/net/lookup.go:366
					_go_fuzz_dep_.CoverTab[6768]++
										isTimeout = true
//line /usr/local/go/src/net/lookup.go:367
					// _ = "end of CoverTab[6768]"
				} else {
//line /usr/local/go/src/net/lookup.go:368
					_go_fuzz_dep_.CoverTab[6769]++
//line /usr/local/go/src/net/lookup.go:368
					if terr, ok := err.(timeout); ok {
//line /usr/local/go/src/net/lookup.go:368
						_go_fuzz_dep_.CoverTab[6770]++
											isTimeout = terr.Timeout()
//line /usr/local/go/src/net/lookup.go:369
						// _ = "end of CoverTab[6770]"
					} else {
//line /usr/local/go/src/net/lookup.go:370
						_go_fuzz_dep_.CoverTab[6771]++
//line /usr/local/go/src/net/lookup.go:370
						// _ = "end of CoverTab[6771]"
//line /usr/local/go/src/net/lookup.go:370
					}
//line /usr/local/go/src/net/lookup.go:370
					// _ = "end of CoverTab[6769]"
//line /usr/local/go/src/net/lookup.go:370
				}
//line /usr/local/go/src/net/lookup.go:370
				// _ = "end of CoverTab[6766]"
//line /usr/local/go/src/net/lookup.go:370
				_go_fuzz_dep_.CoverTab[6767]++
									err = &DNSError{
					Err:		err.Error(),
					Name:		host,
					IsTimeout:	isTimeout,
				}
//line /usr/local/go/src/net/lookup.go:375
				// _ = "end of CoverTab[6767]"
			} else {
//line /usr/local/go/src/net/lookup.go:376
				_go_fuzz_dep_.CoverTab[6772]++
//line /usr/local/go/src/net/lookup.go:376
				// _ = "end of CoverTab[6772]"
//line /usr/local/go/src/net/lookup.go:376
			}
//line /usr/local/go/src/net/lookup.go:376
			// _ = "end of CoverTab[6765]"
		} else {
//line /usr/local/go/src/net/lookup.go:377
			_go_fuzz_dep_.CoverTab[6773]++
//line /usr/local/go/src/net/lookup.go:377
			// _ = "end of CoverTab[6773]"
//line /usr/local/go/src/net/lookup.go:377
		}
//line /usr/local/go/src/net/lookup.go:377
		// _ = "end of CoverTab[6752]"
//line /usr/local/go/src/net/lookup.go:377
		_go_fuzz_dep_.CoverTab[6753]++
							if trace != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:378
			_go_fuzz_dep_.CoverTab[6774]++
//line /usr/local/go/src/net/lookup.go:378
			return trace.DNSDone != nil
//line /usr/local/go/src/net/lookup.go:378
			// _ = "end of CoverTab[6774]"
//line /usr/local/go/src/net/lookup.go:378
		}() {
//line /usr/local/go/src/net/lookup.go:378
			_go_fuzz_dep_.CoverTab[6775]++
								addrs, _ := r.Val.([]IPAddr)
								trace.DNSDone(ipAddrsEface(addrs), r.Shared, err)
//line /usr/local/go/src/net/lookup.go:380
			// _ = "end of CoverTab[6775]"
		} else {
//line /usr/local/go/src/net/lookup.go:381
			_go_fuzz_dep_.CoverTab[6776]++
//line /usr/local/go/src/net/lookup.go:381
			// _ = "end of CoverTab[6776]"
//line /usr/local/go/src/net/lookup.go:381
		}
//line /usr/local/go/src/net/lookup.go:381
		// _ = "end of CoverTab[6753]"
//line /usr/local/go/src/net/lookup.go:381
		_go_fuzz_dep_.CoverTab[6754]++
							return lookupIPReturn(r.Val, err, r.Shared)
//line /usr/local/go/src/net/lookup.go:382
		// _ = "end of CoverTab[6754]"
	}
//line /usr/local/go/src/net/lookup.go:383
	// _ = "end of CoverTab[6737]"
}

// lookupIPReturn turns the return values from singleflight.Do into
//line /usr/local/go/src/net/lookup.go:386
// the return values from LookupIP.
//line /usr/local/go/src/net/lookup.go:388
func lookupIPReturn(addrsi any, err error, shared bool) ([]IPAddr, error) {
//line /usr/local/go/src/net/lookup.go:388
	_go_fuzz_dep_.CoverTab[6777]++
						if err != nil {
//line /usr/local/go/src/net/lookup.go:389
		_go_fuzz_dep_.CoverTab[6780]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:390
		// _ = "end of CoverTab[6780]"
	} else {
//line /usr/local/go/src/net/lookup.go:391
		_go_fuzz_dep_.CoverTab[6781]++
//line /usr/local/go/src/net/lookup.go:391
		// _ = "end of CoverTab[6781]"
//line /usr/local/go/src/net/lookup.go:391
	}
//line /usr/local/go/src/net/lookup.go:391
	// _ = "end of CoverTab[6777]"
//line /usr/local/go/src/net/lookup.go:391
	_go_fuzz_dep_.CoverTab[6778]++
						addrs := addrsi.([]IPAddr)
						if shared {
//line /usr/local/go/src/net/lookup.go:393
		_go_fuzz_dep_.CoverTab[6782]++
							clone := make([]IPAddr, len(addrs))
							copy(clone, addrs)
							addrs = clone
//line /usr/local/go/src/net/lookup.go:396
		// _ = "end of CoverTab[6782]"
	} else {
//line /usr/local/go/src/net/lookup.go:397
		_go_fuzz_dep_.CoverTab[6783]++
//line /usr/local/go/src/net/lookup.go:397
		// _ = "end of CoverTab[6783]"
//line /usr/local/go/src/net/lookup.go:397
	}
//line /usr/local/go/src/net/lookup.go:397
	// _ = "end of CoverTab[6778]"
//line /usr/local/go/src/net/lookup.go:397
	_go_fuzz_dep_.CoverTab[6779]++
						return addrs, nil
//line /usr/local/go/src/net/lookup.go:398
	// _ = "end of CoverTab[6779]"
}

// ipAddrsEface returns an empty interface slice of addrs.
func ipAddrsEface(addrs []IPAddr) []any {
//line /usr/local/go/src/net/lookup.go:402
	_go_fuzz_dep_.CoverTab[6784]++
						s := make([]any, len(addrs))
						for i, v := range addrs {
//line /usr/local/go/src/net/lookup.go:404
		_go_fuzz_dep_.CoverTab[6786]++
							s[i] = v
//line /usr/local/go/src/net/lookup.go:405
		// _ = "end of CoverTab[6786]"
	}
//line /usr/local/go/src/net/lookup.go:406
	// _ = "end of CoverTab[6784]"
//line /usr/local/go/src/net/lookup.go:406
	_go_fuzz_dep_.CoverTab[6785]++
						return s
//line /usr/local/go/src/net/lookup.go:407
	// _ = "end of CoverTab[6785]"
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
	_go_fuzz_dep_.CoverTab[6787]++
						return DefaultResolver.LookupPort(context.Background(), network, service)
//line /usr/local/go/src/net/lookup.go:415
	// _ = "end of CoverTab[6787]"
}

// LookupPort looks up the port for the given network and service.
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error) {
//line /usr/local/go/src/net/lookup.go:419
	_go_fuzz_dep_.CoverTab[6788]++
						port, needsLookup := parsePort(service)
						if needsLookup {
//line /usr/local/go/src/net/lookup.go:421
		_go_fuzz_dep_.CoverTab[6791]++
							switch network {
		case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
//line /usr/local/go/src/net/lookup.go:423
			_go_fuzz_dep_.CoverTab[6793]++
//line /usr/local/go/src/net/lookup.go:423
			// _ = "end of CoverTab[6793]"
		case "":
//line /usr/local/go/src/net/lookup.go:424
			_go_fuzz_dep_.CoverTab[6794]++
								network = "ip"
//line /usr/local/go/src/net/lookup.go:425
			// _ = "end of CoverTab[6794]"
		default:
//line /usr/local/go/src/net/lookup.go:426
			_go_fuzz_dep_.CoverTab[6795]++
								return 0, &AddrError{Err: "unknown network", Addr: network}
//line /usr/local/go/src/net/lookup.go:427
			// _ = "end of CoverTab[6795]"
		}
//line /usr/local/go/src/net/lookup.go:428
		// _ = "end of CoverTab[6791]"
//line /usr/local/go/src/net/lookup.go:428
		_go_fuzz_dep_.CoverTab[6792]++
							port, err = r.lookupPort(ctx, network, service)
							if err != nil {
//line /usr/local/go/src/net/lookup.go:430
			_go_fuzz_dep_.CoverTab[6796]++
								return 0, err
//line /usr/local/go/src/net/lookup.go:431
			// _ = "end of CoverTab[6796]"
		} else {
//line /usr/local/go/src/net/lookup.go:432
			_go_fuzz_dep_.CoverTab[6797]++
//line /usr/local/go/src/net/lookup.go:432
			// _ = "end of CoverTab[6797]"
//line /usr/local/go/src/net/lookup.go:432
		}
//line /usr/local/go/src/net/lookup.go:432
		// _ = "end of CoverTab[6792]"
	} else {
//line /usr/local/go/src/net/lookup.go:433
		_go_fuzz_dep_.CoverTab[6798]++
//line /usr/local/go/src/net/lookup.go:433
		// _ = "end of CoverTab[6798]"
//line /usr/local/go/src/net/lookup.go:433
	}
//line /usr/local/go/src/net/lookup.go:433
	// _ = "end of CoverTab[6788]"
//line /usr/local/go/src/net/lookup.go:433
	_go_fuzz_dep_.CoverTab[6789]++
						if 0 > port || func() bool {
//line /usr/local/go/src/net/lookup.go:434
		_go_fuzz_dep_.CoverTab[6799]++
//line /usr/local/go/src/net/lookup.go:434
		return port > 65535
//line /usr/local/go/src/net/lookup.go:434
		// _ = "end of CoverTab[6799]"
//line /usr/local/go/src/net/lookup.go:434
	}() {
//line /usr/local/go/src/net/lookup.go:434
		_go_fuzz_dep_.CoverTab[6800]++
							return 0, &AddrError{Err: "invalid port", Addr: service}
//line /usr/local/go/src/net/lookup.go:435
		// _ = "end of CoverTab[6800]"
	} else {
//line /usr/local/go/src/net/lookup.go:436
		_go_fuzz_dep_.CoverTab[6801]++
//line /usr/local/go/src/net/lookup.go:436
		// _ = "end of CoverTab[6801]"
//line /usr/local/go/src/net/lookup.go:436
	}
//line /usr/local/go/src/net/lookup.go:436
	// _ = "end of CoverTab[6789]"
//line /usr/local/go/src/net/lookup.go:436
	_go_fuzz_dep_.CoverTab[6790]++
						return port, nil
//line /usr/local/go/src/net/lookup.go:437
	// _ = "end of CoverTab[6790]"
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
	_go_fuzz_dep_.CoverTab[6802]++
						return DefaultResolver.LookupCNAME(context.Background(), host)
//line /usr/local/go/src/net/lookup.go:457
	// _ = "end of CoverTab[6802]"
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
	_go_fuzz_dep_.CoverTab[6803]++
						cname, err := r.lookupCNAME(ctx, host)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:475
		_go_fuzz_dep_.CoverTab[6806]++
							return "", err
//line /usr/local/go/src/net/lookup.go:476
		// _ = "end of CoverTab[6806]"
	} else {
//line /usr/local/go/src/net/lookup.go:477
		_go_fuzz_dep_.CoverTab[6807]++
//line /usr/local/go/src/net/lookup.go:477
		// _ = "end of CoverTab[6807]"
//line /usr/local/go/src/net/lookup.go:477
	}
//line /usr/local/go/src/net/lookup.go:477
	// _ = "end of CoverTab[6803]"
//line /usr/local/go/src/net/lookup.go:477
	_go_fuzz_dep_.CoverTab[6804]++
						if !isDomainName(cname) {
//line /usr/local/go/src/net/lookup.go:478
		_go_fuzz_dep_.CoverTab[6808]++
							return "", &DNSError{Err: errMalformedDNSRecordsDetail, Name: host}
//line /usr/local/go/src/net/lookup.go:479
		// _ = "end of CoverTab[6808]"
	} else {
//line /usr/local/go/src/net/lookup.go:480
		_go_fuzz_dep_.CoverTab[6809]++
//line /usr/local/go/src/net/lookup.go:480
		// _ = "end of CoverTab[6809]"
//line /usr/local/go/src/net/lookup.go:480
	}
//line /usr/local/go/src/net/lookup.go:480
	// _ = "end of CoverTab[6804]"
//line /usr/local/go/src/net/lookup.go:480
	_go_fuzz_dep_.CoverTab[6805]++
						return cname, nil
//line /usr/local/go/src/net/lookup.go:481
	// _ = "end of CoverTab[6805]"
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
	_go_fuzz_dep_.CoverTab[6810]++
						return DefaultResolver.LookupSRV(context.Background(), service, proto, name)
//line /usr/local/go/src/net/lookup.go:499
	// _ = "end of CoverTab[6810]"
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
	_go_fuzz_dep_.CoverTab[6811]++
						cname, addrs, err := r.lookupSRV(ctx, service, proto, name)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:518
		_go_fuzz_dep_.CoverTab[6816]++
							return "", nil, err
//line /usr/local/go/src/net/lookup.go:519
		// _ = "end of CoverTab[6816]"
	} else {
//line /usr/local/go/src/net/lookup.go:520
		_go_fuzz_dep_.CoverTab[6817]++
//line /usr/local/go/src/net/lookup.go:520
		// _ = "end of CoverTab[6817]"
//line /usr/local/go/src/net/lookup.go:520
	}
//line /usr/local/go/src/net/lookup.go:520
	// _ = "end of CoverTab[6811]"
//line /usr/local/go/src/net/lookup.go:520
	_go_fuzz_dep_.CoverTab[6812]++
						if cname != "" && func() bool {
//line /usr/local/go/src/net/lookup.go:521
		_go_fuzz_dep_.CoverTab[6818]++
//line /usr/local/go/src/net/lookup.go:521
		return !isDomainName(cname)
//line /usr/local/go/src/net/lookup.go:521
		// _ = "end of CoverTab[6818]"
//line /usr/local/go/src/net/lookup.go:521
	}() {
//line /usr/local/go/src/net/lookup.go:521
		_go_fuzz_dep_.CoverTab[6819]++
							return "", nil, &DNSError{Err: "SRV header name is invalid", Name: name}
//line /usr/local/go/src/net/lookup.go:522
		// _ = "end of CoverTab[6819]"
	} else {
//line /usr/local/go/src/net/lookup.go:523
		_go_fuzz_dep_.CoverTab[6820]++
//line /usr/local/go/src/net/lookup.go:523
		// _ = "end of CoverTab[6820]"
//line /usr/local/go/src/net/lookup.go:523
	}
//line /usr/local/go/src/net/lookup.go:523
	// _ = "end of CoverTab[6812]"
//line /usr/local/go/src/net/lookup.go:523
	_go_fuzz_dep_.CoverTab[6813]++
						filteredAddrs := make([]*SRV, 0, len(addrs))
						for _, addr := range addrs {
//line /usr/local/go/src/net/lookup.go:525
		_go_fuzz_dep_.CoverTab[6821]++
							if addr == nil {
//line /usr/local/go/src/net/lookup.go:526
			_go_fuzz_dep_.CoverTab[6824]++
								continue
//line /usr/local/go/src/net/lookup.go:527
			// _ = "end of CoverTab[6824]"
		} else {
//line /usr/local/go/src/net/lookup.go:528
			_go_fuzz_dep_.CoverTab[6825]++
//line /usr/local/go/src/net/lookup.go:528
			// _ = "end of CoverTab[6825]"
//line /usr/local/go/src/net/lookup.go:528
		}
//line /usr/local/go/src/net/lookup.go:528
		// _ = "end of CoverTab[6821]"
//line /usr/local/go/src/net/lookup.go:528
		_go_fuzz_dep_.CoverTab[6822]++
							if !isDomainName(addr.Target) {
//line /usr/local/go/src/net/lookup.go:529
			_go_fuzz_dep_.CoverTab[6826]++
								continue
//line /usr/local/go/src/net/lookup.go:530
			// _ = "end of CoverTab[6826]"
		} else {
//line /usr/local/go/src/net/lookup.go:531
			_go_fuzz_dep_.CoverTab[6827]++
//line /usr/local/go/src/net/lookup.go:531
			// _ = "end of CoverTab[6827]"
//line /usr/local/go/src/net/lookup.go:531
		}
//line /usr/local/go/src/net/lookup.go:531
		// _ = "end of CoverTab[6822]"
//line /usr/local/go/src/net/lookup.go:531
		_go_fuzz_dep_.CoverTab[6823]++
							filteredAddrs = append(filteredAddrs, addr)
//line /usr/local/go/src/net/lookup.go:532
		// _ = "end of CoverTab[6823]"
	}
//line /usr/local/go/src/net/lookup.go:533
	// _ = "end of CoverTab[6813]"
//line /usr/local/go/src/net/lookup.go:533
	_go_fuzz_dep_.CoverTab[6814]++
						if len(addrs) != len(filteredAddrs) {
//line /usr/local/go/src/net/lookup.go:534
		_go_fuzz_dep_.CoverTab[6828]++
							return cname, filteredAddrs, &DNSError{Err: errMalformedDNSRecordsDetail, Name: name}
//line /usr/local/go/src/net/lookup.go:535
		// _ = "end of CoverTab[6828]"
	} else {
//line /usr/local/go/src/net/lookup.go:536
		_go_fuzz_dep_.CoverTab[6829]++
//line /usr/local/go/src/net/lookup.go:536
		// _ = "end of CoverTab[6829]"
//line /usr/local/go/src/net/lookup.go:536
	}
//line /usr/local/go/src/net/lookup.go:536
	// _ = "end of CoverTab[6814]"
//line /usr/local/go/src/net/lookup.go:536
	_go_fuzz_dep_.CoverTab[6815]++
						return cname, filteredAddrs, nil
//line /usr/local/go/src/net/lookup.go:537
	// _ = "end of CoverTab[6815]"
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
	_go_fuzz_dep_.CoverTab[6830]++
						return DefaultResolver.LookupMX(context.Background(), name)
//line /usr/local/go/src/net/lookup.go:550
	// _ = "end of CoverTab[6830]"
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
	_go_fuzz_dep_.CoverTab[6831]++
						records, err := r.lookupMX(ctx, name)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:561
		_go_fuzz_dep_.CoverTab[6835]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:562
		// _ = "end of CoverTab[6835]"
	} else {
//line /usr/local/go/src/net/lookup.go:563
		_go_fuzz_dep_.CoverTab[6836]++
//line /usr/local/go/src/net/lookup.go:563
		// _ = "end of CoverTab[6836]"
//line /usr/local/go/src/net/lookup.go:563
	}
//line /usr/local/go/src/net/lookup.go:563
	// _ = "end of CoverTab[6831]"
//line /usr/local/go/src/net/lookup.go:563
	_go_fuzz_dep_.CoverTab[6832]++
						filteredMX := make([]*MX, 0, len(records))
						for _, mx := range records {
//line /usr/local/go/src/net/lookup.go:565
		_go_fuzz_dep_.CoverTab[6837]++
							if mx == nil {
//line /usr/local/go/src/net/lookup.go:566
			_go_fuzz_dep_.CoverTab[6840]++
								continue
//line /usr/local/go/src/net/lookup.go:567
			// _ = "end of CoverTab[6840]"
		} else {
//line /usr/local/go/src/net/lookup.go:568
			_go_fuzz_dep_.CoverTab[6841]++
//line /usr/local/go/src/net/lookup.go:568
			// _ = "end of CoverTab[6841]"
//line /usr/local/go/src/net/lookup.go:568
		}
//line /usr/local/go/src/net/lookup.go:568
		// _ = "end of CoverTab[6837]"
//line /usr/local/go/src/net/lookup.go:568
		_go_fuzz_dep_.CoverTab[6838]++
							if !isDomainName(mx.Host) {
//line /usr/local/go/src/net/lookup.go:569
			_go_fuzz_dep_.CoverTab[6842]++
								continue
//line /usr/local/go/src/net/lookup.go:570
			// _ = "end of CoverTab[6842]"
		} else {
//line /usr/local/go/src/net/lookup.go:571
			_go_fuzz_dep_.CoverTab[6843]++
//line /usr/local/go/src/net/lookup.go:571
			// _ = "end of CoverTab[6843]"
//line /usr/local/go/src/net/lookup.go:571
		}
//line /usr/local/go/src/net/lookup.go:571
		// _ = "end of CoverTab[6838]"
//line /usr/local/go/src/net/lookup.go:571
		_go_fuzz_dep_.CoverTab[6839]++
							filteredMX = append(filteredMX, mx)
//line /usr/local/go/src/net/lookup.go:572
		// _ = "end of CoverTab[6839]"
	}
//line /usr/local/go/src/net/lookup.go:573
	// _ = "end of CoverTab[6832]"
//line /usr/local/go/src/net/lookup.go:573
	_go_fuzz_dep_.CoverTab[6833]++
						if len(records) != len(filteredMX) {
//line /usr/local/go/src/net/lookup.go:574
		_go_fuzz_dep_.CoverTab[6844]++
							return filteredMX, &DNSError{Err: errMalformedDNSRecordsDetail, Name: name}
//line /usr/local/go/src/net/lookup.go:575
		// _ = "end of CoverTab[6844]"
	} else {
//line /usr/local/go/src/net/lookup.go:576
		_go_fuzz_dep_.CoverTab[6845]++
//line /usr/local/go/src/net/lookup.go:576
		// _ = "end of CoverTab[6845]"
//line /usr/local/go/src/net/lookup.go:576
	}
//line /usr/local/go/src/net/lookup.go:576
	// _ = "end of CoverTab[6833]"
//line /usr/local/go/src/net/lookup.go:576
	_go_fuzz_dep_.CoverTab[6834]++
						return filteredMX, nil
//line /usr/local/go/src/net/lookup.go:577
	// _ = "end of CoverTab[6834]"
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
	_go_fuzz_dep_.CoverTab[6846]++
						return DefaultResolver.LookupNS(context.Background(), name)
//line /usr/local/go/src/net/lookup.go:590
	// _ = "end of CoverTab[6846]"
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
	_go_fuzz_dep_.CoverTab[6847]++
						records, err := r.lookupNS(ctx, name)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:601
		_go_fuzz_dep_.CoverTab[6851]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:602
		// _ = "end of CoverTab[6851]"
	} else {
//line /usr/local/go/src/net/lookup.go:603
		_go_fuzz_dep_.CoverTab[6852]++
//line /usr/local/go/src/net/lookup.go:603
		// _ = "end of CoverTab[6852]"
//line /usr/local/go/src/net/lookup.go:603
	}
//line /usr/local/go/src/net/lookup.go:603
	// _ = "end of CoverTab[6847]"
//line /usr/local/go/src/net/lookup.go:603
	_go_fuzz_dep_.CoverTab[6848]++
						filteredNS := make([]*NS, 0, len(records))
						for _, ns := range records {
//line /usr/local/go/src/net/lookup.go:605
		_go_fuzz_dep_.CoverTab[6853]++
							if ns == nil {
//line /usr/local/go/src/net/lookup.go:606
			_go_fuzz_dep_.CoverTab[6856]++
								continue
//line /usr/local/go/src/net/lookup.go:607
			// _ = "end of CoverTab[6856]"
		} else {
//line /usr/local/go/src/net/lookup.go:608
			_go_fuzz_dep_.CoverTab[6857]++
//line /usr/local/go/src/net/lookup.go:608
			// _ = "end of CoverTab[6857]"
//line /usr/local/go/src/net/lookup.go:608
		}
//line /usr/local/go/src/net/lookup.go:608
		// _ = "end of CoverTab[6853]"
//line /usr/local/go/src/net/lookup.go:608
		_go_fuzz_dep_.CoverTab[6854]++
							if !isDomainName(ns.Host) {
//line /usr/local/go/src/net/lookup.go:609
			_go_fuzz_dep_.CoverTab[6858]++
								continue
//line /usr/local/go/src/net/lookup.go:610
			// _ = "end of CoverTab[6858]"
		} else {
//line /usr/local/go/src/net/lookup.go:611
			_go_fuzz_dep_.CoverTab[6859]++
//line /usr/local/go/src/net/lookup.go:611
			// _ = "end of CoverTab[6859]"
//line /usr/local/go/src/net/lookup.go:611
		}
//line /usr/local/go/src/net/lookup.go:611
		// _ = "end of CoverTab[6854]"
//line /usr/local/go/src/net/lookup.go:611
		_go_fuzz_dep_.CoverTab[6855]++
							filteredNS = append(filteredNS, ns)
//line /usr/local/go/src/net/lookup.go:612
		// _ = "end of CoverTab[6855]"
	}
//line /usr/local/go/src/net/lookup.go:613
	// _ = "end of CoverTab[6848]"
//line /usr/local/go/src/net/lookup.go:613
	_go_fuzz_dep_.CoverTab[6849]++
						if len(records) != len(filteredNS) {
//line /usr/local/go/src/net/lookup.go:614
		_go_fuzz_dep_.CoverTab[6860]++
							return filteredNS, &DNSError{Err: errMalformedDNSRecordsDetail, Name: name}
//line /usr/local/go/src/net/lookup.go:615
		// _ = "end of CoverTab[6860]"
	} else {
//line /usr/local/go/src/net/lookup.go:616
		_go_fuzz_dep_.CoverTab[6861]++
//line /usr/local/go/src/net/lookup.go:616
		// _ = "end of CoverTab[6861]"
//line /usr/local/go/src/net/lookup.go:616
	}
//line /usr/local/go/src/net/lookup.go:616
	// _ = "end of CoverTab[6849]"
//line /usr/local/go/src/net/lookup.go:616
	_go_fuzz_dep_.CoverTab[6850]++
						return filteredNS, nil
//line /usr/local/go/src/net/lookup.go:617
	// _ = "end of CoverTab[6850]"
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
	_go_fuzz_dep_.CoverTab[6862]++
						return DefaultResolver.lookupTXT(context.Background(), name)
//line /usr/local/go/src/net/lookup.go:625
	// _ = "end of CoverTab[6862]"
}

// LookupTXT returns the DNS TXT records for the given domain name.
func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error) {
//line /usr/local/go/src/net/lookup.go:629
	_go_fuzz_dep_.CoverTab[6863]++
						return r.lookupTXT(ctx, name)
//line /usr/local/go/src/net/lookup.go:630
	// _ = "end of CoverTab[6863]"
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
	_go_fuzz_dep_.CoverTab[6864]++
						return DefaultResolver.LookupAddr(context.Background(), addr)
//line /usr/local/go/src/net/lookup.go:646
	// _ = "end of CoverTab[6864]"
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
	_go_fuzz_dep_.CoverTab[6865]++
						names, err := r.lookupAddr(ctx, addr)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:657
		_go_fuzz_dep_.CoverTab[6869]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:658
		// _ = "end of CoverTab[6869]"
	} else {
//line /usr/local/go/src/net/lookup.go:659
		_go_fuzz_dep_.CoverTab[6870]++
//line /usr/local/go/src/net/lookup.go:659
		// _ = "end of CoverTab[6870]"
//line /usr/local/go/src/net/lookup.go:659
	}
//line /usr/local/go/src/net/lookup.go:659
	// _ = "end of CoverTab[6865]"
//line /usr/local/go/src/net/lookup.go:659
	_go_fuzz_dep_.CoverTab[6866]++
						filteredNames := make([]string, 0, len(names))
						for _, name := range names {
//line /usr/local/go/src/net/lookup.go:661
		_go_fuzz_dep_.CoverTab[6871]++
							if isDomainName(name) {
//line /usr/local/go/src/net/lookup.go:662
			_go_fuzz_dep_.CoverTab[6872]++
								filteredNames = append(filteredNames, name)
//line /usr/local/go/src/net/lookup.go:663
			// _ = "end of CoverTab[6872]"
		} else {
//line /usr/local/go/src/net/lookup.go:664
			_go_fuzz_dep_.CoverTab[6873]++
//line /usr/local/go/src/net/lookup.go:664
			// _ = "end of CoverTab[6873]"
//line /usr/local/go/src/net/lookup.go:664
		}
//line /usr/local/go/src/net/lookup.go:664
		// _ = "end of CoverTab[6871]"
	}
//line /usr/local/go/src/net/lookup.go:665
	// _ = "end of CoverTab[6866]"
//line /usr/local/go/src/net/lookup.go:665
	_go_fuzz_dep_.CoverTab[6867]++
						if len(names) != len(filteredNames) {
//line /usr/local/go/src/net/lookup.go:666
		_go_fuzz_dep_.CoverTab[6874]++
							return filteredNames, &DNSError{Err: errMalformedDNSRecordsDetail, Name: addr}
//line /usr/local/go/src/net/lookup.go:667
		// _ = "end of CoverTab[6874]"
	} else {
//line /usr/local/go/src/net/lookup.go:668
		_go_fuzz_dep_.CoverTab[6875]++
//line /usr/local/go/src/net/lookup.go:668
		// _ = "end of CoverTab[6875]"
//line /usr/local/go/src/net/lookup.go:668
	}
//line /usr/local/go/src/net/lookup.go:668
	// _ = "end of CoverTab[6867]"
//line /usr/local/go/src/net/lookup.go:668
	_go_fuzz_dep_.CoverTab[6868]++
						return filteredNames, nil
//line /usr/local/go/src/net/lookup.go:669
	// _ = "end of CoverTab[6868]"
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
	_go_fuzz_dep_.CoverTab[6876]++
	// Calling Dial here is scary -- we have to be sure not to
	// dial a name that will require a DNS lookup, or Dial will
	// call back here to translate it. The DNS config parser has
	// already checked that all the cfg.servers are IP
	// addresses, which Dial will use without a DNS lookup.
	var c Conn
	var err error
	if r != nil && func() bool {
//line /usr/local/go/src/net/lookup.go:688
		_go_fuzz_dep_.CoverTab[6879]++
//line /usr/local/go/src/net/lookup.go:688
		return r.Dial != nil
//line /usr/local/go/src/net/lookup.go:688
		// _ = "end of CoverTab[6879]"
//line /usr/local/go/src/net/lookup.go:688
	}() {
//line /usr/local/go/src/net/lookup.go:688
		_go_fuzz_dep_.CoverTab[6880]++
							c, err = r.Dial(ctx, network, server)
//line /usr/local/go/src/net/lookup.go:689
		// _ = "end of CoverTab[6880]"
	} else {
//line /usr/local/go/src/net/lookup.go:690
		_go_fuzz_dep_.CoverTab[6881]++
							var d Dialer
							c, err = d.DialContext(ctx, network, server)
//line /usr/local/go/src/net/lookup.go:692
		// _ = "end of CoverTab[6881]"
	}
//line /usr/local/go/src/net/lookup.go:693
	// _ = "end of CoverTab[6876]"
//line /usr/local/go/src/net/lookup.go:693
	_go_fuzz_dep_.CoverTab[6877]++
						if err != nil {
//line /usr/local/go/src/net/lookup.go:694
		_go_fuzz_dep_.CoverTab[6882]++
							return nil, mapErr(err)
//line /usr/local/go/src/net/lookup.go:695
		// _ = "end of CoverTab[6882]"
	} else {
//line /usr/local/go/src/net/lookup.go:696
		_go_fuzz_dep_.CoverTab[6883]++
//line /usr/local/go/src/net/lookup.go:696
		// _ = "end of CoverTab[6883]"
//line /usr/local/go/src/net/lookup.go:696
	}
//line /usr/local/go/src/net/lookup.go:696
	// _ = "end of CoverTab[6877]"
//line /usr/local/go/src/net/lookup.go:696
	_go_fuzz_dep_.CoverTab[6878]++
						return c, nil
//line /usr/local/go/src/net/lookup.go:697
	// _ = "end of CoverTab[6878]"
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
	_go_fuzz_dep_.CoverTab[6884]++
						if service == "" && func() bool {
//line /usr/local/go/src/net/lookup.go:710
		_go_fuzz_dep_.CoverTab[6888]++
//line /usr/local/go/src/net/lookup.go:710
		return proto == ""
//line /usr/local/go/src/net/lookup.go:710
		// _ = "end of CoverTab[6888]"
//line /usr/local/go/src/net/lookup.go:710
	}() {
//line /usr/local/go/src/net/lookup.go:710
		_go_fuzz_dep_.CoverTab[6889]++
							target = name
//line /usr/local/go/src/net/lookup.go:711
		// _ = "end of CoverTab[6889]"
	} else {
//line /usr/local/go/src/net/lookup.go:712
		_go_fuzz_dep_.CoverTab[6890]++
							target = "_" + service + "._" + proto + "." + name
//line /usr/local/go/src/net/lookup.go:713
		// _ = "end of CoverTab[6890]"
	}
//line /usr/local/go/src/net/lookup.go:714
	// _ = "end of CoverTab[6884]"
//line /usr/local/go/src/net/lookup.go:714
	_go_fuzz_dep_.CoverTab[6885]++
						p, server, err := r.lookup(ctx, target, dnsmessage.TypeSRV, nil)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:716
		_go_fuzz_dep_.CoverTab[6891]++
							return "", nil, err
//line /usr/local/go/src/net/lookup.go:717
		// _ = "end of CoverTab[6891]"
	} else {
//line /usr/local/go/src/net/lookup.go:718
		_go_fuzz_dep_.CoverTab[6892]++
//line /usr/local/go/src/net/lookup.go:718
		// _ = "end of CoverTab[6892]"
//line /usr/local/go/src/net/lookup.go:718
	}
//line /usr/local/go/src/net/lookup.go:718
	// _ = "end of CoverTab[6885]"
//line /usr/local/go/src/net/lookup.go:718
	_go_fuzz_dep_.CoverTab[6886]++
						var cname dnsmessage.Name
						for {
//line /usr/local/go/src/net/lookup.go:720
		_go_fuzz_dep_.CoverTab[6893]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/lookup.go:722
			_go_fuzz_dep_.CoverTab[6899]++
								break
//line /usr/local/go/src/net/lookup.go:723
			// _ = "end of CoverTab[6899]"
		} else {
//line /usr/local/go/src/net/lookup.go:724
			_go_fuzz_dep_.CoverTab[6900]++
//line /usr/local/go/src/net/lookup.go:724
			// _ = "end of CoverTab[6900]"
//line /usr/local/go/src/net/lookup.go:724
		}
//line /usr/local/go/src/net/lookup.go:724
		// _ = "end of CoverTab[6893]"
//line /usr/local/go/src/net/lookup.go:724
		_go_fuzz_dep_.CoverTab[6894]++
							if err != nil {
//line /usr/local/go/src/net/lookup.go:725
			_go_fuzz_dep_.CoverTab[6901]++
								return "", nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:730
			// _ = "end of CoverTab[6901]"
		} else {
//line /usr/local/go/src/net/lookup.go:731
			_go_fuzz_dep_.CoverTab[6902]++
//line /usr/local/go/src/net/lookup.go:731
			// _ = "end of CoverTab[6902]"
//line /usr/local/go/src/net/lookup.go:731
		}
//line /usr/local/go/src/net/lookup.go:731
		// _ = "end of CoverTab[6894]"
//line /usr/local/go/src/net/lookup.go:731
		_go_fuzz_dep_.CoverTab[6895]++
							if h.Type != dnsmessage.TypeSRV {
//line /usr/local/go/src/net/lookup.go:732
			_go_fuzz_dep_.CoverTab[6903]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/lookup.go:733
				_go_fuzz_dep_.CoverTab[6905]++
									return "", nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /usr/local/go/src/net/lookup.go:738
				// _ = "end of CoverTab[6905]"
			} else {
//line /usr/local/go/src/net/lookup.go:739
				_go_fuzz_dep_.CoverTab[6906]++
//line /usr/local/go/src/net/lookup.go:739
				// _ = "end of CoverTab[6906]"
//line /usr/local/go/src/net/lookup.go:739
			}
//line /usr/local/go/src/net/lookup.go:739
			// _ = "end of CoverTab[6903]"
//line /usr/local/go/src/net/lookup.go:739
			_go_fuzz_dep_.CoverTab[6904]++
								continue
//line /usr/local/go/src/net/lookup.go:740
			// _ = "end of CoverTab[6904]"
		} else {
//line /usr/local/go/src/net/lookup.go:741
			_go_fuzz_dep_.CoverTab[6907]++
//line /usr/local/go/src/net/lookup.go:741
			// _ = "end of CoverTab[6907]"
//line /usr/local/go/src/net/lookup.go:741
		}
//line /usr/local/go/src/net/lookup.go:741
		// _ = "end of CoverTab[6895]"
//line /usr/local/go/src/net/lookup.go:741
		_go_fuzz_dep_.CoverTab[6896]++
							if cname.Length == 0 && func() bool {
//line /usr/local/go/src/net/lookup.go:742
			_go_fuzz_dep_.CoverTab[6908]++
//line /usr/local/go/src/net/lookup.go:742
			return h.Name.Length != 0
//line /usr/local/go/src/net/lookup.go:742
			// _ = "end of CoverTab[6908]"
//line /usr/local/go/src/net/lookup.go:742
		}() {
//line /usr/local/go/src/net/lookup.go:742
			_go_fuzz_dep_.CoverTab[6909]++
								cname = h.Name
//line /usr/local/go/src/net/lookup.go:743
			// _ = "end of CoverTab[6909]"
		} else {
//line /usr/local/go/src/net/lookup.go:744
			_go_fuzz_dep_.CoverTab[6910]++
//line /usr/local/go/src/net/lookup.go:744
			// _ = "end of CoverTab[6910]"
//line /usr/local/go/src/net/lookup.go:744
		}
//line /usr/local/go/src/net/lookup.go:744
		// _ = "end of CoverTab[6896]"
//line /usr/local/go/src/net/lookup.go:744
		_go_fuzz_dep_.CoverTab[6897]++
							srv, err := p.SRVResource()
							if err != nil {
//line /usr/local/go/src/net/lookup.go:746
			_go_fuzz_dep_.CoverTab[6911]++
								return "", nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:751
			// _ = "end of CoverTab[6911]"
		} else {
//line /usr/local/go/src/net/lookup.go:752
			_go_fuzz_dep_.CoverTab[6912]++
//line /usr/local/go/src/net/lookup.go:752
			// _ = "end of CoverTab[6912]"
//line /usr/local/go/src/net/lookup.go:752
		}
//line /usr/local/go/src/net/lookup.go:752
		// _ = "end of CoverTab[6897]"
//line /usr/local/go/src/net/lookup.go:752
		_go_fuzz_dep_.CoverTab[6898]++
							srvs = append(srvs, &SRV{Target: srv.Target.String(), Port: srv.Port, Priority: srv.Priority, Weight: srv.Weight})
//line /usr/local/go/src/net/lookup.go:753
		// _ = "end of CoverTab[6898]"
	}
//line /usr/local/go/src/net/lookup.go:754
	// _ = "end of CoverTab[6886]"
//line /usr/local/go/src/net/lookup.go:754
	_go_fuzz_dep_.CoverTab[6887]++
						byPriorityWeight(srvs).sort()
						return cname.String(), srvs, nil
//line /usr/local/go/src/net/lookup.go:756
	// _ = "end of CoverTab[6887]"
}

// goLookupMX returns the MX records for name.
func (r *Resolver) goLookupMX(ctx context.Context, name string) ([]*MX, error) {
//line /usr/local/go/src/net/lookup.go:760
	_go_fuzz_dep_.CoverTab[6913]++
						p, server, err := r.lookup(ctx, name, dnsmessage.TypeMX, nil)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:762
		_go_fuzz_dep_.CoverTab[6916]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:763
		// _ = "end of CoverTab[6916]"
	} else {
//line /usr/local/go/src/net/lookup.go:764
		_go_fuzz_dep_.CoverTab[6917]++
//line /usr/local/go/src/net/lookup.go:764
		// _ = "end of CoverTab[6917]"
//line /usr/local/go/src/net/lookup.go:764
	}
//line /usr/local/go/src/net/lookup.go:764
	// _ = "end of CoverTab[6913]"
//line /usr/local/go/src/net/lookup.go:764
	_go_fuzz_dep_.CoverTab[6914]++
						var mxs []*MX
						for {
//line /usr/local/go/src/net/lookup.go:766
		_go_fuzz_dep_.CoverTab[6918]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/lookup.go:768
			_go_fuzz_dep_.CoverTab[6923]++
								break
//line /usr/local/go/src/net/lookup.go:769
			// _ = "end of CoverTab[6923]"
		} else {
//line /usr/local/go/src/net/lookup.go:770
			_go_fuzz_dep_.CoverTab[6924]++
//line /usr/local/go/src/net/lookup.go:770
			// _ = "end of CoverTab[6924]"
//line /usr/local/go/src/net/lookup.go:770
		}
//line /usr/local/go/src/net/lookup.go:770
		// _ = "end of CoverTab[6918]"
//line /usr/local/go/src/net/lookup.go:770
		_go_fuzz_dep_.CoverTab[6919]++
							if err != nil {
//line /usr/local/go/src/net/lookup.go:771
			_go_fuzz_dep_.CoverTab[6925]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:776
			// _ = "end of CoverTab[6925]"
		} else {
//line /usr/local/go/src/net/lookup.go:777
			_go_fuzz_dep_.CoverTab[6926]++
//line /usr/local/go/src/net/lookup.go:777
			// _ = "end of CoverTab[6926]"
//line /usr/local/go/src/net/lookup.go:777
		}
//line /usr/local/go/src/net/lookup.go:777
		// _ = "end of CoverTab[6919]"
//line /usr/local/go/src/net/lookup.go:777
		_go_fuzz_dep_.CoverTab[6920]++
							if h.Type != dnsmessage.TypeMX {
//line /usr/local/go/src/net/lookup.go:778
			_go_fuzz_dep_.CoverTab[6927]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/lookup.go:779
				_go_fuzz_dep_.CoverTab[6929]++
									return nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /usr/local/go/src/net/lookup.go:784
				// _ = "end of CoverTab[6929]"
			} else {
//line /usr/local/go/src/net/lookup.go:785
				_go_fuzz_dep_.CoverTab[6930]++
//line /usr/local/go/src/net/lookup.go:785
				// _ = "end of CoverTab[6930]"
//line /usr/local/go/src/net/lookup.go:785
			}
//line /usr/local/go/src/net/lookup.go:785
			// _ = "end of CoverTab[6927]"
//line /usr/local/go/src/net/lookup.go:785
			_go_fuzz_dep_.CoverTab[6928]++
								continue
//line /usr/local/go/src/net/lookup.go:786
			// _ = "end of CoverTab[6928]"
		} else {
//line /usr/local/go/src/net/lookup.go:787
			_go_fuzz_dep_.CoverTab[6931]++
//line /usr/local/go/src/net/lookup.go:787
			// _ = "end of CoverTab[6931]"
//line /usr/local/go/src/net/lookup.go:787
		}
//line /usr/local/go/src/net/lookup.go:787
		// _ = "end of CoverTab[6920]"
//line /usr/local/go/src/net/lookup.go:787
		_go_fuzz_dep_.CoverTab[6921]++
							mx, err := p.MXResource()
							if err != nil {
//line /usr/local/go/src/net/lookup.go:789
			_go_fuzz_dep_.CoverTab[6932]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:794
			// _ = "end of CoverTab[6932]"
		} else {
//line /usr/local/go/src/net/lookup.go:795
			_go_fuzz_dep_.CoverTab[6933]++
//line /usr/local/go/src/net/lookup.go:795
			// _ = "end of CoverTab[6933]"
//line /usr/local/go/src/net/lookup.go:795
		}
//line /usr/local/go/src/net/lookup.go:795
		// _ = "end of CoverTab[6921]"
//line /usr/local/go/src/net/lookup.go:795
		_go_fuzz_dep_.CoverTab[6922]++
							mxs = append(mxs, &MX{Host: mx.MX.String(), Pref: mx.Pref})
//line /usr/local/go/src/net/lookup.go:796
		// _ = "end of CoverTab[6922]"

	}
//line /usr/local/go/src/net/lookup.go:798
	// _ = "end of CoverTab[6914]"
//line /usr/local/go/src/net/lookup.go:798
	_go_fuzz_dep_.CoverTab[6915]++
						byPref(mxs).sort()
						return mxs, nil
//line /usr/local/go/src/net/lookup.go:800
	// _ = "end of CoverTab[6915]"
}

// goLookupNS returns the NS records for name.
func (r *Resolver) goLookupNS(ctx context.Context, name string) ([]*NS, error) {
//line /usr/local/go/src/net/lookup.go:804
	_go_fuzz_dep_.CoverTab[6934]++
						p, server, err := r.lookup(ctx, name, dnsmessage.TypeNS, nil)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:806
		_go_fuzz_dep_.CoverTab[6937]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:807
		// _ = "end of CoverTab[6937]"
	} else {
//line /usr/local/go/src/net/lookup.go:808
		_go_fuzz_dep_.CoverTab[6938]++
//line /usr/local/go/src/net/lookup.go:808
		// _ = "end of CoverTab[6938]"
//line /usr/local/go/src/net/lookup.go:808
	}
//line /usr/local/go/src/net/lookup.go:808
	// _ = "end of CoverTab[6934]"
//line /usr/local/go/src/net/lookup.go:808
	_go_fuzz_dep_.CoverTab[6935]++
						var nss []*NS
						for {
//line /usr/local/go/src/net/lookup.go:810
		_go_fuzz_dep_.CoverTab[6939]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/lookup.go:812
			_go_fuzz_dep_.CoverTab[6944]++
								break
//line /usr/local/go/src/net/lookup.go:813
			// _ = "end of CoverTab[6944]"
		} else {
//line /usr/local/go/src/net/lookup.go:814
			_go_fuzz_dep_.CoverTab[6945]++
//line /usr/local/go/src/net/lookup.go:814
			// _ = "end of CoverTab[6945]"
//line /usr/local/go/src/net/lookup.go:814
		}
//line /usr/local/go/src/net/lookup.go:814
		// _ = "end of CoverTab[6939]"
//line /usr/local/go/src/net/lookup.go:814
		_go_fuzz_dep_.CoverTab[6940]++
							if err != nil {
//line /usr/local/go/src/net/lookup.go:815
			_go_fuzz_dep_.CoverTab[6946]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:820
			// _ = "end of CoverTab[6946]"
		} else {
//line /usr/local/go/src/net/lookup.go:821
			_go_fuzz_dep_.CoverTab[6947]++
//line /usr/local/go/src/net/lookup.go:821
			// _ = "end of CoverTab[6947]"
//line /usr/local/go/src/net/lookup.go:821
		}
//line /usr/local/go/src/net/lookup.go:821
		// _ = "end of CoverTab[6940]"
//line /usr/local/go/src/net/lookup.go:821
		_go_fuzz_dep_.CoverTab[6941]++
							if h.Type != dnsmessage.TypeNS {
//line /usr/local/go/src/net/lookup.go:822
			_go_fuzz_dep_.CoverTab[6948]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/lookup.go:823
				_go_fuzz_dep_.CoverTab[6950]++
									return nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /usr/local/go/src/net/lookup.go:828
				// _ = "end of CoverTab[6950]"
			} else {
//line /usr/local/go/src/net/lookup.go:829
				_go_fuzz_dep_.CoverTab[6951]++
//line /usr/local/go/src/net/lookup.go:829
				// _ = "end of CoverTab[6951]"
//line /usr/local/go/src/net/lookup.go:829
			}
//line /usr/local/go/src/net/lookup.go:829
			// _ = "end of CoverTab[6948]"
//line /usr/local/go/src/net/lookup.go:829
			_go_fuzz_dep_.CoverTab[6949]++
								continue
//line /usr/local/go/src/net/lookup.go:830
			// _ = "end of CoverTab[6949]"
		} else {
//line /usr/local/go/src/net/lookup.go:831
			_go_fuzz_dep_.CoverTab[6952]++
//line /usr/local/go/src/net/lookup.go:831
			// _ = "end of CoverTab[6952]"
//line /usr/local/go/src/net/lookup.go:831
		}
//line /usr/local/go/src/net/lookup.go:831
		// _ = "end of CoverTab[6941]"
//line /usr/local/go/src/net/lookup.go:831
		_go_fuzz_dep_.CoverTab[6942]++
							ns, err := p.NSResource()
							if err != nil {
//line /usr/local/go/src/net/lookup.go:833
			_go_fuzz_dep_.CoverTab[6953]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:838
			// _ = "end of CoverTab[6953]"
		} else {
//line /usr/local/go/src/net/lookup.go:839
			_go_fuzz_dep_.CoverTab[6954]++
//line /usr/local/go/src/net/lookup.go:839
			// _ = "end of CoverTab[6954]"
//line /usr/local/go/src/net/lookup.go:839
		}
//line /usr/local/go/src/net/lookup.go:839
		// _ = "end of CoverTab[6942]"
//line /usr/local/go/src/net/lookup.go:839
		_go_fuzz_dep_.CoverTab[6943]++
							nss = append(nss, &NS{Host: ns.NS.String()})
//line /usr/local/go/src/net/lookup.go:840
		// _ = "end of CoverTab[6943]"
	}
//line /usr/local/go/src/net/lookup.go:841
	// _ = "end of CoverTab[6935]"
//line /usr/local/go/src/net/lookup.go:841
	_go_fuzz_dep_.CoverTab[6936]++
						return nss, nil
//line /usr/local/go/src/net/lookup.go:842
	// _ = "end of CoverTab[6936]"
}

// goLookupTXT returns the TXT records from name.
func (r *Resolver) goLookupTXT(ctx context.Context, name string) ([]string, error) {
//line /usr/local/go/src/net/lookup.go:846
	_go_fuzz_dep_.CoverTab[6955]++
						p, server, err := r.lookup(ctx, name, dnsmessage.TypeTXT, nil)
						if err != nil {
//line /usr/local/go/src/net/lookup.go:848
		_go_fuzz_dep_.CoverTab[6958]++
							return nil, err
//line /usr/local/go/src/net/lookup.go:849
		// _ = "end of CoverTab[6958]"
	} else {
//line /usr/local/go/src/net/lookup.go:850
		_go_fuzz_dep_.CoverTab[6959]++
//line /usr/local/go/src/net/lookup.go:850
		// _ = "end of CoverTab[6959]"
//line /usr/local/go/src/net/lookup.go:850
	}
//line /usr/local/go/src/net/lookup.go:850
	// _ = "end of CoverTab[6955]"
//line /usr/local/go/src/net/lookup.go:850
	_go_fuzz_dep_.CoverTab[6956]++
						var txts []string
						for {
//line /usr/local/go/src/net/lookup.go:852
		_go_fuzz_dep_.CoverTab[6960]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /usr/local/go/src/net/lookup.go:854
			_go_fuzz_dep_.CoverTab[6968]++
								break
//line /usr/local/go/src/net/lookup.go:855
			// _ = "end of CoverTab[6968]"
		} else {
//line /usr/local/go/src/net/lookup.go:856
			_go_fuzz_dep_.CoverTab[6969]++
//line /usr/local/go/src/net/lookup.go:856
			// _ = "end of CoverTab[6969]"
//line /usr/local/go/src/net/lookup.go:856
		}
//line /usr/local/go/src/net/lookup.go:856
		// _ = "end of CoverTab[6960]"
//line /usr/local/go/src/net/lookup.go:856
		_go_fuzz_dep_.CoverTab[6961]++
							if err != nil {
//line /usr/local/go/src/net/lookup.go:857
			_go_fuzz_dep_.CoverTab[6970]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:862
			// _ = "end of CoverTab[6970]"
		} else {
//line /usr/local/go/src/net/lookup.go:863
			_go_fuzz_dep_.CoverTab[6971]++
//line /usr/local/go/src/net/lookup.go:863
			// _ = "end of CoverTab[6971]"
//line /usr/local/go/src/net/lookup.go:863
		}
//line /usr/local/go/src/net/lookup.go:863
		// _ = "end of CoverTab[6961]"
//line /usr/local/go/src/net/lookup.go:863
		_go_fuzz_dep_.CoverTab[6962]++
							if h.Type != dnsmessage.TypeTXT {
//line /usr/local/go/src/net/lookup.go:864
			_go_fuzz_dep_.CoverTab[6972]++
								if err := p.SkipAnswer(); err != nil {
//line /usr/local/go/src/net/lookup.go:865
				_go_fuzz_dep_.CoverTab[6974]++
									return nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /usr/local/go/src/net/lookup.go:870
				// _ = "end of CoverTab[6974]"
			} else {
//line /usr/local/go/src/net/lookup.go:871
				_go_fuzz_dep_.CoverTab[6975]++
//line /usr/local/go/src/net/lookup.go:871
				// _ = "end of CoverTab[6975]"
//line /usr/local/go/src/net/lookup.go:871
			}
//line /usr/local/go/src/net/lookup.go:871
			// _ = "end of CoverTab[6972]"
//line /usr/local/go/src/net/lookup.go:871
			_go_fuzz_dep_.CoverTab[6973]++
								continue
//line /usr/local/go/src/net/lookup.go:872
			// _ = "end of CoverTab[6973]"
		} else {
//line /usr/local/go/src/net/lookup.go:873
			_go_fuzz_dep_.CoverTab[6976]++
//line /usr/local/go/src/net/lookup.go:873
			// _ = "end of CoverTab[6976]"
//line /usr/local/go/src/net/lookup.go:873
		}
//line /usr/local/go/src/net/lookup.go:873
		// _ = "end of CoverTab[6962]"
//line /usr/local/go/src/net/lookup.go:873
		_go_fuzz_dep_.CoverTab[6963]++
							txt, err := p.TXTResource()
							if err != nil {
//line /usr/local/go/src/net/lookup.go:875
			_go_fuzz_dep_.CoverTab[6977]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /usr/local/go/src/net/lookup.go:880
			// _ = "end of CoverTab[6977]"
		} else {
//line /usr/local/go/src/net/lookup.go:881
			_go_fuzz_dep_.CoverTab[6978]++
//line /usr/local/go/src/net/lookup.go:881
			// _ = "end of CoverTab[6978]"
//line /usr/local/go/src/net/lookup.go:881
		}
//line /usr/local/go/src/net/lookup.go:881
		// _ = "end of CoverTab[6963]"
//line /usr/local/go/src/net/lookup.go:881
		_go_fuzz_dep_.CoverTab[6964]++

//line /usr/local/go/src/net/lookup.go:885
		n := 0
		for _, s := range txt.TXT {
//line /usr/local/go/src/net/lookup.go:886
			_go_fuzz_dep_.CoverTab[6979]++
								n += len(s)
//line /usr/local/go/src/net/lookup.go:887
			// _ = "end of CoverTab[6979]"
		}
//line /usr/local/go/src/net/lookup.go:888
		// _ = "end of CoverTab[6964]"
//line /usr/local/go/src/net/lookup.go:888
		_go_fuzz_dep_.CoverTab[6965]++
							txtJoin := make([]byte, 0, n)
							for _, s := range txt.TXT {
//line /usr/local/go/src/net/lookup.go:890
			_go_fuzz_dep_.CoverTab[6980]++
								txtJoin = append(txtJoin, s...)
//line /usr/local/go/src/net/lookup.go:891
			// _ = "end of CoverTab[6980]"
		}
//line /usr/local/go/src/net/lookup.go:892
		// _ = "end of CoverTab[6965]"
//line /usr/local/go/src/net/lookup.go:892
		_go_fuzz_dep_.CoverTab[6966]++
							if len(txts) == 0 {
//line /usr/local/go/src/net/lookup.go:893
			_go_fuzz_dep_.CoverTab[6981]++
								txts = make([]string, 0, 1)
//line /usr/local/go/src/net/lookup.go:894
			// _ = "end of CoverTab[6981]"
		} else {
//line /usr/local/go/src/net/lookup.go:895
			_go_fuzz_dep_.CoverTab[6982]++
//line /usr/local/go/src/net/lookup.go:895
			// _ = "end of CoverTab[6982]"
//line /usr/local/go/src/net/lookup.go:895
		}
//line /usr/local/go/src/net/lookup.go:895
		// _ = "end of CoverTab[6966]"
//line /usr/local/go/src/net/lookup.go:895
		_go_fuzz_dep_.CoverTab[6967]++
							txts = append(txts, string(txtJoin))
//line /usr/local/go/src/net/lookup.go:896
		// _ = "end of CoverTab[6967]"
	}
//line /usr/local/go/src/net/lookup.go:897
	// _ = "end of CoverTab[6956]"
//line /usr/local/go/src/net/lookup.go:897
	_go_fuzz_dep_.CoverTab[6957]++
						return txts, nil
//line /usr/local/go/src/net/lookup.go:898
	// _ = "end of CoverTab[6957]"
}

func parseCNAMEFromResources(resources []dnsmessage.Resource) (string, error) {
//line /usr/local/go/src/net/lookup.go:901
	_go_fuzz_dep_.CoverTab[6983]++
						if len(resources) == 0 {
//line /usr/local/go/src/net/lookup.go:902
		_go_fuzz_dep_.CoverTab[6986]++
							return "", errors.New("no CNAME record received")
//line /usr/local/go/src/net/lookup.go:903
		// _ = "end of CoverTab[6986]"
	} else {
//line /usr/local/go/src/net/lookup.go:904
		_go_fuzz_dep_.CoverTab[6987]++
//line /usr/local/go/src/net/lookup.go:904
		// _ = "end of CoverTab[6987]"
//line /usr/local/go/src/net/lookup.go:904
	}
//line /usr/local/go/src/net/lookup.go:904
	// _ = "end of CoverTab[6983]"
//line /usr/local/go/src/net/lookup.go:904
	_go_fuzz_dep_.CoverTab[6984]++
						c, ok := resources[0].Body.(*dnsmessage.CNAMEResource)
						if !ok {
//line /usr/local/go/src/net/lookup.go:906
		_go_fuzz_dep_.CoverTab[6988]++
							return "", errors.New("could not parse CNAME record")
//line /usr/local/go/src/net/lookup.go:907
		// _ = "end of CoverTab[6988]"
	} else {
//line /usr/local/go/src/net/lookup.go:908
		_go_fuzz_dep_.CoverTab[6989]++
//line /usr/local/go/src/net/lookup.go:908
		// _ = "end of CoverTab[6989]"
//line /usr/local/go/src/net/lookup.go:908
	}
//line /usr/local/go/src/net/lookup.go:908
	// _ = "end of CoverTab[6984]"
//line /usr/local/go/src/net/lookup.go:908
	_go_fuzz_dep_.CoverTab[6985]++
						return c.CNAME.String(), nil
//line /usr/local/go/src/net/lookup.go:909
	// _ = "end of CoverTab[6985]"
}

//line /usr/local/go/src/net/lookup.go:910
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/lookup.go:910
var _ = _go_fuzz_dep_.CoverTab
