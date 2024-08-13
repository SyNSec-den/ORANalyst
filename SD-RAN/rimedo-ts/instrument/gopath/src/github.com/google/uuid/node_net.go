// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:7
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:7
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:7
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:7
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:7
)

import "net"

var interfaces []net.Interface	// cached list of interfaces

// getHardwareInterface returns the name and hardware address of interface name.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:13
// If name is "" then the name and hardware address of one of the system's
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:13
// interfaces is returned.  If no interfaces are found (name does not exist or
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:13
// there are no interfaces) then "", nil is returned.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:13
// Only addresses of at least 6 bytes are returned.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:19
func getHardwareInterface(name string) (string, []byte) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:19
	_go_fuzz_dep_.CoverTab[179347]++
										if interfaces == nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:20
		_go_fuzz_dep_.CoverTab[179350]++
											var err error
											interfaces, err = net.Interfaces()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:23
			_go_fuzz_dep_.CoverTab[179351]++
												return "", nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:24
			// _ = "end of CoverTab[179351]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:25
			_go_fuzz_dep_.CoverTab[179352]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:25
			// _ = "end of CoverTab[179352]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:25
		// _ = "end of CoverTab[179350]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:26
		_go_fuzz_dep_.CoverTab[179353]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:26
		// _ = "end of CoverTab[179353]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:26
	// _ = "end of CoverTab[179347]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:26
	_go_fuzz_dep_.CoverTab[179348]++
										for _, ifs := range interfaces {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:27
		_go_fuzz_dep_.CoverTab[179354]++
											if len(ifs.HardwareAddr) >= 6 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:28
			_go_fuzz_dep_.CoverTab[179355]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:28
			return (name == "" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:28
				_go_fuzz_dep_.CoverTab[179356]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:28
				return name == ifs.Name
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:28
				// _ = "end of CoverTab[179356]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:28
			}())
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:28
			// _ = "end of CoverTab[179355]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:28
		}() {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:28
			_go_fuzz_dep_.CoverTab[179357]++
												return ifs.Name, ifs.HardwareAddr
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:29
			// _ = "end of CoverTab[179357]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:30
			_go_fuzz_dep_.CoverTab[179358]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:30
			// _ = "end of CoverTab[179358]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:30
		// _ = "end of CoverTab[179354]"
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:31
	// _ = "end of CoverTab[179348]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:31
	_go_fuzz_dep_.CoverTab[179349]++
										return "", nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:32
	// _ = "end of CoverTab[179349]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:33
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node_net.go:33
var _ = _go_fuzz_dep_.CoverTab
