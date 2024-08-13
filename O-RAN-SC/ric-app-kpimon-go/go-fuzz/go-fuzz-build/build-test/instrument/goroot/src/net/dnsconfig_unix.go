// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js && !windows

// Read system DNS config from /etc/resolv.conf

//line /snap/go/10455/src/net/dnsconfig_unix.go:9
package net

//line /snap/go/10455/src/net/dnsconfig_unix.go:9
import (
//line /snap/go/10455/src/net/dnsconfig_unix.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/dnsconfig_unix.go:9
)
//line /snap/go/10455/src/net/dnsconfig_unix.go:9
import (
//line /snap/go/10455/src/net/dnsconfig_unix.go:9
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/dnsconfig_unix.go:9
)

import (
	"internal/bytealg"
	"net/netip"
	"time"
)

// See resolv.conf(5) on a Linux machine.
func dnsReadConfig(filename string) *dnsConfig {
//line /snap/go/10455/src/net/dnsconfig_unix.go:18
	_go_fuzz_dep_.CoverTab[5724]++
							conf := &dnsConfig{
		ndots:		1,
		timeout:	5 * time.Second,
		attempts:	2,
	}
	file, err := open(filename)
	if err != nil {
//line /snap/go/10455/src/net/dnsconfig_unix.go:25
		_go_fuzz_dep_.CoverTab[528187]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:25
		_go_fuzz_dep_.CoverTab[5730]++
								conf.servers = defaultNS
								conf.search = dnsDefaultSearch()
								conf.err = err
								return conf
//line /snap/go/10455/src/net/dnsconfig_unix.go:29
		// _ = "end of CoverTab[5730]"
	} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:30
		_go_fuzz_dep_.CoverTab[528188]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:30
		_go_fuzz_dep_.CoverTab[5731]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:30
		// _ = "end of CoverTab[5731]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:30
	}
//line /snap/go/10455/src/net/dnsconfig_unix.go:30
	// _ = "end of CoverTab[5724]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:30
	_go_fuzz_dep_.CoverTab[5725]++
							defer file.close()
							if fi, err := file.file.Stat(); err == nil {
//line /snap/go/10455/src/net/dnsconfig_unix.go:32
		_go_fuzz_dep_.CoverTab[528189]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:32
		_go_fuzz_dep_.CoverTab[5732]++
								conf.mtime = fi.ModTime()
//line /snap/go/10455/src/net/dnsconfig_unix.go:33
		// _ = "end of CoverTab[5732]"
	} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:34
		_go_fuzz_dep_.CoverTab[528190]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:34
		_go_fuzz_dep_.CoverTab[5733]++
								conf.servers = defaultNS
								conf.search = dnsDefaultSearch()
								conf.err = err
								return conf
//line /snap/go/10455/src/net/dnsconfig_unix.go:38
		// _ = "end of CoverTab[5733]"
	}
//line /snap/go/10455/src/net/dnsconfig_unix.go:39
	// _ = "end of CoverTab[5725]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:39
	_go_fuzz_dep_.CoverTab[5726]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:39
	_go_fuzz_dep_.CoverTab[786673] = 0
							for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /snap/go/10455/src/net/dnsconfig_unix.go:40
		if _go_fuzz_dep_.CoverTab[786673] == 0 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:40
			_go_fuzz_dep_.CoverTab[528237]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:40
		} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:40
			_go_fuzz_dep_.CoverTab[528238]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:40
		}
//line /snap/go/10455/src/net/dnsconfig_unix.go:40
		_go_fuzz_dep_.CoverTab[786673] = 1
//line /snap/go/10455/src/net/dnsconfig_unix.go:40
		_go_fuzz_dep_.CoverTab[5734]++
								if len(line) > 0 && func() bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
			_go_fuzz_dep_.CoverTab[5737]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
			return (line[0] == ';' || func() bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
				_go_fuzz_dep_.CoverTab[5738]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
				return line[0] == '#'
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
				// _ = "end of CoverTab[5738]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
			}())
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
			// _ = "end of CoverTab[5737]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
		}() {
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
			_go_fuzz_dep_.CoverTab[528191]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:41
			_go_fuzz_dep_.CoverTab[5739]++

									continue
//line /snap/go/10455/src/net/dnsconfig_unix.go:43
			// _ = "end of CoverTab[5739]"
		} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:44
			_go_fuzz_dep_.CoverTab[528192]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:44
			_go_fuzz_dep_.CoverTab[5740]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:44
			// _ = "end of CoverTab[5740]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:44
		}
