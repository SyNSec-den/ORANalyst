// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || windows

//line /usr/local/go/src/net/ipsock_posix.go:7
package net

//line /usr/local/go/src/net/ipsock_posix.go:7
import (
//line /usr/local/go/src/net/ipsock_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/ipsock_posix.go:7
)
//line /usr/local/go/src/net/ipsock_posix.go:7
import (
//line /usr/local/go/src/net/ipsock_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/ipsock_posix.go:7
)

import (
	"context"
	"internal/poll"
	"net/netip"
	"runtime"
	"syscall"
)

// probe probes IPv4, IPv6 and IPv4-mapped IPv6 communication
//line /usr/local/go/src/net/ipsock_posix.go:17
// capabilities which are controlled by the IPV6_V6ONLY socket option
//line /usr/local/go/src/net/ipsock_posix.go:17
// and kernel configuration.
//line /usr/local/go/src/net/ipsock_posix.go:17
//
//line /usr/local/go/src/net/ipsock_posix.go:17
// Should we try to use the IPv4 socket interface if we're only
//line /usr/local/go/src/net/ipsock_posix.go:17
// dealing with IPv4 sockets? As long as the host system understands
//line /usr/local/go/src/net/ipsock_posix.go:17
// IPv4-mapped IPv6, it's okay to pass IPv4-mapped IPv6 addresses to
//line /usr/local/go/src/net/ipsock_posix.go:17
// the IPv6 interface. That simplifies our code and is most
//line /usr/local/go/src/net/ipsock_posix.go:17
// general. Unfortunately, we need to run on kernels built without
//line /usr/local/go/src/net/ipsock_posix.go:17
// IPv6 support too. So probe the kernel to figure it out.
//line /usr/local/go/src/net/ipsock_posix.go:27
func (p *ipStackCapabilities) probe() {
//line /usr/local/go/src/net/ipsock_posix.go:27
	_go_fuzz_dep_.CoverTab[14962]++
							s, err := sysSocket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
							switch err {
	case syscall.EAFNOSUPPORT, syscall.EPROTONOSUPPORT:
//line /usr/local/go/src/net/ipsock_posix.go:30
		_go_fuzz_dep_.CoverTab[14965]++
//line /usr/local/go/src/net/ipsock_posix.go:30
		// _ = "end of CoverTab[14965]"
	case nil:
//line /usr/local/go/src/net/ipsock_posix.go:31
		_go_fuzz_dep_.CoverTab[14966]++
								poll.CloseFunc(s)
								p.ipv4Enabled = true
//line /usr/local/go/src/net/ipsock_posix.go:33
		// _ = "end of CoverTab[14966]"
//line /usr/local/go/src/net/ipsock_posix.go:33
	default:
//line /usr/local/go/src/net/ipsock_posix.go:33
		_go_fuzz_dep_.CoverTab[14967]++
//line /usr/local/go/src/net/ipsock_posix.go:33
		// _ = "end of CoverTab[14967]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:34
	// _ = "end of CoverTab[14962]"
//line /usr/local/go/src/net/ipsock_posix.go:34
	_go_fuzz_dep_.CoverTab[14963]++
							var probes = []struct {
		laddr	TCPAddr
		value	int
	}{

		{laddr: TCPAddr{IP: ParseIP("::1")}, value: 1},

		{laddr: TCPAddr{IP: IPv4(127, 0, 0, 1)}, value: 0},
	}
	switch runtime.GOOS {
	case "dragonfly", "openbsd":
//line /usr/local/go/src/net/ipsock_posix.go:45
		_go_fuzz_dep_.CoverTab[14968]++

//line /usr/local/go/src/net/ipsock_posix.go:49
		probes = probes[:1]
//line /usr/local/go/src/net/ipsock_posix.go:49
		// _ = "end of CoverTab[14968]"
//line /usr/local/go/src/net/ipsock_posix.go:49
	default:
//line /usr/local/go/src/net/ipsock_posix.go:49
		_go_fuzz_dep_.CoverTab[14969]++
//line /usr/local/go/src/net/ipsock_posix.go:49
		// _ = "end of CoverTab[14969]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:50
	// _ = "end of CoverTab[14963]"
//line /usr/local/go/src/net/ipsock_posix.go:50
	_go_fuzz_dep_.CoverTab[14964]++
							for i := range probes {
//line /usr/local/go/src/net/ipsock_posix.go:51
		_go_fuzz_dep_.CoverTab[14970]++
								s, err := sysSocket(syscall.AF_INET6, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
								if err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:53
			_go_fuzz_dep_.CoverTab[14974]++
									continue
//line /usr/local/go/src/net/ipsock_posix.go:54
			// _ = "end of CoverTab[14974]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:55
			_go_fuzz_dep_.CoverTab[14975]++
//line /usr/local/go/src/net/ipsock_posix.go:55
			// _ = "end of CoverTab[14975]"
//line /usr/local/go/src/net/ipsock_posix.go:55
		}
//line /usr/local/go/src/net/ipsock_posix.go:55
		// _ = "end of CoverTab[14970]"
//line /usr/local/go/src/net/ipsock_posix.go:55
		_go_fuzz_dep_.CoverTab[14971]++
								defer poll.CloseFunc(s)
								syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, probes[i].value)
								sa, err := probes[i].laddr.sockaddr(syscall.AF_INET6)
								if err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:59
			_go_fuzz_dep_.CoverTab[14976]++
									continue
//line /usr/local/go/src/net/ipsock_posix.go:60
			// _ = "end of CoverTab[14976]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:61
			_go_fuzz_dep_.CoverTab[14977]++
//line /usr/local/go/src/net/ipsock_posix.go:61
			// _ = "end of CoverTab[14977]"
//line /usr/local/go/src/net/ipsock_posix.go:61
		}
//line /usr/local/go/src/net/ipsock_posix.go:61
		// _ = "end of CoverTab[14971]"
//line /usr/local/go/src/net/ipsock_posix.go:61
		_go_fuzz_dep_.CoverTab[14972]++
								if err := syscall.Bind(s, sa); err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:62
			_go_fuzz_dep_.CoverTab[14978]++
									continue
//line /usr/local/go/src/net/ipsock_posix.go:63
			// _ = "end of CoverTab[14978]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:64
			_go_fuzz_dep_.CoverTab[14979]++
//line /usr/local/go/src/net/ipsock_posix.go:64
			// _ = "end of CoverTab[14979]"
//line /usr/local/go/src/net/ipsock_posix.go:64
		}
//line /usr/local/go/src/net/ipsock_posix.go:64
		// _ = "end of CoverTab[14972]"
//line /usr/local/go/src/net/ipsock_posix.go:64
		_go_fuzz_dep_.CoverTab[14973]++
								if i == 0 {
//line /usr/local/go/src/net/ipsock_posix.go:65
			_go_fuzz_dep_.CoverTab[14980]++
									p.ipv6Enabled = true
//line /usr/local/go/src/net/ipsock_posix.go:66
			// _ = "end of CoverTab[14980]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:67
			_go_fuzz_dep_.CoverTab[14981]++
									p.ipv4MappedIPv6Enabled = true
//line /usr/local/go/src/net/ipsock_posix.go:68
			// _ = "end of CoverTab[14981]"
		}
//line /usr/local/go/src/net/ipsock_posix.go:69
		// _ = "end of CoverTab[14973]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:70
	// _ = "end of CoverTab[14964]"
}

