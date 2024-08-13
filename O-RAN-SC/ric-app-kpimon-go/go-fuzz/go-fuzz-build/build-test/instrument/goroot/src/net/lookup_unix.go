// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || wasip1

//line /snap/go/10455/src/net/lookup_unix.go:7
package net

//line /snap/go/10455/src/net/lookup_unix.go:7
import (
//line /snap/go/10455/src/net/lookup_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/lookup_unix.go:7
)
//line /snap/go/10455/src/net/lookup_unix.go:7
import (
//line /snap/go/10455/src/net/lookup_unix.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/lookup_unix.go:7
)

import (
	"context"
	"internal/bytealg"
	"sync"
	"syscall"
)

var onceReadProtocols sync.Once

// readProtocols loads contents of /etc/protocols into protocols map
//line /snap/go/10455/src/net/lookup_unix.go:18
// for quick access.
//line /snap/go/10455/src/net/lookup_unix.go:20
func readProtocols() {
//line /snap/go/10455/src/net/lookup_unix.go:20
	_go_fuzz_dep_.CoverTab[7256]++
							file, err := open("/etc/protocols")
							if err != nil {
//line /snap/go/10455/src/net/lookup_unix.go:22
		_go_fuzz_dep_.CoverTab[529146]++
//line /snap/go/10455/src/net/lookup_unix.go:22
		_go_fuzz_dep_.CoverTab[7258]++
								return
//line /snap/go/10455/src/net/lookup_unix.go:23
		// _ = "end of CoverTab[7258]"
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:24
		_go_fuzz_dep_.CoverTab[529147]++
//line /snap/go/10455/src/net/lookup_unix.go:24
		_go_fuzz_dep_.CoverTab[7259]++
//line /snap/go/10455/src/net/lookup_unix.go:24
		// _ = "end of CoverTab[7259]"
//line /snap/go/10455/src/net/lookup_unix.go:24
	}
//line /snap/go/10455/src/net/lookup_unix.go:24
	// _ = "end of CoverTab[7256]"
//line /snap/go/10455/src/net/lookup_unix.go:24
	_go_fuzz_dep_.CoverTab[7257]++
							defer file.close()
//line /snap/go/10455/src/net/lookup_unix.go:25
	_go_fuzz_dep_.CoverTab[786715] = 0

							for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /snap/go/10455/src/net/lookup_unix.go:27
		if _go_fuzz_dep_.CoverTab[786715] == 0 {
//line /snap/go/10455/src/net/lookup_unix.go:27
			_go_fuzz_dep_.CoverTab[529182]++
//line /snap/go/10455/src/net/lookup_unix.go:27
		} else {
//line /snap/go/10455/src/net/lookup_unix.go:27
			_go_fuzz_dep_.CoverTab[529183]++
//line /snap/go/10455/src/net/lookup_unix.go:27
		}
//line /snap/go/10455/src/net/lookup_unix.go:27
		_go_fuzz_dep_.CoverTab[786715] = 1
//line /snap/go/10455/src/net/lookup_unix.go:27
		_go_fuzz_dep_.CoverTab[7260]++

								if i := bytealg.IndexByteString(line, '#'); i >= 0 {
//line /snap/go/10455/src/net/lookup_unix.go:29
			_go_fuzz_dep_.CoverTab[529148]++
//line /snap/go/10455/src/net/lookup_unix.go:29
			_go_fuzz_dep_.CoverTab[7263]++
									line = line[0:i]
//line /snap/go/10455/src/net/lookup_unix.go:30
			// _ = "end of CoverTab[7263]"
		} else {
//line /snap/go/10455/src/net/lookup_unix.go:31
			_go_fuzz_dep_.CoverTab[529149]++
//line /snap/go/10455/src/net/lookup_unix.go:31
			_go_fuzz_dep_.CoverTab[7264]++
//line /snap/go/10455/src/net/lookup_unix.go:31
			// _ = "end of CoverTab[7264]"
//line /snap/go/10455/src/net/lookup_unix.go:31
		}
//line /snap/go/10455/src/net/lookup_unix.go:31
		// _ = "end of CoverTab[7260]"
//line /snap/go/10455/src/net/lookup_unix.go:31
		_go_fuzz_dep_.CoverTab[7261]++
								f := getFields(line)
								if len(f) < 2 {
//line /snap/go/10455/src/net/lookup_unix.go:33
			_go_fuzz_dep_.CoverTab[529150]++
//line /snap/go/10455/src/net/lookup_unix.go:33
			_go_fuzz_dep_.CoverTab[7265]++
									continue
//line /snap/go/10455/src/net/lookup_unix.go:34
			// _ = "end of CoverTab[7265]"
		} else {
//line /snap/go/10455/src/net/lookup_unix.go:35
			_go_fuzz_dep_.CoverTab[529151]++
//line /snap/go/10455/src/net/lookup_unix.go:35
			_go_fuzz_dep_.CoverTab[7266]++
//line /snap/go/10455/src/net/lookup_unix.go:35
			// _ = "end of CoverTab[7266]"
//line /snap/go/10455/src/net/lookup_unix.go:35
		}
//line /snap/go/10455/src/net/lookup_unix.go:35
		// _ = "end of CoverTab[7261]"
//line /snap/go/10455/src/net/lookup_unix.go:35
		_go_fuzz_dep_.CoverTab[7262]++
								if proto, _, ok := dtoi(f[1]); ok {
//line /snap/go/10455/src/net/lookup_unix.go:36
			_go_fuzz_dep_.CoverTab[529152]++
//line /snap/go/10455/src/net/lookup_unix.go:36
			_go_fuzz_dep_.CoverTab[7267]++
									if _, ok := protocols[f[0]]; !ok {
//line /snap/go/10455/src/net/lookup_unix.go:37
				_go_fuzz_dep_.CoverTab[529154]++
//line /snap/go/10455/src/net/lookup_unix.go:37
				_go_fuzz_dep_.CoverTab[7269]++
										protocols[f[0]] = proto
//line /snap/go/10455/src/net/lookup_unix.go:38
				// _ = "end of CoverTab[7269]"
			} else {
//line /snap/go/10455/src/net/lookup_unix.go:39
				_go_fuzz_dep_.CoverTab[529155]++
//line /snap/go/10455/src/net/lookup_unix.go:39
				_go_fuzz_dep_.CoverTab[7270]++
//line /snap/go/10455/src/net/lookup_unix.go:39
				// _ = "end of CoverTab[7270]"
//line /snap/go/10455/src/net/lookup_unix.go:39
			}
//line /snap/go/10455/src/net/lookup_unix.go:39
			// _ = "end of CoverTab[7267]"
//line /snap/go/10455/src/net/lookup_unix.go:39
			_go_fuzz_dep_.CoverTab[7268]++
//line /snap/go/10455/src/net/lookup_unix.go:39
			_go_fuzz_dep_.CoverTab[786716] = 0
									for _, alias := range f[2:] {
//line /snap/go/10455/src/net/lookup_unix.go:40
				if _go_fuzz_dep_.CoverTab[786716] == 0 {
//line /snap/go/10455/src/net/lookup_unix.go:40
					_go_fuzz_dep_.CoverTab[529186]++
//line /snap/go/10455/src/net/lookup_unix.go:40
				} else {
//line /snap/go/10455/src/net/lookup_unix.go:40
					_go_fuzz_dep_.CoverTab[529187]++
//line /snap/go/10455/src/net/lookup_unix.go:40
				}
//line /snap/go/10455/src/net/lookup_unix.go:40
				_go_fuzz_dep_.CoverTab[786716] = 1
//line /snap/go/10455/src/net/lookup_unix.go:40
				_go_fuzz_dep_.CoverTab[7271]++
										if _, ok := protocols[alias]; !ok {
//line /snap/go/10455/src/net/lookup_unix.go:41
					_go_fuzz_dep_.CoverTab[529156]++
//line /snap/go/10455/src/net/lookup_unix.go:41
					_go_fuzz_dep_.CoverTab[7272]++
											protocols[alias] = proto
//line /snap/go/10455/src/net/lookup_unix.go:42
					// _ = "end of CoverTab[7272]"
				} else {
//line /snap/go/10455/src/net/lookup_unix.go:43
					_go_fuzz_dep_.CoverTab[529157]++
//line /snap/go/10455/src/net/lookup_unix.go:43
					_go_fuzz_dep_.CoverTab[7273]++
//line /snap/go/10455/src/net/lookup_unix.go:43
					// _ = "end of CoverTab[7273]"
//line /snap/go/10455/src/net/lookup_unix.go:43
				}
//line /snap/go/10455/src/net/lookup_unix.go:43
				// _ = "end of CoverTab[7271]"
			}
//line /snap/go/10455/src/net/lookup_unix.go:44
			if _go_fuzz_dep_.CoverTab[786716] == 0 {
//line /snap/go/10455/src/net/lookup_unix.go:44
				_go_fuzz_dep_.CoverTab[529188]++
//line /snap/go/10455/src/net/lookup_unix.go:44
			} else {
//line /snap/go/10455/src/net/lookup_unix.go:44
				_go_fuzz_dep_.CoverTab[529189]++
//line /snap/go/10455/src/net/lookup_unix.go:44
			}
//line /snap/go/10455/src/net/lookup_unix.go:44
			// _ = "end of CoverTab[7268]"
		} else {
//line /snap/go/10455/src/net/lookup_unix.go:45
			_go_fuzz_dep_.CoverTab[529153]++
//line /snap/go/10455/src/net/lookup_unix.go:45
			_go_fuzz_dep_.CoverTab[7274]++
//line /snap/go/10455/src/net/lookup_unix.go:45
			// _ = "end of CoverTab[7274]"
//line /snap/go/10455/src/net/lookup_unix.go:45
		}
//line /snap/go/10455/src/net/lookup_unix.go:45
		// _ = "end of CoverTab[7262]"
	}
