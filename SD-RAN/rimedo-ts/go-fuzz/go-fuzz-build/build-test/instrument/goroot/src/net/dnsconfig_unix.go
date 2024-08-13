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
	_go_fuzz_dep_.CoverTab[5346]++
							conf := &dnsConfig{
		ndots:		1,
		timeout:	5 * time.Second,
		attempts:	2,
	}
	file, err := open(filename)
	if err != nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:24
		_go_fuzz_dep_.CoverTab[5352]++
								conf.servers = defaultNS
								conf.search = dnsDefaultSearch()
								conf.err = err
								return conf
//line /usr/local/go/src/net/dnsconfig_unix.go:28
		// _ = "end of CoverTab[5352]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:29
		_go_fuzz_dep_.CoverTab[5353]++
//line /usr/local/go/src/net/dnsconfig_unix.go:29
		// _ = "end of CoverTab[5353]"
//line /usr/local/go/src/net/dnsconfig_unix.go:29
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:29
	// _ = "end of CoverTab[5346]"
//line /usr/local/go/src/net/dnsconfig_unix.go:29
	_go_fuzz_dep_.CoverTab[5347]++
							defer file.close()
							if fi, err := file.file.Stat(); err == nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:31
		_go_fuzz_dep_.CoverTab[5354]++
								conf.mtime = fi.ModTime()
//line /usr/local/go/src/net/dnsconfig_unix.go:32
		// _ = "end of CoverTab[5354]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:33
		_go_fuzz_dep_.CoverTab[5355]++
								conf.servers = defaultNS
								conf.search = dnsDefaultSearch()
								conf.err = err
								return conf
//line /usr/local/go/src/net/dnsconfig_unix.go:37
		// _ = "end of CoverTab[5355]"
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:38
	// _ = "end of CoverTab[5347]"
//line /usr/local/go/src/net/dnsconfig_unix.go:38
	_go_fuzz_dep_.CoverTab[5348]++
							for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /usr/local/go/src/net/dnsconfig_unix.go:39
		_go_fuzz_dep_.CoverTab[5356]++
								if len(line) > 0 && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			_go_fuzz_dep_.CoverTab[5359]++
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			return (line[0] == ';' || func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:40
				_go_fuzz_dep_.CoverTab[5360]++
//line /usr/local/go/src/net/dnsconfig_unix.go:40
				return line[0] == '#'
//line /usr/local/go/src/net/dnsconfig_unix.go:40
				// _ = "end of CoverTab[5360]"
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			}())
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			// _ = "end of CoverTab[5359]"
//line /usr/local/go/src/net/dnsconfig_unix.go:40
		}() {
//line /usr/local/go/src/net/dnsconfig_unix.go:40
			_go_fuzz_dep_.CoverTab[5361]++

									continue
//line /usr/local/go/src/net/dnsconfig_unix.go:42
			// _ = "end of CoverTab[5361]"
		} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:43
			_go_fuzz_dep_.CoverTab[5362]++
//line /usr/local/go/src/net/dnsconfig_unix.go:43
			// _ = "end of CoverTab[5362]"
//line /usr/local/go/src/net/dnsconfig_unix.go:43
		}
//line /usr/local/go/src/net/dnsconfig_unix.go:43
		// _ = "end of CoverTab[5356]"
//line /usr/local/go/src/net/dnsconfig_unix.go:43
		_go_fuzz_dep_.CoverTab[5357]++
								f := getFields(line)
								if len(f) < 1 {
//line /usr/local/go/src/net/dnsconfig_unix.go:45
			_go_fuzz_dep_.CoverTab[5363]++
									continue
//line /usr/local/go/src/net/dnsconfig_unix.go:46
			// _ = "end of CoverTab[5363]"
		} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:47
			_go_fuzz_dep_.CoverTab[5364]++
//line /usr/local/go/src/net/dnsconfig_unix.go:47
			// _ = "end of CoverTab[5364]"
//line /usr/local/go/src/net/dnsconfig_unix.go:47
		}
//line /usr/local/go/src/net/dnsconfig_unix.go:47
		// _ = "end of CoverTab[5357]"
