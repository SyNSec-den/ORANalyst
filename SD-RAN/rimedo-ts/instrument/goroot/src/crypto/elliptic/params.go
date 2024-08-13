// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/elliptic/params.go:5
package elliptic

//line /usr/local/go/src/crypto/elliptic/params.go:5
import (
//line /usr/local/go/src/crypto/elliptic/params.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/elliptic/params.go:5
)
//line /usr/local/go/src/crypto/elliptic/params.go:5
import (
//line /usr/local/go/src/crypto/elliptic/params.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/elliptic/params.go:5
)

import "math/big"

// CurveParams contains the parameters of an elliptic curve and also provides
//line /usr/local/go/src/crypto/elliptic/params.go:9
// a generic, non-constant time implementation of Curve.
//line /usr/local/go/src/crypto/elliptic/params.go:9
//
//line /usr/local/go/src/crypto/elliptic/params.go:9
// Note: Custom curves (those not returned by P224(), P256(), P384(), and P521())
//line /usr/local/go/src/crypto/elliptic/params.go:9
// are not guaranteed to provide any security property.
//line /usr/local/go/src/crypto/elliptic/params.go:14
type CurveParams struct {
	P	*big.Int	// the order of the underlying field
	N	*big.Int	// the order of the base point
	B	*big.Int	// the constant of the curve equation
	Gx, Gy	*big.Int	// (x,y) of the base point
	BitSize	int		// the size of the underlying field
	Name	string		// the canonical name of the curve
}

func (curve *CurveParams) Params() *CurveParams {
//line /usr/local/go/src/crypto/elliptic/params.go:23
	_go_fuzz_dep_.CoverTab[7104]++
							return curve
//line /usr/local/go/src/crypto/elliptic/params.go:24
	// _ = "end of CoverTab[7104]"
}

//line /usr/local/go/src/crypto/elliptic/params.go:34
// polynomial returns x³ - 3x + b.
func (curve *CurveParams) polynomial(x *big.Int) *big.Int {
//line /usr/local/go/src/crypto/elliptic/params.go:35
	_go_fuzz_dep_.CoverTab[7105]++
							x3 := new(big.Int).Mul(x, x)
							x3.Mul(x3, x)

							threeX := new(big.Int).Lsh(x, 1)
							threeX.Add(threeX, x)

							x3.Sub(x3, threeX)
							x3.Add(x3, curve.B)
							x3.Mod(x3, curve.P)

							return x3
//line /usr/local/go/src/crypto/elliptic/params.go:46
	// _ = "end of CoverTab[7105]"
}

