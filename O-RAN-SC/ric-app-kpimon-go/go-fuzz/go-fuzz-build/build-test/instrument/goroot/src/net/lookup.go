// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/lookup.go:5
package net

//line /snap/go/10455/src/net/lookup.go:5
import (
//line /snap/go/10455/src/net/lookup.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/lookup.go:5
)
//line /snap/go/10455/src/net/lookup.go:5
import (
//line /snap/go/10455/src/net/lookup.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/lookup.go:5
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
//line /snap/go/10455/src/net/lookup.go:18
// names and numbers for platforms that don't have a complete list of
//line /snap/go/10455/src/net/lookup.go:18
// protocol numbers.
//line /snap/go/10455/src/net/lookup.go:18
//
//line /snap/go/10455/src/net/lookup.go:18
// See https://www.iana.org/assignments/protocol-numbers
//line /snap/go/10455/src/net/lookup.go:18
//
//line /snap/go/10455/src/net/lookup.go:18
// On Unix, this map is augmented by readProtocols via lookupProtocol.
//line /snap/go/10455/src/net/lookup.go:25
var protocols = map[string]int{
	"icmp":		1,
	"igmp":		2,
	"tcp":		6,
	"udp":		17,
	"ipv6-icmp":	58,
}

// services contains minimal mappings between services names and port
//line /snap/go/10455/src/net/lookup.go:33
// numbers for platforms that don't have a complete list of port numbers.
//line /snap/go/10455/src/net/lookup.go:33
//
//line /snap/go/10455/src/net/lookup.go:33
// See https://www.iana.org/assignments/service-names-port-numbers
//line /snap/go/10455/src/net/lookup.go:33
//
//line /snap/go/10455/src/net/lookup.go:33
// On Unix, this map is augmented by readServices via goLookupPort.
//line /snap/go/10455/src/net/lookup.go:39
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
//line /snap/go/10455/src/net/lookup.go:60
// complete. This avoids races on the test hooks.
//line /snap/go/10455/src/net/lookup.go:62
var dnsWaitGroup sync.WaitGroup

const maxProtoLength = len("RSVP-E2E-IGNORE") + 10	// with room to grow

func lookupProtocolMap(name string) (int, error) {
//line /snap/go/10455/src/net/lookup.go:66
	_go_fuzz_dep_.CoverTab[6923]++
						var lowerProtocol [maxProtoLength]byte
						n := copy(lowerProtocol[:], name)
						lowerASCIIBytes(lowerProtocol[:n])
						proto, found := protocols[string(lowerProtocol[:n])]
						if !found || func() bool {
//line /snap/go/10455/src/net/lookup.go:71
		_go_fuzz_dep_.CoverTab[6925]++
//line /snap/go/10455/src/net/lookup.go:71
		return n != len(name)
//line /snap/go/10455/src/net/lookup.go:71
		// _ = "end of CoverTab[6925]"
//line /snap/go/10455/src/net/lookup.go:71
	}() {
//line /snap/go/10455/src/net/lookup.go:71
		_go_fuzz_dep_.CoverTab[528924]++
//line /snap/go/10455/src/net/lookup.go:71
		_go_fuzz_dep_.CoverTab[6926]++
							return 0, &AddrError{Err: "unknown IP protocol specified", Addr: name}
//line /snap/go/10455/src/net/lookup.go:72
		// _ = "end of CoverTab[6926]"
	} else {
//line /snap/go/10455/src/net/lookup.go:73
		_go_fuzz_dep_.CoverTab[528925]++
//line /snap/go/10455/src/net/lookup.go:73
		_go_fuzz_dep_.CoverTab[6927]++
//line /snap/go/10455/src/net/lookup.go:73
		// _ = "end of CoverTab[6927]"
//line /snap/go/10455/src/net/lookup.go:73
	}
//line /snap/go/10455/src/net/lookup.go:73
	// _ = "end of CoverTab[6923]"
//line /snap/go/10455/src/net/lookup.go:73
	_go_fuzz_dep_.CoverTab[6924]++
						return proto, nil
//line /snap/go/10455/src/net/lookup.go:74
	// _ = "end of CoverTab[6924]"
}

// maxPortBufSize is the longest reasonable name of a service
//line /snap/go/10455/src/net/lookup.go:77
// (non-numeric port).
//line /snap/go/10455/src/net/lookup.go:77
// Currently the longest known IANA-unregistered name is
//line /snap/go/10455/src/net/lookup.go:77
// "mobility-header", so we use that length, plus some slop in case
//line /snap/go/10455/src/net/lookup.go:77
// something longer is added in the future.
//line /snap/go/10455/src/net/lookup.go:82
const maxPortBufSize = len("mobility-header") + 10

func lookupPortMap(network, service string) (port int, error error) {
//line /snap/go/10455/src/net/lookup.go:84
	_go_fuzz_dep_.CoverTab[6928]++
						switch network {
	case "tcp4", "tcp6":
//line /snap/go/10455/src/net/lookup.go:86
		_go_fuzz_dep_.CoverTab[528926]++
//line /snap/go/10455/src/net/lookup.go:86
		_go_fuzz_dep_.CoverTab[6931]++
							network = "tcp"
//line /snap/go/10455/src/net/lookup.go:87
		// _ = "end of CoverTab[6931]"
	case "udp4", "udp6":
//line /snap/go/10455/src/net/lookup.go:88
		_go_fuzz_dep_.CoverTab[528927]++
//line /snap/go/10455/src/net/lookup.go:88
		_go_fuzz_dep_.CoverTab[6932]++
							network = "udp"
//line /snap/go/10455/src/net/lookup.go:89
		// _ = "end of CoverTab[6932]"
//line /snap/go/10455/src/net/lookup.go:89
	default:
//line /snap/go/10455/src/net/lookup.go:89
		_go_fuzz_dep_.CoverTab[528928]++
//line /snap/go/10455/src/net/lookup.go:89
		_go_fuzz_dep_.CoverTab[6933]++
//line /snap/go/10455/src/net/lookup.go:89
		// _ = "end of CoverTab[6933]"
	}
//line /snap/go/10455/src/net/lookup.go:90
	// _ = "end of CoverTab[6928]"
//line /snap/go/10455/src/net/lookup.go:90
	_go_fuzz_dep_.CoverTab[6929]++

						if m, ok := services[network]; ok {
//line /snap/go/10455/src/net/lookup.go:92
		_go_fuzz_dep_.CoverTab[528929]++
//line /snap/go/10455/src/net/lookup.go:92
		_go_fuzz_dep_.CoverTab[6934]++
							var lowerService [maxPortBufSize]byte
							n := copy(lowerService[:], service)
							lowerASCIIBytes(lowerService[:n])
							if port, ok := m[string(lowerService[:n])]; ok && func() bool {
//line /snap/go/10455/src/net/lookup.go:96
			_go_fuzz_dep_.CoverTab[6935]++
//line /snap/go/10455/src/net/lookup.go:96
			return n == len(service)
//line /snap/go/10455/src/net/lookup.go:96
			// _ = "end of CoverTab[6935]"
//line /snap/go/10455/src/net/lookup.go:96
		}() {
//line /snap/go/10455/src/net/lookup.go:96
			_go_fuzz_dep_.CoverTab[528931]++
//line /snap/go/10455/src/net/lookup.go:96
			_go_fuzz_dep_.CoverTab[6936]++
								return port, nil
//line /snap/go/10455/src/net/lookup.go:97
			// _ = "end of CoverTab[6936]"
		} else {
//line /snap/go/10455/src/net/lookup.go:98
			_go_fuzz_dep_.CoverTab[528932]++
//line /snap/go/10455/src/net/lookup.go:98
			_go_fuzz_dep_.CoverTab[6937]++
//line /snap/go/10455/src/net/lookup.go:98
			// _ = "end of CoverTab[6937]"
//line /snap/go/10455/src/net/lookup.go:98
		}
//line /snap/go/10455/src/net/lookup.go:98
		// _ = "end of CoverTab[6934]"
	} else {
//line /snap/go/10455/src/net/lookup.go:99
		_go_fuzz_dep_.CoverTab[528930]++
//line /snap/go/10455/src/net/lookup.go:99
		_go_fuzz_dep_.CoverTab[6938]++
//line /snap/go/10455/src/net/lookup.go:99
		// _ = "end of CoverTab[6938]"
//line /snap/go/10455/src/net/lookup.go:99
	}
//line /snap/go/10455/src/net/lookup.go:99
	// _ = "end of CoverTab[6929]"
//line /snap/go/10455/src/net/lookup.go:99
	_go_fuzz_dep_.CoverTab[6930]++
						return 0, &AddrError{Err: "unknown port", Addr: network + "/" + service}
//line /snap/go/10455/src/net/lookup.go:100
	// _ = "end of CoverTab[6930]"
}

// ipVersion returns the provided network's IP version: '4', '6' or 0
//line /snap/go/10455/src/net/lookup.go:103
// if network does not end in a '4' or '6' byte.
//line /snap/go/10455/src/net/lookup.go:105
func ipVersion(network string) byte {
//line /snap/go/10455/src/net/lookup.go:105
	_go_fuzz_dep_.CoverTab[6939]++
						if network == "" {
//line /snap/go/10455/src/net/lookup.go:106
		_go_fuzz_dep_.CoverTab[528933]++
//line /snap/go/10455/src/net/lookup.go:106
		_go_fuzz_dep_.CoverTab[6942]++
							return 0
//line /snap/go/10455/src/net/lookup.go:107
		// _ = "end of CoverTab[6942]"
	} else {
//line /snap/go/10455/src/net/lookup.go:108
		_go_fuzz_dep_.CoverTab[528934]++
//line /snap/go/10455/src/net/lookup.go:108
		_go_fuzz_dep_.CoverTab[6943]++
//line /snap/go/10455/src/net/lookup.go:108
		// _ = "end of CoverTab[6943]"
//line /snap/go/10455/src/net/lookup.go:108
	}
//line /snap/go/10455/src/net/lookup.go:108
	// _ = "end of CoverTab[6939]"
//line /snap/go/10455/src/net/lookup.go:108
	_go_fuzz_dep_.CoverTab[6940]++
						n := network[len(network)-1]
						if n != '4' && func() bool {
//line /snap/go/10455/src/net/lookup.go:110
		_go_fuzz_dep_.CoverTab[6944]++
//line /snap/go/10455/src/net/lookup.go:110
		return n != '6'
//line /snap/go/10455/src/net/lookup.go:110
		// _ = "end of CoverTab[6944]"
//line /snap/go/10455/src/net/lookup.go:110
	}() {
//line /snap/go/10455/src/net/lookup.go:110
		_go_fuzz_dep_.CoverTab[528935]++
//line /snap/go/10455/src/net/lookup.go:110
		_go_fuzz_dep_.CoverTab[6945]++
							n = 0
//line /snap/go/10455/src/net/lookup.go:111
		// _ = "end of CoverTab[6945]"
	} else {
//line /snap/go/10455/src/net/lookup.go:112
		_go_fuzz_dep_.CoverTab[528936]++
//line /snap/go/10455/src/net/lookup.go:112
		_go_fuzz_dep_.CoverTab[6946]++
//line /snap/go/10455/src/net/lookup.go:112
		// _ = "end of CoverTab[6946]"
//line /snap/go/10455/src/net/lookup.go:112
	}
//line /snap/go/10455/src/net/lookup.go:112
	// _ = "end of CoverTab[6940]"
//line /snap/go/10455/src/net/lookup.go:112
	_go_fuzz_dep_.CoverTab[6941]++
						return n
//line /snap/go/10455/src/net/lookup.go:113
	// _ = "end of CoverTab[6941]"
}

// DefaultResolver is the resolver used by the package-level Lookup
//line /snap/go/10455/src/net/lookup.go:116
// functions and by Dialers without a specified Resolver.
//line /snap/go/10455/src/net/lookup.go:118
var DefaultResolver = &Resolver{}

// A Resolver looks up names and numbers.
//line /snap/go/10455/src/net/lookup.go:120
//
//line /snap/go/10455/src/net/lookup.go:120
// A nil *Resolver is equivalent to a zero Resolver.
//line /snap/go/10455/src/net/lookup.go:123
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
//line /snap/go/10455/src/net/lookup.go:159
}

func (r *Resolver) preferGo() bool {
//line /snap/go/10455/src/net/lookup.go:161
	_go_fuzz_dep_.CoverTab[6947]++
//line /snap/go/10455/src/net/lookup.go:161
	return r != nil && func() bool {
//line /snap/go/10455/src/net/lookup.go:161
		_go_fuzz_dep_.CoverTab[6948]++
//line /snap/go/10455/src/net/lookup.go:161
		return r.PreferGo
//line /snap/go/10455/src/net/lookup.go:161
		// _ = "end of CoverTab[6948]"
//line /snap/go/10455/src/net/lookup.go:161
	}()
//line /snap/go/10455/src/net/lookup.go:161
	// _ = "end of CoverTab[6947]"
//line /snap/go/10455/src/net/lookup.go:161
}
func (r *Resolver) strictErrors() bool {
//line /snap/go/10455/src/net/lookup.go:162
	_go_fuzz_dep_.CoverTab[6949]++
//line /snap/go/10455/src/net/lookup.go:162
	return r != nil && func() bool {
//line /snap/go/10455/src/net/lookup.go:162
		_go_fuzz_dep_.CoverTab[6950]++
//line /snap/go/10455/src/net/lookup.go:162
		return r.StrictErrors
//line /snap/go/10455/src/net/lookup.go:162
		// _ = "end of CoverTab[6950]"
//line /snap/go/10455/src/net/lookup.go:162
	}()
//line /snap/go/10455/src/net/lookup.go:162
	// _ = "end of CoverTab[6949]"
//line /snap/go/10455/src/net/lookup.go:162
}

func (r *Resolver) getLookupGroup() *singleflight.Group {
//line /snap/go/10455/src/net/lookup.go:164
	_go_fuzz_dep_.CoverTab[6951]++
						if r == nil {
//line /snap/go/10455/src/net/lookup.go:165
		_go_fuzz_dep_.CoverTab[528937]++
//line /snap/go/10455/src/net/lookup.go:165
		_go_fuzz_dep_.CoverTab[6953]++
							return &DefaultResolver.lookupGroup
//line /snap/go/10455/src/net/lookup.go:166
		// _ = "end of CoverTab[6953]"
	} else {
//line /snap/go/10455/src/net/lookup.go:167
		_go_fuzz_dep_.CoverTab[528938]++
//line /snap/go/10455/src/net/lookup.go:167
		_go_fuzz_dep_.CoverTab[6954]++
//line /snap/go/10455/src/net/lookup.go:167
		// _ = "end of CoverTab[6954]"
//line /snap/go/10455/src/net/lookup.go:167
	}
//line /snap/go/10455/src/net/lookup.go:167
	// _ = "end of CoverTab[6951]"
//line /snap/go/10455/src/net/lookup.go:167
	_go_fuzz_dep_.CoverTab[6952]++
						return &r.lookupGroup
//line /snap/go/10455/src/net/lookup.go:168
	// _ = "end of CoverTab[6952]"
}

// LookupHost looks up the given host using the local resolver.
//line /snap/go/10455/src/net/lookup.go:171
// It returns a slice of that host's addresses.
//line /snap/go/10455/src/net/lookup.go:171
//
//line /snap/go/10455/src/net/lookup.go:171
// LookupHost uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/lookup.go:171
// Resolver.LookupHost.
//line /snap/go/10455/src/net/lookup.go:176
func LookupHost(host string) (addrs []string, err error) {
//line /snap/go/10455/src/net/lookup.go:176
	_go_fuzz_dep_.CoverTab[6955]++
						return DefaultResolver.LookupHost(context.Background(), host)
//line /snap/go/10455/src/net/lookup.go:177
	// _ = "end of CoverTab[6955]"
}

// LookupHost looks up the given host using the local resolver.
//line /snap/go/10455/src/net/lookup.go:180
// It returns a slice of that host's addresses.
//line /snap/go/10455/src/net/lookup.go:182
func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error) {
//line /snap/go/10455/src/net/lookup.go:182
	_go_fuzz_dep_.CoverTab[6956]++

						if host == "" {
//line /snap/go/10455/src/net/lookup.go:184
		_go_fuzz_dep_.CoverTab[528939]++
//line /snap/go/10455/src/net/lookup.go:184
		_go_fuzz_dep_.CoverTab[6959]++
							return nil, &DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
//line /snap/go/10455/src/net/lookup.go:185
		// _ = "end of CoverTab[6959]"
	} else {
//line /snap/go/10455/src/net/lookup.go:186
		_go_fuzz_dep_.CoverTab[528940]++
//line /snap/go/10455/src/net/lookup.go:186
		_go_fuzz_dep_.CoverTab[6960]++
//line /snap/go/10455/src/net/lookup.go:186
		// _ = "end of CoverTab[6960]"
//line /snap/go/10455/src/net/lookup.go:186
	}
//line /snap/go/10455/src/net/lookup.go:186
	// _ = "end of CoverTab[6956]"
//line /snap/go/10455/src/net/lookup.go:186
	_go_fuzz_dep_.CoverTab[6957]++
						if _, err := netip.ParseAddr(host); err == nil {
//line /snap/go/10455/src/net/lookup.go:187
		_go_fuzz_dep_.CoverTab[528941]++
//line /snap/go/10455/src/net/lookup.go:187
		_go_fuzz_dep_.CoverTab[6961]++
							return []string{host}, nil
//line /snap/go/10455/src/net/lookup.go:188
		// _ = "end of CoverTab[6961]"
	} else {
//line /snap/go/10455/src/net/lookup.go:189
		_go_fuzz_dep_.CoverTab[528942]++
//line /snap/go/10455/src/net/lookup.go:189
		_go_fuzz_dep_.CoverTab[6962]++
//line /snap/go/10455/src/net/lookup.go:189
		// _ = "end of CoverTab[6962]"
//line /snap/go/10455/src/net/lookup.go:189
	}
//line /snap/go/10455/src/net/lookup.go:189
	// _ = "end of CoverTab[6957]"
//line /snap/go/10455/src/net/lookup.go:189
	_go_fuzz_dep_.CoverTab[6958]++
						return r.lookupHost(ctx, host)
//line /snap/go/10455/src/net/lookup.go:190
	// _ = "end of CoverTab[6958]"
}