// favoriteAddrFamily returns the appropriate address family for the
//line /usr/local/go/src/net/ipsock_posix.go:73
// given network, laddr, raddr and mode.
//line /usr/local/go/src/net/ipsock_posix.go:73
//
//line /usr/local/go/src/net/ipsock_posix.go:73
// If mode indicates "listen" and laddr is a wildcard, we assume that
//line /usr/local/go/src/net/ipsock_posix.go:73
// the user wants to make a passive-open connection with a wildcard
//line /usr/local/go/src/net/ipsock_posix.go:73
// address family, both AF_INET and AF_INET6, and a wildcard address
//line /usr/local/go/src/net/ipsock_posix.go:73
// like the following:
//line /usr/local/go/src/net/ipsock_posix.go:73
//
//line /usr/local/go/src/net/ipsock_posix.go:73
//   - A listen for a wildcard communication domain, "tcp" or
//line /usr/local/go/src/net/ipsock_posix.go:73
//     "udp", with a wildcard address: If the platform supports
//line /usr/local/go/src/net/ipsock_posix.go:73
//     both IPv6 and IPv4-mapped IPv6 communication capabilities,
//line /usr/local/go/src/net/ipsock_posix.go:73
//     or does not support IPv4, we use a dual stack, AF_INET6 and
//line /usr/local/go/src/net/ipsock_posix.go:73
//     IPV6_V6ONLY=0, wildcard address listen. The dual stack
//line /usr/local/go/src/net/ipsock_posix.go:73
//     wildcard address listen may fall back to an IPv6-only,
//line /usr/local/go/src/net/ipsock_posix.go:73
//     AF_INET6 and IPV6_V6ONLY=1, wildcard address listen.
//line /usr/local/go/src/net/ipsock_posix.go:73
//     Otherwise we prefer an IPv4-only, AF_INET, wildcard address
//line /usr/local/go/src/net/ipsock_posix.go:73
//     listen.
//line /usr/local/go/src/net/ipsock_posix.go:73
//
//line /usr/local/go/src/net/ipsock_posix.go:73
//   - A listen for a wildcard communication domain, "tcp" or
//line /usr/local/go/src/net/ipsock_posix.go:73
//     "udp", with an IPv4 wildcard address: same as above.
//line /usr/local/go/src/net/ipsock_posix.go:73
//
//line /usr/local/go/src/net/ipsock_posix.go:73
//   - A listen for a wildcard communication domain, "tcp" or
//line /usr/local/go/src/net/ipsock_posix.go:73
//     "udp", with an IPv6 wildcard address: same as above.
//line /usr/local/go/src/net/ipsock_posix.go:73
//
//line /usr/local/go/src/net/ipsock_posix.go:73
//   - A listen for an IPv4 communication domain, "tcp4" or "udp4",
//line /usr/local/go/src/net/ipsock_posix.go:73
//     with an IPv4 wildcard address: We use an IPv4-only, AF_INET,
//line /usr/local/go/src/net/ipsock_posix.go:73
//     wildcard address listen.
//line /usr/local/go/src/net/ipsock_posix.go:73
//
//line /usr/local/go/src/net/ipsock_posix.go:73
//   - A listen for an IPv6 communication domain, "tcp6" or "udp6",
//line /usr/local/go/src/net/ipsock_posix.go:73
//     with an IPv6 wildcard address: We use an IPv6-only, AF_INET6
//line /usr/local/go/src/net/ipsock_posix.go:73
//     and IPV6_V6ONLY=1, wildcard address listen.
//line /usr/local/go/src/net/ipsock_posix.go:73
//
//line /usr/local/go/src/net/ipsock_posix.go:73
// Otherwise guess: If the addresses are IPv4 then returns AF_INET,
//line /usr/local/go/src/net/ipsock_posix.go:73
// or else returns AF_INET6. It also returns a boolean value what
//line /usr/local/go/src/net/ipsock_posix.go:73
// designates IPV6_V6ONLY option.
//line /usr/local/go/src/net/ipsock_posix.go:73
//
//line /usr/local/go/src/net/ipsock_posix.go:73
// Note that the latest DragonFly BSD and OpenBSD kernels allow
//line /usr/local/go/src/net/ipsock_posix.go:73
// neither "net.inet6.ip6.v6only=1" change nor IPPROTO_IPV6 level
//line /usr/local/go/src/net/ipsock_posix.go:73
// IPV6_V6ONLY socket option setting.
//line /usr/local/go/src/net/ipsock_posix.go:112
func favoriteAddrFamily(network string, laddr, raddr sockaddr, mode string) (family int, ipv6only bool) {
//line /usr/local/go/src/net/ipsock_posix.go:112
	_go_fuzz_dep_.CoverTab[14982]++
							switch network[len(network)-1] {
	case '4':
//line /usr/local/go/src/net/ipsock_posix.go:114
		_go_fuzz_dep_.CoverTab[14986]++
								return syscall.AF_INET, false
//line /usr/local/go/src/net/ipsock_posix.go:115
		// _ = "end of CoverTab[14986]"
	case '6':
//line /usr/local/go/src/net/ipsock_posix.go:116
		_go_fuzz_dep_.CoverTab[14987]++
								return syscall.AF_INET6, true
//line /usr/local/go/src/net/ipsock_posix.go:117
		// _ = "end of CoverTab[14987]"
//line /usr/local/go/src/net/ipsock_posix.go:117
	default:
//line /usr/local/go/src/net/ipsock_posix.go:117
		_go_fuzz_dep_.CoverTab[14988]++
//line /usr/local/go/src/net/ipsock_posix.go:117
		// _ = "end of CoverTab[14988]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:118
	// _ = "end of CoverTab[14982]"
//line /usr/local/go/src/net/ipsock_posix.go:118
	_go_fuzz_dep_.CoverTab[14983]++

							if mode == "listen" && func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:120
		_go_fuzz_dep_.CoverTab[14989]++
//line /usr/local/go/src/net/ipsock_posix.go:120
		return (laddr == nil || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:120
			_go_fuzz_dep_.CoverTab[14990]++
//line /usr/local/go/src/net/ipsock_posix.go:120
			return laddr.isWildcard()
//line /usr/local/go/src/net/ipsock_posix.go:120
			// _ = "end of CoverTab[14990]"
//line /usr/local/go/src/net/ipsock_posix.go:120
		}())
//line /usr/local/go/src/net/ipsock_posix.go:120
		// _ = "end of CoverTab[14989]"
//line /usr/local/go/src/net/ipsock_posix.go:120
	}() {
//line /usr/local/go/src/net/ipsock_posix.go:120
		_go_fuzz_dep_.CoverTab[14991]++
								if supportsIPv4map() || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:121
			_go_fuzz_dep_.CoverTab[14994]++
//line /usr/local/go/src/net/ipsock_posix.go:121
			return !supportsIPv4()
//line /usr/local/go/src/net/ipsock_posix.go:121
			// _ = "end of CoverTab[14994]"
//line /usr/local/go/src/net/ipsock_posix.go:121
		}() {
//line /usr/local/go/src/net/ipsock_posix.go:121
			_go_fuzz_dep_.CoverTab[14995]++
									return syscall.AF_INET6, false
//line /usr/local/go/src/net/ipsock_posix.go:122
			// _ = "end of CoverTab[14995]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:123
			_go_fuzz_dep_.CoverTab[14996]++
//line /usr/local/go/src/net/ipsock_posix.go:123
			// _ = "end of CoverTab[14996]"
//line /usr/local/go/src/net/ipsock_posix.go:123
		}
