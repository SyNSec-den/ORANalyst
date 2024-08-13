// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:5
// Package proxy provides support for a variety of protocols to proxy network
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:5
// data.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:7
package proxy

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:7
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:7
)

import (
	"errors"
	"net"
	"net/url"
	"os"
	"sync"
)

// A Dialer is a means to establish a connection.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:17
// Custom dialers should also implement ContextDialer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:19
type Dialer interface {
	// Dial connects to the given address via the proxy.
	Dial(network, addr string) (c net.Conn, err error)
}

// Auth contains authentication parameters that specific Dialers may require.
type Auth struct {
	User, Password string
}

// FromEnvironment returns the dialer specified by the proxy-related
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:29
// variables in the environment and makes underlying connections
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:29
// directly.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:32
func FromEnvironment() Dialer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:32
	_go_fuzz_dep_.CoverTab[97014]++
										return FromEnvironmentUsing(Direct)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:33
	// _ = "end of CoverTab[97014]"
}

// FromEnvironmentUsing returns the dialer specify by the proxy-related
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:36
// variables in the environment and makes underlying connections
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:36
// using the provided forwarding Dialer (for instance, a *net.Dialer
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:36
// with desired configuration).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:40
func FromEnvironmentUsing(forward Dialer) Dialer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:40
	_go_fuzz_dep_.CoverTab[97015]++
										allProxy := allProxyEnv.Get()
										if len(allProxy) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:42
		_go_fuzz_dep_.CoverTab[97020]++
											return forward
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:43
		// _ = "end of CoverTab[97020]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:44
		_go_fuzz_dep_.CoverTab[97021]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:44
		// _ = "end of CoverTab[97021]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:44
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:44
	// _ = "end of CoverTab[97015]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:44
	_go_fuzz_dep_.CoverTab[97016]++

										proxyURL, err := url.Parse(allProxy)
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:47
		_go_fuzz_dep_.CoverTab[97022]++
											return forward
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:48
		// _ = "end of CoverTab[97022]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:49
		_go_fuzz_dep_.CoverTab[97023]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:49
		// _ = "end of CoverTab[97023]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:49
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:49
	// _ = "end of CoverTab[97016]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:49
	_go_fuzz_dep_.CoverTab[97017]++
										proxy, err := FromURL(proxyURL, forward)
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:51
		_go_fuzz_dep_.CoverTab[97024]++
											return forward
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:52
		// _ = "end of CoverTab[97024]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:53
		_go_fuzz_dep_.CoverTab[97025]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:53
		// _ = "end of CoverTab[97025]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:53
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:53
	// _ = "end of CoverTab[97017]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:53
	_go_fuzz_dep_.CoverTab[97018]++

										noProxy := noProxyEnv.Get()
										if len(noProxy) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:56
		_go_fuzz_dep_.CoverTab[97026]++
											return proxy
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:57
		// _ = "end of CoverTab[97026]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:58
		_go_fuzz_dep_.CoverTab[97027]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:58
		// _ = "end of CoverTab[97027]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:58
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:58
	// _ = "end of CoverTab[97018]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:58
	_go_fuzz_dep_.CoverTab[97019]++

										perHost := NewPerHost(proxy, forward)
										perHost.AddFromString(noProxy)
										return perHost
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:62
	// _ = "end of CoverTab[97019]"
}

// proxySchemes is a map from URL schemes to a function that creates a Dialer
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:65
// from a URL with such a scheme.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:67
var proxySchemes map[string]func(*url.URL, Dialer) (Dialer, error)

// RegisterDialerType takes a URL scheme and a function to generate Dialers from
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:69
// a URL with that scheme and a forwarding Dialer. Registered schemes are used
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:69
// by FromURL.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:72
func RegisterDialerType(scheme string, f func(*url.URL, Dialer) (Dialer, error)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:72
	_go_fuzz_dep_.CoverTab[97028]++
										if proxySchemes == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:73
		_go_fuzz_dep_.CoverTab[97030]++
											proxySchemes = make(map[string]func(*url.URL, Dialer) (Dialer, error))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:74
		// _ = "end of CoverTab[97030]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:75
		_go_fuzz_dep_.CoverTab[97031]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:75
		// _ = "end of CoverTab[97031]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:75
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:75
	// _ = "end of CoverTab[97028]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:75
	_go_fuzz_dep_.CoverTab[97029]++
										proxySchemes[scheme] = f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:76
	// _ = "end of CoverTab[97029]"
}