//line /usr/local/go/src/net/dnsconfig_unix.go:47
		_go_fuzz_dep_.CoverTab[5358]++
								switch f[0] {
		case "nameserver":
//line /usr/local/go/src/net/dnsconfig_unix.go:49
			_go_fuzz_dep_.CoverTab[5365]++
									if len(f) > 1 && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:50
				_go_fuzz_dep_.CoverTab[5371]++
//line /usr/local/go/src/net/dnsconfig_unix.go:50
				return len(conf.servers) < 3
//line /usr/local/go/src/net/dnsconfig_unix.go:50
				// _ = "end of CoverTab[5371]"
//line /usr/local/go/src/net/dnsconfig_unix.go:50
			}() {
//line /usr/local/go/src/net/dnsconfig_unix.go:50
				_go_fuzz_dep_.CoverTab[5372]++

//line /usr/local/go/src/net/dnsconfig_unix.go:54
				if parseIPv4(f[1]) != nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:54
					_go_fuzz_dep_.CoverTab[5373]++
											conf.servers = append(conf.servers, JoinHostPort(f[1], "53"))
//line /usr/local/go/src/net/dnsconfig_unix.go:55
					// _ = "end of CoverTab[5373]"
				} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:56
					_go_fuzz_dep_.CoverTab[5374]++
//line /usr/local/go/src/net/dnsconfig_unix.go:56
					if ip, _ := parseIPv6Zone(f[1]); ip != nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:56
						_go_fuzz_dep_.CoverTab[5375]++
												conf.servers = append(conf.servers, JoinHostPort(f[1], "53"))
//line /usr/local/go/src/net/dnsconfig_unix.go:57
						// _ = "end of CoverTab[5375]"
					} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:58
						_go_fuzz_dep_.CoverTab[5376]++
//line /usr/local/go/src/net/dnsconfig_unix.go:58
						// _ = "end of CoverTab[5376]"
//line /usr/local/go/src/net/dnsconfig_unix.go:58
					}
//line /usr/local/go/src/net/dnsconfig_unix.go:58
					// _ = "end of CoverTab[5374]"
//line /usr/local/go/src/net/dnsconfig_unix.go:58
				}
//line /usr/local/go/src/net/dnsconfig_unix.go:58
				// _ = "end of CoverTab[5372]"
			} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:59
				_go_fuzz_dep_.CoverTab[5377]++
//line /usr/local/go/src/net/dnsconfig_unix.go:59
				// _ = "end of CoverTab[5377]"
//line /usr/local/go/src/net/dnsconfig_unix.go:59
			}
//line /usr/local/go/src/net/dnsconfig_unix.go:59
			// _ = "end of CoverTab[5365]"

		case "domain":
//line /usr/local/go/src/net/dnsconfig_unix.go:61
			_go_fuzz_dep_.CoverTab[5366]++
									if len(f) > 1 {
//line /usr/local/go/src/net/dnsconfig_unix.go:62
				_go_fuzz_dep_.CoverTab[5378]++
										conf.search = []string{ensureRooted(f[1])}
//line /usr/local/go/src/net/dnsconfig_unix.go:63
				// _ = "end of CoverTab[5378]"
			} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:64
				_go_fuzz_dep_.CoverTab[5379]++
//line /usr/local/go/src/net/dnsconfig_unix.go:64
				// _ = "end of CoverTab[5379]"
//line /usr/local/go/src/net/dnsconfig_unix.go:64
			}
//line /usr/local/go/src/net/dnsconfig_unix.go:64
			// _ = "end of CoverTab[5366]"

		case "search":