//line /snap/go/10455/src/net/dnsconfig_unix.go:44
		// _ = "end of CoverTab[5734]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:44
		_go_fuzz_dep_.CoverTab[5735]++
								f := getFields(line)
								if len(f) < 1 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:46
			_go_fuzz_dep_.CoverTab[528193]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:46
			_go_fuzz_dep_.CoverTab[5741]++
									continue
//line /snap/go/10455/src/net/dnsconfig_unix.go:47
			// _ = "end of CoverTab[5741]"
		} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:48
			_go_fuzz_dep_.CoverTab[528194]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:48
			_go_fuzz_dep_.CoverTab[5742]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:48
			// _ = "end of CoverTab[5742]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:48
		}
//line /snap/go/10455/src/net/dnsconfig_unix.go:48
		// _ = "end of CoverTab[5735]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:48
		_go_fuzz_dep_.CoverTab[5736]++
								switch f[0] {
		case "nameserver":
//line /snap/go/10455/src/net/dnsconfig_unix.go:50
			_go_fuzz_dep_.CoverTab[528195]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:50
			_go_fuzz_dep_.CoverTab[5743]++
									if len(f) > 1 && func() bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:51
				_go_fuzz_dep_.CoverTab[5749]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:51
				return len(conf.servers) < 3
//line /snap/go/10455/src/net/dnsconfig_unix.go:51
				// _ = "end of CoverTab[5749]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:51
			}() {
//line /snap/go/10455/src/net/dnsconfig_unix.go:51
				_go_fuzz_dep_.CoverTab[528201]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:51
				_go_fuzz_dep_.CoverTab[5750]++

//line /snap/go/10455/src/net/dnsconfig_unix.go:55
				if _, err := netip.ParseAddr(f[1]); err == nil {
//line /snap/go/10455/src/net/dnsconfig_unix.go:55
					_go_fuzz_dep_.CoverTab[528203]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:55
					_go_fuzz_dep_.CoverTab[5751]++
											conf.servers = append(conf.servers, JoinHostPort(f[1], "53"))
//line /snap/go/10455/src/net/dnsconfig_unix.go:56
					// _ = "end of CoverTab[5751]"
				} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:57
					_go_fuzz_dep_.CoverTab[528204]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:57
					_go_fuzz_dep_.CoverTab[5752]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:57
					// _ = "end of CoverTab[5752]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:57
				}
//line /snap/go/10455/src/net/dnsconfig_unix.go:57
				// _ = "end of CoverTab[5750]"
			} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:58
				_go_fuzz_dep_.CoverTab[528202]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:58
				_go_fuzz_dep_.CoverTab[5753]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:58
				// _ = "end of CoverTab[5753]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:58
			}
//line /snap/go/10455/src/net/dnsconfig_unix.go:58
			// _ = "end of CoverTab[5743]"

		case "domain":
//line /snap/go/10455/src/net/dnsconfig_unix.go:60
			_go_fuzz_dep_.CoverTab[528196]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:60
			_go_fuzz_dep_.CoverTab[5744]++
									if len(f) > 1 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:61
				_go_fuzz_dep_.CoverTab[528205]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:61
				_go_fuzz_dep_.CoverTab[5754]++
										conf.search = []string{ensureRooted(f[1])}
//line /snap/go/10455/src/net/dnsconfig_unix.go:62
				// _ = "end of CoverTab[5754]"
			} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:63
				_go_fuzz_dep_.CoverTab[528206]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:63
				_go_fuzz_dep_.CoverTab[5755]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:63
				// _ = "end of CoverTab[5755]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:63
			}
//line /snap/go/10455/src/net/dnsconfig_unix.go:63
			// _ = "end of CoverTab[5744]"

		case "search":