//line /usr/local/go/src/net/ipsock_posix.go:123
		// _ = "end of CoverTab[14991]"
//line /usr/local/go/src/net/ipsock_posix.go:123
		_go_fuzz_dep_.CoverTab[14992]++
								if laddr == nil {
//line /usr/local/go/src/net/ipsock_posix.go:124
			_go_fuzz_dep_.CoverTab[14997]++
									return syscall.AF_INET, false
//line /usr/local/go/src/net/ipsock_posix.go:125
			// _ = "end of CoverTab[14997]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:126
			_go_fuzz_dep_.CoverTab[14998]++
//line /usr/local/go/src/net/ipsock_posix.go:126
			// _ = "end of CoverTab[14998]"
//line /usr/local/go/src/net/ipsock_posix.go:126
		}
//line /usr/local/go/src/net/ipsock_posix.go:126
		// _ = "end of CoverTab[14992]"
//line /usr/local/go/src/net/ipsock_posix.go:126
		_go_fuzz_dep_.CoverTab[14993]++
								return laddr.family(), false
//line /usr/local/go/src/net/ipsock_posix.go:127
		// _ = "end of CoverTab[14993]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:128
		_go_fuzz_dep_.CoverTab[14999]++
//line /usr/local/go/src/net/ipsock_posix.go:128
		// _ = "end of CoverTab[14999]"
