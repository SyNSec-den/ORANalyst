// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js && !windows

// Read system DNS config from /etc/resolv.conf

//line /usr/local/go/src/net/dnsconfig_unix.go:9
package net

//line /usr/local/go/src/net/dnsconfig_unix.go:9
import (
//line /usr/local/go/src/net/dnsconfig_unix.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/dnsconfig_unix.go:9
)
//line /usr/local/go/src/net/dnsconfig_unix.go:9
import (
//line /usr/local/go/src/net/dnsconfig_unix.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/dnsconfig_unix.go:9
)

import (
	"internal/bytealg"
	"time"
)

// See resolv.conf(5) on a Linux machine.
func dnsReadConfig(filename string) *dnsConfig {
//line /usr/local/go/src/net/dnsconfig_unix.go:17
	_go_fuzz_dep_.CoverTab[13736]++
							conf := &dnsConfig{
		ndots:		1,
		timeout:	5 * time.Second,
		attempts:	2,
	}
	file, err := open(filename)
	if err != nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:24
		_go_fuzz_dep_.CoverTab[13742]++
								conf.servers = defaultNS
								conf.search = dnsDefaultSearch()
								conf.err = err
								return conf
//line /usr/local/go/src/net/dnsconfig_unix.go:28
		// _ = "end of CoverTab[13742]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:29
		_go_fuzz_dep_.CoverTab[13743]++
//line /usr/local/go/src/net/dnsconfig_unix.go:29
		// _ = "end of CoverTab[13743]"
//line /usr/local/go/src/net/dnsconfig_unix.go:29
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:29
	// _ = "end of CoverTab[13736]"
//line /usr/local/go/src/net/dnsconfig_unix.go:29
	_go_fuzz_dep_.CoverTab[13737]++
							defer file.close()
							if fi, err := file.file.Stat(); err == nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:31
		_go_fuzz_dep_.CoverTab[13744]++
								conf.mtime = fi.ModTime()
//line /usr/local/go/src/net/dnsconfig_unix.go:32
		// _ = "end of CoverTab[13744]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:33
		_go_fuzz_dep_.CoverTab[13745]++
								conf.servers = defaultNS
								conf.search = dnsDefaultSearch()
								conf.err = err
								return conf
//line /usr/local/go/src/net/dnsconfig_unix.go:37
		// _ = "end of CoverTab[13745]"
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:38
	// _ = "end of CoverTab[13737]"
//line /usr/local/go/src/net/dnsconfig_unix.go:38
	_go_fuzz_dep_.CoverTab[13738]++
							for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /usr/local/go/src/net/dnsconfig_unix.go:39
		_go_fuzz_dep_.CoverTab[13746]++
								if len(line) > 0 && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			_go_fuzz_dep_.CoverTab[13749]++
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			return (line[0] == ';' || func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:40
				_go_fuzz_dep_.CoverTab[13750]++
//line /usr/local/go/src/net/dnsconfig_unix.go:40
				return line[0] == '#'
//line /usr/local/go/src/net/dnsconfig_unix.go:40
				// _ = "end of CoverTab[13750]"
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			}())
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			// _ = "end of CoverTab[13749]"
//line /usr/local/go/src/net/dnsconfig_unix.go:40
		}() {
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			_go_fuzz_dep_.CoverTab[13751]++

									continue
//line /usr/local/go/src/net/dnsconfig_unix.go:42
			// _ = "end of CoverTab[13751]"
		} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:43
			_go_fuzz_dep_.CoverTab[13752]++
//line /usr/local/go/src/net/dnsconfig_unix.go:43
			// _ = "end of CoverTab[13752]"
//line /usr/local/go/src/net/dnsconfig_unix.go:43
		}
//line /usr/local/go/src/net/dnsconfig_unix.go:43
		// _ = "end of CoverTab[13746]"
//line /usr/local/go/src/net/dnsconfig_unix.go:43
		_go_fuzz_dep_.CoverTab[13747]++
								f := getFields(line)
								if len(f) < 1 {
//line /usr/local/go/src/net/dnsconfig_unix.go:45
			_go_fuzz_dep_.CoverTab[13753]++
									continue
//line /usr/local/go/src/net/dnsconfig_unix.go:46
			// _ = "end of CoverTab[13753]"
		} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:47
			_go_fuzz_dep_.CoverTab[13754]++
