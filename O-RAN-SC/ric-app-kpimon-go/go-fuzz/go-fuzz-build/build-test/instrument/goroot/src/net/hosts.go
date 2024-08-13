// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/hosts.go:5
package net

//line /snap/go/10455/src/net/hosts.go:5
import (
//line /snap/go/10455/src/net/hosts.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/hosts.go:5
)
//line /snap/go/10455/src/net/hosts.go:5
import (
//line /snap/go/10455/src/net/hosts.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/hosts.go:5
)

import (
	"errors"
	"internal/bytealg"
	"io/fs"
	"net/netip"
	"sync"
	"time"
)

const cacheMaxAge = 5 * time.Second

func parseLiteralIP(addr string) string {
//line /snap/go/10455/src/net/hosts.go:18
	_go_fuzz_dep_.CoverTab[5985]++
						ip, err := netip.ParseAddr(addr)
						if err != nil {
//line /snap/go/10455/src/net/hosts.go:20
		_go_fuzz_dep_.CoverTab[528330]++
//line /snap/go/10455/src/net/hosts.go:20
		_go_fuzz_dep_.CoverTab[5987]++
							return ""
//line /snap/go/10455/src/net/hosts.go:21
		// _ = "end of CoverTab[5987]"
	} else {
//line /snap/go/10455/src/net/hosts.go:22
		_go_fuzz_dep_.CoverTab[528331]++
//line /snap/go/10455/src/net/hosts.go:22
		_go_fuzz_dep_.CoverTab[5988]++
//line /snap/go/10455/src/net/hosts.go:22
		// _ = "end of CoverTab[5988]"
//line /snap/go/10455/src/net/hosts.go:22
	}
//line /snap/go/10455/src/net/hosts.go:22
	// _ = "end of CoverTab[5985]"
//line /snap/go/10455/src/net/hosts.go:22
	_go_fuzz_dep_.CoverTab[5986]++
						return ip.String()
//line /snap/go/10455/src/net/hosts.go:23
	// _ = "end of CoverTab[5986]"
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
//line /snap/go/10455/src/net/hosts.go:52
	_go_fuzz_dep_.CoverTab[5989]++
						now := time.Now()
						hp := testHookHostsPath

						if now.Before(hosts.expire) && func() bool {
//line /snap/go/10455/src/net/hosts.go:56
		_go_fuzz_dep_.CoverTab[5994]++
//line /snap/go/10455/src/net/hosts.go:56
		return hosts.path == hp
//line /snap/go/10455/src/net/hosts.go:56
		// _ = "end of CoverTab[5994]"
//line /snap/go/10455/src/net/hosts.go:56
	}() && func() bool {
//line /snap/go/10455/src/net/hosts.go:56
		_go_fuzz_dep_.CoverTab[5995]++
//line /snap/go/10455/src/net/hosts.go:56
		return len(hosts.byName) > 0
//line /snap/go/10455/src/net/hosts.go:56
		// _ = "end of CoverTab[5995]"
//line /snap/go/10455/src/net/hosts.go:56
	}() {
//line /snap/go/10455/src/net/hosts.go:56
		_go_fuzz_dep_.CoverTab[528332]++
//line /snap/go/10455/src/net/hosts.go:56
		_go_fuzz_dep_.CoverTab[5996]++
							return
//line /snap/go/10455/src/net/hosts.go:57
		// _ = "end of CoverTab[5996]"
	} else {
//line /snap/go/10455/src/net/hosts.go:58
		_go_fuzz_dep_.CoverTab[528333]++
//line /snap/go/10455/src/net/hosts.go:58
		_go_fuzz_dep_.CoverTab[5997]++
//line /snap/go/10455/src/net/hosts.go:58
		// _ = "end of CoverTab[5997]"
//line /snap/go/10455/src/net/hosts.go:58
	}
//line /snap/go/10455/src/net/hosts.go:58
	// _ = "end of CoverTab[5989]"
//line /snap/go/10455/src/net/hosts.go:58
	_go_fuzz_dep_.CoverTab[5990]++
						mtime, size, err := stat(hp)
						if err == nil && func() bool {
//line /snap/go/10455/src/net/hosts.go:60
		_go_fuzz_dep_.CoverTab[5998]++
//line /snap/go/10455/src/net/hosts.go:60
		return hosts.path == hp
//line /snap/go/10455/src/net/hosts.go:60
		// _ = "end of CoverTab[5998]"
//line /snap/go/10455/src/net/hosts.go:60
	}() && func() bool {
//line /snap/go/10455/src/net/hosts.go:60
		_go_fuzz_dep_.CoverTab[5999]++
//line /snap/go/10455/src/net/hosts.go:60
		return hosts.mtime.Equal(mtime)
//line /snap/go/10455/src/net/hosts.go:60
		// _ = "end of CoverTab[5999]"
//line /snap/go/10455/src/net/hosts.go:60
	}() && func() bool {
//line /snap/go/10455/src/net/hosts.go:60
		_go_fuzz_dep_.CoverTab[6000]++
//line /snap/go/10455/src/net/hosts.go:60
		return hosts.size == size
//line /snap/go/10455/src/net/hosts.go:60
		// _ = "end of CoverTab[6000]"
//line /snap/go/10455/src/net/hosts.go:60
	}() {
//line /snap/go/10455/src/net/hosts.go:60
		_go_fuzz_dep_.CoverTab[528334]++
//line /snap/go/10455/src/net/hosts.go:60
		_go_fuzz_dep_.CoverTab[6001]++
							hosts.expire = now.Add(cacheMaxAge)
							return
//line /snap/go/10455/src/net/hosts.go:62
		// _ = "end of CoverTab[6001]"
	} else {
//line /snap/go/10455/src/net/hosts.go:63
		_go_fuzz_dep_.CoverTab[528335]++
//line /snap/go/10455/src/net/hosts.go:63
		_go_fuzz_dep_.CoverTab[6002]++
//line /snap/go/10455/src/net/hosts.go:63
		// _ = "end of CoverTab[6002]"
//line /snap/go/10455/src/net/hosts.go:63
	}
//line /snap/go/10455/src/net/hosts.go:63
	// _ = "end of CoverTab[5990]"
//line /snap/go/10455/src/net/hosts.go:63
	_go_fuzz_dep_.CoverTab[5991]++

						hs := make(map[string]byName)
						is := make(map[string][]string)

						file, err := open(hp)
						if err != nil {
//line /snap/go/10455/src/net/hosts.go:69
		_go_fuzz_dep_.CoverTab[528336]++
//line /snap/go/10455/src/net/hosts.go:69
		_go_fuzz_dep_.CoverTab[6003]++
							if !errors.Is(err, fs.ErrNotExist) && func() bool {
//line /snap/go/10455/src/net/hosts.go:70
			_go_fuzz_dep_.CoverTab[6004]++
//line /snap/go/10455/src/net/hosts.go:70
			return !errors.Is(err, fs.ErrPermission)
//line /snap/go/10455/src/net/hosts.go:70
			// _ = "end of CoverTab[6004]"
//line /snap/go/10455/src/net/hosts.go:70
		}() {
//line /snap/go/10455/src/net/hosts.go:70
			_go_fuzz_dep_.CoverTab[528338]++
//line /snap/go/10455/src/net/hosts.go:70
			_go_fuzz_dep_.CoverTab[6005]++
								return
//line /snap/go/10455/src/net/hosts.go:71
			// _ = "end of CoverTab[6005]"
		} else {
//line /snap/go/10455/src/net/hosts.go:72
			_go_fuzz_dep_.CoverTab[528339]++
//line /snap/go/10455/src/net/hosts.go:72
			_go_fuzz_dep_.CoverTab[6006]++
//line /snap/go/10455/src/net/hosts.go:72
			// _ = "end of CoverTab[6006]"
//line /snap/go/10455/src/net/hosts.go:72
		}
//line /snap/go/10455/src/net/hosts.go:72
		// _ = "end of CoverTab[6003]"
	} else {
//line /snap/go/10455/src/net/hosts.go:73
		_go_fuzz_dep_.CoverTab[528337]++
//line /snap/go/10455/src/net/hosts.go:73
		_go_fuzz_dep_.CoverTab[6007]++
//line /snap/go/10455/src/net/hosts.go:73
		// _ = "end of CoverTab[6007]"
//line /snap/go/10455/src/net/hosts.go:73
	}
//line /snap/go/10455/src/net/hosts.go:73
	// _ = "end of CoverTab[5991]"
//line /snap/go/10455/src/net/hosts.go:73
	_go_fuzz_dep_.CoverTab[5992]++

						if file != nil {
//line /snap/go/10455/src/net/hosts.go:75
		_go_fuzz_dep_.CoverTab[528340]++
//line /snap/go/10455/src/net/hosts.go:75
		_go_fuzz_dep_.CoverTab[6008]++
							defer file.close()
//line /snap/go/10455/src/net/hosts.go:76
		_go_fuzz_dep_.CoverTab[786675] = 0
							for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /snap/go/10455/src/net/hosts.go:77
			if _go_fuzz_dep_.CoverTab[786675] == 0 {
//line /snap/go/10455/src/net/hosts.go:77
				_go_fuzz_dep_.CoverTab[528364]++
//line /snap/go/10455/src/net/hosts.go:77
			} else {
//line /snap/go/10455/src/net/hosts.go:77
				_go_fuzz_dep_.CoverTab[528365]++
//line /snap/go/10455/src/net/hosts.go:77
			}
//line /snap/go/10455/src/net/hosts.go:77
			_go_fuzz_dep_.CoverTab[786675] = 1
//line /snap/go/10455/src/net/hosts.go:77
			_go_fuzz_dep_.CoverTab[6009]++
								if i := bytealg.IndexByteString(line, '#'); i >= 0 {
//line /snap/go/10455/src/net/hosts.go:78
				_go_fuzz_dep_.CoverTab[528342]++
//line /snap/go/10455/src/net/hosts.go:78
				_go_fuzz_dep_.CoverTab[6013]++

									line = line[0:i]
//line /snap/go/10455/src/net/hosts.go:80
				// _ = "end of CoverTab[6013]"
			} else {
//line /snap/go/10455/src/net/hosts.go:81
				_go_fuzz_dep_.CoverTab[528343]++
//line /snap/go/10455/src/net/hosts.go:81
				_go_fuzz_dep_.CoverTab[6014]++
//line /snap/go/10455/src/net/hosts.go:81
				// _ = "end of CoverTab[6014]"
//line /snap/go/10455/src/net/hosts.go:81
			}
//line /snap/go/10455/src/net/hosts.go:81
			// _ = "end of CoverTab[6009]"
//line /snap/go/10455/src/net/hosts.go:81
			_go_fuzz_dep_.CoverTab[6010]++
								f := getFields(line)
								if len(f) < 2 {
//line /snap/go/10455/src/net/hosts.go:83
				_go_fuzz_dep_.CoverTab[528344]++
//line /snap/go/10455/src/net/hosts.go:83
				_go_fuzz_dep_.CoverTab[6015]++
									continue
//line /snap/go/10455/src/net/hosts.go:84
				// _ = "end of CoverTab[6015]"
			} else {
//line /snap/go/10455/src/net/hosts.go:85
				_go_fuzz_dep_.CoverTab[528345]++
//line /snap/go/10455/src/net/hosts.go:85
				_go_fuzz_dep_.CoverTab[6016]++
//line /snap/go/10455/src/net/hosts.go:85
				// _ = "end of CoverTab[6016]"
//line /snap/go/10455/src/net/hosts.go:85
			}
//line /snap/go/10455/src/net/hosts.go:85
			// _ = "end of CoverTab[6010]"
//line /snap/go/10455/src/net/hosts.go:85
			_go_fuzz_dep_.CoverTab[6011]++
								addr := parseLiteralIP(f[0])
								if addr == "" {
//line /snap/go/10455/src/net/hosts.go:87
				_go_fuzz_dep_.CoverTab[528346]++
//line /snap/go/10455/src/net/hosts.go:87
				_go_fuzz_dep_.CoverTab[6017]++
									continue
//line /snap/go/10455/src/net/hosts.go:88
				// _ = "end of CoverTab[6017]"
			} else {
//line /snap/go/10455/src/net/hosts.go:89
				_go_fuzz_dep_.CoverTab[528347]++
//line /snap/go/10455/src/net/hosts.go:89
				_go_fuzz_dep_.CoverTab[6018]++
//line /snap/go/10455/src/net/hosts.go:89
				// _ = "end of CoverTab[6018]"
//line /snap/go/10455/src/net/hosts.go:89
			}
//line /snap/go/10455/src/net/hosts.go:89
			// _ = "end of CoverTab[6011]"
//line /snap/go/10455/src/net/hosts.go:89
			_go_fuzz_dep_.CoverTab[6012]++

								var canonical string
//line /snap/go/10455/src/net/hosts.go:91
			_go_fuzz_dep_.CoverTab[786676] = 0
								for i := 1; i < len(f); i++ {
//line /snap/go/10455/src/net/hosts.go:92
				if _go_fuzz_dep_.CoverTab[786676] == 0 {
//line /snap/go/10455/src/net/hosts.go:92
					_go_fuzz_dep_.CoverTab[528368]++
//line /snap/go/10455/src/net/hosts.go:92
				} else {
//line /snap/go/10455/src/net/hosts.go:92
					_go_fuzz_dep_.CoverTab[528369]++
//line /snap/go/10455/src/net/hosts.go:92
				}
//line /snap/go/10455/src/net/hosts.go:92
				_go_fuzz_dep_.CoverTab[786676] = 1
//line /snap/go/10455/src/net/hosts.go:92
				_go_fuzz_dep_.CoverTab[6019]++
									name := absDomainName(f[i])
									h := []byte(f[i])
									lowerASCIIBytes(h)
									key := absDomainName(string(h))

									if i == 1 {
//line /snap/go/10455/src/net/hosts.go:98
					_go_fuzz_dep_.CoverTab[528348]++
//line /snap/go/10455/src/net/hosts.go:98
					_go_fuzz_dep_.CoverTab[6022]++
										canonical = key
//line /snap/go/10455/src/net/hosts.go:99
					// _ = "end of CoverTab[6022]"
				} else {
//line /snap/go/10455/src/net/hosts.go:100
					_go_fuzz_dep_.CoverTab[528349]++
//line /snap/go/10455/src/net/hosts.go:100
					_go_fuzz_dep_.CoverTab[6023]++
//line /snap/go/10455/src/net/hosts.go:100
					// _ = "end of CoverTab[6023]"
//line /snap/go/10455/src/net/hosts.go:100
				}
//line /snap/go/10455/src/net/hosts.go:100
				// _ = "end of CoverTab[6019]"
//line /snap/go/10455/src/net/hosts.go:100
				_go_fuzz_dep_.CoverTab[6020]++

									is[addr] = append(is[addr], name)

									if v, ok := hs[key]; ok {
//line /snap/go/10455/src/net/hosts.go:104
					_go_fuzz_dep_.CoverTab[528350]++
//line /snap/go/10455/src/net/hosts.go:104
					_go_fuzz_dep_.CoverTab[6024]++
										hs[key] = byName{
						addrs:		append(v.addrs, addr),
						canonicalName:	v.canonicalName,
					}
										continue
//line /snap/go/10455/src/net/hosts.go:109
					// _ = "end of CoverTab[6024]"
				} else {
//line /snap/go/10455/src/net/hosts.go:110
					_go_fuzz_dep_.CoverTab[528351]++
//line /snap/go/10455/src/net/hosts.go:110
					_go_fuzz_dep_.CoverTab[6025]++
//line /snap/go/10455/src/net/hosts.go:110
					// _ = "end of CoverTab[6025]"
//line /snap/go/10455/src/net/hosts.go:110
				}
//line /snap/go/10455/src/net/hosts.go:110
				// _ = "end of CoverTab[6020]"
//line /snap/go/10455/src/net/hosts.go:110
				_go_fuzz_dep_.CoverTab[6021]++

									hs[key] = byName{
					addrs:		[]string{addr},
					canonicalName:	canonical,
				}
//line /snap/go/10455/src/net/hosts.go:115
				// _ = "end of CoverTab[6021]"
			}
//line /snap/go/10455/src/net/hosts.go:116
			if _go_fuzz_dep_.CoverTab[786676] == 0 {
//line /snap/go/10455/src/net/hosts.go:116
				_go_fuzz_dep_.CoverTab[528370]++
//line /snap/go/10455/src/net/hosts.go:116
			} else {
//line /snap/go/10455/src/net/hosts.go:116
				_go_fuzz_dep_.CoverTab[528371]++
//line /snap/go/10455/src/net/hosts.go:116
			}
//line /snap/go/10455/src/net/hosts.go:116
			// _ = "end of CoverTab[6012]"
		}
//line /snap/go/10455/src/net/hosts.go:117
		if _go_fuzz_dep_.CoverTab[786675] == 0 {
//line /snap/go/10455/src/net/hosts.go:117
			_go_fuzz_dep_.CoverTab[528366]++
//line /snap/go/10455/src/net/hosts.go:117
		} else {
//line /snap/go/10455/src/net/hosts.go:117
			_go_fuzz_dep_.CoverTab[528367]++
//line /snap/go/10455/src/net/hosts.go:117
		}
//line /snap/go/10455/src/net/hosts.go:117
		// _ = "end of CoverTab[6008]"
	} else {
//line /snap/go/10455/src/net/hosts.go:118
		_go_fuzz_dep_.CoverTab[528341]++
//line /snap/go/10455/src/net/hosts.go:118
		_go_fuzz_dep_.CoverTab[6026]++
//line /snap/go/10455/src/net/hosts.go:118
		// _ = "end of CoverTab[6026]"
//line /snap/go/10455/src/net/hosts.go:118
	}