//line /usr/local/go/src/net/ipsock_posix.go:128
	}
//line /usr/local/go/src/net/ipsock_posix.go:128
	// _ = "end of CoverTab[14983]"
//line /usr/local/go/src/net/ipsock_posix.go:128
	_go_fuzz_dep_.CoverTab[14984]++

							if (laddr == nil || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:130
		_go_fuzz_dep_.CoverTab[15000]++
//line /usr/local/go/src/net/ipsock_posix.go:130
		return laddr.family() == syscall.AF_INET
//line /usr/local/go/src/net/ipsock_posix.go:130
		// _ = "end of CoverTab[15000]"
//line /usr/local/go/src/net/ipsock_posix.go:130
	}()) && func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:130
		_go_fuzz_dep_.CoverTab[15001]++
//line /usr/local/go/src/net/ipsock_posix.go:130
		return (raddr == nil || func() bool {
									_go_fuzz_dep_.CoverTab[15002]++
//line /usr/local/go/src/net/ipsock_posix.go:131
			return raddr.family() == syscall.AF_INET
//line /usr/local/go/src/net/ipsock_posix.go:131
			// _ = "end of CoverTab[15002]"
//line /usr/local/go/src/net/ipsock_posix.go:131
		}())
//line /usr/local/go/src/net/ipsock_posix.go:131
		// _ = "end of CoverTab[15001]"
