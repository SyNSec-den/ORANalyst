// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/elliptic/nistec.go:5
package elliptic

//line /usr/local/go/src/crypto/elliptic/nistec.go:5
import (
//line /usr/local/go/src/crypto/elliptic/nistec.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/elliptic/nistec.go:5
)
//line /usr/local/go/src/crypto/elliptic/nistec.go:5
import (
//line /usr/local/go/src/crypto/elliptic/nistec.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/elliptic/nistec.go:5
)

import (
	"crypto/internal/nistec"
	"errors"
	"math/big"
)

var p224 = &nistCurve[*nistec.P224Point]{
	newPoint: nistec.NewP224Point,
}

func initP224() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:17
	_go_fuzz_dep_.CoverTab[7002]++
							p224.params = &CurveParams{
		Name:		"P-224",
		BitSize:	224,

		P:	bigFromDecimal("26959946667150639794667015087019630673557916260026308143510066298881"),
		N:	bigFromDecimal("26959946667150639794667015087019625940457807714424391721682722368061"),
		B:	bigFromHex("b4050a850c04b3abf54132565044b0b7d7bfd8ba270b39432355ffb4"),
		Gx:	bigFromHex("b70e0cbd6bb4bf7f321390b94a03c1d356c21122343280d6115c1d21"),
		Gy:	bigFromHex("bd376388b5f723fb4c22dfe6cd4375a05a07476444d5819985007e34"),
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:27
	// _ = "end of CoverTab[7002]"
}

type p256Curve struct {
	nistCurve[*nistec.P256Point]
}

var p256 = &p256Curve{nistCurve[*nistec.P256Point]{
	newPoint: nistec.NewP256Point,
}}

func initP256() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:38
	_go_fuzz_dep_.CoverTab[7003]++
							p256.params = &CurveParams{
		Name:		"P-256",
		BitSize:	256,

		P:	bigFromDecimal("115792089210356248762697446949407573530086143415290314195533631308867097853951"),
		N:	bigFromDecimal("115792089210356248762697446949407573529996955224135760342422259061068512044369"),
		B:	bigFromHex("5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b"),
		Gx:	bigFromHex("6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296"),
		Gy:	bigFromHex("4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5"),
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:48
	// _ = "end of CoverTab[7003]"
}

var p384 = &nistCurve[*nistec.P384Point]{
	newPoint: nistec.NewP384Point,
}

func initP384() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:55
	_go_fuzz_dep_.CoverTab[7004]++
							p384.params = &CurveParams{
		Name:		"P-384",
		BitSize:	384,

		P: bigFromDecimal("394020061963944792122790401001436138050797392704654" +
			"46667948293404245721771496870329047266088258938001861606973112319"),
		N: bigFromDecimal("394020061963944792122790401001436138050797392704654" +
			"46667946905279627659399113263569398956308152294913554433653942643"),
		B: bigFromHex("b3312fa7e23ee7e4988e056be3f82d19181d9c6efe8141120314088" +
			"f5013875ac656398d8a2ed19d2a85c8edd3ec2aef"),
		Gx: bigFromHex("aa87ca22be8b05378eb1c71ef320ad746e1d3b628ba79b9859f741" +
			"e082542a385502f25dbf55296c3a545e3872760ab7"),
		Gy: bigFromHex("3617de4a96262c6f5d9e98bf9292dc29f8f41dbd289a147ce9da31" +
			"13b5f0b8c00a60b1ce1d7e819d7a431d7c90ea0e5f"),
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:70
	// _ = "end of CoverTab[7004]"
}

var p521 = &nistCurve[*nistec.P521Point]{
	newPoint: nistec.NewP521Point,
}

