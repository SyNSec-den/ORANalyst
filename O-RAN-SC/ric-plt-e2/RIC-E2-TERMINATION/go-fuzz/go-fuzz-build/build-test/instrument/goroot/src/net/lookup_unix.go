// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

//line /usr/local/go/src/net/lookup_unix.go:7
package net

//line /usr/local/go/src/net/lookup_unix.go:7
import (
//line /usr/local/go/src/net/lookup_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/lookup_unix.go:7
)
//line /usr/local/go/src/net/lookup_unix.go:7
import (
//line /usr/local/go/src/net/lookup_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/lookup_unix.go:7
)

import (
	"context"
	"internal/bytealg"
	"sync"
	"syscall"
)

var onceReadProtocols sync.Once

// readProtocols loads contents of /etc/protocols into protocols map
//line /usr/local/go/src/net/lookup_unix.go:18
// for quick access.
//line /usr/local/go/src/net/lookup_unix.go:20
func readProtocols() {
//line /usr/local/go/src/net/lookup_unix.go:20
	_go_fuzz_dep_.CoverTab[6990]++
						file, err := open("/etc/protocols")
						if err != nil {
//line /usr/local/go/src/net/lookup_unix.go:22
		_go_fuzz_dep_.CoverTab[6992]++
							return
//line /usr/local/go/src/net/lookup_unix.go:23
		// _ = "end of CoverTab[6992]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:24
		_go_fuzz_dep_.CoverTab[6993]++
//line /usr/local/go/src/net/lookup_unix.go:24
		// _ = "end of CoverTab[6993]"
//line /usr/local/go/src/net/lookup_unix.go:24
	}
//line /usr/local/go/src/net/lookup_unix.go:24
	// _ = "end of CoverTab[6990]"
//line /usr/local/go/src/net/lookup_unix.go:24
	_go_fuzz_dep_.CoverTab[6991]++
						defer file.close()

						for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /usr/local/go/src/net/lookup_unix.go:27
		_go_fuzz_dep_.CoverTab[6994]++

							if i := bytealg.IndexByteString(line, '#'); i >= 0 {
//line /usr/local/go/src/net/lookup_unix.go:29
			_go_fuzz_dep_.CoverTab[6997]++
								line = line[0:i]
//line /usr/local/go/src/net/lookup_unix.go:30
			// _ = "end of CoverTab[6997]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:31
			_go_fuzz_dep_.CoverTab[6998]++
//line /usr/local/go/src/net/lookup_unix.go:31
			// _ = "end of CoverTab[6998]"
//line /usr/local/go/src/net/lookup_unix.go:31
		}
//line /usr/local/go/src/net/lookup_unix.go:31
		// _ = "end of CoverTab[6994]"
//line /usr/local/go/src/net/lookup_unix.go:31
		_go_fuzz_dep_.CoverTab[6995]++
							f := getFields(line)
							if len(f) < 2 {
//line /usr/local/go/src/net/lookup_unix.go:33
			_go_fuzz_dep_.CoverTab[6999]++
								continue
//line /usr/local/go/src/net/lookup_unix.go:34
			// _ = "end of CoverTab[6999]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:35
			_go_fuzz_dep_.CoverTab[7000]++
//line /usr/local/go/src/net/lookup_unix.go:35
			// _ = "end of CoverTab[7000]"
//line /usr/local/go/src/net/lookup_unix.go:35
		}
//line /usr/local/go/src/net/lookup_unix.go:35
		// _ = "end of CoverTab[6995]"
//line /usr/local/go/src/net/lookup_unix.go:35
		_go_fuzz_dep_.CoverTab[6996]++
							if proto, _, ok := dtoi(f[1]); ok {
//line /usr/local/go/src/net/lookup_unix.go:36
			_go_fuzz_dep_.CoverTab[7001]++
								if _, ok := protocols[f[0]]; !ok {
//line /usr/local/go/src/net/lookup_unix.go:37
				_go_fuzz_dep_.CoverTab[7003]++
									protocols[f[0]] = proto
//line /usr/local/go/src/net/lookup_unix.go:38
				// _ = "end of CoverTab[7003]"
			} else {
//line /usr/local/go/src/net/lookup_unix.go:39
				_go_fuzz_dep_.CoverTab[7004]++
//line /usr/local/go/src/net/lookup_unix.go:39
				// _ = "end of CoverTab[7004]"
//line /usr/local/go/src/net/lookup_unix.go:39
			}
//line /usr/local/go/src/net/lookup_unix.go:39
			// _ = "end of CoverTab[7001]"
//line /usr/local/go/src/net/lookup_unix.go:39
			_go_fuzz_dep_.CoverTab[7002]++
								for _, alias := range f[2:] {
//line /usr/local/go/src/net/lookup_unix.go:40
				_go_fuzz_dep_.CoverTab[7005]++
									if _, ok := protocols[alias]; !ok {
//line /usr/local/go/src/net/lookup_unix.go:41
					_go_fuzz_dep_.CoverTab[7006]++
										protocols[alias] = proto
//line /usr/local/go/src/net/lookup_unix.go:42
					// _ = "end of CoverTab[7006]"
				} else {
//line /usr/local/go/src/net/lookup_unix.go:43
					_go_fuzz_dep_.CoverTab[7007]++
//line /usr/local/go/src/net/lookup_unix.go:43
					// _ = "end of CoverTab[7007]"
//line /usr/local/go/src/net/lookup_unix.go:43
				}
//line /usr/local/go/src/net/lookup_unix.go:43
				// _ = "end of CoverTab[7005]"
			}
//line /usr/local/go/src/net/lookup_unix.go:44
			// _ = "end of CoverTab[7002]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:45
			_go_fuzz_dep_.CoverTab[7008]++
//line /usr/local/go/src/net/lookup_unix.go:45
			// _ = "end of CoverTab[7008]"
//line /usr/local/go/src/net/lookup_unix.go:45
		}
//line /usr/local/go/src/net/lookup_unix.go:45
		// _ = "end of CoverTab[6996]"
	}