//line /usr/local/go/src/net/dnsconfig_unix.go:47
			// _ = "end of CoverTab[13754]"
//line /usr/local/go/src/net/dnsconfig_unix.go:47
		}
//line /usr/local/go/src/net/dnsconfig_unix.go:47
		// _ = "end of CoverTab[13747]"
//line /usr/local/go/src/net/dnsconfig_unix.go:47
		_go_fuzz_dep_.CoverTab[13748]++
								switch f[0] {
		case "nameserver":
//line /usr/local/go/src/net/dnsconfig_unix.go:49
			_go_fuzz_dep_.CoverTab[13755]++
									if len(f) > 1 && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:50
				_go_fuzz_dep_.CoverTab[13761]++
//line /usr/local/go/src/net/dnsconfig_unix.go:50
				return len(conf.servers) < 3
//line /usr/local/go/src/net/dnsconfig_unix.go:50
				// _ = "end of CoverTab[13761]"
//line /usr/local/go/src/net/dnsconfig_unix.go:50
			}() {
//line /usr/local/go/src/net/dnsconfig_unix.go:50
				_go_fuzz_dep_.CoverTab[13762]++

//line /usr/local/go/src/net/dnsconfig_unix.go:54
				if parseIPv4(f[1]) != nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:54
					_go_fuzz_dep_.CoverTab[13763]++
											conf.servers = append(conf.servers, JoinHostPort(f[1], "53"))
//line /usr/local/go/src/net/dnsconfig_unix.go:55
					// _ = "end of CoverTab[13763]"
				} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:56
					_go_fuzz_dep_.CoverTab[13764]++
//line /usr/local/go/src/net/dnsconfig_unix.go:56
					if ip, _ := parseIPv6Zone(f[1]); ip != nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:56
						_go_fuzz_dep_.CoverTab[13765]++
												conf.servers = append(conf.servers, JoinHostPort(f[1], "53"))
//line /usr/local/go/src/net/dnsconfig_unix.go:57
						// _ = "end of CoverTab[13765]"
					} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:58
						_go_fuzz_dep_.CoverTab[13766]++
//line /usr/local/go/src/net/dnsconfig_unix.go:58
						// _ = "end of CoverTab[13766]"
//line /usr/local/go/src/net/dnsconfig_unix.go:58
					}
//line /usr/local/go/src/net/dnsconfig_unix.go:58
					// _ = "end of CoverTab[13764]"
//line /usr/local/go/src/net/dnsconfig_unix.go:58
				}
//line /usr/local/go/src/net/dnsconfig_unix.go:58
				// _ = "end of CoverTab[13762]"
			} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:59
				_go_fuzz_dep_.CoverTab[13767]++
//line /usr/local/go/src/net/dnsconfig_unix.go:59
				// _ = "end of CoverTab[13767]"
//line /usr/local/go/src/net/dnsconfig_unix.go:59
			}
//line /usr/local/go/src/net/dnsconfig_unix.go:59
			// _ = "end of CoverTab[13755]"

		case "domain":
//line /usr/local/go/src/net/dnsconfig_unix.go:61
			_go_fuzz_dep_.CoverTab[13756]++
									if len(f) > 1 {
//line /usr/local/go/src/net/dnsconfig_unix.go:62
				_go_fuzz_dep_.CoverTab[13768]++
										conf.search = []string{ensureRooted(f[1])}
//line /usr/local/go/src/net/dnsconfig_unix.go:63
				// _ = "end of CoverTab[13768]"
			} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:64
				_go_fuzz_dep_.CoverTab[13769]++
//line /usr/local/go/src/net/dnsconfig_unix.go:64
				// _ = "end of CoverTab[13769]"
//line /usr/local/go/src/net/dnsconfig_unix.go:64
			}
//line /usr/local/go/src/net/dnsconfig_unix.go:64
			// _ = "end of CoverTab[13756]"

		case "search":