// IsOnCurve implements Curve.IsOnCurve.
//line /usr/local/go/src/crypto/elliptic/params.go:49
//
//line /usr/local/go/src/crypto/elliptic/params.go:49
// Note: the CurveParams methods are not guaranteed to
//line /usr/local/go/src/crypto/elliptic/params.go:49
// provide any security property. For ECDH, use the crypto/ecdh package.
//line /usr/local/go/src/crypto/elliptic/params.go:49
// For ECDSA, use the crypto/ecdsa package with a Curve value returned directly
//line /usr/local/go/src/crypto/elliptic/params.go:49
// from P224(), P256(), P384(), or P521().
//line /usr/local/go/src/crypto/elliptic/params.go:55
func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool {
//line /usr/local/go/src/crypto/elliptic/params.go:55
	_go_fuzz_dep_.CoverTab[7106]++

//line /usr/local/go/src/crypto/elliptic/params.go:58
	if specific, ok := matchesSpecificCurve(curve); ok {
//line /usr/local/go/src/crypto/elliptic/params.go:58
		_go_fuzz_dep_.CoverTab[7109]++
								return specific.IsOnCurve(x, y)
//line /usr/local/go/src/crypto/elliptic/params.go:59
		// _ = "end of CoverTab[7109]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:60
		_go_fuzz_dep_.CoverTab[7110]++
//line /usr/local/go/src/crypto/elliptic/params.go:60
		// _ = "end of CoverTab[7110]"
//line /usr/local/go/src/crypto/elliptic/params.go:60
	}
//line /usr/local/go/src/crypto/elliptic/params.go:60
	// _ = "end of CoverTab[7106]"
//line /usr/local/go/src/crypto/elliptic/params.go:60
	_go_fuzz_dep_.CoverTab[7107]++

							if x.Sign() < 0 || func() bool {
//line /usr/local/go/src/crypto/elliptic/params.go:62
		_go_fuzz_dep_.CoverTab[7111]++
//line /usr/local/go/src/crypto/elliptic/params.go:62
		return x.Cmp(curve.P) >= 0
//line /usr/local/go/src/crypto/elliptic/params.go:62
		// _ = "end of CoverTab[7111]"
//line /usr/local/go/src/crypto/elliptic/params.go:62
	}() || func() bool {
//line /usr/local/go/src/crypto/elliptic/params.go:62
		_go_fuzz_dep_.CoverTab[7112]++
//line /usr/local/go/src/crypto/elliptic/params.go:62
		return y.Sign() < 0
								// _ = "end of CoverTab[7112]"
//line /usr/local/go/src/crypto/elliptic/params.go:63
	}() || func() bool {
//line /usr/local/go/src/crypto/elliptic/params.go:63
		_go_fuzz_dep_.CoverTab[7113]++
//line /usr/local/go/src/crypto/elliptic/params.go:63
		return y.Cmp(curve.P) >= 0
//line /usr/local/go/src/crypto/elliptic/params.go:63
		// _ = "end of CoverTab[7113]"
//line /usr/local/go/src/crypto/elliptic/params.go:63
	}() {
//line /usr/local/go/src/crypto/elliptic/params.go:63
		_go_fuzz_dep_.CoverTab[7114]++
								return false
//line /usr/local/go/src/crypto/elliptic/params.go:64
		// _ = "end of CoverTab[7114]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:65
		_go_fuzz_dep_.CoverTab[7115]++
//line /usr/local/go/src/crypto/elliptic/params.go:65
		// _ = "end of CoverTab[7115]"
//line /usr/local/go/src/crypto/elliptic/params.go:65
	}
//line /usr/local/go/src/crypto/elliptic/params.go:65
	// _ = "end of CoverTab[7107]"
//line /usr/local/go/src/crypto/elliptic/params.go:65
	_go_fuzz_dep_.CoverTab[7108]++

//line /usr/local/go/src/crypto/elliptic/params.go:68
	y2 := new(big.Int).Mul(y, y)
							y2.Mod(y2, curve.P)

							return curve.polynomial(x).Cmp(y2) == 0
//line /usr/local/go/src/crypto/elliptic/params.go:71
	// _ = "end of CoverTab[7108]"
}

// zForAffine returns a Jacobian Z value for the affine point (x, y). If x and
//line /usr/local/go/src/crypto/elliptic/params.go:74
// y are zero, it assumes that they represent the point at infinity because (0,
//line /usr/local/go/src/crypto/elliptic/params.go:74
// 0) is not on the any of the curves handled here.
//line /usr/local/go/src/crypto/elliptic/params.go:77
func zForAffine(x, y *big.Int) *big.Int {
//line /usr/local/go/src/crypto/elliptic/params.go:77
	_go_fuzz_dep_.CoverTab[7116]++
							z := new(big.Int)
							if x.Sign() != 0 || func() bool {
//line /usr/local/go/src/crypto/elliptic/params.go:79
		_go_fuzz_dep_.CoverTab[7118]++
//line /usr/local/go/src/crypto/elliptic/params.go:79
		return y.Sign() != 0
//line /usr/local/go/src/crypto/elliptic/params.go:79
		// _ = "end of CoverTab[7118]"
//line /usr/local/go/src/crypto/elliptic/params.go:79
	}() {
//line /usr/local/go/src/crypto/elliptic/params.go:79
		_go_fuzz_dep_.CoverTab[7119]++
								z.SetInt64(1)
//line /usr/local/go/src/crypto/elliptic/params.go:80
		// _ = "end of CoverTab[7119]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:81
		_go_fuzz_dep_.CoverTab[7120]++
//line /usr/local/go/src/crypto/elliptic/params.go:81
		// _ = "end of CoverTab[7120]"
//line /usr/local/go/src/crypto/elliptic/params.go:81
	}
//line /usr/local/go/src/crypto/elliptic/params.go:81
	// _ = "end of CoverTab[7116]"
//line /usr/local/go/src/crypto/elliptic/params.go:81
	_go_fuzz_dep_.CoverTab[7117]++
							return z
//line /usr/local/go/src/crypto/elliptic/params.go:82
	// _ = "end of CoverTab[7117]"
}