//line /snap/go/10455/src/net/lookup_unix.go:46
	if _go_fuzz_dep_.CoverTab[786715] == 0 {
//line /snap/go/10455/src/net/lookup_unix.go:46
		_go_fuzz_dep_.CoverTab[529184]++
//line /snap/go/10455/src/net/lookup_unix.go:46
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:46
		_go_fuzz_dep_.CoverTab[529185]++
//line /snap/go/10455/src/net/lookup_unix.go:46
	}
//line /snap/go/10455/src/net/lookup_unix.go:46
	// _ = "end of CoverTab[7257]"
}

// lookupProtocol looks up IP protocol name in /etc/protocols and
//line /snap/go/10455/src/net/lookup_unix.go:49
// returns correspondent protocol number.
//line /snap/go/10455/src/net/lookup_unix.go:51
func lookupProtocol(_ context.Context, name string) (int, error) {
//line /snap/go/10455/src/net/lookup_unix.go:51
	_go_fuzz_dep_.CoverTab[7275]++
							onceReadProtocols.Do(readProtocols)
							return lookupProtocolMap(name)
//line /snap/go/10455/src/net/lookup_unix.go:53
	// _ = "end of CoverTab[7275]"
}

func (r *Resolver) lookupHost(ctx context.Context, host string) (addrs []string, err error) {
//line /snap/go/10455/src/net/lookup_unix.go:56
	_go_fuzz_dep_.CoverTab[7276]++
							order, conf := systemConf().hostLookupOrder(r, host)
							if order == hostLookupCgo {
//line /snap/go/10455/src/net/lookup_unix.go:58
		_go_fuzz_dep_.CoverTab[529158]++
//line /snap/go/10455/src/net/lookup_unix.go:58
		_go_fuzz_dep_.CoverTab[7278]++
								return cgoLookupHost(ctx, host)
//line /snap/go/10455/src/net/lookup_unix.go:59
		// _ = "end of CoverTab[7278]"
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:60
		_go_fuzz_dep_.CoverTab[529159]++
//line /snap/go/10455/src/net/lookup_unix.go:60
		_go_fuzz_dep_.CoverTab[7279]++
//line /snap/go/10455/src/net/lookup_unix.go:60
		// _ = "end of CoverTab[7279]"
//line /snap/go/10455/src/net/lookup_unix.go:60
	}
//line /snap/go/10455/src/net/lookup_unix.go:60
	// _ = "end of CoverTab[7276]"
//line /snap/go/10455/src/net/lookup_unix.go:60
	_go_fuzz_dep_.CoverTab[7277]++
							return r.goLookupHostOrder(ctx, host, order, conf)
//line /snap/go/10455/src/net/lookup_unix.go:61
	// _ = "end of CoverTab[7277]"
}

