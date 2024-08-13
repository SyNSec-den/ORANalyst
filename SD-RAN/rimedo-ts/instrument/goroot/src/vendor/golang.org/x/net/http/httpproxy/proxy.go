// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:5
// Package httpproxy provides support for HTTP proxy determination
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:5
// based on environment variables, as provided by net/http's
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:5
// ProxyFromEnvironment function.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:5
// The API is not subject to the Go 1 compatibility promise and may change at
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:5
// any time.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:11
package httpproxy

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:11
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:11
)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:11
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:11
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:11
)

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/idna"
)

// Config holds configuration for HTTP proxy settings. See
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:25
// FromEnvironment for details.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:27
type Config struct {
	// HTTPProxy represents the value of the HTTP_PROXY or
	// http_proxy environment variable. It will be used as the proxy
	// URL for HTTP requests unless overridden by NoProxy.
	HTTPProxy	string

	// HTTPSProxy represents the HTTPS_PROXY or https_proxy
	// environment variable. It will be used as the proxy URL for
	// HTTPS requests unless overridden by NoProxy.
	HTTPSProxy	string

	// NoProxy represents the NO_PROXY or no_proxy environment
	// variable. It specifies a string that contains comma-separated values
	// specifying hosts that should be excluded from proxying. Each value is
	// represented by an IP address prefix (1.2.3.4), an IP address prefix in
	// CIDR notation (1.2.3.4/8), a domain name, or a special DNS label (*).
	// An IP address prefix and domain name can also include a literal port
	// number (1.2.3.4:80).
	// A domain name matches that name and all subdomains. A domain name with
	// a leading "." matches subdomains only. For example "foo.com" matches
	// "foo.com" and "bar.foo.com"; ".y.com" matches "x.y.com" but not "y.com".
	// A single asterisk (*) indicates that no proxying should be done.
	// A best effort is made to parse the string and errors are
	// ignored.
	NoProxy	string

	// CGI holds whether the current process is running
	// as a CGI handler (FromEnvironment infers this from the
	// presence of a REQUEST_METHOD environment variable).
	// When this is set, ProxyForURL will return an error
	// when HTTPProxy applies, because a client could be
	// setting HTTP_PROXY maliciously. See https://golang.org/s/cgihttpproxy.
	CGI	bool
}

// config holds the parsed configuration for HTTP proxy settings.
type config struct {
	// Config represents the original configuration as defined above.
	Config

	// httpsProxy is the parsed URL of the HTTPSProxy if defined.
	httpsProxy	*url.URL

	// httpProxy is the parsed URL of the HTTPProxy if defined.
	httpProxy	*url.URL

	// ipMatchers represent all values in the NoProxy that are IP address
	// prefixes or an IP address in CIDR notation.
	ipMatchers	[]matcher

	// domainMatchers represent all values in the NoProxy that are a domain
	// name or hostname & domain name
	domainMatchers	[]matcher
}

