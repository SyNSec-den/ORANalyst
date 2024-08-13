// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

// Read system port mappings from /etc/services

//line /snap/go/10455/src/net/port_unix.go:9
package net

//line /snap/go/10455/src/net/port_unix.go:9
import (
//line /snap/go/10455/src/net/port_unix.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/port_unix.go:9
)
//line /snap/go/10455/src/net/port_unix.go:9
import (
//line /snap/go/10455/src/net/port_unix.go:9
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/port_unix.go:9
)

import (
	"internal/bytealg"
	"sync"
)

var onceReadServices sync.Once

func readServices() {
//line /snap/go/10455/src/net/port_unix.go:18
	_go_fuzz_dep_.CoverTab[7930]++
						file, err := open("/etc/services")
						if err != nil {
//line /snap/go/10455/src/net/port_unix.go:20
		_go_fuzz_dep_.CoverTab[529562]++
//line /snap/go/10455/src/net/port_unix.go:20
		_go_fuzz_dep_.CoverTab[7932]++
							return
//line /snap/go/10455/src/net/port_unix.go:21
		// _ = "end of CoverTab[7932]"
	} else {
//line /snap/go/10455/src/net/port_unix.go:22
		_go_fuzz_dep_.CoverTab[529563]++
//line /snap/go/10455/src/net/port_unix.go:22
		_go_fuzz_dep_.CoverTab[7933]++
//line /snap/go/10455/src/net/port_unix.go:22
		// _ = "end of CoverTab[7933]"
//line /snap/go/10455/src/net/port_unix.go:22
	}
//line /snap/go/10455/src/net/port_unix.go:22
	// _ = "end of CoverTab[7930]"
//line /snap/go/10455/src/net/port_unix.go:22
	_go_fuzz_dep_.CoverTab[7931]++
						defer file.close()
//line /snap/go/10455/src/net/port_unix.go:23
	_go_fuzz_dep_.CoverTab[786742] = 0

						for line, ok := file.readLine(); ok; line, ok = file.readLine() {
//line /snap/go/10455/src/net/port_unix.go:25
		if _go_fuzz_dep_.CoverTab[786742] == 0 {
//line /snap/go/10455/src/net/port_unix.go:25
			_go_fuzz_dep_.CoverTab[529574]++
//line /snap/go/10455/src/net/port_unix.go:25
		} else {
//line /snap/go/10455/src/net/port_unix.go:25
			_go_fuzz_dep_.CoverTab[529575]++
//line /snap/go/10455/src/net/port_unix.go:25
		}
//line /snap/go/10455/src/net/port_unix.go:25
		_go_fuzz_dep_.CoverTab[786742] = 1
//line /snap/go/10455/src/net/port_unix.go:25
		_go_fuzz_dep_.CoverTab[7934]++

							if i := bytealg.IndexByteString(line, '#'); i >= 0 {
//line /snap/go/10455/src/net/port_unix.go:27
			_go_fuzz_dep_.CoverTab[529564]++
//line /snap/go/10455/src/net/port_unix.go:27
			_go_fuzz_dep_.CoverTab[7939]++
								line = line[:i]
//line /snap/go/10455/src/net/port_unix.go:28
			// _ = "end of CoverTab[7939]"
		} else {
//line /snap/go/10455/src/net/port_unix.go:29
			_go_fuzz_dep_.CoverTab[529565]++
//line /snap/go/10455/src/net/port_unix.go:29
			_go_fuzz_dep_.CoverTab[7940]++
//line /snap/go/10455/src/net/port_unix.go:29
			// _ = "end of CoverTab[7940]"
//line /snap/go/10455/src/net/port_unix.go:29
		}
//line /snap/go/10455/src/net/port_unix.go:29
		// _ = "end of CoverTab[7934]"
//line /snap/go/10455/src/net/port_unix.go:29
		_go_fuzz_dep_.CoverTab[7935]++
							f := getFields(line)
							if len(f) < 2 {
//line /snap/go/10455/src/net/port_unix.go:31
			_go_fuzz_dep_.CoverTab[529566]++
//line /snap/go/10455/src/net/port_unix.go:31
			_go_fuzz_dep_.CoverTab[7941]++
								continue
//line /snap/go/10455/src/net/port_unix.go:32
			// _ = "end of CoverTab[7941]"
		} else {
//line /snap/go/10455/src/net/port_unix.go:33
			_go_fuzz_dep_.CoverTab[529567]++
//line /snap/go/10455/src/net/port_unix.go:33
			_go_fuzz_dep_.CoverTab[7942]++
//line /snap/go/10455/src/net/port_unix.go:33
			// _ = "end of CoverTab[7942]"
//line /snap/go/10455/src/net/port_unix.go:33
		}
//line /snap/go/10455/src/net/port_unix.go:33
		// _ = "end of CoverTab[7935]"
//line /snap/go/10455/src/net/port_unix.go:33
		_go_fuzz_dep_.CoverTab[7936]++
							portnet := f[1]
							port, j, ok := dtoi(portnet)
							if !ok || func() bool {
//line /snap/go/10455/src/net/port_unix.go:36
			_go_fuzz_dep_.CoverTab[7943]++
//line /snap/go/10455/src/net/port_unix.go:36
			return port <= 0
//line /snap/go/10455/src/net/port_unix.go:36
			// _ = "end of CoverTab[7943]"
//line /snap/go/10455/src/net/port_unix.go:36
		}() || func() bool {
//line /snap/go/10455/src/net/port_unix.go:36
			_go_fuzz_dep_.CoverTab[7944]++
//line /snap/go/10455/src/net/port_unix.go:36
			return j >= len(portnet)
//line /snap/go/10455/src/net/port_unix.go:36
			// _ = "end of CoverTab[7944]"
//line /snap/go/10455/src/net/port_unix.go:36
		}() || func() bool {
//line /snap/go/10455/src/net/port_unix.go:36
			_go_fuzz_dep_.CoverTab[7945]++
//line /snap/go/10455/src/net/port_unix.go:36
			return portnet[j] != '/'
//line /snap/go/10455/src/net/port_unix.go:36
			// _ = "end of CoverTab[7945]"
//line /snap/go/10455/src/net/port_unix.go:36
		}() {
//line /snap/go/10455/src/net/port_unix.go:36
			_go_fuzz_dep_.CoverTab[529568]++
//line /snap/go/10455/src/net/port_unix.go:36
			_go_fuzz_dep_.CoverTab[7946]++
								continue
//line /snap/go/10455/src/net/port_unix.go:37
			// _ = "end of CoverTab[7946]"
		} else {
//line /snap/go/10455/src/net/port_unix.go:38
			_go_fuzz_dep_.CoverTab[529569]++
//line /snap/go/10455/src/net/port_unix.go:38
			_go_fuzz_dep_.CoverTab[7947]++
//line /snap/go/10455/src/net/port_unix.go:38
			// _ = "end of CoverTab[7947]"
//line /snap/go/10455/src/net/port_unix.go:38
		}
//line /snap/go/10455/src/net/port_unix.go:38
		// _ = "end of CoverTab[7936]"
//line /snap/go/10455/src/net/port_unix.go:38
		_go_fuzz_dep_.CoverTab[7937]++
							netw := portnet[j+1:]
							m, ok1 := services[netw]
							if !ok1 {
//line /snap/go/10455/src/net/port_unix.go:41
			_go_fuzz_dep_.CoverTab[529570]++
//line /snap/go/10455/src/net/port_unix.go:41
			_go_fuzz_dep_.CoverTab[7948]++
								m = make(map[string]int)
								services[netw] = m
//line /snap/go/10455/src/net/port_unix.go:43
			// _ = "end of CoverTab[7948]"
		} else {
//line /snap/go/10455/src/net/port_unix.go:44
			_go_fuzz_dep_.CoverTab[529571]++
//line /snap/go/10455/src/net/port_unix.go:44
			_go_fuzz_dep_.CoverTab[7949]++
//line /snap/go/10455/src/net/port_unix.go:44
			// _ = "end of CoverTab[7949]"
//line /snap/go/10455/src/net/port_unix.go:44
		}
//line /snap/go/10455/src/net/port_unix.go:44
		// _ = "end of CoverTab[7937]"
//line /snap/go/10455/src/net/port_unix.go:44
		_go_fuzz_dep_.CoverTab[7938]++
//line /snap/go/10455/src/net/port_unix.go:44
		_go_fuzz_dep_.CoverTab[786743] = 0
							for i := 0; i < len(f); i++ {
//line /snap/go/10455/src/net/port_unix.go:45
			if _go_fuzz_dep_.CoverTab[786743] == 0 {
//line /snap/go/10455/src/net/port_unix.go:45
				_go_fuzz_dep_.CoverTab[529578]++
//line /snap/go/10455/src/net/port_unix.go:45
			} else {
//line /snap/go/10455/src/net/port_unix.go:45
				_go_fuzz_dep_.CoverTab[529579]++
//line /snap/go/10455/src/net/port_unix.go:45
			}
//line /snap/go/10455/src/net/port_unix.go:45
			_go_fuzz_dep_.CoverTab[786743] = 1
//line /snap/go/10455/src/net/port_unix.go:45
			_go_fuzz_dep_.CoverTab[7950]++
								if i != 1 {
//line /snap/go/10455/src/net/port_unix.go:46
				_go_fuzz_dep_.CoverTab[529572]++
//line /snap/go/10455/src/net/port_unix.go:46
				_go_fuzz_dep_.CoverTab[7951]++
									m[f[i]] = port
//line /snap/go/10455/src/net/port_unix.go:47
				// _ = "end of CoverTab[7951]"
			} else {
//line /snap/go/10455/src/net/port_unix.go:48
				_go_fuzz_dep_.CoverTab[529573]++
//line /snap/go/10455/src/net/port_unix.go:48
				_go_fuzz_dep_.CoverTab[7952]++
//line /snap/go/10455/src/net/port_unix.go:48
				// _ = "end of CoverTab[7952]"
//line /snap/go/10455/src/net/port_unix.go:48
			}
//line /snap/go/10455/src/net/port_unix.go:48
			// _ = "end of CoverTab[7950]"
		}
//line /snap/go/10455/src/net/port_unix.go:49
		if _go_fuzz_dep_.CoverTab[786743] == 0 {
//line /snap/go/10455/src/net/port_unix.go:49
			_go_fuzz_dep_.CoverTab[529580]++
//line /snap/go/10455/src/net/port_unix.go:49
		} else {
//line /snap/go/10455/src/net/port_unix.go:49
			_go_fuzz_dep_.CoverTab[529581]++
//line /snap/go/10455/src/net/port_unix.go:49
		}
//line /snap/go/10455/src/net/port_unix.go:49
		// _ = "end of CoverTab[7938]"
	}
//line /snap/go/10455/src/net/port_unix.go:50
	if _go_fuzz_dep_.CoverTab[786742] == 0 {
//line /snap/go/10455/src/net/port_unix.go:50
		_go_fuzz_dep_.CoverTab[529576]++
//line /snap/go/10455/src/net/port_unix.go:50
	} else {
//line /snap/go/10455/src/net/port_unix.go:50
		_go_fuzz_dep_.CoverTab[529577]++
//line /snap/go/10455/src/net/port_unix.go:50
	}
//line /snap/go/10455/src/net/port_unix.go:50
	// _ = "end of CoverTab[7931]"
}

// goLookupPort is the native Go implementation of LookupPort.
func goLookupPort(network, service string) (port int, err error) {
//line /snap/go/10455/src/net/port_unix.go:54
	_go_fuzz_dep_.CoverTab[7953]++
						onceReadServices.Do(readServices)
						return lookupPortMap(network, service)
//line /snap/go/10455/src/net/port_unix.go:56
	// _ = "end of CoverTab[7953]"
}

//line /snap/go/10455/src/net/port_unix.go:57
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/port_unix.go:57
var _ = _go_fuzz_dep_.CoverTab
