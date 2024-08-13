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
	_go_fuzz_dep_.CoverTab[15380]++
						file, err := open("/etc/protocols")
						if err != nil {
//line /usr/local/go/src/net/lookup_unix.go:22
		_go_fuzz_dep_.CoverTab[15382]++
							return
//line /usr/local/go/src/net/lookup_unix.go:23
		// _ = "end of CoverTab[15382]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:24
		_go_fuzz_dep_.CoverTab[15383]++
//line /usr/local/go/src/net/lookup_unix.go:24
		// _ = "end of CoverTab[15383]"
//line /usr/local/go/src/net/lookup_unix.go:24
	}
//line /usr/local/go/src/net/lookup_unix.go:24
	// _ = "end of CoverTab[15380]"
//line /usr/local/go/src/net/lookup_unix.go:24
	_go_fuzz_dep_.CoverTab[15381]++
						defer file.close()

						for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /usr/local/go/src/net/lookup_unix.go:27
		_go_fuzz_dep_.CoverTab[15384]++

							if i := bytealg.IndexByteString(line, '#'); i >= 0 {
//line /usr/local/go/src/net/lookup_unix.go:29
			_go_fuzz_dep_.CoverTab[15387]++
								line = line[0:i]
//line /usr/local/go/src/net/lookup_unix.go:30
			// _ = "end of CoverTab[15387]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:31
			_go_fuzz_dep_.CoverTab[15388]++
//line /usr/local/go/src/net/lookup_unix.go:31
			// _ = "end of CoverTab[15388]"
//line /usr/local/go/src/net/lookup_unix.go:31
		}
//line /usr/local/go/src/net/lookup_unix.go:31
		// _ = "end of CoverTab[15384]"
//line /usr/local/go/src/net/lookup_unix.go:31
		_go_fuzz_dep_.CoverTab[15385]++
							f := getFields(line)
							if len(f) < 2 {
//line /usr/local/go/src/net/lookup_unix.go:33
			_go_fuzz_dep_.CoverTab[15389]++
								continue
//line /usr/local/go/src/net/lookup_unix.go:34
			// _ = "end of CoverTab[15389]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:35
			_go_fuzz_dep_.CoverTab[15390]++
//line /usr/local/go/src/net/lookup_unix.go:35
			// _ = "end of CoverTab[15390]"
//line /usr/local/go/src/net/lookup_unix.go:35
		}
//line /usr/local/go/src/net/lookup_unix.go:35
		// _ = "end of CoverTab[15385]"
//line /usr/local/go/src/net/lookup_unix.go:35
		_go_fuzz_dep_.CoverTab[15386]++
							if proto, _, ok := dtoi(f[1]); ok {
//line /usr/local/go/src/net/lookup_unix.go:36
			_go_fuzz_dep_.CoverTab[15391]++
								if _, ok := protocols[f[0]]; !ok {
//line /usr/local/go/src/net/lookup_unix.go:37
				_go_fuzz_dep_.CoverTab[15393]++
									protocols[f[0]] = proto
//line /usr/local/go/src/net/lookup_unix.go:38
				// _ = "end of CoverTab[15393]"
			} else {
//line /usr/local/go/src/net/lookup_unix.go:39
				_go_fuzz_dep_.CoverTab[15394]++
//line /usr/local/go/src/net/lookup_unix.go:39
				// _ = "end of CoverTab[15394]"
//line /usr/local/go/src/net/lookup_unix.go:39
			}
//line /usr/local/go/src/net/lookup_unix.go:39
			// _ = "end of CoverTab[15391]"
//line /usr/local/go/src/net/lookup_unix.go:39
			_go_fuzz_dep_.CoverTab[15392]++
								for _, alias := range f[2:] {
//line /usr/local/go/src/net/lookup_unix.go:40
				_go_fuzz_dep_.CoverTab[15395]++
									if _, ok := protocols[alias]; !ok {
//line /usr/local/go/src/net/lookup_unix.go:41
					_go_fuzz_dep_.CoverTab[15396]++
										protocols[alias] = proto
//line /usr/local/go/src/net/lookup_unix.go:42
					// _ = "end of CoverTab[15396]"
				} else {
//line /usr/local/go/src/net/lookup_unix.go:43
					_go_fuzz_dep_.CoverTab[15397]++
//line /usr/local/go/src/net/lookup_unix.go:43
					// _ = "end of CoverTab[15397]"
//line /usr/local/go/src/net/lookup_unix.go:43
				}
//line /usr/local/go/src/net/lookup_unix.go:43
				// _ = "end of CoverTab[15395]"
			}
//line /usr/local/go/src/net/lookup_unix.go:44
			// _ = "end of CoverTab[15392]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:45
			_go_fuzz_dep_.CoverTab[15398]++
//line /usr/local/go/src/net/lookup_unix.go:45
			// _ = "end of CoverTab[15398]"
//line /usr/local/go/src/net/lookup_unix.go:45
		}
