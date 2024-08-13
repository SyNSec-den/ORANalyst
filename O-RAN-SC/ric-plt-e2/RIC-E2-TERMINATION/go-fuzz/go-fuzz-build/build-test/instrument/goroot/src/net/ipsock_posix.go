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
	_go_fuzz_dep_.CoverTab[6572]++
							s, err := sysSocket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
							switch err {
	case syscall.EAFNOSUPPORT, syscall.EPROTONOSUPPORT:
//line /usr/local/go/src/net/ipsock_posix.go:30
		_go_fuzz_dep_.CoverTab[6575]++
//line /usr/local/go/src/net/ipsock_posix.go:30
		// _ = "end of CoverTab[6575]"
	case nil:
//line /usr/local/go/src/net/ipsock_posix.go:31
		_go_fuzz_dep_.CoverTab[6576]++
								poll.CloseFunc(s)
								p.ipv4Enabled = true
//line /usr/local/go/src/net/ipsock_posix.go:33
		// _ = "end of CoverTab[6576]"
//line /usr/local/go/src/net/ipsock_posix.go:33
	default:
//line /usr/local/go/src/net/ipsock_posix.go:33
		_go_fuzz_dep_.CoverTab[6577]++
//line /usr/local/go/src/net/ipsock_posix.go:33
		// _ = "end of CoverTab[6577]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:34
	// _ = "end of CoverTab[6572]"
//line /usr/local/go/src/net/ipsock_posix.go:34
	_go_fuzz_dep_.CoverTab[6573]++
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
		_go_fuzz_dep_.CoverTab[6578]++

//line /usr/local/go/src/net/ipsock_posix.go:49
		probes = probes[:1]
//line /usr/local/go/src/net/ipsock_posix.go:49
		// _ = "end of CoverTab[6578]"
//line /usr/local/go/src/net/ipsock_posix.go:49
	default:
//line /usr/local/go/src/net/ipsock_posix.go:49
		_go_fuzz_dep_.CoverTab[6579]++
//line /usr/local/go/src/net/ipsock_posix.go:49
		// _ = "end of CoverTab[6579]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:50
	// _ = "end of CoverTab[6573]"
//line /usr/local/go/src/net/ipsock_posix.go:50
	_go_fuzz_dep_.CoverTab[6574]++
							for i := range probes {
//line /usr/local/go/src/net/ipsock_posix.go:51
		_go_fuzz_dep_.CoverTab[6580]++
								s, err := sysSocket(syscall.AF_INET6, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
								if err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:53
			_go_fuzz_dep_.CoverTab[6584]++
									continue
//line /usr/local/go/src/net/ipsock_posix.go:54
			// _ = "end of CoverTab[6584]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:55
			_go_fuzz_dep_.CoverTab[6585]++
//line /usr/local/go/src/net/ipsock_posix.go:55
			// _ = "end of CoverTab[6585]"
//line /usr/local/go/src/net/ipsock_posix.go:55
		}
//line /usr/local/go/src/net/ipsock_posix.go:55
		// _ = "end of CoverTab[6580]"
//line /usr/local/go/src/net/ipsock_posix.go:55
		_go_fuzz_dep_.CoverTab[6581]++
								defer poll.CloseFunc(s)
								syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, probes[i].value)
								sa, err := probes[i].laddr.sockaddr(syscall.AF_INET6)
								if err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:59
			_go_fuzz_dep_.CoverTab[6586]++
									continue
//line /usr/local/go/src/net/ipsock_posix.go:60
			// _ = "end of CoverTab[6586]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:61
			_go_fuzz_dep_.CoverTab[6587]++
//line /usr/local/go/src/net/ipsock_posix.go:61
			// _ = "end of CoverTab[6587]"
//line /usr/local/go/src/net/ipsock_posix.go:61
		}
//line /usr/local/go/src/net/ipsock_posix.go:61
		// _ = "end of CoverTab[6581]"
//line /usr/local/go/src/net/ipsock_posix.go:61
		_go_fuzz_dep_.CoverTab[6582]++
								if err := syscall.Bind(s, sa); err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:62
			_go_fuzz_dep_.CoverTab[6588]++
									continue
//line /usr/local/go/src/net/ipsock_posix.go:63
			// _ = "end of CoverTab[6588]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:64
			_go_fuzz_dep_.CoverTab[6589]++