func initP521() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:77
	_go_fuzz_dep_.CoverTab[7005]++
							p521.params = &CurveParams{
		Name:		"P-521",
		BitSize:	521,

		P: bigFromDecimal("68647976601306097149819007990813932172694353001433" +
			"0540939446345918554318339765605212255964066145455497729631139148" +
			"0858037121987999716643812574028291115057151"),
		N: bigFromDecimal("68647976601306097149819007990813932172694353001433" +
			"0540939446345918554318339765539424505774633321719753296399637136" +
			"3321113864768612440380340372808892707005449"),
		B: bigFromHex("0051953eb9618e1c9a1f929a21a0b68540eea2da725b99b315f3b8" +
			"b489918ef109e156193951ec7e937b1652c0bd3bb1bf073573df883d2c34f1ef" +
			"451fd46b503f00"),
		Gx: bigFromHex("00c6858e06b70404e9cd9e3ecb662395b4429c648139053fb521f8" +
			"28af606b4d3dbaa14b5e77efe75928fe1dc127a2ffa8de3348b3c1856a429bf9" +
			"7e7e31c2e5bd66"),
		Gy: bigFromHex("011839296a789a3bc0045c8a5fb42c7d1bd998f54449579b446817" +
			"afbd17273e662c97ee72995ef42640c550b9013fad0761353c7086a272c24088" +
			"be94769fd16650"),
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:97
	// _ = "end of CoverTab[7005]"
}

// nistCurve is a Curve implementation based on a nistec Point.
//line /usr/local/go/src/crypto/elliptic/nistec.go:100
//
//line /usr/local/go/src/crypto/elliptic/nistec.go:100
// It's a wrapper that exposes the big.Int-based Curve interface and encodes the
//line /usr/local/go/src/crypto/elliptic/nistec.go:100
// legacy idiosyncrasies it requires, such as invalid and infinity point
//line /usr/local/go/src/crypto/elliptic/nistec.go:100
// handling.
//line /usr/local/go/src/crypto/elliptic/nistec.go:100
//
//line /usr/local/go/src/crypto/elliptic/nistec.go:100
// To interact with the nistec package, points are encoded into and decoded from
//line /usr/local/go/src/crypto/elliptic/nistec.go:100
// properly formatted byte slices. All big.Int use is limited to this package.
//line /usr/local/go/src/crypto/elliptic/nistec.go:100
// Encoding and decoding is 1/1000th of the runtime of a scalar multiplication,
//line /usr/local/go/src/crypto/elliptic/nistec.go:100
// so the overhead is acceptable.
//line /usr/local/go/src/crypto/elliptic/nistec.go:110
type nistCurve[Point nistPoint[Point]] struct {
	newPoint	func() Point
	params		*CurveParams
}

// nistPoint is a generic constraint for the nistec Point types.
type nistPoint[T any] interface {
	Bytes() []byte
	SetBytes([]byte) (T, error)
	Add(T, T) T
	Double(T) T
	ScalarMult(T, []byte) (T, error)
	ScalarBaseMult([]byte) (T, error)
}

func (curve *nistCurve[Point]) Params() *CurveParams {
//line /usr/local/go/src/crypto/elliptic/nistec.go:125
	_go_fuzz_dep_.CoverTab[7006]++
							return curve.params
//line /usr/local/go/src/crypto/elliptic/nistec.go:126
	// _ = "end of CoverTab[7006]"
}

func (curve *nistCurve[Point]) IsOnCurve(x, y *big.Int) bool {
//line /usr/local/go/src/crypto/elliptic/nistec.go:129
	_go_fuzz_dep_.CoverTab[7007]++

//line /usr/local/go/src/crypto/elliptic/nistec.go:132
	if x.Sign() == 0 && func() bool {
//line /usr/local/go/src/crypto/elliptic/nistec.go:132
		_go_fuzz_dep_.CoverTab[7009]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:132
		return y.Sign() == 0
//line /usr/local/go/src/crypto/elliptic/nistec.go:132
		// _ = "end of CoverTab[7009]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:132
	}() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:132
		_go_fuzz_dep_.CoverTab[7010]++
								return false
//line /usr/local/go/src/crypto/elliptic/nistec.go:133
		// _ = "end of CoverTab[7010]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:134
		_go_fuzz_dep_.CoverTab[7011]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:134
		// _ = "end of CoverTab[7011]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:134
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:134
	// _ = "end of CoverTab[7007]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:134
	_go_fuzz_dep_.CoverTab[7008]++
							_, err := curve.pointFromAffine(x, y)
							return err == nil
//line /usr/local/go/src/crypto/elliptic/nistec.go:136
	// _ = "end of CoverTab[7008]"
}

