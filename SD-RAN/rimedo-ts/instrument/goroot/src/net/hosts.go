// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/hosts.go:5
package net

//line /usr/local/go/src/net/hosts.go:5
import (
//line /usr/local/go/src/net/hosts.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/hosts.go:5
)
//line /usr/local/go/src/net/hosts.go:5
import (
//line /usr/local/go/src/net/hosts.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/hosts.go:5
)

import (
	"internal/bytealg"
	"sync"
	"time"
)

const cacheMaxAge = 5 * time.Second

func parseLiteralIP(addr string) string {
//line /usr/local/go/src/net/hosts.go:15
	_go_fuzz_dep_.CoverTab[13999]++
						var ip IP
						var zone string
						ip = parseIPv4(addr)
						if ip == nil {
//line /usr/local/go/src/net/hosts.go:19
		_go_fuzz_dep_.CoverTab[14003]++
							ip, zone = parseIPv6Zone(addr)
//line /usr/local/go/src/net/hosts.go:20
		// _ = "end of CoverTab[14003]"
	} else {
//line /usr/local/go/src/net/hosts.go:21
		_go_fuzz_dep_.CoverTab[14004]++
//line /usr/local/go/src/net/hosts.go:21
		// _ = "end of CoverTab[14004]"
//line /usr/local/go/src/net/hosts.go:21
	}
//line /usr/local/go/src/net/hosts.go:21
	// _ = "end of CoverTab[13999]"
//line /usr/local/go/src/net/hosts.go:21
	_go_fuzz_dep_.CoverTab[14000]++
						if ip == nil {
//line /usr/local/go/src/net/hosts.go:22
		_go_fuzz_dep_.CoverTab[14005]++
							return ""
//line /usr/local/go/src/net/hosts.go:23
		// _ = "end of CoverTab[14005]"
	} else {
//line /usr/local/go/src/net/hosts.go:24
		_go_fuzz_dep_.CoverTab[14006]++
//line /usr/local/go/src/net/hosts.go:24
		// _ = "end of CoverTab[14006]"
//line /usr/local/go/src/net/hosts.go:24
	}
//line /usr/local/go/src/net/hosts.go:24
	// _ = "end of CoverTab[14000]"
//line /usr/local/go/src/net/hosts.go:24
	_go_fuzz_dep_.CoverTab[14001]++
						if zone == "" {
//line /usr/local/go/src/net/hosts.go:25
		_go_fuzz_dep_.CoverTab[14007]++
							return ip.String()
//line /usr/local/go/src/net/hosts.go:26
		// _ = "end of CoverTab[14007]"
	} else {
//line /usr/local/go/src/net/hosts.go:27
		_go_fuzz_dep_.CoverTab[14008]++
//line /usr/local/go/src/net/hosts.go:27
		// _ = "end of CoverTab[14008]"
//line /usr/local/go/src/net/hosts.go:27
	}
//line /usr/local/go/src/net/hosts.go:27
	// _ = "end of CoverTab[14001]"
//line /usr/local/go/src/net/hosts.go:27
	_go_fuzz_dep_.CoverTab[14002]++
						return ip.String() + "%" + zone
//line /usr/local/go/src/net/hosts.go:28
	// _ = "end of CoverTab[14002]"
}

type byName struct {
	addrs		[]string
	canonicalName	string
}

// hosts contains known host entries.
var hosts struct {
	sync.Mutex

	// Key for the list of literal IP addresses must be a host
	// name. It would be part of DNS labels, a FQDN or an absolute
	// FQDN.
	// For now the key is converted to lower case for convenience.
	byName	map[string]byName

	// Key for the list of host names must be a literal IP address
	// including IPv6 address with zone identifier.
	// We don't support old-classful IP address notation.
	byAddr	map[string][]string

	expire	time.Time
	path	string
	mtime	time.Time
	size	int64
}

