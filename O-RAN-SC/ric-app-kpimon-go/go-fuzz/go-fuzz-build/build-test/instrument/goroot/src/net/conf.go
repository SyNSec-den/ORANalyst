// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

//line /snap/go/10455/src/net/conf.go:7
package net

//line /snap/go/10455/src/net/conf.go:7
import (
//line /snap/go/10455/src/net/conf.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/conf.go:7
)
//line /snap/go/10455/src/net/conf.go:7
import (
//line /snap/go/10455/src/net/conf.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/conf.go:7
)

import (
	"errors"
	"internal/bytealg"
	"internal/godebug"
	"io/fs"
	"os"
	"runtime"
	"sync"
	"syscall"
)

//line /snap/go/10455/src/net/conf.go:52
// conf is used to determine name resolution configuration.
type conf struct {
	netGo	bool	// prefer go approach, based on build tag and GODEBUG
	netCgo	bool	// prefer cgo approach, based on build tag and GODEBUG

	dnsDebugLevel	int	// from GODEBUG

	preferCgo	bool	// if no explicit preference, use cgo

	goos		string		// copy of runtime.GOOS, used for testing
	mdnsTest	mdnsTest	// assume /etc/mdns.allow exists, for testing
}

// mdnsTest is for testing only.
type mdnsTest int

const (
	mdnsFromSystem	mdnsTest	= iota
	mdnsAssumeExists
	mdnsAssumeDoesNotExist
)

var (
	confOnce	sync.Once	// guards init of confVal via initConfVal
	confVal		= &conf{goos: runtime.GOOS}
)

// systemConf returns the machine's network configuration.
func systemConf() *conf {
//line /snap/go/10455/src/net/conf.go:80
	_go_fuzz_dep_.CoverTab[4738]++
						confOnce.Do(initConfVal)
						return confVal
//line /snap/go/10455/src/net/conf.go:82
	// _ = "end of CoverTab[4738]"
}

// initConfVal initializes confVal based on the environment
//line /snap/go/10455/src/net/conf.go:85
// that will not change during program execution.
//line /snap/go/10455/src/net/conf.go:87
func initConfVal() {
//line /snap/go/10455/src/net/conf.go:87
	_go_fuzz_dep_.CoverTab[4739]++
						dnsMode, debugLevel := goDebugNetDNS()
						confVal.netGo = netGoBuildTag || func() bool {
//line /snap/go/10455/src/net/conf.go:89
		_go_fuzz_dep_.CoverTab[4745]++
//line /snap/go/10455/src/net/conf.go:89
		return dnsMode == "go"
//line /snap/go/10455/src/net/conf.go:89
		// _ = "end of CoverTab[4745]"
//line /snap/go/10455/src/net/conf.go:89
	}()
						confVal.netCgo = netCgoBuildTag || func() bool {
//line /snap/go/10455/src/net/conf.go:90
		_go_fuzz_dep_.CoverTab[4746]++
//line /snap/go/10455/src/net/conf.go:90
		return dnsMode == "cgo"
//line /snap/go/10455/src/net/conf.go:90
		// _ = "end of CoverTab[4746]"
//line /snap/go/10455/src/net/conf.go:90
	}()
						confVal.dnsDebugLevel = debugLevel

						if confVal.dnsDebugLevel > 0 {
//line /snap/go/10455/src/net/conf.go:93
		_go_fuzz_dep_.CoverTab[527549]++
//line /snap/go/10455/src/net/conf.go:93
		_go_fuzz_dep_.CoverTab[4747]++
							defer func() {
//line /snap/go/10455/src/net/conf.go:94
			_go_fuzz_dep_.CoverTab[4748]++
								if confVal.dnsDebugLevel > 1 {
//line /snap/go/10455/src/net/conf.go:95
				_go_fuzz_dep_.CoverTab[527551]++
//line /snap/go/10455/src/net/conf.go:95
				_go_fuzz_dep_.CoverTab[4750]++
									println("go package net: confVal.netCgo =", confVal.netCgo, " netGo =", confVal.netGo)
//line /snap/go/10455/src/net/conf.go:96
				// _ = "end of CoverTab[4750]"
			} else {
//line /snap/go/10455/src/net/conf.go:97
				_go_fuzz_dep_.CoverTab[527552]++
//line /snap/go/10455/src/net/conf.go:97
				_go_fuzz_dep_.CoverTab[4751]++
//line /snap/go/10455/src/net/conf.go:97
				// _ = "end of CoverTab[4751]"
//line /snap/go/10455/src/net/conf.go:97
			}
//line /snap/go/10455/src/net/conf.go:97
			// _ = "end of CoverTab[4748]"
//line /snap/go/10455/src/net/conf.go:97
			_go_fuzz_dep_.CoverTab[4749]++
								switch {
			case confVal.netGo:
//line /snap/go/10455/src/net/conf.go:99
				_go_fuzz_dep_.CoverTab[527553]++
//line /snap/go/10455/src/net/conf.go:99
				_go_fuzz_dep_.CoverTab[4752]++
									if netGoBuildTag {
//line /snap/go/10455/src/net/conf.go:100
					_go_fuzz_dep_.CoverTab[527557]++
//line /snap/go/10455/src/net/conf.go:100
					_go_fuzz_dep_.CoverTab[4756]++
										println("go package net: built with netgo build tag; using Go's DNS resolver")
//line /snap/go/10455/src/net/conf.go:101
					// _ = "end of CoverTab[4756]"
				} else {
//line /snap/go/10455/src/net/conf.go:102
					_go_fuzz_dep_.CoverTab[527558]++
//line /snap/go/10455/src/net/conf.go:102
					_go_fuzz_dep_.CoverTab[4757]++
										println("go package net: GODEBUG setting forcing use of Go's resolver")
//line /snap/go/10455/src/net/conf.go:103
					// _ = "end of CoverTab[4757]"
				}
//line /snap/go/10455/src/net/conf.go:104
				// _ = "end of CoverTab[4752]"
			case !cgoAvailable:
//line /snap/go/10455/src/net/conf.go:105
				_go_fuzz_dep_.CoverTab[527554]++
//line /snap/go/10455/src/net/conf.go:105
				_go_fuzz_dep_.CoverTab[4753]++
									println("go package net: cgo resolver not supported; using Go's DNS resolver")
//line /snap/go/10455/src/net/conf.go:106
				// _ = "end of CoverTab[4753]"
			case confVal.netCgo || func() bool {
//line /snap/go/10455/src/net/conf.go:107
				_go_fuzz_dep_.CoverTab[4758]++
//line /snap/go/10455/src/net/conf.go:107
				return confVal.preferCgo
//line /snap/go/10455/src/net/conf.go:107
				// _ = "end of CoverTab[4758]"
//line /snap/go/10455/src/net/conf.go:107
			}():
//line /snap/go/10455/src/net/conf.go:107
				_go_fuzz_dep_.CoverTab[527555]++
//line /snap/go/10455/src/net/conf.go:107
				_go_fuzz_dep_.CoverTab[4754]++
									println("go package net: using cgo DNS resolver")
//line /snap/go/10455/src/net/conf.go:108
				// _ = "end of CoverTab[4754]"
			default:
//line /snap/go/10455/src/net/conf.go:109
				_go_fuzz_dep_.CoverTab[527556]++
//line /snap/go/10455/src/net/conf.go:109
				_go_fuzz_dep_.CoverTab[4755]++
									println("go package net: dynamic selection of DNS resolver")
//line /snap/go/10455/src/net/conf.go:110
				// _ = "end of CoverTab[4755]"
			}
//line /snap/go/10455/src/net/conf.go:111
			// _ = "end of CoverTab[4749]"
		}()
//line /snap/go/10455/src/net/conf.go:112
		// _ = "end of CoverTab[4747]"
	} else {
//line /snap/go/10455/src/net/conf.go:113
		_go_fuzz_dep_.CoverTab[527550]++
//line /snap/go/10455/src/net/conf.go:113
		_go_fuzz_dep_.CoverTab[4759]++
//line /snap/go/10455/src/net/conf.go:113
		// _ = "end of CoverTab[4759]"
//line /snap/go/10455/src/net/conf.go:113
	}
//line /snap/go/10455/src/net/conf.go:113
	// _ = "end of CoverTab[4739]"
//line /snap/go/10455/src/net/conf.go:113
	_go_fuzz_dep_.CoverTab[4740]++

//line /snap/go/10455/src/net/conf.go:119
	confVal.preferCgo = false

//line /snap/go/10455/src/net/conf.go:122
	if !cgoAvailable {
//line /snap/go/10455/src/net/conf.go:122
		_go_fuzz_dep_.CoverTab[527559]++
//line /snap/go/10455/src/net/conf.go:122
		_go_fuzz_dep_.CoverTab[4760]++
							return
//line /snap/go/10455/src/net/conf.go:123
		// _ = "end of CoverTab[4760]"
	} else {
//line /snap/go/10455/src/net/conf.go:124
		_go_fuzz_dep_.CoverTab[527560]++
//line /snap/go/10455/src/net/conf.go:124
		_go_fuzz_dep_.CoverTab[4761]++
//line /snap/go/10455/src/net/conf.go:124
		// _ = "end of CoverTab[4761]"
//line /snap/go/10455/src/net/conf.go:124
	}
//line /snap/go/10455/src/net/conf.go:124
	// _ = "end of CoverTab[4740]"
//line /snap/go/10455/src/net/conf.go:124
	_go_fuzz_dep_.CoverTab[4741]++

//line /snap/go/10455/src/net/conf.go:127
	if goosPrefersCgo() {
//line /snap/go/10455/src/net/conf.go:127
		_go_fuzz_dep_.CoverTab[527561]++
//line /snap/go/10455/src/net/conf.go:127
		_go_fuzz_dep_.CoverTab[4762]++
							confVal.preferCgo = true
							return
//line /snap/go/10455/src/net/conf.go:129
		// _ = "end of CoverTab[4762]"
	} else {
//line /snap/go/10455/src/net/conf.go:130
		_go_fuzz_dep_.CoverTab[527562]++
//line /snap/go/10455/src/net/conf.go:130
		_go_fuzz_dep_.CoverTab[4763]++
//line /snap/go/10455/src/net/conf.go:130
		// _ = "end of CoverTab[4763]"
//line /snap/go/10455/src/net/conf.go:130
	}
//line /snap/go/10455/src/net/conf.go:130
	// _ = "end of CoverTab[4741]"
//line /snap/go/10455/src/net/conf.go:130
	_go_fuzz_dep_.CoverTab[4742]++

//line /snap/go/10455/src/net/conf.go:133
	switch runtime.GOOS {
	case "plan9", "windows", "js", "wasip1":
//line /snap/go/10455/src/net/conf.go:134
		_go_fuzz_dep_.CoverTab[527563]++
//line /snap/go/10455/src/net/conf.go:134
		_go_fuzz_dep_.CoverTab[4764]++
							return
//line /snap/go/10455/src/net/conf.go:135
		// _ = "end of CoverTab[4764]"
//line /snap/go/10455/src/net/conf.go:135
	default:
//line /snap/go/10455/src/net/conf.go:135
		_go_fuzz_dep_.CoverTab[527564]++
//line /snap/go/10455/src/net/conf.go:135
		_go_fuzz_dep_.CoverTab[4765]++
//line /snap/go/10455/src/net/conf.go:135
		// _ = "end of CoverTab[4765]"
	}
//line /snap/go/10455/src/net/conf.go:136
	// _ = "end of CoverTab[4742]"
//line /snap/go/10455/src/net/conf.go:136
	_go_fuzz_dep_.CoverTab[4743]++

//line /snap/go/10455/src/net/conf.go:142
	_, localDomainDefined := syscall.Getenv("LOCALDOMAIN")
	if localDomainDefined || func() bool {
//line /snap/go/10455/src/net/conf.go:143
		_go_fuzz_dep_.CoverTab[4766]++
//line /snap/go/10455/src/net/conf.go:143
		return os.Getenv("RES_OPTIONS") != ""
//line /snap/go/10455/src/net/conf.go:143
		// _ = "end of CoverTab[4766]"
//line /snap/go/10455/src/net/conf.go:143
	}() || func() bool {
//line /snap/go/10455/src/net/conf.go:143
		_go_fuzz_dep_.CoverTab[4767]++
//line /snap/go/10455/src/net/conf.go:143
		return os.Getenv("HOSTALIASES") != ""
//line /snap/go/10455/src/net/conf.go:143
		// _ = "end of CoverTab[4767]"
//line /snap/go/10455/src/net/conf.go:143
	}() {
//line /snap/go/10455/src/net/conf.go:143
		_go_fuzz_dep_.CoverTab[527565]++
//line /snap/go/10455/src/net/conf.go:143
		_go_fuzz_dep_.CoverTab[4768]++
							confVal.preferCgo = true
							return
//line /snap/go/10455/src/net/conf.go:145
		// _ = "end of CoverTab[4768]"
	} else {
//line /snap/go/10455/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[527566]++
//line /snap/go/10455/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[4769]++
//line /snap/go/10455/src/net/conf.go:146
		// _ = "end of CoverTab[4769]"
//line /snap/go/10455/src/net/conf.go:146
	}
//line /snap/go/10455/src/net/conf.go:146
	// _ = "end of CoverTab[4743]"
//line /snap/go/10455/src/net/conf.go:146
	_go_fuzz_dep_.CoverTab[4744]++

//line /snap/go/10455/src/net/conf.go:150
	if runtime.GOOS == "openbsd" && func() bool {
//line /snap/go/10455/src/net/conf.go:150
		_go_fuzz_dep_.CoverTab[4770]++
//line /snap/go/10455/src/net/conf.go:150
		return os.Getenv("ASR_CONFIG") != ""
//line /snap/go/10455/src/net/conf.go:150
		// _ = "end of CoverTab[4770]"
//line /snap/go/10455/src/net/conf.go:150
	}() {
//line /snap/go/10455/src/net/conf.go:150
		_go_fuzz_dep_.CoverTab[527567]++
//line /snap/go/10455/src/net/conf.go:150
		_go_fuzz_dep_.CoverTab[4771]++
							confVal.preferCgo = true
							return
//line /snap/go/10455/src/net/conf.go:152
		// _ = "end of CoverTab[4771]"
	} else {
//line /snap/go/10455/src/net/conf.go:153
		_go_fuzz_dep_.CoverTab[527568]++
//line /snap/go/10455/src/net/conf.go:153
		_go_fuzz_dep_.CoverTab[4772]++
//line /snap/go/10455/src/net/conf.go:153
		// _ = "end of CoverTab[4772]"
//line /snap/go/10455/src/net/conf.go:153
	}
//line /snap/go/10455/src/net/conf.go:153
	// _ = "end of CoverTab[4744]"
}