//line /snap/go/10455/src/net/dnsconfig_unix.go:65
			_go_fuzz_dep_.CoverTab[528197]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:65
			_go_fuzz_dep_.CoverTab[5745]++
									conf.search = make([]string, 0, len(f)-1)
									for i := 1; i < len(f); i++ {
//line /snap/go/10455/src/net/dnsconfig_unix.go:67
				_go_fuzz_dep_.CoverTab[5756]++
										name := ensureRooted(f[i])
										if name == "." {
//line /snap/go/10455/src/net/dnsconfig_unix.go:69
					_go_fuzz_dep_.CoverTab[528207]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:69
					_go_fuzz_dep_.CoverTab[5758]++
											continue
//line /snap/go/10455/src/net/dnsconfig_unix.go:70
					// _ = "end of CoverTab[5758]"
				} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:71
					_go_fuzz_dep_.CoverTab[528208]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:71
					_go_fuzz_dep_.CoverTab[5759]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:71
					// _ = "end of CoverTab[5759]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:71
				}
//line /snap/go/10455/src/net/dnsconfig_unix.go:71
				// _ = "end of CoverTab[5756]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:71
				_go_fuzz_dep_.CoverTab[5757]++
										conf.search = append(conf.search, name)
//line /snap/go/10455/src/net/dnsconfig_unix.go:72
				// _ = "end of CoverTab[5757]"
			}
//line /snap/go/10455/src/net/dnsconfig_unix.go:73
			// _ = "end of CoverTab[5745]"

		case "options":
//line /snap/go/10455/src/net/dnsconfig_unix.go:75
			_go_fuzz_dep_.CoverTab[528198]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:75
			_go_fuzz_dep_.CoverTab[5746]++
									for _, s := range f[1:] {
//line /snap/go/10455/src/net/dnsconfig_unix.go:76
				_go_fuzz_dep_.CoverTab[5760]++
										switch {
				case hasPrefix(s, "ndots:"):
//line /snap/go/10455/src/net/dnsconfig_unix.go:78
					_go_fuzz_dep_.CoverTab[528209]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:78
					_go_fuzz_dep_.CoverTab[5761]++
											n, _, _ := dtoi(s[6:])
											if n < 0 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:80
						_go_fuzz_dep_.CoverTab[528219]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:80
						_go_fuzz_dep_.CoverTab[5774]++
												n = 0
//line /snap/go/10455/src/net/dnsconfig_unix.go:81
						// _ = "end of CoverTab[5774]"
					} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:82
						_go_fuzz_dep_.CoverTab[528220]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:82
						_go_fuzz_dep_.CoverTab[5775]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:82
						if n > 15 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:82
							_go_fuzz_dep_.CoverTab[528221]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:82
							_go_fuzz_dep_.CoverTab[5776]++
													n = 15
//line /snap/go/10455/src/net/dnsconfig_unix.go:83
							// _ = "end of CoverTab[5776]"
						} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:84
							_go_fuzz_dep_.CoverTab[528222]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:84
							_go_fuzz_dep_.CoverTab[5777]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:84
							// _ = "end of CoverTab[5777]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:84
						}
//line /snap/go/10455/src/net/dnsconfig_unix.go:84
						// _ = "end of CoverTab[5775]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:84
					}
//line /snap/go/10455/src/net/dnsconfig_unix.go:84
					// _ = "end of CoverTab[5761]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:84
					_go_fuzz_dep_.CoverTab[5762]++
											conf.ndots = n
//line /snap/go/10455/src/net/dnsconfig_unix.go:85
					// _ = "end of CoverTab[5762]"
				case hasPrefix(s, "timeout:"):
//line /snap/go/10455/src/net/dnsconfig_unix.go:86
					_go_fuzz_dep_.CoverTab[528210]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:86
					_go_fuzz_dep_.CoverTab[5763]++
											n, _, _ := dtoi(s[8:])
											if n < 1 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:88
						_go_fuzz_dep_.CoverTab[528223]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:88
						_go_fuzz_dep_.CoverTab[5778]++
												n = 1
//line /snap/go/10455/src/net/dnsconfig_unix.go:89
						// _ = "end of CoverTab[5778]"
					} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:90
						_go_fuzz_dep_.CoverTab[528224]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:90
						_go_fuzz_dep_.CoverTab[5779]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:90
						// _ = "end of CoverTab[5779]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:90
					}