// LookupIP looks up host using the local resolver.
//line /snap/go/10455/src/net/lookup.go:193
// It returns a slice of that host's IPv4 and IPv6 addresses.
//line /snap/go/10455/src/net/lookup.go:195
func LookupIP(host string) ([]IP, error) {
//line /snap/go/10455/src/net/lookup.go:195
	_go_fuzz_dep_.CoverTab[6963]++
						addrs, err := DefaultResolver.LookupIPAddr(context.Background(), host)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:197
		_go_fuzz_dep_.CoverTab[528943]++
//line /snap/go/10455/src/net/lookup.go:197
		_go_fuzz_dep_.CoverTab[6966]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:198
		// _ = "end of CoverTab[6966]"
	} else {
//line /snap/go/10455/src/net/lookup.go:199
		_go_fuzz_dep_.CoverTab[528944]++
//line /snap/go/10455/src/net/lookup.go:199
		_go_fuzz_dep_.CoverTab[6967]++
//line /snap/go/10455/src/net/lookup.go:199
		// _ = "end of CoverTab[6967]"
//line /snap/go/10455/src/net/lookup.go:199
	}
//line /snap/go/10455/src/net/lookup.go:199
	// _ = "end of CoverTab[6963]"
//line /snap/go/10455/src/net/lookup.go:199
	_go_fuzz_dep_.CoverTab[6964]++
						ips := make([]IP, len(addrs))
//line /snap/go/10455/src/net/lookup.go:200
	_go_fuzz_dep_.CoverTab[786701] = 0
						for i, ia := range addrs {
//line /snap/go/10455/src/net/lookup.go:201
		if _go_fuzz_dep_.CoverTab[786701] == 0 {
//line /snap/go/10455/src/net/lookup.go:201
			_go_fuzz_dep_.CoverTab[529090]++
//line /snap/go/10455/src/net/lookup.go:201
		} else {
//line /snap/go/10455/src/net/lookup.go:201
			_go_fuzz_dep_.CoverTab[529091]++
//line /snap/go/10455/src/net/lookup.go:201
		}
//line /snap/go/10455/src/net/lookup.go:201
		_go_fuzz_dep_.CoverTab[786701] = 1
//line /snap/go/10455/src/net/lookup.go:201
		_go_fuzz_dep_.CoverTab[6968]++
							ips[i] = ia.IP
//line /snap/go/10455/src/net/lookup.go:202
		// _ = "end of CoverTab[6968]"
	}
//line /snap/go/10455/src/net/lookup.go:203
	if _go_fuzz_dep_.CoverTab[786701] == 0 {
//line /snap/go/10455/src/net/lookup.go:203
		_go_fuzz_dep_.CoverTab[529092]++
//line /snap/go/10455/src/net/lookup.go:203
	} else {
//line /snap/go/10455/src/net/lookup.go:203
		_go_fuzz_dep_.CoverTab[529093]++
//line /snap/go/10455/src/net/lookup.go:203
	}
//line /snap/go/10455/src/net/lookup.go:203
	// _ = "end of CoverTab[6964]"
//line /snap/go/10455/src/net/lookup.go:203
	_go_fuzz_dep_.CoverTab[6965]++
						return ips, nil
//line /snap/go/10455/src/net/lookup.go:204
	// _ = "end of CoverTab[6965]"
}

// LookupIPAddr looks up host using the local resolver.
//line /snap/go/10455/src/net/lookup.go:207
// It returns a slice of that host's IPv4 and IPv6 addresses.
//line /snap/go/10455/src/net/lookup.go:209
func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error) {
//line /snap/go/10455/src/net/lookup.go:209
	_go_fuzz_dep_.CoverTab[6969]++
						return r.lookupIPAddr(ctx, "ip", host)
//line /snap/go/10455/src/net/lookup.go:210
	// _ = "end of CoverTab[6969]"
}

// LookupIP looks up host for the given network using the local resolver.
//line /snap/go/10455/src/net/lookup.go:213
// It returns a slice of that host's IP addresses of the type specified by
//line /snap/go/10455/src/net/lookup.go:213
// network.
//line /snap/go/10455/src/net/lookup.go:213
// network must be one of "ip", "ip4" or "ip6".
//line /snap/go/10455/src/net/lookup.go:217
func (r *Resolver) LookupIP(ctx context.Context, network, host string) ([]IP, error) {
//line /snap/go/10455/src/net/lookup.go:217
	_go_fuzz_dep_.CoverTab[6970]++
						afnet, _, err := parseNetwork(ctx, network, false)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:219
		_go_fuzz_dep_.CoverTab[528945]++
//line /snap/go/10455/src/net/lookup.go:219
		_go_fuzz_dep_.CoverTab[6976]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:220
		// _ = "end of CoverTab[6976]"
	} else {
//line /snap/go/10455/src/net/lookup.go:221
		_go_fuzz_dep_.CoverTab[528946]++
//line /snap/go/10455/src/net/lookup.go:221
		_go_fuzz_dep_.CoverTab[6977]++
//line /snap/go/10455/src/net/lookup.go:221
		// _ = "end of CoverTab[6977]"
//line /snap/go/10455/src/net/lookup.go:221
	}
//line /snap/go/10455/src/net/lookup.go:221
	// _ = "end of CoverTab[6970]"
//line /snap/go/10455/src/net/lookup.go:221
	_go_fuzz_dep_.CoverTab[6971]++
						switch afnet {
	case "ip", "ip4", "ip6":
//line /snap/go/10455/src/net/lookup.go:223
		_go_fuzz_dep_.CoverTab[528947]++
//line /snap/go/10455/src/net/lookup.go:223
		_go_fuzz_dep_.CoverTab[6978]++
//line /snap/go/10455/src/net/lookup.go:223
		// _ = "end of CoverTab[6978]"
	default:
//line /snap/go/10455/src/net/lookup.go:224
		_go_fuzz_dep_.CoverTab[528948]++
//line /snap/go/10455/src/net/lookup.go:224
		_go_fuzz_dep_.CoverTab[6979]++
							return nil, UnknownNetworkError(network)
//line /snap/go/10455/src/net/lookup.go:225
		// _ = "end of CoverTab[6979]"
	}
//line /snap/go/10455/src/net/lookup.go:226
	// _ = "end of CoverTab[6971]"
//line /snap/go/10455/src/net/lookup.go:226
	_go_fuzz_dep_.CoverTab[6972]++

						if host == "" {
//line /snap/go/10455/src/net/lookup.go:228
		_go_fuzz_dep_.CoverTab[528949]++
//line /snap/go/10455/src/net/lookup.go:228
		_go_fuzz_dep_.CoverTab[6980]++
							return nil, &DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
//line /snap/go/10455/src/net/lookup.go:229
		// _ = "end of CoverTab[6980]"
	} else {
//line /snap/go/10455/src/net/lookup.go:230
		_go_fuzz_dep_.CoverTab[528950]++
//line /snap/go/10455/src/net/lookup.go:230
		_go_fuzz_dep_.CoverTab[6981]++
//line /snap/go/10455/src/net/lookup.go:230
		// _ = "end of CoverTab[6981]"
//line /snap/go/10455/src/net/lookup.go:230
	}
//line /snap/go/10455/src/net/lookup.go:230
	// _ = "end of CoverTab[6972]"
//line /snap/go/10455/src/net/lookup.go:230
	_go_fuzz_dep_.CoverTab[6973]++
						addrs, err := r.internetAddrList(ctx, afnet, host)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:232
		_go_fuzz_dep_.CoverTab[528951]++
//line /snap/go/10455/src/net/lookup.go:232
		_go_fuzz_dep_.CoverTab[6982]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:233
		// _ = "end of CoverTab[6982]"
	} else {
//line /snap/go/10455/src/net/lookup.go:234
		_go_fuzz_dep_.CoverTab[528952]++
//line /snap/go/10455/src/net/lookup.go:234
		_go_fuzz_dep_.CoverTab[6983]++
//line /snap/go/10455/src/net/lookup.go:234
		// _ = "end of CoverTab[6983]"
//line /snap/go/10455/src/net/lookup.go:234
	}
//line /snap/go/10455/src/net/lookup.go:234
	// _ = "end of CoverTab[6973]"
//line /snap/go/10455/src/net/lookup.go:234
	_go_fuzz_dep_.CoverTab[6974]++

						ips := make([]IP, 0, len(addrs))
//line /snap/go/10455/src/net/lookup.go:236
	_go_fuzz_dep_.CoverTab[786702] = 0
						for _, addr := range addrs {
//line /snap/go/10455/src/net/lookup.go:237
		if _go_fuzz_dep_.CoverTab[786702] == 0 {
//line /snap/go/10455/src/net/lookup.go:237
			_go_fuzz_dep_.CoverTab[529094]++
//line /snap/go/10455/src/net/lookup.go:237
		} else {
//line /snap/go/10455/src/net/lookup.go:237
			_go_fuzz_dep_.CoverTab[529095]++
//line /snap/go/10455/src/net/lookup.go:237
		}
//line /snap/go/10455/src/net/lookup.go:237
		_go_fuzz_dep_.CoverTab[786702] = 1
//line /snap/go/10455/src/net/lookup.go:237
		_go_fuzz_dep_.CoverTab[6984]++
							ips = append(ips, addr.(*IPAddr).IP)
//line /snap/go/10455/src/net/lookup.go:238
		// _ = "end of CoverTab[6984]"
	}
//line /snap/go/10455/src/net/lookup.go:239
	if _go_fuzz_dep_.CoverTab[786702] == 0 {
//line /snap/go/10455/src/net/lookup.go:239
		_go_fuzz_dep_.CoverTab[529096]++
//line /snap/go/10455/src/net/lookup.go:239
	} else {
//line /snap/go/10455/src/net/lookup.go:239
		_go_fuzz_dep_.CoverTab[529097]++
//line /snap/go/10455/src/net/lookup.go:239
	}
//line /snap/go/10455/src/net/lookup.go:239
	// _ = "end of CoverTab[6974]"
//line /snap/go/10455/src/net/lookup.go:239
	_go_fuzz_dep_.CoverTab[6975]++
						return ips, nil
//line /snap/go/10455/src/net/lookup.go:240
	// _ = "end of CoverTab[6975]"
}

// LookupNetIP looks up host using the local resolver.
//line /snap/go/10455/src/net/lookup.go:243
// It returns a slice of that host's IP addresses of the type specified by
//line /snap/go/10455/src/net/lookup.go:243
// network.
//line /snap/go/10455/src/net/lookup.go:243
// The network must be one of "ip", "ip4" or "ip6".
//line /snap/go/10455/src/net/lookup.go:247
func (r *Resolver) LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error) {
//line /snap/go/10455/src/net/lookup.go:247
	_go_fuzz_dep_.CoverTab[6985]++

//line /snap/go/10455/src/net/lookup.go:252
	ips, err := r.LookupIP(ctx, network, host)
	if err != nil {
//line /snap/go/10455/src/net/lookup.go:253
		_go_fuzz_dep_.CoverTab[528953]++
//line /snap/go/10455/src/net/lookup.go:253
		_go_fuzz_dep_.CoverTab[6988]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:254
		// _ = "end of CoverTab[6988]"
	} else {
//line /snap/go/10455/src/net/lookup.go:255
		_go_fuzz_dep_.CoverTab[528954]++
//line /snap/go/10455/src/net/lookup.go:255
		_go_fuzz_dep_.CoverTab[6989]++
//line /snap/go/10455/src/net/lookup.go:255
		// _ = "end of CoverTab[6989]"
//line /snap/go/10455/src/net/lookup.go:255
	}
//line /snap/go/10455/src/net/lookup.go:255
	// _ = "end of CoverTab[6985]"
//line /snap/go/10455/src/net/lookup.go:255
	_go_fuzz_dep_.CoverTab[6986]++
						ret := make([]netip.Addr, 0, len(ips))
//line /snap/go/10455/src/net/lookup.go:256
	_go_fuzz_dep_.CoverTab[786703] = 0
						for _, ip := range ips {
//line /snap/go/10455/src/net/lookup.go:257
		if _go_fuzz_dep_.CoverTab[786703] == 0 {
//line /snap/go/10455/src/net/lookup.go:257
			_go_fuzz_dep_.CoverTab[529098]++
//line /snap/go/10455/src/net/lookup.go:257
		} else {
//line /snap/go/10455/src/net/lookup.go:257
			_go_fuzz_dep_.CoverTab[529099]++
//line /snap/go/10455/src/net/lookup.go:257
		}
//line /snap/go/10455/src/net/lookup.go:257
		_go_fuzz_dep_.CoverTab[786703] = 1
//line /snap/go/10455/src/net/lookup.go:257
		_go_fuzz_dep_.CoverTab[6990]++
							if a, ok := netip.AddrFromSlice(ip); ok {
//line /snap/go/10455/src/net/lookup.go:258
			_go_fuzz_dep_.CoverTab[528955]++
//line /snap/go/10455/src/net/lookup.go:258
			_go_fuzz_dep_.CoverTab[6991]++
								ret = append(ret, a)
//line /snap/go/10455/src/net/lookup.go:259
			// _ = "end of CoverTab[6991]"
		} else {
//line /snap/go/10455/src/net/lookup.go:260
			_go_fuzz_dep_.CoverTab[528956]++
//line /snap/go/10455/src/net/lookup.go:260
			_go_fuzz_dep_.CoverTab[6992]++
//line /snap/go/10455/src/net/lookup.go:260
			// _ = "end of CoverTab[6992]"
//line /snap/go/10455/src/net/lookup.go:260
		}
//line /snap/go/10455/src/net/lookup.go:260
		// _ = "end of CoverTab[6990]"
	}
//line /snap/go/10455/src/net/lookup.go:261
	if _go_fuzz_dep_.CoverTab[786703] == 0 {
//line /snap/go/10455/src/net/lookup.go:261
		_go_fuzz_dep_.CoverTab[529100]++
//line /snap/go/10455/src/net/lookup.go:261
	} else {
//line /snap/go/10455/src/net/lookup.go:261
		_go_fuzz_dep_.CoverTab[529101]++
//line /snap/go/10455/src/net/lookup.go:261
	}
//line /snap/go/10455/src/net/lookup.go:261
	// _ = "end of CoverTab[6986]"
//line /snap/go/10455/src/net/lookup.go:261
	_go_fuzz_dep_.CoverTab[6987]++
						return ret, nil
//line /snap/go/10455/src/net/lookup.go:262
	// _ = "end of CoverTab[6987]"
}

// onlyValuesCtx is a context that uses an underlying context
//line /snap/go/10455/src/net/lookup.go:265
// for value lookup if the underlying context hasn't yet expired.
//line /snap/go/10455/src/net/lookup.go:267
type onlyValuesCtx struct {
	context.Context
	lookupValues	context.Context
}

var _ context.Context = (*onlyValuesCtx)(nil)

// Value performs a lookup if the original context hasn't expired.
func (ovc *onlyValuesCtx) Value(key any) any {
//line /snap/go/10455/src/net/lookup.go:275
	_go_fuzz_dep_.CoverTab[6993]++
						select {
	case <-ovc.lookupValues.Done():
//line /snap/go/10455/src/net/lookup.go:277
		_go_fuzz_dep_.CoverTab[6994]++
							return nil
//line /snap/go/10455/src/net/lookup.go:278
		// _ = "end of CoverTab[6994]"
	default:
//line /snap/go/10455/src/net/lookup.go:279
		_go_fuzz_dep_.CoverTab[6995]++
							return ovc.lookupValues.Value(key)
//line /snap/go/10455/src/net/lookup.go:280
		// _ = "end of CoverTab[6995]"
	}
//line /snap/go/10455/src/net/lookup.go:281
	// _ = "end of CoverTab[6993]"
}

// withUnexpiredValuesPreserved returns a context.Context that only uses lookupCtx
//line /snap/go/10455/src/net/lookup.go:284
// for its values, otherwise it is never canceled and has no deadline.
//line /snap/go/10455/src/net/lookup.go:284
// If the lookup context expires, any looked up values will return nil.
//line /snap/go/10455/src/net/lookup.go:284
// See Issue 28600.
//line /snap/go/10455/src/net/lookup.go:288
func withUnexpiredValuesPreserved(lookupCtx context.Context) context.Context {
//line /snap/go/10455/src/net/lookup.go:288
	_go_fuzz_dep_.CoverTab[6996]++
						return &onlyValuesCtx{Context: context.Background(), lookupValues: lookupCtx}
//line /snap/go/10455/src/net/lookup.go:289
	// _ = "end of CoverTab[6996]"
}