// goosPreferCgo reports whether the GOOS value passed in prefers
//line /snap/go/10455/src/net/conf.go:156
// the cgo resolver.
//line /snap/go/10455/src/net/conf.go:158
func goosPrefersCgo() bool {
//line /snap/go/10455/src/net/conf.go:158
	_go_fuzz_dep_.CoverTab[4773]++
						switch runtime.GOOS {

//line /snap/go/10455/src/net/conf.go:166
	case "windows", "plan9":
//line /snap/go/10455/src/net/conf.go:166
		_go_fuzz_dep_.CoverTab[527569]++
//line /snap/go/10455/src/net/conf.go:166
		_go_fuzz_dep_.CoverTab[4774]++
							return true
//line /snap/go/10455/src/net/conf.go:167
		// _ = "end of CoverTab[4774]"

//line /snap/go/10455/src/net/conf.go:171
	case "darwin", "ios":
//line /snap/go/10455/src/net/conf.go:171
		_go_fuzz_dep_.CoverTab[527570]++
//line /snap/go/10455/src/net/conf.go:171
		_go_fuzz_dep_.CoverTab[4775]++
							return true
//line /snap/go/10455/src/net/conf.go:172
		// _ = "end of CoverTab[4775]"

//line /snap/go/10455/src/net/conf.go:176
	case "android":
//line /snap/go/10455/src/net/conf.go:176
		_go_fuzz_dep_.CoverTab[527571]++
//line /snap/go/10455/src/net/conf.go:176
		_go_fuzz_dep_.CoverTab[4776]++
							return true
//line /snap/go/10455/src/net/conf.go:177
		// _ = "end of CoverTab[4776]"

	default:
//line /snap/go/10455/src/net/conf.go:179
		_go_fuzz_dep_.CoverTab[527572]++
//line /snap/go/10455/src/net/conf.go:179
		_go_fuzz_dep_.CoverTab[4777]++
							return false
//line /snap/go/10455/src/net/conf.go:180
		// _ = "end of CoverTab[4777]"
	}
//line /snap/go/10455/src/net/conf.go:181
	// _ = "end of CoverTab[4773]"
}

// mustUseGoResolver reports whether a DNS lookup of any sort is
//line /snap/go/10455/src/net/conf.go:184
// required to use the go resolver. The provided Resolver is optional.
//line /snap/go/10455/src/net/conf.go:184
// This will report true if the cgo resolver is not available.
//line /snap/go/10455/src/net/conf.go:187
func (c *conf) mustUseGoResolver(r *Resolver) bool {
//line /snap/go/10455/src/net/conf.go:187
	_go_fuzz_dep_.CoverTab[4778]++
						return c.netGo || func() bool {
//line /snap/go/10455/src/net/conf.go:188
		_go_fuzz_dep_.CoverTab[4779]++
//line /snap/go/10455/src/net/conf.go:188
		return r.preferGo()
//line /snap/go/10455/src/net/conf.go:188
		// _ = "end of CoverTab[4779]"
//line /snap/go/10455/src/net/conf.go:188
	}() || func() bool {
//line /snap/go/10455/src/net/conf.go:188
		_go_fuzz_dep_.CoverTab[4780]++
//line /snap/go/10455/src/net/conf.go:188
		return !cgoAvailable
//line /snap/go/10455/src/net/conf.go:188
		// _ = "end of CoverTab[4780]"
//line /snap/go/10455/src/net/conf.go:188
	}()
//line /snap/go/10455/src/net/conf.go:188
	// _ = "end of CoverTab[4778]"
}

// addrLookupOrder determines which strategy to use to resolve addresses.
//line /snap/go/10455/src/net/conf.go:191
// The provided Resolver is optional. nil means to not consider its options.
//line /snap/go/10455/src/net/conf.go:191
// It also returns dnsConfig when it was used to determine the lookup order.
//line /snap/go/10455/src/net/conf.go:194
func (c *conf) addrLookupOrder(r *Resolver, addr string) (ret hostLookupOrder, dnsConf *dnsConfig) {
//line /snap/go/10455/src/net/conf.go:194
	_go_fuzz_dep_.CoverTab[4781]++
						if c.dnsDebugLevel > 1 {
//line /snap/go/10455/src/net/conf.go:195
		_go_fuzz_dep_.CoverTab[527573]++
//line /snap/go/10455/src/net/conf.go:195
		_go_fuzz_dep_.CoverTab[4783]++
							defer func() {
//line /snap/go/10455/src/net/conf.go:196
			_go_fuzz_dep_.CoverTab[4784]++
								print("go package net: addrLookupOrder(", addr, ") = ", ret.String(), "\n")
//line /snap/go/10455/src/net/conf.go:197
			// _ = "end of CoverTab[4784]"
		}()
//line /snap/go/10455/src/net/conf.go:198
		// _ = "end of CoverTab[4783]"
	} else {
//line /snap/go/10455/src/net/conf.go:199
		_go_fuzz_dep_.CoverTab[527574]++
//line /snap/go/10455/src/net/conf.go:199
		_go_fuzz_dep_.CoverTab[4785]++
//line /snap/go/10455/src/net/conf.go:199
		// _ = "end of CoverTab[4785]"
//line /snap/go/10455/src/net/conf.go:199
	}
//line /snap/go/10455/src/net/conf.go:199
	// _ = "end of CoverTab[4781]"
//line /snap/go/10455/src/net/conf.go:199
	_go_fuzz_dep_.CoverTab[4782]++
						return c.lookupOrder(r, "")
//line /snap/go/10455/src/net/conf.go:200
	// _ = "end of CoverTab[4782]"
}

