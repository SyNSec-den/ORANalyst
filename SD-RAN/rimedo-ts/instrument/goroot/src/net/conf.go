// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

//line /usr/local/go/src/net/conf.go:7
package net

//line /usr/local/go/src/net/conf.go:7
import (
//line /usr/local/go/src/net/conf.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/conf.go:7
)
//line /usr/local/go/src/net/conf.go:7
import (
//line /usr/local/go/src/net/conf.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/conf.go:7
)

import (
	"internal/bytealg"
	"internal/godebug"
	"os"
	"runtime"
	"sync"
	"syscall"
)

// conf represents a system's network configuration.
type conf struct {
	// forceCgoLookupHost forces CGO to always be used, if available.
	forceCgoLookupHost	bool

	netGo	bool	// go DNS resolution forced
	netCgo	bool	// non-go DNS resolution forced (cgo, or win32)

	// machine has an /etc/mdns.allow file
	hasMDNSAllow	bool

	goos		string	// the runtime.GOOS, to ease testing
	dnsDebugLevel	int
}

var (
	confOnce	sync.Once	// guards init of confVal via initConfVal
	confVal		= &conf{goos: runtime.GOOS}
)

// systemConf returns the machine's network configuration.
func systemConf() *conf {
//line /usr/local/go/src/net/conf.go:39
	_go_fuzz_dep_.CoverTab[12813]++
						confOnce.Do(initConfVal)
						return confVal
//line /usr/local/go/src/net/conf.go:41
	// _ = "end of CoverTab[12813]"
}