// affineFromJacobian reverses the Jacobian transform. See the comment at the
//line /usr/local/go/src/crypto/elliptic/params.go:85
// top of the file. If the point is ∞ it returns 0, 0.
//line /usr/local/go/src/crypto/elliptic/params.go:87
func (curve *CurveParams) affineFromJacobian(x, y, z *big.Int) (xOut, yOut *big.Int) {
//line /usr/local/go/src/crypto/elliptic/params.go:87
	_go_fuzz_dep_.CoverTab[7121]++
							if z.Sign() == 0 {
//line /usr/local/go/src/crypto/elliptic/params.go:88
		_go_fuzz_dep_.CoverTab[7123]++
								return new(big.Int), new(big.Int)
//line /usr/local/go/src/crypto/elliptic/params.go:89
		// _ = "end of CoverTab[7123]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:90
		_go_fuzz_dep_.CoverTab[7124]++
//line /usr/local/go/src/crypto/elliptic/params.go:90
		// _ = "end of CoverTab[7124]"
//line /usr/local/go/src/crypto/elliptic/params.go:90
	}
//line /usr/local/go/src/crypto/elliptic/params.go:90
	// _ = "end of CoverTab[7121]"
//line /usr/local/go/src/crypto/elliptic/params.go:90
	_go_fuzz_dep_.CoverTab[7122]++

							zinv := new(big.Int).ModInverse(z, curve.P)
							zinvsq := new(big.Int).Mul(zinv, zinv)

							xOut = new(big.Int).Mul(x, zinvsq)
							xOut.Mod(xOut, curve.P)
							zinvsq.Mul(zinvsq, zinv)
							yOut = new(big.Int).Mul(y, zinvsq)
							yOut.Mod(yOut, curve.P)
							return
//line /usr/local/go/src/crypto/elliptic/params.go:100
	// _ = "end of CoverTab[7122]"
}

// Add implements Curve.Add.
//line /usr/local/go/src/crypto/elliptic/params.go:103
//
//line /usr/local/go/src/crypto/elliptic/params.go:103
// Note: the CurveParams methods are not guaranteed to
//line /usr/local/go/src/crypto/elliptic/params.go:103
// provide any security property. For ECDH, use the crypto/ecdh package.
//line /usr/local/go/src/crypto/elliptic/params.go:103
// For ECDSA, use the crypto/ecdsa package with a Curve value returned directly
//line /usr/local/go/src/crypto/elliptic/params.go:103
// from P224(), P256(), P384(), or P521().
//line /usr/local/go/src/crypto/elliptic/params.go:109
func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/params.go:109
	_go_fuzz_dep_.CoverTab[7125]++

//line /usr/local/go/src/crypto/elliptic/params.go:112
	if specific, ok := matchesSpecificCurve(curve); ok {
//line /usr/local/go/src/crypto/elliptic/params.go:112
		_go_fuzz_dep_.CoverTab[7127]++
								return specific.Add(x1, y1, x2, y2)
//line /usr/local/go/src/crypto/elliptic/params.go:113
		// _ = "end of CoverTab[7127]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:114
		_go_fuzz_dep_.CoverTab[7128]++
//line /usr/local/go/src/crypto/elliptic/params.go:114
		// _ = "end of CoverTab[7128]"
//line /usr/local/go/src/crypto/elliptic/params.go:114
	}
//line /usr/local/go/src/crypto/elliptic/params.go:114
	// _ = "end of CoverTab[7125]"
//line /usr/local/go/src/crypto/elliptic/params.go:114
	_go_fuzz_dep_.CoverTab[7126]++
							panicIfNotOnCurve(curve, x1, y1)
							panicIfNotOnCurve(curve, x2, y2)

							z1 := zForAffine(x1, y1)
							z2 := zForAffine(x2, y2)
							return curve.affineFromJacobian(curve.addJacobian(x1, y1, z1, x2, y2, z2))
//line /usr/local/go/src/crypto/elliptic/params.go:120
	// _ = "end of CoverTab[7126]"
}