// hostLookupOrder determines which strategy to use to resolve hostname.
//line /snap/go/10455/src/net/conf.go:203
// The provided Resolver is optional. nil means to not consider its options.
//line /snap/go/10455/src/net/conf.go:203
// It also returns dnsConfig when it was used to determine the lookup order.
//line /snap/go/10455/src/net/conf.go:206
func (c *conf) hostLookupOrder(r *Resolver, hostname string) (ret hostLookupOrder, dnsConf *dnsConfig) {
//line /snap/go/10455/src/net/conf.go:206
	_go_fuzz_dep_.CoverTab[4786]++
						if c.dnsDebugLevel > 1 {
//line /snap/go/10455/src/net/conf.go:207
		_go_fuzz_dep_.CoverTab[527575]++
//line /snap/go/10455/src/net/conf.go:207
		_go_fuzz_dep_.CoverTab[4788]++
							defer func() {
//line /snap/go/10455/src/net/conf.go:208
			_go_fuzz_dep_.CoverTab[4789]++
								print("go package net: hostLookupOrder(", hostname, ") = ", ret.String(), "\n")
//line /snap/go/10455/src/net/conf.go:209
			// _ = "end of CoverTab[4789]"
		}()
//line /snap/go/10455/src/net/conf.go:210
		// _ = "end of CoverTab[4788]"
	} else {
//line /snap/go/10455/src/net/conf.go:211
		_go_fuzz_dep_.CoverTab[527576]++
//line /snap/go/10455/src/net/conf.go:211
		_go_fuzz_dep_.CoverTab[4790]++
//line /snap/go/10455/src/net/conf.go:211
		// _ = "end of CoverTab[4790]"
//line /snap/go/10455/src/net/conf.go:211
	}
//line /snap/go/10455/src/net/conf.go:211
	// _ = "end of CoverTab[4786]"
//line /snap/go/10455/src/net/conf.go:211
	_go_fuzz_dep_.CoverTab[4787]++
						return c.lookupOrder(r, hostname)
//line /snap/go/10455/src/net/conf.go:212
	// _ = "end of CoverTab[4787]"
}