//line /usr/local/go/src/net/lookup_unix.go:46
	// _ = "end of CoverTab[6991]"
}

// lookupProtocol looks up IP protocol name in /etc/protocols and
//line /usr/local/go/src/net/lookup_unix.go:49
// returns correspondent protocol number.
//line /usr/local/go/src/net/lookup_unix.go:51
func lookupProtocol(_ context.Context, name string) (int, error) {
//line /usr/local/go/src/net/lookup_unix.go:51
	_go_fuzz_dep_.CoverTab[7009]++
						onceReadProtocols.Do(readProtocols)
						return lookupProtocolMap(name)
//line /usr/local/go/src/net/lookup_unix.go:53
	// _ = "end of CoverTab[7009]"
}

func (r *Resolver) lookupHost(ctx context.Context, host string) (addrs []string, err error) {
//line /usr/local/go/src/net/lookup_unix.go:56
	_go_fuzz_dep_.CoverTab[7010]++
						order, conf := systemConf().hostLookupOrder(r, host)
						if !r.preferGo() && func() bool {
//line /usr/local/go/src/net/lookup_unix.go:58
		_go_fuzz_dep_.CoverTab[7012]++
//line /usr/local/go/src/net/lookup_unix.go:58
		return order == hostLookupCgo
//line /usr/local/go/src/net/lookup_unix.go:58
		// _ = "end of CoverTab[7012]"
//line /usr/local/go/src/net/lookup_unix.go:58
	}() {
//line /usr/local/go/src/net/lookup_unix.go:58
		_go_fuzz_dep_.CoverTab[7013]++
							if addrs, err, ok := cgoLookupHost(ctx, host); ok {
//line /usr/local/go/src/net/lookup_unix.go:59
			_go_fuzz_dep_.CoverTab[7015]++
								return addrs, err
//line /usr/local/go/src/net/lookup_unix.go:60
			// _ = "end of CoverTab[7015]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:61
			_go_fuzz_dep_.CoverTab[7016]++
//line /usr/local/go/src/net/lookup_unix.go:61
			// _ = "end of CoverTab[7016]"
//line /usr/local/go/src/net/lookup_unix.go:61
		}
//line /usr/local/go/src/net/lookup_unix.go:61
		// _ = "end of CoverTab[7013]"
//line /usr/local/go/src/net/lookup_unix.go:61
		_go_fuzz_dep_.CoverTab[7014]++

							order = hostLookupFilesDNS
//line /usr/local/go/src/net/lookup_unix.go:63
		// _ = "end of CoverTab[7014]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:64
		_go_fuzz_dep_.CoverTab[7017]++
//line /usr/local/go/src/net/lookup_unix.go:64
		// _ = "end of CoverTab[7017]"
//line /usr/local/go/src/net/lookup_unix.go:64
	}
//line /usr/local/go/src/net/lookup_unix.go:64
	// _ = "end of CoverTab[7010]"
//line /usr/local/go/src/net/lookup_unix.go:64
	_go_fuzz_dep_.CoverTab[7011]++
						return r.goLookupHostOrder(ctx, host, order, conf)
//line /usr/local/go/src/net/lookup_unix.go:65
	// _ = "end of CoverTab[7011]"
}