// addJacobian takes two points in Jacobian coordinates, (x1, y1, z1) and
//line /usr/local/go/src/crypto/elliptic/params.go:123
// (x2, y2, z2) and returns their sum, also in Jacobian form.
//line /usr/local/go/src/crypto/elliptic/params.go:125
func (curve *CurveParams) addJacobian(x1, y1, z1, x2, y2, z2 *big.Int) (*big.Int, *big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/params.go:125
	_go_fuzz_dep_.CoverTab[7129]++

							x3, y3, z3 := new(big.Int), new(big.Int), new(big.Int)
							if z1.Sign() == 0 {
//line /usr/local/go/src/crypto/elliptic/params.go:128
		_go_fuzz_dep_.CoverTab[7135]++
								x3.Set(x2)
								y3.Set(y2)
								z3.Set(z2)
								return x3, y3, z3
//line /usr/local/go/src/crypto/elliptic/params.go:132
		// _ = "end of CoverTab[7135]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:133
		_go_fuzz_dep_.CoverTab[7136]++
//line /usr/local/go/src/crypto/elliptic/params.go:133
		// _ = "end of CoverTab[7136]"
//line /usr/local/go/src/crypto/elliptic/params.go:133
	}
//line /usr/local/go/src/crypto/elliptic/params.go:133
	// _ = "end of CoverTab[7129]"
//line /usr/local/go/src/crypto/elliptic/params.go:133
	_go_fuzz_dep_.CoverTab[7130]++
							if z2.Sign() == 0 {
//line /usr/local/go/src/crypto/elliptic/params.go:134
		_go_fuzz_dep_.CoverTab[7137]++
								x3.Set(x1)
								y3.Set(y1)
								z3.Set(z1)
								return x3, y3, z3
//line /usr/local/go/src/crypto/elliptic/params.go:138
		// _ = "end of CoverTab[7137]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:139
		_go_fuzz_dep_.CoverTab[7138]++
//line /usr/local/go/src/crypto/elliptic/params.go:139
		// _ = "end of CoverTab[7138]"
//line /usr/local/go/src/crypto/elliptic/params.go:139
	}
//line /usr/local/go/src/crypto/elliptic/params.go:139
	// _ = "end of CoverTab[7130]"
//line /usr/local/go/src/crypto/elliptic/params.go:139
	_go_fuzz_dep_.CoverTab[7131]++

							z1z1 := new(big.Int).Mul(z1, z1)
							z1z1.Mod(z1z1, curve.P)
							z2z2 := new(big.Int).Mul(z2, z2)
							z2z2.Mod(z2z2, curve.P)

							u1 := new(big.Int).Mul(x1, z2z2)
							u1.Mod(u1, curve.P)
							u2 := new(big.Int).Mul(x2, z1z1)
							u2.Mod(u2, curve.P)
							h := new(big.Int).Sub(u2, u1)
							xEqual := h.Sign() == 0
							if h.Sign() == -1 {
//line /usr/local/go/src/crypto/elliptic/params.go:152
		_go_fuzz_dep_.CoverTab[7139]++
								h.Add(h, curve.P)
//line /usr/local/go/src/crypto/elliptic/params.go:153
		// _ = "end of CoverTab[7139]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:154
		_go_fuzz_dep_.CoverTab[7140]++
//line /usr/local/go/src/crypto/elliptic/params.go:154
		// _ = "end of CoverTab[7140]"
//line /usr/local/go/src/crypto/elliptic/params.go:154
	}
//line /usr/local/go/src/crypto/elliptic/params.go:154
	// _ = "end of CoverTab[7131]"
//line /usr/local/go/src/crypto/elliptic/params.go:154
	_go_fuzz_dep_.CoverTab[7132]++
							i := new(big.Int).Lsh(h, 1)
							i.Mul(i, i)
							j := new(big.Int).Mul(h, i)

							s1 := new(big.Int).Mul(y1, z2)
							s1.Mul(s1, z2z2)
							s1.Mod(s1, curve.P)
							s2 := new(big.Int).Mul(y2, z1)
							s2.Mul(s2, z1z1)
							s2.Mod(s2, curve.P)
							r := new(big.Int).Sub(s2, s1)
							if r.Sign() == -1 {
//line /usr/local/go/src/crypto/elliptic/params.go:166
		_go_fuzz_dep_.CoverTab[7141]++
								r.Add(r, curve.P)
//line /usr/local/go/src/crypto/elliptic/params.go:167
		// _ = "end of CoverTab[7141]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:168
		_go_fuzz_dep_.CoverTab[7142]++
//line /usr/local/go/src/crypto/elliptic/params.go:168
		// _ = "end of CoverTab[7142]"
//line /usr/local/go/src/crypto/elliptic/params.go:168
	}
//line /usr/local/go/src/crypto/elliptic/params.go:168
	// _ = "end of CoverTab[7132]"
//line /usr/local/go/src/crypto/elliptic/params.go:168
	_go_fuzz_dep_.CoverTab[7133]++
							yEqual := r.Sign() == 0
							if xEqual && func() bool {
//line /usr/local/go/src/crypto/elliptic/params.go:170
		_go_fuzz_dep_.CoverTab[7143]++
//line /usr/local/go/src/crypto/elliptic/params.go:170
		return yEqual
//line /usr/local/go/src/crypto/elliptic/params.go:170
		// _ = "end of CoverTab[7143]"
//line /usr/local/go/src/crypto/elliptic/params.go:170
	}() {
//line /usr/local/go/src/crypto/elliptic/params.go:170
		_go_fuzz_dep_.CoverTab[7144]++
								return curve.doubleJacobian(x1, y1, z1)
//line /usr/local/go/src/crypto/elliptic/params.go:171
		// _ = "end of CoverTab[7144]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:172
		_go_fuzz_dep_.CoverTab[7145]++
//line /usr/local/go/src/crypto/elliptic/params.go:172
		// _ = "end of CoverTab[7145]"
//line /usr/local/go/src/crypto/elliptic/params.go:172
	}
//line /usr/local/go/src/crypto/elliptic/params.go:172
	// _ = "end of CoverTab[7133]"
//line /usr/local/go/src/crypto/elliptic/params.go:172
	_go_fuzz_dep_.CoverTab[7134]++
							r.Lsh(r, 1)
							v := new(big.Int).Mul(u1, i)

							x3.Set(r)
							x3.Mul(x3, x3)
							x3.Sub(x3, j)
							x3.Sub(x3, v)
							x3.Sub(x3, v)
							x3.Mod(x3, curve.P)

							y3.Set(r)
							v.Sub(v, x3)
							y3.Mul(y3, v)
							s1.Mul(s1, j)
							s1.Lsh(s1, 1)
							y3.Sub(y3, s1)
							y3.Mod(y3, curve.P)

							z3.Add(z1, z2)
							z3.Mul(z3, z3)
							z3.Sub(z3, z1z1)
							z3.Sub(z3, z2z2)
							z3.Mul(z3, h)
							z3.Mod(z3, curve.P)

							return x3, y3, z3
//line /usr/local/go/src/crypto/elliptic/params.go:198
	// _ = "end of CoverTab[7134]"
}