func (c *conf) lookupOrder(r *Resolver, hostname string) (ret hostLookupOrder, dnsConf *dnsConfig) {
//line /snap/go/10455/src/net/conf.go:215
	_go_fuzz_dep_.CoverTab[4791]++
	// fallbackOrder is the order we return if we can't figure it out.
	var fallbackOrder hostLookupOrder

	var canUseCgo bool
	if c.mustUseGoResolver(r) {
//line /snap/go/10455/src/net/conf.go:220
		_go_fuzz_dep_.CoverTab[527577]++
//line /snap/go/10455/src/net/conf.go:220
		_go_fuzz_dep_.CoverTab[4803]++

//line /snap/go/10455/src/net/conf.go:224
		switch c.goos {
		case "windows":
//line /snap/go/10455/src/net/conf.go:225
			_go_fuzz_dep_.CoverTab[527579]++
//line /snap/go/10455/src/net/conf.go:225
			_go_fuzz_dep_.CoverTab[4805]++

//line /snap/go/10455/src/net/conf.go:230
			fallbackOrder = hostLookupDNS
//line /snap/go/10455/src/net/conf.go:230
			// _ = "end of CoverTab[4805]"
		default:
//line /snap/go/10455/src/net/conf.go:231
			_go_fuzz_dep_.CoverTab[527580]++
//line /snap/go/10455/src/net/conf.go:231
			_go_fuzz_dep_.CoverTab[4806]++
								fallbackOrder = hostLookupFilesDNS
//line /snap/go/10455/src/net/conf.go:232
			// _ = "end of CoverTab[4806]"
		}
//line /snap/go/10455/src/net/conf.go:233
		// _ = "end of CoverTab[4803]"
//line /snap/go/10455/src/net/conf.go:233
		_go_fuzz_dep_.CoverTab[4804]++
							canUseCgo = false
//line /snap/go/10455/src/net/conf.go:234
		// _ = "end of CoverTab[4804]"
	} else {
//line /snap/go/10455/src/net/conf.go:235
		_go_fuzz_dep_.CoverTab[527578]++
//line /snap/go/10455/src/net/conf.go:235
		_go_fuzz_dep_.CoverTab[4807]++
//line /snap/go/10455/src/net/conf.go:235
		if c.netCgo {
//line /snap/go/10455/src/net/conf.go:235
			_go_fuzz_dep_.CoverTab[527581]++
//line /snap/go/10455/src/net/conf.go:235
			_go_fuzz_dep_.CoverTab[4808]++

								return hostLookupCgo, nil
//line /snap/go/10455/src/net/conf.go:237
			// _ = "end of CoverTab[4808]"
		} else {
//line /snap/go/10455/src/net/conf.go:238
			_go_fuzz_dep_.CoverTab[527582]++
//line /snap/go/10455/src/net/conf.go:238
			_go_fuzz_dep_.CoverTab[4809]++
//line /snap/go/10455/src/net/conf.go:238
			if c.preferCgo {
//line /snap/go/10455/src/net/conf.go:238
				_go_fuzz_dep_.CoverTab[527583]++
//line /snap/go/10455/src/net/conf.go:238
				_go_fuzz_dep_.CoverTab[4810]++

									return hostLookupCgo, nil
//line /snap/go/10455/src/net/conf.go:240
				// _ = "end of CoverTab[4810]"
			} else {
//line /snap/go/10455/src/net/conf.go:241
				_go_fuzz_dep_.CoverTab[527584]++
//line /snap/go/10455/src/net/conf.go:241
				_go_fuzz_dep_.CoverTab[4811]++

//line /snap/go/10455/src/net/conf.go:245
				if bytealg.IndexByteString(hostname, '\\') != -1 || func() bool {
//line /snap/go/10455/src/net/conf.go:245
					_go_fuzz_dep_.CoverTab[4813]++
//line /snap/go/10455/src/net/conf.go:245
					return bytealg.IndexByteString(hostname, '%') != -1
//line /snap/go/10455/src/net/conf.go:245
					// _ = "end of CoverTab[4813]"
//line /snap/go/10455/src/net/conf.go:245
				}() {
//line /snap/go/10455/src/net/conf.go:245
					_go_fuzz_dep_.CoverTab[527585]++
//line /snap/go/10455/src/net/conf.go:245
					_go_fuzz_dep_.CoverTab[4814]++

//line /snap/go/10455/src/net/conf.go:248
					return hostLookupCgo, nil
//line /snap/go/10455/src/net/conf.go:248
					// _ = "end of CoverTab[4814]"
				} else {
//line /snap/go/10455/src/net/conf.go:249
					_go_fuzz_dep_.CoverTab[527586]++
//line /snap/go/10455/src/net/conf.go:249
					_go_fuzz_dep_.CoverTab[4815]++
//line /snap/go/10455/src/net/conf.go:249
					// _ = "end of CoverTab[4815]"
//line /snap/go/10455/src/net/conf.go:249
				}
//line /snap/go/10455/src/net/conf.go:249
				// _ = "end of CoverTab[4811]"
//line /snap/go/10455/src/net/conf.go:249
				_go_fuzz_dep_.CoverTab[4812]++

//line /snap/go/10455/src/net/conf.go:252
				fallbackOrder = hostLookupCgo
									canUseCgo = true
//line /snap/go/10455/src/net/conf.go:253
				// _ = "end of CoverTab[4812]"
			}
//line /snap/go/10455/src/net/conf.go:254
			// _ = "end of CoverTab[4809]"
//line /snap/go/10455/src/net/conf.go:254
		}
//line /snap/go/10455/src/net/conf.go:254
		// _ = "end of CoverTab[4807]"
//line /snap/go/10455/src/net/conf.go:254
	}
//line /snap/go/10455/src/net/conf.go:254
	// _ = "end of CoverTab[4791]"
//line /snap/go/10455/src/net/conf.go:254
	_go_fuzz_dep_.CoverTab[4792]++

//line /snap/go/10455/src/net/conf.go:257
	switch c.goos {
	case "windows", "plan9", "android", "ios":
//line /snap/go/10455/src/net/conf.go:258
		_go_fuzz_dep_.CoverTab[527587]++
//line /snap/go/10455/src/net/conf.go:258
		_go_fuzz_dep_.CoverTab[4816]++
							return fallbackOrder, nil
//line /snap/go/10455/src/net/conf.go:259
		// _ = "end of CoverTab[4816]"
//line /snap/go/10455/src/net/conf.go:259
	default:
//line /snap/go/10455/src/net/conf.go:259
		_go_fuzz_dep_.CoverTab[527588]++
//line /snap/go/10455/src/net/conf.go:259
		_go_fuzz_dep_.CoverTab[4817]++
//line /snap/go/10455/src/net/conf.go:259
		// _ = "end of CoverTab[4817]"
	}
//line /snap/go/10455/src/net/conf.go:260
	// _ = "end of CoverTab[4792]"
//line /snap/go/10455/src/net/conf.go:260
	_go_fuzz_dep_.CoverTab[4793]++

//line /snap/go/10455/src/net/conf.go:268
	dnsConf = getSystemDNSConfig()

	if canUseCgo && func() bool {
//line /snap/go/10455/src/net/conf.go:270
		_go_fuzz_dep_.CoverTab[4818]++
//line /snap/go/10455/src/net/conf.go:270
		return dnsConf.err != nil
//line /snap/go/10455/src/net/conf.go:270
		// _ = "end of CoverTab[4818]"
//line /snap/go/10455/src/net/conf.go:270
	}() && func() bool {
//line /snap/go/10455/src/net/conf.go:270
		_go_fuzz_dep_.CoverTab[4819]++
//line /snap/go/10455/src/net/conf.go:270
		return !errors.Is(dnsConf.err, fs.ErrNotExist)
//line /snap/go/10455/src/net/conf.go:270
		// _ = "end of CoverTab[4819]"
//line /snap/go/10455/src/net/conf.go:270
	}() && func() bool {
//line /snap/go/10455/src/net/conf.go:270
		_go_fuzz_dep_.CoverTab[4820]++
//line /snap/go/10455/src/net/conf.go:270
		return !errors.Is(dnsConf.err, fs.ErrPermission)
//line /snap/go/10455/src/net/conf.go:270
		// _ = "end of CoverTab[4820]"
//line /snap/go/10455/src/net/conf.go:270
	}() {
//line /snap/go/10455/src/net/conf.go:270
		_go_fuzz_dep_.CoverTab[527589]++
//line /snap/go/10455/src/net/conf.go:270
		_go_fuzz_dep_.CoverTab[4821]++

							return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:272
		// _ = "end of CoverTab[4821]"
	} else {
//line /snap/go/10455/src/net/conf.go:273
		_go_fuzz_dep_.CoverTab[527590]++
//line /snap/go/10455/src/net/conf.go:273
		_go_fuzz_dep_.CoverTab[4822]++
//line /snap/go/10455/src/net/conf.go:273
		// _ = "end of CoverTab[4822]"
//line /snap/go/10455/src/net/conf.go:273
	}
//line /snap/go/10455/src/net/conf.go:273
	// _ = "end of CoverTab[4793]"
//line /snap/go/10455/src/net/conf.go:273
	_go_fuzz_dep_.CoverTab[4794]++

						if canUseCgo && func() bool {
//line /snap/go/10455/src/net/conf.go:275
		_go_fuzz_dep_.CoverTab[4823]++
//line /snap/go/10455/src/net/conf.go:275
		return dnsConf.unknownOpt
//line /snap/go/10455/src/net/conf.go:275
		// _ = "end of CoverTab[4823]"
//line /snap/go/10455/src/net/conf.go:275
	}() {
//line /snap/go/10455/src/net/conf.go:275
		_go_fuzz_dep_.CoverTab[527591]++
//line /snap/go/10455/src/net/conf.go:275
		_go_fuzz_dep_.CoverTab[4824]++

//line /snap/go/10455/src/net/conf.go:278
		return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:278
		// _ = "end of CoverTab[4824]"
	} else {
//line /snap/go/10455/src/net/conf.go:279
		_go_fuzz_dep_.CoverTab[527592]++
//line /snap/go/10455/src/net/conf.go:279
		_go_fuzz_dep_.CoverTab[4825]++
//line /snap/go/10455/src/net/conf.go:279
		// _ = "end of CoverTab[4825]"
//line /snap/go/10455/src/net/conf.go:279
	}
//line /snap/go/10455/src/net/conf.go:279
	// _ = "end of CoverTab[4794]"
//line /snap/go/10455/src/net/conf.go:279
	_go_fuzz_dep_.CoverTab[4795]++

//line /snap/go/10455/src/net/conf.go:283
	if c.goos == "openbsd" {
//line /snap/go/10455/src/net/conf.go:283
		_go_fuzz_dep_.CoverTab[527593]++
//line /snap/go/10455/src/net/conf.go:283
		_go_fuzz_dep_.CoverTab[4826]++

//line /snap/go/10455/src/net/conf.go:287
		if errors.Is(dnsConf.err, fs.ErrNotExist) {
//line /snap/go/10455/src/net/conf.go:287
			_go_fuzz_dep_.CoverTab[527595]++
//line /snap/go/10455/src/net/conf.go:287
			_go_fuzz_dep_.CoverTab[4830]++
								return hostLookupFiles, dnsConf
//line /snap/go/10455/src/net/conf.go:288
			// _ = "end of CoverTab[4830]"
		} else {
//line /snap/go/10455/src/net/conf.go:289
			_go_fuzz_dep_.CoverTab[527596]++
//line /snap/go/10455/src/net/conf.go:289
			_go_fuzz_dep_.CoverTab[4831]++
//line /snap/go/10455/src/net/conf.go:289
			// _ = "end of CoverTab[4831]"
//line /snap/go/10455/src/net/conf.go:289
		}
//line /snap/go/10455/src/net/conf.go:289
		// _ = "end of CoverTab[4826]"
//line /snap/go/10455/src/net/conf.go:289
		_go_fuzz_dep_.CoverTab[4827]++

							lookup := dnsConf.lookup
							if len(lookup) == 0 {
//line /snap/go/10455/src/net/conf.go:292
			_go_fuzz_dep_.CoverTab[527597]++
//line /snap/go/10455/src/net/conf.go:292
			_go_fuzz_dep_.CoverTab[4832]++

//line /snap/go/10455/src/net/conf.go:297
			return hostLookupDNSFiles, dnsConf
//line /snap/go/10455/src/net/conf.go:297
			// _ = "end of CoverTab[4832]"
		} else {
//line /snap/go/10455/src/net/conf.go:298
			_go_fuzz_dep_.CoverTab[527598]++
//line /snap/go/10455/src/net/conf.go:298
			_go_fuzz_dep_.CoverTab[4833]++
//line /snap/go/10455/src/net/conf.go:298
			// _ = "end of CoverTab[4833]"
//line /snap/go/10455/src/net/conf.go:298
		}
//line /snap/go/10455/src/net/conf.go:298
		// _ = "end of CoverTab[4827]"
//line /snap/go/10455/src/net/conf.go:298
		_go_fuzz_dep_.CoverTab[4828]++
							if len(lookup) < 1 || func() bool {
//line /snap/go/10455/src/net/conf.go:299
			_go_fuzz_dep_.CoverTab[4834]++
//line /snap/go/10455/src/net/conf.go:299
			return len(lookup) > 2
//line /snap/go/10455/src/net/conf.go:299
			// _ = "end of CoverTab[4834]"
//line /snap/go/10455/src/net/conf.go:299
		}() {
//line /snap/go/10455/src/net/conf.go:299
			_go_fuzz_dep_.CoverTab[527599]++
//line /snap/go/10455/src/net/conf.go:299
			_go_fuzz_dep_.CoverTab[4835]++

								return fallbackOrder, dnsConf
//line /snap/go/10455/src/net/conf.go:301
			// _ = "end of CoverTab[4835]"
		} else {
//line /snap/go/10455/src/net/conf.go:302
			_go_fuzz_dep_.CoverTab[527600]++
//line /snap/go/10455/src/net/conf.go:302
			_go_fuzz_dep_.CoverTab[4836]++
//line /snap/go/10455/src/net/conf.go:302
			// _ = "end of CoverTab[4836]"
//line /snap/go/10455/src/net/conf.go:302
		}
//line /snap/go/10455/src/net/conf.go:302
		// _ = "end of CoverTab[4828]"
//line /snap/go/10455/src/net/conf.go:302
		_go_fuzz_dep_.CoverTab[4829]++
							switch lookup[0] {
		case "bind":
//line /snap/go/10455/src/net/conf.go:304
			_go_fuzz_dep_.CoverTab[527601]++
//line /snap/go/10455/src/net/conf.go:304
			_go_fuzz_dep_.CoverTab[4837]++
								if len(lookup) == 2 {
//line /snap/go/10455/src/net/conf.go:305
				_go_fuzz_dep_.CoverTab[527604]++
//line /snap/go/10455/src/net/conf.go:305
				_go_fuzz_dep_.CoverTab[4842]++
									if lookup[1] == "file" {
//line /snap/go/10455/src/net/conf.go:306
					_go_fuzz_dep_.CoverTab[527606]++
//line /snap/go/10455/src/net/conf.go:306
					_go_fuzz_dep_.CoverTab[4844]++
										return hostLookupDNSFiles, dnsConf
//line /snap/go/10455/src/net/conf.go:307
					// _ = "end of CoverTab[4844]"
				} else {
//line /snap/go/10455/src/net/conf.go:308
					_go_fuzz_dep_.CoverTab[527607]++
//line /snap/go/10455/src/net/conf.go:308
					_go_fuzz_dep_.CoverTab[4845]++
//line /snap/go/10455/src/net/conf.go:308
					// _ = "end of CoverTab[4845]"
//line /snap/go/10455/src/net/conf.go:308
				}
//line /snap/go/10455/src/net/conf.go:308
				// _ = "end of CoverTab[4842]"
//line /snap/go/10455/src/net/conf.go:308
				_go_fuzz_dep_.CoverTab[4843]++

									return fallbackOrder, dnsConf
//line /snap/go/10455/src/net/conf.go:310
				// _ = "end of CoverTab[4843]"
			} else {
//line /snap/go/10455/src/net/conf.go:311
				_go_fuzz_dep_.CoverTab[527605]++
//line /snap/go/10455/src/net/conf.go:311
				_go_fuzz_dep_.CoverTab[4846]++
//line /snap/go/10455/src/net/conf.go:311
				// _ = "end of CoverTab[4846]"
//line /snap/go/10455/src/net/conf.go:311
			}
//line /snap/go/10455/src/net/conf.go:311
			// _ = "end of CoverTab[4837]"
//line /snap/go/10455/src/net/conf.go:311
			_go_fuzz_dep_.CoverTab[4838]++
								return hostLookupDNS, dnsConf
//line /snap/go/10455/src/net/conf.go:312
			// _ = "end of CoverTab[4838]"
		case "file":
//line /snap/go/10455/src/net/conf.go:313
			_go_fuzz_dep_.CoverTab[527602]++
//line /snap/go/10455/src/net/conf.go:313
			_go_fuzz_dep_.CoverTab[4839]++
								if len(lookup) == 2 {
//line /snap/go/10455/src/net/conf.go:314
				_go_fuzz_dep_.CoverTab[527608]++
//line /snap/go/10455/src/net/conf.go:314
				_go_fuzz_dep_.CoverTab[4847]++
									if lookup[1] == "bind" {
//line /snap/go/10455/src/net/conf.go:315
					_go_fuzz_dep_.CoverTab[527610]++
//line /snap/go/10455/src/net/conf.go:315
					_go_fuzz_dep_.CoverTab[4849]++
										return hostLookupFilesDNS, dnsConf
//line /snap/go/10455/src/net/conf.go:316
					// _ = "end of CoverTab[4849]"
				} else {
//line /snap/go/10455/src/net/conf.go:317
					_go_fuzz_dep_.CoverTab[527611]++
//line /snap/go/10455/src/net/conf.go:317
					_go_fuzz_dep_.CoverTab[4850]++
//line /snap/go/10455/src/net/conf.go:317
					// _ = "end of CoverTab[4850]"
//line /snap/go/10455/src/net/conf.go:317
				}
//line /snap/go/10455/src/net/conf.go:317
				// _ = "end of CoverTab[4847]"
//line /snap/go/10455/src/net/conf.go:317
				_go_fuzz_dep_.CoverTab[4848]++

									return fallbackOrder, dnsConf
//line /snap/go/10455/src/net/conf.go:319
				// _ = "end of CoverTab[4848]"
			} else {
//line /snap/go/10455/src/net/conf.go:320
				_go_fuzz_dep_.CoverTab[527609]++
//line /snap/go/10455/src/net/conf.go:320
				_go_fuzz_dep_.CoverTab[4851]++
//line /snap/go/10455/src/net/conf.go:320
				// _ = "end of CoverTab[4851]"
//line /snap/go/10455/src/net/conf.go:320
			}
//line /snap/go/10455/src/net/conf.go:320
			// _ = "end of CoverTab[4839]"
//line /snap/go/10455/src/net/conf.go:320
			_go_fuzz_dep_.CoverTab[4840]++
								return hostLookupFiles, dnsConf
//line /snap/go/10455/src/net/conf.go:321
			// _ = "end of CoverTab[4840]"
		default:
//line /snap/go/10455/src/net/conf.go:322
			_go_fuzz_dep_.CoverTab[527603]++
//line /snap/go/10455/src/net/conf.go:322
			_go_fuzz_dep_.CoverTab[4841]++

								return fallbackOrder, dnsConf
//line /snap/go/10455/src/net/conf.go:324
			// _ = "end of CoverTab[4841]"
		}
//line /snap/go/10455/src/net/conf.go:325
		// _ = "end of CoverTab[4829]"

//line /snap/go/10455/src/net/conf.go:329
	} else {
//line /snap/go/10455/src/net/conf.go:329
		_go_fuzz_dep_.CoverTab[527594]++
//line /snap/go/10455/src/net/conf.go:329
		_go_fuzz_dep_.CoverTab[4852]++
//line /snap/go/10455/src/net/conf.go:329
		// _ = "end of CoverTab[4852]"
//line /snap/go/10455/src/net/conf.go:329
	}
//line /snap/go/10455/src/net/conf.go:329
	// _ = "end of CoverTab[4795]"
//line /snap/go/10455/src/net/conf.go:329
	_go_fuzz_dep_.CoverTab[4796]++

//line /snap/go/10455/src/net/conf.go:332
	if stringsHasSuffix(hostname, ".") {
//line /snap/go/10455/src/net/conf.go:332
		_go_fuzz_dep_.CoverTab[527612]++
//line /snap/go/10455/src/net/conf.go:332
		_go_fuzz_dep_.CoverTab[4853]++
							hostname = hostname[:len(hostname)-1]
//line /snap/go/10455/src/net/conf.go:333
		// _ = "end of CoverTab[4853]"
	} else {
//line /snap/go/10455/src/net/conf.go:334
		_go_fuzz_dep_.CoverTab[527613]++
//line /snap/go/10455/src/net/conf.go:334
		_go_fuzz_dep_.CoverTab[4854]++
//line /snap/go/10455/src/net/conf.go:334
		// _ = "end of CoverTab[4854]"
//line /snap/go/10455/src/net/conf.go:334
	}
//line /snap/go/10455/src/net/conf.go:334
	// _ = "end of CoverTab[4796]"
//line /snap/go/10455/src/net/conf.go:334
	_go_fuzz_dep_.CoverTab[4797]++
						if canUseCgo && func() bool {
//line /snap/go/10455/src/net/conf.go:335
		_go_fuzz_dep_.CoverTab[4855]++
//line /snap/go/10455/src/net/conf.go:335
		return stringsHasSuffixFold(hostname, ".local")
//line /snap/go/10455/src/net/conf.go:335
		// _ = "end of CoverTab[4855]"
//line /snap/go/10455/src/net/conf.go:335
	}() {
//line /snap/go/10455/src/net/conf.go:335
		_go_fuzz_dep_.CoverTab[527614]++
//line /snap/go/10455/src/net/conf.go:335
		_go_fuzz_dep_.CoverTab[4856]++

//line /snap/go/10455/src/net/conf.go:340
		return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:340
		// _ = "end of CoverTab[4856]"
	} else {
//line /snap/go/10455/src/net/conf.go:341
		_go_fuzz_dep_.CoverTab[527615]++
//line /snap/go/10455/src/net/conf.go:341
		_go_fuzz_dep_.CoverTab[4857]++
//line /snap/go/10455/src/net/conf.go:341
		// _ = "end of CoverTab[4857]"
//line /snap/go/10455/src/net/conf.go:341
	}
//line /snap/go/10455/src/net/conf.go:341
	// _ = "end of CoverTab[4797]"
//line /snap/go/10455/src/net/conf.go:341
	_go_fuzz_dep_.CoverTab[4798]++

						nss := getSystemNSS()
						srcs := nss.sources["hosts"]

//line /snap/go/10455/src/net/conf.go:347
	if errors.Is(nss.err, fs.ErrNotExist) || func() bool {
//line /snap/go/10455/src/net/conf.go:347
		_go_fuzz_dep_.CoverTab[4858]++
//line /snap/go/10455/src/net/conf.go:347
		return (nss.err == nil && func() bool {
//line /snap/go/10455/src/net/conf.go:347
			_go_fuzz_dep_.CoverTab[4859]++
//line /snap/go/10455/src/net/conf.go:347
			return len(srcs) == 0
//line /snap/go/10455/src/net/conf.go:347
			// _ = "end of CoverTab[4859]"
//line /snap/go/10455/src/net/conf.go:347
		}())
//line /snap/go/10455/src/net/conf.go:347
		// _ = "end of CoverTab[4858]"
//line /snap/go/10455/src/net/conf.go:347
	}() {
//line /snap/go/10455/src/net/conf.go:347
		_go_fuzz_dep_.CoverTab[527616]++
//line /snap/go/10455/src/net/conf.go:347
		_go_fuzz_dep_.CoverTab[4860]++
							if canUseCgo && func() bool {
//line /snap/go/10455/src/net/conf.go:348
			_go_fuzz_dep_.CoverTab[4862]++
//line /snap/go/10455/src/net/conf.go:348
			return c.goos == "solaris"
//line /snap/go/10455/src/net/conf.go:348
			// _ = "end of CoverTab[4862]"
//line /snap/go/10455/src/net/conf.go:348
		}() {
//line /snap/go/10455/src/net/conf.go:348
			_go_fuzz_dep_.CoverTab[527618]++
//line /snap/go/10455/src/net/conf.go:348
			_go_fuzz_dep_.CoverTab[4863]++

//line /snap/go/10455/src/net/conf.go:352
			return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:352
			// _ = "end of CoverTab[4863]"
		} else {
//line /snap/go/10455/src/net/conf.go:353
			_go_fuzz_dep_.CoverTab[527619]++
//line /snap/go/10455/src/net/conf.go:353
			_go_fuzz_dep_.CoverTab[4864]++
//line /snap/go/10455/src/net/conf.go:353
			// _ = "end of CoverTab[4864]"
//line /snap/go/10455/src/net/conf.go:353
		}
//line /snap/go/10455/src/net/conf.go:353
		// _ = "end of CoverTab[4860]"
//line /snap/go/10455/src/net/conf.go:353
		_go_fuzz_dep_.CoverTab[4861]++

							return hostLookupFilesDNS, dnsConf
//line /snap/go/10455/src/net/conf.go:355
		// _ = "end of CoverTab[4861]"
	} else {
//line /snap/go/10455/src/net/conf.go:356
		_go_fuzz_dep_.CoverTab[527617]++
//line /snap/go/10455/src/net/conf.go:356
		_go_fuzz_dep_.CoverTab[4865]++
//line /snap/go/10455/src/net/conf.go:356
		// _ = "end of CoverTab[4865]"
//line /snap/go/10455/src/net/conf.go:356
	}
//line /snap/go/10455/src/net/conf.go:356
	// _ = "end of CoverTab[4798]"
//line /snap/go/10455/src/net/conf.go:356
	_go_fuzz_dep_.CoverTab[4799]++
						if nss.err != nil {
//line /snap/go/10455/src/net/conf.go:357
		_go_fuzz_dep_.CoverTab[527620]++
//line /snap/go/10455/src/net/conf.go:357
		_go_fuzz_dep_.CoverTab[4866]++

//line /snap/go/10455/src/net/conf.go:360
		return fallbackOrder, dnsConf
//line /snap/go/10455/src/net/conf.go:360
		// _ = "end of CoverTab[4866]"
	} else {
//line /snap/go/10455/src/net/conf.go:361
		_go_fuzz_dep_.CoverTab[527621]++
//line /snap/go/10455/src/net/conf.go:361
		_go_fuzz_dep_.CoverTab[4867]++
//line /snap/go/10455/src/net/conf.go:361
		// _ = "end of CoverTab[4867]"
//line /snap/go/10455/src/net/conf.go:361
	}
//line /snap/go/10455/src/net/conf.go:361
	// _ = "end of CoverTab[4799]"
//line /snap/go/10455/src/net/conf.go:361
	_go_fuzz_dep_.CoverTab[4800]++

						var hasDNSSource bool
						var hasDNSSourceChecked bool

						var filesSource, dnsSource bool
						var first string
//line /snap/go/10455/src/net/conf.go:367
	_go_fuzz_dep_.CoverTab[786647] = 0
						for i, src := range srcs {
//line /snap/go/10455/src/net/conf.go:368
		if _go_fuzz_dep_.CoverTab[786647] == 0 {
//line /snap/go/10455/src/net/conf.go:368
			_go_fuzz_dep_.CoverTab[527667]++
//line /snap/go/10455/src/net/conf.go:368
		} else {
//line /snap/go/10455/src/net/conf.go:368
			_go_fuzz_dep_.CoverTab[527668]++
//line /snap/go/10455/src/net/conf.go:368
		}
//line /snap/go/10455/src/net/conf.go:368
		_go_fuzz_dep_.CoverTab[786647] = 1
//line /snap/go/10455/src/net/conf.go:368
		_go_fuzz_dep_.CoverTab[4868]++
							if src.source == "files" || func() bool {
//line /snap/go/10455/src/net/conf.go:369
			_go_fuzz_dep_.CoverTab[4872]++
//line /snap/go/10455/src/net/conf.go:369
			return src.source == "dns"
//line /snap/go/10455/src/net/conf.go:369
			// _ = "end of CoverTab[4872]"
//line /snap/go/10455/src/net/conf.go:369
		}() {
//line /snap/go/10455/src/net/conf.go:369
			_go_fuzz_dep_.CoverTab[527622]++
//line /snap/go/10455/src/net/conf.go:369
			_go_fuzz_dep_.CoverTab[4873]++
								if canUseCgo && func() bool {
//line /snap/go/10455/src/net/conf.go:370
				_go_fuzz_dep_.CoverTab[4877]++
//line /snap/go/10455/src/net/conf.go:370
				return !src.standardCriteria()
//line /snap/go/10455/src/net/conf.go:370
				// _ = "end of CoverTab[4877]"
//line /snap/go/10455/src/net/conf.go:370
			}() {
//line /snap/go/10455/src/net/conf.go:370
				_go_fuzz_dep_.CoverTab[527624]++
//line /snap/go/10455/src/net/conf.go:370
				_go_fuzz_dep_.CoverTab[4878]++

									return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:372
				// _ = "end of CoverTab[4878]"
			} else {
//line /snap/go/10455/src/net/conf.go:373
				_go_fuzz_dep_.CoverTab[527625]++
//line /snap/go/10455/src/net/conf.go:373
				_go_fuzz_dep_.CoverTab[4879]++
//line /snap/go/10455/src/net/conf.go:373
				// _ = "end of CoverTab[4879]"
//line /snap/go/10455/src/net/conf.go:373
			}
//line /snap/go/10455/src/net/conf.go:373
			// _ = "end of CoverTab[4873]"
//line /snap/go/10455/src/net/conf.go:373
			_go_fuzz_dep_.CoverTab[4874]++
								if src.source == "files" {
//line /snap/go/10455/src/net/conf.go:374
				_go_fuzz_dep_.CoverTab[527626]++
//line /snap/go/10455/src/net/conf.go:374
				_go_fuzz_dep_.CoverTab[4880]++
									filesSource = true
//line /snap/go/10455/src/net/conf.go:375
				// _ = "end of CoverTab[4880]"
			} else {
//line /snap/go/10455/src/net/conf.go:376
				_go_fuzz_dep_.CoverTab[527627]++
//line /snap/go/10455/src/net/conf.go:376
				_go_fuzz_dep_.CoverTab[4881]++
									hasDNSSource = true
									hasDNSSourceChecked = true
									dnsSource = true
//line /snap/go/10455/src/net/conf.go:379
				// _ = "end of CoverTab[4881]"
			}
//line /snap/go/10455/src/net/conf.go:380
			// _ = "end of CoverTab[4874]"
//line /snap/go/10455/src/net/conf.go:380
			_go_fuzz_dep_.CoverTab[4875]++
								if first == "" {
//line /snap/go/10455/src/net/conf.go:381
				_go_fuzz_dep_.CoverTab[527628]++
//line /snap/go/10455/src/net/conf.go:381
				_go_fuzz_dep_.CoverTab[4882]++
									first = src.source
//line /snap/go/10455/src/net/conf.go:382
				// _ = "end of CoverTab[4882]"
			} else {
//line /snap/go/10455/src/net/conf.go:383
				_go_fuzz_dep_.CoverTab[527629]++
//line /snap/go/10455/src/net/conf.go:383
				_go_fuzz_dep_.CoverTab[4883]++
//line /snap/go/10455/src/net/conf.go:383
				// _ = "end of CoverTab[4883]"
//line /snap/go/10455/src/net/conf.go:383
			}
//line /snap/go/10455/src/net/conf.go:383
			// _ = "end of CoverTab[4875]"
//line /snap/go/10455/src/net/conf.go:383
			_go_fuzz_dep_.CoverTab[4876]++
								continue
//line /snap/go/10455/src/net/conf.go:384
			// _ = "end of CoverTab[4876]"
		} else {
//line /snap/go/10455/src/net/conf.go:385
			_go_fuzz_dep_.CoverTab[527623]++
//line /snap/go/10455/src/net/conf.go:385
			_go_fuzz_dep_.CoverTab[4884]++
//line /snap/go/10455/src/net/conf.go:385
			// _ = "end of CoverTab[4884]"
//line /snap/go/10455/src/net/conf.go:385
		}
//line /snap/go/10455/src/net/conf.go:385
		// _ = "end of CoverTab[4868]"
//line /snap/go/10455/src/net/conf.go:385
		_go_fuzz_dep_.CoverTab[4869]++

							if canUseCgo {
//line /snap/go/10455/src/net/conf.go:387
			_go_fuzz_dep_.CoverTab[527630]++
//line /snap/go/10455/src/net/conf.go:387
			_go_fuzz_dep_.CoverTab[4885]++
								switch {
			case hostname != "" && func() bool {
//line /snap/go/10455/src/net/conf.go:389
				_go_fuzz_dep_.CoverTab[4893]++
//line /snap/go/10455/src/net/conf.go:389
				return src.source == "myhostname"
//line /snap/go/10455/src/net/conf.go:389
				// _ = "end of CoverTab[4893]"
//line /snap/go/10455/src/net/conf.go:389
			}():
//line /snap/go/10455/src/net/conf.go:389
				_go_fuzz_dep_.CoverTab[527632]++
//line /snap/go/10455/src/net/conf.go:389
				_go_fuzz_dep_.CoverTab[4886]++

//line /snap/go/10455/src/net/conf.go:392
				if isLocalhost(hostname) || func() bool {
//line /snap/go/10455/src/net/conf.go:392
					_go_fuzz_dep_.CoverTab[4894]++
//line /snap/go/10455/src/net/conf.go:392
					return isGateway(hostname)
//line /snap/go/10455/src/net/conf.go:392
					// _ = "end of CoverTab[4894]"
//line /snap/go/10455/src/net/conf.go:392
				}() || func() bool {
//line /snap/go/10455/src/net/conf.go:392
					_go_fuzz_dep_.CoverTab[4895]++
//line /snap/go/10455/src/net/conf.go:392
					return isOutbound(hostname)
//line /snap/go/10455/src/net/conf.go:392
					// _ = "end of CoverTab[4895]"
//line /snap/go/10455/src/net/conf.go:392
				}() {
//line /snap/go/10455/src/net/conf.go:392
					_go_fuzz_dep_.CoverTab[527635]++
//line /snap/go/10455/src/net/conf.go:392
					_go_fuzz_dep_.CoverTab[4896]++
										return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:393
					// _ = "end of CoverTab[4896]"
				} else {
//line /snap/go/10455/src/net/conf.go:394
					_go_fuzz_dep_.CoverTab[527636]++
//line /snap/go/10455/src/net/conf.go:394
					_go_fuzz_dep_.CoverTab[4897]++
//line /snap/go/10455/src/net/conf.go:394
					// _ = "end of CoverTab[4897]"
//line /snap/go/10455/src/net/conf.go:394
				}
//line /snap/go/10455/src/net/conf.go:394
				// _ = "end of CoverTab[4886]"
//line /snap/go/10455/src/net/conf.go:394
				_go_fuzz_dep_.CoverTab[4887]++
									hn, err := getHostname()
									if err != nil || func() bool {
//line /snap/go/10455/src/net/conf.go:396
					_go_fuzz_dep_.CoverTab[4898]++
//line /snap/go/10455/src/net/conf.go:396
					return stringsEqualFold(hostname, hn)
//line /snap/go/10455/src/net/conf.go:396
					// _ = "end of CoverTab[4898]"
//line /snap/go/10455/src/net/conf.go:396
				}() {
//line /snap/go/10455/src/net/conf.go:396
					_go_fuzz_dep_.CoverTab[527637]++
//line /snap/go/10455/src/net/conf.go:396
					_go_fuzz_dep_.CoverTab[4899]++
										return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:397
					// _ = "end of CoverTab[4899]"
				} else {
//line /snap/go/10455/src/net/conf.go:398
					_go_fuzz_dep_.CoverTab[527638]++
//line /snap/go/10455/src/net/conf.go:398
					_go_fuzz_dep_.CoverTab[4900]++
//line /snap/go/10455/src/net/conf.go:398
					// _ = "end of CoverTab[4900]"
//line /snap/go/10455/src/net/conf.go:398
				}
//line /snap/go/10455/src/net/conf.go:398
				// _ = "end of CoverTab[4887]"
//line /snap/go/10455/src/net/conf.go:398
				_go_fuzz_dep_.CoverTab[4888]++
									continue
//line /snap/go/10455/src/net/conf.go:399
				// _ = "end of CoverTab[4888]"
			case hostname != "" && func() bool {
//line /snap/go/10455/src/net/conf.go:400
				_go_fuzz_dep_.CoverTab[4901]++
//line /snap/go/10455/src/net/conf.go:400
				return stringsHasPrefix(src.source, "mdns")
//line /snap/go/10455/src/net/conf.go:400
				// _ = "end of CoverTab[4901]"
//line /snap/go/10455/src/net/conf.go:400
			}():
//line /snap/go/10455/src/net/conf.go:400
				_go_fuzz_dep_.CoverTab[527633]++
//line /snap/go/10455/src/net/conf.go:400
				_go_fuzz_dep_.CoverTab[4889]++

//line /snap/go/10455/src/net/conf.go:405
				// We don't parse mdns.allow files. They're rare. If one
				// exists, it might list other TLDs (besides .local) or even
				// '*', so just let libc deal with it.
				var haveMDNSAllow bool
				switch c.mdnsTest {
				case mdnsFromSystem:
//line /snap/go/10455/src/net/conf.go:410
					_go_fuzz_dep_.CoverTab[527639]++
//line /snap/go/10455/src/net/conf.go:410
					_go_fuzz_dep_.CoverTab[4902]++
										_, err := os.Stat("/etc/mdns.allow")
										if err != nil && func() bool {
//line /snap/go/10455/src/net/conf.go:412
						_go_fuzz_dep_.CoverTab[4907]++
//line /snap/go/10455/src/net/conf.go:412
						return !errors.Is(err, fs.ErrNotExist)
//line /snap/go/10455/src/net/conf.go:412
						// _ = "end of CoverTab[4907]"
//line /snap/go/10455/src/net/conf.go:412
					}() {
//line /snap/go/10455/src/net/conf.go:412
						_go_fuzz_dep_.CoverTab[527643]++
//line /snap/go/10455/src/net/conf.go:412
						_go_fuzz_dep_.CoverTab[4908]++

											return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:414
						// _ = "end of CoverTab[4908]"
					} else {
//line /snap/go/10455/src/net/conf.go:415
						_go_fuzz_dep_.CoverTab[527644]++
//line /snap/go/10455/src/net/conf.go:415
						_go_fuzz_dep_.CoverTab[4909]++
//line /snap/go/10455/src/net/conf.go:415
						// _ = "end of CoverTab[4909]"
//line /snap/go/10455/src/net/conf.go:415
					}
//line /snap/go/10455/src/net/conf.go:415
					// _ = "end of CoverTab[4902]"
//line /snap/go/10455/src/net/conf.go:415
					_go_fuzz_dep_.CoverTab[4903]++
										haveMDNSAllow = err == nil
//line /snap/go/10455/src/net/conf.go:416
					// _ = "end of CoverTab[4903]"
				case mdnsAssumeExists:
//line /snap/go/10455/src/net/conf.go:417
					_go_fuzz_dep_.CoverTab[527640]++
//line /snap/go/10455/src/net/conf.go:417
					_go_fuzz_dep_.CoverTab[4904]++
										haveMDNSAllow = true
//line /snap/go/10455/src/net/conf.go:418
					// _ = "end of CoverTab[4904]"
				case mdnsAssumeDoesNotExist:
//line /snap/go/10455/src/net/conf.go:419
					_go_fuzz_dep_.CoverTab[527641]++
//line /snap/go/10455/src/net/conf.go:419
					_go_fuzz_dep_.CoverTab[4905]++
										haveMDNSAllow = false
//line /snap/go/10455/src/net/conf.go:420
					// _ = "end of CoverTab[4905]"
//line /snap/go/10455/src/net/conf.go:420
				default:
//line /snap/go/10455/src/net/conf.go:420
					_go_fuzz_dep_.CoverTab[527642]++
//line /snap/go/10455/src/net/conf.go:420
					_go_fuzz_dep_.CoverTab[4906]++
//line /snap/go/10455/src/net/conf.go:420
					// _ = "end of CoverTab[4906]"
				}
//line /snap/go/10455/src/net/conf.go:421
				// _ = "end of CoverTab[4889]"
//line /snap/go/10455/src/net/conf.go:421
				_go_fuzz_dep_.CoverTab[4890]++
									if haveMDNSAllow {
//line /snap/go/10455/src/net/conf.go:422
					_go_fuzz_dep_.CoverTab[527645]++
//line /snap/go/10455/src/net/conf.go:422
					_go_fuzz_dep_.CoverTab[4910]++
										return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:423
					// _ = "end of CoverTab[4910]"
				} else {
//line /snap/go/10455/src/net/conf.go:424
					_go_fuzz_dep_.CoverTab[527646]++
//line /snap/go/10455/src/net/conf.go:424
					_go_fuzz_dep_.CoverTab[4911]++
//line /snap/go/10455/src/net/conf.go:424
					// _ = "end of CoverTab[4911]"
//line /snap/go/10455/src/net/conf.go:424
				}
//line /snap/go/10455/src/net/conf.go:424
				// _ = "end of CoverTab[4890]"
//line /snap/go/10455/src/net/conf.go:424
				_go_fuzz_dep_.CoverTab[4891]++
									continue
//line /snap/go/10455/src/net/conf.go:425
				// _ = "end of CoverTab[4891]"
			default:
//line /snap/go/10455/src/net/conf.go:426
				_go_fuzz_dep_.CoverTab[527634]++
//line /snap/go/10455/src/net/conf.go:426
				_go_fuzz_dep_.CoverTab[4892]++

									return hostLookupCgo, dnsConf
//line /snap/go/10455/src/net/conf.go:428
				// _ = "end of CoverTab[4892]"
			}
//line /snap/go/10455/src/net/conf.go:429
			// _ = "end of CoverTab[4885]"
		} else {
//line /snap/go/10455/src/net/conf.go:430
			_go_fuzz_dep_.CoverTab[527631]++
//line /snap/go/10455/src/net/conf.go:430
			_go_fuzz_dep_.CoverTab[4912]++
//line /snap/go/10455/src/net/conf.go:430
			// _ = "end of CoverTab[4912]"
//line /snap/go/10455/src/net/conf.go:430
		}
//line /snap/go/10455/src/net/conf.go:430
		// _ = "end of CoverTab[4869]"
//line /snap/go/10455/src/net/conf.go:430
		_go_fuzz_dep_.CoverTab[4870]++

							if !hasDNSSourceChecked {
//line /snap/go/10455/src/net/conf.go:432
			_go_fuzz_dep_.CoverTab[527647]++
//line /snap/go/10455/src/net/conf.go:432
			_go_fuzz_dep_.CoverTab[4913]++
								hasDNSSourceChecked = true
//line /snap/go/10455/src/net/conf.go:433
			_go_fuzz_dep_.CoverTab[786648] = 0
								for _, v := range srcs[i+1:] {
//line /snap/go/10455/src/net/conf.go:434
				if _go_fuzz_dep_.CoverTab[786648] == 0 {
//line /snap/go/10455/src/net/conf.go:434
					_go_fuzz_dep_.CoverTab[527671]++
//line /snap/go/10455/src/net/conf.go:434
				} else {
//line /snap/go/10455/src/net/conf.go:434
					_go_fuzz_dep_.CoverTab[527672]++
//line /snap/go/10455/src/net/conf.go:434
				}
//line /snap/go/10455/src/net/conf.go:434
				_go_fuzz_dep_.CoverTab[786648] = 1
//line /snap/go/10455/src/net/conf.go:434
				_go_fuzz_dep_.CoverTab[4914]++
									if v.source == "dns" {
//line /snap/go/10455/src/net/conf.go:435
					_go_fuzz_dep_.CoverTab[527649]++
//line /snap/go/10455/src/net/conf.go:435
					_go_fuzz_dep_.CoverTab[4915]++
										hasDNSSource = true
										break
//line /snap/go/10455/src/net/conf.go:437
					// _ = "end of CoverTab[4915]"
				} else {
//line /snap/go/10455/src/net/conf.go:438
					_go_fuzz_dep_.CoverTab[527650]++
//line /snap/go/10455/src/net/conf.go:438
					_go_fuzz_dep_.CoverTab[4916]++
//line /snap/go/10455/src/net/conf.go:438
					// _ = "end of CoverTab[4916]"
//line /snap/go/10455/src/net/conf.go:438
				}
//line /snap/go/10455/src/net/conf.go:438
				// _ = "end of CoverTab[4914]"
			}
//line /snap/go/10455/src/net/conf.go:439
			if _go_fuzz_dep_.CoverTab[786648] == 0 {
//line /snap/go/10455/src/net/conf.go:439
				_go_fuzz_dep_.CoverTab[527673]++
//line /snap/go/10455/src/net/conf.go:439
			} else {
//line /snap/go/10455/src/net/conf.go:439
				_go_fuzz_dep_.CoverTab[527674]++
//line /snap/go/10455/src/net/conf.go:439
			}
//line /snap/go/10455/src/net/conf.go:439
			// _ = "end of CoverTab[4913]"
		} else {
//line /snap/go/10455/src/net/conf.go:440
			_go_fuzz_dep_.CoverTab[527648]++
//line /snap/go/10455/src/net/conf.go:440
			_go_fuzz_dep_.CoverTab[4917]++
//line /snap/go/10455/src/net/conf.go:440
			// _ = "end of CoverTab[4917]"
//line /snap/go/10455/src/net/conf.go:440
		}
//line /snap/go/10455/src/net/conf.go:440
		// _ = "end of CoverTab[4870]"
//line /snap/go/10455/src/net/conf.go:440
		_go_fuzz_dep_.CoverTab[4871]++

//line /snap/go/10455/src/net/conf.go:445
		if !hasDNSSource {
//line /snap/go/10455/src/net/conf.go:445
			_go_fuzz_dep_.CoverTab[527651]++
//line /snap/go/10455/src/net/conf.go:445
			_go_fuzz_dep_.CoverTab[4918]++
								dnsSource = true
								if first == "" {
//line /snap/go/10455/src/net/conf.go:447
				_go_fuzz_dep_.CoverTab[527653]++
//line /snap/go/10455/src/net/conf.go:447
				_go_fuzz_dep_.CoverTab[4919]++
									first = "dns"
//line /snap/go/10455/src/net/conf.go:448
				// _ = "end of CoverTab[4919]"
			} else {
//line /snap/go/10455/src/net/conf.go:449
				_go_fuzz_dep_.CoverTab[527654]++
//line /snap/go/10455/src/net/conf.go:449
				_go_fuzz_dep_.CoverTab[4920]++
//line /snap/go/10455/src/net/conf.go:449
				// _ = "end of CoverTab[4920]"
//line /snap/go/10455/src/net/conf.go:449
			}
//line /snap/go/10455/src/net/conf.go:449
			// _ = "end of CoverTab[4918]"
		} else {
//line /snap/go/10455/src/net/conf.go:450
			_go_fuzz_dep_.CoverTab[527652]++
//line /snap/go/10455/src/net/conf.go:450
			_go_fuzz_dep_.CoverTab[4921]++
//line /snap/go/10455/src/net/conf.go:450
			// _ = "end of CoverTab[4921]"
//line /snap/go/10455/src/net/conf.go:450
		}
//line /snap/go/10455/src/net/conf.go:450
		// _ = "end of CoverTab[4871]"
	}
//line /snap/go/10455/src/net/conf.go:451
	if _go_fuzz_dep_.CoverTab[786647] == 0 {
//line /snap/go/10455/src/net/conf.go:451
		_go_fuzz_dep_.CoverTab[527669]++
//line /snap/go/10455/src/net/conf.go:451
	} else {
//line /snap/go/10455/src/net/conf.go:451
		_go_fuzz_dep_.CoverTab[527670]++
//line /snap/go/10455/src/net/conf.go:451
	}
//line /snap/go/10455/src/net/conf.go:451
	// _ = "end of CoverTab[4800]"
//line /snap/go/10455/src/net/conf.go:451
	_go_fuzz_dep_.CoverTab[4801]++

//line /snap/go/10455/src/net/conf.go:455
	switch {
	case filesSource && func() bool {
//line /snap/go/10455/src/net/conf.go:456
		_go_fuzz_dep_.CoverTab[4926]++
//line /snap/go/10455/src/net/conf.go:456
		return dnsSource
//line /snap/go/10455/src/net/conf.go:456
		// _ = "end of CoverTab[4926]"
//line /snap/go/10455/src/net/conf.go:456
	}():
//line /snap/go/10455/src/net/conf.go:456
		_go_fuzz_dep_.CoverTab[527655]++
//line /snap/go/10455/src/net/conf.go:456
		_go_fuzz_dep_.CoverTab[4922]++
							if first == "files" {
//line /snap/go/10455/src/net/conf.go:457
			_go_fuzz_dep_.CoverTab[527659]++
//line /snap/go/10455/src/net/conf.go:457
			_go_fuzz_dep_.CoverTab[4927]++
								return hostLookupFilesDNS, dnsConf
//line /snap/go/10455/src/net/conf.go:458
			// _ = "end of CoverTab[4927]"
		} else {
//line /snap/go/10455/src/net/conf.go:459
			_go_fuzz_dep_.CoverTab[527660]++
//line /snap/go/10455/src/net/conf.go:459
			_go_fuzz_dep_.CoverTab[4928]++
								return hostLookupDNSFiles, dnsConf
//line /snap/go/10455/src/net/conf.go:460
			// _ = "end of CoverTab[4928]"
		}
//line /snap/go/10455/src/net/conf.go:461
		// _ = "end of CoverTab[4922]"
	case filesSource:
//line /snap/go/10455/src/net/conf.go:462
		_go_fuzz_dep_.CoverTab[527656]++
//line /snap/go/10455/src/net/conf.go:462
		_go_fuzz_dep_.CoverTab[4923]++
							return hostLookupFiles, dnsConf
//line /snap/go/10455/src/net/conf.go:463
		// _ = "end of CoverTab[4923]"
	case dnsSource:
//line /snap/go/10455/src/net/conf.go:464
		_go_fuzz_dep_.CoverTab[527657]++
//line /snap/go/10455/src/net/conf.go:464
		_go_fuzz_dep_.CoverTab[4924]++
							return hostLookupDNS, dnsConf
//line /snap/go/10455/src/net/conf.go:465
		// _ = "end of CoverTab[4924]"
//line /snap/go/10455/src/net/conf.go:465
	default:
//line /snap/go/10455/src/net/conf.go:465
		_go_fuzz_dep_.CoverTab[527658]++
//line /snap/go/10455/src/net/conf.go:465
		_go_fuzz_dep_.CoverTab[4925]++
//line /snap/go/10455/src/net/conf.go:465
		// _ = "end of CoverTab[4925]"
	}
//line /snap/go/10455/src/net/conf.go:466
	// _ = "end of CoverTab[4801]"
//line /snap/go/10455/src/net/conf.go:466
	_go_fuzz_dep_.CoverTab[4802]++

//line /snap/go/10455/src/net/conf.go:469
	return fallbackOrder, dnsConf
//line /snap/go/10455/src/net/conf.go:469
	// _ = "end of CoverTab[4802]"
}