func (r *Resolver) lookupIP(ctx context.Context, network, host string) (addrs []IPAddr, err error) {
//line /usr/local/go/src/net/lookup_unix.go:68
	_go_fuzz_dep_.CoverTab[7018]++
						if r.preferGo() {
//line /usr/local/go/src/net/lookup_unix.go:69
		_go_fuzz_dep_.CoverTab[7021]++
							return r.goLookupIP(ctx, network, host)
//line /usr/local/go/src/net/lookup_unix.go:70
		// _ = "end of CoverTab[7021]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:71
		_go_fuzz_dep_.CoverTab[7022]++
//line /usr/local/go/src/net/lookup_unix.go:71
		// _ = "end of CoverTab[7022]"
//line /usr/local/go/src/net/lookup_unix.go:71
	}
//line /usr/local/go/src/net/lookup_unix.go:71
	// _ = "end of CoverTab[7018]"
//line /usr/local/go/src/net/lookup_unix.go:71
	_go_fuzz_dep_.CoverTab[7019]++
						order, conf := systemConf().hostLookupOrder(r, host)
						if order == hostLookupCgo {
//line /usr/local/go/src/net/lookup_unix.go:73
		_go_fuzz_dep_.CoverTab[7023]++
							if addrs, err, ok := cgoLookupIP(ctx, network, host); ok {
//line /usr/local/go/src/net/lookup_unix.go:74
			_go_fuzz_dep_.CoverTab[7025]++
								return addrs, err
//line /usr/local/go/src/net/lookup_unix.go:75
			// _ = "end of CoverTab[7025]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:76
			_go_fuzz_dep_.CoverTab[7026]++
//line /usr/local/go/src/net/lookup_unix.go:76
			// _ = "end of CoverTab[7026]"
//line /usr/local/go/src/net/lookup_unix.go:76
		}
//line /usr/local/go/src/net/lookup_unix.go:76
		// _ = "end of CoverTab[7023]"
//line /usr/local/go/src/net/lookup_unix.go:76
		_go_fuzz_dep_.CoverTab[7024]++

							order = hostLookupFilesDNS
//line /usr/local/go/src/net/lookup_unix.go:78
		// _ = "end of CoverTab[7024]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:79
		_go_fuzz_dep_.CoverTab[7027]++
//line /usr/local/go/src/net/lookup_unix.go:79
		// _ = "end of CoverTab[7027]"
//line /usr/local/go/src/net/lookup_unix.go:79
	}
//line /usr/local/go/src/net/lookup_unix.go:79
	// _ = "end of CoverTab[7019]"
//line /usr/local/go/src/net/lookup_unix.go:79
	_go_fuzz_dep_.CoverTab[7020]++
						ips, _, err := r.goLookupIPCNAMEOrder(ctx, network, host, order, conf)
						return ips, err
//line /usr/local/go/src/net/lookup_unix.go:81
	// _ = "end of CoverTab[7020]"
}

