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
	_go_fuzz_dep_.CoverTab[4423]++
						confOnce.Do(initConfVal)
						return confVal
//line /usr/local/go/src/net/conf.go:41
	// _ = "end of CoverTab[4423]"
}

func initConfVal() {
//line /usr/local/go/src/net/conf.go:44
	_go_fuzz_dep_.CoverTab[4424]++
						dnsMode, debugLevel := goDebugNetDNS()
						confVal.dnsDebugLevel = debugLevel
						confVal.netGo = netGo || func() bool {
//line /usr/local/go/src/net/conf.go:47
		_go_fuzz_dep_.CoverTab[4431]++
//line /usr/local/go/src/net/conf.go:47
		return dnsMode == "go"
//line /usr/local/go/src/net/conf.go:47
		// _ = "end of CoverTab[4431]"
//line /usr/local/go/src/net/conf.go:47
	}()
						confVal.netCgo = netCgo || func() bool {
//line /usr/local/go/src/net/conf.go:48
		_go_fuzz_dep_.CoverTab[4432]++
//line /usr/local/go/src/net/conf.go:48
		return dnsMode == "cgo"
//line /usr/local/go/src/net/conf.go:48
		// _ = "end of CoverTab[4432]"
//line /usr/local/go/src/net/conf.go:48
	}()
						if !confVal.netGo && func() bool {
//line /usr/local/go/src/net/conf.go:49
		_go_fuzz_dep_.CoverTab[4433]++
//line /usr/local/go/src/net/conf.go:49
		return !confVal.netCgo
//line /usr/local/go/src/net/conf.go:49
		// _ = "end of CoverTab[4433]"
//line /usr/local/go/src/net/conf.go:49
	}() && func() bool {
//line /usr/local/go/src/net/conf.go:49
		_go_fuzz_dep_.CoverTab[4434]++
//line /usr/local/go/src/net/conf.go:49
		return (runtime.GOOS == "windows" || func() bool {
//line /usr/local/go/src/net/conf.go:49
			_go_fuzz_dep_.CoverTab[4435]++
//line /usr/local/go/src/net/conf.go:49
			return runtime.GOOS == "plan9"
//line /usr/local/go/src/net/conf.go:49
			// _ = "end of CoverTab[4435]"
//line /usr/local/go/src/net/conf.go:49
		}())
//line /usr/local/go/src/net/conf.go:49
		// _ = "end of CoverTab[4434]"
//line /usr/local/go/src/net/conf.go:49
	}() {
//line /usr/local/go/src/net/conf.go:49
		_go_fuzz_dep_.CoverTab[4436]++

//line /usr/local/go/src/net/conf.go:58
		confVal.netCgo = true
//line /usr/local/go/src/net/conf.go:58
		// _ = "end of CoverTab[4436]"
	} else {
//line /usr/local/go/src/net/conf.go:59
		_go_fuzz_dep_.CoverTab[4437]++
//line /usr/local/go/src/net/conf.go:59
		// _ = "end of CoverTab[4437]"
//line /usr/local/go/src/net/conf.go:59
	}
//line /usr/local/go/src/net/conf.go:59
	// _ = "end of CoverTab[4424]"
//line /usr/local/go/src/net/conf.go:59
	_go_fuzz_dep_.CoverTab[4425]++

						if confVal.dnsDebugLevel > 0 {
//line /usr/local/go/src/net/conf.go:61
		_go_fuzz_dep_.CoverTab[4438]++
							defer func() {
//line /usr/local/go/src/net/conf.go:62
			_go_fuzz_dep_.CoverTab[4439]++
								if confVal.dnsDebugLevel > 1 {
//line /usr/local/go/src/net/conf.go:63
				_go_fuzz_dep_.CoverTab[4441]++
									println("go package net: confVal.netCgo =", confVal.netCgo, " netGo =", confVal.netGo)
//line /usr/local/go/src/net/conf.go:64
				// _ = "end of CoverTab[4441]"
			} else {
//line /usr/local/go/src/net/conf.go:65
				_go_fuzz_dep_.CoverTab[4442]++
//line /usr/local/go/src/net/conf.go:65
				// _ = "end of CoverTab[4442]"
//line /usr/local/go/src/net/conf.go:65
			}
//line /usr/local/go/src/net/conf.go:65
			// _ = "end of CoverTab[4439]"
//line /usr/local/go/src/net/conf.go:65
			_go_fuzz_dep_.CoverTab[4440]++
								switch {
			case confVal.netGo:
//line /usr/local/go/src/net/conf.go:67
				_go_fuzz_dep_.CoverTab[4443]++
									if netGo {
//line /usr/local/go/src/net/conf.go:68
					_go_fuzz_dep_.CoverTab[4446]++
										println("go package net: built with netgo build tag; using Go's DNS resolver")
//line /usr/local/go/src/net/conf.go:69
					// _ = "end of CoverTab[4446]"
				} else {
//line /usr/local/go/src/net/conf.go:70
					_go_fuzz_dep_.CoverTab[4447]++
										println("go package net: GODEBUG setting forcing use of Go's resolver")
//line /usr/local/go/src/net/conf.go:71
					// _ = "end of CoverTab[4447]"
				}
//line /usr/local/go/src/net/conf.go:72
				// _ = "end of CoverTab[4443]"
			case confVal.forceCgoLookupHost:
//line /usr/local/go/src/net/conf.go:73
				_go_fuzz_dep_.CoverTab[4444]++
									println("go package net: using cgo DNS resolver")
//line /usr/local/go/src/net/conf.go:74
				// _ = "end of CoverTab[4444]"
			default:
//line /usr/local/go/src/net/conf.go:75
				_go_fuzz_dep_.CoverTab[4445]++
									println("go package net: dynamic selection of DNS resolver")
//line /usr/local/go/src/net/conf.go:76
				// _ = "end of CoverTab[4445]"
			}
//line /usr/local/go/src/net/conf.go:77
			// _ = "end of CoverTab[4440]"
		}()
//line /usr/local/go/src/net/conf.go:78
		// _ = "end of CoverTab[4438]"
	} else {
//line /usr/local/go/src/net/conf.go:79
		_go_fuzz_dep_.CoverTab[4448]++
//line /usr/local/go/src/net/conf.go:79
		// _ = "end of CoverTab[4448]"
//line /usr/local/go/src/net/conf.go:79
	}
//line /usr/local/go/src/net/conf.go:79
	// _ = "end of CoverTab[4425]"
//line /usr/local/go/src/net/conf.go:79
	_go_fuzz_dep_.CoverTab[4426]++

//line /usr/local/go/src/net/conf.go:84
	if runtime.GOOS == "darwin" || func() bool {
//line /usr/local/go/src/net/conf.go:84
		_go_fuzz_dep_.CoverTab[4449]++
//line /usr/local/go/src/net/conf.go:84
		return runtime.GOOS == "ios"
//line /usr/local/go/src/net/conf.go:84
		// _ = "end of CoverTab[4449]"
//line /usr/local/go/src/net/conf.go:84
	}() {
//line /usr/local/go/src/net/conf.go:84
		_go_fuzz_dep_.CoverTab[4450]++
							confVal.forceCgoLookupHost = true
							return
//line /usr/local/go/src/net/conf.go:86
		// _ = "end of CoverTab[4450]"
	} else {
//line /usr/local/go/src/net/conf.go:87
		_go_fuzz_dep_.CoverTab[4451]++
//line /usr/local/go/src/net/conf.go:87
		// _ = "end of CoverTab[4451]"
//line /usr/local/go/src/net/conf.go:87
	}
//line /usr/local/go/src/net/conf.go:87
	// _ = "end of CoverTab[4426]"
//line /usr/local/go/src/net/conf.go:87
	_go_fuzz_dep_.CoverTab[4427]++

						if runtime.GOOS == "windows" || func() bool {
//line /usr/local/go/src/net/conf.go:89
		_go_fuzz_dep_.CoverTab[4452]++
//line /usr/local/go/src/net/conf.go:89
		return runtime.GOOS == "plan9"
//line /usr/local/go/src/net/conf.go:89
		// _ = "end of CoverTab[4452]"
//line /usr/local/go/src/net/conf.go:89
	}() {
//line /usr/local/go/src/net/conf.go:89
		_go_fuzz_dep_.CoverTab[4453]++
							return
//line /usr/local/go/src/net/conf.go:90
		// _ = "end of CoverTab[4453]"
	} else {
//line /usr/local/go/src/net/conf.go:91
		_go_fuzz_dep_.CoverTab[4454]++
//line /usr/local/go/src/net/conf.go:91
		// _ = "end of CoverTab[4454]"
//line /usr/local/go/src/net/conf.go:91
	}
//line /usr/local/go/src/net/conf.go:91
	// _ = "end of CoverTab[4427]"
//line /usr/local/go/src/net/conf.go:91
	_go_fuzz_dep_.CoverTab[4428]++

//line /usr/local/go/src/net/conf.go:96
	_, localDomainDefined := syscall.Getenv("LOCALDOMAIN")
	if os.Getenv("RES_OPTIONS") != "" || func() bool {
//line /usr/local/go/src/net/conf.go:97
		_go_fuzz_dep_.CoverTab[4455]++
//line /usr/local/go/src/net/conf.go:97
		return os.Getenv("HOSTALIASES") != ""
							// _ = "end of CoverTab[4455]"
//line /usr/local/go/src/net/conf.go:98
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:98
		_go_fuzz_dep_.CoverTab[4456]++
//line /usr/local/go/src/net/conf.go:98
		return confVal.netCgo
							// _ = "end of CoverTab[4456]"
//line /usr/local/go/src/net/conf.go:99
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:99
		_go_fuzz_dep_.CoverTab[4457]++
//line /usr/local/go/src/net/conf.go:99
		return localDomainDefined
							// _ = "end of CoverTab[4457]"
//line /usr/local/go/src/net/conf.go:100
	}() {
//line /usr/local/go/src/net/conf.go:100
		_go_fuzz_dep_.CoverTab[4458]++
							confVal.forceCgoLookupHost = true
							return
//line /usr/local/go/src/net/conf.go:102
		// _ = "end of CoverTab[4458]"
	} else {
//line /usr/local/go/src/net/conf.go:103
		_go_fuzz_dep_.CoverTab[4459]++
//line /usr/local/go/src/net/conf.go:103
		// _ = "end of CoverTab[4459]"
//line /usr/local/go/src/net/conf.go:103
	}
//line /usr/local/go/src/net/conf.go:103
	// _ = "end of CoverTab[4428]"
//line /usr/local/go/src/net/conf.go:103
	_go_fuzz_dep_.CoverTab[4429]++

//line /usr/local/go/src/net/conf.go:107
	if runtime.GOOS == "openbsd" && func() bool {
//line /usr/local/go/src/net/conf.go:107
		_go_fuzz_dep_.CoverTab[4460]++
//line /usr/local/go/src/net/conf.go:107
		return os.Getenv("ASR_CONFIG") != ""
//line /usr/local/go/src/net/conf.go:107
		// _ = "end of CoverTab[4460]"
//line /usr/local/go/src/net/conf.go:107
	}() {
//line /usr/local/go/src/net/conf.go:107
		_go_fuzz_dep_.CoverTab[4461]++
							confVal.forceCgoLookupHost = true
							return
//line /usr/local/go/src/net/conf.go:109
		// _ = "end of CoverTab[4461]"
	} else {
//line /usr/local/go/src/net/conf.go:110
		_go_fuzz_dep_.CoverTab[4462]++
//line /usr/local/go/src/net/conf.go:110
		// _ = "end of CoverTab[4462]"
//line /usr/local/go/src/net/conf.go:110
	}
//line /usr/local/go/src/net/conf.go:110
	// _ = "end of CoverTab[4429]"
//line /usr/local/go/src/net/conf.go:110
	_go_fuzz_dep_.CoverTab[4430]++

						if _, err := os.Stat("/etc/mdns.allow"); err == nil {
//line /usr/local/go/src/net/conf.go:112
		_go_fuzz_dep_.CoverTab[4463]++
							confVal.hasMDNSAllow = true
//line /usr/local/go/src/net/conf.go:113
		// _ = "end of CoverTab[4463]"
	} else {
//line /usr/local/go/src/net/conf.go:114
		_go_fuzz_dep_.CoverTab[4464]++
//line /usr/local/go/src/net/conf.go:114
		// _ = "end of CoverTab[4464]"
//line /usr/local/go/src/net/conf.go:114
	}
//line /usr/local/go/src/net/conf.go:114
	// _ = "end of CoverTab[4430]"
}

