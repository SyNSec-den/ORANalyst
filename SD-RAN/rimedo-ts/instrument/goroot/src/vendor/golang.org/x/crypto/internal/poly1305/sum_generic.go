// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides the generic implementation of Sum and MAC. Other files
// might provide optimized assembly implementations of some of this code.

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:8
package poly1305

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:8
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:8
)

import "encoding/binary"

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:28
func sumGeneric(out *[TagSize]byte, msg []byte, key *[32]byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:28
	_go_fuzz_dep_.CoverTab[20766]++
												h := newMACGeneric(key)
												h.Write(msg)
												h.Sum(out)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:31
	// _ = "end of CoverTab[20766]"
}

func newMACGeneric(key *[32]byte) macGeneric {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:34
	_go_fuzz_dep_.CoverTab[20767]++
												m := macGeneric{}
												initialize(key, &m.macState)
												return m
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:37
	// _ = "end of CoverTab[20767]"
}

// macState holds numbers in saturated 64-bit little-endian limbs. That is,
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:40
// the value of [x0, x1, x2] is x[0] + x[1] * 2⁶⁴ + x[2] * 2¹²⁸.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:42
type macState struct {
	// h is the main accumulator. It is to be interpreted modulo 2¹³⁰ - 5, but
	// can grow larger during and after rounds. It must, however, remain below
	// 2 * (2¹³⁰ - 5).
	h	[3]uint64
	// r and s are the private key components.
	r	[2]uint64
	s	[2]uint64
}

type macGeneric struct {
	macState

	buffer	[TagSize]byte
	offset	int
}

// Write splits the incoming message into TagSize chunks, and passes them to
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:59
// update. It buffers incomplete chunks.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:61
func (h *macGeneric) Write(p []byte) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:61
	_go_fuzz_dep_.CoverTab[20768]++
												nn := len(p)
												if h.offset > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:63
		_go_fuzz_dep_.CoverTab[20772]++
													n := copy(h.buffer[h.offset:], p)
													if h.offset+n < TagSize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:65
			_go_fuzz_dep_.CoverTab[20774]++
														h.offset += n
														return nn, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:67
			// _ = "end of CoverTab[20774]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:68
			_go_fuzz_dep_.CoverTab[20775]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:68
			// _ = "end of CoverTab[20775]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:68
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:68
		// _ = "end of CoverTab[20772]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:68
		_go_fuzz_dep_.CoverTab[20773]++
													p = p[n:]
													h.offset = 0
													updateGeneric(&h.macState, h.buffer[:])
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:71
		// _ = "end of CoverTab[20773]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:72
		_go_fuzz_dep_.CoverTab[20776]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:72
		// _ = "end of CoverTab[20776]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:72
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:72
	// _ = "end of CoverTab[20768]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:72
	_go_fuzz_dep_.CoverTab[20769]++
												if n := len(p) - (len(p) % TagSize); n > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:73
		_go_fuzz_dep_.CoverTab[20777]++
													updateGeneric(&h.macState, p[:n])
													p = p[n:]
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:75
		// _ = "end of CoverTab[20777]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:76
		_go_fuzz_dep_.CoverTab[20778]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:76
		// _ = "end of CoverTab[20778]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:76
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:76
	// _ = "end of CoverTab[20769]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:76
	_go_fuzz_dep_.CoverTab[20770]++
												if len(p) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:77
		_go_fuzz_dep_.CoverTab[20779]++
													h.offset += copy(h.buffer[h.offset:], p)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:78
		// _ = "end of CoverTab[20779]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:79
		_go_fuzz_dep_.CoverTab[20780]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:79
		// _ = "end of CoverTab[20780]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:79
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:79
	// _ = "end of CoverTab[20770]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:79
	_go_fuzz_dep_.CoverTab[20771]++
												return nn, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:80
	// _ = "end of CoverTab[20771]"
}

// Sum flushes the last incomplete chunk from the buffer, if any, and generates
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:83
// the MAC output. It does not modify its state, in order to allow for multiple
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:83
// calls to Sum, even if no Write is allowed after Sum.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:86
func (h *macGeneric) Sum(out *[TagSize]byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:86
	_go_fuzz_dep_.CoverTab[20781]++
												state := h.macState
												if h.offset > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:88
		_go_fuzz_dep_.CoverTab[20783]++
													updateGeneric(&state, h.buffer[:h.offset])
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:89
		// _ = "end of CoverTab[20783]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:90
		_go_fuzz_dep_.CoverTab[20784]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:90
		// _ = "end of CoverTab[20784]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:90
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:90
	// _ = "end of CoverTab[20781]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:90
	_go_fuzz_dep_.CoverTab[20782]++
												finalize(out, &state.h, &state.s)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:91
	// _ = "end of CoverTab[20782]"
}

