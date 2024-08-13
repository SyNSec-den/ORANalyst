// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Stuff that exists in std, but we can't use due to being a dependency
// of net, for go/build deps_test policy reasons.

//line /usr/local/go/src/net/netip/leaf_alts.go:8
package netip

//line /usr/local/go/src/net/netip/leaf_alts.go:8
import (
//line /usr/local/go/src/net/netip/leaf_alts.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/netip/leaf_alts.go:8
)
//line /usr/local/go/src/net/netip/leaf_alts.go:8
import (
//line /usr/local/go/src/net/netip/leaf_alts.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/netip/leaf_alts.go:8
)

func stringsLastIndexByte(s string, b byte) int {
//line /usr/local/go/src/net/netip/leaf_alts.go:10
	_go_fuzz_dep_.CoverTab[3546]++
							for i := len(s) - 1; i >= 0; i-- {
//line /usr/local/go/src/net/netip/leaf_alts.go:11
		_go_fuzz_dep_.CoverTab[3548]++
								if s[i] == b {
//line /usr/local/go/src/net/netip/leaf_alts.go:12
			_go_fuzz_dep_.CoverTab[3549]++
									return i
//line /usr/local/go/src/net/netip/leaf_alts.go:13
			// _ = "end of CoverTab[3549]"
		} else {
//line /usr/local/go/src/net/netip/leaf_alts.go:14
			_go_fuzz_dep_.CoverTab[3550]++
//line /usr/local/go/src/net/netip/leaf_alts.go:14
			// _ = "end of CoverTab[3550]"
//line /usr/local/go/src/net/netip/leaf_alts.go:14
		}
//line /usr/local/go/src/net/netip/leaf_alts.go:14
		// _ = "end of CoverTab[3548]"
	}
//line /usr/local/go/src/net/netip/leaf_alts.go:15
	// _ = "end of CoverTab[3546]"
//line /usr/local/go/src/net/netip/leaf_alts.go:15
	_go_fuzz_dep_.CoverTab[3547]++
							return -1
//line /usr/local/go/src/net/netip/leaf_alts.go:16
	// _ = "end of CoverTab[3547]"
}

func beUint64(b []byte) uint64 {
//line /usr/local/go/src/net/netip/leaf_alts.go:19
	_go_fuzz_dep_.CoverTab[3551]++
							_ = b[7]
							return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
//line /usr/local/go/src/net/netip/leaf_alts.go:22
	// _ = "end of CoverTab[3551]"
}

func bePutUint64(b []byte, v uint64) {
//line /usr/local/go/src/net/netip/leaf_alts.go:25
	_go_fuzz_dep_.CoverTab[3552]++
							_ = b[7]
							b[0] = byte(v >> 56)
							b[1] = byte(v >> 48)
							b[2] = byte(v >> 40)
							b[3] = byte(v >> 32)
							b[4] = byte(v >> 24)
							b[5] = byte(v >> 16)
							b[6] = byte(v >> 8)
							b[7] = byte(v)
//line /usr/local/go/src/net/netip/leaf_alts.go:34
	// _ = "end of CoverTab[3552]"
}

func bePutUint32(b []byte, v uint32) {
//line /usr/local/go/src/net/netip/leaf_alts.go:37
	_go_fuzz_dep_.CoverTab[3553]++
							_ = b[3]
							b[0] = byte(v >> 24)
							b[1] = byte(v >> 16)
							b[2] = byte(v >> 8)
							b[3] = byte(v)
//line /usr/local/go/src/net/netip/leaf_alts.go:42
	// _ = "end of CoverTab[3553]"
}

func leUint16(b []byte) uint16 {
//line /usr/local/go/src/net/netip/leaf_alts.go:45
	_go_fuzz_dep_.CoverTab[3554]++
							_ = b[1]
							return uint16(b[0]) | uint16(b[1])<<8
//line /usr/local/go/src/net/netip/leaf_alts.go:47
	// _ = "end of CoverTab[3554]"
}

func lePutUint16(b []byte, v uint16) {
//line /usr/local/go/src/net/netip/leaf_alts.go:50
	_go_fuzz_dep_.CoverTab[3555]++
							_ = b[1]
							b[0] = byte(v)
							b[1] = byte(v >> 8)
//line /usr/local/go/src/net/netip/leaf_alts.go:53
	// _ = "end of CoverTab[3555]"
}

//line /usr/local/go/src/net/netip/leaf_alts.go:54
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/netip/leaf_alts.go:54
var _ = _go_fuzz_dep_.CoverTab