func initConfVal() {
//line /usr/local/go/src/net/conf.go:44
	_go_fuzz_dep_.CoverTab[12814]++
						dnsMode, debugLevel := goDebugNetDNS()
						confVal.dnsDebugLevel = debugLevel
						confVal.netGo = netGo || func() bool {
//line /usr/local/go/src/net/conf.go:47
		_go_fuzz_dep_.CoverTab[12821]++
//line /usr/local/go/src/net/conf.go:47
		return dnsMode == "go"
//line /usr/local/go/src/net/conf.go:47
		// _ = "end of CoverTab[12821]"
//line /usr/local/go/src/net/conf.go:47
	}()
						confVal.netCgo = netCgo || func() bool {
//line /usr/local/go/src/net/conf.go:48
		_go_fuzz_dep_.CoverTab[12822]++
//line /usr/local/go/src/net/conf.go:48
		return dnsMode == "cgo"
//line /usr/local/go/src/net/conf.go:48
		// _ = "end of CoverTab[12822]"
//line /usr/local/go/src/net/conf.go:48
	}()
						if !confVal.netGo && func() bool {
//line /usr/local/go/src/net/conf.go:49
		_go_fuzz_dep_.CoverTab[12823]++
//line /usr/local/go/src/net/conf.go:49
		return !confVal.netCgo
//line /usr/local/go/src/net/conf.go:49
		// _ = "end of CoverTab[12823]"
//line /usr/local/go/src/net/conf.go:49
	}() && func() bool {
//line /usr/local/go/src/net/conf.go:49
		_go_fuzz_dep_.CoverTab[12824]++
//line /usr/local/go/src/net/conf.go:49
		return (runtime.GOOS == "windows" || func() bool {
//line /usr/local/go/src/net/conf.go:49
			_go_fuzz_dep_.CoverTab[12825]++
//line /usr/local/go/src/net/conf.go:49
			return runtime.GOOS == "plan9"
//line /usr/local/go/src/net/conf.go:49
			// _ = "end of CoverTab[12825]"
//line /usr/local/go/src/net/conf.go:49
		}())
//line /usr/local/go/src/net/conf.go:49
		// _ = "end of CoverTab[12824]"
//line /usr/local/go/src/net/conf.go:49
	}() {
//line /usr/local/go/src/net/conf.go:49
		_go_fuzz_dep_.CoverTab[12826]++

//line /usr/local/go/src/net/conf.go:58
		confVal.netCgo = true
//line /usr/local/go/src/net/conf.go:58
		// _ = "end of CoverTab[12826]"
	} else {
//line /usr/local/go/src/net/conf.go:59
		_go_fuzz_dep_.CoverTab[12827]++
//line /usr/local/go/src/net/conf.go:59
		// _ = "end of CoverTab[12827]"
//line /usr/local/go/src/net/conf.go:59
	}
//line /usr/local/go/src/net/conf.go:59
	// _ = "end of CoverTab[12814]"
//line /usr/local/go/src/net/conf.go:59
	_go_fuzz_dep_.CoverTab[12815]++

						if confVal.dnsDebugLevel > 0 {
//line /usr/local/go/src/net/conf.go:61
		_go_fuzz_dep_.CoverTab[12828]++
							defer func() {
//line /usr/local/go/src/net/conf.go:62
			_go_fuzz_dep_.CoverTab[12829]++
								if confVal.dnsDebugLevel > 1 {
//line /usr/local/go/src/net/conf.go:63
				_go_fuzz_dep_.CoverTab[12831]++
									println("go package net: confVal.netCgo =", confVal.netCgo, " netGo =", confVal.netGo)
//line /usr/local/go/src/net/conf.go:64
				// _ = "end of CoverTab[12831]"
			} else {
//line /usr/local/go/src/net/conf.go:65
				_go_fuzz_dep_.CoverTab[12832]++
//line /usr/local/go/src/net/conf.go:65
				// _ = "end of CoverTab[12832]"
//line /usr/local/go/src/net/conf.go:65
			}
//line /usr/local/go/src/net/conf.go:65
			// _ = "end of CoverTab[12829]"
//line /usr/local/go/src/net/conf.go:65
			_go_fuzz_dep_.CoverTab[12830]++
								switch {
			case confVal.netGo:
//line /usr/local/go/src/net/conf.go:67
				_go_fuzz_dep_.CoverTab[12833]++
									if netGo {
//line /usr/local/go/src/net/conf.go:68
					_go_fuzz_dep_.CoverTab[12836]++
										println("go package net: built with netgo build tag; using Go's DNS resolver")
//line /usr/local/go/src/net/conf.go:69
					// _ = "end of CoverTab[12836]"
				} else {
//line /usr/local/go/src/net/conf.go:70
					_go_fuzz_dep_.CoverTab[12837]++
										println("go package net: GODEBUG setting forcing use of Go's resolver")
//line /usr/local/go/src/net/conf.go:71
					// _ = "end of CoverTab[12837]"
				}
//line /usr/local/go/src/net/conf.go:72
				// _ = "end of CoverTab[12833]"
			case confVal.forceCgoLookupHost:
//line /usr/local/go/src/net/conf.go:73
				_go_fuzz_dep_.CoverTab[12834]++
									println("go package net: using cgo DNS resolver")
//line /usr/local/go/src/net/conf.go:74
				// _ = "end of CoverTab[12834]"
			default:
//line /usr/local/go/src/net/conf.go:75
				_go_fuzz_dep_.CoverTab[12835]++
									println("go package net: dynamic selection of DNS resolver")
//line /usr/local/go/src/net/conf.go:76
				// _ = "end of CoverTab[12835]"
			}
//line /usr/local/go/src/net/conf.go:77
			// _ = "end of CoverTab[12830]"
		}()
//line /usr/local/go/src/net/conf.go:78
		// _ = "end of CoverTab[12828]"
	} else {
//line /usr/local/go/src/net/conf.go:79
		_go_fuzz_dep_.CoverTab[12838]++
//line /usr/local/go/src/net/conf.go:79
		// _ = "end of CoverTab[12838]"
//line /usr/local/go/src/net/conf.go:79
	}
//line /usr/local/go/src/net/conf.go:79
	// _ = "end of CoverTab[12815]"
//line /usr/local/go/src/net/conf.go:79
	_go_fuzz_dep_.CoverTab[12816]++

//line /usr/local/go/src/net/conf.go:84
	if runtime.GOOS == "darwin" || func() bool {
//line /usr/local/go/src/net/conf.go:84
		_go_fuzz_dep_.CoverTab[12839]++
//line /usr/local/go/src/net/conf.go:84
		return runtime.GOOS == "ios"
//line /usr/local/go/src/net/conf.go:84
		// _ = "end of CoverTab[12839]"
//line /usr/local/go/src/net/conf.go:84
	}() {
//line /usr/local/go/src/net/conf.go:84
		_go_fuzz_dep_.CoverTab[12840]++
							confVal.forceCgoLookupHost = true
							return
//line /usr/local/go/src/net/conf.go:86
		// _ = "end of CoverTab[12840]"
	} else {
//line /usr/local/go/src/net/conf.go:87
		_go_fuzz_dep_.CoverTab[12841]++
//line /usr/local/go/src/net/conf.go:87
		// _ = "end of CoverTab[12841]"
//line /usr/local/go/src/net/conf.go:87
	}
//line /usr/local/go/src/net/conf.go:87
	// _ = "end of CoverTab[12816]"
//line /usr/local/go/src/net/conf.go:87
	_go_fuzz_dep_.CoverTab[12817]++

						if runtime.GOOS == "windows" || func() bool {
//line /usr/local/go/src/net/conf.go:89
		_go_fuzz_dep_.CoverTab[12842]++
//line /usr/local/go/src/net/conf.go:89
		return runtime.GOOS == "plan9"
//line /usr/local/go/src/net/conf.go:89
		// _ = "end of CoverTab[12842]"
//line /usr/local/go/src/net/conf.go:89
	}() {
//line /usr/local/go/src/net/conf.go:89
		_go_fuzz_dep_.CoverTab[12843]++
							return
//line /usr/local/go/src/net/conf.go:90
		// _ = "end of CoverTab[12843]"
	} else {
//line /usr/local/go/src/net/conf.go:91
		_go_fuzz_dep_.CoverTab[12844]++
//line /usr/local/go/src/net/conf.go:91
		// _ = "end of CoverTab[12844]"
//line /usr/local/go/src/net/conf.go:91
	}
//line /usr/local/go/src/net/conf.go:91
	// _ = "end of CoverTab[12817]"
//line /usr/local/go/src/net/conf.go:91
	_go_fuzz_dep_.CoverTab[12818]++

//line /usr/local/go/src/net/conf.go:96
	_, localDomainDefined := syscall.Getenv("LOCALDOMAIN")
	if os.Getenv("RES_OPTIONS") != "" || func() bool {
//line /usr/local/go/src/net/conf.go:97
		_go_fuzz_dep_.CoverTab[12845]++
//line /usr/local/go/src/net/conf.go:97
		return os.Getenv("HOSTALIASES") != ""
							// _ = "end of CoverTab[12845]"
//line /usr/local/go/src/net/conf.go:98
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:98
		_go_fuzz_dep_.CoverTab[12846]++
//line /usr/local/go/src/net/conf.go:98
		return confVal.netCgo
							// _ = "end of CoverTab[12846]"
//line /usr/local/go/src/net/conf.go:99
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:99
		_go_fuzz_dep_.CoverTab[12847]++
//line /usr/local/go/src/net/conf.go:99
		return localDomainDefined
							// _ = "end of CoverTab[12847]"
//line /usr/local/go/src/net/conf.go:100
	}() {
//line /usr/local/go/src/net/conf.go:100
		_go_fuzz_dep_.CoverTab[12848]++
							confVal.forceCgoLookupHost = true
							return
//line /usr/local/go/src/net/conf.go:102
		// _ = "end of CoverTab[12848]"
	} else {
//line /usr/local/go/src/net/conf.go:103
		_go_fuzz_dep_.CoverTab[12849]++
//line /usr/local/go/src/net/conf.go:103
		// _ = "end of CoverTab[12849]"
//line /usr/local/go/src/net/conf.go:103
	}
//line /usr/local/go/src/net/conf.go:103
	// _ = "end of CoverTab[12818]"
//line /usr/local/go/src/net/conf.go:103
	_go_fuzz_dep_.CoverTab[12819]++

//line /usr/local/go/src/net/conf.go:107
	if runtime.GOOS == "openbsd" && func() bool {
//line /usr/local/go/src/net/conf.go:107
		_go_fuzz_dep_.CoverTab[12850]++
//line /usr/local/go/src/net/conf.go:107
		return os.Getenv("ASR_CONFIG") != ""
//line /usr/local/go/src/net/conf.go:107
		// _ = "end of CoverTab[12850]"
//line /usr/local/go/src/net/conf.go:107
	}() {
//line /usr/local/go/src/net/conf.go:107
		_go_fuzz_dep_.CoverTab[12851]++
							confVal.forceCgoLookupHost = true
							return
//line /usr/local/go/src/net/conf.go:109
		// _ = "end of CoverTab[12851]"
	} else {
//line /usr/local/go/src/net/conf.go:110
		_go_fuzz_dep_.CoverTab[12852]++
//line /usr/local/go/src/net/conf.go:110
		// _ = "end of CoverTab[12852]"
//line /usr/local/go/src/net/conf.go:110
	}
//line /usr/local/go/src/net/conf.go:110
	// _ = "end of CoverTab[12819]"
//line /usr/local/go/src/net/conf.go:110
	_go_fuzz_dep_.CoverTab[12820]++

						if _, err := os.Stat("/etc/mdns.allow"); err == nil {
//line /usr/local/go/src/net/conf.go:112
		_go_fuzz_dep_.CoverTab[12853]++
							confVal.hasMDNSAllow = true
//line /usr/local/go/src/net/conf.go:113
		// _ = "end of CoverTab[12853]"
	} else {
//line /usr/local/go/src/net/conf.go:114
		_go_fuzz_dep_.CoverTab[12854]++
//line /usr/local/go/src/net/conf.go:114
		// _ = "end of CoverTab[12854]"
//line /usr/local/go/src/net/conf.go:114
	}
//line /usr/local/go/src/net/conf.go:114
	// _ = "end of CoverTab[12820]"
}