// [rMask0, rMask1] is the specified Poly1305 clamping mask in little-endian. It
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:94
// clears some bits of the secret coefficient to make it possible to implement
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:94
// multiplication more efficiently.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:97
const (
	rMask0	= 0x0FFFFFFC0FFFFFFF
	rMask1	= 0x0FFFFFFC0FFFFFFC
)

// initialize loads the 256-bit key into the two 128-bit secret values r and s.
func initialize(key *[32]byte, m *macState) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:103
	_go_fuzz_dep_.CoverTab[20785]++
												m.r[0] = binary.LittleEndian.Uint64(key[0:8]) & rMask0
												m.r[1] = binary.LittleEndian.Uint64(key[8:16]) & rMask1
												m.s[0] = binary.LittleEndian.Uint64(key[16:24])
												m.s[1] = binary.LittleEndian.Uint64(key[24:32])
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:107
	// _ = "end of CoverTab[20785]"
}

// uint128 holds a 128-bit number as two 64-bit limbs, for use with the
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:110
// bits.Mul64 and bits.Add64 intrinsics.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:112
type uint128 struct {
	lo, hi uint64
}

func mul64(a, b uint64) uint128 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:116
	_go_fuzz_dep_.CoverTab[20786]++
												hi, lo := bitsMul64(a, b)
												return uint128{lo, hi}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:118
	// _ = "end of CoverTab[20786]"
}

func add128(a, b uint128) uint128 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:121
	_go_fuzz_dep_.CoverTab[20787]++
												lo, c := bitsAdd64(a.lo, b.lo, 0)
												hi, c := bitsAdd64(a.hi, b.hi, c)
												if c != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:124
		_go_fuzz_dep_.CoverTab[20789]++
													panic("poly1305: unexpected overflow")
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:125
		// _ = "end of CoverTab[20789]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:126
		_go_fuzz_dep_.CoverTab[20790]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:126
		// _ = "end of CoverTab[20790]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:126
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:126
	// _ = "end of CoverTab[20787]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:126
	_go_fuzz_dep_.CoverTab[20788]++
												return uint128{lo, hi}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:127
	// _ = "end of CoverTab[20788]"
}

func shiftRightBy2(a uint128) uint128 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:130
	_go_fuzz_dep_.CoverTab[20791]++
												a.lo = a.lo>>2 | (a.hi&3)<<62
												a.hi = a.hi >> 2
												return a
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:133
	// _ = "end of CoverTab[20791]"
}