//line /usr/local/go/src/net/ipsock_posix.go:131
	}() {
//line /usr/local/go/src/net/ipsock_posix.go:131
		_go_fuzz_dep_.CoverTab[15003]++
								return syscall.AF_INET, false
//line /usr/local/go/src/net/ipsock_posix.go:132
		// _ = "end of CoverTab[15003]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:133
		_go_fuzz_dep_.CoverTab[15004]++
//line /usr/local/go/src/net/ipsock_posix.go:133
		// _ = "end of CoverTab[15004]"
//line /usr/local/go/src/net/ipsock_posix.go:133
	}
//line /usr/local/go/src/net/ipsock_posix.go:133
	// _ = "end of CoverTab[14984]"
//line /usr/local/go/src/net/ipsock_posix.go:133
	_go_fuzz_dep_.CoverTab[14985]++
							return syscall.AF_INET6, false
//line /usr/local/go/src/net/ipsock_posix.go:134
	// _ = "end of CoverTab[14985]"
}

func internetSocket(ctx context.Context, net string, laddr, raddr sockaddr, sotype, proto int, mode string, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) (fd *netFD, err error) {
//line /usr/local/go/src/net/ipsock_posix.go:137
	_go_fuzz_dep_.CoverTab[15005]++
							if (runtime.GOOS == "aix" || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[15007]++
//line /usr/local/go/src/net/ipsock_posix.go:138
		return runtime.GOOS == "windows"
//line /usr/local/go/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[15007]"
//line /usr/local/go/src/net/ipsock_posix.go:138
	}() || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[15008]++
//line /usr/local/go/src/net/ipsock_posix.go:138
		return runtime.GOOS == "openbsd"
//line /usr/local/go/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[15008]"
//line /usr/local/go/src/net/ipsock_posix.go:138
	}()) && func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[15009]++
//line /usr/local/go/src/net/ipsock_posix.go:138
		return mode == "dial"