// Double implements Curve.Double.
//line /usr/local/go/src/crypto/elliptic/params.go:201
//
//line /usr/local/go/src/crypto/elliptic/params.go:201
// Note: the CurveParams methods are not guaranteed to
//line /usr/local/go/src/crypto/elliptic/params.go:201
// provide any security property. For ECDH, use the crypto/ecdh package.
//line /usr/local/go/src/crypto/elliptic/params.go:201
// For ECDSA, use the crypto/ecdsa package with a Curve value returned directly
//line /usr/local/go/src/crypto/elliptic/params.go:201
// from P224(), P256(), P384(), or P521().
//line /usr/local/go/src/crypto/elliptic/params.go:207
func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/params.go:207
	_go_fuzz_dep_.CoverTab[7146]++

//line /usr/local/go/src/crypto/elliptic/params.go:210
	if specific, ok := matchesSpecificCurve(curve); ok {
//line /usr/local/go/src/crypto/elliptic/params.go:210
		_go_fuzz_dep_.CoverTab[7148]++
								return specific.Double(x1, y1)
//line /usr/local/go/src/crypto/elliptic/params.go:211
		// _ = "end of CoverTab[7148]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:212
		_go_fuzz_dep_.CoverTab[7149]++
//line /usr/local/go/src/crypto/elliptic/params.go:212
		// _ = "end of CoverTab[7149]"
//line /usr/local/go/src/crypto/elliptic/params.go:212
	}
//line /usr/local/go/src/crypto/elliptic/params.go:212
	// _ = "end of CoverTab[7146]"
//line /usr/local/go/src/crypto/elliptic/params.go:212
	_go_fuzz_dep_.CoverTab[7147]++
							panicIfNotOnCurve(curve, x1, y1)

							z1 := zForAffine(x1, y1)
							return curve.affineFromJacobian(curve.doubleJacobian(x1, y1, z1))
//line /usr/local/go/src/crypto/elliptic/params.go:216
	// _ = "end of CoverTab[7147]"
}