func (r *Resolver) lookupIP(ctx context.Context, network, host string) (addrs []IPAddr, err error) {
//line /snap/go/10455/src/net/lookup_unix.go:64
	_go_fuzz_dep_.CoverTab[7280]++
							if r.preferGo() {
//line /snap/go/10455/src/net/lookup_unix.go:65
		_go_fuzz_dep_.CoverTab[529160]++
//line /snap/go/10455/src/net/lookup_unix.go:65
		_go_fuzz_dep_.CoverTab[7283]++
								return r.goLookupIP(ctx, network, host)
//line /snap/go/10455/src/net/lookup_unix.go:66
		// _ = "end of CoverTab[7283]"
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:67
		_go_fuzz_dep_.CoverTab[529161]++
//line /snap/go/10455/src/net/lookup_unix.go:67
		_go_fuzz_dep_.CoverTab[7284]++
//line /snap/go/10455/src/net/lookup_unix.go:67
		// _ = "end of CoverTab[7284]"
//line /snap/go/10455/src/net/lookup_unix.go:67
	}
//line /snap/go/10455/src/net/lookup_unix.go:67
	// _ = "end of CoverTab[7280]"
//line /snap/go/10455/src/net/lookup_unix.go:67
	_go_fuzz_dep_.CoverTab[7281]++
							order, conf := systemConf().hostLookupOrder(r, host)
							if order == hostLookupCgo {
//line /snap/go/10455/src/net/lookup_unix.go:69
		_go_fuzz_dep_.CoverTab[529162]++
//line /snap/go/10455/src/net/lookup_unix.go:69
		_go_fuzz_dep_.CoverTab[7285]++
								return cgoLookupIP(ctx, network, host)
//line /snap/go/10455/src/net/lookup_unix.go:70
		// _ = "end of CoverTab[7285]"
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:71
		_go_fuzz_dep_.CoverTab[529163]++
//line /snap/go/10455/src/net/lookup_unix.go:71
		_go_fuzz_dep_.CoverTab[7286]++
//line /snap/go/10455/src/net/lookup_unix.go:71
		// _ = "end of CoverTab[7286]"
//line /snap/go/10455/src/net/lookup_unix.go:71
	}
//line /snap/go/10455/src/net/lookup_unix.go:71
	// _ = "end of CoverTab[7281]"
//line /snap/go/10455/src/net/lookup_unix.go:71
	_go_fuzz_dep_.CoverTab[7282]++
							ips, _, err := r.goLookupIPCNAMEOrder(ctx, network, host, order, conf)
							return ips, err
//line /snap/go/10455/src/net/lookup_unix.go:73
	// _ = "end of CoverTab[7282]"
}