var netdns = godebug.New("netdns")

// goDebugNetDNS parses the value of the GODEBUG "netdns" value.
//line /snap/go/10455/src/net/conf.go:474
// The netdns value can be of the form:
//line /snap/go/10455/src/net/conf.go:474
//
//line /snap/go/10455/src/net/conf.go:474
//	1       // debug level 1
//line /snap/go/10455/src/net/conf.go:474
//	2       // debug level 2
//line /snap/go/10455/src/net/conf.go:474
//	cgo     // use cgo for DNS lookups
//line /snap/go/10455/src/net/conf.go:474
//	go      // use go for DNS lookups
//line /snap/go/10455/src/net/conf.go:474
//	cgo+1   // use cgo for DNS lookups + debug level 1
//line /snap/go/10455/src/net/conf.go:474
//	1+cgo   // same
//line /snap/go/10455/src/net/conf.go:474
//	cgo+2   // same, but debug level 2
//line /snap/go/10455/src/net/conf.go:474
//
//line /snap/go/10455/src/net/conf.go:474
// etc.
//line /snap/go/10455/src/net/conf.go:486
func goDebugNetDNS() (dnsMode string, debugLevel int) {
//line /snap/go/10455/src/net/conf.go:486
	_go_fuzz_dep_.CoverTab[4929]++
						goDebug := netdns.Value()
						parsePart := func(s string) {
//line /snap/go/10455/src/net/conf.go:488
		_go_fuzz_dep_.CoverTab[4932]++
							if s == "" {
//line /snap/go/10455/src/net/conf.go:489
			_go_fuzz_dep_.CoverTab[527661]++
//line /snap/go/10455/src/net/conf.go:489
			_go_fuzz_dep_.CoverTab[4934]++
								return
//line /snap/go/10455/src/net/conf.go:490
			// _ = "end of CoverTab[4934]"
		} else {
//line /snap/go/10455/src/net/conf.go:491
			_go_fuzz_dep_.CoverTab[527662]++
//line /snap/go/10455/src/net/conf.go:491
			_go_fuzz_dep_.CoverTab[4935]++
//line /snap/go/10455/src/net/conf.go:491
			// _ = "end of CoverTab[4935]"
//line /snap/go/10455/src/net/conf.go:491
		}
//line /snap/go/10455/src/net/conf.go:491
		// _ = "end of CoverTab[4932]"
//line /snap/go/10455/src/net/conf.go:491
		_go_fuzz_dep_.CoverTab[4933]++
							if '0' <= s[0] && func() bool {
//line /snap/go/10455/src/net/conf.go:492
			_go_fuzz_dep_.CoverTab[4936]++
//line /snap/go/10455/src/net/conf.go:492
			return s[0] <= '9'
//line /snap/go/10455/src/net/conf.go:492
			// _ = "end of CoverTab[4936]"
//line /snap/go/10455/src/net/conf.go:492
		}() {
//line /snap/go/10455/src/net/conf.go:492
			_go_fuzz_dep_.CoverTab[527663]++
//line /snap/go/10455/src/net/conf.go:492
			_go_fuzz_dep_.CoverTab[4937]++
								debugLevel, _, _ = dtoi(s)
//line /snap/go/10455/src/net/conf.go:493
			// _ = "end of CoverTab[4937]"
		} else {
//line /snap/go/10455/src/net/conf.go:494
			_go_fuzz_dep_.CoverTab[527664]++
//line /snap/go/10455/src/net/conf.go:494
			_go_fuzz_dep_.CoverTab[4938]++
								dnsMode = s
//line /snap/go/10455/src/net/conf.go:495
			// _ = "end of CoverTab[4938]"
		}
//line /snap/go/10455/src/net/conf.go:496
		// _ = "end of CoverTab[4933]"
	}
//line /snap/go/10455/src/net/conf.go:497
	// _ = "end of CoverTab[4929]"
//line /snap/go/10455/src/net/conf.go:497
	_go_fuzz_dep_.CoverTab[4930]++
						if i := bytealg.IndexByteString(goDebug, '+'); i != -1 {
//line /snap/go/10455/src/net/conf.go:498
		_go_fuzz_dep_.CoverTab[527665]++
//line /snap/go/10455/src/net/conf.go:498
		_go_fuzz_dep_.CoverTab[4939]++
							parsePart(goDebug[:i])
							parsePart(goDebug[i+1:])
							return
//line /snap/go/10455/src/net/conf.go:501
		// _ = "end of CoverTab[4939]"
	} else {
//line /snap/go/10455/src/net/conf.go:502
		_go_fuzz_dep_.CoverTab[527666]++
//line /snap/go/10455/src/net/conf.go:502
		_go_fuzz_dep_.CoverTab[4940]++
//line /snap/go/10455/src/net/conf.go:502
		// _ = "end of CoverTab[4940]"
//line /snap/go/10455/src/net/conf.go:502
	}
//line /snap/go/10455/src/net/conf.go:502
	// _ = "end of CoverTab[4930]"
//line /snap/go/10455/src/net/conf.go:502
	_go_fuzz_dep_.CoverTab[4931]++
						parsePart(goDebug)
						return
//line /snap/go/10455/src/net/conf.go:504
	// _ = "end of CoverTab[4931]"
}