// doubleJacobian takes a point in Jacobian coordinates, (x, y, z), and
//line /usr/local/go/src/crypto/elliptic/params.go:219
// returns its double, also in Jacobian form.
//line /usr/local/go/src/crypto/elliptic/params.go:221
func (curve *CurveParams) doubleJacobian(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/params.go:221
	_go_fuzz_dep_.CoverTab[7150]++

							delta := new(big.Int).Mul(z, z)
							delta.Mod(delta, curve.P)
							gamma := new(big.Int).Mul(y, y)
							gamma.Mod(gamma, curve.P)
							alpha := new(big.Int).Sub(x, delta)
							if alpha.Sign() == -1 {
//line /usr/local/go/src/crypto/elliptic/params.go:228
		_go_fuzz_dep_.CoverTab[7157]++
								alpha.Add(alpha, curve.P)
//line /usr/local/go/src/crypto/elliptic/params.go:229
		// _ = "end of CoverTab[7157]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:230
		_go_fuzz_dep_.CoverTab[7158]++
//line /usr/local/go/src/crypto/elliptic/params.go:230
		// _ = "end of CoverTab[7158]"
//line /usr/local/go/src/crypto/elliptic/params.go:230
	}
//line /usr/local/go/src/crypto/elliptic/params.go:230
	// _ = "end of CoverTab[7150]"
//line /usr/local/go/src/crypto/elliptic/params.go:230
	_go_fuzz_dep_.CoverTab[7151]++
							alpha2 := new(big.Int).Add(x, delta)
							alpha.Mul(alpha, alpha2)
							alpha2.Set(alpha)
							alpha.Lsh(alpha, 1)
							alpha.Add(alpha, alpha2)

							beta := alpha2.Mul(x, gamma)

							x3 := new(big.Int).Mul(alpha, alpha)
							beta8 := new(big.Int).Lsh(beta, 3)
							beta8.Mod(beta8, curve.P)
							x3.Sub(x3, beta8)
							if x3.Sign() == -1 {
//line /usr/local/go/src/crypto/elliptic/params.go:243
		_go_fuzz_dep_.CoverTab[7159]++
								x3.Add(x3, curve.P)
//line /usr/local/go/src/crypto/elliptic/params.go:244
		// _ = "end of CoverTab[7159]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:245
		_go_fuzz_dep_.CoverTab[7160]++
//line /usr/local/go/src/crypto/elliptic/params.go:245
		// _ = "end of CoverTab[7160]"
//line /usr/local/go/src/crypto/elliptic/params.go:245
	}
//line /usr/local/go/src/crypto/elliptic/params.go:245
	// _ = "end of CoverTab[7151]"
//line /usr/local/go/src/crypto/elliptic/params.go:245
	_go_fuzz_dep_.CoverTab[7152]++
							x3.Mod(x3, curve.P)

							z3 := new(big.Int).Add(y, z)
							z3.Mul(z3, z3)
							z3.Sub(z3, gamma)
							if z3.Sign() == -1 {
//line /usr/local/go/src/crypto/elliptic/params.go:251
		_go_fuzz_dep_.CoverTab[7161]++
								z3.Add(z3, curve.P)
//line /usr/local/go/src/crypto/elliptic/params.go:252
		// _ = "end of CoverTab[7161]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:253
		_go_fuzz_dep_.CoverTab[7162]++
//line /usr/local/go/src/crypto/elliptic/params.go:253
		// _ = "end of CoverTab[7162]"
//line /usr/local/go/src/crypto/elliptic/params.go:253
	}
//line /usr/local/go/src/crypto/elliptic/params.go:253
	// _ = "end of CoverTab[7152]"
//line /usr/local/go/src/crypto/elliptic/params.go:253
	_go_fuzz_dep_.CoverTab[7153]++
							z3.Sub(z3, delta)
							if z3.Sign() == -1 {
//line /usr/local/go/src/crypto/elliptic/params.go:255
		_go_fuzz_dep_.CoverTab[7163]++
								z3.Add(z3, curve.P)
//line /usr/local/go/src/crypto/elliptic/params.go:256
		// _ = "end of CoverTab[7163]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:257
		_go_fuzz_dep_.CoverTab[7164]++
//line /usr/local/go/src/crypto/elliptic/params.go:257
		// _ = "end of CoverTab[7164]"
//line /usr/local/go/src/crypto/elliptic/params.go:257
	}
//line /usr/local/go/src/crypto/elliptic/params.go:257
	// _ = "end of CoverTab[7153]"
//line /usr/local/go/src/crypto/elliptic/params.go:257
	_go_fuzz_dep_.CoverTab[7154]++
							z3.Mod(z3, curve.P)

							beta.Lsh(beta, 2)
							beta.Sub(beta, x3)
							if beta.Sign() == -1 {
//line /usr/local/go/src/crypto/elliptic/params.go:262
		_go_fuzz_dep_.CoverTab[7165]++
								beta.Add(beta, curve.P)
//line /usr/local/go/src/crypto/elliptic/params.go:263
		// _ = "end of CoverTab[7165]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:264
		_go_fuzz_dep_.CoverTab[7166]++
//line /usr/local/go/src/crypto/elliptic/params.go:264
		// _ = "end of CoverTab[7166]"
//line /usr/local/go/src/crypto/elliptic/params.go:264
	}
//line /usr/local/go/src/crypto/elliptic/params.go:264
	// _ = "end of CoverTab[7154]"
//line /usr/local/go/src/crypto/elliptic/params.go:264
	_go_fuzz_dep_.CoverTab[7155]++
							y3 := alpha.Mul(alpha, beta)

							gamma.Mul(gamma, gamma)
							gamma.Lsh(gamma, 3)
							gamma.Mod(gamma, curve.P)

							y3.Sub(y3, gamma)
							if y3.Sign() == -1 {
//line /usr/local/go/src/crypto/elliptic/params.go:272
		_go_fuzz_dep_.CoverTab[7167]++
								y3.Add(y3, curve.P)
//line /usr/local/go/src/crypto/elliptic/params.go:273
		// _ = "end of CoverTab[7167]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:274
		_go_fuzz_dep_.CoverTab[7168]++
//line /usr/local/go/src/crypto/elliptic/params.go:274
		// _ = "end of CoverTab[7168]"
//line /usr/local/go/src/crypto/elliptic/params.go:274
	}
//line /usr/local/go/src/crypto/elliptic/params.go:274
	// _ = "end of CoverTab[7155]"
//line /usr/local/go/src/crypto/elliptic/params.go:274
	_go_fuzz_dep_.CoverTab[7156]++
							y3.Mod(y3, curve.P)

							return x3, y3, z3
//line /usr/local/go/src/crypto/elliptic/params.go:277
	// _ = "end of CoverTab[7156]"
}

