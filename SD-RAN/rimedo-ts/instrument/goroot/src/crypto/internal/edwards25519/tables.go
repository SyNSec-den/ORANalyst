// Copyright (c) 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:5
package edwards25519

//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:5
)
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:5
)

import (
	"crypto/subtle"
)

// A dynamic lookup table for variable-base, constant-time scalar muls.
type projLookupTable struct {
	points [8]projCached
}

// A precomputed lookup table for fixed-base, constant-time scalar muls.
type affineLookupTable struct {
	points [8]affineCached
}

// A dynamic lookup table for variable-base, variable-time scalar muls.
type nafLookupTable5 struct {
	points [8]projCached
}

// A precomputed lookup table for fixed-base, variable-time scalar muls.
type nafLookupTable8 struct {
	points [64]affineCached
}

//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:33
// Builds a lookup table at runtime. Fast.
func (v *projLookupTable) FromP3(q *Point) {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:34
	_go_fuzz_dep_.CoverTab[9305]++

//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:37
	v.points[0].FromP3(q)
	tmpP3 := Point{}
	tmpP1xP1 := projP1xP1{}
	for i := 0; i < 7; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:40
		_go_fuzz_dep_.CoverTab[9306]++

//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:44
		v.points[i+1].FromP3(tmpP3.fromP1xP1(tmpP1xP1.Add(q, &v.points[i])))
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:44
		// _ = "end of CoverTab[9306]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:45
	// _ = "end of CoverTab[9305]"
}

// This is not optimised for speed; fixed-base tables should be precomputed.
func (v *affineLookupTable) FromP3(q *Point) {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:49
	_go_fuzz_dep_.CoverTab[9307]++

//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:52
	v.points[0].FromP3(q)
	tmpP3 := Point{}
	tmpP1xP1 := projP1xP1{}
	for i := 0; i < 7; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:55
		_go_fuzz_dep_.CoverTab[9308]++

										v.points[i+1].FromP3(tmpP3.fromP1xP1(tmpP1xP1.AddAffine(q, &v.points[i])))
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:57
		// _ = "end of CoverTab[9308]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:58
	// _ = "end of CoverTab[9307]"
}

// Builds a lookup table at runtime. Fast.
func (v *nafLookupTable5) FromP3(q *Point) {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:62
	_go_fuzz_dep_.CoverTab[9309]++

//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:65
	v.points[0].FromP3(q)
	q2 := Point{}
	q2.Add(q, q)
	tmpP3 := Point{}
	tmpP1xP1 := projP1xP1{}
	for i := 0; i < 7; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:70
		_go_fuzz_dep_.CoverTab[9310]++
										v.points[i+1].FromP3(tmpP3.fromP1xP1(tmpP1xP1.Add(&q2, &v.points[i])))
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:71
		// _ = "end of CoverTab[9310]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:72
	// _ = "end of CoverTab[9309]"
}

// This is not optimised for speed; fixed-base tables should be precomputed.
func (v *nafLookupTable8) FromP3(q *Point) {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:76
	_go_fuzz_dep_.CoverTab[9311]++
									v.points[0].FromP3(q)
									q2 := Point{}
									q2.Add(q, q)
									tmpP3 := Point{}
									tmpP1xP1 := projP1xP1{}
									for i := 0; i < 63; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:82
		_go_fuzz_dep_.CoverTab[9312]++
										v.points[i+1].FromP3(tmpP3.fromP1xP1(tmpP1xP1.AddAffine(&q2, &v.points[i])))
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:83
		// _ = "end of CoverTab[9312]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:84
	// _ = "end of CoverTab[9311]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:89
// Set dest to x*Q, where -8 <= x <= 8, in constant time.
func (v *projLookupTable) SelectInto(dest *projCached, x int8) {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:90
	_go_fuzz_dep_.CoverTab[9313]++

									xmask := x >> 7
									xabs := uint8((x + xmask) ^ xmask)

									dest.Zero()
									for j := 1; j <= 8; j++ {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:96
		_go_fuzz_dep_.CoverTab[9315]++

										cond := subtle.ConstantTimeByteEq(xabs, uint8(j))
										dest.Select(&v.points[j-1], dest, cond)
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:99
		// _ = "end of CoverTab[9315]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:100
	// _ = "end of CoverTab[9313]"
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:100
	_go_fuzz_dep_.CoverTab[9314]++

									dest.CondNeg(int(xmask & 1))
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:102
	// _ = "end of CoverTab[9314]"
}

// Set dest to x*Q, where -8 <= x <= 8, in constant time.
func (v *affineLookupTable) SelectInto(dest *affineCached, x int8) {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:106
	_go_fuzz_dep_.CoverTab[9316]++

									xmask := x >> 7
									xabs := uint8((x + xmask) ^ xmask)

									dest.Zero()
									for j := 1; j <= 8; j++ {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:112
		_go_fuzz_dep_.CoverTab[9318]++

										cond := subtle.ConstantTimeByteEq(xabs, uint8(j))
										dest.Select(&v.points[j-1], dest, cond)
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:115
		// _ = "end of CoverTab[9318]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:116
	// _ = "end of CoverTab[9316]"
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:116
	_go_fuzz_dep_.CoverTab[9317]++

									dest.CondNeg(int(xmask & 1))
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:118
	// _ = "end of CoverTab[9317]"
}

// Given odd x with 0 < x < 2^4, return x*Q (in variable time).
func (v *nafLookupTable5) SelectInto(dest *projCached, x int8) {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:122
	_go_fuzz_dep_.CoverTab[9319]++
									*dest = v.points[x/2]
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:123
	// _ = "end of CoverTab[9319]"
}

// Given odd x with 0 < x < 2^7, return x*Q (in variable time).
func (v *nafLookupTable8) SelectInto(dest *affineCached, x int8) {
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:127
	_go_fuzz_dep_.CoverTab[9320]++
									*dest = v.points[x/2]
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:128
	// _ = "end of CoverTab[9320]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:129
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/edwards25519/tables.go:129
var _ = _go_fuzz_dep_.CoverTab