func (r *Resolver) lookupPort(ctx context.Context, network, service string) (int, error) {
//line /usr/local/go/src/net/lookup_unix.go:84
	_go_fuzz_dep_.CoverTab[7028]++
						if !r.preferGo() && func() bool {
//line /usr/local/go/src/net/lookup_unix.go:85
		_go_fuzz_dep_.CoverTab[7030]++
//line /usr/local/go/src/net/lookup_unix.go:85
		return systemConf().canUseCgo()
//line /usr/local/go/src/net/lookup_unix.go:85
		// _ = "end of CoverTab[7030]"
//line /usr/local/go/src/net/lookup_unix.go:85
	}() {
//line /usr/local/go/src/net/lookup_unix.go:85
		_go_fuzz_dep_.CoverTab[7031]++
							if port, err, ok := cgoLookupPort(ctx, network, service); ok {
//line /usr/local/go/src/net/lookup_unix.go:86
			_go_fuzz_dep_.CoverTab[7032]++
								if err != nil {
//line /usr/local/go/src/net/lookup_unix.go:87
				_go_fuzz_dep_.CoverTab[7034]++

//line /usr/local/go/src/net/lookup_unix.go:90
				if port, err := goLookupPort(network, service); err == nil {
//line /usr/local/go/src/net/lookup_unix.go:90
					_go_fuzz_dep_.CoverTab[7035]++
										return port, nil
//line /usr/local/go/src/net/lookup_unix.go:91
					// _ = "end of CoverTab[7035]"
				} else {
//line /usr/local/go/src/net/lookup_unix.go:92
					_go_fuzz_dep_.CoverTab[7036]++
//line /usr/local/go/src/net/lookup_unix.go:92
					// _ = "end of CoverTab[7036]"
//line /usr/local/go/src/net/lookup_unix.go:92
				}
//line /usr/local/go/src/net/lookup_unix.go:92
				// _ = "end of CoverTab[7034]"
			} else {
//line /usr/local/go/src/net/lookup_unix.go:93
				_go_fuzz_dep_.CoverTab[7037]++
//line /usr/local/go/src/net/lookup_unix.go:93
				// _ = "end of CoverTab[7037]"
//line /usr/local/go/src/net/lookup_unix.go:93
			}
//line /usr/local/go/src/net/lookup_unix.go:93
			// _ = "end of CoverTab[7032]"
//line /usr/local/go/src/net/lookup_unix.go:93
			_go_fuzz_dep_.CoverTab[7033]++
								return port, err
//line /usr/local/go/src/net/lookup_unix.go:94
			// _ = "end of CoverTab[7033]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:95
			_go_fuzz_dep_.CoverTab[7038]++
//line /usr/local/go/src/net/lookup_unix.go:95
			// _ = "end of CoverTab[7038]"
//line /usr/local/go/src/net/lookup_unix.go:95
		}
//line /usr/local/go/src/net/lookup_unix.go:95
		// _ = "end of CoverTab[7031]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:96
		_go_fuzz_dep_.CoverTab[7039]++
//line /usr/local/go/src/net/lookup_unix.go:96
		// _ = "end of CoverTab[7039]"
//line /usr/local/go/src/net/lookup_unix.go:96
	}
//line /usr/local/go/src/net/lookup_unix.go:96
	// _ = "end of CoverTab[7028]"
//line /usr/local/go/src/net/lookup_unix.go:96
	_go_fuzz_dep_.CoverTab[7029]++
						return goLookupPort(network, service)
//line /usr/local/go/src/net/lookup_unix.go:97
	// _ = "end of CoverTab[7029]"
}

func (r *Resolver) lookupCNAME(ctx context.Context, name string) (string, error) {
//line /usr/local/go/src/net/lookup_unix.go:100
	_go_fuzz_dep_.CoverTab[7040]++
							order, conf := systemConf().hostLookupOrder(r, name)
							if !r.preferGo() && func() bool {
//line /usr/local/go/src/net/lookup_unix.go:102
		_go_fuzz_dep_.CoverTab[7042]++
//line /usr/local/go/src/net/lookup_unix.go:102
		return order == hostLookupCgo
//line /usr/local/go/src/net/lookup_unix.go:102
		// _ = "end of CoverTab[7042]"
//line /usr/local/go/src/net/lookup_unix.go:102
	}() {
//line /usr/local/go/src/net/lookup_unix.go:102
		_go_fuzz_dep_.CoverTab[7043]++
								if cname, err, ok := cgoLookupCNAME(ctx, name); ok {
//line /usr/local/go/src/net/lookup_unix.go:103
			_go_fuzz_dep_.CoverTab[7044]++
									return cname, err
//line /usr/local/go/src/net/lookup_unix.go:104
			// _ = "end of CoverTab[7044]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:105
			_go_fuzz_dep_.CoverTab[7045]++
//line /usr/local/go/src/net/lookup_unix.go:105
			// _ = "end of CoverTab[7045]"
//line /usr/local/go/src/net/lookup_unix.go:105
		}
//line /usr/local/go/src/net/lookup_unix.go:105
		// _ = "end of CoverTab[7043]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:106
		_go_fuzz_dep_.CoverTab[7046]++
//line /usr/local/go/src/net/lookup_unix.go:106
		// _ = "end of CoverTab[7046]"
//line /usr/local/go/src/net/lookup_unix.go:106
	}
//line /usr/local/go/src/net/lookup_unix.go:106
	// _ = "end of CoverTab[7040]"
//line /usr/local/go/src/net/lookup_unix.go:106
	_go_fuzz_dep_.CoverTab[7041]++
							return r.goLookupCNAME(ctx, name, order, conf)
//line /usr/local/go/src/net/lookup_unix.go:107
	// _ = "end of CoverTab[7041]"
}