//line /usr/local/go/src/net/ipsock_posix.go:64
			// _ = "end of CoverTab[6589]"
//line /usr/local/go/src/net/ipsock_posix.go:64
		}
//line /usr/local/go/src/net/ipsock_posix.go:64
		// _ = "end of CoverTab[6582]"
//line /usr/local/go/src/net/ipsock_posix.go:64
		_go_fuzz_dep_.CoverTab[6583]++
								if i == 0 {
//line /usr/local/go/src/net/ipsock_posix.go:65
			_go_fuzz_dep_.CoverTab[6590]++
									p.ipv6Enabled = true
//line /usr/local/go/src/net/ipsock_posix.go:66
			// _ = "end of CoverTab[6590]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:67
			_go_fuzz_dep_.CoverTab[6591]++
									p.ipv4MappedIPv6Enabled = true
//line /usr/local/go/src/net/ipsock_posix.go:68
			// _ = "end of CoverTab[6591]"
		}
//line /usr/local/go/src/net/ipsock_posix.go:69
		// _ = "end of CoverTab[6583]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:70
	// _ = "end of CoverTab[6574]"
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
	_go_fuzz_dep_.CoverTab[6592]++
							switch network[len(network)-1] {
	case '4':
//line /usr/local/go/src/net/ipsock_posix.go:114
		_go_fuzz_dep_.CoverTab[6596]++
								return syscall.AF_INET, false
//line /usr/local/go/src/net/ipsock_posix.go:115
		// _ = "end of CoverTab[6596]"
	case '6':
//line /usr/local/go/src/net/ipsock_posix.go:116
		_go_fuzz_dep_.CoverTab[6597]++
								return syscall.AF_INET6, true
//line /usr/local/go/src/net/ipsock_posix.go:117
		// _ = "end of CoverTab[6597]"
//line /usr/local/go/src/net/ipsock_posix.go:117
	default:
//line /usr/local/go/src/net/ipsock_posix.go:117
		_go_fuzz_dep_.CoverTab[6598]++
//line /usr/local/go/src/net/ipsock_posix.go:117
		// _ = "end of CoverTab[6598]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:118
	// _ = "end of CoverTab[6592]"
//line /usr/local/go/src/net/ipsock_posix.go:118
	_go_fuzz_dep_.CoverTab[6593]++

							if mode == "listen" && func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:120
		_go_fuzz_dep_.CoverTab[6599]++
//line /usr/local/go/src/net/ipsock_posix.go:120
		return (laddr == nil || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:120
			_go_fuzz_dep_.CoverTab[6600]++
//line /usr/local/go/src/net/ipsock_posix.go:120
			return laddr.isWildcard()
//line /usr/local/go/src/net/ipsock_posix.go:120
			// _ = "end of CoverTab[6600]"
//line /usr/local/go/src/net/ipsock_posix.go:120
		}())
//line /usr/local/go/src/net/ipsock_posix.go:120
		// _ = "end of CoverTab[6599]"
//line /usr/local/go/src/net/ipsock_posix.go:120
	}() {
//line /usr/local/go/src/net/ipsock_posix.go:120
		_go_fuzz_dep_.CoverTab[6601]++
								if supportsIPv4map() || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:121
			_go_fuzz_dep_.CoverTab[6604]++
//line /usr/local/go/src/net/ipsock_posix.go:121
			return !supportsIPv4()
//line /usr/local/go/src/net/ipsock_posix.go:121
			// _ = "end of CoverTab[6604]"
//line /usr/local/go/src/net/ipsock_posix.go:121
		}() {
//line /usr/local/go/src/net/ipsock_posix.go:121
			_go_fuzz_dep_.CoverTab[6605]++
									return syscall.AF_INET6, false
//line /usr/local/go/src/net/ipsock_posix.go:122
			// _ = "end of CoverTab[6605]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:123
			_go_fuzz_dep_.CoverTab[6606]++
//line /usr/local/go/src/net/ipsock_posix.go:123
			// _ = "end of CoverTab[6606]"
//line /usr/local/go/src/net/ipsock_posix.go:123
		}
//line /usr/local/go/src/net/ipsock_posix.go:123
		// _ = "end of CoverTab[6601]"