//line /usr/local/go/src/net/dnsconfig_unix.go:66
			_go_fuzz_dep_.CoverTab[13757]++
									conf.search = make([]string, 0, len(f)-1)
									for i := 1; i < len(f); i++ {
//line /usr/local/go/src/net/dnsconfig_unix.go:68
				_go_fuzz_dep_.CoverTab[13770]++
										name := ensureRooted(f[i])
										if name == "." {
//line /usr/local/go/src/net/dnsconfig_unix.go:70
					_go_fuzz_dep_.CoverTab[13772]++
											continue
//line /usr/local/go/src/net/dnsconfig_unix.go:71
					// _ = "end of CoverTab[13772]"
				} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:72
					_go_fuzz_dep_.CoverTab[13773]++
//line /usr/local/go/src/net/dnsconfig_unix.go:72
					// _ = "end of CoverTab[13773]"
//line /usr/local/go/src/net/dnsconfig_unix.go:72
				}
//line /usr/local/go/src/net/dnsconfig_unix.go:72
				// _ = "end of CoverTab[13770]"
//line /usr/local/go/src/net/dnsconfig_unix.go:72
				_go_fuzz_dep_.CoverTab[13771]++
										conf.search = append(conf.search, name)
//line /usr/local/go/src/net/dnsconfig_unix.go:73
				// _ = "end of CoverTab[13771]"
			}
//line /usr/local/go/src/net/dnsconfig_unix.go:74
			// _ = "end of CoverTab[13757]"

		case "options":
//line /usr/local/go/src/net/dnsconfig_unix.go:76
			_go_fuzz_dep_.CoverTab[13758]++
									for _, s := range f[1:] {
//line /usr/local/go/src/net/dnsconfig_unix.go:77
				_go_fuzz_dep_.CoverTab[13774]++
										switch {
				case hasPrefix(s, "ndots:"):
//line /usr/local/go/src/net/dnsconfig_unix.go:79
					_go_fuzz_dep_.CoverTab[13775]++
											n, _, _ := dtoi(s[6:])
											if n < 0 {
//line /usr/local/go/src/net/dnsconfig_unix.go:81
						_go_fuzz_dep_.CoverTab[13788]++
												n = 0
//line /usr/local/go/src/net/dnsconfig_unix.go:82
						// _ = "end of CoverTab[13788]"
					} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:83
						_go_fuzz_dep_.CoverTab[13789]++
//line /usr/local/go/src/net/dnsconfig_unix.go:83
						if n > 15 {
//line /usr/local/go/src/net/dnsconfig_unix.go:83
							_go_fuzz_dep_.CoverTab[13790]++
													n = 15
//line /usr/local/go/src/net/dnsconfig_unix.go:84
							// _ = "end of CoverTab[13790]"
						} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:85
							_go_fuzz_dep_.CoverTab[13791]++
//line /usr/local/go/src/net/dnsconfig_unix.go:85
							// _ = "end of CoverTab[13791]"
//line /usr/local/go/src/net/dnsconfig_unix.go:85
						}
//line /usr/local/go/src/net/dnsconfig_unix.go:85
						// _ = "end of CoverTab[13789]"
//line /usr/local/go/src/net/dnsconfig_unix.go:85
					}
//line /usr/local/go/src/net/dnsconfig_unix.go:85
					// _ = "end of CoverTab[13775]"
//line /usr/local/go/src/net/dnsconfig_unix.go:85
					_go_fuzz_dep_.CoverTab[13776]++
											conf.ndots = n
//line /usr/local/go/src/net/dnsconfig_unix.go:86
					// _ = "end of CoverTab[13776]"
				case hasPrefix(s, "timeout:"):
//line /usr/local/go/src/net/dnsconfig_unix.go:87
					_go_fuzz_dep_.CoverTab[13777]++
											n, _, _ := dtoi(s[8:])
											if n < 1 {
//line /usr/local/go/src/net/dnsconfig_unix.go:89
						_go_fuzz_dep_.CoverTab[13792]++
												n = 1
//line /usr/local/go/src/net/dnsconfig_unix.go:90
						// _ = "end of CoverTab[13792]"
					} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:91
						_go_fuzz_dep_.CoverTab[13793]++