//line /snap/go/10455/src/net/dnsconfig_unix.go:90
					// _ = "end of CoverTab[5763]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:90
					_go_fuzz_dep_.CoverTab[5764]++
											conf.timeout = time.Duration(n) * time.Second
//line /snap/go/10455/src/net/dnsconfig_unix.go:91
					// _ = "end of CoverTab[5764]"
				case hasPrefix(s, "attempts:"):
//line /snap/go/10455/src/net/dnsconfig_unix.go:92
					_go_fuzz_dep_.CoverTab[528211]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:92
					_go_fuzz_dep_.CoverTab[5765]++
											n, _, _ := dtoi(s[9:])
											if n < 1 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:94
						_go_fuzz_dep_.CoverTab[528225]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:94
						_go_fuzz_dep_.CoverTab[5780]++
												n = 1
//line /snap/go/10455/src/net/dnsconfig_unix.go:95
						// _ = "end of CoverTab[5780]"
					} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:96
						_go_fuzz_dep_.CoverTab[528226]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:96
						_go_fuzz_dep_.CoverTab[5781]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:96
						// _ = "end of CoverTab[5781]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:96
					}
//line /snap/go/10455/src/net/dnsconfig_unix.go:96
					// _ = "end of CoverTab[5765]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:96
					_go_fuzz_dep_.CoverTab[5766]++
											conf.attempts = n
//line /snap/go/10455/src/net/dnsconfig_unix.go:97
					// _ = "end of CoverTab[5766]"
				case s == "rotate":
//line /snap/go/10455/src/net/dnsconfig_unix.go:98
					_go_fuzz_dep_.CoverTab[528212]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:98
					_go_fuzz_dep_.CoverTab[5767]++
											conf.rotate = true
//line /snap/go/10455/src/net/dnsconfig_unix.go:99
					// _ = "end of CoverTab[5767]"
				case s == "single-request" || func() bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:100
					_go_fuzz_dep_.CoverTab[5782]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:100
					return s == "single-request-reopen"
//line /snap/go/10455/src/net/dnsconfig_unix.go:100
					// _ = "end of CoverTab[5782]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:100
				}():
//line /snap/go/10455/src/net/dnsconfig_unix.go:100
					_go_fuzz_dep_.CoverTab[528213]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:100
					_go_fuzz_dep_.CoverTab[5768]++

//line /snap/go/10455/src/net/dnsconfig_unix.go:106
					conf.singleRequest = true
//line /snap/go/10455/src/net/dnsconfig_unix.go:106
					// _ = "end of CoverTab[5768]"
				case s == "use-vc" || func() bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
					_go_fuzz_dep_.CoverTab[5783]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
					return s == "usevc"
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
					// _ = "end of CoverTab[5783]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
				}() || func() bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
					_go_fuzz_dep_.CoverTab[5784]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
					return s == "tcp"
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
					// _ = "end of CoverTab[5784]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
				}():
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
					_go_fuzz_dep_.CoverTab[528214]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:107
					_go_fuzz_dep_.CoverTab[5769]++

//line /snap/go/10455/src/net/dnsconfig_unix.go:114
					conf.useTCP = true
//line /snap/go/10455/src/net/dnsconfig_unix.go:114
					// _ = "end of CoverTab[5769]"
				case s == "trust-ad":
//line /snap/go/10455/src/net/dnsconfig_unix.go:115
					_go_fuzz_dep_.CoverTab[528215]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:115
					_go_fuzz_dep_.CoverTab[5770]++
											conf.trustAD = true
//line /snap/go/10455/src/net/dnsconfig_unix.go:116
					// _ = "end of CoverTab[5770]"
				case s == "edns0":
//line /snap/go/10455/src/net/dnsconfig_unix.go:117
					_go_fuzz_dep_.CoverTab[528216]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:117
					_go_fuzz_dep_.CoverTab[5771]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:117
					// _ = "end of CoverTab[5771]"

//line /snap/go/10455/src/net/dnsconfig_unix.go:120
				case s == "no-reload":
//line /snap/go/10455/src/net/dnsconfig_unix.go:120
					_go_fuzz_dep_.CoverTab[528217]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:120
					_go_fuzz_dep_.CoverTab[5772]++
											conf.noReload = true