func (r *Resolver) lookupPort(ctx context.Context, network, service string) (int, error) {
//line /snap/go/10455/src/net/lookup_unix.go:76
	_go_fuzz_dep_.CoverTab[7287]++

//line /snap/go/10455/src/net/lookup_unix.go:79
	if !systemConf().mustUseGoResolver(r) {
//line /snap/go/10455/src/net/lookup_unix.go:79
		_go_fuzz_dep_.CoverTab[529164]++
//line /snap/go/10455/src/net/lookup_unix.go:79
		_go_fuzz_dep_.CoverTab[7289]++
								port, err := cgoLookupPort(ctx, network, service)
								if err != nil {
//line /snap/go/10455/src/net/lookup_unix.go:81
			_go_fuzz_dep_.CoverTab[529166]++
//line /snap/go/10455/src/net/lookup_unix.go:81
			_go_fuzz_dep_.CoverTab[7291]++

//line /snap/go/10455/src/net/lookup_unix.go:84
			if port, err := goLookupPort(network, service); err == nil {
//line /snap/go/10455/src/net/lookup_unix.go:84
				_go_fuzz_dep_.CoverTab[529168]++
//line /snap/go/10455/src/net/lookup_unix.go:84
				_go_fuzz_dep_.CoverTab[7292]++
										return port, nil
//line /snap/go/10455/src/net/lookup_unix.go:85
				// _ = "end of CoverTab[7292]"
			} else {
//line /snap/go/10455/src/net/lookup_unix.go:86
				_go_fuzz_dep_.CoverTab[529169]++
//line /snap/go/10455/src/net/lookup_unix.go:86
				_go_fuzz_dep_.CoverTab[7293]++
//line /snap/go/10455/src/net/lookup_unix.go:86
				// _ = "end of CoverTab[7293]"
//line /snap/go/10455/src/net/lookup_unix.go:86
			}
//line /snap/go/10455/src/net/lookup_unix.go:86
			// _ = "end of CoverTab[7291]"
		} else {
//line /snap/go/10455/src/net/lookup_unix.go:87
			_go_fuzz_dep_.CoverTab[529167]++
//line /snap/go/10455/src/net/lookup_unix.go:87
			_go_fuzz_dep_.CoverTab[7294]++
//line /snap/go/10455/src/net/lookup_unix.go:87
			// _ = "end of CoverTab[7294]"
//line /snap/go/10455/src/net/lookup_unix.go:87
		}
//line /snap/go/10455/src/net/lookup_unix.go:87
		// _ = "end of CoverTab[7289]"
//line /snap/go/10455/src/net/lookup_unix.go:87
		_go_fuzz_dep_.CoverTab[7290]++
								return port, err
//line /snap/go/10455/src/net/lookup_unix.go:88
		// _ = "end of CoverTab[7290]"
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:89
		_go_fuzz_dep_.CoverTab[529165]++
//line /snap/go/10455/src/net/lookup_unix.go:89
		_go_fuzz_dep_.CoverTab[7295]++
//line /snap/go/10455/src/net/lookup_unix.go:89
		// _ = "end of CoverTab[7295]"
//line /snap/go/10455/src/net/lookup_unix.go:89
	}
//line /snap/go/10455/src/net/lookup_unix.go:89
	// _ = "end of CoverTab[7287]"
//line /snap/go/10455/src/net/lookup_unix.go:89
	_go_fuzz_dep_.CoverTab[7288]++
							return goLookupPort(network, service)
//line /snap/go/10455/src/net/lookup_unix.go:90
	// _ = "end of CoverTab[7288]"
}

