// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/netip/uint128.go:5
package netip

//line /snap/go/10455/src/net/netip/uint128.go:5
import (
//line /snap/go/10455/src/net/netip/uint128.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/netip/uint128.go:5
)
//line /snap/go/10455/src/net/netip/uint128.go:5
import (
//line /snap/go/10455/src/net/netip/uint128.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/netip/uint128.go:5
)

import "math/bits"

// uint128 represents a uint128 using two uint64s.
//line /snap/go/10455/src/net/netip/uint128.go:9
//
//line /snap/go/10455/src/net/netip/uint128.go:9
// When the methods below mention a bit number, bit 0 is the most
//line /snap/go/10455/src/net/netip/uint128.go:9
// significant bit (in hi) and bit 127 is the lowest (lo&1).
//line /snap/go/10455/src/net/netip/uint128.go:13
type uint128 struct {
	hi	uint64
	lo	uint64
}

// mask6 returns a uint128 bitmask with the topmost n bits of a
//line /snap/go/10455/src/net/netip/uint128.go:18
// 128-bit number.
//line /snap/go/10455/src/net/netip/uint128.go:20
func mask6(n int) uint128 {
//line /snap/go/10455/src/net/netip/uint128.go:20
	_go_fuzz_dep_.CoverTab[4473]++
							return uint128{^(^uint64(0) >> n), ^uint64(0) << (128 - n)}
//line /snap/go/10455/src/net/netip/uint128.go:21
	// _ = "end of CoverTab[4473]"
}

// isZero reports whether u == 0.
//line /snap/go/10455/src/net/netip/uint128.go:24
//
//line /snap/go/10455/src/net/netip/uint128.go:24
// It's faster than u == (uint128{}) because the compiler (as of Go
//line /snap/go/10455/src/net/netip/uint128.go:24
// 1.15/1.16b1) doesn't do this trick and instead inserts a branch in
//line /snap/go/10455/src/net/netip/uint128.go:24
// its eq alg's generated code.
//line /snap/go/10455/src/net/netip/uint128.go:29
func (u uint128) isZero() bool {
//line /snap/go/10455/src/net/netip/uint128.go:29
	_go_fuzz_dep_.CoverTab[4474]++
//line /snap/go/10455/src/net/netip/uint128.go:29
	return u.hi|u.lo == 0
//line /snap/go/10455/src/net/netip/uint128.go:29
	// _ = "end of CoverTab[4474]"
//line /snap/go/10455/src/net/netip/uint128.go:29
}

// and returns the bitwise AND of u and m (u&m).
func (u uint128) and(m uint128) uint128 {
//line /snap/go/10455/src/net/netip/uint128.go:32
	_go_fuzz_dep_.CoverTab[4475]++
							return uint128{u.hi & m.hi, u.lo & m.lo}
//line /snap/go/10455/src/net/netip/uint128.go:33
	// _ = "end of CoverTab[4475]"
}

// xor returns the bitwise XOR of u and m (u^m).
func (u uint128) xor(m uint128) uint128 {
//line /snap/go/10455/src/net/netip/uint128.go:37
	_go_fuzz_dep_.CoverTab[4476]++
							return uint128{u.hi ^ m.hi, u.lo ^ m.lo}
//line /snap/go/10455/src/net/netip/uint128.go:38
	// _ = "end of CoverTab[4476]"
}

// or returns the bitwise OR of u and m (u|m).
func (u uint128) or(m uint128) uint128 {
//line /snap/go/10455/src/net/netip/uint128.go:42
	_go_fuzz_dep_.CoverTab[4477]++
							return uint128{u.hi | m.hi, u.lo | m.lo}
//line /snap/go/10455/src/net/netip/uint128.go:43
	// _ = "end of CoverTab[4477]"
}

// not returns the bitwise NOT of u.
func (u uint128) not() uint128 {
//line /snap/go/10455/src/net/netip/uint128.go:47
	_go_fuzz_dep_.CoverTab[4478]++
							return uint128{^u.hi, ^u.lo}
//line /snap/go/10455/src/net/netip/uint128.go:48
	// _ = "end of CoverTab[4478]"
}

// subOne returns u - 1.
func (u uint128) subOne() uint128 {
//line /snap/go/10455/src/net/netip/uint128.go:52
	_go_fuzz_dep_.CoverTab[4479]++
							lo, borrow := bits.Sub64(u.lo, 1, 0)
							return uint128{u.hi - borrow, lo}
//line /snap/go/10455/src/net/netip/uint128.go:54
	// _ = "end of CoverTab[4479]"
}

// addOne returns u + 1.
func (u uint128) addOne() uint128 {
//line /snap/go/10455/src/net/netip/uint128.go:58
	_go_fuzz_dep_.CoverTab[4480]++
							lo, carry := bits.Add64(u.lo, 1, 0)
							return uint128{u.hi + carry, lo}
//line /snap/go/10455/src/net/netip/uint128.go:60
	// _ = "end of CoverTab[4480]"
}

// halves returns the two uint64 halves of the uint128.
//line /snap/go/10455/src/net/netip/uint128.go:63
//
//line /snap/go/10455/src/net/netip/uint128.go:63
// Logically, think of it as returning two uint64s.
//line /snap/go/10455/src/net/netip/uint128.go:63
// It only returns pointers for inlining reasons on 32-bit platforms.
//line /snap/go/10455/src/net/netip/uint128.go:67
func (u *uint128) halves() [2]*uint64 {
//line /snap/go/10455/src/net/netip/uint128.go:67
	_go_fuzz_dep_.CoverTab[4481]++
							return [2]*uint64{&u.hi, &u.lo}
//line /snap/go/10455/src/net/netip/uint128.go:68
	// _ = "end of CoverTab[4481]"
}

// bitsSetFrom returns a copy of u with the given bit
//line /snap/go/10455/src/net/netip/uint128.go:71
// and all subsequent ones set.
//line /snap/go/10455/src/net/netip/uint128.go:73
func (u uint128) bitsSetFrom(bit uint8) uint128 {
//line /snap/go/10455/src/net/netip/uint128.go:73
	_go_fuzz_dep_.CoverTab[4482]++
							return u.or(mask6(int(bit)).not())
//line /snap/go/10455/src/net/netip/uint128.go:74
	// _ = "end of CoverTab[4482]"
}

// bitsClearedFrom returns a copy of u with the given bit
//line /snap/go/10455/src/net/netip/uint128.go:77
// and all subsequent ones cleared.
//line /snap/go/10455/src/net/netip/uint128.go:79
func (u uint128) bitsClearedFrom(bit uint8) uint128 {
//line /snap/go/10455/src/net/netip/uint128.go:79
	_go_fuzz_dep_.CoverTab[4483]++
							return u.and(mask6(int(bit)))
//line /snap/go/10455/src/net/netip/uint128.go:80
	// _ = "end of CoverTab[4483]"
}

//line /snap/go/10455/src/net/netip/uint128.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/netip/uint128.go:81
var _ = _go_fuzz_dep_.CoverTab