func (curve *nistCurve[Point]) pointFromAffine(x, y *big.Int) (p Point, err error) {
//line /usr/local/go/src/crypto/elliptic/nistec.go:139
	_go_fuzz_dep_.CoverTab[7012]++

//line /usr/local/go/src/crypto/elliptic/nistec.go:142
	if x.Sign() == 0 && func() bool {
//line /usr/local/go/src/crypto/elliptic/nistec.go:142
		_go_fuzz_dep_.CoverTab[7016]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:142
		return y.Sign() == 0
//line /usr/local/go/src/crypto/elliptic/nistec.go:142
		// _ = "end of CoverTab[7016]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:142
	}() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:142
		_go_fuzz_dep_.CoverTab[7017]++
								return curve.newPoint(), nil
//line /usr/local/go/src/crypto/elliptic/nistec.go:143
		// _ = "end of CoverTab[7017]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:144
		_go_fuzz_dep_.CoverTab[7018]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:144
		// _ = "end of CoverTab[7018]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:144
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:144
	// _ = "end of CoverTab[7012]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:144
	_go_fuzz_dep_.CoverTab[7013]++

							if x.Sign() < 0 || func() bool {
//line /usr/local/go/src/crypto/elliptic/nistec.go:146
		_go_fuzz_dep_.CoverTab[7019]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:146
		return y.Sign() < 0
//line /usr/local/go/src/crypto/elliptic/nistec.go:146
		// _ = "end of CoverTab[7019]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:146
	}() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:146
		_go_fuzz_dep_.CoverTab[7020]++
								return p, errors.New("negative coordinate")
//line /usr/local/go/src/crypto/elliptic/nistec.go:147
		// _ = "end of CoverTab[7020]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:148
		_go_fuzz_dep_.CoverTab[7021]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:148
		// _ = "end of CoverTab[7021]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:148
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:148
	// _ = "end of CoverTab[7013]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:148
	_go_fuzz_dep_.CoverTab[7014]++
							if x.BitLen() > curve.params.BitSize || func() bool {
//line /usr/local/go/src/crypto/elliptic/nistec.go:149
		_go_fuzz_dep_.CoverTab[7022]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:149
		return y.BitLen() > curve.params.BitSize
//line /usr/local/go/src/crypto/elliptic/nistec.go:149
		// _ = "end of CoverTab[7022]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:149
	}() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:149
		_go_fuzz_dep_.CoverTab[7023]++
								return p, errors.New("overflowing coordinate")
//line /usr/local/go/src/crypto/elliptic/nistec.go:150
		// _ = "end of CoverTab[7023]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:151
		_go_fuzz_dep_.CoverTab[7024]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:151
		// _ = "end of CoverTab[7024]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:151
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:151
	// _ = "end of CoverTab[7014]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:151
	_go_fuzz_dep_.CoverTab[7015]++

							byteLen := (curve.params.BitSize + 7) / 8
							buf := make([]byte, 1+2*byteLen)
							buf[0] = 4
							x.FillBytes(buf[1 : 1+byteLen])
							y.FillBytes(buf[1+byteLen : 1+2*byteLen])
							return curve.newPoint().SetBytes(buf)
//line /usr/local/go/src/crypto/elliptic/nistec.go:158
	// _ = "end of CoverTab[7015]"
}

func (curve *nistCurve[Point]) pointToAffine(p Point) (x, y *big.Int) {
//line /usr/local/go/src/crypto/elliptic/nistec.go:161
	_go_fuzz_dep_.CoverTab[7025]++
							out := p.Bytes()
							if len(out) == 1 && func() bool {
//line /usr/local/go/src/crypto/elliptic/nistec.go:163
		_go_fuzz_dep_.CoverTab[7027]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:163
		return out[0] == 0
//line /usr/local/go/src/crypto/elliptic/nistec.go:163
		// _ = "end of CoverTab[7027]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:163
	}() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:163
		_go_fuzz_dep_.CoverTab[7028]++

//line /usr/local/go/src/crypto/elliptic/nistec.go:166
		return new(big.Int), new(big.Int)
//line /usr/local/go/src/crypto/elliptic/nistec.go:166
		// _ = "end of CoverTab[7028]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:167
		_go_fuzz_dep_.CoverTab[7029]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:167
		// _ = "end of CoverTab[7029]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:167
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:167
	// _ = "end of CoverTab[7025]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:167
	_go_fuzz_dep_.CoverTab[7026]++
							byteLen := (curve.params.BitSize + 7) / 8
							x = new(big.Int).SetBytes(out[1 : 1+byteLen])
							y = new(big.Int).SetBytes(out[1+byteLen:])
							return x, y
//line /usr/local/go/src/crypto/elliptic/nistec.go:171
	// _ = "end of CoverTab[7026]"
}