//line /usr/local/go/src/net/dnsconfig_unix.go:66
			_go_fuzz_dep_.CoverTab[5367]++
									conf.search = make([]string, 0, len(f)-1)
									for i := 1; i < len(f); i++ {
//line /usr/local/go/src/net/dnsconfig_unix.go:68
				_go_fuzz_dep_.CoverTab[5380]++
										name := ensureRooted(f[i])
										if name == "." {
//line /usr/local/go/src/net/dnsconfig_unix.go:70
					_go_fuzz_dep_.CoverTab[5382]++
											continue
//line /usr/local/go/src/net/dnsconfig_unix.go:71
					// _ = "end of CoverTab[5382]"
				} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:72
					_go_fuzz_dep_.CoverTab[5383]++
//line /usr/local/go/src/net/dnsconfig_unix.go:72
					// _ = "end of CoverTab[5383]"
//line /usr/local/go/src/net/dnsconfig_unix.go:72
				}
//line /usr/local/go/src/net/dnsconfig_unix.go:72
				// _ = "end of CoverTab[5380]"
//line /usr/local/go/src/net/dnsconfig_unix.go:72
				_go_fuzz_dep_.CoverTab[5381]++
										conf.search = append(conf.search, name)
//line /usr/local/go/src/net/dnsconfig_unix.go:73
				// _ = "end of CoverTab[5381]"
			}
//line /usr/local/go/src/net/dnsconfig_unix.go:74
			// _ = "end of CoverTab[5367]"

		case "options":
//line /usr/local/go/src/net/dnsconfig_unix.go:76
			_go_fuzz_dep_.CoverTab[5368]++
									for _, s := range f[1:] {
//line /usr/local/go/src/net/dnsconfig_unix.go:77
				_go_fuzz_dep_.CoverTab[5384]++
										switch {
				case hasPrefix(s, "ndots:"):
//line /usr/local/go/src/net/dnsconfig_unix.go:79
					_go_fuzz_dep_.CoverTab[5385]++
											n, _, _ := dtoi(s[6:])
											if n < 0 {
//line /usr/local/go/src/net/dnsconfig_unix.go:81
						_go_fuzz_dep_.CoverTab[5398]++
												n = 0
//line /usr/local/go/src/net/dnsconfig_unix.go:82
						// _ = "end of CoverTab[5398]"
					} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:83
						_go_fuzz_dep_.CoverTab[5399]++
//line /usr/local/go/src/net/dnsconfig_unix.go:83
						if n > 15 {
//line /usr/local/go/src/net/dnsconfig_unix.go:83
							_go_fuzz_dep_.CoverTab[5400]++
													n = 15
//line /usr/local/go/src/net/dnsconfig_unix.go:84
							// _ = "end of CoverTab[5400]"
						} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:85
							_go_fuzz_dep_.CoverTab[5401]++
//line /usr/local/go/src/net/dnsconfig_unix.go:85
							// _ = "end of CoverTab[5401]"
//line /usr/local/go/src/net/dnsconfig_unix.go:85
						}
//line /usr/local/go/src/net/dnsconfig_unix.go:85
						// _ = "end of CoverTab[5399]"
//line /usr/local/go/src/net/dnsconfig_unix.go:85
					}
//line /usr/local/go/src/net/dnsconfig_unix.go:85
					// _ = "end of CoverTab[5385]"
//line /usr/local/go/src/net/dnsconfig_unix.go:85
					_go_fuzz_dep_.CoverTab[5386]++
											conf.ndots = n
//line /usr/local/go/src/net/dnsconfig_unix.go:86
					// _ = "end of CoverTab[5386]"
				case hasPrefix(s, "timeout:"):
//line /usr/local/go/src/net/dnsconfig_unix.go:87
					_go_fuzz_dep_.CoverTab[5387]++
											n, _, _ := dtoi(s[8:])
											if n < 1 {
//line /usr/local/go/src/net/dnsconfig_unix.go:89
						_go_fuzz_dep_.CoverTab[5402]++
												n = 1
//line /usr/local/go/src/net/dnsconfig_unix.go:90
						// _ = "end of CoverTab[5402]"
					} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:91
						_go_fuzz_dep_.CoverTab[5403]++
//line /usr/local/go/src/net/dnsconfig_unix.go:91
						// _ = "end of CoverTab[5403]"
//line /usr/local/go/src/net/dnsconfig_unix.go:91
					}
//line /usr/local/go/src/net/dnsconfig_unix.go:91
					// _ = "end of CoverTab[5387]"