func (r *Resolver) lookupCNAME(ctx context.Context, name string) (string, error) {
//line /snap/go/10455/src/net/lookup_unix.go:93
	_go_fuzz_dep_.CoverTab[7296]++
							order, conf := systemConf().hostLookupOrder(r, name)
							if order == hostLookupCgo {
//line /snap/go/10455/src/net/lookup_unix.go:95
		_go_fuzz_dep_.CoverTab[529170]++
//line /snap/go/10455/src/net/lookup_unix.go:95
		_go_fuzz_dep_.CoverTab[7298]++
								if cname, err, ok := cgoLookupCNAME(ctx, name); ok {
//line /snap/go/10455/src/net/lookup_unix.go:96
			_go_fuzz_dep_.CoverTab[529172]++
//line /snap/go/10455/src/net/lookup_unix.go:96
			_go_fuzz_dep_.CoverTab[7299]++
									return cname, err
//line /snap/go/10455/src/net/lookup_unix.go:97
			// _ = "end of CoverTab[7299]"
		} else {
//line /snap/go/10455/src/net/lookup_unix.go:98
			_go_fuzz_dep_.CoverTab[529173]++
//line /snap/go/10455/src/net/lookup_unix.go:98
			_go_fuzz_dep_.CoverTab[7300]++
//line /snap/go/10455/src/net/lookup_unix.go:98
			// _ = "end of CoverTab[7300]"
//line /snap/go/10455/src/net/lookup_unix.go:98
		}
//line /snap/go/10455/src/net/lookup_unix.go:98
		// _ = "end of CoverTab[7298]"
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:99
		_go_fuzz_dep_.CoverTab[529171]++
//line /snap/go/10455/src/net/lookup_unix.go:99
		_go_fuzz_dep_.CoverTab[7301]++
//line /snap/go/10455/src/net/lookup_unix.go:99
		// _ = "end of CoverTab[7301]"
//line /snap/go/10455/src/net/lookup_unix.go:99
	}
//line /snap/go/10455/src/net/lookup_unix.go:99
	// _ = "end of CoverTab[7296]"
//line /snap/go/10455/src/net/lookup_unix.go:99
	_go_fuzz_dep_.CoverTab[7297]++
							return r.goLookupCNAME(ctx, name, order, conf)
//line /snap/go/10455/src/net/lookup_unix.go:100
	// _ = "end of CoverTab[7297]"
}