// lookupIPAddr looks up host using the local resolver and particular network.
//line /snap/go/10455/src/net/lookup.go:292
// It returns a slice of that host's IPv4 and IPv6 addresses.
//line /snap/go/10455/src/net/lookup.go:294
func (r *Resolver) lookupIPAddr(ctx context.Context, network, host string) ([]IPAddr, error) {
//line /snap/go/10455/src/net/lookup.go:294
	_go_fuzz_dep_.CoverTab[6997]++

						if host == "" {
//line /snap/go/10455/src/net/lookup.go:296
		_go_fuzz_dep_.CoverTab[528957]++
//line /snap/go/10455/src/net/lookup.go:296
		_go_fuzz_dep_.CoverTab[7004]++
							return nil, &DNSError{Err: errNoSuchHost.Error(), Name: host, IsNotFound: true}
//line /snap/go/10455/src/net/lookup.go:297
		// _ = "end of CoverTab[7004]"
	} else {
//line /snap/go/10455/src/net/lookup.go:298
		_go_fuzz_dep_.CoverTab[528958]++
//line /snap/go/10455/src/net/lookup.go:298
		_go_fuzz_dep_.CoverTab[7005]++
//line /snap/go/10455/src/net/lookup.go:298
		// _ = "end of CoverTab[7005]"
//line /snap/go/10455/src/net/lookup.go:298
	}
//line /snap/go/10455/src/net/lookup.go:298
	// _ = "end of CoverTab[6997]"
//line /snap/go/10455/src/net/lookup.go:298
	_go_fuzz_dep_.CoverTab[6998]++
						if ip, err := netip.ParseAddr(host); err == nil {
//line /snap/go/10455/src/net/lookup.go:299
		_go_fuzz_dep_.CoverTab[528959]++
//line /snap/go/10455/src/net/lookup.go:299
		_go_fuzz_dep_.CoverTab[7006]++
							return []IPAddr{{IP: IP(ip.AsSlice()).To16(), Zone: ip.Zone()}}, nil
//line /snap/go/10455/src/net/lookup.go:300
		// _ = "end of CoverTab[7006]"
	} else {
//line /snap/go/10455/src/net/lookup.go:301
		_go_fuzz_dep_.CoverTab[528960]++
//line /snap/go/10455/src/net/lookup.go:301
		_go_fuzz_dep_.CoverTab[7007]++
//line /snap/go/10455/src/net/lookup.go:301
		// _ = "end of CoverTab[7007]"
//line /snap/go/10455/src/net/lookup.go:301
	}
//line /snap/go/10455/src/net/lookup.go:301
	// _ = "end of CoverTab[6998]"
//line /snap/go/10455/src/net/lookup.go:301
	_go_fuzz_dep_.CoverTab[6999]++
						trace, _ := ctx.Value(nettrace.TraceKey{}).(*nettrace.Trace)
						if trace != nil && func() bool {
//line /snap/go/10455/src/net/lookup.go:303
		_go_fuzz_dep_.CoverTab[7008]++
//line /snap/go/10455/src/net/lookup.go:303
		return trace.DNSStart != nil
//line /snap/go/10455/src/net/lookup.go:303
		// _ = "end of CoverTab[7008]"
//line /snap/go/10455/src/net/lookup.go:303
	}() {
//line /snap/go/10455/src/net/lookup.go:303
		_go_fuzz_dep_.CoverTab[528961]++
//line /snap/go/10455/src/net/lookup.go:303
		_go_fuzz_dep_.CoverTab[7009]++
							trace.DNSStart(host)
//line /snap/go/10455/src/net/lookup.go:304
		// _ = "end of CoverTab[7009]"
	} else {
//line /snap/go/10455/src/net/lookup.go:305
		_go_fuzz_dep_.CoverTab[528962]++
//line /snap/go/10455/src/net/lookup.go:305
		_go_fuzz_dep_.CoverTab[7010]++
//line /snap/go/10455/src/net/lookup.go:305
		// _ = "end of CoverTab[7010]"
//line /snap/go/10455/src/net/lookup.go:305
	}
//line /snap/go/10455/src/net/lookup.go:305
	// _ = "end of CoverTab[6999]"
//line /snap/go/10455/src/net/lookup.go:305
	_go_fuzz_dep_.CoverTab[7000]++

//line /snap/go/10455/src/net/lookup.go:309
	resolverFunc := r.lookupIP
	if alt, _ := ctx.Value(nettrace.LookupIPAltResolverKey{}).(func(context.Context, string, string) ([]IPAddr, error)); alt != nil {
//line /snap/go/10455/src/net/lookup.go:310
		_go_fuzz_dep_.CoverTab[528963]++
//line /snap/go/10455/src/net/lookup.go:310
		_go_fuzz_dep_.CoverTab[7011]++
							resolverFunc = alt
//line /snap/go/10455/src/net/lookup.go:311
		// _ = "end of CoverTab[7011]"
	} else {
//line /snap/go/10455/src/net/lookup.go:312
		_go_fuzz_dep_.CoverTab[528964]++
//line /snap/go/10455/src/net/lookup.go:312
		_go_fuzz_dep_.CoverTab[7012]++
//line /snap/go/10455/src/net/lookup.go:312
		// _ = "end of CoverTab[7012]"
//line /snap/go/10455/src/net/lookup.go:312
	}
//line /snap/go/10455/src/net/lookup.go:312
	// _ = "end of CoverTab[7000]"
//line /snap/go/10455/src/net/lookup.go:312
	_go_fuzz_dep_.CoverTab[7001]++

//line /snap/go/10455/src/net/lookup.go:319
	lookupGroupCtx, lookupGroupCancel := context.WithCancel(withUnexpiredValuesPreserved(ctx))

	lookupKey := network + "\000" + host
	dnsWaitGroup.Add(1)
	ch := r.getLookupGroup().DoChan(lookupKey, func() (any, error) {
//line /snap/go/10455/src/net/lookup.go:323
		_go_fuzz_dep_.CoverTab[7013]++
							return testHookLookupIP(lookupGroupCtx, resolverFunc, network, host)
//line /snap/go/10455/src/net/lookup.go:324
		// _ = "end of CoverTab[7013]"
	})
//line /snap/go/10455/src/net/lookup.go:325
	// _ = "end of CoverTab[7001]"
//line /snap/go/10455/src/net/lookup.go:325
	_go_fuzz_dep_.CoverTab[7002]++

						dnsWaitGroupDone := func(ch <-chan singleflight.Result, cancelFn context.CancelFunc) {
//line /snap/go/10455/src/net/lookup.go:327
		_go_fuzz_dep_.CoverTab[7014]++
							<-ch
							dnsWaitGroup.Done()
							cancelFn()
//line /snap/go/10455/src/net/lookup.go:330
		// _ = "end of CoverTab[7014]"
	}
//line /snap/go/10455/src/net/lookup.go:331
	// _ = "end of CoverTab[7002]"
//line /snap/go/10455/src/net/lookup.go:331
	_go_fuzz_dep_.CoverTab[7003]++
						select {
	case <-ctx.Done():
//line /snap/go/10455/src/net/lookup.go:333
		_go_fuzz_dep_.CoverTab[7015]++

//line /snap/go/10455/src/net/lookup.go:341
		if r.getLookupGroup().ForgetUnshared(lookupKey) {
//line /snap/go/10455/src/net/lookup.go:341
			_go_fuzz_dep_.CoverTab[528965]++
//line /snap/go/10455/src/net/lookup.go:341
			_go_fuzz_dep_.CoverTab[7021]++
								lookupGroupCancel()
//line /snap/go/10455/src/net/lookup.go:342
			_curRoutineNum8_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/net/lookup.go:342
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum8_)
								go func() {
//line /snap/go/10455/src/net/lookup.go:343
				_go_fuzz_dep_.CoverTab[7022]++
//line /snap/go/10455/src/net/lookup.go:343
				defer func() {
//line /snap/go/10455/src/net/lookup.go:343
					_go_fuzz_dep_.CoverTab[7023]++
//line /snap/go/10455/src/net/lookup.go:343
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum8_)
//line /snap/go/10455/src/net/lookup.go:343
					// _ = "end of CoverTab[7023]"
//line /snap/go/10455/src/net/lookup.go:343
				}()
//line /snap/go/10455/src/net/lookup.go:343
				dnsWaitGroupDone(ch, func() { _go_fuzz_dep_.CoverTab[7024]++; // _ = "end of CoverTab[7024]" })
//line /snap/go/10455/src/net/lookup.go:343
				// _ = "end of CoverTab[7022]"
//line /snap/go/10455/src/net/lookup.go:343
			}()
//line /snap/go/10455/src/net/lookup.go:343
			// _ = "end of CoverTab[7021]"
		} else {
//line /snap/go/10455/src/net/lookup.go:344
			_go_fuzz_dep_.CoverTab[528966]++
//line /snap/go/10455/src/net/lookup.go:344
			_go_fuzz_dep_.CoverTab[7025]++
//line /snap/go/10455/src/net/lookup.go:344
			_curRoutineNum9_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/net/lookup.go:344
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum9_)
								go func() {
//line /snap/go/10455/src/net/lookup.go:345
				_go_fuzz_dep_.CoverTab[7026]++
//line /snap/go/10455/src/net/lookup.go:345
				defer func() {
//line /snap/go/10455/src/net/lookup.go:345
					_go_fuzz_dep_.CoverTab[7027]++
//line /snap/go/10455/src/net/lookup.go:345
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum9_)
//line /snap/go/10455/src/net/lookup.go:345
					// _ = "end of CoverTab[7027]"
//line /snap/go/10455/src/net/lookup.go:345
				}()
//line /snap/go/10455/src/net/lookup.go:345
				dnsWaitGroupDone(ch, lookupGroupCancel)
//line /snap/go/10455/src/net/lookup.go:345
				// _ = "end of CoverTab[7026]"
//line /snap/go/10455/src/net/lookup.go:345
			}()
//line /snap/go/10455/src/net/lookup.go:345
			// _ = "end of CoverTab[7025]"
		}
//line /snap/go/10455/src/net/lookup.go:346
		// _ = "end of CoverTab[7015]"
//line /snap/go/10455/src/net/lookup.go:346
		_go_fuzz_dep_.CoverTab[7016]++
							ctxErr := ctx.Err()
							err := &DNSError{
			Err:		mapErr(ctxErr).Error(),
			Name:		host,
			IsTimeout:	ctxErr == context.DeadlineExceeded,
		}
		if trace != nil && func() bool {
//line /snap/go/10455/src/net/lookup.go:353
			_go_fuzz_dep_.CoverTab[7028]++
//line /snap/go/10455/src/net/lookup.go:353
			return trace.DNSDone != nil
//line /snap/go/10455/src/net/lookup.go:353
			// _ = "end of CoverTab[7028]"
//line /snap/go/10455/src/net/lookup.go:353
		}() {
//line /snap/go/10455/src/net/lookup.go:353
			_go_fuzz_dep_.CoverTab[528967]++
//line /snap/go/10455/src/net/lookup.go:353
			_go_fuzz_dep_.CoverTab[7029]++
								trace.DNSDone(nil, false, err)
//line /snap/go/10455/src/net/lookup.go:354
			// _ = "end of CoverTab[7029]"
		} else {
//line /snap/go/10455/src/net/lookup.go:355
			_go_fuzz_dep_.CoverTab[528968]++
//line /snap/go/10455/src/net/lookup.go:355
			_go_fuzz_dep_.CoverTab[7030]++
//line /snap/go/10455/src/net/lookup.go:355
			// _ = "end of CoverTab[7030]"
//line /snap/go/10455/src/net/lookup.go:355
		}
//line /snap/go/10455/src/net/lookup.go:355
		// _ = "end of CoverTab[7016]"
//line /snap/go/10455/src/net/lookup.go:355
		_go_fuzz_dep_.CoverTab[7017]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:356
		// _ = "end of CoverTab[7017]"
	case r := <-ch:
//line /snap/go/10455/src/net/lookup.go:357
		_go_fuzz_dep_.CoverTab[7018]++
							dnsWaitGroup.Done()
							lookupGroupCancel()
							err := r.Err
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:361
			_go_fuzz_dep_.CoverTab[528969]++
//line /snap/go/10455/src/net/lookup.go:361
			_go_fuzz_dep_.CoverTab[7031]++
								if _, ok := err.(*DNSError); !ok {
//line /snap/go/10455/src/net/lookup.go:362
				_go_fuzz_dep_.CoverTab[528971]++
//line /snap/go/10455/src/net/lookup.go:362
				_go_fuzz_dep_.CoverTab[7032]++
									isTimeout := false
									if err == context.DeadlineExceeded {
//line /snap/go/10455/src/net/lookup.go:364
					_go_fuzz_dep_.CoverTab[528973]++
//line /snap/go/10455/src/net/lookup.go:364
					_go_fuzz_dep_.CoverTab[7034]++
										isTimeout = true
//line /snap/go/10455/src/net/lookup.go:365
					// _ = "end of CoverTab[7034]"
				} else {
//line /snap/go/10455/src/net/lookup.go:366
					_go_fuzz_dep_.CoverTab[528974]++
//line /snap/go/10455/src/net/lookup.go:366
					_go_fuzz_dep_.CoverTab[7035]++
//line /snap/go/10455/src/net/lookup.go:366
					if terr, ok := err.(timeout); ok {
//line /snap/go/10455/src/net/lookup.go:366
						_go_fuzz_dep_.CoverTab[528975]++
//line /snap/go/10455/src/net/lookup.go:366
						_go_fuzz_dep_.CoverTab[7036]++
											isTimeout = terr.Timeout()
//line /snap/go/10455/src/net/lookup.go:367
						// _ = "end of CoverTab[7036]"
					} else {
//line /snap/go/10455/src/net/lookup.go:368
						_go_fuzz_dep_.CoverTab[528976]++
//line /snap/go/10455/src/net/lookup.go:368
						_go_fuzz_dep_.CoverTab[7037]++
//line /snap/go/10455/src/net/lookup.go:368
						// _ = "end of CoverTab[7037]"
//line /snap/go/10455/src/net/lookup.go:368
					}
//line /snap/go/10455/src/net/lookup.go:368
					// _ = "end of CoverTab[7035]"
//line /snap/go/10455/src/net/lookup.go:368
				}
//line /snap/go/10455/src/net/lookup.go:368
				// _ = "end of CoverTab[7032]"
//line /snap/go/10455/src/net/lookup.go:368
				_go_fuzz_dep_.CoverTab[7033]++
									err = &DNSError{
					Err:		err.Error(),
					Name:		host,
					IsTimeout:	isTimeout,
				}
//line /snap/go/10455/src/net/lookup.go:373
				// _ = "end of CoverTab[7033]"
			} else {
//line /snap/go/10455/src/net/lookup.go:374
				_go_fuzz_dep_.CoverTab[528972]++
//line /snap/go/10455/src/net/lookup.go:374
				_go_fuzz_dep_.CoverTab[7038]++
//line /snap/go/10455/src/net/lookup.go:374
				// _ = "end of CoverTab[7038]"
//line /snap/go/10455/src/net/lookup.go:374
			}
//line /snap/go/10455/src/net/lookup.go:374
			// _ = "end of CoverTab[7031]"
		} else {
//line /snap/go/10455/src/net/lookup.go:375
			_go_fuzz_dep_.CoverTab[528970]++
//line /snap/go/10455/src/net/lookup.go:375
			_go_fuzz_dep_.CoverTab[7039]++
//line /snap/go/10455/src/net/lookup.go:375
			// _ = "end of CoverTab[7039]"
//line /snap/go/10455/src/net/lookup.go:375
		}
//line /snap/go/10455/src/net/lookup.go:375
		// _ = "end of CoverTab[7018]"
//line /snap/go/10455/src/net/lookup.go:375
		_go_fuzz_dep_.CoverTab[7019]++
							if trace != nil && func() bool {
//line /snap/go/10455/src/net/lookup.go:376
			_go_fuzz_dep_.CoverTab[7040]++
//line /snap/go/10455/src/net/lookup.go:376
			return trace.DNSDone != nil
//line /snap/go/10455/src/net/lookup.go:376
			// _ = "end of CoverTab[7040]"
//line /snap/go/10455/src/net/lookup.go:376
		}() {
//line /snap/go/10455/src/net/lookup.go:376
			_go_fuzz_dep_.CoverTab[528977]++
//line /snap/go/10455/src/net/lookup.go:376
			_go_fuzz_dep_.CoverTab[7041]++
								addrs, _ := r.Val.([]IPAddr)
								trace.DNSDone(ipAddrsEface(addrs), r.Shared, err)
//line /snap/go/10455/src/net/lookup.go:378
			// _ = "end of CoverTab[7041]"
		} else {
//line /snap/go/10455/src/net/lookup.go:379
			_go_fuzz_dep_.CoverTab[528978]++
//line /snap/go/10455/src/net/lookup.go:379
			_go_fuzz_dep_.CoverTab[7042]++
//line /snap/go/10455/src/net/lookup.go:379
			// _ = "end of CoverTab[7042]"
//line /snap/go/10455/src/net/lookup.go:379
		}
//line /snap/go/10455/src/net/lookup.go:379
		// _ = "end of CoverTab[7019]"
//line /snap/go/10455/src/net/lookup.go:379
		_go_fuzz_dep_.CoverTab[7020]++
							return lookupIPReturn(r.Val, err, r.Shared)
//line /snap/go/10455/src/net/lookup.go:380
		// _ = "end of CoverTab[7020]"
	}
//line /snap/go/10455/src/net/lookup.go:381
	// _ = "end of CoverTab[7003]"
}

// lookupIPReturn turns the return values from singleflight.Do into
//line /snap/go/10455/src/net/lookup.go:384
// the return values from LookupIP.
//line /snap/go/10455/src/net/lookup.go:386
func lookupIPReturn(addrsi any, err error, shared bool) ([]IPAddr, error) {
//line /snap/go/10455/src/net/lookup.go:386
	_go_fuzz_dep_.CoverTab[7043]++
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:387
		_go_fuzz_dep_.CoverTab[528979]++
//line /snap/go/10455/src/net/lookup.go:387
		_go_fuzz_dep_.CoverTab[7046]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:388
		// _ = "end of CoverTab[7046]"
	} else {
//line /snap/go/10455/src/net/lookup.go:389
		_go_fuzz_dep_.CoverTab[528980]++
//line /snap/go/10455/src/net/lookup.go:389
		_go_fuzz_dep_.CoverTab[7047]++
//line /snap/go/10455/src/net/lookup.go:389
		// _ = "end of CoverTab[7047]"
//line /snap/go/10455/src/net/lookup.go:389
	}
//line /snap/go/10455/src/net/lookup.go:389
	// _ = "end of CoverTab[7043]"
//line /snap/go/10455/src/net/lookup.go:389
	_go_fuzz_dep_.CoverTab[7044]++
						addrs := addrsi.([]IPAddr)
						if shared {
//line /snap/go/10455/src/net/lookup.go:391
		_go_fuzz_dep_.CoverTab[528981]++
//line /snap/go/10455/src/net/lookup.go:391
		_go_fuzz_dep_.CoverTab[7048]++
							clone := make([]IPAddr, len(addrs))
							copy(clone, addrs)
							addrs = clone
//line /snap/go/10455/src/net/lookup.go:394
		// _ = "end of CoverTab[7048]"
	} else {
//line /snap/go/10455/src/net/lookup.go:395
		_go_fuzz_dep_.CoverTab[528982]++
//line /snap/go/10455/src/net/lookup.go:395
		_go_fuzz_dep_.CoverTab[7049]++
//line /snap/go/10455/src/net/lookup.go:395
		// _ = "end of CoverTab[7049]"
//line /snap/go/10455/src/net/lookup.go:395
	}
//line /snap/go/10455/src/net/lookup.go:395
	// _ = "end of CoverTab[7044]"
//line /snap/go/10455/src/net/lookup.go:395
	_go_fuzz_dep_.CoverTab[7045]++
						return addrs, nil
//line /snap/go/10455/src/net/lookup.go:396
	// _ = "end of CoverTab[7045]"
}