func (r *Resolver) lookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error) {
//line /usr/local/go/src/net/lookup_unix.go:110
	_go_fuzz_dep_.CoverTab[7047]++
							return r.goLookupSRV(ctx, service, proto, name)
//line /usr/local/go/src/net/lookup_unix.go:111
	// _ = "end of CoverTab[7047]"
}

func (r *Resolver) lookupMX(ctx context.Context, name string) ([]*MX, error) {
//line /usr/local/go/src/net/lookup_unix.go:114
	_go_fuzz_dep_.CoverTab[7048]++
							return r.goLookupMX(ctx, name)
//line /usr/local/go/src/net/lookup_unix.go:115
	// _ = "end of CoverTab[7048]"
}

func (r *Resolver) lookupNS(ctx context.Context, name string) ([]*NS, error) {
//line /usr/local/go/src/net/lookup_unix.go:118
	_go_fuzz_dep_.CoverTab[7049]++
							return r.goLookupNS(ctx, name)
//line /usr/local/go/src/net/lookup_unix.go:119
	// _ = "end of CoverTab[7049]"
}

func (r *Resolver) lookupTXT(ctx context.Context, name string) ([]string, error) {
//line /usr/local/go/src/net/lookup_unix.go:122
	_go_fuzz_dep_.CoverTab[7050]++
							return r.goLookupTXT(ctx, name)
//line /usr/local/go/src/net/lookup_unix.go:123
	// _ = "end of CoverTab[7050]"
}

func (r *Resolver) lookupAddr(ctx context.Context, addr string) ([]string, error) {
//line /usr/local/go/src/net/lookup_unix.go:126
	_go_fuzz_dep_.CoverTab[7051]++
							order, conf := systemConf().hostLookupOrder(r, "")
							if !r.preferGo() && func() bool {
//line /usr/local/go/src/net/lookup_unix.go:128
		_go_fuzz_dep_.CoverTab[7053]++
//line /usr/local/go/src/net/lookup_unix.go:128
		return order == hostLookupCgo
//line /usr/local/go/src/net/lookup_unix.go:128
		// _ = "end of CoverTab[7053]"
//line /usr/local/go/src/net/lookup_unix.go:128
	}() {
//line /usr/local/go/src/net/lookup_unix.go:128
		_go_fuzz_dep_.CoverTab[7054]++
								if ptrs, err, ok := cgoLookupPTR(ctx, addr); ok {
//line /usr/local/go/src/net/lookup_unix.go:129
			_go_fuzz_dep_.CoverTab[7055]++
									return ptrs, err
//line /usr/local/go/src/net/lookup_unix.go:130
			// _ = "end of CoverTab[7055]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:131
			_go_fuzz_dep_.CoverTab[7056]++
//line /usr/local/go/src/net/lookup_unix.go:131
			// _ = "end of CoverTab[7056]"
//line /usr/local/go/src/net/lookup_unix.go:131
		}
//line /usr/local/go/src/net/lookup_unix.go:131
		// _ = "end of CoverTab[7054]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:132
		_go_fuzz_dep_.CoverTab[7057]++
//line /usr/local/go/src/net/lookup_unix.go:132
		// _ = "end of CoverTab[7057]"
//line /usr/local/go/src/net/lookup_unix.go:132
	}
//line /usr/local/go/src/net/lookup_unix.go:132
	// _ = "end of CoverTab[7051]"
//line /usr/local/go/src/net/lookup_unix.go:132
	_go_fuzz_dep_.CoverTab[7052]++
							return r.goLookupPTR(ctx, addr, conf)
//line /usr/local/go/src/net/lookup_unix.go:133
	// _ = "end of CoverTab[7052]"
}