// isLocalhost reports whether h should be considered a "localhost"
//line /snap/go/10455/src/net/conf.go:507
// name for the myhostname NSS module.
//line /snap/go/10455/src/net/conf.go:509
func isLocalhost(h string) bool {
//line /snap/go/10455/src/net/conf.go:509
	_go_fuzz_dep_.CoverTab[4941]++
						return stringsEqualFold(h, "localhost") || func() bool {
//line /snap/go/10455/src/net/conf.go:510
		_go_fuzz_dep_.CoverTab[4942]++
//line /snap/go/10455/src/net/conf.go:510
		return stringsEqualFold(h, "localhost.localdomain")
//line /snap/go/10455/src/net/conf.go:510
		// _ = "end of CoverTab[4942]"
//line /snap/go/10455/src/net/conf.go:510
	}() || func() bool {
//line /snap/go/10455/src/net/conf.go:510
		_go_fuzz_dep_.CoverTab[4943]++
//line /snap/go/10455/src/net/conf.go:510
		return stringsHasSuffixFold(h, ".localhost")
//line /snap/go/10455/src/net/conf.go:510
		// _ = "end of CoverTab[4943]"
//line /snap/go/10455/src/net/conf.go:510
	}() || func() bool {
//line /snap/go/10455/src/net/conf.go:510
		_go_fuzz_dep_.CoverTab[4944]++
//line /snap/go/10455/src/net/conf.go:510
		return stringsHasSuffixFold(h, ".localhost.localdomain")
//line /snap/go/10455/src/net/conf.go:510
		// _ = "end of CoverTab[4944]"
//line /snap/go/10455/src/net/conf.go:510
	}()
//line /snap/go/10455/src/net/conf.go:510
	// _ = "end of CoverTab[4941]"
}

// isGateway reports whether h should be considered a "gateway"
//line /snap/go/10455/src/net/conf.go:513
// name for the myhostname NSS module.
//line /snap/go/10455/src/net/conf.go:515
func isGateway(h string) bool {
//line /snap/go/10455/src/net/conf.go:515
	_go_fuzz_dep_.CoverTab[4945]++
						return stringsEqualFold(h, "_gateway")
//line /snap/go/10455/src/net/conf.go:516
	// _ = "end of CoverTab[4945]"
}

// isOutbound reports whether h should be considered a "outbound"
//line /snap/go/10455/src/net/conf.go:519
// name for the myhostname NSS module.
//line /snap/go/10455/src/net/conf.go:521
func isOutbound(h string) bool {
//line /snap/go/10455/src/net/conf.go:521
	_go_fuzz_dep_.CoverTab[4946]++
						return stringsEqualFold(h, "_outbound")
//line /snap/go/10455/src/net/conf.go:522
	// _ = "end of CoverTab[4946]"
}

//line /snap/go/10455/src/net/conf.go:523
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/conf.go:523
var _ = _go_fuzz_dep_.CoverTab