// canUseCgo reports whether calling cgo functions is allowed
//line /usr/local/go/src/net/conf.go:117
// for non-hostname lookups.
//line /usr/local/go/src/net/conf.go:119
func (c *conf) canUseCgo() bool {
//line /usr/local/go/src/net/conf.go:119
	_go_fuzz_dep_.CoverTab[12855]++
						ret, _ := c.hostLookupOrder(nil, "")
						return ret == hostLookupCgo
//line /usr/local/go/src/net/conf.go:121
	// _ = "end of CoverTab[12855]"
}

// hostLookupOrder determines which strategy to use to resolve hostname.
//line /usr/local/go/src/net/conf.go:124
// The provided Resolver is optional. nil means to not consider its options.
//line /usr/local/go/src/net/conf.go:124
// It also returns dnsConfig when it was used to determine the lookup order.
//line /usr/local/go/src/net/conf.go:127
func (c *conf) hostLookupOrder(r *Resolver, hostname string) (ret hostLookupOrder, dnsConfig *dnsConfig) {
//line /usr/local/go/src/net/conf.go:127
	_go_fuzz_dep_.CoverTab[12856]++
						if c.dnsDebugLevel > 1 {
//line /usr/local/go/src/net/conf.go:128
		_go_fuzz_dep_.CoverTab[12871]++
							defer func() {
//line /usr/local/go/src/net/conf.go:129
			_go_fuzz_dep_.CoverTab[12872]++
								print("go package net: hostLookupOrder(", hostname, ") = ", ret.String(), "\n")
//line /usr/local/go/src/net/conf.go:130
			// _ = "end of CoverTab[12872]"
		}()
//line /usr/local/go/src/net/conf.go:131
		// _ = "end of CoverTab[12871]"
	} else {
//line /usr/local/go/src/net/conf.go:132
		_go_fuzz_dep_.CoverTab[12873]++
//line /usr/local/go/src/net/conf.go:132
		// _ = "end of CoverTab[12873]"
//line /usr/local/go/src/net/conf.go:132
	}
//line /usr/local/go/src/net/conf.go:132
	// _ = "end of CoverTab[12856]"
//line /usr/local/go/src/net/conf.go:132
	_go_fuzz_dep_.CoverTab[12857]++
						fallbackOrder := hostLookupCgo
						if c.netGo || func() bool {
//line /usr/local/go/src/net/conf.go:134
		_go_fuzz_dep_.CoverTab[12874]++
//line /usr/local/go/src/net/conf.go:134
		return r.preferGo()
//line /usr/local/go/src/net/conf.go:134
		// _ = "end of CoverTab[12874]"
//line /usr/local/go/src/net/conf.go:134
	}() {
//line /usr/local/go/src/net/conf.go:134
		_go_fuzz_dep_.CoverTab[12875]++
							switch c.goos {
		case "windows":
//line /usr/local/go/src/net/conf.go:136
			_go_fuzz_dep_.CoverTab[12876]++

//line /usr/local/go/src/net/conf.go:141
			fallbackOrder = hostLookupDNS
//line /usr/local/go/src/net/conf.go:141
			// _ = "end of CoverTab[12876]"
		default:
//line /usr/local/go/src/net/conf.go:142
			_go_fuzz_dep_.CoverTab[12877]++
								fallbackOrder = hostLookupFilesDNS
//line /usr/local/go/src/net/conf.go:143
			// _ = "end of CoverTab[12877]"
		}
//line /usr/local/go/src/net/conf.go:144
		// _ = "end of CoverTab[12875]"
	} else {
//line /usr/local/go/src/net/conf.go:145
		_go_fuzz_dep_.CoverTab[12878]++
//line /usr/local/go/src/net/conf.go:145
		// _ = "end of CoverTab[12878]"
//line /usr/local/go/src/net/conf.go:145
	}
//line /usr/local/go/src/net/conf.go:145
	// _ = "end of CoverTab[12857]"
//line /usr/local/go/src/net/conf.go:145
	_go_fuzz_dep_.CoverTab[12858]++
						if c.forceCgoLookupHost || func() bool {
//line /usr/local/go/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[12879]++
//line /usr/local/go/src/net/conf.go:146
		return c.goos == "android"
//line /usr/local/go/src/net/conf.go:146
		// _ = "end of CoverTab[12879]"
//line /usr/local/go/src/net/conf.go:146
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[12880]++
//line /usr/local/go/src/net/conf.go:146
		return c.goos == "windows"
//line /usr/local/go/src/net/conf.go:146
		// _ = "end of CoverTab[12880]"
//line /usr/local/go/src/net/conf.go:146
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[12881]++
//line /usr/local/go/src/net/conf.go:146
		return c.goos == "plan9"
//line /usr/local/go/src/net/conf.go:146
		// _ = "end of CoverTab[12881]"
//line /usr/local/go/src/net/conf.go:146
	}() {
//line /usr/local/go/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[12882]++
							return fallbackOrder, nil
//line /usr/local/go/src/net/conf.go:147
		// _ = "end of CoverTab[12882]"
	} else {
//line /usr/local/go/src/net/conf.go:148
		_go_fuzz_dep_.CoverTab[12883]++
//line /usr/local/go/src/net/conf.go:148
		// _ = "end of CoverTab[12883]"
//line /usr/local/go/src/net/conf.go:148
	}
//line /usr/local/go/src/net/conf.go:148
	// _ = "end of CoverTab[12858]"
//line /usr/local/go/src/net/conf.go:148
	_go_fuzz_dep_.CoverTab[12859]++
						if bytealg.IndexByteString(hostname, '\\') != -1 || func() bool {
//line /usr/local/go/src/net/conf.go:149
		_go_fuzz_dep_.CoverTab[12884]++
//line /usr/local/go/src/net/conf.go:149
		return bytealg.IndexByteString(hostname, '%') != -1
//line /usr/local/go/src/net/conf.go:149
		// _ = "end of CoverTab[12884]"
//line /usr/local/go/src/net/conf.go:149
	}() {
//line /usr/local/go/src/net/conf.go:149
		_go_fuzz_dep_.CoverTab[12885]++

//line /usr/local/go/src/net/conf.go:152
		return fallbackOrder, nil
//line /usr/local/go/src/net/conf.go:152
		// _ = "end of CoverTab[12885]"
	} else {
//line /usr/local/go/src/net/conf.go:153
		_go_fuzz_dep_.CoverTab[12886]++
//line /usr/local/go/src/net/conf.go:153
		// _ = "end of CoverTab[12886]"
//line /usr/local/go/src/net/conf.go:153
	}
//line /usr/local/go/src/net/conf.go:153
	// _ = "end of CoverTab[12859]"
//line /usr/local/go/src/net/conf.go:153
	_go_fuzz_dep_.CoverTab[12860]++

						conf := getSystemDNSConfig()
						if conf.err != nil && func() bool {
//line /usr/local/go/src/net/conf.go:156
		_go_fuzz_dep_.CoverTab[12887]++
//line /usr/local/go/src/net/conf.go:156
		return !os.IsNotExist(conf.err)
//line /usr/local/go/src/net/conf.go:156
		// _ = "end of CoverTab[12887]"
//line /usr/local/go/src/net/conf.go:156
	}() && func() bool {
//line /usr/local/go/src/net/conf.go:156
		_go_fuzz_dep_.CoverTab[12888]++
//line /usr/local/go/src/net/conf.go:156
		return !os.IsPermission(conf.err)
//line /usr/local/go/src/net/conf.go:156
		// _ = "end of CoverTab[12888]"
//line /usr/local/go/src/net/conf.go:156
	}() {
//line /usr/local/go/src/net/conf.go:156
		_go_fuzz_dep_.CoverTab[12889]++

//line /usr/local/go/src/net/conf.go:161
		return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:161
		// _ = "end of CoverTab[12889]"
	} else {
//line /usr/local/go/src/net/conf.go:162
		_go_fuzz_dep_.CoverTab[12890]++
//line /usr/local/go/src/net/conf.go:162
		// _ = "end of CoverTab[12890]"
//line /usr/local/go/src/net/conf.go:162
	}
//line /usr/local/go/src/net/conf.go:162
	// _ = "end of CoverTab[12860]"
//line /usr/local/go/src/net/conf.go:162
	_go_fuzz_dep_.CoverTab[12861]++

						if conf.unknownOpt {
//line /usr/local/go/src/net/conf.go:164
		_go_fuzz_dep_.CoverTab[12891]++
							return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:165
		// _ = "end of CoverTab[12891]"
	} else {
//line /usr/local/go/src/net/conf.go:166
		_go_fuzz_dep_.CoverTab[12892]++
//line /usr/local/go/src/net/conf.go:166
		// _ = "end of CoverTab[12892]"
//line /usr/local/go/src/net/conf.go:166
	}
//line /usr/local/go/src/net/conf.go:166
	// _ = "end of CoverTab[12861]"
//line /usr/local/go/src/net/conf.go:166
	_go_fuzz_dep_.CoverTab[12862]++

//line /usr/local/go/src/net/conf.go:170
	if c.goos == "openbsd" {
//line /usr/local/go/src/net/conf.go:170
		_go_fuzz_dep_.CoverTab[12893]++

//line /usr/local/go/src/net/conf.go:174
		if os.IsNotExist(conf.err) {
//line /usr/local/go/src/net/conf.go:174
			_go_fuzz_dep_.CoverTab[12897]++
								return hostLookupFiles, conf
//line /usr/local/go/src/net/conf.go:175
			// _ = "end of CoverTab[12897]"
		} else {
//line /usr/local/go/src/net/conf.go:176
			_go_fuzz_dep_.CoverTab[12898]++
//line /usr/local/go/src/net/conf.go:176
			// _ = "end of CoverTab[12898]"
//line /usr/local/go/src/net/conf.go:176
		}
//line /usr/local/go/src/net/conf.go:176
		// _ = "end of CoverTab[12893]"
//line /usr/local/go/src/net/conf.go:176
		_go_fuzz_dep_.CoverTab[12894]++

							lookup := conf.lookup
							if len(lookup) == 0 {
//line /usr/local/go/src/net/conf.go:179
			_go_fuzz_dep_.CoverTab[12899]++

//line /usr/local/go/src/net/conf.go:184
			return hostLookupDNSFiles, conf
//line /usr/local/go/src/net/conf.go:184
			// _ = "end of CoverTab[12899]"
		} else {
//line /usr/local/go/src/net/conf.go:185
			_go_fuzz_dep_.CoverTab[12900]++
//line /usr/local/go/src/net/conf.go:185
			// _ = "end of CoverTab[12900]"
//line /usr/local/go/src/net/conf.go:185
		}
//line /usr/local/go/src/net/conf.go:185
		// _ = "end of CoverTab[12894]"
//line /usr/local/go/src/net/conf.go:185
		_go_fuzz_dep_.CoverTab[12895]++
							if len(lookup) < 1 || func() bool {
//line /usr/local/go/src/net/conf.go:186
			_go_fuzz_dep_.CoverTab[12901]++
//line /usr/local/go/src/net/conf.go:186
			return len(lookup) > 2
//line /usr/local/go/src/net/conf.go:186
			// _ = "end of CoverTab[12901]"
//line /usr/local/go/src/net/conf.go:186
		}() {
//line /usr/local/go/src/net/conf.go:186
			_go_fuzz_dep_.CoverTab[12902]++
								return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:187
			// _ = "end of CoverTab[12902]"
		} else {
//line /usr/local/go/src/net/conf.go:188
			_go_fuzz_dep_.CoverTab[12903]++
//line /usr/local/go/src/net/conf.go:188
			// _ = "end of CoverTab[12903]"
//line /usr/local/go/src/net/conf.go:188
		}
//line /usr/local/go/src/net/conf.go:188
		// _ = "end of CoverTab[12895]"
//line /usr/local/go/src/net/conf.go:188
		_go_fuzz_dep_.CoverTab[12896]++
							switch lookup[0] {
		case "bind":
//line /usr/local/go/src/net/conf.go:190
			_go_fuzz_dep_.CoverTab[12904]++
								if len(lookup) == 2 {
//line /usr/local/go/src/net/conf.go:191
				_go_fuzz_dep_.CoverTab[12909]++
									if lookup[1] == "file" {
//line /usr/local/go/src/net/conf.go:192
					_go_fuzz_dep_.CoverTab[12911]++
										return hostLookupDNSFiles, conf
//line /usr/local/go/src/net/conf.go:193
					// _ = "end of CoverTab[12911]"
				} else {
//line /usr/local/go/src/net/conf.go:194
					_go_fuzz_dep_.CoverTab[12912]++
//line /usr/local/go/src/net/conf.go:194
					// _ = "end of CoverTab[12912]"
//line /usr/local/go/src/net/conf.go:194
				}
//line /usr/local/go/src/net/conf.go:194
				// _ = "end of CoverTab[12909]"
//line /usr/local/go/src/net/conf.go:194
				_go_fuzz_dep_.CoverTab[12910]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:195
				// _ = "end of CoverTab[12910]"
			} else {
//line /usr/local/go/src/net/conf.go:196
				_go_fuzz_dep_.CoverTab[12913]++
//line /usr/local/go/src/net/conf.go:196
				// _ = "end of CoverTab[12913]"
//line /usr/local/go/src/net/conf.go:196
			}
//line /usr/local/go/src/net/conf.go:196
			// _ = "end of CoverTab[12904]"
//line /usr/local/go/src/net/conf.go:196
			_go_fuzz_dep_.CoverTab[12905]++
								return hostLookupDNS, conf
//line /usr/local/go/src/net/conf.go:197
			// _ = "end of CoverTab[12905]"
		case "file":
//line /usr/local/go/src/net/conf.go:198
			_go_fuzz_dep_.CoverTab[12906]++
								if len(lookup) == 2 {
//line /usr/local/go/src/net/conf.go:199
				_go_fuzz_dep_.CoverTab[12914]++
									if lookup[1] == "bind" {
//line /usr/local/go/src/net/conf.go:200
					_go_fuzz_dep_.CoverTab[12916]++
										return hostLookupFilesDNS, conf
//line /usr/local/go/src/net/conf.go:201
					// _ = "end of CoverTab[12916]"
				} else {
//line /usr/local/go/src/net/conf.go:202
					_go_fuzz_dep_.CoverTab[12917]++
//line /usr/local/go/src/net/conf.go:202
					// _ = "end of CoverTab[12917]"
//line /usr/local/go/src/net/conf.go:202
				}
//line /usr/local/go/src/net/conf.go:202
				// _ = "end of CoverTab[12914]"
//line /usr/local/go/src/net/conf.go:202
				_go_fuzz_dep_.CoverTab[12915]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:203
				// _ = "end of CoverTab[12915]"
			} else {
//line /usr/local/go/src/net/conf.go:204
				_go_fuzz_dep_.CoverTab[12918]++
//line /usr/local/go/src/net/conf.go:204
				// _ = "end of CoverTab[12918]"
//line /usr/local/go/src/net/conf.go:204
			}
//line /usr/local/go/src/net/conf.go:204
			// _ = "end of CoverTab[12906]"
//line /usr/local/go/src/net/conf.go:204
			_go_fuzz_dep_.CoverTab[12907]++
								return hostLookupFiles, conf
//line /usr/local/go/src/net/conf.go:205
			// _ = "end of CoverTab[12907]"
		default:
//line /usr/local/go/src/net/conf.go:206
			_go_fuzz_dep_.CoverTab[12908]++
								return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:207
			// _ = "end of CoverTab[12908]"
		}
//line /usr/local/go/src/net/conf.go:208
		// _ = "end of CoverTab[12896]"
	} else {
//line /usr/local/go/src/net/conf.go:209
		_go_fuzz_dep_.CoverTab[12919]++
//line /usr/local/go/src/net/conf.go:209
		// _ = "end of CoverTab[12919]"
//line /usr/local/go/src/net/conf.go:209
	}
//line /usr/local/go/src/net/conf.go:209
	// _ = "end of CoverTab[12862]"
//line /usr/local/go/src/net/conf.go:209
	_go_fuzz_dep_.CoverTab[12863]++

//line /usr/local/go/src/net/conf.go:212
	if stringsHasSuffix(hostname, ".") {
//line /usr/local/go/src/net/conf.go:212
		_go_fuzz_dep_.CoverTab[12920]++
							hostname = hostname[:len(hostname)-1]
//line /usr/local/go/src/net/conf.go:213
		// _ = "end of CoverTab[12920]"
	} else {
//line /usr/local/go/src/net/conf.go:214
		_go_fuzz_dep_.CoverTab[12921]++
//line /usr/local/go/src/net/conf.go:214
		// _ = "end of CoverTab[12921]"
//line /usr/local/go/src/net/conf.go:214
	}
//line /usr/local/go/src/net/conf.go:214
	// _ = "end of CoverTab[12863]"
//line /usr/local/go/src/net/conf.go:214
	_go_fuzz_dep_.CoverTab[12864]++
						if stringsHasSuffixFold(hostname, ".local") {
//line /usr/local/go/src/net/conf.go:215
		_go_fuzz_dep_.CoverTab[12922]++

//line /usr/local/go/src/net/conf.go:220
		return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:220
		// _ = "end of CoverTab[12922]"
	} else {
//line /usr/local/go/src/net/conf.go:221
		_go_fuzz_dep_.CoverTab[12923]++
//line /usr/local/go/src/net/conf.go:221
		// _ = "end of CoverTab[12923]"
//line /usr/local/go/src/net/conf.go:221
	}
//line /usr/local/go/src/net/conf.go:221
	// _ = "end of CoverTab[12864]"
//line /usr/local/go/src/net/conf.go:221
	_go_fuzz_dep_.CoverTab[12865]++

						nss := getSystemNSS()
						srcs := nss.sources["hosts"]

//line /usr/local/go/src/net/conf.go:227
	if os.IsNotExist(nss.err) || func() bool {
//line /usr/local/go/src/net/conf.go:227
		_go_fuzz_dep_.CoverTab[12924]++
//line /usr/local/go/src/net/conf.go:227
		return (nss.err == nil && func() bool {
//line /usr/local/go/src/net/conf.go:227
			_go_fuzz_dep_.CoverTab[12925]++
//line /usr/local/go/src/net/conf.go:227
			return len(srcs) == 0
//line /usr/local/go/src/net/conf.go:227
			// _ = "end of CoverTab[12925]"
//line /usr/local/go/src/net/conf.go:227
		}())
//line /usr/local/go/src/net/conf.go:227
		// _ = "end of CoverTab[12924]"
//line /usr/local/go/src/net/conf.go:227
	}() {
//line /usr/local/go/src/net/conf.go:227
		_go_fuzz_dep_.CoverTab[12926]++
							if c.goos == "solaris" {
//line /usr/local/go/src/net/conf.go:228
			_go_fuzz_dep_.CoverTab[12928]++

								return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:230
			// _ = "end of CoverTab[12928]"
		} else {
//line /usr/local/go/src/net/conf.go:231
			_go_fuzz_dep_.CoverTab[12929]++
//line /usr/local/go/src/net/conf.go:231
			// _ = "end of CoverTab[12929]"
//line /usr/local/go/src/net/conf.go:231
		}
//line /usr/local/go/src/net/conf.go:231
		// _ = "end of CoverTab[12926]"
//line /usr/local/go/src/net/conf.go:231
		_go_fuzz_dep_.CoverTab[12927]++

							return hostLookupFilesDNS, conf
//line /usr/local/go/src/net/conf.go:233
		// _ = "end of CoverTab[12927]"
	} else {
//line /usr/local/go/src/net/conf.go:234
		_go_fuzz_dep_.CoverTab[12930]++
//line /usr/local/go/src/net/conf.go:234
		// _ = "end of CoverTab[12930]"
//line /usr/local/go/src/net/conf.go:234
	}
//line /usr/local/go/src/net/conf.go:234
	// _ = "end of CoverTab[12865]"
//line /usr/local/go/src/net/conf.go:234
	_go_fuzz_dep_.CoverTab[12866]++
						if nss.err != nil {
//line /usr/local/go/src/net/conf.go:235
		_go_fuzz_dep_.CoverTab[12931]++

//line /usr/local/go/src/net/conf.go:239
		return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:239
		// _ = "end of CoverTab[12931]"
	} else {
//line /usr/local/go/src/net/conf.go:240
		_go_fuzz_dep_.CoverTab[12932]++
//line /usr/local/go/src/net/conf.go:240
		// _ = "end of CoverTab[12932]"
//line /usr/local/go/src/net/conf.go:240
	}
//line /usr/local/go/src/net/conf.go:240
	// _ = "end of CoverTab[12866]"
//line /usr/local/go/src/net/conf.go:240
	_go_fuzz_dep_.CoverTab[12867]++

						var mdnsSource, filesSource, dnsSource bool
						var first string
						for _, src := range srcs {
//line /usr/local/go/src/net/conf.go:244
		_go_fuzz_dep_.CoverTab[12933]++
							if src.source == "myhostname" {
//line /usr/local/go/src/net/conf.go:245
			_go_fuzz_dep_.CoverTab[12937]++
								if isLocalhost(hostname) || func() bool {
//line /usr/local/go/src/net/conf.go:246
				_go_fuzz_dep_.CoverTab[12940]++
//line /usr/local/go/src/net/conf.go:246
				return isGateway(hostname)
//line /usr/local/go/src/net/conf.go:246
				// _ = "end of CoverTab[12940]"
//line /usr/local/go/src/net/conf.go:246
			}() || func() bool {
//line /usr/local/go/src/net/conf.go:246
				_go_fuzz_dep_.CoverTab[12941]++
//line /usr/local/go/src/net/conf.go:246
				return isOutbound(hostname)
//line /usr/local/go/src/net/conf.go:246
				// _ = "end of CoverTab[12941]"
//line /usr/local/go/src/net/conf.go:246
			}() {
//line /usr/local/go/src/net/conf.go:246
				_go_fuzz_dep_.CoverTab[12942]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:247
				// _ = "end of CoverTab[12942]"
			} else {
//line /usr/local/go/src/net/conf.go:248
				_go_fuzz_dep_.CoverTab[12943]++
//line /usr/local/go/src/net/conf.go:248
				// _ = "end of CoverTab[12943]"
//line /usr/local/go/src/net/conf.go:248
			}
//line /usr/local/go/src/net/conf.go:248
			// _ = "end of CoverTab[12937]"
//line /usr/local/go/src/net/conf.go:248
			_go_fuzz_dep_.CoverTab[12938]++
								hn, err := getHostname()
								if err != nil || func() bool {
//line /usr/local/go/src/net/conf.go:250
				_go_fuzz_dep_.CoverTab[12944]++
//line /usr/local/go/src/net/conf.go:250
				return stringsEqualFold(hostname, hn)
//line /usr/local/go/src/net/conf.go:250
				// _ = "end of CoverTab[12944]"
//line /usr/local/go/src/net/conf.go:250
			}() {
//line /usr/local/go/src/net/conf.go:250
				_go_fuzz_dep_.CoverTab[12945]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:251
				// _ = "end of CoverTab[12945]"
			} else {
//line /usr/local/go/src/net/conf.go:252
				_go_fuzz_dep_.CoverTab[12946]++
//line /usr/local/go/src/net/conf.go:252
				// _ = "end of CoverTab[12946]"
//line /usr/local/go/src/net/conf.go:252
			}
//line /usr/local/go/src/net/conf.go:252
			// _ = "end of CoverTab[12938]"
//line /usr/local/go/src/net/conf.go:252
			_go_fuzz_dep_.CoverTab[12939]++
								continue
//line /usr/local/go/src/net/conf.go:253
			// _ = "end of CoverTab[12939]"
		} else {
//line /usr/local/go/src/net/conf.go:254
			_go_fuzz_dep_.CoverTab[12947]++
//line /usr/local/go/src/net/conf.go:254
			// _ = "end of CoverTab[12947]"
//line /usr/local/go/src/net/conf.go:254
		}
//line /usr/local/go/src/net/conf.go:254
		// _ = "end of CoverTab[12933]"
//line /usr/local/go/src/net/conf.go:254
		_go_fuzz_dep_.CoverTab[12934]++
							if src.source == "files" || func() bool {
//line /usr/local/go/src/net/conf.go:255
			_go_fuzz_dep_.CoverTab[12948]++
//line /usr/local/go/src/net/conf.go:255
			return src.source == "dns"
//line /usr/local/go/src/net/conf.go:255
			// _ = "end of CoverTab[12948]"
//line /usr/local/go/src/net/conf.go:255
		}() {
//line /usr/local/go/src/net/conf.go:255
			_go_fuzz_dep_.CoverTab[12949]++
								if !src.standardCriteria() {
//line /usr/local/go/src/net/conf.go:256
				_go_fuzz_dep_.CoverTab[12953]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:257
				// _ = "end of CoverTab[12953]"
			} else {
//line /usr/local/go/src/net/conf.go:258
				_go_fuzz_dep_.CoverTab[12954]++
//line /usr/local/go/src/net/conf.go:258
				// _ = "end of CoverTab[12954]"
//line /usr/local/go/src/net/conf.go:258
			}
//line /usr/local/go/src/net/conf.go:258
			// _ = "end of CoverTab[12949]"
//line /usr/local/go/src/net/conf.go:258
			_go_fuzz_dep_.CoverTab[12950]++
								if src.source == "files" {
//line /usr/local/go/src/net/conf.go:259
				_go_fuzz_dep_.CoverTab[12955]++
									filesSource = true
//line /usr/local/go/src/net/conf.go:260
				// _ = "end of CoverTab[12955]"
			} else {
//line /usr/local/go/src/net/conf.go:261
				_go_fuzz_dep_.CoverTab[12956]++
//line /usr/local/go/src/net/conf.go:261
				if src.source == "dns" {
//line /usr/local/go/src/net/conf.go:261
					_go_fuzz_dep_.CoverTab[12957]++
										dnsSource = true
//line /usr/local/go/src/net/conf.go:262
					// _ = "end of CoverTab[12957]"
				} else {
//line /usr/local/go/src/net/conf.go:263
					_go_fuzz_dep_.CoverTab[12958]++
//line /usr/local/go/src/net/conf.go:263
					// _ = "end of CoverTab[12958]"
//line /usr/local/go/src/net/conf.go:263
				}
//line /usr/local/go/src/net/conf.go:263
				// _ = "end of CoverTab[12956]"
//line /usr/local/go/src/net/conf.go:263
			}
//line /usr/local/go/src/net/conf.go:263
			// _ = "end of CoverTab[12950]"
//line /usr/local/go/src/net/conf.go:263
			_go_fuzz_dep_.CoverTab[12951]++
								if first == "" {
//line /usr/local/go/src/net/conf.go:264
				_go_fuzz_dep_.CoverTab[12959]++
									first = src.source
//line /usr/local/go/src/net/conf.go:265
				// _ = "end of CoverTab[12959]"
			} else {
//line /usr/local/go/src/net/conf.go:266
				_go_fuzz_dep_.CoverTab[12960]++
//line /usr/local/go/src/net/conf.go:266
				// _ = "end of CoverTab[12960]"
//line /usr/local/go/src/net/conf.go:266
			}
//line /usr/local/go/src/net/conf.go:266
			// _ = "end of CoverTab[12951]"
//line /usr/local/go/src/net/conf.go:266
			_go_fuzz_dep_.CoverTab[12952]++
								continue
//line /usr/local/go/src/net/conf.go:267
			// _ = "end of CoverTab[12952]"
		} else {
//line /usr/local/go/src/net/conf.go:268
			_go_fuzz_dep_.CoverTab[12961]++
//line /usr/local/go/src/net/conf.go:268
			// _ = "end of CoverTab[12961]"
//line /usr/local/go/src/net/conf.go:268
		}
//line /usr/local/go/src/net/conf.go:268
		// _ = "end of CoverTab[12934]"
//line /usr/local/go/src/net/conf.go:268
		_go_fuzz_dep_.CoverTab[12935]++
							if stringsHasPrefix(src.source, "mdns") {
//line /usr/local/go/src/net/conf.go:269
			_go_fuzz_dep_.CoverTab[12962]++

//line /usr/local/go/src/net/conf.go:273
			mdnsSource = true
								continue
//line /usr/local/go/src/net/conf.go:274
			// _ = "end of CoverTab[12962]"
		} else {
//line /usr/local/go/src/net/conf.go:275
			_go_fuzz_dep_.CoverTab[12963]++
//line /usr/local/go/src/net/conf.go:275
			// _ = "end of CoverTab[12963]"
//line /usr/local/go/src/net/conf.go:275
		}
//line /usr/local/go/src/net/conf.go:275
		// _ = "end of CoverTab[12935]"
//line /usr/local/go/src/net/conf.go:275
		_go_fuzz_dep_.CoverTab[12936]++

							return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:277
		// _ = "end of CoverTab[12936]"
	}
//line /usr/local/go/src/net/conf.go:278
	// _ = "end of CoverTab[12867]"
//line /usr/local/go/src/net/conf.go:278
	_go_fuzz_dep_.CoverTab[12868]++

//line /usr/local/go/src/net/conf.go:283
	if mdnsSource && func() bool {
//line /usr/local/go/src/net/conf.go:283
		_go_fuzz_dep_.CoverTab[12964]++
//line /usr/local/go/src/net/conf.go:283
		return c.hasMDNSAllow
//line /usr/local/go/src/net/conf.go:283
		// _ = "end of CoverTab[12964]"
//line /usr/local/go/src/net/conf.go:283
	}() {
//line /usr/local/go/src/net/conf.go:283
		_go_fuzz_dep_.CoverTab[12965]++
							return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:284
		// _ = "end of CoverTab[12965]"
	} else {
//line /usr/local/go/src/net/conf.go:285
		_go_fuzz_dep_.CoverTab[12966]++
//line /usr/local/go/src/net/conf.go:285
		// _ = "end of CoverTab[12966]"
//line /usr/local/go/src/net/conf.go:285
	}
//line /usr/local/go/src/net/conf.go:285
	// _ = "end of CoverTab[12868]"
//line /usr/local/go/src/net/conf.go:285
	_go_fuzz_dep_.CoverTab[12869]++

//line /usr/local/go/src/net/conf.go:289
	switch {
	case filesSource && func() bool {
//line /usr/local/go/src/net/conf.go:290
		_go_fuzz_dep_.CoverTab[12971]++
//line /usr/local/go/src/net/conf.go:290
		return dnsSource
//line /usr/local/go/src/net/conf.go:290
		// _ = "end of CoverTab[12971]"
//line /usr/local/go/src/net/conf.go:290
	}():
//line /usr/local/go/src/net/conf.go:290
		_go_fuzz_dep_.CoverTab[12967]++
							if first == "files" {
//line /usr/local/go/src/net/conf.go:291
			_go_fuzz_dep_.CoverTab[12972]++
								return hostLookupFilesDNS, conf
//line /usr/local/go/src/net/conf.go:292
			// _ = "end of CoverTab[12972]"
		} else {
//line /usr/local/go/src/net/conf.go:293
			_go_fuzz_dep_.CoverTab[12973]++
								return hostLookupDNSFiles, conf
//line /usr/local/go/src/net/conf.go:294
			// _ = "end of CoverTab[12973]"
		}
//line /usr/local/go/src/net/conf.go:295
		// _ = "end of CoverTab[12967]"
	case filesSource:
//line /usr/local/go/src/net/conf.go:296
		_go_fuzz_dep_.CoverTab[12968]++
							return hostLookupFiles, conf
//line /usr/local/go/src/net/conf.go:297
		// _ = "end of CoverTab[12968]"
	case dnsSource:
//line /usr/local/go/src/net/conf.go:298
		_go_fuzz_dep_.CoverTab[12969]++
							return hostLookupDNS, conf
//line /usr/local/go/src/net/conf.go:299
		// _ = "end of CoverTab[12969]"
//line /usr/local/go/src/net/conf.go:299
	default:
//line /usr/local/go/src/net/conf.go:299
		_go_fuzz_dep_.CoverTab[12970]++
//line /usr/local/go/src/net/conf.go:299
		// _ = "end of CoverTab[12970]"
	}
//line /usr/local/go/src/net/conf.go:300
	// _ = "end of CoverTab[12869]"
//line /usr/local/go/src/net/conf.go:300
	_go_fuzz_dep_.CoverTab[12870]++

//line /usr/local/go/src/net/conf.go:303
	return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:303
	// _ = "end of CoverTab[12870]"
}

