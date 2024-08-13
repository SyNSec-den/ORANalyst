// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

//line /snap/go/10455/src/net/ipsock_posix.go:7
package net

//line /snap/go/10455/src/net/ipsock_posix.go:7
import (
//line /snap/go/10455/src/net/ipsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/ipsock_posix.go:7
)
//line /snap/go/10455/src/net/ipsock_posix.go:7
import (
//line /snap/go/10455/src/net/ipsock_posix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/ipsock_posix.go:7
)

import (
	"context"
	"internal/poll"
	"net/netip"
	"runtime"
	"syscall"
)

// probe probes IPv4, IPv6 and IPv4-mapped IPv6 communication
//line /snap/go/10455/src/net/ipsock_posix.go:17
// capabilities which are controlled by the IPV6_V6ONLY socket option
//line /snap/go/10455/src/net/ipsock_posix.go:17
// and kernel configuration.
//line /snap/go/10455/src/net/ipsock_posix.go:17
//
//line /snap/go/10455/src/net/ipsock_posix.go:17
// Should we try to use the IPv4 socket interface if we're only
//line /snap/go/10455/src/net/ipsock_posix.go:17
// dealing with IPv4 sockets? As long as the host system understands
//line /snap/go/10455/src/net/ipsock_posix.go:17
// IPv4-mapped IPv6, it's okay to pass IPv4-mapped IPv6 addresses to
//line /snap/go/10455/src/net/ipsock_posix.go:17
// the IPv6 interface. That simplifies our code and is most
//line /snap/go/10455/src/net/ipsock_posix.go:17
// general. Unfortunately, we need to run on kernels built without
//line /snap/go/10455/src/net/ipsock_posix.go:17
// IPv6 support too. So probe the kernel to figure it out.
//line /snap/go/10455/src/net/ipsock_posix.go:27
func (p *ipStackCapabilities) probe() {
//line /snap/go/10455/src/net/ipsock_posix.go:27
	_go_fuzz_dep_.CoverTab[6838]++
							s, err := sysSocket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
							switch err {
	case syscall.EAFNOSUPPORT, syscall.EPROTONOSUPPORT:
//line /snap/go/10455/src/net/ipsock_posix.go:30
		_go_fuzz_dep_.CoverTab[528875]++
//line /snap/go/10455/src/net/ipsock_posix.go:30
		_go_fuzz_dep_.CoverTab[6841]++
//line /snap/go/10455/src/net/ipsock_posix.go:30
		// _ = "end of CoverTab[6841]"
	case nil:
//line /snap/go/10455/src/net/ipsock_posix.go:31
		_go_fuzz_dep_.CoverTab[528876]++
//line /snap/go/10455/src/net/ipsock_posix.go:31
		_go_fuzz_dep_.CoverTab[6842]++
								poll.CloseFunc(s)
								p.ipv4Enabled = true
//line /snap/go/10455/src/net/ipsock_posix.go:33
		// _ = "end of CoverTab[6842]"
//line /snap/go/10455/src/net/ipsock_posix.go:33
	default:
//line /snap/go/10455/src/net/ipsock_posix.go:33
		_go_fuzz_dep_.CoverTab[528877]++
//line /snap/go/10455/src/net/ipsock_posix.go:33
		_go_fuzz_dep_.CoverTab[6843]++
//line /snap/go/10455/src/net/ipsock_posix.go:33
		// _ = "end of CoverTab[6843]"
	}
//line /snap/go/10455/src/net/ipsock_posix.go:34
	// _ = "end of CoverTab[6838]"
//line /snap/go/10455/src/net/ipsock_posix.go:34
	_go_fuzz_dep_.CoverTab[6839]++
							var probes = []struct {
		laddr	TCPAddr
		value	int
	}{

		{laddr: TCPAddr{IP: ParseIP("::1")}, value: 1},

		{laddr: TCPAddr{IP: IPv4(127, 0, 0, 1)}, value: 0},
	}
	switch runtime.GOOS {
	case "dragonfly", "openbsd":
//line /snap/go/10455/src/net/ipsock_posix.go:45
		_go_fuzz_dep_.CoverTab[528878]++
//line /snap/go/10455/src/net/ipsock_posix.go:45
		_go_fuzz_dep_.CoverTab[6844]++

//line /snap/go/10455/src/net/ipsock_posix.go:49
		probes = probes[:1]
//line /snap/go/10455/src/net/ipsock_posix.go:49
		// _ = "end of CoverTab[6844]"
//line /snap/go/10455/src/net/ipsock_posix.go:49
	default:
//line /snap/go/10455/src/net/ipsock_posix.go:49
		_go_fuzz_dep_.CoverTab[528879]++
//line /snap/go/10455/src/net/ipsock_posix.go:49
		_go_fuzz_dep_.CoverTab[6845]++
//line /snap/go/10455/src/net/ipsock_posix.go:49
		// _ = "end of CoverTab[6845]"
	}
//line /snap/go/10455/src/net/ipsock_posix.go:50
	// _ = "end of CoverTab[6839]"
//line /snap/go/10455/src/net/ipsock_posix.go:50
	_go_fuzz_dep_.CoverTab[6840]++
//line /snap/go/10455/src/net/ipsock_posix.go:50
	_go_fuzz_dep_.CoverTab[786700] = 0
							for i := range probes {
//line /snap/go/10455/src/net/ipsock_posix.go:51
		if _go_fuzz_dep_.CoverTab[786700] == 0 {
//line /snap/go/10455/src/net/ipsock_posix.go:51
			_go_fuzz_dep_.CoverTab[528920]++
//line /snap/go/10455/src/net/ipsock_posix.go:51
		} else {
//line /snap/go/10455/src/net/ipsock_posix.go:51
			_go_fuzz_dep_.CoverTab[528921]++
//line /snap/go/10455/src/net/ipsock_posix.go:51
		}
//line /snap/go/10455/src/net/ipsock_posix.go:51
		_go_fuzz_dep_.CoverTab[786700] = 1
//line /snap/go/10455/src/net/ipsock_posix.go:51
		_go_fuzz_dep_.CoverTab[6846]++
								s, err := sysSocket(syscall.AF_INET6, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
								if err != nil {
//line /snap/go/10455/src/net/ipsock_posix.go:53
			_go_fuzz_dep_.CoverTab[528880]++
//line /snap/go/10455/src/net/ipsock_posix.go:53
			_go_fuzz_dep_.CoverTab[6850]++
									continue
//line /snap/go/10455/src/net/ipsock_posix.go:54
			// _ = "end of CoverTab[6850]"
		} else {
//line /snap/go/10455/src/net/ipsock_posix.go:55
			_go_fuzz_dep_.CoverTab[528881]++
//line /snap/go/10455/src/net/ipsock_posix.go:55
			_go_fuzz_dep_.CoverTab[6851]++
//line /snap/go/10455/src/net/ipsock_posix.go:55
			// _ = "end of CoverTab[6851]"
//line /snap/go/10455/src/net/ipsock_posix.go:55
		}
//line /snap/go/10455/src/net/ipsock_posix.go:55
		// _ = "end of CoverTab[6846]"
//line /snap/go/10455/src/net/ipsock_posix.go:55
		_go_fuzz_dep_.CoverTab[6847]++
								defer poll.CloseFunc(s)
								syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, probes[i].value)
								sa, err := probes[i].laddr.sockaddr(syscall.AF_INET6)
								if err != nil {
//line /snap/go/10455/src/net/ipsock_posix.go:59
			_go_fuzz_dep_.CoverTab[528882]++
//line /snap/go/10455/src/net/ipsock_posix.go:59
			_go_fuzz_dep_.CoverTab[6852]++
									continue
//line /snap/go/10455/src/net/ipsock_posix.go:60
			// _ = "end of CoverTab[6852]"
		} else {
//line /snap/go/10455/src/net/ipsock_posix.go:61
			_go_fuzz_dep_.CoverTab[528883]++
//line /snap/go/10455/src/net/ipsock_posix.go:61
			_go_fuzz_dep_.CoverTab[6853]++
//line /snap/go/10455/src/net/ipsock_posix.go:61
			// _ = "end of CoverTab[6853]"
//line /snap/go/10455/src/net/ipsock_posix.go:61
		}
//line /snap/go/10455/src/net/ipsock_posix.go:61
		// _ = "end of CoverTab[6847]"
//line /snap/go/10455/src/net/ipsock_posix.go:61
		_go_fuzz_dep_.CoverTab[6848]++
								if err := syscall.Bind(s, sa); err != nil {
//line /snap/go/10455/src/net/ipsock_posix.go:62
			_go_fuzz_dep_.CoverTab[528884]++
//line /snap/go/10455/src/net/ipsock_posix.go:62
			_go_fuzz_dep_.CoverTab[6854]++
									continue
//line /snap/go/10455/src/net/ipsock_posix.go:63
			// _ = "end of CoverTab[6854]"
		} else {
//line /snap/go/10455/src/net/ipsock_posix.go:64
			_go_fuzz_dep_.CoverTab[528885]++
//line /snap/go/10455/src/net/ipsock_posix.go:64
			_go_fuzz_dep_.CoverTab[6855]++
//line /snap/go/10455/src/net/ipsock_posix.go:64
			// _ = "end of CoverTab[6855]"
//line /snap/go/10455/src/net/ipsock_posix.go:64
		}
//line /snap/go/10455/src/net/ipsock_posix.go:64
		// _ = "end of CoverTab[6848]"
//line /snap/go/10455/src/net/ipsock_posix.go:64
		_go_fuzz_dep_.CoverTab[6849]++
								if i == 0 {
//line /snap/go/10455/src/net/ipsock_posix.go:65
			_go_fuzz_dep_.CoverTab[528886]++
//line /snap/go/10455/src/net/ipsock_posix.go:65
			_go_fuzz_dep_.CoverTab[6856]++
									p.ipv6Enabled = true
//line /snap/go/10455/src/net/ipsock_posix.go:66
			// _ = "end of CoverTab[6856]"
		} else {
//line /snap/go/10455/src/net/ipsock_posix.go:67
			_go_fuzz_dep_.CoverTab[528887]++
//line /snap/go/10455/src/net/ipsock_posix.go:67
			_go_fuzz_dep_.CoverTab[6857]++
									p.ipv4MappedIPv6Enabled = true
//line /snap/go/10455/src/net/ipsock_posix.go:68
			// _ = "end of CoverTab[6857]"
		}
//line /snap/go/10455/src/net/ipsock_posix.go:69
		// _ = "end of CoverTab[6849]"
	}
//line /snap/go/10455/src/net/ipsock_posix.go:70
	if _go_fuzz_dep_.CoverTab[786700] == 0 {
//line /snap/go/10455/src/net/ipsock_posix.go:70
		_go_fuzz_dep_.CoverTab[528922]++
//line /snap/go/10455/src/net/ipsock_posix.go:70
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:70
		_go_fuzz_dep_.CoverTab[528923]++
//line /snap/go/10455/src/net/ipsock_posix.go:70
	}
//line /snap/go/10455/src/net/ipsock_posix.go:70
	// _ = "end of CoverTab[6840]"
}