// FromEnvironment returns a Config instance populated from the
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:82
// environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:82
// lowercase versions thereof).
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:82
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:82
// The environment values may be either a complete URL or a
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:82
// "host[:port]", in which case the "http" scheme is assumed. An error
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:82
// is returned if the value is a different form.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:89
func FromEnvironment() *Config {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:89
	_go_fuzz_dep_.CoverTab[34994]++
										return &Config{
		HTTPProxy:	getEnvAny("HTTP_PROXY", "http_proxy"),
		HTTPSProxy:	getEnvAny("HTTPS_PROXY", "https_proxy"),
		NoProxy:	getEnvAny("NO_PROXY", "no_proxy"),
		CGI:		os.Getenv("REQUEST_METHOD") != "",
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:95
	// _ = "end of CoverTab[34994]"
}

func getEnvAny(names ...string) string {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:98
	_go_fuzz_dep_.CoverTab[34995]++
										for _, n := range names {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:99
		_go_fuzz_dep_.CoverTab[34997]++
											if val := os.Getenv(n); val != "" {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:100
			_go_fuzz_dep_.CoverTab[34998]++
												return val
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:101
			// _ = "end of CoverTab[34998]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:102
			_go_fuzz_dep_.CoverTab[34999]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:102
			// _ = "end of CoverTab[34999]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:102
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:102
		// _ = "end of CoverTab[34997]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:103
	// _ = "end of CoverTab[34995]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:103
	_go_fuzz_dep_.CoverTab[34996]++
										return ""
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:104
	// _ = "end of CoverTab[34996]"
}

// ProxyFunc returns a function that determines the proxy URL to use for
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:107
// a given request URL. Changing the contents of cfg will not affect
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:107
// proxy functions created earlier.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:107
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:107
// A nil URL and nil error are returned if no proxy is defined in the
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:107
// environment, or a proxy should not be used for the given request, as
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:107
// defined by NO_PROXY.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:107
//
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:107
// As a special case, if req.URL.Host is "localhost" or a loopback address
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:107
// (with or without a port number), then a nil URL and nil error will be returned.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:117
func (cfg *Config) ProxyFunc() func(reqURL *url.URL) (*url.URL, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:117
	_go_fuzz_dep_.CoverTab[35000]++

										cfg1 := &config{
		Config: *cfg,
	}
										cfg1.init()
										return cfg1.proxyForURL
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:123
	// _ = "end of CoverTab[35000]"
}

func (cfg *config) proxyForURL(reqURL *url.URL) (*url.URL, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:126
	_go_fuzz_dep_.CoverTab[35001]++
										var proxy *url.URL
										if reqURL.Scheme == "https" {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:128
		_go_fuzz_dep_.CoverTab[35005]++
											proxy = cfg.httpsProxy
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:129
		// _ = "end of CoverTab[35005]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:130
		_go_fuzz_dep_.CoverTab[35006]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:130
		if reqURL.Scheme == "http" {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:130
			_go_fuzz_dep_.CoverTab[35007]++
												proxy = cfg.httpProxy
												if proxy != nil && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:132
				_go_fuzz_dep_.CoverTab[35008]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:132
				return cfg.CGI
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:132
				// _ = "end of CoverTab[35008]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:132
			}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:132
				_go_fuzz_dep_.CoverTab[35009]++
													return nil, errors.New("refusing to use HTTP_PROXY value in CGI environment; see golang.org/s/cgihttpproxy")
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:133
				// _ = "end of CoverTab[35009]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:134
				_go_fuzz_dep_.CoverTab[35010]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:134
				// _ = "end of CoverTab[35010]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:134
			}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:134
			// _ = "end of CoverTab[35007]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:135
			_go_fuzz_dep_.CoverTab[35011]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:135
			// _ = "end of CoverTab[35011]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:135
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:135
		// _ = "end of CoverTab[35006]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:135
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:135
	// _ = "end of CoverTab[35001]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:135
	_go_fuzz_dep_.CoverTab[35002]++
										if proxy == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:136
		_go_fuzz_dep_.CoverTab[35012]++
											return nil, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:137
		// _ = "end of CoverTab[35012]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:138
		_go_fuzz_dep_.CoverTab[35013]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:138
		// _ = "end of CoverTab[35013]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:138
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:138
	// _ = "end of CoverTab[35002]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:138
	_go_fuzz_dep_.CoverTab[35003]++
										if !cfg.useProxy(canonicalAddr(reqURL)) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:139
		_go_fuzz_dep_.CoverTab[35014]++
											return nil, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:140
		// _ = "end of CoverTab[35014]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:141
		_go_fuzz_dep_.CoverTab[35015]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:141
		// _ = "end of CoverTab[35015]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:141
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:141
	// _ = "end of CoverTab[35003]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:141
	_go_fuzz_dep_.CoverTab[35004]++

										return proxy, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:143
	// _ = "end of CoverTab[35004]"
}

func parseProxy(proxy string) (*url.URL, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:146
	_go_fuzz_dep_.CoverTab[35016]++
										if proxy == "" {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:147
		_go_fuzz_dep_.CoverTab[35020]++
											return nil, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:148
		// _ = "end of CoverTab[35020]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:149
		_go_fuzz_dep_.CoverTab[35021]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:149
		// _ = "end of CoverTab[35021]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:149
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:149
	// _ = "end of CoverTab[35016]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:149
	_go_fuzz_dep_.CoverTab[35017]++

										proxyURL, err := url.Parse(proxy)
										if err != nil || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:152
		_go_fuzz_dep_.CoverTab[35022]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:152
		return (proxyURL.Scheme != "http" && func() bool {
												_go_fuzz_dep_.CoverTab[35023]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:153
			return proxyURL.Scheme != "https"
												// _ = "end of CoverTab[35023]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:154
		}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:154
			_go_fuzz_dep_.CoverTab[35024]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:154
			return proxyURL.Scheme != "socks5"
												// _ = "end of CoverTab[35024]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:155
		}())
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:155
		// _ = "end of CoverTab[35022]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:155
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:155
		_go_fuzz_dep_.CoverTab[35025]++

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:159
		if proxyURL, err := url.Parse("http://" + proxy); err == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:159
			_go_fuzz_dep_.CoverTab[35026]++
												return proxyURL, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:160
			// _ = "end of CoverTab[35026]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:161
			_go_fuzz_dep_.CoverTab[35027]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:161
			// _ = "end of CoverTab[35027]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:161
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:161
		// _ = "end of CoverTab[35025]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:162
		_go_fuzz_dep_.CoverTab[35028]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:162
		// _ = "end of CoverTab[35028]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:162
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:162
	// _ = "end of CoverTab[35017]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:162
	_go_fuzz_dep_.CoverTab[35018]++
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:163
		_go_fuzz_dep_.CoverTab[35029]++
											return nil, fmt.Errorf("invalid proxy address %q: %v", proxy, err)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:164
		// _ = "end of CoverTab[35029]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:165
		_go_fuzz_dep_.CoverTab[35030]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:165
		// _ = "end of CoverTab[35030]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:165
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:165
	// _ = "end of CoverTab[35018]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:165
	_go_fuzz_dep_.CoverTab[35019]++
										return proxyURL, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:166
	// _ = "end of CoverTab[35019]"
}

// useProxy reports whether requests to addr should use a proxy,
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:169
// according to the NO_PROXY or no_proxy environment variable.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:169
// addr is always a canonicalAddr with a host and port.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:172
func (cfg *config) useProxy(addr string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:172
	_go_fuzz_dep_.CoverTab[35031]++
										if len(addr) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:173
		_go_fuzz_dep_.CoverTab[35038]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:174
		// _ = "end of CoverTab[35038]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:175
		_go_fuzz_dep_.CoverTab[35039]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:175
		// _ = "end of CoverTab[35039]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:175
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:175
	// _ = "end of CoverTab[35031]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:175
	_go_fuzz_dep_.CoverTab[35032]++
										host, port, err := net.SplitHostPort(addr)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:177
		_go_fuzz_dep_.CoverTab[35040]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:178
		// _ = "end of CoverTab[35040]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:179
		_go_fuzz_dep_.CoverTab[35041]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:179
		// _ = "end of CoverTab[35041]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:179
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:179
	// _ = "end of CoverTab[35032]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:179
	_go_fuzz_dep_.CoverTab[35033]++
										if host == "localhost" {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:180
		_go_fuzz_dep_.CoverTab[35042]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:181
		// _ = "end of CoverTab[35042]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:182
		_go_fuzz_dep_.CoverTab[35043]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:182
		// _ = "end of CoverTab[35043]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:182
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:182
	// _ = "end of CoverTab[35033]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:182
	_go_fuzz_dep_.CoverTab[35034]++
										ip := net.ParseIP(host)
										if ip != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:184
		_go_fuzz_dep_.CoverTab[35044]++
											if ip.IsLoopback() {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:185
			_go_fuzz_dep_.CoverTab[35045]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:186
			// _ = "end of CoverTab[35045]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:187
			_go_fuzz_dep_.CoverTab[35046]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:187
			// _ = "end of CoverTab[35046]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:187
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:187
		// _ = "end of CoverTab[35044]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:188
		_go_fuzz_dep_.CoverTab[35047]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:188
		// _ = "end of CoverTab[35047]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:188
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:188
	// _ = "end of CoverTab[35034]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:188
	_go_fuzz_dep_.CoverTab[35035]++

										addr = strings.ToLower(strings.TrimSpace(host))

										if ip != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:192
		_go_fuzz_dep_.CoverTab[35048]++
											for _, m := range cfg.ipMatchers {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:193
			_go_fuzz_dep_.CoverTab[35049]++
												if m.match(addr, port, ip) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:194
				_go_fuzz_dep_.CoverTab[35050]++
													return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:195
				// _ = "end of CoverTab[35050]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:196
				_go_fuzz_dep_.CoverTab[35051]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:196
				// _ = "end of CoverTab[35051]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:196
			}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:196
			// _ = "end of CoverTab[35049]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:197
		// _ = "end of CoverTab[35048]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:198
		_go_fuzz_dep_.CoverTab[35052]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:198
		// _ = "end of CoverTab[35052]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:198
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:198
	// _ = "end of CoverTab[35035]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:198
	_go_fuzz_dep_.CoverTab[35036]++
										for _, m := range cfg.domainMatchers {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:199
		_go_fuzz_dep_.CoverTab[35053]++
											if m.match(addr, port, ip) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:200
			_go_fuzz_dep_.CoverTab[35054]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:201
			// _ = "end of CoverTab[35054]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:202
			_go_fuzz_dep_.CoverTab[35055]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:202
			// _ = "end of CoverTab[35055]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:202
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:202
		// _ = "end of CoverTab[35053]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:203
	// _ = "end of CoverTab[35036]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:203
	_go_fuzz_dep_.CoverTab[35037]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:204
	// _ = "end of CoverTab[35037]"
}

func (c *config) init() {
	if parsed, err := parseProxy(c.HTTPProxy); err == nil {
		c.httpProxy = parsed
	}
	if parsed, err := parseProxy(c.HTTPSProxy); err == nil {
		c.httpsProxy = parsed
	}

	for _, p := range strings.Split(c.NoProxy, ",") {
		p = strings.ToLower(strings.TrimSpace(p))
		if len(p) == 0 {
			continue
		}

		if p == "*" {
			c.ipMatchers = []matcher{allMatch{}}
			c.domainMatchers = []matcher{allMatch{}}
			return
		}

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:228
		if _, pnet, err := net.ParseCIDR(p); err == nil {
			c.ipMatchers = append(c.ipMatchers, cidrMatch{cidr: pnet})
			continue
		}

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:234
		phost, pport, err := net.SplitHostPort(p)
		if err == nil {
			if len(phost) == 0 {

				continue
			}
			if phost[0] == '[' && phost[len(phost)-1] == ']' {
				phost = phost[1 : len(phost)-1]
			}
		} else {
			phost = p
		}

		if pip := net.ParseIP(phost); pip != nil {
			c.ipMatchers = append(c.ipMatchers, ipMatch{ip: pip, port: pport})
			continue
		}

		if len(phost) == 0 {

			continue
		}

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:261
		if strings.HasPrefix(phost, "*.") {
			phost = phost[1:]
		}
		matchHost := false
		if phost[0] != '.' {
			matchHost = true
			phost = "." + phost
		}
		if v, err := idnaASCII(phost); err == nil {
			phost = v
		}
		c.domainMatchers = append(c.domainMatchers, domainMatch{host: phost, port: pport, matchHost: matchHost})
	}
}

var portMap = map[string]string{
	"http":		"80",
	"https":	"443",
	"socks5":	"1080",
}

// canonicalAddr returns url.Host but always with a ":port" suffix
func canonicalAddr(url *url.URL) string {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:283
	_go_fuzz_dep_.CoverTab[35056]++
										addr := url.Hostname()
										if v, err := idnaASCII(addr); err == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:285
		_go_fuzz_dep_.CoverTab[35059]++
											addr = v
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:286
		// _ = "end of CoverTab[35059]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:287
		_go_fuzz_dep_.CoverTab[35060]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:287
		// _ = "end of CoverTab[35060]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:287
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:287
	// _ = "end of CoverTab[35056]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:287
	_go_fuzz_dep_.CoverTab[35057]++
										port := url.Port()
										if port == "" {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:289
		_go_fuzz_dep_.CoverTab[35061]++
											port = portMap[url.Scheme]
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:290
		// _ = "end of CoverTab[35061]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:291
		_go_fuzz_dep_.CoverTab[35062]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:291
		// _ = "end of CoverTab[35062]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:291
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:291
	// _ = "end of CoverTab[35057]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:291
	_go_fuzz_dep_.CoverTab[35058]++
										return net.JoinHostPort(addr, port)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:292
	// _ = "end of CoverTab[35058]"
}

// Given a string of the form "host", "host:port", or "[ipv6::address]:port",
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:295
// return true if the string includes a port.
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:297
func hasPort(s string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:297
	_go_fuzz_dep_.CoverTab[35063]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:297
	return strings.LastIndex(s, ":") > strings.LastIndex(s, "]")
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:297
	// _ = "end of CoverTab[35063]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:297
}

func idnaASCII(v string) (string, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:299
	_go_fuzz_dep_.CoverTab[35064]++

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:309
	if isASCII(v) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:309
		_go_fuzz_dep_.CoverTab[35066]++
											return v, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:310
		// _ = "end of CoverTab[35066]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:311
		_go_fuzz_dep_.CoverTab[35067]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:311
		// _ = "end of CoverTab[35067]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:311
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:311
	// _ = "end of CoverTab[35064]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:311
	_go_fuzz_dep_.CoverTab[35065]++
										return idna.Lookup.ToASCII(v)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:312
	// _ = "end of CoverTab[35065]"
}

func isASCII(s string) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:315
	_go_fuzz_dep_.CoverTab[35068]++
										for i := 0; i < len(s); i++ {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:316
		_go_fuzz_dep_.CoverTab[35070]++
											if s[i] >= utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:317
			_go_fuzz_dep_.CoverTab[35071]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:318
			// _ = "end of CoverTab[35071]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:319
			_go_fuzz_dep_.CoverTab[35072]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:319
			// _ = "end of CoverTab[35072]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:319
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:319
		// _ = "end of CoverTab[35070]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:320
	// _ = "end of CoverTab[35068]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:320
	_go_fuzz_dep_.CoverTab[35069]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:321
	// _ = "end of CoverTab[35069]"
}

// matcher represents the matching rule for a given value in the NO_PROXY list
type matcher interface {
	// match returns true if the host and optional port or ip and optional port
	// are allowed
	match(host, port string, ip net.IP) bool
}

// allMatch matches on all possible inputs
type allMatch struct{}

func (a allMatch) match(host, port string, ip net.IP) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:334
	_go_fuzz_dep_.CoverTab[35073]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:335
	// _ = "end of CoverTab[35073]"
}

type cidrMatch struct {
	cidr *net.IPNet
}

func (m cidrMatch) match(host, port string, ip net.IP) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:342
	_go_fuzz_dep_.CoverTab[35074]++
										return m.cidr.Contains(ip)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:343
	// _ = "end of CoverTab[35074]"
}

type ipMatch struct {
	ip	net.IP
	port	string
}

func (m ipMatch) match(host, port string, ip net.IP) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:351
	_go_fuzz_dep_.CoverTab[35075]++
										if m.ip.Equal(ip) {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:352
		_go_fuzz_dep_.CoverTab[35077]++
											return m.port == "" || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:353
			_go_fuzz_dep_.CoverTab[35078]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:353
			return m.port == port
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:353
			// _ = "end of CoverTab[35078]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:353
		}()
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:353
		// _ = "end of CoverTab[35077]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:354
		_go_fuzz_dep_.CoverTab[35079]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:354
		// _ = "end of CoverTab[35079]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:354
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:354
	// _ = "end of CoverTab[35075]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:354
	_go_fuzz_dep_.CoverTab[35076]++
										return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:355
	// _ = "end of CoverTab[35076]"
}

type domainMatch struct {
	host	string
	port	string

	matchHost	bool
}

func (m domainMatch) match(host, port string, ip net.IP) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:365
	_go_fuzz_dep_.CoverTab[35080]++
										if strings.HasSuffix(host, m.host) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:366
		_go_fuzz_dep_.CoverTab[35082]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:366
		return (m.matchHost && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:366
			_go_fuzz_dep_.CoverTab[35083]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:366
			return host == m.host[1:]
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:366
			// _ = "end of CoverTab[35083]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:366
		}())
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:366
		// _ = "end of CoverTab[35082]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:366
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:366
		_go_fuzz_dep_.CoverTab[35084]++
											return m.port == "" || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:367
			_go_fuzz_dep_.CoverTab[35085]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:367
			return m.port == port
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:367
			// _ = "end of CoverTab[35085]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:367
		}()
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:367
		// _ = "end of CoverTab[35084]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:368
		_go_fuzz_dep_.CoverTab[35086]++
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:368
		// _ = "end of CoverTab[35086]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:368
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:368
	// _ = "end of CoverTab[35080]"
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:368
	_go_fuzz_dep_.CoverTab[35081]++
										return false
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:369
	// _ = "end of CoverTab[35081]"
}

//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:370
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/net/http/httpproxy/proxy.go:370
var _ = _go_fuzz_dep_.CoverTab
