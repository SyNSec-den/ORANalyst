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
	_go_fuzz_dep_.CoverTab[5609]++
						var ip IP
						var zone string
						ip = parseIPv4(addr)
						if ip == nil {
//line /usr/local/go/src/net/hosts.go:19
		_go_fuzz_dep_.CoverTab[5613]++
							ip, zone = parseIPv6Zone(addr)
//line /usr/local/go/src/net/hosts.go:20
		// _ = "end of CoverTab[5613]"
	} else {
//line /usr/local/go/src/net/hosts.go:21
		_go_fuzz_dep_.CoverTab[5614]++
//line /usr/local/go/src/net/hosts.go:21
		// _ = "end of CoverTab[5614]"
//line /usr/local/go/src/net/hosts.go:21
	}
//line /usr/local/go/src/net/hosts.go:21
	// _ = "end of CoverTab[5609]"
//line /usr/local/go/src/net/hosts.go:21
	_go_fuzz_dep_.CoverTab[5610]++
						if ip == nil {
//line /usr/local/go/src/net/hosts.go:22
		_go_fuzz_dep_.CoverTab[5615]++
							return ""
//line /usr/local/go/src/net/hosts.go:23
		// _ = "end of CoverTab[5615]"
	} else {
//line /usr/local/go/src/net/hosts.go:24
		_go_fuzz_dep_.CoverTab[5616]++
//line /usr/local/go/src/net/hosts.go:24
		// _ = "end of CoverTab[5616]"
//line /usr/local/go/src/net/hosts.go:24
	}
//line /usr/local/go/src/net/hosts.go:24
	// _ = "end of CoverTab[5610]"
//line /usr/local/go/src/net/hosts.go:24
	_go_fuzz_dep_.CoverTab[5611]++
						if zone == "" {
//line /usr/local/go/src/net/hosts.go:25
		_go_fuzz_dep_.CoverTab[5617]++
							return ip.String()
//line /usr/local/go/src/net/hosts.go:26
		// _ = "end of CoverTab[5617]"
	} else {
//line /usr/local/go/src/net/hosts.go:27
		_go_fuzz_dep_.CoverTab[5618]++
//line /usr/local/go/src/net/hosts.go:27
		// _ = "end of CoverTab[5618]"
//line /usr/local/go/src/net/hosts.go:27
	}
//line /usr/local/go/src/net/hosts.go:27
	// _ = "end of CoverTab[5611]"
//line /usr/local/go/src/net/hosts.go:27
	_go_fuzz_dep_.CoverTab[5612]++
						return ip.String() + "%" + zone
//line /usr/local/go/src/net/hosts.go:28
	// _ = "end of CoverTab[5612]"
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
	_go_fuzz_dep_.CoverTab[5619]++
						now := time.Now()
						hp := testHookHostsPath

						if now.Before(hosts.expire) && func() bool {
//line /usr/local/go/src/net/hosts.go:61
		_go_fuzz_dep_.CoverTab[5624]++
//line /usr/local/go/src/net/hosts.go:61
		return hosts.path == hp
//line /usr/local/go/src/net/hosts.go:61
		// _ = "end of CoverTab[5624]"
//line /usr/local/go/src/net/hosts.go:61
	}() && func() bool {
//line /usr/local/go/src/net/hosts.go:61
		_go_fuzz_dep_.CoverTab[5625]++
//line /usr/local/go/src/net/hosts.go:61
		return len(hosts.byName) > 0
//line /usr/local/go/src/net/hosts.go:61
		// _ = "end of CoverTab[5625]"
//line /usr/local/go/src/net/hosts.go:61
	}() {
//line /usr/local/go/src/net/hosts.go:61
		_go_fuzz_dep_.CoverTab[5626]++
							return
//line /usr/local/go/src/net/hosts.go:62
		// _ = "end of CoverTab[5626]"
	} else {
//line /usr/local/go/src/net/hosts.go:63
		_go_fuzz_dep_.CoverTab[5627]++
//line /usr/local/go/src/net/hosts.go:63
		// _ = "end of CoverTab[5627]"
//line /usr/local/go/src/net/hosts.go:63
	}
//line /usr/local/go/src/net/hosts.go:63
	// _ = "end of CoverTab[5619]"
