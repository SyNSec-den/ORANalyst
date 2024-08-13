// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:5
package proxy

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:5
)

import (
	"context"
	"net"
	"strings"
)

// A PerHost directs connections to a default Dialer unless the host name
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:13
// requested matches one of a number of exceptions.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:15
type PerHost struct {
	def, bypass	Dialer

	bypassNetworks	[]*net.IPNet
	bypassIPs	[]net.IP
	bypassZones	[]string
	bypassHosts	[]string
}

// NewPerHost returns a PerHost Dialer that directs connections to either
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:24
// defaultDialer or bypass, depending on whether the connection matches one of
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:24
// the configured rules.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:27
func NewPerHost(defaultDialer, bypass Dialer) *PerHost {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:27
	_go_fuzz_dep_.CoverTab[96949]++
											return &PerHost{
		def:	defaultDialer,
		bypass:	bypass,
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:31
	// _ = "end of CoverTab[96949]"
}

// Dial connects to the address addr on the given network through either
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:34
// defaultDialer or bypass.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:36
func (p *PerHost) Dial(network, addr string) (c net.Conn, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:36
	_go_fuzz_dep_.CoverTab[96950]++
											host, _, err := net.SplitHostPort(addr)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:38
		_go_fuzz_dep_.CoverTab[96952]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:39
		// _ = "end of CoverTab[96952]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:40
		_go_fuzz_dep_.CoverTab[96953]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:40
		// _ = "end of CoverTab[96953]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:40
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:40
	// _ = "end of CoverTab[96950]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:40
	_go_fuzz_dep_.CoverTab[96951]++

											return p.dialerForRequest(host).Dial(network, addr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:42
	// _ = "end of CoverTab[96951]"
}

// DialContext connects to the address addr on the given network through either
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:45
// defaultDialer or bypass.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:47
func (p *PerHost) DialContext(ctx context.Context, network, addr string) (c net.Conn, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:47
	_go_fuzz_dep_.CoverTab[96954]++
											host, _, err := net.SplitHostPort(addr)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:49
		_go_fuzz_dep_.CoverTab[96957]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:50
		// _ = "end of CoverTab[96957]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:51
		_go_fuzz_dep_.CoverTab[96958]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:51
		// _ = "end of CoverTab[96958]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:51
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:51
	// _ = "end of CoverTab[96954]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:51
	_go_fuzz_dep_.CoverTab[96955]++
											d := p.dialerForRequest(host)
											if x, ok := d.(ContextDialer); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:53
		_go_fuzz_dep_.CoverTab[96959]++
												return x.DialContext(ctx, network, addr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:54
		// _ = "end of CoverTab[96959]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:55
		_go_fuzz_dep_.CoverTab[96960]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:55
		// _ = "end of CoverTab[96960]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:55
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:55
	// _ = "end of CoverTab[96955]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:55
	_go_fuzz_dep_.CoverTab[96956]++
											return dialContext(ctx, d, network, addr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:56
	// _ = "end of CoverTab[96956]"
}

func (p *PerHost) dialerForRequest(host string) Dialer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:59
	_go_fuzz_dep_.CoverTab[96961]++
											if ip := net.ParseIP(host); ip != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:60
		_go_fuzz_dep_.CoverTab[96965]++
												for _, net := range p.bypassNetworks {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:61
			_go_fuzz_dep_.CoverTab[96968]++
													if net.Contains(ip) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:62
				_go_fuzz_dep_.CoverTab[96969]++
														return p.bypass
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:63
				// _ = "end of CoverTab[96969]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:64
				_go_fuzz_dep_.CoverTab[96970]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:64
				// _ = "end of CoverTab[96970]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:64
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:64
			// _ = "end of CoverTab[96968]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:65
		// _ = "end of CoverTab[96965]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:65
		_go_fuzz_dep_.CoverTab[96966]++
												for _, bypassIP := range p.bypassIPs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:66
			_go_fuzz_dep_.CoverTab[96971]++
													if bypassIP.Equal(ip) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:67
				_go_fuzz_dep_.CoverTab[96972]++
														return p.bypass
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:68
				// _ = "end of CoverTab[96972]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:69
				_go_fuzz_dep_.CoverTab[96973]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:69
				// _ = "end of CoverTab[96973]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:69
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:69
			// _ = "end of CoverTab[96971]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:70
		// _ = "end of CoverTab[96966]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:70
		_go_fuzz_dep_.CoverTab[96967]++
												return p.def
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:71
		// _ = "end of CoverTab[96967]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:72
		_go_fuzz_dep_.CoverTab[96974]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:72
		// _ = "end of CoverTab[96974]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:72
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:72
	// _ = "end of CoverTab[96961]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:72
	_go_fuzz_dep_.CoverTab[96962]++

											for _, zone := range p.bypassZones {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:74
		_go_fuzz_dep_.CoverTab[96975]++
												if strings.HasSuffix(host, zone) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:75
			_go_fuzz_dep_.CoverTab[96977]++
													return p.bypass
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:76
			// _ = "end of CoverTab[96977]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:77
			_go_fuzz_dep_.CoverTab[96978]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:77
			// _ = "end of CoverTab[96978]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:77
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:77
		// _ = "end of CoverTab[96975]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:77
		_go_fuzz_dep_.CoverTab[96976]++
												if host == zone[1:] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:78
			_go_fuzz_dep_.CoverTab[96979]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:81
			return p.bypass
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:81
			// _ = "end of CoverTab[96979]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:82
			_go_fuzz_dep_.CoverTab[96980]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:82
			// _ = "end of CoverTab[96980]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:82
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:82
		// _ = "end of CoverTab[96976]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:83
	// _ = "end of CoverTab[96962]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:83
	_go_fuzz_dep_.CoverTab[96963]++
											for _, bypassHost := range p.bypassHosts {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:84
		_go_fuzz_dep_.CoverTab[96981]++
												if bypassHost == host {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:85
			_go_fuzz_dep_.CoverTab[96982]++
													return p.bypass
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:86
			// _ = "end of CoverTab[96982]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:87
			_go_fuzz_dep_.CoverTab[96983]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:87
			// _ = "end of CoverTab[96983]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:87
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:87
		// _ = "end of CoverTab[96981]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:88
	// _ = "end of CoverTab[96963]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:88
	_go_fuzz_dep_.CoverTab[96964]++
											return p.def
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:89
	// _ = "end of CoverTab[96964]"
}

// AddFromString parses a string that contains comma-separated values
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:92
// specifying hosts that should use the bypass proxy. Each value is either an
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:92
// IP address, a CIDR range, a zone (*.example.com) or a host name
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:92
// (localhost). A best effort is made to parse the string and errors are
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:92
// ignored.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:97
func (p *PerHost) AddFromString(s string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:97
	_go_fuzz_dep_.CoverTab[96984]++
											hosts := strings.Split(s, ",")
											for _, host := range hosts {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:99
		_go_fuzz_dep_.CoverTab[96985]++
												host = strings.TrimSpace(host)
												if len(host) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:101
			_go_fuzz_dep_.CoverTab[96990]++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:102
			// _ = "end of CoverTab[96990]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:103
			_go_fuzz_dep_.CoverTab[96991]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:103
			// _ = "end of CoverTab[96991]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:103
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:103
		// _ = "end of CoverTab[96985]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:103
		_go_fuzz_dep_.CoverTab[96986]++
												if strings.Contains(host, "/") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:104
			_go_fuzz_dep_.CoverTab[96992]++

													if _, net, err := net.ParseCIDR(host); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:106
				_go_fuzz_dep_.CoverTab[96994]++
														p.AddNetwork(net)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:107
				// _ = "end of CoverTab[96994]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:108
				_go_fuzz_dep_.CoverTab[96995]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:108
				// _ = "end of CoverTab[96995]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:108
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:108
			// _ = "end of CoverTab[96992]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:108
			_go_fuzz_dep_.CoverTab[96993]++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:109
			// _ = "end of CoverTab[96993]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:110
			_go_fuzz_dep_.CoverTab[96996]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:110
			// _ = "end of CoverTab[96996]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:110
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:110
		// _ = "end of CoverTab[96986]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:110
		_go_fuzz_dep_.CoverTab[96987]++
												if ip := net.ParseIP(host); ip != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:111
			_go_fuzz_dep_.CoverTab[96997]++
													p.AddIP(ip)
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:113
			// _ = "end of CoverTab[96997]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:114
			_go_fuzz_dep_.CoverTab[96998]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:114
			// _ = "end of CoverTab[96998]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:114
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:114
		// _ = "end of CoverTab[96987]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:114
		_go_fuzz_dep_.CoverTab[96988]++
												if strings.HasPrefix(host, "*.") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:115
			_go_fuzz_dep_.CoverTab[96999]++
													p.AddZone(host[1:])
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:117
			// _ = "end of CoverTab[96999]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:118
			_go_fuzz_dep_.CoverTab[97000]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:118
			// _ = "end of CoverTab[97000]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:118
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:118
		// _ = "end of CoverTab[96988]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:118
		_go_fuzz_dep_.CoverTab[96989]++
												p.AddHost(host)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:119
		// _ = "end of CoverTab[96989]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:120
	// _ = "end of CoverTab[96984]"
}

// AddIP specifies an IP address that will use the bypass proxy. Note that
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:123
// this will only take effect if a literal IP address is dialed. A connection
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:123
// to a named host will never match an IP.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:126
func (p *PerHost) AddIP(ip net.IP) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:126
	_go_fuzz_dep_.CoverTab[97001]++
											p.bypassIPs = append(p.bypassIPs, ip)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:127
	// _ = "end of CoverTab[97001]"
}

// AddNetwork specifies an IP range that will use the bypass proxy. Note that
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:130
// this will only take effect if a literal IP address is dialed. A connection
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:130
// to a named host will never match.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:133
func (p *PerHost) AddNetwork(net *net.IPNet) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:133
	_go_fuzz_dep_.CoverTab[97002]++
											p.bypassNetworks = append(p.bypassNetworks, net)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:134
	// _ = "end of CoverTab[97002]"
}

// AddZone specifies a DNS suffix that will use the bypass proxy. A zone of
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:137
// "example.com" matches "example.com" and all of its subdomains.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:139
func (p *PerHost) AddZone(zone string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:139
	_go_fuzz_dep_.CoverTab[97003]++
											if strings.HasSuffix(zone, ".") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:140
		_go_fuzz_dep_.CoverTab[97006]++
												zone = zone[:len(zone)-1]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:141
		// _ = "end of CoverTab[97006]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:142
		_go_fuzz_dep_.CoverTab[97007]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:142
		// _ = "end of CoverTab[97007]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:142
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:142
	// _ = "end of CoverTab[97003]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:142
	_go_fuzz_dep_.CoverTab[97004]++
											if !strings.HasPrefix(zone, ".") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:143
		_go_fuzz_dep_.CoverTab[97008]++
												zone = "." + zone
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:144
		// _ = "end of CoverTab[97008]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:145
		_go_fuzz_dep_.CoverTab[97009]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:145
		// _ = "end of CoverTab[97009]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:145
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:145
	// _ = "end of CoverTab[97004]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:145
	_go_fuzz_dep_.CoverTab[97005]++
											p.bypassZones = append(p.bypassZones, zone)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:146
	// _ = "end of CoverTab[97005]"
}

// AddHost specifies a host name that will use the bypass proxy.
func (p *PerHost) AddHost(host string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:150
	_go_fuzz_dep_.CoverTab[97010]++
											if strings.HasSuffix(host, ".") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:151
		_go_fuzz_dep_.CoverTab[97012]++
												host = host[:len(host)-1]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:152
		// _ = "end of CoverTab[97012]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:153
		_go_fuzz_dep_.CoverTab[97013]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:153
		// _ = "end of CoverTab[97013]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:153
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:153
	// _ = "end of CoverTab[97010]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:153
	_go_fuzz_dep_.CoverTab[97011]++
											p.bypassHosts = append(p.bypassHosts, host)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:154
	// _ = "end of CoverTab[97011]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:155
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/per_host.go:155
var _ = _go_fuzz_dep_.CoverTab