//line /usr/local/go/src/net/dnsconfig_unix.go:91
					_go_fuzz_dep_.CoverTab[5388]++
											conf.timeout = time.Duration(n) * time.Second
//line /usr/local/go/src/net/dnsconfig_unix.go:92
					// _ = "end of CoverTab[5388]"
				case hasPrefix(s, "attempts:"):
//line /usr/local/go/src/net/dnsconfig_unix.go:93
					_go_fuzz_dep_.CoverTab[5389]++
											n, _, _ := dtoi(s[9:])
											if n < 1 {
//line /usr/local/go/src/net/dnsconfig_unix.go:95
						_go_fuzz_dep_.CoverTab[5404]++
												n = 1
//line /usr/local/go/src/net/dnsconfig_unix.go:96
						// _ = "end of CoverTab[5404]"
					} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:97
						_go_fuzz_dep_.CoverTab[5405]++
//line /usr/local/go/src/net/dnsconfig_unix.go:97
						// _ = "end of CoverTab[5405]"
//line /usr/local/go/src/net/dnsconfig_unix.go:97
					}
//line /usr/local/go/src/net/dnsconfig_unix.go:97
					// _ = "end of CoverTab[5389]"
//line /usr/local/go/src/net/dnsconfig_unix.go:97
					_go_fuzz_dep_.CoverTab[5390]++
											conf.attempts = n
//line /usr/local/go/src/net/dnsconfig_unix.go:98
					// _ = "end of CoverTab[5390]"
				case s == "rotate":
//line /usr/local/go/src/net/dnsconfig_unix.go:99
					_go_fuzz_dep_.CoverTab[5391]++
											conf.rotate = true
//line /usr/local/go/src/net/dnsconfig_unix.go:100
					// _ = "end of CoverTab[5391]"
				case s == "single-request" || func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:101
					_go_fuzz_dep_.CoverTab[5406]++
//line /usr/local/go/src/net/dnsconfig_unix.go:101
					return s == "single-request-reopen"
//line /usr/local/go/src/net/dnsconfig_unix.go:101
					// _ = "end of CoverTab[5406]"
//line /usr/local/go/src/net/dnsconfig_unix.go:101
				}():
//line /usr/local/go/src/net/dnsconfig_unix.go:101
					_go_fuzz_dep_.CoverTab[5392]++

//line /usr/local/go/src/net/dnsconfig_unix.go:107
					conf.singleRequest = true
//line /usr/local/go/src/net/dnsconfig_unix.go:107
					// _ = "end of CoverTab[5392]"
				case s == "use-vc" || func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					_go_fuzz_dep_.CoverTab[5407]++
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					return s == "usevc"
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					// _ = "end of CoverTab[5407]"
//line /usr/local/go/src/net/dnsconfig_unix.go:108
				}() || func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					_go_fuzz_dep_.CoverTab[5408]++
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					return s == "tcp"
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					// _ = "end of CoverTab[5408]"
//line /usr/local/go/src/net/dnsconfig_unix.go:108
				}():
//line /usr/local/go/src/net/dnsconfig_unix.go:108
					_go_fuzz_dep_.CoverTab[5393]++

//line /usr/local/go/src/net/dnsconfig_unix.go:115
					conf.useTCP = true
//line /usr/local/go/src/net/dnsconfig_unix.go:115
					// _ = "end of CoverTab[5393]"
				case s == "trust-ad":
//line /usr/local/go/src/net/dnsconfig_unix.go:116
					_go_fuzz_dep_.CoverTab[5394]++
											conf.trustAD = true
//line /usr/local/go/src/net/dnsconfig_unix.go:117
					// _ = "end of CoverTab[5394]"
				case s == "edns0":
//line /usr/local/go/src/net/dnsconfig_unix.go:118
					_go_fuzz_dep_.CoverTab[5395]++
//line /usr/local/go/src/net/dnsconfig_unix.go:118
					// _ = "end of CoverTab[5395]"

//line /usr/local/go/src/net/dnsconfig_unix.go:121
				case s == "no-reload":
