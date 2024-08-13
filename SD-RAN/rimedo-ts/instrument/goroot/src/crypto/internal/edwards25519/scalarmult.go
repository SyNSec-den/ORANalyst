// Copyright (c) 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:5
package edwards25519

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:5
)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:5
)

import "sync"

// basepointTable is a set of 32 affineLookupTables, where table i is generated
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:9
// from 256i * basepoint. It is precomputed the first time it's used.
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:11
func basepointTable() *[32]affineLookupTable {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:11
	_go_fuzz_dep_.CoverTab[9271]++
									basepointTablePrecomp.initOnce.Do(func() {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:12
		_go_fuzz_dep_.CoverTab[9273]++
										p := NewGeneratorPoint()
										for i := 0; i < 32; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:14
			_go_fuzz_dep_.CoverTab[9274]++
											basepointTablePrecomp.table[i].FromP3(p)
											for j := 0; j < 8; j++ {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:16
				_go_fuzz_dep_.CoverTab[9275]++
												p.Add(p, p)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:17
				// _ = "end of CoverTab[9275]"
			}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:18
			// _ = "end of CoverTab[9274]"
		}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:19
		// _ = "end of CoverTab[9273]"
	})
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:20
	// _ = "end of CoverTab[9271]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:20
	_go_fuzz_dep_.CoverTab[9272]++
									return &basepointTablePrecomp.table
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:21
	// _ = "end of CoverTab[9272]"
}

var basepointTablePrecomp struct {
	table		[32]affineLookupTable
	initOnce	sync.Once
}

// ScalarBaseMult sets v = x * B, where B is the canonical generator, and
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:29
// returns v.
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:29
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:29
// The scalar multiplication is done in constant time.
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:33
func (v *Point) ScalarBaseMult(x *Scalar) *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:33
	_go_fuzz_dep_.CoverTab[9276]++
									basepointTable := basepointTable()

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:47
	digits := x.signedRadix16()

									multiple := &affineCached{}
									tmp1 := &projP1xP1{}
									tmp2 := &projP2{}

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:54
	v.Set(NewIdentityPoint())
	for i := 1; i < 64; i += 2 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:55
		_go_fuzz_dep_.CoverTab[9279]++
										basepointTable[i/2].SelectInto(multiple, digits[i])
										tmp1.AddAffine(v, multiple)
										v.fromP1xP1(tmp1)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:58
		// _ = "end of CoverTab[9279]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:59
	// _ = "end of CoverTab[9276]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:59
	_go_fuzz_dep_.CoverTab[9277]++

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:62
	tmp2.FromP3(v)
									tmp1.Double(tmp2)
									tmp2.FromP1xP1(tmp1)
									tmp1.Double(tmp2)
									tmp2.FromP1xP1(tmp1)
									tmp1.Double(tmp2)
									tmp2.FromP1xP1(tmp1)
									tmp1.Double(tmp2)
									v.fromP1xP1(tmp1)

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:73
	for i := 0; i < 64; i += 2 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:73
		_go_fuzz_dep_.CoverTab[9280]++
										basepointTable[i/2].SelectInto(multiple, digits[i])
										tmp1.AddAffine(v, multiple)
										v.fromP1xP1(tmp1)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:76
		// _ = "end of CoverTab[9280]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:77
	// _ = "end of CoverTab[9277]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:77
	_go_fuzz_dep_.CoverTab[9278]++

									return v
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:79
	// _ = "end of CoverTab[9278]"
}

// ScalarMult sets v = x * q, and returns v.
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:82
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:82
// The scalar multiplication is done in constant time.
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:85
func (v *Point) ScalarMult(x *Scalar, q *Point) *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:85
	_go_fuzz_dep_.CoverTab[9281]++
									checkInitialized(q)

									var table projLookupTable
									table.FromP3(q)

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:98
	digits := x.signedRadix16()

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:101
	multiple := &projCached{}
	tmp1 := &projP1xP1{}
	tmp2 := &projP2{}
	table.SelectInto(multiple, digits[63])

	v.Set(NewIdentityPoint())
	tmp1.Add(v, multiple)
	for i := 62; i >= 0; i-- {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:108
		_go_fuzz_dep_.CoverTab[9283]++
											tmp2.FromP1xP1(tmp1)
											tmp1.Double(tmp2)
											tmp2.FromP1xP1(tmp1)
											tmp1.Double(tmp2)
											tmp2.FromP1xP1(tmp1)
											tmp1.Double(tmp2)
											tmp2.FromP1xP1(tmp1)
											tmp1.Double(tmp2)
											v.fromP1xP1(tmp1)
											table.SelectInto(multiple, digits[i])
											tmp1.Add(v, multiple)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:119
		// _ = "end of CoverTab[9283]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:120
	// _ = "end of CoverTab[9281]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:120
	_go_fuzz_dep_.CoverTab[9282]++
										v.fromP1xP1(tmp1)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:122
	// _ = "end of CoverTab[9282]"
}