//line /usr/local/go/src/net/hosts.go:63
	_go_fuzz_dep_.CoverTab[5620]++
						mtime, size, err := stat(hp)
						if err == nil && func() bool {
//line /usr/local/go/src/net/hosts.go:65
		_go_fuzz_dep_.CoverTab[5628]++
//line /usr/local/go/src/net/hosts.go:65
		return hosts.path == hp
//line /usr/local/go/src/net/hosts.go:65
		// _ = "end of CoverTab[5628]"
//line /usr/local/go/src/net/hosts.go:65
	}() && func() bool {
//line /usr/local/go/src/net/hosts.go:65
		_go_fuzz_dep_.CoverTab[5629]++
//line /usr/local/go/src/net/hosts.go:65
		return hosts.mtime.Equal(mtime)
//line /usr/local/go/src/net/hosts.go:65
		// _ = "end of CoverTab[5629]"
//line /usr/local/go/src/net/hosts.go:65
	}() && func() bool {
//line /usr/local/go/src/net/hosts.go:65
		_go_fuzz_dep_.CoverTab[5630]++
//line /usr/local/go/src/net/hosts.go:65
		return hosts.size == size
//line /usr/local/go/src/net/hosts.go:65
		// _ = "end of CoverTab[5630]"
//line /usr/local/go/src/net/hosts.go:65
	}() {
//line /usr/local/go/src/net/hosts.go:65
		_go_fuzz_dep_.CoverTab[5631]++
							hosts.expire = now.Add(cacheMaxAge)
							return
//line /usr/local/go/src/net/hosts.go:67
		// _ = "end of CoverTab[5631]"
	} else {
//line /usr/local/go/src/net/hosts.go:68
		_go_fuzz_dep_.CoverTab[5632]++
//line /usr/local/go/src/net/hosts.go:68
		// _ = "end of CoverTab[5632]"
//line /usr/local/go/src/net/hosts.go:68
	}
//line /usr/local/go/src/net/hosts.go:68
	// _ = "end of CoverTab[5620]"
//line /usr/local/go/src/net/hosts.go:68
	_go_fuzz_dep_.CoverTab[5621]++

						hs := make(map[string]byName)
						is := make(map[string][]string)

						var file *file
						if file, _ = open(hp); file == nil {
//line /usr/local/go/src/net/hosts.go:74
		_go_fuzz_dep_.CoverTab[5633]++
							return
//line /usr/local/go/src/net/hosts.go:75
		// _ = "end of CoverTab[5633]"
	} else {
//line /usr/local/go/src/net/hosts.go:76
		_go_fuzz_dep_.CoverTab[5634]++
//line /usr/local/go/src/net/hosts.go:76
		// _ = "end of CoverTab[5634]"
//line /usr/local/go/src/net/hosts.go:76
	}
//line /usr/local/go/src/net/hosts.go:76
	// _ = "end of CoverTab[5621]"
//line /usr/local/go/src/net/hosts.go:76
	_go_fuzz_dep_.CoverTab[5622]++
						for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /usr/local/go/src/net/hosts.go:77
		_go_fuzz_dep_.CoverTab[5635]++
							if i := bytealg.IndexByteString(line, '#'); i >= 0 {
//line /usr/local/go/src/net/hosts.go:78
			_go_fuzz_dep_.CoverTab[5639]++

								line = line[0:i]
//line /usr/local/go/src/net/hosts.go:80
			// _ = "end of CoverTab[5639]"
		} else {
//line /usr/local/go/src/net/hosts.go:81
			_go_fuzz_dep_.CoverTab[5640]++
//line /usr/local/go/src/net/hosts.go:81
			// _ = "end of CoverTab[5640]"
//line /usr/local/go/src/net/hosts.go:81
		}
//line /usr/local/go/src/net/hosts.go:81
		// _ = "end of CoverTab[5635]"
//line /usr/local/go/src/net/hosts.go:81
		_go_fuzz_dep_.CoverTab[5636]++
							f := getFields(line)
							if len(f) < 2 {
//line /usr/local/go/src/net/hosts.go:83
			_go_fuzz_dep_.CoverTab[5641]++
								continue
//line /usr/local/go/src/net/hosts.go:84
			// _ = "end of CoverTab[5641]"
		} else {
//line /usr/local/go/src/net/hosts.go:85
			_go_fuzz_dep_.CoverTab[5642]++
//line /usr/local/go/src/net/hosts.go:85
			// _ = "end of CoverTab[5642]"
//line /usr/local/go/src/net/hosts.go:85
		}
//line /usr/local/go/src/net/hosts.go:85
		// _ = "end of CoverTab[5636]"