// ipAddrsEface returns an empty interface slice of addrs.
func ipAddrsEface(addrs []IPAddr) []any {
//line /snap/go/10455/src/net/lookup.go:400
	_go_fuzz_dep_.CoverTab[7050]++
						s := make([]any, len(addrs))
//line /snap/go/10455/src/net/lookup.go:401
	_go_fuzz_dep_.CoverTab[786704] = 0
						for i, v := range addrs {
//line /snap/go/10455/src/net/lookup.go:402
		if _go_fuzz_dep_.CoverTab[786704] == 0 {
//line /snap/go/10455/src/net/lookup.go:402
			_go_fuzz_dep_.CoverTab[529102]++
//line /snap/go/10455/src/net/lookup.go:402
		} else {
//line /snap/go/10455/src/net/lookup.go:402
			_go_fuzz_dep_.CoverTab[529103]++
//line /snap/go/10455/src/net/lookup.go:402
		}
//line /snap/go/10455/src/net/lookup.go:402
		_go_fuzz_dep_.CoverTab[786704] = 1
//line /snap/go/10455/src/net/lookup.go:402
		_go_fuzz_dep_.CoverTab[7052]++
							s[i] = v
//line /snap/go/10455/src/net/lookup.go:403
		// _ = "end of CoverTab[7052]"
	}
//line /snap/go/10455/src/net/lookup.go:404
	if _go_fuzz_dep_.CoverTab[786704] == 0 {
//line /snap/go/10455/src/net/lookup.go:404
		_go_fuzz_dep_.CoverTab[529104]++
//line /snap/go/10455/src/net/lookup.go:404
	} else {
//line /snap/go/10455/src/net/lookup.go:404
		_go_fuzz_dep_.CoverTab[529105]++
//line /snap/go/10455/src/net/lookup.go:404
	}
//line /snap/go/10455/src/net/lookup.go:404
	// _ = "end of CoverTab[7050]"
//line /snap/go/10455/src/net/lookup.go:404
	_go_fuzz_dep_.CoverTab[7051]++
						return s
//line /snap/go/10455/src/net/lookup.go:405
	// _ = "end of CoverTab[7051]"
}

// LookupPort looks up the port for the given network and service.
//line /snap/go/10455/src/net/lookup.go:408
//
//line /snap/go/10455/src/net/lookup.go:408
// LookupPort uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/lookup.go:408
// Resolver.LookupPort.
//line /snap/go/10455/src/net/lookup.go:412
func LookupPort(network, service string) (port int, err error) {
//line /snap/go/10455/src/net/lookup.go:412
	_go_fuzz_dep_.CoverTab[7053]++
						return DefaultResolver.LookupPort(context.Background(), network, service)
//line /snap/go/10455/src/net/lookup.go:413
	// _ = "end of CoverTab[7053]"
}

// LookupPort looks up the port for the given network and service.
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error) {
//line /snap/go/10455/src/net/lookup.go:417
	_go_fuzz_dep_.CoverTab[7054]++
						port, needsLookup := parsePort(service)
						if needsLookup {
//line /snap/go/10455/src/net/lookup.go:419
		_go_fuzz_dep_.CoverTab[528983]++
//line /snap/go/10455/src/net/lookup.go:419
		_go_fuzz_dep_.CoverTab[7057]++
							switch network {
		case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
//line /snap/go/10455/src/net/lookup.go:421
			_go_fuzz_dep_.CoverTab[528985]++
//line /snap/go/10455/src/net/lookup.go:421
			_go_fuzz_dep_.CoverTab[7059]++
//line /snap/go/10455/src/net/lookup.go:421
			// _ = "end of CoverTab[7059]"
		case "":
//line /snap/go/10455/src/net/lookup.go:422
			_go_fuzz_dep_.CoverTab[528986]++
//line /snap/go/10455/src/net/lookup.go:422
			_go_fuzz_dep_.CoverTab[7060]++
								network = "ip"
//line /snap/go/10455/src/net/lookup.go:423
			// _ = "end of CoverTab[7060]"
		default:
//line /snap/go/10455/src/net/lookup.go:424
			_go_fuzz_dep_.CoverTab[528987]++
//line /snap/go/10455/src/net/lookup.go:424
			_go_fuzz_dep_.CoverTab[7061]++
								return 0, &AddrError{Err: "unknown network", Addr: network}
//line /snap/go/10455/src/net/lookup.go:425
			// _ = "end of CoverTab[7061]"
		}
//line /snap/go/10455/src/net/lookup.go:426
		// _ = "end of CoverTab[7057]"
//line /snap/go/10455/src/net/lookup.go:426
		_go_fuzz_dep_.CoverTab[7058]++
							port, err = r.lookupPort(ctx, network, service)
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:428
			_go_fuzz_dep_.CoverTab[528988]++
//line /snap/go/10455/src/net/lookup.go:428
			_go_fuzz_dep_.CoverTab[7062]++
								return 0, err
//line /snap/go/10455/src/net/lookup.go:429
			// _ = "end of CoverTab[7062]"
		} else {
//line /snap/go/10455/src/net/lookup.go:430
			_go_fuzz_dep_.CoverTab[528989]++
//line /snap/go/10455/src/net/lookup.go:430
			_go_fuzz_dep_.CoverTab[7063]++
//line /snap/go/10455/src/net/lookup.go:430
			// _ = "end of CoverTab[7063]"
//line /snap/go/10455/src/net/lookup.go:430
		}
//line /snap/go/10455/src/net/lookup.go:430
		// _ = "end of CoverTab[7058]"
	} else {
//line /snap/go/10455/src/net/lookup.go:431
		_go_fuzz_dep_.CoverTab[528984]++
//line /snap/go/10455/src/net/lookup.go:431
		_go_fuzz_dep_.CoverTab[7064]++
//line /snap/go/10455/src/net/lookup.go:431
		// _ = "end of CoverTab[7064]"
//line /snap/go/10455/src/net/lookup.go:431
	}
//line /snap/go/10455/src/net/lookup.go:431
	// _ = "end of CoverTab[7054]"
//line /snap/go/10455/src/net/lookup.go:431
	_go_fuzz_dep_.CoverTab[7055]++
						if 0 > port || func() bool {
//line /snap/go/10455/src/net/lookup.go:432
		_go_fuzz_dep_.CoverTab[7065]++
//line /snap/go/10455/src/net/lookup.go:432
		return port > 65535
//line /snap/go/10455/src/net/lookup.go:432
		// _ = "end of CoverTab[7065]"
//line /snap/go/10455/src/net/lookup.go:432
	}() {
//line /snap/go/10455/src/net/lookup.go:432
		_go_fuzz_dep_.CoverTab[528990]++
//line /snap/go/10455/src/net/lookup.go:432
		_go_fuzz_dep_.CoverTab[7066]++
							return 0, &AddrError{Err: "invalid port", Addr: service}
//line /snap/go/10455/src/net/lookup.go:433
		// _ = "end of CoverTab[7066]"
	} else {
//line /snap/go/10455/src/net/lookup.go:434
		_go_fuzz_dep_.CoverTab[528991]++
//line /snap/go/10455/src/net/lookup.go:434
		_go_fuzz_dep_.CoverTab[7067]++
//line /snap/go/10455/src/net/lookup.go:434
		// _ = "end of CoverTab[7067]"
//line /snap/go/10455/src/net/lookup.go:434
	}
//line /snap/go/10455/src/net/lookup.go:434
	// _ = "end of CoverTab[7055]"
//line /snap/go/10455/src/net/lookup.go:434
	_go_fuzz_dep_.CoverTab[7056]++
						return port, nil
//line /snap/go/10455/src/net/lookup.go:435
	// _ = "end of CoverTab[7056]"
}

// LookupCNAME returns the canonical name for the given host.
//line /snap/go/10455/src/net/lookup.go:438
// Callers that do not care about the canonical name can call
//line /snap/go/10455/src/net/lookup.go:438
// LookupHost or LookupIP directly; both take care of resolving
//line /snap/go/10455/src/net/lookup.go:438
// the canonical name as part of the lookup.
//line /snap/go/10455/src/net/lookup.go:438
//
//line /snap/go/10455/src/net/lookup.go:438
// A canonical name is the final name after following zero
//line /snap/go/10455/src/net/lookup.go:438
// or more CNAME records.
//line /snap/go/10455/src/net/lookup.go:438
// LookupCNAME does not return an error if host does not
//line /snap/go/10455/src/net/lookup.go:438
// contain DNS "CNAME" records, as long as host resolves to
//line /snap/go/10455/src/net/lookup.go:438
// address records.
//line /snap/go/10455/src/net/lookup.go:438
//
//line /snap/go/10455/src/net/lookup.go:438
// The returned canonical name is validated to be a properly
//line /snap/go/10455/src/net/lookup.go:438
// formatted presentation-format domain name.
//line /snap/go/10455/src/net/lookup.go:438
//
//line /snap/go/10455/src/net/lookup.go:438
// LookupCNAME uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/lookup.go:438
// Resolver.LookupCNAME.
//line /snap/go/10455/src/net/lookup.go:454
func LookupCNAME(host string) (cname string, err error) {
//line /snap/go/10455/src/net/lookup.go:454
	_go_fuzz_dep_.CoverTab[7068]++
						return DefaultResolver.LookupCNAME(context.Background(), host)
//line /snap/go/10455/src/net/lookup.go:455
	// _ = "end of CoverTab[7068]"
}

// LookupCNAME returns the canonical name for the given host.
//line /snap/go/10455/src/net/lookup.go:458
// Callers that do not care about the canonical name can call
//line /snap/go/10455/src/net/lookup.go:458
// LookupHost or LookupIP directly; both take care of resolving
//line /snap/go/10455/src/net/lookup.go:458
// the canonical name as part of the lookup.
//line /snap/go/10455/src/net/lookup.go:458
//
//line /snap/go/10455/src/net/lookup.go:458
// A canonical name is the final name after following zero
//line /snap/go/10455/src/net/lookup.go:458
// or more CNAME records.
//line /snap/go/10455/src/net/lookup.go:458
// LookupCNAME does not return an error if host does not
//line /snap/go/10455/src/net/lookup.go:458
// contain DNS "CNAME" records, as long as host resolves to
//line /snap/go/10455/src/net/lookup.go:458
// address records.
//line /snap/go/10455/src/net/lookup.go:458
//
//line /snap/go/10455/src/net/lookup.go:458
// The returned canonical name is validated to be a properly
//line /snap/go/10455/src/net/lookup.go:458
// formatted presentation-format domain name.
//line /snap/go/10455/src/net/lookup.go:471
func (r *Resolver) LookupCNAME(ctx context.Context, host string) (string, error) {
//line /snap/go/10455/src/net/lookup.go:471
	_go_fuzz_dep_.CoverTab[7069]++
						cname, err := r.lookupCNAME(ctx, host)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:473
		_go_fuzz_dep_.CoverTab[528992]++
//line /snap/go/10455/src/net/lookup.go:473
		_go_fuzz_dep_.CoverTab[7072]++
							return "", err
//line /snap/go/10455/src/net/lookup.go:474
		// _ = "end of CoverTab[7072]"
	} else {
//line /snap/go/10455/src/net/lookup.go:475
		_go_fuzz_dep_.CoverTab[528993]++
//line /snap/go/10455/src/net/lookup.go:475
		_go_fuzz_dep_.CoverTab[7073]++
//line /snap/go/10455/src/net/lookup.go:475
		// _ = "end of CoverTab[7073]"
//line /snap/go/10455/src/net/lookup.go:475
	}
//line /snap/go/10455/src/net/lookup.go:475
	// _ = "end of CoverTab[7069]"
//line /snap/go/10455/src/net/lookup.go:475
	_go_fuzz_dep_.CoverTab[7070]++
						if !isDomainName(cname) {
//line /snap/go/10455/src/net/lookup.go:476
		_go_fuzz_dep_.CoverTab[528994]++
//line /snap/go/10455/src/net/lookup.go:476
		_go_fuzz_dep_.CoverTab[7074]++
							return "", &DNSError{Err: errMalformedDNSRecordsDetail, Name: host}
//line /snap/go/10455/src/net/lookup.go:477
		// _ = "end of CoverTab[7074]"
	} else {
//line /snap/go/10455/src/net/lookup.go:478
		_go_fuzz_dep_.CoverTab[528995]++
//line /snap/go/10455/src/net/lookup.go:478
		_go_fuzz_dep_.CoverTab[7075]++
//line /snap/go/10455/src/net/lookup.go:478
		// _ = "end of CoverTab[7075]"
//line /snap/go/10455/src/net/lookup.go:478
	}
//line /snap/go/10455/src/net/lookup.go:478
	// _ = "end of CoverTab[7070]"
//line /snap/go/10455/src/net/lookup.go:478
	_go_fuzz_dep_.CoverTab[7071]++
						return cname, nil
//line /snap/go/10455/src/net/lookup.go:479
	// _ = "end of CoverTab[7071]"
}

// LookupSRV tries to resolve an SRV query of the given service,
//line /snap/go/10455/src/net/lookup.go:482
// protocol, and domain name. The proto is "tcp" or "udp".
//line /snap/go/10455/src/net/lookup.go:482
// The returned records are sorted by priority and randomized
//line /snap/go/10455/src/net/lookup.go:482
// by weight within a priority.
//line /snap/go/10455/src/net/lookup.go:482
//
//line /snap/go/10455/src/net/lookup.go:482
// LookupSRV constructs the DNS name to look up following RFC 2782.
//line /snap/go/10455/src/net/lookup.go:482
// That is, it looks up _service._proto.name. To accommodate services
//line /snap/go/10455/src/net/lookup.go:482
// publishing SRV records under non-standard names, if both service
//line /snap/go/10455/src/net/lookup.go:482
// and proto are empty strings, LookupSRV looks up name directly.
//line /snap/go/10455/src/net/lookup.go:482
//
//line /snap/go/10455/src/net/lookup.go:482
// The returned service names are validated to be properly
//line /snap/go/10455/src/net/lookup.go:482
// formatted presentation-format domain names. If the response contains
//line /snap/go/10455/src/net/lookup.go:482
// invalid names, those records are filtered out and an error
//line /snap/go/10455/src/net/lookup.go:482
// will be returned alongside the remaining results, if any.
//line /snap/go/10455/src/net/lookup.go:496
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error) {
//line /snap/go/10455/src/net/lookup.go:496
	_go_fuzz_dep_.CoverTab[7076]++
						return DefaultResolver.LookupSRV(context.Background(), service, proto, name)
//line /snap/go/10455/src/net/lookup.go:497
	// _ = "end of CoverTab[7076]"
}

// LookupSRV tries to resolve an SRV query of the given service,
//line /snap/go/10455/src/net/lookup.go:500
// protocol, and domain name. The proto is "tcp" or "udp".
//line /snap/go/10455/src/net/lookup.go:500
// The returned records are sorted by priority and randomized
//line /snap/go/10455/src/net/lookup.go:500
// by weight within a priority.
//line /snap/go/10455/src/net/lookup.go:500
//
//line /snap/go/10455/src/net/lookup.go:500
// LookupSRV constructs the DNS name to look up following RFC 2782.
//line /snap/go/10455/src/net/lookup.go:500
// That is, it looks up _service._proto.name. To accommodate services
//line /snap/go/10455/src/net/lookup.go:500
// publishing SRV records under non-standard names, if both service
//line /snap/go/10455/src/net/lookup.go:500
// and proto are empty strings, LookupSRV looks up name directly.
//line /snap/go/10455/src/net/lookup.go:500
//
//line /snap/go/10455/src/net/lookup.go:500
// The returned service names are validated to be properly
//line /snap/go/10455/src/net/lookup.go:500
// formatted presentation-format domain names. If the response contains
//line /snap/go/10455/src/net/lookup.go:500
// invalid names, those records are filtered out and an error
//line /snap/go/10455/src/net/lookup.go:500
// will be returned alongside the remaining results, if any.
//line /snap/go/10455/src/net/lookup.go:514
func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error) {
//line /snap/go/10455/src/net/lookup.go:514
	_go_fuzz_dep_.CoverTab[7077]++
						cname, addrs, err := r.lookupSRV(ctx, service, proto, name)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:516
		_go_fuzz_dep_.CoverTab[528996]++
//line /snap/go/10455/src/net/lookup.go:516
		_go_fuzz_dep_.CoverTab[7082]++
							return "", nil, err
//line /snap/go/10455/src/net/lookup.go:517
		// _ = "end of CoverTab[7082]"
	} else {
//line /snap/go/10455/src/net/lookup.go:518
		_go_fuzz_dep_.CoverTab[528997]++
//line /snap/go/10455/src/net/lookup.go:518
		_go_fuzz_dep_.CoverTab[7083]++
//line /snap/go/10455/src/net/lookup.go:518
		// _ = "end of CoverTab[7083]"
//line /snap/go/10455/src/net/lookup.go:518
	}
//line /snap/go/10455/src/net/lookup.go:518
	// _ = "end of CoverTab[7077]"
//line /snap/go/10455/src/net/lookup.go:518
	_go_fuzz_dep_.CoverTab[7078]++
						if cname != "" && func() bool {
//line /snap/go/10455/src/net/lookup.go:519
		_go_fuzz_dep_.CoverTab[7084]++
//line /snap/go/10455/src/net/lookup.go:519
		return !isDomainName(cname)
//line /snap/go/10455/src/net/lookup.go:519
		// _ = "end of CoverTab[7084]"
//line /snap/go/10455/src/net/lookup.go:519
	}() {
//line /snap/go/10455/src/net/lookup.go:519
		_go_fuzz_dep_.CoverTab[528998]++
//line /snap/go/10455/src/net/lookup.go:519
		_go_fuzz_dep_.CoverTab[7085]++
							return "", nil, &DNSError{Err: "SRV header name is invalid", Name: name}
//line /snap/go/10455/src/net/lookup.go:520
		// _ = "end of CoverTab[7085]"
	} else {
//line /snap/go/10455/src/net/lookup.go:521
		_go_fuzz_dep_.CoverTab[528999]++
//line /snap/go/10455/src/net/lookup.go:521
		_go_fuzz_dep_.CoverTab[7086]++
//line /snap/go/10455/src/net/lookup.go:521
		// _ = "end of CoverTab[7086]"
//line /snap/go/10455/src/net/lookup.go:521
	}
//line /snap/go/10455/src/net/lookup.go:521
	// _ = "end of CoverTab[7078]"
//line /snap/go/10455/src/net/lookup.go:521
	_go_fuzz_dep_.CoverTab[7079]++
						filteredAddrs := make([]*SRV, 0, len(addrs))