//line /snap/go/10455/src/net/dnsconfig_unix.go:121
					// _ = "end of CoverTab[5772]"
				default:
//line /snap/go/10455/src/net/dnsconfig_unix.go:122
					_go_fuzz_dep_.CoverTab[528218]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:122
					_go_fuzz_dep_.CoverTab[5773]++
											conf.unknownOpt = true
//line /snap/go/10455/src/net/dnsconfig_unix.go:123
					// _ = "end of CoverTab[5773]"
				}
//line /snap/go/10455/src/net/dnsconfig_unix.go:124
				// _ = "end of CoverTab[5760]"
			}
//line /snap/go/10455/src/net/dnsconfig_unix.go:125
			// _ = "end of CoverTab[5746]"

		case "lookup":
//line /snap/go/10455/src/net/dnsconfig_unix.go:127
			_go_fuzz_dep_.CoverTab[528199]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:127
			_go_fuzz_dep_.CoverTab[5747]++

//line /snap/go/10455/src/net/dnsconfig_unix.go:131
			conf.lookup = f[1:]
//line /snap/go/10455/src/net/dnsconfig_unix.go:131
			// _ = "end of CoverTab[5747]"

		default:
//line /snap/go/10455/src/net/dnsconfig_unix.go:133
			_go_fuzz_dep_.CoverTab[528200]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:133
			_go_fuzz_dep_.CoverTab[5748]++
									conf.unknownOpt = true
//line /snap/go/10455/src/net/dnsconfig_unix.go:134
			// _ = "end of CoverTab[5748]"
		}
//line /snap/go/10455/src/net/dnsconfig_unix.go:135
		// _ = "end of CoverTab[5736]"
	}
//line /snap/go/10455/src/net/dnsconfig_unix.go:136
	if _go_fuzz_dep_.CoverTab[786673] == 0 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:136
		_go_fuzz_dep_.CoverTab[528239]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:136
	} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:136
		_go_fuzz_dep_.CoverTab[528240]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:136
	}
//line /snap/go/10455/src/net/dnsconfig_unix.go:136
	// _ = "end of CoverTab[5726]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:136
	_go_fuzz_dep_.CoverTab[5727]++
							if len(conf.servers) == 0 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:137
		_go_fuzz_dep_.CoverTab[528227]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:137
		_go_fuzz_dep_.CoverTab[5785]++
								conf.servers = defaultNS
//line /snap/go/10455/src/net/dnsconfig_unix.go:138
		// _ = "end of CoverTab[5785]"
	} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:139
		_go_fuzz_dep_.CoverTab[528228]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:139
		_go_fuzz_dep_.CoverTab[5786]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:139
		// _ = "end of CoverTab[5786]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:139
	}
//line /snap/go/10455/src/net/dnsconfig_unix.go:139
	// _ = "end of CoverTab[5727]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:139
	_go_fuzz_dep_.CoverTab[5728]++
							if len(conf.search) == 0 {
//line /snap/go/10455/src/net/dnsconfig_unix.go:140
		_go_fuzz_dep_.CoverTab[528229]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:140
		_go_fuzz_dep_.CoverTab[5787]++
								conf.search = dnsDefaultSearch()
//line /snap/go/10455/src/net/dnsconfig_unix.go:141
		// _ = "end of CoverTab[5787]"
	} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:142
		_go_fuzz_dep_.CoverTab[528230]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:142
		_go_fuzz_dep_.CoverTab[5788]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:142
		// _ = "end of CoverTab[5788]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:142
	}
//line /snap/go/10455/src/net/dnsconfig_unix.go:142
	// _ = "end of CoverTab[5728]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:142
	_go_fuzz_dep_.CoverTab[5729]++
							return conf
//line /snap/go/10455/src/net/dnsconfig_unix.go:143
	// _ = "end of CoverTab[5729]"
}