//line /usr/local/go/src/net/ipsock_posix.go:123
		_go_fuzz_dep_.CoverTab[6602]++
								if laddr == nil {
//line /usr/local/go/src/net/ipsock_posix.go:124
			_go_fuzz_dep_.CoverTab[6607]++
									return syscall.AF_INET, false
//line /usr/local/go/src/net/ipsock_posix.go:125
			// _ = "end of CoverTab[6607]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:126
			_go_fuzz_dep_.CoverTab[6608]++
//line /usr/local/go/src/net/ipsock_posix.go:126
			// _ = "end of CoverTab[6608]"
//line /usr/local/go/src/net/ipsock_posix.go:126
		}
//line /usr/local/go/src/net/ipsock_posix.go:126
		// _ = "end of CoverTab[6602]"
//line /usr/local/go/src/net/ipsock_posix.go:126
		_go_fuzz_dep_.CoverTab[6603]++
								return laddr.family(), false
//line /usr/local/go/src/net/ipsock_posix.go:127
		// _ = "end of CoverTab[6603]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:128
		_go_fuzz_dep_.CoverTab[6609]++
//line /usr/local/go/src/net/ipsock_posix.go:128
		// _ = "end of CoverTab[6609]"
//line /usr/local/go/src/net/ipsock_posix.go:128
	}
//line /usr/local/go/src/net/ipsock_posix.go:128
	// _ = "end of CoverTab[6593]"
//line /usr/local/go/src/net/ipsock_posix.go:128
	_go_fuzz_dep_.CoverTab[6594]++

							if (laddr == nil || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:130
		_go_fuzz_dep_.CoverTab[6610]++
//line /usr/local/go/src/net/ipsock_posix.go:130
		return laddr.family() == syscall.AF_INET
//line /usr/local/go/src/net/ipsock_posix.go:130
		// _ = "end of CoverTab[6610]"
//line /usr/local/go/src/net/ipsock_posix.go:130
	}()) && func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:130
		_go_fuzz_dep_.CoverTab[6611]++
//line /usr/local/go/src/net/ipsock_posix.go:130
		return (raddr == nil || func() bool {
									_go_fuzz_dep_.CoverTab[6612]++
//line /usr/local/go/src/net/ipsock_posix.go:131
			return raddr.family() == syscall.AF_INET
//line /usr/local/go/src/net/ipsock_posix.go:131
			// _ = "end of CoverTab[6612]"
//line /usr/local/go/src/net/ipsock_posix.go:131
		}())
//line /usr/local/go/src/net/ipsock_posix.go:131
		// _ = "end of CoverTab[6611]"
//line /usr/local/go/src/net/ipsock_posix.go:131
	}() {
//line /usr/local/go/src/net/ipsock_posix.go:131
		_go_fuzz_dep_.CoverTab[6613]++
								return syscall.AF_INET, false
//line /usr/local/go/src/net/ipsock_posix.go:132
		// _ = "end of CoverTab[6613]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:133
		_go_fuzz_dep_.CoverTab[6614]++
//line /usr/local/go/src/net/ipsock_posix.go:133
		// _ = "end of CoverTab[6614]"
//line /usr/local/go/src/net/ipsock_posix.go:133
	}
//line /usr/local/go/src/net/ipsock_posix.go:133
	// _ = "end of CoverTab[6594]"
//line /usr/local/go/src/net/ipsock_posix.go:133
	_go_fuzz_dep_.CoverTab[6595]++
							return syscall.AF_INET6, false
//line /usr/local/go/src/net/ipsock_posix.go:134
	// _ = "end of CoverTab[6595]"
}

func internetSocket(ctx context.Context, net string, laddr, raddr sockaddr, sotype, proto int, mode string, ctrlCtxFn func(context.Context, string, string, syscall.RawConn) error) (fd *netFD, err error) {
//line /usr/local/go/src/net/ipsock_posix.go:137
	_go_fuzz_dep_.CoverTab[6615]++
							if (runtime.GOOS == "aix" || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6617]++
//line /usr/local/go/src/net/ipsock_posix.go:138
		return runtime.GOOS == "windows"
//line /usr/local/go/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[6617]"
//line /usr/local/go/src/net/ipsock_posix.go:138
	}() || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6618]++
//line /usr/local/go/src/net/ipsock_posix.go:138
		return runtime.GOOS == "openbsd"