//line /snap/go/10455/src/net/lookup.go:522
	_go_fuzz_dep_.CoverTab[786705] = 0
						for _, addr := range addrs {
//line /snap/go/10455/src/net/lookup.go:523
		if _go_fuzz_dep_.CoverTab[786705] == 0 {
//line /snap/go/10455/src/net/lookup.go:523
			_go_fuzz_dep_.CoverTab[529106]++
//line /snap/go/10455/src/net/lookup.go:523
		} else {
//line /snap/go/10455/src/net/lookup.go:523
			_go_fuzz_dep_.CoverTab[529107]++
//line /snap/go/10455/src/net/lookup.go:523
		}
//line /snap/go/10455/src/net/lookup.go:523
		_go_fuzz_dep_.CoverTab[786705] = 1
//line /snap/go/10455/src/net/lookup.go:523
		_go_fuzz_dep_.CoverTab[7087]++
							if addr == nil {
//line /snap/go/10455/src/net/lookup.go:524
			_go_fuzz_dep_.CoverTab[529000]++
//line /snap/go/10455/src/net/lookup.go:524
			_go_fuzz_dep_.CoverTab[7090]++
								continue
//line /snap/go/10455/src/net/lookup.go:525
			// _ = "end of CoverTab[7090]"
		} else {
//line /snap/go/10455/src/net/lookup.go:526
			_go_fuzz_dep_.CoverTab[529001]++
//line /snap/go/10455/src/net/lookup.go:526
			_go_fuzz_dep_.CoverTab[7091]++
//line /snap/go/10455/src/net/lookup.go:526
			// _ = "end of CoverTab[7091]"
//line /snap/go/10455/src/net/lookup.go:526
		}
//line /snap/go/10455/src/net/lookup.go:526
		// _ = "end of CoverTab[7087]"
//line /snap/go/10455/src/net/lookup.go:526
		_go_fuzz_dep_.CoverTab[7088]++
							if !isDomainName(addr.Target) {
//line /snap/go/10455/src/net/lookup.go:527
			_go_fuzz_dep_.CoverTab[529002]++
//line /snap/go/10455/src/net/lookup.go:527
			_go_fuzz_dep_.CoverTab[7092]++
								continue
//line /snap/go/10455/src/net/lookup.go:528
			// _ = "end of CoverTab[7092]"
		} else {
//line /snap/go/10455/src/net/lookup.go:529
			_go_fuzz_dep_.CoverTab[529003]++
//line /snap/go/10455/src/net/lookup.go:529
			_go_fuzz_dep_.CoverTab[7093]++
//line /snap/go/10455/src/net/lookup.go:529
			// _ = "end of CoverTab[7093]"
//line /snap/go/10455/src/net/lookup.go:529
		}
//line /snap/go/10455/src/net/lookup.go:529
		// _ = "end of CoverTab[7088]"
//line /snap/go/10455/src/net/lookup.go:529
		_go_fuzz_dep_.CoverTab[7089]++
							filteredAddrs = append(filteredAddrs, addr)
//line /snap/go/10455/src/net/lookup.go:530
		// _ = "end of CoverTab[7089]"
	}
//line /snap/go/10455/src/net/lookup.go:531
	if _go_fuzz_dep_.CoverTab[786705] == 0 {
//line /snap/go/10455/src/net/lookup.go:531
		_go_fuzz_dep_.CoverTab[529108]++
//line /snap/go/10455/src/net/lookup.go:531
	} else {
//line /snap/go/10455/src/net/lookup.go:531
		_go_fuzz_dep_.CoverTab[529109]++
//line /snap/go/10455/src/net/lookup.go:531
	}
//line /snap/go/10455/src/net/lookup.go:531
	// _ = "end of CoverTab[7079]"
//line /snap/go/10455/src/net/lookup.go:531
	_go_fuzz_dep_.CoverTab[7080]++
						if len(addrs) != len(filteredAddrs) {
//line /snap/go/10455/src/net/lookup.go:532
		_go_fuzz_dep_.CoverTab[529004]++
//line /snap/go/10455/src/net/lookup.go:532
		_go_fuzz_dep_.CoverTab[7094]++
							return cname, filteredAddrs, &DNSError{Err: errMalformedDNSRecordsDetail, Name: name}
//line /snap/go/10455/src/net/lookup.go:533
		// _ = "end of CoverTab[7094]"
	} else {
//line /snap/go/10455/src/net/lookup.go:534
		_go_fuzz_dep_.CoverTab[529005]++
//line /snap/go/10455/src/net/lookup.go:534
		_go_fuzz_dep_.CoverTab[7095]++
//line /snap/go/10455/src/net/lookup.go:534
		// _ = "end of CoverTab[7095]"
//line /snap/go/10455/src/net/lookup.go:534
	}
//line /snap/go/10455/src/net/lookup.go:534
	// _ = "end of CoverTab[7080]"
//line /snap/go/10455/src/net/lookup.go:534
	_go_fuzz_dep_.CoverTab[7081]++
						return cname, filteredAddrs, nil
//line /snap/go/10455/src/net/lookup.go:535
	// _ = "end of CoverTab[7081]"
}

// LookupMX returns the DNS MX records for the given domain name sorted by preference.
//line /snap/go/10455/src/net/lookup.go:538
//
//line /snap/go/10455/src/net/lookup.go:538
// The returned mail server names are validated to be properly
//line /snap/go/10455/src/net/lookup.go:538
// formatted presentation-format domain names. If the response contains
//line /snap/go/10455/src/net/lookup.go:538
// invalid names, those records are filtered out and an error
//line /snap/go/10455/src/net/lookup.go:538
// will be returned alongside the remaining results, if any.
//line /snap/go/10455/src/net/lookup.go:538
//
//line /snap/go/10455/src/net/lookup.go:538
// LookupMX uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/lookup.go:538
// Resolver.LookupMX.
//line /snap/go/10455/src/net/lookup.go:547
func LookupMX(name string) ([]*MX, error) {
//line /snap/go/10455/src/net/lookup.go:547
	_go_fuzz_dep_.CoverTab[7096]++
						return DefaultResolver.LookupMX(context.Background(), name)
//line /snap/go/10455/src/net/lookup.go:548
	// _ = "end of CoverTab[7096]"
}

// LookupMX returns the DNS MX records for the given domain name sorted by preference.
//line /snap/go/10455/src/net/lookup.go:551
//
//line /snap/go/10455/src/net/lookup.go:551
// The returned mail server names are validated to be properly
//line /snap/go/10455/src/net/lookup.go:551
// formatted presentation-format domain names. If the response contains
//line /snap/go/10455/src/net/lookup.go:551
// invalid names, those records are filtered out and an error
//line /snap/go/10455/src/net/lookup.go:551
// will be returned alongside the remaining results, if any.
//line /snap/go/10455/src/net/lookup.go:557
func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error) {
//line /snap/go/10455/src/net/lookup.go:557
	_go_fuzz_dep_.CoverTab[7097]++
						records, err := r.lookupMX(ctx, name)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:559
		_go_fuzz_dep_.CoverTab[529006]++
//line /snap/go/10455/src/net/lookup.go:559
		_go_fuzz_dep_.CoverTab[7101]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:560
		// _ = "end of CoverTab[7101]"
	} else {
//line /snap/go/10455/src/net/lookup.go:561
		_go_fuzz_dep_.CoverTab[529007]++
//line /snap/go/10455/src/net/lookup.go:561
		_go_fuzz_dep_.CoverTab[7102]++
//line /snap/go/10455/src/net/lookup.go:561
		// _ = "end of CoverTab[7102]"
//line /snap/go/10455/src/net/lookup.go:561
	}
//line /snap/go/10455/src/net/lookup.go:561
	// _ = "end of CoverTab[7097]"
//line /snap/go/10455/src/net/lookup.go:561
	_go_fuzz_dep_.CoverTab[7098]++
						filteredMX := make([]*MX, 0, len(records))
//line /snap/go/10455/src/net/lookup.go:562
	_go_fuzz_dep_.CoverTab[786706] = 0
						for _, mx := range records {
//line /snap/go/10455/src/net/lookup.go:563
		if _go_fuzz_dep_.CoverTab[786706] == 0 {
//line /snap/go/10455/src/net/lookup.go:563
			_go_fuzz_dep_.CoverTab[529110]++
//line /snap/go/10455/src/net/lookup.go:563
		} else {
//line /snap/go/10455/src/net/lookup.go:563
			_go_fuzz_dep_.CoverTab[529111]++
//line /snap/go/10455/src/net/lookup.go:563
		}
//line /snap/go/10455/src/net/lookup.go:563
		_go_fuzz_dep_.CoverTab[786706] = 1
//line /snap/go/10455/src/net/lookup.go:563
		_go_fuzz_dep_.CoverTab[7103]++
							if mx == nil {
//line /snap/go/10455/src/net/lookup.go:564
			_go_fuzz_dep_.CoverTab[529008]++
//line /snap/go/10455/src/net/lookup.go:564
			_go_fuzz_dep_.CoverTab[7106]++
								continue
//line /snap/go/10455/src/net/lookup.go:565
			// _ = "end of CoverTab[7106]"
		} else {
//line /snap/go/10455/src/net/lookup.go:566
			_go_fuzz_dep_.CoverTab[529009]++
//line /snap/go/10455/src/net/lookup.go:566
			_go_fuzz_dep_.CoverTab[7107]++
//line /snap/go/10455/src/net/lookup.go:566
			// _ = "end of CoverTab[7107]"
//line /snap/go/10455/src/net/lookup.go:566
		}
//line /snap/go/10455/src/net/lookup.go:566
		// _ = "end of CoverTab[7103]"
//line /snap/go/10455/src/net/lookup.go:566
		_go_fuzz_dep_.CoverTab[7104]++
							if !isDomainName(mx.Host) {
//line /snap/go/10455/src/net/lookup.go:567
			_go_fuzz_dep_.CoverTab[529010]++
//line /snap/go/10455/src/net/lookup.go:567
			_go_fuzz_dep_.CoverTab[7108]++
								continue
//line /snap/go/10455/src/net/lookup.go:568
			// _ = "end of CoverTab[7108]"
		} else {
//line /snap/go/10455/src/net/lookup.go:569
			_go_fuzz_dep_.CoverTab[529011]++
//line /snap/go/10455/src/net/lookup.go:569
			_go_fuzz_dep_.CoverTab[7109]++
//line /snap/go/10455/src/net/lookup.go:569
			// _ = "end of CoverTab[7109]"
//line /snap/go/10455/src/net/lookup.go:569
		}
//line /snap/go/10455/src/net/lookup.go:569
		// _ = "end of CoverTab[7104]"
//line /snap/go/10455/src/net/lookup.go:569
		_go_fuzz_dep_.CoverTab[7105]++
							filteredMX = append(filteredMX, mx)
//line /snap/go/10455/src/net/lookup.go:570
		// _ = "end of CoverTab[7105]"
	}
//line /snap/go/10455/src/net/lookup.go:571
	if _go_fuzz_dep_.CoverTab[786706] == 0 {
//line /snap/go/10455/src/net/lookup.go:571
		_go_fuzz_dep_.CoverTab[529112]++
//line /snap/go/10455/src/net/lookup.go:571
	} else {
//line /snap/go/10455/src/net/lookup.go:571
		_go_fuzz_dep_.CoverTab[529113]++
//line /snap/go/10455/src/net/lookup.go:571
	}
//line /snap/go/10455/src/net/lookup.go:571
	// _ = "end of CoverTab[7098]"
//line /snap/go/10455/src/net/lookup.go:571
	_go_fuzz_dep_.CoverTab[7099]++
						if len(records) != len(filteredMX) {
//line /snap/go/10455/src/net/lookup.go:572
		_go_fuzz_dep_.CoverTab[529012]++
//line /snap/go/10455/src/net/lookup.go:572
		_go_fuzz_dep_.CoverTab[7110]++
							return filteredMX, &DNSError{Err: errMalformedDNSRecordsDetail, Name: name}
//line /snap/go/10455/src/net/lookup.go:573
		// _ = "end of CoverTab[7110]"
	} else {
//line /snap/go/10455/src/net/lookup.go:574
		_go_fuzz_dep_.CoverTab[529013]++
//line /snap/go/10455/src/net/lookup.go:574
		_go_fuzz_dep_.CoverTab[7111]++
//line /snap/go/10455/src/net/lookup.go:574
		// _ = "end of CoverTab[7111]"
//line /snap/go/10455/src/net/lookup.go:574
	}
//line /snap/go/10455/src/net/lookup.go:574
	// _ = "end of CoverTab[7099]"
//line /snap/go/10455/src/net/lookup.go:574
	_go_fuzz_dep_.CoverTab[7100]++
						return filteredMX, nil
//line /snap/go/10455/src/net/lookup.go:575
	// _ = "end of CoverTab[7100]"
}

// LookupNS returns the DNS NS records for the given domain name.
//line /snap/go/10455/src/net/lookup.go:578
//
//line /snap/go/10455/src/net/lookup.go:578
// The returned name server names are validated to be properly
//line /snap/go/10455/src/net/lookup.go:578
// formatted presentation-format domain names. If the response contains
//line /snap/go/10455/src/net/lookup.go:578
// invalid names, those records are filtered out and an error
//line /snap/go/10455/src/net/lookup.go:578
// will be returned alongside the remaining results, if any.
//line /snap/go/10455/src/net/lookup.go:578
//
//line /snap/go/10455/src/net/lookup.go:578
// LookupNS uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/lookup.go:578
// Resolver.LookupNS.
//line /snap/go/10455/src/net/lookup.go:587
func LookupNS(name string) ([]*NS, error) {
//line /snap/go/10455/src/net/lookup.go:587
	_go_fuzz_dep_.CoverTab[7112]++
						return DefaultResolver.LookupNS(context.Background(), name)
//line /snap/go/10455/src/net/lookup.go:588
	// _ = "end of CoverTab[7112]"
}

// LookupNS returns the DNS NS records for the given domain name.
//line /snap/go/10455/src/net/lookup.go:591
//
//line /snap/go/10455/src/net/lookup.go:591
// The returned name server names are validated to be properly
//line /snap/go/10455/src/net/lookup.go:591
// formatted presentation-format domain names. If the response contains
//line /snap/go/10455/src/net/lookup.go:591
// invalid names, those records are filtered out and an error
//line /snap/go/10455/src/net/lookup.go:591
// will be returned alongside the remaining results, if any.
//line /snap/go/10455/src/net/lookup.go:597
func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error) {
//line /snap/go/10455/src/net/lookup.go:597
	_go_fuzz_dep_.CoverTab[7113]++
						records, err := r.lookupNS(ctx, name)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:599
		_go_fuzz_dep_.CoverTab[529014]++
//line /snap/go/10455/src/net/lookup.go:599
		_go_fuzz_dep_.CoverTab[7117]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:600
		// _ = "end of CoverTab[7117]"
	} else {
//line /snap/go/10455/src/net/lookup.go:601
		_go_fuzz_dep_.CoverTab[529015]++
//line /snap/go/10455/src/net/lookup.go:601
		_go_fuzz_dep_.CoverTab[7118]++
//line /snap/go/10455/src/net/lookup.go:601
		// _ = "end of CoverTab[7118]"
//line /snap/go/10455/src/net/lookup.go:601
	}
//line /snap/go/10455/src/net/lookup.go:601
	// _ = "end of CoverTab[7113]"
//line /snap/go/10455/src/net/lookup.go:601
	_go_fuzz_dep_.CoverTab[7114]++
						filteredNS := make([]*NS, 0, len(records))
//line /snap/go/10455/src/net/lookup.go:602
	_go_fuzz_dep_.CoverTab[786707] = 0
						for _, ns := range records {
//line /snap/go/10455/src/net/lookup.go:603
		if _go_fuzz_dep_.CoverTab[786707] == 0 {
//line /snap/go/10455/src/net/lookup.go:603
			_go_fuzz_dep_.CoverTab[529114]++
//line /snap/go/10455/src/net/lookup.go:603
		} else {
//line /snap/go/10455/src/net/lookup.go:603
			_go_fuzz_dep_.CoverTab[529115]++
//line /snap/go/10455/src/net/lookup.go:603
		}
//line /snap/go/10455/src/net/lookup.go:603
		_go_fuzz_dep_.CoverTab[786707] = 1
//line /snap/go/10455/src/net/lookup.go:603
		_go_fuzz_dep_.CoverTab[7119]++
							if ns == nil {
//line /snap/go/10455/src/net/lookup.go:604
			_go_fuzz_dep_.CoverTab[529016]++
//line /snap/go/10455/src/net/lookup.go:604
			_go_fuzz_dep_.CoverTab[7122]++
								continue
//line /snap/go/10455/src/net/lookup.go:605
			// _ = "end of CoverTab[7122]"
		} else {
//line /snap/go/10455/src/net/lookup.go:606
			_go_fuzz_dep_.CoverTab[529017]++
//line /snap/go/10455/src/net/lookup.go:606
			_go_fuzz_dep_.CoverTab[7123]++
//line /snap/go/10455/src/net/lookup.go:606
			// _ = "end of CoverTab[7123]"
//line /snap/go/10455/src/net/lookup.go:606
		}
//line /snap/go/10455/src/net/lookup.go:606
		// _ = "end of CoverTab[7119]"
//line /snap/go/10455/src/net/lookup.go:606
		_go_fuzz_dep_.CoverTab[7120]++
							if !isDomainName(ns.Host) {
//line /snap/go/10455/src/net/lookup.go:607
			_go_fuzz_dep_.CoverTab[529018]++
//line /snap/go/10455/src/net/lookup.go:607
			_go_fuzz_dep_.CoverTab[7124]++
								continue
//line /snap/go/10455/src/net/lookup.go:608
			// _ = "end of CoverTab[7124]"
		} else {
//line /snap/go/10455/src/net/lookup.go:609
			_go_fuzz_dep_.CoverTab[529019]++
//line /snap/go/10455/src/net/lookup.go:609
			_go_fuzz_dep_.CoverTab[7125]++
//line /snap/go/10455/src/net/lookup.go:609
			// _ = "end of CoverTab[7125]"
//line /snap/go/10455/src/net/lookup.go:609
		}
//line /snap/go/10455/src/net/lookup.go:609
		// _ = "end of CoverTab[7120]"
//line /snap/go/10455/src/net/lookup.go:609
		_go_fuzz_dep_.CoverTab[7121]++
							filteredNS = append(filteredNS, ns)
//line /snap/go/10455/src/net/lookup.go:610
		// _ = "end of CoverTab[7121]"
	}