//line /snap/go/10455/src/net/hosts.go:118
	// _ = "end of CoverTab[5992]"
//line /snap/go/10455/src/net/hosts.go:118
	_go_fuzz_dep_.CoverTab[5993]++

						hosts.expire = now.Add(cacheMaxAge)
						hosts.path = hp
						hosts.byName = hs
						hosts.byAddr = is
						hosts.mtime = mtime
						hosts.size = size
//line /snap/go/10455/src/net/hosts.go:125
	// _ = "end of CoverTab[5993]"
}

// lookupStaticHost looks up the addresses and the canonical name for the given host from /etc/hosts.
func lookupStaticHost(host string) ([]string, string) {
//line /snap/go/10455/src/net/hosts.go:129
	_go_fuzz_dep_.CoverTab[6027]++
						hosts.Lock()
						defer hosts.Unlock()
						readHosts()
						if len(hosts.byName) != 0 {
//line /snap/go/10455/src/net/hosts.go:133
		_go_fuzz_dep_.CoverTab[528352]++
//line /snap/go/10455/src/net/hosts.go:133
		_go_fuzz_dep_.CoverTab[6029]++
							if hasUpperCase(host) {
//line /snap/go/10455/src/net/hosts.go:134
			_go_fuzz_dep_.CoverTab[528354]++
//line /snap/go/10455/src/net/hosts.go:134
			_go_fuzz_dep_.CoverTab[6031]++
								lowerHost := []byte(host)
								lowerASCIIBytes(lowerHost)
								host = string(lowerHost)
//line /snap/go/10455/src/net/hosts.go:137
			// _ = "end of CoverTab[6031]"
		} else {
//line /snap/go/10455/src/net/hosts.go:138
			_go_fuzz_dep_.CoverTab[528355]++
//line /snap/go/10455/src/net/hosts.go:138
			_go_fuzz_dep_.CoverTab[6032]++
//line /snap/go/10455/src/net/hosts.go:138
			// _ = "end of CoverTab[6032]"
//line /snap/go/10455/src/net/hosts.go:138
		}
//line /snap/go/10455/src/net/hosts.go:138
		// _ = "end of CoverTab[6029]"
//line /snap/go/10455/src/net/hosts.go:138
		_go_fuzz_dep_.CoverTab[6030]++
							if byName, ok := hosts.byName[absDomainName(host)]; ok {
//line /snap/go/10455/src/net/hosts.go:139
			_go_fuzz_dep_.CoverTab[528356]++
//line /snap/go/10455/src/net/hosts.go:139
			_go_fuzz_dep_.CoverTab[6033]++
								ipsCp := make([]string, len(byName.addrs))
								copy(ipsCp, byName.addrs)
								return ipsCp, byName.canonicalName
//line /snap/go/10455/src/net/hosts.go:142
			// _ = "end of CoverTab[6033]"
		} else {
//line /snap/go/10455/src/net/hosts.go:143
			_go_fuzz_dep_.CoverTab[528357]++
//line /snap/go/10455/src/net/hosts.go:143
			_go_fuzz_dep_.CoverTab[6034]++
//line /snap/go/10455/src/net/hosts.go:143
			// _ = "end of CoverTab[6034]"
//line /snap/go/10455/src/net/hosts.go:143
		}
//line /snap/go/10455/src/net/hosts.go:143
		// _ = "end of CoverTab[6030]"
	} else {
//line /snap/go/10455/src/net/hosts.go:144
		_go_fuzz_dep_.CoverTab[528353]++
//line /snap/go/10455/src/net/hosts.go:144
		_go_fuzz_dep_.CoverTab[6035]++
//line /snap/go/10455/src/net/hosts.go:144
		// _ = "end of CoverTab[6035]"
//line /snap/go/10455/src/net/hosts.go:144
	}
//line /snap/go/10455/src/net/hosts.go:144
	// _ = "end of CoverTab[6027]"
//line /snap/go/10455/src/net/hosts.go:144
	_go_fuzz_dep_.CoverTab[6028]++
						return nil, ""
//line /snap/go/10455/src/net/hosts.go:145
	// _ = "end of CoverTab[6028]"
}

