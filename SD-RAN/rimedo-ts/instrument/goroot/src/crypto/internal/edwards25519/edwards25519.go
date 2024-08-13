// Copyright (c) 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:5
package edwards25519

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:5
)
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:5
)

import (
	"crypto/internal/edwards25519/field"
	"errors"
)

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:14
type projP1xP1 struct {
	X, Y, Z, T field.Element
}

type projP2 struct {
	X, Y, Z field.Element
}

// Point represents a point on the edwards25519 curve.
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:22
//
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:22
// This type works similarly to math/big.Int, and all arguments and receivers
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:22
// are allowed to alias.
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:22
//
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:22
// The zero value is NOT valid, and it may be used only as a receiver.
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:28
type Point struct {
	// The point is internally represented in extended coordinates (X, Y, Z, T)
	// where x = X/Z, y = Y/Z, and xy = T/Z per https://eprint.iacr.org/2008/522.
	x, y, z, t	field.Element

	// Make the type not comparable (i.e. used with == or as a map key), as
	// equivalent points can be represented by different Go values.
	_	incomparable
}

type incomparable [0]func()

func checkInitialized(points ...*Point) {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:40
	_go_fuzz_dep_.CoverTab[9153]++
										for _, p := range points {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:41
		_go_fuzz_dep_.CoverTab[9154]++
											if p.x == (field.Element{}) && func() bool {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:42
			_go_fuzz_dep_.CoverTab[9155]++
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:42
			return p.y == (field.Element{})
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:42
			// _ = "end of CoverTab[9155]"
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:42
		}() {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:42
			_go_fuzz_dep_.CoverTab[9156]++
												panic("edwards25519: use of uninitialized Point")
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:43
			// _ = "end of CoverTab[9156]"
		} else {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:44
			_go_fuzz_dep_.CoverTab[9157]++
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:44
			// _ = "end of CoverTab[9157]"
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:44
		}
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:44
		// _ = "end of CoverTab[9154]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:45
	// _ = "end of CoverTab[9153]"
}

type projCached struct {
	YplusX, YminusX, Z, T2d field.Element
}

type affineCached struct {
	YplusX, YminusX, T2d field.Element
}

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:58
func (v *projP2) Zero() *projP2 {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:58
	_go_fuzz_dep_.CoverTab[9158]++
										v.X.Zero()
										v.Y.One()
										v.Z.One()
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:62
	// _ = "end of CoverTab[9158]"
}

// identity is the point at infinity.
var identity, _ = new(Point).SetBytes([]byte{
	1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

// NewIdentityPoint returns a new Point set to the identity.
func NewIdentityPoint() *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:71
	_go_fuzz_dep_.CoverTab[9159]++
										return new(Point).Set(identity)
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:72
	// _ = "end of CoverTab[9159]"
}

// generator is the canonical curve basepoint. See TestGenerator for the
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:75
// correspondence of this encoding with the values in RFC 8032.
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:77
var generator, _ = new(Point).SetBytes([]byte{
	0x58, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66})

// NewGeneratorPoint returns a new Point set to the canonical generator.
func NewGeneratorPoint() *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:84
	_go_fuzz_dep_.CoverTab[9160]++
										return new(Point).Set(generator)
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:85
	// _ = "end of CoverTab[9160]"
}

func (v *projCached) Zero() *projCached {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:88
	_go_fuzz_dep_.CoverTab[9161]++
										v.YplusX.One()
										v.YminusX.One()
										v.Z.One()
										v.T2d.Zero()
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:93
	// _ = "end of CoverTab[9161]"
}

func (v *affineCached) Zero() *affineCached {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:96
	_go_fuzz_dep_.CoverTab[9162]++
										v.YplusX.One()
										v.YminusX.One()
										v.T2d.Zero()
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:100
	// _ = "end of CoverTab[9162]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:105
// Set sets v = u, and returns v.
func (v *Point) Set(u *Point) *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:106
	_go_fuzz_dep_.CoverTab[9163]++
										*v = *u
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:108
	// _ = "end of CoverTab[9163]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:113
// Bytes returns the canonical 32-byte encoding of v, according to RFC 8032,
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:113
// Section 5.1.2.
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:115
func (v *Point) Bytes() []byte {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:115
	_go_fuzz_dep_.CoverTab[9164]++
	// This function is outlined to make the allocations inline in the caller
										// rather than happen on the heap.
										var buf [32]byte
										return v.bytes(&buf)
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:119
	// _ = "end of CoverTab[9164]"
}

func (v *Point) bytes(buf *[32]byte) []byte {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:122
	_go_fuzz_dep_.CoverTab[9165]++
										checkInitialized(v)

										var zInv, x, y field.Element
										zInv.Invert(&v.z)
										x.Multiply(&v.x, &zInv)
										y.Multiply(&v.y, &zInv)

										out := copyFieldElement(buf, &y)
										out[31] |= byte(x.IsNegative() << 7)
										return out
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:132
	// _ = "end of CoverTab[9165]"
}

var feOne = new(field.Element).One()

// SetBytes sets v = x, where x is a 32-byte encoding of v. If x does not
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:137
// represent a valid point on the curve, SetBytes returns nil and an error and
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:137
// the receiver is unchanged. Otherwise, SetBytes returns v.
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:137
//
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:137
// Note that SetBytes accepts all non-canonical encodings of valid points.
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:137
// That is, it follows decoding rules that match most implementations in
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:137
// the ecosystem rather than RFC 8032.
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:144
func (v *Point) SetBytes(x []byte) (*Point, error) {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:144
	_go_fuzz_dep_.CoverTab[9166]++

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:153
	y, err := new(field.Element).SetBytes(x)
	if err != nil {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:154
		_go_fuzz_dep_.CoverTab[9169]++
											return nil, errors.New("edwards25519: invalid point encoding length")
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:155
		// _ = "end of CoverTab[9169]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:156
		_go_fuzz_dep_.CoverTab[9170]++
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:156
		// _ = "end of CoverTab[9170]"
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:156
	}
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:156
	// _ = "end of CoverTab[9166]"
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:156
	_go_fuzz_dep_.CoverTab[9167]++

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:163
	y2 := new(field.Element).Square(y)
										u := new(field.Element).Subtract(y2, feOne)

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:167
	vv := new(field.Element).Multiply(y2, d)
										vv = vv.Add(vv, feOne)

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:171
	xx, wasSquare := new(field.Element).SqrtRatio(u, vv)
	if wasSquare == 0 {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:172
		_go_fuzz_dep_.CoverTab[9171]++
											return nil, errors.New("edwards25519: invalid point encoding")
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:173
		// _ = "end of CoverTab[9171]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:174
		_go_fuzz_dep_.CoverTab[9172]++
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:174
		// _ = "end of CoverTab[9172]"
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:174
	}
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:174
	// _ = "end of CoverTab[9167]"
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:174
	_go_fuzz_dep_.CoverTab[9168]++

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:177
	xxNeg := new(field.Element).Negate(xx)
										xx = xx.Select(xxNeg, xx, int(x[31]>>7))

										v.x.Set(xx)
										v.y.Set(y)
										v.z.One()
										v.t.Multiply(xx, y)

										return v, nil
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:185
	// _ = "end of CoverTab[9168]"
}

func copyFieldElement(buf *[32]byte, v *field.Element) []byte {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:188
	_go_fuzz_dep_.CoverTab[9173]++
										copy(buf[:], v.Bytes())
										return buf[:]
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:190
	// _ = "end of CoverTab[9173]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:195
func (v *projP2) FromP1xP1(p *projP1xP1) *projP2 {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:195
	_go_fuzz_dep_.CoverTab[9174]++
										v.X.Multiply(&p.X, &p.T)
										v.Y.Multiply(&p.Y, &p.Z)
										v.Z.Multiply(&p.Z, &p.T)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:199
	// _ = "end of CoverTab[9174]"
}

func (v *projP2) FromP3(p *Point) *projP2 {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:202
	_go_fuzz_dep_.CoverTab[9175]++
										v.X.Set(&p.x)
										v.Y.Set(&p.y)
										v.Z.Set(&p.z)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:206
	// _ = "end of CoverTab[9175]"
}

func (v *Point) fromP1xP1(p *projP1xP1) *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:209
	_go_fuzz_dep_.CoverTab[9176]++
										v.x.Multiply(&p.X, &p.T)
										v.y.Multiply(&p.Y, &p.Z)
										v.z.Multiply(&p.Z, &p.T)
										v.t.Multiply(&p.X, &p.Y)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:214
	// _ = "end of CoverTab[9176]"
}

func (v *Point) fromP2(p *projP2) *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:217
	_go_fuzz_dep_.CoverTab[9177]++
										v.x.Multiply(&p.X, &p.Z)
										v.y.Multiply(&p.Y, &p.Z)
										v.z.Square(&p.Z)
										v.t.Multiply(&p.X, &p.Y)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:222
	// _ = "end of CoverTab[9177]"
}

// d is a constant in the curve equation.
var d, _ = new(field.Element).SetBytes([]byte{
	0xa3, 0x78, 0x59, 0x13, 0xca, 0x4d, 0xeb, 0x75,
	0xab, 0xd8, 0x41, 0x41, 0x4d, 0x0a, 0x70, 0x00,
	0x98, 0xe8, 0x79, 0x77, 0x79, 0x40, 0xc7, 0x8c,
	0x73, 0xfe, 0x6f, 0x2b, 0xee, 0x6c, 0x03, 0x52})
var d2 = new(field.Element).Add(d, d)

func (v *projCached) FromP3(p *Point) *projCached {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:233
	_go_fuzz_dep_.CoverTab[9178]++
										v.YplusX.Add(&p.y, &p.x)
										v.YminusX.Subtract(&p.y, &p.x)
										v.Z.Set(&p.z)
										v.T2d.Multiply(&p.t, d2)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:238
	// _ = "end of CoverTab[9178]"
}

func (v *affineCached) FromP3(p *Point) *affineCached {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:241
	_go_fuzz_dep_.CoverTab[9179]++
										v.YplusX.Add(&p.y, &p.x)
										v.YminusX.Subtract(&p.y, &p.x)
										v.T2d.Multiply(&p.t, d2)

										var invZ field.Element
										invZ.Invert(&p.z)
										v.YplusX.Multiply(&v.YplusX, &invZ)
										v.YminusX.Multiply(&v.YminusX, &invZ)
										v.T2d.Multiply(&v.T2d, &invZ)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:251
	// _ = "end of CoverTab[9179]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:256
// Add sets v = p + q, and returns v.
func (v *Point) Add(p, q *Point) *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:257
	_go_fuzz_dep_.CoverTab[9180]++
										checkInitialized(p, q)
										qCached := new(projCached).FromP3(q)
										result := new(projP1xP1).Add(p, qCached)
										return v.fromP1xP1(result)
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:261
	// _ = "end of CoverTab[9180]"
}

// Subtract sets v = p - q, and returns v.
func (v *Point) Subtract(p, q *Point) *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:265
	_go_fuzz_dep_.CoverTab[9181]++
										checkInitialized(p, q)
										qCached := new(projCached).FromP3(q)
										result := new(projP1xP1).Sub(p, qCached)
										return v.fromP1xP1(result)
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:269
	// _ = "end of CoverTab[9181]"
}