//line /usr/local/go/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[15009]"
//line /usr/local/go/src/net/ipsock_posix.go:138
	}() && func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[15010]++
//line /usr/local/go/src/net/ipsock_posix.go:138
		return raddr.isWildcard()
//line /usr/local/go/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[15010]"
//line /usr/local/go/src/net/ipsock_posix.go:138
	}() {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[15011]++
								raddr = raddr.toLocal(net)
//line /usr/local/go/src/net/ipsock_posix.go:139
		// _ = "end of CoverTab[15011]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:140
		_go_fuzz_dep_.CoverTab[15012]++
//line /usr/local/go/src/net/ipsock_posix.go:140
		// _ = "end of CoverTab[15012]"
//line /usr/local/go/src/net/ipsock_posix.go:140
	}
//line /usr/local/go/src/net/ipsock_posix.go:140
	// _ = "end of CoverTab[15005]"
//line /usr/local/go/src/net/ipsock_posix.go:140
	_go_fuzz_dep_.CoverTab[15006]++
							family, ipv6only := favoriteAddrFamily(net, laddr, raddr, mode)
							return socket(ctx, net, family, sotype, proto, ipv6only, laddr, raddr, ctrlCtxFn)
//line /usr/local/go/src/net/ipsock_posix.go:142
	// _ = "end of CoverTab[15006]"
}

func ipToSockaddrInet4(ip IP, port int) (syscall.SockaddrInet4, error) {
//line /usr/local/go/src/net/ipsock_posix.go:145
	_go_fuzz_dep_.CoverTab[15013]++
							if len(ip) == 0 {
//line /usr/local/go/src/net/ipsock_posix.go:146
		_go_fuzz_dep_.CoverTab[15016]++
								ip = IPv4zero
//line /usr/local/go/src/net/ipsock_posix.go:147
		// _ = "end of CoverTab[15016]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:148
		_go_fuzz_dep_.CoverTab[15017]++
//line /usr/local/go/src/net/ipsock_posix.go:148
		// _ = "end of CoverTab[15017]"
//line /usr/local/go/src/net/ipsock_posix.go:148
	}
//line /usr/local/go/src/net/ipsock_posix.go:148
	// _ = "end of CoverTab[15013]"
//line /usr/local/go/src/net/ipsock_posix.go:148
	_go_fuzz_dep_.CoverTab[15014]++
							ip4 := ip.To4()
							if ip4 == nil {
//line /usr/local/go/src/net/ipsock_posix.go:150
		_go_fuzz_dep_.CoverTab[15018]++
								return syscall.SockaddrInet4{}, &AddrError{Err: "non-IPv4 address", Addr: ip.String()}
//line /usr/local/go/src/net/ipsock_posix.go:151
		// _ = "end of CoverTab[15018]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:152
		_go_fuzz_dep_.CoverTab[15019]++
//line /usr/local/go/src/net/ipsock_posix.go:152
		// _ = "end of CoverTab[15019]"
//line /usr/local/go/src/net/ipsock_posix.go:152
	}
//line /usr/local/go/src/net/ipsock_posix.go:152
	// _ = "end of CoverTab[15014]"
//line /usr/local/go/src/net/ipsock_posix.go:152
	_go_fuzz_dep_.CoverTab[15015]++
							sa := syscall.SockaddrInet4{Port: port}
							copy(sa.Addr[:], ip4)
							return sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:155
	// _ = "end of CoverTab[15015]"
}