// ScalarMult implements Curve.ScalarMult.
//line /usr/local/go/src/crypto/elliptic/params.go:280
//
//line /usr/local/go/src/crypto/elliptic/params.go:280
// Note: the CurveParams methods are not guaranteed to
//line /usr/local/go/src/crypto/elliptic/params.go:280
// provide any security property. For ECDH, use the crypto/ecdh package.
//line /usr/local/go/src/crypto/elliptic/params.go:280
// For ECDSA, use the crypto/ecdsa package with a Curve value returned directly
//line /usr/local/go/src/crypto/elliptic/params.go:280
// from P224(), P256(), P384(), or P521().
//line /usr/local/go/src/crypto/elliptic/params.go:286
func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/params.go:286
	_go_fuzz_dep_.CoverTab[7169]++

//line /usr/local/go/src/crypto/elliptic/params.go:289
	if specific, ok := matchesSpecificCurve(curve); ok {
//line /usr/local/go/src/crypto/elliptic/params.go:289
		_go_fuzz_dep_.CoverTab[7172]++
								return specific.ScalarMult(Bx, By, k)
//line /usr/local/go/src/crypto/elliptic/params.go:290
		// _ = "end of CoverTab[7172]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:291
		_go_fuzz_dep_.CoverTab[7173]++
//line /usr/local/go/src/crypto/elliptic/params.go:291
		// _ = "end of CoverTab[7173]"
//line /usr/local/go/src/crypto/elliptic/params.go:291
	}
//line /usr/local/go/src/crypto/elliptic/params.go:291
	// _ = "end of CoverTab[7169]"
//line /usr/local/go/src/crypto/elliptic/params.go:291
	_go_fuzz_dep_.CoverTab[7170]++
							panicIfNotOnCurve(curve, Bx, By)

							Bz := new(big.Int).SetInt64(1)
							x, y, z := new(big.Int), new(big.Int), new(big.Int)

							for _, byte := range k {
//line /usr/local/go/src/crypto/elliptic/params.go:297
		_go_fuzz_dep_.CoverTab[7174]++
								for bitNum := 0; bitNum < 8; bitNum++ {
//line /usr/local/go/src/crypto/elliptic/params.go:298
			_go_fuzz_dep_.CoverTab[7175]++
									x, y, z = curve.doubleJacobian(x, y, z)
									if byte&0x80 == 0x80 {
//line /usr/local/go/src/crypto/elliptic/params.go:300
				_go_fuzz_dep_.CoverTab[7177]++
										x, y, z = curve.addJacobian(Bx, By, Bz, x, y, z)
//line /usr/local/go/src/crypto/elliptic/params.go:301
				// _ = "end of CoverTab[7177]"
			} else {
//line /usr/local/go/src/crypto/elliptic/params.go:302
				_go_fuzz_dep_.CoverTab[7178]++
//line /usr/local/go/src/crypto/elliptic/params.go:302
				// _ = "end of CoverTab[7178]"
//line /usr/local/go/src/crypto/elliptic/params.go:302
			}
//line /usr/local/go/src/crypto/elliptic/params.go:302
			// _ = "end of CoverTab[7175]"
//line /usr/local/go/src/crypto/elliptic/params.go:302
			_go_fuzz_dep_.CoverTab[7176]++
									byte <<= 1
//line /usr/local/go/src/crypto/elliptic/params.go:303
			// _ = "end of CoverTab[7176]"
		}
//line /usr/local/go/src/crypto/elliptic/params.go:304
		// _ = "end of CoverTab[7174]"
	}