//line /usr/local/go/src/net/lookup_unix.go:45
		// _ = "end of CoverTab[15386]"
	}
//line /usr/local/go/src/net/lookup_unix.go:46
	// _ = "end of CoverTab[15381]"
}

// lookupProtocol looks up IP protocol name in /etc/protocols and
//line /usr/local/go/src/net/lookup_unix.go:49
// returns correspondent protocol number.
//line /usr/local/go/src/net/lookup_unix.go:51
func lookupProtocol(_ context.Context, name string) (int, error) {
//line /usr/local/go/src/net/lookup_unix.go:51
	_go_fuzz_dep_.CoverTab[15399]++
						onceReadProtocols.Do(readProtocols)
						return lookupProtocolMap(name)
//line /usr/local/go/src/net/lookup_unix.go:53
	// _ = "end of CoverTab[15399]"
}

func (r *Resolver) lookupHost(ctx context.Context, host string) (addrs []string, err error) {
//line /usr/local/go/src/net/lookup_unix.go:56
	_go_fuzz_dep_.CoverTab[15400]++
						order, conf := systemConf().hostLookupOrder(r, host)
						if !r.preferGo() && func() bool {
//line /usr/local/go/src/net/lookup_unix.go:58
		_go_fuzz_dep_.CoverTab[15402]++
//line /usr/local/go/src/net/lookup_unix.go:58
		return order == hostLookupCgo
//line /usr/local/go/src/net/lookup_unix.go:58
		// _ = "end of CoverTab[15402]"
//line /usr/local/go/src/net/lookup_unix.go:58
	}() {
//line /usr/local/go/src/net/lookup_unix.go:58
		_go_fuzz_dep_.CoverTab[15403]++
							if addrs, err, ok := cgoLookupHost(ctx, host); ok {
//line /usr/local/go/src/net/lookup_unix.go:59
			_go_fuzz_dep_.CoverTab[15405]++
								return addrs, err
//line /usr/local/go/src/net/lookup_unix.go:60
			// _ = "end of CoverTab[15405]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:61
			_go_fuzz_dep_.CoverTab[15406]++
//line /usr/local/go/src/net/lookup_unix.go:61
			// _ = "end of CoverTab[15406]"
//line /usr/local/go/src/net/lookup_unix.go:61
		}
//line /usr/local/go/src/net/lookup_unix.go:61
		// _ = "end of CoverTab[15403]"
//line /usr/local/go/src/net/lookup_unix.go:61
		_go_fuzz_dep_.CoverTab[15404]++

							order = hostLookupFilesDNS
//line /usr/local/go/src/net/lookup_unix.go:63
		// _ = "end of CoverTab[15404]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:64
		_go_fuzz_dep_.CoverTab[15407]++
//line /usr/local/go/src/net/lookup_unix.go:64
		// _ = "end of CoverTab[15407]"
//line /usr/local/go/src/net/lookup_unix.go:64
	}
//line /usr/local/go/src/net/lookup_unix.go:64
	// _ = "end of CoverTab[15400]"
//line /usr/local/go/src/net/lookup_unix.go:64
	_go_fuzz_dep_.CoverTab[15401]++
						return r.goLookupHostOrder(ctx, host, order, conf)
//line /usr/local/go/src/net/lookup_unix.go:65
	// _ = "end of CoverTab[15401]"
}

