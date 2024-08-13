// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm)

// Read system port mappings from /etc/services

//line /usr/local/go/src/net/port_unix.go:9
package net

//line /usr/local/go/src/net/port_unix.go:9
import (
//line /usr/local/go/src/net/port_unix.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/port_unix.go:9
)
//line /usr/local/go/src/net/port_unix.go:9
import (
//line /usr/local/go/src/net/port_unix.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/port_unix.go:9
)

import (
	"internal/bytealg"
	"sync"
)

var onceReadServices sync.Once

func readServices() {
//line /usr/local/go/src/net/port_unix.go:18
	_go_fuzz_dep_.CoverTab[7650]++
						file, err := open("/etc/services")
						if err != nil {
//line /usr/local/go/src/net/port_unix.go:20
		_go_fuzz_dep_.CoverTab[7652]++
							return
//line /usr/local/go/src/net/port_unix.go:21
		// _ = "end of CoverTab[7652]"
	} else {
//line /usr/local/go/src/net/port_unix.go:22
		_go_fuzz_dep_.CoverTab[7653]++
//line /usr/local/go/src/net/port_unix.go:22
		// _ = "end of CoverTab[7653]"
//line /usr/local/go/src/net/port_unix.go:22
	}
//line /usr/local/go/src/net/port_unix.go:22
	// _ = "end of CoverTab[7650]"
//line /usr/local/go/src/net/port_unix.go:22
	_go_fuzz_dep_.CoverTab[7651]++
						defer file.close()

						for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /usr/local/go/src/net/port_unix.go:25
		_go_fuzz_dep_.CoverTab[7654]++

							if i := bytealg.IndexByteString(line, '#'); i >= 0 {
//line /usr/local/go/src/net/port_unix.go:27
			_go_fuzz_dep_.CoverTab[7659]++
								line = line[:i]
//line /usr/local/go/src/net/port_unix.go:28
			// _ = "end of CoverTab[7659]"
		} else {
//line /usr/local/go/src/net/port_unix.go:29
			_go_fuzz_dep_.CoverTab[7660]++
//line /usr/local/go/src/net/port_unix.go:29
			// _ = "end of CoverTab[7660]"
//line /usr/local/go/src/net/port_unix.go:29
		}
//line /usr/local/go/src/net/port_unix.go:29
		// _ = "end of CoverTab[7654]"
//line /usr/local/go/src/net/port_unix.go:29
		_go_fuzz_dep_.CoverTab[7655]++
							f := getFields(line)
							if len(f) < 2 {
//line /usr/local/go/src/net/port_unix.go:31
			_go_fuzz_dep_.CoverTab[7661]++
								continue
//line /usr/local/go/src/net/port_unix.go:32
			// _ = "end of CoverTab[7661]"
		} else {
//line /usr/local/go/src/net/port_unix.go:33
			_go_fuzz_dep_.CoverTab[7662]++
//line /usr/local/go/src/net/port_unix.go:33
			// _ = "end of CoverTab[7662]"
//line /usr/local/go/src/net/port_unix.go:33
		}
//line /usr/local/go/src/net/port_unix.go:33
		// _ = "end of CoverTab[7655]"
//line /usr/local/go/src/net/port_unix.go:33
		_go_fuzz_dep_.CoverTab[7656]++
							portnet := f[1]
							port, j, ok := dtoi(portnet)
							if !ok || func() bool {
//line /usr/local/go/src/net/port_unix.go:36
			_go_fuzz_dep_.CoverTab[7663]++
//line /usr/local/go/src/net/port_unix.go:36
			return port <= 0
//line /usr/local/go/src/net/port_unix.go:36
			// _ = "end of CoverTab[7663]"
//line /usr/local/go/src/net/port_unix.go:36
		}() || func() bool {
//line /usr/local/go/src/net/port_unix.go:36
			_go_fuzz_dep_.CoverTab[7664]++
//line /usr/local/go/src/net/port_unix.go:36
			return j >= len(portnet)
//line /usr/local/go/src/net/port_unix.go:36
			// _ = "end of CoverTab[7664]"
//line /usr/local/go/src/net/port_unix.go:36
		}() || func() bool {
//line /usr/local/go/src/net/port_unix.go:36
			_go_fuzz_dep_.CoverTab[7665]++
//line /usr/local/go/src/net/port_unix.go:36
			return portnet[j] != '/'
//line /usr/local/go/src/net/port_unix.go:36
			// _ = "end of CoverTab[7665]"
//line /usr/local/go/src/net/port_unix.go:36
		}() {
//line /usr/local/go/src/net/port_unix.go:36
			_go_fuzz_dep_.CoverTab[7666]++
								continue
//line /usr/local/go/src/net/port_unix.go:37
			// _ = "end of CoverTab[7666]"
		} else {
//line /usr/local/go/src/net/port_unix.go:38
			_go_fuzz_dep_.CoverTab[7667]++
//line /usr/local/go/src/net/port_unix.go:38
			// _ = "end of CoverTab[7667]"
//line /usr/local/go/src/net/port_unix.go:38
		}
//line /usr/local/go/src/net/port_unix.go:38
		// _ = "end of CoverTab[7656]"
//line /usr/local/go/src/net/port_unix.go:38
		_go_fuzz_dep_.CoverTab[7657]++
							netw := portnet[j+1:]
							m, ok1 := services[netw]
							if !ok1 {
//line /usr/local/go/src/net/port_unix.go:41
			_go_fuzz_dep_.CoverTab[7668]++
								m = make(map[string]int)
								services[netw] = m
//line /usr/local/go/src/net/port_unix.go:43
			// _ = "end of CoverTab[7668]"
		} else {
//line /usr/local/go/src/net/port_unix.go:44
			_go_fuzz_dep_.CoverTab[7669]++
//line /usr/local/go/src/net/port_unix.go:44
			// _ = "end of CoverTab[7669]"
//line /usr/local/go/src/net/port_unix.go:44
		}
//line /usr/local/go/src/net/port_unix.go:44
		// _ = "end of CoverTab[7657]"
//line /usr/local/go/src/net/port_unix.go:44
		_go_fuzz_dep_.CoverTab[7658]++
							for i := 0; i < len(f); i++ {
//line /usr/local/go/src/net/port_unix.go:45
			_go_fuzz_dep_.CoverTab[7670]++
								if i != 1 {
//line /usr/local/go/src/net/port_unix.go:46
				_go_fuzz_dep_.CoverTab[7671]++
									m[f[i]] = port
//line /usr/local/go/src/net/port_unix.go:47
				// _ = "end of CoverTab[7671]"
			} else {
//line /usr/local/go/src/net/port_unix.go:48
				_go_fuzz_dep_.CoverTab[7672]++
//line /usr/local/go/src/net/port_unix.go:48
				// _ = "end of CoverTab[7672]"
//line /usr/local/go/src/net/port_unix.go:48
			}
//line /usr/local/go/src/net/port_unix.go:48
			// _ = "end of CoverTab[7670]"
		}
//line /usr/local/go/src/net/port_unix.go:49
		// _ = "end of CoverTab[7658]"
	}
//line /usr/local/go/src/net/port_unix.go:50
	// _ = "end of CoverTab[7651]"
}

// goLookupPort is the native Go implementation of LookupPort.
func goLookupPort(network, service string) (port int, err error) {
//line /usr/local/go/src/net/port_unix.go:54
	_go_fuzz_dep_.CoverTab[7673]++
						onceReadServices.Do(readServices)
						return lookupPortMap(network, service)
//line /usr/local/go/src/net/port_unix.go:56
	// _ = "end of CoverTab[7673]"
}

//line /usr/local/go/src/net/port_unix.go:57
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/port_unix.go:57
var _ = _go_fuzz_dep_.CoverTab