//line /snap/go/10455/src/net/lookup.go:611
	if _go_fuzz_dep_.CoverTab[786707] == 0 {
//line /snap/go/10455/src/net/lookup.go:611
		_go_fuzz_dep_.CoverTab[529116]++
//line /snap/go/10455/src/net/lookup.go:611
	} else {
//line /snap/go/10455/src/net/lookup.go:611
		_go_fuzz_dep_.CoverTab[529117]++
//line /snap/go/10455/src/net/lookup.go:611
	}
//line /snap/go/10455/src/net/lookup.go:611
	// _ = "end of CoverTab[7114]"
//line /snap/go/10455/src/net/lookup.go:611
	_go_fuzz_dep_.CoverTab[7115]++
						if len(records) != len(filteredNS) {
//line /snap/go/10455/src/net/lookup.go:612
		_go_fuzz_dep_.CoverTab[529020]++
//line /snap/go/10455/src/net/lookup.go:612
		_go_fuzz_dep_.CoverTab[7126]++
							return filteredNS, &DNSError{Err: errMalformedDNSRecordsDetail, Name: name}
//line /snap/go/10455/src/net/lookup.go:613
		// _ = "end of CoverTab[7126]"
	} else {
//line /snap/go/10455/src/net/lookup.go:614
		_go_fuzz_dep_.CoverTab[529021]++
//line /snap/go/10455/src/net/lookup.go:614
		_go_fuzz_dep_.CoverTab[7127]++
//line /snap/go/10455/src/net/lookup.go:614
		// _ = "end of CoverTab[7127]"
//line /snap/go/10455/src/net/lookup.go:614
	}
//line /snap/go/10455/src/net/lookup.go:614
	// _ = "end of CoverTab[7115]"
//line /snap/go/10455/src/net/lookup.go:614
	_go_fuzz_dep_.CoverTab[7116]++
						return filteredNS, nil
//line /snap/go/10455/src/net/lookup.go:615
	// _ = "end of CoverTab[7116]"
}

// LookupTXT returns the DNS TXT records for the given domain name.
//line /snap/go/10455/src/net/lookup.go:618
//
//line /snap/go/10455/src/net/lookup.go:618
// LookupTXT uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/lookup.go:618
// Resolver.LookupTXT.
//line /snap/go/10455/src/net/lookup.go:622
func LookupTXT(name string) ([]string, error) {
//line /snap/go/10455/src/net/lookup.go:622
	_go_fuzz_dep_.CoverTab[7128]++
						return DefaultResolver.lookupTXT(context.Background(), name)
//line /snap/go/10455/src/net/lookup.go:623
	// _ = "end of CoverTab[7128]"
}

// LookupTXT returns the DNS TXT records for the given domain name.
func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error) {
//line /snap/go/10455/src/net/lookup.go:627
	_go_fuzz_dep_.CoverTab[7129]++
						return r.lookupTXT(ctx, name)
//line /snap/go/10455/src/net/lookup.go:628
	// _ = "end of CoverTab[7129]"
}

// LookupAddr performs a reverse lookup for the given address, returning a list
//line /snap/go/10455/src/net/lookup.go:631
// of names mapping to that address.
//line /snap/go/10455/src/net/lookup.go:631
//
//line /snap/go/10455/src/net/lookup.go:631
// The returned names are validated to be properly formatted presentation-format
//line /snap/go/10455/src/net/lookup.go:631
// domain names. If the response contains invalid names, those records are filtered
//line /snap/go/10455/src/net/lookup.go:631
// out and an error will be returned alongside the remaining results, if any.
//line /snap/go/10455/src/net/lookup.go:631
//
//line /snap/go/10455/src/net/lookup.go:631
// When using the host C library resolver, at most one result will be
//line /snap/go/10455/src/net/lookup.go:631
// returned. To bypass the host resolver, use a custom Resolver.
//line /snap/go/10455/src/net/lookup.go:631
//
//line /snap/go/10455/src/net/lookup.go:631
// LookupAddr uses context.Background internally; to specify the context, use
//line /snap/go/10455/src/net/lookup.go:631
// Resolver.LookupAddr.
//line /snap/go/10455/src/net/lookup.go:643
func LookupAddr(addr string) (names []string, err error) {
//line /snap/go/10455/src/net/lookup.go:643
	_go_fuzz_dep_.CoverTab[7130]++
						return DefaultResolver.LookupAddr(context.Background(), addr)
//line /snap/go/10455/src/net/lookup.go:644
	// _ = "end of CoverTab[7130]"
}

// LookupAddr performs a reverse lookup for the given address, returning a list
//line /snap/go/10455/src/net/lookup.go:647
// of names mapping to that address.
//line /snap/go/10455/src/net/lookup.go:647
//
//line /snap/go/10455/src/net/lookup.go:647
// The returned names are validated to be properly formatted presentation-format
//line /snap/go/10455/src/net/lookup.go:647
// domain names. If the response contains invalid names, those records are filtered
//line /snap/go/10455/src/net/lookup.go:647
// out and an error will be returned alongside the remaining results, if any.
//line /snap/go/10455/src/net/lookup.go:653
func (r *Resolver) LookupAddr(ctx context.Context, addr string) ([]string, error) {
//line /snap/go/10455/src/net/lookup.go:653
	_go_fuzz_dep_.CoverTab[7131]++
						names, err := r.lookupAddr(ctx, addr)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:655
		_go_fuzz_dep_.CoverTab[529022]++
//line /snap/go/10455/src/net/lookup.go:655
		_go_fuzz_dep_.CoverTab[7135]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:656
		// _ = "end of CoverTab[7135]"
	} else {
//line /snap/go/10455/src/net/lookup.go:657
		_go_fuzz_dep_.CoverTab[529023]++
//line /snap/go/10455/src/net/lookup.go:657
		_go_fuzz_dep_.CoverTab[7136]++
//line /snap/go/10455/src/net/lookup.go:657
		// _ = "end of CoverTab[7136]"
//line /snap/go/10455/src/net/lookup.go:657
	}
//line /snap/go/10455/src/net/lookup.go:657
	// _ = "end of CoverTab[7131]"
//line /snap/go/10455/src/net/lookup.go:657
	_go_fuzz_dep_.CoverTab[7132]++
						filteredNames := make([]string, 0, len(names))
//line /snap/go/10455/src/net/lookup.go:658
	_go_fuzz_dep_.CoverTab[786708] = 0
						for _, name := range names {
//line /snap/go/10455/src/net/lookup.go:659
		if _go_fuzz_dep_.CoverTab[786708] == 0 {
//line /snap/go/10455/src/net/lookup.go:659
			_go_fuzz_dep_.CoverTab[529118]++
//line /snap/go/10455/src/net/lookup.go:659
		} else {
//line /snap/go/10455/src/net/lookup.go:659
			_go_fuzz_dep_.CoverTab[529119]++
//line /snap/go/10455/src/net/lookup.go:659
		}
//line /snap/go/10455/src/net/lookup.go:659
		_go_fuzz_dep_.CoverTab[786708] = 1
//line /snap/go/10455/src/net/lookup.go:659
		_go_fuzz_dep_.CoverTab[7137]++
							if isDomainName(name) {
//line /snap/go/10455/src/net/lookup.go:660
			_go_fuzz_dep_.CoverTab[529024]++
//line /snap/go/10455/src/net/lookup.go:660
			_go_fuzz_dep_.CoverTab[7138]++
								filteredNames = append(filteredNames, name)
//line /snap/go/10455/src/net/lookup.go:661
			// _ = "end of CoverTab[7138]"
		} else {
//line /snap/go/10455/src/net/lookup.go:662
			_go_fuzz_dep_.CoverTab[529025]++
//line /snap/go/10455/src/net/lookup.go:662
			_go_fuzz_dep_.CoverTab[7139]++
//line /snap/go/10455/src/net/lookup.go:662
			// _ = "end of CoverTab[7139]"
//line /snap/go/10455/src/net/lookup.go:662
		}
//line /snap/go/10455/src/net/lookup.go:662
		// _ = "end of CoverTab[7137]"
	}
//line /snap/go/10455/src/net/lookup.go:663
	if _go_fuzz_dep_.CoverTab[786708] == 0 {
//line /snap/go/10455/src/net/lookup.go:663
		_go_fuzz_dep_.CoverTab[529120]++
//line /snap/go/10455/src/net/lookup.go:663
	} else {
//line /snap/go/10455/src/net/lookup.go:663
		_go_fuzz_dep_.CoverTab[529121]++
//line /snap/go/10455/src/net/lookup.go:663
	}
//line /snap/go/10455/src/net/lookup.go:663
	// _ = "end of CoverTab[7132]"
//line /snap/go/10455/src/net/lookup.go:663
	_go_fuzz_dep_.CoverTab[7133]++
						if len(names) != len(filteredNames) {
//line /snap/go/10455/src/net/lookup.go:664
		_go_fuzz_dep_.CoverTab[529026]++
//line /snap/go/10455/src/net/lookup.go:664
		_go_fuzz_dep_.CoverTab[7140]++
							return filteredNames, &DNSError{Err: errMalformedDNSRecordsDetail, Name: addr}
//line /snap/go/10455/src/net/lookup.go:665
		// _ = "end of CoverTab[7140]"
	} else {
//line /snap/go/10455/src/net/lookup.go:666
		_go_fuzz_dep_.CoverTab[529027]++
//line /snap/go/10455/src/net/lookup.go:666
		_go_fuzz_dep_.CoverTab[7141]++
//line /snap/go/10455/src/net/lookup.go:666
		// _ = "end of CoverTab[7141]"
//line /snap/go/10455/src/net/lookup.go:666
	}
//line /snap/go/10455/src/net/lookup.go:666
	// _ = "end of CoverTab[7133]"
//line /snap/go/10455/src/net/lookup.go:666
	_go_fuzz_dep_.CoverTab[7134]++
						return filteredNames, nil
//line /snap/go/10455/src/net/lookup.go:667
	// _ = "end of CoverTab[7134]"
}

// errMalformedDNSRecordsDetail is the DNSError detail which is returned when a Resolver.Lookup...
//line /snap/go/10455/src/net/lookup.go:670
// method receives DNS records which contain invalid DNS names. This may be returned alongside
//line /snap/go/10455/src/net/lookup.go:670
// results which have had the malformed records filtered out.
//line /snap/go/10455/src/net/lookup.go:673
var errMalformedDNSRecordsDetail = "DNS response contained records which contain invalid names"

// dial makes a new connection to the provided server (which must be
//line /snap/go/10455/src/net/lookup.go:675
// an IP address) with the provided network type, using either r.Dial
//line /snap/go/10455/src/net/lookup.go:675
// (if both r and r.Dial are non-nil) or else Dialer.DialContext.
//line /snap/go/10455/src/net/lookup.go:678
func (r *Resolver) dial(ctx context.Context, network, server string) (Conn, error) {
//line /snap/go/10455/src/net/lookup.go:678
	_go_fuzz_dep_.CoverTab[7142]++
	// Calling Dial here is scary -- we have to be sure not to
	// dial a name that will require a DNS lookup, or Dial will
	// call back here to translate it. The DNS config parser has
	// already checked that all the cfg.servers are IP
	// addresses, which Dial will use without a DNS lookup.
	var c Conn
	var err error
	if r != nil && func() bool {
//line /snap/go/10455/src/net/lookup.go:686
		_go_fuzz_dep_.CoverTab[7145]++
//line /snap/go/10455/src/net/lookup.go:686
		return r.Dial != nil
//line /snap/go/10455/src/net/lookup.go:686
		// _ = "end of CoverTab[7145]"
//line /snap/go/10455/src/net/lookup.go:686
	}() {
//line /snap/go/10455/src/net/lookup.go:686
		_go_fuzz_dep_.CoverTab[529028]++
//line /snap/go/10455/src/net/lookup.go:686
		_go_fuzz_dep_.CoverTab[7146]++
							c, err = r.Dial(ctx, network, server)
//line /snap/go/10455/src/net/lookup.go:687
		// _ = "end of CoverTab[7146]"
	} else {
//line /snap/go/10455/src/net/lookup.go:688
		_go_fuzz_dep_.CoverTab[529029]++
//line /snap/go/10455/src/net/lookup.go:688
		_go_fuzz_dep_.CoverTab[7147]++
							var d Dialer
							c, err = d.DialContext(ctx, network, server)
//line /snap/go/10455/src/net/lookup.go:690
		// _ = "end of CoverTab[7147]"
	}
//line /snap/go/10455/src/net/lookup.go:691
	// _ = "end of CoverTab[7142]"
//line /snap/go/10455/src/net/lookup.go:691
	_go_fuzz_dep_.CoverTab[7143]++
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:692
		_go_fuzz_dep_.CoverTab[529030]++
//line /snap/go/10455/src/net/lookup.go:692
		_go_fuzz_dep_.CoverTab[7148]++
							return nil, mapErr(err)
//line /snap/go/10455/src/net/lookup.go:693
		// _ = "end of CoverTab[7148]"
	} else {
//line /snap/go/10455/src/net/lookup.go:694
		_go_fuzz_dep_.CoverTab[529031]++
//line /snap/go/10455/src/net/lookup.go:694
		_go_fuzz_dep_.CoverTab[7149]++
//line /snap/go/10455/src/net/lookup.go:694
		// _ = "end of CoverTab[7149]"
//line /snap/go/10455/src/net/lookup.go:694
	}
//line /snap/go/10455/src/net/lookup.go:694
	// _ = "end of CoverTab[7143]"
//line /snap/go/10455/src/net/lookup.go:694
	_go_fuzz_dep_.CoverTab[7144]++
						return c, nil
//line /snap/go/10455/src/net/lookup.go:695
	// _ = "end of CoverTab[7144]"
}

// goLookupSRV returns the SRV records for a target name, built either
//line /snap/go/10455/src/net/lookup.go:698
// from its component service ("sip"), protocol ("tcp"), and name
//line /snap/go/10455/src/net/lookup.go:698
// ("example.com."), or from name directly (if service and proto are
//line /snap/go/10455/src/net/lookup.go:698
// both empty).
//line /snap/go/10455/src/net/lookup.go:698
//
//line /snap/go/10455/src/net/lookup.go:698
// In either case, the returned target name ("_sip._tcp.example.com.")
//line /snap/go/10455/src/net/lookup.go:698
// is also returned on success.
//line /snap/go/10455/src/net/lookup.go:698
//
//line /snap/go/10455/src/net/lookup.go:698
// The records are sorted by weight.
//line /snap/go/10455/src/net/lookup.go:707
func (r *Resolver) goLookupSRV(ctx context.Context, service, proto, name string) (target string, srvs []*SRV, err error) {
//line /snap/go/10455/src/net/lookup.go:707
	_go_fuzz_dep_.CoverTab[7150]++
						if service == "" && func() bool {
//line /snap/go/10455/src/net/lookup.go:708
		_go_fuzz_dep_.CoverTab[7154]++
//line /snap/go/10455/src/net/lookup.go:708
		return proto == ""
//line /snap/go/10455/src/net/lookup.go:708
		// _ = "end of CoverTab[7154]"
//line /snap/go/10455/src/net/lookup.go:708
	}() {
//line /snap/go/10455/src/net/lookup.go:708
		_go_fuzz_dep_.CoverTab[529032]++
//line /snap/go/10455/src/net/lookup.go:708
		_go_fuzz_dep_.CoverTab[7155]++
							target = name
//line /snap/go/10455/src/net/lookup.go:709
		// _ = "end of CoverTab[7155]"
	} else {
//line /snap/go/10455/src/net/lookup.go:710
		_go_fuzz_dep_.CoverTab[529033]++
//line /snap/go/10455/src/net/lookup.go:710
		_go_fuzz_dep_.CoverTab[7156]++
							target = "_" + service + "._" + proto + "." + name
//line /snap/go/10455/src/net/lookup.go:711
		// _ = "end of CoverTab[7156]"
	}
//line /snap/go/10455/src/net/lookup.go:712
	// _ = "end of CoverTab[7150]"
//line /snap/go/10455/src/net/lookup.go:712
	_go_fuzz_dep_.CoverTab[7151]++
						p, server, err := r.lookup(ctx, target, dnsmessage.TypeSRV, nil)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:714
		_go_fuzz_dep_.CoverTab[529034]++
//line /snap/go/10455/src/net/lookup.go:714
		_go_fuzz_dep_.CoverTab[7157]++
							return "", nil, err
//line /snap/go/10455/src/net/lookup.go:715
		// _ = "end of CoverTab[7157]"
	} else {
//line /snap/go/10455/src/net/lookup.go:716
		_go_fuzz_dep_.CoverTab[529035]++
//line /snap/go/10455/src/net/lookup.go:716
		_go_fuzz_dep_.CoverTab[7158]++
//line /snap/go/10455/src/net/lookup.go:716
		// _ = "end of CoverTab[7158]"
//line /snap/go/10455/src/net/lookup.go:716
	}
//line /snap/go/10455/src/net/lookup.go:716
	// _ = "end of CoverTab[7151]"
//line /snap/go/10455/src/net/lookup.go:716
	_go_fuzz_dep_.CoverTab[7152]++
						var cname dnsmessage.Name
//line /snap/go/10455/src/net/lookup.go:717
	_go_fuzz_dep_.CoverTab[786709] = 0
						for {
//line /snap/go/10455/src/net/lookup.go:718
		if _go_fuzz_dep_.CoverTab[786709] == 0 {
//line /snap/go/10455/src/net/lookup.go:718
			_go_fuzz_dep_.CoverTab[529122]++
//line /snap/go/10455/src/net/lookup.go:718
		} else {
//line /snap/go/10455/src/net/lookup.go:718
			_go_fuzz_dep_.CoverTab[529123]++
//line /snap/go/10455/src/net/lookup.go:718
		}
//line /snap/go/10455/src/net/lookup.go:718
		_go_fuzz_dep_.CoverTab[786709] = 1
//line /snap/go/10455/src/net/lookup.go:718
		_go_fuzz_dep_.CoverTab[7159]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /snap/go/10455/src/net/lookup.go:720
			_go_fuzz_dep_.CoverTab[529036]++
//line /snap/go/10455/src/net/lookup.go:720
			_go_fuzz_dep_.CoverTab[7165]++
								break
//line /snap/go/10455/src/net/lookup.go:721
			// _ = "end of CoverTab[7165]"
		} else {
//line /snap/go/10455/src/net/lookup.go:722
			_go_fuzz_dep_.CoverTab[529037]++
//line /snap/go/10455/src/net/lookup.go:722
			_go_fuzz_dep_.CoverTab[7166]++
//line /snap/go/10455/src/net/lookup.go:722
			// _ = "end of CoverTab[7166]"
//line /snap/go/10455/src/net/lookup.go:722
		}
