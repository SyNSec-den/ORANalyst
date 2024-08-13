// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Stuff that exists in std, but we can't use due to being a dependency
// of net, for go/build deps_test policy reasons.

//line /snap/go/10455/src/net/netip/leaf_alts.go:8
package netip

//line /snap/go/10455/src/net/netip/leaf_alts.go:8
import (
//line /snap/go/10455/src/net/netip/leaf_alts.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/netip/leaf_alts.go:8
)
//line /snap/go/10455/src/net/netip/leaf_alts.go:8
import (
//line /snap/go/10455/src/net/netip/leaf_alts.go:8
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/netip/leaf_alts.go:8
)

func stringsLastIndexByte(s string, b byte) int {
//line /snap/go/10455/src/net/netip/leaf_alts.go:10
	_go_fuzz_dep_.CoverTab[3882]++
//line /snap/go/10455/src/net/netip/leaf_alts.go:10
	_go_fuzz_dep_.CoverTab[786624] = 0
							for i := len(s) - 1; i >= 0; i-- {
//line /snap/go/10455/src/net/netip/leaf_alts.go:11
		if _go_fuzz_dep_.CoverTab[786624] == 0 {
//line /snap/go/10455/src/net/netip/leaf_alts.go:11
			_go_fuzz_dep_.CoverTab[527014]++
//line /snap/go/10455/src/net/netip/leaf_alts.go:11
		} else {
//line /snap/go/10455/src/net/netip/leaf_alts.go:11
			_go_fuzz_dep_.CoverTab[527015]++
//line /snap/go/10455/src/net/netip/leaf_alts.go:11
		}
//line /snap/go/10455/src/net/netip/leaf_alts.go:11
		_go_fuzz_dep_.CoverTab[786624] = 1
//line /snap/go/10455/src/net/netip/leaf_alts.go:11
		_go_fuzz_dep_.CoverTab[3884]++
								if s[i] == b {
//line /snap/go/10455/src/net/netip/leaf_alts.go:12
			_go_fuzz_dep_.CoverTab[527012]++
//line /snap/go/10455/src/net/netip/leaf_alts.go:12
			_go_fuzz_dep_.CoverTab[3885]++
									return i
//line /snap/go/10455/src/net/netip/leaf_alts.go:13
			// _ = "end of CoverTab[3885]"
		} else {
//line /snap/go/10455/src/net/netip/leaf_alts.go:14
			_go_fuzz_dep_.CoverTab[527013]++
//line /snap/go/10455/src/net/netip/leaf_alts.go:14
			_go_fuzz_dep_.CoverTab[3886]++
//line /snap/go/10455/src/net/netip/leaf_alts.go:14
			// _ = "end of CoverTab[3886]"
//line /snap/go/10455/src/net/netip/leaf_alts.go:14
		}
//line /snap/go/10455/src/net/netip/leaf_alts.go:14
		// _ = "end of CoverTab[3884]"
	}
//line /snap/go/10455/src/net/netip/leaf_alts.go:15
	if _go_fuzz_dep_.CoverTab[786624] == 0 {
//line /snap/go/10455/src/net/netip/leaf_alts.go:15
		_go_fuzz_dep_.CoverTab[527016]++
//line /snap/go/10455/src/net/netip/leaf_alts.go:15
	} else {
//line /snap/go/10455/src/net/netip/leaf_alts.go:15
		_go_fuzz_dep_.CoverTab[527017]++
//line /snap/go/10455/src/net/netip/leaf_alts.go:15
	}
//line /snap/go/10455/src/net/netip/leaf_alts.go:15
	// _ = "end of CoverTab[3882]"
//line /snap/go/10455/src/net/netip/leaf_alts.go:15
	_go_fuzz_dep_.CoverTab[3883]++
							return -1
//line /snap/go/10455/src/net/netip/leaf_alts.go:16
	// _ = "end of CoverTab[3883]"
}

func beUint64(b []byte) uint64 {
//line /snap/go/10455/src/net/netip/leaf_alts.go:19
	_go_fuzz_dep_.CoverTab[3887]++
							_ = b[7]
							return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
//line /snap/go/10455/src/net/netip/leaf_alts.go:22
	// _ = "end of CoverTab[3887]"
}

func bePutUint64(b []byte, v uint64) {
//line /snap/go/10455/src/net/netip/leaf_alts.go:25
	_go_fuzz_dep_.CoverTab[3888]++
							_ = b[7]
							b[0] = byte(v >> 56)
							b[1] = byte(v >> 48)
							b[2] = byte(v >> 40)
							b[3] = byte(v >> 32)
							b[4] = byte(v >> 24)
							b[5] = byte(v >> 16)
							b[6] = byte(v >> 8)
							b[7] = byte(v)
//line /snap/go/10455/src/net/netip/leaf_alts.go:34
	// _ = "end of CoverTab[3888]"
}

func bePutUint32(b []byte, v uint32) {
//line /snap/go/10455/src/net/netip/leaf_alts.go:37
	_go_fuzz_dep_.CoverTab[3889]++
							_ = b[3]
							b[0] = byte(v >> 24)
							b[1] = byte(v >> 16)
							b[2] = byte(v >> 8)
							b[3] = byte(v)
//line /snap/go/10455/src/net/netip/leaf_alts.go:42
	// _ = "end of CoverTab[3889]"
}

func leUint16(b []byte) uint16 {
//line /snap/go/10455/src/net/netip/leaf_alts.go:45
	_go_fuzz_dep_.CoverTab[3890]++
							_ = b[1]
							return uint16(b[0]) | uint16(b[1])<<8
//line /snap/go/10455/src/net/netip/leaf_alts.go:47
	// _ = "end of CoverTab[3890]"
}

func lePutUint16(b []byte, v uint16) {
//line /snap/go/10455/src/net/netip/leaf_alts.go:50
	_go_fuzz_dep_.CoverTab[3891]++
							_ = b[1]
							b[0] = byte(v)
							b[1] = byte(v >> 8)
//line /snap/go/10455/src/net/netip/leaf_alts.go:53
	// _ = "end of CoverTab[3891]"
}

//line /snap/go/10455/src/net/netip/leaf_alts.go:54
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/netip/leaf_alts.go:54
var _ = _go_fuzz_dep_.CoverTab