func (r *Resolver) lookupIP(ctx context.Context, network, host string) (addrs []IPAddr, err error) {
//line /usr/local/go/src/net/lookup_unix.go:68
	_go_fuzz_dep_.CoverTab[15408]++
						if r.preferGo() {
//line /usr/local/go/src/net/lookup_unix.go:69
		_go_fuzz_dep_.CoverTab[15411]++
							return r.goLookupIP(ctx, network, host)
//line /usr/local/go/src/net/lookup_unix.go:70
		// _ = "end of CoverTab[15411]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:71
		_go_fuzz_dep_.CoverTab[15412]++
//line /usr/local/go/src/net/lookup_unix.go:71
		// _ = "end of CoverTab[15412]"
//line /usr/local/go/src/net/lookup_unix.go:71
	}
//line /usr/local/go/src/net/lookup_unix.go:71
	// _ = "end of CoverTab[15408]"
//line /usr/local/go/src/net/lookup_unix.go:71
	_go_fuzz_dep_.CoverTab[15409]++
						order, conf := systemConf().hostLookupOrder(r, host)
						if order == hostLookupCgo {
//line /usr/local/go/src/net/lookup_unix.go:73
		_go_fuzz_dep_.CoverTab[15413]++
							if addrs, err, ok := cgoLookupIP(ctx, network, host); ok {
//line /usr/local/go/src/net/lookup_unix.go:74
			_go_fuzz_dep_.CoverTab[15415]++
								return addrs, err
//line /usr/local/go/src/net/lookup_unix.go:75
			// _ = "end of CoverTab[15415]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:76
			_go_fuzz_dep_.CoverTab[15416]++
//line /usr/local/go/src/net/lookup_unix.go:76
			// _ = "end of CoverTab[15416]"
//line /usr/local/go/src/net/lookup_unix.go:76
		}
//line /usr/local/go/src/net/lookup_unix.go:76
		// _ = "end of CoverTab[15413]"
//line /usr/local/go/src/net/lookup_unix.go:76
		_go_fuzz_dep_.CoverTab[15414]++

							order = hostLookupFilesDNS
//line /usr/local/go/src/net/lookup_unix.go:78
		// _ = "end of CoverTab[15414]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:79
		_go_fuzz_dep_.CoverTab[15417]++
//line /usr/local/go/src/net/lookup_unix.go:79
		// _ = "end of CoverTab[15417]"
//line /usr/local/go/src/net/lookup_unix.go:79
	}
//line /usr/local/go/src/net/lookup_unix.go:79
	// _ = "end of CoverTab[15409]"
//line /usr/local/go/src/net/lookup_unix.go:79
	_go_fuzz_dep_.CoverTab[15410]++
						ips, _, err := r.goLookupIPCNAMEOrder(ctx, network, host, order, conf)
						return ips, err
//line /usr/local/go/src/net/lookup_unix.go:81
	// _ = "end of CoverTab[15410]"
}