// favoriteAddrFamily returns the appropriate address family for the
//line /snap/go/10455/src/net/ipsock_posix.go:73
// given network, laddr, raddr and mode.
//line /snap/go/10455/src/net/ipsock_posix.go:73
//
//line /snap/go/10455/src/net/ipsock_posix.go:73
// If mode indicates "listen" and laddr is a wildcard, we assume that
//line /snap/go/10455/src/net/ipsock_posix.go:73
// the user wants to make a passive-open connection with a wildcard
//line /snap/go/10455/src/net/ipsock_posix.go:73
// address family, both AF_INET and AF_INET6, and a wildcard address
//line /snap/go/10455/src/net/ipsock_posix.go:73
// like the following:
//line /snap/go/10455/src/net/ipsock_posix.go:73
//
//line /snap/go/10455/src/net/ipsock_posix.go:73
//   - A listen for a wildcard communication domain, "tcp" or
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     "udp", with a wildcard address: If the platform supports
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     both IPv6 and IPv4-mapped IPv6 communication capabilities,
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     or does not support IPv4, we use a dual stack, AF_INET6 and
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     IPV6_V6ONLY=0, wildcard address listen. The dual stack
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     wildcard address listen may fall back to an IPv6-only,
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     AF_INET6 and IPV6_V6ONLY=1, wildcard address listen.
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     Otherwise we prefer an IPv4-only, AF_INET, wildcard address
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     listen.
//line /snap/go/10455/src/net/ipsock_posix.go:73
//
//line /snap/go/10455/src/net/ipsock_posix.go:73
//   - A listen for a wildcard communication domain, "tcp" or
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     "udp", with an IPv4 wildcard address: same as above.
//line /snap/go/10455/src/net/ipsock_posix.go:73
//
//line /snap/go/10455/src/net/ipsock_posix.go:73
//   - A listen for a wildcard communication domain, "tcp" or
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     "udp", with an IPv6 wildcard address: same as above.
//line /snap/go/10455/src/net/ipsock_posix.go:73
//
//line /snap/go/10455/src/net/ipsock_posix.go:73
//   - A listen for an IPv4 communication domain, "tcp4" or "udp4",
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     with an IPv4 wildcard address: We use an IPv4-only, AF_INET,
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     wildcard address listen.
//line /snap/go/10455/src/net/ipsock_posix.go:73
//
//line /snap/go/10455/src/net/ipsock_posix.go:73
//   - A listen for an IPv6 communication domain, "tcp6" or "udp6",
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     with an IPv6 wildcard address: We use an IPv6-only, AF_INET6
//line /snap/go/10455/src/net/ipsock_posix.go:73
//     and IPV6_V6ONLY=1, wildcard address listen.
//line /snap/go/10455/src/net/ipsock_posix.go:73
//
//line /snap/go/10455/src/net/ipsock_posix.go:73
// Otherwise guess: If the addresses are IPv4 then returns AF_INET,
//line /snap/go/10455/src/net/ipsock_posix.go:73
// or else returns AF_INET6. It also returns a boolean value what
//line /snap/go/10455/src/net/ipsock_posix.go:73
// designates IPV6_V6ONLY option.
//line /snap/go/10455/src/net/ipsock_posix.go:73
//
//line /snap/go/10455/src/net/ipsock_posix.go:73
// Note that the latest DragonFly BSD and OpenBSD kernels allow
//line /snap/go/10455/src/net/ipsock_posix.go:73
// neither "net.inet6.ip6.v6only=1" change nor IPPROTO_IPV6 level
//line /snap/go/10455/src/net/ipsock_posix.go:73
// IPV6_V6ONLY socket option setting.
//line /snap/go/10455/src/net/ipsock_posix.go:112
func favoriteAddrFamily(network string, laddr, raddr sockaddr, mode string) (family int, ipv6only bool) {
//line /snap/go/10455/src/net/ipsock_posix.go:112
	_go_fuzz_dep_.CoverTab[6858]++
							switch network[len(network)-1] {
	case '4':
//line /snap/go/10455/src/net/ipsock_posix.go:114
		_go_fuzz_dep_.CoverTab[528888]++
//line /snap/go/10455/src/net/ipsock_posix.go:114
		_go_fuzz_dep_.CoverTab[6862]++
								return syscall.AF_INET, false
//line /snap/go/10455/src/net/ipsock_posix.go:115
		// _ = "end of CoverTab[6862]"
	case '6':
//line /snap/go/10455/src/net/ipsock_posix.go:116
		_go_fuzz_dep_.CoverTab[528889]++
//line /snap/go/10455/src/net/ipsock_posix.go:116
		_go_fuzz_dep_.CoverTab[6863]++
								return syscall.AF_INET6, true
//line /snap/go/10455/src/net/ipsock_posix.go:117
		// _ = "end of CoverTab[6863]"
//line /snap/go/10455/src/net/ipsock_posix.go:117
	default:
//line /snap/go/10455/src/net/ipsock_posix.go:117
		_go_fuzz_dep_.CoverTab[528890]++
//line /snap/go/10455/src/net/ipsock_posix.go:117
		_go_fuzz_dep_.CoverTab[6864]++
//line /snap/go/10455/src/net/ipsock_posix.go:117
		// _ = "end of CoverTab[6864]"
	}
//line /snap/go/10455/src/net/ipsock_posix.go:118
	// _ = "end of CoverTab[6858]"
//line /snap/go/10455/src/net/ipsock_posix.go:118
	_go_fuzz_dep_.CoverTab[6859]++

							if mode == "listen" && func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:120
		_go_fuzz_dep_.CoverTab[6865]++
//line /snap/go/10455/src/net/ipsock_posix.go:120
		return (laddr == nil || func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:120
			_go_fuzz_dep_.CoverTab[6866]++
//line /snap/go/10455/src/net/ipsock_posix.go:120
			return laddr.isWildcard()
//line /snap/go/10455/src/net/ipsock_posix.go:120
			// _ = "end of CoverTab[6866]"
//line /snap/go/10455/src/net/ipsock_posix.go:120
		}())