var netdns = godebug.New("netdns")

// goDebugNetDNS parses the value of the GODEBUG "netdns" value.
//line /usr/local/go/src/net/conf.go:308
// The netdns value can be of the form:
//line /usr/local/go/src/net/conf.go:308
//
//line /usr/local/go/src/net/conf.go:308
//	1       // debug level 1
//line /usr/local/go/src/net/conf.go:308
//	2       // debug level 2
//line /usr/local/go/src/net/conf.go:308
//	cgo     // use cgo for DNS lookups
//line /usr/local/go/src/net/conf.go:308
//	go      // use go for DNS lookups
//line /usr/local/go/src/net/conf.go:308
//	cgo+1   // use cgo for DNS lookups + debug level 1
//line /usr/local/go/src/net/conf.go:308
//	1+cgo   // same
//line /usr/local/go/src/net/conf.go:308
//	cgo+2   // same, but debug level 2
//line /usr/local/go/src/net/conf.go:308
//
//line /usr/local/go/src/net/conf.go:308
// etc.
//line /usr/local/go/src/net/conf.go:320
func goDebugNetDNS() (dnsMode string, debugLevel int) {
//line /usr/local/go/src/net/conf.go:320
	_go_fuzz_dep_.CoverTab[12974]++
						goDebug := netdns.Value()
						parsePart := func(s string) {
//line /usr/local/go/src/net/conf.go:322
		_go_fuzz_dep_.CoverTab[12977]++
							if s == "" {
//line /usr/local/go/src/net/conf.go:323
			_go_fuzz_dep_.CoverTab[12979]++
								return
//line /usr/local/go/src/net/conf.go:324
			// _ = "end of CoverTab[12979]"
		} else {
//line /usr/local/go/src/net/conf.go:325
			_go_fuzz_dep_.CoverTab[12980]++
//line /usr/local/go/src/net/conf.go:325
			// _ = "end of CoverTab[12980]"
//line /usr/local/go/src/net/conf.go:325
		}
//line /usr/local/go/src/net/conf.go:325
		// _ = "end of CoverTab[12977]"
//line /usr/local/go/src/net/conf.go:325
		_go_fuzz_dep_.CoverTab[12978]++
							if '0' <= s[0] && func() bool {
//line /usr/local/go/src/net/conf.go:326
			_go_fuzz_dep_.CoverTab[12981]++
//line /usr/local/go/src/net/conf.go:326
			return s[0] <= '9'
//line /usr/local/go/src/net/conf.go:326
			// _ = "end of CoverTab[12981]"
//line /usr/local/go/src/net/conf.go:326
		}() {
//line /usr/local/go/src/net/conf.go:326
			_go_fuzz_dep_.CoverTab[12982]++
								debugLevel, _, _ = dtoi(s)
//line /usr/local/go/src/net/conf.go:327
			// _ = "end of CoverTab[12982]"
		} else {
//line /usr/local/go/src/net/conf.go:328
			_go_fuzz_dep_.CoverTab[12983]++
								dnsMode = s
//line /usr/local/go/src/net/conf.go:329
			// _ = "end of CoverTab[12983]"
		}
//line /usr/local/go/src/net/conf.go:330
		// _ = "end of CoverTab[12978]"
	}
//line /usr/local/go/src/net/conf.go:331
	// _ = "end of CoverTab[12974]"
//line /usr/local/go/src/net/conf.go:331
	_go_fuzz_dep_.CoverTab[12975]++
						if i := bytealg.IndexByteString(goDebug, '+'); i != -1 {
//line /usr/local/go/src/net/conf.go:332
		_go_fuzz_dep_.CoverTab[12984]++
							parsePart(goDebug[:i])
							parsePart(goDebug[i+1:])
							return
//line /usr/local/go/src/net/conf.go:335
		// _ = "end of CoverTab[12984]"
	} else {
//line /usr/local/go/src/net/conf.go:336
		_go_fuzz_dep_.CoverTab[12985]++
//line /usr/local/go/src/net/conf.go:336
		// _ = "end of CoverTab[12985]"
//line /usr/local/go/src/net/conf.go:336
	}
//line /usr/local/go/src/net/conf.go:336
	// _ = "end of CoverTab[12975]"
//line /usr/local/go/src/net/conf.go:336
	_go_fuzz_dep_.CoverTab[12976]++
						parsePart(goDebug)
						return
//line /usr/local/go/src/net/conf.go:338
	// _ = "end of CoverTab[12976]"
}