//line /usr/local/go/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[6618]"
//line /usr/local/go/src/net/ipsock_posix.go:138
	}()) && func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6619]++
//line /usr/local/go/src/net/ipsock_posix.go:138
		return mode == "dial"
//line /usr/local/go/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[6619]"
//line /usr/local/go/src/net/ipsock_posix.go:138
	}() && func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6620]++
//line /usr/local/go/src/net/ipsock_posix.go:138
		return raddr.isWildcard()
//line /usr/local/go/src/net/ipsock_posix.go:138
		// _ = "end of CoverTab[6620]"
//line /usr/local/go/src/net/ipsock_posix.go:138
	}() {
//line /usr/local/go/src/net/ipsock_posix.go:138
		_go_fuzz_dep_.CoverTab[6621]++
								raddr = raddr.toLocal(net)
//line /usr/local/go/src/net/ipsock_posix.go:139
		// _ = "end of CoverTab[6621]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:140
		_go_fuzz_dep_.CoverTab[6622]++
//line /usr/local/go/src/net/ipsock_posix.go:140
		// _ = "end of CoverTab[6622]"
//line /usr/local/go/src/net/ipsock_posix.go:140
	}
//line /usr/local/go/src/net/ipsock_posix.go:140
	// _ = "end of CoverTab[6615]"
//line /usr/local/go/src/net/ipsock_posix.go:140
	_go_fuzz_dep_.CoverTab[6616]++
							family, ipv6only := favoriteAddrFamily(net, laddr, raddr, mode)
							return socket(ctx, net, family, sotype, proto, ipv6only, laddr, raddr, ctrlCtxFn)
//line /usr/local/go/src/net/ipsock_posix.go:142
	// _ = "end of CoverTab[6616]"
}

func ipToSockaddrInet4(ip IP, port int) (syscall.SockaddrInet4, error) {
//line /usr/local/go/src/net/ipsock_posix.go:145
	_go_fuzz_dep_.CoverTab[6623]++
							if len(ip) == 0 {
//line /usr/local/go/src/net/ipsock_posix.go:146
		_go_fuzz_dep_.CoverTab[6626]++
								ip = IPv4zero
//line /usr/local/go/src/net/ipsock_posix.go:147
		// _ = "end of CoverTab[6626]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:148
		_go_fuzz_dep_.CoverTab[6627]++
//line /usr/local/go/src/net/ipsock_posix.go:148
		// _ = "end of CoverTab[6627]"
//line /usr/local/go/src/net/ipsock_posix.go:148
	}
//line /usr/local/go/src/net/ipsock_posix.go:148
	// _ = "end of CoverTab[6623]"
//line /usr/local/go/src/net/ipsock_posix.go:148
	_go_fuzz_dep_.CoverTab[6624]++
							ip4 := ip.To4()
							if ip4 == nil {
//line /usr/local/go/src/net/ipsock_posix.go:150
		_go_fuzz_dep_.CoverTab[6628]++
								return syscall.SockaddrInet4{}, &AddrError{Err: "non-IPv4 address", Addr: ip.String()}
//line /usr/local/go/src/net/ipsock_posix.go:151
		// _ = "end of CoverTab[6628]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:152
		_go_fuzz_dep_.CoverTab[6629]++
//line /usr/local/go/src/net/ipsock_posix.go:152
		// _ = "end of CoverTab[6629]"
//line /usr/local/go/src/net/ipsock_posix.go:152
	}
//line /usr/local/go/src/net/ipsock_posix.go:152
	// _ = "end of CoverTab[6624]"
//line /usr/local/go/src/net/ipsock_posix.go:152
	_go_fuzz_dep_.CoverTab[6625]++
							sa := syscall.SockaddrInet4{Port: port}
							copy(sa.Addr[:], ip4)
							return sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:155
	// _ = "end of CoverTab[6625]"
}