func readHosts() {
//line /usr/local/go/src/net/hosts.go:57
	_go_fuzz_dep_.CoverTab[14009]++
						now := time.Now()
						hp := testHookHostsPath

						if now.Before(hosts.expire) && func() bool {
//line /usr/local/go/src/net/hosts.go:61
		_go_fuzz_dep_.CoverTab[14014]++
//line /usr/local/go/src/net/hosts.go:61
		return hosts.path == hp
//line /usr/local/go/src/net/hosts.go:61
		// _ = "end of CoverTab[14014]"
//line /usr/local/go/src/net/hosts.go:61
	}() && func() bool {
//line /usr/local/go/src/net/hosts.go:61
		_go_fuzz_dep_.CoverTab[14015]++
//line /usr/local/go/src/net/hosts.go:61
		return len(hosts.byName) > 0
//line /usr/local/go/src/net/hosts.go:61
		// _ = "end of CoverTab[14015]"
//line /usr/local/go/src/net/hosts.go:61
	}() {
//line /usr/local/go/src/net/hosts.go:61
		_go_fuzz_dep_.CoverTab[14016]++
							return
//line /usr/local/go/src/net/hosts.go:62
		// _ = "end of CoverTab[14016]"
	} else {
//line /usr/local/go/src/net/hosts.go:63
		_go_fuzz_dep_.CoverTab[14017]++
//line /usr/local/go/src/net/hosts.go:63
		// _ = "end of CoverTab[14017]"
//line /usr/local/go/src/net/hosts.go:63
	}
//line /usr/local/go/src/net/hosts.go:63
	// _ = "end of CoverTab[14009]"
//line /usr/local/go/src/net/hosts.go:63
	_go_fuzz_dep_.CoverTab[14010]++
						mtime, size, err := stat(hp)
						if err == nil && func() bool {
//line /usr/local/go/src/net/hosts.go:65
		_go_fuzz_dep_.CoverTab[14018]++
//line /usr/local/go/src/net/hosts.go:65
		return hosts.path == hp
//line /usr/local/go/src/net/hosts.go:65
		// _ = "end of CoverTab[14018]"
//line /usr/local/go/src/net/hosts.go:65
	}() && func() bool {
//line /usr/local/go/src/net/hosts.go:65
		_go_fuzz_dep_.CoverTab[14019]++
//line /usr/local/go/src/net/hosts.go:65
		return hosts.mtime.Equal(mtime)
//line /usr/local/go/src/net/hosts.go:65
		// _ = "end of CoverTab[14019]"
//line /usr/local/go/src/net/hosts.go:65
	}() && func() bool {
//line /usr/local/go/src/net/hosts.go:65
		_go_fuzz_dep_.CoverTab[14020]++
//line /usr/local/go/src/net/hosts.go:65
		return hosts.size == size
//line /usr/local/go/src/net/hosts.go:65
		// _ = "end of CoverTab[14020]"
//line /usr/local/go/src/net/hosts.go:65
	}() {
//line /usr/local/go/src/net/hosts.go:65
		_go_fuzz_dep_.CoverTab[14021]++
							hosts.expire = now.Add(cacheMaxAge)
							return
//line /usr/local/go/src/net/hosts.go:67
		// _ = "end of CoverTab[14021]"
	} else {
//line /usr/local/go/src/net/hosts.go:68
		_go_fuzz_dep_.CoverTab[14022]++
//line /usr/local/go/src/net/hosts.go:68
		// _ = "end of CoverTab[14022]"
//line /usr/local/go/src/net/hosts.go:68
	}
//line /usr/local/go/src/net/hosts.go:68
	// _ = "end of CoverTab[14010]"
//line /usr/local/go/src/net/hosts.go:68
	_go_fuzz_dep_.CoverTab[14011]++

						hs := make(map[string]byName)
						is := make(map[string][]string)

						var file *file
						if file, _ = open(hp); file == nil {
//line /usr/local/go/src/net/hosts.go:74
		_go_fuzz_dep_.CoverTab[14023]++
							return
//line /usr/local/go/src/net/hosts.go:75
		// _ = "end of CoverTab[14023]"
	} else {
//line /usr/local/go/src/net/hosts.go:76
		_go_fuzz_dep_.CoverTab[14024]++
//line /usr/local/go/src/net/hosts.go:76
		// _ = "end of CoverTab[14024]"
//line /usr/local/go/src/net/hosts.go:76
	}
//line /usr/local/go/src/net/hosts.go:76
	// _ = "end of CoverTab[14011]"
//line /usr/local/go/src/net/hosts.go:76
	_go_fuzz_dep_.CoverTab[14012]++
						for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /usr/local/go/src/net/hosts.go:77
		_go_fuzz_dep_.CoverTab[14025]++
							if i := bytealg.IndexByteString(line, '#'); i >= 0 {
//line /usr/local/go/src/net/hosts.go:78
			_go_fuzz_dep_.CoverTab[14029]++

								line = line[0:i]
//line /usr/local/go/src/net/hosts.go:80
			// _ = "end of CoverTab[14029]"
		} else {
//line /usr/local/go/src/net/hosts.go:81
			_go_fuzz_dep_.CoverTab[14030]++
//line /usr/local/go/src/net/hosts.go:81
			// _ = "end of CoverTab[14030]"
//line /usr/local/go/src/net/hosts.go:81
		}
//line /usr/local/go/src/net/hosts.go:81
		// _ = "end of CoverTab[14025]"
//line /usr/local/go/src/net/hosts.go:81
		_go_fuzz_dep_.CoverTab[14026]++
							f := getFields(line)
							if len(f) < 2 {
//line /usr/local/go/src/net/hosts.go:83
			_go_fuzz_dep_.CoverTab[14031]++
								continue
//line /usr/local/go/src/net/hosts.go:84
			// _ = "end of CoverTab[14031]"
		} else {
//line /usr/local/go/src/net/hosts.go:85
			_go_fuzz_dep_.CoverTab[14032]++
//line /usr/local/go/src/net/hosts.go:85
			// _ = "end of CoverTab[14032]"
//line /usr/local/go/src/net/hosts.go:85
		}
//line /usr/local/go/src/net/hosts.go:85
		// _ = "end of CoverTab[14026]"
//line /usr/local/go/src/net/hosts.go:85
		_go_fuzz_dep_.CoverTab[14027]++
							addr := parseLiteralIP(f[0])
							if addr == "" {
//line /usr/local/go/src/net/hosts.go:87
			_go_fuzz_dep_.CoverTab[14033]++
								continue
//line /usr/local/go/src/net/hosts.go:88
			// _ = "end of CoverTab[14033]"
		} else {
//line /usr/local/go/src/net/hosts.go:89
			_go_fuzz_dep_.CoverTab[14034]++
//line /usr/local/go/src/net/hosts.go:89
			// _ = "end of CoverTab[14034]"
//line /usr/local/go/src/net/hosts.go:89
		}
//line /usr/local/go/src/net/hosts.go:89
		// _ = "end of CoverTab[14027]"
//line /usr/local/go/src/net/hosts.go:89
		_go_fuzz_dep_.CoverTab[14028]++

							var canonical string
							for i := 1; i < len(f); i++ {
//line /usr/local/go/src/net/hosts.go:92
			_go_fuzz_dep_.CoverTab[14035]++
								name := absDomainName(f[i])
								h := []byte(f[i])
								lowerASCIIBytes(h)
								key := absDomainName(string(h))

								if i == 1 {
//line /usr/local/go/src/net/hosts.go:98
				_go_fuzz_dep_.CoverTab[14038]++
									canonical = key
//line /usr/local/go/src/net/hosts.go:99
				// _ = "end of CoverTab[14038]"
			} else {
//line /usr/local/go/src/net/hosts.go:100
				_go_fuzz_dep_.CoverTab[14039]++
//line /usr/local/go/src/net/hosts.go:100
				// _ = "end of CoverTab[14039]"
//line /usr/local/go/src/net/hosts.go:100
			}
//line /usr/local/go/src/net/hosts.go:100
			// _ = "end of CoverTab[14035]"
//line /usr/local/go/src/net/hosts.go:100
			_go_fuzz_dep_.CoverTab[14036]++

								is[addr] = append(is[addr], name)

								if v, ok := hs[key]; ok {
//line /usr/local/go/src/net/hosts.go:104
				_go_fuzz_dep_.CoverTab[14040]++
									hs[key] = byName{
					addrs:		append(v.addrs, addr),
					canonicalName:	v.canonicalName,
				}
									continue
//line /usr/local/go/src/net/hosts.go:109
				// _ = "end of CoverTab[14040]"
			} else {
//line /usr/local/go/src/net/hosts.go:110
				_go_fuzz_dep_.CoverTab[14041]++
//line /usr/local/go/src/net/hosts.go:110
				// _ = "end of CoverTab[14041]"
//line /usr/local/go/src/net/hosts.go:110
			}
//line /usr/local/go/src/net/hosts.go:110
			// _ = "end of CoverTab[14036]"
//line /usr/local/go/src/net/hosts.go:110
			_go_fuzz_dep_.CoverTab[14037]++

								hs[key] = byName{
				addrs:		[]string{addr},
				canonicalName:	canonical,
			}
//line /usr/local/go/src/net/hosts.go:115
			// _ = "end of CoverTab[14037]"
		}
//line /usr/local/go/src/net/hosts.go:116
		// _ = "end of CoverTab[14028]"
	}