//line /usr/local/go/src/crypto/elliptic/params.go:305
	// _ = "end of CoverTab[7170]"
//line /usr/local/go/src/crypto/elliptic/params.go:305
	_go_fuzz_dep_.CoverTab[7171]++

							return curve.affineFromJacobian(x, y, z)
//line /usr/local/go/src/crypto/elliptic/params.go:307
	// _ = "end of CoverTab[7171]"
}

// ScalarBaseMult implements Curve.ScalarBaseMult.
//line /usr/local/go/src/crypto/elliptic/params.go:310
//
//line /usr/local/go/src/crypto/elliptic/params.go:310
// Note: the CurveParams methods are not guaranteed to
//line /usr/local/go/src/crypto/elliptic/params.go:310
// provide any security property. For ECDH, use the crypto/ecdh package.
//line /usr/local/go/src/crypto/elliptic/params.go:310
// For ECDSA, use the crypto/ecdsa package with a Curve value returned directly
//line /usr/local/go/src/crypto/elliptic/params.go:310
// from P224(), P256(), P384(), or P521().
//line /usr/local/go/src/crypto/elliptic/params.go:316
func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int) {
//line /usr/local/go/src/crypto/elliptic/params.go:316
	_go_fuzz_dep_.CoverTab[7179]++

//line /usr/local/go/src/crypto/elliptic/params.go:319
	if specific, ok := matchesSpecificCurve(curve); ok {
//line /usr/local/go/src/crypto/elliptic/params.go:319
		_go_fuzz_dep_.CoverTab[7181]++
								return specific.ScalarBaseMult(k)
//line /usr/local/go/src/crypto/elliptic/params.go:320
		// _ = "end of CoverTab[7181]"
	} else {
//line /usr/local/go/src/crypto/elliptic/params.go:321
		_go_fuzz_dep_.CoverTab[7182]++
//line /usr/local/go/src/crypto/elliptic/params.go:321
		// _ = "end of CoverTab[7182]"
//line /usr/local/go/src/crypto/elliptic/params.go:321
	}
//line /usr/local/go/src/crypto/elliptic/params.go:321
	// _ = "end of CoverTab[7179]"
//line /usr/local/go/src/crypto/elliptic/params.go:321
	_go_fuzz_dep_.CoverTab[7180]++

							return curve.ScalarMult(curve.Gx, curve.Gy, k)
//line /usr/local/go/src/crypto/elliptic/params.go:323
	// _ = "end of CoverTab[7180]"
}

func matchesSpecificCurve(params *CurveParams) (Curve, bool) {
//line /usr/local/go/src/crypto/elliptic/params.go:326
	_go_fuzz_dep_.CoverTab[7183]++
							for _, c := range []Curve{p224, p256, p384, p521} {
//line /usr/local/go/src/crypto/elliptic/params.go:327
		_go_fuzz_dep_.CoverTab[7185]++
								if params == c.Params() {
//line /usr/local/go/src/crypto/elliptic/params.go:328
			_go_fuzz_dep_.CoverTab[7186]++
									return c, true
//line /usr/local/go/src/crypto/elliptic/params.go:329
			// _ = "end of CoverTab[7186]"
		} else {
//line /usr/local/go/src/crypto/elliptic/params.go:330
			_go_fuzz_dep_.CoverTab[7187]++
//line /usr/local/go/src/crypto/elliptic/params.go:330
			// _ = "end of CoverTab[7187]"
//line /usr/local/go/src/crypto/elliptic/params.go:330
		}
//line /usr/local/go/src/crypto/elliptic/params.go:330
		// _ = "end of CoverTab[7185]"
	}
//line /usr/local/go/src/crypto/elliptic/params.go:331
	// _ = "end of CoverTab[7183]"
//line /usr/local/go/src/crypto/elliptic/params.go:331
	_go_fuzz_dep_.CoverTab[7184]++
							return nil, false
//line /usr/local/go/src/crypto/elliptic/params.go:332
	// _ = "end of CoverTab[7184]"
}

//line /usr/local/go/src/crypto/elliptic/params.go:333
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/elliptic/params.go:333
var _ = _go_fuzz_dep_.CoverTab