// concurrentThreadsLimit returns the number of threads we permit to
//line /usr/local/go/src/net/lookup_unix.go:136
// run concurrently doing DNS lookups via cgo. A DNS lookup may use a
//line /usr/local/go/src/net/lookup_unix.go:136
// file descriptor so we limit this to less than the number of
//line /usr/local/go/src/net/lookup_unix.go:136
// permitted open files. On some systems, notably Darwin, if
//line /usr/local/go/src/net/lookup_unix.go:136
// getaddrinfo is unable to open a file descriptor it simply returns
//line /usr/local/go/src/net/lookup_unix.go:136
// EAI_NONAME rather than a useful error. Limiting the number of
//line /usr/local/go/src/net/lookup_unix.go:136
// concurrent getaddrinfo calls to less than the permitted number of
//line /usr/local/go/src/net/lookup_unix.go:136
// file descriptors makes that error less likely. We don't bother to
//line /usr/local/go/src/net/lookup_unix.go:136
// apply the same limit to DNS lookups run directly from Go, because
//line /usr/local/go/src/net/lookup_unix.go:136
// there we will return a meaningful "too many open files" error.
//line /usr/local/go/src/net/lookup_unix.go:146
func concurrentThreadsLimit() int {
//line /usr/local/go/src/net/lookup_unix.go:146
	_go_fuzz_dep_.CoverTab[7058]++
							var rlim syscall.Rlimit
							if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim); err != nil {
//line /usr/local/go/src/net/lookup_unix.go:148
		_go_fuzz_dep_.CoverTab[7061]++
								return 500
//line /usr/local/go/src/net/lookup_unix.go:149
		// _ = "end of CoverTab[7061]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:150
		_go_fuzz_dep_.CoverTab[7062]++
//line /usr/local/go/src/net/lookup_unix.go:150
		// _ = "end of CoverTab[7062]"
//line /usr/local/go/src/net/lookup_unix.go:150
	}
//line /usr/local/go/src/net/lookup_unix.go:150
	// _ = "end of CoverTab[7058]"
//line /usr/local/go/src/net/lookup_unix.go:150
	_go_fuzz_dep_.CoverTab[7059]++
							r := int(rlim.Cur)
							if r > 500 {
//line /usr/local/go/src/net/lookup_unix.go:152
		_go_fuzz_dep_.CoverTab[7063]++
								r = 500
//line /usr/local/go/src/net/lookup_unix.go:153
		// _ = "end of CoverTab[7063]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:154
		_go_fuzz_dep_.CoverTab[7064]++
//line /usr/local/go/src/net/lookup_unix.go:154
		if r > 30 {
//line /usr/local/go/src/net/lookup_unix.go:154
			_go_fuzz_dep_.CoverTab[7065]++
									r -= 30
//line /usr/local/go/src/net/lookup_unix.go:155
			// _ = "end of CoverTab[7065]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:156
			_go_fuzz_dep_.CoverTab[7066]++
//line /usr/local/go/src/net/lookup_unix.go:156
			// _ = "end of CoverTab[7066]"
//line /usr/local/go/src/net/lookup_unix.go:156
		}
//line /usr/local/go/src/net/lookup_unix.go:156
		// _ = "end of CoverTab[7064]"
//line /usr/local/go/src/net/lookup_unix.go:156
	}
//line /usr/local/go/src/net/lookup_unix.go:156
	// _ = "end of CoverTab[7059]"
//line /usr/local/go/src/net/lookup_unix.go:156
	_go_fuzz_dep_.CoverTab[7060]++
							return r
//line /usr/local/go/src/net/lookup_unix.go:157
	// _ = "end of CoverTab[7060]"
}

//line /usr/local/go/src/net/lookup_unix.go:158
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/lookup_unix.go:158
var _ = _go_fuzz_dep_.CoverTab