//line /snap/go/10455/src/net/ipsock_posix.go:120
		// _ = "end of CoverTab[6865]"
//line /snap/go/10455/src/net/ipsock_posix.go:120
	}() {
//line /snap/go/10455/src/net/ipsock_posix.go:120
		_go_fuzz_dep_.CoverTab[528891]++
//line /snap/go/10455/src/net/ipsock_posix.go:120
		_go_fuzz_dep_.CoverTab[6867]++
								if supportsIPv4map() || func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:121
			_go_fuzz_dep_.CoverTab[6870]++
//line /snap/go/10455/src/net/ipsock_posix.go:121
			return !supportsIPv4()
//line /snap/go/10455/src/net/ipsock_posix.go:121
			// _ = "end of CoverTab[6870]"
//line /snap/go/10455/src/net/ipsock_posix.go:121
		}() {
//line /snap/go/10455/src/net/ipsock_posix.go:121
			_go_fuzz_dep_.CoverTab[528893]++
//line /snap/go/10455/src/net/ipsock_posix.go:121
			_go_fuzz_dep_.CoverTab[6871]++
									return syscall.AF_INET6, false
//line /snap/go/10455/src/net/ipsock_posix.go:122
			// _ = "end of CoverTab[6871]"
		} else {
//line /snap/go/10455/src/net/ipsock_posix.go:123
			_go_fuzz_dep_.CoverTab[528894]++
//line /snap/go/10455/src/net/ipsock_posix.go:123
			_go_fuzz_dep_.CoverTab[6872]++
//line /snap/go/10455/src/net/ipsock_posix.go:123
			// _ = "end of CoverTab[6872]"
//line /snap/go/10455/src/net/ipsock_posix.go:123
		}