// FromURL returns a Dialer given a URL specification and an underlying
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:79
// Dialer for it to make network requests.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:81
func FromURL(u *url.URL, forward Dialer) (Dialer, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:81
	_go_fuzz_dep_.CoverTab[97032]++
										var auth *Auth
										if u.User != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:83
		_go_fuzz_dep_.CoverTab[97036]++
											auth = new(Auth)
											auth.User = u.User.Username()
											if p, ok := u.User.Password(); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:86
			_go_fuzz_dep_.CoverTab[97037]++
												auth.Password = p
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:87
			// _ = "end of CoverTab[97037]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:88
			_go_fuzz_dep_.CoverTab[97038]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:88
			// _ = "end of CoverTab[97038]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:88
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:88
		// _ = "end of CoverTab[97036]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:89
		_go_fuzz_dep_.CoverTab[97039]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:89
		// _ = "end of CoverTab[97039]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:89
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:89
	// _ = "end of CoverTab[97032]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:89
	_go_fuzz_dep_.CoverTab[97033]++

										switch u.Scheme {
	case "socks5", "socks5h":
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:92
		_go_fuzz_dep_.CoverTab[97040]++
											addr := u.Hostname()
											port := u.Port()
											if port == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:95
			_go_fuzz_dep_.CoverTab[97043]++
												port = "1080"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:96
			// _ = "end of CoverTab[97043]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:97
			_go_fuzz_dep_.CoverTab[97044]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:97
			// _ = "end of CoverTab[97044]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:97
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:97
		// _ = "end of CoverTab[97040]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:97
		_go_fuzz_dep_.CoverTab[97041]++
											return SOCKS5("tcp", net.JoinHostPort(addr, port), auth, forward)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:98
		// _ = "end of CoverTab[97041]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:98
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:98
		_go_fuzz_dep_.CoverTab[97042]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:98
		// _ = "end of CoverTab[97042]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:99
	// _ = "end of CoverTab[97033]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:99
	_go_fuzz_dep_.CoverTab[97034]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:103
	if proxySchemes != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:103
		_go_fuzz_dep_.CoverTab[97045]++
											if f, ok := proxySchemes[u.Scheme]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:104
			_go_fuzz_dep_.CoverTab[97046]++
												return f(u, forward)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:105
			// _ = "end of CoverTab[97046]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:106
			_go_fuzz_dep_.CoverTab[97047]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:106
			// _ = "end of CoverTab[97047]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:106
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:106
		// _ = "end of CoverTab[97045]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:107
		_go_fuzz_dep_.CoverTab[97048]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:107
		// _ = "end of CoverTab[97048]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:107
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:107
	// _ = "end of CoverTab[97034]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:107
	_go_fuzz_dep_.CoverTab[97035]++

										return nil, errors.New("proxy: unknown scheme: " + u.Scheme)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:109
	// _ = "end of CoverTab[97035]"
}

var (
	allProxyEnv	= &envOnce{
		names: []string{"ALL_PROXY", "all_proxy"},
	}
	noProxyEnv	= &envOnce{
		names: []string{"NO_PROXY", "no_proxy"},
	}
)

// envOnce looks up an environment variable (optionally by multiple
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:121
// names) once. It mitigates expensive lookups on some platforms
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:121
// (e.g. Windows).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:121
// (Borrowed from net/http/transport.go)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:125
type envOnce struct {
	names	[]string
	once	sync.Once
	val	string
}

func (e *envOnce) Get() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:131
	_go_fuzz_dep_.CoverTab[97049]++
										e.once.Do(e.init)
										return e.val
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:133
	// _ = "end of CoverTab[97049]"
}

func (e *envOnce) init() {
	for _, n := range e.names {
		e.val = os.Getenv(n)
		if e.val != "" {
			return
		}
	}
}

// reset is used by tests
func (e *envOnce) reset() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:146
	_go_fuzz_dep_.CoverTab[97050]++
										e.once = sync.Once{}
										e.val = ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:148
	// _ = "end of CoverTab[97050]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:149
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/proxy.go:149
var _ = _go_fuzz_dep_.CoverTab