// canUseCgo reports whether calling cgo functions is allowed
//line /usr/local/go/src/net/conf.go:117
// for non-hostname lookups.
//line /usr/local/go/src/net/conf.go:119
func (c *conf) canUseCgo() bool {
//line /usr/local/go/src/net/conf.go:119
	_go_fuzz_dep_.CoverTab[4465]++
						ret, _ := c.hostLookupOrder(nil, "")
						return ret == hostLookupCgo
//line /usr/local/go/src/net/conf.go:121
	// _ = "end of CoverTab[4465]"
}

// hostLookupOrder determines which strategy to use to resolve hostname.
//line /usr/local/go/src/net/conf.go:124
// The provided Resolver is optional. nil means to not consider its options.
//line /usr/local/go/src/net/conf.go:124
// It also returns dnsConfig when it was used to determine the lookup order.
//line /usr/local/go/src/net/conf.go:127
func (c *conf) hostLookupOrder(r *Resolver, hostname string) (ret hostLookupOrder, dnsConfig *dnsConfig) {
//line /usr/local/go/src/net/conf.go:127
	_go_fuzz_dep_.CoverTab[4466]++
						if c.dnsDebugLevel > 1 {
//line /usr/local/go/src/net/conf.go:128
		_go_fuzz_dep_.CoverTab[4481]++
							defer func() {
//line /usr/local/go/src/net/conf.go:129
			_go_fuzz_dep_.CoverTab[4482]++
								print("go package net: hostLookupOrder(", hostname, ") = ", ret.String(), "\n")
//line /usr/local/go/src/net/conf.go:130
			// _ = "end of CoverTab[4482]"
		}()
//line /usr/local/go/src/net/conf.go:131
		// _ = "end of CoverTab[4481]"
	} else {
//line /usr/local/go/src/net/conf.go:132
		_go_fuzz_dep_.CoverTab[4483]++
//line /usr/local/go/src/net/conf.go:132
		// _ = "end of CoverTab[4483]"
//line /usr/local/go/src/net/conf.go:132
	}
//line /usr/local/go/src/net/conf.go:132
	// _ = "end of CoverTab[4466]"
//line /usr/local/go/src/net/conf.go:132
	_go_fuzz_dep_.CoverTab[4467]++
						fallbackOrder := hostLookupCgo
						if c.netGo || func() bool {
//line /usr/local/go/src/net/conf.go:134
		_go_fuzz_dep_.CoverTab[4484]++
//line /usr/local/go/src/net/conf.go:134
		return r.preferGo()
//line /usr/local/go/src/net/conf.go:134
		// _ = "end of CoverTab[4484]"
//line /usr/local/go/src/net/conf.go:134
	}() {
//line /usr/local/go/src/net/conf.go:134
		_go_fuzz_dep_.CoverTab[4485]++
							switch c.goos {
		case "windows":
//line /usr/local/go/src/net/conf.go:136
			_go_fuzz_dep_.CoverTab[4486]++

//line /usr/local/go/src/net/conf.go:141
			fallbackOrder = hostLookupDNS
//line /usr/local/go/src/net/conf.go:141
			// _ = "end of CoverTab[4486]"
		default:
//line /usr/local/go/src/net/conf.go:142
			_go_fuzz_dep_.CoverTab[4487]++
								fallbackOrder = hostLookupFilesDNS
//line /usr/local/go/src/net/conf.go:143
			// _ = "end of CoverTab[4487]"
		}
//line /usr/local/go/src/net/conf.go:144
		// _ = "end of CoverTab[4485]"
	} else {
//line /usr/local/go/src/net/conf.go:145
		_go_fuzz_dep_.CoverTab[4488]++
//line /usr/local/go/src/net/conf.go:145
		// _ = "end of CoverTab[4488]"
//line /usr/local/go/src/net/conf.go:145
	}
//line /usr/local/go/src/net/conf.go:145
	// _ = "end of CoverTab[4467]"
//line /usr/local/go/src/net/conf.go:145
	_go_fuzz_dep_.CoverTab[4468]++
						if c.forceCgoLookupHost || func() bool {
//line /usr/local/go/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[4489]++
//line /usr/local/go/src/net/conf.go:146
		return c.goos == "android"
//line /usr/local/go/src/net/conf.go:146
		// _ = "end of CoverTab[4489]"
//line /usr/local/go/src/net/conf.go:146
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[4490]++
//line /usr/local/go/src/net/conf.go:146
		return c.goos == "windows"
//line /usr/local/go/src/net/conf.go:146
		// _ = "end of CoverTab[4490]"
//line /usr/local/go/src/net/conf.go:146
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[4491]++
//line /usr/local/go/src/net/conf.go:146
		return c.goos == "plan9"
//line /usr/local/go/src/net/conf.go:146
		// _ = "end of CoverTab[4491]"
//line /usr/local/go/src/net/conf.go:146
	}() {
//line /usr/local/go/src/net/conf.go:146
		_go_fuzz_dep_.CoverTab[4492]++
							return fallbackOrder, nil
//line /usr/local/go/src/net/conf.go:147
		// _ = "end of CoverTab[4492]"
	} else {
//line /usr/local/go/src/net/conf.go:148
		_go_fuzz_dep_.CoverTab[4493]++
//line /usr/local/go/src/net/conf.go:148
		// _ = "end of CoverTab[4493]"
//line /usr/local/go/src/net/conf.go:148
	}
//line /usr/local/go/src/net/conf.go:148
	// _ = "end of CoverTab[4468]"
//line /usr/local/go/src/net/conf.go:148
	_go_fuzz_dep_.CoverTab[4469]++
						if bytealg.IndexByteString(hostname, '\\') != -1 || func() bool {
//line /usr/local/go/src/net/conf.go:149
		_go_fuzz_dep_.CoverTab[4494]++
//line /usr/local/go/src/net/conf.go:149
		return bytealg.IndexByteString(hostname, '%') != -1
//line /usr/local/go/src/net/conf.go:149
		// _ = "end of CoverTab[4494]"
//line /usr/local/go/src/net/conf.go:149
	}() {
//line /usr/local/go/src/net/conf.go:149
		_go_fuzz_dep_.CoverTab[4495]++

//line /usr/local/go/src/net/conf.go:152
		return fallbackOrder, nil
//line /usr/local/go/src/net/conf.go:152
		// _ = "end of CoverTab[4495]"
	} else {
//line /usr/local/go/src/net/conf.go:153
		_go_fuzz_dep_.CoverTab[4496]++
//line /usr/local/go/src/net/conf.go:153
		// _ = "end of CoverTab[4496]"
//line /usr/local/go/src/net/conf.go:153
	}
//line /usr/local/go/src/net/conf.go:153
	// _ = "end of CoverTab[4469]"
//line /usr/local/go/src/net/conf.go:153
	_go_fuzz_dep_.CoverTab[4470]++

						conf := getSystemDNSConfig()
						if conf.err != nil && func() bool {
//line /usr/local/go/src/net/conf.go:156
		_go_fuzz_dep_.CoverTab[4497]++
//line /usr/local/go/src/net/conf.go:156
		return !os.IsNotExist(conf.err)
//line /usr/local/go/src/net/conf.go:156
		// _ = "end of CoverTab[4497]"
//line /usr/local/go/src/net/conf.go:156
	}() && func() bool {
//line /usr/local/go/src/net/conf.go:156
		_go_fuzz_dep_.CoverTab[4498]++
//line /usr/local/go/src/net/conf.go:156
		return !os.IsPermission(conf.err)
//line /usr/local/go/src/net/conf.go:156
		// _ = "end of CoverTab[4498]"
//line /usr/local/go/src/net/conf.go:156
	}() {
//line /usr/local/go/src/net/conf.go:156
		_go_fuzz_dep_.CoverTab[4499]++

//line /usr/local/go/src/net/conf.go:161
		return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:161
		// _ = "end of CoverTab[4499]"
	} else {
//line /usr/local/go/src/net/conf.go:162
		_go_fuzz_dep_.CoverTab[4500]++
//line /usr/local/go/src/net/conf.go:162
		// _ = "end of CoverTab[4500]"
//line /usr/local/go/src/net/conf.go:162
	}
//line /usr/local/go/src/net/conf.go:162
	// _ = "end of CoverTab[4470]"
//line /usr/local/go/src/net/conf.go:162
	_go_fuzz_dep_.CoverTab[4471]++

						if conf.unknownOpt {
//line /usr/local/go/src/net/conf.go:164
		_go_fuzz_dep_.CoverTab[4501]++
							return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:165
		// _ = "end of CoverTab[4501]"
	} else {
//line /usr/local/go/src/net/conf.go:166
		_go_fuzz_dep_.CoverTab[4502]++
//line /usr/local/go/src/net/conf.go:166
		// _ = "end of CoverTab[4502]"
//line /usr/local/go/src/net/conf.go:166
	}
//line /usr/local/go/src/net/conf.go:166
	// _ = "end of CoverTab[4471]"
//line /usr/local/go/src/net/conf.go:166
	_go_fuzz_dep_.CoverTab[4472]++

//line /usr/local/go/src/net/conf.go:170
	if c.goos == "openbsd" {
//line /usr/local/go/src/net/conf.go:170
		_go_fuzz_dep_.CoverTab[4503]++

//line /usr/local/go/src/net/conf.go:174
		if os.IsNotExist(conf.err) {
//line /usr/local/go/src/net/conf.go:174
			_go_fuzz_dep_.CoverTab[4507]++
								return hostLookupFiles, conf
//line /usr/local/go/src/net/conf.go:175
			// _ = "end of CoverTab[4507]"
		} else {
//line /usr/local/go/src/net/conf.go:176
			_go_fuzz_dep_.CoverTab[4508]++
//line /usr/local/go/src/net/conf.go:176
			// _ = "end of CoverTab[4508]"
//line /usr/local/go/src/net/conf.go:176
		}
//line /usr/local/go/src/net/conf.go:176
		// _ = "end of CoverTab[4503]"
//line /usr/local/go/src/net/conf.go:176
		_go_fuzz_dep_.CoverTab[4504]++

							lookup := conf.lookup
							if len(lookup) == 0 {
//line /usr/local/go/src/net/conf.go:179
			_go_fuzz_dep_.CoverTab[4509]++

//line /usr/local/go/src/net/conf.go:184
			return hostLookupDNSFiles, conf
//line /usr/local/go/src/net/conf.go:184
			// _ = "end of CoverTab[4509]"
		} else {
//line /usr/local/go/src/net/conf.go:185
			_go_fuzz_dep_.CoverTab[4510]++
//line /usr/local/go/src/net/conf.go:185
			// _ = "end of CoverTab[4510]"
//line /usr/local/go/src/net/conf.go:185
		}
//line /usr/local/go/src/net/conf.go:185
		// _ = "end of CoverTab[4504]"
//line /usr/local/go/src/net/conf.go:185
		_go_fuzz_dep_.CoverTab[4505]++
							if len(lookup) < 1 || func() bool {
//line /usr/local/go/src/net/conf.go:186
			_go_fuzz_dep_.CoverTab[4511]++
//line /usr/local/go/src/net/conf.go:186
			return len(lookup) > 2
//line /usr/local/go/src/net/conf.go:186
			// _ = "end of CoverTab[4511]"
//line /usr/local/go/src/net/conf.go:186
		}() {
//line /usr/local/go/src/net/conf.go:186
			_go_fuzz_dep_.CoverTab[4512]++
								return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:187
			// _ = "end of CoverTab[4512]"
		} else {
//line /usr/local/go/src/net/conf.go:188
			_go_fuzz_dep_.CoverTab[4513]++
//line /usr/local/go/src/net/conf.go:188
			// _ = "end of CoverTab[4513]"
//line /usr/local/go/src/net/conf.go:188
		}
//line /usr/local/go/src/net/conf.go:188
		// _ = "end of CoverTab[4505]"
//line /usr/local/go/src/net/conf.go:188
		_go_fuzz_dep_.CoverTab[4506]++
							switch lookup[0] {
		case "bind":
//line /usr/local/go/src/net/conf.go:190
			_go_fuzz_dep_.CoverTab[4514]++
								if len(lookup) == 2 {
//line /usr/local/go/src/net/conf.go:191
				_go_fuzz_dep_.CoverTab[4519]++
									if lookup[1] == "file" {
//line /usr/local/go/src/net/conf.go:192
					_go_fuzz_dep_.CoverTab[4521]++
										return hostLookupDNSFiles, conf
//line /usr/local/go/src/net/conf.go:193
					// _ = "end of CoverTab[4521]"
				} else {
//line /usr/local/go/src/net/conf.go:194
					_go_fuzz_dep_.CoverTab[4522]++
//line /usr/local/go/src/net/conf.go:194
					// _ = "end of CoverTab[4522]"
//line /usr/local/go/src/net/conf.go:194
				}
//line /usr/local/go/src/net/conf.go:194
				// _ = "end of CoverTab[4519]"
//line /usr/local/go/src/net/conf.go:194
				_go_fuzz_dep_.CoverTab[4520]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:195
				// _ = "end of CoverTab[4520]"
			} else {
//line /usr/local/go/src/net/conf.go:196
				_go_fuzz_dep_.CoverTab[4523]++
//line /usr/local/go/src/net/conf.go:196
				// _ = "end of CoverTab[4523]"
//line /usr/local/go/src/net/conf.go:196
			}
//line /usr/local/go/src/net/conf.go:196
			// _ = "end of CoverTab[4514]"
//line /usr/local/go/src/net/conf.go:196
			_go_fuzz_dep_.CoverTab[4515]++
								return hostLookupDNS, conf
//line /usr/local/go/src/net/conf.go:197
			// _ = "end of CoverTab[4515]"
		case "file":
//line /usr/local/go/src/net/conf.go:198
			_go_fuzz_dep_.CoverTab[4516]++
								if len(lookup) == 2 {
//line /usr/local/go/src/net/conf.go:199
				_go_fuzz_dep_.CoverTab[4524]++
									if lookup[1] == "bind" {
//line /usr/local/go/src/net/conf.go:200
					_go_fuzz_dep_.CoverTab[4526]++
										return hostLookupFilesDNS, conf
//line /usr/local/go/src/net/conf.go:201
					// _ = "end of CoverTab[4526]"
				} else {
//line /usr/local/go/src/net/conf.go:202
					_go_fuzz_dep_.CoverTab[4527]++
//line /usr/local/go/src/net/conf.go:202
					// _ = "end of CoverTab[4527]"
//line /usr/local/go/src/net/conf.go:202
				}
//line /usr/local/go/src/net/conf.go:202
				// _ = "end of CoverTab[4524]"
//line /usr/local/go/src/net/conf.go:202
				_go_fuzz_dep_.CoverTab[4525]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:203
				// _ = "end of CoverTab[4525]"
			} else {
//line /usr/local/go/src/net/conf.go:204
				_go_fuzz_dep_.CoverTab[4528]++
//line /usr/local/go/src/net/conf.go:204
				// _ = "end of CoverTab[4528]"
//line /usr/local/go/src/net/conf.go:204
			}
//line /usr/local/go/src/net/conf.go:204
			// _ = "end of CoverTab[4516]"
//line /usr/local/go/src/net/conf.go:204
			_go_fuzz_dep_.CoverTab[4517]++
								return hostLookupFiles, conf
//line /usr/local/go/src/net/conf.go:205
			// _ = "end of CoverTab[4517]"
		default:
//line /usr/local/go/src/net/conf.go:206
			_go_fuzz_dep_.CoverTab[4518]++
								return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:207
			// _ = "end of CoverTab[4518]"
		}
//line /usr/local/go/src/net/conf.go:208
		// _ = "end of CoverTab[4506]"
	} else {
//line /usr/local/go/src/net/conf.go:209
		_go_fuzz_dep_.CoverTab[4529]++
//line /usr/local/go/src/net/conf.go:209
		// _ = "end of CoverTab[4529]"
//line /usr/local/go/src/net/conf.go:209
	}
//line /usr/local/go/src/net/conf.go:209
	// _ = "end of CoverTab[4472]"
//line /usr/local/go/src/net/conf.go:209
	_go_fuzz_dep_.CoverTab[4473]++

//line /usr/local/go/src/net/conf.go:212
	if stringsHasSuffix(hostname, ".") {
//line /usr/local/go/src/net/conf.go:212
		_go_fuzz_dep_.CoverTab[4530]++
							hostname = hostname[:len(hostname)-1]
//line /usr/local/go/src/net/conf.go:213
		// _ = "end of CoverTab[4530]"
	} else {
//line /usr/local/go/src/net/conf.go:214
		_go_fuzz_dep_.CoverTab[4531]++
//line /usr/local/go/src/net/conf.go:214
		// _ = "end of CoverTab[4531]"
//line /usr/local/go/src/net/conf.go:214
	}
//line /usr/local/go/src/net/conf.go:214
	// _ = "end of CoverTab[4473]"
//line /usr/local/go/src/net/conf.go:214
	_go_fuzz_dep_.CoverTab[4474]++
						if stringsHasSuffixFold(hostname, ".local") {
//line /usr/local/go/src/net/conf.go:215
		_go_fuzz_dep_.CoverTab[4532]++

//line /usr/local/go/src/net/conf.go:220
		return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:220
		// _ = "end of CoverTab[4532]"
	} else {
//line /usr/local/go/src/net/conf.go:221
		_go_fuzz_dep_.CoverTab[4533]++
//line /usr/local/go/src/net/conf.go:221
		// _ = "end of CoverTab[4533]"
//line /usr/local/go/src/net/conf.go:221
	}
//line /usr/local/go/src/net/conf.go:221
	// _ = "end of CoverTab[4474]"
//line /usr/local/go/src/net/conf.go:221
	_go_fuzz_dep_.CoverTab[4475]++

						nss := getSystemNSS()
						srcs := nss.sources["hosts"]

//line /usr/local/go/src/net/conf.go:227
	if os.IsNotExist(nss.err) || func() bool {
//line /usr/local/go/src/net/conf.go:227
		_go_fuzz_dep_.CoverTab[4534]++
//line /usr/local/go/src/net/conf.go:227
		return (nss.err == nil && func() bool {
//line /usr/local/go/src/net/conf.go:227
			_go_fuzz_dep_.CoverTab[4535]++
//line /usr/local/go/src/net/conf.go:227
			return len(srcs) == 0
//line /usr/local/go/src/net/conf.go:227
			// _ = "end of CoverTab[4535]"
//line /usr/local/go/src/net/conf.go:227
		}())
//line /usr/local/go/src/net/conf.go:227
		// _ = "end of CoverTab[4534]"
//line /usr/local/go/src/net/conf.go:227
	}() {
//line /usr/local/go/src/net/conf.go:227
		_go_fuzz_dep_.CoverTab[4536]++
							if c.goos == "solaris" {
//line /usr/local/go/src/net/conf.go:228
			_go_fuzz_dep_.CoverTab[4538]++

								return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:230
			// _ = "end of CoverTab[4538]"
		} else {
//line /usr/local/go/src/net/conf.go:231
			_go_fuzz_dep_.CoverTab[4539]++
//line /usr/local/go/src/net/conf.go:231
			// _ = "end of CoverTab[4539]"
//line /usr/local/go/src/net/conf.go:231
		}
//line /usr/local/go/src/net/conf.go:231
		// _ = "end of CoverTab[4536]"
//line /usr/local/go/src/net/conf.go:231
		_go_fuzz_dep_.CoverTab[4537]++

							return hostLookupFilesDNS, conf
//line /usr/local/go/src/net/conf.go:233
		// _ = "end of CoverTab[4537]"
	} else {
//line /usr/local/go/src/net/conf.go:234
		_go_fuzz_dep_.CoverTab[4540]++
//line /usr/local/go/src/net/conf.go:234
		// _ = "end of CoverTab[4540]"
//line /usr/local/go/src/net/conf.go:234
	}
//line /usr/local/go/src/net/conf.go:234
	// _ = "end of CoverTab[4475]"
//line /usr/local/go/src/net/conf.go:234
	_go_fuzz_dep_.CoverTab[4476]++
						if nss.err != nil {
//line /usr/local/go/src/net/conf.go:235
		_go_fuzz_dep_.CoverTab[4541]++

//line /usr/local/go/src/net/conf.go:239
		return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:239
		// _ = "end of CoverTab[4541]"
	} else {
//line /usr/local/go/src/net/conf.go:240
		_go_fuzz_dep_.CoverTab[4542]++
//line /usr/local/go/src/net/conf.go:240
		// _ = "end of CoverTab[4542]"
//line /usr/local/go/src/net/conf.go:240
	}
//line /usr/local/go/src/net/conf.go:240
	// _ = "end of CoverTab[4476]"
//line /usr/local/go/src/net/conf.go:240
	_go_fuzz_dep_.CoverTab[4477]++

						var mdnsSource, filesSource, dnsSource bool
						var first string
						for _, src := range srcs {
//line /usr/local/go/src/net/conf.go:244
		_go_fuzz_dep_.CoverTab[4543]++
							if src.source == "myhostname" {
//line /usr/local/go/src/net/conf.go:245
			_go_fuzz_dep_.CoverTab[4547]++
								if isLocalhost(hostname) || func() bool {
//line /usr/local/go/src/net/conf.go:246
				_go_fuzz_dep_.CoverTab[4550]++
//line /usr/local/go/src/net/conf.go:246
				return isGateway(hostname)
//line /usr/local/go/src/net/conf.go:246
				// _ = "end of CoverTab[4550]"
//line /usr/local/go/src/net/conf.go:246
			}() || func() bool {
//line /usr/local/go/src/net/conf.go:246
				_go_fuzz_dep_.CoverTab[4551]++
//line /usr/local/go/src/net/conf.go:246
				return isOutbound(hostname)
//line /usr/local/go/src/net/conf.go:246
				// _ = "end of CoverTab[4551]"
//line /usr/local/go/src/net/conf.go:246
			}() {
//line /usr/local/go/src/net/conf.go:246
				_go_fuzz_dep_.CoverTab[4552]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:247
				// _ = "end of CoverTab[4552]"
			} else {
//line /usr/local/go/src/net/conf.go:248
				_go_fuzz_dep_.CoverTab[4553]++
//line /usr/local/go/src/net/conf.go:248
				// _ = "end of CoverTab[4553]"
//line /usr/local/go/src/net/conf.go:248
			}
//line /usr/local/go/src/net/conf.go:248
			// _ = "end of CoverTab[4547]"
//line /usr/local/go/src/net/conf.go:248
			_go_fuzz_dep_.CoverTab[4548]++
								hn, err := getHostname()
								if err != nil || func() bool {
//line /usr/local/go/src/net/conf.go:250
				_go_fuzz_dep_.CoverTab[4554]++
//line /usr/local/go/src/net/conf.go:250
				return stringsEqualFold(hostname, hn)
//line /usr/local/go/src/net/conf.go:250
				// _ = "end of CoverTab[4554]"
//line /usr/local/go/src/net/conf.go:250
			}() {
//line /usr/local/go/src/net/conf.go:250
				_go_fuzz_dep_.CoverTab[4555]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:251
				// _ = "end of CoverTab[4555]"
			} else {
//line /usr/local/go/src/net/conf.go:252
				_go_fuzz_dep_.CoverTab[4556]++
//line /usr/local/go/src/net/conf.go:252
				// _ = "end of CoverTab[4556]"
//line /usr/local/go/src/net/conf.go:252
			}
//line /usr/local/go/src/net/conf.go:252
			// _ = "end of CoverTab[4548]"
//line /usr/local/go/src/net/conf.go:252
			_go_fuzz_dep_.CoverTab[4549]++
								continue
//line /usr/local/go/src/net/conf.go:253
			// _ = "end of CoverTab[4549]"
		} else {
//line /usr/local/go/src/net/conf.go:254
			_go_fuzz_dep_.CoverTab[4557]++
//line /usr/local/go/src/net/conf.go:254
			// _ = "end of CoverTab[4557]"
//line /usr/local/go/src/net/conf.go:254
		}
//line /usr/local/go/src/net/conf.go:254
		// _ = "end of CoverTab[4543]"
//line /usr/local/go/src/net/conf.go:254
		_go_fuzz_dep_.CoverTab[4544]++
							if src.source == "files" || func() bool {
//line /usr/local/go/src/net/conf.go:255
			_go_fuzz_dep_.CoverTab[4558]++
//line /usr/local/go/src/net/conf.go:255
			return src.source == "dns"
//line /usr/local/go/src/net/conf.go:255
			// _ = "end of CoverTab[4558]"
//line /usr/local/go/src/net/conf.go:255
		}() {
//line /usr/local/go/src/net/conf.go:255
			_go_fuzz_dep_.CoverTab[4559]++
								if !src.standardCriteria() {
//line /usr/local/go/src/net/conf.go:256
				_go_fuzz_dep_.CoverTab[4563]++
									return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:257
				// _ = "end of CoverTab[4563]"
			} else {
//line /usr/local/go/src/net/conf.go:258
				_go_fuzz_dep_.CoverTab[4564]++
//line /usr/local/go/src/net/conf.go:258
				// _ = "end of CoverTab[4564]"
//line /usr/local/go/src/net/conf.go:258
			}
//line /usr/local/go/src/net/conf.go:258
			// _ = "end of CoverTab[4559]"
//line /usr/local/go/src/net/conf.go:258
			_go_fuzz_dep_.CoverTab[4560]++
								if src.source == "files" {
//line /usr/local/go/src/net/conf.go:259
				_go_fuzz_dep_.CoverTab[4565]++
									filesSource = true
//line /usr/local/go/src/net/conf.go:260
				// _ = "end of CoverTab[4565]"
			} else {
//line /usr/local/go/src/net/conf.go:261
				_go_fuzz_dep_.CoverTab[4566]++
//line /usr/local/go/src/net/conf.go:261
				if src.source == "dns" {
//line /usr/local/go/src/net/conf.go:261
					_go_fuzz_dep_.CoverTab[4567]++
										dnsSource = true
//line /usr/local/go/src/net/conf.go:262
					// _ = "end of CoverTab[4567]"
				} else {
//line /usr/local/go/src/net/conf.go:263
					_go_fuzz_dep_.CoverTab[4568]++
//line /usr/local/go/src/net/conf.go:263
					// _ = "end of CoverTab[4568]"
//line /usr/local/go/src/net/conf.go:263
				}
//line /usr/local/go/src/net/conf.go:263
				// _ = "end of CoverTab[4566]"
//line /usr/local/go/src/net/conf.go:263
			}
//line /usr/local/go/src/net/conf.go:263
			// _ = "end of CoverTab[4560]"
//line /usr/local/go/src/net/conf.go:263
			_go_fuzz_dep_.CoverTab[4561]++
								if first == "" {
//line /usr/local/go/src/net/conf.go:264
				_go_fuzz_dep_.CoverTab[4569]++
									first = src.source
//line /usr/local/go/src/net/conf.go:265
				// _ = "end of CoverTab[4569]"
			} else {
//line /usr/local/go/src/net/conf.go:266
				_go_fuzz_dep_.CoverTab[4570]++
//line /usr/local/go/src/net/conf.go:266
				// _ = "end of CoverTab[4570]"
//line /usr/local/go/src/net/conf.go:266
			}
//line /usr/local/go/src/net/conf.go:266
			// _ = "end of CoverTab[4561]"
//line /usr/local/go/src/net/conf.go:266
			_go_fuzz_dep_.CoverTab[4562]++
								continue
//line /usr/local/go/src/net/conf.go:267
			// _ = "end of CoverTab[4562]"
		} else {
//line /usr/local/go/src/net/conf.go:268
			_go_fuzz_dep_.CoverTab[4571]++
//line /usr/local/go/src/net/conf.go:268
			// _ = "end of CoverTab[4571]"
//line /usr/local/go/src/net/conf.go:268
		}
//line /usr/local/go/src/net/conf.go:268
		// _ = "end of CoverTab[4544]"
//line /usr/local/go/src/net/conf.go:268
		_go_fuzz_dep_.CoverTab[4545]++
							if stringsHasPrefix(src.source, "mdns") {
//line /usr/local/go/src/net/conf.go:269
			_go_fuzz_dep_.CoverTab[4572]++

//line /usr/local/go/src/net/conf.go:273
			mdnsSource = true
								continue
//line /usr/local/go/src/net/conf.go:274
			// _ = "end of CoverTab[4572]"
		} else {
//line /usr/local/go/src/net/conf.go:275
			_go_fuzz_dep_.CoverTab[4573]++
//line /usr/local/go/src/net/conf.go:275
			// _ = "end of CoverTab[4573]"
//line /usr/local/go/src/net/conf.go:275
		}
//line /usr/local/go/src/net/conf.go:275
		// _ = "end of CoverTab[4545]"
//line /usr/local/go/src/net/conf.go:275
		_go_fuzz_dep_.CoverTab[4546]++

							return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:277
		// _ = "end of CoverTab[4546]"
	}
//line /usr/local/go/src/net/conf.go:278
	// _ = "end of CoverTab[4477]"
//line /usr/local/go/src/net/conf.go:278
	_go_fuzz_dep_.CoverTab[4478]++

//line /usr/local/go/src/net/conf.go:283
	if mdnsSource && func() bool {
//line /usr/local/go/src/net/conf.go:283
		_go_fuzz_dep_.CoverTab[4574]++
//line /usr/local/go/src/net/conf.go:283
		return c.hasMDNSAllow
//line /usr/local/go/src/net/conf.go:283
		// _ = "end of CoverTab[4574]"
//line /usr/local/go/src/net/conf.go:283
	}() {
//line /usr/local/go/src/net/conf.go:283
		_go_fuzz_dep_.CoverTab[4575]++
							return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:284
		// _ = "end of CoverTab[4575]"
	} else {
//line /usr/local/go/src/net/conf.go:285
		_go_fuzz_dep_.CoverTab[4576]++
//line /usr/local/go/src/net/conf.go:285
		// _ = "end of CoverTab[4576]"
//line /usr/local/go/src/net/conf.go:285
	}
//line /usr/local/go/src/net/conf.go:285
	// _ = "end of CoverTab[4478]"
//line /usr/local/go/src/net/conf.go:285
	_go_fuzz_dep_.CoverTab[4479]++

//line /usr/local/go/src/net/conf.go:289
	switch {
	case filesSource && func() bool {
//line /usr/local/go/src/net/conf.go:290
		_go_fuzz_dep_.CoverTab[4581]++
//line /usr/local/go/src/net/conf.go:290
		return dnsSource
//line /usr/local/go/src/net/conf.go:290
		// _ = "end of CoverTab[4581]"
//line /usr/local/go/src/net/conf.go:290
	}():
//line /usr/local/go/src/net/conf.go:290
		_go_fuzz_dep_.CoverTab[4577]++
							if first == "files" {
//line /usr/local/go/src/net/conf.go:291
			_go_fuzz_dep_.CoverTab[4582]++
								return hostLookupFilesDNS, conf
//line /usr/local/go/src/net/conf.go:292
			// _ = "end of CoverTab[4582]"
		} else {
//line /usr/local/go/src/net/conf.go:293
			_go_fuzz_dep_.CoverTab[4583]++
								return hostLookupDNSFiles, conf
//line /usr/local/go/src/net/conf.go:294
			// _ = "end of CoverTab[4583]"
		}
//line /usr/local/go/src/net/conf.go:295
		// _ = "end of CoverTab[4577]"
	case filesSource:
//line /usr/local/go/src/net/conf.go:296
		_go_fuzz_dep_.CoverTab[4578]++
							return hostLookupFiles, conf
//line /usr/local/go/src/net/conf.go:297
		// _ = "end of CoverTab[4578]"
	case dnsSource:
//line /usr/local/go/src/net/conf.go:298
		_go_fuzz_dep_.CoverTab[4579]++
							return hostLookupDNS, conf
//line /usr/local/go/src/net/conf.go:299
		// _ = "end of CoverTab[4579]"
//line /usr/local/go/src/net/conf.go:299
	default:
//line /usr/local/go/src/net/conf.go:299
		_go_fuzz_dep_.CoverTab[4580]++
//line /usr/local/go/src/net/conf.go:299
		// _ = "end of CoverTab[4580]"
	}
//line /usr/local/go/src/net/conf.go:300
	// _ = "end of CoverTab[4479]"
//line /usr/local/go/src/net/conf.go:300
	_go_fuzz_dep_.CoverTab[4480]++

//line /usr/local/go/src/net/conf.go:303
	return fallbackOrder, conf
//line /usr/local/go/src/net/conf.go:303
	// _ = "end of CoverTab[4480]"
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
	_go_fuzz_dep_.CoverTab[4584]++
						goDebug := netdns.Value()
						parsePart := func(s string) {
//line /usr/local/go/src/net/conf.go:322
		_go_fuzz_dep_.CoverTab[4587]++
							if s == "" {
//line /usr/local/go/src/net/conf.go:323
			_go_fuzz_dep_.CoverTab[4589]++
								return
//line /usr/local/go/src/net/conf.go:324
			// _ = "end of CoverTab[4589]"
		} else {
//line /usr/local/go/src/net/conf.go:325
			_go_fuzz_dep_.CoverTab[4590]++
//line /usr/local/go/src/net/conf.go:325
			// _ = "end of CoverTab[4590]"
//line /usr/local/go/src/net/conf.go:325
		}
//line /usr/local/go/src/net/conf.go:325
		// _ = "end of CoverTab[4587]"
//line /usr/local/go/src/net/conf.go:325
		_go_fuzz_dep_.CoverTab[4588]++
							if '0' <= s[0] && func() bool {
//line /usr/local/go/src/net/conf.go:326
			_go_fuzz_dep_.CoverTab[4591]++
//line /usr/local/go/src/net/conf.go:326
			return s[0] <= '9'
//line /usr/local/go/src/net/conf.go:326
			// _ = "end of CoverTab[4591]"
//line /usr/local/go/src/net/conf.go:326
		}() {
//line /usr/local/go/src/net/conf.go:326
			_go_fuzz_dep_.CoverTab[4592]++
								debugLevel, _, _ = dtoi(s)
//line /usr/local/go/src/net/conf.go:327
			// _ = "end of CoverTab[4592]"
		} else {
//line /usr/local/go/src/net/conf.go:328
			_go_fuzz_dep_.CoverTab[4593]++
								dnsMode = s
//line /usr/local/go/src/net/conf.go:329
			// _ = "end of CoverTab[4593]"
		}
//line /usr/local/go/src/net/conf.go:330
		// _ = "end of CoverTab[4588]"
	}
//line /usr/local/go/src/net/conf.go:331
	// _ = "end of CoverTab[4584]"
//line /usr/local/go/src/net/conf.go:331
	_go_fuzz_dep_.CoverTab[4585]++
						if i := bytealg.IndexByteString(goDebug, '+'); i != -1 {
//line /usr/local/go/src/net/conf.go:332
		_go_fuzz_dep_.CoverTab[4594]++
							parsePart(goDebug[:i])
							parsePart(goDebug[i+1:])
							return
//line /usr/local/go/src/net/conf.go:335
		// _ = "end of CoverTab[4594]"
	} else {
//line /usr/local/go/src/net/conf.go:336
		_go_fuzz_dep_.CoverTab[4595]++
//line /usr/local/go/src/net/conf.go:336
		// _ = "end of CoverTab[4595]"
//line /usr/local/go/src/net/conf.go:336
	}
//line /usr/local/go/src/net/conf.go:336
	// _ = "end of CoverTab[4585]"
//line /usr/local/go/src/net/conf.go:336
	_go_fuzz_dep_.CoverTab[4586]++
						parsePart(goDebug)
						return
//line /usr/local/go/src/net/conf.go:338
	// _ = "end of CoverTab[4586]"
}