//line /snap/go/10455/src/net/ipsock_posix.go:123
		// _ = "end of CoverTab[6867]"
//line /snap/go/10455/src/net/ipsock_posix.go:123
		_go_fuzz_dep_.CoverTab[6868]++
								if laddr == nil {
//line /snap/go/10455/src/net/ipsock_posix.go:124
			_go_fuzz_dep_.CoverTab[528895]++
//line /snap/go/10455/src/net/ipsock_posix.go:124
			_go_fuzz_dep_.CoverTab[6873]++
									return syscall.AF_INET, false
//line /snap/go/10455/src/net/ipsock_posix.go:125
			// _ = "end of CoverTab[6873]"
		} else {
//line /snap/go/10455/src/net/ipsock_posix.go:126
			_go_fuzz_dep_.CoverTab[528896]++
//line /snap/go/10455/src/net/ipsock_posix.go:126
			_go_fuzz_dep_.CoverTab[6874]++
//line /snap/go/10455/src/net/ipsock_posix.go:126
			// _ = "end of CoverTab[6874]"
//line /snap/go/10455/src/net/ipsock_posix.go:126
		}
//line /snap/go/10455/src/net/ipsock_posix.go:126
		// _ = "end of CoverTab[6868]"
//line /snap/go/10455/src/net/ipsock_posix.go:126
		_go_fuzz_dep_.CoverTab[6869]++
								return laddr.family(), false