// isLocalhost reports whether h should be considered a "localhost"
//line /usr/local/go/src/net/conf.go:341
// name for the myhostname NSS module.
//line /usr/local/go/src/net/conf.go:343
func isLocalhost(h string) bool {
//line /usr/local/go/src/net/conf.go:343
	_go_fuzz_dep_.CoverTab[12986]++
						return stringsEqualFold(h, "localhost") || func() bool {
//line /usr/local/go/src/net/conf.go:344
		_go_fuzz_dep_.CoverTab[12987]++
//line /usr/local/go/src/net/conf.go:344
		return stringsEqualFold(h, "localhost.localdomain")
//line /usr/local/go/src/net/conf.go:344
		// _ = "end of CoverTab[12987]"
//line /usr/local/go/src/net/conf.go:344
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:344
		_go_fuzz_dep_.CoverTab[12988]++
//line /usr/local/go/src/net/conf.go:344
		return stringsHasSuffixFold(h, ".localhost")
//line /usr/local/go/src/net/conf.go:344
		// _ = "end of CoverTab[12988]"
//line /usr/local/go/src/net/conf.go:344
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:344
		_go_fuzz_dep_.CoverTab[12989]++
//line /usr/local/go/src/net/conf.go:344
		return stringsHasSuffixFold(h, ".localhost.localdomain")
//line /usr/local/go/src/net/conf.go:344
		// _ = "end of CoverTab[12989]"
//line /usr/local/go/src/net/conf.go:344
	}()
//line /usr/local/go/src/net/conf.go:344
	// _ = "end of CoverTab[12986]"
}

// isGateway reports whether h should be considered a "gateway"
//line /usr/local/go/src/net/conf.go:347
// name for the myhostname NSS module.
//line /usr/local/go/src/net/conf.go:349
func isGateway(h string) bool {
//line /usr/local/go/src/net/conf.go:349
	_go_fuzz_dep_.CoverTab[12990]++
						return stringsEqualFold(h, "_gateway")
//line /usr/local/go/src/net/conf.go:350
	// _ = "end of CoverTab[12990]"
}

// isOutbound reports whether h should be considered a "outbound"
//line /usr/local/go/src/net/conf.go:353
// name for the myhostname NSS module.
//line /usr/local/go/src/net/conf.go:355
func isOutbound(h string) bool {
//line /usr/local/go/src/net/conf.go:355
	_go_fuzz_dep_.CoverTab[12991]++
						return stringsEqualFold(h, "_outbound")
//line /usr/local/go/src/net/conf.go:356
	// _ = "end of CoverTab[12991]"
}

//line /usr/local/go/src/net/conf.go:357
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/conf.go:357
var _ = _go_fuzz_dep_.CoverTab