// isLocalhost reports whether h should be considered a "localhost"
//line /usr/local/go/src/net/conf.go:341
// name for the myhostname NSS module.
//line /usr/local/go/src/net/conf.go:343
func isLocalhost(h string) bool {
//line /usr/local/go/src/net/conf.go:343
	_go_fuzz_dep_.CoverTab[4596]++
						return stringsEqualFold(h, "localhost") || func() bool {
//line /usr/local/go/src/net/conf.go:344
		_go_fuzz_dep_.CoverTab[4597]++
//line /usr/local/go/src/net/conf.go:344
		return stringsEqualFold(h, "localhost.localdomain")
//line /usr/local/go/src/net/conf.go:344
		// _ = "end of CoverTab[4597]"
//line /usr/local/go/src/net/conf.go:344
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:344
		_go_fuzz_dep_.CoverTab[4598]++
//line /usr/local/go/src/net/conf.go:344
		return stringsHasSuffixFold(h, ".localhost")
//line /usr/local/go/src/net/conf.go:344
		// _ = "end of CoverTab[4598]"
//line /usr/local/go/src/net/conf.go:344
	}() || func() bool {
//line /usr/local/go/src/net/conf.go:344
		_go_fuzz_dep_.CoverTab[4599]++
//line /usr/local/go/src/net/conf.go:344
		return stringsHasSuffixFold(h, ".localhost.localdomain")
//line /usr/local/go/src/net/conf.go:344
		// _ = "end of CoverTab[4599]"
//line /usr/local/go/src/net/conf.go:344
	}()
//line /usr/local/go/src/net/conf.go:344
	// _ = "end of CoverTab[4596]"
}

// isGateway reports whether h should be considered a "gateway"
//line /usr/local/go/src/net/conf.go:347
// name for the myhostname NSS module.
//line /usr/local/go/src/net/conf.go:349
func isGateway(h string) bool {
//line /usr/local/go/src/net/conf.go:349
	_go_fuzz_dep_.CoverTab[4600]++
						return stringsEqualFold(h, "_gateway")
//line /usr/local/go/src/net/conf.go:350
	// _ = "end of CoverTab[4600]"
}

// isOutbound reports whether h should be considered a "outbound"
//line /usr/local/go/src/net/conf.go:353
// name for the myhostname NSS module.
//line /usr/local/go/src/net/conf.go:355
func isOutbound(h string) bool {
//line /usr/local/go/src/net/conf.go:355
	_go_fuzz_dep_.CoverTab[4601]++
						return stringsEqualFold(h, "_outbound")
//line /usr/local/go/src/net/conf.go:356
	// _ = "end of CoverTab[4601]"
}

//line /usr/local/go/src/net/conf.go:357
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/conf.go:357
var _ = _go_fuzz_dep_.CoverTab