//line /snap/go/10455/src/net/ipsock_posix.go:127
		// _ = "end of CoverTab[6869]"
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:128
		_go_fuzz_dep_.CoverTab[528892]++
//line /snap/go/10455/src/net/ipsock_posix.go:128
		_go_fuzz_dep_.CoverTab[6875]++
//line /snap/go/10455/src/net/ipsock_posix.go:128
		// _ = "end of CoverTab[6875]"
//line /snap/go/10455/src/net/ipsock_posix.go:128
	}
//line /snap/go/10455/src/net/ipsock_posix.go:128
	// _ = "end of CoverTab[6859]"
//line /snap/go/10455/src/net/ipsock_posix.go:128
	_go_fuzz_dep_.CoverTab[6860]++

							if (laddr == nil || func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:130
		_go_fuzz_dep_.CoverTab[6876]++
//line /snap/go/10455/src/net/ipsock_posix.go:130
		return laddr.family() == syscall.AF_INET
//line /snap/go/10455/src/net/ipsock_posix.go:130
		// _ = "end of CoverTab[6876]"
//line /snap/go/10455/src/net/ipsock_posix.go:130
	}()) && func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:130
		_go_fuzz_dep_.CoverTab[6877]++
//line /snap/go/10455/src/net/ipsock_posix.go:130
		return (raddr == nil || func() bool {
									_go_fuzz_dep_.CoverTab[6878]++
//line /snap/go/10455/src/net/ipsock_posix.go:131
			return raddr.family() == syscall.AF_INET
//line /snap/go/10455/src/net/ipsock_posix.go:131
			// _ = "end of CoverTab[6878]"
//line /snap/go/10455/src/net/ipsock_posix.go:131
		}())
//line /snap/go/10455/src/net/ipsock_posix.go:131
		// _ = "end of CoverTab[6877]"
//line /snap/go/10455/src/net/ipsock_posix.go:131
	}() {
//line /snap/go/10455/src/net/ipsock_posix.go:131
		_go_fuzz_dep_.CoverTab[528897]++
//line /snap/go/10455/src/net/ipsock_posix.go:131
		_go_fuzz_dep_.CoverTab[6879]++
								return syscall.AF_INET, false
//line /snap/go/10455/src/net/ipsock_posix.go:132
		// _ = "end of CoverTab[6879]"
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:133
		_go_fuzz_dep_.CoverTab[528898]++
//line /snap/go/10455/src/net/ipsock_posix.go:133
		_go_fuzz_dep_.CoverTab[6880]++
//line /snap/go/10455/src/net/ipsock_posix.go:133
		// _ = "end of CoverTab[6880]"
//line /snap/go/10455/src/net/ipsock_posix.go:133
	}
//line /snap/go/10455/src/net/ipsock_posix.go:133
	// _ = "end of CoverTab[6860]"
//line /snap/go/10455/src/net/ipsock_posix.go:133
	_go_fuzz_dep_.CoverTab[6861]++
							return syscall.AF_INET6, false
//line /snap/go/10455/src/net/ipsock_posix.go:134
	// _ = "end of CoverTab[6861]"
}

func internetSocket(ctx context.Context, net string, laddr, raddr sockaddr, sotype, proto int, mode string, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) (fd *netFD, err error) {
//line /snap/go/10455/src/net/ipsock_posix.go:137
	_go_fuzz_dep_.CoverTab[6881]++
							if (runtime.GOOS == "aix" || func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6883]++
//line /snap/go/10455/src/net/ipsock_posix.go:138
		return runtime.GOOS == "windows"
//line /snap/go/10455/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[6883]"
//line /snap/go/10455/src/net/ipsock_posix.go:138
	}() || func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6884]++
//line /snap/go/10455/src/net/ipsock_posix.go:138
		return runtime.GOOS == "openbsd"
//line /snap/go/10455/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[6884]"
//line /snap/go/10455/src/net/ipsock_posix.go:138
	}()) && func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6885]++
//line /snap/go/10455/src/net/ipsock_posix.go:138
		return mode == "dial"
//line /snap/go/10455/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[6885]"
//line /snap/go/10455/src/net/ipsock_posix.go:138
	}() && func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6886]++
//line /snap/go/10455/src/net/ipsock_posix.go:138
		return raddr.isWildcard()
//line /snap/go/10455/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[6886]"
//line /snap/go/10455/src/net/ipsock_posix.go:138
	}() {
//line /snap/go/10455/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[528899]++
//line /snap/go/10455/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6887]++
								raddr = raddr.toLocal(net)
//line /snap/go/10455/src/net/ipsock_posix.go:139
		// _ = "end of CoverTab[6887]"
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:140
		_go_fuzz_dep_.CoverTab[528900]++
//line /snap/go/10455/src/net/ipsock_posix.go:140
		_go_fuzz_dep_.CoverTab[6888]++
//line /snap/go/10455/src/net/ipsock_posix.go:140
		// _ = "end of CoverTab[6888]"
//line /snap/go/10455/src/net/ipsock_posix.go:140
	}