// lookupStaticAddr looks up the hosts for the given address from /etc/hosts.
func lookupStaticAddr(addr string) []string {
//line /snap/go/10455/src/net/hosts.go:149
	_go_fuzz_dep_.CoverTab[6036]++
						hosts.Lock()
						defer hosts.Unlock()
						readHosts()
						addr = parseLiteralIP(addr)
						if addr == "" {
//line /snap/go/10455/src/net/hosts.go:154
		_go_fuzz_dep_.CoverTab[528358]++
//line /snap/go/10455/src/net/hosts.go:154
		_go_fuzz_dep_.CoverTab[6039]++
							return nil
//line /snap/go/10455/src/net/hosts.go:155
		// _ = "end of CoverTab[6039]"
	} else {
//line /snap/go/10455/src/net/hosts.go:156
		_go_fuzz_dep_.CoverTab[528359]++
//line /snap/go/10455/src/net/hosts.go:156
		_go_fuzz_dep_.CoverTab[6040]++
//line /snap/go/10455/src/net/hosts.go:156
		// _ = "end of CoverTab[6040]"
//line /snap/go/10455/src/net/hosts.go:156
	}
//line /snap/go/10455/src/net/hosts.go:156
	// _ = "end of CoverTab[6036]"
//line /snap/go/10455/src/net/hosts.go:156
	_go_fuzz_dep_.CoverTab[6037]++
						if len(hosts.byAddr) != 0 {
//line /snap/go/10455/src/net/hosts.go:157
		_go_fuzz_dep_.CoverTab[528360]++
//line /snap/go/10455/src/net/hosts.go:157
		_go_fuzz_dep_.CoverTab[6041]++
							if hosts, ok := hosts.byAddr[addr]; ok {
//line /snap/go/10455/src/net/hosts.go:158
			_go_fuzz_dep_.CoverTab[528362]++
//line /snap/go/10455/src/net/hosts.go:158
			_go_fuzz_dep_.CoverTab[6042]++
								hostsCp := make([]string, len(hosts))
								copy(hostsCp, hosts)
								return hostsCp
//line /snap/go/10455/src/net/hosts.go:161
			// _ = "end of CoverTab[6042]"
		} else {
//line /snap/go/10455/src/net/hosts.go:162
			_go_fuzz_dep_.CoverTab[528363]++
//line /snap/go/10455/src/net/hosts.go:162
			_go_fuzz_dep_.CoverTab[6043]++
//line /snap/go/10455/src/net/hosts.go:162
			// _ = "end of CoverTab[6043]"
//line /snap/go/10455/src/net/hosts.go:162
		}
//line /snap/go/10455/src/net/hosts.go:162
		// _ = "end of CoverTab[6041]"
	} else {
//line /snap/go/10455/src/net/hosts.go:163
		_go_fuzz_dep_.CoverTab[528361]++
//line /snap/go/10455/src/net/hosts.go:163
		_go_fuzz_dep_.CoverTab[6044]++
//line /snap/go/10455/src/net/hosts.go:163
		// _ = "end of CoverTab[6044]"
//line /snap/go/10455/src/net/hosts.go:163
	}
//line /snap/go/10455/src/net/hosts.go:163
	// _ = "end of CoverTab[6037]"
//line /snap/go/10455/src/net/hosts.go:163
	_go_fuzz_dep_.CoverTab[6038]++
						return nil
//line /snap/go/10455/src/net/hosts.go:164
	// _ = "end of CoverTab[6038]"
}

//line /snap/go/10455/src/net/hosts.go:165
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/hosts.go:165
var _ = _go_fuzz_dep_.CoverTab