func (v *projP1xP1) Add(p *Point, q *projCached) *projP1xP1 {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:272
	_go_fuzz_dep_.CoverTab[9182]++
										var YplusX, YminusX, PP, MM, TT2d, ZZ2 field.Element

										YplusX.Add(&p.y, &p.x)
										YminusX.Subtract(&p.y, &p.x)

										PP.Multiply(&YplusX, &q.YplusX)
										MM.Multiply(&YminusX, &q.YminusX)
										TT2d.Multiply(&p.t, &q.T2d)
										ZZ2.Multiply(&p.z, &q.Z)

										ZZ2.Add(&ZZ2, &ZZ2)

										v.X.Subtract(&PP, &MM)
										v.Y.Add(&PP, &MM)
										v.Z.Add(&ZZ2, &TT2d)
										v.T.Subtract(&ZZ2, &TT2d)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:289
	// _ = "end of CoverTab[9182]"
}

func (v *projP1xP1) Sub(p *Point, q *projCached) *projP1xP1 {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:292
	_go_fuzz_dep_.CoverTab[9183]++
										var YplusX, YminusX, PP, MM, TT2d, ZZ2 field.Element

										YplusX.Add(&p.y, &p.x)
										YminusX.Subtract(&p.y, &p.x)

										PP.Multiply(&YplusX, &q.YminusX)
										MM.Multiply(&YminusX, &q.YplusX)
										TT2d.Multiply(&p.t, &q.T2d)
										ZZ2.Multiply(&p.z, &q.Z)

										ZZ2.Add(&ZZ2, &ZZ2)

										v.X.Subtract(&PP, &MM)
										v.Y.Add(&PP, &MM)
										v.Z.Subtract(&ZZ2, &TT2d)
										v.T.Add(&ZZ2, &TT2d)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:309
	// _ = "end of CoverTab[9183]"
}

