// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:5
)

import (
	"encoding/binary"
)

// NewUUID returns a Version 1 UUID based on the current NodeID and clock
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:11
// sequence, and the current time.  If the NodeID has not been set by SetNodeID
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:11
// or SetNodeInterface then it will be set automatically.  If the NodeID cannot
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:11
// be set NewUUID returns nil.  If clock sequence has not been set by
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:11
// SetClockSequence then it will be set automatically.  If GetTime fails to
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:11
// return the current NewUUID returns nil and an error.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:11
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:11
// In most cases, New should be used.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:19
func NewUUID() (UUID, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:19
	_go_fuzz_dep_.CoverTab[179529]++
										var uuid UUID
										now, seq, err := GetTime()
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:22
		_go_fuzz_dep_.CoverTab[179532]++
											return uuid, err
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:23
		// _ = "end of CoverTab[179532]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:24
		_go_fuzz_dep_.CoverTab[179533]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:24
		// _ = "end of CoverTab[179533]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:24
	// _ = "end of CoverTab[179529]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:24
	_go_fuzz_dep_.CoverTab[179530]++

										timeLow := uint32(now & 0xffffffff)
										timeMid := uint16((now >> 32) & 0xffff)
										timeHi := uint16((now >> 48) & 0x0fff)
										timeHi |= 0x1000

										binary.BigEndian.PutUint32(uuid[0:], timeLow)
										binary.BigEndian.PutUint16(uuid[4:], timeMid)
										binary.BigEndian.PutUint16(uuid[6:], timeHi)
										binary.BigEndian.PutUint16(uuid[8:], seq)

										nodeMu.Lock()
										if nodeID == zeroID {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:37
		_go_fuzz_dep_.CoverTab[179534]++
											setNodeInterface("")
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:38
		// _ = "end of CoverTab[179534]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:39
		_go_fuzz_dep_.CoverTab[179535]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:39
		// _ = "end of CoverTab[179535]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:39
	// _ = "end of CoverTab[179530]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:39
	_go_fuzz_dep_.CoverTab[179531]++
										copy(uuid[10:], nodeID[:])
										nodeMu.Unlock()

										return uuid, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:43
	// _ = "end of CoverTab[179531]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:44
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version1.go:44
var _ = _go_fuzz_dep_.CoverTab