func (r *Resolver) lookupPort(ctx context.Context, network, service string) (int, error) {
//line /usr/local/go/src/net/lookup_unix.go:84
	_go_fuzz_dep_.CoverTab[15418]++
						if !r.preferGo() && func() bool {
//line /usr/local/go/src/net/lookup_unix.go:85
		_go_fuzz_dep_.CoverTab[15420]++
//line /usr/local/go/src/net/lookup_unix.go:85
		return systemConf().canUseCgo()
//line /usr/local/go/src/net/lookup_unix.go:85
		// _ = "end of CoverTab[15420]"
//line /usr/local/go/src/net/lookup_unix.go:85
	}() {
//line /usr/local/go/src/net/lookup_unix.go:85
		_go_fuzz_dep_.CoverTab[15421]++
							if port, err, ok := cgoLookupPort(ctx, network, service); ok {
//line /usr/local/go/src/net/lookup_unix.go:86
			_go_fuzz_dep_.CoverTab[15422]++
								if err != nil {
//line /usr/local/go/src/net/lookup_unix.go:87
				_go_fuzz_dep_.CoverTab[15424]++

//line /usr/local/go/src/net/lookup_unix.go:90
				if port, err := goLookupPort(network, service); err == nil {
//line /usr/local/go/src/net/lookup_unix.go:90
					_go_fuzz_dep_.CoverTab[15425]++
										return port, nil
//line /usr/local/go/src/net/lookup_unix.go:91
					// _ = "end of CoverTab[15425]"
				} else {
//line /usr/local/go/src/net/lookup_unix.go:92
					_go_fuzz_dep_.CoverTab[15426]++
//line /usr/local/go/src/net/lookup_unix.go:92
					// _ = "end of CoverTab[15426]"
//line /usr/local/go/src/net/lookup_unix.go:92
				}
//line /usr/local/go/src/net/lookup_unix.go:92
				// _ = "end of CoverTab[15424]"
			} else {
//line /usr/local/go/src/net/lookup_unix.go:93
				_go_fuzz_dep_.CoverTab[15427]++
//line /usr/local/go/src/net/lookup_unix.go:93
				// _ = "end of CoverTab[15427]"
//line /usr/local/go/src/net/lookup_unix.go:93
			}
//line /usr/local/go/src/net/lookup_unix.go:93
			// _ = "end of CoverTab[15422]"
//line /usr/local/go/src/net/lookup_unix.go:93
			_go_fuzz_dep_.CoverTab[15423]++
								return port, err
//line /usr/local/go/src/net/lookup_unix.go:94
			// _ = "end of CoverTab[15423]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:95
			_go_fuzz_dep_.CoverTab[15428]++
//line /usr/local/go/src/net/lookup_unix.go:95
			// _ = "end of CoverTab[15428]"
//line /usr/local/go/src/net/lookup_unix.go:95
		}
//line /usr/local/go/src/net/lookup_unix.go:95
		// _ = "end of CoverTab[15421]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:96
		_go_fuzz_dep_.CoverTab[15429]++
//line /usr/local/go/src/net/lookup_unix.go:96
		// _ = "end of CoverTab[15429]"
//line /usr/local/go/src/net/lookup_unix.go:96
	}
//line /usr/local/go/src/net/lookup_unix.go:96
	// _ = "end of CoverTab[15418]"
//line /usr/local/go/src/net/lookup_unix.go:96
	_go_fuzz_dep_.CoverTab[15419]++
						return goLookupPort(network, service)
//line /usr/local/go/src/net/lookup_unix.go:97
	// _ = "end of CoverTab[15419]"
}