//line /usr/local/go/src/net/dnsconfig_unix.go:91
						// _ = "end of CoverTab[13793]"
//line /usr/local/go/src/net/dnsconfig_unix.go:91
					}
//line /usr/local/go/src/net/dnsconfig_unix.go:91
					// _ = "end of CoverTab[13777]"
//line /usr/local/go/src/net/dnsconfig_unix.go:91
					_go_fuzz_dep_.CoverTab[13778]++
											conf.timeout = time.Duration(n) * time.Second
//line /usr/local/go/src/net/dnsconfig_unix.go:92
					// _ = "end of CoverTab[13778]"
				case hasPrefix(s, "attempts:"):
//line /usr/local/go/src/net/dnsconfig_unix.go:93
					_go_fuzz_dep_.CoverTab[13779]++
											n, _, _ := dtoi(s[9:])
											if n < 1 {
//line /usr/local/go/src/net/dnsconfig_unix.go:95
						_go_fuzz_dep_.CoverTab[13794]++
												n = 1
//line /usr/local/go/src/net/dnsconfig_unix.go:96
						// _ = "end of CoverTab[13794]"
					} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:97
						_go_fuzz_dep_.CoverTab[13795]++
//line /usr/local/go/src/net/dnsconfig_unix.go:97
						// _ = "end of CoverTab[13795]"
//line /usr/local/go/src/net/dnsconfig_unix.go:97
					}
//line /usr/local/go/src/net/dnsconfig_unix.go:97
					// _ = "end of CoverTab[13779]"
//line /usr/local/go/src/net/dnsconfig_unix.go:97
					_go_fuzz_dep_.CoverTab[13780]++
											conf.attempts = n
//line /usr/local/go/src/net/dnsconfig_unix.go:98
					// _ = "end of CoverTab[13780]"
				case s == "rotate":
//line /usr/local/go/src/net/dnsconfig_unix.go:99
					_go_fuzz_dep_.CoverTab[13781]++
											conf.rotate = true
//line /usr/local/go/src/net/dnsconfig_unix.go:100
					// _ = "end of CoverTab[13781]"
				case s == "single-request" || func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:101
					_go_fuzz_dep_.CoverTab[13796]++
//line /usr/local/go/src/net/dnsconfig_unix.go:101
					return s == "single-request-reopen"
//line /usr/local/go/src/net/dnsconfig_unix.go:101
					// _ = "end of CoverTab[13796]"
//line /usr/local/go/src/net/dnsconfig_unix.go:101
				}():
//line /usr/local/go/src/net/dnsconfig_unix.go:101
					_go_fuzz_dep_.CoverTab[13782]++

//line /usr/local/go/src/net/dnsconfig_unix.go:107
					conf.singleRequest = true
//line /usr/local/go/src/net/dnsconfig_unix.go:107
					// _ = "end of CoverTab[13782]"
				case s == "use-vc" || func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					_go_fuzz_dep_.CoverTab[13797]++
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					return s == "usevc"
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					// _ = "end of CoverTab[13797]"
//line /usr/local/go/src/net/dnsconfig_unix.go:108
				}() || func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					_go_fuzz_dep_.CoverTab[13798]++
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					return s == "tcp"
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					// _ = "end of CoverTab[13798]"
//line /usr/local/go/src/net/dnsconfig_unix.go:108
				}():
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					_go_fuzz_dep_.CoverTab[13783]++

//line /usr/local/go/src/net/dnsconfig_unix.go:115
					conf.useTCP = true
//line /usr/local/go/src/net/dnsconfig_unix.go:115
					// _ = "end of CoverTab[13783]"
				case s == "trust-ad":
//line /usr/local/go/src/net/dnsconfig_unix.go:116
					_go_fuzz_dep_.CoverTab[13784]++
											conf.trustAD = true
//line /usr/local/go/src/net/dnsconfig_unix.go:117
					// _ = "end of CoverTab[13784]"
				case s == "edns0":