// basepointNafTable is the nafLookupTable8 for the basepoint.
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:125
// It is precomputed the first time it's used.
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:127
func basepointNafTable() *nafLookupTable8 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:127
	_go_fuzz_dep_.CoverTab[9284]++
										basepointNafTablePrecomp.initOnce.Do(func() {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:128
		_go_fuzz_dep_.CoverTab[9286]++
											basepointNafTablePrecomp.table.FromP3(NewGeneratorPoint())
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:129
		// _ = "end of CoverTab[9286]"
	})
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:130
	// _ = "end of CoverTab[9284]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:130
	_go_fuzz_dep_.CoverTab[9285]++
										return &basepointNafTablePrecomp.table
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:131
	// _ = "end of CoverTab[9285]"
}

var basepointNafTablePrecomp struct {
	table		nafLookupTable8
	initOnce	sync.Once
}

// VarTimeDoubleScalarBaseMult sets v = a * A + b * B, where B is the canonical
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:139
// generator, and returns v.
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:139
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:139
// Execution time depends on the inputs.
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:143
func (v *Point) VarTimeDoubleScalarBaseMult(a *Scalar, A *Point, b *Scalar) *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:143
	_go_fuzz_dep_.CoverTab[9287]++
										checkInitialized(A)

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:160
	basepointNafTable := basepointNafTable()
										var aTable nafLookupTable5
										aTable.FromP3(A)

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:165
	aNaf := a.nonAdjacentForm(5)
										bNaf := b.nonAdjacentForm(8)

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:169
	i := 255
	for j := i; j >= 0; j-- {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:170
		_go_fuzz_dep_.CoverTab[9290]++
											if aNaf[j] != 0 || func() bool {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:171
			_go_fuzz_dep_.CoverTab[9291]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:171
			return bNaf[j] != 0
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:171
			// _ = "end of CoverTab[9291]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:171
		}() {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:171
			_go_fuzz_dep_.CoverTab[9292]++
												break
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:172
			// _ = "end of CoverTab[9292]"
		} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:173
			_go_fuzz_dep_.CoverTab[9293]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:173
			// _ = "end of CoverTab[9293]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:173
		}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:173
		// _ = "end of CoverTab[9290]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:174
	// _ = "end of CoverTab[9287]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:174
	_go_fuzz_dep_.CoverTab[9288]++

										multA := &projCached{}
										multB := &affineCached{}
										tmp1 := &projP1xP1{}
										tmp2 := &projP2{}
										tmp2.Zero()

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:185
	for ; i >= 0; i-- {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:185
		_go_fuzz_dep_.CoverTab[9294]++
											tmp1.Double(tmp2)

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:189
		if aNaf[i] > 0 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:189
			_go_fuzz_dep_.CoverTab[9297]++
												v.fromP1xP1(tmp1)
												aTable.SelectInto(multA, aNaf[i])
												tmp1.Add(v, multA)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:192
			// _ = "end of CoverTab[9297]"
		} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:193
			_go_fuzz_dep_.CoverTab[9298]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:193
			if aNaf[i] < 0 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:193
				_go_fuzz_dep_.CoverTab[9299]++
													v.fromP1xP1(tmp1)
													aTable.SelectInto(multA, -aNaf[i])
													tmp1.Sub(v, multA)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:196
				// _ = "end of CoverTab[9299]"
			} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:197
				_go_fuzz_dep_.CoverTab[9300]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:197
				// _ = "end of CoverTab[9300]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:197
			}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:197
			// _ = "end of CoverTab[9298]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:197
		}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:197
		// _ = "end of CoverTab[9294]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:197
		_go_fuzz_dep_.CoverTab[9295]++

											if bNaf[i] > 0 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:199
			_go_fuzz_dep_.CoverTab[9301]++
												v.fromP1xP1(tmp1)
												basepointNafTable.SelectInto(multB, bNaf[i])
												tmp1.AddAffine(v, multB)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:202
			// _ = "end of CoverTab[9301]"
		} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:203
			_go_fuzz_dep_.CoverTab[9302]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:203
			if bNaf[i] < 0 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:203
				_go_fuzz_dep_.CoverTab[9303]++
													v.fromP1xP1(tmp1)
													basepointNafTable.SelectInto(multB, -bNaf[i])
													tmp1.SubAffine(v, multB)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:206
				// _ = "end of CoverTab[9303]"
			} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:207
				_go_fuzz_dep_.CoverTab[9304]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:207
				// _ = "end of CoverTab[9304]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:207
			}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:207
			// _ = "end of CoverTab[9302]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:207
		}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:207
		// _ = "end of CoverTab[9295]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:207
		_go_fuzz_dep_.CoverTab[9296]++

											tmp2.FromP1xP1(tmp1)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:209
		// _ = "end of CoverTab[9296]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:210
	// _ = "end of CoverTab[9288]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:210
	_go_fuzz_dep_.CoverTab[9289]++

										v.fromP2(tmp2)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:213
	// _ = "end of CoverTab[9289]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:214
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/edwards25519/scalarmult.go:214
var _ = _go_fuzz_dep_.CoverTab