//line /usr/local/go/src/net/dnsconfig_unix.go:121
					_go_fuzz_dep_.CoverTab[5396]++
											conf.noReload = true
//line /usr/local/go/src/net/dnsconfig_unix.go:122
					// _ = "end of CoverTab[5396]"
				default:
//line /usr/local/go/src/net/dnsconfig_unix.go:123
					_go_fuzz_dep_.CoverTab[5397]++
											conf.unknownOpt = true
//line /usr/local/go/src/net/dnsconfig_unix.go:124
					// _ = "end of CoverTab[5397]"
				}
//line /usr/local/go/src/net/dnsconfig_unix.go:125
				// _ = "end of CoverTab[5384]"
			}
//line /usr/local/go/src/net/dnsconfig_unix.go:126
			// _ = "end of CoverTab[5368]"

		case "lookup":
//line /usr/local/go/src/net/dnsconfig_unix.go:128
			_go_fuzz_dep_.CoverTab[5369]++

//line /usr/local/go/src/net/dnsconfig_unix.go:132
			conf.lookup = f[1:]
//line /usr/local/go/src/net/dnsconfig_unix.go:132
			// _ = "end of CoverTab[5369]"

		default:
//line /usr/local/go/src/net/dnsconfig_unix.go:134
			_go_fuzz_dep_.CoverTab[5370]++
									conf.unknownOpt = true
//line /usr/local/go/src/net/dnsconfig_unix.go:135
			// _ = "end of CoverTab[5370]"
		}
//line /usr/local/go/src/net/dnsconfig_unix.go:136
		// _ = "end of CoverTab[5358]"
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:137
	// _ = "end of CoverTab[5348]"
//line /usr/local/go/src/net/dnsconfig_unix.go:137
	_go_fuzz_dep_.CoverTab[5349]++
							if len(conf.servers) == 0 {
//line /usr/local/go/src/net/dnsconfig_unix.go:138
		_go_fuzz_dep_.CoverTab[5409]++
								conf.servers = defaultNS
//line /usr/local/go/src/net/dnsconfig_unix.go:139
		// _ = "end of CoverTab[5409]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:140
		_go_fuzz_dep_.CoverTab[5410]++
//line /usr/local/go/src/net/dnsconfig_unix.go:140
		// _ = "end of CoverTab[5410]"
//line /usr/local/go/src/net/dnsconfig_unix.go:140
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:140
	// _ = "end of CoverTab[5349]"
//line /usr/local/go/src/net/dnsconfig_unix.go:140
	_go_fuzz_dep_.CoverTab[5350]++
							if len(conf.search) == 0 {
//line /usr/local/go/src/net/dnsconfig_unix.go:141
		_go_fuzz_dep_.CoverTab[5411]++
								conf.search = dnsDefaultSearch()
//line /usr/local/go/src/net/dnsconfig_unix.go:142
		// _ = "end of CoverTab[5411]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:143
		_go_fuzz_dep_.CoverTab[5412]++
//line /usr/local/go/src/net/dnsconfig_unix.go:143
		// _ = "end of CoverTab[5412]"
//line /usr/local/go/src/net/dnsconfig_unix.go:143
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:143
	// _ = "end of CoverTab[5350]"
//line /usr/local/go/src/net/dnsconfig_unix.go:143
	_go_fuzz_dep_.CoverTab[5351]++
							return conf
//line /usr/local/go/src/net/dnsconfig_unix.go:144
	// _ = "end of CoverTab[5351]"
}

func dnsDefaultSearch() []string {
//line /usr/local/go/src/net/dnsconfig_unix.go:147
	_go_fuzz_dep_.CoverTab[5413]++
							hn, err := getHostname()
							if err != nil {
//line /usr/local/go/src/net/dnsconfig_unix.go:149
		_go_fuzz_dep_.CoverTab[5416]++

								return nil
//line /usr/local/go/src/net/dnsconfig_unix.go:151
		// _ = "end of CoverTab[5416]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:152
		_go_fuzz_dep_.CoverTab[5417]++
//line /usr/local/go/src/net/dnsconfig_unix.go:152
		// _ = "end of CoverTab[5417]"
//line /usr/local/go/src/net/dnsconfig_unix.go:152
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:152
	// _ = "end of CoverTab[5413]"
//line /usr/local/go/src/net/dnsconfig_unix.go:152
	_go_fuzz_dep_.CoverTab[5414]++
							if i := bytealg.IndexByteString(hn, '.'); i >= 0 && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:153
		_go_fuzz_dep_.CoverTab[5418]++
//line /usr/local/go/src/net/dnsconfig_unix.go:153
		return i < len(hn)-1
//line /usr/local/go/src/net/dnsconfig_unix.go:153
		// _ = "end of CoverTab[5418]"
//line /usr/local/go/src/net/dnsconfig_unix.go:153
	}() {
//line /usr/local/go/src/net/dnsconfig_unix.go:153
		_go_fuzz_dep_.CoverTab[5419]++
								return []string{ensureRooted(hn[i+1:])}
//line /usr/local/go/src/net/dnsconfig_unix.go:154
		// _ = "end of CoverTab[5419]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:155
		_go_fuzz_dep_.CoverTab[5420]++
//line /usr/local/go/src/net/dnsconfig_unix.go:155
		// _ = "end of CoverTab[5420]"
//line /usr/local/go/src/net/dnsconfig_unix.go:155
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:155
	// _ = "end of CoverTab[5414]"
//line /usr/local/go/src/net/dnsconfig_unix.go:155
	_go_fuzz_dep_.CoverTab[5415]++
							return nil
//line /usr/local/go/src/net/dnsconfig_unix.go:156
	// _ = "end of CoverTab[5415]"
}

func hasPrefix(s, prefix string) bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:159
	_go_fuzz_dep_.CoverTab[5421]++
							return len(s) >= len(prefix) && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:160
		_go_fuzz_dep_.CoverTab[5422]++
//line /usr/local/go/src/net/dnsconfig_unix.go:160
		return s[:len(prefix)] == prefix
//line /usr/local/go/src/net/dnsconfig_unix.go:160
		// _ = "end of CoverTab[5422]"
//line /usr/local/go/src/net/dnsconfig_unix.go:160
	}()
//line /usr/local/go/src/net/dnsconfig_unix.go:160
	// _ = "end of CoverTab[5421]"
}