func ipToSockaddrInet6(ip IP, port int, zone string) (syscall.SockaddrInet6, error) {
//line /usr/local/go/src/net/ipsock_posix.go:158
	_go_fuzz_dep_.CoverTab[15020]++

//line /usr/local/go/src/net/ipsock_posix.go:169
	if len(ip) == 0 || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:169
		_go_fuzz_dep_.CoverTab[15023]++
//line /usr/local/go/src/net/ipsock_posix.go:169
		return ip.Equal(IPv4zero)
//line /usr/local/go/src/net/ipsock_posix.go:169
		// _ = "end of CoverTab[15023]"
//line /usr/local/go/src/net/ipsock_posix.go:169
	}() {
//line /usr/local/go/src/net/ipsock_posix.go:169
		_go_fuzz_dep_.CoverTab[15024]++
								ip = IPv6zero
//line /usr/local/go/src/net/ipsock_posix.go:170
		// _ = "end of CoverTab[15024]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:171
		_go_fuzz_dep_.CoverTab[15025]++
//line /usr/local/go/src/net/ipsock_posix.go:171
		// _ = "end of CoverTab[15025]"
//line /usr/local/go/src/net/ipsock_posix.go:171
	}
//line /usr/local/go/src/net/ipsock_posix.go:171
	// _ = "end of CoverTab[15020]"
//line /usr/local/go/src/net/ipsock_posix.go:171
	_go_fuzz_dep_.CoverTab[15021]++

//line /usr/local/go/src/net/ipsock_posix.go:174
	ip6 := ip.To16()
	if ip6 == nil {
//line /usr/local/go/src/net/ipsock_posix.go:175
		_go_fuzz_dep_.CoverTab[15026]++
								return syscall.SockaddrInet6{}, &AddrError{Err: "non-IPv6 address", Addr: ip.String()}
//line /usr/local/go/src/net/ipsock_posix.go:176
		// _ = "end of CoverTab[15026]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:177
		_go_fuzz_dep_.CoverTab[15027]++
//line /usr/local/go/src/net/ipsock_posix.go:177
		// _ = "end of CoverTab[15027]"
//line /usr/local/go/src/net/ipsock_posix.go:177
	}
//line /usr/local/go/src/net/ipsock_posix.go:177
	// _ = "end of CoverTab[15021]"
//line /usr/local/go/src/net/ipsock_posix.go:177
	_go_fuzz_dep_.CoverTab[15022]++
							sa := syscall.SockaddrInet6{Port: port, ZoneId: uint32(zoneCache.index(zone))}
							copy(sa.Addr[:], ip6)
							return sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:180
	// _ = "end of CoverTab[15022]"
}

func ipToSockaddr(family int, ip IP, port int, zone string) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/ipsock_posix.go:183
	_go_fuzz_dep_.CoverTab[15028]++
							switch family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/ipsock_posix.go:185
		_go_fuzz_dep_.CoverTab[15030]++
								sa, err := ipToSockaddrInet4(ip, port)
								if err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:187
			_go_fuzz_dep_.CoverTab[15035]++
									return nil, err
//line /usr/local/go/src/net/ipsock_posix.go:188
			// _ = "end of CoverTab[15035]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:189
			_go_fuzz_dep_.CoverTab[15036]++
//line /usr/local/go/src/net/ipsock_posix.go:189
			// _ = "end of CoverTab[15036]"
//line /usr/local/go/src/net/ipsock_posix.go:189
		}
//line /usr/local/go/src/net/ipsock_posix.go:189
		// _ = "end of CoverTab[15030]"
//line /usr/local/go/src/net/ipsock_posix.go:189
		_go_fuzz_dep_.CoverTab[15031]++
								return &sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:190
		// _ = "end of CoverTab[15031]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/ipsock_posix.go:191
		_go_fuzz_dep_.CoverTab[15032]++
								sa, err := ipToSockaddrInet6(ip, port, zone)
								if err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:193
			_go_fuzz_dep_.CoverTab[15037]++
									return nil, err
//line /usr/local/go/src/net/ipsock_posix.go:194
			// _ = "end of CoverTab[15037]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:195
			_go_fuzz_dep_.CoverTab[15038]++
//line /usr/local/go/src/net/ipsock_posix.go:195
			// _ = "end of CoverTab[15038]"
//line /usr/local/go/src/net/ipsock_posix.go:195
		}
//line /usr/local/go/src/net/ipsock_posix.go:195
		// _ = "end of CoverTab[15032]"