//line /usr/local/go/src/net/hosts.go:117
	// _ = "end of CoverTab[14012]"
//line /usr/local/go/src/net/hosts.go:117
	_go_fuzz_dep_.CoverTab[14013]++

						hosts.expire = now.Add(cacheMaxAge)
						hosts.path = hp
						hosts.byName = hs
						hosts.byAddr = is
						hosts.mtime = mtime
						hosts.size = size
						file.close()
//line /usr/local/go/src/net/hosts.go:125
	// _ = "end of CoverTab[14013]"
}

// lookupStaticHost looks up the addresses and the canonical name for the given host from /etc/hosts.
func lookupStaticHost(host string) ([]string, string) {
//line /usr/local/go/src/net/hosts.go:129
	_go_fuzz_dep_.CoverTab[14042]++
						hosts.Lock()
						defer hosts.Unlock()
						readHosts()
						if len(hosts.byName) != 0 {
//line /usr/local/go/src/net/hosts.go:133
		_go_fuzz_dep_.CoverTab[14044]++
							if hasUpperCase(host) {
//line /usr/local/go/src/net/hosts.go:134
			_go_fuzz_dep_.CoverTab[14046]++
								lowerHost := []byte(host)
								lowerASCIIBytes(lowerHost)
								host = string(lowerHost)
//line /usr/local/go/src/net/hosts.go:137
			// _ = "end of CoverTab[14046]"
		} else {
//line /usr/local/go/src/net/hosts.go:138
			_go_fuzz_dep_.CoverTab[14047]++
//line /usr/local/go/src/net/hosts.go:138
			// _ = "end of CoverTab[14047]"
//line /usr/local/go/src/net/hosts.go:138
		}
//line /usr/local/go/src/net/hosts.go:138
		// _ = "end of CoverTab[14044]"
//line /usr/local/go/src/net/hosts.go:138
		_go_fuzz_dep_.CoverTab[14045]++
							if byName, ok := hosts.byName[absDomainName(host)]; ok {
//line /usr/local/go/src/net/hosts.go:139
			_go_fuzz_dep_.CoverTab[14048]++
								ipsCp := make([]string, len(byName.addrs))
								copy(ipsCp, byName.addrs)
								return ipsCp, byName.canonicalName
//line /usr/local/go/src/net/hosts.go:142
			// _ = "end of CoverTab[14048]"
		} else {
//line /usr/local/go/src/net/hosts.go:143
			_go_fuzz_dep_.CoverTab[14049]++
//line /usr/local/go/src/net/hosts.go:143
			// _ = "end of CoverTab[14049]"
//line /usr/local/go/src/net/hosts.go:143
		}
//line /usr/local/go/src/net/hosts.go:143
		// _ = "end of CoverTab[14045]"
	} else {
//line /usr/local/go/src/net/hosts.go:144
		_go_fuzz_dep_.CoverTab[14050]++
//line /usr/local/go/src/net/hosts.go:144
		// _ = "end of CoverTab[14050]"
//line /usr/local/go/src/net/hosts.go:144
	}
//line /usr/local/go/src/net/hosts.go:144
	// _ = "end of CoverTab[14042]"
//line /usr/local/go/src/net/hosts.go:144
	_go_fuzz_dep_.CoverTab[14043]++
						return nil, ""
//line /usr/local/go/src/net/hosts.go:145
	// _ = "end of CoverTab[14043]"
}