//line /usr/local/go/src/net/hosts.go:85
		_go_fuzz_dep_.CoverTab[5637]++
							addr := parseLiteralIP(f[0])
							if addr == "" {
//line /usr/local/go/src/net/hosts.go:87
			_go_fuzz_dep_.CoverTab[5643]++
								continue
//line /usr/local/go/src/net/hosts.go:88
			// _ = "end of CoverTab[5643]"
		} else {
//line /usr/local/go/src/net/hosts.go:89
			_go_fuzz_dep_.CoverTab[5644]++
//line /usr/local/go/src/net/hosts.go:89
			// _ = "end of CoverTab[5644]"
//line /usr/local/go/src/net/hosts.go:89
		}
//line /usr/local/go/src/net/hosts.go:89
		// _ = "end of CoverTab[5637]"
//line /usr/local/go/src/net/hosts.go:89
		_go_fuzz_dep_.CoverTab[5638]++

							var canonical string
							for i := 1; i < len(f); i++ {
//line /usr/local/go/src/net/hosts.go:92
			_go_fuzz_dep_.CoverTab[5645]++
								name := absDomainName(f[i])
								h := []byte(f[i])
								lowerASCIIBytes(h)
								key := absDomainName(string(h))

								if i == 1 {
//line /usr/local/go/src/net/hosts.go:98
				_go_fuzz_dep_.CoverTab[5648]++
									canonical = key
//line /usr/local/go/src/net/hosts.go:99
				// _ = "end of CoverTab[5648]"
			} else {
//line /usr/local/go/src/net/hosts.go:100
				_go_fuzz_dep_.CoverTab[5649]++
//line /usr/local/go/src/net/hosts.go:100
				// _ = "end of CoverTab[5649]"
//line /usr/local/go/src/net/hosts.go:100
			}
//line /usr/local/go/src/net/hosts.go:100
			// _ = "end of CoverTab[5645]"
//line /usr/local/go/src/net/hosts.go:100
			_go_fuzz_dep_.CoverTab[5646]++

								is[addr] = append(is[addr], name)

								if v, ok := hs[key]; ok {
//line /usr/local/go/src/net/hosts.go:104
				_go_fuzz_dep_.CoverTab[5650]++
									hs[key] = byName{
					addrs:		append(v.addrs, addr),
					canonicalName:	v.canonicalName,
				}
									continue
//line /usr/local/go/src/net/hosts.go:109
				// _ = "end of CoverTab[5650]"
			} else {
//line /usr/local/go/src/net/hosts.go:110
				_go_fuzz_dep_.CoverTab[5651]++
//line /usr/local/go/src/net/hosts.go:110
				// _ = "end of CoverTab[5651]"
//line /usr/local/go/src/net/hosts.go:110
			}
//line /usr/local/go/src/net/hosts.go:110
			// _ = "end of CoverTab[5646]"
//line /usr/local/go/src/net/hosts.go:110
			_go_fuzz_dep_.CoverTab[5647]++

								hs[key] = byName{
				addrs:		[]string{addr},
				canonicalName:	canonical,
			}
//line /usr/local/go/src/net/hosts.go:115
			// _ = "end of CoverTab[5647]"
		}
//line /usr/local/go/src/net/hosts.go:116
		// _ = "end of CoverTab[5638]"
	}
//line /usr/local/go/src/net/hosts.go:117
	// _ = "end of CoverTab[5622]"
//line /usr/local/go/src/net/hosts.go:117
	_go_fuzz_dep_.CoverTab[5623]++

						hosts.expire = now.Add(cacheMaxAge)
						hosts.path = hp
						hosts.byName = hs
						hosts.byAddr = is
						hosts.mtime = mtime
						hosts.size = size
						file.close()
//line /usr/local/go/src/net/hosts.go:125
	// _ = "end of CoverTab[5623]"
}