// updateGeneric absorbs msg into the state.h accumulator. For each chunk m of
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:136
// 128 bits of message, it computes
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:136
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:136
//	h₊ = (h + m) * r  mod  2¹³⁰ - 5
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:136
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:136
// If the msg length is not a multiple of TagSize, it assumes the last
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:136
// incomplete chunk is the final one.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:143
func updateGeneric(state *macState, msg []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:143
	_go_fuzz_dep_.CoverTab[20792]++
												h0, h1, h2 := state.h[0], state.h[1], state.h[2]
												r0, r1 := state.r[0], state.r[1]

												for len(msg) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:147
		_go_fuzz_dep_.CoverTab[20794]++
													var c uint64

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:157
		if len(msg) >= TagSize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:157
			_go_fuzz_dep_.CoverTab[20798]++
														h0, c = bitsAdd64(h0, binary.LittleEndian.Uint64(msg[0:8]), 0)
														h1, c = bitsAdd64(h1, binary.LittleEndian.Uint64(msg[8:16]), c)
														h2 += c + 1

														msg = msg[TagSize:]
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:162
			// _ = "end of CoverTab[20798]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:163
			_go_fuzz_dep_.CoverTab[20799]++
														var buf [TagSize]byte
														copy(buf[:], msg)
														buf[len(msg)] = 1

														h0, c = bitsAdd64(h0, binary.LittleEndian.Uint64(buf[0:8]), 0)
														h1, c = bitsAdd64(h1, binary.LittleEndian.Uint64(buf[8:16]), c)
														h2 += c

														msg = nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:172
			// _ = "end of CoverTab[20799]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:173
		// _ = "end of CoverTab[20794]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:173
		_go_fuzz_dep_.CoverTab[20795]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:197
		h0r0 := mul64(h0, r0)
													h1r0 := mul64(h1, r0)
													h2r0 := mul64(h2, r0)
													h0r1 := mul64(h0, r1)
													h1r1 := mul64(h1, r1)
													h2r1 := mul64(h2, r1)

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:209
		if h2r0.hi != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:209
			_go_fuzz_dep_.CoverTab[20800]++
														panic("poly1305: unexpected overflow")
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:210
			// _ = "end of CoverTab[20800]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:211
			_go_fuzz_dep_.CoverTab[20801]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:211
			// _ = "end of CoverTab[20801]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:211
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:211
		// _ = "end of CoverTab[20795]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:211
		_go_fuzz_dep_.CoverTab[20796]++
													if h2r1.hi != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:212
			_go_fuzz_dep_.CoverTab[20802]++
														panic("poly1305: unexpected overflow")
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:213
			// _ = "end of CoverTab[20802]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:214
			_go_fuzz_dep_.CoverTab[20803]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:214
			// _ = "end of CoverTab[20803]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:214
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:214
		// _ = "end of CoverTab[20796]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:214
		_go_fuzz_dep_.CoverTab[20797]++

													m0 := h0r0
													m1 := add128(h1r0, h0r1)
													m2 := add128(h2r0, h1r1)
													m3 := h2r1

													t0 := m0.lo
													t1, c := bitsAdd64(m1.lo, m0.hi, 0)
													t2, c := bitsAdd64(m2.lo, m1.hi, c)
													t3, _ := bitsAdd64(m3.lo, m2.hi, c)

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:241
		h0, h1, h2 = t0, t1, t2&maskLow2Bits
													cc := uint128{t2 & maskNotLow2Bits, t3}

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:246
		h0, c = bitsAdd64(h0, cc.lo, 0)
													h1, c = bitsAdd64(h1, cc.hi, c)
													h2 += c

													cc = shiftRightBy2(cc)

													h0, c = bitsAdd64(h0, cc.lo, 0)
													h1, c = bitsAdd64(h1, cc.hi, c)
													h2 += c
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:254
		// _ = "end of CoverTab[20797]"

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:259
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:259
	// _ = "end of CoverTab[20792]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:259
	_go_fuzz_dep_.CoverTab[20793]++

												state.h[0], state.h[1], state.h[2] = h0, h1, h2
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:261
	// _ = "end of CoverTab[20793]"
}

const (
	maskLow2Bits	uint64	= 0x0000000000000003
	maskNotLow2Bits	uint64	= ^maskLow2Bits
)

// select64 returns x if v == 1 and y if v == 0, in constant time.
func select64(v, x, y uint64) uint64 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:270
	_go_fuzz_dep_.CoverTab[20804]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:270
	return ^(v-1)&x | (v-1)&y
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:270
	// _ = "end of CoverTab[20804]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:270
}

// [p0, p1, p2] is 2¹³⁰ - 5 in little endian order.
const (
	p0	= 0xFFFFFFFFFFFFFFFB
	p1	= 0xFFFFFFFFFFFFFFFF
	p2	= 0x0000000000000003
)

// finalize completes the modular reduction of h and computes
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:279
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:279
//	out = h + s  mod  2¹²⁸
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:282
func finalize(out *[TagSize]byte, h *[3]uint64, s *[2]uint64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:282
	_go_fuzz_dep_.CoverTab[20805]++
												h0, h1, h2 := h[0], h[1], h[2]

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:290
	hMinusP0, b := bitsSub64(h0, p0, 0)
												hMinusP1, b := bitsSub64(h1, p1, b)
												_, b = bitsSub64(h2, p2, b)

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:295
	h0 = select64(b, h0, hMinusP0)
												h1 = select64(b, h1, hMinusP1)

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:304
	h0, c := bitsAdd64(h0, s[0], 0)
												h1, _ = bitsAdd64(h1, s[1], c)

												binary.LittleEndian.PutUint64(out[0:8], h0)
												binary.LittleEndian.PutUint64(out[8:16], h1)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:308
	// _ = "end of CoverTab[20805]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:309
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_generic.go:309
var _ = _go_fuzz_dep_.CoverTab