//line /snap/go/10455/src/net/ipsock_posix.go:140
	// _ = "end of CoverTab[6881]"
//line /snap/go/10455/src/net/ipsock_posix.go:140
	_go_fuzz_dep_.CoverTab[6882]++
							family, ipv6only := favoriteAddrFamily(net, laddr, raddr, mode)
							return socket(ctx, net, family, sotype, proto, ipv6only, laddr, raddr, ctrlCtxFn)
//line /snap/go/10455/src/net/ipsock_posix.go:142
	// _ = "end of CoverTab[6882]"
}

func ipToSockaddrInet4(ip IP, port int) (syscall.SockaddrInet4, error) {
//line /snap/go/10455/src/net/ipsock_posix.go:145
	_go_fuzz_dep_.CoverTab[6889]++
							if len(ip) == 0 {
//line /snap/go/10455/src/net/ipsock_posix.go:146
		_go_fuzz_dep_.CoverTab[528901]++
//line /snap/go/10455/src/net/ipsock_posix.go:146
		_go_fuzz_dep_.CoverTab[6892]++
								ip = IPv4zero
//line /snap/go/10455/src/net/ipsock_posix.go:147
		// _ = "end of CoverTab[6892]"
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:148
		_go_fuzz_dep_.CoverTab[528902]++
//line /snap/go/10455/src/net/ipsock_posix.go:148
		_go_fuzz_dep_.CoverTab[6893]++
//line /snap/go/10455/src/net/ipsock_posix.go:148
		// _ = "end of CoverTab[6893]"
//line /snap/go/10455/src/net/ipsock_posix.go:148
	}
//line /snap/go/10455/src/net/ipsock_posix.go:148
	// _ = "end of CoverTab[6889]"
//line /snap/go/10455/src/net/ipsock_posix.go:148
	_go_fuzz_dep_.CoverTab[6890]++
							ip4 := ip.To4()
							if ip4 == nil {
//line /snap/go/10455/src/net/ipsock_posix.go:150
		_go_fuzz_dep_.CoverTab[528903]++
//line /snap/go/10455/src/net/ipsock_posix.go:150
		_go_fuzz_dep_.CoverTab[6894]++
								return syscall.SockaddrInet4{}, &AddrError{Err: "non-IPv4 address", Addr: ip.String()}
//line /snap/go/10455/src/net/ipsock_posix.go:151
		// _ = "end of CoverTab[6894]"
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:152
		_go_fuzz_dep_.CoverTab[528904]++
//line /snap/go/10455/src/net/ipsock_posix.go:152
		_go_fuzz_dep_.CoverTab[6895]++
//line /snap/go/10455/src/net/ipsock_posix.go:152
		// _ = "end of CoverTab[6895]"
//line /snap/go/10455/src/net/ipsock_posix.go:152
	}
//line /snap/go/10455/src/net/ipsock_posix.go:152
	// _ = "end of CoverTab[6890]"
//line /snap/go/10455/src/net/ipsock_posix.go:152
	_go_fuzz_dep_.CoverTab[6891]++
							sa := syscall.SockaddrInet4{Port: port}
							copy(sa.Addr[:], ip4)
							return sa, nil
//line /snap/go/10455/src/net/ipsock_posix.go:155
	// _ = "end of CoverTab[6891]"
}