// lookupStaticHost looks up the addresses and the canonical name for the given host from /etc/hosts.
func lookupStaticHost(host string) ([]string, string) {
//line /usr/local/go/src/net/hosts.go:129
	_go_fuzz_dep_.CoverTab[5652]++
						hosts.Lock()
						defer hosts.Unlock()
						readHosts()
						if len(hosts.byName) != 0 {
//line /usr/local/go/src/net/hosts.go:133
		_go_fuzz_dep_.CoverTab[5654]++
							if hasUpperCase(host) {
//line /usr/local/go/src/net/hosts.go:134
			_go_fuzz_dep_.CoverTab[5656]++
								lowerHost := []byte(host)
								lowerASCIIBytes(lowerHost)
								host = string(lowerHost)
//line /usr/local/go/src/net/hosts.go:137
			// _ = "end of CoverTab[5656]"
		} else {
//line /usr/local/go/src/net/hosts.go:138
			_go_fuzz_dep_.CoverTab[5657]++
//line /usr/local/go/src/net/hosts.go:138
			// _ = "end of CoverTab[5657]"
//line /usr/local/go/src/net/hosts.go:138
		}
//line /usr/local/go/src/net/hosts.go:138
		// _ = "end of CoverTab[5654]"
//line /usr/local/go/src/net/hosts.go:138
		_go_fuzz_dep_.CoverTab[5655]++
							if byName, ok := hosts.byName[absDomainName(host)]; ok {
//line /usr/local/go/src/net/hosts.go:139
			_go_fuzz_dep_.CoverTab[5658]++
								ipsCp := make([]string, len(byName.addrs))
								copy(ipsCp, byName.addrs)
								return ipsCp, byName.canonicalName
//line /usr/local/go/src/net/hosts.go:142
			// _ = "end of CoverTab[5658]"
		} else {
//line /usr/local/go/src/net/hosts.go:143
			_go_fuzz_dep_.CoverTab[5659]++
//line /usr/local/go/src/net/hosts.go:143
			// _ = "end of CoverTab[5659]"
//line /usr/local/go/src/net/hosts.go:143
		}
//line /usr/local/go/src/net/hosts.go:143
		// _ = "end of CoverTab[5655]"
	} else {
//line /usr/local/go/src/net/hosts.go:144
		_go_fuzz_dep_.CoverTab[5660]++
//line /usr/local/go/src/net/hosts.go:144
		// _ = "end of CoverTab[5660]"
//line /usr/local/go/src/net/hosts.go:144
	}
//line /usr/local/go/src/net/hosts.go:144
	// _ = "end of CoverTab[5652]"
//line /usr/local/go/src/net/hosts.go:144
	_go_fuzz_dep_.CoverTab[5653]++
						return nil, ""
//line /usr/local/go/src/net/hosts.go:145
	// _ = "end of CoverTab[5653]"
}

// lookupStaticAddr looks up the hosts for the given address from /etc/hosts.
func lookupStaticAddr(addr string) []string {
//line /usr/local/go/src/net/hosts.go:149
	_go_fuzz_dep_.CoverTab[5661]++
						hosts.Lock()
						defer hosts.Unlock()
						readHosts()
						addr = parseLiteralIP(addr)
						if addr == "" {
//line /usr/local/go/src/net/hosts.go:154
		_go_fuzz_dep_.CoverTab[5664]++
							return nil
//line /usr/local/go/src/net/hosts.go:155
		// _ = "end of CoverTab[5664]"
	} else {
//line /usr/local/go/src/net/hosts.go:156
		_go_fuzz_dep_.CoverTab[5665]++
//line /usr/local/go/src/net/hosts.go:156
		// _ = "end of CoverTab[5665]"
//line /usr/local/go/src/net/hosts.go:156
	}
//line /usr/local/go/src/net/hosts.go:156
	// _ = "end of CoverTab[5661]"
//line /usr/local/go/src/net/hosts.go:156
	_go_fuzz_dep_.CoverTab[5662]++
						if len(hosts.byAddr) != 0 {
//line /usr/local/go/src/net/hosts.go:157
		_go_fuzz_dep_.CoverTab[5666]++
							if hosts, ok := hosts.byAddr[addr]; ok {
//line /usr/local/go/src/net/hosts.go:158
			_go_fuzz_dep_.CoverTab[5667]++
								hostsCp := make([]string, len(hosts))
								copy(hostsCp, hosts)
								return hostsCp
//line /usr/local/go/src/net/hosts.go:161
			// _ = "end of CoverTab[5667]"
		} else {
//line /usr/local/go/src/net/hosts.go:162
			_go_fuzz_dep_.CoverTab[5668]++
//line /usr/local/go/src/net/hosts.go:162
			// _ = "end of CoverTab[5668]"
//line /usr/local/go/src/net/hosts.go:162
		}
//line /usr/local/go/src/net/hosts.go:162
		// _ = "end of CoverTab[5666]"
	} else {
//line /usr/local/go/src/net/hosts.go:163
		_go_fuzz_dep_.CoverTab[5669]++
//line /usr/local/go/src/net/hosts.go:163
		// _ = "end of CoverTab[5669]"
//line /usr/local/go/src/net/hosts.go:163
	}
//line /usr/local/go/src/net/hosts.go:163
	// _ = "end of CoverTab[5662]"
//line /usr/local/go/src/net/hosts.go:163
	_go_fuzz_dep_.CoverTab[5663]++
						return nil
//line /usr/local/go/src/net/hosts.go:164
	// _ = "end of CoverTab[5663]"
}

//line /usr/local/go/src/net/hosts.go:165
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/hosts.go:165
var _ = _go_fuzz_dep_.CoverTab