// lookupStaticAddr looks up the hosts for the given address from /etc/hosts.
func lookupStaticAddr(addr string) []string {
//line /usr/local/go/src/net/hosts.go:149
	_go_fuzz_dep_.CoverTab[14051]++
						hosts.Lock()
						defer hosts.Unlock()
						readHosts()
						addr = parseLiteralIP(addr)
						if addr == "" {
//line /usr/local/go/src/net/hosts.go:154
		_go_fuzz_dep_.CoverTab[14054]++
							return nil
//line /usr/local/go/src/net/hosts.go:155
		// _ = "end of CoverTab[14054]"
	} else {
//line /usr/local/go/src/net/hosts.go:156
		_go_fuzz_dep_.CoverTab[14055]++
//line /usr/local/go/src/net/hosts.go:156
		// _ = "end of CoverTab[14055]"
//line /usr/local/go/src/net/hosts.go:156
	}
//line /usr/local/go/src/net/hosts.go:156
	// _ = "end of CoverTab[14051]"
//line /usr/local/go/src/net/hosts.go:156
	_go_fuzz_dep_.CoverTab[14052]++
						if len(hosts.byAddr) != 0 {
//line /usr/local/go/src/net/hosts.go:157
		_go_fuzz_dep_.CoverTab[14056]++
							if hosts, ok := hosts.byAddr[addr]; ok {
//line /usr/local/go/src/net/hosts.go:158
			_go_fuzz_dep_.CoverTab[14057]++
								hostsCp := make([]string, len(hosts))
								copy(hostsCp, hosts)
								return hostsCp
//line /usr/local/go/src/net/hosts.go:161
			// _ = "end of CoverTab[14057]"
		} else {
//line /usr/local/go/src/net/hosts.go:162
			_go_fuzz_dep_.CoverTab[14058]++
//line /usr/local/go/src/net/hosts.go:162
			// _ = "end of CoverTab[14058]"
//line /usr/local/go/src/net/hosts.go:162
		}
//line /usr/local/go/src/net/hosts.go:162
		// _ = "end of CoverTab[14056]"
	} else {
//line /usr/local/go/src/net/hosts.go:163
		_go_fuzz_dep_.CoverTab[14059]++
//line /usr/local/go/src/net/hosts.go:163
		// _ = "end of CoverTab[14059]"
//line /usr/local/go/src/net/hosts.go:163
	}
//line /usr/local/go/src/net/hosts.go:163
	// _ = "end of CoverTab[14052]"
//line /usr/local/go/src/net/hosts.go:163
	_go_fuzz_dep_.CoverTab[14053]++
						return nil
//line /usr/local/go/src/net/hosts.go:164
	// _ = "end of CoverTab[14053]"
}

//line /usr/local/go/src/net/hosts.go:165
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/hosts.go:165
var _ = _go_fuzz_dep_.CoverTab