func (r *Resolver) lookupCNAME(ctx context.Context, name string) (string, error) {
//line /usr/local/go/src/net/lookup_unix.go:100
	_go_fuzz_dep_.CoverTab[15430]++
							order, conf := systemConf().hostLookupOrder(r, name)
							if !r.preferGo() && func() bool {
//line /usr/local/go/src/net/lookup_unix.go:102
		_go_fuzz_dep_.CoverTab[15432]++
//line /usr/local/go/src/net/lookup_unix.go:102
		return order == hostLookupCgo
//line /usr/local/go/src/net/lookup_unix.go:102
		// _ = "end of CoverTab[15432]"
//line /usr/local/go/src/net/lookup_unix.go:102
	}() {
//line /usr/local/go/src/net/lookup_unix.go:102
		_go_fuzz_dep_.CoverTab[15433]++
								if cname, err, ok := cgoLookupCNAME(ctx, name); ok {
//line /usr/local/go/src/net/lookup_unix.go:103
			_go_fuzz_dep_.CoverTab[15434]++
									return cname, err
//line /usr/local/go/src/net/lookup_unix.go:104
			// _ = "end of CoverTab[15434]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:105
			_go_fuzz_dep_.CoverTab[15435]++
//line /usr/local/go/src/net/lookup_unix.go:105
			// _ = "end of CoverTab[15435]"
//line /usr/local/go/src/net/lookup_unix.go:105
		}
//line /usr/local/go/src/net/lookup_unix.go:105
		// _ = "end of CoverTab[15433]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:106
		_go_fuzz_dep_.CoverTab[15436]++
//line /usr/local/go/src/net/lookup_unix.go:106
		// _ = "end of CoverTab[15436]"
//line /usr/local/go/src/net/lookup_unix.go:106
	}
//line /usr/local/go/src/net/lookup_unix.go:106
	// _ = "end of CoverTab[15430]"
//line /usr/local/go/src/net/lookup_unix.go:106
	_go_fuzz_dep_.CoverTab[15431]++
							return r.goLookupCNAME(ctx, name, order, conf)
//line /usr/local/go/src/net/lookup_unix.go:107
	// _ = "end of CoverTab[15431]"
}

func (r *Resolver) lookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error) {
//line /usr/local/go/src/net/lookup_unix.go:110
	_go_fuzz_dep_.CoverTab[15437]++
							return r.goLookupSRV(ctx, service, proto, name)
//line /usr/local/go/src/net/lookup_unix.go:111
	// _ = "end of CoverTab[15437]"
}

func (r *Resolver) lookupMX(ctx context.Context, name string) ([]*MX, error) {
//line /usr/local/go/src/net/lookup_unix.go:114
	_go_fuzz_dep_.CoverTab[15438]++
							return r.goLookupMX(ctx, name)
//line /usr/local/go/src/net/lookup_unix.go:115
	// _ = "end of CoverTab[15438]"
}

func (r *Resolver) lookupNS(ctx context.Context, name string) ([]*NS, error) {
//line /usr/local/go/src/net/lookup_unix.go:118
	_go_fuzz_dep_.CoverTab[15439]++
							return r.goLookupNS(ctx, name)
//line /usr/local/go/src/net/lookup_unix.go:119
	// _ = "end of CoverTab[15439]"
}

func (r *Resolver) lookupTXT(ctx context.Context, name string) ([]string, error) {
//line /usr/local/go/src/net/lookup_unix.go:122
	_go_fuzz_dep_.CoverTab[15440]++
							return r.goLookupTXT(ctx, name)
//line /usr/local/go/src/net/lookup_unix.go:123
	// _ = "end of CoverTab[15440]"
}