func (curve *nistCurve[Point]) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/nistec.go:174
	_go_fuzz_dep_.CoverTab[7030]++
							p1, err := curve.pointFromAffine(x1, y1)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:176
		_go_fuzz_dep_.CoverTab[7033]++
								panic("crypto/elliptic: Add was called on an invalid point")
//line /usr/local/go/src/crypto/elliptic/nistec.go:177
		// _ = "end of CoverTab[7033]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:178
		_go_fuzz_dep_.CoverTab[7034]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:178
		// _ = "end of CoverTab[7034]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:178
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:178
	// _ = "end of CoverTab[7030]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:178
	_go_fuzz_dep_.CoverTab[7031]++
							p2, err := curve.pointFromAffine(x2, y2)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:180
		_go_fuzz_dep_.CoverTab[7035]++
								panic("crypto/elliptic: Add was called on an invalid point")
//line /usr/local/go/src/crypto/elliptic/nistec.go:181
		// _ = "end of CoverTab[7035]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:182
		_go_fuzz_dep_.CoverTab[7036]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:182
		// _ = "end of CoverTab[7036]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:182
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:182
	// _ = "end of CoverTab[7031]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:182
	_go_fuzz_dep_.CoverTab[7032]++
							return curve.pointToAffine(p1.Add(p1, p2))
//line /usr/local/go/src/crypto/elliptic/nistec.go:183
	// _ = "end of CoverTab[7032]"
}

func (curve *nistCurve[Point]) Double(x1, y1 *big.Int) (*big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/nistec.go:186
	_go_fuzz_dep_.CoverTab[7037]++
							p, err := curve.pointFromAffine(x1, y1)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:188
		_go_fuzz_dep_.CoverTab[7039]++
								panic("crypto/elliptic: Double was called on an invalid point")
//line /usr/local/go/src/crypto/elliptic/nistec.go:189
		// _ = "end of CoverTab[7039]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:190
		_go_fuzz_dep_.CoverTab[7040]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:190
		// _ = "end of CoverTab[7040]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:190
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:190
	// _ = "end of CoverTab[7037]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:190
	_go_fuzz_dep_.CoverTab[7038]++
							return curve.pointToAffine(p.Double(p))
//line /usr/local/go/src/crypto/elliptic/nistec.go:191
	// _ = "end of CoverTab[7038]"
}