//line /snap/go/10455/src/net/lookup.go:722
		// _ = "end of CoverTab[7159]"
//line /snap/go/10455/src/net/lookup.go:722
		_go_fuzz_dep_.CoverTab[7160]++
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:723
			_go_fuzz_dep_.CoverTab[529038]++
//line /snap/go/10455/src/net/lookup.go:723
			_go_fuzz_dep_.CoverTab[7167]++
								return "", nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /snap/go/10455/src/net/lookup.go:728
			// _ = "end of CoverTab[7167]"
		} else {
//line /snap/go/10455/src/net/lookup.go:729
			_go_fuzz_dep_.CoverTab[529039]++
//line /snap/go/10455/src/net/lookup.go:729
			_go_fuzz_dep_.CoverTab[7168]++
//line /snap/go/10455/src/net/lookup.go:729
			// _ = "end of CoverTab[7168]"
//line /snap/go/10455/src/net/lookup.go:729
		}
//line /snap/go/10455/src/net/lookup.go:729
		// _ = "end of CoverTab[7160]"
//line /snap/go/10455/src/net/lookup.go:729
		_go_fuzz_dep_.CoverTab[7161]++
							if h.Type != dnsmessage.TypeSRV {
//line /snap/go/10455/src/net/lookup.go:730
			_go_fuzz_dep_.CoverTab[529040]++
//line /snap/go/10455/src/net/lookup.go:730
			_go_fuzz_dep_.CoverTab[7169]++
								if err := p.SkipAnswer(); err != nil {
//line /snap/go/10455/src/net/lookup.go:731
				_go_fuzz_dep_.CoverTab[529042]++
//line /snap/go/10455/src/net/lookup.go:731
				_go_fuzz_dep_.CoverTab[7171]++
									return "", nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /snap/go/10455/src/net/lookup.go:736
				// _ = "end of CoverTab[7171]"
			} else {
//line /snap/go/10455/src/net/lookup.go:737
				_go_fuzz_dep_.CoverTab[529043]++
//line /snap/go/10455/src/net/lookup.go:737
				_go_fuzz_dep_.CoverTab[7172]++
//line /snap/go/10455/src/net/lookup.go:737
				// _ = "end of CoverTab[7172]"
//line /snap/go/10455/src/net/lookup.go:737
			}
//line /snap/go/10455/src/net/lookup.go:737
			// _ = "end of CoverTab[7169]"
//line /snap/go/10455/src/net/lookup.go:737
			_go_fuzz_dep_.CoverTab[7170]++
								continue
//line /snap/go/10455/src/net/lookup.go:738
			// _ = "end of CoverTab[7170]"
		} else {
//line /snap/go/10455/src/net/lookup.go:739
			_go_fuzz_dep_.CoverTab[529041]++
//line /snap/go/10455/src/net/lookup.go:739
			_go_fuzz_dep_.CoverTab[7173]++
//line /snap/go/10455/src/net/lookup.go:739
			// _ = "end of CoverTab[7173]"
//line /snap/go/10455/src/net/lookup.go:739
		}
//line /snap/go/10455/src/net/lookup.go:739
		// _ = "end of CoverTab[7161]"
//line /snap/go/10455/src/net/lookup.go:739
		_go_fuzz_dep_.CoverTab[7162]++
							if cname.Length == 0 && func() bool {
//line /snap/go/10455/src/net/lookup.go:740
			_go_fuzz_dep_.CoverTab[7174]++
//line /snap/go/10455/src/net/lookup.go:740
			return h.Name.Length != 0
//line /snap/go/10455/src/net/lookup.go:740
			// _ = "end of CoverTab[7174]"
//line /snap/go/10455/src/net/lookup.go:740
		}() {
//line /snap/go/10455/src/net/lookup.go:740
			_go_fuzz_dep_.CoverTab[529044]++
//line /snap/go/10455/src/net/lookup.go:740
			_go_fuzz_dep_.CoverTab[7175]++
								cname = h.Name
//line /snap/go/10455/src/net/lookup.go:741
			// _ = "end of CoverTab[7175]"
		} else {
//line /snap/go/10455/src/net/lookup.go:742
			_go_fuzz_dep_.CoverTab[529045]++
//line /snap/go/10455/src/net/lookup.go:742
			_go_fuzz_dep_.CoverTab[7176]++
//line /snap/go/10455/src/net/lookup.go:742
			// _ = "end of CoverTab[7176]"
//line /snap/go/10455/src/net/lookup.go:742
		}
//line /snap/go/10455/src/net/lookup.go:742
		// _ = "end of CoverTab[7162]"
//line /snap/go/10455/src/net/lookup.go:742
		_go_fuzz_dep_.CoverTab[7163]++
							srv, err := p.SRVResource()
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:744
			_go_fuzz_dep_.CoverTab[529046]++
//line /snap/go/10455/src/net/lookup.go:744
			_go_fuzz_dep_.CoverTab[7177]++
								return "", nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /snap/go/10455/src/net/lookup.go:749
			// _ = "end of CoverTab[7177]"
		} else {
//line /snap/go/10455/src/net/lookup.go:750
			_go_fuzz_dep_.CoverTab[529047]++
//line /snap/go/10455/src/net/lookup.go:750
			_go_fuzz_dep_.CoverTab[7178]++
//line /snap/go/10455/src/net/lookup.go:750
			// _ = "end of CoverTab[7178]"
//line /snap/go/10455/src/net/lookup.go:750
		}
//line /snap/go/10455/src/net/lookup.go:750
		// _ = "end of CoverTab[7163]"
//line /snap/go/10455/src/net/lookup.go:750
		_go_fuzz_dep_.CoverTab[7164]++
							srvs = append(srvs, &SRV{Target: srv.Target.String(), Port: srv.Port, Priority: srv.Priority, Weight: srv.Weight})
//line /snap/go/10455/src/net/lookup.go:751
		// _ = "end of CoverTab[7164]"
	}
//line /snap/go/10455/src/net/lookup.go:752
	// _ = "end of CoverTab[7152]"
//line /snap/go/10455/src/net/lookup.go:752
	_go_fuzz_dep_.CoverTab[7153]++
						byPriorityWeight(srvs).sort()
						return cname.String(), srvs, nil
//line /snap/go/10455/src/net/lookup.go:754
	// _ = "end of CoverTab[7153]"
}

// goLookupMX returns the MX records for name.
func (r *Resolver) goLookupMX(ctx context.Context, name string) ([]*MX, error) {
//line /snap/go/10455/src/net/lookup.go:758
	_go_fuzz_dep_.CoverTab[7179]++
						p, server, err := r.lookup(ctx, name, dnsmessage.TypeMX, nil)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:760
		_go_fuzz_dep_.CoverTab[529048]++
//line /snap/go/10455/src/net/lookup.go:760
		_go_fuzz_dep_.CoverTab[7182]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:761
		// _ = "end of CoverTab[7182]"
	} else {
//line /snap/go/10455/src/net/lookup.go:762
		_go_fuzz_dep_.CoverTab[529049]++
//line /snap/go/10455/src/net/lookup.go:762
		_go_fuzz_dep_.CoverTab[7183]++
//line /snap/go/10455/src/net/lookup.go:762
		// _ = "end of CoverTab[7183]"
//line /snap/go/10455/src/net/lookup.go:762
	}
//line /snap/go/10455/src/net/lookup.go:762
	// _ = "end of CoverTab[7179]"
//line /snap/go/10455/src/net/lookup.go:762
	_go_fuzz_dep_.CoverTab[7180]++
						var mxs []*MX
//line /snap/go/10455/src/net/lookup.go:763
	_go_fuzz_dep_.CoverTab[786710] = 0
						for {
//line /snap/go/10455/src/net/lookup.go:764
		if _go_fuzz_dep_.CoverTab[786710] == 0 {
//line /snap/go/10455/src/net/lookup.go:764
			_go_fuzz_dep_.CoverTab[529126]++
//line /snap/go/10455/src/net/lookup.go:764
		} else {
//line /snap/go/10455/src/net/lookup.go:764
			_go_fuzz_dep_.CoverTab[529127]++
//line /snap/go/10455/src/net/lookup.go:764
		}
//line /snap/go/10455/src/net/lookup.go:764
		_go_fuzz_dep_.CoverTab[786710] = 1
//line /snap/go/10455/src/net/lookup.go:764
		_go_fuzz_dep_.CoverTab[7184]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /snap/go/10455/src/net/lookup.go:766
			_go_fuzz_dep_.CoverTab[529050]++
//line /snap/go/10455/src/net/lookup.go:766
			_go_fuzz_dep_.CoverTab[7189]++
								break
//line /snap/go/10455/src/net/lookup.go:767
			// _ = "end of CoverTab[7189]"
		} else {
//line /snap/go/10455/src/net/lookup.go:768
			_go_fuzz_dep_.CoverTab[529051]++
//line /snap/go/10455/src/net/lookup.go:768
			_go_fuzz_dep_.CoverTab[7190]++
//line /snap/go/10455/src/net/lookup.go:768
			// _ = "end of CoverTab[7190]"
//line /snap/go/10455/src/net/lookup.go:768
		}
//line /snap/go/10455/src/net/lookup.go:768
		// _ = "end of CoverTab[7184]"
//line /snap/go/10455/src/net/lookup.go:768
		_go_fuzz_dep_.CoverTab[7185]++
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:769
			_go_fuzz_dep_.CoverTab[529052]++
//line /snap/go/10455/src/net/lookup.go:769
			_go_fuzz_dep_.CoverTab[7191]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /snap/go/10455/src/net/lookup.go:774
			// _ = "end of CoverTab[7191]"
		} else {
//line /snap/go/10455/src/net/lookup.go:775
			_go_fuzz_dep_.CoverTab[529053]++
//line /snap/go/10455/src/net/lookup.go:775
			_go_fuzz_dep_.CoverTab[7192]++
//line /snap/go/10455/src/net/lookup.go:775
			// _ = "end of CoverTab[7192]"
//line /snap/go/10455/src/net/lookup.go:775
		}
//line /snap/go/10455/src/net/lookup.go:775
		// _ = "end of CoverTab[7185]"
//line /snap/go/10455/src/net/lookup.go:775
		_go_fuzz_dep_.CoverTab[7186]++
							if h.Type != dnsmessage.TypeMX {
//line /snap/go/10455/src/net/lookup.go:776
			_go_fuzz_dep_.CoverTab[529054]++
//line /snap/go/10455/src/net/lookup.go:776
			_go_fuzz_dep_.CoverTab[7193]++
								if err := p.SkipAnswer(); err != nil {
//line /snap/go/10455/src/net/lookup.go:777
				_go_fuzz_dep_.CoverTab[529056]++
//line /snap/go/10455/src/net/lookup.go:777
				_go_fuzz_dep_.CoverTab[7195]++
									return nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /snap/go/10455/src/net/lookup.go:782
				// _ = "end of CoverTab[7195]"
			} else {
//line /snap/go/10455/src/net/lookup.go:783
				_go_fuzz_dep_.CoverTab[529057]++
//line /snap/go/10455/src/net/lookup.go:783
				_go_fuzz_dep_.CoverTab[7196]++
//line /snap/go/10455/src/net/lookup.go:783
				// _ = "end of CoverTab[7196]"
//line /snap/go/10455/src/net/lookup.go:783
			}
//line /snap/go/10455/src/net/lookup.go:783
			// _ = "end of CoverTab[7193]"
//line /snap/go/10455/src/net/lookup.go:783
			_go_fuzz_dep_.CoverTab[7194]++
								continue
//line /snap/go/10455/src/net/lookup.go:784
			// _ = "end of CoverTab[7194]"
		} else {
//line /snap/go/10455/src/net/lookup.go:785
			_go_fuzz_dep_.CoverTab[529055]++
//line /snap/go/10455/src/net/lookup.go:785
			_go_fuzz_dep_.CoverTab[7197]++
//line /snap/go/10455/src/net/lookup.go:785
			// _ = "end of CoverTab[7197]"
//line /snap/go/10455/src/net/lookup.go:785
		}
//line /snap/go/10455/src/net/lookup.go:785
		// _ = "end of CoverTab[7186]"
//line /snap/go/10455/src/net/lookup.go:785
		_go_fuzz_dep_.CoverTab[7187]++
							mx, err := p.MXResource()
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:787
			_go_fuzz_dep_.CoverTab[529058]++
//line /snap/go/10455/src/net/lookup.go:787
			_go_fuzz_dep_.CoverTab[7198]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /snap/go/10455/src/net/lookup.go:792
			// _ = "end of CoverTab[7198]"
		} else {
//line /snap/go/10455/src/net/lookup.go:793
			_go_fuzz_dep_.CoverTab[529059]++
//line /snap/go/10455/src/net/lookup.go:793
			_go_fuzz_dep_.CoverTab[7199]++
//line /snap/go/10455/src/net/lookup.go:793
			// _ = "end of CoverTab[7199]"
//line /snap/go/10455/src/net/lookup.go:793
		}
//line /snap/go/10455/src/net/lookup.go:793
		// _ = "end of CoverTab[7187]"
//line /snap/go/10455/src/net/lookup.go:793
		_go_fuzz_dep_.CoverTab[7188]++
							mxs = append(mxs, &MX{Host: mx.MX.String(), Pref: mx.Pref})
//line /snap/go/10455/src/net/lookup.go:794
		// _ = "end of CoverTab[7188]"

	}
//line /snap/go/10455/src/net/lookup.go:796
	// _ = "end of CoverTab[7180]"
//line /snap/go/10455/src/net/lookup.go:796
	_go_fuzz_dep_.CoverTab[7181]++
						byPref(mxs).sort()
						return mxs, nil
//line /snap/go/10455/src/net/lookup.go:798
	// _ = "end of CoverTab[7181]"
}

// goLookupNS returns the NS records for name.
func (r *Resolver) goLookupNS(ctx context.Context, name string) ([]*NS, error) {
//line /snap/go/10455/src/net/lookup.go:802
	_go_fuzz_dep_.CoverTab[7200]++
						p, server, err := r.lookup(ctx, name, dnsmessage.TypeNS, nil)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:804
		_go_fuzz_dep_.CoverTab[529060]++
//line /snap/go/10455/src/net/lookup.go:804
		_go_fuzz_dep_.CoverTab[7203]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:805
		// _ = "end of CoverTab[7203]"
	} else {
//line /snap/go/10455/src/net/lookup.go:806
		_go_fuzz_dep_.CoverTab[529061]++
//line /snap/go/10455/src/net/lookup.go:806
		_go_fuzz_dep_.CoverTab[7204]++
//line /snap/go/10455/src/net/lookup.go:806
		// _ = "end of CoverTab[7204]"
//line /snap/go/10455/src/net/lookup.go:806
	}
//line /snap/go/10455/src/net/lookup.go:806
	// _ = "end of CoverTab[7200]"
//line /snap/go/10455/src/net/lookup.go:806
	_go_fuzz_dep_.CoverTab[7201]++
						var nss []*NS
//line /snap/go/10455/src/net/lookup.go:807
	_go_fuzz_dep_.CoverTab[786711] = 0
						for {
//line /snap/go/10455/src/net/lookup.go:808
		if _go_fuzz_dep_.CoverTab[786711] == 0 {
//line /snap/go/10455/src/net/lookup.go:808
			_go_fuzz_dep_.CoverTab[529130]++
//line /snap/go/10455/src/net/lookup.go:808
		} else {
//line /snap/go/10455/src/net/lookup.go:808
			_go_fuzz_dep_.CoverTab[529131]++
//line /snap/go/10455/src/net/lookup.go:808
		}
//line /snap/go/10455/src/net/lookup.go:808
		_go_fuzz_dep_.CoverTab[786711] = 1
//line /snap/go/10455/src/net/lookup.go:808
		_go_fuzz_dep_.CoverTab[7205]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /snap/go/10455/src/net/lookup.go:810
			_go_fuzz_dep_.CoverTab[529062]++
//line /snap/go/10455/src/net/lookup.go:810
			_go_fuzz_dep_.CoverTab[7210]++
								break
//line /snap/go/10455/src/net/lookup.go:811
			// _ = "end of CoverTab[7210]"
		} else {
//line /snap/go/10455/src/net/lookup.go:812
			_go_fuzz_dep_.CoverTab[529063]++
//line /snap/go/10455/src/net/lookup.go:812
			_go_fuzz_dep_.CoverTab[7211]++
//line /snap/go/10455/src/net/lookup.go:812
			// _ = "end of CoverTab[7211]"
//line /snap/go/10455/src/net/lookup.go:812
		}
//line /snap/go/10455/src/net/lookup.go:812
		// _ = "end of CoverTab[7205]"
//line /snap/go/10455/src/net/lookup.go:812
		_go_fuzz_dep_.CoverTab[7206]++
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:813
			_go_fuzz_dep_.CoverTab[529064]++
//line /snap/go/10455/src/net/lookup.go:813
			_go_fuzz_dep_.CoverTab[7212]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /snap/go/10455/src/net/lookup.go:818
			// _ = "end of CoverTab[7212]"
		} else {
//line /snap/go/10455/src/net/lookup.go:819
			_go_fuzz_dep_.CoverTab[529065]++
//line /snap/go/10455/src/net/lookup.go:819
			_go_fuzz_dep_.CoverTab[7213]++
//line /snap/go/10455/src/net/lookup.go:819
			// _ = "end of CoverTab[7213]"
//line /snap/go/10455/src/net/lookup.go:819
		}
//line /snap/go/10455/src/net/lookup.go:819
		// _ = "end of CoverTab[7206]"