func (r *Resolver) lookupAddr(ctx context.Context, addr string) ([]string, error) {
//line /usr/local/go/src/net/lookup_unix.go:126
	_go_fuzz_dep_.CoverTab[15441]++
							order, conf := systemConf().hostLookupOrder(r, "")
							if !r.preferGo() && func() bool {
//line /usr/local/go/src/net/lookup_unix.go:128
		_go_fuzz_dep_.CoverTab[15443]++
//line /usr/local/go/src/net/lookup_unix.go:128
		return order == hostLookupCgo
//line /usr/local/go/src/net/lookup_unix.go:128
		// _ = "end of CoverTab[15443]"
//line /usr/local/go/src/net/lookup_unix.go:128
	}() {
//line /usr/local/go/src/net/lookup_unix.go:128
		_go_fuzz_dep_.CoverTab[15444]++
								if ptrs, err, ok := cgoLookupPTR(ctx, addr); ok {
//line /usr/local/go/src/net/lookup_unix.go:129
			_go_fuzz_dep_.CoverTab[15445]++
									return ptrs, err
//line /usr/local/go/src/net/lookup_unix.go:130
			// _ = "end of CoverTab[15445]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:131
			_go_fuzz_dep_.CoverTab[15446]++
//line /usr/local/go/src/net/lookup_unix.go:131
			// _ = "end of CoverTab[15446]"
//line /usr/local/go/src/net/lookup_unix.go:131
		}
//line /usr/local/go/src/net/lookup_unix.go:131
		// _ = "end of CoverTab[15444]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:132
		_go_fuzz_dep_.CoverTab[15447]++
//line /usr/local/go/src/net/lookup_unix.go:132
		// _ = "end of CoverTab[15447]"
//line /usr/local/go/src/net/lookup_unix.go:132
	}
//line /usr/local/go/src/net/lookup_unix.go:132
	// _ = "end of CoverTab[15441]"
//line /usr/local/go/src/net/lookup_unix.go:132
	_go_fuzz_dep_.CoverTab[15442]++
							return r.goLookupPTR(ctx, addr, conf)
//line /usr/local/go/src/net/lookup_unix.go:133
	// _ = "end of CoverTab[15442]"
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
	_go_fuzz_dep_.CoverTab[15448]++
							var rlim syscall.Rlimit
							if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim); err != nil {
//line /usr/local/go/src/net/lookup_unix.go:148
		_go_fuzz_dep_.CoverTab[15451]++
								return 500
//line /usr/local/go/src/net/lookup_unix.go:149
		// _ = "end of CoverTab[15451]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:150
		_go_fuzz_dep_.CoverTab[15452]++
//line /usr/local/go/src/net/lookup_unix.go:150
		// _ = "end of CoverTab[15452]"
//line /usr/local/go/src/net/lookup_unix.go:150
	}
//line /usr/local/go/src/net/lookup_unix.go:150
	// _ = "end of CoverTab[15448]"
//line /usr/local/go/src/net/lookup_unix.go:150
	_go_fuzz_dep_.CoverTab[15449]++
							r := int(rlim.Cur)
							if r > 500 {
//line /usr/local/go/src/net/lookup_unix.go:152
		_go_fuzz_dep_.CoverTab[15453]++
								r = 500
//line /usr/local/go/src/net/lookup_unix.go:153
		// _ = "end of CoverTab[15453]"
	} else {
//line /usr/local/go/src/net/lookup_unix.go:154
		_go_fuzz_dep_.CoverTab[15454]++
//line /usr/local/go/src/net/lookup_unix.go:154
		if r > 30 {
//line /usr/local/go/src/net/lookup_unix.go:154
			_go_fuzz_dep_.CoverTab[15455]++
									r -= 30
//line /usr/local/go/src/net/lookup_unix.go:155
			// _ = "end of CoverTab[15455]"
		} else {
//line /usr/local/go/src/net/lookup_unix.go:156
			_go_fuzz_dep_.CoverTab[15456]++
//line /usr/local/go/src/net/lookup_unix.go:156
			// _ = "end of CoverTab[15456]"
//line /usr/local/go/src/net/lookup_unix.go:156
		}
//line /usr/local/go/src/net/lookup_unix.go:156
		// _ = "end of CoverTab[15454]"
//line /usr/local/go/src/net/lookup_unix.go:156
	}
//line /usr/local/go/src/net/lookup_unix.go:156
	// _ = "end of CoverTab[15449]"
//line /usr/local/go/src/net/lookup_unix.go:156
	_go_fuzz_dep_.CoverTab[15450]++
							return r
//line /usr/local/go/src/net/lookup_unix.go:157
	// _ = "end of CoverTab[15450]"
}

//line /usr/local/go/src/net/lookup_unix.go:158
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/lookup_unix.go:158
var _ = _go_fuzz_dep_.CoverTab