// normalizeScalar brings the scalar within the byte size of the order of the
//line /usr/local/go/src/crypto/elliptic/nistec.go:194
// curve, as expected by the nistec scalar multiplication functions.
//line /usr/local/go/src/crypto/elliptic/nistec.go:196
func (curve *nistCurve[Point]) normalizeScalar(scalar []byte) []byte {
//line /usr/local/go/src/crypto/elliptic/nistec.go:196
	_go_fuzz_dep_.CoverTab[7041]++
							byteSize := (curve.params.N.BitLen() + 7) / 8
							if len(scalar) == byteSize {
//line /usr/local/go/src/crypto/elliptic/nistec.go:198
		_go_fuzz_dep_.CoverTab[7044]++
								return scalar
//line /usr/local/go/src/crypto/elliptic/nistec.go:199
		// _ = "end of CoverTab[7044]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:200
		_go_fuzz_dep_.CoverTab[7045]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:200
		// _ = "end of CoverTab[7045]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:200
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:200
	// _ = "end of CoverTab[7041]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:200
	_go_fuzz_dep_.CoverTab[7042]++
							s := new(big.Int).SetBytes(scalar)
							if len(scalar) > byteSize {
//line /usr/local/go/src/crypto/elliptic/nistec.go:202
		_go_fuzz_dep_.CoverTab[7046]++
								s.Mod(s, curve.params.N)
//line /usr/local/go/src/crypto/elliptic/nistec.go:203
		// _ = "end of CoverTab[7046]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:204
		_go_fuzz_dep_.CoverTab[7047]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:204
		// _ = "end of CoverTab[7047]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:204
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:204
	// _ = "end of CoverTab[7042]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:204
	_go_fuzz_dep_.CoverTab[7043]++
							out := make([]byte, byteSize)
							return s.FillBytes(out)
//line /usr/local/go/src/crypto/elliptic/nistec.go:206
	// _ = "end of CoverTab[7043]"
}

func (curve *nistCurve[Point]) ScalarMult(Bx, By *big.Int, scalar []byte) (*big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/nistec.go:209
	_go_fuzz_dep_.CoverTab[7048]++
							p, err := curve.pointFromAffine(Bx, By)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:211
		_go_fuzz_dep_.CoverTab[7051]++
								panic("crypto/elliptic: ScalarMult was called on an invalid point")
//line /usr/local/go/src/crypto/elliptic/nistec.go:212
		// _ = "end of CoverTab[7051]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:213
		_go_fuzz_dep_.CoverTab[7052]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:213
		// _ = "end of CoverTab[7052]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:213
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:213
	// _ = "end of CoverTab[7048]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:213
	_go_fuzz_dep_.CoverTab[7049]++
							scalar = curve.normalizeScalar(scalar)
							p, err = p.ScalarMult(p, scalar)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:216
		_go_fuzz_dep_.CoverTab[7053]++
								panic("crypto/elliptic: nistec rejected normalized scalar")
//line /usr/local/go/src/crypto/elliptic/nistec.go:217
		// _ = "end of CoverTab[7053]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:218
		_go_fuzz_dep_.CoverTab[7054]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:218
		// _ = "end of CoverTab[7054]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:218
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:218
	// _ = "end of CoverTab[7049]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:218
	_go_fuzz_dep_.CoverTab[7050]++
							return curve.pointToAffine(p)
//line /usr/local/go/src/crypto/elliptic/nistec.go:219
	// _ = "end of CoverTab[7050]"
}

func (curve *nistCurve[Point]) ScalarBaseMult(scalar []byte) (*big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/nistec.go:222
	_go_fuzz_dep_.CoverTab[7055]++
							scalar = curve.normalizeScalar(scalar)
							p, err := curve.newPoint().ScalarBaseMult(scalar)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:225
		_go_fuzz_dep_.CoverTab[7057]++
								panic("crypto/elliptic: nistec rejected normalized scalar")
//line /usr/local/go/src/crypto/elliptic/nistec.go:226
		// _ = "end of CoverTab[7057]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:227
		_go_fuzz_dep_.CoverTab[7058]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:227
		// _ = "end of CoverTab[7058]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:227
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:227
	// _ = "end of CoverTab[7055]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:227
	_go_fuzz_dep_.CoverTab[7056]++
							return curve.pointToAffine(p)
//line /usr/local/go/src/crypto/elliptic/nistec.go:228
	// _ = "end of CoverTab[7056]"
}

// CombinedMult returns [s1]G + [s2]P where G is the generator. It's used
//line /usr/local/go/src/crypto/elliptic/nistec.go:231
// through an interface upgrade in crypto/ecdsa.
//line /usr/local/go/src/crypto/elliptic/nistec.go:233
func (curve *nistCurve[Point]) CombinedMult(Px, Py *big.Int, s1, s2 []byte) (x, y *big.Int) {
//line /usr/local/go/src/crypto/elliptic/nistec.go:233
	_go_fuzz_dep_.CoverTab[7059]++
							s1 = curve.normalizeScalar(s1)
							q, err := curve.newPoint().ScalarBaseMult(s1)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:236
		_go_fuzz_dep_.CoverTab[7063]++
								panic("crypto/elliptic: nistec rejected normalized scalar")
//line /usr/local/go/src/crypto/elliptic/nistec.go:237
		// _ = "end of CoverTab[7063]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:238
		_go_fuzz_dep_.CoverTab[7064]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:238
		// _ = "end of CoverTab[7064]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:238
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:238
	// _ = "end of CoverTab[7059]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:238
	_go_fuzz_dep_.CoverTab[7060]++
							p, err := curve.pointFromAffine(Px, Py)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:240
		_go_fuzz_dep_.CoverTab[7065]++
								panic("crypto/elliptic: CombinedMult was called on an invalid point")
//line /usr/local/go/src/crypto/elliptic/nistec.go:241
		// _ = "end of CoverTab[7065]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:242
		_go_fuzz_dep_.CoverTab[7066]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:242
		// _ = "end of CoverTab[7066]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:242
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:242
	// _ = "end of CoverTab[7060]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:242
	_go_fuzz_dep_.CoverTab[7061]++
							s2 = curve.normalizeScalar(s2)
							p, err = p.ScalarMult(p, s2)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:245
		_go_fuzz_dep_.CoverTab[7067]++
								panic("crypto/elliptic: nistec rejected normalized scalar")
//line /usr/local/go/src/crypto/elliptic/nistec.go:246
		// _ = "end of CoverTab[7067]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:247
		_go_fuzz_dep_.CoverTab[7068]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:247
		// _ = "end of CoverTab[7068]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:247
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:247
	// _ = "end of CoverTab[7061]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:247
	_go_fuzz_dep_.CoverTab[7062]++
							return curve.pointToAffine(p.Add(p, q))
//line /usr/local/go/src/crypto/elliptic/nistec.go:248
	// _ = "end of CoverTab[7062]"
}

func (curve *nistCurve[Point]) Unmarshal(data []byte) (x, y *big.Int) {
//line /usr/local/go/src/crypto/elliptic/nistec.go:251
	_go_fuzz_dep_.CoverTab[7069]++
							if len(data) == 0 || func() bool {
//line /usr/local/go/src/crypto/elliptic/nistec.go:252
		_go_fuzz_dep_.CoverTab[7072]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:252
		return data[0] != 4
//line /usr/local/go/src/crypto/elliptic/nistec.go:252
		// _ = "end of CoverTab[7072]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:252
	}() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:252
		_go_fuzz_dep_.CoverTab[7073]++
								return nil, nil
//line /usr/local/go/src/crypto/elliptic/nistec.go:253
		// _ = "end of CoverTab[7073]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:254
		_go_fuzz_dep_.CoverTab[7074]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:254
		// _ = "end of CoverTab[7074]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:254
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:254
	// _ = "end of CoverTab[7069]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:254
	_go_fuzz_dep_.CoverTab[7070]++

							_, err := curve.newPoint().SetBytes(data)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:257
		_go_fuzz_dep_.CoverTab[7075]++
								return nil, nil
//line /usr/local/go/src/crypto/elliptic/nistec.go:258
		// _ = "end of CoverTab[7075]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:259
		_go_fuzz_dep_.CoverTab[7076]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:259
		// _ = "end of CoverTab[7076]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:259
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:259
	// _ = "end of CoverTab[7070]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:259
	_go_fuzz_dep_.CoverTab[7071]++

//line /usr/local/go/src/crypto/elliptic/nistec.go:263
	byteLen := (curve.params.BitSize + 7) / 8
							x = new(big.Int).SetBytes(data[1 : 1+byteLen])
							y = new(big.Int).SetBytes(data[1+byteLen:])
							return x, y
//line /usr/local/go/src/crypto/elliptic/nistec.go:266
	// _ = "end of CoverTab[7071]"
}

func (curve *nistCurve[Point]) UnmarshalCompressed(data []byte) (x, y *big.Int) {
//line /usr/local/go/src/crypto/elliptic/nistec.go:269
	_go_fuzz_dep_.CoverTab[7077]++
							if len(data) == 0 || func() bool {
//line /usr/local/go/src/crypto/elliptic/nistec.go:270
		_go_fuzz_dep_.CoverTab[7080]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:270
		return (data[0] != 2 && func() bool {
//line /usr/local/go/src/crypto/elliptic/nistec.go:270
			_go_fuzz_dep_.CoverTab[7081]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:270
			return data[0] != 3
//line /usr/local/go/src/crypto/elliptic/nistec.go:270
			// _ = "end of CoverTab[7081]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:270
		}())
//line /usr/local/go/src/crypto/elliptic/nistec.go:270
		// _ = "end of CoverTab[7080]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:270
	}() {
//line /usr/local/go/src/crypto/elliptic/nistec.go:270
		_go_fuzz_dep_.CoverTab[7082]++
								return nil, nil
//line /usr/local/go/src/crypto/elliptic/nistec.go:271
		// _ = "end of CoverTab[7082]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:272
		_go_fuzz_dep_.CoverTab[7083]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:272
		// _ = "end of CoverTab[7083]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:272
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:272
	// _ = "end of CoverTab[7077]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:272
	_go_fuzz_dep_.CoverTab[7078]++
							p, err := curve.newPoint().SetBytes(data)
							if err != nil {
//line /usr/local/go/src/crypto/elliptic/nistec.go:274
		_go_fuzz_dep_.CoverTab[7084]++
								return nil, nil
//line /usr/local/go/src/crypto/elliptic/nistec.go:275
		// _ = "end of CoverTab[7084]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:276
		_go_fuzz_dep_.CoverTab[7085]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:276
		// _ = "end of CoverTab[7085]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:276
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:276
	// _ = "end of CoverTab[7078]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:276
	_go_fuzz_dep_.CoverTab[7079]++
							return curve.pointToAffine(p)
//line /usr/local/go/src/crypto/elliptic/nistec.go:277
	// _ = "end of CoverTab[7079]"
}

func bigFromDecimal(s string) *big.Int {
//line /usr/local/go/src/crypto/elliptic/nistec.go:280
	_go_fuzz_dep_.CoverTab[7086]++
							b, ok := new(big.Int).SetString(s, 10)
							if !ok {
//line /usr/local/go/src/crypto/elliptic/nistec.go:282
		_go_fuzz_dep_.CoverTab[7088]++
								panic("crypto/elliptic: internal error: invalid encoding")
//line /usr/local/go/src/crypto/elliptic/nistec.go:283
		// _ = "end of CoverTab[7088]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:284
		_go_fuzz_dep_.CoverTab[7089]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:284
		// _ = "end of CoverTab[7089]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:284
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:284
	// _ = "end of CoverTab[7086]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:284
	_go_fuzz_dep_.CoverTab[7087]++
							return b
//line /usr/local/go/src/crypto/elliptic/nistec.go:285
	// _ = "end of CoverTab[7087]"
}

func bigFromHex(s string) *big.Int {
//line /usr/local/go/src/crypto/elliptic/nistec.go:288
	_go_fuzz_dep_.CoverTab[7090]++
							b, ok := new(big.Int).SetString(s, 16)
							if !ok {
//line /usr/local/go/src/crypto/elliptic/nistec.go:290
		_go_fuzz_dep_.CoverTab[7092]++
								panic("crypto/elliptic: internal error: invalid encoding")
//line /usr/local/go/src/crypto/elliptic/nistec.go:291
		// _ = "end of CoverTab[7092]"
	} else {
//line /usr/local/go/src/crypto/elliptic/nistec.go:292
		_go_fuzz_dep_.CoverTab[7093]++
//line /usr/local/go/src/crypto/elliptic/nistec.go:292
		// _ = "end of CoverTab[7093]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:292
	}
//line /usr/local/go/src/crypto/elliptic/nistec.go:292
	// _ = "end of CoverTab[7090]"
//line /usr/local/go/src/crypto/elliptic/nistec.go:292
	_go_fuzz_dep_.CoverTab[7091]++
							return b
//line /usr/local/go/src/crypto/elliptic/nistec.go:293
	// _ = "end of CoverTab[7091]"
}

//line /usr/local/go/src/crypto/elliptic/nistec.go:294
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/elliptic/nistec.go:294
var _ = _go_fuzz_dep_.CoverTab
