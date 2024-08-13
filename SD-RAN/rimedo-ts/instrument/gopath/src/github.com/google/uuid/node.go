// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:5
)

import (
	"sync"
)

var (
	nodeMu	sync.Mutex
	ifname	string	// name of interface being used
	nodeID	[6]byte	// hardware for version 1 UUIDs
	zeroID	[6]byte	// nodeID with only 0's
)

// NodeInterface returns the name of the interface from which the NodeID was
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:18
// derived.  The interface "user" is returned if the NodeID was set by
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:18
// SetNodeID.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:21
func NodeInterface() string {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:21
	_go_fuzz_dep_.CoverTab[179328]++
										defer nodeMu.Unlock()
										nodeMu.Lock()
										return ifname
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:24
	// _ = "end of CoverTab[179328]"
}

// SetNodeInterface selects the hardware address to be used for Version 1 UUIDs.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:27
// If name is "" then the first usable interface found will be used or a random
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:27
// Node ID will be generated.  If a named interface cannot be found then false
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:27
// is returned.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:27
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:27
// SetNodeInterface never fails when name is "".
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:33
func SetNodeInterface(name string) bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:33
	_go_fuzz_dep_.CoverTab[179329]++
										defer nodeMu.Unlock()
										nodeMu.Lock()
										return setNodeInterface(name)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:36
	// _ = "end of CoverTab[179329]"
}

func setNodeInterface(name string) bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:39
	_go_fuzz_dep_.CoverTab[179330]++
										iname, addr := getHardwareInterface(name)
										if iname != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:41
		_go_fuzz_dep_.CoverTab[179333]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:41
		return addr != nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:41
		// _ = "end of CoverTab[179333]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:41
	}() {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:41
		_go_fuzz_dep_.CoverTab[179334]++
											ifname = iname
											copy(nodeID[:], addr)
											return true
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:44
		// _ = "end of CoverTab[179334]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:45
		_go_fuzz_dep_.CoverTab[179335]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:45
		// _ = "end of CoverTab[179335]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:45
	// _ = "end of CoverTab[179330]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:45
	_go_fuzz_dep_.CoverTab[179331]++

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:50
	if name == "" {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:50
		_go_fuzz_dep_.CoverTab[179336]++
											ifname = "random"
											randomBits(nodeID[:])
											return true
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:53
		// _ = "end of CoverTab[179336]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:54
		_go_fuzz_dep_.CoverTab[179337]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:54
		// _ = "end of CoverTab[179337]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:54
	// _ = "end of CoverTab[179331]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:54
	_go_fuzz_dep_.CoverTab[179332]++
										return false
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:55
	// _ = "end of CoverTab[179332]"
}

// NodeID returns a slice of a copy of the current Node ID, setting the Node ID
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:58
// if not already set.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:60
func NodeID() []byte {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:60
	_go_fuzz_dep_.CoverTab[179338]++
										defer nodeMu.Unlock()
										nodeMu.Lock()
										if nodeID == zeroID {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:63
		_go_fuzz_dep_.CoverTab[179340]++
											setNodeInterface("")
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:64
		// _ = "end of CoverTab[179340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:65
		_go_fuzz_dep_.CoverTab[179341]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:65
		// _ = "end of CoverTab[179341]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:65
	// _ = "end of CoverTab[179338]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:65
	_go_fuzz_dep_.CoverTab[179339]++
										nid := nodeID
										return nid[:]
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:67
	// _ = "end of CoverTab[179339]"
}

// SetNodeID sets the Node ID to be used for Version 1 UUIDs.  The first 6 bytes
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:70
// of id are used.  If id is less than 6 bytes then false is returned and the
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:70
// Node ID is not set.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:73
func SetNodeID(id []byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:73
	_go_fuzz_dep_.CoverTab[179342]++
										if len(id) < 6 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:74
		_go_fuzz_dep_.CoverTab[179344]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:75
		// _ = "end of CoverTab[179344]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:76
		_go_fuzz_dep_.CoverTab[179345]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:76
		// _ = "end of CoverTab[179345]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:76
	// _ = "end of CoverTab[179342]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:76
	_go_fuzz_dep_.CoverTab[179343]++
										defer nodeMu.Unlock()
										nodeMu.Lock()
										copy(nodeID[:], id)
										ifname = "user"
										return true
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:81
	// _ = "end of CoverTab[179343]"
}

// NodeID returns the 6 byte node id encoded in uuid.  It returns nil if uuid is
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:84
// not valid.  The NodeID is only well defined for version 1 and 2 UUIDs.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:86
func (uuid UUID) NodeID() []byte {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:86
	_go_fuzz_dep_.CoverTab[179346]++
										var node [6]byte
										copy(node[:], uuid[10:])
										return node[:]
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:89
	// _ = "end of CoverTab[179346]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:90
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/node.go:90
var _ = _go_fuzz_dep_.CoverTab