//line /usr/local/go/src/net/dnsconfig_unix.go:118
					_go_fuzz_dep_.CoverTab[13785]++
//line /usr/local/go/src/net/dnsconfig_unix.go:118
					// _ = "end of CoverTab[13785]"

//line /usr/local/go/src/net/dnsconfig_unix.go:121
				case s == "no-reload":
//line /usr/local/go/src/net/dnsconfig_unix.go:121
					_go_fuzz_dep_.CoverTab[13786]++
											conf.noReload = true
//line /usr/local/go/src/net/dnsconfig_unix.go:122
					// _ = "end of CoverTab[13786]"
				default:
//line /usr/local/go/src/net/dnsconfig_unix.go:123
					_go_fuzz_dep_.CoverTab[13787]++
											conf.unknownOpt = true
//line /usr/local/go/src/net/dnsconfig_unix.go:124
					// _ = "end of CoverTab[13787]"
				}
//line /usr/local/go/src/net/dnsconfig_unix.go:125
				// _ = "end of CoverTab[13774]"
			}
//line /usr/local/go/src/net/dnsconfig_unix.go:126
			// _ = "end of CoverTab[13758]"

		case "lookup":
//line /usr/local/go/src/net/dnsconfig_unix.go:128
			_go_fuzz_dep_.CoverTab[13759]++

//line /usr/local/go/src/net/dnsconfig_unix.go:132
			conf.lookup = f[1:]
//line /usr/local/go/src/net/dnsconfig_unix.go:132
			// _ = "end of CoverTab[13759]"

		default:
//line /usr/local/go/src/net/dnsconfig_unix.go:134
			_go_fuzz_dep_.CoverTab[13760]++
									conf.unknownOpt = true
//line /usr/local/go/src/net/dnsconfig_unix.go:135
			// _ = "end of CoverTab[13760]"
		}
//line /usr/local/go/src/net/dnsconfig_unix.go:136
		// _ = "end of CoverTab[13748]"
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:137
	// _ = "end of CoverTab[13738]"