func ipToSockaddrInet6(ip IP, port int, zone string) (syscall.SockaddrInet6, error) {
//line /snap/go/10455/src/net/ipsock_posix.go:158
	_go_fuzz_dep_.CoverTab[6896]++

//line /snap/go/10455/src/net/ipsock_posix.go:169
	if len(ip) == 0 || func() bool {
//line /snap/go/10455/src/net/ipsock_posix.go:169
		_go_fuzz_dep_.CoverTab[6899]++
//line /snap/go/10455/src/net/ipsock_posix.go:169
		return ip.Equal(IPv4zero)
//line /snap/go/10455/src/net/ipsock_posix.go:169
		// _ = "end of CoverTab[6899]"
//line /snap/go/10455/src/net/ipsock_posix.go:169
	}() {
//line /snap/go/10455/src/net/ipsock_posix.go:169
		_go_fuzz_dep_.CoverTab[528905]++
//line /snap/go/10455/src/net/ipsock_posix.go:169
		_go_fuzz_dep_.CoverTab[6900]++
								ip = IPv6zero
//line /snap/go/10455/src/net/ipsock_posix.go:170
		// _ = "end of CoverTab[6900]"
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:171
		_go_fuzz_dep_.CoverTab[528906]++
//line /snap/go/10455/src/net/ipsock_posix.go:171
		_go_fuzz_dep_.CoverTab[6901]++
//line /snap/go/10455/src/net/ipsock_posix.go:171
		// _ = "end of CoverTab[6901]"
//line /snap/go/10455/src/net/ipsock_posix.go:171
	}
//line /snap/go/10455/src/net/ipsock_posix.go:171
	// _ = "end of CoverTab[6896]"
//line /snap/go/10455/src/net/ipsock_posix.go:171
	_go_fuzz_dep_.CoverTab[6897]++

//line /snap/go/10455/src/net/ipsock_posix.go:174
	ip6 := ip.To16()
	if ip6 == nil {
//line /snap/go/10455/src/net/ipsock_posix.go:175
		_go_fuzz_dep_.CoverTab[528907]++
//line /snap/go/10455/src/net/ipsock_posix.go:175
		_go_fuzz_dep_.CoverTab[6902]++
								return syscall.SockaddrInet6{}, &AddrError{Err: "non-IPv6 address", Addr: ip.String()}
//line /snap/go/10455/src/net/ipsock_posix.go:176
		// _ = "end of CoverTab[6902]"
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:177
		_go_fuzz_dep_.CoverTab[528908]++
//line /snap/go/10455/src/net/ipsock_posix.go:177
		_go_fuzz_dep_.CoverTab[6903]++
//line /snap/go/10455/src/net/ipsock_posix.go:177
		// _ = "end of CoverTab[6903]"
//line /snap/go/10455/src/net/ipsock_posix.go:177
	}
//line /snap/go/10455/src/net/ipsock_posix.go:177
	// _ = "end of CoverTab[6897]"
//line /snap/go/10455/src/net/ipsock_posix.go:177
	_go_fuzz_dep_.CoverTab[6898]++
							sa := syscall.SockaddrInet6{Port: port, ZoneId: uint32(zoneCache.index(zone))}
							copy(sa.Addr[:], ip6)
							return sa, nil
//line /snap/go/10455/src/net/ipsock_posix.go:180
	// _ = "end of CoverTab[6898]"
}

func ipToSockaddr(family int, ip IP, port int, zone string) (syscall.Sockaddr, error) {
//line /snap/go/10455/src/net/ipsock_posix.go:183
	_go_fuzz_dep_.CoverTab[6904]++
							switch family {
	case syscall.AF_INET:
//line /snap/go/10455/src/net/ipsock_posix.go:185
		_go_fuzz_dep_.CoverTab[528909]++
//line /snap/go/10455/src/net/ipsock_posix.go:185
		_go_fuzz_dep_.CoverTab[6906]++
								sa, err := ipToSockaddrInet4(ip, port)
								if err != nil {
//line /snap/go/10455/src/net/ipsock_posix.go:187
			_go_fuzz_dep_.CoverTab[528912]++
//line /snap/go/10455/src/net/ipsock_posix.go:187
			_go_fuzz_dep_.CoverTab[6911]++
									return nil, err
//line /snap/go/10455/src/net/ipsock_posix.go:188
			// _ = "end of CoverTab[6911]"
		} else {
//line /snap/go/10455/src/net/ipsock_posix.go:189
			_go_fuzz_dep_.CoverTab[528913]++
//line /snap/go/10455/src/net/ipsock_posix.go:189
			_go_fuzz_dep_.CoverTab[6912]++
//line /snap/go/10455/src/net/ipsock_posix.go:189
			// _ = "end of CoverTab[6912]"
//line /snap/go/10455/src/net/ipsock_posix.go:189
		}
//line /snap/go/10455/src/net/ipsock_posix.go:189
		// _ = "end of CoverTab[6906]"
//line /snap/go/10455/src/net/ipsock_posix.go:189
		_go_fuzz_dep_.CoverTab[6907]++
								return &sa, nil
//line /snap/go/10455/src/net/ipsock_posix.go:190
		// _ = "end of CoverTab[6907]"
	case syscall.AF_INET6:
//line /snap/go/10455/src/net/ipsock_posix.go:191
		_go_fuzz_dep_.CoverTab[528910]++
//line /snap/go/10455/src/net/ipsock_posix.go:191
		_go_fuzz_dep_.CoverTab[6908]++
								sa, err := ipToSockaddrInet6(ip, port, zone)
								if err != nil {
//line /snap/go/10455/src/net/ipsock_posix.go:193
			_go_fuzz_dep_.CoverTab[528914]++
//line /snap/go/10455/src/net/ipsock_posix.go:193
			_go_fuzz_dep_.CoverTab[6913]++
									return nil, err
//line /snap/go/10455/src/net/ipsock_posix.go:194
			// _ = "end of CoverTab[6913]"
		} else {
//line /snap/go/10455/src/net/ipsock_posix.go:195
			_go_fuzz_dep_.CoverTab[528915]++
//line /snap/go/10455/src/net/ipsock_posix.go:195
			_go_fuzz_dep_.CoverTab[6914]++
//line /snap/go/10455/src/net/ipsock_posix.go:195
			// _ = "end of CoverTab[6914]"
//line /snap/go/10455/src/net/ipsock_posix.go:195
		}
//line /snap/go/10455/src/net/ipsock_posix.go:195
		// _ = "end of CoverTab[6908]"
//line /snap/go/10455/src/net/ipsock_posix.go:195
		_go_fuzz_dep_.CoverTab[6909]++
								return &sa, nil
//line /snap/go/10455/src/net/ipsock_posix.go:196
		// _ = "end of CoverTab[6909]"
//line /snap/go/10455/src/net/ipsock_posix.go:196
	default:
//line /snap/go/10455/src/net/ipsock_posix.go:196
		_go_fuzz_dep_.CoverTab[528911]++
//line /snap/go/10455/src/net/ipsock_posix.go:196
		_go_fuzz_dep_.CoverTab[6910]++
//line /snap/go/10455/src/net/ipsock_posix.go:196
		// _ = "end of CoverTab[6910]"
	}
//line /snap/go/10455/src/net/ipsock_posix.go:197
	// _ = "end of CoverTab[6904]"
//line /snap/go/10455/src/net/ipsock_posix.go:197
	_go_fuzz_dep_.CoverTab[6905]++
							return nil, &AddrError{Err: "invalid address family", Addr: ip.String()}
//line /snap/go/10455/src/net/ipsock_posix.go:198
	// _ = "end of CoverTab[6905]"
}