func ipToSockaddrInet6(ip IP, port int, zone string) (syscall.SockaddrInet6, error) {
//line /usr/local/go/src/net/ipsock_posix.go:158
	_go_fuzz_dep_.CoverTab[6630]++

//line /usr/local/go/src/net/ipsock_posix.go:169
	if len(ip) == 0 || func() bool {
//line /usr/local/go/src/net/ipsock_posix.go:169
		_go_fuzz_dep_.CoverTab[6633]++
//line /usr/local/go/src/net/ipsock_posix.go:169
		return ip.Equal(IPv4zero)
//line /usr/local/go/src/net/ipsock_posix.go:169
		// _ = "end of CoverTab[6633]"
//line /usr/local/go/src/net/ipsock_posix.go:169
	}() {
//line /usr/local/go/src/net/ipsock_posix.go:169
		_go_fuzz_dep_.CoverTab[6634]++
								ip = IPv6zero
//line /usr/local/go/src/net/ipsock_posix.go:170
		// _ = "end of CoverTab[6634]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:171
		_go_fuzz_dep_.CoverTab[6635]++
//line /usr/local/go/src/net/ipsock_posix.go:171
		// _ = "end of CoverTab[6635]"
//line /usr/local/go/src/net/ipsock_posix.go:171
	}
//line /usr/local/go/src/net/ipsock_posix.go:171
	// _ = "end of CoverTab[6630]"
//line /usr/local/go/src/net/ipsock_posix.go:171
	_go_fuzz_dep_.CoverTab[6631]++

//line /usr/local/go/src/net/ipsock_posix.go:174
	ip6 := ip.To16()
	if ip6 == nil {
//line /usr/local/go/src/net/ipsock_posix.go:175
		_go_fuzz_dep_.CoverTab[6636]++
								return syscall.SockaddrInet6{}, &AddrError{Err: "non-IPv6 address", Addr: ip.String()}
//line /usr/local/go/src/net/ipsock_posix.go:176
		// _ = "end of CoverTab[6636]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:177
		_go_fuzz_dep_.CoverTab[6637]++
//line /usr/local/go/src/net/ipsock_posix.go:177
		// _ = "end of CoverTab[6637]"
//line /usr/local/go/src/net/ipsock_posix.go:177
	}
//line /usr/local/go/src/net/ipsock_posix.go:177
	// _ = "end of CoverTab[6631]"
//line /usr/local/go/src/net/ipsock_posix.go:177
	_go_fuzz_dep_.CoverTab[6632]++
							sa := syscall.SockaddrInet6{Port: port, ZoneId: uint32(zoneCache.index(zone))}
							copy(sa.Addr[:], ip6)
							return sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:180
	// _ = "end of CoverTab[6632]"
}

func ipToSockaddr(family int, ip IP, port int, zone string) (syscall.Sockaddr, error) {
//line /usr/local/go/src/net/ipsock_posix.go:183
	_go_fuzz_dep_.CoverTab[6638]++
							switch family {
	case syscall.AF_INET:
//line /usr/local/go/src/net/ipsock_posix.go:185
		_go_fuzz_dep_.CoverTab[6640]++
								sa, err := ipToSockaddrInet4(ip, port)
								if err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:187
			_go_fuzz_dep_.CoverTab[6645]++
									return nil, err
//line /usr/local/go/src/net/ipsock_posix.go:188
			// _ = "end of CoverTab[6645]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:189
			_go_fuzz_dep_.CoverTab[6646]++
//line /usr/local/go/src/net/ipsock_posix.go:189
			// _ = "end of CoverTab[6646]"
//line /usr/local/go/src/net/ipsock_posix.go:189
		}
//line /usr/local/go/src/net/ipsock_posix.go:189
		// _ = "end of CoverTab[6640]"
//line /usr/local/go/src/net/ipsock_posix.go:189
		_go_fuzz_dep_.CoverTab[6641]++
								return &sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:190
		// _ = "end of CoverTab[6641]"
	case syscall.AF_INET6:
//line /usr/local/go/src/net/ipsock_posix.go:191
		_go_fuzz_dep_.CoverTab[6642]++
								sa, err := ipToSockaddrInet6(ip, port, zone)
								if err != nil {
//line /usr/local/go/src/net/ipsock_posix.go:193
			_go_fuzz_dep_.CoverTab[6647]++
									return nil, err
//line /usr/local/go/src/net/ipsock_posix.go:194
			// _ = "end of CoverTab[6647]"
		} else {
//line /usr/local/go/src/net/ipsock_posix.go:195
			_go_fuzz_dep_.CoverTab[6648]++
//line /usr/local/go/src/net/ipsock_posix.go:195
			// _ = "end of CoverTab[6648]"
//line /usr/local/go/src/net/ipsock_posix.go:195
		}
//line /usr/local/go/src/net/ipsock_posix.go:195
		// _ = "end of CoverTab[6642]"
//line /usr/local/go/src/net/ipsock_posix.go:195
		_go_fuzz_dep_.CoverTab[6643]++
								return &sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:196
		// _ = "end of CoverTab[6643]"
//line /usr/local/go/src/net/ipsock_posix.go:196
	default:
//line /usr/local/go/src/net/ipsock_posix.go:196
		_go_fuzz_dep_.CoverTab[6644]++
//line /usr/local/go/src/net/ipsock_posix.go:196
		// _ = "end of CoverTab[6644]"
	}