func (r *Resolver) lookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error) {
//line /snap/go/10455/src/net/lookup_unix.go:103
	_go_fuzz_dep_.CoverTab[7302]++
							return r.goLookupSRV(ctx, service, proto, name)
//line /snap/go/10455/src/net/lookup_unix.go:104
	// _ = "end of CoverTab[7302]"
}

func (r *Resolver) lookupMX(ctx context.Context, name string) ([]*MX, error) {
//line /snap/go/10455/src/net/lookup_unix.go:107
	_go_fuzz_dep_.CoverTab[7303]++
							return r.goLookupMX(ctx, name)
//line /snap/go/10455/src/net/lookup_unix.go:108
	// _ = "end of CoverTab[7303]"
}

func (r *Resolver) lookupNS(ctx context.Context, name string) ([]*NS, error) {
//line /snap/go/10455/src/net/lookup_unix.go:111
	_go_fuzz_dep_.CoverTab[7304]++
							return r.goLookupNS(ctx, name)
//line /snap/go/10455/src/net/lookup_unix.go:112
	// _ = "end of CoverTab[7304]"
}

func (r *Resolver) lookupTXT(ctx context.Context, name string) ([]string, error) {
//line /snap/go/10455/src/net/lookup_unix.go:115
	_go_fuzz_dep_.CoverTab[7305]++
							return r.goLookupTXT(ctx, name)
//line /snap/go/10455/src/net/lookup_unix.go:116
	// _ = "end of CoverTab[7305]"
}

func (r *Resolver) lookupAddr(ctx context.Context, addr string) ([]string, error) {
//line /snap/go/10455/src/net/lookup_unix.go:119
	_go_fuzz_dep_.CoverTab[7306]++
							order, conf := systemConf().addrLookupOrder(r, addr)
							if order == hostLookupCgo {
//line /snap/go/10455/src/net/lookup_unix.go:121
		_go_fuzz_dep_.CoverTab[529174]++
//line /snap/go/10455/src/net/lookup_unix.go:121
		_go_fuzz_dep_.CoverTab[7308]++
								return cgoLookupPTR(ctx, addr)
//line /snap/go/10455/src/net/lookup_unix.go:122
		// _ = "end of CoverTab[7308]"
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:123
		_go_fuzz_dep_.CoverTab[529175]++
//line /snap/go/10455/src/net/lookup_unix.go:123
		_go_fuzz_dep_.CoverTab[7309]++
//line /snap/go/10455/src/net/lookup_unix.go:123
		// _ = "end of CoverTab[7309]"
//line /snap/go/10455/src/net/lookup_unix.go:123
	}
//line /snap/go/10455/src/net/lookup_unix.go:123
	// _ = "end of CoverTab[7306]"
//line /snap/go/10455/src/net/lookup_unix.go:123
	_go_fuzz_dep_.CoverTab[7307]++
							return r.goLookupPTR(ctx, addr, order, conf)
//line /snap/go/10455/src/net/lookup_unix.go:124
	// _ = "end of CoverTab[7307]"
}