func ensureRooted(s string) string {
//line /usr/local/go/src/net/dnsconfig_unix.go:163
	_go_fuzz_dep_.CoverTab[5423]++
							if len(s) > 0 && func() bool {
//line /usr/local/go/src/net/dnsconfig_unix.go:164
		_go_fuzz_dep_.CoverTab[5425]++
//line /usr/local/go/src/net/dnsconfig_unix.go:164
		return s[len(s)-1] == '.'
//line /usr/local/go/src/net/dnsconfig_unix.go:164
		// _ = "end of CoverTab[5425]"
//line /usr/local/go/src/net/dnsconfig_unix.go:164
	}() {
//line /usr/local/go/src/net/dnsconfig_unix.go:164
		_go_fuzz_dep_.CoverTab[5426]++
								return s
//line /usr/local/go/src/net/dnsconfig_unix.go:165
		// _ = "end of CoverTab[5426]"
	} else {
//line /usr/local/go/src/net/dnsconfig_unix.go:166
		_go_fuzz_dep_.CoverTab[5427]++
//line /usr/local/go/src/net/dnsconfig_unix.go:166
		// _ = "end of CoverTab[5427]"
//line /usr/local/go/src/net/dnsconfig_unix.go:166
	}
//line /usr/local/go/src/net/dnsconfig_unix.go:166
	// _ = "end of CoverTab[5423]"
//line /usr/local/go/src/net/dnsconfig_unix.go:166
	_go_fuzz_dep_.CoverTab[5424]++
							return s + "."
//line /usr/local/go/src/net/dnsconfig_unix.go:167
	// _ = "end of CoverTab[5424]"
}

//line /usr/local/go/src/net/dnsconfig_unix.go:168
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/dnsconfig_unix.go:168
var _ = _go_fuzz_dep_.CoverTab