//line /usr/local/go/src/net/ipsock_posix.go:197
	// _ = "end of CoverTab[6638]"
//line /usr/local/go/src/net/ipsock_posix.go:197
	_go_fuzz_dep_.CoverTab[6639]++
							return nil, &AddrError{Err: "invalid address family", Addr: ip.String()}
//line /usr/local/go/src/net/ipsock_posix.go:198
	// _ = "end of CoverTab[6639]"
}

func addrPortToSockaddrInet4(ap netip.AddrPort) (syscall.SockaddrInet4, error) {
//line /usr/local/go/src/net/ipsock_posix.go:201
	_go_fuzz_dep_.CoverTab[6649]++

//line /usr/local/go/src/net/ipsock_posix.go:204
	addr := ap.Addr()
	if !addr.Is4() {
//line /usr/local/go/src/net/ipsock_posix.go:205
		_go_fuzz_dep_.CoverTab[6651]++
								return syscall.SockaddrInet4{}, &AddrError{Err: "non-IPv4 address", Addr: addr.String()}
//line /usr/local/go/src/net/ipsock_posix.go:206
		// _ = "end of CoverTab[6651]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:207
		_go_fuzz_dep_.CoverTab[6652]++
//line /usr/local/go/src/net/ipsock_posix.go:207
		// _ = "end of CoverTab[6652]"
//line /usr/local/go/src/net/ipsock_posix.go:207
	}
//line /usr/local/go/src/net/ipsock_posix.go:207
	// _ = "end of CoverTab[6649]"
//line /usr/local/go/src/net/ipsock_posix.go:207
	_go_fuzz_dep_.CoverTab[6650]++
							sa := syscall.SockaddrInet4{
		Addr:	addr.As4(),
		Port:	int(ap.Port()),
	}
							return sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:212
	// _ = "end of CoverTab[6650]"
}

func addrPortToSockaddrInet6(ap netip.AddrPort) (syscall.SockaddrInet6, error) {
//line /usr/local/go/src/net/ipsock_posix.go:215
	_go_fuzz_dep_.CoverTab[6653]++

//line /usr/local/go/src/net/ipsock_posix.go:222
	addr := ap.Addr()
	if !addr.IsValid() {
//line /usr/local/go/src/net/ipsock_posix.go:223
		_go_fuzz_dep_.CoverTab[6655]++
								return syscall.SockaddrInet6{}, &AddrError{Err: "non-IPv6 address", Addr: addr.String()}
//line /usr/local/go/src/net/ipsock_posix.go:224
		// _ = "end of CoverTab[6655]"
	} else {
//line /usr/local/go/src/net/ipsock_posix.go:225
		_go_fuzz_dep_.CoverTab[6656]++
//line /usr/local/go/src/net/ipsock_posix.go:225
		// _ = "end of CoverTab[6656]"
//line /usr/local/go/src/net/ipsock_posix.go:225
	}
//line /usr/local/go/src/net/ipsock_posix.go:225
	// _ = "end of CoverTab[6653]"
//line /usr/local/go/src/net/ipsock_posix.go:225
	_go_fuzz_dep_.CoverTab[6654]++
							sa := syscall.SockaddrInet6{
		Addr:	addr.As16(),
		Port:	int(ap.Port()),
		ZoneId:	uint32(zoneCache.index(addr.Zone())),
	}
							return sa, nil
//line /usr/local/go/src/net/ipsock_posix.go:231
	// _ = "end of CoverTab[6654]"
}

//line /usr/local/go/src/net/ipsock_posix.go:232
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/ipsock_posix.go:232
var _ = _go_fuzz_dep_.CoverTab