// concurrentThreadsLimit returns the number of threads we permit to
//line /snap/go/10455/src/net/lookup_unix.go:127
// run concurrently doing DNS lookups via cgo. A DNS lookup may use a
//line /snap/go/10455/src/net/lookup_unix.go:127
// file descriptor so we limit this to less than the number of
//line /snap/go/10455/src/net/lookup_unix.go:127
// permitted open files. On some systems, notably Darwin, if
//line /snap/go/10455/src/net/lookup_unix.go:127
// getaddrinfo is unable to open a file descriptor it simply returns
//line /snap/go/10455/src/net/lookup_unix.go:127
// EAI_NONAME rather than a useful error. Limiting the number of
//line /snap/go/10455/src/net/lookup_unix.go:127
// concurrent getaddrinfo calls to less than the permitted number of
//line /snap/go/10455/src/net/lookup_unix.go:127
// file descriptors makes that error less likely. We don't bother to
//line /snap/go/10455/src/net/lookup_unix.go:127
// apply the same limit to DNS lookups run directly from Go, because
//line /snap/go/10455/src/net/lookup_unix.go:127
// there we will return a meaningful "too many open files" error.
//line /snap/go/10455/src/net/lookup_unix.go:137
func concurrentThreadsLimit() int {
//line /snap/go/10455/src/net/lookup_unix.go:137
	_go_fuzz_dep_.CoverTab[7310]++
							var rlim syscall.Rlimit
							if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim); err != nil {
//line /snap/go/10455/src/net/lookup_unix.go:139
		_go_fuzz_dep_.CoverTab[529176]++
//line /snap/go/10455/src/net/lookup_unix.go:139
		_go_fuzz_dep_.CoverTab[7313]++
								return 500
//line /snap/go/10455/src/net/lookup_unix.go:140
		// _ = "end of CoverTab[7313]"
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:141
		_go_fuzz_dep_.CoverTab[529177]++
//line /snap/go/10455/src/net/lookup_unix.go:141
		_go_fuzz_dep_.CoverTab[7314]++
//line /snap/go/10455/src/net/lookup_unix.go:141
		// _ = "end of CoverTab[7314]"
//line /snap/go/10455/src/net/lookup_unix.go:141
	}
//line /snap/go/10455/src/net/lookup_unix.go:141
	// _ = "end of CoverTab[7310]"
//line /snap/go/10455/src/net/lookup_unix.go:141
	_go_fuzz_dep_.CoverTab[7311]++
							r := rlim.Cur
							if r > 500 {
//line /snap/go/10455/src/net/lookup_unix.go:143
		_go_fuzz_dep_.CoverTab[529178]++
//line /snap/go/10455/src/net/lookup_unix.go:143
		_go_fuzz_dep_.CoverTab[7315]++
								r = 500
//line /snap/go/10455/src/net/lookup_unix.go:144
		// _ = "end of CoverTab[7315]"
	} else {
//line /snap/go/10455/src/net/lookup_unix.go:145
		_go_fuzz_dep_.CoverTab[529179]++
//line /snap/go/10455/src/net/lookup_unix.go:145
		_go_fuzz_dep_.CoverTab[7316]++
//line /snap/go/10455/src/net/lookup_unix.go:145
		if r > 30 {
//line /snap/go/10455/src/net/lookup_unix.go:145
			_go_fuzz_dep_.CoverTab[529180]++
//line /snap/go/10455/src/net/lookup_unix.go:145
			_go_fuzz_dep_.CoverTab[7317]++
									r -= 30
//line /snap/go/10455/src/net/lookup_unix.go:146
			// _ = "end of CoverTab[7317]"
		} else {
//line /snap/go/10455/src/net/lookup_unix.go:147
			_go_fuzz_dep_.CoverTab[529181]++
//line /snap/go/10455/src/net/lookup_unix.go:147
			_go_fuzz_dep_.CoverTab[7318]++
//line /snap/go/10455/src/net/lookup_unix.go:147
			// _ = "end of CoverTab[7318]"
//line /snap/go/10455/src/net/lookup_unix.go:147
		}
//line /snap/go/10455/src/net/lookup_unix.go:147
		// _ = "end of CoverTab[7316]"
//line /snap/go/10455/src/net/lookup_unix.go:147
	}
//line /snap/go/10455/src/net/lookup_unix.go:147
	// _ = "end of CoverTab[7311]"
//line /snap/go/10455/src/net/lookup_unix.go:147
	_go_fuzz_dep_.CoverTab[7312]++
							return int(r)
//line /snap/go/10455/src/net/lookup_unix.go:148
	// _ = "end of CoverTab[7312]"
}

//line /snap/go/10455/src/net/lookup_unix.go:149
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/lookup_unix.go:149
var _ = _go_fuzz_dep_.CoverTab