//line /snap/go/10455/src/net/lookup.go:819
		_go_fuzz_dep_.CoverTab[7207]++
							if h.Type != dnsmessage.TypeNS {
//line /snap/go/10455/src/net/lookup.go:820
			_go_fuzz_dep_.CoverTab[529066]++
//line /snap/go/10455/src/net/lookup.go:820
			_go_fuzz_dep_.CoverTab[7214]++
								if err := p.SkipAnswer(); err != nil {
//line /snap/go/10455/src/net/lookup.go:821
				_go_fuzz_dep_.CoverTab[529068]++
//line /snap/go/10455/src/net/lookup.go:821
				_go_fuzz_dep_.CoverTab[7216]++
									return nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /snap/go/10455/src/net/lookup.go:826
				// _ = "end of CoverTab[7216]"
			} else {
//line /snap/go/10455/src/net/lookup.go:827
				_go_fuzz_dep_.CoverTab[529069]++
//line /snap/go/10455/src/net/lookup.go:827
				_go_fuzz_dep_.CoverTab[7217]++
//line /snap/go/10455/src/net/lookup.go:827
				// _ = "end of CoverTab[7217]"
//line /snap/go/10455/src/net/lookup.go:827
			}
//line /snap/go/10455/src/net/lookup.go:827
			// _ = "end of CoverTab[7214]"
//line /snap/go/10455/src/net/lookup.go:827
			_go_fuzz_dep_.CoverTab[7215]++
								continue
//line /snap/go/10455/src/net/lookup.go:828
			// _ = "end of CoverTab[7215]"
		} else {
//line /snap/go/10455/src/net/lookup.go:829
			_go_fuzz_dep_.CoverTab[529067]++
//line /snap/go/10455/src/net/lookup.go:829
			_go_fuzz_dep_.CoverTab[7218]++
//line /snap/go/10455/src/net/lookup.go:829
			// _ = "end of CoverTab[7218]"
//line /snap/go/10455/src/net/lookup.go:829
		}
//line /snap/go/10455/src/net/lookup.go:829
		// _ = "end of CoverTab[7207]"
//line /snap/go/10455/src/net/lookup.go:829
		_go_fuzz_dep_.CoverTab[7208]++
							ns, err := p.NSResource()
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:831
			_go_fuzz_dep_.CoverTab[529070]++
//line /snap/go/10455/src/net/lookup.go:831
			_go_fuzz_dep_.CoverTab[7219]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /snap/go/10455/src/net/lookup.go:836
			// _ = "end of CoverTab[7219]"
		} else {
//line /snap/go/10455/src/net/lookup.go:837
			_go_fuzz_dep_.CoverTab[529071]++
//line /snap/go/10455/src/net/lookup.go:837
			_go_fuzz_dep_.CoverTab[7220]++
//line /snap/go/10455/src/net/lookup.go:837
			// _ = "end of CoverTab[7220]"
//line /snap/go/10455/src/net/lookup.go:837
		}
//line /snap/go/10455/src/net/lookup.go:837
		// _ = "end of CoverTab[7208]"
//line /snap/go/10455/src/net/lookup.go:837
		_go_fuzz_dep_.CoverTab[7209]++
							nss = append(nss, &NS{Host: ns.NS.String()})
//line /snap/go/10455/src/net/lookup.go:838
		// _ = "end of CoverTab[7209]"
	}
//line /snap/go/10455/src/net/lookup.go:839
	// _ = "end of CoverTab[7201]"
//line /snap/go/10455/src/net/lookup.go:839
	_go_fuzz_dep_.CoverTab[7202]++
						return nss, nil
//line /snap/go/10455/src/net/lookup.go:840
	// _ = "end of CoverTab[7202]"
}

// goLookupTXT returns the TXT records from name.
func (r *Resolver) goLookupTXT(ctx context.Context, name string) ([]string, error) {
//line /snap/go/10455/src/net/lookup.go:844
	_go_fuzz_dep_.CoverTab[7221]++
						p, server, err := r.lookup(ctx, name, dnsmessage.TypeTXT, nil)
						if err != nil {
//line /snap/go/10455/src/net/lookup.go:846
		_go_fuzz_dep_.CoverTab[529072]++
//line /snap/go/10455/src/net/lookup.go:846
		_go_fuzz_dep_.CoverTab[7224]++
							return nil, err
//line /snap/go/10455/src/net/lookup.go:847
		// _ = "end of CoverTab[7224]"
	} else {
//line /snap/go/10455/src/net/lookup.go:848
		_go_fuzz_dep_.CoverTab[529073]++
//line /snap/go/10455/src/net/lookup.go:848
		_go_fuzz_dep_.CoverTab[7225]++
//line /snap/go/10455/src/net/lookup.go:848
		// _ = "end of CoverTab[7225]"
//line /snap/go/10455/src/net/lookup.go:848
	}
//line /snap/go/10455/src/net/lookup.go:848
	// _ = "end of CoverTab[7221]"
//line /snap/go/10455/src/net/lookup.go:848
	_go_fuzz_dep_.CoverTab[7222]++
						var txts []string
//line /snap/go/10455/src/net/lookup.go:849
	_go_fuzz_dep_.CoverTab[786712] = 0
						for {
//line /snap/go/10455/src/net/lookup.go:850
		if _go_fuzz_dep_.CoverTab[786712] == 0 {
//line /snap/go/10455/src/net/lookup.go:850
			_go_fuzz_dep_.CoverTab[529134]++
//line /snap/go/10455/src/net/lookup.go:850
		} else {
//line /snap/go/10455/src/net/lookup.go:850
			_go_fuzz_dep_.CoverTab[529135]++
//line /snap/go/10455/src/net/lookup.go:850
		}
//line /snap/go/10455/src/net/lookup.go:850
		_go_fuzz_dep_.CoverTab[786712] = 1
//line /snap/go/10455/src/net/lookup.go:850
		_go_fuzz_dep_.CoverTab[7226]++
							h, err := p.AnswerHeader()
							if err == dnsmessage.ErrSectionDone {
//line /snap/go/10455/src/net/lookup.go:852
			_go_fuzz_dep_.CoverTab[529074]++
//line /snap/go/10455/src/net/lookup.go:852
			_go_fuzz_dep_.CoverTab[7234]++
								break
//line /snap/go/10455/src/net/lookup.go:853
			// _ = "end of CoverTab[7234]"
		} else {
//line /snap/go/10455/src/net/lookup.go:854
			_go_fuzz_dep_.CoverTab[529075]++
//line /snap/go/10455/src/net/lookup.go:854
			_go_fuzz_dep_.CoverTab[7235]++
//line /snap/go/10455/src/net/lookup.go:854
			// _ = "end of CoverTab[7235]"
//line /snap/go/10455/src/net/lookup.go:854
		}
//line /snap/go/10455/src/net/lookup.go:854
		// _ = "end of CoverTab[7226]"
//line /snap/go/10455/src/net/lookup.go:854
		_go_fuzz_dep_.CoverTab[7227]++
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:855
			_go_fuzz_dep_.CoverTab[529076]++
//line /snap/go/10455/src/net/lookup.go:855
			_go_fuzz_dep_.CoverTab[7236]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /snap/go/10455/src/net/lookup.go:860
			// _ = "end of CoverTab[7236]"
		} else {
//line /snap/go/10455/src/net/lookup.go:861
			_go_fuzz_dep_.CoverTab[529077]++
//line /snap/go/10455/src/net/lookup.go:861
			_go_fuzz_dep_.CoverTab[7237]++
//line /snap/go/10455/src/net/lookup.go:861
			// _ = "end of CoverTab[7237]"
//line /snap/go/10455/src/net/lookup.go:861
		}
//line /snap/go/10455/src/net/lookup.go:861
		// _ = "end of CoverTab[7227]"
//line /snap/go/10455/src/net/lookup.go:861
		_go_fuzz_dep_.CoverTab[7228]++
							if h.Type != dnsmessage.TypeTXT {
//line /snap/go/10455/src/net/lookup.go:862
			_go_fuzz_dep_.CoverTab[529078]++
//line /snap/go/10455/src/net/lookup.go:862
			_go_fuzz_dep_.CoverTab[7238]++
								if err := p.SkipAnswer(); err != nil {
//line /snap/go/10455/src/net/lookup.go:863
				_go_fuzz_dep_.CoverTab[529080]++
//line /snap/go/10455/src/net/lookup.go:863
				_go_fuzz_dep_.CoverTab[7240]++
									return nil, &DNSError{
					Err:	"cannot unmarshal DNS message",
					Name:	name,
					Server:	server,
				}
//line /snap/go/10455/src/net/lookup.go:868
				// _ = "end of CoverTab[7240]"
			} else {
//line /snap/go/10455/src/net/lookup.go:869
				_go_fuzz_dep_.CoverTab[529081]++
//line /snap/go/10455/src/net/lookup.go:869
				_go_fuzz_dep_.CoverTab[7241]++
//line /snap/go/10455/src/net/lookup.go:869
				// _ = "end of CoverTab[7241]"
//line /snap/go/10455/src/net/lookup.go:869
			}
//line /snap/go/10455/src/net/lookup.go:869
			// _ = "end of CoverTab[7238]"
//line /snap/go/10455/src/net/lookup.go:869
			_go_fuzz_dep_.CoverTab[7239]++
								continue
//line /snap/go/10455/src/net/lookup.go:870
			// _ = "end of CoverTab[7239]"
		} else {
//line /snap/go/10455/src/net/lookup.go:871
			_go_fuzz_dep_.CoverTab[529079]++
//line /snap/go/10455/src/net/lookup.go:871
			_go_fuzz_dep_.CoverTab[7242]++
//line /snap/go/10455/src/net/lookup.go:871
			// _ = "end of CoverTab[7242]"
//line /snap/go/10455/src/net/lookup.go:871
		}
//line /snap/go/10455/src/net/lookup.go:871
		// _ = "end of CoverTab[7228]"
//line /snap/go/10455/src/net/lookup.go:871
		_go_fuzz_dep_.CoverTab[7229]++
							txt, err := p.TXTResource()
							if err != nil {
//line /snap/go/10455/src/net/lookup.go:873
			_go_fuzz_dep_.CoverTab[529082]++
//line /snap/go/10455/src/net/lookup.go:873
			_go_fuzz_dep_.CoverTab[7243]++
								return nil, &DNSError{
				Err:	"cannot unmarshal DNS message",
				Name:	name,
				Server:	server,
			}
//line /snap/go/10455/src/net/lookup.go:878
			// _ = "end of CoverTab[7243]"
		} else {
//line /snap/go/10455/src/net/lookup.go:879
			_go_fuzz_dep_.CoverTab[529083]++
//line /snap/go/10455/src/net/lookup.go:879
			_go_fuzz_dep_.CoverTab[7244]++
//line /snap/go/10455/src/net/lookup.go:879
			// _ = "end of CoverTab[7244]"
//line /snap/go/10455/src/net/lookup.go:879
		}
//line /snap/go/10455/src/net/lookup.go:879
		// _ = "end of CoverTab[7229]"
//line /snap/go/10455/src/net/lookup.go:879
		_go_fuzz_dep_.CoverTab[7230]++

//line /snap/go/10455/src/net/lookup.go:883
		n := 0
//line /snap/go/10455/src/net/lookup.go:883
		_go_fuzz_dep_.CoverTab[786713] = 0
							for _, s := range txt.TXT {
//line /snap/go/10455/src/net/lookup.go:884
			if _go_fuzz_dep_.CoverTab[786713] == 0 {
//line /snap/go/10455/src/net/lookup.go:884
				_go_fuzz_dep_.CoverTab[529138]++
//line /snap/go/10455/src/net/lookup.go:884
			} else {
//line /snap/go/10455/src/net/lookup.go:884
				_go_fuzz_dep_.CoverTab[529139]++
//line /snap/go/10455/src/net/lookup.go:884
			}
//line /snap/go/10455/src/net/lookup.go:884
			_go_fuzz_dep_.CoverTab[786713] = 1
//line /snap/go/10455/src/net/lookup.go:884
			_go_fuzz_dep_.CoverTab[7245]++
								n += len(s)
//line /snap/go/10455/src/net/lookup.go:885
			// _ = "end of CoverTab[7245]"
		}
//line /snap/go/10455/src/net/lookup.go:886
		if _go_fuzz_dep_.CoverTab[786713] == 0 {
//line /snap/go/10455/src/net/lookup.go:886
			_go_fuzz_dep_.CoverTab[529140]++
//line /snap/go/10455/src/net/lookup.go:886
		} else {
//line /snap/go/10455/src/net/lookup.go:886
			_go_fuzz_dep_.CoverTab[529141]++
//line /snap/go/10455/src/net/lookup.go:886
		}
//line /snap/go/10455/src/net/lookup.go:886
		// _ = "end of CoverTab[7230]"
//line /snap/go/10455/src/net/lookup.go:886
		_go_fuzz_dep_.CoverTab[7231]++
							txtJoin := make([]byte, 0, n)
//line /snap/go/10455/src/net/lookup.go:887
		_go_fuzz_dep_.CoverTab[786714] = 0
							for _, s := range txt.TXT {
//line /snap/go/10455/src/net/lookup.go:888
			if _go_fuzz_dep_.CoverTab[786714] == 0 {
//line /snap/go/10455/src/net/lookup.go:888
				_go_fuzz_dep_.CoverTab[529142]++
//line /snap/go/10455/src/net/lookup.go:888
			} else {
//line /snap/go/10455/src/net/lookup.go:888
				_go_fuzz_dep_.CoverTab[529143]++
//line /snap/go/10455/src/net/lookup.go:888
			}
//line /snap/go/10455/src/net/lookup.go:888
			_go_fuzz_dep_.CoverTab[786714] = 1
//line /snap/go/10455/src/net/lookup.go:888
			_go_fuzz_dep_.CoverTab[7246]++
								txtJoin = append(txtJoin, s...)
//line /snap/go/10455/src/net/lookup.go:889
			// _ = "end of CoverTab[7246]"
		}
//line /snap/go/10455/src/net/lookup.go:890
		if _go_fuzz_dep_.CoverTab[786714] == 0 {
//line /snap/go/10455/src/net/lookup.go:890
			_go_fuzz_dep_.CoverTab[529144]++
//line /snap/go/10455/src/net/lookup.go:890
		} else {
//line /snap/go/10455/src/net/lookup.go:890
			_go_fuzz_dep_.CoverTab[529145]++
//line /snap/go/10455/src/net/lookup.go:890
		}
//line /snap/go/10455/src/net/lookup.go:890
		// _ = "end of CoverTab[7231]"
//line /snap/go/10455/src/net/lookup.go:890
		_go_fuzz_dep_.CoverTab[7232]++
							if len(txts) == 0 {
//line /snap/go/10455/src/net/lookup.go:891
			_go_fuzz_dep_.CoverTab[529084]++
//line /snap/go/10455/src/net/lookup.go:891
			_go_fuzz_dep_.CoverTab[7247]++
								txts = make([]string, 0, 1)
//line /snap/go/10455/src/net/lookup.go:892
			// _ = "end of CoverTab[7247]"
		} else {
//line /snap/go/10455/src/net/lookup.go:893
			_go_fuzz_dep_.CoverTab[529085]++
//line /snap/go/10455/src/net/lookup.go:893
			_go_fuzz_dep_.CoverTab[7248]++
//line /snap/go/10455/src/net/lookup.go:893
			// _ = "end of CoverTab[7248]"
//line /snap/go/10455/src/net/lookup.go:893
		}
//line /snap/go/10455/src/net/lookup.go:893
		// _ = "end of CoverTab[7232]"
//line /snap/go/10455/src/net/lookup.go:893
		_go_fuzz_dep_.CoverTab[7233]++
							txts = append(txts, string(txtJoin))
//line /snap/go/10455/src/net/lookup.go:894
		// _ = "end of CoverTab[7233]"
	}
//line /snap/go/10455/src/net/lookup.go:895
	// _ = "end of CoverTab[7222]"
//line /snap/go/10455/src/net/lookup.go:895
	_go_fuzz_dep_.CoverTab[7223]++
						return txts, nil
//line /snap/go/10455/src/net/lookup.go:896
	// _ = "end of CoverTab[7223]"
}

func parseCNAMEFromResources(resources []dnsmessage.Resource) (string, error) {
//line /snap/go/10455/src/net/lookup.go:899
	_go_fuzz_dep_.CoverTab[7249]++
						if len(resources) == 0 {
//line /snap/go/10455/src/net/lookup.go:900
		_go_fuzz_dep_.CoverTab[529086]++
//line /snap/go/10455/src/net/lookup.go:900
		_go_fuzz_dep_.CoverTab[7252]++
							return "", errors.New("no CNAME record received")
//line /snap/go/10455/src/net/lookup.go:901
		// _ = "end of CoverTab[7252]"
	} else {
//line /snap/go/10455/src/net/lookup.go:902
		_go_fuzz_dep_.CoverTab[529087]++
//line /snap/go/10455/src/net/lookup.go:902
		_go_fuzz_dep_.CoverTab[7253]++
//line /snap/go/10455/src/net/lookup.go:902
		// _ = "end of CoverTab[7253]"
//line /snap/go/10455/src/net/lookup.go:902
	}
//line /snap/go/10455/src/net/lookup.go:902
	// _ = "end of CoverTab[7249]"
//line /snap/go/10455/src/net/lookup.go:902
	_go_fuzz_dep_.CoverTab[7250]++
						c, ok := resources[0].Body.(*dnsmessage.CNAMEResource)
						if !ok {
//line /snap/go/10455/src/net/lookup.go:904
		_go_fuzz_dep_.CoverTab[529088]++
//line /snap/go/10455/src/net/lookup.go:904
		_go_fuzz_dep_.CoverTab[7254]++
							return "", errors.New("could not parse CNAME record")
//line /snap/go/10455/src/net/lookup.go:905
		// _ = "end of CoverTab[7254]"
	} else {
//line /snap/go/10455/src/net/lookup.go:906
		_go_fuzz_dep_.CoverTab[529089]++
//line /snap/go/10455/src/net/lookup.go:906
		_go_fuzz_dep_.CoverTab[7255]++
//line /snap/go/10455/src/net/lookup.go:906
		// _ = "end of CoverTab[7255]"
//line /snap/go/10455/src/net/lookup.go:906
	}
//line /snap/go/10455/src/net/lookup.go:906
	// _ = "end of CoverTab[7250]"
//line /snap/go/10455/src/net/lookup.go:906
	_go_fuzz_dep_.CoverTab[7251]++
						return c.CNAME.String(), nil
//line /snap/go/10455/src/net/lookup.go:907
	// _ = "end of CoverTab[7251]"
}

//line /snap/go/10455/src/net/lookup.go:908
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/lookup.go:908
var _ = _go_fuzz_dep_.CoverTab