func addrPortToSockaddrInet4(ap netip.AddrPort) (syscall.SockaddrInet4, error) {
//line /snap/go/10455/src/net/ipsock_posix.go:201
	_go_fuzz_dep_.CoverTab[6915]++

//line /snap/go/10455/src/net/ipsock_posix.go:204
	addr := ap.Addr()
	if !addr.Is4() {
//line /snap/go/10455/src/net/ipsock_posix.go:205
		_go_fuzz_dep_.CoverTab[528916]++
//line /snap/go/10455/src/net/ipsock_posix.go:205
		_go_fuzz_dep_.CoverTab[6917]++
								return syscall.SockaddrInet4{}, &AddrError{Err: "non-IPv4 address", Addr: addr.String()}
//line /snap/go/10455/src/net/ipsock_posix.go:206
		// _ = "end of CoverTab[6917]"
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:207
		_go_fuzz_dep_.CoverTab[528917]++
//line /snap/go/10455/src/net/ipsock_posix.go:207
		_go_fuzz_dep_.CoverTab[6918]++
//line /snap/go/10455/src/net/ipsock_posix.go:207
		// _ = "end of CoverTab[6918]"
//line /snap/go/10455/src/net/ipsock_posix.go:207
	}
//line /snap/go/10455/src/net/ipsock_posix.go:207
	// _ = "end of CoverTab[6915]"
//line /snap/go/10455/src/net/ipsock_posix.go:207
	_go_fuzz_dep_.CoverTab[6916]++
							sa := syscall.SockaddrInet4{
		Addr:	addr.As4(),
		Port:	int(ap.Port()),
	}
							return sa, nil
//line /snap/go/10455/src/net/ipsock_posix.go:212
	// _ = "end of CoverTab[6916]"
}

func addrPortToSockaddrInet6(ap netip.AddrPort) (syscall.SockaddrInet6, error) {
//line /snap/go/10455/src/net/ipsock_posix.go:215
	_go_fuzz_dep_.CoverTab[6919]++

//line /snap/go/10455/src/net/ipsock_posix.go:222
	addr := ap.Addr()
	if !addr.IsValid() {
//line /snap/go/10455/src/net/ipsock_posix.go:223
		_go_fuzz_dep_.CoverTab[528918]++
//line /snap/go/10455/src/net/ipsock_posix.go:223
		_go_fuzz_dep_.CoverTab[6921]++
								return syscall.SockaddrInet6{}, &AddrError{Err: "non-IPv6 address", Addr: addr.String()}
//line /snap/go/10455/src/net/ipsock_posix.go:224
		// _ = "end of CoverTab[6921]"
	} else {
//line /snap/go/10455/src/net/ipsock_posix.go:225
		_go_fuzz_dep_.CoverTab[528919]++
//line /snap/go/10455/src/net/ipsock_posix.go:225
		_go_fuzz_dep_.CoverTab[6922]++
//line /snap/go/10455/src/net/ipsock_posix.go:225
		// _ = "end of CoverTab[6922]"
//line /snap/go/10455/src/net/ipsock_posix.go:225
	}
//line /snap/go/10455/src/net/ipsock_posix.go:225
	// _ = "end of CoverTab[6919]"
//line /snap/go/10455/src/net/ipsock_posix.go:225
	_go_fuzz_dep_.CoverTab[6920]++
							sa := syscall.SockaddrInet6{
		Addr:	addr.As16(),
		Port:	int(ap.Port()),
		ZoneId:	uint32(zoneCache.index(addr.Zone())),
	}
							return sa, nil
//line /snap/go/10455/src/net/ipsock_posix.go:231
	// _ = "end of CoverTab[6920]"
}

//line /snap/go/10455/src/net/ipsock_posix.go:232
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/ipsock_posix.go:232
var _ = _go_fuzz_dep_.CoverTab