func (v *projP1xP1) AddAffine(p *Point, q *affineCached) *projP1xP1 {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:312
	_go_fuzz_dep_.CoverTab[9184]++
										var YplusX, YminusX, PP, MM, TT2d, Z2 field.Element

										YplusX.Add(&p.y, &p.x)
										YminusX.Subtract(&p.y, &p.x)

										PP.Multiply(&YplusX, &q.YplusX)
										MM.Multiply(&YminusX, &q.YminusX)
										TT2d.Multiply(&p.t, &q.T2d)

										Z2.Add(&p.z, &p.z)

										v.X.Subtract(&PP, &MM)
										v.Y.Add(&PP, &MM)
										v.Z.Add(&Z2, &TT2d)
										v.T.Subtract(&Z2, &TT2d)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:328
	// _ = "end of CoverTab[9184]"
}

func (v *projP1xP1) SubAffine(p *Point, q *affineCached) *projP1xP1 {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:331
	_go_fuzz_dep_.CoverTab[9185]++
										var YplusX, YminusX, PP, MM, TT2d, Z2 field.Element

										YplusX.Add(&p.y, &p.x)
										YminusX.Subtract(&p.y, &p.x)

										PP.Multiply(&YplusX, &q.YminusX)
										MM.Multiply(&YminusX, &q.YplusX)
										TT2d.Multiply(&p.t, &q.T2d)

										Z2.Add(&p.z, &p.z)

										v.X.Subtract(&PP, &MM)
										v.Y.Add(&PP, &MM)
										v.Z.Subtract(&Z2, &TT2d)
										v.T.Add(&Z2, &TT2d)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:347
	// _ = "end of CoverTab[9185]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:352
func (v *projP1xP1) Double(p *projP2) *projP1xP1 {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:352
	_go_fuzz_dep_.CoverTab[9186]++
										var XX, YY, ZZ2, XplusYsq field.Element

										XX.Square(&p.X)
										YY.Square(&p.Y)
										ZZ2.Square(&p.Z)
										ZZ2.Add(&ZZ2, &ZZ2)
										XplusYsq.Add(&p.X, &p.Y)
										XplusYsq.Square(&XplusYsq)

										v.Y.Add(&YY, &XX)
										v.Z.Subtract(&YY, &XX)

										v.X.Subtract(&XplusYsq, &v.Y)
										v.T.Subtract(&ZZ2, &v.Z)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:367
	// _ = "end of CoverTab[9186]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:372
// Negate sets v = -p, and returns v.
func (v *Point) Negate(p *Point) *Point {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:373
	_go_fuzz_dep_.CoverTab[9187]++
										checkInitialized(p)
										v.x.Negate(&p.x)
										v.y.Set(&p.y)
										v.z.Set(&p.z)
										v.t.Negate(&p.t)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:379
	// _ = "end of CoverTab[9187]"
}

// Equal returns 1 if v is equivalent to u, and 0 otherwise.
func (v *Point) Equal(u *Point) int {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:383
	_go_fuzz_dep_.CoverTab[9188]++
										checkInitialized(v, u)

										var t1, t2, t3, t4 field.Element
										t1.Multiply(&v.x, &u.z)
										t2.Multiply(&u.x, &v.z)
										t3.Multiply(&v.y, &u.z)
										t4.Multiply(&u.y, &v.z)

										return t1.Equal(&t2) & t3.Equal(&t4)
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:392
	// _ = "end of CoverTab[9188]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:397
// Select sets v to a if cond == 1 and to b if cond == 0.
func (v *projCached) Select(a, b *projCached, cond int) *projCached {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:398
	_go_fuzz_dep_.CoverTab[9189]++
										v.YplusX.Select(&a.YplusX, &b.YplusX, cond)
										v.YminusX.Select(&a.YminusX, &b.YminusX, cond)
										v.Z.Select(&a.Z, &b.Z, cond)
										v.T2d.Select(&a.T2d, &b.T2d, cond)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:403
	// _ = "end of CoverTab[9189]"
}

// Select sets v to a if cond == 1 and to b if cond == 0.
func (v *affineCached) Select(a, b *affineCached, cond int) *affineCached {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:407
	_go_fuzz_dep_.CoverTab[9190]++
										v.YplusX.Select(&a.YplusX, &b.YplusX, cond)
										v.YminusX.Select(&a.YminusX, &b.YminusX, cond)
										v.T2d.Select(&a.T2d, &b.T2d, cond)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:411
	// _ = "end of CoverTab[9190]"
}

// CondNeg negates v if cond == 1 and leaves it unchanged if cond == 0.
func (v *projCached) CondNeg(cond int) *projCached {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:415
	_go_fuzz_dep_.CoverTab[9191]++
										v.YplusX.Swap(&v.YminusX, cond)
										v.T2d.Select(new(field.Element).Negate(&v.T2d), &v.T2d, cond)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:418
	// _ = "end of CoverTab[9191]"
}

// CondNeg negates v if cond == 1 and leaves it unchanged if cond == 0.
func (v *affineCached) CondNeg(cond int) *affineCached {
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:422
	_go_fuzz_dep_.CoverTab[9192]++
										v.YplusX.Swap(&v.YminusX, cond)
										v.T2d.Select(new(field.Element).Negate(&v.T2d), &v.T2d, cond)
										return v
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:425
	// _ = "end of CoverTab[9192]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:426
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/edwards25519/edwards25519.go:426
var _ = _go_fuzz_dep_.CoverTab