func dnsDefaultSearch() []string {
//line /snap/go/10455/src/net/dnsconfig_unix.go:146
	_go_fuzz_dep_.CoverTab[5789]++
							hn, err := getHostname()
							if err != nil {
//line /snap/go/10455/src/net/dnsconfig_unix.go:148
		_go_fuzz_dep_.CoverTab[528231]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:148
		_go_fuzz_dep_.CoverTab[5792]++

								return nil
//line /snap/go/10455/src/net/dnsconfig_unix.go:150
		// _ = "end of CoverTab[5792]"
	} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:151
		_go_fuzz_dep_.CoverTab[528232]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:151
		_go_fuzz_dep_.CoverTab[5793]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:151
		// _ = "end of CoverTab[5793]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:151
	}
//line /snap/go/10455/src/net/dnsconfig_unix.go:151
	// _ = "end of CoverTab[5789]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:151
	_go_fuzz_dep_.CoverTab[5790]++
							if i := bytealg.IndexByteString(hn, '.'); i >= 0 && func() bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:152
		_go_fuzz_dep_.CoverTab[5794]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:152
		return i < len(hn)-1
//line /snap/go/10455/src/net/dnsconfig_unix.go:152
		// _ = "end of CoverTab[5794]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:152
	}() {
//line /snap/go/10455/src/net/dnsconfig_unix.go:152
		_go_fuzz_dep_.CoverTab[528233]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:152
		_go_fuzz_dep_.CoverTab[5795]++
								return []string{ensureRooted(hn[i+1:])}
//line /snap/go/10455/src/net/dnsconfig_unix.go:153
		// _ = "end of CoverTab[5795]"
	} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:154
		_go_fuzz_dep_.CoverTab[528234]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:154
		_go_fuzz_dep_.CoverTab[5796]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:154
		// _ = "end of CoverTab[5796]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:154
	}
//line /snap/go/10455/src/net/dnsconfig_unix.go:154
	// _ = "end of CoverTab[5790]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:154
	_go_fuzz_dep_.CoverTab[5791]++
							return nil
//line /snap/go/10455/src/net/dnsconfig_unix.go:155
	// _ = "end of CoverTab[5791]"
}

func hasPrefix(s, prefix string) bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:158
	_go_fuzz_dep_.CoverTab[5797]++
							return len(s) >= len(prefix) && func() bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:159
		_go_fuzz_dep_.CoverTab[5798]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:159
		return s[:len(prefix)] == prefix
//line /snap/go/10455/src/net/dnsconfig_unix.go:159
		// _ = "end of CoverTab[5798]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:159
	}()
//line /snap/go/10455/src/net/dnsconfig_unix.go:159
	// _ = "end of CoverTab[5797]"
}

func ensureRooted(s string) string {
//line /snap/go/10455/src/net/dnsconfig_unix.go:162
	_go_fuzz_dep_.CoverTab[5799]++
							if len(s) > 0 && func() bool {
//line /snap/go/10455/src/net/dnsconfig_unix.go:163
		_go_fuzz_dep_.CoverTab[5801]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:163
		return s[len(s)-1] == '.'
//line /snap/go/10455/src/net/dnsconfig_unix.go:163
		// _ = "end of CoverTab[5801]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:163
	}() {
//line /snap/go/10455/src/net/dnsconfig_unix.go:163
		_go_fuzz_dep_.CoverTab[528235]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:163
		_go_fuzz_dep_.CoverTab[5802]++
								return s
//line /snap/go/10455/src/net/dnsconfig_unix.go:164
		// _ = "end of CoverTab[5802]"
	} else {
//line /snap/go/10455/src/net/dnsconfig_unix.go:165
		_go_fuzz_dep_.CoverTab[528236]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:165
		_go_fuzz_dep_.CoverTab[5803]++
//line /snap/go/10455/src/net/dnsconfig_unix.go:165
		// _ = "end of CoverTab[5803]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:165
	}
//line /snap/go/10455/src/net/dnsconfig_unix.go:165
	// _ = "end of CoverTab[5799]"
//line /snap/go/10455/src/net/dnsconfig_unix.go:165
	_go_fuzz_dep_.CoverTab[5800]++
							return s + "."
//line /snap/go/10455/src/net/dnsconfig_unix.go:166
	// _ = "end of CoverTab[5800]"
}

//line /snap/go/10455/src/net/dnsconfig_unix.go:167
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/dnsconfig_unix.go:167
var _ = _go_fuzz_dep_.CoverTab