//line /usr/local/go/src/net/dnsconfig_unix.go:137
	_go_fuzz_dep_.CoverTab[13739]++
							if len(conf.servers) == 0 {
//line /usr/local/go/src/net/dnsconfig_unix.go:138
		_go_fuzz_dep_.CoverTab[13799]++
								conf.servers = defaultNS
//line /usr/local/go/src/net/dnsconfig_unix.go:139
		// _ = "end of CoverTab[13799]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:140
		_go_fuzz_dep_.CoverTab[13800]++
//line /usr/local/go/src/net/dnsconfig_unix.go:140
		// _ = "end of CoverTab[13800]"
//line /usr/local/go/src/net/dnsconfig_unix.go:140
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:140
	// _ = "end of CoverTab[13739]"
//line /usr/local/go/src/net/dnsconfig_unix.go:140
	_go_fuzz_dep_.CoverTab[13740]++
							if len(conf.search) == 0 {
//line /usr/local/go/src/net/dnsconfig_unix.go:141
		_go_fuzz_dep_.CoverTab[13801]++
								conf.search = dnsDefaultSearch()
//line /usr/local/go/src/net/dnsconfig_unix.go:142
		// _ = "end of CoverTab[13801]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:143
		_go_fuzz_dep_.CoverTab[13802]++
//line /usr/local/go/src/net/dnsconfig_unix.go:143
		// _ = "end of CoverTab[13802]"
//line /usr/local/go/src/net/dnsconfig_unix.go:143
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:143
	// _ = "end of CoverTab[13740]"
//line /usr/local/go/src/net/dnsconfig_unix.go:143
	_go_fuzz_dep_.CoverTab[13741]++
							return conf
//line /usr/local/go/src/net/dnsconfig_unix.go:144
	// _ = "end of CoverTab[13741]"
}

func dnsDefaultSearch() []string {
//line /usr/local/go/src/net/dnsconfig_unix.go:147
	_go_fuzz_dep_.CoverTab[13803]++
							hn, err := getHostname()
							if err != nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:149
		_go_fuzz_dep_.CoverTab[13806]++

								return nil
//line /usr/local/go/src/net/dnsconfig_unix.go:151
		// _ = "end of CoverTab[13806]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:152
		_go_fuzz_dep_.CoverTab[13807]++
//line /usr/local/go/src/net/dnsconfig_unix.go:152
		// _ = "end of CoverTab[13807]"
//line /usr/local/go/src/net/dnsconfig_unix.go:152
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:152
	// _ = "end of CoverTab[13803]"
//line /usr/local/go/src/net/dnsconfig_unix.go:152
	_go_fuzz_dep_.CoverTab[13804]++
							if i := bytealg.IndexByteString(hn, '.'); i >= 0 && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:153
		_go_fuzz_dep_.CoverTab[13808]++
//line /usr/local/go/src/net/dnsconfig_unix.go:153
		return i < len(hn)-1
//line /usr/local/go/src/net/dnsconfig_unix.go:153
		// _ = "end of CoverTab[13808]"
//line /usr/local/go/src/net/dnsconfig_unix.go:153
	}() {
//line /usr/local/go/src/net/dnsconfig_unix.go:153
		_go_fuzz_dep_.CoverTab[13809]++
								return []string{ensureRooted(hn[i+1:])}
//line /usr/local/go/src/net/dnsconfig_unix.go:154
		// _ = "end of CoverTab[13809]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:155
		_go_fuzz_dep_.CoverTab[13810]++
//line /usr/local/go/src/net/dnsconfig_unix.go:155
		// _ = "end of CoverTab[13810]"
//line /usr/local/go/src/net/dnsconfig_unix.go:155
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:155
	// _ = "end of CoverTab[13804]"
//line /usr/local/go/src/net/dnsconfig_unix.go:155
	_go_fuzz_dep_.CoverTab[13805]++
							return nil
//line /usr/local/go/src/net/dnsconfig_unix.go:156
	// _ = "end of CoverTab[13805]"
}

func hasPrefix(s, prefix string) bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:159
	_go_fuzz_dep_.CoverTab[13811]++
							return len(s) >= len(prefix) && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:160
		_go_fuzz_dep_.CoverTab[13812]++
//line /usr/local/go/src/net/dnsconfig_unix.go:160
		return s[:len(prefix)] == prefix
//line /usr/local/go/src/net/dnsconfig_unix.go:160
		// _ = "end of CoverTab[13812]"
//line /usr/local/go/src/net/dnsconfig_unix.go:160
	}()
//line /usr/local/go/src/net/dnsconfig_unix.go:160
	// _ = "end of CoverTab[13811]"
}

func ensureRooted(s string) string {
//line /usr/local/go/src/net/dnsconfig_unix.go:163
	_go_fuzz_dep_.CoverTab[13813]++
							if len(s) > 0 && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:164
		_go_fuzz_dep_.CoverTab[13815]++
//line /usr/local/go/src/net/dnsconfig_unix.go:164
		return s[len(s)-1] == '.'
//line /usr/local/go/src/net/dnsconfig_unix.go:164
		// _ = "end of CoverTab[13815]"
//line /usr/local/go/src/net/dnsconfig_unix.go:164
	}() {
//line /usr/local/go/src/net/dnsconfig_unix.go:164
		_go_fuzz_dep_.CoverTab[13816]++
								return s
//line /usr/local/go/src/net/dnsconfig_unix.go:165
		// _ = "end of CoverTab[13816]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:166
		_go_fuzz_dep_.CoverTab[13817]++
//line /usr/local/go/src/net/dnsconfig_unix.go:166
		// _ = "end of CoverTab[13817]"
//line /usr/local/go/src/net/dnsconfig_unix.go:166
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:166
	// _ = "end of CoverTab[13813]"
//line /usr/local/go/src/net/dnsconfig_unix.go:166
	_go_fuzz_dep_.CoverTab[13814]++
							return s + "."
//line /usr/local/go/src/net/dnsconfig_unix.go:167
	// _ = "end of CoverTab[13814]"
}

//line /usr/local/go/src/net/dnsconfig_unix.go:168
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/dnsconfig_unix.go:168
var _ = _go_fuzz_dep_.CoverTab