//line /usr/local/go/src/net/ipsock_posix.go:195
		_go_fuzz_dep_.CoverTab[15033]++
								return &sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:196
		// _ = "end of CoverTab[15033]"
//line /usr/local/go/src/net/ipsock_posix.go:196
	default:
//line /usr/local/go/src/net/ipsock_posix.go:196
		_go_fuzz_dep_.CoverTab[15034]++
//line /usr/local/go/src/net/ipsock_posix.go:196
		// _ = "end of CoverTab[15034]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:197
	// _ = "end of CoverTab[15028]"
//line /usr/local/go/src/net/ipsock_posix.go:197
	_go_fuzz_dep_.CoverTab[15029]++
							return nil, &AddrError{Err: "invalid address family", Addr: ip.String()}
//line /usr/local/go/src/net/ipsock_posix.go:198
	// _ = "end of CoverTab[15029]"
}

func addrPortToSockaddrInet4(ap netip.AddrPort) (syscall.SockaddrInet4, error) {
//line /usr/local/go/src/net/ipsock_posix.go:201
	_go_fuzz_dep_.CoverTab[15039]++

//line /usr/local/go/src/net/ipsock_posix.go:204
	addr := ap.Addr()
	if !addr.Is4() {
//line /usr/local/go/src/net/ipsock_posix.go:205
		_go_fuzz_dep_.CoverTab[15041]++
								return syscall.SockaddrInet4{}, &AddrError{Err: "non-IPv4 address", Addr: addr.String()}
//line /usr/local/go/src/net/ipsock_posix.go:206
		// _ = "end of CoverTab[15041]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:207
		_go_fuzz_dep_.CoverTab[15042]++
//line /usr/local/go/src/net/ipsock_posix.go:207
		// _ = "end of CoverTab[15042]"
//line /usr/local/go/src/net/ipsock_posix.go:207
	}
//line /usr/local/go/src/net/ipsock_posix.go:207
	// _ = "end of CoverTab[15039]"
//line /usr/local/go/src/net/ipsock_posix.go:207
	_go_fuzz_dep_.CoverTab[15040]++
							sa := syscall.SockaddrInet4{
		Addr:	addr.As4(),
		Port:	int(ap.Port()),
	}
							return sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:212
	// _ = "end of CoverTab[15040]"
}

func addrPortToSockaddrInet6(ap netip.AddrPort) (syscall.SockaddrInet6, error) {
//line /usr/local/go/src/net/ipsock_posix.go:215
	_go_fuzz_dep_.CoverTab[15043]++

//line /usr/local/go/src/net/ipsock_posix.go:222
	addr := ap.Addr()
	if !addr.IsValid() {
//line /usr/local/go/src/net/ipsock_posix.go:223
		_go_fuzz_dep_.CoverTab[15045]++
								return syscall.SockaddrInet6{}, &AddrError{Err: "non-IPv6 address", Addr: addr.String()}
//line /usr/local/go/src/net/ipsock_posix.go:224
		// _ = "end of CoverTab[15045]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:225
		_go_fuzz_dep_.CoverTab[15046]++
//line /usr/local/go/src/net/ipsock_posix.go:225
		// _ = "end of CoverTab[15046]"
//line /usr/local/go/src/net/ipsock_posix.go:225
	}
//line /usr/local/go/src/net/ipsock_posix.go:225
	// _ = "end of CoverTab[15043]"
//line /usr/local/go/src/net/ipsock_posix.go:225
	_go_fuzz_dep_.CoverTab[15044]++
							sa := syscall.SockaddrInet6{
		Addr:	addr.As16(),
		Port:	int(ap.Port()),
		ZoneId:	uint32(zoneCache.index(addr.Zone())),
	}
							return sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:231
	// _ = "end of CoverTab[15044]"
}

//line /usr/local/go/src/net/ipsock_posix.go:232
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/ipsock_posix.go:232
var _ = _go_fuzz_dep_.CoverTab
